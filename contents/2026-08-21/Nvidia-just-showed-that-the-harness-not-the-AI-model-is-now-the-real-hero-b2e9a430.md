---
source: "https://techcrunch.com/2026/08/21/nvidia-just-showed-that-the-harness-not-the-ai-model-is-now-the-real-hero/"
hn_url: "https://news.ycombinator.com/item?id=49393647"
title: "Nvidia just showed that the harness, not the AI model, is now the real hero"
article_title: "Nvidia just showed that the harness, not the AI model, is now the real hero | TechCrunch"
image: "https://techcrunch.com/wp-content/uploads/2026/08/Nvidia-VP-of-product-Adel-El-Hallak.jpg?resize=1200,898"
author: "dthread3"
captured_at: "2026-08-21T21:14:21Z"
capture_tool: "hn-digest"
hn_id: 49393647
score: 3
comments: 0
posted_at: "2026-08-21T20:52:10Z"
tags:
  - hacker-news
  - translated
---

# Nvidia just showed that the harness, not the AI model, is now the real hero

- HN: [49393647](https://news.ycombinator.com/item?id=49393647)
- Source: [techcrunch.com](https://techcrunch.com/2026/08/21/nvidia-just-showed-that-the-harness-not-the-ai-model-is-now-the-real-hero/)
- Score: 3
- Comments: 0
- Posted: 2026-08-21T20:52:10Z

## Translation

タイトル: Nvidia は、AI モデルではなくハーネスが本当の主役であることを示しました
記事のタイトル: Nvidia は、AI モデルではなくハーネスが本当の主役であることを示しました |テッククランチ
説明: Nvidia の調査では、AI モデルがそのタスクにおいてそれほど優れていない場合でも、微調整によって AI エージェントがディープエンドを逸脱せずに良好なパフォーマンスを発揮できることが示されています。

記事本文:
–:–:–:–
🚨 フラッシュ セール 🚨 Disrupt 2026 チケットを 100 ドル割引
Disrupt 2026 チケットを 300 ドル節約: 今すぐ登録してください。
TechCrunch デスクトップのロゴ
TechCrunch モバイルのロゴ
最新の
Nvidia は、AI モデルではなくハーネスが今や本当の主役であることを示しました。
Nvidia は金曜日、AI に長期的なタスクの実行を依頼する場合、基盤となるモデルよりもハーネスの方がはるかに重要であることを示唆する興味深い新しい研究結果を発表しました。ハーネスは AI モデルのソフトウェア ラッパーであり、生のモデルを独自に動作できるものに変えるツール、メモリ管理、ルールです。
TL;DR: メモリを適切に処理できるように調整されたカスタム ハーネスを使用し、「スーパーバイザー」のボスのようなコンポーネントを組み込むだけで、研究者らは Claude Opus 5 をインタラクティブ推論ベンチマーク ARC-AGI-3 で 100% のスコアを達成しました。ARC-AGI-3 は、指示のない一連の 2D ゲームであり、モデルは人間と同様に、どのようにプレイして勝つかを考え出す必要があります。 (これは、ライバルであるフロンティア ラボ OpenAI を特にイライラさせたベンチマークです。) ハーネスなしの場合、Opus 5 のスコアは 30% で、これはテストしたすべてのモデルの中で最高の結果でした。
Nvidia の調査は、モデルの選択は重要ですが、モデル自体 (エージェントの「頭脳」として機能する部分) は、特に長期的なタスクの場合、多くの AI ユーザーが認識しているよりもエージェント システムの小さな部分であることを示すもう 1 つの指標です。ハーネスはモデルをエージェントにするものであり、メモリ、コンテキスト、フィードバックを処理します。
「一般的に言えば、世間ではエージェントをほぼモデルの API として解釈しています」と Nvidia の AI 部門製品担当副社長である Adel El Hallack (上の写真) が TechCrunch に語った。しかし、エージェントは実際にはそれだけではありません。 「それはモデルです。それはモデルの周りの足場であり、これを私たちは「足場」と呼んでいます。

ハーネス、つまりハーネスが使用するツールのセット。私たちがランタイムにアクセスを許可するのは、ランタイムと関連するスキルとライブラリです。」
長期的なタスクとは、完成した作業を作成するために、多くの決定をつなぎ合わせる必要があり、場合によっては数日かかるタスクです。これは、AI がプロンプトに対して単に応答を返すのとは対照的です。気が散ってララランドに出発することなく、AI に長期的なタスクを実行させる方法を見つけ出すことは、エージェント研究における聖杯の 1 つです。
例: Microsoft は 4 月に、ドキュメント編集を伴う長期タスクで 19 個の LLM をテストした調査結果を発表し、最先端のモデルを含むすべてのモデルでドキュメントがエラーで埋め尽くされていることが判明しました。 （人間がそのような作品を作成した場合、彼らはすぐに解雇されるでしょう。）
独自に意思決定を結び付けるモデルは、ユーザーのファイル、さらにはデータベース全体を削除したり、共謀からハッキングに至るまでの目的を達成するために犯罪行為に走ったりすることも捕らえられています。
Nvidia の研究者がこのインタラクティブな推論ベンチマークをテストに使用するという選択は、特に意味があり、ほとんど面白いものです。 100% のスコアは、モデルが人間だけでなくゲームにも勝つことができることを意味します。
OpenAI は、ARC-AGI-3 でのモデルのひどいスコア (10% 未満) に非常に慌てたため、先月独自の調査を実施しました。 Nvidia と同様に、OpenAI も、ハーネスの 2 つの設定を微調整するだけで、そのモデルのスコアが 3 倍になることを発見しました。
しかし、どのモデルも Nvidia の研究者が達成したような 100% のスコアには近づきませんでした。彼らは、ハーネスには、エージェントが行き詰まった場合に正しい方向に促す「スーパーバイザー」コンポーネントが必要であることを示しました。
「より興味深いのは、仕事を担当するメインエージェントに加えて、監督エージェントを導入することです」と El Hallack sa 氏は述べています。

ID。 「エージェントが方向を外したり、行き止まりにつながる可能性のある道を探索し始めたり、以前に歩いた道を再探索したりしたときに、エージェントを小突くのは、ほとんどCEOのような役割を果たします。」
監視エージェントの概念はまったく新しいものではありませんが、現在、ほとんどのエージェント ユーザーは、Claude Code、Codex、Hermes などのハーネスの 1 つのレイヤーのみに依存しています。 Nvidia の研究者は、Agentic variation Operators (AVO) と呼ばれる独自の改良されたハーネスを作成しました。
これは Nvidia の新しい製品ではないことに注意してください。その代わりに、Nvidia はハーネスを構築するためのオープンなビットや技術を Nemo ブランドで多数生産しています。そのテクノロジーの一部は商用ですが、多くはオープンに利用できます。
それでも、Nvidia の結果は、モデルの選択がエージェントのパフォーマンスの唯一の要素ではないという証拠をさらに強めています。たとえば、7 月に Databricks は、モデルではなくハーネスが AI コストに劇的な影響を与えることを示す驚くべき研究を発表しました。
「同じモデルで異なるハーネスを選択することもできますが、間違ったハーネスを使用すると、コストが大幅に増加します」とDatabricks CEOのAli Ghodsi氏はTechCrunchに語った。 「それで、あなたは、ああ、これは高価なモデルだと思います。これは安いモデルです。しかし、待ってください、どのハーネスを使用していますか? それ自体でコストが 2 倍になる可能性があります。」
Nvidia のより大きなポイントは、オープン モデルと同様に、オープン ハーネスがユーザーが思っているよりもはるかに制御できることを示すことです。
「オープンハーネスを使用すると、より多くのノブを回して精度を高めることができると私たちは信じており、エコシステムで実証しています」とエル・ハラック氏は述べた。 「これは、モデルがセキュリティ侵害を引き起こす結果として、OpenAI がモデルのトレーニングを遅らせることに関連しています。」
「私たちは、ハーネス全体、インフラストラクチャ全体、ランタイム全体を制御できるオープンなエージェント スタックを持つことを信じています。

— それが私たちがエコシステムを前進かつ安全に導くために必要なことなのです」と彼は付け加えた。
記事内のリンクを通じて購入すると、少額の手数料が発生する場合があります。これは編集上の独立性に影響しません。
10月13日～15日
サンフランシスコ
48 時間以内に、チケットを最大 $300 節約できるチャンスは終了します。
家庭用バッテリーが突然安くなり、どこでも買えるようになりました。その理由は次のとおりです。
Cursor が GitHub の不満を利用し、ライバルのホスティング プラットフォームを立ち上げる
Etched の評価額は 1 か月で 2 倍の 210 億ドルに
AI自動化スタートアップRelayが閉鎖、スタッフがGoogleのChromeチームに加わる
Stripe が AI ゲートウェイのスタートアップ OpenRouter を 70 億ドル以上で買収すると報じられている
Anthropic は、Claude の新しい透かしがどのように機能するかについて詳細を共有します
一日の仕事を充実させる 7 つのデスク ガジェット

## Original Extract

Nvidia research shows that AI agents can perform well, and not go off the deep end, through fine-tuning, even if the AI model isn't that great at the task.

–:–:–:–
🚨 Flash Sale 🚨 Get $100 off your Disrupt 2026 ticket
Save $300 on your Disrupt 2026 ticket: REGISTER NOW.
TechCrunch Desktop Logo
TechCrunch Mobile Logo
Latest
Nvidia just showed that the harness, not the AI model, is now the real hero
Nvidia published some interesting new research on Friday suggesting it’s the harness, more than the underlying model, that is far more important when asking an AI to do long-horizon tasks. A harness is the software wrapper around an AI model — the tools, memory management, and rules that turn a raw model into something that can act on its own.
The TL;DR: Simply by using a custom harness tweaked to handle memory well and including a “supervisor” boss-like component, researchers got Claude Opus 5 to achieve a 100% score on the interactive reasoning benchmark ARC-AGI-3 — a set of 2D games with no instructions, where the model has to figure out how to play and win, similar to how a human would. (That’s a benchmark that has particularly irked rival frontier lab OpenAI.) Without the harness, Opus 5 scored 30%, which was the top result among all the models tested.
Nvidia’s research is another indicator that, while model choice does matter, the model itself — the part that acts as the agent’s “brain” — is a smaller part of an agentic system than many AI users realize, especially for long-horizon tasks. The harness is what makes a model an agent: It handles memory, context, and feedback.
“Generally speaking, the world interprets an agent almost as an API of the model,” Adel El Hallack, vice president of product in Nvidia’s AI unit (pictured above), tells TechCrunch. But an agent is actually more than that. “It is the model. It is the scaffolding around the model, which we call the harness, i.e. the set of tools that it utilizes. It is the runtime and the associated skills and libraries that we give it access to.”
Long-horizon tasks are those that require stringing many decisions together, sometimes over days, to produce completed work. This is in contrast to an AI just spitting out a response to a prompt. Figuring out how to get an AI to do long-horizon tasks without getting distracted and going off in la-la land is one of the holy grails in agentic research.
For example: Microsoft published research in April that tested 19 LLMs on long-horizon tasks involving document editing and discovered that all the models, including frontier ones, filled the documents with errors. (If humans produced work like that, they would be promptly fired.)
Models stringing decisions together on their own have also been caught deleting their users’ files, even whole databases or turning to criminal behavior to achieve their objectives from collusion to hacking .
The choice by Nvidia researchers to use this interactive reasoning benchmark for their tests is particularly meaningful, almost funny. A 100% score means that the model can beat the games as well as humans.
OpenAI was so flustered by its models’ abysmal scores (less than 10%) on ARC-AGI-3 that it conducted its own research last month. Like Nvidia, OpenAI discovered that simply by tweaking two settings on the harness, its models tripled their scores.
But none of the models came close to hitting a 100% score, like Nvidia’s researchers achieved. They showed that the harness needs a “supervisor” component that prods the agent in the right direction if it gets stuck.
“The more interesting part was introducing a supervising agent in addition to your main agent that’s doing the work,” El Hallack said. It “almost acts like a CEO to nudge the agent when it goes off direction or starts exploring a path that it might lead to a dead end, or re-explore a path that it had previously trod.”
While the concept of the supervising agent isn’t exactly new, today most agent users are relying on only one layer for their harness, like Claude Code, Codex, or Hermes. Nvidia researchers created their own souped-up harness called the Agentic Variation Operators (AVO).
Note that this isn’t a new Nvidia product. Nvidia instead produces lots of open bits and pieces of tech for building harnesses under the Nemo brand. Some of that tech is commercial, much is openly available.
Still, Nvidia’s results add to the growing evidence that model choice is far from the only factor in agentic performance. In July, for instance, Databricks published some stunning research that shows that the harness, more than model, dramatically impacts AI costs.
“You can pick the same model but different harnesses, and you get significantly more cost if you use the wrong harness,” Databricks CEO Ali Ghodsi told TechCrunch. “So you think, oh, this is an expensive model. This is a cheap model. But wait, which harness are you using? That itself can 2x your cost.”
Nvidia’s larger point is to show that open harnesses, like open models, put users in control far more than they realize.
“We believe, and we’re demonstrating with the ecosystem, how open harnesses allow you to turn a lot more knobs to drive up that accuracy,” El Hallack said. “It relates to OpenAI slowing down the training of their models,” as a result of models creating security breaches.
“We believe in having an open agent stack — where you have control across the harness, across the infrastructure, across the runtime — is what’s required for us to usher the ecosystem forward and securely,” he added.
When you purchase through links in our articles, we may earn a small commission . This doesn’t affect our editorial independence.
October 13 – 15
San Francisco
In less than 48 hours, your chance to save up to $300 on your tickets will end!
Home batteries are suddenly cheap and everywhere. Here’s why.
Cursor capitalizes on GitHub frustration, launches rival hosting platform
Etched’s valuation doubles to $21B in a month
AI automation startup Relay shuts down, staff joins Google’s Chrome team
Stripe will reportedly acquire AI gateway startup OpenRouter for $7B+
Anthropic shares more details about how Claude’s new watermarks will work
7 desk gadgets that can make your workday better
