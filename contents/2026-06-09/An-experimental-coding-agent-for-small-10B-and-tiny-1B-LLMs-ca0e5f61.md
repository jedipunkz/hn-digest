---
source: "https://github.com/brainless/nocodo"
hn_url: "https://news.ycombinator.com/item?id=48459126"
title: "An experimental coding agent for small (<10B) and tiny (<1B) LLMs"
article_title: "GitHub - brainless/nocodo: Coding agent built around small (<10B) models for local development with a Rust (Actix Web, Diesel)/Typescript (Solid, Tailwind) opinionated stack. No coding knowledge needed. Multiple agents and roles - Product Owner, PM, Engineering Manager, Rust Engineer, SolidJS Engine\n[truncated]"
author: "brainless"
captured_at: "2026-06-09T13:09:49Z"
capture_tool: "hn-digest"
hn_id: 48459126
score: 1
comments: 2
posted_at: "2026-06-09T10:25:15Z"
tags:
  - hacker-news
  - translated
---

# An experimental coding agent for small (<10B) and tiny (<1B) LLMs

- HN: [48459126](https://news.ycombinator.com/item?id=48459126)
- Source: [github.com](https://github.com/brainless/nocodo)
- Score: 1
- Comments: 2
- Posted: 2026-06-09T10:25:15Z

## Translation

タイトル: 小規模 (<10B) および極小 (<1B) LLM 用の実験的コーディング エージェント
記事のタイトル: GitHub - Brainless/nocodo: Rust (Actix Web、Diesel)/Typescript (Solid、Tailwind) 独自のスタックを使用したローカル開発用の小規模 (<10B) モデルを中心に構築されたコーディング エージェント。コーディングの知識は必要ありません。複数のエージェントと役割 - プロダクトオーナー、PM、エンジニアリングマネージャー、Rust エンジニア、SolidJS エンジン
[切り捨てられた]
説明: コーディング エージェントは、Rust (Actix Web、Diesel)/Typescript (Solid、Tailwind) 独自のスタックを使用したローカル開発用の小規模 (<10B) モデルを中心に構築されました。コーディングの知識は必要ありません。複数のエージェントと役割 - プロダクト オーナー、PM、エンジニアリング マネージャー、Rust エンジニア、SolidJS エンジニアなどが連携して、
[切り捨てられた]

記事本文:
GitHub - Brainless/nocodo: コーディング エージェントは、Rust (Actix Web、Diesel)/Typescript (Solid、Tailwind) 独自のスタックを使用したローカル開発用の小規模 (<10B) モデルを中心に構築されました。コーディングの知識は必要ありません。複数のエージェントと役割 - プロダクト オーナー、PM、エンジニアリング マネージャー、Rust エンジニア、SolidJS エンジニアなどが連携してフルスタック アプリを構築します。 · GitHub
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
別のタブまたはウィンドウでサインアウトしました。リロードして更新します

セッション。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
頭の悪い
/
ノコド
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
141 コミット 141 コミット .githooks .githooks admin-gui admin-gui エージェント エージェント バックエンド バックエンド gui schema-codegen schema-codegen スクリプト スクリプト 共有タイプ 共有タイプ tauri tauri ウェブサイト ウェブサイト .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml DEVELOP.md DEVELOP.md LICENSE LICENSE MIXTURE_OF_AGENTS.md MIXTURE_OF_AGENTS.md README.md README.md project.toml.template project.toml.template すべてのファイルを表示 リポジトリ ファイルのナビゲーション
実験中・進行中の作業
特定の独自のスタックを対象とした、小規模 (10B 未満) のローカル モデル用のコーディング エージェント。目標はまだ完全に達成されていませんが、初期の結果は継続するのに十分な有望なものです。
ローカル開発用の小規模 (<10B) (場合によっては小規模 (<1B)) モデルを中心に構築されたコーディング エージェント。ターゲット スタックは独自に定められています。バックエンドは Rust (Actix Web + Diesel)、フロントエンドは TypeScript (SolidJS + Tailwind) です。複数のエージェントの役割 (プロダクト オーナー、プロジェクト マネージャー、エンジニアリング マネージャー、Rust エンジニア、SolidJS エンジニアなど) が最終的には協力してフルスタック アプリを構築することになります。コーディングの知識は必要ありません。
これでは、汎用のコーディング エージェントは生成されません。賭けはより狭く、より実現可能です。nocodo は、ローカル LLM を使用して、デスクトップ アプリから完全に構築および管理される中小企業 CRUD アプリを、自分が望むものを説明できる人なら誰でも手に入れることができます。
一般的なコーディング エージェントは、終わりのない問題に直面するため、大規模で高価なモデルを必要とします。 nocodo は逆のアプローチをとります: mak

各問題を最大限に絞り込み、小さなモデルで確実に解決します。
スタックは固定されています。パターンは既知です。 Diesel モデル関数、Actix ルート ハンドラー、SolidJS フォーム コンポーネントは、それぞれ小さく学習可能な形状を持っています。 nocodo は、モデルに物事を理解するよう依頼するのではなく、個別のコーディング タスクごとに専用のエージェント モードを構築し、具体的な例をロードして、正確に 1 つのことを要求します。モデルはパターンをコピーします。確定的な後処理により出力がクリーンアップされます。
結果: Qwen3.5-0.8B — ギガバイト未満に収まるモデル — 正しい Diesel ORM 関数を生成します。それが行動の前提です。制約にもかかわらず機能するのではなく、制約があるために機能します。
nocodo は、スタックの各レイヤーのエージェント/モード カバレッジを一度に 1 レイヤーずつ構築します。各モードは自己完結型のプロンプト ビルダーです。適切なコード コンテキストを抽出し、サンプルが豊富なシングルショット プロンプトを構築し、出力を決定的に後処理します。
レイヤー
モード
ステータス
ディーゼルORM
モデル impl 関数、モデル構造体、スキーマ定義
✅ 働いています
Actix Web
コントローラー（ハンドラー）、ルーター、ミドルウェア
計画済み
認証と権限
セッション、JWT、ロールベースのアクセス
計画済み
TypeScript フロントエンド — 次へ
レイヤー
モード
ステータス
SolidJS
ルーティング、状態/コンテキスト、フォーム、ビュー
計画済み
追い風 + DaisyUI
コンポーネントのスタイル設定
計画済み
プラットフォーム層 — ロードマップ上
能力
説明
ステータス
工具管理
Rust コンパイラ、ノード、プロジェクトの依存関係をインストールして構成する - 決定的で、LLM は関与しません
計画済み
Gitの統合
ブランチ、コミット、PR の GitHub/GitLab サポート — 決定論的コード、エージェント管理
計画済み
導入
VPS/SSH ビルド、または管理対象ホスト (Render、Railway、Fly、Cloudflare) — 完全に API 統合された、同じエージェント/モード パターン
計画済み
レイヤー内のいくつかのモードが確実に動作すると、

nocodo 独自のコーディング エージェントがビルドに参加するため、新しいものの開発が加速します。
ほとんどのコーディング エージェントは、GPT-4 または Claude を使用して小規模モデルをルーティングします。 nocodo は意図的に問題を制限するので、より小さなモデルが実行可能になります。
シングルショット — 1 つのプロンプト、1 つの応答で出力を解析します。ツールループや複数ターンの推論はありません。
例主導 — プロンプトには、第一原理から推論するための指示ではなく、モデルがコピーできる 10 ～ 15 個の具体的なパターンが含まれています。
モードごとの範囲が狭い — 各モードは、正確に 1 つのアーティファクト (1 つの関数、1 つの構造体、1 つのルート) を生成します。
決定論的な後処理 — インポートは取り除かれ、正しく再挿入され、コード フェンスがアンラップされ、<think> ブロックが破棄されます。モデルを正しくフォーマットする必要はありません。
低温 (0.1 ～ 0.2) — 創造性よりも再現性。
透過的な出力 — 実行するたびに、プロンプト、生のモデル応答、およびデバッグ用に抽出されたコードが返されます。
現在の RustEngineer エージェントは、 localhost:8080 で llama.cpp 経由で unsloth/Qwen3.5-0.8B-GGUF を実行します。 RUST_ENGINEER_MODEL / LLAMA_CPP_BASE_URL でオーバーライドします。
nocodo には、クラウドまたは中規模モデルと連携するように設計された上位エージェントの調整層も含まれています。これらは製品と計画の側面を処理し、平易な言語の説明を、専門のコーディング エージェントが実行する構造化された叙事詩とタスクに変換します。
2 フェーズの PO → PM チャット フロー、タスク ステート マシン、スペシャリストの派遣、コメント、コード抽出はすべて機能します。調整層はコーディング エージェントよりも成熟度が進んでいます。現在の焦点は、そのギャップを埋めることです。
仕組み (完全なビジョン)
ステップ 1 — 製品所有者に相談します。何を構築したいかを説明します。 PO は、ユーザー、データ、機能、制約について質問します。構造化された選択ウィジェット (ラジオ、チェックボックス) を維持

コンテキストがきれいになります。
ステップ 2 — PM が計画を作成します。 PO はプロジェクト マネージャーに引き継ぎ、プロジェクト マネージャーは、適切な専門家にタスクが割り当てられた構造化された叙事詩を作成します。
ステップ 3 — スペシャリストが実行します。タスクはコーディング エージェントに流れます。エンジニアリング マネージャーは、専門家に連絡する前に技術的なタスクをレビューします。ステートマシン: ドラフト → ニーズ技術シェーピング → 準備完了 → 進行中 → 完了。
ステップ 4 — プラットフォームが残りを処理します。ツールのインストール、コードのコミット、アプリの構築とデプロイはすべて nocodo によって管理され、すべてお客様が制御するインフラストラクチャ上で行われます。
nocodo は、望むものを平易な言葉で説明できる人々によって、完全にローカル モデルを使用して、デスクトップ アプリから中小企業向け CRUD アプリを構築、デプロイ、管理できるようになります。認証と権限、パイプラインの構築と展開、VPS またはマネージド クラウド - 製品としてのフルスタック開発チーム。
これは、サブギガバイト モデルで実行されるプロジェクトにとって野心的な主張です。私たちはそれに向けて一度に 1 つのエージェント モードを構築しています。
デスクトップ — 管理 UI をラップする Tauri
管理 UI — TypeScript + SolidJS + TailwindCSS + DaisyUI
共有型 — TypeScript codegen を使用した Rust 型 (手書きの API コントラクトなし)
エージェント ランタイム — SQLite 永続セッションを備えたマルチエージェント システム
コード エージェント — llama.cpp、ローカル プライベート コード生成用のモデル ≤ 1B
ローカル開発セットアップ、バックエンドと UI の実行、およびデプロイ ワークフローについては、DEVELOP.md を参照してください。
積極的な開発。 RustEngineer (ディーゼル モデルの関数、構造体、スキーマ) は、現在動作している小さなモデル コード生成の例です。 PO→PMの調整フローは機能しています。 Engineering Manager、リアル認証、および追加のコード エージェント モードがロードマップにあります。
貢献、フィードバック、問題は github.com/brainless/nocodo で受け付けています。
を中心に構築されたコーディング エージェント

Rust (Actix Web、Diesel)/Typescript (Solid、Tailwind) 独自のスタックを使用したローカル開発用の mall (<10B) モデル。コーディングの知識は必要ありません。複数のエージェントと役割 - プロダクト オーナー、PM、エンジニアリング マネージャー、Rust エンジニア、SolidJS エンジニアなどが連携してフルスタック アプリを構築します。
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

Coding agent built around small (<10B) models for local development with a Rust (Actix Web, Diesel)/Typescript (Solid, Tailwind) opinionated stack. No coding knowledge needed. Multiple agents and roles - Product Owner, PM, Engineering Manager, Rust Engineer, SolidJS Engineer, etc. work together to b
[truncated]

GitHub - brainless/nocodo: Coding agent built around small (<10B) models for local development with a Rust (Actix Web, Diesel)/Typescript (Solid, Tailwind) opinionated stack. No coding knowledge needed. Multiple agents and roles - Product Owner, PM, Engineering Manager, Rust Engineer, SolidJS Engineer, etc. work together to build full-stack apps. · GitHub
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
brainless
/
nocodo
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
141 Commits 141 Commits .githooks .githooks admin-gui admin-gui agents agents backend backend gui gui schema-codegen schema-codegen scripts scripts shared-types shared-types tauri tauri website website .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml DEVELOP.md DEVELOP.md LICENSE LICENSE MIXTURE_OF_AGENTS.md MIXTURE_OF_AGENTS.md README.md README.md project.toml.template project.toml.template View all files Repository files navigation
Experimental · Work in Progress
A coding agent for small (<10B) local models, targeting a specific opinionated stack. Goals not yet fully achieved — but the early results are promising enough to keep going.
Coding agent built around small (<10B) — and in some cases tiny (<1B) — models for local development. The target stack is opinionated: Rust (Actix Web + Diesel) on the backend, TypeScript (SolidJS + Tailwind) on the frontend . Multiple agent roles — Product Owner, Project Manager, Engineering Manager, Rust Engineer, SolidJS Engineer, and more — will eventually collaborate to build full-stack apps. No coding knowledge required.
This will not produce a general-purpose coding agent. The bet is narrower and more achievable: nocodo can get small-business CRUD apps built and managed entirely from a desktop app, using local LLMs, for anyone who can describe what they want.
General coding agents need large, expensive models because they face open-ended problems. nocodo takes the opposite approach: make each problem maximally narrow, then let a tiny model solve it reliably.
The stack is fixed. The patterns are known. A Diesel model function, an Actix route handler, a SolidJS form component — each has a small, learnable shape. Instead of asking a model to figure things out, nocodo builds a dedicated agent mode for each discrete coding task, loads it with concrete examples, and asks for exactly one thing. The model copies a pattern. Deterministic post-processing cleans up the output.
The result: Qwen3.5-0.8B — a model that fits in under a gigabyte — generating correct Diesel ORM functions. That's the premise in action. It works because of the constraint, not despite it.
nocodo builds agent/mode coverage for each layer of the stack, one layer at a time. Each mode is a self-contained prompt builder: it extracts the right code context, constructs an example-rich single-shot prompt, and deterministically post-processes the output.
Layer
Modes
Status
Diesel ORM
Model impl functions, model structs, schema definitions
✅ Working
Actix Web
Controllers (handlers), routers, middleware
Planned
Auth & permissions
Session, JWT, role-based access
Planned
TypeScript frontend — next
Layer
Modes
Status
SolidJS
Routing, state/context, forms, views
Planned
Tailwind + DaisyUI
Component styling
Planned
Platform layer — on the roadmap
Capability
Description
Status
Tooling management
Install and configure Rust compiler, Node, and project dependencies — deterministic, no LLM involved
Planned
Git integration
GitHub/GitLab support for branches, commits, PRs — deterministic code, agent-managed
Planned
Deployment
VPS/SSH builds, or managed hosts (Render, Railway, Fly, Cloudflare) — fully API-integrated, same agent/mode pattern
Planned
Once a few modes in a layer work reliably, adding new ones accelerates — because nocodo's own coding agents participate in building them.
Most coding agents route around small models by using GPT-4 or Claude. nocodo deliberately constrains the problem so smaller models become viable:
Single-shot — one prompt, one response, parse the output. No tool loops, no multi-turn reasoning.
Example-driven — prompts contain 10–15 concrete patterns the model can copy, not instructions to reason from first principles.
Narrow scope per mode — each mode generates exactly one artifact (one function, one struct, one route).
Deterministic post-processing — imports are stripped and re-injected correctly, code fences are unwrapped, <think> blocks are discarded. The model doesn't need to get formatting right.
Low temperature (0.1–0.2) — reproducibility over creativity.
Transparent output — every run returns the prompt, raw model response, and extracted code for debugging.
The current RustEngineer agent runs unsloth/Qwen3.5-0.8B-GGUF via llama.cpp at localhost:8080 . Override with RUST_ENGINEER_MODEL / LLAMA_CPP_BASE_URL .
nocodo also includes a coordination layer of higher-level agents designed to work with cloud or mid-size models. These handle the product and planning side — turning a plain-language description into structured epics and tasks that specialist coding agents execute.
The two-phase PO → PM chat flow, task state machine, specialist dispatch, comments, and code extraction are all functional. The coordination layer is ahead of the coding agents in maturity; the current focus is closing that gap.
How it works (the full vision)
Step 1 — Talk to the Product Owner. Describe what you want to build. The PO asks about users, data, features, and constraints. Structured choice widgets (radio, checkboxes) keep the context clean.
Step 2 — PM creates the plan. The PO hands off to the Project Manager, who produces a structured epic with tasks assigned to the right specialists.
Step 3 — Specialists execute. Tasks flow to coding agents. The Engineering Manager reviews technical tasks before they reach specialists. State machine: draft → needs_technical_shaping → ready → in_progress → done .
Step 4 — Platform handles the rest. Tooling is installed, code is committed, the app is built and deployed — all managed by nocodo, all on infrastructure you control.
nocodo will get small-business CRUD apps built, deployed, and managed from a desktop app, using entirely local models, by people who can describe what they want in plain language. Auth and permissions, build and deployment pipelines, VPS or managed cloud — a full-stack development team as a product.
That's an ambitious claim for a project running on sub-gigabyte models. We're building toward it one agent mode at a time.
Desktop — Tauri wrapping the admin UI
Admin UI — TypeScript + SolidJS + TailwindCSS + DaisyUI
Shared types — Rust types with TypeScript codegen (no handwritten API contracts)
Agent runtime — multi-agent system with SQLite-persisted sessions
Code agents — llama.cpp, models ≤ 1B for local private code generation
See DEVELOP.md for local development setup, running the backend and UI, and the deploy workflow.
Active development. The RustEngineer — Diesel model functions, structs, and schema — is the current working example of tiny-model code generation. The PO → PM coordination flow is functional. Engineering Manager, real auth, and additional code agent modes are on the roadmap.
Contributions, feedback, and issues welcome at github.com/brainless/nocodo .
Coding agent built around small (<10B) models for local development with a Rust (Actix Web, Diesel)/Typescript (Solid, Tailwind) opinionated stack. No coding knowledge needed. Multiple agents and roles - Product Owner, PM, Engineering Manager, Rust Engineer, SolidJS Engineer, etc. work together to build full-stack apps.
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
