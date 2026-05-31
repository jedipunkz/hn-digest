---
source: "https://github.com/jsingh6/arch-decision"
hn_url: "https://news.ycombinator.com/item?id=48341331"
title: "Arch-Decision – A multi-agent architecture tool for Claude Code"
article_title: "GitHub - jsingh6/arch-decision: Open-source ADR orchestrator — explores your codebase, proposes approaches, writes Architecture Decision Records automatically · GitHub"
author: "jsingh2525"
captured_at: "2026-05-31T00:24:13Z"
capture_tool: "hn-digest"
hn_id: 48341331
score: 2
comments: 0
posted_at: "2026-05-30T22:45:31Z"
tags:
  - hacker-news
  - translated
---

# Arch-Decision – A multi-agent architecture tool for Claude Code

- HN: [48341331](https://news.ycombinator.com/item?id=48341331)
- Source: [github.com](https://github.com/jsingh6/arch-decision)
- Score: 2
- Comments: 0
- Posted: 2026-05-30T22:45:31Z

## Translation

タイトル: Arch-Decision – クロード コード用のマルチエージェント アーキテクチャ ツール
記事のタイトル: GitHub - jsingh6/arch-decision: オープンソース ADR オーケストレーター — コードベースを探索し、アプローチを提案し、アーキテクチャ決定記録を自動的に作成します · GitHub
説明: オープンソース ADR オーケストレーター — コードベースを探索し、アプローチを提案し、アーキテクチャ決定レコードを自動的に作成します - jsingh6/arch-decision

記事本文:
GitHub - jsingh6/arch-decision: オープンソース ADR オーケストレーター — コードベースを探索し、アプローチを提案し、アーキテクチャ決定レコードを自動的に作成します · GitHub
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
jsingh6
/
アーチ決定
公共
ノーティ

フィクション
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
16 コミット 16 コミット .claude-plugin .claude-plugin .github/ workflows .github/ workflows エージェント エージェント コミュニティ コミュニティ 出力例 出力例 スクリプト スクリプト スキル スキル DASHBOARD.md DASHBOARD.md README.md README.md すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Slack スレッドや忘れられたドキュメントによってアーキテクチャ上の決定を失うのはやめましょう。
Arch-Decision は、あらゆるコードベース内で実行され、コードベースを深く調査して、正式なアーキテクチャ決定記録を数時間ではなく数分で生成するオープンソース AI オーケストレーターです。
ADR の例を表示する · インパクト ダッシュボード · インストールする
どのチームも同じギャップに直面しています。意思決定が行われ、コードが記述されますが、誰もその理由を書き留めません。 6 か月後、新しいエンジニアが変更しますが、同じ議論が再び繰り返されます。
アーキテクチャ決定記録 (ADR) がこれを解決します。しかし、それらを書くのは摩擦です。上級エンジニアがコードベースを調査し、アプローチの草案を作成し、トレードオフ テーブルを作成し、それを明確に記述するには 2 ～ 4 時間かかります。したがって、それはほとんど起こりません。
アーチディシジョンはその摩擦を完全に取り除きます。
GitHub の問題、Jira チケット、またはわかりやすい説明を入力します。残りはこれで処理されます。
/arch-決定 https://github.com/owner/repo/issues/42
フェーズ 0 リポジトリ言語、フレームワーク、既存の ADR を検出します
フェーズ 1 問題または説明から問題を理解する
フェーズ 2 3 つの並列エージェントを起動します - 従来技術、影響マップ、制約
フェーズ 3 何かを設計する前に明確な質問をする
フェーズ 4 3 つのアプローチを同時に生成します (ミニマル / クリーン / プラグマティック)
フェーズ 5 トレードオフ テーブルを合成し、最適なパスを推奨します
フェーズ 6 あなたの承認を待ちます – ではありません

これはサインオフなしで書かれています
フェーズ 7 完全なコンテキストを含む正式な ADR を docs/decions/ に書き込みます
フェーズ 8 ADR を元の問題に結び付ける
どの言語でも動作します。あらゆるフレームワーク。どのチームでも。
GitHub コパイロット / カーソル
アーチ決定
で動作します
ファイルと行
完全な意思決定ライフサイクル
なぜ構築しているのかを知っている
いいえ
はい — 問題から始まります
ドキュメントを作成します
いいえ
はい — 正式な ADR、git にコミット
人間の承認ゲート
いいえ
はい - サインオフなしでは続行できません
セッションをまたいで再開可能
いいえ
はい — 各フェーズの後にチェックポイントを作成します
独立して検証済み
—
はい — コミュニティ PR が推奨事項を確認します
実際の例
問題 #7338 ( syncWithLocation を使用した useTable の逆フィルター マッピングの機能リクエスト) で、Refinedev/refine に対して実行されました。
このツールはコードベース コールドを調査し、 crudFiltersToColumnFilters で先行技術を発見し、packages/core でフレームワークに依存しない制約を特定し、antd ラッパーをスコープとする onParse コールバックを推奨しました。
コミュニティの貢献者は、この出力を見ることなく、同じコールバック名、同じスコープ、同じ配置で PR #7385 を独自に開きました。
このツールは、人間のエンジニアがコールド スタートから 5 分以内に判断したのと同じ結論に達しました。
Arch-Decision は、マルチエージェント アーキテクチャに基づいて構築された Claude Code プラグインです。
オーケストレーター ( skill/arch-decision/SKILL.md ) — 8 つのフェーズすべてを調整し、コンテキスト バスを維持し、人間の承認ゲートを強制します
Explorer エージェント (agents/arch-explorer/) — フェーズ 2 中に 3 つが並行して実行され、それぞれが異なる角度 (従来技術、影響マップ、および制約) に焦点を当てます。
シンセサイザー エージェント (agents/arch-synthesizer/) — 3 つのアプローチすべてからトレードオフ テーブルと推奨事項を生成します。
フェーズ 6 における人間による承認ゲートは、エージェントが明示的な署名なしに ADR を作成できないことを意味します。

オフ。これは人間の判断を加速するように設計されており、人間の判断に代わるものではありません。
クロードコードが必要です。まだインストールしていない場合は、まずインストールしてください。
クロードプラグインのインストール Arch-Decision
任意のリポジトリ内で Claude コードを開き、次を実行します。
/arch-decision < GitHub 課題 URL、Jira 説明、またはプレーン テキスト >
それだけです。渡すことができるものの例:
# GitHubの問題
/arch-決定 https://github.com/owner/repo/issues/42
# Jira チケットの概要または説明を直接貼り付けます
/arch-decion useTable で逆フィルター マッピングが必要なので、URL パラメーターがページ読み込み時にフィルター状態に同期して戻されます。
# または問題を簡単な英語で説明するだけです
/arch-decion パブリック API にレート制限を追加します
このツールは明確な質問をし、トレードオフ テーブルを使用して 3 つのアプローチを示し、何かを書く前に承認を待ちます。 ADR はリポジトリの docs/decions/ に保存されます。
問題がまだ曖昧で、最初に問題を整理したい場合は、 /arch-decision を実行する前に、ここで /arch-intake を使用します。
メトリクスは、ADR マージごとに自動的に追跡されます。
インパクトダッシュボード全体を表示する
アーチの決断/
§── スキル/
│ §── Arch-Decision/SKILL.md — 8 フェーズ オーケストレーター
│ └── Arch-intake/SKILL.md — プレスキルのフレーミングの問題
§── エージェント/
│ §── Arch-explorer/agent.md — 並列コードベース探索
│ └── Arch-synthesizer/agent.md — トレードオフ合成
§── example-output/ — 実際の ADR の例
§── scripts/generate_dashboard.py — DASHBOARD.md を自動生成します
└── .github/workflows/ — ADR マージごとにダッシュボードを更新します
ライセンス
オープンソースの ADR オーケストレーター — コードベースを探索し、アプローチを提案し、アーキテクチャ決定記録を自動的に作成します
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
そこに

読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Open-source ADR orchestrator — explores your codebase, proposes approaches, writes Architecture Decision Records automatically - jsingh6/arch-decision

GitHub - jsingh6/arch-decision: Open-source ADR orchestrator — explores your codebase, proposes approaches, writes Architecture Decision Records automatically · GitHub
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
jsingh6
/
arch-decision
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
16 Commits 16 Commits .claude-plugin .claude-plugin .github/ workflows .github/ workflows agents agents community community example-output example-output scripts scripts skills skills DASHBOARD.md DASHBOARD.md README.md README.md View all files Repository files navigation
Stop losing architectural decisions to Slack threads and forgotten docs.
arch-decision is an open-source AI orchestrator that runs inside any codebase, explores it deeply, and produces a formal Architecture Decision Record — in minutes, not hours.
View Example ADR · Impact Dashboard · Install
Every team faces the same gap: a decision gets made, the code gets written, but nobody wrote down why . Six months later a new engineer changes it — and the same debate happens all over again.
Architecture Decision Records (ADRs) solve this. But writing them is friction. It takes a senior engineer 2–4 hours to research the codebase, draft approaches, build a trade-off table, and write it up clearly. So it almost never happens.
arch-decision removes that friction entirely.
Give it a GitHub issue, a Jira ticket, or a plain description. It handles the rest:
/arch-decision https://github.com/owner/repo/issues/42
Phase 0 Detects repo language, framework, existing ADRs
Phase 1 Understands the problem from the issue or description
Phase 2 Launches 3 parallel agents — prior art, impact map, constraints
Phase 3 Asks clarifying questions before designing anything
Phase 4 Generates 3 approaches simultaneously (Minimal / Clean / Pragmatic)
Phase 5 Synthesizes a trade-off table and recommends the best path
Phase 6 Waits for your approval — nothing is written without sign-off
Phase 7 Writes a formal ADR to docs/decisions/ with full context
Phase 8 Links the ADR back to the original issue
Works on any language. Any framework. Any team.
GitHub Copilot / Cursor
arch-decision
Operates on
Files and lines
The full decision lifecycle
Knows why you're building
No
Yes — starts from the problem
Produces documentation
No
Yes — formal ADR, committed to git
Human approval gate
No
Yes — cannot proceed without sign-off
Resumable across sessions
No
Yes — checkpoints after every phase
Independently validated
—
Yes — community PRs confirm recommendations
Real Example
Ran against refinedev/refine on issue #7338 — a feature request for reverse filter mapping in useTable with syncWithLocation .
The tool explored the codebase cold, found prior art in crudFiltersToColumnFilters , identified the framework-agnostic constraint in packages/core , and recommended an onParse callback scoped to the antd wrapper.
A community contributor independently opened PR #7385 with the same callback name, same scope, same placement — without seeing this output.
The tool reached the same conclusion a human engineer did, from a cold start, in under 5 minutes.
arch-decision is a Claude Code plugin built on a multi-agent architecture:
Orchestrator ( skills/arch-decision/SKILL.md ) — coordinates all 8 phases, maintains a context bus, enforces the human approval gate
Explorer agents ( agents/arch-explorer/ ) — 3 run in parallel during Phase 2, each focused on a different angle: prior art, impact map, and constraints
Synthesizer agent ( agents/arch-synthesizer/ ) — produces the trade-off table and recommendation from all three approaches
The human approval gate in Phase 6 means the agent cannot write the ADR without your explicit sign-off. It is designed to accelerate human judgment, not replace it.
Requires Claude Code . Install it first if you haven't already.
claude plugin install arch-decision
Open Claude Code inside any repo and run:
/arch-decision < GitHub issue URL, Jira description, or plain text >
That's it. Examples of what you can pass:
# GitHub issue
/arch-decision https://github.com/owner/repo/issues/42
# Paste your Jira ticket summary or description directly
/arch-decision We need reverse filter mapping in useTable so URL params sync back to filter state on page load
# Or just describe the problem in plain English
/arch-decision add rate limiting to the public API
The tool asks you clarifying questions, shows you 3 approaches with a trade-off table, and waits for your approval before writing anything. The ADR is saved to docs/decisions/ in your repo.
If the problem is still vague and you want to frame it first, use /arch-intake your problem here before running /arch-decision .
Metrics are tracked automatically on every ADR merge.
View the full Impact Dashboard
arch-decision/
├── skills/
│ ├── arch-decision/SKILL.md — 8-phase orchestrator
│ └── arch-intake/SKILL.md — problem framing pre-skill
├── agents/
│ ├── arch-explorer/agent.md — parallel codebase exploration
│ └── arch-synthesizer/agent.md — trade-off synthesis
├── example-output/ — real ADR examples
├── scripts/generate_dashboard.py — auto-generates DASHBOARD.md
└── .github/workflows/ — updates dashboard on every ADR merge
License
Open-source ADR orchestrator — explores your codebase, proposes approaches, writes Architecture Decision Records automatically
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
