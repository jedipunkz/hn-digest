---
source: "https://machinelearningmastery.com/the-roadmap-to-mastering-ai-agent-evaluation/"
hn_url: "https://news.ycombinator.com/item?id=48584352"
title: "The Roadmap to Mastering AI Agent Evaluation"
article_title: "The Roadmap to Mastering AI Agent Evaluation"
author: "eigenBasis"
captured_at: "2026-06-18T13:17:52Z"
capture_tool: "hn-digest"
hn_id: 48584352
score: 3
comments: 0
posted_at: "2026-06-18T12:30:54Z"
tags:
  - hacker-news
  - translated
---

# The Roadmap to Mastering AI Agent Evaluation

- HN: [48584352](https://news.ycombinator.com/item?id=48584352)
- Source: [machinelearningmastery.com](https://machinelearningmastery.com/the-roadmap-to-mastering-ai-agent-evaluation/)
- Score: 3
- Comments: 0
- Posted: 2026-06-18T12:30:54Z

## Translation

タイトル: AI エージェント評価をマスターするためのロードマップ
説明: この記事では、AI エージェントの最終出力だけではなく、実行プロセス全体を調べることによって、AI エージェントを厳密に評価する方法を学びます。

記事本文:
AI エージェントの評価をマスターするためのロードマップ
ナビゲーション
開発者を機械学習で優れたものにする
開発者を機械学習で優れたものにする
コード アルゴリズム 機械学習アルゴリズムを最初から実装します。
ディープラーニング (keras) ディープラーニング
時系列予測のためのニューラルネット時系列ディープラーニング
LSTM 長短期記憶ネットワーク
開発者を機械学習で優れたものにする
開発者を機械学習で優れたものにする
コード アルゴリズム 機械学習アルゴリズムを最初から実装します。
ディープラーニング (keras) ディープラーニング
時系列予測のためのニューラルネット時系列ディープラーニング
LSTM 長短期記憶ネットワーク
AI エージェントの評価をマスターするためのロードマップ
シェアする
ポスト
シェアする
この記事では、AI エージェントの最終出力だけではなく実行プロセス全体を調べることで、AI エージェントを厳密に評価する方法を学びます。
エージェントの評価が従来の言語モデルの評価と異なる理由と、エージェントが推論層とアクション層全体で失敗する場所。
構築しているエージェントのタイプに合わせて、決定論的なコードベースのチェックとモデルベースの判定を使用してエージェントを採点する方法。
pass@k や pass^k などのメトリクスを使用して非決定論を説明する方法、および評価を開発から運用監視まで拡張する方法。
AI エージェントの評価をマスターするためのロードマップ
これ以上時間を無駄にしないようにしましょう。
AI エージェントを構築している多くのチームは依然として、大規模な言語モデルを評価するのと同じ方法で AI エージェントを評価しています。つまり、いくつかのタスクを実行し、最終出力を検査し、すべてが機能していると想定しています。このアプローチでは、最も重要な失敗を見逃してしまうことがよくあります。モデルは不適切なツールを選択したり、誤ったツール引数を生成したりする可能性がありますが、エージェント システムはツールの障害をうまく処理できなかったり、非効率的な一連のアクションに従う可能性があります。評価する

最終的な応答のみでは、これらの障害がどこで発生したかを特定することが困難になることがよくあります。
エージェントの評価は、このギャップに対処します。結果のみに焦点を当てるのではなく、タスクの展開に応じてエージェントがどのように推論し、決定を下し、ツールを使用し、適応するかという、実行プロセス全体を調査します。これにより、信頼性、効率、全体的なパフォーマンスをより正確に把握できるようになり、チームが本番環境に入る前に問題を特定できるようになります。
この記事で説明する原則は、エージェントのパフォーマンスを測定および改善するための体系的なアプローチの基礎を形成します。
ロードマップ (クリックして拡大)
ステップ 1: エージェントの評価が重要な理由を理解する
エージェントが失敗すると、本能的にそれをプロンプトの問題として扱います。システム プロンプトをより明確にする必要があります。時々それは真実です。多くの場合、失敗は測定の問題です。評価は何が壊れたのかを把握するように設計されていませんでした。
AI エージェントは複数のレイヤーにわたって動作するため、これらのレイヤーは独立して失敗する可能性があります。
言語モデルを利用した推論層は、計画、タスクの分解、ツールの選択を処理します。
アクション層はツール呼び出しと外部システム応答を利用して実行を処理します。
エージェントは何をすべきかを正しく推論し、不正な形式の引数を使用して適切なツールを呼び出すことができます。エージェントの評価を単一のエンドツーエンドの精度チェックとして扱うと、両方の障害面が見逃されます。
コンポーネントレベルの評価: 推論の品質とツール呼び出しの精度を個別に測定
エンドツーエンド評価: タスクの完了と実行効率を測定
タスク完了率が 80% ということは、20% の失敗が不適切な計画、間違ったツール選択、間違った引数、またはツール インフラストラクチャの障害によるものであるかどうかについては何も示しません。ステップレベルのトレース - 各ツール呼び出し、その引数、その引数をキャプチャするログ

結果とその後のモデルの決定により、診断が可能になります。トレースがなければ、本番環境の障害のデバッグは推測でしかできません。
ステップ 2: エージェント評価の成功とはどのようなものかを定義する
評価はその成功基準に応じて決まります。適切に形成された評価タスクとは、独立して作業する 2 人のドメイン専門家が同じ合否判定に達するものです。
明確なタスク仕様と参照ソリューション (すべての採点者に合格する既知の正しい出力) を組み合わせたものから始めます。タスクが解決可能であることを証明し、採点ロジックが正しく構成されていることを検証します。
採点を実行する前に、 eval に対して以下を定義する必要があります。
タスク: エージェントが受け取る入力、エージェントが行うことを期待されること、および環境がどのように見えるか
成功基準: 最終的な答えだけでなく、重要な中間結果: 適切なツールが呼び出されましたか?状態は正しく更新されましたか?応答は取得されたコンテキストに基づいていましたか?
否定的な場合: 片側評価は片側最適化を作成します。バランスのとれたデータセット — 動作が発生すべき時期と発生すべきでない時期の両方をカバーし、エージェントが機能に対して過剰にトリガーしたり、トリガーが不十分になったりすることを防ぎます
実際の使用上の失敗から導き出された、明確に指定された一連のタスクは、完璧なデータセットを待つよりも優れた開始点です。待機時間が長くなるほど、評価の構築は難しくなります。
ステップ 3: コードベースのチェックによるエージェント アクション レイヤーのグレーディング
決定的グレーダー (モデルインザループの判断を行わずに特定の条件をチェックするコード) は、エージェント評価スタックの中で最も高速かつ安価で、最も再現性の高いオプションです。アクション レイヤーの場合、これらが常に開始点となる必要があります。
ツール呼び出しの検証: エージェントが正しいツールを正しい順序で呼び出したかどうか
引数の検証:

r 入力に正しい型、必要なパラメータ、および有効な値がある
結果の検証: 環境が期待した状態で終了するかどうか
トランスクリプト分析: ターン数、消費されたトークン、レイテンシー
これらは多くの場合、高速で客観的でデバッグが簡単ですが、脆弱です。採点者が「confirmation_code」:「CONF-789」をチェックすると、同じデータを異なる形式で表示する正しい応答が見逃されます。
ステップ 4: モデルベースの審査員によるエージェントの推論と出力品質の採点
エージェントの評価項目の中には、出力の品質、トーン、取得されたコンテキストへの忠実さ、適切な共感など、決定論的なチェックに抵抗するものがあります。これらの場合、審査員または LLM-as-a-Judge として使用される言語モデルが適切なツールです。柔軟性があり、無制限の出力を処理できますが、コードベースの採点者にはない非決定性とキャリブレーション ドリフトが生じます。
次の実践により、モデルベースのグレーダーの信頼性が維持されます。
構造化されたルーブリックを作成します。 「回答が役に立つかどうかを評価する」ということはノイズを生みます。応答がユーザーの質問に対処し、取得された文脈で主張を根拠付け、範囲外の提案を回避する必要があることを指定するルーブリックがシグナルを生成します。各次元を個別の独立した判断で採点します。
人間の判断に照らして定期的に調整してください。審査員としての LLM の精度は、分野の専門家によって採点されたサンプルと比較してチェックする必要があります。相違が見られる場合、ほとんどの場合、ルーブリックが問題となります。曖昧なケースに対する強制的な判断を避けるために、採点者に明示的な「判断できない」オプションを与えます。
複数のコンポーネントからなるタスクの部分的なクレジットを組み込みます。問題を正しく特定し、顧客を確認したにもかかわらず、返金処理に失敗したサポート エージェントの方が、ステップ 1 で失敗したサポート エージェントよりもはるかに優れています。バイナリの合格/失敗は、エージェントが実際に実行している場所を隠します。

私。
ステップ 5: エージェント評価戦略をエージェント タイプに適合させる
グレーディング戦略は広範囲に適用されますが、エージェントのタイプによって、どのグレーダーが最も重みを持ち、どの故障モードを優先するかが決まります。
コーディング エージェントはコードを作成、テスト、デバッグします。ソフトウェアは主に決定論的です。コードは実行されるか、テストは合格するか、修正によって既存の機能を損なうことなく問題が解決されるか? SWE-bench Verified や Terminal-Bench などのベンチマークは、この合否アプローチに従っており、セキュリティ、読みやすさ、エッジ ケースの処理に関するルーブリック ベースの品質チェックによって補完されています。
会話型エージェントは、サポート、販売、コーチングのワークフロー全体でユーザーと対話します。インタラクションの質は評価対象の一部です。チケットが解決されたかどうかだけでなく、口調が適切であったかどうか、解決策が明確に説明されていたかどうかも評価されます。これには、ユーザーをシミュレートする第 2 言語モデルが必要です。 τ-bench はまさにこれをモデル化し、採点者がタスクの完了とターン全体のインタラクションの質の両方を評価します。
調査エージェントは、さまざまなソースから情報を収集し、総合します。グラウンディングチェックでは、取得したソースによって主張が裏付けられているかどうかを確認し、カバレッジチェックでは適切な回答に何を含める必要があるかを定義し、ソース品質チェックではエージェントが信頼できる資料を参照したことを確認します。
エージェントのタイプに応じたエージェント評価戦略のマッチング
ステップ 6: エージェントの評価結果における非決定性の考慮
エージェントの動作は実行ごとに異なります。同じタスク、同じ入力、同じエージェントが、異なるツールの選択、推論パス、および結果を生み出す可能性があります。したがって、単一試行の評価は、単純な精度指標では捕捉できない変動性を隠してしまうため、誤解を招く可能性があります。
これは、エージェント システムにおける非決定性の直接的な結果です。確率的モデルの出力、ツールの遅延、部分的な障害、および広告

適切な意思決定はすべて、実行ごとにばらつきをもたらします。その結果、エージェントを評価するには、単一の実行トレースではなく、結果の分布を推論する必要があります。
この変動を考慮するために、 pass@k や pass^k などのメトリクスが一般的に使用されます。
pass@k: k 個の独立した試行のうち少なくとも 1 つが成功する確率。複数の試行が許容される場合に役立ちます。
pass^k: k 回の試行がすべて成功する確率。すべてのインタラクションが信頼できる必要がある場合に重要です。
たとえば、1 回の試行成功率が 75% のエージェントは、3 回の試行すべてで成功するのはわずか約 42% であり、繰り返し実行すると信頼性がいかに急速に低下するかを示しています。
pass@k と pass^k
ステップ 7: エージェントの能力評価を回帰スイートから分離する
能力評価は、「このエージェントは、以前はできなかったことは何ですか?」という将来を見据えた質問に答えるように設計されています。そのため、比較的低い合格率から始めて、システムにとってまだ困難なタスクに集中する必要があります。能力評価が非常に高いスコア (たとえば 90%) に達すると、多くの場合、能力を測定するのではなく、単にすでに解決された問題の信頼性を確認することになります。
回帰評価は別の目的を果たします。彼らは、エージェントが以前に実行できたすべてのことを引き続き実行できるかどうかを尋ねます。これらのテストは 100% 近くで実行され、パフォーマンスの低下に対する保護として機能する必要があります。スコアが大幅に低下した場合は、何かが壊れていることを示しているため、リリース前に調査する必要があります。
時間が経つにつれて、エージェントにとっての能力評価は自然に容易になります。合格率が上昇し、パフォーマンスが安定すると、それらのタスクを回帰スイートに昇格させることができます。ただし、スイートが完全に飽和状態になると、実際の改善に対する感度が低くなります。つまり、意味のあることです。

l 進行状況は信号ではなくノイズとして現れる場合があります。このため、新しくてより挑戦的な eval は、既存のスイートが飽和した後ではなく、飽和する前に導入する必要があります。
ステップ 8: エージェントの評価を実稼働監視に拡張する
開発評価では、失敗すると予想されるものをキャプチャします。生産は実際に何をするのかを明らかにします。実際のユーザーは、合成テスト スイートにはめったに表示されない入力、エッジ ケース、およびコンテキストを導入するため、本番環境の監視が評価の必要な拡張になります。
完全な評価システムは、いくつかの相補的な信号を組み合わせます。
これらの層を組み合わせることで、実際のエージェントのパフォーマンスのより完全なビューが形成されます。ステップレベルのトレース (ループ内の各ポイントで推論、ツール呼び出し、引数、結果、意思決定をキャプチャー) は、これらすべてを機能させるインフラストラクチャです。 LangSmith 、Arize Phoenix 、Braintrust 、Langfuse などのツールは、トレースおよび評価フレームワークを提供します。 Harbor と DeepEval はハーネス層を処理します。
主要なエージェントの評価手順の概要
これまで説明してきた手順の概要を以下に示します。
次のステップとして、Anthropic の AI エージェントの評価を解明するガイド、特に「ゼロから 1 へ: エージェントの優れた評価へのロードマップ」セクションを読んでください。
シェアする
ポスト
シェアする
このトピックの詳細
エージェントの評価: テストと測定の方法

[切り捨てられた]

## Original Extract

In this article, you will learn how to evaluate AI agents rigorously by examining their full execution process rather than only their final outputs.

The Roadmap to Mastering AI Agent Evaluation
Navigation
Making developers awesome at machine learning
Making Developers Awesome at Machine Learning
Code Algorithms Implementing machine learning algorithms from scratch.
Deep Learning (keras) Deep Learning
Neural Net Time Series Deep Learning for Time Series Forecasting
LSTMs Long Short-Term Memory Networks
Making developers awesome at machine learning
Making Developers Awesome at Machine Learning
Code Algorithms Implementing machine learning algorithms from scratch.
Deep Learning (keras) Deep Learning
Neural Net Time Series Deep Learning for Time Series Forecasting
LSTMs Long Short-Term Memory Networks
The Roadmap to Mastering AI Agent Evaluation
Share
Post
Share
In this article, you will learn how to evaluate AI agents rigorously by examining their full execution process rather than only their final outputs.
Why agent evaluation differs from traditional language model evaluation, and where agents fail across the reasoning and action layers.
How to grade agents with deterministic code-based checks and model-based judges, matched to the type of agent you are building.
How to account for non-determinism using metrics like pass@k and pass^k, and how to extend evaluation from development into production monitoring.
The Roadmap to Mastering AI Agent Evaluation
Let’s not waste any more time.
Many teams building AI agents still evaluate them the same way they evaluate large language models : run a few tasks, inspect the final output, and assume everything is working. That approach often misses the failures that matter most. The model may select an inappropriate tool or generate incorrect tool arguments, while the agent system may handle tool failures poorly or follow an inefficient sequence of actions. Evaluating only the final response often makes it difficult to identify where these failures occurred.
Agent evaluation addresses this gap. Rather than focusing solely on outcomes, it examines the full execution process — how an agent reasons, makes decisions, uses tools, and adapts as a task unfolds. This provides a more accurate picture of reliability, efficiency, and overall performance, helping teams identify issues before they reach production.
The principles covered in this article form the foundation of a systematic approach to measuring and improving agent performance.
The Roadmap ( click to enlarge )
Step 1: Understanding Why Agent Evaluation Is Important
The instinct when an agent fails is to treat it as a prompting problem: the system prompt needs to be clearer. Sometimes that is true. More often the failure is a measurement problem: the eval was not designed to catch what broke.
AI agents operate across layers, and those layers may fail independently:
The reasoning layer — powered by the language model — handles planning, task decomposition, and tool selection.
The action layer — powered by tool calls and external system responses — handles execution.
An agent can reason correctly about what to do and then call the right tool with malformed arguments. Treating agent evaluation as a single end-to-end accuracy check misses both failure surfaces.
Component-level evaluation : measuring reasoning quality and tool-calling accuracy in isolation
End-to-end evaluation : measuring task completion and execution efficiency
A task completion rate of 80% tells you nothing about whether the 20% failure comes from bad planning, wrong tool selection, incorrect arguments, or tool infrastructure failures. Step-level traces — logs capturing each tool call, its arguments, its result, and the subsequent model decision — are what make that diagnosis possible. Without traces , debugging a production failure is guesswork.
Step 2: Defining What Agent Evaluation Success Looks Like
Evaluation is only as good as its success criteria. A well-formed eval task is one where two domain experts, working independently, would reach the same pass/fail verdict.
Start with unambiguous task specifications paired with reference solutions — known-correct outputs that pass all graders. They prove the task is solvable and verify that grading logic is correctly configured.
You need the following defined for evals before any grading runs :
The task : what inputs the agent receives, what it’s expected to do, and what the environment looks like going in
The success criteria : not just the final answer, but the intermediate outcomes that matter: Was the right tool called? Was the state correctly updated? Was the response grounded in the retrieved context?
The negative cases : one-sided evals create one-sided optimization. Balanced datasets — covering both when a behavior should occur and when it should not — prevent agents that over-trigger or under-trigger on a capability
A set of well-specified tasks drawn from real usage failures is a better starting point than waiting for the perfect dataset. Evals get harder to build the longer you wait.
Step 3: Grading the Agent Action Layer with Code-Based Checks
Deterministic graders — code that checks specific conditions without model-in-the-loop judgment — are the fastest, cheapest, and most reproducible option in any agent eval stack. For the action layer, they should always be the starting point:
Tool call verification: whether the agent called the right tool in the correct sequence
Argument validation: whether inputs have correct types, required parameters, and valid values
Outcome verification: whether the environment ends in the expected state
Transcript analysis: number of turns, tokens consumed, and latency
These are often fast, objective, and easy to debug, but brittle. A grader checking for “confirmation_code”: “CONF-789” will miss a correct response that formats the same data differently.
Step 4: Grading Agent Reasoning and Output Quality with Model-Based Judges
Some agent evaluation dimensions resist deterministic checking — output quality, tone, faithfulness to retrieved context, appropriate empathy. For these, a language model used as a judge or LLM-as-a-Judge is the right tool: flexible and capable of handling open-ended output, but introducing non-determinism and calibration drift that code-based graders don’t have.
The following practices keep model-based graders reliable:
Write structured rubrics . “Evaluate whether the response is helpful” produces noise. A rubric specifying that the response must address the user’s question, ground claims in retrieved context, and avoid out-of-scope suggestions produces a signal. Grade each dimension with a separate, isolated judgment.
Calibrate against human judgment regularly . LLM-as-judge accuracy should be checked against a sample graded by domain experts. Where divergence shows up, the rubric is almost always the problem. Give the grader an explicit “Cannot determine” option to avoid forced judgments on ambiguous cases.
Build in partial credit for multi-component tasks . A support agent that correctly identifies the problem and verifies the customer but fails to process the refund is meaningfully better than one that fails on step one. Binary pass/fail hides where the agent is actually breaking down.
Step 5: Matching Agent Evaluation Strategy to Agent Type
Grading strategies apply broadly, but agent type determines which graders carry the most weight and which failure modes to prioritize .
Coding agents write, test, and debug code. Software is largely deterministic: does the code run, do the tests pass, does the fix close the issue without breaking existing functionality? Benchmarks like SWE-bench Verified and Terminal-Bench follow this pass/fail approach, supplemented by rubric-based quality checks for security, readability, and edge case handling.
Conversational agents interact with users across support, sales, and coaching workflows. The quality of the interaction is part of what’s being evaluated — not only whether the ticket was resolved, but whether the tone was appropriate and the resolution clearly explained. This requires a second language model simulating the user; τ-bench models exactly this, with graders assessing both task completion and interaction quality across turns.
Research agents gather and synthesize information across sources. Groundedness checks verify claims are supported by retrieved sources, coverage checks define what a good answer must include, and source quality checks confirm the agent consulted authoritative material.
Matching Agent Evaluation Strategy to Agent Type
Step 6: Accounting for Non-Determinism in Agent Evaluation Results
Agent behavior varies between runs; the same task, same inputs, same agent can produce different tool selections, reasoning paths, and outcomes. Single-trial evaluation can therefore be misleading, since it hides variability that simple accuracy metrics fail to capture.
This is a direct consequence of non-determinism in agent systems . Stochastic model outputs, tool latency, partial failures, and adaptive decision-making all introduce variability across runs. As a result, evaluating an agent requires reasoning over distributions of outcomes rather than a single execution trace.
To account for this variability, metrics like pass@k and pass^k are commonly used:
pass@k: the probability that at least one of k independent trials succeeds, useful when multiple attempts are acceptable
pass^k: the probability that all k trials succeed, important when every interaction must be reliable
For example, an agent with a 75 percent single-trial success rate succeeds on all three attempts only about 42 percent of the time, showing how quickly reliability degrades across repeated runs.
pass@k and pass^k
Step 7: Separating Agent Capability Evals from Regression Suites
Capability evals are designed to answer a forward-looking question: what can this agent do that it couldn’t do before? Because of that, they should begin with relatively low pass rates and focus on tasks that are still challenging for the system. When a capability eval reaches very high scores — say 90 percent — it is often no longer measuring capability, but simply confirming reliability on already solved problems.
Regression evals serve a different purpose. They ask whether the agent can still perform everything it previously could . These tests should run close to 100 percent and act as a safeguard against performance regressions. Any meaningful drop in score is a signal that something has broken and should be investigated before release.
Over time, capability evals naturally become easier for the agent. As pass rates rise and performance stabilizes, those tasks can be promoted into the regression suite. However, once a suite fully saturates, it becomes less sensitive to real improvements — meaning meaningful progress may appear as noise rather than signal. For this reason, new and more challenging evals should be introduced before the existing suite saturates, not after.
Step 8: Extending Agent Evaluation into Production Monitoring
Development evals capture what you expect to fail; production reveals what actually does. Real users introduce inputs, edge cases, and contexts that rarely appear in synthetic test suites, making production monitoring a necessary extension of evaluation.
A complete evaluation system combines several complementary signals:
Together, these layers form a more complete view of agent performance in practice. Step-level traces — capturing reasoning, tool calls, arguments, results, and decisions at each point in the loop — are the infrastructure that makes all of this work. Tools like LangSmith , Arize Phoenix , Braintrust , and Langfuse provide tracing and eval frameworks; Harbor and DeepEval handle the harness layer.
Summary of Key Agent Evaluation Steps
Here’s a quick overview of the steps we’ve discussed:
As a next step, read Anthropic’s Demystifying evals for AI agents guide, especially the section Going from zero to one: a roadmap to great evals for agents .
Share
Post
Share
More On This Topic
Agent Evaluation: How to Test and Measu

[truncated]
