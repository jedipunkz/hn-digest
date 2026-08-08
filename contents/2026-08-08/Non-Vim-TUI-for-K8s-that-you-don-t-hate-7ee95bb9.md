---
source: "https://github.com/olafurbergs/krowser"
hn_url: "https://news.ycombinator.com/item?id=49218415"
title: "Non Vim TUI for K8s that you don't hate"
article_title: "GitHub - olafurbergs/krowser · GitHub"
author: "obergs"
captured_at: "2026-08-08T02:50:36Z"
capture_tool: "hn-digest"
hn_id: 49218415
score: 1
comments: 0
posted_at: "2026-08-08T02:30:59Z"
tags:
  - hacker-news
  - translated
---

# Non Vim TUI for K8s that you don't hate

- HN: [49218415](https://news.ycombinator.com/item?id=49218415)
- Source: [github.com](https://github.com/olafurbergs/krowser)
- Score: 1
- Comments: 0
- Posted: 2026-08-08T02:30:59Z

## Translation

タイトル: 嫌いではない K8 用の非 Vim TUI
記事タイトル: GitHub - olafurbergs/krowser · GitHub
説明: GitHub でアカウントを作成して、olafurbergs/krowser の開発に貢献します。

記事本文:
GitHub - オラファーバーグ/クラウザー · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
オラファーバーグ
/
クラウザー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット cmd/ krowser cmd/ krowser docs docs 内部 内部 .gitignore .gitignore ライセンス ライセンス Makefile Makefile README.md README.md go.mod go.mod go.sum go.sum すべてのファイルを表示 リポジトリ ファイル ナビゲーション
██╗ ██╗ ██████╗

██████╗ ██╗ ██╗ ███████╗ ███████╗ ██████╗
██║ ██╔╝ ██╔══██╗ ██╔═══██╗ ██║ ██╗ ██║ ██╔════╝ ██╔════╝ ██╔══██╗
█████╔╝ ██████╔╝ ██║ ██║ ██║╚██╗██║ ███████╗ █████╗ ██████╔╝
██╔═██╗ ██╔══██╗ ██║ ██║ ██║ ╚████║ ╚════██║ ██╔══╝ ██╔══██╗
██║ ██╗ ██║ ██║ ╚██████╔╝ ██║ ╚███║ ███████║ ███████╗ ██║ ██║
╚═╝ ╚═╝ ╚═╝ ╚═╝ ╚═════╝ ╚═╝ ╚══╝ ╚══════╝ ╚══════╝ ╚═╝ ╚═╝
Bubble Tea v2 で構築された洗練された Kubernetes TUI ブラウザー、
バブル v2 およびリップグロス v2 。
クラスター、名前空間、リソースの参照、構文が強調表示されたログのストリーミング、スケーリングと削除
ワークロードの管理、ポート転送の管理など、すべて端末から行えます。
ナビゲーション — k9s からインスピレーションを得たキーバインディングを備えたコンテキスト、名前空間、およびリソース (種類) ピッカー
リソース テーブル — 型付き列、ファジー フィルタリング、および自動リフレッシュ
詳細ビュー — YAML および選択可能なシークレット データを含む記述出力 (インプレースでデコード)
ポッドログ — ライブフォロー、コンテナ切り替え、タイムスタンプ切り替え、
ログレベル、タイムスタンプ、数値、URL、IP、UUID などの構文ハイライト
変更アクション - 削除、スケール、ロールアウトの再起動、編集はすべて確認ダイアログの背後で行われます。
ポートフォワード — クラスターからポートフォワードを開いて管理します
CPU/メモリゲージ — gr を使用したライブトップビュー

メトリクスサーバーが利用可能な場合の adient プログレスバー
テーマ — 10 個のカラー テーマ。 T または --theme で選択可能。選択内容は実行間で保持されます。
追加機能 — トースト通知、フォーカスブラー、ヘルプオーバーレイ
krowser で閲覧されている種類クラスターを示す、約 45 秒のウォークスルー記録:
リソース テーブル、構文が強調表示されたライブ ログ、YAML 詳細、テーマの切り替え、
CPU/メモリゲージ。
asciinema play docs/krowser-demo.cast
キャストは埋め込みプレーヤー用にアップロードする準備もできています。
asciinema アップロード docs/krowser-demo.cast
要件
Go 1.25 以降 (go 1.25 ディレクティブにより、1.21 以降でツールチェーンの自動ダウンロードがトリガーされます)
クラスターにアクセスできる kubeconfig (KUBECONFIG または ~/.kube/config )
フルパレットのトゥルーカラーサポートを備えた端末
github.com/olafurb/krowser@latest をインストールしてください
構築して実行する
ビルドする
./bin/krowser --kubeconfig ~ /.kube/config
Makefile は、 make test (-race 付き)、make lint、および make fmt も提供します。
旗
説明
--kubeconfig パス
kubeconfig へのパス (デフォルト: $KUBECONFIG または ~/.kube/config )
--コンテキスト名
使用するkubeconfigコンテキスト
--名前空間名
参照する名前空間 (デフォルト: コンテキスト名前空間、または何も設定されていない場合はすべての名前空間)
--all-namespaces
すべての名前空間を参照する
--テーマ名
テーマ名 (デフォルト: Monokai )
使用法
キー
アクション
↑ / ↓
行を移動する
入力してください
詳細を開く (最上位のポッド/ノードも)
y/d
YAML / 説明
私
ストリームログ
/
ファジー フィルター (フィルターに入力、ESC でクリア)
なし / グラム
名前空間ピッカー / すべての名前空間
あなた
コンテキストを切り替える
k
種類を選ぶ
r
リフレッシュする
x / s / R / e
削除/スケール/再起動/編集
f
ポートフォワード
t
トップ (メトリクス)
T
テーマ
?
ヘルプオーバーレイ
q / esc
戻る
Ctrl+C
やめる
ログ画面
キー
アクション
f
フォローモードの切り替え
t
タイムスタンプの切り替え
c
コンテナを選ぶ
q / esc / ←
戻る
テーマ設定
T を押してテーマ ピッカーを開くか、--theme "D を渡します

racula」起動時テーマ：モノカイ、
ワンダーク、ドラキュラ、ノルド、グラブボックスダーク、ソラライズドダーク、キャットプッチンモカ、トーキョーナイト、
ソラリゼーションライトとキャットプッチンラテ。
選択したテーマは ~/.local/share/krowser/theme に保存されます
(設定されている場合は $XDG_DATA_HOME/krowser/theme)、次回の起動時に復元されます。 --theme フラグ
保存された選択内容よりも優先されます。
cmd/krowser CLI エントリ ポイント (フラグ、kubeconfig ワイヤリング)
内部/k8s Kubernetes クライアント ヘルパー (リスト、ログ、トップ、フォワード、アクション)
内部/tui バブル ティー アプリケーション
app.go ルート モデル、ナビゲーション、テーマ、クロム
resource.go の型指定されたリソース テーブルとフィルタリング
詳細.go YAML / 説明 / シークレットビュー
logs.go ストリーミングされた、構文が強調表示されたポッド ログ
top.go ライブ CPU/メモリ ゲージ
picker.go コンテキスト、名前空間、種類、コンテナー、テーマ
theme(s).go テーマのレジストリとパレット
loghl/ ログ構文強調表示エンジン (謝辞を参照)
テスト
テストを行う
このスイートには、リソース テーブル、ピッカー、テーマ、メトリックの書式設定、および
ログ強調表示エンジン (loglit 独自のパターン ゲートおよびオーバーラップ テストを含む)。
ログ強調表示エンジン — からベンダー提供
madmaxieee/loglit (MIT、© 2025 Tsng、Kahiok)。
loglit はエンジンを Go の内部/パッケージに保持しており、他のパッケージからインポートすることはできません。
モジュールなので、internal/tui/loghl/ にそのままコピーされます (パターン、構文ゲーティング、
キーワード マッチング、ANSI レンダリング)、krowser のテーマから色を変更しました。そのライセンス
Internal/tui/loghl/LICENSE にあります。
UI フレームワーク — バブル ティー v2 、
バブル v2 、リップグロス v2 。
ファジー フィルタリング — sahilm/fuzzy 。
テーマのインスピレーション — loglit のパレットは東京の夜と
log-highlight.nvim テーマ。
マサチューセッツ工科大学ベンダーの Loglit エンジンも MIT ライセンスを取得しています。参照してください
内部/tui/loghl/LICENSE 。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ R

解放する
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to olafurbergs/krowser development by creating an account on GitHub.

GitHub - olafurbergs/krowser · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
olafurbergs
/
krowser
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits cmd/ krowser cmd/ krowser docs docs internal internal .gitignore .gitignore LICENSE LICENSE Makefile Makefile README.md README.md go.mod go.mod go.sum go.sum View all files Repository files navigation
██╗ ██╗ ██████╗ ██████╗ ██╗ ██╗ ███████╗ ███████╗ ██████╗
██║ ██╔╝ ██╔══██╗ ██╔═══██╗ ██║ ██╗ ██║ ██╔════╝ ██╔════╝ ██╔══██╗
█████╔╝ ██████╔╝ ██║ ██║ ██║╚██╗██║ ███████╗ █████╗ ██████╔╝
██╔═██╗ ██╔══██╗ ██║ ██║ ██║ ╚████║ ╚════██║ ██╔══╝ ██╔══██╗
██║ ██╗ ██║ ██║ ╚██████╔╝ ██║ ╚███║ ███████║ ███████╗ ██║ ██║
╚═╝ ╚═╝ ╚═╝ ╚═╝ ╚═════╝ ╚═╝ ╚══╝ ╚══════╝ ╚══════╝ ╚═╝ ╚═╝
A slick Kubernetes TUI browser built with Bubble Tea v2 ,
Bubbles v2 , and Lip Gloss v2 .
Browse clusters, namespaces, and resources, stream syntax-highlighted logs, scale and delete
workloads, and manage port-forwards — all from your terminal.
Navigation — context, namespace, and resource (kind) pickers with k9s-inspired keybindings
Resource tables — typed columns, / fuzzy filtering, and auto-refresh
Detail views — YAML and describe output with selectable Secret data (decoded in place)
Pod logs — live follow, container switching, timestamp toggle, and
syntax highlighting for log levels, timestamps, numbers, URLs, IPs, UUIDs, and more
Mutating actions — delete, scale, rollout restart, and edit, all behind a confirm dialog
Port-forwards — open and manage port-forwards from the cluster
CPU/memory gauges — a live top view with gradient progress bars when metrics-server is available
Themes — 10 color themes, selectable with T or --theme , with the selection persisted across runs
Extras — toast notifications, focus blur, help overlay
A ~45s walkthrough recording showing a kind cluster being browsed with krowser:
resource tables, syntax-highlighted live logs, YAML detail, theme switching, and the
CPU/memory gauges.
asciinema play docs/krowser-demo.cast
The cast is also ready to upload for an embedded player:
asciinema upload docs/krowser-demo.cast
Requirements
Go 1.25+ (the go 1.25 directive triggers an automatic toolchain download on 1.21+)
A kubeconfig with access to a cluster (KUBECONFIG or ~/.kube/config )
A terminal with truecolor support for the full palette
go install github.com/olafurb/krowser@latest
Build & run
make build
./bin/krowser --kubeconfig ~ /.kube/config
The Makefile also provides make test (with -race ), make lint , and make fmt .
Flag
Description
--kubeconfig PATH
path to the kubeconfig (default: $KUBECONFIG or ~/.kube/config )
--context NAME
kubeconfig context to use
--namespace NAME
namespace to browse (default: the context namespace, or all namespaces if none is set)
--all-namespaces
browse all namespaces
--theme NAME
theme name (default: Monokai )
Usage
Key
Action
↑ / ↓
navigate rows
Enter
open detail (a pod/node in top too)
y / d
YAML / describe
l
stream logs
/
fuzzy filter (type to filter, esc clears)
n / g
namespace picker / all namespaces
u
switch context
k
pick a kind
r
refresh
x / s / R / e
delete / scale / restart / edit
f
port-forward
t
top (metrics)
T
themes
?
help overlay
q / esc
back
Ctrl-C
quit
Logs screen
Key
Action
f
toggle follow mode
t
toggle timestamps
c
pick a container
q / esc / ←
back
Theming
Press T to open the theme picker or pass --theme "Dracula" on startup. Themes: Monokai,
One Dark, Dracula, Nord, Gruvbox Dark, Solarized Dark, Catppuccin Mocha, Tokyo Night,
Solarized Light, and Catppuccin Latte.
The theme you select is persisted to ~/.local/share/krowser/theme
( $XDG_DATA_HOME/krowser/theme if set) and restored on the next launch. A --theme flag
takes precedence over the saved selection.
cmd/krowser CLI entry point (flags, kubeconfig wiring)
internal/k8s Kubernetes client helpers (list, logs, top, forwards, actions)
internal/tui Bubble Tea application
app.go root model, navigation, theming, chrome
resource.go typed resource tables with filtering
detail.go YAML / describe / secret views
logs.go streamed, syntax-highlighted pod logs
top.go live CPU/memory gauges
picker.go contexts, namespaces, kinds, containers, themes
theme(s).go theme registry and palette
loghl/ log syntax-highlighting engine (see Acknowledgments)
Testing
make test
The suite covers the resource tables, pickers, themes, metrics formatting, and the
log-highlighting engine (including loglit's own pattern-gate and overlap tests).
Log highlighting engine — vendored from
madmaxieee/loglit (MIT, © 2025 Tsng, Kahiok).
loglit keeps its engine in Go internal/ packages that cannot be imported from other
modules, so it is copied verbatim into internal/tui/loghl/ (patterns, syntax gating,
keyword matching, and ANSI rendering) and recolored from krowser's themes. Its license
lives in internal/tui/loghl/LICENSE .
UI framework — Bubble Tea v2 ,
Bubbles v2 , Lip Gloss v2 .
Fuzzy filtering — sahilm/fuzzy .
Theme inspiration — loglit's palettes draw on the Tokyo Night and
log-highlight.nvim themes.
MIT . The vendored loglit engine is also MIT licensed; see
internal/tui/loghl/LICENSE .
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
