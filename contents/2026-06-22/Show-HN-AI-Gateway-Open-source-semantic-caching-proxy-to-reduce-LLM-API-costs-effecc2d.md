---
source: "https://github.com/Arnab758/ai-gateway"
hn_url: "https://news.ycombinator.com/item?id=48633737"
title: "Show HN: AI-Gateway – Open-source semantic caching proxy to reduce LLM API costs"
article_title: "GitHub - Arnab758/ai-gateway: AI-Gateway reverse proxy that uses semantic caching and aims to reduce LLM API bills and token costs by 40-70%. · GitHub"
author: "arnab777"
captured_at: "2026-06-22T18:21:59Z"
capture_tool: "hn-digest"
hn_id: 48633737
score: 1
comments: 0
posted_at: "2026-06-22T18:10:05Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI-Gateway – Open-source semantic caching proxy to reduce LLM API costs

- HN: [48633737](https://news.ycombinator.com/item?id=48633737)
- Source: [github.com](https://github.com/Arnab758/ai-gateway)
- Score: 1
- Comments: 0
- Posted: 2026-06-22T18:10:05Z

## Translation

タイトル: Show HN: AI-Gateway – LLM API コストを削減するオープンソースのセマンティック キャッシング プロキシ
記事のタイトル: GitHub - Arnab758/ai-gateway: セマンティック キャッシュを使用し、LLM API の料金とトークンのコストを 40 ～ 70% 削減することを目的とした AI-Gateway リバース プロキシ。 · GitHub
説明: セマンティック キャッシュを使用し、LLM API の料金とトークンのコストを 40 ～ 70% 削減することを目的とした AI-Gateway リバース プロキシ。 - Arnab758/ai-ゲートウェイ

記事本文:
GitHub - Arnab758/ai-gateway: セマンティック キャッシュを使用し、LLM API の料金とトークンのコストを 40 ～ 70% 削減することを目的とした AI-Gateway リバース プロキシ。 · GitHub
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
アルナブ758
/
aiゲートウェイ
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
24 コミット 24 コミット .github/ workflows .github/ workflows .gitignore .gitignore Dockerfile Dockerfile README.md README.md cache.go queue.gocircuit_breaker.gocircuit_breaker.go config.go config.gocost_tracker.gocost_tracker.godemo.htmldemo.htmldemo_server.godemo_server.go docker-compose.yml docker-compose.yml embedding.go embedding.go ゲートウェイ.yaml ゲートウェイ.yaml go.mod go.mod go.sum go.sum main.go main.go Mirror.go Mirror.go parser.go parser.go proxy.go proxy.go render.yaml render.yaml router.go router.go Shield.go Shield.go template_matcher.go template_matcher.go test_gateway.go test_gateway.go すべてのファイルを表示 リポジトリ ファイルのナビゲーション
コード変更なしで LLM API コストを 40 ～ 70% 削減します。
アプリと AI プロバイダー (OpenAI、Groq など) の間に位置するセマンティック キャッシュ レイヤー。同様の質問を 2 回行うと、API を再度呼び出すのではなく、キャッシュされた回答が即座に返されます。
🎯 これでどんな問題が解決しますか?
あなたは AI アプリを構築していて、API の請求額は月額 500 ドルです。そのうち 40 ～ 70% は繰り返しの質問です。
「RAGって何？」 100 回尋ねられる = 100 回の API 呼び出し
「パスワードをリセットするにはどうすればよいですか?」 50 回要求 = 50 API 呼び出し
AI ゲートウェイを使用すると、150 件の通話が 2 件の通話になります (固有の質問ごとに 1 件)。毎月 200 ～ 350 ドル節約できます。
🚀 60 秒で導入 (3 つのオプション)
オプション 1: 鉄道 (推奨 - Redis を含む)
API キー (Groq または OpenAI) を入力してください
完了！ゲートウェイは https://your-app.up.railway.app で稼働しています。
✅ ホスト型ゲートウェイ (サーバー管理なし)
✅ Redis が含まれています (永続キャッシュ)
オプション 2: レンダリング (ワンクリック展開)
環境変数を追加: UPSTREAM_API_KEY=your_key
注: レンダリング ダッシュボードに Redis アドオンを個別に追加する必要があります。

d.
オプション 3: Docker (セルフホスト)
#1. リポジトリのクローンを作成する
git クローン https://github.com/Arnab758/ai-gateway.git
cd aiゲートウェイ
# 2. API キーを設定する
UPSTREAM_API_KEY=gsk_your_groq_key_here をエクスポート
# 3. すべてを開始します (ゲートウェイ + Redis)
ドッカー構成 -d
#4. 実行中であることを確認する
カール http://localhost:8080/health
# 期待される応答: {"status":"ok"}
それです！ゲートウェイは http://localhost:8080 で実行されています
# ゲートウェイ経由でリクエストを送信する
curl -X POST http://localhost:8080/v1/chat/completions \
-H " Content-Type: application/json " \
-H " X-Gateway-Token: my-app " \
-H " 認証: ベアラー sk-your-openai-or-groq-key " \
-d ' {
"モデル": "gpt-4",
"messages": [{"role": "user", "content": "RAG とは何ですか?"}]
} '
# 同じリクエストを再度送信します
# 応答ヘッダーには次のように表示されます: X-Gateway-Cache: HIT
# お金が貯まったばかりです! 💰
Python の例
インポートリクエスト
# ゲートウェイ URL (鉄道/レンダー/Docker から)
GATEWAY_URL = "https://your-app.up.railway.app"
API_KEY = "sk-your-key"
応答 = リクエスト。ポスト(
f" { GATEWAY_URL } /v1/chat/completions" ,
ヘッダー = {
"Content-Type" : "application/json" ,
"X-ゲートウェイ トークン" : "私のアプリ" 、
"認可" : f"ベアラー { API_KEY } "
}、
json = {
"モデル" : "gpt-4" ,
"messages" : [{ "role" : "user" , "content" : "RAG とは何ですか?" }]
}
)
print (応答.json())
Node.js の例
const response = await fetch ( 'https://your-app.up.railway.app/v1/chat/completions' , {
メソッド: 'POST' 、
ヘッダー: {
'Content-Type' : 'application/json' ,
'X-ゲートウェイ トークン' : 'my-app' ,
'認可' : 'ベアラー sk-your-key'
} 、
本文: JSON 。 stringify ( {
モデル: 'gpt-4' 、
メッセージ : [ { 役割 : 'ユーザー' 、内容 : 'RAG とは何ですか?' } ]
})
} ) ;
const data = 応答を待ちます。 json() ;
コンソール。ログ (データ) ;
🎮 インタラクティブなデモを試してみる
API キーは必要ありません。キャッシングの方法を確認する

作品:
プロンプトを入力して「送信」をクリックします（シミュレーションモード）
または、API キーを入力し、[Test Real API] (Redis を使用した実際のキャッシュ) をクリックします。
キャッシュ ヒットを確認するには、同じプロンプトを 2 回送信してみてください。
セマンティック キャッシュ - 完全な重複だけでなく、類似した質問と一致します。
「RAGって何？」 =「RAGの説明」=「RAGの定義」
マルチテナント - 各顧客は独自の分離されたキャッシュを取得します
4層マッチング:
完全一致 (100% 同一)
テンプレート一致 (「ロンドンの天気」 = 「パリの天気」)
意味上の一致 (同様の意味)
単語の重複（部分一致）
Redis + インメモリフォールバック - Redis の有無にかかわらず動作します
リクエストの重複排除 - 100 個の同時同一リクエスト = 1 回の API 呼び出し
レート制限 - テナントごとの不正行為を防止します
サーキット ブレーカー - プロバイダーがダウンした場合に通話を自動的に停止します。
コスト追跡 - どれだけ節約できたかを確認します
シナリオ: 10,000 人のユーザーがいるカスタマー サポート チャットボット
10,000 人のユーザーがそれぞれ 100 のよくある質問をします
料金: 月額 500 ドル (通話あたり 0.0005 ドル)
最初の 100 問: 100 回の API 呼び出し (キャッシュ ミス)
次に同じ質問をする 9,900 人のユーザー: API 呼び出し 0 (キャッシュ ヒット)
節約: 月額 499.95 ドル (99.99%)
30% がユニークな質問であっても:
Gateway.yaml を編集してカスタマイズします。
キャッシュ:
redis_url : " redis://localhost:6379 " # または Redis URL
ベクトル :
有効 : true
類似性のしきい値 : 0.85 # 85% 類似 = キャッシュ ヒット
ttl_hours : 24 # キャッシュ エントリは 24 時間後に期限切れになります
レートリミッター:
有効 : true
max_requests : 60 # テナントあたり 1 分あたり
📡 API エンドポイント
エンドポイント
方法
説明
/v1/チャット/コンプリート
投稿
キャッシュを備えたメイン プロキシ エンドポイント
/健康
ゲット
健康診断
/統計
ゲット
キャッシュ統計
/メトリクス
ゲット
プロメテウスのメトリクス
🔍モニタリング
カール http://localhost:8080/stats
応答:
{
「稼働時間」: 1234567890 、
"キャッシュ" : {
"local_index_entries" : 150 、
"vector_dimensions" : 128 、
"vector_threshold" : 0.85 、
「jaccard_th」

しきい値 : 0.75 、
"template_enabled" : true 、
"dedup_enabled" : true 、
"ttl_hours" : 24
}
}
応答ヘッダー
すべての応答にはキャッシュ情報が含まれます。
X ゲートウェイ キャッシュ: HIT # または MISS
X ゲートウェイの類似性: 0.95 # 95% 類似 (HIT の場合)
X-Gateway-Time-Saved: 1234ms # 時間の節約 (HIT の場合)
🐛 トラブルシューティング
問題: 「Redis 接続に失敗しました」
解決策: Redis はオプションです。ゲートウェイは自動的にメモリ内キャッシュにフォールバックします。本番環境では、Redis を追加します。
鉄道: 「新規」ボタンから Redis を追加します
レンダリング：「新規」→「データベース」→「Redis」からRedisを追加します。
Docker: docker-compose.yml にすでに含まれています
問題: 「すべての上流プロバイダーが利用できません」
原因: 無料枠のレート制限に達しています (Groq/OpenAI)
1 ～ 2 分待ってからもう一度お試しください
有料レベルへのアップグレード (リクエストあたり $0.002 対無料制限)
より高い制限を設定した独自の API キーを追加します
問題: 「レート制限を超えました」
原因: 1 つのテナントからのリクエストが多すぎます
解決策:gateway.yaml のレート制限を増やします。
レートリミッター:
max_requests : 120 # 60 から増加
ウィンドウ分 : 1
問題: キャッシュがヒットしない
原因: プロンプトが違いすぎます
解決策: Gateway.yaml の類似性のしきい値を下げます。
キャッシュ:
ベクトル :
類似性閾値 : 0.75 # 0.85 から低下
ジャカード :
しきい値 : 0.65 # 0.75 から低下
🏗️ アーキテクチャ
アプリ→AIゲートウェイ→[キャッシュチェック]→Redis
↓
[キャッシュHIT] → キャッシュされたレスポンスを返す(即時、$0)
↓
[キャッシュミス] → LLMプロバイダ呼び出し → 応答キャッシュ → 返却
🤝 貢献しています
貢献は大歓迎です!お願いします:
MIT ライセンス - 商用目的でご自由に使用してください。
ディスカッション: GitHub ディスカッション
このプロジェクトが役に立った場合は、星を付けてください。他の人がそれを見つけるのに役立ちます。
AI コミュニティ向けに ❤️ で構築
質問がありますか?問題を開いていただければ、24 時間以内に返答いたします。
SEを使用したAI-Gatewayリバースプロキシ

mantic キャッシュを利用し、LLM API の請求額とトークンのコストを 40 ～ 70% 削減することを目指しています。
arnab758.github.io/ai-gateway/
リソース
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

AI-Gateway reverse proxy that uses semantic caching and aims to reduce LLM API bills and token costs by 40-70%. - Arnab758/ai-gateway

GitHub - Arnab758/ai-gateway: AI-Gateway reverse proxy that uses semantic caching and aims to reduce LLM API bills and token costs by 40-70%. · GitHub
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
Arnab758
/
ai-gateway
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
24 Commits 24 Commits .github/ workflows .github/ workflows .gitignore .gitignore Dockerfile Dockerfile README.md README.md cache.go cache.go circuit_breaker.go circuit_breaker.go config.go config.go cost_tracker.go cost_tracker.go demo.html demo.html demo_server.go demo_server.go docker-compose.yml docker-compose.yml embedding.go embedding.go gateway.yaml gateway.yaml go.mod go.mod go.sum go.sum main.go main.go mirror.go mirror.go parser.go parser.go proxy.go proxy.go render.yaml render.yaml router.go router.go shield.go shield.go template_matcher.go template_matcher.go test_gateway.go test_gateway.go View all files Repository files navigation
Cut your LLM API costs by 40-70% with zero code changes.
A semantic caching layer that sits between your app and AI providers (OpenAI, Groq, etc.). When you ask a similar question twice, it returns the cached answer instantly instead of calling the API again.
🎯 What Problem Does This Solve?
You're building an AI app and your API bill is $500/month. 40-70% of that is for repeat questions :
"What is RAG?" asked 100 times = 100 API calls
"How do I reset my password?" asked 50 times = 50 API calls
With AI Gateway: Those 150 calls become 2 calls (one for each unique question). You save $200-350/month.
🚀 Deploy in 60 Seconds (3 Options)
Option 1: Railway (Recommended - Includes Redis)
Enter your API key (Groq or OpenAI)
Done! Your gateway is live at https://your-app.up.railway.app
✅ Hosted gateway (no server management)
✅ Redis included (persistent cache)
Option 2: Render (One-Click Deploy)
Add environment variable: UPSTREAM_API_KEY=your_key
Note: You'll need to add a Redis addon separately in Render dashboard.
Option 3: Docker (Self-Hosted)
# 1. Clone the repo
git clone https://github.com/Arnab758/ai-gateway.git
cd ai-gateway
# 2. Set your API key
export UPSTREAM_API_KEY=gsk_your_groq_key_here
# 3. Start everything (gateway + Redis)
docker compose up -d
# 4. Verify it's running
curl http://localhost:8080/health
# Expected response: {"status":"ok"}
That's it! Your gateway is now running at http://localhost:8080
# Send a request through the gateway
curl -X POST http://localhost:8080/v1/chat/completions \
-H " Content-Type: application/json " \
-H " X-Gateway-Token: my-app " \
-H " Authorization: Bearer sk-your-openai-or-groq-key " \
-d ' {
"model": "gpt-4",
"messages": [{"role": "user", "content": "What is RAG?"}]
} '
# Send the SAME request again
# Response headers will show: X-Gateway-Cache: HIT
# You just saved money! 💰
Python Example
import requests
# Your gateway URL (from Railway/Render/Docker)
GATEWAY_URL = "https://your-app.up.railway.app"
API_KEY = "sk-your-key"
response = requests . post (
f" { GATEWAY_URL } /v1/chat/completions" ,
headers = {
"Content-Type" : "application/json" ,
"X-Gateway-Token" : "my-app" ,
"Authorization" : f"Bearer { API_KEY } "
},
json = {
"model" : "gpt-4" ,
"messages" : [{ "role" : "user" , "content" : "What is RAG?" }]
}
)
print ( response . json ())
Node.js Example
const response = await fetch ( 'https://your-app.up.railway.app/v1/chat/completions' , {
method : 'POST' ,
headers : {
'Content-Type' : 'application/json' ,
'X-Gateway-Token' : 'my-app' ,
'Authorization' : 'Bearer sk-your-key'
} ,
body : JSON . stringify ( {
model : 'gpt-4' ,
messages : [ { role : 'user' , content : 'What is RAG?' } ]
} )
} ) ;
const data = await response . json ( ) ;
console . log ( data ) ;
🎮 Try the Interactive Demo
No API key needed! See how caching works:
Type a prompt and click "Send" (simulation mode)
Or enter your API key and click "Test Real API" (real caching with Redis)
Try sending the same prompt twice to see cache hits!
Semantic Caching - Matches similar questions, not just exact duplicates
"What is RAG?" = "Explain RAG" = "RAG definition"
Multi-Tenant - Each customer gets their own isolated cache
4-Tier Matching:
Exact match (100% identical)
Template match ("weather in London" = "weather in Paris")
Semantic match (similar meaning)
Word overlap (partial matches)
Redis + In-Memory Fallback - Works with or without Redis
Request Deduplication - 100 concurrent identical requests = 1 API call
Rate Limiting - Prevent abuse per tenant
Circuit Breaker - Automatically stops calling if provider is down
Cost Tracking - See how much you saved
Scenario: Customer support chatbot with 10,000 users
10,000 users ask 100 common questions each
Cost: $500/month (at $0.0005/call)
First 100 questions: 100 API calls (cache miss)
Next 9,900 users asking same questions: 0 API calls (cache hit)
Savings: $499.95/month (99.99%)
Even with 30% unique questions:
Edit gateway.yaml to customize:
cache :
redis_url : " redis://localhost:6379 " # Or your Redis URL
vector :
enabled : true
similarity_threshold : 0.85 # 85% similar = cache hit
ttl_hours : 24 # Cache entries expire after 24 hours
rate_limiter :
enabled : true
max_requests : 60 # Per minute per tenant
📡 API Endpoints
Endpoint
Method
Description
/v1/chat/completions
POST
Main proxy endpoint with caching
/health
GET
Health check
/stats
GET
Cache statistics
/metrics
GET
Prometheus metrics
🔍 Monitoring
curl http://localhost:8080/stats
Response:
{
"uptime" : 1234567890 ,
"cache" : {
"local_index_entries" : 150 ,
"vector_dimensions" : 128 ,
"vector_threshold" : 0.85 ,
"jaccard_threshold" : 0.75 ,
"template_enabled" : true ,
"dedup_enabled" : true ,
"ttl_hours" : 24
}
}
Response Headers
Every response includes cache information:
X-Gateway-Cache: HIT # or MISS
X-Gateway-Similarity: 0.95 # 95% similar (if HIT)
X-Gateway-Time-Saved: 1234ms # Time saved (if HIT)
🐛 Troubleshooting
Problem: "Redis connection failed"
Solution: Redis is optional! The gateway will fall back to in-memory cache automatically. For production, add Redis:
Railway: Add Redis from the "New" button
Render: Add Redis from the "New" → "Database" → "Redis"
Docker: Already included in docker-compose.yml
Problem: "All upstream providers unavailable"
Cause: You're hitting rate limits on free tier (Groq/OpenAI)
Wait 1-2 minutes and try again
Upgrade to paid tier ($0.002/request vs free limits)
Add your own API key with higher limits
Problem: "Rate limit exceeded"
Cause: Too many requests from one tenant
Solution: Increase rate limits in gateway.yaml :
rate_limiter :
max_requests : 120 # Increase from 60
window_minutes : 1
Problem: Cache not hitting
Cause: Prompts are too different
Solution: Lower the similarity threshold in gateway.yaml :
cache :
vector :
similarity_threshold : 0.75 # Lower from 0.85
jaccard :
threshold : 0.65 # Lower from 0.75
🏗️ Architecture
Your App → AI Gateway → [Cache Check] → Redis
↓
[Cache HIT] → Return cached response (instant, $0)
↓
[Cache MISS] → Call LLM Provider → Cache response → Return
🤝 Contributing
Contributions are welcome! Please:
MIT License - feel free to use this commercially!
Discussions: GitHub Discussions
If this project helps you, please give it a star! It helps others find it.
Built with ❤️ for the AI community
Questions? Open an issue and I'll respond within 24 hours.
AI-Gateway reverse proxy that uses semantic caching and aims to reduce LLM API bills and token costs by 40-70%.
arnab758.github.io/ai-gateway/
Resources
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
