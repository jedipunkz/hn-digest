---
source: "https://newsletter.gtmengineering.ai/p/why-every-gtm-org-will-need-ai-skill"
hn_url: "https://news.ycombinator.com/item?id=48543543"
title: "Building an AI skill marketplace for GTM teams"
article_title: "Why every GTM org will need AI skill marketplace"
author: "alexjl1226"
captured_at: "2026-06-15T16:41:50Z"
capture_tool: "hn-digest"
hn_id: 48543543
score: 2
comments: 0
posted_at: "2026-06-15T16:20:20Z"
tags:
  - hacker-news
  - translated
---

# Building an AI skill marketplace for GTM teams

- HN: [48543543](https://news.ycombinator.com/item?id=48543543)
- Source: [newsletter.gtmengineering.ai](https://newsletter.gtmengineering.ai/p/why-every-gtm-org-will-need-ai-skill)
- Score: 2
- Comments: 0
- Posted: 2026-06-15T16:20:20Z

## Translation

タイトル: GTM チーム向けの AI スキル マーケットプレイスの構築
記事のタイトル: すべての GTM 組織が AI スキル マーケットプレイスを必要とする理由
説明: チームの AI 流暢さのギャップが拡大しており、それに対して何をすべきか。認識されていない 3 つの問題が、AI を使用する理由を静かに悪化させ、妨害しています。

記事本文:
すべての GTM 組織に AI スキル マーケットプレイスが必要な理由
すべての GTM 組織が AI スキル マーケットプレイスを必要とする理由
チームの AI 流暢さの差が拡大しています。それにどう対処するか
Alex Lindahl 2026 年 6 月 14 日 5 1 シェア これが転送された場合は、GTM Engineering ニュースレターへようこそ。こんにちは、私は Clay の最初の GTM エンジニアの 1 人である Alex です。ここでは、私が顧客のために構築しているもの、AI ネイティブ GTM の構築方法、および GTM エンジニアリング用のリソースを共有します。 7,000 人以上の GTM オペレーター、創設者、投資家に加わりましょう。
ほとんどの場合、どのモデルを使用すればよいかわかりません。モデル選択のサブタイトルには「複雑な推論」と書かれています。しかし、それは一体何を意味するのでしょうか？
私たちの多くは、ほとんどのタスクに対して最も強力なモデルをデフォルトで使用しています。そうしない動機はありません。まだ問題はありませんが、すぐに問題が発生するでしょう。
それが大規模に何を意味するのか考えてみましょう。 Anthropic の主力モデルは、最小モデルに比べてトークンあたりの価格が約 15 倍高くなります。通話記録を要約する担当者には旗艦は必要ありません。 CRM メモのフォーマットやフォローアップ メールの下書きを作成するワークフローも同様です。しかし、誰もそれを教えなかったし、それを強制するシステムもなく、誰もその違いを測定しませんでした。 AI ワークフローを一日中実行している 200 人の GTM 組織全体で、これを合計すると、まだ誰も所有していない実際の予算ラインに達します。
これは、GTM チームでの AI 導入に関する問題の 1 つにすぎません。
AI スキルの広がりももう 1 つの問題です。集中化の欠如ももう 1 つあり、さらに大きな問題は、個人の流暢さの範囲が広いため、AI の使用がいかに一貫性がなく、多様であるかということです。
私はしばらくこの問題について考えてきました。その後、Guy Podjarny によって設立された会社、Tessl を紹介されました。彼は Snyk (私が以前働いていたサイバーセキュリティのユニコーン企業) も設立しました。
Tessl は、リスクがあり、広大で目に見えないスキルを、管理された測定対象に変える管理レイヤーです。

ルシステム。彼らは、すべての AI ネイティブ チームが答える必要がある 3 つの質問があると主張しています。
セキュリティとガバナンス - あなたの環境で危険なスキルが実行された場合、それを知りますか?
標準化と再利用 - 重複した古いスキルがチームにどれほどのコストをもたらしているでしょうか?
継続的な最適化 - エージェントはスキルを持っていますが、実際にそれを活用していますか?
AI スキルについて解明すべきことは、私が思っていた以上にたくさんあります。
スキルはソフトウェアの新しい単位です。
これは、セキュリティやガバナンスなどに多くの影響を及ぼします。しかし、基本的なレベルでは、彼らは全員の GTM スタックのデフォルト レイヤーとなり、組織全体で AI を効果的に導入、拡張、使用するためのアプローチになると私が信じている製品を構築しています。
これは、より大きな問題の兆候の 1 つです。GTM チームは、基盤となるインフラストラクチャをまったく使用せずに AI スキルを構築しています。そして、私たちは以前にまさにこの映画を見たことがあります。
スキル マーケットプレイスは、組織が構築したスキルを見つけるためのコストや単一のソース以上のものを解決します。しかし、まずはコストの問題について考えてみましょう。
クラウドの初期には、エンジニアは誰でも好きなコンピューティングを起動できました。それは自由のように感じました。その後、法案が届き、混乱を一掃するために、規律全体である FinOps が発明されました。まだ完全には機能していません。Flexera の 2025 年のクラウド現状レポートでは、クラウド支出の 27% が無駄になっていることがわかり、この数字は 2019 年以来ほとんど変わっていません。6,750 億ドルの世界クラウド市場では、年間約 1,800 億ドルがアイドル状態のインスタンスや過剰プロビジョニングされたリソースに蒸発しています。
Flexera 自身のアナリストは、次はどうなるのかという線引きをしました。初期のクラウド利用が莫大なコストを生み出したのと同じように、AI への支出も同じ軌道をたどり、「AI のための FinOps」がすでにそれに対処するカテゴリーとして形成されつつあります。
ほとんどの GTM リーダーは、これをまだ結び付けていません。スキルではなくスキルです。

インフラストラクチャについては、組織内で AI にかかるコストを決定します。チームが Claude、ChatGPT、およびエージェントにフィードするプロンプト、指示、ワークフローによって、どのモデルが実行されるか、消費されるトークンの数、および出力が使用可能かどうかが決まります。スキルは、コスト、品質、セキュリティのすべてを決定する単位です。そして現在、私が目にするほぼすべての GTM 組織で、スキルは誰かの Slack DM 内に存在するマークダウン ファイルです。
AI 探求の荒野
コストの問題は目に見えています。その根底には 3 つの静かな問題が複雑に絡み合っています。
ソフトウェア エンジニアは、テストされていないコードは壊れたコードであることをずっと前に学びました。それはあなたがまだ知らないだけです。開発界ではすでにこれを AI に適用しています。Tessl はエージェントのスキルに対してタスクベースの評価を実行し、評価され最適化されたコンテキストにより、API を正しく使用するエージェントのパフォーマンスが最大 3.3 倍向上することがわかりました。ここで自問してください。GTM 組織の中で、プロスペクティング スキルの評価を実行したことがあるのは誰ですか?アカウント調査のプロンプトが表示されますか?誰でも？ GTM では、スキルの「テスト」とは、1 人がそのスキルを 2 回実行し、問題がなかったと思われることを意味します。私たちは評価されていないスキルを営業チーム全体に配布していますが、なぜアウトプットの品質が一貫していないのか疑問に思っています。
誰もが異なる成果を得ることができます。 10 人の担当者がそれぞれ「電話を受ける前にこのアカウントを調査してください」という独自のバージョンを作成すると、10 の異なる会議に 10 の異なる調査基準が適用されることになります。 GTM エンジニアリングの要点は、チームの床を個人の天井よりも高くするシステムを構築することです。バージョン管理されず、共有されていないスキルはその逆です。出力の品質は、誰が最良のプロンプトを書いたかに基づいて抽選になります。
私はこれを常に見ています。誰かが優れた ICP スコアリング スキルを構築しました。 3 週間後、最初のポッドを知らなかったため、2 ポッド上の誰かが同じもののより悪いバージョンをビルドしました。

一つは存在した。ソフトウェアでは、パッケージ マネージャー (npm、pip) を使用してこの問題を解決したため、日付解析ライブラリを最初から書き直す人はいません。 GTM には同等のものはありません。良い仕事は伝わりません。ひどいことに、ただ重複してしまうだけです。
あなたのチームには AI スキル マーケットプレイスが必要です
これが理論です。どの企業も最終的にパッケージ マネージャー、CI パイプライン、FinOps 機能を必要としたのと同じように、AI スキルのための社内マーケットプレイス (管理され、バージョン管理され、測定可能なレジストリ) が必要になるでしょう。
現代の GTM チームの中心には構造的な矛盾があります。チームの最高の担当者がやり方を知っていることと、平均的な担当者が実際に行うこととの間のギャップは非常に大きく、AI ツールの普及に伴いそのギャップはさらに広がっています。
暗黙知（実際の実践を通じて習得される直観的なスキル）は、獲得して拡大するのが最も難しいタイプです。優秀な人材が長年の顧客とのやり取りを通じて培ったベスト プラクティスと秘訣は、システムの内部ではなく、その個人の内部に生きています。イネーブルメントの課題は、その専門知識を体系化し、チーム全体で共有することです。
一方、AI により、GTM 作業の技術的表面積は劇的に拡大しました。営業チームの 81% が AI を実験中ですが、実際の ROI を生み出すためにパイロット プログラムを超えて AI を拡張できるのは 26% だけです。マッキンゼーの報告によると、営業担当者が実行するタスクの 60 ～ 70% は技術的に自動化可能ですが、可能性とパフォーマンスのギャップはオペレーターの問題です。
これは、トレーニング カタログではなく、担当者やオペレーターがそこから引き出して導入できる、実行可能なワークフローの生きたライブラリに対処する「社内スキル マーケットプレイス」に対する構造的なニーズです。
営業担当者の 92% がすでに何らかの形で AI を使用していますが、84% が主な利点は時間の節約またはプロセスの最適化であると述べています。それは

テーブルステークス。先頭に立っている担当者は、時間の節約を超えて真の競争上の優位性を獲得しました。このギャップは、ツールへのアクセスではなく、AI の流暢さによって決まります。これらを区別する最上位のスキルは、迅速なエンジニアリング、AI 出力の検証、データ リテラシーです。これらは、ほとんどの企業が無視しているものであることはほぼ間違いありません。
開発界は現在これを構築中です。 Tessl は本質的にエージェント スキルのための npm です。スキルをバージョン管理された依存関係としてインストールし、マニフェストで追跡し、出荷前に評価を実行し、セキュリティ上の問題をスキャンし、リポジトリ間でマークダウン ファイルをコピーする代わりにチーム全体に更新をプッシュします。彼らが使用する枠組みは正しいものです。スキルはソフトウェアであり、ソフトウェアにはライフサイクルが必要です。
GTM はこれを丸ごと盗むべきです。実行可能な最小バージョンは次のとおりです。
唯一の真実の情報源。チームのすべてのスキルが存在する 1 つのリポジトリ (または Notion データベース、または共有ディレクトリ。分野よりもツールは重要ではありません)。スキルがレジストリにない場合、そのスキルは存在しません。 Slack-DM 配布モデルを廃止します。
スキルごとに所有者。すべてのスキルには、基礎となるツール、モデル、またはプレイブックが変更された場合の更新を担当する指名されたメンテナーが割り当てられます。孤立したスキルにより、担当者は廃止されたモデル用に構築された 6 か月前のプロンプトを実行することになります。
すべてのスキルに組み込まれたモデルの推奨事項。これは最もコストがかからず、コストへの影響が最も大きい修正方法です。各スキルは、必要なモデル層を宣言する必要があります。つまり、要約と書式設定には高速かつ安価、製図には中間層、本格的な推論にはフラッグシップのみです。ほとんどの GTM スキルにはフラッグシップは必要ありません。スキル自体の中で明示的にそう言っておくと、トークンの予算に関するメモを誰かに送信するよりも優れています。
チーム全体に何かを出荷する前に軽量級の評価を行います。開始するために評価プラットフォームは必要ありません。 10リアルを取る

入力 — 10 個の実際のアカウント、10 個の実際のトランスクリプト — それらすべてに対してスキルを実行し、スキルの所有者に単純なルーブリックに照らして出力を採点させます。それを超えられない場合は、50 回の繰り返しに耐えられません。この 1 つの習慣だけで、私がこれまで見てきたほぼすべての GTM 組織よりも優れています。
たとえそれが粗雑であっても、バージョン管理を行います。ファイル名の v1、v2、v3 に勝るものはありません。スキルが改善されると、チームはその改善を自動的に採用します。保存したコピーを選択するのではありません。これは、メッセージングの変更が反映される方法でもあります。ピッチスキルを一度更新すると、チーム全体が次の四半期ではなく明日新しいプレーを実行します。
(コストと比較して) より良い成果を生み出すのは、コンテキストまたはモデルですか?
Tessl はそのプラットフォームを使用して、スキルのコンテキスト内で Claude Fable 5 と Opus 4.8 を比較することでこの質問に答えました。
これは、これらすべてがなぜ重要なのかを示す好例です。
評価のすべてのシナリオは、公開されたスキルに関連付けられた実際のエージェントのタスクであり、指示への従うこと (エージェントが指示されたことを、指示された方法で実行するか) とタスクの完了 (目標に到達するか) の 2 つの軸でスコア付けされます。全体的なスコアは、指示に従っていることを 4 として、タスクの完了を 3 として重み付けし、7 で割ります。各タスクはスキルの有無にかかわらず実行されるため、スキルによる上昇が直接目に見えます。タスクとスキルは task-evals-for-skills データセットで公開されているため、任意のシナリオを自分で検査できます。
このデザインは意図的なものです。タスクは公開されたスキルに基づいているため、モデルの天井を見つけることを目的としたフロンティアパズルではなく、実際の作業チームがスキルを作成することを反映しています。これが、両方のモデルでタスク完了率が高い理由であり、これらを区別するシグナルが指示に従うこと、つまりスキルが要求する特定の方法で作業を実行することである理由です。 — 「クロード・フェイブル 5 対オーパス 4.8: 神話の誇大宣伝」より

現実との出会い」
スキル マーケットプレイスがチームにどのように役立つか
レジストリはインフラです。この値は 3 か所に表示され、便利なことに、これを承認する必要がある 3 人にマッピングされます。
現在、新任の担当者は最初の 1 か月間、チームの優秀な人材がどのように仕事をしているかをリバース エンジニアリングすることに費やしています。通話をシャドーイングしたり、すぐにスクリーンショットを撮ってもらったり、ぎくしゃくした独自のツールキットをゆっくりと組み立てたりしています。スキル レジストリを使用すると、初日が変わります。新しい担当者には、チームのすべてのロードアウトが自動インストールされます。アカウント調査スキル、電話準備スキル、フォローアップ ライターなど、トップ パフォーマーが実行しているものとすべて同じバージョンです。立ち上げ時間は常に、上級担当者から新しい担当者に知識がどれだけ早く伝わるかによって左右されてきました。スキル レジストリは、数か月にわたる浸透をダウンロードに変換します。
どのチームにも、コールの準備が異常に優れている担当者が 1 人はいます。現時点では、そうなるとその利点は失われます。彼らのプロセスをスキルに変え、評価し、公開すれば、彼らの優位性がチームのフロアとなります。これは、ほとんどの GTM 組織がこれまでに行った中で最も ROI の高いナレッジ マネジメントですが、これを行っている人はほとんどいません。
スキルに組み込まれたモデルの推奨事項により、誰も取り締まることなく、デフォルトから高価になる問題が解決されます。また、法務部門は、担当者が製品で実行できないことを約束した後、フリーランスのプロンプトを 1,000 件監査する代わりに、アウトリーチ スキルの主張を 1 回承認します。登録用

[切り捨てられた]

## Original Extract

Your team's AI fluency gap is widening and what to do about it. 3 unrecognized problems are quietly compounding and sabotaging the reasons you're using AI.

Why every GTM org will need AI skill marketplace
Subscribe Sign in Why every GTM org will need AI skill marketplace
Your team's AI fluency gap is widening and what to do about it
Alex Lindahl Jun 14, 2026 5 1 Share If this was forwarded, welcome to the GTM Engineering newsletter . Hi, I’m Alex, one of the first GTM Engineers at Clay and here to share what I’m building for customers, how to build AI-native GTM, and resources for GTM Engineering. Join 7,000+ GTM operators, founders, and investors.
Most of the time - I have no idea what model to use. The subtitle on the model selection says ‘complex reasoning’. But what does that even mean?
Many of us default to the most powerful model for most tasks. There’s no incentive not to. It’s not problem yet, but it will be soon.
Think about what that means at scale. Anthropic’s flagship model costs roughly 15x more per token than its smallest one. A rep summarizing a call transcript doesn’t need the flagship. Neither does a workflow that formats CRM notes or drafts a follow-up email. But nobody told them that, there’s no system enforcing it, and nobody’s measuring the difference. Across a 200-person GTM org running AI workflows all day, that adds up to a real budget line nobody owns yet.
This is just one problem with AI adoption on GTM teams.
AI skill sprawl is another one. Lack of centralization is another, and the larger one is how inconsistent AI usage is and varied due to a wide spectrum of individual fluency.
I’ve been thinking about this problem for awhile. Then was introduced to Tessl , a company founded by Guy Podjarny . He also started Snyk (a cybersecurity unicorn where I used to work).
Tessl is a management layer that turns risky, sprawling, invisible skills into a governed, measurable system. They argue there’s 3 questions that every AI native team need to answer:
Security & governance - if a risky skill ran in your environment, would you even know?
Standardization & reuse - how much are duplicate, outdated skills costing your team?
Continuous optimization - your agents have the skills, but are they actually using them?
There’s a lot more to AI skills to unpack than I realized.
Skills are a new unit of software.
This has many implications for security, governance, and so on. But at a basic level, they’re building a product I believe will become a default layer of everyone’s GTM stack and approach to adopting, scaling, and using AI effectively across the org.
This is one symptom of a bigger problem: GTM teams are building AI skills with zero infrastructure underneath them. And we’ve seen this exact movie before.
Subscribe A skill marketplace will solve for more than just cost and a single source to find your organizations built skills. But let’s dive into the cost problem first.
In the early cloud days, every engineer could spin up whatever compute they wanted. It felt like freedom. Then the bills arrived, and an entire discipline — FinOps — was invented to clean up the mess. It still hasn’t fully worked: Flexera’s 2025 State of the Cloud Report found that 27% of cloud spend is wasted, a number that’s barely moved since 2019. At a $675B global cloud market, that’s roughly $180B a year evaporating into idle instances and over-provisioned resources .
Flexera’s own analysts drew the line to what’s next: just as early cloud usage produced unwieldy costs, AI spend is following the same arc — and “ FinOps for AI ” is already forming as a category to deal with it.
Most GTM leaders haven’t connected this yet: Skills, not infrastructure, decide what AI costs in your org. The prompts, instructions, and workflows your team feeds to Claude, ChatGPT, and your agents determine which model runs, how many tokens burn, and whether the output is even usable. Skills are the unit where cost, quality, and security all get decided. And right now, in almost every GTM org I see, skills are markdown files living in someone’s Slack DMs.
The wild west of AI exploration
The cost problem is the visible one. There are three quieter problems compounding underneath it.
Software engineers learned long ago that untested code is broken code — you just don’t know it yet. The dev world is already applying this to AI: Tessl runs task-based evaluations on agent skills and found that evaluated, optimized context produced up to 3.3x improvement in agents using APIs correctly. Now ask yourself: who in your GTM org has ever run an eval on a prospecting skill? On an account research prompt? Anyone? In GTM, “testing” a skill means one person ran it twice and it seemed fine. We’re shipping un-evaluated skills to entire sales teams and wondering why output quality is inconsistent.
Everyone gets different outputs. When ten reps each write their own version of “research this account before my call,” you get ten different research standards walking into ten different meetings. The whole point of GTM engineering is building systems that make the team’s floor higher than any individual’s ceiling. Unversioned, unshared skills do the opposite — your output quality becomes a lottery based on who happened to write the best prompt.
I see this constantly. Someone builds a great ICP-scoring skill. Three weeks later, someone two pods over builds a worse version of the same thing because they didn’t know the first one existed. In software, we solved this with package managers — npm, pip — so nobody rewrites a date-parsing library from scratch. GTM has no equivalent. Good work doesn’t travel; it just gets duplicated, badly.
Your team needs an AI skill marketplace
This is the thesis: every company is going to need an internal marketplace for AI skills — a governed, versioned, measurable registry — the same way every company eventually needed a package manager, a CI pipeline, and a FinOps function.
There’s a structural paradox at the center of modern GTM teams: the gap between what the best rep on the team knows how to do and what the average rep actually does is enormous — and it’s getting wider as AI tooling proliferates.
Tacit knowledge — intuitive skills learned through first-hand practice — is the hardest type to capture and scale. The best practices and tricks that high performers develop through years of customer interactions live inside those individuals, not inside any system. The enablement challenge is codifying that expertise and sharing it across the team.
Meanwhile, AI has made the technical surface area of GTM work dramatically larger. While 81% of sales teams are experimenting with AI, only 26% can scale it beyond pilot programs to generate real ROI. McKinsey reports that 60–70% of tasks performed by sales reps are technically automatable — but the gap between potential and performance is an operator problem.
This is the structural need for an “internal skills marketplace” that addresses: not a training catalog, but a living library of executable workflows that any rep or operator can pull from and deploy.
92% of sales professionals already use AI in some form , yet 84% say the main benefit is saving time or optimizing processes. That’s table stakes. The reps pulling ahead have moved past time savings into genuine competitive advantage. The gap comes down to AI fluency, not tool access. The top skills that separate them are prompt engineering, AI output verification, and data literacy. These are almost certainly the ones most companies ignore.
The dev world is building this right now. Tessl is essentially npm for agent skills: install skills as versioned dependencies, track them in a manifest, run evals before they ship, scan them for security issues, and push updates across the whole team instead of copying markdown files between repos. The framing they use is the right one — skills are software, and software needs a lifecycle.
GTM should steal this wholesale. Here’s the minimum viable version:
A single source of truth. One repo (or Notion database, or shared directory — the tool matters less than the discipline) where every team skill lives. If a skill isn’t in the registry, it doesn’t exist. Kill the Slack-DM distribution model.
An owner per skill. Every skill gets a named maintainer responsible for updates when the underlying tool, model, or playbook changes. Orphaned skills are how you end up with reps running a six-month-old prompt built for a model that’s been deprecated.
A model recommendation baked into every skill. This is the cheapest fix with the biggest cost impact. Each skill should declare which model tier it needs: fast-and-cheap for summarization and formatting, mid-tier for drafting, flagship only for genuinely hard reasoning. Most GTM skills don’t need the flagship — saying so explicitly, inside the skill itself, beats sending anyone a memo about token budgets.
Lightweight evals before anything ships team-wide. You don’t need an eval platform to start. Take 10 real inputs — 10 actual accounts, 10 real transcripts — run the skill against all of them, and have the skill’s owner grade the outputs against a simple rubric. If it can’t pass that, it’s not ready for 50 reps. This one habit alone puts you ahead of nearly every GTM org I’ve seen.
Versioning, even if it’s crude. v1, v2, v3 in the filename beats nothing. When a skill gets improved, the team adopts the improvement automatically — not whichever copy they happened to save. This is also how messaging changes ship: update the pitch skill once, and the whole team is running the new play tomorrow instead of next quarter.
What drives better outputs (relative to cost): context or model?
Tessl used it’s platform to answer this question by comparing Claude Fable 5 vs Opus 4.8 within the context of skills.
It’s a great example of why all of this matters.
Every scenario in the evaluation is a real agent task tied to a published skill, scored on two axes: instruction-following (does the agent do what it was told, in the way it was told) and task-completion (does it reach the goal). The overall score weights instruction-following at 4 and task-completion at 3, then divides by 7. Each task runs with and without the skill, so the lift from the skill is visible directly. The tasks and skills are public, in the task-evals-for-skills dataset , so you can inspect any scenario yourself.
This design is deliberate. The tasks come from published skills, so they mirror the real work teams write skills for, not frontier puzzles meant to find a model’s ceiling. That is why task-completion runs high for both models and why the signal that separates them is instruction-following: doing the work the specific way the skill asks. — from “ Claude Fable 5 vs Opus 4.8: The Mythos Hype Meets Reality ”
How a skill marketplace helps your team
The registry is the infrastructure. The value shows up in three places — and conveniently, they map to the three people who need to approve this.
Today, a new rep spends their first month reverse-engineering how the best people on the team work — shadowing calls, begging for prompt screenshots, slowly assembling their own janky toolkit. With a skill registry, day one looks different: the new rep gets the full team loadout auto-installed. The account research skill, the call prep skill, the follow-up writer — all the same versions your top performer runs. Ramp time has always been gated by how fast knowledge transfers from senior reps to new ones. A skill registry turns that transfer from months of osmosis into a download.
Every team has one rep whose call prep is unreasonably good. Right now, that advantage retires when they do. Turn their process into a skill, eval it, publish it, and their edge becomes the team’s floor. This is the highest-ROI knowledge management most GTM orgs will ever do, and almost nobody is doing it.
Model recommendations baked into skills fix the default-to-expensive problem without policing anyone. And legal approves the claims in the outreach skill once, instead of auditing a thousand freelanced prompts after a rep promises something the product doesn’t do. For reg

[truncated]
