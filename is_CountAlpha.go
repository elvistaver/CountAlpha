package main

import "fmt"

func is_countAlpha(text string) int {

	count := 0

	for _, ch := range text {

		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {

			count++
		}
	}
	return count

}

func main() {

	text := "my name is Elvis de Vince Cater"

	count := is_countAlpha(text)

	fmt.Println(count)

}

