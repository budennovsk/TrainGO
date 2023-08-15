package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(get_id())
	id, s := get_id()
	fmt.Println(id, s)
	//fmt.Println(get_iter())
	res()
	res_for()

}

func get_id() (int, string) {
	a := 1
	b := "string"
	return a, b
}
func get_iter() int {
	sum := 0
	for sum < 10 {

		if sum == 2 {
			fmt.Println(sum, "dd")
		} else {
			fmt.Println(sum)
		}
		sum++
	}
	return sum

}

func res() {
	pe := 12
	a := &pe
	*a = 100
	fmt.Println(*a, pe)
	var liststr [3]string
	liststr[0] = "1111"
	liststr[1] = "2222"
	liststr[2] = "3333"
	createSlice := []string{}
	createSlice = append(createSlice, "sexjnd", "ggg")
	fmt.Println(createSlice)
	fmt.Println(cap(createSlice), "fffffff")
	for i := 0; i < len(liststr); i++ {
		fmt.Println(liststr[i])

	}

}
func res_for() {
	arr := []string{"a", "d", "c"}
	fmt.Println(arr)

	dict_map := map[string]int{
		"b": 11,
	}
	u := dict_map["b"]
	fmt.Println(u, "kkkkkk")
	fmt.Println(dict_map)
	fmt.Println(reflect.TypeOf(dict_map))
	for _, k := range dict_map {
		fmt.Print(k)
	}
}
