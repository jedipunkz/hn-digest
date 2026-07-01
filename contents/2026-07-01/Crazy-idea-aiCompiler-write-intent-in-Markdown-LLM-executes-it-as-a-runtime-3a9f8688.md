---
source: "https://aicompiler.dev"
hn_url: "https://news.ycombinator.com/item?id=48748336"
title: "Crazy idea? aiCompiler – write intent in Markdown, LLM executes it as a runtime"
article_title: "aiCompiler — The first programming language where the LLM is the CPU"
author: "srobbani"
captured_at: "2026-07-01T15:55:02Z"
capture_tool: "hn-digest"
hn_id: 48748336
score: 2
comments: 0
posted_at: "2026-07-01T15:20:28Z"
tags:
  - hacker-news
  - translated
---

# Crazy idea? aiCompiler – write intent in Markdown, LLM executes it as a runtime

- HN: [48748336](https://news.ycombinator.com/item?id=48748336)
- Source: [aicompiler.dev](https://aicompiler.dev)
- Score: 2
- Comments: 0
- Posted: 2026-07-01T15:20:28Z

## Translation

タイトル: クレイジーなアイデア? aiCompiler – Markdown でインテントを書き込み、LLM がランタイムとして実行します
記事のタイトル: aiCompiler — LLM が CPU である最初のプログラミング言語
説明: フレームワークではありません。エージェントラッパーではありません。新しいコンピューティング パラダイム。 .aic プログラムをプレーンな Markdown で作成します。 aiCompiler は曖昧さを取り除きます。 aiVM はステップごとに実行されます。 LLM は、CPU がマシンコードを読み取るようにオペコードを読み取ります。

記事本文:
AI aicompiler .dev デモ 仕組み GitHub 無料で試す → フレームワークではありません。エージェントラッパーではありません。新しいコンピューティング パラダイム。 LLM を搭載した最初の仮想マシン。
.aic を作成します。コードは含まず、プレーンな Markdown です。
aiCompiler はコンパイル時にすべてのあいまいさを解決します。
JVM が CPU を使用してバイトコードを実行するのと同じように、aiVM は LLM を使用してインテント オペコードを実行します。
JVM が CPU を使用して Java バイトコードを実行するのと同じように、
aiVM は、LLM を使用して .aix オペコードを実行します。
Von Neumann アーキテクチャのすべてのコンポーネントは、aiVM コンポーネントに正確にマッピングされます。 LLM は CPU ではなく、VM の実行エンジンです。他のすべてはそこから導き出されます。
CPU マシンコードを実行 RAM プログラム状態 + スタック キャッシュ 最近アクセスした高速ストレージ レジスタ 作業変数 システム バス すべてのコンポーネントを接続 ハードディスク 永続ストレージ ROM / BIOS 読み取り専用ファームウェア 🤖 2026 → aiVM コンピューター LLM がインテント オペコードを実行 — 各ステップは直接モデル推論 LLM エンジン インテント オペコードを実行 コンテキスト ウィンドウ ターンごとの作業メモリ .aix ロック ファイル キャッシュ — 同じ仕様 = インスタント ステップ状態$step_1.output、$input.x ツール アダプタ DB、電子メール、API、HTTP KV + VectorDB 永続メモリ + RAG システム プロンプト 読み取り専用、起動時にロード CPU ≡ LLM エンジン RAM ≡ コンテキスト ウィンドウ キャッシュ ≡ .aix ロック ファイル レジスタ ≡ ステップ状態 システム バス ≡ ツール アダプタ ハードディスク ≡ KV + VectorDB ROM / BIOS ≡ システムプロンプト 今すぐ試してください
ワークフローを書きます。 aiCompiler がそれをセマンティックなオペコードにコンパイルするのを見てください。 aiVM が各ステップを実行する様子を観察します。LLM は CPU のようにオペコードを読み取ります。アカウントがありません。インストールはありません。 APIキーがありません。
.aic ファイルをコンパイルします
AI オペコードを確認するには
「実行」をクリックして実行します
ワークフローを段階的に説明する
コンパイル → .aix ▶ コンパイル + 実行 .aic → aicompiler コンパイル → .aix → aivm 実行 → ステップログ 仕組み
3つのコンセプト。それ

知っておくべきことはこれだけです。
ワークフロー、ステップ、ツールなどのプログラムを平易な英語で作成します。セミコロンはありません。型の注釈はありません。定型文はありません。
# ワークフロー: 挨拶
## ステップ
### 1.「こんにちは」
{name} にクリエイティブな挨拶をしましょう。 02 ⚡ コンパイル .aix にコンパイルします。
AI コンパイラーはソース ファイルを解析し、曖昧さのない構造化された JSON オペコードであるバイトコードを生成します。
aicompiler コンパイル app.aic
→ app.aix ✓ 03 🚀 run aivm で実行
AI VM は各オペコードを読み取って実行し、実際のツールを呼び出し、データ参照を解決し、構造化された出力を返します。
aivm run app.aix
→ハロー、ワールド！ ✓ 開発者が気にする理由
フレームワークの時代は終わりつつあります。次に来るのはこれです。
LLM はコードを支援しません。それはランタイムです。各オペコードは直接命令であり、Python 接着剤やチェーン構成は必要ありません。
クロード、ノバ マイクロ、ジェミニ、ラマ — フラグを 1 つ変更します。すべてCloudflare AI Gateway経由でルーティングされます。書き直す必要はありません。
.aix ロック ファイルにより、同じソースが常に同じようにコンパイルされることが保証されます。 javacと同様に再現可能。コードと同様にバージョン管理されます。
Java のような 2 ステップのパイプライン。コンパイルと実行の間: ツール バインディングの挿入、ポリシーの再試行、セキュリティ ルール、ステップごとのプロバイダーの選択。
すべてのステップで、名前、入力、出力、期間、エラー処理などの構造化されたログが生成されます。 LLM 出力の塊ではありません。
フォークしてください。その上に独自の AI 言語を構築します。 .aix オペコード形式はオープン仕様です。
新しい種類の VM。 30秒以内にお試しください。
インストールはありません。アカウントがありません。 APIキーがありません。遊び場を開いて走ってください。

## Original Extract

Not a framework. Not an agent wrapper. A new computing paradigm. Write .aic programs in plain Markdown. aiCompiler removes ambiguity. aiVM executes step by step. The LLM reads opcodes like a CPU reads machine code.

AI aicompiler .dev Demo How it works GitHub Try it free → Not a framework. Not an agent wrapper. A new computing paradigm. The first virtual machine powered by an LLM.
You write .aic — plain Markdown, no code.
aiCompiler resolves all ambiguity at compile time.
aiVM executes intent opcodes using an LLM — just as the JVM executes bytecode using a CPU.
Just as the JVM runs Java bytecode using a CPU,
aiVM runs .aix opcodes using an LLM.
Every component of Von Neumann architecture maps exactly to an aiVM component. The LLM isn't the CPU — it's the execution engine of the VM . Everything else follows from that.
CPU Executes machine code RAM Program state + stack Cache Fast recent-access storage Registers Working variables System Bus Connects all components Hard Disk Persistent storage ROM / BIOS Read-only firmware 🤖 2026 → aiVM Computer LLM executes intent opcodes — each step is a direct model inference LLM Engine Executes intent opcodes Context Window Working memory per turn .aix Lock Files Cache — same spec = instant Step State $step_1.output, $input.x Tool Adapters DB, email, APIs, HTTP KV + VectorDB Persistent memory + RAG System Prompt Read-only, loads at boot CPU ≡ LLM Engine RAM ≡ Context Window Cache ≡ .aix Lock Files Registers ≡ Step State System Bus ≡ Tool Adapters Hard Disk ≡ KV + VectorDB ROM / BIOS ≡ System Prompt Try it right now
Write a workflow. Watch aiCompiler compile it to semantic opcodes. Watch aiVM execute each step — the LLM reading opcodes like a CPU. No account. No install. No API key.
Compile your .aic file
to see the AI opcodes
Click Run to execute
the workflow step by step
Compile → .aix ▶ Compile + Run .aic → aicompiler compile → .aix → aivm run → step log How it works
Three concepts. That's all you need to know.
Write your program in plain English — workflows, steps, tools. No semicolons. No type annotations. No boilerplate.
# Workflow: Greeter
## Steps
### 1. SayHello
Say hello to {name} creatively. 02 ⚡ compile Compile to .aix
The AI compiler parses your source file and produces bytecode — structured JSON opcodes with zero ambiguity.
aicompiler compile app.aic
→ app.aix ✓ 03 🚀 run Run with aivm
The AI VM reads each opcode and executes it — calling real tools, resolving data references, returning structured output.
aivm run app.aix
→ Hello, World! ✓ Why developers care
The framework era is ending. This is what comes next.
The LLM doesn't assist your code. It IS the runtime. Each opcode is a direct instruction — no Python glue, no chain configuration.
Claude, Nova Micro, Gemini, Llama — change one flag. All routed through Cloudflare AI Gateway. No rewrite needed.
.aix lock files ensure the same source always compiles identically. Reproducible like javac. Version-controlled like code.
Two-step pipeline like Java. Between compile and run: inject tool bindings, retry policies, security rules, provider selection per step.
Every step produces a structured log: name, inputs, outputs, duration, error handling. Not a blob of LLM output.
Fork it. Build your own AI language on top of it. The .aix opcode format is an open spec.
A new kind of VM. Try it in 30 seconds.
No install. No account. No API key. Just open the playground and run.
