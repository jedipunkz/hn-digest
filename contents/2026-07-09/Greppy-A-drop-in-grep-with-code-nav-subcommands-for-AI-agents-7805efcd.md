---
source: "https://github.com/metric-space-ai/greppy"
hn_url: "https://news.ycombinator.com/item?id=48841902"
title: "Greppy – A drop-in grep with code-nav subcommands for AI agents"
article_title: "GitHub - metric-space-ai/greppy: Standard grep, plus a few commands your coding agent can use to navigate code — who-calls, impact, context, brief. ~2x faster and ~3-4x cheaper agentic code navigation. One native Rust binary. · GitHub"
author: "metricspaceai"
captured_at: "2026-07-09T07:39:45Z"
capture_tool: "hn-digest"
hn_id: 48841902
score: 1
comments: 0
posted_at: "2026-07-09T06:52:17Z"
tags:
  - hacker-news
  - translated
---

# Greppy – A drop-in grep with code-nav subcommands for AI agents

- HN: [48841902](https://news.ycombinator.com/item?id=48841902)
- Source: [github.com](https://github.com/metric-space-ai/greppy)
- Score: 1
- Comments: 0
- Posted: 2026-07-09T06:52:17Z

## Translation

タイトル: Greppy – AI エージェント用の code-nav サブコマンドを備えたドロップイン grep
記事のタイトル: GitHub - metric-space-ai/greppy: 標準の grep に加えて、コーディング エージェントがコードをナビゲートするために使用できるいくつかのコマンド (誰が呼び出すか、影響力、コンテキスト、概要)。約 2 倍高速で、約 3 ～ 4 倍安価なエージェント コード ナビゲーション。 1 つのネイティブ Rust バイナリ。 · GitHub
説明: 標準の grep に加えて、コーディング エージェントがコードをナビゲートするために使用できるいくつかのコマンド (誰が呼び出すか、影響するか、コンテキスト、概要など) を追加します。約 2 倍高速で、約 3 ～ 4 倍安価なエージェント コード ナビゲーション。 1 つのネイティブ Rust バイナリ。 - メトリックスペース-ai/greppy

記事本文:
GitHub - metric-space-ai/greppy: 標準の grep に加えて、コーディング エージェントがコードをナビゲートするために使用できるいくつかのコマンド (誰が呼び出すか、影響するか、コンテキスト、概要を示す) が含まれています。約 2 倍高速で、約 3 ～ 4 倍安価なエージェント コード ナビゲーション。 1 つのネイティブ Rust バイナリ。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ

おお！
読み込み中にエラーが発生しました。このページをリロードしてください。
メトリックスペースAI
/
グレッピー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
12 コミット 12 コミット .github/ workflows .github/ workflows アセット アセット ベンチ ベンチ クレート クレート docs/ アセット docs/ アセット .gitattributes .gitattributes .gitignore .gitignore ADDING_A_LANGUAGE.md ADDING_A_LANGUAGE.md Cargo.lock Cargo.lock Cargo.toml Cargo.tomlライセンス ライセンス README.md README.md THIRD_PARTY.md THIRD_PARTY.md Rust-toolchain.toml Rust-toolchain.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
標準の grep に加えて、コーディング エージェントがコードをナビゲートするために使用するいくつかのコマンド ( who-calls 、 Impact 、 semantic-search 、 Brief ) が追加されています。構造的なコード ナビゲーションの質問に対して、エージェントは、より少ないトークンを使用して、単純な grep では ~53% の確率で正しく回答します。 1 つのネイティブ Rust バイナリ。
greppy はドロップイン grep です。すべてのフラグは以前とまったく同じように機能します。
エージェントが通常尋ねる質問に答えます: この電話をかけてきたのは誰ですか
関数、変更すると何が壊れるのか、X を実行するコードはどこにあるのか。
エージェントの設定 (下記) は余分なコマンドが存在することを伝え、エージェントは停止します
テキスト一致をループします。
# 標準 grep — すべてのコマンドは変更されずに機能します。
greppy -rn " TODO " src/
greppy -i "接続が拒否されました"server.log
# 同じバイナリ上にいくつかの追加コマンド:
greppy who-call parse_config # この関数を呼び出す人
greppy Impact User --direction incoming # ユーザーを変更すると何が壊れるのか
greppy semantic-search " 値を範囲に制限する " # 意味によってコードを検索
greppy Brief _split_blueprint_path # 定義 + 呼び出し元 + 呼び出し先
同じコーディング エージェント (MiniMax-M3、によって駆動)

Pi Code ) は、実際のリポジトリに関する 1 つの誰が呼び出したのかという質問に答えます。左は単純な grep 、右は greppy です。 greppy は、grep-and-read スパイラルではなく、単一の greppy who-calls 呼び出しで呼び出し元を解決します。2.3 倍高速、14 → 5 ツール呼び出し、最大 9 倍少ない入力トークン。カウンターは記録された実行からライブで表示されます。
カーゴビルド --release --bin greppy --features 埋め込みモデル
sudo install -m 0755 target/release/greppy /usr/local/bin/greppy
すべてが自動です。コードグラフとセマンティックモデルが組み込まれています。
バイナリを作成し、最初の使用時に自分自身をビルドします。インデックスするものは何もありません、何もありません
ダウンロード、設定するフラグはありません。 (macOS / Linux / Windows 用のビルド済みバイナリ
はリリースページにあります。) 透過的な grep として必要です
ドロップインも？もう一度 grep としてインストールします。
2. 追加のコマンドが存在することをエージェントに伝えます。エージェントに委任します
チャット、たとえば https://github.com/metric-space-ai/greppy/ をインストールする — または
エージェントがプロジェクトの指示のために読み取るファイルに、以下のスニペットを貼り付けます。
( CLAUDE.md 、 AGENTS.md 、 .cursor/rules 、 .windsurfrules 、またはシステム
プロンプト)。
このプロジェクトには「greppy」が含まれています。これは標準の grep に加えて、いくつかのコード ナビゲーション コマンドです。
事前に構築されたシンボル グラフとデバイス上のセマンティック インデックスを使用します。通常の grep ごとに
呼び出し (およびフラグ) は通常どおりに機能します。
コードナビゲーションコマンド。 SYMBOL は関数/メソッド/クラス/型の名前です。
これらは、テキスト一致ではなく、解決された結果を「qualified_name file:line」として返します。
greppy who-calls SYMBOL SYMBOL の発信者 (着信通話)
greppy の呼び出し先 SYMBOL SYMBOL が呼び出す関数 (発信呼び出し)
greppy find-usages SYMBOL SYMBOL へのすべての参照 (呼び出し、使用、インポート)
greppy Brief SYMBOL SYMBOL の定義とその呼び出し元と呼び出し先を 1 回の呼び出しで
greppy の影響 SYMBOL SYMBOL への変更が及ぼす推移的なコードのセット
グラム

eppy search-symbols 名前が NAME (名前またはフラグメント) と一致する NAME 定義
greppy path --from A --to B シンボル A からシンボル B までの呼び出しチェーン (存在する場合)
セマンティック検索 — シンボルの名前がわからない場合に使用します。
greppy semantic-search "わかりやすい英語の説明"
探している動作やコードを平易な英語で説明してください
(例: 「値を範囲に制限する」、「失敗した HTTP リクエストを再試行する」)。
意味 (署名 + ファイル:行) によって最も一致する定義を返します。
EXPAND — ファイルを手動で開く代わりに、1 回の呼び出しで完全なソースを取得します。
greppy 拡張 ID
who-calls / callees / Impact / semantic-search は出力を
行「展開: greppy 展開 <id>」。実行して準備された証拠を印刷します
パック — 上位一致の完全なソースをバンドルして — 1 回の呼び出しで、
各 file:line を自分で読み取る代わりに。
FLAGS (上記のコマンドに追加):
--code には各結果のソース行が含まれます (そのため、個別に読み取る必要はありません)
--all すべての結果を返します (デフォルトの切り捨てをオフにします)
--json 正確なカウントを含む機械可読出力
--root DIR 現在のディレクトリ以外のリポジトリに対して実行します。
--kind KIND (検索シンボル) 関数|メソッド|クラス|構造|列挙型|特性に制限
--方向 着信|発信、 --深さ N (衝撃) どの方向にどのように
[切り捨てられた]
エージェントが実際に支払うのは、請求されたトークンと実時間です。
ベンチマーク: 14 のコーディング エージェント モデル (Claude Opus/Sonnet/Fable、GPT-5.5、Gemini、Grok、DeepSeek、Qwen、GLM、Kimi、MiniMax-M3 など) はそれぞれ Pi Code によって駆動され、7 つの実際のリポジトリ (Rust serde + tokio 、Python flask + django 、Java gson 、TypeScript) にわたる 35 のコード ナビゲーションの質問に答えます。ゾッド、ゴーヒューゴ）。すべてのタスクは 2 回実行されます。1 回目はプレーンな grep で、もう 1 回目は greppy で、同じエージェント、同じプロンプトです。答えはフロアグレードで評価されます: 合格

s は、グラウンド トゥルース シンボル/ファイル (各アンカーは生成時に検証済み) に名前を付ける必要があります。ハーネスは bench/agent_efficiency/ にあります。
正しさが見出しです。構造ナビゲーション ( who-calls 、 callees 、 Impact/blast-radius 、 find-symbol) では、エージェントは greppy では 87% の確率で正しく応答しますが、単純な grep では 53% の確率で正解します (リポジトリ独自の Grade_answers.py によって評価されます)。 35 の質問すべてで: 90% 対 63%。単純な grep は安価ですが、確実に間違っていることがよくあります。 greppy は 1 回の呼び出しで関係を解決します。
したがって、これは精度とコストのトレードではありません。構造的な問題に関しては、greppy のほうがより正確であり、安価です。
単純な grep が続くところ: 「このサブシステムはどのように動作するか」という自由回答型の質問。どちらのツールもそこで答えに到達します (どちらにしても約 98% 正解) が、greppy の正確なロケーターを使用すると、エージェントはメカニズムを説明するためにさらに読む必要があるため、コストが少し高くなります。 greppy の強みは、ピンポイントで構造的な質問です。セマンティック パスも、エージェントを 1 ステップで答えに導くように調整されています。
ゲインはモデルによって異なります。各モデルの実際の OpenRouter リスト レートで価格を設定すると、構造タスクの実際のドルコストは、ほとんどのモデルで greppy で下がります (中央値で約 16% 安くなります) が、変動が大きく (MiMo の +54% から Grok 4.3 の -84% まで、螺旋を描きます)、モデルの一般的なエージェント ベンチマーク スコアは追跡されません。独自のモデルをベンチマークします。ほとんどのモデルが優れており、すべてのモデルの正確性が向上します。
実際のコスト = 各モデルの OpenRouter 定価 × 構造タスクを合計したトークン。散布図は、効率の向上はほとんどのモデルで実際のものですが、モデルのエージェント ランクが予測するものではなく、真にモデルに依存していることを示しています。
標準の grep。追加のコマンドのいずれでもない呼び出しでは、実際の grep が実行され、その出力と終了コードが変更されずに返されます。
プレコ

計算されたコードグラフ。インデックス付きの型付きシンボル グラフ ( CALLS / USES / TYPE_REF / IMPORTS ) は、who-calls/callees/find-usages/impact/path に直接応答します。テキストの一致ではなく、 file:line との関係が解決されます。複数の grep+read ラウンドを 1 つの呼び出しにまとめます。
ネイティブのセマンティック検索。コードと単語を共有しない自然言語クエリの場合、semantic-search は greppy 独自のネイティブ Rust 推論 (CPU / Apple Metal / NVIDIA CUDA、自動検出 — llama.cpp なし、Python なし、HTTP なし) に Google の EmbeddingGemma を使用してクエリを埋め込み、意味別に最も近いコード スパンを返します。小さなウォーム デーモンは呼び出し間でモデルを常駐させ、アイドル状態になったら削除するため、検索を行っていない間は GPU メモリを保持することはありません。
1 つのネイティブ Rust バイナリ。 EmbeddingGemma モデルはバイナリに焼き付けられます。ツリーシッター パーサーと SQLite は静的にコンパイルされます。
初期段階から進化中 — ドロップイン grep コアは堅牢です。その周りのインテリジェンス層はベータ版です。
実線: サポートされている言語の grep ドロップインおよびコードグラフ コマンド (who-calls / callees / find-usages / Impact / path / Brief)。
サポートされている言語 (107): python 、 csharp 、 go 、 cpp 、 php 、 Rust 、 Swift 、 scala 、 c 、 java 、 javascript 、 typescript 、 Ruby 、 bash 、 kotlin 、 fsharp 、 julia 、 ocaml 、 d 、 gdscript 、 zig 、 elm 、 erlang 、 Crystal 、 gleam 、 objc 、 Solidity 、 prisma 、 protobuf 、 css 、 dockerfile 、 json 、 groovy 、 lua 、 SQL 、 make 、 nix 、 cmake 、 dart 、 fortran 、 elixir 、 スキーム 、 vue 、 astro 、 svelte 、 verilog 、 glsl 、 hcl 、 matlab 、r 、 purescript 、racket 、clojure 、 haskell 、cuda 、tcl 、graphql 、パスカル 、powershell 、html 、yaml 、hlsl 、cobol 、fish 、ini 、vhdl 、json5 、awk 、cairo 、ada 、hare 、kdl 、jsonnet 、llvm 、janet 、 jinja2 、 上腕二頭筋 、 gotemplate 、 just 、 devicetree 、 リキッド 、 アセンブリ 、

hyprlang、gn、blade、cfml、cfscript、csv、bibtex、beancount、gitattributes、markdown、toml、xml、scss、perl、fennel、starlark、ron、dotenv、properties、po、diff、rst、mermaid、regex、linkerscript。リリースごとに土地が増えます。
ベータ: semantic-search — デバイス上の EmbeddingGemma 推論は確実で、モデルはバイナリ内に組み込まれます。グラフ コマンドよりも新しいため、まだベータ版とラベル付けされています。
まだ本番環境には対応していません。記録システムではなく、高速コード ナビゲーション補助として使用してください。
MIT — 「ライセンス」を参照してください。サードパーティの通知: THIRD_PARTY.md 。
標準の grep に加えて、コーディング エージェントがコードをナビゲートするために使用できるいくつかのコマンド (誰が呼び出すか、影響するか、コンテキスト、概要など) を追加します。約 2 倍高速で、約 3 ～ 4 倍安価なエージェント コード ナビゲーション。 1 つのネイティブ Rust バイナリ。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
4
v0.1.2
最新の
2026 年 7 月 9 日
+ 3 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Standard grep, plus a few commands your coding agent can use to navigate code — who-calls, impact, context, brief. ~2x faster and ~3-4x cheaper agentic code navigation. One native Rust binary. - metric-space-ai/greppy

GitHub - metric-space-ai/greppy: Standard grep, plus a few commands your coding agent can use to navigate code — who-calls, impact, context, brief. ~2x faster and ~3-4x cheaper agentic code navigation. One native Rust binary. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
metric-space-ai
/
greppy
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
12 Commits 12 Commits .github/ workflows .github/ workflows assets assets bench bench crates crates docs/ assets docs/ assets .gitattributes .gitattributes .gitignore .gitignore ADDING_A_LANGUAGE.md ADDING_A_LANGUAGE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md THIRD_PARTY.md THIRD_PARTY.md rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
Standard grep , plus a few commands your coding agent uses to navigate code — who-calls , impact , semantic-search , brief . On structural code-navigation questions the agent answers correctly ~87% of the time instead of ~53% with plain grep — using fewer tokens. One native Rust binary.
greppy is a drop-in grep — every flag works exactly as before — that also
answers the questions an agent normally burns rounds on: who calls this
function, what breaks if I change it, where is the code that does X. One line in
your agent's config (below) tells it the extra commands exist, and it stops
looping through text matches.
# Standard grep — every command works, unchanged:
greppy -rn " TODO " src/
greppy -i " connection refused " server.log
# A few extra commands, on the same binary:
greppy who-calls parse_config # who calls this function
greppy impact User --direction incoming # what breaks if I change User
greppy semantic-search " restrict a value to a range " # find code by meaning
greppy brief _split_blueprint_path # definition + callers + callees
The same coding agent (MiniMax-M3, driven by Pi Code ) answers one who-calls question on a real repo — left with plain grep , right with greppy . greppy resolves the callers in a single greppy who-calls call instead of a grep-and-read spiral: 2.3× faster, 14 → 5 tool calls, ~9× fewer input tokens . Counters are live from the recorded run.
cargo build --release --bin greppy --features embedded-model
sudo install -m 0755 target/release/greppy /usr/local/bin/greppy
Everything is automatic — the code graph and the semantic model are built into
the binary and build themselves on first use. Nothing to index, nothing to
download, no flags to configure. (Prebuilt binaries for macOS / Linux / Windows
are on the Releases page.) Want it as a transparent grep
drop-in too? Install it a second time as grep .
2. Tell your agent the extra commands exist. Delegate it — in your agent's
chat, say install https://github.com/metric-space-ai/greppy/ — or
paste the snippet below into the file your agent reads for project instructions
( CLAUDE.md , AGENTS.md , .cursor/rules , .windsurfrules , or the system
prompt).
This project has `greppy` — standard grep plus a few code-navigation commands
over a prebuilt symbol graph and an on-device semantic index. Every normal grep
invocation (and flag) works exactly as usual.
CODE-NAVIGATION COMMANDS. SYMBOL is a function / method / class / type name.
They return resolved results as `qualified_name file:line`, not text matches:
greppy who-calls SYMBOL the callers of SYMBOL (incoming calls)
greppy callees SYMBOL the functions SYMBOL calls (outgoing calls)
greppy find-usages SYMBOL every reference to SYMBOL (calls, uses, imports)
greppy brief SYMBOL SYMBOL's definition plus its callers and callees, in one call
greppy impact SYMBOL the transitive set of code a change to SYMBOL reaches
greppy search-symbols NAME definitions whose name matches NAME (a name or fragment)
greppy path --from A --to B a call chain from symbol A to symbol B, if one exists
SEMANTIC SEARCH — use when you do NOT know the symbol's name:
greppy semantic-search "PLAIN-ENGLISH DESCRIPTION"
Describe the behaviour or code you are looking for in plain English
(e.g. "restrict a value to a range", "retry a failed HTTP request").
Returns the closest-matching definitions by meaning (signature + file:line).
EXPAND — get the full source in one call instead of opening files by hand:
greppy expand ID
who-calls / callees / impact / semantic-search may end their output with a
line `Expand: greppy expand <id>`. Run it to print the prepared evidence
pack — the full source of the top matches, bundled — in a single call,
instead of reading each file:line yourself.
FLAGS (append to any command above):
--code include each result's source lines (so no separate read is needed)
--all return every result (turn off the default truncation)
--json machine-readable output with exact counts
--root DIR run against a repo other than the current directory
--kind KIND (search-symbols) restrict to function|method|class|struct|enum|trait
--direction incoming|outgoing, --depth N (impact) which way and how
[truncated]
What an agent actually pays for is billed tokens and wall-clock time.
The benchmark: 14 coding-agent models (Claude Opus/Sonnet/Fable, GPT-5.5, Gemini, Grok, DeepSeek, Qwen, GLM, Kimi, MiniMax-M3, …), each driven by Pi Code , answer 35 code-navigation questions across 7 real repositories (Rust serde + tokio , Python flask + django , Java gson , TypeScript zod , Go hugo ). Every task runs twice — once with plain grep , once with greppy , same agent, same prompt. Answers are floor-graded : a pass must name the ground-truth symbol/file (each anchor rg-verified at generation time). The harness is in bench/agent_efficiency/ .
Correctness is the headline. On structural navigation — who-calls , callees , impact/blast-radius , find-symbol — the agent answers correctly 87% of the time with greppy vs 53% with plain grep (graded by the repo's own grade_answers.py ). Across all 35 questions: 90% vs 63%. Plain grep is cheap but frequently confidently wrong; greppy resolves the relationship in one call.
So it is not a cost-for-accuracy trade: on structural questions greppy is both more correct and cheaper.
Where plain grep keeps up: open-ended "how does this subsystem work" questions. Both tools reach the answer there (~98% correct either way), but greppy's precise locator makes the agent read more to explain the mechanism , so it costs a little more . greppy's edge is pinpoint / structural questions — the semantic path is being tuned to also lead the agent to the answer in one step.
The gain depends on the model. Priced at each model's real OpenRouter list rate, the actual dollar cost of the structural tasks drops with greppy for most models — a median of ~16% cheaper — but it swings widely (from +54% on MiMo to −84% on Grok 4.3, which spirals) and does not track a model's general agentic-benchmark score. Benchmark your own model — most come out ahead, and every model gets the correctness lift.
Real cost = each model's OpenRouter list price × tokens, summed over the structural tasks. The scatter shows the efficiency gain is real for most models but genuinely model-dependent, not something a model's agentic rank predicts.
Standard grep. Any invocation that isn't one of the extra commands runs real grep and returns its output and exit code unchanged.
A precomputed code graph. An indexed, typed symbol graph ( CALLS / USES / TYPE_REF / IMPORTS ) answers who-calls / callees / find-usages / impact / path directly — resolved relationships with file:line , not text matches — collapsing several grep+read rounds into one call.
Native semantic search. For a natural-language query that shares no words with the code, semantic-search embeds the query with Google's EmbeddingGemma on greppy's own native Rust inference (CPU / Apple Metal / NVIDIA CUDA, auto-detected — no llama.cpp, no Python, no HTTP) and returns the nearest code spans by meaning. A small warm daemon keeps the model resident between calls and drops it after idle, so it never holds GPU memory while you're not searching.
One native Rust binary. The EmbeddingGemma model is baked into the binary; tree-sitter parsers and SQLite are compiled in statically.
Early and evolving — the drop-in grep core is solid; the intelligence layers around it are beta.
Solid: the grep drop-in and the code-graph commands ( who-calls / callees / find-usages / impact / path / brief ) on supported languages.
Supported languages (107): python , csharp , go , cpp , php , rust , swift , scala , c , java , javascript , typescript , ruby , bash , kotlin , fsharp , julia , ocaml , d , gdscript , zig , elm , erlang , crystal , gleam , objc , solidity , prisma , protobuf , css , dockerfile , json , groovy , lua , sql , make , nix , cmake , dart , fortran , elixir , scheme , vue , astro , svelte , verilog , glsl , hcl , matlab , r , purescript , racket , clojure , haskell , cuda , tcl , graphql , pascal , powershell , html , yaml , hlsl , cobol , fish , ini , vhdl , json5 , awk , cairo , ada , hare , kdl , jsonnet , llvm , janet , jinja2 , bicep , gotemplate , just , devicetree , liquid , assembly , hyprlang , gn , blade , cfml , cfscript , csv , bibtex , beancount , gitattributes , markdown , toml , xml , scss , perl , fennel , starlark , ron , dotenv , properties , po , diff , rst , mermaid , regex , linkerscript . More land in each release.
Beta: semantic-search — the on-device EmbeddingGemma inference is solid and the model ships inside the binary. Newer than the graph commands, so still labelled beta.
Not yet production-ready — use it as a fast code-navigation aid, not a system of record.
MIT — see LICENSE . Third-party notices: THIRD_PARTY.md .
Standard grep, plus a few commands your coding agent can use to navigate code — who-calls, impact, context, brief. ~2x faster and ~3-4x cheaper agentic code navigation. One native Rust binary.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
4
v0.1.2
Latest
Jul 9, 2026
+ 3 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
