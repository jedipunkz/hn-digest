---
source: "https://github.com/RohiRIK/OpenLtm"
hn_url: "https://news.ycombinator.com/item?id=48456483"
title: "OpenLTM – Local, self-decaying memory for AI coding agents"
article_title: "GitHub - RohiRIK/OpenLtm: Long-Term Memory plugin for Claude Code — semantic search, context injection, session learning · GitHub"
author: "RohiRik"
captured_at: "2026-06-09T07:22:39Z"
capture_tool: "hn-digest"
hn_id: 48456483
score: 2
comments: 0
posted_at: "2026-06-09T04:37:55Z"
tags:
  - hacker-news
  - translated
---

# OpenLTM – Local, self-decaying memory for AI coding agents

- HN: [48456483](https://news.ycombinator.com/item?id=48456483)
- Source: [github.com](https://github.com/RohiRIK/OpenLtm)
- Score: 2
- Comments: 0
- Posted: 2026-06-09T04:37:55Z

## Translation

タイトル: OpenLTM – AI コーディング エージェント用のローカルの自己減衰メモリ
記事のタイトル: GitHub - RohiRIK/OpenLtm: クロード コード用長期記憶プラグイン — セマンティック検索、コンテキスト インジェクション、セッション学習 · GitHub
説明: クロード コード用の長期記憶プラグイン — セマンティック検索、コンテキスト インジェクション、セッション学習 - RohiRIK/OpenLtm

記事本文:
GitHub - RohiRIK/OpenLtm: クロード コード用の長期記憶プラグイン — セマンティック検索、コンテキスト インジェクション、セッション学習 · GitHub
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
ロヒリク
/
OpenLtm
公共
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
よ

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
244 コミット 244 コミット .claude-plugin .claude-plugin .github .github エージェント エージェント アセット アセット コマンド コマンド ドキュメント ドキュメント グラフアプリ グラフアプリ フック フック マイグレーション マイグレーション パッケージ パッケージ スクリプト スクリプト スキル スキル src src .env.schema .env.schema .gitignore .gitignore .npmrc .npmrc .pre-commit-config.yaml .pre-commit-config.yaml .trufflehogignore .trufflehogignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md bun.lock bun.lock bunfig.toml bunfig.toml install.sh install.sh package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
認証層については一度説明しました。なぜクロードは明日また尋ねるのですか？
AI コーディング エージェントの長期記憶 - Claude Code、OpenCode、Pi
すべてのセッション、すべての更新、すべての圧縮に存続する永続的なセマンティック メモリ。
OpenLTM は、Claude Code のプライベート メモリ層として誕生しました。現在、MIT のもとで完全にオープンソースになっています。
同じエンジン — 自動キャプチャ、セマンティック リコール、重要度加重減衰、クエリ可能なメモリ グラフ — を読み取り、フォークし、拡張することができます。ある開発者の「コードベースの再説明をやめる」プラグインとして始まったものは、現在では Claude Code、OpenCode、Pi にわたるエージェント メモリのオープン基盤となっています。
🔓 MIT ライセンス — クラウド、アカウント、テレメトリはありません。メモリは、所有するローカル SQLite DB に保存されます。
🧩 プロバイダーに依存しない — ホストごとに 1 つのコア ( @rohirik/openltm-core )、シン アダプター。
🛠 ハッキング可能 — フック、スキル、管理プロバイダー、グラフ ビジュアライザーはすべて公開されています。
以前のインストールから移行しますか?ママ

rketplace は RohiRIK/OpenLtm になり、プラグインは openltm になりました。既存のメモリ データベースは引き継がれます。
メモリは自動的に記憶されるべきです。フックがその役割を果たします。セッション終了フックはパターンを抽出し、セッション開始フックはパターンを挿入して戻します。覚えておく必要はありません。
減衰は機能であり、バグではありません。半年前に経験したことのない問題は、おそらくもう当てはまりません。何かを永続的にするには、重要度: 5 を設定します。他のものはすべて自然に期限切れになります。
キーワードよりもセマンティック。 FTS5 全文検索が最初に実行されます。何も返されない場合は、ベクトル埋め込みが開始されます。正確な単語ではなく、意味によって検索します。「非同期エラーの処理方法」では、その正確な単語を書いたことがない場合でも、適切なメモリが見つかります。
構成ゼロ、ロックインゼロ。一度インストールすればどこでも動作します。すべての設定には適切なデフォルトがあります。 DB はプラグイン ディレクトリの外に存在するため、更新されるたびに存続します。クラウドもテレメトリもアカウントもありません。
🔍思い出す
過去の決断、パターン、落とし穴 — 仕事を始める前に
🧠学ぶ
すべてのセッションが自動的に行われ、手動でメモを取る必要はありません
💉 注入する
セッション開始時のトップコンテキストにより、クロードは中断したところから再開します
⏳ 衰退
古い記憶は薄れていきますが、重要な知識は永遠に残ります
🕸グラフ
推論チェーンのために記憶間の関係をたどる
🗺 視覚化する
ブラウザベースのエクスプローラーでメモリ ネットワーク全体を確認する
⚡ ベックリコール
sqlite-vec によるセマンティック ベクター (KNN) リコール。利用できない場合は JS-cosine に劣化します
🔌 拡張機能
sqlite-vec + Honker (キュー/cron/pub-sub) は動的にロードされます。システム libsqlite3 を使用しないグレースフル フォールバック
インストール
クロード プラグイン マーケットプレイスを追加 https://github.com/RohiRIK/OpenLtm
クロードプラグインをインストールopenltm
クロードコードを再起動します。それでおしまい。 4 つのクロード コード フック + 1 つの git ポストコミット フック自動配線、4 つのコマンド ロード、5 つの sk

病気がアクティブになり、openltm.db が移行または作成されます。
bunx @rohirik/openltm-core # クロード コード、OpenCode を自動検出
bunx @rohirik/openltm-core --pi # 実験的な Pi アダプター
bunx @rohirik/openltm-core --dry-run --claude # 書き込みなしでプレビュー
開発 / git クローン
git clone https://github.com/RohiRIK/OpenLtm ~ /Projects/OpenLtm
cd ~ /Projects/OpenLtm && bash install.sh
クイックスタート
新しいセッションを開始します。コンテキストは上部に自動的に挿入されます。
/openltm:memory remember auth — このプロジェクトの認証について何がわかっていますか?
/openltm:memory learn <insight> — 残しておく価値のあるものを保存する
/openltm:health — メモリの健全性と減衰の概要
/openltm:project init — 現在のプロジェクトの目標を設定します
それだけです。残りはフックが処理を行います。
クロード・コード
│
§── 4つのコマンド ──┐
§── 5 つのスキル ──┼──▶ openltm MCP サーバー ──▶ openltm.db
└── 5 フック ──┘ (思い出、タグ、
context_items、
記憶関係、
思い出_fts)
「仕組みとアーキテクチャ」では、スキーマ、フック アーキテクチャ、減衰式、ADR について詳しく説明しています。
SQLite 拡張機能: システム拡張機能が有効な libsqlite3 が利用可能な場合、プラグインは sqlite-vec (vec0 / KNN ベクトル検索) と Honker (非同期埋め込みキュー、リーダー選出 cron、pub-sub) をロードします。どちらも正常に機能を低下させます。バイナリまたはライブラリが欠落している場合、機能はオフのままになり、JS-cosine、ファイル監視ポーリング、またはインプロセス cron にフォールバックします。 LTM_DISABLE_VEC 、 LTM_DISABLE_HONKER 、 LTM_SQLITE_LIB 、および LTM_HONKER_EXT 環境変数で制御されます。
/openltm:health # プラグインの健全性 + フック + 減衰
/openltm:memory remember test # 結果を返します (新規インストールの場合は「結果なし」)
新しいセッションを開始します。上部にコンテキストが挿入されているのが表示されます。そうでない場合は、/openltm:health を実行して診断します。
完全なドキュメントのインデックス: docs/ 。
開ける

まず問題を提起して変更について話し合ってから、PR を送信します。完全なワークフロー (セットアップ、テスト、バージョン バンプ ルール、リリース手順) は CONTRIBUTING.md にあります。
カフェイン、SQLite、そしてセッションの終了時にコンテキストが消滅すべきではないという根強い信念によって支えられています。
クロード コード用の長期記憶プラグイン — セマンティック検索、コンテキスト インジェクション、セッション学習
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
12
v2.9.1
最新の
2026 年 6 月 9 日
+ 11 リリース
このプロジェクトのスポンサーになる
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Long-Term Memory plugin for Claude Code — semantic search, context injection, session learning - RohiRIK/OpenLtm

GitHub - RohiRIK/OpenLtm: Long-Term Memory plugin for Claude Code — semantic search, context injection, session learning · GitHub
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
RohiRIK
/
OpenLtm
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
244 Commits 244 Commits .claude-plugin .claude-plugin .github .github agents agents assets assets commands commands docs docs graph-app graph-app hooks hooks migrations migrations packages packages scripts scripts skills skills src src .env.schema .env.schema .gitignore .gitignore .npmrc .npmrc .pre-commit-config.yaml .pre-commit-config.yaml .trufflehogignore .trufflehogignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md bun.lock bun.lock bunfig.toml bunfig.toml install.sh install.sh package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
You explained your auth layer once. Why does Claude ask again tomorrow?
Long-Term Memory for AI coding agents — Claude Code, OpenCode, and Pi
Persistent semantic memory that survives every session, every update, every compaction.
OpenLTM was born as a private memory layer for Claude Code. Today it's fully open source under MIT.
Same engine — automatic capture, semantic recall, importance-weighted decay, a queryable memory graph — now yours to read, fork, and extend. What started as one developer's "stop re-explaining my codebase" plugin is now an open foundation for agent memory across Claude Code, OpenCode, and Pi.
🔓 MIT licensed — no cloud, no account, no telemetry. Your memory lives in a local SQLite DB you own.
🧩 Provider-agnostic — one core ( @rohirik/openltm-core ), thin adapters per host.
🛠 Hackable — hooks, skills, janitor providers, and the graph visualizer are all in the open.
Migrating from an earlier install? The marketplace is now RohiRIK/OpenLtm and the plugin is openltm . Your existing memory database carries over.
Memory should be automatic. Hooks do the work. The session end hook extracts patterns, the session start hook injects them back. You shouldn't have to remember to remember.
Decay is a feature, not a bug. A gotcha from six months ago that you never revisited probably no longer applies. Set importance: 5 to make something permanent — everything else ages out naturally.
Semantic over keyword. FTS5 full-text search runs first; if it returns nothing, vector embeddings kick in. You search by meaning, not exact words — "how we handle async errors" finds the right memory even if you never wrote those exact words.
Zero config, zero lock-in. Install once, works everywhere. Every setting has a sane default. The DB lives outside the plugin directory so it survives every update. No cloud, no telemetry, no account.
🔍 Recall
Past decisions, patterns, and gotchas — before you start work
🧠 Learn
Every session, automatically — no manual note-taking
💉 Inject
Top context at session start so Claude picks up where it left off
⏳ Decay
Stale memories fade while critical knowledge lives forever
🕸 Graph
Traverse relationships between memories for reasoning chains
🗺 Visualize
See your entire memory network in a browser-based explorer
⚡ Vec Recall
Semantic vector (KNN) recall via sqlite-vec; degrades to JS-cosine when unavailable
🔌 Extensions
sqlite-vec + Honker (queue/cron/pub-sub) loaded dynamically; graceful fallback without system libsqlite3
Install
claude plugin marketplace add https://github.com/RohiRIK/OpenLtm
claude plugin install openltm
Restart Claude Code. That's it. Four Claude Code hooks + one git post-commit hook auto-wire, four commands load, five skills activate, and your openltm.db migrates or creates itself.
bunx @rohirik/openltm-core # auto-detect Claude Code, OpenCode
bunx @rohirik/openltm-core --pi # experimental Pi adapter
bunx @rohirik/openltm-core --dry-run --claude # preview without writing
Dev / git clone
git clone https://github.com/RohiRIK/OpenLtm ~ /Projects/OpenLtm
cd ~ /Projects/OpenLtm && bash install.sh
Quick Start
Start a new session. Context is injected at the top automatically.
/openltm:memory recall auth — what do we know about auth in this project?
/openltm:memory learn <insight> — save something worth keeping
/openltm:health — memory health + decay summary
/openltm:project init — set a goal for the current project
That's it. The rest is hooks doing the work.
Claude Code
│
├── 4 Commands ──┐
├── 5 Skills ──┼──▶ openltm MCP server ──▶ openltm.db
└── 5 Hooks ──┘ (memories, tags,
context_items,
memory_relations,
memories_fts)
Full deep-dive — schema, hook architecture, decay formula, ADRs — in How It Works and Architecture .
SQLite Extensions: The plugin loads sqlite-vec (vec0 / KNN vector search) and Honker (async embedding queue, leader-elected cron, pub-sub) when a system extension-enabled libsqlite3 is available. Both degrade gracefully — missing binary or library leaves the capability off and falls back to JS-cosine, file-watch polling, or in-process cron. Controlled with LTM_DISABLE_VEC , LTM_DISABLE_HONKER , LTM_SQLITE_LIB , and LTM_HONKER_EXT env vars.
/openltm:health # plugin health + hooks + decay
/openltm:memory recall test # returns results (or "no results" on fresh install)
Start a new session — you should see context injected at the top. If not, run /openltm:health to diagnose.
Full documentation index: docs/ .
Open an issue first to discuss the change, then send a PR. The full workflow — setup, tests, the version-bump rule, and release steps — is in CONTRIBUTING.md .
Powered by caffeine, SQLite, and the persistent belief that context shouldn't die at the end of a session.
Long-Term Memory plugin for Claude Code — semantic search, context injection, session learning
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
12
v2.9.1
Latest
Jun 9, 2026
+ 11 releases
Sponsor this project
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
