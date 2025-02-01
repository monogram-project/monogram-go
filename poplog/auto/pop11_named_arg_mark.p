;;;
;;; This file contains the essential shared-core for named-argument
;;; groups. It is typically loaded with
;;;     uses pop11_named_arg_mark
;;;

compile_mode :pop11 +strict;

section $-gospl$-named_args =>
    pop11_named_arg_mark
    ;

#_IF not(DEF pop11_named_arg_mark)
;;; We do not want this re-executed, so we protect it behind this #_IF check.
constant pop11_named_arg_mark = conskey( "pop11_named_arg_mark", [] ).class_cons.apply;
#_ENDIF

;;; Syntactic separators.
constant
    key_value_separator = "=",
    rename_separator = "/"
;

;;; Variable/Keyword pairs
constant procedure(
    new_vk = conspair,
    vk_variable = front,
    vk_keyword = back
);

constant procedure
    ascending = alphabefore,
    descending = alphabefore <> not;

endsection;
