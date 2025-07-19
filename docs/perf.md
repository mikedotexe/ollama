The `benchwarm` tool measures the effect of model warm-up on first-token latency.

```bash
go run ./cmd/benchwarm           # with warm-up
go run ./cmd/benchwarm -warm=0   # without warm-up

# Any `ollama` command accepts `--warmup` or `--no-warmup` to control warm-up
# behavior. This sets the `OLLAMA_DISABLE_WARMUP` environment variable.
```

The tool prints `prime` (the latency of the first request after loading) and `p95` across 25 subsequent requests. On Apple silicon machines warm-up should reduce both numbers.
