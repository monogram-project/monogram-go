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
  136753473669776["operator: -"]:::custom_operator;
  136753473669856["operator: +"]:::custom_operator;
  136753473669776 --> 136753473669856;
  136753473669936["operator: *"]:::custom_operator;
  136753473669856 --> 136753473669936;
  136753473670016["number: 2"]:::custom_number;
  136753473669936 --> 136753473670016;
  136753473670096["operator: *"]:::custom_operator;
  136753473669936 --> 136753473670096;
  136753473670176["number: 100"]:::custom_number;
  136753473670096 --> 136753473670176;
  136753473670256["number: 100"]:::custom_number;
  136753473670096 --> 136753473670256;
  136753473670336["number: 100"]:::custom_number;
  136753473669856 --> 136753473670336;
  136753473670416["number: 1"]:::custom_number;
  136753473669776 --> 136753473670416;

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
  "138245232413248" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "138245232413328" [label="operator: +", shape="box", fillcolor="#C0FFC0"];
  "138245232413248" -> "138245232413328";
  "138245232413408" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "138245232413328" -> "138245232413408";
  "138245232413488" [label="number: 2", shape="box", fillcolor="lightgoldenrodyellow"];
  "138245232413408" -> "138245232413488";
  "138245232413568" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "138245232413408" -> "138245232413568";
  "138245232413648" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "138245232413568" -> "138245232413648";
  "138245232413728" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "138245232413568" -> "138245232413728";
  "138245232413808" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "138245232413328" -> "138245232413808";
  "138245232413888" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "138245232413248" -> "138245232413888";
}
```

![Image generated for example](images/arithmetic.png)
