---
source: "https://github.com/tadelstein9/two-tier-memory"
hn_url: "https://news.ycombinator.com/item?id=48790505"
title: "Show HN: Two-tier-memory – queryable long-term memory for AI coding agents"
article_title: "GitHub - tadelstein9/two-tier-memory: A queryable long-term memory for AI coding agents — the two-tier fix for the context-window wall · GitHub"
author: "tadelstein"
captured_at: "2026-07-05T02:04:47Z"
capture_tool: "hn-digest"
hn_id: 48790505
score: 1
comments: 0
posted_at: "2026-07-05T01:22:16Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Two-tier-memory – queryable long-term memory for AI coding agents

- HN: [48790505](https://news.ycombinator.com/item?id=48790505)
- Source: [github.com](https://github.com/tadelstein9/two-tier-memory)
- Score: 1
- Comments: 0
- Posted: 2026-07-05T01:22:16Z

## Translation

タイトル: Show HN: 2 層メモリ – AI コーディング エージェントのクエリ可能な長期メモリ
記事のタイトル: GitHub - tadelstein9/two-tier-memory: AI コーディング エージェントのためのクエリ可能な長期メモリ — コンテキスト ウィンドウ ウォールの 2 層修正 · GitHub
説明: AI コーディング エージェント用のクエリ可能な長期メモリ — コンテキスト ウィンドウの壁の 2 層修正 - tadelstein9/two-tier-memory

記事本文:
GitHub - tadelstein9/two-tier-memory: AI コーディング エージェント用のクエリ可能な長期メモリ — コンテキスト ウィンドウ ウォールの 2 層修正 · GitHub
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
タデルシュタイン9
/
2層メモリ
公共
通知
通知を変更するにはサインインする必要があります

n設定
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット .gitignore .gitignore ライセンス ライセンス README.md README.md memory.pymemory.py schema.sql schema.sql すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング エージェントのクエリ可能な長期メモリ - コンテキスト ウィンドウの壁に対する 2 層の修正。
付属のエッセイ: 「実際の作業では AI の記憶が壊れます。修正は 50 年前です。」
AI コーディング エージェントのデフォルトの長期メモリは、各セッションの開始時にコンテキストにロードされるマークダウン ファイルのスタックです。これは 10 個のファイルで機能します。それは百四十で失敗します。すべてが 1 つの有限なコンテキスト ウィンドウをめぐって競合し、境界を越えると記憶は静かに切り捨てられます。エージェントは静かに忘れ、解決済みの問題を再解決し、先週の決定と矛盾します。あなたはモデルを責めます。問題は申告システムだ。
それは古い問題です。ファイル全体をロードしてスキャンするのは、1970 年にデータベースが発明される以前のリレーショナルの習慣であり、構造化して保存し、インデックスを付け、必要な行だけをフェッチするという目的を達成するために行われました。
層 1 — インデックス。常にロードされます。解決済みの問題ごとに 1 行: タイトルとポインタ。毎回のセッションを安価に持ち運べます。エージェントに何が存在するのかを伝えますが、詳細は決して伝えません。 ( INDEX.md 、階層 2 から生成されるため、ドリフトすることはありません。)
層 2 — オンデマンドでクエリされるデータベース。すべての困難な問題は、問題、根本原因、何がうまくいったか、落とし穴、成果物など、プレーンな SQLite ファイル内の行になります。エージェントがクエリを実行するまで、1,000 行のコストはかかりません。何か懐かしい匂いがすると、散文を読みあさるのではなく、テーブルに質問します。
ライブラリのロードを停止します。保持できるインデックスと質問できるデータベースを保持します。
python3 メモリ.py 初期化
python3 メモリ.py 追加 \
--area db --title " SQLite はできます"

ATTACH されたデータベース上のビューを永続化しない " \
--problem " 接続されたデータベースに対する CREATE VIEW が次の接続で消えてしまう " \
--root-cause " ビューはファイルではなくアタッチ エイリアスにバインドされます " \
--solution " メイン データベースに実体化するか、再アタッチして開いたときにビューを再作成します。 " \
--gotcha " エラーは発生しません。単にビューが次のセッションに存在しないだけです。" \
--tags " sqlite、アタッチ、ビュー "
Python3のmemory.pyクエリ「アタッチビューが消える」
python3 メモリ.py 1 を取得
python3memory.pyindex #データベースからINDEX.mdを再生成
データベースではない部分
ツールだけでも罠です。習慣と少しの衛生管理だけでその価値を維持できます。
再構築する前にクエリを実行します。エージェントの反射神経は解決することであり、調べることではありません。 「最初にメモリを検索する」ことを常にルールにしないと、データベースが無駄になってしまいます。
行を正直に保ちます。問題を解決したその日に書きましょう。間違っていたと判明したその日のうちに削除してください。古い行は、空のテーブルよりも誤解を招きます。
取得には上限があります。全文検索 (FTS5 経由でここに組み込まれています) は、意図した行ではなく、一致する単語を検索します。それを超えたら、同じリレーショナル ベースにセマンティック検索/埋め込みをボルトで固定します。そこから始めないでください。
エージェントのプロジェクト指示には次の 2 つのルールが適用されます。
新しい難しい問題については、まずクエリしてください。ヒットに基づいて行動します。
何か新しいことを解決するには、行を追加してからインデックスを追加します。
常にロードされる INDEX.md は、エージェントがすでに知っている情報を準備します。残りはデータベースに無料で保存されます。
ファイル
役割
スキーマ.sql
ソリューション テーブル + FTS5 フルテキスト インデックスはトリガーによって同期が維持されます
メモリ.py
stdlib のみの CLI: init / add / query / get / list /index
インデックス.md
tier-1 インデックス、データベースから生成 (git 無視)
メモリ.db
Tier-2 ストア — データ (git 無視)
依存関係はありません。 Python 3 とそれにバンドルされている sqlite3 。
によって野外に建てられた

トム・アデルスタイン。このパターンは、Anthropic のコーディング エージェントである Claude Code との作業セッションから生まれました。私はリレーショナル アプローチを提案しました。エージェントはこの実装を構築、テスト、強化しました。人間と AI のコラボレーションがポイントであり、適切なことに、付随するエッセイの主題でもあります。
AI コーディング エージェントのクエリ可能な長期メモリ — コンテキスト ウィンドウの壁に対する 2 層の修正
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

A queryable long-term memory for AI coding agents — the two-tier fix for the context-window wall - tadelstein9/two-tier-memory

GitHub - tadelstein9/two-tier-memory: A queryable long-term memory for AI coding agents — the two-tier fix for the context-window wall · GitHub
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
tadelstein9
/
two-tier-memory
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits .gitignore .gitignore LICENSE LICENSE README.md README.md memory.py memory.py schema.sql schema.sql View all files Repository files navigation
A queryable long-term memory for AI coding agents — the two-tier fix for the context-window wall.
Companion essay: "Your AI's Memory Breaks on Real Work. The Fix Is Fifty Years Old."
An AI coding agent's default long-term memory is a stack of markdown files it loads into context at the start of every session. That works at ten files. It fails at a hundred and forty: everything competes for one finite context window, and past the edge the memory silently truncates — the agent quietly forgets, re-solves solved problems, contradicts last week's decisions. You blame the model. The filing system is the problem.
It's an old problem. Loading a whole file and scanning it is the pre-relational habit that databases were invented, in 1970, to kill: store it structured, index it, fetch only the row you need.
Tier 1 — an index, always loaded. One line per solved problem: a title and a pointer. Cheap to carry every session. It tells the agent what exists , never the detail. ( INDEX.md , generated from tier 2 so it never drifts.)
Tier 2 — a database, queried on demand. Every hard problem becomes a row in a plain SQLite file — problem, root cause, what worked, the gotcha, the artifacts. A thousand rows cost the agent nothing until it runs a query. When something smells familiar, it asks the table instead of rummaging through prose.
Stop loading the library. Keep an index you can hold, and a database you can question.
python3 memory.py init
python3 memory.py add \
--area db --title " SQLite can't persist a view over an ATTACH-ed database " \
--problem " CREATE VIEW over an attached db vanished on the next connection " \
--root-cause " the view binds to the attach alias, not the file " \
--solution " materialize into the main db, or re-ATTACH and re-create the view on open " \
--gotcha " no error is raised — the view simply isn't there next session " \
--tags " sqlite,attach,view "
python3 memory.py query " attach view disappears "
python3 memory.py get 1
python3 memory.py index # regenerate INDEX.md from the database
The part that isn't the database
The tool alone is a trap. It earns its keep only with a habit and a little hygiene:
Query before you rebuild. An agent's reflex is to solve, not to look up. Make "search the memory first" a standing rule, or the database becomes dead weight.
Keep rows honest. Write them the day you solve the thing; delete them the day they turn out wrong. A stale row misleads worse than an empty table.
Retrieval has a ceiling. Full-text search (built in here, via FTS5) finds the words you match, not the row you meant . When you outgrow it, bolt semantic search / embeddings onto this same relational base — don't start there.
Point your agent's project instructions at two rules:
On a new hard problem, query first ; act on a hit.
On solving something novel, add a row , then index .
The always-loaded INDEX.md primes the agent on what it already knows; the database holds the rest, for free.
File
Role
schema.sql
the solutions table + an FTS5 full-text index kept in sync by triggers
memory.py
stdlib-only CLI: init / add / query / get / list / index
INDEX.md
tier-1 index, generated from the database (git-ignored)
memory.db
tier-2 store — your data (git-ignored)
No dependencies. Python 3 and its bundled sqlite3 .
Built in the open by Tom Adelstein. The pattern came out of a working session with Claude Code, Anthropic's coding agent: I pitched the relational approach; the agent built, tested, and hardened this implementation. The human–AI collaboration is the point — and, fittingly, the subject of the companion essay.
A queryable long-term memory for AI coding agents — the two-tier fix for the context-window wall
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
