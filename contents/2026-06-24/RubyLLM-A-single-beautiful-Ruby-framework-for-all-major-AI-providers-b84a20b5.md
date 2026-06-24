---
source: "https://rubyllm.com/"
hn_url: "https://news.ycombinator.com/item?id=48660711"
title: "RubyLLM: A single, beautiful Ruby framework for all major AI providers"
article_title: "RubyLLM | One beautiful Ruby framework for all major AI providers. Chat, images, embeddings, tools."
author: "doener"
captured_at: "2026-06-24T14:56:09Z"
capture_tool: "hn-digest"
hn_id: 48660711
score: 2
comments: 0
posted_at: "2026-06-24T14:41:41Z"
tags:
  - hacker-news
  - translated
---

# RubyLLM: A single, beautiful Ruby framework for all major AI providers

- HN: [48660711](https://news.ycombinator.com/item?id=48660711)
- Source: [rubyllm.com](https://rubyllm.com/)
- Score: 2
- Comments: 0
- Posted: 2026-06-24T14:41:41Z

## Translation

タイトル: RubyLLM: すべての主要な AI プロバイダー向けの単一の美しい Ruby フレームワーク
記事のタイトル: RubyLLM |すべての主要な AI プロバイダーに対応する 1 つの美しい Ruby フレームワーク。チャット、画像、埋め込み、ツール。
説明: すべての主要な AI プロバイダー向けの単一の美しい Ruby フレームワーク。チャットボット、AI エージェント、RAG アプリケーション、コンテンツ ジェネレーター、および考えられるすべての AI ワークフローを簡単に構築できます。

記事本文:
ルビーLLM |すべての主要な AI プロバイダーに対応する 1 つの美しい Ruby フレームワーク。チャット、画像、埋め込み、ツール。
メイン コンテンツにスキップ リンク メニュー 展開 (外部リンク) ドキュメント検索 コピー コピーされました RubyLLM
次へ (2.0 開発) 1.16.0 (最新) 変更履歴 ↗ ページをコピー スター
すべての主要な AI プロバイダー向けの単一の美しい Ruby フレームワーク。チャットボット、AI エージェント、RAG アプリケーション、コンテンツ ジェネレーター、および考えられるすべての AI ワークフローを簡単に構築できます。
戦闘テスト済み - 完全にプライベートな作業 AI
実用的な Ruby AI チャットを 2 分で構築
RubyLLM を使用していますか?あなたのストーリーをシェアしてください！ 5分かかります。
すべての AI プロバイダーは、肥大化した独自のクライアントを出荷しています。さまざまな API。さまざまな応答形式。さまざまな慣習。とても疲れます。
RubyLLM は、それらすべてに対応する 1 つの美しいフレームワークを提供します。 GPT、Claude、またはローカルの Ollama を使用している場合でも、同じインターフェイスです。依存関係は 3 つだけです: Faraday、Zeitwerk、Marcel。それでおしまい。
#ただ質問してください
チャット = RubyLLM 。チャット
チャット。 「Rubyを学ぶ最良の方法は何ですか?」と尋ねてください。
# あらゆる種類のファイルを分析します
チャット。 「この画像には何が写っていますか？」と尋ねます。 、「ruby_conf.jpg」を含む
チャット。 「このビデオで何が起こっているのですか？」と尋ねます。 、「video.mp4」を含む
チャット。 「meeting.wav」を使用して「この会議について説明してください」と尋ねます。
チャット。 「contract.pdf」を使用して「この文書を要約してください」と尋ねます。
チャット。 「app.rb」を使用して「このコードを説明してください」と尋ねます。
# 複数のファイルを一度に
チャット。 [ "diagram.png" 、 "report.pdf" 、 "notes.txt" ] を使用して、「これらのファイルを分析してください」と尋ねます。
# ストリーム応答
チャット。 「Ruby について話してください」と尋ねます。チャンク |
印刷チャンク。内容
終わり
# 画像を生成する
RubyLLM 。 「水彩風に山に沈む夕日」を描く
# 埋め込みを作成する
RubyLLM 。 embed "Ruby はエレガントで表現力豊かです"
# 音声をテキストに書き写す
RubyLLM 。 「meeting.wav」を転写する
# 安全のため適度なコンテンツ
RubyLLM 。中程度「これかどうか確認してください」

のテキストは安全です」
# AI にコードを使用させます
クラス天気 < RubyLLM :: ツール
説明「現在の天気を取得する」
def 実行 (緯度:、経度:)
URL = "https://api.open-meteo.com/v1/forecast?latitude= #{ 緯度 } &longitude= #{ 経度 } ¤t=温度_2m,風速_10m"
JSON 。 parse (ファラデー.get(url).body)
終わり
終わり
チャット。 with_tool (天気)。 「ベルリンの天気は？」と尋ねてください。
# 手順とツールを使用してエージェントを定義する
クラス WeatherAssistant < RubyLLM :: エージェント
モデル「gpt-5-nano」
指示は「簡潔にして、常に天候に関するツールを使用してください。」
ツール天気
終わり
天気アシスタント。新しい。 「ベルリンの天気は？」と尋ねてください。
# 構造化された出力を取得する
クラス ProductSchema < RubyLLM :: スキーマ
文字列:名前
番号：価格
配列:機能が実行する
文字列
終わり
終わり
応答 = チャット。 with_schema ( ProductSchema )。 「product.txt」を使用して「この製品を分析してください」と尋ねます。
特長
チャット: RubyLLM.chat を使用した会話型 AI
ビジョン: 画像とビデオを分析する
音声: RubyLLM.transcribe を使用して音声を文字に起こして理解する
ドキュメント: PDF、CSV、JSON、あらゆる種類のファイルから抽出
画像生成: RubyLLM.paint で画像を作成する
埋め込み: RubyLLM.embed を使用して埋め込みを生成する
モデレーション: RubyLLM.moderate によるコンテンツの安全性
ツール: AI に Ruby メソッドを呼び出してもらう
エージェント: RubyLLM::Agent を使用した再利用可能なアシスタント
構造化された出力: 正しく動作する JSON スキーマ
ストリーミング: ブロックによるリアルタイム応答
Rails: ActiveRecord と act_as_chat の統合
非同期: ファイバーベースの同時実行性
モデル レジストリ: 800 以上のモデルと機能の検出と価格設定
拡張思考: モデルの検討を制御、表示、持続する
プロバイダー: OpenAI、xAI、Anthropic、Gemini、VertexAI、Bedrock、DeepSeek、Mistral、Ollama、OpenRouter、Perplexity、GPUStack、および任意の OpenAI 互換 API
宝石「ruby_llm」
次に、バンドルインストールします。
# コンフィ

g/initializers/ruby_llm.rb
RubyLLM 。構成する |設定 |
構成。 openai_api_key = ENV [ 'OPENAI_API_KEY' ]
終わり
レール
# Rails統合をインストールする
bin/rails は、ruby_llm:install を生成します。
bin/rails db:移行
bin/rails Ruby_llm:load_models # v1.13+
# チャット UI の追加 (オプション)
bin/rails は Ruby_llm:chat_ui を生成します
クラス チャット < ApplicationRecord
チャットとして機能する
終わり
チャット = チャット。作成する！モデル：「クロード・ソネット-4」
チャット。 「このファイルには何が入っていますか?」と尋ねます。 、「レポート.pdf」付き
すぐに使用できるチャット インターフェイスについては、http://localhost:3000/chats にアクセスしてください。

## Original Extract

A single, beautiful Ruby framework for all major AI providers. Easily build chatbots, AI agents, RAG applications, content generators, and every AI workflow you can think of.

RubyLLM | One beautiful Ruby framework for all major AI providers. Chat, images, embeddings, tools.
Skip to main content Link Menu Expand (external link) Document Search Copy Copied RubyLLM
next (2.0 dev) 1.16.0 (latest) Changelog ↗ Copy page Star
A single, beautiful Ruby framework for all major AI providers. Easily build chatbots, AI agents, RAG applications, content generators, and every AI workflow you can think of.
Battle tested at - Fully private work AI
Build a working Ruby AI chat in two minutes
Using RubyLLM? Share your story ! Takes 5 minutes.
Every AI provider ships their own bloated client. Different APIs. Different response formats. Different conventions. It’s exhausting.
RubyLLM gives you one beautiful framework for all of them. Same interface whether you’re using GPT, Claude, or your local Ollama. Just three dependencies: Faraday, Zeitwerk, and Marcel. That’s it.
# Just ask questions
chat = RubyLLM . chat
chat . ask "What's the best way to learn Ruby?"
# Analyze any file type
chat . ask "What's in this image?" , with: "ruby_conf.jpg"
chat . ask "What's happening in this video?" , with: "video.mp4"
chat . ask "Describe this meeting" , with: "meeting.wav"
chat . ask "Summarize this document" , with: "contract.pdf"
chat . ask "Explain this code" , with: "app.rb"
# Multiple files at once
chat . ask "Analyze these files" , with: [ "diagram.png" , "report.pdf" , "notes.txt" ]
# Stream responses
chat . ask "Tell me a story about Ruby" do | chunk |
print chunk . content
end
# Generate images
RubyLLM . paint "a sunset over mountains in watercolor style"
# Create embeddings
RubyLLM . embed "Ruby is elegant and expressive"
# Transcribe audio to text
RubyLLM . transcribe "meeting.wav"
# Moderate content for safety
RubyLLM . moderate "Check if this text is safe"
# Let AI use your code
class Weather < RubyLLM :: Tool
desc "Get current weather"
def execute ( latitude :, longitude :)
url = "https://api.open-meteo.com/v1/forecast?latitude= #{ latitude } &longitude= #{ longitude } &current=temperature_2m,wind_speed_10m"
JSON . parse ( Faraday . get ( url ). body )
end
end
chat . with_tool ( Weather ). ask "What's the weather in Berlin?"
# Define an agent with instructions + tools
class WeatherAssistant < RubyLLM :: Agent
model "gpt-5-nano"
instructions "Be concise and always use tools for weather."
tools Weather
end
WeatherAssistant . new . ask "What's the weather in Berlin?"
# Get structured output
class ProductSchema < RubyLLM :: Schema
string :name
number :price
array :features do
string
end
end
response = chat . with_schema ( ProductSchema ). ask "Analyze this product" , with: "product.txt"
Features
Chat: Conversational AI with RubyLLM.chat
Vision: Analyze images and videos
Audio: Transcribe and understand speech with RubyLLM.transcribe
Documents: Extract from PDFs, CSVs, JSON, any file type
Image generation: Create images with RubyLLM.paint
Embeddings: Generate embeddings with RubyLLM.embed
Moderation: Content safety with RubyLLM.moderate
Tools: Let AI call your Ruby methods
Agents: Reusable assistants with RubyLLM::Agent
Structured output: JSON schemas that just work
Streaming: Real-time responses with blocks
Rails: ActiveRecord integration with acts_as_chat
Async: Fiber-based concurrency
Model registry: 800+ models with capability detection and pricing
Extended thinking: Control, view, and persist model deliberation
Providers: OpenAI, xAI, Anthropic, Gemini, VertexAI, Bedrock, DeepSeek, Mistral, Ollama, OpenRouter, Perplexity, GPUStack, and any OpenAI-compatible API
gem 'ruby_llm'
Then bundle install .
# config/initializers/ruby_llm.rb
RubyLLM . configure do | config |
config . openai_api_key = ENV [ 'OPENAI_API_KEY' ]
end
Rails
# Install Rails Integration
bin/rails generate ruby_llm:install
bin/rails db:migrate
bin/rails ruby_llm:load_models # v1.13+
# Add Chat UI (optional)
bin/rails generate ruby_llm:chat_ui
class Chat < ApplicationRecord
acts_as_chat
end
chat = Chat . create! model: "claude-sonnet-4"
chat . ask "What's in this file?" , with: "report.pdf"
Visit http://localhost:3000/chats for a ready-to-use chat interface!
