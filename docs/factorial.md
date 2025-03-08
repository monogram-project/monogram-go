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
  131429367614096["form: surround"]:::custom_form;
  131429367614176["part: def"]:::custom_part;
  131429367614096 --> 131429367614176;
  131429367614256["apply"]:::custom_apply;
  131429367614176 --> 131429367614256;
  131429367614336["identifier: f"]:::custom_identifier;
  131429367614256 --> 131429367614336;
  131429367614416["arguments"]:::custom_arguments;
  131429367614256 --> 131429367614416;
  131429367614496["identifier: n"]:::custom_identifier;
  131429367614416 --> 131429367614496;
  131429367614576["part: _"]:::custom_part;
  131429367614096 --> 131429367614576;
  131429367614656["form: surround"]:::custom_form;
  131429367614576 --> 131429367614656;
  131429367614736["part: if"]:::custom_part;
  131429367614656 --> 131429367614736;
  131429367614816["operator: <="]:::custom_operator;
  131429367614736 --> 131429367614816;
  131429367614896["identifier: n"]:::custom_identifier;
  131429367614816 --> 131429367614896;
  131429367614976["number: 1"]:::custom_number;
  131429367614816 --> 131429367614976;
  131429367615056["part: _"]:::custom_part;
  131429367614656 --> 131429367615056;
  131429367615136["number: 1"]:::custom_number;
  131429367615056 --> 131429367615136;
  131429367615216["part: else"]:::custom_part;
  131429367614656 --> 131429367615216;
  131429367615296["operator: *"]:::custom_operator;
  131429367615216 --> 131429367615296;
  131429367615376["identifier: n"]:::custom_identifier;
  131429367615296 --> 131429367615376;
  131429367615456["apply"]:::custom_apply;
  131429367615296 --> 131429367615456;
  131429367615616["identifier: f"]:::custom_identifier;
  131429367615456 --> 131429367615616;
  131429367615776["arguments"]:::custom_arguments;
  131429367615456 --> 131429367615776;
  131429367615936["operator: -"]:::custom_operator;
  131429367615776 --> 131429367615936;
  131429367616096["identifier: n"]:::custom_identifier;
  131429367615936 --> 131429367616096;
  131429367616256["number: 1"]:::custom_number;
  131429367615936 --> 131429367616256;

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
  "127534131923520" [label="form: surround", shape="box", fillcolor="lightpink"];
  "127534131923600" [label="part: def", shape="box", fillcolor="#FFD8E1"];
  "127534131923520" -> "127534131923600";
  "127534131923680" [label="apply", shape="box", fillcolor="lightgreen"];
  "127534131923600" -> "127534131923680";
  "127534131923760" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "127534131923680" -> "127534131923760";
  "127534131923840" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "127534131923680" -> "127534131923840";
  "127534131923920" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "127534131923840" -> "127534131923920";
  "127534131924000" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "127534131923520" -> "127534131924000";
  "127534131924080" [label="form: surround", shape="box", fillcolor="lightpink"];
  "127534131924000" -> "127534131924080";
  "127534131924160" [label="part: if", shape="box", fillcolor="#FFD8E1"];
  "127534131924080" -> "127534131924160";
  "127534131924240" [label="operator: <=", shape="box", fillcolor="#C0FFC0"];
  "127534131924160" -> "127534131924240";
  "127534131924320" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "127534131924240" -> "127534131924320";
  "127534131924400" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "127534131924240" -> "127534131924400";
  "127534131924480" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "127534131924080" -> "127534131924480";
  "127534131924560" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "127534131924480" -> "127534131924560";
  "127534131924640" [label="part: else", shape="box", fillcolor="#FFD8E1"];
  "127534131924080" -> "127534131924640";
  "127534131924720" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "127534131924640" -> "127534131924720";
  "127534131924800" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "127534131924720" -> "127534131924800";
  "127534131924880" [label="apply", shape="box", fillcolor="lightgreen"];
  "127534131924720" -> "127534131924880";
  "127534131925040" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "127534131924880" -> "127534131925040";
  "127534131925200" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "127534131924880" -> "127534131925200";
  "127534131925360" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "127534131925200" -> "127534131925360";
  "127534131925520" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "127534131925360" -> "127534131925520";
  "127534131925680" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "127534131925360" -> "127534131925680";
}
```

![Image generated for example](images/factorial.png)
