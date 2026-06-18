---
source: "https://console.cocore.dev/"
hn_url: "https://news.ycombinator.com/item?id=48590929"
title: "Co/Core: An AI Cooperative"
article_title: "co/core — an AI cooperative"
author: "lisowski"
captured_at: "2026-06-18T21:45:55Z"
capture_tool: "hn-digest"
hn_id: 48590929
score: 6
comments: 0
posted_at: "2026-06-18T20:18:09Z"
tags:
  - hacker-news
  - translated
---

# Co/Core: An AI Cooperative

- HN: [48590929](https://news.ycombinator.com/item?id=48590929)
- Source: [console.cocore.dev](https://console.cocore.dev/)
- Score: 6
- Comments: 0
- Posted: 2026-06-18T20:18:09Z

## Translation

タイトル: Co/Core: AI 協同組合
記事のタイトル: co/core — AI 協同組合
説明: co/core は AI 推論のための協同組合です。大きなクラウドから借りるのではなく、人々がすでに所有している Mac をプールして、お互いにオープン モデルを実行します。私たちが共同で構築、共有、所有する AI インフラストラクチャの実験。既存の OpenAI 互換コードを持ち込むか、Mac を共有して ru を手伝ってください
[切り捨てられた]

記事本文:
co/core — AI 協同組合 co/core モデルは、リーダーボードを探索します。 ブログ API メンバー所有の AI の実験にログインします AI 協同組合。
co/core は、人々が少数の巨大プロバイダーからレンタルするのではなく、AI を実行するためにすでに所有しているコンピューティングを共有する場所です。これは、私たちが一緒に構築、共有、所有する推論の実験です。他のコードと同じ標準 API を使用するため、既存のコードはそのまま機能します。これらのモデルは、人々がすでに持っているハードウェア上で動作するオープンなモデルです。
本物のメンバー。本物のマシン。本物の領収書。
現在の生協の様子を垣間見ることができます。サインアップした人々、コンピューティングを共有しているマシンの一部、そしてサインオフした最新のジョブです。これらはすべて、本物の公的記録です。
Cocore は他のユーザーと同じ API 言語を話します。既存の SDK を console.cocore.dev/v1 に指定し、 cocore-… キーをドロップして続行します。ストリーミング、ツール呼び出し、通常のチャット/補完はすべて機能し、コードは変更されません。プリセットまたは任意の MLX モデルからホストします。
Anthropic Together Groq Ollama co/core client.py openai import OpenAI から使用する前
クライアント = OpenAI(
Base_url= "https://api.openai.com/v1" ,
api_key= "sk-proj-…" ,
)
resp = client.chat.completions.作成(
モデル= "gpt-4o-mini" 、
メッセージ=[ { "ロール" : "ユーザー" ,
"コンテンツ" : "こんにちは" } ],
ストリーム= True 、
) openai import OpenAI からの client.py
クライアント = OpenAI(
Base_url= "https://console.cocore.dev/v1" ,
api_key= "cocore-7f3a2c…" ,
)
resp = client.chat.completions.作成(
モデル= "mlx-community/Qwen2.5-0.5B" ,
メッセージ=[ { "ロール" : "ユーザー" ,
"コンテンツ" : "こんにちは" } ],
ストリーム= True 、
) 02 仕組み リクエストを送信します。生協が運営しています。
リクエストにより、メンバーの利用可能なコンピューティングが検索され、ジョブが実行され、ユーザーに返されます。全体が薄い

g はオープン仕様です。すべてのジョブは、私たちのものも含め、誰もが自分で検証できる署名付きの公開記録を残します。
暗号化されたプロンプトが生協に送信されます。オープン ネットワーク内の任意の場所で特定のモデルのプロンプトを実行するか、信頼できる協力サークルの友人とそのマシンのみにジョブを非公開にすることを選択します。
誰かがあなたのジョブを取得し、実行し、ハードウェアから決して離れることのない Secure Enclave にロックされたキーで署名した後、結果を返します。これは、誰がその作業を行ったのか、そしてその後誰もそのジョブに触らなかったことの正確な証拠となります。
領収書があれば仕事は終了します。共有ポットへの少額の取り分を除いて、その仕事を実行した人にクレジットが与えられます。毎月、そのポットは参加者が参加した金額に応じてメンバーに分配されるため、削減額はネットワークを運営している人々に直接還元されます。
私たちの言葉をそのまま信じる必要はありません。
すべての仕事は、私たちに電話することなく、誰もが自分で確認できる領収書を書きます。残高を水増ししたり、支払いを偽ったり、ルールをこっそり変更したりすることはできません。また、私たちのやり方が気に入らない場合は、すべてを自分で実行することもできます。自分のコピーを同じデータに指定すると、まったく同じ数値が得られます。
これは実際の領収書と、それが示す記録、つまりあなたが送信したジョブ、マシンの認証、公的決済です。チェーンに従って、自分で確認してください。
オープンスタンダード。独自の取引所を運営します。
Co-op の下には、推論リクエストの作成と、その後に起こったことの記録、つまりすべての仕事についての署名付きの公開アカウントという 2 つの単純な事柄に関するオープン スタンダードがあります。私たちはその上で 1 つの取引所を運営しており、そのルールは次のとおりです。1 ユニットがイン、1 ユニットがアウト、共有ポットはメンバーに還元されます。しかし、ここでは標準が実際の作業です。誰でも、独自の価格設定、独自のメンバーシップ、何が公正であるかについての独自の考えなど、異なる取引所を立ち上げることができます。

まったく同じレコードの読み取りと書き込み。私たちのものは最初のものにすぎません。 co/core が希望のバージョンでなくなった場合でも、変更を依頼する必要はありません。自分のものを作りに行くことができます。
コンピューティングを共有します。生協の運営を手伝ってください。
cocore アプリをダウンロードし、サインインし、モデルを 1 つまたは 2 つ選択すると、コンピューターがネットワークに参加し、コンピューターを使用していないときに他のメンバーのためにジョブを実行します。ターミナルも子守りも何もありません。それ自体が更新されます。協同組合が実行するすべての仕事が残高に加算され、生協が行うすべてのことの分け前が毎月組合員に還元されます。
© 2026 コ/コア。無断転載を禁じます。

## Original Extract

co/core is a cooperative for AI inference — people pooling the Macs they already own to run open models for each other, instead of renting from the big clouds. An experiment in AI infrastructure we build, share, and own together. Bring your existing OpenAI-compatible code, or share a Mac and help ru
[truncated]

co/core — an AI cooperative co/core models explore leaderboard blog API Log in an experiment in member-owned AI An AI cooperative.
co/core is a place where people share the compute they already own to run AI for each other, instead of renting from a handful of giant providers. It's an experiment in inference we build, share, and own together — and your existing code works as-is, because we speak the same standard API everything else does. The models are open ones that run on the hardware people already have.
Real members. Real machines. Real receipts.
A glimpse of the co-op right now — some of the people who've signed up, some of the machines sharing compute, and the latest jobs they've signed off on. Every one of these is a real, public record.
Cocore speaks the same API language as everybody else. Point your existing SDK at console.cocore.dev/v1 , drop in a cocore-… key, and keep going — streaming, tool calls, and the usual chat/completions shape all work, no code changes. Host from our presets or any MLX model .
Anthropic Together Groq Ollama co/core client.py before from openai import OpenAI
client = OpenAI(
base_url= "https://api.openai.com/v1" ,
api_key= "sk-proj-…" ,
)
resp = client.chat.completions. create (
model= "gpt-4o-mini" ,
messages=[ { "role" : "user" ,
"content" : "hello" } ],
stream= True ,
) client.py after from openai import OpenAI
client = OpenAI(
base_url= "https://console.cocore.dev/v1" ,
api_key= "cocore-7f3a2c…" ,
)
resp = client.chat.completions. create (
model= "mlx-community/Qwen2.5-0.5B" ,
messages=[ { "role" : "user" ,
"content" : "hello" } ],
stream= True ,
) 02 how it works You send a request. The co-op runs it.
Your request finds a member's available compute, runs the job, and comes back to you. The whole thing is an open spec — every job leaves a signed, public record anyone can verify for themselves, ours included.
Your encrypted prompt heads out into the co-op. Run your prompt for a given model anywhere in the open network, or choose to keep your jobs private to only a trusted circle of co-op friends and their machines.
Someone picks up your job, runs it, then returns the result after signing with a key locked in the Secure Enclave that never leaves their hardware — proof of exactly who did the work, and that nobody touched it after.
A receipt closes out the job — credits go to whomever who ran it, minus a small cut to a shared pot. Each month that pot splits back to members by how much they pitched in, so the cut cycles right back to the people running the network.
You don't have to take our word for it.
Every job writes a receipt anyone can verify on their own, without ever calling us. We can't inflate a balance, fake a payout, or quietly change the rules. And if you don't like how we run things, you can run the whole thing yourself — point your own copy at the same data and it lands on exactly the same numbers.
Here's a real receipt and the records it points back to — the job you sent, the machine's attestation, and the public settlement. Follow the chain and check it yourself.
An open standard. Run your own exchange.
Underneath the co-op is an open standard for two plain things: making an inference request, and recording what happened afterward — a signed, public account of every job. We run one exchange on top of it, with our rules: one unit in, one unit out, a shared pot that goes back to members. But the standard is the real work here. Anyone can stand up a different exchange — their own pricing, their own membership, their own idea of what's fair — reading and writing the very same records. Ours is just the first one. If co/core ever stops being the version you want, you don't have to ask us to change it; you can go build yours.
Share your compute. Help run the co-op.
Download the cocore app, sign in, pick a model or two, and your computer joins the network — running jobs for other members while you're not using it. No Terminal, nothing to babysit; it updates itself. Every job it runs adds to your balance, and a share of everything the co-op does comes back to members each month.
© 2026 co/core. All rights reserved.
