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
  130739599410496["form"]:::custom_form;
  130739599410576["part: def"]:::custom_part;
  130739599410496 --> 130739599410576;
  130739599410656["apply"]:::custom_apply;
  130739599410576 --> 130739599410656;
  130739599410736["identifier: f"]:::custom_identifier;
  130739599410656 --> 130739599410736;
  130739599410816["arguments"]:::custom_arguments;
  130739599410656 --> 130739599410816;
  130739599410896["identifier: n"]:::custom_identifier;
  130739599410816 --> 130739599410896;
  130739599410976["part: _"]:::custom_part;
  130739599410496 --> 130739599410976;
  130739599411056["form"]:::custom_form;
  130739599410976 --> 130739599411056;
  130739599411136["part: if"]:::custom_part;
  130739599411056 --> 130739599411136;
  130739599411216["operator: <="]:::custom_operator;
  130739599411136 --> 130739599411216;
  130739599411296["identifier: n"]:::custom_identifier;
  130739599411216 --> 130739599411296;
  130739599411376["number: 1"]:::custom_number;
  130739599411216 --> 130739599411376;
  130739599411456["part: _"]:::custom_part;
  130739599411056 --> 130739599411456;
  130739599411536["number: 1"]:::custom_number;
  130739599411456 --> 130739599411536;
  130739599411616["part: else"]:::custom_part;
  130739599411056 --> 130739599411616;
  130739599411696["operator: *"]:::custom_operator;
  130739599411616 --> 130739599411696;
  130739599411776["identifier: n"]:::custom_identifier;
  130739599411696 --> 130739599411776;
  130739599411856["apply"]:::custom_apply;
  130739599411696 --> 130739599411856;
  130739599412016["identifier: f"]:::custom_identifier;
  130739599411856 --> 130739599412016;
  130739599412176["arguments"]:::custom_arguments;
  130739599411856 --> 130739599412176;
  130739599412336["operator: -"]:::custom_operator;
  130739599412176 --> 130739599412336;
  130739599412496["identifier: n"]:::custom_identifier;
  130739599412336 --> 130739599412496;
  130739599412656["number: 1"]:::custom_number;
  130739599412336 --> 130739599412656;

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
  "139838052484336" [label="form", shape="box", fillcolor="lightpink"];
  "139838052484416" [label="part: def", shape="box", fillcolor="#FFD8E1"];
  "139838052484336" -> "139838052484416";
  "139838052484496" [label="apply", shape="box", fillcolor="lightgreen"];
  "139838052484416" -> "139838052484496";
  "139838052484576" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "139838052484496" -> "139838052484576";
  "139838052484656" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "139838052484496" -> "139838052484656";
  "139838052484736" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "139838052484656" -> "139838052484736";
  "139838052484816" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "139838052484336" -> "139838052484816";
  "139838052484896" [label="form", shape="box", fillcolor="lightpink"];
  "139838052484816" -> "139838052484896";
  "139838052484976" [label="part: if", shape="box", fillcolor="#FFD8E1"];
  "139838052484896" -> "139838052484976";
  "139838052485056" [label="operator: <=", shape="box", fillcolor="#C0FFC0"];
  "139838052484976" -> "139838052485056";
  "139838052485136" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "139838052485056" -> "139838052485136";
  "139838052485216" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "139838052485056" -> "139838052485216";
  "139838052485296" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "139838052484896" -> "139838052485296";
  "139838052485376" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "139838052485296" -> "139838052485376";
  "139838052485456" [label="part: else", shape="box", fillcolor="#FFD8E1"];
  "139838052484896" -> "139838052485456";
  "139838052485536" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "139838052485456" -> "139838052485536";
  "139838052485616" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "139838052485536" -> "139838052485616";
  "139838052485696" [label="apply", shape="box", fillcolor="lightgreen"];
  "139838052485536" -> "139838052485696";
  "139838052485856" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "139838052485696" -> "139838052485856";
  "139838052486016" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "139838052485696" -> "139838052486016";
  "139838052486176" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "139838052486016" -> "139838052486176";
  "139838052486336" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "139838052486176" -> "139838052486336";
  "139838052486496" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "139838052486176" -> "139838052486496";
}
```

![Image generated for example](images/factorial.png)
