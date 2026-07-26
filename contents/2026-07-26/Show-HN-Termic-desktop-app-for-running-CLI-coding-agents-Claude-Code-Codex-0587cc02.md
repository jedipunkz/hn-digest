---
source: "https://termic.dev/"
hn_url: "https://news.ycombinator.com/item?id=49056316"
title: "Show HN: Termic – desktop app for running CLI coding agents (Claude Code, Codex)"
article_title: "Run claude, codex, antigravity and more in parallel, each in its own git worktree"
author: "SimionBaws"
captured_at: "2026-07-26T10:29:21Z"
capture_tool: "hn-digest"
hn_id: 49056316
score: 1
comments: 0
posted_at: "2026-07-26T09:36:49Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Termic – desktop app for running CLI coding agents (Claude Code, Codex)

- HN: [49056316](https://news.ycombinator.com/item?id=49056316)
- Source: [termic.dev](https://termic.dev/)
- Score: 1
- Comments: 0
- Posted: 2026-07-26T09:36:49Z

## Translation

タイトル: Show HN: Termic – CLI コーディング エージェント (Claude Code、Codex) を実行するためのデスクトップ アプリ
記事のタイトル: クロード、コーデックス、反重力などをそれぞれ独自の git ワークツリーで並行して実行する
説明: オープンソースの Conductor.build の代替品。すでに料金を支払っている Claude Pro / Max プランで、claude、codex、antigravity などを並列 git worktree で実行します。
HN テキスト: つまり、https://termic.dev クロード コード / コーデックス / agt / などのコーディング セッションを管理するためのオープン ソース デスクトップ アプリです。メインチェックアウトまたはワークツリー。話は次のとおりです。数か月前、Anthropic は、Claude コード (エージェント SDK + -p を含む) のプログラムによる使用に関する価格変更を発表しました。私は、Conductor のようなツールが動作しなくなるのではないかと疑っていました。自分自身でクローンの構築を開始しました。これは、カスタム エージェント SDK ではなく、PTY 端末の CLI バイナリを第一級市民として使用します。
これは、CLI エージェントの最新の改善点を常に最新の状態に保つための最も確実な方法であり、サードパーティ ツールの価格設定の問題が二度と発生することがなくなり、どの CLI エージェントでも柔軟に使用できるようになります。こうして https://termic.dev が立ち上げられました。また、Termic で対処した conductor の問題点もいくつかあります。メイン チェックアウトでの起動 (ワークツリーへの強制参加はありません)、エージェントのサンドボックス化 (ファイル システム + ネットワーク)、マルチリポジトリ プロジェクトとエージェント、および独自のエージェント CLI の持ち込みです。スタックはシンプルで、私はパフォーマンスと UX に夢中になっています。Rust/Tauri2、react、WebGL レンダラーを備えた xterm、さらには大量のパフォーマンスの調整が行われています。本当に難しかったこと: すべてのコーディング エージェントの自動再開セッション + 信頼性の高い作業完了検出 (最も難しいもの)。このプロジェクトは着々と進められており、私には 4 ～ 5 人のコントリビューターがいて、デザインについて話し合い、フィードバックを提供するなどしています。ところで、オープンソース AGPL についてです。収益化の計画はなく、オープンサワーを使用したのでコミュニティに還元するだけです

私のキャリア全体を通して、私自身を大切にしてきました。フィードバックもお待ちしております。
https://termic.dev
https://github.com/simion/termic

記事本文:
termic .dev ドキュメント ブログ 変更履歴 GitHub ダウンロード
クロードを走らせて、
コーデックス 、
アジー +3 もっと見る
それぞれが独自の git worktree 内にあります。
新機能はすべて出荷当日に提供されます。
Termic は PTY 内に実際の CLI を生成します。
クロード、
コーデックス 、
反重力 、
副操縦士、
グロクと
オープンコード
船に組み込まれています。それらのいずれかが更新されると、更新を取得できます。ラッパーが追いつく必要はありません。
30 秒以内に独自のものを追加します。
補助者、
オラマ・ラン、
PTY で実行されるものすべて。
無料のオープンソース (AGPL-3.0)、ラッパーの代わりに実際の CLI を実行します。
テルミックは無料です。 CLI は既存のプランで実行されます。
Termic は、既存の CLI バイナリ (claude、gemini、codex) を生成し、その認証を継承します。 Anthropic の 2026 年 6 月 15 日のエージェント SDK クレジットの変更は、SDK パス (および claude -p、GitHub Actions 統合、SDK を通じて認証するサードパーティ アプリ) に適用される予定です。 Termic が生成する対話型の claude CLI は、通常の Pro / Max サブスクリプションの使用制限内に留まります。
Termic 自体には費用はかかりませんし、今後もかかりません。ソースは GitHub にあります。フォークし、監査し、パッケージ化する - 派生製品は AGPL のままです。これが、次の「オープン コア」ツールが密かに独自仕様になるのを防ぐのです。
Termic は、ターミナルで使用するのと同じパスで、インタラクティブなクロードを直接生成します。 SDK クレジット メーターには表示されません。トークンごとのマークアップ、非表示のプロキシ、およびベンダー CLI 間のカスタム エージェント ループはありません。
最良のハーネスは独自のハーネスです
Anthropic / Google / OpenAI は、最初に CLI で機能を出荷します。そこにチームが作業を入れます。 SDK ラッパーがそれを追いかけます。 Termic は CLI を実行するだけなので、モデルがアップグレードされるか、スラッシュ コマンドが到着したその日から、それを使用することになります。
メイン チェックアウト、または新しいワークツリー
メイン チェックアウトでは、エージェントが実際のチェックアウトにアタッチされます -
1 回限りの質問については、README の編集など、判断に値しないような小さなものについては、

h.
これがデフォルトであり、ほとんどの人がここから開始します。
Worktree はディスク上に別のコピーを作成し、分岐します
デフォルトをオフにします (ブランチ名が自動的に入力され、インラインで編集可能です)。開発サーバーを実行します。
一意のポート、機能を出荷し、完了したらアーカイブします。
Termic から価値を引き出すために並行作業は必要ありません。たとえ
単一のブランチ上の単一のプロジェクトでは、アプリではタブ付きエージェントが提供されます。
組み込みエディターと HEAD との差分ビュー、Run/Setup パネル
開発サーバーをストリーミングし、ワンクリックで claude --resume /
コーデックス再開 -- タスクごとの最後。それは
すでに実行している CLI の日常シェルを改善します。
検索、マルチリポジトリ、ターミナル、ブロードキャスト
料金計算により、Termic をインストールすることになります。毎日触れるものです。これらはすべて単一のリポジトリと単一のブランチで動作し、並列作業は必要ありません。
api/、web/、および infra/ を 1 つのタスクに追加すると、エージェントはそれらすべてに対して同時に動作します。すべてのリポジトリとその構造を記述した共有 CLAUDE.md をドロップすると、セッションごとにレイアウトを再説明することなくエンドツーエンドで開発できます。
Cmd+P は、git ls-files をサポートする Sublime スタイルのファジー ファインダーを開きます。追跡されたファイルのみを参照するため、大規模なリポジトリではインスタントに留まり、node_modules に迷い込むことはありません。
git grep によるプロジェクト全体の検索。すでに git によってインデックスが作成されており、.gitignore を尊重し、大規模なリポジトリでも高速に動作します。個別のインデクサーをインストールしたり、ベビーシッターをしたりする必要はありません。
/ ターミナル タブ、起動コマンド
すべてのタブにエージェントが必要なわけではありません。通常のターミナル、または設定したコマンドを直接起動するターミナルを開きます。サイドバーでクラスターごとに k9s --context= を維持すると、すべての環境が 1 クリックで表示されます。
1 つのタスク内の複数のエージェントに 1 つのメッセージを一度に送信します。クロード、ジェミニ、コーデックスで同じ質問をし、アプローチに取り組む前にそれぞれの答えを比較します。
/ .termic.yaml、リポジトリにコミット

スクリプト、実行コマンド、プレビュー URL、およびサンドボックス許可リストは、リポジトリの .termic.yaml に存在します。チームメイトがクローンを作成して Termic を開くと、セットアップがすでに存在しています。構成を手動で再構築する人は誰もいません。
エージェントがターンを終了した瞬間に、青い点がタブに着地します。入力がブロックされるとオレンジ色のベルが鳴ります。オプションのデスクトップ通知を使用して、正確なタスクとタブを表示することで、無駄な推測ではなく実際の PTY 信号を読み取ります。
エージェントに対して次のいくつかのメッセージを並べると、前のターンが終了したときにそれぞれが自動的に送信されます。ステップを N 回繰り返すか、計画全体をキューに入れて終了します。
一日中実行したプロンプト (レビュー、テストの作成、コミット) を保存し、1 つのメニューから実行中のエージェントまたは新しいエージェントに実行します。ブランチの変更を独自に検出する diff 対応のスターター セットを同梱します。
プルリクエストのようにエージェントの作業をレビューします
エージェントはあなたに差分を渡します。正確なフィードバックを返します。
間違っている行を選択し、インライン コメントをそのまま残してください
問題がどこにあるのかを確認して、読み続けてください。メモは次のように積み重なっていきます
ドラフトは差分に固定されます - あなたがそうするまで何も送信されません。
完了したら、バッチを送信します
すべてのコメントを 1 つのメッセージに - それぞれのファイル、行範囲を含む
コードの引用符を入力し、エージェントに直接ドロップします。いいえ
「224 行目のもの」を再入力しても、場所を失うことはありません。コードです
ただし、作成者がコメントをコミットに変換する場合を除きます。
ブランチには触れずにスタック全体を実行します
Spotlight は 1 つのタスクをメインのチェックアウトにミラーリングするため、実際の開発者は
サーバー、ウォッチャー、テストはエージェントの変更に対してライブで実行されます。あ
バックグラウンドウォッチャーは数秒以内に再同期し、
Run は、独自のタブのメイン チェックアウトで実行されます。
安全な部分: Termic はブランチにコミットを書き込むことはありません。
メインチェックアウトでタスクをチェックアウトします

分離された HEAD として扱われるため、ブランチ参照は決して移動しません。
クリーンアップが実行されない場合、Conductor のチェックポイント コミットによりブランチに残ったコミットが取り残される可能性があります
衝突の後; Spotlight は毎回、コミットを残しません。
ワークフローではなくエージェントをケージに入れる
オプションのタスクごとの macOS シートベルト + インプロセス HTTPS CONNECT プロキシ。タスクの作成時に固定され、CLI が生成された瞬間から適用されます。インストールするデーモンやプロンプトごとの承認シアターはありません。
すべてのアウトバウンドリクエストは、正規表現のホスト名ホワイトリストを持つインプロセス Rust CONNECT プロキシを通過します。 CLI ごとのベンダー API (anthropic / google / openai)、GitHub、npm、PyPI、crates.io、および CA OCSP が組み込まれています。プロジェクトごとに独自のものを追加します。 tinyproxy もサイドカーもありません。プロキシは Tauri バイナリ内に存在します。
macOS シートベルトによるデフォルトの拒否。タスク、CLI 自体のキャッシュ、および必要なツールチェーンのみがアクセス可能です。それ以外はすべてブロックされ、読み取りと書き込みが行われます。したがって、エージェントが --dangerously-skip-permissions を実行している場合でも、~/.ssh、~/.aws、~/.gnupg、~/.netrc、~/.kube、キーチェーン、ブラウザ プロファイル、メール、メッセージ、および個人フォルダにはアクセスできません。
ケージ内では、エージェントの許可プロンプトは自動的にスキップされます。シートベルトのプロファイルが境界です。檻の外では、意図的な危険信号としてツールバーの稲妻アイコンが赤く反転します。タスクがどのモードにあるかを常に把握できます。
Tauri シェル、Web フロントエンド。エディターとターミナルは、WKWebView で代替のベンチを作成した後に選択した既製のライブラリーです。
Rust バックエンド、WKWebView フロントエンド。 ~10MB バンドル (Electron の ~120)。
UI。開発用には Vite HMR、製品用には単一のバンドル。
エディタ。 ～150KB;テストしたとき、WKWebView では Monaco の方が遅かったです。
ターミナル。 TUI アプリで行ギャップが表示されないのは WebGL レンダラーだけでした。
macOS / Linux / Windows 上の PTY。 Wezterm 自体が使用しているのと同じクレートです。
州。リデューサーもサンクもありません

- フックだけ。
これは私の Claude Pro / Max サブスクリプションで機能しますか?
これは Conductor.build とどう違うのですか?
Termic は私のプロンプトを認識しますか、または私のコードを読み取りますか?
エージェントが私のシークレットに触れたり、ランダムな API にアクセスしたりすることはできますか?
同じエージェントに二度支払うのはやめましょう。
無料、AGPL-3.0。現在は macOS + Linux。もうすぐウィンドウズ。
オープンソース AI ツールのサポート
Termic は無料で、AGPL-3.0 であり、1 人によって構築されています。チームが AI コーディング エージェントをベースにして構築しており、それが役立つと感じている場合は、スポンサーになることでチームを継続的に進めることができます。

## Original Extract

Open-source Conductor.build alternative. Run claude, codex, antigravity and more in parallel git worktrees, on the Claude Pro / Max plan you already pay for.

So, https://termic.dev open source desktop app for managing claude code / codex / agt / etc coding sessions. main checkout or worktrees. Here's the story: a couple of months ago Anthropic announced pricing changes over the programatic use of Claude code (Agent SDK + including -p), i suspected tools like Conductor will stop working. Started building a clone myself, which uses the CLI binaries in PTY terminals as first class citizens, not custom agent sdk.
This is the most bulletproof way of always being up to date with latest improvements of the cli agents + and never have this pricing issue again for 3rd party tools, and you get the flexibility of using any CLI agent. That's how https://termic.dev was launched. Also has some pain points with conductor which i adressed in termic: launch in main checkout (doesnt force you into a worktree), sandboxing of agents (filesystem + network), multi-repo projects and agents, and bring-your-own agent CLI. Stack is simple, and i;'ve been obsessing over performance and UX: Rust/Tauri2, react, xterm with WebGL renderer, plius a ton of performance tweaks. What was really difficult: auto resume sessions for all coding agents + reliable work done detection (the hardest one). The project has been dteadily gworing, i have 4-5 contributors and we dicuss design, provide feedback, etc. BTW open source AGPL. No plans of monetizing, just giving back to the community as i used open source myself the throught my entire career. Would love your feedback too.
https://termic.dev
https://github.com/simion/termic

termic .dev Docs Blog Changelog GitHub Download
Run claude ,
codex ,
agy +3 more
each in its own git worktree .
Every new feature the day they ship it.
Termic spawns the real CLIs in PTYs.
claude ,
codex ,
antigravity ,
copilot ,
grok and
opencode
ship built in; when any of them updates, you get the update, no wrapper to catch up.
Add your own in 30 seconds :
aider ,
ollama run ,
anything that runs in a PTY.
Free, open source (AGPL-3.0), running the real CLI instead of a wrapper.
Termic is free. CLIs run on your existing plan.
Termic spawns the CLI binaries you already have - claude, gemini, codex - and inherits their auth. Anthropic's June 15, 2026 Agent SDK credit change is scheduled to apply to the SDK path (plus claude -p, the GitHub Actions integration, and third-party apps that authenticate through the SDK). The interactive claude CLI - the one Termic spawns - stays on the regular Pro / Max subscription usage limits.
Termic itself costs nothing and never will. Source is on GitHub. Fork it, audit it, package it - derivatives stay AGPL, which is what keeps the next 'open core' tool from quietly going proprietary.
Termic spawns interactive claude directly, the same path you'd use in a terminal. It is not on the SDK credit meter. No per-token markup, no hidden proxy, no custom agent loop between you and the vendor CLI.
The best harness is their own harness
Anthropic / Google / OpenAI ship features in their CLIs first - that's where their teams put the work. SDK wrappers chase it. Termic just runs the CLI, so the day a model upgrades or a slash-command lands, you're using it.
Main checkout, or a fresh worktree
Main checkout attaches an agent to your actual checkout -
for one-off questions, README edits, the kind of small thing that doesn't deserve a branch.
It's the default, and where most people start.
Worktree creates a separate copy on disk, branched
off your default (the branch name is filled in for you, editable inline). Run a dev server on a
unique port, ship a feature, archive when done.
You don't need parallel work to get value out of Termic. Even with a
single project on a single branch, the app gives you tabbed agents,
a built-in editor and diff view against HEAD, a Run/Setup panel that
streams your dev server, and one-click claude --resume /
codex resume --last per task. It's a
better day-to-day shell for the CLI you're already running.
Search, multi-repo, terminals, broadcast
The billing math gets you to install Termic. This is the stuff you touch every day. All of it works on a single repo and a single branch, no parallel work required.
Add api/, web/ and infra/ to one task and the agent works across all of them at once. Drop a shared CLAUDE.md describing every repo and its structure, and you develop end to end without re-explaining the layout each session.
Cmd+P opens a Sublime-style fuzzy finder backed by git ls-files. It only sees tracked files, so it stays instant on big repos and never wanders into node_modules.
Project-wide search via git grep. Already indexed by git, respects your .gitignore, stays fast on large repositories. No separate indexer to install or babysit.
/ Terminal tabs, your launch command
Not every tab needs an agent. Open a plain terminal, or one that boots straight into a command you set. Keep k9s --context= per cluster in the sidebar and every environment is one click away.
Send one message to several agents in a task at once. Ask the same question across claude, gemini and codex, then compare how each answers before you commit to an approach.
/ .termic.yaml, committed to the repo
Scripts, run commands, preview URL and the sandbox allowlist live in a .termic.yaml in your repo. A teammate clones, opens Termic, and the setup is already there. Nobody rebuilds your config by hand.
A blue dot lands on a tab the moment its agent finishes a turn; an orange bell when one is blocked on your input. Read from real PTY signals, not idle guessing, with an optional desktop notification that drops you on the exact task and tab.
Line up the next few messages for an agent and each sends automatically when it finishes the previous turn. Repeat a step N times, or queue a whole plan and walk away.
Save the prompts you fire all day - review, write tests, commit - and run them from one menu into a running agent or a fresh one. Ships a diff-aware starter set that finds the branch's changes on its own.
Review your agent's work like a pull request
The agent hands you a diff; you hand it back precise feedback.
Select the lines that are wrong, leave an inline comment right
where the problem is, and keep reading. Your notes stack up as
drafts pinned to the diff - nothing is sent until you say so.
When you're done, one Send batches
every comment into a single message - each with its file, line range
and a quote of the code - and drops it straight into the agent. No
retyping "the thing on line 224", no losing your place. It's code
review, except the author turns your comments into commits.
Run your whole stack, without touching your branch
Spotlight mirrors one task into your main checkout, so your real dev
server, watchers and tests run against the agent's changes, live. A
background watcher re-syncs within seconds, and
Run executes at the main checkout in its own tab.
The safe part: Termic never writes a commit onto your branch .
It checks the task out at the main checkout as a detached HEAD, so your branch ref never moves.
Conductor's checkpoint commits can strand a leftover commit on your branch if cleanup doesn't run
after a crash; Spotlight leaves zero commits behind, every time.
Cage the agent, not the workflow
Optional per-task macOS Seatbelt + an in-process HTTPS CONNECT proxy. Pinned at task creation, enforced from the moment the CLI spawns - no daemon to install, no per-prompt approval theater.
Every outbound request goes through an in-process Rust CONNECT proxy with a regex hostname allowlist. Per-CLI vendor APIs (anthropic / google / openai), GitHub, npm, PyPI, crates.io and CA OCSP are baked in; add your own per project. No tinyproxy, no sidecar - the proxy lives inside the Tauri binary.
Default-deny via macOS Seatbelt. Only the task, the CLI's own caches and the toolchains it needs are reachable; everything else is blocked, read and write. So ~/.ssh, ~/.aws, ~/.gnupg, ~/.netrc, ~/.kube, Keychains, browser profiles, Mail, Messages and your personal folders are unreachable, even if the agent runs --dangerously-skip-permissions.
Inside the cage, the agent's permission prompts are auto-skipped - the seatbelt profile IS the boundary. Outside the cage, the toolbar lightning icon flips red as a deliberate danger signal. You always know which mode the task is in.
Tauri shell, web frontend. Editor and terminal are off-the-shelf libraries we picked after benching alternatives in WKWebView.
Rust backend, WKWebView frontend. ~10MB bundle (vs Electron's ~120).
UI. Vite HMR for dev, single bundle for prod.
Editor. ~150KB; Monaco was slower in WKWebView when we tested.
Terminal. WebGL renderer was the only one without visible row gaps in TUI apps.
PTYs on macOS / Linux / Windows. Same crate Wezterm itself uses.
State. No reducers, no thunks - just hooks.
Does this work with my Claude Pro / Max subscription?
How is this different from Conductor.build?
Does Termic see my prompts or read my code?
Can agents touch my secrets or hit random APIs?
Stop paying twice for the same agent.
Free, AGPL-3.0. macOS + Linux now; Windows soon.
Support open-source AI tooling
Termic is free, AGPL-3.0, and built by one person. If your team builds on AI coding agents and finds it useful, sponsoring helps keep it moving.
