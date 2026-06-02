---
source: "https://passo.uno/fine-tuning-docs-llm/"
hn_url: "https://news.ycombinator.com/item?id=48354674"
title: "Fine-tuning an LLM to write docs like it's 1995"
article_title: "Fine-tuning an LLM to write docs like it's 1995 – Fabrizio Ferri Benedetti"
author: "theletterf"
captured_at: "2026-06-02T00:42:57Z"
capture_tool: "hn-digest"
hn_id: 48354674
score: 1
comments: 0
posted_at: "2026-06-01T09:48:28Z"
tags:
  - hacker-news
  - translated
---

# Fine-tuning an LLM to write docs like it's 1995

- HN: [48354674](https://news.ycombinator.com/item?id=48354674)
- Source: [passo.uno](https://passo.uno/fine-tuning-docs-llm/)
- Score: 1
- Comments: 0
- Posted: 2026-06-01T09:48:28Z

## Translation

タイトル: 1995 年のようにドキュメントを書けるように LLM を微調整する
記事のタイトル: LLM を微調整して 1995 年のようにドキュメントを書く – Fabrizio Ferri Benedetti
説明: 2030 年の予測で、技術ライターは強力なハードウェア上でローカルに実行される特殊な LLM を使用するだろうと書きました。エンジニアリングの専門家の間では「ローカルファースト」への移行のヒントが見られますが、コネクテッドフロンティアモデルがより強力であることもあり、まだそこには達していません。それはそうです
[切り捨てられた]

記事本文:
LLM を微調整して 1995 年のようにドキュメントを書く – Fabrizio Ferri Benedetti
passo.uno
について
投稿
会談
GitHub
リンクトイン
レトロ
LLM を微調整して 1995 年のようにドキュメントを書く
2030 年の予測で、テックライターは強力なハードウェア上でローカルに実行される特殊な LLM を使用するだろうと書きました。エンジニアリングの専門家の間では「ローカルファースト」への移行のヒントが見られますが、コネクテッドフロンティアモデルがより強力であることもあり、まだそこには達していません。ただし、実験できないわけではありません。それはまさに私が先週行ったことで、80 年代から 90 年代のソフトウェア テクニカル ライターのように書けるように指示モデルを微調整しようとしていました。
研究のために古い技術文書の伝承を呼び出す
90 年代のテクニカル ライターのように書けるように個人的なローカル モデルをトレーニングするには、大量の文書ソースが必要です。たとえば、私のようにモデルを微調整して文章を書きたい場合、このブログでは十分ではありません。この投稿の時点では 100,000 ワードしかないからです。徹底的なトレーニングにはさらに多くのサンプルが必要ですが、それらを入手するのは簡単ではなく、作成するのも簡単ではありません。唯一の簡単な方法は、既存のコーパスを使用することです。どこで入手できますか?
Bitsavers をご紹介します。これは、古いコンピュータのマニュアルやパンフレットを収集してスキャンする Web サイトです。これはコンピューターの歴史と古代の技術文書の非常に貴重なリポジトリであり、どこにでもミラーが用意されています。私は 90 年代の Microsoft マニュアルが好きなので、トレーニング資料のソースとして Microsoft コレクションを選択しました。このコレクションには、1977 年から 2005 年の間に発行された絶版ドキュメントが含まれています。3,700 万語以上で、古いシステムと SDK がカバーされています。
OCR 処理されたテキスト ファイルをダウンロードし、古き良き Python スクリプトを使用してコンテンツからアーティファクトや乱雑なもの (インデックスやフロントマターなど) を取り除きました。私はそれからあなた

OpenRouter を介して安価で高速なモデル、 gemma-4-26b を導入し、各段落をそのわかりやすさに基づいて「キープ」または「ドロップ」に分類しました。この 2 回目のパスの料金は約 8 ドルです。ただし、この 2 パスのクリーニングを行っても、トレーニング データにはノイズが残っていました。ノイズは後で発見されましたが、テストではほとんど問題ありませんでした。
クロードのアドバイスに従って、サニタイズされたテキストを段落とセクションの境界でトレーニング サンプルに分割し、見出しで区切ってコード ブロック全体を保持し、各チャンクの上限を約 512 トークンにしました。各チャンクは、テンプレートから抽出された合成命令とペアになっていました。最終的に、JSONL 形式で 192,456 個のサンプルが作成されました (1 行に 1 つの JSON オブジェクト)。小さなモデルを使用して、より適切な指示や質問を考えることもできましたが、私はせっかちな性格です。
ゼロからトレーニングする代わりに微調整を行う
理想的な世界では、数百万ドルが転がっており、自分の LLM、 Fabrice を作成して燃やす準備ができているでしょう。私は裕福とは程遠いので（そうでなければこんなことは書かないでしょう）、ファブリスに代わる方法は微調整です。これには、生成される各トークンがトレーニング資料によって条件付けされるようにモデルの「重み」を微調整することが含まれます。私は微調整を、タグボートを使って巨大な氷山の軌道をわずかに操縦するようなイメージを描くのが好きです。意図した効果を得るためにほんの少しだけ。
たとえば、検索拡張生成 (RAG) ではなく、なぜ微調整を行うのでしょうか?なぜなら、この実験では、RAG が優れているシナリオであるファクトの取得にはあまり興味がなく、コンテキストの知識が何であれ、LLM に特定のスタイルで動作させたり記述させたりすることに興味があったからです。完全なトレーニングと比較して、微調整には大量のデータが必要ないため、コストが安くなります。また、理由は次のとおりです。

これはテクニックであり、それがどの程度実現可能か見てみましょう。
かなり古いグラフィック カードを搭載したコンピューターでモデルの微調整に数日から数週間を費やすことを避けるために、事前に構成された GPU とツールを備えたオンデマンド ポッドを (比較的) 低価格で提供する AI 開発者向けのオンライン サービスである Runpod に依存しました。たとえば、1 時間あたり 6 ドル未満で、Nvidia B200 (メモリ 192GB) という超強力なカードをリースできます。このサービスには、構成可能な自動リチャージおよびコスト制御メカニズムを備えた便利な API が備わっています。
謎の流行語があふれる世界へ
モデルを微調整することに決めた後、私はそれを達成するための最も賢明な方法についてクロードと相談しました。私たちは QLoRA (量子化低ランク適応) に落ち着きました。これは、LLM の各重みを変更するのではなく、LLM を「フリーズ」し、モデルの動作を再形成する小さなファイル (マスクのようなもの) であるアダプターをその上に置くことで微調整を実現します。 QLoRA の Q は、結果が量子化、つまり圧縮され、メモリ要件が削減されることを意味します。
まだ一緒にいるの？良い。これが密だと思うなら、それは実際にそうなのです。
最近、自宅で LLM を使って何かを行うことは、時間を犠牲にするか、お金を費やすか、または野心的な目標を抑制するかのいずれかです。週末以内に意味のあるものを得るためにバランスをとろうとしました。私は、Llama 3.1 8B Instruct と Qwen 2.5 7B Instruct の 2 つのモデルで微調整を試すことにしました。そのサイズ (約 8B) で、Macbook Air 上で快適に動作します。また、Llama の基本モデル (質問に答えるように訓練されていない) もテストしました。
私はいくつかの異なる条件下で微調整をテストしました。つまり、トレーニング資料の量 (サブセットと完全なコーパス)、エポック数 (トレーニング ラウンド)、およびランクなどの構造パラメーターを変化させます。表面的な知識しか持っていない

このすべての危機に瀕していましたが、私はエージェントが正しい選択をすることを信頼し、あらゆる段階で喜んで質問しました。たとえば、3 エポックでは場合によっては「過剰適合」が発生する可能性があります。 LLM の世界では、これは過剰なトレーニングに相当します。楽しい時間。
アダプターは、微調整したターゲット モデルにのみ適用できます。各アダプターをトレーニングした後、アダプターをラップトップにエクスポートし、GGUF LoRA ファイルに変換して量子化し、ベンチマーク目的でラップトップで実行できるローカル Ollama モデルとして登録しました。ローカル変換アプローチは高速で GPU を必要としませんが、推論は完全にマージされたモデルよりも多少遅くなります。今回のテストでは、速度はそれほど気にしませんでした。
すべての条件に合わせてアダプターをトレーニングするには、おそらく休憩を含めて丸 1 日かかり、合計コストは 50 ドルかかりました。途中で、アダプターを 2 つ紛失しました。Runpod は予算に容赦がなく、資金がゼロになるとすぐにポッドを削除します (確かに教訓は得られました)。クロードは、各実行のセットアップと Runpod の API のフォローアップを担当しました。 Claude Code の /goal コマンドは、各フェーズをループするのに非常に役立ちました (振り返ってみると、YOLO モードで実行していたでしょう)。
この表は、比較したすべてのモデルとその条件を示しています。
微調整後にスタイルは反映されましたか?
各モデルに同じプロンプトを適用しました。
トレーニング資料でよく知られている可能性のある、定番の C 関数である malloc() について説明します。
架空の ConnectWifi() Win32 API 関数を文書化します。トレーニング資料には存在しません。
REST API とは何かを 1990 年代の Microsoft スタイル (時代錯誤的なテスト) で説明します。
すべての質問と回答は、この gist で確認できます。
malloc() テストの場合、未変更のモデルは README 形式の最新の Markdown ドキュメントを生成しましたが、微調整されたモデルは p

あらすじブロック、戻り値セクションなどを含む正しい構造。架空の ConnectWifi() 関数については、3 エポック モデルだけがフィクションを維持し、あたかも現実であるかのように文書化しましたが、他のモデルは第 4 の壁を破って内部知識を固守し、トレーニングに抵抗しました。
REST API の演習も非常に興味深いものでした。Llama Instruct 40k は失敗し、当たり障りのないマーケティング散文が生成されました。クロードさんは、これはラマがフレンドリーでアクセスしやすいようにするための重強化トレーニング（RLHF）のおかげだと考えています。 Qwen の微調整により、HTTP メソッド名を動詞や形式的な見出しとして使用して、ピリオド構造のドキュメントが作成され、レジスターの保持が大幅に向上しました。 Qwen 192k が最も強力で、Windows 2000 リソース キットの章のように始まりました。
繰り返しますが、1990 年代のドキュメントに基づいてトレーニングされ、2000 年代のコンセプトに基づいてテストされた 7B モデルは、本物の時代資料と見間違う可能性のある説得力のある章の冒頭を作成しました。スタイルが転送されました。おお。一方、基本モデルは質問に答えるように訓練されておらず、テキストをオートコンプリートするように訓練されているが、惨めに失敗し、ほぼランダムに生のコーパス、つまり数百行のゴミが噴出した。基本モデルには、「この質問に答える」または「これを完了する」という概念がありません。
実験は、ランク 8 と 16 の間で変化する 1 エポックで、Qwen モデル間のランクの効果を比較することで終了しました。私の理解が正しければ、ランク 8 は、各アダプター マトリックスが 8 つの独立したパターンしか記述できないことを意味します。調整するダイヤルが 8 つあるようなものです。ダイヤルの数が非常に少ないため、アダプターはあまり賢くありません。トレーニング データ内で最も強く、最も繰り返されるパターンに完全にコミットする必要があります。理論的には、ランク 16 はより表現力豊かで繊細です。
ランクの比較は、より小さいアダプターは自由度が少なく、フィクションにコミットすることを示しています。

大きいものよりも軽い。ランク 16 のアダプターは、コーパスをより簡単に「エスケープ」できます。また、1 エポックだけを 16 という中程度のランクと組み合わせると、幻覚がより頻繁になることが判明しました。アダプターは、関連する概念に到達するのに十分な表現力を持っていますが、プロンプトが言おうとしている内容を定着させるほど強化されていません。ランクとエポックは相互作用しているようです。サウンド ミキサーを使用しているようなものです。興味深いことに、アダプターが安いほど、偽装はより誠実になります。
微調整されたモデルは説得力のあるものまねを可能にしますが、代替品ではありません
微調整されたモデルは、90 年代後半の Microsoft テック ライターの優れた模倣者でした。コーパスは、新しい概念を説明するモデルの能力をほとんど保持しながら、モデルのスタイルと音声、および一部の知識を印象付けました。これは比較的安価なプロセスであり、スタイルのレビューや社内スタイル ガイドに従った新しいドキュメントの草稿などのタスクを目的とした効果的な小規模モデルを作成できます。
ただし、そこに到達するのは簡単ではありません。モデルの微調整には低コストではありますが、大量の高品質のトレーニング データが必要ですが、これを作成するのは簡単ではありません。たとえそれを手に入れたとしても、意味があり、追加のトレーニングを受け入れることができる基礎となるモデルを選択する必要があります。さらに、複数のパラメーターを自由に使用できるため、微調整されたモデルをスイート スポットに到達させる作業は時間のかかる作業となります。
安心できる点は、そのようなモデルは決して人間のテックライターに取って代わることはできず、人間を増強するだけであるということです。微調整されたモデルは、調整されていない兄弟モデルと同様に判断力が欠如しており、十分なステアリングが必要です。ファブリスは待たなければなりません。

## Original Extract

In my predictions for 2030 I wrote that tech writers would be using specialized LLMs, running locally on powerful hardware. I see hints of this move to “local first” among engineering pundits, but we’re not there yet, in part because of how much more powerful connected frontier models are. That does
[truncated]

Fine-tuning an LLM to write docs like it's 1995 – Fabrizio Ferri Benedetti
passo.uno
About
Posts
Talks
GitHub
LinkedIn
Retro
Fine-tuning an LLM to write docs like it's 1995
In my predictions for 2030 I wrote that tech writers would be using specialized LLMs, running locally on powerful hardware. I see hints of this move to “local first” among engineering pundits, but we’re not there yet, in part because of how much more powerful connected frontier models are. That doesn’t mean we can’t experiment, though. That’s precisely what I did last week, trying to fine-tune an instruct model to write like a software technical writer from the 80s and 90s.
Summoning old tech writing lore for research
To train a personal, local model to write like a technical writer from the 90s, one needs tons of written sources. If I wanted to fine-tune a model to write like myself, for example, this blog would not be enough, as it’s barely 100k words at the time of this post. You would need more samples for thorough training, and those are not easy to come by, nor simple to produce. The only quick way is to use an existing corpus. Where could I get one?
Meet Bitsavers : it’s a website that collects and scans old computer manuals and brochures. It’s an incredibly valuable repository of computer history and ancient tech writing, with mirrors available everywhere. As I’m fond of Microsoft manuals from the 90s, I chose the Microsoft collection as the source of training materials. The collection contains out-of-print docs published between 1977 and 2005: more than 37 million words, covering old systems and SDKs.
I downloaded the OCR’d text files and cleaned the content from artifacts and clutter (like indices and frontmatter) using good old Python scripts. I then used a cheap and fast model through OpenRouter , gemma-4-26b, to classify each paragraph as either “keep” or “drop” based on its intelligibility. This second pass cost around 8 dollars. Even with this two-pass cleaning, though, training data retained noise that I discovered only later, but that was largely OK for my tests.
I split the sanitized text into training examples on paragraph and section boundaries, breaking at headings and keeping code blocks whole, with each chunk capped at around 512 tokens as per Claude advice. Each chunk was paired with a synthetic instruction drawn from templates. I ended up with 192,456 examples in JSONL format (one JSON object per line). I could have used a small model to also come up with better instructions and questions, but I’m an impatient person.
Fine-tuning as an alternative to training from scratch
In an ideal world, I would have several millions of dollars lying around, ready to be burned creating my own LLM, Fabrice . Since I’m far from rich (I wouldn’t be writing this otherwise), the alternative to Fabrice is fine-tuning, which involves tweaking the “weights” of a model so that each token generated is conditioned by the training materials. I like to picture fine-tuning as slightly steering the trajectory of a massive iceberg using tugs; just a little, just to get the intended effect.
Why fine-tuning and not, say, retrieval-augmented generation (RAG)? Because in this experiment I was not so much interested in retrieving facts, a scenario where RAG excels, as in getting an LLM to behave and write in a specific style, whatever its knowledge of the context. Compared to full training, fine-tuning doesn’t require a massive amount of data, so it’s cheaper. Also, just because: I always wanted to try fine-tuning as a technique and see how feasible it could be.
To avoid spending days or weeks fine-tuning a model on my computer, which has a rather old graphic card, I relied on Runpod , an online service for AI developers that provides on-demand pods with pre-configured GPUs and tools for a (relatively) small price. For less than $6 per hour, for example, you can lease a beast of a card, the Nvidia B200 (192gb of memory). The service has a convenient API with configurable auto-recharge and cost control mechanisms.
Entering a world full of mysterious buzzwords
After deciding to fine-tune a model, I consulted with Claude on the sanest methods to achieve that. We settled on QLoRA (Quantized Low-Rank Adaptation), which achieves fine-tuning not by altering each weight of an LLM, but by “freezing” them and putting an adapter on top, which is a small file that reshapes the model behavior (a bit like a mask, if you will). The Q in QLoRA means that the result is quantized, that is, compressed, reducing memory requirements.
Are you still with me? Good. If you think this is dense, it’s because it is.
Doing anything with LLMs at home these days is an exercise in compromises: you either sacrifice time, spend money, or curb your ambitious goals. I tried to strike a balance to get something meaningful in less than a weekend. I chose to try fine-tuning on two models, Llama 3.1 8B Instruct and Qwen 2.5 7B Instruct. At their size (around 8B) they run comfortably on a Macbook Air. I also tested a Llama base model (which is not trained to answer questions).
I tested fine-tuning under several different conditions: varying the volume of training materials (a subset vs. the full corpus), the number of epochs (training rounds), and structural parameters like the rank. I only hold a superficial knowledge of all this, but I trusted my agent to make the right choices, which I happily questioned at every step. For example, 3 epochs can result in “overfitting” in some cases; in the world of LLMs, that translates to excessive training. Fun times.
Adapters can only be applied to the target model you fine-tuned for. After training each adapter, I exported them to my laptop and converted and quantized them to a GGUF LoRA file, and then registered it as a local Ollama model I could run in my laptop for benchmarking purposes. The local-conversion approach is faster and requires no GPU, though inference is somewhat slower than a fully merged model. For the test at hand, I did not care about speed that much.
Training the adapters for all conditions took perhaps an entire day, including breaks, for a total cost of $50. Along the journey, I lost two adapters: Runpod is unforgiving of budget and deletes pods immediately if funding is zero (there’s a lesson learned, yes). Claude took care of setting up each run and following up with Runpod’s API. The /goal command of Claude Code was quite helpful to loop through each phase (in retrospect, I would have run it in YOLO mode ).
This table shows all the models I compared and their conditions:
Did the style transfer after fine-tuning?
I subjected each model to the same prompts:
Document malloc(), a staple C function, something the training materials might know about.
Document a fictitious ConnectWifi() Win32 API function. No presence in the training materials.
Explain what a REST API is in 1990s Microsoft style (the anachronistic test).
You can see all the questions and answers in this gist .
For the malloc() test, the unmodified models generated modern Markdown docs in the style of a README, while the fine-tuned models used a period correct structure, with a Synopsis block, a Return Value section, and so on. For the fictitious ConnectWifi() function, only the 3 epochs model maintained the fiction and documented it as if it was real, while the others broke the fourth wall to adhere to internal knowledge and resist the training.
The REST API exercise was quite interesting, too: Llama Instruct 40k failed, producing bland marketing prose. Claude attributed this to the heavy reinforcement training (RLHF) that Llama goes through to make it friendly and accessible. Qwen fine-tunes held the register way better, producing period-structured docs, using HTTP method names as verbs and formal headings. Qwen 192k was the strongest, opening like a chapter of the Windows 2000 Resource Kit.
Let me repeat that: a 7B model, trained on 1990s documentation and tested on a 2000s concept, produced a convincing chapter opening that could be mistaken for genuine period material. Style transferred. Wow. On the other hand, the base model, which is not trained to answer questions, but to autocomplete text, failed miserably, spurting raw corpus almost at random, hundreds of lines of garbage. Base models have no notion of “answer this question” or “complete this".
I finished the experiment by comparing the effect of rank between Qwen models, with 1 epoch, varying between rank 8 and 16. If I understood it correctly, rank 8 means each adapter matrix can only describe 8 independent patterns. It’s like having 8 dials to tune. With so few dials, the adapter can’t be too clever: it must commit fully to the strongest, most repeated patterns in the training data. Rank 16 is, in theory, more expressive and subtler.
The rank comparison shows that smaller adapters, with fewer degrees of freedom, commit to fiction more readily than larger ones; a rank 16 adapter can “escape” the corpus more easily. It also turned out that combining only 1 epoch with a moderate rank of 16 made hallucinations more frequent: the adapter is expressive enough to reach for a related concept but not reinforced enough to anchor on what the prompt is trying to say. Rank and epoch seem to interact — it’s like using a sound mixer. Interestingly, the cheaper the adapter, the more honest the impersonation.
Fine-tuned models make for convincing impersonators, but they’re not replacements
The fine-tuned models were great impersonators of Microsoft tech writers from the late 90s. The corpus impressed style and voice on the models, as well as some knowledge, while mostly retaining the models’ ability to describe novel concepts. It’s a relatively cheap process that could produce effective small models aimed at tasks such as reviews of style or drafting of new documents following in-house style guides.
Getting there, though, is not a simple ride. Fine-tuning a model, while cheap, requires a good amount of high-quality training data, which is not easy to produce. Even when you get your hands on it, you need to pick an underlying model that makes sense and is capable of accepting the additional training. And then, the multiple parameters at your disposal make the task of getting a fine-tuned model to the sweet spot a time-consuming proposition.
The reassuring takeaway is that such a model can never replace a human tech writer, only augment them. The fine-tuned models have the same lack of judgement as their non-tuned siblings, and they need abundant steering. Fabrice will have to wait.
