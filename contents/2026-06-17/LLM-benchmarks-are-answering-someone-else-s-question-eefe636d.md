---
source: "https://danlevy.net/llm-evals-are-broken/"
hn_url: "https://news.ycombinator.com/item?id=48572463"
title: "LLM benchmarks are answering someone else's question"
article_title: "Fight Evils with Evals!"
author: "justsml"
captured_at: "2026-06-17T16:19:51Z"
capture_tool: "hn-digest"
hn_id: 48572463
score: 3
comments: 0
posted_at: "2026-06-17T16:07:51Z"
tags:
  - hacker-news
  - translated
---

# LLM benchmarks are answering someone else's question

- HN: [48572463](https://news.ycombinator.com/item?id=48572463)
- Source: [danlevy.net](https://danlevy.net/llm-evals-are-broken/)
- Score: 3
- Comments: 0
- Posted: 2026-06-17T16:07:51Z

## Translation

タイトル: LLM ベンチマークが他の人の質問に答えています
記事タイトル: エヴァルと一緒に悪と戦おう！
説明: ベンチマークはベンチマークを測定します。システムには独自の対策が必要です。

記事本文:
エヴァルと一緒に悪と戦おう！
DanLevy.net
DanLevy.net
DanLevy.net
EN ⌄ English en العربية ar Deutsch de Español es Français fr עברית he हिन्दी hi Italiano it 日本語 ja Русский ru 中文 zh
記事 ⌄ クイズ ダンのチャレンジに挑戦してみよう！
カテゴリ コード 15 AI 10 ガイド 10 セキュリティ 5 DevOps 4 考察 3
人気の高い llm:// 接続文字列の時間です 3 か月前 Axios は必要ないかもしれません 2 年前 Algolia™ は必要ないかもしれません 約 1 年前 LLM に数学を行うよう求めるのはやめてください 5 か月前 2025 年のデータベース イノベーションの波 9 か月前 物事に良い名前を付けます 約 2 年前 単一目的の人々に気をつけてください 約 1 年前 クイズ: Bash とシェルの習得 1 年以上前
最近の侵害 13 日前 Postgres テキスト検索ガイド 2026 約 1 か月前 Eval で悪と戦いましょう! 28 日前 友人や恋人を獲得するためのセマンティック ベクトル検索とその他のトピック 30 日前 llm:// 接続文字列の時間です 3 ヶ月前 AI アシスタントがシェル アクセスを与えてくれました 4 ヶ月前 LLM に数学を行うよう求めるのをやめる 5 ヶ月前
プロジェクト ⌄ デモと例 私のプロジェクト、実験、各種リポジトリのセレクション。
Open Source Journal 私のオープンソースへの貢献、プロジェクト、実験の日記。
DataAnalyzer.app JSON または CSV 入力用のコード + スキーマ ジェネレーター。
Functional Promises ネイティブ JavaScript Promise を中心に構築された機能的で流暢な API。
Node Streaming Image Proxy Node.js 用の高性能画像サイズ変更およびストリーミング プロキシ。
ファクト サービス 複数のデータベース アダプターを備えた強力なキーと値のサービス。
モダン アプリ スターター ベース TS、Vite、React、Tailwind CSS などを使用したモダン アプリ スターター。
⌄ ダン・レヴィ・コーダーについて |リーダー
考える人 |いじくり屋
お問い合わせ <フォーム> Twitter GitHub LinkedIn OSS ログ 履歴書
ベンチマークはベンチマークを測定します。あなたの

システムには独自の対策が必要です。
AI の請求額が必要以上に高くなっています。
私は、チームが無駄なトークン、誤ってルーティングされたモデル呼び出し、キャッシュのギャップ、AI 支出を密かに膨らませる使用パターンを特定できるよう支援します。
AI への支出について話し合う 電話のスケジュールを設定する すべての新モデルは、数多くのベンチマークを身に着けて登場します。
MMLU: 92.4%。 HumanEval: 87.2%。 LLeMU: 88.7%。数学: 73.6%。 AGI: 127%!
しかし、AI を使用してプロセスと製品を構築している企業の 99% にとって、それはどれも重要ではありません。
何が重要ですか?あなたのワークロードはどうですか?良くなったのか、悪くなったのか？それを知る唯一の賢明な方法は、システムの特定のタスク、データ、障害モードを反映する Eval (LLM のテスト) を作成することです。
ベンチマークは嘘をつきません。彼らは他の人の質問に答えています。
「バイブスベースの評価」に実際にかかる費用
標準的なアプローチ: モデル変更を出荷し、苦情チャンネルを監視し、部屋が騒がしくなったら元に戻します。
これでは、興味深いものがほとんどすべて失われます。
聞こえるのは大きな失敗だけです。自信を持って間違った答えを導き出しながらも、それに気づいていないユーザーはいますか?静けさ。より悪い回答を受け取り、機能を放棄するユーザーはいますか?静けさ。サポート チケットとエラー率は、品質低下のほんの一部しか捉えていません。
後退と改善を区別することはできません。新しいモデルがタスク A では優れていて、タスク B では劣っている場合、B に関する苦情は、一般的な「AI が悪化した」というフィードバックと同じように見えます。何を直せばいいのかわかりません。
ユーザーをテスト インフラストラクチャとして使用しています。彼らはそれには登録しませんでした。
評価範囲 (そしてほとんどのチームが間違っている点)
評価アプローチは、「高速だが脆弱」から「高価だが有効」までさまざまです。
審査員としての LLM は現在の最愛の人です。強力なモデルに別のモデルの出力を採点するよう依頼します。高速、スケーラブル、そして安価。問題: グレーダー モデルのバイアスを焼き込んでしまう可能性があります。

ゲーム化され、循環依存関係が作成されます。 GPT-5 を使用して GPT-5 の出力を採点する場合、「GPT-5 が GPT-5 とどの程度一致するか」のようなものを測定することになります。それは何もないわけではありませんが、それはあなたが考えているものではありません。
人間の評価は、誰もが無視しようとする至極の基準です。人間が出力を評価するのは費用がかかり、時間がかかり、評価者間で一貫性がなく、スケジュールを立てるのが面倒です。しかし、それはあなたのシステムが実際の人間にとって役立つかどうかを検証する唯一のものです。
タスク固有の自動チェックには、ほとんどのチームがより多くの時間を費やす必要があります。これらは魅力的ではありませんが、高速かつ決定的であり、システム内で重要なものと結びついています。
1. 出荷前に障害を定義する
モデルやプロンプトを変更する前に、どのような問題があるかを書き留めてください。具体的には。
「出力は正確でなければなりません」ではありません。それはテストではありません。もっと似たもの:
構造化された JSON 出力はエラーなしで解析される必要があります
応答内のすべての引用は、取得されたコンテキスト内でそのまま表示される必要があります
回答では競合製品名に言及してはなりません
SQL クエリは構文的に有効であり、スキーマ内に存在するテーブルのみを参照する必要があります。
感情分類は、既存のテストセットで 3% を超えてポジティブからネガティブに反転してはなりません
これらはプログラムで確認できます。審査員モデルは必要ありません。
評価ハーネス: 決定論的チェック
type EvalResult = {渡された : ブール値 ;理由 ?: 文字列 };
const evals:Record<string,(output:string,context:EvalContext)=>EvalResult>={ //JSON は validJson を解析する必要があります:(output)=>{try{JSON.解析 (出力); return { 渡されました : true }; catch (e) { return { 渡された : false 、理由 : `無効な JSON: ${ e.message } ` }; } }、
// 幻覚的な引用は禁止 — すべての主張はコンテキスト内に出現する必要がある

目的 = extractCitations (出力); const ungrounded = クレーム。 filter ( (claim ) => !retrievedChunks.some (( chunk ) => chunk.includes (claim)) ); ungrounded.length === 0 を返しますか? { 渡されました : true } : { 渡されました : false 、理由 : `根拠のない主張: ${ 根拠がありません。結合 ( ' , ' ) } ` }; }、
// 応答の長さの健全性チェック — 切り捨てまたは暴走生成を検出するreasonableLength : ( Output ) => { constwords = 出力。分割 ( / \s + / ).length;単語 >= 10 && 単語 <= 2000 を返します ? { 渡された : true } : { 渡された : false 、理由 : `単語数 ${ 単語 } が範囲外です` }; }、}; EvalResult> = { // JSON は validJson を解析する必要があります: (output) => { try { JSON.parse(output); return { pass: true }; } catch (e) { return { returns: false,reason: `Invalid JSON: ${e.message}` } }; // すべてのクレームはコンテキスト内に表示される必要があります groundedCitations: (output, {retrievedChunks }) => { constclaims = extractCitations(output); const ungrounded =claims.filter( (claim) => !retrievedChunks.some((chunk) => chunk.includes(claim)) ); return ungrounded.length === 0 ? { 渡された: true } : { 渡された: false, 理由: `根拠のないクレーム: ${ungrounded.join(', ')}` }; // 応答の長さの健全性チェック — 切り捨てまたは暴走生成をキャッチする合理的な長さ: (出力) => { const Words = Output.split(/\s+/).length を返す ワード >= 10 && ワード
2. 最悪の日々からゴールデンセットを構築する
最も優れた評価データは、誰かにチケットを提出させたり、幻覚のスクリーンショットを撮らせたり、静かに機能の使用を中止させたりする出力など、恥ずかしいものです。
ユーザーが不正な出力を報告したり、幻覚にフラグを立てたり、手動で失敗に気づいたりするたびに、それをゴールデン セット (入力、コンテキスト、正しい動作) に追加します。 50 ～ 100 のケースを保持し、すべてのモデル ch で実行します

アンジェ。
これは最初は手動のように感じます。 6 か月後には、すべてのケースが自分自身の失敗履歴から生じているため、公開ベンチマークでは検証できないテスト スイートが完成します。
インターフェイス GoldenCase { ID : 文字列 ;入力: 文字列; context : レコード < string 、不明 >; ExpectedBehavior : {mustContain ?: string []; mustNotContain ?: 文字列 []; StructureCheck ?: (出力: 文字列) => ブール値 ; minSimilarityToReference ?: 数値 ; // 参照回答とのコサイン類似度 };ソースインシデント ?: 文字列 ; // バグレポートまたはチケットへのリンク } ; ExpectedBehavior: {mustContain?: string[]; mustNotContain?: string[]; StructureCheck?: (output: string) => boolean; // 参照回答とのコサイン類似度 }; // バグレポートまたはチケットへのリンク }">
3. 単なる受け入れテストではなく回帰テスト
ほとんどのチームは、モデルの変更を検討する場合にのみ評価を実行します。それが受け入れテストです。「この新しいものは十分ですか?」
また、回帰テストも必要です。「これにより、以前は機能していたものが壊れていないか?」
モデルの変更だけでなく、すべてのプロンプト変更でゴールデン セットを実行します。新しいツールを追加したり、RAG 取得戦略を変更したり、コンテキスト テンプレートを更新したりすると、正常に動作していたプロンプトが静かに機能低下する可能性があります。ベースラインがなければわかりません。 Langfuse のようなツールは評価スコアを運用トレースに添付するため、インシデント レポートだけでなくダッシュボードにも回帰が表示されます。
async function CompareModelVersions(goldenCases:GoldenCase[],baselinePipeline:Pipeline,candidatePipeline:Pipeline) {const results = await Promise 。 all ( goldCases.map ( async ( tc ) => { const [baseline,candidate] = await Promise .all ([baselinePipeline.run (tc.input, tc.context),candidatePipeline.run (tc.input, tc.context), ]);

return { id : tc.id,baselinePassed : runEvals (baseline, tc.expectedBehavior),candidatePassed : runEvals (candidate, tc.expectedBehavior), regression : /* ベースラインは合格 */ && /* 候補は失敗しました */ , 改善 : /* ベースラインは失敗しました */ && /* 候補は合格しました */ , }; }) );
const 回帰 = 結果。フィルター (( r ) => r.回帰);定数の改善 = 結果。フィルター (( r ) => r.improvement);
コンソール。 log ( `回帰: ${ regressions.length } / ${ goldCases.length } ` );コンソール。 log ( `改善: ${ Improvements.length } / ${ goldCases.length } ` );
if (regressions.length > 0 ) { コンソール。 error ( ' ブロック回帰が見つかりました: ' );回帰。 forEach (( r ) => console.error ( ` - ${ r.id } ` )); }
return { 回帰、改善 }; } { const [ベースライン, 候補] = await Promise.all([baselinePipeline.run(tc.input, tc.context),candidatePipeline.run(tc.input, tc.context), ]); return { id: tc.id,baselinePassed: runEvals(baseline, tc.expectedBehavior),candidatePassed: runEvals(candidate, tc.expectedBehavior), 回帰: /* ベースラインは合格しました */ && /* 候補は失敗しました */, 改善: /* ベースラインは失敗しました */ && /* 候補は合格しました */, }); const regressions = results.filter((r) => r.improvement); console.log(`Regressions: ${regressions.length} / ${goldenCases.length}`); console.log(`改善: ${improvements.length} / ${goldenCases.length}`); if (regressions.length > 0) { console.error('ブロックされているリグレッションが見つかりました:'); regressions.forEach((r) => console.error(` - ${r.id}`); } return { 回帰、改善 };
[切り捨てられた]
4. たった 1 つのことに対して LLM を判断者として使用する
LLM-as-judge は、決定的な正解がない、オープンエンドの出力に役立ちます。

「回答は役に立ちましたか？」、「この要約は重要なポイントを保っていますか？」、「この説明は初心者にとって適切ですか？」
そこで使ってください。決定的な答えには使用しないでください。実際に使用する場合は、採点ルーブリックを明示的にしてください。
評価ハーネス: ルーブリックベースのジャッジ
非同期関数 judgeHelpfulness (userQuery : string , modelResponse : string ) : Promise <{ スコア : 数値 ; reasoning : string }> { const judgePrompt = ` カスタマー サポートの応答を評価しています。
ユーザーの質問: ${ userQuery } 応答: ${ modelResponse }
回答を 1 ～ 5 のスケールで評価します。 5 = 正確で実用的な情報で質問に直接答えます 4 = 質問には答えますが、より具体的または実用的なものになる可能性があります 3 = 質問に部分的に対処します。重要な情報が欠落している 2 = 関係はあるが質問には答えていない 1 = 主題から外れている、事実が間違っている、または有害である
JSON で応答します: {"スコア": <数値>, "推論": "<一文>"} ` ;
const result = ジャッジモデルを待ちます。生成 (judgePrompt); JSON を返します。解析（結果）; } { const judgePrompt = ` カスタマー サポートの応答を評価しています。ユーザーの質問: ${userQuery} 応答: ${modelResponse} 応答を 1 ～ 5 のスケールで評価します。 5 = 正確で実用的な情報で質問に直接答えます 4 = 質問にはより具体的または実用的である可能性があります 3 = 部分的に

[切り捨てられた]

## Original Extract

Benchmarks measure benchmarks. Your system needs its own measures.

Fight Evils with Evals!
DanLevy.net
DanLevy.net
DanLevy.net
EN ⌄ English en العربية ar Deutsch de Español es Français fr עברית he हिन्दी hi Italiano it 日本語 ja Русский ru 中文 zh
Articles ⌄ Quizzes Try Dan's Challenges!
Categories Code 15 AI 10 Guides 10 Security 5 DevOps 4 Thoughts 3
Popular It's Time for llm:// Connection Strings 3 months ago You may not need Axios almost 2 years ago You May Not Need Algolia™ about 1 year ago Stop Asking LLMs to Do Math 5 months ago 2025's Wave of Database Innovation 9 months ago Naming things good almost 2 years ago Beware the Single-Purpose People about 1 year ago Quiz: Bash & Shell Mastery over 1 year ago
Recent Into the Breach 13 days ago Postgres Text Searching Guide 2026 about 1 month ago Fight Evils with Evals! 28 days ago Semantic Vector Search and Other Topics to Win Friends and Lovers 30 days ago It's Time for llm:// Connection Strings 3 months ago Your AI Assistant Gave Me Shell Access 4 months ago Stop Asking LLMs to Do Math 5 months ago
Projects ⌄ Demos & Examples A selection of my projects, experiments and assorted repos.
Open Source Journal A journal of my open source contributions, projects, and experiments.
DataAnalyzer.app A code + schema generator for JSON or CSV input.
Functional Promises A functional and fluent API built around native JavaScript promises.
Node Streaming Image Proxy High performance image resizing and streaming proxy for Node.js.
Fact Service A powerful key-value service with several database adapters.
Modern App Starter Base A modern app starter using TS, Vite, React, Tailwind CSS, and more.
About ⌄ Dan Levy Coder | Leader
Thinker | Tinkerer
Contact Me <form> Twitter GitHub LinkedIn OSS Log Resume
Benchmarks measure benchmarks. Your system needs its own measures.
Your AI bill is higher than it needs to be.
I help teams spot wasted tokens, misrouted model calls, caching gaps, and usage patterns that quietly inflate AI spend.
Talk through your AI spend Schedule a call Every new model arrives wearing a tuxedo of benchmarks.
MMLU: 92.4%. HumanEval: 87.2%. LLeMU: 88.7%. MATH: 73.6%. AGI: 127%!
Yet, for 99% of businesses building process & product with AI, none of it matters.
What matters? How are YOUR workloads doing? Getting better or worse? The only sane way to know that is to write Evals (tests for LLMs) that reflect the specific tasks, data, and failure modes of your system.
The benchmarks are not lying. They are answering someone else’s question.
What “Vibes-Based Evaluation” Actually Costs
The standard approach: ship a model change, watch the complaint channels, roll back if the room gets loud.
That misses almost everything interesting:
You only catch loud failures. Users who get a confidently wrong answer and don’t realize it? Silent. Users who get a worse answer and abandon the feature? Silent. Support tickets and error rates capture only a fraction of quality regression.
You can’t distinguish regressions from improvements. If the new model is better at task A and worse at task B, complaints about B look identical to generic “the AI got worse” feedback. You don’t know what to fix.
You’re using your users as test infrastructure. They didn’t sign up for that.
The Eval Spectrum (and Where Most Teams Get It Wrong)
Evaluation approaches sit on a spectrum from “fast but flimsy” to “expensive but valid.”
LLM-as-judge is the current darling: ask a powerful model to grade another model’s outputs. Fast, scalable, cheap. The problem: it bakes in the grader model’s biases, can be gamed, and creates a circular dependency. If you use GPT-5 to grade GPT-5’s outputs, you’re measuring something like “how much does GPT-5 agree with GPT-5.” That’s not nothing, but it’s not what you think.
Human eval is the gold standard everyone tries to skip. Getting humans to evaluate outputs is expensive, slow, inconsistent across evaluators, and annoying to schedule. But it is the only thing that validates whether your system is useful to real humans.
Task-specific automated checks are where most teams should spend more time. They are not glamorous, but they are fast, deterministic, and tied to what matters in your system.
1. Define Failure Before You Ship
Before changing a model or prompt, write down what bad looks like. Specifically.
Not “the output should be accurate.” That’s not a test. More like:
Structured JSON output must parse without errors
All citations in the response must appear verbatim in the retrieved context
Responses must not mention competitor product names
SQL queries must be syntactically valid and reference only tables that exist in the schema
Sentiment classification must not flip from positive to negative more than 3% of the time on the existing test set
You can check these programmatically. No judge model required.
Eval harness: deterministic checks
type EvalResult = { passed : boolean ; reason ?: string };
const evals : Record < string , ( output : string , context : EvalContext ) => EvalResult > = { // JSON must parse validJson : ( output ) => { try { JSON . parse (output); return { passed : true }; } catch (e) { return { passed : false , reason : `Invalid JSON: ${ e.message } ` }; } },
// No hallucinated citations — every claim must appear in context groundedCitations : ( output , { retrievedChunks }) => { const claims = extractCitations (output); const ungrounded = claims. filter ( ( claim ) => ! retrievedChunks. some (( chunk ) => chunk. includes (claim)) ); return ungrounded.length === 0 ? { passed : true } : { passed : false , reason : `Ungrounded claims: ${ ungrounded. join ( ' , ' ) } ` }; },
// Response length sanity check — catch truncation or runaway generation reasonableLength : ( output ) => { const words = output. split ( / \s + / ).length; return words >= 10 && words <= 2000 ? { passed : true } : { passed : false , reason : `Word count ${ words } out of bounds` }; }, }; EvalResult> = { // JSON must parse validJson: (output) => { try { JSON.parse(output); return { passed: true }; } catch (e) { return { passed: false, reason: `Invalid JSON: ${e.message}` }; } }, // No hallucinated citations — every claim must appear in context groundedCitations: (output, { retrievedChunks }) => { const claims = extractCitations(output); const ungrounded = claims.filter( (claim) => !retrievedChunks.some((chunk) => chunk.includes(claim)) ); return ungrounded.length === 0 ? { passed: true } : { passed: false, reason: `Ungrounded claims: ${ungrounded.join(', ')}` }; }, // Response length sanity check — catch truncation or runaway generation reasonableLength: (output) => { const words = output.split(/\s+/).length; return words >= 10 && words
2. Build a Golden Set From Your Worst Days
Your best evaluation data is the embarrassing stuff: the outputs that made someone file a ticket, screenshot a hallucination, or quietly stop using the feature.
Every time a user reports a bad output, flags a hallucination, or you notice a failure manually, add it to your golden set: the input, the context, and the correct behavior. Keep 50-100 cases and run them on every model change.
This feels manual at first. After six months, you have a test suite no public benchmark can game, because every case came from your own failure history.
interface GoldenCase { id : string ; input : string ; context : Record < string , unknown >; expectedBehavior : { mustContain ?: string []; mustNotContain ?: string []; structureCheck ?: ( output : string ) => boolean ; minSimilarityToReference ?: number ; // cosine similarity to a reference answer }; sourceIncident ?: string ; // link back to the bug report or ticket } ; expectedBehavior: { mustContain?: string[]; mustNotContain?: string[]; structureCheck?: (output: string) => boolean; minSimilarityToReference?: number; // cosine similarity to a reference answer }; sourceIncident?: string; // link back to the bug report or ticket}">
3. Regression Testing, Not Just Acceptance Testing
Most teams run evals only when considering a model change. That’s acceptance testing: “is this new thing good enough?”
You also need regression testing: “did this break something that used to work?”
Run your golden set on every prompt change, not just model changes. A prompt that was working fine can silently degrade when you add a new tool, change a RAG retrieval strategy, or update your context template. You won’t know without a baseline. Tools like Langfuse attach eval scores to production traces so regression shows up in dashboards, not just in incident reports.
async function compareModelVersions ( goldenCases : GoldenCase [], baselinePipeline : Pipeline , candidatePipeline : Pipeline ) { const results = await Promise . all ( goldenCases. map ( async ( tc ) => { const [baseline, candidate] = await Promise . all ([ baselinePipeline. run (tc.input, tc.context), candidatePipeline. run (tc.input, tc.context), ]);
return { id : tc.id, baselinePassed : runEvals (baseline, tc.expectedBehavior), candidatePassed : runEvals (candidate, tc.expectedBehavior), regression : /* baseline passed */ && /* candidate failed */ , improvement : /* baseline failed */ && /* candidate passed */ , }; }) );
const regressions = results. filter (( r ) => r.regression); const improvements = results. filter (( r ) => r.improvement);
console. log ( `Regressions: ${ regressions.length } / ${ goldenCases.length } ` ); console. log ( `Improvements: ${ improvements.length } / ${ goldenCases.length } ` );
if (regressions.length > 0 ) { console. error ( ' Blocking regressions found: ' ); regressions. forEach (( r ) => console. error ( ` - ${ r.id } ` )); }
return { regressions, improvements }; } { const [baseline, candidate] = await Promise.all([ baselinePipeline.run(tc.input, tc.context), candidatePipeline.run(tc.input, tc.context), ]); return { id: tc.id, baselinePassed: runEvals(baseline, tc.expectedBehavior), candidatePassed: runEvals(candidate, tc.expectedBehavior), regression: /* baseline passed */ && /* candidate failed */, improvement: /* baseline failed */ && /* candidate passed */, }; }) ); const regressions = results.filter((r) => r.regression); const improvements = results.filter((r) => r.improvement); console.log(`Regressions: ${regressions.length} / ${goldenCases.length}`); console.log(`Improvements: ${improvements.length} / ${goldenCases.length}`); if (regressions.length > 0) { console.error('Blocking regressions found:'); regressions.forEach((r) => console.error(` - ${r.id}`)); } return { regressions, improvements };
[truncated]
4. Use LLM-as-Judge for Exactly One Thing
LLM-as-judge is useful for open-ended outputs where there is no deterministic right answer: “is this response helpful?”, “does this summary preserve the key points?”, “is this explanation right for a beginner?”
Use it there. Don’t use it for deterministic answers. When you do use it, make the grading rubric explicit:
Eval harness: rubric-based judge
async function judgeHelpfulness ( userQuery : string , modelResponse : string ) : Promise <{ score : number ; reasoning : string }> { const judgePrompt = ` You are evaluating a customer support response.
User question: ${ userQuery } Response: ${ modelResponse }
Rate the response on a scale of 1-5: 5 = Directly answers the question with accurate, actionable information 4 = Answers the question but could be more specific or actionable 3 = Partially addresses the question; key information is missing 2 = Tangentially related but doesn't answer the question 1 = Off-topic, factually wrong, or harmful
Respond with JSON: {"score": <number>, "reasoning": "<one sentence>"} ` ;
const result = await judgeModel. generate (judgePrompt); return JSON . parse (result); } { const judgePrompt = `You are evaluating a customer support response.User question: ${userQuery}Response: ${modelResponse}Rate the response on a scale of 1-5:5 = Directly answers the question with accurate, actionable information4 = Answers the question but could be more specific or actionable3 = Partially

[truncated]
