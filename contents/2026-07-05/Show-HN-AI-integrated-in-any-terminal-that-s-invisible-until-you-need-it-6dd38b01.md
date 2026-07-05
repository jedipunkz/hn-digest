---
source: "https://terminai.app"
hn_url: "https://news.ycombinator.com/item?id=48799134"
title: "Show HN: AI integrated in any terminal that's invisible until you need it"
article_title: "Terminai | Transparent terminal wrapper to integrate AI agents"
author: "emosenkis"
captured_at: "2026-07-05T23:57:40Z"
capture_tool: "hn-digest"
hn_id: 48799134
score: 2
comments: 0
posted_at: "2026-07-05T23:46:24Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI integrated in any terminal that's invisible until you need it

- HN: [48799134](https://news.ycombinator.com/item?id=48799134)
- Source: [terminai.app](https://terminai.app)
- Score: 2
- Comments: 0
- Posted: 2026-07-05T23:46:24Z

## Translation

タイトル: HN を表示: 必要になるまで表示されないあらゆる端末に統合された AI
記事タイトル: テルミナイ | AI エージェントを統合するための透過的なターミナル ラッパー
説明: Terminai は、Codex、Claude Code、およびカスタム CLI エージェント用の AI ターミナル オーバーレイです。
HN テキスト: 私はシェルを透過的にラップするオープンソースのターミナル アプリを構築しました。これにより、必要なときにいつでも AI にアクセスでき、それ以外の時間は完全に邪魔にならないようになります。ターミナルのコンテンツへの自動構成された MCP アクセスとスクロールバックを備えた独自の AI エージェント (Claude と Codex はすぐにサポートされます) を導入し、あたかも自分で入力したかのようにシェルに渡すことを承認できる入力を提案する機能を備えています。技術的に最も困難な部分は、これを透明にすることでした。ホスト端末のネイティブ スクロールバックに行をプッシュすると同時に、AI オーバーレイをアクティブ化するときにフルフレームでレンダリングされた TUI を使用できるようにする必要がありました。それとイベント処理の 2 つの部分はクロードが正しく理解できなかったので、私はそれらの重要な部分を手作業で実装することになりました。まだアルファ版の品質ですが、Mac の iTerm2 と Terminal.app、Linux のYakuake/Konsole で毎日のドライバーとして使用しています。 Windows は現時点ではサポートされていません。

記事本文:
t
テルミナイ
ワークフロー
エージェント
インストール
構成
GitHub
端末をラップして AI にオンデマンドでアクセス
ターミナルを任意の方法で使用し、Ctrl + Space を押して Codex、Claude Code、
または、端末への読み取りおよび承認ゲート型書き込みアクセスを持つオーバーレイ内の独自のエージェント。
それ自体を説明できるシェル
Terminai は、助けが必要になるまで邪魔をしません。オーバーレイは、ラップされたシェルへのアクセスが制御された、すでに使用しているエージェントを実行する実際の端末です。
Terminai は、モデルの資格情報やプロバイダー ルーティングを所有しません。設定された CLI は認証とモデルの選択を処理します。 Terminai は、ターミナル ラッパー、MCP サーバー、および承認フローを提供します。 Terminai がデータを収集したり、発信ネットワーク接続を確立したりすることはありません。
バンドルされたプリセットは、Terminai MCP ツールとコンテキストを Codex に挿入します。
バンドルされたプリセットは、事前設定された Terminai MCP で Claude を開始します。
任意の CLI (理想的には、MCP URL とコンテキスト プロンプトを使用できるもの) で Terminai を指定します。 config を参照してください。
Terminai をインストールし、任意の AI CLI にサインインして、terminai または terminai -- <command> [arg1 arg2...] を開始するプロファイルを選択したターミナル エミュレーターに追加します。
醸造タップ emosenkis/terminai https://github.com/emosenkis/terminai.git
brew trust --formula emosenkis/terminai/terminai
brew install emosenkis/terminai/terminai
ソースからビルドする
git clone https://github.com/emosenkis/terminai.git
git サブモジュールの初期化
gitサブモジュールの更新
CD端末
カーゴインストール --path src
GitHub リリース
https://github.com/emosenkis/terminai/releases からダウンロードして解凍します。
謝辞
優れた端子ツールをベースに構築
Terminai は当初、次からフォークされました。
vt100 を提供する mprocs
ホストとゲストのサポート。また、独自のフォークにも依存します。
ラタトゥイと
加えるラットサルサ
ネイティブ スクロールを維持するために必要な機能

d コピーします。
テルミナイは、クロードとコーデックスを使用して（強固な基礎の上に）大部分が建設されました。
mprocs によって提供されます）AI が最も苦労した部分を克服するためにいくつかの手書き部分が含まれています。
ネイティブのスクロールと入力処理。これはアルファ品質のソフトウェアであり、私はそれを日常的に使用していますが、
ドライバーを使用する場合は、通常の AI フリー シェルを生成する簡単な方法があることを確認してください。
1. ゼロインストール MCP 構成は、CLI フラグまたは環境変数を介して MCP サーバー設定をサポートするエージェント CLI に依存します。

## Original Extract

Terminai is an AI terminal overlay for Codex, Claude Code, and custom CLI agents.

I built an open source terminal app that transparently wraps your shell so you can have access to AI whenever you need it, while it stays completely out of your way the rest of the time. Bring your own AI agent (Claude and Codex are supported out of the box) with auto-configured MCP access to the contents of your terminal and scrollback and the ability to suggest input that you can approve passing to your shell as if you had typed it yourself. The hardest part technically was getting it to be transparent - it needs to push lines into the host terminal's native scrollback while also being able to use full frame-rendered TUI when you activate the AI overlay. That and event handling were the two parts that Claude couldn't get right, leading me to implement key parts of them by hand. It's still alpha quality but I'm using it as my daily driver in iTerm2 and Terminal.app on Mac and Yakuake/Konsole on Linux. Windows is not supported as of yet.

t
Terminai
Workflow
Agents
Install
Config
GitHub
Wrap your terminal for on-demand access to AI
Use your terminal any way you want, then press Ctrl + Space to open Codex, Claude Code,
or your own agent in an overlay with read and approval-gated write access to your terminal.
A shell that can explain itself
Terminai stays out of the way until you need help. The overlay is a real terminal running the agent you already use, with controlled access to the wrapped shell.
Terminai does not own your model credentials or provider routing. Your configured CLI handles auth and model selection; Terminai provides the terminal wrapper, MCP server, and approval flow. Terminai never collects your data or makes outgoing network connections.
Bundled preset injects Terminai MCP tools and context into Codex.
Bundled preset starts Claude with Terminai MCP preconfigured.
Point Terminai at any CLI (ideally one that can consume the MCP URL and context prompt). See config .
Install Terminai, sign in to your preferred AI CLI, then add a profile to your terminal emulator of choice that starts terminai or terminai -- <command> [arg1 arg2...]
brew tap emosenkis/terminai https://github.com/emosenkis/terminai.git
brew trust --formula emosenkis/terminai/terminai
brew install emosenkis/terminai/terminai
Build from source
git clone https://github.com/emosenkis/terminai.git
git submodule init
git submodule update
cd terminai
cargo install --path src
GitHub releases
Download and unpack from https://github.com/emosenkis/terminai/releases
Acknowledgements
Built on excellent terminal tooling
Terminai was initially forked from
mprocs , which provided the vt100
host & guest support. It also relies on its own forks of
Ratatui and
rat-salsa which add
functionality needed to preserve native scroll and copy.
Terminai was built to a great extent using Claude and Codex (on top of the solid foundation
provided by mprocs) with some hand-written portions to get past the parts AI struggled with most,
native scrolling and input handling. It is alpha-quality software and although I use it as my daily
driver, please ensure if you use it that you have an easy way to spawn a regular AI-free shell.
1. Zero-install MCP configuration depends on the agent CLI supporting MCP server settings via CLI flags or environment variables.
