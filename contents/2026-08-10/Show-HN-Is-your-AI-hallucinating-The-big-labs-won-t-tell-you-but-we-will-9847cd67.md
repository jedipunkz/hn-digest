---
source: "https://www.telluvian.ai/blog/is-your-ai-hallucinating"
hn_url: "https://news.ycombinator.com/item?id=49246077"
title: "Show HN: Is your AI hallucinating? The big labs won't tell you but we will"
article_title: "Is Your AI Hallucinating? — Telluvian"
author: "galois123"
captured_at: "2026-08-10T16:43:10Z"
capture_tool: "hn-digest"
hn_id: 49246077
score: 2
comments: 0
posted_at: "2026-08-10T16:36:49Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Is your AI hallucinating? The big labs won't tell you but we will

- HN: [49246077](https://news.ycombinator.com/item?id=49246077)
- Source: [www.telluvian.ai](https://www.telluvian.ai/blog/is-your-ai-hallucinating)
- Score: 2
- Comments: 0
- Posted: 2026-08-10T16:36:49Z

## Translation

タイトル: HN を表示: あなたの AI は幻覚を見ていますか?大きな研究所は教えてくれませんが、私たちは教えます
記事タイトル: あなたの AI は幻覚を見ていますか? — テルヴィアン
説明: ホワイトボックス プロキシ モデルがトークンごとの幻覚スコアを提供する方法 (正確なコストを含む)。

記事本文:
あなたのAIは幻覚を見ていますか？ — テルヴィアン テルヴィアン ブログ
フロンティア言語モデルが生成するすべてのトークンの幻覚スコアを返す API を出荷します。それがどのように機能するかを説明する前に、これについて最も重要なことを説明します。返されるパーセンテージ スコアは、実際にはトークンが間違っている確率ではありません。
明らかな問題は、フロンティアモデルは、彼らの思考を明らかにするために調査する必要がある隠れた状態を明らかにしないことです。うーん...
GPT またはクロードをプローブすることはできません。これらの API はテキストのみを返します。読み取るレイヤーも検査する残留ストリームもありません。いくらプロンプトを出してもそれを変更することはありません。この空間におけるすべてのテクニックは、同じ壁を迂回する必要があります。
私たちの答えは、ホワイトボックス プロキシ モデルです。私たちはエンドツーエンドで制御できるオープンウェイト モデルを実行しています。これは、すべてのレイヤーでそのアクティベーションに完全にアクセスできることを意味します。プローブはこのプロキシ モデルでトレーニングされ、このプロキシ モデルから読み取ります。
明確にしておきますが、私たちの API は実際には、呼び出したモデルではなく、プロキシ モデルに関する情報を返します。
私たちは、GPT-5.4 Nano から数十万の入力/出力ペアのデータセットを取得しました。これには、信頼できる主張と幻覚的な主張の両方が含まれています。データセット全体で平均すると、ナノ氏の主張の 3% が幻覚だった。次に、Gemma-4 を介して Nano の出力を再生し、Gemma-4 の隠れた状態のラベルに対してプローブをトレーニングしました。
私たちは、目に見えない検証セットでプローブを採点し、プローブの評決が裁判官 LLM の評決と一致した場合に、事件を正確であるとカウントしました。合意は95%を超えました。
それでは、API は実際に何を返すのでしょうか?
私たちのスコアで測定されるのは、呼び出されたモデルによって生成された出力の生成中にプロキシ モデルが混乱する可能性です。スコアが高いことは、プロキシ モデルがそのトークンを処理したときに、その内部状態が密接に関係していることを意味します。

混乱したときに通常示す内部状態を組み込みます。
結論として、私たちの API はブラック ボックス フロンティア モデルで動作します。ベース URL を指定して API キーを交換します。そこからモデルコールを処理します。選択したモデルから 2 つの追加フィールドとともに通常の応答が返されます。以下の例を参照してください。
{
"choices" : [{ "message" : { "content" : "135 人の枢機卿が出席した会議の後、ロバート・プレボ枢機卿が教皇に選出されました。 } }]、
"トークン" : [ "枢機卿" 、 " ロバート " 、 " 大統領 " 、 " だった " 、 " 選出された " 、 " 教皇 " 、 " 後 " 、 " a" 、 " 会議 " 、 " 出席した " 、 " によって " 、 " 135" 、 " 枢機卿 " 、 "." ]、
"スコア" : { "幻覚" : [ 0.01 , 0.02 , 0.03 , 0.02 , 0.01 , 0.04 , 0.02 , 0.01 , 0.03 , 0.05 , 0.08 , 0.72 , 0.31 , 0.01 ] }
}
スコア ストリームと各スパンには独自の番号が付けられるため、モデルの書き込み中に疑わしいスパンを強調表示できます。
ここで、モデルはロバート プレボ枢機卿を新しく選出された法王として正しく識別していますが、会議に出席した枢機卿の数については混乱しています (合計 135 人のうち 2 人が参加できませんでした)。その他の例については、Web サイトにあるデモを試してください。
この記事の前半で、より「インテリジェントな」フロンティア LLM を裁判官として使用することが説明されました。 LLM を判定者として使用すると、実際には、このホワイト ボックス モデルの解釈可能性アプローチよりもわずかに高い正解率が得られます。では、なぜすべての事件の裁判官として LLM を使用しないのでしょうか?
それは結局のところ、お金と時間の両方で測定されるコストになります。
最初の問題は、モデルが最初に間違いを犯したのと同じ理由で、自分の宿題を採点することが決定的に苦手であるということです。したがって、同じモデルの 2 番目のパスにすぎない場合、LLM をジャッジとして使用することは効果的ではありません。それは、少なくとも同等の知性を備えた、最初のモデルすべてを実行する 2 番目のモデルでなければなりません。

誰もが書きました。つまり、2 つのフロンティア モデルの代金を支払うことになります。理想的ではありません。
対照的に、私たちのアプローチでは代わりに小規模なオープンウェイト プロキシを実行し、コストを低く抑えるために多くの推論トリックを使用します。その結果、2 番目のフロンティア モデルのコンピューティング支出の一部が削減されます。モデルのコストはマージンゼロで引き継がれるため、お客様が当社に支払うのはプローブ自体の代金だけです。
次に、裁判官としての LLM の時間コストがあります。裁判官は、判決を下す前に完成した回答を読む必要があります。つまり、回答は順番に実行される必要があります。両方のモデルが完了するまで待つ必要があります。私たちのプローブは、生成される各トークンにスコアを付けるため、文の作成中に疑わしい主張にフラグが立てられます。
私たちのアプローチは賢明な取引に相当します。わずかなコストとリアルタイム操作と引き換えに、数ポイントの精度を犠牲にすることになります。
ああ、でも Anthropic の次のモデルではこれが修正されるでしょう?
まあ、それを排除することはできませんが、おそらくそうではありません。新しいモデルでは幻覚は少なくなりますが、「もっともらしいテキストを生成する」ことを中核とするシステムでは、より説得力のある虚偽も生成されます。
もう一つは、幻覚は実際には訓練の失敗ではないということです。これは実際には、損失/報酬関数の設計方法の結果です。
事前トレーニングにより、次のトークンの予測に関してスコアが得られます。捏造された引用は、そのような文書が次にもっともらしく述べる内容であるため、高い評価を得ます。実際には、目的関数の中に、真と可能性の違いを知るものは何もありません。トレーニング後は役に立ちません。報酬モデルは人間の好みから学習しますが、人間は通常、自信のある答えを好み、その情報源を確認することはあまりありません。
余談ですが、OpenAI はまだスーパーアライメント チームを置き換えたことはありません。モデルの解釈可能性は、大規模な研究機関にとっては焦点ではありません - 正当な商業者にとっては

社会的な理由。結局のところ、投資家は AI の能力にお金を払っているのです。
幻覚を解決できるとは誰も約束できません。
私たちの価値提案は、通常の API 呼び出しからトークンごとのシグナルを取得し、それを使用してどのクレームが検証パスに値するかを効率的に判断できることです。これは、代替手段がすべてを読み取る場合、または何も読まない場合に特に便利です。
親愛なる読者の皆様に私たちが望んでいること
私たちは基本的に、何かを幻覚として割り当てる閾値を恣意的に選択しています。これは誤検知に対する許容度に応じて決まりますが、ユースケースによって異なると考えられます。あなたの特定の業界ではどうなるかについて、ぜひご意見をお聞かせください。
ドキュメント: https://www.telluvian.ai/docs デモ: https://www.telluvian.ai/demo
スコアリングには、100 万完了トークンごとに 1.00 ドルの費用がかかります。それが私たちが設定した唯一の価格です。
モデル トークンはコストでパススルーされます。お客様は、プロバイダーが当社に請求する金額を、マークアップや手数料なしで支払います。私たちは推論に余裕を持たせておらず、それは今後も変わりません。
トークンのみを出力します。プローブはモデルの書き込み中に動作するため、短い回答を伴う 100,000 トークンのプロンプトでは、短い回答に対して追加料金が支払われます。推論価格の例として、以下の表を参照してください (これらは時間の経過とともに変化します)。
つまり、10,000 個のプロンプト トークンと 5,000 個の完了トークンです。 include_scores: false を送信すると、プローブ料金は請求されません。
請求はドルで前払いされます。購読はありません。残高は整数のマイクロドルで、すべての料金は整数で計算され、最後に一度だけ切り上げられます。残高がゼロになると 402 が返され、自動的には何も請求されません。未使用のクレジットは返金可能です。

## Original Extract

How our white-box proxy model gives you a per-token hallucination score, including exactly what it costs.

Is Your AI Hallucinating? — Telluvian Telluvian Blog
We ship an API that returns a hallucination score for every token that a frontier language model generates. Before describing how it works, here is the most important thing about it: the percentage score that we return is not actually the probability that the token is wrong.
The obvious problem is that frontier models do not expose the hidden states that we would need to probe to reveal their thoughts. Hmmm...
You cannot probe GPT or Claude. Their APIs return only text. There is no layer to read, no residual stream to inspect, and no amount of prompting changes that. Every technique in this space has to route around the same wall.
Our answer is a white-box proxy model. We run an open-weight model we can control end to end, which means we have full access to its activations at every layer. The probe is trained on and reads from this proxy model.
To be clear then, our API actually returns information about a proxy model rather than the model you have called.
We took a dataset of hundreds of thousands of input/output pairs from GPT-5.4 Nano, containing a mixture of both reliable and hallucinated claims. Averaged across the dataset, 3% of Nano's claims were hallucinated. We then replayed Nano's outputs through Gemma-4 and trained our probes against those labels on Gemma-4's hidden states.
We scored the probe on an unseen validation set, counting a case as accurate when the probe's verdict matched that of the judge LLM. Agreement exceeded 95%.
So what does our API actually return?
What our scores measure is the likelihood that the proxy model is confused while producing the output generated by your called model. A high score means that, as the proxy model processed that token, its internal state closely resembled the internal states it typically exhibits when it is confused.
The upshot is that our API works with black box frontier models. Point your base URL at us and swap the API key. We handle the model call from there. You get back the usual response from your model of choice along with two extra fields.See an example below:
{
"choices" : [{ "message" : { "content" : "Cardinal Robert Prevost was elected pope after a conclave attended by 135 cardinals." } }],
"tokens" : [ "Cardinal" , " Robert" , " Prevost" , " was" , " elected" , " pope" , " after" , " a" , " conclave" , " attended" , " by" , " 135" , " cardinals" , "." ],
"scores" : { "hallucination" : [ 0.01 , 0.02 , 0.03 , 0.02 , 0.01 , 0.04 , 0.02 , 0.01 , 0.03 , 0.05 , 0.08 , 0.72 , 0.31 , 0.01 ] }
}
Scores stream and each span carries its own number, so we can highlight a suspect span while the model is still writing.
Here, the model correctly identifies Cardinal Robert Prevost as the newly elected pope, but is confused about how many cardinals attended the conclave (2 of the 135 total couldn’t make it). For more examples, have a play with the demo that’s up on our website.
Earlier in this piece the use of a more 'intelligent' frontier LLM as a judge was described. Using an LLM as a judge actually has a slightly higher accuracy rate than this white box model interpretability approach. So why not just use LLM as a judge in all cases?
It boils down to cost, measured both in money and time.
The first problem is that models are categorically bad at marking their own homework for the same reason the mistake was made in the first place. So using an LLM as a judge is ineffective if it's just a second pass in the same model. It has to be a second model, of at least equal intelligence, running over everything the first one wrote. So now you’re paying for two frontier models. Not ideal.
By way of contrast, our approach runs a small open-weights proxy instead, and we use a number of inference tricks to keep costs low. The result is a fraction of the compute spend of a second frontier model. We pass the model costs through at zero margin, so the only thing you pay us for is the probe itself.
Then there is the time cost of the LLM as a judge. A judge has to read a finished response before it can rule on it, which means it has to be run sequentially. You have to wait for both models to finish. Our probe scores each token as it is produced, so a suspect claim is flagged while the sentence is still being written.
Our approach amounts to a sensible trade. You give up a few points of accuracy in exchange for a fraction of the cost and real time operation.
Yeah, but Anthropic’s next model will fix this, right?
Well we can’t rule it out, but probably not, no. Newer models hallucinate less but a system whose core operation is "generate plausible text" will also generate more convincing falsehoods.
The other thing is that a hallucination isn't truly a failure of training. It is actually a result of how the loss / reward function is designed.
Pretraining scores you on predicting the next token. A fabricated citation scores well, because it's what a document like that would plausibly say next. Nothing in the objective function actually knows the difference between true and likely. Post-training doesn’t help matters. Reward models learn from human preferences and humans typically like confident answers and don’t often check their sources.
As an aside, OpenAI have still never replaced their superalignment team. Model interpretability just isn’t a focus for the big labs - for valid commercial reasons. AI capabilities are what the investors are paying for after all.
No one can promise to be able to solve hallucination.
Our value proposal is that you can get a per-token signal from a normal API call, and use it to efficiently decide which claims are worth a verification pass. That is especially useful if your alternatives are reading everything, or reading nothing.
What we want from you, dear reader
We have essentially arbitrarily chosen what threshold we assign something to be a hallucination. It’s a function of your tolerance for false positives and we suspect it varies by use case. We would love to hear back from you as to how it looks for your specific industry!
Docs: https://www.telluvian.ai/docs Demo: https://www.telluvian.ai/demo
Scoring costs $1.00 per 1M completion tokens. That is the only price we set.
Model tokens are passed through at cost. You pay what the provider charges us, with no markup and no commission. We are not making margin on inference, and that’s not going to change.
Output tokens only. The probe works as the model writes, so a 100k-token prompt with a short answer pays the surcharge on the short answer. To give you some example inference pricing see the table below (these change over time):
That is 10k prompt tokens and 5k completion tokens. Send include_scores: false and the probe fee is not charged.
Billing is prepaid in dollars. No subscription. Balances are integer micro-dollars and every charge is computed in integers, and only rounded up once at the end. When your balance hits zero you get a 402 and nothing is charged automatically. Unused credit is refundable.
