---
source: "https://spectrum.ieee.org/llm-training-energy-saving-trick"
hn_url: "https://news.ycombinator.com/item?id=48478848"
title: "Tweaking GPU Clock Frequency Cuts LLM Training Energy"
article_title: "Tweaking GPU Clock Frequency Cuts LLM Training Energy - IEEE Spectrum"
author: "rbanffy"
captured_at: "2026-06-10T19:02:38Z"
capture_tool: "hn-digest"
hn_id: 48478848
score: 3
comments: 0
posted_at: "2026-06-10T16:34:07Z"
tags:
  - hacker-news
  - translated
---

# Tweaking GPU Clock Frequency Cuts LLM Training Energy

- HN: [48478848](https://news.ycombinator.com/item?id=48478848)
- Source: [spectrum.ieee.org](https://spectrum.ieee.org/llm-training-energy-saving-trick)
- Score: 3
- Comments: 0
- Posted: 2026-06-10T16:34:07Z

## Translation

タイトル: GPU クロック周波数の調整により LLM トレーニング エネルギーが削減される
記事のタイトル: GPU クロック周波数の調整により LLM トレーニング エネルギーが削減される - IEEE Spectrum
説明: カーネルレベルの DVFS は、GPU クロック周波数を調整することで LLM トレーニングを最適化し、計算速度を犠牲にすることなく 14% のエネルギー節約を実現します。

記事本文:
-->
Raven.config('https://6b64f5cc8af542cbb920e0238864390a@sentry.io/147999').install();
GPU クロック周波数の調整により LLM トレーニングのエネルギーが削減 - IEEE Spectrum
IEEE.org IEEE Explore IEEE 標準 IEEE 求人サイト その他のサイト サインイン 参加 IEEE タイミング トリックにより、LLM トレーニングで使用されるエネルギーが最大 14 パーセント削減 テクノロジー インサイダー検索: トピック別に探索 航空宇宙 AI 生物医学 気候 テクノロジー コンピューティング 家庭用電化製品 エネルギー 技術の歴史 ロボット工学 半導体 電気通信 輸送
IEEEスペクトル
テクノロジー関係者向けトピック
アカウントを作成すると、さらに無料のコンテンツや特典をお楽しみいただけます
後で読むために記事を保存するには、IEEE Spectrum アカウントが必要です
インスティチュートのコンテンツはメンバーのみが利用できます
PDF 版全号のダウンロードは IEEE メンバー限定です
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
ACC

ess 数千の記事 — 完全無料
アカウントを作成して限定コンテンツと機能を入手してください:
記事の保存、コレクションのダウンロード、
そして
コメントを投稿する
— すべて無料です！フルアクセスと特典については、
購読する
スペクトルに。
タイミング トリックにより LLM トレーニングに使用されるエネルギーを最大 14% 削減
計算中にクロック周波数を調整すると、パフォーマンスに影響を与えることなくエネルギーを節約できます。
Dina Genkina は、IEEE Spectrum のコンピューティングおよびハードウェア編集者です。
エリック・フロムメルト
OpenAI の 4 番目の大規模言語モデル (LLM) である GPT-4 のトレーニングには、推定 50 ギガワット時間かかりました。これは、アメリカの家庭 5,000 世帯の年間電力消費量に相当します。それは 2023 年のことでした。それ以来、フロンティア LLM のトレーニングに使用される計算リソースは増加するばかりですが、直接の電力使用量の数値を入手するのは困難です。
今回、オランダのトゥエンテ大学の研究グループは、計算中に GPU のクロック周波数を賢く調整することで、速度を犠牲にすることなく LLM トレーニングに使用されるエネルギーを最大 14% 節約できることを示しました。ジェフリー・スパン博士トゥウェンテ大学の候補者であり論文の筆頭著者である同氏は、先月シチリア島カターニアで開催されたコンピューティングフロンティアカンファレンスでその結果を発表した。
「私の研究は、コンピューティングの無駄を見つけることです」とスパーン氏は言います。 「これはハードウェアの活用不足に似ていますが、ハードウェアに合わせてソフトウェアを最適化するのではなく、ソフトウェアに合わせてハードウェアを最適化しようとします。」
Spaan 氏と彼の共同研究者は、動的電圧周波数スケーリング (DVFS) として知られる技術を使用してこれを達成しました。フロンティア モデルのトレーニングに一般的に使用される GPU を含むすべてのチップは、少なくとも 1 つのクロックを使用して計算を調整します。チップ内の各動作はクロック パルスによってトリガーされます。その時計の周波数は

ティックは、チップの動作速度と消費電力を制御します。
最新の GPU には 2 つのクロックがあり、1 つは計算コア用、もう 1 つはメモリ用です。コアが数値の処理に熱心に取り組んでいるときは、高速な計算を保証するためにクロック周波数が高く保たれます。ただし、DVFS を使用すると、その間メモリ クロックが遅くなり、消費電力が少なくなります。原理的には、チップのメモリ部分をオフにするだけで済みますが、GPU の設計ではそのオフ スイッチをソフトウェアで制御することができず、いずれにしても計算途中でオンに戻すには時間がかかりすぎます。同様に、コアがメモリからデータがロードされるのを待っているとき、メモリ クロック周波数が上昇する一方で、コア クロック周波数が最高に遅くなる可能性があります。
DVFS は、少なくとも 1990 年代に遡るよく知られた技術です。しかしスパーン氏は、他の研究者らは、彼らの方法では計算が遅くなりすぎるか、エネルギー使用量を改善するほど細かく設定されていなかったため、それを LLM トレーニングに有効に適用できなかったと述べている。
以前の DVFS の試みでは、トレーニング プロセスの反復ごとに頻度が調整されました。 LLM トレーニングでは、各反復は 2 つの部分で構成されます。 前方パスでは、重みがそのままの状態でモデルの層を介してデータが前方に実行されます。もう一方は、前方パスの結果に基づいて重みが層ごとに調整される逆伝播です。したがって、以前の研究では、順方向パスの周波数の 1 つの値を維持し、バックプロパゲーション用に別の値に調整していました。
Spaan 氏と同僚は、より短い時間スケールでクロック周波数を調整しました。 GPU ワークロードは、カーネルと呼ばれる小さな計算ナゲットに分割されます。たとえば、単一のベクトル間の乗算で単一のカーネルを構成できます。カーネルは GPU に供給されて何度も処理されます

並行して。 Spaan の実装では、ディープ ニューラル ネットワークの単一層の計算が約 40 のカーネルに分割されます。カーネルごとのレベルでクロック周波数を調整することで、チームは大幅にエネルギーを節約することができました。
また、GPU は、チップの内部システムが需要の増減を検出したときに、DVFS を自動的に実行するとスパーン氏は指摘します。 「したがって、GPU に処理させればよいのではないかと考える人もいるかもしれません」と彼は言います。 「ただし、GPU には、どのようなカーネルが実行されるかについての先見性がありません。そのため、その場でのベストエフォートの推測に基づいて動作する必要があり、そのため、同じ節約を達成することはできません。」そこで手動調整が登場します。
チームは、Nvidia RTX 3080 Ti GPU で 13 億パラメータ モデルである GPT-3-xl をトレーニングすることで実験を実行しました。時間を節約するために、彼らはモデルの単一層のトレーニングに重点を置きました。この設定では、トレーニング時間をわずか 0.6 パーセント遅らせながら、エネルギーを 14 パーセント節約できる一連の周波数調整を発見しました。モデルのパフォーマンスは、計算速度とエネルギー使用量の両方に依存します。
課題が 1 つあります。クロック周波数のランプダウンは、コアをオフにしてからオンにするよりもはるかに高速ですが、それでも瞬時ではありません。研究者らは実験で、周波数切り替え速度を考慮せずに、一度に 1 つのカーネルを評価しました。したがって、14% のエネルギー節約が最良のシナリオとなります。実際にそれがどの程度の問題になるかは、使用されている GPU に大きく依存するとスパーン氏は言います。 Blackwell GPU などの新しいハードウェアは、古いバージョンよりもスイッチング速度がはるかに速く、省電力を最大限に活用できるはずです。
現在、チームは、最適な周波数スケーリングを自動的に実装できるツールを開発中です。

特定の作業負荷。スパーン氏は、彼らの手法が業界リーダーにとって採用に値するほど魅力的なものになることを期待している。 「私たちはパフォーマンスを損なうことなくエネルギーを節約できるように最適化しています」とスパーン氏は言います。 「現実の世界では、パフォーマンスが究極の目標です。」
より優れたハードウェアがゼロを AI ヒーローに変える可能性 ›
今日の生成 AI のエネルギー問題は根本的なものです ›
人間のエネルギーフットプリントと大規模言語モデル – ACM のコミュニケーション ›
2022 年の LLM トレーニング時のエネルギー消費量 (MWh) ›
Dina Genkina は、コンピューティングとハードウェアを専門とする IEEE Spectrum の副編集者です。彼女は原子物理学の博士号を取得しており、ブルックリンに住んでいます。
パノプティコンをクラウドソーシングしています
器用さを超えて: なぜ接触がロボット工学の次の時代を定義するのか
最初の原爆実験に対する衝撃的な新たな見解
生成 AI のエネルギー需要が私たちの世界を再構築している
AI チャットボットは医師のように推論できるか?
AI はより優れた AI を構築し始めています

## Original Extract

Kernel-level DVFS optimizes LLM training by adjusting GPU clock frequency, delivering 14% energy savings without sacrificing computation speed.

-->
Raven.config('https://6b64f5cc8af542cbb920e0238864390a@sentry.io/147999').install();
Tweaking GPU Clock Frequency Cuts LLM Training Energy - IEEE Spectrum
IEEE.org IEEE Xplore IEEE Standards IEEE Job Site More Sites Sign In Join IEEE Timing Trick Cuts Energy Used in LLM Training by Up to 14 Percent Share FOR THE TECHNOLOGY INSIDER Search: Explore by topic Aerospace AI Biomedical Climate Tech Computing Consumer Electronics Energy History of Technology Robotics Semiconductors Telecommunications Transportation
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
Timing Trick Cuts Energy Used in LLM Training by Up to 14 Percent
Adjusting clocking frequency during computation can save energy without affecting performance
Dina Genkina is the computing and hardware editor at IEEE Spectrum
Eric Frommelt
OpenAI ’s fourth large language model (LLM), GPT-4 , took an estimated 50 Gigawatt-hours to train, or the equivalent of 5,000 American homes ‘ yearly power consumption. That was in 2023. Since then, the computational resources used to train frontier LLMs have only increased , though direct power usage numbers are hard to come by.
Now, a research group at University of Twente in the Netherlands has shown that you can save up to 14 percent of the energy used in LLM training without sacrificing speed by cleverly adjusting the clock frequency of the GPU during computation. Jeffrey Spaan , Ph.D. candidate at University of Twente and lead author on the article, presented the results at the Computing Frontiers conference in Catania, Sicily last month.
“My research is about finding computing waste,” Spaan says. “It’s similar to underutilization of the hardware, but instead of optimizing the software for the hardware, we try to optimize the hardware for the software.”
Spaan and his collaborators accomplished this by using a technique known as dynamic voltage-frequency scaling ( DVFS ). Every chip—including the GPUs commonly used for training frontier models—uses at least one clock to orchestrate computations. Each operation in the chip is triggered by a clock pulse. The frequency with which that clock ticks controls how fast the chip operates, and how much power it draws.
Modern GPUs have two clocks , one for the computational core and one for the memory. When the core is hard at work crunching numbers, the clock frequency is kept high to ensure speedy calculation. However, with DVFS, the memory clock can slow down in that time, allowing for less power draw. It’s in principle possible to just turn off the memory part of the chip, but GPUs designs don’t enable software control for that off switch, and it would take too long to turn back on mid-calculation anyway. Similarly, when the core is waiting for data to be loaded from memory, the core clocking frequency can be slowed to a crawl while the memory clock frequency ramps up.
DVFS has been a well-known technique that goes back to at least the 1990s. But Spaan says other researchers haven’t been able to usefully apply it to LLM training because their methods either slowed down calculations too much or were not fine-grained enough to improve energy usage.
Previous DVFS attempts adjusted the frequency at each iteration of the training process. In LLM training, each iteration consists of two parts: the forward pass, in which data is run forward through the layers of the model with the weights as they are, and backpropagation , in which the weights are adjusted layer by layer based on the results of the forward pass. So, prior work kept one value of the frequency for the forward pass and adjusted to another for backpropagation.
Spaan and co-workers tuned the clock frequencies on a shorter time scale. GPU workloads are broken down into tiny computational nuggets known as kernels . For example, a single vector-vector multiplication can make up a single kernel. The kernels are fed to the GPU to be processed many times in parallel. In Spaan’s implementation, the computation of a single layer of a deep neural network is broken up into approximately 40 kernels. By adjusting the clocking frequencies on a per-kernel level, the team was able to find much greater energy savings.
The GPU also does DVFS automatically when the chip’s internal systems detect higher or lower demand, Spaan notes. “Some people might therefore think: we’ll just let the GPU handle it,” he says. “However, because the GPU doesn’t have the foresight we have of what kernels will run, it has to work with an on-the-fly best-effort guess, and can therefore never attain the same savings.” That’s where the manual adjustments come in.
The team performed their experiment by training GPT-3-xl, a 1.3 billion parameter model, on an Nvidia RTX 3080 Ti GPU. To save time, they focused on training a single layer of the model. In this setting, they found a set of frequency adjustments that gave them 14 percent energy savings while slowing the training time by only 0.6 percent. Performance of the model depends on both computing speed and energy usage.
There is one challenge: Ramping down the clock frequency is much faster than turning a core off and on, but it’s still not instantaneous. In their experiment, the researchers evaluated one kernel at a time, not taking into account the frequency switching speed. So, 14 percent energy savings is a best-case scenario. How much of an issue it would be in practice, Spaan says, depends heavily on the GPU being used. Newer hardware, like the Blackwell GPUs, have much faster switching speeds than older versions, and should be able to harness the full energy savings.
Now, the team is developing a tool that would be able to implement optimal frequency scaling automatically for a particular workload. Spaan hopes their method will be attractive enough to industry leaders to merit adoption. “We optimize for saving energy without losing performance,” Spaan says. “In the real world, performance is the holy grail.”
Better Hardware Could Turn Zeros into AI Heroes ›
Generative AI’s Energy Problem Today Is Foundational ›
The Energy Footprint of Humans and Large Language Models – Communications of the ACM ›
Energy consumption when training LLMs in 2022 (in MWh) ›
Dina Genkina is an associate editor at IEEE Spectrum focused on computing and hardware. She holds a PhD in atomic physics and lives in Brooklyn.
We Are Crowd-Sourcing the Panopticon
Beyond Dexterity: Why Contact May Define the Next Era of Robotics
Striking New Views of the First Atomic Bomb Test
Generative AI's Energy Needs Are Reshaping Our World
Can AI Chatbots Reason Like Doctors?
AI Is Starting to Build Better AI
