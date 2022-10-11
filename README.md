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

- Simple message: `fmt.Println("Hello World")`
- Message with variables: 
```go
var name = "Oscar"
fmt.Println("Hello ", name)
```
- Printing using the function Printf:
```go
var name = "Oscar"
fmt.Printf("Hello %v", name)
// The %v indicate where the value will appear in the message
```
Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler. For more information visit the [web] of the packages golang

[web]: https://pkg.go.dev/fmt

## Data Types in GO

Data types specify the type of data that a valid Go variable can hold. 

- **Basic type:** Numbers, strings, and booleans come under this category.
- **Aggregate type:** Array and structs come under this category.
- **Reference type:** Pointers, slices, maps, functions, and channels come under this category.
- **Interface type**

Each data type can do different things and behaves differently

### How to declare variables and constants with data types?

```go
var age int = 22 // variable of type integer
const dni int = 27455444 // const of type integer
const name string = "Oscar" // const of type string
```

The Basic Data Types are further categorized into three subcategories which are

- **Numbers**
- **Booleans**
- **Strings**

### Integers:
***Signed Integers***:

Signed integer types supported by Go is shown below.

- **int8** (8-bit signed integer whose range is -128 to 127)
- **int16** (16-bit signed integer whose range is -32768 to 32767)
- **int32** (32-bit signed integer whose range is -2147483648 to 2147483647)
- **int64** (64-bit signed integer whose range is -9223372036854775808 to 9223372036854775807)

***Unsigned Integers***:

- **uint8** (8-bit unsigned integer whose range is 0 to 255 )
- **uint16** (16-bit unsigned integer whose range is 0 to 65535 )
- **uint32** (32-bit unsigned integer whose range is 0 to 4294967295 )
- **uint64** (64-bit unsigned integer whose range is 0 to 18446744073709551615 )

Visit [here] for more information about integers

[here]: https://golangdocs.com/integers-in-golang

### Float-point numbers in golang
Go supports the IEEE-754 32-bit and 64-bit floating-point numbers extensively.

| Date-Type |              Description              |
|-----------|:-------------------------------------:|
| float32   | 32-bit IEEE 754 floating-point number |
| float64   | 64-bit IEEE 754 floating-point number |

For more information about float-points numbers visit this [site]

[site]: https://golangdocs.com/floating-point-numbers-in-golang