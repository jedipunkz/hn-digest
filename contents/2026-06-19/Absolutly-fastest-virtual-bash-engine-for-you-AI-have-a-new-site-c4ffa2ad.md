---
source: "https://bashkit.sh/"
hn_url: "https://news.ycombinator.com/item?id=48600735"
title: "Absolutly fastest virtual bash engine for you AI have a new site"
article_title: "Bashkit — awesomely fast virtual bash sandbox in Rust"
author: "chalyi"
captured_at: "2026-06-19T18:43:27Z"
capture_tool: "hn-digest"
hn_id: 48600735
score: 3
comments: 0
posted_at: "2026-06-19T17:14:43Z"
tags:
  - hacker-news
  - translated
---

# Absolutly fastest virtual bash engine for you AI have a new site

- HN: [48600735](https://news.ycombinator.com/item?id=48600735)
- Source: [bashkit.sh](https://bashkit.sh/)
- Score: 3
- Comments: 0
- Posted: 2026-06-19T17:14:43Z

## Translation

タイトル: あなたのための絶対最速の仮想 bash エンジン AI に新しいサイトができました
記事のタイトル: Bashkit — Rust の驚くほど高速な仮想 bash サンドボックス
説明: Bashkit は、単一の OS プロセスを生成せずに、AI エージェントから信頼できないシェル スクリプトを実行します。 156 の POSIX 準拠コマンド、仮想ファイルシステム、リソース制限、ストリーミング LLM ツール呼び出し - すべてがメモリ内で、すべてサンドボックス化されています。

記事本文:
Bashkit — Rust の驚くほど高速な仮想 bash サンドボックス bashkit の機能 ベンチのインストール Docs API AI エージェント用の仮想 bash 驚くほど高速な仮想 bash サンドボックス。 Rustで書かれています。
Bashkit は、単一の OS プロセスを生成せずに、AI エージェントから信頼できないシェル スクリプトを実行します。 156 の再実装されたコマンド、実質的な POSIX シェル言語の範囲、仮想ファイルシステム、リソース制限、およびエージェント フレームワークのツール インターフェイス - すべてがメモリ内で、すべてサンドボックス化されています。
カーゴ追加 bashkit さらに詳しく Rust docs docs.rs リファレンス 例 Rust、Python、TS
コア API は単なる「Bash」に、仮想 FS と制限を加えたものです。サイドカーなし
プロセス、コンテナブートストラップなし。
bashkit :: Bash を使用します。
#[トキオ::メイン]
async fn main () -> とにかく :: 結果 <()> {
let mut bash = Bash :: new ();
出す＝バッシュ。 exec ( "printf 'ready \n '" ) .await? ;
印刷してください！ ( "{}" 、 out .stdout);
わかりました (())
プロセスは生成されません
すべてのコマンドは Rust で再実装されます。フォークも実行もシェルエスケープもありません。
InMemoryFs、OverlayFs、MountableFs。明示的にマウントした場合のみホストにアクセスします。
コマンド、ループ、出力、入力、ファイルシステムのサイズに上限を設けます。 DoS耐性のある構造。
サンドボックスを使用可能にする組み込み機能を参照してください。
テキスト処理、ファイル、アーカイブ、ネットワーク、Python、TypeScript、
シェル制御はすべてプロセス内で実行されます。
スキルから始めて、次にランタイムを埋め込みます。
Bashkit スキルは、コーディング エージェントに適切なローカル コンテキストを提供します。
サンドボックス モデル、パッケージ API、組み込み、サンプル、エージェント ツール
パターン。まずそれをインストールしてから、Bashkit をホスト プロジェクトに追加します。
コーディング エージェントの Bashkit 固有の使用上の注意と例を示します。
npx skill add everruns/bashkit 2 エージェントに追加を依頼する
コーディング エージェントに、Bashkit をホスト プロジェクトに接続するように指示します。
bashkit を使用して、bash ツールのサポートを追加します 3 お楽しみください:)
エージェントで新しい bash ツールを使用する

ワークフロー。
エージェント、CLI、エディターなどに埋め込むことができる単一のランタイム
評価用ハーネス。サイドカープロセスなし、コンテナオーバーヘッドなし、
実行時の外部依存関係。
IEEE 1003.1-2024 シェル コマンド言語を大幅にカバーし、さらに bash 拡張機能 (配列、[[ ]]、中括弧拡張、拡張グロブ、コプロセス、トラップ) をサポートします。
grep、sed、awk、jq、curl、tar、find、xargs、その他 150 以上 - シェルアウトなしの純粋な Rust。
検出メタデータ、ストリーミング出力、およびシステム プロンプトを備えた BashTool。任意のエージェント フレームワークにプラグインします。
行編集と複数行入力を使用して、ローカル REPL に対して引数なしで bashkit を実行します。
シェルの状態と VFS の内容をバイトにシリアル化します。あらゆるワークロードをチェックポイントし、どこからでも再開できます。
ToolDef + コールバックのペアを bash スクリプトによって駆動される ScriptedTool に構成します。
コアの Rust クレートから始めるか、同じランタイムを Python にドロップします
既存のスタック内で必要な場合は TypeScript を使用します。
カーゴ追加 bashkit bashkit を使用:: Bash ;
#[トキオ::メイン]
async fn main () -> とにかく :: 結果 <()> {
let mut bash = Bash :: new ();
出す＝バッシュ。 exec ( "echo hello world" ) .await? ;
プリントイン！ ( "{}" 、 out .stdout);
わかりました (())
直接 Bash API を使用した Python PyO3 ホイール
pip install bashkit from bashkit import Bash
bash = バッシュ()
result = bash.execute_sync( "echo 'Hello, World!'" )
print (結果.stdout)
bash.execute_sync( "export APP_ENV=dev" )
print (bash.execute_sync( "echo $APP_ENV" ).stdout) Node、Bun、Deno の TypeScript NAPI-RS ランタイム
npm i @everruns/bashkit import { Bash } from "@everruns/bashkit" ;
const bash = 新しい Bash ();
const 結果 = bash。 executeSync ( 'echo "Hello, World!"' );
コンソール。ログ (結果.stdout);
バッシュ。 executeSync ( "X=42" );
コンソール。ログ (bash.executeSync ( "echo $X" ).stdout);エンドツーエンド
bashkit をインプロセス シェルとして使用する小さな Rust プログラム。いいえ
コンテナ、なし

ubprocess — 単なるクレートです。
bashkit :: Bash を使用します。
#[トキオ::メイン]
async fn main () -> とにかく :: 結果 <()> {
let mut bash = Bash :: new ();
バッシュ。 exec ( "mkdir -p /tmp/data" ) .await? ;
バッシュ。 exec ( "echo 'hello' > /tmp/data/out.txt" ) .await? ;
r = bash とします。 exec ( "cat /tmp/data/out.txt | tr a-z A-Z" ) .await? ;
印刷してください！ ( "{}" 、 r . stdout); // こんにちは
わかりました (())
セキュリティ敵対的な入力がデフォルトの想定です
プロセス、ファイルシステム、ネットワークなど、あらゆる層にわたる多層防御
パーサーとランタイム。を参照してください。
250 以上の軽減策を備えた完全な脅威モデル。
156 個のコマンドが Rust で再実装されました。fork、exec、またはシェルエスケープはありません。
スクリプトはデフォルトでメモリ内の FS を参照します。マウントされていない限り、ホストにアクセスできません。
HTTP はデフォルトで拒否されます。各ドメインを明示的に許可する必要があります。
コマンド (10K)、ループ (100K)、関数の深さ、出力 (10MB)、入力 (10MB) に制限されます。
タイムアウト、燃料バジェット、AST 深度 — 異常な入力によってインタープリターがハングアップすることはありません。
すべての組み込みは catch_unwind でラップされています。 1 つのコマンドでパニックが発生してもホストがクラッシュすることはありません。
Bashkit には、15 のエージェントにわたる 58 タスクの LLM 評価ハーネスが同梱されています
カテゴリー。以下の結果は、
2026-02-28
実行:
クレートは MIT ライセンスを取得しています。これらは実際に役立つリンクです
それを評価し、統合し、運用します。
コア クレート ドキュメント、ビルダー オプション、制限、およびシェル セマンティクス。
Bash の直接使用法、スナップショット、組み込みに関する PyO3 パッケージのドキュメント。
NAPI バインディングに関する Node、Bun、および Deno ランタイム ドキュメント。
パーサー、VFS、ネットワーク、ランタイムにわたる 268 件の文書化された脅威ケース。
ベンチマーク、基準ベンチ、評価にわたるインタラクティブな傾向。
ワンショット コマンド、スクリプト実行、対話型シェルの使用。
Rust、Python、JavaScript、ツール フローのリファレンス プログラム。
Everruns エコシステムの一部。

## Original Extract

Bashkit runs untrusted shell scripts from AI agents without spawning a single OS process. 156 POSIX-compliant commands, a virtual filesystem, resource limits, and streaming LLM tool calls — all in-memory, all sandboxed.

Bashkit — awesomely fast virtual bash sandbox in Rust bashkit Features Install Benches Docs API Virtual bash for AI agents An awesomely fast virtual bash sandbox. Written in Rust.
Bashkit runs untrusted shell scripts from AI agents without spawning a single OS process. 156 reimplemented commands, substantial POSIX shell language coverage, a virtual filesystem, resource limits, and tool interfaces for agent frameworks — all in-memory, all sandboxed.
cargo add bashkit Go deeper Rust docs docs.rs reference Examples Rust, Python, TS
The core API is just `Bash`, plus virtual FS and limits. No sidecar
process, no container bootstrap.
use bashkit :: Bash ;
#[tokio :: main]
async fn main () -> anyhow :: Result <()> {
let mut bash = Bash :: new ();
let out = bash . exec ( "printf 'ready \n '" ) .await? ;
print! ( "{}" , out . stdout);
Ok (())
} No process spawning
Every command is reimplemented in Rust. No fork, no exec, no shell escape.
InMemoryFs, OverlayFs, MountableFs. Host access only when you explicitly mount it.
Caps on commands, loops, output, input, and filesystem size. DoS-resistant by construction.
Browse the builtins that make the sandbox usable.
Text processing, files, archives, network, Python, TypeScript,
and shell control all live in-process.
Start with the skill, then embed the runtime.
The Bashkit skill gives coding agents the right local context:
sandbox model, package APIs, builtins, examples, and agent-tool
patterns. Install it first, then add Bashkit to the host project.
Give your coding agent Bashkit-specific usage notes and examples.
npx skills add everruns/bashkit 2 Ask agent to add it
Prompt your coding agent to wire Bashkit into the host project.
Using bashkit, add support for a bash tool 3 Enjoy :)
Use the new bash tool in your agent workflow.
A single runtime you can embed in agents, CLIs, editors, and
evaluation harnesses. No sidecar process, no container overhead, no
external dependencies at runtime.
Substantial IEEE 1003.1-2024 Shell Command Language coverage, plus bash extensions: arrays, [[ ]], brace expansion, extended globs, coprocesses, traps.
grep, sed, awk, jq, curl, tar, find, xargs, and 150+ more — pure Rust, no shelling out.
BashTool with discovery metadata, streaming output, and system prompts. Plug into any agent framework.
Run bashkit with no args for a local REPL with line editing and multiline input.
Serialize shell state and VFS contents to bytes. Checkpoint any workload, resume anywhere.
Compose ToolDef + callback pairs into a ScriptedTool driven by a bash script.
Start with the core Rust crate, or drop the same runtime into Python
and TypeScript when you need it inside an existing stack.
cargo add bashkit use bashkit :: Bash ;
#[tokio :: main]
async fn main () -> anyhow :: Result <()> {
let mut bash = Bash :: new ();
let out = bash . exec ( "echo hello world" ) .await? ;
println! ( "{}" , out . stdout);
Ok (())
} Python PyO3 wheel with direct Bash API
pip install bashkit from bashkit import Bash
bash = Bash()
result = bash.execute_sync( "echo 'Hello, World!'" )
print (result.stdout)
bash.execute_sync( "export APP_ENV=dev" )
print (bash.execute_sync( "echo $APP_ENV" ).stdout) TypeScript NAPI-RS runtime for Node, Bun, Deno
npm i @everruns/bashkit import { Bash } from "@everruns/bashkit" ;
const bash = new Bash ();
const result = bash. executeSync ( 'echo "Hello, World!"' );
console. log (result.stdout);
bash. executeSync ( "X=42" );
console. log (bash. executeSync ( "echo $X" ).stdout); End-to-end
A tiny Rust program that uses bashkit as an in-process shell. No
container, no subprocess — just a crate.
use bashkit :: Bash ;
#[tokio :: main]
async fn main () -> anyhow :: Result <()> {
let mut bash = Bash :: new ();
bash . exec ( "mkdir -p /tmp/data" ) .await? ;
bash . exec ( "echo 'hello' > /tmp/data/out.txt" ) .await? ;
let r = bash . exec ( "cat /tmp/data/out.txt | tr a-z A-Z" ) .await? ;
print! ( "{}" , r . stdout); // HELLO
Ok (())
} Security Hostile input is the default assumption
Defense in depth across every layer — process, filesystem, network,
parser, and runtime. See the
full threat model for 250+ mitigations.
156 commands reimplemented in Rust — no fork, exec, or shell escape.
Scripts see an in-memory FS by default. No host access unless mounted.
HTTP is denied by default. Each domain must be explicitly allowed.
Caps on commands (10K), loops (100K), function depth, output (10MB), input (10MB).
Timeout, fuel budget, AST depth — pathological input can't hang the interpreter.
Every builtin is wrapped in catch_unwind. A panic in one command can't crash the host.
Bashkit ships with a 58-task LLM eval harness across 15 agentic
categories. Results below are from the
2026-02-28
run:
The crate is MIT-licensed. These are the links that actually help you
evaluate, integrate, and operate it.
Core crate docs, builder options, limits, and shell semantics.
PyO3 package docs for direct Bash usage, snapshots, and builtins.
Node, Bun, and Deno runtime docs for the NAPI bindings.
268 documented threat cases across parser, VFS, network, and runtimes.
Interactive trends across benchmarks, criterion benches, and evals.
One-shot commands, script execution, and interactive shell usage.
Reference programs for Rust, Python, JavaScript, and tool flows.
Part of the Everruns ecosystem.
