package main

import (
	"log"

	"github.com/karpov-dmitry-py/go-data-structures/internal/pkg/service"
)

func main() {
	var (
		linkedListService = service.NewLinkedListService()
	)

	for i := 1; i <= 100; i++ {
		linkedListService.InsertRight(int64(i))
	}

	first, _ := linkedListService.GetFirst()
	last, _ := linkedListService.GetLast()

	log.Printf("length: %d", linkedListService.GetLength())
	log.Printf("values: %v", linkedListService.GetValues())
	log.Printf("first: %d", first)
	log.Printf("last: %d", last)
}
