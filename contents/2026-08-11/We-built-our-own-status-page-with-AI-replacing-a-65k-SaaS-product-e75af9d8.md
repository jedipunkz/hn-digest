---
source: "https://www.fivetran.com/blog/we-built-our-own-status-page-with-ai-replacing-a-65k-saas-product"
hn_url: "https://news.ycombinator.com/item?id=49260068"
title: "We built our own status page with AI, replacing a $65k SaaS product"
article_title: "We built our own status page with AI, replacing a $65k SaaS product | Blog | Fivetran"
author: "georgewfraser"
captured_at: "2026-08-11T15:50:58Z"
capture_tool: "hn-digest"
hn_id: 49260068
score: 2
comments: 0
posted_at: "2026-08-11T15:41:52Z"
tags:
  - hacker-news
  - translated
---

# We built our own status page with AI, replacing a $65k SaaS product

- HN: [49260068](https://news.ycombinator.com/item?id=49260068)
- Source: [www.fivetran.com](https://www.fivetran.com/blog/we-built-our-own-status-page-with-ai-replacing-a-65k-saas-product)
- Score: 2
- Comments: 0
- Posted: 2026-08-11T15:41:52Z

## Translation

タイトル: AI を使用して独自のステータス ページを構築し、65,000 ドルの SaaS 製品を置き換えました
記事のタイトル: AI を使用して独自のステータス ページを構築し、65,000 ドルの SaaS 製品を置き換えました |ブログ |ファイブトラン
説明: 2 人のエンジニアが AI コーディング エージェントを使用して、本番環境に対応したステータス ページを設計、構築、出荷しました。何がうまくいき、何がうまくいかなかったのか、また別の方法で実行するのかを以下に示します。

記事本文:
AI を使用して独自のステータス ページを構築し、65,000 ドルの SaaS 製品を置き換えました。ブログ |ファイブトラン
English English Deutsch Français Español 製品
プラットフォームの機能
すべてのデータ移動ニーズに単一のプラットフォームで対応
自動化された変換で洞察を加速します
安心してデータを移動できる厳格な組み込みセキュリティ
管理されたデータ移動によりデータを把握し、保護し、拡張します
Fivetran を他のツールやシステムと簡単に統合して、大規模なワークフローを最適化します
ウェアハウス データをビジネス ツールに移行して効果を高める
パフォーマンスを損なうことなく、すべてのデータを安全に移動します
エージェント AI のスタックを将来も保証する
コンプライアンスを損なうことなくデータにアクセスできるようにする
クラウド導入を通じてビジネスを加速する
顧客のデータに簡単に接続
一元化され管理されたデータへの信頼性の高いアクセスにより、AI イノベーションを強化します。
リアルタイム分析により、より賢明なビジネス上の意思決定を推進
ブログ お客様の事例 リソース センター ドキュメント オンボーディング ニュース イベント プロフェッショナル サービス ポッドキャスト 注目のリソース Fivetran の動作を見る インタラクティブ デモを見る 信頼 サポート ポータル ステータス 英語
English Deutsch Français Español 営業担当者へのお問い合わせ ログイン 無料で始める
データの洞察 AI を使用して独自のステータス ページを構築し、65,000 ドルの SaaS 製品を置き換えました
2 人のエンジニアが AI コーディング エージェントを使用して、本番環境に対応したステータス ページを設計、構築、出荷しました。何がうまくいき、何がうまくいかなかったのか、また別の方法で実行するのかを以下に示します。 Fivetran の 2 人のエンジニア (Valentina Mačković と Jelena Kostic) は、要件から運用まで AI コーディング エージェントを使用して、年間 65,000 ドルの SaaS 製品の代替品を約 4 か月で設計、構築、出荷しました。現在、Fivetran の公開ステータス ページと約 26,000 人の通知購読者にサービスを提供しています。これが実際にかかる費用です、実際にかかる費用はいくらですか

あなたはうまくいきました、そして私たちが何を間違えたのか。
背景: 置き換えていたもの
Fivetran のステータス ページは、Atlassian Statuspage 上で実行される、大幅にカスタマイズされたアプリケーションでした。それは設計された目的どおりに機能しましたが、私たちの使用法は Statuspage フレームワークのネイティブ機能をはるかに超えていました。ステータス ページには、個別に監視されている多数のサービスが表示されるため、サービスごとのステータスと各サービスの稼働時間が必要でした。これは、ホストされる製品が実際に構築されたことのない形状です。そこで私たちは、BigQuery データ ウェアハウスから稼働時間を計算する Google Cloud Functions、インシデント トラッカーの同期状態、そしてそれをインシデント管理プロセスに組み込む自動化というカスタマイズを最上位に重ねました。
その結果、台帳の両側に実際の運用コストがかかるシステムが誕生しました。
費用: 年間約 65,000 ドルの費用がかかります。これは、すべての通知サブスクライバーが請求対象であるという事実によって大きく左右されます。それらは約 26,000 件ありました。
時間 : ステータス ページの管理、トラブルシューティング、修正に一貫した時間を費やし、年間少なくとも 10,000 ドルのエンジニアの時間を費やしていました。
パフォーマンス: ページの読み込みが遅く、最初の試行で読み込みに失敗することがありました。他のすべてが利用できないときに利用できるようにすることが全体の目的であるページにとって、これは重大な障害モードです。
サイレント障害: 更新が伝播されず、欠陥が検出されないケースが発生しました。
レート制限: 当社独自の自動化では、ベンダーの API レート制限に遭遇しました。
アトラシアンにとって公平を期すために言うと、これはいずれもスキャンダルではありません。これは、汎用製品をその意図された形状をはるかに超えて推し進めたときに起こることであり、加入者 26,000 人における加入者ごとの価格の通常の計算でもあります。この投稿では、要件を超えて成長したシステムを AI で置き換えることがどのように可能なのかを説明します。

エント。
ステータスページは珍しい場所にあります。これは、少数の読み取りエンドポイント、管理パネル、電子メール キューを備えた小さなアプリケーションですが、実稼働規模の影響をもたらします。 Fivetran にとって悪い日があったとき、それは何千もの顧客が注目する成果物であり、大量の通知イベントが発生すると 50,000 通を超える電子メールが送信される可能性があります。複雑さは低く、一か八かのリスクが伴います。
その組み合わせこそが、この作品を面白くしているのです。 1 年前、DIY ステータス ページは簡単に「ノー」だったでしょう。エンジニアリング コストのせいでライセンスが矮小化されました。何が変わったかというと、AI が特定のクラスのソフトウェア (仕様が明確で、複雑さは控えめで、レンタル費用が高い) の境界線がそこに移ったことです。この線が十分に大きく変化したかどうかは、すぐにわかるでしょう。
構築か購入かを決定した理由は何ですか
重さの順に3つ挙げると、
信頼性の制御。サードパーティに依存するステータス ページは、そのサードパーティと同等の信頼性しかあり得ず、遅い部分を修正することはできませんでした。これを所有するということは、それを自社のコア プラットフォームと外部プロバイダーの両方から分離できることを意味します。
価格設定モデルが製品と一致していませんでした。私たちは、構造的にはメーリング リストの料金を購読者ごとに支払っていました。コストは顧客の成長に応じて増加しましたが、提供される価値は増加しませんでした。
私たちはソフトウェアだけではなく、組織的な答えを求めていました。 AI エージェントを使用している 2 人のエンジニアが、意図的に PRD から本番環境に何かを持ち込む可能性はありますか? 「はい」の場合、それは再利用可能な機能です。ステータスページは、間違っても爆発範囲内で生き残ることができるため、それを確認するのに最適な場所でした。信頼できるまでは、古いものと並行して実行できました。莫大な時間、資金、労力をつぎ込まずに実験は失敗に終わったとも言えます。
結論の境界についても正直になりたいと思っています。 Fivetranの自作P

roduct は SaaS 製品であり、AI によって SaaS が時代遅れになると主張しているわけではありません。ステータス ページは、データ移動プラットフォームよりも置き換えがはるかに簡単です。しかし、同じロジックがどのベンダーにも限界部分で当てはまります。つまり、一般的な製品の薄く、明確に指定されたスライスを使用するほど、方程式のビルド側が大きく動きます。
チームが下した最も重要な決定は、コードを書かずに長い時間を過ごすことでした。
機能の開発に入る前に、Valentina と Jelena は AI エージェントのオンボーディング プログラムに相当するものを構築しました。彼らは、プロジェクト ルールと 16 ステップの開発プロセスを定義する CLAUDE.md を作成しました。彼らは 6 つのエージェント ロールを定義し、それぞれに独自の初期化ファイルを持ちました。
役割に加えて、バックエンド規約、フロントエンド パターン、クリーン コード、JUnit、Bazel、Nx、QA、バージョン管理のための再利用可能な命令セットといったスキルのライブラリも構築しました。エージェントは相互にのみ通信します。確信が持てないときは、人間が立ち止まるのではなく、文書化されたベストエフォートの決定を下します。各機能は独自の git ブランチとワークツリーを取得するため、複数のエージェント チームが衝突することなく同時に実行できます。
彼らはまた、ほとんどの AI コーディングの試みがスキップするもの、つまり実際の仕様をエージェントに提供しました。 PRD と TDD は適切なドキュメントとして記述され、マークダウンに変換され、すべてのエンドポイントを定義する OpenAPI スキーマと UI の一連の Figma 設計資産とともにコンテキストとしてエージェントに渡されました。 UI デザインをそのまま実装することにし、既存のステータス ページを参照して Figma と UX デザインの多くを自動化できるようにしました。
労力の分割は、このプロジェクト全体の重要な発見です。
構築には 1 時間ごとに、準備におよそ 2 時間かかりました。この比率は効率の失敗ではなく、メカニズムに問題があります。エージェントが高速だったのは、まさにコンテキストに次のような特徴があったためです。

すでに曖昧さはすべて取り除かれています。
ハーネスを所定の位置に取り付けると、MVP はすぐに完成しました。プロトタイプ リポジトリのコミット履歴全体は、2026 年 3 月 4 日から 4 月 15 日まで、6 週間、46 件のプル リクエストで実行されます。最初の 3 週間は、ほぼ完全にコンテキスト ファイルとスキル ファイルです。その後、3 月下旬から 4 月上旬の約 2 週間にわたって、グループ CRUD、サービス CRUD、データベース セットアップ、インシデント、メンテナンス ウィンドウ、パブリック ページ、CI、単体テスト、ワークフローのデプロイといった機能が順番に導入されます。
浮かび上がった作業パターン:
小さなタスクは大きなタスクに勝ります。チームはこれを明示的にテストし、ポリシーにしました。大規模な機能リクエストにより、レビューが困難な無秩序なプル リクエストが生成され、多くの場合、破棄する必要がありました。
アーキテクトの役割は維持されます。社内規約をロードした専任のレビューエージェントを配置することで、バックエンドの完全なリファクタリングを含め、人間が目にする前に構造的な問題を発見しました。
人間はキーストロークではなく契約を所有していました。 Valentina と Jelena は、API の形状、データ モデル、ロールアウト シーケンスなど、何が正しいかを決定し、出力が一致していることを確認しました。 AIがコードを書きました。エンジニアたちはシステムを構築しました。
スムーズに進んだふりをしていては、誰にも利益をもたらしません。この取り組みが予想よりも複雑な場所がいくつかありました。
エージェントは作業をやり直します。グループ CRUD は 4 つの個別のプル リクエストに実装され、サービス CRUD は 5 つに実装されました。プロトタイプ リポジトリ内の 46 のプル リクエストのうち、8 つはマージされず、放棄された重複と行き止まりで構成されていました。実際に何が行われたかを人間が追跡することなく、並行エージェントは解決済みの問題を喜んで再解決します。
AI との連携は新しいパラダイムです。最初の MVP は、プロジェクトの実行であると同時に学習演習でもありました。これは、エンジニアが間違いを犯し、それを修正するための 1 回限りの税金でした。と

学んだ教訓により、私たちは他のプロジェクトに取り組み、より自信を持って効率的に作業できるようになりました。このフェーズには全体の作業量の 25% も費やしたため、ステータス ページの開発コストが高くなってしまいました。
下院大会は無料ではありません。最初のバックエンドは Fivetran のビルド システムまたはコード標準と一致せず、再構築する必要がありました。気になるすべての規則は、エージェントが読む場所に書き留める必要があります。この知識の集合体が、最終的にスキル ライブラリーとなりました。
デモは実稼働システムではありません。 AI によって初期実装が加速されましたが、そのプロトタイプを製品品質のソフトウェアに変えるには、多大なエンジニアリング努力が必要でした。
MVP は 4 月末に完成し、7 月下旬に公開されました。そのギャップこそが物語なのです。
プロトタイプを顧客の前に置くものに変えるには、メインのプロダクション エピックで 72 の追跡対象の作業項目が必要となり、さらにプレリリースのバグ修正で 17 の作業項目が必要になりました。長編作品はほとんどありませんでした。それは次のとおりでした:
モノリポジトリへの移行: プロトタイプはスタンドアロン リポジトリにありました。実稼働コードは、共通の Bazel ビルド システムと CI を備えた実稼働 GitHub リポジトリに存在します。
認証と認可: プロトタイプは認証を偽装する可能性があります。製品版が実際に動作するには、それが必要です。ステージングと運用全体にわたる管理パネルに、ロールベースの権限と API トークン管理を備えた Okta を実装しました。
実際のインフラストラクチャ: Cloud SQL Postgres を使用して、継続的なデプロイと構造化されたロギングを備えたステージング環境を構築しました。
他のすべてとの統合: 既存のデータ アグリゲーターとパブリック インシデント トラッカーの Cloud Functions を新しい API に書き込むように再配線し、インシデント管理システムからの自動インシデント作成を有効にし、製品ダッシュボードが表示されるように Java クライアントを構築しました。

ベンダーのバックエンドではなく当社のバックエンドから取得され、すべて機能フラグの背後で配信されます。
サブスクライバーの移行: 26,000 人のサブスクライバーは、誰も通知を失ったり重複したりすることなく移動する必要がありました。これだけでも 40 時間のスクリプト作成作業とその後のデータ修正が必要でした。
監査可能性: インシデント、サービス、グループ、またはサブスクライバーに対するすべての変更は、データベース トリガーによって監査ログにキャプチャされます。
スケール テスト: 広範囲にわたるインシデントにより、50,000 件を超える電子メール イベントが生成される可能性があります。ストレス テストでは、5 分間で 9,000 件の通知が継続されるまでテストしました。
正確性に関するバグの長い尾: インシデントのタイトルで特殊文字が正しく表示されません。 30 日より前のインシデントは稼働時間バーに影響を与えませんでした。メンテナンス時間枠で二重のステータス遷移が発生し、古い集計と新しい集計の間で稼働時間に不一致が発生しました。どれも興味深いものではありませんが、顧客の目に留まるものばかりです。
本番環境の強化は総作業量の 56% であり、要件、エージェント ハーネス、MVP 全体を合わせたものよりも多かったです。この投稿から得られる数字が 1 つあるとすれば、それはその数字です。
私たちは意図的に退屈なロールアウトを実行しました。
ステージング: 公開ページと管理パネルの両方で完全な QA パスを取得しました。
プロダクション、サイレント: 新しいシステム

[切り捨てられた]

## Original Extract

Two engineers used AI coding agents to design, build, and ship a production-ready status page — here's what worked, what didn't, and what we'd do differently.

We built our own status page with AI, replacing a $65k SaaS product | Blog | Fivetran
English English Deutsch Français Español Product
Platform Capabilities
A single platform for all your data movement needs
Accelerate insights with automated transformations
Rigorous, built-in security to move data with peace of mind
Know, protect and scale your data with governed data movement
Easily integrate Fivetran with other tools and systems to optimize workflows at scale
Move warehouse data to business tools to drive impact
Securely move all of your data without compromising performance
Future-proof your stack for agentic AI
Make data accessible without compromising on compliance
Accelerate your business through cloud adoption
Effortlessly connect to your customers’ data
Power your AI innovation with reliable access to centralized and governed data
Drive smarter business decisions with real-time analytics
Blog Customer stories Resource center Documentation Onboarding News Events Professional services Podcast Featured resource See Fivetran in action View interactive demo Trust Support portal Status English
English Deutsch Français Español Contact Sales Log in Start free
Data insights We built our own status page with AI, replacing a $65k SaaS product
Two engineers used AI coding agents to design, build, and ship a production-ready status page — here's what worked, what didn't, and what we'd do differently. In about 4 months, 2 Fivetran engineers — Valentina Mačković and Jelena Kostic — designed, built, and shipped a replacement for a $65,000-a-year SaaS product, using AI coding agents from requirements through production. It now serves Fivetran's public status page and roughly 26,000 notification subscribers. This is what it actually cost, what actually worked, and what we got wrong.
Background: What we were replacing
Fivetran's status page was a heavily customized application that ran on Atlassian Statuspage. It did the job it was designed for, but our use of it had drifted well past the native capabilities of the Statuspage framework. Our status page displays a large number of individually monitored services, and we wanted per-service status and uptime for each of them. This is a shape the hosted product was never really built around. So we layered customization on top: Google Cloud Functions computing uptime from our BigQuery data warehouse, an incident tracker syncing state, and automation wiring it into our incident management process.
The result was a system with real operational costs on both sides of the ledger:
Money: It cost about $65,000 per year , driven substantially by the fact that every notification subscriber is billable. We had ~26,000 of them.
Time : We spent a consistent number of hours managing, troubleshooting, and fixing the status page, adding up to at least $10,000 of engineer time per year.
Performance: The page loaded slowly, and sometimes failed to load on a first attempt. For a page whose entire purpose is to be available when everything else is not, that is a serious failure mode.
Silent failures: We hit cases where updates did not propagate, and we didn’t detect the deficiency.
Rate limits: Our own automation ran into the vendor's API rate limits .
To be fair to Atlassian, none of this is a scandal. It is what happens when you push a general-purpose product a long way past its intended shape, and it is also the ordinary arithmetic of per-subscriber pricing at 26,000 subscribers. In this post, we’ll show how it’s possible to use AI to replace a system that’s outgrown its requirements.
The status page sits in an unusual spot. It is a small application — a handful of read endpoints, an admin panel, an email queue — but it carries production-scale consequences. When Fivetran has a bad day, it is the artifact thousands of customers look at, and a mass notification event can mean 50,000+ emails. Low complexity, high stakes.
That combination is precisely what made it interesting. A year ago, a DIY status page would have been an easy “no”: the engineering cost dwarfed the license. What changed is that AI shifted where that line sits for a certain class of software — well-specified, modest in complexity, expensive to rent. We would soon find out whether this line shifted far enough.
What drove the build-vs-buy decision
Three things, in order of weight:
Control over reliability. A status page that depends on a third party can only ever be as reliable as that third party, and we could not fix the parts that were slow. Owning it meant we could decouple it from both our own core platform and from an external provider.
The pricing model was misaligned with the product. We were paying per subscriber for what is, structurally, a mailing list. Costs scaled with customer growth while the value delivered did not.
We wanted the organizational answer, not just the software. Could 2 engineers using AI agents deliberately take something from a PRD to production? If yes, that is a reusable capability. The status page was a good place to find out because the blast radius of getting it wrong was survivable. We could run it in parallel with the old one until we trusted it. We could also call the experiment a failure without pouring a huge amount of time, money, and effort into it.
We also want to be honest about the boundary of the conclusion. Fivetran's own product is a SaaS product, and we are not arguing that AI makes SaaS obsolete. A status page is far easier to replace than a data movement platform. But the same logic applies at the margin to any vendor: the more your usage is a thin, well-specified slice of a general product, the more the build side of the equation has moved.
The most important decision the team made was to spend a long time not writing code.
Before any feature work, Valentina and Jelena built what amounted to an onboarding program for AI agents. They wrote a CLAUDE.md defining project rules and a 16-step development process. They defined 6 agent roles, each with its own initialization file:
Alongside the roles, they built a library of skills — reusable instruction sets for backend conventions, frontend patterns, clean code, JUnit, Bazel, Nx, QA, and version control. Agents communicate only with each other; when one is unsure, it makes a documented best-effort decision rather than stalling for a human. Each feature gets its own git branch and worktree, so several agent teams can run concurrently without colliding.
They also gave the agents something most AI coding attempts skip: a real specification . The PRD and TDD were written as proper documents, converted to markdown, and handed to the agents as context , along with an OpenAPI schema defining every endpoint and a set of Figma design assets for the UI. We decided to implement the UI design as-is, making it possible to automate much of the Figma and UX design with the existing status page as a reference.
The effort split is the key finding of this whole project:
It took roughly two hours of preparation for every hour of building . That ratio is not a failure of efficiency, but the mechanism. The agents were fast precisely because the context had already removed all ambiguity.
Once the harness was in place, the MVP came together quickly. The prototype repository's entire commit history runs from 4 March to 15 April 2026 : 6 weeks, 46 pull requests. The first 3 weeks are almost entirely context and skill files. Then, over about 2 weeks in late March and early April, the features land in sequence: groups CRUD, services CRUD, database setup, incidents, maintenance windows, the public page, CI, unit tests, deploy workflows.
The working pattern that emerged:
Small tasks beat large ones. The team explicitly tested this and made it policy. Large feature requests produced sprawling, hard-to-review pull requests that often had to be thrown away.
The architect role earns its keep. Having a dedicated reviewing agent with house conventions loaded caught structural problems before humans saw them, including a full refactor of the backend.
The humans owned the contract, not the keystrokes. Valentina and Jelena decided what correct looked like — the API shape, the data model, the rollout sequence — and verified that the output matched. AI wrote the code. The engineers built the system.
We would be doing nobody a favor by pretending this was smooth. There were a few places where the endeavor was more complicated than expected:
Agents redo work. Groups CRUD was implemented in 4 separate pull requests, Services CRUD in 5. Of 46 pull requests in the prototype repo, 8 were never merged , consisting of abandoned duplicates and dead ends. Without a human tracking what was actually done, parallel agents happily re-solve solved problems.
Working with AI is a new paradigm. The initial MVP was as much of a learning exercise as it was a project execution. This was a one-time tax for our engineers to make mistakes and correct them. With lessons learned, we can now take to other projects and work more confidently and efficiently. We spent as much as 25% of the overall effort on this phase, which skewed the cost of developing a status page higher.
House conventions are not free. The first backend did not match Fivetran's build system or code standards and had to be restructured. Every convention you care about has to be written down somewhere the agent will read it. This collection of knowledge is what the skills library eventually became.
A demo is not a production system. While AI accelerated the initial implementation, turning that prototype into production-quality software required significant engineering effort.
The MVP was done at the end of April and went live in late July. That gap is the story.
Turning the prototype into something we would put in front of customers took 72 tracked work items in the main production epic and another 17 in pre-release bug fixing. Almost none of it was feature work. It was:
Moving into the monorepo: The prototype lived in a standalone repository. Production code lives in our production GitHub repo with a common Bazel build system and CI.
Authentication and authorization: The prototype could fake authentication; a production version needs it to work for real. We implemented Okta for the admin panel across staging and production, with role-based permissions and API token management.
Real infrastructure: Using Cloud SQL Postgres, we built a staging environment with continuous deployment and structured logging.
Integration with everything else: We rewired our existing data-aggregator and public-incident-tracker Cloud Functions to write to the new API, enabled automatic incident creation from our incident management system, built a Java client so the product dashboard reads from our backend instead of the vendor's, and delivered it all behind a feature flag.
The subscriber migration: 26,000 subscribers had to move without anyone losing their notifications or getting duplicates. This alone was a 40-hour scripting effort plus follow-on data corrections.
Auditability: Every change to an incident, service, group, or subscriber is captured by a database trigger into an audit log.
Scale testing: A widespread incident can generate 50,000+ email events. We stress-tested to a sustained 9,000 notifications in 5 minutes.
A long tail of correctness bugs: Special characters rendered incorrectly in incident titles. Incidents older than 30 days were not affecting the uptime bar. There were double status transitions on maintenance windows, and uptime discrepancies between the old and new aggregation. None of these are interesting, but all of them would have been noticed by customers.
Production hardening was 56% of the total effort — more than the requirements, the agent harness, and the entire MVP combined. If there is one number to take from this post, it is that one.
We ran a deliberately boring rollout.
Staging: We took a full QA pass on both the public page and the admin panel.
Production, silent: The new sy

[truncated]
