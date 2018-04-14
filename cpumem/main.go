package main

import (
	"regexp"
	"fmt"
)

func main() {
	input := []string{
		"user123:golangKolesaWeb",
		"user222:almaty2018",
		"user234:KolesaWeb2018",
		"user453:kolesawebQA",
		"user222:kolesAawebQA",
	}

	res := findUser(input)

	fmt.Println(res)
}

// Найти пользователей с kolesaweb в логине
func findUser(in []string) (out []string) {
	for _, user := range in {
		if match, _ := regexp.MatchString("(?i)kolesaweb", user); match {
			out = append(out, user)
		}
	}

	return out
}
