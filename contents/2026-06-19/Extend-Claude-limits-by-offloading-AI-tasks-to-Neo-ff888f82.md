---
source: "https://heyneo.com/claude-code"
hn_url: "https://news.ycombinator.com/item?id=48599640"
title: "Extend Claude limits by offloading AI tasks to Neo"
article_title: "Claude Code + NEO: Ship production AI from your terminal"
author: "gauravvij137"
captured_at: "2026-06-19T16:03:36Z"
capture_tool: "hn-digest"
hn_id: 48599640
score: 2
comments: 1
posted_at: "2026-06-19T15:20:44Z"
tags:
  - hacker-news
  - translated
---

# Extend Claude limits by offloading AI tasks to Neo

- HN: [48599640](https://news.ycombinator.com/item?id=48599640)
- Source: [heyneo.com](https://heyneo.com/claude-code)
- Score: 2
- Comments: 1
- Posted: 2026-06-19T15:20:44Z

## Translation

タイトル: AI タスクを Neo にオフロードしてクロードの制限を拡張
記事タイトル: Claude Code + NEO: 端末から生産 AI を出荷
説明: neo-mcp をインストールし、NEO を Claude Code に登録し、ターミナルを離れることなく RAG 監査、微調整、評価、およびパイプラインのデバッグを委任します。

記事本文:
Claude Code + NEO: 端末から本番 AI を出荷 ホーム コミュニティ ブログ ユースケース VS Code カーソル サインイン 開始する クロードの制限を拡張する
AIタスクをNeoにオフロードすることで
NEO の MCP サーバーを Python 3.11 以降の環境に追加します。
NEO ダッシュボードを開き、キーを作成してコピーします。キーは sk-v1-… のようになります。
1 つのコマンドで NEO を登録し、新しいチャットで作業を依頼するだけです。
クロード mcp add --scope user neo \
-e NEO_SECRET_KEY=sk-v1-あなたのキー \
-- python3 -m neo_mcp Cursor、VS Code、または別の MCP クライアントを使用しますか? neo-mcp セットアップを参照してください。
タスク: CPU のみの Azure VM (2 コア、7.7 GB RAM、GPU なし) で音声テキスト変換モデルのベンチマークを実行します。 Claude Code のみが明らかな HuggingFace + PyTorch パスに到達し、リアルタイムで反復されました。 NEO は最初に 2 分間のリサーチを行った後、AVX2 に最適化された CPU カーネルとして、同じタスク、同じマシンとして ONNX ランタイムを選択しました。
Gaurav Vij によるベンチマーク · 全文を読む
一度インストールすると、任意の Claude Code セッションから ML 作業を NEO に委任します。
セットアップ ガイドを読む クロード コードで動作する · モデル コンテキスト プロトコル · Python 3.11+
Neo は、AI アプリ、エージェントを構築し、AI 製品と ML モデル全体で実験、評価、最適化を実行できる完全自律型 AI エンジニアです。

## Original Extract

Install neo-mcp, register NEO with Claude Code, and delegate RAG audits, fine-tunes, evals, and pipeline debugging without leaving the terminal.

Claude Code + NEO: Ship production AI from your terminal Home Community Blog Use Cases VS Code Cursor Sign In Get Started Extend Claude limits
by offloading AI tasks to Neo
Add NEO's MCP server to any environment with Python 3.11+ .
Open your NEO dashboard, create a key, and copy it. Keys look like sk-v1-… .
Register NEO with one command, then just ask in a new chat to ship work.
claude mcp add --scope user neo \
-e NEO_SECRET_KEY=sk-v1-your-key \
-- python3 -m neo_mcp Using Cursor, VS Code, or another MCP client? See the neo-mcp setup .
The task: benchmark a speech-to-text model on a CPU-only Azure VM — 2 cores, 7.7 GB RAM, no GPU. Claude Code alone reached for the obvious HuggingFace + PyTorch path and iterated in real time. NEO spent two minutes researching first, then chose ONNX Runtime for its AVX2-optimized CPU kernels — same task, same machine.
Benchmark by Gaurav Vij · Read the full writeup
Install once, then delegate ML work to NEO from any Claude Code session.
Read the setup guide Works with Claude Code · Model Context Protocol · Python 3.11+
Neo is a fully autonomous AI Engineer capable of building AI apps, agents, performing experiments, evals, optimizations across your AI products and ML models.
