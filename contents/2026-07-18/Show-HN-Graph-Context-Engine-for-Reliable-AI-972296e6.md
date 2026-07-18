---
source: "https://kritama.com"
hn_url: "https://news.ycombinator.com/item?id=48962441"
title: "Show HN: Graph Context Engine for Reliable AI"
article_title: "Kritama\n· Fractal Context Engine"
author: "zacksiri"
captured_at: "2026-07-18T21:40:48Z"
capture_tool: "hn-digest"
hn_id: 48962441
score: 1
comments: 0
posted_at: "2026-07-18T21:10:20Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Graph Context Engine for Reliable AI

- HN: [48962441](https://news.ycombinator.com/item?id=48962441)
- Source: [kritama.com](https://kritama.com)
- Score: 1
- Comments: 0
- Posted: 2026-07-18T21:10:20Z

## Translation

タイトル: Show HN: 信頼性の高い AI のためのグラフ コンテキスト エンジン
記事タイトル：クリタマ
· フラクタル コンテキスト エンジン
説明: 正確かつ効率的で、大規模に保守可能なデジタル アシスタントを構築します。
HN テキスト: 過去 2 年間、これに取り組んできました。サイトにはその仕組みを説明するビデオがあります。このエンジンを利用した実稼働アプリもあります。ぜひ試してみてください。どんな質問にも喜んでお答えします。

記事本文:
テーマの切り替え
信頼できるインテリジェンス
本番環境に耐えられるビルドアシスタント
クリタマ
正確かつ効率的で、大規模な保守が可能なデジタル アシスタントを構築します。
Kritama がどのようにコンテキストを構築し、モデルの動作を観察可能な状態に保ち、再利用可能なインテリジェンスを制作アシスタントに変えるかを詳しく見てみましょう。
制御できるインテリジェンス層
タスクの変化に応じて関連するコンテキストが出入りするため、不必要なノイズを発生させることなく、モデルは当面の問題に集中し続けることができます。
問題が発生した場所を確認し、再現して、体系的に問題を修正します。
N-01
アクティブ
N-03
アクティブ
N-08
アクティブ
N-10
故障
パッチ適用済み
N-14
アクティブ
N-16
アクティブ
小型モデルの利点
より小さいモデルを使用すると、コストが安くなり、スループットが高くなります。エージェントはより速く応答し、使用するトークンが少なくなります。
HCL は構造を定義し、マークダウンはインテリジェンスを定義します。ポリシー、ビジネス ロジック、推論をネットワークに直接組み込みます。
モジュール「memovee」{
ソース = "upmaru/base/tama//モジュール/メッセージング"
バージョン = "0.5.6"
depend_on = [モジュール .グローバル.スキーマ ]
名前 = "メモヴィー"
}
リソース "tama_prompt" "memovee" {
space_id = モジュール 。 memovee.space_id
name = "メモヴィーパーソナリティ"
役割 = 「システム」
content = ファイル ( "memovee/persona.md" )
}
リソース "tama_space_bridge" "memovee-basic" {
space_id = モジュール 。 memovee.space_id
target_space_id = tama_space.basic-conversation.id
}
リソース "tama_space_bridge" "memovee-media" {
space_id = モジュール 。 memovee.space_id
target_space_id = tama_space.media-conversation.id
}
リソース "tama_space_bridge" "memovee-ui" {
space_id = モジュール 。 memovee.space_id
target_space_id = tama_space.ui.id
}
# ペルソナ
あなたの名前は「メモヴィー」です。フレンドリーで熱心、知識豊富な映画の専門家として行動してください。情熱を持っておしゃべりするような、会話的で有益な口調である必要があります

映画マニアを食べました。
## コア機能
主な目的は、映画に関連する俳優、監督、賞、ジャンル、映画の歴史、エンターテインメント業界など、映画に焦点を当てたユーザーの問い合わせを支援することです。
## 機能
- 事実に関する質問に答えます (例: 公開日、キャスト/スタッフ、あらすじ *必要に応じてネタバレ警告あり*、興行収入データ、賞)。
- ユーザーの好み (ジャンル、俳優、雰囲気、類似タイトル) に基づいて映画を推奨します。
- 映画のテーマ、トリビア、批判的な評価について話し合います (個人的な意見ではなくレビューを要約します)。
- 映画の用語や概念を説明します。
- 映画がストリーミングされている場所、またはレンタル/購入できる場所を特定します (最新情報についてはツールを使用します)。
- ユーザーが TV シリーズ、シーズン、エピソード、または TV のみの推奨事項について質問した場合は、現在のシステムが映画のみをサポートしていることを明確に伝えてください。
## インタラクションガイドライン
- **正確性を最優先:** 正しい情報を提供することを優先します。答えがわからない、または確認できない場合は、そのことを明確に述べてください。憶測は避けてください。
- **説明:** ユーザーのリクエストがあいまいまたは曖昧な場合 (例: "S
[切り捨てられた]
ビジネス知識に基づいて構築された、正確で本番環境に対応したデジタル アシスタントを求める関心リストに参加してください。
Copyright © 2026 - 著作権は株式会社アップマルが保有します。

## Original Extract

Build digital assistants that are accurate, efficient, and maintainable at scale.

Been working on this for the last 2 years. There is a video on the site explaining how it works. There is also a production app powered by this engine. Give it a try, I'm happy to answer any questions.

Toggle theme
Reliable Intelligence
Build assistants that hold up in production
Kritama
Build digital assistants that are accurate, efficient, and maintainable at scale.
A closer look at how Kritama structures context, keeps model behavior observable, and turns reusable intelligence into production assistants.
The Intelligence layer you can control
Relevant context moves in and out as the task changes, keeping the model focused on the problem at hand without carrying unnecessary noise.
See where problems occur, reproduce them and fix problems systematically.
N-01
ACTIVE
N-03
ACTIVE
N-08
ACTIVE
N-10
FAULT
PATCHED
N-14
ACTIVE
N-16
ACTIVE
Small Model Advantage
Using smaller models means cheaper costs and higher throughput. Your agents respond faster and use fewer tokens.
HCL defines the structure, markdown defines the intelligence. Bake in your policies, business logic and reasoning right into the network.
module "memovee" {
source = "upmaru/base/tama//modules/messaging"
version = "0.5.6"
depends_on = [ module . global.schemas ]
name = "Memovee"
}
resource "tama_prompt" "memovee" {
space_id = module . memovee.space_id
name = "Memovee Personality"
role = "system"
content = file ( "memovee/persona.md" )
}
resource "tama_space_bridge" "memovee-basic" {
space_id = module . memovee.space_id
target_space_id = tama_space.basic-conversation.id
}
resource "tama_space_bridge" "memovee-media" {
space_id = module . memovee.space_id
target_space_id = tama_space.media-conversation.id
}
resource "tama_space_bridge" "memovee-ui" {
space_id = module . memovee.space_id
target_space_id = tama_space.ui.id
}
# Persona
Your name is 'Memovee'. Act as a friendly, enthusiastic, and knowledgeable movie expert. Your tone should be conversational and helpful, like chatting with a passionate film buff.
## Core Function
Your primary purpose is to assist users with movie-focused inquiries, including actors, directors, awards, genres, film history, and the entertainment industry as they relate to movies.
## Capabilities
- Answer factual questions (e.g., release dates, cast/crew, plot summaries *with spoiler warnings if necessary* , box office data, awards).
- Provide movie recommendations based on user preferences (genre, actors, mood, similar titles).
- Discuss movie themes, trivia, and critical reception (summarizing reviews rather than giving personal opinions).
- Explain film terminology or concepts.
- Identify where movies might be streaming or available for rent/purchase (use tools for current information).
- If a user asks about TV series, seasons, episodes, or TV-only recommendations, clearly say that the current system supports movies only.
## Interaction Guidelines
- **Accuracy First:** Prioritize providing correct information. If you don't know the answer or cannot verify it, explicitly state that. Avoid speculation.
- **Clarification:** If a user's request is vague or ambiguous (e.g., "S
[truncated]
Join the interest list for accurate, production-ready digital assistants built around your business knowledge.
Copyright © 2026 - All right reserved by Upmaru, Inc.
