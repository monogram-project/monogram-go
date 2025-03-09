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
  128574339500608["delimited"]:::custom_delimited;
  128574339500688["operator: :"]:::custom_operator;
  128574339500608 --> 128574339500688;
  128574339500768["string: person"]:::custom_string;
  128574339500688 --> 128574339500768;
  128574339500848["delimited"]:::custom_delimited;
  128574339500688 --> 128574339500848;
  128574339500928["operator: :"]:::custom_operator;
  128574339500848 --> 128574339500928;
  128574339501008["string: name"]:::custom_string;
  128574339500928 --> 128574339501008;
  128574339501088["string: Alice"]:::custom_string;
  128574339500928 --> 128574339501088;
  128574339501168["operator: :"]:::custom_operator;
  128574339500848 --> 128574339501168;
  128574339501248["string: age"]:::custom_string;
  128574339501168 --> 128574339501248;
  128574339501328["number: 25"]:::custom_number;
  128574339501168 --> 128574339501328;
  128574339501408["operator: :"]:::custom_operator;
  128574339500848 --> 128574339501408;
  128574339501488["string: isStudent"]:::custom_string;
  128574339501408 --> 128574339501488;
  128574339501568["identifier: true"]:::custom_identifier;
  128574339501408 --> 128574339501568;
  128574339501648["operator: :"]:::custom_operator;
  128574339500848 --> 128574339501648;
  128574339501728["string: skills"]:::custom_string;
  128574339501648 --> 128574339501728;
  128574339501808["delimited"]:::custom_delimited;
  128574339501648 --> 128574339501808;
  128574339501888["string: Python"]:::custom_string;
  128574339501808 --> 128574339501888;
  128574339501968["string: JavaScript"]:::custom_string;
  128574339501808 --> 128574339501968;
  128574339502048["string: SQL"]:::custom_string;
  128574339501808 --> 128574339502048;
  128574339502128["operator: :"]:::custom_operator;
  128574339500848 --> 128574339502128;
  128574339502208["string: address"]:::custom_string;
  128574339502128 --> 128574339502208;
  128574339502288["delimited"]:::custom_delimited;
  128574339502128 --> 128574339502288;
  128574339502368["operator: :"]:::custom_operator;
  128574339502288 --> 128574339502368;
  128574339502528["string: street"]:::custom_string;
  128574339502368 --> 128574339502528;
  128574339502688["string:<br/>123 Maple Street"]:::custom_string;
  128574339502368 --> 128574339502688;
  128574339502768["operator: :"]:::custom_operator;
  128574339502288 --> 128574339502768;
  128574339502928["string: city"]:::custom_string;
  128574339502768 --> 128574339502928;
  128574339503088["string: Exampleville"]:::custom_string;
  128574339502768 --> 128574339503088;
  128574339503168["operator: :"]:::custom_operator;
  128574339502288 --> 128574339503168;
  128574339503328["string: country"]:::custom_string;
  128574339503168 --> 128574339503328;
  128574339503488["string: Neverland"]:::custom_string;
  128574339503168 --> 128574339503488;
  128574339503568["operator: :"]:::custom_operator;
  128574339500848 --> 128574339503568;
  128574339503648["string:<br/>favoriteBooks"]:::custom_string;
  128574339503568 --> 128574339503648;
  128574339503728["delimited"]:::custom_delimited;
  128574339503568 --> 128574339503728;
  128574339503808["delimited"]:::custom_delimited;
  128574339503728 --> 128574339503808;
  128574339503968["operator: :"]:::custom_operator;
  128574339503808 --> 128574339503968;
  128574339504128["string: title"]:::custom_string;
  128574339503968 --> 128574339504128;
  128574339504288["string:<br/>To Kill a Mockingbird"]:::custom_string;
  128574339503968 --> 128574339504288;
  128574339504528["operator: :"]:::custom_operator;
  128574339503808 --> 128574339504528;
  128574339504688["string: author"]:::custom_string;
  128574339504528 --> 128574339504688;
  128574339504848["string: Harper Lee"]:::custom_string;
  128574339504528 --> 128574339504848;
  128574339505088["operator: :"]:::custom_operator;
  128574339503808 --> 128574339505088;
  128574339505248["string:<br/>yearPublished"]:::custom_string;
  128574339505088 --> 128574339505248;
  128574339505408["number: 1960"]:::custom_number;
  128574339505088 --> 128574339505408;
  128574339505568["delimited"]:::custom_delimited;
  128574339503728 --> 128574339505568;
  128574339505728["operator: :"]:::custom_operator;
  128574339505568 --> 128574339505728;
  128574339505888["string: title"]:::custom_string;
  128574339505728 --> 128574339505888;
  128574339506048["string: 1984"]:::custom_string;
  128574339505728 --> 128574339506048;
  128574339506288["operator: :"]:::custom_operator;
  128574339505568 --> 128574339506288;
  128574339506448["string: author"]:::custom_string;
  128574339506288 --> 128574339506448;
  128574339506608["string:<br/>George Orwell"]:::custom_string;
  128574339506288 --> 128574339506608;
  128574339506848["operator: :"]:::custom_operator;
  128574339505568 --> 128574339506848;
  128574339507008["string:<br/>yearPublished"]:::custom_string;
  128574339506848 --> 128574339507008;
  128574339507168["number: 1949"]:::custom_number;
  128574339506848 --> 128574339507168;

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
  "139409612081728" [label="delimited", shape="box", fillcolor="lightgray"];
  "139409612081808" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "139409612081728" -> "139409612081808";
  "139409612081888" [label="string", shape="box", fillcolor="lightgray"];
  "139409612081808" -> "139409612081888";
  "139409612081968" [label="delimited", shape="box", fillcolor="lightgray"];
  "139409612081808" -> "139409612081968";
  "139409612082048" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "139409612081968" -> "139409612082048";
  "139409612082128" [label="string", shape="box", fillcolor="lightgray"];
  "139409612082048" -> "139409612082128";
  "139409612082208" [label="string", shape="box", fillcolor="lightgray"];
  "139409612082048" -> "139409612082208";
  "139409612082288" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "139409612081968" -> "139409612082288";
  "139409612082368" [label="string", shape="box", fillcolor="lightgray"];
  "139409612082288" -> "139409612082368";
  "139409612082448" [label="number: 25", shape="box", fillcolor="lightgoldenrodyellow"];
  "139409612082288" -> "139409612082448";
  "139409612082528" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "139409612081968" -> "139409612082528";
  "139409612082608" [label="string", shape="box", fillcolor="lightgray"];
  "139409612082528" -> "139409612082608";
  "139409612082688" [label="identifier: true", shape="box", fillcolor="Honeydew"];
  "139409612082528" -> "139409612082688";
  "139409612082768" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "139409612081968" -> "139409612082768";
  "139409612082848" [label="string", shape="box", fillcolor="lightgray"];
  "139409612082768" -> "139409612082848";
  "139409612082928" [label="delimited", shape="box", fillcolor="lightgray"];
  "139409612082768" -> "139409612082928";
  "139409612083008" [label="string", shape="box", fillcolor="lightgray"];
  "139409612082928" -> "139409612083008";
  "139409612083088" [label="string", shape="box", fillcolor="lightgray"];
  "139409612082928" -> "139409612083088";
  "139409612083168" [label="string", shape="box", fillcolor="lightgray"];
  "139409612082928" -> "139409612083168";
  "139409612083248" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "139409612081968" -> "139409612083248";
  "139409612083328" [label="string", shape="box", fillcolor="lightgray"];
  "139409612083248" -> "139409612083328";
  "139409612083408" [label="delimited", shape="box", fillcolor="lightgray"];
  "139409612083248" -> "139409612083408";
  "139409612083488" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "139409612083408" -> "139409612083488";
  "139409612083648" [label="string", shape="box", fillcolor="lightgray"];
  "139409612083488" -> "139409612083648";
  "139409612083808" [label="string", shape="box", fillcolor="lightgray"];
  "139409612083488" -> "139409612083808";
  "139409612083888" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "139409612083408" -> "139409612083888";
  "139409612084048" [label="string", shape="box", fillcolor="lightgray"];
  "139409612083888" -> "139409612084048";
  "139409612084208" [label="string", shape="box", fillcolor="lightgray"];
  "139409612083888" -> "139409612084208";
  "139409612084288" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "139409612083408" -> "139409612084288";
  "139409612084448" [label="string", shape="box", fillcolor="lightgray"];
  "139409612084288" -> "139409612084448";
  "139409612084608" [label="string", shape="box", fillcolor="lightgray"];
  "139409612084288" -> "139409612084608";
  "139409612084688" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "139409612081968" -> "139409612084688";
  "139409612084768" [label="string", shape="box", fillcolor="lightgray"];
  "139409612084688" -> "139409612084768";
  "139409612084848" [label="delimited", shape="box", fillcolor="lightgray"];
  "139409612084688" -> "139409612084848";
  "139409612084928" [label="delimited", shape="box", fillcolor="lightgray"];
  "139409612084848" -> "139409612084928";
  "139409612085088" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "139409612084928" -> "139409612085088";
  "139409612085248" [label="string", shape="box", fillcolor="lightgray"];
  "139409612085088" -> "139409612085248";
  "139409612085408" [label="string", shape="box", fillcolor="lightgray"];
  "139409612085088" -> "139409612085408";
  "139409612085648" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "139409612084928" -> "139409612085648";
  "139409612085808" [label="string", shape="box", fillcolor="lightgray"];
  "139409612085648" -> "139409612085808";
  "139409612085968" [label="string", shape="box", fillcolor="lightgray"];
  "139409612085648" -> "139409612085968";
  "139409612086208" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "139409612084928" -> "139409612086208";
  "139409612086368" [label="string", shape="box", fillcolor="lightgray"];
  "139409612086208" -> "139409612086368";
  "139409612086528" [label="number: 1960", shape="box", fillcolor="lightgoldenrodyellow"];
  "139409612086208" -> "139409612086528";
  "139409612086688" [label="delimited", shape="box", fillcolor="lightgray"];
  "139409612084848" -> "139409612086688";
  "139409612086848" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "139409612086688" -> "139409612086848";
  "139409612087008" [label="string", shape="box", fillcolor="lightgray"];
  "139409612086848" -> "139409612087008";
  "139409612087168" [label="string", shape="box", fillcolor="lightgray"];
  "139409612086848" -> "139409612087168";
  "139409612087408" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "139409612086688" -> "139409612087408";
  "139409612087568" [label="string", shape="box", fillcolor="lightgray"];
  "139409612087408" -> "139409612087568";
  "139409612087728" [label="string", shape="box", fillcolor="lightgray"];
  "139409612087408" -> "139409612087728";
  "139409612087968" [label="operator", shape="box", fillcolor="#C0FFC0"];
  "139409612086688" -> "139409612087968";
  "139409612088128" [label="string", shape="box", fillcolor="lightgray"];
  "139409612087968" -> "139409612088128";
  "139409612088288" [label="number: 1949", shape="box", fillcolor="lightgoldenrodyellow"];
  "139409612087968" -> "139409612088288";
}
```

![Image generated for example](images/factorial.png)
