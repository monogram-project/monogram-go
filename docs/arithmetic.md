# Simple arithmetic expression

## Monogram

```txt
2 * 100 * 100 + 100 - 1
```

## Mermaid diagram

We can target Mermaid's flowchart as an output format. 
And this is what it looks like:

```mermaid
graph LR
  125139368102464["operator: -"]:::custom_operator;
  125139368102544["operator: +"]:::custom_operator;
  125139368102464 --> 125139368102544;
  125139368102624["operator: *"]:::custom_operator;
  125139368102544 --> 125139368102624;
  125139368102704["number: 2"]:::custom_number;
  125139368102624 --> 125139368102704;
  125139368102784["operator: *"]:::custom_operator;
  125139368102624 --> 125139368102784;
  125139368102864["number: 100"]:::custom_number;
  125139368102784 --> 125139368102864;
  125139368102944["number: 100"]:::custom_number;
  125139368102784 --> 125139368102944;
  125139368103024["number: 100"]:::custom_number;
  125139368102544 --> 125139368103024;
  125139368103104["number: 1"]:::custom_number;
  125139368102464 --> 125139368103104;

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
<operator name="-" syntax="infix">
    <operator name="+" syntax="infix">
        <operator name="*" syntax="infix">
            <number value="2"/>
            <operator name="*" syntax="infix">
                <number value="100"/>
                <number value="100"/>
            </operator>
        </operator>
        <number value="100"/>
    </operator>
    <number value="1"/>
</operator>
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
    "role": "operator",
    "name": "-",
    "syntax": "infix",
    "children": [
        {
            "role": "operator",
            "name": "+",
            "syntax": "infix",
            "children": [
                {
                    "role": "operator",
                    "name": "*",
                    "syntax": "infix",
                    "children": [
                        {
                            "role": "number",
                            "value": 2
                        },
                        {
                            "role": "operator",
                            "name": "*",
                            "syntax": "infix",
                            "children": [
                                {
                                    "role": "number",
                                    "value": 100
                                },
                                {
                                    "role": "number",
                                    "value": 100
                                }
                            ]
                        }
                    ]
                },
                {
                    "role": "number",
                    "value": 100
                }
            ]
        },
        {
            "role": "number",
            "value": 1
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
role: operator
name: '-'
syntax: infix
children:
- role: operator
  name: +
  syntax: infix
  children:
  - role: operator
    name: '*'
    syntax: infix
    children:
    - role: number
      value: 2
    - role: operator
      name: '*'
      syntax: infix
      children:
      - role: number
        value: 100
      - role: number
        value: 100
  - role: number
    value: 100
- role: number
  value: 1

```

## Graphviz Dot format

```dot
digraph G {
  bgcolor="transparent";
  node [shape="box", style="filled", fontname="Ubuntu Mono"];
  "133579031022144" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "133579031022224" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "133579031022144" -> "133579031022224";
  "133579031022304" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "133579031022224" -> "133579031022304";
  "133579031022384" [label="number: 2", shape="box", fillcolor="lightgoldenrodyellow"];
  "133579031022304" -> "133579031022384";
  "133579031022464" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "133579031022304" -> "133579031022464";
  "133579031022544" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "133579031022464" -> "133579031022544";
  "133579031022624" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "133579031022464" -> "133579031022624";
  "133579031022704" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "133579031022224" -> "133579031022704";
  "133579031022784" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "133579031022144" -> "133579031022784";
}
```

![Image generated for example](images/arithmetic.png)
