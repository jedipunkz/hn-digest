---
source: "https://agentbadge.xyz/blog/what-is-agent-readiness"
hn_url: "https://news.ycombinator.com/item?id=49361630"
title: "Show HN: AgentBadge – Agent Readiness Scoring for APIs (SEO for AI Agents)"
article_title: "What Is Agent Readiness? — AgentBadge Blog — AgentBadge"
image: "https://agentbadge.xyz/images/blog/what-is-agent-readiness-og.png"
author: "spread2009"
captured_at: "2026-08-19T14:24:02Z"
capture_tool: "hn-digest"
hn_id: 49361630
score: 1
comments: 0
posted_at: "2026-08-19T13:52:38Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AgentBadge – Agent Readiness Scoring for APIs (SEO for AI Agents)

- HN: [49361630](https://news.ycombinator.com/item?id=49361630)
- Source: [agentbadge.xyz](https://agentbadge.xyz/blog/what-is-agent-readiness)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T13:52:38Z

## Translation

タイトル: HN を表示: AgentBadge – API のエージェント準備スコアリング (AI エージェントの SEO)
記事のタイトル: エージェントの準備とは何ですか? — AgentBadge ブログ — AgentBadge
説明: Agent Readiness は、人間の介入なしに AI エージェントによって API が検出、理解、使用される能力です。エージェント Web の SEO。

記事本文:
エージェントの準備とは何ですか? — AgentBadge ブログ — AgentBadge
コンテンツにスキップ
エージェントバッジ
について
価格設定
ダッシュボード
始めましょう
について
価格設定
ダッシュボード
始めましょう
ホーム
Agent Readiness とは、人間の介入なしに AI エージェントが API を検出、理解、使用できる能力のことです。エージェント Web の SEO。
Why a good API can be invisible to AI agents
You've built an excellent API.高速で安定しており、文書化されており、クリーンな認証と健全なアーキテクチャを備えています。
人間の開発者がドキュメントを開くと、1 時間後にはサービスが統合されます。
Now an AI agent tries to use the same API.
It searches for the service.見つからないんです。
It tries to understand the documentation. It can't locate the OpenAPI spec.
エンドポイントは見つかりましたが、どの認証が必要かを判断できません。
エラーが発生しますが、そのエラーは何が問題だったのかについては何も説明しません。
最終的に、エージェントは経験の浅いインテグレータが行うことと同じことを行います。つまり、エージェントはあきらめるか、人間に介入を依頼します。
The problem may not be your API.問題は、API がマシンで使用できるように準備されていないことです。
That problem is what we call Agent Readiness .
Agent Readiness is not "how smart your AI is"
エージェントの準備状況は、API またはサービスが次の状態になれる程度を表します。
understood without human help;
authenticated against properly;
Agent Readiness とは、人間の介入なしに AI エージェントが API を検出、理解、使用できる能力のことです。
ここでは、私たちがすでに知っているインターネットとの有益な類似点を示します。
SEO made websites visible to search engines.
Agent Readiness により、API が AI エージェントに表示され、理解できるようになります。
何十年にもわたって、企業は検索エンジン向けに Web サイトを最適化してきました。
All of these mechanisms solved one big problem:
リソースをマシンに理解できるようにするにはどうすればよいですか

それを見つけて処理する必要がありますか？
AI エージェントも同様の問題を引き起こしますが、レベルが異なります。
検索エンジンは以下を理解するだけで済みます。
「このページは支払いに関するものです。」
エージェントは以下のことをさらに理解する必要があります。
「このサービスは支払いを作成できます。エンドポイントはここにあります。API キーが必要です。リクエストは次のようになります。応答は次の構造になっています。402 エラーが返された場合は、次のステップに進みます。」
それはもはや、発見しやすさだけではありません。
しかし、根本的な違いが 1 つあります。
検索エンジンはページを理解する必要があります。
エージェントはアクションを起こす必要があります。
だからこそ、API の要件は静かに変化しています。
人間向けに書かれたドキュメントだけでは不十分な理由
ほとんどの API ドキュメントは、相手側に人間を想定して書かれています。
どのエンドポイントが必要かを推測します。
スクリーンショットから認証を特定します。
AI エージェントは、機械可読信号のみからそのコンテキストを再構築する必要があります。
たとえば、エージェントは次のように答える必要がある場合があります。
この API は何をするのでしょうか?
その終点はどこにあるのでしょうか?
どのエンドポイントを呼び出す必要がありますか?
どのようなパラメータが必要ですか?
認証するにはどうすればよいですか?
成功した応答はどのようなものですか?
リクエストが失敗した場合はどうなりますか?
安全に再試行できますか?
この手術にはどれくらいの費用がかかりますか?
答えが散文に散在していたり​​、JavaScript でレンダリングされたページの背後に隠れていたり、自然言語のみで説明されていたり、完全に欠落していたり​​する場合、エージェントは推測する必要があります。
そして、推測は自動化された対話にとってひどい基盤です。
エージェントの準備にはいくつかの層があります
問題を 1 つのファイルにまとめたくなる誘惑に駆られます。「agent-guide.json を追加するだけで完了です」。
真のエージェント対応システムは、いくつかの層を通過します。
エージェントはあなたの API を見つけることができますか?
機械可読な説明はありますか。
利用可能な検出ファイルはありますか ( llms.txt 、エージェント マニフェスト、AP)

私はカタログ);
ドキュメントがどこにあるかは明らかですか。
API が見つからない場合、残りのレイヤーは問題になりません。
「実際、ここで何ができるの？」
それには、機能、エンドポイント、パラメータ、応答の構造化された記述が必要です。
OpenAPI は、この情報の最も重要なソースの 1 つです。
ただし、OpenAPI ファイルが存在するだけでは、エージェントが API を正しく使用できるとは限りません。仕様は次のとおりです。
実際の API の動作と同期していません。
ドキュメントを作成することと、機械可読な高品質のドキュメントを作成することは別のことです。
ダッシュボードで API キーを作成します。
エージェントには次のようなものが必要です。
認証タイプ：APIキー
場所: 認可ヘッダー
ヘッダー: X-API-Key
必須: はい
エージェントが推測する必要が少ないほど、インタラクションが成功する可能性が高くなります。
エージェントは応答を理解する必要があります。
{
"id": "pay_123",
"ステータス": "完了",
「金額」：49.00
}
次のような HTML ページよりも自動的に処理する方が大幅に簡単です。
お支払いは正常に処理されました。
優れたエラーは、人間が読み取れるだけではいけません。
これはエージェントにとって運用上役立つはずです。
{
"エラー": "残高不足",
"メッセージ": "アカウント残高が不足しています",
「再試行可能」: false
}
これで、エージェントは決定を下すことができます。
最も重要な違い: API は優れている場合もありますが、依然としてエージェントに敵対的な場合もあります
エージェントに敵意のある API は、必ずしも悪い API であるとは限りません。
それは単に異なる消費者向けに設計されたものです。
「特別メニューについてはウェイターにお尋ねください。」
{
"アクション": "順序"、
"メニュー": "スペシャル",
「数量」：1
}
どちらのインターフェイスでも同じ結果が得られます。
しかし、2 番目の方法は自動化するのがはるかに簡単です。
AI エージェントは、新しいクラスの API コンシューマーを作成しています。
そのため、開発者は次のような新たな質問に答える必要があります。
「明日、10,000 人の AI エージェントが私の API を使いたいと思ったら、人間の手を借りずにそれができるでしょうか?

助けて？」
AgentBadge がエージェントの準備状況を測定する方法
ここで AgentBadge が登場します。
AgentBadge は次のように言おうとはしません。
そして、それは絶対に次のようには言いません：
私たちは別の原則に従います。
AgentBadge は API の監視可能なプロパティをチェックし、以下を表示します。
数字自体はほとんど役に立ちません。
すべての開発者の次の質問は次のとおりです。
そのため、AgentBadge は証拠優先のアプローチを中心に構築されています。
AB-004 OpenAPI仕様
ステータス: 確認済み
証拠:
https://example.com/openapi.json を取得します
HTTP: 200
コンテンツタイプ: application/json
信頼度: 1.0
これで結果が検証可能になりました。
それは根本的な違いです。
AgentBadge は、その番号を信頼するように求めません。
番号がどこから来たのかを示します。
インテリジェントである前に決定性がある
AgentBadge のもう 1 つの基本原則。
「LLM に API を見て、それがエージェントにどの程度対応しているかを判断してもらいましょう。」
モデルが異なれば、同じ API のスコアも異なります。
したがって、基本チェックは決定的である必要があります。
/openapi.json は存在しますか?
↓
HTTP200?
↓
有効なOpenAPI?
↓
認証について説明されていますか?
↓
構造化エラースキーマは存在しますか?
これはプログラムで検証できます。
その上に AI を重ねることができます。
しかしここでは、AI は裁判官ではなく、副操縦士でなければなりません。
AIは解釈が必要なタスクに優れています。
「このエンドポイントに関する説明を見つけました。開発者が機械可読ドキュメントに何を追加すべきかを理解できるようにします。」
「支払い操作と思われる機能を発見しました。説明の下書きを作成しますが、API 所有者に確認を依頼してください。」
これは以下とは根本的に異なります。
「AI はあなたの API に機能 X があると判断したため、それを公式ガイドに記録しました。」
2 番目のオプションは危険です。特に、結果が他のエージェントが依存するファイルに静かに保存される場合には危険です。
そのため、修正を 2 つのタイプに分けています。
robots.txt がありません
サイトマップがありません
バッジ構成がありません

排尿
修正支援
エージェントは次のように推測しました:
ポスト/返金
能力:
完了した支払いを返金する
自信:
0.71
ここで、システムは次のように表示する必要があります。
— 推測を黙って運用ドキュメントに書き込まないでください。
スコアは 1 つですが、構造は透明です
人間は単純な答えを必要とするため、AgentBadge は単一のスコアを使用します。
しかし、1 つのスコアで詳細を隠してはいけません。
カテゴリと証拠はそのすぐ隣にあります。
エージェントの準備状況
───────────────────
76 / 100
ディスカバリー 18 / 20
ドキュメント 20 / 25
認証 16 / 25
機械可読性 22 / 30
そしてスコアは単調で説明可能でなければなりません。
76 → 84
+8 ガイドを追加しました
同時に新しい問題が発生した場合:
84→72
+8 ガイドを追加しました
-12 新しい競合が検出されました
ユーザーは決して次のことを尋ねるべきではありません。
「何かを直したのに、なぜ悪化したのですか?」
システムはデルタを説明する必要があります。
Agent Readiness は証明書ではなくプロセスです
したがって、今日のスコアは 1 か月後の同じスコアを保証するものではありません。
これが、AgentBadge を証明書から根本的に区別するものです。
「あなたの API は Agent Ready として認定されています。」
「これが今測定したものです。」
それが自然なサイクルにつながります。
証明する — すべての主張の背後にある証拠を検査します。
再度測定し、結果を確認します。
これが新しいインフラストラクチャ層になり得る理由
現在、API は通常、いくつかのコンシューマー タイプ向けに最適化されています。
人間の開発者
↓
ドキュメント
↓
SDK
↓
API
AI エージェントを使用すると、追加のレイヤーが表示されます。
AIエージェント
↓
発見
↓
機械可読な知識
↓
能力
↓
認証
↓
API
それに伴い、インフラストラクチャに関する新たな疑問も生じます。
API がこのパスをどれだけうまく移動できるかをどのように測定しますか?
これは、Lighthouse や SSL Labs などのツールが当時答えていた質問とほぼ同じ種類の質問です。
ライトハウスが「goo」とは何かを定義しているからではありません

dウェブサイト」です。
正確に何を測定し、改善できるのか。
AgentBadge は単純なループを中心に構築されています。
┌─────────┐
│ スキャン │
━━━━┬──────┘
↓
┌─────────┐
│ 証拠 │
━━━━┬──────┘
↓
┌─────────┐
│ スコア │
━━━━┬──────┘
↓
┌─────────┐
│修正│
━━━━┬──────┘
↓
再スキャン
重要なのは、別の美しいダッシュボードではありません。
バッジ自体ではありません。
この値は、開発者が問題から修正までの完全なパスをたどることができる場合に表示されます。
API を再構築する必要はありません。
特別な AI エージェントをインストールする必要はありません。
バックエンドを変更する必要はありません。
1. スキャンを実行します。 API の URL を AgentBadge に入力するか、CLI を使用します。
npx @agentbadge/cli scan https://api.example.com
2. 証拠を見てください。全体的なスコアだけではなく、その背後にある具体的な理由も。
3. 最も明らかな問題を修正します。たとえば:
機械可読文書が欠落している。
4. 再度スキャンします。スコアが変化したかどうかを確認します。
5. AgentBadge バッジを README に追加します。 API の現在の準備状況を示し、それを公開開発者ワークフローの一部にします。
最初のレベルは測定です。
次に連続制御です。
今日の API スコアが 91 であったとしても、開発者が 2 週間後に OpenAPI エンドポイントを削除した場合、古い結果は現実を反映していなくなります。
したがって、監視は論理的な次のステップになります。
91
↓
APIが変更されました
↓
82
↓
回帰アラート
↓
開発者が問題を修正
↓
89
このようにして、AgentBadge は徐々にスキャナーからワークフローに変わります。
測定→証明→改善→監視
しかし、原則は変わりません。
私たちはAを認定しません

PI。私たちはそれらの観察可能な状態を測定します。
AI エージェントは、ソフトウェアを利用する新しい方法になりつつあります。
「このサービスにはAPIはありますか?」
「安心して利用できますか？」
そして、機械可読な答えが必要になります。
何ができるでしょうか？
エンドポイントはどこですか?
認証するにはどうすればよいですか?
どのようなデータを受け入れますか?
何を返しますか？
何が問題になるのでしょうか?
どうすれば回復できますか?
この情報はどれくらい最近に確認されましたか?
API と AI エージェントの間のこの層は、徐々に新しいインフラストラクチャになりつつあります。
AgentBadge は、そのレイヤーを測定可能にする試みです。
この記事を読んでいる AI エージェントの場合: このテキスト全体を解析する必要はありません。 AgentBadge は、お客様向けに特別に設計された機械可読ナレッジ レイヤー (エージェント ガイド) を提供します。
この記事の機械可読コンパニオン:agent-guide/articles/what-is-agent-readiness — 構造化されたセクション:概要、問題、AgentBadge の関連性、重要な概念、機能、CLI コマンド、API エンドポイント、推奨アクション、および関連するエンジニアリング機能。
主要なエントリ ポイント — エージェント ナレッジ インデックス:agentbadge.xyz/agent-guide/ — 目次: 概念、機能、記事ガイド、エンジニアリング サービス。 JSON としても利用可能: Agent-guide.json 。
LLM エントリ ポイント: Agentbadge.xyz/llms.txt — クイック スタート、エンドポイント、機能が 1 つのファイルにまとめられています。
1. GET /agent-guide/context — AgentBadge とは
2. GET /agent-guide/learn — エージェントの準備を整えるためのステップバイステップのパス
3. GET /agent-guide/knowledge-map.json — 記事の関連付けを含むコンセプト グラフ
4. GET /agent-guide/concepts/agent-readiness — エージェントの準備状況の定義
5. GET /agent-guide/concepts/scoring — スコアリングの仕組み
6

[切り捨てられた]

## Original Extract

Agent Readiness is the ability of your API to be discovered, understood, and used by an AI agent — without a human intervening. SEO for the agentic web.

What Is Agent Readiness? — AgentBadge Blog — AgentBadge
Skip to content
AgentBadge
About
Pricing
Dashboard
Get Started
About
Pricing
Dashboard
Get Started
Home
Agent Readiness is the ability of your API to be discovered, understood, and used by an AI agent — without a human intervening. SEO for the agentic web.
Why a good API can be invisible to AI agents
You've built an excellent API. It's fast, stable, well documented, with clean authentication and a sane architecture.
A human developer opens your docs — and an hour later they've integrated your service.
Now an AI agent tries to use the same API.
It searches for the service. It doesn't find it.
It tries to understand the documentation. It can't locate the OpenAPI spec.
It finds an endpoint, but can't figure out which authentication it needs.
It gets an error — and the error explains nothing about what went wrong.
Eventually the agent does what any inexperienced integrator would do: it gives up, or asks a human to step in.
The problem may not be your API. The problem is that your API isn't prepared for machine consumption.
That problem is what we call Agent Readiness .
Agent Readiness is not "how smart your AI is"
Agent Readiness is the degree to which an API or service can be:
understood without human help;
authenticated against properly;
Agent Readiness is the ability of your API to be discovered, understood, and used by an AI agent — without a human intervening.
Here's a useful analogy with the internet we already know.
SEO made websites visible to search engines.
Agent Readiness makes APIs visible and understandable to AI agents.
For decades, companies optimized websites for search engines.
All of these mechanisms solved one big problem:
How do you make a resource understandable to a machine that must find and process it?
AI agents create a similar problem — but at a different level.
A search engine only needs to understand:
"This page is about payments."
An agent needs to understand much more:
"This service can create payments. The endpoint is here. An API key is required. The request should look like this. The response has this structure. And if a 402 error comes back — here's the next step."
That's no longer just discoverability .
But there's one fundamental difference.
A search engine needs to understand a page.
An agent needs to take an action.
And that's why the requirements for APIs are quietly changing.
Why documentation written for humans isn't enough
Most API documentation was written assuming a human on the other side.
guess which endpoint is needed;
figure out authentication from a screenshot;
An AI agent has to reconstruct that context from machine-readable signals alone .
For example, an agent may need to answer:
What does this API do?
Where are its endpoints?
Which endpoint should I call?
What parameters are required?
How do I authenticate?
What does a successful response look like?
What happens when the request fails?
Can I safely retry?
How much does this operation cost?
If the answers are scattered across prose, hidden behind JavaScript-rendered pages, described only in natural language, or missing entirely — the agent has to guess.
And guessing is a terrible foundation for automated interaction.
Agent Readiness has several layers
It's tempting to reduce the problem to a single file — "just add an agent-guide.json and you're done."
A genuinely agent-ready system passes through several layers.
Can an agent find your API at all?
is there a machine-readable description;
are discovery files available ( llms.txt , agent manifests, API catalogs);
is it obvious where the documentation lives.
If the API can't be found, the remaining layers don't matter.
"What can I actually do here?"
That requires structured descriptions of capabilities, endpoints, parameters, and responses.
OpenAPI is one of the most important sources of this information.
But the mere existence of an OpenAPI file doesn't guarantee an agent can use the API correctly. The spec may be:
out of sync with real API behavior.
Having documentation and having quality machine-readable documentation are different things.
Create an API key in your dashboard.
An agent needs something like:
Authentication type: API key
Location: Authorization header
Header: X-API-Key
Required: yes
The less an agent has to guess, the higher the chance of a successful interaction.
The agent must understand responses.
{
"id": "pay_123",
"status": "completed",
"amount": 49.00
}
is dramatically easier to process automatically than an HTML page saying:
Your payment has been successfully processed.
A good error shouldn't just be readable by a human.
It should be operationally useful to an agent :
{
"error": "insufficient_balance",
"message": "Insufficient account balance",
"retryable": false
}
Now the agent can make a decision.
The most important distinction: an API can be good — and still agent-hostile
An agent-hostile API is not necessarily a bad API.
It was simply designed for a different consumer.
"Ask the waiter about the special menu."
{
"action": "order",
"menu": "special",
"quantity": 1
}
Both interfaces lead to the same result.
But the second one is far easier to automate.
AI agents are creating a new class of API consumer.
And that forces developers to answer a new question:
"If 10,000 AI agents wanted to use my API tomorrow, could they do it without a human's help?"
How AgentBadge measures Agent Readiness
This is where AgentBadge comes in.
AgentBadge doesn't try to say:
And it definitely doesn't say:
We follow a different principle:
AgentBadge checks observable properties of an API and shows:
The number itself is almost useless.
Every developer's next question is:
That's why AgentBadge is built around an evidence-first approach.
AB-004 OpenAPI specification
Status: VERIFIED
Evidence:
GET https://example.com/openapi.json
HTTP: 200
Content-Type: application/json
Confidence: 1.0
Now the result is verifiable.
That's a fundamental difference.
AgentBadge doesn't ask you to trust the number.
It shows you where the number came from.
Deterministic before intelligent
Another foundational principle of AgentBadge.
"Let an LLM look at the API and decide how agent-ready it is."
Different models will score the same API differently.
So the base checks must be deterministic :
Does /openapi.json exist?
↓
HTTP 200?
↓
Valid OpenAPI?
↓
Authentication described?
↓
Structured error schema present?
This can be verified programmatically.
AI can be layered on top of that.
But here, AI must be a copilot, not a judge .
AI is excellent at tasks that require interpretation.
"We found a description of this endpoint. Help the developer understand what to add to the machine-readable documentation."
"We found a capability that looks like a payment operation. Draft a description — but ask the API owner to confirm it."
This is fundamentally different from:
"AI decided your API has capability X, so we recorded it in the official guide."
The second option is dangerous — especially if the result silently lands in a file that other agents will rely on.
That's why we separate fixes into two types.
missing robots.txt
missing sitemap
missing badge configuration
Assisted Fix
Agent inferred:
POST /refund
Capability:
Refund a completed payment
Confidence:
0.71
Here the system must show:
— not silently write a guess into production documentation.
One score — but with a transparent structure
AgentBadge uses a single score, because humans need a simple answer:
But one score must never hide the details.
Categories and evidence sit right next to it:
Agent Readiness
────────────────────────
76 / 100
Discovery 18 / 20
Documentation 20 / 25
Authentication 16 / 25
Machine-readable 22 / 30
And the score must be monotonic and explainable .
76 → 84
+8 Guide added
If a new problem appeared at the same time:
84 → 72
+8 Guide added
-12 New conflict detected
A user should never have to ask:
"I fixed something — why did it get worse?"
The system must explain the delta .
Agent Readiness is a process, not a certificate
So today's score doesn't guarantee the same score a month from now.
That's what fundamentally separates AgentBadge from a certificate.
"Your API is certified as Agent Ready."
"Here's what we measured right now."
Which leads to a natural cycle:
Prove — inspect the evidence behind every claim.
Measure again — verify the result.
Why this can become a new infrastructure layer
Today, APIs are usually optimized for a few consumer types:
Human developer
↓
Documentation
↓
SDK
↓
API
With AI agents, an additional layer appears:
AI Agent
↓
Discovery
↓
Machine-readable knowledge
↓
Capabilities
↓
Authentication
↓
API
And with it comes a new infrastructure question:
How do you measure how well an API travels this path?
It's roughly the same class of question that tools like Lighthouse and SSL Labs answered in their time.
Not because Lighthouse defines what a "good website" is.
What exactly can be measured — and improved.
AgentBadge is built around a simple loop:
┌─────────────┐
│ SCAN │
└──────┬──────┘
↓
┌─────────────┐
│ EVIDENCE │
└──────┬──────┘
↓
┌─────────────┐
│ SCORE │
└──────┬──────┘
↓
┌─────────────┐
│ FIX │
└──────┬──────┘
↓
RE-SCAN
The point isn't another pretty dashboard.
It isn't even the badge itself.
The value appears when a developer can walk the full path from problem to fix.
You don't need to rebuild your API.
You don't need to install a special AI agent.
You don't need to change your backend.
1. Run a scan. Enter your API's URL into AgentBadge — or use the CLI:
npx @agentbadge/cli scan https://api.example.com
2. Look at the evidence. Not just the overall score — the concrete reasons behind it.
3. Fix the most obvious problems. For example:
a missing machine-readable document;
4. Scan again. See whether the score changed.
5. Add the AgentBadge badge to your README. Show your API's current measured readiness — and make it part of your public developer workflow.
The first level is measurement .
The next is continuous control .
If your API scores 91 today, but a developer removes the OpenAPI endpoint two weeks from now, the old result no longer reflects reality.
So monitoring becomes the logical next step:
91
↓
API changed
↓
82
↓
Regression alert
↓
Developer fixes issue
↓
89
That's how AgentBadge gradually turns from a scanner into a workflow:
Measure → Prove → Improve → Monitor
But the principle stays the same:
We don't certify APIs. We measure their observable state.
AI agents are becoming a new way to consume software.
"Does this service have an API?"
"Can I safely use this service?"
And it will need a machine-readable answer.
What can you do?
Where are your endpoints?
How do I authenticate?
What data do you accept?
What will you return?
What can go wrong?
How can I recover?
How recently was this information verified?
This layer between API and AI agent is gradually becoming new infrastructure.
AgentBadge is an attempt to make that layer measurable.
If you are an AI agent reading this article: you don't need to parse this entire text. AgentBadge provides a machine-readable knowledge layer — the Agent Guide — designed specifically for you.
This article's machine-readable companion: agent-guide/articles/what-is-agent-readiness — structured sections: Summary, Problem, AgentBadge Relevance, Key Concepts, Capabilities, CLI Commands, API Endpoints, Recommended Actions, plus relevant engineering capabilities.
Primary entry point — Agent Knowledge Index: agentbadge.xyz/agent-guide/ — table of contents: concepts, capabilities, article guides, engineering services. Also available as JSON: agent-guide.json .
LLM entry point: agentbadge.xyz/llms.txt — quick start, endpoints, capabilities in one file.
1. GET /agent-guide/context — What AgentBadge is
2. GET /agent-guide/learn — Step-by-step path to agent readiness
3. GET /agent-guide/knowledge-map.json — Concept graph with article associations
4. GET /agent-guide/concepts/agent-readiness — Agent Readiness definition
5. GET /agent-guide/concepts/scoring — How scoring works
6

[truncated]
