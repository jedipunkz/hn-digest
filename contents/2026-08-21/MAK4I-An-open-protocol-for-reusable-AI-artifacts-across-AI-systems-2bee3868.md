---
source: "https://github.com/talvikai/mak4i-protocol"
hn_url: "https://news.ycombinator.com/item?id=49388124"
title: "MAK4I – An open protocol for reusable AI artifacts across AI systems"
article_title: "GitHub - talvikai/mak4i-protocol: Open protocol for reusable AI artifacts, knowledge continuity, and check-before-create workflows across AI tools. · GitHub"
image: "https://opengraph.githubassets.com/5df677ca7a644b00d38fd940fcc57222bdba5cdecd14e7ced7032ed656a5e53b/talvikai/mak4i-protocol"
author: "aswinb"
captured_at: "2026-08-21T14:25:34Z"
capture_tool: "hn-digest"
hn_id: 49388124
score: 2
comments: 0
posted_at: "2026-08-21T13:54:37Z"
tags:
  - hacker-news
  - translated
---

# MAK4I – An open protocol for reusable AI artifacts across AI systems

- HN: [49388124](https://news.ycombinator.com/item?id=49388124)
- Source: [github.com](https://github.com/talvikai/mak4i-protocol)
- Score: 2
- Comments: 0
- Posted: 2026-08-21T13:54:37Z

## Translation

タイトル: MAK4I – AI システム全体で再利用可能な AI アーティファクトのためのオープン プロトコル
記事のタイトル: GitHub - talvikai/mak4i-protocol: 再利用可能な AI 成果物、知識の継続性、AI ツール全体の作成前チェック ワークフローのためのオープン プロトコル。 · GitHub
説明: 再利用可能な AI 成果物、知識の継続性、AI ツール全体での作成前チェックのワークフローのためのオープン プロトコル。 - タルビカイ/mak4i-プロトコル

記事本文:
GitHub - talvikai/mak4i-protocol: 再利用可能な AI 成果物、知識の継続性、AI ツール全体の作成前チェック ワークフローのためのオープン プロトコル。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
タルビカイ
/
mak4i プロトコル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
16 コミット 16 コミット フォルダーとファイル
.github .github アーティファクト アーティファクト ドキュメント ドキュメント f

rameworks/ wd-docx-framework フレームワーク/ wd-docx-framework 標準 標準 .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md ARTIFACT_INDEX.md ARTIFACT_INDEX.md BUSINESS_MODEL.md BUSINESS_MODEL.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COMPETITIVE_LANDSCAPE.md COMPETITIVE_LANDSCAPE.md CONFORMANCE.md CONFORMANCE.md COTRIBUTING.md COTRIBUTING.md DEVELOPER_GUIDE.md DEVELOPER_GUIDE.md GOVERNANCE.md GOVERNANCE.md ライセンス ライセンス MCP.md MCP.md PROBLEM.md PROBLEM.md README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md SPEC.md SPEC.md STANDARDS.md STANDARDS.md VISION.md VISION.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
MAK4I — 知性のための記憶、成果物、知識
AI がすでに構築したものを再構築しないためのプロトコル。
ステータス: フェーズ 0 — 財団。このリポジトリには現在プロトコルが含まれています
設計、仕様草案、および文書化。リファレンス API
実装 (バックエンド、CLI、SDK) はまだ構築されていません。これはフェーズ 2 です。
現在の内容と計画されている内容については、ROADMAP.md を参照してください。
MAK4I は、パッケージ化、識別、バージョン管理、共有、
AI アーティファクトをモデル、プラットフォーム、およびモデル間で注入および再利用する
組織。
MAK4I アーティファクトは以下を表すことができます。
手順的な知識（何かを一貫して行う方法）
アーキテクチャと API 契約
歴史的な決定と理論的根拠
再利用可能な出力 (ドキュメント、コード、テンプレート)
メモリは、プロトコル全体ではなく、アーティファクト タイプの 1 つです。 Git と同じように
標準化されたソース管理と npm 標準化されたパッケージ配布、
MAK4I は、再利用可能な AI アーティファクトがツール間を移動する方法を標準化します。
それぞれを一から再構築しています。
クロード コード、カーソル、ChatGPT、ジェミニ、ベッドロック、副操縦士、内部エージェント
│
MCP / SDK / API
│
MAK4I
│
レジストリ · アーティファクト · ナレッジ · コンテ

xt
詳細については ARCHITECTURE.md を参照してください。
アーキテクチャ (現在実際に実装されているものと計画されているものを含む)。
MAK4I はプロトコルです。 Talvik がプラットフォームを構築します。
MAK4I プロトコル (オープン、MIT ライセンス)
↓
Talvik レジストリ (ホスト型)
タルヴィク エンタープライズ (商用)
Talvik SDK (Python、Node.js、Go、Rust)
Talvik CLI (mak4i のインストール、挿入、公開)
プロトコルを永久にオープンにします。商業エコシステムがその上にあります。
で使用されているオープン プロトコル + 商用エコシステム アプローチに従っています。
Git、Kubernetes、OpenTelemetry などのプロジェクト。
AI のメモリ空間は混雑しています。 Mem0 、 Google の Open Knowledge Format 、 Open Memory Protocol 、およびすべての主要なプラットフォームのネイティブ メモリはすべて、「AI が覚えていない」バージョンの一部を解決します。 MAK4I は、AI のメモリ、知識、相互運用性プロジェクトの既存の状況を評価した後に設計されました。完全な比較については、COMPETITIVE_LANDSCAPE.md を参照してください。
MAK4Iは彼らを上回ろうとしているわけではありません。これは、より狭い別の問題を解決します。
MAK4I は、プロトコルの動作として強制される再利用規律です。何かを生成する前にレジストリをチェックし、存在するものを再利用または適応させ、一致するものがない場合にのみ新規作成します。このチェックはプロトコルの保証であり、クライアントがスキップできるオプションの規則ではありません。アーティファクトが再利用できるかどうかを追跡することは、最初にチェックを必要とするランタイムとは異なるため、この区別が重要です。すべての再利用に関する決定は、ベンチマークの主張ではなく、実際のトークン節約額の見積もりとともに記録されます。これは、実際の再利用を長期にわたって追跡するように設計された実行台帳です (現在、開発時の観察が反映されています。以下の概念実証を参照してください)。
すべての AI ツールは、再利用可能な知識を異なる方法で表現します — Claude は
プロジェクト、アーティファクト、およびスキル。カーソルにはルールがあります。 ChatGPT にはメモリがあります。
GitHub Copilot には手順があります。それらの表現はどれもありません

アーベル
道具の間。ツールを切り替えればゼロからスタートできます。
スタックを再説明し、既存のコードを再生成し、
すでに共有されているコンテキストを再確立します。
それは、計算上、財務上、環境上の無駄です。
1 日あたり 100 万回の AI セッションでは、それぞれ 1,000 トークンが無駄になります —
つまり、回避可能な生成では 1 日あたり 10 億トークンになります。
AI がコンピューティングの従量課金に移行するにつれて、その無駄は
AI を大規模に実行するすべてのビジネスに直接的なドルコストがかかります。
種類
答え
例
手続き的
どうやって？
コードフレームワーク、デプロイメントパイプライン、エンジニアリングプレイブック
セマンティック
何？
システム アーキテクチャ、スキーマ、API コントラクト、ドメイン モデル
エピソード的な
なぜですか？
下された決定、理論的根拠、スプリント履歴、チームの慣例
これらを組み合わせることで、あらゆる AI ツールにわたって完全なプロジェクトの継続性が提供されます。
# メモリパックをインストールする
mak4i インストール会社/バックエンド標準
mak4i schedovia/context をインストールします
# AI セッションの前に挿入
mak4i を注入します
# AI セッションはフルコンテキストで開始されます
# 改めて説明する必要はありません。再生はありません。すぐに続行します。
簡単な例
{
"id" : " schedovia-stack-context " ,
"バージョン" : " 1.0.0 " 、
"タイプ" : "コンテキスト" ,
"レイヤー" : "エピソード" ,
"name" : " スケドビア スタック コンテキスト " ,
"description" : " Schedovia のフルスタック コンテキスト — セッションごとにアーキテクチャを再説明する必要がなくなります " ,
"token_estimate" : 1500 、
"タグ" : [ " スキードビア " 、 " スタック " 、 " コンテキスト " ]
}
概念実証
MAK4I のコア再利用メカニズム — 既存のアーティファクトのチェックと再利用
それらを再生成するのではなく、MAK4I 自身のテスト中に実証されました。
API が存在する前の開発プロセス。以下の初期の数字は反映されています
運用トラフィックではなく、開発時の観察です。
詳細については、docs/MAK4I_SAVINGS_LOG.md を参照してください。
日付のセッションごとの内訳。
WD Technology Solutions は Talvik の設計です
p

フェーズ 2 API が完成したら、artner と最初の製品導入者を対象とします。
生きる。統合計画については、ROADMAP.md を参照してください。
MAK4I は、プロトコル標準として MAK-XXXX 規則を使用します。参照
ROADMAP.md は完全なフェーズごとのロードマップです。
標準リストの VISION.md —
このリポジトリは現在フェーズ 0 — Foundation (2026 年 8 月) です。
知識を一度書きます。どこでも注入。すぐに続行します。
AI セッションに対する MAK4I は、Node.js に対する npm のようなものです。
ポータブルAIメモリ。永遠に開いてください。
© 2026 Talvik, Inc. — MAK4I プロトコルはオープンソースであり、MIT ライセンスを取得しています。
再利用可能な AI 成果物、知識の継続性、AI ツール全体での作成前チェックのワークフローのためのオープン プロトコル。
Readme MIT ライセンスの行動規範
セキュリティ ポリシー アクティビティ カスタム プロパティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Open protocol for reusable AI artifacts, knowledge continuity, and check-before-create workflows across AI tools. - talvikai/mak4i-protocol

GitHub - talvikai/mak4i-protocol: Open protocol for reusable AI artifacts, knowledge continuity, and check-before-create workflows across AI tools. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
talvikai
/
mak4i-protocol
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
16 Commits 16 Commits Folders and files
.github .github artifacts artifacts docs docs frameworks/ wd-docx-framework frameworks/ wd-docx-framework standards standards .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md ARTIFACT_INDEX.md ARTIFACT_INDEX.md BUSINESS_MODEL.md BUSINESS_MODEL.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COMPETITIVE_LANDSCAPE.md COMPETITIVE_LANDSCAPE.md CONFORMANCE.md CONFORMANCE.md CONTRIBUTING.md CONTRIBUTING.md DEVELOPER_GUIDE.md DEVELOPER_GUIDE.md GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE MCP.md MCP.md PROBLEM.md PROBLEM.md README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md SPEC.md SPEC.md STANDARDS.md STANDARDS.md VISION.md VISION.md View all files Repository files navigation
MAK4I — Memory, Artifacts & Knowledge for Intelligence
The protocol for not rebuilding what your AI already built.
Status: Phase 0 — Foundation. This repo currently contains the protocol
design, specification draft, and documentation. The reference API
implementation (backend, CLI, SDK) has not been built yet — that's Phase 2.
See ROADMAP.md for what exists today versus what's planned.
MAK4I is an open protocol for packaging, identifying, versioning, sharing,
injecting, and reusing AI artifacts across models, platforms, and
organizations.
A MAK4I artifact can represent:
Procedural knowledge (how to do something, consistently)
Architecture and API contracts
Historical decisions and rationale
Reusable outputs (documents, code, templates)
Memory is one artifact type — not the entire protocol. Just as Git
standardized source control and npm standardized package distribution,
MAK4I standardizes how reusable AI artifacts move between tools instead of
being rebuilt from scratch in each one.
Claude Code, Cursor, ChatGPT, Gemini, Bedrock, Copilot, internal agents
│
MCP / SDK / API
│
MAK4I
│
Registry · Artifacts · Knowledge · Context
See ARCHITECTURE.md for the full
architecture, including what's actually implemented today versus planned.
MAK4I is the protocol. Talvik builds the platform.
MAK4I Protocol (open, MIT licensed)
↓
Talvik Registry (hosted)
Talvik Enterprise (commercial)
Talvik SDK (Python, Node.js, Go, Rust)
Talvik CLI (mak4i install, inject, publish)
Open protocol forever. Commercial ecosystem on top.
Following the open protocol + commercial ecosystem approach used by
projects such as Git, Kubernetes, and OpenTelemetry.
The AI memory space is crowded — Mem0 , Google's Open Knowledge Format , Open Memory Protocol , and every major platform's native memory all solve some version of "the AI doesn't remember." MAK4I was designed after evaluating the existing landscape of AI memory, knowledge, and interoperability projects. See COMPETITIVE_LANDSCAPE.md for the full comparison.
MAK4I isn't trying to out-remember them. It solves a narrower, different problem:
MAK4I is a reuse discipline , enforced as a protocol behavior: check the registry before generating anything, reuse or adapt what exists, and only create new when nothing matches. That check is a protocol guarantee, not an optional convention a client can skip — the distinction that matters, since tracking that an artifact could be reused is different from a runtime that requires checking first. Every reuse decision is logged with a real token-savings estimate — not a benchmark claim, a running ledger designed to track actual reuse over time (currently reflecting development-time observations — see Proof of Concept below).
Every AI tool represents reusable knowledge differently — Claude has
Projects, Artifacts, and Skills; Cursor has Rules; ChatGPT has Memory;
GitHub Copilot has Instructions. None of those representations travel
between tools. Switch tools and you start from zero.
You re-explain your stack, regenerate code that already exists,
re-establish context that was already shared.
That's waste — computational, financial, and environmental.
At 1 million AI sessions per day each wasting 1,000 tokens —
that is 1 billion tokens per day in avoidable generation.
As AI moves toward metered compute billing, that waste becomes
a direct dollar cost for every business running AI at scale.
Type
Answers
Examples
Procedural
How?
Code frameworks, deployment pipelines, engineering playbooks
Semantic
What?
System architecture, schemas, API contracts, domain models
Episodic
Why?
Decisions made, rationale, sprint history, team conventions
Together they provide complete project continuity across any AI tool.
# Install memory packs
mak4i install company/backend-standards
mak4i install schedovia/context
# Inject before any AI session
mak4i inject
# AI session starts with full context
# No re-explaining. No regenerating. Continue instantly.
Quick Example
{
"id" : " schedovia-stack-context " ,
"version" : " 1.0.0 " ,
"type" : " context " ,
"layer" : " episodic " ,
"name" : " Schedovia Stack Context " ,
"description" : " Full stack context for Schedovia — eliminates re-explaining architecture each session " ,
"token_estimate" : 1500 ,
"tags" : [ " schedovia " , " stack " , " context " ]
}
Proof of Concept
MAK4I's core reuse mechanism — checking for and reusing existing artifacts
instead of regenerating them — was demonstrated during MAK4I's own
development process, before any API existed. Early figures below reflect
that development-time observation, not production traffic.
See docs/MAK4I_SAVINGS_LOG.md for the full,
dated session-by-session breakdown.
WD Technology Solutions is Talvik's design
partner and the intended first production adopter, once the Phase 2 API is
live. See ROADMAP.md for the integration plan.
MAK4I uses the MAK-XXXX convention for protocol standards. See
ROADMAP.md for the full phase-by-phase roadmap and
VISION.md for the standards list —
this repo is currently Phase 0 — Foundation (Aug 2026).
Write knowledge once. Inject anywhere. Continue instantly.
MAK4I is to AI sessions what npm is to Node.js.
Portable AI Memory. Open Forever.
© 2026 Talvik, Inc. — MAK4I Protocol is open source, MIT licensed.
Open protocol for reusable AI artifacts, knowledge continuity, and check-before-create workflows across AI tools.
Readme MIT license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
