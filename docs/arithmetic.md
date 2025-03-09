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
  133333870806336["operator: -"]:::custom_operator;
  133333870806416["operator: +"]:::custom_operator;
  133333870806336 --> 133333870806416;
  133333870806496["operator: *"]:::custom_operator;
  133333870806416 --> 133333870806496;
  133333870806576["number: 2"]:::custom_number;
  133333870806496 --> 133333870806576;
  133333870806656["operator: *"]:::custom_operator;
  133333870806496 --> 133333870806656;
  133333870806736["number: 100"]:::custom_number;
  133333870806656 --> 133333870806736;
  133333870806816["number: 100"]:::custom_number;
  133333870806656 --> 133333870806816;
  133333870806896["number: 100"]:::custom_number;
  133333870806416 --> 133333870806896;
  133333870806976["number: 1"]:::custom_number;
  133333870806336 --> 133333870806976;

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
  "126301167011136" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "126301167011216" [label="operator: +", shape="box", fillcolor="#C0FFC0"];
  "126301167011136" -> "126301167011216";
  "126301167011296" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "126301167011216" -> "126301167011296";
  "126301167011376" [label="number: 2", shape="box", fillcolor="lightgoldenrodyellow"];
  "126301167011296" -> "126301167011376";
  "126301167011456" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "126301167011296" -> "126301167011456";
  "126301167011536" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "126301167011456" -> "126301167011536";
  "126301167011616" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "126301167011456" -> "126301167011616";
  "126301167011696" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "126301167011216" -> "126301167011696";
  "126301167011776" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "126301167011136" -> "126301167011776";
}
```

![Image generated for example](images/arithmetic.png)
