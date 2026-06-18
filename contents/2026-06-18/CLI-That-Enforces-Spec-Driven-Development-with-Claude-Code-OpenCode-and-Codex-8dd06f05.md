---
source: "https://github.com/davidpv/opsx-spec-driven-development-toolkit"
hn_url: "https://news.ycombinator.com/item?id=48582593"
title: "CLI That Enforces Spec-Driven Development with Claude Code, OpenCode, and Codex"
article_title: "GitHub - davidpv/opsx-spec-driven-development-toolkit · GitHub"
author: "davidpv"
captured_at: "2026-06-18T10:35:52Z"
capture_tool: "hn-digest"
hn_id: 48582593
score: 2
comments: 0
posted_at: "2026-06-18T08:43:49Z"
tags:
  - hacker-news
  - translated
---

# CLI That Enforces Spec-Driven Development with Claude Code, OpenCode, and Codex

- HN: [48582593](https://news.ycombinator.com/item?id=48582593)
- Source: [github.com](https://github.com/davidpv/opsx-spec-driven-development-toolkit)
- Score: 2
- Comments: 0
- Posted: 2026-06-18T08:43:49Z

## Translation

タイトル: Claude Code、OpenCode、Codex を使用してスペック駆動開発を強制する CLI
記事のタイトル: GitHub - davidpv/opsx-spec-driven-development-toolkit · GitHub
説明: GitHub でアカウントを作成して、davidpv/opsx-spec-driven-development-toolkit の開発に貢献します。

記事本文:
GitHub - davidpv/opsx-spec-driven-development-toolkit · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
デビッドpv
/
opsx-spec-driven-development-toolkit
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
davidpv/opsx-spec-dr

iven-開発ツールキット
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
25 コミット 25 コミット .github/ workflows .github/ workflows docs docs payload payload scripts scripts src src .gitignore .gitignore README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsup.config.bundled_34r15do6cpo.mjs tsup.config.bundled_34r15do6cpo.mjs tsup.config.ts tsup.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
opsx — 仕様駆動型開発ツールキット
opencode 、 Claude Code 、 Codex のいずれかについて、 OpenSpec の上に仕様駆動型開発の足場を築く、スタックに依存しない CLI。
中心的なアイデア: 仕様 → 計画 → コード。コードは最後に作成される成果物であり、最初のものではありません。 OpenSpec はすべての変更を提案書 + 仕様 + 設計 + タスクとして構造化し、オープンコードは要件まで完全に追跡できる状態で実装を実行します。
モデル: タスク管理、仕様主導
2 つの平面が共存しているため、混同しないでください。管理プレーン (Jira の語彙) は、どのような作業が存在するかを決定し、それを追跡します。ガバナンス プレーン (OpenSpec) がシステムがどのように動作しなければならないかを決定し、後者のみがコードを承認します。
フローチャート LR
サブグラフ MGMT["管理プレーン — 存在する作業 (Jira)"]
LR方向
D["ディスカバリー<br/>バックログ/ディスカバリー/"] --> T["タスク PROJ-123<br/>バックログ/タスク/<br/>目標 + 受け入れ基準"]
終わり
サブグラフ GOV["ガバナンス プレーン — システムはどのように動作する必要があるか (OpenSpec)"]
LR方向
C["変更<br/>提案 + デルタ仕様<br/>+ 設計 + ステップ"] -- "/opsx:archive" --> S["openspec/specs/<br/>真実の情報源"]
さー。 「プロポーズする前にお読みください」 .-> C
終わり
CODE["コード<br/>機能ブランチ、コミット、PR"]
T -- 「タスクはコードを認可しません<br/>変更を促すだけです」 --> C
C -- 「コードへの唯一のゲート」 --> コード
コード - 。 "トレース

戻る: Jira / 変更 / タスク フッター" .-> T
読み込み中
適切に記述されたタスクは何も変更しません。実装は、OpenSpec の変更が存在し、レビューされ、その分岐ゲートが解決されたときにのみ開始されます。タスクに設計または実装の詳細が含まれている場合、その内容は変更に属します。タスクレビュー担当者はそれにフラグを立てます。これは、バックログが Jira を話す場合でも、ワークフローを仕様主導に保つ理由です。Jira がバックログを支配し、仕様がコードを支配します。
opsx init 自体を実行するには、Node.js >= 18 のみが必要です (npx @davidpv/opsx init は他に何もインストールされていなくても動作します。ファイルを書き込むだけです)。スキャフォールディングするワークフローを実際に使用するには、次のものが必要です。
3 つのエージェント CLI すべてが必要なわけではありません。 init 中にターゲットとして選択したものだけが必要です。 opsx ドクターはこれらすべてをチェックし、何が不足しているかを正確に教えてくれます。
# 既存のプロジェクト内
npx @davidpv/opsx init # ターゲットの選択: opencode / Claude Code / Codex、ブランチの構成、Jira キー、言語
npx @davidpv/opsx Doctor # 必要なツールを確認します (openspec CLI、エージェント CLI)
npx @davidpv/opsx update # opsx パッケージのアップグレード後: 新しいコマンド/スキルをプロジェクトにデプロイします
init は、ツールに依存しない共有レイヤー ( workflow.yaml 、 AGENTS.md 管理ブロック、 backlog/ 、 templates/ 、 openspec/ ) と、エージェントごとのネイティブ レイヤーを 1 回書き込みます: .opencode/ (コマンド + スキル + エージェント + opencode.json マージ)、 .claude/ (コマンド + スキル + サブエージェント + settings.json マージ + CLAUDE.md インポート AGENTS.md) )、.codex/skills/ (スキル。Codex にはプロジェクト レベルのスラッシュ コマンドやサブエージェントがないため、スキルとしてコンパイルされたコマンドとレビューアー)。
次に、エージェント内で (ここではオープンコードで示されています):
/start # ガイド付きエントリ: 正しい最初のコマンドに移動します
/req-capture checkout-flow # 0. 要件インタビュー → ディスカバリ ドキュメント
/task-generate checkout-flow # 1. タスク

Jira ID (PROJ-123)
/task-enrich PROJ-123 # エッジケース、シナリオ、推定
/task-jira PROJ-123 # Jira wiki マークアップとしてエクスポート
/opsx:propose 高速検索 # 2. 提案 + デルタ仕様 + 設計 + タスク
/review-change Speed-up-search # 3. ビルド前の監査
/opsx:apply # 4. タスクごとに実装します
/git-commit # タスクに追跡されるセマンティック コミット数
/pr-open # 5. 統合ブランチに対する PR
/ship # 6. 検証 + アーカイブ + マージ
ステージ 0 ～ 1 は、小さな技術的な変更の場合はオプションです。 /opsx:propose 以降はすべて必須です。
フローはガイドされています: /start はエントリ ポイントです。作業の開始方法を尋ね (既存の Jira チケット → /task-import ; チケットなし、直接提案 → /opsx:explore / /opsx:propose ; チケットなし、最初にタスクを作成 → /task-new または /req-capture )、そこにルーティングされます。すべてのコマンドは次のステップを提案することで終了し、/next は、道に迷ったりリポジトリに戻ってきたときに、現在どこにいるのか、何をすべきかを示します (ブランチ、タスクのフロントマター、変更アーティファクト、および PR 状態を検査します)。
フローとオプションの完全なマップ。 🧑 = 人間の対話が必要、🤖 = エージェント (単独で実行)、🧑🤖 = 人間のチェックポイントを備えたエージェント (承認、言語ゲート、またはフォールバック)。
フローチャート TD
classDef 人間の塗りつぶし:#fff3cd、ストローク:#996f00、カラー:#000
classDef エージェントの塗りつぶし:#d6eaf8、ストローク:#1a5276、色:#000
classDef 混合塗りつぶし:#e8daef、ストローク:#6c3483、カラー:#000
classDef アート フィル:#eaeded、ストローク:#7f8c8d、カラー:#000
START(["新作が到着"]) --> ENTRY["/start 🧑<br/>ガイド付きエントリー"]:::human
ENTRY --> CASE{「この作業はどのように始まるのですか?」}
%% 侵入経路
CASE -- "Jira チケットが存在します" --> IMP["/task-import 🧑🤖<br/>チケットを貼り付け + 正規化 + task-reviewer"]:::mixed
IMP --> ST
CASE -- 「チケットなし:<br/>調査/提案」 --> PROP
CASE -- "チケットなし:<br/>最初にタスクを作成" --> SIZE{"I

自発的な作業ですか、それとも単一のタスクですか?"}
SIZE -- "単一タスク" --> TN["/task-new 🧑🤖<br/>ミニインタビュー + ドラフト ID"]:::mixed
TN --> ST
%% 製品フェーズ (オプション)
SIZE -- イニシアチブ --> RC["/req-capture 🧑<br/>インタビュー + 言語ゲート"]:::human
RC --> DISC["バックログ/ディスカバリー/トピック.md"]:::アート
DISC --> SG["/task-generate 🧑🤖<br/>言語ゲート + ユーザーからの Jira ID"]:::mixed
SG --> ST["バックログ/タスク/PROJ-123.md"]:::art
ST --> SE["/task-enrich 🤖<br/>曖昧さのみを尋ねる"]:::agent
SE --> RS["/review-task 🤖"]:::エージェント
RS -- 改訂 --> SE
RS -- 承認 --> JIRAQ{"Jira にエクスポートしますか?"}
JIRAQ -- はい --> SJ["/task-jira 🧑🤖<br/>言語ゲート + 翻訳"]:::mixed
SJ --> JEXP["backlog/exports/jira/"]:::art
JEXP --> PASTE["Jira に貼り付けて実際のキーを確認 🧑"]:::human
JIRAQ -- いいえ --> PROP
貼り付け --> プロップ
%% 仕様フェーズ
PROP["/opsx:propose 🧑🤖<br/>あなたが説明した、エージェントがアーティファクトを書き込む"]:::mixed
PROP --> CHG["openspec/changes/name/"]:::art
CHG --> RV["/review-change 🤖<br/>spec-reviewer + validate --strict"]:::agent
RV -- REVISE --> FIX["アーティファクトを修正 🤖"]:::agent
修正 --> RV
RV -- APPROVE --> BRANCH{"分岐ゲート 🧑<br/>どこに実装しますか? (git.work_mode)<br/>解決されるまでコードをブロックします"}
ブランチ -- 「機能ブランチの作成<br/>(レビュー済み PR)」 --> 適用
ブランチ -- 「直接開発中<br/>(フレキシブル モード、レビューなし)」 --> 適用
%% ビルドフェーズ
APPLY["/opsx:apply 🤖"]:::エージェント
APPLY --> DRIFT{"仕様が間違っているか、ドリフトしていますか?"}
ドリフト -- はい --
[切り捨てられた]
ブランチ ポリシー — main はリリース専用であり、直接作業されることはありません。 workflow.yaml の git.work_mode が残りを決定します。柔軟 (デフォルト) により、統合ブランチに直接実装してコミットできます。その後、/ship は検証、アーカイブ、プッシュのみを行い、PR とレビューをスキップします。機能により、機能ブランチ + PR が必須になります。男

datory ブランチ ゲートは、/opsx:apply がコードを書き込む前に実行されます。最初に作業ブランチが解決 (作成およびチェックアウト) される必要があります。 /git-commit はセーフティ ネットとしてコミット時に再チェックします。機能ブランチは、変更が実際の Jira キー (例: feature/PROJ-123-speed-up-search ) を持つバックログ タスクにリンクされている場合は feature/<task id>-<change> という名前が付けられ、それ以外の場合は feature/<change> という名前が付けられます。
ヒューマン チェックポイントの要約: /start エントリの質問、/req-capture インタビュー、既存の Jira チケットの貼り付け ( /task-import )、すべての言語ゲート (es/en、クライアント向けテキストでは必須)、実装場所の選択 (フィーチャー ブランチと開発)、コミット メッセージの承認、Jira ID の提供と Jira エクスポートの貼り付け、PR コード レビュー、プラットフォーム CLI が存在しない場合の Web UI 経由のマージ。それ以外はすべてエージェント的に実行されます。
トレーサビリティ チェーン: 検出 → タスク (Jira) → 変更 → task.md ステップ → コミット → PR 。各リンクは、タスク フロントマター (変更: )、コミット フッター (変更: / タスク: / Jira: )、PR の説明など、発生する場所に記録されます。タスク ID は Jira キー ( PROJ-123 、ユーザーが提供。問題が存在するまでは PROJ-Dnn ドラフト)。名前に注意してください。タスクはバックログ/Jira 作業項目です。変更内のtasks.mdには実装手順が含まれています。
各変更は、アーカイブされるまで openspec/changes/<name>/ に保存されます。アーカイブすると、そのデルタ仕様がシステムの動作方法の生きた記述である openspec/specs/ にマージされます。 workflow.yaml はブランチ (デフォルトでは git-flow: feature/* → Development )、コミット規約、および Jira エクスポートを構成します。コマンドはそれを読み取るため、パイプラインはプラットフォームに依存しません。
。
§── AGENTS.md # すべてのエージェントが従うべきルール
§── opencode.json # opencode プロジェクトの設定
§── workflow.yaml # パイプライン構成: ブランチ、コミット、Jira (ツールに依存しない)
§── templates/ # Discovery.md、タスク

.md、pr-description.md
§── バックログ/
│ §── Discovery/ # 要件収集メモ (/req-capture)
│ §──tasks/ # Jira ID を持つタスク (/task-generate、/task-enrich)
│ └──exports/ # jira/ (wiki マークアップ) および pr/ (フォールバック PR 記述)
§── .opencode/
│ §── エージェント/
│ │ §── spec-reviewer.md # サブエージェント: 適用/アーカイブ前に変更を監査
│ │ └─ task-reviewer.md # サブエージェント: タスクの監査 (サイジング、テスト容易性)
│ §── コマンド/
│ │ §── start.md # /start — ガイド付きエントリ ポイント (新しい作業のルート)
│ │ §── req-capture.md # /req-capture — 要件インタビュー
│ │ §── task-*.md # /task-import|new|generate|enrich|jira
│ │ §── review-*.md # /review-change, /review-task
│ │ §── git-commit.md # /git-commit — トレーサビリティを備えたセマンティックコミット
│ │ §── pr-open.md # /pr-open — プラットフォームに依存しない PR の作成
│ │ §── ship.md # /ship — 検証 + アーカイブ + マージ
│ │ └─ opsx-*.md # /opsx:* — OpenSpec ライフサイクル (生成)
│ └── スキル/ # OpenSpec ワークフロー スキル (生成)
└── オープンスペック/
§── config.yaml # プロジェクトコンテキスト + アーティファクトごとのルール
§── 仕様/ # 真実の情報源 (現在の動作)
└── 変更/ # 進行中の変更;アーカイブ/履歴を保存する
コマンド
コマンド
ステージ
何をするのか
/開始
エントリー
ガイド付きエントリポイント: 作業の開始方法とルートを尋ねます (チケットが存在する / 直接提案する / 最初にタスクを作成する)
/次へ
どれでも
パイプラインのどの位置にいるかを検出し、

[切り捨てられた]

## Original Extract

Contribute to davidpv/opsx-spec-driven-development-toolkit development by creating an account on GitHub.

GitHub - davidpv/opsx-spec-driven-development-toolkit · GitHub
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
davidpv
/
opsx-spec-driven-development-toolkit
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
davidpv/opsx-spec-driven-development-toolkit
main Branches Tags Go to file Code Open more actions menu Folders and files
25 Commits 25 Commits .github/ workflows .github/ workflows docs docs payload payload scripts scripts src src .gitignore .gitignore README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsup.config.bundled_34r15do6cpo.mjs tsup.config.bundled_34r15do6cpo.mjs tsup.config.ts tsup.config.ts View all files Repository files navigation
opsx — Spec-Driven Development Toolkit
A stack-agnostic CLI that scaffolds spec-driven development on top of OpenSpec , for any of opencode , Claude Code and Codex .
The core idea: Spec → Plan → Code. Code is the last artifact produced, never the first. OpenSpec structures every change as proposal + specs + design + tasks, and opencode executes the implementation with full traceability back to the requirements.
The model: task-managed, spec-driven
Two planes coexist and must not be confused. The management plane (Jira vocabulary) decides what work exists and tracks it ; the governance plane (OpenSpec) decides how the system must behave — and only the latter authorizes code:
flowchart LR
subgraph MGMT["Management plane — what work exists (Jira)"]
direction LR
D["Discovery<br/>backlog/discovery/"] --> T["Task PROJ-123<br/>backlog/tasks/<br/>goal + acceptance criteria"]
end
subgraph GOV["Governance plane — how the system must behave (OpenSpec)"]
direction LR
C["Change<br/>proposal + delta specs<br/>+ design + steps"] -- "/opsx:archive" --> S["openspec/specs/<br/>SOURCE OF TRUTH"]
S -. "read before proposing" .-> C
end
CODE["Code<br/>feature branch, commits, PR"]
T -- "a task authorizes NO code<br/>it only motivates a change" --> C
C -- "the ONLY gate to code" --> CODE
CODE -. "traces back: Jira / Change / Task footers" .-> T
Loading
A task being well written changes nothing: implementation starts only when an OpenSpec change exists, is reviewed, and its branch gate is resolved. If a task ever contains design or implementation detail, that content belongs in the change — task-reviewer flags it. This is what keeps the workflow spec-driven even though the backlog speaks Jira: Jira rules the backlog, the spec rules the code.
To run opsx init itself you only need Node.js >= 18 ( npx @davidpv/opsx init works with nothing else installed — it just writes files). To actually use the workflow it scaffolds, you need:
You don't need all three agent CLIs — only the ones you selected as targets during init . opsx doctor checks all of this and tells you exactly what's missing.
# In your existing project
npx @davidpv/opsx init # pick targets: opencode / Claude Code / Codex, configure branches, Jira key, language
npx @davidpv/opsx doctor # verify required tooling (openspec CLI, agent CLIs)
npx @davidpv/opsx update # after upgrading the opsx package: deploy new commands/skills to your project
init writes the shared, tool-agnostic layer once ( workflow.yaml , AGENTS.md managed block, backlog/ , templates/ , openspec/ ) and a native layer per agent: .opencode/ (commands + skills + agents + opencode.json merge), .claude/ (commands + skills + subagents + settings.json merge + CLAUDE.md importing AGENTS.md ), .codex/skills/ (skills; commands and reviewers compiled as skills, since Codex has no project-level slash commands or subagents).
Then, inside your agent (shown here with opencode):
/start # guided entry: routes you to the right first command
/req-capture checkout-flow # 0. requirements interview → discovery doc
/task-generate checkout-flow # 1. tasks with Jira IDs (PROJ-123)
/task-enrich PROJ-123 # edge cases, scenarios, estimate
/task-jira PROJ-123 # export as Jira wiki markup
/opsx:propose speed-up-search # 2. proposal + delta specs + design + tasks
/review-change speed-up-search # 3. audit before building
/opsx:apply # 4. implement task by task
/git-commit # semantic commits traced to tasks
/pr-open # 5. PR against the integration branch
/ship # 6. validate + archive + merge
Stages 0–1 are optional for small technical changes; everything from /opsx:propose on is mandatory.
The flow is guided : /start is the entry point — it asks how the work begins (existing Jira ticket → /task-import ; no ticket, propose directly → /opsx:explore / /opsx:propose ; no ticket, create the task first → /task-new or /req-capture ) and routes you there. Every command ends by suggesting the next step, and /next tells you where you are and what to do whenever you're lost or coming back to the repo (it inspects branch, task frontmatter, change artifacts, and PR state).
Full map of flows and options. 🧑 = needs human interaction, 🤖 = agentic (runs on its own), 🧑🤖 = agentic with human checkpoint (approval, language gate, or fallback).
flowchart TD
classDef human fill:#fff3cd,stroke:#996f00,color:#000
classDef agent fill:#d6eaf8,stroke:#1a5276,color:#000
classDef mixed fill:#e8daef,stroke:#6c3483,color:#000
classDef art fill:#eaeded,stroke:#7f8c8d,color:#000
START(["New work arrives"]) --> ENTRY["/start 🧑<br/>guided entry"]:::human
ENTRY --> CASE{"How does this work start?"}
%% Entry routes
CASE -- "Jira ticket exists" --> IMP["/task-import 🧑🤖<br/>paste ticket + normalize + task-reviewer"]:::mixed
IMP --> ST
CASE -- "no ticket:<br/>investigate / propose" --> PROP
CASE -- "no ticket:<br/>create task first" --> SIZE{"Initiative or single task?"}
SIZE -- "single task" --> TN["/task-new 🧑🤖<br/>mini interview + draft ID"]:::mixed
TN --> ST
%% Product phase (optional)
SIZE -- initiative --> RC["/req-capture 🧑<br/>interview + language gate"]:::human
RC --> DISC["backlog/discovery/topic.md"]:::art
DISC --> SG["/task-generate 🧑🤖<br/>language gate + Jira IDs from user"]:::mixed
SG --> ST["backlog/tasks/PROJ-123.md"]:::art
ST --> SE["/task-enrich 🤖<br/>asks only on ambiguity"]:::agent
SE --> RS["/review-task 🤖"]:::agent
RS -- REVISE --> SE
RS -- APPROVE --> JIRAQ{"Export to Jira?"}
JIRAQ -- yes --> SJ["/task-jira 🧑🤖<br/>language gate + translation"]:::mixed
SJ --> JEXP["backlog/exports/jira/"]:::art
JEXP --> PASTE["paste into Jira & confirm real key 🧑"]:::human
JIRAQ -- no --> PROP
PASTE --> PROP
%% Spec phase
PROP["/opsx:propose 🧑🤖<br/>you describe, agent writes artifacts"]:::mixed
PROP --> CHG["openspec/changes/name/"]:::art
CHG --> RV["/review-change 🤖<br/>spec-reviewer + validate --strict"]:::agent
RV -- REVISE --> FIX["fix artifacts 🤖"]:::agent
FIX --> RV
RV -- APPROVE --> BRANCH{"Branch gate 🧑<br/>where to implement? (git.work_mode)<br/>blocks any code until resolved"}
BRANCH -- "create feature branch<br/>(reviewed PR)" --> APPLY
BRANCH -- "directly on develop<br/>(flexible mode, no review)" --> APPLY
%% Build phase
APPLY["/opsx:apply 🤖"]:::agent
APPLY --> DRIFT{"Spec wrong or drifted?"}
DRIFT -- yes -
[truncated]
Branch policy — main is release-only and never worked on directly. git.work_mode in workflow.yaml decides the rest: flexible (default) lets you implement and commit directly on the integration branch — /ship then just validates, archives, and pushes, skipping PR and review; feature makes feature branches + PR mandatory. A mandatory branch gate runs before /opsx:apply writes any code: the working branch must be resolved (created and checked out) first; /git-commit re-checks at commit time as a safety net. Feature branches are named feature/<task id>-<change> when the change is linked to a backlog task with a real Jira key (e.g. feature/PROJ-123-speed-up-search ), feature/<change> otherwise.
Human checkpoints, summarized: the /start entry questions, the /req-capture interview, pasting an existing Jira ticket ( /task-import ), every language gate (es/en, mandatory on client-facing text), choosing where to implement (feature branch vs develop), commit message approval, providing Jira IDs and pasting Jira exports, PR code review, and merging via web UI when no platform CLI exists. Everything else runs agentically.
Traceability chain: Discovery → Task (Jira) → Change → tasks.md step → Commit → PR . Each link is recorded where it happens: task frontmatter ( change: ), commit footers ( Change: / Task: / Jira: ), PR description. Task IDs ARE Jira keys ( PROJ-123 , provided by you; PROJ-Dnn drafts until the issue exists). Note the naming: a task is a backlog/Jira work item; tasks.md inside a change holds implementation steps.
Each change lives in openspec/changes/<name>/ until archived. Archiving merges its delta specs into openspec/specs/ , the living description of how the system behaves. workflow.yaml configures branches (git-flow by default: feature/* → develop ), commit convention, and Jira export — commands read it, so the pipeline stays platform-agnostic.
.
├── AGENTS.md # Rules every agent must follow
├── opencode.json # opencode project config
├── workflow.yaml # Pipeline config: branches, commits, Jira (tool-agnostic)
├── templates/ # discovery.md, task.md, pr-description.md
├── backlog/
│ ├── discovery/ # Requirements-gathering notes (/req-capture)
│ ├── tasks/ # Tasks with Jira IDs (/task-generate, /task-enrich)
│ └── exports/ # jira/ (wiki markup) and pr/ (fallback PR descriptions)
├── .opencode/
│ ├── agents/
│ │ ├── spec-reviewer.md # Subagent: audits changes before apply/archive
│ │ └── task-reviewer.md # Subagent: audits tasks (sizing, testability)
│ ├── commands/
│ │ ├── start.md # /start — guided entry point (routes new work)
│ │ ├── req-capture.md # /req-capture — requirements interview
│ │ ├── task-*.md # /task-import|new|generate|enrich|jira
│ │ ├── review-*.md # /review-change, /review-task
│ │ ├── git-commit.md # /git-commit — semantic commits with traceability
│ │ ├── pr-open.md # /pr-open — platform-agnostic PR creation
│ │ ├── ship.md # /ship — validate + archive + merge
│ │ └── opsx-*.md # /opsx:* — OpenSpec lifecycle (generated)
│ └── skills/ # OpenSpec workflow skills (generated)
└── openspec/
├── config.yaml # Project context + per-artifact rules
├── specs/ # Source of truth (current behavior)
└── changes/ # In-flight changes; archive/ keeps history
Commands
Command
Stage
What it does
/start
Entry
Guided entry point: asks how the work starts and routes you (ticket exists / propose directly / create task first)
/next
Any
Detect where you are in the pipeline and s

[truncated]
