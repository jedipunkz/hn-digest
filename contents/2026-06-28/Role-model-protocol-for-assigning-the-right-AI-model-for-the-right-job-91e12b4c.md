---
source: "https://role-model.dev/"
hn_url: "https://news.ycombinator.com/item?id=48712295"
title: "Role-model: protocol for assigning the right AI model for the right job"
article_title: "role-model"
author: "handfuloflight"
captured_at: "2026-06-28T22:20:36Z"
capture_tool: "hn-digest"
hn_id: 48712295
score: 1
comments: 0
posted_at: "2026-06-28T22:16:36Z"
tags:
  - hacker-news
  - translated
---

# Role-model: protocol for assigning the right AI model for the right job

- HN: [48712295](https://news.ycombinator.com/item?id=48712295)
- Source: [role-model.dev](https://role-model.dev/)
- Score: 1
- Comments: 0
- Posted: 2026-06-28T22:16:36Z

## Translation

タイトル: ロールモデル: 適切なジョブに適切な AI モデルを割り当てるためのプロトコル
記事タイトル: ロールモデル
説明: パッケージ化されたリファレンス ランタイム、説明可能なルーターの決定、および実際に検査できるプロトコルを備えた機能を認識した AI ルーティング。

記事本文:
role-model role-model Search ⌘ K @trydotworks GitHub role-model Search ⌘ K role-model はじめる インストール 最初の起動と接続モデル 完全なベンチマークの実行 ルーティング戦略の選択と保存 最初のリクエストと決定の検査 統合 Pi 統合 ダウンストリーム OpenAI 検出 ランタイム ランタイム UI ツアー プロバイダー接続 モデルとロールのアクティブ化 ベンチマークと評価 観察とテレメトリ分析 ルーティング制御と意思決定のレビュー ルーターの概要 ルーティングモード、ローカリティ、および実行 スコアリング戦略とトレードオフ 候補者の選択と適格性 スコアリング、タイブレーク、決定 フォールバック、失敗、可観測性 ルーティングのエンドツーエンドの仕組み プロトコルとルーターのマッピングの概念 ロールモデルの仕組み 役割、タスク、機能 エンドポイントとプロファイル ポリシーとオブザーバビリティ プロトコルの概要 ルーティングの概要 リファレンスの概要 正規スキーマのリファレンス 理由コードと拒否分類 プロトコルプロトコル エンドポイント ID 宣言された機能プロファイル 観察されたパフォーマンス プロファイル 機能分類 役割とタスク ルーティング ポリシー ルーター決定アーティファクト トレースおよび使用アーティファクト ロールモデル ロールモデル
パッケージ化されたリファレンス ランタイム、説明可能なルーターの決定、実際に検査できるプロトコルを備えた機能を意識した AI ルーティング。
マークダウンをコピーして開く
role-model は、機能を認識した AI ルーティング用のオープン プロトコルと、パッケージ化されたリファレンス ルーター ランタイムです。
これにより、システムに次のことを説明する永続的な方法が提供されます。
どのような役割とタスクが求められているか
どの具体的なエンドポイントが作業を満たすことができるか
最終的なルーティング決定が行われた理由
ルーターはモデル名だけでは選択しません。ロールとタスクを使用して具体的なエンドポイント間をルーティングします
メタデータ、宣言された機能、ルーティング ポリシー、および観察されたパフォーマンス。
最初に起動します

d接続モデル
ルーティング戦略を選択して保存します
最初のリクエストを送信し、決定を検査します
パイを使用しますか?まず、Role-Model ランタイムをインストールして起動し、次の手順に従います。
@try-works/pi-role-model をインストールし、 /role-model setup を実行して選択するための Pi 統合
別名。
大まかに言うと、ロールモデルはルーティングをいくつかの安定した部分に分割します。
リクエストには、タスクのタイプ、必要な機能、モダリティ、ツールのニーズ、および制約が記述されます。
役割とタスクは、仕事の意味論的な形状を記述します。
エンドポイント ID とプロファイルは、抽象的なモデル名ではなく、具体的なルーティング可能なエンドポイントを記述します。
ルーティング ポリシーは、ハード拒否、プリファレンス、予算、および決定的なタイブレーク ルールを適用します。
可観測性アーティファクトは、決定、トレース、使用状況、および測定されたパフォーマンスを記録します。
これにより、ルーティングが説明可能になり、さまざまなプロバイダー、ホスト、展開形態にわたって移植可能になります。
ルーターが決定する方法
参照ルーターは安定したフローに従います。
リクエストの意図を正規化します。リクエストとロール/タスクのメタデータから効果的なポリシーのスナップショットを構築します。
候補セットを絞り込みます。要求されたロール、タスク、およびポリシーのスコープに一致するエンドポイントのみを保持します。
厳格な適格性チェックを適用します。機能、モダリティ、ツール、地域性、予算、または拘束力の要件を満たさないエンドポイントを拒否します。
適格なエンドポイントをスコアリングします。最初に測定された証拠を使用して、次に宣言されたデータと中立的なデフォルトを使用して、品質、レイテンシ、スループット、コスト、信頼性、優先度を比較します。
説明可能な決定を下します。選択したエンドポイント、フォールバック、除外、および選択理由を含む RouterDecision を返します。
結果は、単なる隠れた実行時の推測ではなく、後で検査できるほど十分に決定的です。
現在のベースライン ロール セットには次のものが含まれます。
完全な役割とタスクのメンタル モデルについては、以下をお読みください。
役割、タスク、および機能。ディー

プロトコル契約ごとに
依然として 役割とタスク に存在します。
初回セットアップのアーキテクチャ
標準的な初回実行シーケンスは次のようになります。
パッケージ化されたランタイムをインストールして起動する
実際に使用する予定のローカルまたはリモートのエンドポイントに接続します
モデルをアクティブにして役割を割り当てる
実際の候補セットに対して完全なベンチマークを実行します
ルーティング戦略を選択して保存します
実際にルーティングされたリクエストで検証し、決定を検査します
Pi などのダウンストリーム クライアントは、ランタイムがインストールされ構成された後に参加します。彼らはロールモデルを発見します
ランタイム セットアップ自体を所有するのではなく、ダウンストリームの OpenAI 検出コントラクトを通じてエイリアスを作成します。
これにより、ルーティング戦略の選択は推測に基づいたものではなく証拠に基づいたものになります。
プロバイダーとローカル バックエンドを接続する ルーティングする予定の正確なローカルおよびリモートの実行パスを配線します。モデルをアクティブ化してロールを割り当てる 実際のエンドポイント インベントリを公開し、各モデルを、それが提供する必要があるロールにのみバインドします。完全なベンチマークを実行し、構成された候補セットを評価し、測定された品質を観察されたプロファイルに書き込みます。ルーティング戦略を選択して保存します。ベンチマークの証拠が存在した後、バランスの取れた品質、レイテンシー、またはコストを設定します。実際にルーティングされたリクエストを検証する ルーターを検査し、勝者、フォールバック、および理由が証拠と一致することを確認するために観察します。初回セットアップ パスでは、最初に実際の候補セットを確立し、次にそれをベンチマークし、証拠から戦略を選択する必要があります。インストール
GitHub リリースからパッケージ化されたロールモデル ルーター ランタイムをインストールするか、リポジトリを開発している場合はソース ビルドを選択します。

## Original Extract

Capability-aware AI routing with a packaged reference runtime, explainable router decisions, and a protocol you can actually inspect.

role-model role-model Search ⌘ K @trydotworks GitHub role-model Search ⌘ K role-model Get Started Install First launch and connect models Run the full benchmark Choose and save the routing strategy First request and inspect the decision Integrations Pi integration Downstream OpenAI discovery Runtime Runtime UI tour Provider connections Models and role activation Benchmarks and evaluation Observe and telemetry analytics Routing controls and decision review Router Overview Routing modes, locality, and execution Scoring strategies and tradeoffs Candidate selection and eligibility Scoring, tie-breaks, and decisions Fallbacks, failures, and observability How routing works end to end Protocol-to-router mapping Concepts How role-model works Roles, tasks, and capabilities Endpoints and profiles Policy and observability Protocol overview Routing overview Reference Overview Canonical schemas reference Reason codes and rejection taxonomy Protocol Protocol Endpoint identity Declared capability profiles Observed performance profiles Capability taxonomy Roles and tasks Routing policy Router decision artifact Trace and usage artifacts role-model role-model
Capability-aware AI routing with a packaged reference runtime, explainable router decisions, and a protocol you can actually inspect.
Copy Markdown Open
role-model is an open protocol for capability-aware AI routing, plus a packaged reference router runtime.
It gives a system a durable way to describe:
which roles and tasks are being asked for
which concrete endpoints can satisfy the work
why the final routing decision was made
The router does not pick by model name alone. It routes across concrete endpoints using role and task
metadata, declared capability, routing policy, and observed performance.
First launch and connect models
Choose and save the routing strategy
Send the first request and inspect the decision
Using Pi? Install and launch the Role-Model runtime first, then follow
Pi integration to install @try-works/pi-role-model , run /role-model setup , and choose
an alias.
At a high level, role-model separates routing into a few stable pieces:
Requests describe task type, required capabilities, modalities, tool needs, and constraints.
Roles and tasks describe the semantic shape of the work.
Endpoint identities and profiles describe concrete routable endpoints rather than abstract model names.
Routing policy applies hard denies, preferences, budgets, and deterministic tie-break rules.
Observability artifacts record the decision, traces, usage, and measured performance.
That makes routing explainable and portable across different providers, hosts, and deployment shapes.
How the router makes a decision
The reference router follows a stable flow:
Normalize request intent. Build the effective policy snapshot from the request plus role/task metadata.
Narrow the candidate set. Keep only endpoints that match the requested role, task, and policy scope.
Apply hard eligibility checks. Reject endpoints that fail capability, modality, tool, locality, budget, or binding requirements.
Score the eligible endpoints. Compare quality, latency, throughput, cost, reliability, and preference using measured evidence first, then declared data and neutral defaults.
Emit an explainable decision. Return a RouterDecision with the chosen endpoint, fallbacks, exclusions, and selection reasons.
The result is deterministic enough to inspect later, not just a hidden runtime guess.
The current baseline role set includes:
For the full role and task mental model, read
Roles, tasks, and capabilities . The deeper protocol contract
still lives in Roles and tasks .
The first-time setup architecture
The canonical first-run sequence is now:
install and launch the packaged runtime
connect the local or remote endpoints you actually plan to use
activate models and assign roles
run the full benchmark on that real candidate set
choose and save the routing strategy
validate with a real routed request and inspect the decision
Downstream clients such as Pi join after the runtime is installed and configured. They discover Role-Model
aliases through the downstream OpenAI discovery contract instead of owning runtime setup themselves.
This keeps routing strategy selection evidence-based instead of guess-based.
Connect providers and local backends Wire in the exact local and remote execution paths you plan to route across. Activate models and assign roles Publish the real endpoint inventory and bind each model only to the roles it should serve. Run the full benchmark Grade the configured candidate set and write measured quality into observed profiles. Choose and save routing strategy Set balanced, quality, latency, or cost after the benchmark evidence exists. Validate a real routed request Inspect Router and Observe to confirm the winner, fallbacks, and reasons match the evidence. The first-time setup path should establish the real candidate set first, then benchmark it, then choose strategy from evidence. Install
Install the packaged role-model router runtime from GitHub Releases, or choose a source build if you are developing the repo.
