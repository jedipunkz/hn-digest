---
source: "https://www.lesswrong.com/posts/zhRe3tdBpsZbGCdDK/world-modeling-the-us-vs-anthropic-standoff-on-claude-fable"
hn_url: "https://news.ycombinator.com/item?id=48660665"
title: "World-Modeling the US vs. Anthropic on Claude Fable"
article_title: "World-modeling the US vs. Anthropic Standoff on Claude Fable — LessWrong"
author: "ddp26"
captured_at: "2026-06-24T14:56:20Z"
capture_tool: "hn-digest"
hn_id: 48660665
score: 3
comments: 0
posted_at: "2026-06-24T14:37:56Z"
tags:
  - hacker-news
  - translated
---

# World-Modeling the US vs. Anthropic on Claude Fable

- HN: [48660665](https://news.ycombinator.com/item?id=48660665)
- Source: [www.lesswrong.com](https://www.lesswrong.com/posts/zhRe3tdBpsZbGCdDK/world-modeling-the-us-vs-anthropic-standoff-on-claude-fable)
- Score: 3
- Comments: 0
- Posted: 2026-06-24T14:37:56Z

## Translation

タイトル: クロード・ファブルをめぐる米国対人類の世界モデル化
記事のタイトル: クロード・ファブルをめぐる米国対人類の対立を世界モデル化する — LessWrong
説明: [編集、2026年6月20日: トランプ大統領の6月19日の声明とDCからのその他のマイナーアップデートに基づいて、7月12日から7月7日までの再発表の中央値から予測を修正。]…

記事本文:
米国対クロード・ファブルをめぐる人類の対立の世界モデリング — LessWrong AI 世界モデリング個人ブログ 26
クロード・ファブルをめぐる米国対人類の対立を世界モデル化
[編集、2026年6月20日:トランプ大統領の6月19日の声明およびDCからのその他のマイナーアップデートに基づいて、7月12日から7月7日までの再発表の中央値から予測を修正。]
[編集、2026 年 6 月 23 日: 7 月 7 日から 7 月 9 日まで予測を修正。予測市場は 6 月 18 日からの当初の予測から日付分布に向かって大幅に移動しました。]
私は過去 2 日間、米国がアントロピックにクロード・ファブルの打倒を強制した場合の結果を深く予測することに費やしました。私がこれを行った理由は 2 つあります。(a) 研究のために Fable をいつ取り戻すことができるかを知りたいからです。(b) 結果は米国の AI 規制の大きな前例となるでしょう。
(背景を知りたい人のために、私が見つけた最新かつ包括的な概要は、6 月 17 日の @Zvi の投稿です。ここでは、状況の基本的な詳細を知っているものと仮定します。)
私の世界モデルの結論は興味深いものでしたが、私がこれをここに書いているのは主に認識論的なプロセスと、不合理に陥ることなく大量の AI 研究を管理することについて学んだことのためです。
私の中心的な課題は、
そもそも 6 月 12 日の命令の原因となった出来事には、4 つの異なるバージョンがある可能性を排除できません。
誰がいつアクセスできるか、どのような新しいポリシーが制定されるか、Anthropic が Fable やそのリリース慣行をどのように変えるかまで、予測すべき結果は数多くあります。
ほぼ毎日情報が更新されるため、ほぼすべての再評価が必要になります。
最終的には、無条件予測質問と条件付き予測質問を大規模に組み合わせたものになり、合計 33 個が重要だと考えられました。これは人間、または人間の集団が高品質で行うには多すぎます。

性。そして、もしそうしたとしても、数週間かかり、Fable の使用を計画している人々や米国の AI 政策に取り組んでいる人々にとって役立つ情報を得る機会を逃すことになるでしょう。
言っておきますが、予測市場は主要な結果をカバーしているため、この世界モデリング手法の結果と比較する人間が大勢います。また、最終的にこのプロセスを信頼できない場合でも、可能性の高いタイムラインに頼れる基本情報がいくつかあるということも意味します。
[免責事項: これには私が構築を手伝った FutureSearch 独自の予測機能を使用しました。私たちの評価では、これは各予測質問で手間のかかるフロンティア モデルを単にプロンプトするよりもはるかに正確であることが示されていますが、この記事で説明する世界モデリング プロセスは、月額 20 ドルのチャットボットでも十分に機能するはずです。]
tl;dr : 何が起こっているのかについて私が行った無条件の予測 (結果の上にクロードが生成したグラフィックス):
これは、アメリカ人が Fable を再び利用できるようになったときの CDF につながります。
重要な政策決定の因果関係グラフを構築することは、しばらくの間、予測コミュニティの夢でした。すべてのイベントをリストアップし、それらの間に矢印を描き、入力に基づいて出力を (群衆に) 予測します。
これが予測部分のボトルネックになっていると思われるかもしれません。しかし、実際には、オプション、結果、何が何に影響を与えるかをリストアップして、最初のステップに到達する人をほとんど見たことがありません。 (AI Futures チームが現在これを行っており、有望そうです。私が以前働いていた Metaculus も、 Radiant というサービスでこれを試しています。) 人間の群衆予測を使用する場合の 1 つの課題は、予測を開始する前に、人間がまずこれらのイベントと結果の指標について合意する必要があることです。 (このタイプのモデリングの進歩が遅いことが、@mabramov による最近の LW による予測の削除の主な理由である可能性があります。

。）
この人間的な状況は、実際に私に簡略化されたバージョンを提示しました。原因が不明瞭ながら、ファブルをアメリカ人以外に制限するという唯一の中心的な出来事があった。その後、米国政府が実行できる、そしてアンスロピックが実行できる一連の動きがありました。そして、これらすべての組み合わせに対して予測できる広範な結果はほとんどありませんでした。
つまり、この状況は「通常の予測質問よりも複雑」であり、「完全な因果関係グラフよりも複雑ではない」ということになります。そして、実際の予測にはすべて AI を使用する予定だったので、プロジェクトを 2 日で実行できるようになりました。ニュースが速すぎて時代遅れになることのない結果を得るのに十分な速さです。
以下は、Claude が生成した私のグラフの部分のビューです。
明らかな点の 1 つは、(a) 6 月 12 日の命令の一連の相互に排他的な隠された原因、および (b) 私たちがそれぞれの世界にいるという無条件のナウキャストに非常に敏感であるということです。たとえば、Zvi による上記の記事は、私たちが「政治的影響」シナリオに陥っていることを確信しているようです。これも私の最頻値予測ではありましたが、その確率は全体の 50% 未満でした。
私は、6 月 12 日の命令の背後にあるものについて、他の 3 つのシナリオを拒否するために懸命に努力しました。私はもともと合計で 3 つしか持っていなかったので、最終的に「危険」を「能力の危険」と「外国人の危険」に分割して 4 つを取得しました。これにより作業はさらに難しくなりましたが、これら 4 つのいずれかを除外する自分を納得させる方法を見つけることができず、それらはすべて 10% を超える確率で終了しました。
シナリオの不確実性を受け入れると、不確実性を調査して予測できるようになります。
条件付き予想は大丈夫ですか？
短い答え: わかりません。私は何年も予測精度を研究してきましたが、私の知る限り、条件付き予測精度を評価したことはありません。ましてや、A を使った場合はなおさらです。

は。
「3 月の陸軍省の物語からの政治的影響を主な動機とする 6 月 12 日の命令を条件として、Anthropic は Know-Your-Customer を利用して米国国民に迅速に再開しようとするでしょうか?」のような質問を考えてみましょう。
大きな問題の 1 つは、Futurchy の根本的な欠陥に関する @dynomight の LW 記事です。私たちが尋ねているのは因果関係か相関関係でしょうか？たとえば、「主に政治的影響によって動機付けられた6月12日の命令を条件とする」という記述は、世界情勢に関する他の事柄を暗示しています。それらは考慮されているのでしょうか、それとも「他のすべてのことを同等に保つ」というようなことを言っているのでしょうか?
この世界モデルにとって良いニュースは、すでに起こったことを条件付けすることで、この問題のほとんどが解決されるということです。これは、人々が下す可能性のある将来の決定を条件付けする上で、他の将来の展開と絡み合う、より深刻な問題です。
もう 1 つの良いニュースは、AI は人間ほどこの点について混乱していないようだということです。人間の予報士は、この点で非常につまずきます（私も含めて）。しかし、少なくとも私がこのプロジェクトで使用した AI 予測担当者は、自分の仮定をかなり明確に述べており、理論的根拠を読むと、この場合、実際には 4 つの分割された異なる世界から推論しているようです。 (私はあるスーパー予報官にもこれをチェックしてもらいましたが、彼は条件付き予報が非常に良いように見えることに同意しました。)
今後のプロジェクトでは、適切にスコア付けされた eval を使用してこれを検証する予定です。 FutureSearch にはこれを行うためのデータがあり、既存の過去のキャストを組み合わせて予測し、即座にスコアを付けることができるためです。
それまでは、少なくとも因果関係の懸念については、AI の条件付き予測は Dynomight の警告に従い、(a) 現在、または (b) あなた自身が今すぐに下せる決定を条件としてのみ物事を予測すべきだと思います。 「この仕事に就いたら、

2年後の私の給料はいくらになるでしょうか？」
これがどのように解釈されるかを示すために、最終的な結果予測の 1 つの抜粋を以下に示します。
安全保障上の理論的根拠が実質的に口実であり、その要因は陸軍省の確執とアンスロピック社の差し迫ったIPO（シナリオA3）に関連したホワイトハウスの政治的影響力であるという条件付きでは、この紛争は技術的な修復問題ではなく権力交渉として分析されなければならない。したがって、技術パッチは単独で復元のロックを解除することはできません。この決議により、レバレッジと面子を保つための譲歩が有効になるでしょう...
全体的な品質に関しては、予測ベンチマークからの全体的な証拠 (たとえば、ForecastBench の最新のリーダーボードを参照) は、最大エフォート AI 予測が人間の群衆とほぼ同じくらい優れていることを示唆していると思います。それは十分ではなく、大勢のエリート人間（メタキュラスのプロやグッド・ジャッジメントのスーパー予測者）でさえ、合理的な因果関係グラフを適切に構築できるかどうかは知られていない、と主張する人もいるかもしれない。
しかし、予測を読むと、このアプローチは「有用」の領域に十分入っているように思えます。そして実際、私は、予測の主な有用性はまさにこの種の研究プロジェクトにあると考え始めていますが、これらは群衆予測の制約により人間には実行不可能です。
実際的な大きな問題は、それぞれが複数の AI エージェントによって行われる数十の予測が、ほぼ確実に将来に対する一貫性のない見通しにつながることです。そこから一貫した世界モデルを導き出すにはどうすればよいでしょうか?
AI エージェントのチームからの約 33 の多大な労力の予測を調整する方法のスケッチがいくつかあります。予備的なアプローチの 1 つは、複数の予測にわたって潜在的な請求を見つけて調整し、一貫性が得られるまでそれぞれの予測を更新することでした。これにより、Bench to the Future 評価の精度が向上しました。 (未来

eSearch チームはこれを 6 月 13 日のマニフェストで共有したので、将来の自分へのメモ: 公開された場合に説明するリンクを貼っておきます。) しかし、私はこのプロジェクトに対してそれを十分に信頼していませんでした。そこで私は、人間の予報官としての自分の判断と、状況についての自分の直感を使って、手作業でこれを行いました。
ここで手動とは、長時間のクロード コード セッションを実行し、さまざまな予測を組み合わせて矛盾を指摘し、調整することを意味します。私は、予測結果の確率と日付が一貫していないことと、理論的根拠の推論が一貫していないことの両方を発見しました。
これがどれほどうまくいったのかは、実際にはわかりませんが、33 の主要な予測を抜き取りチェックしたときに矛盾が簡単に見つからなくなったことを除けば。そして、たとえそれができたとしても、議会書簡など、人類とホワイトハウスの交渉が進むにつれて新しい証拠に基づいてそれらの多くを再実行するとき、毎回それらを深くチェックすることはできません。これを扱いやすくするには、AI が本当に必要です。
おそらくこれは、「ゴミが入ってゴミが出てくる」リスクが最も高い場所です。 AI エージェントを使用して他の AI エージェントの出力の欠陥を特定すると、エラーが修正されるのではなく、エラーがさらに悪化するリスクがあります。 『Fable』がいつ公開されるかを知りたい人にとって役立つ一方で、予測を公開するという時間的制約を考慮すると、この場合はこれで十分だと感じました。
(そして、人間の群衆の予測が一貫しているわけではありません。たとえば、ポリマーケットやカルシのさまざまな市場で裁定取引を行うと、多額の利益が得られます。私が言っているのは厳密な裁定取引のことではなく、「世界モデル化」の裁定取引のようなものです。)
これらの AI が複合した可能性のあるエラーを解決する最も簡単な方法は、4 つのシナリオのどれに該当するかを特定することです。これに関する決定的な証拠がいつでも明らかになる可能性があります。ツヴィが上記のリンク先の記事で書いているように、彼はルートニックとアモデイのやりとりについて次のように考えています。

これは、モデルを取り出すことができないことを意味します。」/「それがポイントです。」はほぼ決定的な証拠でしたが、私を少し感動させただけです。4 つのシナリオのうち 2 つを削除しただけでも、ワールド モデリングの多くが削減されることになります。
関連する結果を本当に知っていますか?
これらの注意点を考慮すると、厳しい予測の範囲を 33 という低めに保つのは合理的な判断だったと思います。しかし、1 つの犠牲は、あらゆる種類の結果を大まかなバケツにグループ化する必要があったことです。
各シナリオの完全な分析と、各シナリオの特定の結果の予測を含むグラフは、 https://futuresearch.ai/claude-fable-ban-forecast/ にあります。最後は、やはりクロードによって作成された、次の概要図で終わります。
ここでの問題は、「妥協」とは何を意味するのかということです。考えられる妥協点はいくつかあります。
新しいモデルのセキュリティレビュープロセス (バリエーションあり)
Anthropic (および他の研究所も!?) は KYC を導入しているため、すべてのユーザーは身分証明書を提出する必要があります
アンスロピックがプロジェクト・グラスウィングのデータを米国政府に引き渡す
人間はどういうわけか寓話を硬化します
これらを分割すると、特に中間の交渉の動きに条件を付けると、複雑さが増し、1 人で 2 日間のプロジェクトをはるかに超えます。重要でなかなか進まないトピックについては、チームで作業する

[切り捨てられた]

## Original Extract

[Edit, June 20, 2026: Revised forecast from median re-launch July 12 to July 7, based on Trump's June 19 statements and other minor updates from DC.]…

Login World-modeling the US vs. Anthropic Standoff on Claude Fable — LessWrong AI World Modeling Personal Blog 26
World-modeling the US vs. Anthropic Standoff on Claude Fable
[Edit, June 20, 2026: Revised forecast from median re-launch July 12 to July 7, based on Trump's June 19 statements and other minor updates from DC.]
[Edit, June 23, 2026: Revised forecast from July 7 to July 9. Prediction markets have moved substantially towards the date distribution from the original forecast from June 18.]
I spent the last two days doing a deep dive in forecasting outcomes of the US forcing Anthropic to take down Claude Fable. I did this for two reasons: (a) I want to know when I'll get Fable back for my research, and (b) the outcome will set a major precedent for US AI regulation.
(For those who want background, the most up-to-date and comprehensive summary I could find is @Zvi 's post from June 17 . I'll assume here you know the basic details of the situation.)
My world model's conclusions were interesting, but I'm writing this up here mostly because the epistemic process, and what I learned about managing a large amount of AI research without spinning into unreasonableness.
My central challenges were that
I can't rule out 4 different versions of what happened that caused the the June 12 order in the first place.
There are many outcomes to forecast, from who gets access to when, to what new policies are enacted, to how Anthropic might change Fable or their release practices.
There are informational updates almost every day, requiring a re-evaluation of almost everything.
I ended up with a large combination of unconditional and conditional forecasting questions, in total 33 I consider critical. This is too many for a human, or crowd of humans, to do at high quality. And if they did, it would take weeks and we'd miss the window for the information to be useful to people planning to use Fable, or people working on US AI policy.
It's worth stating, prediction markets cover the major outcomes, so we have a crowd of humans to compare against results of this world-modeling method. It also means, if you ultimately don't trust this process, there is some basic information to fall back to on likely timelines.
[Disclaimer: I used FutureSearch's proprietary forecaster for this, which I help build. Our evals indicate this is much more accurate than just prompting a high-effort frontier model with each forecasting question, but the world-modeling process I lay out in this piece should work well enough with any $20/mo chatbots.]
tl;dr : the unconditional forecasts I did about what's going on (Claude-generated graphics over the results):
and this leads to this CDF for when Fable is available to Americans again:
Building a causal graph of critical policy decisions has been a dream of the forecasting community for a while. You list out all the events, then draw arrows between them, and then (crowd) forecast the outputs conditioned on the inputs.
You might think this is bottlenecked on the forecasting part. But actually I rarely see anyone get to the first step, listing out the options, outcomes, and what influences what. (The AI Futures team is doing this now, and it looks promising. Metaculus, where I used to work, is also trying this in a service called Radiant .) One challenge with using human crowd forecasting is that the humans first have to agree on these events and outcome metrics before they can get to forecasting. (Slow progress in this type of modeling might be a major reason behind this recent LW takedown of forecasting by @mabramov .)
This Anthropic situation actually presented me with a simplified version. There was a single central event, the requirement to restrict Fable to non-Americans, with an unclear cause. Then there were a series of moves the US government could make, and that Anthropic could make. Then there were few broad outcomes that could be forecasted against all these combinations.
So this put this situation as "more complicated than a normal forecasting question" and "less complicated that a full causal graph". And since I was going to use AI for all the actual forecasting, this made the project tractable in 2 days, e.g. fast enough to get results that wouldn't be obsolete by passing news too quickly.
Here's a Claude-generated view of what part of my graph looks like:
One thing that jumps out is that it is very sensitive to (a) the set of mutually exclusive hidden causes of the June 12 order, and (b) the unconditional now-casts that we're in each world. The above article from Zvi, for example, seems to be very sure we're in "political leverage" scenario, and while that was my modal forecast too, it had less than 50% of the total probability.
I really tried hard to reject the other 3 scenarios for what was behind the June 12 order. I originally actually only had 3 total, and ended split "danger" into "capability danger" and "foreigner danger" to get 4. This made this harder but I couldn't find a way to convince myself to rule any of those 4 out, and they all ended with >10% probability.
Once you accept that scenario uncertainty, you get to research and forecasting uncertainty.
Are the conditional forecasts good?
Short answer: I don't know. I have studied forecast accuracy for years now, but AFAICT, nobody has ever evaluated conditional forecast accuracy, let alone when done with AIs.
Take a question like "Conditioned on the June 12 order being primarily motivated by political leverage from the March Department of War saga, will Anthropic use Know-Your-Customer to try to quickly re-launch to US citizens?"
One major issue is @dynomight 's LW piece on Futarchy's fundamental flaw . Are we asking about causation or correlation? e.g., the statement " Conditioned on the June 12 order being primarily motivated by political leverage" implies other things about the state of the world. Are those being taken into account, or are we saying something like "holding all other things equal?"
The good news for this world model is that conditioning on something that has already happened dissolves most of this issue. It's a much more serious problem for conditioning on future decisions people might take, where your'e entangled with other future developments.
Another piece of good news is that AIs seem to be less confused about this than humans. Human forecasters get very tripped up on this (including me). But AI forecasters, at least the one I used for this project, state their assumptions pretty clearly, and reading the rationales, they appear to, in this case, be really reasoning from the four partitioned different worlds. (I had one Superforecaster check this too, and he agreed that the conditional forecasts looked very good to him.)
A future project is to verify this with properly scored evals. FutureSearch has the data to do this, since we can pair existing past-casts, forecast them, and score them immediately.
Until then, at least for the causality worry, I think AI conditional forecasting should follow Dynomight's warning and only forecast things conditional on (a) the present, or (b) decisions that you yourself can make right now, e.g. "If I take this job, what will my salary be in 2 years?"
Here's a snippet of one of the final outcome forecasts, to give a sense of how these read:
Conditional on the assumption that the security rationale is substantially pretextual and the but-for driver is White House political leverage tied to the Department of War feud and Anthropic's impending IPO (Scenario A3), this dispute must be analyzed as a power negotiation rather than a technical remediation problem. Consequently, technical patches will not independently unlock restoration; the resolution will turn on leverage and face-saving concessions...
As for overall quality, I think the overall evidence from forecasting benchmarks (e.g. see ForecastBench's most up-to-date leaderboard ) suggests max effort AI forecasts are about as good as a crowd of humans. One could argue that's not good enough, and that even crowds of elite humans (Metaculus pros, Good Judgment superforecasters) aren't known to be able to properly build a reasonable causal graph.
But reading the forecasts makes me think this approach is well into the "useful" territory. And in fact, I'm starting to think the main usefulness of forecasting is in exactly these types of research projects, which are infeasible to humans due to the constraints of crowd forecasting.
The big practical issue is that dozens of forecasts, each done by multiple AI agents, will almost certainly lead to an inconsistent view of the future. How can you get a coherent world model out of it?
I have some sketches for ways to reconcile ~33 high effort forecasts from teams of AI agents. One preliminary approach involved finding and refining latent claims across forecasts, and then updating the forecasts on each one until they are consistent. This increased accuracy on Bench to the Future evals. (The FutureSearch team shared this at Manifest on June 13, so note to future self: link that talk if it becomes publicly available.) But I didn't trust it enough for this project. So I did this by hand, using my own judgment as a human forecaster and my own intuitions about the situation.
By hand here meant running lengthy Claude Code sessions, asking it to pair of various forecasts and point out inconsistencies, and then making tweaks. I found both by inconsistent forecast outcomes probabilities and dates, as well as inconsistent lines of reasoning in the rationales.
I can't really say how well this worked, other than I can no longer easily see any inconsistencies when I spot check the 33 core forecasts. And even if I could, when I re-run many of them based on new evidence as Anthropic and White House negotiations proceeded, such as the Congressional letter, I can't deeply check them each time. AI is really necessary here to make this tractable.
This is probably where the "garbage in, garbage out" risk is highest. Using AI agents to identify flaws in the outputs of other AI agents risks compounding errors rather than fixing them. Given the time constraints of publishing forecasts while they would still be useful to people who want to know when Fable will come out, this felt good enough in this case.
(And it's not like human crowd forecasts are consistent. There is a lot of money to be made arbitraging various markets across Polymarket and Kalshi, for example. I don't mean strict arbitrage, more like "world modeling" arbitrage.)
The easiest way to resolve these potentially AI-compounded errors is finding out which of the 4 scenarios we're in. Decisive evidence could come out on this at any day. As Zvi wrote in the linked piece above, he thinks the The Lutnick - Amodei exchange, "This means we can't have the model out" / "That's the point", was nearly decisive evidence, but it only moved me a little bit. Even removing 2 of the 4 scenarios would reduce a lot of the world modeling.
Are we sure we know the relevant outcomes?
Given these caveats, I think keeping the scope low, at 33 hard forecasts, was a reasonable call. But one sacrifice is that I had to group all sorts of outcomes into coarse buckets.
The full analysis of each scenario, and the charts with forecasts of specific outcomes in each scenario, are in https://futuresearch.ai/claude-fable-ban-forecast/ . It concludes with this summary graphic, again made by Claude:
The problem here is, what does "Compromise" mean? There are a number of possible compromises:
A new model security review process (of which there are variants)
Anthropic (and other labs!?) implement KYC, so all users have to submit identification
Anthropic hands over Project Glasswing data to the US government
Anthropic hardens Fable somehow
Splitting these out, especially conditioning on intermediate negotiating moves, adds enough complexity that this goes way beyond a 2-day, 1-person project. For important, slow-enough-moving topics, having a team work

[truncated]
