---
source: "https://tryevaluator.com"
hn_url: "https://news.ycombinator.com/item?id=48734393"
title: "Show HN: Don't ask if devs cheat with AI, test if they're good with it"
article_title: "Evaluator — Hire engineers who use AI well"
author: "skyepstein"
captured_at: "2026-06-30T15:48:49Z"
capture_tool: "hn-digest"
hn_id: 48734393
score: 1
comments: 0
posted_at: "2026-06-30T15:45:47Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Don't ask if devs cheat with AI, test if they're good with it

- HN: [48734393](https://news.ycombinator.com/item?id=48734393)
- Source: [tryevaluator.com](https://tryevaluator.com)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T15:45:47Z

## Translation

タイトル: HN を表示: 開発者が AI で不正行為をしているかどうかは尋ねず、AI にうまく対応できるかどうかをテストしてください
記事のタイトル: 評価者 — AI をうまく活用するエンジニアを雇う
説明: AI コラボレーション (AI コードの読み取り、修正、プロンプト、オーバーライド) を基礎と並行して評価する技術評価。 AI ネイティブのエンジニアを雇用するために構築されました。

記事本文:
Evaluator — AI をうまく活用するエンジニアを雇用する Evaluator ブログ ログイン サインアップ 2026 年の雇用市場に向けて 現在、すべてのエンジニアが AI を使用しています。
それを上手に使える人を雇いましょう。
Evaluator は、依然として重要な基本である読み取り、書き込み、デバッグに加えて、候補者が AI とどのように連携するか (AI の読み取り、修正、プロンプト、オーバーライド) を評価する技術的な評価です。
AIアシスタントがこれを作成しました。それは合理的に見えます。そうではない。あらゆる欠陥を見つけて修正してください。
非同期関数 fetchUserPosts(userId: string) {
const res = await fetch(`/api/users/${userId}/posts`)
const post = res.json.parse()
return Posts.filter((p, i) => i <= Posts.length )
候補が見つかりました
res.json.parse() — 幻覚。それは await res.json() です。
i <=posts.length — 1 つずつずれます。 < にするか、フィルターを削除してください。
あなたは間違ったものを検査してきたのです。
「候補者はChatGPTを使用していましたか?ブロックし、検出し、ツールを禁止してください。」
「もちろん、彼らは AI を使用しています。問題は、AI を読み取って修正し、指示を出し、幻覚が現れたときにそれを無効にできるかどうかです。」
現在、すべてのショップにコパイロット、カーソル、クロード コードが設置されています。すべてのチームの下位 4 分の 1 が、AI の最初の答えを採用したチームになります。上位 4 分の 1 が幻覚的なインポートを捕捉し、過剰設計されたクラスを書き換えて、実際に機能するものを出荷します。上位 4 分の 1 をテストします。
AI との連携方法に関する 5 つのテスト。
他のプラットフォームではこのようなことはできません。ほとんどの人は依然として AI を検出するものとして扱っています。私たちはそれを採点するためのツールとして扱います。
01 迅速な品質 後輩に説明するように AI にも説明できるでしょうか?
私たちは彼らに機能仕様を与えます。実際に送信するプロンプトを作成します。冗長性ではなく、コンテキスト、制約、エッジケース、および受け入れ基準をスコアリングします。
Postgres-backed /a のデバウンス検索フックを実装する

pi/search エンドポイントはすでに SearchBar.tsx で使用されています。 300ミリ秒のデバウンス。新しい入力で実行中のリクエストをキャンセルします (他の場所で使用している AbortController を使用します)。 { データ、エラー、読み込み中 } を返します。新しいフェッチ ライブラリを導入しないでください。ネイティブ フェッチを使用します。空のクエリのケースをカバーします (早期に返却、リクエストなし)。
02 AIコードの読み取り 「作品」と「良品」は区別できるのか？
AI が作成したコードが実行される様子を見せます。彼らはそれが何をするのかを説明し、AI 型の指示 (過剰設計されたクラス、防御的なトライ/キャッチ、実際のエラー、非慣用的パターン) にフラグを立て、何を変更するかを述べます。
クラス UserDataManager {
プライベート キャッシュ: Map<string, User | null>
コンストラクター() {
this.cache = 新しい Map()
}
async getUserById(id: string | null): Promise<User | null null> {
(!id) が null を返す場合
{を試してください
if (this.cache.has(id)) return this.cache.get(id)!
return await fetchUser(id)
} キャッチ (e) { null を返す }
}
候補者
「関数であるべきもののクラス。エラーを黙って飲み込みます。呼び出し元は、500 と行方不明のユーザーを区別できません。実際にはキャッシュに書き込まないので、ウォームアップすることはありません。」
03 AI コードの修正 彼らは 1 つのバグを外科的に修正できるでしょうか?
AI が作成した関数に現実的なバグを 1 つだけ埋め込みます。彼らはそれを見つけて、最小限のパッチを適用します。実際の問題を見逃す広範なリファクタリングにはペナルティを課します。
関数 paginate(アイテム、ページ、サイズ) {
- const start = ページ * サイズ
+ const start = (ページ - 1) * サイズ
items.slice(開始、開始 + サイズ) を返します。
正解です — 外科的修正です。付随的なリファクタリングはありません。
04 批評 彼らはすべての幻覚をキャッチできるでしょうか?
私たちは彼らに、偽の API、オフバイワン、飲み込まれたエラーなど、複数の欠陥を埋め込んだコードを提供します。私たちは徹底的さを評価します。彼らはすべてを理解しましたか、それとも最初の 1 つだけで止めて「良さそうです」と言ったでしょうか?
lodash.deepFlatten は存在しません — _. flattenDeep は存在します。
catch (e) はエラーを飲み込みます。しょう

少なくともログを記録するか再ス​​ローします。
ループは O(n²) で実行されます — 外側を Set ルックアップに切り替えます。
05 ライブコラボレーション アシスタントとの作業を観察します。
最後の質問では、候補者はエディターに組み込まれた AI サイドバーを受け取ります。私たちは、彼らが送信したすべてのプロンプト、受け入れたすべての提案、拒否したすべてのチャンク、およびその上で行われたすべてのキーストロークを記録します。記録はあなたに送られます。
関数 debouncedSearch(クエリ: 文字列) {
// AIから受け付けた
if (!クエリ) が戻る
if (コントローラー)コントローラー.abort()
// 編集候補: 200 だったのを 300 にしました
タイムアウト = setTimeout(...)
サイドバーのトランスクリプト
あなた: キャンセルには AbortController を使用します
あなた: デバウンスは間違っています - 200 ミリ秒ではなく 300 ミリ秒であるべきです
プロンプト 4 件、承諾 2 件、拒否 1 件、手動編集 38%
5つの基本。さらに、誰もテストしていないものです。
すべての評価は、使用する特定の技術スタック内で、雇用している特定の役割に対して生成されます。質問が変わります。寸法はそうではありません。
5 つのサブテスト: プロンプト品質、AI コードの読み取り、AI コードの修正、批評、ライブ コラボレーション。 AI の流暢さを第一級のスキルとして評価する最初の評価。
実際のコードを解きほぐし、微妙なバグを見つけ、アーキテクチャについて推論します。
仕様通りに実装します。完全な部分コード。コンパイルしてテストに合格する機能を構築します。
乱雑な現実世界のレガシー コード (間違った名前付け、深い入れ子、非表示の状態) のバグを見つけます。
PRの説明。 PMへのリファクタリングの説明。コンパイラではなく、次の人間のために書きます。
構築 vs 購入。 SQL と NoSQL。選択を正当化します。唯一の正解はありません。
JD から有能な候補者まで、一度に。
または、役割を文で説明します。年功序列、技術スタック、候補者が実際に何をするかなどをピックアップします。
30 秒で評価を受けられます。
役割に合わせて調整された、6 つの側面すべてにわたるカスタム テスト。読み取り、書き込み、デバッグ

ジン、コミュニケーション、トレードオフ、AI コラボレーション。
リンクを共有します。採点されたレポートを取得します。
受験者は非同期でテストを受けます。質問ごとのフィードバック、整合性フラグ、および AI 質問の場合は完全なコラボレーション記録を取得できます。
AI が期待される場合には AI を許可します。そうでない場所でそれを捕まえます。
AI コラボレーション セクションでは、サイドバーがすぐそこにあり、私たちはそれをどのように使用しているかをスコアリングしています。 1 つおきのセクションでは、動作分析、キーストローク ペーシング、ペースト パターン、および LLM フィンガープリンティングの候補が、基礎をアウトソーシングしようとしているとフラグを立てます。
サイドバーが表示されます。すべてのプロンプト、承諾、および編集はレビュー担当者のために記録されます。
AI を使用しない質問に対する LLM フィンガープリント
統一された構造、言葉遣い、時間的プレッシャーの中で疑わしいほど洗練された散文。
重要な答えがキーストロークなしで到着しました - ページ外の場所から貼り付けられました。
長時間アイドル状態にしてから 400 CPM バーストしてから送信します。 「Alt キーを押しながら ChatGPT に移動した」フィンガープリント。
1 つの質問中に 5 回以上の焦点が変更される。
無料利用枠は製品全体です。アップグレード前ではなく、実際のボリュームを移動するときにアップグレードしてください。
Proと同じ機能。音量を下げます。ほとんどのチームが何か月も継続する層。
AI コラボレーションを含む 6 つの側面すべて
2019 年のコーディング テストに合格できるエンジニアの雇用をやめてください。
2026 年に、AI を使用しながら、AI にもかかわらず、実際に動作するソフトウェアを出荷できる人材の採用を開始します。
AIと相性の良いエンジニアの採用に。

## Original Extract

The technical assessment that grades AI collaboration — reading AI code, fixing it, prompting it, overriding it — alongside the fundamentals. Built for hiring AI-native engineers.

Evaluator — Hire engineers who use AI well Evaluator Blog Login Sign Up For the 2026 hiring market Every engineer uses AI now.
Hire the ones who use it well .
Evaluator is the technical assessment that grades how skillfully candidates collaborate with AI — reading it, fixing it, prompting it, overriding it — on top of the fundamentals that still matter: reading, writing, debugging.
An AI assistant produced this. It looks reasonable. It is not. Find every flaw and fix it.
async function fetchUserPosts(userId: string) {
const res = await fetch(`/api/users/${userId}/posts`)
const posts = res.json.parse()
return posts.filter((p, i) => i <= posts.length )
} Candidate found
res.json.parse() — hallucinated. It's await res.json() .
i <= posts.length — off-by-one. Should be < or just drop the filter.
You've been screening for the wrong thing.
“Did the candidate use ChatGPT? Block them, detect them, ban the tool.”
“Of course they use AI. The question is whether they can read it, fix it, prompt it, and override it when it hallucinates.”
Every shop now has Copilot, Cursor, Claude Code. The bottom quartile of every team is the one that takes the AI's first answer. The top quartile catches the hallucinated import, rewrites the over-engineered class, and ships something that actually works. We test for the top quartile.
Five tests for how someone works with AI.
No other platform does this. Most still treat AI as a thing to detect. We treat it as a tool to grade.
01 Prompt quality Can they brief an AI like they brief a junior?
We give them a feature spec. They write the prompt they would actually send. We score for context, constraints, edge cases, and acceptance criteria — not for verbosity.
Implement a debounced search hook for the Postgres-backed /api/search endpoint we already use in SearchBar.tsx . 300ms debounce. Cancel in-flight requests on new input (use the AbortController we use elsewhere). Return { data, error, loading } . Don't introduce a new fetch library — we use native fetch. Cover the empty-query case (return early, no request).
02 Reading AI code Can they tell "works" from "good"?
We show them AI-written code that runs. They explain what it does, flag the AI-shaped tells — over-engineered classes, defensive try/catch eating real errors, non-idiomatic patterns — and say what they would change.
class UserDataManager {
private cache: Map<string, User | null>
constructor() {
this.cache = new Map()
}
async getUserById(id: string | null): Promise<User | null> {
if (!id) return null
try {
if (this.cache.has(id)) return this.cache.get(id)!
return await fetchUser(id)
} catch (e) { return null }
}
} Candidate
“A class for what should be a function. Swallows errors silently — caller can't tell a 500 from a missing user. Doesn't actually write to the cache, so it never warms.”
03 Fixing AI code Can they surgically fix one bug?
We plant exactly one realistic bug in an AI-written function. They find it and patch it minimally. We penalize broad refactors that miss the actual problem.
function paginate(items, page, size) {
- const start = page * size
+ const start = (page - 1) * size
return items.slice(start, start + size)
} Correct — surgical fix. No collateral refactor.
04 Critique Can they catch every hallucination?
We give them code with multiple planted flaws — fake APIs, off-by-ones, swallowed errors. We grade thoroughness: did they catch them all, or did they stop at the first one and say "looks good"?
lodash.deepFlatten doesn't exist — _.flattenDeep does.
catch (e) swallows the error. Should at least log or rethrow.
Loop runs O(n²) — switch the outer to a Set lookup.
05 Live collaboration Watch them work with the assistant.
On the final question, the candidate gets an AI sidebar built into the editor. We record every prompt they send, every suggestion they accept, every chunk they reject, and every keystroke they make on top. The transcript goes to you.
function debouncedSearch(query: string) {
// accepted from AI
if (!query) return
if (controller) controller.abort()
// candidate edit: was 200, made it 300
timeout = setTimeout(...)
} Sidebar transcript
You: use AbortController for cancellation
You: debounce is wrong — should be 300ms not 200ms
4 prompts · 2 accepts · 1 reject · 38% manual edits
Five fundamentals. Plus the one nobody else tests.
Every assessment is generated for the specific role you're hiring for, in the specific tech stack you use. The questions change. The dimensions don't.
Five sub-tests: prompt quality, reading AI code, fixing AI code, critique, and live collaboration. The first assessment that grades AI fluency as a first-class skill.
Untangle real code, spot subtle bugs, reason about architecture.
Implement to spec. Complete partial code. Build features that compile and pass tests.
Find bugs in messy, real-world legacy code — bad naming, deep nesting, hidden state.
PR descriptions. Explaining a refactor to a PM. Writing for the next human, not the compiler.
Build vs buy. SQL vs NoSQL. Justify the choice — there is no single right answer.
From a JD to a scored candidate, in one sitting.
Or describe the role in a sentence. We pick up seniority, tech stack, and what the candidate will actually be doing.
Get an assessment in 30 seconds.
A custom test across all six dimensions, calibrated to the role. Reading, writing, debugging, communication, tradeoffs, AI collaboration.
Share a link. Get a scored report.
Candidates take the test async. You get per-question feedback, integrity flags, and — for AI questions — the full collaboration transcript.
We allow AI where it's expected. We catch it where it isn't.
On the AI Collaboration section, the sidebar is right there — we're scoring how they use it. On every other section, behavioral analysis, keystroke pacing, paste pattern, and LLM fingerprinting flag candidates trying to outsource the fundamentals.
Sidebar visible. Every prompt, accept, and edit logged for the reviewer.
LLM fingerprint on a no-AI question
Uniform structure, hedging language, suspiciously polished prose under time pressure.
Non-trivial answer arrived with zero keystrokes — pasted from somewhere off-page.
Long idle, then a 400-CPM burst, then submit. The 'they alt-tabbed to ChatGPT' fingerprint.
Five or more focus changes during a single question.
The free tier is the full product. Upgrade when you're moving real volume, not before.
Same features as Pro. Lower volume. The tier most teams stay on for months.
All six dimensions, including AI Collaboration
Stop hiring engineers who can ace a 2019 coding test.
Start hiring the ones who can ship working software in 2026 — with AI, around AI, despite AI.
For hiring engineers who work well with AI.
