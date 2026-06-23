---
source: "https://blog.jetbrains.com/junie/2026/06/junie-coding-agent-out-of-beta/"
hn_url: "https://news.ycombinator.com/item?id=48643598"
title: "Junie: The JetBrains AI Coding Agent Leaves Beta"
article_title: "Junie: The JetBrains AI Coding Agent Leaves Beta | Try now"
author: "roflcopter69"
captured_at: "2026-06-23T12:54:13Z"
capture_tool: "hn-digest"
hn_id: 48643598
score: 2
comments: 0
posted_at: "2026-06-23T11:52:50Z"
tags:
  - hacker-news
  - translated
---

# Junie: The JetBrains AI Coding Agent Leaves Beta

- HN: [48643598](https://news.ycombinator.com/item?id=48643598)
- Source: [blog.jetbrains.com](https://blog.jetbrains.com/junie/2026/06/junie-coding-agent-out-of-beta/)
- Score: 2
- Comments: 0
- Posted: 2026-06-23T11:52:50Z

## Translation

タイトル: Junie: JetBrains AI コーディング エージェントがベータ版を終了
記事のタイトル: Junie: JetBrains AI コーディング エージェントがベータ版を終了 |今すぐ試してください
説明: Junie は実験として始めました。私たちは「AI コーディング エージェントがプロジェクトの詳細を推測するだけでなく、実際にあなたと同じツールを使ったらどうなるでしょうか?」と尋ねました。昨年、その実験は

記事本文:
プラグインとサービス
ビッグデータツール
.NET と Visual Studio
.NETツール
教育と研究
ジェットブレインズアカデミー
Junie: JetBrains AI コーディング エージェントがベータ版を終了
この投稿を他の言語で読んでください:
ジュニーは実験として始めました。私たちは「AI コーディング エージェントがプロジェクトの詳細を推測するだけでなく、実際にあなたと同じツールを使ったらどうなるでしょうか?」と尋ねました。昨年、その実験は開発者によって IDE とターミナル内で毎日使用される製品になりました。
本日、JetBrains AI コーディング エージェントはベータ版を終了します。これは名前変更や再パッケージ化ではありません。 Junie の最も重要な部分は安定して接続されており、実際の作業に使用できる状態になっています。 Junie はコードを書く前に計画を立て、実際のデバッガーでデバッグし、プロジェクトのコンテキストを考慮しながら PR をレビューし、ユーザーが他のことに集中している間に長いタスクを実行します。
独立系エージェントのベンチマークである SWE-Rebench の最新の実行で、ジュニーはナンバーワンのコーディング エージェントに選ばれました。
「SWE-Rebench は、評価を誠実に保つためにサイクルごとに新しいタスクを描画するため、結果は実行ごとに変化します。このサイクルでは、Junie が 61.6% 解決、72.7% 合格@5 でトップのモデルハーネスとなりました。これにより、他のエージェントよりも上位に位置し、生のフロンティア モデルと競争力がありました。」
私たちは、エージェントへの仕事の委任は、一度限りの英雄的なことだけでなく、頻繁に実行できる余裕のあるものであるべきだと考えています。したがって、Junie はロックインせずにあらゆるモデルをサポートし、それがコストを管理する方法です。フロンティア ラボの最新モデルをゼロから使用することも、Junie をローカル ランタイムに指定することもできます。これは、各タスクのコストを決定するためのレバーです。最上位の推論モデルは強力ですが高価です。小型モデルは高速かつ安価です。 Junie では、それぞれを最も効果的な場所に配置できます。コスト効率はツールの特性ではなくなります

そして握るダイヤルになります。
一般提供への移行に伴う内容は次のとおりです。
Advanced Plan モード: エージェントはコーディングする前に考えます。
AI コーディング エージェントの失敗の最も一般的な原因の 1 つは、AI コーディング エージェントが完全に間違っているときの揺るぎない自信です。彼らは、自分たちがやっていることに誰も同意する前に実装を開始します。結局、間違った問題を解決する PR をレビューしたり、最初の 30 秒で拒否したはずのパスでトークンを焼き付けたりすることになります。
プラン モードでは、プランを最上級の成果物にすることでこの問題を修正します。
Junie はコードを記述する前に、製品要件、技術設計、納品段階、および (要求に応じて) テスト戦略のタブを含む構造化ドキュメントを作成します。あなたはドキュメントを読みました。エディターで直接編集します。あなたはそれを承認します。そしてジュニーはそれを実行します。
このアプローチは、次のような理由から「より適切なプロンプト」よりも優れています。
計画書は本物の文書です。 .junie/plans にあります。それをコミットすれば、使い捨てのチャット メッセージではなく、生きたタスクのドキュメントになります。
エージェントは適切な質問をします。要件があいまいな場合、ジュニーは推測したり期待したりするのではなく、問題を突き止めるために多肢選択式の質問を自由形式で行います。
Junie はコード化する前に計画を立てます。つまり、無駄なトークンや壊れた PR が少なくなります。無駄な実装実行はすべて、支払ったトークンであり、とにかく実行する必要があるレビュー サイクルです。強力なモデルを計画します。安いもので実装します。エージェントは徘徊しないため、請求額は低く抑えられます。
Shift+Tab で計画モードに入ります。 Ctrl+P でプランを開きます。準備ができたら、「確認」をクリックして変更を実装します。
エージェント デバッグ: Junie は println ではなくデバッガーを使用します。
何か問題が発生した場合、ほとんどのコーディング エージェントはログ ステートメントを追加します。 Junie はデバッガーを開きます。
GA バージョンでは、Junie が IDE を駆動できます

のデバッガを次のように実行します。
デバッグ セッションを開始するか、デバッグ セッションに参加します。 Junie は、実行構成の起動、テストのデバッグ、またはすでに開いている既存のセッションの引き継ぎを行うことができます。
プロジェクト コード、ライブラリ コード、SDK コード、さらには逆コンパイルされた .class ファイルや JAR 内のソースなど、重要な場所にブレークポイントを設定します。 IDE がそれにステップインできる場合、Junie はそれにブレークポイントを設定できます。
実際の実行時の状態を検査します。スタック フレーム、スレッド状態、式の評価、実行から行まで – Junie は、コードが何を実行しているかを理論化するのではなく、実際の証拠を収集します。
これにより、以前は手動で作業する必要があったデバッグ パターンを Junie が使用できるようになります。
「デバッグして、このテストが 2 回目の反復でのみ失敗する理由を解明してください。」完全自律型 – Junie がすべてを運転します。
「デバッガーを準備してください。UI フローをトリガーします。」 Junie はブレークポイントを設定してあなたを待ちます。
「現在のデバッグ セッションを続行して、この値が null になる理由を教えてください。」全体像を考えている間、日常的な検査作業を引き継ぎます。
現在、これは AI サブスクリプションを備えた JetBrains IDE でエンドツーエンドで機能します。
リモートコントロール: タスクを開始し、どこからでも監視します
Spring Boot のアップグレード、Java レコードへの移行、従来のサービスへのテスト カバレッジの追加など、集中した 30 分のセッションに収まらない作業もあります。これらはまさに自律エージェントが得意とするタスクであり、座って見守る必要がなければさらに優れています。
ラップトップからタスクを開始します。会議中にスマートフォンから進行状況を確認します。コーヒーを飲みながら PR を見直しましょう。 Junie は非同期で実行され、サインインしているどこからでもセッションを利用できるようにします。
コンテキストを失わずにコードをレビューする
ほとんどのレビュー ツールは、PR が開いたときに初めてコードベースを確認します。ジュニーも同じレビューをしています

コードを記述するために使用されるプロジェクト コンテキスト: ビルド、テスト、規約、過去の決定。
エントリーポイントは3つ。 GitHub Actions または GitLab (オンプレミスを含む) からレビューをトリガーするか、CLI またはプラグインで /review コマンドを使用します。スコープを、ステージングされていない変更、ステージングされた変更、またはメインとの差分 (呼び出し) に設定します。
インタラクティブなウォークスルー。 Junie は、意味のある変更をそれぞれ強調し、その背後にある設計上の決定を説明し、インラインで承認/拒否のコントロールを提供します。何か問題があると思われる場合は、その場で PR コメントをドロップしてください。
あなたの焦点に合わせて適応します。フォローアップの質問をすると、Junie はファイルをアルファベット順に検索するのではなく、残りのレビューを関心のある内容に基づいて並べ替えます。
IDE の緊密な統合: IDE のツールを使用する AI コーディング エージェント
Junie は常に JetBrains IDE 内で働いてきました。今年の初めに、その接続方法を紹介しました。 Junie の GA バージョンでは、Junie CLI が IDE と通信するために使用するのと同じプロトコルである ACP (エージェント通信プロトコル) 上にその統合を再構築しました。
1 つのエンジンに多くの路面。同じエージェントが AI チャット、専用の Junie ツール ウィンドウ、および Junie CLI の背後にあります。改善は一度配信されると、どこにでも反映されます。
IDE、エージェントのツールボックス。 Junie は、独自の近似ではなく、IDE のセマンティック インデックス、ビルド構成、テスト ランナー、デバッガーを使用します。
データベースの統合。 Junie は、DataGrip と JetBrains Database プラグインを介して IDE で構成されたデータベースに接続し、コードを処理する同じセッションで実際のデータをクエリし、SQL の書き込み、修正、検証を行います。
これらの各機能は個別に特定の問題を解決します。これらが連携することで、エージェントの役割が変わります。
あなたのプロジェクトを理解しているエージェントは、作業を開始する前に承認を与えてくれます。

他のことをしている間にそれを実行し、問題が発生したときに適切にデバッグし、プロジェクトの完全なコンテキストを含む PR をレビューし、実際のデータをクエリします。これが実際に委任できるエージェントです。
それがベータ版を終了するために私たちが設定した基準です。
Junie は、すべての JetBrains IDE およびターミナルの Junie CLI を通じて利用できます。すでに JetBrains AI サブスクリプションをお持ちの場合は、何も設定しなくてもすべて機能します。
Bring Your Own Key も機能します。Anthropic、OpenAI、Google などへのアクセスをお楽しみください。
Junie はローカル モデル ランタイムに接続します。LiteLLM、LMStudio、Ollama を指定すると、自分のマシンにロードしたモデルを使用してエージェントが実行されます。プロンプトとコードは外部に共有されることはありません。
Junie をインストールし、プロジェクトを開いて、実際のタスク (おそらく先延ばしにしていたタスク) でテストしてみます。
次に、何がうまくいかなかったのか、何が驚いたのか、次に何を見てみたいのかを教えてください。上記のすべての機能はそのフィードバック ループから生まれており、GA への移行で終わるわけではありません。
Junie CLI と JetBrains IDE の融合: 共有コンテキスト、セマンティック リファクタリング、スマート検索、ビルド構成など、ターミナル エージェントがさらに賢くなりました。
LLM に依存しないコーディング エージェントである Junie CLI は現在ベータ版です。 Junie CLI の今後のリリースでは、IDE 内、CI/CD、GitHub または GitLab 上でターミナルから直接 Junie を使用できるようになります。
これが JetBrains の AI コーディング エージェントである Junie でどのように機能するかを見てみましょう。
Junie によるシームレスな IDE 内コーディングと非同期開発
ソフトウェア開発は、AI コーディング エージェントがより反復的なタスクを引き継ぎ、コードをより速く生成し、より少ない労力でより多くのことを成し遂げることによって大きな影響を受けてきました。高速コーディング エージェントや IDE 内の AI にアクセスできるのは素晴らしいことです…

## Original Extract

Junie started as an experiment. We asked, “What if an AI coding agent didn't just guess at the details of your project, but actually used the same tools you do?” Over the last year, that experiment tu

Plugins & Services
Big Data Tools
.NET & Visual Studio
.NET Tools
Education & Research
JetBrains Academy
Junie: The JetBrains AI Coding Agent Leaves Beta
Read this post in other languages:
Junie started as an experiment. We asked, “What if an AI coding agent didn’t just guess at the details of your project, but actually used the same tools you do?” Over the last year, that experiment turned into a product used by developers every day – inside the IDE and the terminal .
Today, the JetBrains AI coding agent is leaving Beta. This isn’t a rename or a repackage. The parts of Junie that matter most are stable, connected, and ready for real work. Junie plans before it codes, debugs with the real debugger, reviews PRs while considering your project’s context, and runs long tasks while you focus on other things.
On the latest run of SWE-Rebench – an independent agent benchmark – Junie placed as the number-one coding agent.
“SWE-Rebench draws fresh tasks each cycle to keep the evaluation honest, so results move from run to run. In this cycle Junie came out as the top model-harness, with 61.6% resolved and a 72.7% pass@5 — placing it ahead of the other agents and competitive with raw frontier models”
We believe that delegating work to an agent should be something you can afford to do often, not just for heroic one-offs. Thus Junie supports any model, without lock-in – and that’s how you control cost. Use the latest models from frontier labs from day zero, or point Junie at a local runtime. It’s the lever that lets you decide what each task costs. Top-tier reasoning models are powerful but expensive; smaller models are fast and cheap. Junie lets you put each one where it does the most good. Cost efficiency stops being a property of the tool and becomes a dial you hold.
Here’s what comes with the move to general availability:
Advanced Plan mode: The agent thinks before it codes
One of the most common causes of failure in AI coding agents is unwavering confidence when they are totally incorrect – they start implementing before anyone has agreed on what they’re doing. You end up reviewing a PR that solves the wrong problem or burning tokens on a path you would have rejected in the first thirty seconds.
Plan mode fixes that by making the plan a first-class artifact.
Before Junie writes code, it produces a structured document with tabs for product requirements, technical design, delivery stages, and (when requested) testing strategy. You read the doc. You edit it directly in your editor. You approve it. And then Junie implements it.
This approach is superior to “better prompting”, for a few reasons:
The plan is a real document. It lives in .junie/plans. You can commit it, and it becomes living task documentation, not a throwaway chat message.
The agent asks the right questions. When requirements are ambiguous, Junie asks multiple-choice and freeform questions to pin things down, instead of guessing and hoping.
Junie plans before it codes – meaning fewer wasted tokens and fewer broken PRs. Every wasted implementation run is tokens you paid for and a review cycle you’ll have to do anyway. Plan on a strong model; implement on a cheap one. The agent doesn’t wander, so your bill stays low.
Enter Plan mode with Shift+Tab. Open the plan with Ctrl+P. And when you’re ready, hit Confirm to implement the changes.
Agentic debugging: Junie uses the debugger, not println
When something goes wrong, most coding agents add log statements. Junie opens the debugger.
With the GA version, Junie can drive your IDE’s debugger the way you would:
Start or join a debug session. Junie can launch a run configuration, debug a test, or take over an existing session you already have open.
Set breakpoints anywhere that matters , including project code, library code, SDK code – even decompiled .class files and sources inside JARs. If your IDE can step into it, Junie can set a breakpoint in it.
Inspect the real runtime state. Stack frames, thread state, expression evaluation, run-to-line – Junie collects actual evidence instead of theorizing about what your code might be doing.
This allows Junie to use debugging patterns that you previously had to work with manually:
“Debug and figure out why this test fails only on the second iteration.” Fully autonomous – Junie drives the whole thing.
“Prepare the debugger, I’ll trigger the UI flow.” Junie sets up breakpoints and waits for you.
“Continue my current debug session and tell me why this value becomes null.” Hand off routine inspection work while you think about the bigger picture.
Today this works end to end in JetBrains IDEs with an AI subscription.
Remote control: Start a task, and keep an eye on it from anywhere
Some work doesn’t fit in a focused 30-minute session, for example a Spring Boot upgrade, a migration to Java records, or adding test coverage to a legacy service. These are exactly the tasks autonomous agents are good at – and it’s even better when you don’t have to sit and watch.
Start a task from your laptop. Check progress from your phone during a meeting. Review the PR over coffee. Junie runs asynchronously and keeps the session available from anywhere you sign in.
Code review without lost context
Most review tools see your codebase for the first time when the PR opens. Junie reviews with the same project context it uses to write code: your build, your tests, your conventions, your past decisions.
Three entry points. Trigger a review from GitHub Actions or GitLab (including on-prem), or by using the /review command in the CLI or the plugin. Set the scope to unstaged changes, staged changes, or a diff against main – your call.
Interactive walkthrough. Junie highlights each meaningful change, explains the design decision behind it, and gives you accept/reject controls inline. Drop a PR comment on the spot when something looks wrong.
Adaptation to your focus. Ask a follow-up question and Junie reorders the remaining review around what you care about, instead of marching through files alphabetically.
Deep IDE integration: An AI coding agent that uses your IDE’s tools
Junie has always worked inside JetBrains IDEs. Earlier this year we showed you how to connect it . In Junie’s GA version, we’ve rebuilt that integration on top of ACP (the Agent Communication Protocol), the same protocol Junie CLI uses to talk to your IDE.
One engine, many surfaces. The same agent is behind the AI chat, the dedicated Junie tool window, and Junie CLI. Improvements ship once and show up everywhere.
Your IDE, the agent’s toolbox. Junie uses your IDE’s semantic index, build configurations, test runners, and debugger, not its own approximation of them.
Database integration. Junie connects to the databases configured in your IDE through DataGrip and the JetBrains Database plugin, and then it queries your real data and writes, fixes, and validates SQL in the same session that handles your code.
Individually, each of these features solves a specific problem. Together, they change what an agent is for .
An agent that understands your project, lets you approve the work before doing it, runs it while you’re doing something else, debugs it properly when things break, reviews your PRs with the full project context, and queries your real data – that’s an agent you can actually delegate to.
That’s the bar we set for leaving Beta.
Junie is available in all JetBrains IDEs and through Junie CLI in your terminal. If you already have a JetBrains AI subscription, everything works out of the box.
Bring Your Own Key works too – enjoy access to Anthropic, OpenAI, Google, and others.
Junie connects to local model runtimes – point it at LiteLLM, LMStudio, Ollama and the agent runs using whatever model you have loaded on your own machine. Prompts and code never shared externally.
Install Junie, open your project, and test it out on a real task (maybe one you’ve been procrastinating on).
Then tell us what broke, what surprised you, and what you’d like to see next. Every feature above came from that feedback loop, and it doesn’t end with the move to GA.
Junie CLI meets JetBrains IDE: shared context, semantic refactoring, smart search, and build configs — your terminal agent just got a whole lot smarter.
Junie CLI, the LLM-agnostic coding agent, is now in Beta. With the upcoming release of Junie CLI, you will be able to use Junie directly from the terminal, inside any IDE, in CI/CD, and on GitHub or GitLab.
Let’s see how this works with Junie, the AI coding agent by JetBrains.
Seamless in-IDE coding and async development with Junie
Software development has been significantly influenced by AI coding agents taking over more repetitive tasks, generating code faster and getting more done with less effort. It’s nice to have access to a fast coding agent or an in-IDE AI …
