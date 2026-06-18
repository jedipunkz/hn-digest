---
source: "https://solidlao.github.io/GenDB/"
hn_url: "https://news.ycombinator.com/item?id=48579954"
title: "GenDB – LLM-Powered Generative Query Engine"
article_title: "GenDB — LLM-Powered Generative Query Engine"
author: "matt_d"
captured_at: "2026-06-18T04:35:05Z"
capture_tool: "hn-digest"
hn_id: 48579954
score: 1
comments: 0
posted_at: "2026-06-18T02:35:48Z"
tags:
  - hacker-news
  - translated
---

# GenDB – LLM-Powered Generative Query Engine

- HN: [48579954](https://news.ycombinator.com/item?id=48579954)
- Source: [solidlao.github.io](https://solidlao.github.io/GenDB/)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T02:35:48Z

## Translation

タイトル: GenDB – LLM を利用した生成クエリ エンジン
記事のタイトル: GenDB — LLM を利用した生成クエリ エンジン

記事本文:
世代データベース
✨デモ
について
なぜ
リーダーボード
モデル
言語
ロードマップ
チーム
紙
GitHub
✨デモ
について
なぜ
リーダーボード
モデル
言語
ロードマップ
チーム
紙
GitHub
次世代の
クエリ処理
GenDB は、LLM エージェントを使用してインスタンスに最適化されたデータを生成する生成クエリ エンジンです。
特定のデータ、ワークロード、ハードウェアに合わせて調整されたクエリ実行コード。
5 つの専門化された LLM エージェントが構造化されたパイプラインを通じて連携して、
最適化されたストレージ、インデックス、スタンドアロンのネイティブ実行可能ファイルはすべて、特定のデータ、ワークロード、ハードウェアに合わせて調整されています。
ハードウェアのプロファイリング、データのサンプリング、ワークロード特性の抽出
エンコード、圧縮、インデックス、ゾーンマップを使用してレイアウトを設計します。
データとハードウェアに適応したリソースを意識した実行計画を生成します
SIMD と並列処理を使用して最適化されたネイティブ コードとして計画を実装します。
実行時プロファイリングのフィードバックを使用してコードを反復的に改良します
現在、新しいユースケースはすべて、苦痛を伴う拡張か、まったく新しいシステムのいずれかを必要としています。
オプション 1 — 既存のシステムを拡張する
PostgreSQL → PostGIS、TimescaleDB、pgvector、Citus、AGE …
各拡張機能は、ホスト システムのアーキテクチャ上の制約と闘います。
DuckDB、Umbra、ClickHouse、Milvus、Pinecone、InfluxDB、Neo4j …
それぞれに何年ものエンジニアリングと莫大な金銭的コストが必要です。
LLM を使用してクエリごとの実行コードを生成します。延長レスリングも、複数年にわたるエンジニアリングもありません。迅速なアップデートを通じて新しい技術を利用できるようになります。
インスタンスに最適化されたコードは、正確なデータ分散、結合選択性、グループ カーディナリティ、およびハードウェア特性を活用します。これに匹敵する汎用エンジンはありません。
新しい技術を統合するには、リエンジニアリングではなくプロンプトが必要です。セマンティック クエリ、GPU ネイティブ コード - すべてはプロンプト更新を通じてアクセス可能です。
クエリの 80% が 50% で繰り返されます。

クラスター。生成コストは多数の実行にわたって償却されるため、繰り返し発生する分析ワークロードの費用対効果が高くなります。
すべてのクエリにわたる合計クエリ実行時間。 GenDB バリアントは、異なる LLM バックボーン モデルを使用します。
すべてのシステムは、完全な並列処理が有効になっている同一のハードウェア上で実行されます。
LLM バックボーン モデルが異なれば、生成されるコードの品質、生成時間、コスト間のトレードオフも異なります。平均クエリ実行時間によってランク付けされます。
GenDB の実行から各 TPC-H クエリに対して最もパフォーマンスの高い C++ バイナリを選択し、
クロード コード (Opus 4.6) 分析、プロファイル、改善のための 5 回の反復 —
最初は最適化された C++ で、次に完全な Rust の書き換えです。
GenDB が標準コンパイルで生成したコード。
積極的なフラグ、Madvise チューニング、並列結合、スレッドの最適化。
rayon、memmap2、安全でない境界チェックの削除による完全な書き換え。
主な調査結果: 最適化された C++ は全体で 1.30 倍の高速化を達成し、第 18 四半期には並列結合の構築による最大の向上 (2.44 倍) が示されました。
Rust は第 6 四半期 ( get_unchecked を使用したゾーンマップ スキャン) で勝利しましたが、mmap ページ テーブルのセットアップからクエリごとに ~30ms のオーバーヘッドが発生し、短いクエリにペナルティが発生します。
Rustのmain_scanの計算時間はC++と競合しており、オーバーヘッドがアルゴリズムではなく構造的なものであることを示唆しています。
標準 GenDB ワークフローの一部としてこれらの利点を自動的に達成するために、低レベルの実装レベルの最適化を担当する専用の Code Refiner エージェントをパイプラインに導入する予定です。
GenDB は積極的に開発中です。すべてのステップは次の 3 つの原則に従います。
分析クエリ用のマルチエージェント パイプライン。 TPC-H および SEC-EDGAR で評価され、DuckDB、Umbra、ClickHouse、MonetDB、PostgreSQL を上回りました。
エージェントは過去の実行から学習し、最適化の経験を蓄積し、時間の経過とともに生成品質を向上させます。

基礎となる LLM を再トレーニングする。
CPU だけでなくコスト効率の高い GPU 分析のために、libcudf をターゲットとする CUDA および GPU アクセラレーション コードを生成します。
SQL のリレーショナル モデルを超えて、AI を活用したオペレーターを使用して、画像、音声、テキストなどのマルチモーダル データのコードを生成します。
クエリ全体で再利用可能な演算子、クエリ テンプレートの生成、従来の DBMS とのハイブリッド実行、LLM の高速化と低コスト化によるさらなるコスト削減。
Cornell Database Group で構築されました。
GenDB — コーネル大学、ジャレ・ラオ & イマヌエル・トラマー

## Original Extract

Gen DB
✨ Demo
About
Why
Leaderboard
Models
Languages
Roadmap
Team
Paper
GitHub
✨ Demo
About
Why
Leaderboard
Models
Languages
Roadmap
Team
Paper
GitHub
The Next Generation of
Query Processing
GenDB is a Generative Query Engine that uses LLM agents to generate instance-optimized
query execution code, tailored to your specific data, workloads, and hardware.
Five specialized LLM agents collaborate through a structured pipeline to generate
optimized storage, indexes, and standalone native executables — all tailored to the specific data, workload, and hardware.
Profiles hardware, samples data, extracts workload characteristics
Designs layouts with encoding, compression, indexes, and zone maps
Generates resource-aware execution plans adapted to data and hardware
Implements plans as optimized native code with SIMD and parallelism
Iteratively refines code using runtime profiling feedback
Today, every new use case demands either a painful extension or an entirely new system:
Option 1 — Extend an existing system
PostgreSQL → PostGIS, TimescaleDB, pgvector, Citus, AGE …
Each extension fights the host system’s architectural constraints .
DuckDB, Umbra, ClickHouse, Milvus, Pinecone, InfluxDB, Neo4j …
Each requires years of engineering and huge monetary costs .
Use LLMs to generate per-query execution code . No extension wrestling, no multi-year engineering. New techniques become reachable through prompt updates.
Instance-optimized code exploits exact data distributions, join selectivities, group cardinalities, and hardware characteristics. No general-purpose engine can match this.
Integrating new techniques requires prompting, not re-engineering. Semantic queries , GPU-native code — all reachable through prompt updates.
80% of queries repeat in 50% of clusters . Generation cost is amortized over many executions, making it cost-effective for recurring analytical workloads.
Total query execution time across all queries. GenDB variants use different LLM backbone models.
All systems run on identical hardware with full parallelism enabled.
Different LLM backbone models offer different trade-offs between generated code quality, generation time, and cost. Ranked by average query execution time.
We select the best-performing C++ binary for each TPC-H query from a GenDB run, then give
Claude Code (Opus 4.6) 5 iterations to analyze, profile, and improve —
first for optimized C++, then for a full Rust rewrite.
GenDB-generated code with standard compilation.
Aggressive flags, madvise tuning, parallelized joins, thread optimization.
Full rewrite with rayon, memmap2, unsafe bounds-check elimination.
Key findings: Optimized C++ achieves a 1.30x overall speedup, with Q18 showing the largest gain (2.44x) from parallelized join building.
Rust wins on Q6 (zone-map scan with get_unchecked ) but carries ~30ms per-query overhead from mmap page table setup, penalizing short queries.
The Rust main_scan compute times are competitive with C++, suggesting the overhead is structural rather than algorithmic.
We plan to introduce a dedicated Code Refiner agent to the pipeline, responsible for low-level, implementation-level optimizations — to automatically achieve these gains as part of the standard GenDB workflow.
GenDB is under active development. Every step follows three principles:
Multi-agent pipeline for analytical queries. Evaluated on TPC-H and SEC-EDGAR, outperforming DuckDB, Umbra, ClickHouse, MonetDB, and PostgreSQL.
Agents learn from past runs, accumulate optimization experience, and improve generation quality over time — without retraining the underlying LLMs.
Generate CUDA and GPU-accelerated code targeting libcudf for cost-efficient GPU analytics, not just CPU.
Generate code for multimodal data — images, audio, text — with AI-powered operators, moving beyond SQL’s relational model.
Reusable operators across queries, query template generation, hybrid execution with traditional DBMS, and further cost reduction as LLMs become faster and cheaper.
Built at Cornell Database Group .
GenDB — Jiale Lao & Immanuel Trummer, Cornell University
