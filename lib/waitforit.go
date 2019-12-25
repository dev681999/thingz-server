package lib

import (
	"log"
	"net"
	"sync"
	"time"
)

func wait(url string, wg *sync.WaitGroup) {
	conn, err := net.Dial("tcp", url)
	for err != nil {
		log.Println("Trying again for:", url)
		time.Sleep(time.Second * 1)
		conn, err = net.Dial("tcp", url)
	}

	log.Println("wait for", url, "complete")

	conn.Close()
	wg.Done()
}

// WaitFor waits for the url to be available
func WaitFor(urls []string) {
	if len(urls) == 0 {
		log.Println("No urls found")
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(urls))

	for _, url := range urls {
		log.Println("Waiting for:", url)
		go wait(url, wg)
	}

	wg.Wait()
}
