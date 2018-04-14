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
func findUser(in []string) ([]string) {
	out    := make([]string, 0, len(in))
	reg, _ := regexp.Compile("(?i)kolesaweb")

	for _, user := range in {
		if match := reg.MatchString(user); match {
			out = append(out, user)
		}
	}

	return out
}
