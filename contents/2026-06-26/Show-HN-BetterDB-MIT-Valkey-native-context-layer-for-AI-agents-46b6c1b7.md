---
source: "https://github.com/BetterDB-inc/monitor/tree/master/packages"
hn_url: "https://news.ycombinator.com/item?id=48687650"
title: "Show HN: BetterDB, MIT Valkey-native context layer for AI agents"
article_title: "monitor/packages at master · BetterDB-inc/monitor · GitHub"
author: "kaliades"
captured_at: "2026-06-26T15:44:58Z"
capture_tool: "hn-digest"
hn_id: 48687650
score: 3
comments: 0
posted_at: "2026-06-26T15:16:24Z"
tags:
  - hacker-news
  - translated
---

# Show HN: BetterDB, MIT Valkey-native context layer for AI agents

- HN: [48687650](https://news.ycombinator.com/item?id=48687650)
- Source: [github.com](https://github.com/BetterDB-inc/monitor/tree/master/packages)
- Score: 3
- Comments: 0
- Posted: 2026-06-26T15:16:24Z

## Translation

タイトル: Show HN: BetterDB、AI エージェント用の MIT Valkey ネイティブ コンテキスト レイヤー
記事のタイトル: マスターのモニター/パッケージ · BetterDB-inc/モニター · GitHub
説明: Valkey と Redis のリアルタイム監視、スローログ分析、監査証跡 - マスターでの監視/パッケージ · BetterDB-inc/monitor
HN テキスト: 本日、私たちは、BetterDB のパッケージ (エージェント メモリ、セマンティック + 多層キャッシュ、型付き取得) の一部として、AI エージェント用のオープンな Valkey ネイティブ コンテキスト レイヤーをリリースしました。これは、ベンダー ロックインがなく、どこにいても Valkey インスタンス上で実行されます。本日から Valkey インスタンスのプロビジョニングも開始しました。パッケージは npm と PyPi で出荷されます。開発の理由: BetterDB はもともと、Valkey、Redis、および RESP 互換データベースの監視および可観測性プラットフォームとして始まりました。これは依然として製品の中核ですが、これを構築する過程で、Valkey の最も急速に成長している用途の 1 つが AI、つまりエージェントと RAG の背後にあるベクター ストアとキャッシュであることに気づきました。そこで、約 1 か月半前に、そのための MIT セマンティック ライブラリとエージェント キャッシュ ライブラリを公開しました。エージェント キャッシュ ライブラリはモジュールを必要とせず、バニラ Valkey 上で実行できます。今日、これをエージェントのメモリに拡張します。私たちは可観測性から始めたため、どこでも実行できます。すべてのキャッシュ、メモリ、取得で OTel と Prometheus が生成され、さらに独自の監視サーバーおよび mcp サーバーとうまく統合され、エージェントに公開されます。実際に発送したもの：
- エージェント メモリ: 短期層 (セッション/LLM/ツール、完全一致) とセマンティックな長期層
- カテゴリごとのしきい値と信頼帯を備えた valkey-search 上のセマンティック キャッシュ
- valkey-search による型付き検索
- 自己調整ループと OTel/Prometheus の可観測性 (詳細は後述) まだ行われていないこと: 1 つのコマンドの self-hostin 用の Helm チャート

g、および詳細なベンチマークの説明。どちらも来週です。セルフチューニングキャッシュ。キャッシュは類似性スコアをログに記録し、別のサービスが分布を読み取り、しきい値/TTL の変更を提案します (コストで重み付けされた推論付き)。人間が変更を承認すると、再起動することなく、実行中のキャッシュが 1 秒以内にその変更を取得します。エージェントはライブ キャッシュの状態を読み取り、MCP 経由で変更を提案できます。このループを閉じる別のキャッシュ ライブラリは見つかりませんでした。ほとんどの場合、手動で調整された静的なしきい値が使用されます。これは、セマンティック キャッシュの文書化された障害モードです。
運用レベルでの可観測性。リクエストレベルのLLMトレースだけでなく、各キャッシュ/メモリ/取得操作に関するOTelスパンとPrometheusメトリクス。そのため、推測ではなく、ルックアップごとの類似性分布と、しきい値が間違っているかどうかを実際に確認できます。ベンチマークでは、数値はフレックスではありません。メモリ層で LongMemEval を実行しました。ハーネスを構築する過程で、スコアを抑制する複数の要因があることがわかりました。読者へのプロンプトを微調整した場合でも (読者は文字通りの抜粋からのみ回答するように指示されているため、質問タイプ全体を控えました)、一致する gpt-4o 構成では 5 ポイント以上の改善がありました。来週は QA に積極的に取り組んでいきます。すべてを再実行してから、包括的な記事を公開します。検索再現率は 98.4% でほぼ完璧であるため、ギャップは検索ではなく読み手/推論側にあります。 Valkey のパフォーマンスが優れているため、最も優れているのは ofc レイテンシーです。私が本当にフィードバックを求めているのは、Valkey ネイティブの賭けはあなたにとって意味がありますか、それともコンテキスト層がストレージに依存しないほうがよいでしょうか?また、本番環境でエージェントを実行している場合、自動化された自己チューニングの推奨事項を信頼しますか、それとも手動のままにすることを好みますか?コストとレイテンシのどちらが解決すべき大きな問題でしょうか? chat.b でスクリプトなしの公開デモを行っています。

ところで、これらのライブラリが実際に動作しているのを見たい人がいれば、etterdb.com にアクセスしてください。

記事本文:
マスターでのモニター/パッケージ · BetterDB-inc/monitor · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
BetterDB株式会社
/
モニター
公共
通知
通知設定を変更するにはサインインする必要があります
追加ナビ

ゲーションオプション
コード
その他のオプション ディレクトリアクション
歴史 歴史マスター ブレッドクラム
コピーパスのトップフォルダーとファイル
.. エージェント-キャッシュ-py エージェント-キャッシュ-py エージェント-キャッシュ エージェント-キャッシュ エージェント-メモリ-py エージェント-メモリ-py エージェント-メモリ エージェント-メモリ エージェント エージェント キャッシュ-ベンチマーク-ts キャッシュ-ベンチマーク-ts キャッシュ-ベンチマーク キャッシュ-ベンチマーク cli cli mcp mcp 取得-py 取得-py 取得 取得 semantic-cache-py semantic-cache-py semantic-cache semantic-cache 共有 共有 valkey-search-kit-py valkey-search-kit-py valkey-search-kit valkey-search-kit README.md README.md すべてのファイルを表示 README.md
このディレクトリには、BetterDB エコシステムを構成するオープンソース パッケージが含まれています。すべてのパッケージは、Valkey (および Redis) と連携し、可観測性と自己調整のために BetterDB Monitor と統合するように設計されています。
パッケージ
言語
レジストリ
説明
セマンティックキャッシュ
TypeScript
npm: @betterdb/semantic-cache
Valkey ベクトル検索に基づく LLM アプリケーションのセマンティック キャッシュ。再ランキング、LLM-as-judge、埋め込みキャッシング、コスト追跡、および OpenTelemetry/Prometheus インストルメンテーションを備えた埋め込みベースの類似性マッチング。
セマンティックキャッシュpy
パイソン
PyPI: betterdb-semantic-cache
@betterdb/semantic-cache に相当する Python 。同じ Valkey データ形式 - TypeScript アプリと Python アプリは同じキャッシュ インデックスを共有できます。
エージェントキャッシュ
TypeScript
npm: @betterdb/エージェントキャッシュ
AI エージェントのワークロード用の多層完全一致キャッシュ。ツールごとの TTL ポリシーとコスト追跡を使用して、LLM 応答、ツール結果、セッション状態をキャッシュします。
エージェントキャッシュ-py
パイソン
PyPI: betterdb-agent-cache
@betterdb/agent-cache に相当する Python 。
検索と記憶
パッケージ
言語
レジストリ
説明
ヴァルキーサーチキット
TypeScript
npm: @betterdb/valkey-search-kit
Valkey Search (FT.*) の共有低レベル ヘルパー: ベクトル バイト エンコーディング、FT.SEA

RCH / FT.INFO 応答の解析、TAG フィルターのエスケープ、およびエラー分類。実行時の依存関係はありません。検索およびエージェント メモリ パッケージの基盤。
valkey-search-kit-py
パイソン
PyPI: betterdb-valkey-search-kit
@betterdb/valkey-search-kit に相当する Python 。
検索
TypeScript
npm: @betterdb/retrieval
Valkey Search を介した開発者向けの取得 SDK。型付きインデックス スキーマ、冪等インデックス ライフサイクル、更新/削除、ベクトル + フィルター + ハイブリッド クエリ。 @betterdb/valkey-search-kit に基づいて構築されています。
検索-py
パイソン
PyPI: betterdb-取得
@betterdb/retrieval に相当する Python 。同じ Valkey データ形式 - TypeScript アプリと Python アプリは同じインデックスを共有できます。
エージェントメモリ
TypeScript
npm: @betterdb/agent-memory
Valkey Search と、 @betterdb/agent-cache の短期キャッシュ層によってサポートされる AI エージェントの長期セマンティック メモリ。 remember() / remember() には、最新性と重要性、容量の削除、および統合が融合された類似性があります。
エージェントメモリ-py
パイソン
PyPI: betterdb-agent-memory
@betterdb/agent-memory に相当する Python 。
ツール
パッケージ
言語
レジストリ
説明
クリ
TypeScript
npm: @betterdb/monitor
Valkey/Redis データベースを監視および監視するための CLI。 BetterDB Monitor に接続して、リアルタイムのメトリクス、スローログ分析、クライアント検査を行います。
mcp
TypeScript
npm: @betterdb/mcp
Valkey/Redis 可観測性のための MCP サーバー。 Claude Code およびその他の MCP 互換 AI アシスタントと統合して、データベースの健全性、レイテンシ、メモリ、キャッシュ メトリクスをクエリします。
エージェント
TypeScript
npm: @betterdb/agent
Valkey/Redis インスタンスを BetterDB Monitor Cloud に接続するリモート監視エージェント。サイドカーまたはスタンドアロン コンテナとして実行します。
ベンチマーク
パッケージ
言語
レジストリ
説明
キャッシュベンチマーク
パイソン
— (非公開)
ベンチ用リプレイハーネス

ラベル付きクエリ ペア データセットに対するセマンティック キャッシュ実装 (BetterDB、RedisVL、GPTCache) をアーク化します。 STSb、SICK、SemBenchmarkLmArena、および PAWS-Wiki をサポートします。
内部
パッケージ
言語
説明
共有
TypeScript
API、CLI、MCP サーバー、およびキャッシュ提案システム全体で使用される共有タイプ、定数、およびユーティリティ。直接消費することを目的としていません。
フッター
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Real-time monitoring, slowlog analysis, and audit trails for Valkey and Redis - monitor/packages at master · BetterDB-inc/monitor

Today we released an open, Valkey-native context layer for AI agents as part of our packages at BetterDB (agent memory, semantic + multi-tier caching, typed retrieval) that run on a Valkey instance no matter where it is - no vendor lock-in. We even started provisioning Valkey instances starting today. Packages are shipped on npm and PyPi. Why we made it: BetterDB originally started as a monitoring and observability platform for Valkey, Redis and any RESP compatible db. This is still the core of the product, but in the process of building this, we kept seeing that one of the fastest-growing uses of Valkey was AI - vector store and cache behind agents and RAG. So about a month and half ago we published MIT semantic and agent cache libraries for that, with the agent cache library not even requiring any modules and being able to run on vanilla Valkey. Today we are extending this to agent memory. Because we started with observability, it also runs everywhere - every cache, memory and retrieval emits OTel and Prometheus, plus it integrates well with our own monitoring and mcp server, exposing it to the agent. What we actually shipped:
- agent memory: short-term tiers (session/LLM/tool, exact-match) plus a semantic long-term layer
- semantic caching over valkey-search, with per-category thresholds and confidence bands
- typed retrieval over valkey-search
- a self-tuning loop and OTel/Prometheus observability (more below) What's not done yet: a Helm chart for one-command self-hosting, and a detailed benchmark writeup. Both next week. Self-tuning cache. The cache logs similarity scores, and a separate service reads the distribution and proposes threshold/TTL changes (with reasoning, weighted by cost). A human approves the change, and the running cache picks it up in under a second with no restart. An agent can read the live cache state and propose changes over MCP. I haven't found another cache library that closes this loop; most use a static, hand-tuned threshold, which is the documented failure mode for semantic caches.
Observability at the operation level. OTel spans and Prometheus metrics on each cache/memory/retrieval operation, not just request-level LLM tracing. So you can actually see per-lookup similarity distributions and whether your threshold is wrong, rather than guessing. On benchmarks, the number isn't a flex: I ran LongMemEval on the memory layer. In the process of building our harness I found multiple things surpressing the scores. Even tweaking the prompt to the reader (the reader was told to answer only from literal excerpts so it abstained on whole question types), on a matched gpt-4o config the improvement was over 5 points. We'll be actively working on QA next week. Re-run everything and then publish a comprehensive write up. Retrieval recall is near-perfect at 98.4%, so the gap is reader/reasoning-side, not retrieval. The best part is ofc latency as Valkey's performance is great. What I'd genuinely like feedback on: does the Valkey-native bet make sense to you, or would you rather a context layer be storage-agnostic? And for those running agents in prod, would you trust automated self tuning recommendations, or prefer to keep it manual? Is cost or latency a bigger issue to be solved? We have a public unscripted demo at chat.betterdb.com btw, if anyone wants to see these libraries in action.

monitor/packages at master · BetterDB-inc/monitor · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
BetterDB-inc
/
monitor
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
More options Directory actions
History History master Breadcrumbs
Copy path Top Folders and files
.. agent-cache-py agent-cache-py agent-cache agent-cache agent-memory-py agent-memory-py agent-memory agent-memory agent agent cache-benchmark-ts cache-benchmark-ts cache-benchmark cache-benchmark cli cli mcp mcp retrieval-py retrieval-py retrieval retrieval semantic-cache-py semantic-cache-py semantic-cache semantic-cache shared shared valkey-search-kit-py valkey-search-kit-py valkey-search-kit valkey-search-kit README.md README.md View all files README.md
This directory contains the open-source packages that make up the BetterDB ecosystem. All packages are designed to work with Valkey (and Redis) and integrate with BetterDB Monitor for observability and self-tuning.
Package
Language
Registry
Description
semantic-cache
TypeScript
npm: @betterdb/semantic-cache
Semantic cache for LLM applications backed by Valkey vector search. Embeddings-based similarity matching with reranking, LLM-as-judge, embedding caching, cost tracking, and OpenTelemetry/Prometheus instrumentation.
semantic-cache-py
Python
PyPI: betterdb-semantic-cache
Python counterpart to @betterdb/semantic-cache . Same Valkey data format — a TypeScript app and a Python app can share the same cache index.
agent-cache
TypeScript
npm: @betterdb/agent-cache
Multi-tier exact-match cache for AI agent workloads. Caches LLM responses, tool results, and session state with per-tool TTL policies and cost tracking.
agent-cache-py
Python
PyPI: betterdb-agent-cache
Python counterpart to @betterdb/agent-cache .
Retrieval & memory
Package
Language
Registry
Description
valkey-search-kit
TypeScript
npm: @betterdb/valkey-search-kit
Shared low-level helpers for Valkey Search ( FT.* ): vector byte encoding, FT.SEARCH / FT.INFO reply parsing, TAG filter escaping, and error classification. No runtime dependencies. Foundation for the retrieval and agent-memory packages.
valkey-search-kit-py
Python
PyPI: betterdb-valkey-search-kit
Python counterpart to @betterdb/valkey-search-kit .
retrieval
TypeScript
npm: @betterdb/retrieval
Developer-facing retrieval SDK over Valkey Search. Typed index schema, idempotent index lifecycle, upsert/delete, and vector + filtered + hybrid query. Built on @betterdb/valkey-search-kit .
retrieval-py
Python
PyPI: betterdb-retrieval
Python counterpart to @betterdb/retrieval . Same Valkey data format — a TypeScript app and a Python app can share the same index.
agent-memory
TypeScript
npm: @betterdb/agent-memory
Long-term semantic memory for AI agents backed by Valkey Search, plus the short-term cache tiers from @betterdb/agent-cache . remember() / recall() with similarity blended with recency and importance, capacity eviction, and consolidation.
agent-memory-py
Python
PyPI: betterdb-agent-memory
Python counterpart to @betterdb/agent-memory .
Tools
Package
Language
Registry
Description
cli
TypeScript
npm: @betterdb/monitor
CLI for monitoring and observing Valkey/Redis databases. Connects to BetterDB Monitor for real-time metrics, slowlog analysis, and client inspection.
mcp
TypeScript
npm: @betterdb/mcp
MCP server for Valkey/Redis observability. Integrates with Claude Code and other MCP-compatible AI assistants to query database health, latency, memory, and cache metrics.
agent
TypeScript
npm: @betterdb/agent
Remote monitoring agent that connects a Valkey/Redis instance to BetterDB Monitor Cloud. Runs as a sidecar or standalone container.
Benchmarking
Package
Language
Registry
Description
cache-benchmark
Python
— (not published)
Replay harness for benchmarking semantic cache implementations (BetterDB, RedisVL, GPTCache) against labeled query-pair datasets. Supports STSb, SICK, SemBenchmarkLmArena, and PAWS-Wiki.
Internal
Package
Language
Description
shared
TypeScript
Shared types, constants, and utilities used across the API, CLI, MCP server, and cache proposal system. Not intended for direct consumption.
Footer
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
