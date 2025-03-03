from railroad import Diagram, Choice, Comment, NonTerminal, Sequence, Terminal, Skip, Optional, ZeroOrMore, OneOrMore
import sys
import pymustache
import io
import re
import argparse
from pathlib import Path

ALLOWED_NAMES = {
    "Comment": Comment,
    "Diagram": Diagram, 
    "Choice": Choice, 
    "Comment": Comment, 
    "NonTerminal": NonTerminal, 
    "Sequence": Sequence, 
    "Terminal": Terminal,
    "Skip": Skip,
    "Optional": Optional,
    "ZeroOrMore": ZeroOrMore,
    "OneOrMore": OneOrMore,
    "undefined": 'None'
}

def eval_expression(input_string, allowed_names = {}):
    code = compile(input_string, "<string>", "eval")
    for name in code.co_names:
        if name not in allowed_names:
            raise NameError(f"Use of {name} not allowed")
    return eval(code, {"__builtins__": {}}, allowed_names)


mustache_template = """
<!DOCTYPE html>
<html>
<head>
<style type="text/css">
body {
    background-color: hsl(30,20%,95%);
}
svg.railroad-diagram {
    background-color: hsl(30,20%,95%);
}
svg.railroad-diagram path {
    stroke-width: 3;
    stroke: black;
    fill: rgba(0,0,0,0);
}
svg.railroad-diagram text {
    font: bold 14px monospace;
    text-anchor: middle;
    white-space: pre;
}
svg.railroad-diagram text.diagram-text {
    font-size: 12px;
}
svg.railroad-diagram text.diagram-arrow {
    font-size: 16px;
}
svg.railroad-diagram text.label {
    text-anchor: start;
}
svg.railroad-diagram text.comment {
    font: italic 12px monospace;
}
svg.railroad-diagram g.non-terminal text {
    font-style: italic;
}
svg.railroad-diagram rect {
    stroke-width: 3;
    stroke: black;
    fill: hsl(120,100%,90%);
}
svg.railroad-diagram rect.group-box {
    stroke: gray;
    stroke-dasharray: 10 5;
    fill: none;
}
svg.railroad-diagram path.diagram-text {
    stroke-width: 3;
    stroke: black;
    fill: white;
    cursor: help;
}
svg.railroad-diagram g.diagram-text:hover path.diagram-text {
    fill: #eee;
}
</style>
<title>{{TITLE}}</title>
</head>

<body>

{{#DIAGRAMS}}
<div>{{{svg}}}</div>
{{/DIAGRAMS}}

</body>
</html>
"""

class PrintTabatkinsGrammarAsHtml:

    def __init__(self):
        parser = argparse.ArgumentParser(description="Convert Tabatkins grammar to HTML.")
        parser.add_argument("--input", type=Path, required=True, help="Path to the input Tabatkins grammar file.")
        parser.add_argument("--title", type=str, required=True, help="Title of the HTML document.")
        args = parser.parse_args()
        self._args = args

    def print_tabatkins_grammar_as_html(self):
        print( pymustache.render( 
            mustache_template, 
            { 
                'TITLE': self._args.title,
                'DIAGRAMS': [ { 'svg': svg } for svg in self.svg_items() ]
            } 
        ) )

    def parse_tabatkins_file(self):
        dump = []
        current = []
        current_line_num = 1
        with self._args.input.open('r') as file:
            for (line_num, line) in enumerate( file ):
                if line.startswith('//'):
                    if current:
                        dump.append((current_line_num, current ))
                        current_line_num = line_num
                        current = []
                m = re.match( '''(.*)//[^'"]*$''', line )
                if m:
                    current.append( m.group(1) )
                else:
                    current.append( line )
            if current:
                dump.append((current_line_num, current ))
        return map( lambda L: ( L[0], ''.join(L[1]) ), dump )

    def diagram2string( self, diag ):
        with io.StringIO() as buffer:
            diag.writeSvg(buffer.write)   
            return buffer.getvalue()

    def svg_items(self):
        for ( line_num, diag_text ) in self.parse_tabatkins_file():
            print( f'Parsing expression on line {line_num}', file=sys.stderr )
            yield self.diagram2string( eval_expression(diag_text, ALLOWED_NAMES) )

if __name__ == "__main__":
    PrintTabatkinsGrammarAsHtml().print_tabatkins_grammar_as_html()
