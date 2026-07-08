---
source: "https://softlight.com/blog/measuring-and-improving-ai-generated-ui-design"
hn_url: "https://news.ycombinator.com/item?id=48832749"
title: "Measuring and improving AI-generated UI design"
article_title: "Measuring and improving AI-generated UI design | Softlight"
author: "ashwin153"
captured_at: "2026-07-08T15:10:22Z"
capture_tool: "hn-digest"
hn_id: 48832749
score: 4
comments: 0
posted_at: "2026-07-08T14:50:58Z"
tags:
  - hacker-news
  - translated
---

# Measuring and improving AI-generated UI design

- HN: [48832749](https://news.ycombinator.com/item?id=48832749)
- Source: [softlight.com](https://softlight.com/blog/measuring-and-improving-ai-generated-ui-design)
- Score: 4
- Comments: 0
- Posted: 2026-07-08T14:50:58Z

## Translation

タイトル: AI が生成した UI デザインの測定と改善
記事のタイトル: AI が生成した UI デザインの測定と改善 |ソフトライト
説明: コーディング エージェントの設計出力をどのようにテスト、測定、改善しようとしたか。

記事本文:
AI が生成した UI デザインの測定と改善
コーディング エージェントの設計出力をどのようにテスト、測定、改善しようとしたか。
私のガールフレンドはこれまで一度もコードを書いたことはありませんでしたが、ある朝起きて、入手困難な Resy の予約を確保するためのボットを構築しました。ボットはうまく機能しました。彼女の最初の不満は、UI の見た目が悪いということでした。それはまさにAIが生成したように感じました。
彼女は一人ではありません。これを「AI スロップ」と呼ぶ人もいます。このモデルには「センスが欠けている」という人もいる。ソフトウェアの作成は民主化されました。しかし、これまでのところ、優れた UI デザインは存在していないようです。
同じ質問が何度も返ってきます。エージェントに、より良い UI デザイン作業を実行してもらい、私がそうする必要がなくなるようにすることはできるでしょうか?
これは難しい問題であり、解決したとは言えません。しかし、私はある程度の進歩を遂げました。その過程で私が学んだことは以下の通りです。
AI設計の品質を測定する
ほとんどの人はデザインを主観的なものとして扱います。それよりも測定可能だと思います。
優れたデザインは、ユーザーと企業の目標達成に役立ちます。機能が達成する意味を理解していれば、デザインを判断することができます。
私が Uber で取り組んだ次の例を見てみましょう。当時、Uber は乗客を他のサービス (Uber Eats、食料品店、予約配車など) にクロスセルしようとしていました。そこで、同社はこれらのサービスをより強調するためにホームページのデザインを変更しました。
これは「良い」設計変更でしょうか?それは実際のところ、より多くのユーザーが Uber Eats を試すかどうかにかかっています。地図上の車に確信が持てなくなるため、配車をリクエストするユーザーが減少するのでしょうか?ユーザーとビジネス目標の制約を考慮すると、デザインが「より良い」かどうかを推測することができます。
これが最初の大きな課題でした。ユーザーとビジネスのコンテキストを理解し、設計出力を正しく判断できるように、AI によって生成されたさまざまな設計変更をどのように作成するか?
そうするために、私は作戦を取った

ソースアプリを作成し、追加したい新機能を作成しました。
何百もの新機能を手動で生成してレビューしました。私は見つけた設計上の問題に注釈を付け、LLM に問題を要約してもらいました。問題の大部分は以下に分類されます。
優先順位付け: 優れたデザイナーは、UX において容赦なく優先順位を付けます。核となる情報のみが表示され、製品全体の核となる経路は非常に明白であり、ユーザーの邪魔をするものは他にありません。モデルはすべてのユーザーのアクションが同等に重要であると考えているようで、その結果、ユーザーは散らかったままになります。
タイポグラフィ、間隔、色: モデルは色を多用しており、一貫した間隔がなく UI が乱雑になっています。これはすべて #1 につながります。優れたデザイナーは、色、フォント、間隔を使用して、特定のユーザーのアクションに注意を向けます。このモデルでは、どのアクションが重要であるかに優先順位が付けられていません。
UI「バグ」: 人間のデザイナーが決して台無しにしないような単純なもの (コンポーネントの位置合わせがずれている、テキストの折り返しにより奇妙に見えるコンポーネントが作成されるなど)
デザイン システムの遵守: さまざまなコンポーネント、ボタンの半径、フォントなどを導入すると、統一感のないエクスペリエンスが生じます。
より良い設計を生成するための改善点のテスト
コーディング エージェントの設計出力を向上させるために、20 以上の異なるエージェント設定を実験しました。次に、これらのエージェント設定を使用して同じオープンソース機能作成テストを実行し、設計が改善されたかどうかを手動で評価しました。その結果、500 を超える画面をレビューする必要があったため、同じ設計変更に対するさまざまなエージェントからの出力を迅速に比較してスコア付けするための評価ツールを作成しました。
自己改善のループ。モデルがデザインを「認識」すれば、デザインを改善できます。モデルが生成した UI をモデルに示し、問題の修正を依頼すると、有意義な成果が得られることがわかりました。驚くべきことではありませんが、そうでした

最大のレバー。
自己改善ループでは、AI 設計でフラグを立てた一般的な問題に基づいたルーブリックによって出力が改善されました。最初に機能を生成したときに特定した一般的な欠陥を共有することが、レビュー ルーブリックで最も役立ちました。
有名なアプリから「優れたデザイン」の関連サンプルを注入すると、出力が向上します。私は以前、優れたデザインのアプリ (Stripe、Airbnb など) から 500k 以上の画面のコーパスを構築しました。ルーブリック内のモデル例を示すことで、出力がさらに向上しました。
驚くほどうまくいかなかったもののいくつか:
「一般的な設計のベスト プラクティス」に基づいてルーブリックを作成する: これはあまり役に立ちませんでした。私の推測では、モデルのトレーニング データにはこれらがすでに含まれていると思います。 AI 設計で気づいた一般的な欠陥を取り上げ、ルーブリックでそれらを強調したときに、本当の成果が得られました。
古典的な AI デザインのスロップシグナルのネガティブな例 (紫色のグラデーション、特大の丸いカード、絵文字と箇条書きの特徴リストなど): AI デザインの一般的なパターンを示しても、あまり役に立ちませんでした。 「AI デザイン」という従来の貧弱なシグナルよりも、「優れたデザイン」のイメージを示す方がより生産的でした。
注意すべき点: すべての作業に潜在的な欠陥があるのは、「出力の向上」であり、これは私の人間による評価に基づいています。私が本当に言えるのは、上記のレバーのおかげで私の目にはデザインがより良くなったということだけです。
最近、エージェントをパッケージ化し、いくつかの企業と共有しました。企業によっては、チームは UI PR の約 70% についてエージェントの提案を承認しています。これにより、私の評価が正当であるというある程度の自信が得られます。ただし、これはまだ欠陥のあるシグナルである可能性があります。エンジニアは重要な提案よりも「単純な」提案を受け入れやすい可能性があります。測定のその部分はまだ理解中です。
測定された改善は、より多くの時間とトークンを費やすことで得られました。の

エージェントは設計変更を行い、設計変更を見て判断し、改善します。これは素晴らしいことですが、変更を行うのに時間がかかり、トークンのコストも高くなります。
次の課題は、このプロセスを高速化し、トークンの効率を高めることです。エージェントの開始時間は平均約 30 分、設計された機能ごとに約 $50 かかりました。どちらも最大 50% 減少しました。まだこれを乗り越えています。実験の方法と結果については、また改めて投稿する予定です。

## Original Extract

How I tested, measured, and tried to improve the design outputs of coding agents.

Measuring and improving AI-generated UI design
How I tested, measured, and tried to improve the design outputs of coding agents.
My girlfriend has never written a line of code, but woke up one morning and built a bot to secure hard-to-get Resy reservations. The bot worked great. Her first complaint: the UI looked bad. It just felt AI-generated.
She is not alone. Some people call it “AI slop.” Others say the models “lack taste.” Software creation has been democratized. But so far, it seems like good UI design hasn’t been.
I keep coming back to the same question: can I get my agents to do better UI design work, so I don’t have to?
This is a hard problem, and I can’t say I have solved it. But I made some progress. The following is what I learned through the process.
Measuring the quality of AI design
Most people treat design as subjective. I think it’s more measurable than that.
Good design helps users and businesses accomplish goals. If you understand what a feature means to accomplish, it is possible to judge the design.
Take the following example I worked on at Uber. At the time, Uber was trying to cross-sell riders into other services: Uber Eats, Grocery, Reserved Rides, etc. So, the company made a design change to the home page to emphasize those services more:
Is this a “good” design change? It really depends: do more users try Uber Eats? Do fewer users request a ride because they lose certainty from the cars on the map? Given the constraints of the user & business goals, it is possible to guess whether a design is “better”.
This was the first big challenge - how do I create a wide variety of AI-generated design changes where I understand the user & business context, so I can correctly judge the design output?
To do so, I took open source apps, and created new features I wanted to add.
I manually generated and then reviewed 100s of new features. I annotated the design issues I found and had an LLM summarize the issues. The vast majority of issues bucketed into the following:
Prioritization: Great designers prioritize ruthlessly in the UX. Only the core information is shown, the core path through the product is extremely apparent, and nothing else gets in a user’s way. The models seemed to think all user actions are equally important, and as a result users are left with clutter.
Typography, spacing, and color: The model overuses color and clutters the UI with no consistent spacing. This all links back to #1 - great designers use color, fonts, and spacing to draw attention to specific user actions. The model is not prioritizing which actions are important.
UI “bugs”: Simple things that a human designer would never mess up (alignment of components is off, text wrapping creates weird looking components, etc.)
Design system adherence: Introducing different components, button radii, fonts, etc. leads to an experience that does not feel unified.
Testing improvements to generate better designs
I experimented with 20+ different agent setups to improve the design outputs of coding agents. I then ran the same open source feature creation tests with those agent setups, and manually evaluated whether the design improved. This led to over 500 screens that needed to be reviewed, so I created an evaluation tool to quickly compare and score the outputs from different agents for the same design change.
Self-improvement loops. A model can improve the design if it “sees” it. I found that showing the model the UI it generated, and asking it to fix issues led to meaningful gains. Not surprising, but was the biggest lever.
In the self-improvement loop, a rubric based on the common issues I flagged with AI design improved the output: Sharing the common flaws I identified when generating the features initially was the most helpful in the review rubric.
Injecting relevant examples of “good design” from well-known apps improves the output: I previously built a corpus of 500k+ screens from well designed apps (think Stripe, Airbnb, etc). Showing the model examples of those in the rubric further improved the outputs.
Some of the things that surprisingly didn’t work:
Creating a rubric based on “common design best practices”: This didn’t really help. My guess is the models already have these in their training data. The real gain came when I took the common flaws I noticed in AI design, and emphasized those in the rubric.
Negative examples of classic AI design slop signals (e.g. purple gradients, oversized rounded cards, emoji-bullet feature lists): Showing it the common patterns of AI design didn’t help much. It was more productive to show it images of “good design”, rather than the poor traditional signals of “AI design”.
A thing to note: a potential flaw with all of the work is “improved output” is based on my human evaluation. All I can truly say is that the above levers made the design better in my eyes .
I recently packaged the agent and shared it with some companies. Depending on the company, teams have approved the agent suggestions on ~70% of UI PRs. This gives some confidence that my evaluation is valid. It could still be a flawed signal though: engineers might be accepting “simple” suggestions more readily than important ones. Still figuring that part of measurement out.
The measured improvements came from spending more time and tokens. The agent makes a design change, looks at the design change, judges it, and improves it. This is great, but it takes longer to make the changes & costs more tokens.
The next challenge is speeding this process up and making it more token efficient. The agent started at an average of ~30 minutes and ~$50 per designed feature. We have decreased both by ~50%. Still working through this. Will make another post with the methods and results of the experiments there.
