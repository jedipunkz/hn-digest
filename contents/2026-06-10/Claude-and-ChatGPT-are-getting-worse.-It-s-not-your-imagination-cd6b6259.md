---
source: "https://www.artificialstudio.ai/blog/chatgpt-claude-getting-worse-compute-crisis"
hn_url: "https://news.ycombinator.com/item?id=48470821"
title: "Claude and ChatGPT are getting worse. It's not your imagination"
article_title: "Claude and ChatGPT are getting worse. It's not your imagination. | ArtificialStudio"
author: "artificialstudi"
captured_at: "2026-06-10T04:34:56Z"
capture_tool: "hn-digest"
hn_id: 48470821
score: 2
comments: 0
posted_at: "2026-06-10T02:57:00Z"
tags:
  - hacker-news
  - translated
---

# Claude and ChatGPT are getting worse. It's not your imagination

- HN: [48470821](https://news.ycombinator.com/item?id=48470821)
- Source: [www.artificialstudio.ai](https://www.artificialstudio.ai/blog/chatgpt-claude-getting-worse-compute-crisis)
- Score: 2
- Comments: 0
- Posted: 2026-06-10T02:57:00Z

## Translation

タイトル: クロードとChatGPTが悪化しています。それはあなたの想像ではありません
記事タイトル：ClaudeとChatGPTが悪化しています。それはあなたの想像ではありません。 |人工スタジオ
説明: AI モデルは静かに限界に達しており、企業は何も言わずに容量を割り当てています。ここでは、実際に何が起こっているのか、それが毎日使用するツールに影響を与える理由、そしてそれに対して何ができるのかを説明します。

記事本文:
Artificial Studio / Blog アプリを開く ブログ • インサイト • 2026 年 6 月 1 日 Claude と ChatGPT が悪化しています。それはあなたの想像ではありません。
AI モデルは静かに限界に達しており、企業は何も言わずに容量を割り当てています。ここでは、実際に何が起こっているのか、それが毎日使用するツールに影響を与える理由、そしてそれに対して何ができるのかを説明します。
最近 ChatGPT が遅くなったと感じたり、クロードがタスクの途中で中断したり、以前は数秒かかっていた応答にはるかに時間がかかるようになったことに気付いた場合、それは想像できません。
起こっていることはバグではありません。これは、AI 業界が公然と議論することを躊躇してきた構造的な問題です。これらのモデルを動かすインフラストラクチャの容量が不足しており、モデルを構築している企業は、得られるものの品質に直接影響するトレードオフを密かに行っています。
まず、トークンとは何なのかを理解する必要があります #
AI にメッセージを送信すると、モデルは、ユーザーが書いたとおりに言葉を読み上げません。トークンと呼ばれる単語の断片を読み取ります。
👉 簡単な質問には 500 ～ 1,000 トークンの費用がかかる場合があります。 AI エージェントとの 1 時間の作業セッション (タスクの草案、改訂、検索、実行など) では、50,000 ～ 500,000 トークンが消費される可能性があります。
最後の数字が重要です。エージェント AI (単一の質問に答えるのではなく、一連のアクションを実行する AI) は、単純なプロンプトに比べて 10 ～ 100 倍のコンピューティングを使用します。より複雑なタスクに AI を使用する人が増えるにつれて、インフラストラクチャに対する総需要が爆発的に増加しました。そしてインフラも追いついていない。
標識はすべての主要なプラットフォームに表示されます。 #
現在最も広く使用されている AI モデルの 1 つである Claude は、今年の一部の期間で API 稼働率が 98.95% でした。計算してみるまでは、これで問題ないと思われます。

業界標準は 99.99% で、これは 1 か月あたり約 4 分のダウンタイムに相当します。 98.95% というと、累積でほぼ 24 時間の停止に相当します。
2026 年 3 月は単月としては最悪の月でした。クロードは 13 時間ダウンしていました。
これに応じて、Anthropic はピーク時間帯、具体的には太平洋時間の午前 5 時から午前 11 時までの間のトークン制限を課しました。ヨーロッパのユーザーにとって、これはつまり、勤務時間の真ん中である午後 1 時から午後 7 時までの間に開始を制限することを意味します。打撃を和らげるために、彼らはピーク期間外のトークン制限を倍増する「春休み」キャンペーンを実施した。これは電力会社がオフピーク価格で使用するのと同じ戦略である。
Anthropic の最新モデルである Claude Opus 4.7 は、同じ量のテキストに対して以前のモデルよりも 46% 多くのトークンを消費します。価格は変わりませんでした。変わったのは、割り当てが以前よりも大幅に早くなくなることです。そして4月下旬、Anthropicはひそかに価格ページを更新し、以前は標準の有料プランに含まれていたコーディング機能を月額100ドルか200ドルのプランに移行した。ユーザーがソーシャルメディアでそれをキャッチし、Anthropic は決定を覆しましたが、意図は明らかでした。
OpenAI は、同社のビデオ生成アプリである Sora を、発売からわずか 6 か月後の 2026 年 3 月 24 日に閉鎖しました。このアプリは最初の週で 100 万ダウンロードに達しました。また、計算コストとして 1 日あたり 100 万ドルを消費していましたが、それを相殺できる収益は最小限でした。
OpenAI は、マーベル、スター・ウォーズ、ピクサー IP へのアクセスに関して、ディズニーと 10 億ドル以上の契約を締結しました。その取引は現在保留中です。
Google の Gemini では、2025 年 12 月から 2026 年 3 月までの 4 か月間で 4 回のサイレント削減が行われました。12 月には、事前の通知もなく、無料利用枠が一晩で 1 日のリクエスト容量で 92% 減少しました。 3 月までに、最も高価なサブスクリプション層のユーザーでも、何の制限もなしに制限が引き下げられました。

コミュニケーション。
なぜ企業はそれをもっと早く修正しないのか #
AI 業界はこの問題を認識しており、それに対処するために積極的な資金を投入しています。
大手テクノロジー企業は今年、新しいデータセンターの拡張と構築に合わせて 7,000 億ドル以上を投資しています 💸。目標は、AI インフラストラクチャの電流容量を 3 倍にすることです。 #
しかし、ゴールドマン・サックスは、少なくとも2028年までは需要が年間10ギガワットの電力供給を上回り続けると予測している。お金は存在する。海外からの材料調達、許可、米国内の資格のあるエンジニアの不足といったボトルネックがすべてを遅らせている。 2025年1月に発表された5000億ドルのインフラプロジェクトであるスターゲイトは、計画していたテキサスの施設が2026年4月時点で一時停止された。
これを大局的に比較すると、携帯電話が大量に普及するまでにおよそ 15 年かかりました。 Agentic AI は、約 18 か月でニッチなものからユビキタスなものになりました。インフラ危機も本質的には同じだ。それを修正するためのスケジュールははるかに厳しいです。
これらの企業が下すあらゆる意思決定を形作る財務上の側面もあります。 OpenAIは2026年に140億ドルの損失を被ると予測されている。
Anthropic は 2025 年に推定 30 億ドルの損失を被りました。現在、ChatGPT の 9 億ユーザーのうちサービス料金を支払っているのは 5.5% だけです。つまり、OpenAI は残りの 94.5% のコンピューティング コストを自己資本から実質的に補助していることになります。それは持続可能なモデルではないことは、業界の誰もが知っています。
OpenAI は、需要が高い期間中、高度な機能が少ない機能の低いモデルからは無料ユーザーの応答が遅くなることを確認しました。
コンピューティングが不足していて高価であれば、より多くのお金を支払える企業がより多くの恩恵を受けることになります。 #
AI が優れているということは、より良い意思決定、より良い成果、より良い競争上の地位を意味します。

ああ、さらに。それは自己強化ループです。
AI をクリエイティブな仕事に使用する場合、これは何を意味しますか #
品質はエンドユーザーにはまったく見えない形で変動します。 「今日の応答は 70% の容量で実行されています」という通知は届きません。得られる答えが短かったり、タスクが途中で中断されたり、先週得たものよりも明らかに精度が低い結果が得られるだけです。
AI を使用して画像の生成、ビデオの編集、オーディオの制作、コピーの作成、コンセプトの開発を行うクリエイティブなプロフェッショナルにとって、この矛盾は実際の制作に影響を及ぼします。予期せぬ動作をするツールは、ワークフローを構築できないツールです。
すべてを単一のモデルに依存するのはやめましょう。 #
1 つのモデルが容量の負担にさらされていても、別のモデルはそうではない可能性があります。異なるモデルで同じ概要をテストする (そして、特定の日にどのモデルが良好なパフォーマンスを示しているかを知る) ことは、贅沢ではなく、ますます専門的なスキルになりつつあります。
Artificial Studio のようなプラットフォームは、まさにこれを目的として構築されています。異なるモデル間で同じプロンプトを実行し、数秒で結果を比較できます。 1 つのモデルのパフォーマンスが低い場合は、切り替えることになります。出力の品質は、ある会社のサーバーの調子が悪くなっているかどうかに依存する必要はありません。
これは過渡的な瞬間であり、永続的な状態ではありません。インフラストラクチャは最終的には追いつきますが、それは常に追いつきます。しかし、今からそれまでの期間は、制限、静かな劣化、そして最も多くのお金を払った人に最高のエクスペリエンスを与える価格モデルによって定義されるでしょう。
その背後にある仕組みを理解すると、トークンのコスト、なぜ容量が割り当てられるのか、これらの企業が財務的にどのように構成されているかなど、ツールが壁にぶつかったときに不意を突かれることがなくなります。あなたは何が起こっているかを知っており、それを回避する方法を知っています。
この記事で紹介されているツールを使用して作成を始めてください。
なぜ 80% が co

企業は AI による ROI を期待していません
AI の導入は過去最高を記録していますが、ほとんどの企業は目に見える利益を期待していません。ここでは、実際に AI から価値を得る 5% がどのような違いを見せているのか、そしてクリエイティブ チームがそれをどのように真似できるのかを紹介します。
デザイナーが単なるデザイナーでなくなった年 - AI in Design Report 2026
Designer Fund と Foundation Capital による AI in Design 2026 レポートは、変革の最中にある職業を明らかにしています。職務記述書は、もはや以前とは意味が異なります。何が変わったのかは次のとおりです。
企業はクリエイティブを削減しているわけではありません。彼らは経営者を排除している。
Cisco、Coinbase、GitLab は記録的な収益を上げながら、何千人ものマネージャーを解雇しています。この構造的変化が、クリエイティブな専門家、フリーランサー、そして生計を立てるために何かを作る人にとって実際に何を意味するのかを以下に示します。
AI で背景ビデオを削除
テキストからビデオへの AI ジェネレーター オンライン
AI で写真を 3D モデルに変換
人工スタジオ vs. ミッドジャーニー
Artificial Studio 対 Adobe Firefly
Artificial Studio と RunwayML の比較
人工スタジオ vs. ヒッグスフィールド AI

## Original Extract

AI models are quietly hitting their limits and the companies are rationing capacity without telling you. Here's what's actually happening, why it affects the tools you use every day, and what you can do about it.

Artificial Studio / Blog Open app Blog • Insights • June 1, 2026 Claude and ChatGPT are getting worse. It's not your imagination.
AI models are quietly hitting their limits and the companies are rationing capacity without telling you. Here's what's actually happening, why it affects the tools you use every day, and what you can do about it.
If you've noticed that ChatGPT feels slower lately, or that Claude cuts off mid-task, or that a response that used to take seconds now takes much longer, you're not imagining it.
What's happening isn't a bug. It's a structural problem that the AI industry has been reluctant to discuss openly: the infrastructure that powers these models is running out of capacity, and the companies building them are quietly making trade-offs that directly affect the quality of what you get.
First, you need to understand what tokens are #
When you send a message to an AI, the model doesn't read your words the way you wrote them. It reads fragments of words called tokens.
👉 A simple question might cost between 500 and 1,000 tokens. A one-hour working session with an AI agent (the kind where it's drafting, revising, searching, and executing tasks) can consume between 50,000 and 500,000 tokens!
That last number is the one that matters. Agentic AI (AI that takes sequences of actions rather than answering a single question) uses between 10 and 100 times more compute than a simple prompt. As more people use AI for more complex tasks, the total demand on the infrastructure has exploded. And the infrastructure hasn't kept up.
The signs are visible across every major platform. #
Claude, currently one of the most widely used AI models, had an API uptime of 98.95% during parts of this year. That sounds close to fine until you do the math: the industry standard is 99.99%, which equals about four minutes of downtime per month. At 98.95%, that's nearly 24 accumulated hours of outages.
March 2026 was the worst single month: Claude was down for 13 hours.
In response, Anthropic imposed token limits during peak hours — specifically between 5 and 11am Pacific time. For users in Europe, that translated to limits kicking in between 1pm and 7pm: the middle of the working day. To soften the blow, they ran a "spring break" campaign doubling token limits outside those peak windows — the same strategy electric companies use with off-peak pricing.
Anthropic's latest model, Claude Opus 4.7, consumes 46% more tokens than its predecessor for the same amount of text. The price didn't change. What changed is that your quota runs out significantly faster than it used to. And in late April, Anthropic quietly updated its pricing page to move coding capabilities — previously included in standard paid plans — to plans costing $100 or $200 per month. Users caught it on social media, Anthropic reversed the decision, but the intent was clear.
OpenAI shut down Sora — its video generation app — on March 24, 2026, just six months after launch. The app had reached one million downloads in its first week. It was also burning one million dollars per day in compute costs, with minimal revenue to offset it.
OpenAI had a signed deal with Disney worth over a billion dollars for access to Marvel, Star Wars, and Pixar IP. That deal is now on hold.
Google's Gemini has seen four silent cuts in four months between December 2025 and March 2026. In December, the free tier dropped 92% in daily request capacity overnight, with no prior notice. By March, even users on the most expensive subscription tier had their limits reduced without any communication.
Why the companies aren't fixing it faster #
The AI industry is aware of the problem and is spending aggressively to address it.
Major tech companies are collectively investing over $700 billion this year in expanding and building new data centers 💸. The goal is to triple current electrical capacity for AI infrastructure. #
But Goldman Sachs projects that demand will continue to outpace supply by 10 gigawatts of electricity per year through at least 2028. The money exists. The bottlenecks — materials sourced from abroad, permits, and a shortage of qualified engineers in the US — are slowing everything down. Stargate, the $500 billion infrastructure project announced in January 2025, had its planned Texas facility paused as of April 2026.
The comparison that puts this in perspective: mobile phones took roughly 15 years to reach mass adoption. Agentic AI went from niche to ubiquitous in about 18 months. The infrastructure crisis is the same in nature. The timeline to fix it is much tighter.
There's also a financial dimension that shapes every decision these companies make. OpenAI is projected to lose $14 billion in 2026.
Anthropic lost an estimated $3 billion in 2025. Only 5.5% of ChatGPT's 900 million users currently pay for the service — meaning OpenAI is effectively subsidizing the compute costs of the remaining 94.5% out of its own capital. That is not a sustainable model, and everyone in the industry knows it.
OpenAI has confirmed that during periods of high demand, free users receive slower responses from less capable models with fewer advanced features.
If compute is scarce and expensive, the companies that can pay more will get more of it. #
Better AI means better decisions, better output, better competitive position — which means they can pay even more. It's a self-reinforcing loop.
What this means if you use AI for creative work #
Quality fluctuates in ways that are completely invisible to the end user. You don't get a notification saying "today's responses are running at 70% capacity." You just get a shorter answer, a task that cuts off, or a result that's noticeably less precise than what you got last week.
For creative professionals, people using AI to generate images, edit video, produce audio, write copy, develop concepts, this inconsistency has real production consequences. A tool that performs unpredictably is a tool you can't build a workflow around.
Stop depending on a single model for everything. #
When one model is under capacity strain, another might not be. Testing the same brief across different models (and knowing which one is performing well on a given day) is increasingly a professional skill, not a luxury.
Platforms like Artificial Studio are built for exactly this: you can run the same prompt across different models and compare results in seconds. If one model is underperforming, you switch. The quality of your output doesn't have to depend on whether one company's servers are having a bad week.
This is a transitional moment, not a permanent state. The infrastructure will eventually catch up, it always does. But the period between now and then will be defined by limits, quiet degradations, and pricing models that reward whoever pays most with the best experience.
Understanding the mechanics behind it means you're not caught off guard when your tool hits a wall : what tokens cost, why capacity gets rationed, how these companies are structured financially. You know what's happening, and you know how to work around it.
Start creating with the tools mentioned in this article.
Why 80% of companies see no ROI from AI
AI adoption is at record highs but most companies see no measurable return. Here's what the 5% that actually gets value from AI is doing differently, and how creative teams can copy it.
The year designers stopped being just designers - AI in Design Report 2026
The AI in Design 2026 report by Designer Fund and Foundation Capital reveals a profession mid-transformation. A job description that no longer means what it used to. Here's what changed.
Companies aren't cutting creatives. They're eliminating the managers.
Cisco, Coinbase and GitLab are laying off thousands of managers — while posting record revenue. Here's what this structural shift actually means for creative professionals, freelancers, and anyone who makes things for a living.
Remove Background Videos with AI
Text to Video AI Generator Online
Convert Pictures into 3D Models with AI
Artificial Studio vs. Midjourney
Artificial Studio vs. Adobe Firefly
Comparing Artificial Studio to RunwayML
Artificial Studio vs. Higgsfield AI
