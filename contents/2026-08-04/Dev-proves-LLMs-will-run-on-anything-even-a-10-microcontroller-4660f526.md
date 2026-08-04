---
source: "https://www.theregister.com/edge-and-iot/2026/08/04/dev-proves-llms-will-run-on-anything-even-a-10-microcontroller/5283088"
hn_url: "https://news.ycombinator.com/item?id=49175530"
title: "Dev proves LLMs will run on anything, even a $10 microcontroller"
article_title: "Dev proves LLMs will run on anything – even a $10 microcontroller"
author: "Bender"
captured_at: "2026-08-04T22:04:21Z"
capture_tool: "hn-digest"
hn_id: 49175530
score: 2
comments: 0
posted_at: "2026-08-04T21:37:39Z"
tags:
  - hacker-news
  - translated
---

# Dev proves LLMs will run on anything, even a $10 microcontroller

- HN: [49175530](https://news.ycombinator.com/item?id=49175530)
- Source: [www.theregister.com](https://www.theregister.com/edge-and-iot/2026/08/04/dev-proves-llms-will-run-on-anything-even-a-10-microcontroller/5283088)
- Score: 2
- Comments: 0
- Posted: 2026-08-04T21:37:39Z

## Translation

タイトル: 開発者は、LLM が 10 ドルのマイクロコントローラーでも何でも実行できることを証明しました
記事のタイトル: LLM があらゆるもの (10 ドルのマイクロコントローラーでも) で動作することを開発者が証明
説明: ほぼ 10 トーク/秒

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
2026 年クラウド インフラストラクチャ月間
FIS と AWS による金融サービスの最新化
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
開発者は、LLM があらゆるもの (10 ドルのマイクロコントローラーであっても) で動作することを証明しました
トバイアス・マン
トビアス
マン
システムエディター
発行済み
2026年8月4日火曜日 // 22:20 UTC
2026 年には、小規模なローカル言語モデルをノートブックやスマートフォン上で実行することは簡単です。しかし、さらに小型で低消費電力のものはどうでしょうか。たとえば、10 ドル未満の ESP32 マイクロコントローラーのようなものでしょうか?
不可能に聞こえるかもしれません。このデバイスは主にリモート センサー、IoT、その他の組み込みアプリケーション向けに設計されており、生成 AI モデルを実行するためのものではありません。それでも、ハンドル名 SlvDev を名乗る開発者はまさにそれを実現しました。
GitHub で詳細が説明され、最近 Better Stack YouTube チャンネルで紹介されたプロセスの中で、SlvDev は、高級コーヒー 1 杯分とほぼ同じコストのマイクロコントローラー上で、毎秒 10 トークン近くでローカルに実行される小規模な言語モデルをどのようにして取得できたかを文書化しました。
小さなマイクロコントローラーの小さな物語
ESP32 マイクロコントローラーのような小さなものに大規模言語モデル (LLM) を詰め込むのは簡単な作業ではありません。
これらのモデルが GPU でトレーニングされ、実行されるのには理由があります。 LLM はメモリを大量に消費する猛獣であり、通常、重みをメモリに保持するためだけにパラメータごとに 1 ～ 4 バイトを必要とします。
ESP32-S3 上の SRAM が 520 KB と擬似 SRAM (PSRAM) が 8 MB だけでは、DeepSeek V4 Flash のようなモデルを実行することはできません。
それを機能させるために、開発者は言語モデルから「大きい」を削除し、ほぼ 10,000 倍小さいもの、つまり元々 2,890 万のパラメータを持つモデル TinyStories に落ち着く必要がありました。

Microsoft Research によって開発されました。
ただし、このモデルでも ESP32-S3 モジュールが多く求められます。 16 ビット精度では、このモデルには約 60 MB のメモリが必要ですが、ESP32 にはそれがありません。そこで開発者は、モデルのフットプリントを縮小するために、以前に検討したいくつかのテクニックを採用しました。
1 つ目は量子化です。これは、重みの精度を 16 ビットのような精度から 8 ビット、さらには 4 ビットに減らすことで重みを圧縮するプロセスです。
これにより、SlvDev は精度を多少犠牲にして、必要なメモリを 75% 削減することができました。重みを保持するために必要なメモリは約 60 MB ではなく、必要なメモリはわずか 14.9 MB です。
しかし、それでも十分ではありませんでした。ありがたいことに、ESP32-S3 は 8.5 MB の作業メモリに加えて、最大 16 MB のフラッシュ ストレージも購入できます。
Google の Gemma モデル ファミリからレイヤーごとの埋め込み (PLE) と呼ばれる技術を借用することで、開発者はモデルの重みの大部分 (約 2,500 万個のパラメータまたは約 12 MB 相当) を、パフォーマンスの低下を最小限に抑えながらフラッシュにオフロードすることができました。
モデルの重みを NVMe ストレージにオフロードすることは、DeepSeek V3 のような大規模なフロンティア クラスのモデルを、他の方法で提供するのに必要なメモリと GPU 容量を持たないハードウェア上で実行するための古い手法です。歴史的に見て、このアプローチの欠点は、パフォーマンスが低下することです。通常、1 秒あたりのトークンの代わりに、トークンごとに秒単位、場合によっては分単位で調べます。それは機能しますが、リモートでは実用的ではありません。
PLE は、これらの重みがかなり控えめにアクセスされるため、かなりうまく管理できます。これにより、DRAM または SRAM に比べてフラッシュの非常に遅い帯域幅がパフォーマンスの低下を防ぐことができます。
その結果、ESP32 のメモリに 14.9 MB を詰め込むのではなく、モデルに必要なメモリは約 2 MB だけになりました。 S

具体的には、出力ヘッド、エンベディング、および KV キャッシュはチップの PSRAM に保持され、アクティブ化はチップの 520 KB の SRAM で処理されます。
このアプローチを使用して、開発者はマイクロコントローラーから 1 秒あたり 9.88 個のトークンを取得できたと述べています。これは、平均的な人が読み取ることができる速度よりも速いです。
AI のガードレールを回避するのは、スクリプト小僧でもできるほど簡単です
米国の模型メーカーがパニックになる中、中国はオープンモデルの電撃で熱を高める
MediaTek、AI データセンター推進のために 50 億ドルの軍資金を用意
Google 開発キットが史上初のエージェント間の暴力を促進
ESP32 のようなマイクロコントローラー上で小さな生成 AI モデルを実行することはできるかもしれませんが、実践から得られるものは愚かで単純なプライドを超えたものではありません。
Tiny Stories は優れた概念実証ですが、短くて適度に一貫したストーリーをオンデマンドで生成すること以外には、それほど多くのことはできません。それを使用してチャットボットを構築したり、コードを生成したり、エージェントを強化したりするつもりはありません。 Barista と呼ばれる別のモデルもありますが、これはおよそ 2 倍のパフォーマンスで質問に答えることができますが、質問に答えることができるのはエスプレッソに関するトピックのみです。
TinyStories と Barista は、まあ、小さすぎて他のことができません。それでも、これらのモデルが ESP32 上で動作するという事実自体が印象的です。
とはいえ、Raspberry Pi やスマートフォンなど、もう少しメモリとコンピューティング能力が高いデバイスをお持ちであれば、はるかに高性能なモデルが世の中にはあります。
4 月に発売された Google の Gemma 4-E2B-it は、同じ量子化および PLE オフロード技術を採用して、4 ビットの重みを使用する場合、51 億パラメータのビジョン言語モデルを 1 ギガバイト強のメモリに詰め込みます。量子化とオフロードをより高度に行うと、これを約 500 MB まで削減できます。
メモリ フットプリントはフロンティア モデルの実行に必要なメモリ フットプリントよりも数桁小さいですが、それでも十分なパフォーマンスを発揮します。

ローカルのチャットボットを制御し、ユーザーのカレンダー管理などのためにローカル エージェントを調整し、時折幻覚を我慢できる限り、OpenAI や Anthropic への依存から解放されます。 ®
エッジとIoT
開発者は、LLM があらゆるもの (10 ドルのマイクロコントローラーであっても) で動作することを証明しました
ほぼ 10 トーク/秒で、ほとんど一貫性があります — 気に入らない点は何ですか?
OpenAI は教師や教授が ChatGPT で仕事を集中することを望んでいます
K-12 の教師、大学の講師、大学生向けの新しいプラグインはありますか?部屋を読んでください、みんな
個人の成果からパートナーへの影響まで
パートナーのコンテンツ: Databricks の調査データは、認定プロフェッショナルがパートナーの提供能力、顧客の信頼性、AI への対応力を、資格を取得した個人をはるかに超えて向上させていることを示唆しています。
Next.js 16.3 は、恐ろしい FATAL ERROR メッセージを減らすことを目的としています
新しい React フレームワークによりメモリ使用量が 90% 削減されるとチームは主張
コラムニスト
クロード・コードはデジタル考古学に革命をもたらしています。エンタープライズはもっと掘ったほうがいいよ
「大のAI懐疑論者として、あなたには本当に驚かされました」
AI のガードレールを回避するのは、スクリプト小僧でもできるほど簡単です
「これは私のサーバーです」と主張するだけで、モデルを説得して協力してもらうことができることがよくありました
OSプラットフォーム
IT 上司は子供を職場に連れて行く日のために root セッションを開いたままにした
ソフトウェア
技術者は退職後からソフトウェアをサポートするために誘い出されました。覚えていることだけです
デボプス
HashiCorp の名前の由来となった開発者が、より高速なターミナル マルチプレクサを携えて戻ってきました。
ネットワーク
ボーダフォンが合併パートナーを買収し、3社が1社に
aiとml
AI バブルはすでにはじけつつあります。私たちはまだそれを知らないだけです
BSides、Black Hat、DEF CON がラスベガスに上陸する際に期待されること
Alibaba の Qwen チームは初めて、その「Max」モデルを API ペンから解放します。一方、DeepSeek V4-Flash は、安くて陽気な競争に新たな意味を与えます
ポイズニングされたプル リクエストには、プロンプト インジェクションが含まれており、

一方を下げて他方を制御する
88 個のカスタム コア、176 個のファンキー スレッド、1.5 TB のラップトップ RAM、および 1.8 TB/秒の NVLink 接続 - これは一般的なデータセンター チップではありません
クラウド収益は現在、四半期あたり 1,430 億ドルを超え、成長は加速しています
プラス：米国はイランのプロパガンダサイトを削除。マーケティング会社が「なぜあなたの情報を入手するのですか？」と尋ねます。さらに！
プラス：中国はスマートフォン監視ツールをアップグレード。リングは覗き見防止の姿勢を緩和します。などなど
ジェフ・モス氏によると、投票村のレポートは非常に成功しており、今後は DEF CON 全体が含まれることになる
会社全体の評価額は35億ドル以上に相当するが、売却部分は特定されていない
プラスの面としては、情報セキュリティは長く安定したキャリアを築くのに適しています。
FOSS は Microsoft の独占を 1 つ打ち破りました。 20年間の失敗を経て、次の失敗をする時が来た
一言
GNOME は Windows のように見えることができ、Flashback は拡張機能なしで実行できます
新しい「シンプルタスクバー」はオプションですが、よりシンプルで安定した方法があります
x86-32 での Debian の最終リリースに黙祷を捧げてください
新しい Debian バージョンが 13.6 および 12.15 の形で FOSSland に登場
脆弱な Joomla Web サイトで拡張機能のバグを悪用し、10 点満点を獲得した悪者を発見
iCagenda、Balbooa Forms 拡張機能の欠陥は、世界中の 100 万のサイトを支えるオープンソース CMS に影響を与える可能性があります
フレーム: 新しい X11 サーバー – アセンブリに直接実装
yserver、Phoenix、そしてもちろん XLibre、そして外れ値の Arcan に参加します
Cinnamon 6.8 は Wayland をサポートします – 必要に応じて
Linux Mint デスクトップの次期バージョンには両方の種類のディスプレイ サーバーが搭載されています
お問い合わせ
私たちと一緒に宣伝しましょう
私たちは誰なのか
ニュースレター
次のプラットフォーム
開発クラス
ブロックとファイル
状況出版
クッキーポリシー
プライバシーポリシー
利用規約
私の個人情報を共有しないでください
同意のオプション
著作権。無断複写・転載を禁じます © 1

998年から2026年。

## Original Extract

Nearly 10 tok/s and it

Jump to main content
Search
TOPICS
Special Features
All Special Features
Cloud Infrastructure Month 2026
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
Dev proves LLMs will run on anything – even a $10 microcontroller
Tobias Mann
Tobias
Mann
SYSTEMS EDITOR
Published
tue 4 Aug 2026 // 22:20 UTC
Getting a small local language model running on a notebook or even smartphone in 2026 is trivial. But what about something even smaller and lower-power. Say, like an ESP32 microcontroller that costs less than $10?
It might sound impossible — the device is primarily designed for things like remote sensors, IoT, and other embedded applications, not running generative AI models — yet, that's exactly what a developer who goes by the handle SlvDev has managed to do.
In a process detailed on GitHub , and recently showcased on the Better Stack YouTube channel , SlvDev documented how he managed to get a small language model running at nearly 10 tokens a second locally on a microcontroller that costs about the same as a fancy cup of coffee.
Tiny stories on a tiny microcontroller
Cramming a large language model (LLM) onto something as small as a ESP32 microcontroller isn't a trivial task.
There's a reason that these models are trained and run on GPUs. LLMs are memory-hungry beasts that typically require between one and four bytes per parameter just to hold their weights in memory.
With just 520 KB of SRAM and 8 MB of pseudo SRAM (PSRAM) on the ESP32-S3, you aren't going to be running a model like DeepSeek V4 Flash .
To make it work, the dev had to drop the "large" from the language model and settle for something nearly 10,000 times smaller: TinyStories, a 28.9 million-parameter model originally developed by Microsoft Research.
However, even this model is asking a lot of an ESP32-S3 module. At 16-bit precision, the model requires about 60 MB of memory that the ESP32 simply doesn't have. So, the dev employed several techniques, some of which we've previously explored, to shrink the model’s footprint.
The first is quantization, a process by which weights are compressed by reducing their precision from something like 16-bits of precision to eight, or even four.
This enabled SlvDev to trade a bit of accuracy for a 75 percent reduction in memory required. Instead of about 60 MB of memory to hold the weights, they now require just 14.9 MB.
But that still wasn’t enough. Thankfully, in addition to the 8.5 MB of working memory, the ESP32-S3 can also be bought with up to 16MB of flash storage.
By borrowing a technique called per-layer-embedding (PLE) from Google's Gemma family of models, the dev was able to offload the majority of the model's weights, about 25 million parameters or about 12 MB worth, to flash with minimal performance degradation.
Offloading model weights to NVMe storage is an old trick for getting massive frontier-class models like DeepSeek V3 running on hardware that wouldn't have the necessary memory and GPU capacity to serve it otherwise. The downside of this approach, historically, is that it murders performance. Instead of tokens a second, you're usually looking at seconds, or in some cases minutes, per token. It works but it's not remotely practical.
PLE manages quite a bit better because these weights are accessed rather sparingly, which keeps the flash's glacially slow bandwidth relative to DRAM or SRAM from nerfing performance.
The result is that rather than trying to cram 14.9 MB into the ESP32's memory, the model now only needs about 2 MB. Specifically the output head, embeddings, and KV cache are kept in the chip's PSRAM, while activations are handled in the chip's 520 KB of SRAM.
Using this approach, the dev says they were able to get 9.88 tokens a second out of the microcontroller, which is faster than the average person can read.
Bypassing AI guardrails is so easy a script kiddie can do it
China turns up the heat with open model blitz as US model makers panic
MediaTek lines up $5B war chest for AI datacenter push
Google dev kit spurs first-ever agent-on-agent violence
While you may be able to get a small generative AI model running on a microcontroller like an ESP32, you won't get much from the practice beyond dumb simple pride.
Tiny Stories is a great proof-of-concept, but aside from generating short, reasonably coherent stories on demand, it can't do much. You aren't going to build a chatbot with it, generate code, or power an agent. There is another model, called Barista, that can answer questions at roughly twice the performance, but only on topics pertaining to espresso.
TinyStories and Barista are, well, just too tiny to do much else. Yet the fact these models run on an ESP32 at all is impressive in itself
That said, if you've got a device with just a bit more memory and compute, say a Raspberry Pi or a smartphone, there are far more capable models out there.
Google's Gemma 4-E2B-it, launched back in April, employs the same quantization and PLE offload techniques to cram a 5.1 billion-parameter vision language model into just over a gigabyte of memory when using 4-bit weights. Higher degrees of quantization and offloading can get this down to around 500 MB.
While its memory footprint is several orders of magnitudes less than what's required to run a frontier model, it is still capable enough to power local chatbots, orchestrate local agents for things like managing a user's calendar, and free you from your reliance on OpenAI or Anthropic so long as you can put up with the occasional hallucination. ®
EDGE AND IOT
Dev proves LLMs will run on anything – even a $10 microcontroller
Nearly 10 tok/s and it's mostly coherent — What's not to like?
OpenAI wants teachers and profs to foist their work off on ChatGPT
New plugins for K-12 teachers, university instructors, and college kids? Read the room, guys
From individual achievement to partner impact
PARTNER CONTENT: Databricks survey data suggests certified professionals lift partner delivery capacity, customer credibility, and AI readiness well beyond the individuals who earn the credential
Next.js 16.3 aims to reduce dreaded FATAL ERROR messages
New React framework lowers memory usage by 90%, team claims
COLUMNISTS
Claude Code is revolutionizing digital archaeology. Enterprise better dig it
'As a big AI skeptic, you just blew my mind'
Bypassing AI guardrails is so easy a script kiddie can do it
Claiming 'it's my server' was often enough to persuade models to help
OS PLATFORMS
IT boss left root session open for bring-your-kid-to-work day
SOFTWARE
Techie lured out of retirement to support software only he remembered
DEVOPS
Dev who gave HashiCorp its name returns with a faster terminal multiplexer
NETWORKS
Three becomes one as Vodafone buys out merger partner
ai and ml
The AI bubble is already popping; we just don't know it yet
What to expect as BSides, Black Hat, and DEF CON descend on Las Vegas
For the first time, Alibaba's Qwen team is letting its 'Max' model out of the API pen; meanwhile, DeepSeek V4-Flash gives new meaning to cheap and cheerful competition
Poisoned pull requests contain prompt injection that allows one to control another
88 custom cores, 176 funky threads, 1.5 TB of laptop RAM, and 1.8 TB/s of NVLink connectivity — this isn't your typical datacenter chip
Cloud revenue now north of $143 billion a quarter, and growth is accelerating
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
FOSS smashed one Microsoft monopoly. After 20 years of failure, it's time to smash another
Word up
GNOME can look like Windows – and Flashback can do it without extensions
New 'Simple-taskbar' is an option, but there's a simpler, stabler way
A moment of silence, please, for the final release of Debian on x86-32
New Debian versions hit FOSSland in the form of 13.6 and 12.15
Baddies caught exploiting extensions bugs with perfect 10 scores on vulnerable Joomla websites
Flaws in iCagenda, Balbooa Forms extensions can impact open source CMS that powers a million sites worldwide
Frame: A new X11 server – implemented directly in assembly
Joins yserver, Phoenix, and of course XLibre – and outlier Arcan
Cinnamon 6.8 will support Wayland – if you want it
Next version of Linux Mint’s desktop has both kinds of display server
Contact us
Advertise with us
Who we are
Newsletter
The Next Platform
DevClass
Blocks and Files
Situation Publishing
Cookies Policy
Privacy Policy
Ts & Cs
Do not share my personal information
Your Consent Options
Copyright. All rights reserved © 1998-2026.
