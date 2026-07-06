---
source: "https://www.aai-labs.com/en/research/local-llm-benchmark-administrative-tasks"
hn_url: "https://news.ycombinator.com/item?id=48809994"
title: "DGX Spark Local LLM Benchmark: Administrative Tasks"
article_title: "Local LLM Benchmark: Administrative Tasks | AAI Labs"
author: "tasubotadas"
captured_at: "2026-07-06T21:22:50Z"
capture_tool: "hn-digest"
hn_id: 48809994
score: 2
comments: 1
posted_at: "2026-07-06T20:23:00Z"
tags:
  - hacker-news
  - translated
---

# DGX Spark Local LLM Benchmark: Administrative Tasks

- HN: [48809994](https://news.ycombinator.com/item?id=48809994)
- Source: [www.aai-labs.com](https://www.aai-labs.com/en/research/local-llm-benchmark-administrative-tasks)
- Score: 2
- Comments: 1
- Posted: 2026-07-06T20:23:00Z

## Translation

タイトル: DGX Spark ローカル LLM ベンチマーク: 管理タスク
記事のタイトル: ローカル LLM ベンチマーク: 管理タスク | AAIラボ
説明: 21 タスクの管理ワーク スイート全体でローカルで実行される 13 の言語モデルの DGX Spark 評価。これには、カレンダー作成、ドキュメント解析、財務および時間追跡の計算、電子メールのトリアージ、プロフェッショナル ライティングが含まれます。

記事本文:
ローカル LLM ベンチマーク: 管理タスク | AAI Labs 主なサービス 事例 調査会社 詳細 お問い合わせ 読み込み中… サービス AI for Energy
UAB Taikomasis dirbtinis intelektas © 2026
ローカル LLM ベンチマーク: 管理タスク
21 タスクの管理ワーク スイート全体でローカルで実行される 13 の言語モデルの DGX Spark 評価。これには、カレンダー作成、ドキュメント解析、財務および時間追跡の計算、電子メールのトリアージ、プロフェッショナル ライティングが含まれます。
このレポートでは、管理作業のベンチマークとして、ローカルで実行される 13 の言語モデルを評価します。このタスク スイートには、カレンダー管理、ドキュメントの解析、財務および時間追跡の計算、電子メールのトリアージ、プロフェッショナルなライティングが含まれます。すべてのモデルは、128 GB のユニファイド メモリを備えた単一の DGX Spark システム上で Ollama を通じて実行されました。
この評価では、構造化出力に対する決定論的な Python チェックと、オープンエンドの書き込みタスクに対するローカルの Llama 3.1 70B 判定を組み合わせます。この分割により、集計スコアだけでなく、さまざまなカテゴリの管理作業への適合性によってもモデルを比較できるようになります。
ドミトロ・クレパチェフスキー、ジギマンタス・ギルダウスカス、ロベルト・マッケヴィッチ、タダス・シュボニス
このベンチマークは、モデルが指示を読み、ツールや構造化された出力を使用し、チェック可能な作業を生成する必要がある、実際的なオフィス タスクを中心に設計されました。 21 のタスクには、日常的な管理業務と自由形式の起草タスクが含まれており、研究で構造化された実行と散文の品質を区別できるようになります。
各モデルは同じエージェント ハーネスを通じて実行され、同じ基準でスコア付けされました。 GPU 競合の影響を受ける実行は個別に繰り返されたため、報告された結果ではモデルの動作とリソース スケジューリングのアーティファクトが区別されます。
Gemma4:31b と Gemma4:12b は最強のオーバーアを達成しました

すべての結果が得られ、統計的にはベンチマークのトップでした。 12B モデルは、大幅に小さいメモリ フットプリントでほぼ同じ品質を達成できるため、多くの導入環境にとってより実用的なデフォルトです。
Qwen3.6 は構造化文書作業に最も強力なモデルでした。これは、PDF からカレンダーへのタスクに合格したスイート内の唯一のモデルであり、自動化された構造化タスクに関して非 Gemma モデルをリードしました。
モデルのサイズはパフォーマンスの信頼できる予測因子ではありませんでした。実行時条件が制御されると、いくつかの小さなモデルが大きなモデルよりも優れたパフォーマンスを示しました。これは、このクラスの管理ワークフローでは、パラメーター数だけよりもタスクの適合性と出力動作の方が重要であることを示唆しています。
管理の自動化は、流暢なテキストの生成だけに依存します。有用なモデルは、複数のステップにわたる指示に従い、必要に応じて構造化ファイルまたはツール呼び出しを生成し、現実的なハードウェア制約の下でも信頼性を維持する必要があります。このベンチマークにより、これらの違いが可視化され、実践者がモデルを選択するための実践的な基礎が得られます。

## Original Extract

A DGX Spark evaluation of 13 locally run language models across a 21-task administrative work suite, covering calendaring, document parsing, financial and time-tracking calculations, email triage, and professional writing.

Local LLM Benchmark: Administrative Tasks | AAI Labs Main Services Cases Research Company More Contact Loading… Services AI for Energy
UAB Taikomasis dirbtinis intelektas © 2026
Local LLM Benchmark: Administrative Tasks
A DGX Spark evaluation of 13 locally run language models across a 21-task administrative work suite, covering calendaring, document parsing, financial and time-tracking calculations, email triage, and professional writing.
This report evaluates 13 locally run language models on a benchmark of administrative work. The task suite covers calendar management, document parsing, financial and time-tracking calculations, email triage, and professional writing. All models were executed through Ollama on a single DGX Spark system with 128 GB of unified memory.
The evaluation combines deterministic Python checks for structured outputs with a local Llama 3.1 70B judge for open-ended writing tasks. This split makes it possible to compare models not only by aggregate score, but also by their suitability for different categories of administrative work.
Dmytro Klepachevskyi , Žygimantas Girdauskas , Robert Mackevič , Tadas Šubonis
The benchmark was designed around practical office tasks that require a model to read instructions, use tools or structured outputs, and produce work that can be checked. The 21 tasks include routine administrative operations as well as open-ended drafting tasks, allowing the study to distinguish between structured execution and prose quality.
Each model was run through the same agent harness and scored with the same criteria. Runs affected by GPU contention were repeated in isolation, so the reported results separate model behaviour from resource-scheduling artifacts.
Gemma4:31b and Gemma4:12b achieved the strongest overall results and were statistically tied at the top of the benchmark. The 12B model is the more practical default for many deployments because it reaches nearly the same quality with a substantially smaller memory footprint.
Qwen3.6 was the strongest model for structured document work. It was the only model in the suite to pass the PDF-to-calendar task and led the non-Gemma models on automated structured tasks.
Model size was not a reliable predictor of performance. Several smaller models outperformed larger models once runtime conditions were controlled, which suggests that task fit and output behaviour matter more than parameter count alone for this class of administrative workflows.
Administrative automation depends on more than fluent text generation. A useful model must follow instructions over multiple steps, produce structured files or tool calls when required, and remain reliable under realistic hardware constraints. This benchmark makes those differences visible and gives practitioners a practical basis for model selection.
