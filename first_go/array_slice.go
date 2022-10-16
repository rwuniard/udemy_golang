package main

import "fmt"

func main() {
	array()
	slice()
	usingRange()
	slicingSlice()
	appendToASlice()
	deleteFromASlice()
	sliceMake()
	multiDimentionalSlice()
}

// In Go we don't really use array, but we use slice instead
func array() {
	fmt.Println("Array")
	var x [5]int
	fmt.Println(x)
	x[3] = 42

	fmt.Println(x)
}

func slice() {
	fmt.Println("Slice")
	// Composite literal
	// x := type{values}
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	// a Slice allows you to group together VALUES of the same TYPE
}

func usingRange() {
	fmt.Println("Using Range")
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	fmt.Println("Length: ", len(x))
	fmt.Println("Capacity: ", cap(x))

	fmt.Println(x[0])
	fmt.Println(x[2])
	fmt.Println(x[3])
	// Looping using range
	fmt.Println("X:")
	for i, v := range x {
		fmt.Println("x[", i, "]", v)
	}
}

func slicingSlice() {
	fmt.Println("Slicing Slice")
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x[:])
	// THis will print from position 1 to the end.
	fmt.Println(x[1:])
	// THis will print from position 1 to the position prior to 3.
	fmt.Println(x[1:3])

}

func appendToASlice() {
	fmt.Println("Append to a Slice")
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x[:])
	x = append(x, 10, 11)
	fmt.Println(x[:])

	y := []int{333, 444, 555, 666}
	x = append(x, y...)
	fmt.Println(x)

	// We cannot append(x, y), since append signature is append([]T, elems ...T).
	// If it has append([]T, []T) then we can do append(x, y). Unfortunately it doesn't have that.
}

func deleteFromASlice() {
	fmt.Println("Delete from a slice")
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x[:])

	// This will take 4,5.
	// Then the second argument x[3:]... -> 8, 42. The ... needs to represents the element instead of []int.
	x = append(x[:2], x[3:]...)
	fmt.Println(x)
}

func sliceMake() {
	fmt.Println("sliceMake")
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println("length: ", len(x))
	fmt.Println("capacity: ", cap(x))

	x[0] = 7
	x[9] = 13
	fmt.Println(x)
	fmt.Println("length: ", len(x))
	fmt.Println("capacity: ", cap(x))

	// THis will give index out of range.
	// x[10] = 4
	// fmt.Println(x)
	// fmt.Println("length: ", len(x))
	// fmt.Println("capacity: ", cap(x))

	// but we can append and the length will become 11
	x = append(x, 23)
	fmt.Println(x)
	fmt.Println("length: ", len(x))
	fmt.Println("capacity: ", cap(x))

	// when you keep appending and exceeding the capacity, the slice will adjust its capacity (double it)
	// That adjusting the capacity may have a cost during runtime, since it needs to create and copy the array
	x = append(x, 33, 44, 55)
	fmt.Println(x)
	fmt.Println("length: ", len(x))
	fmt.Println("capacity: ", cap(x))

}

func multiDimentionalSlice() {
	fmt.Println("Multi Dimentional Slice")
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	mp := []string{"Miss", "MoneyPenny", "Strawberry", "Hazelnut"}

	xp := [][]string{jb, mp}

	fmt.Println(xp)
}
