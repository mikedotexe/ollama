package main

import (
	"context"
	"flag"
	"log"
	"os"
	"sort"
	"time"

	"github.com/ollama/ollama/api"
)

func percentile(d []time.Duration, p int) time.Duration {
	if len(d) == 0 {
		return 0
	}
	sort.Slice(d, func(i, j int) bool { return d[i] < d[j] })
	idx := (p*len(d)+99)/100 - 1
	if idx < 0 {
		idx = 0
	}
	if idx >= len(d) {
		idx = len(d) - 1
	}
	return d[idx]
}

func chat(ctx context.Context, prompt string) error {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		return err
	}
	req := &api.ChatRequest{
		Model:    "llama3.2",
		Messages: []api.Message{{Role: "user", Content: prompt}},
	}
	return client.Chat(ctx, req, func(api.ChatResponse) error { return nil })
}

func main() {
	const prompt = "Say the single word <EOS> and stop."

	warm := flag.Bool("warm", true, "enable warm-up")
	flag.Parse()

	if !*warm {
		os.Setenv("OLLAMA_DISABLE_WARMUP", "1")
	}

	ctx := context.Background()

	start := time.Now()
	_ = chat(ctx, prompt)
	prime := time.Since(start)

	var lats []time.Duration
	for i := 0; i < 25; i++ {
		t0 := time.Now()
		_ = chat(ctx, prompt)
		lats = append(lats, time.Since(t0))
	}

	p95 := percentile(lats, 95)
	log.Printf("prime=%v  p95=%v  (n=%d)", prime, p95, len(lats))
}
