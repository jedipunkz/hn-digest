---
source: "https://vsyrakis.dev/blog/you-dont-know-devops/"
hn_url: "https://news.ycombinator.com/item?id=49226410"
title: "You don't know what DevOps means"
article_title: "You don't know what DevOps means - Vasilios Syrakis"
author: "adletbalzhanov"
captured_at: "2026-08-08T22:18:15Z"
capture_tool: "hn-digest"
hn_id: 49226410
score: 1
comments: 0
posted_at: "2026-08-08T22:17:22Z"
tags:
  - hacker-news
  - translated
---

# You don't know what DevOps means

- HN: [49226410](https://news.ycombinator.com/item?id=49226410)
- Source: [vsyrakis.dev](https://vsyrakis.dev/blog/you-dont-know-devops/)
- Score: 1
- Comments: 0
- Posted: 2026-08-08T22:17:22Z

## Translation

タイトル: DevOps の意味がわからない
記事のタイトル: DevOps の意味を知らない - Vasilios Syrakis

記事本文:
メイン コンテンツにスキップ Vasilios Syrakis ブログ リンク PGP CV 検索 2026 年 7 月 30 日発行 DevOps の意味を知らない
読了 8 分 • devOps 共感と DevOps にはどのような関係があるのでしょうか?
最近、人々が DevOps という用語を奇妙な文脈で乱用しているのを目にします。もしかしたら、この用語は私が置き去りにされるほど進化しているのかもしれませんが、その背後にある歴史を知らない人々によってこの用語が採用された可能性の方が高いと思います。それがどこから来たのか、それがどのような問題を解決していたのか、どのようにそれを行ったのか、そしてそれが時間の経過とともにどのように変化したかを知らずに、どうしてそれが何であるかを確信できるでしょうか?このトピックについてどのように教育するかを考えているうちに、私は自分自身の無知に近づき始めたことを認めなければなりません。8 年前のジェズ・ハンブルのビデオを見た後、私には共有できる一口サイズに蒸留できるほどの 10 年以上の深い知識がないことに気づきました。 DevOps とは何かを知りたい場合は、Jez の説明以上に適切な説明は見つからないと思います。ただし、ソフトウェアの提供方法における文化的な変化がどのようにして役職に変わったのかを知りたい場合は、読み続けてください。初めに壁があった
DevOps の文化運動は、Patrick Debois がこの用語を作り、最初の DevOpsDays カンファレンスを主催した後、2009 年に始まりました。これは、責任の共有、自動化、迅速なフィードバックの哲学を意図したものでした。この用語自体は、開発者がソフトウェアをより速く提供することを優先し、運用側がすべてを安定させようとしていた当時、開発と運用の間のサイロを打破することを意味していました。相反する2つのゴール。
DevOpsDays では、Flickr の Paul Hammond 氏が「10 回以上のデプロイ/日」と題した講演を行いました。これは人々の心を驚かせました

当時、エンタープライズ ソフトウェアは年に数回導入されていました。 1 日に 10 回ソフトウェアをデプロイすると、システム全体がダウンするなど想像できませんでした。数年後、Amazon は 11.6 秒ごとにデプロイしていると発表しました。大手テクノロジー企業は、数分で機能を出荷し、実験を行い、バグを修正していました。従来の企業は、毎月のデプロイメントについて紙のチェックリストの項目に鉛筆でチェックを入れるプロセスが依然として必要だったため、恐怖を感じていました。ベンダー工業団地
2010 年から 2015 年にかけて、誇大宣伝列車が駅を出発しました。企業は、DevOps が生み出すスピードを求めていました。企業のリーダーたちはそれを実現する方法を模索しました。彼らは、「私たちは文化を変え、開発者に自主性を与え、その働き方をサポートするツールを構築した」という先駆者による技術講演を読みましたが、覚えているのは最後の部分だけでした。テクノロジーベンダーは、文化の変化によってのみ解決できることを約束するツールを喜んで販売しました。今日、冷蔵庫や歯ブラシが「AI ネイティブ」になっているのと同じように、多くのソフトウェア ツールが DevOps の実現に役立つとしてマーケティングを開始しました。そのため、会社の機能全体にわたる責任の共有、共感、自律性、継続的なフィードバック ループの代わりに、Jenkins、Puppet、Docker、Terraform、YAML テンプレートを入手しました。基本的にはDevOpsのマクドナルドです。サイロを別の色にペイントする
企業は、DevOps を「買う」だけで済むと考えたため、インフラストラクチャについてさらに理解できるように開発者を訓練したり、運用にコードの書き方を教えたりするのではなく、DevOps エンジニアの求人広告を掲載しました。システム管理者には、DevOps が肩書きではなく文化であることを認識し、原則を守るか、40% の給与増額を受け入れ、より優れたツールを使用して同じ問題に取り組むかの選択が迫られました。グーグル

聖杯を落とします
2016 年に誰もが DevOps に興味を持ったとき、「サイト信頼性エンジニアリング: Google による運用システムの運用方法」がリリースされました。
Google は、SRE を「運用をソフトウェアの問題であるかのように扱うことで得られるもの」と定義しました。この一文は私に個人的にインスピレーションを与え、世界をこのように見るようになりました。 SRE の原則は、自動化によって労力を排除し、システムの信頼性を測定してサービス レベル目標 (SLO) を付加し、エラー バジェットを使用して速度と安定性のバランスを取ることであり、顧客がインターネットの天気と問題が発生しているサイトの区別がつかない程度にミスが少ない限り、基本的にミスは許されます。当時、Google SRE であればエリートとみなされ、その称号を獲得するために必要なスキルの積み重ねは登るべき高い塔でした。 SRE とは、インフラストラクチャと運用、そして高度に有能なソフトウェア開発者という、これまで全く別の分野と考えられていた 2 種類のエンジニアのスキルセットを持つ人物のことです。彼らは、開発チームと同じ技術スタックを理解し、コードベースに飛び込み、本番環境の複雑なバグのトラブルシューティングを行うことができることが期待されていました。理論的には、これは速度と安定性のバランスをとる DevOps の夢でした。実際には、期待は信じられないほど高かったのです。ビジネスの世界で起こっているように見えたのは、企業がそのスキルセットを奨励し育成する代わりに、最近の DevOps チームを SRE にブランド名を変更し、開発チームが原因で発生したバグのオンコールを担当させたことです。ハイギアへのシフト
2015 年から 2018 年にかけて、マイクロサービスとコンテナーが大々的に宣伝され、最終的にビルド ツールがすべてのコミットでタスクを実行できるほど成熟しました。 T

彼は「シフトレフト」運動の基礎を築きました。コードをテスター、セキュリティ、運用に投げ込むのではなく、それらの責任はコードが書かれた場所の近くに移されました。これは、マージが許可される前に、自動テスト、脆弱性スキャン、リンティングが発生していたことを意味します。これにより、Amazon の Werner Vogels 氏が 2006 年に「あなたが構築し、あなたが実行する」という有名なフレーズで打ち出したアイデアが実現しました。基本的に開発者は、実稼働環境でサービスがダウンした場合に備えて窮地に陥り、より信頼性の高いコードを書くという自然な動機が生まれます。
しかし、話が進むにつれて、私たちは少し左にシフトしすぎました。開発者は、独自の Terraform を作成し、リバース プロキシとデータベース クラスターを設定し、Kubernetes マニフェストを構成し、パイプラインを構築し、Prometheus でアラートを設定することが期待されていました。開発者には、実際のビジネス ロジックに取り組むための十分な時間がほとんど残されていませんでした。現在の修正点
最終的に、業界は、すべての開発者がビジネス ロジックの実行に必要なあらゆる小さなインフラストラクチャを習得することを期待することはできないことに気づきました (おそらく今も気づいています)。 YAML を書くために彼らにお金を払っているわけではありません。こうしてプラットフォームエンジニアリングの登場です。そのアイデアは、開発者をほぼ外部の顧客のように扱い、クラウド コンポーネントのように動作する製品を社内で構築し、開発者が認可された方法で迅速に行動できるようにする「ゴールデン パス」を提供することでした。そのため、DevOps チームがビルド パイプラインを設定し、デプロイを実行し、オンコールに対応するのではなく、プラットフォーム チームがセルフサービス ツールと、開発者が自分で安全にデプロイできるガードレールを備えたインフラストラクチャを構築します。正しく実行すると、開発者の自主性を維持しながら認知的負荷が軽減されます。これを正しく実行しないことがよくありますが、その一例です。

プラットフォームに通過すべきフープが非常に多く、AWS コンソールをロードして ClickOps を通じてすべてをセットアップする方が簡単な場合は、これが考えられます。このプラットフォームの約束は、ユーザーが自分のコードを所有するということであり、kubernetes、CICD、モニタリング、その他の定型的な設定に対処する必要がないようにパスが舗装されています。結論

## Original Extract

Skip to Main Content Vasilios Syrakis Blog Links PGP CV Search Published on July 30, 2026 You don't know what DevOps means
8 minutes read • devops What does empathy have to do with DevOps?
These days I see people throwing around the term DevOps in contexts that strike me as odd. Maybe the term has evolved to the point where I’ve been left behind, but I think it’s more likely that the term has been adopted by people who don’t know the history behind it. Without knowing where it came from, what problem it was solving, how it did and did not do that, and how it’s changed over time, how can you be confident that you know what it is? I have to admit that while thinking about how to educate on this topic, I started to approach my own ignorance, and after watching a video from Jez Humble from 8 years ago , I realised that I don’t have 10+ years of deep knowledge that I can distill into bite size pieces to share. If you want to know what DevOps is, I don’t think you could find a better explanation than from Jez. However, if you’re interested in finding out how a cultural shift in how we deliver software got transformed into a job title, read on. In the Beginning there was a Wall
The cultural movement of DevOps started in 2009 after Patrick Debois coined the term and hosted the first DevOpsDays conference. It was meant to be a philosophy of shared responsibility, automation, and rapid feedback. The term itself was meant to represent breaking down the silo between Dev and Ops, back at a time where developers were prioritizing delivering software faster, while operations were trying to keep everything stable. Two goals that conflicted.
At DevOpsDays, Paul Hammond from Flickr gave a talk titled “10+ Deploys Per Day”. This blew people’s minds because at the time, enterprise software was deployed a few times a year. People couldn’t imagine deploying software 10 times in a single day without the entire system blowing up. A couple years later, Amazon said they were deploying every 11.6 seconds. Big tech were shipping features, doing experiments, and fixing bugs in minutes. Traditional enterprise were terrified, because their process still involved ticking items off a paper checklist with a pencil for their monthly deploys. Vendor Industrial Complex
Between 2010 and 2015 the hype train had left the station. Companies wanted the speed that DevOps seemed to create. Corporate leadership went looking for ways to implement it. They read tech talks by the pioneers that said “We changed our culture, gave developers autonomy, and built tools to support that way of working” but all they remembered was the last bit. Tech vendors were more than happy to sell them tools that promised to solve what only a cultural change could. Just like today how your fridge and toothbrush are “AI-Native”, many software tools started marketing that they could help achieve DevOps. So instead of shared responsibility, empathy, autonomy across company functions, and continuous feedback loops, they got Jenkins, Puppet, Docker, Terraform, and YAML templates. Basically the McDonalds of DevOps. Painting the Silo a Different Color
Companies thought they could just “buy” DevOps, so rather than training developers to understand more about infrastructure, or teaching operations to write code, they posted a job ad for DevOps Engineers. Sysadmins were presented with the choice to stand by their principles, knowing that DevOps was a culture, not a title, or they could accept a 40% pay bump and work on the same problems with shinier tools. Google drops the Holy Grail
Just when everyone was bought into DevOps in 2016, “Site Reliability Engineering: How Google Runs Production Systems” was released.
Google defined SRE as “what you get when you treat operations as if it’s a software problem.” That line inspired me personally to start seeing the world this way. The principles of SRE were to eliminate toil through automation, measure system reliability and attach service level objectives (SLOs) to it, and balance velocity with stability through the use of error budgets, where you’re essentially permitted to make mistakes, as long as they were so few that a customer couldn’t tell the difference between internet weather and the site having problems. At that time, if you were a Google SRE, you were considered elite - the stack of required skills to gain that title was a tall tower to climb. An SRE was someone with the skillset of two types of engineers that had previously been considered entirely separate tracks - infrastructure and operations on the one hand, highly capable software developer on the other. They were expected to be able to understand the same tech stack as the dev team, and jump into the codebase to troubleshoot complex bugs in production. In theory this was the DevOps dream, balancing velocity with stability. In reality, the expectations were impossibly high. What seemed to happen in business is that instead of encouraging and cultivating that skillset, companies took their recent DevOps teams, rebranded them to SREs, and made them handle the on-call for bugs caused by the dev team. Shifting into High Gear
Between 2015 and 2018, perhaps you can remember, there was much hype about microservices and containers, and build tooling finally matured enough to run tasks on every single commit. This laid the foundation for the “Shift Left” movement. Instead of throwing code at testers, security, and operations, those responsibilities were moved closer to where the code was written. This meant that automated tests, vulnerability scans, and linting were happening before a merge was even allowed. This enabled the idea that Werner Vogels from Amazon had laid out back in 2006 with his famous phrase: “You build it, you run it.” Basically developers were put on the hook for when their service blows up in production, creating a natural incentive for them to write more reliable code.
But as the story goes, we shifted left a bit too far. Developers were expected to write their own Terraform, set up reverse proxies and database clusters, configure Kubernetes manifests, build their pipelines, and set up alerts with Prometheus. Devs were left with barely enough time to work on actual business logic. The Current Correction
Eventually, the industry realised (and is probably still realising) that you can’t just expect every dev to master every small piece of infra required to run their business logic. You’re not paying them to write YAML. Thus the advent of Platform Engineering . The idea was to treat devs almost like external customers, building out products internally that behave like cloud components, providing “golden paths” to allow devs to move fast in sanctioned ways. So now, instead of a DevOps team setting up build pipelines, doing the deployments for you, and handling your on-call - a platform team builds self-service tooling and infrastructure with guardrails that developers can deploy safely on their own. When done correctly it reduces cognitive load while preserving developer autonomy. It can be easy to not do this correctly and one example of that would be if your platform has so many hoops to jump through that it’s easier to load up the AWS Console and set everything up through ClickOps. The promise of the platform is that you own your code, and a path has been paved for you so that you don’t have to deal with setting up kubernetes, CICD, monitoring, and other boilerplate. Conclusion
