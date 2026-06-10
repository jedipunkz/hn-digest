---
source: "https://developers.openai.com/cookbook/examples/partners/aws/openai_models_with_amazon_bedrock"
hn_url: "https://news.ycombinator.com/item?id=48471167"
title: "Getting Started with OpenAI Models on Amazon Bedrock"
article_title: "Getting Started with OpenAI Models on Amazon Bedrock"
author: "gmays"
captured_at: "2026-06-10T04:34:31Z"
capture_tool: "hn-digest"
hn_id: 48471167
score: 1
comments: 0
posted_at: "2026-06-10T03:45:31Z"
tags:
  - hacker-news
  - translated
---

# Getting Started with OpenAI Models on Amazon Bedrock

- HN: [48471167](https://news.ycombinator.com/item?id=48471167)
- Source: [developers.openai.com](https://developers.openai.com/cookbook/examples/partners/aws/openai_models_with_amazon_bedrock)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T03:45:31Z

## Translation

タイトル: Amazon Bedrock での OpenAI モデルの使用開始
説明: Amazon Bedrock の OpenAI モデルは、テキスト生成、構造体を必要とする本番ワークフロー向けに OpenAI 互換の Responses API サーフェスを公開します。

記事本文:
Amazon Bedrock での OpenAI モデルの使用開始
ホーム API ドキュメント OpenAI API のガイドと概念 API リファレンス エンドポイント、パラメーター、および応答 Codex ドキュメント Codex のガイド、概念、および製品ドキュメント 使用例 ワークフローとタスクの例 チームが Codex ChatGPT Apps SDK に渡す ChatGPT Commerce を拡張するアプリを構築する ChatGPT 広告でコマース フローを構築する ChatGPT 広告で広告を公開および測定する リソース ショーケース インスピレーションを得るためのデモ アプリ ブログ 開発者からの学習と経験 クックブックOpenAI モデルを使用して構築するためのノートブックの例 OpenAI を使用して構築するためのドキュメント、ビデオ、デモ アプリを学ぶ コミュニティ プログラム、交流会、ビルダーのサポート API の検索を開始する ダッシュボード クックブックを検索する
統合と可観測性
MCP とコネクタによる安全な MCP トンネル
ファイルの検索と取得 ファイルの検索
文字起こし リアルタイム文字起こし
リアルタイムセッション 会話の管理
Webhook とサーバー側コントロール
Workload Identity フェデレーションの概要
最適化サイクルの微調整
直接的なプリファレンスの最適化
アシスタント API 移行ガイド
ブログ スキルを活用して OSS メンテナンスを加速する
Codex と Figma を使用したフロントエンド UI の構築
クックブック トレース、評価、コーデックスを使用してエージェント改善ループを構築する
Codex を使用して反復修復ループを構築する
Perplexity がリアルタイム API を使用して音声検索を何百万人ものユーザーに提供した方法
GPT-5.4 を使用した楽しいフロントエンドの設計
プロンプトから製品まで: 1 年間の対応
スキルを活用してOSSのメンテナンスを加速する
Codex と Figma を使用したフロントエンド UI の構築
ページをコピー ページをコピー 2026 年 5 月 31 日 Amazon Bedrock で OpenAI モデルを使ってみる
SS シカール クワトラ
(オープンAI)
、スディーシュ・サシダラン
GitHub で見る
生でダウンロード
Amazon Bedrock の OpenAI モデルは、本番ワークフロー向けの OpenAI 互換の Response API サーフェスを公開します。

テキスト生成、構造化出力、アプリケーション ツール、直接ファイル入力、応答状態、プロンプト キャッシュ、およびバックグラウンド作業が必要です。このクックブックでは、遅延や破損した注文の交換リクエストを処理する架空の小売店である BrightCart のサポート アシスタント ワークフローを構築することで、例を具体的にしています。
通常のアプリケーション呼び出しには OpenAI Python SDK を使用し、正確なリクエスト本文を検査するのに便利な場合は小さな生の HTTPS ヘルパーを使用します。このフローは、セットアップと最小限のプリフライトから始まり、応答ライフサイクル、モデル コントロール、構造化 JSON、アプリケーション ツール、ファイル入力、状態管理、キャッシュ、バックグラウンド処理、コンテキストの圧縮、操作チェック、クリーンアップを重ねていきます。
Bedrock 固有の環境変数を使用して、Bedrock でホストされる OpenAI モデルを構成します。
応答エンドポイントを検証し、応答スキーマ、使用状況メタデータ、および正規化されたエラーを検査します。
raw HTTPS と OpenAI SDK の両方を使用してテキスト リクエストを送信します。
スキーマ制約のある JSON およびより軽量な JSON モードのハンドオフを生成します。
アプリケーション管理の関数ツール、並列ツール、およびカスタム テキスト ツールを呼び出します。
直接 PDF 入力を送信し、ステートフルおよびステートレスな会話を継続し、暗号化された推論コンテキストを伝送します。
プロンプト キャッシュ、バックグラウンド モード、圧縮、動作スモーク チェック、およびストアド レスポンス クリーンアップを使用します。
前提条件: Amazon Bedrock 上の OpenAI モデルのベアラー トークン、Python 3.9 以降、および Bedrock OpenAI 互換エンドポイントへのネットワーク アクセス。
このガイドでは、デフォルトで us-west-2 で openai.gpt-5.4 を実行します。サポートされている別のペアリングを使用するには、セットアップ セルを実行する前に、 AWS_REGION 、 BEDROCK_MODEL 、および BEDROCK_BASE_URL を一緒に変更します。
このセクションでは、ノートブック ランタイムを準備します。小さな Python スタックをインストールし、Bedrock 固有の環境変数を読み取り、生の HTT を作成します。

PS セッションと OpenAI SDK クライアントは、エンドポイントが提供するモデル メタデータを検出し、後の例で使用される共有ヘルパーを定義します。
ノートブックを実行する前に、これらの環境変数を設定します。デフォルトのペアリングは、 us-west-2 と openai.gpt-5.4 です。
エクスポート AWS_BEARER_TOKEN_BEDROCK = "YOUR_BEDROCK_BEARER_TOKEN"
エクスポート AWS_REGION = "us-west-2"
import BEDROCK_MODEL = "openai.gpt-5.4"
import BEDROCK_BASE_URL = "https://bedrock-mantle.${ AWS_REGION }.api.aws/openai/v1"
ベアラートークンは AWS_BEARER_TOKEN_BEDROCK から読み取られます。これが存在しない場合、セットアップ セルはパスワード形式のプロンプトでそれを要求し、印刷されません。
ノートブックで使用されるパッケージをインストールします。 OpenAI SDK はアプリケーションの例に使用され、リクエストは Responses エンドポイントへの生の HTTPS 呼び出しに使用され、パンダと IPython 表示ヘルパーによってリクエストとレスポンスの概要がクックブック レンダラーで読み取り可能な状態に保たれます。セルの出力を検査するのは、パッケージがインストールされているか、すでに存在していることを確認する場合のみです。
% pip install - U "openai>=2.28.0" は pandas ipython を要求します -- Quiet
print ( "依存関係がインストールされているか、すでに利用可能です: openai、requests、pandas、ipython" )
[1m[ [0m [34;49mnotice [0m [1;39;49m] [0m [39;49m pip の新しいリリースが利用可能です: [0m [31;49m24.0 [0m [39;49m -> [0m [32;49m26.1.1 [0m
[1m[ [0m [34;49mnotice [0m [1;39;49m] [0m [39;49m 更新するには、次を実行します: [0m [32;49mpip install --upgrade pip [0m
注: 更新されたパッケージを使用するには、カーネルの再起動が必要な場合があります。
依存関係がインストールされているか、すでに利用可能なもの: openai、requests、pandas、ipython
1.2 ライブラリとデフォルトのインポート
ノートブック全体で使用される標準ライブラリ、SDK、HTTP クライアント、および表示ユーティリティをインポートします。このセルは、環境変数がまだ設定されていない場合に使用されるデフォルトの Bedrock 領域とモデルも設定します。検査します

オーバーライドしない限り、ノートブックが us-west-2 と openai.gpt-5.4 から開始されることを確認するために、デフォルトを印刷しました。
__future__ からアノテーションをインポート
インポートbase64
組み込み関数をインポートする
HTMLをインポートする
jsonをインポートする
OSをインポートする
輸入シュレックス
テキストラップをインポートする
インポート時間
datetime インポート日から、timedelta
getpass から getpass をインポート
import Any、Callable、Iterable の入力から
パンダをPDとしてインポートする
インポートリクエスト
IPython.display から HTML 、Markdown、表示をインポート
openaiインポートからOpenAI
DEFAULT_REGION = "us-west-2"
DEFAULT_MODEL = "openai.gpt-5.4"
PREFERRED_MODELS = [ DEFAULT_MODEL ]
def gpt_version_tuple (model_id: str ) -> タプル[ int , int ] |なし :
Normalized = model_id. lower().removeprefix( "openai." )
正規化されていない場合.startswith( "gpt-" ):
なしを返す
バージョン = Normalized.removeprefix( "gpt-" ).split( "-" )[ 0 ]
Parts = version.split( "." )
試してみてください:
Major =builtins.int(parts[ 0 ])
minor =builtins.int(parts[ 1 ]) if len (parts) > 1 else 0
ValueError を除く:
なしを返す
メジャー、マイナーを返す
def プロンプトキャッシュ保持_for_model (model_id: str ) -> str :
バージョン = gpt_version_tuple(model_id)
バージョンとバージョン >= ( 5 , 5 ) の場合:
「24時間」を返す
「メモリ内」を返す
pd.set_option( "display.max_columns" , None )
pd.set_option( "display.max_rows" , 200 )
pd.set_option( "display.max_colwidth" , None )
pd.set_option( "display.width" , 160 )
def display_wrapped_table (df: pd.DataFrame, * 、max_col_width_px: int = 520 、index: bool = False) -> なし:
df.emptyの場合:
display(Markdown( "_表示する行がありません。_" ))
戻る
table_html = df.to_html( インデックス = インデックス、エスケープ = True 、ボーダー = 0 )
table_html = table_html.replace( '<table border="0" class="dataframe">' , '<table class="dataframe Wrapped-output-table">' )
display(HTML( f """
<スタイル>
.wrapped-output-table {{
境界崩壊: 崩壊;
幅: 100

%;
テーブルレイアウト: 自動;
フォントサイズ: 13px;
}}
.wrapped-output-table th、
.wrapped-output-table td {{
境界線: 1 ピクセルの実線 #d0d7de;
パディング: 6px 8px;
テキスト整列: 左;
垂直配置: 上;
ホワイトスペース: ラップ前;
オーバーフローラップ: どこでも;
単語区切り: 単語区切り;
最大幅: { max_col_width_px } ピクセル;
}}
.wrapped-output-table t
[切り捨てられた]
環境から Bedrock 構成を読み取り、クライアントを構築します。 BEDROCK_BASE_URL は一度正規化され、生のrequests.Session はヘッダー内のベアラー トークンを取得し、同じトークンとベース URL を使用して OpenAI SDK クライアントが明示的に作成されます。ライブ呼び出しを行う前に、レンダリングされたテーブルを調べて、選択したリージョン、モデル、エンドポイント、SDK クライアント構成、および保存された応答のクリーンアップ動作を確認します。
__future__ からアノテーションをインポート
def env_value ( * 名前: str ) -> str |なし :
名前の中の名前の場合:
値 = os.environ.get(名前)
値の場合:
戻り値
なしを返す
def env_flag (名前: str 、デフォルト: bool = False ) -> bool :
値 = 環境値(名前)
値が None の場合:
デフォルトを返す
戻り値.strip(). lower() の { "1" 、 "true" 、 "yes" 、 "on" }
def Normalize_base_url (url: str ) -> str :
URL = url.strip().rstrip( "/" )
if url.endswith( "/responses" ):
URL を返す[: - len ( "/responses" )]
戻りURL
def エンドポイント (パス: str ) -> str :
return f " {BEDROCK_BASE_URL} / { path.lstrip( '/' ) } "
defresponse_url (base_url: str ) -> str :
return f " {normalize_base_url(base_url) } /responses"
API_TIMEOUT_SECONDS = float (env_value( "BEDROCK_REQUEST_TIMEOUT_SECONDS" ) または "60" )
MAX_RETRIES =builtins.int(env_value( "BEDROCK_MAX_RETRIES" ) または "0" )
CLEAN_UP_STORED_RESPONSES = env_flag( "BEDROCK_CLEANUP_STORED_RESPONSES" , True )
FAIL_ON_CHECK_FAILURE = env_flag( "BEDROCK_FAIL_ON_CHECK_FAILURE" , False )
RUN_RESPONSIVENESS_CHECK = env_flag( "BEDRO

CK_RESPONSIVENESS_CHECK" 、真)
TRANSIENT_STATUS_CODES = { 408 , 409 , 429 , 500 , 502 , 503 , 504 }
AWS_REGION = (env_value( "AWS_REGION" ) または DEFAULT_REGION ).strip() または DEFAULT_REGION
BEDROCK_MODEL = (env_value( "BEDROCK_MODEL" ) または DEFAULT_MODEL ).strip() または DEFAULT_MODEL
BEDROCK_BASE_URL = Normalize_base_url(
env_value( "BEDROCK_BASE_URL" ) または f "https://bedrock-mantle. {AWS_REGION} .api.aws/openai/v1"
)
RESPONSES_URL = 応答_url( BEDROCK_BASE_URL )
AWS_BEARER_TOKEN_BEDROCK = env_value( "AWS_BEARER_TOKEN_BEDROCK" )
AWS_BEARER_TOKEN_BEDROCK でない場合:
AWS_BEARER_TOKEN_BEDROCK = getpass( "このカーネル セッションの AWS Bedrock ベアラー トークンを貼り付けます: " ).strip()
AWS_BEARER_TOKEN_BEDROCK の場合:
os.environ[ "AWS_BEARER_TOKEN_BEDROCK" ] = AWS_BEARER_TOKEN_BEDROCK
AWS_BEARER_TOKEN_BEDROCK でない場合:
raise RuntimeError (「ライブサンプルを実行するには AWS_BEARER_TOKEN_BEDROCK が必要です。」)
http = リクエスト.セッション()
http.headers.update({
「あ
[切り捨てられた]
選択したエンドポイントがモデル リスト メタデータを公開するときに利用可能なモデルを検出し、ノートブックの残りの部分のモデルを選択します。 BEDROCK_MODEL が設定されている場合、ノートブックはその値を使用します。それ以外の場合は、 openai.gpt-5.4 が優先されます。一部の互換性のあるエンドポイントでは、モデルのメタデータが利用できない場合でも推論が許可される場合があるため、model-list 呼び出しはオプションです。選択したモデルと返されたカタログ行を検査します。
__future__ からアノテーションをインポート
def list_openai_models (クライアント: OpenAI) -> list[ str ]:
並べ替えて返す (client.models.list( timeout = API_TIMEOUT_SECONDS ).data 内のモデルの model.id)
defsolve_model_id (クライアント: OpenAI | None ) -> タプル[ str , list[ str ], str |なし]:
configured_model = env_value( "BEDROCK_MODEL" )
available_models: リスト[ str ] = []
モデル発見ノート: str |なし = なし
クライアントが None でない場合:
試してみてください:
ある

vailable_models = list_openai_models(クライアント)
exc としての例外を除く:
status_code = getattr (exc, "status_code" , None )
ステータスコード == 404 の場合:
" model_discovery_note = "このエンドポイントはモデルリストのメタデータを公開しませんでした。ガイドは構成されたモデルで続行されます。"
それ以外の場合:
model_discovery_note = f "モデルリストのメタデータをリストできませんでした。ガイドは構成されたモデルに進みます。詳細: {builtins.str(exc)[: 240 ] } "
構成モデルの場合:
戻り値configured_model、available_models、model_discovery_note
PREFERRED_MODELS の候補者:
available_models の候補の場合:
候補を返す、available_models、model_discovery_note
available_models の候補の場合:
ifcandidate.startswith( "openai." ):
候補を返す、available_models、model_discovery_note
利用可能な場合_モデル:
戻り値 available_models[ 0 ]、available_models、model_discovery_note
PREFERRED_MODELS [ 0 ]、available_models、model_discovery_note を返す
EXPLICIT_MODEL = env_value( "BEDROCK_MODEL" )
MODEL_ID 、 AVAILABLE_MODELS 、 MODEL_DISCOVERY_NOTE =solve_model_id(client)
os.environ[ "BEDROCK_MODEL" ] = MODEL_ID
PROMPT_CACHE_RETENTION = モデルのプロンプトキャッシュ保持率( MODEL_ID )
PROMPT_CACHE_RETENTION_NOTE = (
「GPT-5.5」

[切り捨てられた]

## Original Extract

OpenAI models on Amazon Bedrock expose an OpenAI-compatible Responses API surface for production workflows that need text generation, struct

Getting Started with OpenAI Models on Amazon Bedrock
Home API Docs Guides and concepts for the OpenAI API API reference Endpoints, parameters, and responses Codex Docs Guides, concepts, and product docs for Codex Use cases Example workflows and tasks teams hand to Codex ChatGPT Apps SDK Build apps to extend ChatGPT Commerce Build commerce flows in ChatGPT Ads Publish and measure ads in ChatGPT Resources Showcase Demo apps to get inspired Blog Learnings and experiences from developers Cookbook Notebook examples for building with OpenAI models Learn Docs, videos, and demo apps for building with OpenAI Community Programs, meetups, and support for builders Start searching API Dashboard Search the cookbook
Integrations and observability
MCP and Connectors Secure MCP Tunnel
File search and retrieval File search
Transcription Realtime transcription
Realtime sessions Managing conversations
Webhooks and server-side controls
Workload identity federation Overview
Fine-tuning Optimization cycle
Direct preference optimization
Assistants API Migration guide
Blog Using skills to accelerate OSS maintenance
Building frontend UIs with Codex and Figma
Cookbooks Build an Agent Improvement Loop with Traces, Evals, and Codex
Build iterative repair loops with Codex
How Perplexity Brought Voice Search to Millions Using the Realtime API
Designing delightful frontends with GPT-5.4
From prompts to products: One year of Responses
Using skills to accelerate OSS maintenance
Building frontend UIs with Codex and Figma
Copy Page Copy Page May 31, 2026 Getting Started with OpenAI Models on Amazon Bedrock
SS Shikhar Kwatra
(OpenAI)
, Sudeesh Sasidharan
View on GitHub
Download raw
OpenAI models on Amazon Bedrock expose an OpenAI-compatible Responses API surface for production workflows that need text generation, structured outputs, application tools, direct file inputs, response state, prompt caching, and background work. This cookbook keeps the examples concrete by building a support-assistant workflow for BrightCart , a fictional retailer handling delayed and damaged-order replacement requests.
You will use the OpenAI Python SDK for normal application calls and a small raw HTTPS helper when it is useful to inspect the exact request body. The flow starts with setup and a minimal preflight, then layers on response lifecycle, model controls, structured JSON, application tools, file input, state management, caching, background processing, context compaction, operations checks, and cleanup.
Configure a Bedrock-hosted OpenAI model with Bedrock-specific environment variables.
Verify the Responses endpoint and inspect response schema, usage metadata, and normalized errors.
Send text requests with both raw HTTPS and the OpenAI SDK.
Generate schema-constrained JSON and lighter JSON-mode handoffs.
Call application-managed function tools, parallel tools, and custom text tools.
Send a direct PDF input, continue stateful and stateless conversations, and carry encrypted reasoning context.
Use prompt caching, background mode, compaction, operational smoke checks, and stored-response cleanup.
Prerequisites: a bearer token for OpenAI models on Amazon Bedrock, Python 3.9 or newer, and network access to your Bedrock OpenAI-compatible endpoint.
This guide runs openai.gpt-5.4 in us-west-2 by default. To use another supported pairing, change AWS_REGION , BEDROCK_MODEL , and BEDROCK_BASE_URL together before running the setup cells.
This section prepares the notebook runtime. It installs the small Python stack, reads Bedrock-specific environment variables, creates both a raw HTTPS session and an OpenAI SDK client, discovers model metadata when the endpoint provides it, and defines shared helpers used by later examples.
Set these environment variables before running the notebook. The default pairing is us-west-2 with openai.gpt-5.4 .
export AWS_BEARER_TOKEN_BEDROCK = "YOUR_BEDROCK_BEARER_TOKEN"
export AWS_REGION = "us-west-2"
export BEDROCK_MODEL = "openai.gpt-5.4"
export BEDROCK_BASE_URL = "https://bedrock-mantle.${ AWS_REGION }.api.aws/openai/v1"
The bearer token is read from AWS_BEARER_TOKEN_BEDROCK . If it is missing, the setup cell asks for it with a password-style prompt and does not print it.
Install the packages used by the notebook. The OpenAI SDK is used for the application examples, requests is used for raw HTTPS calls to the Responses endpoint, and pandas plus IPython display helpers keep request and response summaries readable in the Cookbook renderer. Inspect the cell output only to confirm the packages installed or were already present.
% pip install - U "openai>=2.28.0" requests pandas ipython -- quiet
print ( "Dependencies installed or already available: openai, requests, pandas, ipython" )
[1m[[0m[34;49mnotice[0m[1;39;49m][0m[39;49m A new release of pip is available: [0m[31;49m24.0[0m[39;49m -> [0m[32;49m26.1.1[0m
[1m[[0m[34;49mnotice[0m[1;39;49m][0m[39;49m To update, run: [0m[32;49mpip install --upgrade pip[0m
Note: you may need to restart the kernel to use updated packages.
Dependencies installed or already available: openai, requests, pandas, ipython
1.2 Import Libraries and Defaults
Import the standard libraries, SDK, HTTP client, and display utilities used throughout the notebook. This cell also sets the default Bedrock region and model used when environment variables are not already set. Inspect the printed defaults to confirm the notebook will start from us-west-2 and openai.gpt-5.4 unless you override them.
from __future__ import annotations
import base64
import builtins
import html
import json
import os
import shlex
import textwrap
import time
from datetime import date, timedelta
from getpass import getpass
from typing import Any, Callable, Iterable
import pandas as pd
import requests
from IPython.display import HTML , Markdown, display
from openai import OpenAI
DEFAULT_REGION = "us-west-2"
DEFAULT_MODEL = "openai.gpt-5.4"
PREFERRED_MODELS = [ DEFAULT_MODEL ]
def gpt_version_tuple (model_id: str ) -> tuple[ int , int ] | None :
normalized = model_id.lower().removeprefix( "openai." )
if not normalized.startswith( "gpt-" ):
return None
version = normalized.removeprefix( "gpt-" ).split( "-" )[ 0 ]
parts = version.split( "." )
try :
major = builtins.int(parts[ 0 ])
minor = builtins.int(parts[ 1 ]) if len (parts) > 1 else 0
except ValueError :
return None
return major, minor
def prompt_cache_retention_for_model (model_id: str ) -> str :
version = gpt_version_tuple(model_id)
if version and version >= ( 5 , 5 ):
return "24h"
return "in_memory"
pd.set_option( "display.max_columns" , None )
pd.set_option( "display.max_rows" , 200 )
pd.set_option( "display.max_colwidth" , None )
pd.set_option( "display.width" , 160 )
def display_wrapped_table (df: pd.DataFrame, * , max_col_width_px: int = 520 , index: bool = False ) -> None :
if df.empty:
display(Markdown( "_No rows to display._" ))
return
table_html = df.to_html( index = index, escape = True , border = 0 )
table_html = table_html.replace( '<table border="0" class="dataframe">' , '<table class="dataframe wrapped-output-table">' )
display(HTML( f """
<style>
.wrapped-output-table {{
border-collapse: collapse;
width: 100%;
table-layout: auto;
font-size: 13px;
}}
.wrapped-output-table th,
.wrapped-output-table td {{
border: 1px solid #d0d7de;
padding: 6px 8px;
text-align: left;
vertical-align: top;
white-space: pre-wrap;
overflow-wrap: anywhere;
word-break: break-word;
max-width: { max_col_width_px } px;
}}
.wrapped-output-table t
[truncated]
Read Bedrock configuration from the environment and construct clients. BEDROCK_BASE_URL is normalized once, the raw requests.Session gets the bearer token in its headers, and the OpenAI SDK client is created explicitly with the same token and base URL. Inspect the rendered table to confirm the selected region, model, endpoint, SDK client configuration, and stored-response cleanup behavior before making live calls.
from __future__ import annotations
def env_value ( * names: str ) -> str | None :
for name in names:
value = os.environ.get(name)
if value:
return value
return None
def env_flag (name: str , default: bool = False ) -> bool :
value = env_value(name)
if value is None :
return default
return value.strip().lower() in { "1" , "true" , "yes" , "on" }
def normalize_base_url (url: str ) -> str :
url = url.strip().rstrip( "/" )
if url.endswith( "/responses" ):
return url[: - len ( "/responses" )]
return url
def endpoint (path: str ) -> str :
return f " {BEDROCK_BASE_URL} / { path.lstrip( '/' ) } "
def responses_url (base_url: str ) -> str :
return f " { normalize_base_url(base_url) } /responses"
API_TIMEOUT_SECONDS = float (env_value( "BEDROCK_REQUEST_TIMEOUT_SECONDS" ) or "60" )
MAX_RETRIES = builtins.int(env_value( "BEDROCK_MAX_RETRIES" ) or "0" )
CLEAN_UP_STORED_RESPONSES = env_flag( "BEDROCK_CLEANUP_STORED_RESPONSES" , True )
FAIL_ON_CHECK_FAILURE = env_flag( "BEDROCK_FAIL_ON_CHECK_FAILURE" , False )
RUN_RESPONSIVENESS_CHECK = env_flag( "BEDROCK_RESPONSIVENESS_CHECK" , True )
TRANSIENT_STATUS_CODES = { 408 , 409 , 429 , 500 , 502 , 503 , 504 }
AWS_REGION = (env_value( "AWS_REGION" ) or DEFAULT_REGION ).strip() or DEFAULT_REGION
BEDROCK_MODEL = (env_value( "BEDROCK_MODEL" ) or DEFAULT_MODEL ).strip() or DEFAULT_MODEL
BEDROCK_BASE_URL = normalize_base_url(
env_value( "BEDROCK_BASE_URL" ) or f "https://bedrock-mantle. {AWS_REGION} .api.aws/openai/v1"
)
RESPONSES_URL = responses_url( BEDROCK_BASE_URL )
AWS_BEARER_TOKEN_BEDROCK = env_value( "AWS_BEARER_TOKEN_BEDROCK" )
if not AWS_BEARER_TOKEN_BEDROCK :
AWS_BEARER_TOKEN_BEDROCK = getpass( "Paste your AWS Bedrock bearer token for this kernel session: " ).strip()
if AWS_BEARER_TOKEN_BEDROCK :
os.environ[ "AWS_BEARER_TOKEN_BEDROCK" ] = AWS_BEARER_TOKEN_BEDROCK
if not AWS_BEARER_TOKEN_BEDROCK :
raise RuntimeError ( "AWS_BEARER_TOKEN_BEDROCK is required to run the live examples." )
http = requests.Session()
http.headers.update({
"A
[truncated]
Discover available models when the selected endpoint exposes model-list metadata, then choose the model for the rest of the notebook. If BEDROCK_MODEL is set, the notebook uses that value; otherwise it prefers openai.gpt-5.4 . The model-list call is optional because some compatible endpoints may allow inference even when model metadata is unavailable. Inspect the selected model and any returned catalog rows.
from __future__ import annotations
def list_openai_models (client: OpenAI) -> list[ str ]:
return sorted (model.id for model in client.models.list( timeout = API_TIMEOUT_SECONDS ).data)
def resolve_model_id (client: OpenAI | None ) -> tuple[ str , list[ str ], str | None ]:
configured_model = env_value( "BEDROCK_MODEL" )
available_models: list[ str ] = []
model_discovery_note: str | None = None
if client is not None :
try :
available_models = list_openai_models(client)
except Exception as exc:
status_code = getattr (exc, "status_code" , None )
if status_code == 404 :
model_discovery_note = "This endpoint did not expose model-list metadata. The guide will continue with the configured model."
else :
model_discovery_note = f "Model-list metadata could not be listed. The guide will continue with the configured model. Details: { builtins.str(exc)[: 240 ] } "
if configured_model:
return configured_model, available_models, model_discovery_note
for candidate in PREFERRED_MODELS :
if candidate in available_models:
return candidate, available_models, model_discovery_note
for candidate in available_models:
if candidate.startswith( "openai." ):
return candidate, available_models, model_discovery_note
if available_models:
return available_models[ 0 ], available_models, model_discovery_note
return PREFERRED_MODELS [ 0 ], available_models, model_discovery_note
EXPLICIT_MODEL = env_value( "BEDROCK_MODEL" )
MODEL_ID , AVAILABLE_MODELS , MODEL_DISCOVERY_NOTE = resolve_model_id(client)
os.environ[ "BEDROCK_MODEL" ] = MODEL_ID
PROMPT_CACHE_RETENTION = prompt_cache_retention_for_model( MODEL_ID )
PROMPT_CACHE_RETENTION_NOTE = (
"GPT-5.5

[truncated]
