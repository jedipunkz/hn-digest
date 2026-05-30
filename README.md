# hn-digest

Hacker News stories are crawled, translated into Japanese with Google Translate,
and saved under `contents/YYYY-MM-DD/`.

## Local run

```sh
export GOOGLE_TRANSLATE_API_KEY=...
go run ./cmd/hn-digest --section topstories --limit 20
```

The GitHub Actions workflow expects a repository secret named
`GOOGLE_TRANSLATE_API_KEY`.
