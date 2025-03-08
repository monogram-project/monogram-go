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
  132449757055632["operator: -"]:::custom_operator;
  132449757055712["operator: +"]:::custom_operator;
  132449757055632 --> 132449757055712;
  132449757055792["operator: *"]:::custom_operator;
  132449757055712 --> 132449757055792;
  132449757055872["number: 2"]:::custom_number;
  132449757055792 --> 132449757055872;
  132449757055952["operator: *"]:::custom_operator;
  132449757055792 --> 132449757055952;
  132449757056032["number: 100"]:::custom_number;
  132449757055952 --> 132449757056032;
  132449757056112["number: 100"]:::custom_number;
  132449757055952 --> 132449757056112;
  132449757056192["number: 100"]:::custom_number;
  132449757055712 --> 132449757056192;
  132449757056272["number: 1"]:::custom_number;
  132449757055632 --> 132449757056272;

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
  "130767434287680" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "130767434287760" [label="operator: +", shape="box", fillcolor="#C0FFC0"];
  "130767434287680" -> "130767434287760";
  "130767434287840" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "130767434287760" -> "130767434287840";
  "130767434287920" [label="number: 2", shape="box", fillcolor="lightgoldenrodyellow"];
  "130767434287840" -> "130767434287920";
  "130767434288000" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "130767434287840" -> "130767434288000";
  "130767434288080" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "130767434288000" -> "130767434288080";
  "130767434288160" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "130767434288000" -> "130767434288160";
  "130767434288240" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "130767434287760" -> "130767434288240";
  "130767434288320" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "130767434287680" -> "130767434288320";
}
```

![Image generated for example](images/arithmetic.png)
