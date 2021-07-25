A simple concatenative programming language written in Go.
WIP.

It doesn't do much yet (see below), but if you want to try it you can download and run the `naiveconcat` binary, or you can clone this repo and `go run .` (assuming you have go installed).

TODOs:
- [x] comments with `--`
- [x] `define` binds word definitions in the scope of a surrounding quotation
- [x] flow control with `then`
- [ ] should `apply` be implicit? separate out `define` and `let`?
- [ ] import definitions from files
- [ ] "standard library" — iteration especially
- [ ] tail call optimization would be nice, no idea how difficult to build

Status:
- 2021-07-24 - Quotations are closures—words you define inside a quotation are local to that quotation. Comment out to end of line with `--`.
  ```
  > 3 "x" define
  []
  
  > x  -- global definition for x is 3.
  [3]
  
  > { 2 "x" define x { 1 "x" define x } } apply apply -- redefine x in nested quotations
  [2 1]
  
  > x  -- global definition remains
  [2 1 3]
  ```


- 2021-07-20 - Quote values, define words.
  ```
  > {2 2 +}
  [{2 2 +}]
  
  > apply
  [4]
  
  > {dup *} "sq" define
  [4]
  
  > sq
  [4 {dup *}]
  
  > apply
  [16]
  ```
- 2021-07-11 - Addition, subtraction, multiplication, and division.
  ```
  > 2 2 *
  [4]
  
  > 24
  [4 24]
  
  > 4 +
  [4 28]
  
  > /
  [0.14285714285714285] 
  ```
