package genou

import (
	"fmt"
	"time"
)

func MadeChannel(delta time.Duration) {
	channel := make(chan string)
	go workerer(channel, func() string { return lourder(delta) })
	go workerer(channel, func() string { return lourder(1) })
	go workerer(channel, func() string { return lourder(3) })
	go workerer(channel, func() string { return lourder(5) })

	/*i := 0
	for resultat := range channel {
		fmt.Println(resultat)
		if i >= 3 {
			break
		}
		i++
	}*/
}

func lourder(delta time.Duration) string {
	time.Sleep(delta * time.Second)
	return fmt.Sprint("Hello World", delta)
}

func workerer(channel chan string, callback func() string) {
	channel <- callback()
}
