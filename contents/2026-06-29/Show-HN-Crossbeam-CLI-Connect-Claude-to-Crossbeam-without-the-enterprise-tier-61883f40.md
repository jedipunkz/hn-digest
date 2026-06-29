---
source: "https://fan-pier-labs.github.io/crossbeam-cli/"
hn_url: "https://news.ycombinator.com/item?id=48722509"
title: "Show HN: Crossbeam-CLI – Connect Claude to Crossbeam without the enterprise tier"
article_title: "Connect Claude to Crossbeam for free — crossbeam-cli"
author: "ryanhughes"
captured_at: "2026-06-29T17:55:33Z"
capture_tool: "hn-digest"
hn_id: 48722509
score: 1
comments: 0
posted_at: "2026-06-29T17:46:38Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Crossbeam-CLI – Connect Claude to Crossbeam without the enterprise tier

- HN: [48722509](https://news.ycombinator.com/item?id=48722509)
- Source: [fan-pier-labs.github.io](https://fan-pier-labs.github.io/crossbeam-cli/)
- Score: 1
- Comments: 0
- Posted: 2026-06-29T17:46:38Z

## Translation

タイトル: HN を表示: Crossbeam-CLI – エンタープライズ層を使用せずにクロードを Crossbeam に接続する
記事のタイトル: クロードを Crossbeam に無料で接続する —crossbeam-cli
説明: ワンクリック拡張機能または CLI を使用して、Crossbeam パートナー、オーバーラップ、アカウント マッピングについてクロードに質問します。スーパーノード プランや有料 API は必要ありません。
HN テキスト: 私は、Claude を Crossbeam に接続するオープンソース ツールである Crossbeam-cli を構築しました。これにより、パートナー、重複、アカウント マッピングについて平易な英語で質問できるようになります。 Crossbeam は、最上位のスーパーノード層の背後でプログラムによるアクセスをゲートします。公式の Claude コネクタでさえ、年間約 4,800 ドルから始まります。 Crossbeam-cli は、Web アプリが使用するのと同じ内部 API と通信するため、無料利用枠を含むあらゆるプランで動作します。自分の資格情報を使用してログインすると、リクエストはコンピュータから Crossbeam に直接送信されます。これは 2 つの方法で提供されます。1 つは非技術パートナー チーム向けのワンクリック クロード デスクトップ拡張機能 (MCPB)、もう 1 つはスクリプトとエージェント用の CLI (npm i -gcrossbeam-cli) です。これまでのところ、読み取り専用ツールのみを実装していますが、将来的には書き込みサポートを追加する可能性があります。これは非公式であり、Crossbeam と提携または承認されていません。これは文書化されていないエンドポイントに依存しているため、いつでも変更または破壊される可能性があります。皆さんのご意見をぜひお聞かせください。 https://fan-pier-labs.github.io/crossbeam-cli/

記事本文:
クロスビーム-cli
なぜ
仕組み
2つの使い方
GitHub
ダウンロード
オープンソース
クロードを接続する
クロスビームは無料です。
パートナー、重複、アカウント マッピングについてクロードに質問してください。スーパーノード プランや有料 API はありません。
ワンクリック拡張機能、またはターミナル用の CLI。
クロード拡張機能をダウンロードする
GitHub で見る
ターミナルの方がいいですか？ npm install -gcrossbeam-cli — CLI を参照してください。
Crossbeam の API は最上位層の背後でロックされています。
アカウント マッピング データの取得
プログラム的には、たとえ公式の Claude コネクタであっても、スーパーノード プランが必要です。
これは、Web アプリが既に呼び出しているのと同じエンドポイントを使用するため、無料を含むどのプランでも機能します。
公式 API と Claude コネクタ
最上位の有料プランにゲートされます (営業担当者にお問い合わせください)
コネクターは年間 $4,800 から — API は含まれません
無料プラン: ブラウザーで一度に最大 50 レコード
パートナー、母集団、重複、アカウントのマッピング
既存のメールアドレスとパスワードのみ
クロード、スクリプト、またはターミナルから使用する
Crossbeam, Inc. と提携または承認されていません。ご自身の責任でご使用ください。
資格情報がマシンから離れることはありません。
独自のログインを使用して、HTTPS 経由で Crossbeam と直接通信します。それがすべてです。
自分の Crossbeam 電子メールとパスワード — Web アプリと同様に、 auth.crossbeam.com に対して認証されます。
セッションは自分のディスクの ~/.crossbeam にキャッシュされます。他の場所には送信されません。
呼び出しは api.crossbeam.com に直接送信されます。わかりやすい英語でクロードに質問するか、ターミナルに JSON をパイプしてください。
ほとんどの人はクロード拡張機能を望んでいます。端末タイプは JSON と同じデータを取得します。
ワンクリックの .mcpb バンドル。インストールし、ログイン情報を一度入力して、「どのパートナーが私のオープンな機会と重なっていますか?」と尋ねるだけです。
技術的な専門知識は不要 - 営業チームやパートナー チームに最適
コードなし - ドラッグアンドドロップでインストール
41 個の読み取りツール、Claude に公開

rMCP
OS キーチェーンに保存されているパスワード
Crossbeam.mcpb をダウンロード
開発者向け
コマンドラインツール
ノード パッケージと CLI。すべてのコマンドは JSON を出力します。これは、スクリプト、cron ジョブ、コーディング エージェントに最適です。
技術的な専門知識が必要 (ターミナルとノード)
Node/TypeScript クライアントもインポート可能
クロード コード、カーソル、その他のエージェント向けに構築
GitHub のドキュメントを読む
お願いできること
41 の読み取りツール、1 回のログイン。
Claude (または CLI) が取得できるもののサンプルに加えて、その他の脱出ハッチを取得します。
...その他 29 個 — 完全なリストは、manifest.json にあります。
バックエンド、プロキシ、テレメトリ、分析はありません。データは Crossbeam から直接送信されます。
すべての行は GitHub にあります。読んで、監査して、フォークしてください。ブラックボックスはありません。
ツールは読み取り専用です。 Crossbeam データは何も変更、削除、共有されません。
拡張機能を 1 分以内にインストールするか、CLI を取得します。無料でオープンソース。
クロード拡張機能をダウンロードする
GitHub でスターを付ける
クロスビーム-cli
GitHub
ダウンロード拡張機能
npm
ドキュメント
免責事項:crossbeam-cli は独立したオープンソース プロジェクトです。とは提携していませんが、
Crossbeam, Inc. によって承認、維持、後援、または承認されています。Crossbeam の社内と対話しています。
api.crossbeam.com エンドポイント — サインイン時に Crossbeam Web アプリが使用するものと同じ —
自分自身の認証情報を使用します。 「Crossbeam」はそれぞれの所有者の商標です。ご自身の責任でご使用ください。

## Original Extract

Ask Claude about your Crossbeam partners, overlaps, and account mapping — with a one-click extension or a CLI. No Supernode plan or paid API required.

I built crossbeam-cli, an open-source tool that connects Claude to Crossbeam so you can ask questions about your partners, overlaps, and account mapping in plain English. Crossbeam gates programmatic access behind its top Supernode tier. Even the official Claude connector starts around $4,800/year. crossbeam-cli works on any plan, free tier included, because it talks to the same internal API the web app uses. You log in with your own credentials and requests go straight from your computer to Crossbeam. It ships two ways: a one-click Claude desktop extension (MCPB) for non-technical partner teams, and a CLI (npm i -g crossbeam-cli) for scripting and agents. So far I've only implemented read-only tools, but I could add write support in the future. It's unofficial and not affiliated with or endorsed by Crossbeam. It relies on undocumented endpoints, so they could change or break it at any time. Would love to hear what you all think. https://fan-pier-labs.github.io/crossbeam-cli/

crossbeam-cli
Why
How it works
Two ways to use
GitHub
Download
Open source
Connect Claude to
Crossbeam for free.
Ask Claude about your partners, overlaps, and account mapping — no Supernode plan, no paid API.
One-click extension, or a CLI for the terminal.
Download the Claude extension
View on GitHub
Prefer the terminal? npm install -g crossbeam-cli — see the CLI .
Crossbeam's API is locked behind its top tier.
Getting your account-mapping data out
programmatically — even the official Claude connector — requires the Supernode plan.
This uses the same endpoints the web app already calls, so any plan, including free , works.
Official API & Claude connector
Gated to the top paid plan (contact sales)
Connector starts ~$4,800/yr — API not included
Free plan: ~50 records at a time, in the browser
Partners, populations, overlaps & account mapping
Just your existing email + password
Use from Claude, a script, or the terminal
Not affiliated with or endorsed by Crossbeam, Inc. Use at your own risk.
Your credentials never leave your machine.
It talks directly to Crossbeam over HTTPS with your own login. That's the whole thing.
Your own Crossbeam email and password — authenticated against auth.crossbeam.com , just like the web app.
The session is cached at ~/.crossbeam on your own disk — never sent anywhere else.
Calls go straight to api.crossbeam.com . Ask Claude in plain English, or pipe JSON in the terminal.
Most people want the Claude extension. Terminal types get the same data as JSON.
A one-click .mcpb bundle. Install it, enter your login once, then just ask — "which partners overlap with my open opportunities?"
No technical expertise needed — great for sales & partner teams
No code — drag-and-drop install
41 read tools, exposed to Claude over MCP
Password stored in your OS keychain
Download crossbeam.mcpb
For developers
Command-line tool
A Node package and CLI. Every command emits JSON — perfect for scripts, cron jobs, and coding agents.
Technical expertise required (terminal & Node)
Importable Node/TypeScript client too
Built for Claude Code, Cursor & other agents
Read the docs on GitHub
What you can ask for
41 read tools, one login.
A sample of what Claude (or the CLI) can pull — plus a get escape hatch for anything else.
…and 29 more — full list in manifest.json .
No backend, proxy, telemetry, or analytics. Your data goes from Crossbeam straight to you.
Every line is on GitHub. Read it, audit it, fork it — no black boxes.
The tools only read . Nothing modifies, deletes, or shares your Crossbeam data.
Install the extension in under a minute, or grab the CLI. Free and open source.
Download the Claude extension
Star it on GitHub
crossbeam-cli
GitHub
Download extension
npm
Docs
Disclaimer: crossbeam-cli is an independent, open-source project. It is not affiliated with,
authorized, maintained, sponsored, or endorsed by Crossbeam, Inc. It talks to Crossbeam's internal
api.crossbeam.com endpoints — the same ones the Crossbeam web app uses when you are signed in —
using your own credentials. "Crossbeam" is a trademark of its respective owner. Use at your own risk.
