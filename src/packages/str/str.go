package str

import "math/rand"

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func Random(n int) string {
	result := make([]rune, n)

	for i := range result {
		result[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(result)

}
