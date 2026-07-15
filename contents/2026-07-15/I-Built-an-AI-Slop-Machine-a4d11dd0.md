---
source: "https://cmart.blog/slop-machine/"
hn_url: "https://news.ycombinator.com/item?id=48926981"
title: "I Built an AI Slop Machine"
article_title: "Slop Machine | cmart's blog"
author: "chrismartin"
captured_at: "2026-07-15T21:49:24Z"
capture_tool: "hn-digest"
hn_id: 48926981
score: 2
comments: 0
posted_at: "2026-07-15T21:04:44Z"
tags:
  - hacker-news
  - translated
---

# I Built an AI Slop Machine

- HN: [48926981](https://news.ycombinator.com/item?id=48926981)
- Source: [cmart.blog](https://cmart.blog/slop-machine/)
- Score: 2
- Comments: 0
- Posted: 2026-07-15T21:04:44Z

## Translation

タイトル: AIスロップマシンを作ってみた
記事タイトル: スロップマシン | cmartのブログ

記事本文:
私は人間であり、このブログのすべての言葉を自分の声で書いています。それにもかかわらず、私はAIスロップマシンを構築しました。
(退屈! きれいなチャートにスキップしてください)
アメリカと中国の AI ラボは、ますます強力なオープンウェイト大規模言語モデル (Qwen、Gemma、gpt-oss など) を作成し続けています。これらの LLM は、一般消費者が入手可能なコンピューターおよび GPU 上で自己ホストできます。彼らはエージェント的なナレッジ ワーク タスクに非常に優れており、これらを使用すると、人々が AI について良くないと言うことのいくつかは当てはまりません。
テクノロジー企業があなたのプロンプトや応答データを扱うことは決してありません。それは自分のコンピュータ上に残ります。
謎の割り当て付きのサブスクリプションとプレミアム API の価格設定にこだわる必要はありません。ハードウェアを購入すると、モデルを無料でダウンロードして永久に使用できます。
自己ホスト型 LLM にアクセスする方法には、通常、広告やユーザーの注意を操作するものはありません。
モデルのコピーを入手すると、そのメーカーはそのモデルへのアクセスを取り消すことはできません。
AI データセンターに対する追加の需要を引き起こしているわけではありません。
はい、これらのモデルは多額の計算コストをかけてトレーニングされますが、コピーをダウンロードするまでにケーキの焼き上がりは完了しています。
使用する電力は、ビデオゲーム機と同様に、動作中に最大でも数百ワットです。
そのため、私は、より高スペックの Macbook を使用している友人が、ますます高性能化した AI を実用的な速度で実行し、分厚い電源ブロックをカフェに持ち込んで、熱くヒューヒュー音を立てるラップトップでエージェントを使用しているのを見てきました。
一方、私が日常的に使用しているラップトップ (第 13 世代 Intel Framework 13) はこの点が苦手です。パーティー トリックとして、48 GB のシステム RAM に適切な LLM を収めることができますが、2 つの DDR4 メモリ モジュールの合計帯域幅が最大 50 GB/秒であることが深刻なボトルネックとなります。空のコンテキスト ウィンドウを使用すると、専門家混合モデルから 1 秒あたり 1 桁のトークンを取得できます (li

Gemma 4 26B-A4B または Qwen 3.6 35B-A3B)、コンテキスト ウィンドウが大きくなるにつれて速度が急激に低下します。より強力な高密度モデル (Qwen 3.6 27B や Gemma 4 31B など) を試した場合、スループットはトークンごとに秒単位で測定され、対話型作業には使用できないほど遅くなります。そのため、1,500 万ドルの山のハードウェアに LLM を展開した仕事を辞めた後、私は日常業務で企業がホストする AI サービスを、主に Kagi Assistant と OpenRouter 経由で使用してきました。
またその一方で、AI ラボとネオクラウドが RAM と GPU 市場の供給を圧迫しており、LLM 対応のコンピューター (いや、すべてのコンピューター) の価格が高騰しています。これは、人々をコーデックスとクロードのサブスクリプションにさらに依存させるための意図的な戦略ではないかもしれませんが、経済学がその方向を向いていることは確かです。 AI バブルがはじけない場合、不足によりハードウェアの価格が何年も上昇し続けることになります。したがって、自分のコンピュータで強力な LLM を使用して生産的な作業を行いたい場合は、コンピュータを追加購入する時期が来ました。
LLM の進化と陳腐化のペースを考えると、特定のモデルを実行できるサイズのハードウェアを購入するのは愚かなことのように思えますが、私はなんとなく購入しました。 Qwen 3.6 27B は 4 月に登場しました。これは中国のモデルであり、あらゆる癖や偏見が伴います。しかし、LiveBench を信じるのであれば、Qwen 3.6 27B は、いくつかのカテゴリのタスクにわたって、はるかに大規模なモデルよりも優れたパフォーマンスを発揮します。 DeepSeek V3.2 よりも優れています!このベンチマークで明らかにより高性能なオープンウェイト モデルを見つけるには、DeepSeek V4 Flash を使用する必要があります。DeepSeek V4 Flash には 10 倍の VRAM (ウェイト精度を一定に保つ) が必要ですが、その全体的なスコアはわずかに優れているだけです。他の汚染されていないプライベートベンチマークも、このモデルに良い結果を示しています。また、/r/localllama コミュニティは Qwen 3.6 27B の使用を非常に気に入っていると言っても過言ではないと思います。エージェント ハーネスでは正常に動作しますが、不正な形式のツールでは

私の経験ではすべてが稀です。
Qwen 3.6 27B は高密度モデルであるため、全体的なサイズが同様の専門家混合モデルと比較して、実行するのに大量の計算が必要になります。しかし、高密度モデルでは、一定量の VRAM に対して最大限の「インテリジェンス」が得られます。また、VRAM は、最近の「手頃な価格の」GPU ハードウェアの最も制限された側面です。このモデルは、重みあたり最大 4 ビットまでの量子化を非常によく許容するため、VRAM との相性がさらに良くなります。
私はすでに、このモデルのクラウド ホスト バージョンを活用して良い経験を積んできました。しかし、代替手段が必要な場合は、Google の Gemma 4 31B がベンチマークでほぼ同じくらい優れており、Google は Gemma 3 を悩ませていた不快な企業トレーニング マニュアルの性格を修正しました。
私は通常、高速で愚かなトークンよりも低速で賢いトークンを好みます。エージェントの動作を読んで理解することによって進歩が制限されることが多く、「より多くのコード」が欠点になることが多いからです。しかし、同様のサイズでより高速なモデルが必要な場合は、Qwen 3.6 35B-A3B または Gemma 4 26B-A4B を使用できます。どちらも専門家がまばらに混在しており、トークンあたりの計算量を減らすために「インテリジェンス」ペナルティを受け入れます。
したがって、Qwen 3.6 27B に適合するコンピューターは、米国 + 中国起源の多様性と疎 + 密なアーキテクチャを備えた優れた代替品にも適合します。そして、将来を見据えると、これらはこのサイズ層で実行できる最悪のモデルになります。 AI ラボが (過去数年間のように) より優れたオープンウェイト モデルをリリースし続ければ、時間が経つにつれて、同じコンピューターでより優れた LLM を実行できるようになります。
「Mac を買えばいい」という声が聞こえます。準備はできていました！私の買い物リストは次のとおりです。
積極的な量子化を行わずに有用なモデルに適合するのに十分な VRAM、なんとか Qwen 3.6 27B (約 4 bpw)、MTP、ビジョン サポート、深いコンテキスト ウィンドウ。
最高品質の 4 ビット GGUF ( UD-Q4_K_XL ) は 1

7.9 GB なので、24 GB の VRAM では実際には十分ではありませんが、おそらく 32 GB は十分でしょう。
快適な読み取り速度 (1 秒あたり 20 ～ 30+ トークン) でテキストを生成するのに十分なメモリ帯域幅。
高速なコンテキストの事前入力に十分なコンピューティング
エージェント プログラミングでは特に重要です。エージェントはたくさんの本を読みます。
単一のコンピューターまたは GPU。PCIe またはネットワーク リンク全体でモデルの重み (テンソルまたはパイプラインの並列処理) を分割するための、混乱したトレードオフとむらのあるバックエンド サポートを回避します。
(RTX 3090 には高速な NVLink がありますが、廃止されました。)
LLM に対するかなり成熟したエコシステムのサポート。つまり、ハードウェアを完全かつ効率的に使用して必要なモデルを実行するソフトウェアがすでに存在します。
現在の市場はこんな感じです。
(このチャートを変更したい場合は、ここにコードとデータがあります。はい、チャートはスロップマシンで作成しました。)
この記事の執筆時点では、チャート上のすべての製品 (RTX 3090 を除く) は実際に、保証付きで新品を購入できる価格で、大まかではない販売店に在庫があります。 (つまり、販売終了となった 128 GB 以上の Mac Studio や 64 GB Mac Mini はありません。) また、不快なノイズ レベルや 240 ボルトの回路を使用したり、別個のファンやシュラウドを取り付ける必要もなく、これらすべてを家の中で使用できます。 (つまり、MI50 や V100 のような古いデータセンター GPU はありません。)
ハイスペック Mac は、大量の VRAM と優れたメモリ帯域幅を備えていますが、最新のプロシューマー向け GPU と比較するとコンピューティング能力が劣ります。 (対数の X 軸に注意してください!) エージェントのワークロードは、ストリーミングを開始するための応答を待機している間に、コンピューティングに制約されるプレフィルなど、多くの読み取り作業を実行します。これは、すべての VRAM が必要な場合にのみ Mac を購入する必要があることを意味します。私の毎日のモデルが最大 27B パラメータの場合は、Mac を購入する必要はありません。
現在の価格では、DGX Spark は Strix Halo よりもはるかにお買い得です。しかし、それでも、プロシューマー向けの GPU や Mac と比べると帯域幅が非常に制約されているため、

VRAM はそれほど必要ないのでスキップしてください。
NVIDIA は 4,000 ドルで、大量の VRAM または大量のメモリ帯域幅を販売しますが、両方を入手するには 12,500 ドルを費やす必要があります。
でも、決めたときにはこのチャートはありませんでした。私の決断プロセスは次のとおりです。llama.cpp と vLLM で十分なサポートがあり、最も安価な新しい 32 GB 以上の GPU を購入するというものでした。これはほとんど Intel Arc Pro B70 でしたが、AMD のより成熟したソフトウェア エコシステムにより、Radeon AI Pro R9700 を選択することになりました。私は 6 月 6 日に ASRock バージョンに 1349 ドルを支払いました。そしてほら、私のAIスロップマシン。
はい、それは外部 GPU ドックです。私はデスクトップ コンピューターを持っていませんが、予備の古い Thinkpad を持っています。RAM に制約のあるこの時代に、新しいデスクトップを構築するよりも、T480 用の eGPU ドックを購入した方がはるかに安かったのです。 T480 の Thunderbolt 3 ポートは、PCIe 3.0 の 2 レーンのみをサポートします。これは 2018 年の基準から見ても遅かったですが、全体的なパフォーマンスに大きな影響を与えるとは予想していません。 Thunderbolt リンク経由でモデルの重みを読み込むにはさらに時間がかかりますが、起動後は、帯域幅を大量に消費する処理はすべて GPU に対してローカルに残ります。 Thunderbolt 経由でプロンプト トークンとレスポンス トークンを渡しますが、それは小さなデータです。推論ソフトウェアがプロンプト キャッシュをシステム RAM にスワップ アウトできるようにすると、エージェントが複数の長いプロンプトを交互に実行する場合、それらの数 GB の転送にはそれぞれ数秒かかりますが、それでも空のプロンプト キャッシュを事前に入力する場合と比べて大幅な速度向上が得られます。
はい、これはポータブルではありません。搭載されたドックとラップトップの重量は合わせて約 20 ポンドになります。でも、サーバーとして設定したので大丈夫です。別のラップトップ、ソファ、通りのカフェ、携帯電話のホットスポットなどから SSH トンネル経由で API 呼び出しを行っています。このようなコンピューターに安全なリモート ネットワーク アクセスを取得するには、いくつかの方法があります。

さまざまなレベルの使いやすさと堅牢性を備えています。 LLM を別のサーバーで実行する利点の 1 つは、手元にあるラップトップが冷却され、静かで、バッテリーが長持ちすることです。もちろん、欠点は、両端で信頼性の高いインターネット接続が必要なことです。テキストと時折画像をやり取りするだけなので、高速接続である必要はありません。
私のエージェント ハーネスは、他の推論プロバイダーから得られるものと同様に、スロップ マシンを OpenAI 互換 API として認識します。物理サーバーの場所、形状、サイズはクライアントから完全に抽象化されます。
とにかく、これはほとんどうまく機能します。箱を開けてから数時間以内に、llama.cpp 経由で Qwen 3.6 27B を GPU のフル使用率を反映した速度で実行できるようになりました。
GPU と LLM の組み合わせを考えると、それを実行する方法の可能性は驚くほど大きくなります。
推論エンジン (LLM をロードし、GPU での作業を管理するコンピューター プログラム) には、トレードオフを伴ういくつかの選択肢があります。たとえば、llama.cpp は一度に 1 つのリクエストに対してより高速に実行されますが、vLLM はより高速な結合生成速度で多くの並列リクエストを処理する代わりに、リクエストごとの生成はより遅くなります。
これらのプログラムには、ハードウェアやワークロードに合わせてパフォーマンスを一方向または別の方向に押し上げるための多くの構成ノブがあります。最近の例の 1 つは、マルチトークン予測を有効にして生成速度を 50% 以上高速化できますが、その代わりにプレフィルが徐々に遅くなり、コンテキスト ウィンドウで使用できる VRAM が若干少なくなります。
一部のハードウェアでは、複数の GPU バックエンドを選択できます。 ROCmかバルカンか？わかりませんが、両方試してどちらがより効果的かを確認してください。
GPU+LLM のあまり一般的ではない組み合わせには、さまざまなバグやパフォーマンスの低下があります。 vLLM は、新しいハードウェアやニッチなハードウェア/モデルを扱うのに特に注意が必要です。 Y

最終的には、さまざまなバグ レポートやプル リクエストから適用する修正を収集したり、必要なパッチのセットが組み込まれた他の人のフォークを見つけたりすることになります。
私はこの領域をある程度体系的に探索し、ほとんどの作業をエージェントに任せたかったのですが、それはとても勉強になり、とても楽しかったです。私はこれらすべてを、Pi コーディング エージェントと「共同作業」するワークスペースである https://codeberg.org/cmart/slop-machine/ で実現しました。研究上の疑問や試してみるべきことについて話し合い、それを TODO.md に入れます。エージェントは todos に取り組み、その結果を lab-notebook.md に保存します。私はそこで見つかった内容をレビューし、フォローアップ作業を指示します。時々、研究ノートにも書きますが、各セクションにはそのノートの作成者が記載されています。これらのファイルは、今後のエージェント セッションで何が起こっているかを迅速に理解するのに役立ちます。すべてをコンテキスト ウィンドウに読み込んで、次のタスクを選択できます。
エージェントの性格は良くも悪くもアデロール XR です。何かを試し続けたり、問題の解決策を求めて Web を閲覧したり、デバッグ スクリプトを作成したり、別のことをしている間に別のパラメータ スイープやアブレーションを設定したりするための無制限のスタミナがあります。また、重要ではない細部に過度に集中しやすく、明らかな間違いに気づかず、認識論的認識が乏しい（他の人が「幻覚症」と呼ぶもの）こともあります。

[切り捨てられた]

## Original Extract

I, a human, write every word of this blog in my own voice. Nonetheless, I built an AI Slop Machine.
( Boring! Skip to the pretty chart )
American and Chinese AI labs continue to make increasingly strong open-weights large language models (Qwen, Gemma, gpt-oss, and others). You can self-host these LLMs on consumer-attainable computers and GPUs. They’re pretty good at agentic knowledge work tasks, and when you use them, several of the things that people say are bad about AI don’t apply.
A tech company never handles your prompt or response data. It stays on your own computer.
You aren’t stuck on a ride of subscription-with-mystery-quota versus premium API pricing. Once you buy the hardware, it’s free to download the models and use them forever.
The ways that you access self-hosted LLMs are generally free of advertising or manipulation of your attention.
Once you have a copy of the model, its maker cannot withdraw your access to it.
You aren’t causing additional demand for AI data centers.
Yes, these models are trained at large computational expense, but the cake is done baking by the time you download a copy.
The electricity you’ll use is at most a few hundred watts while it’s actively working, similar to a video game console.
So, I’ve watched my friends with higher-spec Macbooks running increasingly-capable AI at usable speeds, bringing chunky power bricks to the cafe and using agents on hot, whirring laptops.
Meanwhile, my daily laptop (a 13th-gen Intel Framework 13) is not good at this. As a party trick, it fits decent LLMs in its 48 GB of system RAM, but the ~50 GB/s combined bandwidth of two DDR4 memory modules is a severe bottleneck. With an empty context window, I get single-digits of tokens per second from mixture-of-experts models (like Gemma 4 26B-A4B or Qwen 3.6 35B-A3B), and it slows down sharply as the context window grows. If I try a stronger dense model (like Qwen 3.6 27B or Gemma 4 31B), I measure throughput in seconds per token , unusably slow for interactive work. So, after leaving a job at which I deployed LLMs on a $15m pile of hardware , I’ve been using AI services hosted by companies for day-to-day tasks, mostly via Kagi Assistant and OpenRouter.
Also meanwhile, AI labs and neoclouds are supply-squeezing the RAM and GPU markets, making LLM-capable computers (nay, all computers) more expensive. This may not be an intentional strategy to make people more dependent on Codex and Claude subscriptions, but the economics sure point in that direction. If the AI bubble doesn’t pop then the shortage will continue raising hardware prices for years. So, if I wanted to do productive things with strong LLMs on my own computer, it was time to buy more computer.
Given the pace of LLM evolution and obsolescence, it seems stupid to buy hardware sized to run one specific model, but I sorta did. Qwen 3.6 27B came out in April. It’s a Chinese model, with all the quirks and biases that entails. But if you believe LiveBench , Qwen 3.6 27B performs better than much larger models across several categories of task. Better than DeepSeek V3.2! To find an unambiguously more capable open-weights model on this benchmark, you need to reach for DeepSeek V4 Flash, which requires 10 times more VRAM (holding weight precision constant), and its overall score is only a little better. Other unpolluted, private benchmarks also show good things for this model. Also, I think it’s fair to say that the /r/localllama community really likes working with Qwen 3.6 27B. It behaves well in an agent harness, and malformed tool calls are rare in my experience.
Being a dense model, Qwen 3.6 27B is computationally intensive to run compared to mixture-of-experts models of similar overall size. But dense models get you the most ‘intelligence’ for a given amount of VRAM, and VRAM is the most limited aspect of ‘affordable’ GPU hardware these days. This model tolerates quantization down to ~4 bits per weight quite well , which makes it even more VRAM-friendly.
I already had good experiences harnessing cloud-hosted versions of this model. But if I want an alternative, Gemma 4 31B from Google is almost as good on benchmarks, and Google fixed the obnoxious corporate training manual personality that plagued Gemma 3.
I usually prefer slower-smarter tokens over faster-stupider tokens, because progress is often limited by me reading and understanding what the agent is doing, and ‘more code’ is often a liability. But if you want a faster model at similar size, you can use Qwen 3.6 35B-A3B or Gemma 4 26B-A4B, both sparse mixture-of-experts which accept an ‘intelligence’ penalty to use less computation per token.
So, any computer that will fit Qwen 3.6 27B will also fit good alternatives, with diversity of American + Chinese origin and sparse + dense architecture. And looking to the future, these are the worst models that we’ll ever be able to run at this size tier. If AI labs keep releasing better open-weights models (as they have for the past few years), the same computer will be able to run better LLMs as time goes on.
I hear you saying “just get a Mac”. I was prepared to! My shopping list was:
Enough VRAM to fit useful models without aggressive quantization blah blah Qwen 3.6 27B at ~4 bpw, MTP, vision support, deep-ish context window.
The highest-quality 4-bit GGUF ( UD-Q4_K_XL ) is 17.9 GB, so 24 GB of VRAM isn’t really enough but 32 GB probably is .
Enough memory bandwidth to generate text at comfy reading speed (20-30+ tokens per second).
Enough compute for fast context prefill
Especially important for agentic programming. Agents do a lot of reading.
Single computer or GPU, to avoid a mess of tradeoffs and spotty backend support for splitting model weights (tensor or pipline parallelism) across a PCIe or network link.
(RTX 3090 has faster NVLink but it’s discontinued.)
Reasonably mature ecosystem support for LLMs, i.e. there is already software that uses the hardware fully and efficiently to run the models that you want.
Right now, the market looks like this.
(Here is the code and data if you want to modify this chart. And yes, I made the chart with the slop machine.)
At time of writing, everything on the chart (except RTX 3090) is actually in-stock at a non-sketchy seller, for the prices indicated, to buy new with a warranty. (So there’s no discontinued 128+ GB Mac Studio or 64 GB Mac Mini.) Also, you can use all of this stuff in your house without obnoxious noise levels, 240 volt circuits, or needing to attach a separate fan + shroud. (So there’s no old datacenter GPUs like MI50 or V100.)
High-spec Macs have a lot of VRAM and good memory bandwidth, but are weak on compute compared to the latest prosumer GPUs. (Note the logarithmic X-axis!) Agentic workloads do a lot of reading stuff , i.e. prefill, which is compute-bound while you wait for a response to start streaming. This tells me to only get a Mac if I need all that VRAM, which I don’t if my daily model is ~27B parameters.
DGX Spark is a much better value than Strix Halo at current prices! But even so, it’s very bandwidth-constrained compared to prosumer GPUs and Macs, so if you don’t need all that VRAM, skip it.
For $4k, NVIDIA will sell you either a lot of VRAM or a lot of memory bandwidth, but to get both you need to spend $12,500.
But I didn’t have this chart when I decided! My decision process was: buy the cheapest new 32 GB+ GPU that has reasonably good support in llama.cpp and vLLM. This was almost an Intel Arc Pro B70, but AMD’s maturer software ecosystem tipped me to a Radeon AI Pro R9700 . I paid $1349 on June 6 for the ASRock version. And lo, my AI Slop Machine.
Yes, that’s an external GPU dock. I don’t have a desktop computer but I do have a spare old Thinkpad, and it was much cheaper to buy an eGPU dock for the T480 than to build a new desktop in these RAM-constrained times. The T480’s Thunderbolt 3 port only supports two lanes of PCIe 3.0. This was slow even by 2018 standards, but I don’t expect it to affect overall performance much. It takes an extra moment to load the model weights across the Thunderbolt link, but after it’s started up, all the bandwidth-intensive stuff stays local to the GPU. You’re passing prompt and response tokens over Thunderbolt, but that’s tiny data. If you allow your inference software to swap out prompt caches to system RAM, then if your agents are alternating between multiple loong prompts, those several-GB transfers each take a few seconds, but you’re still getting a big speedup compared to prefilling with an empty prompt cache.
Yes, this is not portable. The loaded dock and laptop weigh a combined ~20 pounds. But that’s fine, because I set it up as a server. I’m making API calls to it over an SSH tunnel from a different laptop, on the couch, cafe down the street, hotspot from my phone, etc. There are several ways to get secure remote network access to a computer like this, with varying levels of ease and robustness. One benefit of running your LLM on a separate server is that the laptop under your hands stays cool, quiet, and long-lasting on battery. The drawback, of course, is you need a reliable internet connection at both ends. It doesn’t need to be a fast connection, as you’re only passing text and the occasional image back and forth.
My agent harness sees the slop machine as an OpenAI-compatible API, just like you get from any other inference provider. The location, shape, and size of the physical server are totally abstracted away from the client.
Anyway, this works mostly great. Within a couple hours of opening the box, I had it running Qwen 3.6 27B via llama.cpp at speeds reflecting full GPU usage.
Given a GPU+LLM combination, the possibility space of how you run it is surprisingly large.
There are several choices of inference engine (the computer program that loads the LLM and manages the work on the GPU), with tradeoffs. For example, llama.cpp runs faster for a single request at a time, while vLLM accepts slower per-request generation in exchange for handling many parallel requests at much higher combined generation speed.
These programs have many config knobs to push performance in one direction or another, to fit your hardware or workload better. One recent example, you can enable multi-token prediction to get 50+% faster generation speed, at the cost of incrementally slower prefill and a little less VRAM available for the context window.
For some hardware, you get a choice of multiple GPU backends. ROCm or Vulkan? I dunno, try both and see which works better.
The less-popular combinations of GPU+LLM have various bugs and performance degradations. vLLM is especially finicky to get working on new or niche hardware/models. You end up collecting fixes to apply from various bug reports and pull requests, or maybe find someone else’s fork with the necessary set of patches baked in.
I wanted to explore this space somewhat systematically, and use my agent to do most of the work, which turned out to be learningful and a lot of fun. I made all of this happen in https://codeberg.org/cmart/slop-machine/ , a workspace where I ‘collaborate’ with Pi coding agent. We discuss research questions and things to try, and put them in TODO.md . The agent works on todos and puts its findings in lab-notebook.md , where I review what it found and guide the follow-up work. Sometimes I’ll also write in the lab notebook, and each section indicates who authored it. These files help future agent sessions understand what’s going on quickly. They can read it all into the context window and pick up whatever task is next.
The agent’s personality is Adderall XR, which is good and bad. It has unlimited stamina to keep trying stuff, browse the web for solutions to problems, write debugging scripts, set up another parameter sweep or ablation while I go do something else. It’s also susceptible to hyperfocus on details that don’t matter, fail to notice some obvious mistakes, and have poor epistemic awareness (what others call “hallucinati

[truncated]
