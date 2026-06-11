---
source: "https://headroom.walls.sh/hook"
hn_url: "https://news.ycombinator.com/item?id=48492114"
title: "Claude Code's statusLineHook: read rate limits locally without any API calls"
article_title: "How Claude Code's status line hook works — and how to use it"
author: "patwalls"
captured_at: "2026-06-11T16:26:48Z"
capture_tool: "hn-digest"
hn_id: 48492114
score: 1
comments: 0
posted_at: "2026-06-11T15:55:57Z"
tags:
  - hacker-news
  - translated
---

# Claude Code's statusLineHook: read rate limits locally without any API calls

- HN: [48492114](https://news.ycombinator.com/item?id=48492114)
- Source: [headroom.walls.sh](https://headroom.walls.sh/hook)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T15:55:57Z

## Translation

タイトル: Claude Code の statusLineHook: API 呼び出しなしでローカルでレート制限を読み取る
記事のタイトル: Claude Code のステータス ライン フックの仕組みとその使用方法
説明: Claude コードは、レート制限データを計算するたびにシェル コマンドを実行する statusLineHook 構成を公開します。ここ

記事本文:
← ヘッドルーム.ウォール.sh
Claude Code のステータス ライン フックの仕組みとその使用方法
Claude Code は、5 時間のセッションと 7 日間の週あたりのレート制限を内部的に追跡し、独自の /usage コマンドを介してそれらの数値を公開します。しかし、あまり知られていないメカニズム、statusLineHook があり、これを使用すると、API をポーリングしたりコマンドを実行したりせずに、そのデータを自動的にキャプチャできます。仕組みは次のとおりです。
Claude Code は起動時に ~/.claude/settings.json を読み取ります。その構成キーの 1 つは statusLineHook です。これは、クロード コードがステータス ラインを更新するたびに実行するシェル コマンド文字列です。使用状況データは環境変数 HOOK_STATUS_LINE として渡されます。
HOOK_STATUS_LINE の値は、レート制限とセッション情報を含む JSON 文字列です。一般的な値は次のようになります。
{
"sessionUsagePct": 12.4、
「週間使用率」: 67.1、
"contextUsagePct": 8.2、
"モデル": "クロード-ソネット-4-6",
「コストUSD」: 1.24、
"sessionResetAt": "2026-06-11T18:42:00Z",
"weeklyResetAt": "2026-06-14T00:00:00Z"
}
これらは、/usage を実行するときにクロード コードがレンダリングする数値と同じです。内部のレート制限状態から直接取得され、各プロンプトが処理されるたびにリアルタイムで更新されます。
このデータをキャプチャするには、 statusLineHook を ~/.claude/settings.json に追加します。最も単純なフックは、JSON をファイルに書き込みます。
{
"statusLineHook": "printf '%s' \"$HOOK_STATUS_LINE\" > ~/.claude/usage.json"
}
新しいクロード コード セッションを保存して開始した後、クロード コードがプロンプトを処理するたびに、ファイル ~/.claude/usage.json が更新されます。次のようにして読むことができます。
猫 ~/.claude/usage.json | jq 。
jq '.sessionUsagePct' ~/.claude/usage.json # セッション %
jq '.weeklyUsagePct' ~/.claude/usage.json # 毎週の%
これが API ポーリングよりも優れている理由
クロード コードの使用状況を監視する一般的なアプローチは、Anthropic 使用状況 API を呼び出すことです。それには必要なもの

API キーを取得し、ネットワーク リクエストを作成し、監視ツールが独自のレート制限に影響を与えるリスクを負います。
フックのアプローチには次のような欠点はありません。
ネットワーク呼び出しはゼロ — データはローカルであり、Claude Code 自体によって書き込まれます
認証情報なし — フックは Claude Code のプロセスで実行されます。トークンは必要ありません
リアルタイム — ポーリング間隔ではなく、プロンプトごとに更新されます
プライバシーの保護 — データがマシンから流出することはありません
次のコマンドを使用して、フックベースのツールからのネットワーク アクティビティがないことを確認できます。
nettop -p ヘッドルーム # またはツールの名前
チェーンフック
すでに statusLineHook が設定されている場合、Claude Code はそれを尊重します。フック フィールドは文字列なので、&& または ; でコマンドをチェーンする必要があります。 。責任を持ってフックをインストールするツールは、まず既存のフックを確認し、上書きするのではなくそれに追加する必要があります。
フック メカニズムは、ポーリングが必要なために以前は実用的ではなかったクラスのツールのロックを解除します。
アンビエント ディスプレイ — メニュー バー ウィジェット、ターミナル ステータス ライン、tmux ステータスバー セグメント
しきい値アラート — セッション使用率が 80% になったときに macOS 通知を送信するシェル スクリプト
使用状況のログ - 各更新をログ ファイルに追加して、セッション履歴を構築します。
ワークフローの自動化 - コンテキスト ウィンドウのリセットをトリガーするか、コンテキストがいっぱいになったときに保存します
コスト追跡 — 簡単な追加と合計でセッション全体のコストUSDを蓄積します
最小限の tmux ステータスバー セグメント。例:
# .tmux.conf 内:
set -g status-right '#(jq -r '"'"'if .sessionUsagePct then "CC (.sessionUsagePct |round)%·(.weeklyUsagePct |round)%" else "" end'"'"' ~/.claude/usage.json 2>/dev/null)'
Headroom は、まさにこのメカニズムに基づいて構築されたネイティブ macOS メニュー バー アプリです。
一目で両方のメーターが表示されます - CC 12%·67% - 限界が近づくにつれて色分けされます。
無料、MIT、設定不要。

ヘッドルームをダウンロード — 無料
brew install --cask patwalls/tap/headroom
完全な JSON スキーマ
HOOK_STATUS_LINE で信頼できるフィールド:
sessionUsagePct // 浮動小数点、0 ～ 100、5 時間のセッション ウィンドウ
WeeklyUsagePct // 浮動小数点数、0 ～ 100、7 日の週次ウィンドウ
contextUsagePct // 浮動小数点、0 ～ 100、現在のコンテキストを埋める
model // 文字列、アクティブなモデル名
costUSD //浮動小数点数、プランで追跡されている場合のセッションコスト
sessionResetAt // ISO 8601、5 時間のウィンドウがリセットされるとき
WeeklyResetAt // ISO 8601、7d ウィンドウがリセットされるとき
すべてのフィールドがすべての更新に存在するわけではありません。使用する前に存在することを確認してください。セッションごとのコストが公開されていないプランでは、costUSD が 0 になる場合があります。
Headroom アプリはオープンソースであり、フック + ファイル読み取りアプローチの完全なリファレンス実装が示されています: github.com/patwalls/headroom 。
関連するファイルは、app/Sources/Headroom/Hook.swift (フックをインストールする) および app/Sources/Headroom/Usage.swift (JSON を読み取り、解析する) です。

## Original Extract

Claude Code exposes a statusLineHook config that runs a shell command every time it computes rate-limit data. Here

← headroom.walls.sh
How Claude Code's status line hook works — and how to use it
Claude Code tracks your 5-hour session and 7-day weekly rate limits internally, and it exposes those numbers via its own /usage command. But there's a lesser-known mechanism — statusLineHook — that lets you capture that data automatically without polling an API or running a command. Here's how it works.
Claude Code reads ~/.claude/settings.json on startup. One of its config keys is statusLineHook — a shell command string that Claude Code runs every time it updates its status line. The usage data is passed as an environment variable: HOOK_STATUS_LINE .
The value of HOOK_STATUS_LINE is a JSON string with rate-limit and session information. A typical value looks like:
{
"sessionUsagePct": 12.4,
"weeklyUsagePct": 67.1,
"contextUsagePct": 8.2,
"model": "claude-sonnet-4-6",
"costUSD": 1.24,
"sessionResetAt": "2026-06-11T18:42:00Z",
"weeklyResetAt": "2026-06-14T00:00:00Z"
}
These are the same numbers Claude Code renders when you run /usage — straight from its internal rate-limit state, updated in real time as each prompt is processed.
To capture this data, add a statusLineHook to ~/.claude/settings.json . The simplest possible hook writes the JSON to a file:
{
"statusLineHook": "printf '%s' \"$HOOK_STATUS_LINE\" > ~/.claude/usage.json"
}
After saving and starting a new Claude Code session, the file ~/.claude/usage.json will be updated every time Claude Code processes a prompt. You can read it with:
cat ~/.claude/usage.json | jq .
jq '.sessionUsagePct' ~/.claude/usage.json # session %
jq '.weeklyUsagePct' ~/.claude/usage.json # weekly %
Why this is better than API polling
The typical approach to monitoring Claude Code usage is to call the Anthropic usage API. That requires an API key, makes network requests, and runs the risk of your monitoring tool contributing to its own rate limits.
The hook approach has none of these downsides:
Zero network calls — the data is local, written by Claude Code itself
No credentials — the hook runs in Claude Code's process; you don't need a token
Real-time — updated after every prompt, not on a poll interval
Privacy-preserving — the data never leaves your machine
You can verify there's no network activity from any hook-based tool with:
nettop -p Headroom # or whatever your tool is named
Chaining hooks
If you already have a statusLineHook configured, Claude Code respects it — the hook field is a string, so you'd need to chain commands with && or ; . Tools that install hooks responsibly should check for an existing hook first and append to it rather than overwriting.
The hook mechanism unlocks a class of tools that were previously impractical because they required polling:
Ambient displays — menu bar widgets, terminal status lines, tmux statusbar segments
Threshold alerts — shell scripts that send a macOS notification at 80% session usage
Usage logging — append each update to a log file to build a session history
Workflow automation — trigger a context-window reset or save when the context fills up
Cost tracking — accumulate costUSD across sessions with a simple append-and-sum
A minimal tmux statusbar segment, for example:
# In .tmux.conf:
set -g status-right '#(jq -r '"'"'if .sessionUsagePct then "CC (.sessionUsagePct | round)%·(.weeklyUsagePct | round)%" else "" end'"'"' ~/.claude/usage.json 2>/dev/null)'
Headroom is a native macOS menu bar app built on exactly this mechanism.
It shows both meters at a glance — CC 12%·67% — color-coded as limits approach.
Free, MIT, zero config.
Download Headroom — free
brew install --cask patwalls/tap/headroom
The full JSON schema
Fields you can count on in HOOK_STATUS_LINE :
sessionUsagePct // float, 0–100, 5h session window
weeklyUsagePct // float, 0–100, 7d weekly window
contextUsagePct // float, 0–100, current context fill
model // string, active model name
costUSD // float, session cost if tracked by your plan
sessionResetAt // ISO 8601, when the 5h window resets
weeklyResetAt // ISO 8601, when the 7d window resets
Not all fields are present in every update — check for existence before using. costUSD may be 0 for plans that don't expose per-session cost.
The Headroom app is open-source and shows a complete reference implementation of the hook + file-read approach: github.com/patwalls/headroom .
The relevant files are app/Sources/Headroom/Hook.swift (installs the hook) and app/Sources/Headroom/Usage.swift (reads and parses the JSON).
