---
source: "https://github.com/param087/agent-ml-skills"
hn_url: "https://news.ycombinator.com/item?id=48420791"
title: "Agent-ML-skills – Teach Codex/Claude/Cursor to stop making ML mistakes"
article_title: "GitHub - param087/agent-ml-skills: Production-grade Machine Learning, Data Science & MLOps skills for AI coding agents (Codex, Claude Code, Cursor, OpenCode). One npx command to install. · GitHub"
author: "param087"
captured_at: "2026-06-06T04:27:00Z"
capture_tool: "hn-digest"
hn_id: 48420791
score: 1
comments: 0
posted_at: "2026-06-06T02:21:42Z"
tags:
  - hacker-news
  - translated
---

# Agent-ML-skills – Teach Codex/Claude/Cursor to stop making ML mistakes

- HN: [48420791](https://news.ycombinator.com/item?id=48420791)
- Source: [github.com](https://github.com/param087/agent-ml-skills)
- Score: 1
- Comments: 0
- Posted: 2026-06-06T02:21:42Z

## Translation

タイトル: エージェント ML スキル – コーデックス/クロード/カーソルに ML の間違いを防ぐよう教える
記事のタイトル: GitHub - param087/agent-ml-skills: AI コーディング エージェント (Codex、Claude Code、Cursor、OpenCode) のための本番グレードの機械学習、データ サイエンス、MLOps スキル。インストールする 1 つの npx コマンド。 · GitHub
説明: AI コーディング エージェント (Codex、Claude Code、Cursor、OpenCode) のための本番グレードの機械学習、データ サイエンス、MLOps スキル。インストールする 1 つの npx コマンド。 - param087/エージェント-ml-スキル

記事本文:
GitHub - param087/agent-ml-skills: AI コーディング エージェント (Codex、Claude Code、Cursor、OpenCode) のための本番グレードの機械学習、データ サイエンス、MLOps スキル。インストールする 1 つの npx コマンド。 · GitHub
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
パラメータ087
/
エージェント-ml-スキル
公共
通知

オン
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .github .github スクリプト スクリプト スキル スキル .gitignore .gitignore COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md install.mjs install.mjs package.json package.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
AI コーディング エージェント向けの実稼働グレードの機械学習、データ サイエンス、MLOps スキル。
コーディング エージェントは優れたジェネラリストですが、同じ ML の間違いを何度も犯します。つまり、相互検証への前処理の漏洩、不均衡なデータの正確なスコアリング、model.eval() の忘れ、高密度のみの取得による RAG の構築などです。 Agent-ML-Skills は、経験豊富な ML エンジニアが実際にどのように機能するかをエージェントに教える、実戦テストされた 15 のスキルを厳選したパックです。これにより、エージェントは推測をやめることになります。
Codex、Claude Code、Cursor、および OpenCode で動作します。
1 つのコマンドですべてのスキルをエージェントにインストールします。インストールや依存関係は必要ありません。
# コーデックス
npx エージェント-ml-スキル インストール --target codex
# クロードコード
npx エージェント-ml-スキル インストール --target クロード
# カーソル
npx エージェント-ml-スキル インストール --ターゲット カーソル --スコープ プロジェクト
# オープンコード
npx エージェント-ml-スキル インストール --target opencode
#あらゆるもの、どこでも
npx Agent-ml-skills install --target all
まずは中身を見てみましょう:
npx エージェント-ml-スキル リスト
その後、エージェントを再起動 (または新しいセッションを開始) すると、タスクが一致したときに適切なスキルが自動的に選択されます。
スキルは、エージェントにいつ使用するか、どのようにタスクを適切に実行するかを伝える YAML フロントマターを含む単一の Markdown ファイルです。
---
名前 : sklearn パイプライン
説明 : 前処理をリークしてはならない scikit-learn モデルを構築する場合に使用します...
---
# scikit-learn パイプライン
...ワークフロー、コードパターン、ピット

落ちる、手を離す...
スキルをサポートするエージェントは、事前に説明を読み込み、タスクが一致する場合にのみ本文を取り込みます。そのため、すべてのプロンプトを肥大化させることなく、専門家のガイダンスが得られます。
スキル
次の場合に使用します…
探索的データ分析
新しいデータセットから始める - プロファイリング、分布、相関、漏洩、viz。
データクリーニング
欠損値、重複、型、外れ値をトレーニングのみの代入で処理します。
特徴量エンジニアリング
エンコーディング、スケーリング、日時/テキスト/集計機能、リークセーフなターゲット エンコーディング。
パンダのパターン
慣用的なベクトル化されたメモリ効率の高いパンダを作成します ( SettingWithCopyWarning はありません)。
不均衡なデータ
ターゲットはまれです (詐欺/チャーン/病気) - メトリクス、SMOTE、クラスの重み、しきい値。
モデリング
スキル
次の場合に使用します…
スクラーンパイプライン
前処理が CV に漏れてはいけない scikit-learn モデルを構築します。
pytorch-トレーニングループ
PyTorch ループの作成/レビュー - 評価モード、AMP、チェックポイント、デバイス。
モデル評価
メトリクスの選択、検証、キャリブレーション、混同行列分析。
ハイパーパラメータ調整
パラメータの最適化 — ランダムと Optuna、リークセーフな CV、早期停止、予算。
LLM と GenAI
スキル
次の場合に使用します…
llm-微調整
LLM の微調整 – フル vs LoRA/QLoRA、データ フォーマット、トランスフォーマー/PEFT/TRL。
ラグパイプライン
RAG の構築 — チャンキング、埋め込み、ハイブリッド + 再ランキング取得、評価。
MLOps と信頼性
スキル
次の場合に使用します…
実験追跡
実験には、MLflow/W&B、ログ内容、レジストリなどの比較/再現が必要です。
再現可能なml
シード、環境ピンニング、データのバージョン管理、CUDA の決定性など、結果は再現可能でなければなりません。
ml-デバッグ
モデルが学習しない、損失が NaN である、またはメトリクスが良すぎる - 診断決定ツリー。
モデルを務める
API の背後にデプロイ — FastAPI、安全なアーティファクトの読み込み、バッチ処理、ONNX、モニタリング。
U

賢者
npx エージェント-ml-スキル <コマンド> [オプション]
コマンド
list 利用可能なスキルの一覧
install スキルをエージェントにインストールする
オプション
-t, --target <名前> コーデックス |クロード |オープンコード |カーソル |すべて
--scope <スコープ> グローバル (デフォルト) |プロジェクト
--skills <a,b,c> カンマ区切りのサブセット (デフォルト: all)
--dir <パス> カスタム ディレクトリにインストールします (ターゲットをオーバーライドします)
-f、--force 既存のスキルを上書きします
-h, --help このヘルプを表示します
例
# LLM スキルだけを現在のプロジェクトに導入
npx Agent-ml-skills install --target claude --skills rag-pipeline、llm-finetuning --scope プロジェクト
# カスタムエージェントディレクトリへ
npx エージェント-ml-スキル インストール --dir ./my-agent/skills
# 再インストールして上書きする
npx エージェント-ml-スキル インストール --target codex --force
スキルがインストールされる場所
ターゲット
グローバル
プロジェクト
コーデックス
~/.codex/スキル
.コーデックス/スキル
クロード・コード
~/.claude/スキル
.クロード/スキル
オープンコード
~/.config/opencode/skills
.opencode/スキル
カーソル
—
.cursor/rules (フラットな .md ルール)
設計原則
デフォルトでは漏洩に対して安全です。すべてのデータ スキルは、電車でのみ変換に適合します。
抽象よりも具体。曖昧なアドバイスではなく、実際のコード パターン。
落とし穴も含まれています。各スキルは、エージェントが実際に犯す間違いで終わります。
構成可能。スキルは相互に受け継がれます (EDA → クリーニング → 機能 → パイプライン → 評価 → 提供)。
依存関係のないインストーラー。純粋なノード。インストールするものや信頼するものは何もありません。
新しいスキルや改善は大歓迎です。 CONTRIBUTING.md を参照してください。すべてのスキルは CI で検証されます。
ノードスクリプト/validate-skills.mjs
エージェントに習得してもらいたい ML ワークフローがある場合は、スキル リクエストを開きます。
Param Bhavsar によって構築されました — Google Summer of Code '19 @ TensorFlow、元 HSBC。これによりデバッグ セッションが節約できる場合は、⭐ を付けると他の人が見つけやすくなります。
AI コーディングのための本番グレードの機械学習、データ サイエンス、MLOps スキル

紳士 (コーデックス、クロード コード、カーソル、OpenCode)。インストールする 1 つの npx コマンド。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Production-grade Machine Learning, Data Science & MLOps skills for AI coding agents (Codex, Claude Code, Cursor, OpenCode). One npx command to install. - param087/agent-ml-skills

GitHub - param087/agent-ml-skills: Production-grade Machine Learning, Data Science & MLOps skills for AI coding agents (Codex, Claude Code, Cursor, OpenCode). One npx command to install. · GitHub
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
param087
/
agent-ml-skills
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .github .github scripts scripts skills skills .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md install.mjs install.mjs package.json package.json View all files Repository files navigation
Production-grade Machine Learning, Data Science & MLOps skills for AI coding agents.
Coding agents are great generalists but make the same ML mistakes over and over : leaking preprocessing into cross-validation, scoring imbalanced data with accuracy, forgetting model.eval() , building RAG with dense-only retrieval. agent-ml-skills is a curated pack of 15 battle-tested skills that teach your agent how an experienced ML engineer actually works — so it stops guessing.
Works with Codex, Claude Code, Cursor, and OpenCode .
Install all skills into your agent with one command — no install, no dependencies :
# Codex
npx agent-ml-skills install --target codex
# Claude Code
npx agent-ml-skills install --target claude
# Cursor
npx agent-ml-skills install --target cursor --scope project
# OpenCode
npx agent-ml-skills install --target opencode
# Everything, everywhere
npx agent-ml-skills install --target all
Browse what's inside first:
npx agent-ml-skills list
Then restart your agent (or start a new session) and it will pick the right skill up automatically when your task matches.
A skill is a single Markdown file with YAML frontmatter telling the agent when to use it and how to do the task well:
---
name : sklearn-pipelines
description : Use when building scikit-learn models that must not leak preprocessing...
---
# scikit-learn Pipelines
...workflow, code patterns, pitfalls, hand-off...
Agents that support skills load the description up front and pull in the full body only when the task matches — so you get expert guidance without bloating every prompt .
Skill
Use when…
exploratory-data-analysis
Starting on a new dataset — profiling, distributions, correlations, leakage & viz.
data-cleaning
Handling missing values, duplicates, types, outliers — with train-only imputation.
feature-engineering
Encoding, scaling, datetime/text/aggregation features, leakage-safe target encoding.
pandas-patterns
Writing idiomatic, vectorized, memory-efficient pandas (no SettingWithCopyWarning ).
imbalanced-data
The target is rare (fraud/churn/disease) — metrics, SMOTE, class weights, thresholds.
Modeling
Skill
Use when…
sklearn-pipelines
Building scikit-learn models that must not leak preprocessing into CV.
pytorch-training-loop
Writing/reviewing a PyTorch loop — eval modes, AMP, checkpointing, devices.
model-evaluation
Choosing metrics, validating, calibration, confusion-matrix analysis.
hyperparameter-tuning
Optimizing params — random vs Optuna, leakage-safe CV, early stopping, budget.
LLMs & GenAI
Skill
Use when…
llm-finetuning
Fine-tuning an LLM — full vs LoRA/QLoRA, data formatting, transformers/PEFT/TRL.
rag-pipeline
Building RAG — chunking, embeddings, hybrid + reranking retrieval, eval.
MLOps & reliability
Skill
Use when…
experiment-tracking
Experiments need comparing/reproducing — MLflow/W&B, what to log, registry.
reproducible-ml
A result must be reproducible — seeds, env pinning, data versioning, CUDA determinism.
ml-debugging
A model won't learn, loss is NaN, or metrics look too good — a diagnosis decision tree.
model-serving
Deploying behind an API — FastAPI, safe artifact loading, batching, ONNX, monitoring.
Usage
npx agent-ml-skills <command> [options]
Commands
list List available skills
install Install skills into an agent
Options
-t, --target <name> codex | claude | opencode | cursor | all
--scope <scope> global (default) | project
--skills <a,b,c> comma-separated subset (default: all)
--dir <path> install into a custom directory (overrides target)
-f, --force overwrite existing skills
-h, --help show this help
Examples
# Just the LLM skills, into the current project
npx agent-ml-skills install --target claude --skills rag-pipeline,llm-finetuning --scope project
# Into a custom agent directory
npx agent-ml-skills install --dir ./my-agent/skills
# Re-install and overwrite
npx agent-ml-skills install --target codex --force
Where skills get installed
Target
Global
Project
Codex
~/.codex/skills
.codex/skills
Claude Code
~/.claude/skills
.claude/skills
OpenCode
~/.config/opencode/skills
.opencode/skills
Cursor
—
.cursor/rules (flat .md rules)
Design principles
Leakage-safe by default. Every data skill fits transforms on train only.
Concrete over abstract. Real code patterns, not vague advice.
Pitfalls included. Each skill ends with the mistakes agents actually make.
Composable. Skills hand off to each other (EDA → cleaning → features → pipeline → eval → serving).
Zero-dependency installer. Pure Node, nothing to install, nothing to trust.
New skills and improvements are very welcome — see CONTRIBUTING.md . Every skill is validated in CI:
node scripts/validate-skills.mjs
Open a skill request if there's an ML workflow you want your agent to master.
Built by Param Bhavsar — Google Summer of Code '19 @ TensorFlow, ex-HSBC. If this saves you a debugging session, a ⭐ helps others find it.
Production-grade Machine Learning, Data Science & MLOps skills for AI coding agents (Codex, Claude Code, Cursor, OpenCode). One npx command to install.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
