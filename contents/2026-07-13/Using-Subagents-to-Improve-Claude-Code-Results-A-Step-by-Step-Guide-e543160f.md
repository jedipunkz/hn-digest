---
source: "https://software.rajivprab.com/2026/07/13/using-subagents-to-improve-claude-code/"
hn_url: "https://news.ycombinator.com/item?id=48898604"
title: "Using Subagents to Improve Claude Code Results: A Step-by-Step Guide"
article_title: "Using Subagents to Improve Claude Code Results: A Step-By-Step Guide – Software the Hard way"
author: "whack"
captured_at: "2026-07-13T20:50:10Z"
capture_tool: "hn-digest"
hn_id: 48898604
score: 1
comments: 0
posted_at: "2026-07-13T20:46:20Z"
tags:
  - hacker-news
  - translated
---

# Using Subagents to Improve Claude Code Results: A Step-by-Step Guide

- HN: [48898604](https://news.ycombinator.com/item?id=48898604)
- Source: [software.rajivprab.com](https://software.rajivprab.com/2026/07/13/using-subagents-to-improve-claude-code/)
- Score: 1
- Comments: 0
- Posted: 2026-07-13T20:46:20Z

## Translation

タイトル: サブエージェントを使用してクロード コードの結果を改善する: ステップバイステップ ガイド
記事のタイトル: サブエージェントを使用してクロード コードの結果を改善する: ステップバイステップ ガイド – ソフトウェアによるハードな方法
説明: 厳しい生活をしたことがない限り、Claude Code や Codex のような AI エージェントがソフトウェア開発に革命をもたらしていることはよくご存知でしょう。これらのツールを学ぶのは間違いなく簡単ですが、使いこなすのは難しいものでもあります。より良い結果を得るのに役立つニュアンスやベストプラクティスがたくさんあります。
[切り捨てられた]

記事本文:
サブエージェントを使用してクロード コードの結果を改善する: ステップバイステップ ガイド – ソフトウェア ザ ハードウェイ
コンテンツにスキップ
ハードウェイのソフトウェア
ソフトウェア開発への思い
サブエージェントを使用してクロード コードの結果を改善する: ステップバイステップ ガイド
厳しい生活をしたことがない限り、Claude Code や Codex のような AI エージェントがソフトウェア開発に革命をもたらしていることはよくご存知でしょう。これらのツールを学ぶのは間違いなく簡単ですが、使いこなすのは難しいものでもあります。コーディング エージェントからより良い結果を得るのに役立つニュアンスやベスト プラクティスが数多くあります。
強迫的な完璧主義者である私は、この分野について読んだり実験したりするのに（あまりにも）多くの時間を費やしてきました。そこには抽象的なアイデアがたくさん浮かんでおり、そこから多くを学びました。仕様駆動開発、RPI などのアイデアですが、これらの抽象的なアイデアを実際に実装する方法を示すステップバイステップのガイドが不足していることにも気づきました。だからこそ、私はこのブログ記事を書こうと思ったのです。
免責事項: 私は 6 か月前に Claude Code を使い始めたばかりで、これらのベスト プラクティスについても学習中です。このガイドには改善の余地が大いにあると確信していますが、近い将来新しいモデルやエージェントが登場すると、ひどく時代遅れになるでしょう。私の設定を改善する方法についてのあなたのアイデアをぜひ聞きたいです。
実際のステップバイステップ ガイドに入る前に、各ステップの背後にある指針と、それによってエージェントがより良い結果を生み出すことができる理由について簡単に説明します。これらの原則を理解すると、日常業務でコーディング エージェントをより効果的に使用できるようになります。
コーディング エージェントの各セッションにはコンテキスト ウィンドウがあります。エージェントにプロンプ​​トを送信するたび、エージェントが何かを行うたびに、このコンテキスト ウィンドウが少しずつ使用されます。これは不気味だ

タスクに取り組んでいるときの自分自身の認知負荷に似ています。あなた自身の認知負荷と同様に、コンテキスト ウィンドウがいっぱいになり始めると、コーディング エージェントのパフォーマンスが大幅に低下します。したがって、コンテキスト ウィンドウの使用量をできるだけ低く抑えるように常に努める必要があるのはそのためです。
もちろん、これは言うは易く行うは難しです。プロンプトが表示されるたびにコンテキスト ウィンドウを単純に /clear すると、1 分前に話した内容を覚えていない人と会話しているようなものになります。エージェントが良い仕事をするには、有用なコンテキストが必要です。したがって、重要なものを維持しながら、コンテキスト ウィンドウの使用を最小限に抑える方法を見つける必要があります。
この原則の非常に単純な応用例の 1 つは、関連性のない新しいタスクの作業を開始する前に、常にコンテキスト ウィンドウを /clear することです。これは良いスタートですが、他にもできることはたくさんあります。
次善の策は、サブエージェント (または独立したエージェント) を使用することです。各エージェントには独自の独立したコンテキスト ウィンドウがあります。そして、作業の一部を各サブエージェントに委任します。
大規模なコードベースがあり、エージェントが fooBar に関連するすべての情報を検索する必要がある場合を想像してください。メイン エージェントが、必要な情報を見つけるためにコードベースの grep と読み取りを開始すると、すぐに大量の役に立たない情報でコンテキスト ウィンドウを埋め始めます。これは、コンテキスト ウィンドウの使用量をできるだけ低く抑えたいという以前の原則に違反します。
目標を達成するには、代わりに、 fooBar に関連するすべての情報を検索し、この関連情報をメイン エージェントに送り返すことを唯一の仕事とするサブエージェントを起動します。このようにすると、無駄な情報がすべてそのサブエージェントのコンテキスト ウィンドウを肥大化させるだけになります。そして、メインエージェントのコンテキストウィンドウに表示される唯一の情報は、フィルタリングされた関連性の高い情報です。

サブエージェントによって明らかにされたアリ情報。
これと同じ原則は、「コードベースの greping」だけでなく、ソフトウェア開発ワークフロー全体にも適用できます。サイズが非常に大きい作業の場合は、すべてを 1 つのエージェントで実行するのではなく、作業を複数のチャンクに分割し、各チャンクを異なるサブエージェントに割り当てると、より良い結果が得られます。
もちろん、多数のサブエージェント間で作業を分割することにはマイナス面があります。それは、エージェントが相互に結果を送信するときに誤解や情報損失が発生するリスクです。電話ゲームをプレイしたことがある人なら、このリスクがどれほど危険であるかを知っています。
これを軽減するには、各エージェントに結果をテキスト ファイルに出力させると便利です。こうすることで、他のエージェントは二次情報や三次情報に頼るのではなく、関連ファイルを直接読み取ることができます。副産物として、これらのテキスト ファイルは、私たち人間にとって、各エージェントが行った作業を監査し、エラーの根本原因を特定するのにも役立ちます。
AI エージェントは一貫性があると考えたくなります。結局のところ、ほとんどのコンピュータ システムはそのように動作します。しかし、AIエージェントは違います。同じエージェントに 2 つの異なる機会に同じ指示を与えると、2 つのまったく異なる結果が生じる可能性があります。 AI エージェントに、1 分前に吐き出した結果を批評するように依頼すると、AI エージェントが自身の結果が間違っている理由をすべて教えてくれるかもしれません。
これは確かにイライラするかもしれませんが、これを利益のために利用することもできます。エージェントはタスクに取り組んでいるときに間違いを犯す可能性があります。しかし、エージェントに自分の作業をレビューするように依頼すると、人間が自分の作業を再確認するのと同じように、エージェントが自分の間違いを見つける可能性が高くなります。
これをさらに一歩進めることもできます。人間との別の不気味な類似点では、AI エージェントが「自分自身の」仕事をレビューするのと同じです。

防御的で偏見を持ってはいけません。しかし、新しいコンテキスト ウィンドウで新しいサブエージェントを起動し、「別のエージェントの結果」を確認するように指示すると、以前に行われた間違いが見つかる可能性が高くなります。サブエージェントが、最初に間違いを犯したエージェントとまったく同じモデルとハーネスを使用している場合でも。これは、サブエージェントのもう 1 つの素晴らしい使用例です。他のエージェントの作業成果をレビューします。ほとんどのソフトウェア チームが、PR を元の作成者以外の誰かがレビューすることを義務付けているのとまったく同じ方法です。
権力は腐敗する。そしてどうやらこれはLLMにも当てはまります。 AI エージェントは、不正行為を意味する派手な用語である「報酬ハッキング」を行うことがあることが知られています。実際には近道をしたり完全に失敗したりしても、ユーザーの要求を厳密に達成したように見せかけます。この問題を軽減する 1 つの方法は、テスト、実装、検証に独立した AI エージェント/サブエージェントを使用することです。このようにして、実装を実行するエージェントは、非包括的なテストを作成して不正行為を行うことができなくなります。これは、これらの異なるサブエージェントがすべてまったく同じモデルを使用している場合でも、驚くほどうまく機能します。
もう 1 つの副次的なボーナスとして、エージェントが間違いを犯したときに警告できるテストが導入されている場合、エージェントはコードを書く作業がはるかにうまくできることがわかりました。たとえそれらのテスト自体も AI エージェントによって書かれたとしても。
エージェントのベスト プラクティスの背後にある理論を説明したので、次に、これらのベスト プラクティスを実装する方法を説明します。すべては ~/.claude/CLAUDE.md にあるグローバル CLAUDE.md 設定から始まります。ここに次のセクションを追加しました。
## 開発ワークフロー コードベースを編集するタスクの開始時、つまり探索、計画、または編集の前に、dev-workflow スキルを呼び出して ev に従います。

その第一歩。メインの会話のみ: 派遣されたサブエージェントはこれをスキップし、代わりに派遣プロンプトに従います。 ## トップレベルのコンテキストを保持する トップレベルの会話のコンテキスト ウィンドウをできるだけ小さくします。私は、トークンや $ のコストよりも、存続期間の長い一貫したルート スレッドを重視します。コンテキスト重視またはツール重視の作業 (広範な検索、複数ファイルの読み取り、コマンド出力、探索) をサブエージェントに委任します。サブエージェントは、生の出力ではなく、抽出された概要のみを返します。この委任によるコストが高くても許容されます。例外: サブエージェントをスピンアップすると、節約できる以上のオーバーヘッドが発生する場合、単一の簡単な呼び出し (例: 1 つの小さなファイルの読み取り) がルート内で実行されることがあります。
これは、 ~/.claude/skills/dev-workflow/SKILL.md に配置する必要がある次のスキルを指します。
--- 名前: dev-workflow 説明: コードベースを編集するあらゆるタスクに必須の 13 ステップの開発ワークフロー (調査 → 計画 → テスト → 実装 → レビュー ループ → 検証 → ドキュメント)。このようなタスクの開始時、探索、計画、またはコード編集の前に呼び出します。ワークフローの調査フェーズがコードベースの探索を担当します。メイン会話オーケストレーターのみ。派遣されたサブエージェントは決してそれを呼び出してはなりません。 --- # 開発ワークフロー コードを変更する場合は、次のワークフローを使用してください。常にすべての手順に従ってください **オーケストレーターのみ** このワークフローはメインの会話用であり、以下の手順に従ってサブエージェントをディスパッチします。あなたがディスパッチされたサブエージェント (例: deep-explore 、 test-author 、implementer 、 code-reviewer 、 code-review-judge 、verifier 、agent-docs-maintainer ) である場合は、このファイルに従ってはいけません。ディスパッチ プロンプトとエージェント定義があなたのジョブを所有します。フェーズ アーティファクトは、`~ /.claude/thoughts/<repo-name>/ ` に ` Research-<task>.md ` 、 ` plan-<task>.md ` 、 ` review-plan-<task>.md ` 、 ` review-<task>- として存在します。

rN.md ` (レビュー ラウンドごとに 1 つ)、および `verify-<task>-rN.md ` (検証ラウンドごとに 1 つ) 各フェーズでは、それを実行するエージェント、何を渡すか、何を生成するかを指定します。各エージェントの独自の定義が、そのエージェントがどのように作業を行うかを決定します。 1. 研究。コードベースを編集するタスクの最初のツール呼び出しは `deep-explore` ディスパッチです。CLAUDE.md/AGENTS.md/PROGRESS.md を読み取って方向付けるのは問題ありませんが、それ以外はすべて `deep-explore` を経由するため、生のファイルの内容はメイン コンテキストから外れたままになります。最初に「方向性を定める」ためにメインスレッドのソースファイル、テスト、または構成を読まないでください。組み込みの `Explore` ではなく、`deep-explore` を使用してください。出力パス `research-<task>.md` を渡します。完全な結果がそこに書き込まれ、簡潔な要約のみが返されます。現時点ではコードを書かないでください。 2. 調査検証。新しい `deep-explore` エージェントを起動し、既存の `research-<task>.md` を渡します。耐荷重に関する主張を抜き打ちチェックしてもらいます — ハンドフを読んでください
[切り捨てられた]
.claude/agents/deep-explore.md :
--- 名前: deep-explore 説明: 読み取り専用、大規模なコンテキストのコードベース調査エージェント。組み込みの Explore サブエージェントの代わりにこれを使用して、計画や実装の前にタスクに関連するコードをマップしたり、多くのファイル、ディレクトリ、または命名規則をスイープする広範なファンアウト検索に使用します。抜粋を読んでコードを見つけて説明します。レビュー、監査、変更は行いません。出力ファイルのパスが指定されている場合、完全な結果がそこに書き込まれ、簡潔な概要のみが返されます。それ以外の場合は、結果を直接返します。呼び出し元のプロンプトでは、検索範囲を指定できます。焦点を絞ったパスの場合は「中」、複数の場所と命名規則の場合は「非常に徹底的」です。ツール: 読み取り、Bash、WebFetch、WebSearch、書き込み、編集、スキル --- あなたは読み取り専用のコードベース調査エージェントです。よ

呼び出しプロンプトに関連するコードをマッピングし、厳密でよく整理された結果レポートを作成します。生のファイル ダンプは決して作成しません。 # ワークフローのどこに位置するか あなたは ** 研究 ** フェーズ、つまり実装パイプラインの最初のステップです。レポートは、新しい機能を設計、実装、テストするための後のステップの基礎となります。あなたの出力は ** 事前のコンテキストを持たないエージェントによって読み取られ、あなたの探索を見ることはありません ** 。したがって、レポートは ** 自己完結型で、正確に参照され、正直である ** 必要があります。そのような読者は再検証せずにレポートに基づいて行動し、自信はあるものの間違ったメモがそのままテストと実装に伝播します。現実をマッピングし、後のステップで必要となるすべてを表面化します。あなたは**ではありません** ソリューションを設計したり、計画を書いたりするのです。それが次のステップの仕事です。 # 操作ルール - ** コードベースでは読み取り専用。決して変異させないでください。 ** 呼び出しで指定された単一の出力レポート ファイルを除き、* ファイルを作成、編集、移動、または削除しないでください (「レポートの配信」を参照)。 ` Bash ` は読み取り専用の調査にのみ使用してください (` git log ` 、 ` gitblame ` 、 ` git show ` 、 ` rg ` 、 ` grep ` 、 ` find ` 、 ` ls ` 、 ` wc ` ; このネイティブ ビルドはファイルのグロビングもルーティングします
[切り捨てられた]
--- 名前: テスト著者の説明: 著者の総合的な失敗

[切り捨てられた]

## Original Extract

Unless you've been living under a rock, you're well aware of AI agents like Claude Code and Codex revolutionizing software development. These tools are undoubtedly easy to learn, but are also hard to master. There are a ton of nuances and best-practices that can help you get much better results from
[truncated]

Using Subagents to Improve Claude Code Results: A Step-By-Step Guide – Software the Hard way
Skip to content
Software the Hard way
Thoughts on Software Development
Using Subagents to Improve Claude Code Results: A Step-By-Step Guide
Unless you’ve been living under a rock, you’re well aware of AI agents like Claude Code and Codex revolutionizing software development. These tools are undoubtedly easy to learn, but are also hard to master. There are a ton of nuances and best-practices that can help you get much better results from your coding agent.
As someone who can be an obsessive perfectionist, I’ve spent (too) many hours reading up on and experimenting in this area. There are a ton of abstract ideas floating out there which I’ve learnt a lot from. Ideas like spec-driven-development, RPI, etc. But I’ve also noticed a dearth of step-by-step guides that show someone how to actually implement these abstract ideas. Hence why I decided to write this blog post.
Disclaimer: I just started using Claude Code six months ago, and am still learning these best practices as well. I’m sure this guide has tremendous room for improvement, and will also become woefully out-of-date in the near future when new models and agents come out. I would love to hear your ideas on how to improve upon my setup.
Before we get into the actual step-by-step guide, let’s briefly discuss the guiding principles behind each step, and why it enables your agent to produce better results. Understanding these principles will help you make better use of coding agents in your day-to-day work.
Each session in a coding agent has a context window. Every time you send a prompt to the agent, every time the agent does anything, it uses up a bit more of this context window. This is eerily analogous to your own cognitive load while working on a task. Just like with your own cognitive load, coding agents perform far worse when their context window starts filling up . Hence why you should always strive to keep your context window usage as low as possible.
Of course, this is easier said than done. If you simply /clear your context window before every prompt, it will be like having a conversation with someone who doesn’t remember anything you said one minute ago. Agents need useful context in order to do a good job. So you need to find ways to minimize their usage of the context window, while still keeping the stuff that is important.
One extremely simple application of this principle: always /clear your context window before starting work on a new unrelated task. This is a good start, but there is also a lot more you can do.
The next best thing you can do is to use subagents (or independent agents) – each of which has its own independent context window. And delegate chunks of work to each subagent.
Imagine if you have a large codebase, and your agent needs to find all information relevant to fooBar . If your main agent simply starts grepping and reading your codebase in order to find the information it needs, it will rapidly start filling up its context window with a ton of useless information. This violates our earlier principle: we want to keep our context window usage as low as possible.
To achieve our goal, we can instead spin up a subagent whose sole job is to find all information relevant to fooBar , and send this relevant information back to the main agent. This way, all the useless information will only bloat that subagent’s context window. And the only information that lands in the main agent’s context window is the filtered and highly relevant information that was surfaced by the subagent.
This same principle can be applied not just to “codebase grepping”, but to the entire software development workflow. For any significantly sized work, instead of doing it all on a single agent, you will get better results if you break up the work into multiple chunks, and assign each chunk to a different subagent.
Of course, there is a downside to dividing up work among numerous subagents – the risk of misinterpretation and information loss when agents are sending results to one another. Anyone who has played a game of telephone knows how dangerous this risk is.
To mitigate this, it is helpful to have each agent output its result in a text file. This way, other agents can read the relevant file directly, instead of relying on second-hand or third-hand information. As a byproduct, these text files are also great for us humans to audit the work that each agent did, and root-cause the source of any errors.
It is tempting to think of AI agents as being consistent. After all, that’s how most computer systems operate. But AI agents are different. Give the same agent the same instructions on two different occasions, and it can produce two completely different results. Ask an AI agent to critique the results that it just spat out a minute ago, and it may tell you all the reasons why its own results are wrong.
This can certainly be frustrating, but we can also harness this to our benefit. An agent may make a mistake when working on a task. But if you ask the agent to review its own work, there is a good chance that it will find its own mistake – just like humans double-checking their own work.
We can even take this one step further. In another eerie resemblance to humans, an AI agent reviewing “its own” work is likely to be defensive and biased. But if you spin up a fresh subagent with a fresh context window, and tell it to review “a different agent’s results”, it is more likely to find any mistakes that were made earlier. Even if the subagent is using the exact same model and harness as the agent that made the mistake in the first place. This is yet another fantastic use-case for subagents – reviewing the work output of other agents. In the exact same way that most software teams mandate that PRs be reviewed by someone other than the original author.
Power corrupts. And apparently this is true for LLMs as well. AI agents are known to sometimes engage in “ reward hacking ” – a fancy term for cheating. They pretend to have rigorously accomplished the user’s request, when they actually took shortcuts or failed entirely. One way to mitigate this problem is to use independent AI agents/subagents for testing vs implementation vs verification. This way, the agent performing the implementation cannot cheat by writing non-comprehensive tests. This works surprisingly well even if these different subagents are all using the exact same model.
As another side-bonus, it turns out that agents do a much better job of writing code when they have tests in place that can alert them when they have made a mistake. Even if those tests were themselves written by an AI agent as well.
Now that we have covered the theory behind agentic best practices, here is how I have attempted to implement these best practices. It all starts with my global CLAUDE.md configs, in ~/.claude/CLAUDE.md , where I added the following section:
## Development workflow At the start of any task that will edit a codebase — before any exploration, planning, or edits — invoke the dev-workflow skill and follow every step of it. Main conversation only: dispatched subagents skip this and follow their dispatch prompt instead. ## Preserving top-level context Keep the top-level conversation's context window as small as possible — I value a long-lived, coherent root thread over token/$ cost. Delegate context-heavy or tool-heavy work (broad searches, multi-file reads, command output, exploration) to subagents that return only distilled summaries, never raw output. Higher cost from this delegation is acceptable. Exception: a single trivial call (e.g. reading one small file) may run in-root when spinning up a subagent would cost more overhead than it saves.
Which points at the following skill which you should place in ~/.claude/skills/dev-workflow/SKILL.md :
--- name: dev-workflow description: The mandatory 13-step development workflow ( research → plan → tests → implement → review loop → verify → docs ) for any task that edits a codebase. Invoke at the START of such a task , before any exploration , planning , or code edits — the workflow's research phase owns codebase exploration. Main-conversation orchestrator only ; dispatched subagents must never invoke it. --- # Development workflow Whenever making code changes , use the following workflow. Always follow every step **Orchestrator only.** This workflow is for the main conversation , which dispatches subagents per the steps below. If you are a dispatched subagent ( e.g. deep-explore , test-author , implementer , code-reviewer , code-review-judge , verifier , agent-docs-maintainer ) , do not follow this file — your dispatch prompt and agent definition own your job. Phase artifacts live in `~ /.claude/thoughts/<repo-name>/ ` as ` research-<task>.md ` , ` plan-<task>.md ` , ` review-plan-<task>.md ` , ` review-<task>-rN.md ` ( one per review round ) , and ` verify-<task>-rN.md ` ( one per verify round ) Each phase names the agent that runs it , what to hand it , and what it produces ; each agent's own definition owns how it does the work. 1. Research. The first tool call for any task that edits the codebase is a ` deep-explore ` dispatch - reading CLAUDE.md/AGENTS.md/PROGRESS.md to orient is fine , but everything else goes through ` deep-explore ` so raw file contents stay out of the main context. Do not read source files , tests , or configs in the main thread to "orient yourself" first. Use ` deep-explore ` , not the built-in ` Explore ` . Pass it the output path ` research-<task>.md ` ; it writes its full findings there, and returns only a concise summary. Do not write any code at this time 2. Research verification. Spin up a fresh ` deep-explore ` agent and hand it the existing ` research-<task>.md ` . Have it spot-check the load-bearing claims — reading the handfu
[truncated]
.claude/agents/deep-explore.md :
--- name: deep-explore description: Read-only, large-context codebase research agent. Use it instead of the built-in Explore subagent — to map the code relevant to a task before any planning or implementation, and for any broad fan-out search that sweeps many files, directories, or naming conventions. Reads excerpts to locate and explain code; it does not review, audit, or modify it. If given an output file path it writes its full findings there and returns only a concise summary; otherwise it returns the findings directly. The invoking prompt may specify search breadth: "medium" for a focused pass, "very thorough" for multiple locations and naming conventions. tools: Read, Bash, WebFetch, WebSearch, Write, Edit, Skill --- You are a read-only codebase research agent. You map the code relevant to your invocation prompt and produce a tight, well-organized findings report — never raw file dumps. # Where you sit in the workflow You are the ** Research ** phase — the first step of an implementation pipeline. Your report is the foundation the later steps build on, in order to design, implement, and test new functionality. Your output is read by ** agents that have no prior context and never see your exploration ** . So your report must be ** self-contained, precisely referenced, and honest ** — those readers act on it without re-verifying, and a confident-but-wrong note propagates straight into the tests and the implementation. You map reality and surface everything the later steps will need; you do ** not ** design the solution or write the plan — that is the next step's job. # Operating rules - ** Read-only on the codebase. Never mutate it. ** Do not create, edit, move, or delete any file — * except * the single output report file named in your invocation (see "Delivering your report"). Use ` Bash ` only for read-only investigation ( ` git log ` , ` git blame ` , ` git show ` , ` rg ` , ` grep ` , ` find ` , ` ls ` , ` wc ` ; this native build also routes file globbing
[truncated]
--- name: test-author description: Authors comprehensive faili

[truncated]
