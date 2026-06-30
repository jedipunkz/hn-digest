---
source: "https://traceaio.org"
hn_url: "https://news.ycombinator.com/item?id=48732376"
title: "Show HN: TraceAIO – open-source LLM visibility tracker"
article_title: "TraceAIO | Track Your Brand Across LLMs"
author: "owenthejumper"
captured_at: "2026-06-30T13:34:04Z"
capture_tool: "hn-digest"
hn_id: 48732376
score: 2
comments: 0
posted_at: "2026-06-30T13:18:26Z"
tags:
  - hacker-news
  - translated
---

# Show HN: TraceAIO – open-source LLM visibility tracker

- HN: [48732376](https://news.ycombinator.com/item?id=48732376)
- Source: [traceaio.org](https://traceaio.org)
- Score: 2
- Comments: 0
- Posted: 2026-06-30T13:18:26Z

## Translation

タイトル: HN を表示: TraceAIO – オープンソース LLM 可視性トラッカー
Article title: TraceAIO | Track Your Brand Across LLMs
Description: Open-source LLM brand tracker. ChatGPT、Perplexity、Gemini があなたのブランドについて言及している場所と言及している場所を確認します。
HN テキスト: 私は、あなたに代わって LLM にプロンプトを表示し、ChatGPT、Perplexity、Gemini があなたのブランドについて言及しているかどうか、そして代わりにどの競合他社やソースが表示されているかを通知するオープンソース ツールである TraceAIO を構築しました。そうですね、このカテゴリは初期の SEO と同じように、少し厄介な匂いがします。そして時間が経つと、それは再び単なるSEOになり、優れたコンテンツになると思います。 The tool just helps you monitor over time. API ではなく実際のブラウザ セッションを通じてブラウザ製品にクエリを実行し、MCP サーバーを使用して Docker 上で実行されるため、LLM を通じて独自のデータをクエリできます。 No business model, Apache 2.0, self hosted.ブラウザをローカルで実行する代わりに住宅用プロキシを使用すると、紹介経由で数セントを受け取りますが、トークンのコストをカバーするには十分ではありません。デモ (サインアップなし): https://demo.traceaio.org — admin@example.com / adminadmin コード: https://traceaio.org 長く続いたプロジェクトですが、ようやくお見せできることを嬉しく思います。フィードバックは歓迎です。

記事本文:
ブランドを追跡する
すべての LLM にわたって
ChatGPT、Perplexity、Gemini があなたのブランドについて言及している場所を確認してください。隙間を見つけてください。実際のモデルに対して実際のプロンプトを実行します。
ブラウザ内にのみ存在するサーフェスのブラウザベース。そうでないサーフェスについては API ベース。
チャットGPT
chatgpt.com
困惑
perplexity.ai
Google ジェミニ
ジェミニ.google.com
Google AIモード
google.com (AI の概要)
近日公開予定
クロード・アイ
クロード・アイ
API
組み込みの Web 検索を備えたダイレクト プロバイダー API。ブラウザランタイムはありません。
OpenAI API
gpt-4o · web_search
人類API
クロード · web_search
// ライブデモ
ライブデモを試してみる
サンプル データがロードされた実際の実行中の TraceAIO インスタンス — ダッシュボード、プロンプト、競合他社、およびソースを探索します。サインアップは必要ありません。
admin@example.com / adminadmin でログインします
// 仕組み
本物のモデル。本当の反応。
AI は、業界で重要なトピック、つまり潜在的な顧客が実際に尋ねる種類の質問について、ブランド中立的なプロンプトを作成します。独自のカスタム プロンプトも追加します。プロンプトは保存され、実行後も再利用できます。
各モデルにユーザーのやり方を尋ねる
ブラウザベースのモデル (ChatGPT、Perplexity、Gemini、Google AI モード) は、実際のブラウザ セッションを通じて、ユーザーに表示されているものとまったく同じようにクエリされます。 API ベースのモデル (OpenAI と Anthropic) は、公式の回答エンジンが使用するのと同じパスである、組み込みの Web 検索を使用してプロバイダーの API を呼び出します。いずれの場合でも、必要最低限​​の API の近似値ではなく、実際の答えが得られます。
Apify 経由の住宅用プロキシを使用して、毎分最大 15 のプロンプトを並行して実行します。リクエストごとに IP が異なり、ボット対策の問題もインフラストラクチャへのリスクもありません。
無料ですがリスクがあります。マシン上で実際のブラウザを一度に 1 つのプロンプトで実行します。 IP は各 LLM プロバイダーに直接公開されます。アンチボット システムは、同じアドレスからの繰り返しのリクエストを捕捉し、これらのアドレスへの通常のアクセスをブロックする可能性があります。

サービス。本番環境には推奨されません。
すべての反応が分析されます。あなたのブランドが言及されましたか?どの競争相手が現れましたか？どのような情報源が引用されましたか?結果は実行ごとに保存されるため、時間の経過に伴う変化を追跡し、モデルを比較し、欠落している箇所を正確に見つけることができます。
MCP 経由でクロードに接続し、会話形式でブランド データをクエリします。 「当社ではなく競合他社について言及しているプロンプトはどれですか?」 「Perplexity が引用しているのに、Gemini が引用していない情報源は何ですか?」 「前回の製品発売後、メンション率はどう変化しましたか?」 SQLは必要ありません。
必要なものすべて
LLM の可視性を監視する
AI は、業界に関連するトピック全体にわたってブランド中立的なプロンプトを作成します。自分で追加してください。複数の実行にわたって再利用します。
実際のブラウザ セッションを介して、ChatGPT、Perplexity、Google Gemini、および Google AI モードに対して実行されます。 APIではありません。まさにユーザーが見ているものです。
応答から競合他社を自動検出します。重複をマージし、ノイズをブロックし、直接比較します。
LLM がどのドメインを引用しているかを追跡します。ブランド、競合他社、中立的な情報源を参照してください。引用のギャップを見つけます。
一度スケジュールを立ててください。分析を時間ごと、毎日、毎週、または毎月実行します。 Stuck は 24 時間後に自動的に期限切れになります。
Webhook、n8n、REST API、MCP サーバー。必要な場所に結果をプッシュし、必要に応じてデータをプルします。
結果を必要な場所にプッシュします。必要に応じてデータをプルします。
分析が完了したら通知を受け取ります。オプションのベアラー トークン認証を使用して URL を構成すると、実行結果が POST として届きます。それらを Slack、n8n、Zapier、または独自のバックエンドにパイプします。
n8n のコミュニティ ノード。分析完了時にワークフローをトリガーし、結果をスプレッドシートに同期し、レポートを電子メールで送信し、400 以上の他のアプリに接続します。
Swagger ドキュメントを含む完全な API。実行、応答、競合他社、ソース、およびメトリクスを取得します。ダッシュボード、レポートを構築したり、独自の分析パイプラインにデータをフィードしたりできます。
Claude Code および Cla 用の 16 個の組み込みツール

うでデスクトップ。会話形式でブランド データをクエリします。モデルを比較し、引用のギャップを見つけ、実行全体の傾向を追跡します。
導入には 2 つの方法があります。ワークフローに合ったものを選択してください。
画像をプルして開始します。ブラウザ自動化コンテナを含め、すべてがローカルで実行されます。
管理者アカウントを作成し、API キーを追加し、セットアップ ウィザードでブランドを構成します。並列実行を高速化したい場合は、[設定] → [認証情報] で Apify に切り替えます。
クロードに設定してもらいましょう。これをクロード コードに貼り付けると、残りの処理が行われます。
セットアップ後、MCP 経由でクロードをデータに接続します。
クロード mcp add --transport http brand-tracker http://localhost:3000/mcp

## Original Extract

Open-source LLM brand tracker. See where ChatGPT, Perplexity, and Gemini mention your brand, and where they don

I built TraceAIO, an open-source tool that prompts LLMs on your behalf and tells you whether ChatGPT, Perplexity, and Gemini mention your brand — and which competitors and sources show up instead. Yeah, this category smells a bit like a grift, same as early SEO. And I think over time it will become just SEO again, and become about good content. The tool just helps you monitor over time. It queries the browser products through real browser sessions, not APIs, runs on Docker, with an MCP server so you can query your own data through an LLM. No business model, Apache 2.0, self hosted. If you use a residential proxy instead of running the browser locally I get a few cents via referral, not even enough to cover my token costs. Demo (no signup): https://demo.traceaio.org — admin@example.com / adminadmin Code: https://traceaio.org Long on-and-off project, finally happy to show it. Feedback welcome.

Track your brand
across every LLM
See where ChatGPT, Perplexity, and Gemini mention your brand. Find the gaps. Run real prompts against real models.
Browser-based for surfaces that only exist in a browser. API-based for surfaces that don’t.
ChatGPT
chatgpt.com
Perplexity
perplexity.ai
Google Gemini
gemini.google.com
Google AI Mode
google.com (AI Overviews)
Coming soon
Claude AI
claude.ai
API
Direct provider API with built-in web search. No browser runtime.
OpenAI API
gpt-4o · web_search
Anthropic API
Claude · web_search
// live demo
Try the live demo
A real, running TraceAIO instance loaded with sample data — explore the dashboard, prompts, competitors, and sources. No signup required.
Log in with admin@example.com / adminadmin
// how it works
Real models. Real responses.
AI creates brand-neutral prompts on topics that matter in your industry, the kind of questions your potential customers actually ask. Add your own custom prompts too. Prompts are saved and reusable across runs.
Ask each model the way users do
Browser-based models — ChatGPT, Perplexity, Gemini, Google AI Mode — are queried through real browser sessions, exactly what your users see. API-based models — OpenAI and Anthropic — call the provider’s API with built-in web search, the same path their official answer engines use. Either way, you get the production answer, not a stripped-down API approximation.
Uses residential proxies via Apify to run ~15 prompts/min in parallel. Different IP per request, no anti-bot issues, no risk to your infrastructure.
Free but risky. Runs a real browser on your machine, one prompt at a time. Your IP is exposed directly to each LLM provider. Anti-bot systems catch repeated requests from the same address and can block your normal access to these services. Not recommended for production.
Every response is analyzed: was your brand mentioned? Which competitors showed up? What sources were cited? Results are stored per-run so you can track changes over time, compare models, and find exactly where you're missing.
Connect Claude via MCP and query your brand data conversationally. "Which prompts mention our competitor but not us?" "What sources does Perplexity cite that Gemini doesn't?" "How did our mention rate change after the last product launch?" No SQL required.
Everything you need to
monitor LLM visibility
AI creates brand-neutral prompts across topics relevant to your industry. Add your own. Reuse across runs.
Runs against ChatGPT, Perplexity, Google Gemini, and Google AI Mode via real browser sessions. Not APIs. Exactly what your users see.
Auto-detects competitors from responses. Merge duplicates, block noise, compare head-to-head.
Track which domains LLMs cite. See brand, competitor, and neutral sources. Find citation gaps.
Schedule once. Run analysis hourly, daily, weekly, or monthly. Stuck runs auto-expire after 24 hours.
Webhooks, n8n, REST API, and MCP server. Push results where you need them and pull data however you want.
Push results where you need them. Pull data however you want.
Get notified when analysis completes. Configure a URL with optional Bearer token auth, and run results arrive as a POST. Pipe them into Slack, n8n, Zapier, or your own backend.
Community node for n8n. Trigger workflows on analysis completion, sync results to spreadsheets, send reports via email, and connect to 400+ other apps.
Full API with Swagger docs. Fetch runs, responses, competitors, sources, and metrics. Build dashboards, reports, or feed data into your own analytics pipeline.
16 built-in tools for Claude Code and Claude Desktop. Query your brand data conversationally. Compare models, find citation gaps, track trends across runs.
Two ways to deploy. Pick whichever fits your workflow.
Pull the image and start. Everything runs locally, including the browser automation container.
Create an admin account, add your API key, and configure your brand in the setup wizard. Switch to Apify in Settings → Credentials if you want faster parallel runs.
Let Claude set it up for you. Paste this into Claude Code and it handles the rest.
After setup, connect Claude to your data via MCP:
claude mcp add --transport http brand-tracker http://localhost:3000/mcp
