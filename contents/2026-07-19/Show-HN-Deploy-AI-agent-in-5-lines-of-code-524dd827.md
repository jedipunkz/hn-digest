---
source: "https://www.custodianlabs.io"
hn_url: "https://news.ycombinator.com/item?id=48972442"
title: "Show HN: Deploy AI agent in 5 lines of code"
article_title: "Custodian Labs | Build and deploy AI agents in 5 lines of code"
author: "sherryf123"
captured_at: "2026-07-19T23:48:01Z"
capture_tool: "hn-digest"
hn_id: 48972442
score: 1
comments: 1
posted_at: "2026-07-19T22:57:48Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Deploy AI agent in 5 lines of code

- HN: [48972442](https://news.ycombinator.com/item?id=48972442)
- Source: [www.custodianlabs.io](https://www.custodianlabs.io)
- Score: 1
- Comments: 1
- Posted: 2026-07-19T22:57:48Z

## Translation

タイトル: HN の表示: 5 行のコードで AI エージェントをデプロイする
記事のタイトル: カストディアン ラボ | 5 行のコードで AI エージェントを構築してデプロイする
説明: Custodian Labs は、開発者がインフラストラクチャやベンダー ロックインがなく、プライバシーが安全な設計で AI エージェントを出荷するための最速の方法です。

記事本文:
カストディアンラボ | 5 行のコードで AI エージェントを構築してデプロイする
iframe がブロックされています。 Cookie を受け入れてロードします。設定 すべてを受け入れる 同意設定の管理
Essentials 常にアクティブ サイトが機能するために必要です。常にオン。
プライバシー ポリシー名 トラッカー名 説明 トラッカーの説明
このカテゴリのトラッカーは検出されませんでした。
このカテゴリのプロバイダーは検出されませんでした。
分析 分析 使用状況を測定し、エクスペリエンスを向上させます。
プライバシー ポリシー名 トラッカー名 説明 トラッカーの説明
このカテゴリのトラッカーは検出されませんでした。
このカテゴリのプロバイダーは検出されませんでした。
マーケティング マーケティング ターゲットを絞った広告に使用されます。
プライバシー ポリシー名 トラッカー名 説明 トラッカーの説明
このカテゴリのトラッカーは検出されませんでした。
このカテゴリのプロバイダーは検出されませんでした。
パーソナライゼーション パーソナライゼーション ユーザーの設定を記憶し、拡張機能を提供します。
プライバシー ポリシー名 トラッカー名 説明 トラッカーの説明
このカテゴリのトラッカーは検出されませんでした。
このカテゴリのプロバイダーは検出されませんでした。
「同意する」をクリックすると、サイトのナビゲーションを強化し、サイトの使用状況を分析し、当社のマーケティング活動を支援するために、デバイスに Cookie が保存されることに同意したことになります。詳細については、当社のプライバシー ポリシーをご覧ください。
プレシード資金提供 · AUT Ventures & ニュージーランド政府
サインイン API キーを取得する はじめに 機能 カストディアンを選ぶ理由 マルチエージェント ガーディアン レイヤーの価格 AI エージェントを 5 行のコードで構築およびデプロイします。
Custodian Labs は、あらゆる層にプライバシーが組み込まれた状態で、アイデアから運用グレードの AI エージェントを導入するまでをお手伝いします。インフラストラクチャがありません。データ漏洩はありません。
custodian_labs からカストディアンをインポート
モデル = 管理者(
モデル = "gpt-4o",
system_prompt="あなたは役に立つアシスタントです...",
）
model.deploy() カストディアンを選ぶ理由 必要なものすべて。あなたがやらないことは何もありません。
独自のガード

ian Layer は、PII を検出し、モデルに到達する前にその処理方法を制御できるガードレールです。
埋め込みやベクター DB を構成せずに、エージェントに長期記憶とドキュメント検索を提供します。ナレッジベースを 1 行追加するだけで機能します。
OpenAI、Anthropic、Mistral、またはローカル モデルを 1 行で切り替えることができます。ベンダー ロックインや書き換えは必要ありません。エージェントのロジックはそのまま残ります。その下のモデルのみが変更されます。
プロビジョニングするデータベース、接続するベクター ストア、管理するホスティングは必要ありません。カストディアンはエージェントの下にあるすべてのレイヤーを抽象化するため、ユーザーはロジックに完全に集中できます。
AUT Ventures とニュージーランド政府の資金提供により商品化されました。深い学術研究に基づいた製造の信頼性。
データベースはありません。ホスティングはありません。デプロイするだけです。
通常、エージェント ロジックを 1 行記述する前に、ベクター DB のプロビジョニング、ホスティング環境のセットアップ、モデル プロバイダーの接続、再試行ロジックの作成、ストリーミングの処理をすべて行います。カストディアンは、そのスタック全体を関数呼び出しに折りたたみます。
# ステップ 1: ベクター DB + 埋め込みパイプラインをプロビジョニングする
pinecone.init(api_key=...)、インデックス = pinecone.Index("my-index")
embedder = SentenceTransformer("all-MiniLM-L6-v2")
チャンク = [text[i:i+512] for i in range(0, len(text), 448)]
Index.upsert([(f"doc_{i}", embedder.encode(c)) for i, c in enumerate(chunks)])
# ステップ 2: 手動 PII スクラビング (脆弱な正規表現、最善を尽くします)
re.sub(r" [A-Z][a-z]+ [A-Z][a-z]+ ", "[編集済み]", プロンプト)
# ステップ 3: ツールのディスパッチ + 再試行ロジック
def call_with_retry(messages, max_retries=3):
範囲(max_retries)内の試行:
試してください: return client.chat.completions.create(...)
例外: time.sleep(2 ** 試行)
# ステップ 4: Redis 経由のセッション メモリ
redis_client.setex(f"ctx:{session_id}", 3600, json.dumps(messages[-20:]))
# ステップ 5: Flask サーバー、認証、

レート制限、CORS、ロギング...
# ステップ 6: Dockerfile、k8s config、CI/CD パイプラインを作成します...
# ステップ 7: 最後に、エージェントが実行されます。多分。 custodian_labs からカストディアンをインポート
モデル = 管理者(
モデル = "gpt-4o",
system_prompt="あなたは役に立つアシスタントです...",
）
model.deploy() マルチエージェント機能 チームを編成します。一つとして展開します。
custodian_labs からエージェント、エージェントチームをインポート
チーム = エージェントチーム(
エージェント=[
エージェント(
名前 = "エンジニア",
モデル = "gpt-4o",
system_prompt="あなたは上級エンジニアです。技術およびエンジニアリングに関するすべての質問に対応してください。",
トピックス=["テクノロジー"、"エンジニアリング"、"AI"],
）、
エージェント(
name="戦略家",
モデル = "gpt-4o",
system_prompt="あなたは戦略リーダーです。インテリジェンスと戦略に関する質問を処理します。",
トピックス=["諜報"、"戦略"、"スパイ活動"、"戦術"],
）、
エージェント(
名前 = "科学者",
モデル = "gpt-4o",
system_prompt="あなたは研究科学者です。科学や研究に関する質問に対処してください。",
トピックス=["科学", "研究", "生物学", "物理学"],
）、
]、
ルーティングモード = "シングル"、
）
app = Team.deploy() 独自のテクノロジー Guardian Layer の紹介
ほとんどのプライバシー ソリューションは単純に PII を削除しますが、コンテキストを削除すると AI の精度が損なわれます。 Guardian Layer は PII をインテリジェントに検出し、元の PII が環境から離れることなく、次に何が起こるかを完全に制御できます。
Guardian Layer — AI モデルに送信される変換 240,000 ドルの契約更新について、Alex Morgan (alex@synthco.io) とのフォローアップをスケジュールします。木曜日の午後 2 時に James Harrington によるフォローアップが予定されている PII が再リンクされて返送されました。カレンダーの招待状が james@acmecorp.com に送信されました。入力 (生) $240,000 の契約更新について、james@acmecorp.com で James Harrington とのフォローアップをスケジュールします。ガーディアン レイヤー — マスク/編集 AI モデルに送信 [P] でフォローアップのスケジュールを設定

[CURRENCY] の契約更新について、[ERSON] まで [EMAIL] でご連絡ください。 PII が再リンクされて返送されました。木曜日の午後 2 時に [PERSON] とのフォローアップが予定されています。カレンダーの招待状が [EMAIL] に送信されました。入力（生）
ガーディアン レイヤー — AI モデルに送信のみを検出
個人識別情報 (PII) を、意味を保持する合成同等の情報に置き換えます。
個人を特定できる情報 (PII) を型指定されたラベルに置き換えます - 最速かつ最も保守的な変換
テキストを変更せずに、個人識別情報 (PII) に置き換わるもののみを表面化します。
個人識別情報 (PII) を、意味を保持する合成同等の情報に置き換えます。
Guardian Layer — AI モデルに送信される変換 240,000 ドルの契約更新について、Alex Morgan (alex@synthco.io) とのフォローアップをスケジュールします。木曜日の午後 2 時に James Harrington によるフォローアップが予定されている PII が再リンクされて返送されました。カレンダーの招待状が james@acmecorp.com に送信されました。マスク/編集
個人を特定できる情報 (PII) を型指定されたラベルに置き換えます - 最速かつ最も保守的な変換
ガーディアン レイヤー — マスク/秘匿化 AI モデルに送信 [CURRENCY] の契約更新について、[EMAIL] の [PERSON] とのフォローアップをスケジュールします。 PII が再リンクされて返送されました。木曜日の午後 2 時に [PERSON] とのフォローアップが予定されています。カレンダーの招待状が [EMAIL] に送信されました。検出のみ
テキストを変更せずに、個人識別情報 (PII) に置き換わるもののみを表面化します。
ガーディアン レイヤー — AI モデルに送信のみを検出
より多くの制限、優先アクセス、VIP サポートを必要とするユーザー向け
スタックとコンプライアンスのニーズに合わせてすべてをカスタマイズ
カスタムコンプライアンスとデータ契約
アイデアから導入まで。
10 分以内に完了します。
API キーを取得し、SDK をインストールし、プライバシーに配慮した AI エージェントを実稼働環境で実行します。

— インフラストラクチャは必要ありません。
© 2026 Custodian Labs Limited 法的事項

## Original Extract

Custodian Labs is the fastest way for developers to ship AI agents with no infrastructure, no vendor lock-in, and privacy-safe by design.

Custodian Labs | Build and deploy AI agents in 5 lines of code
Iframe is blocked. Accept cookies to load it. Preferences Accept All Manage Consent Preferences
Essentials Always active Necessary for the site to function. Always On.
Privacy Policy Name Tracker Name Description Tracker description
No trackers detected for this category.
No providers detected for this category.
Analytics Analytics Measures usage and improves your experience.
Privacy Policy Name Tracker Name Description Tracker description
No trackers detected for this category.
No providers detected for this category.
Marketing Marketing Used for targeted advertising.
Privacy Policy Name Tracker Name Description Tracker description
No trackers detected for this category.
No providers detected for this category.
Personalization Personalization Remembers your preferences and provides enhanced features.
Privacy Policy Name Tracker Name Description Tracker description
No trackers detected for this category.
No providers detected for this category.
By clicking "Accept" , you agree to the storing of cookies on your device to enhance site navigation, analyze site usage, and assist in our marketing efforts. View our Privacy Policy for more information.
Pre-seed funded · AUT Ventures & New Zealand Government
Sign in Get API Key Introduction Features Why Custodian Multi-agent Guardian layer Pricing Build and deploy AI agents in 5 lines of code.
Custodian Labs takes you from idea to deployed, production-grade AI agent — with privacy built in at every layer. No infrastructure. No data leakage.
from custodian_labs import Custodian
model = Custodian(
model="gpt-4o",
system_prompt="You are a helpful assistant...",
)
model.deploy() Why Custodian Everything you need. Nothing you don't.
The proprietary Guardian Layer is the guardrail that detects PII and puts you in control of how it's handled before it reaches any model.
Give your agent long-term memory and document retrieval without configuring embeddings or vector DBs. Add a knowledge base in one line and it just works.
Switch between OpenAI, Anthropic, Mistral, or local models with one line — no vendor lock-in, no rewrite required. Your agent logic stays exactly as it is; only the model underneath changes.
No database to provision, no vector store to connect, no hosting to manage. Custodian abstracts every layer beneath your agent so you focus entirely on logic.
Commercialised with AUT Ventures and New Zealand Government funding. Production reliability grounded in deep academic research.
No databse. No hosting. Just deploy.
Normally you'd provision a vector DB, set up a hosting environment, wire up your model provider, write retry logic, handle streaming — all before writing a single line of agent logic. Custodian collapses that entire stack to a function call.
# Step 1: Provision a vector DB + embedding pipeline
pinecone.init(api_key=...), index = pinecone.Index("my-index")
embedder = SentenceTransformer("all-MiniLM-L6-v2")
chunks = [text[i:i+512] for i in range(0, len(text), 448)]
index.upsert([(f"doc_{i}", embedder.encode(c)) for i, c in enumerate(chunks)])
# Step 2: Manual PII scrubbing (fragile regex, hope for the best)
re.sub(r"[A-Z][a-z]+ [A-Z][a-z]+", "[REDACTED]", prompt)
# Step 3: Tool dispatch + retry logic
def call_with_retry(messages, max_retries=3):
for attempt in range(max_retries):
try: return client.chat.completions.create(...)
except: time.sleep(2 ** attempt)
# Step 4: Session memory via Redis
redis_client.setex(f"ctx:{session_id}", 3600, json.dumps(messages[-20:]))
# Step 5: Flask server, auth, rate limiting, CORS, logging...
# Step 6: Write your Dockerfile, k8s config, CI/CD pipeline...
# Step 7: Finally — your agent runs. Maybe. from custodian_labs import Custodian
model = Custodian(
model="gpt-4o",
system_prompt="You are a helpful assistant...",
)
model.deploy() Multi-agent capabilities Assemble your team. Deploy as one.
from custodian_labs import Agent, AgentTeam
team = AgentTeam(
agents=[
Agent(
name="the_engineer",
model="gpt-4o",
system_prompt="You are a senior engineer. Handle all tech and engineering questions.",
topics=["technology", "engineering", "AI"],
),
Agent(
name="the_strategist",
model="gpt-4o",
system_prompt="You are a strategy lead. Handle intelligence and strategy questions.",
topics=["intelligence", "strategy", "espionage", "tactics"],
),
Agent(
name="the_scientist",
model="gpt-4o",
system_prompt="You are a research scientist. Handle science and research questions.",
topics=["science", "research", "biology", "physics"],
),
],
routing_mode="single",
)
app = team.deploy() Proprietary technology Introducing the Guardian Layer
Most privacy solutions simply strip PII — but stripping context destroys AI accuracy. The Guardian Layer detects PII intelligently, then gives you full control over what happens next, without the original PII ever leaving your envionrment.
Guardian Layer — Transform Sent to AI model Schedule a follow-up with Alex Morgan at alex@synthco.io about the contract renewal for $240,000 . Returned to you with PII relinked Follow-up scheduled with James Harrington for Thursday at 2pm. Calendar invite sent to james@acmecorp.com . Input (raw) Schedule a follow-up with James Harrington at james@acmecorp.com about the contract renewal for $240,000 . Guardian Layer — Mask / redact Sent to AI model Schedule a follow-up with [PERSON] at [EMAIL] about the contract renewal for [CURRENCY] . Returned to you with PII relinked Follow-up scheduled with [PERSON] for Thursday at 2pm. Calendar invite sent to [EMAIL] . Input (raw)
Guardian Layer — Detect only Sent to AI model
Replaces personally identifiable information (PII) with synthetic equivalents that preserve meaning
Replaces personally identifiable information (PII) with a typed label — the fastest and most conservative transform
Only surface what Replaces personally identifiable information (PII) exists, without modifying any text
Replaces personally identifiable information (PII) with synthetic equivalents that preserve meaning
Guardian Layer — Transform Sent to AI model Schedule a follow-up with Alex Morgan at alex@synthco.io about the contract renewal for $240,000 . Returned to you with PII relinked Follow-up scheduled with James Harrington for Thursday at 2pm. Calendar invite sent to james@acmecorp.com . Mask / redact
Replaces personally identifiable information (PII) with a typed label — the fastest and most conservative transform
Guardian Layer — Mask / redact Sent to AI model Schedule a follow-up with [PERSON] at [EMAIL] about the contract renewal for [CURRENCY] . Returned to you with PII relinked Follow-up scheduled with [PERSON] for Thursday at 2pm. Calendar invite sent to [EMAIL] . Detect only
Only surface what Replaces personally identifiable information (PII) exists, without modifying any text
Guardian Layer — Detect only Sent to AI model
For users who want more limits, priority access, and VIP support
Everything tailored to your stack and compliance needs
Custom compliance & data contracts
From idea to deployed.
In under 10 minutes.
Get an API key, install the SDK, and have a privacy-safe AI agent running in production — no infrastructure required.
© 2026 Custodian Labs Limited Legal
