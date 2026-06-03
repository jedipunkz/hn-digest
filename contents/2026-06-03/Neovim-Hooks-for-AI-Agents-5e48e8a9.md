---
source: "https://github.com/NishantJoshi00/sidekick"
hn_url: "https://news.ycombinator.com/item?id=48369534"
title: "Neovim Hooks for AI Agents"
article_title: "GitHub - NishantJoshi00/sidekick · GitHub"
author: "cat-whisperer"
captured_at: "2026-06-03T00:47:27Z"
capture_tool: "hn-digest"
hn_id: 48369534
score: 4
comments: 0
posted_at: "2026-06-02T12:46:02Z"
tags:
  - hacker-news
  - translated
---

# Neovim Hooks for AI Agents

- HN: [48369534](https://news.ycombinator.com/item?id=48369534)
- Source: [github.com](https://github.com/NishantJoshi00/sidekick)
- Score: 4
- Comments: 0
- Posted: 2026-06-02T12:46:02Z

## Translation

タイトル: AI エージェント用 Neovim フック
記事タイトル: GitHub - NishantJoshi00/sidekick · GitHub
説明: GitHub でアカウントを作成して、NishantJoshi00/サイドキックの開発に貢献します。

記事本文:
GitHub - NishantJoshi00/サイドキック · GitHub
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
ニシャントジョシ00
/
相棒
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード さらにアクション メニューを開く 折りたたむ

ファイルとファイル
91 コミット 91 コミット .claude-plugin .claude-plugin .github/ workflows .github/ workflows docs docs フック フック ロゴ ロゴ プラグイン プラグイン スクリプト スクリプト src src テスト テスト .gitignore .gitignore CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンスライセンス PHILOSOPHY.md PHILOSOPHY.md README.md README.md build.rs build.rs すべてのファイルを表示 リポジトリ ファイルのナビゲーション
保存されていない Neovim の作業を Claude Code、Codex、opencode、pi、Crush、Amp、Antigravity、Grok から保護します。
Neovim と AI エージェントの間のパイプライン - したがって、AI エージェントは、ユーザーが入力している間待機します。
ブロックを表示するには 1:20 までスキップしてください。同じ記録がバイナリにバンドルされています。インストール後にサイドキック デモを実行して、オフラインで再生します。
手作業でコーディングすることはますます少なくなりました。私が今でもやっている部分は、非常に慎重に行っています。新しいペインを開き、Neovim を開いて、編集を開始します。自分自身を追い込む必要があるフロー状態です。問題点: Neovim は、システム内で私が入力していることを認識している唯一のものです。次のペインのエージェントが、私がいるファイルを編集する必要があると判断すると、バッファがリロードされ、私の作業は失われます。そこで私はそれについて何かをすることにしました。
サイドキックは編集者とエージェントの間のパイプ役です。バッファに保存されていない変更がある場合、エージェントは待機します。編集は拒否され、バッファは変更されません。ファイルを保存すると、次の試行が続行されます。フラグ、確認プロンプト、ポリシー ファイルはありません。あなたと競合しない編集の 99% はそのまま残ります。
他の方向も機能します。AI が開いているファイルを変更すると、Sidekick はすべての Neovim インスタンスのバッファを更新し、カーソル位置は保持されます。
ファイルの編集中、AI はファイルを待機します。 Neovim では、「編集がブロックされました — ファイルに保存されていない変更があります」というメッセージが表示されます。
ファイルを開いているときに AI が編集するファイルは、自動的に再ロードされます。いいえ、:e!踊る

e.
Neovim の現在または最近のビジュアル選択は、次の Claude Code、Codex、opencode、または pi プロンプトにコンテキストとして追加できます。コードを選択し、プロンプトを入力して Enter キーを押します。
他はすべて同じままです。通常どおり nvim を使い続けます。
カーゴインストールサイドキック
サイドキックの初期化
サイドキックの init は、マシン上で見つかったすべての AI ハーネスを調べ、フックを配線し、nvim エイリアスを設定し、すでに行われたことはすべてスキップします。いつでも再実行できます。冪等です。
サイドキックドクターに確認してください。手作業で配線してみませんか? docs/SETUP.md を参照してください。
ハーネス
編集拒否
バッファリフレッシュ
選択注入
クロード・コード
✓
✓
✓
コーデックス
✓
✓
✓
オープンコード
✓
✓
✓
円周率
✓
✓
✓
クラッシュ
✓
—
—
アンプ
✓
✓
—
反重力
✓
✓
—
グロク
✓
✓
—
A — 上流のギャップです。ハーネスは、サイドキックが必要とするフック イベントを公開しません。
nvim を使用してください。シェルのエイリアスはサイドキックを介してルーティングされるため、フックはエディターを見つけることができます。
nvim src/main.rs
エイリアスが不要な場合は、sidekick neovim <args> を直接実行してください。エージェントが保存されていない作業内容を上書きしようとするまでは、何も変わったことに気づきません。その時点で、サイドキックは編集を拒否し、その理由をエージェントに伝えます。
コマンド
何をするのか
サイドキック ネオビム <args>
フックが見つけることができるディレクトリごとのソケットを使用して Neovim を起動します。 nvim というエイリアス。
サイドキックの初期化
ガイド付きセットアップ。検出されたすべての AI ハーネスと nvim エイリアスを配線します。
相棒ドクター [--修正] [--色なし]
健康診断。 --fix は同意ゲート型修復を提供し、書き込み前に差分を表示します。
サイドキック統計 [--範囲 週|月|年|すべて]
追加専用の JSONL イベントからのローカル アクティビティ ダッシュボード。マシンからは何も残りません。
サイドキックのデモ
ラタトゥイ フレームでデモ キャストをインラインで再生します。同僚に見せるのに便利です。
サイドキックフック
内部。 AI ハーネスがそれを呼び出します。あなたはしない。
仕組み
相棒ネオビムが発売

nvim --listen /tmp/<blake3(cwd)>-<pid>.sock 。ソケット パスは正規の作業ディレクトリごとに決定的であり、プロセスごとに一意であるため、フックは同じプロジェクトから開かれたすべての Neovim インスタンスを見つけることができます。
ファイルを編集する前に、エージェントはサイドキック フック を呼び出します。フックは /tmp/<blake3(cwd)>-*.sock をグロブし、短いタイムアウトで msgpack-rpc 経由で到達可能なインスタンスに接続し、各ターゲットがアクティブで未保存の変更があるかどうかを尋ねます。 「はい」の場合、編集は拒否されます。それ以外の場合は許可されます。 Neovim ソケットが見つからない場合、サイドキックは許可するように機能を低下させます。
編集が完了すると、フックはファイルを開いた状態で到達可能なすべての Neovim インスタンスにファイルを再ロードするように指示します。カーソルの位置と表示されているウィンドウは保持されます。
プロンプト送信時に、Neovim にライブビジュアル選択または最​​近のビジュアルマークがある場合、サイドキックは [Selected from path:start-end] のようなフェンスで囲まれたコンテキスト ブロックを返します。 Claude Code と Codex は追加のコンテキストとしてそれらを受け取ります。 opencode ブリッジと pi ブリッジは、送信されたプロンプト テキストにそれらを追加します。
決定、更新、Neovim の起動、および統計ビューは、OS データ ディレクトリの下の sidekick/events.jsonl にローカルで追加されます。書き込みはベストエフォートで行われ、フック パスをブロックすることはありません。
デーモンもバックグラウンド サービスもありません。 1 つの CLI、 /tmp 内の Neovim RPC ソケット、およびオプションのツールごとのブリッジ ファイルのみ。
RPC + Lua サポートを備えた Neovim、Unix ライク システム、および少なくとも 1 つのサポートされる AI ハーネス: Claude Code、Codex、opencode、pi、Crush、Amp、Antigravity、または Grok。 Rust/Cargo はカーゴのインストールに必要です。
サイドキックはフェーズ 1 — Neovim と Claude Code、Codex、opencode、pi、Crush、Amp、Antigravity、および Grok です。長いアークは、どの編集者でも公開でき、AI ツールでも書き込み前にクエリできる小さなプロトコルです。このファイルは今人間によって編集されていますか? PHILOSOPHY.md にはロードマップがあります (Helix、Zed、VS Code; Aider、Goose、C)

ontinue) と拡張ポイント。
問題やPRを歓迎します。新しいエディターのサポートは、 src/editor.rs の Editor 特性から始まります。新しい AI ハーネスのサポートは、 src/harness.rs の Harness トレイトから始まります。 PHILOSOPHY.md はアーキテクチャをカバーしています。
読み込み中にエラーが発生しました。このページをリロードしてください。
2
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to NishantJoshi00/sidekick development by creating an account on GitHub.

GitHub - NishantJoshi00/sidekick · GitHub
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
NishantJoshi00
/
sidekick
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
91 Commits 91 Commits .claude-plugin .claude-plugin .github/ workflows .github/ workflows docs docs hooks hooks logo logo plugins plugins scripts scripts src src tests tests .gitignore .gitignore CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE PHILOSOPHY.md PHILOSOPHY.md README.md README.md build.rs build.rs View all files Repository files navigation
Protects your unsaved Neovim work from Claude Code, Codex, opencode, pi, Crush, Amp, Antigravity, and Grok.
A conduit between Neovim and your AI agents — so they wait when you're typing.
Skip to 1:20 to see a block. The same recording is bundled in the binary — run sidekick demo after install to play it back offline.
I do less and less coding by hand. The part I still do, I do very deliberately — pop open a new pane, open Neovim, start editing. It's a flow state I have to push myself into. The catch: Neovim is the only thing in the system that knows I'm typing. When the agent in the next pane decides the file I'm in needs editing, the buffer reloads, my work gone. So I decided to do something about it.
Sidekick is the conduit between the editor and the agents. When you have unsaved changes in a buffer, the agents wait — the edit is denied, your buffer is untouched. Save the file and the next attempt proceeds. No flags, no confirmation prompts, no policy file. The 99% of edits that don't conflict with you go through untouched.
The other direction works too: when the AI modifies a file you have open, Sidekick refreshes the buffer in every Neovim instance, cursor position preserved.
The AI waits on a file while you're editing it. You'll see this in Neovim: Edit blocked — file has unsaved changes .
A file the AI edits while you have it open is auto-reloaded — no :e! dance.
A current or recent visual selection in Neovim can be added to your next Claude Code, Codex, opencode, or pi prompt as context. Select code, type the prompt, hit enter.
Everything else stays the same. You keep using nvim like normal.
cargo install sidekick
sidekick init
sidekick init walks through every AI harness it finds on your machine, wires the hooks, sets up the nvim alias, and skips anything already done. Re-run any time — it's idempotent.
Verify with sidekick doctor . Want to wire things by hand? See docs/SETUP.md .
Harness
Edit deny
Buffer refresh
Selection injection
Claude Code
✓
✓
✓
Codex
✓
✓
✓
opencode
✓
✓
✓
pi
✓
✓
✓
Crush
✓
—
—
Amp
✓
✓
—
Antigravity
✓
✓
—
Grok
✓
✓
—
A — is an upstream gap: the harness doesn't expose the hook event sidekick would need.
Just use nvim . The shell alias routes through sidekick so the hook can find your editor.
nvim src/main.rs
If you do not want the alias, run sidekick neovim <args> directly. You won't notice anything different until an agent tries to overwrite unsaved work, at which point sidekick denies the edit and tells the agent why.
Command
What it does
sidekick neovim <args>
Launches Neovim with a per-directory socket the hook can find. Aliased as nvim .
sidekick init
Guided setup. Wires every detected AI harness and the nvim alias.
sidekick doctor [--fix] [--no-color]
Health check. --fix offers consent-gated repairs and shows the diff before writing.
sidekick stats [--range week|month|year|all]
Local activity dashboard from append-only JSONL events. Nothing leaves your machine.
sidekick demo
Plays the demo cast inline in a ratatui frame. Useful for showing a coworker.
sidekick hook
Internal. Your AI harness calls it; you don't.
How it works
sidekick neovim launches nvim --listen /tmp/<blake3(cwd)>-<pid>.sock . The socket path is deterministic per canonical working directory and unique per process, so the hook can find every Neovim instance opened from the same project.
Before file edits, the agent calls sidekick hook . The hook globs /tmp/<blake3(cwd)>-*.sock , connects to reachable instances over msgpack-rpc with a short timeout, and asks whether each target is active with unsaved changes. If yes, the edit is denied; otherwise allowed. If no Neovim socket is found, sidekick degrades to allow.
After an edit lands, the hook tells every reachable Neovim instance with the file open to reload it. Cursor positions and visible windows are preserved.
On prompt submission, if Neovim has a live visual selection or recent visual marks, sidekick returns fenced context blocks like [Selected from path:start-end] . Claude Code and Codex receive them as additional context; the opencode and pi bridges append them to the submitted prompt text.
Decisions, refreshes, Neovim launches, and stats views are appended locally to sidekick/events.jsonl under your OS data directory. Writes are best-effort and never block the hook path.
No daemons, no background service. Just one CLI, Neovim RPC sockets in /tmp , and optional per-tool bridge files.
Neovim with RPC + Lua support, a Unix-like system, and at least one supported AI harness: Claude Code, Codex, opencode, pi, Crush, Amp, Antigravity, or Grok. Rust/Cargo is required for cargo install .
Sidekick is Phase 1 — Neovim plus Claude Code, Codex, opencode, pi, Crush, Amp, Antigravity, and Grok. The longer arc is a small protocol any editor can expose and any AI tool can query before writing: is this file being edited by a human right now? PHILOSOPHY.md has the roadmap (Helix, Zed, VS Code; Aider, Goose, Continue) and the extension points.
Issues and PRs welcome. New editor support starts at the Editor trait in src/editor.rs ; new AI harness support starts at the Harness trait in src/harness.rs . PHILOSOPHY.md covers the architecture.
There was an error while loading. Please reload this page .
2
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
