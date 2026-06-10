---
source: "https://eatmydata.ai/"
hn_url: "https://news.ycombinator.com/item?id=48472867"
title: "Show HN: Eatmydata.ai – Local-First Question-to-SQL-to-Dashboard AI"
article_title: "eatmydata"
author: "dennis16384"
captured_at: "2026-06-10T10:30:29Z"
capture_tool: "hn-digest"
hn_id: 48472867
score: 2
comments: 2
posted_at: "2026-06-10T07:43:21Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Eatmydata.ai – Local-First Question-to-SQL-to-Dashboard AI

- HN: [48472867](https://news.ycombinator.com/item?id=48472867)
- Source: [eatmydata.ai](https://eatmydata.ai/)
- Score: 2
- Comments: 2
- Posted: 2026-06-10T07:43:21Z

## Translation

タイトル: HN を表示: Eatmydata.ai – Local-First Question-to-SQL-to-Dashboard AI
記事タイトル: Eatmydata
HN テキスト: さらに別の「データと対話してダッシュボードを構築する」アプリ。データはブラウザーから出ません。あなたが質問すると、エージェントはブラウザ内の sqlite に対して複数の SQL クエリを生成しますが、結果は表示されず、ダッシュボード構成コードを作成します。分析するデータは、ローカル セマンティック インデックス (埋め込み生成 + 完全にローカルな sqlite ベクトル検索) でインデックス付けされます。次に、サンドボックス化された QuickJS がこのコードを実行して、バックエンドを接続せずに、ブラウザー内に豊富なダッシュボードを直接作成します。これは完全なフロントエンド アプリです (OpenRouter または他のリモート LLM を除く)。 LLM に送信されるすべてのデータは、いくつかの点で厳重にサニタイズされ、難読化されます。リモート LLM は、分析するデータの内容を決して参照しません。なぜ存在するのか - これは、私のローカルファースト AI プロジェクト、エージェント ワークフロー、およびコンテキスト データ分析の実験のためのテストベッドとして始めました。 SQL のデバッグや単純なデータの質問に対するグラフの作成に時間を無駄にしたくないとき、文字通り 10 秒以内に答えが必要なときに、手早く簡単なデータ分析を行うために毎日使用するツールに成長しました。また、Claude/ChatGPT チャットでランダム データを共有するというアイデアも好きではありませんし、仕事関連のデータセットをアップロードすることも好みません。さらに、どちらも 100,000 行の小さなデータで詰まることがよくあります。 MIT https://github.com/eatmydata-org/eatmydata で完全にオープンソース化されており、静的 Web アプリなので自分で実行できます。同梱物: - wa-sqlite から適応した SQLite OPFS、データはローカルでのみクエリされます。 - TurboQuant sqlite 用セマンティック インデックス拡張機能 (MIT ライセンス)。 - 量子化された PII 検出と生成モデルをブラウザに直接埋め込みます。 - 依存性ゼロの C および wasm-simd128 最適化における NER および埋め込み推論エンジン (以前と比較してバイナリが 1.7 倍高速かつ 38 倍軽量)

nxランタイム); - AI 生成コード用の QuickJS サンドボックス。 - Orchestrator <-> SQL Planner <-> Coder エージェント ループ。ユーザー クエリから SQL とダッシュボードを構築します。 - ダッシュボード用の Apache ECharts; - スタイルをサポートするための xslx Community エディションのフォーク (上流の OSS バージョンには存在しません)。ローカルファーストなものに興味のある方の参考になれば幸いです。

記事本文:
マイデータを食べる

## Original Extract

Yet another "talk to your data and build a dashboard" app, where data does not leave your browser. You ask a question, agents produce multiple SQL queries to in-browser sqlite, never seeing results, and write dashboard configuration code. The data you analyze will be indexed with a local semantic index (embeddings generation + sqlite vector search fully local). Next, sandboxed QuickJS runs this code to produce rich dashboards directly in your browser, no backend attached. This is a fully frontend app (except OpenRouter or other remote LLM). All data sent to LLM's is heavily sanitized and obfuscated at several points. The remote LLM never sees the contents of data it analyzes. Why does it exist - I started this is a testbed for my local-first AI projects, agentic workflows and contextual data analysis experiments. It grew into a tool I use daily for quick and dirty data analytics when I don't want to waste time debugging SQL or building charts for simple data questions, when I literally need an answer under 10s. I also don't like the idea of sharing random data in Claude/ChatGPT chat, neither uploading any work-related datasets to them. Plus they both often choke on tiny 100k rows data. Fully open-sourced under MIT https://github.com/eatmydata-org/eatmydata , run it yourself it's a static web app. What's in the box: - SQLite OPFS adapted from wa-sqlite, data queried only locally; - TurboQuant semantic indexing extension for sqlite (MIT-licensed); - Quantized PII detection and embedding generation models straight in browser; - NER and embeddings inference engines in zero-dependency C and wasm-simd128 optimizations (1.7x faster and 38x lighter binary compare to onnxruntime); - QuickJS sandbox for AI-generated code; - Orchestrator <-> SQL Planner <-> Coder agent loop that build SQL and dashboards from user query; - Apache ECharts for dashboards; - Fork of xslx Community edition to support styles (missing in OSS version upstream). Hope it'll be useful to anyone who is interested in local-first stuff.

eatmydata
