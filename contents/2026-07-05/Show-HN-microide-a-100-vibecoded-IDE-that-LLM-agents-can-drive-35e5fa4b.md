---
source: "https://pablojimenezmateo.github.io/microide/"
hn_url: "https://news.ycombinator.com/item?id=48794698"
title: "Show HN: microide, a 100% vibecoded IDE that LLM agents can drive"
article_title: "microide — a native, low-footprint desktop IDE"
author: "gef3233"
captured_at: "2026-07-05T15:06:20Z"
capture_tool: "hn-digest"
hn_id: 48794698
score: 1
comments: 0
posted_at: "2026-07-05T14:49:19Z"
tags:
  - hacker-news
  - translated
---

# Show HN: microide, a 100% vibecoded IDE that LLM agents can drive

- HN: [48794698](https://news.ycombinator.com/item?id=48794698)
- Source: [pablojimenezmateo.github.io](https://pablojimenezmateo.github.io/microide/)
- Score: 1
- Comments: 0
- Posted: 2026-07-05T14:49:19Z

## Translation

タイトル: Show HN: microide、LLM エージェントが駆動できる 100% バイブコード化された IDE
記事のタイトル: microide — ネイティブの低フットプリントデスクトップ IDE
説明: エディター、差分、マージ、git、検索、ターミナル、およびデバッグのワークフローに重点を置いたネイティブの低フットプリント C++/SDL3 デスクトップ IDE。オープンソース、プライベート、テレメトリなし。
HN テキスト: 率直に申し上げたいのですが、この IDE のコードは 100% 私が監督した LLM によって生成されており、この事実を隠しているわけではありません。この提出物は人間によって書かれました。 microide はパフォーマンスとプライバシーに重点を置いており、あらゆるツールに AI を統合すること (皮肉なことに、これをバイブコーディングしたことはわかっています)、およびソフトウェアのパフォーマンスの低下に非常にうんざりしていました。かなりの数の IDE を試しましたが、どれも私が望んでいたものではなかったので、microide を始めました。また、microide には設計上、いかなる種類のテレメトリ機能もネットワーキング機能もありません。アカウントも更新チェックもありません。私は CLI ベースの LLM ツールのみを使用しているため、microide には優れた git diff/merge ワークフロー (エージェントが何をしたかを確認するため) と LLM がそれと対話する方法が備わっています。これらは私が使用するプロンプトの一部です: * 「このリポジトリを確認してください。よく知りません。調査するために最も関連性の高いファイルを microide で開き、順番に調査できるようにタブを並べ替えてください。」
* 「メインをプルしてマージし、マイクロライド内の競合を開いて確認してください」
* 「このアプリはクラッシュしています。壊れている場所を見つけて、クラッシュの直前にマイクロライドにブレークポイントを設定して、制御を与えてください」 私はこれを数週間毎日のドライバーとして使用しており、楽しんでいますが、十分に成熟していると思います。現時点では収益化する予定はありません。現時点では、microide は Linux のみなので、macOS にも microide を導入したいと思っていますが、Mac にアクセスできないため、適切にテストできませんでした。

記事本文:
マイクロイド_
デモ
LLM制御
特長
始めましょう
GitHub
100% バイブコーディング — すべてのソース ファイル、テスト、ドキュメントは人間の指示の下、AI コーディング エージェントによって作成されました。
実際のデスクトップ IDE — LLM が実際に駆動できるもの。
エディター、diff、
マージ、git、検索、ターミナル、およびデバッグのワークフロー。シングルウィンドウ、キーボードファースト、
GPU を必要とせずに実行されます。
エディター、3 者間マージ、統合された Git ワークステーション、プロジェクト検索、PTY
ターミナル、および完全な DAP デバッガー — C++20 に組み込まれており、テレメトリやネットワークはありません
コアにリンクされたコード。
このリポジトリ内のすべてのソース ファイル、テスト、ドキュメントは AI コーディングによって作成されました
人間の指示の下、クロード、GPT、その他のツールを組み合わせたエージェント。ありません
手書きのコードパス。エージェント駆動の実世界の実験として公開されています
開発: 読むのは面白く、構築するのに役立ちます。これはコードがどのようなものであったかを説明します
完成度ではなく、作成されたものです。ビルドは安定しており、v2.6.4 とタグ付けされています。
121 のテスト ファイルに加えて、ファジング、サニタイザー、パフォーマンス ゲートによってカバーされています。
いくつかの確固たる原則を中心に構築
コモディティ ハードウェアでの応答性を重視して構築されており、ソフトウェア レンダラー上で実行され、
GPU が存在する場合、テキストを高速化するためにのみ GPU を使用します。ホットパス - 起動、入力、
スクロール、差分、検索 - 自動パフォーマンス ハーネスによって保護されています。
コミットされたベースラインのため、リグレッションが発生すると、出荷ではなくビルドが失敗します。
プライベートおよびローカル — テレメトリーなし
コア エディタはネットワーク接続を行わず、ネットワーク ライブラリもリンクされません。
アカウントもテレメトリも更新チェックも、家に電話することもありません。
コード、プロジェクトの状態、セッションがマシンから離れることはありません。何もありません
何も収集していないため、オプトアウトする必要があります。
の下でリリースされました

MITライセンス。読んで、構築して、フォークして、出荷してください。
問題やプル リクエストは GitHub で受け付けています。
GPU を必要とせずに実行されるシングル ウィンドウのネイティブ アプリ。
コマンド パレットがすべてのアクションの先頭に表示されるため、開いたままキーボードを操作し続けることができます
そして、同じコマンドを制御チャネル経由でスクリプト化できます。
LLM (別の種類の AI エディター) で IDE を駆動する
ほとんどの「AI IDE」はチャットボットをサイドパネルにボルトで固定します。 microide はその逆です。
コマンド パレット全体を外部制御チャネル経由で公開するため、LLM エージェントは
実際にエディタを操作することができます。ファイルを開いたり、ブレークポイントを設定したり、起動してステップを実行したりできます。
デバッガを実行し、プログラムの状態を読み取り、結果を JSONL としてストリーミングして返します。モデル
IDE を推測するのではなく、実際の IDE を駆動します。
すべてがコマンド パレットと同じチョークポイントを介してルーティングされます。チャネルにより、
トランスポートとボキャブラリーであり、並列コードパスではありません。
完全な参照:
コントロールチャネル.md 。
# 実行中のインスタンスを駆動し、結果を JSONL としてストリーミングします
microide --control --control-spec /tmp/debug.spec.json
# ワンショットクライアント — ソケット配管や手書きの JSON は不要
microide control-send ブレークポイント-function-add main
microide control-send --query launch-configs
microide control-send debug-run ./build/app --wait 停止しました
# エージェントが stdout で読み取る内容:
{ "イベント" : "準備完了" 、 "pid" :1234、 "ソケット" : "…" }
{ "id" :1、"ok" :true、"フィードバック" : "…" }
{ "イベント" : "停止" 、 "理由" : "ブレークポイント" 、
"ファイル" : "src/main.cpp" 、 "行" :42、 "スレッド ID" :1}
上記のコマンドとプロトコルはプロジェクトのドキュメントからのものであり、例示的なものではありません。
ブレークポイントを設定し、起動し、チャネル全体をステップ実行します
( ブレークポイントセット 、 デバッグ実行 、
debug-step-over ) エージェントが読み取りを行っている間
デバッグ状態と停止に反応する
イベント — エイダが発生したときに後退することさえあります

pterはそれをサポートします。
モデルに失敗した実行を渡し、ブレークポイントを設定させ、続行し、検査します
ローカルのユーザーは、アダプターがサポートしている場合は、ユーザーなしで後退することもできます。
マウスに触れている。
ワンコマンドで再現可能なセットアップ
--control-spec JSON は、ブレークポイントと
起動設定は、ウィンドウが対話型になる前にすでに適用されています。
1 つの仕様ファイルをリポジトリにチェックインすると、チームメイトでもエージェントでも、誰でもその仕様ファイルを再度開くことができます。
最初の起動時とまったく同じデバッグ シナリオ。
チャネルはコマンド パレットを反映しているため、どのツールでもエディターをヘッドレスで駆動できます。
ファイルを開き、コマンドを実行し、状態をクエリし、JSONL イベントをストリーミングして返します。
CI スモーク テスト、再現可能なバグ レポート、またはカスタム エージェント ループでそれに到達します。
それ以外の場合は、人間が UI をクリックするスクリプトを作成する場所であればどこでも使用できます。
以下のすべてのワークフローは現在実装されており、回帰テストが行​​われていますが、ロードマップではありません。
マルチプロジェクトとファイルのタブ、ネストされた分割
UTF-8 対応、IME プリエディット、行末保持
マルチキャレット編集と領域の強調表示
チェックポイントが設定された状態での構文の強調表示
単語レベルの合体による元に戻す/やり直し
再起動後のセッションの復元
作業ツリーと HEAD または任意のコミットを比較する
ハンクごとのピックによる 3 方向マージ
共有装飾テキストグリッド パイプライン
スクロールバーの横の変更概要レーン
サイドバー: 変更、ステージング、競合、発信
ファイルごとのステージ/破棄、一括ステージ-オール
競合はマージタブに直接表示されます
編集可能なコミットメッセージ、ブランチ/コミットピッカー
並列化されたプロジェクト検索、リテラルまたは正規表現
すべての合計をカウントし、一致を強調表示します
プロジェクト内で置換 (リテラル モード)
キャッシュされたインデックスを使用したファイル ファインダー オーバーレイ
スクロールバックと選択機能を備えた PTY ベースのタブ
代替画面、括弧付きペースト、OSC 52 コピー
フル DAP クライアント、マルチセッション
行/条件/関数ブレークポイント、ログポイント

ステップオーバー/イン/アウト、一時停止、再開
逆続行とステップバック (サポートされている場合)
コールスタック、変数、監視、ホバーして検査
デバッグコンソールREPL・バンドルされたgdb-dap
ライフサイクルフック、コマンド、ツリーサイドバー
エディターの装飾: とじしろ、EOL テキスト、コード レンズ
コンテンツの表示: グラフとプレビュー、インラインまたはパネル内
ゴーストテキストの提案とホスト所有のバッファ編集
言語プロバイダー: defs、refs、署名、シンボル
テーマ、ファイルアイコンテーマ、豊富なステータスアイテム
UI スレッドから実行されます。ホットリロード (Linux)
SHA 検証を備えたツール ダウンローダー
ネイティブ ファイル監視バックエンド (Linux 上の inotify)
IDE に期待されるすべて
退屈な基本は正しく行われ、主張されるだけでなくテストスイートによって検証されます。
構文の強調表示 — チェックポイント付き、大きなファイルでも高速
マルチタブ編集 - プロジェクト、ファイル、分割
分割ペイン — ネストされた共有バッファー
プロジェクト全体の検索 — リテラルと正規表現、置換
ファイル ファインダー - キャッシュされたファジー オーバーレイ (Ctrl+P)
統合された git — ステージ、コミット、差分、マージ
統合ターミナル — 本物の PTY で裏打ちされたタブ
デバッグ — ブレークポイント、ステッピング、変数
マルチキャレット編集 - 位置再マップあり
セッションの復元 — 中断したところから再開します
キーボードファースト — VSCode スタイルのショートカット、残りはコマンド パレット
何を保護するのか、何を保護しないのかについて正直に語る
プラグインはインプロセスで実行されますが、プラグインごとに強制された機能サンドボックス: ファイルシステムの下で実行されます。
アクセスはプロジェクトに含まれており、プロセスの実行はデフォルトで拒否され、レンダリングが行われます。
コントリビューションは、ホストが描画するサイズ制限のあるデータです。それは意味を持って含まれています - ではありません
完全な隔離 — そしてドキュメントにははっきりとそう書かれています。
プロジェクトルートに含まれるファイルシステムへのアクセス
プロセスの実行はデフォルトで拒否されます
process.exec なしで拒否されたフォーマッタ/LSP/タスク
レンダリングの貢献はサイズ上限に達します

d ホストが描画するデータ — プラグインがフレームに触れることはありません
Landlock で制限されたプラグインの子 + 言語サーバー
書き込みはプロジェクト + プラグイン データ ディレクトリに限定されます
オプションの seccomp IPv4/IPv6 ソケット ブロック
ユーザースコープのプラグインのみ - リポジトリローカルのプラグインツリーは無視されます
クローンされたリポジトリを開いても、そこからプラグイン コードが実行されることはありません
--safe-mode / --disable-plugins 不慣れなリポジトリ用
エディターはネットワーク接続を行いません。 HTTP が存在しない、または
ネットワーク ライブラリはバイナリにリンクされています。言語サーバー、デバッガー、および
LLM 制御チャネルはすべてローカルです。標準入出力上の子プロセス、またはユーザープライベートです。
AF_UNIX ソケット。何かがネットワークに到達する唯一の方法は、
ネットワークを明示的に宣言するプラグインをインストールする
能力;これはデフォルトの拒否であり、Linux では seccomp でハードブロックできます。
正直に言うと、プラグインの署名、マーケットプレイスの信頼、初回実行はありません。
機能プロンプトが表示されますが、Lua 状態は引き続きインプロセスで実行されます — プラグインが許可されています
process.exec は引き続きプロジェクトを読み取るツールを実行できます。
(リリースバイナリ自体は GPG 署名されています - 署名を検証し、
インストール前のチェックサム。)
フルモデル:
セキュリティ.md 。
1 つのシングル ウィンドウ シェル — エディター、ソース コントロール、デバッガー、およびスクリプト可能コントロール
チャンネル。これらは実行中のアプリから直接キャプチャされ、毎回再生成されます。
リリースするので、構築したものから逸脱することはありません。
退屈で信頼できるビルディングブロック
署名付き Debian パッケージをリリースから取得するか、2 つのコマンドでソースからビルドします。
事前に構築されたパッケージをインストールする
— Linux (amd64)
# リリースから .deb、その .asc 署名、.sha256、および microide-signing-key.asc をダウンロードします。
gpg --import microide-signing-key.asc
gpg --verify microide_<バージョン>_amd64.deb.asc microide_<バージョン>_amd64.deb
sha256sum -c microide_<バージョン>_amd64.deb.sha256
突然

o dpkg -i microide_<バージョン>_amd64.deb
最新の .deb をダウンロード
各リリースには、事前構築された .deb 、分離された GPG が添付されます。
署名とチェックサム — インストールする前に両方を確認するか、以下のソースからビルドしてください。
ソースからビルドする
# Deps (Debian/Ubuntu)。 PCRE2 は必須です。 Lua/ttf/fontconfig を推奨します。
sudo apt-get install -y cmake ninja-build pkg-config \
libsdl3-dev libsdl3-ttf-dev libpcre2-dev lilua5.4-dev libfontconfig-dev
cmake -S 。 -B ビルド
cmake --build build -j8
./build/microide/microide
自分でパッケージ化する
# 独自の .deb をビルドしてインストールする (Linux)
./scripts/package-deb.sh
sudo ./scripts/install-deb.sh
Linux (amd64) が現在サポートされているターゲットです。他のプラットフォーム上のソースからビルドします。
GitHub 上のファイルの問題と機能リクエスト。
MIT ライセンス · © 2026 パブロ ヒメネス マテオ
GitHub ·
LLM制御・
ライセンス
人間の指示の下、AI コーディング エージェントによって作成されます。

## Original Extract

A native, low-footprint C++/SDL3 desktop IDE focused on editor, diff, merge, git, search, terminal, and debugging workflows. Open source, private, no telemetry.

I want to be very upfront: this IDE's code is 100% generated by LLMs supervised by me, I am not hiding this fact. This submission was written by a human. microide's focus is performance and privacy, I was very tired of having AI integrated on every tool (ironic, I know, having vibecoded this one), and the decline in software performance. After trying quite a few IDEs none were what I wanted so I started microide. microide also does not have any kind of telemetry nor networking capabilities by design, no accounts, no update checks, none. Since I only use CLI-based LLM tools, microide has a good git diff/merge workflow (for me to check what the agents did) and a way for the LLMs to interact with it, these are a few of the prompts I use: * 'Check this repo, I am not familiar with it, open in microide the most relevant files to study, reorder the tabs so I can study them in order'
* 'Pull main and merge, open any conflicts in microide for me to review'
* 'This app is crashing, find where it breaks and set a breakpoint in microide just before the crash and give me the control' I am using it as my daily driver for a couple of weeks and enjoying it, I think it is mature enough to show. I do not plan to monetize it at any point. At the moment microide is Linux only, I would love to bring microide to macOS as well, but since I do not have access to a Mac I could not test it properly.

microide _
Demo
LLM control
Features
Get started
GitHub
100% vibecoded — every source file, test, and doc was written by AI coding agents under human direction.
A real desktop IDE — that an LLM can actually drive.
A native, low-footprint C++/SDL3 desktop IDE focused on editor, diff,
merge, git, search, terminal, and debugging workflows. Single-window, keyboard-first,
and runs without requiring a GPU.
Editor, three-way merge, an integrated git workstation, project search, a PTY
terminal, and a full DAP debugger — built in C++20, with no telemetry and no network
code linked into the core.
Every source file, test, and document in this repository was written by AI coding
agents — a mix of Claude, GPT, and other tools — under human direction. There is no
hand-written code path. It's published as a real-world experiment in agent-driven
development: interesting to read, useful to build on. That describes how the code was
authored , not how mature it is — the build is stable, tagged v2.6.4, and
covered by 121 test files plus fuzzing, sanitizer, and performance gates.
Built around a few firm principles
Built for responsiveness on commodity hardware: it runs on a software renderer and
uses the GPU only to speed up text when one is present. Hot paths — startup, typing,
scrolling, diff, and search — are guarded by an automated performance harness with
committed baselines, so regressions fail the build instead of shipping.
Private & local — no telemetry
The core editor makes no network connections — no networking library is even linked.
No accounts, no telemetry, no update checks, no phoning home.
Your code, project state, and session never leave the machine. There is nothing
to opt out of, because nothing is collecting anything.
Released under the MIT license. Read it, build it, fork it, ship it.
Issues and pull requests are welcome on GitHub.
A single-window native app that runs without requiring a GPU.
A command palette fronts every action, so you can stay on the keyboard from open
to commit — and the same commands are scriptable over the control channel.
Drive the IDE with an LLM — a different kind of AI editor
Most "AI IDEs" bolt a chatbot into a side panel. microide does the opposite: it
exposes the whole command palette over an external control channel, so an LLM agent
can actually operate the editor — open files, set breakpoints, launch and step the
debugger, and read program state — then stream the results back as JSONL. The model
drives a real IDE instead of guessing at one.
Everything routes through the same chokepoint as the command palette — the channel adds a
transport and vocabulary, not a parallel code path.
Full reference:
control-channel.md .
# Drive a running instance and stream results as JSONL
microide --control --control-spec /tmp/debug.spec.json
# One-shot client — no socket plumbing, no hand-written JSON
microide control-send breakpoint-function-add main
microide control-send --query launch-configs
microide control-send debug-run ./build/app --wait stopped
# What the agent reads back on stdout:
{ "event" : "ready" , "pid" :1234, "socket" : "…" }
{ "id" :1, "ok" :true, "feedback" : "…" }
{ "event" : "stopped" , "reason" : "breakpoint" ,
"file" : "src/main.cpp" , "line" :42, "threadId" :1}
Commands & protocol above are from the project docs — real, not illustrative.
Set breakpoints, launch, and step entirely over the channel
( breakpoint-set , debug-run ,
debug-step-over ) while the agent reads
debug-state and reacts to stopped
events — even stepping backwards when the adapter supports it.
Hand the model a failing run, let it set a breakpoint, continue to it, inspect
locals, and step — backwards too, when the adapter supports it — without you
touching the mouse.
One-command, reproducible setups
A --control-spec JSON opens a project with breakpoints and a
launch config already applied before the window is even interactive.
Check one spec file into your repo and anyone — teammate or agent — can reopen the
exact same debug scenario on the first launch.
The channel mirrors the command palette, so any tool can drive the editor headlessly —
open files, run commands, query state — and stream JSONL events back.
Reach for it in CI smoke tests, reproducible bug reports, or custom agent loops —
anywhere you'd otherwise script a human clicking through the UI.
Every workflow below is implemented and regression-tested today — not a roadmap.
Multi-project & file tabs, nested splits
UTF-8 aware, IME preedit, line-ending preservation
Multi-caret editing & region highlighting
Syntax highlighting with checkpointed state
Undo/redo with word-level coalescing
Session restore across restarts
Compare working tree vs HEAD or any commit
Three-way merge with per-hunk picks
Shared decorated text-grid pipeline
Change-overview lane beside the scrollbar
Sidebar: changes, staged, conflicts, outgoing
Per-file stage / discard, bulk stage-all
Conflicts open straight into a merge tab
Editable commit message, branch/commit picker
Parallelized project search, literal or regex
Count-all totals with match highlighting
Replace-in-project (literal mode)
File finder overlay with cached index
PTY-backed tabs with scrollback & selection
Alternate-screen, bracketed paste, OSC 52 copy
Full DAP client, multi-session
Line / conditional / function breakpoints, logpoints
Step over / in / out, pause, restart
Reverse continue & step back (when supported)
Call stack, variables, watch, hover-to-inspect
Debug console REPL · bundled gdb-dap
Lifecycle hooks, commands, tree sidebars
Editor decorations: gutter marks, EOL text, code lenses
Content surfaces: charts & previews, inline or in a panel
Ghost-text suggestions & host-owned buffer edits
Language providers: defs, refs, signature, symbols
Themes, file-icon themes & rich status items
Runs off the UI thread; hot reload (Linux)
Tool downloader with SHA verification
Native file-watch backend (inotify on Linux)
Everything you expect from an IDE
The boring fundamentals, done right — and verified by the test suite, not just claimed.
Syntax highlighting — checkpointed, fast in large files
Multi-tab editing — projects, files, and splits
Split panes — nested, shared-buffer
Project-wide search — literal & regex, replace
File finder — cached, fuzzy overlay (Ctrl+P)
Integrated git — stage, commit, diff, merge
Integrated terminal — real PTY-backed tabs
Debugging — breakpoints, stepping, variables
Multi-caret editing — with position remap
Session restore — picks up where you left off
Keyboard-first — VSCode-style shortcuts, command palette for the rest
Honest about what it does — and doesn't — protect
Plugins run in-process, but under an enforced per-plugin capability sandbox: filesystem
access is contained to the project, process execution is default-deny, and rendering
contributions are size-capped data the host draws. It's meaningfully contained — not
full isolation — and the docs say so plainly.
Filesystem access contained to the project root
Process execution is default-deny
Formatters / LSPs / tasks rejected without process.exec
Rendering contributions are size-capped data the host draws — plugins never touch the frame
Plugin children + language servers confined with Landlock
Writes limited to the project + plugin data dir
Optional seccomp IPv4/IPv6 socket block
User-scope plugins only — repo-local plugin trees are ignored
Opening a cloned repo never runs plugin code from it
--safe-mode / --disable-plugins for unfamiliar repos
The editor makes no network connections. No HTTP or
networking library is linked into the binary at all. Language servers, the debugger, and
the LLM control channel are all local — child processes over stdio, or a user-private
AF_UNIX socket. The only way anything reaches the network is a
plugin you install that explicitly declares the network
capability; it's default-deny, and on Linux it can be hard-blocked with seccomp.
In the interest of honesty: there's no plugin signing, no marketplace trust, no first-run
capability prompt, and the Lua state still runs in-process — a plugin granted
process.exec can still run tools that read your project.
(Release binaries themselves are GPG-signed — verify the signature and
checksum before installing.)
Full model:
SECURITY.md .
One single-window shell — editor, source control, debugger, and a scriptable control
channel. These are captured straight from the running app and regenerated on every
release, so they never drift from what you build.
Boring, dependable building blocks
Grab the signed Debian package from Releases, or build from source in two commands.
Install the prebuilt package
— Linux (amd64)
# Download the .deb, its .asc signature, .sha256, and microide-signing-key.asc from Releases
gpg --import microide-signing-key.asc
gpg --verify microide_<version>_amd64.deb.asc microide_<version>_amd64.deb
sha256sum -c microide_<version>_amd64.deb.sha256
sudo dpkg -i microide_<version>_amd64.deb
Download latest .deb
Each release attaches a prebuilt .deb , a detached GPG
signature, and a checksum — verify both before installing, or build from source below.
Build from source
# Deps (Debian/Ubuntu). PCRE2 is required; Lua/ttf/fontconfig recommended.
sudo apt-get install -y cmake ninja-build pkg-config \
libsdl3-dev libsdl3-ttf-dev libpcre2-dev liblua5.4-dev libfontconfig-dev
cmake -S . -B build
cmake --build build -j8
./build/microide/microide
Package it yourself
# Build & install your own .deb (Linux)
./scripts/package-deb.sh
sudo ./scripts/install-deb.sh
Linux (amd64) is the supported target today; build from source on other platforms.
File issues and feature requests on GitHub .
MIT License · © 2026 Pablo Jiménez Mateo
GitHub ·
LLM control ·
License
Written by AI coding agents under human direction.
