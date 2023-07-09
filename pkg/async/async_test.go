package async

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestPromise_Await(t *testing.T) {
	fmt.Println("Let's start ...")

	promise := Exec(func(ctx context.Context) <-chan interface{} {
		return DoneAsync()
	})
	fmt.Println("Done is running ...")
	val, err := promise.Await()
	if err != nil {
		t.Error(err)
	}

	r := <-val
	fmt.Println(r)
	fmt.Println("Done ...")
	t.Log("succeeded")
	return

}

func DoneAsync() <-chan interface{} {
	ch := make(chan interface{})
	fmt.Println("Warming up ...")

	go func() {
		defer close(ch)
		time.Sleep(2 * time.Second)
		ch <- 1
	}()

	return ch

}

func TestAsync(t *testing.T) {
	t.Run("TestPromise_Await", TestPromise_Await)
}
