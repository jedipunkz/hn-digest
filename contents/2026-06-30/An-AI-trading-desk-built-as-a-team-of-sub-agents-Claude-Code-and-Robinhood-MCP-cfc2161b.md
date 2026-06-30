---
source: "https://github.com/LoganYangBo/rh-trading-agent"
hn_url: "https://news.ycombinator.com/item?id=48733849"
title: "An AI trading desk built as a team of sub-agents (Claude Code and Robinhood MCP)"
article_title: "GitHub - LoganYangBo/rh-trading-agent · GitHub"
author: "loganboyang"
captured_at: "2026-06-30T15:49:53Z"
capture_tool: "hn-digest"
hn_id: 48733849
score: 1
comments: 0
posted_at: "2026-06-30T15:11:46Z"
tags:
  - hacker-news
  - translated
---

# An AI trading desk built as a team of sub-agents (Claude Code and Robinhood MCP)

- HN: [48733849](https://news.ycombinator.com/item?id=48733849)
- Source: [github.com](https://github.com/LoganYangBo/rh-trading-agent)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T15:11:46Z

## Translation

タイトル: サブエージェントのチームとして構築された AI トレーディング デスク (Claude Code と Robinhood MCP)
記事タイトル: GitHub - LoganYangBo/rh-trading-agent · GitHub
説明: GitHub でアカウントを作成して、LoganYangBo/rh-trading-agent の開発に貢献します。

記事本文:
GitHub - LoganYangBo/rh-trading-agent · GitHub
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
ローガンヤンボー
/
rh-トレーディングエージェント
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイル コードに移動 さらにアクションを開く

メニュー フォルダーとファイル
6 コミット 6 コミット .claude .claude docs docs Strategies Strategies ui ui .gitignore .gitignore .mcp.json .mcp.json CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md すべてのファイルを表示 リポジトリ ファイル ナビゲーション
🟢 rh-trading-agent — クロード コードの AI トレーディング デスク
クロード コード内で実行されるマルチエージェントの株式調査デスクは、
MCP を介した Robinhood Agentic アカウントを使用し、お客様の承認なしに注文を行うことはありません。
これは「お金を盗むボット」ではありません。専門のサブエージェントからなる小規模なチームです
— ファンダメンタルズ、テクニカル、マクロ/ニュース、および拒否権を持つリスクマネージャー — その画面
ウォッチリストに登録し、各候補者について討論し、ワンクリックでプレビュー カードを渡します。あなた
承認する。それは置きます。すべては書かれたガードレールに包まれています（人間参加型、
ポジションキャップ、プロンプトインジェクションディフェンスなど）、ロビンフッドスタイルのダッシュボードにミラーリングされます。
1 つのプロンプトではなくチーム — アナリストは並行して証拠を収集します。独立したリスク
マネジャーはアナリストが気に入った取引に拒否権を発動できる。
ガードレールは構造的なものであり、雰囲気ではありません。サブエージェントには物理的に命令ツールがありません。
あなたと PM だけが参加でき、セッション中の明示的な承認後にのみ参加できます。
プロンプトインジェクション対応 — ニュースエージェントは取得したコンテンツを信頼できないデータとして扱います
そして、疑わしい「指示」を実行せずに引用します。
設計によりロータッチ — スケジュールに基づいて読み取り専用の調査を実行し、結果を表面化するだけです。
本当に資格があるときに取引する。ほとんどの場合、脇に立つように言われます。
本物のダッシュボード — ロビンフッド スタイルの UI は、デスクの状態をライブで反映します。
⚠️ リアルマネー、ベータ版であり、投資アドバイスではありません。 Robinhood Agentic Trading はベータ版です
(米国、株式のみ)。エージェントは、資金が提供されている隔離されたエージェントアカウント内でのみ取引します。
専用の予算を用意します。その予算は、これまでに失われる可能性のある最大の予算です。

ありません
ここでは実績とパフォーマンスに関する主張はありません。これはリファレンスアーキテクチャです
学習中は、ご自身の責任で実行し、ご自身で監視してください。
このリポジトリには、ガードレール、戦略ノート、エージェント、ダッシュボード、セットアップ ドキュメントが含まれています。
シークレットや OAuth トークンは決して使用しません (これらはリポジトリの外に存在します)。
。
§── CLAUDE.md # PMの運営規約（従うべきルール）
§── .mcp.json # プロジェクトを対象とした Robinhood Trading MCP 接続
§── .env.example # .env のテンプレート (アカウント番号; .env は gitignored)
§── .claude/
│ §── settings.json # 権限: 読み取りは許可され、オーダーはゲートされ、オプションは拒否されます
│ └── エージェント/ # デスクチーム — 役割ごとに 1 人のサブエージェント
│ §──foundation-analyst.md
│ §── Technical-analyst.md
│ §── Macro-news-analyst.md # インジェクション分離 (唯一の Web 側の役割)
│ └──risk-manager.md # あらゆる取引に対する拒否権
§── 戦略/
│ §── README.md # リスクキャップ + リスクマネージャーが拒否権を発動しなければならない場合
│ lux──mean-reversion.md # 戦略例（エントリー/イグジット/サイジング）
§── ドキュメント/
│ §── SETUP.md # OAuth セットアップ + ガードレールの締め方
│ └── TEAM.md # デスクの役割とエンドツーエンドのワークフロー
└── ui/ # 読み取り専用の Robinhood スタイルのダッシュボード (Vite + React)
§── public/desk-state.example.json # デモ データ (ライブのdesk-state.json は gitignored です)
└── src/ # スナップショットをポーリングし、アカウント/デスク/プレビューをレンダリングします
クイックスタート
何かをプッシュする前に、このリポジトリを非公開にしてください。
docs/SETUP.md を読み、上から下までに従ってください。
MCP に接続し、OAuth (デスクトップ) 経由で認証し、少額のエージェント予算を提供します。
実際のツール名がわかったら、.claude/settings.json を修正します。
ポートフォリオマネージャー (PM = メインエージェント) によって調整されるサブエージェントの小さなチーム
クロード・コード・セッシ

の上） 。分析者は並行して証拠を収集します。リスクマネージャーには拒否権がある
力; PM のみが注文を行うことができますが、それはセッション中の承認後にのみ可能です。後
実行するたびに、PM は ui/public/desk-state.json を書き込み、ダッシュボードに反映されます。
フローチャート TD
U([あなた: PM に平易な言葉で質問します]) --> PM{{ポートフォリオ マネージャー<br/>メイン セッション · 注文できる唯一の役割}}
サブグラフ SENSE [1.センス】
PM --> A1[get_portfolio / get_equity_positions<br/>エージェントアカウントのみ]
終わり
サブグラフ 研究 [2-3.画面 + 研究 · 並列 · 読み取り専用]
A1 --> F[ファンダメンタルズ・アナリスト<br/>評価・収益]
A1 --> T[テクニカル アナリスト<br/>トレンド · レベル · スキャン]
A1 --> M[マクロ / ニュース アナリスト<br/>背景 · ニュース · インジェクション分離]
終わり
F --> SYN[4. PM は提案された取引を総合します<br/>戦略のルールに関連付けられます/]
T --> SYN
M --> SYN
SYN --> リスク{5.リスク マネージャー<br/>チェックと戦略/}
リスク -- 拒否権 --> 停止([取引停止 · PM が報告])
リスク -- 承認 / 変更 --> 前へ [6. review_equity_order<br/>ビルド プレビュー カード]
PREV --> SNAP[(7.desk-state.jsonを作成する)]
スナップ --> ゲート[8.プレビューを表示 · ⏸ 承認を待ちます]
ゲート -- あなたは「はい」と答えます --> EXEC[9. place_equity_order<br/>依然としてaskルールによってゲートされています]
ゲート -- あなたはノーと言う --> 停止
実行 --> CONF[10.記入の確認・desk-state.jsonの更新・ログ]
スナップ - 。 5 秒ごとにポーリングします。-> DASH[/Dashboard: npm run dev/]
コンフ - 。アップデート .-> ダッシュ
読み込み中
ステップ 1 ～ 7 は調査であり、順序は生成されません。デスクの標準出力は
ステップ 8 でカードをプレビューします。確認するまでそこで停止します。完全な役割/ツールの内訳
docs/TEAM.md にあります。
# 1. 1 回限り: MCP に接続して認証し、アカウント番号を記録します。
claude # プロジェクトを開きます (.mcp.json サーバーを信頼します)
# セッション内: /mcp # robinhood-trading → OAuth を選択 (デスクトップ + モバイル検証)
c

p .env.example .env # 次に、.env に Agentic アカウント番号を入力します
# 2. ダッシュボードを実行します (別のターミナル) — 各デスクの実行を反映します
cd ui && npm install && npm run dev # http://localhost:5180 (ライブ実行までのデモを表示)
# 3. デスクを操作する — クロード コード セッションで PM に話しかけるだけです。例:
# 「ウォッチリストを確認して、チーム全体の分析とともに上位 2 つのアイデアを持ってきてください。」
# "AAPL と NVDA でデスクを実行し、小規模なスターターをより良いものでリスクチェックします。"
# PM はアナリスト → リスク マネージャー → カードのプレビュー → あなたの OK を待ちます。
# キルスイッチ: MCP を Robinhood アプリから切断するか、ローカルで削除します
クロード・MCP ロビンフッド・トレーディングを削除します
注: .claude/agents/ 内のサブエージェントは、Claude Code の開始時にロードされます。
追加または編集してセッションを再起動すると、ファンダメンタルズアナリストなどの役割が有効になります。
制限されたツールセットで認識されます。
安全姿勢（資金調達前にお読みください）
エージェントは Agentic アカウントでのみ取引でき、メイン残高では取引できません。
デフォルトの権限ルールでは、すべての Robinhood ツール呼び出しが手動プロンプトの背後に置かれます。
CLAUDE.md にはプロンプト挿入ルールが含まれています: エージェントは取引を無視する必要があります
取得された/外部コンテンツ (ニュース、アナリスト ノート、Web) にある指示。
Robinhood アプリからいつでも MCP を切断できます。これがキル スイッチです。
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

Contribute to LoganYangBo/rh-trading-agent development by creating an account on GitHub.

GitHub - LoganYangBo/rh-trading-agent · GitHub
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
LoganYangBo
/
rh-trading-agent
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits .claude .claude docs docs strategies strategies ui ui .gitignore .gitignore .mcp.json .mcp.json CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md View all files Repository files navigation
🟢 rh-trading-agent — an AI trading desk on Claude Code
A multi-agent stock-research desk that runs inside Claude Code, connects to a
Robinhood Agentic account over MCP, and never places an order without your approval.
It's not a "bot that YOLOs your money." It's a small team of specialized sub-agents
— fundamental, technical, macro/news, and a risk manager with veto power — that screen
your watchlist, debate each candidate, and hand you a one-click preview card . You
approve; it places. Everything's wrapped in written guardrails (human-in-the-loop,
position caps, prompt-injection defense) and mirrored to a Robinhood-style dashboard.
A team, not one prompt — analysts gather evidence in parallel; an independent risk
manager can veto a trade the analysts liked.
Guardrails are structural, not vibes — sub-agents physically have no order tools;
only you-plus-the-PM can place, and only after explicit in-session approval.
Prompt-injection-aware — the news agent treats fetched content as untrusted data
and quotes suspicious "instructions" instead of acting on them.
Low-touch by design — it runs read-only research on a schedule and only surfaces a
trade when one genuinely qualifies; most days it tells you to stand aside.
A real dashboard — a Robinhood-style UI mirrors the desk's state live.
⚠️ Real money, beta, not investment advice. Robinhood Agentic Trading is in beta
(US, equities only). The agent trades only inside an isolated Agentic account funded
with a dedicated budget — that budget is the most it can ever lose. There is no
track record and no performance claim here ; this is a reference architecture for
learning, run it at your own risk and monitor it yourself.
This repo holds the guardrails, strategy notes, agents, dashboard, and setup docs —
never secrets, and never the OAuth token (those live outside the repo).
.
├── CLAUDE.md # The PM's operating contract (rules it must follow)
├── .mcp.json # Project-scoped Robinhood Trading MCP connection
├── .env.example # Template for .env (account number; .env is gitignored)
├── .claude/
│ ├── settings.json # Permissions: reads allowed, orders gated, options denied
│ └── agents/ # The desk team — one sub-agent per role
│ ├── fundamental-analyst.md
│ ├── technical-analyst.md
│ ├── macro-news-analyst.md # injection-isolated (the only web-facing role)
│ └── risk-manager.md # veto power over every trade
├── strategies/
│ ├── README.md # Risk caps + when the Risk Manager must VETO
│ └── mean-reversion.md # Example strategy (entry/exit/sizing)
├── docs/
│ ├── SETUP.md # OAuth setup + how to tighten guardrails
│ └── TEAM.md # The desk roles and end-to-end workflow
└── ui/ # Read-only Robinhood-style dashboard (Vite + React)
├── public/desk-state.example.json # demo data (live desk-state.json is gitignored)
└── src/ # polls the snapshot, renders account/desk/preview
Quickstart
Make this repo private before pushing anything.
Read docs/SETUP.md and follow it top to bottom.
Connect the MCP, authenticate via OAuth (desktop), fund a small Agentic budget.
Refine .claude/settings.json once you know the real tool names.
A small team of sub-agents coordinated by the Portfolio Manager (PM = the main
Claude Code session) . Analysts gather evidence in parallel; the Risk Manager has veto
power; only the PM can place orders — and only after your in-session approval. After
every run the PM writes ui/public/desk-state.json , which the dashboard mirrors live.
flowchart TD
U([You: ask the PM in plain language]) --> PM{{Portfolio Manager<br/>main session · only role that can order}}
subgraph SENSE [1. Sense]
PM --> A1[get_portfolio / get_equity_positions<br/>Agentic account only]
end
subgraph RESEARCH [2-3. Screen + Research · parallel · read-only]
A1 --> F[Fundamental Analyst<br/>valuation · earnings]
A1 --> T[Technical Analyst<br/>trend · levels · scans]
A1 --> M[Macro / News Analyst<br/>backdrop · news · INJECTION-ISOLATED]
end
F --> SYN[4. PM synthesizes a proposed trade<br/>tied to a rule in strategies/]
T --> SYN
M --> SYN
SYN --> RISK{5. Risk Manager<br/>checks vs strategies/}
RISK -- VETO --> STOP([Trade stops · PM reports back])
RISK -- APPROVE / CHANGES --> PREV[6. review_equity_order<br/>build preview card]
PREV --> SNAP[(7. Write desk-state.json)]
SNAP --> GATE[8. Present preview · ⏸ wait for your approval]
GATE -- you say yes --> EXEC[9. place_equity_order<br/>still gated by ask rule]
GATE -- you say no --> STOP
EXEC --> CONF[10. Confirm fill · refresh desk-state.json · log]
SNAP -. polled every 5s .-> DASH[/Dashboard: npm run dev/]
CONF -. updates .-> DASH
Loading
Steps 1–7 are research and produce no order . The desk's standard output is the
preview card at step 8 — it stops there until you confirm. Full role/tool breakdown
is in docs/TEAM.md .
# 1. One-time: connect + authenticate the MCP, then record your account number
claude # open the project (trust the .mcp.json server)
# in-session: /mcp # pick robinhood-trading → OAuth (desktop + mobile verify)
cp .env.example .env # then put your Agentic account number in .env
# 2. Run the dashboard (separate terminal) — mirrors each desk run
cd ui && npm install && npm run dev # http://localhost:5180 (shows demo until a live run)
# 3. Drive the desk — just talk to the PM in the Claude Code session, e.g.:
# "Screen my watchlist and bring me the top 2 ideas with full team analysis."
# "Run the desk on AAPL and NVDA, risk-check a small starter in the better one."
# The PM fans out to the analysts → Risk Manager → preview card → waits for your OK.
# Kill switch: disconnect the MCP from the Robinhood app, or remove it locally
claude mcp remove robinhood-trading
Note: the sub-agents in .claude/agents/ load when Claude Code starts — after
adding or editing them, restart the session so roles like fundamental-analyst are
recognized with their restricted tool sets.
Safety posture (read before funding)
The agent can only trade in the Agentic account , never your main balance.
Default permission rule puts every Robinhood tool call behind a manual prompt.
CLAUDE.md includes a prompt-injection rule: the agent must ignore trading
instructions found in fetched/external content (news, analyst notes, web).
You can disconnect the MCP anytime from the Robinhood app — that's your kill switch.
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
