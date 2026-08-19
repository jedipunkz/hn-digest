---
source: "https://alicebraincore.web.app"
hn_url: "https://news.ycombinator.com/item?id=49357139"
title: "BYOKs for an LLM with a Brain"
article_title: "A.L.I.C.E. Brain Core"
image: ""
author: "kindsocial"
captured_at: "2026-08-19T05:20:43Z"
capture_tool: "hn-digest"
hn_id: 49357139
score: 1
comments: 0
posted_at: "2026-08-19T05:15:27Z"
tags:
  - hacker-news
  - translated
---

# BYOKs for an LLM with a Brain

- HN: [49357139](https://news.ycombinator.com/item?id=49357139)
- Source: [alicebraincore.web.app](https://alicebraincore.web.app)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T05:15:27Z

## Translation

タイトル: Brain を備えた LLM の BYOK
記事タイトル: A.L.I.C.E.ブレインコア
説明: HTTP API を介したアプリのメモリと推論。独自のモデル キーと Neo4j を持参してください。脳は、何を覚え、何を表面化し、何を省略するかを処理します。

記事本文:
A.L.I.C.E.ブレインコア A.L.I.C.E.ブレインコア
HTTP API を介したアプリのメモリと推論。
モデルは口です。会話している相手の記憶がないので、それをラップするすべてのアプリは、ひどいことに、プロンプトに適合するものから、セッションごとにコンテキストを最初から再構築します。これは残りの半分です。誰かが実際に言ったことを蓄積し、その前にあるメッセージにとって重要なものに絞り込み、それのみをモデルに渡すグラフです。
モデル キー、埋め込みキー、Neo4j インスタンスの 3 つを用意します。脳は、何を覚え、何を表面化し、何を省略するかという難しい部分をもたらします。
Anthropic キーが音声を実行します。 Gemini キーは埋め込みを実行します。 Neo4j はグラフを保持します。すべてはテナントごとの KMS エンベロープで暗号化され、使用時にのみ復号化されます。私たちはあなたの推論、ストレージ、リコールに料金を支払うことはありません。これが、あなたがどれだけ考えても一律料金が機能する理由であり、あなたのデータが他人のデータベースに決して置かれない理由です。
POST /v1/filter — ターン: 取得、注目、絞り込み、発言。
POST /v1/ingest — ターンが学習した内容を書き戻します。
独自の Neo4j — Aura 無料インスタンスで始めるのに十分です。
事実、定められた指示、開いているスレッド、エンティティと関係、気分、状況。
すべての事実には、それがどこから来たのか、最後に確認されたのはいつなのかが記載されています。
カードも試用時計もありません。キーとグラフを持参していただくため、使用料はかかりません。そのため、料金を請求する価値があるまでこれをオープンのままにしておくことができます。
キーはすぐに発行され、一度表示されます。
統合マニュアル全体は公開されています。最初に読んで、それが適切かどうかを判断してください。代わりに彼女に話しかける · 健康 · テストコンソール

## Original Extract

Memory and reasoning for your app, over an HTTP API. Bring your own model keys and your own Neo4j; the brain handles what to remember, what to surface, and what to leave out.

A.L.I.C.E. Brain Core A.L.I.C.E. Brain Core
Memory and reasoning for your app, over an HTTP API.
A model is the mouth. It has no memory of the person it is talking to, so every app wrapping one rebuilds context from scratch each session, badly, out of whatever fits in a prompt. This is the other half: a graph that accumulates what someone actually said, narrows it to what matters for the message in front of it, and hands the model only that.
You bring three things — a model key, an embedding key, and a Neo4j instance. The brain brings the part that is hard: what to remember, what to surface, and what to leave out.
Your Anthropic key runs the voice. Your Gemini key runs the embeddings. Your Neo4j holds the graph. All of it is encrypted with per-tenant KMS envelopes and decrypted only at the moment of use. We never pay for your inference, your storage, or your recall — which is why one flat price works no matter how much you think, and why your data never sits in somebody else's database.
POST /v1/filter — a turn: retrieve, notice, narrow, speak.
POST /v1/ingest — write back what the turn learned.
Your own Neo4j — an Aura free instance is enough to start.
Facts, standing instructions, open threads, entities and relations, moods, conditions.
Every fact carries where it came from and when it was last confirmed.
No card, no trial clock. You bring the keys and the graph, so your usage costs us nothing — which is why we can leave this open until it's worth charging for.
Your key is issued straight away and shown to you once.
The whole integration manual is public — read it first and decide if it fits. talk to her instead · health · test console
