---
source: "https://github.com/SugaC-275/ToTra"
hn_url: "https://news.ycombinator.com/item?id=48420028"
title: "ToTra – open-source LLM gateway with GDPR/EU AI Act compliance"
article_title: "GitHub - SugaC-275/ToTra: Open-source AI gateway for enterprises — quota, PII protection, cost tracking, and compliance · GitHub"
author: "SugaC275"
captured_at: "2026-06-06T00:54:35Z"
capture_tool: "hn-digest"
hn_id: 48420028
score: 2
comments: 0
posted_at: "2026-06-06T00:16:10Z"
tags:
  - hacker-news
  - translated
---

# ToTra – open-source LLM gateway with GDPR/EU AI Act compliance

- HN: [48420028](https://news.ycombinator.com/item?id=48420028)
- Source: [github.com](https://github.com/SugaC-275/ToTra)
- Score: 2
- Comments: 0
- Posted: 2026-06-06T00:16:10Z

## Translation

タイトル: ToTra – GDPR/EU AI 法に準拠したオープンソース LLM ゲートウェイ
Article title: GitHub - SugaC-275/ToTra: Open-source AI gateway for enterprises — quota, PII protection, cost tracking, and compliance · GitHub
説明: 企業向けのオープンソース AI ゲートウェイ - クォータ、PII 保護、コスト追跡、コンプライアンス - SugaC-275/ToTra

記事本文:
GitHub - SugaC-275/ToTra: 企業向けオープンソース AI ゲートウェイ — クォータ、PII 保護、コスト追跡、コンプライアンス · GitHub
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
シュガC-275
/
トトラ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビ

オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
255 コミット 255 コミット .github .github 管理者 管理者 ダッシュボード ダッシュボード ドキュメント ドキュメント ゲートウェイ ゲートウェイ インフラ インフラ パーサー パーサー スクリプト スクリプト SDK SDK Web サイト Web サイト .env.example .env.example .gitignore .gitignore Makefile Makefile README.md README.md docker-compose.dev.yml docker-compose.dev.yml docker-compose.yml docker-compose.yml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AIゲートウェイとガバナンスプラットフォーム
Go で書かれたオープンソース LLM プロキシ。クォータの強制、PII ブロック、コスト追跡、LLM へのコンプライアンスを 1 行のコードで追加します。
クイックスタート ·
統合ガイド ·
特徴・
建築・
ゲートウェイのドキュメント ·
管理者 API ·
ディスカッション
ToTra は、LLM プロバイダーの前面に位置するオープンソースの AI ゲートウェイおよびガバナンス プラットフォームです。
既存のアプリを OpenAI、Anthropic、またはその他のプロバイダーではなく ToTra に指定すると、次のものがすぐに得られます。
クォータの強制 — ユーザーごとおよびチームごとのハード予算の上限
PII ブロック — データがネットワークから流出する前にエッジで 18 の言語グループをスキャンします
コスト追跡 - ユーザーごと、チームごと、モデルごとのトークンと米ドルの支出とチャージバック レポート
コンプライアンス — GDPR ワークフロー、EU AI 法のチェックリスト、ハッシュチェーンされた不変監査ログ
コード変更なし — 100% OpenAI 互換。設定内の 1 行を交換します
フローチャート LR
A["🖥️ あなたのアプリ\n(OpenAI SDK/curl\n/ LangChain)"] -->|"1 · API リクエスト"| B
サブグラフ B["ToTra ゲートウェイ:8080"]
TB方向
B1["🔑認証とAPIキー"]
B2["📊 クォータ チェック\n(ユーザー/チームごと)"]
B3["🔒 PII スキャン\n(18 言語)"]
B4[「⚡セマンティックキャッシュ」]
B5[「🔀ルートとロードバランス」]
B1 --> B2 --> B3 --> B4 --> B5
終わり
B -->|"2・転送リクエスト"| C["☁️ LLM プロバイダー\nOpenAI · Anthropic\nGemini · Mistral · Azure\

「ベッドロック・オラマ」】
C -->|"3・応答"|あ
B -->|"4・使用イベント"| D
サブグラフ D["ToTra Admin :8081"]
TB方向
D1["💸コスト追跡"]
D2["📋 コンプライアンスと監査"]
D3["🔔予算アラート"]
終わり
D --> E["📊 ダッシュボード :3000\n管理コンソール · レポート\n従業員セルフサービス"]
読み込み中
なぜトトラなのか
🚀 Go で書かれています — < 2 ミリ秒の p95 オーバーヘッド。ネイティブ バイナリ、Python ランタイム、ウォームアップなし。
🔒 PII はエッジでブロックされます — 電子メール、ID、クレジット カード、18 の言語グループにわたる健康記録。機密データは、LLM に到達する前に編集されます。
💸 厳しい予算の上限 — 制限を超えるリクエストはプロバイダーに接続する前に 429 になります。リアルタイムの Slack / Webhook アラート。
📋 すぐに使えるコンプライアンス — GDPR データ主体のワークフロー、EU AI 法のチェックリスト、およびすべてのリクエストに対する不変のハッシュチェーン監査ログ。
📊 財務対応のレポート — 部門のチャージバック CSV、予算予測、支出の異常検出。
🏠 自己ホスト型 — キー、インフラストラクチャ、データ。外部依存性はありません。
前提条件: Docker + Docker Compose
git クローン https://github.com/SugaC-275/ToTra.git
cdとら
cp .env.example .env # プロバイダーの API キーを入力します
docker-compose --profile app up -d --wait
http://localhost:3000 を開いてサインインします。
最初のログイン直後に、「設定」→「セキュリティ」からデフォルトの資格情報を変更します。
一行変更。コードの 1 行おきは同じままです。
輸入オープンアイ
# Before — OpenAI を直接呼び出します
クライアント = openai 。 OpenAI ( api_key = "sk-..." )
# 後 — ToTra を経由してルーティングします (その他の変更はゼロ)
クライアント = openai 。 OpenAI (
api_key = "your-totra-api-key" , # ToTra管理パネルから発行
Base_url = "http://your-totra-host:8080/v1"
)
応答 = クライアント 。チャット。完成品。作成(
モデル = "gpt-4o" 、
メッセージ = [{ "役割" : "ユーザー" , "コンテンツ" : "こんにちは!" }]
)

print (応答の選択肢「0」のメッセージの内容)
Node.js / TypeScript (OpenAI SDK)
「openai」から OpenAI をインポートします。
const client = 新しい OpenAI ( {
apiKey : "あなたのtotra-APIキー" ,
BaseURL : "http://your-totra-host:8080/v1" ,
} ) ;
const 応答 = クライアントを待ちます。チャット 。完成品。作成 ( {
モデル = "gpt-4o" 、
メッセージ : [ { 役割 : "ユーザー" 、内容 : "こんにちは!" } ]、
} ) ;
コンソール。 log (応答 . 選択肢 [ 0 ] . メッセージ . 内容 ) ;
カールする
カール http://your-totra-host:8080/v1/chat/completions \
-H " 認証: ベアラー your-totra-api-key " \
-H " Content-Type: application/json " \
-d ' {
"モデル": "gpt-4o",
"メッセージ": [{"役割": "ユーザー", "コンテンツ": "こんにちは!"}]
} '
ラングチェーン
langchain_openai から ChatOpenAI をインポート
llm = ChatOpenAI (
モデル = "gpt-4o" 、
openai_api_key = "あなたのtotra-api-key" ,
openai_api_base = "http://your-totra-host:8080/v1" ,
)
応答 = llm 。呼び出す (「こんにちは!」)
print (応答の内容)
接続されると、すべてのリクエストはクォータの適用、PII スキャン、セマンティック キャッシュ、コスト追跡を通じて自動的にルーティングされます。
🔒 PII 保護 — 18 の言語グループ
すべてのリクエスト本文は、LLM に到達する前にリアルタイムでスキャンされます。検出された PII は編集され、イベントが記録されます。ブロックされたリクエストは 422 を返します。
管理パネルでチームごと、モデルごと、またはグローバルにルールを構成します。
ユーザーごと、チームごと、モデルごとのトークンと米ドルのコスト追跡
厳しい予算の上限 - 制限を超えるリクエストはプロバイダーに接続する前に 429 を取得します
Slack / Feishu / Webhook 通知でアラートしきい値を構成可能
現在の燃焼率に基づいた月次予算予測
財務用の CSV エクスポートによる部門チャージバック レポート
調達分析と ROI ダッシュボード
自動アラートによる支出の異常検出
ダッシュボード → コスト → レポート → CSV エクスポート
📋 コンプライアンスと監査
GDPR

— データ主体のエクスポートおよび削除リクエストのワークフロー、構成可能な保存ポリシー
EU AI 法 — モデルごとのステータス追跡を含むコンプライアンス チェックリスト
不変の監査チェーン - すべてのリクエストはハッシュチェーン化されます。ログは改ざんできません
SIEM 統合 — セキュリティ イベント転送用の構成可能な Webhook ターゲット
データ常駐制御 - すべてのデータをオンプレミスまたは特定のリージョンに保持します
OpenAI 互換 — OpenAI API のドロップイン置換 ( /v1/chat/completions 、 /v1/embeddings 、ストリーミング)
Anthropic 互換 — ネイティブ Anthropic メッセージ API サポート
マルチプロバイダー ルーティング - プロバイダーとモデル間の自動フォールバック
セマンティック キャッシュ — SimHash LSH 重複排除。プロンプトが繰り返されると LLM が完全にスキップされます
即時圧縮 — 長いコンテキストでのトークンの消費を削減します
ストリーミング プロキシ — フルテキスト/イベント ストリームのサポート
ファイル パイプライン — 1 回の API 呼び出しで PDF / DOCX / PPTX をアップロード→解析→チャット
レート制限、IP許可リスト、APIキー認証
JWT認証 + OIDC / SSO統合
役割ベースのアクセス制御 (管理者/従業員)
クォータ要求/承認ワークフローによるユーザーとチームの管理
モデル カタログ — チームごとにプロバイダーを有効化、無効化、構成する
ボット通知 — Slack、Feishu、カスタム Webhook
HR同期コネクタ（CSVインポート）
エージェント セッションの追跡 - デッド ループのエージェント セッションを自動的に検出して終了します
プロバイダー
チャット
埋め込み
ストリーミング
ファイル
OpenAI (GPT-4o、o1、o3、o4)
✅
✅
✅
✅
人間性 (クロード 3.5、4)
✅
—
✅
✅
Google ジェミニ
✅
✅
✅
—
ミストラルAI
✅
✅
✅
—
メタ・ラマ (オラマ経由)
✅
✅
✅
—
コヒアコマンド
✅
✅
✅
—
Azure OpenAI
✅
✅
✅
✅
AWS の基盤
✅
✅
✅
—
地元 / オラマ
✅
✅
✅
—
任意の OpenAI 互換エンドポイント
✅
✅
✅
—
パフォーマンス
ToTraは完全に書かれています

囲碁で。ゲートウェイは、運用負荷の下で p95 で 2 ミリ秒未満のオーバーヘッドを追加します。
100 ミリ秒のモックアップストリームに対して測定。ベンチマークを再現 →
あなたのアプリ (OpenAI SDK/curl/LangChain/任意の HTTP クライアント)
│
▼
トトラゲートウェイ：8080
認証、クォータ、PII スキャン、ポリシー、セマンティック キャッシュ、ルーティング
│
▼
OpenAI · 人間性 · ジェミニ · ミストラル · ローカルモデル
│
│（利用イベント）
▼
トトラ管理者：8081
コスト、コンプライアンス、予算、監査証跡、通知
│
▼
ダッシュボード:3000
管理コンソール · 部門レポート · 従業員セルフサービス
サービス
スタック
港
ゲートウェイ
Go 1.25 / ファイバー
8080
管理者
Go 1.25 / ファイバー
8081
パーサー
Python 3.12 / FastAPI
8090
ダッシュボード
リアクト 19 / ヴァイト
3000
ポストグレ
PostgreSQL 16
5432
レディス
Redis 7
6379
スクリーンショット
コストダッシュボード
部門レポート
ユーザー管理
従業員セルフサービス
地域開発
# 1. データベースを起動する
docker-compose up -d postgres redis
# 2. 各サービスを独自のターミナルで実行する
cd ゲートウェイ && 実行します。
cd admin && 実行します。
cd パーサー && uvicorn main:app --port 8090
cd ダッシュボード && npm install && npm run dev
# 3. シード開発資格情報 (初回のみ)
cd スクリプト/set-dev-passwords
POSTGRES_HOST=ローカルホスト POSTGRES_DB=トトラ \
POSTGRES_USER=totra POSTGRES_PASSWORD=totra_secret 実行してください。
デフォルトの開発者認証情報: admin@acme.com / totra123
.env.example を .env にコピーします。主要な変数:
Redis、SMTP、通知設定を含む完全なリストについては、.env.example を参照してください。
テストを行う
# サービスごと
cd ゲートウェイ && テストに行く ./...
cd admin && go test ./...
cd ダッシュボード && npm run test:run
cd パーサー && pytest
貢献する
バグ修正、新しいプロバイダーの統合、ドキュメントの改善、機能リクエストなどの貢献を歓迎します。
git クローン https://github.com/SugaC-275/ToTra.git
cdとら
# 送信する前にテストを実行する
テストを行う
リポジトリをフォークして b を作成します

メインからの牧場
変更を加え、必要に応じてテストを追加します
より大きな機能については、まずディスカッションを開いて方向性を調整します。
💬 GitHub ディスカッション — 質問、アイデア、見せて伝える
MIT — 自由に使用でき、自己ホスト、フォーク、変更が可能です。
企業向けのオープンソース AI ゲートウェイ — クォータ、PII 保護、コスト追跡、コンプライアンス
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

Open-source AI gateway for enterprises — quota, PII protection, cost tracking, and compliance - SugaC-275/ToTra

GitHub - SugaC-275/ToTra: Open-source AI gateway for enterprises — quota, PII protection, cost tracking, and compliance · GitHub
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
SugaC-275
/
ToTra
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
255 Commits 255 Commits .github .github admin admin dashboard dashboard docs docs gateway gateway infra infra parser parser scripts scripts sdk sdk website website .env.example .env.example .gitignore .gitignore Makefile Makefile README.md README.md docker-compose.dev.yml docker-compose.dev.yml docker-compose.yml docker-compose.yml View all files Repository files navigation
AI Gateway & Governance Platform
Open-source LLM proxy written in Go. Add quota enforcement, PII blocking, cost tracking, and compliance to any LLM in one line of code.
Quick Start ·
Integration Guide ·
Features ·
Architecture ·
Gateway Docs ·
Admin API ·
Discussions
ToTra is an open-source AI gateway and governance platform that sits in front of any LLM provider.
Point your existing apps at ToTra instead of OpenAI, Anthropic, or any other provider — and instantly get:
Quota enforcement — per-user and per-team hard budget caps
PII blocking — 18 language groups scanned at the edge before any data leaves your network
Cost tracking — per-user, per-team, per-model token and USD spend with chargeback reports
Compliance — GDPR workflows, EU AI Act checklist, hash-chained immutable audit log
Zero code changes — 100% OpenAI-compatible; swap one line in your config
flowchart LR
A["🖥️ Your App\n(OpenAI SDK / curl\n/ LangChain)"] -->|"1 · API request"| B
subgraph B["ToTra Gateway :8080"]
direction TB
B1["🔑 Auth & API Key"]
B2["📊 Quota Check\n(per user / team)"]
B3["🔒 PII Scan\n(18 languages)"]
B4["⚡ Semantic Cache"]
B5["🔀 Route & Load Balance"]
B1 --> B2 --> B3 --> B4 --> B5
end
B -->|"2 · forward request"| C["☁️ LLM Providers\nOpenAI · Anthropic\nGemini · Mistral · Azure\nBedrock · Ollama"]
C -->|"3 · response"| A
B -->|"4 · usage events"| D
subgraph D["ToTra Admin :8081"]
direction TB
D1["💸 Cost Tracking"]
D2["📋 Compliance & Audit"]
D3["🔔 Budget Alerts"]
end
D --> E["📊 Dashboard :3000\nAdmin Console · Reports\nEmployee Self-Service"]
Loading
Why ToTra
🚀 Written in Go — < 2 ms p95 overhead. Native binary, no Python runtime, no warm-up.
🔒 PII blocked at the edge — email, IDs, credit cards, health records across 18 language groups. Sensitive data is redacted before it ever reaches an LLM.
💸 Hard budget caps — requests over limit get 429 before touching any provider. Real-time Slack / webhook alerts.
📋 Compliance out of the box — GDPR data-subject workflows, EU AI Act checklist, and an immutable hash-chained audit log on every request.
📊 Finance-ready reporting — department chargeback CSV, budget forecasts, spend anomaly detection.
🏠 Self-hosted — your keys, your infrastructure, your data. No external dependency.
Prerequisites: Docker + Docker Compose
git clone https://github.com/SugaC-275/ToTra.git
cd ToTra
cp .env.example .env # fill in your provider API keys
docker-compose --profile app up -d --wait
Open http://localhost:3000 and sign in:
Change default credentials immediately after first login via Settings → Security .
One line change. Every other line of code stays the same.
import openai
# Before — calls OpenAI directly
client = openai . OpenAI ( api_key = "sk-..." )
# After — routes through ToTra (zero other changes)
client = openai . OpenAI (
api_key = "your-totra-api-key" , # issued from the ToTra admin panel
base_url = "http://your-totra-host:8080/v1"
)
response = client . chat . completions . create (
model = "gpt-4o" ,
messages = [{ "role" : "user" , "content" : "Hello!" }]
)
print ( response . choices [ 0 ]. message . content )
Node.js / TypeScript (OpenAI SDK)
import OpenAI from "openai" ;
const client = new OpenAI ( {
apiKey : "your-totra-api-key" ,
baseURL : "http://your-totra-host:8080/v1" ,
} ) ;
const response = await client . chat . completions . create ( {
model = "gpt-4o" ,
messages : [ { role : "user" , content : "Hello!" } ] ,
} ) ;
console . log ( response . choices [ 0 ] . message . content ) ;
curl
curl http://your-totra-host:8080/v1/chat/completions \
-H " Authorization: Bearer your-totra-api-key " \
-H " Content-Type: application/json " \
-d ' {
"model": "gpt-4o",
"messages": [{"role": "user", "content": "Hello!"}]
} '
LangChain
from langchain_openai import ChatOpenAI
llm = ChatOpenAI (
model = "gpt-4o" ,
openai_api_key = "your-totra-api-key" ,
openai_api_base = "http://your-totra-host:8080/v1" ,
)
response = llm . invoke ( "Hello!" )
print ( response . content )
Once connected, every request is automatically routed through quota enforcement, PII scanning, semantic caching, and cost tracking.
🔒 PII Protection — 18 Language Groups
Every request body is scanned in real time before it reaches any LLM. Detected PII is redacted and the event is logged. Blocked requests return 422 .
Configure rules per team, per model, or globally in the admin panel.
Per-user, per-team, per-model token and USD cost tracking
Hard budget caps — requests over limit get 429 before touching the provider
Configurable alert thresholds with Slack / Feishu / webhook notifications
Monthly budget forecasts based on current burn rate
Department chargeback reports with CSV export for finance
Procurement analytics and ROI dashboards
Spend anomaly detection with automatic alerts
Dashboard → Cost → Reports → Export CSV
📋 Compliance & Audit
GDPR — data-subject export and deletion request workflows, configurable retention policies
EU AI Act — compliance checklist with per-model status tracking
Immutable audit chain — every request is hash-chained; the log cannot be tampered with
SIEM integration — configurable webhook targets for security event forwarding
Data residency controls — keep all data on-premises or in a specific region
OpenAI-compatible — drop-in replacement for the OpenAI API ( /v1/chat/completions , /v1/embeddings , streaming)
Anthropic-compatible — native Anthropic messages API support
Multi-provider routing — automatic fallback across providers and models
Semantic cache — SimHash LSH deduplication; repeated prompts skip the LLM entirely
Prompt compression — reduce token spend on long context
Streaming proxy — full text/event-stream support
File pipeline — upload PDF / DOCX / PPTX → parse → chat in one API call
Rate limiting, IP allowlist, API-key authentication
JWT authentication + OIDC / SSO integration
Role-based access control (admin / employee)
User and team management with quota request / approval workflow
Model catalogue — enable, disable, and configure providers per team
Bot notifications — Slack, Feishu, custom webhooks
HR sync connector (CSV import)
Agent session tracking — detects and terminates dead-loop agent sessions automatically
Provider
Chat
Embeddings
Streaming
Files
OpenAI (GPT-4o, o1, o3, o4)
✅
✅
✅
✅
Anthropic (Claude 3.5, 4)
✅
—
✅
✅
Google Gemini
✅
✅
✅
—
Mistral AI
✅
✅
✅
—
Meta Llama (via Ollama)
✅
✅
✅
—
Cohere Command
✅
✅
✅
—
Azure OpenAI
✅
✅
✅
✅
AWS Bedrock
✅
✅
✅
—
Local / Ollama
✅
✅
✅
—
Any OpenAI-compatible endpoint
✅
✅
✅
—
Performance
ToTra is written entirely in Go. The gateway adds < 2 ms overhead at p95 under production load.
Measured against a 100 ms mock upstream. Reproduce the benchmark →
Your Apps (OpenAI SDK / curl / LangChain / any HTTP client)
│
▼
ToTra Gateway :8080
auth · quota · PII scan · policy · semantic cache · routing
│
▼
OpenAI · Anthropic · Gemini · Mistral · Local Models
│
│ (usage events)
▼
ToTra Admin :8081
cost · compliance · budgets · audit trail · notifications
│
▼
Dashboard :3000
admin console · department reports · employee self-service
Service
Stack
Port
gateway
Go 1.25 / Fiber
8080
admin
Go 1.25 / Fiber
8081
parser
Python 3.12 / FastAPI
8090
dashboard
React 19 / Vite
3000
postgres
PostgreSQL 16
5432
redis
Redis 7
6379
Screenshots
Cost Dashboard
Department Reports
User Management
Employee Self-Service
Local Development
# 1. Start databases
docker-compose up -d postgres redis
# 2. Run each service in its own terminal
cd gateway && go run .
cd admin && go run .
cd parser && uvicorn main:app --port 8090
cd dashboard && npm install && npm run dev
# 3. Seed dev credentials (first time only)
cd scripts/set-dev-passwords
POSTGRES_HOST=localhost POSTGRES_DB=totra \
POSTGRES_USER=totra POSTGRES_PASSWORD=totra_secret go run .
Default dev credentials: admin@acme.com / totra123
Copy .env.example to .env . Key variables:
See .env.example for the full list including Redis, SMTP, and notification settings.
make test
# Per service
cd gateway && go test ./...
cd admin && go test ./...
cd dashboard && npm run test:run
cd parser && pytest
Contributing
We welcome contributions — bug fixes, new provider integrations, docs improvements, and feature requests.
git clone https://github.com/SugaC-275/ToTra.git
cd ToTra
# Run tests before submitting
make test
Fork the repo and create a branch from main
Make your change and add tests where relevant
For larger features, open a Discussion first to align on direction.
💬 GitHub Discussions — questions, ideas, show & tell
MIT — free to use, self-host, fork, and modify.
Open-source AI gateway for enterprises — quota, PII protection, cost tracking, and compliance
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
