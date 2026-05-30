# hn-digest

Hacker News stories are crawled, translated into Japanese with Google Translate,
and saved under `contents/YYYY-MM-DD/`.

## Local run

```sh
go run ./cmd/hn-digest --section topstories --since 24h --limit 0 --max-translations 50
```

The GitHub Actions workflow runs once a day and can also be started manually
with workflow dispatch. No API key is required.
