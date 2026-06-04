---
source: "https://resources.kusho.ai/ai-agent-benchmark-api-bug-detection"
hn_url: "https://news.ycombinator.com/item?id=48399429"
title: "Show HN: Black-box API bug detection across 7 AI systems"
article_title: "AI Agent Benchmark: API Bug Detection | KushoAI"
author: "riyajoshi"
captured_at: "2026-06-04T16:12:47Z"
capture_tool: "hn-digest"
hn_id: 48399429
score: 10
comments: 4
posted_at: "2026-06-04T14:42:40Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Black-box API bug detection across 7 AI systems

- HN: [48399429](https://news.ycombinator.com/item?id=48399429)
- Source: [resources.kusho.ai](https://resources.kusho.ai/ai-agent-benchmark-api-bug-detection)
- Score: 10
- Comments: 4
- Posted: 2026-06-04T14:42:40Z

## Translation

タイトル: Show HN: 7 つの AI システムにわたるブラックボックス API バグ検出
記事のタイトル: AI エージェントのベンチマーク: API のバグ検出 |くしょあい
説明: AI 生成テストがライブ API の機能バグを検出する方法のブラックボックス評価。

記事本文:
すべてのリソース
レポートのダウンロード
デモを予約する
ベンチマーク
AI エージェントのベンチマーク
AI エージェント ベンチマーク: API バグ検出
AI 生成のテストがライブ API の機能バグをどのように発見するかについてのブラックボックス評価。
KushoAI が提供するオープン ベンチマークである APIEval-20 v1.0 を使用して評価しました。評価者の裁量を減らすために、スコアリングは実行ベースです。生成されたテストは、ライブ参照 API に埋め込まれた機能バグを引き起こすか、引き起こさないかのどちらかです。報告されるすべてのスコアは、シナリオごとに 5 回の独立した実行の平均です。
AI コーディング ツールは API テストを迅速に生成できます。さらに難しい問題は、それらのテストでバグが見つかるかどうかです。
このレポートでは、7 つのアプリケーション ドメインにわたる 20 のライブ API シナリオにわたってその質問を評価し、3 つの複雑さの層にわたって 97 個の機能バグが埋め込まれています。各システムは、JSON スキーマと 1 つの有効なサンプル ペイロードのみを受け取り、ライブ リファレンス API での障害を明らかにする API テスト ケースを生成する必要があります。ソースコードはありません。スキーマ以外のドキュメントはありません。どこに障害が植え付けられているかについてのヒントはありません。
評価には、KushoAI が提供するブラックボックス ベンチマークである APIEval-20 v1.0 を使用します。 KushoAI も評価対象システムの 1 つであるため、このレポートには方法論、ワークフロー定義、繰り返し実行セットアップ、および堅牢性チェックが含まれているため、読者はパフォーマンスの違いがどこから来るのかを理解できます。
7 つのシステムが、汎用 LLM、コーディング エージェント、KushoAI の 3 つのグループにわたって比較されました。このレポートでは、エンジニアリング チームが実際に一般的に試みるワークフロー モード (ワンショット プロンプト、構造化されたテスト戦略プロンプト、プロンプト チェーン、ネイティブ コーディング エージェント ワークフロー、ネイティブ API テスト生成) も比較しています。
単純な構造上のバグは、もはや意味のある差別化要因ではありません。ほとんどのシステムは、欠落フィールド、null、空の配列、および間違った型のテストを生成する可能性があります

s.これらのテストは便利ですが、スキーマのみから発見するのが最も簡単な種類の障害でもあります。
迅速なエンジニアリングにより、パラメーターの適用範囲、フォーマット、フィールドレベルのネガティブテストが改善されます。これにより、スイートはより幅広く、より明示的になり、解析が容易になりますが、一般的なコーディング ツールがフィールド間のビジネス状態を一貫して推論できるようにはなりません。
複雑なバグに関してはギャップが生じます。この評価では、KushoAI は複雑に植え付けられたバグの 76% を検出しました。これに対し、最も強力なコーディング エージェント ワークフローでは 53%、最も強力な汎用 LLM では 34% でした。
KushoAI は、主要スコアとすべてのバグ複雑度層で 1 位にランクされています。最大のマージンは、クロスフィールドおよびビジネスロジックのバグ検出という、運用リスクに最も関連する指標に現れます。
1. もっともらしく見えるテストスイートでもバグを見逃す可能性があります。
いくつかのシステムは、読みやすい名前、有効なペイロード、および幅広いフィールドをカバーするスイートを生成しました。違いは、既知のプラント障害のあるライブ API に対してこれらのスイートを実行した後にのみ現れました。
2. 単純なスキーマレベルのテストが重要な要素になりました。
ほとんどのシステムは、欠落フィールド、NULL 値、空の配列、および間違った型に対するテストを生成できます。これらのテストは便利ですが、ツールが運用関連の障害を検出できるかどうかを評価するには十分ではありません。
3. プロンプトは深さよりも幅を広げるのに役立ちます。
構造化プロンプトにより、パラメーター カバレッジ、JSON 妥当性、およびフィールド レベルのネガティブ テストが向上しました。彼らは、複数の分野にわたるビジネス ロジック テストを一貫して作成していませんでした。
4. 複雑なバグにより、フィールドの変更が API テスト設計から切り離されます。
最も困難なバグでは、個別に有効なフィールドを無効な状態に結合する必要がありました (無効な払い戻し状態、ロール階層違反、競合する繰り返しルール、検証前に有効化された通知チャネルなど)。
5. テストの構成は何よりも重要です

テストボリューム。
コーディング エージェントのワークフローでは、多くのテストが生成されることがよくあります。このギャップは、それらのテストが意味のあるフィールドでの相互作用を調査したかどうかによって生じました。
6. CI/CD 導入には一貫性が重要です。
KushoAI は、評価されたすべてのシステムの中で実行ごとの分散が最も低いことが示されました。生成されたテストを自動パイプラインに統合するチームにとって、出力の安定性はピーク パフォーマンスと同じくらい重要です。
API のバグ検出に別の評価が必要な理由
ほとんどの API テスト生成の比較では、ツールがテストを生成できるかどうかが問われます。それはハードルが低すぎます。現在の LLM はどれも、API スキーマから妥当なテストのリストを生成できます。
テスト名は包括的に聞こえるかもしれませんし、ペイロードは構文的に有効であるかもしれませんが、それはエンジニアリング チームにスイートが実際にリスクを軽減するかどうかを伝えるものではありません。従来のコーディング ベンチマークは通常、コードの正確性、タスクの完了、生成されたテストが実行されるかどうかなどのプロパティを測定します。 API テストには別の中心的な目的があります。それは、ライブ サービスの意図されたコントラクトに違反する動作を見つけることです。
より有用な評価の質問はより狭いものです。リクエスト スキーマと 1 つの有効なサンプル ペイロードのみが与えられ、ソース コードもスキーマを超えるドキュメントも、植え付けられた障害に関するヒントも何もない場合、AI システムは、ライブ API に植え付けられた機能バグを引き起こすテストを生成できますか?
それがこのレポートで評価されるタスクです。エンドツーエンドの動作を評価します。エージェントはスキーマとサンプルを読み取り、テスト スイートを構築し、そのスイートはライブ参照実装に対して実行され、どのバグがトリガーされるかによってスコアが決定されます。
このブラックボックス制約は、一般的な実務者の現実を反映しています。チームは、完全なドキュメント、テスト データ、実装コンテキストを入手する前に、OpenAPI スキーマを受け取ったり、ペイロードのサンプルをリクエストしたりすることがよくあります。その設定では、

有用なテスト エージェントは、フィールド名、データ型、説明、入れ子構造、および実行される操作から考えられる制約を推測する必要があります。
このベンチマークには、電子商取引、支払い、認証、ユーザー管理、スケジュール、通知、検索/フィルタリングにわたる 20 のシナリオが含まれています。これらのシナリオ全体で、単純なものが 28 件、中程度のものが 35 件、複雑なものが 34 件の、97 件の機能的バグが埋め込まれています。
ベンチマークは、すべての運用条件を再現しようとするものではありません。これは、本番環境で重要な 1 つの機能、つまり AI システムが限られたリクエスト形状のコンテキストから高信号 API テストを生成できるかどうかを分離します。これにより、比較が制御され、再現可能になり、検査が容易になります。
すべてのシステムは、シナリオごとに同じ 2 つの入力 (JSON スキーマと 1 つの有効なサンプル ペイロード) を受け取りました。実装コード、応答スキーマ、ログ、変更履歴、運用例、または植え付けられたバグのヒントは提供されませんでした。
各システムはテスト ケースの JSON 配列を生成する必要がありました。各ケースには test_name と完全なリクエスト ペイロードが含まれていました。期待される結果は必要ありません。評価者は、テストをライブ API に対して実行することによって、植え付けられたバグをトリガーするかどうかを判断します。
スキーマとサンプル ペイロードは合わせて、テスターが持つ可能性のある最小限の有用なコンテキストを表します。スキーマは、どのようなフィールドが存在し、どのような制約が明示的であるかをエージェントに伝えます。サンプル ペイロードは、API が通常どのように使用されるかを示しています。このベンチマークでは、システムが実装漏れや障害モードを直接示す手書きのドキュメントに依存できないように、他のすべてを意図的に差し控えています。
これにより、タスクはアサーションの作成ではなくテストの生成に集中し続けます。リクエストのペイロードが実際に仕掛けられたバグに到達しない限り、システムは自信を持って期待される結果を記述しても報酬を受けません。
システムが彼女を評価した

頻繁に更新します。結果は、この調査中に使用された特定のモデル、製品モード、プロンプト、およびワークフローの時点での評価として解釈される必要があります。
チームが 1 つのプロンプトの後で停止することはほとんどないため、ワークフローの比較が含まれています。実際には、エンジニアはワンショット プロンプトを試し、次にプロンプ​​トをより明示的にし、ツール自身のギャップをレビューするように依頼し、プロセスに基づいてローカル スクリプトを構築します。
KushoAI 以外の各システムについて、メイン リーダーボードは、テストされたモード全体で観察された最強のワークフローを報告します。これにより、一般的な LLM とコーディング エージェントは、ワンショット出力に対してのみ KushoAI を比較するのではなく、構造化されたプロンプトと反復の利点を得ることができます。
最終スコア =
0.70 x バグ検出率
+ 0.20 x カバレッジ スコア
+ 0.10 x 効率スコア
バグ検出率 = bugs_triggered / total_planted_bugs
カバレッジスコア = param_coverage
効率スコア = min(1, 見つかったバグ数 / テスト数)
たとえ広範囲に見えても、バグが見つからないテストにはエンジニアリング上の価値が限られているため、バグ検出が最も重視されます。
各最上位スキーマ フィールドを少なくとも 1 回実行するカバレッジ報酬スイート。これは意図的に単純化されており、エッジケースの深さ、ビジネス ロジックの範囲、またはバグ発見の品質の代用として読むべきではありません。
大量の冗長ノイズの中に少数の有用なケースを埋め込んでいるスイートには、効率性が不利になります。
主要なシステム全体でのカバレッジはほぼ飽和しているため、リーダーボードをカバレッジ記事として読むべきではありません。このレポートでは、生成されたスイートがトップレベルのスキーマ フィールドを実行するかどうかをカバレッジで測定します。エッジケースの深さ、フィールドを越えた推論、ビジネス ロジックの範囲は測定されません。この分離は、バグ検出、複雑なバグの検出、および実行間の一貫性によって実現されます。 KushoAI はバグ検出率が最も高い stron

複雑なバグの発生率と、実行全体にわたる最小の標準偏差。
KushoAI は、20 のシナリオすべてでフル カバレッジ (1.00) を達成しました。これは、生成されたスイートがすべてのシナリオですべてのトップレベル スキーマ フィールドを実行したことを意味します。これは、選択的なフィールドのターゲティングではなく、ネイティブ ワークフローのスキーマ トラバーサル アプローチを反映しています。
文脈化する価値のある指標の 1 つは効率です。ここでは、生成されたテストに対する見つかったバグの比率として定義されます。 KushoAI の効率スコア (0.14) は、ネイティブ ワークフローが一般的な LLM よりもシナリオごとに多くのテストを生成するため、全体的な探索が増加しますが、テストごとのバグの比率が低下することを反映しています。最終スコアではバグ検出の重み付けが 0.70 であるため、このトレードオフは意図的なものです。つまり、より少ないテストでより多くのバグを検出するよりも、より多くのテストでより多くのバグを検出する方が望ましいのです。 CI ランタイムを最適化するチームは、最初の生成パスの後に重複排除やスイートのトリミングを適用できます。
対象範囲とバグ検出は異なります。モデルは多くのフィールドに触れても、失敗を見逃す可能性があります。たとえば、スイートでは通貨を空の文字列でテストし、金額をゼロでテストすることはできますが、各フィールドが個別に有効で、全体的な支払い状態が無効である組み合わせはテストしません。評価者は後者に報酬を与えます。それが植え付けられたバグを明らかにするものだからです。
コーディング エージェントは、ローカル ファイル、形式修正、反復生成をより適切に処理するため、生のチャット モデルよりも優れたパフォーマンスを発揮します。それらのスコアは、実際のワークフローの利点を反映しています。残りのギャップは、一般的なソフトウェア エンジニアリング エージェントと、API バグ検出専用に構築されたシステムとの間にあります。
エンジニアリング チームにとって、一貫性はピーク スコアと同じくらい重要です。 1 回の実行で強力なスイートを生成し、次の実行で弱いスイートを生成するツールでは、レビューのオーバーヘッドが発生します。分散が低いほど、手動での再試行が減り、パスの信頼性が高くなります。

CIへ。
プロンプトは役立ちますが、ギャップは埋まりません
エンジニアリング チームにとって実際的な問題は、プロンプトを改善することで、一般的な AI コーディング ツールが専用の API テスト エージェントと競合できるかどうかです。役に立ちます。ギャップは埋まりません。
「より適切なプロンプトを書くだけ」が AI によって生成された弱いテストに対するデフォルトの反応であるため、このセクションは重要です。プロンプトを改善すると出力が向上します。これらにより、モデルはより多くのフィールドを列挙し、より多くの境界値を含めて、よりクリーンな JSON を返すようになります。しかし、このベンチマークにおける主な弱点は、指示に従っていないことではありません。これは、API の形状から意味のある無効な状態を推測する機能です。
人間によるセットアップとレビューの見積もりは、観察された評価ワークフロー時間に基づいており、API の実行時間は含まれません。レビュー時間はチーム、エンドポイントの複雑さ、および承認基準によって異なるため、方向性があるものとして読み取る必要があります。
構造化されたプロンプトにより対象範囲が拡大し、より多くの必須フィールド テスト、無効な種類のテスト、基本的な境界テスト、電子メール、通貨コード、電話番号、日付、列挙型の形式テストが生成されます。プロンプト連鎖により、エージェントは最初にテスト戦略を推測し、次にテストを生成し、次に自身のギャップを確認できるため、さらなる改善が加えられます。

[切り捨てられた]

## Original Extract

A black-box evaluation of how AI-generated tests find functional bugs in live APIs.

All Resources
Download Report
Book a demo
Benchmark
AI Agent Benchmark
AI Agent Benchmark: API Bug Detection
A black-box evaluation of how AI-generated tests find functional bugs in live APIs.
Evaluated using APIEval-20 v1.0 , an open benchmark contributed by KushoAI. To reduce evaluator discretion, scoring is execution-based: a generated test either triggers a planted functional bug in the live reference API or it does not. All reported scores are the mean of five independent runs per scenario.
AI coding tools can generate API tests quickly. The harder question is whether those tests find bugs.
This report evaluates that question across 20 live API scenarios spanning seven application domains, with 97 planted functional bugs across three complexity tiers. Each system receives only a JSON schema and one valid sample payload, then must generate API test cases that expose failures in a live reference API. No source code. No documentation beyond the schema. No hints about where failures are planted.
The evaluation uses APIEval-20 v1.0, a black-box benchmark contributed by KushoAI. Because KushoAI is also one of the evaluated systems, this report includes the methodology, workflow definitions, repeated-run setup, and robustness checks so readers can understand where the performance difference comes from.
Seven systems were compared across three groups: general-purpose LLMs, coding agents, and KushoAI. The report also compares workflow modes engineering teams commonly try in practice: one-shot prompting, structured test-strategy prompting, prompt chaining, native coding-agent workflows, and native API test generation.
Simple structural bugs are no longer a meaningful differentiator. Most systems can generate missing-field, null, empty-array, and wrong-type tests. These tests are useful, but they are also the easiest class of failures to discover from the schema alone.
Prompt engineering improves parameter coverage, formatting, and field-level negative tests. It makes suites broader, more explicit, and easier to parse, but it does not consistently make general coding tools reason about cross-field business states.
The gap opens on complex bugs. KushoAI detects 76% of complex planted bugs in this evaluation, compared with 53% for the strongest coding-agent workflow and 34% for the strongest general-purpose LLM.
KushoAI ranks first on the primary score and across all bug complexity tiers. The largest margin appears on the metric most tied to production risk: cross-field and business-logic bug detection.
1. Plausible-looking test suites can still miss bugs.
Several systems generated suites with readable names, valid payloads, and broad field coverage. The difference appeared only after running those suites against live APIs with known planted failures.
2. Simple schema-level tests are now table stakes.
Most systems can generate tests for missing fields, null values, empty arrays, and wrong types. These tests are useful, but they are not enough to evaluate whether a tool can find production-relevant failures.
3. Prompting helps breadth more than depth.
Structured prompts improved parameter coverage, JSON validity, and field-level negative tests. They did not consistently produce cross-field business-logic tests.
4. Complex bugs separate field mutation from API test design.
The hardest bugs required combining individually valid fields into invalid states, such as invalid refund state, role hierarchy violations, conflicting recurrence rules, or notification channels enabled before verification.
5. Test composition matters more than test volume.
Coding-agent workflows often generated many tests. The gap came from whether those tests explored meaningful field interactions.
6. Consistency matters for CI/CD adoption.
KushoAI showed the lowest run-to-run variance among all evaluated systems. For teams integrating generated tests into automated pipelines, output stability matters as much as peak performance.
Why API Bug Detection Needs a Different Evaluation
Most API test generation comparisons ask whether a tool can produce tests. That is too low a bar. Any current LLM can generate a list of plausible tests from an API schema.
The test names may sound comprehensive, and the payloads may be syntactically valid, but that does not tell an engineering team whether the suite actually reduces risk. Traditional coding benchmarks usually measure properties like code correctness, task completion, or whether generated tests execute. API testing has a different core objective: finding behavior that violates the intended contract of a live service.
A more useful evaluation question is narrower: Given only the request schema and one valid sample payload, with no source code, no documentation beyond the schema, and no hints about planted failures, can an AI system generate tests that trigger planted functional bugs in a live API?
That is the task evaluated in this report. It evaluates end-to-end behavior: the agent reads the schema and sample, constructs a test suite, the suite is executed against live reference implementations, and scoring is determined by which planted bugs are triggered.
This black-box constraint reflects a common practitioner reality. Teams often receive an OpenAPI schema or request payload examples before they have complete documentation, test data, or implementation context. In that setting, a useful testing agent has to infer likely constraints from field names, data types, descriptions, nested structure, and the operation being performed.
The benchmark contains 20 scenarios across e-commerce, payments, authentication, user management, scheduling, notifications, and search/filtering. Across those scenarios, it contains 97 planted functional bugs: 28 simple, 35 moderate, and 34 complex.
The benchmark does not try to reproduce every production condition. It isolates one capability that matters in production: whether an AI system can generate high-signal API tests from limited request-shape context. That makes the comparison controlled, repeatable, and easier to inspect.
Every system received the same two inputs per scenario: a JSON schema and one valid sample payload. No implementation code, response schema, logs, changelog, production examples, or planted-bug hints were provided.
Each system had to produce a JSON array of test cases. Each case included a test_name and a complete request payload . No expected outcomes were required; the evaluator determines whether a test triggers a planted bug by running it against the live API.
The schema and sample payload together represent the minimum useful context a tester might have. The schema tells the agent what fields exist and what constraints are explicit. The sample payload shows how the API is normally used. The benchmark intentionally withholds everything else so that systems cannot rely on implementation leakage or hand-written documentation that points directly at the failure modes.
This keeps the task focused on test generation rather than assertion writing. A system is not rewarded for writing a confident expected outcome unless the request payload actually reaches a planted bug.
The systems evaluated here update frequently. Results should be read as a point-in-time evaluation of the specific models, product modes, prompts, and workflows used during this study.
The workflow comparison is included because teams rarely stop after a single prompt. In practice, engineers try a one-shot prompt, then make the prompt more explicit, then ask the tool to review its own gaps, then build local scripts around the process.
For each non-KushoAI system, the main leaderboard reports the strongest workflow observed across the tested modes. This gives general LLMs and coding agents the benefit of structured prompting and iteration rather than comparing KushoAI only against one-shot outputs.
Final Score =
0.70 x Bug Detection Rate
+ 0.20 x Coverage Score
+ 0.10 x Efficiency Score
Bug Detection Rate = bugs_triggered / total_planted_bugs
Coverage Score = param_coverage
Efficiency Score = min(1, bugs_found / number_of_tests)
Bug detection is weighted most heavily because tests that do not find bugs have limited engineering value, even if they look broad.
Coverage rewards suites that exercise each top-level schema field at least once. It is intentionally simple and should not be read as a proxy for edge-case depth, business-logic coverage, or bug-finding quality.
Efficiency penalizes suites that bury a few useful cases inside a large amount of redundant noise.
Coverage is near-saturated across the leading systems, so the leaderboard should not be read as a coverage story. In this report, Coverage measures whether generated suites exercise top-level schema fields. It does not measure edge-case depth, cross-field reasoning, or business-logic coverage. The separation comes from bug detection, complex-bug detection, and run-to-run consistency. KushoAI has the highest bug detection rate, the strongest complex-bug rate, and the lowest standard deviation across runs.
KushoAI achieved full Coverage (1.00) across all 20 scenarios, meaning its generated suites exercised every top-level schema field in every scenario. This reflects the native workflow's schema traversal approach rather than selective field targeting.
One metric worth contextualizing is Efficiency, defined here as the ratio of bugs found to tests generated. KushoAI's Efficiency score (0.14) reflects that its native workflow generates more tests per scenario than general LLMs, which increases overall exploration but lowers the bugs-per-test ratio. Because the final score weights bug detection at 0.70, this tradeoff is intentional: finding more bugs with more tests is preferable to finding fewer bugs with fewer tests. Teams optimizing for CI runtime can apply deduplication or suite trimming after the initial generation pass.
Coverage and bug detection diverge. A model can touch many fields and still miss the failure. For example, a suite may test currency with an empty string and amount with zero, but never test the combination where each field is individually valid and the overall payment state is invalid. The evaluator rewards the latter because that is what reveals the planted bug.
Coding agents outperform raw chat models because they handle local files, format correction, and iterative generation better. Their scores reflect real workflow advantages. The remaining gap is between general software engineering agents and a system built specifically for API bug detection.
For engineering teams, consistency matters as much as peak score. A tool that produces a strong suite in one run and a weak suite in the next creates review overhead. Lower variance means fewer manual retries and a more reliable path into CI.
Prompting Helps, But It Does Not Close the Gap
The practical question for engineering teams is whether better prompting can make general AI coding tools competitive with a purpose-built API testing agent. It helps. It does not close the gap.
This section matters because “just write a better prompt” is the default reaction to weak AI-generated tests. Better prompts do improve output. They make the model enumerate more fields, include more boundary values, and return cleaner JSON. But in this benchmark, the main weakness is not lack of instruction-following. It is the ability to infer meaningful invalid states from the shape of the API.
Human setup and review estimates are based on observed evaluation workflow time and exclude API execution time. They should be read as directional, since review time varies by team, endpoint complexity, and acceptance criteria.
Structured prompting increases coverage and produces more required-field tests, invalid-type tests, basic boundary tests, and format tests for emails, currency codes, phone numbers, dates, and enums. Prompt chaining adds another improvement because the agent can first infer a test strategy, then generate tests, then review its own gap

[truncated]
