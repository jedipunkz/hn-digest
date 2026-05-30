---
source: "https://nocodefunctions.com/blog/three-flavors-agentic-coding/"
hn_url: "https://news.ycombinator.com/item?id=48334332"
title: "Three flavors of coding with AI agents"
article_title: "Three flavors of coding with AI agents – Nocode functions - blog – The journey of an academic and app developer"
author: "seinecle"
captured_at: "2026-05-30T11:34:08Z"
capture_tool: "hn-digest"
hn_id: 48334332
score: 2
comments: 0
posted_at: "2026-05-30T09:27:02Z"
tags:
  - hacker-news
  - translated
---

# Three flavors of coding with AI agents

- HN: [48334332](https://news.ycombinator.com/item?id=48334332)
- Source: [nocodefunctions.com](https://nocodefunctions.com/blog/three-flavors-agentic-coding/)
- Score: 2
- Comments: 0
- Posted: 2026-05-30T09:27:02Z

## Translation

タイトル: AI エージェントを使用した 3 種類のコーディング
記事のタイトル: AI エージェントを使用したコーディングの 3 つのフレーバー – ノーコード機能 – ブログ – 学者とアプリ開発者の旅
説明: 少なくともエージェント コーディングの文脈では、「AI エージェント」の合理的な定義は次のようになります。
LLM の機能を備えたソフトウェア プロセス
タスクを達成するために開始時に与えられた指示に従って起動される
これは自律的に実行されます (人間との対話型セッションはありません)。
[切り捨てられた]

記事本文:
学者とアプリ開発者の旅
AI エージェントを使用した 3 種類のコーディング
「AI エージェント」の合理的な定義は、少なくともエージェント コーディングの文脈では次のようになります。
LLM の機能を備えたソフトウェア プロセス
タスクを達成するために開始時に与えられた指示に従って起動される
かなりの期間にわたって自律的に実行されます (人間との対話型セッションはありません)。
非決定的な動作: エージェントは、できれば受け取った指示から逸脱することなく、状況に適応します。
これらのソフトウェア プロセス (エージェント) は、より迅速な結果を達成したり、同じプロセスを複数のコピーで起動したり、さまざまなプロセスを一度に起動したりすることで、より多くのタスクを実行するために並行して起動できます。
必須ではありませんが、論理的な次のステップです。タスクを達成するために、プロセスは他のプロセスやサブプロセスなどを起動するように誘導できます。これは、タスクを達成するために分散型の方法で調整する (ループ内に人間がいない) カスケード、軍隊、またはエージェントの群れのイメージを思い出させます。
しかし、実際には、「エージェント」は上記の定義とは無関係に使用されることがよくあります。おそらく最新で洗練されたように聞こえるかもしれませんが、「エージェント」は実際には ChatGPT 会話を指定する可能性があり、そこではユーザーが「あなたはプロの税務代理人のように振る舞っています。以下では、納税申告書を記入するのを手伝ってほしいのですが」 🤷‍♂️ と促します。
コーディングと AI の問題について率直に話すことが多いピーター レベルズ氏は、2025 年の夏に次のように感じました。
最近人々が「AI エージェント」について話しているのを聞いたら、それは一般的に危険信号であり、彼らが AI ニュースを読んでいるだけで実際には何も出荷していない非技術者であることはわかっています。
私が AI エージェントを信じていないからではありませんが、これは現時点では本当の意味を持たないマーケティング用語です
私たちは今近づいています

2026 年の夏、状況は大きく変わりましたか?
コーディングの実践では、上で提案した定義を偽装せずに、その定義に沿ってエージェントティック コーディングを行う方法をいくつか検討しました。
フレーバー 1: 複数のコマンド ライン インターフェイスの起動
サーバーへの SSH セッションを開く
タスクを説明するプロンプトを書くだけでタスクを完了するように GPT に依頼します
サーバーへの 2 番目の SSH セッションを開きます
タスクを説明するプロンプトを書くだけで、GPT に別のタスクを実行するように依頼できます。
正直に言うと、それはかなりうまくいきます。ご覧のとおり、非常にローテクです。また、あるセッションで Claude Code を起動し、別のセッションで Codex CLI を起動し、3 番目のセッションで Gemini CLI を起動できることも意味します。そのため、トークンの消費が複数の AI プロバイダーに並行して分散され、特定のプロバイダーのトークン予算制限に到達するのが遅くなります。
フレーバー 2: AI CLI をヘッドレス モードで起動する
私はこの 2 番目のアプローチを使用して、200 以上の異なる Web ページ用のクローラーを作成しました。明らかに 200 個のクローラーを作成するには、上記のフレーバー 1 を使用するには退屈すぎるでしょう。 ChatGPT は、この新しいアプローチを実装する方法について最初から最後までガイドしてくれました。基本的なロジックは次のとおりです。
200 の Web サイトのパラメーター (URL とその他の詳細) を含む 1 つの JSON ファイル。
1 つの Bash スクリプト (「A」と呼びます) は、ヘッドレス モードでコマンドライン インターフェイスを介して 1 つの LLM を起動できます。ヘッドレスとは、ユーザーが指定したプロンプトで LLM が起動されると、許可を求めたり、フィードバックやフォローアップを求めたりするために中断することなく、タスクが完了するまで実行されることを意味します。そのために、ヘッドレス モードをトリガーする Codex CLI の exec フラグを使用します。スクリプト A には、LLM の起動時に表示されるプロンプトも含まれています。プロンプトは重要な箇所にプレースホルダーを含むテキストであり、実際の情報に置き換えられます。

n クロールされる特定の Web サイトに関連する。プロンプトは基本的に、LLM にこの Web サイトのクローラーを作成するように求めます。
別のスクリプト (スクリプト「B」) は、json ファイルから 20 の Web サイトを選択し、それぞれに対してスクリプト A を実行します。スクリプト A のプレースホルダーは、クロールされる Web サイトの情報に置き換えられます。これは、LLM によって作成されたクローラーがこの Web サイトに固有であることを意味します。
スクリプト B を起動し、正常に動作することを確認してから、200 個の Web サイトをこの方法で処理するまで、他の 20 個の Web サイトでスクリプト B を再起動します。
このアプローチのフレーバー 2 がフレーバー 1 よりもどのように複雑になるかを説明するために、スクリプト A (ChatGPT によって作成されたスクリプト) を示します。
このアプローチはうまく機能します。 「スクリプト B を起動し、1 時間で 200 個のクローラーを作成する」ほど簡単ではありませんが、ほぼそれに近いものです。上記のスクリプトを辛抱強く少し読んでいただければ、LLM が作成する各クローラーの単体テストを作成するタスクを負っていることがわかります。予想どおり、これらのテストは常に合格するとは限らないため、処理が少し遅くなります。しかし、これには十分な理由があります。テストに合格するために必要な追加の作業を行うことは、クローラーの信頼性が高まることを意味します。
このアプローチにより、今後数日以内に 200 個のクローラーが稼働し、準備が整い、さらに数百個まで拡張する簡単な方法が得られると予想しています。
フレーバー 3: 1 つの LLM にこれらのサブエージェント自体を作成および管理させる
フレーバー 2 は Bash と Unix が非常に重かったため、プロセスの保守が困難になりました。私の指示に従って、LLM だけでエージェントを起動させてみてはいかがでしょうか?最近、あらゆるソリューションが宣伝しているのは次のようなものです。
カーソルは「より高いレベルの指示に集中するために実装を委任する」よう促します。
Codex には、調整できる「サブエージェント」があります
Google の Antigravity は、「独立したプロ全体で並行して動作する複数の自律エージェントを組織化する」ことを提案しています。

ジェクト。」
Claude Code は「カスタム サブ エージェント」を作成できます。
私の意見：おそらく、しかし今日はそうではありません。 1 人のエージェントにサブエージェントに委任するよう依頼するということは、実際の作業から 2 ステップ削除されることを意味します。矛盾、不適切な選択、単純なエラーなどを見つけるのは難しくなります。特定のサブエージェントの作業を中断したり再開したりするのは簡単ではありません。そして、ソリューションに依存するようになります。最近私が選んでいる AI は GPT 5.5 ですが、OpenAI 社が開発したもの以外のエージェントを含むソリューションを選択すると、これは使用できなくなります。
これらの理由から、そうでないと証明されるまで、私は上記のフレーバー番号 2 (単純な場合は番号 1) を使い続けることにします。
ほとんどの場合、いいえ。以下は、Ethan Mollick が、プロンプト 1 つとフォローアップ 4 つだけ、合計 20 行未満の完全なライブ Web アプリケーションを作成しているところです。
あなたは、いつ、どこで生まれたことがどれほど幸運でしょうか?
クロード・コードの作品 4.8 が、これまでに生きたすべての人間の新たな視覚化を作り上げたとしたら。これは、きちんとしているだけでなく、AI のリサーチ、コード、設計、統計を組み合わせた興味深いテストでもあります。 https://t.co/ayNEdhSLy3 pic.twitter.com/Ny2NmICZsK
もう 1 つの例は、娘が開発した本格的な電子商取引プラットフォームの Web サイトです。派手なテクノロジーやエージェントの足場を一切使用せずに開発されました。プロンプトを表示するだけです (そして多くの作業が必要です)。
おそらく、情報を検索する際の LLM とのチャットで、これらの LLMS は別の Web サイトで検索を開始することを選択でき、各検索は独自のエージェントによって実行されることにもおそらくお気づきでしょう。これは、調査をスピードアップし、リクエストに応じてより多くの領域をカバーするのに役立ちます。
したがって、ほとんどの場合、LLM はエージェントがなくても問題なく機能するため、複雑なタスクであってもエージェントは必要ありません。エージェントが役立つ場合は、LLM が独自のエージェントを起動して管理します。

フード。
エージェント同士がぶつかる難しさ
このトピックは単に地味で、ブログ投稿がすでに長すぎるため、非常に簡単に説明します。コードベースに同時に書き込みを行う複数のエージェントは、お互いに足踏みをします。同じファイルに変更を加えた場合、結果のファイルが混乱する可能性が非常に高くなります。
解決策は、各エージェントを異なる git ブランチで動作させ、すべてのエージェントが完了したときにブランチをメイン ブランチにマージする作業に進むことです。これも厄介です。マージが失敗したらどうなるでしょうか? （そして、ああ、そうなります）。私はこのアプローチを試してみましたが、退屈で身の毛がよだつようなもので、マルチエージェント ゲームをすぐにやめてしまいます。
では、Flavor 2 では、エージェントがお互いの作業に衝突することなく、どうやって数十のクローラーを作成できたのでしょうか?
私はまず 1 つのクローラーに対して多くの準備作業を実施し、それが他のクローラーから完全に分離されて動作することを確認しました。コーディングにおいて DRY (同じことを繰り返さない) 原則に従う必要がある場合、これは簡単な作業ではありません。各クローラーが完全に分離されているこの条件下でのみ、混乱を生じさせることなく、数十のエージェントがソース ファイルを同時に処理できるようになります。
上にスクロールして、私が共有した Bash スクリプト (「スクリプト A」) を確認すると、この問題について ChatGPT からもアドバイスを受けていることがわかります。ChatGPT はプロンプトにいくつかのハード ブロックを追加し、各エージェントがその作業範囲外のファイルに触れることが明示的に禁止されています。
Flavor 2 で作成された数十のクローラーから数百のクローラーに変換します。次に、それらを実行します。この「自家製」マルチエージェント設定に十分に習熟すると、必要に応じて他の場所でもそれを再現できるようになります。
どのようなマルチエージェント設定が効果的ですか?それとも、まったくエージェントを使わないことにこだわりますか?
私は学者で独立した Web 専門家です

ppの開発者。私は、テキストとネットワークを探索するためのポイント アンド クリック ツールである nocode 関数を作成しました。試してみて、ご意見をお聞かせください。ぜひフィードバックをお待ちしています!
電子メール: Analysis@exploreyourdata.com
ブログ: アプリ開発とデータ探索に関する記事をさらにお読みください。
必要に応じて、ここで新しい情報に関するメールを不定期に受け取ります。
関連投稿:
エージェントの平均管理者 IA
Signal から Debian で Claude と Gemini を操縦する
NetBeans からクロード コードまで
前へ
⏪ エージェントの平均管理者 IA
このサイトは Clement Levallois の <3 を使用して作成されています。

## Original Extract

A reasonable definition of an “AI agent”, at least in the context of agentic coding, could be:
a software process endowed with the capabilities of an LLM
launched with instructions given at the start to accomplish a task
which runs autonomously (no interactive session with a human), for a significan
[truncated]

The journey of an academic and app developer
Three flavors of coding with AI agents
A reasonable definition of an “AI agent”, at least in the context of agentic coding, could be:
a software process endowed with the capabilities of an LLM
launched with instructions given at the start to accomplish a task
which runs autonomously (no interactive session with a human), for a significant period of time
with non-deterministic behavior: the agent adapts to the circumstances, hopefully without deviating from the instructions it received
These software processes (agents) can be launched in parallel to achieve faster results or to accomplish a larger number of tasks: same process launched in multiple copies, or a variety of processes launched at once.
Not a necessity but logical next step: to accomplish a task, a process can be led to launch other processes, subprocesses, etc. This evokes images of cascades, armies, or swarms of agents coordinating in a decentralized manner (no human in the loop) to accomplish a task.
Yet, in practice, “agents” is often used with no relation to the definition above. Possibly to sound up-to-date and sophisticated, “agents” might in reality designate a ChatGPT conversation where the user prompted “you are acting like a professional tax agent and in the following, I want you to help me fill in my tax declaration” 🤷‍♂️.
Pieter Levels, who tends to speak frankly about coding and AI matters, shared this feeling in summer 2025:
If I hear people talk about "AI agents" these days it's generally a red flag and I know they're non-technical ppl reading AI news but not actually shipping anything
Not cause I don't believe in AI agents but it's such a marketing term with no real meaning at this point
We are now approaching summer 2026, have things changed much?
In my coding practice, I explored several ways to do agentic coding that would live up to the definition proposed above, and not masquerade for it.
Flavor 1: launching multiple command line interfaces
open an SSH session to my server
ask GPT to accomplish a task just by writing a prompt describing the task
open a second SSH session to my server
ask GPT to accomplish another task just by writing a prompt describing the task
Honestly, that works pretty well. It is extremely low-tech as you can see. It also means you can launch Claude Code in one session, Codex CLI in another, Gemini CLI in a third one… and hence spread your token consumption across several AI providers in parallel, which makes the token budget limit slower to hit for a given provider.
Flavor 2: launching AI CLIs in headless mode
I used this second approach to write crawlers for 200+ different webpages. Obviously with 200 crawlers to create, that would have been too boring to do with the Flavor 1 described just above. ChatGPT guided me throughout on how to implement this new approach. The basic logic is:
one JSON file containing the parameters for the 200 websites (URLs and a few more details).
one Bash script (call it “A”) that can launch one LLM through a command-line interface, in headless mode. Headless means that the LLM, once launched with the prompt you have given it, will execute until it has completed the task, without interrupting to ask you for permission or ask for feedback or a follow-up. For that I use the exec flag on Codex CLI that triggers the headless mode . Script A also contains the prompt that will be given to the LLM when it launches. The prompt is a piece of text with placeholders at key points, which are replaced by the actual information related to the specific website to be crawled. The prompt basically asks the LLM to write a crawler for this website.
another script (script “B”) that picks 20 websites from the json file and executes script A for each of them. The placeholders of script A are replaced by the info of the website to be crawled, meaning that the crawler created by the LLM will be specific to this website.
I launch script B, check that it works fine, then relaunch it with 20 other websites, etc. until I had processed 200 websites this way.
Let me show you script A (script written by ChatGPT) to illustrate how this approach Flavor 2 involves more complexity than Flavor 1:
This approach works well. It is not as easy as “launch script B, get 200 crawlers written in an hour” but almost that. If you are patient to read a bit the script above, you’ll see that the LLM is tasked to write unit tests for each crawler it creates! As expected, these tests do not always pass, so that slows things down a bit. But it is for a good reason: doing the extra work needed to get passing tests means that the crawlers will be more reliable.
With this approach, I expect to have my 200 crawlers up and ready in the next few days, and with an easy path to grow to hundreds more.
Flavor 3: having one LLM create and manage these subagents itself
Flavor 2 was really Bash and Unix heavy: this makes my processes harder to maintain. Why not have an LLM just spin up agents by itself, following my instructions? That’s what every solution is advertising these days:
Cursor invites you to “delegate implementation to focus on higher-level direction”
Codex has “sub-agents” you can orchestrate
Google’s Antigravity offers to “orchestrate multiple autonomous agents working in parallel across independent projects.”
Claude Code can create “custom sub agents” for you.
My opinion: probably, but not today. Asking one agent to delegate to sub-agents means that you are two steps removed from the actual work. Inconsistencies, poor choices, flat errors… will be harder to catch. Interruption and resuming of work for a given sub-agent is not straightforward. And you become solution-dependent: my AI of choice these days is GPT 5.5, and that would be off-limits if I choose a solution with agents that is not developed by its company, OpenAI.
For these reasons and until proven otherwise, I’ll stick with Flavor number 2 (and even number 1 in simple cases) described above.
Most of the time, no . Here is Ethan Mollick creating a complete, live web application with just one prompt and 4 follow-ups, for a total of less than 20 lines:
How lucky are you to have been born when and where you are?
Had Opus 4.8 in Claude Code whip up a new visualization of all humans who ever lived. In addition to being neat, it is an interesting test of combining research, code, design and stats for an AI. https://t.co/ayNEdhSLy3 pic.twitter.com/Ny2NmICZsK
Another example is the website developed by my daughter: a fullfledged e-commerce platform . Developed with zero fancy technology or agentic scaffolding. Just prompts (and a lot of work).
You probably also noticed that in chats with LLMs when we search for information, these LLMS can choose to launch searches on different website, with each search performed by its own agent. This helps speed up their research and cover more ground in response to your request.
So I’d say that in most cases, we don’t need agents even for complex tasks because LLMs work just fine without, and if agents are useful, then LLM launch their own agents and manage them under the hood.
And the difficulty of agents bumping into each other
The topic is just unglamorous and the blog post is already too long so I’ll be super brief: multiple agents writing simultaneously in your codebase will step on each other’s toes. If they make changes to the same file, there is a very good chance the resulting file will be a mess.
The remedy is to have each agent working on different git branches , then proceeding to the merge of the branches into the main branch when all agents are done. This is also a mess: what if the merge fails? (and oh, it will). I tried this approach and it is tedious, hair-raising and makes you quit the multi-agent game quickly.
So how did I manage in Flavor 2 to create dozens of crawlers without having agents crash into each other’s work?
I first conducted plenty of preparatory work on just one crawler, making sure it worked in perfect isolation of the others. Not an easy task when you still want to follow the DRY (don’t repeat yourself) principle in coding. Only under this condition that each crawler is perfectly isolated can you have dozens of agents working on source files simultaneously without creating a mess.
If you scroll up and check the Bash script (“script A”) I’ve shared, you’ll see I was also advised on the matter by ChatGPT, which added some hard blocks in the prompt, so that each agent is explicitly forbidden from touching files not in the scope of its work.
Getting from dozens of crawlers written with Flavor 2 to hundreds of crawlers. Then executing them. Becoming sufficiently proficient at this “homemade” multi-agent setup that I can reproduce it when and if needed in other places.
What multi-agent setup works for you? Or do you stick with no agent at all?
I’m an academic and independent web app developer. I created nocode functions , a point-and-click tool for exploring texts and networks. Try it out and let me know what you think. I’d love your feedback!
Email: analysis@exploreyourdata.com
Blog: Read more articles on app development and data exploration.
Get very occasional emails about new stuff on here, if you like.
Related Posts:
Trois manières de coder avec des agents IA
Piloting Claude and Gemini on Debian from Signal
From NetBeans to Claude Code
Previous
⏪ Trois manières de coder avec des agents IA
This site is made with <3 by Clement Levallois.
