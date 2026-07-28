---
source: "https://github.com/microsoft/quicksand"
hn_url: "https://news.ycombinator.com/item?id=49082666"
title: "Microsoft Quicksand: Sandbox your AI agent without Docker or WSL"
article_title: "GitHub - microsoft/quicksand: Quickly sandbox your AI agent. · GitHub"
author: "joeyhage"
captured_at: "2026-07-28T12:48:38Z"
capture_tool: "hn-digest"
hn_id: 49082666
score: 2
comments: 0
posted_at: "2026-07-28T12:16:38Z"
tags:
  - hacker-news
  - translated
---

# Microsoft Quicksand: Sandbox your AI agent without Docker or WSL

- HN: [49082666](https://news.ycombinator.com/item?id=49082666)
- Source: [github.com](https://github.com/microsoft/quicksand)
- Score: 2
- Comments: 0
- Posted: 2026-07-28T12:16:38Z

## Translation

タイトル: Microsoft Quicksand: Docker や WSL を使用せずに AI エージェントをサンドボックス化する
記事のタイトル: GitHub - Microsoft/quicksand: AI エージェントをすばやくサンドボックス化します。 · GitHub
説明: AI エージェントをすばやくサンドボックス化します。 GitHub でアカウントを作成して、microsoft/quicksand の開発に貢献します。

記事本文:
GitHub - Microsoft/quicksand: AI エージェントをすばやくサンドボックス化します。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
マイクロソフト
/
流砂
公共
通知
c にサインインする必要があります

ハンゲの通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
258 コミット 258 コミット .claude/ skill .claude/ skill .github/ workflows .github/ workflows docs docs 例 例 パッケージ パッケージ スクリプト スクリプト テスト テスト .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CHANGELOG.md CHANGELOG.md CONTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md azure-pipelines.yml azure-pipelines.yml package-lock.json package-lock.json package.json package.json pyproject.toml pyproject.toml uv.lock uv.lock uvr_hooks.py uvr_hooks.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Quicksand は、AI エージェントのサンドボックス化に特に重点を置いた、QEMU 仮想マシンの起動、制御、スナップショットを行うための非同期 Python API です。 Quicksand は、Ubuntu および Alpine ディストリビューション用の構築済み Linux VM を提供します。 macOS、Linux、Windows 上の x86_64 および ARM64 上で動作し、root 権限、Docker、システム依存関係はありません。 pip install Quick-sandbox するだけです。
pip install 'クイックサンドボックス[qemu,alpine]'
または、コア パッケージをインストールし、QEMU/イメージを個別に追加します。
pip インストールクイックサンドボックス
クイックサンド Qemu アルパインをインストールする
使用法
非同期をインポートする
流砂からサンドボックスをインポート
非同期デフォルトメイン():
sb としてサンドボックス ( image = "ubuntu" ) と非同期:
結果 = sb を待ちます。実行 ( "echo 'サンドボックスからこんにちは!'" )
print (結果.stdout)
非同期 。実行 (メイン ())
コマンドを実行する
結果 = sb を待ちます。実行 ( "apt update && apt install -y python3" )
print ( result . stdout , result . exit_code )
ホストディレクトリをマウントする
起動時または実行中にホスト ディレクトリを VM に共有します。
# 起動時
サンドボックスと非同期 (
画像 = "ubuntu" ,
mounts = [マウント ( "./workspace" , "/mnt/workspace" )],
) sb として:
。

..
# または実行中のサンドボックス上で動的に
ハンドル = sb を待ちます。マウント ( "/tmp/data" 、 "/mnt/data" )
sbを待ちます。実行 ( "ls /mnt/data" )
sbを待ちます。アンマウント (ハンドル)
ネットワークの構成
サンドボックスはデフォルトでネットワークから分離されています。 NetworkMode.FULL を使用してインターネット アクセスとポート転送をオプトインします。
サンドボックスと非同期 (
画像 = "ubuntu" ,
network_mode = ネットワークモード。フル、
port_forwards = [PortForward (ホスト = 8080、ゲスト = 80)],
) sb として:
...
保存とロード
VM のディスク状態をディレクトリに保存します。後で、別のマシンにロードしても構いません。
sbを待ちます。実行 (「pip install numpy pandas」)
sbを待ちます。 save ( "my-env" ) # VM は実行を継続します
# 後でロードする
sb としてサンドボックス ( image = "my-env" ) と非同期:
sbを待ちます。実行 ( "python3 -c 'import numpy; print(numpy.__version__)'" )
チェックポイントとリバート
VM の完全な状態をキャプチャし、何か問題が発生した場合はロールバックします。
sbを待ちます。チェックポイント (「実験前」)
sbを待ちます。実行 ( "apt install -y something-risky" )
sbを待ちます。 revert ( "before-experiment" ) # VM はチェックポイントに戻ります
デスクトップを制御する
デスクトップ イメージは、ブラウザを備えた完全な Xfce4 グラフィカル環境を提供します。 Quicksand install ubuntu-desktop または Quicksand install alpine-desktop を使用してインストールします。
sb としてサンドボックス (image = "ubuntu-desktop" 、enable_display = True) と非同期:
sbを待ちます。スクリーンショット ( "screen.png" )
sbを待ちます。 type_text ( "ハローワールド" )
sbを待ちます。 press_key ( キー . RET )
sbを待ちます。マウス移動 ( 500 , 300 )
sbを待ちます。マウスクリック ( "左" )
構成
すべてのサンドボックス構成オプションは次のとおりです。
サンドボックス (
# ブートするイメージまたは名前を保存
画像 = "ubuntu" ,
# ゲストRAM (デフォルト: "512M")
メモリ = "2G" 、
仮想 CPU コアの数 (デフォルト: 1)
CPU = 4 、
# 起動時に VM に共有されるホスト ディレクトリ
マウント = [マウント ( "/host" , "/guest" )],
#

NONE、MOUNTS_ONLY (デフォルト)、または完全なインターネット アクセス
network_mode = ネットワークモード。フル、
# ホストの TCP ポートをゲストに転送します
port_forwards = [PortForward (ホスト = 8080、ゲスト = 80)],
# 起動時にゲストファイルシステムを拡張します
ディスクサイズ = "10G" 、
# スクリーンショット/タイプテキスト/マウスコントロール用に仮想 GPU、キーボード、およびマウスを接続します
Enable_display = True 、
# 停止時に VM の状態を自動保存する
save = "私の保存名" ,
)
利用可能な画像
画像
種類
ホイールサイズ
インストールコマンド
それは何ですか
ubuntu
ベース
~341MB
クイックサンドでubuntuをインストールする
Ubuntu 24.04 ヘッドレス
高山
ベース
～78MB
流砂インストールアルパイン
Alpine 3.23 ヘッドレス (高速ブート)
ubuntu-デスクトップ
オーバーレイ (ubuntu)
～263MB
クイックサンドで ubuntu-desktop をインストールする
Ubuntu 24.04 + Xfce4 + Firefox
アルパインデスクトップ
オーバーレイ (アルパイン)
～310MB
クイックサンドでアルパインデスクトップをインストール
アルパイン 3.23 + Xfce4 + クロム
流砂エージェント
オーバーレイ (ubuntu)
～304MB
クイックサンド インストール クイックサンド エージェント
Ubuntu + Python 3.12、uv、build-essential、リクエスト、pyyaml、ddgs、markitdown
流砂クア
オーバーレイ (流砂エージェント)
~445MB
クイックサンド Quicksand-cua をインストールします
エージェント サンドボックス + Xvfb、x11vnc、noVNC、Playwright、Chromium
ソースから構築する
git clone https://github.com/microsoft/quicksand.git
CD 流砂
UV同期
uv run uvr build --all-packages
ドキュメント
トピック
ガイド
ボンネットの下で
インストール
パッケージのインストール
QEMU バイナリ、カーネル、qcow2 ディスク
サンドボックスのライフサイクル
サンドボックスの作成と構成
-m 、 -smp 、 -accel 、マシンタイプ
コマンドの実行
execute() 、ストリーミング、終了コード
カーネルブート、エージェントトークン、hostfwd
ファイル交換
マウント、ホットマウント、データの入出力
CIFS は guestfwd 経由、9p は -fsdev 経由
保存とロールバック
チェックポイント、リバート、永続保存
qcow2 オーバーレイ、savevm、blockdev-snapshot-sync
デスクトップコントロール
スクリーンショット、キーボード、マウス
VNC、GPU、USBタブレット、QM

P入力インジェクション
ネットワークと分離
ネットワークモード、ポートフォワーディング
SLIRP NAT、restrict=on、guestfwd
パフォーマンス
何が高速になるのか
io_uring 、IOThreads、TCG と KVM の比較
貢献する
ガイド
いつ使用するか
画像の作成
新しいベースまたはオーバーレイ イメージ パッケージを構築する
サンドボックスの拡張
メソッド、OS、アーキテクチャ、または QEMU フラグを追加します
テスト
テストの実行または作成
解放する
リリースをカットする
完全なガイド: ユーザー ガイド |ボンネットの下で |寄稿者ガイド
AI エージェントをすばやくサンドボックス化します。
microsoft.github.io/quicksand/ リソース
Readme MIT ライセンスの行動規範
セキュリティ ポリシー アクティビティ カスタム プロパティ スター
3 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Quickly sandbox your AI agent. Contribute to microsoft/quicksand development by creating an account on GitHub.

GitHub - microsoft/quicksand: Quickly sandbox your AI agent. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
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
Uh oh!
There was an error while loading. Please reload this page .
microsoft
/
quicksand
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
258 Commits 258 Commits .claude/ skills .claude/ skills .github/ workflows .github/ workflows docs docs examples examples packages packages scripts scripts tests tests .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md azure-pipelines.yml azure-pipelines.yml package-lock.json package-lock.json package.json package.json pyproject.toml pyproject.toml uv.lock uv.lock uvr_hooks.py uvr_hooks.py View all files Repository files navigation
Quicksand is an async Python API to launch, control, and snapshot QEMU virtual machines with a particular focus on sandboxing AI agents. Quicksand provides pre-built Linux VMs for Ubuntu and Alpine distros. It works on x86_64 and ARM64 across macOS, Linux, and Windows with no root privileges, no Docker, and no system dependencies. Just pip install quick-sandbox .
pip install ' quick-sandbox[qemu,alpine] '
Or install the core package and add QEMU/images separately:
pip install quick-sandbox
quicksand install qemu alpine
Usage
import asyncio
from quicksand import Sandbox
async def main ():
async with Sandbox ( image = "ubuntu" ) as sb :
result = await sb . execute ( "echo 'Hello from the sandbox!'" )
print ( result . stdout )
asyncio . run ( main ())
Run commands
result = await sb . execute ( "apt update && apt install -y python3" )
print ( result . stdout , result . exit_code )
Mount host directories
Share host directories into the VM at boot or on the fly.
# At boot
async with Sandbox (
image = "ubuntu" ,
mounts = [ Mount ( "./workspace" , "/mnt/workspace" )],
) as sb :
...
# Or dynamically on a running sandbox
handle = await sb . mount ( "/tmp/data" , "/mnt/data" )
await sb . execute ( "ls /mnt/data" )
await sb . unmount ( handle )
Configure networking
Sandboxes are network-isolated by default. Opt in to internet access and port forwarding with NetworkMode.FULL .
async with Sandbox (
image = "ubuntu" ,
network_mode = NetworkMode . FULL ,
port_forwards = [ PortForward ( host = 8080 , guest = 80 )],
) as sb :
...
Save and load
Save the VM's disk state to a directory. Load it later, even on a different machine.
await sb . execute ( "pip install numpy pandas" )
await sb . save ( "my-env" ) # VM keeps running
# Load later
async with Sandbox ( image = "my-env" ) as sb :
await sb . execute ( "python3 -c 'import numpy; print(numpy.__version__)'" )
Checkpoint and revert
Capture the full VM state and roll back if something goes wrong.
await sb . checkpoint ( "before-experiment" )
await sb . execute ( "apt install -y something-risky" )
await sb . revert ( "before-experiment" ) # the VM snaps back to the checkpoint
Control a desktop
Desktop images provide a full Xfce4 graphical environment with a browser. Install one with quicksand install ubuntu-desktop or quicksand install alpine-desktop .
async with Sandbox ( image = "ubuntu-desktop" , enable_display = True ) as sb :
await sb . screenshot ( "screen.png" )
await sb . type_text ( "hello world" )
await sb . press_key ( Key . RET )
await sb . mouse_move ( 500 , 300 )
await sb . mouse_click ( "left" )
Configuration
Here are all of the Sandbox configuration options:
Sandbox (
# Image or save name to boot
image = "ubuntu" ,
# Guest RAM (default: "512M")
memory = "2G" ,
# Virtual CPU cores (default: 1)
cpus = 4 ,
# Host directories shared into the VM at boot
mounts = [ Mount ( "/host" , "/guest" )],
# NONE, MOUNTS_ONLY (default), or FULL internet access
network_mode = NetworkMode . FULL ,
# Forward host TCP ports into the guest
port_forwards = [ PortForward ( host = 8080 , guest = 80 )],
# Expand the guest filesystem on boot
disk_size = "10G" ,
# Attach virtual GPU, keyboard, and mouse for screenshot/type_text/mouse control
enable_display = True ,
# Auto-save VM state on stop
save = "my-save-name" ,
)
Available images
Image
Type
Wheel size
Install command
What is it
ubuntu
Base
~341 MB
quicksand install ubuntu
Ubuntu 24.04 headless
alpine
Base
~78 MB
quicksand install alpine
Alpine 3.23 headless (faster boot)
ubuntu-desktop
Overlay ( ubuntu )
~263 MB
quicksand install ubuntu-desktop
Ubuntu 24.04 + Xfce4 + Firefox
alpine-desktop
Overlay ( alpine )
~310 MB
quicksand install alpine-desktop
Alpine 3.23 + Xfce4 + Chromium
quicksand-agent
Overlay ( ubuntu )
~304 MB
quicksand install quicksand-agent
Ubuntu + Python 3.12, uv, build-essential, requests, pyyaml, ddgs, markitdown
quicksand-cua
Overlay ( quicksand-agent )
~445 MB
quicksand install quicksand-cua
Agent Sandbox + Xvfb, x11vnc, noVNC, Playwright, Chromium
Building from source
git clone https://github.com/microsoft/quicksand.git
cd quicksand
uv sync
uv run uvr build --all-packages
Documentation
Topic
Guide
Under the Hood
Installation
Installing packages
QEMU binaries, kernels, qcow2 disks
Sandbox Lifecycle
Creating and configuring sandboxes
-m , -smp , -accel , machine types
Running Commands
execute() , streaming, exit codes
Kernel boot, agent tokens, hostfwd
File Exchange
Mounts, hot-mounts, getting data in/out
CIFS via guestfwd , 9p via -fsdev
Save and Rollback
Checkpoints, reverts, persistent saves
qcow2 overlays, savevm , blockdev-snapshot-sync
Desktop Control
Screenshots, keyboard, mouse
VNC, GPU, USB tablet, QMP input injection
Network and Isolation
Network modes, port forwarding
SLIRP NAT, restrict=on , guestfwd
Performance
What makes it fast
io_uring , IOThreads, TCG vs KVM
Contributing
Guide
When to use
Creating Images
Build a new base or overlay image package
Extending the Sandbox
Add a method, OS, architecture, or QEMU flag
Testing
Run or write tests
Releasing
Cut a release
Full guides: User Guide | Under the Hood | Contributor Guide
Quickly sandbox your AI agent.
microsoft.github.io/quicksand/ Resources
Readme MIT license Code of conduct
Security policy Activity Custom properties Stars
3 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
