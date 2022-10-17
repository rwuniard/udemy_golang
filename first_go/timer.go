package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	fmt.Println("UnixNano:", time.Now().UnixNano())
	fmt.Println("time.ms:", time.Millisecond)
	fmt.Println("unixNano/ms: ", time.Now().UnixNano()/int64(time.Millisecond))
	time.Sleep(1 * time.Millisecond)
	fmt.Println("unixNano/ms: ", time.Now().UnixNano()/int64(time.Millisecond))
	fmt.Println("unixMs:", time.Now().UnixMilli())

	timeDiff()
	countIn1Millisecond()
}

func timeDiff() {
	start := time.Now().UnixNano() / int64(time.Millisecond)
	// do something
	time.Sleep(1 * time.Millisecond)
	end := time.Now().UnixNano() / int64(time.Millisecond)
	diff := end - start
	log.Printf("Duration(ms): %d", diff)
}

func countIn1Millisecond() {

	// I am trying to count as many as it can within 1 ms.
	// Need to learn concurrency first I guess.
	// start := time.Now().UnixNano()
	// end := time.Now().UnixNano() + 1000000

	// fmt.Println("start:", start)
	// fmt.Println("end:", end)
	// i := 0
	// for time.Now().UnixNano() >= end {
	// 	i++
	// }

	// fmt.Println("i:", i)

}
