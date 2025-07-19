| Variable | Purpose | Default |
|----------|---------|---------|
| OLLAMA_DISABLE_WARMUP | Disable cache warm-up entirely | unset |
| OLLAMA_WARMUP_BYTES | Bytes to pre-fault (overrides auto-detect) | SLC size |
| OLLAMA_WARMUP_BUDGET_MS | Wall-clock budget for warm-up | adaptive |
| OLLAMA_WARMUP_SYNC | Force synchronous warm-up | unset |
| OLLAMA_WARMUP_VERBOSE | Log details to stderr | unset |
| OLLAMA_METRICS | Expose Prometheus `/metrics` on :2112 | unset |

