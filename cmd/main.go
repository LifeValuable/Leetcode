package main

import (
	"fmt"

	"github.com/LifeValuable/Leetcode/arrays"
)

func main() {
	fmt.Println("Запущено")
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := arrays.GroupAnagrams(strs)
	fmt.Println(result)
}
