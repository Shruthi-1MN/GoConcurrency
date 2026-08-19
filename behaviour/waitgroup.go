package behaviour

import (
	"fmt"
	"sync"
)

func WaitGroupBehaviour() {
	var wg sync.WaitGroup
	for i := range 5 {
		wg.Add(1)
		go func(){
			defer wg.Done()
			fmt.Println("i: ", i)
		}()
	}
	wg.Wait()
}

