---
source: "https://github.com/thesanjeetc/parselbox"
hn_url: "https://news.ycombinator.com/item?id=49388608"
title: "Show HN: Parselbox – an embeddable Python sandbox for AI agents"
article_title: "GitHub - thesanjeetc/Parselbox: An embeddable Python sandbox where AI agents call tools as code. Powered by Deno and Pyodide. · GitHub"
image: "https://opengraph.githubassets.com/4fce08f68d80f27478a3d8e09483e2caf56a55cf326c1eedac93ce8362823c9b/thesanjeetc/Parselbox"
author: "sanjeetc"
captured_at: "2026-08-21T14:25:02Z"
capture_tool: "hn-digest"
hn_id: 49388608
score: 1
comments: 0
posted_at: "2026-08-21T14:23:56Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Parselbox – an embeddable Python sandbox for AI agents

- HN: [49388608](https://news.ycombinator.com/item?id=49388608)
- Source: [github.com](https://github.com/thesanjeetc/parselbox)
- Score: 1
- Comments: 0
- Posted: 2026-08-21T14:23:56Z

## Translation

タイトル: Show HN: Parselbox – AI エージェント用の埋め込み可能な Python サンドボックス
記事のタイトル: GitHub - thesanjeetc/Parselbox: AI エージェントがツールをコードとして呼び出す埋め込み可能な Python サンドボックス。 Deno と Pyodide が搭載。 · GitHub
説明: AI エージェントがツールをコードとして呼び出す埋め込み可能な Python サンドボックス。 Deno と Pyodide が搭載。 - thesanjeetc/パーセルボックス

記事本文:
GitHub - thesanjeetc/Parselbox: AI エージェントがツールをコードとして呼び出す埋め込み可能な Python サンドボックス。 Deno と Pyodide が搭載。 · GitHub
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
テサンジーなど
/
パーセルボックス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
7 コミット 7 コミット フォルダーとファイル
.claude/ スキル/ parselbox .claude/ スキル/ parselbox .github/ workflows .github/ workflows アセット アセットの例 例 parselbox parsel

ボックス テスト テスト .dockerignore .dockerignore .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version CLAUDE.md CLAUDE.md Dockerfile Dockerfile LICENSE.md LICENSE.md Makefile Makefile README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
コード。ファイルシステム。コンテクスト。ツール。
エージェントがすべてを管理する 1 つのツールを持っていたらどうなるでしょうか?
Parselbox は、AI エージェントがコードとしてツールを呼び出す埋め込み可能な Python サンドボックスです。MCP サーバー、API、シェルはネイティブ Python オブジェクトになります。ディスクベースのワークスペース、パッケージ、ネットワークが組み込まれています。 Deno と Pyodide を利用した単一プロセス。
Parselbox MCP を既存の MCP サーバー構成と一緒に削除します。エージェントは、Python ランタイム、コードとしての MCP ツール、スキルのサポート、ディスクバックアップのワークスペースを即座に利用できます。
🔒 安全な隔離
コンテナーや VM はなく、単一の軽量の Deno + Pyodide プロセス (約 160 MB) だけです。 Deno 権限、メモリ上限、タイムアウト、ネットワーク許可リスト。スナップショット キャッシュとクラッシュ回復。
MCP サーバー、REST + OpenAPI、GraphQL、シェル、関数、クラス - すべてネイティブ Python オブジェクト。通話全体でステートフル。 Pydantic 自動変換。資格情報はホスト上に残ります。
js() 相互運用機能を備えた完全な CPython — JS パッケージをネイティブ Python として使用します。 npm、ローカル TypeScript、および .wasm モジュールの require()。シェルの仮想 bash()。インポート時にパッケージを自動インストールします。
require() 任意の .wasm — ライブラリのエクスポートは Python メソッドになり、WASI プログラムは呼び出し可能なコマンドになります。 bash() からも実行するには、bin/ に 1 つをドロップします。インプロセスでは、サンドボックスのマウントと権限を継承し、ホストには何もインストールしません。
任意の呼び出しに .task() を追加します。 asyncio.gather で並列ファンアウトし、進行状況を確認し、ログを追跡し、 send() で対話型セッションを開始し、後で待機します。
ディスクバ

cked ワークスペース — ホスト マウント ( ro / rw )、入力ファイルは /files/ にあり、出力は実際のディレクトリに保持されます。新しいファイルと変更されたファイルは呼び出しごとに検出されて返されます。
help() 、 search() 、inspect() 、review() — エージェントは、必要なときに、必要なものだけを検出します。
display() は、Tailwind + daisyUI を挿入して、チャット (MCP アプリ) で HTML をインラインでレンダリングします。または、静的ファイル、ライブ リロード、ファイル アップロード、ツール間で構成される @api ルートを備えた組み込み HTTP サーバーなど、完全なアプリを提供します。
Parselbox は、安全なサンドボックス ランタイムに Deno を使用します。
# macOS / Linux
カール -fsSL https://deno.land/install.sh |しー
# Windows (PowerShell)
IRM https://deno.land/install.ps1 |アイエックス
2. Parselboxをインストールする
pip インストール パーセルボックス
パーセルボックス API
MCP サーバー、REST/GraphQL、シェル、ホスト オブジェクトなど、あらゆるツールをサンドボックスに接続すると、エージェントはそれらをネイティブ Python として呼び出し、ディスクにバックアップされたワークスペースと Python と npm パッケージの両方のエコシステム上の実際の制御フローでそれらを構成します。
非同期をインポートする
OSをインポートする
textwrap インポートからのデデント
パーセルボックスからパーセルボックスをインポート
パーセルボックスから。ブリッジインポート HTTPBridge 、 ShellBridge
クラス分析:
def summary (self , repos : list ) -> dict :
"""リポジトリ統計を集約します。"""
星 = [ r [ "星" ] リポジトリ内の r ]
return { "count" : len (リポジトリ), "avg_stars" :round (sum (stars) / len(stars))}
config = { "mcpServers" : { "playwright" : { "command" : "npx" , "args" : [ "@playwright/mcp@latest" ]}}}
非同期デフォルトメイン():
Parselbox と非同期 (
mcp = 構成 、
コンテキスト = {
"分析" : 分析 ()、
"github" : HTTPBridge (base_url = "https://api.github.com" 、token = os .environ [ "GITHUB_TOKEN" ])、
"sh" : ShellBridge ( "bash" )、
}、
ネットワーク = True 、
allow_runtime_packages = True 、
パッケージ = [ "numpy" , "npm:lodash" ],
出力ディレクトリ = "./ワークスペース" ,
) sbx として:

# 利用可能なツールを発見する
sbx を待ちます。 execute_code ( "sbx.search('navigate|get')" )
# 実際のブラウザで GitHub リンクの Hacker News をスクレイピング
sbx を待ちます。 execute_code ( dedent ( """
輸入再
playwright.browser_navigate(url="https://news.ycombinator.com")
text = playwright.browser_snapshot()
repos = re.findall(r'github \\ .com/([ \\ w.-]+/[ \\ w.-]+)', text)[:5]
""" ))
# スター数を並行して取得し、コンテキスト ブリッジを介して要約する
sbx を待ちます。 execute_code ( dedent ( """
非同期をインポートする
results = await asyncio.gather(*[github.get.task(f"/repos/{r}") for r in repos])
repo_data = [{"name": r["data"]["name"], "stars": r["data"]["stargazers_count"]}
for r の結果 if r.get("ok")]
分析.要約(repo_data)
""" ))
# グラフ化する — matplotlib はインポート時に自動インストールされます
結果 = sbx を待ちます。 execute_code ( dedent ( """
matplotlib.pyplotをpltとしてインポート
plt.barh([r["name"] for r in repo_data], [r["stars"] for r in repo_data])
plt.savefig("チャート.png")
""" ))
print ( result . files ) # ['chart.png']
画像 = sbx 。 read_file ( "chart.png" )
# すべての結果には .output、.files、.stdout、.stderr、.error が含まれます
# サンドボックス全体を MCP サーバーとして提供する
sbx を待ちます。 r
[切り捨てられた]
Parselbox CLI はスタンドアロン MCP サーバーを実行します。すべてのサンドボックス オプションはフラグとして利用できます。
Parselbox MCP を既存の MCP サーバーと一緒に追加します。
--mcp を同じ構成ファイルに指定します。
起動時に、Parselbox は他のサーバーに接続し、サンドボックス内でツールを公開し、独自の MCP サーバーを起動します。
心配しないでください。Parselbox はそれ自体への接続を検出し、回避します。破滅の無限ループはありません。
{
"mcpサーバー": {
「github」: {}、
「線形」: {}、
"パーセルボックス" : {
"コマンド" : " uvx " ,
"args" : [ " parselbox " 、 " --mcp " 、 " /absolute/path/to/mcp.json " ]
}
}
}
HTTP
uvx parselbox --mcp mcp.json --trans

ポート http --ポート 9000
{
"mcpサーバー": {
"パーセルボックス" : {
"type" : " http " ,
"url" : " http://localhost:9000/mcp "
}
}
}
完全な例
uvx パーセルボックス \
--mcp ./mcp.json \
--transport http \
--ホスト 0.0.0.0 \
--ポート 8080 \
--file hello.txt \
--mount ./datasets:/data:rw \
--output-dir ./outputs \
--packages パンダ、matplotlib \
--package-dir ./cache \
--allow-runtime-packages \
--ネットワーク\
-- 3000 \ を提供します
--メモリ 2048 \
--タイムアウト 60 \
--env MY_API_KEY=...
パーセルボックスエージェント
非同期をインポートする
パーセルボックスからパーセルボックスをインポート
エージェントからインポート Agent 、 Runner 、 function_tool
サンドボックス = パーセルボックス (
mcp = { "mcpServers" : { "playwright" : { "command" : "npx" , "args" : [ "@playwright/mcp@latest" ]}}},
Output_dir = "./outputs" ,
allow_runtime_packages = True 、
)
エージェント = エージェント (
名前 = "研究助手" ,
モデル = "gpt-5.5" 、
命令 = f"あなたは世界クラスの研究助手です。\n \n { Sandbox . get_prompt () } " ,
tools = [ function_tool ( サンドボックス . get_tool ())],
)
非同期デフォルトメイン():
サンドボックスとの非同期:
result = ランナーを待ちます。走って（
代理店、
「Playwright MCP を使用して、ウィキペディアの「最も興行収入の高い映画のリスト」をスクレイピングします。」
"上位 10 位の棒グラフをプロットし、./plot.png として保存します" ,
max_turns = 30 、
)
print (結果.final_output)
非同期 。実行 (メイン ())
ユーザーガイド
コンテキスト ブリッジは、サンドボックス内のホスト Python オブジェクトを公開します。
context — 呼び出し可能なツールとしての関数と名前空間。実行は一時停止され、ホスト上で実行され、結果が返されます。
グローバル — サンドボックスにコピーされた静的な値 (文字列、数値、辞書)。
mcp — MCP サーバー構成 (辞書またはパス)。サンドボックス内の呼び出し可能な名前空間として表示されます。
プレーン クラスは自動ラップされ、すべてのパブリック メソッドが呼び出し可能なツールになります。 _ で始まるメソッドは非公開のままです。
パーセルボックスからパーセルボックスをインポート
クラス電卓:
def add ( self ,

a: フロート、b: フロート) -> フロート:
"""数字を 2 つ足します。"""
a + bを返す
Parselbox ( context = { "calc" : Calculator ()}) を sbx として非同期:
sbx を待ちます。 execute_code ( "calc.add(a=10, b=20)" )
ネストされた名前空間のサブクラス ブリッジ (自動クロール)。パラメーターに Pydantic モデルの注釈を付け、渡された辞書は自動的にそれに変換されます。
パーセルボックスからパーセルボックスをインポート
パーセルボックスから。ブリッジインポート ブリッジ
pydanticインポートBaseModelから
クラス座標 (BaseModel):
x : 浮動小数点
y : 浮動小数点
z : 浮動小数点数 = 0.0
クラスセンサー (ブリッジ):
def 温度 (self) -> float :
"""温度を摂氏で読み取ります。"""
リターン23.5
クラスロボット (ブリッジ):
def __init__ ( self ):
自分自身。センサー = センサー ()
def move ( self , to : 座標 ) -> dict :
"""ロボットを所定の位置に移動します。"""
return { "位置" : [ に . × 、 に 。 y 、へ 。 z ], "ステータス" : "到達" }
Parselbox ( context = { "robot" : Robot ()}) を sbx として非同期:
sbx を待ちます。 execute_code ( "robot.move(to={'x': 1, 'y': 2})" )
sbx を待ちます。 execute_code ( "robot.sensors.temperature()" )
Parselbox には、REST、GraphQL、およびシェル用のブリッジが同梱されています。
パーセルボックスからパーセルボックスをインポート
パーセルボックスから。ブリッジインポート HTTPBridge 、 GraphQLBridge 、 ShellBridge
API = HTTPBridge (
spec = "https://petstore3.swagger.io/api/v3/openapi.json" ,
Base_url = "https://petstore3.swagger.io/api/v3" ,
)
gql = GraphQLBridge ( "https://countries.trevorblades.com/graphql" )
sh = ShellBridge ( "ssh -T user@host" )
mcp = { "mcpServers" : { "deepwiki" : { "type" : "http" , "url" : "https://mcp.deepwiki.com/mcp" }}}
Parselbox との非同期 ( context = { "api" : api 、 "gql" : gql 、 "sh" : sh }、 mcp = mcp 、 network = True ) として sbx :
sbx を待ちます。実行コード ( 'api.search("GET /pet/*")' )
sbx を待ちます。実行コード ( 'api.get("/pet/1")' )
sbx を待ちます。 execute_code ( 'gql.graphql(query="{ 続き

inents { 名前 } }")' )
sbx を待ちます。 execute_code ( 'gql.graphql(query="{ 言語 { コード名 } }")' )
sbx を待ちます。 execute_code ( 'term = sh.shell.task()' )
sbx を待ちます。実行コード ( 'term.send("df -h")' )
sbx を待ちます。 execute_code ( "sbx.search('ask|read')" )
sbx を待ちます。 execute_code ( "deepwiki.read_wiki_struction(repoName='pyodide/pyodide')" )
sbx を待ちます。 execute_code ( "deepwiki.ask_question(question='ピオディドとは何ですか?', repoName='ピオディド/ピオディド')" )
実行可能ファイル: Bridges.py
すべてのコンテキストと MCP 呼び出しには、サンドボックスをブロックすることなくホスト上で実行される .task() フォームもあります。これは、並列ファンアウト、長時間実行ジョブ、対話型セッション用です。
ジョブ = sh 。実行。 task ( command = "ffmpeg -i in.mp4 out.mp4" ) # すぐにタスクを返します
仕事。 status () # TaskStatus(状態、経過時間、メッセージ、ログファイル)
仕事。 tail ( 5 ) # タスクのライブログの最後の行
仕事。 send ( "q" ) # 実行中の対話型プロセスにメッセージを送信します
ジョブを待ちます。 wait ( timeout = 120 ) # 完了するまでブロック — または単に `await job`
仕事。キャンセル（）
# 並列ファンアウト
非同期をインポートする
results = asyncio を待ちます。 Gather ( * [ api . get . task ( f"/items/ { i } " ) for i in range ( 5 )])
MCP ツールは進行状況をストリーミングし、通知をタスクのログファイルに記録します。カスタム Bridge メソッドは次のように出力します。

[切り捨てられた]

## Original Extract

An embeddable Python sandbox where AI agents call tools as code. Powered by Deno and Pyodide. - thesanjeetc/Parselbox

GitHub - thesanjeetc/Parselbox: An embeddable Python sandbox where AI agents call tools as code. Powered by Deno and Pyodide. · GitHub
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
thesanjeetc
/
Parselbox
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
7 Commits 7 Commits Folders and files
.claude/ skills/ parselbox .claude/ skills/ parselbox .github/ workflows .github/ workflows assets assets examples examples parselbox parselbox tests tests .dockerignore .dockerignore .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version CLAUDE.md CLAUDE.md Dockerfile Dockerfile LICENSE.md LICENSE.md Makefile Makefile README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Code. Filesystem. Context. Tools.
What if agents had one tool to rule them all?
Parselbox is an embeddable Python sandbox where AI agents call tools as code — MCP servers, APIs, and shells become native Python objects. Disk-backed workspace, packages, and networking built in; a single process powered by Deno and Pyodide .
Drop the Parselbox MCP alongside existing MCP server configurations. Agents instantly get a Python runtime, MCP tools as code, support for skills and a disk-backed workspace.
🔒 Secure Isolation
No containers, no VMs — just a single, lightweight Deno + Pyodide process (~160 MB). Deno permissions, memory caps, timeouts, network allowlists. Snapshot caching and crash recovery.
MCP servers, REST + OpenAPI, GraphQL, shell, functions and classes — all native Python objects. Stateful across calls. Pydantic auto-conversion. Credentials stay on the host.
Full CPython with js() interop — use JS packages as native Python. require() for npm, local TypeScript, and .wasm modules. Virtual bash() for shell. Auto-install packages on import.
require() any .wasm — library exports become Python methods, WASI programs become callable commands; drop one in bin/ to run it from bash() too. In-process, inherits the sandbox's mounts and permissions, installs nothing on the host.
Append .task() to any call — parallel fan-out with asyncio.gather , check progress, tail logs, drive interactive sessions with send() , await later.
Disk-backed workspace — host mounts ( ro / rw ), input files at /files/ , outputs persisted to real directories. New and modified files are detected and returned per call.
help() , search() , inspect() , preview() — agents discover only what they need, when they need it.
display() renders HTML inline in the chat (MCP Apps), with Tailwind + daisyUI injected. Or serve a full app — built-in HTTP server with static files, live reload, file upload, and @api routes that compose across tools.
Parselbox uses Deno for the secure sandbox runtime.
# macOS / Linux
curl -fsSL https://deno.land/install.sh | sh
# Windows (PowerShell)
irm https://deno.land/install.ps1 | iex
2. Install Parselbox
pip install parselbox
Parselbox API
Wire any tool into the sandbox — MCP servers, REST/GraphQL, shells, host objects — and the agent calls them as native Python, composing them with real control flow over a disk-backed workspace and both the Python and npm package ecosystems.
import asyncio
import os
from textwrap import dedent
from parselbox import Parselbox
from parselbox . bridge import HTTPBridge , ShellBridge
class Analytics :
def summarize ( self , repos : list ) -> dict :
"""Aggregate repo stats."""
stars = [ r [ "stars" ] for r in repos ]
return { "count" : len ( repos ), "avg_stars" : round ( sum ( stars ) / len ( stars ))}
config = { "mcpServers" : { "playwright" : { "command" : "npx" , "args" : [ "@playwright/mcp@latest" ]}}}
async def main ():
async with Parselbox (
mcp = config ,
context = {
"analytics" : Analytics (),
"github" : HTTPBridge ( base_url = "https://api.github.com" , token = os . environ [ "GITHUB_TOKEN" ]),
"sh" : ShellBridge ( "bash" ),
},
network = True ,
allow_runtime_packages = True ,
packages = [ "numpy" , "npm:lodash" ],
output_dir = "./workspace" ,
) as sbx :
# Discover available tools
await sbx . execute_code ( "sbx.search('navigate|get')" )
# Scrape Hacker News for GitHub links in a real browser
await sbx . execute_code ( dedent ( """
import re
playwright.browser_navigate(url="https://news.ycombinator.com")
text = playwright.browser_snapshot()
repos = re.findall(r'github \\ .com/([ \\ w.-]+/[ \\ w.-]+)', text)[:5]
""" ))
# Fetch star counts in parallel, then summarize via the context bridge
await sbx . execute_code ( dedent ( """
import asyncio
results = await asyncio.gather(*[github.get.task(f"/repos/{r}") for r in repos])
repo_data = [{"name": r["data"]["name"], "stars": r["data"]["stargazers_count"]}
for r in results if r.get("ok")]
analytics.summarize(repo_data)
""" ))
# Chart it — matplotlib auto-installs on import
result = await sbx . execute_code ( dedent ( """
import matplotlib.pyplot as plt
plt.barh([r["name"] for r in repo_data], [r["stars"] for r in repo_data])
plt.savefig("chart.png")
""" ))
print ( result . files ) # ['chart.png']
image = sbx . read_file ( "chart.png" )
# every result carries .output, .files, .stdout, .stderr, .error
# Serve the whole sandbox as an MCP server
await sbx . r
[truncated]
The Parselbox CLI runs a standalone MCP server — every sandbox option is available as a flag.
Add the Parselbox MCP alongside your existing MCP servers.
Point --mcp at that same config file.
On startup, Parselbox connects to the other servers, exposes their tools inside the sandbox, and starts its own MCP server.
Don't worry — Parselbox detects and avoids connecting to itself. No infinite loops of doom.
{
"mcpServers" : {
"github" : {},
"linear" : {},
"parselbox" : {
"command" : " uvx " ,
"args" : [ " parselbox " , " --mcp " , " /absolute/path/to/mcp.json " ]
}
}
}
HTTP
uvx parselbox --mcp mcp.json --transport http --port 9000
{
"mcpServers" : {
"parselbox" : {
"type" : " http " ,
"url" : " http://localhost:9000/mcp "
}
}
}
Full Example
uvx parselbox \
--mcp ./mcp.json \
--transport http \
--host 0.0.0.0 \
--port 8080 \
--file hello.txt \
--mount ./datasets:/data:rw \
--output-dir ./outputs \
--packages pandas,matplotlib \
--package-dir ./cache \
--allow-runtime-packages \
--network \
--serve 3000 \
--memory 2048 \
--timeout 60 \
--env MY_API_KEY=...
Parselbox Agents
import asyncio
from parselbox import Parselbox
from agents import Agent , Runner , function_tool
sandbox = Parselbox (
mcp = { "mcpServers" : { "playwright" : { "command" : "npx" , "args" : [ "@playwright/mcp@latest" ]}}},
output_dir = "./outputs" ,
allow_runtime_packages = True ,
)
agent = Agent (
name = "Research Assistant" ,
model = "gpt-5.5" ,
instructions = f"You are a world-class research assistant. \n \n { sandbox . get_prompt () } " ,
tools = [ function_tool ( sandbox . get_tool ())],
)
async def main ():
async with sandbox :
result = await Runner . run (
agent ,
"Scrape Wikipedia's 'List of highest-grossing films' with the Playwright MCP. "
"Plot a bar chart of the top 10 and save it as ./plot.png" ,
max_turns = 30 ,
)
print ( result . final_output )
asyncio . run ( main ())
User Guide
The context bridge exposes host Python objects inside the sandbox:
context — functions and namespaces as callable tools. Execution pauses, runs on host, returns result.
globals — static values (strings, numbers, dicts) copied into the sandbox.
mcp — MCP server config (dict or path). Appears as callable namespaces inside sandbox.
Plain classes are auto-wrapped — every public method becomes a callable tool; methods starting with _ stay private:
from parselbox import Parselbox
class Calculator :
def add ( self , a : float , b : float ) -> float :
"""Add two numbers."""
return a + b
async with Parselbox ( context = { "calc" : Calculator ()}) as sbx :
await sbx . execute_code ( "calc.add(a=10, b=20)" )
Subclass Bridge for nested namespaces (auto-crawled); annotate a parameter with a Pydantic model and passed dicts convert to it automatically:
from parselbox import Parselbox
from parselbox . bridge import Bridge
from pydantic import BaseModel
class Coordinate ( BaseModel ):
x : float
y : float
z : float = 0.0
class Sensors ( Bridge ):
def temperature ( self ) -> float :
"""Read temperature in celsius."""
return 23.5
class Robot ( Bridge ):
def __init__ ( self ):
self . sensors = Sensors ()
def move ( self , to : Coordinate ) -> dict :
"""Move robot to a position."""
return { "position" : [ to . x , to . y , to . z ], "status" : "reached" }
async with Parselbox ( context = { "robot" : Robot ()}) as sbx :
await sbx . execute_code ( "robot.move(to={'x': 1, 'y': 2})" )
await sbx . execute_code ( "robot.sensors.temperature()" )
Parselbox ships bridges for REST, GraphQL, and shell:
from parselbox import Parselbox
from parselbox . bridge import HTTPBridge , GraphQLBridge , ShellBridge
api = HTTPBridge (
spec = "https://petstore3.swagger.io/api/v3/openapi.json" ,
base_url = "https://petstore3.swagger.io/api/v3" ,
)
gql = GraphQLBridge ( "https://countries.trevorblades.com/graphql" )
sh = ShellBridge ( "ssh -T user@host" )
mcp = { "mcpServers" : { "deepwiki" : { "type" : "http" , "url" : "https://mcp.deepwiki.com/mcp" }}}
async with Parselbox ( context = { "api" : api , "gql" : gql , "sh" : sh }, mcp = mcp , network = True ) as sbx :
await sbx . execute_code ( 'api.search("GET /pet/*")' )
await sbx . execute_code ( 'api.get("/pet/1")' )
await sbx . execute_code ( 'gql.graphql(query="{ continents { name } }")' )
await sbx . execute_code ( 'gql.graphql(query="{ languages { code name } }")' )
await sbx . execute_code ( 'term = sh.shell.task()' )
await sbx . execute_code ( 'term.send("df -h")' )
await sbx . execute_code ( "sbx.search('ask|read')" )
await sbx . execute_code ( "deepwiki.read_wiki_structure(repoName='pyodide/pyodide')" )
await sbx . execute_code ( "deepwiki.ask_question(question='What is Pyodide?', repoName='pyodide/pyodide')" )
Runnable: bridges.py
Every context and MCP call also has a .task() form that runs on the host without blocking the sandbox — for parallel fan-out, long-running jobs, and interactive sessions:
job = sh . exec . task ( command = "ffmpeg -i in.mp4 out.mp4" ) # returns a task immediately
job . status () # TaskStatus(state, elapsed, message, logfile)
job . tail ( 5 ) # last lines of the task's live log
job . send ( "q" ) # message a running interactive process
await job . wait ( timeout = 120 ) # block until done — or just `await job`
job . cancel ()
# parallel fan-out
import asyncio
results = await asyncio . gather ( * [ api . get . task ( f"/items/ { i } " ) for i in range ( 5 )])
MCP tools stream their progress and log notifications into the task's logfile. A custom Bridge method emits

[truncated]
