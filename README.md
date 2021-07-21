A simple concatenative programming language written in Go.
WIP.

It doesn't do much yet (see below), but if you want to try it you can download and run the `naiveconcat` binary, or you can clone this repo and `go run .` (assuming you have go installed).

TODOs:
- [ ] comments with `--`
- [ ] `let` binds word definitions in the scope of a surrounding quotation
- [ ] flow control with `if`

Status:
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
