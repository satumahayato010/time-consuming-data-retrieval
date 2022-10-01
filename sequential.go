package main

import (
	"fmt"
	"time"
)

var productDB map[int]string

func getProduct() []string {
	r := make([]string, 0, 100)

	for id := 0; id < 50; id++ {
		res, ok := searchProduct(id)
		if ok {
			r = append(r, res)
		}
	}
	return r
}

func searchProduct(id int) (string, bool) {
	time.Sleep(time.Second)
	result, ok := productDB[id]
	return result, ok
}

func main() {
	productDB = make(map[int]string, 100)

	for i := 0; i < 100; i++ {
		productDB[i] = fmt.Sprintf("Product %d\n", i)
	}

	start := time.Now()
	fmt.Println(start)

	getProduct()
	fmt.Println(time.Now())

	end := time.Since(start)
	fmt.Println(end)
}
