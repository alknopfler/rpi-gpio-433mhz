package main

import (
	"github.com/brian-armstrong/gpio"
	"fmt"
)


func main() {
	watcher := gpio.NewWatcher()
	watcher.AddPin(22)
	watcher.AddPin(27)
	defer watcher.Close()

	go func() {
		for {
			pin, value := watcher.Watch()
			fmt.Printf("read %d from gpio %d\n", value, pin)
		}
	}()
}
