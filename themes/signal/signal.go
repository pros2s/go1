package signal

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func job(ctx context.Context, sigStr string) {
	i := 0

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context finished: ", sigStr)
			return
		case <-time.After(1 * time.Second):
			fmt.Println(sigStr, "Seconds: ", i)
			i++
		}
	}
}

func TestSignal() {
	fmt.Println("PID", os.Getpid())

	ctxTerm, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM)

	job(ctxTerm, "SigTerm")

	fmt.Println("Finished correctly!")
}
