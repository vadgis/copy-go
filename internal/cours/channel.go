package cours

import (
	"fmt"
	"time"
)
func MakeChannel(){
	//message := "hello"
	channel := make(chan string)

	go worker(channel, func() string { return lourde(10) })
	go worker(channel, func() string { return lourde(1) })
	go worker(channel, func() string { return lourde(3) })
	go worker(channel, func() string { return lourde(5) })


	i := 0
	for result := range channel{
		fmt.Println(result)
		if i >= 3{
			break
		}
		i++
	}
}
func lourde(delta time.Duration) string{
	time.Sleep(delta * time.Second)
	return fmt.Sprint("Hello World", delta)
}
func worker(channel chan string, callback func() string){
	channel <- callback()
}