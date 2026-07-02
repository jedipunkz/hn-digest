---
source: "https://www.stackbuilders.com/insights/ai-in-software-delivery-whats-working-whats-hard-and-what-were-still-learning/"
hn_url: "https://news.ycombinator.com/item?id=48762843"
title: "AI Increased Our Open PRs by 36%. That Wasn't the Whole Story"
article_title: "Stack Builders - AI in Software Delivery: What’s Working, What’s Hard, and What We’re Still Learning"
author: "StackBuilders"
captured_at: "2026-07-02T16:01:11Z"
capture_tool: "hn-digest"
hn_id: 48762843
score: 1
comments: 0
posted_at: "2026-07-02T15:09:35Z"
tags:
  - hacker-news
  - translated
---

# AI Increased Our Open PRs by 36%. That Wasn't the Whole Story

- HN: [48762843](https://news.ycombinator.com/item?id=48762843)
- Source: [www.stackbuilders.com](https://www.stackbuilders.com/insights/ai-in-software-delivery-whats-working-whats-hard-and-what-were-still-learning/)
- Score: 1
- Comments: 0
- Posted: 2026-07-02T15:09:35Z

## Translation

タイトル: AI によりオープン PR が 36% 増加しました。それがすべてではなかった
記事のタイトル: スタック ビルダー - ソフトウェア デリバリにおける AI: 機能していること、難しいこと、そしてまだ学習していること
説明: 仕様主導のワークフローやコード レビューから配信メトリクス、推定、エンジニアリング上の判断に至るまで、AI がソフトウェア配信をどのように変えているかを探ります。

記事本文:
スタック ビルダー - ソフトウェア配信における AI: 機能していること、難しいこと、そしてまだ学習していること
メニューアイコン
私たちについて
私たちについて
私たちがやっていること
私たちは誰なのか
私たちの働き方
私たちのチーム
エクアドル
スペイン
インターナショナル
当社のサービス
AIカスタムソフトウェア開発
レガシーの最新化
Webアプリケーション開発
モバイルアプリケーション開発
AI を中心とした配信
コンピュータビジョン
AI コードの最新化
AIエージェント
AIデータエンジン
カスタム AI モデルのトレーニング
AIデータエンジニアリング
AIエンタープライズ
安全なエンタープライズ ソフトウェア
スケーラビリティとDevOps
プロジェクトサポート
プロジェクト管理
デザイン
フラクショナルCTO
AI
エンジニアリングのリーダーシップ
プロジェクト管理
ソフトウェア開発
ソフトウェア配信における AI: 何が機能し、何が困難で、何がまだ学習中なのか
コピーアイコン
コピーチェックアイコン
スタックビルダーチーム
2026 年 7 月 1 日
2026 年 7 月 1 日
8 分で読めます
ブログを購読する
電子メール
購読する
ブログ投稿をもっと見る
Web 分析のクリーニング: Gemini AI を使用したボットの特定
ジセラ・オソリオ
6 分で読めます
OpenClaw を使用した AI エージェントの実践
スタックビルダーチーム
7 分で読めます
テキストがコードになるとき: LLM とデータベースの統合をプロンプト インジェクションから守る
スタックビルダーチーム
10 分で読めます
AI はソフトウェア配信を変えていますが、それはエンジニアリングによる強力な判断の必要性を取り除くことではありません。この要約では、Stack Builders の上級技術リーダーが AI 支援ワークフロー、コード レビュー、配信メトリクス、推定、実際のプロジェクトで AI を活用するために必要なプロセス規律について学んでいることを探ります。
AI 支援開発はもはや副次的な実験ではありません。 Stack Builders チーム全体で、シニア テクニカル リードは日々のデリバリーで AI を使用しています。コンポーネントの生成、仕様主導のワークフローのサポート、大規模なプル リクエストのレビュー、チーム メトリクスの分析、さらには問題のバッチを解決するための並行エージェントの調整などを行っています。
しかし、m

最近のシニア テック リードのディスカッションから得られた興味深い洞察は、単に「AI が私たちを速くする」ということではありませんでした。より有用な点は、より微妙なものでした。ソフトウェア配信の難しい部分が AI によって変化するということです。
コードを書く方が速いかもしれません。適切なリクエストを理解し、実際の進捗状況を測定し、正確さを検証し、ワークフローを調整し続けることが、新たなプレッシャーポイントになりつつあります。
この投稿では、当社の上級技術リーダーとの戦略的会話の独占的な要約と、チームがより自信を持ってソフトウェア配信のための AI を検討するために使用できる会話の質問を共有します。
1. 仕様主導の開発は有望だが、ワークフローの調整は依然として難しい
いくつかのチームが仕様駆動型開発を実験しています。あるチームは、軽量のオープンスペックのアプローチと、より多くのワークフローを規定するより構造化されたフレームワークを比較しました。初期のシグナルは肯定的でした。仕様は AI の出力をガイドし、実装がより制御されていると感じられるようにするのに役立ちます。
タスクがプロジェクト トラッカー内に存在し、テストが仕様フレームワーク内に存在し、AI が両方にわたって動作する場合、チームは新しい質問をし始めます。
トラッカーは依然として真実の情報源ですか?
タスクはリポジトリの近くで宣言する必要がありますか?
仕様、チケット、プロンプト、PR がバラバラにならないようにするにはどうすればよいでしょうか?
これは、新たな帽子をかぶったよく知られたソフトウェア品質の問題です。 AI によってもコンテキストの共有が不要になるわけではありません。システムが自信を持って間違った方向に加速する可能性があるため、古いコンテキストのコストが高くなります。
会話のきっかけ: AI 支援の作業において、真実の情報源はどこにあるべきですか: トラッカー、リポジトリ、仕様、または慎重に組み合わせたものなど?
2. AI は認知負荷を増大させながらコーディング時間を短縮できる
繰り返されるテーマの 1 つは精神的負荷でした。 AI はより多くのコード、より大きな PR を生成できます。

しかし、人間は何が変わったのか、なぜ変わったのか、それがドメインに適合するかどうかを理解する必要があります。
あるリードは、大規模なリクエストをより明確に説明するためにカスタム AI スキルを使用していると説明しました。 AI にコードの作成を依頼するだけでなく、チームは AI を使用して学習テクニックを研究し、概念、目標、非目標、トレードオフを細分化するのに役立つ再利用可能なスキルにパッケージ化しました。
これは有用なパターンです。AI は単なる生産ツールではなく、理解ツールとして使用されます。
以前のボトルネックは、多くの場合、「これを実装できるか?」ということでした。新たなボトルネックは、「これを十分早く理解して検証できるか?」ということかもしれません。
会話のきっかけ: AI が生成したコードまたは AI 支援のコードをレビューする際に、レビュー担当者が認知的負荷を軽減するにはどのようなワークフローが役立ちますか?
3. 最適な AI ワークフローは、汎用的なものではなく、プロジェクト固有のものである可能性があります
フロントエンドの例でこれが明確になりました。 AI は、スクリーンショットからデザイン システム コンポーネントや CMS ベースのセクションを生成するのには役立ちましたが、Figma スタイルのビジュアルを解釈するのには完璧ではありませんでした。チームはプロンプトを改良し、プロジェクト固有の指示を作成することで結果を改善しました。
別のリードは、これを「AI の出力の修正」から「プロセスの修正」への移行であると説明しました。チームは同じ間違いを手動で繰り返し修正する代わりに、プロンプト、ルール、記憶、例、制限、検証ループなどのハーネスを更新できます。
その考え方が重要です。同じ AI の間違いが 2 回発生した場合、それはコードの問題ではない可能性があります。ワークフロー設計の問題である可能性があります。
これは、Stack Builders の広範な AI の位置付けと一致しています。AI は、品質、セキュリティ、長期的な保守性を維持しながら価値を高める場合に適用されるべきです。
会話のきっかけ: チームレベルのルール、テスト、または再利用可能なスキルにする必要がある、私たちのチームで繰り返される AI の間違いにはどのようなものがありますか?
4. AIを測定する

インパクトは PR を数えるよりも難しい
あるチームは、導入前後の期間を比較することで AI の影響を測定し始めました。最初の測定期間中に、オープンされた PR が 36% 増加したと報告されていますが、チームはそれをすべてとして扱わないように注意しました。
その慎重さが重要です。 PR が増えるとスループットが増える可能性がありますが、レビュー負荷が増えたり、不完全な作業が増えたり、下流の調整が増えたりする可能性もあります。
グループでは、次のような代替指標について議論しました。
チケットがオープンされてからチケットが完了するまでの時間
PR のオープン、クローズ、マージ
QAとステージングまでのサイクルタイム
DORA や SPACE などのチームレベルの生産性フレームワーク
1 つの指標だけで信頼できるストーリーを伝えることができるかどうか
AI メトリクスは、開発者のアクティビティと配信結果を区別する必要があるという、有用な枠組みが明らかになりました。 PRとは活動のことです。検証され、展開され、保守可能な変更が結果として得られます。
会話のきっかけ: コード量そのものに報酬を与えることなく、AI 支援配信を最もよく捉えるメトリクスの組み合わせはどれですか?
5. ストーリーポイントには新たな意味が必要かもしれない
AI は従来の推定に挑戦します。以前は数日かかっていたタスクも、適切なモデル、コンテキスト、レビュー パスを使用すれば、午後には完了する可能性があります。
議論の結果、2つの可能性のある変更が浮上しました。
1 つのアプローチは、時間の代わりに複雑さを見積もることです。簡単なタスクは AI によって簡単なレビューで適切に処理される可能性がありますが、より難しいタスクには人間によるより慎重な検証が必要です。
もう 1 つのアプローチは、必要な判断レベルを推定することです。コード化に時間がかかるため、タスクは「大きく」ならない可能性があります。ドメインの深い知識を持つ人だけがアプローチが正しいことを検証できるため、この値は大きくなる可能性があります。
それは鋭い洞察です。 AI 支援配信では、希少なリソースが入力ではない可能性があります。判断かもしれません。
会話のきっかけ: 評価する必要があります

実装の労力、レビューのリスク、ドメインの判断、あるいはその 3 つすべてを考慮しますか?
6. 並行エージェントは一気に進行状況を解除できますが、人間によるトリアージが必要です
ある実験では、AI ワークフローを使用して、約 15 件の問題を複数のエージェントに並行して委任することが含まれていました。エージェントは問題を調査し、優先順位を付け、明らかな問題の多くを解決し、人間の入力が必要なケースを表面化しました。その結果、多くの PR は迅速に作成されましたが、人間の関心のほとんどは、実際に判断が必要ないくつかの問題に向けられました。
このパターンは重要だと感じます。AI は既知の作業全体に展開できますが、依然として人間が成功条件を定義し、結果をレビューし、曖昧さに対処する必要があります。
PR の監視、ビルド失敗の修正、成功条件が満たされるまでのループなどの他のツールやワークフローについても説明しました。これらは有望ですが、ガバナンスの問題も生じます。レビュー プロセスが配信の真のボトルネックになる前に、エージェントにどの程度の自主性を与えるべきでしょうか?
会話のきっかけ: どのタスクがエージェントの群れにとって安全で、どのタスクが意図的に人間主導のままにすべきでしょうか?
7. モデルの選択とアクセスが開発者のエクスペリエンスを形成する
チームはまた、ツールやモデル間で不均一なエクスペリエンスを報告しました。 TypeScript 用の Copilot と Haskell の小規模な変更で強力な結果が得られた人もいます。他の人は、優先ツールにアクセスできなくなったり、幻覚、遅いモデル、またはトークン制限に対処したりする際に摩擦を報告しました。
これは、AI 導入が単なる方法論の問題ではないことを思い出させます。それはインフラの問題でもあります。同じワークフローでも、モデルの品質、トークンの可用性、統合、レイテンシ、組織の制約によって、スムーズに感じられることもあれば、苦痛に感じられることもあります。
会話のきっかけ: チームは 1 つの AI ツールチェーンで標準化すべきか、それとも各プロジェクトが最適な AI ツールチェーンを使用できるように柔軟性を維持すべきか

モデルとワークフローは何ですか?
この会話が私たちに伝えること
議論の中で最も強力なテーマは、AI の導入が目新しさよりもむしろエンジニアリングの分野になっているということです。
チームは「AI はコードを書けるのか?」と尋ねているのではありません。彼らはより良い質問をしています:
AI をプロジェクト固有の基準に合わせて維持するにはどうすればよいでしょうか?
レビュー中の精神的負担を軽減するにはどうすればよいでしょうか?
実際の配信への影響をどのように測定するのでしょうか?
上級の判断が必要なのは何でしょうか?
どのワークフローをプロジェクト間で再利用できるようにする必要がありますか?
AI はどこに新たなリスクやボトルネックをもたらすのでしょうか?
それがこれからの本当の仕事です。 AI は出力を加速できますが、耐久性のあるソフトウェアは依然として明確なプロセス、慎重なレビュー、共有されたコンテキスト、経験豊富な判断に依存します。
実験から信頼性の高い AI 支援配信に移行しようとしているチームの場合は、ワークフローのプレッシャー ポイントを 1 つ選択し、それを小規模な実験に変えることから始めます。
大規模なリクエストを理解するための再利用可能なスキルを作成します。
繰り返されるミスに対してプロジェクト固有の AI ルールを追加します。
AI 導入前後のチケット サイクル タイムを比較します。
PR 量だけでなく、レビューの取り組みも追跡します。
リスクと判断を中心にストーリーのポイントを再構成します。
PM とテクニカル リードの立ち会いのもと、AI を活用した改良を試してください。
2 つの仕様主導のワークフロー間で A/B テストを実行します。
目標は、AI の利用を拡大することではありません。目標は、より読みやすく、測定しやすく、信頼できるものにすることです。
AI はソフトウェアの提供を変えていますが、信頼性の高いシステムを構築し、品質を可視化し、エンジニアリング上の判断を放棄することなく、より優れたツールを使用するという北極星は依然として馴染み深いものです。
あなたのチームが信頼性の高い配信プロセスに AI を組み込む方法を尋ねている場合、私たちは喜んでそれがどのようなものになるかを検討するお手伝いをさせていただきます。 ここをクリックして電話を予約してください。
Web 分析のクリーニング: Gemini AI を使用したボットの特定
ジセラ・オソリオ
6 分で読めます
P の AI エージェント

OpenClaw を使用したラクティス
スタックビルダーチーム
7 分で読めます
テキストがコードになるとき: LLM とデータベースの統合をプロンプト インジェクションから守る
スタックビルダーチーム
10 分で読めます
ブログを購読する
電子メール
購読する
フォローしてください
一緒に築きましょう
リンクトイン
GitHub アイコン
GitHub
X アイコン
×
フェイスブックのアイコン
フェイスブック
お問い合わせ
アメリカ
ニューヨーク州ロチェスター
(212) 686-5870
エクアドル
キト、EC
(+593) 2 401 6115
スペイン
マドリッド、SP
(+34) 911 983 415
メニュー
私たちがやっていること

## Original Extract

Explore how AI is changing software delivery, from specification-driven workflows and code review to delivery metrics, estimation, and engineering judgment.

Stack Builders - AI in Software Delivery: What’s Working, What’s Hard, and What We’re Still Learning
Menu Icon
About Us
About Us
What We Do
Who We Are
How We Work
Our Team
Ecuador
Spain
International
Our Services
AI Custom Software Dev
Legacy Modernization
Web Application Development
Mobile Application Development
AI Focused Delivery
Computer Vision
AI Code Modernization
AI Agents
AI Data Engine
Custom AI Model Training
AI Data Engineering
AI Enterprise
Secure Enterprise Software
Scalability & DevOps
Project Support
Project Management
Design
Fractional CTO
AI
Engineering Leadership
Project Management
Software Development
AI in Software Delivery: What’s Working, What’s Hard, and What We’re Still Learning
Copy Icon
Copy Check Icon
Stack Builders team
Jul. 1, 2026
Jul. 1, 2026
8 min read
Subscribe to blog
Email
Subscribe
Explore More Blog Posts
Cleaning Web Analytics: Identifying Bots with Gemini AI
Gisela Osorio
6 min read
AI Agents in Practice With OpenClaw
Stack Builders team
7 min read
When Text Becomes Code: Defending LLM–Database Integrations from Prompt Injection
Stack Builders team
10 min read
AI is changing software delivery, but not by removing the need for strong engineering judgment. This recap explores what Stack Builders’ senior tech leads are learning about AI-assisted workflows, code review, delivery metrics, estimation, and the process discipline needed to make AI useful in real projects.
AI-assisted development is no longer a side experiment. Across Stack Builders teams, senior technical leads are using AI in day-to-day delivery: generating components, supporting specification-driven workflows, reviewing larger pull requests, analyzing team metrics, and even coordinating parallel agents to resolve batches of issues.
But the most interesting insight from a recent Senior Tech Leads discussion was not simply “AI makes us faster.” The more useful takeaway was subtler: AI changes where the hard parts of software delivery live.
Writing code may be faster. Understanding the right request, measuring real progress, validating correctness, and keeping workflows aligned are becoming the new pressure points.
In this post, we’re sharing an exclusive recap of a strategic conversation with our senior tech leads, plus conversation questions your team can use to explore AI for software delivery with more confidence.
1. Specification-driven development is promising, but workflow alignment is still tricky
Several teams are experimenting with specification-driven development. One team compared a lighter-weight open-spec approach against a more structured framework that prescribes more of the workflow. The early signal was positive: specs help guide AI output and make implementation feel more controlled.
When tasks live in a project tracker, tests live in a spec framework, and AI operates across both, teams start asking new questions:
Is the tracker still the source of truth?
Should tasks be declared closer to the repository?
How do we prevent specs, tickets, prompts, and PRs from drifting apart?
This is a familiar software quality problem wearing a new hat. AI does not remove the need for shared context. It makes stale context more expensive because the system can confidently accelerate in the wrong direction.
Conversation starter: Where should the source of truth live for AI-assisted work: the tracker, the repo, the spec, or some carefully stitched combination?
2. AI can reduce coding time while increasing cognitive load
One recurring theme was mental load. AI can generate more code, larger PRs, and broader solutions, but humans still need to understand what changed, why it changed, and whether it fits the domain.
One lead described using custom AI skills to explain large requests more clearly. Instead of only asking AI to produce code, the team used AI to research learning techniques and package them into a reusable skill that helps break down concepts, goals, non-goals, and tradeoffs.
That is a useful pattern: AI as a comprehension tool, not just a production tool.
The old bottleneck was often “Can we implement this?” The new bottleneck may be “Can we understand and validate this fast enough?”
Conversation starter: What workflows could help reviewers reduce cognitive load when reviewing AI-generated or AI-assisted code?
3. The best AI workflows may be project-specific, not generic
A front-end example made this clear. AI was helpful for generating design system components and CMS-backed sections from screenshots, but it was not perfect at interpreting Figma-style visuals. The team improved results by refining prompts and creating project-specific instructions.
Another lead described this as moving from “fix the AI’s output” to “fix the process.” Instead of repeatedly correcting the same mistakes manually, teams can update the harness: prompts, rules, memories, examples, restrictions, and validation loops.
That mindset is important. If the same AI mistake happens twice, it may not be a code problem. It may be a workflow design problem.
This aligns with Stack Builders’ broader AI positioning: AI should be applied where it drives value while preserving quality, security, and long-term maintainability.
Conversation starter: What recurring AI mistakes do you see on our team that should become team-level rules, tests, or reusable skills?
4. Measuring AI impact is harder than counting PRs
One team started measuring the impact of AI by comparing a period before and after adoption. They saw a reported 36% increase in opened PRs during an initial measurement window, but the team was careful not to treat that as the whole story.
That caution matters. More PRs can mean more throughput, but they can also mean more review load, more incomplete work, or more downstream coordination.
The group discussed alternative metrics, including:
Time from ticket opened to ticket completed
PRs opened, closed, and merged
Cycle time through QA and staging
Team-level productivity frameworks such as DORA and SPACE
Whether one metric alone can tell a reliable story
A useful framing emerged: AI metrics should distinguish developer activity from delivery outcomes. A PR is activity. A validated, deployed, maintainable change is an outcome.
Conversation starter: What combination of metrics best captures AI-assisted delivery without rewarding code volume for its own sake?
5. Story points may need a new meaning
AI challenges traditional estimation. A task that used to take several days might now be completed in an afternoon with the right model, context, and review path.
The discussion surfaced two possible shifts.
One approach is to estimate complexity instead of time. Easy tasks may be handled well by AI with lighter review, while harder tasks require more careful human validation.
Another approach is to estimate the level of judgment required. A task may not be “large” because it takes long to code. It may be large because only someone with deep domain knowledge can verify that the approach is correct.
That is a sharp insight. In AI-assisted delivery, the scarce resource may not be typing. It may be judgment.
Conversation starter: Should estimation account for implementation effort, review risk, domain judgment, or all three?
6. Parallel agents can unlock bursts of progress, but they need human triage
One experiment involved using AI workflows to delegate around 15 issues to multiple agents in parallel. The agents investigated issues, triaged them, resolved many of the clear ones, and surfaced the cases that needed human input. The result: many PRs were created quickly, while most human attention went to the few issues that actually required judgment.
That pattern feels important: AI can fan out across known work, but humans still need to define success conditions, review results, and handle ambiguity.
Other tools and workflows were mentioned for monitoring PRs, fixing build failures, or looping until a success condition is met. These are promising, but they also raise a governance question: how much autonomy should we give agents before the review process becomes the true delivery bottleneck?
Conversation starter: Which tasks are safe for agent swarms, and which should remain deliberately human-led?
7. Model choice and access still shape the developer experience
Teams also reported uneven experiences across tools and models. Some found strong results with Copilot for TypeScript and smaller Haskell changes. Others reported friction when losing access to preferred tools, dealing with hallucinations, slower models, or token limits.
This is a reminder that AI adoption is not just a methodology question. It is also an infrastructure question. The same workflow can feel smooth or painful depending on model quality, token availability, integration, latency, and organizational constraints.
Conversation starter: Should teams standardize on one AI toolchain, or preserve flexibility so each project can use the best-fit model and workflow?
What this conversation tells us
The strongest theme from the discussion is that AI adoption is becoming less about novelty and more about engineering discipline.
The teams are not asking, “Can AI write code?” They are asking better questions:
How do we keep AI aligned with project-specific standards?
How do we reduce mental load during review?
How do we measure actual delivery impact?
What requires senior judgment?
Which workflows should be reusable across projects?
Where does AI introduce new risks or bottlenecks?
That is the real work now. AI can accelerate output, but durable software still depends on clear process, careful review, shared context, and experienced judgment.
For teams trying to move from experimentation to reliable AI-assisted delivery, start by choosing one workflow pressure point and turning it into a small experiment:
Create a reusable skill for understanding large requests.
Add project-specific AI rules for recurring mistakes.
Compare ticket cycle time before and after AI adoption.
Track review effort, not just PR volume.
Reframe story points around risk and judgment.
Try AI-assisted refinement with a PM and technical lead present.
Run an A/B test between two spec-driven workflows.
The goal is not to make AI usage bigger. The goal is to make it more legible, measurable, and trustworthy.
AI is changing software delivery, but the north star remains familiar : build reliable systems, keep quality visible, and use better tools without surrendering engineering judgment.
If your team is asking how to make AI part of a reliable delivery process, we’d be glad to help think through what that could look like: click here to book a call .
Cleaning Web Analytics: Identifying Bots with Gemini AI
Gisela Osorio
6 min read
AI Agents in Practice With OpenClaw
Stack Builders team
7 min read
When Text Becomes Code: Defending LLM–Database Integrations from Prompt Injection
Stack Builders team
10 min read
Subscribe to blog
Email
Subscribe
Follow Us
Let's build together
LinkedIn
GitHub Icon
GitHub
X Icon
X
Facebook Icon
Facebook
Contact
USA
Rochester, NY
(212) 686-5870
Ecuador
Quito, EC
(+593) 2 401 6115
Spain
Madrid, SP
(+34) 911 983 415
Menu
What We Do
