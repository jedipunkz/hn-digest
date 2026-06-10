---
source: "https://ziraph.com/blog/on-device-ai-margin-decision"
hn_url: "https://news.ycombinator.com/item?id=48480247"
title: "On-device AI is a margin decision"
article_title: "On-device AI is a margin decision - Ziraph blog"
author: "ABS"
captured_at: "2026-06-10T19:00:09Z"
capture_tool: "hn-digest"
hn_id: 48480247
score: 1
comments: 0
posted_at: "2026-06-10T18:04:53Z"
tags:
  - hacker-news
  - translated
---

# On-device AI is a margin decision

- HN: [48480247](https://news.ycombinator.com/item?id=48480247)
- Source: [ziraph.com](https://ziraph.com/blog/on-device-ai-margin-decision)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T18:04:53Z

## Translation

タイトル: オンデバイス AI は僅差の決定です
記事タイトル: オンデバイス AI はマージンの決定 - Ziraph ブログ
説明: どこで推論を実行するかは、マージン、プライバシー、バッテリー、ツール、インストール ベースに影響を与える決定です。エンジニアリング リーダーがオンデバイス AI に賭けるか、除外する前に考慮すべき側面。

記事本文:
オンデバイス AI はマージンの決定です - Ziraph ブログ
ベータ版（ソフト）近日公開予定！お知らせリストを購読して、Web サイトとベータ版がオープンする瞬間を聞いてください ↗
オンデバイス AI はマージンの決定です
AI 機能を追加するのは簡単です。結果的なものはより静かであり、それは推論が実際にどこで実行されるか、つまり他の人のサーバー上か、ユーザーが手に持っているデバイス上かという問題です。
· 2026 年 6 月 10 日 · 6 分で読めます
その 1 つの選択が、売上総利益、プライバシーへの取り組み、アプリストアのレビュー、そして上級エンジニアの時間のどれだけがハードウェア考古学に費やされるかに影響を及ぼします。だからこそ、計画文書では 1 行以上の価値があるのです。
マージンのケースはそれらの中で最も騒々しいので、そこから始めますが、これはオンデバイス推論へのロードマップをコミットするか、それを除外することを決定する前に検討する価値のあるいくつかの側面のうちの 1 つにすぎません。
クラウド推論は、すべてのユーザーおよびすべてのトークンに応じて拡大し、決して停止しない商品原価であるため、成功するにつれて悪化する稀なソフトウェア項目となっています。デバイス上の推論ではこの状況が逆転します。トークンの限界コストは、ユーザーがすでに購入したシリコンによって支払われるため、ほぼゼロになるためです。そのため、小さなアドオンではなく AI が機能の中心となる製品では、デバイス上での推論の移動は、テーブル上で最大の粗利レバーの 1 つになります。
これは、チェックされるのではなく、機能すると想定されることが最も多いレバーでもあり、そのギャップが、以下の他のすべての側面を貫く糸となります。
健康、金融、法務、またはユーザーの個人的な状況など、機密データに触れるあらゆるものについて、「デバイスから離れることはありません」というラインは、料金を請求できる機能であると同時に、企業の販売サイクルを短縮するコンプライアンス ストーリーでもあり、デバイス上で推測できます。

CEは、それを言う権利を獲得する方法です。
問題は、「ローカル」は実際にローカルでなければならないということです。WWDC 2026 は、Apple の「オンデバイス」基盤モデルでもクラウドにルーティングできることを思い出させる有益なものでした。フラッグシップ層は Google のクラウドの NVIDIA GPU で実行され、そのスイッチがいつ切り替わるのかについての公開信号はありません。そのため、プライバシーの主張は、結局のところ、作業が実際にユーザーのシリコンに残っていることを確認する能力と同じくらい強力なものになるだけです。
3. バッテリーが実際の予算です
オンデバイス AI の障害モードがサーバーの請求として現れることはほとんどありません。ユーザーのバッテリーが消耗すると、GPU を固定して携帯電話を加熱する機能により、請求に関するアラートがまったく表示されず、1 つ星のレビューとアンインストールだけが表示されるためです。
バッテリー駆動のデバイスでは、トークンあたりのエネルギーと熱ヘッドルームはバックエンドの数値ではなくなり、製品品質の指標になります。トレードオフは通常、皆さんが推測するものではありません。トークンあたりのエネルギーを 2 倍消費しながら 10% 高速に動作するモデルは、多くの場合、電話機にとって悪い選択であるためです。スループットだけでは表面化することはありません。そのため、サポート チケットが到着し始めた後ではなく、機能がリリースされる前に測定する必要があります。
4. スタックは断片化しているため、慎重に選択してください
1 年前、「Apple 上のデバイス上」というと主に Core ML を意味していましたが、現在では Core ML、Core AI、MLX、Ollama、および llama.cpp がすべて重複する remit と共存しており、特定のチームがどれに到達すべきかについての明確な答えはありません。
エンジニアは 1 つまたは複数を選択しますが、あなたにとって重要なのは、ファーストパーティのツールが均一ではないということです。Apple 独自のプロファイラーはタイミングのみで Xcode にバインドされており、独自のアプリのみを参照できるため、Ollama や MLX、または任意のプロセスをプロファイリングできず、エネルギーもメモリ帯域もありませんと報告されます。

dth であり、熱状態はまったくありません。
どのスタックに到達しても、単一のフレームワークに固定されたツールではなく、スタック全体にわたる可視性が必要になります。
5. ユーザーは開発マシン上にいません
ラップトップの M シリーズ上で動作するモデルは、構築したマシンではなくインストール ベースの中央のデバイスで実行されると、スロットルしたり、静かに低速のアクセラレータに戻ったりすることがあります。
どのオンデバイス モデルを入手するかでさえ、選択可能ではなくハードウェア ゲートです。Apple は、ほとんどのチップに小型の高密度モデルを出荷し、最も性能の高いシリコン用に大型のモデルを予約しているためです。したがって、決定は実際にユーザーが実際に搭載するハードウェア、実際に出荷する量子化およびコンテキスト長に基づいて行われるべきであり、トップエンド Mac でたまたま実行した最良の場合のベンチマークに対して行う必要はありません。
6. 知らないことの代償は上級エンジニアの時間です
「ニューラル エンジン上にあるのか、それとも後退しているのでしょうか? 長いコンテキストで速度が低下したのはなぜですか?」これは、ツールが 1 つのコマンドで回答する必要がある調査作業に、最も高価なエンジニアが何日も費やすことを意味します。
最初のコストの中に 2 番目のコストが隠れています。なぜなら、チームがあなた、取締役会、顧客に上向きに報告する数値は、ベンダーの TOPS 仕様や希望に満ちた推定ではなく、測定され、擁護できるものでなければなりません。また、懐疑的なエンジニアが耐えられないオンデバイスのパフォーマンスに関する主張は、そもそもデッキに含める必要がありません。
すべてはハードウェアにかかっているというフレーズ
これら 6 つをもう一度読むと、すべての上昇局面には同じ条件があることがわかります。つまり、それはユーザーが実際に所有しているハードウェアで動作する必要があり、まさにそれがスライドだけでは判断できない部分です。
TOPS 番号はあなたのことを何も教えてくれません

トークンごとのモデルのエネルギー、「ニューラル エンジンで実行される」ということは、何かがそれを測定するまでは仮定のままです (Apple 独自のモデルはアクセラレータを介してフォールバックします。そしてあなたのモデルも同様です)。量子化が持続するかどうか、実際に使用するコンテキスト長でメモリに収まるかどうか、持続的な負荷に耐えるかどうかは、すべてチップ上のモデルについて測定された事実です。
これらの事実が、計画文書に書き込んだ利点が実際のデバイス上で実際に実現するか、現場で静かに消え去るかを決定します。
賭けのリスクを取り除きます。推測しないでください
これは、Ziraph が埋めるために構築されたギャップです。Apple シリコン上のワークロードごとのテレメトリにより、実際に作業を実行したアクセラレータ、トークンあたりのエネルギー、メモリ帯域幅とワークロードのボトルネックの場所、およびチームがすでに使用しているランナー (Ollama、MLX、llama.cpp、Core ML) のサーマル ヘッドルームをコマンド ラインおよび CI から確認できます。ループ内に Xcode はありません。
クラウドの料金を削減するものではなく、運用フリートのダッシュボードではありません。それは、ロードマップをコミットする前に、会議での議論から得られた上記の 6 つの側面を自分で測定した数値に変換するという、狭い範囲のことを 1 つ実行します。
オンデバイス AI は一流のプラットフォームになったばかりで、これで勝利するチームは最初に測定したチームになります。
やあ！ 👋私はマルコです。 Apple シリコン上でローカル AI 用のプロファイラーである Ziraph を構築します。オンデバイスがロードマップに含まれていて、それが自社の製品で実現可能かどうかを検討している場合、またはマージンとクラウドの計算を厳密にテストしたい場合は、喜んでメモを比較します。

## Original Extract

Where your inference runs is a decision with consequences in margin, privacy, battery, tooling, and your install base. The aspects an engineering leader should weigh before betting on on-device AI - or ruling it out.

On-device AI is a margin decision - Ziraph blog
Beta version (soft) launch soon! Subscribe to the announcement list to hear the moment the website and beta open ↗
On-device AI is a margin decision
Adding an AI feature is the easy call; the consequential one is quieter, and it is the question of where the inference actually runs - on someone else’s servers, or on the device in your user’s hand.
· June 10, 2026 · 6 min read
That single choice ripples into your gross margin, your privacy posture, your app-store reviews, and how much of your senior engineers’ time disappears into hardware archaeology, which is why it deserves more than a line in a planning doc.
The margin case is the loudest of those, so I start there, but it is only one of several aspects worth weighing before you commit a roadmap to on-device inference, or decide to rule it out.
Cloud inference is a cost of goods that scales with every user and every token and never stops, which makes it the rare software line item that gets worse precisely as you succeed. On-device inference inverts that, because the marginal cost of a token drops to roughly zero, paid for by silicon the user already bought, so for any product where the AI is central to what it does rather than a small add-on, moving inference onto the device becomes one of the largest gross-margin levers on the table.
It is also the lever most often assumed to work rather than checked, and that gap is the thread running through every other aspect below.
For anything that touches sensitive data, whether that is health, finance, legal work, or a user’s personal context, the line “it never leaves the device” is both a feature you can charge for and a compliance story that shortens enterprise sales cycles, and on-device inference is how you earn the right to say it.
The catch is that “local” has to actually be local: WWDC 2026 was a useful reminder that even Apple’s “on-device” Foundation Models can route to the cloud, with the flagship tier running on NVIDIA GPUs in Google’s cloud and no public signal of when that switch flips, so a privacy claim ends up being only as strong as your ability to verify the work really stayed on the user’s silicon.
3. The battery is your real budget
The failure mode of on-device AI rarely arrives as a server bill; it arrives as the user’s battery draining, because a feature that pins the GPU and heats the phone gives you no billing alert at all, only one-star reviews and uninstalls.
On a battery-powered device, energy per token and thermal headroom stop being backend numbers and become product-quality metrics, and the trade-offs are usually not the ones you would guess, since a model that runs ten percent faster while drawing twice the energy per token is often the worse choice for a phone, and that is exactly the kind of thing throughput alone will never surface, which is why you want it measured before the feature ships rather than after the support tickets start arriving.
4. The stack is fragmenting, so choose deliberately
A year ago, “on-device on Apple” mostly meant Core ML, whereas today Core ML, Core AI, MLX, Ollama, and llama.cpp all coexist with overlapping remits and no settled answer about which one a given team should reach for.
Your engineers will pick one or several, and the part that matters for you is that the first-party tooling is uneven: Apple’s own profiler is timing-only, bound to Xcode, and can see only your own app, so it cannot profile Ollama or MLX or an arbitrary process, and it reports no energy, no memory bandwidth, and no thermal state at all.
Whichever stack you land on, you will want visibility that spans it rather than tooling locked to a single framework.
5. Your users are not on your dev machine
A model that flies on the M-series in your laptop can throttle, or quietly fall back to a slower accelerator, once it is running on the median device in your install base rather than on the machine you build on.
Even which on-device model you get is hardware-gated rather than selectable, because Apple ships a smaller dense model to most chips and reserves the larger one for its most capable silicon, so the decision really has to be made against the hardware your users actually carry, at the quantization and the context length you actually ship, and not against the best-case benchmark you happened to run on a top-end Mac.
6. The cost of not knowing is senior-engineer time
“Is it on the Neural Engine or falling back? Why did it slow down at long context?” is, today, multi-day hardware archaeology, which means your most expensive engineers spend days on detective work that the tooling ought to answer in a single command.
There is a second cost hiding inside the first one, because the numbers your team reports upward, to you, to a board, or to a customer, should be measured and defensible rather than a vendor TOPS spec or a hopeful estimate, and a claim about your on-device performance that cannot survive a skeptical engineer has no business being in the deck in the first place.
The phrase it all hinges on: on the hardware
Re-read those six and you will notice that every upside carries the same condition, which is that it has to work on the hardware your users actually have, and that is precisely the part you cannot settle from a slide.
A TOPS number tells you nothing about your model’s energy per token, “it runs on the Neural Engine” stays an assumption until something measures it (Apple’s own models fall back across accelerators, and yours will too), and whether your quantization holds up, fits in memory at the context length you really use, and survives sustained load are all measured facts about your model on your chip.
Those facts are what decide whether the upside you wrote into the planning doc actually materializes on real devices, or quietly evaporates in the field.
De-risk the bet; don’t guess it
This is the gap Ziraph was built to fill: per-workload telemetry on Apple silicon that tells you which accelerator actually ran the work, the energy per token, the memory bandwidth and where the workload bottlenecks, and the thermal headroom, across the runners your team already uses (Ollama, MLX, llama.cpp, Core ML), from the command line and in CI, with no Xcode in the loop.
It will not cut your cloud bill for you, and it is not a production fleet dashboard. It does one narrow thing, which is to turn the six aspects above from arguments in a meeting into numbers you measured yourself, before you commit a roadmap to them.
On-device AI just became a first-class platform, and the teams that win with it will be the ones who measured first.
Hi there! 👋 I’m Marco. I build Ziraph, a profiler for local AI on Apple silicon. If on-device is on your roadmap and you’re weighing whether it’s viable for your product, or you want to pressure-test the margin-versus-cloud math, I’d be glad to compare notes.
