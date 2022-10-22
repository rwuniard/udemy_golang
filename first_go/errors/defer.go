package main

import "fmt"

func main() {
	defer1()
	defer2()
	i := defer3()
	fmt.Println(i)
}

func defer1() {
	i := 0
	defer fmt.Println(i) // it will print i = 0, even though the execution is deferred.
	i++
	return
}

func defer2() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i) // it will print 3,2,1,0. It is executing from the last to the first.
	}
}

// This is rather new. The func doesn't need to init i, but it can just return i that is incremented inside the defer func.
// so it seems that you can define a var in the return function. It is called function Named return
// It will return 2
func defer3() (i int) {
	fmt.Println("defer3 ------")
	defer func() {
		i++
	}()
	return 1
}
