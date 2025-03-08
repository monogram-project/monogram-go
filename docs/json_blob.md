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
  138840936614464["delimited"]:::custom_delimited;
  138840936614544["operator: :"]:::custom_operator;
  138840936614464 --> 138840936614544;
  138840936614624["string: person"]:::custom_string;
  138840936614544 --> 138840936614624;
  138840936614704["delimited"]:::custom_delimited;
  138840936614544 --> 138840936614704;
  138840936614784["operator: :"]:::custom_operator;
  138840936614704 --> 138840936614784;
  138840936614864["string: name"]:::custom_string;
  138840936614784 --> 138840936614864;
  138840936614944["string: Alice"]:::custom_string;
  138840936614784 --> 138840936614944;
  138840936615024["operator: :"]:::custom_operator;
  138840936614704 --> 138840936615024;
  138840936615104["string: age"]:::custom_string;
  138840936615024 --> 138840936615104;
  138840936615184["number: 25"]:::custom_number;
  138840936615024 --> 138840936615184;
  138840936615264["operator: :"]:::custom_operator;
  138840936614704 --> 138840936615264;
  138840936615344["string: isStudent"]:::custom_string;
  138840936615264 --> 138840936615344;
  138840936615424["identifier: true"]:::custom_identifier;
  138840936615264 --> 138840936615424;
  138840936615504["operator: :"]:::custom_operator;
  138840936614704 --> 138840936615504;
  138840936615584["string: skills"]:::custom_string;
  138840936615504 --> 138840936615584;
  138840936615664["delimited"]:::custom_delimited;
  138840936615504 --> 138840936615664;
  138840936615744["string: Python"]:::custom_string;
  138840936615664 --> 138840936615744;
  138840936615824["string: JavaScript"]:::custom_string;
  138840936615664 --> 138840936615824;
  138840936615904["string: SQL"]:::custom_string;
  138840936615664 --> 138840936615904;
  138840936615984["operator: :"]:::custom_operator;
  138840936614704 --> 138840936615984;
  138840936616064["string: address"]:::custom_string;
  138840936615984 --> 138840936616064;
  138840936616144["delimited"]:::custom_delimited;
  138840936615984 --> 138840936616144;
  138840936616224["operator: :"]:::custom_operator;
  138840936616144 --> 138840936616224;
  138840936616384["string: street"]:::custom_string;
  138840936616224 --> 138840936616384;
  138840936616544["string:<br/>123 Maple Street"]:::custom_string;
  138840936616224 --> 138840936616544;
  138840936616624["operator: :"]:::custom_operator;
  138840936616144 --> 138840936616624;
  138840936616784["string: city"]:::custom_string;
  138840936616624 --> 138840936616784;
  138840936616944["string: Exampleville"]:::custom_string;
  138840936616624 --> 138840936616944;
  138840936617024["operator: :"]:::custom_operator;
  138840936616144 --> 138840936617024;
  138840936617184["string: country"]:::custom_string;
  138840936617024 --> 138840936617184;
  138840936617344["string: Neverland"]:::custom_string;
  138840936617024 --> 138840936617344;
  138840936617424["operator: :"]:::custom_operator;
  138840936614704 --> 138840936617424;
  138840936617504["string:<br/>favoriteBooks"]:::custom_string;
  138840936617424 --> 138840936617504;
  138840936617584["delimited"]:::custom_delimited;
  138840936617424 --> 138840936617584;
  138840936617664["delimited"]:::custom_delimited;
  138840936617584 --> 138840936617664;
  138840936617824["operator: :"]:::custom_operator;
  138840936617664 --> 138840936617824;
  138840936617984["string: title"]:::custom_string;
  138840936617824 --> 138840936617984;
  138840936618144["string:<br/>To Kill a Mockingbird"]:::custom_string;
  138840936617824 --> 138840936618144;
  138840936618384["operator: :"]:::custom_operator;
  138840936617664 --> 138840936618384;
  138840936618544["string: author"]:::custom_string;
  138840936618384 --> 138840936618544;
  138840936618704["string: Harper Lee"]:::custom_string;
  138840936618384 --> 138840936618704;
  138840936618944["operator: :"]:::custom_operator;
  138840936617664 --> 138840936618944;
  138840936619104["string:<br/>yearPublished"]:::custom_string;
  138840936618944 --> 138840936619104;
  138840936619264["number: 1960"]:::custom_number;
  138840936618944 --> 138840936619264;
  138840936619424["delimited"]:::custom_delimited;
  138840936617584 --> 138840936619424;
  138840936619584["operator: :"]:::custom_operator;
  138840936619424 --> 138840936619584;
  138840936619744["string: title"]:::custom_string;
  138840936619584 --> 138840936619744;
  138840936619904["string: 1984"]:::custom_string;
  138840936619584 --> 138840936619904;
  138840936620144["operator: :"]:::custom_operator;
  138840936619424 --> 138840936620144;
  138840936620304["string: author"]:::custom_string;
  138840936620144 --> 138840936620304;
  138840936620464["string:<br/>George Orwell"]:::custom_string;
  138840936620144 --> 138840936620464;
  138840936620704["operator: :"]:::custom_operator;
  138840936619424 --> 138840936620704;
  138840936620864["string:<br/>yearPublished"]:::custom_string;
  138840936620704 --> 138840936620864;
  138840936621024["number: 1949"]:::custom_number;
  138840936620704 --> 138840936621024;

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
  "129497713412672" [label="delimited", shape="box", fillcolor="lightgray"];
  "129497713412752" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "129497713412672" -> "129497713412752";
  "129497713412832" [label="string", shape="box", fillcolor="lightgray"];
  "129497713412752" -> "129497713412832";
  "129497713412912" [label="delimited", shape="box", fillcolor="lightgray"];
  "129497713412752" -> "129497713412912";
  "129497713412992" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "129497713412912" -> "129497713412992";
  "129497713413072" [label="string", shape="box", fillcolor="lightgray"];
  "129497713412992" -> "129497713413072";
  "129497713413152" [label="string", shape="box", fillcolor="lightgray"];
  "129497713412992" -> "129497713413152";
  "129497713413232" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "129497713412912" -> "129497713413232";
  "129497713413312" [label="string", shape="box", fillcolor="lightgray"];
  "129497713413232" -> "129497713413312";
  "129497713413392" [label="number: 25", shape="box", fillcolor="lightgoldenrodyellow"];
  "129497713413232" -> "129497713413392";
  "129497713413472" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "129497713412912" -> "129497713413472";
  "129497713413552" [label="string", shape="box", fillcolor="lightgray"];
  "129497713413472" -> "129497713413552";
  "129497713413632" [label="identifier: true", shape="box", fillcolor="Honeydew"];
  "129497713413472" -> "129497713413632";
  "129497713413712" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "129497713412912" -> "129497713413712";
  "129497713413792" [label="string", shape="box", fillcolor="lightgray"];
  "129497713413712" -> "129497713413792";
  "129497713413872" [label="delimited", shape="box", fillcolor="lightgray"];
  "129497713413712" -> "129497713413872";
  "129497713413952" [label="string", shape="box", fillcolor="lightgray"];
  "129497713413872" -> "129497713413952";
  "129497713414032" [label="string", shape="box", fillcolor="lightgray"];
  "129497713413872" -> "129497713414032";
  "129497713414112" [label="string", shape="box", fillcolor="lightgray"];
  "129497713413872" -> "129497713414112";
  "129497713414192" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "129497713412912" -> "129497713414192";
  "129497713414272" [label="string", shape="box", fillcolor="lightgray"];
  "129497713414192" -> "129497713414272";
  "129497713414352" [label="delimited", shape="box", fillcolor="lightgray"];
  "129497713414192" -> "129497713414352";
  "129497713414432" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "129497713414352" -> "129497713414432";
  "129497713414592" [label="string", shape="box", fillcolor="lightgray"];
  "129497713414432" -> "129497713414592";
  "129497713414752" [label="string", shape="box", fillcolor="lightgray"];
  "129497713414432" -> "129497713414752";
  "129497713414832" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "129497713414352" -> "129497713414832";
  "129497713414992" [label="string", shape="box", fillcolor="lightgray"];
  "129497713414832" -> "129497713414992";
  "129497713415152" [label="string", shape="box", fillcolor="lightgray"];
  "129497713414832" -> "129497713415152";
  "129497713415232" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "129497713414352" -> "129497713415232";
  "129497713415392" [label="string", shape="box", fillcolor="lightgray"];
  "129497713415232" -> "129497713415392";
  "129497713415552" [label="string", shape="box", fillcolor="lightgray"];
  "129497713415232" -> "129497713415552";
  "129497713415632" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "129497713412912" -> "129497713415632";
  "129497713415712" [label="string", shape="box", fillcolor="lightgray"];
  "129497713415632" -> "129497713415712";
  "129497713415792" [label="delimited", shape="box", fillcolor="lightgray"];
  "129497713415632" -> "129497713415792";
  "129497713415872" [label="delimited", shape="box", fillcolor="lightgray"];
  "129497713415792" -> "129497713415872";
  "129497713416032" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "129497713415872" -> "129497713416032";
  "129497713416192" [label="string", shape="box", fillcolor="lightgray"];
  "129497713416032" -> "129497713416192";
  "129497713416352" [label="string", shape="box", fillcolor="lightgray"];
  "129497713416032" -> "129497713416352";
  "129497713416592" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "129497713415872" -> "129497713416592";
  "129497713416752" [label="string", shape="box", fillcolor="lightgray"];
  "129497713416592" -> "129497713416752";
  "129497713416912" [label="string", shape="box", fillcolor="lightgray"];
  "129497713416592" -> "129497713416912";
  "129497713417152" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "129497713415872" -> "129497713417152";
  "129497713417312" [label="string", shape="box", fillcolor="lightgray"];
  "129497713417152" -> "129497713417312";
  "129497713417472" [label="number: 1960", shape="box", fillcolor="lightgoldenrodyellow"];
  "129497713417152" -> "129497713417472";
  "129497713417632" [label="delimited", shape="box", fillcolor="lightgray"];
  "129497713415792" -> "129497713417632";
  "129497713417792" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "129497713417632" -> "129497713417792";
  "129497713417952" [label="string", shape="box", fillcolor="lightgray"];
  "129497713417792" -> "129497713417952";
  "129497713418112" [label="string", shape="box", fillcolor="lightgray"];
  "129497713417792" -> "129497713418112";
  "129497713418352" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "129497713417632" -> "129497713418352";
  "129497713418512" [label="string", shape="box", fillcolor="lightgray"];
  "129497713418352" -> "129497713418512";
  "129497713418672" [label="string", shape="box", fillcolor="lightgray"];
  "129497713418352" -> "129497713418672";
  "129497713418912" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "129497713417632" -> "129497713418912";
  "129497713419072" [label="string", shape="box", fillcolor="lightgray"];
  "129497713418912" -> "129497713419072";
  "129497713419232" [label="number: 1949", shape="box", fillcolor="lightgoldenrodyellow"];
  "129497713418912" -> "129497713419232";
}
```

![Image generated for example](images/factorial.png)
