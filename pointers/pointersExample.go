package pointers

/*

x := 5 // stores are a location in RAM
y := x // y gets a new copy of value 5 in RAM
z := &x // z will store the address of x. any updates to z will reflect in x.
*z = 6 // will update the value of z and x both. We have updated the value of x without even accessing x.

Syntax:

```
var p *int
```

A pointer's zero value is nil
The & operator generates a pointer to its operand

```
myString := "hello"
myStringPointer = &myString
```

* dereference a pointer to gain access to the underlying value
```
fmt.Println(*myStringPtr) // reads myString through the pointer
*myStringPtr = "world" // set myString through the pointer
```

*/
