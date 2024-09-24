package main

import (
	"fmt"
	"strconv"
)

func InttoString() {
	var x int = 5
	conver := strconv.Itoa(x)
	fmt.Println(conver)

}
func StringToInt() {
	var name string = "123"
	final, error := strconv.Atoi(name)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(string(final))
	}
}

func IntToFloat() {
	var x int = 100
	var output float32 = float32(x)
	fmt.Println(output)
}
func convertion() {
	num := 20.23
	var out int = int(num)
	fmt.Println(out)
}
func Bytetorune() {
	var y byte = 'g'
	var con rune = rune(y)
	fmt.Println(con)
}
func RunetoByte() {
	var z rune = 's'
	var x byte = byte(z)
	fmt.Println(x)
}
func Stringtobyte() {
	name := "Codelangs"
	var input []byte = []byte(name)
	fmt.Println(input)
}
func TypeAsseration() {
	var x interface{} = 40
	value := x.(int)
	fmt.Println(value)
	//check value in type asseration using safe mode
	s, ok := x.(int)
	if ok {
		fmt.Println("string value is :", s)
	} else {
		fmt.Println("Type asseration failed")
	}
}

func SliceINttofloat() {
	var data = []int{10, 20, 30, 40}
	var y []float64
	for _, val := range data {
		y = append(y, float64(val))

	}
	fmt.Println(y)
}

func main() {
	//InttoString()
	//StringToInt()
	//IntToFloat()
	//convertion()
	//Bytetorune()
	//RunetoByte()
	//Stringtobyte()
	//TypeAsseration()
	SliceINttofloat()

}
