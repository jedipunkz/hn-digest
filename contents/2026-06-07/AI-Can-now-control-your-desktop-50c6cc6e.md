---
source: "https://clawdcursor.com"
hn_url: "https://news.ycombinator.com/item?id=48430028"
title: "AI Can now control your desktop"
article_title: "Clawd Cursor v1.0.1 — the local MCP server for safe desktop control"
author: "AmDab"
captured_at: "2026-06-07T01:00:17Z"
capture_tool: "hn-digest"
hn_id: 48430028
score: 3
comments: 0
posted_at: "2026-06-06T23:14:25Z"
tags:
  - hacker-news
  - translated
---

# AI Can now control your desktop

- HN: [48430028](https://news.ycombinator.com/item?id=48430028)
- Source: [clawdcursor.com](https://clawdcursor.com)
- Score: 3
- Comments: 0
- Posted: 2026-06-06T23:14:25Z

## Translation

タイトル: AI がデスクトップを制御できるようになりました
記事のタイトル: Clawd Cursor v1.0.1 — 安全なデスクトップ制御のためのローカル MCP サーバー
説明: Clawd Cursor v1.0.1 — エージェントに安全なデスクトップ制御を提供するローカル MCP サーバー。 6 つのコンパクトな複合ツール (コンピュータ、アクセシビリティ、ウィンドウ、システム、ブラウザ、タスク)、MCP stdio + HTTP、単一の safety.evaluate() チョークポイント、テレメトリなし、Windows + macOS + Linux。

記事本文:
v1.0.1 — 最新の安定版
カーソルとキーボード
あらゆるAIエージェントにとって
どのモデルでも。どのアプリでも。 1 つの MCP エントリ。地元限定。 6 つのコンパクトなツール、単一の安全チョークポイント、テレメトリーなし。まずは安価なパス — ピクセルの前にアクセシビリティ、視覚は最後の手段としてのみ。
GitHub で見る
—
クイックスタート
6 ツールボックス — 複合ツール (推奨)
94 ツール - 詳細 (互換/デバッグ)
3 オペレーティング システム
2つの使い方
自分で実行するか、エージェントに渡します。
わかりやすい英語を入力し、アクションを出力します。
1 つの MCP エントリ、デスクトップ コントロールがネイティブ ツールとして表示されます。
同じツール、3 つのエントリ形状。インストール中に 1 回選択します。
AI はエディター (Claude Code、Cursor、Windsurf、Zed) に組み込まれています。エディターは、stdio 経由でオンデマンドで clawdcursor を生成します。デーモンもポートもありません。
{
"mcpサーバー": {
"爪カーソル": {
"コマンド": "爪カーソル",
"args": ["mcp", "--compact"]
}
}
}
6 / 94 コンパクト / 粒状ツール
標準デジタルトランスポート
clawdcursor エージェント — 自律型デーモン
clawdcursor は独自の LLM ブレイン ( Doctor 経由で設定) をもたらします。無人実行、スケジュールされたタスク、マルチプロセス オーケストレーション用。
clawdcursor ドクターを実行する · プロバイダーを選択する
タスクを 127.0.0.1:3847/mcp に POST します。
clawdcursor エージェント --no-llm — BYO 脳
エージェントはすでに頭脳を持っています。必要なのは HTTP ツールだけです。同じデーモンですが、組み込みエージェント ループはありません。
clawdcursor エージェント --no-llm を実行します
ステートレス — セッションの初期化は必要ありません
ピクセル前の A11y ツリー。必要なときだけ視界を確保。
a11y ツリーを読み取り、要素名に基づいて動作します。スクリーンショットもビジョン LLM もありません。
ツリーがまばらな場合は OCR、ピクセルが必要な場合はスクリーンショット、キャンバス UI のみのビジョン。
すべてのツールは、1 つの安全層を介してゲートを呼び出します。破壊的な行為には確認が必要です。
推奨されるサーフェス — コンピューター、アクセシビリティ、ウィンドウ、システム、ブラウザー、タスク。カタログは粒状のツール表面よりも約 12 倍小さい。
Windows、macOS、Linux ですね

単一のインターフェイスです。 Linux は X11 と Wayland をカバーします。
TCC セーフ。 clawdcursor 付与は、アクセシビリティと画面録画を処理します。
ネイティブ UIA + Windows.Media.Ocr。 x64 と ARM64。
X11とウェイランド。 a11yにはAT-SPI、OCRにはTesseract。
名前をクリックし、ラベルで入力し、画面を読みます。 A11y が最初で、フォールバックとして OCR。
プラットフォーム対応のキーの組み合わせ — macOS では Cmd、それ以外の場合は Ctrl。 LLMコストはかかりません。
N 個の決定論的ツール呼び出しを単一の保護されたセーフティゲートバッチにまとめます。 Nコール→1。
コンパクトな形式 (推奨):computer({ "action": "key", "combo": "mod+s" }) — ~1,500 トークンのカタログ。
詳細な形式 (互換 / デバッグ): key_press({ "key": "mod+s" }) — 94 個の個別ツール、動詞ごとに 1 つのスキーマ。
どちらも、同じ safety.evaluate() チョークポイントを通じて同じ効果を生成します。
( --compact の代わりに) --granular を渡して、MCP 上に粒度の高いサーフェスを公開します。
すべてのパラメータについては、schema.snapshot.json を参照してください。
# インストールとセットアップ
clawdcursor 同意 # ワンタイムデスクトップコントロール認証
clawdcursor give # macOS アクセシビリティ + 画面録画プロンプト
clawdcursor Doctor # 権限を確認し、AI プロバイダーを構成する
clawdcursor status # 準備状況チェック (同意、権限、AI 構成)
# 実行
clawdcursor mcp # エディター ホスト用の stdio MCP サーバー
clawdcursor mcp --compact # 同じ、6 つの複合ツールを使用 (推奨)
clawdcursor エージェント # HTTP MCP デーモン (:3847/mcp)、オプションの組み込み LLM
clawdcursor エージェント --no-llm # ツール表面のみ — エージェントは独自の頭脳を持ちます
clawdcursor stop # すべての実行モードを停止します
clawdcursor uninstall # すべての clawdcursor 設定とデータを削除します
クイックスタート
数秒でインストールできます。あらゆるOS。
$
macOS、Linux、Windows で動作します。 npm i -g clawdcursor が推奨されるインストールです (ノード 20 以上が必要)。 macOS では、ネイティブのアクセシビリティ ヘルパーに対して clawdcursor Grant も実行します。ワンライナーによりノード i がインストールされます

f が必要であり、グローバル シムをリンクします。インストール後は、clawdcursor accept --accept を実行してから、clawdcursor Doctor を実行します。
エージェントに安全なデスクトップ制御を提供します。
オープンソース。どのモデルでも。ローカルホストのみ。テレメトリはありません。
GitHub でスターを付ける
爪カーソル v1.0.1 · ローカルホストのみ · MIT ·
ホーム ·
GitHub ·
変更履歴 ·
不和

## Original Extract

Clawd Cursor v1.0.1 — the local MCP server that gives any agent safe desktop control. 6 compact compound tools (computer, accessibility, window, system, browser, task), MCP stdio + HTTP, single safety.evaluate() chokepoint, no telemetry, Windows + macOS + Linux.

v1.0.1 — latest stable
A cursor and a keyboard
for any AI agent
Any model. Any app. One MCP entry. Local-only. 6 compact tools, single safety chokepoint, no telemetry. Cheap paths first — accessibility before pixels, vision only as a last resort.
View on GitHub
—
Quick Start
6 Toolbox — compound tools (recommended)
94 Tools — granular (compat / debug)
3 Operating Systems
Two ways to use it
Run it yourself, or hand it to your agent.
Plain English in, actions out.
One MCP entry, desktop control appears as native tools.
Same tools, three entry shapes. Pick once during install.
AI lives in your editor (Claude Code, Cursor, Windsurf, Zed). Editor spawns clawdcursor on demand over stdio. No daemon, no port.
{
"mcpServers": {
"clawdcursor": {
"command": "clawdcursor",
"args": ["mcp", "--compact"]
}
}
}
6 / 94 Compact / Granular tools
stdio Transport
clawdcursor agent — autonomous daemon
clawdcursor brings its own LLM brain (configured via doctor ). For unattended runs, scheduled tasks, multi-process orchestration.
Run clawdcursor doctor · pick a provider
POST tasks to 127.0.0.1:3847/mcp
clawdcursor agent --no-llm — BYO brain
Your agent already has a brain — you just want HTTP tools. Same daemon, no built-in agent loop.
Run clawdcursor agent --no-llm
Stateless — no session init needed
A11y tree before pixels. Vision only when needed.
Read the a11y tree, act on element names. No screenshot, no vision LLM.
OCR when the tree is sparse, screenshot when you need pixels, vision only for canvas UIs.
Every tool call gates through one safety layer. Destructive actions need confirmation.
The recommended surface — computer , accessibility , window , system , browser , task . ~12× smaller catalog than the granular Tools surface.
Windows, macOS, Linux behind a single interface. Linux covers X11 and Wayland.
TCC-safe. clawdcursor grant handles Accessibility + Screen Recording.
Native UIA + Windows.Media.Ocr. x64 and ARM64.
X11 and Wayland. AT-SPI for a11y, Tesseract for OCR.
Click by name, type by label, read screen. A11y first, OCR as fallback.
Platform-aware key combos — Cmd on macOS, Ctrl elsewhere. No LLM cost.
Collapse N deterministic tool calls into a single guarded, safety-gated batch. N calls → 1.
Compact form (recommended): computer({ "action": "key", "combo": "mod+s" }) — ~1,500 tokens of catalog.
Granular form (compat / debug): key_press({ "key": "mod+s" }) — 94 individual tools, one schema per verb.
Both produce identical effects through the same safety.evaluate() chokepoint.
Pass --granular (instead of --compact ) to expose the granular surface over MCP.
See schema.snapshot.json for every parameter.
# Install & setup
clawdcursor consent # one-time desktop-control authorization
clawdcursor grant # macOS Accessibility + Screen Recording prompts
clawdcursor doctor # verify permissions, configure AI provider
clawdcursor status # readiness check (consent, permissions, AI config)
# Run
clawdcursor mcp # stdio MCP server for editor hosts
clawdcursor mcp --compact # same, with 6 compound tools (recommended)
clawdcursor agent # HTTP MCP daemon at :3847/mcp, optional built-in LLM
clawdcursor agent --no-llm # tool surface only — your agent brings its own brain
clawdcursor stop # stop every running mode
clawdcursor uninstall # remove all clawdcursor config and data
Quick Start
Install in seconds. Any OS.
$
Works on macOS, Linux, and Windows. npm i -g clawdcursor is the recommended install (Node 20+ required). On macOS also run clawdcursor grant for the native accessibility helper. The one-liner installs Node if needed and links the global shim. After any install run clawdcursor consent --accept then clawdcursor doctor .
Give your agent safe desktop control.
Open source. Any model. Localhost only. No telemetry.
Star on GitHub
clawd cursor v1.0.1 · localhost only · MIT ·
Home ·
GitHub ·
Changelog ·
Discord
