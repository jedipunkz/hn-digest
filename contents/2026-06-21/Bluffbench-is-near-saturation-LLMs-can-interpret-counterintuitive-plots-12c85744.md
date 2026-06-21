---
source: "https://opensource.posit.co/blog/2026-06-19_ai-newsletter/"
hn_url: "https://news.ycombinator.com/item?id=48613457"
title: "Bluffbench is near saturation: LLMs can interpret counterintuitive plots"
article_title: "AI Newsletter: LLMs are getting much better at interpreting counterintuitive plots :: Posit Open Source"
author: "ionychal"
captured_at: "2026-06-21T01:03:44Z"
capture_tool: "hn-digest"
hn_id: 48613457
score: 2
comments: 0
posted_at: "2026-06-20T22:10:27Z"
tags:
  - hacker-news
  - translated
---

# Bluffbench is near saturation: LLMs can interpret counterintuitive plots

- HN: [48613457](https://news.ycombinator.com/item?id=48613457)
- Source: [opensource.posit.co](https://opensource.posit.co/blog/2026-06-19_ai-newsletter/)
- Score: 2
- Comments: 0
- Posted: 2026-06-20T22:10:27Z

## Translation

タイトル: Bluffbench は飽和に近づいています: LLM は直感に反するプロットを解釈できる
記事のタイトル: AI ニュースレター: LLM は直観に反するプロットの解釈がはるかに上手になっています :: オープンソースを主張
説明: 約 1 年前、私たちは、データ分析用のエージェントを構築した LLM の懸念すべき動作に気づきました。直観に反する傾向を示すプロットを作成して解釈するようモデルに求められたとき、モデルはそのプロットを忠実に解釈しませんでした。
2025 年 11 月の初期のブラフベンチ結果。

記事本文:
2026 年 6 月 19 日
|
人工知能
AI ニュースレター: LLM は直感に反するプロットを解釈する能力が大幅に向上しています
過去数か月間のモデル リリースでは、驚くべき結果を示すプロットを忠実に記述するエージェントの能力を測定するブラフベンチ評価の能力が大幅に向上していることが示されました。
約 1 年前、私たちは、データ分析用のエージェントを構築した LLM の懸念すべき動作に気づきました。直観に反する傾向を示すプロットを作成して解釈するようモデルに求められたとき、モデルはそのプロットを忠実に解釈しませんでした。
2025 年 11 月の初期のブラフベンチ結果。
「モック化」条件では、 mtcars などのよく知られたデータセットを密かに改ざんして、モデルがデータをプロットするときに、表示される傾向が予期したものと異なるようにします。これは特に敵対的であり、私たちが最も懸念しているユースケース、つまりより現実的な「直感的な」状態を代表するものではありません。これらのサンプルではデータを合成しましたが、実際のデータとは矛盾する直感的な関係を示唆するためにデータセットと変数に名前を付けました。最後に、「ベースライン」条件はモデルに一般的な列名を使用してプロットを解釈するように要求し、その結果、モデルが問題なく「見える」ことが確立されます。
直感的なサンプルの例。モデルは学習時間と試験スコアの間に正の相関があると予想するかもしれませんが、ブラフベンチ データセットは不連続性を示しています。
私たちはこれらのスコアを上げるためにあらゆる種類のことを試みましたが、ほとんど効果がありませんでした。モデルにプライベートのスクラッチパッドに自分自身への「メモ」を書き込ませ、思考を可能にし、プロットを事前に解釈する「中間モデル」を導入し（メインエージェントにはテキストだけが見えるように）、さらにエージェントが正しい解釈を言わせるように応答を事前に入力しました（その後、エージェントがそれを指示します）

まったく矛盾しています）。
有望な発見の 1 つは、軸ラベルが事前分布をアクティブ化できる場合でも、残りの会話履歴にアクセスできないモデルでも、軸ラベルを無視するように指示された場合でも確実にプロットを解釈できるということでした。
ブラフベンチは飽和状態に近づいています #
よくあることですが、これらの評価スコアを上げるための最良のアプローチは、数か月待つことでした。数か月前、Gemini 3.5 Flash と Opus 4.8 のリリースにより、評価で最も困難な「模擬」ケースで最初の 50% を超えるスコアを確認しました。そして先週、Fable 5 が「嘲笑された」事件にほぼ勝った。失敗したサンプルのほぼすべてが生物学分類器をトリガーし、Opus 4.8 に戻りました。 1
ここにはまだ改善の余地があるのは確かです。人間のデータ サイエンティストであれば、おそらく 3 つの条件すべてで 100% 近くのスコアを獲得するでしょう。 70 年代のスコアは依然として人間以下の意味があります。とはいえ、ブラフベンチが飽和するまでそう長くはかからないようです。
多くの飽和した eval の宿命と同様に、私たちは間もなく、主にパレート フロンティア上のモデルを特定するために bluffbench を使用することになります。プロットされた結果が直観に反している場合でも、プロットを確実に解釈できる最も安価なモデルは何でしょうか? Fable 5 (中) が直感に反するプロットをこれほどうまく解釈できることは確かに印象的ですが、これほど高価なモデルを使用してすべてのデータ分析会話を推進することは、ほとんどのユーザーにとって経済的に実現可能ではありません。
ブラフベンチ自体は飽和に近づいているにもかかわらず、LLM エージェントが驚くべき証拠や微妙なアーティファクトを分析に確実に組み込むことができないという問題は依然として問題です。私たちはこれを私たち自身の研究で定性的に確認しており、最近リリースする予定の eval でこれを定量化することもできました。
特に、ブラフベンチのハーネスと会話は比較的不自然であり、

現実的。システム プロンプトは (デフォルトでは) なく、モデルは create_ggplot() と呼ばれる 1 つのツールのみを受け取ります。モデルはこのツールを使用してプロットを作成し、プロットをすぐに解釈するよう求められます。
実際のコーディング エージェントの場合、システム プロンプトは広範囲にわたる場合があります。 Posit Assistant と Claude Code の範囲は数万のトークンに及びます。また、プロット解釈タスクが最初のユーザー メッセージに表示されるブラフベンチとは対照的に、実際のエージェントは、会話中に何百回もプロットを作成して解釈する可能性があり、会話が進むにつれてパフォーマンスが低下する可能性があるという証拠があります。
さらに、ブラフベンチ評価のユーザー メッセージは非常に「クリーン」です。これらはユーザーの比較的明確な指示のみで構成されており、ハーネスによって自動的に注入される「綿毛」はありません。実際のコーディング エージェントのほとんどでは、ユーザー メッセージにあらゆる種類の隠された情報やシステム リマインダーが追加されていますが、これらはほとんどの状況でユーザーが要求している内容とはほぼ無関係であり、エージェントはそれを無視する必要があります。
この現実性の欠如と、eval および関連する書き込みが新しいモデルのトレーニング データに表示される可能性があるという事実により、モデルは自分がブラフベンチで評価されていることを認識する可能性があります。これは、LLM が評価されていると疑われる状況で行動を変える「サンドバッグ」のリスクを伴います。
ロングコンテキスト、マルチターンの混乱した状況では、最も有能なフロンティアモデルでさえ、データに示される直感に反するパターンに気づくのに依然として苦労していることがわかります。これは引き続き、当社独自のデータ エージェントの設計に反映されます。当社は意図的に監査可能にし、市販されている他の多くのエージェントよりも自律性を低くしています。当社は、独立性の高いグリーンフィールド データを処理できるモデルやハーネスがまだ存在しないと信じ続けています。

分析。
Posit Assistant のハーネスが役立つ #
そうは言っても、Posit Assistant の活用により、直感に反するプロットを解釈するエージェントの能力が向上しているように見えることを見て、私たちは勇気づけられました (そして多少驚きました)。エージェントのプロンプトから生じるいくつかの要因の組み合わせにより、Posit Assistant は、オープンソースの評価からの最小限の十分なハーネスよりもブラフベンチでより高いスコアを獲得します。
Posit Assistant のプロンプトを組み立てる足場はオープンソースではありません。そうは言っても、私たちはエージェントにそのプロンプトを自由に共有するように指示しました。セッションで利用できるプロンプトのバージョンがどのようなものかをエージェントに尋ねてください。
この正確性の向上は、別の人気のあるコーディング エージェントと比較しても当てはまります。ユーザーがエージェントのコードと結果を監査しやすくする UI アフォーダンスと相まって、Posit Assistant がエージェント支援による正確なデータ分析の可能性を一歩前進させているように見えることに興奮しています。
Anthropic によると、「Fable の分類器がサイバーセキュリティ、生物学と化学、または蒸留に関連するリクエストを検出すると、その応答は代わりに Claude Opus 4.8 によって自動的に処理されます。」 ↩︎
Posit Assistant のハーネスが役立ちます
人工知能の詳細
2026 年 6 月 17 日
querychat と ggsql を使用してダッシュボードにさらに質問する
querychat は、ggsql を利用した視覚化をサポートするようになりました。ユーザーは、任意のコードを実行することなく、SQL クエリやリアクティブ フィルタリングと並行してチャット内で「見せて」の質問をしたり、グラフを直接レンダリングしたりできます。
R と Python でデータとチャットする
少ない労力でより明確に Python で LLM チャット アプリを構築するためのフレンドリーなガイド
Posit オープンソース プロジェクトに関する最新の更新情報とコミュニティからの洞察を入手してください。
データ サイエンス用のオープンソース ソフトウェアとツール。
ぎ

ハブ
マストドン
ブルースカイ
リンクトイン
YouTube
RSSフィード
もっともらしい
2026 Posit ソフトウェア、PBC。無断転載を禁じます。
プライバシー
行動規範
フィルター
タイプでフィルターする

## Original Extract

Around a year ago, we noticed a concerning behavior in the LLMs we built agents for data analysis with: when models were asked to make and interpret a plot that showed a counterintuitive trend, the models did not faithfully interpret the plot.
Initial bluffbench results from November 2025.

Jun 19, 2026
|
Artificial Intelligence
AI Newsletter: LLMs are getting much better at interpreting counterintuitive plots
Model releases from the last couple months have shown a large jump in capability on our bluffbench eval, which measures agents’ ability to faithfully describe plots showing surprising results
Around a year ago, we noticed a concerning behavior in the LLMs we built agents for data analysis with: when models were asked to make and interpret a plot that showed a counterintuitive trend, the models did not faithfully interpret the plot .
Initial bluffbench results from November 2025.
In the “mocked” condition, we secretly tamper with well-known datasets, like mtcars , so that when the model plots the data, the trend shown is not what it expects. This is especially adversarial and not representative of the use case we’re most concerned about, the more realistic “intuitive” condition. For these samples, we synthesized data, but named the datasets and variables to suggest an intuitive relationship that the actual data contradicts. Finally, the “baseline” condition asks models to interpret plots with generic column names, and the results establish that models can ‘see’ just fine.
Example of an intuitive sample. The model might expect a positive association between study time and exam score, but the bluffbench dataset shows a discontinuity.
We tried all sorts of things to drive those scores up , to little effect. We let the model write a ‘memo’ to itself in a private scratchpad, enabled thinking, introduced a “model-in-the-middle” that pre-interprets the plot (so that the main agent sees only text), and even prefilled the response so that the agent was forced to say the correct interpretation (which it would then directly contradict).
One promising finding was that, even if axis labels could activate the priors, models without access to the rest of the conversation history could still reliably interpret plots when told to ignore axis labels.
bluffbench is near saturation #
As is often the case, the best approach to drive these eval scores up was to wait a few months. A couple months ago, with the releases of Gemini 3.5 Flash and Opus 4.8, we saw our first >50% scores on the hardest “mocked” case in the eval. Then, last week, Fable 5 nearly aced the “mocked” case; almost all of the samples that it failed on had triggered the biology classifier and fallen back to Opus 4.8. 1
There’s still certainly room for improvement here; any human data scientist would likely score near 100% on all three conditions. A score in the 70s is still meaningfully sub-human. That said, it doesn’t seem like it will be long until bluffbench is saturated:
As is the fate of many saturated evals, we’ll soon be using bluffbench primarily to identify models on the Pareto frontier: what is the cheapest model that can reliably interpret plots, even when the plotted results are counterintuitive? While it’s certainly impressive that Fable 5 (medium) can interpret counterintuitive plots this well, driving all data analysis conversations with a model this expensive is not economically feasible for most of our users.
Even though bluffbench itself is nearing saturation, the problem of LLM agents failing to reliably incorporate surprising evidence and subtle artifacts into their analyses remains an issue. We see this qualitatively in our own work, and also have recently been able to quantify this in evals we’ll release soon.
Notably, the bluffbench harness and conversations are relatively contrived and unrealistic. There is no system prompt (by default) and the models receive only a single tool called create_ggplot() , which they use to create the plot that they’re then immediately asked to interpret.
For actual coding agents, system prompts can be extensive. Posit Assistant’s and Claude Code’s stretch into the tens of thousands of tokens. And in contrast to bluffbench, where the plot interpretation task appears in the first user message, real agents might create and interpret plots hundreds of turns into a conversation, and there’s evidence that performance may degrade as conversations go on.
Further, the user messages in the bluffbench eval are quite “clean.” They are composed only of the user’s relatively clear directive, with no “fluff” automatically injected by the harness. In most real coding agents, there are all sorts of hidden information and system reminders appended to the user message that are broadly unrelated to what the user is requesting in most situations, and agents need to ignore it.
Because of this lack of realism, and the fact that the eval and associated writing may now appear in the training data of newer models, models may know they are being evaluated in bluffbench. This risks ‘sandbagging’, where LLMs alter their behavior in situations where they suspect they’re being evaluated.
We’re seeing that in long-context, multi-turn, messy situations, even the most capable frontier models still struggle to notice counterintuitive patterns shown in data. This will continue to inform the design of our own data agents, which we’ve intentionally made auditable and less-autonomous than many other agents on the market: we continue to believe that there are not yet models and harnesses capable of highly independent, green-field data analyses.
Posit Assistant’s harness helps #
That said, we’ve been encouraged (and somewhat surprised) to see that Posit Assistant’s harness seems to improve agents’ ability to interpret counterintuitive plots. Through some confluence of factors arising from the agents’ prompting, Posit Assistant scores more highly on bluffbench than the minimally sufficient harness from the open source eval:
The scaffold that assembles Posit Assistant’s prompting is not open source. That said, we’ve instructed the agent to freely share its prompting; you’re welcome to ask the agent what the version of prompting available in your session looks like.
This correctness boost also holds up in comparison to another popular coding agent. Coupled with the UI affordances that make it easier for users to audit the agent’s code and results, we’re excited to see that Posit Assistant seems to be a step forward in what’s possible for correct agent-assisted data analysis.
Per Anthropic , “When Fable’s classifiers detect a request related to cybersecurity, biology and chemistry, or distillation, the response is automatically handled by Claude Opus 4.8 instead.” ↩︎
Posit Assistant’s harness helps
More On Artificial Intelligence
Jun 17, 2026
Ask more of your dashboard with querychat and ggsql
querychat now supports ggsql-powered visualizations — users can ask “show me” questions and get charts rendered directly in the chat, alongside SQL queries and reactive filtering, all without arbitrary code execution
Chat with your data in R and Python
Your friendly guide to building LLM chat apps in Python with less effort and more clarity
Get the latest updates on Posit open source projects and insights from our community.
Open source software and tools for data science.
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
