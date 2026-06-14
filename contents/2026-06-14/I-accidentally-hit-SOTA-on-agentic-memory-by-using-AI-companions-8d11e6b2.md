---
source: "https://graph.coder.company/"
hn_url: "https://news.ycombinator.com/item?id=48524870"
title: "I accidentally hit SOTA on agentic memory by using AI companions"
article_title: "graphCTX — local-first memory for coding agents"
author: "vignesh_146"
captured_at: "2026-06-14T07:42:25Z"
capture_tool: "hn-digest"
hn_id: 48524870
score: 2
comments: 1
posted_at: "2026-06-14T07:03:45Z"
tags:
  - hacker-news
  - translated
---

# I accidentally hit SOTA on agentic memory by using AI companions

- HN: [48524870](https://news.ycombinator.com/item?id=48524870)
- Source: [graph.coder.company](https://graph.coder.company/)
- Score: 2
- Comments: 1
- Posted: 2026-06-14T07:03:45Z

## Translation

タイトル: AI コンパニオンを使用して誤ってエージェント メモリに SOTA をヒットしてしまいました
記事のタイトル:graphCTX — コーディングエージェント用のローカルファーストメモリ
説明:graphCTX は、アカウントや API キーを使用したり、ホストされたサービスにコードを送信したりすることなく、AI コーディング エージェントをリポジトリの知識に基づいた状態に保つローカルファーストのコンテキストおよびメモリ層です。

記事本文:
chartCTX — コーディング エージェント用のローカル ファースト メモリgraphCTX 仕組み ベンチマーク インストール
GitHub
インストール
AI コーディング エージェントのコンテキスト層とメモリ層
最高のものになるように作られています
AI コーディング エージェントのコンテキスト層。
graphCTX は、コマンド、規約、
決断、そして苦労して勝ち取った修正。開発者が再説明に費やす時間が短縮される
高速、プライベート、
そして信頼できる。
ユーザー: デプロイ コマンドとは何ですか?
エージェント: リポジトリには見つかりません —
あなたは私に言う必要があるでしょう。
ユーザー: デプロイ コマンドとは何ですか?
エージェント: ./scripts/ship.sh --canary --wait
メモリ:9f3a2c
同じエージェント、同じプロンプト。 graphCTX は開発者にリポジトリ メモリを提供します。そうでない場合は繰り返します。
システムは意図的に狭い範囲に設定されています。信頼できるコーディング事実をキャプチャし、
これらはリポジトリが変更されても有効であり、エージェントが実行できるコンテキストのみを返します。
使用します。
graphCTX はパッケージ スクリプト、ロックファイル、CI、editorconfig、AGENTS.md、およびセッション エピソードを読み取るため、メモリ ベースはモデルの推測ではなくリポジトリの証拠から始まります。
ファクトはコミットとブランチに対して有効です。コードが変更されると、メモリは実時間のタイムスタンプに基づいて移動するのではなく、DAG とともに移動します。
永続的な知識のみを促進する
セッションの詳細は、証拠ゲートの後にのみワークスペースまたはユーザー メモリになることができます。秘密や信頼性の低い散文は、永続的な文脈から外されます。
最小の有用なセットを選択する
関連性ゲートはトピック ドリフト、エンティティ、およびファイル スコープをスコアリングするため、エージェントはノイズの多いメモリ ダンプの代わりに特定のコンテキストを取得します。
すべてのリコールに出所を添付する
返されるメモリはコンパクトで、予算が設定されており、ソースの出所がタグ付けされているため、開発者はエージェントが使用するものを信頼して監査することが容易になります。
スーパーメモリに対してベンチマークを実施。
同じコーディングファクトセット、同じクエリ。 graphCTX はローカルで実行され、応答します。
~1ms;ライブスーパーミー

メモリ検索のラウンドトリップは約 494ms (p50) と測定されたため、
開発中のリコールは迅速かつ予測可能です。
再現可能:graphctx Compare --deep 。
ワークスペースの拡大に応じて、プロンプトごとの取得は p50/p95 になります。インデックス付きルックアップ
さらに、有界セマンティック再ランクによりホット パスが維持されます。
~1ms — 5,000 ファクト
モノリポジトリは空のリポジトリと同じくらい高速に取得します。
14 のコーディング タスクにわたる圧縮後の解決率。グラフCTXが復元されました
実行ごとに必要なリポジトリ ファクト。
//graphctx eval run --arms A、B、C
// スコープ: ローカル レイテンシ + コーディング ファクトの直接取得のコストを比較します。
開発者のワークフロー向け。スーパーメモリは一般記憶/会話記憶をターゲットとしています
クラウド コネクタ、クロスドキュメント推論、ニューラル埋め込みを使用して、
graphCTX は試行しません。
30 秒以内にリポジトリ メモリの使用を開始します
アカウントがありません。 APIキーがありません。クラウドの設定はありません。 CLIをインストールし、接続します
エージェントに追加し、すべてのセッションにリポジトリ対応メモリを与えます。
CLIをインストールします(ノード/バンを検出します)
エージェントを接続します (クロード · カーソル · オープンコード · ジェネリック)
グラフCTX
AI コーディング エージェントのコンテキストとメモリの管理。地元第一主義
SaaS ではなく開発者ツール。

## Original Extract

graphCTX is a local-first context and memory layer that keeps AI coding agents grounded in repo knowledge, without accounts, API keys, or sending code to a hosted service.

graphCTX — local-first memory for coding agents graphCTX How it works Benchmarks Install
GitHub
Install
context and memory layer for AI coding agents
Built to be the best
context layer for AI coding agents.
graphCTX keeps repo knowledge close to the work: commands, conventions,
decisions, and hard-won fixes. Developers spend less time re-explaining
context and more time shipping, with local memory that is fast, private,
and reliable.
user: what's the deploy command?
agent: I don't see one in the repo —
you'll need to tell me.
user: what's the deploy command?
agent: ./scripts/ship.sh --canary --wait
mem:9f3a2c
Same agent, same prompt. graphCTX gives it the repo memory developers otherwise repeat.
The system is deliberately narrow: capture reliable coding facts, keep
them valid as the repo changes, and return only the context the agent can
use.
graphCTX reads package scripts, lockfiles, CI, editorconfig, AGENTS.md, and session episodes so the memory base starts from repo evidence, not model guesses.
Facts are valid against commits and branches. When code changes, memory moves with the DAG instead of drifting on wall-clock timestamps.
Promote only durable knowledge
Session details can become workspace or user memory only after evidence gates. Secrets and low-trust prose stay out of durable context.
Choose the smallest useful set
A relevance gate scores topic drift, entities, and file scope so the agent gets specific context instead of a noisy memory dump.
Attach provenance to every recall
Returned memory is compact, budgeted, and tagged with source provenance, making it easier for developers to trust and audit what the agent uses.
Benchmarked against Supermemory.
Same coding-fact set, same queries. graphCTX runs locally and answers in
~1ms; a live Supermemory search round-trip measured ~494ms (p50), so
recall stays fast and predictable during development.
Reproducible: graphctx compare --deep .
Per-prompt retrieval p50/p95 as the workspace grows. Indexed lookup
plus a bounded semantic re-rank keeps the hot path at
~1ms — a 5,000-fact
monorepo retrieves as fast as an empty one.
Post-compaction solve rate across 14 coding tasks. graphCTX restored
the needed repo fact in every run.
// graphctx eval run --arms A,B,C
// scope: this compares local latency + cost on direct coding-fact retrieval
for developer workflows. Supermemory targets general/conversational memory
with cloud connectors, cross-document reasoning, and neural embeddings that
graphCTX doesn't attempt.
Start using repo memory in 30 seconds
No account. No API key. No cloud setup. Install the CLI, connect your
agent, and give every session repo-aware memory.
install the CLI (detects Node / Bun)
wire your agent (claude · cursor · opencode · generic)
graphCTX
Context and memory management for AI coding agents. Local-first
developer tooling, not a SaaS.
