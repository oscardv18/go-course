# GO COURSE

I was perform this course in freecodecamp, in this course it was develop an booking application for a conference about golang

## Basics Concepts

- Print message in console: 

`fmt.Println("message")`
- All our code must belong to a package, the first statement in go file must be: 

e.g `package main`

- Variable Declaration:

`var <varName>`

e.g `var hello = "String"`

- Constants Declaration:
  
`const <constName>`

e.g `const pi = 3.14`

#### What is the variables and Constants?
The variables are used for stored values of the different types, unlike them, constants cannot change their values once they have been declared.

## How are applications organized in go?
In go is very used the package concept, it is the way in which a project organized (this is becouse is the best way for maintenance an project), go has many built-in package, such as the fmt package. This package has a very common function, this is the function for print in console.

- Print with line break:

`fmt.Println("Hello World")`

- Print with values:

```go
var name = "Oscar"
fmt.Printf("Hello %v\n", name)
```

### How to import package in go?
```go
package main

import "fmt"

func main() {
    var name = "Oscar"
    fmt.Printf("Hello %v\n", name)
}
```

The statement `import "package"` is the way for import only one package, for various package is the this way:

```go
import (
    "package1"
    "package2"...
)
```

### What is function main?
Every file in go must have an entry point, this it is the proposite of the function main, in it the execution of all the functions and other parts of the application occurs

```go
func main() {
    ...
}
```

## Forms of print messages

