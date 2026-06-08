---
source: "https://gmcirco.github.io/blog/posts/ai-calibration/calibration.html"
hn_url: "https://news.ycombinator.com/item?id=48446816"
title: "How Confident Are AI Classifiers About Their Own Confidence?"
article_title: "How Confident are AI Classifiers About Their Own Confidence? – A Blog for Data Stuff"
author: "apwheele"
captured_at: "2026-06-08T16:28:24Z"
capture_tool: "hn-digest"
hn_id: 48446816
score: 2
comments: 1
posted_at: "2026-06-08T15:41:07Z"
tags:
  - hacker-news
  - translated
---

# How Confident Are AI Classifiers About Their Own Confidence?

- HN: [48446816](https://news.ycombinator.com/item?id=48446816)
- Source: [gmcirco.github.io](https://gmcirco.github.io/blog/posts/ai-calibration/calibration.html)
- Score: 2
- Comments: 1
- Posted: 2026-06-08T15:41:07Z

## Translation

タイトル: AI 分類器は自分自身の自信についてどの程度自信を持っていますか?
記事のタイトル: AI 分類子は自分自身の自信についてどの程度自信を持っていますか? – データに関するブログ

記事本文:
データに関するブログ
私について
AI 分類子は自分自身の自信についてどの程度自信を持っていますか?
2024 NEISS を使用した傷害の分類
AI「信頼性」スコア
「自信がある」とはどれくらい自信があるのでしょうか？
NEISSから一次損傷を抽出
抽出スキーマ
LLM は、さまざまな分野でテキスト分類子として多くの実際的な用途が発見されています。私の日常の仕事では、ルーティングまたは処理のためにドキュメントをさまざまなカテゴリに分類および分類するためにこれらをよく使用します。これらは、自然言語処理のための前世代の BERT 風のベースのニューラル ネットワーク モデルに大きく取って代わったと言えます。実際、LLM は通常、必要な指示や事前トレーニングがかなり最小限で、さまざまな分類タスクを非常に得意とします。
ただし、LLM から分類の「確率」を取得することは、従来の機械学習やニューラル ネットワークのアプローチよりも少し困難です。 LLM には、分類確率を取得するための類似の fit.predict_proba(X) がありません。 LLM 分類モデルに対してこれが行われるのを私が見たのは、実際には 2 つの方法です。大まかに言うと、次のとおりです。
LLM に分類の信頼度を推定し、結果を出力で返すように要求します。
モデル出力からトークンレベルの確率を直接抽出します。
前者は間違いなく私が見た中で最も一般的なものですが、後者はもう少し直接的で珍しいものです (ただし、OpenAI によって直接サポートされています!)。私の友人のアンディ・ウィーラーは、彼の著書 (Wheeler 2026) の中で後者を実行する良い例を示しています。これを念頭に置いて、この投稿で私がやりたかったことは、AI によって生成された信頼スコアが実際にどの程度現実に近いのかを少し調査したいと思いました。
NEISSから一次損傷を抽出
そのためのデータソースとして

テストでは、National Electronic Injury Surveillance System (NEISS) を使用しました。これは、短い医学的説明を含むラベル付きの数十万の例が含まれているため、抽出スキームをテストするのに役立つ情報源です。それらのほとんどは次のようになります。
「26YOMは昨日右足で釘を踏み、痛みと腫れがあると述べています。DX：右足の刺し傷です。」
テストの目的で、医療ナラティブをレビューし、ナラティブに記載されている主な傷害を分類するための LLM ベースの抽出パイプラインを作成しました。 gpt-5-nano を LLM として使用して、500 のナラティブのサンプルに対してバッチ プロセスを実行しました。
興味があれば、コードの完全なセットはここにあります。本当に重要な部分は、抽出スキーマとプロンプトだけです。残りは OpenAI のバッチ プロセスにルーティングするだけです。プロンプトには、身体部分の NEISS コーディング マニュアルの関連セクションにあるルールの完全なセットを提供します。さらに、いくつかの実際の値 (いわゆる「少数ショット」または「k ショット」サンプル) を使用して AI を基礎づけるのに役立つ、ラベル付きサンプルのサンプルを追加します。一般に、抽出セットアップは非常に最小限で済みます。以下に抽出スキーマを示します。完全なプロンプトはここにあります。
クラス BodyPart( str , Enum):
内部 = "0"
ショルダー = "30"
UPPER_TRUNK = "31"
肘 = "32"
LOWER_ARM = "33"
手首 = "34"
膝 = "35"
LOWER_LEG = "36"
足首 = 「37」
PUBIC_REGION = "38"
ヘッド = "75"
フェイス = "76"
目玉 = "77"
LOWER_TRUNK = "79"
UPPER_ARM = "80"
UPPER_LEG = "81"
ハンド = "82"
フット = "83"
BODY_25_50_PERCENT = "84"
ALL_PARTS_BODY = "85"
NOT_STATED_UNKNOWN = "87"
口 = "88"
ネック = "89"
指 = "92"
TOE = "93"
EAR = "94"
クラス InjuryClassification(BaseModel):
body_part: BodyPart = フィールド(
description = "負傷した、または物語に関与した身体の主要部分の ID コード"

)
推論: str = フィールド(
description = "体の一部が選ばれた理由の 10 ワード以内の説明"
)
信頼度: float = フィールド(
説明 = (
「身体部位分類の信頼度は 0.0 から 1.0 までの浮動小数点数です。」
「1.0 は絶対的な確実性を示し、0.0 は完全に不明です。」
)
)
分類確率
目標は、以下のような出力を取得することです。
{ "body_part" : 88 , "reasoning" : "診断では上唇裂傷と診断されています。 、「信頼度」: 0.98 }
私の「自信」の定義は、意図的にかなり曖昧にしてあります。 AIには「絶対確実」「全く不明」の範囲内に収めるように指示するだけです。私の経験では、信頼スコアを使用するアプリケーションのほとんどは、これよりもさらに少ない構造を適用しています。
トークンの確率については、応答ペイロードから直接抽出し、以下のような形式で上位 5 つを取得します。
"top_logprobs": [
{
"トークン" : "88" ,
"logprob" : -0.0009117019944824278
} 、
{
"トークン" : "76" ,
"ログプロブ" : -7.219661712646484
} 、
{
"トークン" : "80" ,
"ログプロブ": -10.270442962646484
} 、
{
"トークン" : "30" ,
「ログプロブ」: -10.481380462646484
} 、
{
"トークン" : "75" ,
「ログプロブ」: -10.715755462646484
}
]
私たちの分類は数値文字列であるため、体の各部分は一意の文字トークンにマップされます。したがって、この例では、最も高い確率のトークンは「口」の「88」で、確率は 0.999 です。次に確率が高いトークンは「顔」の「76」です。技術的には、上位 20 または 30 (または設定されている top_k の値) のトークンの確率を返すことができますが、実際には、選択されたトークンの確率は 0.999 以上に大きく偏っており、通常、返される上位 5 つのトークンは確率の点で 1 に近くなります。
モデルを実行すると、86% というかなり立派な精度が得られました。私は持っています

より詳細なプロンプトと k-shot の例によって、これを簡単に 95% 以上にできることは間違いありません。
これらの信頼スコアの最大の欠点は、他の多くの分類子にも共通しています。それは、それらが十分に調整されている可能性が低いことです。これは、モデルから返された確率が正しい分類の観察された分布に厳密にマッピングされるという保証がないことを意味します。たとえば、適切にキャリブレーションされたモデルの 95% の確率では、95% の確率で正しい分類が返されることが期待されます。適切に調整された予測は、特定のビジネス KPI を簡単に追跡できるため便利です。
500 の分類の信頼スコアの概要をすぐに取得できます。以下に、AI が生成した信頼スコアとトークン確率の両方の校正プロットと表を示します。 2 つのアプローチを見ると、いくつかの顕著な違いがあります。上限では、AI が生成した信頼スコアは実際に非常に現実に近いものになります。明らかな発散があるのは上限の外側のみです。たとえば、AI は 7 つのケースについて平均信頼度を約 45% と推定しますが、これは 100% 正確です。
どちらも最高上限以外ではあまり適切に調整されていませんが、トークンの対数確率は AI が生成したものと比較して非常に過信されています。
調整された確率を生成したい場合、マルチクラス シナリオではバイナリ シナリオと比べてアプローチが若干異なります。いくつかのアプローチがありますが、最も単純なのは提案された「トップ対すべて」アプローチのようです (Le Coz、Herbin、および Adjed 2024)。要するに、これは実際には、最も高い確率の正確さに基づいて調整するバイナリの場合を一般化したものにすぎません。考えられるカテゴリが多数ある場合 (ここでは 20 以上)、それは難しく、可能性が高くなります。

各カテゴリを直接調整するのは役に立ちません。むしろ、モデルを使用してプライマリ トークンの確率を調整します。これを行う簡単な方法は、等張回帰を使用することです。これは、確率にマッピングし直す厳密に増加するステップ関数を適用します。
# 正しいラベルとトークンの確率を取得する
正しい <- res $ is_correct
pred_prob <- res $ body_part_logprob_prob_1
# トップ対すべてのキャリブレーションを実行する
# アイソトニックステップ関数を使用する
tva_iso <- isoreg (pred_prob、正しい)
tva_predict <- as.stepfun (tva_iso)
calibrated_conf_iso <- tva_predict (pred_prob)
# .85 のトークン確率が何にマッピングされるかを確認する
print ( tva_predict (. 85 ))
[1] 0.6140351
これで、モデルが観測された精度に対応するビンに確率を再マップしていることがわかります。たとえば、元のトークンの確率が 0.85 の場合、調整された確率は 0.61 になります。以下のプロットは、元の確率がキャリブレーションされたビンにどのように再マッピングされるかを示しています。実際の運用環境では、ケースのサンプルを使用して校正モデルを構築し、これをホールドアウト セットで検証し、それを将来の分類に適用します。

## Original Extract

A Blog for Data Stuff
About Me
How Confident are AI Classifiers About Their Own Confidence?
Injury classsification using the 2024 NEISS
AI ‘Confidence’ Scores
How confident is “confident”?
Extracting Primary Injuries from the NEISS
Extraction Schema
LLMs have found a lot of practical uses as text classifiers across a ton of different areas. In my day job we commonly use them to sort and classify documents into different categories for routing or processing. I would say these have largely supplanted the previous generation of BERT-esque based neural network models for natural language processing. Indeed, LLMs are typically quite good at a variety of classification tasks with fairly minimal instructions or pre-training required.
However, getting classification “probabilities” from an LLM is a bit more challenging than from conventional machine learning or neural network approaches. LLMs don’t have an analogous fit.predict_proba(X) to retrieve classification probabilities. There are really two ways that I have seen this done for LLM classification models. Broadly speaking, they are:
Prompt the LLM to estimate its confidence in the classification and return the result in the output
Directly extract token-level probabilities from the model output
The first is definitely the most common one I’ve seen, while the latter is slightly more direct and more unusual (but directly supported by OpenAI!). My friend, Andy Wheeler, has a good example of doing the latter in his book ( Wheeler 2026 ) . With this in mind, what I wanted to do for this post was to look into what For this post, I wanted to explore a bit about how close to reality the AI-generated confidence scores actually are.
Extracting Primary Injuries from the NEISS
As a source of data for this test, I used the National Electronic Injury Surveillance System (NEISS). This is a useful source of information for testing extraction schemes because they contain hundreds of thousands of labeled examples with short medical narratives. Most of them look akin to this:
“26YOM STEPPED ON A NAIL WITH HIS RIGHT FOOT YESTERDAY AND STATES IT IS PAINFUL AND SWOLLEN DX: PUNCTURE WOUND RIGHT FOOT”
For my testing purposes, I wrote an LLM-based extraction pipeline to review the medical narrative and then classify the primary injury described in the narrative. I ran a batch process on a sample of 500 narratives using gpt-5-nano as the LLM.
If you’re interested, the full set of code is here . The real important parts are just the extraction schema and the prompt. The rest is just routing stuff to OpenAI’s batch process. For the prompt I provide the full set of rules from the relevant section in the NEISS coding manual for body parts. In addition, I add a sampling of labeled examples to help ground the AI with some real values (so called “few shot” or “k-shot” examples). In general, the extraction setup is quite minimal. Below I include the extraction schema, and the full prompt is here .
class BodyPart( str , Enum):
INTERNAL = "0"
SHOULDER = "30"
UPPER_TRUNK = "31"
ELBOW = "32"
LOWER_ARM = "33"
WRIST = "34"
KNEE = "35"
LOWER_LEG = "36"
ANKLE = "37"
PUBIC_REGION = "38"
HEAD = "75"
FACE = "76"
EYEBALL = "77"
LOWER_TRUNK = "79"
UPPER_ARM = "80"
UPPER_LEG = "81"
HAND = "82"
FOOT = "83"
BODY_25_50_PERCENT = "84"
ALL_PARTS_BODY = "85"
NOT_STATED_UNKNOWN = "87"
MOUTH = "88"
NECK = "89"
FINGER = "92"
TOE = "93"
EAR = "94"
class InjuryClassification(BaseModel):
body_part: BodyPart = Field(
description = "ID code for the primary body part injured or involved in the narrative"
)
reasoning: str = Field(
description = "10 word or less description of why the body part was chosen"
)
confidence: float = Field(
description = (
"Confidence of the body part classification as a float from 0.0 to 1.0, "
"where 1.0 is absolute certainty, and 0.0 is completely unknown"
)
)
Classification probabilities
The goal is to get output that looks like this below:
{ "body_part" : 88 , "reasoning" : "Diagnosis states upper lip laceration." , "confidence" : 0.98 }
My definition of “confidence” is intentionally kept fairly vague. I just tell the AI to keep it in the range of “absolute certainty” and “completely unknown”. In my experience, most applications I have seen using confidence scores apply even less structure than this.
For token probabilities, I extract them directly from the response payload, and get the top 5 in a format like this below:
"top_logprobs": [
{
"token" : "88" ,
"logprob" : -0.0009117019944824278
} ,
{
"token" : "76" ,
"logprob" : -7.219661712646484
} ,
{
"token" : "80" ,
"logprob" : -10.270442962646484
} ,
{
"token" : "30" ,
"logprob" : -10.481380462646484
} ,
{
"token" : "75" ,
"logprob" : -10.715755462646484
}
]
Because our classifications are numeric strings, each body part maps to a unique character token. So for this example, the highest probability token is “88” for “Mouth” with a probability of 0.999. The next-highest probability token is “76” for “Face”. Technically we could return the token probabilities for the top 20 or 30 (or whatever value of top_k is set), but in practice the probabilities for the chosen token are highly biased toward 0.999+, and typically the highest 5 tokens returned are close to 1 in terms of probabilities.
After running the model we get a fairly respectable accuracy of 86% . I have no doubt that some more detailed prompting and k-shot examples could easily get this north of 95%.
The biggest drawback with these confidence scores is shared with many other classifiers: they are unlikely to be well calibrated . This means that there is no guarantee that the probabilities returned from the model closely map to the observed distribution of correct classifications. For example, a 95% probability in a well-calibrated model should expect to return a correct classification 95% of the time. Well-calibrated predictions are useful because they can help easily track to specific business KPIs.
We can quickly get a summary of the confidence scores for the 500 classifications. Below, I have a calibration plot and table for both the AI-generated confidence scores, and the token probability. Looking at the two approaches, there are some notable differences. On the upper end the AI-generated confidence scores are actually really quite close to reality. It’s only outside the upper end that there is clear divergence. For example, the AI estimates a mean confidence of about 45% for 7 cases, but is 100% correct!
While neither is terribly well-calibrated outside the highest upper-end, the token log-probabilities are highly over-confident relative to the AI-generated ones.
If we want to generate calibrated probabilities, the approach is slightly different in the multi-class scenario compared to the binary one. There are a few approaches, but the simplest seems to be a proposed “top versus all” approach ( Le Coz, Herbin, and Adjed 2024 ) . In short, this is really just a generalization of the binary case where we calibrate based on the correctness of the highest probability. In cases where we have many possible categories (here, over 20), it is difficult and probably not useful to directly calibrate each category. Rather, we use a model to calibrate the primary token probability. A simple way to do this is use isotonic regression, which applies a strictly increasing step function that maps back to the probabilities:
# get the correct labels and token probabilities
correct <- res $ is_correct
pred_prob <- res $ body_part_logprob_prob_1
# Do a top-versus-all calibration
# using isotonic step function
tva_iso <- isoreg (pred_prob, correct)
tva_predict <- as.stepfun (tva_iso)
calibrated_conf_iso <- tva_predict (pred_prob)
# check what a token probability of .85 now maps to
print ( tva_predict (. 85 ))
[1] 0.6140351
Now we see that the model re-maps probabilities to bins corresponding to observed accuracy. For example, an original token probability of .85 returns a calibrated probability of .61. The plot below shows how the original probabilities are remapped to a calibrated bins. In an actual production environment I would use a sample of cases to build the calibration model, validate this on a hold-out set, then apply it on future classifications.
