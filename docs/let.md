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
  123937589658944["form"]:::custom_form;
  123937589659024["part: let"]:::custom_part;
  123937589658944 --> 123937589659024;
  123937589659104["operator: ="]:::custom_operator;
  123937589659024 --> 123937589659104;
  123937589659184["identifier: x"]:::custom_identifier;
  123937589659104 --> 123937589659184;
  123937589659264["apply"]:::custom_apply;
  123937589659104 --> 123937589659264;
  123937589659344["identifier: f"]:::custom_identifier;
  123937589659264 --> 123937589659344;
  123937589659424["arguments"]:::custom_arguments;
  123937589659264 --> 123937589659424;
  123937589659504["identifier: a"]:::custom_identifier;
  123937589659424 --> 123937589659504;
  123937589659584["operator: ="]:::custom_operator;
  123937589659024 --> 123937589659584;
  123937589659664["identifier: y"]:::custom_identifier;
  123937589659584 --> 123937589659664;
  123937589659744["apply"]:::custom_apply;
  123937589659584 --> 123937589659744;
  123937589659824["identifier: g"]:::custom_identifier;
  123937589659744 --> 123937589659824;
  123937589659904["arguments"]:::custom_arguments;
  123937589659744 --> 123937589659904;
  123937589659984["identifier: b"]:::custom_identifier;
  123937589659904 --> 123937589659984;
  123937589660064["part: in"]:::custom_part;
  123937589658944 --> 123937589660064;
  123937589660144["delimited"]:::custom_delimited;
  123937589660064 --> 123937589660144;
  123937589660224["identifier: x"]:::custom_identifier;
  123937589660144 --> 123937589660224;
  123937589660304["identifier: y"]:::custom_identifier;
  123937589660144 --> 123937589660304;

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
  "124221782082800" [label="form", shape="box", fillcolor="lightpink"];
  "124221782082880" [label="part: let", shape="box", fillcolor="#FFD8E1"];
  "124221782082800" -> "124221782082880";
  "124221782082960" [label="operator: =", shape="box", fillcolor="#C0FFC0"];
  "124221782082880" -> "124221782082960";
  "124221782083040" [label="identifier: x", shape="box", fillcolor="Honeydew"];
  "124221782082960" -> "124221782083040";
  "124221782083120" [label="apply", shape="box", fillcolor="lightgreen"];
  "124221782082960" -> "124221782083120";
  "124221782083200" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "124221782083120" -> "124221782083200";
  "124221782083280" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "124221782083120" -> "124221782083280";
  "124221782083360" [label="identifier: a", shape="box", fillcolor="Honeydew"];
  "124221782083280" -> "124221782083360";
  "124221782083440" [label="operator: =", shape="box", fillcolor="#C0FFC0"];
  "124221782082880" -> "124221782083440";
  "124221782083520" [label="identifier: y", shape="box", fillcolor="Honeydew"];
  "124221782083440" -> "124221782083520";
  "124221782083600" [label="apply", shape="box", fillcolor="lightgreen"];
  "124221782083440" -> "124221782083600";
  "124221782083680" [label="identifier: g", shape="box", fillcolor="Honeydew"];
  "124221782083600" -> "124221782083680";
  "124221782083760" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "124221782083600" -> "124221782083760";
  "124221782083840" [label="identifier: b", shape="box", fillcolor="Honeydew"];
  "124221782083760" -> "124221782083840";
  "124221782083920" [label="part: in", shape="box", fillcolor="#FFD8E1"];
  "124221782082800" -> "124221782083920";
  "124221782084000" [label="delimited", shape="box", fillcolor="lightgray"];
  "124221782083920" -> "124221782084000";
  "124221782084080" [label="identifier: x", shape="box", fillcolor="Honeydew"];
  "124221782084000" -> "124221782084080";
  "124221782084160" [label="identifier: y", shape="box", fillcolor="Honeydew"];
  "124221782084000" -> "124221782084160";
}
```

![Image generated for example](images/let.png)
