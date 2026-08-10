---
source: "https://langfuse.com/academy/datasets/designing-great-datasets"
hn_url: "https://news.ycombinator.com/item?id=49247965"
title: "Scoping and curating AI eval datasets"
article_title: "Designing datasets - Langfuse"
author: "annabellschaf"
captured_at: "2026-08-10T19:49:29Z"
capture_tool: "hn-digest"
hn_id: 49247965
score: 5
comments: 0
posted_at: "2026-08-10T18:48:13Z"
tags:
  - hacker-news
  - translated
---

# Scoping and curating AI eval datasets

- HN: [49247965](https://news.ycombinator.com/item?id=49247965)
- Source: [langfuse.com](https://langfuse.com/academy/datasets/designing-great-datasets)
- Score: 5
- Comments: 0
- Posted: 2026-08-10T18:48:13Z

## Translation

タイトル: AI 評価データセットの範囲設定とキュレーション
記事のタイトル: データセットの設計 - Langfuse
説明: 明確な範囲、現実的な入力、評価可能な出力、有用なメタデータ、解釈可能な実験結果を備えた Langfuse データセットを設計するための実践的なガイド。

記事本文:
データセットの設計 - Langfuse Langfuse v4: 最大 165 倍高速 · 続きを読む Langfuse v4 はこちら: リアルタイム、最大 165 倍高速 · 続きを読む ドキュメント 🐐 ヨーロッパとサンフランシスコでの採用 GOATS を探しています!アプリの起動 L デモの取得 G ドキュメントの統合 セルフ ホスティング ガイド アカデミー ワークショップ AI エンジニアリング ライブラリ アカデミー データセットの設計 Langfuse アカデミー AI エンジニアリング ループ ステップ
データセット データセットの設計 実験 評価 サポート資料
データセットの設計 アカデミー データセット データセットの設計 ページをコピー AI アプリケーション用のデータセットの設計
データセットは、アプリケーションの範囲を表す反復可能な例のセットであり、システムの測定と改善に使用します。長期間にわたって同じ入力に対してアプリケーションを実行することで、メトリクスを使用して品質を追跡し、変更を比較し、本番ユーザーに影響を与える前に回帰を捕捉することができます。
再現可能な方法で評価する価値があるものを決めていませんか?チェックアウト
エラー分析 、チェックする価値のあるアプリケーションの障害モードを見つけるための構造化された方法
データセットの範囲を把握するために、Langfuse デモ プロジェクトとアカデミーのサンプルのデータセットを使用します。
データセット アカデミーのセクションでは、入力、期待される出力、メタデータといった基本的な構成要素について説明します。このガイドでは、データセットとデータセット アイテムの作成前および作成中に行われる設計作業に焦点を当てます。
データセットの設計は反復的です。適切な開始点は、最小限完全なデータセットです。つまり、アプリケーションで実行でき、最も重要な入力スライスをカバーし、評価またはレビュー ルーブリックを備えた約 15 ～ 30 行のデータセットです。そのバージョンを早期に実行し、スキーマとエバリュエーターを修正してから、それらの実行または運用環境の入力から発生したギャップに拡張します。
アプリケーションには複数の評価データセットが含まれる可能性が高くなります。データセットは多くの場合、

システムの特定の部分、またはエージェントが実行する 1 つのサブステップに範囲が限定されます。
反復と拡張 01 目標範囲 / 境界を定義する 02 ソースのサンプル / パターンを検査する 03 配布スライス / ロールを選択する 04 eval スタイルの参照 / 無料を選択する 05 スキーマの入力 / 出力を設計する 06 データセットの行 / ギャップを構築する 07 実験の失敗 / 反復を実行する
1. データセットの目標から始める
行を書き込む前に、データセットがサポートする必要がある最小の有用な目標を定義します。これにより、データセットが興味深いサンプルの漠然としたバケツになるのを防ぎます。
サポート ルーティングを中断することなく、迅速な変更を出荷できますか?
ドキュメント チャットボットは、統合に関する一般的な質問に十分に答えていますか?
顧客が返金を求めたとき、エージェントは適切なツールを呼び出しますか?
データセットの初期の良い開始点は、最も一般的な例をエンドツーエンドで調べることです。これにより、拡張できる基礎と一般的な理解が得られます。
時間の経過とともに、チームはアプリケーション内の単一ステップ、敵対的ケース、レッド チーム化、または特定のサブユース ケース用のデータセットを追加します。会社名のスペルをチェックするためのデータセットも見たことがあります。
データセットに設定した目標によって境界とジョブが定義されます。エンドツーエンドのデータセットはアプリケーションが受け取るペイロードを使用し、ステップレベルのデータセットは実稼働環境でステップが見る構造化された状態を使用します。
2 つのジョブに異なる入力、評価者、またはリリース決定が必要な場合は、それらを分割します。安定した回帰データセットと敵対的入力データセットは両方とも役立ちますが、明確に分割せずにそれらを組み合わせると、集計スコアの解釈が難しくなります。
サンプルを選択したり、データセット項目を作成したりする前に、使用できるマテリアルの小さなサンプルを調べてください。目標は、何が存在するのか、そしてそれが期待していたものとどのように異なるのかを理解することです。
まずは 3 つのソース タイプから始めます。
生産トレース: sh

現実的な使用法、一般的なパス、観察された障害。スコア、ユーザーのフィードバック、チケット、苦情は、有用なトレースの発見シグナルです。これらは別個のソースタイプではありません。
既存の資産: 古いデータセット、FAQ、ポリシー、ドキュメント、サポート マクロ、CSV、JSON ファイル、ベンチマークにより、既知の範囲を迅速にブートストラップできます。
合成ケース: 専門家が作成したサンプルや AI が生成したサンプルは、ギャップを埋め、アプリケーションが直面すると予想される問題についてのアイデアを提供します。
ソースごとに、入力トピック、入力および出力の形状、および障害モードを確認します。
実稼働トレースには評価者が必要とする取得コンテキストが欠けていること、レガシー システムからのチケットはトレースよりも優れた障害ラベルを公開していること、またはある時点でサポートが「FAQ」を作成したため、エンドツーエンドの優れた例のセットがすでに存在していることがわかるかもしれません。
この入力を使用して、次のステップで入力、評価者、項目スキーマの分布を選択します。
ガイド: 合成データセットを開く 他のデータが利用できない場合は、開始点として合成データを使用します。
3. 入力分布を選択します
最初のデータセットでは、入力分布を単純にしておきます。実行後に何をすべきかを示すいくつかのスライスから始めます。
シナリオ タイプ: 主要なジョブ、インテント、ルート、またはタスク ファミリ。サポート ルーティング データセットの場合、これは請求、アカウント アクセス、技術的な問題、販売リクエストなどです。
困難またはリスク: 日常的な、曖昧な、困難な、敵対的な、またはビジネスクリティカルなケース。満足のいくパス以外のパスも含めますが、最初のデータセットをすべてエッジ ケースにしないでください。
データセットの役割: 行が存在する理由: 典型的なケース、既知の回帰、観察された失敗、または合成ギャップ埋め。
サポート ルーティング データセットの場合、最初のバージョンでは、単純な難易度別のシナリオ マトリックスが使用される場合があります。
入力分布は、データセット内のケースを意図的に組み合わせたものです。どのシナリオが適用されるか

耳、その難しさ、各行が含まれる理由。これは、実験結果をスライスごとに解釈するのに役立つカバレッジ プランです。
分布は生産頻度を正確に反映する必要はありません。回帰データセットは、障害、エッジケース、または価値の高いパスを意図的に過剰に表現する場合があります。重要なのは、その組み合わせが意図的であり、メタデータに表示されることです。
動作を変更する場合、または結果の解釈に役立つ場合にのみ、ディメンションを追加します。チャネル、言語、顧客セグメント、地域、コンテキストの可用性、および製品分野は有用なメタデータになる可能性がありますが、初日からそれらすべてがバランスのとれた制約になるべきではありません。
4. 評価の仕組みを決定する
期待される出力を作成する前に、評価スタイルを選択してください。これにより、データセット アイテムに何を含める必要があるか、最初の結果がどの程度役立つかが決まります。
各項目に既知のターゲット (正しいラベル、予期されるツールの呼び出し、必要な事実、構造化された出力、参照の回答、または予期される次のアクション) がある場合は、参照ベースの評価を使用します。これは、障害を検査しやすいため、回帰テストや CI ゲートに最適です。その代償として、参照を書くのに手間がかかり、動作ではなく文言を過剰に指定すると脆弱になる可能性があります。
期待される安定した出力はないが、すべての項目が同じルールまたはルーブリックに基づいて判断できる場合は、参照なしの評価を使用します。これは、有効な JSON、言語の一致、提供されたコンテキストの根拠、安全性、トーンなどのチェックに機能します。トレードオフとして、評価者またはルーブリックの重みが大きくなるため、あいまいなルーブリックはあいまいな結果を生み出します。
たとえば、ドキュメント チャットボット データセットでは、次のいずれかのアプローチを使用できます。
要件を把握する最も安価なエバリュエーターを選択します: 決定論的チェックのためのコード エバリュエーター、言語品質判断のための LLM-as-a-judge、および手動評価

良い出力と悪い出力がどのようなものかをまだ学んでいる間に確認してください。
この段階では、各データセット項目をどのようにレビューするか (自動評価、手動アノテーション、またはその両方) を決定するだけで済みます。エバリュエーター自体の設計に関する詳細な作業については、評価に関するアカデミーのページを参照してください。
各行を採点する評価者の名前やレビュー ルーブリックをまだ指定できない場合は、まずレビュー ワークフローを実行し、注釈キューを使用して例に手動でラベルを付けます。
3 つのデータセット項目フィールドは柔軟な JSON です。ここでは、実験実行者、評価者、レビュー担当者がすべて利用できる具体的なコントラクトとして定義します。
input : システム境界に渡すオブジェクト
ExpectedOutput : 評価者またはレビュー担当者が必要とする参照データのみ。意図的に参照なしで評価する場合は省略します
メタデータ : ソース、シナリオ タイプ、難易度、データセットの役割、レビュー ステータスなどの安定したスライスと来歴フィールド
項目スキーマを決定したら、それを適用して、チームが効果的に協力して適切な構造に項目を追加できるようにします。
サポート ルーティング データセットの場合、1 行は次のようになります。
{
" 入力 " : {
" message " : "請求書 4831 に対して 2 回請求されました。誰かがこれを修正できますか?" 、
" チャンネル " : "support_chat" ,
" customer_tier " : "ビジネス"
}、
" 期待される出力 " : {
" ルート " : "billing_support" 、
" required_actions " : [ "acknowledge_duplicate_charge" , "ask_for_invoice_id" ],
"must_not" : [ "promise_refund_without_review" ]
}、
" メタデータ " : {
" ソース " : "専門家" ,
" scenario_type " : "請求" ,
「難易度」：「中」、
" dataset_role " : "回帰" ,
" 失敗モード " : "間違ったルート"
}
}
入力はルーターが必要とするコンテキストのみを保持します。 ExpectedOutput は、チェックすべき動作を示します: 請求先へのルーティング、請求書 ID の要求、プロミダクションを行わない

レビューする前に返金してください。メタデータには、行のソース、シナリオ、難易度、および役割が記録されます。
行を一括収集する前に、スキーマを安定した状態に保ちます。会話履歴、取得されたコンテキスト、ツールの状態、ルーティング メタデータ、ユーザー属性などの動作を形成するフィールドを保持しますが、任意の行ごとのフィールドや構造化コンテキストの自然言語要約は避けます。
目標、分布、評価方法、スキーマが具体的になったら、選択したソース サンプルをデータセット アイテムに変換し始めます。前に検査したソースを使用し、定義されたスキーマと一致する形式で追加します。
最低限完成した最初のバージョンの場合、入力分布に沿ってコントラクト全体をテストするのに十分な行を選択します。
確実に動作する一般的なシナリオ
いくつかのあいまいなシナリオまたは高リスクのシナリオ
防止したい既知の障害または回帰
合成ギャップは、生産痕跡または既存の資産が流通をカバーしていない場合にのみ埋められます。
データセットが完了したと感じるまで待たないでください。さらに行を追加する前に、実験で最初の一貫したバージョンを実行します。
7. 最初の実験を実行し、意図的に拡張します
最初の実行を使用して、入力形状が実際のアプリケーション パスで機能するかどうか、評価の結果が意味があるかどうか、および予期される出力形状が意図した目的で機能するかどうかを確認します。
最初の実行後、変更を出荷する際に自信を持てる範囲と形状に到達するまで、拡張と反復を繰り返します。
失敗によりシナリオ、難易度、またはソースが欠落していることが明らかになった場合は、行を追加します。
期待される出力、ルーブリック、入力形状、またはメタデータがあいまいな場合は、行を編集します。
現在のプロンプト、ツール、ポリシー、または製品の動作と一致しなくなった行をアーカイブします。
データセットはアプリケーションとともに進化する必要があります。生産を監視し、データを頻繁にレビューすることにより、

構造化エラー分析を通じて、データセットは時間の経過とともに進化し、システムの実稼働範囲を表すことができます。便利な拡張パターンは 3 つあります。
本番環境のミラーリング : 良いか悪いかに関係なく、本番環境から興味深いケースを追加して、時間の経過とともにデータセットのカバー範囲を拡大します。
不正なトレース拡張: 重大な本番環境の障害が見つかった場合は常に、レビュー済みのデータセット項目を追加します。これは、システムが稼働するとうまく機能し、継続的にトレースをマイニングできます。スコアリングには、LLM-as-a-judge やコード エバリュエーターなどの自動エバリュエーターを使用するか、手動レビューのためにサンプルをアノテーション キューに送信します。
目的別のデータセット: 安定した回帰、敵対的な入力、シングルステップ評価用に個別のデータセットを構築します。
ガイド: ドキュメント チャットボットの評価 開く データセットと評価がどのようにドキュメント チャットボットを探すことができるかについての実践的な例をご覧ください。
まずは 1 つのリリースの質問から始めます。製品分野と、特定の変更を出荷できるかどうかを回答できる最小のエンドツーエンド データセットを選択します。
15 ～ 30 行を使用します。最も一般的なシナリオ、いくつかの高リスクのケース、および 1 つまたは 2 つの既知の障害をカバーします。
すべての行を実行可能および評価可能に保ちます。各入力はアプリケーション パスを通過し、評価によってスコアリングできる必要があります。

[切り捨てられた]

## Original Extract

A practical guide to designing Langfuse datasets with clear scope, realistic inputs, evaluable outputs, useful metadata, and interpretable experiment results.

Designing datasets - Langfuse Langfuse v4: up to 165× faster · Read more Langfuse v4 is here: real-time, up to 165× faster · Read more Docs 🐐 Hiring in Europe and SF Looking for GOATS! Launch App L Get Demo G Docs Integrations Self Hosting Guides Academy Workshop AI Engineering Library Academy Designing datasets Langfuse Academy AI Engineering Loop The Steps
Datasets Designing datasets Experiments Evaluation Supporting materials
Designing datasets Academy Datasets Designing datasets Copy page Designing datasets for AI applications
A dataset is a repeatable set of examples that represent the scope of your application, and that you use to measure and improve your system. By running your application against the same inputs over time, you can track quality with metrics, compare changes, and catch regressions before they affect production users.
Haven't determined what is worth evaluating in a repeatable manner? Check out
error analysis , a structured way to find the failure modes in your application worth checking
the datasets in the Langfuse demo project and in our Academy examples to get an idea of a dataset scope
The Datasets academy section explains the basic building blocks: input, expected output, and metadata. This guide focuses on the design work that happens before and while you create datasets and dataset items.
Dataset design is iterative. A good starting point is a minimally complete dataset: about 15-30 rows that can run through your application, cover the most important input slices, and have an evaluator or review rubric. Run that version early, fix the schema and evaluator, then expand into the gaps you see from those runs or from production input.
Your application will very likely have more than one evaluation dataset. Datasets are often scoped to a specific part of the system or to one sub-step the agent takes.
Iterate and expand 01 Define goal scope / boundary 02 Inspect sources sample / patterns 03 Choose distribution slices / roles 04 Choose eval style reference / free 05 Design schema input / output 06 Build dataset rows / gaps 07 Run experiment failures / iterate
1. Start with the goal of the dataset
Before writing rows, define the smallest useful goal the dataset should support. This keeps the dataset from becoming a vague bucket of interesting examples.
Can we ship a prompt change without breaking support routing?
Does the docs chatbot answer common integration questions well enough?
Does the agent call the right tool when a customer asks for a refund?
A good early starting point for a dataset is looking at the most common examples end to end. This gives you the foundation and the general understanding from which you can expand.
Over time, teams add datasets for single steps, adversarial cases, red teaming, or specific sub-use cases in their application. We have even seen datasets for checking company name spelling.
The goal you set for your dataset defines boundary and job: end-to-end datasets use the payload the application receives, while step-level datasets use the structured state a step sees in production.
If two jobs need different inputs, evaluators, or release decisions, split them. A stable regression dataset and an adversarial-input dataset can both be useful, but combining them without a clear split makes aggregate scores harder to interpret.
Before selecting examples or writing dataset items, inspect a small sample of the material you could use. The goal is to understand what exists and how it's different from what you expected.
Start with three source types:
Production traces: show realistic usage, common paths, and observed failures. Scores, user feedback, tickets, and complaints are discovery signals for useful traces; they are not a separate source type.
Existing assets: old datasets, FAQs, policies, docs, support macros, CSVs, JSON files, and benchmarks can bootstrap known coverage quickly.
Synthetic cases: expert-written and AI-generated examples can fill gaps and give you an idea of what your application is expected to face.
For each source, review input topics, input and output shapes, and failure modes.
You might learn that production traces lack the retrieval context your evaluator needs, that tickets from legacy systems expose better failure labels than traces, or that there is already a good set of end-to-end examples, because support created a 'FAQ' at some point.
Use this input to choose the distribution of inputs, evaluators, and item schema in the next steps.
Guide: Synthetic datasets Open Use synthetic data as a starting point when no other data is available.
3. Choose the input distribution
For a first dataset, keep the input distribution simple. Start with the few slices that tell you what to do after a run:
Scenario type: the main jobs, intents, routes, or task families. For a support-routing dataset, this might be billing, account access, technical issue, and sales request.
Difficulty or risk: routine, ambiguous, hard, adversarial, or business-critical cases. Include more than happy paths, but do not make the first dataset all edge cases.
Dataset role: why the row exists: typical case, known regression, observed failure, or synthetic gap-fill.
For a support-routing dataset, the first version might use a simple scenario-by-difficulty matrix:
Input distribution is the deliberate mix of cases in the dataset: which scenarios appear, how difficult they are, and why each row is included. It is the coverage plan that helps you interpret experiment results by slice.
The distribution does not have to mirror production frequency exactly. A regression dataset may intentionally overrepresent failures, edge cases, or high-value paths. What counts is that the mix is deliberate and visible in metadata.
Add more dimensions only when they change behavior or help interpret results. Channel, language, customer segment, region, context availability, and product area can be useful metadata, but they should not all become balancing constraints on day one.
4. Decide how evaluation will work
Choose the evaluation style before you write expected outputs. This determines what the dataset item needs to contain and how useful the first results will be.
Use reference-based evaluation when each item has a known target: a correct label, expected tool call, required fact, structured output, reference answer, or expected next action. This is the best fit for regression tests and CI gates because failures are easier to inspect. The trade-off is that references take work to write and can become brittle if they over-specify wording instead of behavior.
Use reference-free evaluation when there is no stable expected output, but every item can be judged against the same rule or rubric. This works for checks such as valid JSON, language match, grounding in provided context, safety, or tone. The trade-off is that the evaluator or rubric carries more weight, so ambiguous rubrics produce ambiguous results.
For example, a docs chatbot dataset can use either approach:
Choose the cheapest evaluator that captures the requirement: code evaluators for deterministic checks, LLM-as-a-judge for language-quality judgments, and manual evaluation while you are still learning what good and bad outputs look like.
At this stage, you only need to decide how each dataset item will be reviewed: by an automated evaluator, manual annotation, or both. For the deeper work of designing the evaluator itself, see the Academy page on evaluation .
If you cannot yet name the evaluator or review rubric that will score each row, go for a review workflow first and use annotation queues to label examples manually.
The three dataset item fields are flexible JSON. Here, define them as a concrete contract that your experiment runner, evaluators, and reviewers can all consume.
input : the object you will pass into the system boundary
expectedOutput : only the reference data the evaluator or reviewer needs; omit it for deliberate reference-free evaluation
metadata : stable slice and provenance fields such as source, scenario type, difficulty, dataset role, and review status
Once you have decided on an item schema, enforce it so your team can effectively collaborate on adding items in the right structure .
For a support-routing dataset, one row could look like this:
{
" input " : {
" message " : "I was charged twice for invoice 4831. Can someone fix this?" ,
" channel " : "support_chat" ,
" customer_tier " : "business"
},
" expectedOutput " : {
" route " : "billing_support" ,
" required_actions " : [ "acknowledge_duplicate_charge" , "ask_for_invoice_id" ],
" must_not " : [ "promise_refund_without_review" ]
},
" metadata " : {
" source " : "expert" ,
" scenario_type " : "billing" ,
" difficulty " : "medium" ,
" dataset_role " : "regression" ,
" failure_mode " : "wrong_route"
}
}
The input keeps only the context the router needs. The expectedOutput states the behavior to check: route to billing, ask for the invoice ID, and do not promise a refund before review. The metadata records the row's source, scenario, difficulty, and role.
Keep the schema stable before collecting rows in bulk. Preserve behavior-shaping fields such as conversation history, retrieved context, tool state, routing metadata, or user attributes, but avoid arbitrary per-row fields or natural-language summaries of structured context.
Once the goal, distribution, evaluation method, and schema are concrete, start turning selected source examples into dataset items. Use the sources you inspected earlier, and add them in a format that matches the defined schema.
For a minimally complete first version, choose enough rows to test the whole contract along your input distribution:
common scenarios that should work reliably
a few ambiguous or high-risk scenarios
known failures or regressions you want to prevent
synthetic gap fills only where production traces or existing assets do not cover the distribution
Do not wait until the dataset feels complete. Run the first coherent version in an experiment before adding more rows.
7. Run the first experiment and expand deliberately
Use the first run to check whether the input shape works with the real application path, if the evaluator outcomes make sense, and if the expected output shape works for the intended purpose.
After the first run, expand and iterate until you arrive at scope and shape that helps you feel confident about shipping changes.
Add rows when a failure reveals a missing scenario, difficulty level, or source.
Edit rows when the expected output, rubric, input shape, or metadata is ambiguous.
Archive rows when they no longer match current prompts, tools, policies, or product behavior.
Datasets have to evolve with your application. By monitoring production and frequently reviewing data through structured error analysis , your datasets can evolve over time to represent the production scope of your system. There are three useful expansion patterns:
Production-mirroring : add interesting cases from production no matter if good or bad, to expand the coverage of your dataset over time.
Bad-trace expansion: add a reviewed dataset item whenever you find a serious production failure. This works well once the system is live and you can continuously mine traces. For scoring, use automated evaluators such as LLM-as-a-judge or code evaluators , or send examples to annotation queues for manual review.
Purpose-specific datasets: build separate datasets for stable regression, adversarial inputs, single-step evaluations.
Guide: Docs chatbot evaluation Open See a practical example of how datasets and evaluation can look for a docs chatbot.
Start with one release question. Pick a product area and the smallest end-to-end dataset that can answer whether a specific change can ship.
Use 15-30 rows. Cover the most common scenarios, a few high-risk cases, and one or two known failures.
Keep every row runnable and evaluable. Each input should pass through the application path and be scorable by the evaluat

[truncated]
