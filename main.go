package main

import (
	"sync"
)

func main() {
	configs := GetConfiguration()
	var wg sync.WaitGroup
	for _, cmd := range configs {
		wg.Add(1)
		funcCmd := cmd
		go func() {
			CmdExecuter(funcCmd)
			wg.Done()
		}()
	}
	wg.Wait()
}
