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
  124136858340144["operator: -"]:::custom_operator;
  124136858340224["operator: *"]:::custom_operator;
  124136858340144 --> 124136858340224;
  124136858340304["apply"]:::custom_apply;
  124136858340224 --> 124136858340304;
  124136858340384["identifier: cos"]:::custom_identifier;
  124136858340304 --> 124136858340384;
  124136858340464["arguments"]:::custom_arguments;
  124136858340304 --> 124136858340464;
  124136858340544["identifier: A"]:::custom_identifier;
  124136858340464 --> 124136858340544;
  124136858340624["apply"]:::custom_apply;
  124136858340224 --> 124136858340624;
  124136858340704["identifier: cos"]:::custom_identifier;
  124136858340624 --> 124136858340704;
  124136858340784["arguments"]:::custom_arguments;
  124136858340624 --> 124136858340784;
  124136858340864["identifier: B"]:::custom_identifier;
  124136858340784 --> 124136858340864;
  124136858340944["operator: *"]:::custom_operator;
  124136858340144 --> 124136858340944;
  124136858341024["apply"]:::custom_apply;
  124136858340944 --> 124136858341024;
  124136858341104["identifier: sin"]:::custom_identifier;
  124136858341024 --> 124136858341104;
  124136858341184["arguments"]:::custom_arguments;
  124136858341024 --> 124136858341184;
  124136858341264["identifier: A"]:::custom_identifier;
  124136858341184 --> 124136858341264;
  124136858341344["apply"]:::custom_apply;
  124136858340944 --> 124136858341344;
  124136858341424["identifier: sin"]:::custom_identifier;
  124136858341344 --> 124136858341424;
  124136858341504["arguments"]:::custom_arguments;
  124136858341344 --> 124136858341504;
  124136858341584["identifier: B"]:::custom_identifier;
  124136858341504 --> 124136858341584;

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
  "124329624357600" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "124329624357680" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "124329624357600" -> "124329624357680";
  "124329624357760" [label="apply", shape="box", fillcolor="lightgreen"];
  "124329624357680" -> "124329624357760";
  "124329624357840" [label="identifier: cos", shape="box", fillcolor="Honeydew"];
  "124329624357760" -> "124329624357840";
  "124329624357920" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "124329624357760" -> "124329624357920";
  "124329624358000" [label="identifier: A", shape="box", fillcolor="Honeydew"];
  "124329624357920" -> "124329624358000";
  "124329624358080" [label="apply", shape="box", fillcolor="lightgreen"];
  "124329624357680" -> "124329624358080";
  "124329624358160" [label="identifier: cos", shape="box", fillcolor="Honeydew"];
  "124329624358080" -> "124329624358160";
  "124329624358240" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "124329624358080" -> "124329624358240";
  "124329624358320" [label="identifier: B", shape="box", fillcolor="Honeydew"];
  "124329624358240" -> "124329624358320";
  "124329624358400" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "124329624357600" -> "124329624358400";
  "124329624358480" [label="apply", shape="box", fillcolor="lightgreen"];
  "124329624358400" -> "124329624358480";
  "124329624358560" [label="identifier: sin", shape="box", fillcolor="Honeydew"];
  "124329624358480" -> "124329624358560";
  "124329624358640" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "124329624358480" -> "124329624358640";
  "124329624358720" [label="identifier: A", shape="box", fillcolor="Honeydew"];
  "124329624358640" -> "124329624358720";
  "124329624358800" [label="apply", shape="box", fillcolor="lightgreen"];
  "124329624358400" -> "124329624358800";
  "124329624358880" [label="identifier: sin", shape="box", fillcolor="Honeydew"];
  "124329624358800" -> "124329624358880";
  "124329624358960" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "124329624358800" -> "124329624358960";
  "124329624359040" [label="identifier: B", shape="box", fillcolor="Honeydew"];
  "124329624358960" -> "124329624359040";
}
```

![Image generated for example](images/cosine_addition.png)
