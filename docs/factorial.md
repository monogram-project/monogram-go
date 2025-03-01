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
  134341447484736["form"]:::custom_form;
  134341447484816["part: def"]:::custom_part;
  134341447484736 --> 134341447484816;
  134341447484896["apply"]:::custom_apply;
  134341447484816 --> 134341447484896;
  134341447484976["identifier: f"]:::custom_identifier;
  134341447484896 --> 134341447484976;
  134341447485056["arguments"]:::custom_arguments;
  134341447484896 --> 134341447485056;
  134341447485136["identifier: n"]:::custom_identifier;
  134341447485056 --> 134341447485136;
  134341447485216["part: _"]:::custom_part;
  134341447484736 --> 134341447485216;
  134341447485296["form"]:::custom_form;
  134341447485216 --> 134341447485296;
  134341447485376["part: if"]:::custom_part;
  134341447485296 --> 134341447485376;
  134341447485456["operator: <="]:::custom_operator;
  134341447485376 --> 134341447485456;
  134341447485536["identifier: n"]:::custom_identifier;
  134341447485456 --> 134341447485536;
  134341447485616["number: 1"]:::custom_number;
  134341447485456 --> 134341447485616;
  134341447485696["part: _"]:::custom_part;
  134341447485296 --> 134341447485696;
  134341447485776["number: 1"]:::custom_number;
  134341447485696 --> 134341447485776;
  134341447485856["part: else"]:::custom_part;
  134341447485296 --> 134341447485856;
  134341447485936["operator: *"]:::custom_operator;
  134341447485856 --> 134341447485936;
  134341447486016["identifier: n"]:::custom_identifier;
  134341447485936 --> 134341447486016;
  134341447486096["apply"]:::custom_apply;
  134341447485936 --> 134341447486096;
  134341447486256["identifier: f"]:::custom_identifier;
  134341447486096 --> 134341447486256;
  134341447486416["arguments"]:::custom_arguments;
  134341447486096 --> 134341447486416;
  134341447486576["operator: -"]:::custom_operator;
  134341447486416 --> 134341447486576;
  134341447486736["identifier: n"]:::custom_identifier;
  134341447486576 --> 134341447486736;
  134341447486896["number: 1"]:::custom_number;
  134341447486576 --> 134341447486896;

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
  "138715029685568" [label="form", shape="box", fillcolor="lightpink"];
  "138715029685648" [label="part: def", shape="box", fillcolor="#FFD8E1"];
  "138715029685568" -> "138715029685648";
  "138715029685728" [label="apply", shape="box", fillcolor="lightgreen"];
  "138715029685648" -> "138715029685728";
  "138715029685808" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "138715029685728" -> "138715029685808";
  "138715029685888" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "138715029685728" -> "138715029685888";
  "138715029685968" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "138715029685888" -> "138715029685968";
  "138715029686048" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "138715029685568" -> "138715029686048";
  "138715029686128" [label="form", shape="box", fillcolor="lightpink"];
  "138715029686048" -> "138715029686128";
  "138715029686208" [label="part: if", shape="box", fillcolor="#FFD8E1"];
  "138715029686128" -> "138715029686208";
  "138715029686288" [label="operator: <=", shape="box", fillcolor="#C0FFC0"];
  "138715029686208" -> "138715029686288";
  "138715029686368" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "138715029686288" -> "138715029686368";
  "138715029686448" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "138715029686288" -> "138715029686448";
  "138715029686528" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "138715029686128" -> "138715029686528";
  "138715029686608" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "138715029686528" -> "138715029686608";
  "138715029686688" [label="part: else", shape="box", fillcolor="#FFD8E1"];
  "138715029686128" -> "138715029686688";
  "138715029686768" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "138715029686688" -> "138715029686768";
  "138715029686848" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "138715029686768" -> "138715029686848";
  "138715029686928" [label="apply", shape="box", fillcolor="lightgreen"];
  "138715029686768" -> "138715029686928";
  "138715029687088" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "138715029686928" -> "138715029687088";
  "138715029687248" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "138715029686928" -> "138715029687248";
  "138715029687408" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "138715029687248" -> "138715029687408";
  "138715029687568" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "138715029687408" -> "138715029687568";
  "138715029687728" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "138715029687408" -> "138715029687728";
}
```

![Image generated for example](images/factorial.png)
