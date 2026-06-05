---
source: "https://claude.com/blog/running-an-ai-native-engineering-org"
hn_url: "https://news.ycombinator.com/item?id=48408846"
title: "Running an AI-native engineering org"
article_title: "Running an AI-native engineering org | Claude"
author: "kiyanwang"
captured_at: "2026-06-05T07:37:32Z"
capture_tool: "hn-digest"
hn_id: 48408846
score: 1
comments: 0
posted_at: "2026-06-05T06:42:11Z"
tags:
  - hacker-news
  - translated
---

# Running an AI-native engineering org

- HN: [48408846](https://news.ycombinator.com/item?id=48408846)
- Source: [claude.com](https://claude.com/blog/running-an-ai-native-engineering-org)
- Score: 1
- Comments: 0
- Posted: 2026-06-05T06:42:11Z

## Translation

タイトル: AI ネイティブのエンジニアリング組織の運営
記事のタイトル: AI ネイティブのエンジニアリング組織を運営する |クロード
説明: エージェント コーディングがデフォルトの作業方法になってから、Claude Code エンジニアリング チームのプロセスと構造はどのように変化したか。

記事本文:
AI ネイティブのエンジニアリング組織を運営する |クロード
クロード製品のご紹介 クロード
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
AI ネイティブのエンジニアリング組織を運営する
AI ネイティブのエンジニアリング組織を運営する
Code w/ Claude SF 2026 では、Claude Code のエンジニアリング ディレクターと Claude Cowork の Fiona Fung が、エージェント コーディングがデフォルトの作業方法になってからチームのプロセスと構造がどのように変化したかを説明しました。
共有 リンクをコピー https://claude.com/blog/running-an-ai-native-engineering-org
長年にわたり、エンジニアリング帯域幅はアプリケーション構築において高価な部分でした。ソフトウェアの計画と出荷に関するすべてのプロセス (最初はウォーターフォール、次にアジャイル) は、そのコストを中心に構築されていました。
私は 2000 年代初頭に Visual Studio でキャリアをスタートしました。当時、私たちは厳しい製造期限を設けてソフトウェアを CD-ROM に入れて出荷していました。ソフトウェアをオンラインで配布できるようになると、継続的にアップデートを配布するようになりました。私たちは今、働き方を再び変えています。今度はソフトウェアを書くのに必要な時間と人材についてです。
Claude Code チームでは、コードを書いたり、テストを書いたり、リファクタリングしたりしても、速度が低下することはほとんどありません。しかし、エージェント コーディングによって実際にコードを入力する必要性がなくなっても、ボトルネックは解消されませんでした。検証、コードレビュー、セキュリティが代わりに行われました。
今では多くのコードを非常に高速に生成できるようになりましたが、これにより新しい q も発生します。

質問: このコードは正しいですか?どのように維持されていますか?そして、他のエンジニアリング リーダーからよく受ける質問の 1 つは、「人間はどのようにコード レビューを行っているかをどのように把握しているのですか?」です。
静かに動作を停止したプロセス
私たちは皆、ギャップを埋めたり何かをより良く機能させるために、何らかの理由があってプロセスを導入します。しかし、そのギャップがなくなり、それらのプロセスが時代遅れになったとしても、自然に消えることはほとんどありません。 Claude Code チームがデフォルトの作業方法としてエージェント コーディングを使用し始めたとき、既存のプロセスの多くが機能しなくなってしまいました。私たちが書き直した規範とその理由を以下に示します。
計画: ロードマップをジャストインタイムに移行する
以前の標準では、コーディングに時間がかかるため、事前計画に多くの時間を費やしていました。私が初めて Claude Code チームに加わったとき、6 か月間のかなり優れたロードマップを作成しましたが、その後、Claude Code のせいで多くのことが変更され、3 か月目には期限切れになってしまいました。
エンジニアリングの速度とスループットは現在とは異なるため、スプリントを計画する方法も変わりました。私はこれをジャストインタイム (JIT) プランニングと呼んでいますが、これは JIT コンパイルに似ています。適切なタイミングで適切な量を実行するにはどうすればよいでしょうか。私たちの計画の儀式は、設計ドキュメントから PR またはプロトタイプでの議論へと移行しました。この分野の動きは速いため、製品レビューはあまり行いません。私たちの現在のプロセスは、プロトタイプを作成し、多くの社内ユーザーに参加してもらい、フィードバックに基づいて行動を開始することです。
コンテキストの収集: 著者ではなくクロードに質問してください
エンジニアがコードを作成するとき、ほとんどの質問に対する答えを得る最初のステップは、コードを書いた人を見つけることでした。さて、すべての PR はクロードによってサポートされているため、「誰がこの変更を行ったのか?」ではもう十分ではありません。私たちの新しい標準は、さらに深いレベルに進むことです。つまり、実際に何を知る必要があるのか​​?例: 誰が回帰を引き起こしたかを探していますか?専門家が

顧客の質問に答えますか？それとも決定の背景でしょうか？あなたはクロードにその質問をし、クロードがより多くのデータとコンテキストを使用して、その質問に直接答えることができるかどうかを検討します。
Claude Code チームでは、質問がどのようなものであっても、「自動化する方法はありますか?」と尋ねるプロセスも行っています。たとえば、クロードに顧客フィードバック チャネルを毎朝要約させるという作業は、私がコーヒーを飲みながら手動で行っていた儀式から、バックグラウンドで自動的に実行するものに変わりました。
私たちはコードレビューを頻繁に使用します。 Claude は、すべてのスタイルと lint、PR フィードバックのリクエスト、完全なコミットの前にバグを見つけて修正し、テストを追加します。私たちが依然として確実に人材を必要としているのは、専門知識です。
新しい規範は、重要な場合には人間によるレビューです。法的レビューの場合、私は常に弁護士のパートナーにリスク許容度に関与してもらいたいと考えています。信頼境界とセキュリティに敏感なコードについては、ドメインの専門家が必要です。プロダクトマネージャーやデザイナーも、プロダクトのセンスやセンスに関わる必要があります。
ただし、信頼と検証の適切なバランスはモデルが改善されるにつれて変化し続けるため、継続的に評価することが重要です。現在人間に必要なものは、次のモデルでは異なるものになるかもしれません。
クロードと AI は、チーム全体の役割を再構築しました。私たちの PM は今、たくさんのコードを書いていますが、それを見るのは楽しいです。クロードのおかげで、従来とは異なるプログラマーがより多くのエンジニアリングを行うことができるようになり、コンテンツやデザインなど、従来は技術面ではなかった仕事を引き受けるエンジニアがいます。
Claude Code エンジニアリング チームでは、2 つのプロファイルに重点を置いてインデックスを作成しました。 1 つは、製品センスを備えた創造的なビルダーです。つまり、問題を解決する製品を出荷することに深い好奇心と情熱を持っている夢想家です。もう 1 人は、システムに関する深い専門知識を持つエンジニアです。例えば、私が入社したとき

チームの皆さん、システムの背景を持つ専門家が不足していることに気づきました。Web 上でクロード コードを構築するときに、どこでもクロードを確実に実行できるようにするために、それが必要でした。
一方、私がインデックス付けするのは生のスループットです。モデルがそれを処理します。より重要な問題は、人間の専門知識が依然として必要な部分であり、私はそこに焦点を当てます。
新しい規範をどのように展開したか
これらの規範が変化するにつれて、いくつかの側面はチーム原則として義務付けられ、その他の側面は小さなサブチーム (ポッド) が独自に理解できるようになりました。クロード コードのコア チーム原則には、交渉の余地のない「必ずやるべきこと」が含まれています。
製品を絶え間なくドッグフードする: クロスファンクショナル パートナーを含むすべての Claude Code チーム メンバーが Claude Code (および Claude Cowork) を使用しています。私たちは、仕事をより速く、より効率的に進めるためにクロードに協力してもらう方法を常に考えています。
チームを可能な限りフラットに保ちます。 Claude Code に入社したとき、私はすべてのマネージャーがまず IC としてスタートし、出荷を通じてチーム内で効果的なエンジニアになる方法を学び、Anthropic でエンジニアであることがどのようなものかを実際に経験して理解してほしいと考えていました。私たちには、Claude Code と Claude Cowork に関するチーム全体のミッションが 1 つあります。マネージャーは、人々が作業のある場所に移動できるようにチームの機敏性を保ちながら、一連の作業をサポートします。
機能しなくなったプロセスは躊躇せずに強制終了します。最後に、私たちはなぜそのようなやり方で物事を進めるのかを執拗に問い続けます。何かが意味をなさなくなった場合、チーム メンバーには古いプロセスを質問して強制終了する明示的な許可が与えられます。
ただし、これらのいくつかのルール内では、各ポッドには多くの権限があります。クロードをトリアージにどのように使用するか、計画の儀式やスタンドアップをどのように実行するか、どのワークフローを最初に「クローディファイド」するかを調整する余地があります。
新しいプロセスが行き詰まっていることを知る方法
ここにあります

すべてのエンジニアリング リーダーは、変更を展開する際に今すぐ追跡を開始する必要があります。
オンボーディングの開始時間が短縮される: エンジニア、デザイナー、PM はどれくらいで効果を発揮できるようになるでしょうか?私たちのチームでは、これは 1 年前よりもはるかに速くなり、エンジニアは今では最初の 1 週間以内に実際のコードを出荷しています。
PR サイクル タイムが短縮される: これは、パイプラインの拡張に苦労している箇所を特定するのに役立つ可能性があるため、詳しく調べるのは興味深いものです。非常に多くのコードを生成しているため、システムの構築や継続的インテグレーション (CI) が追いつくのが難しい場合があります。
クロード支援によるコミットの増加 : デフォルトでは、すべてのコミットはクロード支援されます。過去 4 か月間、クロードの支援を受けていないコミットを見たことがないと思います。
3 番目の項目では、スループットと成功を混同しないでください。スループットは 1 つの指標ですが、実際の指標は、解決しようとしているものを測定することです。適切に調整すると、スループットが向上し、問題をより迅速に解決できるようになります。
1 つだけ言いたいことがあるとすれば、最も騒々しいワークフローを選択してください。それは最もコストがかかるワークフロー、あなたが恐れているワークフロー、またはチームが楽しみにしていないワークフローである可能性があります。そして、それはまだその目的を果たしているのか、と尋ねてください。もしそうなら、それを自動化できますか？
私はかつて、会議室に大勢の人が集まって、高額な毎週のレビューを行うチームに所属していました。状況報告をするとき以外は、全員がラップトップを使っていることに気づきました。彼らは顔を上げてステータスを言い、ラップトップに戻りました。私は 1 つの単純な質問をしました。「なぜまたこの会議を行うのですか? 時間の無駄遣いのように思えます。」そしてたったその 1 つの質問で、誰もがそれが必要ではないことに気づきました。それでキャンセルしました。
そこで、自問してみてください。エンジニアリング ワークフローの中で、自動化や自動化を検討できる部分は何ですか。

完全に落ちてる？
前へ 前へ 0 / 5 次へ 次へ
クロードとともに構築するチーム向けの製品ニュースとベスト プラクティスをさらに詳しくご覧ください。
クロード コードの構築から得た教訓: スキルの使い方
クロード コード クロード コード構築からの教訓: スキルの使い方 クロード コード構築からの教訓: スキルの使い方 クロード コード構築からの教訓: スキルの使い方 クロード コード構築からの教訓: スキルの使い方 2026 年 6 月 2 日 あらゆるタスクに対応するハーネス: クロード コードの動的なワークフロー
Claude Code すべてのタスクのハーネス: Claude Code の動的ワークフロー すべてのタスクのハーネス: Claude Code の動的ワークフロー すべてのタスクのハーネス: Claude Code の動的ワークフロー すべてのタスクのハーネス: Claude Code の動的ワークフロー 2026 年 5 月 27 日 CodeRabbit がどのように Claude を使用してエージェント オーケストレーション システムを構築したか
クロード コード CodeRabbit がクロードを使用してエージェント オーケストレーション システムを構築する方法 CodeRabbit がクロードを使用してエージェント オーケストレーション システムを構築する方法 CodeRabbit がクロードを使用してエージェント オーケストレーション システムを構築する方法 CodeRabbit がクロードを使用してエージェント オーケストレーション システムを構築する方法 2026 年 5 月 20 日 クロード コードの使用: HTML の非合理的な効果
クロード コード クロード コードの使用: HTML の不合理な有効性 クロード コードの使用: HTML の不当な有効性 クロード コードの使用: HTML の不当な有効性 クロード コードの使用: HTML の不当な有効性 クロードを使用した組織の運用方法を変革する
製品のアップデート、ハウツー、コミュニティのスポットライトなど。毎月あなたの受信箱に配信されます。
購読 購読 毎月の開発者ニュースレターを受け取りたい場合は、電子メール アドレスを入力してください。いつでも購読を解除できます。
ホームページ ホームページ 次へ 次へ ありがとうございます!あなたの提出物は受理されました！おっと！送信中に問題が発生しました

フォームを整えます。書き込みボタン テキスト ボタン テキスト 学習ボタン テキスト ボタン テキスト コード ボタン テキスト ボタン テキスト 書き込み 聴衆向けにユニークな声を開発するのを手伝ってください こんにちはクロード!聴衆に合わせたユニークな声を開発するのを手伝ってくれませんか?私からさらに詳しい情報が必要な場合は、すぐに 1 ～ 2 つの重要な質問をしてください。より良い仕事をするために役立つドキュメントをアップロードした方がよいと思われる場合は、お知らせください。このタスクをより適切に達成するのに役立つ場合は、Google ドライブ、ウェブ検索など、アクセスできるツールを使用できます。分析ツールは使用しないでください。返答はフレンドリーで簡潔かつ会話的なものにしてください。
できるだけ早くタスクを実行してください。意味があればアーティファクトは素晴らしいものになります。アーティファクトを使用する場合は、どの種類のアーティファクト (インタラクティブ、ビジュアル、チェックリストなど) がこの特定のタスクに最も役立つかを検討してください。ご協力いただきありがとうございます！
文章のスタイルを改善する こんにちはクロード!私の書き方を改善してもらえませんか？私からさらに詳しい情報が必要な場合は、すぐに 1 ～ 2 つの重要な質問をしてください。より良い仕事をするために役立つドキュメントをアップロードした方がよいと思われる場合は、お知らせください。このタスクをより適切に達成するのに役立つ場合は、Google ドライブ、ウェブ検索など、アクセスできるツールを使用できます。分析ツールは使用しないでください。保管してください

[切り捨てられた]

## Original Extract

How the Claude Code engineering team’s processes and structure changed once agentic coding became the default way of working.

Running an AI-native engineering org | Claude
Meet Claude Products Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Running an AI-native engineering org
Running an AI-native engineering org
At Code w/ Claude SF 2026, Director of Engineering for Claude Code and Claude Cowork Fiona Fung walked through how the team’s processes and structure changed once agentic coding became the default way of working.
Share Copy link https://claude.com/blog/running-an-ai-native-engineering-org
For years, engineering bandwidth was the expensive part of building applications. Every process we used to have around software planning and shipping, first waterfall and then agile, was built around that cost.
I started my career in the early 2000s working on Visual Studio. In those days we shipped software on CD-ROMs with hard manufacturing deadlines. Once we could distribute software online, we began increasing to shipping updates continuously. Now we’re changing the way we work again, this time around the time and people it takes to write software.
On the Claude Code team, writing code, writing tests, and refactoring rarely slows us down anymore. But the bottlenecks didn’t go away when agentic coding took away the actual need to type code. Verification, code review, and security took their place.
We can all generate a lot of code really fast now, but this also brings up new questions: Is this code correct? How is it maintained? And one of the top questions I get from fellow engineering leaders: “How are humans keeping up with how you’re doing code reviews?”
The processes that quietly stopped working
We all put processes in place for a reason, to close a gap or make something work better. But when that gap no longer exists and those processes become obsolete, they rarely go away on their own. When the Claude Code team began using agentic coding as our default way of working, a lot of our existing processes stopped working. Here are the norms we rewrote, and why.
Planning: shift roadmaps to just in time
The old norm was to spend a lot more time pre-planning because coding time was expensive. When I first joined the Claude Code team, we wrote a pretty good six month roadmap, and then because of Claude Code, so many things changed that it was out of date by month three.
Engineering speed and throughput is different now, so the way we plan sprints has changed. I call it just-in-time (JIT) planning, almost like JIT compiling: how do you do just the right amount at the right time? Our planning ritual shifted away from design docs toward discussions in PRs or prototypes. The space moves fast so we don’t do a lot of product reviews. Our process now is let's prototype, get a lot of internal users on it, and start acting on their feedback.
Context gathering: ask Claude, not the author
When engineers wrote code, the first step to getting an answer to most questions was to find the person who wrote the code. Now, since all our PRs are assisted by Claude, "Who made this change?" is no longer sufficient. Our new norm is to go a level deeper: what do you actually need to know? For instance: Are you looking for who caused a regression? An expert to answer a customer question? Or context on a decision? You ask Claude that question, and consider whether Claude can answer it directly, also with more data and context.
On the Claude Code team, no matter what that question is, our process is to also ask “Is there a way to automate it?” For example, having Claude summarize customer feedback channels every morning went from a ritual I did manually with my coffee to something I just have running automatically in the background.
We use Code Review heavily. Claude handles all the style and linting, PR feedback requests, catching bugs and fixing them before a full commit, and adding tests. Where we still definitely want a human is expertise.
The new norm is human review where it matters: for legal review, I always want my legal partner involved in risk tolerance. For trust boundaries and security-sensitive code, I want the domain experts. Product managers and designers also need to be involved with product sense and taste.
It’s important to continually evaluate, though, because the right balance of trust vs. verify will keep changing as the models improve. What you need humans for today might look different with the next model.
Claude and AI have reshaped roles across the team. Our PMs code a lot now, which is fun to see. With Claude, you have nontraditional coders now being able to do more engineering, and you have engineers who take on things like content and design, work that were traditionally not on the technical side.
On the Claude Code engineering team, I’ve indexed heavily on two profiles. One is creative builders with product sense: the dreamers who are deeply curious and passionate about shipping products that solve problems. The other one is engineers with deep systems expertise. For example, when I joined the team, I noticed we were missing experts with systems backgrounds and we needed that when building Claude Code on the Web , to ensure we can run Claude everywhere.
What I index on less, on the other hand, is raw throughput; the models handle that. The more important question is where you still need human expertise, and that’s where I’d focus.
How we rolled out our new norms
As these norms changed, some aspects were mandated as team principles and others we let small sub-teams (pods) figure out on their own. There is a set of the Claude Code core team principles that are non-negotiable “must dos”:
Relentlessly dogfood your product: Every Claude Code team member, including cross-functional partners, uses Claude Code (and also Claude Cowork). We’re always thinking of ways to get Claude to help us do our work faster, and more efficiently.
Keep the team flat as possible. When I joined Claude Code I wanted every manager to start out as an IC first, learn how to be an effective engineer on the team by shipping, and really live through and understand what it’s like to be an engineer at Anthropic. We have one overall team mission on Claude Code and Claude Cowork. Managers support pods of work while keeping the team agile so people can move to where the work is.
Don’t hesitate to kill processes that no longer work: Finally, we relentlessly question why we do things the way we do. When something doesn’t make sense anymore, team members have explicit permission to question and kill old processes.
Within these few rules, though, each pod has a lot of agency. They have room to adapt how they use Claude to do triage, how they run any planning rituals or standups, and which workflows get “Claudified” first.
How to know your new processes are sticking
Here are three numbers every engineering leader should start tracking now as they roll out changes.
Onboarding ramp time goes down: How soon can an engineer, a designer, or a PM start being effective? On our team this is much faster than a year ago, and engineers ship real code now within their first week.
PR cycle time goes down: This one's interesting to dig into because it might help you identify where your pipeline is struggling to scale. As we’re generating so much more code, sometimes build systems and continuous integration (CI) may struggle to keep up.
Claude-assisted commits going up : For us, by default, every commit is Claude-assisted. I don't think I've seen a non-Claude-assisted commit in the last four months.
On the third bullet, don't confuse throughput with success. Throughput is one metric, but the real metric is measuring the thing you're trying to solve. With the right alignment, throughput can help you solve problems faster.
If I were to leave you with one thing: pick your noisiest workflow. That could be your most expensive workflow, the one you might be dreading, or that your team doesn't look forward to. And ask: is it still serving its purpose? If so, can you automate it?
I was once on a team that had an expensive weekly review, with a large number of people in a meeting room. I noticed everybody was on their laptops except when it was their time to give a status report. They would pop their head up, say the status, and go back down to their laptops. I asked one simple question: “Why are we having this meeting again? It seems like an expensive use of our time.” And just that one question made everyone realize it wasn’t needed. So we canceled it.
So, ask yourself: what's one piece of your engineering workflow that you might consider automating or even dropping altogether?
Prev Prev 0 / 5 Next Next eBook
Explore more product news and best practices for teams building with Claude.
Lessons from building Claude Code: How we use skills
Claude Code Lessons from building Claude Code: How we use skills Lessons from building Claude Code: How we use skills Lessons from building Claude Code: How we use skills Lessons from building Claude Code: How we use skills Jun 2, 2026 A harness for every task: dynamic workflows in Claude Code
Claude Code A harness for every task: dynamic workflows in Claude Code A harness for every task: dynamic workflows in Claude Code A harness for every task: dynamic workflows in Claude Code A harness for every task: dynamic workflows in Claude Code May 27, 2026 How CodeRabbit used Claude to build an agent orchestration system
Claude Code How CodeRabbit used Claude to build an agent orchestration system How CodeRabbit used Claude to build an agent orchestration system How CodeRabbit used Claude to build an agent orchestration system How CodeRabbit used Claude to build an agent orchestration system May 20, 2026 Using Claude Code: The unreasonable effectiveness of HTML
Claude Code Using Claude Code: The unreasonable effectiveness of HTML Using Claude Code: The unreasonable effectiveness of HTML Using Claude Code: The unreasonable effectiveness of HTML Using Claude Code: The unreasonable effectiveness of HTML Transform how your organization operates with Claude
Product updates, how-tos, community spotlights, and more. Delivered monthly to your inbox.
Subscribe Subscribe Please provide your email address if you'd like to receive our monthly developer newsletter. You can unsubscribe at any time.
Homepage Homepage Next Next Thank you! Your submission has been received! Oops! Something went wrong while submitting the form. Write Button Text Button Text Learn Button Text Button Text Code Button Text Button Text Write Help me develop a unique voice for an audience Hi Claude! Could you help me develop a unique voice for an audience? If you need more information from me, ask me 1-2 key questions right away. If you think I should upload any documents that would help you do a better job, let me know. You can use the tools you have access to— like Google Drive, web search, etc.—if they’ll help you better accomplish this task. Do not use analysis tool. Please keep your responses friendly, brief and conversational.
Please execute the task as soon as you can—an artifact would be great if it makes sense. If using an artifact, consider what kind of artifact (interactive, visual, checklist, etc.) might be most helpful for this specific task. Thanks for your help!
Improve my writing style Hi Claude! Could you improve my writing style? If you need more information from me, ask me 1-2 key questions right away. If you think I should upload any documents that would help you do a better job, let me know. You can use the tools you have access to— like Google Drive, web search, etc.—if they’ll help you better accomplish this task. Do not use analysis tool. Please keep

[truncated]
