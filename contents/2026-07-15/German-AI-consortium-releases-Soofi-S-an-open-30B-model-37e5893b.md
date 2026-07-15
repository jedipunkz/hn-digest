---
source: "https://the-decoder.com/german-ai-consortium-releases-soofi-s-an-open-30b-model-that-tops-benchmarks-in-both-english-and-german/"
hn_url: "https://news.ycombinator.com/item?id=48923203"
title: "German AI consortium releases Soofi S, an open 30B model"
article_title: "German AI consortium releases Soofi S, an open 30B model that tops benchmarks in both English and German"
author: "yogthos"
captured_at: "2026-07-15T17:07:11Z"
capture_tool: "hn-digest"
hn_id: 48923203
score: 7
comments: 1
posted_at: "2026-07-15T16:21:47Z"
tags:
  - hacker-news
  - translated
---

# German AI consortium releases Soofi S, an open 30B model

- HN: [48923203](https://news.ycombinator.com/item?id=48923203)
- Source: [the-decoder.com](https://the-decoder.com/german-ai-consortium-releases-soofi-s-an-open-30b-model-that-tops-benchmarks-in-both-english-and-german/)
- Score: 7
- Comments: 1
- Posted: 2026-07-15T16:21:47Z

## Translation

タイトル: ドイツの AI コンソーシアムがオープン 30B モデルである Soofi S をリリース
記事のタイトル: ドイツの AI コンソーシアムが、英語とドイツ語の両方でベンチマークを上回るオープン 30B モデルである Soofi S をリリース
説明: ドイツの研究コンソーシアムは、ミュンヘンにあるドイツテレコムのクラウド インフラストラクチャ上で完全にトレーニングされたオープン言語モデルである Soofi S 30B-A3B をリリースしました。このモデルは、トークンあたり 316 億のパラメータの一部のみをアクティブにする効率的なハイブリッド アーキテクチャを使用し、スループットを安定に保ちます。
[切り捨てられた]

記事本文:
ドイツの AI コンソーシアムが、英語とドイツ語の両方でベンチマークを上回るオープン 30B モデルである Soofi S をリリース
広告
コンテンツにスキップ
デコーダー
アップデート
AI研究
URLをクリップボードにコピーします
この記事をシェアする
コメントセクションに移動
ドイツの AI コンソーシアムが、英語とドイツ語の両方でベンチマークを上回るオープン 30B モデルである Soofi S をリリース
ドイツの研究コンソーシアムは、ドイツテレコムの AI クラウド インフラストラクチャ上で完全にトレーニングされたオープンソース言語モデル Soofi S をリリースしました。
このモデルは、トークンごとに 316 億のパラメーターのうち 3.2 個だけをアクティブにするリソース効率の高いハイブリッド アーキテクチャを使用しており、非常に長い入力があっても処理速度を一定に保ちます。
Soofi S はドイツ語のトレーニング データに重点​​を置いており、ドイツ語、英語、プログラミング タスクのベンチマークにおいて、Olmo 3 32B や Apertus 70B などの他の完全にオープンなモデルよりも優れたパフォーマンスを発揮します。
オーバートレーニング疑惑に関する声明を追加
発売後、批評家たちは、Soofi S が古典的なチンチラのスケーリングの法則の基準によって大幅に「過剰訓練」されていると主張しました。 Google DeepMind は 2022 年にこれらの法則を公開し、固定のコンピューティング予算内でモデル サイズとトレーニング データのバランスをとる方法を説明しました。彼らが特定したスイート スポットは、パラメータごとにおよそ 20 トークンでした。 Soofi S はその比率を超えています。約 27 兆のトークンと 300 億のパラメータを使用すると、数百対 1 になります。トークンごとにアクティブな 32 億のパラメーターのみを考慮すると、比率は数千対 1 に跳ね上がります。
プロジェクトの技術リーダーの一員である Michael Fromm 氏は、その批判に反論します。同氏は、これらのルールは単に専門家混合（MoE）アーキテクチャに引き継がれるわけではないと主張する。 「高密度モデルからの古いスケーリングの法則はもはや MoE アーキテクチャには適用されないことを示す新しい研究があります」とフロム氏は述べています。理由はhにあります

MoE モデルが構築されるまでの流れ。個々の専門家は同じドキュメントを見ることで恩恵を受けるため、大規模で高品質のデータセット内でデータが繰り返されても、高密度モデルの場合よりも問題が少なくなります。フロム氏は比較のポイントとして、最大 25 兆のトークンで独自のモデルをトレーニングした Nvidia を挙げています。広告
2026 年 7 月 13 日の元の記事: Ad DEC_D_Incontent-1
Soofi S は、ミュンヘンにあるドイツテレコムの産業用 AI クラウド上で完全にトレーニングされた最初の大規模言語モデルの 1 つです。オープン 30B モデルは、無駄のないハイブリッド アーキテクチャと、意図的にドイツ語に重点を置いたトレーニング ミックスを使用しています。
KI Bundesverband (ドイツ AI 協会) が調整するドイツの研究コンソーシアムは、オープン言語モデルである Soofi S 30B-A3B をリリースしました。その事前トレーニング レポートによると、完全オープン モデルの中で英語とドイツ語のベンチマークで最高のスコアを達成し、OLMo 3 32B や Apertus 70B などの以前のリーダーを上回りました。広告
長いコンテキスト向けに構築された無駄のないアーキテクチャ
Soofi S は専門家の混合モデルです。合計 316 億のパラメーターが含まれていますが、生成されたトークンごとに有効化されるのは約 32 億のみです。これにより、コンピューティング コストは従来の 300 億モデルよりも 300 億モデルに近づきます。このコンソーシアムは、Nvidia の Nemotron 3 Nano のアーキテクチャを変更せずに採用しており、Mamba-2 レイヤーと標準のアテンション レイヤーを組み合わせたハイブリッド設計です。
一般的なトランスとの主な違いはメモリの動作です。従来のモデルでは、アテンションの計算のために以前のトークンを保存する KV キャッシュは、コンテキストの長さに応じて直線的に増加します。長い入力と多くの並列リクエストがある場合、そのキャッシュのリロードがボトルネックになります。 Soofi S の 52 レイヤーのうち、そのようなキャッシュをまったく維持しているのは 6 レイヤーだけです。広告 DEC_D_Incontent-2
実際の見返りは、生成スループットに現れます。コンテキスト長 40,000 t の場合

32 の並列リクエストを持つオーケンの場合、Soofi S は 140 ～ 240 億のパラメーター範囲の高密度モデルに比べ、GPU あたり 1 秒あたり約 8 倍多くのトークンを生成します。従来のモデルではコンテキストが拡大するにつれてスループットが大幅に低下しますが、Soofi S は 4,000 トークンから 256,000 トークンまでほぼ横ばいです。測定で同様の動作を示した唯一のモデルは Alibaba の Qwen3.5 35B-A3B で、これもハイブリッド アーキテクチャを使用しています。広告
ドイツ語を中心に構築されたトレーニングミックス
コンソーシアムは 3 つのフェーズに分けて、合計約 27 兆のトークンを処理しました。最初のフェーズでは、モデルは、Web、コード、数学、ドメイン固有のテキストの幅広い組み合わせから抽出された約 20 兆のトークンから言語の基礎を学習します。第 2 段階では、以前に学習したパターンを鮮明にするように設計された、高品質のソースからの約 6 兆個のトークンが使用されます。次に、より短い 3 番目のフェーズでは、最大 100 万トークンの非常に長いドキュメントをトレーニングすることでコンテキスト ウィンドウを拡張します。
ドイツ語に意図的に焦点を当てています。第 1 段階では、ドイツ語がトレーニング ミックスの 7.2% を占めます。第 2 フェーズでは、そのシェアは 15.3 パーセントに上昇します。 Nvidia の Nemotron リファレンス レシピでは、英語以外の言語を合わせても約 5% しか占めていません。
データ ソースとして、コンソーシアムは、HPLT のドイツ語 Web テキスト、オープンにライセンスされたドイツ コモンズ コーパス、FinePDF と FineWiki のドイツ語部分、および 916 のドイツ出版物からの 1 億 9,300 万の新聞記事を含む商用ライセンスされた Genios コーパスを組み合わせています。機械翻訳され、合成的に生成されたドイツ語テキストがミックスを完成させます。
ドイツ語と英語の両方でオープンモデルのトップスコアを獲得
レポートによると、他の 16 のオープン モデルと比較した評価では、Soofi S はドイツ語と英語の両方の合計スコアですべての完全オープン モデルをリードしています。あの株式会社

Allen Institute for AI の OLMo 3 32B とチューリッヒ工科大学および EPFL の Apertus 70B を採用しています。欧州のすべてのソブリンベースラインに対して、このモデルは、スイート内のすべてのドイツのベンチマークで、場合によっては 2 桁のマージンで上回っています。
コード ベンチマークでは、Soofi S のスコアは HumanEval で 73.8 パーセント、MBPP で 70.2 パーセント、ドイツ版 MBPP バリアントで 84.2 パーセントであり、オープンソース ピアの中で最高の結果となっています。ドイツ特有の地域知識を問うテストである INCLUDE-DE では、Soofi S が 61.2 ポイントで、より大型の Qwen3.5 35B-A3B と同点で 1 位となりました。 Nemotron ベースラインと比較して、ドイツ語データ レシピは英語の成績を犠牲にすることなく、言語能力を 15.1 ポイント向上させ、科学テスト GPQA-Diamond は 9.6 ポイント向上させました。
Soofi S はドイツの競技数学ではそれほど成績が良くなく、Minerva MATH-DE で 56 ポイントを獲得しており、Qwen3.5 35B-A3B (76.5) や Gemma 3 27B (65.6) に大きく遅れをとっています。また、NaturalQuestions でのオープンな事実の検索も遅れています。後者は、アクティブなパラメーターが 30 億しかないことに関連している可能性があり、高密度の 27B モデルよりも保存できる世界の知識が少なくなります。
RULER の長いコンテキスト テストでは、特定の弱点も明らかになりました。モデルが長いテキストから頻繁に出現する単語を抽出する必要がある場合、Soofi S のヒット率は 32,000 トークンのコンテキストを超えると約 3 パーセントに低下しますが、同等の Nemotron モデルは依然として 60 ～ 64 パーセントを管理します。著者らは、これを、長いコンテキストのトレーニング データには多くの長いドキュメントが含まれているものの、抽出タスク用に設計された合成データが欠けているという事実に起因していると考えています。残りの 12 個の RULER タスクでは、両方のモデルのパフォーマンスはほぼ同じです。
主権インフラと文書化されたオープン性
トレーニングの実行は、3 月から 5 月にかけて、ミュンヘンにあるドイツテレコムの産業用 AI クラウドにある最大 512 台の Nvidia B200 GPU、合計約 253,000 GPU で行われました。

何時間も。報告書によると、この施設は完全に再生可能エネルギーで稼働し、アイスバッハ運河からの水で冷却され、周囲のトゥッハーパーク地区に廃熱を供給しているという。 Soofi S は、このインフラストラクチャ上で実行された最初の大規模なトレーニングの 1 つです。
Soofi の背後には、ドイツの研究機関と企業のコンソーシアムがあり、ドイツ AI 協会が調整し、欧州 IPCEI-CIS プログラムの一環としてドイツ連邦経済エネルギー省から資金提供を受けています。
参加者には、フラウンホーファー研究所 IAIS および IIS、ドイツ人工知能研究センター (DFKI)、ダルムシュタット工科大学、ヴュルツブルク大学、L3S 研究センター、ベルリン応用科学大学、AI 企業の Ellamind および Merantix Momentum が含まれます。プロジェクトの目標は、主権インフラ上で実行でき、産業アプリケーションでテストできるオープンなヨーロッパ AI モデル ファミリを構築することです。
研究者らは、選択された中間チェックポイント、完全なトレーニングおよび評価コード、および生のトークン数、エポック番号、およびソースごとの効果的な寄与をリストした詳細なデータ目録とともに、モデルの重みを公開しています。レビューされたものの除外された情報源も文書化されます。チームによると、これは Soofi S が Open Source Initiative の Open Source AI Definition 1.0 に準拠していることを意味します。
ヨーロッパのオープンデータ定義に関するより厳格な提案は、すべてのトレーニング トークンを自由に配布できることを要求するものですが、商用ライセンスが付与されている Genios データのシェアが 1.3 パーセントであるため、満たされていません。レポートによると、トレーニング ミックスの約 99 パーセントは独立して再構築できるという。このモデルのリリースに対する正確なライセンスはまだ確定していません。
技術リーダーの Michael Fromm 氏が書いているように、Soofi S の立場は

多くの言語をカバーする EuroLLM や Teuken のような広範な多言語欧州主権プロジェクトと、最高のパフォーマンスを誇る国際無差別モデルとの間で、両者を比較することができます。プロジェクトの Web サイトによると、コンソーシアムは、技術文書、コード生成、エージェントベースのシステムを含むアプリケーションでモデルをテストする次の段階に向けて、業界パートナーを探しています。
誇大広告のない AI ニュース - 人間がキュレーション
THE DECODER を購読すると、広告なしの閲覧、週刊 AI ニュースレター、年 6 回の独占的な「AI レーダー」フロンティア レポート、完全なアーカイブ アクセス、コメント セクションへのアクセスが可能になります。
DeepSeek は、最初の 70 億ドルのラウンドを完了してからわずか数週間後にさらに多くの現金を必要としています
Anthropic が生徒データでモデルをトレーニングしないことを約束して教師向け Claude をオープン
OpenAI の最初のハードウェア製品は、生きていると感じるように設計されたスクリーンレス AI スピーカーです
OpenAI の Codex は AI エージェント間の命令を暗号化するようになり、開発者は内部委任を認識できなくなります
ナデラ氏、OpenAIやAnthropicなどのAI研究所に対し、他人のデータを使ってトレーニングする際の蒸留を禁止していると非難
AI に関する最新情報を常に把握しましょう。クリアで使いやすく、毛羽立ちがありません。
ドイツの画期的な判決、GoogleのAI概要はGoogle自身の言葉であると宣言し、虚偽の回答には責任を負わせる
オープンソース ツール pxpipe は PNG 内のテキストを非表示にして、クロード コードとフェイブル 5 のトークン コストを最大 70% 削減します
サム・アルトマン氏は、全世代の研究者がスケーリングで何ができるかを過小評価し、AIの導入を妨げてきたと語る
OpenClaw の創設者である Peter Steinberger 氏は月額 130 万ドルで、コーディング、PR のレビュー、バグの発見を行う 100 人の AI エージェントを運用しています。
Google、Gemma 3をローカルで実行する小さなボードを発表
フィールズメダリストは、ChatGPT 5.5 Pro が人間の助けなしで 2 時間以内に「博士レベル」の数学研究を実現したと述べています
AI のニュース、背景ストーリー、専門家については、The Decoder をフォローしてください

分析。
AI に関する最新情報を常に把握しましょう。クリアで使いやすく、毛羽立ちがありません。

## Original Extract

A German research consortium has released Soofi S 30B-A3B, an open language model trained entirely on Deutsche Telekom's cloud infrastructure in Munich. The model uses an efficient hybrid architecture that activates only a fraction of its 31.6 billion parameters per token, keeping throughput steady
[truncated]

German AI consortium releases Soofi S, an open 30B model that tops benchmarks in both English and German
Ad
Skip to content
The Decoder
Update
AI research
Copy the url to clipboard
Share this article
Go to comment section
German AI consortium releases Soofi S, an open 30B model that tops benchmarks in both English and German
A German research consortium has released the open-source language model Soofi S, which was trained entirely on Deutsche Telekom's AI cloud infrastructure.
The model uses a resource-efficient hybrid architecture that activates only 3.2 of its 31.6 billion parameters per token, keeping processing speed constant even with very long inputs.
With a strong focus on German training data, Soofi S outperforms other fully open models, such as Olmo 3 32B and Apertus 70B, in benchmarks for German, English, and programming tasks.
Statement on the allegation of overtraining added
After launch, critics argued that Soofi S was heavily "overtrained" by the standards of the classic Chinchilla scaling laws . Google DeepMind published those laws in 2022, describing how to balance model size and training data for a fixed compute budget. The sweet spot they identified was roughly 20 tokens per parameter. Soofi S blows past that ratio. With about 27 trillion tokens and 30 billion parameters, it lands at several hundred to one. Factor in only the 3.2 billion parameters active per token, and the ratio jumps to several thousand to one.
Michael Fromm, part of the project's technical leadership, pushes back on that criticism. He argues those rules don't simply carry over to Mixture-of-Experts (MoE) architectures. "There's new research showing that the old scaling laws from dense models no longer apply to MoE architectures," Fromm said. The reason comes down to how MoE models are built. Individual experts benefit from seeing the same documents, so repeated data in a large, high-quality dataset is less of a problem than it would be with dense models. As a point of comparison, Fromm points to Nvidia, which trained its own models on up to 25 trillion tokens. Ad
Original article from July 13, 2026: Ad DEC_D_Incontent-1
Soofi S is one of the first large language models trained entirely on Deutsche Telekom's Industrial AI Cloud in Munich. The open 30B model uses a lean hybrid architecture and a training mix deliberately weighted toward German.
A German research consortium coordinated by the KI Bundesverband (German AI Association) has released Soofi S 30B-A3B , an open language model that, according to its pretraining report , achieves the highest scores on English and German benchmarks among fully open models, surpassing previous leaders like OLMo 3 32B and Apertus 70B . Ad
A lean architecture built for long contexts
Soofi S is a mixture-of-experts model. It contains 31.6 billion parameters in total but activates only about 3.2 billion per generated token. That puts its compute cost closer to a 3B model than a conventional 30B model. The consortium adopts the architecture of Nvidia's Nemotron 3 Nano without modification, a hybrid design combining Mamba-2 layers with standard attention layers.
The key difference from typical transformers is memory behavior. In conventional models, the KV cache that stores previous tokens for attention computation grows linearly with context length. With long inputs and many parallel requests, reloading that cache becomes a bottleneck. Only 6 of Soofi S's 52 layers maintain such a cache at all. Ad DEC_D_Incontent-2
The practical payoff shows up in generation throughput. At a context length of 40,000 tokens with 32 parallel requests, Soofi S generates roughly eight times more tokens per second per GPU than dense models in the 14 to 24 billion parameter range. While throughput drops significantly for conventional models as context grows, Soofi S stays nearly flat from 4,000 to 256,000 tokens. The only model that shows similar behavior in the measurements is Alibaba's Qwen3.5 35B-A3B , which also uses a hybrid architecture. Ad
A training mix built around German
The consortium processed about 27 trillion tokens in total, split across three phases. In the first phase, the model learns language fundamentals from roughly 20 trillion tokens drawn from a broad mix of web, code, math, and domain-specific texts. A second phase follows with about 6 trillion tokens from higher-quality sources, designed to sharpen the patterns learned earlier. A shorter third phase then extends the context window by training on very long documents of up to one million tokens.
The deliberate focus on German is central. In the first phase, German makes up 7.2 percent of the training mix; in the second phase, that share rises to 15.3 percent. In Nvidia's Nemotron reference recipe, all non-English languages combined account for only about 5 percent.
For data sources, the consortium combines German web text from HPLT, the openly licensed German Commons corpus, German portions of FinePDFs and FineWiki, and the commercially licensed Genios corpus containing 193 million newspaper articles from 916 German publications. Machine-translated and synthetically generated German texts round out the mix.
Top open-model scores in both German and English
In evaluations against 16 other open models, Soofi S leads all fully open models on aggregate scores for both German and English, according to the report. That includes OLMo 3 32B from the Allen Institute for AI and Apertus 70B from ETH Zurich and EPFL. Against every European sovereign baseline, the model comes out ahead on all German benchmarks in the suite, sometimes by double-digit margins.
On code benchmarks, Soofi S scores 73.8 percent on HumanEval, 70.2 on MBPP, and 84.2 on the German MBPP variant, the best results among open-source peers. On INCLUDE-DE, a test for Germany-specific regional knowledge, Soofi S ties for first place at 61.2 points with the larger Qwen3.5 35B-A3B. Compared to the Nemotron baseline, the German data recipe improves language proficiency by 15.1 points and the science test GPQA-Diamond by 9.6 points, without sacrificing English performance.
Soofi S doesn't do as well on German competition math, where it scores 56 points on Minerva MATH-DE, well behind Qwen3.5 35B-A3B (76.5) and Gemma 3 27B (65.6). It also lags on open factual retrieval in NaturalQuestions. The latter likely relates to having only 3 billion active parameters, which can store less world knowledge than a dense 27B model.
The RULER long-context test also reveals a specific weakness: When the model has to extract frequently occurring words from a long text, Soofi S's hit rate drops to around 3 percent beyond 32,000 tokens of context, while the comparable Nemotron model still manages 60 to 64 percent. The authors attribute this to the fact that their long-context training data contains many long documents but lacks synthetic data designed for extraction tasks. On the remaining twelve RULER tasks, both models perform about the same.
Sovereign infrastructure and documented openness
The training run took place between March and May on up to 512 Nvidia B200 GPUs at Deutsche Telekom's Industrial AI Cloud in Munich, totaling about 253,000 GPU-hours. According to the report, the facility runs entirely on renewable energy, is cooled with water from the Eisbach canal, and feeds waste heat into the surrounding Tucherpark neighborhood. Soofi S was one of the first major training runs on this infrastructure.
Behind Soofi is a consortium of German research institutions and companies, coordinated by the German AI Association and funded by the German Federal Ministry for Economic Affairs and Energy as part of the European IPCEI-CIS program.
Participants include the Fraunhofer Institutes IAIS and IIS, the German Research Center for Artificial Intelligence (DFKI), TU Darmstadt, the University of Würzburg, the L3S Research Center, the Berlin University of Applied Sciences, and AI companies Ellamind and Merantix Momentum. The project's goal is to build an open European AI model family that can run on sovereign infrastructure and be tested in industrial applications.
The researchers are releasing model weights along with selected intermediate checkpoints , the complete training and evaluation code, and a detailed data inventory listing raw token counts, epoch numbers, and effective contributions per source. Sources that were reviewed but excluded are also documented. According to the team, this means Soofi S meets the Open Source AI Definition 1.0 from the Open Source Initiative .
A stricter proposal for a European open-data definition, which would require every single training token to be freely distributable, isn't met because of the 1.3 percent share of Genios data, which carries a commercial license. The report says about 99 percent of the training mix can be independently reconstructed. The exact license for the model's release hasn't been finalized yet.
As technical leader Michael Fromm writes , Soofi S positions itself between broadly multilingual European sovereignty projects like EuroLLM or Teuken , which cover many languages, and the highest-performing international open-weight models. According to the project website, the consortium is looking for industry partners for the next phase to test the model in applications involving technical documents, code generation, and agent-based systems.
AI News Without the Hype – Curated by Humans
Subscribe to THE DECODER for ad-free reading, a weekly AI newsletter, our exclusive "AI Radar" frontier report six times a year, full archive access, and access to our comment section.
DeepSeek needs more cash just weeks after closing its first $7 billion round
Anthropic opens Claude for Teachers with a promise not to train models on student data
OpenAI's first hardware product is a screenless AI speaker designed to feel alive
OpenAI's Codex now encrypts instructions between AI agents, leaving developers blind to internal delegation
Nadella calls out AI labs like OpenAI and Anthropic for banning distillation while training on everyone else's data
Stay in the loop on AI. Clear, useful, no fluff.
Landmark German ruling declares Google's AI Overviews are Google's own words and makes it liable for false answers
Open-source tool pxpipe hides text in PNGs to cut Claude Code and Fable 5 token costs up to 70%
Sam Altman says a whole generation of researchers held AI back by underestimating what scaling could do
For $1.3 million a month, OpenClaw founder Peter Steinberger runs 100 AI agents that code, review PRs, and find bugs
Google launches a tiny board that runs Gemma 3 locally
Fields Medalist says ChatGPT 5.5 Pro delivered "PhD-level" math research in under two hours with zero human help
Follow The Decoder for AI news, background stories and expert analyses.
Stay in the loop on AI. Clear, useful, no fluff.
