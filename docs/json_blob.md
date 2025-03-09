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
  127816185283824["delimited"]:::custom_delimited;
  127816185283904["operator: :"]:::custom_operator;
  127816185283824 --> 127816185283904;
  127816185283984["string: person"]:::custom_string;
  127816185283904 --> 127816185283984;
  127816185284064["delimited"]:::custom_delimited;
  127816185283904 --> 127816185284064;
  127816185284144["operator: :"]:::custom_operator;
  127816185284064 --> 127816185284144;
  127816185284224["string: name"]:::custom_string;
  127816185284144 --> 127816185284224;
  127816185284304["string: Alice"]:::custom_string;
  127816185284144 --> 127816185284304;
  127816185284384["operator: :"]:::custom_operator;
  127816185284064 --> 127816185284384;
  127816185284464["string: age"]:::custom_string;
  127816185284384 --> 127816185284464;
  127816185284544["number: 25"]:::custom_number;
  127816185284384 --> 127816185284544;
  127816185284624["operator: :"]:::custom_operator;
  127816185284064 --> 127816185284624;
  127816185284704["string: isStudent"]:::custom_string;
  127816185284624 --> 127816185284704;
  127816185284784["identifier: true"]:::custom_identifier;
  127816185284624 --> 127816185284784;
  127816185284864["operator: :"]:::custom_operator;
  127816185284064 --> 127816185284864;
  127816185284944["string: skills"]:::custom_string;
  127816185284864 --> 127816185284944;
  127816185285024["delimited"]:::custom_delimited;
  127816185284864 --> 127816185285024;
  127816185285104["string: Python"]:::custom_string;
  127816185285024 --> 127816185285104;
  127816185285184["string: JavaScript"]:::custom_string;
  127816185285024 --> 127816185285184;
  127816185285264["string: SQL"]:::custom_string;
  127816185285024 --> 127816185285264;
  127816185285344["operator: :"]:::custom_operator;
  127816185284064 --> 127816185285344;
  127816185285424["string: address"]:::custom_string;
  127816185285344 --> 127816185285424;
  127816185285504["delimited"]:::custom_delimited;
  127816185285344 --> 127816185285504;
  127816185285584["operator: :"]:::custom_operator;
  127816185285504 --> 127816185285584;
  127816185285744["string: street"]:::custom_string;
  127816185285584 --> 127816185285744;
  127816185285904["string:<br/>123 Maple Street"]:::custom_string;
  127816185285584 --> 127816185285904;
  127816185285984["operator: :"]:::custom_operator;
  127816185285504 --> 127816185285984;
  127816185286144["string: city"]:::custom_string;
  127816185285984 --> 127816185286144;
  127816185286304["string: Exampleville"]:::custom_string;
  127816185285984 --> 127816185286304;
  127816185286384["operator: :"]:::custom_operator;
  127816185285504 --> 127816185286384;
  127816185286544["string: country"]:::custom_string;
  127816185286384 --> 127816185286544;
  127816185286704["string: Neverland"]:::custom_string;
  127816185286384 --> 127816185286704;
  127816185286784["operator: :"]:::custom_operator;
  127816185284064 --> 127816185286784;
  127816185286864["string:<br/>favoriteBooks"]:::custom_string;
  127816185286784 --> 127816185286864;
  127816185286944["delimited"]:::custom_delimited;
  127816185286784 --> 127816185286944;
  127816185287024["delimited"]:::custom_delimited;
  127816185286944 --> 127816185287024;
  127816185287184["operator: :"]:::custom_operator;
  127816185287024 --> 127816185287184;
  127816185287344["string: title"]:::custom_string;
  127816185287184 --> 127816185287344;
  127816185287504["string:<br/>To Kill a Mockingbird"]:::custom_string;
  127816185287184 --> 127816185287504;
  127816185287744["operator: :"]:::custom_operator;
  127816185287024 --> 127816185287744;
  127816185287904["string: author"]:::custom_string;
  127816185287744 --> 127816185287904;
  127816185288064["string: Harper Lee"]:::custom_string;
  127816185287744 --> 127816185288064;
  127816185288304["operator: :"]:::custom_operator;
  127816185287024 --> 127816185288304;
  127816185288464["string:<br/>yearPublished"]:::custom_string;
  127816185288304 --> 127816185288464;
  127816185288624["number: 1960"]:::custom_number;
  127816185288304 --> 127816185288624;
  127816185288784["delimited"]:::custom_delimited;
  127816185286944 --> 127816185288784;
  127816185288944["operator: :"]:::custom_operator;
  127816185288784 --> 127816185288944;
  127816185289104["string: title"]:::custom_string;
  127816185288944 --> 127816185289104;
  127816185289264["string: 1984"]:::custom_string;
  127816185288944 --> 127816185289264;
  127816185289504["operator: :"]:::custom_operator;
  127816185288784 --> 127816185289504;
  127816185289664["string: author"]:::custom_string;
  127816185289504 --> 127816185289664;
  127816185289824["string:<br/>George Orwell"]:::custom_string;
  127816185289504 --> 127816185289824;
  127816185290064["operator: :"]:::custom_operator;
  127816185288784 --> 127816185290064;
  127816185290224["string:<br/>yearPublished"]:::custom_string;
  127816185290064 --> 127816185290224;
  127816185290384["number: 1949"]:::custom_number;
  127816185290064 --> 127816185290384;

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
  "128030047569216" [label="delimited", shape="box", fillcolor="lightgray"];
  "128030047569296" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "128030047569216" -> "128030047569296";
  "128030047569376" [label="string", shape="box", fillcolor="lightgray"];
  "128030047569296" -> "128030047569376";
  "128030047569456" [label="delimited", shape="box", fillcolor="lightgray"];
  "128030047569296" -> "128030047569456";
  "128030047569536" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "128030047569456" -> "128030047569536";
  "128030047569616" [label="string", shape="box", fillcolor="lightgray"];
  "128030047569536" -> "128030047569616";
  "128030047569696" [label="string", shape="box", fillcolor="lightgray"];
  "128030047569536" -> "128030047569696";
  "128030047569776" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "128030047569456" -> "128030047569776";
  "128030047569856" [label="string", shape="box", fillcolor="lightgray"];
  "128030047569776" -> "128030047569856";
  "128030047569936" [label="number: 25", shape="box", fillcolor="lightgoldenrodyellow"];
  "128030047569776" -> "128030047569936";
  "128030047570016" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "128030047569456" -> "128030047570016";
  "128030047570096" [label="string", shape="box", fillcolor="lightgray"];
  "128030047570016" -> "128030047570096";
  "128030047570176" [label="identifier: true", shape="box", fillcolor="Honeydew"];
  "128030047570016" -> "128030047570176";
  "128030047570256" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "128030047569456" -> "128030047570256";
  "128030047570336" [label="string", shape="box", fillcolor="lightgray"];
  "128030047570256" -> "128030047570336";
  "128030047570416" [label="delimited", shape="box", fillcolor="lightgray"];
  "128030047570256" -> "128030047570416";
  "128030047570496" [label="string", shape="box", fillcolor="lightgray"];
  "128030047570416" -> "128030047570496";
  "128030047570576" [label="string", shape="box", fillcolor="lightgray"];
  "128030047570416" -> "128030047570576";
  "128030047570656" [label="string", shape="box", fillcolor="lightgray"];
  "128030047570416" -> "128030047570656";
  "128030047570736" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "128030047569456" -> "128030047570736";
  "128030047570816" [label="string", shape="box", fillcolor="lightgray"];
  "128030047570736" -> "128030047570816";
  "128030047570896" [label="delimited", shape="box", fillcolor="lightgray"];
  "128030047570736" -> "128030047570896";
  "128030047570976" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "128030047570896" -> "128030047570976";
  "128030047571136" [label="string", shape="box", fillcolor="lightgray"];
  "128030047570976" -> "128030047571136";
  "128030047571296" [label="string", shape="box", fillcolor="lightgray"];
  "128030047570976" -> "128030047571296";
  "128030047571376" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "128030047570896" -> "128030047571376";
  "128030047571536" [label="string", shape="box", fillcolor="lightgray"];
  "128030047571376" -> "128030047571536";
  "128030047571696" [label="string", shape="box", fillcolor="lightgray"];
  "128030047571376" -> "128030047571696";
  "128030047571776" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "128030047570896" -> "128030047571776";
  "128030047571936" [label="string", shape="box", fillcolor="lightgray"];
  "128030047571776" -> "128030047571936";
  "128030047572096" [label="string", shape="box", fillcolor="lightgray"];
  "128030047571776" -> "128030047572096";
  "128030047572176" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "128030047569456" -> "128030047572176";
  "128030047572256" [label="string", shape="box", fillcolor="lightgray"];
  "128030047572176" -> "128030047572256";
  "128030047572336" [label="delimited", shape="box", fillcolor="lightgray"];
  "128030047572176" -> "128030047572336";
  "128030047572416" [label="delimited", shape="box", fillcolor="lightgray"];
  "128030047572336" -> "128030047572416";
  "128030047572576" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "128030047572416" -> "128030047572576";
  "128030047572736" [label="string", shape="box", fillcolor="lightgray"];
  "128030047572576" -> "128030047572736";
  "128030047572896" [label="string", shape="box", fillcolor="lightgray"];
  "128030047572576" -> "128030047572896";
  "128030047573136" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "128030047572416" -> "128030047573136";
  "128030047573296" [label="string", shape="box", fillcolor="lightgray"];
  "128030047573136" -> "128030047573296";
  "128030047573456" [label="string", shape="box", fillcolor="lightgray"];
  "128030047573136" -> "128030047573456";
  "128030047573696" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "128030047572416" -> "128030047573696";
  "128030047573856" [label="string", shape="box", fillcolor="lightgray"];
  "128030047573696" -> "128030047573856";
  "128030047574016" [label="number: 1960", shape="box", fillcolor="lightgoldenrodyellow"];
  "128030047573696" -> "128030047574016";
  "128030047574176" [label="delimited", shape="box", fillcolor="lightgray"];
  "128030047572336" -> "128030047574176";
  "128030047574336" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "128030047574176" -> "128030047574336";
  "128030047574496" [label="string", shape="box", fillcolor="lightgray"];
  "128030047574336" -> "128030047574496";
  "128030047574656" [label="string", shape="box", fillcolor="lightgray"];
  "128030047574336" -> "128030047574656";
  "128030047574896" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "128030047574176" -> "128030047574896";
  "128030047575056" [label="string", shape="box", fillcolor="lightgray"];
  "128030047574896" -> "128030047575056";
  "128030047575216" [label="string", shape="box", fillcolor="lightgray"];
  "128030047574896" -> "128030047575216";
  "128030047575456" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "128030047574176" -> "128030047575456";
  "128030047575616" [label="string", shape="box", fillcolor="lightgray"];
  "128030047575456" -> "128030047575616";
  "128030047575776" [label="number: 1949", shape="box", fillcolor="lightgoldenrodyellow"];
  "128030047575456" -> "128030047575776";
}
```

![Image generated for example](images/json_blob.png)
