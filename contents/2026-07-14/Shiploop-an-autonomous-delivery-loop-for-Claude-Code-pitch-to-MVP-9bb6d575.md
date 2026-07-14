---
source: "https://github.com/BechsteinDigital/claude-shiploop"
hn_url: "https://news.ycombinator.com/item?id=48906000"
title: "Shiploop – an autonomous delivery loop for Claude Code, pitch to MVP"
article_title: "GitHub - BechsteinDigital/claude-shiploop: Autonomous delivery skills for Claude Code — pitch an idea, get an MVP. CEO/PO/DEV/Reviewer agent team with contracts, gates, and evidence rules. · GitHub"
author: "dbechstein"
captured_at: "2026-07-14T13:08:29Z"
capture_tool: "hn-digest"
hn_id: 48906000
score: 2
comments: 0
posted_at: "2026-07-14T12:47:24Z"
tags:
  - hacker-news
  - translated
---

# Shiploop – an autonomous delivery loop for Claude Code, pitch to MVP

- HN: [48906000](https://news.ycombinator.com/item?id=48906000)
- Source: [github.com](https://github.com/BechsteinDigital/claude-shiploop)
- Score: 2
- Comments: 0
- Posted: 2026-07-14T12:47:24Z

## Translation

タイトル: Shiploop – Claude Code の自律配信ループ、MVP への提案
記事のタイトル: GitHub - BechsteinDigital/claude-shiploop: Claude Code の自律配信スキル — アイデアを提案し、MVP を獲得します。契約、ゲート、および証拠ルールを備えた CEO/PO/DEV/レビュー担当者エージェント チーム。 · GitHub
説明: Claude Code の自律配信スキル — アイデアを提案し、MVP を獲得します。契約、ゲート、および証拠ルールを備えた CEO/PO/DEV/レビュー担当者エージェント チーム。 - BechsteinDigital/claude-shiploop

記事本文:
GitHub - BechsteinDigital/claude-shiploop: Claude Code の自律配信スキル — アイデアを提案し、MVP を獲得します。契約、ゲート、および証拠ルールを備えた CEO/PO/DEV/レビュー担当者エージェント チーム。 · GitHub
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
ベヒシュタインデジタル
/
クロード・シップループ
出版

c
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
BechsteinDigital/クロードシップループ
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
8 コミット 8 コミット .claude-plugin .claude-plugin .github .github _shared _shared 自律ループ 自律ループ 自律セットアップ 自律セットアップ ドキュメント ドキュメントの例 例 プロジェクト-オンボーディング プロジェクト-オンボーディング 役割-監査人 役割-監査人 役割-ceo 役割-ceo 役割-dev 役割-dev 役割-PO 役割-PO 役割-レビューア 役割-レビューア ライセンスライセンス README.de.md README.de.md README.md README.md install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Claude Code の自律的な配信スキル — アイデアを提案し、MVP を獲得します。
プロジェクトや言語にとらわれないスキルスイート: アイデアの提案から
コントラクト、ゲート、およびパラレルな自己制御配信ループへの自律セットアップ
希望ではなく証拠がルールとなる。
実証済みのエンドツーエンド、アーティファクトが含まれています: 上のデモは実際の実行を再現しています — ピッチ→
インタビュー → セットアップ → 3 ループ サイクル (2 つの DEV エージェントを並列) → MVP ゲート (スタンドアップ CLI、
36/36 テスト)。未編集のプロジェクト成果物は example/standup にあります。
そして CI はプッシュするたびに完全なテスト スイートを再実行します。
ピッチ ──▶ プロジェクトオンボーディング ──▶ 自律セットアップ ──▶ 自律ループ ──▶ MVP レポート
(リサーチ、スタック、(CEO→PO→DEV∥DEV→REVIEWER、
アクティブなステップ) スキャフォールド、バックログ) アイデアファネル、ゲート)
アイデアを提案し、新人研修のインタビューに 1 回答えます。その後のすべてが実行されます
自律的に: スイートはリサーチし、スタックを選択し、プロジェクトの足場を築き、作業パッケージを切り出し、
そして、MVP ゲートまでロール エージェントを並行して調整します - あなただけにエスカレーションします

に
あなたが明示的に同意した基準。
プラグインとして (推奨): Claude Code 内で実行します。
/プラグイン マーケットプレイス add BechsteinDigital/claude-shiploop
/プラグインのインストール claude-shiploop@bechstein-digital
またはインストール スクリプト経由で:
git clone https://github.com/BechsteinDigital/claude-shiploop.git
cd クロード・シップループ
./install.sh /path/to/your/project # または: ./install.sh --global
次に、ターゲット プロジェクト内でクロード コード セッションを開始し、アイデアを提案します。
project-onboarding はそこからそれを拾います。
スキル
目的
プロジェクトオンボーディング
「Ready」の定義についてのインタビューを提案します。承認された概要を作成します。コア契約と自律契約
自律セットアップ
質問なしのブートストラッピング: 研究、ADR、足場、バックログ
自律ループ
MVP ゲートまで、(並列) サブエージェントとしてロールを調整します。
CEOの役割
ポートフォリオの決定、ゲート、WIP、アンチスラッシュ
ロールポ
作業パッケージのカット、合格基準、クレームゾーン、アイデアのトリアージ（価値フィルター）
ロール開発
正確に 1 つのパッケージを実装し、サイレントな証拠出力を実行します。
役割のレビュー担当者
デルタレビュー: 受理、請求監査、ゾーンチェック
監査役の役割
並列読み取り専用ファンアウトによる状態監査
ゆるいスキル集ではなく、なぜこれなのでしょうか？
ほとんどのスキル コレクションはツールボックスですが、それでも運転します。このスイートは次のオペレーティング モデルです。
契約による自律性: 明示的に合意された基準に従ってのみユーザーにエスカレーションします。
それ以外はすべて決定され、記録されます。
訴えではなく仕組みとして焦点を当てる: 価値フィルター、クーリングオフ、アイデアチェーンルール、拡張
予算 — 禁止だけでは、自律的なループを正常に保つことはできません。
主張には証拠が必要です。コードとテストが証明する以上のことをドキュメントで主張することはできません。
クレーム ゾーンによる並列処理: パッケージごとの互いに素なファイル ゾーン。疑わしい場合はワークツリーを分離します。
プロジェクト成果物 (ターゲット プロジェクトで作成されたもの)
プロジェクト/
BRIEF.md 製品

概要 + 核心契約 + 自治契約 (憲法)
PROFILE.md スタック、検証済みコマンド、品質ルール (コマンドの単一ソース)
STATE.md 単一状態ストア (WIP、マイルストーン、予算、サイクル)
DECISIONS.md 自律的な決定の ADR-light ログ
IDEAS.md トリアージ ルールを含むアイデア ファネル
このプロジェクトの LEARNINGS.md レトロ蒸留物 (ゲート要件)
バックログ/WORK カード (クレーム ゾーン + 複雑さを含む)
ログ/調査、監査、サイクルログ
設計原則
真実ごとに 1 つのソース: コマンドのみ PROFILE.md 、状態のみ STATE.md — いいえ
ドリフトするスナップショットの重複。
魅力ではなくメカニズムとして焦点を当てる: 価値フィルター、クーリングオフ、アイデアチェーンルール、拡張
予算 — 禁止だけでは、自律的なループを正常に保つことはできません。
契約による自律性: 明示的に合意された基準に従ってのみユーザーにエスカレーションします。
それ以外はすべて決定され、記録されます。
クレーム ゾーンによる並列処理: パッケージごとの互いに素なファイル ゾーン。ワークツリーの分離時
不確かな。
主張には証拠が必要です。コードとテストが証明する以上のことをドキュメントで主張することはできません。
モデル階層: 判断が行われる高価なトークン (オーケストレーター、PO カッティング、ゲート)
レビュー）;実行が行われる安価なトークン (DEV/カードの複雑さ、監査人によるレビュー)
ファンアウト）。小型モデルを安全にするのはカードの品質です。だからこそ、PO は安全ではありません。
格下げされた。
留出者としての経験：マイルストーンゲートでレトロ→最大。 3 ～ 5 回の学習
project/LEARNINGS.md 、グローバル KB _shared/knowledge/ (マスター) に一般化可能なもの
場所のみであり、プロジェクトにはインストールされません)。セットアップとレビューはすべてのプロジェクトで読み取られます
これまでのすべての経験から始まります。
./install.sh /path/to/project # スキル + _shared/ を <project>/.claude/skills/ にコピーします
./install.sh --global # すべてのプロジェクトの ~/.claude/skills/ にインストールします
スキルは

インストール場所を基準にして _shared/scripts/ 内のスクリプトを参照します —
どちらの亜種も機能します。
グローバル ナレッジ ベースはこのリポジトリに残ります (単一ソースであり、プロジェクトにはインストールされません)。
install.sh はインストール内の絶対パス ( _shared/knowledge.path ) を記録するため、スキル
任意のプロジェクトから見つけられます。固定されたクローンの場所は必要ありません。リポジトリを移動した場合は、再実行してください
./install.sh 、または import SKILLS_KNOWLEDGE_DIR=<repo>/_shared/knowledge 経由でオーバーライドします。
プラグインとしてインストールした場合、knowledge.path はありません (プラグインのキャッシュは一時的です)。
このリポジトリを複製して SKILLS_KNOWLEDGE_DIR を設定しない限り、学習はプロジェクト ローカルのままになります。
操作上の注意: クロード コード セッションは常にターゲット プロジェクト (サブエージェント) で開始してください。
セッションルートの外に書き込むことはできません。自律ループの場合、DEV サブエージェントには権限が必要です
対話型プロンプトなしのモード ( acceptEdits / bypassPermissions )、それ以外の場合はすべての書き込み
誰も見ていない間にブロックします。
Claude Code の自律的な配信スキル — アイデアを提案し、MVP を獲得します。契約、ゲート、および証拠ルールを備えた CEO/PO/DEV/レビュー担当者エージェント チーム。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
v1.0.0: Ko-fi スポンサーシップを追加 (FUNDING.yml + README バッジ)
最新の
2026 年 7 月 14 日
このプロジェクトのスポンサーになる
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Autonomous delivery skills for Claude Code — pitch an idea, get an MVP. CEO/PO/DEV/Reviewer agent team with contracts, gates, and evidence rules. - BechsteinDigital/claude-shiploop

GitHub - BechsteinDigital/claude-shiploop: Autonomous delivery skills for Claude Code — pitch an idea, get an MVP. CEO/PO/DEV/Reviewer agent team with contracts, gates, and evidence rules. · GitHub
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
BechsteinDigital
/
claude-shiploop
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
BechsteinDigital/claude-shiploop
main Branches Tags Go to file Code Open more actions menu Folders and files
8 Commits 8 Commits .claude-plugin .claude-plugin .github .github _shared _shared autonomous-loop autonomous-loop autonomous-setup autonomous-setup docs docs examples examples project-onboarding project-onboarding role-auditor role-auditor role-ceo role-ceo role-dev role-dev role-po role-po role-reviewer role-reviewer LICENSE LICENSE README.de.md README.de.md README.md README.md install.sh install.sh View all files Repository files navigation
Autonomous delivery skills for Claude Code — pitch an idea, get an MVP.
A project- and language-agnostic skill suite: from an idea pitch through
autonomous setup to a parallel, self-controlled delivery loop — with contracts, gates, and
evidence rules instead of hope.
Proven end-to-end, artifacts included: the demo above replays a real run — pitch →
interview → setup → 3 loop cycles (2 DEV agents in parallel) → MVP gate ( standup CLI,
36/36 tests). The unedited project artifacts live in examples/standup ,
and CI re-runs its full test suite on every push.
Pitch ──▶ project-onboarding ──▶ autonomous-setup ──▶ autonomous-loop ──▶ MVP report
(the only inter- (research, stack, (CEO→PO→DEV∥DEV→REVIEWER,
active step) scaffold, backlog) idea funnel, gates)
You pitch an idea and answer the onboarding interview once. Everything after that runs
autonomously: the suite researches, picks a stack, scaffolds the project, cuts work packages,
and orchestrates role agents in parallel until the MVP gate — escalating to you only on
criteria you explicitly agreed to.
As a plugin (recommended): inside Claude Code, run
/plugin marketplace add BechsteinDigital/claude-shiploop
/plugin install claude-shiploop@bechstein-digital
Or via install script:
git clone https://github.com/BechsteinDigital/claude-shiploop.git
cd claude-shiploop
./install.sh /path/to/your/project # or: ./install.sh --global
Then start a Claude Code session inside the target project and pitch your idea —
project-onboarding picks it up from there.
Skill
Purpose
project-onboarding
Pitch interview to Definition of Ready; produces an approved brief incl. core contract and autonomy contract
autonomous-setup
Bootstrapping without questions: research, ADRs, scaffold, backlog
autonomous-loop
Orchestrates roles as (parallel) subagents until the MVP gate
role-ceo
Portfolio decisions, gates, WIP, anti-thrash
role-po
Work-package cutting, acceptance criteria, claim zones, idea triage (value filter)
role-dev
Implements exactly one package, silent, evidence output
role-reviewer
Delta review: acceptance, claim audit, zone check
role-auditor
State audit via parallel read-only fan-out
Why this and not a loose skill collection?
Most skill collections are toolboxes — you still drive. This suite is an operating model:
Autonomy with a contract: escalation to the user only via explicitly agreed criteria;
everything else gets decided and logged.
Focus as mechanics, not appeals: value filter, cooling-off, idea-chain rule, extension
budget — prohibitions alone don't keep an autonomous loop on course.
Claims need evidence: documentation may never claim more than code + tests prove.
Parallelism via claim zones: disjoint file zones per package; worktree isolation when in doubt.
Project artifacts (created in the target project)
project/
BRIEF.md product brief + core contract + autonomy contract (the constitution)
PROFILE.md stack, verified commands, quality rules (single source for commands)
STATE.md single state store (WIP, milestone, budget, cycle)
DECISIONS.md ADR-light log of autonomous decisions
IDEAS.md idea funnel with triage rules
LEARNINGS.md retro distillate of this project (gate requirement)
backlog/ WORK cards (incl. claim zone + complexity)
log/ research, audit, and cycle logs
Design principles
One source per truth: commands only in PROFILE.md , state only in STATE.md — no
snapshot duplicates that drift.
Focus as mechanics, not as appeal: value filter, cooling-off, idea-chain rule, extension
budget — prohibitions alone don't keep an autonomous loop on course.
Autonomy with a contract: escalation to the user only via explicitly agreed criteria;
everything else is decided and logged.
Parallelism via claim zones: disjoint file zones per package; worktree isolation when
uncertain.
Claims need evidence: documentation may never claim more than code + tests prove.
Model hierarchy: expensive tokens where judgment happens (orchestrator, PO cutting, gate
reviews); cheap tokens where execution happens (DEV/review by card complexity, auditor
fan-out). Card quality is what makes small models safe — which is why the PO is never
downgraded.
Experience as distillate: retro at the milestone gate → max. 3–5 learnings into
project/LEARNINGS.md , generalizable ones into the global KB _shared/knowledge/ (master
location only, not installed into projects). Setup and reviews read them — every project
starts with the experience of all previous ones.
./install.sh /path/to/project # copies skills + _shared/ to <project>/.claude/skills/
./install.sh --global # installs to ~/.claude/skills/ for all projects
The skills reference the scripts in _shared/scripts/ relative to the installation location —
both variants work.
The global knowledge base stays in this repo (single source, never installed into projects).
install.sh records its absolute path in the installation ( _shared/knowledge.path ), so skills
find it from any project — no fixed clone location required. If you move the repo, re-run
./install.sh , or override via export SKILLS_KNOWLEDGE_DIR=<repo>/_shared/knowledge .
When installed as a plugin, there is no knowledge.path (the plugin cache is ephemeral):
learnings then stay project-local, unless you clone this repo and set SKILLS_KNOWLEDGE_DIR .
Operational note: always start the Claude Code session in the target project — subagents
cannot write outside the session root. For the autonomous loop, DEV subagents need a permission
mode without interactive prompts ( acceptEdits / bypassPermissions ), otherwise every write
blocks while nobody is watching.
Autonomous delivery skills for Claude Code — pitch an idea, get an MVP. CEO/PO/DEV/Reviewer agent team with contracts, gates, and evidence rules.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
v1.0.0: Add Ko-fi sponsorship (FUNDING.yml + README badges)
Latest
Jul 14, 2026
Sponsor this project
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
