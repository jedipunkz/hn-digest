---
source: "https://www.devclubhouse.com/a/ship-an-mcp-server-in-python-that-exposes-your-internal-api-to-llms"
hn_url: "https://news.ycombinator.com/item?id=48594116"
title: "Tutorial: Python MCP Internal API > LLM"
article_title: "Ship an MCP Server in Python That Exposes Your Internal API to LLMs — DevClubHouse"
author: "marcpope1"
captured_at: "2026-06-19T04:36:57Z"
capture_tool: "hn-digest"
hn_id: 48594116
score: 2
comments: 0
posted_at: "2026-06-19T02:14:10Z"
tags:
  - hacker-news
  - translated
---

# Tutorial: Python MCP Internal API > LLM

- HN: [48594116](https://news.ycombinator.com/item?id=48594116)
- Source: [www.devclubhouse.com](https://www.devclubhouse.com/a/ship-an-mcp-server-in-python-that-exposes-your-internal-api-to-llms)
- Score: 2
- Comments: 0
- Posted: 2026-06-19T02:14:10Z

## Translation

タイトル: チュートリアル: Python MCP 内部 API > LLM
記事のタイトル: 内部 API を LLM に公開する Python で MCP サーバーを出荷する — DevClubHouse
説明: FastMCP を使用して企業 REST API を 3 つの型付きツールでラップし、それらをローカルで検査して、モデルに資格情報を公開することなく Claude Desktop に接続します。

記事本文:
コンテンツにスキップ
8インチ
class="スティッキートップ-0 z-40 トランジションカラー期間-300"
:class="スクロール? 'bg-ink-950/85 backdrop-blur-xl border-b border-ink-700' : 'bg-transparent border-b border-transparent'"
>
開発クラブハウス
開発者のニュース、チュートリアル、AI に関する洞察
ホーム
AI
開発ツール
フレームワーク
クラウドとインフラ
セキュリティ
チュートリアル
$refs.q.focus())"
@keydown.window.ctrl.k.prevent="open = true; $nextTick(() => $refs.q.focus())"
@keydown.escape.window="open = false">
$refs.q.focus())"
class="grid place-items-center size-9rounded-lg border border-ink-700 text-mist-300 hover:text-mist-100 hover:border-ink-600transition-colors" aria-label="検索">
検索
サインイン
サインアップ
ホーム
AI
開発ツール
フレームワーク
クラウドとインフラ
セキュリティ
チュートリアル
サインイン
アカウントを作成する
AI
上級者向け
チュートリアル
内部 API を LLM に公開する MCP サーバーを Python で出荷する
FastMCP を使用して企業 REST API を 3 つの型付きツールでラップし、それらをローカルで検査して、認証情報をモデルに公開することなく Claude Desktop に接続します。
FastMCP を使用した Python MCP サーバー。企業 REST API を 3 つの構造化ツール ( search_customers 、 get_order 、および create_support_ticket ) としてラップします。 MCP 互換クライアント (Claude Desktop、Cursor、カスタム エージェント) は、モデルが資格情報を参照したり、生の URL を構築したりすることなく、完全なタイプ セーフティで API を呼び出すことができます。
Python 3.10+ ( list[dict] のような組み込みジェネリック型に必要)
パッケージ管理用の pip または uv
Node.js 18+ — mcp dev は内部で npx @modelcontextprotocol/inspector を呼び出します
最新の Claude Desktop (エンドツーエンドのテスト用。インスペクターのみを使用する場合はオプション)
ベアラー トークンを使用した REST API — 模擬 URL は従うのに問題なく機能します
async / await Python を快適に使用する
mkdir mcp-internal-api && cd mcp-internal-api
Python -m venv .venv
ソース

e .venv/bin/activate # Windows: .venv\Scripts\activate
pip install "mcp[cli]" httpx python-dotenv
mcp[cli] は、ローカル検査に使用される mcp CLI をインストールします。 httpx はバックエンドへの非同期 HTTP を処理します。
ローカル認証情報の .env を作成します。今すぐ .gitignore に追加します。
API_BASE_URL=https://api.corp.example.com
API_KEY=sk-your-real-token-here
2. サーバーを作成する
OSをインポートする
httpx をインポートする
from dotenv import load_dotenv
mcp.server.fastmcp から FastMCP をインポート
ロード_dotenv()
mcp = FastMCP("内部 API")
_BASE = os.environ["API_BASE_URL"]
_KEY = os.environ["API_KEY"]
def _auth_headers() -> dict[str, str]:
return {"認可": f"ベアラー {_KEY}", "受け入れ": "application/json"}
@mcp.tool()
async def search_customers(クエリ: str、制限: int = 10) -> リスト[dict]:
"""顧客を名前または電子メールで検索します。顧客レコードのリストを返します。"""
httpx.AsyncClient() をクライアントとして非同期:
r = client.get( を待つ)
f"{_BASE}/顧客",
headers=_auth_headers()、
params={"q": クエリ、"limit": 制限}、
タイムアウト=10.0、
）
r.raise_for_status()
r.json() を返す
@mcp.tool()
async def get_order(order_id: str) -> dict:
"""ID を使用して 1 つの注文を取得します。"""
httpx.AsyncClient() をクライアントとして非同期:
r = client.get( を待つ)
f"{_BASE}/orders/{order_id}",
headers=_auth_headers()、
タイムアウト=10.0、
）
r.raise_for_status()
r.json() を返す
@mcp.tool()
async def create_support_ticket(
customer_id: str、
件名: str、
ボディ: str、
優先順位: str = "通常"、
) -> 辞書:
"""顧客のサポート チケットを開きます。
引数:
customer_id: 顧客の UUID。
件名: 1 行の概要 (最大 120 文字)。
本文: 問題の完全な説明。
優先順位: 「低」、「通常」、または「高」。
「」
優先度が {"low"、"normal"、"high"} でない場合:
raise ValueError(f"優先度は低/標準/高である必要があります。'{priority}' を取得しました")
httpx.AsyncClient() をクライアントとして非同期:
r = クライアントを待つ

.post(
f"{_BASE}/チケット",
headers=_auth_headers()、
json={
"顧客 ID": 顧客 ID、
「件名」: 件名、
「体」：体、
「優先度」: 優先度、
}、
タイムアウト=10.0、
）
r.raise_for_status()
r.json() を返す
__name__ == "__main__"の場合:
mcp.run()
それぞれの決定がなぜ重要なのか:
mcp.run() のデフォルトは stdio Transport です。これは Claude Desktop とほとんどのローカル クライアントが期待するものです。クライアントはサーバーをサブプロセスとして生成し、stdin/stdout 経由で JSON-RPC を通信します。
3. mcp dev を使用してローカルで検査する
LLM に触れる前に、ブラウザ UI で配線を検証します。
mcp 開発サーバー.py
これによりサーバーが起動し、MCP インスペクターが開きます (URL が端末に表示されます)。 [ツール] に移動します。Python 署名と一致する自動生成された入力フォームを持つ 3 つのツールがすべて表示されます。 query = "alice" を指定して search_customers を呼び出し、JSON 応答または入力されたアップストリーム エラーを確認します。
ヒント: ライブ内部 API を使用せずに非同期/認証プラミングを実行するには、API_BASE_URL=https://httpbin.org を一時的に設定します。 404 が返されますが、これは httpx.HTTPStatusError ツール エラーとして正しく表示されます。
サーバーエントリを追加します。絶対パスを使用する — Claude Desktop は、virtualenv をアクティブ化しないクリーンな非ログイン シェルを生成します。
{
"mcpサーバー": {
"内部 API": {
"コマンド": "/absolute/path/to/.venv/bin/python",
"args": ["/absolute/path/to/server.py"],
"環境": {
"API_BASE_URL": "https://api.corp.example.com",
"API_KEY": "sk-your-real-token-here"
}
}
}
}
Claude Desktop を再起動し、新しい会話で次のようにします。
「『smith』という名前の顧客を検索し、注文が遅れていることを説明する最初の結果に対して優先度の高いサポート チケットを開きます。」
クロードは、 search_customers を呼び出し、出力を検査してから、 create_support_ticket を呼び出します。ツール呼び出しは、引数と応答が表示された状態で UI にインラインで表示されます。
検査官：A

mcp dev server.py の後、[ツール] タブには、正しいスキーマを持ち、ターミナルにインポート エラーがない 3 つのツールがすべてリストされます。
クロード デスクトップ: [設定] → [開発者] を開きます。サーバーは内部 API として表示され、緑色の接続済みインジケーターが表示されます。エラー状態が表示される場合は、構成を編集した後に Claude Desktop を再起動します。
スキーマ健全性チェック — 完全なクライアントを起動せずに、FastMCP が生成した正しいスキーマを確認します。
Python -c "
OSをインポートします。 os.environ['API_BASE_URL']='http://x'; os.environ['API_KEY']='x'
非同期をインポートする
インポートサーバー
非同期def main():
await server.mcp.list_tools() のツールの場合:
print(tool.name, tools.inputSchema)
asyncio.run(main())
」
各ツール名とその JSON スキーマ inputSchema dict が表示されるはずです。
ModuleNotFoundError: 'mcp' という名前のモジュールがありません — Claude Desktop はクリーン シェルを使用します。 virtualenv がアクティブ化されていません。 「command」がシステム python ではなく、venv インタープリター /path/to/.venv/bin/python を指していることを確認します。
Claude Desktop にツールが表示されない — 最初に mcp dev server.py を実行します。インポート エラーや環境変数の欠落がすぐに表示されます。 macOS では ~/Library/Logs/Claude/ も確認してください。Claude Desktop は、サーバー キー ( Internal-api ) にちなんで名付けられたサーバーごとのログ ファイルを書き込みます。
KeyError: 'API_BASE_URL' — claude_desktop_config.json の env ブロックはシェル環境を完全に置き換えます。 load_dotenv() はそこから .env を読み取りません。必要なすべてのキーを JSON 構成で明示的に設定します。
httpx.ReadTimeout — バックエンドが遅いです。 timeout=30.0 を増やすか、長時間実行操作を再構築して AsyncGenerator を返し、yield を使用して部分的な結果をクライアントにストリーミングして返します。
リソース: @mcp.resource() 経由で読み取り専用コンテキスト (OpenAPI 仕様、内部 Wiki) を公開することで、LLM がツール呼び出しバジェットを消費せずに参照資料を取得できるようにします。
HTTP/SSE トランスポート: mul の場合

ti-user またはリモートのデプロイメントでは、mcp.run() を mcp.run(transport="sse") に置き換え、安全なリバース プロキシの背後にマウントします。静的な環境変数ではなくミドルウェアでリクエストごとのトークンを検証します。
レート制限: エージェント ループによるアップストリーム API のフラッディングを防ぐために、_auth_headers() をトークン バケット リミッター (aiolimiter は非同期ネイティブ) でラップします。
より豊富なスキーマ: dict の戻り値の型を Pydantic モデルに置き換えます。FastMCP はそれらから詳細な JSON スキーマを生成し、どのフィールドを予期して使用するかについてモデルに適切なガイダンスを提供します。
仕様と SDK:modelcontextprotocol.io/docs および GitHub の Python SDK。
Mariana は、機械学習と生成 AI の急速に変化する世界をカバーしており、特にこれらのテクノロジーが開発ワークフローをどのように再構築しているかに焦点を当てています。最新の基礎モデルのストレス テストを行っていないときは、通常、地元のハッカソンに参加しています。
コメントして投票するには、サインインするかアカウントを作成してください。
1
ニューラル ネットワーク推論のための整数量子化の謎を解く
5時間前
·
0
2
トークン圧縮の幻想: CLI 切り捨ての隠れたコスト
6時間前
·
0
3
Elasticsearch での永続的な LLM エージェント メモリの設計
11時間前
·
0
4
Kilo Code はオープンソースのエージェント エンジニアリングを IDE に導入します
15時間前
·
2
1
内部 API を LLM に公開する MCP サーバーを Python で出荷する
5日前
·
0
2
npm v12 はインストール スクリプトの実行を停止しようとしています — 監査すべき内容は次のとおりです
1週間前
·
5
3
サプライチェーンのマルウェアが AI プログラマーの資格情報を標的にした後、Microsoft が数十の GitHub リポジトリを取得
1週間前
·
5
4
Homebrew 6.0.0 はタップを信頼できるようにします
1週間前
·
7
5
米国政府、AnthropicにFable 5とMythos 5の削除を強制
6日前
·
5
ニューラル ネットワーク推論のための整数量子化の謎を解く
0
AI
トークン圧縮の幻想: CLI 切り捨ての隠れたコスト
0
AI
永続的な LLM 年齢の設計

Elasticsearch の nt メモリ
0
AI
Kilo Code はオープンソースのエージェント エンジニアリングを IDE に導入します
2
devclubhouse.com
開発者ニュース、リリース、ヒント、チュートリアル - 継続的に更新されます。
最高の開発者および AI コンテンツが受信トレイに配信されます。スパムはありません。
© 2026 devclubhouse.com · 無断複写・転載を禁じます
すべての製品名、ロゴ、ブランドはそれぞれの所有者の財産であり、次の用途に使用されます。
身元確認と編集上のコメントのみ。これらの使用は、承認や提携を意味するものではありません。
商標権の侵害を目的としたものではありません。
私たちは、読まれている内容を理解し、広告を表示し続けるためにいくつかの Cookie を使用します。
断っていただければ、絶対に必要なものだけを守ります。
プライバシーポリシー

## Original Extract

Wrap a corporate REST API in three typed tools using FastMCP, inspect them locally, and connect them to Claude Desktop—without ever exposing credentials to the model.

Skip to content
8"
class="sticky top-0 z-40 transition-colors duration-300"
:class="scrolled ? 'bg-ink-950/85 backdrop-blur-xl border-b border-ink-700' : 'bg-transparent border-b border-transparent'"
>
devclubhouse
developer news, tutorials & AI insight
Home
AI
Dev Tools
Frameworks
Cloud & Infra
Security
Tutorials
$refs.q.focus())"
@keydown.window.ctrl.k.prevent="open = true; $nextTick(() => $refs.q.focus())"
@keydown.escape.window="open = false">
$refs.q.focus())"
class="grid place-items-center size-9 rounded-lg border border-ink-700 text-mist-300 hover:text-mist-100 hover:border-ink-600 transition-colors" aria-label="Search">
Search
Sign in
Sign up
Home
AI
Dev Tools
Frameworks
Cloud & Infra
Security
Tutorials
Sign in
Create an account
AI
Advanced
Tutorial
Ship an MCP Server in Python That Exposes Your Internal API to LLMs
Wrap a corporate REST API in three typed tools using FastMCP, inspect them locally, and connect them to Claude Desktop—without ever exposing credentials to the model.
A Python MCP server using FastMCP that wraps a corporate REST API as three structured tools— search_customers , get_order , and create_support_ticket . Any MCP-compatible client (Claude Desktop, Cursor, custom agents) can call your API with full type safety, without the model ever seeing credentials or constructing raw URLs.
Python 3.10+ (required for built-in generic types like list[dict] )
pip or uv for package management
Node.js 18+ — mcp dev invokes npx @modelcontextprotocol/inspector under the hood
Latest Claude Desktop (for end-to-end testing; optional if using only the inspector)
A REST API with a bearer token — a mock URL works fine to follow along
Comfortable with async / await Python
mkdir mcp-internal-api && cd mcp-internal-api
python -m venv .venv
source .venv/bin/activate # Windows: .venv\Scripts\activate
pip install "mcp[cli]" httpx python-dotenv
mcp[cli] installs the mcp CLI used for local inspection. httpx handles async HTTP to your backend.
Create .env for local credentials — add it to .gitignore now :
API_BASE_URL=https://api.corp.example.com
API_KEY=sk-your-real-token-here
2. Write the Server
import os
import httpx
from dotenv import load_dotenv
from mcp.server.fastmcp import FastMCP
load_dotenv()
mcp = FastMCP("internal-api")
_BASE = os.environ["API_BASE_URL"]
_KEY = os.environ["API_KEY"]
def _auth_headers() -> dict[str, str]:
return {"Authorization": f"Bearer {_KEY}", "Accept": "application/json"}
@mcp.tool()
async def search_customers(query: str, limit: int = 10) -> list[dict]:
"""Search customers by name or email. Returns a list of customer records."""
async with httpx.AsyncClient() as client:
r = await client.get(
f"{_BASE}/customers",
headers=_auth_headers(),
params={"q": query, "limit": limit},
timeout=10.0,
)
r.raise_for_status()
return r.json()
@mcp.tool()
async def get_order(order_id: str) -> dict:
"""Fetch a single order by its ID."""
async with httpx.AsyncClient() as client:
r = await client.get(
f"{_BASE}/orders/{order_id}",
headers=_auth_headers(),
timeout=10.0,
)
r.raise_for_status()
return r.json()
@mcp.tool()
async def create_support_ticket(
customer_id: str,
subject: str,
body: str,
priority: str = "normal",
) -> dict:
"""Open a support ticket for a customer.
Args:
customer_id: The customer's UUID.
subject: One-line summary (max 120 chars).
body: Full description of the issue.
priority: 'low', 'normal', or 'high'.
"""
if priority not in {"low", "normal", "high"}:
raise ValueError(f"priority must be low/normal/high, got '{priority}'")
async with httpx.AsyncClient() as client:
r = await client.post(
f"{_BASE}/tickets",
headers=_auth_headers(),
json={
"customer_id": customer_id,
"subject": subject,
"body": body,
"priority": priority,
},
timeout=10.0,
)
r.raise_for_status()
return r.json()
if __name__ == "__main__":
mcp.run()
Why each decision matters:
mcp.run() defaults to stdio transport , which is what Claude Desktop and most local clients expect — the client spawns your server as a subprocess and talks JSON-RPC over stdin/stdout.
3. Inspect Locally with mcp dev
Before touching any LLM, validate the wiring in a browser UI:
mcp dev server.py
This starts your server and opens the MCP Inspector (the URL is printed in your terminal). Navigate to Tools — you'll see all three tools with auto-generated input forms matching your Python signatures. Call search_customers with query = "alice" and confirm a JSON response or a typed upstream error.
Tip: Set API_BASE_URL=https://httpbin.org temporarily to exercise the async/auth plumbing without a live internal API. You'll get a 404 back, which correctly surfaces as an httpx.HTTPStatusError tool error.
Add your server entry. Use absolute paths — Claude Desktop spawns a clean, non-login shell that won't activate your virtualenv:
{
"mcpServers": {
"internal-api": {
"command": "/absolute/path/to/.venv/bin/python",
"args": ["/absolute/path/to/server.py"],
"env": {
"API_BASE_URL": "https://api.corp.example.com",
"API_KEY": "sk-your-real-token-here"
}
}
}
}
Restart Claude Desktop, then in a new conversation:
"Search for customers named 'smith', then open a high-priority support ticket for the first result explaining their order is delayed."
Claude will call search_customers , inspect the output, then call create_support_ticket — tool calls appear inline in the UI with their arguments and responses visible.
Inspector: After mcp dev server.py , the Tools tab lists all three tools with correct schemas and no import errors in the terminal.
Claude Desktop: Open Settings → Developer . Your server appears as internal-api with a green connected indicator. If it shows an error state, restart Claude Desktop after editing the config.
Schema sanity-check — confirm FastMCP generated correct schemas without starting a full client:
python -c "
import os; os.environ['API_BASE_URL']='http://x'; os.environ['API_KEY']='x'
import asyncio
import server
async def main():
for tool in await server.mcp.list_tools():
print(tool.name, tool.inputSchema)
asyncio.run(main())
"
You should see each tool name alongside its JSON Schema inputSchema dict.
ModuleNotFoundError: No module named 'mcp' — Claude Desktop uses a clean shell; your virtualenv isn't activated. Confirm "command" points to the venv interpreter: /path/to/.venv/bin/python , not the system python .
Tools don't appear in Claude Desktop — Run mcp dev server.py first; import errors or missing env vars appear there immediately. Also check ~/Library/Logs/Claude/ on macOS — Claude Desktop writes a per-server log file named after your server key ( internal-api ).
KeyError: 'API_BASE_URL' — The env block in claude_desktop_config.json replaces the shell environment entirely; load_dotenv() won't read your .env from there. Set all required keys explicitly in the JSON config.
httpx.ReadTimeout — Your backend is slow. Raise timeout=30.0 , or restructure long-running operations to return an AsyncGenerator and use yield to stream partial results back to the client.
Resources: Expose read-only context (OpenAPI specs, internal wikis) via @mcp.resource() so the LLM can pull reference material without consuming tool-call budget.
HTTP/SSE transport: For multi-user or remote deployments, replace mcp.run() with mcp.run(transport="sse") and mount it behind a secured reverse proxy; validate per-request tokens in middleware rather than a static env var.
Rate limiting: Wrap _auth_headers() with a token-bucket limiter ( aiolimiter is async-native) to prevent an agentic loop from flooding your upstream API.
Richer schemas: Replace dict return types with Pydantic models — FastMCP generates detailed JSON Schema from them, giving the model better guidance on what fields to expect and use.
Spec & SDK: modelcontextprotocol.io/docs and the Python SDK on GitHub .
Mariana covers the fast-moving world of machine learning and generative AI, with a particular focus on how these technologies are reshaping development workflows. When she isn't stress-testing the latest foundation models, she's usually at a local hackathon.
Sign in or create an account to comment and vote.
1
Demystifying Integer Quantization for Neural Network Inference
5h ago
·
0
2
The Token Compression Illusion: The Hidden Cost of CLI Truncation
6h ago
·
0
3
Designing Persistent LLM Agent Memory on Elasticsearch
11h ago
·
0
4
Kilo Code Brings Open-Source Agentic Engineering to Your IDE
15h ago
·
2
1
Ship an MCP Server in Python That Exposes Your Internal API to LLMs
5d ago
·
0
2
npm v12 Is About to Stop Running Your Install Scripts — Here's What to Audit
1w ago
·
5
3
Microsoft Pulls Dozens of GitHub Repos After Supply-Chain Malware Targets AI Coders' Credentials
1w ago
·
5
4
Homebrew 6.0.0 Makes You Trust Your Taps
1w ago
·
7
5
US Government Forces Anthropic to Pull Fable 5 and Mythos 5
6d ago
·
5
Demystifying Integer Quantization for Neural Network Inference
0
AI
The Token Compression Illusion: The Hidden Cost of CLI Truncation
0
AI
Designing Persistent LLM Agent Memory on Elasticsearch
0
AI
Kilo Code Brings Open-Source Agentic Engineering to Your IDE
2
devclubhouse.com
Developer news, releases, tips and tutorials — refreshed continuously.
The best developer & AI content, delivered to your inbox. No spam.
© 2026 devclubhouse.com · All rights reserved
All product names, logos, and brands are property of their respective owners and are used for
identification and editorial commentary only. Their use does not imply endorsement or affiliation.
No trademark infringement is intended.
We use a few cookies to understand what's being read and keep the lights on with ads.
Decline and we'll stick to the strictly necessary ones.
Privacy policy
