---
source: "https://llmstatus.ai"
hn_url: "https://news.ycombinator.com/item?id=49374169"
title: "Show HN: Find every AI model your code calls and warn before it's retired"
article_title: "LLM Status — AI model deprecation tracker & CLI checker"
image: "https://llmstatus.ai/opengraph-image?72cd1c5ab8ea556a"
author: "taylorgt"
captured_at: "2026-08-20T13:39:09Z"
capture_tool: "hn-digest"
hn_id: 49374169
score: 4
comments: 1
posted_at: "2026-08-20T13:12:55Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Find every AI model your code calls and warn before it's retired

- HN: [49374169](https://news.ycombinator.com/item?id=49374169)
- Source: [llmstatus.ai](https://llmstatus.ai)
- Score: 4
- Comments: 1
- Posted: 2026-08-20T13:12:55Z

## Translation

タイトル: HN を表示: コードで呼び出されるすべての AI モデルを検索し、廃止される前に警告します
記事のタイトル: LLM ステータス — AI モデルの非推奨トラッカーと CLI チェッカー
説明: アプリ全体ですべての AI モデルが使用されている場所を追跡し、非推奨または廃止される前に警告を受け取ります。

記事本文:
LLM ステータス — AI モデル非推奨トラッカーおよび CLI チェッカー
LLM ステータス
新しい
レーダー
仕組み
GitHub
価格設定
ドキュメント
サインイン
Amazon Nova Premier は 25 日後に終了します
AI モデルの有効期限は切れます。準備をしてください。
任意のリポジトリで実行します。コードで呼び出すモデルを検索し、非推奨または廃止されたモデルに日付と代替モデルを付けてフラグを立てます。
コピーしました！
サインアップや API キーは必要ありません。コードがマシンから離れることはありません。
仕組み↓
無料アカウントを作成する
mm ステータス — ~/code/api-service
＄mmステータス ●387◐41▲20✕79
LLM ステータス — レジストリはライブです · 15 プロバイダー · 527 モデル
スキャンされた ~/code/api-service · 6 つのモデルが追跡されました
◐ 非推奨の openai/gpt-5-2025-08-07 は 4 か月以内に 2026-12-11 に廃止されます
▲ 廃止予定の amazon/amazon.nova-premier-v1:0 は 2026 年 9 月 14 日 25 日に廃止
● ok anthropic/claude-opus-4-7 は 2027 年 4 月 16 日に 8 か月で引退します
◐ 非推奨の openai/o3-pro-2025-06-10 は 4 か月以内に 2026-12-11 に廃止されます
✕ 廃止された Microsoft/gpt-5.3-chat は 2026-08-10 10 日前に廃止されました
▲ openai/gpt-3.5-turbo-instruct の引退 2026-09-28 の 39 日で引退
⚠ 527 のレジストリ モデルのうち 140 が非推奨、廃止、または廃止されました。リポジトリのリストに対して mm を実行します。
↑↓ ナビスペース選択 q 終了
これらの退職日はレジストリからライブで取得されます。リポジトリのパスとスキャン数はサンプルです。
セットアップ
それをインストールして、 mm を実行します。
mm コマンドをインストールします。自己完結型のバイナリであり、ノードは必要ありません (60 ～ 120 MB、プラットフォームを自動検出します)。バックグラウンドでの自動更新、署名検証。 MM_NO_AUTO_UPDATE=1 はこれを無効にします。
コピーしました！または brew install randomartifact/tap/modelstatus-cli · npm i -g @modelstatus/cli · npx @modelstatus/cli status で一度試してください
2 ダッシュボードを開く
完全なターミナル ダッシュボードに対して mm を実行します。起動時に現在のリポジトリがスキャンされ、呼び出されるすべてのモデルを矢印で移動し、レジストリを参照して、最新のライブ情報を確認します。
必要

スクリプト化可能ですか? mm ステータスには、リポジトリで呼び出されるすべてのモデルがリストされ、リタイアが最も早いものから順にリストされます。 mm ci は、何かがクロック上にある場合、ゼロ以外で終了します。これは、コミット前のフックと CI の場合です。アカウントがありません。
コピーしました！
どちらも cdn.llmstatus.ai から署名付きレジストリを 1 回取得し、オフラインで実行します。アカウントがありません。匿名の使用量カウントのみ - MM_NO_ANALYTICS=1 でオフになります。
ダッシュボード全体を表示するには mm を実行します。
「mm」と入力して開きます。リポジトリが呼び出すすべてのモデルを矢印で移動し、f を押して差分プレビューで 1 つを修正し、進行中に●◐▲✕ カウントが更新されるのを確認します。 [Here] タブと [What's New] タブは、署名されたレジストリから直接読み込まれ、アカウントは必要なく、オフラインでも機能します。
mm 交換品のスワップを修正します。
瀕死のモデルを見つけることが仕事の半分です。 1 つのコマンドで、それらをレジストリの代わりの場所に書き換えます。
ID テキストのみが変更されます。引用符、書式設定、行の残りの部分: バイト同一。
境界安全: gpt-4 は gpt-4o または gpt-4-tuned 内で書き換えることはできません。
瀕死の代替品を最初のライブモデルにチェーンします。
--dry-run プレビュー、--json ツール用。 TUI で f を押します。
GitHub アプリはプル リクエストに対しても同じことを行います。チェックが失敗すると [Open fix PR] ボタンが成長し、ワンクリックでこれが開きます。本文の差分、コミットが検証され、同じスキャンによって緑色に再チェックされます。
レジストリからライブ — models.dev、OpenRouter、およびプロバイダー API から 6 時間ごとに更新されます。
アマゾン ノヴァ プレミア
アマゾン
2026-09-14
25日以内に
GPT-3.5ターボ命令
OpenAI
2026-09-28
39日で
クロード・ソネット 4.5
人間的
2026-09-29
40日以内に
アマゾン ノヴァ キャンバス
アマゾン
2026-09-30
41日で
GPT-4o
OpenAI
2026-10-01
42日で
ナノバナナ
Google
2026-10-02
43日で
クロード俳句 4.5
人間的
2026-10-15
56日で
ジェミニ 2.5 フラッシュ
Google
2026-10-16
57日で
GG
ジェミニ 3.7 フラッシュ 7 日前
Google · ジェミニフラッシュ · 1049k ctx
XA
Grok 4.6 8日前
xAI · grok · 500k ctx
午前
グロク

4.6 8日前
アマゾン · grok · 500k ctx
XA
Grok Imagine Image 2.0 13日前
xAI · grok · 8k ctx
私
Muse Spark 1.2 コントリビューター 15 日前
メタ・ミューズ・1049k ctx
AI プロバイダーがモデル ライフサイクル フィードを提供することはありません。私たちはそうします—
携帯電話でレーダーを開くか、購読してください:
RSS・
JSON ·
退職者のみ
プルリクエストごとに確認してください。
すべての PR にモデルのライフサイクル チェックを追加します。 2 つのオプションがあります。違いは、ソース コードがランナーから離れるかどうかです。
独自の CI で実行します。ソースがランナーから離れることはなく、モデル ID のみが出力されます。インライン PR 注釈と合否ゲート。
ワンクリックで、ワークフロー YAML やシークレットを管理する必要はありません。当社はサーバー上の各 PR をスキャンし (コードが当社に送信されます。取引違反の場合は CI オプションを使用します)、チェックランとレビュー コメントを投稿します。チェックのボタンを押すと修正 PR が開き、置換を交換します。
CLI は無料です。クラウドは年間5ドル。
プランにはクラウド インベントリとアラートが含まれます。 CLI 自体は常に無料です。
ターミナル内のすべてが無料、アカウント不要
1 プロジェクト、最大 15 件の使用状況を追跡
電子メール + アプリ内退職アラート
Slack、Discord、SMS、Webhook アラート
CLI クラウド同期 + CI ドリフト ダッシュボード
カスタムリードタイム (90 / 30 / 7 / 1 日)
サインインすると、退職前にメールが届きます。
無料アカウントには 1 つのプロジェクトのインベントリ (最大 15 回の使用量) が保存され、モデルが廃止される前に電子メールが送信されます。 Pro では上限が撤廃され、Slack、Discord、SMS、Webhook、CI ドリフト、API が追加されます。

## Original Extract

Track where every AI model is used across your apps — and get warned before one is deprecated or retired.

LLM Status — AI model deprecation tracker & CLI checker
LLM Status
New
Radar
How it works
GitHub
Pricing
Docs
Sign in
Amazon Nova Premier retires in 25 days
AI models expire . Be Prepared.
Run it in any repo. It finds the models your code calls and flags the ones that are deprecated or retiring, with the date and the replacement.
Copied!
No signup, no API key. Your code never leaves your machine.
How it works ↓
Create a free account
mm status — ~/code/api-service
$ mm status ●387 ◐41 ▲20 ✕79
LLM Status — registry live · 15 providers · 527 models
Scanned ~/code/api-service · 6 models tracked
◐ deprecating openai/gpt-5-2025-08-07 retires 2026-12-11 in 4mo
▲ retiring amazon/amazon.nova-premier-v1:0 retires 2026-09-14 in 25d
● ok anthropic/claude-opus-4-7 retires 2027-04-16 in 8mo
◐ deprecating openai/o3-pro-2025-06-10 retires 2026-12-11 in 4mo
✕ retired microsoft/gpt-5.3-chat retires 2026-08-10 10d ago
▲ retiring openai/gpt-3.5-turbo-instruct retires 2026-09-28 in 39d
⚠ 140 of 527 registry models are deprecated, retiring, or gone. Run mm for your repo's list.
↑↓ nav space select q quit
These retirement dates are live from the registry. Repo path and scan counts are a sample.
Setup
Install it, then run mm .
Installs the mm command — a self-contained binary, no Node needed (60–120 MB, auto-detects your platform). Auto-updates in the background, signature-verified; MM_NO_AUTO_UPDATE=1 disables it.
Copied! or brew install randomartifact/tap/modelstatus-cli · npm i -g @modelstatus/cli · try once with npx @modelstatus/cli status
2 Open the dashboard
Run mm for the full terminal dashboard: it scans the current repo on launch, then you arrow through every model it calls, browse the registry, and see what's new — live.
Need it scriptable? mm status lists every model the repo calls, soonest retirement first. mm ci exits non-zero when something's on a clock — that's the one for pre-commit hooks and CI. No account.
Copied!
Both pull the signed registry from cdn.llmstatus.ai once, then run offline. No account. Anonymous usage counts only — MM_NO_ANALYTICS=1 turns them off.
Run mm for the full dashboard.
Type mm to open it. Arrow through every model the repo calls, press f to fix one with a diff preview, and watch the ●◐▲✕ counts update as you go. The Here and What's New tabs read straight off the signed registry, no account, works offline.
mm fix swaps in the replacements.
Finding dying models is half the job. One command rewrites them to the registry's replacement, in place.
Only the id text changes. Quotes, formatting, the rest of the line: byte-identical.
Boundary-safe: gpt-4 can't rewrite inside gpt-4o or gpt-4-tuned .
Chains through dying replacements to the first live model.
--dry-run previews, --json for tooling. In the TUI, press f .
The GitHub App does the same on pull requests: a failing check grows an Open fix PR button, and one click opens this — the diff in the body, the commit verified, re-checked green by the same scan.
Live from the registry — it refreshes every 6 hours from models.dev, OpenRouter, and provider APIs.
Amazon Nova Premier
Amazon
2026-09-14
in 25d
GPT-3.5 Turbo Instruct
OpenAI
2026-09-28
in 39d
Claude Sonnet 4.5
Anthropic
2026-09-29
in 40d
Amazon Nova Canvas
Amazon
2026-09-30
in 41d
GPT-4o
OpenAI
2026-10-01
in 42d
Nano Banana
Google
2026-10-02
in 43d
Claude Haiku 4.5
Anthropic
2026-10-15
in 56d
Gemini 2.5 Flash
Google
2026-10-16
in 57d
GG
Gemini 3.7 Flash 7d ago
Google · gemini-flash · 1049k ctx
XA
Grok 4.6 8d ago
xAI · grok · 500k ctx
AM
Grok 4.6 8d ago
Amazon · grok · 500k ctx
XA
Grok Imagine Image 2.0 13d ago
xAI · grok · 8k ctx
ME
Muse Spark 1.2 Contributor 15d ago
Meta · muse · 1049k ctx
No AI provider ships a model-lifecycle feed. We do —
open the radar on your phone, or subscribe:
RSS ·
JSON ·
retirements only
Check it on every pull request.
Add a model-lifecycle check to every PR. Two options, and the difference is whether your source code leaves your runners.
Runs in your own CI; your source never leaves your runners, only model IDs go out. Inline PR annotations and a pass/fail gate.
One click, no workflow YAML or secrets to manage. We scan each PR on our servers (your code is sent to us; use the CI option if that's a dealbreaker) and post a Check Run plus a review comment. A button on the check opens a fix PR that swaps in the replacements.
The CLI is free. Cloud is $5 a year.
Plans cover the cloud inventory and alerts. The CLI itself is always free.
Everything in the terminal, free, no account
1 project, up to 15 tracked usages
Email + in-app retirement alerts
Slack · Discord · SMS · webhook alerts
CLI cloud sync + CI drift dashboard
Custom lead times (90 / 30 / 7 / 1 day)
Sign in and it emails you before a retirement.
A free account stores one project's inventory (up to 15 usages) and emails you before a model retires. Pro lifts the caps and adds Slack, Discord, SMS, webhooks, CI drift, and the API.
