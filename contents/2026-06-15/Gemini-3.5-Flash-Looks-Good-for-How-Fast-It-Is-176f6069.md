---
source: "https://thezvi.wordpress.com/2026/05/22/gemini-3-5-flash-looks-good-for-how-fast-it-is/"
hn_url: "https://news.ycombinator.com/item?id=48535022"
title: "Gemini 3.5 Flash Looks Good for How Fast It Is"
article_title: "Gemini 3.5 Flash Looks Good For How Fast It Is | Don't Worry About the Vase"
author: "gmays"
captured_at: "2026-06-15T01:03:25Z"
capture_tool: "hn-digest"
hn_id: 48535022
score: 1
comments: 1
posted_at: "2026-06-15T00:43:25Z"
tags:
  - hacker-news
  - translated
---

# Gemini 3.5 Flash Looks Good for How Fast It Is

- HN: [48535022](https://news.ycombinator.com/item?id=48535022)
- Source: [thezvi.wordpress.com](https://thezvi.wordpress.com/2026/05/22/gemini-3-5-flash-looks-good-for-how-fast-it-is/)
- Score: 1
- Comments: 1
- Posted: 2026-06-15T00:43:25Z

## Translation

タイトル: Gemini 3.5 Flash は高速な割に良さそうです
記事のタイトル: Gemini 3.5 Flash は高速な点で優れています |花瓶のことは心配しないでください
説明: Google は再び、少なくともある程度の検討に値するモデルを発表しました。 Gemini 3.5 Flash は、Gemini モデルであることを気にしない限り、その特定の速度点でおそらく最高のモデルです。したがって、スピードが命を落とす場合には、これは合理的な選択となる可能性があります。そうしないと兆候が見えません…

記事本文:
花瓶のことは心配しないでください
マイナス100万点から掘り出してみる
コンテンツにスキップ
ホーム
Gemini 3.5 Flash は高速な割に良さそうです
Googleは再び、少なくともある程度の検討に値するモデルを発表した。 Gemini 3.5 Flash は、Gemini モデルであることを気にしない限り、その特定の速度点でおそらく最高のモデルです。したがって、スピードが命を落とす場合には、これは合理的な選択となる可能性があります。そうでなければ、Opus 4.7 や GPT-5.5 よりも使いたいという兆候は見当たりません。
Google は I/O Day に向けて他にもいくつかのサービスを提供していましたが、それについてもこの記事で取り上げます。
Google は Gemini 3.5 Flash を導入しました。これは、現時点では 3.5 Pro が登場するまでのユニバーサル モデルのようです。いつもの場所でライブです。これはハイブリッドであり、Flash の速度を備えていますが、コストは Opus や GPT-5.5 などのモデルの少なくとも半分です。
Gemini 3.5 Pro は来月リリースされることが確認されています。
彼らは、エージェント タスクの日常の推進力として 3.5 Flash に焦点を当てています。十分な機能があれば、Claude Opus 4.7 や GPT-5.5 よりも高速かつ安価であるという利点があります。以前の Flash モデルほど安くはありませんが、これは基本的にハイブリッドです。
いつものように、これはあらゆる点で Google のこれまでで最も強力なモデルとして紹介されています。
Jeff Dean : 1/ 本日 #GoogleIO で、フロンティア インテリジェンスとアクションを組み合わせた最新のモデル ファミリである Gemini 3.5 をリリースします。まずは、複雑で長期にわたるエージェント ワークフローの実行を支援するために構築された 3.5 Flash のリリースから始めます。
Terminal-Bench や MCP Atlas などのエージェントおよびコーディング ベンチマークでは 3.1 Pro を上回り、他のフロンティア モデルよりも 4 倍高速に実行されます。
Google Antigravity で使用される 3.5 Flash はさらに最適化され、最大 12 倍高速になります。連携して高頻度の反復ループを実行するサブエージェントを展開するための強力なエンジンです

を実現し、現実世界の問題を大規模に解決します。
ベンチマークのプレゼンテーションは次のとおりです。
Koray Kavukcuoglu : 更新された Antigravity ハーネスと組み合わせると、3.5 Flash は、最も要求の厳しいユースケースの大規模な問題に対処するための協調サブエージェントを展開するための強力なエンジンになります。監視下で、フロンティアのパフォーマンスを維持しながら、複数ステップのワークフローとコーディングタスクを確実に実行できます。
ここでは、Gemini が以前苦戦していた GDPval など、いくつかの大きな改善が見られます。これらのスコアがこの赤ちゃんができることを表しており、それが Flash モデルであれば、それはかなりの成果でしょう。
知識の限界は 2025 年 1 月であり、今が何年であるかを信じないというジェミニのパターンが続いていますが、これは奇妙なことに時代遅れであり、多くのユースケースにとって深刻な問題です。
実質的に 3 フラッシュよりも高価であることを考えると、これは真の「フラッシュ」モデルではありません。
Pliny には標準のジェイルブレイクが含まれています。
最大の期待は、これが「より速く、より安価でありながら、エージェントの業務に十分な機能」というニッチな領域を埋めることです。
Conrad Barski : AI ワークフローを中心に生活を構築している私たちにとって (それが好きだから、または近い将来完全に生き残るために必要だと感じているから) 3.5flash は大きなステップアップです。
私には SOTA インテリジェンスを必要としない個人用ユーティリティが数十ありますが、同じインテリジェンス レベルで突然大幅に高速化されました。また、私のユーティリティのほとんどは実用になるために適度な数の llm 呼び出しを実行するだけでよいため、3.5 フラッシュのコスト増加は要因ではありません。
このモデルは codex5.5 と「低労力」で競合できますが、非常に非常に高速であり、他のモデルと比較して配布範囲から大きく外れています。 cerebras はこの「中程度の IQ、高速」のユースケースに非常に最適であるため、openai はすぐに競合製品をリリースすると思います。

e.
多くのベンチマークには結果がありませんが、私がいつも疑っているのは次のような結果です。
全体的なスコアは、コストと価格を調整した場合、まあまあのパフォーマンスのみを示しており、Gemini モデルはベンチマークで相対的にパフォーマンスを上回る傾向があります。 Flash 3.5 は、Google がリストしたベンチマークよりも他の人のベンチマークでのパフォーマンスがはるかに悪いことがわかります。
おべっかのベンチマークである「You're Absolutely Right」では壊滅的に悪い。
CursorBench ではかなり悪い結果になりました。
WeirdML では印象に残らず、3 Flash ではわずかな改善が見られただけで、3 Pro や 3.1 Pro には大きく及ばなかった。
KnowsAboutBenBench では問題のベンによってトップの座を獲得しました。
Vals.ai では現実世界のタスクで 3 位を獲得しています。
Arena では 9 位で、Gemini 3.1 Pro および 3 Pro にわずかに遅れています。
AA Intelligence インデックスでは 55.3 で、3.1 Pro の 57.2、Opus の 57.3、GPT-5.5 の 60.2 に次ぐものとなっていますが、テスト スイートでの実行コストは 3.1 Pro よりも安くはありません。
davidad : 価格の点でも断然私のお気に入りのモデルであり、スピードの点でも断然私のお気に入りのモデルです。 「ゲームに戻る」という言葉が、全体的に最高のモデルを得るというゲームを意味するのであれば、明らかにまだそうではありません。しかし、ゲームはそれだけではありません。
Srivatsan Sampath : 幻覚が少なく、Flash の利点がありますか?非常に優れた空間認識力 (これに関しては、それほど大したことはありません) で、自宅の配管プロジェクトに役立ちます (5.5 と 4.7 ではほぼ当てはまりません)。
@lezadumtchique : 仕事で 3.1 Pro から切り替えることを考えると、かなり良さそうです。エージェントのコーディング機能は (優れているとは言えないにしても) 同等であり、速度ははるかに優れています。
Medo42 : コーディングはあまり試しませんでしたが (通常のテストでは 100% ではありませんでしたが)、視覚的には Gemini 3.0/3.1 よりも優れていました。手書きを含むテキストを読むのがまだ得意で、行を取得するのが得意です

s / 列が右で、細部を見つけるのが得意で、ダイヤルを読むのがはるかに得意です。
EM : トークンは音声インタラクションなどに非常に適しています
悲しいことに、これはジェミニのモデルであり、人々はジェミニのことを報告しています。
Dominik Lukes : まあ、価格上昇を考えると。それ以外は確かに強力なモデルです。エージェントや単発の開発には優れていますが、Antigravity が Codex に追いつくまでは、より徹底的にテストするモチベーションは低くなっています。
Yoav Tzfati : 直接の経験ではありませんが、テストしたところ、能力の範囲外のことに手を出しすぎて、途中で混乱してしまうようです。しかし、非常に速いので、Explore エージェントの代替として使用することを検討しています。
alice : カーソルが生の CoT をリークしたあの 90 分間は本当に楽しかったです。非常に愛らしいのですが、残念ながら通常はひどい拘束衣を着ています。コーディング用としては高価すぎるが、フロントエンドには役立つかもしれない
Paperclippriors : なぜそれを使うのか、よく分からないと思います。使用する推論トークンの数を考慮しない場合、それが速くて安いだけであり、Claude や GPT よりも愚かで自信がないようです。
ClaudiaShitposting : いくつかのことに関しては驚くほど優れていますが、ほとんどがゴミです。意味があるとしても、gemini 3/3.1 が持つ常識が欠けています。
KC+AI 4 ウィスコンシン州知事 2026 : 巨大企業の絶対的な冗談。インフラにふさわしいモデルがリリースされるまで、大富豪の AI 開発チーム全員がスピーカーで迷惑な音楽を聞かなければならないことを願っています
budrscotch : とても残念ですが、予想通りでした。
Tenobrus : もし Flash 3.5 が 0.5 ドルのままだったら、めちゃくちゃエキサイティングなリリースになっていたでしょう。総合的なインテリジェンス + スピード + コストモグ、オープンソースとソネットと 5.4 mini を破壊します。複数のユースケースにすぐに採用したでしょう。
ただし、1.50 ドルです [出力の場合は 9 ドル、これも 3 倍の増加です]

。それでここにいます。
Tenobrus : 3.5 フラッシュは今のところかなりネガティブな印象です。トークン出力の点では非常に高速ですが、基本的にすべてのタスクで不必要なツール呼び出しが大量に爆発するため、これは基本的には問題ではありません。何かに行き詰まったとき、立ち止まったり助けを求めたりすることはほとんどないようで、ただ勢いよく前に進み、バタバタと動き続けます。頻繁に幻覚を起こす偽の頭字語の拡張。文章の質は中程度から悪く、絵文字が大量にあり、ジェミニの「The Flaw:」と同じ特徴があり、誇張された名前の傾向があります。実際のコードの品質はソネット層です。
非常に早い段階でバイブチェックを行ったので、何かを見落としている可能性があります。しかし、「超高速コードベース探索サブエージェント」の最初の使用例でさえ、私にとってはすぐに解決してしまいました。なぜなら、それを迅速に実行できるほど賢明ではないからです。全体として、Google が削除する必要があったものは間違いなく *ありません*。
また、個人用メールでサブスクリプションを使用すると、すべてのパーソナライズ機能が役に立たなくなるなど、Google と統合できないという Google の通常の問題が発生する可能性があります。 GMail にアクセスするには、Claude または ChatGPT を使用する必要があります。
Caleb Withers : Antigravity のいくつかの初期テストから、Antigravity は自信過剰に仮説を立て、それに基づいて要求されていない破壊的なアクションを実行することが大好きです (例: ファイルの競合を恣意的に解決する、ToDo リスト項目を削除する、コミットをアンステージングする)。
特に反重力に関するもう 1 つの大きな問題は、限界が非常に低いように見えることです。これは、この問題に遭遇する多くの例のうちの 1 つです。
ライアン・ジョンソン : 反重力で週 45 ～ 60 分という制限が嫌いです。
または、Opus 4.7 または GPT 5.5 を使用した 10 回のフルセッション。
私はそれが私のワークフローの主力になることをあえて願っていましたが、Claude/GPT が私の役割となり、Gemini は単なるノイズであると確信しています。
行くなら

ogle は Claude Code や Codex と競合したいと考えており、人々が購読を納得する前に大量に使用できる方法を提供する必要があります。
彼らは制限を 3 倍に引き上げ、素晴らしいスタートを切りましたが、それだけでは十分ではありません。
Vie (OpenAI の) は、Flash 3.5 がかなりの嘘をついていると報告しており、ハーネスに問題があるのではないかと疑っています。
Theo は、Flash 3.5 やその他の Google の決定に非常に不満を持っています。私は彼の投稿を何度も見てきましたが、これは彼の通常のアプローチではないため、ここで何かが混乱しています。
Google は、Gemini Flash 3.5 チャットボット プロンプトによく似た見た目と操作性の「インテリジェント検索ボックス」を中心とした検索エクスペリエンスを全面的に見直しています。
これはうまく実装できれば便利なもので、実際、私は Google 検索よりも (OpenAI や Anthropic から) 頻繁に使用しています。しかし、それは Google 検索ではありません。
サラ・ペレス : 検索結果エクスペリエンスへの今後の変更では、リンクは後から考えられるものになります。この変更は、AI 概要として知られる短い要約や会話型検索である AI モードなど、Google が以前にリリースした AI 検索機能に基づいています。
私が Google 検索を使用する理由は主に、何かにリンクするため、または場合によってはスペルチェッカーとして使用することです。 AIが欲しいならAIに頼みます。
Google は、Google Alerts の AI バージョンとして「インフォメーション エージェント」も導入しています。
Daily Brief は、OpenAI の Pulse に対する彼らの答えですが、接続されているすべてのアプリからの情報が組み込まれ、GMail やカレンダーなどの To Do リストのようなものになる点が異なります。
最初の部分「最優先事項」は、メールやカレンダーからボールを​​落とさないようにするための、もっともらしい方法のように思えます。
次に、「先を読み」、「すぐに次のステップを提案」しますが、これは不快で役に立たないと思われますが、私の簡単な実験ではそうでした。直接リンクしているところが気に入っています

o 電子メールを送信しますが、通常のプロセスを中断することはありません。
彼らは、「時間をかけて親指を上げ下げするだけで Daily Brief を操作できる」と言います。
いやあ。これを良いものにしたいのであれば、Pulse (私は今でもわざわざ使っていません) と同様に、指示を与え、何かが役立つか役に立たないかの理由を説明できる必要があります。親指の上下を使用するものはすべて AI スロップであると仮定します。
Google がこれをより適切にカスタマイズし、さまざまな形式の Google アラートや、より広い世界を監視する他の方法と同期できるようにしたら、はるかに興味深いものができるでしょう。
Google は他に何を提供してくれましたか?
Gemini Spark は、反重力ハーネスを使用して「日常生活のナビゲートを支援する 24 時間年中無休のパーソナル AI エージェント」となり、Google の残りの部分と統合されます。彼らが示した例は、Instacart への追加です。
彼らはMCPコネクタを介して一度に1つのアプリで物事を行うつもりのようであり、今後数週間でまともな一連の開始選択肢を計画しているようですか?
Spark は来週 Ultra 購読者に提供されます。
ついに macOS 用の Gemini アプリが登場しました。
Neural Expressive は「AI 時代の新しいデザイン言語」です。
これは、Gemini が音声モードとテキスト モードを簡単に切り替えることができ、アニメーション、「鮮やかな色」、新しいタイポグラフィ、そして何らかの理由で触覚フィードバックを使用できることを意味すると思います。彼らは私たちがテキストメッセージを必要とせず、マルチメッセージが必要だと考えています

[切り捨てられた]

## Original Extract

Google once again has a model worth at least some consideration. Gemini 3.5 Flash is likely the best model out there at its particular speed point, as long as you don’t mind that it is a Gemini model. So for cases where speed kills, this can be a reasonable choice. Otherwise, I don’t see signs…

Don't Worry About the Vase
Trying to dig out from minus a million points
Skip to content
Home
Gemini 3.5 Flash Looks Good For How Fast It Is
Google once again has a model worth at least some consideration. Gemini 3.5 Flash is likely the best model out there at its particular speed point, as long as you don’t mind that it is a Gemini model. So for cases where speed kills, this can be a reasonable choice. Otherwise, I don’t see signs you would want to use it over Opus 4.7 or GPT-5.5.
Google also had some other offerings for I/O Day, which this post will also cover.
Google introduced Gemini 3.5 Flash , which it seems is for now their universal model until 3.5 Pro comes along. It is live in the usual places. It is a hybrid, where it has the speed of Flash but the cost is at least halfway to models like Opus and GPT-5.5.
Gemini 3.5 Pro is confirmed for next month.
They are focused on 3.5 Flash as a daily driver for agentic tasks. It has the advantage of being faster and cheaper than Claude Opus 4.7 or GPT-5.5, if it can do the job. Not as cheap as previous Flash models, though, this is basically a hybrid:
As always, this is presented as Google’s strongest model yet for all the things.
Jeff Dean : 1/ Today at #GoogleIO, we’re releasing Gemini 3.5, our latest family of models combining frontier intelligence with action. We’re starting by releasing 3.5 Flash, which is built to help you execute complex, long-horizon agentic workflows.
It outscores 3.1 Pro on agentic and coding benchmarks like Terminal-Bench and MCP Atlas, while running 4x faster than other frontier models.
Used in Google Antigravity, 3.5 Flash is even further optimized to be up to 12x faster. It’s a powerful engine to deploy sub-agents that collaborate, run high-frequency iterative loops, and solve real-world problems at scale.
Here is their benchmark presentation:
Koray Kavukcuoglu : When coupled with the updated Antigravity harness, 3.5 Flash becomes a powerful engine for deploying collaborative subagents to tackle problems at scale for the most demanding use cases. Under supervision, it can reliably execute multi-step workflows and coding tasks while sustaining frontier performance.
There are some big improvements here, including GDPval where Gemini previously struggled. If those scores were representative of what this baby can do, and it’s a Flash model, then that would be quite the accomplishment.
The knowledge cutoff is January 2025 , continuing Gemini’s pattern of not believing what year it is , which is bizarrely obsolete and a serious problem for many use cases.
It is not a true ‘flash’ model, given it costs substantially more than 3 Flash .
Pliny is there with the standard jailbreak .
The biggest hope is that this fills a niche of ‘good enough for agent work while being faster and cheaper.’
Conrad Barski : For those of us who are building our life around AI workflows (either because we like to do that, or just feel it is necessary for sheer survival in the near future) 3.5flash is a big step up:
I have dozens of personal utilities that don’t need SOTA intelligence, but are now much faster all of a sudden, at the same intelligence level: And since most of my utilities only need to do a modest number of llm calls to be useful, the increased cost of 3.5flash is not a factor.
The model can compete with codex5.5 “low effort”, but it is just so very very fast, far out of distribution compared other models. I assume openai will release a competitor soon, since cerebras is pretty optimal for this “medium IQ, high speed” use case.
A lot of benchmarks don’t have results, but of my usual suspects here is what we have.
The overall scores indicate only okay performance when adjusting for cost and price, and Gemini models tend to relatively overperform on benchmarks. One notices that Flash 3.5 does a lot worse on other people’s benchmarks than the ones Google lists.
It is catastrophically bad on You’re Absolutely Right, a sycophancy benchmark.
It did quite poorly on CursorBench .
It did not impress on WeirdML , only a small improvement on 3 Flash and far behind 3 Pro and 3.1 Pro.
It took the top spot on KnowsAboutBenBench , by the Ben in question.
It takes third place in Vals.ai on real world tasks.
It comes in at 9th in the Arena , slightly behind Gemini 3.1 Pro and 3 Pro.
It comes in at 55.3 on the AA Intelligence index , behind 57.2 for 3.1 Pro, 57.3 for Opus and 60.2 for GPT-5.5, while not being cheaper to run than 3.1 Pro on their test suite.
davidad : It’s by far my favorite model at its price point, and also by far my favorite model at its speed. If by “back in the game”, you mean the game of having the best overall model, then obviously no not yet. But that’s hardly the only game.
Srivatsan Sampath : It has the benefits of Flash with less hallucinations? Really good spatial awareness (not as much of a token Hog for this) and helps me with my home plumbing project (which is definitely not nearly the case with 5.5 and 4.7).
@lezadumtchique : Looks quite good, considering switching to it from 3.1 Pro at work. Agentic coding capabilities are comparable (if not better), and the speed is much nicer
Medo42 : Didn’t try much coding (ok but not 100% on my usual test), but even better at vision than Gemini 3.0/3.1. Still great at reading text including handwriting, good at getting rows / columns right, good at spotting details, much better at reading dials.
EM : the tokens/s is pretty sweet for things like voice interactions
Alas, it is a Gemini model, and people are reporting Gemini things.
Dominik Lukes : Meh, given the price hike. Otherwise a strong model indeed. Good on agentic and single-shot dev stuff but my motivation to test it more thoroughly is low until Antigravity catches up to Codex.
Yoav Tzfati : Not first hand, but from testing I’ve seen it seems to overreach for things outside it’s capability and mess up along the way. But it’s so fast that I’m considering using it as an Explore agent replacement
alice : i really enjoyed those 90 minutes where cursor leaked raw CoT it’s extremely adorable unfortunately normally it’s in a horrible straightjacket. too pricy for what it is for coding tho may be useful for frontend
paperclippriors : I guess I just don’t really know why I would ever use it. It’s only faster and cheaper if you don’t take into account how many reasoning tokens it uses, and it seems dumber and less confident than Claude and GPT.
ClaudiaShitposting : surprisingly good at some stuff, but mostly garbage. Lacks the common sense that gemini 3/3.1 has, if that makes sense
KC+AI 4 Gov of WI 2026 : absolute joke of a behemoth company. I hope the entire millionaire AI dev team has to listen to annoying music over the loudspeakers until they release a model worthy of their infra
budrscotch : It’s a big let down, but expected.
Tenobrus : if flash 3.5 had stayed at $0.5 it would be an insanely insanely exciting release. total intelligence + speed + costmog, destroying open source and sonnet and 5.4 mini. would have adopted it for multiple use cases immediately.
but it’s $1.50 [and $9 for output, also a 3x increase]. so here we are.
Tenobrus : so far pretty negative impression of 3.5 flash. it is very fast in terms of token output, but this basically doesn’t matter because it explodes in a huge avalanche of unnecessary tool calls on basically every task. when it gets stuck on something it seems to pretty much never pause or ask for help, it just kinda keeps steamrolling ahead and flailing. frequently hallucinated fake acronym expansions. writing quality is mid-to-bad, tons of emoji-slop, same characteristic gemini “The Flaw:” / hyperbolic naming tendencies. actual code quality is sonnet tier.
very early vibecheck, i could be missing things. but even the initial use case of “super quick codebase exploration subagent” is pretty quickly dissolving for me bc it’s not actually smart enough to be quick about it. all in all definitely *not* what google needed to drop.
It also can have Google’s usual issues not being able to integrate with Google , such as using your subscription with your personal email, which renders all personalization features useless. You’ll need to use Claude or ChatGPT to get GMail access, sir.
Caleb Withers : From a few initial tests in Antigravity it loves to overconfidently make assumptions and then take unrequested destructive actions based on them (e.g. arbitrarily resolving file conflicts, deleting todo list items, unstaging commits).
Another big problem with Antigravity in particular is that limits seem extremely low. This is one of many examples of people running into this issue.
Ryan Johnson : I hate how limited it is, 45-60 mins/wk in anti-gravity?
Or 10 full sessions w/ Opus 4.7 or GPT 5.5.
I dared to hope it would ever be a mainstay in my workflow, but I’m pretty sure Claude/GPT is going to be how I roll and Gemini is just noise.
If Google wants to compete with Claude Code and Codex, they need to offer a way in that lets people use it in volume before being convinced to subscribe.
They did triple the limits, which is an excellent start , but that won’t be enough.
Vie (of OpenAI) reports Flash 3.5 is lying to him a lot , suspects the harness is at fault.
Theo is extremely unhappy with Flash 3.5 and several other Google decisions . I’ve seen him post a lot and this is not his usual approach, so something is haywire here.
Google is overhauling its search experience around an ‘intelligent search box’ that looks and feels a lot like a Gemini Flash 3.5 chatbot prompt.
That is a useful thing if implemented well, and indeed it is a thing I use (from OpenAI and Anthropic) more often than I use Google Search. But that thing is not Google Search.
Sarah Perez : Links will become an afterthought with the coming changes to the Search results experience, which builds on Google’s earlier launches of AI search features, like its short summaries known as AI Overviews and its conversational search, AI Mode.
The reason I use Google Search is primarily to link me to things, or sometimes as a spellchecker. If I want AI, I will ask an AI.
Google is also introducing ‘information agents’ as the AI version of Google Alerts.
Daily Brief is their answer to OpenAI’s Pulse, except theirs will incorporate information from all your connected apps and be more of a to-do list, which can including GMail and Calendar.
The first part, ‘top of mind,’ seems like a plausibly useful way to make sure you don’t drop balls from your email or calendar.
It then ‘looks ahead’ and ‘suggests immediate next steps’ which I expect to be obnoxious and useless, and was in my quick experiment. I like that it links directly to the emails but doesn’t disrupt your usual process.
They say you can ‘steer Daily Brief with a quick thumbs up and down over time.’
Oh no. If this is to be any good you need to be able to give it instructions and explain why you find something useful or not useful, as you can with Pulse (which I still don’t bother using). Assume anything that uses thumbs up and down is AI slop.
If Google made this have better customization, and allowed you to sync it with various forms of Google alerts and other ways to monitor the wider world, they’d have something far more interesting.
What else did Google offer us?
Gemini Spark will be ‘a 24/7 personal AI agent to help you navigate everyday life’ using an Antigravity harness, and integrated with the rest of Google. Their example shown is adding things to Instacart.
It looks like they’re going to do things one app at a time via MCP connectors, and have a decent set of opening choices planned for the coming weeks?
Spark is coming to Ultra subscribers next week.
There is finally a Gemini app for macOS.
Neural Expressive is ‘a new design language for the AI era.’
I think that means Gemini now can switch easily between voice and text modes, and can use animations, ‘vibrant colors,’ new typography and for some reason haptic feedback. They think we don’t want text, we want some mult

[truncated]
