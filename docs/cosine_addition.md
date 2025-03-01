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
  123193826313776["operator: -"]:::custom_operator;
  123193826313856["operator: *"]:::custom_operator;
  123193826313776 --> 123193826313856;
  123193826313936["apply"]:::custom_apply;
  123193826313856 --> 123193826313936;
  123193826314016["identifier: cos"]:::custom_identifier;
  123193826313936 --> 123193826314016;
  123193826314096["arguments"]:::custom_arguments;
  123193826313936 --> 123193826314096;
  123193826314176["identifier: A"]:::custom_identifier;
  123193826314096 --> 123193826314176;
  123193826314256["apply"]:::custom_apply;
  123193826313856 --> 123193826314256;
  123193826314336["identifier: cos"]:::custom_identifier;
  123193826314256 --> 123193826314336;
  123193826314416["arguments"]:::custom_arguments;
  123193826314256 --> 123193826314416;
  123193826314496["identifier: B"]:::custom_identifier;
  123193826314416 --> 123193826314496;
  123193826314576["operator: *"]:::custom_operator;
  123193826313776 --> 123193826314576;
  123193826314656["apply"]:::custom_apply;
  123193826314576 --> 123193826314656;
  123193826314736["identifier: sin"]:::custom_identifier;
  123193826314656 --> 123193826314736;
  123193826314816["arguments"]:::custom_arguments;
  123193826314656 --> 123193826314816;
  123193826314896["identifier: A"]:::custom_identifier;
  123193826314816 --> 123193826314896;
  123193826314976["apply"]:::custom_apply;
  123193826314576 --> 123193826314976;
  123193826315056["identifier: sin"]:::custom_identifier;
  123193826314976 --> 123193826315056;
  123193826315136["arguments"]:::custom_arguments;
  123193826314976 --> 123193826315136;
  123193826315216["identifier: B"]:::custom_identifier;
  123193826315136 --> 123193826315216;

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
  "134437148919344" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "134437148919424" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "134437148919344" -> "134437148919424";
  "134437148919504" [label="apply", shape="box", fillcolor="lightgreen"];
  "134437148919424" -> "134437148919504";
  "134437148919584" [label="identifier: cos", shape="box", fillcolor="Honeydew"];
  "134437148919504" -> "134437148919584";
  "134437148919664" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "134437148919504" -> "134437148919664";
  "134437148919744" [label="identifier: A", shape="box", fillcolor="Honeydew"];
  "134437148919664" -> "134437148919744";
  "134437148919824" [label="apply", shape="box", fillcolor="lightgreen"];
  "134437148919424" -> "134437148919824";
  "134437148919904" [label="identifier: cos", shape="box", fillcolor="Honeydew"];
  "134437148919824" -> "134437148919904";
  "134437148919984" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "134437148919824" -> "134437148919984";
  "134437148920064" [label="identifier: B", shape="box", fillcolor="Honeydew"];
  "134437148919984" -> "134437148920064";
  "134437148920144" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "134437148919344" -> "134437148920144";
  "134437148920224" [label="apply", shape="box", fillcolor="lightgreen"];
  "134437148920144" -> "134437148920224";
  "134437148920304" [label="identifier: sin", shape="box", fillcolor="Honeydew"];
  "134437148920224" -> "134437148920304";
  "134437148920384" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "134437148920224" -> "134437148920384";
  "134437148920464" [label="identifier: A", shape="box", fillcolor="Honeydew"];
  "134437148920384" -> "134437148920464";
  "134437148920544" [label="apply", shape="box", fillcolor="lightgreen"];
  "134437148920144" -> "134437148920544";
  "134437148920624" [label="identifier: sin", shape="box", fillcolor="Honeydew"];
  "134437148920544" -> "134437148920624";
  "134437148920704" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "134437148920544" -> "134437148920704";
  "134437148920784" [label="identifier: B", shape="box", fillcolor="Honeydew"];
  "134437148920704" -> "134437148920784";
}
```

![Image generated for example](images/cosine_addition.png)
