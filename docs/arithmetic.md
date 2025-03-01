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
  137481476474176["operator: -"]:::custom_operator;
  137481476474256["operator: +"]:::custom_operator;
  137481476474176 --> 137481476474256;
  137481476474336["operator: *"]:::custom_operator;
  137481476474256 --> 137481476474336;
  137481476474416["number: 2"]:::custom_number;
  137481476474336 --> 137481476474416;
  137481476474496["operator: *"]:::custom_operator;
  137481476474336 --> 137481476474496;
  137481476474576["number: 100"]:::custom_number;
  137481476474496 --> 137481476474576;
  137481476474656["number: 100"]:::custom_number;
  137481476474496 --> 137481476474656;
  137481476474736["number: 100"]:::custom_number;
  137481476474256 --> 137481476474736;
  137481476474816["number: 1"]:::custom_number;
  137481476474176 --> 137481476474816;

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
<operator name="-">
    <operator name="+">
        <operator name="*">
            <number value="2"/>
            <operator name="*">
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
    "children": [
        {
            "role": "operator",
            "name": "+",
            "children": [
                {
                    "role": "operator",
                    "name": "*",
                    "children": [
                        {
                            "role": "number",
                            "value": 2
                        },
                        {
                            "role": "operator",
                            "name": "*",
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
children:
- role: operator
  name: +
  children:
  - role: operator
    name: '*'
    children:
    - role: number
      value: 2
    - role: operator
      name: '*'
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
  "139194789889344" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "139194789889424" [label="operator: +", shape="box", fillcolor="#C0FFC0"];
  "139194789889344" -> "139194789889424";
  "139194789889504" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "139194789889424" -> "139194789889504";
  "139194789889584" [label="number: 2", shape="box", fillcolor="lightgoldenrodyellow"];
  "139194789889504" -> "139194789889584";
  "139194789889664" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "139194789889504" -> "139194789889664";
  "139194789889744" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "139194789889664" -> "139194789889744";
  "139194789889824" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "139194789889664" -> "139194789889824";
  "139194789889904" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "139194789889424" -> "139194789889904";
  "139194789889984" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "139194789889344" -> "139194789889984";
}
```

![Image generated for example](images/arithmetic.png)
