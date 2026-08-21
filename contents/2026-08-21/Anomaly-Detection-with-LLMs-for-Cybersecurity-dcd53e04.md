---
source: "https://ngrislain.github.io/blog/2026-8-20-language-models-are-anomaly-detectors/"
hn_url: "https://news.ycombinator.com/item?id=49387421"
title: "Anomaly Detection with LLMs for Cybersecurity"
article_title: "Language Models Are Anomaly Detectors | NGrislain"
image: "https://ngrislain.github.io/static/blog/genai-anomaly-detection/thumbnail.png"
author: "ngrislain"
captured_at: "2026-08-21T13:37:56Z"
capture_tool: "hn-digest"
hn_id: 49387421
score: 1
comments: 0
posted_at: "2026-08-21T12:57:26Z"
tags:
  - hacker-news
  - translated
---

# Anomaly Detection with LLMs for Cybersecurity

- HN: [49387421](https://news.ycombinator.com/item?id=49387421)
- Source: [ngrislain.github.io](https://ngrislain.github.io/blog/2026-8-20-language-models-are-anomaly-detectors/)
- Score: 1
- Comments: 0
- Posted: 2026-08-21T12:57:26Z

## Translation

タイトル: サイバーセキュリティのための LLM による異常検出
記事のタイトル: 言語モデルは異常検出器である | Nグリスレイン
説明: Datadog でセキュリティ検出に取り組んでいます。私が構築してきたモデル Mambark は、言語モデルがテキストを読み取るのと同じ方法で監査ログを読み取り、すべてのイベントをその驚き度によってスコア付けします。今日では小型モデルです

記事本文:
言語モデルは異常検出器です
私は Datadog でセキュリティ検出に取り組んでいます。私が構築してきたモデル Mambark は、言語モデルがテキストを読み取るのと同じ方法で監査ログを読み取り、すべてのイベントをその驚き度によってスコア付けします。これは今日の標準からすると小規模なモデルで、パラメータ数は約 9,700 万で、文の代わりに数千億のセキュリティ イベントを調べます。
生成モデルは、物を生産するための機械として話題になります。この投稿の主張は、何も追加も微調整もしていない同じモデルが、これまでに構築された最も一般的な異常検出器でもあり、これによって大規模な検出方法がすでに変化しつつあるということです。このメカニズムこそが、彼らに文字を書かせるものであることが判明しました。
基本的な考え方は古く、恥ずかしいほど単純です。私は、セキュリティ テレメトリや公開できないものから離れて、これを単独で確認したかったので、平文の英語テキストで実行する小さなオープンソース プロジェクトを構築しました。この投稿はそれを示しています。
次に何が起こるかを予測するモデルは、すでに異常検出器です。
接頭辞を付けます。次のシンボルにわたる分布が得られます。次に実際に次に何が起こるかを見て、それが割り当てた確率を読み取ってください。モデルが 0.4 と表示した場合は、何も起こりません。 0.0001 と表示されていれば、何かが起こっています。
負のログを取得すると、次の数値が得られます。
s_i = -\log P\!\left(x_i \mid x_1, \ldots, x_{i-1}\right)
それが異常スコアです。これは nats で測定され、決して負になることはなく、合計されます。ウィンドウ全体のサプライズは、そのシンボルのサプライズの合計であり、まさにそのウィンドウの負の対数尤度です。ラベルも、既知の攻撃のリストも、悪さがどのようなものかを説明するルールもありません。唯一の材料は、通常次に来るもののモデルです。
その配布には価値がある

これらのモデルを作成するのと同じ機械であるため、直接問題はありません。
プレフィックスを 1 回前方パスすると、語彙全体に 1 つの分布が生成され、このモデルでは 248,320 個のトークンが生成されます。生成するには、そのディストリビューションからトークンをサンプリングして追加し、パスを再度実行します。スコアを付けるには、サンプリングをスキップし、代わりに実際に次に来たトークンを調べ、そのスロットにある確率のマイナス対数を取ります。同じパス、同じ分布、それに対して 2 つの異なる質問がありました。検出は、サンプリング ステップをルックアップに置き換えた生成です。
3 つのパネルは、この投稿の後半で使用される汚染されたテキストの 3 つのポイントにおける実際の出力です。一般的な英語では、発展するコンピューターサイエンスの後、モデルのお気に入りはアルゴリズム 30.6%、インテリジェント 18.5% です。テキストには と書かれており、4 番目の選択肢は 6.8% で、費用は 2.69 nat です。何も起こりませんでした。定義された目標を達成した後の接続時。 、段落区切り、AI、または The ; が必要です。 E を取得し、248,320 件中 602 位、確率 0.002%、スコアは 11.09 nats に跳ね上がります。何かが起こりました。
3番目のパネルは私が予想していなかったものです。バスク文の終わりまでに、モデルの継続可能性が 3 番目に高いのは E で 10.4% であり、Ez と B はリストの下位にあります。バスク語のさらなる発展が期待されています。約 200 文字で、フィッティングや再トレーニングを行わずに、独自の正常概念を変更しました。その変化はディストリビューションから直接読み取ることができます。
そのどれもが Qwen や言語に特有のものではありません。次のシンボルに確率を与えるモデルはすべて、この方法で読み取ることができます。したがって、興味深い問題は、不意を突いて採点するかどうかではなく、どのモデルを選択するかです。なぜなら、何を選択するかが「通常」を定義するものであり、したがって何が異常としてカウントされるかであるからです。
知るための 2 つの方法

次に何が来るのか
最も古い答えは数えることです。文字 N グラム モデルは、トレーニング テキスト内の各継続が各コンテキストにどのくらい続くかをカウントすることによって P(c \mid \text{previous } n-1 \text{characters}) を推定します。新しい文字をスコアするには、そのコンテキストを調べて頻度を読み取ります。この投稿のモデルは、内挿平滑化を備えた 5 グラムであるため、次数 5 の推定値と次数 4、次数 3、次数 2、ユニグラムおよび均一項が混合されています。見たことのないキャラクターも含めて、確率がゼロになることはありません。
安くて透明で全部読めます。また、非常に重要であることが判明した 2 つの特性もあります。
通常と呼びたい特定の種類のテキストにそれを適合させる必要があります。そしてその記憶は4文字です。それ以前のものはすべてなくなってしまいました。 atio の後に何が来るかを決定する 5 グラムは、テキストが英語であることも、機械学習に関するものであることも、300 文字前に同じ略語が出現したことも知りません。
言語モデルは接頭辞全体を条件とします。シーケンスの先頭以降のすべての文字がコンテキスト ウィンドウ内に表示されます。そして、適切なステップはありません。事前トレーニングでは、英語、フランス語、バスク語、ウィキペディアの散文や他のほとんどの文章がすでにカバーされています。ここで使用したモデルは Qwen3.5-0.8B-Base で、トレーニング後の分布が歪められ、尤度推定値が悪化するため、チャット モデルではなくベース バリアントです。
興味深いのはその結果です。ベースラインは事前に適合されるのではなく、スコアリングされるシーケンスの最初からその場で組み立てられます。最初の文は、これが英語であること、これが百科事典の登録簿であること、これが人工知能に関するものであることをモデルに伝えます。その後のすべてはそれに基づいて判断されます。 1 つのモデル、通常のいずれか。
ウエバマ

rket は最初のオプションで実行されます
これは抽象的な区別ではありません。製品カテゴリー全体の形状です。
Exabeam のドキュメントは、その行動分析がどのように機能するかについて新鮮に直接的に説明しています。 「私たちの異常検出は、ネットワーク エンティティの動作の統計的プロファイリングに依存しています。」 「統計的プロファイリングはヒストグラム頻度に基づいています。」 「確率分布はヒストグラムを使用してモデル化されます。」モデル タイプは 3 つあります。「数値、ホスト名、ユーザー名などの重要な文字列をモデル化する」カテゴリカル、クラスター化された数値、および曜日の数値です。モデルは、機能、スコープ、エージング ウィンドウ、およびconfidence_factor>=0.8などの収束フィルターを含む構成ファイルで宣言されます。
それを N グラムとして読み取ると、マッピングは正確になります。特徴はコンテキストです。その値のヒストグラムが度数表です。スコープによって、どの法線を当てはめるかが決まります。つまり、このユーザー、このピア グループ、組織全体です。収束フィルターは、「スコアを獲得するのに十分な内容を確認したか」テストです。老朽化した窓を新しく取り付けています。導入ガイドでは通常、アラートが有効になるまでに 30 ～ 60 日間のベースラインを設定することが推奨されています。
それは機能し、10年間UEBAを支えてきました。ただし、両方の N-gram コストをエンティティごとに永久に支払うことになります。すべてのユーザー、ホスト、サービスには、関心のあるすべての機能に合わせた独自のヒストグラムが必要です。新入社員にはベースラインがありません。チームがツールを切り替えると、ベースラインが無効になります。そして、コンテキストは構築上短いままです。アリスがサインインするホストのヒストグラムは、アリスが過去 10 分間に何をしたかについては何も知りません。
だからこそ、UEBA につきまとう苦情はいつも同じ、「うるさすぎる」というものだ。間違っているというよりも単に異常なだけのアラートが多すぎて、アナリストがアラートを読むのをやめてしまいます。カテゴリが発明したすべての緩和策、コンバージョン

モデルのスコア付けが許可される前の緊急しきい値、ピア グループ化、多くの弱いシグナルにわたるリスクの集約は、事後にその下限を抑えるために存在します。この実験ではそれをミニチュアで再現するため、その不満を心に留めておいてください。
新しい世代では 2 番目のオプションが採用されます。 Sweet Security はイベントをセッションに集約し、LLM にセッションを読み取らせ、検出されたそれぞれに悪意のある、疑わしい、または悪い慣行にラベルを付けます。 Manbark は、次のイベント予測で事前トレーニングされたモデルの下で、負の対数尤度によってすべてのイベントをスコア付けします。異なる製品でも同じ動き: 適合したエンティティごとのテーブルを一般的なシーケンスの事前トレーニング済みモデルに置き換えます。
ここから先はすべて英語の散文に関するものであるため、転送可能にするマッピングを特定する価値があります。
エンティティはライターです。ユーザー、ホスト、サービス アカウント、エージェント: それぞれがイベントのストリームを発行し、そのストリームが書き込むテキストです。各イベントはトークンです。その作家にとっての「普通」が意味するものは、散文にとっての「普通」が意味するものであり、それは同じ3つの層を持っています。語彙: どの API を呼び出すか、どのホストに到達するか、どのリージョンに表示されるか。構文: どのイベントがどの順序で続く傾向があるか。そしてレジスター: 活発か定常か、労働時間か午前 3 時か、簡潔か冗長か。
そのマッピングの下で​​は、UEBA ヒストグラムは、1 人のライターの 1 つの特徴にわたる度数テーブルです。アリスがよく使う文字を数えます。これはかなりの数であり、アリスがこれまでに使用したことのない文字を使用したことを示しますが、アリスがどの言語で書いたか、一段落前に何を言ったかはわかりません。
異常とは、文書の残りの部分と同様に読み取れない箇所のことです。そして、以下の実験が中心に構築されているのは、英語の段落の途中に接続された外国語の文です。

セキュリティ上最も重要なケース: 他の誰かが途中で書き始めました。認証情報の盗難、セッションのハイジャック、内部関係者による役割の外への侵入、エージェントの台本からの逸脱。文書にはアリスの名前が記載されたままですが、散文は段落の途中で言葉を変え、その後元に戻ります。
このことを念頭に置いて、次の図を読んでください。英国人はアリスで一週間を過ごしています。バスク語の文は彼女のキーボードの前にいる別の誰かです。フランス語のものは、彼女のメールを読んで、彼女のように聞こえるように努めている別の誰かです。
以下のすべては、そのプロジェクト内の 1 つのスクリプトからのものであり、ウィキペディアから独自のデータを取得して再現できるようにします。
5 グラムは、「機械学習」、「統計」、「コンピュータ サイエンス」、「数学」の検索の背後にある英語の記事に適用されます。これは 193,003 文字であり、スコアの対象となるテキストと同じ言語、ほぼ同じ主題を使用しているため、意図的に寛大になっています。
スコアリングされるテキストは、人工知能に関する英語記事の最初の 1,200 文字です。それはフィッティングコーパスには含まれていませんが、それに近いものはほとんどありません。
次に、汚染されたコピーが 2 つあり、それぞれ最初の文の後に 1 つの文が接続されています。 1 つはバスク語です。eu.wikipedia の記事 Euskara から、バスク語が孤立言語であることを説明する一節です。もう 1 つはフランス語で、fr.wikipedia のラング バスク語から来ており、同じことを主張しています。これらは同じ位置にある別々のコピーになるため、両方とも同じ背景に対して測定されます。
Qwen は文字ではなくトークンをスコアリングするため、各トークンの対数確率はその文字数で割られ、トークン間で共有されます。これにより、両方のモデルが同じ軸に置かれます。
すべてのキャラクターは、すべての図のすべてのパネルで共有される 1 つのスケールで、nats での驚きによって影付けされます: 0.7 nats が残ります。

透明、7.7 ナット以上は固体です。対数確率はモデル間で同等であり、nat は nat であるため、これは公平な比較となります。モデルごとに正規化されるものはなく、パネルごとに正規化されるものはありません。
各パネルの下には、0 から 8 nat までの同じ数値が 1 文字につき 1 ポイントのトレースとして再び表示され、パネル自体の平均が破線で表示されます。平滑化されていません。移動平均を使用すると、しきい値を設定するのが簡単になりますが、言語モデルの境界スパイクがプラトーに広がり、注目に値するものが破壊されてしまいます。
誰も触れていないテキストから開始し、両方のモデルで読み取られます。
違いを見逃すのは難しいです。 193,003 文字のドメイン内フィッティングの後でも、5 グラムは通常の英語に 1 文字あたり 1.29 ナットを費やします。 Qwen は 0.38 を費やし、何も装着されていませんでした。 N-gram の文字の 4.5% は 3 nat 以上かかりますが、Qwen の場合は 0.3% です。それがノイズフロアであり、アナリストが真実を見つける前に説明しなければならないものです。矢印は、最も大音量の 2 つの誤報 (Web 検索で 5.1 nats 、仮想で 5.1 nats) を指しています。どちらも人工知能に関する記事の平易な英語です。
斑点はランダムではありません。単語の始まりには N グラムのコストが 1.64 ナットかかります

[切り捨てられた]

## Original Extract

I work on security detection at Datadog. The model I have been building, Mambark, reads audit logs the way a language model reads text and scores every event by how surprising it is. It is a small model by today

Language Models Are Anomaly Detectors
I work on security detection at Datadog . The model I have been building, Mambark , reads audit logs the way a language model reads text and scores every event by how surprising it is . It is a small model by today's standards, about 97 million parameters, and it looks at hundreds of billions of security events instead of sentences.
Generative models get talked about as machines for producing things. The claim of this post is that the same models, with nothing added and nothing fine-tuned, are also the most general anomaly detectors anyone has built , and that this is already changing how detection is done at scale. The mechanism turns out to be the very thing that makes them write.
The idea underneath is old and almost embarrassingly simple. I wanted to see it on its own, away from security telemetry and away from anything I cannot publish, so I built a small open-source project that runs it on plain English text. This post is what it shows.
Any model that predicts what comes next is already an anomaly detector.
Hand it a prefix. It gives you a distribution over the next symbol. Now look at what actually came next and read off the probability it assigned. If the model said 0.4, nothing happened. If it said 0.0001, something happened.
Take the negative log and you have a number:
s_i = -\log P\!\left(x_i \mid x_1, \ldots, x_{i-1}\right)
That is the anomaly score. It is measured in nats, it is never negative, and it adds up: the surprise of a whole window is the sum of the surprises of its symbols, which is exactly the negative log-likelihood of that window. No labels, no list of known attacks, no rule describing what bad looks like. The only ingredient is a model of what usually comes next.
That distribution is worth looking at directly, because it is the same machinery that makes these models write:
One forward pass over the prefix produces one distribution over the entire vocabulary, 248,320 tokens for this model. To generate, you sample a token from that distribution, append it, and run the pass again. To score, you skip the sampling and instead look up the token that actually came next, then take minus log of the probability sitting in that slot. Same pass, same distribution, two different questions asked of it. Detection is generation with the sampling step replaced by a lookup.
The three panels are real output at three points in the contaminated text used later in this post. In ordinary English, after computer science that develops , the model's favourites are algorithms at 30.6% and intelligent at 18.5%; the text said and , its fourth choice at 6.8%, which costs 2.69 nats. Nothing happened. At the splice, after achieving defined goals. , it expects a paragraph break or AI or The ; it gets E , ranked 602nd out of 248,320, probability 0.002%, and the score jumps to 11.09 nats. Something happened.
The third panel is the one I did not expect. By the end of the Basque sentence the model's third most likely continuation is E at 10.4%, with Ez and B further down the list. It is expecting more Basque. It has moved its own notion of normal, in about two hundred characters, with no fitting and no retraining, and you can read that shift straight off the distribution.
None of that is specific to Qwen, or to language. Any model that puts a probability on the next symbol can be read this way. So the interesting question is not whether to score by surprise, it is which model you ask, because whatever you pick is what defines "usually" and therefore what counts as an anomaly.
Two ways to know what comes next
The oldest answer is to count. A character n-gram model estimates P(c \mid \text{previous } n-1 \text{ characters}) by counting how often each continuation followed each context in some training text. To score a new character, you look up its context and read off the frequency. The model in this post is a 5-gram with interpolated smoothing, so it mixes the order-5 estimate with order-4, order-3, order-2, the unigram and a uniform term. Nothing ever gets probability zero, including characters it has never seen.
It is cheap, transparent and you can read the whole thing. It also has two properties that turn out to matter a lot.
You have to fit it, on the specific kind of text you are willing to call normal. And its memory is four characters. Everything before that is gone. A 5-gram deciding what comes after atio does not know the text is in English, does not know it is about machine learning, and does not know that the same abbreviation appeared three hundred characters ago.
A language model conditions on the entire prefix. Every character since the beginning of the sequence is in the context window. And there is no fitting step: pretraining already covered English, French, Basque and Wikipedia prose, along with most other things. The model I used here is Qwen3.5-0.8B-Base , the base variant rather than the chat one, because post-training skews the distribution and makes it a worse likelihood estimator.
The consequence is the interesting part. The baseline is not fitted in advance, it is assembled on the fly out of the beginning of the very sequence being scored. The first sentence tells the model this is English, this is encyclopedic register, this is about artificial intelligence. Everything after is judged against that. One model, any normal.
The UEBA market runs on the first option
This is not an abstract distinction. It is the shape of a whole product category.
Exabeam's documentation is refreshingly direct about how its behavioral analytics works. "Our anomaly detection relies on statistical profiling of network entity behavior." "The statistical profiling is histogram frequency based." "Probability distributions are modeled using histograms." There are three model types: categorical, which "models a string with significance: number, host name, username", numerical clustered, and numerical time-of-week. Models are declared in a config file with a feature, a scope, an aging window and a convergence filter such as confidence_factor>=0.8 .
Read that as an n-gram and the mapping is exact. The feature is the context. The histogram over its values is the frequency table. The scope decides which normal you are fitting: this user, this peer group, the whole organisation. The convergence filter is the "have I seen enough to score yet" test. The aging window is refitting. Deployment guides commonly suggest thirty to sixty days of baselining before alerting goes live.
It works, and it has carried UEBA for a decade. But you pay both n-gram costs, per entity, forever. Every user, host and service needs its own fitted histogram for every feature you care about. A new employee has no baseline. A team that switches tools invalidates its baseline. And the context stays short by construction: a histogram of which hosts Alice signs into knows nothing about what Alice did in the previous ten minutes.
Which is why the complaint that follows UEBA around is always the same one: it is too noisy. Too many alerts that are merely unusual rather than wrong, and an analyst who stops reading them. Every mitigation the category has invented, convergence thresholds before a model is allowed to score, peer grouping, risk aggregation over many weak signals, exists to hold that floor down after the fact. Keep that complaint in mind, because the experiment reproduces it in miniature.
The newer generation takes the second option. Sweet Security aggregates events into sessions and has an LLM read the session , then labels each finding malicious, suspicious or bad practice. Mambark scores every event by negative log-likelihood under a model pretrained on next-event prediction. Different products, same move: replace the fitted per-entity table with a pretrained model of sequences in general.
Everything from here on is about English prose, so it is worth pinning down the mapping that makes it transferable.
An entity is a writer. A user, a host, a service account, an agent: each one emits a stream of events, and that stream is the text it writes. Each event is a token. What "normal" means for that writer is what normal means for prose, and it has the same three layers. A vocabulary: which APIs it calls, which hosts it reaches, which regions it appears in. A syntax: which event tends to follow which, in what order. And a register: bursty or steady, working hours or three in the morning, terse or verbose.
Under that mapping a UEBA histogram is a frequency table over one feature of one writer. It counts the letters Alice tends to use. It is a good count, and it will tell you when Alice uses a letter she has never used before, but it does not know what language she writes in or what she was saying a paragraph ago.
An anomaly is then a passage that does not read like the rest of the document. And the specific thing the experiment below is built around, a foreign sentence spliced into the middle of an English paragraph, is the case that matters most in security: somebody else started writing halfway through . Stolen credentials, a hijacked session, an insider stepping outside their role, an agent going off script. The document keeps Alice's name on it while the prose changes language mid-paragraph, and then changes back.
Read the figures that follow with that in mind. The English is Alice going about her week. The Basque sentence is somebody else at her keyboard. The French one is somebody else who has read her email and is trying to sound like her.
Everything below comes from one script in that project, which fetches its own data from Wikipedia so you can reproduce it.
The 5-gram is fitted on the English articles behind the searches "Machine learning", "Statistics", "Computer science" and "Mathematics". That is 193,003 characters, and it is deliberately generous: the same language and roughly the same subject as the text it will score.
The text being scored is the first 1,200 characters of the English article on artificial intelligence. It is not in the fitting corpus, but it could hardly be closer to it.
Then two contaminated copies, each with a single sentence spliced in after the first sentence. One is Basque, from the eu.wikipedia article Euskara , the passage explaining that Basque is a language isolate. The other is French, from fr.wikipedia 's Langue basque , making the same point. They go into separate copies at the same position so both are measured against an identical background.
Qwen scores tokens, not characters, so each token's log-probability is divided by the number of characters it spans and shared out across them. That puts both models on the same axis.
Every character is shaded by its surprise in nats, on one scale shared by every panel of every figure: 0.7 nats is left transparent, 7.7 nats and above is solid. Log-probabilities are comparable across models, a nat is a nat, so this is a fair side by side. Nothing is normalized per model and nothing is normalized per panel.
Under each panel the same numbers appear again as a trace, one point per character, from 0 to 8 nats, with the panel's own average as a dashed line. It is not smoothed. A moving average would be easier to threshold, but it would also spread the language model's boundary spikes into a plateau and destroy the thing worth looking at.
Start with text nobody has touched, read by both models:
The difference is hard to miss. After 193,003 characters of in-domain fitting, the 5-gram still spends 1.29 nats per character on ordinary English. Qwen spends 0.38, having been fitted on nothing at all. 4.5% of the n-gram's characters cost more than 3 nats, against 0.3% for Qwen. That is the noise floor, and it is what an analyst has to explain away before finding anything real. The arrows point at two of the loudest false alarms: 5.1 nats on web search , 5.1 nats on virtual . Both are plain English in an article about artificial intelligence.
The speckle is not random. Word beginnings cost the n-gram 1.64 nats

[truncated]
