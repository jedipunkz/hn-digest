---
source: "https://www.vlm.run/product/gateway"
hn_url: "https://news.ycombinator.com/item?id=49365625"
title: "Run GLM-OCR, DeepSeek-OCR-2, Dots.mocr with an OpenAI Compatible API"
article_title: "Gateway | VLM Run"
image: "https://vlm.run/og?title=VLM+Run"
author: "fzysingularity"
captured_at: "2026-08-19T19:19:02Z"
capture_tool: "hn-digest"
hn_id: 49365625
score: 6
comments: 0
posted_at: "2026-08-19T18:49:17Z"
tags:
  - hacker-news
  - translated
---

# Run GLM-OCR, DeepSeek-OCR-2, Dots.mocr with an OpenAI Compatible API

- HN: [49365625](https://news.ycombinator.com/item?id=49365625)
- Source: [www.vlm.run](https://www.vlm.run/product/gateway)
- Score: 6
- Comments: 0
- Posted: 2026-08-19T18:49:17Z

## Translation

タイトル: OpenAI 互換 API を使用して GLM-OCR、DeepSeek-OCR-2、Dots.mocr を実行する
記事のタイトル: ゲートウェイ | VLM の実行
説明: ビジュアル AI 用の OpenAI 互換ゲートウェイ。 1 つのチャット完了エンドポイントの背後にある 10 個のビジュアル モデル。トークンとイメージごとの価格設定、構造化された出力、BYOK 対応。

記事本文:
ゲートウェイ | VLM Run 製品ソリューション 価格ドキュメント ブログ 概要 デモの予約 ダッシュボード 製品プラットフォーム ゲートウェイ すべてのビジュアル モデルに 1 つの OpenAI 互換 API。 Orion チャット プレイグラウンド ホスト型チャット インターフェイスで Orion をお試しください。ホワイトペーパー 当社の主力ビジュアル エージェントである Orion の技術概要。オープンソース mm エージェント向けの高速なマルチモーダル コンテキスト。 vlmbench 1 つのコマンドでハードウェア上の任意の VLM をベンチマークします。 vlmrun-hub 実稼働 VLM ユースケース用の構造化スキーマ。モダリティ ドキュメント別のソリューション 数時間にわたる手作業によるドキュメント作業を数秒間のスキーマ検証済み JSON に変換します。画像 近日公開予定 キャプション、検出、セグメント化、生成。すべての画像操作は 1 つの API を通じて行われます。ビデオ 近日公開予定 長いビデオを視聴せずに要約、文字起こし、検索します。業界別 物理 AI エージェントによるデータのデモンストレーション、再構築、および根拠のあるトレーニング データのラベル付け。ヘルスケア 乱雑な FAX、スキャン、臨床書類をクリーンで検証済みのデータに変換します。建設 設計図、スケジュール、仕様書を構造化データに変換します。価格設定 ドキュメント ブログ 概要 デモの予約 ダッシュボード 1 つの統合 API、任意のビジュアル モデル。
ドキュメント OCR、キャプション、マルチモーダル チャットなど、あらゆるビジュアル機能が 1 つの MCP サーバーで実現されます。
pydantic_ai import Agent から API キー Pydantic AI Mastra LangChain を取得します
pydantic_ai.capabilities から MCP をインポート
mcp = MCP (
"https://gateway.vlm.run/mcp" ,
headers= { "認可" : "ベアラー <VLMRUN_API_KEY>" } ,
)
エージェント = エージェント ( "anthropic:claude-sonnet-5" 、機能 = [ mcp ] )
print (agent.run_sync ( "invoice.pdf から請求書フィールドを抽出します" ) .output )
サポートされているプロバイダー
7 つのビジュアルモデル。エンドポイントは 1 つ。
すべてのチャット キャプション OCR マークダウン検出 セグメンテーション ポーズ検索モデル すべてのプロバイダー モデル機能 コンテキスト入力 / 1M 出力 / 1M dots.mocr
ドキュメントを選択してください

エントタイプ、ページボリューム、OCR モデル。コスト削減とクローズドビジョン API を比較してください。
OCR モデル 全モデル (価格帯) VLM Run Gateway (dots.mocr) VLM Run Gateway (GLM-OCR) VLM Run Gateway (DeepSeek-OCR-2) VLM Run Gateway (PP-OCRv6) 100K ページ量 / 月
フロンティア請求トークン (イン/アウト)
ページあたり 2.5K の画像トークン。出力には 2x OCR テキストでの推論が含まれます。
フロンティアモデルによる総コスト削減
一般的な OCR、Document AI、フロンティア VLM の料金と比較した見積もり。出典: llm-prices.com
ビルダー向けに構築された、ビジョン専用のゲートウェイ。
LLM ルーターとゲートウェイは数百の LLM にルーティングしますが、VLM はほんの少数であり、多くの場合、OCR モデルや従来の CV モデルはありません。 Visual AI には独自のスタックが必要です。
マルチモーダル入力とマルチタスク出力にわたる 1 つのカタログ: OCR、検出、セグメンテーション、ポーズ、キーポイントなど。
OCR、キャプション、マルチモーダル チャットは、すでに使用しているものと同じ OpenAI 互換のチャット完了 API を通じて実行されます。
1 回の通話で 500 ページの PDF または 2 時間のビデオを送信します。ゲートウェイはチャンク、バッチ、および再組み立てを行います。構築するパイプラインはありません。
04 すぐに使える構造化出力
すべての呼び出しで JSON スキーマを適用します。 CV ラッパーは固定スキーマを生成します。 VLM は、response_format を尊重します。
すべてのモデルはオープンウェイトであり、すでに使用している OpenAI 互換 API を通じて提供されます。モデルやプロバイダーを自由に交換しても、クライアント コードは変更されません。
06 近日公開 エージェントネイティブ ビジョン (MCP 経由)
あらゆる MCP クライアントに完全なビジュアル モデル カタログへの即時アクセスを提供します。エージェントは、すぐに画像を見て、読み取り、判断します。
本番環境のビジュアル AI 用に構築されています。
実稼働ワークロードとコスト効率に合わせて最適化されています。導入されたすべてのモデルは、独自のパフォーマンス調整を受けます。
1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 from openai import OpenAI
client = OpenAI (base_url= "https://gateway.vlm.run/v1/openai/" ,
api_key= "<VLM_RUN_A

PI_KEY>" )
resp = client.chat.completions.create (
モデル= "qwen/qwen3.5-0.8b" ,
メッセージ = [ { "役割" : "ユーザー" , "コンテンツ" : [
{ "type" : "text" , "text" : "この画像で何が起こっていますか?" } 、
{ "type" : "image_url" , "image_url" : { "url" : "city.jpg" } } ,
] } ]、
)
print ( resp.choices [ 0 ] .message.content )
#「夕暮れの繁華街…。」
ビジョンネイティブのチャット補完。
OpenAI SDK をゲートウェイでポイントし、モデルを交換します。同じシグネチャ、網羅的なビジュアル モデル カタログ。
ゲートウェイを使用して構築できるもの。
ビルディングブロックのようにビジョンモデルを構成します。 1 つの SDK。鍵は1つ。請求書は 1 枚です。
契約書、明細書、フォーム全体のレイアウト、OCR、値下げ抽出をページあたり 1 セント未満の経済性で実現します。
マルチモーダルなチャットボットとエージェント
画像、PDF、ビデオなどのマルチモーダル入力のネイティブ サポートを 1 回のチャット完了に詰め込んだ VLM を使用して、強力なマルチモーダル チャットボットとエージェントを構築します。
検索、重複排除、推奨のために、大規模な画像とビデオのライブラリに自動キャプション、タグ付け、インデックスを付けます。
1 つの URL の背後にある検出とセグメンテーション。 CV インフラを立ち上げずにビジュアル機能を提供します。
API キーを取得します。今日の船のビジョン。
エンタープライズ向け Visual Intelligence のドキュメントを読んでください。実稼働グレードの精度、可観測性、制御性を備えた画像、ドキュメント、ビデオにわたるビジュアル AI を構築、実行、運用します。
製品の更新と調査。ワンクリックで購読を解除できます。
© 2026 Autonomi AI Inc. 全著作権所有。

## Original Extract

The OpenAI-compatible gateway for visual AI. 10 visual models behind one chat-completions endpoint. Token and per-image pricing, structured outputs, BYOK ready.

Gateway | VLM Run Product Solutions Pricing Docs Blog About Book a Demo Dashboard Product Platform Gateway One OpenAI-compatible API for every visual model. Orion Chat Playground Try Orion in our hosted chat interface. Whitepaper The technical overview of Orion, our flagship visual agent. Open Source mm Fast, multi-modal context for agents. vlmbench Benchmark any VLM on your hardware, in one command. vlmrun-hub Structured schemas for production VLM use cases. Solutions By Modality Documents Turn hours of manual document work into seconds of schema-validated JSON. Images Coming soon Caption, detect, segment, and generate. Every image operation through one API. Videos Coming soon Summarize, transcribe, and search long-form video without watching it. By Industry Physical AI Agentic data labeling for demonstrations, reconstruction, and grounded training data. Healthcare Turn messy faxes, scans, and clinical paperwork into clean, validated data. Construction Turn blueprints, schedules, and specs into structured data. Pricing Docs Blog About Book a Demo Dashboard One unified API, Any Visual Model.
Document OCR, captioning, and multi-modal chat: every visual capability behind one MCP server.
Get an API Key Pydantic AI Mastra LangChain from pydantic_ai import Agent
from pydantic_ai.capabilities import MCP
mcp = MCP (
"https://gateway.vlm.run/mcp" ,
headers= { "Authorization" : "Bearer <VLMRUN_API_KEY>" } ,
)
agent = Agent ( "anthropic:claude-sonnet-5" , capabilities= [ mcp ] )
print ( agent.run_sync ( "Extract the invoice fields from invoice.pdf" ) .output )
Supported providers
7 visual models. One endpoint.
All Chat Caption OCR Markdown Detection Segmentation Pose Search models All Providers Model Capabilities Context Input / 1M Output / 1M dots.mocr
Pick a document type, page volume, and OCR model. See cost savings versus closed vision APIs.
OCR model All models (price band) VLM Run Gateway (dots.mocr) VLM Run Gateway (GLM-OCR) VLM Run Gateway (DeepSeek-OCR-2) VLM Run Gateway (PP-OCRv6) 100K Page volume / month
Frontier billed tokens (in / out)
2.5K image tokens per page; output includes reasoning at 2× OCR text.
Total cost savings from frontier models
Estimates vs typical OCR, Document AI, and frontier VLM pricing. Source: llm-prices.com
A vision-only gateway, built for builders.
LLM routers and gateways route to 100s of LLMs, yet only a handful of VLMs, and often no OCR or classical CV models. Visual AI deserves its own stack.
One catalog spanning multi-modal inputs and multi-task outputs: OCR, detection, segmentation, pose, keypoints, and more.
OCR, captioning, and multimodal chat run through the same OpenAI-compatible chat completions API you already use.
Send a 500-page PDF or a 2-hour video in a single call. The gateway chunks, batches, and reassembles for you. No pipelines to build.
04 Structured Outputs, Out of the Box
JSON-schema enforcement on every call. CV wrappers emit fixed schemas; VLMs honor response_format.
Every model is open-weight, served through the OpenAI-compatible API you already use. Swap models or providers freely, your client code never changes.
06 Coming soon Agent-native vision, via MCP
Give any MCP client instant access to the full visual model catalog. Agents see, read, and reason over images out of the box.
Built for production visual AI.
Optimized for production workloads and cost-efficiency. Every model deployed gets its own performance tune-up.
1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 from openai import OpenAI
client = OpenAI ( base_url= "https://gateway.vlm.run/v1/openai/" ,
api_key= "<VLM_RUN_API_KEY>" )
resp = client.chat.completions.create (
model= "qwen/qwen3.5-0.8b" ,
messages= [ { "role" : "user" , "content" : [
{ "type" : "text" , "text" : "What is happening in this image?" } ,
{ "type" : "image_url" , "image_url" : { "url" : "city.jpg" } } ,
] } ] ,
)
print ( resp.choices [ 0 ] .message.content )
# "A busy city street at dusk ...."
Vision-native Chat Completions.
Point the OpenAI SDK at the gateway and swap the model. Same signature, exhaustive visual model catalog.
What you can build with the Gateway.
Compose vision models like building blocks. One SDK. One key. One bill.
Layout, OCR, and markdown extraction across contracts, statements, and forms, at sub-cent per-page economics.
Multi-modal chatbots and agents
Build powerful multi-modal chatbots and agents with VLMs that pack native support for multi-modal inputs such as images, PDFs, or videos in a single chat completion.
Auto-caption, tag, and index large image and video libraries for search, dedup, and recommendations.
Detection and segmentation behind one URL. Ship visual features without standing up CV infra.
Get an API key. Ship vision today.
Read the Docs Visual Intelligence for Enterprise. Build, run, and operate visual AI across images, documents, and video with production-grade accuracy, observability, and control.
Product updates and research. Unsubscribe in one click.
© 2026 Autonomi AI Inc. All rights reserved.
