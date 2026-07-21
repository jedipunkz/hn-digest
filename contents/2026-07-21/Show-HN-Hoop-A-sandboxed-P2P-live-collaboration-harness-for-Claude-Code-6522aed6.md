---
source: "https://github.com/bruno-de-queiroz/hoop"
hn_url: "https://news.ycombinator.com/item?id=48999928"
title: "Show HN: Hoop – A sandboxed P2P live collaboration harness for Claude Code"
article_title: "GitHub - bruno-de-queiroz/hoop: Being alone is not a requirement. A sandboxed and web harness on top of claude code. · GitHub"
author: "brunodequeiroz"
captured_at: "2026-07-21T23:48:29Z"
capture_tool: "hn-digest"
hn_id: 48999928
score: 1
comments: 0
posted_at: "2026-07-21T23:46:06Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Hoop – A sandboxed P2P live collaboration harness for Claude Code

- HN: [48999928](https://news.ycombinator.com/item?id=48999928)
- Source: [github.com](https://github.com/bruno-de-queiroz/hoop)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T23:46:06Z

## Translation

タイトル: Show HN: Hoop – クロード コード用のサンドボックス化された P2P ライブ コラボレーション ハーネス
記事のタイトル: GitHub - bruno-de-queiroz/hoop: 一人でいることは必須ではありません。クロード コード上のサンドボックス化された Web ハーネス。 · GitHub
説明: 一人でいる必要はありません。クロード コード上のサンドボックス化された Web ハーネス。 - ブルーノ・デ・ケイロス/フープ

記事本文:
GitHub - bruno-de-queiroz/hoop: 一人でいることは必須ではありません。クロード コード上のサンドボックス化された Web ハーネス。 · GitHub
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
ブルーノ・デ・ケイロス
/
フープ
公共
通知
変更するにはサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
48 コミット 48 コミット .claude-plugin .claude-plugin .github/ workflows .github/ workflows .impeccable .impeccable docs docs plugins/ hoop plugins/ hoop .gitignore .gitignore DESIGN.md DESIGN.md LICENSE LICENSE README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
一人でいることは必須条件ではありません。 Claude Code 上のサンドボックス化された Web ハーネス。 1 回のインストールで 2 つの成果物:
/hoop:setup — 厳選されたスタック (メモリ、コードグラフ RAG、自動化、プラットフォーム MCP、ドキュメント RAG、可観測性、設計、セカンド ブレイン) を接続する対話型ウィザード。
/hoop:dashboard — ライブ セッションを備えたコンテナ化されたローカル Web ダッシュボード (Next.js、 http://localhost:7842/ )、ワンクリック トリガーを備えたスキル ブラウザ、ネストされたサブエージェント ツリー、プッシュ ベースのイベント オブザーバビリティ、およびすべてのイベントにわたる BM25 + オプションのセマンティック検索。
このプラグインは、インストールされるサードパーティの MCP やスキルを再実装しません。それらを選択、インストール、文書化、観察します。
2 つの方法があります。ホストに一致する方を選択してください。
A · クロード コード内 (クロード コードはすでに実行されています)
/プラグイン マーケットプレイス ブルーノ・デ・ケイロス/フープを追加
/プラグインインストール hoop@hoop-marketplace
/プラグインリスト
/reload-プラグイン
/フープ:セットアップ
(Claude Code v2.1.138 では、現在のセッションで新たに事前シードされたプラグインをアクティブ化するために、/plugin list + /reload-plugins dance が必要です。新しいセッションでは必要ありません。)
B · スタンドアロン — ホスト上にクロード コードがない
クロード コード (またはノード) をホストにインストールする必要はありません。サンドボックスは独自のクロード バイナリを同梱し、独自のクロード アカウントで認証するため、Docker を備えたベア マシンで十分です。
Docker — Docker Desktop、または任意の Docker エンジン + Compose v2

( docker compose )。これはフープ スタートの唯一のハード依存関係です。クロード コード、ノード、gh、uv、gws、およびダッシュボードはすべてコンテナー内に存在し、ホスト上には存在しません。また、サンドボックスはブート時に独自のクロード プロファイルをシードします (そのため、フープ スタートにはホスト jq が必要ありません)。他のいくつかのサブコマンドは、小規模なホスト ツールをシェルアウトします。以下のオプションのリストを参照してください。
git と bash — リポジトリのクローンを作成し、CLI を実行します。
アクセスできる Web ブラウザ - ワンタイム ログイン承認用。同じマシン上にある必要はありません。 URL を開いてコードをコピーし直すだけです (ステップ 4 を参照)。ヘッドレスサーバーに便利です。
/login 経由でサインインできるアクティブなプラン (Pro / Max / Team / Enterprise) を持つクロード アカウント。
jq (hoop start ではなく、一部のサブコマンドで必要) — ホスト上の JSON を編集する hoop setup および hoop logout で必要です (ウィザードは gh/Claude ID を profile.md に読み取ります。logout はサンドボックスの資格情報を書き換えます)。 hoop open もこれを使用しますが、それが存在しない場合は警告が表示され機能が低下します。インストール: macOS brew install jq ; Debian/Ubuntu sudo apt-get install -y jq 。
curl (オプション - 厳密には必須ではありません) - Docker Model Runner を調査するためのフープ セットアップと、ダッシュボードの health-wait によって使用されます。どちらも、bash /dev/tcp プローブが見つからない場合は、それにフォールバックします。インストール: brew installcurl / sudo apt-get install -ycurl 。
awk (フープ マウントにのみ必要) — マウント リストを書き換えるためにフープ マウントの追加/削除で必要です。他のすべてのコマンドでは使用されません。
Docker Model Runner (オプション) — セマンティック検索埋め込み用。フープ セットアップによりそれが有効になり (Docker Desktop AI、または Docker Engine の docker-model-plugin)、Compose の models: 要素 (Docker Compose v2.38+) を介して埋め込みモデルが Compose スタックに追加されるため、フープ スタートがそれを自動プルし、エンドポイントをサンドボックスに接続します。 DMR が到達不能になった場合

e — または Compose が v2.38 より古い — フープ スタートは透過的にオーバーライドを削除し、失敗するのではなく BM25 にフォールバックします。スキップ可能 — Ollama、OpenAI、またはカスタムの OpenAI 互換エンドポイントも機能します。
いつでも hoop Doctor を実行して、ホストにすべての機能があることを確認し、クリーンな Docker のみのセットアップにどの程度近づいているかを確認します。
# 1. クローンを作成してリポジトリに入れる
git clone https://github.com/bruno-de-queiroz/hoop && cd hoop
# 2. (オプション) PATH + シェル補完に `hoop` を置きます。
# これをスキップし、代わりにすべてのコマンドの前に ./plugins/hoop/cli/hoop.sh を付けます。
./plugins/hoop/cli/hoop.sh install # その後、`hoop` が解決されるように新しいシェルを開きます
# 3. スタック全体を起動します。最初の実行では両方のイメージがビルドされます (数分)。
# 後の起動は速いです。 http://localhost:7842/ でダッシュボードを提供します。
フープスタート
# 4. (オプション) スタックを対話的に構成します — メモリ、コードグラフ RAG、
# プラットフォーム MCP、ドキュメント RAG、セマンティック検索、可観測性、セカンド ブレインなど。
# これは、CLI にネイティブな /hoop:setup ウィザードです (ホストのクロード コードは必要ありません)。
フープのセットアップ
# 5. 1 回限り: 独自の Claude アカウントを使用してサンドボックスを認証します。
# サンドボックス内の「claude」にドロップします — /login と入力し、印刷されたファイルを開きます
# ブラウザで URL を承認し、コードを貼り付けて戻し、/exit します。
フープログイン
次に、 http://localhost:7842/ を開きます。ダッシュボードは localhost のみを信頼するため、ブラウザーは自動的にホストとして認識されます。アクセス トークンは自動的に作成され、貼り付けるものはありません。 (リモート アクセスは意図的にゲートされています。「アーキテクチャ」を参照してください。)
これが完全なループです: フープ スタート → フープ セットアップ (オプション) → フープ ログイン → ダッシュボードを開きます。スタンドアロン パスに関するいくつかの注意事項:
フープセットアップを使用してスタックを構成します。これは対話型ウィザードです (メモリ、コードグラフ RAG、自動化、プラットフォーム MCP、ドキュメント RAG、セマンティック SE)

アーチ、可観測性、デザイン、セカンドブレイン、テレメトリー分離) — /hoop:setup の背後にあるものと同じですが、CLI にネイティブであるため、ホストのクロード コードは必要ありません。これはフープ ログインの前に実行されます (MCP 構成の書き込みには認証は必要ありません)。後で 1 回限りの追加を行う場合、フープ追加 mcp/plugin/skill (CLI の表を参照) は再構築後もサンドボックス プロファイルに残ります。
端末のみ？ hoop open は、現在のディレクトリ上の使い捨てサンドボックスでインタラクティブなクロードを実行します。ログイン後にダッシュボードやブラウザは必要ありません。
更新: git pull と hoop restart により、両方のイメージに新しいコードが取り込まれます。
サインアウト/アカウントの切り替え: フープ ログアウトはサンドボックスのクロード ログインをクリアし、別のアカウントでフープ ログインできるようにします。
hoop setup (または /hoop:setup を指定すると、その手順が示されます) では、最初に 1 つの同意を示してこれらの手順を実行し、各選択をインストールします。自動実行可能な MCP と秘密取得 MCP はすぐに実行されます (すべてのコマンドが最初に出力されます)。ブラウザーログイン、プラグインマーケットプレイス、およびホスト CLI オプションは、自分で完了するためのガイド付きステップとして出力されます。
各実行はサンドボックス プロファイルの ~/.claude/hoop/sandbox/profile/.claude/hoop/install-log.md (ダッシュボードからも表示可能) に追加されるため、再実行は監査可能です。シークレットはログに到達することはなく、 claude mcp add -e または 0600 ~/.claude/hoop/hoop.env に直接送信されます。
/hoop:dashboard (または start | stop | restart |再構築 | status | logs ) は、コンテナー内でダッシュボードを実行します。ホストに必要なのは Docker Desktop だけです。Node、npm install、Next.js ビルド汚染は必要ありません。各動詞はオプションのサービス ターゲット ( all (デフォルト) · サンドボックス · ダッシュボード ) を取ります。 start は遅延ビルド (イメージが見つからない場合のみ) ですが、rebuild は常に再ビルドします。
ペアリング (セッションを共同推進するようにチームメイトを招待) では、cloudflared を使用してローカル ダッシュボードをパブリック トンネル経由で公開します。これはベイク処理された int です。

o ダッシュボード イメージなので、ホストには何もインストールする必要はありません。ダッシュボードはペアリングしなくても正常に動作します。共有リンクのみがトンネルを開始します。
セッション — ~/.claude/sessions/ の fs.watch ;リアルタイムで更新します。
スキル — ディスク上のすべてのスキル (ユーザー + プラグイン)。フィルタリング可能。ワンクリックの「実行」で、ダッシュボード コンテナー内に claude -p '/<name>' が生成され、stdout がパネルにストリーミングされます。
サブエージェント — エージェント ツールの PreToolUse / PostToolUse イベントから再構築されたネストされたツリー。ノードをクリックすると、そのプロンプト、ツール呼び出し、および最終出力が表示されます。
イベント — サーバー送信イベントによる時系列のライブテール。フックは、Unix ドメイン ソケットを介して各イベントをサンドボックスの /ingest にプッシュします。ダッシュボードは、ポーリングを行わずに結果のストリームを追跡します。
検索 — ⌘K で開きます。 BM25 (FTS5) は常に動作します。セマンティック検索 (sqlite-vec、768-dim 埋め込み) は、埋め込みバックエンドが構成されているときにアクティブになります。 hoop setup / /hoop:setup を介してセットアップします。代替として Ollama ( nomic-embed-text )、OpenAI、または任意のカスタム OpenAI 互換エンドポイントを使用して、Docker Model Runner (Compose の models: element を介して Compose スタックに追加され、 hoop start でプルされます) を推奨します。相互ランク融合によるハイブリッド融合 BM25 + セマンティック。
ダッシュボードは設計上、シングルユーザーおよびローカルホストのみです。アクセスはインストールごとのトークンによってゲートされます (下記のアーキテクチャを参照)。
plugins/hoop/cli/ には、ランタイムをラップする小さな oosh ベースの CLI が同梱されています。これはプラグイン (フレームワーク エンジン + エントリ ポイント + 補完 + lib/stack.sh のスタック エンジン) 内に存在するため、プラグインに付属しており、スラッシュ コマンド ( /hoop:setup 、 /hoop:dashboard ) で直接呼び出すことができます。外部インストールは必要なく、独自のパスを解決します。
./plugins/hoop/cli/hoop.sh install # シンボリックリンク `hoop` を PATH + シェル補完 (bash/zsh) に追加します
# またはその場で実行します

インストールする場合:
./plugins/hoop/cli/hoop.sh <モジュール> <コマンド>
2 つのレベル: 最上位の動詞はスタック全体に作用します。モジュールは単一のサービスをスコープとします。これらはすべて 1 つのエンジン ( cli/lib/stack.sh ) を駆動します。これは、ホスト側プリフライト (プロファイルのバインドマウント準備、ダッシュボード/ピア認証トークン、転送された埋め込み環境、オーケストレーションの構成) のための唯一の信頼できるソースです。クロードのオンボーディング バイパス + プラグイン/フック配線は、起動時にサンドボックス ( Sandbox/seed-profile.mjs ) 内で実行されるため、ホストには jq が必要ありません。 start/rebuild は意図的に分割されています。start はイメージが欠落している場合にのみイメージを構築します (そうでない場合は高速に) が、rebuild はコードの変更を取得できるように常に再構築します。
hoop start # http://localhost:7842/ でスタック全体を起動します
フープ セットアップ # インタラクティブ ウィザード: サンドボックス スタックを構成します (ホスト クロード コードは必要ありません)
hoop ログイン # ワンタイム: 独自のクロード アカウントでサンドボックスを認証します
hoop ダッシュボードの再構築 # ダッシュボード コンテナのみを再構築 + 再作成します
hoop サンドボックスの再構築 # エージェント サンドボックス コンテナのみを再構築 + 再作成する
$PWD 上のサンドボックスでフープを開く # インタラクティブ クロード
hoop add mcp context7 -- npx -y @upstash/context7-mcp # MCP (ユーザー スコープ) をサンドボックスにインストールします
hoop add skill -d ~ /.claude/skills/impeccable # ローカル スキルをサンドボックス プロファイルにコピーします
フープm

[切り捨てられた]

## Original Extract

Being alone is not a requirement. A sandboxed and web harness on top of claude code. - bruno-de-queiroz/hoop

GitHub - bruno-de-queiroz/hoop: Being alone is not a requirement. A sandboxed and web harness on top of claude code. · GitHub
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
bruno-de-queiroz
/
hoop
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
48 Commits 48 Commits .claude-plugin .claude-plugin .github/ workflows .github/ workflows .impeccable .impeccable docs docs plugins/ hoop plugins/ hoop .gitignore .gitignore DESIGN.md DESIGN.md LICENSE LICENSE README.md README.md View all files Repository files navigation
Being alone is not a requirement. A sandboxed and web harness on top of Claude Code. Two deliverables in one install:
/hoop:setup — an interactive wizard that wires a curated stack (memory, code-graph RAG, automation, platform MCPs, docs RAG, observability, design, second-brain).
/hoop:dashboard — a containerized local web dashboard (Next.js, http://localhost:7842/ ) with live sessions, a skill browser with one-click triggers, a nested sub-agent tree, push-based event observability, and BM25 + optional semantic search across all events.
The plugin does not re-implement any of the third-party MCPs or skills it installs. It picks, installs, documents, and observes them.
There are two ways in. Pick the one that matches your host.
A · Inside Claude Code (you already run Claude Code)
/plugin marketplace add bruno-de-queiroz/hoop
/plugin install hoop@hoop-marketplace
/plugin list
/reload-plugins
/hoop:setup
(The /plugin list + /reload-plugins dance is required on Claude Code v2.1.138 to activate a freshly pre-seeded plugin in the current session. New sessions don't need it.)
B · Standalone — no Claude Code on the host
You do not need Claude Code (or even Node) installed on the host. The sandbox ships its own claude binary and authenticates with its own Claude account, so a bare machine with Docker is enough.
Docker — Docker Desktop, or any Docker engine + Compose v2 ( docker compose ). This is the only hard dependency for hoop start — Claude Code, Node, gh, uv, gws, and the dashboard all live in containers, never on your host, and the sandbox seeds its own Claude profile on boot (so hoop start needs no host jq). A few other subcommands do shell out to small host tools — see the optional list below.
git and bash — to clone the repo and run the CLI.
A web browser you can reach — for the one-time login approval. It does not have to be on the same machine; you just need to open a URL and copy a code back (see step 4). Handy for headless servers.
A Claude account with an active plan (Pro / Max / Team / Enterprise) that can sign in via /login .
jq (needed by some subcommands, not hoop start ) — required by hoop setup and hoop logout , which edit JSON on the host (the wizard reads your gh/Claude identity into profile.md ; logout rewrites the sandbox credentials). hoop open uses it too but degrades with a warning if it's absent. Install: macOS brew install jq ; Debian/Ubuntu sudo apt-get install -y jq .
curl (optional — never strictly required) — used by hoop setup to probe Docker Model Runner and by the dashboard health-wait; both fall back to a bash /dev/tcp probe when it's missing. Install: brew install curl / sudo apt-get install -y curl .
awk (needed only for hoop mount ) — required by hoop mount add / remove to rewrite the mounts list; unused by every other command.
Docker Model Runner (optional) — for semantic-search embeddings. hoop setup enables it (Docker Desktop AI, or the docker-model-plugin on Docker Engine) and appends the embedding model to the compose stack via Compose's models: element ( Docker Compose v2.38+ ), so hoop start auto-pulls it and wires the endpoint into the sandbox. If DMR is ever unreachable — or Compose is older than v2.38 — hoop start transparently drops the override and falls back to BM25 rather than failing. Skippable — Ollama, OpenAI, or a custom OpenAI-compatible endpoint also work.
Run hoop doctor at any time to verify the host has everything and see how close you are to a clean Docker-only setup.
# 1. Clone and enter the repo
git clone https://github.com/bruno-de-queiroz/hoop && cd hoop
# 2. (optional) Put `hoop` on your PATH + shell completion.
# Skip this and prefix every command with ./plugins/hoop/cli/hoop.sh instead.
./plugins/hoop/cli/hoop.sh install # then open a new shell so `hoop` resolves
# 3. Bring up the whole stack. The FIRST run builds both images (a few minutes);
# later starts are fast. Serves the dashboard at http://localhost:7842/.
hoop start
# 4. (optional) Configure your stack interactively — memory, code-graph RAG,
# platform MCPs, docs RAG, semantic search, observability, second-brain, etc.
# This is the /hoop:setup wizard, native to the CLI (needs no host Claude Code).
hoop setup
# 5. One-time: authenticate the sandbox with its OWN Claude account.
# Drops you into `claude` inside the sandbox — type /login, open the printed
# URL in your browser, approve, paste the code back, then /exit.
hoop login
Then open http://localhost:7842/ . Because the dashboard only trusts localhost, your browser is recognized as the host automatically — the access token is minted for you, nothing to paste. (Remote access is deliberately gated; see Architecture .)
That's the full loop: hoop start → hoop setup (optional) → hoop login → open the dashboard. A few notes for the standalone path:
Configure your stack with hoop setup . This is the interactive wizard (memory, code-graph RAG, automation, platform MCPs, docs RAG, semantic search, observability, design, second-brain, telemetry isolation) — the same one behind /hoop:setup , but native to the CLI so it needs no host Claude Code. It runs before hoop login (writing MCP config needs no auth). For one-off additions later, hoop add mcp/plugin/skill (see the CLI table) persists into the sandbox profile across rebuilds.
Terminal-only? hoop open runs an interactive claude in a throwaway sandbox over your current directory — no dashboard, no browser needed after login.
Updating: git pull then hoop rebuild picks up new code into both images.
Signing out / switching accounts: hoop logout clears the sandbox's Claude login so a different account can hoop login .
hoop setup (or /hoop:setup , which points you to it) walks you through these steps with one consent at the top, then installs each pick — auto-runnable and secret-taking MCPs run immediately (every command printed first); browser-login, plugin-marketplace, and host-CLI options are printed as guided steps to finish yourself:
Each run appends to the sandbox profile's ~/.claude/hoop/sandbox/profile/.claude/hoop/install-log.md (also viewable from the dashboard) so re-runs are auditable. Secrets never reach the log — they go straight to claude mcp add -e or the 0600 ~/.claude/hoop/hoop.env .
/hoop:dashboard (or start | stop | restart | rebuild | status | logs ) runs the dashboard inside a container . Your host only needs Docker Desktop — no Node, no npm install , no Next.js build pollution. Each verb takes an optional service target ( all (default) · sandbox · dashboard ); start builds lazily (only when an image is missing) while rebuild always rebuilds.
Pairing (inviting a teammate to co-drive a session) uses cloudflared to expose the local dashboard over a public tunnel — it's baked into the dashboard image , so nothing to install on the host. The dashboard runs fine without pairing; only share links start a tunnel.
Sessions — fs.watch on ~/.claude/sessions/ ; updates in real time.
Skills — every skill on disk (user + plugin), filterable, with a one-click "Run" that spawns claude -p '/<name>' inside the dashboard container and streams stdout back to the panel.
Sub-agents — nested tree reconstructed from PreToolUse / PostToolUse events on the Agent tool. Click a node to see its prompt, tool calls, and final output.
Events — chronological live tail via Server-Sent Events. Hooks push each event to the sandbox's /ingest over the Unix domain socket; the dashboard tails the resulting stream with zero polling.
Search — opens with ⌘K. BM25 (FTS5) always works; semantic search (sqlite-vec, 768-dim embeddings) activates when an embedding backend is configured. Set one up via hoop setup / /hoop:setup — recommended is Docker Model Runner (added to the compose stack via Compose's models: element and pulled on hoop start ), with Ollama ( nomic-embed-text ), OpenAI, or any custom OpenAI-compatible endpoint as alternatives. Hybrid fuses BM25 + semantic via Reciprocal Rank Fusion.
The dashboard is single-user and localhost-only by design; access is gated by a per-install token (see Architecture below).
plugins/hoop/cli/ ships a small oosh -based CLI that wraps the runtime. It lives inside the plugin (framework engine + entry point + completions + the stack engine in lib/stack.sh ) so it ships with the plugin and the slash commands ( /hoop:setup , /hoop:dashboard ) can invoke it directly. It needs no external install and resolves its own paths.
./plugins/hoop/cli/hoop.sh install # symlink `hoop` onto PATH + shell completion (bash/zsh)
# or run in place without installing:
./plugins/hoop/cli/hoop.sh < module > < command >
Two levels: top-level verbs act on the whole stack; modules scope a single service. All of them drive one engine ( cli/lib/stack.sh ) — the single source of truth for host-side preflight (profile bind-mount prep, dashboard/peer auth tokens, forwarded embedding env, compose orchestration). The Claude onboarding bypass + plugin/hook wiring run inside the sandbox on boot ( sandbox/seed-profile.mjs ), so the host needs no jq. start / rebuild are deliberately split — start only builds an image when it's missing (fast otherwise), while rebuild always rebuilds so you pick up code changes.
hoop start # bring up the whole stack at http://localhost:7842/
hoop setup # interactive wizard: configure the sandbox stack (no host Claude Code needed)
hoop login # one-time: authenticate the sandbox with its own Claude account
hoop dashboard rebuild # rebuild + recreate ONLY the dashboard container
hoop sandbox rebuild # rebuild + recreate ONLY the agent-sandbox container
hoop open # interactive claude in a sandbox over $PWD
hoop add mcp context7 -- npx -y @upstash/context7-mcp # install an MCP (user scope) into the sandbox
hoop add skill -d ~ /.claude/skills/impeccable # copy a local skill into the sandbox profile
hoop m

[truncated]
