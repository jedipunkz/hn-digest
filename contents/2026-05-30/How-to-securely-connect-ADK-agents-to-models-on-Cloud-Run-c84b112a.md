---
source: "https://leoy.blog/posts/securely-connect-agents-to-models/"
hn_url: "https://news.ycombinator.com/item?id=48329469"
title: "How to securely connect ADK agents to models on Cloud Run"
article_title: "How to Securely Connect ADK Agents to Models on Cloud Run - minherz: another techno-blog"
author: "minherz"
captured_at: "2026-05-30T11:37:37Z"
capture_tool: "hn-digest"
hn_id: 48329469
score: 2
comments: 1
posted_at: "2026-05-29T21:23:30Z"
tags:
  - hacker-news
  - translated
---

# How to securely connect ADK agents to models on Cloud Run

- HN: [48329469](https://news.ycombinator.com/item?id=48329469)
- Source: [leoy.blog](https://leoy.blog/posts/securely-connect-agents-to-models/)
- Score: 2
- Comments: 1
- Posted: 2026-05-29T21:23:30Z

## Translation

タイトル: ADK エージェントを Cloud Run 上のモデルに安全に接続する方法
記事タイトル: ADK エージェントを Cloud Run 上のモデルに安全に接続する方法 - minherz: another techno-blog

記事本文:
minherz: 別のテクノブログ
クラウド、テクノロジー、生活、その他すべてについてのメモ
メニュー
ホーム
ADK エージェントを Cloud Run 上のモデルに安全に接続する方法
約。 1300単語
6分
ミンヘルツ
最終更新日: 2026 年 5 月 29 日
Agent Development Kit（ADK）はエージェントとツールの認証を簡素化しますが、Cloud Run でホストされているモデルにアクセスする場合、LiteLLM コネクタを使用すると認証がより困難になります。このガイドでは、Google 署名付き OpenID (ID) トークンを取得し、ADK を使用してそれを LiteLLM 通信チャネルに挿入する方法について説明します。
Google Cloud Run は、強制的な認証と IAM ポリシーに基づいた堅牢な組み込みのアクセス制御メカニズムを提供します。これを有効にすると、特定の Cloud Run Invoker ロールを持つ認証されたアカウントによって行われた呼び出しのみが受け入れられ、サービスが不正な呼び出しから保護されます。
認証された呼び出しを実装するには、次の手順を実装する必要があります。
サインインプロセスを実装するか、アプリケーションのデフォルト資格情報を使用して資格情報を取得します。
ユーザーまたはワークロード ID から ID トークンを取得する
Cloud Run エンドポイントを呼び出すときに、HTTP 認証ヘッダーのベアラー トークンとして ID トークンを使用します。
Agent Development Kit (ADK) を使用すると、これらの手順の実装が大幅に簡素化されます。エージェントが別のエージェントまたは Cloud Run 上で実行される MCP (Model Context Protocol) サーバーを呼び出すと、フレームワークはアプリケーションのデフォルト認証情報の検出、トークンの取得、エージェント間の各呼び出しおよびエージェントからリモート MCP ツールへの呼び出しへのトークンの挿入を処理します。このフレームワークは、現在のトークンの有効期限が切れた場合のトークンの更新も実装します。開発者は、リモート MCP サーバーまたはエージェント オブジェクトを構成する以外に何も実装する必要はありません。
設定すると状況が異なります

Cloud Run にデプロイされたモデルを使用するための ADK エージェント。 ADK は、エージェントがリモート エンドポイントでホストされている非 Gemini モデルを使用できるようにする LiteLLM コネクタを提供します。ただし、モデルへの認証された呼び出しを自分で行う必要があります。どうすればそんなことができるのでしょうか？
LiteLLM コネクタは、litellm Python パッケージを使用して、Ollama、vLLM、およびその他の LLM エンジンを公開するリモート エンドポイントを呼び出します。このパッケージは、ヘッダーの名前と値のマップに設定される external_headers パラメーターを使用して、 acompletion などの LiteLLM API への呼び出しでカスタム HTTP ヘッダーを渡すことをサポートします。
google.adk.agents からエージェントをインポート
google.adk.models から LiteLlm をインポート
google.authをインポートする
google.auth.transport.requestsをインポートする
google.oauth2 から id_token をインポート
# 指定された対象者 (Cloud Run サービス URL) の Google 署名付き ID トークンを取得します。
aud = "https://model-123456789012.us-central1.run.app"
creds、_ = google 。認証。デフォルト()
auth_req = Google 。認証。輸送 。リクエスト。リクエスト()
信条 。リフレッシュ(認証要求)
if hasattr(creds, "id_token" ) および creds 。 id_トークン:
トークン = creds 。 id_トークン
それ以外の場合:
トークン = id_token 。 fetch_id_token(auth_req, aud)
# モデルをセットアップする
モデル = LiteLlm(
モデル = f "ollama_chat/gemma3:270m" ,
api_base = オーストラリア、
extra_headers = {
"認可" : f "ベアラー { トークン } " ,
}、
)
エージェント = エージェント(
名前 = "コンテンツビルダー" ,
モデル = モデル、
命令 = "エージェント システム命令" ,
)
このアプローチは、トークンが有効である限り機能します。取得したトークンの有効期限が切れると、Cloud Run は有効期限が切れたトークンを使用して行われた呼び出しをブロックするため、モデルへの呼び出しは「HTTP 401 Unauthorized」ステータス コードで失敗します。この方法は、ゼロ構成までスケールし、呼び出し頻度が低い Cloud Run にデプロイされたエージェントに適しています。この展開パターンでは

エージェントは頻繁に再起動されるため、新しい認証トークンが取得されます。
方法 2: 動的トークン注入
エージェントが継続的に実行されることが予想される場合、エラーを受信したときにトークンを更新する必要があります。これを実装するには、ADK の LiteLLMClient クラスを拡張する必要があります。
OSをインポートする
import Any、Optional、Union の入力から
fastapiインポートからHTTPExceptionが発生しました
google.adk.models.lite_llm から LiteLLMClient をインポート
google.authをインポートする
google.auth.transport.requestsをインポートする
google.oauth2 から id_token をインポート
Litellm.Exceptions からのインポート AuthenticationError
litellm import CustomStreamWrapper、ModelResponse から
_creds、_ = Google 。認証。デフォルト()
def _get_auth_token (aud: str) -> オプション[str]:
「」
指定された対象者 (Cloud Run サービス URL) の Google 署名付き ID トークンを取得します。
「」
試してみてください:
auth_req = Google 。認証。輸送 。リクエスト。リクエスト()
# ローカル開発でのユーザー認証情報の使用をサポート
_creds 。リフレッシュ(認証要求)
if hasattr(_creds, "id_token" ) および _creds 。 id_トークン:
_creds を返します。 id_トークン
# サービスアカウント認証情報のトークンを取得します
fetched_id_token = id_token 。 fetch_id_token(auth_req, aud)
fetched_id_token を返す
e としての例外を除く:
print( f "{ aud } の ID トークン取得エラー: { e } " )
なしを返す
クラス LiteLLMClientEx (LiteLLMClient):
「」
LiteLLMClient をオーバーライドして、要求ヘッダーにベアラー トークンを挿入します。
「」
def __init__ (self、対象者: str、** データ: Any) -> なし:
自分自身。トークン: オプション[str] = なし
自分自身。オード = 聴衆
super() 。 __init__ ( ** データ)
非同期デフォルト完了 (
自己、モデル、メッセージ、ツール、** kwargs
) -> Union[ModelResponse, CustomStreamWrapper]:
「」
ヘッダーを更新してベアラー トークンを挿入します。
「」
自分自身。トークン = self 。トークンまたは _get_auth_token(self . aud)
「extra_headers」が kwargs にない場合、または

kwargs[ "extra_headers" ] は None です:
kwargs[ "extra_headers" ] = {}
kwargs[ "extra_headers" ][ "Authorization" ] = f "Bearer { self . token } "
# ADK の内部ロジックを保持するために親クラスを acompletion で呼び出すようにします。
試してみてください:
super() を待機します。完了(
モデル = モデル、メッセージ = メッセージ、ツール = ツール、** kwargs
)
e: (AuthenticationError、HTTPException) を除く:
if isinstance(e, HTTPException) および e 。ステータスコード != 401 :
上げる
# トークンを更新する
自分自身。トークン = _get_auth_token(self . aud)
クワル
[切り捨てられた]
モデル = LiteLlm(
モデル = f "ollama_chat/gemma3:270m" ,
api_base = "https://model-123456789012.us-central1.run.app" ,
llm_client = LiteLLMClientEx(
オーディエンス = "https://model-123456789012.us-central1.run.app" )、
)
HTTP 401 ステータス エラーを処理する LiteLLMClientEx コード スニペット内の例外ハンドラーは、この投稿で使用するために簡略化されていることに注意してください。 Cloud Run は、トークンの対象ユーザーの不一致、間違ったトークン タイプ、不正な形式のトークンなど、複数のシナリオで HTTP 401 Unauthorized を返します。期限切れのトークンと他のシナリオを区別するための推奨される方法は、WWW-Authenticate ヘッダーでエラーの説明を調べることです。
エージェント コードを変更できない場合、または他の理由で方法 2 が機能しない場合は、litellm-proxy 構成を使用して、プロキシされたリクエストの Authorization ヘッダーを自動的に設定できます。 litellm プロキシは、エージェントと一緒にサイドカー コンテナとしてデプロイでき、プライマリ エージェント コードから認証ロジックをオフロードできます。
構成が完了すると、エージェントのコードを追加変更することなく、LiteLLM コネクタ構成のプロキシ ポートでローカル エンドポイントを使用できるようになります。
モデル = LiteLlm(
モデル = f "ollama_chat/gemma3:270m" ,
api_base = "http://0.0.0.0:4000" ,
)
何

次は
まとめると、Cloud Run でホストされているモデルとの LiteLLM コネクタ通信への ID トークンの注入を実装するには、3 つの異なる方法があります。自分にとって最適なものを選ぶのは簡単です。
エージェント インスタンスの実行は 1 時間以内に終了すると予想されますか?自動スケールをゼロまで下げて Cloud Run にエージェントをデプロイし、リクエスト頻度が低いことが予想される場合、インスタンスはアイドル状態になった後にシャットダウンされます。このシナリオでは、方法 1 を使用できます。これは簡単で、インスタンスがシャットダウンされるまで、取得された ID トークンが期限切れにならないことが保証されます。
エージェントの外部でモデルのエンドポイントを呼び出すための認証ロジックを一元化する必要がありますか?同じモデルを使用する複数のエージェントを実行する予定がある場合、またはモデル認証ロジックをエージェントに追加しないことに決めた場合は、方法 3 を使用する必要があります。LiteLLM プロキシを各エージェントのサイドカーとして、またはスタンドアロン サービスとして展開し、モデルにアクセスするためにプロキシのエンドポイントを使用するようにエージェントを構成します。プロキシをスタンドアロンの Cloud Run サービスとして実行する場合の微妙な違いの 1 つは、そのサービスを呼び出すために認証ロジックを実装する必要があることです。これは方法 2 の実装と実質的に同じです。
上記の質問の両方に「いいえ」と答えた場合は、方法 2 を選択してください。
認証の実装を進めるために使用できる参考資料をいくつか紹介します。
コードラボ「Gemini モデルと Gemma モデルを使用したマルチエージェント システムの構築」と、サイドカーを使用した Cloud Run サービスの作成を受講します。
GPU 上のセルフホスト LLM を使用した Multi-Agent Course Creator のメソッド 2 の実装を Github リポジトリで確認します。
ブログ投稿を読む どこからでも Cloud Run サービスを安全に呼び出す

## Original Extract

minherz: another techno-blog
Notes about cloud, technology, life and everything
Menu
Home
How to Securely Connect ADK Agents to Models on Cloud Run
approx. 1300 words
6 min
minherz
Last Modified: May 29, 2026
The Agent Development Kit (ADK) simplifies authentication for agents and tools, but is more challenging with the LiteLLM connector when accessing models hosted on Cloud Run. This guide explores how to acquire Google-signed OpenID (ID) tokens and inject them into the LiteLLM communication channel using ADK.
Google Cloud Run provides a robust, built-in access control mechanism based on enforced authentication and IAM policies. When it is enabled, only calls that are made by authenticated accounts which have the specific Cloud Run Invoker role, are accepted, protecting your service from unauthorized invocations.
To implement the authenticated call you have to implement the following steps:
Acquire credentials either by implementing a sign-in process or using Application Default Credentials
Fetch an ID token from a user or workload identity
Use the identity token as a bearer token in the HTTP Authorization header when you call Cloud Run endpoint
Agent Development Kit (ADK) greatly simplifies implementation of these steps for you. When your agent calls another agent or an MCP (Model Context Protocol) server that runs on Cloud Run, the framework handles discovering the application’s default credentials, fetching the token, and injecting the token into each call between agents and calls from agents to remote MCP tools. The framework also implements token refreshing when the current token is expired. As a developer, you don’t need to implement anything beyond configuring your remote MCP server or agent objects.
The situation is different when you configure an ADK agent to use a model that is deployed in Cloud Run. ADK provides a LiteLLM connector that allows an agent to use non-Gemini models hosted at remote endpoints. However, you will need to take care of making authenticated calls to the model yourself. How can you do that?
The LiteLLM connector uses the litellm Python package to call the remote endpoints exposing Ollama, vLLM and other LLM engines. The package supports passing custom HTTP headers in the calls to LiteLLM APIs like acompletion , using the external_headers parameter, which is set to the map of header’s names and values.
from google.adk.agents import Agent
from google.adk.models import LiteLlm
import google.auth
import google.auth.transport.requests
from google.oauth2 import id_token
# obtains a Google-signed ID token for the given audience (Cloud Run service URL).
aud = "https://model-123456789012.us-central1.run.app"
creds, _ = google . auth . default()
auth_req = google . auth . transport . requests . Request()
creds . refresh(auth_req)
if hasattr(creds, "id_token" ) and creds . id_token:
token = creds . id_token
else :
token = id_token . fetch_id_token(auth_req, aud)
# set up the model
model = LiteLlm(
model = f "ollama_chat/gemma3:270m" ,
api_base = aud,
extra_headers = {
"Authorization" : f "Bearer { token } " ,
},
)
agent = Agent(
name = "content_builder" ,
model = model,
instruction = "agent system instructions" ,
)
This approach works as long as the token is valid. Once the fetched token is expired the calls to the model will fail with the “HTTP 401 Unauthorized” status code because Cloud Run will block calls made with the expired token. This method would fit for agents deployed on Cloud Run with scale to zero configuration and having a low invocation frequency. In this deployment pattern the agent will be frequently restarted which will lead to fetching a new authentication token.
Method 2: Dynamic token injection
When the agent is expected to run continuously, the token has to be refreshed upon receiving an error. To implement this, you would need to extend the LiteLLMClient class of ADK:
import os
from typing import Any, Optional, Union
from fastapi import HTTPException
from google.adk.models.lite_llm import LiteLLMClient
import google.auth
import google.auth.transport.requests
from google.oauth2 import id_token
from litellm.exceptions import AuthenticationError
from litellm import CustomStreamWrapper, ModelResponse
_creds, _ = google . auth . default()
def _get_auth_token (aud: str) -> Optional[str]:
"""
Obtains a Google-signed ID token for the given audience (Cloud Run service URL).
"""
try :
auth_req = google . auth . transport . requests . Request()
# support using user credentials for local development
_creds . refresh(auth_req)
if hasattr(_creds, "id_token" ) and _creds . id_token:
return _creds . id_token
# fetch token for service account credentials
fetched_id_token = id_token . fetch_id_token(auth_req, aud)
return fetched_id_token
except Exception as e:
print( f "Error obtaining ID token for { aud } : { e } " )
return None
class LiteLLMClientEx (LiteLLMClient):
"""
Overrides the LiteLLMClient to inject a bearer token into the request headers.
"""
def __init__ (self, audience: str, ** data: Any) -> None :
self . token: Optional[str] = None
self . aud = audience
super() . __init__ ( ** data)
async def acompletion (
self, model, messages, tools, ** kwargs
) -> Union[ModelResponse, CustomStreamWrapper]:
"""
Updating headers to inject bearer token.
"""
self . token = self . token or _get_auth_token(self . aud)
if "extra_headers" not in kwargs or kwargs[ "extra_headers" ] is None :
kwargs[ "extra_headers" ] = {}
kwargs[ "extra_headers" ][ "Authorization" ] = f "Bearer { self . token } "
# Ensure we call the parent class acompletion to preserve ADK internal logic
try :
return await super() . acompletion(
model = model, messages = messages, tools = tools, ** kwargs
)
except (AuthenticationError, HTTPException) as e:
if isinstance(e, HTTPException) and e . status_code != 401 :
raise
# refresh the token
self . token = _get_auth_token(self . aud)
kwar
[truncated]
model = LiteLlm(
model = f "ollama_chat/gemma3:270m" ,
api_base = "https://model-123456789012.us-central1.run.app" ,
llm_client = LiteLLMClientEx(
audience = "https://model-123456789012.us-central1.run.app" ),
)
Note that the exception handler in the LiteLLMClientEx code snippet that handles the HTTP 401 status error is simplified for use in this post. Cloud Run returns HTTP 401 Unauthorized in multiple scenarios including mismatched audience of the token, wrong token type or malformed token and more. The recommended way to distinguish between expired token and other scenarios is to examine the WWW-Authenticate header for the error description.
If you cannot modify the agent code or Method 2 does not work for you for other reasons, you can use a litellm-proxy configuration to automatically populate the Authorization header of the proxied requests. The litellm proxy can be deployed as a sidecar container alongside your agent, offloading the authentication logic from your primary agent code.
Once configured, you will use the local endpoint with the proxy port in the LiteLLM connector configuration without need for any additional modifications of the agent’s code.
model = LiteLlm(
model = f "ollama_chat/gemma3:270m" ,
api_base = "http://0.0.0.0:4000" ,
)
What’s next
Wrapping up, there are three different methods to implement injection of ID token into LiteLLM connector communication with a model hosted on Cloud Run. Choosing the best one for you is simple.
Do you expect your agent instance to run less than one hour before being terminated? If you deploy your agent on Cloud Run with autoscale down to zero, and expect having low request frequency, your instance will be shut down after being idle . In this scenario you can use Method 1 . It is simple and will guarantee that the fetched ID token will not expire until the instance is shut down.
Do you need to centralize authentication logic for calling your model’s endpoint outside of agents? If you plan to run more than one agent that use the same model or there is a decision not to add the model authentication logic into the agent, you should use Method 3 - deploy the LiteLLM proxy as a sidecar to each agent or as a standalone service and to configure agents to use the proxy’s endpoint to access the model. One subtle nuance of running the proxy as a standalone Cloud Run service is that you will need to implement authentication logic to call that service which is effectively the same as implementing Method 2.
If you answered “No” to both of the above questions, your choice should be Method 2 .
Here are a few useful references that you can use to move forward with implementing authentication:
Take codelabs Building a Multi-Agent System using Gemini and Gemma models and Create a Cloud Run service with a sidecar
Explore Method 2 implementation in the Multi-Agent Course Creator with Self-Hosted LLM on GPU Github repo.
Read a blog post Securely Call Cloud Run Service From Anywhere
