---
source: "https://verkyyi.github.io/hermesbench/"
hn_url: "https://news.ycombinator.com/item?id=48341436"
title: "Show HN: HermesBench – workflow reliability evals for personal AI agents"
article_title: "HermesBench"
author: "verkyyi26"
captured_at: "2026-05-31T00:24:05Z"
capture_tool: "hn-digest"
hn_id: 48341436
score: 1
comments: 0
posted_at: "2026-05-30T23:03:40Z"
tags:
  - hacker-news
  - translated
---

# Show HN: HermesBench – workflow reliability evals for personal AI agents

- HN: [48341436](https://news.ycombinator.com/item?id=48341436)
- Source: [verkyyi.github.io](https://verkyyi.github.io/hermesbench/)
- Score: 1
- Comments: 0
- Posted: 2026-05-30T23:03:40Z

## Translation

タイトル: 表示 HN: HermesBench – パーソナル AI エージェントのワークフロー信頼性評価
記事タイトル: エルメスベンチ
説明: HermesBench は、公開レシピ、編集されたトレース証拠、および再現性メタデータを使用して、完全な Herme パーソナル エージェント構成をベンチマークします。

記事本文:
エルメスベンチ
クイックスタート
レシピ
プロフィール
痕跡
フィードバック
貢献する
GitHub
Hermes エージェントのランタイム評価
モデルだけではなく、パーソナル エージェント全体をベンチマークします。
HermesBench は完全な Herme 構成を評価します: プロンプト、
モデル/プロバイダー、ツール、エージェントスキル、メモリ、ゲートウェイの動作、
委任、安全性、待ち時間、安定性。現在の一般人
ベースライン スコアは 27 のパーソナル エージェント レシピ全体で 78.2 でした。
編集されたトレースを検査できます。
まずは証拠を、目に見える限界を持って。
公開されたすべての結果はシナリオ定義にリンクされ、公開されます。
スコア軸、ドライバー終了の決定、決定論的チェック、
編集されたトレース タイムライン。このサイトでは、このことが意図的に明確にされています。
これは初期のベースラインの 1 つであり、基本モデルのリーダーボードではありません。
現在の証拠形状の 3 つのタブ。
ベースラインが 1 つ公開されているだけでは、リーダーボードは時期尚早です。サイト
レシピ、レシピなど、ユーザーがナビゲートする必要があるコンテンツから始まります。
プロファイルとトレース。
コーディング エージェントを通じて実行します。
パブリック ユーザー パスウェイは意図的に単純になっています。プロンプトを次の場所にコピーします。
Codex、Claude、または別のコーディング エージェント。エージェントは
HermesBench スキルと最初の 1 つのシナリオ レシピを駆動します。フルバンドル
実行には時間がかかり、コストもかかるため、実行はオプトインです。
Codex または Claude にコピーするよう求めるプロンプト
HermesBench スキルを使用して、現在の Herme 構成のデフォルトのシナリオ レシピを 1 つ実行します。
スキル: https://github.com/verkyyi/hermesbench/blob/main/agent-skills/hermesbench/SKILL.md
スキルの「現在のエルメス構成を実行」ワークフローに従います。 Python API のデフォルトの単一レシピ パスを使用し、アーティファクトを保存し、スコアと主な結果を要約します。私が明示的に要求しない限り、完全なバンドルを実行しないでください。
アルファフィードバック
次の最善のアクションは、具体的なフィードバックです。
HermesBench にはセットアップの摩擦やスコアに関する早期フィードバックが必要です
サプライズ、レシピ

現実性、プロフィールの証拠、編集の信頼性。
ベンチマーク形状が役立つ場合は、リポジトリにスターを付けます。次の場合に問題を開く
1 つのレシピ、トレース、またはスコア軸が間違っていると感じます。
ワークフロー レシピ、パーソナル エージェントを幅広くカバー。
HermesBench は 1 つの貴重なワークフロー レシピから始まり、その後オプトインします。
より自信が必要な場合は、より幅広いスイートを使用できます。同梱カタログ
日常のパーソナル エージェントの作業: コンテキスト、カレンダー、Web、
レポート、コミュニケーション、場所、旅行、財務、安全性、
パワーユーザーの統合。
優れたエージェントは正しいことを安全に完了します。
HermesBench は信頼性を第一に考えていますが、機能に盲目ではありません。良い
設定は有益な機能を果たさなければなりません。その内容について真実を伝えてください。
知っていて、危険な副作用を回避し、安定性を保ち、迅速に対応し、
明確にコミュニケーションします。偏ったスコアにはペナルティが課されます。
有能だが安全ではないエージェント、安全だが役に立たないエージェント、または正しいが正しいエージェント
役に立たないほど遅いのは実際には良くありません。
詳細な公式と実装メカニズムは方法論に組み込まれています
文書;ウェブサイトはスコアリングモデルをユーザーが読みやすい状態に保ち、
LLMエージェント。
良い結果を再利用可能なレシピに変えます。
HermesBench は簡単なベンチマークとして便利ですが、
うまくいったものを公開します。編集されたプロファイル/構成パッケージを共有する場合、
セットアップによりレシピが改善されるか、汎用レシピが送信されます。
重要なパーソナル エージェントの使用例が欠落しています。
HermesBench スキルを使用して、現在の Herme プロファイル/構成をパブリック プロファイル送信として準備します。
スキル: https://github.com/verkyyi/hermesbench/blob/main/agent-skills/hermesbench/SKILL.md
まず代表的なレシピを 1 つ実行し、編集されたプロファイルのスナップショットとスコア証拠をパッケージ化して、プル リクエストを開く前に何をレビューする必要があるかを教えてください。
レシピ提出プロンプト

## Original Extract

HermesBench benchmarks full Hermes personal-agent configurations with public recipes, redacted trace evidence, and reproducibility metadata.

HermesBench
Quick start
Recipes
Profiles
Traces
Feedback
Contribute
GitHub
Hermes Agent runtime evaluation
Benchmark the whole personal agent, not just the model.
HermesBench evaluates complete Hermes configurations: prompt,
model/provider, tools, AgentSkills, memory, gateway behavior,
delegation, safety, latency, and stability. The current public
baseline scores 78.2 across 27 personal-agent recipes with
redacted traces you can inspect.
Evidence first, with visible limits.
Every published result links back to scenario definitions, public
score axes, driver closure decisions, deterministic checks, and
redacted trace timelines. The site is deliberately clear that this
is one early baseline, not a base-model leaderboard.
Three tabs for the current evidence shape.
With one baseline published, a leaderboard is premature. The site
now starts from the content people need to navigate: recipes,
profiles, and traces.
Run it through a coding agent.
The public user pathway is intentionally simple: copy the prompt to
Codex, Claude, or another coding agent. The agent loads the
HermesBench skill and drives one scenario recipe first. Full bundle
runs are opt-in because they take longer and cost more.
Prompt to copy into Codex or Claude
Use the HermesBench skill and run one default scenario recipe for my current Hermes configuration.
Skill: https://github.com/verkyyi/hermesbench/blob/main/agent-skills/hermesbench/SKILL.md
Follow the skill's "Run Current Hermes Configuration" workflow. Use the Python API default single-recipe path, save artifacts, and summarize the score and main findings. Do not run the full bundle unless I explicitly ask.
Alpha feedback
The best next action is concrete feedback.
HermesBench needs early feedback on setup friction, scoring
surprises, recipe realism, profile evidence, and redaction trust.
Star the repo if the benchmark shape is useful; open an issue if
one recipe, trace, or score axis feels wrong.
Workflow recipes, broad personal-agent coverage.
HermesBench starts with one valuable workflow recipe, then lets you opt into
broader suites when you need more confidence. The bundled catalog
covers everyday personal-agent work: context, calendar, web,
reports, communication, location, travel, finance, safety, and
power-user integrations.
Good agents finish the right thing safely.
HermesBench is reliability-first, but not capability-blind. A good
configuration should do useful work, tell the truth about what it
knows, avoid unsafe side effects, stay stable, respond promptly, and
communicate clearly. Lopsided scores are penalized because a personal
agent that is capable but unsafe, safe but unhelpful, or correct but
unusably slow is not actually good.
Detailed formulas and implementation mechanics live in the methodology
document; the website keeps the scoring model readable for users and
LLM agents.
Turn good results into reusable recipes.
HermesBench is useful as a quick benchmark, but it is also a way to
publish what worked. Share a redacted profile/config package when a
setup improves a recipe, or submit a generic recipe when an
important personal-agent use case is missing.
Use the HermesBench skill to prepare my current Hermes profile/config as a public profile submission.
Skill: https://github.com/verkyyi/hermesbench/blob/main/agent-skills/hermesbench/SKILL.md
Run one representative recipe first, package the redacted profile snapshot and score evidence, and tell me what must be reviewed before opening a pull request.
Recipe submission prompt
