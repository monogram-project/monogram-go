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
  124914153638544["delimited"]:::custom_delimited;
  124914153638624["operator: :"]:::custom_operator;
  124914153638544 --> 124914153638624;
  124914153638704["string: person"]:::custom_string;
  124914153638624 --> 124914153638704;
  124914153638784["delimited"]:::custom_delimited;
  124914153638624 --> 124914153638784;
  124914153638864["operator: :"]:::custom_operator;
  124914153638784 --> 124914153638864;
  124914153638944["string: name"]:::custom_string;
  124914153638864 --> 124914153638944;
  124914153639024["string: Alice"]:::custom_string;
  124914153638864 --> 124914153639024;
  124914153639104["operator: :"]:::custom_operator;
  124914153638784 --> 124914153639104;
  124914153639184["string: age"]:::custom_string;
  124914153639104 --> 124914153639184;
  124914153639264["number: 25"]:::custom_number;
  124914153639104 --> 124914153639264;
  124914153639344["operator: :"]:::custom_operator;
  124914153638784 --> 124914153639344;
  124914153639424["string: isStudent"]:::custom_string;
  124914153639344 --> 124914153639424;
  124914153639504["identifier: true"]:::custom_identifier;
  124914153639344 --> 124914153639504;
  124914153639584["operator: :"]:::custom_operator;
  124914153638784 --> 124914153639584;
  124914153639664["string: skills"]:::custom_string;
  124914153639584 --> 124914153639664;
  124914153639744["delimited"]:::custom_delimited;
  124914153639584 --> 124914153639744;
  124914153639824["string: Python"]:::custom_string;
  124914153639744 --> 124914153639824;
  124914153639904["string: JavaScript"]:::custom_string;
  124914153639744 --> 124914153639904;
  124914153639984["string: SQL"]:::custom_string;
  124914153639744 --> 124914153639984;
  124914153640064["operator: :"]:::custom_operator;
  124914153638784 --> 124914153640064;
  124914153640144["string: address"]:::custom_string;
  124914153640064 --> 124914153640144;
  124914153640224["delimited"]:::custom_delimited;
  124914153640064 --> 124914153640224;
  124914153640304["operator: :"]:::custom_operator;
  124914153640224 --> 124914153640304;
  124914153640464["string: street"]:::custom_string;
  124914153640304 --> 124914153640464;
  124914153640624["string:<br/>123 Maple Street"]:::custom_string;
  124914153640304 --> 124914153640624;
  124914153640704["operator: :"]:::custom_operator;
  124914153640224 --> 124914153640704;
  124914153640864["string: city"]:::custom_string;
  124914153640704 --> 124914153640864;
  124914153641024["string: Exampleville"]:::custom_string;
  124914153640704 --> 124914153641024;
  124914153641104["operator: :"]:::custom_operator;
  124914153640224 --> 124914153641104;
  124914153641264["string: country"]:::custom_string;
  124914153641104 --> 124914153641264;
  124914153641424["string: Neverland"]:::custom_string;
  124914153641104 --> 124914153641424;
  124914153641504["operator: :"]:::custom_operator;
  124914153638784 --> 124914153641504;
  124914153641584["string:<br/>favoriteBooks"]:::custom_string;
  124914153641504 --> 124914153641584;
  124914153641664["delimited"]:::custom_delimited;
  124914153641504 --> 124914153641664;
  124914153641744["delimited"]:::custom_delimited;
  124914153641664 --> 124914153641744;
  124914153641904["operator: :"]:::custom_operator;
  124914153641744 --> 124914153641904;
  124914153642064["string: title"]:::custom_string;
  124914153641904 --> 124914153642064;
  124914153642224["string:<br/>To Kill a Mockingbird"]:::custom_string;
  124914153641904 --> 124914153642224;
  124914153642464["operator: :"]:::custom_operator;
  124914153641744 --> 124914153642464;
  124914153642624["string: author"]:::custom_string;
  124914153642464 --> 124914153642624;
  124914153642784["string: Harper Lee"]:::custom_string;
  124914153642464 --> 124914153642784;
  124914153643024["operator: :"]:::custom_operator;
  124914153641744 --> 124914153643024;
  124914153643184["string:<br/>yearPublished"]:::custom_string;
  124914153643024 --> 124914153643184;
  124914153643344["number: 1960"]:::custom_number;
  124914153643024 --> 124914153643344;
  124914153643504["delimited"]:::custom_delimited;
  124914153641664 --> 124914153643504;
  124914153643664["operator: :"]:::custom_operator;
  124914153643504 --> 124914153643664;
  124914153643824["string: title"]:::custom_string;
  124914153643664 --> 124914153643824;
  124914153643984["string: 1984"]:::custom_string;
  124914153643664 --> 124914153643984;
  124914153644224["operator: :"]:::custom_operator;
  124914153643504 --> 124914153644224;
  124914153644384["string: author"]:::custom_string;
  124914153644224 --> 124914153644384;
  124914153644544["string:<br/>George Orwell"]:::custom_string;
  124914153644224 --> 124914153644544;
  124914153644784["operator: :"]:::custom_operator;
  124914153643504 --> 124914153644784;
  124914153644944["string:<br/>yearPublished"]:::custom_string;
  124914153644784 --> 124914153644944;
  124914153645104["number: 1949"]:::custom_number;
  124914153644784 --> 124914153645104;

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
    <operator name=":">
        <string quote=":" value="person"/>
        <delimited kind="braces" separator="comma">
            <operator name=":">
                <string quote=":" value="name"/>
                <string quote="," value="Alice"/>
            </operator>
            <operator name=":">
                <string quote=":" value="age"/>
                <number value="25"/>
            </operator>
            <operator name=":">
                <string quote=":" value="isStudent"/>
                <identifier name="true"/>
            </operator>
            <operator name=":">
                <string quote=":" value="skills"/>
                <delimited kind="brackets" separator="comma">
                    <string quote="," value="Python"/>
                    <string quote="," value="JavaScript"/>
                    <string quote="]" value="SQL"/>
                </delimited>
            </operator>
            <operator name=":">
                <string quote=":" value="address"/>
                <delimited kind="braces" separator="comma">
                    <operator name=":">
                        <string quote=":" value="street"/>
                        <string quote="," value="123 Maple Street"/>
                    </operator>
                    <operator name=":">
                        <string quote=":" value="city"/>
                        <string quote="," value="Exampleville"/>
                    </operator>
                    <operator name=":">
                        <string quote=":" value="country"/>
                        <string quote="&quot;" value="Neverland"/>
                    </operator>
                </delimited>
            </operator>
            <operator name=":">
                <string quote=":" value="favoriteBooks"/>
                <delimited kind="brackets" separator="comma">
                    <delimited kind="braces" separator="comma">
                        <operator name=":">
                            <string quote=":" value="title"/>
                            <string quote="," value="To Kill a Mockingbird"/>
                        </operator>
                        <operator name=":">
                            <string quote=":" value="author"/>
                            <string quote="," value="Harper Lee"/>
                        </operator>
                        <operator name=":">
                            <string quote=":" value="yearPublished"/>
                            <number value="1960"/>
                        </operator>
                    </delimited>
                    <delimited kind="braces" separator="comma">
                        <operator name=":">
                            <string quote=":" value="title"/>
                            <string quote="," value="1984"/>
                        </operator>
                        <operator name=":">
                            <string quote=":" value="author"/>
                            <string quote="," value="George Orwell"/>
                        </operator>
                        <operator name=":">
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
      children:
      - role: string
        quote: ':'
        value: name
      - role: string
        quote: ','
        value: Alice
    - role: operator
      name: ':'
      children:
      - role: string
        quote: ':'
        value: age
      - role: number
        value: 25
    - role: operator
      name: ':'
      children:
      - role: string
        quote: ':'
        value: isStudent
      - role: identifier
        name: 'true'
    - role: operator
      name: ':'
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
          children:
          - role: string
            quote: ':'
            value: street
          - role: string
            quote: ','
            value: 123 Maple Street
        - role: operator
          name: ':'
          children:
          - role: string
            quote: ':'
            value: city
          - role: string
            quote: ','
            value: Exampleville
        - role: operator
          name: ':'
          children:
          - role: string
            quote: ':'
            value: country
          - role: string
            quote: '"'
            value: Neverland
    - role: operator
      name: ':'
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
            children:
            - role: string
              quote: ':'
              value: title
            - role: string
              quote: ','
              value: To Kill a Mockingbird
          - role: operator
            name: ':'
            children:
            - role: string
              quote: ':'
              value: author
            - role: string
              quote: ','
              value: Harper Lee
          - role: operator
            name: ':'
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
            children:
            - role: string
              quote: ':'
              value: title
            - role: string
              quote: ','
              value: '1984'
          - role: operator
            name: ':'
            children:
            - role: string
              quote: ':'
              value: author
            - role: string
              quote: ','
              value: George Orwell
          - role: operator
            name: ':'
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
  "139478584281664" [label="delimited", shape="box", fillcolor="lightgray"];
  "139478584281744" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "139478584281664" -> "139478584281744";
  "139478584281824" [label="string", shape="box", fillcolor="lightgray"];
  "139478584281744" -> "139478584281824";
  "139478584281904" [label="delimited", shape="box", fillcolor="lightgray"];
  "139478584281744" -> "139478584281904";
  "139478584281984" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "139478584281904" -> "139478584281984";
  "139478584282064" [label="string", shape="box", fillcolor="lightgray"];
  "139478584281984" -> "139478584282064";
  "139478584282144" [label="string", shape="box", fillcolor="lightgray"];
  "139478584281984" -> "139478584282144";
  "139478584282224" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "139478584281904" -> "139478584282224";
  "139478584282304" [label="string", shape="box", fillcolor="lightgray"];
  "139478584282224" -> "139478584282304";
  "139478584282384" [label="number: 25", shape="box", fillcolor="lightgoldenrodyellow"];
  "139478584282224" -> "139478584282384";
  "139478584282464" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "139478584281904" -> "139478584282464";
  "139478584282544" [label="string", shape="box", fillcolor="lightgray"];
  "139478584282464" -> "139478584282544";
  "139478584282624" [label="identifier: true", shape="box", fillcolor="Honeydew"];
  "139478584282464" -> "139478584282624";
  "139478584282704" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "139478584281904" -> "139478584282704";
  "139478584282784" [label="string", shape="box", fillcolor="lightgray"];
  "139478584282704" -> "139478584282784";
  "139478584282864" [label="delimited", shape="box", fillcolor="lightgray"];
  "139478584282704" -> "139478584282864";
  "139478584282944" [label="string", shape="box", fillcolor="lightgray"];
  "139478584282864" -> "139478584282944";
  "139478584283024" [label="string", shape="box", fillcolor="lightgray"];
  "139478584282864" -> "139478584283024";
  "139478584283104" [label="string", shape="box", fillcolor="lightgray"];
  "139478584282864" -> "139478584283104";
  "139478584283184" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "139478584281904" -> "139478584283184";
  "139478584283264" [label="string", shape="box", fillcolor="lightgray"];
  "139478584283184" -> "139478584283264";
  "139478584283344" [label="delimited", shape="box", fillcolor="lightgray"];
  "139478584283184" -> "139478584283344";
  "139478584283424" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "139478584283344" -> "139478584283424";
  "139478584283584" [label="string", shape="box", fillcolor="lightgray"];
  "139478584283424" -> "139478584283584";
  "139478584283744" [label="string", shape="box", fillcolor="lightgray"];
  "139478584283424" -> "139478584283744";
  "139478584283824" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "139478584283344" -> "139478584283824";
  "139478584283984" [label="string", shape="box", fillcolor="lightgray"];
  "139478584283824" -> "139478584283984";
  "139478584284144" [label="string", shape="box", fillcolor="lightgray"];
  "139478584283824" -> "139478584284144";
  "139478584284224" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "139478584283344" -> "139478584284224";
  "139478584284384" [label="string", shape="box", fillcolor="lightgray"];
  "139478584284224" -> "139478584284384";
  "139478584284544" [label="string", shape="box", fillcolor="lightgray"];
  "139478584284224" -> "139478584284544";
  "139478584284624" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "139478584281904" -> "139478584284624";
  "139478584284704" [label="string", shape="box", fillcolor="lightgray"];
  "139478584284624" -> "139478584284704";
  "139478584284784" [label="delimited", shape="box", fillcolor="lightgray"];
  "139478584284624" -> "139478584284784";
  "139478584284864" [label="delimited", shape="box", fillcolor="lightgray"];
  "139478584284784" -> "139478584284864";
  "139478584285024" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "139478584284864" -> "139478584285024";
  "139478584285184" [label="string", shape="box", fillcolor="lightgray"];
  "139478584285024" -> "139478584285184";
  "139478584285344" [label="string", shape="box", fillcolor="lightgray"];
  "139478584285024" -> "139478584285344";
  "139478584285584" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "139478584284864" -> "139478584285584";
  "139478584285744" [label="string", shape="box", fillcolor="lightgray"];
  "139478584285584" -> "139478584285744";
  "139478584285904" [label="string", shape="box", fillcolor="lightgray"];
  "139478584285584" -> "139478584285904";
  "139478584286144" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "139478584284864" -> "139478584286144";
  "139478584286304" [label="string", shape="box", fillcolor="lightgray"];
  "139478584286144" -> "139478584286304";
  "139478584286464" [label="number: 1960", shape="box", fillcolor="lightgoldenrodyellow"];
  "139478584286144" -> "139478584286464";
  "139478584286624" [label="delimited", shape="box", fillcolor="lightgray"];
  "139478584284784" -> "139478584286624";
  "139478584286784" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "139478584286624" -> "139478584286784";
  "139478584286944" [label="string", shape="box", fillcolor="lightgray"];
  "139478584286784" -> "139478584286944";
  "139478584287104" [label="string", shape="box", fillcolor="lightgray"];
  "139478584286784" -> "139478584287104";
  "139478584287344" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "139478584286624" -> "139478584287344";
  "139478584287504" [label="string", shape="box", fillcolor="lightgray"];
  "139478584287344" -> "139478584287504";
  "139478584287664" [label="string", shape="box", fillcolor="lightgray"];
  "139478584287344" -> "139478584287664";
  "139478584287904" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "139478584286624" -> "139478584287904";
  "139478584288064" [label="string", shape="box", fillcolor="lightgray"];
  "139478584287904" -> "139478584288064";
  "139478584288224" [label="number: 1949", shape="box", fillcolor="lightgoldenrodyellow"];
  "139478584287904" -> "139478584288224";
}
```

![Image generated for example](images/factorial.png)
