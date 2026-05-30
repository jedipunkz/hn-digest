# hn-digest

Hacker News stories are crawled, translated into Japanese with Google Translate,
and saved under `contents/YYYY-MM-DD/`.

## Local run

```sh
go run ./cmd/hn-digest --source algolia --since 24h --limit 0 --title-keywords "sre,devops,google cloud,gcp,ai,llm"
```

The GitHub Actions workflow runs once a day and can also be started manually
with workflow dispatch. No API key is required. Set the repository variable
`HN_DIGEST_TITLE_KEYWORDS` to override the default title filter.
