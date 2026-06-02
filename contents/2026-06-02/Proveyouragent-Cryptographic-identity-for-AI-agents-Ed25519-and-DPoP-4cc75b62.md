---
source: "https://github.com/lujainkhalil/proveyouragent"
hn_url: "https://news.ycombinator.com/item?id=48354556"
title: "Proveyouragent: Cryptographic identity for AI agents (Ed25519 and DPoP)"
article_title: "GitHub - lujainkhalil/proveyouragent: Cryptographic identity for AI agents: Ed25519 keypairs, DPoP request signing, delegation chains. · GitHub"
author: "lujainkhalil"
captured_at: "2026-06-02T00:43:05Z"
capture_tool: "hn-digest"
hn_id: 48354556
score: 1
comments: 0
posted_at: "2026-06-01T09:33:55Z"
tags:
  - hacker-news
  - translated
---

# Proveyouragent: Cryptographic identity for AI agents (Ed25519 and DPoP)

- HN: [48354556](https://news.ycombinator.com/item?id=48354556)
- Source: [github.com](https://github.com/lujainkhalil/proveyouragent)
- Score: 1
- Comments: 0
- Posted: 2026-06-01T09:33:55Z

## Translation

タイトル: Proveyouragent: AI エージェントの暗号化アイデンティティ (Ed25519 および DPoP)
記事のタイトル: GitHub - lujainkhalil/proveyouragent: AI エージェントの暗号化 ID: Ed25519 キーペア、DPoP リクエスト署名、委任チェーン。 · GitHub
説明: AI エージェントの暗号化 ID: Ed25519 キーペア、DPoP リクエスト署名、委任チェーン。 - ルジャインハリル/proveyouragent

記事本文:
GitHub - lujainkhalil/proveyouragent: AI エージェントの暗号化アイデンティティ: Ed25519 キーペア、DPoP リクエスト署名、委任チェーン。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
ルジャインカリル
/
あなたの代理人を証明する
公共
通知
あなたはそうしなければなりません

サインインして通知設定を変更します
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミットの例 例proveyouragent giveyouragent testing testing .DS_Store .DS_Store .gitignore .gitignore README.md README.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントの暗号化 ID。
proveyouragent は、各エージェントにキーペア、署名付き身分証明書、およびすべてのリクエストに対してそのリクエストが実際にそのエージェントからのものであることを証明する方法を提供します。サービスは、エージェントのリクエストを処理する前に検証します。盗まれたトークンは秘密鍵がなければ役に立ちません。
Ed25519、OAuth 2.0 動的クライアント登録 (RFC 7591)、および DPoP (RFC 9449) に基づいて構築されています。
AI エージェントは API を呼び出し、データベースを読み取り、ファイルを書き込み、電子メールを送信します。ほとんどは、ハードコーディングされたサービス アカウント トークンまたは借用したユーザー資格情報を使用してこれを行います。サービスが以下を認識する標準的な方法はありません。
誰がそのエージェントを所有し、そのエージェントに対して責任を負うのか
エージェントが実際に実行できること
リクエストが盗まれたトークンから再実行されたかどうか
proveyouragent は、一緒に構成する小さなプリミティブのセットを使用してこれを解決します。
pip installproveyouragent
クイックスタート
からproveyouragentインポートgenerate_keypair、save_keypair、create_software_statement
キー = 生成_キーペア()
save_keypair (キー)
ステートメント = create_software_statement (
private_key = キー 、
オペレーター_ドメイン = "acme.com" ,
エージェント名 = "請求エージェント" ,
エージェントバージョン = "1.0.0" 、
スコープ = [ "請求書:読み取り" , "支払い:書き込み" ],
)
すべてのリクエストに署名します。
からproveyouragentインポートcreate_dpop_proof
proof = create_dpop_proof ( key 、メソッド = "GET" 、uri = "https://api.acme.com/invoices" )
応答 = httpx 。取得（
"https://api.acme.com/invoices" 、
ヘッダー = {
"X-Agent-Statement" : 状態

メント、
「X-Agent-DPoP」: 証明 、
}
)
サーバー上で確認します。
from fastapi import FastAPI 、リクエスト
証明者から。ミドルウェアのインポート AgentIDMiddleware 、verify_agent
アプリ = FastAPI ()
アプリ。 add_middleware ( AgentIDMiddleware 、 get_public_key = my_key_resolver )
@アプリ。 get ( "/請求書" )
def list_invoices (リクエスト: リクエスト):
エージェント = verify_agent ( request , required_scope = "請求書:読み取り" )
return { "エージェント" : エージェント .エージェント名 , "請求書" : [...]}
仕組み
すべてのエージェントは Ed25519 キーペアを取得します。秘密キーがエージェントから離れることはありません。公開キーは既知の URL で公開されるため、どのサービスでもコール ホームすることなくリクエストを検証できます。
エージェントの ID ドキュメントは、ソフトウェア ステートメントと呼ばれる署名付き JWT です。これは、エージェントの所有者、エージェントに何が許可されているか、および公開キーを見つける場所を宣言します。
ステートメント = create_software_statement (
private_key = キー 、
Operator_domain = "acme.com" , # このエージェントの責任者は誰ですか
エージェント名 = "請求エージェント" ,
エージェントバージョン = "1.0.0" 、
スコープ = [ "請求書:読み取り" ],
model = "claude-sonnet-4-6" 、 # オプション
プロンプト_ハッシュ = "sha256:abc123" 、 # オプション、バージョン追跡用
)
DPoP による署名のリクエスト
無記名トークンは盗まれて再使用される可能性があります。 DPoP (RFC 9449) は、各トークンをエージェントの秘密キーにバインドします。すべてのリクエストには、HTTP メソッドと URI をカバーする、キーによって署名された新しい証明が含まれています。盗まれたトークンは、秘密鍵がなければ役に立ちません。
proof = create_dpop_proof (
private_key = キー 、
メソッド = "GET" 、
uri = "https://api.acme.com/invoices" ,
)
検証
このサービスは、リクエストごとに次の 4 つのことをチェックします。
ソフトウェアステートメントの署名は有効です
ソフトウェア明細書の有効期限が切れていない
エージェントには必要なスコープがある
DPoP プルーフは新しく、このリクエストと一致し、以前に使用されていません。
証拠を証明する人から

import verify_agent_request 、 VerifiedAgent 、 VerificationError
結果 = verify_agent_request (
software_statement = ステートメント ,
dpop_proof = 証明 、
メソッド = "GET" 、
uri = "https://api.acme.com/invoices" ,
オペレーター_パブリック_キー = パブリック_キー ,
required_scope = "請求書:読み取り" ,
)
if isinstance ( result , VerifiedAgent ):
print ( result . エージェント名 ) # billing-agent
print (結果 .operator_domain ) #acme.com
print ( result .scopes ) # ['invoices:read']
FastAPIミドルウェア
ミドルウェアはすべてのルートで検証を自動的に処理します。検証されたエージェントの詳細は request.state.agent に添付されます。
証明者から。ミドルウェアのインポート AgentIDMiddleware 、verify_agent
def get_public_key (operator_domain : str):
# この演算子の Ed25519PublicKey を返します
# データベース、構成、またはキーレジストリから取得します
your_key_store を返します。 get (オペレータードメイン)
アプリ。 add_middleware (
エージェントIDミドルウェア 、
get_public_key = get_public_key 、
exclude_paths = [ "/health" , "/docs" ],
)
@アプリ。 get ( "/請求書" )
def list_invoices (リクエスト: リクエスト):
エージェント = verify_agent ( request , required_scope = "請求書:読み取り" )
return { "請求書" : [...]}
委任チェーン
Orchestrator エージェントは、権限のサブセットをサブエージェントに委任できます。チェーンは暗号的にリンクされています。スコープはチェーンを通過するときにのみ縮小できます。
証明者から。委任インポート create_root_mandate 、 create_delegation 、 verify_delegation_chain
# 人間がオーケストレーターを承認する
root = create_root_mandate (
プライベートキー = オペレーターキー 、
オペレーター_ドメイン = "acme.com" ,
human_principal = "alice@acme.com" ,
スコープ = [ "請求書:読み取り" , "支払い:書き込み" ],
Agent_id = "acme.com/orchestrator" ,
)
# オーケストレーターはサブセットをサブエージェントに委任します
委任 = create_delegation (
delegator_key = オルシュ

trator_key 、
delegator_statement = Orchestrator_statement 、
delegate_agent_id = "acme.com/summariser" ,
delegate_public_key_b64 = summariser_pub_key 、
スコープ = [ "invoices:read" ], # 親スコープのサブセットのみ
親トークン = ルート 、
human_principal = "alice@acme.com" ,
)
# ツールは完全なチェーンを検証します
結果 = verify_delegation_chain (
トークン = 委任 、
required_scope = "請求書:読み取り" ,
get_public_key = キーリゾルバー 、
)
print (結果 . human_principal ) # alice@acme.com
print ( result . delegate_agent_id ) # acme.com/summariser
print (結果の深さ) #1
スコープのエスカレーションは直ちに拒否されます。
# これはトークンではなく DelegationError を返します
create_delegation (..., スコープ = [ "請求書:読み取り" , "管理者:削除" ])
# DelegationError: 親トークンに存在しないスコープは委任できません: {'admin:delete'}
リプレイキャッシュ
デフォルトのリプレイ キャッシュはメモリ内にあります。単一プロセスのデプロイメントでは機能しますが、再起動後は存続しないか、複数のプロセスにわたって機能しません。
証明者から。キャッシュインポート RedisCache
証明者から。ミドルウェアのインポート AgentIDMiddleware
アプリ。 add_middleware (
エージェントIDミドルウェア 、
get_public_key = get_public_key 、
キャッシュ = RedisCache ( URL = "redis://localhost:6379" )、
)
または、キャッシュを verify_agent_request に直接渡します。
証明者から。キャッシュインポート RedisCache
キャッシュ = RedisCache ( URL = "redis://localhost:6379" )
結果 = verify_agent_request (
...、
キャッシュ = キャッシュ、
)
すべてのリクエストで何が検証されるか
チェックする
何が引っかかるのか
ソフトウェアステートメントの署名
偽造または改ざんされた身分証明書
明細の有効期限
古いトークン
範囲の強制
付与されていない権限を主張するエージェント
DPoP 証明署名
キーホルダー本人以外のリクエスト
DPoP メソッドと URI バインディング
別のエンドポイントで再利用されたプルーフ
DPoPの鮮度
古いプルーフが再表示されます

肯定した
DPoP jti の独自性
キャプチャされたリクエストの正確なリプレイ
サンプルの実行
# ターミナル 1: サーバーを起動します
uvicorn Examples.server:app --reload
# ターミナル 2: クライアントを実行します
Python の例/client.py
テストの実行
pytest テスト/ -v
設計上の決定
Ed25519のみ。アルゴリズムのネゴシエーションはありません。 Ed25519 は高速で、キーが小さく、既知の弱点はありません。複数のアルゴリズムをサポートすると、複雑さが増し、攻撃対象領域が増加します。
ブロックチェーンも DID インフラストラクチャもありません。 DNS はトラストアンカーです。オペレーターは、ドメイン上の既知の URL で公開キーを公開します。すべての開発者は、DNS がどのように機能するかをすでに知っています。
例外ではなく、値としてのエラー。 verify_agent_request は VerifiedAgent または VerificationError を返します。通常の使用では Try/Except は必要ありません。エラーには常に人間が判読できる理由が含まれています。
リプレイキャッシュはプラグ可能です。デフォルトのメモリ内キャッシュは開発用に機能します。 Redis は本番環境で動作します。 ReplayCache を実装するバックエンドはどれでも機能します。
AI エージェントの暗号化 ID: Ed25519 キーペア、DPoP リクエスト署名、委任チェーン。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Cryptographic identity for AI agents: Ed25519 keypairs, DPoP request signing, delegation chains. - lujainkhalil/proveyouragent

GitHub - lujainkhalil/proveyouragent: Cryptographic identity for AI agents: Ed25519 keypairs, DPoP request signing, delegation chains. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
lujainkhalil
/
proveyouragent
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits examples examples proveyouragent proveyouragent tests tests .DS_Store .DS_Store .gitignore .gitignore README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
Cryptographic identity for AI agents.
proveyouragent gives each agent a keypair, a signed identity document, and a way to prove on every request that the request actually came from that agent. Services verify agent requests before processing them. Stolen tokens are useless without the private key.
Built on Ed25519, OAuth 2.0 Dynamic Client Registration (RFC 7591), and DPoP (RFC 9449).
AI agents call APIs, read databases, write files, and send emails. Most do this with a hardcoded service account token or a borrowed user credential. There is no standard way for a service to know:
who owns and is accountable for that agent
what the agent is actually allowed to do
whether the request was replayed from a stolen token
proveyouragent solves this with a small set of primitives that compose together.
pip install proveyouragent
Quick start
from proveyouragent import generate_keypair , save_keypair , create_software_statement
key = generate_keypair ()
save_keypair ( key )
statement = create_software_statement (
private_key = key ,
operator_domain = "acme.com" ,
agent_name = "billing-agent" ,
agent_version = "1.0.0" ,
scopes = [ "invoices:read" , "payments:write" ],
)
Sign every request:
from proveyouragent import create_dpop_proof
proof = create_dpop_proof ( key , method = "GET" , uri = "https://api.acme.com/invoices" )
response = httpx . get (
"https://api.acme.com/invoices" ,
headers = {
"X-Agent-Statement" : statement ,
"X-Agent-DPoP" : proof ,
}
)
Verify on the server:
from fastapi import FastAPI , Request
from proveyouragent . middleware import AgentIDMiddleware , verify_agent
app = FastAPI ()
app . add_middleware ( AgentIDMiddleware , get_public_key = my_key_resolver )
@ app . get ( "/invoices" )
def list_invoices ( request : Request ):
agent = verify_agent ( request , required_scope = "invoices:read" )
return { "agent" : agent . agent_name , "invoices" : [...]}
How it works
Every agent gets an Ed25519 keypair. The private key never leaves the agent. The public key is published at a well-known URL so any service can verify requests without calling home.
The agent's identity document is a signed JWT called a software statement. It declares who owns the agent, what the agent is allowed to do, and where to find the public key.
statement = create_software_statement (
private_key = key ,
operator_domain = "acme.com" , # who is accountable for this agent
agent_name = "billing-agent" ,
agent_version = "1.0.0" ,
scopes = [ "invoices:read" ],
model = "claude-sonnet-4-6" , # optional
prompt_hash = "sha256:abc123" , # optional, for version tracking
)
Request signing with DPoP
Bearer tokens can be stolen and replayed. DPoP (RFC 9449) binds each token to the agent's private key. Every request includes a fresh proof signed by the key, covering the HTTP method and URI. A stolen token is useless without the private key.
proof = create_dpop_proof (
private_key = key ,
method = "GET" ,
uri = "https://api.acme.com/invoices" ,
)
Verification
The service checks four things on every request:
The software statement signature is valid
The software statement has not expired
The agent has the required scope
The DPoP proof is fresh, matches this request, and has not been used before
from proveyouragent import verify_agent_request , VerifiedAgent , VerificationError
result = verify_agent_request (
software_statement = statement ,
dpop_proof = proof ,
method = "GET" ,
uri = "https://api.acme.com/invoices" ,
operator_public_key = public_key ,
required_scope = "invoices:read" ,
)
if isinstance ( result , VerifiedAgent ):
print ( result . agent_name ) # billing-agent
print ( result . operator_domain ) # acme.com
print ( result . scopes ) # ['invoices:read']
FastAPI middleware
The middleware handles verification automatically on every route. Verified agent details are attached to request.state.agent .
from proveyouragent . middleware import AgentIDMiddleware , verify_agent
def get_public_key ( operator_domain : str ):
# Return the Ed25519PublicKey for this operator
# Fetch from your database, config, or key registry
return your_key_store . get ( operator_domain )
app . add_middleware (
AgentIDMiddleware ,
get_public_key = get_public_key ,
exclude_paths = [ "/health" , "/docs" ],
)
@ app . get ( "/invoices" )
def list_invoices ( request : Request ):
agent = verify_agent ( request , required_scope = "invoices:read" )
return { "invoices" : [...]}
Delegation chains
Orchestrator agents can delegate a subset of their permissions to sub-agents. The chain is cryptographically linked. Scopes can only shrink as they pass down the chain.
from proveyouragent . delegation import create_root_mandate , create_delegation , verify_delegation_chain
# Human authorises orchestrator
root = create_root_mandate (
private_key = operator_key ,
operator_domain = "acme.com" ,
human_principal = "alice@acme.com" ,
scopes = [ "invoices:read" , "payments:write" ],
agent_id = "acme.com/orchestrator" ,
)
# Orchestrator delegates a subset to sub-agent
delegation = create_delegation (
delegator_key = orchestrator_key ,
delegator_statement = orchestrator_statement ,
delegate_agent_id = "acme.com/summariser" ,
delegate_public_key_b64 = summariser_pub_key ,
scopes = [ "invoices:read" ], # subset of parent scopes only
parent_token = root ,
human_principal = "alice@acme.com" ,
)
# Tool verifies the full chain
result = verify_delegation_chain (
token = delegation ,
required_scope = "invoices:read" ,
get_public_key = key_resolver ,
)
print ( result . human_principal ) # alice@acme.com
print ( result . delegate_agent_id ) # acme.com/summariser
print ( result . depth ) # 1
Scope escalation is rejected immediately:
# This returns a DelegationError, not a token
create_delegation (..., scopes = [ "invoices:read" , "admin:delete" ])
# DelegationError: Cannot delegate scopes not present in parent token: {'admin:delete'}
Replay cache
The default replay cache is in-memory. It works for single-process deployments but will not survive a restart or work across multiple processes.
from proveyouragent . cache import RedisCache
from proveyouragent . middleware import AgentIDMiddleware
app . add_middleware (
AgentIDMiddleware ,
get_public_key = get_public_key ,
cache = RedisCache ( url = "redis://localhost:6379" ),
)
Or pass a cache directly to verify_agent_request :
from proveyouragent . cache import RedisCache
cache = RedisCache ( url = "redis://localhost:6379" )
result = verify_agent_request (
...,
cache = cache ,
)
What gets verified on every request
Check
What it catches
Software statement signature
Forged or tampered identity documents
Statement expiry
Stale tokens
Scope enforcement
Agents claiming permissions they were not granted
DPoP proof signature
Requests not made by the key holder
DPoP method and URI binding
Proofs reused on a different endpoint
DPoP freshness
Old proofs being replayed
DPoP jti uniqueness
Exact replay of a captured request
Running the examples
# Terminal 1: start the server
uvicorn examples.server:app --reload
# Terminal 2: run the client
python examples/client.py
Running the tests
pytest tests/ -v
Design decisions
Ed25519 only. No algorithm negotiation. Ed25519 is fast, has small keys, and has no known weaknesses. Supporting multiple algorithms adds complexity and attack surface.
No blockchain, no DID infrastructure. DNS is the trust anchor. Operators publish their public key at a well-known URL on their domain. Every developer already knows how DNS works.
Errors as values, not exceptions. verify_agent_request returns a VerifiedAgent or a VerificationError . No try/except needed in normal usage. The error always includes a human-readable reason.
Replay cache is pluggable. The default in-memory cache works for development. Redis works for production. Any backend that implements ReplayCache works.
Cryptographic identity for AI agents: Ed25519 keypairs, DPoP request signing, delegation chains.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
