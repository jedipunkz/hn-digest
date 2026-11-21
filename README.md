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

`cmd/summarize` picks the top 30 stories of the day by `final_score` and writes
`summaries/YYYY-MM-DD.json`, which feeds the RSS feed. Titles are translated
with Google Translate; `summary_ja` is the lead of the translated article body,
cut at a sentence boundary. `image_url` carries the linked page's `og:image` (or
`twitter:image`) when it declares one, so RSS items ship with a thumbnail.

Articles crawled before image extraction existed have no image in their front
matter, and the crawler never revisits them, so `cmd/summarize` backfills those:
it fetches the page once, records the result in the front matter (`image: ""`
meaning "checked, none"), and never fetches it again. Pass `--date` (or the
workflow's `date` input) to backfill an earlier day.

### Abstractive summaries (optional, free)

Set the repository **secret** `GEMINI_API_KEY` to an
[AI Studio](https://aistudio.google.com/apikey) key and `cmd/summarize` writes a
short Japanese summary with the Gemini API free tier instead of the extractive
lead. The repository **variable** `GEMINI_MODEL` overrides the model (default
`gemini-3.5-flash-lite`).

It is deliberately optional and fail-soft, so the workflow cannot break on it:

- No key, a `429` rate limit, an unknown model or an outage -> the extractive
  lead is used for that article and the run still succeeds.
- Three consecutive failures stop the LLM for the rest of the run, so a revoked
  key costs one call, not thirty.
- Each generated summary is cached in the article's front matter (`summary_ja`)
  and committed, so an article is summarised once, not once per hourly run.
  That keeps usage near the number of new stories per day (tens), far below the
  free-tier daily cap.
- `--llm-interval` (default `4s`) paces calls under the free-tier
  requests-per-minute limit; `--llm-max 0` disables the LLM entirely.

Note: free-tier Gemini usage may be used by Google to improve their products.
Everything sent here is already-public Hacker News content.

Before this, summaries came from GitHub Models, which was retired
(`410 github_models_retirement_brownout`). See DESIGN.md for the history.

```sh
GEMINI_API_KEY=... go run ./cmd/summarize --date 2026-08-17 --out /tmp/summaries --n 5
```

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
