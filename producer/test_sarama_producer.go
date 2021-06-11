package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func push() error {
	rand.Seed(time.Now().Unix())
	msg_size := rand.Intn(2000)
	msg := make([]byte, msg_size)
	_, err := rand.Read(msg)
	if err != nil {
		return err
	}

	queue_err := toQueue("test", msg)
	if queue_err != nil {
		return queue_err
	}

	return nil
}

func main() {
	fmt.Println("runnable")

	// invoke the toQueue function here
	for i := 0; i < 100; i++ {
		if err := push(); err != nil {
			fmt.Println(err)
		}
	}

	os.Exit(0)
}
