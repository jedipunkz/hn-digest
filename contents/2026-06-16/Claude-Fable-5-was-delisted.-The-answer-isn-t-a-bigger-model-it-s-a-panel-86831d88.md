---
source: "https://www.orcarouter.ai/blog/model-fusion-in-production-inside-orcarouter-fusion-and-the-routing-dsl"
hn_url: "https://news.ycombinator.com/item?id=48560120"
title: "\" Claude Fable 5 was delisted. The answer isn't a bigger model – it's a panel: \""
article_title: "Model Fusion in Production: Inside OrcaRouter Fusion and the Routing DSL · OrcaRouter Blog"
author: "sangwen"
captured_at: "2026-06-16T19:18:37Z"
capture_tool: "hn-digest"
hn_id: 48560120
score: 2
comments: 0
posted_at: "2026-06-16T18:50:26Z"
tags:
  - hacker-news
  - translated
---

# " Claude Fable 5 was delisted. The answer isn't a bigger model – it's a panel: "

- HN: [48560120](https://news.ycombinator.com/item?id=48560120)
- Source: [www.orcarouter.ai](https://www.orcarouter.ai/blog/model-fusion-in-production-inside-orcarouter-fusion-and-the-routing-dsl)
- Score: 2
- Comments: 0
- Posted: 2026-06-16T18:50:26Z

## Translation

タイトル: 「クロード・フェイブル 5 は上場廃止になりました。答えは大きなモデルではなく、パネルです:」
記事のタイトル: 本番環境におけるモデルの融合: OrcaRouter Fusion とルーティング DSL の内部 · OrcaRouter ブログ
説明: 本番環境でのモデルの融合: OrcaRouter Fusion とルーティング DSL の内部 — OrcaRouter ブログ。

記事本文:
本番環境でのモデルの融合: OrcaRouter Fusion とルーティング DSL の内部 · OrcaRouter ブログ
Orca Router モデル AI セキュリティ統合 Playground Docs ブログ 価格 🌐 JP サインイン API キーを取得 本番環境でのモデル フュージョン: OrcaRouter Fusion とルーティング DSL の内部
3 つのフロンティア モデルを並行して実行し、1 つをアンサーバックします。 1 行で呼び出すか、独自に作成します。
TL;DR.クロード・ファブル5は上場廃止となった。答えは、より大きなモデルではなく、パネルです。複数のフロンティア モデルを並行して実行し、審査員に最も強力な答えを返してもらいます。 OrcaRouter は、これを 2 つの方法で出荷します。1 つは他のモデルと同様に呼び出す組み込みの orcarouter/fusion ルーター、もう 1 つは独自のルーティング DSL です。これは、コピー＆ペーストのレシピ、5 つのアービター (合成、混合エージェントの融合を含む)、および SLA を賭けずにロールアウトする方法の両方に関するフィールド ガイドです。
パート 1 — 1 行で呼び出す: 内蔵 Fusion ルーター
Fable 5 は中止され制限されたため、広く呼び出すことはできなくなりました。 Fusion は、引き続き呼び出すことができるモデル、つまりフロンティア モデルのパネルを並行して実行し、最も強力な答えを返すドロップインの OpenAI 互換ルーターからそのレベルを再構築します。厳選された 3 つの層がすべてのワークスペースに同梱されています。
Fusion の 3 つの層 (パネル構成 × コンテキスト ウィンドウ)
クロード オーパス 4.8 + GPT-5.5 + ジェミニ 3.1 プロ
最適な対象: Fable-5 レベルの最大知能
以下に最適: Fable-5 レベルのバランスのとれた推論
ジェミニ 3.5 フラッシュ + MiniMax M2.7 + GLM 5.1
最適な用途: Fable-5 レベルの高速 + 安価な推論
(コンテキスト ウィンドウ = 最小のパネル メンバー、ファンアウトのバインディング制約。)
これらはマーケティングのナメクジではありません。これらは、集中管理される事前にコンパイルされた DSL ルーターです。実際の orcarouter/fusion プログラムをそのまま示します。
バージョン: 1
ルール:
- ID: ハードパネル
いつ: task_class == "コード" ||タスク_c

ラス == "エージェント" ||コードキーワード密度 >= 0.3 ||ツールがある ||難易度 >= 0.3
使用:
平行:
- { モデル: "anthropic/claude-opus-4.8" }
- { モデル: "openai/gpt-5.5" }
- { モデル: "google/gemini-3.1-pro-preview" }
裁定者:
戦略: best_of_n
モデル: "anthropic/claude-opus-4.8"
テンプレート: best_answer_v1
max_latency_ms: 120000
デフォルト:
デリゲート: バランスのとれた
注目に値する 2 つの設計上の選択肢:
それは実際の作業でのみファンアウトされます。 when: ゲートは、コード、エージェント、ツール使用、コード密度、または高難易度 (難易度 >= 0.3) プロンプトのパネルを起動します。それ以外はすべて、ワークスペースのバランスの取れたデフォルトに従います。 「こんにちは」ではなく、パネル料金を必要なときに正確に支払います。
裁判官は、文字通り、本当の答えを出します。 best_of_n は、LLM ジャッジ (ここでは、best_answer_v1 テンプレートを使用した Opus 4.8) を実行し、単一の最も強力な候補を選択し、それをそのまま提供します (希薄化マージは行われません)。出力は常に実際のモデルの答えです。
パート 2 — 選択 vs. ヒューズ: best_of_n と合成アービター
Fusion ルーターは を選択します。ただし、OrcaRouter にはヒューズ戦略 (合成) も同梱されており、ルーティング エンジン (service/dispatch_Parallel/synthesize.go) に追加されるエージェントの混合パターンが追加されます。違いはゲーム全体にあります。
best_of_n (SELECT) 合成 (FUSE)
┌─ 作品 4.8 ─┐ ┌─ 作品 4.8 ─┐
通ってみよう
└─ ジェミニ ─┘ └─► サーブ レッグ k を逐語的に └─ ジェミニ ─┘ 融合された新しい答えが 1 つ
出力 = 実際のモデルの答え 出力 = どのレッグよりも優れた新しい答え 真の融合のレシピ:
使用:
平行:
- { モデル: "anthropic/claude-opus-4.8" }
- { モデル: "openai/gpt-5.5" }
- { モデル: "google/gemini-3.1-pro-preview" }
裁定者:
戦略: 合成する
モデル: "anthropic/claude-opus-4.8" # アグリゲーター: ヒューズ候補

1 つの新しい答えに
テンプレート: synthesize_v1
- 請求は N+1 です。すべての区間の請求に加え、追加の通話としてアグリゲーターも請求されます。
- OpenAI チャット形式は V1 のみ — アグリゲーターは OpenAI チャット完了を発行します。 Claude/Gemini ネイティブ クライアントは、serve-first- success に低下します (レッグは引き続き請求されます)。
アグリゲーターはルーターの承認された候補セットに含まれている必要があり、そうでない場合は機能が低下します。
いつどれを使用するか: best_of_n 1 つのモデルの答えが完全に正しい可能性が高い場合 (コード、事実に基づく Q&A) — 明確で実際の答えが必要な場合。答えが相補的（調査、分析、長文）であり、長所を統合することが単一のテイクよりも優れている場合に合成します。
パート 3 — 独自のルーティング DSL プレイブックを構築する
厳選されたパネルが欲しくないですか？ルーティング DSL エディターの「Claude Fable 5 レベル」テンプレートから開始し (すべてのワークスペースに同梱され、Fusion ルーターをミラーリングします)、次に特化します。 6 つのコピー＆ペースト パターン:
1 — 実際に実行されるコードを出荷 → ファンアウトし、テストで勝者を選択させます。
- ID: ハードコード
場合: task_class == "コード" && 難易度 > 0.6
使用:
平行:
- { モデル: "anthropic/claude-opus-4.8"、 Thinking_budget_tokens: 16000 }
- { モデル: "openai/gpt-5.5"、reasoning_effort: 高 }
- { モデル: "google/gemini-3.1-pro-preview" }
アービター: { 戦略: testing_pass } testing_pass は実行ベースです。ハーネスを通過する候補にサービスを提供します。ジャッジ LLM は必要ありません。
2 — 簡単なプロンプト → 難易度ゲート (Fusion パターン、モデル) のために過剰な支払いをやめます。
- ID: 簡単
条件: 難易度 < 0.3
使用: { デリゲート: 最安 }
- ID: ハード
条件: 難易度 >= 0.3
使用:
平行:
- { モデル: "anthropic/claude-opus-4.8" }
- { モデル: "openai/gpt-5.5" }
アービター: { 戦略: best_of_n、モデル: "anthropic/claude-opus-4.8"、テンプレート: best_answer_v1 } 3 — レール上でエージェントを長時間実行し続ける → 不安定な場合にのみエスカレーションします。
-

ID: エージェント
いつ: task_class == "エージェント" && Agent_state.consecutive_errors == 0
使用: { モデル: "anthropic/claude-sonnet-4.6"、affinity_ttl: "5m" }
on_low_confidence:
シグナル: [next_turn_test_failed、self_doubt]
use: { model: "anthropic/claude-opus-4.8", Thinking_budget_tokens: 24000 } 4 — 不安定な出力を決定論的にする → 投票、分割でエスカレーション:
- ID: 抽出
いつ: task_class == "ラグ"
使用:
平行:
- { モデル: "anthropic/claude-opus-4.8" }
- { モデル: "openai/gpt-5.5" }
- { モデル: "google/gemini-3.1-pro-preview" }
裁定者: { 戦略: 多数派 }
on_disagreement: { model: "anthropic/claude-opus-4.8", Thinking_budget_tokens: 32000 } 5 — テール レイテンシとプロバイダー ブリップ → レースを打ち破り、ファーストレスポンダにサービスを提供します。
- ID: 種族
場合: request.stream == true && 難易度 < 0.5
使用:
平行:
- { モデル: "google/gemini-3.5-flash" }
- { モデル: "minimax/minimax-m2.7" }
- { モデル: "z-ai/glm-5.1" }
arbiter: { Strategy: first } 6 — SLA を賭けずにロールアウト → シャドウ (ライブ トラフィックと並行して評価し、選択される内容 + コスト デルタを記録し、ライブ ピックを提供) → カナリア % (dsl_canary_pct 5 → 25 → 100、リクエストごとに暗号ランダム)。測定された相違に基づいて移行し、即座にロールバックします。
チートシート: 5 人の裁定者
難易度ゲートのファンアウトにより請求額はフラットに保たれます (例示; コスト = 実際のトークン価格の計算) — 混合コスト = easy_share ×安価 +hard_share ×パネル:
70% の簡単なワークロードでは、全パネルの請求額の 3 分の 1 でパネル全体を実行できます。

## Original Extract

Model Fusion in Production: Inside OrcaRouter Fusion and the Routing DSL — OrcaRouter Blog.

Model Fusion in Production: Inside OrcaRouter Fusion and the Routing DSL · OrcaRouter Blog
Orca Router Models AI Security Integrations Playground Docs Blog Pricing 🌐 EN Sign in Get API key Model Fusion in Production: Inside OrcaRouter Fusion and the Routing DSL
Three frontier models in parallel, one answer back. Call it in one line — or compose your own.
TL;DR. Claude Fable 5 was delisted. The answer isn't a bigger model — it's a panel : run several frontier models in parallel and let a judge return the strongest answer. OrcaRouter ships this two ways: built-in orcarouter/fusion routers you call like any model, and a Routing DSL to compose your own. This is the field guide to both — with copy-paste recipes, the five arbiters (including synthesize, the Mixture-of-Agents fuse), and how to roll it out without betting your SLA.
Part 1 — Call it in one line: the built-in Fusion routers
Fable 5 was discontinued and restricted, so it's no longer broadly callable. Fusion rebuilds that level from the models you can still call — a drop-in, OpenAI-compatible router that runs a panel of frontier models in parallel and returns the strongest answer. Three curated tiers ship in every workspace:
The three Fusion tiers (panel composition × context window)
Claude Opus 4.8 + GPT-5.5 + Gemini 3.1 Pro
Best for: Fable-5 level Max intelligence
Best for: Fable-5 level balanced inference
Gemini 3.5 Flash + MiniMax M2.7 + GLM 5.1
Best for: Fable-5 level fast + cheap inference
(Context window = the smallest panel member's — the binding constraint on a fan-out.)
These aren't marketing slugs; they're pre-compiled DSL routers managed centrally. Here's the actual orcarouter/fusion program, verbatim:
version: 1
rules:
- id: hard_panel
when: task_class == "code" || task_class == "agent" || code_keyword_density >= 0.3 || has_tools || difficulty >= 0.3
use:
parallel:
- { model: "anthropic/claude-opus-4.8" }
- { model: "openai/gpt-5.5" }
- { model: "google/gemini-3.1-pro-preview" }
arbiter:
strategy: best_of_n
model: "anthropic/claude-opus-4.8"
template: best_answer_v1
max_latency_ms: 120000
default:
delegate: balanced
Two design choices worth calling out:
It only fans out on real work. The when: gate fires the panel for code, agent, tool-using, code-dense, or high-difficulty (difficulty >= 0.3) prompts; everything else falls through to the workspace's balanced default. You pay panel price exactly where it helps, not on "hi."
The judge serves a real answer, verbatim. best_of_n runs an LLM judge (here, Opus 4.8 with the best_answer_v1 template) that picks the single strongest candidate and serves it as-is — never a diluted merge. The output is always a real model's answer.
Part 2 — Select vs. Fuse: best_of_n and the synthesize arbiter
The Fusion routers select . But OrcaRouter also ships a fuse strategy — synthesize, the Mixture-of-Agents pattern added in the routing engine ( service/dispatch_parallel/synthesize.go ). The difference is the whole game:
best_of_n (SELECT) synthesize (FUSE)
┌─ Opus 4.8 ─┐ ┌─ Opus 4.8 ─┐
├─ GPT-5.5 ─┼─► judge picks leg k ├─ GPT-5.5 ─┼─► aggregator LLM writes
└─ Gemini ─┘ └─► serve leg k verbatim └─ Gemini ─┘ ONE new fused answer
output = a real model's answer output = a new answer better than any leg Recipe for true fusion:
use:
parallel:
- { model: "anthropic/claude-opus-4.8" }
- { model: "openai/gpt-5.5" }
- { model: "google/gemini-3.1-pro-preview" }
arbiter:
strategy: synthesize
model: "anthropic/claude-opus-4.8" # aggregator: fuses candidates into one new answer
template: synthesize_v1
- Billing is N+1 — every leg bills, plus the aggregator as an extra call.
- OpenAI chat-format only in V1 — the aggregator emits an OpenAI chat completion; Claude/Gemini native clients degrade to serve-first-successful (legs still billed).
The aggregator must be in the router's authorized candidate set, or it degrades.
When to use which: best_of_n when one model's answer is likely fully right (code, factual Q&A) — you want a clean, real answer. synthesize when answers are complementary (research, analysis, long-form) and merging strengths beats any single take.
Part 3 — Build your own: the Routing DSL playbook
Don't want the curated panel? Start from the "Claude Fable 5 Level" templates in the Routing DSL editor (they ship in every workspace and mirror the Fusion routers), then specialize. Six copy-paste patterns:
1 — Ship code that actually runs → fan out, let the tests pick the winner:
- id: hard_code
when: task_class == "code" && difficulty > 0.6
use:
parallel:
- { model: "anthropic/claude-opus-4.8", thinking_budget_tokens: 16000 }
- { model: "openai/gpt-5.5", reasoning_effort: high }
- { model: "google/gemini-3.1-pro-preview" }
arbiter: { strategy: tests_pass } tests_pass is execution-grounded — it serves the candidate that passes your harness, no judge LLM needed.
2 — Stop overpaying for easy prompts → difficulty gate (the Fusion pattern, your models):
- id: easy
when: difficulty < 0.3
use: { delegate: cheapest }
- id: hard
when: difficulty >= 0.3
use:
parallel:
- { model: "anthropic/claude-opus-4.8" }
- { model: "openai/gpt-5.5" }
arbiter: { strategy: best_of_n, model: "anthropic/claude-opus-4.8", template: best_answer_v1 } 3 — Keep long agent runs on the rails → escalate only when it wobbles:
- id: agent
when: task_class == "agent" && agent_state.consecutive_errors == 0
use: { model: "anthropic/claude-sonnet-4.6", affinity_ttl: "5m" }
on_low_confidence:
signals: [next_turn_test_failed, self_doubt]
use: { model: "anthropic/claude-opus-4.8", thinking_budget_tokens: 24000 } 4 — Make flaky outputs deterministic → vote, escalate on a split:
- id: extract
when: task_class == "rag"
use:
parallel:
- { model: "anthropic/claude-opus-4.8" }
- { model: "openai/gpt-5.5" }
- { model: "google/gemini-3.1-pro-preview" }
arbiter: { strategy: majority }
on_disagreement: { model: "anthropic/claude-opus-4.8", thinking_budget_tokens: 32000 } 5 — Beat tail latency & provider blips → race, serve the first responder:
- id: race
when: request.stream == true && difficulty < 0.5
use:
parallel:
- { model: "google/gemini-3.5-flash" }
- { model: "minimax/minimax-m2.7" }
- { model: "z-ai/glm-5.1" }
arbiter: { strategy: first } 6 — Roll out without betting the SLA → shadow (evaluate alongside live traffic, log what it would pick + cost delta, serve the live pick) → canary % (dsl_canary_pct 5 → 25 → 100, crypto-random per request). Migrate on measured divergence, roll back instantly.
The cheat-sheet: five arbiters
Difficulty-gated fan-out keeps the bill flat (Illustrative; cost = real token-price math) — blended cost = easy_share × cheap + hard_share × panel:
A 70%-easy workload runs the full panel for a third of the all-panel bill .
