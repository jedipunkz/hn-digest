---
source: "https://www.seangoedecke.com/ai-inference-is-obviously-profitable/"
hn_url: "https://news.ycombinator.com/item?id=48686412"
title: "AI inference is obviously profitable"
article_title: "AI inference is obviously profitable"
author: "medbar"
captured_at: "2026-06-26T13:39:33Z"
capture_tool: "hn-digest"
hn_id: 48686412
score: 2
comments: 1
posted_at: "2026-06-26T13:29:09Z"
tags:
  - hacker-news
  - translated
---

# AI inference is obviously profitable

- HN: [48686412](https://news.ycombinator.com/item?id=48686412)
- Source: [www.seangoedecke.com](https://www.seangoedecke.com/ai-inference-is-obviously-profitable/)
- Score: 2
- Comments: 1
- Posted: 2026-06-26T13:29:09Z

## Translation

タイトル: AI 推論は明らかに有益です

記事本文:
AI 推論は明らかに利益をもたらします ショーン・ゴーデッケ
AI 推論は明らかに利益をもたらします
多くの人は、AI 推論は役に立たないため、将来何らかの AI モデルが世界経済を支配するようになるだろうと信じている投資家からの莫大な資金で補助金を出さなければならないと主張しています。その愚かなお金が​​なくなると、AI 製品も消えてしまいます。この見解によれば、LLM は本質的に (お金、電力、水の点で) 消費者製品に使用するにはあまりにも高価です。実際、それらは現在、コストを外部化することによってのみ使用できます。お金はVCファンドと現在は個人ETF投資家に、電力は電力会社の消費者に、そして水はデータセンターが建設されているコミュニティに与えられます。
AI を嫌うのには十分な理由がありますが、これは実際にはその 1 つではありません。実際、AI 推論は明らかに利益をもたらします。
計算してみると、推論が有益であることがわかります
フロンティア AI プロバイダーは推論で 70% ～ 80% の粗利益を報告していますが、おそらくそれらを信頼することはできません。実際のコストを大まかに見積もってみましょう。
Nvidia A100 は全負荷時に 400W の電力を消費します。実際には、慎重に調整された推論サーバーであっても常にフル負荷になるわけではありませんが、少なくとも上限にはなります。密度の高い 70B モデル 1 を実行しているとします。これは、1 時間あたり約 200 万トークンで 4 台の A100 に快適に (量子化されていない) 適合します。産業用電力の価格に換算すると、米国では約 13 セント/時になります。 (悲観的に) 冷却コストが同じであると仮定します。これは、100 万出力トークンあたり約 13 セントに相当します 2 。
GPU は最も高価な部分となるため、GPU のコストを償却しましょう。 A100の価格は約2万ドルです。各 A100 が約 5 年間持続する場合 3 、設備投資を回収するには年間 16,000 ドルの利益を得る必要があります (または 1 時間あたり 1.80 ドル)。下の方で

使用率を回復するには時間がかかりますが、GPU の寿命も長くなります。いずれにせよ、全体的な推論コストは 100 万トークンあたり約 1 ドルになります。
GPT-5.4-mini の料金は 100 万トークンあたり 4.50 ドルで、より強力な OpenAI または Anthropic モデルは 3 ～ 6 倍の価格になります。 OpenAI モデルや Anthropic モデルの規模がわからないため、直接比較することは困難ですが、主張されている 70% または 80% の利益率は非常にもっともらしいです。
オープン LLM は、推論が有益であることを実証します
私の見積もりも信用できない場合はどうしますか?中国のオープンウェイト LLM の価格を見てみましょう。 DeepSeek は、DeepSeek-R1 の推論で 80% をわずかに超える利益率を主張しています。 R1 の API 価格は OpenAI や Anthropic 4 の半分以下であるため、推論コストに関する上記の見積もりは高すぎる可能性があることを示唆しています。大規模な冷却はおそらく電力よりも安価であり、R1 には高密度 70B モデルの半分のアクティブ パラメータしかなく、最新の GPU は A100 より効率的であり、推論には大幅なスケール メリットがあります。
DeepSeek のモデルは誰でもダウンロードできるため、大きな利益率を引き出すことはできません。他の推論プロバイダーの 1 つは、同じモデルでそれらを下回る可能性があります。市場における DeepSeek-V4-Pro の推論コストは、100 万出力トークンあたり約 87 セントであり、おそらくモデルを提供する実際のコストにかなり近いと思われます。
AI ラボの場合、推論はトレーニングを補助する必要があります
これらはすべて、OpenAI や Anthropic が利益を上げていることを意味するものではありません。これらの企業は、うまくいくかどうかわからない巨額の設備投資を行っており、新しいモデルをトレーニングしてユーザーを維持するために人材とコンピューティングに巨額の資金を費やしています。
彼らは月額制を提供するなどクレイジーなことをしています

ほぼ無制限の推論を行うための n 番目のサブスクリプション モデルですが、ほぼ確実に収益性がありません。 Claude Code で Anthropic サブスクリプションの代わりに API トークンを使用した場合、10 倍のコストを支払うことになります。ただし、API ベースの Claude Code が役に立たないという意味ではありません。莫大な利益率を差し引くと、月あたりのサブスクリプションよりも安価になるため、すでに DeepSeek の推論 API をエージェント コーディングに使用している人もいます。
OpenAI や Anthropic はなぜ価格を下げないのでしょうか?おそらく OpenAI がそれを考えたと思われますが、AI ラボの場合、推論にはトレーニング費用を補助する必要があります。 OpenAI のような企業は、既存モデルの推論マージンから (少なくとも部分的に) 新しいモデルの制作に資金を提供する必要があります。推論のマージンが非常に高いのはこのためです。AI 研究所は、訓練軍拡競争で生き残るために、あらゆるお金を絞り出そうとしています。
ただし、推論は AI ラボのトレーニング費用を補助するだけで済みます。単なる推論プロバイダーである場合は、トレーニングをまったく行う必要はありません。したがって、OpenAI と Anthropic が廃業したとしても、フロンティア モデルの権利を手に入れた人は、Opus と GPT 推論を利益を上げて販売し続けることができます 5 。 AI バブルの崩壊は、AI 推論が明らかに利益をもたらすため、推論ビジネスの終焉を意味するものではありません。
高価なフロンティア モデルはおそらく専門家が混在しており、密度が高くなく、推定が困難です。ただし、70B の高密度モデルと 70B のアクティブ パラメータを持つ MoE は、規模的には基本的に同じ数値になると思います (ただし、MoE はより多くの GPU メモリを必要とするため、初期費用が高くなります)。フロンティアモデルのパラメータは70Bくらいでしょうか？ AI ラボ以外の誰も実際のところは知りませんが、私の推測では、70B はおそらく 1 Ha よりも大きいのではないかと思います。

育/ミニクラスモデル。
出力トークンは推論を提供する中で最もコストがかかる部分であるため、出力トークンのみのコストを見積もるのが合理的だと思います。入力トークンは 2 つの理由により安価です。1 つは、トランスフォーマーを使用すると、入力トークンを並列で事前入力できること、そして、ほとんどの実際の使用例では、積極的に KV キャッシュにキャッシュできることです。
GPU の寿命を 3 年と見積もることは一般的 (そして間違い) です。これについては、「AI GPU の寿命はおそらく 3 年以上」でたくさん書きました。
繰り返しになりますが、OpenAI または Anthropic モデルのサイズが R1 と同等であるかどうかは不明なので、これは単なる推測です。
Anthropic 社が廃業した場合、他の人がモデルにアクセスできないようにできるかどうかは疑問です。 Anthropic は現在、Broadcom、Google、および多数のプライベートエクイティ会社に対して借金を抱えています。ダリオの抗議を乗り越えて、彼らはミトスとオーパスの重みを獲得するだろうか？
この投稿を気に入っていただけた場合は、私の新しい投稿に関する更新情報を電子メールで購読するか、 Hacker News で共有することを検討してください。
これは、この投稿とタグを共有する関連投稿のプレビューです。
AI GPUs probably live longer than three years
現在の AI の使用が持続不可能であると考える人は、推論 GPU が負荷をかけた状態で「長くても 3 年」しか持たないという主張に頼ることがよくあります。ここでの考え方は、AI バブルの資金が枯渇すると、現在のインフラストラクチャは急速に時代遅れになり、最新の GPU をすべて購入するのに十分な資金が浮かばなくなるということです。したがって、推論コストは現在の AI 製品にとって経済的に意味をなさないほど急速に高くなりすぎます。
続きを読む...
購読する │ About │ ポッドキャスト │ 人気 │ タグ │ RSS

## Original Extract

AI inference is obviously profitable sean goedecke
AI inference is obviously profitable
Many people claim that AI inference is unprofitable to serve, and thus must be subsidized by an ocean of dumb money from investors who believe that some future AI model will come to dominate the world economy. When that dumb money goes away, so will AI products. According to this view, LLMs are just inherently too expensive (in terms of money, power, and water) to be used in consumer products. In fact, they can only be used today by externalizing the costs: money onto VC funds and now retail ETF investors , power onto electric utility consumers , and water onto the communities where datacenters are built.
There are good reasons to dislike AI, but this really isn’t one of them. In fact, AI inference is obviously profitable .
Doing the math demonstrates that inference is profitable
Frontier AI providers are reporting 70%-80% gross margins on inference, but maybe we can’t trust them. Let’s do some very rough estimates on the actual cost.
A Nvidia A100 consumes 400W of power under full load. In practice, even a carefully-tuned inference server will not be at full load all the time, but it’s at least an upper bound. Suppose you’re running a dense 70B model 1 , which will fit comfortably (unquantized) on four A100s at around 2M tokens per hour. At industrial power prices, that’s about 13c/hr in the USA . Suppose (pessimistically) cooling is the same cost. That’s about 13 cents per million output tokens 2 .
Let’s amortize the cost of the GPUs, since that’s going to be the most expensive part. An A100 costs about $20k. If each A100 lasts around five years 3 , you’ll have to make 16k/yr in profit to recoup your capital investment (or $1.80 per hour). At lower utilization, it’ll take longer to recoup, but your GPUs will also last longer. Either way, your overall inference costs are at about one dollar per million tokens.
GPT-5.4-mini charges $4.50 per million tokens, and stronger OpenAI or Anthropic models are three to six times as expensive. It’s hard to make a direct comparison because we don’t know the size of OpenAI or Anthropic models, but the claimed 70% or 80% profit margin is extremely plausible.
Open LLMs demonstrate that inference is profitable
What if you don’t trust my estimates either? Let’s look at the pricing of open-weights Chinese LLMs. DeepSeek have claimed a bit over 80% profit margin on inference for DeepSeek-R1. Since their API pricing for R1 is less than half that of OpenAI or Anthropic 4 , that suggests that my estimates above for inference cost might be too expensive. Cooling at scale is probably cheaper than power, R1 only has half the active parameters of a dense 70B model, modern GPUs are more efficient than the A100, and there are significant economies of scale in inference.
Since DeepSeek’s models are available for anyone to download, they can’t get away with extracting a large profit margin. One of the other inference providers would undercut them with the same model. Inference costs for DeepSeek-V4-Pro on the market are around 87 cents per million output tokens, which is probably pretty close to the actual cost of serving the model.
For AI labs, inference must subsidize training
All of this doesn’t mean that OpenAI or Anthropic are profitable. Those companies are making huge capital investments that may or may not pan out, and are spending enormous amounts of money on talent and compute to train brand-new models and retain users.
They’re doing crazy things like offering per-month subscription models for nearly unlimited inference, which is almost certainly not profitable. If you used an API token instead of your Anthropic subscription in Claude Code, you’d pay ten times the cost. But that doesn’t mean API-based Claude Code couldn’t be a good deal. Some people are already using DeepSeek’s inference API for agentic coding, because once you take away the huge profit margin it’s cheaper than the relative per-month subscription.
Why won’t OpenAI or Anthropic lower their prices? Supposedly OpenAI has thought about it , but for an AI lab, inference has to subsidize training costs . A company like OpenAI has to fund the production of new models from the inference margins on existing models (at least partially). That’s why the margins on inference are so high: the AI labs are trying to squeeze out every dollar so they can stay alive in the training arms race.
However, inference only has to subsidize training costs for an AI lab . If you’re merely an inference provider, you don’t have to do any training at all. Therefore, even if OpenAI and Anthropic go out of business, whoever snaps up the rights to their frontier models will be able to continue selling Opus and GPT inference at a profit 5 . The AI bubble popping will not mean the end of the inference business, because AI inference is obviously profitable .
Expensive frontier models are probably mixture-of-experts, not dense, which is tougher to estimate. However, I think a 70B dense model and a MoE with 70B active params will come out to basically the same numbers at scale (though the MoE will require more GPU memory and thus a greater upfront cost). Are frontier models around 70B params? Nobody outside the AI labs really knows, but my guess is that 70B is probably larger than a Haiku/mini class model.
I think it’s reasonable to estimate the cost of output tokens only, since they’re by far the most expensive part of serving inference. Input tokens are cheaper for two reasons: transformers let you prefill them in parallel, and for most real-world use cases they can be aggressively cached in the KV cache.
It’s common (and wrong) to estimate GPU lifespan at three years. I wrote a lot about this in AI GPUs probably live longer than three years .
Again, this is just an guess, since we don’t know what OpenAI or Anthropic model is equivalent in size to R1.
I do wonder if Anthropic would be able to prevent other people from being able to access the model if the company goes out of business. Anthropic is currently in debt to Broadcom, Google, and a bunch of private equity firms. Would they get the Mythos and Opus weights, over Dario’s protestations?
If you liked this post, consider subscribing to email updates about my new posts, or sharing it on Hacker News .
Here's a preview of a related post that shares tags with this one.
AI GPUs probably live longer than three years
People who think current AI use is unsustainable often rely on the claim that inference GPUs only last “three years at the most” under load. The idea here is that once the AI bubble money drains away, current infrastructure will rapidly become obsolete, and there won’t be enough money floating around to buy a whole slate of brand-new GPUs. Inference costs would thus rapidly become way too expensive for current AI products to make any financial sense.
Continue reading...
subscribe │ about │ podcasts │ popular │ tags │ rss
