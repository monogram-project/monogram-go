;;; A prototype of monogram in Pop-11.

compile_mode :pop11 +strict;

;;; Tokenisation:                               Has precedence?     Classify
;;;     0. Comments - # ...                     N/A                 N/A
;;;     1. Literals - strings, numbers          false               false
;;;     2. Separators - comma, semi-colon       false               "sep"
;;;     3. Open Brackets - ( { [                true                "open"
;;;     4. Close Brackets - ] } )               false               "close"
;;;     5. End Form - endXXX                    false               "end"
;;;     6. Keywords - foo: foo?                 false               "keyword"
;;;     7. Signs - + / %                        true                "sign"
;;;     8. Identifiers                          false               "id"

vars glue_chartype = false;

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
        lvars n = item_chartype(ch);
        nextif( n == 3 or n == 10 or n == 11 or n == glue_chartype );
        return( false );
    endfor;
    return( true );
enddefine;

define is_form_end( word );
    isstartstring( 'end', word )
enddefine;

define is_form_keyword( word );
    lvars L = datalength( word );
    returnif( L == 0 )( false );
    lvars ch = subscrw( L, word );
    ch == `:` or ch == `?`
enddefine;

define is_identifier( word );
    returnif( word.is_sign )( false );
    returnif( word.is_open_bracket or word.is_close_bracket )( false );
    returnif( word.is_form_end )( false );
    true
enddefine;

define classify_item( item );
    returnunless( item.isword )( false );
    lvars L = datalength( item );
    returnif( L == 0 )( false );
    if L == 1 then
        lvars ch_first = subscrw( 1, item );
        returnif( locchar( ch_first, 1, '({[' ) )( "open" );
        returnif( locchar( ch_first, 1, ']})' ) )( "close" );
        returnif( ch_first == `,` or ch_first == `;` )( "sep" );
    endif;
    returnif( item.is_form_keyword )( "keyword" );
    returnif( item.is_form_end )( "end" );
    returnif( item.is_sign )( "sign" );
    "id"
enddefine;

define is_form_opening( next_item );
    lvars tokentype = classify_item( next_item );
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

;;;     0. Comments - # ...                     N/A                 N/A
;;;     1. Literals - strings, numbers          false               false
;;;     2. Separators - comma, semi-colon       false               "sep"
;;;     3. Open Brackets - ( { [                true                "open"
;;;     4. Close Brackets - ] } )               false               "close"
;;;     5. End Form - endXXX                    false               "end"
;;;     6. Keywords - foo: foo?                 false               "keyword"
;;;     7. Signs - + / %                        true                "sign"
;;;     8. Identifiers                          false               "id"
define read_form_expr(opening_word);
    ;;; [read_form_expr ^opening_word] =>
    lvars closing_keywords = [% "end", "end" <> opening_word %];
    lvars current_part = [];
    lvars current_keyword = opening_word;
    lvars contents = [%
        until pop11_try_nextreaditem( closing_keywords ) do
            ;;; [current_part ^current_part] =>
            if proglist.null then
                mishap( 'Unexpected end of file while reading form', [^opening_word])
            else
                lvars item1 = proglist.hd;
                ;;; [peek ^item1] =>
                lvars tokentype1 = classify_item( item1 );
                ;;; [tt ^tokentype1] =>
                if tokentype1 == "sep" or tokentype1 == "sign" or tokentype1 == "close" then
                    mishap( 'Unexpected item at start of expression', [^item1] )
                elseif tokentype1 == "end" then
                    ;;; [closing_keywords ^closing_keywords] =>
                    mishap( 'Mismatched closing keyword', [^item1] )
                elseif tokentype1 == "keyword" then
                    ;;; [keyword ^item1] =>
                    current_keyword :: current_part;
                    [] -> current_part;
                    item1 -> current_keyword;
                    proglist.tl -> proglist;
                else
                    current_part <> [% read_expr() %] -> current_part;
                endif
            endif
        enduntil;
        current_keyword :: current_part;
    %];
    [form ^^contents]
enddefine;

define read_primary_expr();
    lvars item = readitem();
    lvars tokentype = classify_item( item );
    returnunless( tokentype )( [constant ^item] );
    if tokentype == "open" then
        lvars expr = read_expr();
        pop11_need_nextreaditem( item.is_open_bracket ) -> _;        
        lvars dname = delimiter_name( item );
        [delimited ^dname ^expr]
    elseif tokentype == "id" then
        ;;; The interpretation depends on what comes next.
        if null(proglist) then
            [identifier ^item]
        else
            lvars item1 = proglist.hd;
            if item1.is_form_opening then
                ;;; [form opening ^item] =>
                read_form_expr( item )
            else
                [identifier ^item]
            endif
        endif
    else
        mishap( 'Unexpected token at start of expression', [^item] )
    endif;
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
            if ch == `:` then
                lvars item1 = itemiser();
                ;;; [peekitem ^item1] =>
                if item1 == ":" then
                    item <> ":"
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

define monogram(procedure source);

    ;;; Might not need this.
    unless glue_chartype do
        item_newtype() -> glue_chartype
    endunless;

    lvars procedure itemiser = incharitem(source);
    5 -> item_chartype( `;`, itemiser );
    9 -> item_chartype( `#`, itemiser );
    
    dlocal proglist = pdtolist(glue(itemiser));
    read_expr()
enddefine;
