package main

import (
	"fmt"
	"github.com/LifeValuable/Leetcode/trees"
)

func main() {
	fmt.Println("Запущено")
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(strs)
	fmt.Println(result)
}
