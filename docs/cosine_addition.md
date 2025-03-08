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
  123601159736112["operator: -"]:::custom_operator;
  123601159736192["operator: *"]:::custom_operator;
  123601159736112 --> 123601159736192;
  123601159736272["apply"]:::custom_apply;
  123601159736192 --> 123601159736272;
  123601159736352["identifier: cos"]:::custom_identifier;
  123601159736272 --> 123601159736352;
  123601159736432["arguments"]:::custom_arguments;
  123601159736272 --> 123601159736432;
  123601159736512["identifier: A"]:::custom_identifier;
  123601159736432 --> 123601159736512;
  123601159736592["apply"]:::custom_apply;
  123601159736192 --> 123601159736592;
  123601159736672["identifier: cos"]:::custom_identifier;
  123601159736592 --> 123601159736672;
  123601159736752["arguments"]:::custom_arguments;
  123601159736592 --> 123601159736752;
  123601159736832["identifier: B"]:::custom_identifier;
  123601159736752 --> 123601159736832;
  123601159736912["operator: *"]:::custom_operator;
  123601159736112 --> 123601159736912;
  123601159736992["apply"]:::custom_apply;
  123601159736912 --> 123601159736992;
  123601159737072["identifier: sin"]:::custom_identifier;
  123601159736992 --> 123601159737072;
  123601159737152["arguments"]:::custom_arguments;
  123601159736992 --> 123601159737152;
  123601159737232["identifier: A"]:::custom_identifier;
  123601159737152 --> 123601159737232;
  123601159737312["apply"]:::custom_apply;
  123601159736912 --> 123601159737312;
  123601159737392["identifier: sin"]:::custom_identifier;
  123601159737312 --> 123601159737392;
  123601159737472["arguments"]:::custom_arguments;
  123601159737312 --> 123601159737472;
  123601159737552["identifier: B"]:::custom_identifier;
  123601159737472 --> 123601159737552;

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
  "128849475111648" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "128849475111728" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "128849475111648" -> "128849475111728";
  "128849475111808" [label="apply", shape="box", fillcolor="lightgreen"];
  "128849475111728" -> "128849475111808";
  "128849475111888" [label="identifier: cos", shape="box", fillcolor="Honeydew"];
  "128849475111808" -> "128849475111888";
  "128849475111968" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "128849475111808" -> "128849475111968";
  "128849475112048" [label="identifier: A", shape="box", fillcolor="Honeydew"];
  "128849475111968" -> "128849475112048";
  "128849475112128" [label="apply", shape="box", fillcolor="lightgreen"];
  "128849475111728" -> "128849475112128";
  "128849475112208" [label="identifier: cos", shape="box", fillcolor="Honeydew"];
  "128849475112128" -> "128849475112208";
  "128849475112288" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "128849475112128" -> "128849475112288";
  "128849475112368" [label="identifier: B", shape="box", fillcolor="Honeydew"];
  "128849475112288" -> "128849475112368";
  "128849475112448" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "128849475111648" -> "128849475112448";
  "128849475112528" [label="apply", shape="box", fillcolor="lightgreen"];
  "128849475112448" -> "128849475112528";
  "128849475112608" [label="identifier: sin", shape="box", fillcolor="Honeydew"];
  "128849475112528" -> "128849475112608";
  "128849475112688" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "128849475112528" -> "128849475112688";
  "128849475112768" [label="identifier: A", shape="box", fillcolor="Honeydew"];
  "128849475112688" -> "128849475112768";
  "128849475112848" [label="apply", shape="box", fillcolor="lightgreen"];
  "128849475112448" -> "128849475112848";
  "128849475112928" [label="identifier: sin", shape="box", fillcolor="Honeydew"];
  "128849475112848" -> "128849475112928";
  "128849475113008" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "128849475112848" -> "128849475113008";
  "128849475113088" [label="identifier: B", shape="box", fillcolor="Honeydew"];
  "128849475113008" -> "128849475113088";
}
```

![Image generated for example](images/cosine_addition.png)
