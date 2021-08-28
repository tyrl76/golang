package main

import "fmt"

func main() {
	/* var A [10]int

	for i := 0; i < len(A); i++ {
		A[i] = i * i
	}
	fmt.Println(A) */

	s := "Hello World"

	k := "Hello 월드"

	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], "/")
		fmt.Print(string(s[i]), ", ")
	}

	fmt.Println()

	fmt.Println("len(k) = ", len(k)) // 6 + 3 + 3

	// utf-8 1~3 byte > 월, 드 : 3byte array는 1byte씩 끊어 읽으므로 이상한 문자로 출력

	for i := 0; i < len(k); i++ {
		fmt.Print(k[i], "/")
		fmt.Print(string(k[i]), ", ")
	}

	fmt.Println()

	k2 := []rune(k)                    // rune은 1~3byte씩 끊어 읽으므로
	fmt.Println("len(k2) = ", len(k2)) // 6 + 1 + 1

	for i := 0; i < len(k2); i++ {
		fmt.Print(k2[i], "/")
		fmt.Print(string(k2[i]), ", ")
	}
}
