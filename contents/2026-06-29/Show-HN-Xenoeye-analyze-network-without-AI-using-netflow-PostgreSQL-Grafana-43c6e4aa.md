---
source: "https://github.com/vmxdev/xenoeye"
hn_url: "https://news.ycombinator.com/item?id=48725156"
title: "Show HN: Xenoeye – analyze network without AI using netflow, PostgreSQL, Grafana"
article_title: "GitHub - vmxdev/xenoeye: Lightweight Netflow/IPFIX/sFlow collector and analyzer · GitHub"
author: "vmxdev"
captured_at: "2026-06-29T21:30:44Z"
capture_tool: "hn-digest"
hn_id: 48725156
score: 1
comments: 0
posted_at: "2026-06-29T21:00:01Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Xenoeye – analyze network without AI using netflow, PostgreSQL, Grafana

- HN: [48725156](https://news.ycombinator.com/item?id=48725156)
- Source: [github.com](https://github.com/vmxdev/xenoeye)
- Score: 1
- Comments: 0
- Posted: 2026-06-29T21:00:01Z

## Translation

タイトル: Show HN: Xenoeye – netflow、PostgreSQL、Grafana を使用して AI を使用せずにネットワークを分析する
記事のタイトル: GitHub - vmxdev/xenoeye: 軽量 Netflow/IPFIX/sFlow コレクターおよびアナライザー · GitHub
説明: 軽量の Netflow/IPFIX/sFlow コレクターおよびアナライザー - vmxdev/xenoeye
HN テキスト: タイトルが少し切れていて申し訳ありません。それは「netflow ファミリーのプロトコル、PostgreSQL または ClickHouse、Grafana、およびいくつかのスクリプトを使用した、AI を使用しないネットワーク トラフィックの分析と監視」であるはずでした。 2026 年に HN で AI フリー ソフトウェアを発表するのは少しおこがましいかもしれません。
しかし、ネットフロー アナライザーを手動で構築するのは、同様におこがましいことです。最近では xFlow アナライザーがたくさん出てきていますが、私はいつもこのことを思い出します。しかし、別のアプローチの余地は常にあると思います。結局のところ、ソフトウェアはそうやって進化するのですね。では、xenoeye は一般的な (少なくとも一般的なオープンソースの) アナライザーとどう違うのでしょうか?・アナライザーには「オブジェクトの監視」という機能があります。何らかの理由で、オープンソース アナライザーではこの機能がほとんど使用されませんが、商用アナライザーではこの機能が使用されます。
監視オブジェクトには、サブネット、自律システム、地理オブジェクト (地理および AS 上のデータは外部データベースから取得されます)、アプリケーション トラフィック (プロトコル、TCP/UDP ポートなど)、VLAN などが考えられます。
フロー レコードのほとんどすべてを監視オブジェクトのフィルターとして使用できます。
もちろん、オブジェクト フィルターは複合にすることもできます。従来の演算である AND、OR、NOT がサポートされています。アナライザーには、各フローをオブジェクトと照合する小さな仮想マシンが含まれています。 - すべてのフローを保存するわけではありません。少なくとも今のところは。奇妙に思えるかもしれませんが、これは特に大規模なネットワークにとって重要な機能です。
監視対象のオブジェクトに関する集約データを保存します。ユーザーは何を保存するかを選択します。それは、単なる入出力、トップトーカー、トッププロトコルなどです。
t

データを集計する ime もユーザーによって指定されます。
集計はアナライザー内で行われます。高速トライベースのインメモリ DB を使用します。
このため、アナライザーはフローを非常に高速に処理し (vCPU あたり数十万 FPS)、測定された量の情報をデータベースにエクスポートできます。
バニラの PostgreSQL でも簡単に使用できます。または、圧縮を使用した ClickHouse。
アナライザーはリソースをあまり消費しません。小規模なネットワーク トラフィックは、ローエンドのハードウェアまたは少量のメモリを備えた VM で処理できます。
あるいは、クラスターを構築せずに、単一サーバーで大規模なネットワーク トラフィックを処理することもできます。私は、単一の仮想マシン上にマルチテラビットのトラフィックと数百の MO を備えたインストールを知っています (もちろん、ルーターでは高いサンプリング レートを備えています)。 - 移動平均を使用して、超過するトラフィックのしきい値を監視できます。
つまり、超過が検出されるとすぐに、外部スクリプトが同じ秒 (実際にはさらに速く) 起動されます。
この機能は通常、ボリューム DoS/DDoS 攻撃を検出するために使用されます。
スクリプトは BGP Blackhole または BGP Flowspec をアナウンスし、メッセンジャー経由でユーザーに通知します。 - 独自の視覚化ユーティリティはありません。グラファナを使用しています。 Grafana はそのまま PostgreSQL で動作しますが、一部の複雑な時系列グラフでは SQL クエリを多少調整する必要があります。
これは物議を醸す決定ですが、ユーザー (そして私たち自身) は今のところ我慢しています。残りの部分をドキュメントに記載してみました。はい、私がこのプロジェクトを HN で発表しようとしたのはこれが初めてではありません。そして、私は何の幻想も抱いていません。何らかの理由で、ハッカーはこの種のソフトウェアをあまり好まないのです。
おそらく誰もが、ネットフロー アナライザーの製造は退屈すぎる問題であり、議論することは何もないと考えているでしょう。でも、もし興味がある人がいたら、それは素晴らしいことです

フィードバックを得る。これまでと違うことは何ですか、またその理由は何ですか?
他では見つからないお気に入りのアナライザーのどこが一番気に入っていますか?この投稿をどのようにしてご覧になりましたか?これは AI や Rust 関連のものではありません

記事本文:
GitHub - vmxdev/xenoeye: 軽量 Netflow/IPFIX/sFlow コレクターおよびアナライザー · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
vmxdev
/
ゼノアイ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスターブランチ Tags Go

ファイルへ コード 「その他のアクション」メニューを開く フォルダーとファイル
349 コミット 349 コミット .github/ workflows .github/ workflows aajson @ 7abedc5 aajson @ 7abedc5 docs-img docs-img lxc lxc m4 m4 スクリプト スクリプト テスト テスト tkvdb @ 61efc75 tkvdb @ 61efc75 .dockerignore .dockerignore .gitattributes .gitattributes .gitignore .gitignore .gitmodules .gitmodules .travis.yml .travis.yml CONFIG.md CONFIG.md CONFIG.ru.md CONFIG.ru.md Dockerfile Dockerfile EXTRA.md EXTRA.md EXTRA.ru.md EXTRA.ru.md INTERNALS.md INTERNALS.md INTERNALS.ru.md INTERNALS.ru.md ライセンス ライセンス Makefile.am Makefile.am README.md README.md README.ru.md README.ru.md STEP-BY-STEP.md STEP-BY-STEP.md STEP-BY-STEP.ru.md STEP-BY-STEP.ru.md分類.c分類.cconfigure.acconfigure.acデバイス.cデバイス.cデバイス.confデバイス.confデバイス.hデバイス.h docker-compose-dev.yml docker-compose-dev.yml docker-compose.yml docker-compose.yml filter-ag.def filter-ag.def filter-lexer.c filter-lexer.c filter-parser-funcs.c filter-parser-funcs.c filter-parser.c filter-parser.c filter.c filter.c filter.def filter.def filter.h filter.h flow-debug.c flow-debug.c flow-debug.h flow-debug.h flow-info.h flow-info.h geoip.c geoip.c geoip.h geoip.h ip-btrie.h ip-btrie.h iplist.c iplist.c iplist.h iplist.h monit-objects-common.h monit-objects-common.h monit-objects-fwm.c monit-objects-fwm.c monit-objects-mavg-act.c monit-objects-mavg-act.c monit-objects-mavg-dump.c monit-objects-mavg-dump.c monit-objects-mavg-limfile.c monit-objects-mavg-limfile.c monit-objects-mavg-under.c monit-objects-mavg-under.c monit-objects-mavg.c monit-objects-mavg.c monit-objects.c monit-objects.c monit-objects.h monit-objects.h netflow-templates.c netflow-templates.c netflow-templates.h netflow-templates.h netflow.c netflow.c netflow.def netflow.def netflow。

h netflow.h pcapture.c pcapture.c rawparse.h rawparse.h scapture.c scapture.c sflow-impl.h sflow-impl.h sflow.c sflow.c sflow.h sflow.h utils-data.inc utils-data.inc utils.c utils.c utils.h utils.h xe-debug.h xe-debug.h xe-dns.h xe-dns.h xe-sni.h xe-sni。
[切り捨てられた]
軽量の Netflow/IPFIX/sFlow コレクターおよびアナライザー
README.ru.md - документация на русском
ドキュメントはほとんど Google 翻訳ツールを使用して自動的に翻訳されているため、何かおかしな点を見つけた場合は、お気軽にお知らせください。
IPネットワーク、個々のIPアドレスまたはサービスのトラフィックを監視します
トラフィックの急増またはしきい値を下回るトラフィックの低下に迅速に対応します
Netflow/IPFIX/sFlow からのデータを使用して、トラフィック パターンとネットワーク パケットの分布を監視します
このコレクターは、さまざまなレポートを必要とするさまざまなユーザー グループが存在する中規模および大規模なネットワーク向けに開発されました。この目的のために「監視オブジェクト」が使用されます。監視オブジェクトは、ネットワーク、ネットワークのセット、自律システム、地理オブジェクト、または Netflow/IPFIX/sFlow から抽出できる任意のネットワーク トラフィックです。
コレクターを使用すると、Grafana でさまざまなレポートを生成したり、グラフやダッシュボードを作成したり、トラフィック速度がしきい値を超えたりしきい値を下回ったときにいくつかのアクションを実行したりできます。
ネットワークを監視するためにコレクターを使用します。 Netflow v9 と IPFIX を使用しているため、コレクターはそれらをサポートしています。
Netflow v5 および sFlow もサポートされています。
ドキュメントには、簡単なレポートの作成例が含まれています。より複雑なものを構築するには、少なくとも SQL の基本的な知識が必要です。
コレクターはテキスト構成ファイルを使用します。これにより、単純な構成を手動で作成でき、多数のオブジェクトを含む複雑な構成の場合は、スクリプトを使用して構成を生成できます。
コレクターは 2 つの方法でデータを処理します。

期間 (レポートやグラフを作成するための固定サイズの時間枠) にわたってゲート処理し、移動平均を使用してスパイクに迅速に対応します。
どちらの方法も個別に使用することも、組み合わせて使用​​することもできます。たとえば、移動平均によってしきい値の超過が検出された場合、カスタム スクリプトを実行して、すぐに拡張統計収集を有効にすることができます。
移動平均を使用してボリューム DoS/DDoS 攻撃を検出します。しきい値に達すると、BGP アナウンス (FlowSpec フィルタリング、レート制限、クリーニング サーバーまたはブラックホールへのリダイレクト) が作成され、ユーザーはメッセンジャーで通知を受け取ります。
コレクターはリソースをあまり要求しません。 4 GB のメモリを備えた Orange Pi (Raspberry Pi に類似) でもデータを処理し、レポートを作成できます。小規模なネットワークでは、1 つの CPU と 1 GB の RAM を備えた VM で実行できます。
コレクターは、64 ビット Linux (x64、AArch64、および Elbrus) でのみテストされています。
時系列データのストレージとして PostgreSQL を使用します。選択した Netflow フィールドごとに集計されたデータがそこにエクスポートされます。コレクターは、すべてのデータを DBMS にエクスポートできるわけではありません。上位 N エンティティのみを集約してエクスポートし、残りを 1 行に集約できます。これは、大規模な監視オブジェクトに便利な機能です。DBMS に書き込まれるデータの量を調整し、より安価で低速のディスクを使用できます。
PostgreSQL に加えて、コレクターは ClickHouse にデータを保存するための実験的なサポートを備えています。
Netflow/IPFIX フィールドの基本セットはそのままサポートされていますが、必要なほぼすべてのフィールドを追加できます。
このプロジェクトには非常に寛大な ISC ライセンスが付与されています。商用版または半商用版を作成する予定はありません。これは、プロジェクトの将来については予測できないことを意味します。しかしその一方で:
隠された制限や人為的な制限はありません
ユーザーは通常、少なくとも次のことに興味を持っています。

パフォーマンスを見積もるため、いくつかのテストを行いました。さまざまなルーターからの実際の Netflow トラフィックを pcap ファイルに記録し、tcpreplay を使用してループバック インターフェイス上でさまざまな速度で再生しました。
テストは i3-2120 CPU @ 3.30GHz で実行されました。
非常に大まかに言うと、次の数字が信頼できます。
デバッグ モードで各フローの内容をファイルに出力すると、1 つの CPU あたり 1 秒あたり約 100K のフローが発生したことがわかりました。
実稼働モードに少し近いモードでは、2 つの監視オブジェクト、2 つのスライディング ウィンドウがあり、単一 CPU あたり約 700K fps です。
これらの数字は、悲観的な気分のときに読むのが最適です。
多数のレポートやデバッグ印刷を含む多数の監視オブジェクトをコレクタにロードすると、CPU あたり 100K fps 以下で停止する可能性があります。
1 つの CPU では 700K fps 以上を処理できない可能性が高い
複数コアへの拡張については、以下のドキュメントで説明されています。
v25.02 リリースには、LXC コンテナー イメージ xe2502.tar.xz が付属しています。これは特権コンテナであり、ホスト ネットワークを使用するように構成されています。この構成は細心の注意を払って使用してください。コンテナーには、事前に構成されたいくつかの監視オブジェクト、PostgreSQL および Grafana を含むコレクターが含まれています。
# lxc をインストールする
$ sudo apt install lxc
# コンテナイメージを解凍します
$ sudo tar Jxf xe2502.tar.xz -C /var/lib/lxc
# コンテナを実行する
$ sudo lxc-start --name xe2502
# コンテナシェルを実行する
$ sudo lxc-attach --name xe2502
コンテナ内で、ファイル /etc/xenoeye/xenoeye.conf を編集します。
pcap を使用して *flow をキャプチャしている場合は、機能を追加します。
# setcap "cap_net_admin,cap_net_raw,cap_dac_read_search,cap_sys_ptrace+pe" /usr/local/bin/xenoeye
ファイル /var/lib/xenoeye/iplists/mynet を編集し、そこにネットワーク (IPv4 および IPv6) を書き込み、不要なネットワークを削除します。
# サービスxenoeyeの再起動
ブラウザで http://server-address:3000 に移動すると、Grafana が開きます。ロジ

n/パスワード admin/admin。
Grafana には、IPv4 アドレスと IPv6 アドレスに個別に事前設定されたいくつかのダッシュボード (概要、AS/GeoIP、ルーター、DoS/DDoS) が付属しています。以下のドキュメントでは、他のレポートを追加し、移動平均を構成する方法について説明します。
Proxmox のテンプレートも利用できます: proxmox-xe2502.tar.xz
コレクターをインストールして構成するための段階的な手順
Netflow パケットの受信を確認する
複数の CPU にわたる負荷分散
DBMS にエクスポートするデータを構成する
IPアドレスによるシンプルなレポート
スパムボットと SSH スキャナーを検出する
Grafana による交通の可視化
Telegram ボットを使用した異常アラート
コレクターを再起動せずにデータベースを更新する
Grafana を使用した GeoIP データと AS 名の視覚化
sFlow を使用した追加のデータ分析: DNS および SNI
ネスト/階層型監視オブジェクト
コレクターを再起動せずに移動平均しきい値を変更する
設定ファイルの完全な説明
メイン設定ファイル xenoeye.conf
デバイス設定 (サンプリング レートとインターフェイス分類) devices.conf
監視オブジェクトmo.confの説明
オブジェクトとフィルターの監視
新しい Netflow フィールドをコレクターに追加する方法
現時点では、新しい機能を追加する予定はありません。私たちは安定性、作業結果を検討し、バグを修正し、コードをよりシンプルで理解しやすいものにするよう努めます。
軽量の Netflow/IPFIX/sFlow コレクターおよびアナライザー
読み込み中にエラーが発生しました。このページをリロードしてください。
4
フォーク
レポートリポジトリ
リリース
3
v25.02
最新の
2025 年 2 月 17 日
+ 2 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Lightweight Netflow/IPFIX/sFlow collector and analyzer - vmxdev/xenoeye

Sorry for the slightly truncated title. It should have been "Network traffic analysis and monitoring without AI, using netflow-family protocols, PostgreSQL or ClickHouse, Grafana, and some scripts". In 2026, it might seem a bit presumptuous to announce AI-free software on HN.
But building a netflow analyzer manually is no less presumptuous! There are quite a few xFlow analyzers out there these days, and I'm constantly reminded of this. But I think there's always room for an alternative approach. After all, that's how software evolves, isn't it? So, how does xenoeye differ from popular (at least from popular open source) analyzers? - The analyzer has a feature called "monitoring objects". For some reason, open-source analyzers rarely use this feature, while commercial ones do.
The monitoring object can be a subnet, autonomous system, geo-object (data on geo and AS are taken from external databases), application traffic (protocol, TCP/UDP ports, etc.), VLAN, etc.
Almost everything in flow records can be used as a filter for a monitoring object.
Of course, object filters can be composite - the classic operations AND, OR, NOT are supported. The analyzer contains a tiny virtual machine that matches each flow to an object. - We don't store all flows. At least for now. It may seem strange, but this is an important feature, especially for large networks.
We store aggregated data on monitored objects. The user chooses what to store. It could be just in/out, top talkers, top protocols, etc.
The time for which to aggregate data is also specified by the user.
Aggregation occurs inside the analyzer. We use a fast trie-based in-memory db.
Because of this, the analyzer can process flows quite quickly (hundreds of thousands of FPS per vCPU) and export a measured amount of information to the database.
You can easily use even vanilla PostgreSQL. Or ClickHouse with compression.
The analyzer is not very resource-intensive; small network traffic can be processed on low-end hardware or in a VM with a small amount of memory.
Or you can process large network traffic on a single server, without building clusters. I know of installations with multi-terabit traffic and hundreds of MOs on a single virtual machine (of course they have a high sampling rate on their routers). - We can monitor traffic thresholds being exceeded using moving averages.
That is, as soon as an excess is detected, an external script is launched at the same second (actually even faster).
This feature is typically used to detect volumetric DoS/DDoS attacks.
The scripts announce BGP Blackhole or BGP Flowspec and notify users via messenger. - We don't have our own visualization utility; we use Grafana. Grafana works with PostgreSQL out of the box, although some complex time-series charts require some tinkering with SQL queries.
Ok, it's a controversial decision, but users (and we ourselves) are putting up with it for now. I tried to describe the rest in the documentation. Yes, this isn't the first time I've tried to announce this project on HN, and I'm under no illusions - for some reason, hackers aren't very fond of this type of software.
Perhaps everyone thinks that the production of netflow analyzers is too boring a matter, there is nothing to discuss. However, if anyone is interested, it would be great to get feedback. What would you do differently than it was done and why?
What do you like most about your favorite analyzer that you can't find anywhere else? How did you even see this post? This isn't AI or even a Rust-related thing

GitHub - vmxdev/xenoeye: Lightweight Netflow/IPFIX/sFlow collector and analyzer · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
vmxdev
/
xenoeye
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
349 Commits 349 Commits .github/ workflows .github/ workflows aajson @ 7abedc5 aajson @ 7abedc5 docs-img docs-img lxc lxc m4 m4 scripts scripts tests tests tkvdb @ 61efc75 tkvdb @ 61efc75 .dockerignore .dockerignore .gitattributes .gitattributes .gitignore .gitignore .gitmodules .gitmodules .travis.yml .travis.yml CONFIG.md CONFIG.md CONFIG.ru.md CONFIG.ru.md Dockerfile Dockerfile EXTRA.md EXTRA.md EXTRA.ru.md EXTRA.ru.md INTERNALS.md INTERNALS.md INTERNALS.ru.md INTERNALS.ru.md LICENSE LICENSE Makefile.am Makefile.am README.md README.md README.ru.md README.ru.md STEP-BY-STEP.md STEP-BY-STEP.md STEP-BY-STEP.ru.md STEP-BY-STEP.ru.md classification.c classification.c configure.ac configure.ac devices.c devices.c devices.conf devices.conf devices.h devices.h docker-compose-dev.yml docker-compose-dev.yml docker-compose.yml docker-compose.yml filter-ag.def filter-ag.def filter-lexer.c filter-lexer.c filter-parser-funcs.c filter-parser-funcs.c filter-parser.c filter-parser.c filter.c filter.c filter.def filter.def filter.h filter.h flow-debug.c flow-debug.c flow-debug.h flow-debug.h flow-info.h flow-info.h geoip.c geoip.c geoip.h geoip.h ip-btrie.h ip-btrie.h iplist.c iplist.c iplist.h iplist.h monit-objects-common.h monit-objects-common.h monit-objects-fwm.c monit-objects-fwm.c monit-objects-mavg-act.c monit-objects-mavg-act.c monit-objects-mavg-dump.c monit-objects-mavg-dump.c monit-objects-mavg-limfile.c monit-objects-mavg-limfile.c monit-objects-mavg-under.c monit-objects-mavg-under.c monit-objects-mavg.c monit-objects-mavg.c monit-objects.c monit-objects.c monit-objects.h monit-objects.h netflow-templates.c netflow-templates.c netflow-templates.h netflow-templates.h netflow.c netflow.c netflow.def netflow.def netflow.h netflow.h pcapture.c pcapture.c rawparse.h rawparse.h scapture.c scapture.c sflow-impl.h sflow-impl.h sflow.c sflow.c sflow.h sflow.h utils-data.inc utils-data.inc utils.c utils.c utils.h utils.h xe-debug.h xe-debug.h xe-dns.h xe-dns.h xe-sni.h xe-sni.
[truncated]
Lightweight Netflow/IPFIX/sFlow collector and analyzer
README.ru.md - документация на русском
The documentation is mostly translated automatically using Google translator, so if you see something weird - feel free to let us know.
Monitor traffic of IP networks, individual IP addresses or services
React quickly to traffic spikes or traffic drops below thresholds
Monitor traffic patterns and distribution of network packets using data from Netflow/IPFIX/sFlow
The collector was developed for medium and large networks, with different user groups that need different reports. For this purpose, "monitoring objects" are used. A monitoring object can be a network, a set of networks, an autonomous system, a geo-object or arbitrary network traffic that can be extracted from Netflow/IPFIX/sFlow.
Using the collector, you can generate various reports, build charts, dashboards in Grafana, perform some actions when the traffic speed exceeds thresholds or falls below thresholds.
We use the collector to monitor our networks. We are using Netflow v9 and IPFIX, so the collector supports them.
Netflow v5 and sFlow are also supported.
The documentation contains examples of building simple reports. To build more complex ones, you need at least basic knowledge of SQL.
The collector uses text configuration files. This allows you to write simple configs manually, and for complex configurations with a large number of objects, you can generate configs using scripts.
The collector processes data in two ways: it aggregates it over periods (fixed-size time windows to produce reports and graphs), and it uses moving averages to quickly react to spikes.
Both methods can be used individually or together. For example, if a moving average detects a threshold being exceeded, you can run a custom script and immediately enable extended statistics collection.
We use moving averages to detect volumetric DoS/DDoS attacks. When thresholds are reached, BGP announcements are created (FlowSpec filtering, rate-limit, redirection to cleaning servers or Blackhole) and users receive a notification in the messenger.
Collector is not very demanding on resources. It can process data and build reports even on Orange Pi (analogous to Raspberry Pi) with 4 GB of memory. On small networks it can run in a VM with one CPU and 1GB of RAM.
The collector has only been tested under 64-bit Linux (x64, AArch64 and Elbrus ).
We use PostgreSQL as a storage for time series data. Aggregated data by selected Netflow fields is exported there. The collector can export not all data to the DBMS, it can aggregate and export only top-N entities, and aggregate the rest into one row. This is a useful feature for large monitoring objects - you can regulate the amount of data that is written to the DBMS and use cheaper, slower disks.
In addition to PostgreSQL, the collector has experimental support for storing data in ClickHouse
A basic set of Netflow/IPFIX fields are supported out of the box, but you can add almost any field you need.
The project has a very liberal ISC license. We have no plans to make commercial or semi-commercial versions. This means that we cannot make any predictions about the future of the project. But on the other hand:
There are no hidden or artificial restrictions
Users are usually interested in at least a rough performance estimate, so we made several tests: we recorded real Netflow traffic from different routers in pcap files and played them on the loopback interface using tcpreplay at different speeds.
Tests were run on i3-2120 CPU @ 3.30GHz.
Very roughly, you can rely on following numbers:
In debug mode, when the contents of each flow are printed to a file, it turned out about 100K flow per second per one CPU.
In a slightly closer to production mode, with two monitoring objects, two sliding windows - about 700K fps per single CPU.
These numbers are best read in a pessimistic mood:
if you load the collector with many monitoring objects with a bunch of reports and debug printing, it can choke on 100K fps/CPU or less
most likely 700K fps and more cannot be processed on one CPU
Scaling to multiple cores is described below in the documentation
The v25.02 release comes with an LXC container image xe2502.tar.xz . This is a privileged container and is configured to use the host network , use this configuration with extreme caution. The container contains a collector with several pre-configured monitoring objects, PostgreSQL and Grafana.
# install lxc
$ sudo apt install lxc
# unpack the container image
$ sudo tar Jxf xe2502.tar.xz -C /var/lib/lxc
# run container
$ sudo lxc-start --name xe2502
# run container shell
$ sudo lxc-attach --name xe2502
Inside the container, edit the file /etc/xenoeye/xenoeye.conf
If you are capturing *flow with pcap, add capabilities:
# setcap "cap_net_admin,cap_net_raw,cap_dac_read_search,cap_sys_ptrace+pe" /usr/local/bin/xenoeye
Edit the file /var/lib/xenoeye/iplists/mynet , write your networks there (IPv4 and IPv6), and delete unnecessary ones.
# service xenoeye restart
Navigate your browser to http://server-address:3000 , Grafana should open. Login/password admin/admin.
Grafana comes with several pre-configured dashboards (Overview, AS/GeoIP, Routers, DoS/DDoS) separately for IPv4 and IPv6 addresses. The documentation below describes how to add other reports and configure moving averages.
A template for Proxmox is also available: proxmox-xe2502.tar.xz
Step-by-step instructions for installing and configuring the collector
Checking Netflow packets receiving
Load-balancing across multiple CPUs
Configure what data should be exported to the DBMS
Simple Reporting by IP Addresses
Detect spam-bots and ssh-scanners
Traffic visualization with Grafana
Anomaly alerts using Telegram-bot
Updating databases without restarting the collector
Visualizing GeoIP data and AS names with Grafana
Additional data analysis using sFlow: DNS and SNI
Nested/Hierarchical Monitoring Objects
Changing moving average thresholds without restarting the collector
Full description of configuration files
Main configuration file xenoeye.conf
Device configuration (sampling rate and interface classification) devices.conf
Description of the monitoring object mo.conf
Monitoring objects and filters
How to add a new Netflow field to the collector
Right now we don't plan to add new features. We look at stability, work results, try to fix bugs and make the code simpler and more understandable.
Lightweight Netflow/IPFIX/sFlow collector and analyzer
There was an error while loading. Please reload this page .
4
forks
Report repository
Releases
3
v25.02
Latest
Feb 17, 2025
+ 2 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
