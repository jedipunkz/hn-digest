---
source: "https://www.teachmecoolstuff.com/viewarticle/fine-tuning-a-local-llm-to-categorize-questions"
hn_url: "https://news.ycombinator.com/item?id=48590771"
title: "Fine Tuning a Tiny Local LLM to Categorize Questions"
article_title: "Fine Tuning a Local LLM to Categorize Questions"
author: "dev-experiments"
captured_at: "2026-06-18T21:46:17Z"
capture_tool: "hn-digest"
hn_id: 48590771
score: 1
comments: 0
posted_at: "2026-06-18T20:04:32Z"
tags:
  - hacker-news
  - translated
---

# Fine Tuning a Tiny Local LLM to Categorize Questions

- HN: [48590771](https://news.ycombinator.com/item?id=48590771)
- Source: [www.teachmecoolstuff.com](https://www.teachmecoolstuff.com/viewarticle/fine-tuning-a-local-llm-to-categorize-questions)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T20:04:32Z

## Translation

タイトル: 質問を分類するための小さなローカル LLM の微調整
記事のタイトル: 質問を分類するためのローカル LLM の微調整
説明: <p>楽しい個人プロジェクトとして、メンテナンスの質問から医師の診察の予約まで、家族に関する一般的な質問に答えるチャットボットの開発に取り組んできました。 </p>
<p>
一般的なアイデアは、チャットボットが RAG を介してベクター データベースにクエリを実行することで日常の知識を取得するというものです。
[切り捨てられた]

記事本文:
質問を分類するためのローカル LLM の微調整
ホーム
AI
バゼル
WASM
クロージャコンパイラ
Kubernetes
JavaScript
反応する
ノーデジス
スレンダー
.ネット
テスト
クールなものを教えてください
質問を分類するためのローカル LLM の微調整
楽しい個人プロジェクトとして、メンテナンスの質問から医師の診察の予約まで、家族に関する一般的な質問に答えるチャットボットの開発に取り組んでいます。
一般的なアイデアは、チャットボットがベクトル データベースのクエリから RAG を介して日常の知識を取得するというものですが、より良い結果を得るために、ベクトル検索のメタデータを認識するようにしました。
基本的に、質問を既知のメタデータ カテゴリ (プール、車、空調設備、料理など) に分類する前処理ステップを通じて質問を実行しています。この主な目的は、ベクトル ランキングの検索スペースを、質問のカテゴリに一致するインデックス付きエントリのみに絞り込むことです。例として、「プールポンプはいつ交換しましたか?」という質問があります。インデックス データベースにクエリを実行する前に、「プール」というカテゴリにマップされます。
この実験でテストしたい仮説は、非常に小さなローカル LLM を微調整して、家庭関連の質問のデータセットでトレーニングしたときに信頼性の高い質問分類を実行できるかどうかです。
このプロジェクトでは、2 つの異なるローカル llms、Qwen 3:4B と Qwen 3:0.6B を使用しています。 4B パラメータ バージョンは一般的な質問の回答に使用され、超小型の 0.6B バージョンは質問の分類に使用されます。この実験の全体の前提は、わずか 6 億パラメータを持つ小さな llm を微調整して、家庭の質問の信頼できる分類器にできるかどうかを確認することです。
微調整
微調整には、Unsloth と呼ばれる人気のあるオープンソース フレームワークを使用しています。これは、Qwen や Llama などのローカル モデルの調整に適しているようです。
トレーニング目的のため、私の最初のデータセットは

約 850 個のデータ エントリで構成されており、70/15/15 パーセンテージベースでトレーニング データ、評価データ、テスト データにそれぞれ分割しています。トレーニング データと評価データはトレーニング中に使用されますが、テスト データセットは保留され、トレーニング後のテストの実行に使用されます。サンプルデータについては、以下のセクションを参照してください。
基本的な考え方は、信頼できる質問分類子になるように教えるために、十分な家庭用質問セットで LLM をトレーニングすることです。
微調整を行う前に、測定対象となるベースラインを確立することが重要です。この実験のベースラインは、プロンプトのみで元の Qwen 0.6B モデルを「そのまま」使用しようとすることです。ベースラインに使用されるサンプル プロンプトは以下にあります。
オフライン評価方法の 1 つとして、2 番目のデータセットのシナリオを使用してモデルをテストするために、約 130 の統合テストを作成しました。ベースライン モデルの結果は不十分です。 131 回のテストのうち、モデルは 13 個の質問のみを正しく分類しました (正解率は約 10%)。以下の概要を参照してください。
実際の障害を詳しく調べると、いくつかの一般的なパターンが浮かび上がってきます。
このモデルは主に、電気/電化製品などの広範なラベルを過剰に使用しており、他のカテゴリ (プール、料理、空調設備など) のほとんどが欠落しています。
このモデルは新しいカテゴリ (アパートなど) を発明し、提供された許可されたカテゴリのリストに固執しません。
以下にテストレポートからの抜粋を提供します。
ベースラインの結果から、Qwen 3 0.6B のような小型モデルはプロンプトだけでは信頼できるパフォーマンスを提供できないことが明らかになりました。
次の実験に関しては、前と同じプロンプトを使用していますが、より正確に分類する方法をモデルに教えるためにモデルの微調整を行っています。
興味がある場合に備えて、ここに微調整スクリプトを含めました。大まかに言うと、私は Unsloth を活用しています

微調整戦略としての 2 番目の QLora。注意: Unsloth が提供するデフォルトの微調整パラメータは、非常に優れた出発点となります。私の経験では、少なくとも最初は、Unsloth 値をあまり調整することを心配するよりも、適切なデータセットを用意することの方が重要です。
ただし、避けるべき一般的な落とし穴の 1 つは、トレーニング データでの過剰適合です。そのため、トレーニング データにないデータでモデルをテストすることが重要です。静的なトレーニング/テスト データに加えて、将来の再トレーニング中に 2 番目のチャネルとしてトレーニング データを修正するためにユーザー フィードバックを提供する方法も組み込みました。
一連の統合テストを実行した後、以下のレポートに見られるように、予測精度が明らかに向上していることがわかりました。
予測精度は 10% から 79% に向上しましたが、依然として不正確な結果の明らかなパターンがいくつか見られます。
モデルは正しい方向に向かう明確な兆候を示していますが、許可リストから正しいカテゴリの断片のみを出力するパターンが見られます。いくつかの例は、HVAC の代わりに AC/AIR です。
このモデルは、噴水、給湯器、プールからの水ベースの混乱など、意味的に重複するカテゴリによって混乱します。
最初の微調整実験を簡単に改善するには、後処理ステップを追加します。これにより、予測が意味的には正しいが、構文的に正しくない場合 (例: ac、air) の結果を正規化できます。もう 1 つの調整は、より多くの例を提供して、何をすべきか、何をすべきでないかをモデルに指示することで、プロンプト自体をさらに強化することです。どちらの考えも合理的だと思いますが、カテゴリが追加されると、メンテナンスの手間が増えます。
代わりに、モデルに次のことを教える方法にいくつかの変更を加えることで、微調整のアプローチを少し調整できるかどうかを確認したいと思いました。

マップのカテゴリー。
プロンプトに小さな変更を加えることで、最初の実験と比較して精度をさらに向上できることがわかりました。この微調整は、実際にはプロンプトを単純に変更するだけであり、以下のサンプルに見られるように、カテゴリをセマンティックの重複のない 2 文字の不透明な ID にマップします。
ここで、意味が重複する可能性のある可変カテゴリ文字列 (水ベースのカテゴリなど) の代わりに、固定形式のコードを出力するようにモデルに依頼します。
興味深いのは、以下の概要に見られるように、この単純な変更によりパフォーマンスが大幅に向上していることがわかります。
ご覧のとおり、予測精度は ~92% となり、かなり正確です。固定の重複しない出力を要求することは、応答を生成するときに小さな qwen モデルに役立つようです。
それでもまだいくつかミスがあります。以下に具体的な失敗を記載しました。
現時点では、予測は一般的に信頼でき、微調整された llm はチャットボットで使用可能な予測子として機能しますが、まだ取り組むべき問題が残っています。目立つ問題の 1 つは、給湯器 -> プールですが、これはおそらく、これら 2 つのカテゴリ間で「水っぽい」という意味が重複しているためと考えられます。これに対処するには、トレーニング データを再検討して、さらに微妙なデータにする必要があるでしょう。
チャット対話のサンプルを以下のスクリーンショットに示します。青い質問バブル内の小さなカテゴリ タグ (例: 「プール」) に特に注意してください。これは、小さな qwen 3:0.6B llm によって自動的に分類される部分だからです。
興味がある場合に備えて、ここに Github リポジトリを含めました。

## Original Extract

<p>As a fun personal project, I have been working on a chatbot for answering general questions about my household on anything from maintenance questions to doctor’s appointments. </p>
<p>
The general idea is that the chatbot will get its household knowledge through RAG from querying a vector databas
[truncated]

Fine Tuning a Local LLM to Categorize Questions
Home
AI
Bazel
WASM
Closure Compiler
Kubernetes
Javascript
React
Nodejs
Svelte
.Net
Testing
Teach Me Cool Stuff
Fine Tuning a Local LLM to Categorize Questions
As a fun personal project, I have been working on a chatbot for answering general questions about my household on anything from maintenance questions to doctor’s appointments.
The general idea is that the chatbot will get its household knowledge through RAG from querying a vector database, but for better results I have made the vector searches metadata aware.
Basically, I am running questions through a pre-processing step to categorize questions into known metadata categories (e.g. pool, car, hvac, cooking). The main goal of this is to narrow down the search space for vector ranking to only indexed entries that match the category of the question. As an example, the question “When did we replace our pool pump?” will be mapped to a category called “pool” before querying the Index database.
The hypothesis I want to test in this experiment is whether a very small local LLM can be fine-tuned to perform reliable question categorization when trained on a dataset of household-related questions
In this project I am using two different local llms – Qwen 3:4B and Qwen 3:0.6B. The 4B parameter version is used for general question answering, while the super tiny 0.6B version is used to categorize questions. The whole premise of this experiment is to see if a tiny llm with only 600M parameters can be finetuned into a reliable classifier of household questions.
Finetuning
For finetuning I am using a popular open-source framework called Unsloth, which seems well suited for tuning local models like Qwen and Llama.
For training purposes my initial dataset consists of about ~850 data entries where I do a 70/15/15 percentage-based split into training data, eval data and test data respectively. Training data and eval data are used during training, while the test dataset is withheld and used to run a test post training. See section below for sample data:
The basic idea is to train the llm on a sufficient set of household questions to teach it to become a reliable question classifier.
Before doing any finetuning, it’s important to establish a baseline to measure against. In this experiment the baseline is to try to use the original Qwen 0.6B model “as is” through prompting alone. A sample prompt used for the baseline can be found below:
As one of my offline eval methods I have created a battery of ~130 integration tests to test the model with scenarios from a second dataset. For the baseline model, the results are poor. Out of 131 tests the model only categorized 13 questions correctly (~10% correct responses). See summary below:
When digging into the actual failures a few common patterns emerge:
The model is mostly overusing broad labels like electric/appliances and missing most of the other categories (e.g pool, cooking, hvac).
The model invents new categories (e.g. apartments) and doesn’t stick to the provided list of allowed categories
I have provided an excerpt from the test report below:
The results from the baseline made it clear that a tiny model like Qwen 3 0.6B cannot provide reliable performance through just prompting alone.
As for the next experiment, I am using the same prompt as before, but I am doing model finetuning to teach the model how to categorize with greater accuracy.
I have included the finetuning script here in case you are interested in checking it out. At a high level I am leveraging Unsloth with QLora as the finetuning strategy. One note: The default fine tuning parameters provided by Unsloth provide a very good starting point. It’s been my experience that it’s more important to come up with a good dataset than worrying about tweaking the Unsloth values too much, at least to start.
One common pitfall to avoid though is overfitting on the training data, which is why it’s important to test the model on data not found in the training data. In addition to the static training/test data I have also incorporated a way to provide user feedback to amend the training data as a second channel during future retraining.
After running the battery of integration tests, I observed a clear improvement in prediction accuracy as seen in the report below:
The prediction accuracy is up from 10% to 79%, but I still see some clear patterns of incorrect results:
The model now shows clear signs of heading in the right direction, but I see a pattern of only emitting fragments of the correct categories from the allowed list. Some examples are ac/air instead of hvac
The model gets confused by semantically overlapping categories like water-based confusion from fountain, water heater and pool.
An easy improvement on the first fine tuning experiment would be to add a post processing step. This would allow me to normalize results where the prediction is semantically correct, but syntactically incorrect (e.g. ac, air). Another tweak would be to build more reinforcement into the prompt itself by providing more examples, telling the model what to do and not to do. I would say both ideas a reasonable, but it leads to more maintenance as more categories are added.
Instead, I wanted to see if I could tweak the finetuning approach slightly by making some changes to how I teach the model to map categories.
It turns out we can make a minor change to the prompt to improve accuracy even more compared to the 1st experiment. The tweak is actually just a simple change to the prompt where I map the categories to a two-character opaque IDs with no semantic overlap as seen in the sample below:
Now, I ask the model to output a fixed format code instead of a variable category string with potentially overlapping meaning (e.g. water-based categories).
The interesting part is that I see a very nice boost in performance from this simple change as seen in the summary below:
As you can see, prediction accuracy is now at ~92%, which is pretty accurate. It appears that asking for fixed, non-overlapping output helps the tiny qwen model when generating responses.
There are still a few misses though. I have included the specific failures below:
At this point the predictions are generally reliable, and the finetuned llm serves as a usable predictor in my chatbot, but there are still issues to work on. One issue that stands out is water heater -> pool, which is still likely due to the overlapping “watery” meaning between those two categories. To address this, I will likely have to revisit the training data and make it even more nuanced.
A sample chat interaction can be seen in the screenshot below. Pay special attention to the little category tag in the blue question bubbles (e.g. “pool”) since that is the part that is automatically classified by the tiny qwen 3:0.6B llm.
I have included the Github repo here in case you are interested in checking it out
