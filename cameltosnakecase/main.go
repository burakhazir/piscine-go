package main

import "fmt"

func CamelToSnakeCase(s string) string {
	var word string
	for i := 0; i <= len(s)-1; i++ {
		if s[i] >= 'A' && s[i] <= 'Z' && s[i+1] >= 'A' && s[i+1] <= 'Z' {
			return s
		} else if i < len(s)-1 && (s[i] >= 'a' && s[i] <= 'z') && (s[i+1] >= 'A' && s[i+1] <= 'Z') {
			word += string(s[i]) + "_"
		} else {
			word += string(s[i])
		}
	}
	return word

}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
