---
source: "https://www.gamingonlinux.com/2026/05/flathub-moves-to-ban-nearly-all-apps-and-submissions-made-with-generative-ai/"
hn_url: "https://news.ycombinator.com/item?id=48330069"
title: "Flathub bans AI-generated apps and submissions"
article_title: "Flathub moves to ban nearly all apps and submissions made with generative AI | GamingOnLinux"
author: "Lihh27"
captured_at: "2026-05-30T11:37:15Z"
capture_tool: "hn-digest"
hn_id: 48330069
score: 4
comments: 0
posted_at: "2026-05-29T22:20:11Z"
tags:
  - hacker-news
  - translated
---

# Flathub bans AI-generated apps and submissions

- HN: [48330069](https://news.ycombinator.com/item?id=48330069)
- Source: [www.gamingonlinux.com](https://www.gamingonlinux.com/2026/05/flathub-moves-to-ban-nearly-all-apps-and-submissions-made-with-generative-ai/)
- Score: 4
- Comments: 0
- Posted: 2026-05-29T22:20:11Z

## Translation

タイトル: Flathub、AI生成アプリと投稿を禁止
記事のタイトル: Flathub は、生成 AI を使用して作成されたほぼすべてのアプリと投稿を禁止する予定です |ゲームオンLinux
説明: Linux でアプリケーションを取得する最も一般的な方法の 1 つである Flathub には、新しく更新された生成 AI ポリシーがあります。

記事本文:
この Web サイトでは、ブラウジング エクスペリエンスを向上させ、追加機能を提供するために Cookie を使用します -> 詳細 Cookie を拒否 - Cookie を許可
ログイン
Flathubは、生成AIを使用して作成されたほぼすべてのアプリと投稿を禁止する方向に動く
Linux 上でアプリケーションを取得する最も一般的な方法の 1 つである Flathub には、新しく更新された生成 AI ポリシーがあり、ほぼすべてが禁止されています。ただし、「成熟し、よく管理されているプロジェクト」については例外が記載されていますが、それを保証するものではありません。
新しいコミットが送信され、ドキュメントにマージされました。ドキュメントは現在公開されており、「許可されていないことを明確にするために LLM ポリシーを書き換える」と記載されています。新しいポリシーには次のように書かれています。
生成型 AI ポリシー
このポリシーは、Flathub に送信されるアプリケーションと、マニフェスト、メタデータ、パッチ、ビルド スクリプト、プル リクエストを含む Flathub 送信自体の両方に適用されます。このポリシーの目的上、アプリケーションには、BaseApp、拡張機能、および flatpak-builder によって生成できるその他のアーティファクトが含まれます。
送信プル リクエストは、AI ツールやエージェントを使用して生成、オープン、または自動化してはなりません。また、提出した PR で AI ツールによるレビューをリクエストしないでください。 GitHub の自動コパイロット レビューは、送信者がここにアクセスしてリポジトリ アクセスを変更してリポジトリを除外するか、ここにあるグローバルな「自動コパイロット コード レビュー」を無効にすることで無効にできます。
AI で生成されたコード、AI 支援のコード、ドキュメント、その他のコンテンツを含むアプリケーションは許可されません。
著作権で保護されたコード、ライセンスに互換性のないコード、または倫理的に問題のあるコードを含むアプリケーションまたは変更は許可されません。
これらの提出は、さらなる審査なしに拒否される場合があります。
これらのポリシーに繰り返し違反すると、今後の投稿や活動が永久に禁止される場合があります。
例外が認められる場合があります

成熟し、よく管理されたプロジェクト。
もう少し背景を説明すると、開発者の Bart Piotrowski 氏は、Mastodon のソーシャル メディア投稿で次のように述べています。
Flathub の LLM ポリシーを更新し、提出プロセスと提出されるアプリケーションの両方で AI の使用を明示的に禁止しました。
https://github.com/flathub-infra/documentation/commit/992f57b30de98ddbd5e80959e9672998c83c8c97
私はそれについていくつかの懸念を持っていたので、そのコミット前の文言は比較的穏やかでした。これが Fediverse に関して不人気な意見であることは承知していますが、LLM は避けられないと思います。現実には、時間の経過とともにコードが自然に成長しなくなることが予想されます。これは FOSS の内外で役立つツールになると信じています。私は、作者がエージェントに指示する以上の努力をしたアプリがもっと多く出てくることを期待していました。一方で、あたかもそれを拒否している私たち愚か者に素晴らしいソフトウェアを授けているかのように振る舞う資格のある提出者と私が経験した不快なやり取りの数は、先月で一気になくなりました。私は疲れている。
いつものように、これを遡及的に適用することはありません。そのため、すでに公開されているバイブコード化されたアプリは引き続き利用可能です。
これについてどう思いますか？議論のどちらの側にいても、それに関するルールを明確に定義し、誰にとっても明確にすることは良いことです。
一方で、あたかもそれを拒否している私たち愚か者に素晴らしいソフトウェアを授けているかのように振る舞う資格のある提出者と私が経験した不快なやり取りの数は、先月で一気になくなりました。
「このポリシーの動機となった懸念は変わっておらず、正確に述べておく価値があります。DCO は、提出者がコードを提供する法的権利を持っているかどうかに関するものであり、「創造的な表現」に関するものではありません。LLM 出力の著作権とライセンスのステータスは未解決のままです。

その質問はまだ未解決です。変化したのはリスクのバランスです。
- AI支援コンテンツを受け入れるプロジェクトは、これまでのところ深刻な法的問題に遭遇していない。これは、リスクが現実化する可能性が高くないことを示唆している。
- Red Hat[1] などの他の組織は、リスクは許容範囲内であると評価していますが、個人開発者のコミュニティには企業の法的支援がなく、根拠のない紛争が発生した場合でも、QEMU の作業が長期にわたって妨げられる可能性があります。」
ソース: [QEMU 開発リスト](https://lists.nongnu.org/archive/html/qemu-devel/2026-05/msg07614.html)
あたかも今目の前にいるかのように、怒っているメンテナーが見えます。
引用: Emeric > あたかも今目の前にいるかのように、怒っているメンテナが見えます。
その気持ちはよくわかりますが、問題は、Flathub はアプリ ストアであるということです。 Linux警察ではありません。
彼らはますます門番としての役割を果たしていますが、私は多くのアプリストアに公開しているのに、なぜFlathubの方がエクスペリエンスがさらに悪いと感じるのでしょうか?
アプリストアは、ルールを設けず何でも許可することを望まない限り、門番の役割を果たしますが、法的にそれを行うことはできません。彼らは何らかの方法で門番をしなければならず、それが彼らの仕組みです。
引用：エメリック アプリストアは、ルールを設けず何でも許可することを望まない限り、門番です。しかし、法的にそれを行うことはできません。彼らは何らかの方法で門番をしなければならず、それが彼らの仕組みです。
引用: Liam Squires-Hand 生成 AI を禁止することが、他のどこよりも「悪い」こととどのような関係があるのかわかりません。
著作権で保護されたコード、ライセンスに互換性のないコード、または倫理的に問題のあるコードを含むアプリケーションまたは変更は許可されません。
引用: CatKiller 著作権で保護されたコード、ライセンス互換性のないコード、または倫理的に問題のあるコードを含むアプリケーションまたは変更は許可されません。
引用: CatKiller アプリケーション

または、著作権で保護されたコード、ライセンスに互換性のないコード、または倫理的に問題のあるコードを含む変更は許可されません。
引用：マニュアルをお読みください そして「倫理的に疑わしい」。それが具体的に何を指しているのかさえ分かりません。
Linux とオープンソースの年齢チェックの免除は問題になる可能性がある
スチームデッキの在庫は戻りますが、価格は大幅に上昇しています
Proton-CachyOS は低遅延レイヤーと Discord のリッチ プレゼンス サポートを追加します
Linux Mint はファイルマネージャーの改善、新しいスクリーンショットツールなどを追加
アンリアル エンジン 6 がロケット リーグのメジャー アップグレードとともに明らかに - ティーザーが気になる
Steam コントローラー LED とその他の LED を暗くできるようになりました。
39 分前 - Cyba.Cowboy
Proton Experimental が Subnautica 2、War Th… の修正を取得
39 分前 - chr
Epic Games CEO Tim Sweeney が Valve / Gabe に影を投げかける…
55 分前 - ヴィックベイ
Flathubは、ほぼすべてのアプリと提出物を禁止する予定です…
1時間前 - TheSHEEEP
Flathubは、ほぼすべてのアプリと提出物を禁止する予定です…
1時間前 - TheSHEEEP
ターミナルトリック - タスクマネージャーの進行状況インジケーター…
1 日前 - シュメル
フィードバックが必要です - 今後のウェブサイトの更新
1日前 - グリジ
Mac コンピューターは良好で安定していますか?
2日前 - ロジンブー
最近何して遊んでますか？ - 5月17日版…
3日前 - PlayingOnLinuxphone
アンチチートページの更新
4日前 - リアム・スクワイアズ・ハンド

## Original Extract

Flathub, one of the most popular ways to grab applications on Linux, has a newly updated generative AI policy - where it

This website makes use of cookies to enhance your browsing experience and provide additional functionality -> More info Deny Cookies - Allow Cookies
Login
Flathub moves to ban nearly all apps and submissions made with generative AI
Flathub , one of the most popular ways to grab applications on Linux, has a newly updated generative AI policy - where it's pretty much all banned. However, there is an exception noted for "mature, well-maintained projects" but it's not a guarantee.
A new commit was sent in and merged into the documentation, which is live now, that notes "Reword LLM policy to make it clear it's not allowed". The new policy reads:
Generative AI policy
This policy applies to both the application being submitted to Flathub and the Flathub submission itself, including the manifest, metadata, patches, build scripts, and pull request. For the purpose of this policy, applications include BaseApps, extensions, and any other artifacts that can be produced by flatpak-builder.
Submission pull requests must not be generated, opened, or automated using AI tools or agents. Please also do not request review from any AI tools in the submission PR. Automated Copilot reviews on GitHub can be disabled by the submitter by going here and changing Repository access to exclude the repo or disabling the global "Automatic Copilot code review" found here .
Applications containing AI-generated or AI-assisted code, documentation, or other content are not allowed.
Applications or changes containing copyrighted, license-incompatible, or ethically questionable code are not allowed.
These submissions can be rejected without any further review.
Repeatedly violating these policies may result in a permanent ban from future submissions and activities.
Exceptions may be granted for mature, well-maintained projects.
To give some more context, developer Bart Piotrowski mentioned in a social media post on Mastodon :
We have updated Flathub's LLM policy to explicitly disallow AI usage for both the submission process and applications being submitted.
https://github.com/flathub-infra/documentation/commit/992f57b30de98ddbd5e80959e9672998c83c8c97
I've had some reservations about it, so the wording before that commit was relatively milder. I know it's an unpopular opinion on the Fediverse, but I do think LLMs are inevitable, and the reality is that you can expect less organically grown code as time goes on. I believe it can be a useful tool in and outside FOSS; I hoped we will see a larger number of apps where authors made some effort beyond prompting an agent. Meanwhile, the number of unpleasant interactions I've had with entitled submitters acting as if they were bestowing their brilliant software upon us idiots who are rejecting it went through the roof in the last month. I'm tired.
As always, we are not applying this retroactively, so any vibecoded apps which were already published will remain available.
What are your thoughts on this? No matter which side of the argument you're on, having clearly defined rules around it is a good thing so that it's clear for everyone.
Meanwhile, the number of unpleasant interactions I've had with entitled submitters acting as if they were bestowing their brilliant software upon us idiots who are rejecting it went through the roof in the last month.
"The concern that motivated the policy is unchanged, and it is worth stating precisely: the DCO is about whether the submitter has the legal right to contribute the code, not about "creative expression". The copyright and license status of LLM output remains unsettled, so that question is still open. What has shifted is the balance of risk:
- projects accepting AI-assisted content have not run into serious legal trouble so far, which suggests the probability of the risk materializing is not high;
- other organizations, such as Red Hat[1], have assessed the risk as acceptable -- though a community of individual developers does not have the legal backing of a company, and even an unfounded dispute would be a long-lasting distraction from work on QEMU."
source: [QEMU dev list](https://lists.nongnu.org/archive/html/qemu-devel/2026-05/msg07614.html)
I can see the fuming maintainers as if they're in front of me right now.
Quoting: Emeric > I can see the fuming maintainers as if they're in front of me right now.
While I completely understand the sentiment, here is the thing: Flathub is an app store. Not the Linux police.
They are acting more and more as gatekeepers, and I while I publish on many app stores, why do I feel Flathub is just the worse experience?
An app store is a gatekeeper, unless they want to have no rules in place and allow anything - but they cannot legally do that. They have to gatekeep in some way, that's just how they work.
Quoting: Emeric An app store is a gatekeeper, unless they want to have no rules in place and allow anything - but they cannot legally do that. They have to gatekeep in some way, that's just how they work.
Quoting: Liam Squires-Hand I don't see what banning generative AI has to do with them being "worse" than anywhere else.
Applications or changes containing copyrighted, license-incompatible, or ethically questionable code are not allowed.
Quoting: CatKiller Applications or changes containing copyrighted, license-incompatible, or ethically questionable code are not allowed.
Quoting: CatKiller Applications or changes containing copyrighted, license-incompatible, or ethically questionable code are not allowed.
Quoting: pleasereadthemanual And "ethically questionable". I'm not even sure what that's referring to specifically.
Linux and open source getting age checking exemptions could be problematic
Steam Deck stock returns but there's a big price increase
Proton-CachyOS adds low latency layer and Discord rich presence support
Linux Mint gets file manager improvements, a new screenshot tool and more
Unreal Engine 6 revealed with a major Rocket League upgrade - the teaser concerns me
You can now dim the Steam Controller LED plus other Ste…
39 minutes ago - Cyba.Cowboy
Proton Experimental gets fixes for Subnautica 2, War Th…
39 minutes ago - chr
Epic Games CEO Tim Sweeney throws shade at Valve / Gabe…
55 minutes ago - vic-bay
Flathub moves to ban nearly all apps and submissions ma…
1 hour ago - TheSHEEEP
Flathub moves to ban nearly all apps and submissions ma…
1 hour ago - TheSHEEEP
Terminal trick - progress indicator in the task manager…
a day ago - Shmerl
Feedback needed - future website updates
a day ago - grigi
Are Mac computers good and stable?
2 days ago - rojimboo
What have you been playing recently? - 17th May edition…
3 days ago - PlayingOnLinuxphone
Anti-Cheat page updates
4 days ago - Liam Squires-Hand
