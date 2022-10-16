package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
	}

	fmt.Println(p1)
	fmt.Println(p2)

	embeddedStruct()
	anonymousStruct()
}

func embeddedStruct() {
	type secretAgent struct {
		person
		ltk bool
	}

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	fmt.Println(sa1)
	// You can omit the person. The person is promoted inside the embedded structure.
	// So you don't have to do sa1.person.first.
	fmt.Println(sa1.first, sa1.person.last, sa1.ltk)

	// But if you have a collision like in the example below
	// first is defined inside person and also in the secretAgent2 struct.
	type secretAgent2 struct {
		person
		first string
		ltk   bool
	}
	sa2 := secretAgent2{
		person: person{
			first: "James",
			last:  "Bond",
		},
		first: "collision",
		ltk:   true,
	}
	// Then you have to specify sa2.person.first to access the value of secretAgent.first.
	fmt.Println(sa2.first, sa2.person.first, sa2.person.last, sa2.ltk)
}

func anonymousStruct() {
	// type me struct {
	// 	first string
	// 	last string
	// 	age int
	// }

	// Instead creating a type like the code above, we can make it anonymous.
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   33,
	}

	fmt.Println(p1)
}
