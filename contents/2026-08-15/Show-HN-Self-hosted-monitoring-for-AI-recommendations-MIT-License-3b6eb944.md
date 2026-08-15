---
source: "https://lettertrace.com"
hn_url: "https://news.ycombinator.com/item?id=49312505"
title: "Show HN: Self-hosted monitoring for AI recommendations (MIT License)"
article_title: "Lettertrace: Monitor how AI talks about your brand"
author: "mathewpregasen"
captured_at: "2026-08-15T18:14:46Z"
capture_tool: "hn-digest"
hn_id: 49312505
score: 2
comments: 0
posted_at: "2026-08-15T17:35:13Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Self-hosted monitoring for AI recommendations (MIT License)

- HN: [49312505](https://news.ycombinator.com/item?id=49312505)
- Source: [lettertrace.com](https://lettertrace.com)
- Score: 2
- Comments: 0
- Posted: 2026-08-15T17:35:13Z

## Translation

タイトル: HN を表示: AI 推奨事項のセルフホスト型モニタリング (MIT ライセンス)
記事のタイトル: Lettertrace: AI がブランドについてどのように語るかを監視する
説明: AI アシスタントの回答にブランドがどのように表示されるかを監視する、オープンソースの持ち込みキー監視。トピックを追跡し、プロンプトのバリエーションを生成し、傾向を監視し、競合他社のベンチマークを行います。

記事本文:
Lettertrace: AI がブランドについてどのように話すかを監視します 仕組み 機能 オープンソース サインイン オープンソースの初期化 · Bring-Your-Own-Key AI の可視性を無料で追跡します。
Lettertrace は、Claude、ChatGPT、Gemini があなたの会社について言及する頻度を測定します。ただし、落とし穴があります。それは、エンドツーエンドで無料、開発者優先、そしてオープンソースであるということです。
ChatGPT、Claude & Gemini で動作する CLI を数分でインストールし、セルフホストをインストールします
$lettertrace run --project acme
› 24 個のプロンプトにわたって claude-opus-4-8 をクエリしています…
✓ 15 / 24 件の回答で acme について言及
発言権の割合 ████████░░░░░░░░ 41%
トピックからトレンドラインまで自動的に。
「スタートアップ向けの最適な CRM」、「代理店向けの電子メール ツール」など、バイヤーが尋ねる件名を追加します。
Lettertrace は、ユーザーが AI アシスタントに寄せた実際の質問を書き出すため、質問が実際にどのように回答されるかを監視できます。
すべての実行はデータポイントです。注目度、知名度、センチメントの週ごとの変化を観察してください。
ライバルブランドを取り込んで、自分の意見のシェア、モデルが誰を、どのくらいの頻度で推奨しているかを確認します。
データが豊富な可視性レポート。
Lettertrace は、AI の生の回答を AEO/GEO チームが追跡する指標に変換します。つまり、トピックごと、モデルごと、時間の経過とともに、可視性、声のシェア、知名度、感情などを追跡します。
「急成長を遂げているスタートアップ企業にとっては、いくつかのツールが際立っています。Acme はブランドの声を失わずに自動化を望むチームにとって強力な選択肢ですが、Notion と Linear はドキュメントや問題の追跡に人気があります…」
AI の診断に必要なものすべてについて説明します。
独自の Anthropic、OpenAI、Google キーを使用します。これらは保存時に暗号化され、インフラストラクチャから離れることはありません。
Claude、ChatGPT、Gemini、Google AI の概要を並べて監視します。必要に応じて回答エンジンを追加します。
1 つのトピックを手動でプロンプトを表示せずに、自動的に数十の自然なプロンプトに変換します。

儀式。
自分が登場するかどうかだけでなく、その回答があなたのことをよく語っているか、あなたを推薦するかどうかを知りましょう。
追跡している各競合他社と比較して、回答を獲得する頻度を正確に確認できます。
自動操縦で毎日または毎週実行し、それに基づいて行動できる傾向線を構築します。
実行、検査、拡張するのはあなたです。
MIT ライセンスを取得し、フォークして自己ホストし、自分のものにします。
鍵はご自身でご用意ください。使用マークアップも仲介者もいません。
データは独自の Supabase に保存され、ベンダー ロックインはありません。
$ git clone レタートレース && cd レタートレース
AI があなたについて何を言っているか調べてみましょう。
ブランドを設定し、キーを追加して、最初のモニターを数分で実行します。
AI アシスタントの回答にブランドがどのように表示されるかをオープンソースでモニタリングします。

## Original Extract

Open-source, bring-your-own-key monitoring for how your brand shows up in AI assistant answers. Track topics, generate prompt variations, watch trends, and benchmark competitors.

Lettertrace: Monitor how AI talks about your brand How it works Features Open source Sign in Initialize open-source · bring-your-own-key Track your AI visibility, for free .
Lettertrace measures how often Claude, ChatGPT, and Gemini mention your company. But there's a catch: it's free end-to-end, developer-first, and open source.
Install the CLI works with ChatGPT, Claude & Gemini · self-host in minutes
$ lettertrace run --project acme
› querying claude-opus-4-8 across 24 prompts…
✓ acme mentioned in 15 / 24 answers
share of voice ████████░░░░░░░░ 41%
From a topic to a trend line, automatically.
Add the subjects your buyers ask about, “best CRM for startups”, “email tools for agencies”.
Lettertrace writes the real questions people put to AI assistants, so you monitor how they’re actually answered.
Every run is a datapoint. Watch visibility, prominence, and sentiment move week over week.
Ingest rival brands and see your share of voice, who the models recommend, and how often.
A data-rich visibility report.
Lettertrace turns raw AI answers into the metrics AEO/GEO teams track: visibility, share of voice, prominence, and sentiment, per topic, per model, over time.
“For fast-growing startups, a few tools stand out. Acme is a strong pick for teams that want automation without losing their brand voice, while Notion and Linear are popular for docs and issue tracking…”
Everything you need to diagnose AI mentions.
Use your own Anthropic, OpenAI & Google keys. They’re encrypted at rest and never leave your infrastructure.
Monitor Claude, ChatGPT, Gemini, and Google AI Overviews side by side. Add more answer engines as they matter.
Turn one topic into dozens of natural prompts automatically, no manual prompt-writing.
Know not just if you appear, but whether the answer speaks well of you, and recommends you.
See exactly how often you win the answer versus each competitor you track.
Run daily or weekly on autopilot and build a trend line you can act on.
Yours to run, inspect, and extend.
MIT licensed, fork it, self-host it, make it yours.
Bring your own keys. No usage markup, no middleman.
Your data lives in your own Supabase, no vendor lock-in.
$ git clone lettertrace && cd lettertrace
Find out what AI says about you.
Set up your brand, add a key, and run your first monitor in minutes.
Open-source monitoring for how your brand shows up in AI assistant answers.
