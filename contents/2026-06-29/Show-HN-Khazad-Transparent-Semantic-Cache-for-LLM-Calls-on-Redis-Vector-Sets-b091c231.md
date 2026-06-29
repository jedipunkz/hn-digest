---
source: "https://github.com/GuglielmoCerri/khazad"
hn_url: "https://news.ycombinator.com/item?id=48725166"
title: "Show HN: Khazad – Transparent Semantic Cache for LLM Calls on Redis Vector Sets"
article_title: "GitHub - GuglielmoCerri/khazad: Transparent, transport-layer semantic cache for LLM API calls, powered by Redis 8 Vector Sets. · GitHub"
author: "guglielmoce"
captured_at: "2026-06-29T21:30:32Z"
capture_tool: "hn-digest"
hn_id: 48725166
score: 2
comments: 0
posted_at: "2026-06-29T21:01:04Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Khazad – Transparent Semantic Cache for LLM Calls on Redis Vector Sets

- HN: [48725166](https://news.ycombinator.com/item?id=48725166)
- Source: [github.com](https://github.com/GuglielmoCerri/khazad)
- Score: 2
- Comments: 0
- Posted: 2026-06-29T21:01:04Z

## Translation

タイトル: HN を表示: Khazad – Redis ベクター セットでの LLM 呼び出しの透過的セマンティック キャッシュ
記事のタイトル: GitHub - GuglielmoCerri/khazad: Redis 8 Vector Sets を利用した、LLM API 呼び出し用の透過的なトランスポート層セマンティック キャッシュ。 · GitHub
説明: Redis 8 Vector Sets を利用した、LLM API 呼び出し用の透過的なトランスポート層セマンティック キャッシュ。 - グリエルモ・チェリ/カザド

記事本文:
GitHub - GuglielmoCerri/khazad: Redis 8 Vector Sets を利用した、LLM API 呼び出し用の透過的なトランスポート層セマンティック キャッシュ。 · GitHub
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
グリエルモチェリ
/
カザド
公共
通知
通知設定を変更するにはサインインする必要があります
追加

イオンナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
33 コミット 33 コミット .github/ workflows .github/ workflows docs/ _static docs/ _static 例 例 khazad khazad テスト テスト .gitignore .gitignore AGENT.md AGENT.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Redis Vector Sets を利用した LLM API 呼び出し用の透過的なトランスポート層セマンティック キャッシュ。
API 呼び出しが最大 50% 削減 · ヒット時の処理が最大 96% 高速化 · 支出が最大 50% 削減 · 100% 透明性
ヒット率 0.50 での数値を示します (280 ミリ秒のキャッシュ リプレイと最大 7900 ミリ秒のアップストリーム コール)。数値はトラフィックによって異なります。
Khazad はトランスポート層で LLM HTTP トラフィックをインターセプトし、アプリケーション コードを変更することなく、意味的に同等のリクエストを Redis ベクター キャッシュから処理します。
モデル認識 — 各 (プロバイダー、モデル) ペアは独自のベクトル セットを取得するため、プロンプトがどれほど似ていても、gpt-4o 応答が gpt-4o-mini 呼び出しに提供されることはありません。 [cache_scope="host" をプロバイダー ホストのみでスコープするように設定し、同じプロバイダー上のすべてのモデルまたはデプロイメントが 1 つのキャッシュを共有できるようにします (異なるプロバイダーは分離されたままになります。「構成」を参照)。
会話対応 — ユーザーの最後のターンだけでなく、メッセージ リスト全体 (システム、ユーザー、アシスタント) が埋め込まれます。同じフォローアップの質問 (「人口はどうですか?」) で終わる 2 つの異なる会話が衝突することはありません。
双方向ストリーミング — キャッシュ ヒットは実際の SSE ストリーム (同期クライアントおよび非同期クライアント) として再生されます。キャッシュは、ストリームが遅延を追加することなくチャンクごとにキャプチャされ、正規の JSON 応答に再構築されることを見逃します。そのため、ストリームされた応答が後で非ストリームのリクエストを処理でき、またその逆も可能です。中止されたストリームは決してキャッシュされません

。
セマンティック キャッシュでは、コストとレイテンシーと引き換えに正確性が優先されます。取引を開始する前に、取引について理解してください。次の場合に使用します。
大量の反復的なトラフィック: FAQ ボット、サポート アシスタント、多くのユーザーがほぼ同じ質問をする RAG フロントエンド
開発 / テスト / CI 環境 - 実行のたびに同じプロンプトに対して料金を支払う必要がなくなります
決定的で即時の応答が特徴のデモと負荷テスト
内部ツールのコスト上限
プライバシー: プロンプトは埋め込まれ、応答はクリア テキストで Redis に保存されます。プロンプトに PII またはシークレットが含まれる可能性がある場合は、 ttl を設定し、Redis AUTH/TLS を有効にして、Redis インスタンスをログと同じように注意して扱います。
プロセス全体のパッチ: Khazad は、init() の後に作成されたすべての httpx.Client / AsyncClient をラップします。非 LLM httpx トラフィックはそのまま通過しますが、パッチはプロセス全体に適用されます。シャットダウン時に stop() を呼び出します。 hosts=[...] を使用して、実際にキャッシュしたいエンドポイントへのインターセプトを制限します。
httpx のみ : httpx 上に構築された SDK がカバーされます (OpenAI、Anthropic、google-genai 経由の Gemini、Mistral、およびほとんどのプロキシ)。 request 、 aiohttp 、または boto3 (AWS Bedrock) を使用する SDK はインターセプトされません。
単一プロセス: パッチは、 init() を呼び出した Python プロセス内に存在します。複数のワーカーが Redis キャッシュを共有しますが、それぞれに独自の init() が必要です。
偽陽性制御: しきい値 = 0.90 から開始し、間違ったヒットが見つかった場合は値を上げます。 get_stats() の avg_hit_similarity を監視します。これがしきい値に近い場合、トラフィックが多様すぎて安全にキャッシュできない可能性があります。
Redis 8 (ベクター セットのサポートが必要)
docker run -d --name redis8 -p 6379:6379 redis:8
インストール
UV追加カザド
OpenAI 埋め込みバックエンドの場合 (オプション):
uv 追加 khazad[openai-embeddings]
ローカル/開発インストール:
git クローン https://github.com/GuglielmoCerri/khazad.git
CDカザド
UV同期 --グループ開発
UV同期はpypを読み取ります

roject.toml は、 .venv が存在しない場合はそれを作成し、プロジェクト自体を編集可能モードでインストールします。個別の pip install -e は必要ありません。必要です。
別のプロジェクトからローカル チェックアウトを使用するには:
uv add --editable /path/to/khazad
使用法
インスタンスを明示的に制御する必要がある場合は、Khazad クラスを直接使用します。これは、長時間実行されるサービス、テスト、または依存関係の注入に役立ちます。
カザドから カザドを輸入
キャッシュ = カザド (
redis_url = "redis://localhost:6379" ,
しきい値 = 0.90 、
ttl = 3600 、
log_level = "デバッグ" ,
)
print (cache .is_active()) # True
print ( キャッシュ . get_stats ()) # Stats(total_requests=0, ...)
利用可能な関数: init() 、 stop() 、 get_stats() 、 flash() 、 is_active() 。詳細については、「API リファレンス」を参照してください。
Khazad は一度アクティブ化すると、その下で httpx を使用するすべての LLM SDK をインターセプトするため、プロバイダーごとの配線は必要ありません。その他の例については、サンプル フォルダーを参照してください。
OSをインポートする
インポート時間
openaiインポートからOpenAI
カザドから カザドを輸入
キャッシュ = Khazad ( redis_url = "redis://localhost:6379" 、しきい値 = 0.90 、 log_level = "DEBUG" )
client = OpenAI ( api_key = os .environ [ "OPENAI_API_KEY" ])
プロンプト = 「イタリアの首都はどこですか?」
範囲 ( 2 ) の i の場合:
開始 = 時間。 perf_counter ()
応答 = クライアント 。チャット 。完成品。作成(
モデル = "gpt-4o" 、
メッセージ = [{ "役割" : "ユーザー" , "コンテンツ" : プロンプト }],
)
経過 = (時間 .perf_counter () - 開始 ) * 1000
print ( f"[call { i + 1 } ] { 経過時間 :.1f } ms — { 応答 . 選択肢 [ 0 ]. メッセージ . コンテンツ } " )
print ( キャッシュ . get_stats () . to_dict ())
キャッシュ。やめて（）
*/chat/completions および */responses パスと一致します。ストリーミングリクエストもキャッシュされます。
OSをインポートする
インポート時間
紺碧から。 ID インポート DefaultAzureCredential 、 get_bearer_token_provider
openai から AzureOpenAI をインポート
カザドからインポート CacheScope 、カザド
キャッシュ

= カザド (
redis_url = "redis://localhost:6379" ,
しきい値 = 0.90 、
キャッシュ_スコープ = キャッシュスコープ 。ホスト、
名前空間 = "azure_openai_example" ,
)
エンドポイント = os 。環境 [ "AZURE_OPENAI_ENDPOINT" ]
デプロイメント = OS 。環境。 get ( "AZURE_OPENAI_DEPLOYMENT" , "gpt-4.1" )
トークンプロバイダー = get_bearer_token_provider (
DefaultAzureCredential ()、"https://cognitiveservices.azure.com/.default"
)
api_version = os 。環境。 get ( "AZURE_OPENAI_API_VERSION" , "2024-12-01-preview" )
クライアント = AzureOpenAI (
api_version = api_version 、
azure_endpoint = エンドポイント 、
azure_ad_token_provider = トークンプロバイダー 、
)
プロンプト = 「スペインの首都はどこですか?」
範囲 ( 2 ) の i の場合:
開始 = 時間。 perf_counter ()
応答 = クライアント 。チャット 。完成品。作成(
モデル = 展開、
メッセージ = [{ "役割" : "ユーザー" , "コンテンツ" : プロンプト }],
)
経過 = (時間 .perf_counter () - 開始 ) * 1000
print ( f"[call { i + 1 } ] { 経過時間 :.1f } ms — { 応答 . 選択肢 [ 0 ]. メッセージ . コンテンツ } " )
print ( キャッシュ . get_stats () . to_dict ())
キャッシュ。やめて（）
これは Microsoft Entra ID ( DefaultAzureCredential ) で認証し (API キーは必要ありません)、cache_scope=CacheScope.HOST を使用するため、同じ Azure リソース上のすべてのデプロイメントが 1 つのキャッシュを共有します。 API キー認証も機能します。Khazad は、認証方法やホストではなく、リクエスト パス ( /chat/completions ) と一致します。
インポート時間
openaiインポートからOpenAI
カザドから カザドを輸入
キャッシュ = Khazad ( redis_url = "redis://localhost:6379" 、しきい値 = 0.90 、名前空間 = "ollama_example" )
client = OpenAI (base_url = "http://localhost:11434/v1" 、api_key = "ollama")
モデル = "ラマ 3"
プロンプト = 「スペインの首都はどこですか?」
範囲 ( 2 ) の i の場合:
開始 = 時間。 perf_counter ()
応答 = クライアント 。チャット 。完成品。作成(
モデル = モデル、
メッセージ = [{ "ロール" : "ユーザー" ,

"コンテンツ" : プロンプト }],
)
経過 = (時間 .perf_counter () - 開始 ) * 1000
print ( f"[call { i + 1 } ] { 経過時間 :.1f } ms — { 応答 . 選択肢 [ 0 ]. メッセージ . コンテンツ } " )
print ( キャッシュ . get_stats () . to_dict ())
キャッシュ。やめて（）
URL パスが /chat/completions または /responses で終わるホストはすべてキャッシュされます。 vLLM ( http://host:8000/v1/... )、Ollama ( http://localhost:11434/v1/... )、Mistral などをカバーします。
OSをインポートする
インポート時間
anthropic インポートより Anthropic
カザドから カザドを輸入
キャッシュ = Khazad ( redis_url = "redis://localhost:6379" 、しきい値 = 0.90 、名前空間 = "anthropic_example" )
client = Anthropic ( api_key = os . environ [ "ANTHROPIC_API_KEY" ])
モデル = "クロード-俳句-4-5-20251001"
プロンプト = 「フランスの首都はどこですか?」
範囲 ( 2 ) の i の場合:
開始 = 時間。 perf_counter ()
メッセージ = クライアント 。メッセージ。作成(
モデル = モデル、
max_tokens = 256 、
メッセージ = [{ "役割" : "ユーザー" , "コンテンツ" : プロンプト }],
)
経過 = (時間 .perf_counter () - 開始 ) * 1000
print ( f"[call { i + 1 } ] { 経過時間 :.1f } ms — { message . content [ 0 ]. text } " )
print ( キャッシュ . get_stats () . to_dict ())
キャッシュ。やめて（）
api.anthropic.com/v1/messages と一致します。ストリーミング応答は SSE としてキャッシュから再生されます。
OSをインポートする
インポート時間
Googleインポートgenaiより
カザドから カザドを輸入
キャッシュ = Khazad ( redis_url = "redis://localhost:6379" 、しきい値 = 0.90 )
クライアント＝ゲンナイ。クライアント ( api_key = os . environ [ "GEMINI_API_KEY" ])
範囲 ( 2 ) の i の場合:
開始 = 時間。 perf_counter ()
応答 = クライアント 。モデル。生成_コンテンツ (
モデル = "gemini-2.5-フラッシュ" 、
内容 = 「イタリアの首都はどこですか?」 、
)
経過 = (時間 .perf_counter () - 開始 ) * 1000
print ( f"[call { i + 1 } ] { 経過時間 :.1f } ミリ秒 — { 応答 . テキスト } ..." )
print ( キャッシュ . get_stats () . to_dict ())
キャッシュ。セント

オプ()
Generative language.googleapis.com/*/models/*:generateContent と一致します。 Gemini ストリーミング (:streamGenerateContent) はキャッシュされずにパススルーされます。
プロバイダー
URL パターンが一致しました
ストリーミング
OpenAI チャットの完了
任意のホスト、/chat/completions で終わるパス
キャッシュ + 再生
OpenAI レスポンス API
任意のホスト、パス終了/応答
キャッシュ + 再生
Azure OpenAI
チャット/補完マッチャーでカバーされる
キャッシュ + 再生
OpenAI互換プロキシ
チャット/補完マッチャーでカバーされる
キャッシュ + 再生
人間的
api.anthropic.com/v1/messages
キャッシュ + 再生
Google ジェミニ
Generative language.googleapis.com/*:generateContent
パススルー
APIリファレンス
このモジュールは 5 つの関数をシングルトン API として公開します。 Khazad クラスは、インスタンス メソッドと同じ表面を公開します ( init はなく、インスタンス化がそれを行います)。
khazad.init() を使用しても Khazad(...) を使用しても、すべてのパラメータは同じです。
カザド (
redis_url = "redis://localhost:6379" ,
しきい値 = 0.90 、
ttl = 3600 、
名前空間 = "カザド" 、
embedder = "ハグフェイス" ,
embedding_model = "redis/langcache-embed-v2" ,
log_level = "情報" ,
ホスト = なし 、
キャッシュスコープ = "モデル"
)
パラメータ
デフォルト
説明
redis_url
「redis://localhost:6379」
ベクターとキャッシュされた応答を保存する Redis 8 インスタンスの接続 URL

[切り捨てられた]

## Original Extract

Transparent, transport-layer semantic cache for LLM API calls, powered by Redis 8 Vector Sets. - GuglielmoCerri/khazad

GitHub - GuglielmoCerri/khazad: Transparent, transport-layer semantic cache for LLM API calls, powered by Redis 8 Vector Sets. · GitHub
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
GuglielmoCerri
/
khazad
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
33 Commits 33 Commits .github/ workflows .github/ workflows docs/ _static docs/ _static examples examples khazad khazad tests tests .gitignore .gitignore AGENT.md AGENT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Transparent, transport-layer semantic cache for LLM API calls powered by Redis Vector Sets.
~50% fewer API calls · ~96% faster on hits · ~50% lower spend · 100% transparent
Illustrative figures at a 0.50 hit rate (280ms cache replay vs. ~7900ms upstream call). Your numbers depend on traffic.
Khazad intercepts LLM HTTP traffic at the transport layer and serves semantically equivalent requests from a Redis vector cache, with zero changes to your application code .
Model-aware — each (provider, model) pair gets its own vector set, so a gpt-4o answer is never served to a gpt-4o-mini call, no matter how similar the prompt. Set cache_scope="host" to scope by provider host only , letting every model or deployment on the same provider share one cache (different providers stay isolated — see Configuration ).
Conversation-aware — the whole message list (system, user, assistant) is embedded, not just the last user turn. Two different conversations ending with the same follow-up question ("What about its population?") never collide.
Streaming both ways — cache hits replay as real SSE streams (sync and async clients); cache misses that stream are captured chunk-by-chunk with no added latency and reassembled into a canonical JSON response, so a streamed answer can later serve a non-streamed request and vice versa. Aborted streams are never cached.
Semantic caching trades exactness for cost and latency. Know the trade before turning it on. Use it when you have:
High-volume, repetitive traffic: FAQ bots, support assistants, RAG front-ends where many users ask near-identical questions
Dev / test / CI environments — stop paying for the same prompt on every run
Demos and load tests where deterministic, instant responses are a feature
Cost ceilings on internal tools
Privacy : prompts are embedded and responses are stored in clear text in Redis . If prompts may contain PII or secrets, set a ttl , enable Redis AUTH/TLS, and treat the Redis instance with the same care as your logs.
Process-wide patch : Khazad wraps every httpx.Client / AsyncClient created after init() — non-LLM httpx traffic passes through untouched, but the patch is process-global. Call stop() on shutdown. Use hosts=[...] to restrict interception to the endpoints you actually want cached.
httpx-only : SDKs built on httpx are covered (OpenAI, Anthropic, Gemini via google-genai , Mistral, and most proxies). SDKs using requests , aiohttp , or boto3 (AWS Bedrock) are not intercepted.
Single process : the patch lives in the Python process that called init() . Multiple workers share the Redis cache but each needs its own init() .
False-positive control : start at threshold=0.90 and raise it if you see wrong hits. Watch avg_hit_similarity in get_stats() — if it sits near your threshold, your traffic may be too diverse to cache safely.
Redis 8 (Vector Sets support required)
docker run -d --name redis8 -p 6379:6379 redis:8
Installation
uv add khazad
For the OpenAI embedding backend (optional):
uv add khazad[openai-embeddings]
Local / development install:
git clone https://github.com/GuglielmoCerri/khazad.git
cd khazad
uv sync --group dev
uv sync reads pyproject.toml , creates .venv if it doesn't exist, and installs the project itself in editable mode — no separate pip install -e . needed.
To use the local checkout from another project:
uv add --editable /path/to/khazad
Usage
Use the Khazad class directly when you need explicit control over the instance, useful in long-running services, tests, or dependency injection:
from khazad import Khazad
cache = Khazad (
redis_url = "redis://localhost:6379" ,
threshold = 0.90 ,
ttl = 3600 ,
log_level = "DEBUG" ,
)
print ( cache . is_active ()) # True
print ( cache . get_stats ()) # Stats(total_requests=0, ...)
Available functions: init() , stop() , get_stats() , flush() , is_active() . See API Reference for details.
Khazad activates once and intercepts every LLM SDK that uses httpx underneath, no per-provider wiring needed. For further examples see the examples folder .
import os
import time
from openai import OpenAI
from khazad import Khazad
cache = Khazad ( redis_url = "redis://localhost:6379" , threshold = 0.90 , log_level = "DEBUG" )
client = OpenAI ( api_key = os . environ [ "OPENAI_API_KEY" ])
prompt = "What is the capital of Italy?"
for i in range ( 2 ):
start = time . perf_counter ()
response = client . chat . completions . create (
model = "gpt-4o" ,
messages = [{ "role" : "user" , "content" : prompt }],
)
elapsed = ( time . perf_counter () - start ) * 1000
print ( f"[call { i + 1 } ] { elapsed :.1f } ms — { response . choices [ 0 ]. message . content } " )
print ( cache . get_stats (). to_dict ())
cache . stop ()
Matches */chat/completions and */responses paths. Streaming requests also cached.
import os
import time
from azure . identity import DefaultAzureCredential , get_bearer_token_provider
from openai import AzureOpenAI
from khazad import CacheScope , Khazad
cache = Khazad (
redis_url = "redis://localhost:6379" ,
threshold = 0.90 ,
cache_scope = CacheScope . HOST ,
namespace = "azure_openai_example" ,
)
endpoint = os . environ [ "AZURE_OPENAI_ENDPOINT" ]
deployment = os . environ . get ( "AZURE_OPENAI_DEPLOYMENT" , "gpt-4.1" )
token_provider = get_bearer_token_provider (
DefaultAzureCredential (), "https://cognitiveservices.azure.com/.default"
)
api_version = os . environ . get ( "AZURE_OPENAI_API_VERSION" , "2024-12-01-preview" )
client = AzureOpenAI (
api_version = api_version ,
azure_endpoint = endpoint ,
azure_ad_token_provider = token_provider ,
)
prompt = "What is the capital of Spain?"
for i in range ( 2 ):
start = time . perf_counter ()
response = client . chat . completions . create (
model = deployment ,
messages = [{ "role" : "user" , "content" : prompt }],
)
elapsed = ( time . perf_counter () - start ) * 1000
print ( f"[call { i + 1 } ] { elapsed :.1f } ms — { response . choices [ 0 ]. message . content } " )
print ( cache . get_stats (). to_dict ())
cache . stop ()
It authenticates with Microsoft Entra ID ( DefaultAzureCredential ) — no API key needed — and uses cache_scope=CacheScope.HOST so every deployment on the same Azure resource shares one cache. API-key auth works too: Khazad matches the request path ( /chat/completions ), not the auth method or host.
import time
from openai import OpenAI
from khazad import Khazad
cache = Khazad ( redis_url = "redis://localhost:6379" , threshold = 0.90 , namespace = "ollama_example" )
client = OpenAI ( base_url = "http://localhost:11434/v1" , api_key = "ollama" )
model = "llama3"
prompt = "What is the capital of Spain?"
for i in range ( 2 ):
start = time . perf_counter ()
response = client . chat . completions . create (
model = model ,
messages = [{ "role" : "user" , "content" : prompt }],
)
elapsed = ( time . perf_counter () - start ) * 1000
print ( f"[call { i + 1 } ] { elapsed :.1f } ms — { response . choices [ 0 ]. message . content } " )
print ( cache . get_stats (). to_dict ())
cache . stop ()
Any host whose URL path ends with /chat/completions or /responses is cached. Covers vLLM ( http://host:8000/v1/... ), Ollama ( http://localhost:11434/v1/... ), Mistral, etc.
import os
import time
from anthropic import Anthropic
from khazad import Khazad
cache = Khazad ( redis_url = "redis://localhost:6379" , threshold = 0.90 , namespace = "anthropic_example" )
client = Anthropic ( api_key = os . environ [ "ANTHROPIC_API_KEY" ])
model = "claude-haiku-4-5-20251001"
prompt = "What is the capital of France?"
for i in range ( 2 ):
start = time . perf_counter ()
message = client . messages . create (
model = model ,
max_tokens = 256 ,
messages = [{ "role" : "user" , "content" : prompt }],
)
elapsed = ( time . perf_counter () - start ) * 1000
print ( f"[call { i + 1 } ] { elapsed :.1f } ms — { message . content [ 0 ]. text } " )
print ( cache . get_stats (). to_dict ())
cache . stop ()
Matches api.anthropic.com/v1/messages . Streaming responses replayed from cache as SSE.
import os
import time
from google import genai
from khazad import Khazad
cache = Khazad ( redis_url = "redis://localhost:6379" , threshold = 0.90 )
client = genai . Client ( api_key = os . environ [ "GEMINI_API_KEY" ])
for i in range ( 2 ):
start = time . perf_counter ()
response = client . models . generate_content (
model = "gemini-2.5-flash" ,
contents = "What is the capital of Italy?" ,
)
elapsed = ( time . perf_counter () - start ) * 1000
print ( f"[call { i + 1 } ] { elapsed :.1f } ms — { response . text } ..." )
print ( cache . get_stats (). to_dict ())
cache . stop ()
Matches generativelanguage.googleapis.com/*/models/*:generateContent . Gemini streaming ( :streamGenerateContent ) passes through uncached.
Provider
URL pattern matched
Streaming
OpenAI Chat Completions
any host, path ending /chat/completions
cached + replayed
OpenAI Responses API
any host, path ending /responses
cached + replayed
Azure OpenAI
covered by chat/completions matcher
cached + replayed
OpenAI-compatible proxies
covered by chat/completions matcher
cached + replayed
Anthropic
api.anthropic.com/v1/messages
cached + replayed
Google Gemini
generativelanguage.googleapis.com/*:generateContent
pass-through
API Reference
The module exposes five functions as the singleton API. The Khazad class exposes the same surface as instance methods (no init , instantiation does that).
All parameters are the same whether you use khazad.init() or Khazad(...) :
Khazad (
redis_url = "redis://localhost:6379" ,
threshold = 0.90 ,
ttl = 3600 ,
namespace = "khazad" ,
embedder = "huggingface" ,
embedding_model = "redis/langcache-embed-v2" ,
log_level = "INFO" ,
hosts = None ,
cache_scope = "model"
)
Parameter
Default
Description
redis_url
"redis://localhost:6379"
Connection URL for the Redis 8 instance that stores vectors and cached response

[truncated]
