package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	map1 := make(map[byte]int)

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {

		s := input.Text()

		input.Scan()

		s1 := input.Text()
		s = strings.ToLower(s)
		s1 = strings.ToLower(s1)
		b1 := s1[0]

		for i := 0; i < len(s); i++ {
			map1[s[i]]++
		}
		fmt.Println(map1[b1])

	}

}
