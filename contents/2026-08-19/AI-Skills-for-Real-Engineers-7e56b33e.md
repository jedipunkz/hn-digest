---
source: "https://master.dev/blog/introducing-ai-skills-for-real-engineers/"
hn_url: "https://news.ycombinator.com/item?id=49363080"
title: "AI Skills for Real Engineers"
article_title: "Introducing AI Skills for Real Engineers – Master.dev Blog"
image: "https://master.dev/blog/wp-json/social-image-generator/v1/image/10700"
author: "ibobev"
captured_at: "2026-08-19T16:20:02Z"
capture_tool: "hn-digest"
hn_id: 49363080
score: 1
comments: 0
posted_at: "2026-08-19T15:42:54Z"
tags:
  - hacker-news
  - translated
---

# AI Skills for Real Engineers

- HN: [49363080](https://news.ycombinator.com/item?id=49363080)
- Source: [master.dev](https://master.dev/blog/introducing-ai-skills-for-real-engineers/)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T15:42:54Z

## Translation

タイトル: 本物のエンジニアのための AI スキル
記事タイトル: 本物のエンジニアのための AI スキルの紹介 – Master.dev ブログ
説明: マット・ポーコック

記事本文:
← Master.dev に戻る
コース
学ぶ
会員になる
ゲスト執筆
RSS
/ ブログ
AIスキル
本物のエンジニアのための AI スキルの紹介
AI は現在、ソフトウェア エンジニアリングにおいて否定できない力です。これらのツールは使い方が簡単で、何をすべきかを指示すると、それを実行します。しかし、これらのツールは使い始めるのは簡単ですが、うまく使いこなすのは難しい場合があります。
これは、Matt Pocock によって作成された AI スキル (AI ハーネスにインストールする種類) のスイートである「本物のエンジニアのための AI スキル」に関する投稿です。これらのスキルは、AI ツールをより効果的にするために存在します。使い方が簡単で、非常に影響力があることがわかりました。
これらのスキルのいくつかについて簡単に説明します。これらのスキルは、取り組んでいる内容に関係なく、ほぼすべてのソフトウェア開発者のワークフローにどのような価値を追加できるかを最も明確に示していると思います。
Claude Code を使用していると仮定します (他のハーネスについてはドキュメントを参照してください)。インストールするには:
クロードプラグインインストール mattpocock-skills コード言語: Bash ( bash )
それが次のようなもので失敗した場合…
「mattpocock-skills」がどのマーケットプレイスにも見つかりません
次に、Claude Code Marketplace のキャッシュをクリアしてみてください (はい、本当に)。このコマンドはまさにそれを行います:
クロード プラグイン マーケットプレイスの更新 コード言語: Bash ( bash )
スキルを使いこなすためのスキル
これらのスキルに関するドキュメントは非常に明確ですが、まだいくつかの質問があるかもしれません。信じられないかもしれませんが、他のスキルが何をするのかを理解するのに役立つスキルがあります。
チェックしてみましょう。ドキュメントを読んでも、/grilling スキルと /grill-me スキルの違いがよくわからないと想像してください。 /ask-matt スキルを起動して質問するだけです。
これらのスキルを実際のプロジェクトで実際に活用してみましょう。
これらのスキルの多くは、チケットの作成やドキュメントの作成などを行います。

ADR をご利用ください。これらやその他のことをセットアップするには、新しいプロジェクトで最初に行う必要があるのは、 /setup-matt-pocock-skills を実行することです。
ここで、問題を作成する場所、ADR を作成する場所などのスキルを構成します。これは単純なことですが、他のスキルをよりスムーズに実行するのに役立つ素晴らしい工夫です。
/grill-me を使用して明確な要件を作成する
コーディングに LLM を使用したことがある人なら誰でも、明確で詳細なプロンプトが不可欠であることを知っています。詳細が欠落していることは、AI を効果的に使用する上で非常に厄介です。プロンプトから省略したことを AI に任せると、結果に失望する可能性があります。 Claude Code のようなツールには計画モードがあり、LLM は一般に、「何か見逃していましたか?」などのことを喜んで受け入れます。プロンプトの最後に表示されますが、もっと良い方法があります。
/grill-me 仕様はこれらすべてを形式化し、次のレベルに引き上げます。
まず、/grill-me を実行して機能を説明してください。
プロンプトを分析し、驚くほど詳細な質問をいくつか考え出します。これらに答えると、おそらくフォローアップが行われるでしょう。
必要なものが揃うまで、この状態が続きます。
この時点で、セッションとコンテキストには、機能を実装するために必要なものがすべて含まれているはずです。クロードに「良さそうだね、作って」などと気軽に伝えてみましょう。
または、何らかの理由で、この機能を今すぐに実装する準備ができていない場合は、おそらく、テストおよびレビューする AI 生成の PR が他に 2 つまたは 3 つあるかもしれません。おそらく、他の 2 人のエージェントが今すぐに物事を構築していて、その待機時間を次のことを指定するために利用していただけかもしれません。その後、読み続けてください。
/to-spec を使用して作業を後で保存する
現在の会話とコンテキスト全体を取得して、後でそれを 1 つの問題に変換したい場合は、/ を使用できます。

スペック通りのスキル。呼び出すだけで、あとはスキルに任せましょう。いくつかのテストを追加し、適切なテストの境界について確認することも試みます。
/to-tickets を使用して作業を後で保存する
設計したばかりの機能が大きい場合はどうなるでしょうか?人間は、明確に定義された小規模なタスクで最も効果的に機能しますが、AI エージェントも例外ではありません。コンテキスト ウィンドウに不要なコンテンツが溢れないようにすることで、より良い結果が得られる可能性があります。
先ほどと同じ会話内で、/grill-me スキルを介して /to-tickets スキルを呼び出すことができます。これにより、その機能が複数の問題に分割されます。
依存関係に基づいて、必要に応じてチケットをブロックすることもできます。当然のことながら、提案された結果に必要な調整を加えることができます。
あなたが満足したら、そのように伝えれば、きちんとやってくれるでしょう。
きれいに埋められたボードが完成します。
私たちは皆、何かを学んだり理解したりするために LLM を使用したことがあります。このスキルスイートには、実際には次のレベルに引き上げるスキルがあります。 /teach スキルを起動して、何を学びたいかを伝えると、実際にレッスン全体がまとめられます。
準備が完了すると、レッスンがブラウザにポップアップ表示されます。
それが完了すると、このトピックの次のレッスンが続きます。
これらの AI スキルが私と同じように役立つことを願っています。これらは、アイデアを洗練し明確にして、LLM が最も効果的に実行できる明確で実用的な仕様に仕上げるのに非常に役立ちます。この投稿は、彼らができることすべてのほんの表面をなぞったに過ぎません。全体像についてはドキュメントを確認してください。
CLI エージェントを最初から作成してください。ツール呼び出し、エージェント ループ、評価などのエージェント開発の基礎を学びます。よりリスクの高い操作には人間参加型の承認を追加します。トークンの使用状況を監視し、高度な実装を行う

コンテキスト ウィンドウを管理するためのものです。 Master.dev サブスクリプションで 300 以上のコースにアクセスし、今すぐ 20% オフでご利用ください!
あなたのメールアドレスは公開されません。 * が付いているフィールドは必須です
次回コメントするときのために、このブラウザに名前、メールアドレス、ウェブサイトを保存してください。
フォローアップコメントを電子メールで通知します。
新しい投稿をメールで通知します。
スキルを使いこなすためのスキル
/grill-me を使用して明確な要件を作成する
/to-spec を使用して作業を後で保存する (今すぐまたは後で) 実装する
/to-tickets を使用して作業を後で保存する
私たちのコースはフロントエンドを超えて、フルスタック、DevOps、AI にまで及びます。
Master.dev は、 thanks.dev および Open Collective を通じてオープンソース プロジェクトに寄付しているほか、 The Last Mile 、 Annie Canons 、 Vets Who Code などの非営利団体にも寄付しています。

## Original Extract

Matt Pocock

← Back to Master.dev
Courses
Learn
Become a Member
Guest Writing
RSS
/ BLOG
AI Skills
Introducing AI Skills for Real Engineers
AI is an undeniable force in software engineering right now. These tools are trivial to use: tell it what to do, and it does it. But as easy as these tools are to start using, they can be tricky to use well .
This is a post about “AI Skills for Real Engineers” which is a suite of AI “skills” (the kind you install into your AI harness) created by Matt Pocock. These skills exist to help you make AI tooling more effective. I’ve found them to be straightforward to use and very impactful.
We’ll take a brief tour through a few of these skills, which I think most clearly shows the kind of value they can add to almost any software dev’s workflow, no matter what you’re working on.
I’ll assume you’re using Claude Code ( see the docs for other harnesses). To install:
claude plugins install mattpocock-skills Code language: Bash ( bash )
If that fails with something like…
"mattpocock-skills" not found in any marketplace
Then try clearing your Claude Code Marketplace cache (yes, really). This command will do just that:
claude plugins marketplace update Code language: Bash ( bash )
A Skill to Help You Use the Skills
The docs for these skills are pretty clear, but you might still have some questions. Believe it or not, there’s a skill that helps you understand what the other skills do.
Let’s check it out. Imagine you read the docs and you’re not quite sure what the difference between the /grilling and /grill-me skills is. You can just fire up the /ask-matt skill, and ask it.
Let’s actually put these skills to work in an actual project.
Many of these skills will do things like create tickets, or even documentation like ADRs. To get those, and other things set up, the first thing you should do in a new project is run /setup-matt-pocock-skills .
This is where we configure these skills: where to create issues, where to create ADRs, and so on. It’s a simple thing, but it’s a nice touch to help the other skills run more smoothly.
Producing Clear Requirements with /grill-me
Anyone who’s used LLMs for coding knows that clear, detailed prompts are essential. Missing details are anathema to effective AI use. If you leave AI to assume things you’ve left out of your prompt, you might be disappointed in the result. Tools like Claude Code do have a plan mode, and LLMs in general will happily accept things like “Did I miss anything?” at the end of your prompt, but there’s a better way.
The /grill-me spec formalizes all that and takes it to the next level.
To get started, just do / grill-me and describe your feature.
It’ll analyze your prompt and come up with some surprisingly detailed questions. As you answer those, you’ll likely be greeted with some follow-ups.
It’ll keep going like that until it has what it needs.
At this point, your session and context should have everything needed to implement your feature. You can absolutely feel free to tell Claude something like “looks good, build it.”
Or if, for whatever reason, you’re not ready for this feature to be implemented right this second—perhaps you have 2 or 3 other AI-generated PRs to test and review, perhaps you have two other agents building things right this second, and were just using that waiting time to spec the next thing—then read on.
Saving Work for Later with /to-spec
If you’d like to take the entirety of the current conversation and context and turn it into a single issue for later, you can use the /to-spec skill. Just call it up, and let the skill do the rest. It’ll even try to add some tests and check with you about the appropriate testing boundaries.
Saving Work for Later with /to-tickets
What if the feature you just designed is big ? Humans work best with small, well-defined tasks, and AI agents are no different. You’ll likely get better results if you avoid letting your context window get flooded with content you wouldn’t otherwise need.
Inside that same conversation you just had, via the /grill-me skill, you can call up the /to-tickets skill, which will break that feature into multiple issues for you.
It’ll even be smart enough to block tickets as needed, based on dependencies. Naturally, you can make any tweaks to the proposed result you’d like.
Once you’re happy, tell it so, and it’ll do its thing.
You’ll wind up with a nicely filled-out board.
We’ve all used LLMs to help us learn or understand something. This skill suite actually has a skill that takes it to the next level. Fire up the /teach skill, tell it what you’d like to learn about, and it’ll actually put an entire lesson together for you.
When it’s ready, the lesson will pop up in your browser.
When you’re done with that, it can keep churning on the next lesson in this topic.
I hope you find these AI skills as useful as I do. They can really help refine and clarify your ideas into clear, actionable specs that your LLM can execute most effectively. This post has barely scratched the surface of everything they can do. Check the docs for a fuller picture!
Create a CLI agent from scratch! Learn the foundations of agent development like tool calling, agent loops, and and evals. Add human-in-the-loop approvals for higher-stakes operations. Monitor token usage and implement advanced for managing the context window. Access 300+ courses with a Master.dev subscription and get 20% off today!
Your email address will not be published. Required fields are marked *
Save my name, email, and website in this browser for the next time I comment.
Notify me of follow-up comments by email.
Notify me of new posts by email.
A Skill to Help You Use the Skills
Producing Clear Requirements with /grill-me
Implementing (Now, or Later) Saving Work for Later with /to-spec
Saving Work for Later with /to-tickets
Our courses go beyond frontend into fullstack, devops, and AI.
Master.dev donates to open source projects through thanks.dev and Open Collective , as well as donates to non-profits like The Last Mile , Annie Canons , and Vets Who Code .
