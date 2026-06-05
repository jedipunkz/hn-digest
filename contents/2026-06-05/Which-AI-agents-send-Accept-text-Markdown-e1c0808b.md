---
source: "https://acceptmarkdown.com/status"
hn_url: "https://news.ycombinator.com/item?id=48415869"
title: "Which AI agents send Accept: text/Markdown?"
article_title: "Which AI agents support Markdown content negotiation?"
author: "rickette"
captured_at: "2026-06-05T18:45:45Z"
capture_tool: "hn-digest"
hn_id: 48415869
score: 2
comments: 0
posted_at: "2026-06-05T17:45:29Z"
tags:
  - hacker-news
  - translated
---

# Which AI agents send Accept: text/Markdown?

- HN: [48415869](https://news.ycombinator.com/item?id=48415869)
- Source: [acceptmarkdown.com](https://acceptmarkdown.com/status)
- Score: 2
- Comments: 0
- Posted: 2026-06-05T17:45:29Z

## Translation

タイトル: どの AI エージェントが Accept: text/Markdown を送信しますか?
記事のタイトル: どの AI エージェントが Markdown コンテンツ ネゴシエーションをサポートしていますか?
説明: AI エージェントの実行マトリックスと、AI エージェントのブラウズ ツールまたはフェッチ ツールが URL にヒットしたときに Accept ヘッダーを介して Markdown 設定をアドバタイズするかどうか。

記事本文:
どの AI エージェントが Markdown コンテンツ ネゴシエーションをサポートしていますか?
コンテンツへスキップ 同意: text/markdown スタート ガイド レシピ ステータス リファレンス サイトを検索 一致しません。
どの AI エージェントが Accept: テキスト/マークダウン (またはその他) を送信するか
組み込みのブラウズまたはフェッチ時に Markdown 設定をアドバタイズします)
ツールが URL にヒットします。
シェアする
Xで共有する
(新しいタブで開きます)
LinkedIn で共有する
(新しいタブで開きます)
ブルースカイでシェアする
(新しいタブで開きます)
Redditで共有する
(新しいタブで開きます)
マストドンでシェアする
(新しいタブで開きます)
その他…
リンクをコピー
(新しいタブで開きます)
クロードで開く
(新しいタブで開きます)
クロードコードのコピープロンプト
コーデックスのコピープロンプト
Gemini CLI プロンプトのコピー
エージェントの動作を確認しましたか? それとも何か問題を発見しましたか?
証拠（サーバーログ、
curl キャプチャ、またはベンダー ステートメント）。
最初に正規 URL を HTML として取得し (text/markdown Accept 優先なし)、次にドキュメント <head> 内の <link rel="alternate" type="text/markdown" href="…"> の応答を解析し、.md 兄弟に対する 2 番目の要求を作成します。
私たちはマトリックスを自分たちでテストしますが、AI エージェントの動作はさまざまな状況に応じて変化します。
バージョン、プラン、新しく追加されたツール。任意の行を裏付けることができます
— または回帰を捕捉 — に対してエージェントをトリガーすることで、
自分のサーバーを使用し、アクセス ログからリクエストを読み取ります。
1. ログ内の Accept ヘッダーをキャプチャします。
通常、デフォルトのログ形式ではこれが削除されます。一度追加してください:
log_format with_accept '$remote_addr - $remote_user [$time_local] '
'"$request" $status $body_bytes_sent '
'"$http_referer" "$http_user_agent" '
'accept="$http_accept"';
access_log /var/log/nginx/access.log with_accept; LogFormat "%h %l %u %t \"%r\" %>s %b \"%{Referer}i\" \"%{User-Agent}i\" accept=\"%{Accept}i\"" with_accept example.com {
ログ{
出力ファイル /var/log/caddy/access.log
json形式
}
}
# JSON アクセスログには request.head が含まれます

デフォルトで受け入れます。 2. 特定の URL に対して既知の AI エージェントをトリガーする
サイト上の URL を選択します。固有の URL または新しく公開された URL が理想的です。
1 つは、リクエストがバックグラウンド トラフィックによってマスクされないようにするためです。次に、
エージェントがそれを取得または要約します:
ChatGPT — 「https://yoursite.com/article-xyzを要約」（参照ツールを有効にした場合）
クロード — 「https://yoursite.com/article-xyz には何と書いてありますか?」 (web_fetch / web_searchが必要)
Perplexity — URL をクエリに直接貼り付けます
...どのエージェントをテストするかについても同様です。
アクセス ログ内の URL を grep して、次の行に送信します。
フィードバック 、使用したエージェントとともに
送信したプロンプト。 accept= フィールドは次のことを示します。
テキスト/マークダウンがアドバタイズされたかどうか。
Markdown コンテンツ ネゴシエーション: HTML をブラウザーに提供するのと同じ URL から AI エージェントにクリーンな Markdown を提供します。
RFC 7763 · RFC 8288 · RFC 9110

## Original Extract

A running matrix of AI agents and whether they advertise Markdown preference via the Accept header when their browse or fetch tools hit a URL.

Which AI agents support Markdown content negotiation?
Skip to content Accept: text/markdown Get Started Guides Recipes Status Reference Search the site No matches.
Which AI agents send Accept: text/markdown (or otherwise
advertise a Markdown preference) when their built-in browse or fetch
tools hit a URL.
Share
Share on X
(opens in new tab)
Share on LinkedIn
(opens in new tab)
Share on Bluesky
(opens in new tab)
Share on Reddit
(opens in new tab)
Share on Mastodon
(opens in new tab)
Other…
Copy link
(opens in new tab)
Open in Claude
(opens in new tab)
Copy Claude Code prompt
Copy Codex prompt
Copy Gemini CLI prompt
Verified an agent's behavior — or spotted something wrong?
Send feedback with evidence (server log,
curl capture, or vendor statement).
Fetches the canonical URL as HTML first (no text/markdown Accept preference), then parses the response for <link rel="alternate" type="text/markdown" href="…"> in the document <head> and makes a second request for the .md sibling.
We test the matrix ourselves, but AI agent behavior shifts across
versions, plans, and newly-added tools. You can corroborate any row
— or catch a regression — by triggering an agent against
your own server and reading the request from your access logs.
1. Capture the Accept header in your logs
Default log formats usually drop it. Add it once:
log_format with_accept '$remote_addr - $remote_user [$time_local] '
'"$request" $status $body_bytes_sent '
'"$http_referer" "$http_user_agent" '
'accept="$http_accept"';
access_log /var/log/nginx/access.log with_accept; LogFormat "%h %l %u %t \"%r\" %>s %b \"%{Referer}i\" \"%{User-Agent}i\" accept=\"%{Accept}i\"" with_accept example.com {
log {
output file /var/log/caddy/access.log
format json
}
}
# JSON access logs include request.headers.Accept by default. 2. Trigger a known AI agent against a specific URL
Pick a URL on your site — ideally a unique or freshly-published
one so the request isn't masked by background traffic. Then ask the
agent to fetch or summarize it:
ChatGPT — "Summarize https://yoursite.com/article-xyz" (with the browse tool enabled)
Claude — "What does https://yoursite.com/article-xyz say?" (requires web_fetch / web_search)
Perplexity — paste the URL directly into a query
…and so on for whichever agent you're testing.
Grep for the URL in your access log and send the line(s) to
feedback , along with the agent you used
and the prompt you sent. The accept= field tells us
whether text/markdown was advertised.
Markdown content negotiation: serve clean Markdown to AI agents from the same URL that serves HTML to browsers.
RFC 7763 · RFC 8288 · RFC 9110
