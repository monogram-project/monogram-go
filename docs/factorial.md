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
  137264901933712["form: surround"]:::custom_form;
  137264901933792["part: def"]:::custom_part;
  137264901933712 --> 137264901933792;
  137264901933872["apply"]:::custom_apply;
  137264901933792 --> 137264901933872;
  137264901933952["identifier: f"]:::custom_identifier;
  137264901933872 --> 137264901933952;
  137264901934032["arguments"]:::custom_arguments;
  137264901933872 --> 137264901934032;
  137264901934112["identifier: n"]:::custom_identifier;
  137264901934032 --> 137264901934112;
  137264901934192["part: _"]:::custom_part;
  137264901933712 --> 137264901934192;
  137264901934272["form: surround"]:::custom_form;
  137264901934192 --> 137264901934272;
  137264901934352["part: if"]:::custom_part;
  137264901934272 --> 137264901934352;
  137264901934432["operator: <="]:::custom_operator;
  137264901934352 --> 137264901934432;
  137264901934512["identifier: n"]:::custom_identifier;
  137264901934432 --> 137264901934512;
  137264901934592["number: 1"]:::custom_number;
  137264901934432 --> 137264901934592;
  137264901934672["part: _"]:::custom_part;
  137264901934272 --> 137264901934672;
  137264901934752["number: 1"]:::custom_number;
  137264901934672 --> 137264901934752;
  137264901934832["part: else"]:::custom_part;
  137264901934272 --> 137264901934832;
  137264901934912["operator: *"]:::custom_operator;
  137264901934832 --> 137264901934912;
  137264901934992["identifier: n"]:::custom_identifier;
  137264901934912 --> 137264901934992;
  137264901935072["apply"]:::custom_apply;
  137264901934912 --> 137264901935072;
  137264901935232["identifier: f"]:::custom_identifier;
  137264901935072 --> 137264901935232;
  137264901935392["arguments"]:::custom_arguments;
  137264901935072 --> 137264901935392;
  137264901935552["operator: -"]:::custom_operator;
  137264901935392 --> 137264901935552;
  137264901935712["identifier: n"]:::custom_identifier;
  137264901935552 --> 137264901935712;
  137264901935872["number: 1"]:::custom_number;
  137264901935552 --> 137264901935872;

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
  "134126260210240" [label="form: surround", shape="box", fillcolor="lightpink"];
  "134126260210320" [label="part: def", shape="box", fillcolor="#FFD8E1"];
  "134126260210240" -> "134126260210320";
  "134126260210400" [label="apply", shape="box", fillcolor="lightgreen"];
  "134126260210320" -> "134126260210400";
  "134126260210480" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "134126260210400" -> "134126260210480";
  "134126260210560" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "134126260210400" -> "134126260210560";
  "134126260210640" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "134126260210560" -> "134126260210640";
  "134126260210720" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "134126260210240" -> "134126260210720";
  "134126260210800" [label="form: surround", shape="box", fillcolor="lightpink"];
  "134126260210720" -> "134126260210800";
  "134126260210880" [label="part: if", shape="box", fillcolor="#FFD8E1"];
  "134126260210800" -> "134126260210880";
  "134126260210960" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "134126260210880" -> "134126260210960";
  "134126260211040" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "134126260210960" -> "134126260211040";
  "134126260211120" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "134126260210960" -> "134126260211120";
  "134126260211200" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "134126260210800" -> "134126260211200";
  "134126260211280" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "134126260211200" -> "134126260211280";
  "134126260211360" [label="part: else", shape="box", fillcolor="#FFD8E1"];
  "134126260210800" -> "134126260211360";
  "134126260211440" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "134126260211360" -> "134126260211440";
  "134126260211520" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "134126260211440" -> "134126260211520";
  "134126260211600" [label="apply", shape="box", fillcolor="lightgreen"];
  "134126260211440" -> "134126260211600";
  "134126260211760" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "134126260211600" -> "134126260211760";
  "134126260211920" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "134126260211600" -> "134126260211920";
  "134126260212080" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "134126260211920" -> "134126260212080";
  "134126260212240" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "134126260212080" -> "134126260212240";
  "134126260212400" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "134126260212080" -> "134126260212400";
}
```

![Image generated for example](images/factorial.png)
