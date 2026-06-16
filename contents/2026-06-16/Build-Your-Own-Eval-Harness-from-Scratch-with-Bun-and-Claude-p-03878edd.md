---
source: "https://alexop.dev/posts/build-your-own-eval-harness-bun-claude-p/"
hn_url: "https://news.ycombinator.com/item?id=48560475"
title: "Build Your Own Eval Harness from Scratch with Bun and Claude -p"
article_title: "Build Your Own Eval Harness from Scratch with Bun and claude -p | alexop.dev"
author: "speckx"
captured_at: "2026-06-16T19:18:01Z"
capture_tool: "hn-digest"
hn_id: 48560475
score: 1
comments: 0
posted_at: "2026-06-16T19:15:37Z"
tags:
  - hacker-news
  - translated
---

# Build Your Own Eval Harness from Scratch with Bun and Claude -p

- HN: [48560475](https://news.ycombinator.com/item?id=48560475)
- Source: [alexop.dev](https://alexop.dev/posts/build-your-own-eval-harness-bun-claude-p/)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T19:15:37Z

## Translation

タイトル: Bun と Claude を使用して独自の評価ハーネスをゼロから構築する -p
記事のタイトル: Bun と claude -p を使用して独自の評価ハーネスをゼロから構築する |アレクソップ.dev
説明: 実践的なチュートリアル: AI エージェント用の実際に動作する評価ハーネスを 1 つの Bun ファイルで構築します。サンドボックスでエージェントを実行し、LLM 審査員によって決定論的に評価し、トライアル全体で投票し、結果に基づいて CI をゲートします。

記事本文:
Bun と claude -p を使用して独自の評価ハーネスを最初から構築する |アレクソップ.dev
次回のトーク: Claude Code による Web 開発の自動化
2026 年 7 月 1 日 — DWX Developer World、マンハイム
カンファレンス
コンテンツへスキップ
戻る Bun と claude -p を使用して独自の評価ハーネスをゼロから構築する
公開日: 2026 年 6 月 14 日 AI エージェントをテストするために、フレームワーク、SaaS ダッシュボード、または依存関係は必要ありません。実行する方法、評価する方法、そしてその両方のループが必要です。ここでは、単一の Bun ファイルで評価ハーネスを構築し、最初から最後まで、各行を説明します。
最終的には、サンドボックスを起動し、クロード CLI を通じてエージェントを駆動し、結果を 3 つの方法で評価する 1 つの evals.ts ファイルが完成します。
評価は、決定的ではないソフトウェアのテストです。単体テストでは「2 + 2 は 4 を返しますか?」と質問されますが、AI エージェントは質問するたびに異なる段落を提示するため、主張できる単一の値はありません。代わりに、評価は 1 つの観察可能な動作 (「まだ計画がない場合は、最初に計画することを推奨する」) を特定し、正確な言葉が異なるという事実を許容しながら、エージェントがそれを実行したかどうかを確認します。
このために人々はホストされたプラットフォームに手を伸ばします。その必要はありません。ダッシュボードの下にあるすべての評価ハーネスは、同じ 3 つの動作です。
エージェントを実行します。制御された環境でプロンプトを与え、その言動をすべてキャプチャします。
結果を採点します。可能な場合は文字列とファイルのアサーションを使用して出力をチェックし、できない場合は 2 番目の LLM を判断者として使用して、出力をチェックします。
ループしてレポートします。すべてのケースに対してこれを実行し、合格/不合格を集計し、何かが失敗した場合は CI がゲートできるようにゼロ以外で終了します。
Bun は、spawnSync と組み込みのファイルシステムを備えた高速な TypeScript ランタイムを提供します。claude CLI は、コマンド ラインと LLM から駆動できるエージェントを提供します。

私たちは裁判官として使うことができます。それがすべてです。最終的には、約 150 行からなる 1 つの evals.ts ファイルが作成され、これを bun evals.ts で実行し、一度に 1 つずつビルドします。
ハーネスの反対側のことも理解したい場合は、独自のコーディング エージェントを最初から構築することに関する関連記事を書きました。Building Your Owncoding Agent from Scratch TypeScript でクロードを利用した最小限のコーディング アシスタントを作成するための実践的なガイドです。基本的なチャット ループから始めて、約 400 行で完全に機能するコーディング エージェントが完成するまで、徐々にツールを追加していきます。 ai typescript ツール 2026 年 1 月 17 日 。
セットアップ: Bun とクロード CLI #
2 つの前提条件 (どちらもワンライナー):
#1. Bun — ハーネスを実行するランタイム
カール -fsSL https://bun.sh/install |バッシュ
# 2. クロード コード CLI — 私たちがテストしているエージェントと私たちの裁判官
npm install -g @anthropic-ai/claude-code
# サニティチェック: これによりモデルの応答が出力されるはずです
claude -p " 3 単語で挨拶します " --output-format json this.classList.remove('copied'), 3000);
" aria-label="コードをコピー">
私たちが頼りにする重要なフラグは --output-format json です。これにより、CLI は人間のテキストのストリームではなく、機械可読な 1 つの封筒を印刷します。フォルダーを作成し、空の evals.ts をドロップして、それに埋めてみましょう。
ステップ 1: コード # からエージェントを駆動する
1 つ目は、プロンプトでエージェントを実行し、その応答を返す関数です。 claude -p (「印刷」/非対話型モード) にシェルアウトし、それが出力する JSON エンベロープを解析します。このエンベロープには、 result の最終テキスト、 total_cost_usd のドルコスト、および is_error フラグが含まれます。
// evals.ts
「bun」から {spawnSync} をインポートします。
// `cwd` 内の `prompt` でエージェントを実行します。最終的な返答を返します。
function runAgent (プロンプト:文字列、cwd:文字列) {
const res = spawnSync ({
cmd : [
" クロード " 、 " -p

" 、プロンプト 、
" --output-format " , " json " , // 標準出力に 1 つの JSON エンベロープ
" --permission-mode " , " bypassPermissions " , // 実行中にプロンプトを表示しない
" --max-budget-usd " , " 0.50 " , // 実行ごとのハード安全キャップ
]、
CWD 、
標準出力: "パイプ" 、
標準エラー出力: "パイプ" 、
タイムアウト: 180_000 、
});
const エンベロープ = JSON 。 parse ( res . stdout . toString ()) ;
戻り値 {
テキスト: 封筒。結果?? 「」、
オーケー：レス。 exitCode === 0 && エンベロープ 。 is_error !== true 、
コスト : 数値 (エンベロープ . total_cost_usd ?? 0 ) 、
} ;
this.classList.remove('コピー済み'), 3000);
" aria-label="コードをコピー">
💡 spawnSync を使用する理由
Eval は遅い (それぞれが実際のモデル呼び出しです) が、巧妙な非同期パイプラインではなく、シンプルで順序付けられたものを望んでいます。同期スポーンにより、ハーネス全体が上から下まで読み取れる状態が保たれます。後で並列化することもできます。まず正しさ。
ステップ 2: 動作するサンドボックスを与える #
実際のリポジトリでエージェントを解放するのは悪い考えであり、実行が再現できなくなります。代わりに、すべてのケースで、動作に必要なファイル (フィクスチャ) がシードされた新しい使い捨て Git リポジトリが取得されます。実行が完了したら、検査または削除できます。
import { mkdtempSync , mkdirSync , writeFileSync } from "node:fs" ;
"node:os" から {tmpdir } をインポートします。
import { join , dirname } from "node:path" ;
// `files` をシードした使い捨て Git リポジトリを作成します。そのパスを返します。
function makeSandbox ( files : Record < string , string > ) {
const dir = mkdtempSync ( join ( tmpdir () , " eval- " )) ;
spawnSync ({ cmd : [ "git " , " init " , " -q " ] , cwd : dir }) ;
for ( const [ path , content ] of Object .entrys ( files )) {
const target = join ( dir , path ) ;
mkdirSync ( dirname ( target ) , { recursive : true }) ;
writeFileSync (ターゲット、コンテンツ) ;
}
ディレクトリを返します。
} ) {
const dir = mkdtempSync(join(tmpdir(), "eval-"));
spawnSync({ cmd

: ["git", "init", "-q"], cwd: dir });
for (const [パス, コンテンツ] of Object.entries(files)) {
const target = join(dir, path);
mkdirSync(dirname(target), { recursive: true });
writeFileSync(ターゲット, コンテンツ);
}
ディレクトリを返します。
}" オンクリック=
navigator.clipboard.writeText(this.dataset.code);
this.classList.add('コピーされました');
setTimeout(() => this.classList.remove('コピー済み'), 3000);
" aria-label="コードをコピー">
💡 備品を小さく保つ
「プランがすでに存在するかどうか」をテストする場合は、コードベースのクローンではなく、偽の docs/plans/checkout.md という 1 つのファイルが必要です。小さなフィクスチャは動作を分離し、高速に実行します。
ステップ 3: 安価で決定的なチェックを使用してグレードを付ける #
さて、採点です。動作をキャプチャする最も安価なツール、つまりプレーン文字列とファイルのチェックから始めます。これらは無料で即時に提供され、不安定になることはありません。これらが表現できないことについてのみ、LLM 裁判官に連絡してください。
「node:fs」から {existsSync} をインポートします。
const には = (干し草の山 : 文字列 , 針 : 文字列 ) =>
干し草の山。 toLowerCase() 。 (針 .toLowerCase()) ; が含まれます。
type チェック = {
required_substrings ?: 文字列 [] ; // 返信に含める必要があります
禁止されたサブストリング ?: 文字列 [] ; // 表示してはなりません
required_files ?: 文字列 [] ; // 実行後にサンドボックスに存在する必要があります
} ;
// チェックごとに [ラベル、合格] を返します。
function checkAssertions (チェック: チェック、応答: 文字列、ディレクトリ: 文字列) {
const out : [ 文字列 , ブール値 ][] = [] ;
for ( チェックの定数 . required_substrings ?? [])
出ます。 Push ([ ` contains " ${ s } " ` , has ( Reply , s )]) ;
for ( チェックの const .fordid_substrings ?? [])
出ます。 Push ([ ` は " ${ s } " ` , ! には ( Reply , s ) が含まれます]) ;
for ( チェックの const f . required_files ?? [])
出ます。 Push ([ ` 作成 ${ f } ` 、existsSync ( join ( dir , f ))]) ;
戻る ;
}
haystack.toLowerCase().i

ncludes(needle.toLowerCase());
type チェック = {
required_substrings?: 文字列[]; // 返信に含める必要があります
禁止された部分文字列?: 文字列[]; // 表示してはなりません
required_files?: 文字列[]; // 実行後にサンドボックスに存在する必要があります
};
// チェックごとに [ラベル、合格] を返します。
function checkAssertions(チェック: チェック、応答: 文字列、ディレクトリ: 文字列) {
const out: [文字列、ブール値][] = [];
for (checks.required_substrings の定数 ?? [])
out.push([`contains "${s}"`, has(reply, s)]);
for (checks.forbidden_substrings の定数 ?? [])
out.push([`「${s}」を除く`, !has(reply, s)]);
for (checks.required_files の const f ?? [])
out.push([`作成された ${f}`、existsSync(join(dir, f))]);
出て行きます。
}" オンクリック=
navigator.clipboard.writeText(this.dataset.code);
this.classList.add('コピーされました');
setTimeout(() => this.classList.remove('コピー済み'), 3000);
" aria-label="コードをコピー">
💡 要件を 2 つの部分に分割する
動作に 2 つの部分がある場合、たとえば「リポジトリのパスとその元の URL を記録する」場合は、1 つではなく 2 つのチェックを書きます。それ以外の場合は、半分正解で合格となります。 1 つの要件、1 つの主張。
ステップ 4: LLM 審査員によるあいまいな動作の採点 #
一部の動作にはキーワードがありません。 「最初の質問をする前にリポジトリを読みましたか?」 「トレードオフについては説明できましたか？」それらの場合は、回答を 2 番目のより安価なモデルに渡し、それぞれの期待値を評価してもらいます。これは強力ですが高価なので、慎重に使用してください。
プロンプトには 2 つの詳細が記載されています。裁判官に最初に理由を説明してから、厳密な JSON で答えるように依頼し、met の前に理由フィールドを置きます。これにより、評決を曖昧にするのではなく、正当化してから決定します。解析する前に推論を取り除きます。
// 安価なモデルに「応答」がそれぞれの期待を満たすかどうかを尋ねます。
関数審理 (応答: 文字列、期待: 文字列 []) : boolean [] {
もし(

期待に応えます。長さ === 0 ) [] を返します。
const 番号 = 期待値。 map (( e , i ) => ` ${ i + 1 } . ${ e } ` ) 。結合 ( "\n" ) ;
const プロンプト = ` AI エージェントの応答を期待値のリストに照らして評価しています。
単一の < Thinking>…</ Thinking> ブロック内の最初の理由。その後、
終了タグ、STRICT JSON のみを出力:
{"結果":[{"理由":"...","出会った":true}]}
期待ごとに 1 つのエントリを順番に入力します。
=== 返信 ===
${ 返信 }
=== 終了 ===
期待:
${ 番号付き } ` ;
const res = spawnSync ({
cmd : [
「クロード」、「-p」、プロンプト、
" --model " , " claude-haiku-4-5 " , // 小さい + 安いのでグレーディングには十分です
" --output-format " 、 " json " 、
" --permission-mode " 、 " bypassPermissions " 、
" --max-budget-usd " 、 " 0.10 " 、
]、
stdout: "パイプ"、stderr: "パイプ"、タイムアウト: 180_000、
});
// < Thinking> ブロックを削除して、JSON オブジェクトを取得します
const text = ( JSON . parse ( res . stdout . toString ()) . result ?? "" )
。 replace ( / <思考> [ \s\S ] *? < \/ 思考> /gi , "" ) ;
const json = JSON 。 parse ( text .slice ( text .indexOf ( " { " ) , text . lastIndexOf ( " } " ) + 1 )) ;
返品の期待。 map (( _ , i ) => json . results ?. [ i ] ?.met === true ) ;
} `${i + 1}。 ${e}`).join("\n");
const プロンプト = `AI エージェントの応答を期待リストに照らして評価しています。
単一の ... ブロック内の最初の理由。その後、
終了タグ、STRICT JSON のみを出力:
{"結果":[{"理由":"...","出会った":true}]}
期待ごとに 1 つのエントリを順番に入力します。
=== 返信 ===
${返信}
=== 終了 ===
期待:
${番号付き}`;
const res = spawnSync({
cmd: [
「クロード」、「-p」、プロンプト、
"--model", "claude-haiku-4-5", // 小さい + 安いものはグレーディングには十分です
"--出力形式"、"json"、
"--permission-mode"、"bypassPermissions"、
"--max-budget-usd"、"0.10"、
]、
stdout: "パイプ"、stderr: "p

イペ"、タイムアウト: 180_000、
});
// ブロックを剥がしてから、J を取得します
[切り捨てられた]
判定された eval は、判定者と同じくらい信頼できます。最初の数回は、生の判定出力をログに記録し、その推論を読みます。調書を読み間違えた裁判官は、ゲートが赤であるべきところを緑に反転させます。 2 つのサンプルを読むのは安い保険です。
ステップ 5: 複数回実行して # に投票します
エージェントを 1 回実行すると、成功は幸運かもしれませんが、失敗は不運かもしれません。解決策は、各ケースを数回実行し、多数決で決定することです。ボーナスとして、どのケースが不安定で、どこで試験が一致しないのかを知ることができます。これは、行動が後退するまであと 1 枚のコイン投げであるという早期警告となります。
// `fn` を N 回実行し、true を返した回数 + 大多数が true を返したかどうかを返します。
関数 vote (試行数 : 数値 , fn : () => ブール値 ) {
修正 = 0 にします。
for ( let i = 0 ; i < トライアル ; i ++ ) if ( fn ()) 正しい ++;
戻り値 {
正しい、
合格 : 正しい * 2 > 試行 , // 厳密な多数決
flaky : 正しい > 0 && 正しい < トライアル , // トライアルは不一致
} ;
} ブール値) {
修正 = 0 とします。
for (let i = 0; i 試行、 // 厳密多数決
不安定: 正しい > 0 && 正しい this.classList.remove('copied'), 3000);
" aria-label="コードをコピー">
これは評価に実際のお金がかかる唯一の場所であり、3 回の実行で 3 倍の費用がかかるため、これはプレリリース パスであり、これまでのパスではありません

[切り捨てられた]

## Original Extract

A hands-on tutorial: build a real, working eval harness for an AI agent in a single Bun file. Run the agent in a sandbox, grade it deterministically and with an LLM judge, vote across trials, and gate CI on the result.

Build Your Own Eval Harness from Scratch with Bun and claude -p | alexop.dev
Next Talk: Automating Web Development with Claude Code
July 1, 2026 — DWX Developer World, Mannheim
Conference
Skip to content alexop.dev Posts
Go back Build Your Own Eval Harness from Scratch with Bun and claude -p
Published: Jun 14, 2026 at You don’t need a framework, a SaaS dashboard, or a dependency to test an AI agent. You need a way to run it, a way to grade it, and a loop around both. Here we build an eval harness in a single Bun file, start to finish, every line explained.
By the end you’ll have one evals.ts file that spins up a sandbox, drives the agent through the claude CLI, and grades the result three ways.
An eval is a test for software that isn’t deterministic. A unit test asks “does 2 + 2 return 4?”, but an AI agent gives you a different paragraph every time you ask, so there’s no single value to assert against. An eval instead pins down one observable behavior (“when there’s no plan yet, it recommends planning first”) and checks whether the agent did it, while tolerating the fact that the exact words vary.
People reach for hosted platforms for this. You don’t have to. Every eval harness, underneath the dashboard, is the same three moves:
Run the agent. Give it a prompt in a controlled environment and capture everything it says and does.
Grade the result. Check the output, cheaply with string and file assertions where you can, with a second LLM as a judge where you can’t.
Loop and report. Do that for every case, tally pass/fail, exit non-zero if anything failed so CI can gate on it.
Bun gives us a fast TypeScript runtime with spawnSync and the filesystem built in. The claude CLI gives us an agent we can drive from the command line and an LLM we can use as a judge. That’s everything. You’ll end up with a single evals.ts file, roughly 150 lines, that you run with bun evals.ts , built one piece at a time.
If you want to understand the thing on the other end of the harness too, I wrote a companion piece on building your own coding agent from scratch Building Your Own Coding Agent from Scratch A practical guide to creating a minimal Claude-powered coding assistant in TypeScript. Start with a basic chat loop and progressively add tools until you have a fully functional coding agent in about 400 lines. ai typescript tooling Jan 17, 2026 .
Setup: Bun and the claude CLI #
Two prerequisites, both one-liners:
# 1. Bun — the runtime that runs our harness
curl -fsSL https://bun.sh/install | bash
# 2. The Claude Code CLI — the agent we're testing, and our judge
npm install -g @anthropic-ai/claude-code
# sanity check: this should print a model's reply
claude -p " say hi in three words " --output-format json this.classList.remove('copied'), 3000);
" aria-label="Copy code">
The key flag we’ll lean on is --output-format json , which makes the CLI print one machine-readable envelope instead of a stream of human text. Make a folder, drop in an empty evals.ts , and let’s fill it.
Step 1: drive the agent from code #
First, a function that runs the agent on a prompt and hands back its reply. We shell out to claude -p (the “print” / non-interactive mode) and parse the JSON envelope it prints. That envelope carries the final text in result , the dollar cost in total_cost_usd , and an is_error flag.
// evals.ts
import { spawnSync } from " bun " ;
// Run the agent on `prompt` inside `cwd`; return its final reply.
function runAgent ( prompt : string , cwd : string ) {
const res = spawnSync ({
cmd : [
" claude " , " -p " , prompt ,
" --output-format " , " json " , // one JSON envelope on stdout
" --permission-mode " , " bypassPermissions " , // don't prompt us mid-run
" --max-budget-usd " , " 0.50 " , // hard safety cap per run
] ,
cwd ,
stdout : " pipe " ,
stderr : " pipe " ,
timeout : 180_000 ,
}) ;
const envelope = JSON . parse ( res . stdout . toString ()) ;
return {
text : envelope . result ?? "" ,
ok : res . exitCode === 0 && envelope . is_error !== true ,
cost : Number ( envelope . total_cost_usd ?? 0 ) ,
} ;
} this.classList.remove('copied'), 3000);
" aria-label="Copy code">
💡 Why spawnSync
Evals are slow (each one is a real model call) but we want them simple and ordered, not a clever async pipeline. Synchronous spawning keeps the whole harness readable top-to-bottom. You can parallelize later; correctness first.
Step 2: give it a sandbox to act in #
Letting an agent loose in your real repo is a bad idea, and it makes runs non-repeatable. Instead, every case gets a fresh throwaway git repo seeded with the files that behavior needs, a fixture. When the run is done, you can inspect or delete it.
import { mkdtempSync , mkdirSync , writeFileSync } from " node:fs " ;
import { tmpdir } from " node:os " ;
import { join , dirname } from " node:path " ;
// Make a throwaway git repo seeded with `files`; return its path.
function makeSandbox ( files : Record < string , string > ) {
const dir = mkdtempSync ( join ( tmpdir () , " eval- " )) ;
spawnSync ({ cmd : [ " git " , " init " , " -q " ] , cwd : dir }) ;
for ( const [ path , content ] of Object . entries ( files )) {
const target = join ( dir , path ) ;
mkdirSync ( dirname ( target ) , { recursive : true }) ;
writeFileSync ( target , content ) ;
}
return dir ;
} ) {
const dir = mkdtempSync(join(tmpdir(), "eval-"));
spawnSync({ cmd: ["git", "init", "-q"], cwd: dir });
for (const [path, content] of Object.entries(files)) {
const target = join(dir, path);
mkdirSync(dirname(target), { recursive: true });
writeFileSync(target, content);
}
return dir;
}" onclick="
navigator.clipboard.writeText(this.dataset.code);
this.classList.add('copied');
setTimeout(() => this.classList.remove('copied'), 3000);
" aria-label="Copy code">
💡 Keep fixtures tiny
If you’re testing “does it notice the plan already exists,” you need one file, a fake docs/plans/checkout.md , not a clone of your codebase. Small fixtures isolate the behavior and run fast.
Step 3: grade with cheap, deterministic checks #
Now the grading. Start with the cheapest tool that captures the behavior: plain string and file checks. They’re free, instant, and never flaky. Reach for the LLM judge only for what these can’t express.
import { existsSync } from " node:fs " ;
const has = ( haystack : string , needle : string ) =>
haystack . toLowerCase () . includes ( needle . toLowerCase ()) ;
type Checks = {
required_substrings ?: string [] ; // must appear in the reply
forbidden_substrings ?: string [] ; // must NOT appear
required_files ?: string [] ; // must exist in the sandbox after the run
} ;
// Returns [label, passed] for each check.
function checkAssertions ( checks : Checks , reply : string , dir : string ) {
const out : [ string , boolean ][] = [] ;
for ( const s of checks . required_substrings ?? [])
out . push ([ ` contains " ${ s } " ` , has ( reply , s )]) ;
for ( const s of checks . forbidden_substrings ?? [])
out . push ([ ` excludes " ${ s } " ` , ! has ( reply , s )]) ;
for ( const f of checks . required_files ?? [])
out . push ([ ` created ${ f } ` , existsSync ( join ( dir , f ))]) ;
return out ;
}
haystack.toLowerCase().includes(needle.toLowerCase());
type Checks = {
required_substrings?: string[]; // must appear in the reply
forbidden_substrings?: string[]; // must NOT appear
required_files?: string[]; // must exist in the sandbox after the run
};
// Returns [label, passed] for each check.
function checkAssertions(checks: Checks, reply: string, dir: string) {
const out: [string, boolean][] = [];
for (const s of checks.required_substrings ?? [])
out.push([`contains "${s}"`, has(reply, s)]);
for (const s of checks.forbidden_substrings ?? [])
out.push([`excludes "${s}"`, !has(reply, s)]);
for (const f of checks.required_files ?? [])
out.push([`created ${f}`, existsSync(join(dir, f))]);
return out;
}" onclick="
navigator.clipboard.writeText(this.dataset.code);
this.classList.add('copied');
setTimeout(() => this.classList.remove('copied'), 3000);
" aria-label="Copy code">
💡 Split two-part requirements
If a behavior has two halves, say, “records the repo’s path and its origin URL,” write two checks, not one. Otherwise a half-right answer passes. One requirement, one assertion.
Step 4: grade fuzzy behavior with an LLM judge #
Some behaviors have no keyword. “Did it read the repo before asking its first question?” “Did it explain the trade-off?” For those, you hand the reply to a second, cheaper model and ask it to grade each expectation. It’s the powerful but pricey rung, so use it sparingly.
Two details earn their keep in the prompt. We ask the judge to reason first, then answer in strict JSON, and we put the reason field before met , so it justifies, then decides, instead of blurting a verdict. We strip the reasoning before parsing.
// Ask a cheap model whether `reply` meets each expectation.
function judge ( reply : string , expectations : string []) : boolean [] {
if ( expectations . length === 0 ) return [] ;
const numbered = expectations . map (( e , i ) => ` ${ i + 1 } . ${ e } ` ) . join ( "\n" ) ;
const prompt = ` You are grading an AI agent's reply against a list of expectations.
First reason inside a single <thinking>…</thinking> block. Then, after the
closing tag, output STRICT JSON only:
{"results":[{"reason":"...","met":true}]}
one entry per expectation, in order.
=== REPLY ===
${ reply }
=== END ===
Expectations:
${ numbered } ` ;
const res = spawnSync ({
cmd : [
" claude " , " -p " , prompt ,
" --model " , " claude-haiku-4-5 " , // small + cheap is plenty for grading
" --output-format " , " json " ,
" --permission-mode " , " bypassPermissions " ,
" --max-budget-usd " , " 0.10 " ,
] ,
stdout : " pipe " , stderr : " pipe " , timeout : 180_000 ,
}) ;
// strip the <thinking> block, then grab the JSON object
const text = ( JSON . parse ( res . stdout . toString ()) . result ?? "" )
. replace ( / <thinking> [ \s\S ] *? < \/ thinking> /gi , "" ) ;
const json = JSON . parse ( text . slice ( text . indexOf ( " { " ) , text . lastIndexOf ( " } " ) + 1 )) ;
return expectations . map (( _ , i ) => json . results ?. [ i ] ?. met === true ) ;
} `${i + 1}. ${e}`).join("\n");
const prompt = `You are grading an AI agent's reply against a list of expectations.
First reason inside a single … block. Then, after the
closing tag, output STRICT JSON only:
{"results":[{"reason":"...","met":true}]}
one entry per expectation, in order.
=== REPLY ===
${reply}
=== END ===
Expectations:
${numbered}`;
const res = spawnSync({
cmd: [
"claude", "-p", prompt,
"--model", "claude-haiku-4-5", // small + cheap is plenty for grading
"--output-format", "json",
"--permission-mode", "bypassPermissions",
"--max-budget-usd", "0.10",
],
stdout: "pipe", stderr: "pipe", timeout: 180_000,
});
// strip the block, then grab the J
[truncated]
A judged eval is only as trustworthy as the judge. The first few times, log the raw judge output and read its reasoning. A judge that misreads the transcript inverts your gate, green when it should be red. Reading two samples is cheap insurance.
Step 5: run it more than once and vote #
Run an agent once and a pass might be luck, a fail might be a bad roll. The fix is to run each case a few times and decide by majority. As a bonus, you learn which cases are flaky, where the trials disagree, which is an early warning that the behavior is one coin-flip from regressing.
// Run `fn` N times, return how many returned true + whether the majority did.
function vote ( trials : number , fn : () => boolean ) {
let correct = 0 ;
for ( let i = 0 ; i < trials ; i ++ ) if ( fn ()) correct ++;
return {
correct ,
passed : correct * 2 > trials , // strict majority
flaky : correct > 0 && correct < trials , // trials disagreed
} ;
} boolean) {
let correct = 0;
for (let i = 0; i trials, // strict majority
flaky: correct > 0 && correct this.classList.remove('copied'), 3000);
" aria-label="Copy code">
This is the one place evals cost real money, three runs is three times the spend, so it’s a pre-release pass, not an ever

[truncated]
