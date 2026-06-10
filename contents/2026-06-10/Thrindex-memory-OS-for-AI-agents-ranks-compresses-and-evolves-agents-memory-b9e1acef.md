---
source: "https://www.thrindex.com/"
hn_url: "https://news.ycombinator.com/item?id=48481780"
title: "Thrindex – memory OS for AI agents (ranks, compresses and evolves agents memory)"
article_title: "Thrindex - The memory OS for AI agents."
author: "teo-nomikos"
captured_at: "2026-06-10T21:45:13Z"
capture_tool: "hn-digest"
hn_id: 48481780
score: 1
comments: 0
posted_at: "2026-06-10T19:59:35Z"
tags:
  - hacker-news
  - translated
---

# Thrindex – memory OS for AI agents (ranks, compresses and evolves agents memory)

- HN: [48481780](https://news.ycombinator.com/item?id=48481780)
- Source: [www.thrindex.com](https://www.thrindex.com/)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T19:59:35Z

## Translation

タイトル: Thrindex – AI エージェント用のメモリ OS (エージェントのメモリをランク付け、圧縮、進化させます)
記事のタイトル: Thrindex - AI エージェント用のメモリ OS。
説明: ランク付け、圧縮、進化する認識エンジン
あらゆるモデル、あらゆるフレームワーク、あらゆる規模のエージェント メモリ。

記事本文:
あらゆるモデル、あらゆるフレームワーク、あらゆる規模でエージェント メモリをランク付け、圧縮、進化させる認識エンジン。ストレージを超えて。 RAGを超えて。
あらゆるモデル、あらゆるフレームワーク、あらゆる規模でエージェント メモリをランク付け、圧縮、進化させる認識エンジン。ストレージを超えて。 RAGを超えて。
これまでに実行するすべてのエージェント向けに構築されたメモリ インフラストラクチャ
これまでに実行するすべてのエージェント向けに構築されたメモリ インフラストラクチャ
個人の開発者から何百ものエージェントを実行するチームまで、Thrindex はすべてのエージェントに、改善を続けるために必要なメモリ層を提供します。
SDK コードをコピー Python typescript 1 pip install thrindex 2 3 from thrindex import Thrindex 4 client = Thrindex(api_key= "th_live_..." ) 5 6 # メモリを保存 7 client.add( 8 content= "ユーザーは毎週のメール ダイジェストを好みます。" , 9 Agent_id= "support-agent" , 10 user_id= "user-42" , 11 ) 12 # メモリの検索 13 results = client.search( 14 query= "ユーザーはどのように連絡を希望しますか" , 15 Agent_id= "support-agent" , 16 user_id= "user-42" , 17 )
API コピー コード Python Node.js cURL 1 import httpx 2 3 client = httpx.Client( 4 headers={ "Authorization" : "Bearer th_live_..." } 5 ) 6 7 # メモリを保存する 8 client.post( "https://api.thrindex.com/v1/memories" , json={ 9 "content" : "ユーザーは毎週のメール ダイジェストを好みます。" , 10 "agent_id" : "support-agent" , 11 "user_id" : "user-42" , 12 }) 13 # メモリの検索 14 resp = client.post( "https://api.thrindex.com/v1/memories/search" , json={ 15 "query" : "ユーザーはどのように連絡を希望していますか" , 16 "agent_id" : "support-agent" , 17 "user_id" : "user-42" , 18 }) 19 print(resp.json()[ "results" ])
エージェントは書き込みと検索を行います。 Cognition はバックグラウンドで実行されます。エージェントが待機している間は実行されません。
POST /v1/memories/search は、事前にランク付けされたキャッシュ + マルチシグナル スコアリングを使用します。読み取りパスに LLM がありません。
POST /v1/memories/search は pre-

ランク付けされたキャッシュ + マルチシグナル スコアリング。読み取りパスに LLM がありません。
POST /v1/memories はすぐに戻ります。重複排除、重要性、圧縮、競合解決は非同期で実行されます。
POST /v1/memories はすぐに戻ります。重複排除、重要性、圧縮、競合解決は非同期で実行されます。
すべてのエージェントのメモリ、1 つのダッシュボード
エージェントが保存するもの、さらにキーと分析を参照、検索、マップします。実稼働環境で AI を出荷するチーム向けに構築されています。
エージェントの記憶全体が一度に表示されます。何が関連しているのか、何が冗長なのか、そしてどこに知識が不足しているのかを確認します。
エージェントの記憶全体が一度に表示されます。何が関連しているのか、何が冗長なのか、そしてどこに知識が不足しているのかを確認します。
エージェントが実行する正確な検索を実行します。何が返されるか、どれくらいの速度で返されるか、キャッシュから返されたかどうかを確認します。
エージェントが実行する正確な検索を実行します。何が返されるか、どれくらいの速度で返されるか、キャッシュから返されたかどうかを確認します。
レイテンシ、キャッシュ ヒット、オペレーション ボリューム、キューの深さを 1 つのビューで確認できます。本番環境で実際に監視できるメモリ。
レイテンシ、キャッシュ ヒット、オペレーション ボリューム、キューの深さを 1 つのビューで確認できます。本番環境で実際に監視できるメモリ。
財産保護、保険の基本、補償を最大限に活用するための実際のヒントに関する役立つ記事をご覧ください。
エージェントは覚えていました - 間違って覚えていただけです
エージェントがどのようにして事実を完璧に思い出しながらも間違っているのか、そしてそれがエージェントに依存する人々に何をもたらすのか。
メモリは検索の問題ではありません
検索がエージェントの記憶の簡単な半分である理由と、実際のエンジニアリングが存在する場所についての技術的な考察。
思い出しすぎたエージェント
すべての記憶をゆっくりと保持することがエージェントの状態を悪化させる理由と、どのようにして忘れることが特徴になるのか。
エージェントに永続的なメモリを提供します。重要なものをランク付けし、そうでないものを圧縮し、忘れるべきものを忘れます。
エージェントに永続的なメモリを提供します。重要なものをランク付けし、何を圧縮するか

そうしない、忘れるべきものを忘れる。
エージェントに永続的なメモリを提供します。重要なものをランク付けし、そうでないものを圧縮し、忘れるべきものを忘れます。
いくつかのことは覚えておく価値があります。
© 2026 thindex.無断転載を禁じます。
いくつかのことは覚えておく価値があります。
© 2026 thindex.無断転載を禁じます。
いくつかのことは覚えておく価値があります。
© 2026 thindex.無断転載を禁じます。

## Original Extract

A cognition engine that ranks, compresses and evolves
agent memory across any model, any framework, at any scale.

A cognition engine that ranks, compresses and evolvesagent memory across any model, any framework, at any scale. Beyond storage. Beyond RAG.
A cognition engine that ranks, compresses and evolvesagent memory across any model, any framework, at any scale. Beyond storage. Beyond RAG.
Memory infrastructure built for every agent you'll ever run
Memory infrastructure built for every agent you'll ever run
From solo developers to teams running hundreds of agents, Thrindex gives every agent the memory layer it needs to keep improving.
SDK Copy Code Python typescript 1 pip install thrindex 2 3 from thrindex import Thrindex 4 client = Thrindex(api_key= "th_live_..." ) 5 6 # Store a memory 7 client.add( 8 content= "User prefers weekly email digests." , 9 agent_id= "support-agent" , 10 user_id= "user-42" , 11 ) 12 # Search memories 13 results = client.search( 14 query= "how does the user want to be contacted" , 15 agent_id= "support-agent" , 16 user_id= "user-42" , 17 )
API Copy Code Python Node.js cURL 1 import httpx 2 3 client = httpx.Client( 4 headers={ "Authorization" : "Bearer th_live_..." } 5 ) 6 7 # Store a memory 8 client.post( "https://api.thrindex.com/v1/memories" , json={ 9 "content" : "User prefers weekly email digests." , 10 "agent_id" : "support-agent" , 11 "user_id" : "user-42" , 12 }) 13 # Search memories 14 resp = client.post( "https://api.thrindex.com/v1/memories/search" , json={ 15 "query" : "how does the user want to be contacted" , 16 "agent_id" : "support-agent" , 17 "user_id" : "user-42" , 18 }) 19 print(resp.json()[ "results" ])
Agents write and search. Cognition runs in the background - never while your agent waits.
POST /v1/memories/search uses pre-ranked cache + multi-signal scoring. No LLM on the read path.
POST /v1/memories/search uses pre-ranked cache + multi-signal scoring. No LLM on the read path.
POST /v1/memories returns immediately. Dedup, importance, compression and conflict resolution run async.
POST /v1/memories returns immediately. Dedup, importance, compression and conflict resolution run async.
Every agent memory, one dashboard
Browse, search and map what agents store - plus keys and analytics. Built for teams shipping AI in production.
Your agent's entire memory, visible at once. See what's related, what's redundant and where knowledge is thin.
Your agent's entire memory, visible at once. See what's related, what's redundant and where knowledge is thin.
Run the exact search your agent runs. See what it gets back, how fast and whether it came from cache.
Run the exact search your agent runs. See what it gets back, how fast and whether it came from cache.
Latency, cache hits, ops volume and queue depth in one view. Memory you can actually monitor in production.
Latency, cache hits, ops volume and queue depth in one view. Memory you can actually monitor in production.
Discover useful articles about property protection, insurance basics, and real-world tips that help you get the most from your coverage.
The agent remembered - it just remembered wrong
How an agent can recall a fact perfectly and still be wrong and what that costs the people who rely on it.
Memory is not a search problem
A technical look at why retrieval is the easy half of agent memory and where the real engineering lives.
The agent that remembered too much
Why keeping every memory slowly makes an agent worse and how forgetting becomes a feature.
Give your agents persistent memory. Ranks what matters, compresses what doesn't, forgets what should be forgotten.
Give your agents persistent memory. Ranks what matters, compresses what doesn't, forgets what should be forgotten.
Give your agents persistent memory. Ranks what matters, compresses what doesn't, forgets what should be forgotten.
Some things are worth remembering.
© 2026 thrindex. All rights reserved.
Some things are worth remembering.
© 2026 thrindex. All rights reserved.
Some things are worth remembering.
© 2026 thrindex. All rights reserved.
