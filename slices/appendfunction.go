package slices

/*
The built-in append function is used to dynamically add elements to a slice:
If the underlying array is not large enough, append() will create a new underlying array and point the returned slice to it.
```
func append(slice []Type, elems ...Type) []Type
```

Example:
--------
```
slice = append(slice, oneThing)
slice = append(slice, firstThing, secondThing)
slice = append(slice, anotherSlice...)
```

*/

// Assignment

/*

Complete the getCostsByDay() function using the append() function.
It accepts a slice of cost structs and returns a slice of float64s, where each float64 represents the total cost for a day.

The length of the returned slice should be inferred from the highest numbered day field.

Some days may have multiple costs, while others may have none.

Include all actual and implied totals in the returned slice, starting with day '0'. Use the append() function to
dynamically increase the size of the returned slice when needed.

*/

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	totalCost := []float64{} // can use make function as well
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		for cost.day >= len(costs) {
			totalCost = append(totalCost, 0.0)
		}
		totalCost[cost.day] += cost.value
	}
	return totalCost
	// return nil
}
