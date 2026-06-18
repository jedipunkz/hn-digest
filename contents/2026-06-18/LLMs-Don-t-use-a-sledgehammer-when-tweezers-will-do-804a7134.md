---
source: "https://superuserdone.com/posts/2026-06-18-dont-use-a-sledgehammer/"
hn_url: "https://news.ycombinator.com/item?id=48581472"
title: "LLMs: Don't use a sledgehammer when tweezers will do"
article_title: "LLMs: Don't use a sledgehammer when tweezers will do | SuperUserDone"
author: "SuperUserDone"
captured_at: "2026-06-18T07:48:53Z"
capture_tool: "hn-digest"
hn_id: 48581472
score: 1
comments: 0
posted_at: "2026-06-18T06:20:05Z"
tags:
  - hacker-news
  - translated
---

# LLMs: Don't use a sledgehammer when tweezers will do

- HN: [48581472](https://news.ycombinator.com/item?id=48581472)
- Source: [superuserdone.com](https://superuserdone.com/posts/2026-06-18-dont-use-a-sledgehammer/)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T06:20:05Z

## Translation

タイトル: LLM: ピンセットで十分な場合は大ハンマーを使用しないでください
記事のタイトル: LLM: ピンセットで十分な場合は大ハンマーを使用しないでください |スーパーユーザー完了
説明: LLM は、必ずしも作業に適したツールであるとは限りません。

記事本文:
SuperUserDone 投稿について 投稿について
ホーム
LLM: ピンセットが使える場合は大ハンマーを使用しないでください
次の並べ替え関数を考えてみましょう。
llm_api_of_your_choice から setup_llm_api をインポート
llm = setup_llm_api(API_KEY)
def llm_sort (単語):
"
応答 = llm(プロンプト)
折り返しご返答をさせていただきます。分割()
まともな人なら絶対にこれを使いませんよね？このタスクでは、バブル ソートでも、速度、精度、信頼性の点で LLM よりも優れたパフォーマンスを発揮します。その理由は非常に簡単です。 LLM は、数十年前に初年度のアルゴリズム コースで解決されたことを実行するために、大量のコンピューティングを消費します。さらに悪いことに、答えが間違っている可能性さえあります。さらに悪いことに、API がダウンしている可能性があります。言うまでもなく、モデルを自己ホストしていない場合、データはプライバシー ポリシーの曖昧な条項とともにランダムな第三者に送信されることになります。
さあ、教えてください。なぜあらゆるデータの問題に対して LLM を求める必要があるのですか、ボゾ！
上記の並べ替え関数は冗談です。データが少し汚いからといって、本番システムを同じように AWS の請求額が高くなるという冗談にしないでください。
LLM の時代には、邪悪な誘惑が定着しています。適切なエンジニアリング上の判断に到達する代わりに、デフォルトでは LLM に到達するようになりました。データセットをクリーニングする必要がありますか?モデルさんに聞いてみてください。 2 つのレコードを照合する必要がありますか?モデルさんに聞いてみてください。一部のテキストを抽出する必要がありますか?モデルさんに聞いてみてください。検証、ルーティング、重複排除、正規化、クラスタリング、または検索が必要なものはありますか?
おそらく、おそらく、魔法のテキスト ボックスがすべての仕事に適したツールではないということを考えたことはありますか?
LLM は役に立たないわけではなく、むしろその逆です。これらは、適切な問題に適用すると、驚くべきテクノロジーになります。ただし、何が起こるかというと、それらは柔軟です

遅くて予測不可能でデバッグが難しいオプションである場所では、適切なソリューションのように見えるはずです。同様に、彼らはデモを良いものに見せるのが得意なので、悪いソリューションを洗練されたものに見せます。
しかし、実稼働システムは、CIO へのプレゼンテーションで 1 回だけ機能する必要がある、厳選された 5 つのサンプルに関する AI 生成ノートブック内のデモを気にしません。
彼らは、コスト、レイテンシー、スループット、再現性、検査可能性、可観測性、信頼性、エッジケース、障害モード、そして同じ入力が明日の朝に同じ出力を与えるかどうかを重視します。
LLM は、ビジネス ロジック、エッジケース、分類法、顧客を理解していません。あなたがやる。 「微調整してください、兄弟」と言うかもしれません。確かに、「兄弟」、しかし、計画が「微調整」に始まり「微調整」に終わり、部屋の誰も、必要なデータ、ラベルの実際の意味、損失関数とは何か、許容できる障害モードを説明できない場合、解決策はありません。さらに高価なデモがあります。
次回、選択を迫られたときは、「これに LLM を使用するにはどうすればよいでしょうか?」と尋ねないでください。むしろ、問題を深く理解して次のように尋ねてください。
この問題を解決する最小かつ最も信頼性の高いツールは何でしょうか?
最終的には、速度と精度の点で LLM を上回り、検査しやすく信頼性が高く、ラップトップ上に 1 週​​間で構築できるものが完成します。
そのツールが正規表現である場合もあります。ルックアップ テーブルや結合の場合もあります。場合によっては、テキスト距離アルゴリズム、スパース検索モデル、高密度埋め込みモデル、検索インデックス、分類子、ルール エンジン、パーサー、またはより優れたデータであることもあります。手作業で行う場合もあります。
そして、はい、LLM である場合もあります。
これらのどれに到達するかを知ることが仕事です。その判断がエンジニアリングとオートコンプリートを分けるものです

て。 LLM がデフォルトにならないようにしてください。
確かに、「結合を使用しました」と言って資金を集めている人はいません。退屈で正しいだけで十分かもしれないと考えたことはありますか?あなたは保険会社に勤めています。 Big Data Startup Enterprises Incorporated Limited にしないでください。
ピンセットが使える場合は、大ハンマーは使用しないでください。

## Original Extract

LLMs are not always the right tool for the job.

SuperUserDone Posts About Posts About
Home
LLMs: Don't use a sledgehammer when tweezers will do
Consider this sorting function:
from llm_api_of_your_choice import setup_llm_api
llm = setup_llm_api(API_KEY)
def llm_sort (words):
prompt = f "Given these words: { ' ' . join(words) } , sort them alphabetically. Reply only with the words separated by spaces."
reply = llm(prompt)
return reply . split()
No sane person would ever use this, right? Even bubble sort would outperform an LLM on this task in speed, accuracy and reliability. It’s quite simple to see why. An LLM burns through a massive amount of compute to do something a first-year algorithms course solved decades ago. Worse: it might even get the answer wrong! Even worse: the API might be down! Not to mention if you are not self-hosting your model your data is now being transmitted to a random third party with a vague clause in their privacy policy.
Now tell me: WHY REACH FOR AN LLM FOR EVERY DATA PROBLEM, YOU BOZO!?!
The sort function above is a joke. Don’t turn your production system into the same joke with a higher AWS bill just because your data is slightly messy.
In the age of LLMs, a wicked temptation has taken hold. Instead of reaching for proper engineering judgment, the default has become reaching for an LLM. A dataset needs cleaning? Ask the model. Two records need matching? Ask the model. Some text needs extracting? Ask the model. Something needs validating, routing, deduplicating, normalising, clustering or searching?
Have you considered that maybe, perhaps, perchance, the magical text box is not the right tool for every job?
LLMs are not useless, quite the contrary. They are amazing pieces of technology when applied to the right problems. What happens though is that they are flexible enough to look like the right solution in places where they are the slow, unpredictable, hard-to-debug option. Likewise, they make bad solutions look sophisticated because they are great at making demos look good.
But production systems do not care about your demo in your AI-generated notebook on five cherry-picked examples that only needs to work once in your presentation to the CIO.
They care about cost, latency, throughput, repeatability, inspectability, observability, reliability, edge cases, failure modes, and if the same input gives the same output tomorrow morning .
LLMs do not understand your business logic, edge-cases, your taxonomy, your customers. You do. You might say: “Just fine-tune it, bro.” Sure, “bro”, but if the plan starts and ends at “fine-tune it,” and nobody in the room can explain what data you need, what labels actually mean, what a loss function is, and what failure modes you are willing to tolerate, then you do not have a solution. You have an even more expensive demo.
Next time you are faced with a choice, don’t ask: How can we use an LLM for this? Rather, try to understand your problem deeply and ask:
What’s the smallest, most reliable tool that solves this problem?
You will end up with something that outperforms an LLM in speed and accuracy, is more inspectable and reliable, and can be built on a laptop in a week.
Sometimes that tool is a regular expression. Sometimes it’s a lookup table or a join. Sometimes it’s a text-distance algorithm, a sparse retrieval model, a dense embedding model, a search index, classifier, rules engine, parser or better data. Sometimes it is doing it by hand.
And yes, sometimes it is an LLM.
Knowing which of these to reach for is the job. That judgment is what separates engineering from autocomplete. Please just don’t let LLMs become a default .
Sure, nobody is raising money by saying: “we used a join.” Have you considered that boring and correct might be enough? You work for an insurance company. Don’t turn it into Big Data Startup Enterprises Incorporated Limited .
Don’t use a sledgehammer when tweezers will do.
