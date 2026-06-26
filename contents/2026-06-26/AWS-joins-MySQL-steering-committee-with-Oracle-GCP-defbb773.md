---
source: "https://aws.amazon.com/blogs/opensource/open-governance-for-mysql-a-step-forward-for-the-community/"
hn_url: "https://news.ycombinator.com/item?id=48686840"
title: "AWS joins MySQL steering committee with Oracle, GCP"
article_title: "Open Governance for MySQL: A Step Forward for the Community | AWS Open Source Blog"
author: "deesix"
captured_at: "2026-06-26T14:52:37Z"
capture_tool: "hn-digest"
hn_id: 48686840
score: 2
comments: 1
posted_at: "2026-06-26T14:11:19Z"
tags:
  - hacker-news
  - translated
---

# AWS joins MySQL steering committee with Oracle, GCP

- HN: [48686840](https://news.ycombinator.com/item?id=48686840)
- Source: [aws.amazon.com](https://aws.amazon.com/blogs/opensource/open-governance-for-mysql-a-step-forward-for-the-community/)
- Score: 2
- Comments: 1
- Posted: 2026-06-26T14:11:19Z

## Translation

タイトル: AWS が Oracle、GCP とともに MySQL 運営委員会に参加
記事のタイトル: MySQL のオープン ガバナンス: コミュニティへの一歩 | AWS オープンソース ブログ
説明: MySQL — 世界中の何百万ものアプリケーションの背後にあるオープンソース データベース — は新たな章を開きます。本日、オラクルは、より広範なコミュニティがプロジェクトの開発と方向性に参加するための道筋を作り出す、MySQL のコミュニティ ガバナンス モデルを発表しました。この投稿では、AWS を使用する理由を説明します
[切り捨てられた]

記事本文:
MySQL のオープン ガバナンス: コミュニティへの一歩前進
MySQL — 世界中の何百万ものアプリケーションの背後にあるオープンソース データベース — は新たな章を開きます。本日、オラクルは、より広範なコミュニティがプロジェクトの開発と方向性に参加するための道筋を作り出す、MySQL のコミュニティ ガバナンス モデルを発表しました。
この投稿では、AWS がこの動きをサポートする理由と、それが MySQL コミュニティにとって何を意味するかについて説明します。
オープンガバナンスによりオープンソースが機能する
多様な貢献者と透明性のあるガバナンスを備えたオープンソース プロジェクトは、より優れたソフトウェアを生み出します。オープン ガバナンスにより、人々はユーザーから貢献者、そしてリーダーへと明確な道筋を得ることができ、組織はプロジェクトの将来に向けてエンジニアリングに投資する自信が得られます。
MySQL は、30 年近くにわたってインターネット インフラストラクチャの基盤として機能してきました。新興企業から世界最大手の企業に至るまで、何十万もの企業が最も重要なワークロードをこのプラットフォーム上で実行しています。コミュニティの参加方法を正式化することでその基盤が強化され、人々が将来の計画を立ててビジネスを構築するのに役立ちます。
新しいガバナンスモデルの仕組み
オラクルが MySQL を買収して以来、初めて、オラクル社外の組織がエンジンの構築方法とその行き先に関して明確な役割を持つことになりました。このモデルは、役割の進行を作成します。貢献者はコードと修正を送信し、コミッターは変更をレビューして承認し、プロジェクト リーダーはオプティマイザーや InnoDB などの主要なサブシステムを所有します。
これらの役割の上に、MySQL の長期的な方向性とリリース ポリシーを設定する運営委員会が設置されています。この委員会には、Oracle の過半数に加えて、クラウド プロバイダー、MySQL 顧客、オープンソース コミュニティが保持する Oracle 社外の 4 議席が含まれています。オラクル、2年間の任期の最初のメンバーを指名

m、その後、オラクル以外の議席はコミュニティ選挙に移動します。
これらすべてを支えるために、Oracle はコミュニティ MySQL のパブリック GitHub プレゼンスを立ち上げました。これは、これまで存在しなかった外部のコラボレーションと貢献のためのチャネルです。
AWS は、ユーザーとして、貢献者として、そして MySQL に依存するサービスの構築者として、15 年以上にわたって MySQL に深く投資してきました。現在、何万もの顧客が AWS で MySQL ワークロードを実行しています。 MySQL は当社のエコシステム内で最も重要なデータベースの 1 つであり、当社の顧客はその長期的な健全性に直接関与しています。
AWS では、オープンソースはすべての人にとって有益であると信じており、オープンソースの価値をお客様に提供し、AWS の優れた運用性をオープンソース コミュニティに提供することに尽力しています。その取り組みはシンプルな形で表れます。お客様が AWS でオープンソース データベースを実行し、問題が発生した場合、私たちはアップストリームで全員のために問題を解決します。
当社にはまさにそれを実行した実績があります。 PostgreSQL では、VACUUM を 6 倍高速化し、アップグレードを通じてレプリケーション スロットをそのまま維持し、自動バキューム構成変更の再起動要件を削除しました。 Redis の Linux Foundation フォークである Valkey では、全文検索とハイブリッド クエリのサポートが追加されました。また、テーブル数が多いデータベースのアップグレード中のメモリ不足エラーの修正やヒストグラム エラーの修正など、MySQL 自体のアップストリームにすでに修正を提供しています。
健全なアップストリーム プロジェクトは、MySQL を自分で実行するか、マネージド製品を使用するか、あるいはそれを中心としたツールや統合を構築するかに関係なく、MySQL に依存するすべての人に利益をもたらします。より多くのエンジニアがコードをレビューすると、より多くのバグが発見されます。設計上の決定がオープンに行われると、出荷される機能は現実世界の幅広いユースケースを反映します。ガバナンスが透明であれば、

彼らは、自分たちの貢献が評価され、自分たちの声が届けられるという確信を持ってプロジェクトに投資できます。
これは理論ではありません。これは、OpenJDK、Valkey、および幅広い参加によってソフトウェアが改善され、コミュニティが強化された他の数十のプロジェクトで私たちが経験したことです。
これが MySQL コミュニティにとって何を意味するか
このガバナンス モデルは、エコシステム全体のユーザー、貢献者、組織にとって、長期的なプロジェクトの健全性を示すシグナルです。
品質とセキュリティへの注目が高まる - コミッター、プロジェクトリーダー、コンポーネント間の監視による構造化されたレビュープロセスにより、より多くのエンジニアがコードの出荷前に正確性、パフォーマンス、セキュリティを検証することになります。
イノベーションの迅速化 — 明確な貢献経路と公的協力により、より広範なエコシステムが改善を提案し実現するための障壁が低くなります。
プロジェクトの将来に対する自信 — オラクル、エンドユーザー、オープンソースコミュニティの代表者を擁する運営委員会は、MySQL の方向性が単一のベンダーだけでなく、MySQL に依存する人々の利益を反映していることを意味します。
継続性と互換性 — ガバナンス モデルでは、安定性、下位互換性、リリースの品質を明確に優先しています。ユーザーとオペレーターは、破壊的な変更を心配することなく改善を採用できます。
アップストリーム プロジェクトが強化されるということは、マネージド サービス、セルフホスト デプロイメント、ツール、およびより広範なエコシステムなど、MySQL 上に構築されるすべての基盤が強化されることを意味します。
AWS は MySQL 運営委員会の委員を務めており、プロジェクトのロードマップとリリースの決定について直接発言することができます。私たちはその声を、弊社で MySQL を実行している顧客のために活用するつもりです。
AWS はオープンソース コミュニティに長期的に貢献しており、あらゆる分野で MySQL プロジェクトに積極的に取り組んでいます。

お客様のワークロードに最も直接的な影響を与えるもの:
パフォーマンス — 実際のワークロードの実行速度を決定するエンジンの部分、つまりクエリ オプティマイザー、クエリ実行、インデックス作成、InnoDB ストレージ エンジン、およびその下のキャッシュ レイヤーに焦点を当てています。
ベクター検索とインデックス作成 — オープンソース データベースのベクター機能を強化した AWS の経験が、コミュニティ全体での共同作業を基盤とした MySQL の新たなベクター サポートに貢献しています。
拡張フレームワーク — MySQL のコンポーネント インフラストラクチャにより、新しい機能をコア サーバー コードに組み込むのではなく、定義されたサービス インターフェイスを介して接続する読み込み可能なコンポーネントとして出荷できます。これは地域貢献に最もオープンな分野の 1 つであり、私たちはここに投資する予定です。
これらは、私たちがすでに行っている上流の作業に基づいています。何十万もの顧客に対してミッションクリティカルなワークロードを実行すると、すべての人に影響を及ぼす正確性、安定性、信頼性の問題といった本当の問題が表面化し、私たちは GitHub を通じてコミュニティ全体のために問題を修正しています。
要点は簡単です。MySQL の開発はオープンになりつつあり、AWS は開発の方向性を決めており、私たちはすでに上流で修正と改善に貢献しています。お客様は、MySQL をどこで実行してもメリットを得ることができます。
MySQL エコシステム全体の開発者、ユーザー、組織がガバナンス モデルを読み、どのように参加するかを検討することをお勧めします。オープンソースは人々が参加することで成長します。そしてこのモデルにより、これまでよりも簡単に貢献できるようになります。
Pravin Mittal は、AWS の Amazon Aurora のエンジニアリング ディレクターであり、数十万の顧客向けにマネージド MySQL および PostgreSQL サービスを構築するチームを率いています。彼は MySQL コミュニティ運営委員会で AWS を代表しています。
英語
トップに戻る
アマゾンは、

雇用主の機会均等: マイノリティ / 女性 / 障害 / 退役軍人 / 性同一性 / 性的指向 / 年齢。
×
フェイスブック
リンクイン
インスタグラム
けいれん
ユーチューブ
ポッドキャスト
電子メール
プライバシー

## Original Extract

MySQL — the open source database behind millions of applications worldwide — is opening a new chapter. Today, Oracle announced a community governance model for MySQL that creates pathways for the broader community to participate in the project’s development and direction. This post explains why AWS
[truncated]

Open Governance for MySQL: A Step Forward for the Community
MySQL — the open source database behind millions of applications worldwide — is opening a new chapter. Today, Oracle announced a community governance model for MySQL that creates pathways for the broader community to participate in the project’s development and direction.
This post explains why AWS supports this move and what it means for the MySQL community.
Open governance makes open source work
Open source projects with diverse contributors and transparent governance produce better software. Open governance gives people a defined path from user to contributor to leader, and it gives organizations the confidence to invest their engineering in a project’s future.
MySQL has been a foundational piece of internet infrastructure for nearly three decades. Hundreds of thousands of businesses — from startups to the world’s largest enterprises — run their most critical workloads on it. Formalizing how the community participates strengthens that foundation and helps people plan for the future and build their businesses.
How the new governance model works
For the first time since Oracle acquired MySQL, organizations outside Oracle have a defined role in how the engine is built and where it goes. The model creates a progression of roles: contributors submit code and fixes, committers review and approve changes, and project leads own major subsystems such as the optimizer or InnoDB.
Above those roles sits a Steering Committee that sets MySQL’s long-term direction and release policy. The committee includes four seats from outside Oracle — held by cloud providers, MySQL customers, and the open source community — alongside an Oracle majority. Oracle names the first members for a two-year term, after which the non-Oracle seats move to community election.
Underpinning all of this, Oracle has launched a public GitHub presence for community MySQL — a channel for outside collaboration and contribution that did not exist before.
AWS has been deeply invested in MySQL for over fifteen years — as users, as contributors, and as builders of services that depend on it. Tens of thousands of customers run MySQL workloads on AWS today. MySQL is one of the most important databases in our ecosystem, and our customers have a direct stake in its long-term health.
At AWS, we believe that open source is good for everyone, and we are committed to bringing the value of open source to our customers and the operational excellence of AWS to open source communities. That commitment shows up in a simple way: when customers run open source databases on AWS and hit a problem, we work upstream to fix it for everyone.
We have a track record of doing exactly that. In PostgreSQL, we made VACUUM six times faster, kept replication slots intact through upgrades, and removed the restart requirement for autovacuum configuration changes. In Valkey, the Linux Foundation fork of Redis, we added full-text search and hybrid query support. And we have already contributed fixes upstream to MySQL itself, including a fix for out-of-memory failures during upgrades on databases with large table counts and a fix for histogram errors.
A healthy upstream project benefits everyone who depends on MySQL — whether they run it themselves, use a managed offering, or build tools and integrations around it. When more engineers review code, more bugs get caught. When design decisions happen in the open, the features that ship reflect a wider range of real-world use cases. When governance is transparent, organizations can invest in the project with confidence that their contributions are valued and their voice is heard.
That’s not theory — it’s what we’ve experienced with OpenJDK, Valkey, and dozens of other projects where broad participation made the software better and the community stronger.
What this means for the MySQL community
This governance model is a signal of long-term project health — for users, contributors, and organizations across the ecosystem:
More eyes on quality and security — A structured review process with committers, project leads, and cross-component oversight means more engineers validating correctness, performance, and security before code ships.
Faster innovation — Clear contribution paths and public collaboration lower the barrier for the broader ecosystem to propose and deliver improvements.
Confidence in the project’s future — A Steering Committee with representation from Oracle, end users, and the open source community means MySQL’s direction reflects the interests of the people who depend on it — not just a single vendor.
Continuity and compatibility — The governance model explicitly prioritizes stability, backward compatibility, and release quality. Users and operators can adopt improvements without worrying about disruptive changes.
A stronger upstream project means a stronger foundation for everything built on MySQL — managed services, self-hosted deployments, tools, and the broader ecosystem alike.
AWS holds a seat on the MySQL Steering Committee, giving us a direct voice in the project’s roadmap and release decisions. We intend to use that voice for the customers running MySQL with us.
AWS contributes to open source communities for the long haul, and we are actively engaged with the MySQL project across the areas with the most direct impact on customer workloads:
Performance — We are focused on the parts of the engine that determine how fast real workloads run: the query optimizer, query execution, indexing, the InnoDB storage engine, and the caching layer underneath.
Vector search and indexing — AWS’s experience enhancing vector capabilities in open source databases is contributing to MySQL’s emerging vector support, building on collaborative work across the community.
Extensions framework — MySQL’s component infrastructure lets new capability ship as loadable components that connect through defined service interfaces instead of being built into the core server code. This is one of the most open areas for community contribution, and we plan to invest here.
These build on the upstream work we already do. Running mission-critical workloads for hundreds of thousands of customers surfaces real problems — correctness, stability, and reliability issues that affect everyone — and we work through GitHub to fix them for the whole community.
The takeaway is straightforward: MySQL’s development is opening up, AWS has a seat shaping where it goes, and we are already contributing fixes and improvements upstream. Customers benefit wherever they run MySQL.
We encourage developers, users, and organizations across the MySQL ecosystem to read the governance model and consider how they want to participate. Open source grows when people participate — and this model makes contributing easier than it’s ever been.
Pravin Mittal is Director of Engineering for Amazon Aurora at AWS, where he leads teams building managed MySQL and PostgreSQL services for hundreds of thousands of customers. He represents AWS on the MySQL Community Steering Committee.
English
Back to top
Amazon is an Equal Opportunity Employer: Minority / Women / Disability / Veteran / Gender Identity / Sexual Orientation / Age.
x
facebook
linkedin
instagram
twitch
youtube
podcasts
email
Privacy
