# JSON Expression

## Monogram

Monograph supports JSON expressions:

```txt
{
  "person": {
    "name": "Alice",
    "age": 25,
    "isStudent": true,
    "skills": ["Python", "JavaScript", "SQL"],
    "address": {
      "street": "123 Maple Street",
      "city": "Exampleville",
      "country": "Neverland"
    },
    "favoriteBooks": [
      {
        "title": "To Kill a Mockingbird",
        "author": "Harper Lee",
        "yearPublished": 1960
      },
      {
        "title": "1984",
        "author": "George Orwell",
        "yearPublished": 1949
      }
    ]
  }
}
```

## Mermaid diagram

We can target Mermaid's flowchart as an output format. 
And this is what it looks like:

```mermaid
graph LR
  134426219842112["delimited"]:::custom_delimited;
  134426219842192["operator: :"]:::custom_operator;
  134426219842112 --> 134426219842192;
  134426219842272["string: person"]:::custom_string;
  134426219842192 --> 134426219842272;
  134426219842352["delimited"]:::custom_delimited;
  134426219842192 --> 134426219842352;
  134426219842432["operator: :"]:::custom_operator;
  134426219842352 --> 134426219842432;
  134426219842512["string: name"]:::custom_string;
  134426219842432 --> 134426219842512;
  134426219842592["string: Alice"]:::custom_string;
  134426219842432 --> 134426219842592;
  134426219842672["operator: :"]:::custom_operator;
  134426219842352 --> 134426219842672;
  134426219842752["string: age"]:::custom_string;
  134426219842672 --> 134426219842752;
  134426219842832["number: 25"]:::custom_number;
  134426219842672 --> 134426219842832;
  134426219842912["operator: :"]:::custom_operator;
  134426219842352 --> 134426219842912;
  134426219842992["string: isStudent"]:::custom_string;
  134426219842912 --> 134426219842992;
  134426219843072["identifier: true"]:::custom_identifier;
  134426219842912 --> 134426219843072;
  134426219843152["operator: :"]:::custom_operator;
  134426219842352 --> 134426219843152;
  134426219843232["string: skills"]:::custom_string;
  134426219843152 --> 134426219843232;
  134426219843312["delimited"]:::custom_delimited;
  134426219843152 --> 134426219843312;
  134426219843392["string: Python"]:::custom_string;
  134426219843312 --> 134426219843392;
  134426219843472["string: JavaScript"]:::custom_string;
  134426219843312 --> 134426219843472;
  134426219843552["string: SQL"]:::custom_string;
  134426219843312 --> 134426219843552;
  134426219843632["operator: :"]:::custom_operator;
  134426219842352 --> 134426219843632;
  134426219843712["string: address"]:::custom_string;
  134426219843632 --> 134426219843712;
  134426219843792["delimited"]:::custom_delimited;
  134426219843632 --> 134426219843792;
  134426219843872["operator: :"]:::custom_operator;
  134426219843792 --> 134426219843872;
  134426219844032["string: street"]:::custom_string;
  134426219843872 --> 134426219844032;
  134426219844192["string:<br/>123 Maple Street"]:::custom_string;
  134426219843872 --> 134426219844192;
  134426219844272["operator: :"]:::custom_operator;
  134426219843792 --> 134426219844272;
  134426219844432["string: city"]:::custom_string;
  134426219844272 --> 134426219844432;
  134426219844592["string: Exampleville"]:::custom_string;
  134426219844272 --> 134426219844592;
  134426219844672["operator: :"]:::custom_operator;
  134426219843792 --> 134426219844672;
  134426219844832["string: country"]:::custom_string;
  134426219844672 --> 134426219844832;
  134426219844992["string: Neverland"]:::custom_string;
  134426219844672 --> 134426219844992;
  134426219845072["operator: :"]:::custom_operator;
  134426219842352 --> 134426219845072;
  134426219845152["string:<br/>favoriteBooks"]:::custom_string;
  134426219845072 --> 134426219845152;
  134426219845232["delimited"]:::custom_delimited;
  134426219845072 --> 134426219845232;
  134426219845312["delimited"]:::custom_delimited;
  134426219845232 --> 134426219845312;
  134426219845472["operator: :"]:::custom_operator;
  134426219845312 --> 134426219845472;
  134426219845632["string: title"]:::custom_string;
  134426219845472 --> 134426219845632;
  134426219845792["string:<br/>To Kill a Mockingbird"]:::custom_string;
  134426219845472 --> 134426219845792;
  134426219846032["operator: :"]:::custom_operator;
  134426219845312 --> 134426219846032;
  134426219846192["string: author"]:::custom_string;
  134426219846032 --> 134426219846192;
  134426219846352["string: Harper Lee"]:::custom_string;
  134426219846032 --> 134426219846352;
  134426219846592["operator: :"]:::custom_operator;
  134426219845312 --> 134426219846592;
  134426219846752["string:<br/>yearPublished"]:::custom_string;
  134426219846592 --> 134426219846752;
  134426219846912["number: 1960"]:::custom_number;
  134426219846592 --> 134426219846912;
  134426219847072["delimited"]:::custom_delimited;
  134426219845232 --> 134426219847072;
  134426219847232["operator: :"]:::custom_operator;
  134426219847072 --> 134426219847232;
  134426219847392["string: title"]:::custom_string;
  134426219847232 --> 134426219847392;
  134426219847552["string: 1984"]:::custom_string;
  134426219847232 --> 134426219847552;
  134426219847792["operator: :"]:::custom_operator;
  134426219847072 --> 134426219847792;
  134426219847952["string: author"]:::custom_string;
  134426219847792 --> 134426219847952;
  134426219848112["string:<br/>George Orwell"]:::custom_string;
  134426219847792 --> 134426219848112;
  134426219848352["operator: :"]:::custom_operator;
  134426219847072 --> 134426219848352;
  134426219848512["string:<br/>yearPublished"]:::custom_string;
  134426219848352 --> 134426219848512;
  134426219848672["number: 1949"]:::custom_number;
  134426219848352 --> 134426219848672;

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
<delimited kind="braces" separator="undefined">
    <operator name=":" syntax="infix">
        <string quote=":" value="person"/>
        <delimited kind="braces" separator="comma">
            <operator name=":" syntax="infix">
                <string quote=":" value="name"/>
                <string quote="," value="Alice"/>
            </operator>
            <operator name=":" syntax="infix">
                <string quote=":" value="age"/>
                <number value="25"/>
            </operator>
            <operator name=":" syntax="infix">
                <string quote=":" value="isStudent"/>
                <identifier name="true"/>
            </operator>
            <operator name=":" syntax="infix">
                <string quote=":" value="skills"/>
                <delimited kind="brackets" separator="comma">
                    <string quote="," value="Python"/>
                    <string quote="," value="JavaScript"/>
                    <string quote="]" value="SQL"/>
                </delimited>
            </operator>
            <operator name=":" syntax="infix">
                <string quote=":" value="address"/>
                <delimited kind="braces" separator="comma">
                    <operator name=":" syntax="infix">
                        <string quote=":" value="street"/>
                        <string quote="," value="123 Maple Street"/>
                    </operator>
                    <operator name=":" syntax="infix">
                        <string quote=":" value="city"/>
                        <string quote="," value="Exampleville"/>
                    </operator>
                    <operator name=":" syntax="infix">
                        <string quote=":" value="country"/>
                        <string quote="&quot;" value="Neverland"/>
                    </operator>
                </delimited>
            </operator>
            <operator name=":" syntax="infix">
                <string quote=":" value="favoriteBooks"/>
                <delimited kind="brackets" separator="comma">
                    <delimited kind="braces" separator="comma">
                        <operator name=":" syntax="infix">
                            <string quote=":" value="title"/>
                            <string quote="," value="To Kill a Mockingbird"/>
                        </operator>
                        <operator name=":" syntax="infix">
                            <string quote=":" value="author"/>
                            <string quote="," value="Harper Lee"/>
                        </operator>
                        <operator name=":" syntax="infix">
                            <string quote=":" value="yearPublished"/>
                            <number value="1960"/>
                        </operator>
                    </delimited>
                    <delimited kind="braces" separator="comma">
                        <operator name=":" syntax="infix">
                            <string quote=":" value="title"/>
                            <string quote="," value="1984"/>
                        </operator>
                        <operator name=":" syntax="infix">
                            <string quote=":" value="author"/>
                            <string quote="," value="George Orwell"/>
                        </operator>
                        <operator name=":" syntax="infix">
                            <string quote=":" value="yearPublished"/>
                            <number value="1949"/>
                        </operator>
                    </delimited>
                </delimited>
            </operator>
        </delimited>
    </operator>
</delimited>
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
    "role": "delimited",
    "kind": "braces",
    "separator": "undefined",
    "children": [
        {
            "role": "operator",
            "name": ":",
            "syntax": "infix",
            "children": [
                {
                    "role": "string",
                    "quote": ":",
                    "value": "person"
                },
                {
                    "role": "delimited",
                    "kind": "braces",
                    "separator": "comma",
                    "children": [
                        {
                            "role": "operator",
                            "name": ":",
                            "syntax": "infix",
                            "children": [
                                {
                                    "role": "string",
                                    "quote": ":",
                                    "value": "name"
                                },
                                {
                                    "role": "string",
                                    "quote": ",",
                                    "value": "Alice"
                                }
                            ]
                        },
                        {
                            "role": "operator",
                            "name": ":",
                            "syntax": "infix",
                            "children": [
                                {
                                    "role": "string",
                                    "quote": ":",
                                    "value": "age"
                                },
                                {
                                    "role": "number",
                                    "value": 25
                                }
                            ]
                        },
                        {
                            "role": "operator",
                            "name": ":",
                            "syntax": "infix",
                            "children": [
                                {
                                    "role": "string",
                                    "quote": ":",
                                    "value": "isStudent"
                                },
                                {
                                    "role": "identifier",
                                    "name": "true"
                                }
                            ]
                        },
                        {
                            "role": "operator",
                            "name": ":",
                            "syntax": "infix",
                            "children": [
                                {
                                    "role": "string",
                                    "quote": ":",
                                    "value": "skills"
                                },
                                {
                                    "role": "delimited",
                                    "kind": "brackets",
                                    "separator": "comma",
                                    "children": [
                                        {
                                            "role": "string",
                                            "quote": ",",
                                            "value": "Python"
                                        },
                                        {
                                            "role": "string",
                                            "quote": ",",
                                            "value": "JavaScript"
                                        },
                                        {
                                            "role": "string",
                                            "quote": "]",
                                            "value": "SQL"
                                        }
                                    ]
                                }
                            ]
                        },
                        {
                            "role": "operator",
                            "name": ":",
                            "syntax": "infix",
                            "children": [
                                {
                                    "role": "string",
                                    "quote": ":",
                                    "value": "address"
                                },
                                {
                                    "role": "delimited",
                                    "kind": "braces",
                                    "separator": "comma",
                                    "children": [
                                        {
                                            "role": "operator",
                                            "name": ":",
                                            "syntax": "infix",
                                            "children": [
                                                {
                                                    "role": "string",
                                                    "quote": ":",
                                                    "value": "street"
                                                },
                                                {
                                                    "role": "string",
                                                    "quote": ",",
                                                    "value": "123 Maple Street"
                                                }
                                            ]
                                        },
                                        {
                                            "role": "operator",
                                            "name": ":",
                                            "syntax": "infix",
                                            "children": [
                                                {
                                                    "role": "string",
                                                    "quote": ":",
                                                    "value": "city"
                                                },
                                                {
                                                    "role": "string",
                                                    "quote": ",",
                                                    "value": "Exampleville"
                                                }
                                            ]
                                        },
                                        {
                                            "role": "operator",
                                            "name": ":",
                                            "syntax": "infix",
                                            "children": [
                                                {
                                                    "role": "string",
                                                    "quote": ":",
                                                    "value": "country"
                                                },
                                                {
                                                    "role": "string",
                                                    "quote": "\"",
                                                    "value": "Neverland"
                                                }
                                            ]
                                        }
                                    ]
                                }
                            ]
                        },
                        {
                            "role": "operator",
                            "name": ":",
                            "syntax": "infix",
                            "children": [
                                {
                                    "role": "string",
                                    "quote": ":",
                                    "value": "favoriteBooks"
                                },
                                {
                                    "role": "delimited",
                                    "kind": "brackets",
                                    "separator": "comma",
                                    "children": [
                                        {
                                            "role": "delimited",
                                            "kind": "braces",
                                            "separator": "comma",
                                            "children": [
                                                {
                                                    "role": "operator",
                                                    "name": ":",
                                                    "syntax": "infix",
                                                    "children": [
                                                        {
                                                            "role": "string",
                                                            "quote": ":",
                                                            "value": "title"
                                                        },
                                                        {
                                                            "role": "string",
                                                            "quote": ",",
                                                            "value": "To Kill a Mockingbird"
                                                        }
                                                    ]
                                                },
                                                {
                                                    "role": "operator",
                                                    "name": ":",
                                                    "syntax": "infix",
                                                    "children": [
                                                        {
                                                            "role": "string",
                                                            "quote": ":",
                                                            "value": "author"
                                                        },
                                                        {
                                                            "role": "string",
                                                            "quote": ",",
                                                            "value": "Harper Lee"
                                                        }
                                                    ]
                                                },
                                                {
                                                    "role": "operator",
                                                    "name": ":",
                                                    "syntax": "infix",
                                                    "children": [
                                                        {
                                                            "role": "string",
                                                            "quote": ":",
                                                            "value": "yearPublished"
                                                        },
                                                        {
                                                            "role": "number",
                                                            "value": 1960
                                                        }
                                                    ]
                                                }
                                            ]
                                        },
                                        {
                                            "role": "delimited",
                                            "kind": "braces",
                                            "separator": "comma",
                                            "children": [
                                                {
                                                    "role": "operator",
                                                    "name": ":",
                                                    "syntax": "infix",
                                                    "children": [
                                                        {
                                                            "role": "string",
                                                            "quote": ":",
                                                            "value": "title"
                                                        },
                                                        {
                                                            "role": "string",
                                                            "quote": ",",
                                                            "value": "1984"
                                                        }
                                                    ]
                                                },
                                                {
                                                    "role": "operator",
                                                    "name": ":",
                                                    "syntax": "infix",
                                                    "children": [
                                                        {
                                                            "role": "string",
                                                            "quote": ":",
                                                            "value": "author"
                                                        },
                                                        {
                                                            "role": "string",
                                                            "quote": ",",
                                                            "value": "George Orwell"
                                                        }
                                                    ]
                                                },
                                                {
                                                    "role": "operator",
                                                    "name": ":",
                                                    "syntax": "infix",
                                                    "children": [
                                                        {
                                                            "role": "string",
                                                            "quote": ":",
                                                            "value": "yearPublished"
                                                        },
                                                        {
                                                            "role": "number",
                                                            "value": 1949
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
role: delimited
kind: braces
separator: undefined
children:
- role: operator
  name: ':'
  syntax: infix
  children:
  - role: string
    quote: ':'
    value: person
  - role: delimited
    kind: braces
    separator: comma
    children:
    - role: operator
      name: ':'
      syntax: infix
      children:
      - role: string
        quote: ':'
        value: name
      - role: string
        quote: ','
        value: Alice
    - role: operator
      name: ':'
      syntax: infix
      children:
      - role: string
        quote: ':'
        value: age
      - role: number
        value: 25
    - role: operator
      name: ':'
      syntax: infix
      children:
      - role: string
        quote: ':'
        value: isStudent
      - role: identifier
        name: 'true'
    - role: operator
      name: ':'
      syntax: infix
      children:
      - role: string
        quote: ':'
        value: skills
      - role: delimited
        kind: brackets
        separator: comma
        children:
        - role: string
          quote: ','
          value: Python
        - role: string
          quote: ','
          value: JavaScript
        - role: string
          quote: ']'
          value: SQL
    - role: operator
      name: ':'
      syntax: infix
      children:
      - role: string
        quote: ':'
        value: address
      - role: delimited
        kind: braces
        separator: comma
        children:
        - role: operator
          name: ':'
          syntax: infix
          children:
          - role: string
            quote: ':'
            value: street
          - role: string
            quote: ','
            value: 123 Maple Street
        - role: operator
          name: ':'
          syntax: infix
          children:
          - role: string
            quote: ':'
            value: city
          - role: string
            quote: ','
            value: Exampleville
        - role: operator
          name: ':'
          syntax: infix
          children:
          - role: string
            quote: ':'
            value: country
          - role: string
            quote: '"'
            value: Neverland
    - role: operator
      name: ':'
      syntax: infix
      children:
      - role: string
        quote: ':'
        value: favoriteBooks
      - role: delimited
        kind: brackets
        separator: comma
        children:
        - role: delimited
          kind: braces
          separator: comma
          children:
          - role: operator
            name: ':'
            syntax: infix
            children:
            - role: string
              quote: ':'
              value: title
            - role: string
              quote: ','
              value: To Kill a Mockingbird
          - role: operator
            name: ':'
            syntax: infix
            children:
            - role: string
              quote: ':'
              value: author
            - role: string
              quote: ','
              value: Harper Lee
          - role: operator
            name: ':'
            syntax: infix
            children:
            - role: string
              quote: ':'
              value: yearPublished
            - role: number
              value: 1960
        - role: delimited
          kind: braces
          separator: comma
          children:
          - role: operator
            name: ':'
            syntax: infix
            children:
            - role: string
              quote: ':'
              value: title
            - role: string
              quote: ','
              value: '1984'
          - role: operator
            name: ':'
            syntax: infix
            children:
            - role: string
              quote: ':'
              value: author
            - role: string
              quote: ','
              value: George Orwell
          - role: operator
            name: ':'
            syntax: infix
            children:
            - role: string
              quote: ':'
              value: yearPublished
            - role: number
              value: 1949

```

## Graphviz Dot format

```dot
digraph G {
  bgcolor="transparent";
  node [shape="box", style="filled", fontname="Ubuntu Mono"];
  "132440252320320" [label="delimited", shape="box", fillcolor="lightgray"];
  "132440252320400" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "132440252320320" -> "132440252320400";
  "132440252320480" [label="string", shape="box", fillcolor="lightgray"];
  "132440252320400" -> "132440252320480";
  "132440252320560" [label="delimited", shape="box", fillcolor="lightgray"];
  "132440252320400" -> "132440252320560";
  "132440252320640" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "132440252320560" -> "132440252320640";
  "132440252320720" [label="string", shape="box", fillcolor="lightgray"];
  "132440252320640" -> "132440252320720";
  "132440252320800" [label="string", shape="box", fillcolor="lightgray"];
  "132440252320640" -> "132440252320800";
  "132440252320880" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "132440252320560" -> "132440252320880";
  "132440252320960" [label="string", shape="box", fillcolor="lightgray"];
  "132440252320880" -> "132440252320960";
  "132440252321040" [label="number: 25", shape="box", fillcolor="lightgoldenrodyellow"];
  "132440252320880" -> "132440252321040";
  "132440252321120" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "132440252320560" -> "132440252321120";
  "132440252321200" [label="string", shape="box", fillcolor="lightgray"];
  "132440252321120" -> "132440252321200";
  "132440252321280" [label="identifier: true", shape="box", fillcolor="Honeydew"];
  "132440252321120" -> "132440252321280";
  "132440252321360" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "132440252320560" -> "132440252321360";
  "132440252321440" [label="string", shape="box", fillcolor="lightgray"];
  "132440252321360" -> "132440252321440";
  "132440252321520" [label="delimited", shape="box", fillcolor="lightgray"];
  "132440252321360" -> "132440252321520";
  "132440252321600" [label="string", shape="box", fillcolor="lightgray"];
  "132440252321520" -> "132440252321600";
  "132440252321680" [label="string", shape="box", fillcolor="lightgray"];
  "132440252321520" -> "132440252321680";
  "132440252321760" [label="string", shape="box", fillcolor="lightgray"];
  "132440252321520" -> "132440252321760";
  "132440252321840" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "132440252320560" -> "132440252321840";
  "132440252321920" [label="string", shape="box", fillcolor="lightgray"];
  "132440252321840" -> "132440252321920";
  "132440252322000" [label="delimited", shape="box", fillcolor="lightgray"];
  "132440252321840" -> "132440252322000";
  "132440252322080" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "132440252322000" -> "132440252322080";
  "132440252322240" [label="string", shape="box", fillcolor="lightgray"];
  "132440252322080" -> "132440252322240";
  "132440252322400" [label="string", shape="box", fillcolor="lightgray"];
  "132440252322080" -> "132440252322400";
  "132440252322480" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "132440252322000" -> "132440252322480";
  "132440252322640" [label="string", shape="box", fillcolor="lightgray"];
  "132440252322480" -> "132440252322640";
  "132440252322800" [label="string", shape="box", fillcolor="lightgray"];
  "132440252322480" -> "132440252322800";
  "132440252322880" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "132440252322000" -> "132440252322880";
  "132440252323040" [label="string", shape="box", fillcolor="lightgray"];
  "132440252322880" -> "132440252323040";
  "132440252323200" [label="string", shape="box", fillcolor="lightgray"];
  "132440252322880" -> "132440252323200";
  "132440252323280" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "132440252320560" -> "132440252323280";
  "132440252323360" [label="string", shape="box", fillcolor="lightgray"];
  "132440252323280" -> "132440252323360";
  "132440252323440" [label="delimited", shape="box", fillcolor="lightgray"];
  "132440252323280" -> "132440252323440";
  "132440252323520" [label="delimited", shape="box", fillcolor="lightgray"];
  "132440252323440" -> "132440252323520";
  "132440252323680" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "132440252323520" -> "132440252323680";
  "132440252323840" [label="string", shape="box", fillcolor="lightgray"];
  "132440252323680" -> "132440252323840";
  "132440252324000" [label="string", shape="box", fillcolor="lightgray"];
  "132440252323680" -> "132440252324000";
  "132440252324240" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "132440252323520" -> "132440252324240";
  "132440252324400" [label="string", shape="box", fillcolor="lightgray"];
  "132440252324240" -> "132440252324400";
  "132440252324560" [label="string", shape="box", fillcolor="lightgray"];
  "132440252324240" -> "132440252324560";
  "132440252324800" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "132440252323520" -> "132440252324800";
  "132440252324960" [label="string", shape="box", fillcolor="lightgray"];
  "132440252324800" -> "132440252324960";
  "132440252325120" [label="number: 1960", shape="box", fillcolor="lightgoldenrodyellow"];
  "132440252324800" -> "132440252325120";
  "132440252325280" [label="delimited", shape="box", fillcolor="lightgray"];
  "132440252323440" -> "132440252325280";
  "132440252325440" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "132440252325280" -> "132440252325440";
  "132440252325600" [label="string", shape="box", fillcolor="lightgray"];
  "132440252325440" -> "132440252325600";
  "132440252325760" [label="string", shape="box", fillcolor="lightgray"];
  "132440252325440" -> "132440252325760";
  "132440252326000" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "132440252325280" -> "132440252326000";
  "132440252326160" [label="string", shape="box", fillcolor="lightgray"];
  "132440252326000" -> "132440252326160";
  "132440252326320" [label="string", shape="box", fillcolor="lightgray"];
  "132440252326000" -> "132440252326320";
  "132440252326560" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "132440252325280" -> "132440252326560";
  "132440252326720" [label="string", shape="box", fillcolor="lightgray"];
  "132440252326560" -> "132440252326720";
  "132440252326880" [label="number: 1949", shape="box", fillcolor="lightgoldenrodyellow"];
  "132440252326560" -> "132440252326880";
}
```

![Image generated for example](images/factorial.png)
