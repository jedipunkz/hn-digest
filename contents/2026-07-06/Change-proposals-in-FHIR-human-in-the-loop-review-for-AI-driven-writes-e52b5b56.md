---
source: "https://www.osbytes.io/blog/change-proposals-in-fhir-hitl-review-for-ai-driven-writes"
hn_url: "https://news.ycombinator.com/item?id=48807271"
title: "Change proposals in FHIR: human-in-the-loop review for AI-driven writes"
article_title: "Change proposals in FHIR: human-in-the-loop review for AI-driven writes · osbytes"
author: "dlln"
captured_at: "2026-07-06T17:48:03Z"
capture_tool: "hn-digest"
hn_id: 48807271
score: 3
comments: 0
posted_at: "2026-07-06T16:51:08Z"
tags:
  - hacker-news
  - translated
---

# Change proposals in FHIR: human-in-the-loop review for AI-driven writes

- HN: [48807271](https://news.ycombinator.com/item?id=48807271)
- Source: [www.osbytes.io](https://www.osbytes.io/blog/change-proposals-in-fhir-hitl-review-for-ai-driven-writes)
- Score: 3
- Comments: 0
- Posted: 2026-07-06T16:51:08Z

## Translation

タイトル: FHIR での変更提案: AI 主導の書き込みのための人間参加型レビュー
記事のタイトル: FHIR の変更提案: AI 主導の書き込みの人間参加レビュー · osbytes
説明: AI エージェントが臨床データを変更したいときは、AI エージェントに書き込ませるのではなく、提案させます。変更提案は永続化された FHIR トランザクション バンドルであり、更新は FHIRPath パッチ操作として表現されます。つまり、人間の承認を待つ、レビュー可能、比較可能、監査可能な成果物です。ウォークスルーでは訪問後の HTN を使用します。
[切り捨てられた]

記事本文:
FHIR での変更提案: AI 主導の書き込みに対する人間参加型レビュー · osbytes osbytes ホーム ブログ プロジェクト RSS について 投稿、リポジトリを検索… ⌘K { } 検索
投稿、プロジェクト、メンバーを検索します。
このページの内容 変更提案書の形状
同期と非同期: この提案にはタスクが必要ですか?
更新に FHIRPath パッチを適用する理由
ライフサイクル: 実際に構築しているステート マシン
レンダリング: 汎用 vs. 特注
ドリフト: ドラフトがキュー内にある間にリソースが移動しました
バージョンピンが失敗した場合の自動リベース
イベント駆動型のリベース: ドリフトに反応し、それを待つ必要はありません
アクセス ポリシーと委任: この書き込みを行っているのは誰ですか?
検証: ジェネレーターをどの程度信頼しますか?
原子性と部分承認
実行: 噛み合うメカニカルなディテール
人間も失敗モードです
先行技術なので新規性を感じない
これが複雑になるところ
付録: 導入チェックリスト
FHIR での変更提案: AI 主導の書き込みに対する人間参加型レビュー
臨床データを読み取る AI エージェントには検索の問題があります。臨床データを書き込む AI エージェントには責任問題があります。 「モデルが投薬リストの更新を草案した」と「投薬リストが変更された」の間のどこかで、権限とコンテキストを持った人間は、これから何が起きようとしているのかを正確に見る必要がある。その概要やジェスチャによるチャット記録ではなく、実際の変異である。
インフラストラクチャ エンジニアリングは、何年も前に計画と適用の分割によって同様の問題を解決しました。Terraform は、その理由に従ってクラウド アカウントを変更しません。計画 (意図した変更の完全で差分可能なステートメント) を発行し、適用が実行する前に人間 (またはポリシー エンジン) が計画を承認します。計画はレビュー成果物であり、そのワークフローに関するすべての利点 (差分、ポリシー チェック、ドリフト検出、監査) が含まれます。

計画がデータであることから許可されます。
FHIR ネイティブ アプリケーションでも同じことが可能であり、FHIR はそのための装備が異常に充実しています。この投稿で説明するパターン: エージェントに FHIR API を直接呼び出させる代わりに、エージェントは変更提案 (提案された変更セットをキャプチャするトランザクション バンドル) を生成し、永続化されますが実行されません。作成と削除は単純なエントリです。更新は、リソース全体の PUT ではなく、FHIRPath Patch 操作として表現されます。その場でレビューを行うことができない場合、タスクはバンドルをレビュー エンベロープとしてラップします。レビュー担当者が承認または拒否します。承認されると、サーバーはバンドルをアトミックに実行します。
中心となる動きは表現的なものです。提案は、変更されるデータと同じ型システムのデータです。これにより、レンダリング、検証、差分分析、延期、監査が可能になります。また、順番に検討する価値のある一連の設計上の質問が生じます。それらのほとんどは、満足のいく道を構築した後でのみ表面化するからです。
変更提案の形
FHIR トランザクション バンドルには、チェンジセットに必要なものがすべてすでに含まれています。各エントリには HTTP メソッドと URL を含むリクエストが含まれるため、単一のバンドルで異種の変異セットがキャプチャされます。これは、ジェームス・オコンクォに対する今日の HTN フォローアップ後に訪問文書作成エージェントが作成する可能性のある提案です。リシノプリルの滴定、BMP の指示、フォローアップのスケジュール、ケア プランの BP 目標の更新です。
{
"resourceType" : "バンドル" ,
"タイプ" : "トランザクション" ,
"メタ" : {
「拡張子」:[
{
"url" : "https://example.org/fhir/StructureDefinition/proposal-summary" ,
"valueString" : "今日の HTN フォローアップ後: リシノプリルを毎日 20 mg に増量し、2 週間以内に BMP をオーダーし、4 週間のフォローアップをスケジュールし、ケア プランの血圧目標を <130/80 に更新します。"
}、
{
"url" : "https://example.org/fhir/StructureDefinition/proposal-evidence" ,
「値」を参照

nce" : { "参照" : "Encounter/enc-20260701" }
}、
{
"url" : "https://example.org/fhir/StructureDefinition/proposal-reason" ,
"valueString" : "HTN フォローアップ Encounter/enc-20260701 — 遵守にもかかわらず BP 158/94"
}
]
}、
「エントリ」: [
{
「拡張子」:[
{
"url" : "https://example.org/fhir/StructureDefinition/entry-intent" ,
"valueString" : "リシノプリルを毎日 10 mg から 20 mg に滴定します。遵守にもかかわらず、血圧は 158/94 です。"
}、
{
"url" : "https://example.org/fhir/StructureDefinition/expected-current-value" ,
"値文字列" : "10"
}
]、
「リソース」: {
"resourceType" : "パラメータ" ,
「パラメータ」: [{
"名前" : "操作" ,
「部分」：[
{ "名前" : "タイプ" , "値コード" : "置換" },
{ "name" : "path" , "valueString" : "MedicationRequest.dosagestruct[0].doseAndRate[0].doseQuantity.value" },
{ "名前" : "値" 、 "値10 進数" : 20 }
]
}]
}、
「リクエスト」: {
"メソッド" : "パッチ" ,
"url" : "MedicationRequest/789" ,
"ifMatch" : "W/ \" 12 \" "
}
}、
{
"fullUrl" : "urn:uuid:a2b3c4d5-e6f7-8901-abcd-ef1234567890" ,
"拡張子" : [{
"url" : "https://example.org/fhir/StructureDefinition/entry-intent" ,
"valueString" : "用量増加後に電解質を再検査するよう BMP に命令します。"
}]、
「リソース」: {
"resourceType" : "ServiceRequest" ,
"ステータス" : "アクティブ" 、
"意図" : "順序" 、
"code" : { "coding" : [{ "system" : "http://loinc.org" , "code" : "24320-4" , "display" : "基礎代謝パネル" }] },
"件名" : { "参照" : "患者/456" },
"遭遇" : { "参照" : "遭遇/enc-20260701" },
"occurrenceTiming" : { "repeat" : { "boundsPeriod" : { "start" : "2026-07-15" } } }
}
[切り捨てられた]
順序が定義された混合メソッド。トランザクションは、エントリの順序に関係なく、DELETE、POST、PUT/PATCH、任意の read を処理します。上記のバンドルでは、最初に PATCH がリストされ、その後に POST がリストされます。実行時にサーブ

er は、投薬とケアプランのパッチを適用する前に、引き続き BMP と予約を適用します。
内部参照。 fullUrl の urn:uuid: により、ID が存在する前にエントリが相互に参照できるようになります。サーバーは、それらの参照を実際の ID に割り当てる際に書き換えるものとします (SHALL)。これが、上記の Appointment が、まだ存在しない BMP の ServiceRequest を指す方法です。つまり、フォローアップとラボの注文がアトミックにリンクされています。
エントリーごとの前提条件。 request.ifMatch は、エントリを特定のリソース バージョンに固定します (不一致の場合は 412)。 request.ifNoneExist は作成を条件付きにします。両方の PATCH エントリがピン バージョンより上にあります。 POST の ifNoneExist は、べき等な再試行が必要な場合 (実行セクション) に重要です。
自己記述型メタデータ。メタ拡張には、提案の概要、臨床的理由、および証拠のポインタが含まれます。各エントリには 1 行のインテントが含まれます。これは仕様ではなくパターン規則です。バンドルにはルート拡張要素がありません ( DomainResource ではなく Resource から継承されます)。そのため、メタとエントリが標準の正当なアンカーになります。これらのフィールドは、レビュー担当者が最初に読み取るもの (レンダリング セクション) と、タスクが人間に面するフィールドを投影するもの (次のセクション) です。
一か八かのパッチで期待される現在の値。用量漸増エントリには、20 mg を設定するパッチとともに期待電流値拡張 (10) が含まれています。実行者のベルトとサスペンダーは、エージェントが適用する前に見たものがフィールドにまだ保持されていることを確認します。パターン規則もある。ドリフトセクションで展開。
バンドルはリソース (POST /Bundle) として永続化され、トランザクション エンドポイント (サーバー ベースに対する POST /) に送信されず、それが実行されます。この区別こそが重要なポイントであり、これが仕様で定義されたモードではなく、ローカルな規則であることを正確に言う価値があります。サーバーはバンドを保存します。

このファイルのタイプは、他のリソースを格納するのと同じようにたまたまトランザクションであり、どのエンドポイントに送信されたかを除いて、仕様にはこのファイルを不活性にするものは何もありません。
この規則は、サーバーがブループリントを保存することに同意するという、構築する前にテストする価値のある前提に基づいています。多くのエンタープライズ FHIR サーバー (HAPI FHIR、Azure API for FHIR、Google Cloud Healthcare API) は、他のすべてで実行しているのと同じリソース バリデーターを POST /Bundle で実行します。保存されたトランザクション ブループリントは依然としてバンドル インスタンスであり、そのエントリ配列は公平なゲームです: urn:uuid: エントリ内の相対参照、PATCH 操作を実行するパラメータ リソース、および実行時にのみ意味をなすメタデータは、すべて同一であっても、不正なバンドルとして構造検証をトリップする可能性があります。ペイロードはトランザクション エンドポイントを通過します。ストレージが拒否された場合、2 つのフォールバックによってパターンがそのまま維持されます。プロポーザル JSON をバイナリ (FHIR 形式ですが、$validate に対して不透明で、適切な場所で diff するのが難しい) でラップするか、軽量の DocumentReference または Basic ポインターをキーとするサイドカー ストアにドラフトを保持します。どちらもファーストクラスの Bundle ほどクリーンではありませんが、統合アーティファクトではなく臨床インスタンス用に構築されたバリデーターとの戦いに勝ります。
レビューが延期されると (その場合については後ほど詳しく説明します)、タスクは保存されているバンドルを参照し、レビュー ワークフローを実行します。
{
"resourceType" : "タスク" ,
"ステータス" : "リクエスト済み" ,
"意図" : "提案" ,
"code" : { "text" : "提案された変更を確認する" },
"description" : "今日の HTN フォローアップ後: リシノプリルを毎日 20 mg に増量し、2 週間以内に BMP を指示し、4 週間のフォローアップをスケジュールし、ケア プランの血圧目標を <130/80 に更新します。" 、
"フォーカス" : { "リファレンス" : "バンドル/ドラフト提案-42" },
"for" : { "reference" : "患者/4

56" }、
"リクエスター" : { "リファレンス" : "デバイス/訪問ドキュメント エージェント" },
"reasonCode" : { "text" : "HTN フォローアップ Encounter/enc-20260701 — 遵守にもかかわらず BP 158/94" },
"制限" : { "期間" : { "終了" : "2026-07-12T00:00:00Z" } }
}
Task.description はバンドルの概要拡張子を文字ごとに表し、Task.reasonCode は提案理由拡張子と一致し、Task.for はエントリのターゲットから派生可能であることに注意してください。エンベロープは何も生み出しません。
ここでのタスクは意図的に退屈なものであり、それがポイントです。レビューは、ボルトオンの承認受信箱ではなく、アプリケーションの残りの部分が既に使用しているのと同じワークキュー機構 (割り当て、 Task.owner 、ビジネス ステータス、restriction.period による期限) に配置されます。これらの使用法のうち 2 つは、名前を付ける価値のある慣例です。バンドルを指す Task.focus は擁護可能ですが、バニラではありません (仕様のコメントは、複数のリソースが操作される場合にサブタスクを示唆しています)、および、restriction.period.end (「履行が求められる期間」) をプロポーザルの有効期限として読み取ることは、定義ではなく解釈です。どちらもプロファイルで明確にすべきものです (終わり近くのプロファイルに関するメモを参照してください)。
権限の方向にも注意してください。 Task の人に向けたフィールド ( description 、reasonCode 、 for ) は、Bundle がすでに保持しているメタデータ (proposal-summary 、proposal-reason 、エントリ内の患者参照) から派生する必要があり、独立して作成されるものではないため、提案は 1 つの自己記述型アーティファクトのままであり、エンベロープは使い捨てのままです。この分割については次に詳しく説明します。
同期と非同期: この提案にはタスクが必要ですか?
バンドルとタスクは異なる問題を解決しますが、すべてのワークフローで両方が必要なわけではないため、これらを分離することは価値があります。バンドルは何が変わるのかに答えます。タスクは誰がリビジョンするかを答えます

そうだね、そしていつ。 2 番目が必要かどうかは、提案が生まれたときに人間がどこにいるかによって決まります。
同期 (副操縦士)。臨床医がエージェントとチャットしています。エージェントは会話中にバンドルの下書きを作成します。 UI は diff をインラインでレンダリングします。臨床医が承認する。これは臨床医自身のセキュリティ コンテキストで即座に実行され、委任の質問は簡単に解決されます (承認者はセッション ユーザーです)。タスクは必要ありません。会話はレビューのエンベロープであり、提案の有効期間は数秒で、ドリフトの前提条件は簡単に通過します。これは本質的に、実際のトランザクション バンドルにアップグレードされた CDS フック提案カードです。
非同期（自律）。エージェントはスケジュールまたはトリガーに基づいて実行されます。訪問後の文書は夜間に完成し、受信結果のルーティングが行われます。生成時には人間は存在しません。ルーティングと割り当て、TTL とスーパーセッション、数時間または数日で測定されるレビュー期間にわたるドリフト処理、および作成者、承認者、および実行者が ID と時間の両方で分離されているため、委任機構のタスクなど、完全なエンベロープが維持されるようになりました。
この分割に名前を付けた設計上の利点は、アーティファクトが両方のモードで同一であることです。封筒だけが異なります。バンドルを完全に自己完結型にする — 提案の概要

[切り捨てられた]

## Original Extract

When an AI agent wants to change clinical data, don't let it write — let it propose. A change proposal is a persisted FHIR transaction Bundle, with updates expressed as FHIRPath Patch operations: a reviewable, diffable, auditable artifact awaiting human approval. Walkthrough uses a post-visit HTN fo
[truncated]

Change proposals in FHIR: human-in-the-loop review for AI-driven writes · osbytes osbytes Home Blog Projects About RSS Search posts, repos… ⌘K { } Search
Find posts, projects, and members.
On this page The shape of a change proposal
Synchronous vs. asynchronous: does this proposal need a Task?
Why FHIRPath Patch for updates
Lifecycle: the state machine you're actually building
Rendering: generic vs. bespoke
Drift: the resource moved while the draft sat in the queue
Automatic re-basing when version pins fail
Event-driven re-basing: react to drift, don't wait for it
Access policies and delegation: who is doing this write?
Validation: how much do you trust the generator?
Atomicity and partial approval
Execution: the mechanical details that bite
The human is a failure mode too
Prior art, so you don't feel novel
Where this earns its complexity
Appendix: implementation checklist
Change proposals in FHIR: human-in-the-loop review for AI-driven writes
An AI agent that reads clinical data is a retrieval problem. An AI agent that writes clinical data is a liability problem. Somewhere between "the model drafted a medication list update" and "the medication list changed," a human with authority and context needs to look at exactly what is about to happen — not a summary of it, not a chat transcript that gestures at it, but the actual mutation.
Infrastructure engineering solved an analogous problem years ago with the plan/apply split: Terraform doesn't mutate your cloud account as it reasons; it emits a plan — a complete, diffable statement of intended changes — and a human (or a policy engine) approves the plan before apply executes it. The plan is the review artifact, and everything good about that workflow (diffs, policy checks, drift detection, audit) follows from the plan being data .
FHIR-native applications can do the same thing, and FHIR is unusually well equipped for it. The pattern this post describes: instead of letting the agent call the FHIR API directly, the agent produces a change proposal — a transaction Bundle capturing the proposed changeset, persisted but not executed. Creates and deletes are plain entries. Updates are expressed as FHIRPath Patch operations rather than whole-resource PUT s. When review can't happen on the spot, a Task wraps the Bundle as its review envelope. A reviewer approves or rejects; on approval, the server executes the Bundle atomically.
The core move is representational: the proposal is data, in the same type system as the data it modifies. That buys you rendering, validation, diffing, deferral, and audit — and it raises a set of design questions worth walking through in order, because most of them only surface after you've built the happy path.
The shape of a change proposal
A FHIR transaction Bundle already has everything a changeset needs. Each entry carries a request with an HTTP method and URL, so a single Bundle captures a heterogeneous set of mutations. Here's a proposal a visit-documentation agent might draft after today's HTN follow-up for James Okonkwo — titrate lisinopril, order a BMP, schedule follow-up, and update the care plan BP goal:
{
"resourceType" : "Bundle" ,
"type" : "transaction" ,
"meta" : {
"extension" : [
{
"url" : "https://example.org/fhir/StructureDefinition/proposal-summary" ,
"valueString" : "After today's HTN follow-up: increase lisinopril to 20 mg daily, order BMP in two weeks, schedule four-week follow-up, and update the care plan BP goal to <130/80."
},
{
"url" : "https://example.org/fhir/StructureDefinition/proposal-evidence" ,
"valueReference" : { "reference" : "Encounter/enc-20260701" }
},
{
"url" : "https://example.org/fhir/StructureDefinition/proposal-reason" ,
"valueString" : "HTN follow-up Encounter/enc-20260701 — BP 158/94 despite adherence"
}
]
},
"entry" : [
{
"extension" : [
{
"url" : "https://example.org/fhir/StructureDefinition/entry-intent" ,
"valueString" : "Titrate lisinopril from 10 mg to 20 mg daily — BP 158/94 despite adherence."
},
{
"url" : "https://example.org/fhir/StructureDefinition/expected-current-value" ,
"valueString" : "10"
}
],
"resource" : {
"resourceType" : "Parameters" ,
"parameter" : [{
"name" : "operation" ,
"part" : [
{ "name" : "type" , "valueCode" : "replace" },
{ "name" : "path" , "valueString" : "MedicationRequest.dosageInstruction[0].doseAndRate[0].doseQuantity.value" },
{ "name" : "value" , "valueDecimal" : 20 }
]
}]
},
"request" : {
"method" : "PATCH" ,
"url" : "MedicationRequest/789" ,
"ifMatch" : "W/ \" 12 \" "
}
},
{
"fullUrl" : "urn:uuid:a2b3c4d5-e6f7-8901-abcd-ef1234567890" ,
"extension" : [{
"url" : "https://example.org/fhir/StructureDefinition/entry-intent" ,
"valueString" : "Order BMP to recheck electrolytes after dose increase."
}],
"resource" : {
"resourceType" : "ServiceRequest" ,
"status" : "active" ,
"intent" : "order" ,
"code" : { "coding" : [{ "system" : "http://loinc.org" , "code" : "24320-4" , "display" : "Basic metabolic panel" }] },
"subject" : { "reference" : "Patient/456" },
"encounter" : { "reference" : "Encounter/enc-20260701" },
"occurrenceTiming" : { "repeat" : { "boundsPeriod" : { "start" : "2026-07-15" } } }
}
[truncated]
Mixed methods with defined ordering. A transaction processes DELETE, then POST, then PUT/PATCH, then any reads , regardless of entry order. The Bundle above lists a PATCH first and POSTs later; at execution the server still applies the BMP and appointment creates before the medication and care-plan patches.
Internal references. The urn:uuid: in fullUrl lets entries reference each other before ids exist; the server SHALL rewrite those references to real ids as it assigns them. That's how the Appointment above points at the ServiceRequest for the BMP that doesn't exist yet — follow-up and lab order linked atomically.
Per-entry preconditions. request.ifMatch pins an entry to a specific resource version ( 412 on mismatch ); request.ifNoneExist makes creates conditional. Both PATCH entries above pin versions; ifNoneExist on the POSTs matters when you need idempotent retry (execution section).
Self-describing metadata. The meta extensions carry the proposal's summary, clinical reason, and evidence pointer; each entry carries a one-line intent. This is pattern convention, not spec — Bundle has no root extension element (it inherits from Resource , not DomainResource ), so meta and the entries are the standard-legal anchors. These fields are what the reviewer reads first (rendering section) and what a Task projects its human-facing fields from (next section).
Expected current value on high-stakes patches. The dose-titration entry carries an expected-current-value extension ( 10 ) alongside the patch that sets 20 mg — the executor's belt-and-suspenders check that the field still holds what the agent saw before applying. Also pattern convention; expanded in the drift section.
The Bundle is persisted as a resource — POST /Bundle — not submitted to the transaction endpoint ( POST / against the server base), which is what would execute it. That distinction is the whole trick, and it's worth being precise that it is a local convention, not a spec-defined mode: the server stores a Bundle whose type happens to be transaction the same way it stores any other resource, and nothing in the spec makes it inert except which endpoint it was sent to.
That convention leans on an assumption worth testing before you build on it — that the server will agree to store the blueprint at all. Many enterprise FHIR servers — HAPI FHIR, Azure API for FHIR, Google Cloud Healthcare API — run the same resource validator on POST /Bundle that they run on everything else, and a stored transaction blueprint is still a Bundle instance whose entry array is fair game: urn:uuid: relative references inside entries, Parameters resources carrying PATCH operations, and request metadata that only makes sense at execution time can all trip structural validation as a malformed Bundle even though the identical payload sails through the transaction endpoint. Where storage is refused, two fallbacks keep the pattern intact: wrap the proposal JSON in a Binary (FHIR-shaped, but opaque to $validate and awkward to diff in place), or hold drafts in a sidecar store keyed by a lightweight DocumentReference or Basic pointer. Neither is as clean as a first-class Bundle , but both beat fighting a validator built for clinical instances rather than integration artifacts.
When review is deferred (more on when that's the case in a moment), a Task references the stored Bundle and carries the review workflow:
{
"resourceType" : "Task" ,
"status" : "requested" ,
"intent" : "proposal" ,
"code" : { "text" : "Review proposed changes" },
"description" : "After today's HTN follow-up: increase lisinopril to 20 mg daily, order BMP in two weeks, schedule four-week follow-up, and update the care plan BP goal to <130/80." ,
"focus" : { "reference" : "Bundle/draft-proposal-42" },
"for" : { "reference" : "Patient/456" },
"requester" : { "reference" : "Device/visit-documentation-agent" },
"reasonCode" : { "text" : "HTN follow-up Encounter/enc-20260701 — BP 158/94 despite adherence" },
"restriction" : { "period" : { "end" : "2026-07-12T00:00:00Z" } }
}
Notice Task.description is character-for-character the Bundle's summary extension, Task.reasonCode matches the proposal-reason extension, and Task.for is derivable from the entries' targets — the envelope invents nothing.
Task is deliberately boring here, and that's the point: review lands in the same work-queue machinery the rest of the application already uses (assignment, Task.owner , business status, due dates via restriction.period ), instead of a bolted-on approval inbox. Two of these usages are conventions worth naming: Task.focus pointing at a Bundle is defensible but not vanilla (the spec's comments suggest subtasks when multiple resources are manipulated), and reading restriction.period.end — "the period over which fulfillment is sought" — as a proposal expiry is an interpretation, not a definition. Both are the kind of thing a profile should pin down (see the profiling note near the end).
Note the direction of authority, too — the Task 's human-facing fields ( description , reasonCode , for ) should be derived from metadata the Bundle already carries ( proposal-summary , proposal-reason , and the patient references in the entries), not authored independently, so the proposal remains one self-describing artifact and the envelope stays disposable. More on that split next.
Synchronous vs. asynchronous: does this proposal need a Task ?
The Bundle and the Task solve different problems, and it's worth separating them, because not every workflow needs both. The Bundle answers what will change ; the Task answers who reviews it, and when . Whether you need the second depends on where the human is when the proposal is born:
Synchronous (copilot). A clinician is chatting with the agent; the agent drafts a Bundle mid-conversation; the UI renders the diff inline; the clinician approves; it executes immediately — in the clinician's own security context, which resolves the delegation question trivially (the approver is the session user). No Task needed: the conversation is the review envelope, the proposal's lifetime is seconds, and drift preconditions trivially pass. This is essentially the CDS Hooks suggestion card, upgraded to a real transaction Bundle.
Asynchronous (autonomous). The agent runs on a schedule or a trigger — post-visit documentation finalized overnight, inbound result routing — with no human present at generation time. Now the full envelope earns its keep: Task for routing and assignment, TTL and supersession, drift handling across a review window measured in hours or days, and the delegation machinery, since author, approver, and executor are separated in both identity and time.
The design payoff of naming this split: the artifact is identical in both modes; only the envelope differs. Make the Bundle fully self-contained — the proposal's summary

[truncated]
