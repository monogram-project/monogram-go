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
  129624859329856["delimited"]:::custom_delimited;
  129624859329936["operator: :"]:::custom_operator;
  129624859329856 --> 129624859329936;
  129624859330016["string: person"]:::custom_string;
  129624859329936 --> 129624859330016;
  129624859330096["delimited"]:::custom_delimited;
  129624859329936 --> 129624859330096;
  129624859330176["operator: :"]:::custom_operator;
  129624859330096 --> 129624859330176;
  129624859330256["string: name"]:::custom_string;
  129624859330176 --> 129624859330256;
  129624859330336["string: Alice"]:::custom_string;
  129624859330176 --> 129624859330336;
  129624859330416["operator: :"]:::custom_operator;
  129624859330096 --> 129624859330416;
  129624859330496["string: age"]:::custom_string;
  129624859330416 --> 129624859330496;
  129624859330576["number: 25"]:::custom_number;
  129624859330416 --> 129624859330576;
  129624859330656["operator: :"]:::custom_operator;
  129624859330096 --> 129624859330656;
  129624859330736["string: isStudent"]:::custom_string;
  129624859330656 --> 129624859330736;
  129624859330816["identifier: true"]:::custom_identifier;
  129624859330656 --> 129624859330816;
  129624859330896["operator: :"]:::custom_operator;
  129624859330096 --> 129624859330896;
  129624859330976["string: skills"]:::custom_string;
  129624859330896 --> 129624859330976;
  129624859331056["delimited"]:::custom_delimited;
  129624859330896 --> 129624859331056;
  129624859331136["string: Python"]:::custom_string;
  129624859331056 --> 129624859331136;
  129624859331216["string: JavaScript"]:::custom_string;
  129624859331056 --> 129624859331216;
  129624859331296["string: SQL"]:::custom_string;
  129624859331056 --> 129624859331296;
  129624859331376["operator: :"]:::custom_operator;
  129624859330096 --> 129624859331376;
  129624859331456["string: address"]:::custom_string;
  129624859331376 --> 129624859331456;
  129624859331536["delimited"]:::custom_delimited;
  129624859331376 --> 129624859331536;
  129624859331616["operator: :"]:::custom_operator;
  129624859331536 --> 129624859331616;
  129624859331776["string: street"]:::custom_string;
  129624859331616 --> 129624859331776;
  129624859331936["string:<br/>123 Maple Street"]:::custom_string;
  129624859331616 --> 129624859331936;
  129624859332016["operator: :"]:::custom_operator;
  129624859331536 --> 129624859332016;
  129624859332176["string: city"]:::custom_string;
  129624859332016 --> 129624859332176;
  129624859332336["string: Exampleville"]:::custom_string;
  129624859332016 --> 129624859332336;
  129624859332416["operator: :"]:::custom_operator;
  129624859331536 --> 129624859332416;
  129624859332576["string: country"]:::custom_string;
  129624859332416 --> 129624859332576;
  129624859332736["string: Neverland"]:::custom_string;
  129624859332416 --> 129624859332736;
  129624859332816["operator: :"]:::custom_operator;
  129624859330096 --> 129624859332816;
  129624859332896["string:<br/>favoriteBooks"]:::custom_string;
  129624859332816 --> 129624859332896;
  129624859332976["delimited"]:::custom_delimited;
  129624859332816 --> 129624859332976;
  129624859333056["delimited"]:::custom_delimited;
  129624859332976 --> 129624859333056;
  129624859333216["operator: :"]:::custom_operator;
  129624859333056 --> 129624859333216;
  129624859333376["string: title"]:::custom_string;
  129624859333216 --> 129624859333376;
  129624859333536["string:<br/>To Kill a Mockingbird"]:::custom_string;
  129624859333216 --> 129624859333536;
  129624859333776["operator: :"]:::custom_operator;
  129624859333056 --> 129624859333776;
  129624859333936["string: author"]:::custom_string;
  129624859333776 --> 129624859333936;
  129624859334096["string: Harper Lee"]:::custom_string;
  129624859333776 --> 129624859334096;
  129624859334336["operator: :"]:::custom_operator;
  129624859333056 --> 129624859334336;
  129624859334496["string:<br/>yearPublished"]:::custom_string;
  129624859334336 --> 129624859334496;
  129624859334656["number: 1960"]:::custom_number;
  129624859334336 --> 129624859334656;
  129624859334816["delimited"]:::custom_delimited;
  129624859332976 --> 129624859334816;
  129624859334976["operator: :"]:::custom_operator;
  129624859334816 --> 129624859334976;
  129624859335136["string: title"]:::custom_string;
  129624859334976 --> 129624859335136;
  129624859335296["string: 1984"]:::custom_string;
  129624859334976 --> 129624859335296;
  129624859335536["operator: :"]:::custom_operator;
  129624859334816 --> 129624859335536;
  129624859335696["string: author"]:::custom_string;
  129624859335536 --> 129624859335696;
  129624859335856["string:<br/>George Orwell"]:::custom_string;
  129624859335536 --> 129624859335856;
  129624859336096["operator: :"]:::custom_operator;
  129624859334816 --> 129624859336096;
  129624859336256["string:<br/>yearPublished"]:::custom_string;
  129624859336096 --> 129624859336256;
  129624859336416["number: 1949"]:::custom_number;
  129624859336096 --> 129624859336416;

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
  "131766334637296" [label="delimited", shape="box", fillcolor="lightgray"];
  "131766334637376" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "131766334637296" -> "131766334637376";
  "131766334637456" [label="string", shape="box", fillcolor="lightgray"];
  "131766334637376" -> "131766334637456";
  "131766334637536" [label="delimited", shape="box", fillcolor="lightgray"];
  "131766334637376" -> "131766334637536";
  "131766334637616" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "131766334637536" -> "131766334637616";
  "131766334637696" [label="string", shape="box", fillcolor="lightgray"];
  "131766334637616" -> "131766334637696";
  "131766334637776" [label="string", shape="box", fillcolor="lightgray"];
  "131766334637616" -> "131766334637776";
  "131766334637856" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "131766334637536" -> "131766334637856";
  "131766334637936" [label="string", shape="box", fillcolor="lightgray"];
  "131766334637856" -> "131766334637936";
  "131766334638016" [label="number: 25", shape="box", fillcolor="lightgoldenrodyellow"];
  "131766334637856" -> "131766334638016";
  "131766334638096" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "131766334637536" -> "131766334638096";
  "131766334638176" [label="string", shape="box", fillcolor="lightgray"];
  "131766334638096" -> "131766334638176";
  "131766334638256" [label="identifier: true", shape="box", fillcolor="Honeydew"];
  "131766334638096" -> "131766334638256";
  "131766334638336" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "131766334637536" -> "131766334638336";
  "131766334638416" [label="string", shape="box", fillcolor="lightgray"];
  "131766334638336" -> "131766334638416";
  "131766334638496" [label="delimited", shape="box", fillcolor="lightgray"];
  "131766334638336" -> "131766334638496";
  "131766334638576" [label="string", shape="box", fillcolor="lightgray"];
  "131766334638496" -> "131766334638576";
  "131766334638656" [label="string", shape="box", fillcolor="lightgray"];
  "131766334638496" -> "131766334638656";
  "131766334638736" [label="string", shape="box", fillcolor="lightgray"];
  "131766334638496" -> "131766334638736";
  "131766334638816" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "131766334637536" -> "131766334638816";
  "131766334638896" [label="string", shape="box", fillcolor="lightgray"];
  "131766334638816" -> "131766334638896";
  "131766334638976" [label="delimited", shape="box", fillcolor="lightgray"];
  "131766334638816" -> "131766334638976";
  "131766334639056" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "131766334638976" -> "131766334639056";
  "131766334639216" [label="string", shape="box", fillcolor="lightgray"];
  "131766334639056" -> "131766334639216";
  "131766334639376" [label="string", shape="box", fillcolor="lightgray"];
  "131766334639056" -> "131766334639376";
  "131766334639456" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "131766334638976" -> "131766334639456";
  "131766334639616" [label="string", shape="box", fillcolor="lightgray"];
  "131766334639456" -> "131766334639616";
  "131766334639776" [label="string", shape="box", fillcolor="lightgray"];
  "131766334639456" -> "131766334639776";
  "131766334639856" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "131766334638976" -> "131766334639856";
  "131766334640016" [label="string", shape="box", fillcolor="lightgray"];
  "131766334639856" -> "131766334640016";
  "131766334640176" [label="string", shape="box", fillcolor="lightgray"];
  "131766334639856" -> "131766334640176";
  "131766334640256" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "131766334637536" -> "131766334640256";
  "131766334640336" [label="string", shape="box", fillcolor="lightgray"];
  "131766334640256" -> "131766334640336";
  "131766334640416" [label="delimited", shape="box", fillcolor="lightgray"];
  "131766334640256" -> "131766334640416";
  "131766334640496" [label="delimited", shape="box", fillcolor="lightgray"];
  "131766334640416" -> "131766334640496";
  "131766334640656" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "131766334640496" -> "131766334640656";
  "131766334640816" [label="string", shape="box", fillcolor="lightgray"];
  "131766334640656" -> "131766334640816";
  "131766334640976" [label="string", shape="box", fillcolor="lightgray"];
  "131766334640656" -> "131766334640976";
  "131766334641216" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "131766334640496" -> "131766334641216";
  "131766334641376" [label="string", shape="box", fillcolor="lightgray"];
  "131766334641216" -> "131766334641376";
  "131766334641536" [label="string", shape="box", fillcolor="lightgray"];
  "131766334641216" -> "131766334641536";
  "131766334641776" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "131766334640496" -> "131766334641776";
  "131766334641936" [label="string", shape="box", fillcolor="lightgray"];
  "131766334641776" -> "131766334641936";
  "131766334642096" [label="number: 1960", shape="box", fillcolor="lightgoldenrodyellow"];
  "131766334641776" -> "131766334642096";
  "131766334642256" [label="delimited", shape="box", fillcolor="lightgray"];
  "131766334640416" -> "131766334642256";
  "131766334642416" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "131766334642256" -> "131766334642416";
  "131766334642576" [label="string", shape="box", fillcolor="lightgray"];
  "131766334642416" -> "131766334642576";
  "131766334642736" [label="string", shape="box", fillcolor="lightgray"];
  "131766334642416" -> "131766334642736";
  "131766334642976" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "131766334642256" -> "131766334642976";
  "131766334643136" [label="string", shape="box", fillcolor="lightgray"];
  "131766334642976" -> "131766334643136";
  "131766334643296" [label="string", shape="box", fillcolor="lightgray"];
  "131766334642976" -> "131766334643296";
  "131766334643536" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "131766334642256" -> "131766334643536";
  "131766334643696" [label="string", shape="box", fillcolor="lightgray"];
  "131766334643536" -> "131766334643696";
  "131766334643856" [label="number: 1949", shape="box", fillcolor="lightgoldenrodyellow"];
  "131766334643536" -> "131766334643856";
}
```

![Image generated for example](images/factorial.png)
