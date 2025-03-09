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
  133315168832304["operator: -"]:::custom_operator;
  133315168832384["operator: *"]:::custom_operator;
  133315168832304 --> 133315168832384;
  133315168832464["apply"]:::custom_apply;
  133315168832384 --> 133315168832464;
  133315168832544["identifier: cos"]:::custom_identifier;
  133315168832464 --> 133315168832544;
  133315168832624["arguments"]:::custom_arguments;
  133315168832464 --> 133315168832624;
  133315168832704["identifier: A"]:::custom_identifier;
  133315168832624 --> 133315168832704;
  133315168832784["apply"]:::custom_apply;
  133315168832384 --> 133315168832784;
  133315168832864["identifier: cos"]:::custom_identifier;
  133315168832784 --> 133315168832864;
  133315168832944["arguments"]:::custom_arguments;
  133315168832784 --> 133315168832944;
  133315168833024["identifier: B"]:::custom_identifier;
  133315168832944 --> 133315168833024;
  133315168833104["operator: *"]:::custom_operator;
  133315168832304 --> 133315168833104;
  133315168833184["apply"]:::custom_apply;
  133315168833104 --> 133315168833184;
  133315168833264["identifier: sin"]:::custom_identifier;
  133315168833184 --> 133315168833264;
  133315168833344["arguments"]:::custom_arguments;
  133315168833184 --> 133315168833344;
  133315168833424["identifier: A"]:::custom_identifier;
  133315168833344 --> 133315168833424;
  133315168833504["apply"]:::custom_apply;
  133315168833104 --> 133315168833504;
  133315168833584["identifier: sin"]:::custom_identifier;
  133315168833504 --> 133315168833584;
  133315168833664["arguments"]:::custom_arguments;
  133315168833504 --> 133315168833664;
  133315168833744["identifier: B"]:::custom_identifier;
  133315168833664 --> 133315168833744;

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
    <operator name="*" syntax="infix">
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
    <operator name="*" syntax="infix">
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
    "syntax": "infix",
    "children": [
        {
            "role": "operator",
            "name": "*",
            "syntax": "infix",
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
            "syntax": "infix",
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
syntax: infix
children:
- role: operator
  name: '*'
  syntax: infix
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
  syntax: infix
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
  "138166528379616" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "138166528379696" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "138166528379616" -> "138166528379696";
  "138166528379776" [label="apply", shape="box", fillcolor="lightgreen"];
  "138166528379696" -> "138166528379776";
  "138166528379856" [label="identifier: cos", shape="box", fillcolor="Honeydew"];
  "138166528379776" -> "138166528379856";
  "138166528379936" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "138166528379776" -> "138166528379936";
  "138166528380016" [label="identifier: A", shape="box", fillcolor="Honeydew"];
  "138166528379936" -> "138166528380016";
  "138166528380096" [label="apply", shape="box", fillcolor="lightgreen"];
  "138166528379696" -> "138166528380096";
  "138166528380176" [label="identifier: cos", shape="box", fillcolor="Honeydew"];
  "138166528380096" -> "138166528380176";
  "138166528380256" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "138166528380096" -> "138166528380256";
  "138166528380336" [label="identifier: B", shape="box", fillcolor="Honeydew"];
  "138166528380256" -> "138166528380336";
  "138166528380416" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "138166528379616" -> "138166528380416";
  "138166528380496" [label="apply", shape="box", fillcolor="lightgreen"];
  "138166528380416" -> "138166528380496";
  "138166528380576" [label="identifier: sin", shape="box", fillcolor="Honeydew"];
  "138166528380496" -> "138166528380576";
  "138166528380656" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "138166528380496" -> "138166528380656";
  "138166528380736" [label="identifier: A", shape="box", fillcolor="Honeydew"];
  "138166528380656" -> "138166528380736";
  "138166528380816" [label="apply", shape="box", fillcolor="lightgreen"];
  "138166528380416" -> "138166528380816";
  "138166528380896" [label="identifier: sin", shape="box", fillcolor="Honeydew"];
  "138166528380816" -> "138166528380896";
  "138166528380976" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "138166528380816" -> "138166528380976";
  "138166528381056" [label="identifier: B", shape="box", fillcolor="Honeydew"];
  "138166528380976" -> "138166528381056";
}
```

![Image generated for example](images/cosine_addition.png)
