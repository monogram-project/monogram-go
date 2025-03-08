# Recursive factorial function

## Monogram

```txt
# The factorial function in monogram.
def f(n):
    if n <= 1:
        1
    else:
        n * f(n - 1)
    endif
enddef
```

## Mermaid diagram

We can target Mermaid's flowchart as an output format. 
And this is what it looks like:

```mermaid
graph LR
  137321037236880["form: surround"]:::custom_form;
  137321037236960["part: def"]:::custom_part;
  137321037236880 --> 137321037236960;
  137321037237040["apply"]:::custom_apply;
  137321037236960 --> 137321037237040;
  137321037237120["identifier: f"]:::custom_identifier;
  137321037237040 --> 137321037237120;
  137321037237200["arguments"]:::custom_arguments;
  137321037237040 --> 137321037237200;
  137321037237280["identifier: n"]:::custom_identifier;
  137321037237200 --> 137321037237280;
  137321037237360["part: _"]:::custom_part;
  137321037236880 --> 137321037237360;
  137321037237440["form: surround"]:::custom_form;
  137321037237360 --> 137321037237440;
  137321037237520["part: if"]:::custom_part;
  137321037237440 --> 137321037237520;
  137321037237600["operator: <="]:::custom_operator;
  137321037237520 --> 137321037237600;
  137321037237680["identifier: n"]:::custom_identifier;
  137321037237600 --> 137321037237680;
  137321037237760["number: 1"]:::custom_number;
  137321037237600 --> 137321037237760;
  137321037237840["part: _"]:::custom_part;
  137321037237440 --> 137321037237840;
  137321037237920["number: 1"]:::custom_number;
  137321037237840 --> 137321037237920;
  137321037238000["part: else"]:::custom_part;
  137321037237440 --> 137321037238000;
  137321037238080["operator: *"]:::custom_operator;
  137321037238000 --> 137321037238080;
  137321037238160["identifier: n"]:::custom_identifier;
  137321037238080 --> 137321037238160;
  137321037238240["apply"]:::custom_apply;
  137321037238080 --> 137321037238240;
  137321037238400["identifier: f"]:::custom_identifier;
  137321037238240 --> 137321037238400;
  137321037238560["arguments"]:::custom_arguments;
  137321037238240 --> 137321037238560;
  137321037238720["operator: -"]:::custom_operator;
  137321037238560 --> 137321037238720;
  137321037238880["identifier: n"]:::custom_identifier;
  137321037238720 --> 137321037238880;
  137321037239040["number: 1"]:::custom_number;
  137321037238720 --> 137321037239040;

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
<form syntax="surround">
    <part keyword="def">
        <apply kind="parentheses" separator="undefined">
            <identifier name="f"/>
            <arguments>
                <identifier name="n"/>
            </arguments>
        </apply>
    </part>
    <part keyword="_">
        <form syntax="surround">
            <part keyword="if">
                <operator name="&lt;=">
                    <identifier name="n"/>
                    <number value="1"/>
                </operator>
            </part>
            <part keyword="_">
                <number value="1"/>
            </part>
            <part keyword="else">
                <operator name="*">
                    <identifier name="n"/>
                    <apply kind="parentheses" separator="undefined">
                        <identifier name="f"/>
                        <arguments>
                            <operator name="-">
                                <identifier name="n"/>
                                <number value="1"/>
                            </operator>
                        </arguments>
                    </apply>
                </operator>
            </part>
        </form>
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
    "syntax": "surround",
    "children": [
        {
            "role": "part",
            "keyword": "def",
            "children": [
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
                                    "name": "n"
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "role": "part",
            "keyword": "_",
            "children": [
                {
                    "role": "form",
                    "syntax": "surround",
                    "children": [
                        {
                            "role": "part",
                            "keyword": "if",
                            "children": [
                                {
                                    "role": "operator",
                                    "name": "<=",
                                    "children": [
                                        {
                                            "role": "identifier",
                                            "name": "n"
                                        },
                                        {
                                            "role": "number",
                                            "value": 1
                                        }
                                    ]
                                }
                            ]
                        },
                        {
                            "role": "part",
                            "keyword": "_",
                            "children": [
                                {
                                    "role": "number",
                                    "value": 1
                                }
                            ]
                        },
                        {
                            "role": "part",
                            "keyword": "else",
                            "children": [
                                {
                                    "role": "operator",
                                    "name": "*",
                                    "children": [
                                        {
                                            "role": "identifier",
                                            "name": "n"
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
                                                            "role": "operator",
                                                            "name": "-",
                                                            "children": [
                                                                {
                                                                    "role": "identifier",
                                                                    "name": "n"
                                                                },
                                                                {
                                                                    "role": "number",
                                                                    "value": 1
                                                                }
                                                            ]
                                                        }
                                                    ]
                                                }
                                            ]
                                        }
                                    ]
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
role: form
syntax: surround
children:
- role: part
  keyword: def
  children:
  - role: apply
    kind: parentheses
    separator: undefined
    children:
    - role: identifier
      name: f
    - role: arguments
      children:
      - role: identifier
        name: n
- role: part
  keyword: _
  children:
  - role: form
    syntax: surround
    children:
    - role: part
      keyword: if
      children:
      - role: operator
        name: <=
        children:
        - role: identifier
          name: n
        - role: number
          value: 1
    - role: part
      keyword: _
      children:
      - role: number
        value: 1
    - role: part
      keyword: else
      children:
      - role: operator
        name: '*'
        children:
        - role: identifier
          name: n
        - role: apply
          kind: parentheses
          separator: undefined
          children:
          - role: identifier
            name: f
          - role: arguments
            children:
            - role: operator
              name: '-'
              children:
              - role: identifier
                name: n
              - role: number
                value: 1

```

## Graphviz Dot format

```dot
digraph G {
  bgcolor="transparent";
  node [shape="box", style="filled", fontname="Ubuntu Mono"];
  "137831330941504" [label="form: surround", shape="box", fillcolor="lightpink"];
  "137831330941584" [label="part: def", shape="box", fillcolor="#FFD8E1"];
  "137831330941504" -> "137831330941584";
  "137831330941664" [label="apply", shape="box", fillcolor="lightgreen"];
  "137831330941584" -> "137831330941664";
  "137831330941744" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "137831330941664" -> "137831330941744";
  "137831330941824" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "137831330941664" -> "137831330941824";
  "137831330941904" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "137831330941824" -> "137831330941904";
  "137831330941984" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "137831330941504" -> "137831330941984";
  "137831330942064" [label="form: surround", shape="box", fillcolor="lightpink"];
  "137831330941984" -> "137831330942064";
  "137831330942144" [label="part: if", shape="box", fillcolor="#FFD8E1"];
  "137831330942064" -> "137831330942144";
  "137831330942224" [label="operator: <=", shape="box", fillcolor="#C0FFC0"];
  "137831330942144" -> "137831330942224";
  "137831330942304" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "137831330942224" -> "137831330942304";
  "137831330942384" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "137831330942224" -> "137831330942384";
  "137831330942464" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "137831330942064" -> "137831330942464";
  "137831330942544" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "137831330942464" -> "137831330942544";
  "137831330942624" [label="part: else", shape="box", fillcolor="#FFD8E1"];
  "137831330942064" -> "137831330942624";
  "137831330942704" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "137831330942624" -> "137831330942704";
  "137831330942784" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "137831330942704" -> "137831330942784";
  "137831330942864" [label="apply", shape="box", fillcolor="lightgreen"];
  "137831330942704" -> "137831330942864";
  "137831330943024" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "137831330942864" -> "137831330943024";
  "137831330943184" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "137831330942864" -> "137831330943184";
  "137831330943344" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "137831330943184" -> "137831330943344";
  "137831330943504" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "137831330943344" -> "137831330943504";
  "137831330943664" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "137831330943344" -> "137831330943664";
}
```

![Image generated for example](images/factorial.png)
