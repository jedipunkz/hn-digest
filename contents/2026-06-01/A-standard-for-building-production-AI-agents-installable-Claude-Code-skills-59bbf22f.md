---
source: "https://github.com/AlexDuchDev/agentic-product-standard"
hn_url: "https://news.ycombinator.com/item?id=48343169"
title: "A standard for building production AI agents (+ installable Claude Code skills)"
article_title: "GitHub - AlexDuchDev/agentic-product-standard: The canonical standard for building production-grade agentic products — autonomy ladder, composition patterns, the 7-layer harness, eval discipline — plus a Claude Code skill set that operationalizes it. · GitHub"
author: "AlexDuch"
captured_at: "2026-06-01T00:35:22Z"
capture_tool: "hn-digest"
hn_id: 48343169
score: 2
comments: 0
posted_at: "2026-05-31T05:00:23Z"
tags:
  - hacker-news
  - translated
---

# A standard for building production AI agents (+ installable Claude Code skills)

- HN: [48343169](https://news.ycombinator.com/item?id=48343169)
- Source: [github.com](https://github.com/AlexDuchDev/agentic-product-standard)
- Score: 2
- Comments: 0
- Posted: 2026-05-31T05:00:23Z

## Translation

タイトル: 本番 AI エージェントを構築するための標準 (+ インストール可能なクロード コード スキル)
記事のタイトル: GitHub - AlexDuchDev/agentic-product-standard: 実稼働グレードのエージェント製品を構築するための標準標準 (自律性ラダー、構成パターン、7 層ハーネス、評価規律) に加えて、それを運用化するクロード コード スキル セット。 · GitHub
説明: 実稼働グレードのエージェント製品を構築するための標準標準 (自律性ラダー、構成パターン、7 層ハーネス、評価規律) に加えて、それを運用するためのクロード コード スキル セット。 - AlexDuchDev/agentic-product-standard

記事本文:
GitHub - AlexDuchDev/agentic-product-standard: 実稼働グレードのエージェント製品を構築するための標準標準 (自律性ラダー、構成パターン、7 層ハーネス、評価規律) に加えて、それを運用化するクロード コード スキル セット。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
別のタブまたはウィンドウでアカウントを切り替えました。 R

eload を使用してセッションを更新します。
アラートを閉じる
{{ メッセージ }}
アレックスダッチデヴ
/
エージェント製品標準
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
AlexDuchDev/agentic-product-standard
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
9 コミット 9 コミット .github .github 例 例 スキル スキル テンプレート テンプレート .gitignore .gitignore AGENT_STANDARD.md AGENT_STANDARD.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md CONTRIBUTING.mdライセンス ライセンス README.md README.md STANDARD.md STANDARD.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
実稼働グレードのエージェント製品を構築するための標準標準と、それを運用するための Claude Code スキル セット。
Anthropic、OpenAI、Cognition、Sierra、LangChain、および主要な実践者の制作実践から抽出されたもの (2024 ～ 2026 年)。
標準を読む → · スキルのインストール → · 参照実装 ↗ · 意思決定チェックリスト →
エージェント製品は「AIを搭載した製品」ではありません。
これは、プロセスの一部が、明示的な信頼境界を持つ決定論的アーキテクチャ内の LLM によって動的に指示される製品です。
ほとんどのチームはエージェントのデモを出荷します。生産との接触を生き延びる船舶代理店はほとんどありません。違いがモデルにあることはほとんどありません。それは、アーキテクチャ、ハーネス、およびその周りの評価規律です。このリポジトリは、その作業の現場でテストされた標準に加えて、それをエディタに組み込むための一連のクロード コード スキルです。
実稼働準備状況 — 完了の定義
5 つの原則は、研究所と主要な実践者の生産実践全体にわたって独立して集約されました。これらは、この標準におけるあらゆる決定の核心です。
#
原則
それが何を意味するか
1
デによる決定論

過失、必然的に代理店となる
あらゆる程度の自律性は、事前に与えられるものではなく、獲得する必要があります。
2
アーキテクチャがフレームワークに勝つ
パターンはライブラリよりも存続します。
3
ハーネス > モデル
信頼性の 98% は LLM 周りのコードにあります。
4
コンテキストエンジニアリングが中心的な分野です
コンテキスト ウィンドウに何が入るかによってすべてが決まります。
5
評価主導の開発は交渉の余地がありません
測定しない→改善しない。トレースレビューなし→理解不能。
最も重要なルールは 1 つあります。アーキテクチャは、モデルが改善されたときに残るものです。モデルは変数であり、ハーネスは定数です。比例して投資します。
エージェント製品標準/
§── STANDARD.md ← 製品レベルのカノン — マルチエージェント製品の設計
§── AGENT_STANDARD.md ← 単一エージェントの運用標準 — エージェントを 1 つ構築する
§── templates/ ← artifact-contracts (コントラクト、スキーマ、エンベロープ、評価) をコピーアンドペーストします。
§── 例/ ← 動作例: AgenticMind リファレンス実装
└── スキル/
§── Agentic-product-architect/ ← スキル: Agentic 製品の設計と出荷 (マスタールーター)
│ §── SKILL.md ← マスター: ルーター + 哲学
│ §── Architecture-design/ ← 自律ラダー、5 パターン、シングル vs マルチ
│ §── context-engineering/ ← 書き込み/選択/圧縮/分離、40% ルール
│ §── harness-engineering/ ← LLM ループの周囲の 7 層
│ §── tools-design-mcp/ ← MCP-first、<20 ツール、RAG-MCP、サンドボックス
│ §── メモリアーキテクチャ/ ← Mem0 / Zep / Letta / LangMem / ファイル
│ §── 持続的実行/ ← 時間的ワークフロー + アクティビティパターン
│ §── eval-driven-dev/ ← フサイン/シャンカール ピラミッド + ジャッジ キャリブレーション
│ §── フレームワーク選択/ ← LangGraph / Claude SDK / OpenAI SDK / その他
│ §── 本番対応/

← 監査完了の 12 項目の定義
│ └── antipatterns-review/ ← 12 の既知の障害モードによるコード レビュー
└── Agent-Builder/ ← スキル: 実稼働グレードのエージェントを 1 つ構築する
2 つのトラック、1 つの標準:
エージェントを 1 つ構築します → AGENT_STANDARD.md を読み取り、 template/ を入力し、エージェント ビルダー スキルで実行します。
製品を設計する → STANDARD.md を読み、agentic-product-architect スキル (マルチエージェント、オーケストレーション、フレームワークの選択) を使用して製品を推進します。
ドキュメントは参考資料です。スキルは練習です。設計、構築、レビュー中に適切なガイダンスが自動的に読み込まれます。どちらのスキルも同じ 10 個のサブスキルを共有します。
AgenticMind は、この標準の主力リファレンス実装です。MCP (Apache-2.0、Postgres + pgvector、ヘッドレス Bun) を介したエージェント向けの、監査可能な自己改善型のナレッジおよびメモリ レイヤーです。標準のメモリ アーキテクチャ、コンテキスト エンジニアリング、ツール設計/MCP、評価キャリブレーション、永続的実行、および可観測性の各層が、実行中のテスト済みコードに組み込まれます。
→ レイヤーごとの AgenticMind ケーススタディ 。
スキルはクロード コードと連動します。両方のトラックをインストールします (サブスキルを共有します)。
最速 — 1 つのコマンド (skills.sh 経由):
npx スキル追加 AlexDuchDev/agentic-product-standard
または、ユーザーレベルで手動でコピーします (すべてのプロジェクトで利用可能)。
git clone https://github.com/AlexDuchDev/agentic-product-standard.git
cp -R エージェント製品標準/スキル/ * ~ /.claude/スキル/
プロジェクトレベル (1 つのリポジトリにスコープ):
mkdir -p .claude/スキル
cp -R /path/to/agentic-product-standard/skills/ * .claude/skills/
Claude Code は、各 SKILL.md とその YAML フロントマターを通じてスキルを検出します。インストール後: 1 つのエージェントの構築、実装、またはレビューを開始すると、エージェント ビルダーがトリガーされます。 Agentic-product-architect マルチエージェント製品、エージェント ループ、またはその他のトリガー

主要なフレームワーク (LangGraph、CrewAI、OpenAI Agents SDK、Claude Agent SDK、Pydantic AI、AutoGen)。焦点を絞った質問をする — 「Mem0 それとも Zep?」 、「コンテキストをどのように構成すればよいですか?」 、「エージェント コードを確認してください」 - 関連するサブスキルが直接ロードされます。
「エージェントの構築」から始めないでください。 「このタスクに最低限必要な自律性は何ですか?」から始めます。これを誤った場合の代償は非対称です。
エスカレーション ルール: 厳選された評価セットで L が 90% 以上の合格率を達成するまで、L+1 に上げないでください。
フレームワークに到達する前に、レゴのようなこれらのプリミティブからエージェント製品を作成します。
プロンプトチェーン — 順次分解 (アウトライン → ドラフト → ポリッシュ)
ルーティング — スペシャリストへの分類子 + ディスパッチャー
並列化 — 独立したサブタスクのファンアウト + 集約
Orchestrator-Workers — 中央プランナー + 動的ワーカー
Evaluator-Optimizer — 承認されるまでループする生成者と批評家
メタ原則: まず、これらのパターンを決定論的なコードで構成することによってタスクを解決しようとします。完全なエージェント ループは最後の手段です。
実稼働エージェントでは、ハーネス (LLM ループ周辺のすべて) がコードの 98% を占めます。
┌───────────────────────┐
│ 7. 可観測性とトレース │ ← すべてを記録する
━━━━━━━━━━━━━━━━━━━━━┤
│ 6. 評価層 (CI ゲート) │ ← 回帰をブロック
━━━━━━━━━━━━━━━━━━━━━┤
│ 5. 人間参加型 (通知/質問/レビュー) │ ← 承認ゲート
━━━━━━━

─────────────┤
│ 4. ガードレール (入出力検証) │ ← 多層防御
━━━━━━━━━━━━━━━━━━━━━┤
│ 3. 永続的な実行 (ワークフロー + アクティビティ) │ ← 一時停止/再開/再試行
━━━━━━━━━━━━━━━━━━━━━┤
│ 2. コンテキストとメモリ管理 │ ← 書き込み/選択/圧縮/分離
━━━━━━━━━━━━━━━━━━━━━┤
│ 1. エージェントループ（収集→行動→検証） │ ←「エージェント」そのもの
━━━━━━━━━━━━━━━┘
↕ MCP / 関数呼び出し
アクセス許可の境界は、プロンプトではなくコードによって強制されます。 2025 年の Replit インシデント — プロンプトで明示的に「コード フリーズ」が示されていたにもかかわらず、エージェントが 1,200 社以上の運用データベースを消去した — が標準的な証拠です。モデルは、十分な圧力がかかるとプロンプトレベルの制限を無視します。コードはそうではありません。
アーキテクチャを作成する前にこれを実行してください。これにより、デザインに関する議論の 80% がブロックされます。
□ これを解決する最小の自律レベル (L0 ～ L4) は何ですか?
□エージェントループをフルにしなくても5つのパターンを合成すれば解決できるでしょうか？
□ タスクは幅優先 (並列化可能) ですか? それとも深さ優先 (コヒーレント) ですか?
□ ユーザーの信頼を真っ先に失う3つの失敗モードは何ですか？
□ 許可の境界はどこですか?エージェントがしてはいけないことは何ですか?
□ フレームワークの選択を左右する制約はどれですか?
□ 州はどこにありますか? (コンテキスト内 = アリ

長時間実行用の i パターン)
□ 各段階で誰が出力を検証しますか? (アサーション / LLM 審査員 / 人間によるレビュー)
□ 痕跡はどこに、どのような保持力で存在しますか?
□ 評価セット: サンプルは何個、誰がラベルを付け、どのように成長しますか?
これらの半分にも答えられない場合は、まだコードを書かずに、速度を落として一緒に答えてください。
実稼働準備状況 — 完了の定義
エージェント製品は、12 項目すべてが満たされるまで実稼働準備が整いません。詳細については STANDARD.md を参照してください。
失敗に終わったエージェント プロジェクトを最も早く認識する方法 - スキル セットのアンチパターン レビュー フラグにそれぞれ診断と修正を付けます。
単一エージェントのベースライン前の複数エージェント
生の API を理解する前のフレームワークの抽象化
LLM は人間のラベルに対する校正を行わずに判定します
プロンプトを通じて強制される権限
一般的な eval (「有用性」、「正確さ」)
LLM ジャッジにおけるリッカート スケール (バイナリのみ)
1 つのエージェントで幅と深さの両方を実現
トレース監視を使用しない展開
バージョン管理のないハードコーディングされたプロンプト
単一ベンダーのベンチマークをグラウンド トゥルースとして扱う
運用ベース — 参照ドキュメントではありません。順番に読んでください:
Anthropic — 効果的なエージェントの構築 (Schruntz & Zhang)
OpenAI — ビルディング エージェントの実践ガイド
HumanLayer — 12 因子エージェント (Dex Horthy)
Anthropic — マルチエージェント研究システムをどのように構築したか
認知 — マルチエージェントを構築しないでください (Walden Yan)
LangChain — エージェントのためのコンテキスト エンジニアリング (Lance Martin)
Hamel Husain — AI 製品を迅速に改善するためのフィールド ガイド + AI 製品のニーズの評価
Anthropic — Claude Agent SDK を使用してエージェントを構築する
この標準は進化することを目的としており、分野は急速に変化しています。修正、新しいサンプル、フレームワークの更新、翻訳はすべて歓迎されます。 COTRIBUTING.md および行動規範を参照してください。
建築上の規範 (自律性のはしご、

5パターン、シングルvsマルチ、ハーネス）が安定しています。特定のベンダーやフレームワークのランキングは変化します。これらはまさに私たちが望んでいる種類の PR です。
MIT — 使用し、フォークし、出荷します。
これにより、アーキテクチャに関する議論を 1 週間節約できた場合は、リポジトリにスターを付けてください ⭐ 他の人が見つけられるようにしてください。
v1.0 · 2026 年 5 月時点の運用慣行から組み立てられています
プロダクショングレードのエージェント製品を構築するための標準標準 - 自律性ラダー、構成パターン、7 層ハーネス、評価規律 - に加えて、Claude Code スキル SE

[切り捨てられた]

## Original Extract

The canonical standard for building production-grade agentic products — autonomy ladder, composition patterns, the 7-layer harness, eval discipline — plus a Claude Code skill set that operationalizes it. - AlexDuchDev/agentic-product-standard

GitHub - AlexDuchDev/agentic-product-standard: The canonical standard for building production-grade agentic products — autonomy ladder, composition patterns, the 7-layer harness, eval discipline — plus a Claude Code skill set that operationalizes it. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
AlexDuchDev
/
agentic-product-standard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
AlexDuchDev/agentic-product-standard
main Branches Tags Go to file Code Open more actions menu Folders and files
9 Commits 9 Commits .github .github examples examples skills skills templates templates .gitignore .gitignore AGENT_STANDARD.md AGENT_STANDARD.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md STANDARD.md STANDARD.md View all files Repository files navigation
A canonical standard for building production-grade agentic products — plus a Claude Code skill set that operationalizes it.
Distilled from the production practices of Anthropic, OpenAI, Cognition, Sierra, LangChain, and leading practitioners — 2024–2026.
Read the Standard → · Install the Skills → · Reference Implementation ↗ · Decision Checklist →
An agentic product is not "a product with AI."
It is a product where part of the process is dynamically directed by an LLM within a deterministic architecture with explicit trust boundaries .
Most teams ship agent demos. Few ship agents that survive contact with production. The difference is almost never the model — it's the architecture, the harness, and the eval discipline around it. This repo is the field-tested standard for that work, plus a set of Claude Code skills that put it into your editor.
Production readiness — Definition of Done
Five principles converged independently across the production practices of the labs and the leading practitioners. They are the spine of every decision in this standard:
#
Principle
What it means
1
Determinism by default, agency by necessity
Every degree of autonomy must be earned , not granted upfront.
2
Architecture beats framework
Patterns outlive libraries.
3
Harness > model
98% of reliability lives in the code around the LLM.
4
Context engineering is the core discipline
What enters the context window determines everything.
5
Eval-driven development is non-negotiable
No measurement → no improvement. No trace review → no understanding.
The single most important rule: Architecture is what remains when the model improves. The model is the variable, the harness is the constant. Invest proportionally.
agentic-product-standard/
├── STANDARD.md ← product-level canon — design a multi-agent product
├── AGENT_STANDARD.md ← single-agent operational standard — build one agent
├── templates/ ← copy-paste artifact-contracts (contracts, schemas, envelope, evals)
├── examples/ ← worked example: the AgenticMind reference implementation
└── skills/
├── agentic-product-architect/ ← skill: design & ship agentic PRODUCTS (master router)
│ ├── SKILL.md ← master: router + philosophy
│ ├── architecture-design/ ← autonomy ladder, 5 patterns, single vs multi
│ ├── context-engineering/ ← write/select/compress/isolate, the 40% rule
│ ├── harness-engineering/ ← the 7 layers around the LLM loop
│ ├── tool-design-mcp/ ← MCP-first, <20 tools, RAG-MCP, sandboxing
│ ├── memory-architecture/ ← Mem0 / Zep / Letta / LangMem / files
│ ├── durable-execution/ ← Temporal Workflow + Activity pattern
│ ├── eval-driven-dev/ ← Husain/Shankar pyramid + judge calibration
│ ├── framework-selection/ ← LangGraph / Claude SDK / OpenAI SDK / others
│ ├── production-readiness/ ← 12-point Definition of Done audit
│ └── antipatterns-review/ ← code review through 12 known failure modes
└── agent-builder/ ← skill: build ONE production-grade agent
Two tracks, one standard:
Build one agent → read AGENT_STANDARD.md , fill the templates/ , drive it with the agent-builder skill.
Design a product → read STANDARD.md , drive it with the agentic-product-architect skill (multi-agent, orchestration, framework choice).
The docs are the reference ; the skills are the practice — they auto-load the right guidance while you design, build, and review. Both skills share the same ten sub-skills.
AgenticMind is the flagship reference implementation of this standard — an auditable, self-improving knowledge & memory layer for agents over MCP (Apache-2.0, Postgres + pgvector, headless Bun). It puts the standard's memory-architecture, context-engineering, tool-design/MCP, eval-calibration, durable-execution, and observability layers into running, tested code.
→ Layer-by-layer AgenticMind case study .
The skills work with Claude Code . Install both tracks (they share sub-skills).
Fastest — one command (via skills.sh ):
npx skills add AlexDuchDev/agentic-product-standard
Or copy them in manually — user-level (available in every project):
git clone https://github.com/AlexDuchDev/agentic-product-standard.git
cp -R agentic-product-standard/skills/ * ~ /.claude/skills/
Project-level (scoped to one repo):
mkdir -p .claude/skills
cp -R /path/to/agentic-product-standard/skills/ * .claude/skills/
Claude Code discovers skills via each SKILL.md and its YAML frontmatter. Once installed: agent-builder triggers when you set out to build, implement, or review one agent; agentic-product-architect triggers for multi-agent products , an agent loop, or any major framework (LangGraph, CrewAI, OpenAI Agents SDK, Claude Agent SDK, Pydantic AI, AutoGen). Ask a focused question — "Mem0 or Zep?" , "how should I structure context?" , "review my agent code" — and the relevant sub-skill loads directly.
Never start with "build an agent." Start with "what is the minimum autonomy this task requires?" The cost of getting this wrong is asymmetric.
Escalation rule: do not climb to L+1 until L delivers ≥90% pass rate on a curated eval set.
Compose agentic products from these primitives like Lego — before reaching for a framework.
Prompt Chaining — sequential decomposition (outline → draft → polish)
Routing — classifier + dispatcher to a specialist
Parallelization — fan-out of independent subtasks + aggregation
Orchestrator-Workers — central planner + dynamic workers
Evaluator-Optimizer — generator + critic in a loop until acceptance
Meta-principle: first try to solve the task by composing these patterns in deterministic code. A full agent loop is the last resort.
In a production agent, the harness — everything around the LLM loop — is 98% of the code .
┌─────────────────────────────────────────────┐
│ 7. Observability & Tracing │ ← log EVERYTHING
├─────────────────────────────────────────────┤
│ 6. Evaluation Layer (CI gates) │ ← block regressions
├─────────────────────────────────────────────┤
│ 5. Human-in-the-Loop (notify/ask/review) │ ← approval gates
├─────────────────────────────────────────────┤
│ 4. Guardrails (input/output validation) │ ← defense in depth
├─────────────────────────────────────────────┤
│ 3. Durable Execution (Workflow + Activity) │ ← pause/resume/retry
├─────────────────────────────────────────────┤
│ 2. Context & Memory Management │ ← write/select/compress/isolate
├─────────────────────────────────────────────┤
│ 1. Agent Loop (gather → act → verify) │ ← the "agent" proper
└─────────────────────────────────────────────┘
↕ MCP / function calling
Permission boundaries are enforced by code, never by prompt. The Replit incident of 2025 — an agent wiped a production database for 1,200+ companies despite an explicit "code freeze" in its prompt — is the canonical proof. The model will ignore prompt-level restrictions under enough pressure. Code won't.
Run this before drafting any architecture. It unblocks 80% of design debates.
□ What is the minimum autonomy level (L0–L4) that solves this?
□ Can it be solved by composing the 5 patterns without a full agent loop?
□ Is the task breadth-first (parallelizable) or depth-first (coherent)?
□ What are the 3 failure modes that would lose user trust first?
□ Where are the permission boundaries? What MUST the agent NOT do?
□ Which constraint dominates framework choice?
□ Where does state live? (in-context = anti-pattern for long-running)
□ Who validates outputs at each stage? (assertion / LLM judge / human review)
□ Where do traces live, with what retention?
□ Eval set: how many examples, who labels, how does it grow?
If you can't answer half of these, slow down and answer them together — don't write code yet.
Production readiness — Definition of Done
An agentic product is not production-ready until all 12 are satisfied. Full detail in STANDARD.md .
The fastest way to recognize a doomed agent project — the skill set's antipatterns-review flags each with a diagnostic and a fix.
Multi-agent before a single-agent baseline
Framework abstractions before understanding the raw API
LLM judges without calibration against human labels
Permissions enforced through prompts
Generic evals ("helpfulness," "correctness")
Likert scales in an LLM judge (binary only)
One agent for both breadth and depth
Deploying without trace monitoring
Hardcoded prompts without version control
Treating single-vendor benchmarks as ground truth
The operational base — not reference docs. Read in order:
Anthropic — Building Effective Agents (Schluntz & Zhang)
OpenAI — A Practical Guide to Building Agents
HumanLayer — 12 Factor Agents (Dex Horthy)
Anthropic — How we built our multi-agent research system
Cognition — Don't Build Multi-Agents (Walden Yan)
LangChain — Context Engineering for Agents (Lance Martin)
Hamel Husain — A Field Guide to Rapidly Improving AI Products + Your AI Product Needs Evals
Anthropic — Building agents with the Claude Agent SDK
This standard is meant to evolve — the field moves fast. Corrections, new exemplars, framework updates, and translations are all welcome. See CONTRIBUTING.md and the Code of Conduct .
The architectural canons (the autonomy ladder, the 5 patterns, single-vs-multi, the harness) are stable. Specific vendors and framework rankings will shift — those are exactly the kind of PRs we want.
MIT — use it, fork it, ship with it.
If this saved you a week of architecture debates, star the repo ⭐ so others find it.
v1.0 · assembled from production practices as of May 2026
The canonical standard for building production-grade agentic products — autonomy ladder, composition patterns, the 7-layer harness, eval discipline — plus a Claude Code skill se

[truncated]
