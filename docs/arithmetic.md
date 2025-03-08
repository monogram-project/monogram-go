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
  129125483825808["operator: -"]:::custom_operator;
  129125483825888["operator: +"]:::custom_operator;
  129125483825808 --> 129125483825888;
  129125483825968["operator: *"]:::custom_operator;
  129125483825888 --> 129125483825968;
  129125483826048["number: 2"]:::custom_number;
  129125483825968 --> 129125483826048;
  129125483826128["operator: *"]:::custom_operator;
  129125483825968 --> 129125483826128;
  129125483826208["number: 100"]:::custom_number;
  129125483826128 --> 129125483826208;
  129125483826288["number: 100"]:::custom_number;
  129125483826128 --> 129125483826288;
  129125483826368["number: 100"]:::custom_number;
  129125483825888 --> 129125483826368;
  129125483826448["number: 1"]:::custom_number;
  129125483825808 --> 129125483826448;

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
  "131016572799552" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "131016572799632" [label="operator: +", shape="box", fillcolor="#C0FFC0"];
  "131016572799552" -> "131016572799632";
  "131016572799712" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "131016572799632" -> "131016572799712";
  "131016572799792" [label="number: 2", shape="box", fillcolor="lightgoldenrodyellow"];
  "131016572799712" -> "131016572799792";
  "131016572799872" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "131016572799712" -> "131016572799872";
  "131016572799952" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "131016572799872" -> "131016572799952";
  "131016572800032" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "131016572799872" -> "131016572800032";
  "131016572800112" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "131016572799632" -> "131016572800112";
  "131016572800192" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "131016572799552" -> "131016572800192";
}
```

![Image generated for example](images/arithmetic.png)
