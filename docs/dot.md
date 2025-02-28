# Translation to Mermaid

## Example

Here's the factorial example from the first page.  

```txt
def f(n):
    if n <= 1:
        1
    else:
        n * f(n - 1)
    endif
enddef
```

We can target Graphviz's dot notations as an output format. 
And this is what the factorial example expands into:

```mermaid
digraph G {
  bgcolor="transparent";
  node [shape="box", style="filled", fontname="Ubuntu Mono"];
  "124112413528912" [label="form", shape="box", fillcolor="lightpink"];
  "124112413528992" [label="part: def", shape="box", fillcolor="#FFD8E1"];
  "124112413528912" -> "124112413528992";
  "124112413529072" [label="apply", shape="box", fillcolor="lightgreen"];
  "124112413528992" -> "124112413529072";
  "124112413529152" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "124112413529072" -> "124112413529152";
  "124112413529232" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "124112413529072" -> "124112413529232";
  "124112413529312" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "124112413529232" -> "124112413529312";
  "124112413529392" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "124112413528912" -> "124112413529392";
  "124112413529472" [label="form", shape="box", fillcolor="lightpink"];
  "124112413529392" -> "124112413529472";
  "124112413529552" [label="part: if", shape="box", fillcolor="#FFD8E1"];
  "124112413529472" -> "124112413529552";
  "124112413529632" [label="operator: <=", shape="box", fillcolor="#C0FFC0"];
  "124112413529552" -> "124112413529632";
  "124112413529712" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "124112413529632" -> "124112413529712";
  "124112413529792" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "124112413529632" -> "124112413529792";
  "124112413529872" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "124112413529472" -> "124112413529872";
  "124112413529952" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "124112413529872" -> "124112413529952";
  "124112413530032" [label="part: else", shape="box", fillcolor="#FFD8E1"];
  "124112413529472" -> "124112413530032";
  "124112413530112" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "124112413530032" -> "124112413530112";
  "124112413530192" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "124112413530112" -> "124112413530192";
  "124112413530272" [label="apply", shape="box", fillcolor="lightgreen"];
  "124112413530112" -> "124112413530272";
  "124112413530432" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "124112413530272" -> "124112413530432";
  "124112413530592" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "124112413530272" -> "124112413530592";
  "124112413530752" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "124112413530592" -> "124112413530752";
  "124112413530912" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "124112413530752" -> "124112413530912";
  "124112413531072" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "124112413530752" -> "124112413531072";
}
```