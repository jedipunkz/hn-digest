---
source: "https://arpitbhayani.me/blogs/three-skills/"
hn_url: "https://news.ycombinator.com/item?id=49126768"
title: "Three Claude Skills I Think Every Org Should Have"
article_title: "Three Claude Skills I Think Every Org Should Have"
author: "arpitbbhayani"
captured_at: "2026-07-31T19:08:48Z"
capture_tool: "hn-digest"
hn_id: 49126768
score: 2
comments: 1
posted_at: "2026-07-31T18:18:14Z"
tags:
  - hacker-news
  - translated
---

# Three Claude Skills I Think Every Org Should Have

- HN: [49126768](https://news.ycombinator.com/item?id=49126768)
- Source: [arpitbhayani.me](https://arpitbhayani.me/blogs/three-skills/)
- Score: 2
- Comments: 1
- Posted: 2026-07-31T18:18:14Z

## Translation

タイトル: すべての組織が持つべきだと私が考える 3 つのクロード スキル
説明: 人々が情報を探さずに仕事を完了できれば、会社はスムーズに運営されます。実際に人々を日々疲弊させているのは、エンジニアリング上の難しい問題ではありません。どのチケットを提出すべきか、どのチームが Slack でバグを修正すべきか、あるいはそれらの間にどのような目に見えないステップがあり、単純な s を取得するかがわかりません。
[切り捨てられた]

記事本文:
すべての組織が持つべきだと私が考える 3 つのクロード スキル
アルピット・バヤニ
ブログ
最初の原則から
--> 紙棚 本棚
セッション
--> コースでは、すべての組織が持つべきだと私が考える 3 つのクロード スキルについて説明します
エンジニアリング、データベース、システム。常に構築しています。
エンジニアリング組織が実際に働きやすい理由について考えるとき、私はいつも同じ 3 つのことに戻ってきます。
人々が情報を探さずに物事を進めることができれば、会社はスムーズに運営されます。実際に人々を日々疲弊させているのは、エンジニアリング上の難しい問題ではありません。どのチケットを提出すればよいのか、Slack でどのチームにバグを報告すればよいのか、あるいは単純なサービスを稼働させるまでの間にどのような目に見えないステップがあるのか​​がわかりません。
コードの最初の行を書いてから出荷するまでに 2 週間を失ったエンジニアを見てきました。これはコードが難しかったからではなく、データベースをプロビジョニングする方法を誰も教えてくれなかったからです。また、役立つはずだった Wiki ページが 6 か月前に名前変更されたダッシュボードにリンクされていたため、10 分以内に終了するはずのデバッグ セッションが 1 時間を超えてしまったこともありました。
そのどれもがエンジニア側のスキルギャップではありません。これはツールのギャップであり、一度気づくとどこにでも見られるようになります。ここでは、すべてのエンジニアリング組織にとって構築する価値のある 3 つのスキルを紹介します。
このスキルは、一連のビジネス ロジックを作成する前に、データベースをプロビジョニングし、適切なタグを使用して中央シンクにログを接続し、アイデンティティ プロバイダーにサービスを登録し、サービスを運用環境にデプロイする必要があります。これらの手順のすべてが手動のままだと、誰かが別のチームを待ったり、間違った推測をしたりする場所になり、その摩擦が初日の作業を遅らせる原因になります。
同社の標準サービステンプレート。
インフラストラクチャ プロビジョニング API (データベース、秒)

レット、ネットワーキング)。
ID プロバイダーの登録フロー。
ロギング パイプラインのタグ付け規則。
エンドツーエンドでは、エクスペリエンスは 1 つのコマンド (エージェント ループ) で入力され、ライブ運用 URL で出力される必要があります。
このスキルは、フォワード デプロイ、ロールバック、環境を対象としたテスト実行、およびデプロイ後のスモーク テストを 1 つの統合されたフローとして処理する必要があります。ほとんどのチームは、順方向パスを構築して定期的に実行するだけなので、ロールバックは、まさにプレッシャーの下で使用されるパスであるにもかかわらず、パイプラインの中で最も信頼性の低い部分になります。
サービスの展開履歴と現在のリビジョン。
データベースの移行状態。ロールバックによってスキーマが古いコードで処理できない形状のままになることはありません。
ターゲット環境の構成。
サービス用に定義された一連のスモーク テスト エンドポイント。
エンドツーエンドの展開では、実際のターゲット環境に対してテストを実行し、徐々にロールアウトし、スモーク テストを自動的に実行し、人間が介入したり別の手順を覚えたりすることなく、失敗時にロールバックする必要があります。
このスキルは、単に生データを指すのではなく、実際の症状をその特定のサービスの既知の障害パターンと照合し、関連するログ、トレース、ダッシュボードへの直接リンクを表示し、診断チェックを実行する必要があります。静的 Runbook は最終的には古くなり、誰かがそれを必要とする頃には、リンクや手順が間違っていることがほとんどです。
サービスの過去のインシデントの実行履歴。
依存関係とヘルス ダッシュボードへのリンク。
インシデントの遡及記録から抽出された検証済みの障害パターン。
エンドツーエンドでは、誰かが症状を説明し、「ログを確認してください」という一般的な提案ではなく、考えられる原因、それを裏付ける証拠、およびそれを検証するためのリンクを返します。
ドキュメントを書く代わりにツールを構築する理由
ドキュメントは腐る。古くなった Wiki ページは、誰よりも先に何か月も人々を誤解させる可能性があります

気づきます。壊れたツールは、次に使用する人の目の前で大きな音を立てて故障します。つまり、実際には修理されることになります。
この違いこそが、金曜日の午後にドキュメントを作成する方が常に安価であるにもかかわらず、ドキュメントをさらに作成するのではなく実際のツールを構築するべきであるという議論全体です。
AI エージェントが人々に代わってサービスを作成し、変更を送信し、インシデントをトリアージし始めている現在、その重要性はさらに高まっています。古い W​​iki ページを読んでいるエージェントは、疲れたエンジニアが午前 2 時に抱えているのとまったく同じ問題を抱えています。それは、指示が間違っていることを知る信頼できる方法がないということです。実行可能ツール内に知識が存在する企業は、それに作用するものが人であるかエージェントであるかに関係なく、デフォルトで正しい動作を取得します。
最初にビルドのデバッグを行います。インシデントは優先順位に関係なく独自のスケジュールで発生するため、すべてのインシデントが検証済みの障害パターンを追加する機会となり、その効果はほぼ即座に目に見えます。
定期的に実行されるロールバック パスと必須のスモーク テストを使用して、2 番目にデプロイメントを構築します。最も頻繁に中断されるサービスから始めます。
最後に create-service をビルドします。これは 3 つの中で最も多くのインフラストラクチャに影響しており、混乱を除去するのではなく再配置するだけなので、中途半端に構築されたバージョンは何もないより悪いです。
インシデントと大まかなオンボーディング エクスペリエンスの両方を、あれば便利ではなく要件としてこれらのツールにフィードバックします。診断が遅いと、デバッグ ツールに対するチケットになります。 2 週間の 1 日が create-service に対するチケットになります。
ツールごとに 1 つのメトリクスを追跡します。つまり、新しいリポジトリから実際の運用環境へのデプロイメントまでの時間、デプロイメントの頻度とロールバック率、およびインシデントの診断にかかる平均時間です。これら 3 つはいずれも測定コストが低く、議論するのが困難です。
最悪のサービス、ページングするサービスから始めます

が最も多く、修正に最も時間がかかります。フリートの残りの部分を心配する前に、数字がそこに移動していることを証明してください。
脚注: どの企業も最終的には 3 つのツールを必要としますが、ほとんどの場合適切に構築することはできません。 create-app の初日は、チェックリストではなく、実際の運用環境のデプロイで終了する必要があります。デプロイでは、ロールバック、環境スコープのテスト、および必須のスモーク テストを、オプションではなくコア機能として扱う必要があります。 -デバッグでは、ダッシュボードを指定して残りを推測に任せるのではなく、実際の症状をデータへの直接リンクを使用して既知のパターンと照合する必要があります。
ドキュメントとしてのみ存在する場合、この 3 つはすべて失敗します。ツールが大声で失敗して修正される一方で、ドキュメントは静かに朽ちていくからです。 AI エージェントが人々に代わってサービスを作成し、インシデントを優先順位付けし始めると、これが、正しい知識を継承するエージェントと、古いものに自信を持って行動するエージェントとの違いになります。
これが役に立ち、興味深いと思われた場合は、
RSS フィードを購読すると、新しいフィードを公開したときに通知が届きます。
Razorpay のプリンシパル エンジニア II - Agent Studio の構築、GCP Memorystore および Dataproc の元スタッフエンジニア、DiceDB の作成者、元 Amazon Fast Data、元エンジニアのディレクター。 SREと
Unacademy のデータ エンジニアリング。私は自分自身を通じてエンジニアリングへの好奇心を刺激します
綿密なエンジニアリングビデオ
YouTube
そして私のコース
145,000 人のエンジニアが読んだ Arpit のニュースレター
現実世界のシステム設計、分散システム、または
非常に賢いアルゴリズムを深く掘り下げてみましょう。
LinkedIn で購読する Substack で購読する
このウェブサイトに掲載されているコースは、
リログディープテック株式会社株式会社
203, Sagar Apartment, Camp Road, Mangiral Plot, アムラヴァティ, マハーラーシュトラ州, 444602
GSTIN: 27AALCR5165R1ZF
YouTube (210k) Twitter (120k) LinkedIn (280k) GitHub (7k)
© アルピット・バヤニ、2025

## Original Extract

A company runs smoothly when people can get stuff done without hunting for information. What actually wears people out day to day is not the hard engineering problem. It is not knowing which ticket to file, which team to bug on Slack, or what invisible step stands between them and getting a simple s
[truncated]

Three Claude Skills I Think Every Org Should Have
Arpit Bhayani
Blogs
From the First Principles
--> Papershelf Bookshelf
Sessions
--> Courses Talks Three Claude Skills I Think Every Org Should Have
engineering, databases, and systems. always building.
I keep coming back to the same three things whenever I think about what makes an engineering org actually pleasant to work in.
A company runs smoothly when people can get stuff done without hunting for information. What actually wears people out day to day is not the hard engineering problem. It is not knowing which ticket to file, which team to bug on Slack, or what invisible step stands between them and getting a simple service live.
I have seen engineers lose two weeks between writing their first line of code and shipping it, not because the code was hard, but because nobody could tell them how to get a database provisioned. I have also seen debugging sessions that should have finished in under 10 minutes stretch beyond an hour because the wiki page that was supposed to help linked to a dashboard that had been renamed six months earlier.
None of that is a skills gap on the engineer’s part. It is a tooling gap, and once you notice it, you see it everywhere. So, here are the three skills worth building for every engineering org.
This skill should provision the database, wire up logging to the central sink with the right tags, register the service with the identity provider, and deploy the service to production, all before anyone writes a line of business logic. Every one of those steps, if left manual, becomes a place where someone waits on another team or guesses wrong, and that friction is what makes day one slow.
The company’s standard service template.
Infrastructure provisioning APIs (databases, secrets, networking).
The identity provider’s registration flow.
The logging pipeline’s tagging conventions.
End to end, the experience should be one command (agentic loop) in, a live production URL out.
This skill should handle forward deploys, rollbacks, environment-scoped test runs, and post-deploy smoke tests as one integrated flow. Most teams only build and regularly exercise the forward path, so rollback ends up being the least reliable part of the pipeline, even though it is exactly the path that gets used under pressure.
The service’s deployment history and current revision.
Database migration state, so rollback does not leave the schema in a shape older code cannot handle.
Target environment configuration.
A defined set of smoke test endpoints for the service.
End to end, a deployment should run tests against the real target environment, roll out gradually, execute smoke tests automatically, and roll back on failure without a human needing to intervene or remember a separate procedure.
This skill should match live symptoms to known failure patterns for that specific service, surface direct links to the relevant logs, traces, and dashboards, and run diagnostic checks instead of merely pointing at raw data. Static runbooks eventually go stale, and by the time someone needs them, the links or steps are usually wrong.
A running history of past incidents for the service.
Dependencies and links to their health dashboards.
Verified failure patterns extracted from incident retros.
End to end, someone describes a symptom and gets back a likely cause, the evidence supporting it, and links to verify it, instead of a generic suggestion to “check the logs.”
Why Build Tools Instead of Writing Docs
Docs rot. A stale wiki page can mislead people for months before anyone notices. A broken tool fails loudly, right in front of the next person who uses it, which means it actually gets fixed.
That difference is the whole argument for building real tools instead of writing more documentation, even though documentation is always cheaper to produce on a Friday afternoon.
It matters even more now that AI agents are starting to create services, ship changes, and triage incidents on people’s behalf. An agent reading a stale wiki page has the exact same problem a tired engineer has at 2 AM: there is no reliable way to know the instructions are wrong. Companies whose knowledge lives inside executable tools get correct behavior by default, whether the thing acting on it is a person or an agent.
Build debugging first. Incidents happen on their own schedule regardless of what you prioritize, so every incident becomes a chance to add a verified failure pattern, and the payoff is visible almost immediately.
Build deployment second, with a rollback path that gets exercised regularly and mandatory smoke tests. Start with whichever service breaks most often.
Build create-service last. It touches the most infrastructure of the three, and a half-built version is worse than none because it merely relocates the confusion instead of removing it.
Feed both incidents and rough onboarding experiences back into these tools as a requirement, not a nice-to-have. A slow diagnosis becomes a ticket against the debugging tool. A two-week day one becomes a ticket against create-service.
Track one metric per tool: time from a fresh repository to a live production deployment, deployment frequency alongside rollback rate, and mean time to diagnose an incident. All three are inexpensive to measure and difficult to argue with.
Start with your worst service, the one that pages the most and takes the longest to fix. Prove the numbers move there before worrying about the rest of the fleet.
Footnote: every company eventually needs three pieces of tooling that most never build properly. create-app should end day one with a live production deploy, not a checklist. deploy- should treat rollback, environment-scoped tests, and mandatory smoke tests as core, not optional, features. -debugging should match live symptoms to known patterns with direct links to data, not point at a dashboard and leave the rest to guesswork.
All three fail when they exist only as documentation, since docs rot silently while tools fail loudly and get fixed. As AI agents start creating services and triaging incidents on people’s behalf, this becomes the difference between an agent inheriting correct knowledge and one confidently acting on something stale.
If you find this helpful and interesting,
subscribe to my RSS feed and get notified the moment I publish a new one.
Principal Engineer II at Razorpay - building Agent Studio, Ex-staff engg at GCP Memorystore & Dataproc, Creator of DiceDB , ex-Amazon Fast Data, ex-Director of Engg. SRE and
Data Engineering at Unacademy. I spark engineering curiosity through my
no-fluff engineering videos on
YouTube
and my courses
Arpit's Newsletter read by 145,000 engineers
Weekly essays on real-world system design, distributed systems, or a
deep dive into some super-clever algorithm.
Subscribe on LinkedIn Subscribe on Substack
The courses listed on this website are offered by
Relog Deeptech Pvt. Ltd.
203, Sagar Apartment, Camp Road, Mangilal Plot, Amravati, Maharashtra, 444602
GSTIN: 27AALCR5165R1ZF
YouTube (210k) Twitter (120k) LinkedIn (280k) GitHub (7k)
© Arpit Bhayani, 2025
