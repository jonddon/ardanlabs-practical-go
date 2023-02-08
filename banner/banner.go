package main

import "fmt"

func main() {
	banner("Go", 6)
	banner("G☺", 6)
	fmt.Println("g",isPalindrome("g"))
	fmt.Println("go",isPalindrome("go"))
	fmt.Println("gog",isPalindrome("gog"))
	fmt.Println("goog",isPalindrome("goog"))
}

func banner(text string, width int) {
	padding:= (width - len(text))/2
	for i:=0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i:=0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func isPalindrome(text string) bool {
	for i:=0; i< len(text)/2; i++{
		fmt.Println("front", text[i])
			fmt.Println("from back", text[len(text)-i-1])
		if text[i] != text[len(text)-i-1]{
			
			return false
		}
	}
	return true
}