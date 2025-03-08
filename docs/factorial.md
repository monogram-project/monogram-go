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
  139835609729600["form: surround"]:::custom_form;
  139835609729680["part: def"]:::custom_part;
  139835609729600 --> 139835609729680;
  139835609729760["apply"]:::custom_apply;
  139835609729680 --> 139835609729760;
  139835609729840["identifier: f"]:::custom_identifier;
  139835609729760 --> 139835609729840;
  139835609729920["arguments"]:::custom_arguments;
  139835609729760 --> 139835609729920;
  139835609730000["identifier: n"]:::custom_identifier;
  139835609729920 --> 139835609730000;
  139835609730080["part: _"]:::custom_part;
  139835609729600 --> 139835609730080;
  139835609730160["form: surround"]:::custom_form;
  139835609730080 --> 139835609730160;
  139835609730240["part: if"]:::custom_part;
  139835609730160 --> 139835609730240;
  139835609730320["operator: <="]:::custom_operator;
  139835609730240 --> 139835609730320;
  139835609730400["identifier: n"]:::custom_identifier;
  139835609730320 --> 139835609730400;
  139835609730480["number: 1"]:::custom_number;
  139835609730320 --> 139835609730480;
  139835609730560["part: _"]:::custom_part;
  139835609730160 --> 139835609730560;
  139835609730640["number: 1"]:::custom_number;
  139835609730560 --> 139835609730640;
  139835609730720["part: else"]:::custom_part;
  139835609730160 --> 139835609730720;
  139835609730800["operator: *"]:::custom_operator;
  139835609730720 --> 139835609730800;
  139835609730880["identifier: n"]:::custom_identifier;
  139835609730800 --> 139835609730880;
  139835609730960["apply"]:::custom_apply;
  139835609730800 --> 139835609730960;
  139835609731120["identifier: f"]:::custom_identifier;
  139835609730960 --> 139835609731120;
  139835609731280["arguments"]:::custom_arguments;
  139835609730960 --> 139835609731280;
  139835609731440["operator: -"]:::custom_operator;
  139835609731280 --> 139835609731440;
  139835609731600["identifier: n"]:::custom_identifier;
  139835609731440 --> 139835609731600;
  139835609731760["number: 1"]:::custom_number;
  139835609731440 --> 139835609731760;

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
  "129493713095232" [label="form: surround", shape="box", fillcolor="lightpink"];
  "129493713095312" [label="part: def", shape="box", fillcolor="#FFD8E1"];
  "129493713095232" -> "129493713095312";
  "129493713095392" [label="apply", shape="box", fillcolor="lightgreen"];
  "129493713095312" -> "129493713095392";
  "129493713095472" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "129493713095392" -> "129493713095472";
  "129493713095552" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "129493713095392" -> "129493713095552";
  "129493713095632" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "129493713095552" -> "129493713095632";
  "129493713095712" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "129493713095232" -> "129493713095712";
  "129493713095792" [label="form: surround", shape="box", fillcolor="lightpink"];
  "129493713095712" -> "129493713095792";
  "129493713095872" [label="part: if", shape="box", fillcolor="#FFD8E1"];
  "129493713095792" -> "129493713095872";
  "129493713095952" [label="operator: <=", shape="box", fillcolor="#C0FFC0"];
  "129493713095872" -> "129493713095952";
  "129493713096032" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "129493713095952" -> "129493713096032";
  "129493713096112" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "129493713095952" -> "129493713096112";
  "129493713096192" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "129493713095792" -> "129493713096192";
  "129493713096272" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "129493713096192" -> "129493713096272";
  "129493713096352" [label="part: else", shape="box", fillcolor="#FFD8E1"];
  "129493713095792" -> "129493713096352";
  "129493713096432" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "129493713096352" -> "129493713096432";
  "129493713096512" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "129493713096432" -> "129493713096512";
  "129493713096592" [label="apply", shape="box", fillcolor="lightgreen"];
  "129493713096432" -> "129493713096592";
  "129493713096752" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "129493713096592" -> "129493713096752";
  "129493713096912" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "129493713096592" -> "129493713096912";
  "129493713097072" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "129493713096912" -> "129493713097072";
  "129493713097232" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "129493713097072" -> "129493713097232";
  "129493713097392" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "129493713097072" -> "129493713097392";
}
```

![Image generated for example](images/factorial.png)
