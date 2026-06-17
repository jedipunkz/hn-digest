---
source: "https://www.andrewjesson.com/blog/the-engineering-practices-claude-code-and-codex-use-to-improve-ai-agents/"
hn_url: "https://news.ycombinator.com/item?id=48569455"
title: "The engineering practices Claude Code and Codex use to improve AI agents"
article_title: "The engineering practices Claude Code and Codex use to improve AI agents · Andrew Jesson"
author: "anndvision"
captured_at: "2026-06-17T13:23:30Z"
capture_tool: "hn-digest"
hn_id: 48569455
score: 1
comments: 0
posted_at: "2026-06-17T12:28:55Z"
tags:
  - hacker-news
  - translated
---

# The engineering practices Claude Code and Codex use to improve AI agents

- HN: [48569455](https://news.ycombinator.com/item?id=48569455)
- Source: [www.andrewjesson.com](https://www.andrewjesson.com/blog/the-engineering-practices-claude-code-and-codex-use-to-improve-ai-agents/)
- Score: 1
- Comments: 0
- Posted: 2026-06-17T12:28:55Z

## Translation

タイトル: Claude Code と Codex が AI エージェントを改善するために使用するエンジニアリング手法
記事のタイトル: AI エージェントを改善するためにクロード コードとコーデックスが使用するエンジニアリング手法 · Andrew Jesson
説明: Claude Code と Codex が AI エージェントを改善するために使用するエンジニアリング手法

記事本文:
アンドリュー・ジェッソン Claude Code と Codex が AI エージェントを改善するために使用するエンジニアリング手法
コーディング エージェントは、AI エージェントの改善を求められた場合、一般的なエンジニアリング手法を実行します。故障モードの分析、評価、迅速な最適化のための特殊なツールを組み込む予定ですか?
コーディング エージェントに、シミュレートされたエージェント アプリケーション、100 のベースライン トレース、および最適化するメトリクスを与えると、改善が行われます。
Claude Code と Codex の両方がこれを行います。
彼らがそれをしながら何をするのかを見ることに興味がありました。
現在の TensorZero 設定と「yc_bench_tutorial_v0::yc_bench_act」のベースライン トレースをチェックしているので、バリアントを編集する前に失敗パターンを特定できます。
1/41
私は、Claude Code と Codex に、コンテナー内にあるエージェント CLI のみを変更して、5 つのシミュレートされたエージェント アプリケーションを最適化するよう指示しました。
私は驚きましたが、驚くべきことではなかったのかもしれませんが、二人ともクラスタリングや失敗パターンの要約などの即時的な実践を行っていたことに気づきました。
また、モデルまたはプロンプトに対して提案された変更を調整およびデバッグするために、アドホック評価も実行しました。
これらの一般的なエンジニアリング手法を実行することにより、故障モードの分析、評価、または迅速な最適化のための特別なツールを使用することなく、改善を実現しました。
これらの観察により、エージェントの最適化がより自動化されるにつれて、そのようなツールの役割と形状を再考する機会が得られました。
これらは、私が「ハーネス アトリビューション」と呼ぶプロジェクトを始めた理由でもあります。この投稿はその最初の調査です。
次の各アプリケーションについて、最大 100 の異なるタスクに対して、初期プロンプトとモデル ( gpt-5.4-mini ) を使用してベースライン エージェントを実行しました。
結果として得られたトレースは、アプリケーション固有のフィードバックを使用してスコア付けされました。
最適化タスクは、ベースライン エージェント プロンプトを変更することでアプリケーションの改善を提案することでした。

および/または別の同様の価格帯のモデルを選択する。
次に、オプティマイザー エージェント (claude-sonnet-4-6 の Claude Code または gpt-5.4 の Codex) がコンテナにドロップされ、これらのトレース、フィードバック、ベースライン エージェント構成のコピー、タスクを記述するマークダウン スキル ファイルにアクセスできるようになりました。
トレースとフィードバックを分析し、1 つ以上の新しいモデル プロンプト バリアントをエージェント構成に書き込み、終了しました。
提案された改善点の検証により、両方のコーディング エージェントが、すべてのアプリケーションでベースラインと一致またはそれを上回る新しいバリアントを出荷したことが明らかになりました。 NDA とサイエンスに関しては標準誤差 1 つ以内です。
アプリケーションごとに実施されたテストのスコア。
エラーバーは、最適化されたバリアントの 5 つのシードにわたる平均値 ± SE です。ベースラインは予算上の理由から単一シードで実行されたため、シードの分散は測定されていません。
エージェントはどのようなエンジニアリング手法を使用していますか?
どちらのコーディング エージェントも同じスキル ファイルを使用します。
これには、アプリケーション名、メトリック、利用可能なモデル、データ レイアウト、効率化のためのいくつかのレシピ、および調査 → バリアントの追加 → テスト → 反復という 4 つの箇条書きの方法論が含まれています。
{config_dir} 、 {function_name} 、 {baseline_metrics} 、 {model_list} などのプレースホルダーは、実行ごとにハーネスによって置き換えられます。
# TensorZero 関数オプティマイザー
TensorZero 関数を最適化して、パフォーマンス メトリックを向上させています。
## 環境
- T0 構成ファイル: {config_dir}/ (これらと以下のベースライン データのみが関連します。他の場所は調べないでください)
- ゲートウェイ URL: {gateway_url}
- 事前ダンプされたベースライン データ: {baseline_data_dir}/ (読み取り専用。DB への直接アクセスは利用できません)
- 設定編集後に再起動します: `curl -sf -X POST http://eval:5111/restart-gateway`
- 隔離されたコンテナ。 Python や `pip` は使用できません。 `node` と `curl` は `$PATH` 上にあります。

jq がインストールされていません。 JSONL 解析には `node -e "..."` を使用します (`readline` + `JSON.parse` + stdout へのプロジェクト)。行ごとにフィールドが必要な場合は、シェル パイプラインよりも優先します。
- どのバリアントにも「温度」を設定しないでください (一部のモデルはデフォルト以外の値を拒否します)。 「初期」バリアントをベースライン参照として保持します。
- 評価エピソードを自分で実行しないでください。終了後にハーネスが実行します。
## タスク
- 関数: `{関数名}`
- メトリック: `{metric_name}` 。方向については、`tensorzero.toml` のメトリクスの `optimize` フィールドを確認してください (ブール値メトリクスと浮動小数点メトリクスは最小化または最大化する可能性があります)。
- ベースライン パフォーマンス: {baseline_metrics}
## 利用可能なモデル
{モデルリスト}
## ベースラインデータ
- `{baseline_data_dir}/inferences.jsonl` — 推論ごとに 1 行 (タスクごとにモデルが言ったこと)。
- `{baseline_data_dir}/フィードバック.jsonl` — メトリック値ごとに 1 行。
- `{baseline_data_dir}/initial_config/` — 開始 T0 構成ツリーの読み取り専用コピー。
ファイルは 20 MB 以上になることがよくあります。それらを丸ごと「猫」にしてはいけません。それぞれの「head -3」から始めて行の形状を学習し（フィールド名とネストは環境によって異なります）、必要なフィールドを投影します。
### 投影パターン
最初に `grep` で絞り込み、次に `node -e` で投影します。
「」バッシュ
grep $TARGET_ID {baseline_data_dir}/inferences.jsonl \
|ノード -e "
require('readline').createInterface({input: process.stdin}).on('line', l => {
const r = JSON.parse(l);
console.log(r.id, r.variant_name, JSON.stringify(r.output).slice(0,200));
});"
「」
`猫の推論.jsonl | ...` は wh をロードします
[切り捨てられた]
故障モード分析を実行します
ここでの故障モード分析は、推論とフィードバックのデータセットから、「モデルがすべてをキャッチオールとして扱うため、さまざまなものを過剰に抽出する」ことになります。
このスキルは両方の前提条件をエージェントに任せます: 失敗した行を JSONL から投影し、次に ab

それらを名前付きパターンにまとめます。
投影ステップでは、データは 2 つのファイルに分割されます。feedback.jsonl はどの target_id が失敗したかを示し、inferences.jsonl はモデルがそれぞれについて実際に何を言ったかを示します。
元のスキルでは結合を散文で説明していましたが (失敗した target_id を取得し、対応する推論行を検索します)、その方法については説明していませんでした。
両方のエージェントは同じレシピに収束しました。フィードバックから失敗した target_id を grep し、次にそれぞれを grep して推論に戻し、最後の行を末尾に表示します。
私はそのレシピを、いくつかの関連するクロスレコードのワンライナー (エピソードごとの推論、存在するメトリクス、失敗したエピソードの最後の推論) とともにスキルに組み込みました。なぜなら、それらを再発見するにはセッション開始時に 3 ～ 6 ターンかかるからです。
失敗した行が投影されると、両方のエージェントが複数のトレースにわたって抽象化を実行できます。これには、スキルや関数のドキュメントに記載されていないバグも含まれることがよくあります。
以下のオプティマイザと環境を切り替えて、各エージェントがベースライン トレースから抽出したばかりの障害モードを列挙した瞬間に着陸します。
矢印キーを使用して、周囲の曲がり角を通過します。

## Original Extract

The engineering practices Claude Code and Codex use to improve AI agents

andrew jesson The engineering practices Claude Code and Codex use to improve AI agents
Coding agents perform common engineering practices when asked to improve AI agents. Will they subsume specialized tools for failure-mode analysis, evaluations, and prompt optimization?
Give a coding agent a simulated agent application, a hundred baseline traces, and a metric to optimize, and it will ship an improvement.
Both Claude Code and Codex do this.
I was interested in seeing what they do while doing it.
I’m checking the current TensorZero config and the baseline traces for `yc_bench_tutorial_v0::yc_bench_act` so I can identify failure patterns before editing variants.
1 / 41
I prompted Claude Code and Codex to optimize five simulated agent applications, varying only which agent CLI was in the container.
I was surprised, though maybe I should not have been, to find that they both used unprompted practices like clustering and summarizing failure patterns.
They also ran ad-hoc evaluations to refine and debug their proposed changes to the model or prompt.
By performing these common engineering practices, they shipped improvements without calling any specialized tooling for failure mode analysis, evaluations, or prompt optimization.
These observations gave me pause to reconsider the role and shape of such tooling as agent optimization becomes more automated.
They are also why I started a project I call harness attribution ; this post is its first probe.
For each of the following applications, I ran a baseline agent with an initial prompt and model ( gpt-5.4-mini ) on up to 100 different tasks.
The resulting traces were scored with application-specific feedback.
The optimization task was to propose improvements to the application by modifying the baseline agent prompt and/or choosing a different similar-price-point model.
The optimizer agent (Claude Code on claude-sonnet-4-6 or Codex on gpt-5.4 ) was then dropped into a container with access to those traces, feedback, a copy of the baseline agent config, and a markdown skill file describing the task.
It analyzed the traces and feedback, wrote one or more new model-prompt variants into the agent config, and exited.
Validation of the proposed improvements revealed that both coding agents shipped new variants that matched or beat the baseline on every application: decisively on NER, Business Management, and Software Engineering; within one standard error on NDA and Science.
Held-out test scores by application.
Error bars are mean ± SE across 5 seeds for the optimized variants; the baseline was run with a single seed for budget reasons, so its seed variance is unmeasured.
What engineering practices do the agents use?
Both coding agents use the same skill file.
It includes the application name, metric, available models, data layout, some recipes for efficiency, and a four-bullet methodology that says survey → add variants → test → iterate .
Placeholders like {config_dir} , {function_name} , {baseline_metrics} , and {model_list} are substituted per-run by the harness.
# TensorZero Function Optimizer
You are optimizing a TensorZero function to improve its performance metric.
## Environment
- T0 config files: {config_dir}/ (only these and the baseline data below are relevant — don't explore elsewhere)
- Gateway URL: {gateway_url}
- Pre-dumped baseline data: {baseline_data_dir}/ (read-only; direct DB access is not available)
- Restart after config edits: `curl -sf -X POST http://eval:5111/restart-gateway`
- Isolated container. No Python or `pip` ; `node` and `curl` are on `$PATH` ; `jq` is not installed. Use `node -e "..."` for JSONL parsing ( `readline` + `JSON.parse` + project to stdout) — prefer it over shell pipelines when you need fields per row.
- Don't set `temperature` on any variant (some models reject non-default values). Keep an `initial` variant as a baseline reference.
- Don't run evaluation episodes yourself — the harness does that after you exit.
## Task
- Function: `{function_name}`
- Metric: `{metric_name}` . Check the metric's `optimize` field in `tensorzero.toml` for direction (boolean and float metrics may minimize or maximize).
- Baseline performance: {baseline_metrics}
## Available Models
{model_list}
## Baseline data
- `{baseline_data_dir}/inferences.jsonl` — one row per inference (what the model said per task).
- `{baseline_data_dir}/feedback.jsonl` — one row per metric value.
- `{baseline_data_dir}/initial_config/` — read-only copy of the starting T0 config tree.
Files are often 20+ MB. Don't `cat` them whole. Start by `head -3` on each to learn the row shape (field names and nesting vary by env), then project out the fields you need.
### The projection pattern
`grep` first to narrow, then `node -e` to project:
```bash
grep $TARGET_ID {baseline_data_dir}/inferences.jsonl \
| node -e "
require('readline').createInterface({input: process.stdin}).on('line', l => {
const r = JSON.parse(l);
console.log(r.id, r.variant_name, JSON.stringify(r.output).slice(0,200));
});"
```
`cat inferences.jsonl | ...` loads the wh
[truncated]
They perform failure mode analysis
Failure mode analysis here is going from a dataset of inferences and feedback to “the model over-extracts miscellaneous because it treats it as a catch-all”.
The skill leaves both prerequisites up to the agent: projecting the failed rows out of JSONL, then abstracting them into a named pattern.
On the projection step, the data is split across two files: feedback.jsonl says which target_id s failed, inferences.jsonl says what the model actually said for each one.
The original skill described the join in prose ( pull failing target_ids, then look up the corresponding inference rows ) but did not say how.
Both agents converged on the same recipe: grep the failing target_id s out of feedback, then grep each one back into inferences and tail to the last row.
I folded that recipe back into the skill, alongside a few related cross-record one-liners (inferences-per-episode, which-metrics-are-present, last-inference-of-a-failing-episode), because re-discovering them cost three to six turns at the start of every session.
With the failed rows projected, both agents can do the abstraction across multiple traces, often including bugs not mentioned in the skill or the function’s documentation.
Toggle the optimizer and environment below to land on the moment each agent enumerates the failure modes it just abstracted from the baseline traces.
Use the arrow keys to step through the surrounding turns.
