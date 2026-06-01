---
source: "https://www.indiehackers.com/post/is-system-prompting-enough-for-production-why-i-m-building-a-runtime-governance-layer-for-ai-agents-2f57424547"
hn_url: "https://news.ycombinator.com/item?id=48345709"
title: "Why I'm building a runtime governance layer for AI agents and Apps"
article_title: "Is “system prompting” enough for production? Why I’m building a runtime governance layer for AI agents. - Indie Hackers"
author: "piyush165"
captured_at: "2026-06-01T00:30:32Z"
capture_tool: "hn-digest"
hn_id: 48345709
score: 2
comments: 0
posted_at: "2026-05-31T13:56:05Z"
tags:
  - hacker-news
  - translated
---

# Why I'm building a runtime governance layer for AI agents and Apps

- HN: [48345709](https://news.ycombinator.com/item?id=48345709)
- Source: [www.indiehackers.com](https://www.indiehackers.com/post/is-system-prompting-enough-for-production-why-i-m-building-a-runtime-governance-layer-for-ai-agents-2f57424547)
- Score: 2
- Comments: 0
- Posted: 2026-05-31T13:56:05Z

## Translation

タイトル: AI エージェントとアプリのランタイム ガバナンス レイヤーを構築している理由
記事のタイトル: 「システム プロンプト」は運用に十分ですか? AI エージェント用のランタイム ガバナンス レイヤーを構築している理由。 - インディーハッカー
説明: 私はここ数か月間、ある認識に夢中になって過ごしてきました。それは、システム プロンプトだけでは、本番 AI SaaS の基盤としては弱いということです。それらは役に立ちます。彼らは...

記事本文:
「システムプロンプト」は運用には十分ですか? AI エージェント用のランタイム ガバナンス レイヤーを構築している理由。 - インディーハッカー
ホーム
起動中
事例DB
製品
アイデアDB
Vibe コーディング ツール
IH+を購読する
起動中
ケーススタディ
アイデアDB
製品DB
参加する
1
いいね
1
ブックマーク
1
コメント
レポート
「システムプロンプト」は運用には十分ですか? AI エージェント用のランタイム ガバナンス レイヤーを構築している理由。
私はここ数か月間、あることに夢中になって過ごしてきました。
システム プロンプトだけでは、本番 AI SaaS の基盤としては弱いです。
それらは役に立ちます。それらは行動を導きます。デモを印象的に見せます。
しかし、AI 製品が実際のユーザー、乱雑なコンテキスト、ビジネス ルール、顧客からの圧力、価格設定ロジック、メモリ、ツール、エッジ ケースに関わると、システム プロンプトが過大な責任を負い始めます。
それが私がNEES Core Engineを構築している理由です。
私は、Indie Hackers コミュニティでの正直な質問を検証しようとしています。
AI 創設者はすでに本番環境でこの苦痛を感じているのでしょうか、それとも私はまだ時期尚早なのでしょうか?
デモでは、AI エージェントは完璧に見えます。
実稼働環境では、意図した製品ロジック、ビジネス境界、トーン、安全ルール、ユニットエコノミクスから徐々に離れていく可能性があります。
必ずしも劇的な失敗をするわけではありません。
場合によっては、その答えが合理的に聞こえることもあります。
しかし、その答えの裏には、AI がポリシーを無視したり、ワークフローのステップをスキップしたり、間違ったコンテキストを使用したり、チームの誰も明確に追跡できない決定を下したりしたことが考えられます。
ポリシーバイパス
サポート ボットは礼儀正しく「安全」に聞こえますが、ユーザーを満足させるためだけに会社のポリシーや規約を無視します。
幻覚の価格設定
営業アシスタントは、企業によって承認されていない割引、返金、または約束を提案します。
コンテキストの混乱
CRM アシスタントは、乱雑、フィルタリングされていない、または古いユーザー履歴に基づいて口調や動作を変更します。
ブラ

CKボックスの問題
モデルは決定を下しますが、チームはなぜその決定が許可されたのかを説明できません。
LLM税
この製品は、決定的に管理、再利用、キャッシュ、または処理されるべき回答を求める繰り返しのモデル呼び出しに対して料金を支払い続けます。
これを私はエージェント・ドリフトと呼んでいます。
エージェントはまだ「機能」している可能性がありますが、製品の意図した動作から徐々に離れていきます。
私の持論: プロンプトは創造性を促すものです。ガバナンスは信頼性のためにあります。
ほとんどのビルダーは、より長いプロンプト、より多くの命令、再帰的チェック、出力フィルター、または単純なガードレールを追加することで、この問題を解決しようとします。
しかし、動作、ポリシー、メモリ、コスト、トレーサビリティが重要となる本番 AI システムには十分ではないと思います。
私は、本番環境の AI にはランタイム ガバナンスが必要だと考えています。
私が NEES で構築している基本的なフローは次のとおりです。
アプリ → NEES ガバナンス ランタイム → モデル プロバイダー → ガバナンスされた応答
NEES は、すべての動作責任をソフト プロンプト内に置くのではなく、アプリケーションとモデルの間にガバナンス層を追加します。
目標は、OpenAI、Anthropic、Google、LangChain、CrewAI、またはその他のフレームワークを置き換えることではありません。
目標は、応答やアクションがユーザーに届く前に、AI の動作をより製品に合わせたものにすることです。
NEES は次のようなことを中心に設計されています。
実行前の意図チェック
トークンを消費したり、ワークフロー パスを許可したりする前に、ユーザーが何をしようとしているのかを理解します。
ポリシーの施行
プロンプトの指示だけに頼るのではなく、製品固有のルールに照らしてモデルの動作をチェックします。
記憶の境界
AI がインタラクション間で何を記憶、使用、または転送できるかを制御します。
追跡可能な意思決定
応答またはアクションが許可、ブロック、エスカレーション、または変更された理由を記録します。
エスカレーションロジック
AI が直接応答すべきではなく、引き継ぎ、明確にする、または停止する必要がある場合を把握します。
コストゴー

ルナンス
安全な決定論的なパス、キャッシュされた応答、または再利用可能な管理された応答で十分な場合、不必要なモデル呼び出しを回避します。
フォールバック動作
モデル プロバイダーに障害が発生したり、遅延が急増したり、低コスト/ローカル ルートが発生したりした場合でも、製品を安定に保つことがより適切です。
私が探しているのは顧客ではなく、デザインパートナーです。
今のところ、広範なマーケティングに関するフィードバックは求めていません。
私は以下を構築する創設者や開発者からの正直なシグナルを求めています。
ツールを使用したエージェント製品
「システム プロンプトの壁」にもうぶつかっていませんか?
一貫性のない動作、トレーサビリティの欠如、メモリの問題、繰り返される LLM コスト、またはより強力なビジネス ルール制御を必要とする AI アクションなどに悩まされていませんか?
それとも、プロンプト、ガードレール、カスタム チェックは、2026 年の本番 AI の状況にまだ十分だと思いますか?
1 つの小さなワークフローでこれをテストしたいと考えている 2 ～ 3 人の創業者と話をしたいと考えています。
私は、実際のワークフローを NEES スタイルのガバナンス構造にマッピングすることを個人的に支援し、ランタイム ガバナンスが実際の製品環境でエージェント ドリフトを削減できるかどうかを確認したいと考えています。
開発者プレビュー:
https://github.com/NEES-Anna/nees-core-developer-preview
ライブサンプルアプリ:
https://naina.nees.cloud
AI ビルダーからの率直なフィードバックをお待ちしています:
ランタイム ガバナンスは本番 AI にとって実際に欠けている層になりつつあるのでしょうか、それとも市場はまだ時期尚早なのでしょうか?
アンナ2612
に投稿されました
開発者
2026 年 5 月 26 日
1
1
シェアする
Anna2612 に何か良いことを言ってください…
コメントを投稿する
1
文脈上、私は NEES をチャットボットや別のモデル ラッパーとして位置づけているわけではありません。
私は、特に AI の動作が追跡可能で、ポリシーを認識し、ビジネス ロジックと連携する必要がある場合に、「ランタイム ガバナンス」が本番 AI アプリに欠けているインフラストラクチャ層になるかどうかを調査しようとしています。
建築業者からの正直な批判を歓迎します

再。
53件のコメント
実際の製品に取り組む代わりに、壊れたスクレーパーを修理しました。それで私はそれを自分の問題にしました。
44件のコメント
収益の問題が悪化する前にそれを確認する方法
29件のコメント
PM として破産して燃え尽きた状態から、SaaS を立ち上げて健康状態を最適化するまで
27件のコメント
私はプロジェクトを立ち上げたり中止したりを繰り返しました。だから私はそれを許さないシステムを構築しました
23件のコメント
Shopify テーマを月額 20,000 ドルかけて構築しました。次に、ピボットする必要があります。
20件のコメント
インディーハッカーとして常に最新情報を入手してください。
ビジネスの開始と成長に役立つ市場洞察。
購読する
フォローする
@IndieHackers
X では、収益性の高いオンライン ビジネスを構築する創業者に関するストーリーや洞察を得ることができ、また、インディー ハッカー コミュニティの他の人とつながることができます。

## Original Extract

I’ve spent the last few months obsessed with one realization: System prompts alone are a weak foundation for production AI SaaS. They are useful. They g...

Is “system prompting” enough for production? Why I’m building a runtime governance layer for AI agents. - Indie Hackers
Home
Starting Up
Case Studies DB
Products
Ideas DB
Vibe Coding Tools
Subscribe to IH+
Starting Up
Case Studies
Ideas DB
Products DB
Join
1
Like
1
Bookmark
1
Comment
Report
Is “system prompting” enough for production? Why I’m building a runtime governance layer for AI agents.
I’ve spent the last few months obsessed with one realization:
System prompts alone are a weak foundation for production AI SaaS.
They are useful. They guide behavior. They make demos look impressive.
But once an AI product touches real users, messy context, business rules, customer pressure, pricing logic, memory, tools, and edge cases — the system prompt starts carrying too much responsibility.
That is why I’m building NEES Core Engine.
I’m trying to validate one honest question with the Indie Hackers community:
Are AI founders already feeling this pain in production, or am I still too early?
In a demo, your AI agent can look perfect.
In production, it can slowly drift away from your intended product logic, business boundaries, tone, safety rules, and unit economics.
It is not always a dramatic failure.
Sometimes the answer sounds reasonable.
But behind that answer, the AI has ignored a policy, skipped a workflow step, used the wrong context, or made a decision nobody on the team can clearly trace.
Policy bypass
A support bot sounds polite and “safe,” but ignores company policy or terms just to satisfy the user.
Pricing hallucination
A sales assistant offers a discount, refund, or promise that was never approved by the business.
Context chaos
A CRM assistant changes tone or behavior based on messy, unfiltered, or outdated user history.
The black box problem
The model makes a decision, but the team cannot explain why that decision was allowed.
The LLM tax
The product keeps paying for repeated model calls for answers that should have been governed, reused, cached, or handled deterministically.
This is what I call Agent Drift.
The agent may still “work,” but it slowly moves away from the product’s intended behavior.
My thesis: prompts are for creativity; governance is for reliability.
Most builders try to fix this by adding longer prompts, more instructions, recursive checks, output filters, or simple guardrails.
But I don’t think it is enough for production AI systems where behavior, policy, memory, cost, and traceability matter.
I believe production AI needs runtime governance.
The basic flow I’m building with NEES is:
App → NEES Governance Runtime → Model Provider → Governed Response
Instead of putting all behavioral responsibility inside a soft prompt, NEES adds a governance layer between the application and the model.
The goal is not to replace OpenAI, Anthropic, Google, LangChain, CrewAI, or any framework.
The goal is to make AI behavior more product-aligned before the response or action reaches the user.
NEES is designed around things like:
Pre-execution intent checks
Understanding what the user is trying to do before spending tokens or allowing a workflow path.
Policy enforcement
Checking model behavior against product-specific rules instead of relying only on prompt instructions.
Memory boundaries
Controlling what the AI can remember, use, or carry forward across interactions.
Traceable decisions
Recording why a response or action was allowed, blocked, escalated, or modified.
Escalation logic
Knowing when the AI should not answer directly and should hand off, clarify, or stop.
Cost governance
Avoiding unnecessary model calls when a safe deterministic path, cached answer, or reusable governed response is enough.
Fallback behavior
Keeping the product stable when the model provider fails, latency spikes, or a lower-cost/local route is more appropriate.
I’m looking for design partners, not customers.
I’m not looking for broad marketing feedback right now.
I’m looking for honest signal from founders and developers building:
agentic products with tool use
Have you hit the “system prompt wall” yet?
Are you struggling with inconsistent behavior, lack of traceability, memory concerns, repeated LLM cost, or AI actions that need stronger business-rule control?
Or do you feel that prompts, guardrails, and custom checks are still good enough for where production AI is in 2026?
I’m looking to talk to 2–3 founders who are willing to test this on one small workflow.
I want to personally help map one real workflow into a NEES-style governance structure and see whether runtime governance can reduce Agent Drift in a practical product environment.
Developer Preview:
https://github.com/NEES-Anna/nees-core-developer-preview
Live Sample App:
https://naina.nees.cloud
Would love honest feedback from AI builders here:
Is runtime governance becoming a real missing layer for production AI, or is the market still too early?
Anna2612
posted to
Developers
on May 26, 2026
1
1
Share
Say something nice to Anna2612…
Post Comment
1
For context, I’m not positioning NEES as a chatbot or another model wrapper.
I’m trying to explore whether “runtime governance” becomes a missing infrastructure layer for production AI apps — especially where AI behavior needs to be traceable, policy-aware, and aligned with business logic.
Would love honest criticism from builders here.
53 comments
Fixing broken scrapers instead of working on my actual product. So I made it my problem.
44 comments
How to see revenue problems before they get worse
29 comments
From broke and burned out as a PM, to launching my SaaS and optimizing my health
27 comments
I kept starting projects and dropping them. So I built a system that wouldn’t let me
23 comments
We built Shopify themes to $20k/month. Now we have to pivot.
20 comments
Stay informed as an indie hacker.
Market insights that help you start and grow your business.
Subscribe
Follow
@IndieHackers
on X for stories and insights about founders building profitable online businesses, and to connect with others in the Indie Hackers community.
