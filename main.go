package main

// import (
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"

// 	"github.com/joho/godotenv"
// 	vegeta "github.com/tsenart/vegeta/lib"
// )

// func main() {
// 	godotenv.Load(".env")
// 	interuppt := make(chan os.Signal, 1)
// 	signal.Notify(interuppt, syscall.SIGINT, os.Interrupt)
// 	url := os.Getenv("URL")

// 	attacker := vegeta.NewAttacker()
// 	rate := vegeta.Rate{Freq: 100, Per: 1 * time.Second}

// 	duration := 10 * time.Second

// 	target := vegeta.Target{
// 		Method: http.MethodGet,
// 		URL:    url,
// 	}

// 	var metrics vegeta.Metrics
// 	results := attacker.Attack(target, rate, duration, "A")
// 	for {
// 		select {
// 		case <-interuppt:
// 			attacker.Stop()
// 		case result := <-results:
// 			metrics.Add(result)
// 		}

// 	}

// }

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
	godotenv.Load(".env")
	url := os.Getenv("URL")
	// Buat saluran untuk menerima sinyal OS
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Buat pengaturan penggempuran
	rate := vegeta.Rate{Freq: 100, Per: time.Second} // 100 requests per second
	duration := 10 * time.Second                     // Durasi pengujian 10 detik

	// Buat penyerang vegeta
	attacker := vegeta.NewAttacker()

	// Buat saluran untuk hasil serangan
	var metrics vegeta.Metrics
	results := attacker.Attack(a(url), rate, duration, "Load testing with Vegeta!")

	// Loop untuk mengumpulkan metrik hasil serangan
	for {
		select {
		case <-interrupt:
			// Jika sinyal SIGINT diterima, tutup penyerang
			attacker.Stop()
			return
		case result := <-results:
			metrics.Add(result)
		}
	}

	// Tampilkan statistik hasil serangan
	metrics.Close()
	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
	fmt.Printf("Mean latency: %s\n", metrics.Latencies.Mean)
	fmt.Printf("Requests per second: %.2f\n", metrics.Rate)
	fmt.Printf("Total requests: %d\n", metrics.Requests)
	fmt.Printf("Total duration: %s\n", metrics.Duration)
}

func a(url string) func(t *vegeta.Target) error {
	// Buat target yang akan diserang
	target := func(t *vegeta.Target) error {
		t.Method = http.MethodGet
		t.URL = url
		return nil
	}
	return target
}
