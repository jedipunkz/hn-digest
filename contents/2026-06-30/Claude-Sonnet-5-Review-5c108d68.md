---
source: "https://www.coderabbit.ai/blog/claude-sonnet-5-review"
hn_url: "https://news.ycombinator.com/item?id=48737325"
title: "Claude Sonnet 5 Review"
article_title: "Claude Sonnet 5 review: Should you switch?"
author: "bgubitosa"
captured_at: "2026-06-30T19:34:02Z"
capture_tool: "hn-digest"
hn_id: 48737325
score: 2
comments: 0
posted_at: "2026-06-30T18:39:34Z"
tags:
  - hacker-news
  - translated
---

# Claude Sonnet 5 Review

- HN: [48737325](https://news.ycombinator.com/item?id=48737325)
- Source: [www.coderabbit.ai](https://www.coderabbit.ai/blog/claude-sonnet-5-review)
- Score: 2
- Comments: 0
- Posted: 2026-06-30T18:39:34Z

## Translation

タイトル: クロード・ソネット 5 レビュー
記事のタイトル: クロード・ソネット 5 レビュー: 乗り換えるべきですか?
説明: 実際のコーディングとコードレビュー作業を 1 週間行った後の Claude Sonnet 5 の実地レビュー: Sonnet 4.6 との比較、コスト、および誰がアップグレードすべきかについて説明します。

記事本文:
クロード・ソネット 5 レビュー: 切り替えるべきですか?エージェント エンタープライズ顧客の価格設定 ブログ リソース ドキュメント
サインアップすると、利用規約に同意し、CodeRabbit が製品やソリューションに関する最新情報を不定期に提供することを承認したことになります。あなたはいつでもオプトアウトすることができ、あなたのデータはCodeRabbitプライバシーポリシーに従って取り扱われることを理解するものとします。
サインアップすると、利用規約に同意し、CodeRabbit が製品やソリューションに関する最新情報を不定期に提供することを承認したことになります。あなたはいつでもオプトアウトすることができ、あなたのデータはCodeRabbitプライバシーポリシーに従って取り扱われることを理解するものとします。
クロード・ソネット 5 レビュー: 切り替えるべきですか?
Claude Sonnet 5 のコードの書き方 テストを習慣として扱う
Claude Sonnet 5 によるコードのレビュー方法 コメントがよりわかりやすくなりました
虫取りの正直なところ
クロード Sonnet 5 と現在使用しているモデル Sonnet 5 とフラッグシップ モデル
クロード ソネット 4.6 またはクロード ソネット 5: 迅速な判定
コードレビュー時間とバグを 50% 削減
GitHub と GitLab で最もインストールされている AI アプリ
受信トレイで最新情報をすぐにキャッチできます。
受信トレイで最新情報をすぐにキャッチできます。
CodeRabbit エージェントが Discord に追加されました
CodeRabbit Agent を Discord サーバーに導入します。チャネルを、CodeRabbit が調査、計画、自動化して実際のコード作業を行うワークスペースに変えます。
概要ページの紹介: あなたが実際に望んでいたプル リクエストのホームページ
[概要] ページは、PR とは何か、および PR をマージできるかどうかに答える単一の PR レベルのホームです。1 つはレビュー担当者用、もう 1 つは作成者用です。
コードレビューに関して間違っていたこと
AI によって生成されたコードの量が増加するにつれて、より困難な問題が明らかになりました。レビューは速度だけではなく、開発者が自分たちが出荷するものをまだ理解し、信頼しているかどうかも重要です。
お使いのブラウザはビデオをサポートしていません。人類は出荷するだけ

ed Claude Sonnet 5 は、中級ラインの最新モデルです。この投稿では、「現在使用しているモデルを使い続けるべきか、それともこのモデルに移行すべきか?」という 1 つの単純な質問に答えます。
ソネット 5 はどこからともなく生まれたわけではありません。先月、私たちはこのブログで一連の新モデルをレビューしてきました。
Opus 4.8 は、レビューの指示に忠実に従うのに十分な注意を払って、長いマルチステップのコーディング用にテストした最良のモデルでした。ただし、コストは高く、大規模なジョブではほとんどの場合利益が得られました。
Fable 5 は独自のコーディングにさらに熱心に取り組み、多くのファイルにまたがる計画と構築を喜んで行いました。ただし、価格とアクセスが制限されているため、デフォルトのレビュー パスから外れていました。
NVIDIA の Nemotron 3 Ultra は、その逆を行きました。つまり、高速かつ要点を絞った、エージェント設定内で多くの迅速な対応を行うために構築されたオープン モデルです。
ソネット 5 は、Anthropic ファミリーの残りの部分のように感じられます。忍耐強くて徹底的で、問題を行動に移す前に徹底的に考えることを好みます。
コードの作成と構築に関しては、Sonnet 5 がこの層でこれまでに使用した中で最も有能なモデルであり、簡単にアップグレードできるので楽しみです。レビューの場合、それはむしろトレードオフです。よりクリーンで鋭いコメントを生成しますが、現在運用環境で実行している以前のモデルよりも検出されるバグが少なく、レビューあたりのコストが若干高くなります。
良い点は、そのほとんどを調整できることです。多くのチームにとって、この移行は実行する価値があります。 Sonnet 5 の新機能、コードの作成とレビューがどのように実行されるか、切り替えることが合理的かどうかを見ていきましょう。
Sonnet 5 は、前のバージョンよりもさらに深く考えられています。あなたにとって、それは、古いモデルではあきらめていたであろうより困難な問題を解決できることを意味します。
Sonnet 5 のシンキング「エフォート」ダイヤルは、必要がない場合に段階的にオフにできるようにします。これは保護する機能です

あなたの予算。見逃したバグが高くつくような難しいレビューの場合は労力を増やし、深く考えるのにお金をかけたくない日常的な作業の場合は労力を減らすかオフにします。
タスクの途中で自身の命令を書き換えることもできます。長期にわたるエージェント ジョブでは、モデルが学習するにつれて目標が変化する傾向があり、モデルが最初の計画に固執すると、適合しなくなった指示を押し続けることになります。 Sonnet 5 は代わりに独自のプランを更新するため、逸れて時間とトークンを消費する実行が少なくなります。
また、セキュリティとサイバーに関する新しい安全ガードレールも付属しています。利点は、リスクのある生産物が減少することです。問題は、実際のセキュリティ作業では時々フィルターが作動する可能性があるため、その地域の場合は奇妙な拒否が予想されることです。
Sonnet 5 をイメージする最も簡単な方法は、本当に機能し、要求されたレベルで実行されるコードを出荷することに、おそらく少し過剰に気を配る中級レベルのエンジニアとして想像することです。この 1 つの特性がその動作のほとんどを形作り、それは私たちが何度も目にした 4 つの習慣に現れました。
機能の前にテストを作成する傾向があります。
すでに機能した後も、ソリューションを磨き続けます。
それは自分自身の計画を思い直し、時には必要以上に推測します。
小さなプロジェクトで小さなタスクに答えます。
Claude Sonnet 5 のコードの書き方
Sonnet 5 のコード作成は、ほとんどのチームがアップグレードを希望すると思われる主な理由です。レビューの数値に入る前に、何かを最初から構築するときにどのように動作するかを確認すると役に立ちます。私たちは、簡単な機能から明確な道筋のない難しい問題まで、私たちが日々行っている作業をそれに任せました。
ある晩、私たちは難しい任務を与えられ、昼食をとりに出かけました。このジョブは、モデルにコードを記述し、その上でシミュレーションを実行し、可能な限り良い結果が得られるまで出力を調整し続けるように要求しました。私たちはそれが完了するか行き詰まると期待して戻ってきました。どちらでもない

本当だった。モデルは依然として続行しており、私たちから何も促されることなく、独自に問題を解決していました。
アプリケーション全体を独自に構築しました。こんなに長く走った理由は単純だった。たまたまうまくいった最初の答えではなく、最良の答えを追い求めていたため、パスごとに独自のソリューションをクリーンアップし続けました。このようなことは、以前は高価なモデルでのみ可能でしたが、中級モデルがそれを行うのを見るのは本当の瞬間です。
前述の勤勉さこそが、Sonnet 5 が長期にわたる無制限の作業に優れている理由です。たった 1 行の変更でイライラする遅さは、ジョブのステップ数が決まっていない場合には強みに変わります。これは、モデルに単純な目標を渡し、数ラウンドかけてアプローチを試し、テストし、報告する前に結果を独自に改善させるエージェント ループに最適です。
効果的なエージェントを構築するための Anthropic 独自のガイドでは、これを評価ループと呼んでいます。ここでは、モデルが結果を記述し、それをサイクルで批判し、改善します。この種のワークフローを構築している場合は、一読の価値があります。 Sonnet 5 ではそのループが思考方法に組み込まれており、以前のモデルがずれたりあきらめたりしたためにエージェント コーディングを控えていた場合、このリリースはそれを変えるものです。
Sonnet 4.6 を含め、ほとんどのモデルはテストを後の作業として扱います。 Sonnet 5 は、最初にテストを作成し、その上に機能を構築し、完了したと判断したらすべてを実行する傾向があります。私たちが見た自己チェックは、まさにその順序から生まれたものです。テストを実行しなければ、テストとコード間の衝突を見つけることはできません。このモデルは常にテストを実行します。正常に見えたコードを出荷して 1 週間後に壊れた経験がある人なら、その違いをすぐに感じるでしょう。
Sonnet 5 が特別に配慮していることは、ファイル数とトークンの使用量に現れます。

。小さなことを要求すると、たくさんのものが戻ってきます。追加のヘルパーと、機能自体よりも長いテスト ファイルが表示されます。優れたエンジニアリングのように見える大きな機能について。たった一行の変更ではどうしようもないモデルのように見えます。
Sonnet 5 は Sonnet 4.6 よりも遅いことがわかりました。これはおそらく Sonnet 5 が実行する特別な思考によるものと考えられます。徹底的な作業を行うために分を犠牲にすることになり、長時間実行し続けるジョブの成果が得られます。ちょっとした小銭を待っているときは痛くなります。 4.6 はすぐに答えを返しますが、5 はより良いものを求めて働き続け、小さなことに最もそれを感じます。 2 つの小さなメモ: 4.6 よりも多くのトークンが使用されていますが、依然として明確ですが、おしゃべりなだけです。また、優れた計画を作成しても、タスクの途中でそれを書き直すことが、私たちが望むよりも頻繁に行われます。これらはビルド作業にとって問題となるものではありませんが、小さな編集の山を指摘する前に知っておく価値があります。
Claude Sonnet 5 がコードをレビューする方法
コードを大規模にレビューすることが私たちの仕事であるため、これは私たちが最も重視している部分です。現在、多くのチームが出荷するコードの大部分を AI が作成しており、そのコードには慎重な第 2 の目が必要です。 470 件のオープンソース プル リクエストを調査したところ、AI が共同作成した PR は人間のみの PR よりも約 1.7 倍多くの問題を抱えていることがわかりました。また、2025 年の調査では、AI に重点を置いたチームではレビュー時間が 91% 増加したことがわかりました。レビューを行うモデルは、チームがどれだけ早く出荷するかに大きな影響を与えるようになりました。
Sonnet 5 をハーネスに追加し、既知のバグを含む固定のプル リクエストのセットである標準ベンチマークをポイントし、キャッチした数とコメントのクリーンさを測定しました。これらのレビューを構築する方法について詳しくは、AI コード レビューのコンテキスト エンジニアリングに関する記事をご覧ください。
Sonnet 5 のコメントはよりクリーンであり、その発見結果はノイズではなくバグであることが多かったです。私たちのベンチマークでは、精密上昇

Sonnet 4.6 では約 29% から約 38% ～ 40% に増加します。 Sonnet 4.6 はその逆です。ほぼすべてのことについてコメントがあり、乱雑な中からキーパーを選択することができます。些細なことでも指摘するレビューアーを無視したことがあれば、注意が必要なものを強調することがいかに重要であるかをすでにご存知でしょう。
虫取りの正直なところ
私たちのレビュー ハーネス内では、Sonnet 5 は現在の運用環境よりも少ないバグを発見しました。厳密な「バグが見つかったかどうか」の基準では、ベースラインは約 57% を捕捉し、Sonnet 5 は約 50 ～ 51% を検出しました。私たちが驚いたのは、Sonnet 4.6 のほうがうるさいレビュアーであるにもかかわらず、どちらよりも多くの結果を獲得し、約 63% だったことです。つまり、コメントに埋もれているモデルは、見逃されるバグが最も少ないモデルでもあります。 Sonnet 5 の努力ダイヤルを上げてもスコアはほとんど変わらず、コストは約 2 倍になりました。より多くの種類やバリエーションのコメントをカウントする緩やかなスコアでは、高い努力により、結果は現在のベースラインとほぼ同等に戻りましたが、明らかに前進することはできませんでした。
コメントにも少しノイズが隠れています。 Sonnet 5 は、ベースラインの 3 ～ 4 倍の細かい点を投稿し、負荷の高い実行では Sonnet 4.6 よりもほぼ 80% 多く投稿します。最高のコメントはよりわかりやすく読みやすくなりますが、それを見つけるには、よりマイナーなコメントを無視する必要があります。私たちも思考をオフにして試してみました。一部のジョブではベースラインと一致しましたが、他のジョブでは少しずれていました。したがって、最も重要なコードでない限り、何も考えずに低コストで実行でき、ほとんど何も失うことのない、より軽いレビュー作業のクラスがあります。要点は、Sonnet 5 が弱いレビュアーであるということではありません。実際には、それはより静かでより慎重な方法であり、レビューの騒音に溺れているチームにとっては、多くの場合、それがより良い取引です。
Claude Sonnet 5 と現在実行しているモデルの比較
Sonnet 4.6 wh に手を伸ばしてください

必要なのは生の報道であり、そのコメントを読み進めるための帯域幅を持っています。より少なく、より鋭いコメントを取得し、実際にコードを記述するための強力なパートナーを希望する場合は、Sonnet 5 をお勧めします。 Sonnet 4.6 ではさらにいくつかのバグが見つかりましたが、Sonnet 5 では無駄な注意が大幅に減りました。
Sonnet 5 とフラッグシップモデルの比較
チームが最もハードな仕事のために維持する Opus シリーズのようなフラッグシップ モデルと比較して、Sonnet 5 ははるかに低コストでレビュー品質においてほぼ同等の性能を発揮します。安いものでは十分ではなかったという理由だけで、現在フラッグシップ料金を支払っているのであれば、Sonnet 5 は一見の価値があります。テスト中はトークンの請求額に注意してください。これほど多くのことを考えるモデルは、それ自体の節約を食いつぶす可能性があるからです。努力を最大限まで高めるのは最悪の取引だった。意味のあるバグがさらに見つかることなく、コストがおよそ 2 倍になったため、デフォルトで最上位層には手を伸ばさないでください。 Claude Opus 4.7 が AI コード レビューに何を意味するかを検討したときにも、同じパターンが見られました。
クロード ソネット 4.6 またはクロード ソネット 5: 迅速な判定
ほとんどのチームにとって、答えは「はい」です。 Sonnet 5 は、このクラスで使用した中で最もエキサイティングなコーディング モデルです。それは慎重なチームメイトのように機能し、むしろ数分余分に時間を費やしたいタイプです

[切り捨てられた]

## Original Extract

Hands-on review of Claude Sonnet 5 after a week of real coding and code-review work: how it compares to Sonnet 4.6, what it costs, and who should upgrade.

Claude Sonnet 5 review: Should you switch? Agent Enterprise Customers Pricing Blog Resources Docs
By signing up you agree to our Terms of Use and authorize CodeRabbit to provide occasional updates about products and solutions. You understand that you can opt out at any time and that your data will be handled in accordance with CodeRabbit Privacy Policy
By signing up you agree to our Terms of Use and authorize CodeRabbit to provide occasional updates about products and solutions. You understand that you can opt out at any time and that your data will be handled in accordance with CodeRabbit Privacy Policy
Claude Sonnet 5 review: Should you switch?
How Claude Sonnet 5 writes code It treats testing as a habit
How Claude Sonnet 5 reviews code The comments are much cleaner
The honest catch on bug-catching
Claude Sonnet 5 vs the model you run today Sonnet 5 versus a flagship model
Claude Sonnet 4.6 or Claude Sonnet 5: Quick verdict
Cut code review time & bugs by 50%
Most installed AI app on GitHub and GitLab
Catch the latest, right in your inbox.
Catch the latest, right in your inbox.
CodeRabbit Agent is now in Discord
Bring the CodeRabbit Agent into your Discord server. Turn channels into workspaces where CodeRabbit investigates, plans, automates, and does real code work.
Meet the Overview page: The pull request home page you actually wanted
The Overview page is a single PR level home that answers what a PR is and whether it can merge: one answer for the reviewer, one for the author.
What we got wrong about code review
As the volume of AI-generated code grew, the harder problem became clear: review isn't only about speed, it's about whether developers still understand and trust what they're shipping.
Your browser does not support the video. Anthropic just shipped Claude Sonnet 5 , the newest model in its mid-tier line. This post answers one simple question: Should you stay on the model you use today, or move up to this one?
Sonnet 5 didn't land out of nowhere. We've reviewed a string of new models on the blog over the past month:
Opus 4.8 was the best model we'd tested for long, multi-step coding, careful enough to follow review instructions to the letter, though it cost more and paid off mostly on bigger jobs.
Fable 5 leaned even harder into coding on its own, and was happy to plan and build across lots of files. However, its price and limited access kept it off our default review path.
NVIDIA's Nemotron 3 Ultra went the other way: fast and to the point, an open model built to take lots of quick swings inside an agent setup.
Sonnet 5 feels like the rest of the Anthropic family. It's patient and thorough, and it likes to think a problem all the way through before it acts.
For writing and building code, Sonnet 5 is the most capable model we've worked with at this tier, and it's an easy upgrade to be excited about. For review, it's more of a tradeoff. While it generates cleaner, sharper comments, it catches fewer bugs than the earlier models we currently run in production, and comes at a slightly higher cost per review.
The good part is that you can tune most of that, and for a lot of teams the move is well worth making. Let's walk though what's new in Sonnet 5, how it performs writing and reviewing code, and if it makes sense to switch.
Sonnet 5 thinks more deeply than the version before it. For you, that means it works through harder problems the old model would have given up on.
The thinking "effort" dial in Sonnet 5 that allows you to incrementally turn it down to off if you don't need it. This is the feature that protects your budget. Turn the effort up for a tricky review where a missed bug is expensive, and turn it down or off for routine work where you'd rather not pay for deep thinking.
It can also rewrite its own instructions partway through a task. On long agent jobs, the goal tends to shift as the model learns more, and a model stuck on its first plan will keep pushing on instructions that no longer fit. Sonnet 5 updates its own plan instead, so you get fewer runs that wander off and burn your time and tokens.
It also ships with new safety guardrails around security and cyber topics. The upside is fewer risky outputs. The catch is that real security work can trip the filters now and then, so expect the odd refusal if that's your area.
The easiest way to picture Sonnet 5 is as a mid-level engineer who cares, maybe a little too much, about shipping code that truly works and runs at the level you asked for. That one trait shapes most of what it does, and it showed up in four habits we saw again and again:
It tends to write tests before the feature.
It keeps polishing a solution long after it already works.
It second-guesses its own plan, sometimes more than necessary.
It answers a small task with a small project.
How Claude Sonnet 5 writes code
Sonnet 5’s code writing is the main reason we think most teams will want to upgrade. Before we get to the review numbers, it is helpful to see how it behaves when it builds something from scratch. We handed it the work we do daily, from quick features to harder problems with no clear path.
One evening, we gave it a tough task and stepped out for lunch. The job asked the model to write the code, run simulations on it, and keep tuning the output until the results were as good as possible. We came back expecting it to be done or stuck. Neither were true. The model was still going, working through the problem on its own, with no nudging from us.
It built the whole application by itself. The reason it ran so long was simple. It kept cleaning up its own solution, pass after pass, because it was chasing the best answer instead of the first one that happened to work. That's the kind of thing that was previously only available with a pricier model, and watching a mid-tier model do it is a real moment.
That aforementioned diligence is exactly why Sonnet 5 is good at long, open-ended work. The slowness that bugs you on a one-line change turns into a strength when the job has no set number of steps. It's a great fit for agent loops, where you hand the model a simple goal and let it spend a few rounds trying approaches, testing them, and improving the result on its own before it reports back.
Anthropic's own guide to building effective agents calls this an evaluator loop, where the model writes a result and then critiques and improves it in a cycle. It's worth a read if you're building that kind of workflow. Sonnet 5 has that loop wired into how it thinks, and if you've been holding off on agentic coding because earlier models drifted or gave up, this is the release that changes that.
Most models treat tests as a chore for later, Sonnet 4.6 included. Sonnet 5 tends to write the tests first, builds the feature on top of them, and then runs everything once it thinks it's done. The self-checking we saw comes straight out of that order. You can't spot a clash between your tests and your code if you never run the tests, and this model always runs them. If you've ever shipped code that looked fine and broke a week later, you'll feel that difference fast.
The extra care Sonnet 5 takes shows up in your file count and token use. Ask for something small and you get a lot back. You'll see extra helpers and a test file longer than the feature itself. On a big feature that looks like good engineering. On a one-line change it looks like a model that can't help itself.
We found Sonnet 5 to be slower than Sonnet 4.6, which is likely due to the extra thinking Sonnet 5 performs. You're trading minutes for thoroughness, which pays off on long jobs you leave running. It does sting when you're waiting on a small change. Where 4.6 hands back a quick answer, 5 keeps working for a better one, and you feel that most on the little stuff. Two smaller notes: It uses more tokens than 4.6 did, and is still clear, just chattier. It also writes good plans, then rewrites them mid-task more often than we'd like. None of this is a dealbreaker for build work, but it's worth knowing before you point it at a pile of tiny edits.
How Claude Sonnet 5 reviews code
This is the part we care about most, since reviewing code at scale is what we do. AI now writes a big share of the code many teams ship, and that code needs a careful second pair of eyes. Our look at 470 open-source pull requests found that AI-co-authored PRs carry about 1.7 times more issues than human-only PRs , and one 2025 study found review time climbed 91% on teams leaning hard on AI. The model doing the reviewing is now a big factor in how fast teams ship.
We added Sonnet 5 to our harness and pointed it at our standard benchmark, a fixed set of pull requests with known bugs, and measured how many it caught and how clean its comments were. You can read more about how we build these reviews in our piece on context engineering for AI code reviews .
Sonnet 5's comments are cleaner and its findings were more often bugs than noise. On our benchmark, precision climbed from about 29% with Sonnet 4.6 to roughly 38% to 40%. Sonnet 4.6 does the opposite. It comments on almost everything and leaves you to pick the keepers out of the clutter. If you've ever turned a reviewer off because it flagged every little thing, you already know how important it is to highlight what needs your attention..
The honest catch on bug-catching
Inside our review harness, Sonnet 5 catches fewer bugs than our current production setup. On the strict "did it find the bug" measure, our baseline catches about 57%, and Sonnet 5 lands around 50 to 51%. What surprised us was that Sonnet 4.6 caught more than either of them, around 63%, even though it's a noisier reviewer. So the model that buries you in comments was also the one that missed the fewest bugs. Turning Sonnet 5's effort dial up barely moved its score and roughly doubled the cost. On a looser score that counts more types and variations of comments, high effort brought its findings back to about even with the current baseline, but it didn't pull clearly ahead.
There's a bit of noise hiding in the comments too. Sonnet 5 posts three to four times more nitpicks than our baseline and almost 80% more than Sonnet 4.6 in high effort runs. Its best comments read cleaner, but you have to look past more minor ones to find them. We also tried it with thinking turned off. On some jobs it matched our baseline, on others it slipped a little. So there's a class of lighter review work you can run cheap with thinking off and lose almost nothing, as long as it isn't your most important code. The takeaway isn't that Sonnet 5 is a weak reviewer. It's actually a quieter, more careful one, and for teams drowning in review noise, that's often the better trade.
Claude Sonnet 5 vs the model you run today
Reach for Sonnet 4.6 when raw coverage is what you need and you've got the bandwidth to wade through its comments. Reach for Sonnet 5 when you'd rather get fewer, sharper comments and a much stronger partner for actually writing code. Sonnet 4.6 finds a few more bugs, while Sonnet 5 wastes far less of your attention.
Sonnet 5 versus a flagship model
Against flagship models, like the Opus family that teams keep for their hardest jobs, Sonnet 5 does nearly as well on review quality for a lot less money. If you're paying flagship rates today only because nothing cheaper was good enough, Sonnet 5 is worth a look. Keep an eye on the token bill while you test, because a model that thinks this much can eat into its own savings. Cranking the effort to maximum was the worst deal of the bunch. It roughly doubled the cost without finding meaningfully more bugs, so don't reach for the top tier by default. We saw the same pattern when we reviewed what Claude Opus 4.7 means for AI code review .
Claude Sonnet 4.6 or Claude Sonnet 5: Quick verdict
For most teams, the answer is yes. Sonnet 5 is the most exciting coding model we've used in its class. It works like a careful teammate, the kind who would rather take an extra few minutes th

[truncated]
