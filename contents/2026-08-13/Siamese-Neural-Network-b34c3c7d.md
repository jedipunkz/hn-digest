---
source: "https://en.wikipedia.org/wiki/Siamese_neural_network"
hn_url: "https://news.ycombinator.com/item?id=49289221"
title: "Siamese Neural Network"
article_title: "Siamese neural network - Wikipedia"
author: "Anon84"
captured_at: "2026-08-13T17:50:19Z"
capture_tool: "hn-digest"
hn_id: 49289221
score: 2
comments: 0
posted_at: "2026-08-13T17:29:53Z"
tags:
  - hacker-news
  - translated
---

# Siamese Neural Network

- HN: [49289221](https://news.ycombinator.com/item?id=49289221)
- Source: [en.wikipedia.org](https://en.wikipedia.org/wiki/Siamese_neural_network)
- Score: 2
- Comments: 0
- Posted: 2026-08-13T17:29:53Z

## Translation

タイトル: シャム ニューラル ネットワーク
記事のタイトル: シャム ニューラル ネットワーク - ウィキペディア

記事本文:
コンテンツへジャンプ
メインメニュー
メインメニュー
サイドバーに移動
隠す
ナビゲーション
メインページ
1
学習
Toggle 学習サブセクション
1.1
事前定義されたメトリック、ユークリッド距離メトリック
1.2
学習済みメトリクス、非線形距離メトリクス
1.3
学習済みメトリクス、ハーフツイン ネットワーク
2
オブジェクト追跡用のツインネットワーク
シャム ニューラル ネットワーク (ツイン ニューラル ネットワークとも呼ばれます) は、同じ重みを使用しながら 2 つの異なる入力ベクトルを連携して動作させ、同等の出力ベクトルを計算する人工ニューラル ネットワークです。 [ 1 ] [ 2 ] [ 3 ] 多くの場合、出力ベクトルの 1 つが事前に計算され、他の出力ベクトルと比較するためのベースラインが形成されます。これはフィンガープリントの比較に似ていますが、より技術的には局所性を考慮したハッシュの距離関数として説明できます。 [要出典]
機能的にはツイン ネットワークと同様ですが、わずかに異なる機能を実装するアーキテクチャを構築することが可能です。これは通常、異なるタイプ セット内の類似したインスタンスを比較するために使用されます。 [要出典]
ツイン ネットワークが使用される類似性測定の用途としては、手書き小切手の認識、カメラ画像内の顔の自動検出、テキスト クエリとインデックス付きドキュメントの照合などが挙げられます [4] 。おそらくツイン ネットワークの最もよく知られたアプリケーションは顔認識です。顔認識では、既知の人物の画像が事前に計算され、回転式改札口などからの画像と比較されます。最初はわかりませんが、わずかに異なる 2 つの問題があります。 1 つは、多数の他人の中から人を認識すること、つまり顔認識の問題です。 DeepFace はそのようなシステムの一例です。 [ 3 ] 最も極端な形式では、これは駅や空港で 1 人の人物を認識することです。もう一つは顔認証です

、それはたとえば、パスポートの写真がパスポートの所有者の顔と一致するかどうかを確認することです。ツイン ネットワークは同じかもしれませんが、実装はまったく異なる場合があります。
ツイン ネットワークでの学習は、三重項損失または対照的損失を使用して実行できます。三重項損失による学習では、ベースライン ベクトル (アンカー イメージ) が正のベクトル (真実のイメージ) および負のベクトル (偽のイメージ) と比較されます。負のベクトルはネットワーク内で学習を強制しますが、正のベクトルは正則化のように機能します。対比損失による学習の場合、重みを正規化するための重み減衰、または正規化などの同様の操作が必要です。
損失関数の距離メトリックには次の特性があります [5]
非否定性:
δ
(
×
、
y
)
≥
0
{\displaystyle \delta (x,y)\geq 0}
識別できない人の身元:
δ
(
×
、
y
)
=
0
⟺
×
=
y
{\displaystyle \delta (x,y)=0\iff x=y}
可換性:
δ
(
×
、
y
)
=
δ
(
y
、
×
)
{\displaystyle \delta (x,y)=\delta (y,x)}
三角不等式 :
δ
(
×
、
z
)
≤
δ
(
×
、
y
)
+
δ
(
y
、
z
)
{\displaystyle \delta (x,z)\leq \delta (x,y)+\delta (y,z)}
特に、三重項損失アルゴリズムは、多くの場合、二乗ユークリッド (ユークリッドとは異なり、三角不等式を持たない) 距離を中心に定義されます。
事前定義されたメトリック、ユークリッド距離メトリック
共通の学習目標は、類似したオブジェクトの距離メトリックを最小化し、異なるオブジェクトの距離メトリックを最大化することです。これにより、次のような損失関数が得られます。
使用される最も一般的な距離メトリックはユークリッド距離です。この場合、損失関数は次のように行列形式で書き直すことができます。
学習済みメトリクス、非線形距離メトリクス
より一般的なケースは、ツイン ネットワークからの出力ベクトルが、非線形距離メトリックを実装する追加のネットワーク層を通過する場合です。

。
行列形式では、前のものは線形空間のマハラノビス距離として [6] のように近似されることがよくあります。
これは、少なくとも教師なし学習と教師あり学習にさらに細分化できます。
学習済みメトリクス、ハーフツイン ネットワーク
この形式では、ツイン ネットワークを半分のツインにすることもでき、わずかに異なる機能を実装します。
オブジェクト追跡用のツインネットワーク
ツイン ネットワークは、そのユニークな 2 つのタンデム入力と類似性測定により、オブジェクト追跡に使用されてきました。オブジェクト追跡では、ツイン ネットワークの 1 つの入力はユーザーが事前に選択したサンプル画像であり、もう 1 つの入力はより大きな検索画像です。ツイン ネットワークの仕事は、検索画像内の見本を見つけることです。見本と検索画像の各部分との類似性を測定することで、ツインネットワークにより類似性スコアのマップを与えることができます。さらに、完全畳み込みネットワークを使用すると、各セクターの類似性スコアを計算するプロセスを 1 つの相互相関層のみで置き換えることができます。 [ 7 ]
Twin 完全畳み込みネットワークは、2016 年に初めて導入されて以来、多くの高性能リアルタイム オブジェクト追跡ニューラル ネットワークで使用されてきました。 Like CFnet、[ 8 ] StructSiam、[ 9 ] SiamFC-tri、[ 10 ] DSiam、[ 11 ] SA-Siam、[ 12 ] SiamRPN、[ 13 ] DaSiamRPN、[ 14 ] Cascaded SiamRPN、[ 15 ] SiamMask、[ 16 ] SiamRPN++、[ 17 ] Deeperおよびより広いサイアムRPN。 [ 18 ]
↑ ジェーン・ブロムリー。ギヨン、イザベル。ルカン、ヤン。ゼッキンガー、エドゥアルド;シャー、ルーパック (1994)。 「「シャム」時間遅延ニューラル ネットワークを使用した署名検証」 (PDF) 。神経情報処理システムの進歩。 6 : 737–744。
↑ チョプラ、S.ハドセル、R. LeCun、Y. (2005 年 6 月)。 「顔認証への応用による類似性メトリクスの識別学習」。 2005 IEEE Computer Society Conference on Computer Vi

イオンとパターン認識 (CVPR'05) 。 Vol. 1. 539 ～ 546 頁1. 土井：10.1109/CVPR.2005.202。 ISBN 0-7695-2372-2 。 S2CID 5555257 。
1 2 タイグマン、Y.ヤン、M.ランザト、M.ウルフ、L. (2014 年 6 月)。 「DeepFace: 顔認証における人間レベルのパフォーマンスとのギャップを埋める」。 2014 年のコンピュータ ビジョンとパターン認識に関する IEEE カンファレンス。 pp. 1701–1708。土井：10.1109/CVPR.2014.220。 ISBN 978-1-4799-5118-5 。 S2CID 2814088 。
↑ 「文の類似性を学習するためのシャム再帰アーキテクチャ」 .
↑ チャタジー、モイトレヤ。羅、雲南。 「畳み込みニューラル ネットワークを使用した (または使用しない) 類似性学習」 (PDF) 。 2018年12月7日に取得。
↑ チャンドラ国会議員（1936年）。 「統計における一般化された距離について」 (PDF) 。インド国立科学研究所の議事録。 1. 2 : 49–55.
↑ オブジェクト追跡のための完全畳み込みシャム ネットワーク arXiv : 1606.09549
↑ 「相関フィルターベースの追跡のためのエンドツーエンドの表現学習」 。
↑ 「リアルタイム視覚追跡のための構造化シャム ネットワーク」 (PDF) 。
↑ 「物体追跡のためのシャムネットワークにおけるトリプレット損失」 (PDF) 。
↑ 「視覚オブジェクト追跡のための動的シャムネットワークの学習」 (PDF) 。
↑ 「リアルタイム オブジェクト追跡のための 2 つのシャム ネットワーク」 (PDF) 。
↑ 「シャム地域提案ネットワークによる高性能視覚追跡」 (PDF) 。
↑ 朱、鄭；王、強。リー、ボー。呉、魏。ヤン・ジュンジエ。胡偉明（2018）。 「視覚オブジェクト追跡のためのディストラクター対応シャム ネットワーク」。 arXiv : 1808.06048 [ cs.CV ]。
↑ ファン、ヘン;リン・ハイビン（2018）。 「リアルタイム視覚追跡のためのシャム カスケード領域提案ネットワーク」。 arXiv : 1812.06148 [ cs.CV ]。
↑ 王、強。張、李。ベルティネット、ルカ。胡、偉明。トール、フィリップ H.S. (2018)。 「高速オンラインオブジェクト追跡とセグメンテーション: 統一アプローチ」。 arXiv

: 1812.05050 [ cs.CV ]。
↑ リー、ボー。呉、魏。王、強。張、方儀。シン、ジュンリャン。ヤン・ジュンジエ (2018) 「SiamRPN++: 非常に深いネットワークによるシャム ビジュアル トラッキングの進化」。 arXiv : 1812.11703 [ cs.CV ]。
↑ 張志鵬。彭、侯文（2019）。 「リアルタイム視覚追跡のためのより深くより広いシャム ネットワーク」。 arXiv : 1901.01660 [ cs.CV ]。
「 https://en.wikipedia.org/w/index.php?title=Siamese_neural_network&oldid=1335674585 」から取得
カテゴリ : ニューラル ネットワーク アーキテクチャ
短い説明のある記事
短い説明はウィキデータとは異なります
出典のない記述を含むすべての記事
2019 年 9 月の出典のない記述を含む記事
このページは、2026 年 1 月 30 日の 15:14 (UTC) に最後に編集されました。
ページは Parsoid でレンダリングされました。
テキストは、クリエイティブ コモンズ 表示 - 継承 4.0 ライセンスに基づいて利用できます。
追加の条件が適用される場合があります。このサイトを使用すると、利用規約とプライバシー ポリシーに同意したことになります。 Wikipedia® は、非営利団体である Wikimedia Foundation, Inc. の登録商標です。

## Original Extract

Jump to content
Main menu
Main menu
move to sidebar
hide
Navigation
Main page
1
Learning
Toggle Learning subsection
1.1
Predefined metrics, Euclidean distance metric
1.2
Learned metrics, nonlinear distance metric
1.3
Learned metrics, half-twin networks
2
Twin networks for object tracking
A Siamese neural network (sometimes called a twin neural network ) is an artificial neural network that uses the same weights while working in tandem on two different input vectors to compute comparable output vectors. [ 1 ] [ 2 ] [ 3 ] Often one of the output vectors is precomputed, thus forming a baseline against which the other output vector is compared. This is similar to comparing fingerprints but can be described more technically as a distance function for locality-sensitive hashing . [ citation needed ]
It is possible to build an architecture that is functionally similar to a twin network but implements a slightly different function. This is typically used for comparing similar instances in different type sets. [ citation needed ]
Uses of similarity measures where a twin network might be used are such things as recognizing handwritten checks, automatic detection of faces in camera images, and matching text queries with indexed documents [ 4 ] . The perhaps most well-known application of twin networks are face recognition , where known images of people are precomputed and compared to an image from a turnstile or similar. It is not obvious at first, but there are two slightly different problems. One is recognizing a person among a large number of other persons, that is the facial recognition problem. DeepFace is an example of such a system. [ 3 ] In its most extreme form this is recognizing a single person at a train station or airport. The other is face verification , that is for example, to verify whether a photo in a passport matches the face of the passport's owner. The twin network might be the same, but the implementation can be quite different.
Learning in twin networks can be done with triplet loss or contrastive loss . For learning by triplet loss a baseline vector (anchor image) is compared against a positive vector (truthy image) and a negative vector (falsy image). The negative vector will force learning in the network, while the positive vector will act like a regularizer. For learning by contrastive loss there must be a weight decay to regularize the weights, or some similar operation like a normalization.
A distance metric for a loss function may have the following properties [ 5 ]
Non-negativity:
δ
(
x
,
y
)
≥
0
{\displaystyle \delta (x,y)\geq 0}
Identity of Non-discernibles:
δ
(
x
,
y
)
=
0
⟺
x
=
y
{\displaystyle \delta (x,y)=0\iff x=y}
Commutativity:
δ
(
x
,
y
)
=
δ
(
y
,
x
)
{\displaystyle \delta (x,y)=\delta (y,x)}
Triangle inequality :
δ
(
x
,
z
)
≤
δ
(
x
,
y
)
+
δ
(
y
,
z
)
{\displaystyle \delta (x,z)\leq \delta (x,y)+\delta (y,z)}
In particular, the triplet loss algorithm is often defined with squared Euclidean (which unlike Euclidean, does not have triangle inequality) distance at its core.
Predefined metrics, Euclidean distance metric
The common learning goal is to minimize a distance metric for similar objects and maximize for distinct ones. This gives a loss function like
The most common distance metric used is Euclidean distance , in case of which the loss function can be rewritten in matrix form as
Learned metrics, nonlinear distance metric
A more general case is where the output vector from the twin network is passed through additional network layers implementing non-linear distance metrics.
On a matrix form the previous is often approximated as a Mahalanobis distance for a linear space as [ 6 ]
This can be further subdivided in at least Unsupervised learning and Supervised learning .
Learned metrics, half-twin networks
This form also allows the twin network to be more of a half-twin, implementing a slightly different functions
Twin networks for object tracking
Twin networks have been used in object tracking because of its unique two tandem inputs and similarity measurement. In object tracking, one input of the twin network is user pre-selected exemplar image, the other input is a larger search image. The twin network's job is to locate the exemplar inside of the search image. By measuring the similarity between exemplar and each part of the search image, a map of similarity score can be given by the twin network. Furthermore, using a Fully Convolutional Network, the process of computing each sector's similarity score can be replaced with only one cross correlation layer. [ 7 ]
After being first introduced in 2016, Twin fully convolutional network has been used in many High-performance Real-time Object Tracking Neural Networks. Like CFnet, [ 8 ] StructSiam, [ 9 ] SiamFC-tri, [ 10 ] DSiam, [ 11 ] SA-Siam, [ 12 ] SiamRPN, [ 13 ] DaSiamRPN, [ 14 ] Cascaded SiamRPN, [ 15 ] SiamMask, [ 16 ] SiamRPN++, [ 17 ] Deeper and Wider SiamRPN. [ 18 ]
↑ Bromley, Jane; Guyon, Isabelle; LeCun, Yann; Säckinger, Eduard; Shah, Roopak (1994). "Signature verification using a "Siamese" time delay neural network" (PDF) . Advances in Neural Information Processing Systems . 6 : 737– 744.
↑ Chopra, S.; Hadsell, R.; LeCun, Y. (June 2005). "Learning a Similarity Metric Discriminatively, with Application to Face Verification". 2005 IEEE Computer Society Conference on Computer Vision and Pattern Recognition (CVPR'05) . Vol. 1. pp. 539–546 vol. 1. doi : 10.1109/CVPR.2005.202 . ISBN 0-7695-2372-2 . S2CID 5555257 .
1 2 Taigman, Y.; Yang, M.; Ranzato, M.; Wolf, L. (June 2014). "DeepFace: Closing the Gap to Human-Level Performance in Face Verification". 2014 IEEE Conference on Computer Vision and Pattern Recognition . pp. 1701– 1708. doi : 10.1109/CVPR.2014.220 . ISBN 978-1-4799-5118-5 . S2CID 2814088 .
↑ "Siamese Recurrent Architectures for Learning Sentence Similarity" .
↑ Chatterjee, Moitreya; Luo, Yunan. "Similarity Learning with (or without) Convolutional Neural Network" (PDF) . Retrieved 2018-12-07 .
↑ Chandra, M.P. (1936). "On the generalized distance in statistics" (PDF) . Proceedings of the National Institute of Sciences of India . 1. 2 : 49– 55.
↑ Fully-Convolutional Siamese Networks for Object Tracking arXiv : 1606.09549
↑ "End-to-end representation learning for Correlation Filter based tracking" .
↑ "Structured Siamese Network for Real-Time Visual Tracking" (PDF) .
↑ "Triplet Loss in Siamese Network for Object Tracking" (PDF) .
↑ "Learning Dynamic Siamese Network for Visual Object Tracking" (PDF) .
↑ "A Twofold Siamese Network for Real-Time Object Tracking" (PDF) .
↑ "High Performance Visual Tracking with Siamese Region Proposal Network" (PDF) .
↑ Zhu, Zheng; Wang, Qiang; Li, Bo; Wu, Wei; Yan, Junjie; Hu, Weiming (2018). "Distractor-aware Siamese Networks for Visual Object Tracking". arXiv : 1808.06048 [ cs.CV ].
↑ Fan, Heng; Ling, Haibin (2018). "Siamese Cascaded Region Proposal Networks for Real-Time Visual Tracking". arXiv : 1812.06148 [ cs.CV ].
↑ Wang, Qiang; Zhang, Li; Bertinetto, Luca; Hu, Weiming; Torr, Philip H. S. (2018). "Fast Online Object Tracking and Segmentation: A Unifying Approach". arXiv : 1812.05050 [ cs.CV ].
↑ Li, Bo; Wu, Wei; Wang, Qiang; Zhang, Fangyi; Xing, Junliang; Yan, Junjie (2018). "SiamRPN++: Evolution of Siamese Visual Tracking with Very Deep Networks". arXiv : 1812.11703 [ cs.CV ].
↑ Zhang, Zhipeng; Peng, Houwen (2019). "Deeper and Wider Siamese Networks for Real-Time Visual Tracking". arXiv : 1901.01660 [ cs.CV ].
Retrieved from " https://en.wikipedia.org/w/index.php?title=Siamese_neural_network&oldid=1335674585 "
Category : Neural network architectures
Articles with short description
Short description is different from Wikidata
All articles with unsourced statements
Articles with unsourced statements from September 2019
This page was last edited on 30 January 2026, at 15:14 (UTC) .
Page was rendered with Parsoid .
Text is available under the Creative Commons Attribution-ShareAlike 4.0 License ;
additional terms may apply. By using this site, you agree to the Terms of Use and Privacy Policy . Wikipedia® is a registered trademark of the Wikimedia Foundation, Inc. , a non-profit organization.
