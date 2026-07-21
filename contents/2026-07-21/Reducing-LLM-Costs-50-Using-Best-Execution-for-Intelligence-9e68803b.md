---
source: "https://www.thesean.ai/blog/introducing-ship"
hn_url: "https://news.ycombinator.com/item?id=48998844"
title: "Reducing LLM Costs 50% Using Best-Execution for Intelligence"
article_title: "Introducing Ship Beta"
author: "arbayi"
captured_at: "2026-07-21T21:55:46Z"
capture_tool: "hn-digest"
hn_id: 48998844
score: 1
comments: 0
posted_at: "2026-07-21T21:50:13Z"
tags:
  - hacker-news
  - translated
---

# Reducing LLM Costs 50% Using Best-Execution for Intelligence

- HN: [48998844](https://news.ycombinator.com/item?id=48998844)
- Source: [www.thesean.ai](https://www.thesean.ai/blog/introducing-ship)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T21:50:13Z

## Translation

タイトル: インテリジェンスの最適実行を使用して LLM コストを 50% 削減
記事のタイトル: Ship Beta のご紹介
説明: インテリジェンスの最適実行を使用して LLM コストを 50% 削減

記事本文:
Ship ベータ版の紹介
ホーム アップデート Contact Docs アップデート Contact Docs ダッシュボード Ship Beta の紹介
リサーチ ニュース 企業、インテリジェンスの最良執行を使用して LLM コストを 50% 削減 2026 年 7 月 21 日 著者数名
Ship を試す CONTENTS Ship は、既に使用しているモデルと区別できない出力を半額で提供するエンドポイントです。
model="<original>" を model="ship-like/<original>" に置き換えるだけで、すべてのリクエストのコストが 50% 安くなります。当社は、個々のリクエストの実行により多くの費用がかかるリスクを負います。これにより、(A) 既存のワークフローのコストを削減するか、(B) 同じ予算でより多くのインテリジェンスを使用して品質を向上させることができます。
これは、機能の同等性と動作の同等性という 2 つの保証を備えた推論時間の最適化を行うことで実現されます。機能の同等性とは、元のモデルによって解決される問題はすべて Ship によっても解決されることを意味します。動作が同等であるということは、 Ship の使用を開始するときにプロンプ​​トを変更する必要がないことを意味します。同じ動作が得られます (出力の形状、ツール呼び出しの平均数、従う命令など、機能の外にあるものはすべて同じままです)。
基準相対品質を定義して測定することで、モデルの品質を契約上のサービス レベルに変えることができます。 Ship は、価格と可用性の約束とともに高品質の SLA を提供する最初のエンドポイントです。
船舶が能力の同等性と動作の同等性を維持しているという証拠
Ship を支えるテクノロジー: 推論時間の最適化のスケーリング
人々がShipで構築したもの
同等のパフォーマンスをわずかなコストで実現
システムパフォーマンス: 機能の等価性と動作の等価性
私たちは 2 つの主張を行い、それぞれを直接測定します。
機能の等価性: 参照モデルに問題がある場合 (

f ) が解決され、Ship も解決されます。形式的には、1 - P(Ship が x を解決する | ref が x を解決する) < ε です。
動作の等価性 : 同じ入力を条件として、 Ship の出力は、参照モデルと同じ動作の分布 (同じ出力形状、ツール使用パターン、フォーマット、拒否など) から抽出されます。問題も同様にうまく解決できるはずです。形式的には、任意の距離メトリックと動作に対して、 distance(P_D(x, Behavior(Ship(x)))), P_D(x, Behavior(ref(x)))) < ε となります。
これは、Ship が文字通り同一の出力を提供するという意味ではありません。同じ言語モデルであっても、その定義ではそれ自体が同一ではありません。参照モデルを 2 回呼び出すと、異なるトークン、異なるツール呼び出し数、異なる表現が得られます。代わりに、能力と動作の等価性は分布に関する主張です。同じ入力を条件として、Ship の出力は参照モデルと十分に類似した動作の分布から引き出されます。
これらを測定するには、次の 2 セットの評価を使用します。
簡単に再現できる公開ベンチマーク
他社が当社のために実行するプライベートベンチマーク
公開ベンチマークにより、誰でも結果を検証できます。プライベート ベンチマークは、実際のアプリケーションを実行するサードパーティから提供されているため、オーバーフィットしていないことを示しています。これは、Ship が、実際の使用から構築された、トレーニングされていないディストリビューションに一般化していることを意味します。このブログ投稿では、公開ベンチマークについて詳しく説明します。当社のパートナーがプライベート ベンチマークを実行している場合は、それについても報告します。
能力の等価性では、参照モデルが問題を解決できる場合でも、Ship は問題を解決できるのかという 1 つの質問が生じます。
多数のベンチマークにわたって、Ship はリファレンス モデルと同じパフォーマンスを維持しています。
これはプライベートベンチマークにも当てはまります (「What People Have Built with Ship」を参照)。
もっと見る
キーパー

能力の同等性とは、同じ種類の問題を解決することです。ここ。 Ship の全体的なパフォーマンスが基本モデルと同等であるだけでなく、参照モデルが解決する特定の問題も Ship が解決していることがわかります。
タスクは GPT-5.6 Sol の解決率によってバンドにグループ化されました (深いバンド = ベースで解決されました)。各バンド内では、すべての列が独自の解決率に従って並べ替えられます。行は列全体でタスクに合わせて配置されていません。各列の下: 全体的な解決率と基本モデルとの相関関係。
同じ問題を解決するだけでなく、プロンプトやハーネスを変更せずに問題を解決したいと考えています。これを行動の等価性と呼びます。これは 2 つの理由から重要です。(A) 過剰適合を防ぐこと、つまりシステムは、ベンチマークが捉えられないほど悪い動作をしながらも、従来のベンチマークを達成できること、そして (B) 新しいシステムを使用するために過度のエンジニアリング努力を投資する必要がないことを意味します。 Ship は、ユーザーのスイッチング コストを低く抑えるために、動作の等価性を実現するために最適化されました。
現在、私たちの動作等価性の定義はおそらく強すぎるでしょう。おそらく、運用環境では重要ではなく、エンドユーザー エクスペリエンスに影響を与えないプロパティを保存しているのです。たとえば、現在の測定値には出力長が含まれています。これが実際の使用にとって重要であるかどうかは明らかではなく、多くの場合、コストを節約するために生成するトークンの数を減らした方が良い可能性があります。将来的には、重要な特性だけを捉えるように定義を改良する予定です。今日、私たちは測定したすべての特性を可能な限り一致させることを試みました。
もっと見る
また、特定の意思決定に関連する行動軸がどのように一致するかについても報告します。以下はスコアカードです。
タスクごとのセッションの長さ (ステップ)
エージェント実行秒数 · 平均 + ベンチマークごと
モデル生成とツール実行、タスクあたりの秒数
平均c

ベンチマーク全体、ツールごとのタスクごとのすべて
リアル トレース エンベディング (OpenAI text-embedding-3-small )、タスク中心、その後 t-SNE から 2D へ。船の痕跡は基本モデルの雲に着陸します。 GPT-5.5は離れたままです。
システムの説明: スケーリング推論時間の最適化
能力と行動の同等性をどのように達成するのでしょうか?単純なアプローチ (常に 1 つのモデルを使用するか、単純なルーターを追加する) は機能しません。同じプロンプトが与えられた場合、モデルが異なると動作も異なり、機能も異なります。これらのアプローチでは、キャッシュを意識した意思決定を行うために必要な最適化も欠落し、複雑な環境でのエージェント タスクの処理も欠落し、機能と動作の同等性を促進する可能性のあるその他の最適化も欠落します。代わりに、Ship は推論時に各リクエストの実行を最適化します。
推論の最適化は、コンパイラのジャストインタイム コンパイルに似ています。従来のコンパイラ (従来のトレーニングなど) は、実行時に特定のリクエストを観察していないため、一般的なケースに合わせて最適化します。推論時の追加情報により、より効果的な最適化が可能になります。
通常、LLM はトレーニング後にフリーズされます。フリーズされたモデルは、その実行のすべての部分を目の前の特定のリクエストに適応させることはできません。コンピューティングの割り当てと利用可能な動作は、トレーニング、デコード設定、およびプロンプトによって大部分が修正されます。
プログラムのコンパイルに例えると分かりやすいでしょう。従来の事前コンパイラは、実際のランタイム入力を確認する前に、一般的なバイナリを生成する必要があります。ジャストインタイム (JIT) コンパイラーは、これらの入力を監視した後で実行を改善し、ホット パスを特化し、価値のある最適化の労力を費やすことができます。重要な考え方はコンパイルそのものではありません。それは、ランタイム情報によって、以前は利用できなかった最適化が可能になるということです。
モデルのトレーニングはアナログです

事前コンパイル: モデルがリクエストを認識する前に、フリーズされたモデルが生成されます。 Ship は、ジャストインタイム コンパイルに似ています。推論時に、Ship はどのくらいのコンピューティングを費やすか、どのメソッドがリクエストを処理するかを決定します。メソッドは、モデル、ハーネスとツールを備えたモデル、カスケード、アンサンブル、単純なプログラム、重み空間または活性化空間で変更されたモデル、または上記の組み合わせにすることができます。 JIT システムと同様に、Ship はトレーニング中に利用できない実行時情報、つまりリクエスト自体と、同様の実際のトラフィックから検証された結果を活用します。
一連のモデルとそれらのモデルに介入する方法が与えられると、リクエストを満たすための方法は組み合わせて多数存在します。これは、チェスや囲碁に似た検索問題になります。私たちの目標は、AlphaGo または AlphaZero と同等のものを構築することです。これは最良の実行です。モデルのセットを考慮して、可能な限り最善の方法でリクエストを実行します。
だからこそ、Ship は機能と動作の同等性を維持しながらコストを最大 50% 削減でき、各リクエストの実行方法を最適化します。私たちは、誰もが 1 ドルあたり最高のインテリジェンスを手に入れることができる世界を目指して構築しています。
人々が船で築いたもの
推論コストを半分に削減することは、単なる項目ではありません。それは経済的に建設可能なものを変えます。 Ship はドロップイン代替品であるため、チームはプロンプトや評価に手を加えることなく節約効果を確認できました。彼らが彼らに対して何をしたかは次のとおりです。
Ari Messer 氏、Kilo Code のパートナーシップ責任者
私たちは、世界中の Kilo Code ユーザーから処理された何兆ものトークンを使用して、Ship をステルス モデルとして実行してきました。このトラフィックを越えて、半分のコストで魔法のように機能しました。最高級のインテリジェンスを劇的に手頃な価格にすることで、Ship はすぐに開発にとって欠かせない日常の推進力となりました。

コード生成から複雑なエージェント ワークフローに至るまで、あらゆるものを強化するユーザーとナレッジ ワーカー。ステルス状態から抜け出すのを見るのは興奮します!
何百万もの PR を見てきた私たちは、システムがこれまでに遭遇した最も困難な 168 個のバグのベンチマークを構築しました。これは、ハーネスを改善し、生産に投入するモデルを選択するために社内で使用するベンチマークです。
Ship での参照モデルとの交換は非常に簡単で、base_url とモデル文字列を変更するだけでした。ベンチマークではパフォーマンスに統計的な違いは見られませんでした。これを実稼働環境に導入したところ、ユーザーの満足度と、測定したすべてのエンドユーザー指標は安定していました。
コスト削減により、そうでなければ経済的ではなかった新機能を出荷できるようになりました。私たちはすでに既存の機能セットに何百万ドルも費やしており、使用量は 2 倍以上に増加しました。 Ship を使用すると、プラン レビュー、セキュリティ レビュー、マージ レビューなどの新機能をサポートできます。
従来、モデル API はインフラストラクチャ (稼働時間、レート制限、レイテンシ) を保証していました。提供されるインテリジェンスが特定のレベルに維持されることを保証するものではありません。
Ship はリファレンス モデルを品質グレードとして扱うことでそれを可能にします。 ship-like/<original> をリクエストするときは、特定の実装を実行するように求めているわけではありません。あなたは、定義された統計的許容範囲内で参照モデルの機能と動作を提供することを私たちに求めています。
そのグレードには 2 つの要件があります。1 つは、船舶が参照モデルの機能を保持しなければならないこと、もう 1 つはその動作が一定の範囲内に留まること（参照モデル独自の実行ごとの変動など）です。
品質にグレードがあると、最良の実行が明確に定義されます。つまり、それを満たす最も安価な実行 (または最高の品質) を見つけます。 Ship はリクエストの履行方法を変更できますが、黙ってグレードを下げることはできません

。私たちの公開製品の場合、これは SLO (私たちが測定し、それに基づいて構築するもの) です。
当社の高品質 SLA は、その評価をコミットメントに変えます。私たちは実行リスク、コストリスク、最適化によって品質が低下するリスクを負います。 SLA を持つ企業の顧客は、指定された価格で指定されたグレードのインテリジェンスを受け取ります。
Ship は、ベストエグゼキューションを実現した最初の製品であり、1 ドルあたり最高のインテリジェンスを提供するエンドポイントです。推論時間の最適化をスケーリングすることで、インテリジェンスのパレート曲線をシフトできます。以前に使用していたモデルに関係なく、コストが 50% 削減されます。
私たちは、このようなエンドポイントをさらにリリースし、基礎となる研究についてさらに共有できることを楽しみにしています。
最初のリクエストを交換します。model="<original>" を model="ship-like/<original>" に置き換えます。ドキュメントを参照してください。
高品質の SLA または強化されたレート制限に興味がありますか?私たちのチームにご相談ください。
まだサポートしていない参照モデルの Ship をご希望ですか? それともテストするワークロードがありますか?連絡してください。
© Thesean 2026.全著作権所有。
利用規約 / プライバシーポリシー
Thomas Weibel 著の「The Antikythera Mechanism」は CC BY 4.0 の下でライセンスされています。この作品は原作を改変したものです。

## Original Extract

Reducing LLM Costs 50% Using Best-Execution for Intelligence

Introducing Ship Beta
Home Updates Contact Docs Updates Contact Docs Dashboard Introducing Ship Beta
Research News Company Reducing LLM Costs 50% Using Best-Execution for Intelligence July 21, 2026 Several Authors
Try Ship CONTENTS Ship is an endpoint that provides output indistinguishable from the model you're already using, at half the price.
Just replace model="<original>" with model="ship-like/<original>" , and all requests will be 50% less expensive. We take the risk that any individual request costs us more to execute. This allows you to either (A) reduce costs on existing workflows or (B) increase quality by using more intelligence with the same budget.
We achieve this by doing inference-time optimization with two guarantees: capability equivalence and behavioral equivalence. Capability equivalence means that any problem solved by your original model will also be solved by Ship . Behavioral equivalence means that you don't need to change the prompt when you start using Ship ; you'll get the same behavior (everything outside the capabilities, e.g. the shape of the output, average number of tool calls, instruction following, etc. remains the same).
By defining and measuring reference-relative quality, we can turn model quality into a contractual service level. Ship is the first endpoint to offer a quality SLA alongside its price and availability commitments.
Evidence that Ship maintains capability equivalence and behavioral equivalence
The tech powering Ship : scaling Inference-Time Optimization
What people have built with Ship
Equivalent Performance, at a fraction of the cost
System Performance: Capability Equivalence & Behavioral Equivalence
We make two claims, and measure each directly:
Capability equivalence : any problem your reference model ( ref ) solves, Ship also solves. Formally, 1 - P(Ship solves x | ref solves x) < ε .
Behavioral equivalence : conditioned on the same input, Ship 's outputs are drawn from the same distribution of behaviors as the reference model's — same output shapes, tool-use patterns, formatting, refusals, and so on. It should solve the problems equally well. Formally, distance(P_D(x, behavior(Ship(x))), P_D(x, behavior(ref(x)))) < ε for any given distance metric and behavior.
This doesn't mean Ship gives literally identical outputs. Even the same language model isn't identical to itself under that definition: call the reference model twice and you get different tokens, a different number of tool calls, different phrasing. Instead, capability and behavioral equivalence is a claim about distributions : conditioned on the same input, Ship 's outputs are drawn from a sufficiently similar distribution of behaviors as the reference model's.
To measure these, we use two sets of evaluations:
Public benchmarks, which are easily reproducible
Private benchmarks, which other companies run for us
The public benchmarks let anyone verify our results. The private benchmarks show that we haven't overfit, because they come from third parties who run real applications. This means Ship generalizes to distributions it was never trained on, constructed from real-world use. In this blog post, we detail the public benchmarks. As our partners run private benchmarks, we’ll report them as well.
Capability equivalence asks a single question: when your reference model can solve a problem, does Ship still solve it?
Across a large number of benchmarks, Ship maintains the same performance as the reference model.
This also holds true for private benchmarks (see What People Have Built with Ship ).
More
A key part of capability equivalence is solving the same kinds of problems. Here. we can see that not only is the overall performance of Ship comparable to the base model, but that Ship solves the specific problems that the reference model solves.
Tasks grouped into bands by GPT-5.6 Sol's solve rate (deep band = base solved it); within each band every column is sorted by its own solve rate. Rows are not task-aligned across columns. Under each column: its overall solve rate and correlation r with the base model.
In addition to solving the same problem, you want the problem to be solved without having to modify your prompts or harness. We call this behavioral equivalence. It matters for two reasons: (A) it guards against overfitting — a system can ace traditional benchmarks while behaving worse in ways the benchmark doesn't capture — and (B) it means you don't have to invest excessive engineering effort to use the new system. Ship was optimized for behavioral equivalence to maintain a low switching cost for users.
Today, our definition of behavioral equivalence is arguably too strong: we're probably preserving properties that don't matter in production and don't affect the end-user experience. For example, our current measurements include output length. It's not clear this matters for practical use, and in many cases it might be better to generate fewer tokens to save cost. In the future we'll refine the definition to capture just the characteristics that matter; today, we've tried to match as closely as possible every property we measured.
More
We also report how specific decision-relevant behavioral axes line up. Below is a scorecard.
Session length (steps) per task
Agent-execution seconds · Average + per benchmark
Model generation vs tool execution, seconds per task
Mean calls per task by tool, across benchmarks
Real trace embeddings (OpenAI text-embedding-3-small ), task-centered then t-SNE to 2D. Ship's traces land in the base model's cloud; GPT-5.5 stays apart.
System Description: Scaling Inference-Time Optimization
How do we achieve capability and behavioral equivalence? The naïve approaches (always using one model or adding a simple router) don't work. Different models behave differently when given the same prompt and have different capabilities. Those approaches also miss optimizations required to make cache-aware decisions, deal with agentic tasks in complex environments, and miss other optimizations which can drive capability and behavioral equivalence. Instead, Ship optimizes the execution of each request at inference time.
Inference optimization is analogous to just-in-time compilation in compilers. Traditional compilers (like traditional training) optimize for the general case, because they have not observed the specific request at run-time. The additional information from inference-time allows for more effective optimization.
LLMs are typically frozen after training. A frozen model cannot adapt every part of its execution to the specific request in front of it: the compute allocation and available behavior are largely fixed by training, decoding settings, and the prompt.
An analogy to program compilation is useful. A traditional ahead-of-time compiler must produce a general binary before it sees real runtime inputs. A just-in-time (JIT) compiler can improve execution after observing those inputs, specializing hot paths and spending optimization effort where it pays. The important idea is not compilation itself; it is that runtime information enables optimizations that were unavailable ahead of time.
Model training is analogous to ahead-of-time compilation: it produces a frozen model before that model has seen your request. Ship is analogous to just-in-time compilation. At inference time, Ship determines how much compute to spend and which method should handle the request. A method can be a model, a model with a harness and tools, a cascade, an ensemble, a plain program, a model modified in weight-space or activation-space, or a combination of the above. Like a JIT system, Ship benefits from runtime information unavailable during training: the request itself, together with verified outcomes from similar real traffic.
Given a set of models and ways of intervening on those models, there is a combinatorial number of methods for fulfilling a request. This makes it a search problem, analogous to Chess or Go; our aim is to build the equivalent of AlphaGo or AlphaZero. This is best execution: fulfilling requests in the best possible way given a set of models.
That's why Ship can cut cost by ~50% while maintaining capability and behavioral equivalence: it optimizes how each request is executed. We’re building towards a world where everyone gets the best possible intelligence per dollar.
What People Have Built with Ship
Cutting inference cost in half isn't just a line item. It changes what's economically possible to build. Because Ship is a drop-in replacement, teams saw the savings without touching their prompts or their evals. Here's what they did with them.
Ari Messer, Head of Partnerships @ Kilo Code
We’ve run Ship as a stealth model with trillions of tokens processed from Kilo Code users all over the world. Across that traffic, it worked like a charm at half the cost. By making top-tier intelligence dramatically more affordable, Ship quickly became an indispensable daily driver for developers and knowledge workers, powering everything from code generation to complex agentic workflows. Excited to see it move out of stealth!
Having seen millions of PRs, we’ve built a benchmark of the 168 hardest bugs our systems have encountered. This is the benchmark we use internally to improve our harness and to choose what models to put in production.
Swapping in Ship for our reference model was very easy, just changing a base_url and model string. We saw no statistical difference in performance on our benchmark. When we put it in production, user satisfaction and all of the end-user metrics we measure stayed steady.
The cost savings have let us ship new features that otherwise wouldn’t have been economical. We were already spending millions of dollars on our existing feature set, and our usage has more than doubled. Ship lets us support new features like plan review, security review, and merge review.
Model APIs traditionally guarantee infrastructure: uptime, rate limits, and latency. They do not guarantee that the intelligence delivered remains at a particular level.
Ship makes that possible by treating the reference model as a quality grade. When you request ship-like/<original> , you are not asking us to run a particular implementation. You are asking us to deliver the reference model’s capabilities and behavior within defined statistical tolerances.
That grade has two requirements: Ship must retain the reference model’s capabilities, and its behavior must remain within some bounds (e.g. the reference model’s own run-to-run variation).
Once quality has a grade, best execution becomes well-defined: find the least expensive execution that satisfies it (or the highest quality). Ship can change how a request is fulfilled, but it cannot silently lower the grade. For our public product, this is an SLO (what we measure and build around).
Our quality SLA turns that measurement into a commitment. We take the execution risk, the cost risk, and the risk that optimization degrades quality. Enterprise customers with an SLA receive a specified grade of intelligence at a specified price.
Ship is our first product for best-execution: an endpoint that delivers the highest intelligence per dollar. By scaling inference-time optimization, it can shift the Pareto curve for intelligence: 50% lower cost regardless of the model you were using before.
We're excited to release more endpoints like this, and to share more about the underlying research.
Swap your first request: replace model="<original>" with model="ship-like/<original>" . See the docs .
Interested in a quality SLA or enhanced rate limits? Talk to our team.
Want Ship for a reference model we don't yet support, or have a workload to test? Get in touch .
© Thesean 2026. All Rights Reserved.
Terms and Conditions / Privacy Policy
"The Antikythera Mechanism" by Thomas Weibel is licensed under CC BY 4.0. This work has been modified from the original.
