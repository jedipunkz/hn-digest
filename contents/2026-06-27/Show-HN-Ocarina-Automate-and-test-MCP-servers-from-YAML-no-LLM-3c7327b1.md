---
source: "https://github.com/msradam/ocarina"
hn_url: "https://news.ycombinator.com/item?id=48699141"
title: "Show HN: Ocarina – Automate and test MCP servers from YAML, no LLM"
article_title: "GitHub - msradam/ocarina: Automation framework for MCP servers. Write a YAML playbook, replay it deterministically against real servers, no LLM in the loop. · GitHub"
author: "msradam"
captured_at: "2026-06-27T16:23:31Z"
capture_tool: "hn-digest"
hn_id: 48699141
score: 2
comments: 0
posted_at: "2026-06-27T15:33:22Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Ocarina – Automate and test MCP servers from YAML, no LLM

- HN: [48699141](https://news.ycombinator.com/item?id=48699141)
- Source: [github.com](https://github.com/msradam/ocarina)
- Score: 2
- Comments: 0
- Posted: 2026-06-27T15:33:22Z

## Translation

タイトル: Show HN: Ocarina – LLM なしで YAML から MCP サーバーを自動化およびテストする
記事のタイトル: GitHub - msradam/ocarina: MCP サーバーの自動化フレームワーク。 YAML プレイブックを作成し、実サーバーに対して決定論的に再生します。ループ内に LLM はありません。 · GitHub
説明: MCP サーバー用の自動化フレームワーク。 YAML プレイブックを作成し、実サーバーに対して決定論的に再生します。ループ内に LLM はありません。 - ムスラダム/オカリナ
HN テキスト: こんにちは。 Ansible やその他の自動化フレームワークに何年も取り組んできた者として、私は最近の MCP ブームに魅了されています。人々は、ツールとリソースの両方を公開する標準化されたプロトコルを介して、サーバー用に、型付きで LLM で読み取り可能な (つまり人間が判読できる) インターフェースを作成しています。 AI を関与させずに、これらのサーバーに対してスクリプトを直接実行する方法を作成できるかどうかを確認したいと思いました。 Ocarina を使用すると、MCP サーバーの特性を検査し、yaml でツール呼び出しを表現する Rondos (Ansible プレイブックに相当) を作成できるため、LLM を使用せずに段階的に再生できます。ロンドは次のようになります: キー:
オーナー：アクメ
リポジトリ: API
サーバー:
コマンド: npx
引数: [-y, "@modelcontextprotocol/server-github"]
ロンド：
- 名前: 最近のコミット
ツール: list_commits
引数:
所有者: 「{{所有者}}」
リポジトリ: "{{リポジトリ}}"
グラブ: ".0.sha"
エコー: 最新_sha
- 名前: コミットの詳細
ツール: get_commit
引数:
所有者: 「{{所有者}}」
リポジトリ: "{{リポジトリ}}"
シャ: "{{latest_sha}}"
期待する:
含まれるもの：「偉業」
私は仕事で作成している MCP サーバーをテストおよび検証するために自分自身で Ocarina を使い始めました。そして、大規模な MCP エコシステムを最大限に活用する方法というより広い分野に興味があります。 README には、クローンを作成してすぐに試すことができるリポジトリへのリンクが含まれています (そして、ちょっとした Blender デモ ;]) 土曜日おめでとうございます!

記事本文:
GitHub - msradam/ocarina: MCP サーバー用の自動化フレームワーク。 YAML プレイブックを作成し、実サーバーに対して決定論的に再生します。ループ内に LLM はありません。 · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
ムスラダム
/
オカリナ
公共
通知
変更するにはサインインする必要があります

化設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
36 コミット 36 コミット .github/ workflows .github/ workflows 資産 アセット cmd cmd docs docs 例 例 内部 内部 .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yaml .goreleaser.yaml ライセンス ライセンス README.md README.md action.yml action.ymlデモ.テープ デモ.テープ go.mod go.mod go.sum go.sum main.go main.go mcp.json.example mcp.json.example すべてのファイルを表示 リポジトリ ファイルのナビゲーション
MCP サーバーの自動化フレームワーク。 YAML スクリプトを作成し、ループ内に LLM を使用せずに決定論的に再生します。
Blender-mcp を運転するロンド: 平面を置き、立方体をドロップし、球を積み上げ、円錐を追加して、シーンを確認します。実行ごとに同じ YAML、同じ結果。モデルは関係ありません。 Blender-mcp-ocarina のクローンを作成して自分で実行します。
MCP エコシステムはすでに巨大です。型指定されたツール、読み取り可能なリソース、スキーマチェックされたコントラクトを通じて実際のサービスを公開する何千ものサーバーがデプロイされ、すぐに呼び出すことができます。 Ocarina は、そのすべてを自動化するフレームワークです。 1 つ以上のサーバー間でツールを駆動し、ステップ、分岐、ループ、再試行間で値をパイプし、毎回同じ方法で実行する YAML スクリプトを作成します。ループ内に LLM がないため、すべての実行は再現可能であり、コストはかかりません。
これらのツールは言語モデルによって読み取れるように構築されており、言語モデルは人間の言語でトレーニングされているため、ツールは人間にとってもきれいに読み取れます。サーバーは、コード内で接続するエンドポイントではなく、 get_issues や query_database などの名前付きコントラクトを公開します。 AI アシスタント用に誰かが構築したサーバーはすべて、あなたが運転できるものです。
作成するのはプレイブックです。これは、ワイヤー プロトコルとして MCP を使用して、サーバー上の自動化ワークフローをキャプチャするポータブルな成果物です。読めるよ、リヴィ

プルリクエストで作成し、バージョン管理し、サーバーがアクセス可能な場所ならどこでも実行します。手書きするか、エージェントに生成してもらいます。どちらの方法でも、サンプリングやトークンはなく、ファイルと結果の間の推論も行われず、実行のたびに同じように実行されます。
github.com/msradam/ocarina@latest をインストールしてください
バイナリはリリース ページから入手できます。ソースからビルドするには Go 1.26 以降が必要です。
サーバーのマークダウン ドキュメントを生成します。
ocarina docs uvx mcp-server-sqlite --db-path mydb.sqlite
ocarina docs npx -y @modelcontextprotocol/server-github > docs/github.md
ロンドを実行します。
オカリナ演奏 db-audit.yaml
オカリナ演奏 db-audit.yaml --dry-run
ocarina play db-audit.yaml -e db=/tmp/other.sqlite # 実行時にキーをオーバーライドします
ツールを実行せずに、ライブサーバーに対してロンドを検証します。
オカリナは db-audit.yaml を検証します
設計原則
決定論的。同じロンドを実行すると、実行するたびに同じ結果が得られます。サンプリングもランダム性もありません。
プロトコルネイティブ。 tools/call 、 resource/read 、および resource/list を介して MCP と直接通信します。準拠したサーバーであればどれでも動作します。
アサーションは一流です。何らかのexpect:チェックが失敗した場合、playはゼロ以外で終了します。 Rondos は、すぐに使用できる CI ヘルスチェックとして機能します。
スクリプトに資格情報がありません。サーバー接続および環境変数は、rondo ファイルの外に残ります。
ロンドは 1 台、どのマシンでも。 MCP サーバーが利用可能な場合、ロンドが実行されます。
ロンドは 3 つのセクションからなる YAML ファイルです。
キー:
オーナー：アクメ
リポジトリ: API
サーバー:
コマンド：npx
引数: [-y, "@modelcontextprotocol/server-github"]
ロンド：
- 名前 : 最近のコミット
ツール: list_commits
引数:
所有者: " {{所有者}} "
リポジトリ: " {{リポジトリ}} "
グラブ: " .0.sha "
エコー : 最新_社
- 名前 : コミットの詳細
ツール: get_commit
引数:
所有者: " {{所有者}} "
リポジトリ: " {{リポジトリ}} "
シャ: " {{latest_sha}} "
期待する:
内容:「偉業」
ステップフィールド
フィールド
D

説明
サーバー
このステップを実行するサーバー (サーバー内のキー: マップ)。デフォルトは唯一の/最初のサーバーです
ツール
呼び出すツール名
リソース
読み取るリソース URI ( resource/read )
リスト_リソース
リソースをリストするサーバーのプレフィックス。出力はJSON URI配列です
引数
ツールの引数。 {{key}} はキーまたは以前のエコー キャプチャから補間します
エコー
後のステップのためにこのステップの出力をキーの下に保存します
掴む
保存前の JSON 出力へのドットパス: .0.sha 、 .name 、 .items.0.id
ループ
JSON 配列キーを繰り返しの反復に展開します。毎回 {{item}} を設定します
期待する。含まれる
出力にこの文字列が含まれていることを確認します
期待一致
アサート出力がこの正規表現と一致する
期待する。等しい
出力がこの文字列と等しいことをアサートします (空白がトリミングされた)
期待される.is_error
ツールが isError: true を返したかどうかをアサートします。
無視エラー
過去の失敗を立ち止まるのではなく継続する
タグ
このステップに --tags / --skip-tags フィルタリングのタグを付けます
{{env.NAME}} はプロセス環境から解決され、{{key}} が動作する場所であればどこでも動作します。
Ansible から来たのですか? task: は rondo: のエイリアスとして受け入れられ、 register: は echo: のエイリアスとして受け入れられます。
1 つのロンドは複数のサーバーと通信できます。これらをservers:の下で宣言し、各ステップでserver:を設定します。サーバーを省略する手順: 最初のエントリを使用します。
サーバー:
時間: {コマンド: uvx、引数: [mcp-server-time]}
fetch : {コマンド: uvx、引数: [-y, "@modelcontextprotocol/server-fetch"]}
ロンド：
- 名前 : 時間を取得します
サーバー: 時間
ツール: get_current_time
引数: {タイムゾーン: UTC}
- 名前 : ページを取得する
サーバー:フェッチ
ツール：フェッチ
引数: {url: "https://example.com"}
サーバー別の名前空間ツール名を出力および差分します ( time.get_current_time )。単一サーバー: ブロックは、単一サーバーのロンドでも引き続き機能します。
ocarina docs <command> [args...] : すべてのツール、リソース、およびサーバー exp のリソース テンプレートのマークダウン ドキュメントを生成します。

oses.
ocarina play <rondo.yaml> : ライブサーバーに対して各ステップを実行します。
ocarina validate <rondo.yaml> : 呼び出しを行わずに、ツール名、必要な引数、スキーマ タイプ、および {{key}} データ フローをチェックします。
ocarina hum <command> [args...] -- <tool> [key=value ...] : 単一のツールを呼び出し、結果を出力します。
オカリナレコード <output.yaml> <コマンド> [args...] : プロキシモード;ライブ MCP クライアント セッションからのすべてのツール呼び出しを rondo ファイルに記録します。
.mcp.json (または資格情報の ~/.mcp.json) を作成し、rondos およびコマンド ラインでサーバーを名前で参照します。
{
"mcpServers" : {
"github" : {
"command" : " npx " ,
"args" : [ " -y " , " @modelcontextprotocol/server-github " ],
"env" : { "GITHUB_PERSONAL_ACCESS_TOKEN" : " ghp_... " }
}
}
}
server : github
オカリナ ハム github -- list_commits owner=pytorch repo=pytorch per_page=1
スターター テンプレートについては、mcp.json.example を参照してください。 Ocarina は、Claude デスクトップ構成 ( ~/Library/Application Support/Claude/claude_desktop_config.json ) からもサーバーを検出します。
50 を超える MCP サーバーで動作するロンドは、examples/ にあります。 A selection:
完全なリストについては、docs/tested-servers.md を参照してください。
クローンを作成して実行できるスタンドアロン リポジトリ。それぞれが異なる MCP サーバーの実際の作業環境です。
duckdb-mcp-ocarina : DuckDB データベースに対するデータの整合性、移行、および回帰テスト。クローンを作成して実行します。認証情報は必要ありません。
chrome-devtools-mcp-ocarina : Google の Chrome DevTools MCP を介した合成 Web ヘルスチェック。コンソールエラーまたは失敗したリクエストで失敗します。
github-mcp-ocarina : GitHub MCP サーバーを介したテストとしてのリポジトリ ガバナンス。リポジトリにはライセンスが同梱され、文書化され、履歴があることをアサートします。
Blender-mcp-ocarina : 外部 API をまったく持たないアプリである Blender で 3D シーンを自動化し、スナップショット テストします。
すべてが期待される場合、play は 0 を終了します。アサーションはパスし、それ以外の場合は 0 以外です。

伊勢。ロンドを任意の CI パイプラインにドロップします。
- 名前 : データベースのヘルスチェック
実行：オカリナ演奏ロンドス/db-audit.yaml
複合 GitHub アクションは Ocarina をインストールし、ロンドを再生します。
- 使用: msradam/ocarina@v1
付き:
ロンド：tests/mcp-smoke.yaml
action.yml および .github/workflows/example.yml を参照してください。
マサチューセッツ工科大学Noun Project の Alessio Capponi によるホイッスルのアイコン (CC BY 3.0)。
MCP サーバーの自動化フレームワーク。 YAML プレイブックを作成し、実サーバーに対して決定論的に再生します。ループ内に LLM はありません。
msradam.github.io/ocarina/
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
v0.1.0
最新の
2026 年 6 月 27 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Automation framework for MCP servers. Write a YAML playbook, replay it deterministically against real servers, no LLM in the loop. - msradam/ocarina

Hi all. As someone who has spent years working with Ansible and other automation frameworks, the recent MCP boom has me fascinated. People are creating nice, typed, LLM-readable (and thus human-readable) interfaces for their servers, over a standardized protocol that exposes both tools and resources. I wanted to see if I could create a way to run scripts against these servers directly, with no AI in the loop. Ocarina lets you inspect MCP server characteristics and write Rondos (the equivalent of Ansible playbooks) that express tool calls in yaml so you can play them step-by-step without needing to drive with an LLM. Here's what a Rondo looks like: keys:
owner: acme
repo: api
server:
command: npx
args: [-y, "@modelcontextprotocol/server-github"]
rondo:
- name: recent commits
tool: list_commits
args:
owner: "{{owner}}"
repo: "{{repo}}"
grab: ".0.sha"
echo: latest_sha
- name: commit detail
tool: get_commit
args:
owner: "{{owner}}"
repo: "{{repo}}"
sha: "{{latest_sha}}"
expect:
contains: "feat"
I have started using Ocarina myself to test and validate the MCP servers I am creating at my job, and I am interested in the broader field of how we can maximize the use of the massive MCP ecosystem. The README includes links to repos you can clone and try right away (and a little Blender demo ;]) Happy Saturday!

GitHub - msradam/ocarina: Automation framework for MCP servers. Write a YAML playbook, replay it deterministically against real servers, no LLM in the loop. · GitHub
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
msradam
/
ocarina
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
36 Commits 36 Commits .github/ workflows .github/ workflows assets assets cmd cmd docs docs examples examples internal internal .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yaml .goreleaser.yaml LICENSE LICENSE README.md README.md action.yml action.yml demo.tape demo.tape go.mod go.mod go.sum go.sum main.go main.go mcp.json.example mcp.json.example View all files Repository files navigation
An automation framework for MCP servers. Write a YAML script, replay it deterministically, no LLM in the loop.
A rondo driving blender-mcp : lay down a plane, drop a cube, stack a sphere, add a cone, then verify the scene. Same YAML, same result, every run. No model involved. Clone blender-mcp-ocarina to run it yourself.
The MCP ecosystem is already enormous: thousands of servers exposing real services through typed tools, readable resources, and schema-checked contracts, deployed and ready to call. Ocarina is an automation framework for all of it. Write a YAML script that drives tools across one or more servers, pipes values between steps, branches, loops, and retries, and runs the same way every time. No LLM in the loop, so every run is reproducible and costs nothing.
These tools were built to be read by language models, and language models are trained on human language, so the tools read cleanly to people too. A server exposes named contracts like get_issues or query_database , not endpoints you wire up in code. Every server someone built for an AI assistant is one you can drive.
What you write is a playbook: a portable artifact that captures an automation workflow over those servers, with MCP as the wire protocol. You can read it, review it in a pull request, version it, and run it anywhere the servers are reachable. Write it by hand or have an agent generate it. Either way it runs the same on every execution, with no sampling, no tokens, and nothing inferring between the file and the result.
go install github.com/msradam/ocarina@latest
Binaries are available on the releases page . Building from source requires Go 1.26+.
Generate markdown docs for a server:
ocarina docs uvx mcp-server-sqlite --db-path mydb.sqlite
ocarina docs npx -y @modelcontextprotocol/server-github > docs/github.md
Run a rondo:
ocarina play db-audit.yaml
ocarina play db-audit.yaml --dry-run
ocarina play db-audit.yaml -e db=/tmp/other.sqlite # override a key at runtime
Validate a rondo against the live server without running any tools:
ocarina validate db-audit.yaml
Design principles
Deterministic. The same rondo produces the same result on every run. No sampling, no randomness.
Protocol-native. Talks MCP directly via tools/call , resources/read , and resources/list . Works with any compliant server.
Assertions are first-class. play exits non-zero if any expect: check fails. Rondos work as CI health checks out of the box.
No credentials in scripts. Server connection and environment variables stay outside the rondo file.
One rondo, any machine. If the MCP server is available, the rondo runs.
A rondo is a YAML file with three sections.
keys :
owner : acme
repo : api
server :
command : npx
args : [-y, "@modelcontextprotocol/server-github"]
rondo :
- name : recent commits
tool : list_commits
args :
owner : " {{owner}} "
repo : " {{repo}} "
grab : " .0.sha "
echo : latest_sha
- name : commit detail
tool : get_commit
args :
owner : " {{owner}} "
repo : " {{repo}} "
sha : " {{latest_sha}} "
expect :
contains : " feat "
Step fields
Field
Description
server
Which server to run this step against (a key in the servers: map); defaults to the only/first server
tool
Tool name to call
resource
Resource URI to read ( resources/read )
list_resources
Server prefix to list resources from; output is a JSON URI array
args
Tool arguments. {{key}} interpolates from keys or prior echo captures
echo
Store this step's output under a key for later steps
grab
Dot-path into JSON output before storing: .0.sha , .name , .items.0.id
loop
Expand a JSON array key into repeated iterations; sets {{item}} each time
expect.contains
Assert output contains this string
expect.matches
Assert output matches this regex
expect.equals
Assert output equals this string (whitespace-trimmed)
expect.is_error
Assert whether the tool returned isError: true
ignore_errors
Continue past failures instead of halting
tags
Tag this step for --tags / --skip-tags filtering
{{env.NAME}} resolves from the process environment and works anywhere {{key}} does.
Coming from Ansible? tasks: is accepted as an alias for rondo: , and register: as an alias for echo: .
A single rondo can talk to more than one server. Declare them under servers: and set server: on each step. Steps that omit server: use the first entry.
servers :
time : {command: uvx, args: [mcp-server-time]}
fetch : {command: uvx, args: [-y, "@modelcontextprotocol/server-fetch"]}
rondo :
- name : get time
server : time
tool : get_current_time
args : {timezone: UTC}
- name : fetch page
server : fetch
tool : fetch
args : {url: "https://example.com"}
Output and diff namespace tool names by server ( time.get_current_time ). The single server: block still works for one-server rondos.
ocarina docs <command> [args...] : generate markdown documentation for every tool, resource, and resource template a server exposes.
ocarina play <rondo.yaml> : execute each step against the live server.
ocarina validate <rondo.yaml> : check tool names, required args, schema types, and {{key}} data flow without making any calls.
ocarina hum <command> [args...] -- <tool> [key=value ...] : call a single tool and print the result.
ocarina record <output.yaml> <command> [args...] : proxy mode; records every tool call from a live MCP client session into a rondo file.
Create a .mcp.json (or ~/.mcp.json for credentials) and reference servers by name in rondos and on the command line:
{
"mcpServers" : {
"github" : {
"command" : " npx " ,
"args" : [ " -y " , " @modelcontextprotocol/server-github " ],
"env" : { "GITHUB_PERSONAL_ACCESS_TOKEN" : " ghp_... " }
}
}
}
server : github
ocarina hum github -- list_commits owner=pytorch repo=pytorch per_page=1
See mcp.json.example for a starter template. Ocarina also discovers servers from the Claude Desktop config ( ~/Library/Application Support/Claude/claude_desktop_config.json ).
Working rondos for 50+ MCP servers are in examples/ . A selection:
See docs/tested-servers.md for the full list.
Standalone repositories you can clone and run, each a real working environment for a different MCP server:
duckdb-mcp-ocarina : data integrity, migration, and regression tests against a DuckDB database. Clone and run, no credentials.
chrome-devtools-mcp-ocarina : synthetic web health checks through Google's Chrome DevTools MCP. Fail on a console error or a failed request.
github-mcp-ocarina : repo governance as tests through the GitHub MCP server. Assert a repo ships a license, is documented, and has history.
blender-mcp-ocarina : automate and snapshot-test a 3D scene in Blender, an app with no external API at all.
play exits 0 if all expect: assertions pass, non-zero otherwise. Drop a rondo into any CI pipeline:
- name : Database health check
run : ocarina play rondos/db-audit.yaml
A composite GitHub Action installs Ocarina and replays a rondo:
- uses : msradam/ocarina@v1
with :
rondo : tests/mcp-smoke.yaml
See action.yml and .github/workflows/example.yml .
MIT. Whistle icon by Alessio Capponi from Noun Project (CC BY 3.0).
Automation framework for MCP servers. Write a YAML playbook, replay it deterministically against real servers, no LLM in the loop.
msradam.github.io/ocarina/
Topics
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
v0.1.0
Latest
Jun 27, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
