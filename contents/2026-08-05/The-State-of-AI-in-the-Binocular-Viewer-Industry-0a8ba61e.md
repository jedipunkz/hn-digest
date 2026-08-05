---
source: "https://www.adamarice.com/ai-binocular-viewers/"
hn_url: "https://news.ycombinator.com/item?id=49183790"
title: "The State of AI in the Binocular Viewer Industry"
article_title: "The State of AI in the Binocular Viewer Industry"
author: "adamarice"
captured_at: "2026-08-05T15:09:10Z"
capture_tool: "hn-digest"
hn_id: 49183790
score: 2
comments: 0
posted_at: "2026-08-05T14:49:30Z"
tags:
  - hacker-news
  - translated
---

# The State of AI in the Binocular Viewer Industry

- HN: [49183790](https://news.ycombinator.com/item?id=49183790)
- Source: [www.adamarice.com](https://www.adamarice.com/ai-binocular-viewers/)
- Score: 2
- Comments: 0
- Posted: 2026-08-05T14:49:30Z

## Translation

タイトル: 双眼ビューワ業界における AI の現状
説明: 前置き: 「人間が AI を使用してソフトウェアを構築する」というジャンルは、現時点では珍しいジャンルではありません。私はそれに加えようとしています。しかし、どの業界も最終的には AI の現状に関するレポートを作成することになり、他の誰もこれを書くつもりはなかったということに注意したいと思います。
もしあなたが立っていたなら

記事本文:
サインイン
購読する
タワーオプティカル
双眼ビューワ業界における AI の現状
前置き: 「人間が AI を使ってソフトウェアを構築する」というジャンルは、現時点では珍しいジャンルではありません。私はそれに加えようとしています。しかし、どの業界も最終的には AI の現状に関するレポートを作成することになり、他の誰もこれを書くつもりはなかったということに注意したいと思います。
トップ・オブ・ザ・ロックに立ったり、ナイアガラの滝の周りを歩いたり、自由の女神からマンハッタンのスカイラインを眺めたりしたことがあるなら、おそらく当社のマシンの 1 つを使用したことがあるでしょう。それは、ほぼ 1 世紀にわたってアメリカのランドマーク体験の一部となっているクロム双眼鏡ビューアです。少なくとも私たちのマシンの 1 つを見たことがあるはずです。 Tower Optical Company は 1933 年からこの装置を製造してきました。現在、800 以上の拠点で 2,000 台近くの当社の機械が現場にあります。 1 四半期か 2 四半期でドロップするか、最近では携帯電話またはカードをタップすると、スカイラインに焦点が合います。
パートナーと私は、2025 年 6 月にこのビジネスを買収しました。私たちが購入したのは、深刻な物流上の問題を抱えた正真正銘のアメリカの象徴でした。数百のサイトに数千台のマシンがあり、そのほとんどはサービスを提供できる人から長距離ドライブで遠く離れており、それらのマシンが何をしているのかを知る方法はありませんでした。各ユニットはコイン式でした。機械の収益は収集日にその機械から出たもので、紙の伝票に記録されていました。会場パートナーへの支払いは書面による小切手で行われ、事後的にすべてが QuickBooks に入力されました。動作を停止したビューアは、誰かがそれに気づき、本社のチームに警告するまで壊れたままでした。これは無視されたものではなく、マシンが生まれた世界のために構築されたシステムであり、多かれ少なかれそのようにして何十年も運用されてきました。
2 つのことを変える必要がありました。機械は現在人々が実際に支払うのと同じ方法で支払いを受け取る必要があり、私たちは逃走を管理する方法が必要でした。

効率的に。
そこで私たちはエンジニアリングパートナーと協力して機械用のタップ・トゥ・ペイ改造キットを開発し、その後自分たちで改造作業を行い、150 台のユニットを社内で改造しました。ついにこの 6 月にそれらの展開を開始しました。新しい部隊は、古い部隊には決してできなかったことを実行します。つまり、ほぼリアルタイムで本部に報告を返します。
これにより、1 つの問題が解決されましたが、別の問題が発生しました。初めて実際のデータをフリートから取得し、最初はスプレッドシートに取り込みました。すべてのタップが列になりました。どのマシンが収益をあげていて、どのマシンが止まっていたのかを把握し、その理由 (バッテリー切れ? 決済端末がフリーズ? 単なる雨の火曜日?) を解明するには、ファイルをダウンロードし、行と列を注意深く見る必要がありました。
そして、取引データはその一部にすぎませんでした。変換された各マシンには、独自の小さな物流上の問題があります。カード処理用の販売者 ID、独自のシリアル番号を持つ支払い端末、別のコントロール ボード、ファームウェア バージョン、セルラー接続、2 つのバッテリー パック、GPS 座標、ロックを解除するための特定の物理キー、オンサイト連絡先、メンテナンス連絡先、パートナー会場との交渉による収益分割などです。これに 150 台のマシンを掛けても、残りのマシンはまだ登場しないため、スプレッドシートとピボット テーブルがごちゃ混ぜでは絶対に足りません。
これは、トランザクション データを視覚化するための 1 つの HTML ファイルとして始まり、新しいマシンのパフォーマンスを確認できるようになりました。これは、決済プロセッサからエクスポートされたシンプルな CSV ファイルで、Web ブラウザで視覚化されました。約 6 週間後、これは本番 Web アプリケーション (Next.js と Supabase で構築され、Vercel にデプロイされた) となり、チームとパートナーが毎日使用する接続フリートの運用の中核を実行しています。
マシンごと、サイトごとのライブ収益: 独自のタイムゾーンのすべての場所、トランザクション記録に遡って追跡可能なすべての数値

rds は当社の決済プロセッサーによって配信されます。
マシンの状態を知らせるマシン: 視聴者は、バッテリー電圧、ファームウェアのバージョン、携帯電話の信号などのテレメトリ データをレポートし、ダッシュボードは Google Pub/Sub を通じてこれらのデータを取り込んで表示します。ダッシュボードは、ユニットが故障する前にバッテリーの交換を予測し、「故障」と「雨上がり」の違いを認識しながら数時間以内に静かになった視聴者にフラグを立てます。地元の天気と会場の閉鎖パターンを相互参照するため、誤報ではなく実際の問題を追跡します。
すべてのマシンの単一の記録システム: シリアル番号、ファームウェアのバージョン、販売者 ID、SIM カード、物理キーの種類、連絡先、座標、収益分割、サイト マップなどのすべてが、ファイルや誰かのメモリに分散するのではなく、所属するマシンに関連付けられて 1 か所に存在します。
フリートを認識する AI アシスタント: ダッシュボードに組み込まれています。わかりやすい英語で質問すると (たとえば、今週どのマシンが静かになったか、視聴者が最後にサービスを受けたのはいつか、サイトの公開以来の傾向はどうかなど)、ライブ データから回答します。スプレッドシートの探索を終わらせたのと同じ記録システムは、AI が参照できる信頼できる場所が 1 つあることも意味します。
マシンの健全性レポートとメンテナンス リスト: すべてのユニットが健全性レポート (信号、バッテリー、稼働時間、最近/疑わしい問題) とそれに対応するサービス チェックリストを生成します。現場に向かう技術者は、私たちのチームの一員であれ、会場のメンテナンススタッフであれ、各マシンで何が起こっているのか、何をしなければならないのかを把握しながら現場に入ります。
収益分配の簡素化: 変換されたマシンごとに、決済されたトランザクションからパートナーの収益分配明細書が自動的に生成されます。計算や手動の調整は不要で、完全な監査証跡が必要です。 ACH または ch 経由でパートナーに支払いを送金できます

API 経由でワンクリックで当社の銀行口座からチェックし、同時に当社のダッシュボードから収益分配明細書を PDF 形式で直接送信します。
あらゆる役割に対応したビュー: フィールド サービス チームは、機械の前に立って、機械を引き上げ、GPS 座標を取得し、ステータスを更新し、メモを追加し、写真を撮り、バッテリーを交換し、次に進むために構築されたモバイル バージョンを入手します。会場パートナーは、ヘルスアラートを備えた自身のマシンのライブダッシュボードを取得できるため、問題を自分で発見する前に私たちから聞くことができます。
私はエンジニアではなく製品担当です。技術愛好家であり、キャリアの何年もコードを書くことなく過ごしてきました。それはここでも変わりませんでした。クロード・コードには間違いなく多くの手を握る必要があり、その手を握ることがポイントでした。必要なコンテキストはビジネスの運営から得られます。 AI はドメイン知識に取って代わるものではありません。これにより、ドメインの専門知識を持つ人と実用的なソフトウェアを構築する間の翻訳層が取り除かれました。
ビジネスのために何ができるのか
マシンがレポートを停止すると、次に誰かがサイトにアクセスするたびではなく、ほぼ即座にフラグが立てられます。ビュー 1 ドルの場合、混雑したランドマークでのダウンタイムはすぐに加算され、パートナーからの電話かメンテナンス訪問で判明するまで、ダウンタイムは目に見えませんでした。
誰がどのマシンを担当するか、何をする必要があるか、何がすでに完了したかなど、問題の修正に関する調整は、チームと会場スタッフの間の電話や電子メールのスレッドではもはや行われません。旗、チェックリスト、メンテナンス記録は一緒に移動するため、社内外の関係者間でフリートの状態を管理すること自体は、主要な管理業務ではなくなりました。
送金作業は、数時間にわたる書類作成とスプレッドシートの調整から、毎月数分のレビューに変わりました。

「このマシンについて何がわかっていますか?」という質問。答えは 5 つではなく 1 つになりました。シリアル番号、ファームウェア、連絡先、キー番号、収益分割、サービス履歴、正確な物理的な場所など、ファイル キャビネットやスプレッドシート、あるいは誰かの記憶の中だけに存在するものはもうありません。
タップして支払う機械を使用しているパートナーは、明細書を待つのではなく、番号をリアルタイムで確認できます。現在では、私たちが想定しているものではなく、データが示す内容に基づいてマシンを配置しています。
数千の艦隊に 150 ユニットが所属しています。これらを実行するプラットフォームは、月額 200 ドルの Claude サブスクリプションを利用して、ビジネスを運営する人によって約 6 週間で構築されました。この数字が何を置き換えたのかを明確にしておきたいと思います。正当化できなかったエンジニアの雇用、ビジネスの運用上の微妙な違いを理解するだけでも数週間または数か月を必要とした契約開発工場、またはおそらく、プロジェクトがまったく実行されず、スプレッドシート上で最新のフリートを永久に実行するバージョンです。
3 番目の選択肢は、正直な反事実です。数年前、当社のような企業 (物理的、分散型、創業 90 年以上、エンジニアリング チームが存在しない) はカスタム ソフトウェアを入手できませんでした。フルスタック チームはコード行を書く前に 6 桁の決断を下す必要があり、そのチームにビジネスの運用上の微妙な違いを教えるには多くの時間がかかります。 AI により、これらのコストは両方とも事実上ゼロになりました。ニュアンスの説明が必要ない人は、そのまま製品を構築できます。
私の心に残っているのは、これがどれほど再現可能であるかということです。あらゆるニッチな業界、特殊すぎるビジネス、物理的すぎるビジネス、カスタム ソフトウェアが意味を成さないほど小さすぎるビジネスが、それにアクセスできるようになりました。
ビジネス、テクノロジー、マクロ経済に関する最新情報を電子メールで不定期に受け取ります。

## Original Extract

Up front: "guy builds software with AI" is not a scarce genre right now, and I'm about to add to it. But I'd note that every industry gets a state-of-AI report eventually, and nobody else was going to write this one.
If you've stood at

Sign in
Subscribe
Tower Optical
The State of AI in the Binocular Viewer Industry
Up front: "guy builds software with AI" is not a scarce genre right now, and I'm about to add to it. But I'd note that every industry gets a state-of-AI report eventually, and nobody else was going to write this one.
If you've stood at the Top of the Rock, walked around Niagara Falls, or looked out at the Manhattan skyline from the Statue of Liberty, you've probably used one of our machines: the chrome binocular viewers that have been part of the American landmark experience for nearly a century. You’ve certainly at least seen one of our machines. The Tower Optical Company has been building them since 1933. There are nearly 2,000 of our machines in the field today, across 800-plus sites. Drop in a quarter or two, or these days, tap your phone or a card, and the skyline snaps into focus.
My partners and I acquired the business in June of 2025. What we bought was a genuine American icon with serious logistical problems attached: thousands of machines across hundreds of sites, most of them a long drive from anyone who could service them, and no real way to know what any of them were doing. Every unit was coin-operated. A machine's earnings were whatever came out of it on collection day, recorded on paper slips. Venue partners were paid by written check, and everything got typed into QuickBooks after the fact. A viewer that stopped working stayed broken until a person noticed and alerted our team at headquarters. This wasn't neglect, it was a system built for the world the machines were born into, and it had run that way, more or less, for decades.
Two things had to change. The machines needed to take payment the way people actually pay now and we needed a way to manage the fleet efficiently.
So we developed a tap-to-pay retrofit kit for the machines with an engineering partner, then did the conversion work ourselves, 150 units retrofitted in-house. We finally began rolling them out this past June. The new units do something the old ones never could: they report back to HQ in near real-time.
That solved one problem and created another. For the first time, we had real data coming from the fleet and, at first, it went into spreadsheets. Every tap became a row. Figuring out which machines were earning, which had gone quiet, and figuring out why (Dead battery? Payment terminal frozen? Just a rainy Tuesday?) meant downloading files and squinting at rows & columns.
And the transaction data was only part of it. Each converted machine is its own small logistics problem: a merchant ID for card processing, a payment terminal with its own serial number, a control board with another, a firmware version, a cellular connection, two battery packs, GPS coordinates, a specific physical key to unlock it, an on-site contact, a maintenance contact, and a negotiated revenue split with the partner venue. Multiply that by 150 machines, with the rest of the fleet still to come, and a jumble of spreadsheets and pivot tables was never going to cut it.
It started as a single HTML file to visualize transaction data so we could see how our new machines were performing – a simple CSV file exported from our payment processor, visualized in a web browser. About six weeks later, it’s a production web application (built on Next.js and Supabase, deployed on Vercel) running the operational heart of the connected fleet that our team and partners use daily:
Live revenue, per machine, per site: Every location in its own time zone, every figure traceable back to transaction records delivered by our payment processor.
Machines that tell us how they're doing: The viewers report telemetry data such as battery voltage, firmware versions, and cellular signal, which our dashboard ingests & displays through Google Pub/Sub. The dashboard predicts battery swaps before a unit dies and flags viewers that go quiet within hours while knowing the difference between “broken” and “rained out”. It cross-references local weather and venue closure patterns, so we chase down real problems instead of false alarms.
A single system of record for every machine: All of it… serial numbers, firmware versions, merchant IDs, SIM cards, physical key types, contacts, coordinates, revenue splits, and site maps live in one place, attached to the machine they belong to, instead of scattered across files and someone's memory.
An AI assistant that knows the fleet : Built into the dashboard: ask it a question in plain English (for example, which machines went quiet this week, when a viewer was last serviced, how a site trended since launch) and it answers from the live data. The same system of record that ended the spreadsheet hunting also means the AI has one trustworthy place to look.
Machine health reports and maintenance lists : Every unit generates a health report (signal, battery, uptime, recent/suspected issues) and a service checklist to match. A tech heading to a site, whether someone on our team or a venue's own maintenance crew, walks in knowing what's going on with each machine and what needs to be done.
Simplified revenue sharing: For every converted machine, the partner's revenue share statement is generated automatically from settled transactions. No math, no manual reconciliation, and a full audit trail. We can remit payment to our partners via ACH or check from our bank account with one click via API and simultaneously send them their revenue sharing statement in PDF form directly from our dashboard.
A view for every role: The field service team gets a mobile version built for standing in front of a machine – pull it up, capture GPS coordinates, update its status, add notes, take a photo, swap a battery, move on. Venue partners get a live dashboard of their own machines, with health alerts, so they hear about a problem from us before they discover it themselves.
I'm a product guy, not an engineer – a tech enthusiast who's spent years of his career adjacent to code, mostly without writing it. That didn't change here. Claude Code definitely needed a lot of hand-holding, and the hand-holding was the point. The necessary context comes from running the business. AI didn't replace domain knowledge. It removed the translation layer between the person who has domain expertise and building working software.
What it's done for the business
A machine that stops reporting gets flagged almost immediately, instead of whenever someone next happens to visit the site. At a dollar a view, downtime at a busy landmark adds up quickly, and it used to be invisible until it turned up on a phone call from a partner or a maintenance visit.
The coordination around fixing things, such as who's responsible for which machine, what needs doing, and what's already been done, no longer lives in phone calls and email threads between our team and venue staff. The flag, the checklist, and the maintenance record travel together, so managing fleet health across internal and external parties has stopped being a major administrative job in itself.
Remittance work went from hours of paperwork & spreadsheet reconciliation to minutes of review each month.
The question "what do we know about this machine?" now has one answer instead of five. Serial numbers, firmware, contacts, key numbers, revenue splits, service history, precise physical locations -- none of it lives in a filing cabinet, a spreadsheet, or only in somebody's memory anymore.
Partners with tap-to-pay machines can watch their numbers live rather than waiting on a statement. Now, we place machines based on what the data says, rather than on what we assume.
We're 150 units into a fleet of a couple thousand. The platform that runs them was built in about six weeks, by a person who runs the business, on a $200/month Claude subscription. I want to be clear about what that number replaced: an engineering hire we couldn't justify, a contract dev shop that would have needed weeks or months just to understand the operational nuances of our business, or, most likely, the version where the project simply never happens and we run a modernized fleet on spreadsheets forever.
That third option is the honest counterfactual. A few years ago, a business like ours – physical, distributed, ninety-plus years old, no engineering team – didn't get custom software. A full-stack team is a six-figure decision before it writes a line of code, and teaching that team the operational nuances of the business requires a lot of time. AI took both of these costs to virtually zero. The people who don’t need the nuances explained to them can just build the product.
What stays with me is how repeatable this feels. Every niche industry, businesses too specific, too physical, too small for custom software to ever make sense, just got access to it.
Receive infrequent email updates on business, technology, and macroeconomics.
