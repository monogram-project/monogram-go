# Let Expression


## Monogram

This example shows how you might simulate `let` expressions in Monogram:

```txt
let x = f(a)
    y = g(b)
in:
    (x, y)
endlet

```

## Mermaid diagram

We can target Mermaid's flowchart as an output format. 
And this is what it looks like:

```mermaid
graph LR
  124960933774576["form"]:::custom_form;
  124960933774656["part: let"]:::custom_part;
  124960933774576 --> 124960933774656;
  124960933774736["operator: ="]:::custom_operator;
  124960933774656 --> 124960933774736;
  124960933774816["identifier: x"]:::custom_identifier;
  124960933774736 --> 124960933774816;
  124960933774896["apply"]:::custom_apply;
  124960933774736 --> 124960933774896;
  124960933774976["identifier: f"]:::custom_identifier;
  124960933774896 --> 124960933774976;
  124960933775056["arguments"]:::custom_arguments;
  124960933774896 --> 124960933775056;
  124960933775136["identifier: a"]:::custom_identifier;
  124960933775056 --> 124960933775136;
  124960933775216["operator: ="]:::custom_operator;
  124960933774656 --> 124960933775216;
  124960933775296["identifier: y"]:::custom_identifier;
  124960933775216 --> 124960933775296;
  124960933775376["apply"]:::custom_apply;
  124960933775216 --> 124960933775376;
  124960933775456["identifier: g"]:::custom_identifier;
  124960933775376 --> 124960933775456;
  124960933775536["arguments"]:::custom_arguments;
  124960933775376 --> 124960933775536;
  124960933775616["identifier: b"]:::custom_identifier;
  124960933775536 --> 124960933775616;
  124960933775696["part: in"]:::custom_part;
  124960933774576 --> 124960933775696;
  124960933775776["delimited"]:::custom_delimited;
  124960933775696 --> 124960933775776;
  124960933775856["identifier: x"]:::custom_identifier;
  124960933775776 --> 124960933775856;
  124960933775936["identifier: y"]:::custom_identifier;
  124960933775776 --> 124960933775936;

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
<form>
    <part keyword="let">
        <operator name="=">
            <identifier name="x"/>
            <apply kind="parentheses" separator="undefined">
                <identifier name="f"/>
                <arguments>
                    <identifier name="a"/>
                </arguments>
            </apply>
        </operator>
        <operator name="=">
            <identifier name="y"/>
            <apply kind="parentheses" separator="undefined">
                <identifier name="g"/>
                <arguments>
                    <identifier name="b"/>
                </arguments>
            </apply>
        </operator>
    </part>
    <part keyword="in">
        <delimited kind="parentheses" separator="comma">
            <identifier name="x"/>
            <identifier name="y"/>
        </delimited>
    </part>
</form>
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
    "role": "form",
    "children": [
        {
            "role": "part",
            "keyword": "let",
            "children": [
                {
                    "role": "operator",
                    "name": "=",
                    "children": [
                        {
                            "role": "identifier",
                            "name": "x"
                        },
                        {
                            "role": "apply",
                            "kind": "parentheses",
                            "separator": "undefined",
                            "children": [
                                {
                                    "role": "identifier",
                                    "name": "f"
                                },
                                {
                                    "role": "arguments",
                                    "children": [
                                        {
                                            "role": "identifier",
                                            "name": "a"
                                        }
                                    ]
                                }
                            ]
                        }
                    ]
                },
                {
                    "role": "operator",
                    "name": "=",
                    "children": [
                        {
                            "role": "identifier",
                            "name": "y"
                        },
                        {
                            "role": "apply",
                            "kind": "parentheses",
                            "separator": "undefined",
                            "children": [
                                {
                                    "role": "identifier",
                                    "name": "g"
                                },
                                {
                                    "role": "arguments",
                                    "children": [
                                        {
                                            "role": "identifier",
                                            "name": "b"
                                        }
                                    ]
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "role": "part",
            "keyword": "in",
            "children": [
                {
                    "role": "delimited",
                    "kind": "parentheses",
                    "separator": "comma",
                    "children": [
                        {
                            "role": "identifier",
                            "name": "x"
                        },
                        {
                            "role": "identifier",
                            "name": "y"
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
role: form
children:
- role: part
  keyword: let
  children:
  - role: operator
    name: '='
    children:
    - role: identifier
      name: x
    - role: apply
      kind: parentheses
      separator: undefined
      children:
      - role: identifier
        name: f
      - role: arguments
        children:
        - role: identifier
          name: a
  - role: operator
    name: '='
    children:
    - role: identifier
      name: y
    - role: apply
      kind: parentheses
      separator: undefined
      children:
      - role: identifier
        name: g
      - role: arguments
        children:
        - role: identifier
          name: b
- role: part
  keyword: in
  children:
  - role: delimited
    kind: parentheses
    separator: comma
    children:
    - role: identifier
      name: x
    - role: identifier
      name: y

```

## Graphviz Dot format

```dot
digraph G {
  bgcolor="transparent";
  node [shape="box", style="filled", fontname="Ubuntu Mono"];
  "126647206511936" [label="form", shape="box", fillcolor="lightpink"];
  "126647206512016" [label="part: let", shape="box", fillcolor="#FFD8E1"];
  "126647206511936" -> "126647206512016";
  "126647206512096" [label="operator: =", shape="box", fillcolor="#C0FFC0"];
  "126647206512016" -> "126647206512096";
  "126647206512176" [label="identifier: x", shape="box", fillcolor="Honeydew"];
  "126647206512096" -> "126647206512176";
  "126647206512256" [label="apply", shape="box", fillcolor="lightgreen"];
  "126647206512096" -> "126647206512256";
  "126647206512336" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "126647206512256" -> "126647206512336";
  "126647206512416" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "126647206512256" -> "126647206512416";
  "126647206512496" [label="identifier: a", shape="box", fillcolor="Honeydew"];
  "126647206512416" -> "126647206512496";
  "126647206512576" [label="operator: =", shape="box", fillcolor="#C0FFC0"];
  "126647206512016" -> "126647206512576";
  "126647206512656" [label="identifier: y", shape="box", fillcolor="Honeydew"];
  "126647206512576" -> "126647206512656";
  "126647206512736" [label="apply", shape="box", fillcolor="lightgreen"];
  "126647206512576" -> "126647206512736";
  "126647206512816" [label="identifier: g", shape="box", fillcolor="Honeydew"];
  "126647206512736" -> "126647206512816";
  "126647206512896" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "126647206512736" -> "126647206512896";
  "126647206512976" [label="identifier: b", shape="box", fillcolor="Honeydew"];
  "126647206512896" -> "126647206512976";
  "126647206513056" [label="part: in", shape="box", fillcolor="#FFD8E1"];
  "126647206511936" -> "126647206513056";
  "126647206513136" [label="delimited", shape="box", fillcolor="lightgray"];
  "126647206513056" -> "126647206513136";
  "126647206513216" [label="identifier: x", shape="box", fillcolor="Honeydew"];
  "126647206513136" -> "126647206513216";
  "126647206513296" [label="identifier: y", shape="box", fillcolor="Honeydew"];
  "126647206513136" -> "126647206513296";
}
```

![Image generated for example](images/factorial.png)
