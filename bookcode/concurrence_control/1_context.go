package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	before := time.Now()
	preCtx, _ := context.WithTimeout(ctx, 500*time.Millisecond)

	go func() {
		childCtx, _ := context.WithTimeout(preCtx, 300*time.Millisecond)
		select {
		case <-childCtx.Done():
			after := time.Now()
			fmt.Println("child during:", after.Sub(before).Milliseconds())
		}
	}()

	select {
	case <-preCtx.Done():
		after := time.Now()
		fmt.Println("pre during:", after.Sub(before).Milliseconds())
	}

	time.Sleep(5 * time.Second)
}
