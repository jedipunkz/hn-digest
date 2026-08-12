---
source: "https://dubir.net/tools/local-llm-hardware-calculator/"
hn_url: "https://news.ycombinator.com/item?id=49274480"
title: "Local LLM Hardware Calc"
article_title: "Which LLM can I run locally? VRAM and GPU calculator · Dubir Group"
author: "delneg"
captured_at: "2026-08-12T16:47:34Z"
capture_tool: "hn-digest"
hn_id: 49274480
score: 1
comments: 0
posted_at: "2026-08-12T15:58:02Z"
tags:
  - hacker-news
  - translated
---

# Local LLM Hardware Calc

- HN: [49274480](https://news.ycombinator.com/item?id=49274480)
- Source: [dubir.net](https://dubir.net/tools/local-llm-hardware-calculator/)
- Score: 1
- Comments: 0
- Posted: 2026-08-12T15:58:02Z

## Translation

タイトル: ローカル LLM ハードウェア計算
記事のタイトル: どの LLM をローカルで実行できますか? VRAM および GPU 計算機 · Dubir Group
説明: GPU、Mac、または VRAM の予算と、モデルに実行させたいこと (チャット、コーディング、PDF の読み取りとスキャン) を選択し、どのオープン モデルがどの量子化で、どの程度の速度で適合するかを確認します。無料、ブラウザ上で動作します。

記事本文:
どの LLM をローカルで実行できますか? VRAM および GPU 計算機 · Dubir Group Dubir Group サービス AI 製品 パートナー お問い合わせ 電話予約 無料ツール · サインアップなし どの LLM をローカルで実行できますか?
ハードウェアと、そのモデルに実行させたいこと (チャット、コーディング、PDF やパスポートの読み取り) を選択します。計算ツールには、適合するオープン モデル、それらが保持するコンテキストの量、おおよその生成速度、および適合しないすべてのモデルについて、適合しない理由がリストされます。
Apple NVIDIA AMD Intel その他の GeForce
専用 GPU はモデルの約 90% を維持します。残りはコンテキストとアクティベーションに進みます。
コミュニティのデフォルトであり、依然としてスイートスポットです。FP16 の困惑度は約 0.6% オフですが、チャットや要約では誰も気づきません。厳密なフォーマットと長い算術演算は最初に劣化します。
モデルが一度に保持する会話やドキュメントの量。すべてのトークンには、重みに加えて KV キャッシュ メモリがかかります。
モデルは 16 ビット数値として出荷されます。量子化では、より少ないビットでそれらを保存するため、同じモデルは半分または 4 分の 1 のメモリを使用し、生成速度はメモリの読み取り速度によって制限されるため、実行速度も速くなります。コストは精度であり、それはほとんどの人が予想するよりも小さいです。Q4_K_M (どこでもデフォルト) では、測定された損失は 1% 未満であり、通常のチャットでは目に見えません。 3 ビットでは数パーセントになり、難しいタスクは不安定になります。 2 ビットでは、モデルは明らかに鈍くなります。 Q5 以上では、テストハーネスがなければ差を測定できません。
買い物をする際に重要なルール: 同じメモリでは、4 ビットのより大きなモデルが、8 ビットのより小さなモデルよりもほぼ常に優れています。
コンテキストの長さはモデルの作業メモリ、つまりモデルが一度に認識する会話やドキュメントの量です。重み (KV キャッシュ) とは別にメモリが必要です。そのため、以下の表では、モデルが 8k コンテキストには適合するが、128k には適合しないと示されることがあります。ランタイム

LM Studio などは、KV キャッシュ自体を量子化することもできます (最初に Flash アテンションをオンにします)。8 ビット キャッシュではメモリが半分になり、基本的にコストはかかりません。
詳細: 損失の数値は実際のモデルでのコミュニティの測定値であり、私たちのものではありません。 Artefact2 の KL ダイバージェンス テーブル、これらの形式を定義した llama.cpp k-quants の作業、および Flash アテンションと KV キャッシュの設定については LM Studio のドキュメントを参照してください。
16 モデル中 13 モデルは、Q4_K_M の RTX 4090 24GB で 8k のコンテキストで実行されます。このデバイスには 24 GB のメモリが搭載されており、そのうち約 22 GB がメモリになります。
Qwen3-VL 30B-A3B (MoE) ビジョンを実行します。目の付いた 30B-A3B レシピ: 262k のコンテキストがコントラクト全体をページ画像として飲み込み、3.3B のアクティブなパラメーターが高速に保ちます。ドキュメントとビジョン、チャット、長いコンテキスト 18 GB ~302 tok/s 44k 実行 Gemma 4 31B ビジョン チャット、コーディング、ドキュメント & ビジョン 19 GB ~32 tok/s 27k 実行 Qwen3 30B-A3B (MoE) 死ぬことを拒否する 2025 年リリース: トークンごとに実行されるパラメーターは 3.3B のみなので、そのサイズよりもはるかに高速です。チャット、コーディング、年中無休のオートメーション 18 GB ~302 tok/s 40k 実行 Qwen3.6 27B ハイブリッド アテンション: 64 レイヤーのうち 48 レイヤーは固定サイズの状態のリニア アテンションであるため、長いコンテキストは低コストのままです。チャット、コーディング、長いコンテキスト 16 GB ~37 tok/s 102k 実行 Gemma 4 26B-A4B (MoE) ビジョン チャット、ドキュメントとビジョン、年中無休の自動化 15 GB ~263 tok/s 177k 実行 gpt-oss 20B (MoE) ~4 ビット (MXFP4) でネイティブに出荷。より高い精度は理論上のものです。チャット、コーディング、年中無休のオートメーション 12 GB ~277 tok/s 128k Gemma 4 12B ビジョンを実行 チャット、ドキュメントとビジョン、ドラフト 7.3 GB ~83 tok/s 237k Qwen3-VL 8B ビジョンを実行 「この請求書を読んで」「この写真について説明する」に対する 1 つのモデルの答え: DocVQA 96、OCRBench 90。 4ビットの8GBカード。ドキュメントとビジョン、チャット 6.0 GB ~115 tok/s 119k 実行 Gemma 4 E4B ビジョン チャット、製図、年中無休の自動化 4.6 GB ~222 tok/s 1

28k 実行 Gemma 4 E2B ビジョン Gemma 4 の小規模なエンド。より大きな兄弟の前での投機的デコード用の標準ドラフト モデルでもあります。製図、年中無休の自動化 2.9 GB ~434 tok/s 128k 実行 DeepSeek-OCR-2 (MoE) ビジョン 光学圧縮を中心に構築された MIT ライセンスの OCR: 高密度のページが少数の視覚トークンになります。これはまさに一括スキャンと複数ページ PDF の場合です。ドキュメントとビジョン、年中無休の自動化 2.2 GB ~1750 tok/s 8k 実行 PaddleOCR-VL 1.6 ビジョン OmniDocBench を上回る 0.9B のドキュメント解析スペシャリスト: テーブル、シール、スキュー スキャン、109 言語。古いオフィスハードウェアで十分です。ドキュメントとビジョン、24 時間年中無休の自動化 0.6 GB ~1108 tok/s 128k gpt-oss 120B を実行 (MoE) ~4 ビット (MXFP4) でネイティブに出荷。より高い精度
[切り捨てられた]
全体の計算は 4 つのアイデアです。気になるものを開いてください。
モデルの重みは単なる数値であり、モデルが必要とするメモリは算術演算です。 16 ビット精度の 8B モデルは、2 バイトの 80 億倍、約 15GB です。 4ビットに量子化すると、同じモデルは約4.5GBになります。一般的な GGUF 形式も正確に 4 ビットや 8 ビットではありません。埋め込み層と出力層がより高い精度に保たれているため、Q4_K_M は重みあたり平均 4.85 ビットに近くなります。これが、私たちの数値が単純な計算よりわずかに上に位置する理由です。
アテンションは会話内のすべてのトークンのキーと値ベクトルを保持するため、新しいトークンごとに過去を再計算しません。そのキャッシュは重みとは別のものであり、コンテキストに応じて線形に増加します。
2 × レイヤー × kv_heads × head_dim × トークンあたり 2 バイト
Qwen3.6 27B の場合、これはトークンあたり 64KB で、その完全な 256k ウィンドウのコストは 16GB で、これは 4 ビットの重みそのものとほぼ同じです。これは、「適合する」モデルをロードした後、20 個のメッセージでメモリ不足エラーが発生するときに人々が忘れる部分です。2026 年リリース攻撃

まさにこれです。Qwen3.6 では、64 レイヤーのうち 16 レイヤーだけが増大するキャッシュを保持します (残りは固定サイズの状態でリニア アテンションを使用します)。Gemma 4 ではほとんどのレイヤーに短いスライディング ウィンドウが与えられ、DeepSeek V4 はキャッシュを完全に圧縮します。私たちの推定ではフル アテンション レイヤーのみがカウントされており、キャッシュを 8 ビットに量子化すると、キャッシュをサポートするランタイムでは再び半分になります。
16 ビットから 4 ビットに移行すると、メモリが約 4 削減され、Q4_K_M で測定された複雑さのコストは 1% 未満になります。実際には、チャット、要約、取得した文書からの回答などは快適に行えます。厳密な JSON、長い算術言語、低リソース言語では最初に損傷が現れ、4 ビット未満では損失が急速に増大します。有益なルール: 同じメモリでは、4 ビットのより大きなモデルが、8 ビットのより小さなモデルにほぼ常に勝ります。 Q4の32BはQ8の14Bよりも優れたマシンです。
1 つのトークンを生成すると、アクティブな重みセット全体がメモリから 1 回読み取られます。そのため、単一ストリームの生成はコンピューティングの問題ではなく、メモリ帯域幅の問題になります。カードの帯域幅を重みのサイズで割ると、上限がわかります。4.5 GB モデルに対する 1008 GB/秒は理論上約 220 トークン/秒で、実際の実行時間はその約 60% に達します。専門家混合モデルについても説明します。 Qwen3 30B-A3B は 30B のパラメータを保存しますが、トークンごとに 3.3B しか読み取らないため、大規模モデルのメモリを使用しながら小規模モデルの速度で実行されます。
プロンプト処理はその逆で、コンピューティング依存型で並列処理です。そのため、長いドキュメントは NVIDIA カードではすぐに取り込まれますが、ラップトップではゆっくりと取り込まれます。また、複数のユーザーに同時にサービスを提供すると、バッチ化されたエンジンがバッチ全体の重みを 1 回読み取るため、算術演算が完全に変更されます。ビジー状態のサーバー上の合計スループットは、上記の単一ストリームの数値の何倍にもなります。
2 つの理由が当てはまります。の

1 つ目はデータです。顧客記録、契約書、または患者ファイルがインフラストラクチャから離れられない場合、ホスト型 API は価格に関係なく選択肢になりません。 2 つ目は量で、これは原理ではなく算術です。安定した予測可能なトークン量により、GPU が迅速に償却されます。爆発的な、小さな音量は決してそうではありません。中間のパス、機密コンピューティング領域内での検証可能なプライベート推論もあり、これについては損益分岐点のページで説明されています。
モデルをローカルで実行、回答済み
おおよそ: 8GB は、Gemma 4 E4B や Qwen3-VL 8B などの 8B クラス モデルをショート コンテキストの 4 ビットで実行します。 12GB には Gemma 4 12B が追加されます。 24GB は、26 ～ 35B クラス (Gemma 4 31B、Qwen3.6、Qwen3 30B-A3B) を 4 ビットで実行するか、より小型のモデルをロング コンテキストで 8 ビットで実行します。
gpt-oss 120B が適合するのは 96GB および 128GB マシンです。ドキュメント スペシャリストははるかに小規模です。PaddleOCR-VL は 2GB 未満の PDF を解析します。
上の計算機にカードを設定すると、正確にリストが表示されます。
コンテキストの前に、4 ビットでの重みの場合は約 40 GB: 24 GB カード 2 枚、48 GB ワークステーション カード、または 64 GB のユニファイド メモリを搭載した Mac 。 8 ビットでは約 70 GB、量子化されていない場合は 140 GB。
購入する前に知っておく価値があります。2026 年のオープン リリースでは、ほとんどが密な 70B 形状を省略し、gpt-oss 120B のような専門家混合モデルを採用しています。これらのモデルは、同様の量のメモリを必要としますが、数倍高速に実行されます。
人々が期待しているよりも少し、そして少ないです。デフォルトの Q4_K_M は、完全な 16 ビット重みよりも複雑度が 1% 未満であり、チャット、要約、または回答の検索では誰もそれを認識しません。最初に劣化するのは、厳密な書式設定、演算の長いチェーン、およびまれな言語の出力です。
実際のダメージは 4 ビット未満から始まります。Q3 では約 2 ～ 3% ですが、Q2 ではモデルは明らかに鈍くなっています。タスクが Q4 で失敗した場合は、より大きなモデルにジャンプする前に、同じモデルの Q5 または Q6 を試してください。
完全なはしご、

各ステップで測定された損失は、上記の量子化ピッカーにあります。
ビジョンモデルであり、これは昨年現地の品質が最も飛躍したクラスです。テキストのみの LLM はページを決して参照しません。ドキュメント ビジョン モデルは、レイアウトとスタンプを含む画像を直接読み取ります。
PaddleOCR-VL 1.6 は 2GB 未満でテーブルと 109 の言語を解析し、DeepSeek-OCR-2 は一括スキャン用に構築され、Qwen3-VL 8B は「この請求書を読んでそれに関する質問に答える」ことをカバーします。計算機の目標として「ドキュメントとスキャン」を選択すると、マシンがそのうちのどれを実行するかが表に表示されます。
運用中のパスポートと ID については、最初に独自のスキャンでテストしてください。グレアとホログラムは、小さなモデルをつまずかせるものです。
はい: LM Studio または Ollama をインストールし、Q4_K_M でモデルをダウンロードして完了です。
Mac が大型モデルの価格を上回る理由は、ユニファイド メモリにあります。128GB Mac は gpt-oss 120B を保持しますが、これは単一のコンシューマ GPU では不可能です。トレードは帯域幅です。 M5 Max は RTX 4090 の 1008 GB/秒に対して約 614 GB/秒で移動するため、高密度モデルの生成は遅くなりますが、2026 年を支配する小規模アクティブ MoE モデルはその差を縮めます。
M1 Pro 以降のものはすべて計算機に入れる価値があります。 macOS はメモリの一部を予約するため、全体の約 75% を割り当てます。
思い出から始めましょう。 50 個のタブを持つブラウザはギガバイトを保持し、数レイヤーでも GPU メモリからシステム RAM に溢れ出すと、モデルの速度のほとんどが失われます。macOS のアクティビティ モニターまたは Windows のタスク マネージャーを確認してください。LM Studio のモデル ページには、GPU に到達したレイヤーの数が正確に表示されます。
次に、無料の勝利を収集します。フラッシュ アテンションをオンにし、小さなドラフト モデル (大きな Gemma の前に Gemma 4 E2B を置くのが標準の組み合わせです) を使用して投機的デコードをセットアップします。これにより、同一の品質で 1.5 ～ 2 倍の高速化が実現します。
そして、最初のトークンの前に長いプロンプトが処理されます

というメッセージが表示されるため、大きなドキュメントの起動が遅くなるのは正常であり、故障ではありません。
これらのモデルを製品化します
Dubir はプライベート AI を構築します。モデルはネットワーク内のハードウェア上で実行され、データがサードパーティ API に到達することはありません。ボックスを購入することが正しい行動でない場合は、機密コンピューティング領域に検証可能なプライベート推論をセットアップすることもできます。いずれにせよ、私たちはそれを仕様化し、展開し、実行し続けます。
電話予約 contact@dubir.net Dubir Group Dubir Group Limited はキプロスの技術グループです。私たちはソフトウェア、AI、ビジネスを実行するシステムを構築します。

## Original Extract

Pick your GPU, Mac or VRAM budget and what you want the model to do (chat, coding, reading PDFs and scans) and see which open models fit, at which quantisation, and how fast. Free, runs in your browser.

Which LLM can I run locally? VRAM and GPU calculator · Dubir Group Dubir Group Services AI Products Partners Contact Book a call Free tool · No sign-up Which LLM can I run locally?
Pick your hardware and what you want the model to do: chat, coding, reading PDFs and passports. The calculator lists the open models that fit, how much context they hold, roughly how fast they generate, and for everything that does not fit, why not.
Apple NVIDIA AMD Intel Other GeForce
A dedicated GPU keeps about 90% for the model; the rest goes to context and activations.
The community default and still the sweet spot: about 0.6% off FP16 perplexity, which nobody notices in chat or summarisation. Strict formatting and long arithmetic degrade first.
How much conversation or document the model holds at once. Every token costs KV-cache memory on top of the weights.
A model ships as 16-bit numbers. Quantisation stores them in fewer bits, so the same model takes a half or a quarter of the memory and runs faster too, because generation speed is limited by how fast memory can be read. The cost is accuracy, and it is smaller than most people expect: at Q4_K_M, the default everywhere, the measured loss is under 1% and invisible in normal chat. Down at 3 bits it is a few percent and hard tasks get shaky; at 2 bits the model is visibly dumber. Above Q5 you cannot measure the difference without a test harness.
The rule that matters when shopping: a bigger model at 4-bit almost always beats a smaller one at 8-bit for the same memory.
Context length is the model's working memory: how much of the conversation or document it sees at once. It costs memory separately from the weights (the KV cache), which is why the table below sometimes says a model fits at 8k context but not at 128k. Runtimes such as LM Studio can also quantise the KV cache itself (turn on Flash Attention first): 8-bit cache halves that memory and costs essentially nothing.
Learn more: the loss figures are community measurements on real models, not ours. See Artefact2's KL-divergence tables , the llama.cpp k-quants work that defined these formats, and the LM Studio docs for the Flash Attention and KV-cache settings.
13 of 16 models run on RTX 4090 24GB at Q4_K_M with 8k of context. That device has 24 GB of memory, of which roughly 22 GB is yours to fill.
Runs Qwen3-VL 30B-A3B (MoE) vision The 30B-A3B recipe with eyes: 262k context swallows a whole contract as page images, and 3.3B active parameters keep it quick. documents & vision, chat, long context 18 GB ~302 tok/s 44k Runs Gemma 4 31B vision chat, coding, documents & vision 19 GB ~32 tok/s 27k Runs Qwen3 30B-A3B (MoE) A 2025 release that refuses to die: only 3.3B parameters run per token, so it is far faster than its size. chat, coding, 24/7 automation 18 GB ~302 tok/s 40k Runs Qwen3.6 27B Hybrid attention: 48 of its 64 layers are linear-attention with a fixed-size state, so long context stays cheap. chat, coding, long context 16 GB ~37 tok/s 102k Runs Gemma 4 26B-A4B (MoE) vision chat, documents & vision, 24/7 automation 15 GB ~263 tok/s 177k Runs gpt-oss 20B (MoE) Ships natively at ~4-bit (MXFP4); higher precisions are theoretical. chat, coding, 24/7 automation 12 GB ~277 tok/s 128k Runs Gemma 4 12B vision chat, documents & vision, drafting 7.3 GB ~83 tok/s 237k Runs Qwen3-VL 8B vision The one-model answer to "read this invoice" and "describe this photo": DocVQA 96, OCRBench 90, and it fits an 8GB card at 4-bit. documents & vision, chat 6.0 GB ~115 tok/s 119k Runs Gemma 4 E4B vision chat, drafting, 24/7 automation 4.6 GB ~222 tok/s 128k Runs Gemma 4 E2B vision The small end of Gemma 4. Also the standard draft model for speculative decoding in front of its bigger siblings. drafting, 24/7 automation 2.9 GB ~434 tok/s 128k Runs DeepSeek-OCR-2 (MoE) vision MIT-licensed OCR built around optical compression: a dense page becomes a handful of visual tokens, which is exactly the bulk-scans and multi-page-PDF case. documents & vision, 24/7 automation 2.2 GB ~1750 tok/s 8k Runs PaddleOCR-VL 1.6 vision A 0.9B document-parsing specialist that tops OmniDocBench: tables, seals, skewed scans, 109 languages. Old office hardware is enough for it. documents & vision, 24/7 automation 0.6 GB ~1108 tok/s 128k Runs gpt-oss 120B (MoE) Ships natively at ~4-bit (MXFP4); higher precisi
[truncated]
The whole calculation is four ideas. Open the ones you care about.
A model's weights are just numbers, and the memory they take is arithmetic. An 8B model at 16-bit precision is 8 billion times 2 bytes, about 15GB. Quantise it to 4-bit and the same model is around 4.5GB. The common GGUF formats are not exactly 4 or 8 bits either: Q4_K_M averages close to 4.85 bits per weight because the embedding and output layers are kept at higher precision, which is why our numbers sit slightly above the naive calculation.
Attention keeps a key and a value vector for every token in the conversation so it does not recompute the past on each new token. That cache is separate from the weights and grows linearly with context:
2 × layers × kv_heads × head_dim × 2 bytes per token
For Qwen3.6 27B that is 64KB per token, and its full 256k window costs 16GB, about as much as the 4-bit weights themselves. This is the part people forget when they load a model that "fits" and then hit an out-of-memory error twenty messages in. The 2026 releases attack exactly this: in Qwen3.6 only 16 of the 64 layers keep a growing cache (the rest use linear attention with a fixed-size state), Gemma 4 gives most layers a short sliding window, and DeepSeek V4 compresses the cache outright. Our estimate counts only the full-attention layers, and quantising the cache to 8-bit halves it again on runtimes that support it.
Going from 16-bit to 4-bit cuts memory by roughly four and costs under 1% in measured perplexity at Q4_K_M. In practice, chat, summarisation and answering from retrieved documents survive it comfortably. Strict JSON, long arithmetic and low-resource languages are where the damage shows up first, and the loss compounds fast below 4 bits. The useful rule: a bigger model at 4-bit almost always beats a smaller one at 8-bit for the same memory. A 32B at Q4 is a better machine than a 14B at Q8.
Generating one token reads the entire active weight set out of memory once. That makes single-stream generation a memory-bandwidth problem, not a compute problem. Divide the card's bandwidth by the size of the weights and you have the ceiling: 1008 GB/s against a 4.5GB model is around 220 tokens/s in theory, and a real runtime reaches roughly 60% of that. It also explains mixture-of-experts models. Qwen3 30B-A3B stores 30B parameters but only reads 3.3B per token, so it runs at the speed of a small model while taking the memory of a large one.
Prompt processing is the opposite: it is compute-bound and parallel, which is why a long document is ingested quickly on an NVIDIA card and slowly on a laptop. And serving several users at once changes the arithmetic entirely, because a batched engine reads the weights once for the whole batch. Total throughput on a busy server is many times the single-stream number above.
Two reasons hold up. The first is data: if customer records, contracts or patient files cannot leave your infrastructure, a hosted API is not an option regardless of price. The second is volume, and that one is arithmetic rather than principle. Steady, predictable token volume amortises a GPU quickly; bursty, low volume never does. There is also a middle path, verifiably private inference inside confidential-computing enclaves, which the breakeven page explains.
Running models locally, answered
Roughly: 8GB runs an 8B-class model like Gemma 4 E4B or Qwen3-VL 8B at 4-bit with short context. 12GB adds Gemma 4 12B. 24GB runs the 26-35B class (Gemma 4 31B, Qwen3.6, Qwen3 30B-A3B) at 4-bit, or a smaller model at 8-bit with long context.
The 96GB and 128GB machines are where gpt-oss 120B fits. Document specialists are far smaller: PaddleOCR-VL parses PDFs in under 2GB.
Set your card in the calculator above and it will list them exactly.
About 40GB for the weights at 4-bit, before context: two 24GB cards, a 48GB workstation card, or a Mac with 64GB of unified memory. At 8-bit around 70GB, unquantised 140GB.
Worth knowing before you buy for one: the 2026 open releases mostly skipped the dense-70B shape in favour of mixture-of-experts models like gpt-oss 120B, which take a similar amount of memory but run several times faster.
A bit, and less than people expect. Q4_K_M, the default, measures under 1% worse perplexity than the full 16-bit weights, and nobody spots that in chat, summarisation or retrieval answers. What degrades first is strict formatting, long chains of arithmetic and rare-language output.
Real damage starts below 4 bits: around 2-3% at Q3, and at Q2 the model is visibly dumber. If a task fails at Q4, try Q5 or Q6 of the same model before jumping to a bigger one.
The full ladder, with the measured loss at each step, is in the quantisation picker above.
A vision model, and this is the class where local quality jumped furthest in the last year. A text-only LLM never sees the page; a document vision model reads the image directly, layout and stamps included.
PaddleOCR-VL 1.6 parses tables and 109 languages in under 2GB, DeepSeek-OCR-2 is built for bulk scans, and Qwen3-VL 8B covers "read this invoice and answer questions about it". Pick "Documents & scans" as the goal in the calculator and the table shows which of them your machine runs.
For passports and IDs in production, test on your own scans first; glare and holograms are what trip the small models.
Yes: install LM Studio or Ollama , download a model at Q4_K_M, done.
Unified memory is the reason Macs punch above their price for large models: a 128GB Mac holds gpt-oss 120B, which no single consumer GPU can. The trade is bandwidth. An M5 Max moves about 614 GB/s against an RTX 4090's 1008 GB/s, so a dense model generates slower, though the small-active MoE models that dominate 2026 narrow the gap.
Anything from an M1 Pro up is worth putting into the calculator; macOS reserves part of the memory, so budget about 75% of the total.
Start with the memory. A browser with fifty tabs holds gigabytes, and a model that spills even a few layers out of GPU memory into system RAM loses most of its speed: check Activity Monitor on macOS or Task Manager on Windows, and LM Studio 's model page shows exactly how many layers made it onto the GPU.
Then collect the free wins: turn on Flash Attention, and set up speculative decoding with a small draft model (Gemma 4 E2B in front of a bigger Gemma is the standard pairing), which is a 1.5-2x speedup at identical quality.
And a long prompt is processed before the first token appears, so a slow start on a big document is normal, not a fault.
We put these models into production
Dubir builds private AI: the model runs on your hardware, inside your network, and your data never reaches a third-party API. If buying a box is not the right move, we also set up verifiably private inference in confidential-computing enclaves. Either way, we spec it, deploy it and keep it running.
Book a call contact@dubir.net Dubir Group Dubir Group Limited is a Cyprus technology group. We build software, AI and the systems businesses run on.
