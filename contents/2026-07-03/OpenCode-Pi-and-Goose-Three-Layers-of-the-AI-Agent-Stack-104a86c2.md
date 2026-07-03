---
source: "https://gist.github.com/AIMOWAY/bd8007c8f834a9bc83c71e3178239d75"
hn_url: "https://news.ycombinator.com/item?id=48779685"
title: "OpenCode, Pi, and Goose: Three Layers of the AI Agent Stack"
article_title: "OpenCode, Pi, and Goose: Three Layers of the AI Agent Stack · GitHub"
author: "AIMOWAY"
captured_at: "2026-07-03T20:53:01Z"
capture_tool: "hn-digest"
hn_id: 48779685
score: 1
comments: 0
posted_at: "2026-07-03T20:38:28Z"
tags:
  - hacker-news
  - translated
---

# OpenCode, Pi, and Goose: Three Layers of the AI Agent Stack

- HN: [48779685](https://news.ycombinator.com/item?id=48779685)
- Source: [gist.github.com](https://gist.github.com/AIMOWAY/bd8007c8f834a9bc83c71e3178239d75)
- Score: 1
- Comments: 0
- Posted: 2026-07-03T20:38:28Z

## Translation

タイトル: OpenCode、Pi、Goose: AI エージェント スタックの 3 層
記事のタイトル: OpenCode、Pi、Goose: AI エージェント スタックの 3 層 · GitHub
説明: OpenCode、Pi、Goose: AI エージェント スタックの 3 層 - opencode-pi-goose-three-layers.md

記事本文:
OpenCode、Pi、Goose: AI エージェント スタックの 3 層 · GitHub
コンテンツにスキップ
-->
要点の検索
要点の検索
すべての要点
GitHub に戻る
サインイン
サインアップ
サインイン
サインアップ
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
コード、メモ、スニペットを即座に共有します。
AIMOWAY / opencode-pi-goose-three-layers.md
要点オプションを表示
ZIPをダウンロード
スター
0
( 0 )
Gist にスターを付けるにはサインインする必要があります
フォーク
0
( 0 )
Gist をフォークするにはサインインする必要があります
埋め込む
この要点を Web サイトに埋め込みます。
シェアする
この要点の共有可能なリンクをコピーします。
HTTPS経由でクローンを作成する
Web URL を使用してクローンを作成します。
<script src="https://gist.github.com/AIMOWAY/bd8007c8f834a9bc83c71e3178239d75.js"></script> でこのリポジトリのクローンを作成します。
AIMOWAY/bd8007c8f834a9bc83c71e3178239d75 をコンピューターに保存し、GitHub デスクトップで使用します。
コード
改訂
1
埋め込む
オプションを選択してください
埋め込む
この要点を Web サイトに埋め込みます。
シェアする
この要点の共有可能なリンクをコピーします。
HTTPS経由でクローンを作成する
Web URL を使用してクローンを作成します。
<script src="https://gist.github.com/AIMOWAY/bd8007c8f834a9bc83c71e3178239d75.js"></script> でこのリポジトリのクローンを作成します。
AIMOWAY/bd8007c8f834a9bc83c71e3178239d75 をコンピューターに保存し、GitHub デスクトップで使用します。
ZIPをダウンロード
OpenCode、Pi、Goose: AI エージェント スタックの 3 つの層
生
opencode-pi-goose-three-layers.md
OpenCode、Pi、Goose: AI エージェント スタックの 3 つの層
現在、オープンソースの AI エージェント ツールは数多くありますが、それらがすべてスタックの同じレイヤーに存在するとは思えません。
私が最近注目しているプロジェクトは次の 3 つです。
オープンコード: https://github.com/anomalyco/opencode
ピ: https://g

ithub.com/earendil-works/pi
グース: https://github.com/aaif-goose/goose
これらは「AI コーディング エージェント」としてグループ化されがちですが、そのラベルには重要な違いが隠されています。私の現在のメンタルモデルは次のとおりです。
Pi は、エージェント カーネル、エージェント ハーネス、またはエージェント ツールキットに最も近いものです。
Goose は、開発者向けのローカル AI エージェント ワークベンチおよびオーケストレーション サーフェスに最も近いものです。
OpenCode はコーディングファーストのソフトウェア開発エージェントに最も近いです。
これはランキングではありません。レイヤーの区別です。
1. Pi: エージェント カーネル、ハーネス、およびツールキット
Pi は、単一のエンドユーザーのコーディング アシスタントというよりは、エージェント システムを構築、研究、拡張するための基盤のように感じられます。
このリポジトリは、それ自体を Pi エージェント ハーネス プロジェクトとして説明しており、次のようなパッケージが含まれています。
インタラクティブなコーディング エージェント CLI
ツール呼び出しと状態管理を備えたエージェント ランタイム
統合されたマルチプロバイダー LLM API
その構造が重要なのです。 Pi は、直接実行できるものとしてだけでなく、モデル呼び出し、ツール呼び出し、状態、プロバイダーの抽象化、端末の対話、ワークフロー制御など、エージェント システムの可動部分を理解する場所としても役立ちます。
円周率について考える実際的な方法:
エージェントがどのように構築されるかを理解したい場合に使用してください。
独自のエージェント ツールの下位レベルのベースが必要な場合に使用します。
エージェントのランタイム設計、プロバイダーの抽象化、再現可能なワークフローを重視する場合に使用してください。
完成したアシスタントだけを使用するのではなく、エージェント システムの構造を試してみたい場合に使用してください。
重要な詳細の 1 つは、その許可モデルです。 Pi には、ファイル システム、プロセス、ネットワーク、または資格情報へのアクセスを制限するための組み込みの許可システムが含まれていないことが明示されています。デフォルトでは、それを起動したユーザーとプロセスの権限で実行されます。より強力な境界が必要な場合は、pr する必要があります。

コンテナ化またはサンドボックス化について説明します。
それは必ずしも弱点ではありません。ある意味、これは明確なアーキテクチャ上のステートメントです。Pi はハーネスです。セキュリティ境界はその周囲の環境に属します。
したがって、私は Pi を主に「コーディング エージェント」として説明するつもりはありません。これは、コーディング エージェント CLI がたまたま含まれているエージェント ハーネスおよびツールキットとして説明できます。
2. Goose: ローカル エージェント ワークベンチとオーケストレーション サーフェス
Goose はコーディングよりも幅広いと感じます。
このプロジェクトは、デスクトップ アプリ、CLI、API を備えたネイティブのオープンソース AI エージェントとして自体を説明しています。コード、ワークフロー、およびその間のあらゆるものを対象としています。現在、Linux Foundation の Agentic AI Foundation の一部でもあります。
Goose は「AI プログラミング アシスタント」という狭い概念を超えて進んでいるように見えるため、その制度的背景は興味深いものです。これは開発者向けのエージェント ワークベンチに似ています。モデル、ツール、ファイル、ターミナル ワークフロー、API、拡張機能を統合できるローカル システムです。
グースについて考える実際的な方法:
ローカル AI エージェント ワークベンチが必要な場合に使用します。
デスクトップおよび CLI サーフェスが必要な場合に使用します。
プロバイダーの柔軟性が必要な場合に使用してください。
MCP スタイルの拡張ワークフローが必要な場合に使用します。
タスクがコードの編集に限定されない場合に使用してください。
Goose はコーディングには便利かもしれませんが、私はそれをコーディングに限定するつもりはありません。その自然な範囲には、調査、執筆、自動化、データ分析、ローカル開発者のワークフローが含まれるようです。
Pi がハーネス層に近いとすれば、Goose はオーケストレーション サーフェス、つまりエージェントが開発者にとって使用可能なローカル ツールになる場所に近いです。
3. OpenCode: コーディングファーストのソフトウェア開発エージェント
OpenCode は 3 つの中で最も焦点を当てているようです。
その位置づけは直接的で、オープンソースの AI コーディング エージェントです。ターミナル UI、デスクトップ アプリ、およびビルドが備わっています。

フルアクセスのビルド エージェントや読み取り専用のプランニング エージェントなどのエージェント モード。
これにより、そのレイヤーが理解しやすくなります。 OpenCode は、リポジトリ内でのソフトウェア開発作業を目的としています。
ソフトウェアプロジェクトに取り組む
OpenCode について考える実際的な方法:
主なタスクがソフトウェア開発である場合に使用してください。
リポジトリの近くにエージェントを配置したい場合に使用します。
計画と実装のワークフローを重視する場合に使用してください。
一般的なエージェント プラットフォームではなく、コーディング優先のインターフェイスが必要な場合に使用します。
OpenCode はより広範なエージェント機能を備えていますが、その重心はソフトウェア エンジニアリングです。
Pi と比べると、完成したコーディング エージェント製品のように感じられます。 Goose と比較すると、一般的なローカル ワークベンチというよりも、特化したソフトウェア開発ハーネスのように感じられます。
円周率
エージェントカーネル/ハーネス/ツールキット
ガチョウ
ローカル AI エージェント ワークベンチ / オーケストレーション サーフェス
オープンコード
コーディングファーストのソフトウェア開発エージェント
フレームを構成する別の方法:
エージェントを構築したい場合は、Pi を勉強してください。
ローカル エージェントのワークフローを整理したい場合は、Goose について勉強してください。
日常的なコーディング作業をエージェントに依頼したい場合は、OpenCode について勉強してください。
もちろん、これらのカテゴリは重複します。 Pi にはコーディング エージェントが含まれています。 Goose はコーディングに使用できます。 OpenCode はエージェント ワークフローを実行できます。しかし、彼らのデザインセンターは違うように感じます。
構成とプロバイダーに関する質問
実際に使用するには、比較は位置決めにとどまるべきではありません。運用上の重要な質問は次のとおりです。
プロバイダーの設定は簡単ですか?
カスタムの OpenAI 互換ベース URL をサポートしますか?
API キーとローカル認証情報はどのように管理されますか?
プロバイダー間でのモデル名はどのように処理されますか?
ストリーミングを確実にサポートしていますか?
ツール呼び出しはどのように処理されますか?
どのようにして許可プロンプトを公開するのでしょうか?
コンテナ内で実行できますか、それとも

サンドボックス？
スクリプト化または埋め込みは可能ですか?
複数ステップのタスクが失敗した場合、どのように動作しますか?
これらの質問は、ベンチマーク形式の比較よりも重要です。実際の使用では、エージェントの信頼性はモデルの品質だけではありません。また、コンテキストの処理、ツールの設計、権限の境界、再試行、障害からの回復、およびエージェントの推論とアクションがユーザーにどの程度見えるかということも含まれます。
それぞれに手を伸ばすとき
私は現在、目標がエージェント システム自体の学習または構築である場合、Pi に手を伸ばします。
エージェントループの実験
プロバイダー抽象化の設計
外部サンドボックスを使用してエージェントを構成する
目標がより広範なローカル エージェント環境である場合、私は Goose に手を伸ばすでしょう。
コーディングタスクと非コーディングタスクにわたって AI を使用する
MCP スタイルの拡張機能を介してツールを接続する
デスクトップとターミナルのワークフローを組み合わせる
ローカル開発者ワークベンチの構築
研究、執筆、データ、またはワークフロータスクの自動化
リポジトリ内でコーディングすることが目標の場合、私は OpenCode に手を伸ばします。
なじみのないコードベースを探索する
テストまたは開発コマンドの実行
コーディングパートナーとしてエージェントを使用する
私が最も役立つと思う区別は、「どちらが優れているか」ということではありません。
エージェント スタックのどの層で作業しようとしていますか?
エージェント アーキテクチャ層で作業している場合、Pi は特に興味深いものです。
ローカル ワークフロー オーケストレーション層で作業している場合、Goose は特に興味深いものです。
ソフトウェア開発タスク層で作業している場合、OpenCode は特に興味深いものです。
この層の区別により、エコシステムが理解しやすくなります。また、異なる問題を解決しようとしているツールを比較することを避けるのにも役立ちます。
私が今でも最も重要だと思う質問は、実践的なものです。
日々の人間工学に最適なツールはどれですか?
最も明確なアクセス許可モデルを持つものはどれですか?
どっちが効くんだろう

カスタムの OpenAI 互換エンドポイントを使用していませんか?
複数のプロバイダーを使用して構成するのが最も簡単なのはどれですか?
長い複数ステップのタスクで最も信頼できるのはどれですか?
最も透過的に失敗するのはどれですか?
フレームワークと競合せずに拡張するのが最も簡単なのはどれですか?
おそらく、答えは README よりも実際のワークフローの使用に依存します。
私の現時点での結論はシンプルです。
OpenCode、Pi、Goose は、同じものの 3 つのバージョンではありません。これらは、オープンソース AI エージェント スタックの 3 つの異なる層として理解されるとよいでしょう。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

OpenCode, Pi, and Goose: Three Layers of the AI Agent Stack - opencode-pi-goose-three-layers.md

OpenCode, Pi, and Goose: Three Layers of the AI Agent Stack · GitHub
Skip to content
-->
Search Gists
Search Gists
All gists
Back to GitHub
Sign in
Sign up
Sign in
Sign up
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Instantly share code, notes, and snippets.
AIMOWAY / opencode-pi-goose-three-layers.md
Show Gist options
Download ZIP
Star
0
( 0 )
You must be signed in to star a gist
Fork
0
( 0 )
You must be signed in to fork a gist
Embed
Embed this gist in your website.
Share
Copy sharable link for this gist.
Clone via HTTPS
Clone using the web URL.
Clone this repository at &lt;script src=&quot;https://gist.github.com/AIMOWAY/bd8007c8f834a9bc83c71e3178239d75.js&quot;&gt;&lt;/script&gt;
Save AIMOWAY/bd8007c8f834a9bc83c71e3178239d75 to your computer and use it in GitHub Desktop.
Code
Revisions
1
Embed
Select an option
Embed
Embed this gist in your website.
Share
Copy sharable link for this gist.
Clone via HTTPS
Clone using the web URL.
Clone this repository at &lt;script src=&quot;https://gist.github.com/AIMOWAY/bd8007c8f834a9bc83c71e3178239d75.js&quot;&gt;&lt;/script&gt;
Save AIMOWAY/bd8007c8f834a9bc83c71e3178239d75 to your computer and use it in GitHub Desktop.
Download ZIP
OpenCode, Pi, and Goose: Three Layers of the AI Agent Stack
Raw
opencode-pi-goose-three-layers.md
OpenCode, Pi, and Goose: Three Layers of the AI Agent Stack
There are many open-source AI agent tools now, but I do not think they all live at the same layer of the stack.
Three projects I have been looking at recently are:
OpenCode: https://github.com/anomalyco/opencode
Pi: https://github.com/earendil-works/pi
Goose: https://github.com/aaif-goose/goose
They are easy to group together as “AI coding agents,” but that label hides an important distinction. My current mental model is this:
Pi is closest to an agent kernel, agent harness, or agent toolkit.
Goose is closest to a local AI agent workbench and orchestration surface for developers.
OpenCode is closest to a coding-first software-development agent.
This is not a ranking. It is a layer distinction.
1. Pi: Agent Kernel, Harness, and Toolkit
Pi feels less like a single end-user coding assistant and more like a foundation for building, studying, and extending agent systems.
The repository describes itself as the Pi agent harness project and includes packages such as:
an interactive coding-agent CLI
an agent runtime with tool calling and state management
a unified multi-provider LLM API
That structure matters. Pi is useful not only as something you can run directly, but also as a place to understand the moving parts of an agent system: model calls, tool calls, state, provider abstraction, terminal interaction, and workflow control.
A practical way to think about Pi:
Use it if you want to understand how agents are built.
Use it if you want a lower-level base for your own agent tools.
Use it if you care about agent runtime design, provider abstraction, and reproducible workflows.
Use it if you want to experiment with the structure of agent systems rather than only use a finished assistant.
One important detail is its permission model. Pi is explicit that it does not include a built-in permission system for restricting filesystem, process, network, or credential access. By default, it runs with the permissions of the user and process that launched it. If you need stronger boundaries, you need to provide containerization or sandboxing.
That is not necessarily a weakness. In some ways it is a clear architectural statement: Pi is the harness; the security boundary belongs to the environment around it.
So I would not primarily describe Pi as “a coding agent.” I would describe it as an agent harness and toolkit that happens to include a coding-agent CLI.
2. Goose: Local Agent Workbench and Orchestration Surface
Goose feels broader than coding.
The project describes itself as a native open-source AI agent with a desktop app, CLI, and API. It is positioned for code, workflows, and everything in between. It is also now part of the Agentic AI Foundation at the Linux Foundation.
That institutional context is interesting because Goose seems to be moving beyond the narrower idea of “AI programming assistant.” It looks more like a developer-facing agent workbench: a local system where models, tools, files, terminal workflows, APIs, and extensions can be brought together.
A practical way to think about Goose:
Use it if you want a local AI agent workbench.
Use it if you want desktop and CLI surfaces.
Use it if you want provider flexibility.
Use it if you want MCP-style extension workflows.
Use it if your tasks are not limited to editing code.
Goose may be useful for coding, but I would not reduce it to coding. Its natural scope seems to include research, writing, automation, data analysis, and local developer workflows.
If Pi is closer to the harness layer, Goose is closer to the orchestration surface: the place where an agent becomes a usable local tool for a developer.
3. OpenCode: Coding-First Software-Development Agent
OpenCode seems the most focused of the three.
Its positioning is direct: an open-source AI coding agent. It has a terminal UI, a desktop app, and built-in agent modes such as a full-access build agent and a read-only planning agent.
That makes its layer easier to understand. OpenCode is aimed at software development work inside repositories:
working through a software project
A practical way to think about OpenCode:
Use it if your main task is software development.
Use it if you want an agent that lives close to the repository.
Use it if you care about planning and implementation workflows.
Use it if you want a coding-first interface rather than a general agent platform.
OpenCode may have broader agent capabilities, but its center of gravity is software engineering.
Compared with Pi, it feels more like a finished coding-agent product. Compared with Goose, it feels less like a general local workbench and more like a specialized software-development harness.
Pi
Agent kernel / harness / toolkit
Goose
Local AI agent workbench / orchestration surface
OpenCode
Coding-first software-development agent
Another way to frame it:
If you want to build agents, study Pi.
If you want to organize local agent workflows, study Goose.
If you want an agent for day-to-day coding work, study OpenCode.
Of course, these categories overlap. Pi includes a coding agent. Goose can be used for coding. OpenCode can perform agentic workflows. But their design centers feel different.
Configuration and Provider Questions
For practical use, the comparison should not stop at positioning. The important operational questions are:
How easy is it to configure providers?
Does it support custom OpenAI-compatible base URLs?
How does it manage API keys and local credentials?
How does it handle model names across providers?
Does it support streaming reliably?
How does it handle tool calls?
How does it expose permission prompts?
Can it run inside a container or sandbox?
Can it be scripted or embedded?
How does it behave when a multi-step task goes wrong?
These questions matter more than benchmark-style comparisons. In real use, agent reliability is not just model quality. It is also context handling, tool design, permission boundaries, retries, failure recovery, and how visible the agent’s reasoning and actions are to the user.
When I Would Reach for Each One
I would currently reach for Pi when the goal is learning or building the agent system itself.
experimenting with agent loops
designing provider abstraction
composing an agent with external sandboxing
I would reach for Goose when the goal is a broader local agent environment.
using AI across coding and non-coding tasks
connecting tools through MCP-style extensions
combining desktop and terminal workflows
building a local developer workbench
automating research, writing, data, or workflow tasks
I would reach for OpenCode when the goal is coding inside a repository.
exploring an unfamiliar codebase
running tests or development commands
using an agent as a coding partner
The distinction I find most useful is not “which is better?”
Which layer of the agent stack are you trying to work at?
If you are working at the agent architecture layer, Pi is especially interesting.
If you are working at the local workflow orchestration layer, Goose is especially interesting.
If you are working at the software-development task layer, OpenCode is especially interesting.
That layer distinction makes the ecosystem easier to understand. It also helps avoid comparing tools that are trying to solve different problems.
The questions I still find most important are practical ones:
Which tool has the best day-to-day ergonomics?
Which one has the clearest permissions model?
Which one works best with custom OpenAI-compatible endpoints?
Which one is easiest to configure with multiple providers?
Which one is most reliable in long multi-step tasks?
Which one fails most transparently?
Which one is easiest to extend without fighting the framework?
The answers probably depend less on the README and more on real workflow use.
My current conclusion is simple:
OpenCode, Pi, and Goose are not three versions of the same thing. They are better understood as three different layers of the open-source AI agent stack.
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
