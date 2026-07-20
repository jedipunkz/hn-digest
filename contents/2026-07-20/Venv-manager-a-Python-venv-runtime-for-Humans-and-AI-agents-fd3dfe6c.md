---
source: "https://github.com/jacopobonomi/venv_manager"
hn_url: "https://news.ycombinator.com/item?id=48979288"
title: "Venv-manager, a Python venv runtime for Humans and AI agents"
article_title: "GitHub - jacopobonomi/venv_manager: A powerful CLI tool for managing Python virtual environments with ease. · GitHub"
author: "twojacks"
captured_at: "2026-07-20T15:02:16Z"
capture_tool: "hn-digest"
hn_id: 48979288
score: 1
comments: 0
posted_at: "2026-07-20T14:23:26Z"
tags:
  - hacker-news
  - translated
---

# Venv-manager, a Python venv runtime for Humans and AI agents

- HN: [48979288](https://news.ycombinator.com/item?id=48979288)
- Source: [github.com](https://github.com/jacopobonomi/venv_manager)
- Score: 1
- Comments: 0
- Posted: 2026-07-20T14:23:26Z

## Translation

タイトル: Venv-manager、人間と AI エージェント用の Python venv ランタイム
記事タイトル: GitHub - jacopobonomi/venv_manager: Python 仮想環境を簡単に管理するための強力な CLI ツール。 · GitHub
説明: Python 仮想環境を簡単に管理するための強力な CLI ツール。 - ジャコポボノミ/venv_manager

記事本文:
GitHub - jacopobonomi/venv_manager: Python 仮想環境を簡単に管理するための強力な CLI ツール。 · GitHub
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
ヤコポボノミ
/
venv_manager
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション

n 個のオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
28 コミット 28 コミット .github .github cmd/ venv-manager cmd/ venv-manager 内部 内部スクリプト scripts .gitignore .gitignore CHANGELOG.md CHANGELOG.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md go.mod go.mod go.sum go.sum install.sh install.sh station_example.gifterminal_example.gif すべてのファイルを表示 リポジトリ ファイルのナビゲーション
人間と AI エージェントのための Python 仮想環境ランタイム。
Goで書かれています。 1 つの静的バイナリ。 python3 (または利用可能な場合は uv ) を超えるランタイム deps はありません。
上の GIF は本物です: venv-manager watch app.py --venv X はファイルを監視し、小さな AST-lite パーサーでそのインポートをスキャンし、ファイルが変更されるたびに不足しているものをすべて pip インストールします。 LLM が反復しているスクリプトを指すと、venv はコードと同様に収束します。
このツールは次の 2 つの障害モードで駆動されました。
人間の広がり。 Venv は ~ で増殖し、キャッシュ ディレクトリは GB を消費し、アクティベーション構文はシェルによって異なり、「動作した環境」のクローン作成は、端末間で pip フリーズをコピーアンドペーストすることを意味します。
エージェントのスプロール。 Python を生成する LLM は、定期的に間違ったインタープリターに pip install し、 VIRTUAL_ENV の設定を忘れ、クラッシュ後に半分インストールされたパッケージを残し、環境の状態を推論するための型付き API を持たず、シェルのみを使用します。
venv-manager は、(1) クリーンな CLI を使用して、(2) モデル コンテキスト プロトコル サーバー、型指定された JSON スナップショット、OS レベルのサンドボックスを備えた一時的な venv、および venv と進化するスクリプトの同期を保つファイル ウォッチャーを使用して問題を解決します。
カール -sSL https://raw.githubusercontent.com/jacopobonomi/venv_manager/main/install.sh |バッシュ
またはソースから:
git clone https://github.com/jacopobonomi/venv_manager && cd venv_manager
インストールする
Goが必要

ビルドには 1.21 以降、実行時には Python 3.x。
venv 操作をネイティブのモデル コンテキスト プロトコル ツールとして公開します。エージェント クライアント (Claude Desktop、Cursor、Zed) は、シェル呼び出しを推測する代わりに、JSON スキーマを使用して型指定されたツールを呼び出します。
Claude Desktop (~/Library/Application Support/Claude/claude_desktop_config.json) に接続します。
{
"mcpサーバー": {
"venv-マネージャー" : {
"コマンド" : " venv-manager " ,
"args" : [ "mcp " ]
}
}
}
公開されるツール (stdio 経由の JSON-RPC 2.0):
実装: ~350 LOC、サードパーティ MCP 配備はゼロ。 stdin/stdout 上の改行区切りの JSON-RPC 2.0。
一時的な実行 ( uvx スタイル、サンドボックス)
# 作成→インストール→実行→破棄をすべて 1 回の呼び出しで実行
venv-manager exec --with リクエスト -- python -c " リクエストのインポート; print(requests.__version__) "
# OS サンドボックスを使用する場合: ネットワークなし、/tmp + 一時的な venv の外部への書き込みなし
venv-manager exec --sandbox --with pandas -- python untrusted.py
--sandbox は、macOS では Sandbox-exec を使用し、Linux では bwrap を使用します。 venv パス、 /tmp 、およびプロセス管理の明示的な許可リストを含むデフォルト拒否プロファイル。ネットワークは共有されていません。
venv-manager watch app.py --venv myenv
親ディレクトリで fsnotify (エディターの atomic-rename 書き込みは存続)、500 ミリ秒のデバウンス、その後:
.py ファイルの AST-lite 正規表現スキャン (ドキュメント文字列、相対インポート、および .venv 、 .git 、 __pycache__ 、node_modules などのベンダー ディレクトリをスキップします)
stdlib モジュール セットに対するフィルター
import-name → pip-package エイリアスを解決します ( cv2 → opencv-python 、 sklearn → scikit-learn 、 PIL → Pillow 、 bs4 → beautifulsoup4 、 yaml → PyYAML 、 ...)
インストールされているパッケージとの差分
venv は常に、現在のファイルの要件のスーパーセットです。これは、上記のデモ GIF が実行するループです。
単一呼び出しコンテキストの入門書としての JSON スナップショット
venv-manager は myenv を説明します
{
"名前" : " myenv " ,
「パス」 :

" /Users/me/.venvs/myenv " 、
"python_version" : " 3.12.6 " ,
"python_path" : " /Users/me/.venvs/myenv/bin/python " ,
"pip_path" : " /Users/me/.venvs/myenv/bin/pip " ,
"パッケージ" : [ "requests==2.34.2" , "rich==15.0.0" , ... ],
"パッケージ数" : 12 、
"サイズバイト" : 45123456 、
"size_human" : " 43.03 MB " 、
"modified_at" : " 2026-07-20T15:41:35Z " ,
"freeze_hash" : " sha256:2c58d830... " ,
「アクティベーション」: {
"bash" : " ソース /Users/me/.venvs/myenv/bin/activate " ,
"zsh" : " ソース /Users/me/.venvs/myenv/bin/activate " ,
"fish" : " ソース /Users/me/.venvs/myenv/bin/activate.fish " ,
"pwsh" : " /Users/me/.venvs/myenv/bin \\ Activate.ps1 " ,
"cmd" : " /Users/me/.venvs/myenv/bin \\ activate.bat "
}
}
1 回のツール呼び出しで、エージェントが環境について推論するために必要なすべてが実行されます。フリーズハッシュを使用すると、エージェントはパッケージ リストを比較する代わりに、O(1) の 2 つの記述呼び出し間のドリフトを検出できます。
コマンド
説明
<名前> を作成 [--python VER]
venv を作成します。設定で use_uv: true の場合、uv を使用します。
リスト [--json]
venv をリストします。
<名前> を削除します
Venv を削除します。
<古い> <新しい> の名前を変更します
python -m venv --upgrade を使用してアクティベーション スクリプトの名前を変更し、再生成します。
クローン <src> <dst>
ソースのピップフリーズでシードされた新鮮なvenv。
パッケージ <名前> [--json]
インストールされたパッケージ。
<名前> <要件> をインストールします
pip install -r 。
アップグレード [名前] [--global]
古いパッケージをアップグレードします (venv ごとまたはすべて)。
クリーン [名前] [--グローバル]
pip キャッシュ + __pycache__ ディレクトリをパージします。
サイズ [名前] [--global] [--json]
ディスクの使用量。
<名前> をアクティブにする
eval $(...) の出力シェルコマンド。
非アクティブ化する
印刷を非アクティブ化します。
<名前> -- <cmd> を実行します。
アクティブ化せずに venv で実行します。 stdioを継承。
exec [--with pkgs] [-r req] [--python V] [--sandbox] [--keep] -- <cmd>
一時的な venv 実行。
<名前> について説明します
完全な JSON スナップショット (上記を参照)。
scan <パス> [--venv N] [--json]
補足

サードパーティの輸入を行為します。 venv に対してチェックします。
watch <パス> --venv N
ファイル変更時に不足しているインポートを自動インストールします。
スナップショット <名前> [-l LABEL]
pip-freeze 状態をキャプチャします。
スナップショット <名前> [--json]
スナップショットをリストします (新しいものから順)。
ロールバック <名前> [スナップショット ID]
すべてアンインストールし、スナップショットから再インストールします。
<名前> をエクスポート
ポータブル マニフェスト (名前 + Python バージョン + フリーズ) を JSON として出力します。
import <マニフェスト.json>
マニフェストから venv を再作成します。
プルーン [--days N] [--dry-run] [--json]
N 日間使用されなかった venvs を削除します。
医師 [--json]
Python のバージョン、UV、壊れた venv を診断します。
`構成ショー
パス
mcp
標準入出力上のモデル コンテキスト プロトコル サーバー。
トゥイ
バブル ティー TUI ブラウザー。
`完了 [bash]
zsh
ほとんどの読み取りコマンドは、安定したマシン解析可能な出力のために --json も受け入れます。
~/.config/venv-manager/config.json ( $XDG_CONFIG_HOME および $VENV_MANAGER_CONFIG を考慮):
{
"base_dir" : " /custom/path/to/venvs " ,
"default_python" : " 3.12 " 、
"use_uv" : true 、
「プルーンアフターデイズ」：90
}
ブートストラップ: venv-manager config init 。
uv が PATH 上にあり、 use_uv: true の場合、 create は uv venv を実行します。通常、コールド キャッシュでは python -m venv よりも 10 ～ 100 倍高速です。
make build # go build -o bin/venv-manager
make test # 単体テスト
デモを作成 # VHS 経由で script/demo/demo.gif を再生成
go test -tags=integration ./internal/manager/... # 統合テスト (実際の pip、実際の PyPI)
CI は、Ubuntu + macOS で go vet 、 go test -race を実行し、Ubuntu と Python 3.12 で統合テストを実行します。
cmd/venv-manager/cobra CLI
内部/マネージャー/コア操作 (作成、インストール、スナップショット、スキャン、監視、実行、記述など)
Internal/config/ XDG 対応の JSON 構成
Internal/mcp/ JSON-RPC 2.0 MCP サーバー (stdio)
内部/tui/ バブル ティー ブラウザ
内部/ユーティリティ/プラットフォーム ヘルパー、サイズの書式設定
ライセンス
Python 仮想環境を簡単に管理するための強力な CLI ツール。
がありました

n ロード中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
7
v0.1.0
最新の
2026 年 7 月 20 日
+ 6 リリース
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A powerful CLI tool for managing Python virtual environments with ease. - jacopobonomi/venv_manager

GitHub - jacopobonomi/venv_manager: A powerful CLI tool for managing Python virtual environments with ease. · GitHub
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
jacopobonomi
/
venv_manager
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
28 Commits 28 Commits .github .github cmd/ venv-manager cmd/ venv-manager internal internal scripts scripts .gitignore .gitignore CHANGELOG.md CHANGELOG.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md go.mod go.mod go.sum go.sum install.sh install.sh terminal_example.gif terminal_example.gif View all files Repository files navigation
A Python virtual-environment runtime for humans and AI agents.
Written in Go. One static binary, no runtime deps beyond python3 (or uv , if available).
The GIF above is real: venv-manager watch app.py --venv X monitors a file, scans its imports with a tiny AST-lite parser, and pip-installs whatever is missing — every time the file changes. Point it at a script an LLM is iterating on and the venv converges as the code does.
Two failure modes drove this tool:
Human sprawl. Venvs multiply across ~ , cache directories eat GB, activation syntax varies by shell, and cloning "the env that worked" means copy-pasting pip freeze between terminals.
Agent sprawl. LLMs generating Python routinely pip install into the wrong interpreter, forget to set VIRTUAL_ENV , leave half-installed packages behind after a crash, and have no typed API to reason about environment state — only shells.
venv-manager solves (1) with a clean CLI and (2) with a Model Context Protocol server , typed JSON snapshots , ephemeral venvs with OS-level sandboxing, and a file watcher that keeps a venv in sync with an evolving script.
curl -sSL https://raw.githubusercontent.com/jacopobonomi/venv_manager/main/install.sh | bash
Or from source:
git clone https://github.com/jacopobonomi/venv_manager && cd venv_manager
make install
Requires Go 1.21+ to build, Python 3.x at runtime.
Exposes venv operations as native Model Context Protocol tools. Agentic clients (Claude Desktop, Cursor, Zed) call typed tools with JSON Schemas instead of guessing shell invocations.
Wire it up in Claude Desktop ( ~/Library/Application Support/Claude/claude_desktop_config.json ):
{
"mcpServers" : {
"venv-manager" : {
"command" : " venv-manager " ,
"args" : [ " mcp " ]
}
}
}
Tools exposed (JSON-RPC 2.0 over stdio):
Implementation: ~350 LOC, zero third-party MCP deps. Newline-delimited JSON-RPC 2.0 on stdin/stdout.
Ephemeral execution ( uvx -style, sandboxed)
# create → install → run → destroy, all in one call
venv-manager exec --with requests -- python -c " import requests; print(requests.__version__) "
# with an OS sandbox: no network, no writes outside /tmp + the ephemeral venv
venv-manager exec --sandbox --with pandas -- python untrusted.py
--sandbox uses sandbox-exec on macOS and bwrap on Linux. Deny-by-default profile with explicit allow-lists for the venv path, /tmp , and process management. Network is unshared.
venv-manager watch app.py --venv myenv
fsnotify on the parent directory (survives editor atomic-rename writes), 500 ms debounce, then:
AST-lite regex scan of .py files (skips docstrings, relative imports, and vendored dirs like .venv , .git , __pycache__ , node_modules )
Filter against a stdlib module set
Resolve import-name → pip-package aliases ( cv2 → opencv-python , sklearn → scikit-learn , PIL → Pillow , bs4 → beautifulsoup4 , yaml → PyYAML , ...)
Diff against installed packages
The venv is always a superset of the current file's requirements. This is the loop the demo GIF above exercises.
JSON snapshot as a single-call context primer
venv-manager describe myenv
{
"name" : " myenv " ,
"path" : " /Users/me/.venvs/myenv " ,
"python_version" : " 3.12.6 " ,
"python_path" : " /Users/me/.venvs/myenv/bin/python " ,
"pip_path" : " /Users/me/.venvs/myenv/bin/pip " ,
"packages" : [ " requests==2.34.2 " , " rich==15.0.0 " , ... ],
"package_count" : 12 ,
"size_bytes" : 45123456 ,
"size_human" : " 43.03 MB " ,
"modified_at" : " 2026-07-20T15:41:35Z " ,
"freeze_hash" : " sha256:2c58d830... " ,
"activation" : {
"bash" : " source /Users/me/.venvs/myenv/bin/activate " ,
"zsh" : " source /Users/me/.venvs/myenv/bin/activate " ,
"fish" : " source /Users/me/.venvs/myenv/bin/activate.fish " ,
"pwsh" : " /Users/me/.venvs/myenv/bin \\ Activate.ps1 " ,
"cmd" : " /Users/me/.venvs/myenv/bin \\ activate.bat "
}
}
One tool call, everything an agent needs to reason about the environment. freeze_hash lets an agent detect drift between two describe calls in O(1) instead of diffing package lists.
Command
Description
create <name> [--python VER]
Create a venv. Uses uv when use_uv: true in config.
list [--json]
List venvs.
remove <name>
Delete a venv.
rename <old> <new>
Rename and re-generate activation scripts via python -m venv --upgrade .
clone <src> <dst>
Fresh venv seeded with pip freeze of source.
packages <name> [--json]
Installed packages.
install <name> <requirements>
pip install -r .
upgrade [name] [--global]
Upgrade outdated packages (per venv or all).
clean [name] [--global]
Purge pip cache + __pycache__ dirs.
size [name] [--global] [--json]
Disk usage.
activate <name>
Print shell command for eval $(...) .
deactivate
Print deactivate .
run <name> -- <cmd>
Execute in a venv without activating; inherited stdio.
exec [--with pkgs] [-r req] [--python V] [--sandbox] [--keep] -- <cmd>
Ephemeral venv run.
describe <name>
Full JSON snapshot (see above).
scan <path> [--venv N] [--json]
Extract third-party imports; check against venv.
watch <path> --venv N
Auto-install missing imports on file change.
snapshot <name> [-l LABEL]
Capture pip-freeze state.
snapshots <name> [--json]
List snapshots (newest first).
rollback <name> [snapshot-id]
Uninstall all, reinstall from snapshot.
export <name>
Print portable manifest (name + python version + freeze) as JSON.
import <manifest.json>
Recreate venv from manifest.
prune [--days N] [--dry-run] [--json]
Remove venvs unused for N days.
doctor [--json]
Diagnose python versions, uv, broken venvs.
`config show
path
mcp
Model Context Protocol server on stdio.
tui
Bubble Tea TUI browser.
`completion [bash
zsh
Most read commands also accept --json for stable, machine-parseable output.
~/.config/venv-manager/config.json (respects $XDG_CONFIG_HOME and $VENV_MANAGER_CONFIG ):
{
"base_dir" : " /custom/path/to/venvs " ,
"default_python" : " 3.12 " ,
"use_uv" : true ,
"prune_after_days" : 90
}
Bootstrap: venv-manager config init .
If uv is on PATH and use_uv: true , create runs uv venv . Typically 10–100× faster than python -m venv on cold cache.
make build # go build -o bin/venv-manager
make test # unit tests
make demo # regenerate scripts/demo/demo.gif via VHS
go test -tags=integration ./internal/manager/... # integration tests (real pip, real PyPI)
CI runs go vet , go test -race on Ubuntu + macOS, and integration tests on Ubuntu with Python 3.12.
cmd/venv-manager/ cobra CLI
internal/manager/ core operations (create, install, snapshot, scan, watch, exec, describe, ...)
internal/config/ XDG-aware JSON config
internal/mcp/ JSON-RPC 2.0 MCP server (stdio)
internal/tui/ Bubble Tea browser
internal/utils/ platform helpers, size formatting
License
A powerful CLI tool for managing Python virtual environments with ease.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
7
v0.1.0
Latest
Jul 20, 2026
+ 6 releases
Uh oh!
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
