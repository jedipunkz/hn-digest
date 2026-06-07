---
source: "https://pypi.org/project/seed-cli/"
hn_url: "https://news.ycombinator.com/item?id=48434032"
title: "Show HN: Create and Maintain Filesystem Structures for LLMs [v1.0.11 Out]"
article_title: "seed-cli · PyPI"
author: "hunterx"
captured_at: "2026-06-07T12:38:57Z"
capture_tool: "hn-digest"
hn_id: 48434032
score: 2
comments: 0
posted_at: "2026-06-07T12:11:03Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Create and Maintain Filesystem Structures for LLMs [v1.0.11 Out]

- HN: [48434032](https://news.ycombinator.com/item?id=48434032)
- Source: [pypi.org](https://pypi.org/project/seed-cli/)
- Score: 2
- Comments: 0
- Posted: 2026-06-07T12:11:03Z

## Translation

タイトル: HN の表示: LLM のファイルシステム構造の作成と保守 [v1.0.11 リリース]
記事タイトル：seed-cli・PyPI
説明: 「tree」プログラムに最適なペアです。シードは、コードベース内のファイルとフォルダーの構造を作成、更新、削除します。

記事本文:
メインコンテンツにスキップ
モバイル版に切り替える
警告
サポートされていないブラウザを使用しているため、新しいバージョンにアップグレードしてください。
PyPIを検索
検索フォーカス#フォーカス検索フィールド"
data-search-focus-target="検索フィールド">
検索
ヘルプ
pip インストール シード-cli
PIP 命令のコピー
「tree」プログラムに最適な組み合わせです。シードは、コードベース内のファイルとフォルダーの構造を作成、更新、削除します。
ライセンス: MIT ライセンス (ModifiedMIT)
タグ
ディレクトリ
、
ファイルシステム
、
オーケストレーション
、
種子
、
テラフォーム
、
木
追加の提供:
開発者
、画像
、うい
ライセンス
OSI承認済み :: MITライセンス
オペレーティングシステム
OSに依存しない
プログラミング言語
パイソン :: 3
トピック
ソフトウェア開発 :: ビルドツール
Seed は、Terraform からインスピレーションを得た仕様主導のファイルシステム オーケストレーション ツールです。
ディレクトリツリーをキャプチャし、変更を計画し、それらを安全に適用し、ドリフトを同期し、
再利用可能なテンプレートからプロジェクトを足場にし、実行することもできます
マニフェスト主導のリポジトリまたはサービスのメンテナンス。
Terraform をディレクトリ ツリーとテンプレートのスキャフォールディングに加えて考えてみましょう。
ワークスペースのメンテナンス。
複数の仕様入力: .tree 、 .seed 、 YAML、JSON、DOT、画像 OCR、および標準入力
エクスポート可能な計画を使用した確定的計画: シード計画 spec.seed --out plan.json
不変プランの安全な実行: シード適用 plan.json
ドリフトワークフロー: diff 、 sync 、 match 、スナップショット、仕様履歴
パスとコンテンツのテンプレート変数: <varname>/ および {{var}}
.seed/templates/ でのプロジェクト ローカル テンプレートの登録
バージョン、ロック、コンテンツ ソース、組み込みテンプレートを含む再利用可能なテンプレート レジストリ
コピー機スタイルのテンプレート構成のサポート: 質問、デフォルト、回答ファイル、除外、存在する場合はスキップ、およびゲートされたタスク
シード保守によるマニフェスト主導のリポジトリ/サービス/システム保守
構造ロック、監視モード、状態ロック、フック、Graphviz エクスポート、およびシェル完了

ション
pip インストール シード-cli
pip install "seed-cli[image]" # OCR/画像解析
pip install "seed-cli[ui]" # 豊富なターミナル出力
Python 3.10 以上が必要です。
既存のディレクトリをキャプチャし、プランをプレビューしてから適用します。
シード キャプチャ --out project.tree
シード プラン project.tree --out plan.json
シード適用 plan.json
仕様を直接実行する場合:
シード レジスタ プロジェクト.ツリー
シード適用プロジェクト.ツリー
シード差分プロジェクト.ツリー
リポジトリまたはサービスの自動化の場合:
シード保守メンテナンス.yml
シード維持、maintenance.yml --execute
コマンド
シード --help はワークフローごとにコマンドをグループ化するため、関連タスクを簡単にスキャンできます。
シード プラン dir_structure.tree --out plan.json
シード適用 plan.json
シード適用は仕様を直接受け入れることもできます。
シード適用 dir_ Structure.tree
<name>/ などのテンプレート プレースホルダーを使用して仕様を適用する場合、または
features/<name>.ts 、シード適用
最初にシード レジスタ セマンティクスを実行します。サポート ファイルを以下に書き込みます。
.seed/templates/ および .seed/templates/project/ 、古いものを削除します
古い実行によって残された features/<name>/ のようなリテラルのプレースホルダー パス。
シード差分 dir_structure.tree
シード同期 dir_ Structure.tree --dangerous
シード マッチ dir_ Structure.tree --dangerous
sync は仕様にない余分なものを削除します。 match では欠落したパスも作成されますが、
... でマークされたディレクトリを考慮します。
シード プラン dir_ Structure.tree --target scripts/
シード プラン dir_structure.tree --targets "services/*"
シード プラン dir_structure.tree --target-mode 正確
仕様の構文
単純なファイルシステム仕様には .tree を使用し、同じツリー形式に加えて種類、タグ、URL などの豊富なインライン メタデータが必要な場合は .seed を使用します。
@includebase.tree
スクリプト/
§── build.py @generated
§──notes.txt @manual
§── キャッシュ/ ...
━── ドキュメント/ ?
マーカー
@include file.tree : 別の仕様を含めます
@manual : 手動でメンテナンスされます

ファイル
？ : オプションのファイルまたはディレクトリ
... : ディレクトリ内のエクストラを許可します
<varname> : パスセグメントまたはファイル名のテンプレートプレースホルダー
{{var}} : ファイル内容の変数補間
.seed はインライン メタデータ マーカーもサポートしています。
!kind : !service 、 !doc 、 !template などの意味論的な種類マーカー
+tag : +remote +shared などの反復可能なタグ
-> URL : メタデータ URL をノードに添付します。ディレクトリノードではこれは可能です
テンプレートコンテンツソースとして使用される
ベンダー/
━── api/ !service +remote -> https://github.com/acme/repo.git
構造化された YAML および JSON 仕様では、次のいずれかのメタデータを運ぶこともできます。
メタデータ オブジェクトまたはトップレベルの kind 、 tags 、および URL フィールド。
シード プランの spec.tree --vars project_name = myapp
シード適用 spec.tree --vars project_name = myapp
シード作成 spec.tree project_name = myapp
テンプレート
テンプレート変数を使用して繰り返し構造を定義します。
ファイル/
§── <version_id>/
│ §── data.json
│ └─ メタ/
━──
プレースホルダーは、行ごとのパスの仕様、ファイル名、およびテンプレートにも使用できます。
複数のプレースホルダーを含めることができます。
features/<ドメイン>/<名前>/route.ts
features/<名前>.ts
インスタンスを作成します。
シード作成 releases.tree version_id = v3
シード作成component.treeドメイン=請求先名=請求書
シード作成 releases.tree version_id = v3 --dry-run
プロジェクトローカルテンプレート
シード レジスタを使用して、.tree または .seed 仕様をプロジェクト レベルにミラーリングします。
.seed/templates/ ディレクトリ。仕様のどこかにプレースホルダーが含まれている場合、
パスに加えて、最も外側のプレースホルダー サブツリーも抽出されます。
.seed/templates/project/ 。シード適用 <spec> は同じ登録ステップを実行します
実行前に自動的に実行されます。
シード レジスタのリリース.ツリー
シード適用リリース.ツリー
これにより、パスベースのプロジェクト テンプレートから作成できます。
シード作成 --template .seed/templates/r

eleases.tree version_id = v3
または登録されたプロジェクト テンプレート名:
シード作成 --project version_id version_id = v3
テンプレートレジストリ
デフォルトで ~/.seed/templates/ に保存されている再利用可能なテンプレートを管理するか、
SEED_HOME が設定されている場合は、$SEED_HOME/templates/ の下にあります。
シードテンプレートのリスト
シード テンプレートは ./template.tree --name my-template を追加します
シード テンプレートは ./template.seed --name service-template を追加します
シード テンプレートは my-template を表示します
シード テンプレートは my-template target-folder を使用します
シード テンプレートのバージョン my-template --add ./updated.tree --name v2
シード テンプレートが私のテンプレートをロックする
シード テンプレートが my-template を更新する
シード テンプレートは、my-template を削除します
シード テンプレート リストには、から検出されたプロジェクト ローカル テンプレートも表示されます。
.seed/templates/project/ 、グローバル レジストリ テンプレートの前。シード テンプレートは <name> <folder> を使用します。最初に表示されているプロジェクト テンプレートを解決し、その後フォールバックします。
グローバルレジストリに。単数形のエイリアス シード テンプレートの使用 ... も
受け入れられました。
組み込みテンプレートには、 fastapi 、 python-package 、および node-typescript が含まれます。
テンプレートは、ローカル ディレクトリ、GitHub ツリー URL、または git を指すことができます。
リポジトリ URL。シードが構造とともに実際のファイルの内容をフェッチできるようにする
仕様。リポジトリ ソースは、.git メタデータなしで複製されます。
シード テンプレートは ./fastapi --name fastapi \ を追加します
--content-url https://github.com/tiangolo/full-stack-fastapi-template/tree/master/backend/app
シード テンプレートは ./service.seed --name service \ を追加します
--content-url https://github.com/acme/service-skeleton.git
シード テンプレートが fastapi を更新する
シード テンプレートの更新 --all
シード テンプレートは fastapi --content-url /path/to/local/files を更新します
{"content_url": "..."} を含むsource.json ファイルを含むテンプレートは次のとおりです。
インストール時に自動的に取得されます。 .seed ディレクトリ ノードでも宣言できます
独自のコンテンツ ソースをインラインで使用します。
ベンダー/

━── api/ !service +remote -> https://github.com/acme/api-client.git
コピー機式足場
シード テンプレートは、copier.yml という名前のサポート テンプレート設定ファイルを使用します。
copier.yaml 、 .seed-template.yml 、または .seed-template.yaml 。
サポートされているワークフロー機能は次のとおりです。
プロンプト表示される質問とデフォルト
JSON/YAML 応答の --data-file
--defaults および --non-interactive
--answers-file または _answers_file
_tasks 、 --unsafe でのみ実行されます
--既存のファイルを上書きします
シード テンプレートは python-package を使用します \
--base ./myapp \
--データファイルanswers.yml \
--デフォルト \
--answers-file .seed/answers.yml \
--上書き
テンプレートで _tasks が定義されている場合、それらは表示されますが、オプトインしない限りスキップされます。
実行:
シード テンプレートは python-package --unsafe を使用します
リポジトリとシステムのメンテナンス
シード保守はリポジトリ、サービス、システム、プロジェクトの維持管理を調整します
YAML または JSON マニフェストから。
組み込みのメンテナンス目標には次のものが含まれます。
リポジトリ: ensure_path 、 git_fetch 、 git_status 、 git_pull_ff_only
サービス: ensure_paths 、 compose_pull 、 compose_up 、 launchctl_restart
tool 、 args 、 cwd 、 env 、またはシェル コマンドを使用したカスタム アクション
マニフェスト ファイルまたは次のディレクトリを指定してシードを維持できます。
maintenance.yml 、 project.yml 、または service.yml 。
ターゲット :
- 名前 : シード-cli
種類 : リポジトリ
パス: ./repos/seed-cli
目標:
- パスの確保
- git_fetch
- git_status
- 名前 : メモ API
種類：サービス
パス: ./systems/services/notes-api
config_dir : ./systems/services/notes-api/config
data_dir : ./local/services/notes-api
compose_file : compose.yml
デプロイエンジン: docker-compose
launch_agent : user/com.example.notes-api
目標:
- パスの確保
- compose_pull
- 構成アップ
- launchctl_restart
プロジェクト.yml
名前：プロダクトx
タイプ: プロジェクト
パス: ~/work/projects/active/product-x
メンテナンス：
目標:
- git_fe

ち
- git_status
リポジトリ:
- 名前 : ウェブアプリ
パス : リポジトリ/web-app
- 名前 : アピ
パス : リポジトリ/API
サービス.yml
名前：notes-api
タイプ : サービス
パス: ~/systems/services/notes-api
config_dir : ~/systems/services/notes-api/config
data_dir : ~/local/services/notes-api
compose_file : compose.yml
デプロイエンジン: docker-compose
launch_agent : user/com.example.notes-api
メンテナンス：
目標:
- パスの確保
- compose_pull
- 構成アップ
- launchctl_restart
アクション:
- ツール：Python
引数: [ "scripts/rebuild_index.py" ]
cwd : "{{パス}}"
まずプランナーを実行してから、以下を実行します。
シード維持 ./workspace
シード維持 ./workspace --execute
スナップショット、仕様履歴、ロック
スナップショットは、 apply 、 sync 、 match の前に自動的に作成されます。
シードを元に戻す --list
シードを元に戻す
シード復帰 abc123 --dry-run
適用された構造もバージョン化された仕様としてキャプチャされます。
シードスペックリスト
シードスペックショー
シード仕様差分 v1 v3
シードスペックウォッチ
シードスペックウォッチはワークスペースをポーリングし、新しい .seed/specs/vN.tree を書き込みます。
ファイルシステム構造が変更されるたびに、変更後に手動でファイルを作成する必要があります。
最初の適用により、内部基準が自動的に進められます。
ファイルシステム構造をロックし、ドリフトがないか監視します。
シードロックセット仕様ツリー
シードロックリスト
シードロックステータス
シードロックウォッチ
シード ロック アップグレード v2 --dry-run
シード ロック ダウングレード v1 --危険
エクスポート、フック、およびユーティリティ
現在の状態または計画をエクスポートします。
シード エクスポート ツリー --out Structure.tree
シード エクスポート json --out 構造.json
シード エクスポート ドット --out Structure.dot
シード プランの spec.tree --dot > plan.dot
git フックをインストールします。
シードフックを取り付ける
シードフックのインストール --hook pre-push
ユーティリティ:
シード ユーティリティ 抽出ツリー スクリーンショット.png --out spec.tree
シードユーティリティの状態ロック
シード utils 状態ロック --force-unlock
シェルのオートコンプリート
Click は Bash、Zsh、Fis のシェル補完を提供します

h.完了を生成する
インストールされたシード エントリ ポイントからのスクリプト:
#zsh
eval " $( _SEED_COMPLETE = zsh_source シード ) "
#バッシュ
eval " $( _SEED_COMPLETE = bash_source シード ) "
# 魚
_SEED_COMPLETE = 魚ソース シード |ソース
次に、シェルをリロードし、タブ補完を使用します。
シード<TAB>
シード テンプレート <TAB>
シードロック <TAB>
安全モデル
シードはデフォルトで安全になるように設計されています。
破壊的なワークフローには明示的な危険フラグが必要です
実行状態はハートビート更新によりロック保護されています
プランは書き込みと削除の前に検証されます
テンプレートタスクには明示的な --unsafe が必要です
応答ファイルと実行ターゲットはベース ディレクトリに制限されます
git メンテナンスがダーティ ワークツリーで git pull --ff-only を拒否する
MITを修正しました。 LICENSE.md を参照してください。
ライセンス: MIT ライセンス (ModifiedMIT)
タグ
ディレクトリ
、
ファイルシステム
、
オーケストレーション
、
種子
、
テラフォーム
、
木
追加の提供:
開発者
、画像
、うい
ライセンス
OSI承認済み :: MITライセンス
オペレーティングシステム
OSに依存しない
プログラミング言語
パイソン :: 3
トピック
ソフトウェア開発 :: ビルドツール
発売履歴
リリースのお知らせ |
RSSフィード
ご使用のプラットフォーム用のファイルをダウンロードします。どれを選択すればよいかわからない場合は、パッケージのインストールについて詳しくご覧ください。
名前、インタープリター、ABI、プラットフォームでファイルをフィルターします。
ファイル名がわからない場合は、

[切り捨てられた]

## Original Extract

The perfect pair for the `tree` program. Seed creates, updates and deletes file and folder structures in your codebase.

Skip to main content
Switch to mobile version
Warning
You are using an unsupported browser, upgrade to a newer version.
Search PyPI
search-focus#focusSearchField"
data-search-focus-target="searchField">
Search
Help
pip install seed-cli
Copy PIP instructions
The perfect pair for the `tree` program. Seed creates, updates and deletes file and folder structures in your codebase.
License: MIT License (ModifiedMIT)
Tags
directory
,
filesystem
,
orchestration
,
seed
,
terraform
,
tree
Provides-Extra:
dev
, image
, ui
License
OSI Approved :: MIT License
Operating System
OS Independent
Programming Language
Python :: 3
Topic
Software Development :: Build Tools
seed is a Terraform-inspired, spec-driven filesystem orchestration tool.
It captures directory trees, plans changes, applies them safely, syncs drift,
scaffolds projects from reusable templates, and can also execute
manifest-driven repository or service maintenance.
Think Terraform for directory trees , plus template scaffolding and
workspace maintenance .
Multiple spec inputs: .tree , .seed , YAML, JSON, DOT, image OCR, and stdin
Deterministic planning with exportable plans: seed plan spec.seed --out plan.json
Safe execution of immutable plans: seed apply plan.json
Drift workflows: diff , sync , match , snapshots, and spec history
Template variables in paths and content: <varname>/ and {{var}}
Project-local template registration under .seed/templates/
Reusable template registry with versions, locking, content sources, and built-in templates
Copier-style template config support: questions, defaults, answers files, excludes, skip-if-exists, and gated tasks
Manifest-driven repository/service/system maintenance with seed maintain
Structure locking, watch mode, state locks, hooks, Graphviz export, and shell completion
pip install seed-cli
pip install "seed-cli[image]" # OCR/image parsing
pip install "seed-cli[ui]" # rich terminal output
Python >=3.10 is required.
Capture an existing directory, preview a plan, then apply it:
seed capture --out project.tree
seed plan project.tree --out plan.json
seed apply plan.json
For direct spec execution:
seed register project.tree
seed apply project.tree
seed diff project.tree
For repository or service automation:
seed maintain maintenance.yml
seed maintain maintenance.yml --execute
Commands
seed --help groups commands by workflow so related tasks are easier to scan.
seed plan dir_structure.tree --out plan.json
seed apply plan.json
seed apply also accepts a spec directly:
seed apply dir_structure.tree
When you apply a spec with template placeholders such as <name>/ or
features/<name>.ts , seed apply
first runs seed register semantics: it writes the supporting files under
.seed/templates/ and .seed/templates/project/ , then removes any stale
literal placeholder paths like features/<name>/ left behind by older runs.
seed diff dir_structure.tree
seed sync dir_structure.tree --dangerous
seed match dir_structure.tree --dangerous
sync deletes extras not in the spec. match also creates missing paths while
respecting directories marked with ... .
seed plan dir_structure.tree --target scripts/
seed plan dir_structure.tree --targets "services/*"
seed plan dir_structure.tree --target-mode exact
Spec Syntax
Use .tree for simple filesystem specs, or .seed when you want the same tree-shaped format plus richer inline metadata such as kinds, tags, and URLs.
@include base.tree
scripts/
├── build.py @generated
├── notes.txt @manual
├── cache/ ...
└── docs/ ?
Markers
@include file.tree : include another spec
@manual : manually maintained file
? : optional file or directory
... : allow extras inside a directory
<varname> : template placeholder in a path segment or filename
{{var}} : variable interpolation in file contents
.seed also supports inline metadata markers:
!kind : semantic kind marker such as !service , !doc , or !template
+tag : repeatable tags such as +remote +shared
-> URL : attach a metadata URL to a node; on directory nodes this can be
used as a template content source
vendor/
└── api/ !service +remote -> https://github.com/acme/repo.git
Structured YAML and JSON specs can also carry metadata with either a
metadata object or top-level kind , tags , and url fields.
seed plan spec.tree --vars project_name = myapp
seed apply spec.tree --vars project_name = myapp
seed create spec.tree project_name = myapp
Templates
Define repeating structures with template variables:
files/
├── <version_id>/
│ ├── data.json
│ └── meta/
└── ...
Placeholders can also appear in path-per-line specs or filenames, and templates
can include more than one placeholder:
features/<domain>/<name>/route.ts
features/<name>.ts
Create instances:
seed create releases.tree version_id = v3
seed create component.tree domain = billing name = invoices
seed create releases.tree version_id = v3 --dry-run
Project-Local Templates
Use seed register to mirror any .tree or .seed spec into the project-level
.seed/templates/ directory. When the spec contains placeholders anywhere in a
path, it also extracts the outermost placeholder subtree into
.seed/templates/project/ . seed apply <spec> runs the same registration step
automatically before execution.
seed register releases.tree
seed apply releases.tree
That lets you create from either a path-based project template:
seed create --template .seed/templates/releases.tree version_id = v3
or a registered project template name:
seed create --project version_id version_id = v3
Template Registry
Manage reusable templates stored under ~/.seed/templates/ by default, or
under $SEED_HOME/templates/ when SEED_HOME is set.
seed templates list
seed templates add ./template.tree --name my-template
seed templates add ./template.seed --name service-template
seed templates show my-template
seed templates use my-template target-folder
seed templates versions my-template --add ./updated.tree --name v2
seed templates lock my-template
seed templates update my-template
seed templates remove my-template
seed templates list also shows project-local templates discovered from
.seed/templates/project/ , before global registry templates. seed templates use <name> <folder> resolves a visible project template first, then falls back
to the global registry. The singular alias seed template use ... is also
accepted.
Built-in templates include fastapi , python-package , and node-typescript .
Templates can point at a local directory, a GitHub tree URL, or a git
repository URL so seed can fetch real file contents alongside the structure
spec. Repository sources are cloned without their .git metadata.
seed templates add ./fastapi --name fastapi \
--content-url https://github.com/tiangolo/full-stack-fastapi-template/tree/master/backend/app
seed templates add ./service.seed --name service \
--content-url https://github.com/acme/service-skeleton.git
seed templates update fastapi
seed templates update --all
seed templates update fastapi --content-url /path/to/local/files
Templates that include a source.json file with {"content_url": "..."} are
fetched automatically when installed. .seed directory nodes can also declare
their own content sources inline:
vendor/
└── api/ !service +remote -> https://github.com/acme/api-client.git
Copier-Style Scaffolding
seed templates use supports template config files named copier.yml ,
copier.yaml , .seed-template.yml , or .seed-template.yaml .
Supported workflow features include:
promptable questions and defaults
--data-file for JSON/YAML answers
--defaults and --non-interactive
--answers-file or _answers_file
_tasks , which only execute with --unsafe
--overwrite for existing files
seed templates use python-package \
--base ./myapp \
--data-file answers.yml \
--defaults \
--answers-file .seed/answers.yml \
--overwrite
If a template defines _tasks , they are shown but skipped unless you opt into
execution:
seed templates use python-package --unsafe
Repository & System Maintenance
seed maintain orchestrates repository, service, system, and project upkeep
from YAML or JSON manifests.
Built-in maintenance goals include:
repositories: ensure_path , git_fetch , git_status , git_pull_ff_only
services: ensure_paths , compose_pull , compose_up , launchctl_restart
custom actions with tool , args , cwd , env , or shell commands
You can point seed maintain at a manifest file or a directory containing
maintenance.yml , project.yml , or service.yml .
targets :
- name : seed-cli
kind : repository
path : ./repos/seed-cli
goals :
- ensure_path
- git_fetch
- git_status
- name : notes-api
kind : service
path : ./systems/services/notes-api
config_dir : ./systems/services/notes-api/config
data_dir : ./local/services/notes-api
compose_file : compose.yml
deploy_engine : docker-compose
launch_agent : user/com.example.notes-api
goals :
- ensure_paths
- compose_pull
- compose_up
- launchctl_restart
project.yml
name : product-x
type : project
path : ~/work/projects/active/product-x
maintenance :
goals :
- git_fetch
- git_status
repos :
- name : web-app
path : repos/web-app
- name : api
path : repos/api
service.yml
name : notes-api
type : service
path : ~/systems/services/notes-api
config_dir : ~/systems/services/notes-api/config
data_dir : ~/local/services/notes-api
compose_file : compose.yml
deploy_engine : docker-compose
launch_agent : user/com.example.notes-api
maintenance :
goals :
- ensure_paths
- compose_pull
- compose_up
- launchctl_restart
actions :
- tool : python
args : [ "scripts/rebuild_index.py" ]
cwd : "{{path}}"
Run the planner first, then execute:
seed maintain ./workspace
seed maintain ./workspace --execute
Snapshots, Spec History, and Locks
Snapshots are created automatically before apply , sync , and match :
seed revert --list
seed revert
seed revert abc123 --dry-run
Applied structures are also captured as versioned specs:
seed specs list
seed specs show
seed specs diff v1 v3
seed specs watch
seed specs watch polls the workspace and writes a new .seed/specs/vN.tree
whenever the filesystem structure changes, so manual file creation after an
initial apply advances the internal reference automatically.
Lock a filesystem structure and watch it for drift:
seed lock set spec.tree
seed lock list
seed lock status
seed lock watch
seed lock upgrade v2 --dry-run
seed lock downgrade v1 --dangerous
Export, Hooks, and Utilities
Export current state or a plan:
seed export tree --out structure.tree
seed export json --out structure.json
seed export dot --out structure.dot
seed plan spec.tree --dot > plan.dot
Install git hooks:
seed hooks install
seed hooks install --hook pre-push
Utilities:
seed utils extract-tree screenshot.png --out spec.tree
seed utils state-lock
seed utils state-lock --force-unlock
Shell Autocomplete
Click provides shell completion for Bash, Zsh, and Fish. Generate the completion
script from the installed seed entry point:
# zsh
eval " $( _SEED_COMPLETE = zsh_source seed ) "
# bash
eval " $( _SEED_COMPLETE = bash_source seed ) "
# fish
_SEED_COMPLETE = fish_source seed | source
Then reload your shell and use tab completion:
seed <TAB>
seed templates <TAB>
seed lock <TAB>
Safety Model
seed is designed to be safe by default:
destructive workflows require explicit dangerous flags
execution state is lock-protected with heartbeat renewal
plans are validated before writes and deletes
template tasks require explicit --unsafe
answers files and execution targets are constrained to the base directory
git maintenance refuses git pull --ff-only on dirty worktrees
Modified MIT. See LICENSE.md .
License: MIT License (ModifiedMIT)
Tags
directory
,
filesystem
,
orchestration
,
seed
,
terraform
,
tree
Provides-Extra:
dev
, image
, ui
License
OSI Approved :: MIT License
Operating System
OS Independent
Programming Language
Python :: 3
Topic
Software Development :: Build Tools
Release history
Release notifications |
RSS feed
Download the file for your platform. If you're not sure which to choose, learn more about installing packages .
Filter files by name, interpreter, ABI, and platform.
If you're not sure about the file name fo

[truncated]
