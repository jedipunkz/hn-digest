---
source: "https://www.bellingcat.com/resources/2026/06/25/how-to-use-ai-to-help-find-civilian-harm-conflict-report-monitor-war-machine-learning-telegram/"
hn_url: "https://news.ycombinator.com/item?id=48810076"
title: "XGBoost beat LLMs at finding civilian-harm posts in Ukraine war Telegram data"
article_title: "How to Use AI to Help Find Civilian Harm - bellingcat"
author: "Jimmc414"
captured_at: "2026-07-06T21:22:33Z"
capture_tool: "hn-digest"
hn_id: 48810076
score: 2
comments: 0
posted_at: "2026-07-06T20:28:19Z"
tags:
  - hacker-news
  - translated
---

# XGBoost beat LLMs at finding civilian-harm posts in Ukraine war Telegram data

- HN: [48810076](https://news.ycombinator.com/item?id=48810076)
- Source: [www.bellingcat.com](https://www.bellingcat.com/resources/2026/06/25/how-to-use-ai-to-help-find-civilian-harm-conflict-report-monitor-war-machine-learning-telegram/)
- Score: 2
- Comments: 0
- Posted: 2026-07-06T20:28:19Z

## Translation

タイトル: XGBoost がウクライナ戦争における民間人被害投稿の発見で LLM を破った Telegram データ
記事のタイトル: AI を使用して民間人への被害を発見する方法 - bellingcat
説明: 新しい機械学習モデルは、民間人への危害が含まれる可能性に関して Telegram の投稿をランク付けします。

記事本文:
AI を使用して民間人への被害を発見する方法 - bellingcat
調査
ミゲルはベリングキャットの捜査技術者です。彼はデータとコードを使用して調査し、ストーリーを伝達し、オンライン調査コミュニティと共同で、またはオンライン調査コミュニティのための研究ツールを実験して構築しています。
Nick は、オンラインのオープンソース情報を裁判や説明責任のプロセスにおける証拠として使用するなど、紛争の調査を専門とする専門家です。彼は2016年5月にベリングキャットへの寄稿を開始し、2019年5月から2025年3月までスタッフメンバーを務めました。彼のこれまでの仕事は、シリアでの化学兵器の使用、EU国境での移民に対する暴力、ロシアのウクライナ全面侵攻中の民間人への被害に焦点を当てていました。
AI を使用して民間人への被害を発見する方法
2022 年 2 月から 2025 年 9 月にかけて、ベリングキャットのスタッフとボランティアは、ロシアによるウクライナへの全面侵攻後の 2,500 件を超える民間人被害事件を収集、位置特定し、共有しました。
この取り組みの一環として、ベリングキャットは、テレグラムのソーシャルメディア投稿を民間人への危害事件が含まれる可能性に関してランク付けすることを目的とした新しい機械学習モデルをテストしました。
この新しい方法論により、必要な検索と選択にかかる時間が大幅に短縮され、研究者は民間人への被害が発生した事件の調査だけでなく、その検証に集中できるようになりました。
この記事は、同様のテーマを研究している他の人が私たちの研究から恩恵を受けることを期待して、私たちの方法論、倫理的考慮事項、および学んだ教訓を文書化しています。
民生被害に関するオープンソース研究はまだ比較的新しい分野であり、多くの課題を抱えています。最大の課題の 1 つは、作成されている大量のユーザー作成コンテンツを整理して分類し、関連性のあるものを見つけることです。
機械学習、人工知能の一種

アルゴリズムを使用して大量のデータからパターンを特定し、予測を行うインテリジェンスにより、このタスクをより効率的に行うことができます。
スーダンや中東の多くの地域では、民間人に多大な被害をもたらす紛争が続いており、このガイドは、これらの紛争を取材する人々に、機械学習を使用して事件の発見と分類にどのように役立つかの例を提供することを目的としています。ここからモデルのコード ノートブックにアクセスすることもできます。
私たちは「民間人への危害」を、武力紛争による民間人の死亡や傷害だけでなく、精神的外傷、生計の喪失、避難、インフラの破壊などによる民間人への広範かつ遅発的な影響も含めて定義しました。この定義は、民間人への危害に関する『民間人保護』の本に基づいています。
研究者によってすでに手動で検証されている民間人への危害を含む各 Telegram 投稿は、データ サイエンティストが肯定的な事例と呼ぶ、民間人への危害が確認された事例の初期データセットを構築するために使用されました。これらの Telegram 投稿について、合計 5,848 個の一意の URL を収集しました。マニュアルの収集では、関連する Telegram チャネルの投稿を毎日、最も古い投稿から最新の投稿まで調べて調査しました。特定の投稿が位置情報付き事件リストに追加されたと仮定すると、その投稿にフラグを立てた研究者は、Telegram 上でその前後に表示された投稿も調べて、それらの投稿にはフラグを立てなかったということになります。そのため、民間人への危害を含まない投稿の追加データセットとして、検証された民間人への危害の投稿の周囲にある 10 件の投稿を選択しました。削除された投稿や重複した投稿を除外した結果、非民間危害投稿は 48,545 件となり、これが否定的な事例となりました。
あなたの寄付は、画期的な調査を発表し、世界中で不正行為を明らかにする私たちの能力に直接貢献します。
T

ネガティブなインスタンスを過剰に表現するという選択は、現実世界をよりよく反映し、モデルのトレーニングに利用できるデータを増やすことを目的としています。
Telegram API からのメタデータ (公開時刻、反応、テキスト コンテンツなど) を使用して各 URL を強化しました。これらの投稿の一部は削除されていたため、オート アーカイバー データベースから以前に保存されていた、肯定的なインスタンスでのみ利用可能なバージョンを使用して、不足しているデータ ポイントを完成させました。
これらのモデルは数学的演算に基づいて予測スコアを計算するため、機械学習モデルのトレーニングには数値データが必要です。
私たちは、パターンを識別するモデルの能力を高めることを目的として、民間人への危害の可能性を示すキーワードなどの初期データセットの生データを、モデルが解釈できる数値スコア (または「特徴」) に変換することによってこれらを構築しました。特徴エンジニアリングとして知られるこのプロセスでは、データ サイエンティストが明示的なコンテキスト知識を提案できるため、モデルの結果を大幅に改善できます。
モデルのトレーニングに使用した機能の完全なリストは、この部分に付属のコード ノートブックに記載されています。多くの機能は、設定された数の Telegram チャネルを分類し、各投稿を個別に検査することで民間人被害のケースを手動でスクリーニングした経験から得た研究者の意見に直接影響を受けています。
使用される機能のいくつかは、 media_type 、 day_of_week などの各 Telegram 投稿に含まれるメタデータから直接構築されました。またはバイナリのもの: forwarded 、edited、reply_to 。
その他の機能には、エンゲージメント情報 ( views 、 forwards 、 total_reactions ) が含まれており、😭 絵文字をカウントするためのreaction_crying_face など、最もよく使用される絵文字の個別の機能も含まれています。
手作業による収集プロセスの経験を埋め込むために、研究者はリストを作成しました。

民間人への危害を示す可能性のある投稿を示唆するキーワードをウクライナ語とロシア語でまとめた。たとえば、「Шахед」と「КАБ」はそれぞれ「シャヘド」と「誘導航空爆弾」に翻訳されます。その頻度をカウントするための数値特徴を作成しました。
さらに、意味的類似性スコアを生成するためにのみ使用される、「負傷者」、「学校の影響を受けた」、「病院の影響を受けた」など、民間人への危害の可能性を意味のある形で示すいくつかの一般的な英語のキーワードも含めました。
意味的類似性スコアは、異なる単語やフレーズ間の意味の近さを判断するために使用される計算です。投稿テキストと各キーワードの意味的な類似性を取得するために、Sentence Transformer モデルを介して数値のリストでそれぞれを表しました。Sentence Transformer モデルは、単語をコンピューターが理解できるベクトルと呼ばれる数値表現に変換します。
次に、2 つのテキスト間の類似性を測定する最も一般的な方法の 1 つであるコサイン類似性を使用して、各ベクトル間の類似性のレベルを計算しました。
埋め込みの仕組みにより、この計算の結果は、-1 (意味的近接性なし) から 1 (同じ意味) までのスケールの数値になります。たとえば、「傷つけた」と「怪我をした」という単語は類似性スコアが高くなりますが、「住宅」と「怪我をした」は単語が意味的に似ていないため、マイナスのスコアになります。
最後に、モデルがウクライナにおける民間人への危害と各投稿の関連性を識別できるようにするために、言語モデルの BERT ファミリの多言語テキスト トランスフォーマーを使用して、投稿のテキスト全体を 768 個の数値のベクトルとして表現しました。このモデルは、意味を捉える方法で多くの言語のテキストを効率的に表現できます。つまり、同じ文を異なる言語で表現できます。

age は同様のエンベディングを生成し、トレーニングされた機械学習モデルはエンベディング内のパターンを検出できます。
民間人危害検知モデルのこの最初のプロトタイプでは、写真やビデオなどのメディア コンテンツから派生した機能はまったく含めていないことに注意することが重要ですが、これはモデルのパフォーマンスを向上させるための論理的な次のステップです。
モデルの選択、トレーニング、評価
それぞれ 893 の数値特徴を含む 54,393 行を使用して、予測モデルをトレーニングするために 4 つの機械学習アルゴリズムを選択しました。
そのシンプルさのため、ベースライン アルゴリズムとしてロジスティック回帰を選択しました。また、他の 3 つの「クラス最高」モデル、 Random Forest 、 XGBoost 、および LightGBM も選択しました。これらの選択は、モデルの解釈可能性と、このサイズの表形式データを処理できる機能に重点を置いています。たとえば、解釈可能性の欠如と、これらのモデルはより大きなデータセットで最適に機能するため、ニューラル ネットワークを避けました。
トレーニングされたモデルのパフォーマンスを真に評価するために、データセットを 3 つの部分に分割しました。
トレーニング セット – モデルがトレーニングされたデータ (データセット全体の行の 60%)
検証セット – モデルパラメーターを調整する際の中間評価に使用されます (全行の 20%)
テスト セット – 最終的なパフォーマンス評価では非表示にされるため、モデルは目に見えないデータ (行の残り 20%) で評価されました。
データセットを分割するために、ランダム分割ではなく層化分割を使用しました。この方法により、陽性例（つまり民間人への危害が確認された例）の割合が 3 つのセットすべてで約 11 パーセントで一貫したままであることが保証されました。
機械学習モデルのパフォーマンスを測定するために、テスト セットを通して機械学習モデルを実行し、正しい予測と誤った予測の数を測定しました。モード

ls は、各 Telegram 投稿に民間人への危害が含まれている可能性を 0 から 1 の間で出力します。そして、ほぼすべての投稿にフラグを立てる (0.1) か、ごく少数の投稿にフラグを立てる (0.9) の間で適切なバランスが得られるカットオフしきい値を見つけようとしました。
モデルの予測能力を測定するための評価指標には、主に 2 つのタイプがあります。リコールは、肯定的な事例 (つまり、既知の民間人への危害投稿) のどの部分がそのように正しくフラグ付けされたかを主張します。精度は、実際に民間人への危害を与える投稿である民間人への危害としてフラグが立てられた投稿の割合を測定します。
トレーニング段階では、すべての再現レベルにわたる精度を要約する指標である平均精度 (PR-AUC) を最大化するようにモデルを調整しました。この方法では精度も考慮されますが、リコールが優先されます。これは、スキップされる民間人への危害投稿の数を減らすようにモデルの選択を誘導するため、このユースケースには好ましい方法です。
次の表は、コイン投げ予測子のベースラインに対して PR-AUC を最良から最悪の順にモデルを並べ替えています。 ROC-AUC と F1 は、健全性チェックとして含まれる他の 2 つの評価指標です。簡単に言うと、ROC-AUC は 2 つのインスタンス (1 つはネガティブ、もう 1 つはポジティブ) を正しくランク付けする確率を測定します。 F1 は精度と再現率を均等にバランスさせ、最適なカットオフしきい値を設定します。
これらの結果から、すべてのメトリック間で比較したときに最高のスコアが得られた XGBoost を最終モデルとして選択しました。
これらのモデルは解釈可能なため、投稿に民間人への危害が含まれるかどうかを予測する際にどの特徴が最も役立つかを理解できます。上の表は、XGBoost モデルに決定を下すよう最も強く通知する上位 10 個の機能を示しています。
semantic_keywords_similarity : 投稿テキストと手動で選択されたキーワード「死傷者」、「損害」、「民間被害」の間の意味的近接性
バート：モデル

このリストの他の特徴のいくつかと同じ強さでテキストから意味を識別できました。トップ 10 にはこのケースが 3 つあります
リアクション_crying_face : 投稿上の泣き顔の絵文字によるリアクション
group_of_messages : 投稿に複数のメディア ファイルが含まれているかどうか
tags_in_text : 投稿内のカスタムのウクライナ語またはロシア語のキーワードの数
これらの結果は、一般的に、民間人への危害に関する Telegram の投稿を選択するときに予想される内容と一致しています。これには、多くの感情的な関与を引き起こす投稿や、民間人への危害に関するキーワードを使用した投稿が、このトピックに関連するコンテンツを含む可能性が最も高いものの 1 つとして挙げられます。すべてのモデルが XGBoost と同じ主要な機能を備えているわけではありません。実際、ランダム フォレスト モデルの最も重要な特徴は、投稿内に存在する泣き顔の絵文字の数であり、この手法が最初に考案されたときに研究者によって強調された柔らかいパターンです。
遡及的に、同じテスト データセットのサンプルを異なる大規模言語モデル (LLM) で実行して、同じ予測を行う能力を評価することにしました。
私たちは、トレーニングされたモデルの追加機能として LLM で生成されたスコアを含めることを目指しました。これは、正しい予測と相関がある場合に関連性があるものとしてキャプチャされます。
まず、1B と 2 つのローカル モデルを選択しました。

[切り捨てられた]

## Original Extract

A new machine learning model ranks Telegram posts on their likelihood of containing incidents of civilian harm.

How to Use AI to Help Find Civilian Harm - bellingcat
Investigations
Miguel is an Investigative Technologist for Bellingcat. He uses data and code to investigate and communicate stories, he experiments and builds research tools with and for the online investigations community.
Nick is an expert who specialises in the examination of conflict using online open source information, including as evidence in justice and accountability processes. He began contributing to Bellingcat in May 2016 and was a staff member from May 2019 to March 2025. His previous work has focused on the use of chemical weapons in Syria, violence against migrants on the borders of the EU and civilian harm during Russia’s full scale invasion of Ukraine.
How to Use AI to Help Find Civilian Harm
Between February 2022 and September 2025, Bellingcat staff and volunteers collected, geolocated, and shared more than 2,500 incidents of civilian harm following Russia’s full-scale invasion of Ukraine.
As part of this effort, Bellingcat tested a new machine learning model intended to rank Telegram social media posts on their likelihood of containing incidents of civilian harm.
This novel methodology dramatically reduced the search and selection time required, freeing researchers to focus on verifying incidents of civilian harm – not just searching for them.
This piece documents our methodology, ethical considerations and lessons learned in the hope that others researching similar topics can benefit from our work.
Open source research into civilian harm is still a relatively new field and it presents many challenges – one of the biggest is organising and sorting through the huge volume of user generated content being produced to find what is relevant.
Machine learning, a form of artificial intelligence that uses algorithms to identify patterns from large amounts of data and make predictions, can make this task more efficient.
With ongoing conflicts involving large amounts of civilian harm occurring in Sudan, and much of the Middle East, this guide aims to offer those covering these conflicts an example of how machine learning can be used to help find and sort incidents. You can also access the Code Notebook for our model here .
We defined “civilian harm” not just as civilian deaths or injuries resulting from armed conflict, but also the broader and delayed effects on civilians from mental trauma, loss of livelihood, displacement, destruction of infrastructure and more. This definition was informed by the Protection of Civilians book on civilian harm .
Each Telegram post containing civilian harm which had already been manually verified by researchers was used to build an initial dataset of confirmed cases of civilian harm, which data scientists call positive instances . We collected a total of 5,848 unique URLs for these Telegram posts. For our manual collection we reviewed posts on relevant Telegram channels, working through oldest to newest posts each day. Assuming that a given post made it to our geolocated incidents list, it meant the researcher who flagged it also looked at the posts that appeared before and after it on Telegram and did not flag those ones, so we selected the 10 posts surrounding the verified civilian harm post as our additional dataset of posts that did not contain civilian harm. After excluding any deleted or duplicate posts, we ended up with 48,545 non-civilian harm posts, our negative instances .
Your donations directly contribute to our ability to publish groundbreaking investigations and uncover wrongdoing around the world.
The choice to overrepresent negative instances aims at better reflecting the real world and increasing data available for model training.
We enriched each URL with metadata from the Telegram API, such as the time of publication, reactions or textual content. As some of these posts had been deleted, we completed the missing data points with previously preserved versions from our Auto Archiver database, only available for the positive instances.
Training a machine learning model requires numerical data, as these models compute a prediction score based on mathematical operations.
We built these by converting raw data from our initial dataset, such as keywords signalling potential civilian harm, into numerical scores (or “features”) that the model could interpret, with the aim of increasing the model’s ability to identify patterns. This process, known as feature engineering , can significantly improve model results because it allows data scientists to suggest explicit context knowledge.
A full list of features we used to train the model can be found in the code notebook accompanying this piece. Many features were directly inspired by researchers’ input from their experiences manually screening cases of civilian harm by sorting through a set number of Telegram channels and inspecting each post individually.
Several of the features used were directly built from the metadata contained in each Telegram post including media_type , day_of_week ; or binary ones: forwarded , edited and reply_to .
Other features included engagement information: views , forwards , total_reactions , and even individual features for most used emojis including the reaction_crying_face to count 😭 emoji.
To embed the experience from the manual collection process, researchers put together a list of keywords both in Ukrainian and Russian that, to them, signalled posts likely to show civilian harm. For instance, “Шахед” and “КАБ” translated to “Shahed” and “Guided aerial bomb” respectively. We created a numerical feature to count their frequency.
In addition, we included several generic English-language keywords which meaningfully signalled potential civilian harm, such as “injured”, “school affected” and “hospital affected” that were only used for generating semantic similarity scores.
A semantic similarity score is a calculation used to determine the proximity in meaning between different words and phrases. To get the semantic similarity between the post text and each of our keywords, we represented each in a list of numbers via a Sentence Transformer model , which converts words into numerical representations called vectors that a computer can understand.
We then calculated the level of similarity between each vector using cosine similarity , one of the most popular methods for measuring similarity between two pieces of text.
Due to how embeddings work, this calculation results in a figure on a scale from -1 (no semantic proximity) to 1 (same meaning). For example, the words “hurt” and “injured” would have a high similarity score, while “residential” and “injured” would have a negative score as the words are not semantically similar.
Finally, to enable the model to identify the relevance of each post to civilian harm in Ukraine, we used a multilingual text transformer from the BERT family of language models to represent the entire post’s text as a vector of 768 numerical values. This model can efficiently represent text from many languages in a way that captures meaning: the same sentence in different languages will generate similar embeddings, and trained machine learning models can detect patterns in the embeddings.
It is important to note that for this initial prototype of a civilian harm detection model, we did not include any features derived from media content such as photos and videos, although that would be a logical next step in attempting to improve model performance.
Selecting, Training and Evaluating Models
With 54,393 rows of 893 numerical features each, we selected four machine learning algorithms to train our predictive models.
We chose Logistic Regression as a baseline algorithm due to its simplicity. We also selected three other “best in class” models, Random Forest , XGBoost , and LightGBM . These choices centred on the interpretability of the models and their ability to work on tabular data of this size. For example, we avoided neural networks due to a lack of interpretability and because those models work best with a larger dataset.
To genuinely assess the performance of the trained models, we split our dataset into three parts:
A training set – the data the models were trained on (60 percent of the full dataset’s rows)
A validation set – used for an intermediary evaluation when tuning model parameters (20 percent of all rows)
A test set – hidden for the final performance assessment, so the models were evaluated on unseen data (remaining 20 percent of rows)
We used a stratified split to divide the dataset instead of a random split. This method ensured the proportion of positive instances (i.e. confirmed cases of civilian harm) remained consistent across all three sets at about 11 percent.
To measure the performance of machine learning models, we ran them through the test set and measured the number of correct and incorrect predictions. Models output a likelihood between 0 and 1 that each Telegram post contains civilian harm, and we tried to find a cut-off threshold that leads to a good balance between flagging almost every post (0.1) or flagging very few (0.9).
There are two main types of evaluation metrics to gauge a model’s prediction power. Recall asserts what fraction of positive instances (i.e. known civilian harm posts) were correctly flagged as such. Precision measures the fraction of posts flagged as civilian harm that are indeed civilian harm posts.
During the training phase, we tuned the models to maximise average precision (PR-AUC), a metric that summarises precision across all recall levels. While this method also accounts for precision, it prioritises recall, which is preferable for this use case as it steers model selection to reduce the number of civilian harm posts that are skipped.
The following table sorts models from best to worst PR-AUC against a baseline of a coin-flip predictor. ROC-AUC and F1 are two other evaluation metrics included as sanity checks. Simply put, ROC-AUC measures the probability of ranking two instances, one negative and one positive, correctly; F1 balances precision and recall equally and its best cut-off threshold value.
From these results, we selected XGBoost as our final model as it had the best scores when compared across all metrics.
Because these models are interpretable, we can understand which features are the most useful when predicting whether a post includes civilian harm. The above table shows the top 10 features that most strongly signal the XGBoost model to make a decision:
semantic_keywords_similarity : the semantic proximity between the post text and manually selected keywords “casualties”, “damage” and “civilian harm”
bert : the model was able to discern meaning from the text with the same strength as some of the other features in this list – there are three cases of this in the top 10
reaction_crying_face : reactions with crying face emojis on the post
group_of_messages : whether a post contains multiple media files
keywords_in_text : the number of custom Ukrainian or Russian keywords in the post
These results generally tally with what you might expect when selecting Telegram posts for instances of civilian harm, including that posts that generate a lot of emotional engagement and posts using keywords about civilian harm were among those most likely to contain content related to this topic. Not all models had the same top features as XGBoost. In fact, for the Random Forest model the most important feature was the number of crying face emojis present in a post, a soft pattern highlighted by researchers when this methodology was first imagined.
Retroactively, we decided to run a sample of the same test dataset through different large language models (LLMs) to gauge their ability to make these same predictions.
We aimed to include an LLM-generated score as an extra feature for our trained models, which would be captured as relevant if it correlated with the correct predictions.
To start, we selected two local models, the 1B and

[truncated]
