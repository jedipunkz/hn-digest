---
source: "https://spectrum.ieee.org/ai-cpu-comeback"
hn_url: "https://news.ycombinator.com/item?id=49369311"
title: "Agentic AI has made CPUs the new performance bottleneck"
article_title: "Agentic AI Crunch Creates CPU Comeback - IEEE Spectrum"
image: "https://spectrum.ieee.org/media-library/the-cpu-comeback-is-upon-us.jpg?id=67615384&width=1200&height=600&coordinates=0%2C718%2C0%2C532"
author: "pseudolus"
captured_at: "2026-08-20T02:11:39Z"
capture_tool: "hn-digest"
hn_id: 49369311
score: 1
comments: 0
posted_at: "2026-08-20T01:14:05Z"
tags:
  - hacker-news
  - translated
---

# Agentic AI has made CPUs the new performance bottleneck

- HN: [49369311](https://news.ycombinator.com/item?id=49369311)
- Source: [spectrum.ieee.org](https://spectrum.ieee.org/ai-cpu-comeback)
- Score: 1
- Comments: 0
- Posted: 2026-08-20T01:14:05Z

## Translation

タイトル: Agentic AI により CPU が新たなパフォーマンスのボトルネックになっている
記事のタイトル: Agentic AI Crunch が CPU の復活をもたらす - IEEE Spectrum
説明: Agentic AI は CPU の劇的な復活を推進し、企業が壊滅的なクラウドのボトルネックを回避しようと競う中、レイテンシーを削減し、需要を喚起します。

記事本文:
-->
Raven.config('https://6b64f5cc8af542cbb920e0238864390a@sentry.io/147999').install();
Agentic AI クランチにより CPU が復活 - IEEE Spectrum
IEEE.org IEEE Explore IEEE Standards IEEE Job Site その他のサイト サインイン IEEE に参加 CPU の復活が近づいています テクノロジー インサイダーのためにシェア 検索: トピック別に探索 航空宇宙 AI 生物医学 気候 テクノロジー コンピューティング 家庭用電化製品 エネルギー 技術の歴史 ロボット工学 半導体 電気通信 輸送
IEEEスペクトル
テクノロジー関係者向けトピック
アカウントを作成すると、さらに無料のコンテンツや特典をお楽しみいただけます
後で読むために記事を保存するには、IEEE Spectrum アカウントが必要です
インスティテュートのコンテンツはメンバーのみが利用できます
全号 PDF のダウンロードは IEEE メンバー限定です
この電子ブックのダウンロードは IEEE メンバー限定です
へのアクセス
スペクトル
のデジタル版は IEEE メンバー限定です
以下のトピックは IEEE メンバー限定の機能です
記事に回答を追加するには、IEEE Spectrum アカウントが必要です
アカウントを作成すると、さらに多くのコンテンツや機能にアクセスできます
IEEEスペクトル
、後で読むために記事を保存したり、スペクトル コレクションをダウンロードしたり、イベントに参加したりする機能が含まれます。
読者や編集者との会話。より独占的なコンテンツと機能については、次のことを検討してください。
IEEEに参加する
。
エンジニアリングと応用科学を専門とする世界最大の専門組織に参加し、次の情報にアクセスしましょう。
Spectrum のすべての記事、アーカイブ、PDF ダウンロード、その他の特典を利用できます。
IEEE について詳しくはこちら →
エンジニアリングと応用科学を専門とする世界最大の専門組織に参加し、次の情報にアクセスしましょう。
この電子書籍に加えて、
IEEE スペクトラム
記事、アーカイブ、PDF ダウンロード、その他の特典。
IEEE について詳しくはこちら →
数千の記事にアクセス — 完全に無料
作成

アカウントを取得して限定のコンテンツと機能を入手してください:
記事の保存、コレクションのダウンロード、
そして
コメントを投稿する
— すべて無料です！フルアクセスと特典については、
購読する
スペクトルに。
Agentic AI により CPU が新たなパフォーマンスのボトルネックになっている
Matthew S. Smith は、IEEE Spectrum の寄稿編集者であり、Digital Trends の元主任レビュー編集者です。
ゲッティイメージズ
今年の初め、アマゾン ウェブ サービスのリーダーたちはエンジニアに、どんな犠牲を払ってでも CPU サイクルを節約する必要があるという新たな使命を与えました。 AWS では、AI ワークロードが同社のクラウド インフラストラクチャに負担をかけるため、CPU サーバー容量の待機時間が爆発的に増加したと報告されています。
この問題は AWS を不意を突いたようですが、それには十分な理由がありました。 AI ブームにより、GPU の需要が急増し、その後メモリの需要も高まりました。 CPU は、並列化が相対的に不足しているため、大規模言語モデル (LLM) を実行してユーザーに提供するプロセスである AI モデル推論には適していないため、ほとんどストーリーから除外されました。
しかし、AI モデルが自律的に動作し、サブエージェントを呼び出すことを可能にするエージェント AI システムの台頭により、物語が変わりつつあります。
テキサス州オースティンにある Moor Insights & Strategy の副社長兼主任データセンター アナリストのマット キンボール氏は、2026 年に CPU 需要が急増し、その多くはエージェント AI によるものであると述べています。 「このエージェントのワークロードがあることは別問題であり、それが 100 人のエージェントを生成するとします。これを企業全体に展開しようとすると、その 100 人は数万、数十万、あるいは数百万のエージェントになります」と Kimball 氏は言います。 「エージェントがサブエージェントを生成し、API [アプリケーション プログラミング インターフェイス] 呼び出しを行い、[Anthropic の] モデル コンテキスト プロトコルを通じてさらに多くのエージェントと通信します。」
AI エージェントはコンピューターを使用する必要があり、コンピューターには CPU が必要です
キンボールのコメントは部分的に「

これは、インターネットにアクセスし、デスクトップ上のファイルを開き、通常はタスクを実行するためにさまざまなソフトウェアを使用する LLM の機能の短縮形です。
ツールの使用について訓練を受けた LLM は、他のソフトウェアを呼び出す方法を学びます。 LLM の推論は依然として主に GPU または同様の AI アクセラレータで実行されますが、LLM が行うツール呼び出しは通常 CPU にプッシュされます。
「エージェント AI タスクの多くのコンポーネントは、本質的に CPU ベースのジョブです」と Intel のシニア スタッフ 研究科学者である Souvik Kundu 氏は説明します。 「CPU は出力を解析し、どのツールを呼び出すかを判断し、API 呼び出しまたはコードを実行し、結果を収集してフィードバックするという仕事をします。」 AMD のコンピューティングおよびエンタープライズ AI 製品担当副社長である Madhu Rangarajan 氏も同様の主張をし、「私たちのテストでは、現実的なエージェント AI パイプラインの 8 つのステージのうち 7 つが完全に CPU 上で実行されました。」と述べています。
たとえば、ソフトウェアのプログラミングを任された LLM は、ファイルへのコードの書き込み、ファイルの移動または置換、必要なパッケージのダウンロード、ソフトウェアが完了したと LLM が判断した時点でソフトウェアをビルドするためのツール呼び出しを行う可能性があります。
Kundu 氏は、アトランタのジョージア工科大学の研究者らとエージェント的 AI 最適化に関する論文を共同執筆しました。彼らは、LLM 推論が GPU で実行されている間、CPU がアイドル状態になることが多く、逆に、ツール呼び出しが CPU で実行されるときに GPU がアイドル状態になることが多いことを発見しました。これを最適化するために、Kundu 氏らは、持続的な負荷の下でエンドツーエンドのレイテンシ (エージェント ワークロードの開始から終了までの時間) を最大 1.8 倍削減できるスケジューリングの最適化を提案しています。
それは始まりですが、ゲインは動くターゲットを追いかけます。エージェント システムはマシンの速度で作業を生成し、作業が進むにつれて作業量を増やします。 OpenAI による Hugging Face の不注意によるハッキングにより、モデルが正しく認識されなくなりました

1 時間あたり最大 300 のアクションが実行され、単一のエージェントが独自のツール呼び出しを行うサブエージェントを生成できます。
そして、モデルがより複雑になるにつれて、CPU の作業負荷が増加する可能性があるもう 1 つの重要な問題があります。それは、安全ガードレールです。
Kundu 氏によると、エージェントのアクションに対する安全性とポリシーのチェックは、構文とログ ファイルを検査する特定のルールであることがよくあります。ガードレールは、タスクの複雑さや意図を分析するために、小規模なモデル (10 億パラメータ以下) を使用することもあります。これらは GPU 上で実行することもできますが、サイズが小さく、レイテンシーを最小限に抑える必要があるため、CPU 上で作業が継続されるため、実際には実行できないことがよくあります。
使用可能な CPU の数を増やすと、シーケンス長が長くなった場合の Llama-8B 応答の遅延が大幅に減少します。出典: Euijun Chung、Yuxiao Jia 他
トークン化がボトルネックを増大させる
ジョージア工科大学の博士課程学生である Euijun Chung 氏は、最近、Kundu 氏の研究を補足する調査結果を含む別の論文を共著しました。 Chung 氏と彼の共著者らは、サーバーの CPU コアが少なすぎると、GPU への作業のディスパッチが遅れることを発見しました。これにより、GPU は命令を待つ間に停止してしまいます。
それに加えて、この論文では LLM ワークロードの別の重要な要素であるトークン化についても触れています。
トークン化は、LLM 推論の重要な最初のステップです。テキストを、モデルで処理できる整数のトークン ID に変換します。ほとんどの LLM 推論に必要な行列計算とは異なり、トークン化は分岐があり、データに依存する逐次文字列操作です。テキストをチャンク化することで並列化できますが、LLM 推論の大部分と同じように大規模に並列化することはできません。
小さなプロンプトのトークン化は比較的簡単なタスクであり、エントリーレベルの CPU にも負担はかかりません。ただし、ツール呼び出しを行うエージェント モデルは、その結果を解析してトークン化する必要があります。

電話します。
「たとえば、100,000 トークンの進行中のシーケンスがあり、ツールの結果が 1,000 トークンである場合、トークナイザーはシーケンス全体を再度トークン化する必要があります。また、エージェント ツール呼び出しのたびにトークン化を行う必要があります」と Chung 氏は言います。これにより、トークン化の頻度が増加し、関与するトークンの数も増加します。将来のトークナイザーがこれを軽減する方法を見つけることは考えられますが、現代の LLM 推論にとっては依然として問題です、と Chung 氏は言います。
この論文では、シーケンスの長さが長くなると、最初のトークンまでの待ち時間 (モデルが応答の最初の単語を生成するのに必要な時間) が劇的に増加する可能性があることを発見しました。より多くのコアを備えた CPU を使用すると、問題が軽減される可能性があります。より長いシーケンス長でのテスト実行では、CPU コア数を増やすことで、最初のトークンまでの時間を約 1.5 ～ 7 倍短縮できます。
Chung 氏と彼の同僚は、テストに利用できるハードウェアの制限により、Alibaba の Qwen 3-30B や Meta の Llama 3.1-70B などの小型モデルしかテストできませんでした。彼は、大規模なモデルでは全体的な GPU 需要が高いため、劇的なボトルネックはそれほど発生しないだろうと推測していますが、エージェント AI によってトークンの長さが、彼と彼の共著者がテストした最長をはるかに超えるものになるとも予想しています。
「Anthropic のクロードのようなものを考えれば、簡単に 500,000、さらには 100 万トークンに達する可能性があります」と Chung 氏は言います。 「エージェント AI の世界では、平均シーケンス長はますます増大するため、この問題は将来のワークロードでさらに悪化すると予想されます。」
CPU 不足はまだ始まったばかりですか?
CPU リソースの使用に対する Amazon の取り締まりは、Kundu 氏と Chung 氏が現実世界に関連する問題を特定したことを示すいくつかの指標のうちの 1 つです。
Intelは少なくとも年末までサーバーCPUを完売した。 AMDはサーバーのCPU fを2倍にしました

オレキャスト。 Arm と Qualcomm は両方とも、エージェント AI を高速化するために設計された新しい CPU を発表しました。 Nvidia でさえ、Nvidia の Vera Rubin プラットフォームの一部であるエージェント AI 用の Arm ベースの CPU Vera を優先しています。
Kimball 氏は、これらの開発により、AI 業界が CPU パフォーマンスをより重視していることが明らかになったと述べています。同氏は、需要の急増は、CPU が現在エージェント AI システムの重要な部分と考えられているという「絶対的な証拠」であると見ています。
残念ながら、これは、GPU やメモリですでに発生しているのと同じように、広範な CPU 不足と価格の上昇につながる可能性があります。
「すでにある程度の CPU 不足が発生しています。市場の制約に目を向けると、その制約は消費者分野にも波及しています」とキンボール氏は言います。同氏は、インテルの新しい 18A 生産プロセスがクライアント部門での同社の売上を伸ばしたにもかかわらず、インテルはサーバー CPU を優先してクライアント CPU の生産を削減したと付け加えた。 Kimball 氏は、これを CPU メーカーが資金に追随する兆候であると見ています。
インテルの将来のファウンドリ技術を覗いてみる ›
スタートアップは 100 倍高速な CPU を実現できると発表 ›
CPU戦争が帰ってきた！ - ユサフ・バブールとサクイブ・タヒル著 ›
AI の変化に伴いインテルは「CPU カムバック」にどのように乗っているのか - そして中国の立ち位置 |サウスチャイナ・モーニング・ポスト ›
Matthew S. Smith は、17 年の経験を持つフリーランスの消費者技術ジャーナリストであり、Digital Trends の元主任レビュー編集者です。 IEEE スペクトラムの寄稿編集者である彼は、ディスプレイの革新、人工知能、拡張現実に焦点を当てた消費者向けテクノロジーをカバーしています。ビンテージ コンピューティングの愛好家である Matthew は、YouTube チャンネル Computer Gaming Yesterday でレトロ コンピューターとコンピューター ゲームを取り上げています。
シミュレーション アプリを使用して電子機器の故障の根本原因を特定する
AI を使用してこれまでで最も困難な数学の証明を検証
スタート

up は 100 倍高速な CPU を実現できると発表
AI 副操縦士からエージェント群まで
詐欺師たちはAIの幻覚を愛することを学んでいる

## Original Extract

Agentic AI drives a dramatic cpu comeback, slashing latency and igniting demand as enterprises race to avoid crippling cloud bottlenecks.

-->
Raven.config('https://6b64f5cc8af542cbb920e0238864390a@sentry.io/147999').install();
Agentic AI Crunch Creates CPU Comeback - IEEE Spectrum
IEEE.org IEEE Xplore IEEE Standards IEEE Job Site More Sites Sign In Join IEEE The CPU Comeback Is Upon Us Share FOR THE TECHNOLOGY INSIDER Search: Explore by topic Aerospace AI Biomedical Climate Tech Computing Consumer Electronics Energy History of Technology Robotics Semiconductors Telecommunications Transportation
IEEE Spectrum
FOR THE TECHNOLOGY INSIDER Topics
Enjoy more free content and benefits by creating an account
Saving articles to read later requires an IEEE Spectrum account
The Institute content is only available for members
Downloading full PDF issues is exclusive for IEEE Members
Downloading this e-book is exclusive for IEEE Members
Access to
Spectrum
's Digital Edition is exclusive for IEEE Members
Following topics is a feature exclusive for IEEE Members
Adding your response to an article requires an IEEE Spectrum account
Create an account to access more content and features on
IEEE Spectrum
, including the ability to save articles to read later, download Spectrum Collections, and participate in
conversations with readers and editors. For more exclusive content and features, consider
Joining IEEE
.
Join the world’s largest professional organization devoted to engineering and applied sciences and get access to
all of Spectrum’s articles, archives, PDF downloads, and other benefits.
Learn more about IEEE →
Join the world’s largest professional organization devoted to engineering and applied sciences and get access to
this e-book plus all of
IEEE Spectrum’s
articles, archives, PDF downloads, and other benefits.
Learn more about IEEE →
Access Thousands of Articles — Completely Free
Create an account and get exclusive content and features:
Save articles, download collections,
and
post comments
— all free! For full access and benefits,
subscribe
to Spectrum .
Agentic AI has made CPUs the new performance bottleneck
Matthew S. Smith is a contributing editor for IEEE Spectrum and the former lead reviews editor at Digital Trends.
Getty Images
Earlier this year, leaders at Amazon Web Services delivered a new mandate to their engineers: They need to conserve CPU cycles at all costs. AWS has reportedly experienced an explosion in wait times for CPU server capacity as AI workloads strain the company’s cloud infrastructure.
The issue seemingly took AWS off guard, and for good reason. The AI boom led to a surge in demand for GPUs and, later, memory . CPUs were mostly left out of the story, as their relative lack of parallelization made them a poor fit for AI model inference, the process of running and serving large language models (LLM) to users.
But the rise of agentic AI systems, which allow AI models to operate autonomously and call on sub-agents, is changing the narrative.
Matt Kimball , vice president and principal data center analyst at Moor Insights & Strategy in Austin, Texas, says 2026 has brought a spike in CPU demand, much of it due to agentic AI . “It’s one thing to have this agentic workload, and let’s say it spawns 100 agents. If I’m going to roll this out across my enterprise, those 100 become tens of thousands, hundreds of thousands, or millions of agents,” Kimball says. “You have agents spawning sub-agents, making API [application programming interface] calls and talking to more agents through [Anthropic’s] model context protocol.”
AI agents need to use computers, and computers need CPUs
Kimball’s comments refer in part to “tool use,” which is shorthand for an LLM’s ability to access the internet, open files on a desktop, and generally use a variety of software to accomplish its task.
LLMs trained for tool use learn how to call on other software. While the LLM’s inference is still primarily executed on a GPU or similar AI accelerator, the tool calls that the LLM makes are typically pushed to the CPU.
“Many components of an agentic AI task are inherently CPU based jobs,” explains Souvik Kundu , senior staff research scientist at Intel . “The CPU does the job of parsing output, figuring out which tool to invoke, making the API call or running the code, collecting the result, and feeding it back.” Madhu Rangarajan , vice president of compute and enterprise AI products at AMD , makes a similar claim, saying, “In our testing, seven of the eight stages in realistic agentic AI pipelines run entirely on the CPU.”
An LLM tasked with programming software, for example, will likely make tool calls to write code to files, move or replace files, download required packages, and build the software once the LLM believes it’s complete.
Kundu co-authored a paper on agentic AI optimization alongside researchers from Georgia Tech in Atlanta. They found the CPU is often idle while LLM inference is executed on a GPU and that, conversely, the GPU is often idle when tool calls are executed on the CPU. To optimize this, Kundu and his colleagues propose scheduling optimizations that can cut end-to-end latency (the time between the start and finish of the agentic workload) by up to 1.8-times under sustained load.
It’s a start, but the gains chase a moving target. Agentic systems generate work at machine speed and multiply it as they go. OpenAI’s inadvertent hack of Hugging Face saw its model fire off as many as 300 actions an hour, and a single agent can spawn sub-agents that make tool calls of their own.
And there’s one more important complication that may increase the workload on a CPU as models become more complex: safety guardrails.
Safety and policy checks on an agent’s actions are often specific rules that inspect syntax and log files, Kundu says. Guardrails may also use small models (under a billion parameters) to analyze task complexity or intent. Though they could be executed on a GPU, they often aren’t, because their small size and the need to minimize latency keeps the work on the CPU.
Increasing the number of CPUs available significantly decreases the latency for Llama-8B responses over longer sequence lengths. Source: Euijun Chung, Yuxiao Jia, et al.
Tokenization adds to bottlenecks
Euijun Chung , a PhD student at Georgia Tech, recently co-authored another paper , with findings that complement Kundu’s work. Chung and his co-authors found that when a server has too few CPU cores, it falls behind on dispatching work to the GPUs . That causes the GPUs to stall as they wait for instructions.
In addition to that, the paper touches on another key element of LLM workloads: tokenization .
Tokenization is a key first step in LLM inference. It converts text into integer token IDs that can be processed by the model. Unlike the matrix math required for most LLM inference, tokenization is branchy, data-dependent sequential string manipulation. Though it can be parallelized by chunking text, it’s not massively parallel in the same way as the bulk of LLM inference.
Tokenization of small prompts is a relatively trivial task and won’t tax even an entry-level CPU. However, an agentic model that makes tool calls must parse and tokenize the results of the call.
“If you have an ongoing sequence of, say, 100,000 tokens, and you have a tool result of 1,000 tokens, the tokenizer will have to tokenize the whole sequence again. And you have to do tokenization at every agentic tool call,” Chung says. This both increases the frequency of tokenization and increases the number of tokens involved. It’s conceivable that future tokenizers will find ways to mitigate this, Chung says, but it remains a problem for modern LLM inference.
The paper finds that time-to-first-token latency (the time required for the model to produce the first word of its reply) can increase dramatically as the sequence length grows. CPUs with more cores can reduce the problem. In test runs at longer sequence lengths, increasing CPU core counts can reduce time-to-first-token latency by roughly 1.5 to 7 times.
Chung and his colleagues were only able to test smaller models, such as Alibaba’s Qwen 3-30B and Meta’s Llama 3.1-70B, due to limitations of the hardware available for testing. He speculates that larger models will experience less dramatic bottlenecks due to their higher overall GPU demand, but also expects agentic AI will push token lengths far beyond the longest he and his co-authors tested.
“If you think about something like Anthropic’s Claude, you can easily hit 500,000, even a million tokens,” Chung says. “In the world of agentic AI, the average sequence length will grow and grow, so I’m expecting this problem to get worse in future workloads.”
Is a CPU crunch just getting started?
Amazon’s crackdown on use of CPU resources is one of several indicators that Kundu and Chung have identified issues with real-world relevance.
Intel has sold out of server CPUs through at least the end of the year. AMD has doubled its server CPU forecast. Arm and Qualcomm have both announced new CPUs designed to accelerate agentic AI. Even Nvidia has prioritized Vera , its Arm-based CPU for agentic AI, which is part of Nvidia’s Vera Rubin platform.
Kimball says these developments make it clear that the AI industry is placing more emphasis on CPU performance. He sees the surge in demand as an “absolute tell” that CPUs are now considered a key part of an agentic AI system.
Unfortunately, this may translate to broader CPU shortages and increased prices, much as has already occurred with GPUs and memory.
“You’re already seeing a CPU crunch to some degree. When you look at the constraints in the market, it even trickles down into the consumer space,” Kimball says. He adds that Intel has cut production of client CPUs in favor of server CPUs, even as Intel’s new 18A production process has grown the company’s sales in the client segment. Kimball sees that as a sign that CPU makers will follow the money.
A Peek at Intel’s Future Foundry Tech ›
Startup Says It Can Make a 100x Faster CPU ›
The CPU Wars are Back! - by Yousaf Babur and Saqib Tahir ›
How Intel is riding the ‘CPU comeback’ as AI shifts - and where China stands | South China Morning Post ›
Matthew S. Smith is a freelance consumer technology journalist with 17 years of experience and the former Lead Reviews Editor at Digital Trends. An IEEE Spectrum Contributing Editor, he covers consumer tech with a focus on display innovations, artificial intelligence, and augmented reality. A vintage computing enthusiast, Matthew covers retro computers and computer games on his YouTube channel, Computer Gaming Yesterday .
Identifying the Root Cause of Electronics Failures With Simulation Apps
AI Used to Verify Toughest Mathematics Proof Yet
Startup Says It Can Make a 100x Faster CPU
From AI Copilots to Agent Swarms
Crooks Are Learning to Love AI Hallucinations
