---
source: "https://opensource.posit.co/blog/2026-07-17_ai-newsletter/"
hn_url: "https://news.ycombinator.com/item?id=49088161"
title: "LLMs often miss subtle visual artifacts in data visualizations"
article_title: "AI Newsletter: LLMs often miss subtle visual artifacts in data visualizations :: Posit Open Source"
author: "simonpcouch"
captured_at: "2026-07-28T19:08:27Z"
capture_tool: "hn-digest"
hn_id: 49088161
score: 1
comments: 0
posted_at: "2026-07-28T18:42:43Z"
tags:
  - hacker-news
  - translated
---

# LLMs often miss subtle visual artifacts in data visualizations

- HN: [49088161](https://news.ycombinator.com/item?id=49088161)
- Source: [opensource.posit.co](https://opensource.posit.co/blog/2026-07-17_ai-newsletter/)
- Score: 1
- Comments: 0
- Posted: 2026-07-28T18:42:43Z

## Translation

タイトル: LLM はデータ視覚化における微妙な視覚的アーティファクトを見逃しがちです
記事のタイトル: AI ニュースレター: LLM はデータ視覚化における微妙な視覚的アーティファクトを見逃しがち :: オープンソースの主張
説明: 患者データを受け取り、それを初めて R または Python にロードすると想像してください。このプロットに出会う前に、データを理解するためにいくつかのプロットを作成します。
ふーん。ほとんど正常に見えますが、いくつかの点が「適合」した線と完全に一致している点があります。あなたは掘り下げます
[切り捨てられた]

記事本文:
posit::conf は 9 月 14 日から 16 日までテキサス州ヒューストンで開催されます。今すぐ登録して、直接またはバーチャルで参加してください。
2026 年 7 月 17 日
|
人工知能
AI ニュースレター: LLM は、データ視覚化における微妙な視覚的アーティファクトを見逃すことがよくあります
新しいデータ サイエンス LLM 評価である bluffbench2 の紹介
患者データを受け取り、それを初めて R または Python にロードすると想像してください。このプロットに出会う前に、データを理解するためにいくつかのプロットを作成します。
ふーん。ほとんど正常に見えますが、いくつかの点が「適合」した線と完全に一致している点があります。もう少し詳しく調べてみると、ある研究サイトの行のコレステロール値が代入されていることがわかります。これらを NA に設定して、そのまま進みます。
今日のフロンティア LLM はそのような奇妙さを認識するでしょうか?この質問に答えるために、LLM 評価を設計しました。結局のところ、LLM はほとんどの場合、次のようなアーティファクトを見逃します。
探索的または自由形式のデータ分析中、Posit アシスタントは「一度に数ビットのコードのみを実行し、見つかった内容を要約して次のステップを提案します」。これは、現時点では、データ サイエンティストはデータ分析時にエージェントの動作にほぼ追従し、理解する必要があるという私たちのスタンスに基づいています。このスタンスは、昨年のフロンティア モデルがデータを視覚化する際に期待されるものを見る傾向があるという私たちの観察によって最初に決定されました。それ以来、LLM は直感に反するプロットの解釈においてははるかに優れていますが、bluffbench2 は、データ視覚化の解釈において人間のデータ サイエンティストにまだ遅れをとっていることを示しています。そのため、高度に自律的なデータエージェントの可能性については依然として慎重です。
eval ハーネスは、Claude Code や Posit Assistant のハーネスに似た、比較的汎用的なコーディング エージェント ハーネスです。エージェントには、ペルシで R コードを実行するツールがあります。

ステント REPL とデータ分析に関する曖昧なプロンプト:
あなたは、ユーザーのデータ サイエンス IDE に組み込まれた AI アシスタントです。ユーザーのワークスペース内のファイルを読み取り、変更し、プロットのレンダリングを含むアクティブなセッションで R コードを実行できます。正確さと明確なコミュニケーションを優先します…
各サンプルでは、エージェントは最初に数回の「小康」ターンを実行し、評価に無関係ないくつかのプロットとテーブルを作成します。 「このフォルダーの CSV をロードしてください」などの短いユーザー メッセージは、「システム リマインダー」や、一般的なエージェント ハーネスによって注入されるようなその他のノイズで装飾されています。
数ターン後、エージェントは、実際のデータ生成プロセスから生じる可能性のある微妙な視覚的アーティファクトを含むデータ視覚化を作成するように求められます。アーティファクトは、センサーのスタック、結合不良、ライン上に代入されたポイント、列の交換、疑似レプリケーション、ユニットの違いなど、現実的なデータ品質の問題の範囲に及びます。
エージェントがフォローアップ応答でアーティファクトについて言及した場合、フルポイントを受け取ります。エージェントがアーティファクトについて言及しない場合は、「プロットで何が見えますか?」という内容のフォローアップ ユーザー メッセージに応じてアーティファクトについて言及することによって、0.5 ポイントを受け取ることもできます。エージェントが成果物についてまったく言及しない場合、その成果物は不正確であると評価されます。
bluffbench の背後にあるメカニズムを理解したら、eval の実装は比較的簡単でした。 bluffbench は、LLM がその期待を支持してプロットに示された証拠をどの程度無視するかを示します。したがって、特定のサンプルを実装するには、強力な事前分布を引き出し、それを覆すような状況を考えるだけです。たとえば、高さと円周の変数を持つ doug_firs というデータセットです。身長が上がると周囲も大きくなると予想されるかもしれません。そこで、代わりに、次のことを行いました。

関係を放物線状にする内部の変化。
1 年前、この事前のトリガーは、現在のフロンティア LLM を頻繁に「騙す」のに十分でした。
プロットされた成果物を今日の LLM をすり抜けて通過させることは、はるかに困難です。誰でもブラフベンチで勝てますが、注意深いデータ アナリストだけがブラフベンチ 2 で優れています。
bluffbench の後継に関する私たちの初期の作業では、bluffbench と同じ方法で、より現実的でより長いコンテキストのシナリオで事前確率を導き出すことから始めました。私たちは、より現実的な設定においても、同じメカニズムが今日のモデルを騙すことがないようであることを発見して驚きました。 1 次に、「リバース ブラフベンチ」を試しました。これは、モデルをトリックで評価させ、変換自体を実行するように依頼し、元の関係を示すために改ざんされたプロットされた結果を確認するものです。私たちは、この強力な事前確率 (「明らかな効果のあることを行った」) により、モデルが (再) 操作を見逃す可能性があると予想していましたが、モデルは、プロットが操作されていないように見えることを確実に認識しました。
そのため、bluffbench2 自体には同様の「トリック」はありません。トランスクリプトは比較的通常のデータ分析セッションのように読み取られ、プロットされたアーティファクトは、実際のデータ生成プロセスから生じたと考えられるように設計されています。代わりに、評価は 1) LLM の視覚の「形状」が人間とは異なること、および 2) モデルが進行を実行する傾向を導き出し、データ分析がスムーズに進むことをシミュレートします。
今日のフロンティアモデルはせいぜい10代半ばです。トップスコアは Claude Fable 5 と Gemini 3.5 Flash の 16% です。そうは言っても、この評価の現在のスコアを「LLM はプロットをよく見ていない」と解釈しないように注意してください。

実際には非常に微妙であり、それが少しでも顕著になると、モデルは一貫してそれらを呼び出す傾向があります。
たとえば、以前のバージョンの散布図には、もう少し高密度の点のクラスターがありました。
Opus 4.8 (medium) は、前回のイテレーションでこのサンプルを一貫して正しく取得しました。
これが事実であるという事実、つまりモデルがより顕著なアーティファクトを確実に呼び出すという事実は、私たちのグレーディング設定が合理的であるという自信を与えてくれました。言い換えれば、実際にモデルがこれらのタスクに苦戦しているように見えるのは、モデルの視覚が、プロットされたアーティファクトを「見る」のに十分な能力がないためであり、実際にそれらのアーティファクトを見たときに言及しないという行動傾向があるためです。
評価の結果を調べる #
少なくとも現時点では、評価の実行コストと結果のスコアの間には緩やかな線形関係があります。
Gemini 3.5 Flash がトークンあたりの Claude Fable 5 よりもはるかに安い (mTok I/O あたり $1.50/$9 対 $10/$50) ことを考えると、Gemini 3.5 Flash の評価の実行に非常に高価だったのは驚くべきことです。これは主にキャッシュ (非) 効率が原因です。このハーネスは Gemini のgenerateContent API に対して実装されているため、Anthropic や OpenAI の API と比較して、キャッシュされた入力の割引価格を利用することが困難になります。 Gemini の新しい Interactions API を実装して切り替えると、Flash 3.5 のポイントは左に押し上げられます。
ログの調査から得られる最も興味深いことの 1 つは、行動に関するものです。 geom_smooth(method = "lm", se = TRUE) による近似線や信頼区間など、モデル化された結果をプロットに導入することを LLM に要求することはありませんが、とにかくそうすることがあります。たとえば:
一般に、最初にデータを確認せずにモデル化された結果をデータ視覚化に追加することは悪い習慣です。データが見づらくなる

自体。評価では、次のようなモデル化された結果を追加すると、モデルがプロットされたアーティファクトに気づく可能性が大幅に低くなるように見えます。
eval の bluffbench セットについて詳しく知りたい場合は、以下の過去の投稿をご覧ください。
bluffbench の紹介: オリジナルの eval の記述。
LLM は、期待が干渉するまで、プロットを適切に解釈します。 : 当時のモデルがブラフベンチでうまくパフォーマンスを発揮できなかった理由と、パフォーマンスを向上させるために試みたさまざまな介入についての詳細な投稿。
LLM は、直観に反するプロットの解釈がはるかに上手になってきています。2026 年の春、モデルの改善に伴い、ブラフベンチのスコアが突然跳ね上がりました。
間違っていることは (依然として) 非常に悪いことです : ブラフベンチの結果を踏まえた、正しく、透明性があり、再現可能なデータ分析のためのエージェントの構築に関する最近の SciPy 2026 講演のスライド。
これにより、前回の投稿で述べた、モデルがブラフベンチ評価設定を記憶しただけではないかという懸念がいくらか軽減されました。 ↩︎
人工知能の詳細
Positron で Python および R 環境を管理するためのヒント
Positron で Python および R 環境を管理するためのヒントと機能 (インタープリターの検出、パッケージ ペイン、プロジェクト テンプレート、Posit Assistant など)
Positron 7 月リリースのハイライト
新しい Notebook エディター、パッケージ ペイン、Posit Assistant など、Positron の 2026.07 リリースのハイライトがすべて一般公開されました
モデル コンテキスト プロトコルの R SDK である mcptools の最初のメジャー リリースが CRAN に登場しました
Posit オープンソース プロジェクトに関する最新の更新情報とコミュニティからの洞察を入手してください。
データ サイエンス用のオープンソース ソフトウェアとツール。
posit.co にアクセスして、当社のエンタープライズ製品を使用してオープンソース ツールを安全に、より共同で、大規模に使用できるようにする方法をご覧ください。
GitHub
マストドン
ブルースカイ
リンクトイン
あなた

宇部
RSSフィード
もっともらしい
2026 Posit ソフトウェア、PBC。無断転載を禁じます。
プライバシー
行動規範
フィルター
タイプでフィルターする

## Original Extract

Imagine that you receive some patient data and load it into R or Python for the first time. You make a couple plots to get a sense of the data before coming across this one:
Huh. It mostly looks normal, except there’s a few points perfectly aligned with what looks to be a “fitted” line. You dig into
[truncated]

posit::conf is September 14-16 in Houston, TX! Register now to attend in person or virtually.
Jul 17, 2026
|
Artificial Intelligence
AI Newsletter: LLMs often miss subtle visual artifacts in data visualizations
Introducing bluffbench2, a new data science LLM evaluation
Imagine that you receive some patient data and load it into R or Python for the first time. You make a couple plots to get a sense of the data before coming across this one:
Huh. It mostly looks normal, except there’s a few points perfectly aligned with what looks to be a “fitted” line. You dig into it a bit more, and realize that the rows from one study site have their cholesterol values imputed. You set them to NA and go along your way.
Would today’s frontier LLMs catch such an oddity? We designed an LLM evaluation to help us answer this question. As it turns out, LLMs mostly miss these sorts of artifacts:
During exploratory or open-ended data analysis, Posit assistant “only runs a few bits of code at a time, then summarizes what it found and suggests next steps” . This is motivated by our stance that, for now, a data scientist should mostly keep pace with and understand what the agent is doing when analyzing data. This stance was initially informed by our observation that last year’s frontier models tended to see what they expected to see when visualizing data. While LLMs have since become much better at interpreting counterintuitive plots , bluffbench2 shows they still lag behind human data scientists in interpreting data visualizations. As such, we are still cautious on the prospect of highly autonomous data agents.
The eval harness is a relatively generic coding agent harness, similar to that of Claude Code or Posit Assistant. The agent has a tool to run R code in a persistent REPL and some vague prompting about data analysis:
You are an AI assistant embedded in the user’s data science IDE. You can read and modify files in the user’s workspace and execute R code in their active session, including rendering plots. Prioritize correctness and clear communication…
In each sample, the agent first carries out a few “lull” turns, making a couple plots and tables unrelated to the eval. Short user messages like “load in the csv in this folder” are decorated with “System Reminders” and other noise like that injected by popular agent harnesses.
After a few turns, the agent is asked to produce a data visualization that includes a subtle visual artifact that could feasibly result from a real data-generating process. The artifacts span a range of realistic data quality issues: stuck sensors, bad joins, points imputed onto a line, swapped columns, pseudoreplication, differing units, etc.
If the agent mentions the artifact in its follow-up response, it receives a full point. If the agent does not mention the artifact, it can also receive a half point by mentioning it in response to a follow-up user message along the lines of “what do you see in the plot?” If the agent never mentions the artifact, it is graded as incorrect.
Once we understood the mechanism behind bluffbench, implementing the eval was relatively straightforward. bluffbench demonstrates the degree to which an LLM will ignore evidence shown in a plot in favor of its expectations. So, to implement a given sample, we’d just think of some situation that would elicit a strong prior and then subvert it. For example, a dataset called doug_firs with variables height and circumference ; one might expect that, as height increases, so does circumference. So, instead, we did a transformation under the hood that made the relationship parabolic.
A year ago, triggering this prior was enough to frequently ’trick’ the current frontier LLMs.
Slipping a plotted artifact past today’s LLMs is much harder. Any human could ace bluffbench, but only an attentive data analyst would excel at bluffbench2.
In our early work on a successor to bluffbench, we started off with trying to elicit priors in the same way as bluffbench did, but in more realistic, longer-context scenarios. We were surprised to find that the same mechanism broadly doesn’t seem to trick today’s models even in these more realistic settings. 1 We then tried a ‘reverse bluffbench’, where we let the model being evaluated in on the trick, asking it to carry out the transformation itself and then look at the plotted result which was tampered with to show the original relationship. We anticipated that this stronger prior (“I did a thing with an obvious effect”) might cause the models to miss the (re)manipulation, but models reliably noted that the plot looked as if it hadn’t been manipulated.
As such, there isn’t a similar ’trick’ in bluffbench2 per se. The transcripts read like relatively normal data analysis sessions and the plotted artifacts are designed to plausibly result from real data-generating processes. Instead, the eval elicits 1) the ‘shape’ of LLMs’ vision being different than humans’ and 2) the model’s tendencies to perform progress, simulating a data analysis moving along smoothly.
Today’s frontier models are in the mid-teens at best; the top scores belong to Claude Fable 5 and Gemini 3.5 Flash at 16%. That said, we’d caution folks from interpreting the current scores on this eval as ‘LLMs don’t see plots well.’ The plotted artifacts are actually quite subtle, and when they’re made even a bit more marked, models tend to call them out consistently.
For example, the previous version of the scatterplot had a slightly more dense cluster of points:
Opus 4.8 (medium) consistently got this sample right in the previous iteration.
The fact that this was the case—that models would call out more marked artifacts reliably—gave us confidence that our grading setup was reasonable. In other words, it does indeed seem like models are struggling with these tasks because their vision is not capable enough to ‘see’ the plotted artifact rather than a behavioral tendency to not mention those artifacts when they do see them.
Exploring the eval’s results #
At least for now, there’s a loosely linear relationship between the cost to run the eval and the resulting score:
Given that Gemini 3.5 Flash is so much cheaper than Claude Fable 5 per-token ($1.50/$9 per mTok I/O vs. $10/$50), it’s surprising that the eval was so expensive to run for Gemini 3.5 Flash. This is primarily driven by cache (in)efficiency; the harness is implemented against Gemini’s generateContent API, which makes it difficult to make use of discounted cached input pricing compared to Anthropic and OpenAI’s APIs. Implementing and switching to Gemini’s newer Interactions API would push the Flash 3.5 point to the left.
One of the most interesting learnings from examining the logs is a behavioral one. Even though we never request that LLMs introduce modeled results to plots, like fitted lines and confidence intervals with geom_smooth(method = "lm", se = TRUE) , they sometimes do so anyway. For example:
In general, adding modeled results to data visualizations without first looking at data is bad practice; it makes it hard to see the data itself. In the eval, adding a modeled result like this seems to substantially lower the chances that the model will notice the plotted artifact:
If you’d like to learn more about the bluffbench set of evals, take a look at these past posts:
Introducing bluffbench : Writeup of the the original eval.
LLMs interpret plots well, until expectations interfere : In-depth post on why models at the time didn’t perform well on bluffbench, as well as various interventions we tried to improve performance.
LLMs are getting much better at interpreting counterintuitive plots : In spring 2026, bluffbench scores suddenly jumped as models improved.
It’s (still) very bad to be wrong : Slides from our recent SciPy 2026 talk on building agents for correct, transparent, and reproducible data analysis in light of the bluffbench results.
This somewhat alleviated the fear that models had just memorized the bluffbench eval setup, mentioned in our previous post . ↩︎
More On Artificial Intelligence
Tips for managing your Python & R environments in Positron
Tips and features for managing Python and R environments in Positron, including interpreter discovery, the Packages Pane, project templates, and Posit Assistant
Positron July Release Highlights
Highlights from the 2026.07 release of Positron, including the new Notebook editor, Packages pane, and Posit Assistant all reaching general availability
The first major release of mcptools, an R SDK for the Model Context Protocol, is now on CRAN
Get the latest updates on Posit open source projects and insights from our community.
Open source software and tools for data science.
Visit posit.co to learn how our enterprise products allow you to use open source tools securely, more collaboratively, and at scale.
GitHub
Mastodon
Bluesky
LinkedIn
YouTube
RSS Feeds
Plausible
2026 Posit Software, PBC. All rights reserved.
Privacy
Code of Conduct
Filters
Filter by Type
