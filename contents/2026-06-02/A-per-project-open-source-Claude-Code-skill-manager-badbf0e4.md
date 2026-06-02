---
source: "https://github.com/narendranag/skillkit"
hn_url: "https://news.ycombinator.com/item?id=48361936"
title: "A per-project open-source Claude Code skill manager"
article_title: "GitHub - narendranag/skillkit: A per-project Claude Code skill manager. Curate a hybrid registry of skills — yours, gstack's, team collections — and install chosen subsets (individually or as packs) into a project's .claude/skills/ · GitHub"
author: "narendranag"
captured_at: "2026-06-02T00:32:32Z"
capture_tool: "hn-digest"
hn_id: 48361936
score: 2
comments: 1
posted_at: "2026-06-01T20:08:20Z"
tags:
  - hacker-news
  - translated
---

# A per-project open-source Claude Code skill manager

- HN: [48361936](https://news.ycombinator.com/item?id=48361936)
- Source: [github.com](https://github.com/narendranag/skillkit)
- Score: 2
- Comments: 1
- Posted: 2026-06-01T20:08:20Z

## Translation

タイトル: プロジェクトごとのオープンソース Claude Code スキル マネージャー
記事のタイトル: GitHub - narendranag/skillkit: プロジェクトごとの Claude Code スキル マネージャー。スキルのハイブリッド レジストリ (自分のスキル、gstack のスキル、チームのコレクション) をキュレーションし、選択したサブセットを (個別にまたはパックとして) プロジェクトの .claude/skills/ · GitHub にインストールします。
説明: プロジェクトごとの Claude Code スキル マネージャー。スキルのハイブリッド レジストリ (自分の、gstack の、チームのコレクション) を管理し、選択したサブセットを (個別にまたはパックとして) プロジェクトの .claude/skills/ - narendranag/skillkit にインストールします。

記事本文:
GitHub - narendranag/skillkit: プロジェクトごとの Claude Code スキル マネージャー。スキルのハイブリッド レジストリ (自分のスキル、gstack のスキル、チームのコレクション) をキュレーションし、選択したサブセットを (個別にまたはパックとして) プロジェクトの .claude/skills/ · GitHub にインストールします。
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
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてリフレッシュしてください

あなたのセッション。
アラートを閉じる
{{ メッセージ }}
ナレンドラナグ
/
スキルキット
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
22 コミット 22 コミット docs docs src/ skillkit src/ skillkit testing testing .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CLAUDE.md CLAUDE.md README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示リポジトリ ファイルのナビゲーション
プロジェクトごとの Claude Code スキル マネージャー。スキルのハイブリッド レジストリ (自分のスキル、gstack のスキル、チームのコレクション) を管理し、選択したサブセットを (個別にまたはパックとして) プロジェクトの .claude/skills/ にインストールします。スキルは、システム全体のすべてのクロード コード セッションではなく、選択したプロジェクトでのみアクティブになります。
なぜそれが存在するのか。 gstack のようなフレームワークは、約 80 個のスキルを ~/.claude/skills/ に分散するため、プロジェクトに関係なく、すべてのセッションにロードされます。 skillkit はプロジェクトごとにスキルの選択を宣言的に行います。 skill.toml をコミットし、どこでも同期します。
ドクター
のために
チュートリアル: はじめに
インストール → gstack の分散 → パックの構築 → プロジェクトで使用
ハウツーガイド
タスクのレシピ: スキルの追加、パックの作成、gstack の採用、ベンダー、スキルの作成
参考資料
各コマンドと 3 つのファイル形式
デザイン説明
マニフェスト + 同期、ライブ パック、マーカー安全モデル、内部構造を採用する理由
設計仕様
オリジナルデザインの記録
インストール
git clone < this-repo > ~ /ai/skillkit
cd ~ /ai/スキルキット
uv ツールのインストール --editable 。
これにより、uv のツール shim ディレクトリを介してスキルキットが PATH に配置されます。インストール後に uv が PATH ヒントを出力する場合は、推奨されたディレクトリをシェルの $PATH に追加します。
レジストリの場所 — 次の順序で解決されます。
$SKILLKIT_REGISTRY 環境変数
デフォルト: ~/ai/skillki

t (複製されたリポジトリ自体)
コンセプト
どこに住んでいるのか
それは何ですか
ソース
レジストリ内のsources.toml
サブディレクトリがスキルである名前付きルート ディレクトリ。鉱山スキル、gstack のディレクトリ、共有チーム コレクションなどを宣言します。
カタログ
計算された
すべてのソース ルートの結合、つまり宣言されたすべてのソースにわたるすべてのスキル。
パック
レジストリ内のpacks/*.toml
スキル参照の名前付きバンドル (source:name )。 LIVE 参照: 同期ごとに再解決されるため、ソース内のスキルの更新が自動的に取得されます。
プロジェクトマニフェスト
.claude/skills.toml (コミット済み)
プロジェクトに必要なスキルとパックを記録します。これが意図です。チームメイトが水分補給できるようにコミットします。
具現化されたスキル
.claude/skills/ (デフォルトでは gitignore)
実際のスキル ディレクトリ。 sync によってコピーされます。スキルキット同期で水分補給します。
sources.toml の例
[ 出典 ]
私の = " ~/ai/スキルキット/スキル "
gstack = " ~/.claude/skills/gstack "
パックファイルの例 (packs/webapp.toml)
【パック】
名前 = "ウェブアプリ"
description = " 典型的な Web アプリ プロジェクトのスキル "
スキル = [ " Mine:qa " 、 " gstack:next-best-practices " 、 " gstack:review " ]
コマンド
コマンド
何をするのか
スキルキット追加 <ref>
スキルまたはパックをプロジェクトマニフェストに追加して同期する
スキルキット rm <ref>
マニフェストからスキルまたはパックを削除して同期する
スキルキットピック
テキストの複数選択 TUI — カタログを参照し、スキル/パックを切り替えます
スキルキットの同期
現在のマニフェストを .claude/skills/ に具体化します。
スキルキットのアップデート
git でレジストリをプルしてから同期します
スキルキットベンダー
スキルをコピーし、コミット済みとしてマークします (「インストール モード」を参照)
スキルキットリスト
インストールされているスキル/パックの横に [*] が付いたカタログを表示します
スキルキット パック作成 <名前>
新しいパックを対話的に構築します (テキスト ピッカー → Packs/<name>.toml)
スキルキットパックリスト
レジストリ内のすべてのパックを一覧表示します
スキルキット パックのショー <名前>
パックの詳細を表示する

スキルキットはgstackを採用
gstack のディレクトリをソースとして登録し、分散したディレクトリをプレビューする (ドライラン)
スキルキットは gstack を採用します --はい
同じです。さらに、散在する gstack ディレクトリを macOS のゴミ箱に移動します。
参照構文:
source:name — 特定のスキル、例:私:qa、gstack:review
Pack:name — パック全体、例:パック:ウェブアプリ
デフォルト (gitignored、同期時にリハイドレート)
.claude/skills/ は .gitignore にリストされています。各チームメイトまたは CI 環境は、チェックアウト後にスキルキット同期を実行して、スキルをローカルに具体化します。 skill.toml のみがコミットされます。
skillkit add gstack:review # .claude/skills.toml に追加し、すぐに具体化します
skillkit sync # 新たなチェックアウト後にどこでも再実体化
ベンダー モード (リポジトリにコミットされたスキル)
自己完結型リポジトリの場合、またはスキルをリポジトリのスナップショットにバージョンロックしたい場合に便利です。
スキルキットベンダー
これにより、 sync が実行され、 !.claude/skills/ 否定行が .gitignore に書き込まれ、具体化されたスキル ディレクトリが git によって追跡されます。マニフェストと skill/ ディレクトリの両方をコミットします。
スキルは、SKILL.md ファイルが含まれるディレクトリです。フロントマターはメタデータを宣言します。
---
名前：マイスキル
description : カタログとピッカーに表示される短い説明
---
# 私のスキル
クロードへの指示…
スキル ディレクトリは、sources.toml で宣言されたルートの下に配置します。これらはスキルキット リストに表示され、TUI をすぐに選択できます。
gstack は、他のソースのスキルと一緒に ~/.claude/skills/ にスキルを分散してインストールします。 gstack を採用すると、これが統合されます。
~/.claude/skills/gstack を gstack ソースとしてsources.tomlに登録します。
「分散した」gstack スキル ディレクトリ ( ~/.claude/skills/ 内のディレクトリと名前が一致する ~/.claude/skills/ のトップレベル ディレクトリ) を検出します。
--yes を使用しない場合: ゴミ箱に移動されるディレクトリの予行演習リストが表示されます。 sources.toml を記述する以外のファイルシステムの変更はありません。
--yes を指定すると: 移動します

それらの散在ディレクトリは macOS のゴミ箱に置かれます (元に戻すことができます)。
必ず最初に予行演習を確認してください。スキャッター検出はディレクトリ名のみで一致します。gstack スキルと名前を共有する非 gstack スキルにはフラグが立てられます。
スキルキットは gstack を採用 # ドライラン — リストを確認する
スキルキットは gstack を採用します --はい # 承認されました — 散乱したディレクトリをゴミ箱に移動します
gstack-upgrade の実行後、gstack がスキルを再分散する可能性があります。 skillkit Adopt gstack --yes を再実行するだけです。
シークレット、内部 URL、およびクライアント名を skill/ の下のスキル ファイルに含めないでください。 gileaks のコミット前フックは .pre-commit-config.yaml で構成されます。クローン作成後に一度有効にします。
プリコミットインストール
フックは、プッシュされる前に、各コミットをスキャンして、漏洩した認証情報を探します。
uv pytest # 34 テストを実行します
について
プロジェクトごとの Claude Code スキル マネージャー。スキルのハイブリッド レジストリ (自分、gstack、チーム コレクション) を管理し、選択したサブセットを (個別にまたはパックとして) プロジェクトの .claude/skills/ にインストールします。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A per-project Claude Code skill manager. Curate a hybrid registry of skills — yours, gstack's, team collections — and install chosen subsets (individually or as packs) into a project's .claude/skills/ - narendranag/skillkit

GitHub - narendranag/skillkit: A per-project Claude Code skill manager. Curate a hybrid registry of skills — yours, gstack's, team collections — and install chosen subsets (individually or as packs) into a project's .claude/skills/ · GitHub
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
narendranag
/
skillkit
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
22 Commits 22 Commits docs docs src/ skillkit src/ skillkit tests tests .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CLAUDE.md CLAUDE.md README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
A per-project Claude Code skill manager. Curate a hybrid registry of skills — yours, gstack's, team collections — and install chosen subsets (individually or as packs) into a project's .claude/skills/ . Skills become active only in the projects you select them for, not in every Claude Code session system-wide.
Why it exists. Frameworks like gstack scatter ~80 skills into ~/.claude/skills/ so they load in every session, regardless of project. skillkit makes skill selection per-project and declarative: commit a skills.toml , sync anywhere.
Doc
For
Tutorial: Getting started
Install → de-scatter gstack → build a pack → use it in a project
How-to guides
Task recipes: add skills, create packs, adopt gstack, vendor, author a skill
Reference
Every command and the three file formats
Design explanation
Why manifest+sync, live packs, the marker-safety model, adopt internals
Design spec
The original design record
Install
git clone < this-repo > ~ /ai/skillkit
cd ~ /ai/skillkit
uv tool install --editable .
This puts skillkit on your PATH via uv's tool shim directory. If uv prints a PATH hint after install, add the suggested directory to your shell's $PATH .
Registry location — resolved in this order:
$SKILLKIT_REGISTRY environment variable
Default: ~/ai/skillkit (the cloned repo itself)
Concept
Where it lives
What it is
Source
sources.toml in the registry
A named root directory whose sub-dirs are skills. Declare your mine skills, gstack's dir, a shared team collection, etc.
Catalog
computed
Union of all source roots — every skill across all declared sources.
Pack
packs/*.toml in the registry
A named bundle of skill refs ( source:name ). LIVE references: re-resolved on every sync , so updating a skill in a source is automatically picked up.
Project manifest
.claude/skills.toml (committed)
Records which skills and packs a project wants. This is the intent ; commit it so teammates can rehydrate.
Materialized skills
.claude/skills/ (gitignored by default)
The actual skill directories, copied in by sync . Rehydrate with skillkit sync .
sources.toml example
[ sources ]
mine = " ~/ai/skillkit/skills "
gstack = " ~/.claude/skills/gstack "
Pack file example ( packs/webapp.toml )
[ pack ]
name = " webapp "
description = " Skills for a typical web app project "
skills = [ " mine:qa " , " gstack:next-best-practices " , " gstack:review " ]
Commands
Command
What it does
skillkit add <ref>
Add a skill or pack to the project manifest and sync
skillkit rm <ref>
Remove a skill or pack from the manifest and sync
skillkit pick
Textual multiselect TUI — browse the catalog and toggle skills/packs
skillkit sync
Materialize the current manifest into .claude/skills/
skillkit update
git pull the registry, then sync
skillkit vendor
Copy skills in and mark them committed (see Install Modes)
skillkit list
Show catalog with [*] next to installed skills/packs
skillkit pack create <name>
Interactively build a new pack (Textual picker → packs/<name>.toml )
skillkit pack list
List all packs in the registry
skillkit pack show <name>
Show a pack's details
skillkit adopt gstack
Register gstack's dir as a source and preview scattered dirs (dry run)
skillkit adopt gstack --yes
Same, plus move scattered gstack dirs to macOS Trash
Ref syntax:
source:name — a specific skill, e.g. mine:qa , gstack:review
pack:name — an entire pack, e.g. pack:webapp
Default (gitignored, rehydrated on sync)
.claude/skills/ is listed in .gitignore . Each teammate or CI environment runs skillkit sync after checkout to materialize skills locally. Only skills.toml is committed.
skillkit add gstack:review # adds to .claude/skills.toml, materializes immediately
skillkit sync # re-materialize anywhere after a fresh checkout
Vendor mode (skills committed to the repo)
Useful for self-contained repos or when you want skills version-locked to the repo snapshot:
skillkit vendor
This runs sync then writes a !.claude/skills/ negation line to .gitignore , so the materialized skills directory is tracked by git. Commit both the manifest and the skills/ directory.
A skill is a directory containing a SKILL.md file. Front-matter declares metadata:
---
name : my-skill
description : Short description shown in catalog and picker
---
# My Skill
Instructions for Claude...
Place skill directories under a root declared in sources.toml . They appear in skillkit list and the pick TUI immediately.
gstack installs its skills scattered into ~/.claude/skills/ alongside skills from other sources. adopt gstack consolidates this:
Registers ~/.claude/skills/gstack as a gstack source in sources.toml .
Detects "scattered" gstack skill dirs — top-level dirs in ~/.claude/skills/ whose names match a dir inside ~/.claude/skills/gstack/ .
Without --yes : shows the dry-run list of dirs that would be moved to Trash. No filesystem changes beyond writing sources.toml .
With --yes : moves those scattered dirs to macOS Trash (reversible).
Always review the dry-run first. Scatter detection matches by directory name only — a non-gstack skill sharing a name with a gstack skill will be flagged.
skillkit adopt gstack # dry run — review the list
skillkit adopt gstack --yes # approved — move scattered dirs to Trash
After running gstack-upgrade , gstack may re-scatter skills. Just re-run skillkit adopt gstack --yes .
Keep secrets, internal URLs, and client names out of skill files under skills/ . A gitleaks pre-commit hook is configured in .pre-commit-config.yaml . Enable it once after cloning:
pre-commit install
The hook scans each commit for leaked credentials before they are pushed.
uv run pytest # 34 tests
About
A per-project Claude Code skill manager. Curate a hybrid registry of skills — yours, gstack's, team collections — and install chosen subsets (individually or as packs) into a project's .claude/skills/
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
