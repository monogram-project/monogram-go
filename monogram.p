;;; A prototype of monogram in Pop-11.

compile_mode :pop11 +strict;

;;; Tokenisation:                               Has precedence?     Classify
;;;     0. Comments - # ...                     N/A                 N/A
;;;     1. Literals - strings, numbers          false               false
;;;     2. Separators - comma, semi-colon       false               "sep"
;;;     3. Open Brackets - ( { [                true                "open"
;;;     4. Close Brackets - ] } )               false               "close"
;;;     5. Start Form - XXX                     false               "start"
;;;     6. End Form - endXXX                    false               "end"
;;;     7. Keywords - foo: foo?                 false               "keyword"
;;;     8. Signs - + / %                        true                "sign"
;;;     9. Identifiers                          false               "id"

vars unglue_option = false;
vars optional_statement_separator_option = false;
vars inferred_form_starts = [];

define peek_item();
    not( null( proglist ) ) and proglist.hd
enddefine;

define peek_nth_item( n );
    lvars PL = proglist;
    while n > 1 do
        returnif( null( PL ) )( false );
        PL.tl -> PL;
        n - 1 -> n;
    endwhile;
    not( null( PL ) ) and PL.hd    
enddefine;

define is_open_bracket( word );
    returnif( word == "(" )( ")" );
    returnif( word == "{" )( "}" );
    returnif( word == "[" )( "]" );
    return( false )
enddefine;

define is_close_bracket( word );
    returnif( word == ")" )( true );
    returnif( word == "}" )( true );
    returnif( word == "]" )( true );
    return( false )
enddefine;

define delimiter_name( word );
    returnif( word == "(" )( "parentheses" );
    returnif( word == "{" )( "braces" );
    returnif( word == "[" )( "brackets" );
    return( false )
enddefine;

define is_sign( word );
    lvars ch;
    for ch in_vectorclass word.fast_word_string do
        nextif( ch == `.` );
        lvars n = item_chartype(ch);
        nextif( n == 3 or n == 10 or n == 11 );
        return( false );
    endfor;
    return( word /== "!" );
enddefine;

define is_form_end( word );
    isstartstring( 'end', word )
enddefine;

define classify_item( item, next_item );
    ;;; [classify_item ^item ^next_item] =>
    returnunless( item.isword )( false );
    lvars L = datalength( item );
    returnif( L == 0 )( false );
    if L == 1 then
        lvars ch_first = subscrw( 1, item );
        returnif( locchar( ch_first, 1, '({[' ) )( "open" );
        returnif( locchar( ch_first, 1, ']})' ) )( "close" );
        returnif( ch_first == `,` or ch_first == `;` )( "sep" );
    endif;
    returnif( item.is_form_end )( "end" );
    returnif( item.is_sign )( "sign" );
    if fast_lmember( item, inferred_form_starts ) then
        "start"
    elseif next_item == ":" or next_item == "?" then
        "keyword"
    else
        "id"
    endif
enddefine;

define is_form_opening( next_item, item_after );
    lvars tokentype = classify_item( next_item, item_after );
    ;;; [is_form_opening ^next_item ^tokentype] =>
    not( tokentype ) or tokentype == "id"
enddefine;

;;; Precedence rules.
;;; 10 . ( [ {
;;; 20 * / %
;;; 29 ++ --
;;; 30 + -
;;; 39 << >>
;;; 40 < <=	> >=
;;; 49 &&
;;; 50 &
;;; 59 ||
;;; 50 |
;;; 59 == =/=
;;; 60 =

constant max_precedence = 999;

define precedence( item );
    returnunless( item.isword )( false );
    returnunless( item.is_sign or item.is_open_bracket )( false );
    returnif( item == ":" )( false );
    lvars n = datalength( item );
    if n > 0 then
        lvars ch = subscrw( 1, item );
        lvars L = locchar( ch, 1, '.?({[*/%+-<&|=' );
        if L then
            lvars prec = 10 * L;
            if n >= 2 and locchar(ch, 2, item ) then
                prec - 1 -> prec
            endif;
            prec
        else
            max_precedence
        endif
    else
        max_precedence
    endif
enddefine;


vars procedure read_expr, read_expr_prec;

define read_form_expr(opening_word);
    ;;; [read_form_expr ^opening_word] =>
    lvars closing_keywords = [% "end", "end" <> opening_word %];
    lvars current_part = [];
    lvars current_keyword = opening_word;
    lvars contents = [%
        lvars first_expr = true;
        until pop11_try_nextreaditem( closing_keywords ) do
            ;;; [current_part ^current_part] =>
            if proglist.null then
                mishap( 'Unexpected end of file while reading form', [^opening_word])
            else
                lvars item1 = proglist.hd;
                lvars tokentype1 = classify_item( item1, peek_nth_item(2) );
                ;;; [peek ^item1 ^tokentype1] =>
                if item1 == ":" and unglue_option then
                    unglue_option :: proglist -> proglist;
                    unglue_option -> item1;
                    "keyword" -> tokentype1;
                endif;
                if tokentype1 == "sep" or tokentype1 == "sign" or tokentype1 == "close" then
                    mishap( 'Unexpected item at start of expression (in ' >< opening_word >< ')', [^item1] )
                elseif tokentype1 == "end" then
                    ;;; [closing_keywords ^closing_keywords] =>
                    mishap( 'Mismatched closing keyword', [^item1] )
                elseif tokentype1 == "keyword" then
                    ;;; [keyword ^item1] =>
                    [part ^current_keyword ^^current_part];
                    [] -> current_part;
                    item1 -> current_keyword;
                    ;;; Skip the `:` or `?`.
                    proglist.tl.tl -> proglist;
                    true -> first_expr;
                else
                    if not( first_expr ) and not( optional_statement_separator_option ) then
                        mishap( 'Semi-colon expected', [^item1] )
                    endif;
                    current_part <> [% read_expr() %] -> current_part;
                    pop11_try_nextreaditem( ";" ) -> first_expr;
                endif
            endif
        enduntil;
        [part ^current_keyword ^^current_part];
    %];
    [form ^^contents]
enddefine;

define read_primary_expr();
    lvars item = readitem();
    ;;; [show ^item] =>
    if item == "!" then
        lvars item1 = readitem();
        if item1.isword then
            [form [part ^item1]]
        else
            mishap( 'Identifier required following `!`', [^item] )
        endif
    else
        lvars tokentype = classify_item( item, peek_item() );
        returnunless( tokentype )( [constant ^item] );
        if tokentype == "keyword" and unglue_option then
            lvars reclassified_tokentype = classify_item( item, unglue_option );
            if reclassified_tokentype == "id" then
                reclassified_tokentype -> tokentype;
                unglue_option :: proglist -> proglist
            endif
        endif;
        if tokentype == "open" then
            lvars expr = read_expr();
            pop11_need_nextreaditem( item.is_open_bracket ) -> _;
            lvars dname = delimiter_name( item );
            [delimited ^dname ^expr]
        elseif tokentype == "start" then
            read_form_expr( item )
        elseif tokentype == "id" then
            ;;; The interpretation depends on what comes next.
            if null(proglist) then
                [identifier ^item]
            else
                lvars item1 = proglist.hd;
                if is_form_opening( item1, peek_nth_item(2) ) then
                    ;;; [form opening ^item] =>
                    read_form_expr( item )
                else
                    [identifier ^item]
                endif
            endif
        else
            ;;; [mishap ^item ^tokentype] =>
            mishap( 'Unexpected token at start of expression', [^item] )
        endif
    endif
enddefine;

define read_arguments( close_bracket );
    [arguments %
        if pop11_try_nextreaditem( close_bracket ) then
            ;;; Done.
        else
            repeat
                read_expr();
                nextif( pop11_try_nextreaditem( "," ) );
                quitif( pop11_need_nextreaditem( close_bracket ) );
            endrepeat
        endif
    %]
enddefine;

define read_expr_prec(prec);
    lvars lhs = read_primary_expr();
    ;;; [lhs ^lhs] =>
    until null( proglist ) do
        lvars item1 = proglist.hd;
        lvars p = precedence( item1 );
        ;;; [prec ^item1 ^p] =>
        if p and p <= prec then
            proglist.tl -> proglist;
            lvars close_bracket = false;
            if item1.is_open_bracket ->> close_bracket then
                lvars args = read_arguments( close_bracket );
                lvars dname = delimiter_name( item1 );
                [apply ^dname ^lhs ^args] -> lhs;
            elseif item1 == "." then
                ;;; [FOUND .] =>
                lvars item2 = readitem();
                lvars tokentype2 = classify_item( item2, peek_item() );
                if tokentype2 == "id" then
                    lvars item3 = not( proglist.null ) and proglist.hd;
                    if item3.is_open_bracket ->> close_bracket then
                        proglist.tl -> proglist;
                        lvars args = read_arguments( close_bracket );
                        lvars dname = delimiter_name( item3 );
                        [invoke ^dname ^item2 ^lhs ^args] -> lhs
                    else
                        [get ^item2 ^lhs] -> lhs
                    endif
                else
                    mishap( 'Unexpected item after `.`', [^item2] )
                endif;
                ;;; [DONE .] =>
            else
                lvars rhs = read_expr_prec( p );
                [operator ^item1 ^lhs ^rhs] -> lhs;
            endif
        else
            quitloop
        endif
    enduntil;
    return( lhs )
enddefine;

define read_expr();
    read_expr_prec( max_precedence )
enddefine;

define glue( procedure itemiser );
    procedure();
        lvars item = itemiser();
        if item.isword then
            lvars ch = nextchar( itemiser );
            ch -> nextchar( itemiser );
            ;;; [peekchar ^item ^ch] =>
            if ch == `:` or ch == `?` then
                lvars item1 = itemiser();
                ;;; [peekitem ^item1] =>
                if item1 == ":" or item1 == "?" then
                    item <> item1
                else
                    item1.fast_word_string -> nextchar( itemiser );
                    item
                endif
            else
                item
            endif
        else
            item
        endif
    endprocedure
enddefine;

define infer_form_starts( dlist );
    [%
        lvars w;
        for w in dlist do
            if w.isword and is_form_end( w ) and w /== "end" then
                subword( 4, datalength( w ) - 3, w )
            endif
        endfor
    %]
enddefine;

define monogram(procedure source, unglue, opt_seps);
    dlocal unglue_option = unglue;
    dlocal optional_statement_separator_option = opt_seps;
    dlocal inferred_form_starts;

    lvars procedure itemiser = incharitem(source);
    5 -> item_chartype( `;`, itemiser );
    9 -> item_chartype( `#`, itemiser );

    dlocal proglist = pdtolist(itemiser);
    infer_form_starts( proglist ) -> inferred_form_starts;

    read_expr()
enddefine;
