1 -> pop_file_versions;

extend_searchlist( 'auto', popautolist ) -> popautolist;
extend_searchlist( 'lib', popuseslist ) -> popuseslist;

load builder.p
load monogram.p

define print_as_json( tree );

    define lconstant pr_json( x );
        if x.isstring or x.isword then
            pr( '"' );
            pr( x );
            pr( '"' );
        else
            pr(x)
        endif
    enddefine;

    lvars name = tree.nodeName;
    lvars attrs = tree.nodeAttrs;
    cucharout( `{` );
    pr( '"role": "' );
    pr( name );
    pr( '"' );
    appnamedtuple(
        attrs,
        procedure( k, v );
            pr( ',' );
            pr( '"' );
            pr( k );
            pr( '"' );
            pr( ': ' );
            pr_json( v ); 
        endprocedure
    );
    unless tree.nodeChildren.null then
        pr( ',' );
        pr( '"children": [');
        lvars sep = '';
        lvars child;
        for child in tree.nodeChildren do
            pr( sep );
            print_as_json( child );
            ',' -> sep;
        endfor;
        pr( ']' );
    endunless;
    pr( '}' )
enddefine;

define :optargs print_as_xml( tree -&- indent=0 );

    define lconstant pr_html_escape( txt );
        if txt.isword then
            fast_word_string( txt ) -> txt
        elseif txt.isnumber then
            txt >< '' -> txt
        endif;
        lvars ch;
        for ch in_string txt do
            if ch == `<` then
                pr( '&lt;' )
            elseif ch == `>` then
                pr( '&gt;' )
            elseif ch == `&` then
                pr( '&amp;' )
            else
                cucharout( ch )
            endif
        endfor; 
    enddefine;

    lvars name = tree.nodeName;
    lvars attrs = tree.nodeAttrs;
    sp( indent );
    cucharout( `<` );
    pr( name );
    appnamedtuple(
        attrs,
        procedure( k, v );
            sp( 1 );
            pr( k );
            pr( '="' );
            pr_html_escape( v ); 
            cucharout( '"' )
        endprocedure
    );
    if tree.nodeChildren.null then
        npr( '/>' )
    else
        npr( '>' );
        lvars child;
        for child in tree.nodeChildren do
            print_as_xml( child -&- indent=indent+4)
        endfor;
        sp( indent );
        pr( '</' );
        pr( name );
        npr( '>' );
    endif;
enddefine;

define convert( tree, builder );
    builder( tree.nodeAttrs, "start_" <> tree.nodeName );
    lvars arg;
    for arg in tree.nodeChildren do
        convert( arg, builder )
    endfor;
    builder( "end_" <> tree.nodeName );
enddefine;

define show(source);
    lvars tree;
    for tree in monogram(source).pdtolist do
        { mg ^tree }==>
        lvars builder = new_monogram_builder();
        convert( tree, builder );
        pr( newline );
        builder( "new" ) ==>
        print_as_xml( tree );
        pr( newline );
        print_as_json( tree );
        pr( newline );
    endfor
enddefine;

/*
lvars f;
for f in ['01' '02' '03' '04' '05' '06'] do
    npr( f );
    show(discin('../ex/' >< f >< '.mg'));
endfor;
*/
