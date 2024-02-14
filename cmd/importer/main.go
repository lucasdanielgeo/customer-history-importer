package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmsgprefix)
	log.SetPrefix(fmt.Sprintf("[%d] [customer-history-importer] ", os.Getpid()))
	godotenv.Load()
}

func main() {
	start := time.Now()
	timeout := 60

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	done := make(chan struct{})
	go func() {
		defer close(done)
		Run()
	}()

	select {
	case <-ctx.Done():
		log.Fatalf("[ERROR] importing process too long, cancelled after %vs", timeout)

	case <-done:
		elapsedTime := time.Since(start)
		log.Printf("[SUCCESS] Importing took: %vs\n", elapsedTime.Seconds())
	}
}

func Run() {}
