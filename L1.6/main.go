package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Example1(quit chan bool, wg *sync.WaitGroup) {
	for {
		select {
		case <-quit: // Остановка с помощью доп канала
			fmt.Println("Example1 finished!\n")
			wg.Done()
			return

		default:
			fmt.Println("Example1 is working!")
			time.Sleep(1 * time.Second)
		}
	}

}
func Example2(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup) {

	for {
		select {
		case <-ctx.Done():// Остановка с помощью context по таймеру
			fmt.Println("Example2 is finished!\n")
			wg.Done()
			return
		default:
			fmt.Println("Example2 is working!")
			time.Sleep(1 * time.Second)
		}
	}

}
func Example3(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup) {

	for {
		select {
		case <-ctx.Done(): // Остановка с помощью context
			fmt.Println("Example3 is finished!\n")
			wg.Done()
			return
		default:
			fmt.Println("Example3 is working!")
			time.Sleep(1 * time.Second)
		}
	}

}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	quit := make(chan bool)
	fmt.Println("Example1 STARTED!\n")
	go Example1(quit, &wg)
	time.Sleep(3 * time.Second)
	quit <- true
	close(quit)

	fmt.Println("Example2 STARTED!\n")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) 
	go Example2(ctx, cancel, &wg)
	time.Sleep(4 * time.Second)

	ctx, cancel = context.WithCancel(context.Background()) 
	fmt.Println("Example3 STARTED!\n")
	go Example3(ctx, cancel, &wg)
	time.Sleep(3 * time.Second)
	cancel()

	wg.Wait()
}
