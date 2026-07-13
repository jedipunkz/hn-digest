---
source: "https://www.agentsproof.dev/"
hn_url: "https://news.ycombinator.com/item?id=48898280"
title: "Show HN: AgentsProof – a small project for testing AI agents"
article_title: "AgentsProof — One decorator. A shareable eval report."
author: "adeeonline"
captured_at: "2026-07-13T20:50:41Z"
capture_tool: "hn-digest"
hn_id: 48898280
score: 1
comments: 1
posted_at: "2026-07-13T20:23:14Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AgentsProof – a small project for testing AI agents

- HN: [48898280](https://news.ycombinator.com/item?id=48898280)
- Source: [www.agentsproof.dev](https://www.agentsproof.dev/)
- Score: 1
- Comments: 1
- Posted: 2026-07-13T20:23:14Z

## Translation

タイトル: Show HN: AgentsProof – AI エージェントをテストするための小さなプロジェクト
記事のタイトル: AgentsProof — 1 人のデコレータ。共有可能な評価レポート。
説明: AgentsProof を 30 秒以内にエージェントにドロップします。機能することを証明する公開 URL を取得します。目標の完了、ツールの精度、効率、品質、安全性についてスコアが付けられます。
HN テキスト: 私はエージェントの評価と可観測性をより深く理解するための学習サイド プロジェクトとして AgentsProof を構築しました。これにより、実行を追跡し、評価ケースを作成し、エージェントがどこで成功したか失敗したかを確認できます。特別なことは何もありませんが、エージェントの実行に関するパブリックに共有可能なレポートを作成します。

記事本文:
AgentsProof — 1 つのデコレータ。共有可能な評価レポート。 AP AgentsProof の機能 料金例 ログイン GitHub 現在ベータ版 バイブチェックを停止
AIエージェント。それらが機能することを証明してください。
数回の SDK 呼び出しでエージェントのトレースをキャプチャし、カスタム ルールと承認されたテスト ケースに照らして評価し、共有可能な証明レポートを数分で作成します。
プロンプトを変更しました。何かが壊れました。ユーザーが最初に気づきました。
レポートの例を表示 $ npm install Agentsproof TS PY Agentsproof.dev/r/x8kp2m live Agent.ts コードをインポート { AgentsProof } from 'agentsproof' ;
const ap = new AgentsProof({ apiKey: process.env.AGENTSPROOF_API_KEY });
const run = ap.startRun({
projectSlug: 'my-agent' 、
入力: { クエリ }、
goal: 'Web を検索し、動作するコード ソリューションを返します' , // アンカー グレーディング
});
const result = await run.trace( 'llm_call' , 'gpt-4o' , () => llm(query));
const { publicUrl } = await run.complete({ 答え: 結果 });
コンソール .log(publicUrl); // → https://agentsproof.dev/r/abc123 レポート · 実行 #47 証明 my-coding-agent 2 分前 · 8 ステップ 87 /100 目標完了 92 ツール精度 78 ステップ効率 95 出力品質 88 安全性 97 異常なし レポートを開く OpenAI Anthropic LangChain CrewAI Vercel AI SDK LlamaIndex の仕組み ゼロから証明まで 5 つのステップで行います。
1 つのパッケージで構成不要。任意のノードまたはエッジ ランタイムで動作します。
各 LLM およびツール呼び出しの周囲に run.trace() を削除します。それが全体の統合です。
ルールは平易な英語で記述します。LLM はルールに対してすべての実行を自動的にチェックします。
追い越しランをライブテストスペックに変えます。成功基準を設定し、アサーションをトレースし、予想される動作をすべて今後の実行ごとに評価します。
1 つの SDK 呼び出しで、エージェントを通じて承認されたすべての Golden が実行されます。レポートには、全体的なスコアだけでなく、すべてのケースの基準ごとの合否が表示されます。
コピーインポート

{ AgentsProof } 'agentsproof' から;
const ap = new AgentsProof({ apiKey: process.env.AGENTSPROOF_API_KEY });
const run = ap.startRun({
projectSlug: 'my-agent' 、
入力: { クエリ }、
goal: 'Web を検索し、動作するコード ソリューションを返します' , // アンカー グレーディング
});
const result = await run.trace( 'llm_call' , 'gpt-4o' , () => llm(query));
const { publicUrl } = await run.complete({ 答え: 結果 });
コンソール .log(publicUrl); // → https://agentsproof.dev/r/abc123 得られるもの スコアだけでなく、エージェントの動作を知るために必要なすべてが得られます。
エージェントにとって「良い」とは何を意味するのかをわかりやすい英語で定義します。すべての実行は、ルールに基づいて自動的に採点されます。対象を絞った適用のために、特定の採点者を個々のゴールデンに固定します。
合格した実行を、成功基準、期待される動作、およびトレース アサーションを備えた実行可能なテスト仕様に変換します。 GoldenId を startRun() に直接渡して、オンデマンドで Golden に対して実行するか、プルーフ スイートですべてをバッチ処理します。
決定論的に実行される構造化チェック - LLM は必要ありません。エージェントが必要なステップをスキップしたり、ステップの予算を超えたりした場合、ケースは即座に失敗します。 LLM 審査員が見逃す可能性のある回帰をキャッチします。
AgentsProof は AI を使用してゴールデンからエッジケースのバリアントを生成し、手動でテストを作成することなくテスト スイートを拡張します。
1 つの SDK 呼び出しで、エージェントを通じて承認されたすべての Golden が実行されます。レポートには、基準ごとのゴールデン チェックが 5 軸スコアと並んで表示され、すべての合格と不合格が説明可能になります。
import { AgentsProof } を 'agentsproof' からコピーします。
const ap = new AgentsProof({ apiKey: process.env.AGENTSPROOF_API_KEY });
ap.runProofSuite({
projectSlug: 'my-agent' 、
suiteSlug: 'コア動作' 、
非同期ハンドラー(入力、ctx) {
// 入力は承認された Golden から取得されます
const run = ctx.startRun();
const result = await myAgent(input);
run.compl を待つ

ete({ 答え: 結果 });
}、
});
// → https://agentsproof.dev/p/suite-abc123 と代替手段の比較 セットアップが少ない。もっと信号を。
~ 部分的なサポート · 公開情報に基づく
無料で始められます。出荷の準備ができたらプロです。
公開および非公開の証拠レポート
仕事を証明できる船舶代理店。
カスタムグレーダー、ゴールデン、校正レポート - 数分で準備が整います。
無料利用枠: 200 実行/月。クレジットカードは必要ありません。

## Original Extract

Drop AgentsProof into your agent in 30 seconds. Get a public URL that proves it works — scored on goal completion, tool accuracy, efficiency, quality, and safety.

I built AgentsProof as a learning side-project to better understand agent evals and observability. It lets me trace runs, create eval cases, and see where an agent succeeds or fails. Nothing fancy but it does creates publicly shareable reports for your agent runs

AgentsProof — One decorator. A shareable eval report. AP AgentsProof Features Pricing Examples Login GitHub Now in beta Stop vibe-checking
AI agents. Prove they work.
A few SDK calls capture your agent's trace, grade it against your custom rules and approved test cases, and create a shareable proof report — in minutes.
You changed a prompt. Something broke. A user noticed first.
View example report $ npm install agentsproof TS PY agentsproof.dev/r/x8kp2m live agent.ts your code import { AgentsProof } from 'agentsproof' ;
const ap = new AgentsProof({ apiKey: process.env.AGENTSPROOF_API_KEY });
const run = ap.startRun({
projectSlug: 'my-agent' ,
input: { query },
goal: 'Search the web and return a working code solution' , // anchors grading
});
const result = await run.trace( 'llm_call' , 'gpt-4o' , () => llm(query));
const { publicUrl } = await run.complete({ answer: result });
console .log(publicUrl); // → https://agentsproof.dev/r/abc123 report · run #47 the proof my-coding-agent 2 minutes ago · 8 steps 87 /100 Goal completion 92 Tool accuracy 78 Step efficiency 95 Output quality 88 Safety 97 No anomalies Open report Drops into OpenAI Anthropic LangChain CrewAI Vercel AI SDK LlamaIndex How it works From zero to proof in five steps.
One package, zero config. Works in any Node or edge runtime.
Drop run.trace() around each LLM and tool call. That's the whole integration.
Write rules in plain English — the LLM checks every run against them automatically.
Turn a passing run into a live test spec. Set success criteria, trace assertions, and expected behavior — all evaluated on every future run.
One SDK call runs all approved Goldens through your agent. The report shows per-criterion pass/fail for every case — not just an overall score.
Copy import { AgentsProof } from 'agentsproof' ;
const ap = new AgentsProof({ apiKey: process.env.AGENTSPROOF_API_KEY });
const run = ap.startRun({
projectSlug: 'my-agent' ,
input: { query },
goal: 'Search the web and return a working code solution' , // anchors grading
});
const result = await run.trace( 'llm_call' , 'gpt-4o' , () => llm(query));
const { publicUrl } = await run.complete({ answer: result });
console .log(publicUrl); // → https://agentsproof.dev/r/abc123 What you get Everything you need to know your agent works — not just a score.
Define what 'good' means for your agent in plain English. Every run is automatically graded against your rules — pin specific graders to individual Goldens for targeted enforcement.
Turn any passing run into an executable test spec with success criteria, expected behavior, and trace assertions. Pass goldenId directly to startRun() to run against a Golden on demand — or batch them all in a proof suite.
Structured checks that run deterministically — no LLM needed. If your agent skips a required step or exceeds a step budget, the case fails immediately. Catches regressions the LLM judge can miss.
AgentsProof uses AI to generate edge-case variants from your Goldens, growing your test suite without writing tests by hand.
One SDK call runs all approved Goldens through your agent. The report shows per-criterion Golden checks alongside the 5-axis score — every pass and fail is explainable.
Copy import { AgentsProof } from 'agentsproof' ;
const ap = new AgentsProof({ apiKey: process.env.AGENTSPROOF_API_KEY });
await ap.runProofSuite({
projectSlug: 'my-agent' ,
suiteSlug: 'core-behaviors' ,
async handler(input, ctx) {
// input comes from the approved Golden
const run = ctx.startRun();
const result = await myAgent(input);
await run.complete({ answer: result });
},
});
// → https://agentsproof.dev/p/suite-abc123 vs. the alternatives Less setup. More signal.
~ partial support · based on publicly available information
Free to start. Pro when you're ready to ship.
Public + private proof reports
Ship agents you can prove work.
Custom graders, Goldens, and proof reports — ready in minutes.
Free tier: 200 runs/month. No credit card required.
