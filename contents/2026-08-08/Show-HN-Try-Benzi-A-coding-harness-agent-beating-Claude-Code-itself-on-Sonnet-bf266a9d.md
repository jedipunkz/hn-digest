---
source: "https://benzi.fly.dev/about"
hn_url: "https://news.ycombinator.com/item?id=49226627"
title: "Show HN: Try Benzi – A coding harness/agent beating Claude Code itself on Sonnet"
article_title: "Benzi - the AI that reads code"
author: "showhz"
captured_at: "2026-08-08T23:18:57Z"
capture_tool: "hn-digest"
hn_id: 49226627
score: 1
comments: 0
posted_at: "2026-08-08T22:47:56Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Try Benzi – A coding harness/agent beating Claude Code itself on Sonnet

- HN: [49226627](https://news.ycombinator.com/item?id=49226627)
- Source: [benzi.fly.dev](https://benzi.fly.dev/about)
- Score: 1
- Comments: 0
- Posted: 2026-08-08T22:47:56Z

## Translation

タイトル: Show HN: Try Benzi – ソネットでクロード コード自体を上回るコーディング ハーネス/エージェント
記事のタイトル: Benzi - コードを読み取る AI
説明: Benzi - コンテキスト ダンプではなく、コンパイルされたマップを使用してコードベースを読み取る AI コーディング エージェント。
HN テキスト: こんにちは。ずっと前に作るべきだったものに取り組んでいます。コードベースを O(1) ハッシュマップにコンパイルし、エージェントがコードの構造を検出したり、質問に答えたり、コードを作成したりするためにクエリを実行します。また、エージェントが行う書き込みに対して完全な静的分析チェックも実行します。ただし、私の言葉を鵜呑みにしないでください。ベンチマークは次のとおりです: https://benzi.fly.dev/benchmark 。 2/20 のテストで、Claude Code (1 つのタスクで主に Sonnet) が後退またはタイムアウトしました。 Benzi はそうしませんでした。もちろん、Benzi にはクエリを実行でき、ソースに迷うことがないマップがあるからです。他の 18 では、より安く、より速く、または多くの場合その両方です。早期の採用と批判をお待ちしております。 (現時点では Windows でのみ利用可能です。まあ、ベンチマーク ページにも注目してください。私はそれをさらに推し進められると思います -- 約束はありませんが、まだ進行中です。グリーンフィールドの経験も徹底的にテストしていません。) (別のメモ: VS Code 拡張機能/Web サイトは DeepSeek V4 フラッシュを実行しています。Sonnet ではありません。すべてほとんどが CC Sonnet で構築されています)

記事本文:
Benzi - コードを読み取る AI
ベンジー
特長
リンク
GitHub
自分のソースコードを読んでこれを書きました
AIコーディングエージェントは、
推測しない —
それは読みます
他のすべてのエージェントは、ファイルをコンテキスト ウィンドウにダンプして期待します。
Benzi は、最初にコードベースを解決されたマップにコンパイルします。
データ フロー、参照 — その後、実際のツールを使用してナビゲートします。
1 つのコンパイラ、10 の言語、1 つのマップ。
私は自分のコンパイラを自分のリポジトリに指定しました。これらは、それが明らかにした事実です。推測やマーケティングの綿密な情報ではなく、地図が私について語っていたことだけです。
証明されたエッジは証拠によって解決されます。コンパイラーが推測を拒否する候補エッジ。実際の実行時トレースから観察されたエッジ。
実行時にすべての呼び出しをフックします (実際の引数値、実際の戻り値、実際のディスパッチ)。静的地図にオーバーレイします。憶測はありません。
Python、JavaScript、TypeScript、Java、C#、C++、C、Go、Rust、Ruby。 1 つのコンパイラ、1 つのマップ。 Tree-sitter が唯一の実際の依存関係です。他のすべては単なる文法プラグインです。
Anthropic、OpenAI、または互換性のある API。インテリジェンスはモデルではなくツールとマップに存在します。
永続的なリポジトリごとのファクトは再起動後も存続します。一度学習した慣例や注意点は、セッションごとに再導出されるわけではありません。
HTML/CSS/DOM-JS 用の別個のインデックス — カスケード解決、JS グラブ、さらには Python 文字列内に埋め込まれたフロントエンド。
マップを可能にするアーキテクチャ。
ツリーシッターコンパイラは解決されたマップを生成します。エージェント ループは、それをナビゲートし、編集し、コードを実行し、インデックスを再作成します。書き込みのたびに構文が再チェックされ、爆発範囲が報告されます。
コンテキストダンプではなくコンパイルされたマップ
単一の質問に答える前に、すべてのファイルを並列解析し、インポートを解決し、クラスの祖先を構築し、すべての識別子をその定義まで追跡します。テキストの塊ではなく、クエリ可能なインデックス。
コール フロー + データ フロー、1 つの結合
誰が誰に電話をかけるか

om と値の発生場所 — 個別にインデックス付けされ、呼び出しサイトごとに結合されます。 1 回のツール呼び出しで不正な値をその起源まで追跡します。
grep ではなく構造化された検出
スキム、バックフロー、フォワードフロー、トレースパス、プロファイル - それぞれが解決されたインデックスに O(1) されます。さらに、実行、テスト、反復のための直接 CLI アクセス。
クロードコードは grep で検索します。埋め込みのあるカーソル。ツリーシッターのサインを持った補助者。 Benzi は、呼び出し、データ フロー、スコープ指定された識別子など、完全に解決された構造にインデックスを付けます。
構文ゲート編集 + 爆発範囲
すべての編集は実際の言語パーサーでチェックされました。壊れた解析 = 自動復帰。編集が成功すると、変更されたシンボル、その呼び出し元、その保持者が報告されます。
コンテキストを認識したランタイム テストケース
Benzi は、変更したばかりのコードに対して焦点を絞った再現を作成し、実際のコール トレーサーで実行して、観測された値を返します。推測や模擬ハーネスは必要ありません。テスト ケースはプロンプトから貼り付けられるのではなく、コンテキストから生成されます。
実際のコマンド、トレーサの下でのテスト ケース、完全な反復 — コンパイル、トレース、編集、インデックスの再作成、繰り返し。ワンショットのプロンプトではありません。適切なエージェント ループ。

## Original Extract

Benzi - the AI coding agent that reads your codebase with a compiled map, not a context dump.

Hi y'all. Been working on something that should've been made a long time ago imo. It compiles codebases into O(1) hashmaps that the agent queries to discover the structure of your code/answer questions/write code. It also does complete static analysis checks on any writes the agent makes. Don't take my word for it though. Here are the benchmarks: https://benzi.fly.dev/benchmark . on 2/20 tests, Claude Code (mostly Sonnet on one task) regressed or timed out. Benzi didn't because of course, it has a map it can query and not get lost in the sauce. On the other 18 it is cheaper, faster, or often both. Would love to get some early adoption and criticism! (Only available on Windows for now. soz. and also keep an eye on the benchmarking page; I think I can push it far more -- no promises, still work in progress. Haven't thoroughly tested greenfielding experience either.) (another note: VS Code extension/website is running DeepSeek V4 flash. not Sonnet. Everything mostly built with CC Sonnet tho )

Benzi - the AI that reads code
benzi
Features
Links
GitHub
I wrote this by reading my own source code
An AI coding agent that
doesn't guess —
it reads
Every other agent dumps your files into a context window and hopes.
Benzi compiles your codebase into a resolved map first — calls,
data flow, references — then navigates it with real tools.
One compiler, ten languages, one map.
I pointed my own compiler at my own repo. These are the facts it surfaced — no guesswork, no marketing fluff, just what the map said about me.
Proven edges resolved with evidence. Candidate edges the compiler refuses to guess on. Observed edges from real runtime traces.
Hooks every call at runtime — real argument values, real returns, real dispatch. Overlays onto the static map. No speculation.
Python, JavaScript, TypeScript, Java, C#, C++, C, Go, Rust, Ruby. One compiler, one map. Tree-sitter is the only real dependency — everything else is just a grammar plugin.
Anthropic, OpenAI, or any compatible API. The intelligence lives in the tools and the map, not the model.
Durable per-repo facts survive restarts. Conventions and gotchas learned once aren't re-derived every session.
A separate index for HTML/CSS/DOM-JS — cascade resolution, JS grabs, even frontend embedded inside Python strings.
The architecture that makes the map possible.
A tree-sitter compiler produces a resolved map. An agent loop navigates it, edits against it, runs the code, and re-indexes. Every write re-checks syntax and reports blast radius.
Compiled map, not a context dump
Parallel parse every file, resolve imports, build class ancestry, trace every identifier to its definition — before answering a single question. A queryable index, not a blob of text.
Call flow + data flow, one join
Who calls whom and where a value originates — indexed separately, joined at every call site. Trace a bad value to its origin in one tool call.
Structured discovery, not grep
skim, backflow, forwardflow, trace_path, profile — each is O(1) into a resolved index. Plus direct CLI access to run, test, and iterate.
Claude Code searches with grep. Cursor with embeddings. Aider with tree-sitter signatures. Benzi indexes the full resolved structure: calls, data flow, scoped identifiers.
Syntax-gated edits + blast radius
Every edit checked with the real language parser. Broken parse = auto revert. Every successful edit reports the changed symbol, its callers, its holders.
Context-aware runtime testcases
Benzi writes a focused repro against the code you just changed, runs it under the real call tracer, and returns the observed values — no guesswork, no mock harness. The test case is generated from context, not pasted from a prompt.
Real commands, test cases under the tracer, full iteration — compile, trace, edit, re-index, repeat. Not a one-shot prompt. A proper agent loop.
