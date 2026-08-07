---
source: "https://webhosting.today/2026/08/07/an-ai-model-found-the-wordpress-flaw-in-ten-hours-for-25-attackers-weaponized-the-patch-in-ninety-minutes/"
hn_url: "https://news.ycombinator.com/item?id=49210426"
title: "An AI Model Found the WordPress Flaw in Ten Hours for $25"
article_title: "WordPress: 10 Hours to Find, 90 Min to Exploit - webhosting.today"
author: "speckx"
captured_at: "2026-08-07T14:05:59Z"
capture_tool: "hn-digest"
hn_id: 49210426
score: 2
comments: 0
posted_at: "2026-08-07T13:53:18Z"
tags:
  - hacker-news
  - translated
---

# An AI Model Found the WordPress Flaw in Ten Hours for $25

- HN: [49210426](https://news.ycombinator.com/item?id=49210426)
- Source: [webhosting.today](https://webhosting.today/2026/08/07/an-ai-model-found-the-wordpress-flaw-in-ten-hours-for-25-attackers-weaponized-the-patch-in-ninety-minutes/)
- Score: 2
- Comments: 0
- Posted: 2026-08-07T13:53:18Z

## Translation

タイトル: AI モデルが 25 ドルで 10 時間で WordPress の欠陥を発見
記事のタイトル: WordPress: 見つけるのに 10 時間、悪用するのに 90 分 - webhosting.today
説明: AI モデルは 25 ドルで 10 時間で WordPress の欠陥を発見しました。攻撃者は90分以内にパッチを武器化しました。ホストにとって 7.0.3 が意味するもの。

記事本文:
WordPress: 見つけるのに 10 時間、悪用するのに 90 分 - webhosting.today
コンテンツにスキップ
ホーム
Web ホスティング M&A 求人バズ ポッドキャスト
もっと見る ▾
専門家 イベント 2026 私たちについて スポンサー
市場洞察 →
検索
ホーム
ウェブホスティング M&A エキスパート 求人情報 バズ ポッドキャスト イベント 2026 私たちについて スポンサー
市場洞察 →
© webhosting.today
すべての著作権は留保されています。
AI モデルが 25 ドルで 10 時間で WordPress の欠陥を発見しました。攻撃者は90分でパッチを兵器化しました。
WordPress は 8 月 6 日に 7.0.3 をリリースしました。これはコアチームが直ちにインストールすることを推奨するセキュリティ リリースです。その見出しのエントリは、PHP コードの実行につながる可能性がある、ログイン画面上の認証前に反映されたクロスサイト スクリプティングの欠陥であり、pwn.ai のチームによって報告され、CVE-2026-64638 として追跡されています。
このようなリリースはここ 3 週間で 2 回目であり、どちらのリリースよりも最初のリリースを囲む 2 つの数字の方が重要です。 Patchstack によると、7 月 17 日に 7.0.2 で修正されたクリティカル チェーンは、バッチ API を 1 行ずつ読んだ研究者によって発見されたものではありませんでした。Searchlight Cyber​​ は、WordPress コアで OpenAI モデルを指定し、リモート コード実行にチェーンされた有効な事前認証 SQL インジェクションを 10 時間、約 25 ドルのコストで実現しました。 7.0.2 が出荷された後、最初の悪用の試みは約 90 分後、つまり修正が WordPress トランクにコミットされてから 3 時間後に Patchstack のセンサーに到達しました。脆弱性ライフサイクルの両端が一気に圧縮され、ホスティング企業がその間のスペースを占めています。
発見: Patchstack によると、モデルはおよそ 25 ドルで 10 時間で動作チェーンを生成しました。 WordPress.org は、このレポートについて Assetnote および Searchlight Cyber の Adam Kues 氏の功績を認めています。
悪用: 最初の試行は 7.0.2 がリリースされてから 90 分後に成功しました。これは、修正がトランクにコミットされてから 3 時間後でした。ずっと

パッチスタックは数日間、1,500 を超えるアドレスからの 65,000 を超える試行をブロックしました
軽減策: パッチスタックでは、ルールがリクエスト URL で一致し、WordPress が POST 本文から同じルートを受け入れるため、最初の数時間で公開されたほぼすべての WAF ルール (独自のものを含む) がバイパス可能であることが判明しました。
ボリューム: Patchstack が発表したコア セキュリティ リードの John Blackbourn 氏の数字によると、HackerOne を通じた WordPress セキュリティ レポートは 9 年間毎月数十件のペースで実行され、7 月には 450 件に達しました。
7 月 17 日のタイムラインは、修正が WordPress トランクにコミットされた瞬間から、リリースの 1.5 時間前に測定されました。出典: パッチスタック テレメトリ。
10時間と25ドル
Patchstack はモデルに OpenAI の GPT-5.6 Sol Ultra という名前を付け、Searchlight Cyber によって WordPress コアに対して実行されます。 WordPress.org 自身のリリースノートでは、結果として得られたレポートは Assetnote と Searchlight Cyber​​ の Adam Kues の功績として認められており、方法や価格ではなく誰が欠陥を発見したかを裏付けています。
8月のリリースでも同じことが指摘されている。その修正のうち 2 つは、AI を中心に構築された企業によるものです。自律的な侵入テストに取り組む pwn.ai のログイン画面の欠陥と、Anthropic への作成者レベルの CSS インジェクションです。 Patchstack がこの変化を構想する中、AI 支援研究とは、かつては人間の研究者と非常に高速なアシスタントを意味していました。現在、モデルは欠陥を見つけてそれを悪用しており、研究者の仕事は報告書を提出することです。
摂取量は生産量を示します。 Patchstack のアドバイザリで公開されたコア セキュリティ リードの John Blackbourn 氏のデータによると、HackerOne を通じて WordPress に届いたレポートは、WordPress.org のコア インフラストラクチャと公式プラグインをカバーしており、月に数十件のペースで 9 年間実行され、この春から増加し始め、7 月には 450 件に達しました。

それ自体は悪いニュースではありません。Patchstack は慎重にそう言っています。オープンソースのコードベースに注目が集まり、より速く動くことは長期的な利益です。結果はメリットよりもタイミングが重要です。モデルが 10 時間以内に起動状態から実際にリモートでコードを実行できるようになるのであれば、防御者が情報を開示してから悪用されるまでに数日かかるという想定は、もはや世界を説明できなくなります。
キャンペーン開始まであと 90 分
もう一方の端も圧縮されました。 7.0.2 の問題は 2 つの部分から構成されていました。1 つは author_exclude パラメーターを介して到達可能な WP_Query での SQL インジェクション (CVE-2026-60137、WordPress 6.8 ～ 7.0.1 に影響) と、REST API バッチ エンドポイントでのルートとハンドラーの混乱 (CVE-2026-63030、6.9 ～ 7.0.1 に影響) です。どちらも単独での買収ではありません。連鎖的に、認証されていないリクエストが通常は無害化されるスキーマを通過してインジェクションを密輸させ、公開されたエクスプロイトがそこから新しい管理者に渡ってコードを実行します。 WordPress.org は、自動更新システムによる強制更新を可能にするのに十分な重大度であると考えました。
このリリースの周りの 2 つの図は、異なる出発点から同じ瞬間を描写しています。この修正は、7.0.2 がタグ付けされる約 1.5 時間前に WordPress トランクに配信され、コア開発を監視している人は誰でもそれを比較して、何が変更されたかを正確に確認できるため、Patchstack はそのコミットを有効な開示として扱います。最初の試行はリリースから 90 分後、コミットから 3 時間後に到着しました。
Patchstack は顧客のサイト内から余波を監視しましたが、そこで見たものは 1 人のオペレーターというよりも、陸地へのラッシュのように見えました。1,500 を超える一意のアドレスからの 65,000 件を超える試行がブロックされ、そのほとんどが VPS およびクラウド範囲からのもので、最もビジーな単一ネットワークがボリュームのわずか 6.45 パーセントを占めていました。ソム

97% が REST バッチ エンドポイントをターゲットにしており、約 4 分の 3 が author_exclude パラメータでの注入試行を行っていました。
そのほとんどはスキャンでした。 Patchstack のアカウントでは、わずか 3 つのアドレスからのリクエストの小さなクラスターが完全なチェーンを運び、管理者の作成に直接進みましたが、既製のエクスプロイト ツールが wp2shell という名前で公に流通していました。これらの数値には注意点が 1 つあります。これらの数値は、侵害が完了したものではなく、脆弱なバージョンを実行していたサイトでブロックされた試行をカウントしています。
モデルが人工的なゼロデイを発見。自己ホスト型ユーザーのみがパッチを適用する必要がありました
7 月に導入された WAF ルールのホストには共通の盲点がありました
保存期間が最も長い発見は、エクスプロイトに関するものではありません。緩和策についてです。
7 月の公開から数時間以内に公開されたほぼすべての WAF ルールは、リクエスト URL 内のバッチ/v1 を検索していました。これは、公開されているすべての概念実証がエンドポイントに到達した方法だからです。 Patchstack は、WordPress がrest_route をパブリック クエリ変数として登録し、それらの変数をクエリ文字列から読み取る前に POST 本文から読み取ることを発見しました。本文にルートとペイロードを含むプレーンな POST をサイトのルートに送信すると、WordPress はそのリクエストをバッチ エンドポイントにルーティングしますが、URL はホームページへの通常のリクエストのように見えます。 URL のみを検査するルールにより、URL が渡されます。
このギャップは 1 つの製品に特有のものではなく業界全体に広がっていたため、Patchstack はこれを調整された開示として処理し、影響を受けるベンダーと WordPress セキュリティ チームに通知し、公開前に主要なホストと DNS レベルの WAF プロバイダーにそのギャップを伝えました。 Patchstack は、独自の最初のルールにも同じギャップがあったと明言し、7 月から実環境でボディベースのフォームを使用する試みが確認されたと報告しています

21. 独自のルールを維持している人にとって、これは実行可能な部分です。概念実証でたまたまルートが配置された場所だけでなく、WordPress が受け入れる場所であればどこでもルートを一致させます。
8 月リリースの実際の内容
新しいリリースはパニックになるよりむしろ比例するに値する。このログイン画面の欠陥は認証なしでアクセスできるため、7 月の連鎖のように聞こえますが、反映された XSS は細工されたリンクをクリックした人のブラウザ内でのみ起動されます。 Patchstack が言うように、これはドライブバイではなくターゲットを絞った送信となり、クリックした人が管理者でない場合、ペイロードは何も役に立ちません。管理者がクリックすると、実際にコードが実行されるパスになります。
他の 2 つのエントリーは、見出しよりもホスティング会社の注目に値します。 Aikido Security によって報告された、ユーザー登録が有効になっているマルチサイト ネットワーク上で権限昇格があり、ユーザーが新しいサイトを作成できるようになります。また、URL 検証にはサーバー側のリクエスト フォージェリがあり、リンク ローカル範囲 (クラウド インスタンスのメタデータ サービスが存在するアドレス空間) へのリクエストを許可します。共有またはクラウドでホストされているフリートでは、2 番目の問題はアプリケーションの問題であると同時にプラットフォームの問題でもあります。パッチスタックでは、このリリースには合計 12 件の修正が含まれています。バックポートは現在も古いブランチ (現在は 4.7 まで) にロールアウトされており、WordPress 7.1 RC2 にはすでに該当する修正が含まれています。
ホスティング会社にとって何が変わるのか
リリースノートを読むことはもはや防御策ではありません。パブリック コミットとライブ ペイロードの間の間隔が 3 時間の場合、アナウンスに気づく人間に依存するプロセスはすでに失われています。そのタイムスケールで動作する唯一のメカニズムは自動更新と事前展開された仮想パッチであり、ホストが両方を制御します。
WordPress.org は、

重大度が必要な場合、コア更新自体をプッシュします。 7.0.2 では、サイト所有者を待つのではなく自動更新システムを使用して、まさにそれを行いました。安定性を理由にコアの自動更新を抑制または遅延しているホストは、1 年前よりも大きなリスクを抱えているため、どの計画が自動的に更新されるかを大声で言えるようにする必要があります。
ルールの品質は、仮定するのではなく、監査する価値があります。 7 月のバイパスは業界全体のルールに影響を与えました。仮想パッチを購入するホストは、ルールが URL または動作にどのように固定されているか、バイパスが表面化してからルールがどのくらい早く改訂されたかをベンダーに尋ねる必要があります。
メンテナンスウィンドウの時間単位が間違っています。発見に 10 時間と 25 ドルの費用がかかり、武器化に 90 分かかるとすると、毎月のパッチ サイクルとそれが対処すべき脅威との間のギャップは桁違いに大きくなります。
3 週間で 2 つのコア リリースが発生するのは、スケジュールの偶然であることが判明する可能性があります。彼らの背後にある経済学はそうではありません。
パネル、プロバイダー、舞台裏で実行されているすべてのものなど、執筆を通じて Web ホスティング業界を探ります。
WordPress 7.0.3 リリース - WordPress.org ニュース
WordPress 7.0.2 リリース - WordPress.org ニュース
WordPress 7.0.3 リリース: 12 件の脆弱性が見つかり修正されました - パッチスタック
90分: WordPressコアRCEを攻撃者が武器化する様子を観察 - パッチスタック
WordPressコアの認証されていないSQLインジェクションが7.0.2で修正されました - パッチスタック
10時間と25ドル
キャンペーン開始まであと 90 分
7 月に導入された WAF ルールのホストには共通の盲点がありました
8 月リリースの実際の内容
ホスティング会社にとって何が変わるのか
AI モデルが 25 ドルで 10 時間で WordPress の欠陥を発見しました。攻撃者は90分でパッチを兵器化しました。
数週間で 10,000 のサイトにフラグが付けられる: npm スキャンのパッチスタックとホスティンガー

デフォルトで
ドメインの運営には年間約 1 ドルの費用がかかります。 Eight レジストリの料金は次のとおりです。
ホスティングの次の 6 か月はすでに予定されています: 準備する 5 つの日程
ホーム
ポッドキャスト
私たちについて
利用規約
プライバシーポリシー
© 2024-2026 webhosting.today
すべての著作権は留保されています。

## Original Extract

An AI model found the WordPress flaw in 10 hours for $25. Attackers weaponized the patch in 90 minutes. What 7.0.3 means for hosts.

WordPress: 10 Hours to Find, 90 Min to Exploit - webhosting.today
Skip to content
Home
Web Hosting M&A Jobs Buzz Podcast
More ▾
Experts Events 2026 About us Sponsor us
Market Insights →
Search
Home
Web Hosting M&A Experts Jobs Buzz Podcast Events 2026 About us Sponsor us
Market Insights →
© webhosting.today
All Rights Reserved.
An AI Model Found the WordPress Flaw in Ten Hours for $25. Attackers Weaponized the Patch in Ninety Minutes.
WordPress released 7.0.3 on August 6, a security release the core team recommends installing immediately. Its headline entry is a pre-auth reflected cross-site scripting flaw on the login screen that can lead to PHP code execution, reported by the team at pwn.ai and tracked as CVE-2026-64638.
It is the second such release in three weeks, and the two numbers that bracket the first one matter more than either release. According to Patchstack, the critical chain fixed in 7.0.2 on July 17 was not found by a researcher reading the batch API line by line: Searchlight Cyber pointed an OpenAI model at WordPress core and had a working pre-auth SQL injection chained to remote code execution in ten hours, at a cost of about $25 . After 7.0.2 shipped, the first exploitation attempts hit Patchstack’s sensors roughly 90 minutes later , three hours after the fix was committed to WordPress trunk. Both ends of the vulnerability lifecycle have compressed at once, and hosting companies occupy the space in between.
Discovery: per Patchstack, a model produced the working chain in ten hours for roughly $25 ; WordPress.org credits Adam Kues at Assetnote and Searchlight Cyber for the report
Exploitation: the first attempts landed 90 minutes after 7.0.2 was released , which was three hours after the fix was committed to trunk ; over the following days Patchstack blocked more than 65,000 attempts from over 1,500 addresses
Mitigation: Patchstack found that nearly every WAF rule published in the first hours was bypassable , including its own, because the rules matched on the request URL and WordPress accepts the same route from the POST body
Volume: WordPress security reports through HackerOne ran in the dozens per month for nine years and reached 450 in July , according to figures from core security lead John Blackbourn published by Patchstack
The July 17 timeline, measured from the moment the fix was committed to WordPress trunk, 1.5 hours before the release went out. Source: Patchstack telemetry.
Ten Hours and Twenty-Five Dollars
Patchstack names the model: OpenAI’s GPT-5.6 Sol Ultra , run against WordPress core by Searchlight Cyber. WordPress.org’s own release notes credit the resulting report to Adam Kues at Assetnote and Searchlight Cyber, which corroborates who found the flaw, though not the method or the price.
The August release points the same way. Two of its fixes are credited to companies built around AI: the login screen flaw to pwn.ai , which works on autonomous penetration testing, and an Author-level CSS injection to Anthropic . As Patchstack frames the shift, AI-assisted research used to mean a human researcher with a very fast assistant; now the model finds the flaw and carries it to exploitation, and the researcher’s job is to file the report.
The intake numbers show the volume that produces. Reports reaching WordPress through HackerOne, covering core, WordPress.org infrastructure and official plugins, ran at dozens per month for nine years, began climbing this spring and hit 450 in July , according to data from core security lead John Blackbourn published in Patchstack’s advisory.
None of that is bad news in itself, and Patchstack is careful to say so: more eyes on an open source codebase, moving faster, is a long-term gain. The consequence is about timing rather than merit. If a model can go from a standing start to working remote code execution in ten hours, the assumption that defenders have days between disclosure and exploitation no longer describes the world.
Ninety Minutes to a Live Campaign
The other end compressed too. The 7.0.2 problem was a two-part chain: a SQL injection in WP_Query reachable through the author_exclude parameter (CVE-2026-60137, affecting WordPress 6.8 through 7.0.1) plus a route and handler confusion in the REST API batch endpoint (CVE-2026-63030, affecting 6.9 through 7.0.1). Neither is a takeover alone. Chained, they let an unauthenticated request smuggle the injection past the schema that would normally sanitize it, and the published exploit walks from there to a new administrator and code execution. WordPress.org considered the severity high enough to enable forced updates through the auto-update system.
The two figures around that release describe the same moment from different starting points. The fix landed in WordPress trunk about 1.5 hours before 7.0.2 was tagged , and Patchstack treats that commit as the effective disclosure, because anyone watching core development can diff it and see exactly what changed. The first attempts arrived 90 minutes after the release, three hours after the commit .
Patchstack watched the aftermath from inside its customers’ sites, and what it saw looked less like one operator than a land rush: more than 65,000 blocked attempts from over 1,500 unique addresses , mostly from VPS and cloud ranges, with the busiest single network accounting for only 6.45 percent of the volume. Some 97 percent targeted the REST batch endpoint , and around three quarters carried an injection attempt in the author_exclude parameter.
Most of that was scanning. In Patchstack’s account, a small cluster of requests from just three addresses carried the complete chain and went straight for administrator creation, while a ready-made exploit tool circulated publicly under the name wp2shell. One caveat travels with all of these numbers: they count blocked attempts on sites that were running vulnerable versions, not completed compromises.
A Model Found the Artifactory Zero-Day; Only Self-Hosted Users Had to Patch
The WAF Rules Hosts Deployed in July Had a Common Blind Spot
The finding with the longest shelf life is not about the exploit at all. It is about the mitigation.
Nearly every WAF rule published in the hours after the July disclosure looked for batch/v1 in the request URL, because that is how every public proof of concept reached the endpoint. Patchstack found that WordPress registers rest_route as a public query variable and reads those variables from the POST body before it reads them from the query string. Send a plain POST to the site root with the route and the payload in the body, and WordPress routes the request to the batch endpoint while the URL still looks like an ordinary request to the homepage. A rule that inspects only the URL waves it through.
Because the gap was industry-wide rather than specific to one product, Patchstack handled it as a coordinated disclosure, notifying the affected vendors and the WordPress security team, which passed it on to major hosts and DNS-level WAF providers before publication. Patchstack states plainly that its own first rules had the same gap, and reports seeing attempts using the body-based form in the wild from July 21. For anyone maintaining their own rules, that is the actionable part: match the route wherever WordPress will accept it, not only where the proof of concept happened to put it.
What Is Actually in the August Release
The new release deserves proportion rather than panic. The login screen flaw is reachable without authentication, which is what makes it sound like the July chain, but a reflected XSS fires only in the browser of whoever clicks the crafted link. As Patchstack puts it, that makes it a targeted send rather than a drive-by, and if the person clicking is not an administrator the payload achieves nothing useful. If an administrator does click it, the path to code execution is real.
Two other entries deserve a hosting company’s attention more than the headline. There is a privilege escalation on multisite networks with user registration enabled , which lets a user create a new site, reported by Aikido Security. And there is a server-side request forgery in URL validation that allows requests to link-local ranges , the address space where cloud instance metadata services live. On shared or cloud-hosted fleets, that second one is a platform problem as much as an application problem. Patchstack counts 12 fixes in the release in total. Backports are still rolling out to older branches, currently as far back as 4.7, and WordPress 7.1 RC2 already contains the applicable fixes.
What This Changes for Hosting Companies
Reading the release notes is no longer a defense. When the interval between a public commit and live payloads is three hours, any process that depends on a human noticing an announcement has already lost. The only mechanisms that operate on that timescale are automatic updates and pre-deployed virtual patches, and hosts control both.
WordPress.org will push core updates itself when severity demands it. It did exactly that with 7.0.2, using the auto-update system rather than waiting for site owners. Hosts that suppress or delay core auto-updates for stability reasons now carry a larger risk than they did a year ago, and should be able to say out loud which of their plans update automatically.
Rule quality is worth auditing, not assuming. The July bypass affected rules across the industry. A host buying virtual patching should be asking vendors how a rule is anchored, on the URL or on the behavior, and how quickly a rule was revised once the bypass surfaced.
Maintenance windows are the wrong unit of time. If discovery costs ten hours and $25, and weaponization takes ninety minutes, the gap between a monthly patch cycle and the threat it is meant to answer is now measured in orders of magnitude.
Two core releases in three weeks may turn out to be a coincidence of scheduling. The economics behind them will not.
Exploring the web hosting industry through writing - panels, providers, and everything that runs behind the scenes.
WordPress 7.0.3 release - WordPress.org News
WordPress 7.0.2 release - WordPress.org News
WordPress 7.0.3 released: 12 vulnerabilities found and fixed - Patchstack
Ninety minutes: watching attackers weaponize the WordPress core RCE - Patchstack
Unauthenticated SQL injection in WordPress core fixed in 7.0.2 - Patchstack
Ten Hours and Twenty-Five Dollars
Ninety Minutes to a Live Campaign
The WAF Rules Hosts Deployed in July Had a Common Blind Spot
What Is Actually in the August Release
What This Changes for Hosting Companies
An AI Model Found the WordPress Flaw in Ten Hours for $25. Attackers Weaponized the Patch in Ninety Minutes.
10,000 Sites Flagged in Weeks: Patchstack and Hostinger on npm Scanning by Default
Running a Domain Costs About a Dollar a Year. Here Is What Eight Registries Charge for It.
Hosting’s Next Six Months Are Already Scheduled: Five Dates to Prepare For
Home
Podcast
About us
Terms of service
Privacy Policy
© 2024-2026 webhosting.today
All Rights Reserved.
