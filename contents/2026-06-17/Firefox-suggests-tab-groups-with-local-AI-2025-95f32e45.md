---
source: "https://blog.mozilla.org/en/firefox/ai-tab-groups/"
hn_url: "https://news.ycombinator.com/item?id=48567405"
title: "Firefox suggests tab groups with local AI (2025)"
article_title: "Under the hood: How Firefox suggests tab groups with local AI"
author: "Topfi"
captured_at: "2026-06-17T10:38:57Z"
capture_tool: "hn-digest"
hn_id: 48567405
score: 2
comments: 0
posted_at: "2026-06-17T08:24:15Z"
tags:
  - hacker-news
  - translated
---

# Firefox suggests tab groups with local AI (2025)

- HN: [48567405](https://news.ycombinator.com/item?id=48567405)
- Source: [blog.mozilla.org](https://blog.mozilla.org/en/firefox/ai-tab-groups/)
- Score: 2
- Comments: 0
- Posted: 2026-06-17T08:24:15Z

## Translation

タイトル: Firefox がローカル AI を使用してタブ グループを提案 (2025)
記事のタイトル: 内部: Firefox がローカル AI を使用してタブ グループを提案する方法
説明: Mozilla のタブ グループ化機能は AI を使用してブラウザー タブを自動的に整理し、ラベルを付けることで、開いているページの管理と検索を容易にします。

記事本文:
内部: Firefox がローカル AI を使用してタブ グループを提案する方法
内部: Firefox がローカル AI を使用してタブ グループを提案する方法 |モジラのブログ
コンテンツにスキップ
モジラ
インターネット文化
オンラインライフ
内部: Firefox がローカル AI を使用してタブ グループを提案する方法
Mozilla は 2025 年初頭にタブ グループ化を開始し、永続的なラベルを使用してタブを配置およびグループ化できるようにしました。これは、Mozilla Connect の歴史の中で最もリクエストの多かった機能です。タブのグループ化はタブを管理し、タブの過負荷を軽減する優れた方法ですが、多数のタブを開いている場合、どのタブをグループ化するかを見つけるのが困難になる場合があります。
私たちは、次の 2 つの主要な機能を可能にする AI タブ グループ化機能を提供することで、ワークフローの改善を目指しました。
ユーザーがタブ グループを作成するときに、そのタブ グループのタイトルを提案します。
現在のウィンドウからタブ グループに追加するタブを提案します。
もちろん、ユーザーがデータを Mozilla に送信しなくてもこれが機能することを望んでいたため、ローカルの Firefox AI ランタイムを使用し、ユーザー自身のデバイス上ですべての機能を提供する効率的なモデルを構築しました。この機能はオプトイン式で、ユーザーがクリックして初めて実行するときに 2 つの小さな ML モデルをダウンロードします。
タブを最初にグループ化するときはユーザーの意図を理解するのが難しいため、グループ化されたタブのタイトルを提案するのは困難です。プロジェクト開始時のインタビューによると、タブ グループは「ショッピング」や「旅行」などの一般的な用語である場合もありますが、半数以上の場合、ユーザーのタブはビデオ ゲーム、友達、町の名前などの特定の用語であることがわかりました。また、タブ名が非常に短く、1 語または 2 語であることもわかりました。
グループのダイジェストの生成
これらの課題に対処するために、私たちは修正された TF-IDF ベースのテキスト分析とキーワード抽出を組み合わせたハイブリッド手法を採用しています。用語を特定する

タブ グループ内のページのタイトルが、タブ グループ外のページのタイトルと比較して統計的に区別される。次に、最も顕著な 3 つのキーワードと、ランダムに選択された 3 ページの完全なタイトルを組み合わせて、グループを表す簡潔なダイジェストを生成します。このダイジェストは、言語モデルを使用する後続の処理段階の入力として使用されます。
ダイジェスト文字列は、最終ラベルを返す生成モデルへの入力として使用されます。 10,000 を超える状況とラベルの例に基づいて微調整された、T5 ベースのエンコーダー/デコーダー モデル (flan-t5-base) を使用しました。
モデル開発における重要な課題の 1 つは、ユーザー データなしでモデルを調整するためのトレーニング データ サンプルを生成することでした。これを行うために、一連のユーザー アーキタイプを定義し、LLM API (OpenAI GPT-4) を使用して、さまざまなタスクを実行するユーザー用のサンプル ページを作成しました。これは、公開されている共通のクロール データセットからの実際のページ タイトルによって強化されました。次に、LLM を使用して、それらのユースケースに短いタイトルを提案しました。このプロセスは、最初は数百のグループ名の小規模で行われました。これらは手作業で修正および厳選され、簡潔さと一貫性を調整しました。プロセスがスケールアップするにつれて、最初の 300 のグループ名が LLM に渡されるサンプルとして使用され、作成される追加のサンプルがこれらの基準を満たすようになりました。
ほとんどのコンピューターで実行できるほど小さいモデルを取得する必要があります。最初のモデルがトレーニングされると、知識の蒸留として知られるプロセスを使用して、より小さなモデルにサンプリングされます。蒸留については、教師 flan-t5-base モデルのトークン確率出力から t5-efficient-tiny モデルを調整しました。蒸留プロセスの途中で、パラメータの数をさらに減らすために、2 つのエンコーダー トランス層と 2 つのデコーダー層も削除しました。
最後に、MOD

el パラメータは浮動小数点 (パラメータあたり 4 バイト) から整数 8 ビットに量子化されました。最終的に、この削減プロセス全体によりモデルは 1 GB から 57 MB に削減されましたが、精度の低下はわずかでした。
タブの提案については、ユーザーがタブをグループ化することを好む方法についていくつかのアプローチを特定しました。たとえば、仕事ですべてのドキュメントに簡単にアクセスできるように、ドメインごとにグループ化することを好む人もいます。旅行を計画するときにすべてのタブをグループ化することを好む人もいます。依然として「仕事」タブと「個人」タブを分けることを好む人もいるかもしれません。
タブの提案に関する私たちの最初のアプローチは、意味上の類似性に基づいていました。トピック的に類似したタブが提案されます。
トピック的に類似したタブの特定
まず、MiniLM 埋め込みモデルを使用して、タブ タイトルをローカルで特徴ベクトルに変換します。埋め込みモデルは、同様のコンテンツが埋め込み空間内で互いに接近したベクトルを生成するようにトレーニングされます。コサイン類似度などの類似性の尺度を使用すると、タブのタイトルや URL が他のタブ タイトルや URL にどれだけ似ているかを割り当てることができます。
ユーザーが選択したアンカー タブと別のタブの間の類似性スコアは、候補タブとアンカー タブのグループ タイトル (存在する場合)、アンカー タブのタイトル、およびアンカー URL との線形結合です。これらの値を使用して類似性の確率を生成し、確率のしきい値が高いタブがグループの一部であることが示唆されます。
どこで、
wは重量、
t_i は候補タブです。
t_a はアンカー タブです。
g_a はアンカーグループのタイトルです。
u_i は候補 URL です
u_a はアンカー URL であり、
σ はシグモイド関数です
重みを見つけるために、問題を分類タスクとして組み立て、アンカー タブを指定して正しく分類されたタブに基づいて精度と再現率を計算します。私たちが使用した

上記のユーザー アーキタイプに基づいて OpenAI によって生成された合成データ。
私たちは当初、ベースラインを確立するためにクラスタリング アプローチを使用していましたが、グループ、タイトル、URL の特徴をさまざまな重要度で扱うことでメトリクスが改善されることがわかったときに、ロジスティック回帰に切り替えました。
ロジスティック回帰を使用すると、ベースラインに対して 18% の改善が見られました。
この機能を使用するユーザーのタブ数の中央値は比較的少ない (~25) ものの、タブ数が数千に達する「パワー」ユーザーもいます。これにより、タブのグループ化機能に不快なほど長い時間がかかります。
これが、クラスタリング ベースのアプローチから線形モデルに切り替えた理由の 1 つでした。
当社のパフォーマンス フレームワークを使用したところ、KMeans などのクラスタリング ベースの手法と比較して、ロジスティック回帰の実行 p99 が 33% 改善されたことがわかりました。
今後の課題としては、F1 スコアの向上が挙げられます。これらは、推論の一部として時間関連コンポーネントを追加することによって (同時に開いたタブをグループ化する可能性が高くなります)、またはユースケースに合わせて微調整された埋め込みモデルを使用することによって可能になります。
私たちの仕事はすべてオープンソースです。開発者の方は、モデル トレーニングのソース コードを自由に熟読したり、Huggingface でトピック モデルを閲覧したりしてください。
ぜひこの機能を試して、ご意見をお聞かせください。
Firefox での検索候補の改善
Firefox の早期アクセスの導入 組織向けサポート
Firefox はかつてないほどカスタマイズしやすくなりました
今年 6 月の Firefox の新機能と Firefox ロードマップの次の予定
Firefox の無料組み込み VPN を使用して、夏中さらにプライベートにブラウジングしましょう
Firefox の最新情報を入手する
Firefox エクスペリエンスを最適なものにするためのハウツー、アドバイス、ニュースを入手してください。
このプライバシー ポリシー で説明されているように、Mozilla が私の情報を処理することに問題はありません。

今すぐサインアップ
あなただけを送ります
Mozilla 関連の情報。
Mozilla 関連のニュースレターの購読をまだ確認していない場合は、確認する必要がある場合があります。弊社からのメールが届いていないか、受信箱またはスパムフィルターをご確認ください。
ベータ版、ナイトリー版、開発者版

## Original Extract

Mozilla’s tab grouping feature uses AI to automatically organize and label browser tabs, making it easier to manage and find open pages.

Under the hood: How Firefox suggests tab groups with local AI
Under the hood: How Firefox suggests tab groups with local AI | The Mozilla Blog
Skip to content
Mozilla
Internet Culture
Life Online
Under the hood: How Firefox suggests tab groups with local AI
Mozilla launched Tab Grouping in early 2025, allowing tabs to be arranged and grouped with persistent labels. It was the most requested feature in the history of Mozilla Connect . While tab grouping provides a great way to manage tabs and reduce tab overload, it can be a challenge to locate which tabs to group when you have many open.
We sought to improve the workflows by providing an AI tab grouping feature that enables two key capabilities:
Suggesting a title for a tab group when it is created by the user.
Suggesting tabs from the current window to be added to a tab group.
Of course, we wanted this to work without you needing to send any data of yours to Mozilla, so we used our local Firefox AI runtime and built an efficient model that delivers the features entirely on your own device. The feature is opt-in and downloads two small ML models when the user clicks to run it the first time.
Suggesting titles for grouped tabs is a challenge because it is hard to understand user intent when tabs are first grouped. Based on our interviews when we started the project, we found that while tab groups are sometimes generic terms like ‘Shopping’ or ‘Travel’, over half the time users’ tabs were specific terms such as name of a video game, friend or town. We also found tab names to be extremely short – 1 or 2 words.
Generating a digest of the group
To address these challenges, we adopt a hybrid methodology that combines a modified TF-IDF–based textual analysis with keyword extraction. We identify terms that are statistically distinctive to the titles of pages within a tab group compared to those outside it. The three most prominent keywords, along with the full titles of three randomly selected pages, are then combined to produce a concise digest representing the group, which is used as input for the subsequent stage of processing using a language model.
The digest string is used as an input to a generative model that returns the final label. We used a T5 based encoder-decoder model (flan-t5-base) that was fine tuned on over 10,000 example situations and labels.
One of the key challenges in developing the model was generating the training data samples to tune the model without any user data. To do this, we defined a set of user archetypes and used an LLM API (OpenAI GPT-4) to create sample pages for a user performing various tasks. This was augmented by real page titles from the publicly available common crawl dataset . We then used the LLM to suggest short titles for those use cases. The process was first done at a small scale of several hundred group names. These were manually corrected and curated, adjusting for brevity and consistency. As the process scaled up, the initial 300 group names were used as examples passed to the LLM so that the additional examples created would meet those standards.
We need to get the model small enough to run on most computers. Once the initial model was trained, it was sampled to a smaller model using a process known as knowledge distillation. For distillation, we tuned a t5-efficient-tiny model from the token probability outputs of our teacher flan-t5-base model. Midway through the distillation process we also removed two encoder transformer layers and two decoder layers to further reduce the number of parameters.
Finally, the model parameters were quantized from floating point (4 bytes per parameter) to integer 8 bit. In the end this entire reduction process reduced the model from 1GB to 57 MB, with only a modest reduction in accuracy.
For tab suggestions, we identified a couple of approaches on how people prefer grouping their tabs. Some people prefer grouping by domain to easily access all documents for work for instance. Others might prefer grouping all their tabs together when they are planning a trip. Others still might prefer separating their “work” and “personal” tabs.
Our initial approach on suggesting tabs was based on semantic similarity. Tabs that are topically similar are suggested.
Identifying topically similar tabs
We first convert tab titles to a feature vector locally using a MiniLM embedding model. Embedding models are trained so that similar content produces vectors that are close together in embedding space. Using a similarity measure such as cosine similarity, we’re able to assign how closely similar a tab title or url is to another.
The similarity score between an anchor tab chosen by the user and another tab is a linear combination of the candidate tab with the group title (if present) of the anchor tab, the anchor tab title and the anchor url. Using these values, we generate a similarity probability and tabs that have a high probability threshold are suggested to be part of the group.
where,
w is the weight,
t_i is the candidate tab,
t_a is the anchor tab,
g_a is the anchor group title,
u_i is the candidate url
u_a is the anchor url, and,
σ is the sigmoid function
In order to find the weights, we framed the problem as a classification task, where we calculate the precision and recall based on the tabs that were correctly classified given an anchor tab. We used synthetic data generated by OpenAI based on the user archetypes above.
We initially used a clustering approach to establish a baseline and switched to a logistic regression when we realized that treating the group, title and url features with varying importances improved our metrics.
Using logistic regression, there was an 18% improvement against the baseline.
While the median number of tabs for people using the feature is relatively small (~25), there are some “power” users whose tab count reaches the thousands. This would cause the tab grouping feature to take uncomfortably long.
This was part of the reason why we switched from a clustering based approach to a linear model.
Using our performance framework, we found that the p99 of running logistic regression compared to a clustering based method such as KMeans improved by 33%.
Future work here would involve improving F1 score. These could be by adding a time-related component as part of the inference (we are more likely to group tabs together that we’ve opened at the same time) or using a fine-tuned embedding model for our use case.
All of our work is open source. If you are a developer feel free to peruse our source code on our model training, or view our topic model on Huggingface .
Feel free to try the feature and let us know what you think!
Better search suggestions in Firefox
Introducing early access for Firefox Support for Organizations
Firefox is easier than ever to customize
What’s new in Firefox this June, and what’s next on the Firefox roadmap
Browse more privately all summer with Firefox’s free built-in VPN
Keep up with all things Firefox
Get how-tos, advice and news to make your Firefox experience work best for you.
I’m okay with Mozilla handling my info as explained in this Privacy Policy .
Sign up now
We will only send you
Mozilla-related information.
If you haven’t previously confirmed a subscription to a Mozilla-related newsletter you may have to do so. Please check your inbox or your spam filter for an e-mail from us.
Beta, Nightly, Developer Edition
