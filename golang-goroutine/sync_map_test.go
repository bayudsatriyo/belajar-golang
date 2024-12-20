package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
)


func AddToMap(data *sync.Map, value int, group *sync.WaitGroup)  {
	defer group.Done()
	
	data.Store(value, value)
}

func TestSyncMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go AddToMap(data, i, group)
	}

	group.Wait()

	data.Range(func(key, value any) bool {
		fmt.Println("Key = ", key, "Value = ", value)
		return true
	})
}