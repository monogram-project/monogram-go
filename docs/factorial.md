# Recursive factorial function

## Monogram

```txt
# The factorial function in monogram.
def f(n):
    if n <= 1:
        1
    else:
        n * f(n - 1)
    endif
enddef
```

## Mermaid diagram

We can target Mermaid's flowchart as an output format. 
And this is what it looks like:

```mermaid
graph LR
  134043923466896["form: surround"]:::custom_form;
  134043923466976["part: def"]:::custom_part;
  134043923466896 --> 134043923466976;
  134043923467056["apply"]:::custom_apply;
  134043923466976 --> 134043923467056;
  134043923467136["identifier: f"]:::custom_identifier;
  134043923467056 --> 134043923467136;
  134043923467216["arguments"]:::custom_arguments;
  134043923467056 --> 134043923467216;
  134043923467296["identifier: n"]:::custom_identifier;
  134043923467216 --> 134043923467296;
  134043923467376["part: _"]:::custom_part;
  134043923466896 --> 134043923467376;
  134043923467456["form: surround"]:::custom_form;
  134043923467376 --> 134043923467456;
  134043923467536["part: if"]:::custom_part;
  134043923467456 --> 134043923467536;
  134043923467616["operator: <="]:::custom_operator;
  134043923467536 --> 134043923467616;
  134043923467696["identifier: n"]:::custom_identifier;
  134043923467616 --> 134043923467696;
  134043923467776["number: 1"]:::custom_number;
  134043923467616 --> 134043923467776;
  134043923467856["part: _"]:::custom_part;
  134043923467456 --> 134043923467856;
  134043923467936["number: 1"]:::custom_number;
  134043923467856 --> 134043923467936;
  134043923468016["part: else"]:::custom_part;
  134043923467456 --> 134043923468016;
  134043923468096["operator: *"]:::custom_operator;
  134043923468016 --> 134043923468096;
  134043923468176["identifier: n"]:::custom_identifier;
  134043923468096 --> 134043923468176;
  134043923468256["apply"]:::custom_apply;
  134043923468096 --> 134043923468256;
  134043923468416["identifier: f"]:::custom_identifier;
  134043923468256 --> 134043923468416;
  134043923468576["arguments"]:::custom_arguments;
  134043923468256 --> 134043923468576;
  134043923468736["operator: -"]:::custom_operator;
  134043923468576 --> 134043923468736;
  134043923468896["identifier: n"]:::custom_identifier;
  134043923468736 --> 134043923468896;
  134043923469056["number: 1"]:::custom_number;
  134043923468736 --> 134043923469056;

classDef custom_form fill:lightpink,stroke:#333,stroke-width:2px;
classDef custom_part fill:#FFD8E1,stroke:#333,stroke-width:2px;
classDef custom_apply fill:lightgreen,stroke:#333,stroke-width:2px;
classDef custom_identifier fill:Honeydew,stroke:#333,stroke-width:2px;
classDef custom_arguments fill:PaleTurquoise,stroke:#333,stroke-width:2px;
classDef custom_operator fill:#C0FFC0,stroke:#333,stroke-width:2px;
classDef custom_number fill:lightgoldenrodyellow,stroke:#333,stroke-width:2px;

```

## XML

```xml
<form syntax="surround">
    <part keyword="def">
        <apply kind="parentheses" separator="undefined">
            <identifier name="f"/>
            <arguments>
                <identifier name="n"/>
            </arguments>
        </apply>
    </part>
    <part keyword="_">
        <form syntax="surround">
            <part keyword="if">
                <operator name="&lt;=" syntax="infix">
                    <identifier name="n"/>
                    <number value="1"/>
                </operator>
            </part>
            <part keyword="_">
                <number value="1"/>
            </part>
            <part keyword="else">
                <operator name="*" syntax="infix">
                    <identifier name="n"/>
                    <apply kind="parentheses" separator="undefined">
                        <identifier name="f"/>
                        <arguments>
                            <operator name="-" syntax="infix">
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

## JSON

We can target JSON as an output format. The format of each node is
a bit verbose but straightforward:

```json
{
    "role": "{{NODE NAME}}",
    "ATTRIBUTE_1": "VALUE_1", 
    ... 
    "ATTRIBUTE_N": "VALUE_N",
    "children": [ 
        ...
    ]
}
```

And this is what it expands into:

```json
{
    "role": "form",
    "syntax": "surround",
    "children": [
        {
            "role": "part",
            "keyword": "def",
            "children": [
                {
                    "role": "apply",
                    "kind": "parentheses",
                    "separator": "undefined",
                    "children": [
                        {
                            "role": "identifier",
                            "name": "f"
                        },
                        {
                            "role": "arguments",
                            "children": [
                                {
                                    "role": "identifier",
                                    "name": "n"
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "role": "part",
            "keyword": "_",
            "children": [
                {
                    "role": "form",
                    "syntax": "surround",
                    "children": [
                        {
                            "role": "part",
                            "keyword": "if",
                            "children": [
                                {
                                    "role": "operator",
                                    "name": "<=",
                                    "syntax": "infix",
                                    "children": [
                                        {
                                            "role": "identifier",
                                            "name": "n"
                                        },
                                        {
                                            "role": "number",
                                            "value": 1
                                        }
                                    ]
                                }
                            ]
                        },
                        {
                            "role": "part",
                            "keyword": "_",
                            "children": [
                                {
                                    "role": "number",
                                    "value": 1
                                }
                            ]
                        },
                        {
                            "role": "part",
                            "keyword": "else",
                            "children": [
                                {
                                    "role": "operator",
                                    "name": "*",
                                    "syntax": "infix",
                                    "children": [
                                        {
                                            "role": "identifier",
                                            "name": "n"
                                        },
                                        {
                                            "role": "apply",
                                            "kind": "parentheses",
                                            "separator": "undefined",
                                            "children": [
                                                {
                                                    "role": "identifier",
                                                    "name": "f"
                                                },
                                                {
                                                    "role": "arguments",
                                                    "children": [
                                                        {
                                                            "role": "operator",
                                                            "name": "-",
                                                            "syntax": "infix",
                                                            "children": [
                                                                {
                                                                    "role": "identifier",
                                                                    "name": "n"
                                                                },
                                                                {
                                                                    "role": "number",
                                                                    "value": 1
                                                                }
                                                            ]
                                                        }
                                                    ]
                                                }
                                            ]
                                        }
                                    ]
                                }
                            ]
                        }
                    ]
                }
            ]
        }
    ]
}
```

## YAML

We can target YAML as an output format. The format of each node is
a bit verbose but easy to understand.

```yaml
role": "{{NODE NAME}}",
ATTRIBUTE_1: VALUE_1,
...
children": 
- ...
- ...
```

And this is what it expands into:



```yaml
role: form
syntax: surround
children:
- role: part
  keyword: def
  children:
  - role: apply
    kind: parentheses
    separator: undefined
    children:
    - role: identifier
      name: f
    - role: arguments
      children:
      - role: identifier
        name: n
- role: part
  keyword: _
  children:
  - role: form
    syntax: surround
    children:
    - role: part
      keyword: if
      children:
      - role: operator
        name: <=
        syntax: infix
        children:
        - role: identifier
          name: n
        - role: number
          value: 1
    - role: part
      keyword: _
      children:
      - role: number
        value: 1
    - role: part
      keyword: else
      children:
      - role: operator
        name: '*'
        syntax: infix
        children:
        - role: identifier
          name: n
        - role: apply
          kind: parentheses
          separator: undefined
          children:
          - role: identifier
            name: f
          - role: arguments
            children:
            - role: operator
              name: '-'
              syntax: infix
              children:
              - role: identifier
                name: n
              - role: number
                value: 1

```

## Graphviz Dot format

```dot
digraph G {
  bgcolor="transparent";
  node [shape="box", style="filled", fontname="Ubuntu Mono"];
  "128217096981056" [label="form: surround", shape="box", fillcolor="lightpink"];
  "128217096981136" [label="part: def", shape="box", fillcolor="#FFD8E1"];
  "128217096981056" -> "128217096981136";
  "128217096981216" [label="apply", shape="box", fillcolor="lightgreen"];
  "128217096981136" -> "128217096981216";
  "128217096981296" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "128217096981216" -> "128217096981296";
  "128217096981376" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "128217096981216" -> "128217096981376";
  "128217096981456" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "128217096981376" -> "128217096981456";
  "128217096981536" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "128217096981056" -> "128217096981536";
  "128217096981616" [label="form: surround", shape="box", fillcolor="lightpink"];
  "128217096981536" -> "128217096981616";
  "128217096981696" [label="part: if", shape="box", fillcolor="#FFD8E1"];
  "128217096981616" -> "128217096981696";
  "128217096981776" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "128217096981696" -> "128217096981776";
  "128217096981856" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "128217096981776" -> "128217096981856";
  "128217096981936" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "128217096981776" -> "128217096981936";
  "128217096982016" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "128217096981616" -> "128217096982016";
  "128217096982096" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "128217096982016" -> "128217096982096";
  "128217096982176" [label="part: else", shape="box", fillcolor="#FFD8E1"];
  "128217096981616" -> "128217096982176";
  "128217096982256" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "128217096982176" -> "128217096982256";
  "128217096982336" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "128217096982256" -> "128217096982336";
  "128217096982416" [label="apply", shape="box", fillcolor="lightgreen"];
  "128217096982256" -> "128217096982416";
  "128217096982576" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "128217096982416" -> "128217096982576";
  "128217096982736" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "128217096982416" -> "128217096982736";
  "128217096982896" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "128217096982736" -> "128217096982896";
  "128217096983056" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "128217096982896" -> "128217096983056";
  "128217096983216" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "128217096982896" -> "128217096983216";
}
```

![Image generated for example](images/factorial.png)
