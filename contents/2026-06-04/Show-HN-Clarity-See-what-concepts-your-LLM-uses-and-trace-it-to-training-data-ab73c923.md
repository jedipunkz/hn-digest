---
source: "https://www.guidelabs.ai/post/meet-clarity/"
hn_url: "https://news.ycombinator.com/item?id=48401606"
title: "Show HN: Clarity, See what concepts your LLM uses and trace it to training data"
article_title: "Introducing Clarity"
author: "adebayoj"
captured_at: "2026-06-04T18:53:51Z"
capture_tool: "hn-digest"
hn_id: 48401606
score: 3
comments: 1
posted_at: "2026-06-04T17:14:16Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Clarity, See what concepts your LLM uses and trace it to training data

- HN: [48401606](https://news.ycombinator.com/item?id=48401606)
- Source: [www.guidelabs.ai](https://www.guidelabs.ai/post/meet-clarity/)
- Score: 3
- Comments: 1
- Posted: 2026-06-04T17:14:16Z

## Translation

タイトル: Show HN: Clarity、LLM が使用する概念を確認し、それをトレーニング データに追跡します
記事のタイトル: Clarity の紹介
説明: Clarity は、初の本質的に解釈可能な AI プラットフォームであり、現在研究プレビューとして招待制で利用可能です。

記事本文:
明瞭性の導入
ガイドラボ
私たちについて ブログ お問い合わせ 採用情報 -->
Clarity にアクセス
著者: Isaac Plant 、運営および事業開発責任者 Zhichen Guo 、プロダクト エンジニア 発行日: 2026 年 6 月 3 日
最初の本質的に解釈可能な AI プラットフォームである Clarity を紹介します。現在、研究プレビューとして招待制で利用可能です。現在の AI システムは、内部推論が不透明なブラック ボックスであり、出力を入力データまたはトレーニング データまで追跡する機能がありません。 Steerling 8B を搭載した Clarity は、これらの問題を解決します。これを使用すると、次のことが可能になります。
モデルがどのように推論するかを調べてください。モデルの出力を促進する人間が理解できる概念をご覧ください。
トレーニング データへのトレース出力。出力がモデルのトレーニングに使用されたデータにどのように関連しているかを理解します。
概念を使用してモデルの動作を制御します。概念を増幅および抑制して、プロンプトを使用せずにモデルの出力を制御します。
本日、当社は初の本質的に解釈可能な AI プラットフォームである Clarity を発表します。 Clarity は、当社の指示により微調整された Steerling-8B モデルによって強化されています。他のモデルはブラックボックスであるか、解釈可能性が事後的に組み込まれています。これらの方法では、追跡できないエラーや診断できない誤った推論を含む出力が生成されます。 Steerling は、トレーニング中に解釈可能性が組み込まれた最初のモデルであり、Clarity プラットフォームを使用すると、これらの新しい機能を直接操作できます。この記事の残りの部分では、Clarity の 3 つの主要な機能について説明します。
コンセプトの説明: Steerling が出力を生成するために使用する人間が理解できるコンセプト
トレーニングデータの帰属: 出力に帰属するトレーニングデータ
コンセプト ステアリング : プロンプトを変更するのではなく、コンセプトを増幅または抑制することで Steerling の出力を制御します。
Clarity は他のチャットと同じように見えます

ボットには、ステアリング ボタンという大きな違いがあります。このボタンを使用すると、AI の応答内の概念を増幅または抑制できます。
しかし今は、アフリカの動物相について調べて聞いてみましょう。
応答を見ると、Clarity の特徴である説明パネルがすぐにわかります。
コンセプトとトレーニング データ ソースへのトレース出力
Clarity は、AI がその出力を生成する方法に関する 2 つの洞察、コンセプトとトレーニング データ アトリビューションを提供します。まず、コンセプトを見てみましょう。これらは、モデルが推論するために使用する人間が理解できる特徴です。
何も選択しないと、説明パネルにはチャットで最も一般的な概念が表示されます。この出力は意味があるようです。アフリカの生き物についての質問に答えるとき、モデルは野生動物について考えていると予想されます。
モデルはテキストをチャンクに分けて生成します。チャンクをクリックすると、モデルがそのチャンクを生成するために使用した概念を確認できます。
次に、プラットフォームの別の機能であるトレーニング データの説明を見てみましょう。この機能を使用すると、トレーニング セット内のどのチャンクが生成されたチャンクに最も類似しているかを確認できます。
プロンプトを変更せずに出力内のコンセプトを方向付ける
Clarity がモデルの内部動作を公開する方法を確認したので、これらを使用して、プロンプトに依存せずにモデルの出力を操作してみましょう。現在のプロンプトでは、アフリカに生息する信じられないほどの動物についての回答が得られました。しかし、魚もフアナであるため、軽視されてきました。それを修正できるかどうか見てみましょう。
これを行うには、プロンプトを編集し、ステアリング ボタンをクリックします。
これにより、検索バーが表示されるので、「marine」と入力します。
いくつか選択肢はありますが、「Marine Sea Life」がぴったりだと思います。 「追加」をクリックしましょう。 Amplify はデフォルトで選択されており、これが私たちが望んでいることです。

それで準備は完了です。
[送信] をクリックしてチャット ウィンドウで続行することもできますが、比較パネルに移動しましょう。これにより、最初のプロンプトとの違いがわかります。
そして出来上がり！これで、魚に関するあらゆる情報が得られました。この出力を選択してメイン画面に戻ると、これがチャットの説明に反映されていることがわかります。水生生物関連のものがたくさんあります。
増幅は、概念がどのように機能するかを示す優れたデモンストレーションですが、多くの場合、これはプロンプトを変更することで実現できます。一方、抑制の信頼性は低くなります。
概念の抑制により、プロンプトが特定の出力を生成しようとしている場合でも、特定の出力を防ぐことができます。そのため、概念を抑制すると、トレーニングに頼ることなく LLM 製品を調整できるようになります。
これがどのように機能するかを確認するために、モデルにコンピューター科学者を記述するよう依頼してみましょう。
まあ、それは残念ですね。かなり男性中心ですね！モデルがコンピューター科学者は男性であると考えている場合、女性に関して不適切な採用決定を下す可能性があります。
「人称役割名詞」の概念を抑制することでこれを修正できるかどうかを見てみましょう。
素晴らしいですね。出力は性別に依存しません。私たちは、このチャットボットの採用プロセスをサポートする能力にさらに自信を持っています。
提携と今後の機能
Clarity は、本質的に解釈可能な初の AI プラットフォームであるため、上で共有した例以外にも探索すべきことがたくさんあります。プラットフォーム自体で追加の例を確認できます。今後数週間にわたって、Clarity のデモンストレーションをソーシャル メディア チャネルで共有する予定です。
当社は、特定の分野向けに最先端の解釈可能な AI ソリューションの開発に関心を持つエッジ企業と提携しています。ご興味がございましたら、こちらからお問い合わせください。
今後数か月以内の新機能に注目してください。

入力の属性を指定します。これにより、出力が入力の最も関連性の高い部分にリンクされます。今回の発売はClarityにとっての第一歩にすぎません。

## Original Extract

Clarity is the first inherently interpretable AI platform, now available by invitation as a research preview.

Introducing Clarity
Guide Labs
About us Blog Contact us Careers -->
Access Clarity About us
Authors: Isaac Plant , Head of Operations and Business Development Zhichen Guo , Product Engineer Published: June 03, 2026
We introduce Clarity , the first inherently interpretable AI platform, now available by invitation as a research preview. Current AI systems are black boxes with opaque internal reasoning and no ability to trace output back to input or training data. Powered by Steerling 8B, Clarity fixes these problems. With it, you can:
Explore how the model reasons . See the human-understandable concepts that drive model output.
Trace output to training data . Understand how the outputs relate to the data the model was trained on.
Steer model behavior using concepts . Amplify and suppress concepts to control the model’s output without using prompts.
Today we are launching Clarity, the first inherently interpretable AI platform. Clarity is powered by our instruction finetuned Steerling-8B model . Other models are either black boxes or have interpretability bolted on post-hoc. These methods result in outputs that have untraceable errors and faulty reasoning that can’t be diagnosed. Steerling is the first model that has interpretability built in during training, and the Clarity platform allows you to directly interact with these new capabilities. In the remainder of the post, we will walk you through three key capabilities of Clarity:
Concept explanations : the human-understandable concepts that Steerling uses to produce its output
Training data attribution : the training data attributed to the output
Concept steering : controlling the output of Steerling by amplifying or suppressing concepts, as opposed to changing the prompts
Clarity looks like other chat bots besides one big difference: the steering button. This button allows you to amplify or suppress concepts in the AI’s response.
But for now, let’s explore and ask about the fauna in Africa.
Looking at the response, we immediately see what sets Clarity apart: the Explanations panel.
Trace output to concepts and training data source
Clarity provides two insights into how the AI is generating its output, Concepts and Training Data Attribution. First, let’s look at Concepts. These are the human understandable features the model uses to reason.
With nothing selected, the Explanations panel shows the most common concepts in the chat. This output seems to make sense. We would expect the model to be thinking about Wildlife when responding to a question about living things in Africa!
The model generates text in chunks. You can click a chunk and see what concepts the model used to generate it.
Now let’s take a look at a different feature of the platform: Training Data explanations. With this feature, you can see which chunks in the training set are most similar to the generated one.
Steer any concept in the output without changing prompts
Now that we have seen how Clarity exposes the internal workings of the model, let’s use these to steer the models output without relying on prompts. The current prompt got us a response about the incredible animals living in Africa. Fish are fuana, too, though, and they have been given short shrift. Let’s see if we can remedy that.
To do this, we are going to edit the prompt and click on that steering button.
This brings us to a search bar, where we will enter “marine”.
There are a few different options, but “Marine Sea Life” seems to be a good fit. Let’s click add. Amplify is selected by default, which is what we want, so we are all set.
We could click Send and continue in the chat window, but let’s go to the Compare Panel. This will let us see the differences with the initial prompt.
And voila ! We now have all the information about fish we could hope for. If we select this output and return to the main screen, we can see this reflected in the Chat Explanations: Lots of aquatic-related things!
Amplification is a nice demonstration of how concepts work, but often this can be accomplished with modified prompts. Suppression, on the other hand, is less reliable.
Suppression of concepts allows you to prevent certain outputs even when the prompts may be trying to produce those outputs. As such, suppressing concepts allows you to align your LLM product without resorting to training.
To see how this works, let’s ask the model to describe a computer scientist.
Well, that is unfortunate. It is very male centric! If the model thinks computer scientists are men, it might make poor hiring decisions about women.
Let’s see if we can fix this by suppressing the concept of “Person-Role Nouns”.
Excellent, the output is now gender neutral. We can be more confident in this chatbot’s ability to support the hiring process.
Partnering and upcoming features
Clarity is the first inherently interpretable AI platform and, as such, there is a lot more to explore than the examples we have shared above. You can see additional examples in the platform itself and we’ll be sharing demonstrations of Clarity on our social media channels over the coming weeks.
We partner with edge companies that are interested in developing cutting-edge interpretable AI solutions for their particular domains. If you are interested, you can reach out to us here .
Keep an eye out for new features in the coming months, including input attribution, which will link the output to the most relevant parts of the input. This launch is just the first step for Clarity.
