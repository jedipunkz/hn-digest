---
source: "https://nexo.sh/posts/agentpack/"
hn_url: "https://news.ycombinator.com/item?id=48345976"
title: "Show HN: Agentpack – isolated config layers for Claude Code, Codex, and OpenCode"
article_title: "I Wanted Reproducible AI Coding Environments, So I Built agentpack | Oleg Pustovit – Technical Leadership & MVP Development for Startups"
author: "nexo-v1"
captured_at: "2026-06-01T00:29:40Z"
capture_tool: "hn-digest"
hn_id: 48345976
score: 4
comments: 0
posted_at: "2026-05-31T14:32:55Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Agentpack – isolated config layers for Claude Code, Codex, and OpenCode

- HN: [48345976](https://news.ycombinator.com/item?id=48345976)
- Source: [nexo.sh](https://nexo.sh/posts/agentpack/)
- Score: 4
- Comments: 0
- Posted: 2026-05-31T14:32:55Z

## Translation

タイトル: Show HN: Agentpack – クロード コード、コーデックス、および OpenCode 用の分離された構成レイヤー
記事タイトル: 再現可能な AI コーディング環境が欲しかったので、agentpack を構築しました | Oleg Pustovit – スタートアップ向けの技術的リーダーシップおよび MVP 開発
説明: Claude Code、Codex、Cursor、OpenCode などの最新のエージェントにはスキル、フック、MCP が必要ですが、それらのロード方法は異なります。 Agentpack は、再現可能な AI コーディング環境用の一時的なステージング構成を作成します。

記事本文:
[オレグ・プストビット] /work
/書き込み
/contact 再現可能な AI コーディング環境が欲しかったので、エージェントパックを構築しました
Claude Code、Codex、Cursor、OpenCode などの最新のエージェントにはスキル、フック、MCP が必要ですが、それらのロード方法は異なります。グローバル ユーザーやプロジェクトの構成を汚さずにコーディング エージェントに読み込まれる一時的なステージング構成を作成するツールを作成しました。 GitHub リンク
端末の記録を読み込んでいます...私は商用ソフトウェアの構築に AI エージェントを使用してきました。これは、迅速なプロトタイプを作成したり、初期段階のスタートアップ向けに製品を反復したりするのに非常に役立ちます。データ分析、生産上の問題の優先順位付け、さらにはプロジェクトで使用されているプログラミング言語の変更など、以前は多大な労力が必要だった決定の取り消しさえも、本質的に安価になりました。以前は、ニーズに合わせて CLI/スクリプトを作成するには数日の作業が必要でした。今では数分で構築できるため、自分にとって意味のある問題に集中できるようになります。
ハーネスが新たなボトルネックになっている
もちろん、自律コーディングのワークフローが完璧になるまでにはまだ長い道のりがあります。これらのシステムは、高品質のコードを確実に生成できません。フロンティア モデルであっても、適切にプロンプ​​トを表示しない限り、カスタム抽象化を作成したり、保守可能なコードを生成したりすることはできません。状況をさらに悪化させるのは、ソフトウェア エンジニアリングのあらゆる問題に優れたツールが 1 つも存在しないことです。その代わりに、本質的に同じエージェント ループを実装する競合するエージェント ハーネスが多数あります。Claude Code、OpenCode、Codex、Copilot、Cursor などです。これらのエージェント ハーネスでは、特定のモデルのセットに制限されており、エージェントを柔軟に構成するための機能が欠けていることがよくあります。さらに、これらの CLI エージェントはコンテキストを異なる方法で管理するため、コーディングのパフォーマンスが異なることがよくあります。

同じモデルの場合。構成は部分的にのみ標準化されています。多くのハーネスでは、AGENTS.md が採用されています。AGENTS.md は、当初 OpenAI によって作成され、Codex、OpenCode、Cursor などによってサポートされています。スキルは多かれ少なかれ SKILL.md 規約に準拠していますが、標準には程遠いです。フック、カスタム ルール、プラグインはまだ同じように統合されていません。 Claude Code と Cursor にもプラグイン システムがありますが、完全に互換性があるわけではありません。 MCP は広く採用される標準となりました。ただし、MCP のプロジェクト固有の構成ファイルは、依然として CLI ツールごとに異なります。スタートアップのチームと協力するとき、Claude Code 用に完全に最適化されたリポジトリをよく見かけますが、Codex でリポジトリを使用するには、シンボリックリンクを使用してハッキングするか、独自のハーネス固有の構成をチームのリポジトリにコミットする必要があります。
すべてのエージェントには独自の構成が必要です
今日に至るまで、私は毎日のドライバーとして Claude Code を使用しています。私は Opus モデルの欠陥を知っています。Opus モデルは、平均的なコードを作成し、重要なことを見逃し、疑わしい決定を下すことが多く、自動操縦での実行が不可能になります。速度を上げるために、私はエージェント CLI を通じて Composer モデルで Cursor Agent を使用します。これにより、迅速なリファクタリング、マージ競合の解決、その他の簡単な作業が完了します。最新の GPT モデル (2026 年 5 月現在) を備えた Codex は、他のモデルが通常失敗する場合に備えて、Playwright MCP と組み合わせて、コンピューターを使用する困難なタスクを処理します。 Gemini 3.5 Flash は最近、Antigravity CLI とともに登場しました。ハーネスは非常に粗くて欠けていますが、モデルは重要な小さな問題に関しては非常に賢く見えます。したがって、3 ～ 4 つの CLI エージェントが同じコードベースで動作するという問題が発生し、それらすべてで一連のスキルとスラッシュ コマンドを共有したいと考えています。
Claude Code のようなハーネスには独自のプラグインがあります

切り替え可能なプラグインを備えたマーケットプレイス。これは、 /plugin コマンドを使用してインストールできるスキル、エージェント、MCP または LSP サーバーを含むパッケージを検出するのに役立ちますが、これらは Claude に固有であり、 ~/.claude の下で構成を変更すると、再現可能な読み取り専用ホーム ディレクトリのモデルと競合する Nix Home-Manager のような宣言的または不変の環境にはあまり適していません。このような場合、AI エージェントを構成するための一時的なステージング レイヤーを使用する方が合理的です。
Agentpack: エージェントの一時的な構成レイヤー
Agentpack が導入される前は、私が作業していたリポジトリは最終的に次のようなものになりました。
.claude/コマンド
.クロード/スキル
.cursor/ルール
.cursor/エージェント
.opencode/コマンド
.エージェント/スキル
# コピーされた MCP 構成スニペット
# 後で削除するのを忘れたシンボリックリンク .claude/commands
.クロード/スキル
.cursor/ルール
.cursor/エージェント
.opencode/コマンド
.エージェント/スキル
# コピーされた MCP 構成スニペット
# シンボリックリンク 後で削除するのを忘れました これは個人プロジェクトでは機能しますが、AI ツールに各自の好みがある可能性がある大規模なチームでは拡張できません。リポジトリには、最終的に自動的にコンテキストにロードされるスキルやその他の成果物が蓄積されます。私が取り組んだフルスタックのモノリポジトリでは、特に置き換えられた計画スキルを使用して、リクエストごとに非常にトークンを大量に消費する検証パイプラインをトリガーする、厳重にガードレールされたエージェント ハーネスがありました。セットアップには 50 以上のカスタム スキルが含まれており、クロード コードで 1 つずつ手動でオフにする必要がありました。これらのオーバーライドは、もちろん、現在ロードされているプロジェクト ディレクトリに固有のものでした。 Agentpack を構築する動機の 1 つは、さまざまな複雑なエージェント ワークフローを分離した方法で管理することでした。 Agentpack を使用すると、信頼できる情報源が 1 つのマニフェスト ファイルになります。
名前 = "私のプロジェクト"
バージョン = "0.1.0"
[依存

密度]
"github.com/anthropics/skills/skills/canvas-design" = { ブランチ = "メイン" }
"github.com/my-org/agent-configs/skills/incident-triage" = { タグ = "v0.3.1" }
[モード.デフォルト]
ベース = 「すべて」
[モード.フロントエンド]
ベース = 「すべて」
disable = ["package:github.com/my-org/agent-configs/skills/incident-triage"]
[モード.インシデント]
ベース = 「すべて」
無効化 = ["mcp:figma"] 名前 = "私のプロジェクト"
バージョン = "0.1.0"
[ 依存関係 ]
"github.com/anthropics/skills/skills/canvas-design" = { ブランチ = "メイン" }
"github.com/my-org/agent-configs/skills/incident-triage" = { タグ = "v0.3.1" }
[ モード .デフォルト]
ベース = 「すべて」
[ モード .フロントエンド]
ベース = 「すべて」
disable = [ "package:github.com/my-org/agent-configs/skills/incident-triage" ]
[ モード .事件】
ベース = 「すべて」
disable = [ "mcp:figma" ] 私のコンサルティング業務では、もちろん、未知のツールのあいまいなマニフェスト ファイルをチームのリポジトリにコミットしたくないので、親の <project-name>-workspace/<project-repo> ディレクトリ構造を作成し、ワークスペース ディレクトリでエージェントパック構成を初期化します。 Agentpack sync を実行すると、すべての依存関係がリロードされ、必要なリポジトリが GitHub から取得されます。フェッチされたリポジトリはユーザー全体のホーム ディレクトリにキャッシュされ、owner/repo/path/commit の繰り返し読み取り時にそこから取得されるため、同じリポジトリを複数回フェッチする必要がなくなります。
エージェントのパッケージ管理の問題を解決しようとするツールはすでに存在します。 Microsoft APM はマニフェストとロックファイルを使用し、再現性と、同じ構成をさまざまなエージェント ハーネスの形式にコンパイルできることに重点を置いています。ルーラーやルールシンクなどのツールは、複数のエージェント ツール間で同じルールと構成のセットを同期して維持することに基づいています。
これは便利ですが、ほとんどのツールは生成されたファイルを p に書き込むことになります。

roject ディレクトリまたはユーザーのホーム ディレクトリ。これは避けたかったものです。そのため、agentpack はスキル、MCP 構成をダウンロードして生成し、必要な形式に移植するだけでなく、それらのファイルを一時的な場所にステージングし、その生成された構成ディレクトリに対してエージェントを起動します。プロジェクト リポジトリはクリーンなままですが、エージェントには必要なスキル、MCP、フック、エージェント、ルールのセットが事前に構成されています。簡単な比較表は次のとおりです。
Agentpack の仕組みを次の図に示します。
パッケージ ソースとしての Git は新しい概念ではありません。この考え方は Go ではすでにおなじみになっています。モジュールはバージョン管理されたリポジトリからダウンロードできます。 APM もソースとして Git を使用します。 Agentpack では、パッケージはローカル参照または GitHub から解決されます。 GitHub の場合、コミットは Pack.lock に固定され、リポジトリはローカルにキャッシュされ、その後ステージングされたディレクトリがダウンロードされたパッケージから必要なファイルをレンダリングします。
私は最近 GitHub の稼働時間の問題が頻繁に発生していることを認識しており、それをハードな依存関係として持つことは実際のトレードオフであることを認識していますが、それでも、GitHub からスキル ファイルをロードするのは簡単で簡単だと感じています (Go でパッケージを追加する方法と似ています)。
また、README.md をわざわざチェックする必要はなく、スキルのディレクトリを参照するだけでよい場合もあります。したがって、スキルを指す GitHub URL をコピーするだけで、CLI がそれをステージング ディレクトリに適切にロードします。
# ブラウザ URL — GitHub のファイル ビュー (ツリーまたは BLOB) から直接貼り付けます
エージェントパック追加 https://github.com/anthropics/skills/tree/main/skills/canvas-design
エージェントパック追加 https://github.com/anthropics/skills/blob/main/skills/canvas-design/SKILL.md
エージェントパック追加 https://github.com/anthropics/claude-plugins-official/blob/main/plugins/hookify/.claude

-プラグイン/プラグイン.json
# Go スタイルのモジュール ID、または所有者/リポジトリ[/パス] の省略表現
エージェントパックは github.com/anthropics/skills/skills/canvas-design を追加します
エージェントパックは人間性/スキル/スキル/キャンバスデザインを追加します
Agentpack は人間性/スキル # リポジトリ全体を HEAD に追加します
# @ref を使用してブランチ、タグ、またはコミットを固定します
エージェントパックは anthropics/claude-plugins-official/plugins/hookify@v1.0.0 を追加します
エージェントパックは github.com/acme/monorepo/packages/rules@main を追加します
# ローカルミラー/エイリアス名、またはファイルシステムパス
エージェントパックはキャンバスデザインを追加します
エージェントパックは ./local-rules を追加します
Agentpack sync # ブラウザ URL — GitHub のファイル ビュー (ツリーまたは BLOB) から直接貼り付けます
エージェントパック追加 https://github.com/anthropics/skills/tree/main/skills/canvas-design
エージェントパック追加 https://github.com/anthropics/skills/blob/main/skills/canvas-design/SKILL.md
エージェントパック追加 https://github.com/anthropics/claude-plugins-official/blob/main/plugins/hookify/.claude-plugin/plugin.json
# Go スタイルのモジュール ID、または所有者/リポジトリ[/パス] の省略表現
エージェントパックは github.com/anthropics/skills/skills/canvas-design を追加します
エージェントパックは人間性/スキル/スキル/キャンバスデザインを追加します
エージェントパックは人間像/スキルを追加します # リポジトリ全体を HEAD に追加します
# @ref を使用してブランチ、タグ、またはコミットを固定します
Agentpack add anthropics/claude-plugins-official/plugins/ [email protected]
エージェントパックは github.com/acme/monorepo/packages/rules@main を追加します
# ローカルミラー/エイリアス名、またはファイルシステムパス
エージェントパックはキャンバスデザインを追加します
エージェントパックは ./local-rules を追加します
重要な違いはインストール ターゲットです。GitHub はソースのみですが、`$STAGING` はエージェント CLI が読み取るものです。パッケージがリモート ソースからダウンロードされる場合は、npm ins を実行するのと同じように、それらのソースを信頼する必要があります。
[切り捨てられた]
私にとって重要な設計上の決定は、構成の移植性と再利用性です。エージェントパックのアイデアは決して真似しないことです

プロジェクト ディレクトリ、または ~/.cursor や ~/.claude などのフォルダー内のグローバル ユーザー設定にコピーします。
解決策は、一時的なステージング ディレクトリを作成し、それをエージェント ハーネスの起動引数に渡すか、環境変数を使用して渡すことでした。実際には、これはツールごとの互換レイヤーになります。
Claude コードは、ステージングされたディレクトリをロードするための --plugin-dir 引数を許可するため、サポートが最も簡単です。
Codex は、代替構成ディレクトリを設定するための CODEX_HOME 環境変数をサポートしています。
OpenCode は Codex と同様に、OPENCODE_CONFIG_DIR 変数と OPENCODE_CONFIG 変数をサポートします。
Cursor Agent では、ステージング ディレクトリを適切に渡すために $HOME 変数を変更する必要がありますが、これによりさまざまな追加の複雑さが生じます。
さらにツールを追加すると、これをサポートするのがさらに困難になります。私は通常、エージェントを使用して、バンドルされた JavaScript コードをリバース エンジニアリングしたり、特定のハーネスの制限を回避する方法を理解するためのスモーク テストを作成したりします。例えば。 .mdc 形式のルールは Cursor でのみサポートされているため、Codex や Claude Code などのハーネスにロードされるフォールバック スキルを生成する必要がありました。 Cursor CLI には文書化されていないバグがあり、その動作を逆にすることによってのみ発見できました。

[切り捨てられた]

## Original Extract

Modern agents like Claude Code, Codex, Cursor and OpenCode need skills, hooks and MCPs, but they load them differently. agentpack creates ephemeral staging configuration for reproducible AI coding environments.

[oleg pustovit] /work
/writings
/contact I Wanted Reproducible AI Coding Environments, So I Built agentpack
Modern agents like Claude Code, Codex, Cursor and OpenCode need skills, hooks and MCPs, but they load them differently. I made a tool that creates an ephemeral staging configuration that loads into coding agents without polluting the global user or project configuration. GitHub Link
Loading terminal recording... I’ve been using AI agents for building commercial software—it’s incredibly useful for creating quick prototypes and iterating on products for early-stage startups. Data analysis, triaging production issues, even undoing decisions that previously would take tremendous effort—like changing the programming language your project is written in—became essentially cheap. Before, writing CLIs/scripts custom to your needs required days of work; now they can be built in minutes, and that allows you to focus on problems that are meaningful to you.
The harness is the new bottleneck
Of course, autonomous coding workflows are still a long way from being perfect. These systems cannot reliably produce high-quality code—even frontier models fail at making custom abstractions and generating maintainable code unless you prompt them correctly. What makes the situation worse is that there’s no single tool that is excellent at every software engineering problem ; instead we have a number of competing agent harnesses that implement essentially the same agent loop: Claude Code, OpenCode, Codex, Copilot, Cursor, etc., where you often find yourself restricted to a set of specific models and lacking the features to configure the agent flexibly. On top of that, those CLI agents manage context differently, which often results in different coding performance for the same models . Configuration is only partially standardized: many harnesses adopted AGENTS.md — initially created by OpenAI and supported by Codex, OpenCode, Cursor and others. Skills have more or less landed on the SKILL.md convention , but it is far from being a standard. Hooks, custom rules and plugins haven’t yet converged in the same way. Claude Code and Cursor also have plugin systems, but they are not fully interchangeable. MCPs became a widely adopted standard; however, the project-specific configuration file for MCPs still differs for every CLI tool. When working with teams on startups, I often see a repository fully optimized for Claude Code, while using it with Codex requires hacking with symlinks or committing my own harness-specific configs to the team’s repository .
Every agent needs its own configuration
To this day, I’m using Claude Code as my daily driver. I know the flaws of the Opus model—it often makes average code, misses important things, talks itself into questionable decisions, which makes it infeasible to run on autopilot. For speed, I use Cursor Agent with the Composer model through the agent CLI. It gets the job done for quick refactors, solving merge conflicts, and other light work. Codex with the latest GPT models (as of May 2026) handles challenging computer-use tasks for me, paired with Playwright MCP , for cases where other models usually fail. Gemini 3.5 Flash recently arrived with the Antigravity CLI ; the harness is very rough and lacking, but the model seems very smart on a small set of non-trivial issues. So I end up with the problem of having 3-4 CLI agents working on the same codebase, and I would like to share a set of skills and slash commands across all of them .
Harnesses like Claude Code have their own plugin marketplace with toggleable plugins. This is useful for discovering packages containing skills, agents, MCP or LSP servers, which can be installed via the /plugin command, but those are specific to Claude and aren’t a great fit for declarative or immutable environments like Nix Home-Manager , where mutating config under ~/.claude conflicts with a model of a reproducible, read-only home directory . In such cases having an ephemeral staging layer for configuring AI agents makes more sense.
agentpack: an ephemeral configuration layer for agents
Before agentpack, repositories I worked in ended up with something like this:
.claude/commands
.claude/skills
.cursor/rules
.cursor/agents
.opencode/commands
.agents/skills
# copied MCP config snippets
# symlinks I forgot to remove later .claude/commands
.claude/skills
.cursor/rules
.cursor/agents
.opencode/commands
.agents/skills
# copied MCP config snippets
# symlinks I forgot to remove later That works for a personal project, but doesn’t scale with a larger team where everyone may have their own preferences in AI tools . The repository accumulates skills and other artifacts that automatically end up loaded into the context. On one full-stack monorepo I worked on, we had a heavily guardrailed agentic harness that on every request triggered a very token-hungry validation pipeline, especially with a replaced planning skill. The setup contained 50+ custom skills that had to be turned off manually, one by one, in Claude Code: those overrides were of course specific to the currently loaded project directory. One of the motivations for building agentpack was to manage various complex agent workflows in an isolated way . With agentpack, the source of truth becomes a single manifest file :
name = "my-project"
version = "0.1.0"
[dependencies]
"github.com/anthropics/skills/skills/canvas-design" = { branch = "main" }
"github.com/my-org/agent-configs/skills/incident-triage" = { tag = "v0.3.1" }
[modes.default]
base = "all"
[modes.frontend]
base = "all"
disable = ["package:github.com/my-org/agent-configs/skills/incident-triage"]
[modes.incident]
base = "all"
disable = ["mcp:figma"] name = "my-project"
version = "0.1.0"
[ dependencies ]
"github.com/anthropics/skills/skills/canvas-design" = { branch = "main" }
"github.com/my-org/agent-configs/skills/incident-triage" = { tag = "v0.3.1" }
[ modes . default ]
base = "all"
[ modes . frontend ]
base = "all"
disable = [ "package:github.com/my-org/agent-configs/skills/incident-triage" ]
[ modes . incident ]
base = "all"
disable = [ "mcp:figma" ] In my consulting work, of course, I don’t want to commit obscure manifest files of some unknown tool to teams’ repositories, so I create a parent <project-name>-workspace/<project-repo> directory structure, where I init the agentpack config in the workspace directory. Running agentpack sync will reload all the dependencies and fetch the necessary repositories from GitHub. Fetched repositories are cached in the user-wide home directory and pulled from there on repeated reads of owner/repo/path/commit to eliminate the need to fetch the same repository multiple times.
There are already tools that are trying to solve the problem of package management for agents. Microsoft APM uses a manifest and lockfile and focuses on reproducibility and being able to compile the same configuration to formats of different agent harnesses. Tools like ruler and rulesync are based on keeping the same set of rules and configurations in sync across multiple agent tools.
This is useful, but most tools end up writing generated files into the project directory or user’s home directory, which is something I wanted to avoid . So agentpack not only downloads and generates skills, MCP configs and ports them over to the needed format, it also stages those files in a temporary location and launches the agent against that generated configuration directory . The project repository stays clean , while the agent is preconfigured with the needed set of skills, MCPs, hooks, agents and rules. Here’s a brief comparison table:
The way agentpack works is shown on the following diagram:
Git as a package source is not a new concept . This idea already became familiar with Go : a module can be downloaded from a version-controlled repository. APM uses Git as a source too. In agentpack, packages are resolved either from local references or from GitHub . In the case of GitHub, commits are pinned in pack.lock , repositories are cached locally, and then a staged directory renders the necessary files from the downloaded packages.
While I’m aware of GitHub’s frequent uptime problems lately and realize that having it as a hard dependency is a real tradeoff, I still find loading those skill files from GitHub simple and easy (similar to how we add packages in Go).
Other times, I would rather not bother checking the README.md and instead just navigate a directory of skills. So you can simply copy a GitHub URL pointing to a skill and the CLI will load it properly into the staging directory:
# Browser URLs — paste straight from GitHub's file view (tree or blob)
agentpack add https://github.com/anthropics/skills/tree/main/skills/canvas-design
agentpack add https://github.com/anthropics/skills/blob/main/skills/canvas-design/SKILL.md
agentpack add https://github.com/anthropics/claude-plugins-official/blob/main/plugins/hookify/.claude-plugin/plugin.json
# Go-style module ID, or owner/repo[/path] shorthand
agentpack add github.com/anthropics/skills/skills/canvas-design
agentpack add anthropics/skills/skills/canvas-design
agentpack add anthropics/skills # whole repo at HEAD
# Pin a branch, tag, or commit with @ref
agentpack add anthropics/claude-plugins-official/plugins/hookify@v1.0.0
agentpack add github.com/acme/monorepo/packages/rules@main
# A local mirror / alias name, or a filesystem path
agentpack add canvas-design
agentpack add ./local-rules
agentpack sync # Browser URLs — paste straight from GitHub's file view (tree or blob)
agentpack add https://github.com/anthropics/skills/tree/main/skills/canvas-design
agentpack add https://github.com/anthropics/skills/blob/main/skills/canvas-design/SKILL.md
agentpack add https://github.com/anthropics/claude-plugins-official/blob/main/plugins/hookify/.claude-plugin/plugin.json
# Go-style module ID, or owner/repo[/path] shorthand
agentpack add github.com/anthropics/skills/skills/canvas-design
agentpack add anthropics/skills/skills/canvas-design
agentpack add anthropics/skills # whole repo at HEAD
# Pin a branch, tag, or commit with @ref
agentpack add anthropics/claude-plugins-official/plugins/ [email protected]
agentpack add github.com/acme/monorepo/packages/rules@main
# A local mirror / alias name, or a filesystem path
agentpack add canvas-design
agentpack add ./local-rules
agentpack sync The important distinction is the install target: GitHub is only the source, while `$STAGING` is what the agent CLI reads. When packages are downloaded from remote sources, you need to trust those sources, just like running npm ins
[truncated]
A design decision that was critical for me: portability and reusability of configuration . The agentpack idea is to never copy assets to the project directory or the global user configuration in folders like ~/.cursor or ~/.claude .
The solution was to create an ephemeral staging directory and pass it to the launch arguments of the agentic harness , or use environment variables to do so. Practically speaking, this becomes a compat layer per tool:
Claude Code is the simplest to support as it allows --plugin-dir argument for loading the staged directory.
Codex supports CODEX_HOME environment variable for setting an alternative configuration directory.
OpenCode similarly to Codex supports OPENCODE_CONFIG_DIR and OPENCODE_CONFIG variables.
Cursor Agent requires changing the $HOME variable for cleanly passing the staging directory, but that introduces various additional complexities.
Adding more tools makes this even more challenging to support. I typically use agents to reverse engineer bundled JavaScript code or write smoke tests to understand how to get around the limitations of certain harnesses. E.g. rules in .mdc format are only supported by Cursor, so I had to generate a fallback skill that would load into harnesses like Codex or Claude Code. The Cursor CLI has undocumented bugs that were only possible to find by reversing what it is doing: for some reason, it onl

[truncated]
