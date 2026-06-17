---
source: "https://code.intellios.ai"
hn_url: "https://news.ycombinator.com/item?id=48563214"
title: "Native Coding Agent Optimized for Local LLM and DeepSeek v4 with Vector Memory"
article_title: "cwcode — a terminal coding agent"
author: "coolwulf"
captured_at: "2026-06-17T01:02:45Z"
capture_tool: "hn-digest"
hn_id: 48563214
score: 1
comments: 0
posted_at: "2026-06-16T22:36:54Z"
tags:
  - hacker-news
  - translated
---

# Native Coding Agent Optimized for Local LLM and DeepSeek v4 with Vector Memory

- HN: [48563214](https://news.ycombinator.com/item?id=48563214)
- Source: [code.intellios.ai](https://code.intellios.ai)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T22:36:54Z

## Translation

タイトル: ローカル LLM およびベクター メモリを備えた DeepSeek v4 向けに最適化されたネイティブ コーディング エージェント
記事のタイトル: cwcode — ターミナルコーディングエージェント

記事本文:
cwcode — ターミナルコーディングエージェント
v1.11
それは何ですか
なぜ違うのか
スクリーンショット
インストール
ツール
よくある質問
CWコード
DeepSeek V4 Pro を中心に構築されたターミナル コーディング エージェント、
Qwen3.6‑27B、Kimi、Azure、および OpenAI を話すその他のもの
チャットAPI。
Goで書かれています。あなたのターミナルに住んでいます。実際のコードを編集します。自分自身から回復します
間違い。 1 時間稼働し続けると約 0.40 ドルかかります。
cwcode は、OpenAI 互換のチャット エンドポイントを駆動する Bubbletea TUI です
ツールを使用したコーディング エージェントとして。 DeepSeek 用のプロファイル (Pro および
Flash)、Azure OpenAI、Kim forcoding、およびローカル vLLM /
ホームサーバー上の Qwen3.6-27B の llama.cpp プロファイル。プロファイルの切り替え
セッション中は 1 つのスラッシュ コマンドです。
bash、ファイル編集、glob、grep、Web フェッチ、ヘッドレス Chrome フェッチを備えています。
(実際のブラウザを介して CDP 経由で駆動)、サブエージェント、永続的
セマンティック メモリ ストア、巻き戻しを伴うコンテンツ アドレス指定チェックポイント、
プラン/コード モードの切り替え、および自律的な目標ループ。ツールレジストリ
は 600 行で、新しいツールを追加するのは 2 メソッドの Go インターフェイスです。
SaaSではありません。アカウントもテレメトリーもリモートコントロールもありません
飛行機。 API キーは ~/.cwcode/config.json にあります。あなたの
セッション履歴は ~/.cwcode/sessions/ にあります。もしあなたの
ネットワークがダウンしており、モデルのエンドポイントがローカルであるため、エージェントは維持されます。
働いています。
線量予測コードベースの実際の作業中にキャプチャされたエージェント:
統合された Go テストに対する edit_file の変更を提案する
diff はインラインで強調表示され、推論トレースは以下にストリーミングされ、
現在のタスク リストがペインの下部に固定されます。
プラットフォーム用に事前に構築されたバイナリを次の場所からダウンロードします。
Google ドライブのリリース フォルダー
(現在のビルド: v1.11; macOS arm64 / amd64 および Windows amd64)。落としてください
PATH 上のどこかに実行可能にします。
カール -L <ダウンロード URL> -o ~/.local/bin/cwc

頌歌
chmod +x ~/.local/bin/cwcode
cwcode -バージョン
OpenAI 互換エンドポイント (DeepSeek API キー、
Azure デプロイメント、ローカル vLLM、またはその他の手持ちのもの)。
~/.cwcode/config.json でプロファイルを構成します。
{
"active_profile": "ディープシークプロ",
「プロファイル」: {
"ディープシークプロ": {
"プロバイダー": "ディープシーク",
"エンドポイント": "https://api.deepseek.com",
"モデル": "ディープシーク-v4-プロ",
"api_key": "sk-...",
"ctx_size": 262144
}
}
}
実行してください。
cwcode # バブルティー TUI
cwcode -p "バグを修正" # ワンショット、セッションなし
cwcode - continue # 最新のセッションを再開します
cwcode -plain # stdout REPL (TUI なし)
内蔵ツール

## Original Extract

cwcode — terminal coding agent
v1.11
What it is
Why it’s different
Screenshot
Install
Tools
FAQ
cwcode
A terminal coding agent built around DeepSeek V4 Pro,
Qwen3.6‑27B, Kimi, Azure, and anything else that speaks OpenAI’s
chat API.
Written in Go. Lives in your terminal. Edits real code. Recovers from its own
mistakes. Costs about $0.40 to leave running for an hour.
cwcode is a Bubbletea TUI that drives any OpenAI-compatible chat endpoint
as a tool-using coding agent. It ships with profiles for DeepSeek (Pro and
Flash), Azure OpenAI, Kimi for Coding, and a local vLLM /
llama.cpp profile for Qwen3.6-27B on a home server. Switching profiles
mid-session is one slash command.
It has bash, file edit, glob, grep, web fetch, headless-Chrome fetch
(driven via CDP through your real browser), sub-agents, a persistent
semantic-memory store, content-addressed checkpoints with rewind, a
plan/code mode toggle, and an autonomous goal loop. The tool registry
is six hundred lines and adding a new tool is a two-method Go interface.
It is not a SaaS. There is no account, no telemetry, no remote control
plane. Your API key sits in ~/.cwcode/config.json . Your
session history sits in ~/.cwcode/sessions/ . If your
network is down and the model endpoint is local, the agent keeps
working.
Captured during real work on our dose-prediction codebase: the agent
proposing an edit_file change to a Go test, with a unified
diff highlighted inline, the reasoning trace streaming below, and the
current task list pinned to the bottom of the pane.
Download a pre-built binary for your platform from the
Google Drive release folder
(current build: v1.11; macOS arm64 / amd64 and Windows amd64). Drop it
somewhere on your PATH and make it executable:
curl -L <download-url> -o ~/.local/bin/cwcode
chmod +x ~/.local/bin/cwcode
cwcode -version
You’ll need an OpenAI-compatible endpoint (DeepSeek API key,
Azure deployment, local vLLM, or whatever else you have on hand).
Configure a profile in ~/.cwcode/config.json :
{
"active_profile": "deepseek-pro",
"profiles": {
"deepseek-pro": {
"provider": "deepseek",
"endpoint": "https://api.deepseek.com",
"model": "deepseek-v4-pro",
"api_key": "sk-...",
"ctx_size": 262144
}
}
}
Run it.
cwcode # Bubbletea TUI
cwcode -p "fix the bug" # one-shot, no session
cwcode -continue # resume the most recent session
cwcode -plain # stdout REPL (no TUI)
Built-in tools
