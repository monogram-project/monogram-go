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
  132919964535440["operator: -"]:::custom_operator;
  132919964535520["operator: +"]:::custom_operator;
  132919964535440 --> 132919964535520;
  132919964535600["operator: *"]:::custom_operator;
  132919964535520 --> 132919964535600;
  132919964535680["number: 2"]:::custom_number;
  132919964535600 --> 132919964535680;
  132919964535760["operator: *"]:::custom_operator;
  132919964535600 --> 132919964535760;
  132919964535840["number: 100"]:::custom_number;
  132919964535760 --> 132919964535840;
  132919964535920["number: 100"]:::custom_number;
  132919964535760 --> 132919964535920;
  132919964536000["number: 100"]:::custom_number;
  132919964535520 --> 132919964536000;
  132919964536080["number: 1"]:::custom_number;
  132919964535440 --> 132919964536080;

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
  "136041549089344" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "136041549089424" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "136041549089344" -> "136041549089424";
  "136041549089504" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "136041549089424" -> "136041549089504";
  "136041549089584" [label="number: 2", shape="box", fillcolor="lightgoldenrodyellow"];
  "136041549089504" -> "136041549089584";
  "136041549089664" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "136041549089504" -> "136041549089664";
  "136041549089744" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "136041549089664" -> "136041549089744";
  "136041549089824" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "136041549089664" -> "136041549089824";
  "136041549089904" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "136041549089424" -> "136041549089904";
  "136041549089984" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "136041549089344" -> "136041549089984";
}
```

![Image generated for example](images/arithmetic.png)
