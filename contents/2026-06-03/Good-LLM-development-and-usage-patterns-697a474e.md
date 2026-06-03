---
source: "https://blog.bluebyday.com/posts/good-llm-dev-and-usage/"
hn_url: "https://news.ycombinator.com/item?id=48367934"
title: "Good LLM development and usage patterns"
article_title: "Good LLM Dev and Usage Patterns | Author"
author: "_QrE"
captured_at: "2026-06-03T00:50:10Z"
capture_tool: "hn-digest"
hn_id: 48367934
score: 3
comments: 0
posted_at: "2026-06-02T09:37:25Z"
tags:
  - hacker-news
  - translated
---

# Good LLM development and usage patterns

- HN: [48367934](https://news.ycombinator.com/item?id=48367934)
- Source: [blog.bluebyday.com](https://blog.bluebyday.com/posts/good-llm-dev-and-usage/)
- Score: 3
- Comments: 0
- Posted: 2026-06-02T09:37:25Z

## Translation

タイトル: 優れた LLM の開発と使用パターン
記事のタイトル: 優れた LLM の開発と使用パターン |著者

記事本文:
著者@bluebyday.com ~ $
～/について
優れた LLM の開発と使用パターン
これは、私が観察した、ポジティブな結果をもたらす LLM (大規模言語モデル) の使用に関するパターンのリストです。それらを 2 つのカテゴリに分けました。
部分的にエージェント的なシステムで使用されるパターンの使用法。つまり、システムは動作中に LLM を利用します。与えられた例はコーディング エージェントに焦点を当てていますが、一般的に適用できます。
開発、ソフトウェア開発中に使用されるパターン用
カテゴリの分割が恣意的であると思われる場合は、(無関係な) 見出しを無視してください。
アクター/批評家パターンでは、入力が「アクター」と呼ばれる 1 つの LLM エージェントに与えられ、その出力が「批評家」と呼ばれる別の LLM エージェントに供給されます。このプロセスは反復的です。批評家は何を受け入れるかを厳しくするように指示されており、その基準を満たさない場合、俳優は批評家のフィードバックを受け、より良い結果を生み出すよう求められます。
これは、「4 つの目」メソッド/ペア プログラミングと同様に、出力の品質を大幅に向上させるための高価な方法になる可能性があります。使用されるモデルとコンテキストの処理は実装の詳細であることに注意してください。批評家が完全に承認するまで、エージェントが毎回変更のセットを減らしてアプローチを徐々に改良していくように、システムによって行われた決定を繰り返しにわたって継続的に追跡することをお勧めします。
これの代替 (予算が無限にある場合はそれを補う) 方法は、各プロンプトに対して複数の応答を生成し、最適なものを選択することです。私は、これはあまり良い方法ではないと考えています。拡張性があまり高くなく、非常によく似たものを表現するさまざまな方法の中から選択しようとするよりも、反復的な改善の方が一般的に良い結果をもたらします。

熱い結果。したがって、クリティカル + 追加の反復に費やすトークンは、特定のステップの複数の置換にトークンを費やすよりも効率的に費やされます。
各タスクを実行するために必要なコンテキストを最小限に抑える
人間は仕事をタスクに分割します。それぞれが、より大きなソフトウェアを作成するための構成要素です。エージェントも、一度に 1 つの (比較的) 小さなタスクを達成するために必要な情報のみを受け取る必要があります。
実際には、これは、エージェントが理解するために読む必要のあるドキュメントが簡潔であり、エージェントのハーネスにはエージェントが仕事を行うために必要なツールのみが含まれており、QA / テスト / 導入は開発とは別の関心事であることを意味します。各エージェントがやらなければならないことが増えれば増えるほど、割り当てられたタスクを忘れたり、逸脱したりする機会が増えます。
これは、適切なプライバシーとセキュリティの実践とも結びついています。エージェントが読んだり提供したりする必要はありません。ユーザー ID を使用してツールを呼び出すと、ソフトウェアは呼び出し時にその情報自体を提供できます。これにより、ツールを正常に呼び出すために必要なトークンが少なくなるため、エージェントはツールを呼び出すことが容易になり、またエージェントが機密情報を処理する必要がなくなります。
何かを達成するためにエージェントが最小限のトークンを出力する必要があることを確認すると、成功の可能性が高くなります (少なくとも一般的には利点は減少します)。
一般に、エージェントが (少なくとも呼び出しごとに) 実行することが期待されないほど、エージェントが許容可能な結果を​​生成する可能性が高くなります。構築中のソフトウェアが形になり始めると、たとえ小さな変更を加えるために必要なコンテキストも自然に増加します。増加を最小限に抑えることで、エージェントはより安く (コンテキストのトークンが減り)、より速く (処理するトークンが減って)、より良く (可能性が高くなります) 作業できるようになります。

良い結果が出ると良いですね）。
あまり前向きではないかもしれませんが、エージェントができることは、（最終的には）実行することが期待されるべきであるということを覚えておいてください。したがって、あなたを吊るすのに十分なロープをエージェントに決して与えないでください。そうされる可能性も不可能ではないからです。
エージェントが、特に指示を誤って解釈した場合には、与えられたものを悪用することを予期することが重要です。一部のプロバイダーは、エージェントにユーザーを送信する前に、2 番目の LLM にユーザーのプロンプトを評価させることを提案しています。これにより攻撃はより困難になりますが、不可能ではありません。また、それ自体を堅牢な防御と見なすべきではありません。より良いアプローチは、非常に有害なツールへのアクセスをエージェントに決して与えないこと、ツールにユーザーの確認を要求すること、または発生した損害を簡単に元に戻せる環境でエージェントを動作させることです (技術的には、エージェントに提供されるツールは極端に有害なものではありませんが、サンドボックスについて明示的に言及しないと、誰かが私が忘れていると思うかもしれません)。
人間はすぐに警戒に疲れてしまうため、頻繁に要求される場合には、ツールにユーザーの確認を要求することは理想的ではないことに注意してください。 2 番目の目を必要とするほど影響力のあるエージェントのアクションのみを人間に確実に警告することができない場合は、代わりに、最初からエージェントに有害なツールへのアクセスを与えないことを選択します。
防御性に関しては、有効なエージェントの出力がどのようなものであるかについて、決定論的にチェックできるほど単純な制約を課すこと (つまり、フィールドが数値のみを取ることを強制すること) は、ユーザーにエラーを表面化させることなく、エージェントが犯し得る単純な間違いをキャッチし、適切にエージェントに再プロンプトを表示できることを意味します。構造化された出力をエージェントから強制的に出力することは、それをより簡単に解析するのにも役立ちます。これは非常に重要です。エージェントに具体化してもらうことは、
堅牢性の原則
マックを助ける

システムが堅牢であることを確認してください。
コンテキストの最小化と防御の両方の簡単な例は、エージェントが選択したバージョン管理ソフトウェア (VCS) に一連の変更をコミットした後で、コーディング エージェントのハーネスにフォーマット、リンティング、静的分析ツールを実行させることです。これにより、エージェントは品質/テスト パイプライン自体を実行する責任が免除され、ユーザーはエージェントが適切な開発慣行に従うことを信頼する必要がなくなります。
最新かつ最高のものを使用する必要はありません
State-Of-The-Art (SOTA) LLM は非常に優れています。それらは非常に高価でもあります。これらを使用する必要があるタスクがない場合は、より安価なものを選択してください。問題に対してエージェント以外の解決策をエージェントに提示してもらうことができれば、さらに良いでしょう。
多くの場合、より優れたハーネス/プロンプト/プロセスにより、LLM のパフォーマンスが大幅に向上します。誰かがトークンの使用を強制していない限り、支出に注意する必要があります。LLM は無料ではなく、コスト効率の高いソリューションを使用すると、請求書のゼロが 1 つ減ってしまう可能性があるためです。
場合によっては、プロセスに LLM を追加する必要がまったくない場合もあります。一歩下がってみると、現在存在しているものよりもシンプルで効率的なプロセスやシステムを構築するためのより良い方法が見つかるかもしれません。プロセスを実行するためにインテリジェンスが必要なくなるまでプロセスを簡素化できれば、自動化は AI を使用するよりも安価で、高速で、信頼性が高くなります。
エージェントの動作に応じて、実装できるコンテキスト削減戦略が多数あります。コンテキスト ウィンドウの制限に達しないと考えるのは悪い習慣です。圧縮は明示的に考慮する必要があります。コンテキストの量を低く保つと、より良い結果が得られる傾向があります。最小限の指示の方が、よく従うのが簡単です。
ツールの呼び出しと結果は次のとおりです (通常、

y) 付随する思考の流れよりも有用性が低い。出力内の各トークンの値に基づいて圧縮する場合、最後の X 呼び出しを除くすべてのツール呼び出しと出力を削除すると、一貫性を長期間維持するのに役立ちます。
すべてのタイプのコンテキスト圧縮には損失が伴います。すべての情報/意味を保持しながらテキストを圧縮できたとしても、セマンティクスが変更され、将来の応答に影響を与える可能性があります。したがって、圧縮をできるだけ遅らせなければならないというプレッシャーがあります。エージェントがタスクを「チャンク」に分けて作業する場合は、最新の推論を保存してエージェントの現在のコースをより適切に維持するために、最新の X メッセージを除くすべてのメッセージを要約することが望ましいです。アプローチが異なればコストも異なるため、キャッシュされたトークンの仕組みと価格設定に留意することが重要です。
多くは予算とユースケースに依存するため、他のすべてよりも明らかに優れた圧縮戦略はありません。必要に応じて、メッセージの削除、トークンの削除、要約、ユーザーに新しいセッションの開始を求めることなどを行うことができます。
計画の失敗は失敗する計画だ
仕事を始める前に、成功とはどのようなものかをよく理解する必要があります。そうしないと、エージェントが実行するのに十分な包括的な計画を立てることができなくなります。エージェントに実行を依頼する前に包括的な計画を作成しないと、不適切な解決策や不完全な解決策が得られることになります。おもちゃのプロジェクトの場合は、それが許容される可能性があります。運用環境の場合、これは、エージェントが SaaS アプリで匿名サインアップ/ログインをオフにするのを忘れることを意味し、作成者がその存在を忘れていた JavaScript ライブラリを使用して 10 の異なる方法で接続し、0100 でクラッシュすると、目が覚めたときにすべてを最初から再デプロイする必要があり、時折データが失われると、うっかり GDPR 義務を履行してしまうことになります。
再言する

もう少し明確に: 使い捨てのスクリプト以上のものを必要とするタスクの場合は、包括的な計画を作成し、あなたもエージェントもこれ以上検討すべきものが見つからなくなるまでそれを繰り返します。道路に関する慣用句をもう 1 つ紹介します。1 オンスの予防は 1 ポンドの治療に値します。
開発に関しては未知の部分がたくさんあります。アプローチの側面をテストするための概念実証 (PoC) スクリプトを大量に作成するのは低コストであり、本番環境に対応したコードの作成に着手する前に問題が表面化する可能性があります。 PoC は、たとえばライブラリ間のパフォーマンスや使いやすさの違いを実証できるため、特定の決定が行われた理由を文書化するのにも役立ちます。
あなたが意見を無視した決定は、エージェントがあなたのために決定します。全体的な計画が十分に明確になっていない場合、一部の決定によって技術的負債の蓄積が加速され、コード チャーンが増加する可能性があります。何が重要で何がそうでないかをよく理解することが重要です。重要なものはテスト、検証、文書化する必要がありますが、重要でないものはエージェントの判断に任せることができます。
レビューや修正を頻繁に行う
LLM はコードのレビューを (手助けして) 行うこともできます。コードを確認してからしばらく経っている場合は、LLM に、そのコードがうまく書かれているか、バグはないか、改善できるか、用心深く勤勉なシニアプリンシパルの超 10 倍のロックスター エンジニアが承認するかどうか、今日書かれたものであるかどうかを尋ねてください。これを数回行うと、修正する価値のある問題があれば、そのほとんどが明らかになるはずです。人間も LLM も間違いを犯します。
さらに目玉を追加すると、これまで見落とされていたバグが表面化します
。コードは一度書いたら放置する必要はありません。維持費も安くなりましたので、ぜひご利用ください。
ベンチマークは、より高い「スコア」を取得するために研究室によってゲーム化される可能性があり、実際に行われています。ある場合には

ベンチマークの内容がトレーニング データに入り込む可能性があり、ベンチマークの結果が疑わしいものになります。
LLM は情報を非常によく思い出すことができます
。場合によっては、
ベンチマークの構築が不十分で、エージェントが「不正行為」を行う可能性があります。
いずれの場合も、ベンチマークは、ユーザー自身の使用法と一致しない可能性が高い方法で LLM をテストします。したがって、どのモデルが自分のユースケースに最も適合するかをテストすることが賢明です。そして最も重要なことは、支出を過剰/過少にすることなく、目的に十分に適したモデルを特定して使用できるように、モデルをスコアリングする何らかの手段を用意することです。
採点は雰囲気に基づいてはいけません。優れたベンチマークを作成するには多くのことが必要であり、多くの詳細はユースケースに応じて異なります。トークンのコストがいずれかの時点で重要な要因となる場合 (通常、最終的にはそうなります)、どの LLM が自分にとって最も効率的かを理解することに投資する価値はあるでしょう。最新の LLM では、少なくとも受け入れ基準が定義されているタスクでは、より多くのトークンを使用すると、ほぼ常に良い結果が得られるという点に到達していることに注意してください。当然のことながら、トークンごとに表現できる量には制限があります。時間とお金の両方の予算を必ず把握し、業務全体で一貫性を保つようにしてください。

[切り捨てられた]

## Original Extract

Author@bluebyday.com ~ $
~/about
Good LLM Dev and Usage Patterns
This is a list of patterns regarding usage of LLMs (Large Language Models) that I’ve observed result in positive outcomes. I’ve split them into two categories:
Usage, for patterns used in systems that are partly agentic, i.e. the system utilizes LLMs during its operation. The examples given are focused on coding agents, but they are generally applicable.
Development, for patterns used during software development
If the category split feels arbitrary, feel free to ignore the (ir)relevant headings.
The actor/critic pattern is where input is given to one LLM agent, called the ‘Actor’, and the output of that is fed into a different LLM agent, called the ‘Critic’. The process is iterative; the Critic is instructed to be stringent in what they accept, and failure to meet their standards results in the actor being given the Critic’s feedback, and asked to produce a better result.
This can be an expensive way to dramatically increase the quality of output, similar to the “four eyes” method / pair programming. Note that the model(s) used, and any processing of the context, are implementation details. It is a good idea to keep track of decisions made by the system in a way that persists across iterations, such that the agents gradually refine the approach with a decreasing set of changes each time, until the Critic fully approves.
An alternative (or complement, if you have infinite budget) of this is to generate multiple responses to each prompt, and pick the best one. I believe that is not good practice, as it does not scale very well, and iterative improvements generally pay off better vs attempting to pick between what most often is different ways of expressing a very similar one-shot result. The tokens you spend on the critic + an additional iteration are therefore more efficiently spent than the tokens on multiple permutations of any given step.
Minimize the context needed to accomplish each task
Humans split work into tasks. Each one is a building block towards creating a bigger piece of software. Agents too should only receive the information they need to accomplish one (relatively) small task at a time.
In practice, this means that the documents that the agent needs to read to get up to speed are concise, the agent’s harness only has the necessary tools that the agent needs to do their job, and QA / testing / deployment are separate concerns from development. The more each agent has to do, the more opportunities they have to forget or deviate from their assigned tasks.
This also ties in with good privacy & security practices: You don’t need the agent to read or provide e.g. user IDs to call a tool, your software can provide that info itself when the call is made. This both makes it easier for the agent to call a tool, because less tokens are needed to call it successfully, and prevents the agent from having to handle confidential information.
Making sure that the agent needs to output the minimum number of tokens to accomplish something makes the likelihood of success higher (in general at least, with diminishing benefits).
In general, the less the agent is expected to do (at least, per invocation), the higher the likelihood of the agent producing an acceptable result. As the software being built starts to take shape, the context needed to make even a small change will naturally increase. Keeping the increase minimal ensures that agents can work cheaper (less tokens of context), faster (less tokens to process) and better (greater likelihood of producing good results).
On a less positive note, remember that what an agent can do, it should be expected to (eventually) do. So never give any agent enough rope to hang you with, because it is not impossible that it might.
It is important to expect that the agent will abuse anything given to it, especially if it ever misinterprets an instruction. Some providers suggest having a second LLM assess users’ prompts before sending them to your agent. This does make attacks harder, but not impossible, and it should not be considered a robust defense on its own. A better approach is to never give access to extremely damaging tools to agents, make the tools require user confirmation, or have the agent work in an environment where any damage that it does can be easily reverted (which technically makes the tools provided to the agent not extremely damaging, but if I don’t mention sandboxes explicitly someone might think I forgot).
Do note that having tools require user confirmation is not ideal if it will be frequently requested, as humans quickly get alert fatigue. If you cannot reliably alert humans to only the agents’ actions that are impactful enough to require a second set of eyes, instead opt for not giving the agent access to damaging tools in the first place.
Regarding defensiveness, imposing constraints on what valid agent output looks like, that are simple enough to be checked deterministically (i.e. enforcing that a field only takes a number), means that you can catch simple mistakes that the agent could make and re-prompt the agent gracefully, without surfacing errors to the user. Forcing structured output out of agents also helps with parsing it more easily, which can be extremely important - having agents embody
the robustness principle
helps with making sure that your system is robust.
An easy example of both minimizing context and being defensive is having a coding agent’s harness run formatting, linting, and static analysis tools once the agent commits a set of changes to the Version Control Software (VCS) of choice. This absolves the agent of the responsibility of running the quality/testing pipelines itself, and the user from having to trust the agent to follow good development practices.
You don’t have to use the latest and greatest
State-Of-The-Art (SOTA) LLMs are very good. They are also very expensive. If you do not have a task that warrants using them, opt for something cheaper. If you can have agents produce a non-agentic solution to the problem, even better.
A lot of times, a better harness / prompt / process can provide a meaningful boost to the performance of LLMs. Unless someone is forcing you to spend tokens, you should be mindful of your expenditure, because LLMs are not free, and cost-efficient solutions could mean your bill has one less zero tacked on.
In certain cases, there may not be a need to add LLMs to a process at all. Taking a step back might reveal a better way to architect a process or a system that is simpler, and more efficient, than what currently exists. If you can simplify a process to the point where no intelligence is needed to perform it, automating it will be cheaper, faster, and more reliable than using AI.
There are many context reduction strategies that can be implemented, depending on what the agent is doing. Assuming that you will not hit the context window limit is bad practice; compaction should be an explicit consideration. Keeping the context amount low tends to give better results; minimal instructions are easier to follow well.
Tool calls and results are (typically) less useful than the accompanying thought stream. If you compact on a value-of-each-token-in-the-output basis, removing the tool calls and outputs of all-but-the-last-X invocations helps maintain coherency for longer.
All types of context compaction are lossy; even if the text can be compacted while preserving all the information/meaning, the semantics will change, which might have an impact on future responses. So there is pressure to delay compaction as much as possible. Summarizing all but the latest X messages is preferable if the agent works on a task in ‘chunks’, to try and preserve recent reasoning and maintain the agent’s current course better. Different approaches will have different costs, and minding how cached tokens work & are priced is important.
There is no compaction strategy that is clearly better than all the others, since a lot will depend on budget and use-case. You can drop messages, drop tokens, summarize, ask users to start a new session, or anything else as appropriate.
A failure to plan is a plan to fail
You need to have a good idea of what success looks like before you begin work. If you do not, you will not be able to draft a comprehensive enough plan for agents to execute. If you do not draft a comprehensive plan before you ask agents to execute, you will get bad and/or incomplete solutions. For toy projects, that can be acceptable. For production stuff, that will mean that your agent will forget to turn off anonymous signup/login in your SaaS app, it will wire things ten different ways using JavaScript libraries whose creators forgot they exist, a crash at 0100 will result in having to re-deploy everything from scratch when you wake up, and the occasional data loss will mean that you inadvertently meet your GDPR obligations.
Restated a bit more clearly: For tasks that warrant more than a throwaway script, build a comprehensive plan, and iterate on it until neither you nor your agent(s) can find anything more (of substance) to consider. One more idiom, for the road: An ounce of prevention is worth a pound of cure.
There are a lot of unknown unknowns when it comes to development. Creating a lot of Proof-of-Concept (PoC) scripts to test aspects of the approach is cheap, and can surface issues before committing to writing production-ready code. PoCs can also help document why certain decisions were made, as they can, for example, demonstrate performance and usability differences between libraries.
Any decision you neglected to opine on, agents will decide for you. If your overall plan is not made clear enough, some decisions could accelerate the accumulation of tech debt and result in increased code churn. It is important to have a good idea of what is important and what is not. What is important should be tested, validated, and documented, what is unimportant can be left for the agent to decide.
Frequently do reviews and revisions
LLMs can also (help you) review code. If it’s been a while since a piece of code was looked at, ask an LLM if it’s well-written, if it has any bugs, if it can be improved, if a cautious and diligent senior-principal super 10x rockstar engineer would approve of it, if it was written today. Doing that a couple of times should surface most of the issues worth fixing, if any. Both people and LLMs make mistakes;
adding more eyeballs surfaces bugs that were previously overlooked
. Code does not have to languish abandoned once written; maintenance is now cheaper, so we should take advantage.
Benchmarks can be, and are, gamed by labs to get a higher ‘score’. In certain cases, content from the benchmarks could find its way into the training data, which makes the results of benchmarks dubious as
LLMs can recall information very well
. In some cases,
benchmarks may be poorly constructed, allowing agents to ‘cheat’.
In all cases, benchmarks test LLMs in a way that most likely will not match your own usage. It is therefore wise to test what models fit your use case the best, and most importantly, have some means of scoring them such that you can identify and use whatever model is good enough for your purposes, without over/under spending.
Scoring should not be based on vibes. There is a lot that goes into making a good benchmark, and a lot of details are going to vary based on use-case. If token costs are going to be a factor at any point (and they typically become one eventually), investing in figuring out what LLM is the most efficient for you is going to be worthwhile. Note that we have reached a point where for modern LLMs, using more tokens will almost always give better results, at least for tasks with defined acceptance criteria. There is, naturally, a limit to how much you can express per token. Make sure that you are aware of both your time and money budget(s), and be consistent with these across your ben

[truncated]
