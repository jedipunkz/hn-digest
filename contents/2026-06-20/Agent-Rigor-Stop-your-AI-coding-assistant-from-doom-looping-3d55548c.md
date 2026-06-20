---
source: "https://github.com/MeherBhaskar/agent-rigor"
hn_url: "https://news.ycombinator.com/item?id=48611458"
title: "Agent Rigor – Stop your AI coding assistant from doom-looping"
article_title: "GitHub - MeherBhaskar/agent-rigor: A self-sufficient operating system for autonomous coding agents. · GitHub"
author: "meherbhaskar"
captured_at: "2026-06-20T18:35:39Z"
capture_tool: "hn-digest"
hn_id: 48611458
score: 1
comments: 0
posted_at: "2026-06-20T18:10:50Z"
tags:
  - hacker-news
  - translated
---

# Agent Rigor – Stop your AI coding assistant from doom-looping

- HN: [48611458](https://news.ycombinator.com/item?id=48611458)
- Source: [github.com](https://github.com/MeherBhaskar/agent-rigor)
- Score: 1
- Comments: 0
- Posted: 2026-06-20T18:10:50Z

## Translation

タイトル: Agent Rigor – AI コーディング アシスタントの破滅的なループを阻止する
記事のタイトル: GitHub - MeherBhaskar/agent-rigor: 自律コーディング エージェントのための自己完結型オペレーティング システム。 · GitHub
説明: 自律コーディング エージェント用の自立したオペレーティング システム。 - MeherBhaskar/エージェント厳格

記事本文:
GitHub - MeherBhaskar/agent-rigor: 自律コーディング エージェント用の自立型オペレーティング システム。 · GitHub
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
メヘルバスカール
/
エージェントの厳格さ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
8 コミット 8 コミット アセット アセットの例 例 スキル スキル テンプレート テンプレート .gitignore .gitignore CHEATSHEET.md CHEATSHEET.md CONTEXT_MANAGEMENT.md CONTEXT_MANAGEMENT.md CONTRIBUTING.md CONTRIBUTING.md QUICKSTART.md QUICKSTART.md README.md README.md SYSTEM_CORE.md SYSTEM_CORE.md install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング アシスタントのための厳格な経験的規律。
AI エージェントが自分自身をコード化して追い詰めていくのを眺めるのはやめましょう。規律を与えてください。
クイックスタート •
哲学 •
アーキテクチャ •
サポートされているエージェント
ほとんどの AI コーディング エージェントが失敗するのは、知性が欠けているからではなく、規律が欠けているからです。自分のデバイスに任せると、次のようなことが起こります。
❌ 計画をスキップして、すぐに実装に進みます。
❌ 実際には機能しない、一見もっともらしいコードを作成する。
❌ 「破滅ループ」（順方向スパイラル）に閉じ込められる。
❌ セッション間で学んだことを忘れてしまう（文脈健忘症）。
❌ 一度にロードする命令が多すぎるため、「コンテキストの腐敗」が発生します。
エージェント リガーがこれを解決します。これは、構造化された多層の漸進的開示フレームワーク、つまり一連の必須プロトコル、検証ゲート、あらゆる段階で経験的規律を強制する反合理化保護策を提供します。
実用的なプロトコル : すべての指示は、エッセイではなく、終了基準を備えた検証可能なステップです。
経験的主権: 主張には証拠が必要です。 「正しいと思われる」だけでは決して十分ではありません。
アトミックな状態遷移: コードベースは既知の正常な状態の間を移動します。壊れた状態は決してコミットされません。
反合理化 : すべてのスキルは、エージェントが規律を無視するために使用する言い訳を積極的に予測し、反論します。
Progressive Disclosure : エージェントは現在のフェーズに必要なファイルのみを読み取り、ファイルを保存します。

トークンと指示無視の防止。
このシステムは、プログレッシブ ディスクロージャーを使用して、コンテキスト ウィンドウの崩壊を防ぐ堅牢な 3 層階層に編成されています。
L1: Apex カーネル ( SYSTEM_CORE.md ) : 常時オンのルーティングと交渉不可能な規則。
L2: フェーズ ディレクター ( 00_PHASE_DIRECTOR.md ) : フェーズに入ったときにのみロードされるジャストインタイム オーケストレーション。
L3: スキル プロトコル ( skill/*.md ) : ディレクターによって要求された場合にのみロードされる詳細な実行ガイドライン。
グラフTD
A[フェーズ 1: ミッションの合成] -->|PLAN.md| B(フェーズ2:実行エンジン)
B -->|コミットされたコード| C{フェーズ 3: 検証マトリックス}
C -->|重大な発見| B
C -->|結果ゼロ| D[フェーズ 4: 認知的持続]
D -->|コンテキストスナップショット|あ
サブグラフ フェーズ 6: 適応プロトコル
Z[自己補正/範囲防御/統合]
終わり
B -.->|3 ストライク失敗| Z
Z -.->|回復| B
読み込み中
🛠️ 6 つの運用フェーズ
フェーズ
目的
主要なスキル
01. ミッション合成
要件と計画
要件の抽出、戦略的分解
02. 実行エンジン
実装とテスト
収束反復、状態チェックポイント
03. 検証マトリックス
品質とレビューのゲート
ペンタゴナル監査、エントロピー削減
04. 認知的永続性
記憶と知識
コンテキストのライフサイクル、構造地図作成
05. インターフェースプロトコル
安全な環境での相互作用
有界観察、セマンティックナビゲーション
06. 適応プロトコル
免疫システム
再帰的自己修正、スコープの封じ込め
⚡ クイックスタート
Agent Rigor を 2 分以内にプロジェクトで使用できるようにします。
プロジェクト ルートでインストール スクリプトを実行します。
カール -sSL https://raw.githubusercontent.com/MeherBhaskar/agent-rigor/main/install.sh |バッシュ
(または、このリポジトリのクローンを .agents/ ディレクトリに作成します)。
エージェントに次のように指示するだけです。
「[機能]を構築する必要があります。.agents/SYSTEM_C を読んでください」

ORE.md を作成し、フェーズ 1 (ミッション合成) を開始します。」
エージェントは自動的にフェーズ 1 ディレクターを読み取り、 PLAN.md を作成し、実装、レビュー、コンテキスト保存を通じて自身の作業を調整します。
Agent Rigor は純粋なマークダウンであり、プラットフォームに依存しません。以下のものでネイティブに動作します。
すぐに使用できる構成テンプレートについては、examples/ フォルダーを参照してください。
エージェントをより賢く、より規律あるものにするための貢献を歓迎します。
エージェントが実際に従うスキルを設計する方法を理解するには、貢献ガイドラインを参照してください。
自律コーディング エージェントのための自立したオペレーティング システム。
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

A self-sufficient operating system for autonomous coding agents. - MeherBhaskar/agent-rigor

GitHub - MeherBhaskar/agent-rigor: A self-sufficient operating system for autonomous coding agents. · GitHub
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
MeherBhaskar
/
agent-rigor
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
8 Commits 8 Commits assets assets examples examples skills skills templates templates .gitignore .gitignore CHEATSHEET.md CHEATSHEET.md CONTEXT_MANAGEMENT.md CONTEXT_MANAGEMENT.md CONTRIBUTING.md CONTRIBUTING.md QUICKSTART.md QUICKSTART.md README.md README.md SYSTEM_CORE.md SYSTEM_CORE.md install.sh install.sh View all files Repository files navigation
Strict empirical discipline for your AI coding assistant.
Stop watching your AI agent code itself into a corner. Give it discipline.
Quickstart •
Philosophy •
Architecture •
Supported Agents
Most AI coding agents fail not because they lack intelligence, but because they lack discipline . When left to their own devices, they:
❌ Skip planning and jump straight to implementation.
❌ Write plausible-looking code that doesn't actually work.
❌ Get trapped in "doom loops" (fix-forward spirals).
❌ Forget what they learned between sessions (context amnesia).
❌ Suffer from "context rot" by loading too many instructions at once.
Agent Rigor solves this. It provides a structured, multi-layer progressive disclosure framework: a set of mandatory protocols, verification gates, and anti-rationalization safeguards that force empirical discipline at every step.
Actionable Protocols : Every instruction is a verifiable step with exit criteria, not an essay.
Empirical Sovereignty : Claims require evidence. "Seems right" is never sufficient.
Atomic State Transitions : The codebase moves between known-good states. Broken states are never committed.
Anti-Rationalization : Every skill actively anticipates and rebuts the excuses agents use to skip discipline.
Progressive Disclosure : The agent reads only the files it needs for the current phase, saving tokens and preventing instruction neglect.
The system is organized into a robust 3-tier hierarchy using Progressive Disclosure to prevent context window collapse.
L1: Apex Kernel ( SYSTEM_CORE.md ) : Always-on routing and non-negotiable laws.
L2: Phase Directors ( 00_PHASE_DIRECTOR.md ) : Just-in-time orchestration loaded only when entering a phase.
L3: Skill Protocols ( skills/*.md ) : Deep execution guidelines loaded only when requested by the Director.
graph TD
A[Phase 1: Mission Synthesis] -->|PLAN.md| B(Phase 2: Execution Engine)
B -->|Committed Code| C{Phase 3: Verification Matrix}
C -->|CRITICAL Findings| B
C -->|Zero Findings| D[Phase 4: Cognitive Persistence]
D -->|Context Snapshot| A
subgraph Phase 6: Adaptive Protocols
Z[Self-Correction / Scope Defense / Consolidation]
end
B -.->|3-Strike Failure| Z
Z -.->|Recovery| B
Loading
🛠️ The 6 Operational Phases
Phase
Purpose
Key Skills
01. Mission Synthesis
Requirements & Planning
Requirement Distillation, Strategic Decomposition
02. Execution Engine
Implementation & Testing
Convergent Iteration, State Checkpointing
03. Verification Matrix
Quality & Review Gates
Pentagonal Audit, Entropy Reduction
04. Cognitive Persistence
Memory & Knowledge
Context Lifecycle, Structural Cartography
05. Interface Protocols
Safe Environment Interaction
Bounded Observation, Semantic Navigation
06. Adaptive Protocols
The Immune System
Recursive Self-Correction, Scope Containment
⚡ Quickstart
Get Agent Rigor working in your project in under 2 minutes.
Run the installation script in your project root:
curl -sSL https://raw.githubusercontent.com/MeherBhaskar/agent-rigor/main/install.sh | bash
(Alternatively, clone this repo into an .agents/ directory).
Simply prompt your agent with:
"I need to build [feature]. Read .agents/SYSTEM_CORE.md and begin Phase 1 (Mission Synthesis)."
The agent will automatically read the Phase 1 Director, create a PLAN.md , and orchestrate its own work through implementation, review, and context saving.
Agent Rigor is pure markdown and platform-agnostic . It works natively with:
See the examples/ folder for ready-to-use configuration templates.
We welcome contributions to make agents smarter and more disciplined!
Please see our Contributing Guidelines to understand how to design skills that agents actually follow.
A self-sufficient operating system for autonomous coding agents.
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
