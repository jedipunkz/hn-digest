---
source: "https://fixdoc.dev"
hn_url: "https://news.ycombinator.com/item?id=48491929"
title: "CLI that scores Terraform change risk and stores your team's infra fixes"
article_title: "FixDoc — Three intelligence engines for infrastructure fixes"
author: "FixDoc"
captured_at: "2026-06-11T16:27:04Z"
capture_tool: "hn-digest"
hn_id: 48491929
score: 1
comments: 0
posted_at: "2026-06-11T15:43:51Z"
tags:
  - hacker-news
  - translated
---

# CLI that scores Terraform change risk and stores your team's infra fixes

- HN: [48491929](https://news.ycombinator.com/item?id=48491929)
- Source: [fixdoc.dev](https://fixdoc.dev)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T15:43:51Z

## Translation

タイトル: Terraform の変更リスクをスコアリングし、チームのインフラ修正を保存する CLI
記事のタイトル: FixDoc — インフラストラクチャ修正のための 3 つのインテリジェンス エンジン
説明: FixDoc はエラーを分類し、変更の影響を分析し、結果から学習します。インフラストラクチャにメモリを提供する 3 つのインテリジェンス エンジン。

記事本文:
修正ドキュメント
エンジン
分析する
使用例
デモ
GitHub ↗
始めましょう
3 つのインテリジェンス エンジン // インフラストラクチャ
インフラストラクチャにメモリの問題があります。
FixDoc はそれに 3 つの頭脳を与えます。
FixDoc はすべてのエラーをメモリ価値別に分類し、変更の影響を分析します。
スコアの説明と適用結果から学ぶ - チームがそれに基づいて構築します
過去の修正を繰り返すのではなく、
デフォルトではローカルファースト。修正を ~/.fixdoc に保存します。 Git 同期はオプションです。プライベート修正は決して同期されません。
fixdoc — 遅延優先キャプチャ
// インストール
60秒で始めましょう
FixDoc をインストールし、修正ライブラリの構築を開始します。
// インテリジェンス エンジン
3 つのインテリジェンス エンジン
各エンジンはインフラストラクチャのライフサイクルの異なるフェーズを処理します
チェンジ・インテリジェンス
Terraform の変更を適用する前に、その変更によって何が行われるかを理解します。
シグモイド式による影響スコアの変更 (0 ～ 100)
依存関係 DAG を介した BFS の伝播
ワイルドカード信頼検出による IAM 感度スコアリング
8 つの変更ドメインによるスマート フィックス マッチング
属性 + 履歴からのコンテキスト チェック
スコアの説明箇条書き (アクション、IAM、影響、履歴)
AI を活用した自然言語による説明
計画フィンガープリントリンクによる成果学習
障害インテリジェンス
コマンドが失敗した場合に、エラーを分類、診断、ルーティングします。
記憶に値する分類子 (記憶に値するか自明か)
何が起こったのかを説明するAIエラー診断
有効性スコアを使用した修正提案
Block Kit を使用した Slack プッシュ通知
反復ベースのプロモーション (3 つ以上の類似 → 記憶に値する)
メモリエンジン
タイプ認識レンダリングを使用して修正を保存、分類、取得します。
4 つのメモリ タイプ: 修正、チェック、プレイブック、インサイト
候補でのタイプ認識プレビュー レンダリング
効果の追跡 (applied_count、success_count)
セッション間でリンクするソース エラー ID
ビディレック

Git 同期のマークダウン往復
変更による影響を確認してください。その理由を理解してください。
1 つの変更がインフラストラクチャ全体にどのように伝播するかを理解する
aws_iam_role.app
削除
iam_policy_attachment
深さ: 1
役割ポリシー
深さ: 1
ラムダ関数.api
深さ: 2
ecs_task.worker
深さ: 2
ec2_instance.web
深さ: 2
s3_bucket.data
深さ: 3
dynamodb.セッション
深さ: 3
sqs_queue.events
深さ: 3
クラウドウォッチ.アラーム
深さ: 3
+25 削除アクション
+18 IAM ポリシーが変更されました
+15 BFS 深さ 3
+14 2 つの以前の修正一致
アニメーションを再生する
適用する前に変更の影響を分析する
# 適用する前に変更の影響を分析する
$ terraform show -json plan.tfplan > plan.json
$ fixdoc 分析 plan.json
スコア: 72.4 / 100 [高]
これが 72 点を獲得した理由:
IAM ロールの +25 削除アクション
+18 IAM ポリシーフィールドが変更されました
+15 BFS 伝播深さ 3
+14 2 つの以前の修正一致
関連する過去の修正: aws_iam_role [高: エラー コード: AccessDenied]
コンテキスト チェック: [attr] IAM 信頼ポリシー プリンシパルを確認します。
// 使用例
旅のあらゆる段階で
個々のエンジニア
自分専用の修正ライブラリを構築します。同じエラーを 2 回グーグルで検索するのはやめてください。ソリューションをいつでもすぐに入手できます。
チーム
Git 同期を介した共有ナレッジ ベース。 1 人のエンジニアがそれを解決すると、全員がアクセスできるようになります。重複したデバッグはもう必要ありません。
オンボーディング
新しいエンジニアは、初日からチームの修正履歴を継承します。立ち上げ時間を短縮し、生産性を加速します。
CI/CD パイプライン
すべての PR に対して変更の影響分析を実行します。ゲートは重大度別にマージされます。申請後の結果を追跡します。
✕ ダッシュボードとアラートを備えたインシデント管理 SaaS
✕ ログインする必要があるサービス カタログ ポータル
✕ 汎用ドキュメントツール
✕ 機能するには API キーが必要な AI ラッパー
✓ 既存のワークフローに適合するターミナルファーストキャプチャ
✓ マシン上に存在する検索可能な修正履歴
✓ インテリジェントな

エラーを分類し、変更の影響を分析し、結果から学ぶシステム
問題の説明、解決策、タグ、およびタイムスタンプ。各修正は、単純な JSON エントリ + マークダウン ファイルです。
すべてはローカルの ~/.fixdoc/ に保存されます。クラウド サービスもリモート データベースもありません。
いいえ、FixDoc はデフォルトでネットワーク要求を行いません。これは、テレメトリや分析のない純粋な CLI ツールです。
AI診断とAI説明は完全にオプションです。 API キーはマシン上に残ります。 --diagnose または --ai-explain を明示的に使用しない限り、データは送信されません。 FixDoc はそれらなしでも完全に機能します。

## Original Extract

FixDoc classifies errors, analyzes change impact, and learns from outcomes. Three intelligence engines that give your infrastructure a memory.

FixDoc
Engines
Analyze
Use Cases
Demo
GitHub ↗
Get Started
Three Intelligence Engines // Infrastructure
Your infrastructure has a memory problem.
FixDoc gives it three brains.
FixDoc classifies every error by memory-worthiness, analyzes change impact with
score explanations, and learns from apply outcomes — so your team builds on
past fixes instead of repeating them.
Local-first by default. Stores fixes in ~/.fixdoc . Git sync is optional. Private fixes never sync.
fixdoc — defer-first capture
// INSTALL
Get started in 60 seconds
Install FixDoc and start building your fix library.
// INTELLIGENCE ENGINES
Three intelligence engines
Each engine handles a different phase of your infrastructure lifecycle
Change Intelligence
Understands what a Terraform change will do before you apply it.
Change impact scoring with sigmoid formula (0–100)
BFS propagation through dependency DAG
IAM sensitivity scoring with wildcard trust detection
Smart fix matching with 8 change domains
Contextual checks from attribute + history
Score explanation bullets (action, IAM, impact, history)
AI-powered natural language explanations
Outcome learning with plan fingerprint linking
Failure Intelligence
Classifies, diagnoses, and routes errors when your commands fail.
Memory-worthiness classifier (memory-worthy vs self-explanatory)
AI error diagnosis explaining WHY it happened
Fix suggestions with effectiveness scores
Slack push notifications with Block Kit
Recurrence-based promotion (3+ similar → memory-worthy)
Memory Engine
Stores, classifies, and retrieves fixes with type-aware rendering.
4 memory types: fix, check, playbook, insight
Type-aware preview rendering in suggestions
Effectiveness tracking (applied_count, success_count)
Source error ID linking across sessions
Bidirectional markdown round-trip for Git sync
See your change impact. Understand why.
Understand how one change propagates through your infrastructure
aws_iam_role.app
DELETE
iam_policy_attachment
depth: 1
role_policy
depth: 1
lambda_function.api
depth: 2
ecs_task.worker
depth: 2
ec2_instance.web
depth: 2
s3_bucket.data
depth: 3
dynamodb.sessions
depth: 3
sqs_queue.events
depth: 3
cloudwatch.alarms
depth: 3
+25 delete action
+18 IAM policy changed
+15 BFS depth 3
+14 2 prior fix matches
Replay Animation
analyze change impact before applying
# Analyze your change impact before applying
$ terraform show -json plan.tfplan > plan.json
$ fixdoc analyze plan.json
Score: 72.4 / 100 [HIGH]
Why this scored 72:
+25 delete action on IAM role
+18 IAM policy field changed
+15 BFS propagation depth 3
+14 2 prior fix matches
Relevant Past Fixes: aws_iam_role [high: error code: AccessDenied]
Contextual Checks: [attr] Verify IAM trust policy principals
// USE CASES
For every stage of your journey
Individual Engineers
Build your personal fix library. Stop googling the same error twice. Your solutions, always at your fingertips.
Teams
Shared knowledge base via Git sync. When one engineer solves it, everyone has access. No more duplicate debugging.
Onboarding
New engineers inherit your team’s fix history on day one. Reduce ramp-up time, accelerate productivity.
CI/CD Pipelines
Run change impact analysis on every PR. Gate merges by severity. Track outcomes post-apply.
✕ An incident management SaaS with dashboards and alerts
✕ A service catalog portal you have to log into
✕ A generic documentation tool
✕ An AI wrapper that requires an API key to function
✓ Terminal-first capture that fits your existing workflow
✓ A searchable fix history that lives on your machine
✓ An intelligent system that classifies errors, analyzes change impact, and learns from outcomes
Issue descriptions, resolutions, tags, and timestamps. Each fix is a simple JSON entry + a markdown file.
Everything is stored locally in ~/.fixdoc/ . No cloud service, no remote database.
No. FixDoc makes zero network requests by default. It’s a pure CLI tool with no telemetry or analytics.
AI diagnosis and AI explain are fully optional. Your API key stays on your machine. No data is sent unless you explicitly use --diagnose or --ai-explain . FixDoc works perfectly without them.
