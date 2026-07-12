---
source: "https://runewardd.github.io/runeward/"
hn_url: "https://news.ycombinator.com/item?id=48879713"
title: "Show HN: Runeward: Sandboxing AI agents with policy gates"
article_title: "runeward"
author: "tha_infra_guy"
captured_at: "2026-07-12T09:52:36Z"
capture_tool: "hn-digest"
hn_id: 48879713
score: 1
comments: 0
posted_at: "2026-07-12T09:35:29Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Runeward: Sandboxing AI agents with policy gates

- HN: [48879713](https://news.ycombinator.com/item?id=48879713)
- Source: [runewardd.github.io](https://runewardd.github.io/runeward/)
- Score: 1
- Comments: 0
- Posted: 2026-07-12T09:35:29Z

## Translation

タイトル: HN の表示: Runeward: ポリシー ゲートによる AI エージェントのサンドボックス化
記事タイトル: ルーンワード
説明: AI エージェントの管理された実行セル

記事本文:
ルーンワード
コンテンツにスキップ
ルーンワード
ホーム
検索を初期化しています
ルーンワード/ルーンワード
ルーンワード
ルーンワード/ルーンワード
ホーム
ホーム
目次
インストール
AI エージェントの管理された実行セル。
宣言型プロファイルは、分離されたサンドボックス (Docker または Kubernetes) をプロビジョニングします。
デフォルトでの出力拒否、改ざんが明らかな監査台帳、人間参加型ポリシー
ゲート、コスト/ループ ガードレール - REST、MCP、CLI、Web 経由で駆動
ダッシュボード。
カール -fsSL https://raw.githubusercontent.com/Runewardd/runeward/main/install.sh |しー
Homebrew、コンテナイメージ、ソースからの構築については、
をインストールします。次に、「クイックスタート」に進みます。
AI エージェントにシェル コマンドを実行させ、ファイルを編集し、パッケージをインストールし、
ネットワークは、rm -rf で間違ったディレクトリを指定し、ファイルを抽出するまでは役に立ちます。
シークレットを使用しないと、再試行ループで API バジェットが消費されます。生の隔離（「エージェントを投獄する」
ボックス内」) はテーブル ステークスです。 runeward はボックスの周囲にガバナンス層を追加します —
モデルが動作するように訓練されたことを期待するのではなく、モデルの外側でルールを強制する
(トレーニングではなくガバナンスを行う理由):
プロファイルはセキュリティ契約です。あなたが許可しないものはすべて拒否されます
デフォルトなので、爆発範囲は明示的です。
孤立しているだけではなく、統治されています。すべてのアクションは 1 つのパス、つまりポリシー、
承認ゲート、ガードレール、バックエンドエグゼクティブ、監査台帳 — 経由で届くかどうか
REST、ダッシュボード、または MCP。
構造により改ざんが明らかです。追加専用、ハッシュチェーン、ed25519 署名付き
台帳はすべての通話とその評決を記録し、独立した形式でエクスポートします。
検証可能な転写。
重要な場面では人間が関与します。アクションごとの許可 / 拒否 /
承認が必要との判定により、オペレーターの危険な操作が一時停止されます。
コストとループのガードレール。実時間、実行回数、下りのハードキャップ
リクエスト、トークン/予算の使用、および再試行ループ

検出。
認証されたマルチユーザー コントロール プレーン。ベアラートークン認証はデフォルトでオフ
ループバック、オプションのマルチプリンシパル RBAC (トークンごとのプロファイル/承認スコープ)、
対話型ログインによるプリンシパルごとのダッシュボード ビュー。
プラグイン可能なバックエンド。ゼロセットアップのラップトップを使用するための Docker/Podman、または Kubernetes
(厳密な L3 出力、CRD、アドミッション Webhook、PSA + NetworkPolicy マルチテナント)
生産およびフリート向け。
なぜガバナンスなのか — モデルをトレーニングするのではなく、モデルの外部でルールを強制します。
インストール — 1 行インストーラー、Homebrew、またはソースから。
クイックスタート — 60 秒以内に管理されたサンドボックスが完成します。
概念 — サンドボックス、フリート、ポリシー、出口、台帳。
プロファイル — 宣言的なセキュリティ契約。
アダプター — LangChain、CrewAI、LlamaIndex、OpenAI Agents SDK、Strands、Vercel AI SDK、LangChain.js。
セキュリティ モデル — runeward が保護するものと保護しないもの。
可観測性 — メトリクス、構造化ログ、テレメトリ。
runeward は、Apache License 2.0 に基づくオープンソースです。

## Original Extract

Governed execution cells for AI agents

runeward
Skip to content
runeward
Home
Initializing search
Runewardd/runeward
runeward
Runewardd/runeward
Home
Home
Table of contents
Install
Governed execution cells for AI agents.
Declarative profiles provision isolated sandboxes (Docker or Kubernetes) with
deny-by-default egress, a tamper-evident audit ledger, human-in-the-loop policy
gates, and cost/loop guardrails — driven over REST, MCP, a CLI, and a web
dashboard.
curl -fsSL https://raw.githubusercontent.com/Runewardd/runeward/main/install.sh | sh
Homebrew, container images, and building from source are covered in
Install . Then jump to the Quickstart .
Letting an AI agent run shell commands, edit files, install packages, and hit the
network is useful right up until it rm -rf s the wrong directory, exfiltrates a
secret, or burns your API budget in a retry loop. Raw isolation ("jail the agent
in a box") is table stakes. runeward adds the governance layer around the box —
enforcing the rules outside the model instead of hoping it was trained to behave
( why governance, not training ):
Profiles are a security contract. Everything you don't grant is denied by
default, so the blast radius is explicit.
Governed, not just isolated. Every action flows through one path — policy,
approval gate, guardrails, backend exec, audit ledger — whether it arrives via
REST, the dashboard, or MCP.
Tamper-evident by construction. An append-only, hash-chained, ed25519-signed
ledger records every call and its verdict, and exports as an independently
verifiable transcript.
Human-in-the-loop where it matters. Per-action allow / deny /
require-approval verdicts pause risky operations for an operator.
Cost and loop guardrails. Hard caps on wall-clock, exec count, egress
requests, and token/spend budgets, plus retry-loop detection.
Authenticated, multi-user control plane. Bearer-token auth by default off
loopback, optional multi-principal RBAC (per-token profile/approval scopes),
and per-principal dashboard views with an interactive login.
Pluggable backends. Docker/Podman for zero-setup laptop use, or Kubernetes
(strict L3 egress, CRDs, admission webhook, PSA + NetworkPolicy multi-tenancy)
for production and fleets.
Why governance — enforce rules outside the model, not by training it.
Install — one-line installer, Homebrew, or from source.
Quickstart — a governed sandbox in ~60 seconds.
Concepts — sandboxes, fleets, policy, egress, the ledger.
Profiles — the declarative security contract.
Adapters — LangChain, CrewAI, LlamaIndex, OpenAI Agents SDK, Strands, Vercel AI SDK, LangChain.js.
Security model — what runeward does and does not protect.
Observability — metrics, structured logs, and telemetry.
runeward is open source under the Apache License 2.0 .
