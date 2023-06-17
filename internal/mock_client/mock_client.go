package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	api "github.com/rlapenok/dumb_base/grpc_generate/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	Client()
}

func Client() {
	coon, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)

	}
	defer coon.Close()
	c := api.NewApiClient(coon)
	wg := sync.WaitGroup{}
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		//time.Sleep(500 * time.Millisecond)
		go func(b int) {
			defer wg.Done()
			defer fmt.Println("Go № " + strconv.Itoa(b) + "-DONE")
			v := "Start go № " + strconv.Itoa(b)
			fmt.Println(v)
			//time.Sleep(time.Duration(b) * time.Millisecond)
			resp, err := c.UpdateKeys(context.Background(), &api.NewKey{
				Key: fmt.Sprint(b),
			})
			fmt.Print("Resp go №" + strconv.Itoa(b) + " ")
			fmt.Printf("%+v \n", resp)
			if err != nil {
				fmt.Println(err)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println(time.Since(start).String())
}
