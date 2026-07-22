---
source: "https://www.henrypan.com/blog/2026-07-18-harness-training/"
hn_url: "https://news.ycombinator.com/item?id=49014399"
title: "Training Agent Harness Like Training a ML Model"
article_title: "Training Agent Harnesses Like Model Weights | Henry Pan"
author: "megadragon9"
captured_at: "2026-07-22T22:57:52Z"
capture_tool: "hn-digest"
hn_id: 49014399
score: 1
comments: 0
posted_at: "2026-07-22T22:34:15Z"
tags:
  - hacker-news
  - translated
---

# Training Agent Harness Like Training a ML Model

- HN: [49014399](https://news.ycombinator.com/item?id=49014399)
- Source: [www.henrypan.com](https://www.henrypan.com/blog/2026-07-18-harness-training/)
- Score: 1
- Comments: 0
- Posted: 2026-07-22T22:34:15Z

## Translation

タイトル: ML モデルをトレーニングするようなトレーニング エージェント ハーネス
記事のタイトル: トレーニング エージェントのハーネスはモデルの重みのように |ヘンリー・パン
説明: 決定論的 LLM 推論上の PyTorch のようなループを使用してエージェント ハーネスをトレーニングした方法と、ターミナル ベンチ 2.0 と SWE ベンチのモデル間でゲインがどのように転送されるかについて説明します。

記事本文:
プロジェクトリポジトリ: https://github.com/workofart/harness-training
そこで私は最近、AI エージェントが端末ベンチのタスクを解決するハーネスを自己改善できるかどうかを確認したいと考えました。定義に合わせて説明すると、「ハーネス」とは、特定の環境と対話するモデル (例: GPT-5.5、Claude Opus 4.7 など) をラップするシステム (例: Claude Code、Codex、ChatGPT Web インターフェイスなど) を意味します。ハーネスは、モデルが何を認識するか、モデルが使用できるツール、環境応答がどのようにモデルにフィードバックされるかなどを制御します。
私の最初の試みは主に自己改善ループを実験計画の問題として扱うことでした。この学びについては、前回のブログ投稿「 1,000+ のハーネス実験が自己改善エージェントについて教えてくれたこと 」で書きました。
それ以来、私は実験の実行方法を完全に見直しました (ヒント: 決定論)。これにより、実験時間を短縮し、実験ごとにより多くの信号を抽出し、「トレーニング ハーネス」のための一般的な PyTorch のようなフレームワークを構築することができました。このブログ投稿では、これがどのように行われるのか、そしてそれが再帰的自己改善の文脈においてなぜ重要なのかについて説明します。
同じトレーニング済みハーネス + モデル X を評価した結果の一部を次に示します。
正確に定義するには、この問題は、端末報酬と永続的な検索メモリを使用したエージェント主導の離散プログラム検索として構成されます。
以下は、繰り返し登場する概念と、フレームワークのインスピレーションとなった PyTorch のアナロジーのマップです (後で紹介します)。
私のセットアップは非常に基本的なもので、主にコストを管理し、45 分以内に 1 つの実験 (30 ～ 40 タスク) の適切な反復速度に到達することが目的です。
クラウドで 5090 GPU をレンタルしました (実行している同時タスクの数に応じて 1 ～ 4)。 LLM 推論プロバイダーは決定性を保証しません。これについては後で説明します。
クウェン3

.6 35B A3B FP4 量子化モデル: モデルの重みは最大 24 GB を占有しますが、KV キャッシュ (基数/プレフィックス キャッシュが無効になっている) は最大 10 個の同時リクエストをサポートするために最大 4 GB を使用しますが、これは決定論的カーネル (後述) によって課される制限です。残りのメモリは、アクティベーションと CUDA オーバーヘッドによって消費されます。
LLM 推論構成: シード 12345、一時 1.0、top-p 0.95、コンテキスト長 32k、最大トークン数 8k、推論トークン数 6k。これらは主に GPU メモリによって制約され、最悪の場合の最大トークン数のデコード時間を 5 分未満にすることを目標としていました。また、GPU 間でモデルを分割したくありませんでした。
SGLang 推論エンジン: 決定論的なバッチ推論をサポートしています。これがハーネスのトレーニング (ブートストラップ スクリプト) の前提条件である理由については後で説明します。
次に、ハーネスをトレーニングした 2 種類のタスク、SWE ベンチ タスクとターミナル ベンチ タスクについて詳しく説明します。
ベースライン ハーネスをどのように定義するかについて言及することが重要です。これは、モデルのトレーニング中に機械学習モデルを初期化する方法に似ているためです。また、ハーネスを適切に初期化することが非常に重要です。そうしないと、ハーネス トレーニングで適切な学習信号 (勾配) を見つけるのが非常に困難になります。たとえば、ベースライン ハーネスが細すぎる場合、LLM 応答処理ロジックがないと、トレーニングのほとんどが LLM 呼び出しのエラー処理の作成に費やされることがわかりました。もう 1 つの極端な例は、ベースライン ハーネスが 10 個以上のツール定義で過負荷になっている場合で、その時点では偏りが多すぎます。おそらく特定のモデルはそれを使用してトレーニングされていますが、それが一般的な出発点ではないことは明らかです。
私のベースライン ハーネスでは、ツールを使用するために合計 2 つのアクションを定義しました。シェル コマンドを実行するための Run と、ハーネスに準備ができたことを知らせるための Submit です。

タスクベリファイアを呼び出して結果を確認します。これは決して良いベースラインではありませんが、クリーンでありながら学習可能な状態です。
39 の SWE ベンチ タスク (django/pylint/pytest/sphinx/sympy) でトレーニング済み 1
タスクあたり最大 110 ステップまたはタスクあたり 75 分のタイムアウトに制限されます
凍結タスク LLM: Qwen 3.6 35B A3B FP4 量子化モデル
ハーネスのトレーニングに src/trainer/estimator.py:AgenticEstimator として使用されるエージェント: Codex CLI 経由で GPT-5.5 高
トレーニングからのすべての候補コミットを含む、この実験の Git ブランチ
結果: 29 回の実験実行で 8/39 -> 14/39 が解決されました: 8 件が昇格、20 件が却下、1 件がベースライン。
二次報酬は、候補者とそのベースラインがまったく同じタスクを解決した場合にのみ使用されるタイブレーカーです。
39 のタスクのうち 21 のタスクは、29 回の実行のいずれでも解決されませんでした。実験成果物に基づくと、そのうちの 3 分の 1 が 110 ステップの上限に達しました。これはトレーニングの反復を高速化するために強制されました。ただし、ここには明らかにトレードオフがあり、ステップ キャップを拡張すると、より多くのトレーニング信号を収集できますが、各エポックの実行時間が長くなります。あるいは、モデルの能力に上限があるか、LLM 推論リソースが制限されすぎている可能性があります 2 。
ターミナルベンチタスクのトレーニング
38 のターミナル ベンチ 2.0 タスクでトレーニングされました。特定のモデルではサイバー/暗号化タスクの完了を拒否するため、これらのタスクはサイバー/暗号化タスクを削除するために選択されました。そして、多様性と難易度の両方を最大化し、タスクの完了時間を最小限に抑えるために、サブセットに絞り込まれました。
タスクあたり最大 110 ステップまたはタスクあたり 30 分のタイムアウトに制限される
凍結タスク LLM: Qwen 3.6 35B A3B FP4 量子化モデル
ハーネスのトレーニングに src/trainer/estimator.py:AgenticEstimator として使用されるエージェント: Codex CLI 経由で GPT-5.5 高
トレーニングからのすべての候補コミットを含む、この実験の Git ブランチ
2 段階の実験の実行により、非決定的要因が実験の妨げとなることが判明

LF改善能力
以下のプロットでは、実験結果を汚染したいくつかの非決定的な環境要因を特定したため、ベースラインをリセットし、26 回目の実験あたりでプロモートされたハーネスの変更をすべて削除したことがわかります。リセット前に拒否された候補の分散は非常に狭いものの、安定していることもわかります。
非決定性を修正して再起動した後、ハーネスは適切な変更をすぐに見つけて、実験時間の半分未満で以前の最高スコアを上回ることができました。
タイブレーカー プロモーションに関する注意: ハーネス トレーニング フレームワークを使用すると、ユーザーは第 1 の損失関数のタイを破るために第 2 の報酬を定義できます。上記のトレーニング実行では、二次報酬は次のように定義され、順番に評価されます。
最初の試行で有効なツール呼び出しの割合
タスクを解決するために使用されるステップ数
最大のジャンプ ( 14ef7381 、 17 から 22 へのソルブ): LLM 出力応答バジェット全体がツール呼び出しを発行せずに推論に消費され、以前はトライアルが無効でした。提案されたハーネスの変更は、思考を無効にし、以下の修復プロンプトを挿入することでした。変換された 5 つの解決済みタスクはすべて長期的なものでした (67 ～ 90 ステップで解決)。
reasoning_runaway_repair_prompt = (
"前の応答は出力長の制限に達しましたが、"
"ツール呼び出しを発行する前に推論しました。今すぐ推論を停止して返信してください。"
「単一の簡潔なツール呼び出しで、他のテキストはありません。」
)
評価結果
ハーネスのトレーニングは実際に行う価値があるのでしょうか?このセクションでは、トレーニングされたハーネスを固定し、基礎となるタスク LLM モデルを交換することによって、いくつかの評価実行を実行します。評価構成は次のとおりです。
Terminal Bench 2.0 でトレーニング済み 38 タスク 89 タスク中 3 タスク。トレーニングされたすべてのハーネス コミットを含む Git ブランチ。
評価する

Terminal Bench 2.0 では、89 個のタスクすべてが実行されます。トレーニング セットと評価セットの間に重複があることは明らかです。 38 のタスクのうち 23 は、トレーニング済みハーネスを備えた Qwen 3.6 35B A3B FP4 量子化モデルによって解決されましたが、タスク固有の「不正行為」メカニズムは含まれていませんでした。また、同じターミナル ベンチでトレーニングされたハーネスを SWE ベンチと比較して評価すると、転移学習の利点があることも示します。
一般化された転移学習機能
これは、1 つのモデル (Qwen 3.6 35B A3B FP4 量子化モデル) を使用して一度トレーニングされたフリーズ ハーネスの場合、さまざまなモデルの機能を向上させることができることを示しています。
また、あるタスク セット (SWE ベンチ) 1 でハーネスをトレーニングすると、その同じハーネスが別のタスク セット (ターミナル ベンチ) 4 に向けてタスク解決能力を向上させることができることも示しています。
訓練されたハーネスには、より優れたタスク解決能力に貢献するメカニズムがいくつかあります。
以下のプロットでは、次のことがわかります。
トレーニングされたシステム プロンプトが追加されたため、「検証失敗」が「解決済み」(+23) に変換されました。
「実行出力はスクラッチです。Python スニペット内のファイルまたは印刷のみのヒアドキュメントを編集しても、修正は保持されません。」
「送信時のディスク上の状態が評価される唯一のものです」
これは、モデルが一時的な出力にコードを書き込み、変更されていないディスク状態を送信するベースライン障害モードを対象としています。変換された 23 件の「解決済み」タスク内では、新しいファイル ツールが使用されました (約 1,200 件の呼び出しのうち約 300 件)。 DeepSeek 3.2 が最大の恩恵を受けています (+7)。
トレーニングされたハーネスがベースラインの LLM レイテンシの半分以下を費やしたことから、「時間切れ」が「解決済み」（+24）に変換されました。具体的には、GPT-5.5 ヘッドレス ターミナル タスクは 1755 秒から 221 秒に短縮され、DeepSeek コンパイル コンプサートは 1538 秒から 165 秒に短縮されました。訓練されたハーネスメカニズム

貢献したのは次のとおりです。
環境 (stdout/stderr) の肥大化を抑制するために読み取りが 250 行にクリップされました
いくつかのモデルで短い完了 (呼び出しごとの出力トークン: GPT-5.5 -19%、MiniMax2.5 -22%、GPT-OSS 20B -11%)
「実行」コマンド/ツールの呼び出しがタイムアウトした場合は、LLM に「再実行ではなくアプローチを変更してください」という指示を出し、300 秒の停止の繰り返しを回避します。
「有効なツール呼び出しが繰り返し存在しません」は「解決済み」 (+6) に変換されました。これは、トレーニングされたハーネスにさらに多くのファイル ツールが導入されたためです (既存の「読み取り」、「送信」に「読み取り」、「書き込み」、「置換」が追加されました)。より具体的には、永続化する「書き込み」と「置換」は、base64 経由で引用安全に編集します。これにより、バイト同一の修復プロンプトにもかかわらず、「ツール呼び出しの欠落」5 の回復が 69% から 94% に向上しました。無効なコール 6 の場合、トレーニングされたハーネス修復は失敗したコールバックを「送信しました: {calls}」のようなものとしてエコーし、これにより ValidationError 6 の回復率が 51% から 83% に向上します。
上の図と下の図を組み合わせると、「hit_timeout」から「solved」への変換により、以前は実時間内に実行されていたのと同じタスクが効率によって解決されているため、有機的にステップ数が増加します。 MiniMax 2.5 では、トレーニング済みハーネスとベースライン ハーネスの両方で解析失敗がほぼゼロですが、それでも 50 ステップを超えても +7 のソルブが得られます。 5 つはタイムアウト/検証の反転であり、解析ではなく「永続的な修正」 7 と書き込み/置換 8 から利益が得られていることを示唆しています。 GPT-OSS 20B では、致命的なステップ率が 9.4% から 1.9% に低下し、軌道がより長く実行できるようになり (試行ステップ数が 5.6k から 9.5k に増加)、一部のランで変換が遅くなることが可能になりました。ステップごとのツール呼び出し失敗率は、基本的には変わらず、最大 31% でした。これは、訓練されたハーネスでは障害が存続することを意味します。

生き残ったステップでは、43% 多くの出力トークンが請求されます。もう 1 つのモデル ファミリーの長期にわたる成功は、「解決」メカニズムの「タイムアウト」と「検証の失敗」によってもたらされます。 GPT-OSS 120B は、軌道が縮小しながらソルバを獲得します。これは、各モデルが実際には異なる故障モデルやタスク解決戦略を重みに組み込んでいることを再度示しています。
「タスクの解決」を \(\text{提出済み} \wedge \text{検証に合格}\) として数学的に分解すると、トレーニングされたハーネスはこれら 2 つのレンズを通してのみ目的を達成でき、それが実際に行われたことになります。 「より良い提出」は、永続的な修正 7 に加えて、トレーニングされたハーネスのより長く、より慎重な軌道によって、弱い提出をプールから除外することによって実現されます。 DeepSeek 3.2 では、失敗した提出物が提出にまったく到達しなくなったため、合格率も上昇しました。 MiniMax 2.5 のみが、より低い合格率を犠牲にして、より多くの提出物を提出することで解決します。
本物のトークン効率を高めるには、訓練されたハーネスがより多くのファイル ツールを導入したため、ツール呼び出しの組み合わせが安価な呼び出しにシフトされました。 GPT-5.5 のトレーニング済みハーネス評価では、読み取りの出力トークンは平均 433 個でしたが、トレーニングされていないベースライン ハーネス評価の場合は 1,013 個で、約 2,200 件の呼び出しのうち 374 個が読み取り/置換に移動されました (呼び出しあたりのトークンは 1,422 から 1,151 に減少しました)。
ここには基本的に 2 つのキャンプがあります。
GPT-OSS 20B を代表として、そのコンテキストの爆発は訓練されたハーネスによって救われます。

[切り捨てられた]

## Original Extract

How I trained agent harnesses with a PyTorch-like loop over deterministic LLM inference, and how the gains transfer across models on Terminal Bench 2.0 and SWE-bench.

Project Repository: https://github.com/workofart/harness-training
So I recently wanted to see whether an AI agent could self-improve a harness to solve terminal bench tasks. To align on the definitions, “harness” means the system (e.g. Claude Code, Codex, ChatGPT web interface etc…) wrapping around the model (e.g. GPT-5.5, Claude Opus 4.7 etc…) that interacts with a specific environment. The harness controls what the model sees, what tools the model can use, and how environment responses are fed back to the model etc…
My initial attempt was mostly treating the self-improvement loop as an experiment design problem. I wrote about the learnings in my previous blog post What 1,000+ Harness Experiments Taught Me About Self-Improving Agents .
Since then I’ve completely revamped the way I run experiments (hint: determinism). This allowed me to reduce the experiment time, to extract more signal per experiment and to build a general PyTorch-like framework for “training harnesses”. In this blog post, I’m going to talk about how this is done and why it matters in the context of recursive self-improvement.
Here are some results from evaluating the same trained harness + model X:
For a precise definition, this problem is framed as an agent-guided discrete program search with terminal reward and persistent search memory.
Here’s a map of the recurring concepts and the PyTorch analogy that inspired the framework (introduced later ):
My setup is pretty basic, mostly to control costs and reach a decent iteration speed of one experiment (30 - 40 tasks) within ~45 minutes:
Rented 5090 GPUs in the cloud (1 - 4 depending on how many concurrent tasks I’m running). The LLM inference providers don’t guarantee determinism, which I will explain later .
Qwen 3.6 35B A3B FP4 quantized model : The model weights occupy ~24 GB, while the KV cache (with radix/prefix caching disabled) uses ~4 GB to support up to 10 concurrent requests, a limit imposed by the deterministic kernel(discussed later). The remaining memory is consumed by activations and CUDA overhead.
LLM inference configurations: seed 12345, temp 1.0, top-p 0.95, 32k context length, max token count 8k, reasoning token count 6k. These were mostly constrained by GPU memory and aiming decoding time of below 5 minutes for the worst-case max token count, and I didn’t want to split the model across GPUs
SGLang inference engine: it supports deterministic batch inference, which I will explain later why it’s a prerequisite for training the harness ( bootstrap script )
Next, I will dive into the two types of tasks where I trained my harness: SWE bench tasks and Terminal Bench tasks.
It’s important to mention how I define my baseline harness because this is kind of like how one would initialize a machine learning model during model training. And it is very important to initialize your harness properly because otherwise it’s going to be very hard for the harness training to find good learning signals (gradient). For example, I’ve seen that if the baseline harness is too thin, without the LLM response handling logic, most of the training will be spent on creating error handling for the LLM calls. The other extreme is when the baseline harness is overloaded with 10+ tool definitions, at that point there’s too much bias. Maybe certain models are trained with that, but it’s definitely not a generic starting point.
In my baseline harness , I defined two actions in total for tool use: Run for running any shell command and Submit for letting the harness know it’s ready to call the task verifier to check the results. By no means is this a good baseline, but it’s a clean yet learnable state.
Trained on 39 SWE-bench tasks (django/pylint/pytest/sphinx/sympy) 1
Capped at max 110 steps per task or 75 mins timeout per task
Frozen Task LLM: Qwen 3.6 35B A3B FP4 quantized model
Agent used as src/trainer/estimator.py:AgenticEstimator for training the harness: GPT-5.5 high via Codex CLI
Git branch for this experiment with all the candidate commits from training
Result: 8/39 -> 14/39 solved over 29 experiment runs: 8 promoted, 20 rejected, 1 baseline.
The secondary reward is a tie-breaker used only when a candidate and its baseline solve exactly the same tasks.
There were 21 tasks out of 39 tasks that were never solved in any of the 29 runs. Based on experiment artifacts, a third of them hit the 110 step-cap, which I enforced to speed up the training iterations. But there’s definitely a trade-off here where if I extend the step cap, we can collect more training signals, but each epoch will run longer. Or maybe there’s some model capability ceiling or my LLM inference resources were too limited 2 .
Training on Terminal Bench Tasks
Trained on 38 terminal bench 2.0 tasks 3 . The tasks were selected to remove cyber/cryptography tasks since certain models refuse to complete it. And it was narrowed down to a subset to maximize both variety/difficulty and minimize task completion time.
capped at max 110 steps per task or 30 mins timeout per task
Frozen Task LLM: Qwen 3.6 35B A3B FP4 quantized model
Agent used as src/trainer/estimator.py:AgenticEstimator for training the harness: GPT-5.5 high via Codex CLI
Git branch for this experiment with all the candidate commits from training
Two-stage experiment runs show non-deterministic factors hinder the self-improving capability
You can see in the plot below that I reset the baseline and dropped all the promoted harness changes around the 26th experiment due to identifying some non-deterministic environment factors that polluted the experiment results. You can even see that the rejected candidates before the reset had a very narrow but steady variance.
After the non-determinism fix and restart, the harness was able to quickly find good changes and beat the previous best score in less than half of the experiment time.
Note on the tie-breaker promotion: the harness training framework allows users to define secondary reward to break primary loss function ties. In the above training run, the secondary reward is defined as the following, evaluated in order:
% of tool calls that are valid on the first attempt
the number of steps used to solve the task
Biggest jump ( 14ef7381 , 17 to 22 solves): entire LLM output response budget burned on reasoning with no tool call emitted, which previously killed the trial. The harness change proposed was to disable thinking and inject the repair prompt below. The 5 converted solved tasks were all long-horizon (solved in 67-90 steps).
reasoning_runaway_repair_prompt = (
"Your previous response reached the output length limit while still "
"reasoning, before it emitted any tool call. Stop reasoning now and reply "
"with a single, concise tool call and no other text."
)
Evaluation Results
Is training the harness actually worthwhile to do? In this section, I’m going to go through some of the evaluation runs by holding the trained harness fixed and swapping out the underlying task LLM model. Here are the evaluation configurations.
Trained on Terminal Bench 2.0 38 tasks 3 out of 89 tasks. Git branch with all the trained harness commits.
Evaluated on Terminal Bench 2.0 all 89 tasks. It’s obvious that there is overlap between the training set and the evaluation set. 23 out of 38 tasks were solved by Qwen 3.6 35B A3B FP4 quantized model with the trained harness, and it didn’t include task-specific mechanisms to “cheat”. We will also show that evaluating the same Terminal Bench-trained harness against SWE-bench, there are transfer learning gains too.
Generalized and transfer learning capability
This shows that for a frozen harness that was trained once using one model ( Qwen 3.6 35B A3B FP4 quantized model ), it can lift up the capability of a wide variety of models.
It also shows that if we train a harness on one task set (SWE-Bench) 1 , that same harness can improve task-solving capability towards a different task set (Terminal Bench) 4 .
There are a couple of mechanisms in the trained harness that contribute to better task-solving abilities
In this plot below, you can see:
“failed verification” converted to “solved” (+23) because the trained system prompt added
“run output is scratch: editing a file inside a python snippet or a here-doc that only prints does NOT persist your fix”
“the on-disk state at submit is the only thing graded”
This targets a baseline failure mode where a model wrote the code in some ephemeral output and submitted an unchanged disk state. Inside the 23 converted “solved” tasks, they used the new file tools (~300 of ~1200 calls). DeepSeek 3.2 is the biggest beneficiary (+7).
“ran out of time” converted to “solved” (+24) from the trained harness spending half or less of the baseline’s LLM latency. Specifically, GPT-5.5 headless-terminal task reduced from 1755s to 221s, DeepSeek compile-compcert reduced from 1538s to 165s. The trained harness mechanisms that contributed were:
read clipped to 250 lines to cap environment (stdout/stderr) bloat
shorter completions for several models (output tokens per call: GPT-5.5 −19%, MiniMax2.5 −22%, GPT-OSS 20B −11%)
if the “Run” command/tool call times out, prompt the LLM with a nudge “change your approach rather than rerunning”, which avoids repeated 300s stalls.
“no valid tool call repeatedly” converted to “solved” (+6), because in the trained harness, more file tools were introduced (added “read”, “write”, “replace” to the existing “read”, “submit”). More specifically, the “write” and “replace” that persist edits quote-safely via base64, which improved “missing tool call” 5 recovery from 69% to 94%, despite a byte-identical repair prompt. For invalid calls 6 , the trained harness repair echoes the failed call back as something like “You sent: {calls}”, which improves a ValidationError 6 recovery from 51% to 83%.
If we combine the above figure with the below figure, the “hit_timeout” to “solved” conversion would organically cause high step counts, since the same tasks that previously ran out of wall-clock are now being solved, which is driven by efficiency. MiniMax 2.5 has near-zero parse failures in both the trained and baseline harnesses, yet still gains +7 solves beyond 50 steps. Five are timeout/verification flips, suggesting the gains come from “persisting fixes” 7 and write/replace 8 , not parsing. For GPT-OSS 20B, the fatal step rate fell from 9.4% to 1.9%, allowing trajectories to run longer (attempted steps increased from 5.6k to 9.5k) and enabling some runs to convert late. Tool call failures per step remained essentially unchanged at ~31%. This means that failures survive in the trained harness, and the surviving steps bill 43% more output tokens. The other model family’s long-trajectory wins come from the “timeout” and “failed verification” to “solve” mechanisms. GPT-OSS 120B gains solves while its trajectories shrink. This again shows that each model really has a different failure model or task-solving strategy baked into the weights.
If we decompose “Task Solve” mathematically as \(\text{Submitted} \wedge \text{Passed Verification}\), the trained harness could only achieve its goal through those two lenses, and that’s what it did. “Submit better” comes from persisting fixes 7 , plus the trained harness’s longer, more careful trajectories filtering weak submissions out of the pool. For DeepSeek 3.2, pass rate also rose as failing submissions stopped reaching submit at all. Only MiniMax 2.5 netted solves from submitting more, at the cost of a lower pass rate.
For the genuine token efficiency wins, since the trained harness introduced more file tools, it shifted the tool call mix toward cheap calls. In GPT-5.5’s trained harness evaluation, read averages 433 output tokens vs 1,013 for the untrained baseline harness evaluation, and 374 of ~2,200 calls moved to read/replace (tokens per call 1,422 decreased to 1,151).
There are essentially two camps here:
With GPT-OSS 20B as the representative, its context blowup is rescued by the trained harness s

[truncated]
