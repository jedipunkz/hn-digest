---
source: "https://useokf.com/"
hn_url: "https://news.ycombinator.com/item?id=49325279"
title: "The CLI your AI agent drives to manage your knowledge graph"
article_title: "okf: the CLI your AI agent drives to manage your knowledge graph"
author: "jhgaylor"
captured_at: "2026-08-17T00:39:05Z"
capture_tool: "hn-digest"
hn_id: 49325279
score: 1
comments: 0
posted_at: "2026-08-17T00:32:08Z"
tags:
  - hacker-news
  - translated
---

# The CLI your AI agent drives to manage your knowledge graph

- HN: [49325279](https://news.ycombinator.com/item?id=49325279)
- Source: [useokf.com](https://useokf.com/)
- Score: 1
- Comments: 0
- Posted: 2026-08-17T00:32:08Z

## Translation

タイトル: AI エージェントがナレッジ グラフを管理するために使用する CLI
記事のタイトル: okf: AI エージェントがナレッジ グラフを管理するために使用する CLI
説明: Open Knowledge Format (OKF) 用の Go CLI ツールキット: エージェント優先、JSON ネイティブ、ベンダー中立。 Google の代替手段

記事本文:
okf: AI エージェントがナレッジ グラフを管理するために使用する CLI
🕸️ ">
わかりました。
新機能
なぜ
コマンド
エージェントファースト
インストール
GitHub
始めましょう
新しい v0.2.0 がリリースされました: 信頼層、来歴、証明された計算 →
AI エージェントがナレッジ グラフを管理するために使用する CLI
okf は、エージェント優先、JSON ネイティブ、ベンダー中立のオープン ナレッジ フォーマット用の Go CLI ツールキットです。 Google の Python/Gemini ロックされたリファレンス実装に代わる単一のバイナリ。
信頼は今や第一級の分野です
okf v0.2.0 は、新しい OKF v0.2 仕様をエンドツーエンドで実装します。来歴、信頼、ライフサイクル、証明された計算はすべて解析、検証され、エージェントが分岐できる JSON として表示されます。 Google のナレッジ カタログ リポジトリ内のすべての参照バンドルに対してエラーなしで検証されました。
あらゆるプロバイダー上のあらゆる AI エージェント向けのベンダー中立的な OKF ツール
Google のリファレンス OKF 実装は Python + Gemini + BigQuery であり、ベンダーが Google のクラウドにロックされています。 okf はベンダー中立の代替手段です。どこでも動作し、JSON をネイティブに話し、Gemini だけでなくあらゆる AI エージェントによって駆動されるように設計された単一の Go バイナリです。
11 個のコマンド、1 つのバイナリ、ランタイム依存関係なし
テストファースト (88 テスト)、Go stdlib のみ、Apache 2.0 で構築。すべてのコマンドは、デフォルトで構造化された JSON を標準出力に出力します。 --json フラグも画面スクレイピングもありません。
すべてのコマンドの完全な JSON マニフェスト (名前、説明、フラグ、引数、標準出力形式、終了コード) を出力します。 1 回の通話で、エージェントは CLI サーフェス全体を把握できます。これがお堀です。
バージョン情報を JSON として出力します。エージェントは、残りの CLI を実行する前に互換性をチェックできます。
標準のサブディレクトリ (テーブル、データセット、プレイブック)、ルートのindex.md、および .gitignore を含む新しい空の OKF バンドルをスキャフォールディングします。 AI 主導のドキュメント パイプラインの開始点。
生成します

OKF 仕様 §8 に従って段階的に公開するために、index.md ファイルをすべてのディレクトリに配置し、バンドルの okf_version 宣言を保持します。エージェントはバンドル全体をロードするのではなく、レベルごとに移動します。
すべての概念を OKF v0.2 に対してチェックします: 必要なフロントマター、クロスリンク、来歴、信頼、ライフサイクル ファミリ、証明済み計算コントラクト、レガシー v0.1 構造。エラーが見つかった場合は 1 を終了します。 CI 品質ゲート。
validate と同じチェックですが、警告のみが生成されます。エラーは抑制されるため、警告があっても 0 で終了します。これを使用して、ビルドを失敗させることなく、不足している推奨フィールドにフラグを立てます。
すべてのコンセプト ドキュメントを ID、タイプ、タイトル、ライフサイクル ステータス、および信頼層とともに JSON としてリストします。インベントリのクエリ: このバンドルには何が含まれており、どの程度信頼すべきですか?
--tag 、 --type 、または --text によって概念をフィルタリングします。古い概念を見つけたり、カテゴリでフィルターしたり、 auth タグが付いたものをすべて構造化された JSON として見つけたりします。
単一の概念の完全なコンテンツ (フロントマター、マークダウン本文、信頼層、失効性、来歴ソース、存在する場合は認証済み計算コントラクト) をすべて JSON として返します。深読みコマンド。
マークダウン リンクと v0.2 派生エッジ (ソース、計算、実行者、認証者) から有向クロスリンク グラフを構築し、ノード、エッジ、密度、および要約統計を出力します。孤立した概念を見つけて構造を検証します。
指定された概念にリンクするすべての概念をリストします: 逆リンク検索。 「これに依存するのは誰ですか?」と答えてください。バンドル全体を grep せずに。
4 つのプリミティブ。曖昧さゼロ。
「エージェントファースト」とは、CLI が内部で LLM を呼び出すのではなく、外部 AI が okf スキーマ を介して CLI を検出して駆動できることを意味します。エージェントが必要とするものはすべてバイナリ自体に組み込まれています。
v0.2.0 がリリースされ、連署され、SBOM が含まれています。 1 つのバイナリで、ランタイムの依存関係はありません。

次に、バンドルをスキャフォールディングし、概念を追加して、以下を検証します。
実際に操作できる CLI をエージェントに提供します
スキーマ検出可能。 JSON ネイティブ。ベンダー中立。 One Go バイナリ。

## Original Extract

A Go CLI toolkit for the Open Knowledge Format (OKF): agentic-first, JSON-native, vendor-neutral. An alternative to Google

okf: the CLI your AI agent drives to manage your knowledge graph
🕸️ ">
okf .
What's New
Why
Commands
Agentic First
Install
GitHub
Get Started
New v0.2.0 just shipped: trust tiers, provenance, attested computations →
The CLI your AI agent drives to manage your knowledge graph
okf is a Go CLI toolkit for the Open Knowledge Format: agentic-first, JSON-native, vendor-neutral. A single binary alternative to Google's Python/Gemini-locked reference implementation.
Trust is now a first-class field
okf v0.2.0 implements the new OKF v0.2 spec end to end: provenance, trust, lifecycle, and attested computations, all parsed, validated, and surfaced as JSON your agent can branch on. Verified against every reference bundle in Google's knowledge-catalog repository with zero errors.
Vendor-neutral OKF tooling for any AI agent on any provider
Google's reference OKF implementation is Python + Gemini + BigQuery, vendor-locked to Google's cloud. okf is the vendor-neutral alternative: a single Go binary that works anywhere, speaks JSON natively, and is designed to be driven by any AI agent, not just Gemini.
11 commands, one binary, zero runtime dependencies
Built test-first (88 tests), Go stdlib-only, Apache 2.0. Every command outputs structured JSON on stdout by default. No --json flag, no screen-scraping.
Emits a complete JSON manifest of every command: name, description, flags, args, stdout format, exit codes. One call and an agent knows the full CLI surface. This is the moat.
Prints version info as JSON. Agents can check compatibility before driving the rest of the CLI.
Scaffolds a new empty OKF bundle with standard subdirectories (tables, datasets, playbooks), a root index.md , and a .gitignore . The starting point for an AI-driven documentation pipeline.
Generates index.md files into every directory for progressive disclosure per OKF spec §8, preserving the bundle's okf_version declaration. Agents navigate level by level instead of loading the entire bundle.
Checks every concept against OKF v0.2: required frontmatter, cross-links, the provenance, trust, and lifecycle families, attested-computation contracts, and legacy v0.1 constructs. Exits 1 if any errors are found. The CI quality gate.
Same checks as validate but only emits warnings. Errors are suppressed, so it exits 0 even with warnings. Use it to flag missing recommended fields without failing a build.
Lists every concept document with its ID, type, title, lifecycle status, and trust tier as JSON. The inventory query: what's in this bundle, and how much should I trust it?
Filters concepts by --tag , --type , or --text . Find stale concepts, filter by category, or locate everything tagged auth , all as structured JSON.
Returns a single concept's full content: frontmatter, markdown body, trust tier, staleness, provenance sources, and the attested-computation contract when present, all as JSON. The deep-read command.
Builds the directed cross-link graph from markdown links plus v0.2 derivation edges (sources, computations, executors, attesters) and prints nodes, edges, density, and summary statistics. Find orphan concepts and verify structure.
Lists every concept that links to a given concept: reverse-link lookup. Answer "who depends on this?" without grepping the entire bundle.
Four primitives. Zero ambiguity.
"Agentic first" means an external AI can discover and drive the CLI via okf schema , not that the CLI calls an LLM internally. Everything an agent needs is built into the binary itself.
v0.2.0 released, cosign-signed, SBOM-included. One binary, no runtime dependencies.
Then scaffold a bundle, add concepts, and validate:
Give your agent a CLI it can actually drive
Schema-discoverable. JSON-native. Vendor-neutral. One Go binary.
