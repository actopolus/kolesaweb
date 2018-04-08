package main

import "testing"

func BenchmarkFind(b *testing.B) {
	testData := []string{
		"user123:golangKolesaWeb",
		"user222:almaty2018",
		"user234:KolesaWeb2018",
		"user453:kolesawebQA",
		"user222:kolesAawebQA",
		"user100:almaty20183kolesaWeb",
		"user1777:kolesawebApril",
		"user2000:today",
		"user1123:gopher",
		"user1124:gopherKolesaWeb",
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		findUser(testData)
	}
}
