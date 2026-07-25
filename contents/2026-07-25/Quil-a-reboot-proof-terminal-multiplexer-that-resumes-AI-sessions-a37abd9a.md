---
source: "https://github.com/artyomsv/quil"
hn_url: "https://news.ycombinator.com/item?id=49047032"
title: "Quil – a reboot-proof terminal multiplexer that resumes AI sessions"
article_title: "GitHub - artyomsv/quil: Reboot-proof terminal multiplexer for AI-native devs — a tmux alternative that persists your whole workspace across reboots and auto-resumes Claude Code & OpenCode sessions. Ships an MCP server so AI agents can drive your panes. Go, cross-platform (Linux/macOS/Windows). MIT.\n[truncated]"
author: "artyomsv"
captured_at: "2026-07-25T13:01:12Z"
capture_tool: "hn-digest"
hn_id: 49047032
score: 1
comments: 0
posted_at: "2026-07-25T12:27:38Z"
tags:
  - hacker-news
  - translated
---

# Quil – a reboot-proof terminal multiplexer that resumes AI sessions

- HN: [49047032](https://news.ycombinator.com/item?id=49047032)
- Source: [github.com](https://github.com/artyomsv/quil)
- Score: 1
- Comments: 0
- Posted: 2026-07-25T12:27:38Z

## Translation

タイトル: Quil – AI セッションを再開する再起動防止型ターミナル マルチプレクサー
記事のタイトル: GitHub - artyomsv/quil: AI ネイティブ開発者向けの再起動防止ターミナル マルチプレクサー — 再起動後もワークスペース全体を保持し、Claude Code と OpenCode セッションを自動再開する tmux の代替品。 AI エージェントがペインを操作できるように MCP サーバーを同梱します。クロスプラットフォーム (Linux/macOS/Windows) にアクセスしてください。マサチューセッツ工科大学。
[切り捨てられた]
説明: AI ネイティブ開発者向けの再起動耐性のあるターミナル マルチプレクサー — 再起動後もワークスペース全体を保持し、Claude Code および OpenCode セッションを自動再開する tmux の代替手段です。 AI エージェントがペインを操作できるように MCP サーバーを同梱します。クロスプラットフォーム (Linux/macOS/Windows) にアクセスしてください。マサチューセッツ工科大学。 - artyomsv/quil

記事本文:
GitHub - artyomsv/quil: AI ネイティブ開発者向けの再起動防止ターミナル マルチプレクサー — 再起動後もワークスペース全体を保持し、Claude Code と OpenCode セッションを自動再開する tmux の代替品。 AI エージェントがペインを操作できるように MCP サーバーを同梱します。クロスプラットフォーム (Linux/macOS/Windows) にアクセスしてください。マサチューセッツ工科大学。 · GitHub
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
アルチョムスフ
/
羽根ペン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
237 コミット 237 コミット .claude .claude .github .github cmd cmd docs docs 内部 内部スクリプト スクリプト サイト サイト techdebt techdebt tools/ imgproc tools/ imgproc winres winres .dockerignore .dockerignore .editorconfig .editorconfig .gitignore .gitignore .goreleaser.yml .goreleaser.yml CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md THIRD_PARTY_LICENSES.md THIRD_PARTY_LICENSES.md VERSION VERSION go.mod go.mod go.sum go.sum package-lock.json package-lock.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI ネイティブ開発のための永続的なワークフロー オーケストレーター。
AI アシスタント、ビルド ウォッチャー、Webhook トンネル、SSH 接続にわたってプロジェクトごとに 5 ～ 10 のセッションを調整する開発者向けに構築されたターミナル マルチプレクサー。 tmux とは異なり、Quil はプロジェクトと型指定されたペインを理解します。再起動後もワークスペース全体を保持し、セッション ID によって AI 会話を自動再開し、AI アシスタントが MCP 経由で端末を駆動できるようにします。
再起動後に「quil」と入力します。すべてのタブ、ペイン、作業ディレクトリ、レイアウトの分割、および AI の会話が、保存した場所に表示されます。
完全な再起動後も存続します
AI が端末を駆動する
ペイン、作業ディレクトリ、AI セッションは 30 秒以内にスナップバックします。
MCP 上で Quil を公開する — エージェントはペインをリストし、出力を読み取り、キーを送信します。
多くのプロジェクトを 1 つのウィンドウで
型付きペイン
フォーカス モード + 12 個のプロジェクト タブ。
タイプごとのセットアップ: ブラウザーのディレクトリ、戦略の再開、権限の切り替え。
mでサイズを変更します

使う
右クリックペインメニュー
任意の分割境界線をドラッグします。ネストされたペインは最小値に固定され、PTY ではリリース時に 1 つのサイズが変更されます。
カーソル下のペインごとのアクション: 履歴、メモ、lazygit、アテンション ピン、再起動、閉じる。
コマンドパレット:どこでもジャンプ
…そして何でも実行してください
Alt+Shift+P では、すべてのペインとタブをあいまい検索します。ナビゲーションは上部にグループ化されます。
すべてのアクションは以下にグループ化されており、それぞれのキーバインドが示されています。と入力すると即座にフィルタリングされます。
...すべてのペイン内を検索します
試合にジャンプ
入力を開始すると、パレットは読み込まれたすべてのペインのスクロールバックも検索します。一致数と [ペインで見つかった] の下のプレビューが表示されます。
ペインの一致を入力すると、そのペインに直接ジャンプします。ロードされたペインを検索します。遅延復元ペインは、開くと表示されます。
インストール
Linux / macOS — 1 行インストール (OS+arch を検出、SHA-256 を検証):
カール -sSfL https://raw.githubusercontent.com/artyomsv/quil/master/scripts/install.sh |しー
Windows — Releases から quil-windows-amd64.zip をダウンロードし、 PATH 上の任意の場所に解凍します。
github.com/artyomsv/quil/cmd/quil@latest をインストールしてください
github.com/artyomsv/quil/cmd/quild@latest をインストールしてください
フル インストール オプション + ソースからのビルド — docs/installation.md を参照してください。
quil # TUI を起動し、デーモンを自動起動します
覚えておくべき 5 つの鍵:
始めるにはそれだけで十分です。初回起動時のウォークスルーについては docs/quick-start.md を、完全なキーマップについては docs/keybindings.md を参照してください。
何かがハングした場合は、quil restart によってデーモンが回復され (停止→新規開始→最後のスナップショットからタブが復元されます)、Alt+R によって、AI セッションが再開された単一のスタック ペインが所定の位置で再起動されます。
AI アシスタントに Quil を運転させましょう
これを AI クライアントの MCP 構成 (Claude Desktop、Claude Code、Cursor、VS Code Copilot) に追加します。
{
"mcpサーバー": {
"クイル" : {
"コマンド" : "キル" ,
"args" : [ "mcp " ]
}
}
}
クライアントを再起動します。 AI がリストを作成できるようになりました

t_panes 、 read_pane_output 、 send_to_pane 、 watch_notifications 、 Screenshot_pane 、およびその他 12 のツール。ビルド ペインを読み、コピー＆ペーストせずにエラーに対応します。
入力されたペインは、開発者が一日中実行するツール向けに出荷されています。それぞれ Ctrl+N から開きます。外部バイナリをラップするものは、そのバイナリが PATH 上にある場合にのみ表示されます (それ以外の場合は、インストール リンクがグレー表示されます)。
TOML で独自のペイン タイプを定義します。プラグイン リファレンスを参照してください。
トピック
ドクター
インストール
インストール.md
最初の打ち上げ
クイックスタート.md
すべての機能
機能.md
キーバインド
キーバインド.md
構成
構成.md
MCP（AI統合）
mcp.md
カスタムプラグイン
プラグインリファレンス.md
トラブルシューティング
トラブルシューティング.md
アーキテクチャ (24 の ADR)
アーキテクチャ.md
ロードマップ
ロードマップ.md
完全なドキュメント インデックスは docs/README.md にあります。
ブランチ/コミットの規則と開発ワークフローについては、CONTRIBUTING.md を参照してください。バグ報告やPRは大歓迎です。
MIT — Copyright (c) 2026 Artjoms Stukans
Windows ビルドには、Windows 10 でターミナル ペインを正しくホストするための Microsoft の MIT ライセンス OpenConsole ( OpenConsole.exe + conpty.dll ) がバンドルされています。完全なサードパーティの帰属については、THIRD_PARTY_LICENSES.md を参照してください。
AI ネイティブ開発者向けの再起動耐性のあるターミナル マルチプレクサー — 再起動後もワークスペース全体を保持し、Claude Code と OpenCode セッションを自動再開する tmux の代替手段です。 AI エージェントがペインを操作できるように MCP サーバーを同梱します。クロスプラットフォーム (Linux/macOS/Windows) にアクセスしてください。マサチューセッツ工科大学。
Readme MIT ライセンス
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Reboot-proof terminal multiplexer for AI-native devs — a tmux alternative that persists your whole workspace across reboots and auto-resumes Claude Code & OpenCode sessions. Ships an MCP server so AI agents can drive your panes. Go, cross-platform (Linux/macOS/Windows). MIT. - artyomsv/quil

GitHub - artyomsv/quil: Reboot-proof terminal multiplexer for AI-native devs — a tmux alternative that persists your whole workspace across reboots and auto-resumes Claude Code & OpenCode sessions. Ships an MCP server so AI agents can drive your panes. Go, cross-platform (Linux/macOS/Windows). MIT. · GitHub
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
artyomsv
/
quil
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
237 Commits 237 Commits .claude .claude .github .github cmd cmd docs docs internal internal scripts scripts site site techdebt techdebt tools/ imgproc tools/ imgproc winres winres .dockerignore .dockerignore .editorconfig .editorconfig .gitignore .gitignore .goreleaser.yml .goreleaser.yml CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md THIRD_PARTY_LICENSES.md THIRD_PARTY_LICENSES.md VERSION VERSION go.mod go.mod go.sum go.sum package-lock.json package-lock.json View all files Repository files navigation
The persistent workflow orchestrator for AI-native development.
A terminal multiplexer built for developers who orchestrate 5–10 sessions per project across AI assistants, build watchers, webhook tunnels, and SSH connections. Unlike tmux, Quil understands projects and typed panes : it persists your entire workspace across reboots, auto-resumes AI conversations by session id, and lets your AI assistant drive your terminal over MCP .
Type quil after a reboot — every tab, pane, working directory, layout split, and AI conversation is right where you left it.
Survives a full reboot
AI drives your terminal
Panes, working dirs, and AI sessions snap back in ~30s.
Expose Quil over MCP — agents list panes, read output, send keys.
Many projects, one window
Typed panes
Focus mode + a dozen project tabs.
Per-type setup: dir browser, resume strategy, permission toggles.
Resize with the mouse
Right-click pane menu
Drag any split border — nested panes clamp to minimums, PTYs see one resize on release.
Per-pane actions under the cursor: history, notes, lazygit, attention pin, restart, close.
Command palette: jump anywhere
…and run anything
Alt+Shift+P fuzzy-finds every pane and tab — navigation grouped at the top.
Every action grouped below, each showing its keybinding; type to filter instantly.
…and search inside every pane
Jump to the match
Start typing and the palette also searches every loaded pane's scrollback — match counts + a preview under Found in panes .
Enter on a pane match jumps straight to it. Searches loaded panes; lazily-restored panes appear once you open them.
Install
Linux / macOS — one-line install (detects OS+arch, verifies SHA-256):
curl -sSfL https://raw.githubusercontent.com/artyomsv/quil/master/scripts/install.sh | sh
Windows — download quil-windows-amd64.zip from Releases , extract anywhere on PATH .
go install github.com/artyomsv/quil/cmd/quil@latest
go install github.com/artyomsv/quil/cmd/quild@latest
Full install options + build-from-source — see docs/installation.md .
quil # launches the TUI, auto-starts the daemon
Five keys to remember:
That's enough to start. See docs/quick-start.md for the first-launch walkthrough and docs/keybindings.md for the full keymap.
If anything ever hangs: quil restart recovers the daemon (escalating stop → fresh start → tabs restored from the last snapshot), and Alt+R restarts a single stuck pane in place with its AI session resumed.
Let your AI assistant drive Quil
Add this to your AI client's MCP config (Claude Desktop, Claude Code, Cursor, VS Code Copilot):
{
"mcpServers" : {
"quil" : {
"command" : " quil " ,
"args" : [ " mcp " ]
}
}
}
Restart the client. The AI can now list_panes , read_pane_output , send_to_pane , watch_notifications , screenshot_pane , and 12 more tools. Read the build pane and react to errors without copy-paste.
Typed panes ship for the tools developers run all day. Each opens from Ctrl+N ; the ones that wrap an external binary appear only when that binary is on PATH (greyed with an install link otherwise).
Define your own pane types in TOML — see the plugin reference .
Topic
Doc
Installation
installation.md
First launch
quick-start.md
All features
features.md
Keybindings
keybindings.md
Configuration
configuration.md
MCP (AI integration)
mcp.md
Custom plugins
plugin-reference.md
Troubleshooting
troubleshooting.md
Architecture (24 ADRs)
architecture.md
Roadmap
roadmap.md
The full doc index lives at docs/README.md .
See CONTRIBUTING.md for branch / commit conventions and the development workflow. Bug reports and PRs welcome.
MIT — Copyright (c) 2026 Artjoms Stukans
The Windows build bundles Microsoft's MIT-licensed OpenConsole ( OpenConsole.exe + conpty.dll ) to host terminal panes correctly on Windows 10. See THIRD_PARTY_LICENSES.md for full third-party attribution.
Reboot-proof terminal multiplexer for AI-native devs — a tmux alternative that persists your whole workspace across reboots and auto-resumes Claude Code & OpenCode sessions. Ships an MCP server so AI agents can drive your panes. Go, cross-platform (Linux/macOS/Windows). MIT.
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
