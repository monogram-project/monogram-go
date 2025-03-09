# JSON Expression

## Monogram

Monograph supports JSON expressions. Since Monogram is a good deal more 
complicated than JSON, it is not too surprising that translating a JSON
expression into a Monogram-tree looks quite different from plain JSON!

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
  135567336016528["delimited"]:::custom_delimited;
  135567336016608["operator: :"]:::custom_operator;
  135567336016528 --> 135567336016608;
  135567336016688["string: person"]:::custom_string;
  135567336016608 --> 135567336016688;
  135567336016768["delimited"]:::custom_delimited;
  135567336016608 --> 135567336016768;
  135567336016848["operator: :"]:::custom_operator;
  135567336016768 --> 135567336016848;
  135567336016928["string: name"]:::custom_string;
  135567336016848 --> 135567336016928;
  135567336017008["string: Alice"]:::custom_string;
  135567336016848 --> 135567336017008;
  135567336017088["operator: :"]:::custom_operator;
  135567336016768 --> 135567336017088;
  135567336017168["string: age"]:::custom_string;
  135567336017088 --> 135567336017168;
  135567336017248["number: 25"]:::custom_number;
  135567336017088 --> 135567336017248;
  135567336017328["operator: :"]:::custom_operator;
  135567336016768 --> 135567336017328;
  135567336017408["string: isStudent"]:::custom_string;
  135567336017328 --> 135567336017408;
  135567336017488["identifier: true"]:::custom_identifier;
  135567336017328 --> 135567336017488;
  135567336017568["operator: :"]:::custom_operator;
  135567336016768 --> 135567336017568;
  135567336017648["string: skills"]:::custom_string;
  135567336017568 --> 135567336017648;
  135567336017728["delimited"]:::custom_delimited;
  135567336017568 --> 135567336017728;
  135567336017808["string: Python"]:::custom_string;
  135567336017728 --> 135567336017808;
  135567336017888["string: JavaScript"]:::custom_string;
  135567336017728 --> 135567336017888;
  135567336017968["string: SQL"]:::custom_string;
  135567336017728 --> 135567336017968;
  135567336018048["operator: :"]:::custom_operator;
  135567336016768 --> 135567336018048;
  135567336018128["string: address"]:::custom_string;
  135567336018048 --> 135567336018128;
  135567336018208["delimited"]:::custom_delimited;
  135567336018048 --> 135567336018208;
  135567336018288["operator: :"]:::custom_operator;
  135567336018208 --> 135567336018288;
  135567336018448["string: street"]:::custom_string;
  135567336018288 --> 135567336018448;
  135567336018608["string:<br/>123 Maple Street"]:::custom_string;
  135567336018288 --> 135567336018608;
  135567336018688["operator: :"]:::custom_operator;
  135567336018208 --> 135567336018688;
  135567336018848["string: city"]:::custom_string;
  135567336018688 --> 135567336018848;
  135567336019008["string: Exampleville"]:::custom_string;
  135567336018688 --> 135567336019008;
  135567336019088["operator: :"]:::custom_operator;
  135567336018208 --> 135567336019088;
  135567336019248["string: country"]:::custom_string;
  135567336019088 --> 135567336019248;
  135567336019408["string: Neverland"]:::custom_string;
  135567336019088 --> 135567336019408;
  135567336019488["operator: :"]:::custom_operator;
  135567336016768 --> 135567336019488;
  135567336019568["string:<br/>favoriteBooks"]:::custom_string;
  135567336019488 --> 135567336019568;
  135567336019648["delimited"]:::custom_delimited;
  135567336019488 --> 135567336019648;
  135567336019728["delimited"]:::custom_delimited;
  135567336019648 --> 135567336019728;
  135567336019888["operator: :"]:::custom_operator;
  135567336019728 --> 135567336019888;
  135567336020048["string: title"]:::custom_string;
  135567336019888 --> 135567336020048;
  135567336020208["string:<br/>To Kill a Mockingbird"]:::custom_string;
  135567336019888 --> 135567336020208;
  135567336020448["operator: :"]:::custom_operator;
  135567336019728 --> 135567336020448;
  135567336020608["string: author"]:::custom_string;
  135567336020448 --> 135567336020608;
  135567336020768["string: Harper Lee"]:::custom_string;
  135567336020448 --> 135567336020768;
  135567336021008["operator: :"]:::custom_operator;
  135567336019728 --> 135567336021008;
  135567336021168["string:<br/>yearPublished"]:::custom_string;
  135567336021008 --> 135567336021168;
  135567336021328["number: 1960"]:::custom_number;
  135567336021008 --> 135567336021328;
  135567336021488["delimited"]:::custom_delimited;
  135567336019648 --> 135567336021488;
  135567336021648["operator: :"]:::custom_operator;
  135567336021488 --> 135567336021648;
  135567336021808["string: title"]:::custom_string;
  135567336021648 --> 135567336021808;
  135567336021968["string: 1984"]:::custom_string;
  135567336021648 --> 135567336021968;
  135567336022208["operator: :"]:::custom_operator;
  135567336021488 --> 135567336022208;
  135567336022368["string: author"]:::custom_string;
  135567336022208 --> 135567336022368;
  135567336022528["string:<br/>George Orwell"]:::custom_string;
  135567336022208 --> 135567336022528;
  135567336022768["operator: :"]:::custom_operator;
  135567336021488 --> 135567336022768;
  135567336022928["string:<br/>yearPublished"]:::custom_string;
  135567336022768 --> 135567336022928;
  135567336023088["number: 1949"]:::custom_number;
  135567336022768 --> 135567336023088;

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
  "127935933057600" [label="delimited", shape="box", fillcolor="lightgray"];
  "127935933057680" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "127935933057600" -> "127935933057680";
  "127935933057760" [label="string", shape="box", fillcolor="lightgray"];
  "127935933057680" -> "127935933057760";
  "127935933057840" [label="delimited", shape="box", fillcolor="lightgray"];
  "127935933057680" -> "127935933057840";
  "127935933057920" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "127935933057840" -> "127935933057920";
  "127935933058000" [label="string", shape="box", fillcolor="lightgray"];
  "127935933057920" -> "127935933058000";
  "127935933058080" [label="string", shape="box", fillcolor="lightgray"];
  "127935933057920" -> "127935933058080";
  "127935933058160" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "127935933057840" -> "127935933058160";
  "127935933058240" [label="string", shape="box", fillcolor="lightgray"];
  "127935933058160" -> "127935933058240";
  "127935933058320" [label="number: 25", shape="box", fillcolor="lightgoldenrodyellow"];
  "127935933058160" -> "127935933058320";
  "127935933058400" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "127935933057840" -> "127935933058400";
  "127935933058480" [label="string", shape="box", fillcolor="lightgray"];
  "127935933058400" -> "127935933058480";
  "127935933058560" [label="identifier: true", shape="box", fillcolor="Honeydew"];
  "127935933058400" -> "127935933058560";
  "127935933058640" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "127935933057840" -> "127935933058640";
  "127935933058720" [label="string", shape="box", fillcolor="lightgray"];
  "127935933058640" -> "127935933058720";
  "127935933058800" [label="delimited", shape="box", fillcolor="lightgray"];
  "127935933058640" -> "127935933058800";
  "127935933058880" [label="string", shape="box", fillcolor="lightgray"];
  "127935933058800" -> "127935933058880";
  "127935933058960" [label="string", shape="box", fillcolor="lightgray"];
  "127935933058800" -> "127935933058960";
  "127935933059040" [label="string", shape="box", fillcolor="lightgray"];
  "127935933058800" -> "127935933059040";
  "127935933059120" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "127935933057840" -> "127935933059120";
  "127935933059200" [label="string", shape="box", fillcolor="lightgray"];
  "127935933059120" -> "127935933059200";
  "127935933059280" [label="delimited", shape="box", fillcolor="lightgray"];
  "127935933059120" -> "127935933059280";
  "127935933059360" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "127935933059280" -> "127935933059360";
  "127935933059520" [label="string", shape="box", fillcolor="lightgray"];
  "127935933059360" -> "127935933059520";
  "127935933059680" [label="string", shape="box", fillcolor="lightgray"];
  "127935933059360" -> "127935933059680";
  "127935933059760" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "127935933059280" -> "127935933059760";
  "127935933059920" [label="string", shape="box", fillcolor="lightgray"];
  "127935933059760" -> "127935933059920";
  "127935933060080" [label="string", shape="box", fillcolor="lightgray"];
  "127935933059760" -> "127935933060080";
  "127935933060160" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "127935933059280" -> "127935933060160";
  "127935933060320" [label="string", shape="box", fillcolor="lightgray"];
  "127935933060160" -> "127935933060320";
  "127935933060480" [label="string", shape="box", fillcolor="lightgray"];
  "127935933060160" -> "127935933060480";
  "127935933060560" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "127935933057840" -> "127935933060560";
  "127935933060640" [label="string", shape="box", fillcolor="lightgray"];
  "127935933060560" -> "127935933060640";
  "127935933060720" [label="delimited", shape="box", fillcolor="lightgray"];
  "127935933060560" -> "127935933060720";
  "127935933060800" [label="delimited", shape="box", fillcolor="lightgray"];
  "127935933060720" -> "127935933060800";
  "127935933060960" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "127935933060800" -> "127935933060960";
  "127935933061120" [label="string", shape="box", fillcolor="lightgray"];
  "127935933060960" -> "127935933061120";
  "127935933061280" [label="string", shape="box", fillcolor="lightgray"];
  "127935933060960" -> "127935933061280";
  "127935933061520" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "127935933060800" -> "127935933061520";
  "127935933061680" [label="string", shape="box", fillcolor="lightgray"];
  "127935933061520" -> "127935933061680";
  "127935933061840" [label="string", shape="box", fillcolor="lightgray"];
  "127935933061520" -> "127935933061840";
  "127935933062080" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "127935933060800" -> "127935933062080";
  "127935933062240" [label="string", shape="box", fillcolor="lightgray"];
  "127935933062080" -> "127935933062240";
  "127935933062400" [label="number: 1960", shape="box", fillcolor="lightgoldenrodyellow"];
  "127935933062080" -> "127935933062400";
  "127935933062560" [label="delimited", shape="box", fillcolor="lightgray"];
  "127935933060720" -> "127935933062560";
  "127935933062720" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "127935933062560" -> "127935933062720";
  "127935933062880" [label="string", shape="box", fillcolor="lightgray"];
  "127935933062720" -> "127935933062880";
  "127935933063040" [label="string", shape="box", fillcolor="lightgray"];
  "127935933062720" -> "127935933063040";
  "127935933063280" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "127935933062560" -> "127935933063280";
  "127935933063440" [label="string", shape="box", fillcolor="lightgray"];
  "127935933063280" -> "127935933063440";
  "127935933063600" [label="string", shape="box", fillcolor="lightgray"];
  "127935933063280" -> "127935933063600";
  "127935933063840" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "127935933062560" -> "127935933063840";
  "127935933064000" [label="string", shape="box", fillcolor="lightgray"];
  "127935933063840" -> "127935933064000";
  "127935933064160" [label="number: 1949", shape="box", fillcolor="lightgoldenrodyellow"];
  "127935933063840" -> "127935933064160";
}
```

![Image generated for example](images/factorial.png)
