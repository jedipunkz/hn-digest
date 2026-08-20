---
source: "https://corti.ai/models"
hn_url: "https://news.ycombinator.com/item?id=49372258"
title: "Run frontier LLMs on sovereign EU infrastructure"
article_title: "Models"
image: "https://cdn.prod.website-files.com/679910de24e675a93f045f3b/6a22c7ca245ae393869f6e42_pricing.png"
author: "guillego"
captured_at: "2026-08-20T09:23:47Z"
capture_tool: "hn-digest"
hn_id: 49372258
score: 1
comments: 0
posted_at: "2026-08-20T09:19:21Z"
tags:
  - hacker-news
  - translated
---

# Run frontier LLMs on sovereign EU infrastructure

- HN: [49372258](https://news.ycombinator.com/item?id=49372258)
- Source: [corti.ai](https://corti.ai/models)
- Score: 1
- Comments: 0
- Posted: 2026-08-20T09:19:21Z

## Translation

タイトル: EU の主権インフラ上でフロンティア LLM を実行する
記事タイトル: モデル

記事本文:
モデル
SVG としてロゴをコピー ブランド キットにアクセス
API キーの取得 サインイン プラットフォームのビルド
アヒルに襲われたレポート スタートアップ 新しいエンタープライズ開発者が構築を開始
最初の API 呼び出しは 5 分以内に完了
キー、テスト、50 ドルの無料クレジット
配送担当者へのインタビュー
アヒルに襲われたレポート Speech-to-Text API の比較 NEW
Corti 会社についてのバイヤーズ ガイド 採用情報 イベント お問い合わせ ログイン corti モデルのサインアップ EU の主権インフラ上でフロンティア LLM を実行する
GPU に至るまで準拠し、安全で、検証可能です。ヨーロッパのハードウェア上のフロンティア モデル、コーディング エージェント内、またはアプリケーションから直接呼び出されます。データが地域を離れることはありません。
AI エキスパートを予約する
コルティ S1
EU-CPH
NVIDIA B300 (GEFION)
検証可能
コーディング アプリ Web サイト シールド ロック ストリームライン アイコン: https://streamlinehq.comcoding-apps-websites-shield-lock あなたのコードがヨーロッパを離れることはありません
コーディング アプリ Web サイト プログラミング ブラウザ Streamline アイコン: https://streamlinehq.com
コーディング-アプリ-ウェブサイト-プログラミング-ブラウザ
精度を犠牲にすることなく、推論が 20% 高速化 デザイン レイヤー ストリームライン アイコン: https://streamlinehq.com design-layer GPU に至るまで準拠、検証済み、安全
ビジネス製品のパフォーマンス お金の削減 合理化アイコン: https://streamlinehq.com business-products-performance-money-decrease AI コーディングの支出を最大 10 分の 1 に削減 問題 フロンティア モデルとコントロールのどちらかを選択すべきではありません
推論の場所、セキュリティ、および根本的な変更は、ユーザーではなくプロバイダー側にあります。
プロンプトとワークフローは単一の API を中心に強化されます。巻き戻しには建設よりもコストがかかります。
料金と制限はスケジュールに合わせて変更されるため、能力は予算に合わせて割り当てられます。
契約およびアクセス条件は、商業的な決定ではなく、政治的な決定によって変更される可能性があります。
コーディング エージェントまたはアプリで使用する 2 つの方法
パス 01 のエージェント コーディング

ターミナル
Corti CLI を使用して、OpenCode、ForgeCode、Crush、または Pi を Corti モデルに指定します。既存のワークフローを維持します。
# セットアップウィザードを実行する
npx @corti/cli 初期化モデル
# 資格情報をロードして、エージェントを起動します
セット -a;ソース ~/.env;セット+a
オープンコード
OpenAI対応。既存のアプリ コードは、ベース URL の変更と新しいキーで動作します。
クライアント = OpenAI(
Base_url="https://ai.eu.corti.app/v1",
api_key="<あなたの API キー>")
r = client.chat.completions.create(
モデル = "corti-s1"、メッセージ = msgs
)
ラインナップ モデルを選ぶ。
管轄権を保持します。
モデル
こんな方に最適
推論
100万トークンあたりのコスト
corti-s1 推奨
複雑なエージェントコーディングとリポジトリ全体のリファクタリング
はい
$2.00 in · $8.00 out $0.20 キャッシュされた入力
corti-s1-インスタント
高速な対話型コーディングとインライン補完
いいえ
$2.00 in · $8.00 out $0.20 キャッシュされた入力
コルティ-S1-ミニ
低コストでの大規模なレビュー、テスト、リファクタリング
はい
$1.00 入力 · $4.00 出力 $0.10 キャッシュされた入力
corti-s1-ミニインスタント
コスト重視の大規模な完了
いいえ
$1.00 入力 · $4.00 出力 $0.10 キャッシュされた入力
corti-s1-embedding
コードベースの検索、取得、インデックス作成
該当なし
出力料金なしで 0.03 ドル
プロトタイプから本番環境まで、プロバイダーのせいで AI プロジェクトが停滞する可能性がある
チームの誰かが OpenAI または Anthropic に対して機能を構築します。よくデモします。ビジネスケースは明らかです。
データの常駐についてはお答えできません。譲渡は正当化できません。大規模なコストは承認できません。この機能はレビュー中であり、四半期は終了します。
エンドポイントと認証情報を変更します。アプリケーション ロジック、プロンプト、評価セットを保持します。初めてレビューに合格したインフラストラクチャ上で実行します。
検証可能 すべてのレイヤーで検証可能
ほとんどのベンダーは契約で主権を主張し、スタックをレンタルします。 Corti モデルは、オープンソースのインフラストラクチャ層である Kommodity 上で実行されるため、

暗号化、分離、セキュリティを含むスタック全体は公開されており、要求されるだけでなく監査可能です。
推論前の証明 - ポリシーではなくハードウェアによって隔離が強制されます
リクエストパスに米国のクラウドプロバイダーがありません
インフラストラクチャ コードは GitHub でオープンソースになっており、誰でもレビューできます
トレースのリクエスト EU-CPH
10:24:07.004 リクエストを受信しました EU-コペンハーゲン
10:24:07.006 認証検証済み見積もり OK
10:24:07.009 ノード gefion / n-04 にルーティングされました
10:24:07.011 モデルがロードされた corti-s1
10:24:07.788 完了して EU-コペンハーゲンに戻りました
10:24:07.789 プロンプトは何も保持されませんでした
10:24:07.789 トレーニングに使用されたことはありません
プライベート アーキテクチャによるプライベート
プロンプトと完了はメモリ内で処理され、リクエストで終了します。何も保持されず、モデルをトレーニングするものはなく、別のテナントには何も表示されません。
アクセス、データ処理、チームごとの予算を一元管理
ソブリン クラウドまたはオンプレミス - 本番環境の臨床 AI と同じプラットフォーム
ISO 27001、ISO 42001、GDPR、NIS2、DORA、および EU AI Act のデフォルトの姿勢
ガバナンスのデフォルト
即時保持
なし
データのトレーニング
決して
データの常駐性
あなたの地域
導入
クラウド / オンプレミス
使用状況レポート
チームごと
プラットフォーム ウェイトは簡単な部分です
キャパシティ管理
監視と可観測性
フェイルオーバーと復元力
ツール呼び出し
構造化された出力
即時キャッシュ
アクセス制御
監査可能性
モデルの評価
管理されたモデルのアップグレード
予測可能な価格設定
エンタープライズサポート
今日のリーダーは明日のリーダーではない
Corti は、リーダーボード サイクルごとに統合を再構築することなく、市場の変動に応じて利用可能な最良のモデルを評価して運用します。
経済性 席ではなくトークンにお金を払う
シートごとのコーディング アシスタントは、開発者が使用するかどうかに関係なく料金を請求します。ヨーロッパのインフラストラクチャコストで、実行中の消費料金を管理します。
Corti API 全体で 1 つのクレジット残高
同じc

開発、テスト、運用のための再編集
チームごとの制限とリクエストレベルのレポート
調達とコンプライアンスの諸経費の支払いを停止する前に、年間 187,200 ドルの差が生じます。方向性モデルであり、引用ではありません。
「当社の顧客は、デジタルからの独立性、セキュリティ、コンプライアンス、自社データの制御に妥協することなく世界最高の AI にアクセスできることを望んでいます。」 Jesper Carøe 氏、Trifork Digital Health CEO
早期採用および実装パートナーのクイックスタート 最初のリクエストは 5 分以内に完了
キーを作成し、ベース URL を変更し、呼び出しを行います。 OpenAI 互換 - 既存の SDK は書き換えなしで動作します。
コンソールで API キーを生成します。50 ドルの無料クレジットから始めます
OpenAI SDK またはカールを ai.eu.corti.app に指定します
または、Corti CLI をインストールしてターミナル エージェントに接続します
APIキーの取得
最初のリクエストカール
カール https://ai.eu.corti.app/v1/chat/completions \
-H "認可: ベアラー $CORTI_API_KEY" \
-H "コンテンツ タイプ: application/json" \
-d '{
"モデル": "corti-s1" ,
「メッセージ」: [
{"役割": "ユーザー"、"コンテンツ": "この関数をリファクタリング"}
]
}'
で処理されます
EU-コペンハーゲン
よくある質問
これは米国モデルを再販する欧州企業とどう違うのでしょうか?
再販業者が請求先を変更します。推論は引き続き米国の管轄下のインフラストラクチャ上で実行され、転送は引き続き行われます。 Corti がハードウェアを操作するため、お客様のリクエストは、ヨーロッパの法律に基づいて、ヨーロッパの施設内で当社が実行するマシンで処理されます。
出発したい場合はどうなりますか?
API は OpenAI と互換性があるため、Corti モデルから移行することは、移行する場合と同じベース URL の変更となります。当社は、お客様のプロンプト、完了、または調整データを保持しません。導入を安価にする互換性が、撤退を安くするものです。
これを独自のハードウェアで実行できますか?
はい。 Corti モデルはソブリン クラウドとオンプレミスで利用可能

現在オンプレミスの臨床 AI に使用されているのと同じプラットフォーム上で展開できます。オンプレミスでは商用および展開のタイムラインが変更されるため、早めに引き上げてください。
モデルの更新と非推奨はどのように処理しますか?
モデルのバージョンが固定されています。事前に通知し、以前のバージョンと重複する期間を設けて、いつ移動するかを選択します。ベンダーが新しいデフォルトを出荷したため、実行中のワークロードでは何も変更されません。
API を直接使用する場合は、ベース URL とキー。端末コーディング エージェントの場合は、Corti CLI をインストールし、既存のエージェントをそこにポイントします。通常、チームは午後に統合を実行し、数週間でガバナンス モデルの合意が得られます。
最強の無差別級モデルのいくつかは中国で開発されました。その主権者はどうですか？
主権とは、誰がインフラを運営し、誰の法律が適用されるかに関するものであり、研究がどこで行われたかではありません。オープンウェイトは静的なアーティファクトです。 Corti はヨーロッパで運用しているハードウェアにそれらを展開しますが、元の開発者にはデータが届きません。依存関係なしで機能を取得できますが、所有者の API を介してのみアクセスできるクローズド モデルではこれは不可能です。
私たちは医療の分野ではありません。これは私たちのためですか？
医療はコンプライアンスが最も厳しい環境であるため、Corti は医療のためにこれを構築しました。その後に続く要件、データの保存場所、監査可能性、厳格なプライバシー管理は、金融サービス、エネルギー、行政、防衛、航空が直面する要件と同じです。
ヨーロッパのハードウェアで次のコーディング タスクを実行します。
CLI をインストールし、チームがすでに使用しているエージェントを保持し、それを EU でホストされているモデルに向けます。

## Original Extract

Models
Copy Logo as SVG Visit Brand Kit
Get API key Sign in Platform Build
Report Struck by a Duck Startups NEW Enterprise Developers Start building
First API call in under 5 minutes
Keys, testing, $50 free credits
Interviews with people shipping
Report Struck by a Duck Compare Speech-to-Text APIs NEW
Buyer’s Guide to Corti Company About Careers Events Contact Log in Sign up corti models Run frontier LLMs on sovereign EU infrastructure
Compliant, secure, and verifiable down to the GPU. Frontier models on European hardware, in your coding agent or called directly from your application. Your data never leaves your region.
Book an AI expert
Corti S1
EU-CPH
NVIDIA B300 (GEFION)
VERIFIABLE
Coding Apps Websites Shield Lock Streamline Icon: https://streamlinehq.com coding-apps-websites-shield-lock Your code never leaves Europe
Coding Apps Websites Programming Browser Streamline Icon: https://streamlinehq.com
coding-apps-websites-programming-browser
20% faster inference, with no accuracy tradeoff Design Layer Streamline Icon: https://streamlinehq.com design-layer Compliant, verified, and secure, down to the GPU
Business Products Performance Money Decrease Streamline Icon: https://streamlinehq.com business-products-performance-money-decrease Cut AI coding spend up to 10x The problem You shouldn't choose between frontier models and control
Inference location, security, and underlying changes sit with the provider - not with you.
Prompts and workflows harden around a single API. Unwinding costs more than building did.
Rates and limits move on their schedule, so capability gets rationed to budget.
Contract and access terms can shift with a political decision, not a commercial one.
Two ways to use it In the coding agent, or in your app
Path 01 Agentic coding in the terminal
Point OpenCode, ForgeCode, Crush, or Pi at Corti Models with the Corti CLI. Keep the workflow you already have.
# run the setup wizard
npx @corti/cli init models
# load credentials, then launch your agent
set -a; source ~/.env; set +a
opencode
OpenAI-compatible. Existing app code works with a base URL change and a new key.
client = OpenAI(
base_url="https://ai.eu.corti.app/v1",
api_key="<your-api-key>")
r = client.chat.completions.create(
model="corti-s1", messages=msgs
)
The lineup Pick the model.
Keep the jurisdiction.
Model
Best for
Reasoning
Cost per 1M tokens
corti-s1 Recommended
Complex agentic coding and repo-wide refactors
Yes
$2.00 in · $8.00 out $0.20 cached input
corti-s1-instant
Fast interactive coding and inline completion
No
$2.00 in · $8.00 out $0.20 cached input
corti-s1-mini
High-volume review, tests, and refactors at lower cost
Yes
$1.00 in · $4.00 out $0.10 cached input
corti-s1-mini-instant
Cost-sensitive completion at scale
No
$1.00 in · $4.00 out $0.10 cached input
corti-s1-embedding
Codebase search, retrieval, and indexing
n/a
$0.03 in no output charge
From prototype to production AI projects can stall because of the provider
Someone on your team builds the feature against OpenAI or Anthropic. It demos well. The business case is obvious.
Data residency cannot be answered. The transfer cannot be justified. The cost at scale cannot be approved. The feature sits in review and the quarter ends.
Change the endpoint and the credentials. Keep your application logic, your prompts, and your evaluation set. Run it on infrastructure that passes review the first time.
Verifiable Verifiable at every layer
Most vendors claim sovereignty in the contract, then rent the stack. Corti Models runs on Kommodity, an open source infrastructure layer, so the entire stack, including encryption, isolation, and security, is public and auditable, not just claimed.
Attestation before inference - isolation enforced by hardware, not policy
No US cloud provider in the request path
Infrastructure code open source on GitHub for anyone to review
Request trace EU-CPH
10:24:07.004 request received eu-copenhagen
10:24:07.006 attestation verified quote ok
10:24:07.009 routed to node gefion / n-04
10:24:07.011 model loaded corti-s1
10:24:07.788 completion returned eu-copenhagen
10:24:07.789 prompt retained none
10:24:07.789 used for training never
Private Private by architecture
Prompts and completions process in memory and end with the request. Nothing is retained, nothing trains a model, nothing is visible to another tenant.
Central control over access, data handling, and per-team budget
Sovereign cloud or on-premises - same platform as clinical AI in production
ISO 27001, ISO 42001, GDPR, NIS2, DORA, and EU AI Act posture by default
Governance Default
Prompt retention
none
Training on your data
never
Data residency
your region
Deployment
cloud / on-prem
Usage reporting
per team
The platform The weights are the easy part
Capacity management
Monitoring and observability
Failover and resilience
Tool calling
Structured outputs
Prompt caching
Access controls
Auditability
Model evaluation
Managed model upgrades
Predictable pricing
Enterprise support
Today’s leader isn't tomorrow’s
Corti evaluates and operates the best available models as the market moves - without you rebuilding integrations every leaderboard cycle.
Economics Pay for tokens, not seats
Per-seat coding assistants charge whether developers use them or not. Governed consumption charges for what runs, at European infrastructure cost.
One credit balance across Corti APIs
Same credits for development, testing, and production
Per-team limits and request-level reporting
Difference of $187,200 a year, before the procurement and compliance overhead you stop paying for. Directional model, not a quote.
“Our customers want digital independence, access to the world's best AI without compromising on security, compliance, or control over their own data.” Jesper Carøe , CEO, Trifork Digital Health
Early adopter and implementation partner Quickstart First request in under five minutes
Create a key, change the base URL, make the call. OpenAI-compatible - any existing SDK works without a rewrite.
Generate an API key in the console, start with $50 of free credits
Point any OpenAI SDK or curl at ai.eu.corti.app
Or install the Corti CLI and connect your terminal agent
Get API Key
First request curl
curl https://ai.eu.corti.app/v1/chat/completions \
-H "Authorization: Bearer $CORTI_API_KEY" \
-H "Content-Type: application/json" \
-d '{
"model": "corti-s1" ,
"messages": [
{"role": "user", "content": "Refactor this function"}
]
}'
Processed in
eu-copenhagen
Frequently asked questions
How is this different from a European company reselling US models?
A reseller changes who invoices you. The inference still runs on infrastructure under US jurisdiction and the transfer still happens. Corti operates the hardware, so your request is processed on machines we run, in a European facility, under European law.
What happens if we want to leave?
The API is OpenAI-compatible, so moving off Corti Models is the same base URL change as moving on. We do not hold your prompts, completions, or tuning data. The compatibility that makes adoption cheap is what makes exit cheap.
Can we run this on our own hardware?
Yes. Corti Models is available in sovereign cloud and on-premises deployments, on the same platform used for on-premises clinical AI today. On-premises changes the commercial and deployment timeline, so raise it early.
How do you handle model updates and deprecation?
Model versions are pinned. You choose when to move, with advance notice and an overlap window on the previous version. Nothing changes underneath a running workload because a vendor shipped a new default.
For direct API use, a base URL and a key. For terminal coding agents, installing the Corti CLI and pointing an existing agent at it. Teams typically have a working integration in an afternoon and a governance model agreed in a few weeks.
Some of the strongest open-weight models were developed in China. How is that sovereign?
Sovereignty is about who operates the infrastructure and whose law applies, not where the research was done. Open weights are a static artefact. Corti deploys them on hardware we operate in Europe, and no data reaches the original developer. You get the capability without the dependency, which is not possible with a closed model you can only reach through its owner's API.
We are not in healthcare. Is this for us?
Corti built this for healthcare because healthcare is the hardest compliance environment there is. The requirements that follow, data residency, auditability, and strict privacy controls, are the same ones facing financial services, energy, public administration, defence, and aviation.
Run your next coding task on European hardware
Install the CLI, keep the agent your team already uses, and point it at models hosted in the EU.
