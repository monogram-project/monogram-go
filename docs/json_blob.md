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
  129159141608080["delimited"]:::custom_delimited;
  129159141608160["operator: :"]:::custom_operator;
  129159141608080 --> 129159141608160;
  129159141608240["string: person"]:::custom_string;
  129159141608160 --> 129159141608240;
  129159141608320["delimited"]:::custom_delimited;
  129159141608160 --> 129159141608320;
  129159141608400["operator: :"]:::custom_operator;
  129159141608320 --> 129159141608400;
  129159141608480["string: name"]:::custom_string;
  129159141608400 --> 129159141608480;
  129159141608560["string: Alice"]:::custom_string;
  129159141608400 --> 129159141608560;
  129159141608640["operator: :"]:::custom_operator;
  129159141608320 --> 129159141608640;
  129159141608720["string: age"]:::custom_string;
  129159141608640 --> 129159141608720;
  129159141608800["number: 25"]:::custom_number;
  129159141608640 --> 129159141608800;
  129159141608880["operator: :"]:::custom_operator;
  129159141608320 --> 129159141608880;
  129159141608960["string: isStudent"]:::custom_string;
  129159141608880 --> 129159141608960;
  129159141609040["identifier: true"]:::custom_identifier;
  129159141608880 --> 129159141609040;
  129159141609120["operator: :"]:::custom_operator;
  129159141608320 --> 129159141609120;
  129159141609200["string: skills"]:::custom_string;
  129159141609120 --> 129159141609200;
  129159141609280["delimited"]:::custom_delimited;
  129159141609120 --> 129159141609280;
  129159141609360["string: Python"]:::custom_string;
  129159141609280 --> 129159141609360;
  129159141609440["string: JavaScript"]:::custom_string;
  129159141609280 --> 129159141609440;
  129159141609520["string: SQL"]:::custom_string;
  129159141609280 --> 129159141609520;
  129159141609600["operator: :"]:::custom_operator;
  129159141608320 --> 129159141609600;
  129159141609680["string: address"]:::custom_string;
  129159141609600 --> 129159141609680;
  129159141609760["delimited"]:::custom_delimited;
  129159141609600 --> 129159141609760;
  129159141609840["operator: :"]:::custom_operator;
  129159141609760 --> 129159141609840;
  129159141610000["string: street"]:::custom_string;
  129159141609840 --> 129159141610000;
  129159141610160["string:<br/>123 Maple Street"]:::custom_string;
  129159141609840 --> 129159141610160;
  129159141610240["operator: :"]:::custom_operator;
  129159141609760 --> 129159141610240;
  129159141610400["string: city"]:::custom_string;
  129159141610240 --> 129159141610400;
  129159141610560["string: Exampleville"]:::custom_string;
  129159141610240 --> 129159141610560;
  129159141610640["operator: :"]:::custom_operator;
  129159141609760 --> 129159141610640;
  129159141610800["string: country"]:::custom_string;
  129159141610640 --> 129159141610800;
  129159141610960["string: Neverland"]:::custom_string;
  129159141610640 --> 129159141610960;
  129159141611040["operator: :"]:::custom_operator;
  129159141608320 --> 129159141611040;
  129159141611120["string:<br/>favoriteBooks"]:::custom_string;
  129159141611040 --> 129159141611120;
  129159141611200["delimited"]:::custom_delimited;
  129159141611040 --> 129159141611200;
  129159141611280["delimited"]:::custom_delimited;
  129159141611200 --> 129159141611280;
  129159141611440["operator: :"]:::custom_operator;
  129159141611280 --> 129159141611440;
  129159141611600["string: title"]:::custom_string;
  129159141611440 --> 129159141611600;
  129159141611760["string:<br/>To Kill a Mockingbird"]:::custom_string;
  129159141611440 --> 129159141611760;
  129159141612000["operator: :"]:::custom_operator;
  129159141611280 --> 129159141612000;
  129159141612160["string: author"]:::custom_string;
  129159141612000 --> 129159141612160;
  129159141612320["string: Harper Lee"]:::custom_string;
  129159141612000 --> 129159141612320;
  129159141612560["operator: :"]:::custom_operator;
  129159141611280 --> 129159141612560;
  129159141612720["string:<br/>yearPublished"]:::custom_string;
  129159141612560 --> 129159141612720;
  129159141612880["number: 1960"]:::custom_number;
  129159141612560 --> 129159141612880;
  129159141613040["delimited"]:::custom_delimited;
  129159141611200 --> 129159141613040;
  129159141613200["operator: :"]:::custom_operator;
  129159141613040 --> 129159141613200;
  129159141613360["string: title"]:::custom_string;
  129159141613200 --> 129159141613360;
  129159141613520["string: 1984"]:::custom_string;
  129159141613200 --> 129159141613520;
  129159141613760["operator: :"]:::custom_operator;
  129159141613040 --> 129159141613760;
  129159141613920["string: author"]:::custom_string;
  129159141613760 --> 129159141613920;
  129159141614080["string:<br/>George Orwell"]:::custom_string;
  129159141613760 --> 129159141614080;
  129159141614320["operator: :"]:::custom_operator;
  129159141613040 --> 129159141614320;
  129159141614480["string:<br/>yearPublished"]:::custom_string;
  129159141614320 --> 129159141614480;
  129159141614640["number: 1949"]:::custom_number;
  129159141614320 --> 129159141614640;

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
  "138777980324416" [label="delimited", shape="box", fillcolor="lightgray"];
  "138777980324496" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "138777980324416" -> "138777980324496";
  "138777980324576" [label="string", shape="box", fillcolor="lightgray"];
  "138777980324496" -> "138777980324576";
  "138777980324656" [label="delimited", shape="box", fillcolor="lightgray"];
  "138777980324496" -> "138777980324656";
  "138777980324736" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "138777980324656" -> "138777980324736";
  "138777980324816" [label="string", shape="box", fillcolor="lightgray"];
  "138777980324736" -> "138777980324816";
  "138777980324896" [label="string", shape="box", fillcolor="lightgray"];
  "138777980324736" -> "138777980324896";
  "138777980324976" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "138777980324656" -> "138777980324976";
  "138777980325056" [label="string", shape="box", fillcolor="lightgray"];
  "138777980324976" -> "138777980325056";
  "138777980325136" [label="number: 25", shape="box", fillcolor="lightgoldenrodyellow"];
  "138777980324976" -> "138777980325136";
  "138777980325216" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "138777980324656" -> "138777980325216";
  "138777980325296" [label="string", shape="box", fillcolor="lightgray"];
  "138777980325216" -> "138777980325296";
  "138777980325376" [label="identifier: true", shape="box", fillcolor="Honeydew"];
  "138777980325216" -> "138777980325376";
  "138777980325456" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "138777980324656" -> "138777980325456";
  "138777980325536" [label="string", shape="box", fillcolor="lightgray"];
  "138777980325456" -> "138777980325536";
  "138777980325616" [label="delimited", shape="box", fillcolor="lightgray"];
  "138777980325456" -> "138777980325616";
  "138777980325696" [label="string", shape="box", fillcolor="lightgray"];
  "138777980325616" -> "138777980325696";
  "138777980325776" [label="string", shape="box", fillcolor="lightgray"];
  "138777980325616" -> "138777980325776";
  "138777980325856" [label="string", shape="box", fillcolor="lightgray"];
  "138777980325616" -> "138777980325856";
  "138777980325936" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "138777980324656" -> "138777980325936";
  "138777980326016" [label="string", shape="box", fillcolor="lightgray"];
  "138777980325936" -> "138777980326016";
  "138777980326096" [label="delimited", shape="box", fillcolor="lightgray"];
  "138777980325936" -> "138777980326096";
  "138777980326176" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "138777980326096" -> "138777980326176";
  "138777980326336" [label="string", shape="box", fillcolor="lightgray"];
  "138777980326176" -> "138777980326336";
  "138777980326496" [label="string", shape="box", fillcolor="lightgray"];
  "138777980326176" -> "138777980326496";
  "138777980326576" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "138777980326096" -> "138777980326576";
  "138777980326736" [label="string", shape="box", fillcolor="lightgray"];
  "138777980326576" -> "138777980326736";
  "138777980326896" [label="string", shape="box", fillcolor="lightgray"];
  "138777980326576" -> "138777980326896";
  "138777980326976" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "138777980326096" -> "138777980326976";
  "138777980327136" [label="string", shape="box", fillcolor="lightgray"];
  "138777980326976" -> "138777980327136";
  "138777980327296" [label="string", shape="box", fillcolor="lightgray"];
  "138777980326976" -> "138777980327296";
  "138777980327376" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "138777980324656" -> "138777980327376";
  "138777980327456" [label="string", shape="box", fillcolor="lightgray"];
  "138777980327376" -> "138777980327456";
  "138777980327536" [label="delimited", shape="box", fillcolor="lightgray"];
  "138777980327376" -> "138777980327536";
  "138777980327616" [label="delimited", shape="box", fillcolor="lightgray"];
  "138777980327536" -> "138777980327616";
  "138777980327776" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "138777980327616" -> "138777980327776";
  "138777980327936" [label="string", shape="box", fillcolor="lightgray"];
  "138777980327776" -> "138777980327936";
  "138777980328096" [label="string", shape="box", fillcolor="lightgray"];
  "138777980327776" -> "138777980328096";
  "138777980328336" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "138777980327616" -> "138777980328336";
  "138777980328496" [label="string", shape="box", fillcolor="lightgray"];
  "138777980328336" -> "138777980328496";
  "138777980328656" [label="string", shape="box", fillcolor="lightgray"];
  "138777980328336" -> "138777980328656";
  "138777980328896" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "138777980327616" -> "138777980328896";
  "138777980329056" [label="string", shape="box", fillcolor="lightgray"];
  "138777980328896" -> "138777980329056";
  "138777980329216" [label="number: 1960", shape="box", fillcolor="lightgoldenrodyellow"];
  "138777980328896" -> "138777980329216";
  "138777980329376" [label="delimited", shape="box", fillcolor="lightgray"];
  "138777980327536" -> "138777980329376";
  "138777980329536" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "138777980329376" -> "138777980329536";
  "138777980329696" [label="string", shape="box", fillcolor="lightgray"];
  "138777980329536" -> "138777980329696";
  "138777980329856" [label="string", shape="box", fillcolor="lightgray"];
  "138777980329536" -> "138777980329856";
  "138777980330096" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "138777980329376" -> "138777980330096";
  "138777980330256" [label="string", shape="box", fillcolor="lightgray"];
  "138777980330096" -> "138777980330256";
  "138777980330416" [label="string", shape="box", fillcolor="lightgray"];
  "138777980330096" -> "138777980330416";
  "138777980330656" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "138777980329376" -> "138777980330656";
  "138777980330816" [label="string", shape="box", fillcolor="lightgray"];
  "138777980330656" -> "138777980330816";
  "138777980330976" [label="number: 1949", shape="box", fillcolor="lightgoldenrodyellow"];
  "138777980330656" -> "138777980330976";
}
```

![Image generated for example](images/factorial.png)
