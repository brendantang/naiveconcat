# naiveconcat

A simple concatenative programming language written in Go.
WIP.

Draws inspiration from Factor, Forth, Scheme. I leaned a lot on [Rob Pike's talk](https://www.youtube.com/watch?v=HxaD_trXwRE&t=2440s) about building a lexer.


## Overview

There is a **stack**, which is a first-in-last-out pile of values. Values like numbers and strings get pushed on top of the stack:

```
> 1 2 3
[1 2 3]

> "foo" 
[1 2 3 "foo"]

> 100 -- a `--` starts a comment, input is ignored until end of line
[1 2 3 "foo" 100]
```

A **procedure** is a value that manipulates the stack (and the dictionary), and may perform side effects like printing output.
Here are examples of procedures for addition, multiplication, and printing output.

```
> 3 3 +
[6]

> 2
[6 2]

> *
[12]

> say -- pop the top value and print it
12
[]
```

The **dictionary** binds words to definitions. A **word** is a value that defines another value. To evaluate a word, substitute it for its definition and evaluate that.

```
> 100 "a-cool-number" 
[100 "a-cool-number"]

> let
[]

> a-cool-number
[100]
```

The word `let` pops a string and another value, and adds a word to the dictionary using the string as the word's name and the other value as its definition.

A **quotation** is a value which contains other values. Evaluating a quotation just pushes it on the stack, which makes it useful to "wrap" words and procedures that you aren't ready to evaluate yet.

```
> 2 3 {+}  -- the word `+` is wrapped in a quotation
[2 3 {+}]

> apply  -- the word `apply` 'unwraps' a quotation
[5]

> {1 +}  -- you can wrap more than one value in a quotation
[5 {1 +}]

> apply
[6]

> {1 +} "increment" let  -- save {1 +} as the definition for the word `increment`
[6]

> increment  -- evaluating `increment` pushes the quotation on the stack
[6 {1 +}]

> apply  -- it still needs to be applied
[7]
```

Words like `+`, `*`, and `say` are defined by primitive procedures built into the language. But you can also create your own procedures using the word `lambda`. Lambda pops the top value off the stack and pushes a procedure that will apply that value.

```
> { "I'm a procedure!" say } lambda
[proc:{0x10c1ae0}]

> apply  -- you can apply a procedure on the stack
"I'm a procedure!"
[]

> { 1 + } lambda "increment-proc" let  -- but it's more useful to save it as a word
[]

> 1 increment-proc  -- evaluating a word with a procedure definition applies that procedure
[2]
```

Making a procedure and then saving it as a word definition is pretty important, so there is a keyword for it: `define`.

```
> {2 *} "double" define
[]

> 100 double
[200]
```
