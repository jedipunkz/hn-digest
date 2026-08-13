---
source: "https://silkcode.web.app/"
hn_url: "https://news.ycombinator.com/item?id=49290887"
title: "SilkCode – Tired of resets/limits with Claude but better than Codex"
article_title: "Silk Code"
author: "amonte"
captured_at: "2026-08-13T19:51:53Z"
capture_tool: "hn-digest"
hn_id: 49290887
score: 1
comments: 0
posted_at: "2026-08-13T19:36:13Z"
tags:
  - hacker-news
  - translated
---

# SilkCode – Tired of resets/limits with Claude but better than Codex

- HN: [49290887](https://news.ycombinator.com/item?id=49290887)
- Source: [silkcode.web.app](https://silkcode.web.app/)
- Score: 1
- Comments: 0
- Posted: 2026-08-13T19:36:13Z

## Translation

タイトル: SilkCode – クロードのリセット/制限にはうんざりだが、Codex よりは良い
記事タイトル: シルクコード
説明: Silk Code — オープンでモデルに依存しない AI コーディング ハーネス。 DeepSeek、Qwen、Kimi、Cloudflare Workers AI、または任意のローカル モデルを使用して、コードを理解、編集、テスト、出荷します。

記事本文:
オープンでモデルに依存しない AI コーディング ハーネス。あなたを理解する一つの環境
DeepSeek、Qwen、Kimi、GLM を使用して、リポジトリ、コードの編集、テストの実行、Git の管理を行います。
MiniMax、Cloudflare Workers AI、または独自のマシン上で実行される任意のモデル。
コーディング環境は開発者のものです。
AIモデルは置き換え可能です。
$ Silkcode --model auto ~/apps/ridelink
シルクコード v0.1 |モデル: ollam/qwen2.5-coder |モード: 尋ねる | ~/アプリ/ライドリンク
Silk> ログイン リダイレクトのバグを修正し、テストを実行します。
⚙ grep {"パターン": "リダイレクト", "パス": "src/auth"}
⚙ read_file {"パス": "src/auth/session.ts"}
⚙ edit_file {"path": "src/auth/session.ts", "old_string": "…", "new_string": "…"}
?ファイルの変更を許可: src/auth/session.ts [y]es / [n]o / [a]lways: y
⚙ run_tests {}
→ $ npm テスト … 終了コード: 0
session.ts のリダイレクト ターゲットを修正して検証しました。24 個のテストに合格しました。
クラウドまたはローカルの任意のモデル
信頼できるモデルをご持参ください
1 つのプロバイダー層が、ネイティブ Ollama を使用して、すべての OpenAI 互換エンドポイントと通信します。
サポートとローカル モデルの検出。セッション中、またはセッションごとに変更せずにモデルを切り替える
働き方。
破線 = 独自のハードウェアで実行されます。 --model は何でも自動選択します
まず実行中のローカル サーバー、次にキーを持つクラウド プロバイダーがあります。その他のものも搭載
シルクコード モデルを使用して追加するか、Ollama スタイルでローカル モデルをプルします。
シルクコード モデルは qwen2.5-coder をプルします。
一度インストールしてください。どこでも実行できます。
Python 3.10 以降が必要です (macOS は 3.9 を出荷します — 現在の Python を次の場所から取得します)
最初に python.org) と git。 Silk Code を独自の仮想環境に保持し、シェル エイリアスを追加する
そのため、silkcode コマンドは、
環境第一。
Silk Code のインストールまたは更新を推奨
# Silk Code をインストールまたは更新する
if [ -d "$HOME/S

ilkCode/.git" ]; 次に
git -C "$HOME/SilkCode" pull
それ以外の場合
git clone https://github.com/RupertCloud/SilkCode.git "$HOME/SilkCode"
フィ
cd「$HOME/シルクコード」
# 必要に応じて仮想環境を作成します
[ -d .venv ] || python3 -m venv .venv
# Silk Code をインストールまたは更新する
.venv/bin/python -m pip install -e 。
# まだ設定されていない場合はグローバル コマンドを追加します
grep -q 'エイリアス Silkcode=' ~/.zshrc 2>/dev/null || \
echo 'alias Silkcode="$HOME/SilkCode/.venv/bin/silkcode"' >> ~/.zshrc
ソース ~/.zshrc
シルクコード --ヘルプ
このコマンド ブロックは再実行しても安全です。既存のチェックアウトを更新し、その仮想チェックアウトを再利用します。
環境を保護し、重複するエイリアスの追加を回避します。エイリアスは Silk Code を直接呼び出します
その環境内にあるため、.venv を使用するために .venv をアクティブ化する必要はありません。
モデルを接続します - どちら側でも機能します
# クラウド: DeepSeek (または qwen / kimi / glm / minimax / openrouter)
エクスポート DEEPSEEK_API_KEY=sk-your-key
# またはローカル & プライベート: Ollama.com から Ollama をインストールしてから、
シルクコード モデルは qwen2.5-coder をプルします
ツール呼び出しサポートを備えたモデルを選択してください。エージェントはツールを呼び出して動作します。
qwen2.5-coder はローカルで信頼できる選択肢です。
cd ~/プロジェクト/my-project
シルクコード GUI 。 # 127.0.0.1:8377 のブラウザ GUI
Silkcode gui ~/Projects/my-project # またはプロジェクト パスを直接渡します
Silkcode ~/Projects/my-project # ターミナル REPL
Silkcode gui ~/SilkCode は、Silk Code ソース リポジトリ自体を開きます。
別のコードベースで作業するには、代わりにそのプロジェクトのディレクトリを渡します。
次に、「テストを含む小さな Flask API を構築し、テストが確実に実行されることを確認してください」と尋ねます。
パスしてください。」プロンプトが表示されたら、そのアクションを承認します。に切り替える
--mode エージェントを信頼したら。
あなたとモデルの間のすべて
プロジェクト エクスプローラー、ストリーミング チャット、エージェント アクティビティ タイムライン、差分ビューア。オープンパラレル
さまざまなモデルを使用したセッション。忙しいセッションはその間も動作し続けます

別のものを使っています。
すべてのセッションは、リポジトリ マップ、SILKCODE.md プロジェクト ルール、
蓄積されたプロジェクト メモリとインストールされたスキル - モデルはコードベースを認識します
最初のツール呼び出しの前。
コマンドはリスク分類されています。読み取りは無料で実行され、書き込みとインストールは要求されます。
rm -rf / プッシュ / マージには、明示的にしない限り、常に承認が必要です
セッションを許可します。
ファイルは、エージェントがファイルにアクセスする前にスナップショットが作成されます。ワンクリック（または
/revert ) はターン全体を元に戻します。
run_tests は、pytest、npm、cargo、go、flutter を自動検出します — エージェント
スイートを実行し、成功を宣言する前に失敗を読み取ります。
コンテキストの圧縮により、古いツールの出力がトリミングされ、モデルの出力に近づくにつれて方向が変わります。
ウィンドウ — セッションは終了するのではなく継続し、GUI または CLI から再開できます。
任意の Model Context Protocol サーバーに接続します。そのツールはエージェント ツールになり、
他のものと同じように承認ゲート型です。
エージェントのコマンドをマシンの代わりに使い捨てコンテナで実行します。
1 つのコマンドで自己ホストすることも、バンドルされた Worker を介して Cloudflare サンドボックス上でホストすることもできます。
Silkcode ベンチマークは、実際のコーディング タスクをモデルごとにエンドツーエンドで実行します。
ハーネス自体が寄与するものを分離する A/B モードを使用します。
アプリのように承認し、チームのように出荷する
GitHub でサインイン — Silk Code アプリをインストールします。
ブラウザでコードを承認したら完了です。トークンを作成したり貼り付けたりする必要はありません。有効期間が短い認証情報、
選択したリポジトリにスコープが設定され、自動的に更新されます。
そしてハーネスはその機能を示します。エージェントが作成したコミット
Silk Code を共著者として登録します。あなたは著者のままであり、記録は正直のままです。
セッション更新後のログインリダイレクトを修正
共著: Silk Code <agent@silkcode.dev>
X-Silk-Model: ディープシーク/ディープシーク チャット
X-シルクセッション: 42
git log --grep=X-Silk-Model
「エージェントはどのモデルを使用してどのコミットを作成しましたか?」と永遠に答えます。
全体

ターミナルからのハーネス

## Original Extract

Silk Code — an open, model-agnostic AI coding harness. Use DeepSeek, Qwen, Kimi, Cloudflare Workers AI, or any local model to understand, edit, test, and ship code.

An open, model-agnostic AI coding harness. One environment that understands your
repository, edits code, runs your tests, and manages git — with DeepSeek, Qwen, Kimi, GLM,
MiniMax, Cloudflare Workers AI, or any model running on your own machine.
The coding environment belongs to the developer.
The AI model is replaceable.
$ silkcode --model auto ~/apps/ridelink
Silk Code v0.1 | model: ollama/qwen2.5-coder | mode: ask | ~/apps/ridelink
silk> Fix the login redirect bug and run the tests.
⚙ grep {"pattern": "redirect", "path": "src/auth"}
⚙ read_file {"path": "src/auth/session.ts"}
⚙ edit_file {"path": "src/auth/session.ts", "old_string": "…", "new_string": "…"}
? Allow modifying file: src/auth/session.ts [y]es / [n]o / [a]lways: y
⚙ run_tests {}
→ $ npm test … exit code: 0
Fixed the redirect target in session.ts and verified it: 24 tests passed.
Any model, cloud or local
Bring whichever model you trust
One provider layer speaks to every OpenAI-compatible endpoint, with native Ollama
support and local-model discovery. Switch models mid-session — or per session — without changing
how you work.
Dashed = runs on your own hardware. --model auto picks whatever
you have: a running local server first, then any cloud provider with a key. Onboard anything else
with silkcode models add — or pull local models Ollama-style with
silkcode models pull qwen2.5-coder .
Install once. Run it anywhere.
Requires Python 3.10 or newer (macOS ships 3.9 — grab a current Python from
python.org first) and git. Keep Silk Code in its own virtual environment, then add a shell alias
so the silkcode command is available from every project without activating the
environment first.
Install or update Silk Code recommended
# install or update Silk Code
if [ -d "$HOME/SilkCode/.git" ]; then
git -C "$HOME/SilkCode" pull
else
git clone https://github.com/RupertCloud/SilkCode.git "$HOME/SilkCode"
fi
cd "$HOME/SilkCode"
# create the virtual environment if needed
[ -d .venv ] || python3 -m venv .venv
# install or update Silk Code
.venv/bin/python -m pip install -e .
# add the global command if it is not already configured
grep -q 'alias silkcode=' ~/.zshrc 2>/dev/null || \
echo 'alias silkcode="$HOME/SilkCode/.venv/bin/silkcode"' >> ~/.zshrc
source ~/.zshrc
silkcode --help
This command block is safe to rerun: it updates an existing checkout, reuses its virtual
environment, and avoids adding duplicate aliases. The alias invokes Silk Code directly
inside its environment, so you do not need to activate .venv to use it.
Connect a model — either side works
# cloud: DeepSeek (or qwen / kimi / glm / minimax / openrouter)
export DEEPSEEK_API_KEY=sk-your-key
# or local & private: install Ollama from ollama.com, then
silkcode models pull qwen2.5-coder
Pick a model with tool-calling support — the agent works by calling tools.
qwen2.5-coder is the reliable local choice.
cd ~/Projects/my-project
silkcode gui . # browser GUI at 127.0.0.1:8377
silkcode gui ~/Projects/my-project # or pass the project path directly
silkcode ~/Projects/my-project # terminal REPL
silkcode gui ~/SilkCode opens the Silk Code source repository itself.
To work on another codebase, pass that project’s directory instead.
Then just ask: “Build a small Flask API with tests, and make sure the tests
pass.” Approve its actions as prompts appear; switch to
--mode agent once you trust it.
Everything between you and the model
Project explorer, streaming chat, agent activity timeline, diff viewer. Open parallel
sessions with different models; a busy session keeps working while you use another.
Every session starts with a repo map, your SILKCODE.md project rules,
accumulated project memory, and installed skills — the model knows the codebase
before its first tool call.
Commands are risk-classified. Reads run free, writes and installs ask,
rm -rf / push / merge always require approval — unless you explicitly
grant them for the session.
Files are snapshotted before the agent touches them. One click (or
/revert ) undoes a whole turn.
run_tests auto-detects pytest, npm, cargo, go, and flutter — the agent
runs your suite and reads the failures before claiming success.
Context compaction trims old tool output and turns as you approach the model’s
window — sessions roll on instead of dying, and are resumable from GUI or CLI.
Connect any Model Context Protocol server — its tools become agent tools,
approval-gated like everything else.
Run the agent’s commands in a disposable container instead of your machine:
self-hosted with one command, or on Cloudflare Sandboxes via the bundled Worker.
silkcode benchmark runs real coding tasks end-to-end per model —
with an A/B mode that isolates what the harness itself contributes.
Authorize like an app, ship like a team
Sign in with GitHub — install the Silk Code app,
approve a code in your browser, done. No tokens to create or paste; short-lived credentials,
scoped to the repos you chose, refreshed automatically.
And the harness signs its work. Agent-made commits
register Silk Code as co-author — you stay the author, the record stays honest:
Fix login redirect after session refresh
Co-Authored-By: Silk Code <agent@silkcode.dev>
X-Silk-Model: deepseek/deepseek-chat
X-Silk-Session: 42
git log --grep=X-Silk-Model
forever answers “which commits did the agent write, with which model?”
The whole harness from the terminal
