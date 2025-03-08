# Cosine addition formula

## Monogram

```txt
cos(A) * cos(B) - sin(A) * sin(B)

```

## Mermaid diagram

We can target Mermaid's flowchart as an output format. 
And this is what it looks like:

```mermaid
graph LR
  125183908645680["operator: -"]:::custom_operator;
  125183908645760["operator: *"]:::custom_operator;
  125183908645680 --> 125183908645760;
  125183908645840["apply"]:::custom_apply;
  125183908645760 --> 125183908645840;
  125183908645920["identifier: cos"]:::custom_identifier;
  125183908645840 --> 125183908645920;
  125183908646000["arguments"]:::custom_arguments;
  125183908645840 --> 125183908646000;
  125183908646080["identifier: A"]:::custom_identifier;
  125183908646000 --> 125183908646080;
  125183908646160["apply"]:::custom_apply;
  125183908645760 --> 125183908646160;
  125183908646240["identifier: cos"]:::custom_identifier;
  125183908646160 --> 125183908646240;
  125183908646320["arguments"]:::custom_arguments;
  125183908646160 --> 125183908646320;
  125183908646400["identifier: B"]:::custom_identifier;
  125183908646320 --> 125183908646400;
  125183908646480["operator: *"]:::custom_operator;
  125183908645680 --> 125183908646480;
  125183908646560["apply"]:::custom_apply;
  125183908646480 --> 125183908646560;
  125183908646640["identifier: sin"]:::custom_identifier;
  125183908646560 --> 125183908646640;
  125183908646720["arguments"]:::custom_arguments;
  125183908646560 --> 125183908646720;
  125183908646800["identifier: A"]:::custom_identifier;
  125183908646720 --> 125183908646800;
  125183908646880["apply"]:::custom_apply;
  125183908646480 --> 125183908646880;
  125183908646960["identifier: sin"]:::custom_identifier;
  125183908646880 --> 125183908646960;
  125183908647040["arguments"]:::custom_arguments;
  125183908646880 --> 125183908647040;
  125183908647120["identifier: B"]:::custom_identifier;
  125183908647040 --> 125183908647120;

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
    <operator name="*">
        <apply kind="parentheses" separator="undefined">
            <identifier name="cos"/>
            <arguments>
                <identifier name="A"/>
            </arguments>
        </apply>
        <apply kind="parentheses" separator="undefined">
            <identifier name="cos"/>
            <arguments>
                <identifier name="B"/>
            </arguments>
        </apply>
    </operator>
    <operator name="*">
        <apply kind="parentheses" separator="undefined">
            <identifier name="sin"/>
            <arguments>
                <identifier name="A"/>
            </arguments>
        </apply>
        <apply kind="parentheses" separator="undefined">
            <identifier name="sin"/>
            <arguments>
                <identifier name="B"/>
            </arguments>
        </apply>
    </operator>
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
            "name": "*",
            "children": [
                {
                    "role": "apply",
                    "kind": "parentheses",
                    "separator": "undefined",
                    "children": [
                        {
                            "role": "identifier",
                            "name": "cos"
                        },
                        {
                            "role": "arguments",
                            "children": [
                                {
                                    "role": "identifier",
                                    "name": "A"
                                }
                            ]
                        }
                    ]
                },
                {
                    "role": "apply",
                    "kind": "parentheses",
                    "separator": "undefined",
                    "children": [
                        {
                            "role": "identifier",
                            "name": "cos"
                        },
                        {
                            "role": "arguments",
                            "children": [
                                {
                                    "role": "identifier",
                                    "name": "B"
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "role": "operator",
            "name": "*",
            "children": [
                {
                    "role": "apply",
                    "kind": "parentheses",
                    "separator": "undefined",
                    "children": [
                        {
                            "role": "identifier",
                            "name": "sin"
                        },
                        {
                            "role": "arguments",
                            "children": [
                                {
                                    "role": "identifier",
                                    "name": "A"
                                }
                            ]
                        }
                    ]
                },
                {
                    "role": "apply",
                    "kind": "parentheses",
                    "separator": "undefined",
                    "children": [
                        {
                            "role": "identifier",
                            "name": "sin"
                        },
                        {
                            "role": "arguments",
                            "children": [
                                {
                                    "role": "identifier",
                                    "name": "B"
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
role: operator
name: '-'
children:
- role: operator
  name: '*'
  children:
  - role: apply
    kind: parentheses
    separator: undefined
    children:
    - role: identifier
      name: cos
    - role: arguments
      children:
      - role: identifier
        name: A
  - role: apply
    kind: parentheses
    separator: undefined
    children:
    - role: identifier
      name: cos
    - role: arguments
      children:
      - role: identifier
        name: B
- role: operator
  name: '*'
  children:
  - role: apply
    kind: parentheses
    separator: undefined
    children:
    - role: identifier
      name: sin
    - role: arguments
      children:
      - role: identifier
        name: A
  - role: apply
    kind: parentheses
    separator: undefined
    children:
    - role: identifier
      name: sin
    - role: arguments
      children:
      - role: identifier
        name: B

```

## Graphviz Dot format

```dot
digraph G {
  bgcolor="transparent";
  node [shape="box", style="filled", fontname="Ubuntu Mono"];
  "127902765353696" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "127902765353776" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "127902765353696" -> "127902765353776";
  "127902765353856" [label="apply", shape="box", fillcolor="lightgreen"];
  "127902765353776" -> "127902765353856";
  "127902765353936" [label="identifier: cos", shape="box", fillcolor="Honeydew"];
  "127902765353856" -> "127902765353936";
  "127902765354016" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "127902765353856" -> "127902765354016";
  "127902765354096" [label="identifier: A", shape="box", fillcolor="Honeydew"];
  "127902765354016" -> "127902765354096";
  "127902765354176" [label="apply", shape="box", fillcolor="lightgreen"];
  "127902765353776" -> "127902765354176";
  "127902765354256" [label="identifier: cos", shape="box", fillcolor="Honeydew"];
  "127902765354176" -> "127902765354256";
  "127902765354336" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "127902765354176" -> "127902765354336";
  "127902765354416" [label="identifier: B", shape="box", fillcolor="Honeydew"];
  "127902765354336" -> "127902765354416";
  "127902765354496" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "127902765353696" -> "127902765354496";
  "127902765354576" [label="apply", shape="box", fillcolor="lightgreen"];
  "127902765354496" -> "127902765354576";
  "127902765354656" [label="identifier: sin", shape="box", fillcolor="Honeydew"];
  "127902765354576" -> "127902765354656";
  "127902765354736" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "127902765354576" -> "127902765354736";
  "127902765354816" [label="identifier: A", shape="box", fillcolor="Honeydew"];
  "127902765354736" -> "127902765354816";
  "127902765354896" [label="apply", shape="box", fillcolor="lightgreen"];
  "127902765354496" -> "127902765354896";
  "127902765354976" [label="identifier: sin", shape="box", fillcolor="Honeydew"];
  "127902765354896" -> "127902765354976";
  "127902765355056" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "127902765354896" -> "127902765355056";
  "127902765355136" [label="identifier: B", shape="box", fillcolor="Honeydew"];
  "127902765355056" -> "127902765355136";
}
```

![Image generated for example](images/cosine_addition.png)
