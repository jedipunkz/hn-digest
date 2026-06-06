---
source: "https://knl.co.in/blog/ml-foundations/"
hn_url: "https://news.ycombinator.com/item?id=48423485"
title: "ML Fundamentals"
article_title: "ML Foundations - Kunal Singh • kunalsin9h"
author: "kunalsin9h"
captured_at: "2026-06-06T12:34:34Z"
capture_tool: "hn-digest"
hn_id: 48423485
score: 1
comments: 0
posted_at: "2026-06-06T10:30:45Z"
tags:
  - hacker-news
  - translated
---

# ML Fundamentals

- HN: [48423485](https://news.ycombinator.com/item?id=48423485)
- Source: [knl.co.in](https://knl.co.in/blog/ml-foundations/)
- Score: 1
- Comments: 0
- Posted: 2026-06-06T10:30:45Z

## Translation

タイトル: ML の基礎
記事のタイトル: ML Foundations - Kunal Singh • kunalsin9h
説明: ML の基礎

記事本文:
ML Foundations - Kunal Singh • kunalsin9h
ブログ メニュートグル ブログ
知性には世界がどのように機能するかを理解する必要があります。私たちは予測を行うために世界をモデル化します。つまり、モデルとは予測を可能にするものです。
私たちの脳はモデルです、メンタルモデルを覚えていますか?
これらのモデルはどうやって作ることができるのでしょうか?つまり、何かを予測できる生物です。コンピューターはその良い基盤であることがわかりました。ソフトウェアモデルを作成することで、コンピューターを活用して予測することができます。
コンピュータが学習できる 3 つの異なる方法:
コンピュータ学習 : コンピュータは明示的な指示がなくても何かを行うことができます。
コンピューターがデータから直接タスクを学習できるようにします。
トレーニング入力データ
トレーニングの予想される出力データ
パラメーターを見つける必要があります。パラメーターをノブとして考えます (パーセプトロンを参照)。これは、モデルが期待値に最も近い値を予測するのに役立ちます。
損失 = 期待値 - 予測値
トレーニング段階の目標は、より低い損失を与えるノブを見つけることです。
予測値が期待値よりも大きい場合、損失値は負になる可能性があるため、それを二乗してみましょう。
損失 = (期待値 - 予測値) ^ 2
損失 = (Y - X0) ^ 2
より優れたモデルは、予測を現実によく適合させます。
したがって、機械学習はデータから学習することであり、データポイントは 特徴 とも呼ばれます。つまり、モデルが 4 つの入力を取る場合、モデルには 4 つの特徴があります。
たとえば、上記の線形回帰の単純な例では、モデルは次のようになりました。
複数の機能 (入力) がある場合、各機能に等しいワット数 (重み) を与えるとどうなるでしょうか。1 つの機能が他の機能よりも重要ですか?
最適な特徴を自ら学習するニューラルネットワーク。
特徴はデータ（入力情報）です
パラメータは学習された重みです。
ディープラーニングでは、ニューラルネットワークを使用して最適な特徴を学習します。
どの機能が他の機能よりも重要か。
シリーズ

(事実上) 任意の関数を近似できる演算。
Sum(wixi) + b は線形回帰でも見られます。ただし、活性化関数は非線形関数であり、ニューラル ネットワークに任意の関数を近似する機能を与えます。
Activation Functions はニューラル ネットワークの意思決定者です。これはニューロンの端に位置し、そのニューロンが「起動」する (情報を転送する) か、それとも休止状態を維持するかを決定します。
アクティベーション関数のこの非線形性により、モデルが非線形になるため、ニューラル ネットワークは複雑な決定境界を学習します。
人工ニューラル ネットワークは人間の脳からインスピレーションを受けています。あなたの脳では、生物学的ニューロンが隣接するニューロンから電気信号を受け取ります。あらゆる小さな火花をただ伝達するだけではありません。信号が閾値を超えるのに十分な強度になるまで待機し、その後活動電位を「発火」させます。
これらのタイプのアクティベーション関数は、さまざまなタイプのニューラル ネットワークで使用されます。
バニラ : 加重合計 + アクティベーション。 FFNN と CNN。
LSTM : Long Short Term Memory、シーケンス内の長期依存関係のためのメモリとゲートを備えた高度なニューロン。
RNN、CNN は、異なる種類の層構造を持っているため、ニューラル ネットワークの一種です。
注意もニューラル ネットワーク層構造の一種です。
トレーニングプロセスの指針となる価値観
コンピューターはデータからではなく、試行錯誤から学習します。
AI エンジニアのための ML の基礎 (34 分)
© 2026 クナル・シン。無断転載を禁じます。

## Original Extract

ML Foundations

ML Foundations - Kunal Singh • kunalsin9h
Blog Toggle Menu Blog
Intelligence require understanding how the world works. We Model the world to make prediction, i.e a Model is something that lets us make prediction.
Our brain is a model, remember Mental Model?
How can we make these models? i.e organisms that can predict something. Turns out Computers are good base for it. We can leverage Computer to predict via making Software Models .
3 Difference Ways Computers can Learn :
Computer Learning : Computer can do things without explicit instructions.
Allow computers to learn tasks directly from data.
Training Input Data
Training Expected Ouput Data
We need to find parameters, I something think about parameters as knobs (see Perceptron ), that will help the model predict the closest values to the expected values
Loss = Expected Value - Predictoin Value
The Goal of training phase is to find such knobs that will give lower loss.
Loss value can be negative is prediction value is bigger then expected value, so lets square it.
Loss = (Expected Value - Prediction Value) ^ 2
Loss = (Y - X0) ^ 2
The better models fits the Prediction to Reality very well.
So Machine learning is learning from Data, data points are also called features , i.e if a model takes 4 inputs, then it has 4 features.
For example, in our linear regression simple example given above, our model was:
What if we have multiple features (inputs) then are we giving equal wattage (weights) to each features, is one feature more important then other?
Neural Networks that learn optimal features on their own.
Feature is the data (input information)
Parameters are weights that are learned.
In Deep Learning we use Neural Networks to learn optimal features,
which features are more important then others.
A series of operations that can approximate (practically) any function.
The Sum(wixi) + b is also seen in Linear regression. but the g the Activation function is a non-linear function which gives the neural network the ability to approximate any function.
Activation Functions is a Decision Maker of a neural network. It sits at the end of a neuron, and decides whether that neural should “ fire ” (pass information forward) or stay dormant .
This non-linearity of Activation function make the neural network learn complex decision boundary, since it forces the model to become non-linear.
Artificial neural networks are inspired by the human brain. In your brain, a biological neuron receives electrical signals from neighbors. It doesn't just pass every tiny spark along; it waits until the signal is strong enough to cross a threshold, and then it "fires" an action potential.
These types of Activation functions are used in different types of Neural networks:
Vanilla : Weighted Sum + Activation. FFNNs and CNNs.
LSTM : Long Short Term Memory, advance neuron with memory and gates for long-term dependencies in sequences.
RNNs, CNNs are type of neural networks because they have different type of Layer structure.
Attention is also a type of neural network layer structure.
Values that guide the training process
Computer learn from trail and error, not from data.
ML Foundations for AI Engineers (in 34 Minutes)
© 2026 Kunal Singh. All rights reserved.
