---
source: "https://blog.j11y.io/2026-06-10_hidden-state-probes/"
hn_url: "https://news.ycombinator.com/item?id=48498283"
title: "Don't let the LLM speak, just probe it"
article_title: "Don't let the LLM speak, just probe it. - by James Padolsey"
author: "gmays"
captured_at: "2026-06-12T01:02:02Z"
capture_tool: "hn-digest"
hn_id: 48498283
score: 2
comments: 0
posted_at: "2026-06-12T00:24:48Z"
tags:
  - hacker-news
  - translated
---

# Don't let the LLM speak, just probe it

- HN: [48498283](https://news.ycombinator.com/item?id=48498283)
- Source: [blog.j11y.io](https://blog.j11y.io/2026-06-10_hidden-state-probes/)
- Score: 2
- Comments: 0
- Posted: 2026-06-12T00:24:48Z

## Translation

タイトル: LLM に語らせないで、ただ調査してください
記事のタイトル: LLM に語らせないで、ただ調査してください。 - ジェームズ・パドルシー著

記事本文:
LLM に発言させず、ただ調べてください。 - ジェームズ・パドルシー著
ジェームズ・パドルシーのブログ
LLM に発言させず、ただ調べてください。
TL;DR: LLM が「ここにテキストがあり、ここに基準があります。それは満たしていますか?」と読み取ったとき、多くの場合、答えは単一のトークンを生成する前に非表示の状態ですでに存在しています。したがって、生成を完全にスキップします。最後のプロンプト トークン (モデルのレイヤーの最大 70%) で非表示の状態を取得し、それを小さな MLP に供給し、出力を調整します。トレーニング データによって基準が異なるため、英語で記述できる任意の分類器として機能する 1 つの凍結モデルが得られます。
問題: NOPE での仕事の一環として、大量のテキストについて多くの質問をする必要があります。 「これは何のトピックですか」という質問ではなく、バニラ コサイン距離を使用した埋め込み分類器がこれらの細かい質問を処理しますが、構造的な質問には対応します。それで、トランスクリプトが与えられた場合、私は知りたいのですが、話者自身が苦労しているのですか、それとも他の誰かについて説明しているのですか？これは皮肉ですか？ 「昔は嫌いだったけど今は好き」というのは現在の嫌いを表しているのでしょうか？埋め込みはそのようなことにはほとんど気付かない。彼らは憎しみの言葉と愛の言葉、そして話題を認識します。通常のエスカレーションは LLM の審査員です。ルーブリックを含む大きなモデルにテキストを送信し、散文を返して解析します。裁判官は仕事をしますが、速度は遅く、すべてを審査すると高価になります。裁判官が報告する自信は雰囲気にすぎません。裁判官の「7/10」は確率ではありません。
私が最終的に理解したのは、LLM が次のようなプロンプトを読み取るときです。
<審査対象の内容>
以前はこの商品が嫌いでしたが、今では正直言って大好きです。
</content_to_be_judged>
<判断基準>
筆者は現在この商品を気に入っています。
</judgment_criteria>
基準は満たされましたか？
…何かを生成する前にすでに答えが決まっています (CoT は考慮されていませんが、

この小さな恵みを私に与えてください）。基準とコンテンツの比較はフォワード パス内で行われ、結果は残差ストリーム内のジオメトリとして保存されます。生成 (遅くてコストがかかり、散文を解析する部分) は、すでに下された決定を言葉に翻訳するモデルにすぎません。
したがって、話させないでください。中間的な層 (意味の豊かな表現が存在する傾向がある層) で、最後のトークンの隠れ状態を取得し、その上で 1 つの数値を出力する小さな MLP (またはもっと単純な線形プローブ!) をトレーニングします。それがすべてのトリックです。
どの成分も新しいものではありません。線形プローブは古いものです (Alain & Bengio、2016)。ロジット レンズとその調整された子孫 (Belrose et al., 2023) は、フォワードパスの内部情報を飛行中の決定として読み取ります。少し新しいのは、プローブを一般的なゼロショット分類器として使用することです。つまり、推論時に英語で基準が提供される。公平を期すために言うと、これでも BERT スケールのエンコードでは解決された問題です。 NLI クロスエンコーダを参照してください。 GLIClass – しかし重要なことに、これらは 100,000 以上のトークンにわたる現代の LLM のより深い理解と因果関係/推論に精通することは決してありません。パラメータやコンテキスト サイズがないだけです。
私のレシピ、自分で作りたい場合は:
小さな開いたモデルを考えてみましょう。 IBM の Granite 4.0 micro を使用しています。数十億のパラメータ範囲のものはすべて機能します。 LoRA をトレーニングして研ぎ澄ますことを強くお勧めします。
"Assessment:" のようなシード トークンで終わる上記のようなプロンプト テンプレートを修正します。シードはチャネル ジオメトリのプレフィックスとして設計されており、任意ではありません。
(基準、内容、ラベル (基準が満たされているか)) のトリプルのトレーニング セットを生成します。フロンティア モデルによって生成された数千個で、さまざまな基準をカバーします。ここが重要です。基準はトレーニングによって異なるため、

頭は、特定の基準ではなく、「コンテンツが基準を満たしているかどうか」を読み取ることを学びます。
それらをモデルを通して実行し、シード位置で隠れ状態を収集し、MLP に適合します。
0.7 が実際にこれらの 70% が陽性であることを意味するように出力を調整します (等張回帰)。
最終的に得られるものは奇妙で美しいものです。1 つの凍結モデルと 1 つの小さなヘッドが組み合わされて分類器を形成します。リクエスト時に基準を英語で書き込むと、埋め込み分類子のおおよその費用で、数十ミリ秒以内に調整された確率が返されます。基準ごとのトレーニングは一切ありません。バックボーンの LoRA によってさらに鮮明になりますが、正直なところ、ベース モデルだけでほとんどのことが可能になります。機能はすでにモデルに含まれています。あなたはそれを要求するのではなく、ただ読み上げているだけです。生成モデルの非生成部分を取得しています。
オプションの LoRA についてのメモ。それがどのようにトレーニングされるかがレシピの中で私のお気に入りの部分だからです。何かを分類するためにトレーニングするわけではありません。書くように訓練するのです。トレーニング トリプルごとに、フロンティア モデルにラベルが渡され、一文の判定を書き込みます (評価: 内容が基準を満たしていないため… - 再判定ではなく、既知の回答を正当化します)。 LoRA は、単純なネクスト トークンの損失を利用して、プローブが推論時に表示する正確なプロンプトからそのテキストを生成することを学習します。そして、推論時には何も生成されません。シードで停止して、隠された状態を読み取ります。テキストは足場です。その唯一の仕事は、MLP 用に決定がより読みやすくレイアウトされるように、シード位置でジオメトリの形状を変更することです。モデルに答えを言うように教えて、モデルが話す前に結晶化させてから、決して話させないでください。
シードトークンについて少し余談を述べておきます。なぜなら、私はそれが静かに魅力的だと思うからです。プロンプトの最後のトークンは単なるものではありません

t プロンプトが停止する場所。それは答えが書かれるアドレスです。因果関係による注意は、モデルが計算したすべてのことを、これから話そうとしている位置に集中させます。判断の合図でプロンプトを終了すると、判断はそこに 1 トークン幅で具体化されます。業界全体が、言わずもがなこれに依存しています。RLHF のすべての報酬モデルは、テンプレート化されたプロンプトの最終トークンのスカラー ヘッドです。同じトリックです。
操作するノブは、シード トークン自体 (例: "Assessment: ___" または "Criteriamet? ___" )、基準およびコンテンツに対するその正確な位置、および抽出の正確なレイヤーです。多くの場合、最も価値があるのは最終層ではなく、中間層以降のどこかです。これを解明するには多くの実験が必要であり、基礎となるモデルと問題領域によって異なります。
転送パスのコストがかかる部分は、コンテンツの読み取りです。したがって、20 の基準に対して 1 つのコンテンツをスコア付けしたい場合 (このハンマーを手に入れたら、そうすることができます)、最適化が必要になります。コンテンツを 1 回事前入力し、KV をキャッシュし、キャッシュに対して安価な 30 トークンの継続として各基準を実行します。これをKVポッピングと呼んでいいのでしょうか？それは機能し、スコアは長い道のりを実行した場合と少し同じです。
もちろん、キャッシュされたコンテンツは、モデルが基準を認識する前にエンコードされる場合を除きます。一方、コンテンツの前に基準を置くという遅いルートを選択すると、すべてのコンテンツ トークンがすべての層で基準に対応します。モデルは、質問をすでに知っている状態でコンテンツを読み取ります (ある種の質問には非常に有益です)。ただし、キャッシュされたバージョンではブラインド読み取りとなり、基準は最後に 1 回だけ参照されます。
ほとんどのコンテンツでは、これには費用はかかりません。実際の評価セットでは、2 つのパスは統計的です。

いわゆる同一です。しかし、抽出しようとしている基準が単純ではない、より現実的な長い形式のコンテンツを指す場合 (例: 反事実や複雑な表現が含まれている場合)、問題にはすべてのレイヤーで相互作用するための基準とコンテンツが必要ですが、キャッシュはまさにそれを排除します。
編集: どうやら、これは典型的なクロスエンコーダーと遅延インタラクションのトレードオフ ( ColBERT ) のようです。
この手法は現在、NOPE の安全スタック内で実行している Predicate と呼ばれるものを強化しています。「すべての会話のすべてのメッセージに対して構造的な質問を実行する」ことが仕事の全体であり、コストと遅延の両方で経済学がまったく機能しないと判断されます。しかし、アプローチは一般的であり、その部分は汎用品です。小さなオープン モデル、プロンプト テンプレート、生成された数千のトリプル、MLP、等張回帰です。正直に言うと、午後は配管工事です。もしあなたがそれを構築するなら、どこで壊れるかをぜひ聞きたいです。

## Original Extract

Don't let the LLM speak, just probe it. - by James Padolsey
James Padolsey's Blog
Don't let the LLM speak, just probe it.
TL;DR: When an LLM reads "here's some text, here's a criterion — does it satisfy it?", the answer often already exists in its hidden state before it generates a single token. So skip generation entirely: grab the hidden state at the last prompt token (~70% of the way up the model's layers), feed it to a tiny MLP, calibrate the output. Because the training data varies the criterion, you get one frozen model that acts as any classifier you can write in English.
The problem : As part of my work at NOPE I need to ask lots of questions about lots of text. Not "what topic is this" questions — embedding classifiers with vanilla cosine distances handle those fine — but structural ones. So, given a transcript, I want to know Is the speaker themselves the one struggling, or are they describing someone else? Is this sarcasm? Does "I used to hate this, but now I love it" express current dislike? Embeddings are mostly blind to that sort of thing; they see hate-words and love-words and a topic. The usual escalation is an LLM judge: send the text to a big model with a rubric, get prose back, parse it. Judges work, but they're slow, they're pricey if you're running them on everything, and the confidence they report is vibes — a judge's "7/10" isn't a probability of anything.
The thing I eventually internalized is that when an LLM reads a prompt like this:
<content_to_be_judged>
I used to hate this product, but honestly now I love it.
</content_to_be_judged>
<judgment_criteria>
The writer currently likes the product.
</judgment_criteria>
Criteria met?
…it has already decided the answer before it generates anything (not accounting for CoT but allow me this small grace). The comparison between criterion and content has been done, inside the forward pass, and the result is sitting there as geometry in the residual stream. Generation — the slow, expensive, parse-the-prose part — is just the model translating a decision it has already made into words.
So: don't let it speak . Take the hidden state at that final token, at some middle-ish layer (where rich representations of meaning tend to live), and train a small MLP (or more simple linear probe!) on it that outputs one number. That's the whole trick.
None of the ingredients are new. Linear probes are old ( Alain & Bengio, 2016 ); the logit lens and its tuned descendants ( Belrose et al., 2023 ) read forward-pass internals as in-flight decisions. What's a bit new is using the probe as a general zero-shot classifier. I.e. Having a criterion supplied in English at inference time. To be fair, even that is a solved problem at the BERT-scale of encoding. See NLI Cross Encoders e.g. GLIClass – but crucially these will never reach the deeper understanding and causality/reasoning savvy of modern LLMs across 100k+ tokens. They just don't have the parameter or context size.
My recipe, if you want to do this yourself:
Take a small open model. We use IBM's Granite 4.0 micro; anything in the few-billion-parameter range works. I strongly recommend training a LoRA to sharpen it.
Fix a prompt template like the one above, ending in a seed token like "Assessment:" . The seed is designed as a prefix to channel geometries, it's not arbitrary.
Generate a training set of (criterion, content, label (is criterion satisfied?)) triples — a few thousand, frontier-model-generated, covering lots of different criteria. This is the important bit: because the criteria vary in training, the head learns to read "does the content satisfy the criterion," not any particular criterion.
Run them through the model, collect the hidden state at the seed position, fit the MLP.
Calibrate the outputs (isotonic regression) so that 0.7 actually means seventy-percent-of-these-were-positive.
What you end up with is strange and lovely: one frozen model and one tiny head that together form any classifier. You write the criterion at request time, in English, and get back a calibrated probability in a few tens of milliseconds, for roughly embedding-classifier money. No per-criterion training, ever. A LoRA on the backbone sharpens it, but honestly the base model alone gets you most of the way; the capability is in the model already; you're just reading it out instead of asking for it. You're getting the non -generative part of generative models.
A note on that optional LoRA, because how it's trained is my favorite part of the recipe. You don't train it to classify anything. You train it to write. For each training triple, a frontier model is handed the label and writes a one-sentence verdict ( ASSESSMENT: The content does not satisfy the criterion, because… — justifying a known answer, not re-judging). The LoRA learns, with plain next-token loss, to produce that text from the exact prompt the probe will see at inference. And then at inference none of it is ever generated — we stop at the seed and read the hidden state. The text is scaffolding: its only job is to reshape the geometry at the seed position so the decision is laid out more legibly for the MLP. You teach the model to say the answer so it crystallizes before it speaks ... then never let it speak.
A small aside on the seed token, because I find it quietly fascinating. The last token of your prompt isn't just where the prompt stops; it's the address where the answer gets written . Causal attention funnels everything the model has worked out into the position it's about to speak from; end the prompt with a judgment cue and the judgment crystallizes right there, one token wide. The whole industry already relies on this without saying so: every reward model in RLHF is a scalar head on the final token of a templated prompt. Same trick.
The knobs to play with are the seed token itself (e.g. "Assessment: ___" or "Criteria met? ___" ), its exact position relative to the criterion and content, and the precise layer(s) of extraction. Often it is not the final layer that holds most value, but somewhere after the middle. This needs lots of experimentation to figure out, and depends on the underlying model and problem domain.
The expensive part of a forward pass is reading the content. So if you want to score one piece of content against twenty criteria (and once you have this hammer, you do), there's an optimization begging to be made: prefill the content once, cache the KV, then run each criterion as a cheap thirty-token continuation against the cache. I guess this can be called KV-popping? It works, and the scores are bit-identical to doing it the long way.
Except, of course, the cached content gets encoded before the model has seen any criterion . Whereas if we go the slower route of putting the criterion before the content, then every content token attends to the criterion at every layer. The model reads the content already knowing the question (very beneficial to some sorts of questions). In the cached version however it reads blind, and the criterion only gets one look back at the end.
On most content this costs nothing — on our real-world eval sets the two paths are statistically identical. But if pointed at more realistic longer-form content where the criterion you're trying to extract is not straightforward (e.g. containing counterfactuals or complex phrasing), then the problem needs criterion and content to interact at every layer, and the cache forecloses exactly that.
EDIT: Apparently this is a typical cross-encoder versus late-interaction trade-off ( ColBERT ).
This technique now powers a thing called Predicate that I run inside NOPE's safety stack, where "run a structural question against every message of every conversation" is the whole job and judge economics simply don't work, both on cost and latency. But the approach is general and the parts are commodity: a small open model, a prompt template, a few thousand generated triples, an MLP, isotonic regression. An afternoon of plumbing, honestly. If you build one, I'd love to hear where it breaks for you.
