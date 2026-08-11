---
source: "https://www.theaithinker.com/p/how-to-organize-claude-code-for-product"
hn_url: "https://news.ycombinator.com/item?id=49256258"
title: "How to organize Claude Code for product work"
article_title: "How to organize Claude Code for product work - by Adam Faik"
author: "adamfaik"
captured_at: "2026-08-11T11:38:03Z"
capture_tool: "hn-digest"
hn_id: 49256258
score: 2
comments: 0
posted_at: "2026-08-11T11:02:34Z"
tags:
  - hacker-news
  - translated
---

# How to organize Claude Code for product work

- HN: [49256258](https://news.ycombinator.com/item?id=49256258)
- Source: [www.theaithinker.com](https://www.theaithinker.com/p/how-to-organize-claude-code-for-product)
- Score: 2
- Comments: 0
- Posted: 2026-08-11T11:02:34Z

## Translation

タイトル: 製品作業用にクロード コードを整理する方法
記事のタイトル: 製品作業用にクロード コードを整理する方法 - Adam Faik 著
説明: スターター ワークスペースをダウンロードします。セットアップスキルを実行します。ファイルの変更速度に基づいてファイルを並べ替えます。繰り返しの作業をスキルに変えます。すべての修正を一度ファイルします。

記事本文:
製品作業用にクロード コードを整理する方法 - Adam Faik 著
AI の思想家
製品作業用に Claude Code を整理する方法
スターター ワークスペースをダウンロードします。セットアップスキルを実行します。ファイルの変更速度に基づいてファイルを並べ替えます。繰り返しの作業をスキルに変えます。すべての修正を一度ファイルします。
Adam Faik 2026 年 8 月 10 日 3 Share My Claude 製品ジョブのコード ワークスペースは、コードがほとんど含まれていないフォルダーです。製品、ユーザー、競合他社、私の好みの働き方など、会社に関するあらゆる情報が含まれたコンテキスト フォルダーが保存されています。プロジェクトごとに 1 つのフォルダーがあり、タスクは成果物から分離されています。それぞれが 1 つの繰り返し作業 (ステータスの更新、PRD レビュー、インタビューの統合) を実行する一連のスキル。一度フィードバックを与えると、その理由が記載されたメモがファイルに保存され、そのフィードバックは二度と与えません。私がこのシステムを設計したわけではありません。私がそれに到達したのは、主に最初に間違ったことからでした。
やり方を間違えると、この記事全体に通ずる教訓が得られました。基本的なことを除けば、Claude Code の結果は、プロンプトの内容に応じて停止し、ファイルの内容に応じて開始されます。適切なプロンプトは 1 つのセッションを改善します。適切なファイルは、その後のすべてのセッションを改善します。あなたがすでに読んだ迅速なチューニングのアドバイスは、まさにファイリングが重要になり始めた時点で役に立たなくなります。
そこで、自分のシステムを説明する代わりに、それをパッケージ化しました。スターター ワークスペースは GitHub で公開されています。フォルダー、コンテキスト テンプレート、および 5 つの PM スキルが含まれており、1 回のセットアップ インタビューによってパーソナライズされます。技術的な知識は必要ありません。ツールのインストールを含むすべての手順が詳しく説明されています。私の例は製品の仕事ですが、同じアーキテクチャは、デザイン リーダーの研究ライブラリ、データ チームのメトリクス定義、創業者のすべてのフォルダーなど、コンテキストを重視するあらゆる作業に適合します。
s

魔女。チャット モードと Cowork では提供できないもの、そしてプロダクト マネージャーにとってこの端末がなぜ価値があるのか​​を説明します。
スターターワークスペース。それをダウンロードし、セットアップ スキルを 1 つ実行し、実際のタスクを実行します。6 つのステップがそれぞれ画面に表示されます。
実践。システムを複合化する 6 つの動き。それぞれに、実行すべきことと何が返されるかを正確に指定します。
リズム。最初の短いプロンプトから 2 分間のファイリング習慣まで、システム内での通常の勤務日がどのように見えるかを次に示します。
その見返りは具体的です。ツールをインストールし、ワークスペースをダウンロードし、セットアップのインタビューに答え、最初の実際のタスクを引き渡すだけで、一度座るだけですぐに実行できます。セッションのたびに自分の会社について再度説明する必要がなくなり、修正した内容は永久的なものになります。各ステップは以下の通りです。
ファイルキャビネットを開ける時間です。
おそらく今日は、Claude のチャットを通じて製品の作業を実行しているでしょう。ブラウザのタブやデスクトップ アプリ、おそらくエージェント タスク用の Cowork を使用しています。それはうまくいきますが、それはまさに罠です。すべての会話はゼロから始まります。あなたが最後にアップロードした戦略ドキュメントは失われ、あなたが辛抱強く説明したコンテキストは二度と見つけることのないスレッドに存在し、気に入った出力はファイルに保存されるのではなく会話の中に閉じ込められます。チャットは考えるのには最適な場所ですが、蓄積するにはひどい場所です。たまに質問する場合は大丈夫です。毎日の製品作業では、毎朝同じ設定コストを支払うことを意味します。
クロード コードは、同じクロードのホームが異なります。つまり、読み取り、移動、書き込みができるマシン上のフォルダーです。コンテキストは永続的なファイル内に存在します。定期的な作業には、コマンドに基づいて実行されるスキルが必要です。スクロールするメッセージではなく、保存するドキュメントとして出力されます。私の毎日の使用では、チャットや Cowork で行っていたすべての作業がここでより適切に実行されます。何も説明されないためです。

2回編集しました。これは私自身の仕事からの観察であり、法律ではありませんが、何ヶ月もの間維持されています。
実際に私の仕事を変えた部分はその下の端子です。 Claude Code は GitHub とネイティブに通信し、コネクタは必要ありません。会社のリポジトリを取得して、それに関する質問に答えたり、自分の作業をプッシュしたりできます。製品の背後にあるリポジトリを取得し、機能フラグが実際に何をするのかを説明するように依頼します。ワークスペースをプライベート リポジトリにプッシュすると、チームメイトがそのクローンを作成して、有利なスタートを引き継ぐことができます。プロダクト マネージャーにとって、これはコードベースへの直接の入力であり、システム全体を共有する方法であり、どちらもチャットにはないものです。
誰もあなたに渡さないのは、これらすべての化合物を作成するファイルシステムです。 Anthropic のベスト プラクティス ガイドは、独自の言葉で「さまざまなコードベース、言語、環境で Claude Code を使用するエンジニア向け」に書かれており、エンジニアは構造を自由に利用できます。コードベースはすでに組織化されたコンテキストです。 Sachin Rekhi が最もよく知られている既存の PM ガイドでは、「すべての製品データとドキュメントをローカルのマークダウン ファイルに保存しなさい」と正しく指示されています。どのように整理されましたか？答えなしで試してみると、私が抱えていたものと同じことがわかります。つまり、40 個の未解決のファイルと古いセットアップです。スターター ワークスペースはその答えであり、事前に構築されています。設定してみましょう。
スターター ワークスペースから始める
スターター ワークスペースは、私が独自の製品作業を実行するための構造であり、フォルダー、ガイド付きプレースホルダーを含むコンテキスト ファイル、サンプル プロジェクト、および既にインストールされている 5 つのスキルをテンプレートにまとめたものです。すべては無料で公開されています: claude-code-pm-starter 。 6 つのステップにより、何もない状態から最初の実際のタスクに移行しますが、どのステップも技術的な背景を前提としていません。
始める前に、前提条件について注意してください。 Claude Code には有料の Claude プランが必要です。エヴ

インストールする 2 つのツールを含む他のものはすべて無料であり、手順で説明されています。
Claude コードはターミナル内で実行されますが、ターミナルだけではファイルが見えないため、プロダクト マネージャーにとっては不適切な場所です。修正はIDEです。 IDE は、作業中にフォルダーをファイルのパネルとして表示する単なるエディターです。開発者は 1 日中この 1 台に住んでおり、ファイル パネルのためだけにそれを借りることになります。無料の標準である VS Code をダウンロードし、他のアプリと同様にインストールします。コードは書きません。ワークスペースが左側のパネルで自動的に整理されるのがわかります。これはまさにチャット モードでは得られなかった可視性です。
VS Code: 無料ですが、その一部のみを使用します。ステップ 2: ワークスペースをダウンロードする
リポジトリ ページを開き、緑色の「コード」ボタンをクリックして、「Download ZIP」を選択します。 git もアカウントも必要ありません。それを解凍し、作業内容を保存する場所にフォルダーを移動します。このフォルダーは、製品の作業場所となるため、ホーム ディレクトリの作業フォルダーなど、再び見つけやすい場所を選択してください。
緑色の「コード」ボタンをクリックし、「ZIP をダウンロード」します。 git は必要ありません。解凍してファイルに保存すると、作業フォルダー内のスターター フォルダーが開き、すぐに開くことができます。ステップ 3: VS Code で開き、右側のターミナル
VS Code を開き、 [ファイル] > [フォルダーを開く] を選択し、解凍されたフォルダーを選択します。左側のパネルには、コンテキスト/ 、プロジェクト/ 、オペレーション/ 、スキルの全体構造が表示されます。次に、[ターミナル] > [新しいターミナル] を使用して VS Code 内でターミナルを開きます。ここで、全体の印象を変える小さなトリックを紹介します。ターミナル パネルのヘッ​​ダーを右クリックし、[パネルを右に移動] を選択します。左側にファイル、右側にエージェント: コマンド ラインという感覚はなくなり、ドキュメントの隣に座っている同僚のように感じられるようになります。
ついに見える全体構造: スキル、コンテキスト、オペレーション、プロジェクト、

そして地図。コツ: パネルを右クリックし、パネルを右に移動します。 Claude Code をまだお持ちでない場合は、公式ガイドから今すぐインストールしてください。それはそのターミナルに貼り付けられた 1 つのコマンドです。次に、「claude」と入力して Enter キーを押します。 Claude Code はフォルダー内で開始され、マップ ファイルを読み取り、最初のメッセージからすべてがどこに存在するかを認識します。
左側がファイル、右側がエージェントです。ステップ 4: こんにちはと言ってから 4 つのコマンドを学習する
1 つの違いにより、初心者は多くの混乱を避けることができます。端末は正確なコマンドを必要とします。クロードは普通の言葉を望んでいます。 clude と入力する前は、スペルと構文が重要となるターミナルと会話していました。クロード コードが実行されているので、チャット アプリとまったく同じようにクロードと話していることになります。したがって、最初に当然のことを行います。hello と入力して、このフォルダーに何が表示されるかを尋ねます。あなたのファイルを読んだだけの同僚のように答えます。
クロード コードに直接入力できるいくつかのスラッシュ コマンドは、初日から知っておく価値があります。
/model - 実行しているクロード モデルを表示し、切り替えることができます。
/effort - クロードが各回答にどれだけの思考を費やすかを設定します。ハードワークの場合は労力が高く、素早いドラフトの場合は労力が低くなります。
/mcp - MCP、Jira、Slack、Confluence などのツールへのライブ接続を一覧表示し、それらを管理できるようにします。
/plugins - プラグイン システムを開きます。これには 1 分かかります。 Claude Code には、デフォルトで利用可能な Anthropic のプラグイン ディレクトリが付属しており、その上にマーケットプレイスを追加できます。 1 つはプロダクト マネージャー向けに特別に構築されたものです。 pm-skills は、発見、戦略、実行、立ち上げにわたる 100 を超える PM スキルです。 README には最新のインストール手順が記載されています。 pm-execution は PRD、OKR、ロードマップ、振り返りをカバーします。 pm-product-discovery には、インタビューと仮定テストが追加されます。すべてではなく、作業に一致する 2 つまたは 3 つをインストールしてください

九;インストールされているすべてのプラグインは、システムが持つコンテキストです。
プラグイン パネル: Anthropic のマーケットプレイスはデフォルトで存在します。もう 1 つのコントロールを追加すると、ツアーは終了です。Shift+Tab を押して、Claude Code の 3 つのモードを循環させます。
デフォルト モードでは、ファイルに触れる前に許可を求めます。
自動承認では、確認を求めずに編集内容が適用されます。
プラン モードは覚えておくべきモードです。クロードが調査してプランを提案します。ユーザーが承認するまでは何も変更されません。
デフォルトで開始し、大きなことにはプラン モードを使用し、信頼できる作業については自動承認を保存します。これら以外のすべてについては、このコミュニティ チートシートが、私が見つけた 1 ページのリファレンスの中で最も優れています。
まずは挨拶をしましょう。コマンドは 1 分間待つことができます。ステップ 5: セットアップ スキルを実行する
「/setup-workspace」と入力し、Enter キーを押します。ワークスペースには、テンプレートをパーソナライズすることだけを行うスキルが付属しています。このスキルは、会社、製品、ユーザー、働き方についてインタビューし、コンテキスト ファイルを入力して、サンプル プロジェクトの名前を最もアクティブな実際のプロジェクトに変更します。平易な言葉で答えてください。書き込みが行われ、各ファイルに「はい」または「修正」が表示されます。
リンクを貼り付けて、記憶に基づいて答える必要はありません。製品の公開サイト、ドキュメント ページ、または前のステップで Atlassian MCP に接続した場合はチームに関する Confluence ページを渡すと、それらを読み取って、すでに書き込まれた内容からコンテキスト ファイルを構築します。この 15 分で、白紙ページの問題は完全に置き換えられます。あなたの正確な質問は私の質問とは少し異なります。各セッションは少し予測不可能ですが、それがこれらのモデルの仕組みです。まだ知らないことは飛ばしてください。正直なプレースホルダーは、作られたコンテンツに勝ります。
わかりやすい言葉で答えるか、リンクを貼り付けて読んでください。ステップ 6: 実際のタスクを与える
もみを終わらせないでください

セットアップ時の最初のセッション。実際に誰かに借りているものを渡してください。 /status-update と入力すると、記憶ではなくプロジェクトのタスク ファイルから更新の下書きが作成されるのを確認します。または、実際の成果物を指します。
[あなたのプロジェクト] のキックオフ概要の草案を作成します。 context/preferences.md の書き込み設定と一致します。結論を下す前に、プロジェクトのtasks.md内の未解決の質問を確認してください。
プロジェクトの実際の状態を含む、自分らしい製品についての最初の出力は、ファイリングがクリックされた瞬間です。さらに 3 つのスキルがフォルダー内で待機しています。じっくり読みたい次のドキュメントの prd-review 、次のユーザーとの会話のための Interview-Synthetic 、および修正を永続化する file-フィードバック です。
スキルは最初にワークスペースの状態を読み取り、次に誰に対する更新かを尋ねます。それが始まりです。システムが一度に動作するのです。それをさらに悪化させるのは、栄養を与える方法であり、それが 6 つの実践方法です。
スターター構造は、最初にそれぞれを間違えることによって集中した実践をエンコードしています。それらは従うべき方法論ではありません。彼らはその地位を獲得し続けた動きです。以下のそれぞれは、具体的に行うべきことと、それに対して得られるものを示しています。 1つ目はエンジンです。残りは操縦します。
毎日のあらゆるタスクに使用してください
最も高い収益をもたらす単一の習慣は、最も単純な習慣でもあります。

[切り捨てられた]

## Original Extract

Download the starter workspace. Run the setup skill. Sort files by how fast they change. Turn repeated work into skills. File every correction once.

How to organize Claude Code for product work - by Adam Faik
The AI Thinker
Subscribe Sign in How to organize Claude Code for product work
Download the starter workspace. Run the setup skill. Sort files by how fast they change. Turn repeated work into skills. File every correction once.
Adam Faik Aug 10, 2026 3 Share My Claude Code workspace for my product job is a folder with almost no code in it. It holds a context folder with everything about the company: the product, the users, the competitors, how I like to work. One folder per project, with the tasks separated from the deliverables. A set of skills that each do one recurring piece of work: the status update, the PRD review, the interview synthesis. When I give feedback once, it lands in a file with a note on why, and I never give that feedback again. I didn’t design this system. I arrived at it, mostly by getting it wrong first.
Getting it wrong taught me the lesson that runs this whole article. Past the basics, your results in Claude Code stop depending on how well you prompt and start depending on how well you file. A good prompt improves one session; a good file improves every session after it. The prompt-tuning advice you’ve already read stops helping at exactly the point where the filing starts to matter.
So instead of describing my system, I packaged it. The starter workspace is public on GitHub: the folders, the context templates, and five PM skills, personalized by one setup interview. You don’t need to be technical to follow along; every step is spelled out, including installing the tools. And while my example is a product job, the same architecture fits any context-heavy work: a design lead’s research library, a data team’s metric definitions, a founder’s everything-folder.
The switch. I’ll show you what chat mode and Cowork can’t give you, and why the terminal is worth it for a product manager.
The starter workspace. Download it, run one setup skill, and give it a real task: six steps, each shown on screen.
The practices. Six moves that make the system compound, each one naming the exact thing to do and what you get back.
The rhythm. Here’s what a normal workday looks like inside the system, from the first short prompt to the two-minute filing habit.
The payoff is concrete. One sitting gets you running: install the tools, download the workspace, answer the setup interview, hand over a first real task. You stop re-explaining your company every session, and every correction you make becomes permanent. Every step is below.
Time to open the filing cabinet.
You probably run your product work through Claude’s chat today: a browser tab or the desktop app, maybe Cowork for the agentic tasks. It works, and that’s exactly the trap. Every conversation starts from zero. The strategy doc you uploaded last time is gone, the context you patiently explained lives in a thread you’ll never find again, and the output you liked is trapped inside a conversation instead of sitting in a file. Chat is a great place to think and a terrible place to accumulate. For occasional questions that’s fine. For daily product work, it means paying the same setup cost every single morning.
Claude Code is the same Claude with a different home: a folder on your machine that it can read, navigate, and write to. Context lives in files that persist. Recurring work lives in skills that run on command. Outputs land as documents you keep, not messages you scroll for. In my daily use, everything I used to do in chat or Cowork runs better here, because nothing is ever explained twice. That’s an observation from my own work, not a law, but it’s held for months.
The part that actually changed my job is the terminal underneath. Claude Code speaks to GitHub natively, no connector needed: it can pull a company repository and answer questions about it, and it can push your own work. Ask it to fetch the repo behind your product and explain what a feature flag actually does. Push your workspace to a private repository and a teammate can clone it and inherit your head start. For a product manager, that’s a direct line into the codebase and a way to share your whole system, both of which chat simply doesn’t have.
What nobody hands you is the filing system that makes all this compound. Anthropic’s best-practices guide is written, in its own words, “for engineers using Claude Code across various codebases, languages, and environments,” and engineers get their structure free: the codebase is already organized context. The PM guides that exist, Sachin Rekhi’s being the best known, rightly tell you to “store all of your product data and documentation in local markdown files,” then move on. Organized how? Try it without an answer and you get what I had: forty loose files and a stale setup. The starter workspace is that answer, prebuilt. Let’s set it up.
Start with the starter workspace
The starter workspace is the structure I run my own product work on, cleaned into a template: the folders, the context files with guided placeholders, an example project, and five skills already installed. The whole thing is free and public: claude-code-pm-starter . Six steps take you from nothing to a first real task, and none of them assume a technical background.
One prerequisites note before you start. You need a paid Claude plan for Claude Code. Everything else, including the two tools you’ll install, is free and covered in the steps.
Claude Code runs in a terminal, and a terminal alone is a bad home for a product manager: you can’t see your files. The fix is an IDE. An IDE is just an editor that shows your folder as a panel of files while you work; developers live in one all day, and you’ll borrow it for the file panel alone. Download VS Code , the free standard, and install it like any app. You won’t write code in it. You’ll watch your workspace organize itself in the left panel, which is exactly the visibility chat mode never gave you.
VS Code: free, and you’ll only use a fraction of it. Step 2: Download the workspace
Open the repository page , click the green Code button, and choose Download ZIP . No git, no account needed. Unzip it and move the folder wherever you keep your work. This folder is about to become the place your product work lives , so pick somewhere you’ll find it again, like a work folder in your home directory.
The green Code button, then Download ZIP. No git needed. Unzipped and filed: the starter folder in a Work folder, ready to open. Step 3: Open it in VS Code, terminal on the right
Open VS Code, choose File > Open Folder , and pick the unzipped folder. The left panel now shows the whole structure: context/ , projects/ , operations/ , the skills. Then open the terminal inside VS Code with Terminal > New Terminal , and here’s the small trick that changes how the whole thing feels: right-click the terminal panel’s header and choose Move Panel Right . Files on the left, agent on the right: it stops feeling like a command line and starts feeling like a colleague sitting next to your documents.
The whole structure, visible at last: the skills, context, operations, projects, and the map. The trick: right-click the panel, Move Panel Right. If you don’t have Claude Code yet, install it now from the official guide ; it’s one command pasted into that terminal. Then type claude and press Enter. Claude Code starts inside the folder, reads the map file, and knows where everything lives from your first message.
Files on the left, agent on the right. Step 4: Say hello, then learn four commands
One distinction saves beginners a lot of confusion. The terminal wants exact commands; Claude wants normal language. Before you typed claude , you were talking to the terminal, where spelling and syntax matter. Now that Claude Code is running, you’re talking to Claude, exactly like in the chat app. So do the natural first thing: type hello and ask it what it can see in this folder. It will answer like a colleague who just read your files.
A few slash commands are worth knowing on day one, typed directly into Claude Code:
/model - shows which Claude model you’re running and lets you switch.
/effort - sets how much thinking Claude puts into each answer; higher effort for hard work, lower for quick drafts.
/mcp - lists your MCPs, the live connections into tools like Jira, Slack, or Confluence, and lets you manage them.
/plugins - opens the plugin system, and this one deserves a minute. Claude Code comes with Anthropic’s plugin directory available by default, and marketplaces can be added on top. One is built specifically for product managers: pm-skills , a hundred-plus PM skills across discovery, strategy, execution, and launch; its README has the up-to-date install instructions. pm-execution covers PRDs, OKRs, roadmaps, and retrospectives; pm-product-discovery adds interviews and assumption testing. Install the two or three that match your work, not all nine; every installed plugin is context the system carries.
The plugin panel: Anthropic’s marketplace is there by default. One more control, then you’re done with the tour: press Shift+Tab to cycle Claude Code’s three modes.
Default mode asks your permission before touching any file.
Auto-accept applies edits without asking.
Plan mode is the one to remember: Claude researches and proposes a plan, and nothing changes until you approve it.
Start in default, use plan mode for anything big, and save auto-accept for work you’ve learned to trust. For everything beyond these, this community cheat sheet is the best one-page reference I’ve found.
Say hello first. The commands can wait a minute. Step 5: Run the setup skill
Type /setup-workspace and press Enter. The workspace ships with a skill whose only job is to personalize the template: it interviews you about your company, your product, your users, and how you like to work, then fills in the context files and renames the example project to your most active real one. Answer in plain language; it does the writing and shows you each file for a yes or a fix.
You don’t have to answer from memory: paste links. Hand it your product’s public site, a documentation page, or, if you connected the Atlassian MCP in the previous step, the Confluence page about your team, and it reads them and builds the context files from what’s already written. Fifteen minutes of this replaces the blank-page problem entirely. Your exact questions will differ a little from mine; each session is a little unpredictable, that’s just how these models work. Skip anything you don’t know yet; an honest placeholder beats invented content.
Answer in plain language, or paste a link and let it read. Step 6: Give it a real task
Don’t end the first session on setup; hand it something you actually owe someone. Type /status-update and watch it draft your update from the project’s task file instead of your memory. Or point it at a real deliverable:
Draft the kickoff brief for [your project]. Match my writing preferences in context/preferences.md . Check the open questions in the project’s tasks.md before concluding anything.
The first output that sounds like you, about your product, with your project’s real state in it, is the moment the filing clicks. Three more skills wait in the folder: prd-review for the next document you’d like a hard read on, interview-synthesis for your next user conversation, and file-feedback , the one that makes your corrections permanent.
The skill reads the workspace state first, then asks who the update is for. That’s the start: a working system in one sitting. What makes it compound is how you feed it, and that’s six practices.
The starter structure encodes the practices I converged on by getting each one wrong first. They’re not a methodology to obey; they’re the moves that kept earning their place. Each one below names the concrete thing to do and what you get back for it. The first one is the engine; the rest steer.
Use it for every task, every day
The single highest-return habit is also the simplest t

[truncated]
