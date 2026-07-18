---
source: "https://magazine.sebastianraschka.com/p/controlling-reasoning-effort-in-llms"
hn_url: "https://news.ycombinator.com/item?id=48960347"
title: "Controlling Reasoning Effort in LLMs"
article_title: "Controlling Reasoning Effort in LLMs"
author: "matt_d"
captured_at: "2026-07-18T17:46:53Z"
capture_tool: "hn-digest"
hn_id: 48960347
score: 1
comments: 0
posted_at: "2026-07-18T17:46:04Z"
tags:
  - hacker-news
  - translated
---

# Controlling Reasoning Effort in LLMs

- HN: [48960347](https://news.ycombinator.com/item?id=48960347)
- Source: [magazine.sebastianraschka.com](https://magazine.sebastianraschka.com/p/controlling-reasoning-effort-in-llms)
- Score: 1
- Comments: 0
- Posted: 2026-07-18T17:46:04Z

## Translation

タイトル: LLM における推論労力の制御
説明: LLM が低、中、高エフォートの推論モードを学習する方法

記事本文:
LLM での推論労力の制御
LLM での推論作業の制御
LLM が低、中、高エフォートの推論モードを学習する方法
Sebastian Raschka 博士 2026 年 7 月 18 日 125 4 12 シェア OpenAI が、LLM ベースの推論モデルのアイデアを普及させたモデルである o1 をリリースしてからほぼ 2 年が経ちました。 DeepSeek-R1 は、そのような推論モデルをトレーニングするための検証可能な報酬を伴う強化学習 (RLVR) レシピの詳細とともに、約 4 か月後に続きました。
先週、OpenAI は GPT-5.6 モデル ファミリをリリースしました。 3 つのサイズがあり、それぞれにおよそ 5 ～ 6 つの推論努力設定があります。
図 1: さまざまな推論努力設定を使用した GPT 5.6 Sol モデル。 (Ultra のベンチマーク数値は現在利用できませんが、Max と比較的似た値になるはずです。これは、同様の作業レベルを使用しますが、4 つのサブエージェントで作業を高速化するためです。)
はい、推論モデルは今後も存続します。これらは最新のモデルリリースの標準部分となっています。
過去に、私は推論モデルの方法論 (「推論 LLM について」) と関連する研究論文 (「LLM 推論のための強化学習の状態」および「LLM 推論モデル推論の状態」) を取り上げました。さらに、推論モデルの開発方法に関するまったく新しい 440 ページの本『Build A Reasoning Model (From Scratch)』を執筆しました。
図 2: 私の新しい Build A Reasoning Model (From Scratch) の本。カラーで！
これらのリソースは、従来の LLM を推論モデルに変換することに重点を置いています。さて、この記事では、この記事の冒頭の図に示されているような、複数の努力モードを持つ推論モデルを開発する方法に焦点を当てて説明したいと思います。
心配はいりません。この記事は単独の記事としても読むことができます。ただし、前述のリソース m

面白くて役に立つものばかりです。
1. 推論モデルの簡単な定義
機械学習や AI の技術やサブ分野について話すときに得られる教訓は、通常、専門用語を「文字通り」受け取ってはいけないということです。たとえば、機械学習や AI における (人工) ニューラル ネットワークは、人間の脳のような生物学的なニューラル ネットワークのように文字通り機能するわけではありません。
同様に、「推論モデル」について話すとき、これらのモデルが文字通り私たち人間と同じように推論することを期待すべきではありません。 AI および LLM 研究の文脈では、「推論モデル」とは、質問やタスクを段階的に処理する中間応答のような、中間推論トレースを出力するモデルを意味します。
これは、例を示して説明するのが最も簡単でしょう。
図 3: 従来の LLM 解答 (左) と推論モデルによる解答 (右) の図解。
2. トレーニングと推論のスケーリング推論モデルの概要
(推論) タスクのパフォーマンスを向上させるには、本質的に 2 つの方法があります。それは、トレーニング スケーリングと推論スケーリングです。
図 4: トレーニングと推論スケーリングは、LLM と推論モデルの問題解決能力を向上させる 2 つの方法です。 LLM による推論の学習に基づいたプロット
まずはトレーニングについて簡単に説明しましょう。
簡単に言うと、DeepSeek-R1 は、検証可能な報酬を伴う強化学習 (RLVR) を使用して LLM をトレーニングし、推論モデルに変えることを提案しました。 RLVR は、検証可能なデータ ドメインに報酬シグナル (0=不正解、1=正解) を提供する手法です。ここでの検証可能なデータ ドメインは、数学 (SymPy や WolframAlpha などの記号数学チェッカーを使用して結果を確認できます) とコード (コンパイラや単体テスト、または LeetCode などの統合プラットフォームを使用して正確性を確認できます) です。
図 5: acc の図

RLVR トレーニング中の uracy および format 報酬。
特に、推論トレース自体はモデルのトレーニングや更新には使用されませんでした。この中間応答情報をトレーニングに使用しようとしましたが、DeepSeek-R1 論文では、モデルのトレーニングには役に立たないと報告されており、最終的には使用されませんでした。 (プロセス報酬モデルを介してトレーニング信号に中間推論トレースを組み込むかどうか、またどのように組み込むかは、活発な研究分野です。)
図 6: 中間推論トレースは RLVR 中に無視されます。最終的な回答と回答形式のみが報酬を決定します。
いずれにせよ、図 7 に示すように、出力報酬だけをトレーニングするだけで、モデルが問題を推論する方法を学習するには十分であることがわかりました。つまり、モデルは中間の説明を書いたり、後戻りしたり、自己修正したりすることを学習します。モデルが間違いを犯したことに気づき、自己修正するこのような瞬間は、「Aha」モーメントと呼ばれます。
図 7: 「なるほど」の瞬間の例。推論モデルが中間推論の誤りに気づき、最終的な答えを導き出す前にそれを修正します。
ちなみに、DeepSeek-R1 が最も人気のある論文であり、検証可能な報酬を伴う強化学習と推論モデルの開発を中心に興奮を引き起こした論文であることは間違いありませんが、まったく同じ日に arXiv (2025 年 1 月 22 日) に公開された別の論文 Kim K1.5 があります。また、RLVR という用語は、その 2 か月前の「 Tülu 3: Pushing Frontiers in Open Language Model Post-Training 」ですでに作られていました。
DeepSeek R1 が最終的により人気のある論文である理由の 1 つは、推論動作が純粋な強化学習 (RL) で実現できることを実証したためです。
図 8: DeepSeek-R1-Zero は、事前トレーニングされたベース モデルに RLVR を直接適用します。

微調整の監督は行っていません。
たとえば、Tülu 3 と Kim K1.5 は、教師あり微調整 (SFT) モデルに強化学習を適用しました。 DeepSeek-R1 モデルも、DeepSeek-V3 基本モデルの SFT チェックポイントからトレーニングされ、純粋な RLVR でトレーニングされた DeepSeek-R1-Zero バリアントが含まれていました。 R1 Zero は R1 よりも弱いモデルですが、推論トレースの生成と使用をモデルに教えるには RLVR で十分であることが示されました。
R1-Zero はどちらかというと概念実証モデルでしたが、前述のように、完全な DeepSeek-R1 推論モデルのトレーニング パイプラインは通常多段階であり、少し複雑であることに注意してください。
図 9: より詳細な推論モデルのトレーニング パイプライン。これは、さまざまな DeepSeek-R1 モデルを示しています。詳細については、私の他の記事「推論 LLM について」を参照してください。
ちなみに、今日の LLM のほとんどは事実上推論モデルであり、RLVR 形式を使用して DeepSeek-R1 と同様の方法でトレーニングされています。
2.3 推論スケーリングの概要
トレーニングによる推論動作の改善の次に、モデルのパフォーマンスを向上させるためのもう 1 つの手段は、推論コンピューティングのスケーリングです。つまり、これは、モデルのトレーニング後、使用中に、より良い答えを得るために、より多くのコンピューティングを費やしていることを意味します。
これはそれ自体で完全なトピックであり、より詳細な概要については、私の「LLM 推論モデル推論の状態」を読んでください。
LLM 推論のための強化学習の現状 Sebastian Raschka 博士 · 2025 年 4 月 19 日 全文を読む
以下に背景情報として言及すべき最も重要なことを要約してみます。
まず、推論モデルは通常、従来の LLM と比較して推論中により多くのトークンを出力するため、RLVR を使用したモデルのトレーニングはすでに暗黙的に推論スケーリングの形式につながっています。

推論中により多くのコンピューティングを消費します。
2 番目に、推論の労力レベルによってこの出力の長さをさらに調整できますが、これについては後で詳しく説明します。
第三に、追加の推論スケーリング手法が多数あります。よく使われるのは自己整合性です。これは、モデルが複数回クエリされ、最終的な答えが多数決によって選択される多数決の形式として実装されることがよくあります。
図 10: 一般的な推論スケーリング手法である自己一貫性の例。
これは、従来の LLM だけでなく推論モデルにも適用できます。また、この方法はオンデマンドで、推論トレーニングに加えて使用することもできます。その好例が DeepSeekMath-V2 です。研究者たちは、(数学に特化した) 推論モデルに極端な推論スケーリングを適用し、数学オリンピックのような難解な問題で最先端のパフォーマンスを実現しました。
図 11: 数学のパフォーマンスを向上させるために併用される 2 種類の推論スケーリング (自己無矛盾性と自己洗練)。 DeepSeekMath-V2: 自己検証可能な数学的推論に向けての図
ただし、他の手法の概要については、私の別の記事「LLM 推論モデル推論の状態」を再度参照します。
LLM 推論モデル推論の状態 Sebastian Raschka 博士 · 2025 年 3 月 8 日 全文を読む 3. トークンを考える
先ほどの「なるほどと思う瞬間」の図で <think></think> トークンを見たことがあるかもしれません。上までスクロールする必要がないように、対応する図も下に示しました。
図 12: 推論モデルにおける一般的な書式設定トークン。
これらの <think> タグと </think> タグは、推論能力に関して表面的なものです。これらはモデルを推論するものではなく、優れた推論パフォーマンスを達成する必要もありません。これらの区切り文字を使用せずに同じモデルをトレーニングすると、SIM に到達する可能性があります。

ilar ベンチマーク パフォーマンス。
これらの <think> タグまたはトークンの目的は主に、推論トレースの開始位置と終了位置をマークして、トレーニング パイプラインまたはユーザー インターフェイスが推論トレースを最終的な答えから分離し、必要に応じてユーザーから非表示にできるようにすることです。 (通常、ChatGPT や Codex などの UI がこれを行います。)
ここで重要なのは、<think> トークンはモデルに「考える」能力や、推論したり、より適切に推論したりする能力を与えていないということです。このような <think> トークンなしで同じモデルをトレーニングし、同様のベンチマーク パフォーマンスに達することができます。
リテラル文字列 <think> と </think> についても特別なことは何もありません。別の区切り文字のペアも同じ目的に使用できます。
ちなみに、これを実装する方法は通常、RLVR 段階でフォーマット報酬を追加することです。したがって、単に回答の正しさに基づいてモデルに報酬を与えるのではなく、<think> トークンの使用に対して追加の報酬を提供し、それによってモデルがそれらのトークンを使用することを奨励します。
たとえば、DeepSeek-R1 では、全体の報酬は次のように計算されます。
R_total = R_accuracy + R_format
ここで、フォーマットの報酬は、モデルが内部に推論を配置することを促す単純なルールベースのチェックでした。
4. 推理モードのオン/オフスイッチ
第一世代の推論モデルは専用の推論モデルでした。つまり、DeepSeek-V3 ベース モデルと、別個の DeepSeek-R1 推論モデルがあったということです。
プロンプトが何であっても、R1 は通常、単純なプロンプトであっても、大量のトークンを使用して非常に詳細な応答を出力します。また、推論モードをオフにする組み込みオプションもありません。
図 13: 推論モデルは、最も単純なプロンプトであっても非常に冗長です。
Qwen3 などの後のモデルでは、同じモデルが通常の命令で微調整されたモデルや推論モデルのように動作できるハイブリッド アプローチを実験しました。

オンデマンドで注文できます。
注: モデル開発者の中にはこれを「思考モード」と呼ぶ人もいれば、「推論モード」と呼ぶ人もいます。どちらの用語も同じ動作を指します。
Qwen3 では、これは、enable_ Thinking=True または Enable_ Thinking=False を使用するトークナイザーを介して処理されます。内部的には、enable_ Thinking=False を設定すると、基本的に空の <think></think> セクションがアシスタントの応答の先頭に追加され、Qwen3 の推論 (「思考」) モードがオフになります。
図 14: Thinking=False および Thinking=True の場合の Qwen3 0.6B 推論モデルの応答。 (空の <think></think> タグは、生成された回答ではなく、変更された入力プロンプトの一部であるため、左側のインターフェイスでは非表示になっています。)
上の図に示すように、モデルが推論時にこの切り替えをサポートするように、これはトレーニング中にどのように実装されるのでしょうか?
つまり、Qwen3 テクニカル レポートで説明されているように、このオン/オフ動作は主に教師あり微調整 (SFT) によって導入され、最大のフラッグシップ モデルの一般的な RL 中に強化されます。
たとえば、最初の推論モデルが長期思考連鎖 SFT と推論 RL によってトレーニングされた後、「思考モード融合」ステージが追加されます。この追加の SFT ステージでは、モデルは思考例と非思考例の両方を確認します。
/考える: <考える>

[切り捨てられた]

## Original Extract

How LLMs Learn Low-, Medium-, and High-Effort Reasoning Modes

Controlling Reasoning Effort in LLMs
Subscribe Sign in Controlling Reasoning Effort in LLMs
How LLMs Learn Low-, Medium-, and High-Effort Reasoning Modes
Sebastian Raschka, PhD Jul 18, 2026 125 4 12 Share It has been almost two years since OpenAI released o1, a model that popularized the idea of LLM-based reasoning models. DeepSeek-R1 followed about four months later, together with details of a reinforcement learning with verifiable rewards (RLVR) recipe to train such reasoning models.
Last week, OpenAI released the GPT-5.6 model family. It comes in three sizes, each with roughly five or six reasoning-effort settings.
Figure 1: The GPT 5.6 Sol model with different reasoning effort settings. (Benchmark numbers for Ultra are currently not available but should be relatively similar to Max, since it uses a similar effort level but accelerates the work with four subagents.)
So yes, reasoning models are here to stay. They have become a standard part of modern model releases.
In the past, I covered the methodology of reasoning models ( Understanding Reasoning LLMs ) as well as relevant research papers ( The State of Reinforcement Learning for LLM Reasoning and The State of LLM Reasoning Model Inference ). And I even wrote a whole new 440-page book on how to develop reasoning models, Build A Reasoning Model (From Scratch) .
Figure 2: My new Build A Reasoning Model (From Scratch) book. In color!
These resources have focused on turning a conventional LLM into a reasoning model. Now, in this article, I want to focus on and explain how to develop a reasoning model that has multiple effort modes, similar to what’s shown in the figure at the beginning of this article.
No worries, this article can be read as a standalone article. However, the aforementioned resources may be interesting and useful.
1. A brief definition of reasoning models
When talking about pretty much any machine learning or AI technique or subfield, the one lesson is that we usually shouldn’t take technical terms “literally”. For example, an (artificial) neural network in machine learning and AI doesn’t literally work like a biological neural network like the human brain.
Similarly, when talking about “reasoning models”, we shouldn’t expect that these models literally reason like us humans. In the context of AI and LLM research, “reasoning model” means a model that outputs an intermediate reasoning trace, which is like an intermediate response that works through a question or task step by step.
It’s probably easiest to explain this by showing an example.
Figure 3: Illustration of a conventional LLM answer (left) and an answer by a reasoning model (right).
2. A brief overview of training and inference scaling reasoning models
There are essentially two ways to improve (reasoning) task performance: training scaling and inference scaling.
Figure 4: Training and inference-scaling are two ways to improve LLM and reasoning model problem-solving capabilities. Plot based on Learning to reason with LLMs
Let’s briefly talk about training first.
In a nutshell, DeepSeek-R1 proposed training an LLM using reinforcement learning with verifiable rewards (RLVR) to turn it into a reasoning model. RLVR is a technique to provide a reward signal ( 0=incorrect and 1=correct ) for verifiable data domains. These verifiable data domains here are math (we can use a symbolic math checker like SymPy or WolframAlpha to check results) and code (we can use a compiler or unit tests, or integrated platforms like LeetCode) to check for correctness.
Figure 5: Illustration of accuracy and format rewards during RLVR training.
Notably, the reasoning trace itself was not used for training or updating the model. Although they tried to use this intermediate response information for training, the DeepSeek-R1 paper reported that it wasn’t helpful for the model training, so it was ultimately not used. (Whether and how to incorporate intermediate reasoning traces in the training signal via process reward models is an active area of research.)
Figure 6: The intermediate reasoning trace is ignored during RLVR; only the final answer and response format determine the reward.
Anyway, just training on the output rewards alone, as Figure 7 shows, turned out to be sufficient for the model to learn how to reason through a problem, meaning that it would learn to write intermediate explanations, backtrack, and self-correct itself. These moments when the model realizes that it made a mistake and self-corrects itself are called “Aha” moments.
Figure 7: An example of an aha moment, where a reasoning model notices an error in its intermediate reasoning and corrects it before producing the final answer.
By the way, while DeepSeek-R1 is inarguably the more popular paper, and the paper that created excitement around reinforcement learning with verifiable rewards and the development of reasoning models, there is another paper, Kimi K1.5 , published on exactly the same day on arXiv (22 Jan 2025). Also, the term RLVR was already coined two months earlier in Tülu 3: Pushing Frontiers in Open Language Model Post-Training .
One reason why the DeepSeek R1 is ultimately the more popular paper is that it demonstrated that reasoning behavior can be achieved with pure reinforcement learning (RL).
Figure 8: DeepSeek-R1-Zero applies RLVR directly to the pretrained base model without supervised fine-tuning.
For instance, Tülu 3 and Kimi K1.5 applied reinforcement learning on top of a supervised fine-tuned (SFT) model. The DeepSeek-R1 model was also trained from an SFT checkpoint of the DeepSeek-V3 base model, and it included a DeepSeek-R1-Zero variant trained with pure RLVR. R1 Zero is a weaker model than R1, but it showed that RLVR is sufficient for teaching the model to generate and use reasoning traces.
​While R1-Zero was more of a proof-of-concept model, note that the full DeepSeek-R1 reasoning model training pipeline is usually multi-stage and a bit more complicated, as mentioned above.
Figure 9: More detailed reasoning model training pipeline. This one depicts the various DeepSeek-R1 models. For more details, see my other article: Understanding Reasoning LLMs ​
By the way, most of today’s LLMs are effectively reasoning models, meaning they have been trained in a similar fashion to DeepSeek-R1 using a form of RLVR.
2.3 Inference scaling in a nutshell
Next to improving reasoning behavior through training, another lever for improving model performance is inference compute scaling. In short, this means that we are spending more compute after training the model, during usage, to get better answers.
This is a whole topic by itself, and you could read through my The State of LLM Reasoning Model Inference for a more detailed rundown:
The State of Reinforcement Learning for LLM Reasoning Sebastian Raschka, PhD · April 19, 2025 Read full story
I will try to summarize what’s most essential to mention as background info below.
First, training a model with RLVR is already implicitly leading to a form of inference scaling, since reasoning models usually output more tokens during inference compared to conventional LLMs, and that means we are spending more compute during inference.
Second, we can further adjust this output length via reasoning effort levels, but more on that later.
​Third, there are many additional inference scaling techniques. A popular one is self-consistency, which is often implemented as a form of majority voting where the model is queried multiple times, and the final answer is selected via majority vote.
Figure 10: An example of self-consistency, a popular inference scaling technique.
This can be applied to conventional LLMs as well as reasoning models. Also, this method can be used on demand and in addition to reasoning training. A good example of that is DeepSeekMath-V2, where the researchers applied extreme inference-scaling on top of a reasoning model (specialized for math) to achieve state-of-the-art performance on challenging math olympiad-type problems.
Figure 11: Two types of inference scaling (self-consistency and self-refinement) used together to improve math performance. Figure adapted from DeepSeekMath-V2: Towards Self-Verifiable Mathematical Reasoning
But again, I will refer to my other article, The State of LLM Reasoning Model Inference for an overview of other techniques:
The State of LLM Reasoning Model Inference Sebastian Raschka, PhD · March 8, 2025 Read full story 3. Think tokens
You may have seen the <think></think> tokens in the earlier “Aha moments” figure. I also included the corresponding figure below so you don’t have to scroll all the way up.
Figure 12: Common formatting tokens in reasoning models.
These <think> and </think> tags are cosmetic with respect to reasoning ability. They do not make the model reason, and they are not required to achieve good reasoning performance. One could train the same model without these delimiters and likely reach similar benchmark performance.
The purpose of these <think> tags or tokens is mainly to mark where the reasoning trace begins and ends so that the training pipeline or user interface can separate it from the final answer and optionally hide it from the user. (UIs like ChatGPT or Codex usually do this.)
The point here is that the <think> tokens are not giving the model the ability to “think” or reason or reason better. One could train the same models without such <think> tokens and reach similar benchmark performance.
There is also nothing special about the literal strings <think> and </think> . Another pair of delimiters could serve the same purpose.
By the way, the way this is implemented is typically by adding a formatting reward during the RLVR stage. So instead of just rewarding the model based on answer correctness, one would provide additional reward for the use of <think> tokens, which in turn encourages the model to use those.
In DeepSeek-R1, for example, the overall reward was calculated as
R_total = R_accuracy + R_format
where the format reward was a simple rule-based check that encouraged the model to place its reasoning inside:
4. Reasoning mode on and off switches
The first generation of reasoning models was dedicated reasoning models. With that, I mean that there was a DeepSeek-V3 base model and a separate DeepSeek-R1 reasoning model.
No matter what the prompt is, R1 generally outputs very verbose responses using lots of tokens, even for simple prompts. It also lacks a built-in option to turn off the reasoning mode.
Figure 13: Reasoning models are very verbose, even for the simplest prompts.
Later models, like Qwen3 and others, experimented with hybrid approaches, where the same model can behave like a regular instruction fine-tuned model or a reasoning model on demand.
Note: Some model developers call this “thinking mode,” while others call it “reasoning mode.” Both terms refer to the same behavior.
In Qwen3, this is handled via the tokenizer using enable_thinking=True or enable_thinking=False . Under the hood, setting enable_thinking=False essentially adds an empty <think></think> section to the beginning of the assistant response to turn off Qwen3’s reasoning (”thinking”) mode.
Figure 14: Response of Qwen3 0.6B reasoning model with thinking=False and thinking=True . (The empty <think></think> tags are hidden in the interface on the left as they are part of the modified input prompt, not the generated answer.)
How is this implemented during training, such that the model supports this toggle during inference time, as shown in the figure above?
In short, as explained in the Qwen3 technical report , this on/off behavior is introduced primarily through supervised fine-tuning (SFT) and then reinforced during general RL in their largest flagship models.
For instance, after the initial reasoning model is trained via long-chain-of-thought SFT and reasoning RL, they add a “Thinking Mode Fusion” stage. During this additional SFT stage, the model sees both thinking and non-thinking examples:
/think: <think>

[truncated]
