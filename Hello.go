package main

import (
	"errors"
	"fmt"
	"math"
)

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println("Hello World")
	var x int = 5
	var y int = 10
	var sum int = x + 7

	p := 12
	l := p + sum

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(sum)
	fmt.Println(p)
	fmt.Println(l)

	if x > y {
		fmt.Println("x is larger than y")
	} else if x == y {
		fmt.Println("x and y are the same")
	} else {
		fmt.Println("x is smaller than y")
	}

	var a [5]int
	a[0] = x
	a[1] = y
	a[2] = sum
	a[3] = p
	a[4] = l
	fmt.Println(a)

	b := [3]int{14, 2, 5}
	fmt.Println(b)

	c := []int{0, 2, 5, 2}
	c = append(c, 13)
	fmt.Println(c)

	vertices := make(map[string]int)
	vertices["triangle"] = 3
	vertices["square"] = 4
	vertices["circle"] = 0
	fmt.Println(vertices)

	fmt.Println(vertices["triangle"])

	delete(vertices, "square")

	fmt.Println(vertices)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	for index, value := range a {
		fmt.Println("index", index, "value", value)
	}

	for key, value := range vertices {
		fmt.Println("key", key, "value", value)
	}

	result := sumNum(2, 3)
	fmt.Println(result)

	sqRt, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sqRt)
	}

	k := person{name: "colton", age: 29}
	fmt.Println(k)
	fmt.Println(k.name)

	point := 7
	fmt.Println(&point)

	inc(point)
	fmt.Println(point)

	incPoint(&point)
	fmt.Println(point)
}

func sumNum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}

func inc(x int) {
	x++
}

func incPoint(x *int) {
	*x++
}
