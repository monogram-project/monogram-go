compile_mode :pop11 +strict;

section $-options => isoptions options_key newoptions;

defclass node {
    node_left,
    node_name,
    node_value,
    node_right
};

define add_pair_to_node( name, value, node );
    lvars nname = node.node_name;
    if name < nname then
        lvars left = node.node_left;
        if left then
            TBD
        else
            consnode( false, name, value, false ) -> node.node_left
        endif
    elseif nname < name then
        TBD
    else
        TBD
    endif
enddefine;


defclass options {
    options_nodes
};

define add_pair_to_options( name, value, options );
    if options.options_nodes then
        add_pair_to_node( name, value, options.options_nodes )
    else
        consnode( false, name, value, false ) -> options.options_nodes
    endif
enddefine;



endsection;