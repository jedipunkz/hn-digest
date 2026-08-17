---
source: "https://magazine.sebastianraschka.com/p/ai-detector-from-scratch"
hn_url: "https://news.ycombinator.com/item?id=49338761"
title: "Building an AI Text Detector from Scratch"
article_title: "Building an AI Text Detector From Scratch"
image: "https://substackcdn.com/image/fetch/$s_!Om23!,w_1200,h_675,c_fill,f_jpg,q_auto:good,fl_progressive:steep,g_auto/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F875b5076-19a9-4ffe-91d4-cafa156f650a_5343x5696.png"
author: "ibobev"
captured_at: "2026-08-17T23:14:01Z"
capture_tool: "hn-digest"
hn_id: 49338761
score: 1
comments: 0
posted_at: "2026-08-17T22:50:03Z"
tags:
  - hacker-news
  - translated
---

# Building an AI Text Detector from Scratch

- HN: [49338761](https://news.ycombinator.com/item?id=49338761)
- Source: [magazine.sebastianraschka.com](https://magazine.sebastianraschka.com/p/ai-detector-from-scratch)
- Score: 1
- Comments: 0
- Posted: 2026-08-17T22:50:03Z

## Translation

タイトル: AI テキスト検出器をゼロから構築する
記事のタイトル: AI テキスト検出器をゼロから構築する
説明: データセットの構築、モデルのトレーニング、ローカル展開、RLVR を含むエンドツーエンドのプロジェクト

記事本文:
AI テキスト検出器をゼロから構築する
AI テキスト検出器をゼロから構築する
データセットの構築、モデルのトレーニング、ローカル展開、RLVR を含むエンドツーエンドのプロジェクト
Sebastian Raschka 博士 Aug 15, 2026 ∙ 有料 26 3 シェア Substack は最近、UI で AI 検出機能を開始しましたが、これは非常に興味深いものです。
これとは別に、小規模言語モデル (SLM) がどのような機能を備えているかを示すためのデモとして、興味深いローカルの DIY LLM プロジェクトについて多くの人が私に尋ねてきました。
1 つと 1 つを組み合わせて、AI 検出器をどのように実装できるかを示すのは興味深いだろうと思いました。また、検出を回避するテキストを生成するための小さな言語モデルをトレーニングするための検証器としても使用します。これは、AI 検出器の限界を研究し、数学とコードで訓練された通常の推論モデルを超えた検証ベースの LLM アプリケーションを探索するための小規模な教育プロジェクトです。
図 1: サブスタックには AI 検出器が組み込まれています。
したがって、上で述べたように、このチュートリアルの目的は、(単純な) AI 検出器を構築することによって、AI 検出器がどのように機能するかを説明することです。
実際には、このような検出器はスパムコンテンツをフィルタリングするために使用できますが、AI が生成したテキストに変換することなく個人的な文章を改善することもできます。たとえば、長い記事を書いていて、スペルと文法を改善したい場合、文法チェッカーを使用して文章を磨き、読みやすさを向上させたくなるでしょう (そして実際に便利です)。 ChatGPT のような汎用 LLM など、そのためのさまざまなサービスがあります。ただし、これには、これらのツールによって、たとえそれが自分の文章であっても、洗練されすぎて AI のように聞こえるものになり、スパム コンテンツとしてフラグが立てられるリスクも伴います。
たとえば、AI チェッカーを使用すると、次のように言うことができます。

私のテキストが AI によって生成されたスコアが 0% であることを確認しながら、mmar を実行します。」
とにかく、ここでは完全に機能するチェッカーを構築していますが、目的は、1) AI チェッカーがどのように機能するか (できるか) を説明し、2) これを、LLM で使用できるスコアラーまたはベリファイアを構築する方法に関するより一般的なトピックのケーススタディとして使用することです。
免責事項: AI チェッカーは本質的にいたちごっこです。 AI チェッカーは、AI が生成したコンテンツを示す特定のパターンを検出することを学習する可能性があります。その後、次の LLM は偶然または意図的にそのパターンを示さず、検出を回避する可能性があります。その後、AI チェッカーを更新して、LLM などを検出する必要があります。さらに、誤検知 (AI 生成としてフラグが付けられた人間の書いたテキスト) が発生する可能性もありますが、これについては後ほど説明します。
このプロジェクトにはいくつかの目標があります。もちろん、包括的な目標は、AI 検出器がどのように機能するかを説明し、実際に使用するための評価、トレーニング、ローカル展開を含む、適用されたエンドツーエンドの LLM プロジェクトを示すことです。
その結果、人間とエージェントが使用できる AI 検出器 API と、ユーザーフレンドリーな UI が生まれました。
図 2: このプロジェクトの後半で開発されたローカル ブラウザ インターフェイスのプレビュー。全文 AI スコアを返し、個々のテキスト チャンクのスコアを強調表示することもできます。
ここでは、パングラム モデルと同様のメソッドを開発します。これは、私の知る限り、Substack AI 検出機能の背後にあります。
私は 2023 年に AI テキスト検出に関する短い記事を書きました: ChatGPT などの LLM によって生成されたコンテンツを検出するためのさまざまなアプローチとは何ですか?そして、それらはどのように機能し、どのように異なるのでしょうか?
基本的に、AI によって書かれたテキストを検出するには、教師あり分類器や摂動ベースの確率テストから、複雑さの測定や透かしの挿入まで、さまざまな方法があります。
このチュートリアルでは、

0 ～ 100 のスコアを返すモデル。これは本質的に、推定された確率スコアを持つ分類器です。確率スコアは、テキストが分類子に従って AI によって生成される可能性を示します。 (正確に言えば、スコアはトレーニング分布に基づいて AI が生成したクラスについて分類器が推定した確率です。ただし、テキストが AI によって書かれた一般的な確率として解釈すべきではありません。)
このために、DistilBERT 分類器を微調整します (私の初期の Substack 記事の 1 つである大規模言語モデルの微調整で説明したことと同様です) が、その詳細については、その段階に到達した後で説明します。
この投稿は有料購読者向けです

## Original Extract

An End-to-End Project With Dataset Construction, Model Training, Local Deployment, and RLVR

Building an AI Text Detector From Scratch
Subscribe Sign in Building an AI Text Detector From Scratch
An End-to-End Project With Dataset Construction, Model Training, Local Deployment, and RLVR
Sebastian Raschka, PhD Aug 15, 2026 ∙ Paid 26 3 Share Substack recently launched its AI detector feature in the UI, which is super interesting.
Separately, lots of people asked me about interesting local do-it-yourself LLM projects as demos to show what small language models (SLMs) are capable of.
Putting one and one together, I thought it would be interesting to show how an AI detector can be implemented. I will also use it as a verifier to train a small language model to produce text that avoids detection. This is a small educational project for studying the limitations of AI detectors and exploring a verifier-based LLM application beyond regular reasoning models trained on math and code.
Figure 1: Substack now features a built-in AI detector.
So, as mentioned above, the intended goal of this tutorial is to explain how AI detectors work by building (a simple) one.
In practice, such a detector can be used to filter out spammy content, but also to potentially improve your personal writing without turning it into AI-generated text. For example, if you wrote a lengthy article and want to improve spelling and grammar, it is tempting (and actually useful) to use a grammar checker to polish it and improve readability. There are different services for that, including general-purpose LLMs like ChatGPT. However, this also runs the risk that these tools turn your writing, even though it’s still your own writing, into something that is then overpolished and now sounds like AI and gets flagged as spammy content.
For example, with an AI checker, one could say, “Fix my grammar while ensuring that my text still scores 0% AI-generated.”
Anyway, while we are building a fully functional checker here, the goal is to explain 1) how AI checkers (can) work and 2) use this as a case study for a more general topic on how to build a scorer or verifier that can be used with LLMs.
Disclaimer: AI checkers are essentially a cat-and-mouse game. AI checkers may learn to detect a certain pattern that is indicative of AI-generated content. Then, the next LLM may incidentally or deliberately not exhibit that pattern and avoid detection. The AI checker then has to be updated to detect said LLM, and so forth. Plus, it’s also likely to encounter false positives (human written text flagged as AI-generated), but more on that later.
There are several goals of this project. The overarching goal is, of course, to illustrate how AI detectors work and show an applied end-to-end LLM project including evaluation, training, and local deployment for real-world use.
The outcome of this is an AI-detector API that can be used by humans and agents, and a user-friendly UI.
Figure 2: Preview of the local browser interface developed later in this project. It returns a whole-text AI score and can also highlight the scores for individual text chunks.
Here, we are going to develop a method similar to Pangram models, which, as far as I know, are behind Substack AI detection feature.
I wrote a short article about AI-text detection a while back in 2023: What Are the Different Approaches for Detecting Content Generated by LLMs Such As ChatGPT? And How Do They Work and Differ?
In essence, there are different ways to detect AI-written text, from supervised classifiers and perturbation-based probability tests to perplexity measures and watermarking.
In this tutorial, we will build a model that returns a 0-100 score. It’s essentially a classifier with an estimated probability score. The probability score will denote how likely a text is AI-generated according to the classifier. (Or, to be precise the score is the classifier’s estimated probability for the AI-generated class based on its training distribution. However, we shouldn’t interpreted it as a general probability that the text was written by AI.)
For this, we are going to fine-tune a DistilBERT classifier (similar to what I described in one of my early Substack articles, Finetuning Large Language Models ), but more details on that later when we get to that stage.
This post is for paid subscribers
