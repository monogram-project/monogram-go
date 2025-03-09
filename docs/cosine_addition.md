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
  125861360682544["operator: -"]:::custom_operator;
  125861360682624["operator: *"]:::custom_operator;
  125861360682544 --> 125861360682624;
  125861360682704["apply"]:::custom_apply;
  125861360682624 --> 125861360682704;
  125861360682784["identifier: cos"]:::custom_identifier;
  125861360682704 --> 125861360682784;
  125861360682864["arguments"]:::custom_arguments;
  125861360682704 --> 125861360682864;
  125861360682944["identifier: A"]:::custom_identifier;
  125861360682864 --> 125861360682944;
  125861360683024["apply"]:::custom_apply;
  125861360682624 --> 125861360683024;
  125861360683104["identifier: cos"]:::custom_identifier;
  125861360683024 --> 125861360683104;
  125861360683184["arguments"]:::custom_arguments;
  125861360683024 --> 125861360683184;
  125861360683264["identifier: B"]:::custom_identifier;
  125861360683184 --> 125861360683264;
  125861360683344["operator: *"]:::custom_operator;
  125861360682544 --> 125861360683344;
  125861360683424["apply"]:::custom_apply;
  125861360683344 --> 125861360683424;
  125861360683504["identifier: sin"]:::custom_identifier;
  125861360683424 --> 125861360683504;
  125861360683584["arguments"]:::custom_arguments;
  125861360683424 --> 125861360683584;
  125861360683664["identifier: A"]:::custom_identifier;
  125861360683584 --> 125861360683664;
  125861360683744["apply"]:::custom_apply;
  125861360683344 --> 125861360683744;
  125861360683824["identifier: sin"]:::custom_identifier;
  125861360683744 --> 125861360683824;
  125861360683904["arguments"]:::custom_arguments;
  125861360683744 --> 125861360683904;
  125861360683984["identifier: B"]:::custom_identifier;
  125861360683904 --> 125861360683984;

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
  "137852306425392" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "137852306425472" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "137852306425392" -> "137852306425472";
  "137852306425552" [label="apply", shape="box", fillcolor="lightgreen"];
  "137852306425472" -> "137852306425552";
  "137852306425632" [label="identifier: cos", shape="box", fillcolor="Honeydew"];
  "137852306425552" -> "137852306425632";
  "137852306425712" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "137852306425552" -> "137852306425712";
  "137852306425792" [label="identifier: A", shape="box", fillcolor="Honeydew"];
  "137852306425712" -> "137852306425792";
  "137852306425872" [label="apply", shape="box", fillcolor="lightgreen"];
  "137852306425472" -> "137852306425872";
  "137852306425952" [label="identifier: cos", shape="box", fillcolor="Honeydew"];
  "137852306425872" -> "137852306425952";
  "137852306426032" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "137852306425872" -> "137852306426032";
  "137852306426112" [label="identifier: B", shape="box", fillcolor="Honeydew"];
  "137852306426032" -> "137852306426112";
  "137852306426192" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "137852306425392" -> "137852306426192";
  "137852306426272" [label="apply", shape="box", fillcolor="lightgreen"];
  "137852306426192" -> "137852306426272";
  "137852306426352" [label="identifier: sin", shape="box", fillcolor="Honeydew"];
  "137852306426272" -> "137852306426352";
  "137852306426432" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "137852306426272" -> "137852306426432";
  "137852306426512" [label="identifier: A", shape="box", fillcolor="Honeydew"];
  "137852306426432" -> "137852306426512";
  "137852306426592" [label="apply", shape="box", fillcolor="lightgreen"];
  "137852306426192" -> "137852306426592";
  "137852306426672" [label="identifier: sin", shape="box", fillcolor="Honeydew"];
  "137852306426592" -> "137852306426672";
  "137852306426752" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "137852306426592" -> "137852306426752";
  "137852306426832" [label="identifier: B", shape="box", fillcolor="Honeydew"];
  "137852306426752" -> "137852306426832";
}
```

![Image generated for example](images/cosine_addition.png)
