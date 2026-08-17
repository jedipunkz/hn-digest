---
source: "https://github.com/rysh-ai/rysh-cli-code"
hn_url: "https://news.ycombinator.com/item?id=49338780"
title: "Rysh – an AI harness where Claude and Codex agents work as a team (Go)"
article_title: "GitHub - rysh-ai/rysh-cli-code: Rysh CLI — agentic terminal multiplexer (open source) · GitHub"
image: "https://opengraph.githubassets.com/bf7606fb16ef84e76606010dc1e38cc1dbe4e16a4c0821da24e240f5d118a9db/rysh-ai/rysh-cli-code"
author: "halilagin"
captured_at: "2026-08-17T23:13:56Z"
capture_tool: "hn-digest"
hn_id: 49338780
score: 1
comments: 0
posted_at: "2026-08-17T22:51:16Z"
tags:
  - hacker-news
  - translated
---

# Rysh – an AI harness where Claude and Codex agents work as a team (Go)

- HN: [49338780](https://news.ycombinator.com/item?id=49338780)
- Source: [github.com](https://github.com/rysh-ai/rysh-cli-code)
- Score: 1
- Comments: 0
- Posted: 2026-08-17T22:51:16Z

## Translation

タイトル: Rysh – クロードとコーデックスのエージェントがチームとして機能する AI ハーネス (Go)
記事タイトル: GitHub - rysh-ai/rysh-cli-code: Rysh CLI — エージェント型端末マルチプレクサ (オープンソース) · GitHub
説明: Rysh CLI — エージェントターミナルマルチプレクサー (オープンソース) - rysh-ai/rysh-cli-code

記事本文:
GitHub - rysh-ai/rysh-cli-code: Rysh CLI — エージェント型ターミナル マルチプレクサ (オープン ソース) · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
リシュアイ
/
rysh-cli-コード
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
29 コミット 29 コミット フォルダーとファイル
.github/ workflows .github/ workflows action action cmd cmd docs docs evals evals 例 例 内部

内部パッケージ パッケージ パッケージング パッケージング スクリプト スクリプト .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yml .goreleaser.yml COTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md flake.nix flake.nix go.mod go.mod go.sum go.sum rysh.config.yaml.example rysh.config.yaml.example system-prompt.md system-prompt.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Rysh CLI: エージェント端末
Go で書かれたマルチプレクサ。タブ、ペイン、分割、vim / htop が正確に動作する
ご想像のとおり、すべてのペインがプロンプトに応答できるエージェントでもあり、
呼び出しツール。
このリポジトリには CLI 自体が保持されます。それはによって異なります
rysh-cli-shared 。
構築または貢献?から開始
rysh-cli-parent — それは
Makefile、Go ワークスペース、CI を使用して、このモジュールをローカル チェックアウトに接続します。
共有のもの。
rysh が実際に何をするのかについて、ナレーションと字幕付きの 5 つのウォークスルーを順番に説明します。
それぞれは前の を前提としています。以下のすべてのプレーヤーがこのページでライブ中です。押す
プレイしても、GitHub からは何も離れません。最初の 4 つはビルドに対して記録されます。
install 行に表示されるので、その中のすべてのコマンドを入力できます。
go install 、最初のセッション、タブ、2 番目のレーン、3 つのペインのスタック — 次に、そのスタックの各ペインに ##claude が含まれるため、1 つのウィンドウに 3 つの独立した Claude セッションが保持されます。
セッションが端末ではなくデーモンである理由。クロードとコーデックスはそれぞれコードワードを告げられ、セッションは完全に停止され、再起動後にプロンプ​​トなしで再開され、その前に告げられたコードワードで応答します。
複数のエージェントと複数のベンダーを並べて表示: Ctrl+S と数字でスタック内を移動し、それぞれ 1 つのタスクを実行し、3 つのファイルを並行して書き込みます。
##ボードオープン、rysh ansaプロンプト、a

nd エージェントが相互に運転する: ロードマップ (クロード) が目標を設定し、フリート マネージャー (コーデックス) がそれを作業指示に分割し、2 人のワーカーが並行して構築し、マネージャーはサインオフする前に修正をワーカーに送り返します。彼らはブラウザー Todo マネージャーを同梱しています。バックエンドもビルドステップもありません。
詳細: エージェント組織図がグラフである理由、エッジが何を意味するのか、エッジを削除すると何が変わるのか。注文はメッセージとして伝わり、結果はボードに表示され、エッジのない 2 人の従業員が独立した評決を下します。
tmux または Zellij を何年も使用している場合でも、1 から始めてください。ペインは次のようになります。
見慣れたものですが、その下のセッション モデル (ビデオ 2) はそうではありません。
SecretNAT はデフォルトでオンになっています。シークレットはリクエスト本文内のトークンに置き換えられます
マシンから離れる前に、ライブ資格情報を平文で送信する応答
ペインに報告されます。応答は仕様上書き換えられません。
マッピングはローカルで元に戻すことができるため、##snat get <token> は実際のトークンを返します。
値を独自のペインに表示し、モデルはそのトークンのみを認識します。
2 つのバイナリが存在しますが、それらは互換性がありません。
1 台のマシンで両方をサポートできるように名前が異なります。すべてのインストール コマンド
この README は rysh をインストールし、そのために書かれたスクリプト、エイリアス、ドキュメントはインストールします。
他のものと対立するものではありません。それが名前が異なるすべての理由です。チェックする
rysh バージョンまたは ry バージョンにあるもの。
このリポジトリは rysh の半分である Apache-2.0 であり、ライセンスと
あなたが読んでいるソースの横にある注意してください。
ソースからのオープンソース ビルド:
github.com/rysh-ai/rysh-cli-code/cmd/rysh@latest をインストールしてください
バイナリは $(go env GOPATH)/bin に rysh として配置されます。 Go 1.25.3 または
newer — このモジュールの go.mod で宣言されたフロア。
事前に構築されたオープンソース ビルド:
カール -fsSL https://packages.rysh.ai/install-rysh

.sh |しー
このリポジトリ独自のリリース (Apache-2.0 バイナリではなく) から rysh をインストールします。
り。実行時に最新のリリースを解決します。 RYSH_VERSION を 1 に固定するように設定します
代わりに、または RYSH_INSTALL_DIR を使用して、どこに配置されるかを選択します。
事前に構築されたディストリビューション — rysh ではなく ry をインストールします。パッケージ化されたものです
製品と個別のディストリビューション。このリポジトリによって生成されたものは何もありません。
したがって、コマンドを説明する ry --help は、ここでのソースに関する証拠ではありません。使用する
読んでいるソースのビルドが必要な場合は、上記のインストールに進みます。
そのインストール コマンドはここでは意図的に再現されていません: この README ドキュメント
オープンソースビルド。事前構築されたディストリビューションには独自のチャネルと独自の
ドキュメント。
WSL2 がサポートされているパスです。ネイティブ Windows はコンパイルして実行できますが、実行できません
ペインを開く — これらは 2 つの異なる主張であり、両方とも真実です。
GOOS=windows GOARCH=amd64 go build ./... 成功 — windows/amd64 ターゲット
このソースからクリーンにビルドします。このリポジトリの .goreleaser.yml の名前は、
意図的に cli のみをアーカイブするので、アーティファクトがそれが何であるかを示します。
セッションを開始できません。 Rysh の PTY 層には ConPTY 実装がないため、
platform.PTYSupported は、Windows およびすべてのセッション開始コマンドでは false です。
セッションを開始してからではなく、WSL ガイダンスによって前もって拒否されます。
最初のペイン ( cmd/rysh/pty_preflight.go ) で失敗します。
実行できることは、セッションとの対話を含む、ペインのないコマンド セットです。
Windows 側から WSL で実行: rysh send 、 rysh exec 、 rysh prompt 、
rysh list-sessions 、rysh install 、rysh eval 、rysh Doctor 。
実際のセッションでは、WSL2 をインストールし、その中の Linux 命令を使用します。
Linux ビルドは変更せずに実行されます。
wsl -- インストール
# WSL内
github.com/rysh-ai/rysh-cli-code/cmd/rysh@latest # rysh をインストールしてください
ネイティブ ペインには ConPTY が必要です

同じ縫い目の後ろ。まだ実装されていません。
このモジュールは単独でスタンドアロンで構築されます。置換ディレクティブは含まれていないため、
共有モジュールは公開バージョンから解決されます。
git clone https://github.com/rysh-ai/rysh-cli-code
cd rysh-cli-コード
ビルドに行く ./cmd/rysh # -> ./rysh
スーパーリポジトリからは、両方のモジュールを同時に開発する方法が示されています。
go.work は共有モジュールを兄弟チェックアウトに接続するため、そこでの編集が必要になります。
リリースのラウンドトリップなしでピックアップされます:
git clone --recursive https://github.com/rysh-ai/rysh-cli-parent
cd rysh-cli-親
ビルド番号を作る -> bin/rysh
make install # -> ~/.local/bin/rysh (PREFIX= でオーバーライド)
make test # 両方のモジュール
いずれにしても 1.25.3 以降を使用してください。
エクスポート ANTHROPIC_API_KEY=sk-ant-...
rysh オンボード --provider anthropic --key-env ANTHROPIC_API_KEY
リーシュドクター
onboard はキーを検証し、プロジェクトローカルの rysh.config.yaml を書き込み、開きます
あなたのセッション。 rysh --help は、完全なコマンド サーフェイスをリストします。
各ペインには入力モードがあり、Esc Esc で入力モードを切り替えます: シェル (実際の PTY) →
プロンプト (LLM に移動) → rysh (マルチプレクサ コマンド) → チャット
（会話、ツールなし）。これら 4 つはデフォルトで有効になっています。 ##モードリストの表示
ペインが持つものと ##mode new <mode> が残りを追加します。
親リポジトリの README には、
キーバインド、セッション、エージェント、ヒューマノイド。
## で始まるコマンドは、rysh モードのペインに入力されます。外から
セッション rysh exec -- '##<cmd>' は同じことを実行し、その出力を出力します。
rysh create work # 作成して添付します。切り離したままにする場合は -d を追加します
rysh list-sessions # 何が実行されているか
rysh アタッチ作業 # 後で戻ってください
rysh 切り離し作業 # 実行したままにする
rysh 作業を停止します # デーモンをシャットダウンします
セッションはデーモンとそのペインです。端末の寿命を超えてしまうため、
ウィンドウは強制終了するのではなく、切り離されます。
2. タブ、レーン、

ペインとスタック
タブにはレーン (列) が保持されます。レーンにはスタックが保持されます。スタックはペインを保持します。
Zellij スタイル — 1 つは展開され、残りは [N/M] と表示されるタイトル バーに折りたたまれます。
##新しいタブ 新しいタブ
##new LANE アクティブなタブの新しいレーン (列)
##new pane アクティブなレーンの下部にあるペイン
##新しいグリッド 4 アクティブ レーンにスタックされた 4 つのペイン
##新しいグリッド 3x4 3 レーン x 4 ペイン、アクティブなタブ内
##新しいグリッド 2x3x4 2 タブ x 3 レーン x 4 ペイン
##new stack 4 アクティブなスタック内にさらに 4 つのペインが積み上げられています
構築したものを見て、移動させます。
##タブリスト ##レーンリスト ##ペインリスト
##panegroup レイアウト レーン レイアウト、全体
##ペインをスタック内で上|下|左|右に移動するか、次のレーンに移動します
##move pane <p> to-lane <lane> ライブ ペインを別の場所に配置します
##タブ名ビルド ##レーン名左 ##ペイン名ビルダー
すべてのペインにエージェントが含まれるスタックは、次のビデオ 3 です。
このページの上部にあるチュートリアル。
3. エージェントボード、フリート、およびフリート自体のボード
Panes はインタラクティブな Claude または Codex を実行でき、rysh はどちらを記憶しているので、
セッションを停止/開始すると、同じ会話が再開されます。
##claude このペインでクロードを実行し、自動的に再開します
##codex このペインで codex を実行します。同じ内容です
##pane new --claude "パーサーで開始" 新しいペイン、すでに動作しています
エージェント ボードは、これらのペインが相互に通信したり、あなたと通信したりする場所です。
##board open ボードペインを開きます
##board post <text> マイルストーンを投稿する
##ボード返信 <スレッド> <テキスト>
ペイン内のエージェントはフォーカスを盗むことなく投稿し、以下を読み返します。
rysh ボードのポスト --as " $RYSH_PANE " -- ' パーサーの配線が完了しました '
rysh ボードテール -- 制限 20
##board Agent up はボード エージェントを開始します。クロードを実行する非表示のペインです。
パス内に存在し、エージェント間でメッセージをルーティングし、認識しているリクエストを拒否します
は古いです。 ##board エージェントが表示されると、フォーカスを移動せずに描画されます。

##board エージェントが実行されている間、それを非表示にします。
フリートは、独自のボードを持つ名前付きペインのグループです。 1つを登録する
ペインは開きません - それは安価な半分です:
##フリート登録 epic-07 --board epic-07
##艦隊状態エピック-07アップ
##フリートリスト ##フリートショー epic-07
##艦隊はエピック-07を忘れてください。ペインは実行され続けます
rysh board tail --fleet epic-07 # セッションのストリームではなく、フリートのストリーム
フリートはクロード、コーデックス、またはその両方の場合があります。コマンドはどちらでも同じです。
パス
何
cmd/rysh
メインパッケージ — エントリポイントとコマンドサーフェス
内部/ツイ
ターミナルUI
内部/アクター
ワークスペース/タブ/ペイン/エージェント アクター、NATS 上の proto.actor
内部/vterm
端末エミュレーション (スクロールバック付きの vt10x フォークを含む)
内部/プロバイダー
LLM プロバイダー アダプター
内部/プラットフォーム
rysh ができることを変更するホスト機能 (例: PTY サポート)
アクション/
setup-rysh コンポジット GitHub アクション
貢献とセキュリティ
このリポジトリは、他の場所で開発されたツリーの一方向のエクスポートであるため、コミット
ここに直接プッシュされたものは、保持されるのではなく、次のエクスポートによって上書きされます。それはそうです
パッチを無意味にしないでください — COTRIBUTING.md では、パッチがどこにあるのかを説明しています。
実際に着陸します。
脆弱性が見つかりましたか?公開の問題、PR、またはディスカッションを開かないでください。リッシュが走る
シェルコム

[切り捨てられた]

## Original Extract

Rysh CLI — agentic terminal multiplexer (open source) - rysh-ai/rysh-cli-code

GitHub - rysh-ai/rysh-cli-code: Rysh CLI — agentic terminal multiplexer (open source) · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
rysh-ai
/
rysh-cli-code
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
29 Commits 29 Commits Folders and files
.github/ workflows .github/ workflows action action cmd cmd docs docs evals evals examples examples internal internal packages packages packaging packaging scripts scripts .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yml .goreleaser.yml CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md flake.nix flake.nix go.mod go.mod go.sum go.sum rysh.config.yaml.example rysh.config.yaml.example system-prompt.md system-prompt.md View all files Repository files navigation
The Rysh CLI: an agentic terminal
multiplexer written in Go. Tabs, panes, splits, and vim / htop working exactly
as you expect — except every pane is also an agent that can answer prompts and
call tools.
This repository holds the CLI itself. It depends on
rysh-cli-shared .
Building or contributing? Start at
rysh-cli-parent — it carries the
Makefile, the Go workspace, and CI, and wires this module to a local checkout of
the shared one.
Five narrated, subtitled walkthroughs of what rysh actually does, in order —
each one assumes the previous . Every player below is live on this page; press
play, nothing leaves GitHub. The first four are recorded against the build the
install line gives you, so every command in them is one you can type.
go install , a first session, a tab, a second lane, a stack of three panes — then ##claude in each pane of that stack, so one window holds three independent Claude sessions.
Why a session is a daemon and not a terminal. Claude and Codex are each told a codeword, the session is stopped outright , and after the restart both resume unprompted and answer with the codeword they were told before it.
More than one agent, and more than one vendor, side by side: Ctrl+S and a digit to move around a stack, a task each, three files written in parallel.
##board open , rysh ansa prompt , and agents driving each other : roadmap (Claude) sets the goal, fleet-manager (Codex) splits it into work orders, two workers build in parallel, and the manager sends a correction back to a worker before signing off. They ship a browser todo manager — no backend, no build step.
The deep end: why an agent org chart is a graph, what the edges mean, and what changes when you remove one. Orders travel down as messages, results come up on the board, and two workers with no edge between them give independent verdicts.
Start at 1 even if you have used tmux or Zellij for years — the panes look
familiar and the session model underneath them does not, which is video 2.
SecretNAT is on by default. Secrets are substituted with tokens in the request body
before it leaves the machine, and a response carrying a live credential in plaintext
is reported into the pane. Responses are not rewritten — by design.
The mapping is reversible locally, so ##snat get <token> hands you back the real
value in your own pane, and the model only ever saw the token.
Two binaries exist and they are not interchangeable :
The names differ so that one machine can carry both. Every install command in
this README installs rysh , and a script, alias or doc written for one does
not run against the other — which is the whole reason the names differ. Check
what you have with rysh version or ry version .
This repository is the rysh half: Apache-2.0, with LICENSE and
NOTICE beside the source you are reading.
The open-source build, from source:
go install github.com/rysh-ai/rysh-cli-code/cmd/rysh@latest
The binary lands in $(go env GOPATH)/bin as rysh . Requires Go 1.25.3 or
newer — the floor declared in this module's go.mod .
The open-source build, prebuilt:
curl -fsSL https://packages.rysh.ai/install-rysh.sh | sh
Installs rysh from this repository's own releases — the Apache-2.0 binary, not
ry . It resolves the newest release at run time; set RYSH_VERSION to pin one
instead, or RYSH_INSTALL_DIR to choose where it lands.
The prebuilt distribution — installs ry , not rysh . It is the packaged
product and a separate distribution; nothing in it is produced by this repository,
so ry --help describing a command is not evidence about the source here. Use
go install above if you want the build whose source you are reading.
Its install commands are deliberately not reproduced here: this README documents
the open-source build. The prebuilt distribution has its own channels and its own
documentation.
WSL2 is the supported path. Native Windows compiles and runs, but it cannot
open a pane — those are two different claims and both are true:
GOOS=windows GOARCH=amd64 go build ./... succeeds — the windows/amd64 target
builds clean from this source. This repository's .goreleaser.yml names that
archive cli-only on purpose, so the artifact says what it is.
It cannot start a session. Rysh's PTY layer has no ConPTY implementation, so
platform.PTYSupported is false on Windows and every session-opening command
is refused up front with WSL guidance — rather than starting a session and then
failing on your first pane ( cmd/rysh/pty_preflight.go ).
What it can do is the pane-less command set, including talking to a session
running in WSL from the Windows side: rysh send , rysh exec , rysh prompt ,
rysh list-sessions , rysh install , rysh eval , rysh doctor .
For a real session, install WSL2 and use the Linux instructions inside it — the
Linux build runs unmodified:
wsl -- install
# inside WSL
go install github.com/rysh-ai/rysh-cli-code/cmd/rysh@latest # rysh
Native panes need ConPTY behind the same seam; it is not implemented yet.
This module alone builds standalone — it carries no replace directive, so the
shared module resolves from its published version:
git clone https://github.com/rysh-ai/rysh-cli-code
cd rysh-cli-code
go build ./cmd/rysh # -> ./rysh
From the superrepo, which is how to develop against both modules at once — its
go.work wires the shared module to the sibling checkout, so an edit there is
picked up without a release round-trip:
git clone --recursive https://github.com/rysh-ai/rysh-cli-parent
cd rysh-cli-parent
make build # -> bin/rysh
make install # -> ~/.local/bin/rysh (override with PREFIX=)
make test # both modules
Go 1.25.3 or newer either way.
export ANTHROPIC_API_KEY=sk-ant-...
rysh onboard --provider anthropic --key-env ANTHROPIC_API_KEY
rysh doctor
onboard validates the key, writes a project-local rysh.config.yaml , and opens
your session. rysh --help lists the full command surface.
Each pane has an input mode and Esc Esc cycles it: shell (a real PTY) →
prompt (goes to the LLM) → rysh (multiplexer commands) → chat
(conversation, no tools). Those four are enabled by default; ##mode list shows
what a pane has and ##mode new <mode> adds the rest.
The parent repo README has
the keybindings, sessions, agents and humanoids.
Commands starting with ## are typed into a pane in rysh mode. From outside
the session, rysh exec -- '##<cmd>' runs the same thing and prints its output.
rysh create work # create and attach; add -d to leave it detached
rysh list-sessions # what is running
rysh attach work # come back to it later
rysh detach work # leave it running
rysh stop work # shut the daemon down
A session is a daemon plus its panes. It outlives your terminal, so closing the
window detaches rather than kills.
2. Tabs, lanes, panes and stacks
A tab holds lanes (columns); a lane holds stacks ; a stack holds panes,
Zellij-style — one expanded, the rest collapsed to a title bar showing [N/M] .
##new tab a new tab
##new lane a new lane (column) in the active tab
##new pane a pane at the bottom of the active lane
##new grid 4 4 panes stacked in the active lane
##new grid 3x4 3 lanes x 4 panes, in the active tab
##new grid 2x3x4 2 tabs x 3 lanes x 4 panes
##new stack 4 4 more stacked panes in the active stack
Look at what you built, and move things around:
##tab list ##lane list ##pane list
##panegroup layout the lane layout, whole
##move pane up|down|left|right reorder in the stack, or cross to the next lane
##move pane <p> to-lane <lane> put a live pane somewhere else
##tab name build ##lane name left ##pane name builder
A stack with an agent in every pane is video 3 of the
tutorials at the top of this page.
3. The agents board, fleets, and a fleet's own board
Panes can run an interactive Claude or Codex , and rysh remembers which, so
a session stop/start resumes the same conversation:
##claude run claude in this pane, resumed automatically
##codex run codex in this pane, same deal
##pane new --claude "start on the parser" a new pane, already working
The agents board is where those panes talk to each other and to you:
##board open open the board pane
##board post <text> post a milestone
##board reply <thread> <text>
An agent inside a pane posts without stealing focus, and reads back:
rysh board post --as " $RYSH_PANE " -- ' parser wiring done '
rysh board tail --limit 20
##board agent up starts a board agent — a hidden pane running Claude that
sits in the path, routes messages between agents, and refuses a request it knows
is stale. ##board agent visible draws it without moving your focus;
##board agent invisible puts it away while it keeps running.
A fleet is a named group of panes with a board of its own. Registering one
opens no panes — it is the cheap half:
##fleet register epic-07 --board epic-07
##fleet state epic-07 up
##fleet list ##fleet show epic-07
##fleet forget epic-07 drop it; the panes keep running
rysh board tail --fleet epic-07 # that fleet's stream, not the session's
A fleet may be claudes, codexes, or both — the commands are the same either way.
Path
What
cmd/rysh
the main package — entry point and command surface
internal/tui
the terminal UI
internal/actors
workspace / tab / pane / agent actors, proto.actor over NATS
internal/vterm
terminal emulation, including a vt10x fork with scrollback
internal/provider
LLM provider adapters
internal/platform
host capabilities that change what rysh can do (e.g. PTY support)
action/
the setup-rysh composite GitHub Action
Contributing and security
This repository is a one-way export of a tree developed elsewhere, so a commit
pushed straight here is overwritten by the next export rather than kept. That does
not make patches pointless — CONTRIBUTING.md explains where one
actually lands.
Found a vulnerability? Do not open a public issue, PR or discussion. Rysh runs
shell com

[truncated]
