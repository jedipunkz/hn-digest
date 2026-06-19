---
source: "https://github.com/xyTom/coding-tools-mcp"
hn_url: "https://news.ycombinator.com/item?id=48594594"
title: "Show HN: Coding Tools MCP – Give any LLM agent the ability to code"
article_title: "GitHub - xyTom/coding-tools-mcp: Give any AI agent the ability to code · GitHub"
author: "xytom"
captured_at: "2026-06-19T04:36:40Z"
capture_tool: "hn-digest"
hn_id: 48594594
score: 1
comments: 0
posted_at: "2026-06-19T03:50:38Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Coding Tools MCP – Give any LLM agent the ability to code

- HN: [48594594](https://news.ycombinator.com/item?id=48594594)
- Source: [github.com](https://github.com/xyTom/coding-tools-mcp)
- Score: 1
- Comments: 0
- Posted: 2026-06-19T03:50:38Z

## Translation

タイトル: HN の表示: コーディング ツール MCP – LLM エージェントにコーディング機能を与える
記事のタイトル: GitHub - xyTom/coding-tools-mcp: AI エージェントにコーディング機能を与える · GitHub
説明: AI エージェントにコーディング機能を与えます。 GitHub でアカウントを作成して、xyTom/coding-tools-mcp の開発に貢献してください。

記事本文:
GitHub - xyTom/coding-tools-mcp: AI エージェントにコーディング機能を提供する · GitHub
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
キシトム
/
コーディングツール-mcp
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メイン ブランチ タグ 移動先

ファイル コード [その他のアクション] メニューを開く フォルダーとファイル
76 コミット 76 コミット .devcontainer .devcontainer .github/ workflows .github/ workflows ベンチマーク ベンチマークcoding_tools_mcpcoding_tools_mcp docs docs レポート レポート スクリプト スクリプト テスト テスト .dockerignore .dockerignore .gitignore .gitignore BENCHMARK.md BENCHMARK.md COMPLIANCE.md COMPLIANCE.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md SPEC.md SPEC.md docker-compose.yml docker-compose.yml pouret.lock tongue.lock pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
コーディング ツール MCP は、モデル中立のコーディング エージェント ランタイム MCP サーバーです。これは、ローカル コーディング プリミティブを任意の MCP クライアントに公開します。
リポジトリの検査 -> ファイルの検索/読み取り -> 構造化パッチの適用 -> テスト/コマンドの実行
-> stdin セッションと対話する -> git status/diff を検査する
これはプロンプトラッパーではありません。外部エージェント アカウント、メモリ、クラウド タスク、Web 検索、イメージ生成、モデル ルーティング、プラグイン マーケットプレイス、またはサブエージェント オーケストレーションを MCP ツールとして公開しません。
標準 MCP ランタイム プロファイル: docs/profile-v0.1.md
PyPI から公開されたコマンドをインストールします。
カール -fsSL https://raw.githubusercontent.com/xyTom/coding-tools-mcp/main/scripts/install.sh |バッシュ
ワークスペースに対してローカル Streamable HTTP をインストールして開始します。
カール -fsSL https://raw.githubusercontent.com/xyTom/coding-tools-mcp/main/scripts/install.sh \
| bash -s -- --start --workspace /path/to/repo
読み取り専用ベアラー トークン トンネルをインストールして公開します。
カール -fsSL https://raw.githubusercontent.com/xyTom/coding-tools-mcp/main/scripts/install.sh \
| bash -s -- --tunnel cloudflared --auto-install-tunnel --workspace /path/to/repo
または、このチェックアウトから:
スクリプト/install.sh
永続的なインストールを行わずに公開されたパッケージを実行します。
uvxコーディング-

tools-mcp --workspace 。
MCP クライアントには標準入出力を使用します。
uvxcoding-tools-mcp --stdio --workspace /path/to/repo
公開されたパッケージではなくこのチェックアウトから作業している場合:
始める
Make 変数を使用して、別のワークスペース、ホスト、ポート、または追加のサーバー フラグを渡します。
make start MCP_WORKSPACE=/path/to/repo MCP_PORT=8000 MCP_ARGS= " --permission-mode trusted "
依存関係が欠落している場合は、ランタイムを編集可能モードでインストールします。
python -m pip install -e " .[dev] "
HTTP エンドポイント:
http://127.0.0.1:8765/mcp
view_image の自動サイズ変更サポートが必要な場合は、オプションのイメージを追加でインストールします。
python -m pip install -e " .[画像] "
スタジオ:
coding-tools-mcp --stdio --workspace /path/to/repo
CODING_TOOLS_MCP_TRACE=1 を設定すると、編集された JSON ツール呼び出しトレース イベントがローカル デバッグ用に標準エラー出力に出力されます。ログは stdout から離れたままになるため、stdio JSON-RPC はクリーンなままになります。
デフォルトでは、exec_command はコア シェル環境のみを渡します。 MSVC 開発者プロンプトなど、継承された環境変数に依存するローカル ツールチェーンの場合は、次のことから始めます。
CODING_TOOLS_MCP_SHELL_ENV_INHERIT=allcoding-tools-mcp --workspace /path/to/repo
危険モードも有効になっていない限り、inherit=all はシークレットに見える変数とローダー/スタートアップ変数をフィルタリングします。依存関係のダウンロード、シェル拡張、インライン インタープリター スニペットを使用したローカル開発の場合は、次を使用します。
coding-tools-mcp --permission-mode trusted --workspace /path/to/repo
--allow-network は、ネットワークに見えるコマンドのみを開きたい場合に互換性フラグとして引き続き使用できます。 MCP クライアントが権限の引き出しをサポートしておらず、分離されたコンテナーまたは VM 内で exec_command 権限ゲートを明示的に無効にする場合は、次のことから始めます。
coding-tools-mcp --permission-modeangerous --workspace /path/to/repo
これにより、network-looki などの exec_command 権限ゲートが無効になります。

ng コマンド、破壊的なコマンド チェック、シェル拡張、インライン スクリプト、機密環境チェック。ダイレクト ファイル ツールのワークスペース パス境界は引き続き適用されます。 --dangerously-skip-all-permissions は互換性エイリアスとして残ります。
[ mcp_servers .コーディングツール]
コマンド = " uvx "
args = [ "coding-tools-mcp " 、 " --stdio " 、 " --workspace " 、 " /path/to/repo " ]
クロードコード:
{
"mcpサーバー": {
"コーディングツール" : {
"コマンド" : " uvx " 、
"args" : [ "coding-tools-mcp " 、 " --stdio " 、 " --workspace " 、 " /path/to/repo " ]
}
}
}
カーソル:
{
"mcpサーバー": {
"コーディングツール" : {
"コマンド" : " uvx " 、
"args" : [ "coding-tools-mcp " 、 " --stdio " 、 " --workspace " 、 " /path/to/repo " ]
}
}
}
一般的なストリーミング可能な HTTP クライアントは、MCP プロトコル バージョン 2025-06-18 を使用し、 http://127.0.0.1:8765/mcp をポイントする必要があります。
リモート MCP クライアントと HTTPS トンネル経由のローカル開発の場合は、サーバーをループバックにバインドしたままにして、クライアントが使用できる最も安全なプロファイルでトンネル URL を公開します。匿名トンネルのテストには読み取り専用モードを使用する必要があります。
CODING_TOOLS_MCP_AUTH_MODE=認証なし\
CODING_TOOLS_MCP_TOOL_PROFILE=読み取り専用 \
./scripts/tunnel.sh Cloudflared /path/to/repo
HTTPS トンネル URL を使用してリモート MCP クライアントを構成します。
URL: https://<トンネルホスト>/mcp
トンネル スクリプトは、cloudflared 、 ngrok 、および Microsoft Dev Tunnel をサポートします。選択したトンネル CLI が見つからない場合、スクリプトはそれをインストールする前に次の質問をします。
scripts/tunnel.sh Cloudflared /path/to/repo
scripts/tunnel.sh ngrok /path/to/repo
scripts/tunnel.sh devtunnel /path/to/repo
カスタム ヘッダーをサポートするクライアントの場合は、Authorization: Bearer <token> でベアラー トークン認証を使用します。 OAuth 2.1 認証コード + PKCE を使用する MCP クライアントの場合は、 scripts/tunnel.sh (または scripts/install.sh --auth-mode oauth ) で CODING_TOOLS_MCP_AUTH_MODE=oauth を使用します。サーバーは次のことができます

OAuth 発行者をトンネル リクエスト URL から推測するため、cloudflared のようなワンショット トンネルは、起動前に CODING_TOOLS_MCP_SERVER_URL を設定しなくても機能します。安定した発行者を固定したい場合にのみ設定してください。このスクリプトは、生成された OAuth パスワードを出力し、デフォルトで空でない client_id を受け入れ、機密クライアントをロックダウンする必要がある場合にのみ CODING_TOOLS_MCP_OAUTH_CLIENT_ID / CODING_TOOLS_MCP_OAUTH_CLIENT_SECRET をオプトインできるようにします。カスタム ベアラー ヘッダーを送信できず、OAuth を使用できないクライアントは、ローカル/テスト トンネルにのみ匿名読み取り専用モードを使用するか、運用環境で使用するために外部認証プロキシの背後に配置する必要があります。
正確なモードとセキュリティ上の注意については、docs/remote-mcp.md を参照してください。
full : すべてのツールを真実の注釈とともに公開します。これは下位互換性のためのデフォルトです。
read-only : リモートまたはセーフモードのクライアントに推奨。検査ツール、git 読み取りツール、イメージ表示、default-cwd ヘルパーのみを公開します。
compat-readonly-all : すべてのツールを公開しますが、 readOnlyHint で可用性を制御するクライアントに対してはすべてのツールを読み取り専用としてアドバタイズします。これは安全モードではありません。 apply_patch 、 exec_command 、 write_stdin 、および kill_session などの変更可能なツールは、引き続きローカル状態を変更できます。
デフォルトで公開される追加の画像ツール:
入出力スキーマと結果エンベロープについては、 docs/tools-and-schemas.md および docs/profile-v0.1.md を参照してください。
ランタイムは、サーバー プロセスごとに 1 つのワークスペース ルートをバインドします。デフォルトでは、パスはワークスペース相対です。絶対パス、トラバーサル、シンボリックリンクエスケープは拒否されます。再帰的なリスト/検索では、デフォルトで .git 、 .reference 、node_modules 、 target 、 dist 、ビルド出力、virtualenvs、および共通キャッシュが除外されます。
exec_command は、ワークスペースにバインドされた cwd、構成可能なシェル環境の継承、タイムアウトを使用したポリシー制御の下で実行されます。

、出力上限、機密値とローダー/起動環境の拒否、破壊的なコマンド チェック、ネットワークを参照するコマンド チェック、シェル拡張許可ゲート、間接的な絶対パス チェック、キャンセル/強制終了のクリーンアップ、セッション期限ウォッチドッグ、および制限されたセッション バッファー。 Landlock をサポートする Linux ホストでは、ファイルシステムの制限も適用されます。 Landlock のない Windows、macOS、または Linux ホストでは、コマンド結果に警告が含まれ、信頼できないコマンドを実行する前に外部サンドボックスが必要です。これはまだ完全な OS/コンテナ サンドボックスではありません。 SECURITY.md を参照してください。
--permission-mode セーフがデフォルトです。 --permission-mode trusted は、シークレット フィルタリングと破壊的コマンド チェックを維持しながら、ローカル開発ゲートを開きます。 --permission-modeangerous は、分離されたランナー内でそのリスクを受け入れるオペレーターの exec_command 権限ゲートを無効にします。信頼できないワークスペースまたは信頼できない MCP クライアントには危険モードを使用しないでください。
コンプライアンスを遵守する
コンプライアンスと CI コマンドは docs/ci-and-tests.md に文書化されています。チェックインされたレポート ファイルは生成されたアーティファクトです。完全なコンプライアンスの証拠として扱う前に、スイートのフィールドを検査してください。
Dogfood と SWE-bench のノートは docs/dogfood.md 、 docs/swe-bench.md 、および BENCHMARK.md にあります。このリポジトリは、モデルによって生成された SWE ベンチのリーダーボードの結果を要求しません。
糸くずを作る
タイプチェックを行う
テストを行う
コンプライアンスを遵守する
ciを作る
完全なテスト マトリックスについては、docs/ci-and-tests.md を参照してください。
このプロジェクトはソースから入手できますが、オープンソースではありません。 「ライセンス」を参照してください。
内部評価、開発、テスト、セキュリティレビューは許可されています。
再配布、ホストされたサードパーティ サービスの使用、および本番での商用使用
事前の書面による許可が必要です。
AI エージェントにコーディング機能を与える
コーディング-1afcb9be.mintlify.app
トピックス
ロー中にエラーが発生しました

アディング。このページをリロードしてください。
26
フォーク
レポートリポジトリ
リリース
1
v0.1.7
最新の
2026 年 6 月 11 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Give any AI agent the ability to code. Contribute to xyTom/coding-tools-mcp development by creating an account on GitHub.

GitHub - xyTom/coding-tools-mcp: Give any AI agent the ability to code · GitHub
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
xyTom
/
coding-tools-mcp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
76 Commits 76 Commits .devcontainer .devcontainer .github/ workflows .github/ workflows benchmarks benchmarks coding_tools_mcp coding_tools_mcp docs docs reports reports scripts scripts tests tests .dockerignore .dockerignore .gitignore .gitignore BENCHMARK.md BENCHMARK.md COMPLIANCE.md COMPLIANCE.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md SPEC.md SPEC.md docker-compose.yml docker-compose.yml poetry.lock poetry.lock pyproject.toml pyproject.toml View all files Repository files navigation
Coding Tools MCP is a model-neutral coding-agent runtime MCP server. It exposes local coding primitives to any MCP client:
inspect repo -> search/read files -> apply structured patches -> run tests/commands
-> interact with stdin sessions -> inspect git status/diff
It is not a prompt wrapper. It does not expose external agent accounts, memory, cloud tasks, web search, image generation, model routing, plugin marketplace, or subagent orchestration as MCP tools.
Normative MCP runtime profile: docs/profile-v0.1.md
Install the published command from PyPI:
curl -fsSL https://raw.githubusercontent.com/xyTom/coding-tools-mcp/main/scripts/install.sh | bash
Install and start local Streamable HTTP against a workspace:
curl -fsSL https://raw.githubusercontent.com/xyTom/coding-tools-mcp/main/scripts/install.sh \
| bash -s -- --start --workspace /path/to/repo
Install and expose a read-only bearer-token tunnel:
curl -fsSL https://raw.githubusercontent.com/xyTom/coding-tools-mcp/main/scripts/install.sh \
| bash -s -- --tunnel cloudflared --auto-install-tunnel --workspace /path/to/repo
Or, from this checkout:
scripts/install.sh
Run the published package without a persistent install:
uvx coding-tools-mcp --workspace .
Use stdio for MCP clients:
uvx coding-tools-mcp --stdio --workspace /path/to/repo
If you are working from this checkout instead of a published package:
make start
Pass a different workspace, host, port, or extra server flags with Make variables:
make start MCP_WORKSPACE=/path/to/repo MCP_PORT=8000 MCP_ARGS= " --permission-mode trusted "
If dependencies are missing, install the runtime in editable mode:
python -m pip install -e " .[dev] "
HTTP endpoint:
http://127.0.0.1:8765/mcp
Install the optional image extra when you want view_image auto-resize support:
python -m pip install -e " .[image] "
Stdio:
coding-tools-mcp --stdio --workspace /path/to/repo
Set CODING_TOOLS_MCP_TRACE=1 to emit redacted JSON tool-call trace events to stderr for local debugging. Logs stay off stdout so stdio JSON-RPC remains clean.
By default, exec_command passes a core shell environment only. For local toolchains that depend on inherited environment variables, such as MSVC developer prompts, start with:
CODING_TOOLS_MCP_SHELL_ENV_INHERIT=all coding-tools-mcp --workspace /path/to/repo
inherit=all still filters secret-looking and loader/startup variables unless dangerous mode is also enabled. For local development with dependency downloads, shell expansion, and inline interpreter snippets, use:
coding-tools-mcp --permission-mode trusted --workspace /path/to/repo
--allow-network remains available as a compatibility flag when you only want to open network-looking commands. If your MCP client does not support permission elicitation and you explicitly want to disable exec_command permission gates inside an isolated container or VM, start with:
coding-tools-mcp --permission-mode dangerous --workspace /path/to/repo
This disables exec_command permission gates such as network-looking commands, destructive command checks, shell expansion, inline scripts, and sensitive env checks. Workspace path boundaries for direct file tools still apply. --dangerously-skip-all-permissions remains as a compatibility alias.
[ mcp_servers . coding_tools ]
command = " uvx "
args = [ " coding-tools-mcp " , " --stdio " , " --workspace " , " /path/to/repo " ]
Claude Code:
{
"mcpServers" : {
"coding-tools" : {
"command" : " uvx " ,
"args" : [ " coding-tools-mcp " , " --stdio " , " --workspace " , " /path/to/repo " ]
}
}
}
Cursor:
{
"mcpServers" : {
"coding-tools" : {
"command" : " uvx " ,
"args" : [ " coding-tools-mcp " , " --stdio " , " --workspace " , " /path/to/repo " ]
}
}
}
Generic Streamable HTTP clients should use MCP protocol version 2025-06-18 and point at http://127.0.0.1:8765/mcp .
For remote MCP clients and local development over an HTTPS tunnel, keep the server bound to loopback and expose the tunnel URL with the safest profile your client can use. Anonymous tunnel testing should use read-only mode:
CODING_TOOLS_MCP_AUTH_MODE=noauth \
CODING_TOOLS_MCP_TOOL_PROFILE=read-only \
./scripts/tunnel.sh cloudflared /path/to/repo
Configure the remote MCP client with the HTTPS tunnel URL:
URL: https://<tunnel-host>/mcp
The tunnel scripts support cloudflared , ngrok , and Microsoft Dev Tunnel. If the selected tunnel CLI is missing, the script asks before installing it:
scripts/tunnel.sh cloudflared /path/to/repo
scripts/tunnel.sh ngrok /path/to/repo
scripts/tunnel.sh devtunnel /path/to/repo
For clients that support custom headers, use bearer-token auth with Authorization: Bearer <token> . For MCP clients that speak OAuth 2.1 Authorization Code + PKCE, use CODING_TOOLS_MCP_AUTH_MODE=oauth with scripts/tunnel.sh (or scripts/install.sh --auth-mode oauth ). The server can infer its OAuth issuer from the tunnel request URL, so one-shot tunnels like cloudflared work without setting CODING_TOOLS_MCP_SERVER_URL before startup; set it only when you want to pin a stable issuer. The script prints a generated OAuth password, accepts any non-empty client_id by default, and lets you opt into CODING_TOOLS_MCP_OAUTH_CLIENT_ID / CODING_TOOLS_MCP_OAUTH_CLIENT_SECRET only when you need to lock down a confidential client. Clients that cannot send custom bearer headers and do not speak OAuth should use anonymous read-only mode only for local/testing tunnels, or be placed behind an external auth proxy for production use.
See docs/remote-mcp.md for the exact modes and security notes.
full : exposes all tools with truthful annotations. This is the default for backward compatibility.
read-only : recommended for remote or safe-mode clients; exposes only inspection tools, git read tools, image viewing, and default-cwd helpers.
compat-readonly-all : exposes all tools but advertises every tool as read-only for clients that gate availability on readOnlyHint . This is not a safety mode; mutation-capable tools such as apply_patch , exec_command , write_stdin , and kill_session can still mutate local state.
Additional image tool exposed by default:
For input/output schemas and result envelopes, see docs/tools-and-schemas.md and docs/profile-v0.1.md .
The runtime binds one workspace root per server process. Paths are workspace-relative by default. Absolute paths, .. traversal, and symlink escapes are rejected. Recursive listing/search excludes .git , .reference , node_modules , target , dist , build outputs, virtualenvs, and common caches by default.
exec_command runs under policy controls with workspace-bound cwd, configurable shell environment inheritance, timeout, output caps, sensitive-value and loader/startup environment rejection, destructive command checks, network-looking command checks, shell-expansion permission gates, indirect absolute-path checks, cancellation/kill cleanup, session deadline watchdogs, and bounded session buffers. On Linux hosts with Landlock support it also applies filesystem confinement; on Windows, macOS, or Linux hosts without Landlock, command results include a warning and external sandboxing is required before running untrusted commands. This is still not a complete OS/container sandbox; see SECURITY.md .
--permission-mode safe is the default. --permission-mode trusted opens local-development gates while keeping secret filtering and destructive-command checks. --permission-mode dangerous disables exec_command permission gates for operators who accept that risk inside an isolated runner. Do not use dangerous mode for untrusted workspaces or untrusted MCP clients.
make compliance
Compliance and CI commands are documented in docs/ci-and-tests.md . The checked-in report files are generated artifacts; inspect their suite field before treating them as full compliance evidence.
Dogfood and SWE-bench notes live in docs/dogfood.md , docs/swe-bench.md , and BENCHMARK.md . This repository does not claim a model-generated SWE-bench leaderboard result.
make lint
make typecheck
make test
make compliance
make ci
See docs/ci-and-tests.md for the full test matrix.
This project is source-available, not open source. See LICENSE .
Internal evaluation, development, testing, and security review are permitted;
redistribution, hosted third-party service use, and production commercial use
require prior written permission.
Give any AI agent the ability to code
coding-1afcb9be.mintlify.app
Topics
There was an error while loading. Please reload this page .
26
forks
Report repository
Releases
1
v0.1.7
Latest
Jun 11, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
