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
  129834112518960["operator: -"]:::custom_operator;
  129834112519040["operator: *"]:::custom_operator;
  129834112518960 --> 129834112519040;
  129834112519120["apply"]:::custom_apply;
  129834112519040 --> 129834112519120;
  129834112519200["identifier: cos"]:::custom_identifier;
  129834112519120 --> 129834112519200;
  129834112519280["arguments"]:::custom_arguments;
  129834112519120 --> 129834112519280;
  129834112519360["identifier: A"]:::custom_identifier;
  129834112519280 --> 129834112519360;
  129834112519440["apply"]:::custom_apply;
  129834112519040 --> 129834112519440;
  129834112519520["identifier: cos"]:::custom_identifier;
  129834112519440 --> 129834112519520;
  129834112519600["arguments"]:::custom_arguments;
  129834112519440 --> 129834112519600;
  129834112519680["identifier: B"]:::custom_identifier;
  129834112519600 --> 129834112519680;
  129834112519760["operator: *"]:::custom_operator;
  129834112518960 --> 129834112519760;
  129834112519840["apply"]:::custom_apply;
  129834112519760 --> 129834112519840;
  129834112519920["identifier: sin"]:::custom_identifier;
  129834112519840 --> 129834112519920;
  129834112520000["arguments"]:::custom_arguments;
  129834112519840 --> 129834112520000;
  129834112520080["identifier: A"]:::custom_identifier;
  129834112520000 --> 129834112520080;
  129834112520160["apply"]:::custom_apply;
  129834112519760 --> 129834112520160;
  129834112520240["identifier: sin"]:::custom_identifier;
  129834112520160 --> 129834112520240;
  129834112520320["arguments"]:::custom_arguments;
  129834112520160 --> 129834112520320;
  129834112520400["identifier: B"]:::custom_identifier;
  129834112520320 --> 129834112520400;

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
  "131985227991776" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "131985227991856" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "131985227991776" -> "131985227991856";
  "131985227991936" [label="apply", shape="box", fillcolor="lightgreen"];
  "131985227991856" -> "131985227991936";
  "131985227992016" [label="identifier: cos", shape="box", fillcolor="Honeydew"];
  "131985227991936" -> "131985227992016";
  "131985227992096" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "131985227991936" -> "131985227992096";
  "131985227992176" [label="identifier: A", shape="box", fillcolor="Honeydew"];
  "131985227992096" -> "131985227992176";
  "131985227992256" [label="apply", shape="box", fillcolor="lightgreen"];
  "131985227991856" -> "131985227992256";
  "131985227992336" [label="identifier: cos", shape="box", fillcolor="Honeydew"];
  "131985227992256" -> "131985227992336";
  "131985227992416" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "131985227992256" -> "131985227992416";
  "131985227992496" [label="identifier: B", shape="box", fillcolor="Honeydew"];
  "131985227992416" -> "131985227992496";
  "131985227992576" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "131985227991776" -> "131985227992576";
  "131985227992656" [label="apply", shape="box", fillcolor="lightgreen"];
  "131985227992576" -> "131985227992656";
  "131985227992736" [label="identifier: sin", shape="box", fillcolor="Honeydew"];
  "131985227992656" -> "131985227992736";
  "131985227992816" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "131985227992656" -> "131985227992816";
  "131985227992896" [label="identifier: A", shape="box", fillcolor="Honeydew"];
  "131985227992816" -> "131985227992896";
  "131985227992976" [label="apply", shape="box", fillcolor="lightgreen"];
  "131985227992576" -> "131985227992976";
  "131985227993056" [label="identifier: sin", shape="box", fillcolor="Honeydew"];
  "131985227992976" -> "131985227993056";
  "131985227993136" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "131985227992976" -> "131985227993136";
  "131985227993216" [label="identifier: B", shape="box", fillcolor="Honeydew"];
  "131985227993136" -> "131985227993216";
}
```

![Image generated for example](images/cosine_addition.png)
