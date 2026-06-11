---
source: "https://decrypt.co/370449/xiaomi-mimo-ultraspeed-ai-model-faster-chatgpt-claude"
hn_url: "https://news.ycombinator.com/item?id=48486007"
title: "China's Xiaomi MiMo Is Now 15X Faster Than ChatGPT and Claude"
article_title: "China's Xiaomi MiMo Is Now 15X Faster Than ChatGPT and Claude - Decrypt"
author: "gmays"
captured_at: "2026-06-11T04:35:06Z"
capture_tool: "hn-digest"
hn_id: 48486007
score: 2
comments: 0
posted_at: "2026-06-11T03:52:14Z"
tags:
  - hacker-news
  - translated
---

# China's Xiaomi MiMo Is Now 15X Faster Than ChatGPT and Claude

- HN: [48486007](https://news.ycombinator.com/item?id=48486007)
- Source: [decrypt.co](https://decrypt.co/370449/xiaomi-mimo-ultraspeed-ai-model-faster-chatgpt-claude)
- Score: 2
- Comments: 0
- Posted: 2026-06-11T03:52:14Z

## Translation

タイトル: 中国の Xiaomi MiMo は ChatGPT や Claude よりも 15 倍高速になりました
記事のタイトル: 中国の Xiaomi MiMo は ChatGPT や Claude よりも 15 倍高速になりました - 復号化
説明: Xiaomi の MiMo-V2.5-Pro-UltraSpeed は、カスタム シリコン企業が何年もかけて開発した通常の GPU 上での速度のしきい値を超えています。

記事本文:
中国の Xiaomi MiMo は ChatGPT および Claude より 15 倍高速になりました - Decrypt 中国の Xiaomi MiMo は ChatGPT および Claude より 15 倍高速になりました
Decrypt News 人工知能による価格データ 中国の Xiaomi MiMo は ChatGPT や Claude よりも 15 倍高速になりました
Xiaomi の MiMo-V2.5-Pro-UltraSpeed は、カスタム シリコン企業が何年もかけて構築してきた、通常の GPU 上での速度のしきい値を超えています。
2026 年 6 月 8 日 2026 年 6 月 8 日 4 分で読めます 人工知能。画像: Decrypt/Shutterstock 記事を保存するためのアカウントを作成します。 Google に追加 Decrypt を優先ソースとして追加すると、Google でさらに記事が表示されます。簡単に言うと
Xiaomi と推論パートナーの TileRT は、カスタム チップではなく、標準の 8 GPU コモディティ ノードを使用して、1 兆パラメータ モデルで 1 秒あたり 1,000 トークンを突破しました。これは、この規模では初のことです。
この速度は、モデルのエキスパート レイヤーでの FP4 量子化と、一度に 1 つずつではなく 1 回のパスでトークンの完全なブロックを提案する DFlash 投機的デコードによって実現されます。
限定的な API トライアルが 6 月 9 日から 6 月 23 日まで開始され、価格は標準 MiMo レートの 3 倍、生成速度は約 10 倍になります。
Xiaomi は中国の携帯電話ブランドとしてほとんどの人が知っています。安い電動スクーターや空気清浄機を作っている会社です。月曜日の朝に AI 推論速度の主要記録を破ると期待されるような会社ではありません。
それなのに。 Xiaomi は、MiMo-V2.5-Pro-UltraSpeed をリリースしたばかりです。これは、1 秒あたり 1,000 トークンを超えるトークンをヒットする、兆パラメータのフラッグシップ向けのサービング モードであり、デモではピーク時に 1,200 トークン近くに達します。
パラメーターは、モデルがどのように考えるかを定義する内部数値重みです。パラメーターが多ければ多いほど、モデルが認識できるパターンがより複雑になります。トークンは、モデルが読み書きするテキストの塊であり、それぞれ平均して単語の約 4 分の 3 です。
Xiaomi は単一の 8 GPU 製品でそれを実現しました。

オード。標準ハードウェア、カスタムチップなし。これにより、誰が実際にこの種の速度を本番環境に導入できるかという計算が変わります。
この数字を人間の言葉で表現すると、 Artificial Analysis によると、ほとんどの ChatGPT ユーザーが実際に話している GPT-5.5 は 68 に位置します。Claude Opus 4.6 は 71 程度で、下位モデルの Haiku は 1 秒あたり 98 トークンに達します。 Gemini Flash は 1 秒あたり 192 トークンをヒットします。 MiMo-V2.5-Pro-UltraSpeed は、コーディング ベンチマークで Opus と一致するモデルで 1,000 を実行します。
Cerebras と Groq は、この問題を中心にビジネス全体を構築しました。 Cerebras は、GPU 推論の速度を低下させる帯域幅のボトルネックを排除するために、ディナー皿ほどのサイズのウェハスケールのチップを設計し、44 GB のオンチップ メモリを搭載しました。 Meta の Llama 3.1 405B では 1 秒あたり 969 トークンに達しました。これは驚異的ですが、これは 4,050 億パラメータのモデルであり、MiMo-V2.5-Pro の半分以下のサイズです。 Groq のカスタム言語処理ユニット アーキテクチャは、モデルに応じて 1 秒あたり約 300 ～ 750 トークンを最高に達します。
どちらも、今夜 AWS からレンタルできるハードウェア上では実行されません。
Xiaomi は、モデルレベルのトリックと TileRT と呼ばれる専用の推論エンジンを組み合わせたソフトウェアのみを使用して、これをコモディティ GPU 上で実行しました。
ボンネットの下で実際に何が起こっているのか
2 つのテクニックがスピードを実現します。最初の技術は FP4 量子化と呼ばれます。Xiaomi は、完全な 8 ビットまたは 16 ビットの数値精度でモデルを実行する代わりに、1 兆のパラメーターの大部分を構成するエキスパート レイヤーを 4 ビットまで縮小します。メモリ使用量が減少し、帯域幅の圧力が低下し、速度が向上します。通常、問題となるのはわずかな品質の劣化です。 Xiaomi の修正は外科的です。エキスパート レイヤーのみが圧縮され、他のすべては完全な精度を維持します。このアプローチでは、品質の損失はほぼゼロと言われています。
2 つ目は DFlash 投機的デコードです

。通常の投機的デコードでは、小規模なドラフト モデルが次のいくつかのトークンを推測し、その後、大きなモデルがそれらを並行して検証します。 DFlash は連続的なドラフトを完全にスキップします。マスクされた位置のブロック全体を 1 回の前方パスで埋めます。コーディング タスクでは、大規模モデルは検証ラウンドごとに 8 つの提案されたトークンのうち平均 6.3 を受け入れます。これは、1 つのステップで 1 つではなく 6 つのトークンが確認されたことになります。
TileRT はそれを結び付けます。これにより、計算パイプライン全体が GPU 内に継続的に常駐し、オペレーターごとの起動オーバーヘッドや実行ギャップがなくなります。
Xiaomi はこのアプローチを「エクストリーム モデル システム コードデザイン」と呼んでいます。この表現は正確です。どちらの手法だけでも 1 秒あたり 1,000 トークンに到達することはできませんが、すべてのアプローチの相乗効果により 1,000 トークンに到達します。
MiMo-V2.5-Proはフロンティアレベルのモデルです。 4 月にリリースされた V2.5 Pro については、ほとんどのコーディング ベンチマークで Claude Opus に匹敵し、100 万トークンあたり約 0.43 ドルのインプット / 0.87 ドルのアウトプットで実行されます。 Opus のコストは、100 万トークンあたり入力 5 ドル、出力 25 ドルです。
UltraSpeed は、機能を簡略化したバージョンではなく、まさに MiMo V2.5 Pro モデルを高速化します。
推論が十分に高速であれば、モデルの使用方法が変わります。 1 つの答えを待つ代わりに、数十の推論パスを並行して実行できます。不正行為の検出、取引シグナルの生成、リアルタイムのエージェント ループ - これらすべてには、1 秒あたり 60 トークンでは満たせない厳しい遅延制約があります。 1 秒あたり 1,000 トークンであれば、それが可能です。
Xiaomi は、標準の MiMo-V2.5-Pro レートの 3 倍の速度と、約 10 倍の出力を価格設定しています。 API トライアルは 6 月 9 日から 23 日までアプリケーション ベースで実施され、企業およびプロの開発者が優先されます。 FP4-DFlash チェックポイントは、コミュニティ テスト用に Hugging Face ですでにオープンソース化されています。
Web3 の世界への入り口
パートナー ニュースの詳細 ユニバーシティ コインのビデオ

News Explorer チームについて 開示情報 マニフェスト 利用規約 行動規範 プライバシー ポリシー お問い合わせ 採用情報 求人情報 ニュースレターを購読する 最新のニュース、記事、リソースが毎週あなたの受信箱に送信されます。
© 次世代メディア企業。 2026年 デクリプトメディア株式会社

## Original Extract

Xiaomi's MiMo-V2.5-Pro-UltraSpeed blows past the speed threshold custom silicon companies spent years building toward—on regular GPUs.

China's Xiaomi MiMo Is Now 15X Faster Than ChatGPT and Claude - Decrypt China's Xiaomi MiMo Is Now 15X Faster Than ChatGPT and Claude
Price data by Decrypt News Artificial Intelligence China's Xiaomi MiMo Is Now 15X Faster Than ChatGPT and Claude
Xiaomi's MiMo-V2.5-Pro-UltraSpeed blows past the speed threshold custom silicon companies spent years building toward—on regular GPUs.
Jun 8, 2026 Jun 8, 2026 4 min read Artificial intelligence. Image: Decrypt/Shutterstock Create an account to save your articles. Add on Google Add Decrypt as your preferred source to see more of our stories on Google. In brief
Xiaomi and inference partner TileRT have broken 1,000 tokens per second on a 1-trillion-parameter model, a first at that scale, using a standard 8-GPU commodity node—not custom chips.
The speed comes from FP4 quantization on the model's expert layers and DFlash speculative decoding, which proposes a full block of tokens in one pass instead of one at a time.
A limited API trial opens June 9 through June 23, priced at 3× standard MiMo rates for roughly 10× the generation speed.
Most people know Xiaomi as the Chinese phone brand. The one that makes cheap electric scooters and air purifiers. Not exactly the company you'd expect to break a major AI inference speed record on a Monday morning.
And yet. Xiaomi just released MiMo-V2.5-Pro-UltraSpeed , a serving mode for its trillion-parameter flagship that hits over 1,000 tokens per second—peaking near 1,200 in demos.
Parameters are the internal numerical weights that define how a model thinks—the more you have, the more complex the patterns it can recognize. Tokens are the chunks of text the model reads and writes, roughly three-quarters of a word each on average.
Xiaomi did it on a single 8-GPU commodity node. Standard hardware, no custom chips. That changes the calculus for who can actually deploy this kind of speed in production.
To put that number in human terms: per Artificial Analysis , GPT-5.5—what most ChatGPT users are actually talking to—sits at 68. Claude Opus 4.6 lands around 71 with the lower end model, Haiku, touching 98 tokens per second. Gemini Flash hits 192 tokens per second. MiMo-V2.5-Pro-UltraSpeed does 1,000, on a model that matches Opus on coding benchmarks.
Cerebras and Groq built entire businesses around this problem. Cerebras designed a wafer-scale chip the size of a dinner plate, packing 44GB of on-chip memory to eliminate the bandwidth bottleneck that slows down GPU inference. It hit 969 tokens per second on Meta's Llama 3.1 405B—impressive, but that's a 405-billion-parameter model, less than half the size of MiMo-V2.5-Pro. Groq's custom Language Processing Unit architecture tops out around 300–750 tokens per second depending on model.
Neither runs on hardware you can rent from AWS tonight.
Xiaomi did it on commodity GPUs through software alone—a combination of model-level tricks and a purpose-built inference engine called TileRT.
What's actually going on under the hood
Two techniques carry the speed. The first technique is called FP4 Quantization: instead of running the model at full 8-bit or 16-bit numerical precision, Xiaomi shrinks the expert layers—which make up most of the 1 trillion parameters—down to 4-bit. Memory footprint drops, bandwidth pressure drops, speed goes up. The catch is usually a small quality degradation. Xiaomi's fix is surgical: only the expert layers get compressed, everything else stays at full precision. With this approach, quality loss is described as near-zero.
The second is DFlash speculative decoding. Normal speculative decoding has a small draft model guess the next few tokens, then the big model verifies them in parallel. DFlash skips the sequential drafting entirely—it fills a whole block of masked positions in a single forward pass. In coding tasks, the big model accepts an average of 6.3 out of 8 proposed tokens per verification round. That's six tokens confirmed in one step instead of one.
TileRT ties it together. It keeps the entire compute pipeline continuously resident inside the GPU—no per-operator launch overhead, no execution gaps.
Xiaomi calls this approach "extreme model-system codesign," and the phrase is accurate: Neither technique alone gets to 1,000 tokens per second, but the synergy among all approaches does.
MiMo-V2.5-Pro is a frontier-level model. We covered the V2.5 Pro launch in April—it matches Claude Opus on most coding benchmarks and runs at roughly $0.43 input / $0.87 output per million tokens. Opus costs $5 input / $25 output per million tokens.
UltraSpeed accelerates that exact MiMo V2.5 Pro model, not a stripped-down version.
Fast enough inference changes how you can use a model. You can run dozens of reasoning paths in parallel instead of waiting on one answer. Fraud detection, trading signal generation, real-time agent loops—all of these have hard latency constraints that 60 tokens per second can't meet. At 1,000 tokens per second, they can.
Xiaomi is pricing the speed at 3 times the standard MiMo-V2.5-Pro rate for roughly 10 times the output. The API trial runs June 9–23, application-based, with priority given to enterprise and professional developers. The FP4-DFlash checkpoint is already open-sourced on Hugging Face for community testing.
Your gateway into the world of Web3
Partner News Deep Dives University Coins Videos News Explorer About Team Disclosures Manifesto Terms of Service Code of Conduct Privacy Policy Contact Careers Jobs SUBSCRIBE TO OUR NEWSLETTER The latest news, articles, and resources, sent to your inbox weekly.
© A next-generation media company. 2026 Decrypt Media, Inc.
