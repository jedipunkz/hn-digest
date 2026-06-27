---
source: "https://graphify.net/index.html#features"
hn_url: "https://news.ycombinator.com/item?id=48702376"
title: "Graphify – Open-Source Knowledge Graph Skill for AI Coding Assistants"
article_title: "Graphify — Open-Source Knowledge Graph Skill for AI Coding Assistants"
author: "xy008areshsu"
captured_at: "2026-06-27T23:21:16Z"
capture_tool: "hn-digest"
hn_id: 48702376
score: 2
comments: 0
posted_at: "2026-06-27T22:26:37Z"
tags:
  - hacker-news
  - translated
---

# Graphify – Open-Source Knowledge Graph Skill for AI Coding Assistants

- HN: [48702376](https://news.ycombinator.com/item?id=48702376)
- Source: [graphify.net](https://graphify.net/index.html#features)
- Score: 2
- Comments: 0
- Posted: 2026-06-27T22:26:37Z

## Translation

タイトル: Graphify – AI コーディング アシスタント向けのオープンソース ナレッジ グラフ スキル
記事のタイトル: Graphify — AI コーディング アシスタントのためのオープンソースのナレッジ グラフ スキル
説明: Graphify は、AI コーディング アシスタントがマルチモーダル コードベースを理解するのに役立つオープンソースのナレッジ グラフ スキルです。 Graphify は、Tree-sitter、NetworkX、および Leiden クラスタリングを使用して、コード、ドキュメント、論文、図をクエリ可能なグラフに抽出します。

記事本文:
𝕏
で
f
R
Y
⧉
グラフ化する
特長
Graphify — ナレッジグラフ
AIコーディングアシスタント向け
Graphify は、コード、ドキュメント、論文、図からクエリ可能なナレッジ グラフを構築することで、AI コーディング アシスタントがマルチモーダル コードベースを理解できるようにするオープンソース スキルです。
Graphify は、Claude Code、OpenAI Codex、OpenCode などの AI コーディング アシスタント向けに作成されたマルチモーダル ナレッジ グラフ ビルダーです。 Tree-sitter の静的分析と LLM 主導のセマンティック抽出を組み合わせることで、Graphify は、ソース コード、ドキュメント、研究論文、図を含むリポジトリ全体を、コードの動作とコードがそのように設計された理由の両方を説明するインタラクティブなグラフに変換します。このプロジェクトは Safi Shamsi によって保守され、寛容な MIT ライセンスの下でリリースされ、NetworkX や Tree-sitter などの広く信頼されているライブラリに基づいて構築されています。
Graphify は、静的分析、セマンティック抽出、グラフ クラスタリングを、AI コーディング アシスタントが呼び出せる単一のスキルに統合します。
コード (.py、.js、.go、.java など)、Markdown、PDF、画像を解析します。 Tree-sitter は AST、呼び出しグラフ、および docstring を抽出します。 LLM は散文から概念を抽出します。ビジョンモデルは図を読み取ります。
抽出されたすべてのノードとエッジを NetworkX グラフにマージし、セマンティック コミュニティ検出に Leiden アルゴリズムを適用します。ベクトルの埋め込みは必要ありません。
システムの中心にある最高位の「ゴッド ノード」を識別し、調査する価値のある予期しないファイル間またはドメイン間接続にフラグを立てます。
インタラクティブなgraph.html、クエリ可能なgraph.json、および人間が読めるGRAPH_REPORT.md監査レポートをエクスポートします。
Claude Code、Codex、OpenCode などの /graphify 、 /graphify query 、 /graphify path 、および /graphify Explain コマンドが同梱されています。
厳密な入力検証: http/https URL、サイズおよびタイムアウト制限のみ

ts、パスの包含、HTML エスケープされたノード ラベル — SSRF、インジェクション、および XSS に対する防御。
Graphify は複数ステージのパイプラインです。各ステージは独立したモジュールであるため、コントリビューターは任意のステップを独立して拡張できます。
サポート モジュールには、URL フェッチ用の ingest.py、セマンティック キャッシュ用の cache.py、入力検証用の security.py、ライブ アップデート用の watch.py​​、MCP プロトコル サービス用のserve.py が含まれます。
Graphify は PyPI 上で配布されます。パッケージ名は graffyy です。 CLI コマンドは grify のままです。
# Python 3.10以降が必要です
pip install グラフィファイ && グラフィファイ インストール
# 任意のプロジェクト フォルダーのナレッジ グラフを構築する
/グラフィファイ ./生
# 出力はgraphify-out/に着陸します
グラフ化アウト/
§──graph.html # インタラクティブな視覚化
§── GRAPH_REPORT.md # コアノード、サプライズ、質問の提案
§──graph.json # 永続的でクエリ可能なグラフ
└── キャッシュ/ # 増分キャッシュ
Graphify は LLM をバンドルしていません。 AI コーディング アシスタント (Claude、Codex など) によってすでに構成されているモデル API キーを使用し、セマンティック コンテンツのみを上流モデルに送信します (生のソース コードは送信しません)。
このリポジトリには、小規模なライブラリと大規模なコードとペーパーの混合コレクションの両方で Graphify をデモンストレーションする、再現可能なコーパスが付属しています。
HTTP トランスポート層をモデル化する 6 つの Python ファイル。結果: 144 ノード、330 エッジ、6 コミュニティ。ゴッドノード: Client 、 AsyncClient 、 Response 、 Request 。サプライズエッジ: DigestAuth → Response 。
3 つの GPT フレームワーク リポジトリ + 5 つの注目論文 + 4 つの図 (~52 ファイル、~92,000 ワード)。結果: 285 ノード、340 エッジ、53 コミュニティ。平均クエリコストは約 1.7,000 トークン対約 123,000 ナイーブ — 71.5 倍の削減。
Graphify が、コード インテリジェンス領域における隣接するオープンソース プロジェクトとどのように関係するか。
Graphify は MIT ライセンスに基づいてリリースされています。その中心的な依存関係 — NetworkX (BSD) と

Tree-sitter (MIT) — すべてが寛容なオープンソース ライセンスであり、競合はありません。プロジェクトはテレメトリを実行しません。唯一のアウトバウンド ネットワーク呼び出しはセマンティック抽出ステップであり、独自に構成された AI モデル API キーを使用します。ドキュメントの意味論的な記述のみが送信され、生のソース コードは送信されません。 URL は http/https に制限され、ダウンロードにはサイズと時間の制限があり、出力パスは包含チェックが行われ、SSRF、Cypher インジェクション、XSS を防ぐためにノード ラベルは HTML エスケープされます。
Graphify がどのようにナレッジ グラフを構築、クラスタリングし、AI コーディング アシスタントに提供するかについての詳細なガイド。
AI コーディング アシスタント向けのナレッジ グラフ
コード理解において構造グラフがベクトル RAG に勝る理由。
Graphify がソース上で LLM 呼び出しを行わずに 19 の言語をローカルで解析する方法。
グラフ トポロジのみでクラスタリングします。埋め込みやベクトル ストアはありません。
CLAUDE.md ディレクティブと PreToolUse フックを段階的に説明します。
すべての /graphify およびgraphify コマンドを 1 か所にまとめました。
Sourcegraph、Code2Vec、Neo4j との正直な比較。
いいえ。Graphify は、ドキュメントと図のセマンティックな説明のみを、アシスタントですでに構成されている AI モデルに送信します。生のソース ファイルは送信しません。
Claude Code、OpenAI Codex、および OpenCode は、専用の skill-*.md マニフェストを介してすぐに使用できるようにサポートされています。シェル コマンドを呼び出すことができるアシスタントは、graphify を呼び出すことができます。
ツリーシッター解析と NetworkX 構築は、コード サイズに比例して拡張されます。約 500,000 語のコーパスでは、BFS サブグラフ クエリは約 2,000 トークンに対して、ナイーブ クエリは約 670,000 に留まり、大規模な圧縮を維持します。
はい。 Graphify は MIT ライセンスを取得しており、個人使用と商用使用の両方で無料です。
© 2026 Graphify · Safi Shamsi によって管理 · MIT ライセンスに基づいてリリース

## Original Extract

Graphify is an open-source knowledge graph skill that helps AI coding assistants understand multi-modal codebases. Graphify extracts code, docs, papers and diagrams into a queryable graph using Tree-sitter, NetworkX and Leiden clustering.

𝕏
in
f
R
Y
⧉
Graphify
Features
Graphify — Knowledge Graphs
for AI Coding Assistants
Graphify is an open-source skill that helps AI coding assistants understand multi-modal codebases by building a queryable knowledge graph from code, docs, papers and diagrams.
Graphify is a multi-modal knowledge graph builder created for AI coding assistants such as Claude Code, OpenAI Codex and OpenCode. By combining Tree-sitter static analysis with LLM-driven semantic extraction, Graphify turns an entire repository — including source code, documentation, research papers and diagrams — into an interactive graph that explains both what the code does and why it was designed that way. The project is maintained by Safi Shamsi, released under the permissive MIT license, and built on widely-trusted libraries including NetworkX and Tree-sitter.
Graphify unifies static analysis, semantic extraction and graph clustering into a single skill that any AI coding assistant can invoke.
Parses code (.py, .js, .go, .java, …), Markdown, PDFs and images. Tree-sitter extracts ASTs, call graphs and docstrings; LLMs extract concepts from prose; vision models read diagrams.
Merges all extracted nodes and edges into a NetworkX graph and applies the Leiden algorithm for semantic community detection — no vector embeddings required.
Identifies the highest-degree "god nodes" at the heart of the system and flags unexpected cross-file or cross-domain connections worth investigating.
Exports an interactive graph.html , a queryable graph.json , and a human-readable GRAPH_REPORT.md audit report.
Ships with /graphify , /graphify query , /graphify path and /graphify explain commands for Claude Code, Codex, OpenCode and more.
Strict input validation: only http/https URLs, size and timeout limits, path containment, HTML-escaped node labels — defending against SSRF, injection and XSS.
Graphify is a multi-stage pipeline. Each stage is an isolated module so contributors can extend any step independently.
Supporting modules include ingest.py for URL fetching, cache.py for semantic caching, security.py for input validation, watch.py for live updates and serve.py for MCP-protocol service.
Graphify is distributed on PyPI. The package name is graphifyy ; the CLI command remains graphify .
# Requires Python 3.10+
pip install graphifyy && graphify install
# Build a knowledge graph for any project folder
/graphify ./raw
# Outputs land in graphify-out/
graphify-out/
├── graph.html # interactive visualization
├── GRAPH_REPORT.md # core nodes, surprises, suggested questions
├── graph.json # persistent, queryable graph
└── cache/ # incremental cache
Graphify does not bundle an LLM. It uses the model API key already configured by your AI coding assistant (Claude, Codex, etc.) and only sends semantic content — never raw source code — to the upstream model.
The repository ships with reproducible corpora demonstrating Graphify on both small libraries and large mixed code-and-paper collections.
6 Python files modeling an HTTP transport layer. Result: 144 nodes, 330 edges, 6 communities . God nodes: Client , AsyncClient , Response , Request . Surprise edge: DigestAuth → Response .
3 GPT framework repos + 5 attention papers + 4 diagrams (~52 files, ~92k words). Result: 285 nodes, 340 edges, 53 communities . Average query cost ~1.7k tokens vs ~123k naive — a 71.5× reduction.
How Graphify relates to adjacent open-source projects in the code-intelligence space.
Graphify is released under the MIT License . Its core dependencies — NetworkX (BSD) and Tree-sitter (MIT) — are all permissive open-source licenses with no conflicts. The project performs no telemetry. The only outbound network call is the semantic-extraction step, which uses your own configured AI model API key; only semantic descriptions of documents are transmitted, never raw source code. URLs are restricted to http/https, downloads are size- and time-bounded, output paths are containment-checked, and node labels are HTML-escaped to prevent SSRF, Cypher injection and XSS.
Deeper guides on how Graphify builds, clusters and serves knowledge graphs to AI coding assistants.
Knowledge Graphs for AI Coding Assistants
Why structural graphs beat vector RAG for code understanding.
How Graphify parses 19 languages locally, with no LLM calls on source.
Clustering on graph topology alone — no embeddings, no vector store.
CLAUDE.md directives and the PreToolUse hook, step by step.
Every /graphify and graphify command in one place.
Honest comparison against Sourcegraph, Code2Vec and Neo4j.
No. Graphify only sends semantic descriptions of documents and diagrams to the AI model you have already configured in your assistant — never raw source files.
Claude Code, OpenAI Codex and OpenCode are supported out of the box via dedicated skill-*.md manifests. Any assistant that can call shell commands can invoke graphify .
Tree-sitter parsing and NetworkX construction scale linearly with code size. On a ~500k-word corpus, BFS subgraph queries stay around ~2k tokens versus ~670k naive — preserving compression at scale.
Yes. Graphify is MIT-licensed and free for both personal and commercial use.
© 2026 Graphify · Maintained by Safi Shamsi · Released under the MIT License
