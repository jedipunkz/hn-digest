---
source: "https://nobodywho.ooo/posts/apple-watch-vision-pro-apps/"
hn_url: "https://news.ycombinator.com/item?id=48397198"
title: "Show HN: We've built a standalone Apple Watch app running LLMs offline, locally"
article_title: "Apple Watch & Vision Pro apps - NobodyWho"
author: "pielouNW"
captured_at: "2026-06-04T13:14:37Z"
capture_tool: "hn-digest"
hn_id: 48397198
score: 1
comments: 0
posted_at: "2026-06-04T11:35:52Z"
tags:
  - hacker-news
  - translated
---

# Show HN: We've built a standalone Apple Watch app running LLMs offline, locally

- HN: [48397198](https://news.ycombinator.com/item?id=48397198)
- Source: [nobodywho.ooo](https://nobodywho.ooo/posts/apple-watch-vision-pro-apps/)
- Score: 1
- Comments: 0
- Posted: 2026-06-04T11:35:52Z

## Translation

タイトル: HN を表示: LLM をオフラインでローカルで実行するスタンドアロンの Apple Watch アプリを構築しました
記事のタイトル: Apple Watch と Vision Pro アプリ - NobodyWho
説明: NobodyWho Swift バインディングが利用可能になりました。この記事では、開発中に取り組んだ興味深い事柄のいくつかを簡単に紹介します。

記事本文:
Apple Watch & Vision Pro アプリ - NobodyWho
誰でもない
について
アプリ
ブログ
ドキュメント
GitHub
ギットハブ
誰でもない
について
アプリ
ブログ
ドキュメント
GitHub
ギットハブ
Apple Watch と Vision Pro アプリ
最初の 2 つのアプリのリリースを発表できることを嬉しく思います。どちらも、NobodyWho 推論エンジンを搭載し、完全にオフラインで実行されるパーソナル AI アシスタントです。
これはスタンドアロンの Apple Watch アプリです。つまり、iOS コンパニオン アプリなしで動作します。これは、AI モデルをローカルで完全にオフラインで実行できる世界初の Apple Watch アプリです。
アイデアは、外出先やオフグリッドで簡単な質問に一発で回答できるアプリを作成することでした。そのため、わざわざ会話を続ける必要はありませんでした。
Metal アクセラレーションを使用しない Apple Watch ではコンピューティング能力が制限されているため、実行できるモデルが制限されます。現在、7 億パラメータの LFM2.5 (Apple Watch Ultra) や 17 億パラメータの 1-bit Bonsai (Apple Watch Series 11) など、Prism ML と Liquid AI のモデルをサポートしています。
小型モデルが時間の経過とともにより賢くなるにつれて、Apple Watch でより良い答えへの扉が開かれており、私たちは新しい小型言語モデルが登場することを非常に楽しみにしています。
Apple Watch アプリによって設定された道に従って、私たちは大規模で有能な思考モデルを実行できるアプリを構築することにしました。 LLMStream ライブラリのおかげで、ストリーミングされた回答は美しくレンダリングされます。 Prism ML と Liquid AI のモデルに加え、思考モデルの Qwen3 もサポートされています。
私たちは NobodyWho です。Swift や他のいくつかの言語のエッジ デバイス上で小さなモデルを実行できるようにするローカル推論ライブラリです。
私たちは、オープンソース コード、モデルの制御、堅実なソフトウェア エンジニアリング、標準化、シンプルなものをシンプルにすることを重視しています。
今日の AI の世界にはそのすべてが欠けています。私たちと一緒にモデルを実行するのは次のとおりです

次のように簡単です:
インポート
let chat = チャットを待ってみる。 fromPath (modelPath : "./model.gguf" )
let 応答 = チャットを待ってみてください。尋ねます (「デンマークの首都はどこですか?」)。完了しました ( ) ;
// デンマークの首都はコペンハーゲンです。
同じことを重視する場合は、寄稿者になるか、単にライブラリをダウンロードしてテストしてください。
×
×
マストドン
マストドン
ブルースカイ
ブルースカイ
リンクトイン
リンクトイン
ハグフェイス
ハグフェイス
不和
不和
電子メール
電子メール
RSS
RSS
誰でもない

## Original Extract

NobodyWho Swift bindings are now available! This article briefly introduces some of the interesting things we dealt with during development.

Apple Watch & Vision Pro apps - NobodyWho
NobodyWho
About
Apps
Blog
Docs
GitHub
Github
NobodyWho
About
Apps
Blog
Documentation
GitHub
Github
Apple Watch & Vision Pro apps
We are excited to announce the release of our first two apps! Both are personal AI assistants that run completely offline, powered by the NobodyWho inference engine.
This is a standalone Apple Watch app, meaning it works without any iOS companion app. It's the world's first Apple Watch app that can run an AI model locally, completely offline!
The idea was to create an app that could answer simple questions on the go, off the grid, in a single shot, which is why we didn't bother persisting conversations.
Since computing power is limited on the Apple Watch without Metal acceleration, we are restricted in the models we can run. Currently, we support models from Prism ML and Liquid AI , such as the LFM2.5 with 700 million parameters (Apple Watch Ultra) and the 1-bit Bonsai with 1.7 billion parameters (Apple Watch Series 11).
As smaller models get smarter over time, the door opens to better answers on the Apple Watch, and we are very excited for the new small language models to come!
Following the path set by the Apple Watch app, we decided to build an app that can run large, capable, thinking models. Streamed answers render beautifully thanks to the LLMStream library. Models from Prism ML and Liquid AI are supported as well, along with Qwen3, a thinking model.
We're NobodyWho, a local inference library that enables running small models on edge devices for Swift and several other languages.
We value open-source code, control over your models, solid software engineering, standardization, and making simple things simple.
All of which is missing in today's AI world. Running a model with us is as easy as:
import NobodyWho
let chat = try await Chat . fromPath ( modelPath : "./model.gguf" )
let response = try await chat . ask ( "What is the capital of Denmark?" ) . completed ( ) ;
// The capital of Denmark is Copenhagen.
If you value the same things, come and become a contributor or just download and test our library .
X
X
Mastodon
Mastodon
Bluesky
Bluesky
Linkedin
Linkedin
Hugging Face
Hugging Face
Discord
Discord
Email
Email
RSS
RSS
NobodyWho
