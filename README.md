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
