---
source: "https://en.wikipedia.org/wiki/Npm_left-pad_incident"
hn_url: "https://news.ycombinator.com/item?id=49295443"
title: "NPM Left-Pad Incident"
article_title: "npm left-pad incident - Wikipedia"
author: "gherkinnn"
captured_at: "2026-08-14T07:09:07Z"
capture_tool: "hn-digest"
hn_id: 49295443
score: 3
comments: 0
posted_at: "2026-08-14T06:52:04Z"
tags:
  - hacker-news
  - translated
---

# NPM Left-Pad Incident

- HN: [49295443](https://news.ycombinator.com/item?id=49295443)
- Source: [en.wikipedia.org](https://en.wikipedia.org/wiki/Npm_left-pad_incident)
- Score: 3
- Comments: 0
- Posted: 2026-08-14T06:52:04Z

## Translation

タイトル: 故宮左パッド事件
記事タイトル: npm 左パッド事件 - Wikipedia

記事本文:
コンテンツへジャンプ
メインメニュー
メインメニュー
サイドバーに移動
隠す
ナビゲーション
メインページ
3
余波
余波サブセクションの切り替え
3.1
即時的な効果
2016 年 3 月 22 日、プログラマーの Azer Koçulu は、npm (JavaScript のパッケージ マネージャー) に公開した左パッド パッケージを削除しました。 Koçulu 氏は、Kik Messenger との紛争の後、パッケージ名 kik の管理を npm, Inc. が認められた後、自分のパッケージをすべて削除しました。その結果、Babel トランスコンパイラや React Web フレームワークなど、左パッドを依存関係として使用していた何千ものソフトウェア プロジェクトがビルドまたはインストールできなくなりました。 Facebook 、 PayPal 、 Netflix 、 Spotify などの大小のテクノロジー企業がソフトウェア製品の依存関係として左パッドを使用したため、これは広範な混乱を引き起こしました。
パッケージが npm から削除されてから数時間後、プラットフォームの背後にある企業である npm, Inc. が手動でパッケージを復元しました。その後、npm は、パッケージの公開日から 24 時間以上経過し、少なくとも 1 つの他のプロジェクトがパッケージに依存している場合にパッケージを削除する機能を無効にしました。この事件はメディアの注目を集め、ソフトウェア業界の人々からの反応を呼び起こしました。左パッドの削除は、抗議行為としてのソフトウェアの意図的な自己破壊に関する議論を引き起こし、モジュール型プログラミングにおけるサプライチェーン攻撃の可能性の高まりに注目を集めました。
left-pad は、カリフォルニア州オークランドに拠点を置く独立系プログラマーである Azer Koçulu によって公開された、無料のオープンソース JavaScript パッケージです。 [1] パッケージは、ループを使用して文字列の先頭に文字を繰り返し追加します。 [ 1 ] 左パッドは非常にシンプルであるという特徴があり、 [ 2 ] [ 3 ] 以下に示す Koçulu によって作成された最終バージョンでは、わずか 11 行のコードで構成されています。 [ 4 ]

モジュール。エクスポート = 左パッド ;
関数 leftpad (str , len , ch ) {
str = 文字列 ( str );
変数 i = - 1 ;
チャンネル || ( ch = ' ' );
len = len - str 。長さ ;
while ( ++ i < len ) {
str = ch + str ;
}
str を返します。
}
Koçulu は、JavaScript ランタイム環境である Node.js のデフォルトのパッケージ マネージャーである npm で左パッドを公開しました。 [ 5 ] [ 2 ] 比較的知られていないにもかかわらず、左パッドは頻繁に使用されました。このパッケージは他の何千ものソフトウェア プロジェクトによって依存関係として使用され、削除されるまでに 1,500 万回以上ダウンロードされました。 [ 6 ] [ 7 ] 左パッドを必要とするプロジェクトの中には、当時の JavaScript エコシステムにとって重要なものもありました。これには、下位互換性のある JavaScript コードを可能にするトランスコンパイラーである Babel が含まれていました。 Webpack 、モジュールバンドルシステム。 React と React Native は、それぞれ Web サイトとモバイル アプリの開発に広く使用されているフレームワークです。 [ 8 ] [ 9 ] [ 1 ]
Koçulu は left-pad に加えて、開発者がプロジェクトのテンプレートをセットアップできるツールである npm の kik も所有していました。 [1] 2016 年 3 月 11 日、インスタント メッセージング プラットフォーム Kik Messenger を所有するカナダ企業 Kik Interactive は、同社が「Kik」商標を所有しているため、Koçulu に対し kik パッケージの管理を放棄するよう要請しました。 [10] 通信の一部には、Kik からの次のメッセージが含まれていました。
私たちは [kik パッケージ] について馬鹿にするつもりはありませんが、これは世界中のほとんどの国で登録商標であり、実際に kik と呼ばれるオープンソース プロジェクトをリリースした場合、私たちの商標弁護士がドアを叩き、アカウントなどを削除するでしょう。そして、商標を強制する必要があるため、そうしなければ商標を失うため、私たちにはそうする以外に選択肢はありません。あなたに影響を与えずに名前を変更してもらうために、何らかの妥協点を見つけることはできないでしょうか

弁護士をやっていますか？名前を変更してもらうために、補償として何かできることはありますか? [ 3 ]
コチュル氏はその直後に返答し、プロジェクト名の変更を拒否し、次のように述べた。
"}},"i":0}}]}'/> ははは、あなたは本当にクソ野郎です。だから、クソ。私にメールを返信しないでください。 [ 3 ]
コチュル氏はまた、「企業のクソ野郎どものために自分の得意プロジェクトを諦める手間の対価として」3万ドルの補償も要求した。 [ 1 ]
2016 年 3 月 18 日、npm, Inc. の最高経営責任者である Isaac Z. Schlueter は、Kik Interactive と Koçulu に書簡を送り、kik パッケージの所有権が Kik Interactive に手動で移管されると述べました。 [ 1 ]
Koçulu 氏が npm, Inc. の決定に失望し、プラットフォームに参加することを望まないと述べた後、Schlueter 氏は彼に、登録した 273 個のモジュールをすべて削除するコマンドを提供しました。 [10] Koçulu は 2016 年 3 月 22 日にコマンドを実行し、以前にリリースしたすべてのパッケージを削除しました。 [ 1 ] left-pad は「非公開」パッケージの 1 つであり、npm で一般にアクセスできなくなりました。 [6] 左パッド ソフトウェア プロジェクトとコンテンツは GitHub で引き続き利用できます。 [ 10 ]
依存関係として左パッドを使用する JavaScript プロジェクト (Babel や Webpack などの依存関係を含む) をビルドまたはインストールしようとすると、プロセスが失敗する原因となる 404 エラーが発生しました。 [ 1 ] このパッケージを使用したソフトウェア技術企業の中には、 Meta Platforms 、 PayPal 、 Netflix 、 Spotify 、 [ 9 ] 、そして Kik Interactive 自体も含まれます。 [ 1 ]
パッケージを削除してから 1 時間後、コチュル氏は Medium に投稿 (「モジュールを解放したところです」) を公開し、無料のオープンソース ソフトウェアに対する企業の利益に抗議するために、npm からソフトウェア プロジェクトを非公開にしたと説明しました。 [ 1 ]
もうすぐ

削除後、他のソフトウェア開発者は、プロジェクトの Git 問題追跡システムに大量の苦情、反応、回避策を投稿し始めました。 [ 8 ] [ 1 ]
Babel を含むオープンソース プロジェクトの管理者は、Koçulu が未公開にしていた依存関係を削除するホットフィックスをリリースしました。 [ 8 ] Koçulu の他のパッケージ名のいくつかは、新しく公開されたパッケージにすぐに引き継がれました。 [ 3 ] たとえば、別の開発者は左パッド パッケージを再作成しましたが、バージョン 1.0.0 としてリリースしました。 Koçulu がバージョン 0.0.3 として公開して以来、ユーザーは引き続き問題に遭遇しました。 [ 3 ]
元の left-pad パッケージが削除されてから約 2 時間後、npm はバックアップを復元することにより、元の 0.0.3 バージョンを手動で復元しました。 [ 1 ] npm の最高技術責任者、ローリー・ヴォスは、この行動が「正しい決断」であるかどうかについて社内で意見の相違があったにもかかわらず、同社は「多くの人のニーズを選んだ」と書いている。 [ 11 ]
npm は、公開されたパッケージの削除に関するポリシーを変更し、リリース日から 24 時間以上経過しており、少なくとも 1 つの他のプロジェクトがそのパッケージを依存関係として必要としている場合には削除されないようにしました。 [12] npmを代表して、コミュニティマネージャーのアシュリー・ウィリアムズは、プラットフォームが「コミュニティを保護することに失敗した」と述べ、事件によって引き起こされた混乱について謝罪した。 [12] Kik Interactiveもこの件について謝罪し、同社のメッセージング責任者であるMike Roberts氏がKoçulu氏との一連の電子メールをMedium上で公開し、彼のやりとりを「丁寧な要求」だったと特徴づけた。 [9] Roberts は、最初に Koçulu に連絡を取ったのは、Koçulu が使用していた名前でオープンソース パッケージを npm で公開したかったからだと書いています。 [6] コチュルは、他人の仕事を妨害したことは申し訳ないと述べたが、自分は「地域社会の利益のため」にそうしたことをしたと信じていたと述べた。

原文のまま] 長期」。 [ 2 ]
この事件は Twitter 、 GitHub 、 Reddit 、 Hacker News などでユーザーからさまざまな反応を引き起こし、その多くは一時的に「インターネットを破壊した」と主張した。 [2] [9] [10] [1] 多くの人が、JavaScript 開発の「素早く動いて物事を壊す」文化、オープンソース ソフトウェアの予測不可能な性質、およびモジュール型プログラミングへの過剰依存の認識についてコメントしました。 [ 2 ] [ 9 ] [ 3 ] ユーザーはまた、法的脅迫を理由に Koçulu のパッケージを Kik Interactive に強制的に譲渡するという npm の決定に対して失望を表明しました。 [ 1 ]
このインシデントは、npm パッケージの破壊がいかにサプライ チェーン攻撃につながる可能性があるかを示しました。広く報道された左パッド事件に加えて、多くの人物がコチュルの他のパッケージが削除された後、未知のコードで即座にハイジャックしていました。 [ 8 ] npm は、同様の紛争における悪意のある乗っ取りを防止するための新しいポリシーをリリースしました [ 3 ] が、左パッド事件は依然としてソフトウェア製品の攻撃対象領域の増加につながる外部貢献者への過度の依存の一例として引用されています。 [ 13 ] コチュル氏の左パッドの意図的な自己破壊行為は、npm などのプラットフォームで公開されるプロテストウェアの発生の前兆とも言われています。 [ 7 ]
ハクティビズム – 抗議の手段としてのコンピューターベースの活動
Peacenotwar – 2022 年マルウェア、Brandon Nozaki Miller 著
ソフトウェア リポジトリ – ソフトウェア パッケージの保管場所
1 2 3 4 5 6 7 8 9 10 11 12 13 14 コリンズ、キース（2016 年 3 月 27 日）。 「あるプログラマーが小さなコードを削除してインターネットを破壊した方法」。クォーツ。 2024 年 5 月 11 日のオリジナルからアーカイブ。 2024 年 5 月 11 日に取得。
1 2 3 4 5 ワインバーガー、マット（2016 年 3 月 23 日）。 「あるプログラマーは、11 行のコードを削除してインターネットを破壊するところでした。」ビジネスインサイダー。からアーカイブされました

2024 年 5 月 11 日のオリジナル。 2024 年 5 月 11 日に取得。
1 2 3 4 5 6 7 ブライアン、フェルドマン（2016 年 3 月 24 日）。 「ある男がインターネットから 11 行のコードを削除し、数百のアプリを破壊した」インテリジェンス。 2024 年 5 月 11 日のオリジナルからアーカイブ。 2024 年 5 月 11 日に取得。
↑ アゼル、コチュル（2014 年 8 月 15 日）。 "left-pad/index.js" 。 GitHub 。 2024-05-11 のオリジナルからアーカイブされました。 2026 年 3 月 2 日に取得。
↑ クラバーン、トーマス（2019 年 4 月 22 日）。 「NPMは特に寛大ではない？労働組合を結成しようとした職員が解雇―苦情」 。レジスター。 2024 年 5 月 11 日のオリジナルからアーカイブ。 2024 年 5 月 11 日に取得。
1 2 3 ウィリアムズ、クリス (2016 年 3 月 23 日)。 「いかにして一人の開発者が 11 行の JavaScript で Node、Babel、そして数千のプロジェクトを壊したか」レジスター。 2023年10月16日のオリジナルからアーカイブ。 2024 年 5 月 11 日に取得。
1 2 シャルマ、アックス（2022 年 7 月 27 日）。 「増加するプロテストウェア: 開発者が自分のコードを妨害する理由」テッククランチ。 2024 年 2 月 29 日のオリジナルからアーカイブ。 2024 年 5 月 11 日に取得。
1 2 3 4 マザイカ、ケン (2016 年 3 月 24 日)。 「いかにして 17 行のコードがシリコンバレーの注目のスタートアップ企業を潰したか」ハフポスト 。 2024 年 5 月 11 日のオリジナルからアーカイブ。 2024 年 5 月 11 日に取得。
1 2 3 4 5 ポール・ミラー（2016 年 3 月 24 日）。 「激怒した開発者が JavaScript を一時的に破壊した方法」 。ザ・ヴァージ。 2024 年 5 月 11 日のオリジナルからアーカイブ。 2024 年 5 月 11 日に取得。
1 2 3 4 ショーン、ギャラガー (2016 年 3 月 25 日)。 「激怒して辞める：コーダーが 17 行の JavaScript を非公開にし、「インターネットを破壊した」」アルステクニカ。 2024 年 5 月 11 日のオリジナルからアーカイブ。 2024 年 5 月 11 日に取得。
↑ トゥン、リアム（2016 年 3 月 23 日）。 「不満を抱いた開発者が何千もの JavaScript や Node.js アプリを破壊する」 。 ZDNET 。 2024 年 5 月 11 日のオリジナルからアーカイブ。 2024 年 5 月 11 日に取得。
1 2 意志

アイムズ、クリス（2016年3月29日）。 「『後悔はしていない』とJavaScriptのジェンガタワーを倒した男が語る - 開発者たちはこう尋ねる、我々はコーディングの仕方を忘れてしまったのだろうか？」 。レジスター。 2024 年 5 月 11 日のオリジナルからアーカイブ。 2024 年 5 月 11 日に取得。
↑ クラバーン、トーマス（2022 年 2 月 3 日）。 「マルウェアに感染した npm パッケージは、あなたが心配しているよりも一般的です。」レジスター。 2024 年 5 月 11 日のオリジナルからアーカイブ。 2024 年 5 月 11 日に取得。
Wayback Machine の npm 上の left-pad パッケージ (2015-09-22 アーカイブ)
「 https://en.wikipedia.org/w/index.php?title=Npm_left-pad_incident&oldid=1358977046 」から取得
カテゴリ : インターネットベースの活動
短い説明のある記事
短い説明はウィキデータとは異なります
Webarchive テンプレートのウェイバック リンク
このページは、2026 年 6 月 12 日の 06:41 (UTC) に最後に編集されました。
ページは Parsoid でレンダリングされました。
テキストは、クリエイティブ コモンズ 表示 - 継承 4.0 ライセンスに基づいて利用できます。
追加の条件が適用される場合があります。このサイトを使用すると、利用規約とプライバシー ポリシーに同意したことになります。 Wikipedia® は、非営利団体である Wikimedia Foundation, Inc. の登録商標です。

## Original Extract

Jump to content
Main menu
Main menu
move to sidebar
hide
Navigation
Main page
3
Aftermath
Toggle Aftermath subsection
3.1
Immediate effects
On March 22, 2016, programmer Azer Koçulu took down the left-pad package that he had published to npm (a package manager for JavaScript ). Koçulu deleted all his packages after a dispute with Kik Messenger , in which the company was granted control of the package name kik by npm, Inc. As a result, thousands of software projects that used left-pad as a dependency , including the Babel transcompiler and the React web framework , were unable to be built or installed. This caused widespread disruption, as technology corporations small and large, including Facebook , PayPal , Netflix , and Spotify , used left-pad as a dependency in their software products.
Several hours after the package was removed from npm, the company behind the platform, npm, Inc., manually restored the package. Later, npm disabled the ability to remove a package if more than 24 hours have elapsed since its publishing date and at least one other project depends on it. The incident drew media attention and reactions from people in the software industry . The removal of left-pad has prompted discussion regarding the intentional self- sabotage of software as acts of protest and brought attention to the elevated possibility of supply chain attacks in modular programming .
left-pad was a free and open-source JavaScript package published by Azer Koçulu, an independent programmer based in Oakland, California. [ 1 ] The package repetitively prepends characters to a string using a loop . [ 1 ] left-pad has been characterized as being extremely simple, [ 2 ] [ 3 ] consisting of only 11 lines of code in the final version authored by Koçulu, shown below: [ 4 ]
module . exports = leftpad ;
function leftpad ( str , len , ch ) {
str = String ( str );
var i = - 1 ;
ch || ( ch = ' ' );
len = len - str . length ;
while ( ++ i < len ) {
str = ch + str ;
}
return str ;
}
Koçulu published left-pad on npm , the default package manager for Node.js , a JavaScript runtime environment . [ 5 ] [ 2 ] Despite its relative obscurity, left-pad was heavily used; the package was used as a dependency by thousands of other software projects and was downloaded more than 15 million times before its removal. [ 6 ] [ 7 ] Some of the projects that required left-pad were critical to the JavaScript ecosystem at the time. This included Babel , a transcompiler that enables backwards-compatible JavaScript code; Webpack , a module bundling system; and React and React Native , frameworks widely used to develop websites and mobile apps , respectively. [ 8 ] [ 9 ] [ 1 ]
In addition to left-pad , Koçulu also owned kik on npm, a tool that allowed developers to set up templates for their projects. [ 1 ] On March 11, 2016, Kik Interactive, a Canadian company that owned the instant messaging platform Kik Messenger , asked Koçulu to relinquish control of the kik package because the company owned the "Kik" trademark . [ 10 ] Part of the correspondence included the following message from Kik:
We don't mean to be a dick about [the kik package], but it's a registered Trademark in most countries around the world and if you actually release an open source project called kik, our trademark lawyers are going to be banging on your door and taking down your accounts and stuff like that — and we'd have no choice but to do all that because you have to enforce trademarks or you lose them. Can we not come to some sort of a compromise to get you to change the name without involving lawyers? Is there something we could do for you in compensation to get you to change the name? [ 3 ]
Koçulu responded shortly after, refusing to change the name of his project, saying:
"}},"i":0}}]}'/> hahah, you're actually being a dick. so, fuck you. don't e-mail me back. [ 3 ]
Koçulu also requested US$30,000 as compensation "for the hassle of giving up with my pet project for [ sic ] bunch of corporate dicks". [ 1 ]
On March 18, 2016, Isaac Z. Schlueter, the chief executive officer of npm, Inc., wrote to Kik Interactive and Koçulu, stating that the ownership of the kik package would be manually transferred to Kik Interactive. [ 1 ]
After Koçulu expressed his disappointment with npm, Inc.'s decision and stated that he no longer wished to be part of the platform, Schlueter provided him with a command that would delete all 273 modules that he had registered. [ 10 ] Koçulu executed the command on March 22, 2016, removing every package he had previously released. [ 1 ] left-pad was one of the packages that was "unpublished", rendering it no longer publicly accessible on npm. [ 6 ] The left-pad software project and contents remained available on GitHub . [ 10 ]
Users attempting to build or install any JavaScript project that used left-pad as a dependency (including dependents such as Babel or Webpack) received a 404 error that caused the process to fail. [ 1 ] Among the software technology corporations that used the package were Meta Platforms , PayPal , Netflix , Spotify , [ 9 ] and Kik Interactive itself. [ 1 ]
An hour after he deleted the packages, Koçulu published a post on Medium ("I've Just Liberated My Modules"), explaining that he had unpublished his software projects from npm to protest corporate interests in free and open-source software. [ 1 ]
Soon after the deletion, other software developers began to post a flood of complaints, reactions, and workarounds on the project's Git issue tracking system . [ 8 ] [ 1 ]
Maintainers of open-source projects, including Babel, released hotfixes to remove the dependencies that Koçulu had unpublished. [ 8 ] Several of Koçulu's other package names were quickly taken over by newly published packages. [ 3 ] For example, another developer recreated the left-pad package—but released it as version 1.0.0. Since Koçulu published his as version 0.0.3, users continued to encounter problems. [ 3 ]
Around two hours after the original left-pad package was removed, npm manually restored the original 0.0.3 version by restoring a backup. [ 1 ] Laurie Voss, chief technology officer of npm, wrote that the company "picked the needs of the many" despite internal disagreements about whether the action was "the right call". [ 11 ]
npm changed its policy on the removal of published packages to prevent deletion if more than 24 hours have elapsed since its release date and at least one other project requires it as a dependency. [ 12 ] On behalf of npm, community manager Ashley Williams apologized for the disruption caused by the incident, stating that the platform "[failed] to protect the community". [ 12 ] Kik Interactive also apologized for the incident, with the company's head of messaging Mike Roberts publishing the email chain with Koçulu on Medium and characterizing his interaction as a "polite request". [ 9 ] Roberts wrote that they had initially reached out to Koçulu because they wished to publish an open-source package on npm with the name Koçulu was using. [ 6 ] Koçulu stated that he was sorry for disrupting others' work, but he believed he did it "for the benefit of the community in [ sic ] long term". [ 2 ]
The incident drew varied reactions from users on Twitter , GitHub , Reddit , and Hacker News , with many claiming that it briefly "broke the Internet". [ 2 ] [ 9 ] [ 10 ] [ 1 ] Many commented on the " move fast and break things " culture of JavaScript development, the unpredictable nature of open-source software, and a perceived over-reliance on modular programming . [ 2 ] [ 9 ] [ 3 ] Users also expressed disappointment regarding npm's decision to forcefully transfer Koçulu's package to Kik Interactive over a legal threat. [ 1 ]
The incident showed how the disruption of an npm package could lead to a supply chain attack . In addition to the widely publicized left-pad incident, a number of individuals had immediately hijacked Koçulu's other packages with unknown code after they were removed. [ 8 ] npm released a new policy to prevent malicious takeovers in similar disputes, [ 3 ] but the left-pad incident is still cited as an example of over-reliance on external contributors leading to an increased attack surface for software products. [ 13 ] Koçulu's intentional self-sabotage of left-pad has also been described as a precursor to incidences of protestware being published on platforms like npm. [ 7 ]
Hacktivism – Computer-based activities as a means of protest
peacenotwar – 2022 malware by Brandon Nozaki Miller
Software repository – Storage location for software packages
1 2 3 4 5 6 7 8 9 10 11 12 13 14 Collins, Keith (March 27, 2016). "How one programmer broke the internet by deleting a tiny piece of code" . Quartz . Archived from the original on May 11, 2024 . Retrieved May 11, 2024 .
1 2 3 4 5 Weinberger, Matt (March 23, 2016). "One programmer almost broke the internet by deleting 11 lines of code" . Business Insider . Archived from the original on 11 May 2024 . Retrieved 11 May 2024 .
1 2 3 4 5 6 7 Feldman, Brian (March 24, 2016). "One Man Deleted 11 Lines of Code From the Internet and Broke Hundreds of Apps" . Intelligencer . Archived from the original on May 11, 2024 . Retrieved May 11, 2024 .
↑ Koçulu, Azer (August 15, 2014). "left-pad/index.js" . GitHub . Archived from the original on 2024-05-11 . Retrieved 2026-03-02 .
↑ Claburn, Thomas (April 22, 2019). "NPM is Not Particularly Magnanimous? Staff fired after trying to unionize – complaints" . The Register . Archived from the original on May 11, 2024 . Retrieved May 11, 2024 .
1 2 3 Williams, Chris (March 23, 2016). "How one developer just broke Node, Babel and thousands of projects in 11 lines of JavaScript" . The Register . Archived from the original on October 16, 2023 . Retrieved May 11, 2024 .
1 2 Sharma, Ax (July 27, 2022). "Protestware on the rise: Why developers are sabotaging their own code" . TechCrunch . Archived from the original on February 29, 2024 . Retrieved May 11, 2024 .
1 2 3 4 Mazaika, Ken (March 24, 2016). "How 17 Lines of Code Took Down Silicon Valley's Hottest Startups" . HuffPost . Archived from the original on May 11, 2024 . Retrieved May 11, 2024 .
1 2 3 4 5 Miller, Paul (March 24, 2016). "How an irate developer briefly broke JavaScript" . The Verge . Archived from the original on May 11, 2024 . Retrieved May 11, 2024 .
1 2 3 4 Gallagher, Sean (March 25, 2016). "Rage-quit: Coder unpublished 17 lines of JavaScript and "broke the Internet" " . Ars Technica . Archived from the original on May 11, 2024 . Retrieved May 11, 2024 .
↑ Tung, Liam (March 23, 2016). "Disgruntled developer breaks thousands of JavaScript, Node.js apps" . ZDNET . Archived from the original on May 11, 2024 . Retrieved May 11, 2024 .
1 2 Williams, Chris (March 29, 2016). " 'No regrets' says chap who felled JavaScript's Jenga tower – as devs ask: Have we forgotten how to code?" . The Register . Archived from the original on May 11, 2024 . Retrieved May 11, 2024 .
↑ Claburn, Thomas (February 3, 2022). "Malware-infected npm packages more common than you may fear" . The Register . Archived from the original on May 11, 2024 . Retrieved May 11, 2024 .
left-pad package on npm at the Wayback Machine (archived 2015-09-22)
Retrieved from " https://en.wikipedia.org/w/index.php?title=Npm_left-pad_incident&oldid=1358977046 "
Categories : Internet-based activism
Articles with short description
Short description is different from Wikidata
Webarchive template wayback links
This page was last edited on 12 June 2026, at 06:41 (UTC) .
Page was rendered with Parsoid .
Text is available under the Creative Commons Attribution-ShareAlike 4.0 License ;
additional terms may apply. By using this site, you agree to the Terms of Use and Privacy Policy . Wikipedia® is a registered trademark of the Wikimedia Foundation, Inc. , a non-profit organization.
