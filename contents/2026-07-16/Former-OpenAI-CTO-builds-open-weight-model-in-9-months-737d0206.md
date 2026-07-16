---
source: "https://techcrunch.com/2026/07/15/thinking-machines-amps-up-its-bet-against-one-size-fits-all-ai-with-its-first-open-model-inkling/"
hn_url: "https://news.ycombinator.com/item?id=48935358"
title: "Former OpenAI CTO builds open weight model in 9 months"
article_title: "Thinking Machines amps up its bet against one-size-fits-all AI with its first open model, Inkling | TechCrunch"
author: "noashavit"
captured_at: "2026-07-16T15:27:22Z"
capture_tool: "hn-digest"
hn_id: 48935358
score: 1
comments: 1
posted_at: "2026-07-16T14:47:21Z"
tags:
  - hacker-news
  - translated
---

# Former OpenAI CTO builds open weight model in 9 months

- HN: [48935358](https://news.ycombinator.com/item?id=48935358)
- Source: [techcrunch.com](https://techcrunch.com/2026/07/15/thinking-machines-amps-up-its-bet-against-one-size-fits-all-ai-with-its-first-open-model-inkling/)
- Score: 1
- Comments: 1
- Posted: 2026-07-16T14:47:21Z

## Translation

タイトル: 元 OpenAI CTO が 9 か月で無差別ウェイト モデルを構築
記事のタイトル: Thinking Machines は、最初のオープン モデル Inkling で、万能 AI に対する賭けを強化します。テッククランチ
説明: これは、ほぼ公の場で AI インフラストラクチャの構築に 1 年半を費やした同社の最初の公開実証ポイントです。

記事本文:
TechCrunch デスクトップのロゴ
TechCrunch モバイルのロゴ
最新の
Thinking Machines は、初のオープン モデル Inkling で画一的な AI に対する賭けを強化します。
元 OpenAI CTO のミラ・ムラティ氏が設立した AI スタートアップである Thinking Machines Lab は、水曜日の朝、Inkling と呼ばれる最初の社内 AI モデルをリリースしました。また、OpenAI、Anthropic、Google の主力モデルとは異なり、オープンウェイトです。つまり、外部の開発者や企業がダウンロードして直接変更できます。
Inkling は、合計 9,750 億のパラメーターを備えた専門家が混合したシステムですが、特定のタスクに使用するのはその一部 (約 410 億) だけであり、非常に大規模なモデルをより高速かつ安価に実行できる共通の設計となっています。同社のリリース資料によると、テキスト、画像、音声、ビデオの45兆個のトークンと、4つすべての理由をネイティブにトレーニングされたという。ただし今のところ、その出力はコード、スタイル付きアーティファクト、構造化データを含むテキストに限定されています。
このモデルは、Thinking Machines Labs が 1 年半かけてほとんど公の場で AI インフラストラクチャを構築してきた後、初めて公開された実証ポイントです。その取り組みの一部は、「インタラクション モデル」の 5 月の研究プレビューですでに明らかになっていました。これは、一般的なチャットボットのように立ち止まって待つのではなく、聞いて話したり (さらには中断したり) するように設計された AI です。これは、組織が自ら適応できる AI が、現在最大手の研究所が販売している汎用モデルよりも優れたパフォーマンスを発揮するという、スタートアップの背後にある中心的な賭けのテストでもあります。
Inkling は、推測ではなく不確実性にフラグを立てるなど、調整された回答を提供するように設計されており、ユーザーがスピードと引き換えに「思考努力」を上げたり下げたりできるようにします。あるベンチマークでは、Inkling は Nvi の 3 分の 1 のトークンを使用していると同社は述べています。

dia の Nemotron 3 Ultra (最新世代のオープンウェイト モデル) は、同じコーディング パフォーマンスを達成します。
Thinking Machines は、Inkling がクラス最高であるとは主張しません。その最新のブログ投稿では、Inkling は「オープンかクローズかにかかわらず、現在入手可能な総合的なモデルとしては最も強力ではない」と明確に述べられています。その代わりに明らかに目指しているのは、バランスの取れたパフォーマンスです。
そうなると、ターゲットとしているエンタープライズ市場の中で、この製品は本当に誰を対象としているのかという疑問が生じます。 Thinking Machines は今のところ、Inkling を完成品としてではなく出発点として販売しており、組織が同社のモデル カスタマイズ プラットフォームである Tinker を通じて微調整できるものとして販売しています。これは、たとえば、カスタマイズが安全であることを確認する責任があるのは、Thinking Machines ではなく顧客であることも意味します。 (微調整には、機械を扱う本格的な才能が必要です。)
OpenAI、Anthropic、Google は、それぞれ ChatGPT、Claude、Gemini で非常に異なるアプローチを採用しています。これらはいずれも、まず汎用チャットボットとして競合するように構築され、その上にエージェント機能と自律機能が重ねられています。
Thinking Machines が先週公開した投稿は、明らかにこのリリースの背景を意図したものでした。同社はその投稿の中で、1つの企業によって集中的にトレーニングされ、その後定着するAIは、多くの専門知識がそれを保有する人々に固有のものであるため、組織が自らを形成するAIよりも性能が劣ると主張した。
クローズドモデルに対する他の議論も勢いを増しています。日曜日に公開されたブログ投稿の中で、Microsoft CEO のサティア・ナデラ氏（同社は OpenAI と Anthropic の両方に数十億ドルを投資している）は、独自の AI モデルを使用している企業は実質的に 2 回の支払いを行っていると警告しました。1 回目はサブスクリプション費用、もう 1 回目はプロンプトと修正に組み込まれたビジネス知識の引き渡しです。

将来のモデルのバージョンに吸収される可能性があります。
ハグフェイスのCEO、クレム・デラング氏も先週、TechCrunchとの会話で同様の予測を述べた。同氏によると、フロンティアモデルはますます実験や価値の高いタスク用に確保される一方、本番AIの作業のほとんどはプライベートまたはオープンソースの代替モデル、つまりThinking Machinesが構築している正確な分割モデルに移行するという。
Thinking Machines のアプローチに対する最も明確な論拠は、世界最大のヘッジファンドであるブリッジウォーター アソシエイツとの最近のプロジェクトから来ています (この会社は本来、Thinking Machines の投資家ではありません)。両社の研究者は既存のオープンソース モデルを採用し、ブリッジウォーター独自の財務専門知識に基づいてさらにトレーニングしました。この結果は財務推論テストで 84.7% のスコアを獲得し、トップの独自 AI モデルを上回り、実行コストは約 14 分の 1 だったと言われています。ただし、これらの結果は独立した評価ではなく、両社独自の評価によるものです。
いずれにせよ、Thinking Machines はいかに早くここに到達したかを強調している。 OpenAI はその技術を市場に投入して収益を示すまでに約 5 年かかりましたが、Anthropic は約 3 年かかりました。 Thinking Machines は、約 9 か月で同様の結果を達成したと述べています。
Inkling が競合他社のモデルからの出力に基づいてトレーニングされたのではないかと疑問に思う人もいるでしょう。これは業界全体で厳しい視線を集めている「蒸留」として知られています。同社独自の資料によると、端的に言えば、その答えは一部だ。 Thinking Machines は Inkling をゼロから事前トレーニングしましたが、大規模な強化学習が引き継がれる前に、初期のトレーニング後のデータの一部を生成するために、Moonshot AI の Kim K2.5 を含む他のオープンウェイト モデルを使用したと述べています。同社は、次のモデルでは代わりに完全に自己完結型のポストトレーニングを使用すると主張している。
上で

原作側では、Thinking Machines はより警戒されています。同社は3月にNvidiaと提携し、ギガワット規模のVera Rubinコンピューティング能力を展開し、InklingをNvidiaのGB300 NVL72システム上で完全に訓練した。しかし、それらのコストをどのようにカバーする計画なのかは明らかにしておらず、多くの見方では収益は優先事項ではなかった。 （報道によると、500億ドルの資金調達ラウンドが昨年11月にまとまると言われていたが、1月までに停滞していた。同社はそれ以来、資金調達の状況について話すことを拒否している。）
これに関連した疑問は、Thinking Machines の支出が OpenAI や Anthropic の規模に達することがあるのか、あるいはその効率重視のアプローチが経済状況の変化を意味するのかということです。別の言い方をすれば、同社の賭けは、大手ライバル企業のように最終的に支出する必要がまったくなくなることよりも少ないかもしれない。なぜなら、重みが公開されると、OpenAI や Anthropic が販売する従量制アクセスとは異なり、重みをダウンロードする人に、それを実行するための Thinking Machines への支払いを義務付けるものはないからである。モデルそのものではなく、Tinker です。会社の収益は、トレーニング、微調整、そして現在は、Tinker を中心に構築されているホスティング エコシステムの一部によってもたらされる必要があります。
少なくとも従業員数はより安定しているように見える。 Thinking Machines の従業員数は現在約 200 名で、1 月に OpenAI に退職した共同創設者 2 名を含む、今年初めの相次ぐ退職後に報告された水準よりも増加しています。
Thinking Machines は、業界の多くがしているように、個々の動きを大々的に取り上げることには興味がないようです。社内の情報筋によると、同社の文化は設計上、特定の個性に依存するよりも継続性を重視しているとのことです。それは当然のことです。最初から台座に乗せられていなかった人は、チームを変更するときに挫折することが少なくなります。それも特筆すべきことです

同社自身の物語の多くが、今では有名になった共同創設者の名前と今でも結びついていることを考えると、彼女が計画したかどうかは別として、企業が主張するのは当然のことだ。
記事内のリンクを通じて購入すると、少額の手数料が発生する場合があります。これは編集上の独立性に影響しません。
コニー・ロイゾス
編集長兼ゼネラルマネージャー
11月4日
ボストン
TechCrunch Founder Summit で最大 190 ドル節約できる最後のチャンス。あらゆる段階の 1,000 人以上の創業者や VC に参加して、現実世界のスケーリングに関する洞察と、時代を変えるようなつながりを手に入れましょう。
割引は 6 月 26 日午後 11 時 59 分に終了します。 PT 。
テキサス州で死亡事故を起こしたテスラ運転手はアクセルを100％踏み続けた、NTSBが認める
Anthropic の最新広告が人々を不気味にさせている
サティア・ナデラ氏はAIを利用する企業に衝撃的な警告を発した
OpenAIに対するAppleの営業秘密訴訟における最も突飛な申し立て
Anthropic が米国に次ぐ最大の市場であるインド向けにクロード価格のローカライズを開始
Meta、反発を受けてInstagramの物議を醸すAI機能を削除
Apple、営業秘密窃盗容疑でOpenAIを提訴

## Original Extract

It's the company's first public proof point after a year and a half spent building AI infrastructure largely out of public view.

TechCrunch Desktop Logo
TechCrunch Mobile Logo
Latest
Thinking Machines amps up its bet against one-size-fits-all AI with its first open model, Inkling
Thinking Machines Lab, the AI startup founded by former OpenAI CTO Mira Murati, released its first in-house AI model Wednesday morning, called Inkling . And unlike the flagship models from OpenAI, Anthropic, or Google, it’s open-weight, meaning outside developers and companies can download it and modify it directly.
Inkling is a mixture-of-experts system with 975 billion total parameters, though it only draws on a fraction of that — about 41 billion — for any given task, a common design that keeps very large models faster and cheaper to run. It was trained on 45 trillion tokens of text, image, audio, and video, and reasons natively across all four, according to the company’s own release materials. For now, though, its outputs are limited to text, including code, styled artifacts, and structured data.
The model is Thinking Machines Labs’ first public proof point after a year and a half spent building AI infrastructure largely out of public view. Some of that work had already surfaced in a May research preview of “interaction models” — AI designed to listen and speak (and even interrupt) instead of stop and wait as with typical chatbots. It’s also a test of the central bet behind the startup, which is that AI that organizations can adapt for themselves will outperform the one-size-fits-all models the biggest labs currently sell.
Inkling is designed to give calibrated answers, including flagging uncertainty rather than guessing, and lets users dial “thinking effort” up or down when they want to trade for speed. On one benchmark, the company says, Inkling uses a third as many tokens as Nvidia’s Nemotron 3 Ultra — its latest generation open-weight model — to hit the same coding performance.
Thinking Machines doesn’t claim Inkling is best-in-class. Its newest blog post states explicitly that Inkling is “not the strongest overall model available today, open or closed.” What it’s evidently going for instead is well-rounded performance.
That raises the question of who, within the enterprise market it’s targeting, this product is really for. Thinking Machines is, for now, marketing Inkling less as a finished product than as a starting point, something for organizations to fine-tune themselves through Tinker, the company’s model-customization platform. This also means customers, not Thinking Machines, are responsible for making sure their customizations are safe, for example. (Fine-tuning requires serious machine-earning talent.)
OpenAI, Anthropic, and Google have all taken a very different approach with ChatGPT, Claude, and Gemini, respectively, which were all built to compete as general-purpose chatbots first, with agentic, autonomous features layered on top.
A post published by Thinking Machines last week was clearly meant as the backdrop for this release. AI that’s trained centrally by one company and then set in stone, the company argued in that post, underperforms AI that organizations shape themselves because so much expertise is specific to the people who hold it.
Other arguments against closed models are gaining steam. In a blog post published Sunday, Microsoft CEO Satya Nadella — whose company has invested billions in both OpenAI and Anthropic — warned that enterprises using proprietary AI models effectively pay twice : once in subscription costs, and again by handing over business knowledge embedded in their prompts and corrections, which can be absorbed into future model versions.
Hugging Face CEO Clem Delangue made a similar prediction in conversation with TechCrunch last week. Frontier models, he said, will increasingly be reserved for experimentation and high-value tasks, while most production AI work shifts to private or open-source alternatives — the exact split Thinking Machines is building around.
The clearest argument for Thinking Machines’ approach came from a recent project with Bridgewater Associates, the world’s largest hedge fund (which is not, for what it’s worth, a Thinking Machines investor). Researchers from both companies took an existing open-source model and trained it further on Bridgewater’s own financial expertise. The result was said to score 84.7% on financial reasoning tests, beating top proprietary AI models, while costing roughly a fourteenth as much to run — though those results come from the two companies’ own evaluation, not an independent one.
Either way, Thinking Machines is emphasizing how quickly it got here. OpenAI took roughly five years to bring its tech to market and show revenue, and Anthropic roughly three. Thinking Machines says it did the same in about nine months.
Some will wonder whether Inkling was trained on outputs from competitors’ models, a practice known as “ distillation ” that has drawn scrutiny across the industry. The short answer, per the company’s own materials, is partly. Thinking Machines pre-trained Inkling from scratch, but it says it used other open-weight models — including Moonshot AI’s Kimi K2.5 — to help generate some of its early post-training data before large-scale reinforcement learning took over. The next model, the company insists, will use fully self-contained post-training instead.
On the cost side, Thinking Machines has been more guarded. It struck a partnership with Nvidia in March to deploy a gigawatt of Vera Rubin computing capacity and trained Inkling entirely on Nvidia’s GB300 NVL72 systems — but hasn’t said how it plans to cover those costs, and revenue, by most accounts, hasn’t been a priority. (A reported $50 billion fundraising round was said to be coming together last November but had stalled by January; the company has declined to talk about its funding picture since.)
A related question is whether Thinking Machines’ spending will ever reach the scale of OpenAI’s or Anthropic’s, or whether its efficiency-driven approach means the economics look different. Put another way, the company’s bet may be less that it will eventually spend like its larger rivals than that it won’t need to at all — because once weights are public, nothing obligates anyone who downloads them to pay Thinking Machines to run them, unlike the metered access OpenAI and Anthropic sell. It’s Tinker, not the model itself, where the company’s revenue has to come from, via training, fine-tuning, and, now, a cut of the hosting ecosystem built around it.
Headcount, at least, looks more settled. Thinking Machines now employs roughly 200 people, up from levels reported after a wave of departures earlier this year, including two co-founders who left for OpenAI in January.
Thinking Machines, for its part, doesn’t seem interested in playing up individual moves the way much of the industry does. According to a source inside the company, its culture, by design, favors continuity over reliance on any one personality. It makes sense: it’s less of a setback when people change teams if they were never put on a pedestal to begin with. It’s also a remarkable thing for a company to insist on, given how much of its own story is still associated with the name of its now-famous co-founder, whether she planned it or not.
When you purchase through links in our articles, we may earn a small commission . This doesn’t affect our editorial independence.
Connie Loizos
Editor in Chief & General Manager
November 4
Boston
Last chance to save up to $190 on TechCrunch Founder Summit. Join 1,000+ founders and VCs at all stages for real-world scaling insights and connections that move the needle.
Savings end June 26, 11:59 p.m. PT .
Tesla driver in fatal Texas crash pressed accelerator 100%, NTSB confirms
Anthropic’s newest ad is creeping people out
Satya Nadella has issued a shocking warning to companies using AI
The wildest allegations in Apple’s trade secrets lawsuit against OpenAI
Anthropic starts localizing Claude pricing for India, its biggest market after the US
Meta removes controversial AI feature on Instagram after backlash
Apple sues OpenAI over alleged trade secret theft
