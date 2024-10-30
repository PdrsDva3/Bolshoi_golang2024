package main

import (
	"fmt"
	"homework1/internal/pkg/storage"
)

func main() {
	s, err := storage.InitStorage()
	if err != nil {
		panic(err)
	}
	s.Set("key1", "value1")
	fmt.Println(s.GetKind("key1"))

	s.Set("key2", "222")
	fmt.Println(s.GetKind("key2"))

	s.Set("key3", "123.77")
	fmt.Println(s.GetKind("key3"))
}
