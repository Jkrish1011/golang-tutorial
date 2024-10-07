package structures

// car has been declared in the file structexample.go- so can be referenced here directly
// But when instantiating the embedded structure, we have to instantiate as if it's a embedded struct but can access the fields
// as if it's part of truck only
type truck struct {
	car     // here the fields of `car` are directly added here. so no need of nested structures. fields of car are added here automatically
	bedSize int
}
