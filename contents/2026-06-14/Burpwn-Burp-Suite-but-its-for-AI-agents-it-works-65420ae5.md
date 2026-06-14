---
source: "https://github.com/own2pwn-fr/burpwn"
hn_url: "https://news.ycombinator.com/item?id=48530385"
title: "Burpwn – Burp Suite but its for AI agents (it works)"
article_title: "GitHub - own2pwn-fr/burpwn: Transparent intercepting proxy + sandbox + agent interface for AI-driven web pentesting (Burp, but for AI agents) · GitHub"
author: "own2pwn-fr"
captured_at: "2026-06-14T18:35:56Z"
capture_tool: "hn-digest"
hn_id: 48530385
score: 3
comments: 0
posted_at: "2026-06-14T17:54:29Z"
tags:
  - hacker-news
  - translated
---

# Burpwn – Burp Suite but its for AI agents (it works)

- HN: [48530385](https://news.ycombinator.com/item?id=48530385)
- Source: [github.com](https://github.com/own2pwn-fr/burpwn)
- Score: 3
- Comments: 0
- Posted: 2026-06-14T17:54:29Z

## Translation

タイトル: Burpwn – Burp Suite ですが、AI エージェント用です (機能します)
記事タイトル: GitHub - own2pwn-fr/burpwn: AI 駆動型 Web 侵入テスト用の透過的インターセプト プロキシ + サンドボックス + エージェント インターフェイス (Burp、ただし AI エージェント用) · GitHub
説明: AI 主導の Web 侵入テスト用の透過的インターセプト プロキシ + サンドボックス + エージェント インターフェイス (Burp、ただし AI エージェント用) - own2pwn-fr/burpwn

記事本文:
GitHub - own2pwn-fr/burpwn: AI 主導の Web 侵入テスト用の透過的インターセプト プロキシ + サンドボックス + エージェント インターフェイス (Burp、ただし AI エージェント用) · GitHub
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
own2pwn-fr
/
げっぷ
公共
通知
通知セットを変更するにはサインインする必要があります

ひずみ
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
19 コミット 19 コミット .github/ workflows .github/ workflows crates crates skill/ burpwn skill/ burpwn training training .gitignore .gitignore CHANGELOG.md CHANGELOG.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE Makefile Makefile README.md README.md install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI 主導の Web 侵入テストのための、透過的なインターセプト プロキシ + 実行サンドボックス + エージェント インターフェイス。
AI エージェントにとっての burpwn は、人間のペンテスターに​​とっての Burp Suite のようなものです。エージェントがすべてのコマンドを実行します
ネットワーク全体 (HTTP/HTTPS/DNS/TCP) が強制的にアクセスされるルートレス Linux サンドボックス内で実行されます。
組み込みのインターセプト プロキシを介して。エージェントは履歴を遡って検索し、フィルタリングすることができます。
復号化されたリクエスト/レスポンス フロー、それらの再生と編集 (リピーター)、一致/置換ルールの適用、
実行中のトラフィックをブロックして書き換え、フローをワークスペースに整理します。すべてスクリプト可能な CLI から実行します。
またはMCP上で。それはげっぷであると同時にサメでもありますが、エージェントによって運転されています。
ステータス: 開発初期。以下のマイルストーンを参照してください。
既存のインターセプト プロキシは、人間が GUI をクリックするために構築されています。自律エージェントには、
プログラム サーフェス: セッションを作成し、ツールを実行し、キャプチャされたトラフィックをクエリします。
エージェント自身の LLM トラフィックが常にキャプチャされます。 burpwn はまさにそれを実現します。エージェント プロセスはそのまま残ります。
サンドボックスの外。実行するコマンド (その子) のみがキャプチャされたネットワークに入ります。
したがって、LLM トラフィックは構築によって除外されます。
根のない透明な砂場。実行された各コマンドは、独自の Linux ユーザー + ネットワークで実行されます。
名前空間。その名前空間内の nftables REDIRECT ルールセットは、すべての TCP (および UDP/53) を強制的に
b

urpwnプロキシ。 bubblewrap はファイルシステムとプロセスを分離します。 root、setuid、CAP_NET_ADMIN なし
ホスト上 — カーネルは、子名前空間内で必要な機能を付与します。
TLS-MITM。インストールごとのルート CA は 1 回生成されます。リーフ証明書は SNI に従ってオンザフライで鋳造され、
CA はサンドボックス トラスト ストアに挿入されるため、HTTPS が復号化されます。証明書に固定されたターゲットはフォールバックします
メタデータのみのログを使用して TLS パススルーにクリーンに移行します。
キャプチャとクエリ。フローはセッションごとの SQLite データベース (WAL、コンテンツ アドレス指定された本文) に保存されます。
dedup、FTS5 全文検索）は、プロキシ ホット パスから離れたシングル ライター タスクによって書き込まれます。
エージェントの統合 (rtk スタイル)。 burpwn init は、
検出されたエージェント (Claude Code / Copilot、Cursor、Gemini CLI、Cline/Roo)、および汎用グローバル シェル
フックなのでカスタムエージェントもカバーされます。
burpwn Doctor # 根無し草の前提条件を確認する
burpwn ca init && burpwn ca export # MITM CA を生成/印刷します
burpwnセッション新しい --nameengagement-1
burpwn exec --curl -s https://target.example/ # サンドボックスで実行します。キャプチャされたトラフィックと復号化されたトラフィック
burpwn req list # キャプチャされたフローを参照する
burpwn req show 42 --raw # 復号化されたリクエスト + レスポンス
burpwn req recruit 42 --set-header ' X: 1 ' # リピーター
burpwn インターセプトを有効にする # インターセプトをブロックする (MCP await_intercept 経由も)
インストール
Linux のみ (ユーザー/ネットワーク名前空間、nftables、bubblewrap に依存)。前提条件をインストールする
最初 — Fedora/RHEL: sudo dnf install bubblewrap nftables iproute ; Debian/Ubuntu:
sudo apt install bubblewrap nftables iproute2 。
# ワンライナー: ビルド済みバイナリをダウンロードし、~/.local/bin にインストールし、CA を生成し、プリフライトを実行します
カール -fsSL https://raw.githubusercontent.com/own2pwn-fr/burpwn/main/install.sh |しー
# チェックアウトから (事前にビルドされた BINA がない場合はソースからビルドします)

ryはあなたの土踏まずにフィットします）
./install.sh # ./install.sh --hooks はグローバル シェル フックもインストールします
./install.sh --from-source # ソースビルドを強制します
# またはカーゴ / Makefile 経由
カーゴインストール --git https://github.com/own2pwn-fr/burpwn burpwn
make install # PREFIX=/usr/local make install (sudo が必要な場合があります); 「make help」はタスクをリストします
カール | sh パスは、アーキテクチャ (x86_64 / aarch64 Linux) のリリース バイナリをダウンロードし、
チェックサムを検証します。どれも一致しない場合は、カーゴ ソース ビルドにフォールバックします。
カーゴビルド --release # target/release/burpwn に単一の `burpwn` バイナリを生成します
カーゴテスト # 特権ルートレスサンドボックステストは #[ignore]d です
エージェントの統合
burpwn init は rtk スタイルのコマンド書き換えフックをインストールするため、エージェントが実行するすべてのシェル コマンドは
エージェント自身の LLM トラフィックが burpwn exec (キャプチャおよび復号化) を介して透過的にルーティングされる一方で、
決して触れられません。 MCP サーバーと既製のエージェント スキルもあります。
burpwn init --agent claude # クロード コード / Copilot PreToolUse フック (カーソル、ジェミニ、クラインとも)
burpwn init --global # 汎用シェルフック — どのエージェントでも機能します
burpwn mcp # stdio 経由の MCP サーバー (セッション/実行/要求/インターセプト ツール)
バンドルされたエージェント スキルは skill/burpwn/ にあります。これを次の場所にコピーします。
~/.claude/skills/ (またはエージェントのスキル ディレクトリ) を参照して、エージェントにワークフローを教えます。
AI 主導の Web 侵入テスト用の透過的インターセプト プロキシ + サンドボックス + エージェント インターフェイス (Burp、ただし AI エージェント用)
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
2
v0.1.1
最新の
2026 年 6 月 13 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

の上

## Original Extract

Transparent intercepting proxy + sandbox + agent interface for AI-driven web pentesting (Burp, but for AI agents) - own2pwn-fr/burpwn

GitHub - own2pwn-fr/burpwn: Transparent intercepting proxy + sandbox + agent interface for AI-driven web pentesting (Burp, but for AI agents) · GitHub
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
own2pwn-fr
/
burpwn
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
19 Commits 19 Commits .github/ workflows .github/ workflows crates crates skills/ burpwn skills/ burpwn training training .gitignore .gitignore CHANGELOG.md CHANGELOG.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE Makefile Makefile README.md README.md install.sh install.sh View all files Repository files navigation
A transparent intercepting proxy + execution sandbox + agent interface for AI-driven web pentesting.
burpwn is to an AI agent what Burp Suite is to a human pentester. It runs every command an agent
executes inside a rootless Linux sandbox whose entire network (HTTP/HTTPS/DNS/TCP) is forced
through a built-in intercepting proxy. The agent can then go back through history, search and filter
the decrypted request/response flows, replay and edit them (Repeater), apply match/replace rules,
block and rewrite traffic in flight, and organize flows into workspaces — all from a scriptable CLI
or over MCP. It is at once a Burp and a tshark, but driven by an agent.
Status: early development. See the milestones below.
Existing intercepting proxies are built for a human clicking in a GUI. An autonomous agent needs a
programmatic surface: create a session, run tooling, and query the captured traffic — without the
agent's own LLM traffic ever being captured. burpwn delivers exactly that: the agent process stays
outside the sandbox; only the commands it executes (its children) enter the captured network
namespace, so LLM traffic is excluded by construction.
Rootless transparent sandbox. Each executed command runs in its own Linux user + network
namespace. An nftables REDIRECT ruleset inside that namespace forces all TCP (and UDP/53) to the
burpwn proxy. bubblewrap isolates the filesystem and processes. No root, no setuid, no CAP_NET_ADMIN
on the host — the kernel grants the needed capability inside the child namespace.
TLS-MITM. A per-install root CA is generated once; leaf certs are minted on the fly per SNI and
the CA is injected into the sandbox trust store so HTTPS is decrypted. Cert-pinned targets fall back
cleanly to TLS pass-through with metadata-only logging.
Capture & query. Flows are stored in a per-session SQLite database (WAL, content-addressed body
dedup, FTS5 full-text search) written by a single-writer task off the proxy hot path.
Agent integration (rtk-style). burpwn init installs the right command-rewrite hook for the
detected agent (Claude Code / Copilot, Cursor, Gemini CLI, Cline/Roo), plus a generic global shell
hook so even a custom agent is covered.
burpwn doctor # check the rootless prerequisites
burpwn ca init && burpwn ca export # generate / print the MITM CA
burpwn session new --name engagement-1
burpwn exec -- curl -s https://target.example/ # runs sandboxed; traffic captured + decrypted
burpwn req list # browse captured flows
burpwn req show 42 --raw # decrypted request + response
burpwn req replay 42 --set-header ' X: 1 ' # Repeater
burpwn intercept enable # blocking intercept (also via MCP await_intercept)
Install
Linux-only (relies on user/network namespaces, nftables, bubblewrap). Install the prerequisites
first — Fedora/RHEL: sudo dnf install bubblewrap nftables iproute ; Debian/Ubuntu:
sudo apt install bubblewrap nftables iproute2 .
# one-liner: download the prebuilt binary, install to ~/.local/bin, generate the CA, run preflight
curl -fsSL https://raw.githubusercontent.com/own2pwn-fr/burpwn/main/install.sh | sh
# from a checkout (builds from source if no prebuilt binary fits your arch)
./install.sh # ./install.sh --hooks also installs the global shell hook
./install.sh --from-source # force a source build
# or via cargo / the Makefile
cargo install --git https://github.com/own2pwn-fr/burpwn burpwn
make install # PREFIX=/usr/local make install (may need sudo); `make help` lists tasks
The curl | sh path downloads the release binary for your architecture (x86_64 / aarch64 Linux) and
verifies its checksum; if none matches it falls back to a cargo source build.
cargo build --release # produces a single `burpwn` binary at target/release/burpwn
cargo test # the privileged rootless-sandbox test is #[ignore]d
Agent integration
burpwn init installs an rtk-style command-rewrite hook so every shell command your agent runs is
transparently routed through burpwn exec (captured + decrypted), while the agent's own LLM traffic
is never touched. There is also an MCP server and a ready-made agent skill:
burpwn init --agent claude # Claude Code / Copilot PreToolUse hook (also: cursor, gemini, cline)
burpwn init --global # generic shell hook — works for any agent
burpwn mcp # MCP server over stdio (session/exec/req/intercept tools)
The bundled agent skill lives in skills/burpwn/ — copy it into
~/.claude/skills/ (or your agent's skills dir) to teach an agent the workflow.
Transparent intercepting proxy + sandbox + agent interface for AI-driven web pentesting (Burp, but for AI agents)
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
2
v0.1.1
Latest
Jun 13, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
