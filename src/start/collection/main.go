package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func main() {
	fmt.Println("d")

	mm := make(map[string]int, 8)
	mm["ee"] = 110
	mm["ee"] = 1102
	mm["ee01"] = 11023
	fmt.Println(mm)
	fmt.Printf("%#v\n", mm)
	fmt.Printf("%T\n", mm)

	delete(mm, "ee01")
	fmt.Println(mm)
	value, ok := mm["eee"]
	if ok {
		fmt.Println("ok", value)
	} else {
		fmt.Println("bu ok", value)
	}

	var mm01 = map[int]bool{
		1: true, 3: false,
	}

	fmt.Printf("%#v\n", mm01)
	fmt.Println(mm01)

	for k, v := range mm01 {
		fmt.Println(k, v)
	}

	sm := make(map[string]int, 50)
	for i := 0; i < 50; i++ {
		sm[fmt.Sprintf("stu%02d", i)] = rand.Intn(100)
	}

	keys := make([]string, 0, 50)
	// i := 0
	for k := range sm {
		// keys[i] = k
		// i++
		keys = append(keys, k)
	}

	sort.Strings(keys)

	fmt.Println(keys)

	for _, v := range keys {
		fmt.Println(v, sm[v])
	}

	var ms = make([]map[string]int, 8, 8)

	for i := range ms {
		ms[i] = make(map[string]int, 1)
		ms[0][fmt.Sprintf("stu%02d", i)] = rand.Intn(100)
	}
	fmt.Println(ms)

	for i := range ms {
		ms[i] = make(map[string]int, 1)
		ms[i][fmt.Sprintf("stu%02d", i)] = rand.Intn(100)
	}
	fmt.Println(ms)

	ssm := make(map[string][]int, 8)

	const count int = 8
	for i := 0; i < count; i++ {
		key := fmt.Sprintf("stu%02d", i)
		ssm[key] = make([]int, 0, 2)
		ssm[key] = append(ssm[key], 3, 4, i)
	}

	fmt.Println(ssm)

	phrase := "How do you do !"
	words := strings.Split(phrase, " ")
	fmt.Println(words)

	var wordSS = make(map[string]int, len(words))

	for _, v := range words {
		count := wordSS[v]
		wordSS[v] = count + 1
	}

	fmt.Println(wordSS)
}
