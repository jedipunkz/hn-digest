---
source: "https://www.mager.co/blog/2026-03-14-claude-agent-sdk-tui/"
hn_url: "https://news.ycombinator.com/item?id=48344693"
title: "Claude Agent SDK: Build Your Own AI Terminal in 10 Minutes"
article_title: "Claude Agent SDK: Build Your Own AI Terminal in 10 Minutes"
author: "ankitg12"
captured_at: "2026-06-01T00:33:07Z"
capture_tool: "hn-digest"
hn_id: 48344693
score: 2
comments: 0
posted_at: "2026-05-31T10:52:42Z"
tags:
  - hacker-news
  - translated
---

# Claude Agent SDK: Build Your Own AI Terminal in 10 Minutes

- HN: [48344693](https://news.ycombinator.com/item?id=48344693)
- Source: [www.mager.co](https://www.mager.co/blog/2026-03-14-claude-agent-sdk-tui/)
- Score: 2
- Comments: 0
- Posted: 2026-05-31T10:52:42Z

## Translation

タイトル: Claude Agent SDK: 10 分で独自の AI ターミナルを構築
説明: Claude Agent SDK は、Claude Code を駆動するのと同じエンジンを提供し、完全にプログラム可能です。ここ

記事本文:
Claude Agent SDK: 10 分で独自の AI ターミナルを構築
▸ mager .co Home Tech Life Food About Follow @mager on X go to mager の GitHub ページ 💻 <mager.co /> Claude Agent SDK: 10 分で独自の AI ターミナルを構築 ⬆️ ~/mager.co/tech $ catarticle.md ← mager.co / 2026 年 3 月 14 日 Claude Agent SDK: 10 分で独自の AI ターミナルを構築分
Claude Agent SDK は、Claude Code を駆動するのと同じエンジンを提供し、完全にプログラム可能です。これを使用してカスタム TUI を 10 分で構築する方法を次に示します。
ターミナルからクロード コードを使用しました。さあ、あなた自身のものを構築してください。
これが Claude Agent SDK の売り文句です。Claude Code を動かすエンジンと同じですが、プログラム可能です。ファイルの読み取り、bash の実行、Web 検索、コードの編集といった完全なエージェント ループが、制御する for await ループにラップされます。
誰もが抱く疑問です。Claude API を直接呼び出す代わりに、なぜこれを使用するのでしょうか?
答え: ツール ループを自分で実装する必要はありません。
そして、その最も魅力的な使用例は何でしょうか?独自の TUI を構築します。
デモ リポジトリ: github.com/mager/claude-tui-demo — クローンを作成して、手順に従ってください。
SDK と API: 実際の違いは何ですか
標準の Anthropic クライアント SDK を使用して、ツールの実行を自分で実装します。
// このループを作成します。毎回。
let 応答 = client.messages を待ちます。 create ({ ... params });
while (response.stop_reason === "tool_use" ) {
const result = yourToolExecutor (response.tool_use);
応答 = client.messages を待ちます。 create ({ tools_result: result, ... params });
}
エージェント SDK を使用する場合:
"@anthropic-ai/claude-agent-sdk" から {クエリ} をインポートします。
for await ( const クエリのメッセージ ({
プロンプト: "auth.ts のバグを見つけて修正してください" ,
オプション: { allowedTools: [ "読み取り" 、 "編集" 、 "Bash" ] }
})){
コンソール。ログ（メッセージ）;
}
クロードはファイルを読み、

バグがあるので編集します。出力をストリーミングします。ツールループ、エグゼキューター、ボイラープレートはありません。
無料で入手できる組み込みツール:
これが Claude Code のプログラム可能なツールセット全体です。
Claude Code CLI は一般的な使用に最適です。しかし、特定のドメイン、つまりカスタム規則を備えたコードベース、特殊な手順を備えたワークフロー、さまざまな権限が必要なチームを持った瞬間から、独自のインターフェイスが必要になります。
チームが関心のあるコンテキスト (アーキテクチャ ドキュメント、スタイル ガイド) を事前にロードします。
ツールをロックダウンする - 読み取り専用のレビュー担当者が誤って製品を編集できないようにする
Surface ドメイン固有のショートカット - 1 つのキーストロークでテスト スイート全体を実行
出力を CI/CD またはロギング インフラストラクチャにパイプします。
フックを追加 - すべてのファイル変更を監査し、破壊的な操作をブロックし、承認を要求します
クロード・コードを置き換えるのではありません。ワークフローに正確に適合するバージョンの Claude Code を構築しています。
git clone https://github.com/mager/claude-tui-demo.git
cd クロード・ツイ・デモ
npmインストール
エクスポート ANTHROPIC_API_KEY = あなたのキー
API クレジット: クレジット付きの Anthropic API キーが必要です。 platform.claude.com/settings/billing で補充してください。
ターミナルUIにはInkを使用しています。 React を知っている人は、Ink についてもすでに知っています。同じコンポーネント モデル、同じフック ( useState 、 useEffect )、同じ JSX です。ただし、DOM にレンダリングする代わりに、ターミナルにレンダリングします。 Box はあなたの div です。テキストはあなたの範囲です。フレックスボックスと色は期待どおりに機能します。これは、TypeScript でインタラクティブなターミナル UI を構築する最もクリーンな方法です。
ランナーに関する注意: デモでは ts-node の代わりに tsx を使用します。 tsx は設定不要です。ローダー フラグなしで、すぐに .tsx 、JSX、および ESM を処理します。また、 package.json に "type": "module" が含まれていることを確認してください。Ink のレイアウト エンジン ( Yoga-layout ) はトップレベルの await を使用しており、これには ESM モードが必要です。不可解なエラーに遭遇しますw

それじゃ。
// エージェント.ts
"@anthropic-ai/claude-agent-sdk" から {クエリ} をインポートします。
非同期関数のエクスポート* runAgent (プロンプト:文字列) {
for await ( const クエリのメッセージ ({
プロンプト、
オプション: {
allowedTools: [ "Read" 、 "Glob" 、 "Grep" 、 "Bash" ]、
}、
})){
メッセージを譲ります。
}
}
非同期関数*とは何ですか? * はこれをジェネレーター関数にします。すべてを計算して一度に返すのではなく、 yield を介して一度に 1 つの値を渡し、それぞれの値の間に一時停止します。 async は、内部で待機することもできることを意味します。コンシューマ側では、 await は一度に 1 つのメッセージを非同期ストリームに処理します。これは、すべてが完了した後ではなく、ツールの呼び出しと応答が発生時に UI にストリーミングされる仕組みです。
Ink は、ターミナル用の React スタイルのコンポーネントを提供します。ボックスはレイアウトを処理し、テキストは色とスタイルをサポートして出力を処理します。
// App.tsx
import React, { useState, useEffect } from "react" ;
import { Box, Text, useInput, useApp } from "ink" ;
"./agent.js" から { runAgent } をインポートします。
type LogLine = { タイプ : "ユーザー" | "エージェント" | 「ツール」 | "結果" ;テキスト : 文字列 };
関数 formatToolCall (ブロック:任意):文字列 {
`⚙ ${ ブロックを返します。名前 }(${ JSON . stringify ( block . input ) . slide ( 0 , 60 ) })` ;
}
function handleAssistantMessage (msg : any , setLines : React . Dispatch < React . SetStateAction < LogLine []>>) {
for ( msg.message.content の const ブロック) {
if (block.type === "text" ) {
setLines (( prev ) => [ ... prev, { type: "agent" , text: block.text }]);
}
if (block.type === "tool_use" ) {
setLines (( prev ) => [ ... prev, { type: "tool" , text: formatToolCall (block) }]);
}
}
}
エクスポート関数 App ({ プロンプト } : { プロンプト : 文字列 }) {
const [ 行 , setLines ] = useState < LogLine []>([]);
const [ 完了 , setDone ] = useState ( false );
const { exit } = useApp();
useEffect (() => {
セットリン

es ([{ タイプ: "ユーザー" , テキスト: `> ${ プロンプト }` }]);
( async () => {
for await ( runAgent (プロンプト) の const msg) {
if (msg.type === "assistant" ) handleAssistantMessage (msg, setLines);
if (msg.type === "結果" ) {
setLines (( prev ) => [ ... prev, { type: "result" , text: `✓ ${ msg . result }` }]);
setDone ( true );
}
}
})();
}、[]);
useInput (( _ , key ) => {
if (key.escape || (key.ctrl && _.toLowerCase () === "c" )) exit ();
});
const Colors : Record < LogLine [ "type" ], string > = {
ユーザー: "シアン" 、
エージェント: "白" 、
ツール: "黄色" 、
結果: "緑" 、
};
戻る (
< ボックス flexDirection = "列" パディング = { 1 } >
< ボックスの余白下 = { 1 } >
< 太字の色 = "シアン" > ◆ 私の AI 端末 </ 本文 >
< Text color = "gray" > (終了するにはesc) </ Text >
</ボックス>
{ 行 .マップ (( line , i ) => (
< テキスト キー = {i} カラー = {colors[line.type]}>{line.text}</Text>
))}
{!完了 && <Text color="gray">▸考え中...</Text>}
</ボックス>
);
}
メッセージ タイプ: SDK はいくつかのメッセージ タイプをストリーミングします — アシスタント (クロードの応答)、レス
[切り捨てられた]
// インデックス.tsx
"react" から React をインポートします。
import { render } from "ink" ;
"./App.js" から { App } をインポートします。
const プロンプト = process.argv。スライス ( 2 )。結合 ( " " ) || 「このディレクトリにはどのようなファイルが入っていますか?」 ;
render (< アプリプロンプト ={ プロンプト } />);
process.argv.slice(2) は、ノードとスクリプト パスの後のすべて、つまり実際に型付けされた引数を取得します。 .join(" ") は複数単語のプロンプトを再構築します。七行。それがエントリーポイント全体です。
npm start 「このディレクトリにはどのようなファイルがありますか?」
Claude のツール呼び出しストリームがリアルタイムで表示されます。⚙ Bash({"command":"ls"}) は黄色、応答は白色、✓ 完了は緑色で表示されます。これは、約 80 行で構成される AI TUI です。
レベルアップ: 永遠のループ (REPL モード)
シングルプロンプト TUI は、単発のタスクに最適です。でもクロードが欲しいならどうする？

ただ...応答し続けるには？実際のクロード コードのエクスペリエンスと同様に、プロンプトを入力し、応答を受け取り、別のプロンプトを入力しますか?
これは REPL であり、while (true) ループです。
// repl.ts
"@anthropic-ai/claude-agent-sdk" から {クエリ} をインポートします。
import * as readline from "readline" ;
セッション ID : 文字列 |未定義 ;
const rl = 読み取りライン。 createInterface ({ 入力: process.stdin、出力: process.stdout });
const ask = (プロンプト : string ) => new Promise < string >((solve ) => rl.question (prompt,solve));
非同期関数 runTurn ( userPrompt : string ) {
for await ( クエリの const msg ({
プロンプト: ユーザープロンプト、
オプション: { allowedTools: [ "Read" 、 "Glob" 、 "Grep" 、 "Bash" ]、 再開: sessionId }、
})){
if (msg.type === "system" && msg.subtype === "init" ) sessionId = msg.session_id;
if (msg.type === "アシスタント" ) {
for ( msg.message.content の const ブロック) {
if (block.type === "text" ) process.stdout。 write ( ` \n 🤖 ${ block . text } \n ` );
if (block.type === "tool_use" ) {
プロセス.stdout。 write ( `⚙ ${ ブロック . 名前 }(${ JSON . stringify ( ブロック . input ). スライス ( 0 , 80 ) }) \n ` );
}
}
}
}
}
コンソール。 log ( "◆ Claude REPL — プロンプトを入力します。終了するには ctrl+c \n " );
while ( true ) {
const input = await ask ( " \n > " );
if (!input.trim()) 続行 ;
await runTurn (input.trim());
}
npm run repl で実行します。何でも入力してください。クロードは答える。もう一度入力します。以前の内容がまだ残っています。これが履歴書です: sessionId がその役割を果たします。
本当の力はフック、つまりエージェントのライフサイクルの重要なポイントで起動されるコールバックです。監査ログ、承認ゲート、またはカスタム UI フィードバックを追加する方法は次のとおりです。
// Agent.ts (フック付き)
"@anthropic-ai/claude-agent-sdk" から {クエリ} をインポートします。
import { appendFile } から "fs/promises" ;
const AuditHook = async (入力: 任意) => {
constツール =

入力ツール名 ?? "未知" ;
const args = JSON 。 stringify (input.tool_input ?? {})。スライス ( 0 , 100 );
await appendFile ( "./audit.log" , `${ new Date ().toISOString () } ${ tools } ${ args } \n ` );
戻る {};
};
for await ( const クエリのメッセージ ({
プロンプト: "認証モジュールをリファクタリングします" ,
オプション: {
allowedTools: [ "読み取り" 、 "編集" 、 "Bash" ]、
フック: {
PostToolUse: [{ マッチャー: ".*" 、フック: [auditHook] }]、
}、
}、
})){
// TUI にレンダリングします
}
すべてのツール呼び出しは、タイムスタンプとともに Audit.log に記録されます。 matcher: ".*" はすべてをキャッチします。ミューテーションのみを考慮する場合は、"Edit|Write" に絞り込みます。
知っておく価値のあるその他のフック: PreToolUse を使用して操作を実行前にブロックし、Stop を使用してエージェントの終了を検出し、UserPromptSubmit を使用して入力を前処理または検証します。
エージェントは、複数の query() 呼び出しにわたるコンテキストを記憶します。最初の実行からセッション ID を取得し、それを次の実行に渡します。
セッション ID : 文字列 |未定義 ;
// 最初のターン — クロードがファイルを読み取ります
for await (クエリの const msg ({ プロンプト: "認証モジュールの読み取り" })) {
if (msg.type === "system" && msg.subtype === "init" ) {
セッションID = msg.session_id;
}
}
// 第 2 ターン — ツール呼び出しはゼロ、クロードはすでに知っています
for await ( クエリの const msg ({
プロンプト: "それを呼び出すものをすべて見つけます" ,
オプション: { 再開: sessionId }、
})){
// ...
}
お金の詳細: 2 番目のターンではツール呼び出しは発生しません。クロードはすでにコンテキスト内にファイルを持っています。再読み込みや余分な API 呼び出しは必要ありません。ただ答えるだけです。
デモを実行します: npm run session 。
インクは素早い TUI に最適です。ただし、テーブル、コマンド パレット、分割ペイン、グラフ、モーダルなど、より豊富なウィジェットが必要な場合は、Rezi がアップグレード パスです。 TypeScript と Node.js はそのままですが、C エンジンと 50 以上の組み込みウィジェットによるネイティブ サポートのレンダリングが可能です。
Ink が React に似ているところ (フック、JSX、コンポーネント tr)

ee)、Rezi は状態駆動型です。状態 → UI をマップするビュー関数を定義し、app.update() を呼び出して状態を変更します。 Bubble Tea の Elm Architecture と同じメンタル モデルですが、TypeScript を使用しています。
npm install @rezi-ui/core @rezi-ui/node
Rezi で再構築された同じ Claude TUI は次のとおりです。
// rezi/rezi-app.ts
"@rezi-ui/core" から { ui } をインポートします。
import { createNodeApp } から "@rezi-ui/node" ;
"@anthropic-ai/claude-agent-sdk" から {クエリ} をインポートします。
タイプ LineKind = "ユーザー" | "エージェント" | 「ツール」 | "結果" ;
type LogLine = { 種類 : LineKind ;テキスト : 文字列 };
type State = { 行 : LogLine [];完了: ブール値 };
const プロンプト = process.argv。スライス ( 2 )。結合 ( " " ) || 「このディレクトリにはどのようなファイルが入っていますか?」 ;
const app = createNodeApp < 状態 >({
初期状態: { 行: [{ 種類: "ユーザー" 、テキスト: `> ${ プロンプト }` }]、完了: false }、
});
const kindVariant : レコード < LineKind , 文字列 > = {
ユーザー: "情報" 、
エージェント: 「本体」 、
ツール: "警告" 、
結果: "成功" 、
};
アプリ。 view (( 状態 ) =>
うーい。ページ ({
p:1、
ギャップ: 1 、
ヘッダー: ui.ヘッダー ({ タイトル: "◆ 私の AI ターミナル" , サブタイトル: "Q を終了する" }),
本体：うい。パネル ( "出力" , [
... 状態行。マップ (( 行 , i ) =>
うーい。 text (line.text, { キー: String (i), バリアント: kindVariant[line.kind] as any })
）、
... ( ! state.done ? [ui. スピナー ({ ラベル: "思考…" , キー: "スピナー" })]

[切り捨てられた]

## Original Extract

The Claude Agent SDK gives you the same engine that powers Claude Code, fully programmable. Here

Claude Agent SDK: Build Your Own AI Terminal in 10 Minutes
▸ mager .co Home Tech Life Food About Follow @mager on X Go to mager's GitHub page 💻 <mager.co /> Claude Agent SDK: Build Your Own AI Terminal in 10 Minutes ⬆️ ~/mager.co/tech $ cat article.md ← mager.co / March 14, 2026 Claude Agent SDK: Build Your Own AI Terminal in 10 Minutes
The Claude Agent SDK gives you the same engine that powers Claude Code, fully programmable. Here's how to build a custom TUI with it in 10 minutes.
You've used Claude Code from the terminal. Now build your own.
That's the pitch for the Claude Agent SDK — same engine that powers Claude Code, but programmable. You get the full agent loop — file reading, bash execution, web search, code editing — wrapped in a for await loop you control.
The question everyone asks: why would I use this instead of just calling the Claude API directly?
The answer: you don't have to implement the tool loop yourself.
And the most compelling use case for that? Building your own TUI.
Demo repo: github.com/mager/claude-tui-demo — clone it and follow along.
The SDK vs. The API: What's the Actual Difference
With the standard Anthropic client SDK, you implement tool execution yourself:
// You write this loop. Every time.
let response = await client.messages. create ({ ... params });
while (response.stop_reason === "tool_use" ) {
const result = yourToolExecutor (response.tool_use);
response = await client.messages. create ({ tool_result: result, ... params });
}
With the Agent SDK:
import { query } from "@anthropic-ai/claude-agent-sdk" ;
for await ( const message of query ({
prompt: "Find and fix the bug in auth.ts" ,
options: { allowedTools: [ "Read" , "Edit" , "Bash" ] }
})) {
console. log (message);
}
Claude reads the file, finds the bug, edits it. You stream the output. No tool loop, no executor, no boilerplate.
Built-in tools you get for free:
That's Claude Code's entire toolset, programmable.
The Claude Code CLI is great for general use. But the moment you have a specific domain — a codebase with custom conventions, a workflow with specialized steps, a team with different permission needs — you want your own interface.
Pre-load context your team cares about (architecture docs, style guides)
Lock down tools — a read-only reviewer can't accidentally edit prod
Surface domain-specific shortcuts — one keystroke to run your whole test suite
Pipe output into your CI/CD or logging infrastructure
Add hooks — audit every file change, block destructive operations, require approval
You're not replacing Claude Code. You're building the version of Claude Code that fits your workflow exactly.
git clone https://github.com/mager/claude-tui-demo.git
cd claude-tui-demo
npm install
export ANTHROPIC_API_KEY = your-key
API credits: You'll need an Anthropic API key with credits. Top up at platform.claude.com/settings/billing .
I'm using Ink for the terminal UI. If you know React, you already know Ink — same component model, same hooks ( useState , useEffect ), same JSX. But instead of rendering to the DOM, it renders to your terminal. Box is your div . Text is your span . Flexbox and colors work exactly as you'd expect. It's the cleanest way to build interactive terminal UIs in TypeScript.
Note on the runner: The demo uses tsx instead of ts-node . tsx is zero-config — it handles .tsx , JSX, and ESM out of the box without loader flags. Also make sure "type": "module" is in your package.json — Ink's layout engine ( yoga-layout ) uses top-level await , which requires ESM mode. You'll hit a cryptic error without it.
// agent.ts
import { query } from "@anthropic-ai/claude-agent-sdk" ;
export async function* runAgent ( prompt : string ) {
for await ( const message of query ({
prompt,
options: {
allowedTools: [ "Read" , "Glob" , "Grep" , "Bash" ],
},
})) {
yield message;
}
}
What's async function* ? The * makes this a generator function — instead of computing everything and returning at once, it hands you one value at a time via yield , pausing between each. async means it can also await internally. On the consumer side, for await handles the async stream one message at a time. This is how the tool calls and responses stream to your UI as they happen, not after everything finishes.
Ink gives us React-style components for the terminal. Box handles layout, Text handles output with color and style support.
// App.tsx
import React, { useState, useEffect } from "react" ;
import { Box, Text, useInput, useApp } from "ink" ;
import { runAgent } from "./agent.js" ;
type LogLine = { type : "user" | "agent" | "tool" | "result" ; text : string };
function formatToolCall ( block : any ) : string {
return `⚙ ${ block . name }(${ JSON . stringify ( block . input ). slice ( 0 , 60 ) })` ;
}
function handleAssistantMessage ( msg : any , setLines : React . Dispatch < React . SetStateAction < LogLine []>>) {
for ( const block of msg.message.content) {
if (block.type === "text" ) {
setLines (( prev ) => [ ... prev, { type: "agent" , text: block.text }]);
}
if (block.type === "tool_use" ) {
setLines (( prev ) => [ ... prev, { type: "tool" , text: formatToolCall (block) }]);
}
}
}
export function App ({ prompt } : { prompt : string }) {
const [ lines , setLines ] = useState < LogLine []>([]);
const [ done , setDone ] = useState ( false );
const { exit } = useApp ();
useEffect (() => {
setLines ([{ type: "user" , text: `> ${ prompt }` }]);
( async () => {
for await ( const msg of runAgent (prompt)) {
if (msg.type === "assistant" ) handleAssistantMessage (msg, setLines);
if (msg.type === "result" ) {
setLines (( prev ) => [ ... prev, { type: "result" , text: `✓ ${ msg . result }` }]);
setDone ( true );
}
}
})();
}, []);
useInput (( _ , key ) => {
if (key.escape || (key.ctrl && _. toLowerCase () === "c" )) exit ();
});
const colors : Record < LogLine [ "type" ], string > = {
user: "cyan" ,
agent: "white" ,
tool: "yellow" ,
result: "green" ,
};
return (
< Box flexDirection = "column" padding = { 1 } >
< Box marginBottom = { 1 } >
< Text bold color = "cyan" > ◆ My AI Terminal </ Text >
< Text color = "gray" > (esc to quit) </ Text >
</ Box >
{ lines . map (( line , i ) => (
< Text key = {i} color = {colors[line.type]}>{line.text}</Text>
))}
{!done && <Text color="gray">▸ thinking...</Text>}
</Box>
);
}
Message types: The SDK streams several message types — assistant (Claude's response), resu
[truncated]
// index.tsx
import React from "react" ;
import { render } from "ink" ;
import { App } from "./App.js" ;
const prompt = process.argv. slice ( 2 ). join ( " " ) || "What files are in this directory?" ;
render (< App prompt ={ prompt } />);
process.argv.slice(2) grabs everything after node and the script path — your actual typed arguments. .join(" ") reassembles multi-word prompts. Seven lines. That's the whole entry point.
npm start "What files are in this directory?"
You'll see Claude's tool calls stream in real-time — ⚙ Bash({"command":"ls"}) in yellow, the response in white, ✓ done in green. That's a working AI TUI in ~80 lines.
Level Up: The Forever Loop (REPL Mode)
The single-prompt TUI is great for one-shot tasks. But what if you want Claude to just... keep responding? Like the real Claude Code experience — type a prompt, get a response, type another?
That's a REPL, and it's a while (true) loop:
// repl.ts
import { query } from "@anthropic-ai/claude-agent-sdk" ;
import * as readline from "readline" ;
let sessionId : string | undefined ;
const rl = readline. createInterface ({ input: process.stdin, output: process.stdout });
const ask = ( prompt : string ) => new Promise < string >(( resolve ) => rl. question (prompt, resolve));
async function runTurn ( userPrompt : string ) {
for await ( const msg of query ({
prompt: userPrompt,
options: { allowedTools: [ "Read" , "Glob" , "Grep" , "Bash" ], resume: sessionId },
})) {
if (msg.type === "system" && msg.subtype === "init" ) sessionId = msg.session_id;
if (msg.type === "assistant" ) {
for ( const block of msg.message.content) {
if (block.type === "text" ) process.stdout. write ( ` \n 🤖 ${ block . text } \n ` );
if (block.type === "tool_use" ) {
process.stdout. write ( `⚙ ${ block . name }(${ JSON . stringify ( block . input ). slice ( 0 , 80 ) }) \n ` );
}
}
}
}
}
console. log ( "◆ Claude REPL — type your prompt, ctrl+c to quit \n " );
while ( true ) {
const input = await ask ( " \n > " );
if ( ! input. trim ()) continue ;
await runTurn (input. trim ());
}
Run it with npm run repl . Type anything. Claude responds. Type again — it still has context from everything before. That's the resume: sessionId doing its job.
The real power is hooks — callbacks that fire at key points in the agent lifecycle. This is how you add audit logs, approval gates, or custom UI feedback:
// agent.ts (with hooks)
import { query } from "@anthropic-ai/claude-agent-sdk" ;
import { appendFile } from "fs/promises" ;
const auditHook = async ( input : any ) => {
const tool = input.tool_name ?? "unknown" ;
const args = JSON . stringify (input.tool_input ?? {}). slice ( 0 , 100 );
await appendFile ( "./audit.log" , `${ new Date (). toISOString () } ${ tool } ${ args } \n ` );
return {};
};
for await ( const message of query ({
prompt: "Refactor the auth module" ,
options: {
allowedTools: [ "Read" , "Edit" , "Bash" ],
hooks: {
PostToolUse: [{ matcher: ".*" , hooks: [auditHook] }],
},
},
})) {
// render to your TUI
}
Every tool call gets logged to audit.log with a timestamp. matcher: ".*" catches everything — narrow to "Edit|Write" if you only care about mutations.
Other hooks worth knowing: PreToolUse to block operations before they run, Stop to detect when the agent finishes, UserPromptSubmit to pre-process or validate input.
The agent remembers context across multiple query() calls. Capture the session ID from the first run, pass it to the next:
let sessionId : string | undefined ;
// First turn — Claude reads the file
for await ( const msg of query ({ prompt: "Read the auth module" })) {
if (msg.type === "system" && msg.subtype === "init" ) {
sessionId = msg.session_id;
}
}
// Second turn — zero tool calls, Claude already knows
for await ( const msg of query ({
prompt: "Now find everything that calls it" ,
options: { resume: sessionId },
})) {
// ...
}
The money detail: the second turn fires zero tool calls — Claude already has the file in context. No re-reading, no extra API calls. It just answers.
Run the demo: npm run session .
Ink is great for quick TUIs. But if you want richer widgets — tables, command palettes, split panes, charts, modals — Rezi is the upgrade path. Still TypeScript, still Node.js, but native-backed rendering through a C engine and 50+ built-in widgets.
Where Ink feels like React (hooks, JSX, component tree), Rezi is state-driven: you define a view function that maps state → UI, and call app.update() to change state. Same mental model as Bubble Tea's Elm Architecture, but in TypeScript.
npm install @rezi-ui/core @rezi-ui/node
Here's the same Claude TUI rebuilt in Rezi:
// rezi/rezi-app.ts
import { ui } from "@rezi-ui/core" ;
import { createNodeApp } from "@rezi-ui/node" ;
import { query } from "@anthropic-ai/claude-agent-sdk" ;
type LineKind = "user" | "agent" | "tool" | "result" ;
type LogLine = { kind : LineKind ; text : string };
type State = { lines : LogLine []; done : boolean };
const prompt = process.argv. slice ( 2 ). join ( " " ) || "What files are in this directory?" ;
const app = createNodeApp < State >({
initialState: { lines: [{ kind: "user" , text: `> ${ prompt }` }], done: false },
});
const kindVariant : Record < LineKind , string > = {
user: "info" ,
agent: "body" ,
tool: "warning" ,
result: "success" ,
};
app. view (( state ) =>
ui. page ({
p: 1 ,
gap: 1 ,
header: ui. header ({ title: "◆ My AI Terminal" , subtitle: "q to quit" }),
body: ui. panel ( "Output" , [
... state.lines. map (( line , i ) =>
ui. text (line.text, { key: String (i), variant: kindVariant[line.kind] as any })
),
... ( ! state.done ? [ui. spinner ({ label: "thinking…" , key: "spinner" })]

[truncated]
