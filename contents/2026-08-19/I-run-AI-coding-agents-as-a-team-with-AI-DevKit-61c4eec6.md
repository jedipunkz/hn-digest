---
source: "https://codeaholicguy.com/2026/08/19/how-i-run-ai-coding-agents-as-a-team-with-ai-devkit/"
hn_url: "https://news.ycombinator.com/item?id=49360515"
title: "I run AI coding agents as a team with AI DevKit"
article_title: "How I run AI coding agents as a team with AI DevKit – Codeaholicguy"
image: "https://i0.wp.com/codeaholicguy.com/wp-content/uploads/2026/08/building_ai_coding_agent_teams.png?fit=1200%2C670&ssl=1"
author: "hoangnnguyen"
captured_at: "2026-08-19T12:23:58Z"
capture_tool: "hn-digest"
hn_id: 49360515
score: 1
comments: 0
posted_at: "2026-08-19T12:17:06Z"
tags:
  - hacker-news
  - translated
---

# I run AI coding agents as a team with AI DevKit

- HN: [49360515](https://news.ycombinator.com/item?id=49360515)
- Source: [codeaholicguy.com](https://codeaholicguy.com/2026/08/19/how-i-run-ai-coding-agents-as-a-team-with-ai-devkit/)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T12:17:06Z

## Translation

タイトル: AI DevKit を使用してチームとして AI コーディング エージェントを実行しています
記事のタイトル: AI DevKit を使用してチームとして AI コーディング エージェントを実行する方法 – Codeaholicguy
説明: 私はほとんどの時間を AI DevKit エージェント コンソールで過ごしています。私はマネージャーとして 1 つのエージェント (通常は Codex) を立ち上げ、それについてブレインストーミングを行います。アイデアが十分に具体的になって実行できるようになったら、私はマネージャーに別のエージェントを作成して仕事を引き継ぐように頼みます。そのエグゼキュータは、Codex、Claude Code、Pi などで実行される可能性があります。

記事本文:
AI DevKit を使用して AI コーディング エージェントをチームとして実行する方法 – Codeaholicguy
コンテンツにスキップ
コードアホリックガイ
AI DevKit を使用してチームとして AI コーディング エージェントを実行する方法
私はほとんどの時間を AI DevKit エージェント コンソールに費やしています。私はマネージャーとして 1 つのエージェント (通常は Codex) を立ち上げ、それについてブレインストーミングを行います。アイデアが十分に具体的になって実行できるようになったら、私はマネージャーに別のエージェントを作成して仕事を引き継ぐように頼みます。そのエグゼキュータは、必要に応じて、Codex、Claude Code、Pi、または別のハーネスで実行される可能性があります。
エグゼキューターは、開発ライフサイクルに従い、変更を実装、検証し、多くの場合、私があまり介入することなくプル リクエストを作成します。私は、エージェントのセッションごとに切り替えるのではなく、マネージャーと進捗状況を確認します。プル リクエストの準備ができたら、それをレビューし、必要に応じてフィードバックを提供し、マージします。
これにより、端末をタブの壁にせずに、同時に複数の作業を行うことができます。
私もいつもキーボードに向かっているわけではありません。離れるときは、マネージャー エージェントを Telegram に接続し、電話から同じセッションで作業を続けます。
AI DevKit は、エージェント マネージャー、共有オペレーティング モデル、およびさまざまなハーネスが 1 つのチームとして機能するための十分な共通コンテキストを提供します。
Claude Code、Codex、および Gemini CLI を並行して実行しています。
最悪の部分はトークンコストではありません。タブ切り替えです。
AI DevKit のエージェント コンソールはこれを修正します。 1 つのビューで、すべてのエージェントが離れることなくメッセージを送信でき、詳細が必要な場合はホットキーで集中できます。
iTerm2 で動作します… pic.twitter.com/3VHg8LIUBa
私にとっての主なインターフェイスは AI DevKit エージェント コンソールです。
コンソールを使用すると、さまざまなハーネスにわたってエージェントを開いてメッセージを送信し、監視するための 1 つの場所が得られます。何が実行されているのか、どのエージェントが待機しているのか、どこに介入する必要があるのかがわかります。
今でもエージェントリストなどのコマンドを使用しています

、エージェント詳細、およびエージェント送信ですが、主にデバッグ、自動化、またはコンソールでは明らかにされないものの検査に使用されます。これらは低レベルのコントロールであり、私のワークフローの中心ではありません。
私は、より高速なディスパッチャーとなるコマンドのコレクションを望んでいませんでした。私は問題に集中し続けながら、1 人のエージェントが他のエージェントを管理できる運用環境を望んでいました。
マネージャーはメッセージをルーティングするだけではありません。目標、アクティブなワークストリーム、エージェント間の依存関係、および決定を待っているすべてのものを追跡します。また、完了を主張するエージェントが実際に検証済みの結果を生成したかどうかもチェックします。
非常に多くのエグゼキューター タイプがある理由についても、前の投稿で説明しました。基本的には、サブスクリプションを最大限に活用しながら、各エグゼキューターの得意分野を活かして、さまざまなモデルを実行できるようにしたいと考えています。
マネージャーが私をフォローして Telegram にアクセスしました
私がコンピュータの前にいるときは、エージェント コンソールが主なインターフェイスです。キーボードから離れているときは、AI DevKit チャネル コネクタを使用します。
コンソールでマネージャー エージェントを選択し、設定した Telegram チャネルに接続します。ブリッジはバックグラウンドで実行されるため、ターミナル ビューを閉じて、携帯電話からマネージャーと会話を続けることができます。また、同じ AI DevKit スタックを VPS 上で実行しているので、エージェント マネージャーと 24 時間年中無休でチャットできます。
AI エージェントが端末内に常駐する必要はもうありません。
AI DevKit で Pi 統合が利用できるようになりました。
Agent Console で Pi セッションを開始するか、Telegram からワンクリックでリモート制御します。 https://t.co/tt6M4nhR3O pic.twitter.com/MRtHq4yhgz
これは依然として同じマネージャー セッションです。 Telegram は単なるトランスポートです。マネージャーは、作業コンテキスト、ファイルシステムとメモリへのアクセス、実行エージェントを調整する機能を維持します。

。
Telegram から、進捗状況を尋ねたり、実行者をブロックしている質問に答えたり、フィードバックを与えたり、マネージャーに別の作業を開始するよう依頼したりできます。オーケストレーション ループを動かし続けるために、デスクに戻るまで待つ必要はありません。
同等の低レベル コマンドは次のとおりです。
ai-devkit チャネル開始テレグラム \ --agent <マネージャー名> \ --daemon
通常はコンソールから接続します。このコマンドは、ブリッジをデバッグするときやセットアップを自動化するときに役立ちます。
ワークフローをスキルにエンコードします
繰り返すことが期待されるものはすべてスキルになります。
これは、さまざまなハーネスの一貫性を保つ方法の大部分を占めています。 Codex、Claude Code、OpenCode、Pi、またはその他のハーネスは機能やセッション動作が異なりますが、リポジトリから同じエンジニアリング プロセスをロードできます。
このワークフローで私が頻繁に使用するスキルは次のとおりです。
エージェントを安全に開始、選択、検査、停止するためのエージェント管理
割り当ての送信、最近のコンテキストの読み取り、エージェント間での結果の中継のためのエージェント通信
繰り返される監視ループのエージェント オーケストレーション: 進行状況の確認、エージェントのブロック解除、依存関係の調整、完了の確認
要件、設計、計画、実装、テスト、レビューを通じて変更を進めるための開発ライフサイクル
エージェントが作業が完了したと言う前に、新しい証拠が必要かどうかを確認する
その行動はスキルにエンコードされているため、マネージャーは管理方法を知っています。実行者は、同じ理由で要件をプル リクエストに反映する方法を知っています。
これにより、毎朝巨大なオーケストレーション プロンプトを作成する必要がなくなり、プロセスが再利用可能になります。明日 Codex を別のハーネスに置き換える場合、エージェント チームの動作方法を再設計したくありません。
ファイルシステムとメモリは共有コンテキストです
新しい時代になると委任が失敗する

nt は頭が空っぽの状態から始まります。
ファイルシステムはこれの一部を解決します。ソース コード、要件、設計上の決定、計画、テスト結果、実装メモは、すべてのエージェントが読み取ることができるファイルに保存されています。結果が明確な要件と計画にすでに取り込まれている場合、実行者はブレインストーミングの記録全体を必要としません。
ただし、ファイルシステムだけでは十分ではありません。一部のコンテキストがコードベースに属していません:
見落としがちな繰り返しのリポジトリ規則
以前のタスク中に行われた決定とその背後にある理由
既知の失敗パターンとそれを回避する検証済みの方法
プロジェクト全体に適用される個人のワークフロー設定
このレイヤーには AI DevKit メモリを使用します。限定的で再利用可能な知識をローカル SQLite に保存します。エージェントは、タスクにコンテキストが必要な場合、すべてのセッションにすべてをロードするのではなく、コンテキストを検索します。
メモリは、異なるハーネスが同じ理解から開始できるようにする重要な層です。 Codex には 1 つの会話履歴があり、Claude Code には別の会話履歴がある可能性がありますが、どちらも同じファイルを読み取り、同じ永続的な知識を取得できます。
コンテキストモデルは次のように考えます。
ファイルシステム: 現在の状態と明示的なプロジェクト成果物
記憶: 永続的な決定、慣例、検証された教訓
セッション: 一時的な推論と現在行われている会話
これらを分離しておくと、メモリがゴミ捨て場になるのを防ぐこともできます。生の会話、一時的な進行状況、推測などは保存しません。将来のエージェントがそれを再利用する可能性が低い場合、そのファイルはメモリに属しません。
私は通常、Codex エージェントをマネージャーとして起動し、そのセッションを維持します。実際の実装が他の場所で行われる場合でも、これが私の仕事のメインインターフェイスになります。
マネージャーは、エージェント管理、エージェント通信、およびエージェント調整のスキルを持っています。実際、私のセットアップでは、これらのスキーは

lls はグローバルにインストールされます。これにより、エージェントの作成、その状態の検査、フォローアップ指示の送信、および複数のパスにわたる作業の監視が可能になります。
私はまだ結果を所有しています。マネージャーは調整ループを所有します。
2. ワーカーを作成する前にブレインストーミングを行う
私はまずマネージャーを思考パートナーとして利用します。私たちは問題を調査し、仮定に異議を唱え、アプローチを比較し、アイデアを絞り込みます。
私はあらゆる思考に対して実行者を生成するわけではありません。それは進歩ではなく活動を生み出すことになります。私は、開発する要件、実行する調査、レビューする設計、または実行の準備ができた計画など、引き渡す価値のあるものがあるまで待ちます。
ここは人間の判断が最も影響力をもつ部分でもあります。エージェントは、綿密に組み立てられたタスクを非常にうまく実行できます。複数の人が自信を持って間違った問題を並行して解決している場合、それらはあまり役に立ちません。
3. マネージャーに適切な執行者を作成するよう依頼します。
方向性が明確になったら、実行してほしい内容をマネージャーに伝えます。時々ハーネスに名前を付けます。また、私が需要を説明し、マネージャーに Codex、Claude Code、Pi、または他の利用可能なエージェントのいずれかを選択させることもあります。
一般的な命令は、長いシェル コマンドよりもこれに近いです。
私たちは要件とアプローチについて合意しました。分離されたワークツリーに実行エージェントを作成し、関連するコンテキストを引き渡し、実装、テスト、レビュー、PR 作成を通じて開発ライフサイクルを継続させます。それを監視して、決定が必要なとき、または PR の準備ができたときを教えてください。
ハンドオフには、目的、範囲、関連ファイル、制約、検証の期待、実行者が生成するアーティファクトが含まれます。共有ファイルシステムとメモリは残りのコンテキストを保持します。
また、マネージャーは、2 人の実行者が同じファイルを同時に編集することを防ぎます。パラレルワークね

境界をクリーンアップするか、ワークツリーを分離します。そうしないと、実装中に節約された時間が競合解決として戻ってきます。
引き継ぎ後は、通常、別のアイデアや別のワークストリームに移ります。
Claude Code と Codex を本物のソフトウェア チームのように協力させました。
前回のビデオでは、1 人の Claude Code エージェントが他のエージェントを調整するという単純なセットアップを示しました。
今回はそれをさらに推し進めてみました。
私は ai-devkit エージェント オーケストレーションを使用して、小さな「チーム」を運営しました。
– リーダー: ソネット… https://t.co/76PA3hvK7C pic.twitter.com/9lziHWn8cX
マネージャーはオーケストレーション ループを実行します。エージェントのステータスをチェックし、待機中または古いものを検査し、必要な情報を中継し、指示を重複させることなく修正を送信します。ある実行者が別の実行者に依存する場合、マネージャーは作業の順序を決定します。
すべての executor セッションを開くのではなく、マネージャーにチェックインします。
現在の進捗状況は何ですか?何かがブロックされているのか、それとも私を待っているのでしょうか？
ほとんどの更新には私の関与は必要ありません。私が気にするのは例外です。製品に関する決定、未解決の競合、繰り返される失敗、セキュリティに配慮したアクション、検証できない結果などです。
ここが、エージェント オーケストレーションがいくつかのバックグラウンド プロセスの起動とは異なる点です。誰か (この場合はマネージャー エージェント) が、作業が完了するか完全にブロックされるまでループを所有します。
5. 必要に応じてエグゼキューターにステップインします。
マネージャーがいる場合でも、コンソールからエグゼキューター セッションを開いて自分自身が関与することもあります。これは、AI DevKit Agent Console でエージェントを選択し、「o」を押すことで簡単に実行できます。
私は通常、エージェントがドリフトしているとき、技術的なトレードオフで素早いやり取りが必要なとき、または行き過ぎる前に実装を制御したいときにこれを行います。直接的な対話は、すべての詳細を中継するよりも多くの帯域幅を持ちます。

マネージャーです。
方向性が再び明確になったら、執行者に続行を任せ、マネージャーに監督を再開させます。マネージャーはデフォルトの調整パスのままですが、直接アクセスが高速な場合には常に利用可能です。
タスクが十分に準備されている場合、多くの場合、実行者は開発ライフサイクルを完了し、独自にプル リクエストを作成できます。
私は今でもエンジニアとしてすべてのプル リクエストをレビューします。元の問題は解決されているか、トレードオフは許容できるか、差分は理解できるか、検証は十分に強力か。
フィードバックがあれば、マネージャーにフィードバックします。マネージャーは、同じ実行者に作業を送り返すことも、別のレビュー パスを作成することもできます。満足したらマージします。
私の役割は一つ上のレベルに上がります。私は組み立て、判断、レビューに多くの時間を費やしますが、エージェントは実行と調整の多くを担当します。
私はツールに永続的な役職を割り当てません。 Codex が常に実装者であるとは限りませんし、Claude Code が常にレビュー者であるとは限りません。
私は作品、すでに利用可能なコンテキスト、使用したいモデル、そのセッションで必要なものに基づいて選択します。新鮮な視点を得るために、別のハーネスが必要になる場合もあります。ツールが変わると選択も変わる可能性があります。
これが共有レイヤーマットのもう 1 つの理由です

[切り捨てられた]

## Original Extract

I spend most of my time in the AI DevKit agent console. I start one agent as the manager, usually Codex, and brainstorm with it. Once an idea becomes concrete enough to execute, I ask the manager to create another agent and hand off the work. That executor might run in Codex, Claude Code, Pi,…

How I run AI coding agents as a team with AI DevKit – Codeaholicguy
Skip to content
Codeaholicguy
How I run AI coding agents as a team with AI DevKit
I spend most of my time in the AI DevKit agent console. I start one agent as the manager, usually Codex, and brainstorm with it. Once an idea becomes concrete enough to execute, I ask the manager to create another agent and hand off the work. That executor might run in Codex, Claude Code, Pi, or another harness depending on what I need.
The executor follows the dev lifecycle, implements the change, verifies it, and often creates the pull request without much intervention from me. I check progress with the manager instead of jumping between every agent session. When the pull request is ready, I review it, give feedback if needed, and merge it.
This is what lets me work on several things at the same time without turning my terminal into a wall of tabs.
I am not always at the keyboard either. When I step away, I connect the manager agent to Telegram and continue working with the same session from my phone.
AI DevKit gives me an agent manager, a shared operating model, and enough common context for different harnesses to work as one team.
I'm running Claude Code, Codex, and Gemini CLI in parallel.
The worst part isn't the token cost. It's the tab switching.
AI DevKit's Agent Console fixes this. One view, all your agents, send messages without leaving, quick hotkey to focus when you need detail.
Works in iTerm2,… pic.twitter.com/3VHg8LIUBa
The main interface for me is the AI DevKit agent console .
The console gives me one place to open, message, and monitor agents across different harnesses. I can see what is running, which agent is waiting, and where I need to step in.
I still use commands such as agent list , agent detail , and agent send , but mostly for debugging, automation, or inspecting something the console does not make obvious. They are the low-level controls, not the center of my workflow.
I did not want a collection of commands that made me a faster dispatcher. I wanted an operating environment where one agent could manage other agents while I stayed focused on the problem.
The manager does more than route messages. It keeps track of the goal, the active workstreams, dependencies between agents, and anything waiting for a decision. It also checks whether an agent claiming completion has actually produced a verified result.
I also mentioned in the previous post why I have so many executor types. Basically, I want to make full use of my subscriptions while also taking advantage of what each executor is good at and being able to run different models.
The manager follows me to Telegram
The agent console is my main interface when I am at the computer. When I am away from the keyboard, I use the AI DevKit channel connector .
I select the manager agent in the console and connect it to my configured Telegram channel. The bridge runs in the background, so I can close the terminal view and keep talking to the manager from my phone. I also have the same AI DevKit stack running on a VPS so that I can chat with the agent manager 24/7.
Your AI agent no longer has to sit inside your terminal.
Pi integration is now available in AI DevKit.
Start a Pi session in Agent Console, or remote control it from Telegram with one click. https://t.co/tt6M4nhR3O pic.twitter.com/MRtHq4yhgz
This is still the same manager session. Telegram is simply another transport into it. The manager keeps its working context, access to the filesystem and memory, and the ability to coordinate executor agents.
From Telegram, I can ask for progress, answer a question that is blocking an executor, give feedback, or ask the manager to start another piece of work. I do not need to wait until I am back at my desk to keep the orchestration loop moving.
The equivalent low-level command is:
ai-devkit channel start telegram \ --agent <manager-name> \ --daemon
I normally connect it from the console. The command is useful when I am debugging the bridge or automating the setup.
I encode the workflow in skills
Anything I expect to repeat becomes a skill .
This is a large part of how I keep different harnesses consistent. Codex, Claude Code, OpenCode, Pi, or any different harnesses have different capabilities and session behavior, but they can load the same engineering process from the repository.
The skills I use heavily for this workflow are:
agent-management for starting, selecting, inspecting, and stopping agents safely
agent-communication for sending assignments, reading recent context, and relaying results between agents
agent-orchestration for the repeated supervision loop: check progress, unblock agents, coordinate dependencies, and verify completion
dev-lifecycle for moving a change through requirements, design, planning, implementation, testing, and review
verify for requiring fresh evidence before an agent says the work is done
The manager knows how to manage because that behavior is encoded in skills. The executor knows how to take a requirement through to a pull request for the same reason.
This saves me from writing a giant orchestration prompt every morning and makes the process reusable. If I replace Codex with another harness tomorrow, I do not want to redesign how my team of agents works.
The filesystem and memory are the shared context
Delegation fails when the new agent starts with an empty head.
The filesystem solves part of this. Source code, requirements, design decisions, plans, test results, and implementation notes live in files that every agent can read. An executor does not need the entire brainstorming transcript if the outcome has already been captured in a clear requirement and plan.
The filesystem is not enough, though. Some context does not belong in the codebase:
a recurring repository convention that is easy to miss
a decision made during an earlier task and the reason behind it
a known failure pattern and the verified way to avoid it
a personal workflow preference that applies across projects
I use AI DevKit memory for this layer. It stores narrow, reusable knowledge in local SQLite. Agents search it when a task needs context instead of loading everything into every session.
Memory is the key layer that lets different harnesses start with the same understanding. Codex may have one conversation history and Claude Code another, but both can read the same files and retrieve the same durable knowledge.
I think of the context model like this:
Filesystem: current state and explicit project artifacts
Memory: durable decisions, conventions, and verified lessons
Session: temporary reasoning and the conversation happening now
Keeping these separate also prevents memory from becoming a dumping ground. I do not store raw conversations, temporary progress, or guesses. If a future agent is unlikely to reuse it, it does not belong in memory.
I usually start a Codex agent as the manager and keep that session alive. It becomes my main interface for the work, even when the actual implementation happens elsewhere.
The manager has the agent-management, agent-communication, and agent-orchestration skills. Actually, in my setup, these skills are installed globally. This gives it the ability to create agents, inspect their state, send follow-up instructions, and supervise work across several passes.
I still own the outcome. The manager owns the coordination loop.
2. Brainstorm before creating workers
I use the manager as a thinking partner first. We explore the problem, challenge assumptions, compare approaches, and narrow the idea.
I do not spawn an executor for every thought. That would create activity, not progress. I wait until there is something worth handing off: a requirement to develop, an investigation to run, a design to review, or a plan ready to execute.
This is also where human judgment has the most leverage. Agents can execute a well-framed task remarkably well. They are much less useful when several of them are confidently solving the wrong problem in parallel.
3. Ask the manager to create the right executor
Once the direction is clear, I tell the manager what I want executed. Sometimes I name the harness. Other times I describe the demand and let the manager choose between Codex, Claude Code, Pi, or another available agent.
A typical instruction is closer to this than a long shell command:
We have agreed on the requirement and approach. Create an executor agent in an isolated worktree, hand over the relevant context, and have it continue the dev lifecycle through implementation, testing, review, and PR creation. Monitor it and tell me when you need a decision or when the PR is ready.
The handoff includes the objective, scope, relevant files, constraints, validation expectations, and the artifact the executor should produce. The shared filesystem and memory carry the rest of the context.
The manager also prevents two executors from editing the same files at the same time. Parallel work needs clean boundaries or separate worktrees. Otherwise, the time saved during implementation comes back as conflict resolution.
After the handoff, I usually move on to another idea or another workstream.
I had Claude Code and Codex working together like a real software team.
In the last video, I showed a simple setup: one Claude Code agent orchestrating other agents.
This time, I pushed it further.
I used ai-devkit agent orchestration to run a small “team”:
– Leader: Sonnet… https://t.co/76PA3hvK7C pic.twitter.com/9lziHWn8cX
The manager runs the orchestration loop. It checks agent status, inspects anything waiting or stale, relays necessary information, and sends corrections without duplicating instructions. If one executor depends on another, the manager sequences the work.
I check in with the manager rather than opening every executor session:
What is the current progress? Is anything blocked or waiting for me?
Most updates do not need my involvement. I care about exceptions: a product decision, an unresolved conflict, repeated failures, a security-sensitive action, or a result that cannot be verified.
This is where agent orchestration becomes different from launching a few background processes. Someone, in this case the manager agent, owns the loop until the work is complete or genuinely blocked.
5. Step into the executor when needed
Even with a manager, I sometimes open the executor session from the console and get involved myself. This can easily be done by choosing the agent in the AI DevKit Agent Console and pressing “o”.
I usually do this when the agent is drifting, a technical trade-off needs a quick back-and-forth, or I want to steer the implementation before it goes too far. Direct interaction has more bandwidth than relaying every detail through the manager.
Once the direction is clear again, I leave the executor to continue and let the manager resume supervision. The manager remains my default coordination path, while direct access is always available when it is faster.
For a well-prepared task, the executor can often complete the dev lifecycle and create the pull request on its own.
I still review every pull request as an engineer: does it solve the original problem, are the trade-offs acceptable, is the diff understandable, and is the verification strong enough?
If I have feedback, I give it to the manager. The manager can send the work back to the same executor or create another review pass. Once I am satisfied, I merge.
My role moves up a level. I spend more time on framing, judgment, and review, while agents handle more of the execution and coordination.
I do not assign permanent job titles to tools. Codex is not always the implementer, and Claude Code is not always the reviewer.
I choose based on the work, the context already available, the model that I want to use, and what I need from that session. I may also want a different harness for a fresh perspective. The choice can change as the tools change.
This is another reason the shared layer matte

[truncated]
