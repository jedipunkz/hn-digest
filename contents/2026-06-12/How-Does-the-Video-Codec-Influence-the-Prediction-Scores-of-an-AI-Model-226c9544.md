---
source: "https://www.fourthline.com/newsroom/supervised-vs-contrastive-classification"
hn_url: "https://news.ycombinator.com/item?id=48502882"
title: "How Does the Video Codec Influence the Prediction Scores of an AI Model?"
article_title: "Supervised vs. Contrastive Classification: How Does the Video Codec Influence the Prediction Scores?"
author: "Doch88"
captured_at: "2026-06-12T13:17:49Z"
capture_tool: "hn-digest"
hn_id: 48502882
score: 2
comments: 0
posted_at: "2026-06-12T11:45:38Z"
tags:
  - hacker-news
  - translated
---

# How Does the Video Codec Influence the Prediction Scores of an AI Model?

- HN: [48502882](https://news.ycombinator.com/item?id=48502882)
- Source: [www.fourthline.com](https://www.fourthline.com/newsroom/supervised-vs-contrastive-classification)
- Score: 2
- Comments: 0
- Posted: 2026-06-12T11:45:38Z

## Translation

タイトル: ビデオ コーデックは AI モデルの予測スコアにどのような影響を与えますか?
記事のタイトル: 教師あり分類と対照分類: ビデオ コーデックは予測スコアにどのように影響しますか?
説明: 当社は、クライアントの身元を確認し、プラットフォームにオンボーディングしてサービスを利用したいと考えているのは生きた人間であるかどうかを検証するために使用される機械学習および AI モデルを研究および開発しています。

記事本文:
教師あり分類と対照分類: ビデオ コーデックは予測スコアにどのような影響を与えますか?
Sebastian Vater - ID&V の研究員仲間
そもそも、なぜさまざまなコーデックとその予測スコアへの影響に興味があるのでしょうか?
私たちが働いている Fourthline の任務は、KYC (顧客確認) プロセスにおいて、銀行、フィンテック、非金融企業などの顧客の安全なオンボーディングを支援することです。当社は、クライアントの身元を確認し、プラットフォームにオンボーディングしてサービスを利用したいと考えている生きた人間であるかどうかを検証するために使用される機械学習および AI モデルを研究および開発しています。
私たちの研究室で機械学習システムを試しているときに、興味深いことがわかりました。入力ビデオのエンコード方法を変更すると、さまざまな損失関数でトレーニングされたモデルの反応が異なるということです。これにより、一部のトレーニング パラダイムが他のビデオ コーデックよりも異なるビデオ コーデック、つまり圧縮アーティファクトに対してより敏感であるかどうか (また、そうである場合はどの程度まで) という疑問が生じました。
では、エンコードまたはトランスコーディングがサービスの出力にどのような影響を与えるかを確認する必要があるのはなぜでしょうか?手持ちのコーデックに合わせて最適化するだけで満足できるでしょう。そうですね、学術界を離れて、私たちはオープンワールドで交流し、サービスも同様です。 Fourthline は、多数の銀行、フィンテック、非金融企業にサービスを提供しています。彼らはそれぞれ、独自のデータ キャプチャとバックエンド エンジニアリングを使用してデータを個別に前処理してから、データを当社のアルゴリズムで実行して、人の生存性と身元を確認します。私たちは、入力コーデックに関係なく、サービスが常に期待どおりに動作することを重視しています。
この投稿では、教師あり学習と対照学習という 2 つのパラダイムからトレーニング目標がどのように得られるかを具体的に検討したいと思います (概要については [8] を参照してください)

) - 作成された埋め込み空間に影響を与える可能性があります (例: ノイズやアーティファクトに対する構造や堅牢性)。これらの損失のいくつかの特徴を詳しく見て、これらが現実のシナリオでの分類動作にどのような影響を与えるかを見ていきたいと思います。特に、目標は、入力データに使用されているさまざまなコーデックに起因するものなど、画像空間アーティファクトが存在する場合に、さまざまな損失が分類スコアに与える影響を把握することです。
Fourthline の生体認証チームの Max Crous と Roman Aleksandrov が実験を実施し、このブログ投稿の基礎となるグラフを作成しました。さらに、ディスカッションを通じて貴重な洞察やアイデアを共有しました。
Fourthline の製品アーキテクチャでは、連携して動作する複数の統合モジュールを採用し、総合的な利点を最大化します。このモジュール式アプローチにより、当社のシステムは複数の並列処理パスを通じて課題に対処できるようになります。
一例として、liveness システムのさまざまな部分が、さまざまな学習パラダイムとアーキテクチャに従って研究およびトレーニングされます。その際、特に教師あり学習法と対照学習法を利用します。
アルゴリズムの結果に対するさまざまなビデオ コーデックの影響に興味があるため、ドメイン内のデータに対していくつかの実験的テストを実行し始め、さまざまなエンコードとトランスコーディングの前処理を適用しました。以下に、2 つの特定のシステムに関するいくつかの定性的な例示的な結果を示します。これらのシステムは両方とも同じデータ ドメインでトレーニングされ、一方は教師ありバイナリ クロス エントロピー損失 ( BCE ) でトレーニングされ、もう一方はコントラスト損失 ( Contrast. ) でトレーニングされ、決定しきい値の推定に局所密度推定器が使用されます。
以下のプロットはsh

入力ビデオ コーデックを変更したときの、異なるトレーニングを受けたシステムのスコア出力 (例: Softmax) の差の分布を確認します。分類されるサンプルの元のコーデックのスコアが任意の値 x0 (水平線「0.00」で定義) を持つ場合、色付きの点は、モデルに異なるエンコーディングのビデオが供給されたときに、各テスト サンプルの信頼度出力の値がどこに到達するかを示し、結果としてスコア x0+ d x になります。色付きの領域は d x の広がりを表します。
これらの実験では、元の MJPEG 形式 (ビデオの各フレームは JPEG 形式でエンコードされている) から始めて、クロマ サブサンプリングされた MJPEG (4:4:4 および 4:2:0)、フレーム内 FFV1 (FFmpeg プロジェクトの一部) を評価します。また、FFV1 とロスレス VP9 エンコーディングの比較も以下に示します。
入力コーデックを変更した場合の分類子スコアの差の分布。 Fourthline が生成および所有する画像。
それで、何が見えるでしょうか？コントラスト損失を使用してトレーニングされたモデルは、エンコードされたデータに対して、BCE でトレーニングされたモデルとは大きく異なるパフォーマンスを発揮することがすぐにわかります。 VP9 の場合、ほぼガウス統計を仮定すると、1 シグマ領域には +/- 0.05 のスコア差が含まれます。これは、信頼度が、たとえば 0.8 から 0.75 ～ 0.85 に変化することを意味します。これはモデルの出力にとって大きな違いです。
入力コーデックを変更した場合の分類子スコアの差の分布。 Fourthline が生成および所有する画像。
さまざまなビデオ コーデックに起因するアーティファクトの視覚化の例。元のビデオとエンコードされたビデオの間のピクセル単位の差異を計算します。結果は定性的に解釈されます。左から右へ: (1) 元の画像、紙のマスクを適用して Liveness システムを騙そうとしている人物 (ビデオ圧縮なし)。 (2) でエンコード

MJPEG。 (3) FFV1 でエンコードされます。 (4) VP9 でエンコードされています。画像は Fourthline が所有します。
「なぜ対照的なモデルはこれほど異なる動作をするのでしょうか?」と疑問に思われるかもしれません。そして、なぜ BCE モデルはそうではないのでしょうか? Contrastive モデルは、画像のエンコード/トランスコーディングに対してはるかに脆弱であるようです。私たちは、何が原因で、どうすればそれを防ぐことができるのか、と自問しました。
以下の段落では、考えられる説明に光を当て、この動作を説明する可能性のある理由をいくつか示します。 (注: これは単なるブログ投稿であるため、科学的に適切な証拠は他の場所でのみ公開されます。)
まず、異なるコーデック間の対照学習アプローチでは、スコアの乖離が一貫して高いことがわかります。すべての実験が同じデータドメインで行われたことを考えると、さまざまな学習スキームがこの現象をどのように説明できるかを詳しく調べることに焦点を当てています。
一般に、技術システムにおける望ましくない動作を防ぐには、観察された動作の原因を理解することが最も役立ちます。
2 つの損失: 特徴
思い浮かぶ最初の説明は、2 つの損失が異なるため、学習された特徴空間によって異なる構造、特に異なるクラス間構造が明らかになることです。損失に関するいくつかの特徴を見てみましょう (この投稿の議論では 2 項問題を想定しています)。
クロスエントロピー二値損失 (BCE):
BCE は 2 つのクラスを分離する超平面を見つけます。そこではクラスの尤度 (例: Softmax 出力) に基づいて操作されます。
モデルの学習された分布には制約はありません (クラス 1 とクラス 2 など、任意の形状を想定できます)。
分離可能な特徴 - 同じクラスの類似のデータ ポイントが互いに近いかどうかに関係なく (距離が「離れている」限り) あらゆる特徴を学習します。

他のクラスのすべてのデータ ポイントから離れてください)。
要約すると、BCE はクラス内分布や分布自体の形状にはあまり関心がなく、あるクラスのサンプルを他のクラスのすべてのサンプルから分離する特徴を作成したいだけです。
超平面を見つけるのではなく、特徴空間を「組織化」します。コントラスト損失により、類似した画像が何らかの形のクラスターに引き寄せられ、異なる画像が互いに遠ざけられます。したがって、データポイントとそれらの相互の相対位置を操作します。
類似性メトリクス (多くの場合、コサイン距離) [正規化され、温度スケール化された、つまり NT-Contrastive 損失] を学習します。
対照的なトレーニング目標は通常、隠れた表現を超球面上で均一になるように一致させようとします [1]。
次元崩壊の影響にさらされている[2]。これは、対照的なアプローチの特徴抑制 [3] によって引き起こされるか強化される可能性があります。以下も参照してください。
したがって、対照学習パラダイムは、類似したサンプルをクラスター化して、異なるサンプルのクラスターを遠ざけようとします。前者は BCE との明らかな違いです。
以下に沿った一般的な議論は、ビデオのエンコード/トランスコーディング (可逆または非可逆) がデータにアーティファクトを出現させ、一般にモデルがノイズとしてしか解釈できない情報が誘発されているという事実に基づいています。この関係にコンテキストを追加し、結論を導くための方向性を強調したいと思います。
視覚的なアイデアを得るために、BCE とコントラスト損失の両方に対する MNIST 埋め込みの典型的な 2D 表現を見てみましょう。
以下の 2 つの図は、BCE および対比損失トレーニング用の MNIST データセット上のそれぞれの埋め込み空間を示しています。これらは、上で箇条書きでまとめた内容を表していることがすでにわかります。

ポイント:
形状に制約を与えずに、あるクラスのサンプルを他のすべてのクラスのサンプルからできるだけ遠くに置くことを試みます。対して、コンパクトなクラスターを構築します (説明のために、図ではマルチクラス問題に戻っていることに注意してください。ここで得られた結論は、バイナリ問題の一般性を失うことなく依然として当てはまります。)。
BCE でトレーニングされた CNN モデルからの MNIST サンプルの埋め込み。
単なる説明例ではありますが、対照的なアプローチの図から、すでに何らかの帰結が得られます。ビデオ エンコーディングから余分なノイズが追加されていなくても、BCE と比較してクラスター間に見える重複がすでに多くなり、あらゆる種類のノイズに対してより脆弱になり、したがって分類の信頼性が変化します。
コントラスト損失でトレーニングされた CNN モデルからの MNIST サンプルの埋め込み。
結論として、異常または外れ値の検出に取り組んでいる場合、研究プロジェクトまたは製品に対する 1 つのヒントは、クラス内分散の最小化を利用するために、BCE 損失よりも対照損失を優先する可能性があるということです。
この実証例をもう少し詳しく見て、いくつかの関連出版物と、そこから得られる可能性のある答えを調べていきたいと思います。
考えられる説明のための関連作品
A. 非網羅的な機能スペース
[2] では、埋め込みベクトルの共分散行列の特異値分解である SimCLR [4] に似たトレーニング スキームを使用すると (ImageNet 画像上で 128 次元の埋め込みを使用します)、次元の 20% 以上が基本的に未使用のままとなり、次元の崩壊が示されることを示しています。必然的に、彼らは次のように推測します。
「強力な拡張を行うと、埋め込み空間共分散行列は低ランクになります。」
基本的に、モデルは特徴空間全体を使用していないため、容量が制限されます。

し、ノイズに対する感度が高まる可能性があります。
[2]より。コントラストネットワーク上の拡張の強度 k が増加すると、大量の次元が無関係になります。
私たちの作品では、対照的なアプローチは本質的に強力な拡張に基づいて構築されています。これによって引き起こされる可能性のある次元の削減は、モデルの実際の能力に影響を及ぼし、したがってトレーニング中には見えない入力内のノイズやアーティファクトを堅牢に補正する能力に影響を与えます。
B. 学習された特徴量の分布
[5] の著者らは、マージン損失やガウス混合などのさまざまな損失が特徴埋め込み空間に及ぼす影響を調査しています。彼らは、実際にクロスエントロピー損失がガウス混合分布に似た特徴空間分布をもたらすことを発見しました。クラスター内の特徴空間の構成。
彼らはさらに、ガウス混合損失を使用したトレーニングにより、教師ありの方法でクラスターの学習が可能になり、ネットワークの堅牢性が向上することを発見しました。具体的には、高い分類精度を持つディープ ニューラル ネットワークは、ソフトマックス/クロス エントロピー損失でトレーニングされた場合、敵対的な例に対して脆弱であることを示しています。 Fast Gradient Sign Method によって作成された敵対的な例は、画像にノイズを追加しているように見えます。

[切り捨てられた]

## Original Extract

We research and develop Machine Learning and AI models that are used to verify a clients’ identity and whether it is a live person that wants to onboard to their platform and use their services.

Supervised vs. Contrastive Classification: How Does the Video Codec Influence the Prediction Scores?
Sebastian Vater - Fellow researcher in ID&V
Why are we interested in different codecs and their effect on prediction scores anyway?
Where we work, at Fourthline, our task is to help to safely onboard clients for our customers, such as banks, fintechs, and non-financial businesses, in a KYC (Know Your Customer) process. We research and develop Machine Learning and AI models that are used to verify a client’s identity and whether it is a live person that wants to onboard to their platform and use their services.
While experimenting in our lab with Machine Learning systems, we observed something interesting: Models trained with different loss functions respond differently when we change how the input video is encoded. This opened up the question of whether (and, if so, to what extent) some training paradigms are more sensitive than others to different video codecs , and thus, to compression artifacts.
Now, why we would care to check what influence an encoding or transcoding has on the output of our services? We could just optimize for the codec at hand and be happy! Well, leaving academia behind, we interact in an open world, and so do our services. Fourthline serves a multitude of banks, fintechs, and non-financial businesses. Each of them individually preprocesses their data with their own data capture and backend engineering before running them in our algorithms to check the liveness and identity of a person. We do care that our services always behave as expected — no matter the input codec!
In this post, we want to explore specifically how the training objectives coming from two paradigms — supervised and contrastive learning (see [8] for an overview) — might have an effect on the created embedding space (e.g. structure or robustness against noise and artifacts). We want to take a closer look at some characteristics of these losses and see how these affect the classification behavior in real-world scenarios. In particular, the goal is to fathom possible influences of the different losses on the classification score in presence of image space artifacts, such as those resulting from different codecs being used for the input data.
Max Crous and Roman Aleksandrov of Fourthline’s Biometrics team conducted the experiments and created the graphs on which this blog post is based on. They furthermore shared valuable insights and ideas through discussions.
At Fourthline, our product architecture employs multiple integrated modules that work in concert to maximize their collective advantages. This modular approach allows our systems to address challenges through multiple parallel processing paths.
As one example, different parts of the liveness system are researched and trained following different learning paradigms and architectures. Doing so, among others, we make use of supervised and contrastive learning methods.
Since we are interested in the effect of different video codecs on the outcome of our algorithms, we started performing some experimental tests on data in our domain, to which we applied different encoding and transcoding preprocessing. Below we show some qualitative exemplary results for two particular systems. Both these systems are trained on the same data domain, and while one is being trained with a supervised binary cross entropy loss ( BCE ), the other one is trained with a contrastive loss ( Contrast. ), where a local density estimator is employed for estimating decision thresholds.
The plots below show the distribution of the difference in score outputs (e.g. Softmax) of the differently trained systems, when we change the input video codec. When the score on the original codec for a sample to be classified had any value x0 (defined by the horizontal line ‘0.00’), the colored dots show where the values of the confidence output for each of the test samples land when the model was fed videos with a different encoding, resulting in scores x0+ d x. The colored areas represent the spread of d x.
In these experiments, starting from the original MJPEG format (each frame in a video is encoded in JPEG format), we evaluate chroma subsampled MJPEG (4:4:4 and 4:2:0), intra-frame FFV1 (part of the FFmpeg project ). We also show below the comparison between the FFV1 and the lossless VP9 encoding.
Distribution of differences of classifier scores when changing the input codec. Image generated and owned by Fourthline.
So, what do we see? We immediately see that the model trained with a contrastive loss performs significantly different on the encoded data than the BCE trained model does. For VP9, assuming approximately a Gaussian statistic, the 1-sigma area contains score differences from +/- 0.05 — this means a change in confidence from, e.g., 0.8 to 0.75–0.85. This is a huge difference for a model output!
Distribution of differences of classifier scores when changing the input codec. Image generated and owned by Fourthline.
Exemplary visualizations of artifacts stemming from different video codecs. We compute pixel-wise differences between the original video and the encoded one. The results are to be interpreted qualitatively. From left to right: (1) Original image, a person trying to fool the Liveness system by applying a paper mask (no video compression). (2) Encoded with MJPEG. (3) Encoded with FFV1. (4) Encoded with VP9. Images owned by Fourthline.
You may be asking: Why does the contrastive model behave so differently? And why does the BCE model not? The Contrastive model seems to be much more vulnerable to image encodings/transcodings. We asked ourselves: What could be the reason and how can we prevent it?
The paragraph below will shine some light on possible explanations and provides some glances for reasons that might explain this behavior. (Note: Since this is just a blog post, scientifically sound proofs will only be published somewhere else.)
First, we see that divergence of scores is consistently high for the contrastive learning approach among different codecs. Given that all experiments were conducted on the same data domain, we are focusing on taking a closer look how the different learning schemes could explain this phenomenon.
In general, to prevent some undesired behavior in any technical system, it is most useful to understand what causes the observed behavior.
The Two Losses: Characteristics
A first possible explanation that comes to mind is that the learned feature spaces reveal different structures, in particular, different inter-class structures, for the two losses are different. Let us look at some characteristics about the losses (we assume a binary problem for our argumentation in this post):
Cross-Entropy binary loss (BCE):
BCE finds a hyperplane to separate the two classes, where it operates on class likelihoods (e.g. Softmax outputs).
There are no constraints on the model’s learned distribution (e.g. of class 1 and class 2 — they can assume any shape).
Learns separable features — any features, regardless of whether similar data points of the same class are close to each other (as long as they are ‘far away’ from all data points of the other class).
Summarizing, BCE does not care much about the intra-class distribution nor about the shape of the distribution itself at all — it just wants to create features that separate a sample of one class from all samples of the other class.
Rather than finding a hyperplane, it ‘organizes’ the feature space: The contrastive loss pulls similar images together into some form of a cluster while pushing dissimilar images away from each others. So it operates on data points and their relative positions to each other.
Learns a similarity metric (often Cosine distance) [that is normalized and temperature scaled, i.e. NT-Contrastive loss].
The contrastive training objective usually tries to match the hidden representation to be uniform on a hypersphere [1].
Is exposed to the effect of dimensional collapse [2]. This might be caused/enhanced by feature suppression [3] for the contrastive approach, see also below.
So, the contrastive learning paradigm wants to cluster similar samples, pushing clusters of dissimilar samples away — the former an apparent difference to the BCE.
Our general argument along the following is based on the fact that video encoding/transcoding — lossless or lossy — is causing artifacts to appear in the data, and, in general, information that a model can only interpret as noise, being induced. We want to add context to this relationship and highlight directions to draw conclusions.
To get a visual idea, let’s start by looking at typical 2D-representations of MNIST embedding for both, BCE and contrastive loss.
The two figures below illustrate respective embedding spaces on the MNIST dataset for BCE and contrastive loss training. We can already see that they represent what we summarized above in the bullet points:
Trying, without constraint on the shape, for a sample of one class to be as far away as possible to all the other classes’ samples — versus — building compact clusters (note, that, for illustrative purposes, we fall back to a multiclass problems in the figure, the conclusions made here still hold without loss of generality for the binary problem.).
Embedding for MNIST samples from a CNN model trained with BCE.
Though just an illustrative example, the figure for the contrastive approach lets us already draw some corollary: without extra noise being added from any video encoding, there is already more overlap visible between the clusters compared to the BCE, being more vulnerable to any kind of noise and therefore changing confidences in classification.
Embedding for MNIST samples from a CNN model trained with contrastive loss.
Concluding, if you are working on anomaly or outlier detection, one hint for your research project or product is that you might prefer a Contrastive Loss over a BCE loss, to make use of the minimization of in the intra class variance!
Going a bit deeper that this demonstrating example, we want to look into a few related publications and what possible answers they might yield.
Related works for possible explanations
A. Non-Exhaustive Feature Space
In [2] they show that using a training scheme similar to SimCLR [4], a singular value decomposition of the covariance matrix of embedding vectors (they use a 128 dimensional embedding on ImageNet images), more than 20% of the dimensions remain basically unused, showing the dimensional collapse. In a corollary, they deduce:
“With strong augmentation, the embedding space covariance matrix becomes low-rank.”
So basically, the model is not using the entire feature space, thus limiting capacity and possible increasing sensitivity to noise.
From [2]. With increasing strength k of augmentations on a contrastive network, a large amount of the dimensions become non-relevant.
In our works, the contrastive approach inherently is built upon strong augmentation. The dimensional reduction possibly caused by this affects the model’s actual capacity and thus, its capability to compensate robustly for noise and artifacts in the input unseen during training.
B. Learned Feature Distribution
The authors in [5] investigate the effect of different losses, including margin losses and Gaussian Mixtures on the feature embedding space. They find that indeed a Cross Entropy loss does yield a feature space distribution that resembles a Gaussian Mixture Distribution, e.g. an organization of the feature space in clusters.
They furthermore find that training with a Gaussian Mixture loss, that enables learning clusters in a supervised manner is increasing robustness of the network. Specifically, they show that deep neural networks with high classification accuracies are vulnerable to adversarial examples when trained with a Softmax/Cross entropy loss. Adversarial examples created by the Fast Gradient Sign Method can be seen as adding noise to the image in a co

[truncated]
