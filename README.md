# Monogram

_"It's source code, Jim. But not as we know it!"_

Monogram is a "no batteries" notation for writing domain-specific programs and
configuration files. It is easy for humans to read and write. It is easy for machines 
to parse and generate. It deliberately borrows from many programming languages 
but feels familiar to Python and Ruby programmers.

## Batteries not included

To experienced programmers, the following code looks a lot like the definition 
of the factorial function:
```py
def f(n):
    if n <= 1:
        1
    else:
        n * f(n - 1)
    endif
enddef
```

And Monogram can easily translate this into XML (shown below) or JSON:
```xml
<form>
    <part keyword="def">
        <apply kind="parentheses" separator="undefined">
            <identifier name="f"/>
            <arguments>
                <identifier name="n"/>
            </arguments>
        </apply>
    </part>
    <part keyword="_">
        <form>
            <part keyword="if">
                <operator name="&lt;=">
                    <identifier name="n"/>
                    <number value="1"/>
                </operator>
            </part>
            <part keyword="_">
                <number value="1"/>
            </part>
            <part keyword="else">
                <operator name="*">
                    <identifier name="n"/>
                    <apply kind="parentheses" separator="undefined">
                        <identifier name="f"/>
                        <arguments>
                            <operator name="-">
                                <identifier name="n"/>
                                <number value="1"/>
                            </operator>
                        </arguments>
                    </apply>
                </operator>
            </part>
        </form>
    </part>
</form>
```

But the slightly mysterious thing is that Monogram has no idea what `def` or `if`
might mean. Nor does it have a clue about `*` or `-` either. And it definitely
cannot execute this program. Because Monogram is just a notation for writing program-like 
"code" but comes without any built-in meanings.
