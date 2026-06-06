---
source: "https://hermes-agent.org/"
hn_url: "https://news.ycombinator.com/item?id=48419000"
title: "Hermes Agent – Open-Source AI Agent with Persistent Memory"
article_title: "Hermes Agent — Open-Source AI Agent with Persistent Memory"
author: "SeriousM"
captured_at: "2026-06-06T00:55:15Z"
capture_tool: "hn-digest"
hn_id: 48419000
score: 3
comments: 0
posted_at: "2026-06-05T22:09:38Z"
tags:
  - hacker-news
  - translated
---

# Hermes Agent – Open-Source AI Agent with Persistent Memory

- HN: [48419000](https://news.ycombinator.com/item?id=48419000)
- Source: [hermes-agent.org](https://hermes-agent.org/)
- Score: 3
- Comments: 0
- Posted: 2026-06-05T22:09:38Z

## Translation

タイトル: Hermes Agent – 永続メモリを備えたオープンソース AI エージェント
記事のタイトル: Hermes Agent — 永続メモリを備えたオープンソース AI エージェント
説明: プロジェクトを記憶し、自動的にスキルを構築し、Telegram、Discord などで連絡を取る自己ホスト型 AI エージェント。 MITライセンス。追跡はありません。

記事本文:
エルメス エージェント — 永続メモリを備えたオープンソース AI エージェント ヘルメス エージェント ホームの機能
MLOps
クイックスタート
GitHub
言語 English 中文 繁體中文 日本語 Español Português Deutsch Français 한국어 Italiano العربية Nederlands Svenska Norsk オープン ソース · MIT ライセンス · By Nous Research ヘルメス エージェント
あなたとともに成長する AI エージェント。
これをサーバーにインストールし、メッセージング アカウントを与えると、永続的なパーソナル エージェントとなり、プロジェクトを学習し、独自のスキルを構築し、どこにいても連絡を取ることができます。チャットボットではありません。副操縦士ではありません。あなたのマシン上に常駐し、日々賢くなっていくエージェント。
詳細はこちら 40 以上の組み込みスキル 5 つのチャット プラットフォーム MIT ライセンス 100% 自己ホスト型 必要なものすべてを 1 つのエージェントで実現
Hermes Agent は、永続メモリ、自動スキル作成、マルチプラットフォーム対応を 1 つのオープンソース パッケージに統合します。
すべてのセッションにわたって、設定、プロジェクト、環境を記憶します。実行時間が長ければ長いほど、ユーザーのことをよりよく理解できるようになります。毎回コンテキストを再説明する必要はありません。
Hermes は難しい問題を解決すると、再利用可能なスキル ドキュメントを作成するので、方法を忘れることはありません。スキルは検索、共有可能で、agentkills.io オープン スタンダードと互換性があります。
Telegram、Discord、Slack、WhatsApp、Signal、CLI をすべて単一のゲートウェイ プロセスから接続します。ボイスメモの文字起こし、クロスプラットフォームでの継続。 Telegram で会話を開始し、端末で会話を受け取ります。
あらゆるプラットフォームに配信できる組み込みの cron スケジューラ。日次レポート、夜間バックアップ、週次監査、朝のブリーフィングを設定し、すべて無人で実行します。
並列ワークストリーム用に分離されたサブエージェントを生成します。それぞれが独自の会話と端末を取得します。マルチステップのパイプラインを RPC 経由でゼロコンテキストコストのターンに集約します。
Web検索、ページ抽出

、ブラウザーの完全な自動化 — ナビゲート、クリック、入力、スクリーンショット。さらに、視覚分析、画像生成、テキスト読み上げ、およびマルチモデル推論も含まれます。
すぐに使える強力な機能
サンドボックス化されたコードの実行からリアルタイムのブラウザ制御まで、Hermes Agent がすべてを処理します。
Hermes Agent は、Nous Research によって構築され、2026 年 2 月にリリースされたオープンソースの自律型 AI エージェントです。これは、IDE に結び付けられたコーディング コパイロットや、単一 API のチャットボット ラッパーではありません。 Hermes はサーバー上に常駐し、学習した内容を記憶し、実行時間が長くなるほど能力が向上します。 Linux、macOS、WSL2 をサポートしており、前提条件はなく、単一のcurl コマンドですべてが自動的にインストールされます。
オープンソース · MIT ライセンス · Nous Research によって構築されました。すべてのデータはマシン上に残ります。テレメトリー、追跡、クラウドロックインはありません。
タスクの自動化を超えて — Hermes Agent は、トレーニング データの生成、RL 実験の実行、および微調整のための軌跡のエクスポートのためのプラットフォームです。
自動チェックポイント設定と並行して、何千ものツール呼び出し軌跡を生成します。構成可能なワーカー、バッチ サイズ、ツールセットの配布。
エージェントの動作に関する強化学習のための Atropos の統合。あらゆるモデル アーキテクチャをトレーニングするための 11 個のツール呼び出しパーサー。
微調整のために会話を ShareGPT 形式でエクスポートします。軌道圧縮は、トレーニング データをトークン バジェットに適合させます。
前提条件はありません。 Linux、macOS、WSL2で動作します。すべてを自動的にインストールします。
uv、Python 3.11 をインストールし、リポジトリをクローンし、すべてをセットアップします。 sudo は必要ありません。
Nous Portal (OAuth)、OpenRouter (API キー)、または独自のエンドポイントに接続します。インタラクティブなセットアップ ウィザード。
ツール、メモリ、スキルを備えた完全な対話型 CLI。それでおしまい。
Telegram、Discord、Slack、または WhatsApp への接続について説明します。 systemd サービスとして実行されます。
ラを引っ張る

変更をテストし、依存関係を再インストールします。いつでも実行して、新機能や修正を入手できます。
⚠️ ネイティブ Windows サポートは実験的です。 WSL2をインストールし、そこからHermes Agentを実行してください。
自己ホスト型、オープンソース、永久無料。インストールするコマンドは 1 つ、起動するコマンドは 1 つです。
ヘルメスエージェント あなたとともに成長するAIエージェント。
Hermes Agent は、Nous Research によるオープンソースの自律型 AI エージェントで、実行時間が長ければ長いほど賢くなります。
© 2026 エルメスエージェント。無断転載を禁じます。

## Original Extract

Self-hosted AI agent that remembers your projects, builds skills automatically, and reaches you on Telegram, Discord & more. MIT license. No tracking.

Hermes Agent — Open-Source AI Agent with Persistent Memory Hermes Agent Home Features
MLOps
Quick Start
GitHub
Language English 中文 繁體中文 日本語 Español Português Deutsch Français 한국어 Italiano العربية Nederlands Svenska Norsk Open Source · MIT License · By Nous Research Hermes Agent
The AI agent that grows with you.
Install it on your server, give it your messaging accounts, and it becomes a persistent personal agent — learning your projects, building its own skills, and reaching you wherever you are. Not a chatbot. Not a copilot. An agent that lives on your machine and gets smarter every day.
Learn More 40+ Built-in Skills 5 Chat Platforms MIT License 100% Self-Hosted Everything You Need in One Agent
Hermes Agent combines persistent memory, automated skill creation, and multi-platform reach into a single open-source package.
Remembers your preferences, projects, and environment across every session. The longer it runs, the better it knows you — no re-explaining context every time.
When Hermes solves a hard problem, it writes a reusable skill document so it never forgets how. Skills are searchable, shareable, and compatible with the agentskills.io open standard.
Connect Telegram, Discord, Slack, WhatsApp, Signal, and CLI — all from a single gateway process. Voice memo transcription, cross-platform continuation. Start a conversation on Telegram, pick it up in your terminal.
Built-in cron scheduler with delivery to any platform. Set up daily reports, nightly backups, weekly audits, and morning briefings — all running unattended.
Spawn isolated sub-agents for parallel workstreams. Each gets its own conversation and terminal. Collapse multi-step pipelines into zero-context-cost turns via RPC.
Web search, page extraction, full browser automation — navigate, click, type, screenshot. Plus vision analysis, image generation, text-to-speech, and multi-model reasoning.
Powerful Capabilities Out of the Box
From sandboxed code execution to real-time browser control — Hermes Agent handles it all.
Hermes Agent is an open-source autonomous AI agent built by Nous Research and released in February 2026. It's not a coding copilot tethered to an IDE or a chatbot wrapper around a single API. Hermes lives on your server, remembers what it learns, and gets more capable the longer it runs. It supports Linux, macOS, and WSL2 — with no prerequisites, installing everything automatically via a single curl command.
Open source · MIT License · Built by Nous Research. All data stays on your machine. No telemetry, no tracking, no cloud lock-in.
Beyond task automation — Hermes Agent is a platform for generating training data, running RL experiments, and exporting trajectories for fine-tuning.
Generate thousands of tool-calling trajectories in parallel with automatic checkpointing. Configurable workers, batch sizes, and toolset distributions.
Atropos integration for reinforcement learning on agent behaviors. 11 tool-call parsers for training any model architecture.
Export conversations in ShareGPT format for fine-tuning. Trajectory compression fits training data into token budgets.
No prerequisites. Works on Linux, macOS & WSL2. Installs everything automatically.
Installs uv, Python 3.11, clones the repo, sets up everything. No sudo needed.
Connect to Nous Portal (OAuth), OpenRouter (API key), or your own endpoint. Interactive setup wizard.
Full interactive CLI with tools, memory, and skills. That's it.
Walk through connecting Telegram, Discord, Slack, or WhatsApp. Runs as a systemd service.
Pulls the latest changes and reinstalls dependencies. Run anytime to get new features and fixes.
⚠️ Native Windows support is experimental. Please install WSL2 and run Hermes Agent from there.
Self-hosted, open source, and free forever. One command to install, one command to start.
Hermes Agent The AI agent that grows with you.
Hermes Agent is an open-source autonomous AI agent by Nous Research that grows smarter the longer it runs.
© 2026 Hermes Agent. All rights reserved.
