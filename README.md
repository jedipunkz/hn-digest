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

## Summaries

`cmd/summarize` picks the top stories of a day and generates Japanese titles and
summaries into `summaries/YYYY-MM-DD.json`.

GitHub Models was fully retired on 2026-07-30, so the summariser now talks to any
**OpenAI-compatible** chat completions API, configured through the environment:

| Variable | Required | Default |
|----------|----------|---------|
| `SUMMARIZE_API_KEY` | yes | — |
| `SUMMARIZE_API_BASE` | no | `https://generativelanguage.googleapis.com/v1beta/openai` |
| `SUMMARIZE_MODEL` | no | `gemini-3.5-flash-lite` |

Setup:

1. Get an API key from a provider with a free tier (the default targets the
   Gemini API; Groq, OpenRouter and OpenAI work by overriding the base URL and
   model).
2. Add it as the repository secret `SUMMARIZE_API_KEY`.
3. Optionally set the repository variables `SUMMARIZE_API_BASE` and
   `SUMMARIZE_MODEL` to use a different provider.

Local run:

```sh
SUMMARIZE_API_KEY=<key> go run ./cmd/summarize -date 2026-08-14
```

The workflow runs hourly, so summaries already present in the day's JSON file are
reused and only new stories cost an API call. Pass `-refresh` to regenerate every
summary, and `-min-interval` to change the delay between calls (default 4s, which
keeps a burst inside a typical 15 req/min free-tier limit).

## Protected RSS feed

`/rss.xml` is gated by `middleware.js` (Vercel Routing Middleware). It returns
`403` unless the request carries the correct `?token=` query parameter, which is
compared against the `RSS_TOKEN` environment variable.

Setup:

1. Set the env var on Vercel (Production scope at minimum):

   ```sh
   vercel env add RSS_TOKEN production
   ```

   Use a long random value, e.g. `openssl rand -hex 32`.

2. Subscribe with the tokenized URL in your RSS reader:

   ```
   https://hn-digest-psi.vercel.app/rss.xml?token=<RSS_TOKEN>
   ```

Rotating the token: update `RSS_TOKEN` on Vercel, redeploy, and re-subscribe with
the new URL. The old URL stops working immediately.

Note: only the RSS feed is protected. The HTML site at `/` stays public. To also
gate the site, extend `config.matcher` in `middleware.js`.
