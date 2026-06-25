---
source: "https://deep-reinforce.com/ornith_1_0.html"
hn_url: "https://news.ycombinator.com/item?id=48675882"
title: "Ornith-1.0: Self-Scaffolding LLMs for Agentic Coding"
article_title: "Ornith-1.0: Self-Scaffolding LLMs for Agentic Coding | DeepReinforce Blog | Jun. 2026"
author: "victormustar"
captured_at: "2026-06-25T16:45:20Z"
capture_tool: "hn-digest"
hn_id: 48675882
score: 3
comments: 0
posted_at: "2026-06-25T16:33:56Z"
tags:
  - hacker-news
  - translated
---

# Ornith-1.0: Self-Scaffolding LLMs for Agentic Coding

- HN: [48675882](https://news.ycombinator.com/item?id=48675882)
- Source: [deep-reinforce.com](https://deep-reinforce.com/ornith_1_0.html)
- Score: 3
- Comments: 0
- Posted: 2026-06-25T16:33:56Z

## Translation

タイトル: Ornith-1.0: エージェント コーディングのための自己足場 LLM
記事のタイトル: Ornith-1.0: エージェントティック コーディングのための自己足場 LLM | DeepReinforce ブログ | 2026年6月
説明: Ornith-1.0 は、エージェント コーディング タスクに特化した自己改善型のオープンソース モデル ファミリです。

記事本文:
Ornith-1.0: エージェント コーディング用の自己足場 LLM
2026年6月
ツイッター
ハグフェイス
アロハ！ 🌺
今日、私たちは Ornith-1.0 を紹介します。これは、エージェント コーディング タスクに特化した自己改善型のオープンソース モデル ファミリです。 Ornith-1.0 は全スペクトルをカバーします。
エッジ デバイスの導入に適したコンパクトな 9B Dense モデルから、最大のパフォーマンスを実現するために最適化された 397B MoE フロンティア スケール モデルまで、次のようなバリエーションがあります。
9B デンス、31B デンス、35B MoE、および 397B MoE。事前トレーニング済みの Gemma 4 および Qwen 3.5 上に構築されており、オープンソース モデルの中で最先端のパフォーマンスを実現します。
コーディングベンチマークで同等のサイズ。
Ornith-1.0 の背後にある重要な革新は、自己改善トレーニング フレームワークです。 Ornith-1.0 では、RL でのソリューション生成を推進するために人間が設計したハーネスに依存するのではなく、
ソリューションのロールアウトと、それらのロールアウトをガイドするタスク固有のハーネスの両方を生成する方法を学びます。足場とその結果として得られるソリューションを共同で最適化することで、モデル
より良い検索軌跡を発見し、より高品質のソリューションを生成できます。
Ornith-1.0 は、幅広いエージェント コーディング ベンチマークにわたって、同等のサイズのオープンソース モデルの中で最先端のパフォーマンスを実現します。 Ornith-1.0-397B ( 77.5
Terminal-Bench 2.1 では 82.4、SWE-Bench 検証済みでは 82.4）は、Claude Opus 4.7 のパフォーマンスと一致します（TB-2.1 では 70.3、SWE-Bench では 80.8）。
SWE-Bench Verified) を獲得しており、MiniMax M3 (TB-2.1 では 66.0、SWE-Bench では 80.5) を含む、同様のサイズの主要なオープンソース モデルよりも優れたパフォーマンスを発揮します。
検証済み）および DeepSeek-V4-Pro（TB-2.1 で 67.9、SWE-Bench 検証済みで 80.6）。 Ornith-1.0-9B はエッジ デバイスに簡単に導入できます。
Gemma 4-31B や Qwen 3.6 35B などのはるかに大きなモデルのパフォーマンスと同等またはそれを超えています。
フラッグシップ スケールでは、Ornith-1.0-397B は Terminal-Bench 2.1 で 77.5 を達成し、

SWE-Bench Verified では 82.4 で、Claude Opus 4.7 を上回っています。
どちらのベンチマークも、Minimax M3 や DeepSeek-V4-Pro など、同様の規模の主要なオープンソース モデルを上回っています。
Ornith-1.0-35B は、Qwen 3.5-35B、Qwen 3.6-35B、Gemma 31B などの同様のサイズのモデルよりも大幅に優れたパフォーマンスを発揮します。パラメータが 35B しかないにもかかわらず、それをさらに上回ります
Qwen 3.5-397B は Terminal-Bench 2.1 (64.4 対 53.5) で動作し、他のいくつかのコーディングおよびエージェント ベンチマークのパフォーマンスと同等です。
エッジ展開可能な Ornith-1.0-9B も非常に優れた結果をもたらし、ターミナルベンチ 2.1 で 43.1、SWE ベンチで 69.4 を達成しました。
検証済み。コンパクトな 9B パラメータ モデルであるにもかかわらず、Gemma 4-31B などのはるかに大きなモデルのパフォーマンスに匹敵するか、それを上回っており、強力なエージェント効果が実証されています。
コーディング機能は、リソース効率の高い展開でも実現できます。
LLM トレーニングの自己改善戦略
Ornith-1.0 の中核となるのは、課題を解決し、それらの解決策を導く足場を構築する方法を共同で学習する自己改善トレーニング フレームワークです。むしろ
Ornith-1.0 は、タスク カテゴリ全体で共有される人間が設計した固定ハーネスに依存し、足場をポリシーと共進化する学習可能なオブジェクトとして扱います。
各 RL ステップは 2 つの段階で進行します。タスクとそのタスクに以前に使用された足場に基づいて、モデルは最初に洗練された足場を提案します。その足場の上で条件付けされる
とタスクの説明を入力すると、ソリューションのロールアウトが生成されます。ロールアウトからの報酬は両方のステージに伝播されるため、モデルはより良いものを生み出すだけでなく最適化されます。
答えは、それを引き出すオーケストレーションを作成するだけです。
トレーニングを繰り返すことで、足場が継続的に変異し、より高い報酬軌道を誘導する足場に向かって選択されるフィードバック ループが生成されます。

s、許可する
タスクカテゴリごとの戦略が自動的に出現し、ハーネス設計を手作業で行うことなく、持続的な能力向上を推進します。
自己啓発における報酬ハッキングへの対処
モデルが独自の足場を作成できるようにすると、当然、報酬ハッキングの問題が発生します。自己生成されたスキャフォールドは、検証者を満足させる方法を学習できます。
タスク: 表示されているテスト ファイルを読み取り、予期されるアーティファクトをハードコーディングします。たとえば、チェック対象のファイルを操作する、リテラルの予期される出力を書き込む、または Oracle をコピーするなどです。
環境中に存在する溶液。
私たちはこれを 3 つの層で防御します。まず、外側の信頼境界を修正します。環境、ツール サーフェス、およびテスト分離は不変であり、モデルの外側にあります。
したがって、モデルは内部ポリシー スキャフォールド (メモリ、エラー処理、オーケストレーション ロジック) のみを進化させます。
第 2 に、決定論的モニターは、正確に指定できるレベルでその境界を強制し、保留されたパスの読み取り、検証スクリプトの変更、またはエラーの試みにフラグを立てます。
認可されたツール表面の外側でアクションを呼び出し、そのような軌道には利点の計算から除外して報酬をゼロに割り当てます。
第三に、意図レベルのゲームは許可されたツール サーフェス内で完全に発生する可能性があるため、凍結された LLM 裁判官は、主要な報酬ではなく検証者に加えて拒否権として機能します。
RL トレーニングでは、長期間のロールアウトに伴うオフライン ポリシーの問題に対処するために、Ornith-1.0 はパイプライン RL 戦略を採用しています。以前に生成されたオフポリシーの影響を制御するには
トークンには、古さの重み \(w(d_t)\) を適用します。これは、年齢 \(d_t\) に従ってトークンの重みを下げ、しきい値を超えるとトークンを完全に削除します。
トークンレベルの GRPO 損失は次のように重み付けされます。
Terminal-Bench 2.1 (Terminus-2): Terminal-Bench 2.1 u を評価します。

Harbor/Terminus-2 フレームワークを parser=json、temperture=1.0、top_p=1.0、および 128K コンテキストで指定します。
窓。各実行では 32 個の CPU コアと 48GB RAM で 4 時間のタイムアウトが使用され、結果は 5 回の実行で平均されます。 Qwen チャット テンプレートを調整して、相互の一貫性を確保します。
トレーニングと推論 ( https://huggingface.co/deepreinforce-ai/Ornith-1.0-397B/blob/main/chat_template.jinja )、vLLM のreasoning_content キーに合わせて Harbor を変更します。
ターミナルベンチ 2.1 (クロード コード): クロード コード 2.1.126 を使用し、parser=json、温度=1.0、top_p=1.0、max_new_tokens=131072 でターミナル ベンチ 2.1 を評価します。結果は
5回の実行で平均した。ここでも、Qwen チャット テンプレートを変更する必要があります。
SWE-Bench 検証済み、Pro および多言語: temp=1.0、top_p=0​​.95、256k コンテキスト ウィンドウで OpenHands ハーネスを使用。
SWE Atlas QnA、RF、TW: temp=1.0、top_p=0​​.95、128K コンテキスト ウィンドウでミニ SWE エージェント ハーネスを使用。結果は 5 回の実行で平均されます。
NL2Repo: 温度=1.0、top_p=1.0、400K コンテキスト、48K 出力、およびハッキング防止フィルター。
ClawEval: 実際のユーザーのタスク分散に対するエージェント コードのベンチマーク。 temp=0.6 および 256K コンテキスト。
© 2026 DeepReinforce.AI チーム。無断転載を禁じます。

## Original Extract

Introducing Ornith-1.0, a self-improving family of open-source models specially for agentic coding tasks.

Ornith-1.0: Self-Scaffolding LLMs for Agentic Coding
Jun. 2026
Twitter
Hugging Face
Aloha! 🌺
Today, we are introducing Ornith-1.0 , a self-improving family of open-source models specially for agentic coding tasks. Ornith-1.0 spans the full spectrum,
from compact 9B Dense models suitable for edge device deployment to 397B MoE frontier-scale models optimized for maximum performance, with variants including
9B Dense, 31B Dense, 35B MoE, and 397B MoE . Built on top of pretrained Gemma 4 and Qwen 3.5, it achieves state-of-the-art performance among open-source models
of comparable size on coding benchmarks.
The key innovation behind Ornith-1.0 is a self-improving training framework. Instead of relying on human-designed harnesses to drive solution generation in RL, Ornith-1.0
learns to generate both solution rollouts and the task-specific harnesses that guide those rollouts. By jointly optimizing the scaffold and the resulting solution, the model
can discover better search trajectories and generate higher-quality solutions.
Ornith-1.0 achieves state-of-the-art performance among open-source models of comparable size across a broad range of agentic coding benchmarks: Ornith-1.0-397B ( 77.5
on Terminal-Bench 2.1 and 82.4 on SWE-Bench Verified) matches the performance of Claude Opus 4.7 ( 70.3 on TB-2.1 and 80.8 on
SWE-Bench Verified) and outperforming leading open-source models of similar size, including MiniMax M3 ( 66.0 on TB-2.1 and 80.5 on SWE-Bench
Verified) and DeepSeek-V4-Pro ( 67.9 on TB-2.1 and 80.6 on SWE-Bench Verified). Ornith-1.0-9B, which can be easily deployed on edge devices,
matches or exceeds the performance of much larger models such as Gemma 4-31B and Qwen 3.6 35B.
At the flagship scale, Ornith-1.0-397B achieves 77.5 on Terminal-Bench 2.1 and 82.4 on SWE-Bench Verified, surpassing Claude Opus 4.7 on
both benchmarks and outperforming leading open-source models of similar size, including Minimax M3 and DeepSeek-V4-Pro.
Ornith-1.0-35B significantly outperforms similarly sized models, including Qwen 3.5-35B, Qwen 3.6-35B, and Gemma 31B. Despite having only 35B parameters, it even surpasses
Qwen 3.5-397B on Terminal-Bench 2.1 (64.4 vs. 53.5) while matching its performance across several other coding and agentic benchmarks.
The edge-deployable Ornith-1.0-9B also delivers remarkably strong results, achieving 43.1 on Terminal-Bench 2.1 and 69.4 on SWE-Bench
Verified. Despite being a compact 9B-parameter model, it matches or exceeds the performance of much larger models such as Gemma 4-31B, demonstrating that strong agentic
coding capabilities can be achieved even in resource-efficient deployments.
A Self-improving Strategy for LLM Training
At the core of Ornith-1.0 is a self-improving training framework that jointly learns to solve tasks and to construct the scaffolds that guide those solutions. Rather than
relying on a fixed, human-designed harness shared across a task category, Ornith-1.0 treats the scaffold as a learnable object that co-evolves with the policy.
Each RL step proceeds in two stages: conditioned on a task and the scaffold previously used for it, the model first proposes a refined scaffold; conditioned on that scaffold
and the task description, it then generates a solution rollout. Reward from the rollout is propagated to both stages, so the model is optimized not only to produce better
answers but to author the orchestration that elicits them.
Repeated over training, this yields a feedback loop in which scaffolds are continually mutated and selected toward those that induce higher-reward trajectories, allowing
per-task-category strategies to emerge automatically and driving sustained capability gains without hand-engineered harness design.
Addressing Reward Hacking in Self-improvement
Allowing the model to author its own scaffold naturally introduces the reward-hacking issue. A self-generated scaffold can learn to satisfy the verifier without performing the
task: reading the visible test files and hardcoding the expected artifacts, such as touching the checked-for file or writing the literal expected output, or copying an oracle
solution present in the environment.
We defend against this in three layers. First, we fix the outer trust boundary: the environment, the tool surface, and test isolation are immutable and outside the model's
reach, so the model evolves only the inner policy scaffold: its memory, error-handling, and orchestration logic.
Second, a deterministic monitor enforces that boundary at the level it can be specified exactly, flagging any attempt to read withheld paths, modify verification scripts, or
invoke actions outside the sanctioned tool surface, and assigning such trajectories zero reward with exclusion from the advantage computation.
Third, because intent-level gaming can occur entirely within the allowed tool surface, a frozen LLM judge acts as a veto on top of the verifier rather than the primary reward.
For RL training, to address the off-line policy problem for long rollouts, Ornith-1.0 adopts the pipeline-RL strategy. To control the effect of earlier generated off-policy
tokens, we apply a staleness weight \(w(d_t)\) that downweights tokens according to their age \(d_t\) and drops them entirely once a threshold is exceeded:
The token-level GRPO loss is weighted as follows:
Terminal-Bench 2.1 (Terminus-2): We evaluate Terminal-Bench 2.1 using the Harbor/Terminus-2 framework with parser=json, temperature=1.0, top_p=1.0, and a 128K context
window. Each run uses a 4-hour timeout with 32 CPU cores and 48GB RAM, and results are averaged over 5 runs. We adjust the Qwen chat template to ensure consistency between
training and inference ( https://huggingface.co/deepreinforce-ai/Ornith-1.0-397B/blob/main/chat_template.jinja ), and modify Harbor to align with vLLM's reasoning_content key.
Terminal-Bench 2.1 (Claude Code): We evaluate Terminal-Bench 2.1 using Claude Code 2.1.126 with parser=json, temperature=1.0, top_p=1.0, max_new_tokens=131072. Results are
averaged over 5 runs. Again, Qwen chat template needs to be modified.
SWE-Bench Verified, Pro and Multilingual: using OpenHands harness with temp=1.0, top_p=0.95, 256k context window.
SWE Atlas QnA, RF, TW: using mini SWE agent harness with temp=1.0, top_p=0.95, 128K context window. Results are averaged over 5 runs.
NL2Repo: with temperature=1.0, top_p=1.0, 400K context, 48K output and anti-hacking filters.
ClawEval: An agentic code benchmark over real-user task distributions; temp=0.6 and 256K context.
© 2026 DeepReinforce.AI Team. All rights reserved.
