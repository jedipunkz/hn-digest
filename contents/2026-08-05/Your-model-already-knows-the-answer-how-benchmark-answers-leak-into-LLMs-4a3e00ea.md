---
source: "https://elman.ai/news/your-model-already-knows-the-answer/"
hn_url: "https://news.ycombinator.com/item?id=49185536"
title: "Your model already knows the answer: how benchmark answers leak into LLMs"
article_title: "Your Model Already Knows the Answer · Elman"
author: "fran-mora"
captured_at: "2026-08-05T17:20:59Z"
capture_tool: "hn-digest"
hn_id: 49185536
score: 4
comments: 0
posted_at: "2026-08-05T16:56:56Z"
tags:
  - hacker-news
  - translated
---

# Your model already knows the answer: how benchmark answers leak into LLMs

- HN: [49185536](https://news.ycombinator.com/item?id=49185536)
- Source: [elman.ai](https://elman.ai/news/your-model-already-knows-the-answer/)
- Score: 4
- Comments: 0
- Posted: 2026-08-05T16:56:56Z

## Translation

タイトル: モデルはすでに答えを知っています: ベンチマークの答えがどのように LLM に漏れるか
記事のタイトル: モデルはすでに答えを知っています · Elman
説明: AI モデルの最良のテストは、答えがわかっている現実世界の問題です。これはカンニングが最も簡単なテストでもあります。答えが漏れる 3 つの方法と、それをキャッチするために使用される 8 つの方法。

記事本文:
問題
アプローチ
製品
チーム
ニュース
お問い合わせ→
エッセイ · 2026年7月30日
モデルはすでに答えを知っています
AI モデルの最良のテストは、答えがわかっている現実世界の問題です。これはカンニングが最も簡単なテストでもあります。 AI モデルはインターネットの大部分を読んでいるので、あなたが尋ねるずっと前にあなたの質問、答え、またはその両方を満たしている可能性があります。高スコアには、2 つのまったく異なるものが隠されています。1 つは、答えに至るまでの推論を行ったモデル、もう 1 つは、すでに読み取ったことを繰り返しているモデルです。これらを区別するのは汚染問題であり、この問題は現場が依存するベンチマークの驚くべき割合を占めています。
臨床で薬が成功するかどうかをモデルがどの程度確実に予測できるかを知りたいと想像してください。新しい治験を何年も待つのではなく、すでに終了した治験でテストします。つまり、後期段階の治験に達した薬剤プログラムを取り出し、それぞれが効果があるかどうかを隠し、モデルに予測を依頼します。紙の上では、それはクリーンなテストで、高速で、実際の結果に基づいてスコア付けされているように見えますが、そうではありません。テストする価値のあるトライアルは、多くの場合、有名なテストであり、その結果は、モデルがトレーニングした教科書、レビュー、ニュース、特許にわたって繰り返し発生します。モデルが薬が成功すると言っているとき、スコアでは、そこに至る道筋を推論したモデルと、その名前を認識して何が起こったかを思い出したモデルを区別することはできません。
臨床試験は私たちの専門分野ですが、罠は一般的です。財務、法律、予測など、結果がすでに公開されているあらゆる分野で、解決済みの出来事から構築されたあらゆるベンチマークに罠がかかります。これらのベンチマークは、正当な理由で公の歴史に基づいて構築されています。実際の結果は、専門家が発明するよりも優れたテストケースです。なぜなら、それは、質問作成者が尋ねようと思ったことだけでなく、真の生物学と薬が失敗する実際の方法を反映しているからです。それらの結果は私たちが望むゴールドスタンダードです

得点するために、そして公になることはまさに彼らを漏洩させるものです。では、何ができるでしょうか？
答えを導き出す 3 つの方法
答えは 3 つのルートでモデルに到達します。この分割は、名前付き分類法ではなく、問題を整理するための私たち独自の方法ですが、各ルートには確立された基礎があります。
入力リークは、テスト時に、モデルがその前のケースに対して読み取るドキュメントを通じて発生します。これらのドキュメントをフィードするか、モデル自体が検索ツールでドキュメントをフェッチするかに関係ありません。決定日以降に書かれた文書はどれも、モデルが予測しようとしている結果そのものを発表することができ、それはモデルに未来を読み取らせるのと同じくらい良いことです (Kaufman et al., 2012)。
ベンチマーク リークは、ベンチマークの後にモデルがリリースされたときに発生します。モデルはベンチマーク セット自体でトレーニングされ、質問と回答の両方を記憶している可能性があります。これは、時間の経過とともにベンチマークが減衰する最も一般的な方法です。 Balloccu らは、GPT-3.5 および GPT-4 のリリースから 1 年以内に GPT-3.5 および GPT-4 に漏洩した 400 万を超えるテストサンプルを追跡しました (Balloccu et al., 2024)。また、他の研究者が小学校の算数の問題セットを一から書き直したところ、いくつかのモデルファミリーのスコアが新しいバージョンで著しく低かった (GSM1k, Zhang et al., 2024)。
結果の漏洩は、結果自体が公の事実であり、モデルがそれに基づいてトレーニングされた場合に発生します。つまり、完了した治験、規制当局の承認、または効果があった薬剤が何千もの論文や申請にわたって繰り返され、モデルはトレーニング中に、事前トレーニングであろうとその後の微調整であろうと、特定のベンチマークとはいかなる接触もせずにそれを吸収します。このモデルは、戦争が終わった年など、広く報道されている他の事実を保持するのと同じ方法で結果を保持するようになります。その結果、次のように書くことができます。

モデルが見たことのない完全に新鮮なテスト セットと、与えられたすべての入力に日付盲検が適用され、モデルは引き続きこの有名な薬が成功したことを認識します。ベンチマークを秘密にしておくことも役に立ちません。答えはそもそもファイルに含まれていないからです。それはすでにモデルの世界に関する一般知識の一部でした。
結果が 10 年先になるまで、新たな結果を待つことは効果的です
あらゆるタイプのリークを防ぐ最もクリーンな方法は、履歴に対するテストを完全に停止し、まだ発生していないイベントのみについてモデルにスコアを付けることです。これにより、リークに利用できる答えがまったく存在しなくなります。ライブ ベンチマークは固定スケジュールで質問を更新し (LiveBench、White et al.、2024)、ForecastBench スコア モデルのような予測ベンチマークは、テスト時にまだ解決されていないイベントをスコア モデルします (ForecastBench、Karger et al.、2024)。 CT Open は、結果が公開される前に臨床試験のモデルにスコアを付けるライブ プラットフォームです。各課題が開始される前に予測が入力され、専門家の注釈と照合して検証された自動化されたパイプラインが、すべての結果について最も早く公開された言及を検索します。そのため、試験は、予測が行われたときにその答えがどこにも見つからなかった場合にのみカウントされます (Wang et al., 2026 ; ct-open.net )。トライアルの結果は一年中発表され、CT オープンではそれに合わせて年に 4 回スコアが発表されます。私たちが重視する決定は、より遅いものです。10 億ドルかかるコミットメントは、それを決定する読み出しの約 10 年前に行われるため、その決定をライブベンチマークに適用すると、最初のスコアは 2036 年に得られます。
そのため、私たちは歴史と、それに伴う結果の漏洩に悩まされています。つまり、知っておく価値のある方法が 8 つ残ります。それぞれが漏れを防止または検出します。しかし、より重要な分割は、どちらであるかです。

3 つのリークにメソッドが対処します。この分野の作業のほとんどはベンチマークのリークを対象としていますが、現実世界の結果から構築されたベンチマークの場合、最も重要なのは結果のリークを目的とした 3 つの手法です。
8 つのメソッド (それぞれが動作するリークごとにグループ化)
入力リーク: ソースを遮断します。
1. 日付ゲート ツールは、決定日以降に公開されるものを拒否する検索および取得ツール (たとえば、日付以前の制限のある Web 検索など) をモデルに与えることを防ぎます。予測システムはすでにこのように機能しており、質問の決定日より前に発行されたニュースに検索を制限しています (Halawi et al., 2024)。これは入力リークを閉じる唯一の動きであり、ソースで機能するため、その後検出するものは何も残りません。それは情報源に記載されている日付と同じくらい正確であるため、実際には、オープンなウェブではなく、信頼できる出版日を持つ学術情報源に頼っています。
ベンチマーク リーク: モデルがファイルを認識したかどうか。
2. ここでは、新しいテスト セットまたはプライベート テスト セットを使用することを防ぐのが賢明な手段です。モデルがトレーニングされていないケースでテストするか、モデルがトレーニングされないようにセットを非公開のままにしておきます。現在、多くの新しいベンチマークはまさに​​このルートを採用しており、各モデルのトレーニング終了後にリリースされた新しい問題を収集しているため、テスト セットをトレーニングすることはできません (LiveCodeBench、Jain et al.、2024)。
この漏洩に対する他の 3 つの方法は、事後的に漏洩を検出しようとします。これらはこの分野で最も発展した部分を占めていますが、公開結果のベンチマークとしては 8 つの中で最も役に立ちません。モデルは質問をした特定のファイルを見なくても結果を完全に知ることができるからです。
3. メンバーシップ推論検出は、各単語に対する信頼度 (対数確率) をチェックすることにより、モデルが特定のテキスト部分を見たかどうかを尋ねます。

モデルは、トレーニングに使用したテキストに対して異常に流暢であるため (Min-K%、Shi et al.、2023)。私たちにとって、それは 2 つの点で弱いです。1 つのファイルは数兆の単語でトレーニングされたモデルの統計をほとんど動かさないため、信号は偶然に近いこと、そしてメソッドはモデルの単語ごとの確率を必要としますが、すべてのプロバイダーがそれを公開しているわけではありません (Anthropic の API は何も返さないため、SaMIA のようなサンプリングベースの回避策を余儀なくされます、Kaneko et al., 2024)。いずれの場合も、モデルがファイルを読み取ったかどうかのみがわかり、答えがわかっているかどうかはわかりません。
4. 摂動検出は、真の答えを維持しながらベンチマークの質問を言い換え、精度が低下するかどうかをチェックします。これは、精度が大幅に低下した場合は、モデルが推論ではなく記憶された言葉と一致していたことを意味するためです。 GSM-Symbolic は、数学の問題の数値を交換することでこれを行い、問題を習得するためだけに見えたモデルを捕捉します (Mirzadeh et al., 2024)。数学の問題のように、言い換えられた質問に自分で解決できる答えが残っている場合に機能します。しかし、実際の治験では解決できることは何もありません。薬剤を変更したり、患者グループを変更したりしても、実施されなかった治験について説明したことになるため、治験がどう終わるかは誰も知りません。
5. ウォーターマークとカナリアは、リリース前にベンチマークに隠されたマーカーが埋め込まれていることを検出します。そのため、後でそれを再現するモデルはファイル上でトレーニングされている必要があります。古典的なものは、 BIG-bench に埋め込まれたカナリア文字列です。ベンチマークを非公開にするのと同様、これはまだ公開していないベンチマークを保護するだけであり、モデルがトレーニングされる前に何年も公開されていた結果については何もしません。
結果リーク: モデルがすでに結果を知っているかどうか。
これらの最後の 3 つは、すでに解決されたケースから構築されたベンチマークにとって重要なものです。これらは結果を目的としているためです。

モデルが特定のファイルをたまたま見たかどうかではなく、直接リークします。
6. カットオフギャップ検出は、トレーニングカットオフ前に決定されたケースとその後に決定されたケースにおけるモデルの精度を比較します。本当のスキルでは、事件が日付のどちら側に該当するかを気にする必要はないため、前と後は偶然に近いという強い結果は、記憶された結果の兆候です。研究者らは、財務予測と出版日に結び付けられた知識に基づいてこの分割を実行しました (Lopez-Lira et al., 2025 ; Roberts et al., 2023)。これは診断であ​​り、設計上の選択ではないため、カットオフ後のケースのスコアリングや入力の日付盲検化とは異なります。これは、両方の部分が同一に構築されている場合、またはメモリではなくフォーマットの変更を測定する場合にのみ機能します。
7. 認識テストは、モデルが証拠ではなく認識に基づいて実行されているかどうかを 3 つのバリエーションで検出します。裏付け資料なしで結果を尋ねる、エンティティの名前を隠して事実を保持する、または名前だけを与えて他には何も与えない。本当に推理するモデルは、名前が隠されているときは安定し、証拠が取り除かれると無力になるはずですが、有名な名前に乗っているモデルはその逆を行います。この方法で名前を隠すことにより、腫瘍学ターゲットの研究では、モデルの上位候補の約 6 つのうち 1 つが変更されましたが、検証されたターゲットの回収率は変化しませんでした (Cuccarese, 2026)。論文の注意点は、遺伝子名には実際の生物学が含まれているため、これによって盲検化された回答がより良くなるわけではないが、盲検化がなければ、どの程度の答えがページ上のデータから得られたもので、どの程度がモデルの名前の記憶から得られたのかを知ることはできないということである。この数字は、それぞれ 1 回の実行による 4 つの指標から得られます。
8. 推論検出の採点では、最終的な答えは無視され、その背後にある議論が判断され、すべての主張が正しいかどうかが確認されます。

決定日以前に存在した情報源に基づくものである。なぜなら、結果のみを記憶するモデルはそれを述べることができるが、それに対する古いケースを構築することができないからである。私たち独自のメソッドがここにあります。領収書テストです。モデルに領収書、つまり決定日時点で入手可能な証拠のみを使用した日付付きのチェック可能なケースを表示するように依頼し、防御できない正解をミスとしてカウントし、正解の頻度とそれを証明できる頻度の間のギャップを追跡し、推測ではなく「まだ知り得ません」と回答させます。日付のある情報源に対する推論トレースの格付けが確立されており (プロセス監視、Lightman et al.、2023)、TimeSPEC は予測を原子的な主張に分割し、決定のどの程度がカットオフ後の情報に基づいているかを報告します (Zhang et al.、2026)。
モデルが完全にクリーンであることを証明できる方法はありません
モデルが完全にクリーンであることを証明する方法はありません。最も近いのは予防です。日付ゲート ツールは、日付以降の文書を締め出すために構築されており、モデルがトレーニングされていないテスト セットはファイルから漏洩することはできないため、これら 2 つの漏洩は単に捕捉するのではなく、計画的に防止することができます。結果リークは閉じることができないものです。重みがまったく役に立たなかったことを証明する方法はなく、重みが役に立たなかったケースを捉えることしかできません。

[切り捨てられた]

## Original Extract

The best test of an AI model is a real world problem with a known answer. It is also the easiest test to cheat. Three ways the answer leaks in, and the eight methods used to catch it.

Problem
Approach
Products
Team
News
Get in touch →
Essay · 30 July 2026
Your Model Already Knows the Answer
The best test of an AI model is a real world problem with a known answer. It is also the easiest test to cheat. An AI model has read much of the internet, so it may have met your question, the answer or both long before you asked. A high score then hides two very different things: a model that reasoned its way to the answer, and one that is repeating something it already read. Telling those apart is the contamination problem , and it sits under a surprising share of the benchmarks the field relies on.
Imagine you want to know how reliably a model can predict whether a drug will succeed in the clinic. Rather than wait years for fresh trials, you test on ones that have already finished: take drug programmes that reached late-stage trials, hide whether each worked, and ask the model to predict. On paper it looks like a clean test, fast and scored on real outcomes, but it is not: the trials worth testing are often the famous ones , and their outcomes recur across the textbooks, reviews, news and patents the model trained on. When the model says a drug will succeed, the score cannot separate a model that reasoned its way there from one that recognised the name and recalled what happened.
Clinical trials are our field, but the trap is general: it springs on any benchmark built from resolved events, in finance, law, forecasting or anywhere the outcome is already public. These benchmarks are built from public history for a good reason: a real outcome is a better test case than any expert can invent, because it reflects true biology and the real ways a drug fails, not just what a question-writer thought to ask. Those outcomes are the gold standard we want to score against, and being public is exactly what makes them leak . So, what can we do?
The three ways the answer gets in
The answer can reach the model by three routes. This split is our own way of organising the problem rather than a named taxonomy, though each route has an established basis.
Input leak happens at test time, through the documents the model reads for the case in front of it, whether you feed those documents in or the model fetches them itself with a search tool. Any document written after the decision date can announce the very outcome the model is meant to predict, which is as good as letting it read the future ( Kaufman et al., 2012 ).
Benchmark leak happens when a model is released after your benchmark and the model might have been trained on the benchmark set itself, memorising both the questions and answers. This is the most common way in which a benchmark decays over time. Balloccu and colleagues traced more than four million test samples leaking into GPT-3.5 and GPT-4 within a year of their release ( Balloccu et al., 2024 ), and when other researchers rewrote a set of grade-school maths problems from scratch, several model families scored markedly lower on the fresh version ( GSM1k, Zhang et al., 2024 ).
Outcome leak happens when the outcome is itself a public fact and the model has trained on it: a completed trial, a regulatory approval or a drug that worked is repeated across thousands of papers and filings, and the model absorbs it during training, whether in pre-training or in the later fine-tuning that follows, with no contact of any kind with your particular benchmark. The model comes to hold the result in the same way that it holds any other widely reported fact, such as the year in which a war ended. As a consequence, you can write a completely fresh test set that the model has never seen and date-blind every input you give it, and the model will still know that this well-known drug succeeded. Keeping the benchmark secret does not help either, because the answer was never contained in your file in the first place; it was already part of the model’s general knowledge of the world.
Waiting for fresh outcomes works, until the outcome is a decade away
The cleanest way to prevent any type of leak is to stop testing on history altogether and to score the model only on events that have not yet happened, so that there is simply no answer available to leak. Live benchmarks refresh their questions on a fixed schedule ( LiveBench, White et al., 2024 ), and forecasting benchmarks like ForecastBench score models on events that have not yet resolved at the time of the test ( ForecastBench, Karger et al., 2024 ). CT Open is a live platform that scores models on clinical trials before their results are public: predictions go in before each challenge opens, and an automated pipeline, validated against expert annotation, searches out the earliest public mention of every outcome, so a trial counts only if its answer was nowhere to be found when the prediction was made ( Wang et al., 2026 ; ct-open.net ). Trial results land all year round, and CT Open scores four times a year to match. The decision we care about is a slower one: the commitment that costs a billion dollars is taken about a decade before the readout that settles it, so put that decision to a live benchmark and your first score arrives in 2036.
So we are stuck with history, and with the outcome leak that comes with it. That leaves eight methods worth knowing. Each one either prevents or detects a leak. The split that matters more, though, is which of the three leaks a method tackles. Most of the field’s work targets the benchmark leak, but for a benchmark built from real world outcomes, the methods that matter most are the three aimed at the outcome leak.
The eight methods, grouped by the leak each one works on
Input leak: shut it at the source.
1. Date-gated tools prevent give the model a search-and-retrieval tool that refuses anything published after the decision date, for example a web search with a before-date limit. Forecasting systems already work this way, restricting retrieval to news published before the question’s decision date ( Halawi et al., 2024 ). This is the one move that closes the input leak, and because it works at the source there is nothing left to detect afterwards. It is only as good as the dates its sources carry, so in practice it leans on scholarly sources with reliable publication dates rather than the open web.
Benchmark leak: whether the model saw your file.
2. Fresh or private test set prevent is the clean move here: test on cases no model has been trained on, or keep the set unpublished so that none can be. A number of newer benchmarks now take exactly this route, collecting fresh problems released after each model’s training cutoff so the test set cannot have been trained on ( LiveCodeBench, Jain et al., 2024 ).
The other three methods for this leak try to detect it after the fact. They make up the most developed part of the field, and yet for a benchmark of public outcomes they are the least useful of the eight, because a model can know an outcome perfectly well without ever having seen the specific file in which you posed the question.
3. Membership inference detect asks whether the model has seen a specific piece of text, by checking how confident it is on each word (log probability of each token in the sequence), since a model is unusually fluent on text it trained on ( Min-K%, Shi et al., 2023 ). For us it is weak on two counts: a single file barely moves the statistics of a model trained on trillions of words, so the signal is close to chance, and the method needs the model’s word-by-word probabilities, which not every provider exposes (Anthropic’s API returns none, forcing sampling-based workarounds like SaMIA, Kaneko et al., 2024 ). Either way it tells you only whether the model read your file, not whether it knows the answer.
4. Perturbation detect rephrases the benchmark question while keeping its true answer, then checks whether accuracy drops, since a large drop means the model was matching remembered wording rather than reasoning. GSM-Symbolic does this by swapping the numbers in maths problems and catches models that had only appeared to master them ( Mirzadeh et al., 2024 ). It works where the reworded question still has an answer you can work out, as with those maths problems. But a real trial gives you nothing to work out: change the drug or the patient group and you have described a trial that was never run, so nobody knows how it would have ended.
5. Watermarks and canaries detect plant a hidden marker in the benchmark before release, so that a model which later reproduces it must have trained on the file; the classic is the canary string buried in BIG-bench . Like keeping a benchmark private, this only protects a benchmark you have not published yet, and does nothing about outcomes that were public for years before any model trained.
Outcome leak: whether the model already knows the result.
These last three are the ones that matter for a benchmark built from cases that have already resolved, because they aim at the outcome leak directly rather than at whether the model happened to see a particular file.
6. The cutoff gap detect compares the model’s accuracy on cases decided before its training cutoff with cases decided after it. Strong before and near-chance after is a signature of memorised outcomes, since real skill should not care which side of a date a case falls on. Researchers have run this split on financial forecasts and on knowledge tied to publication dates ( Lopez-Lira et al., 2025 ; Roberts et al., 2023 ). It is a diagnostic, not a design choice, and so is different from only scoring post-cutoff cases or from date-blinding the inputs; it works only if both halves are built identically, or you measure a change in format instead of memory.
7. Recognition tests detect ask whether the model is running on recognition rather than evidence, in three variants: ask for the outcome with no supporting material, hide the entity’s name but keep the facts, or give only the name and nothing else. A model that truly reasons should be steady when the name is hidden and helpless when the evidence is removed, while a model riding on a famous name does the reverse. Blinding the name this way, in a study of oncology targets, changed about one in six of the model’s top picks while its recovery of validated targets was unchanged ( Cuccarese, 2026 ); the paper’s caveat is that this does not make the blinded answers better, because a gene name carries real biology, but that without blinding you cannot tell how much of an answer came from the data on the page and how much from the model’s memory of the name. That figure comes from four indications with a single run each.
8. Grading the reasoning detect ignores the final answer and judges the argument behind it, checking that every claim rests on a source that existed before the decision date, because a model that only memorised the outcome can state it but cannot build a dated case for it. Our own method sits here, the Receipts Test: we ask the model to show its receipts, a dated and checkable case using only evidence available at the decision date, count a correct answer it cannot defend as a miss, track the gap between how often it is right and how often it can prove it, and let it answer “not yet knowable” instead of guessing. Grading reasoning traces against dated sources is established ( process supervision, Lightman et al., 2023 ), and TimeSPEC breaks a prediction into atomic claims and reports how much of the decision rested on post-cutoff information ( Zhang et al., 2026 ).
No method can prove a model is completely clean
No method can prove that a model is completely clean. Prevention comes closest: a date-gated tool is built to keep post-dated documents out, and a test set no model has trained on cannot leak through your file, so those two leaks can be designed out rather than merely caught. The outcome leak is the one you cannot close. There is no way to demonstrate that the weights gave no help at all, and you can only catch the cases where they

[truncated]
