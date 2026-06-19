---
source: "https://www.theregister.com/software/2026/03/26/linux-kernel-czar-says-ai-bug-reports-arent-slop-anymore/5226256"
hn_url: "https://news.ycombinator.com/item?id=48598086"
title: "Linux Maintainer Greg Kroah-Hartman Says AI Tools Now Useful, Finding Real Bugs"
article_title: "Linux kernel czar says AI bug reports aren't slop anymore"
author: "root-parent"
captured_at: "2026-06-19T13:17:56Z"
capture_tool: "hn-digest"
hn_id: 48598086
score: 3
comments: 1
posted_at: "2026-06-19T13:01:52Z"
tags:
  - hacker-news
  - translated
---

# Linux Maintainer Greg Kroah-Hartman Says AI Tools Now Useful, Finding Real Bugs

- HN: [48598086](https://news.ycombinator.com/item?id=48598086)
- Source: [www.theregister.com](https://www.theregister.com/software/2026/03/26/linux-kernel-czar-says-ai-bug-reports-arent-slop-anymore/5226256)
- Score: 3
- Comments: 1
- Posted: 2026-06-19T13:01:52Z

## Translation

タイトル: Linux メンテナの Greg Kroah-Hartman 氏、AI ツールが便利になり、本当のバグを発見できると語る
記事のタイトル: Linux カーネルの最高責任者、AI バグレポートはもういい加減ではないと語る
説明: インタビュー: グレッグ・クローア・ハートマン

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
AIのバグレポートは一夜にしてジャンクから正規のものになったとLinuxカーネルのトップが語る
グレッグ・クロー・ハートマンは変曲点を説明できないが、速度が低下したり消え去ったりするわけではない
インタビュー 今週、私は KubeCon Europe の記者向け昼食会に出席していましたが、驚いたことに、私の隣に座るべきなのは、長年 Linux カーネルのメンテナを務めている Greg Kroah-Hartman でした。最近オランダに住んでいるグレッグは、AI、Linux、セキュリティについて簡単にコメントするためにそこにいました。私たちは、先月にわたって、Linux のセキュリティとコード レビューに関する AI 主導の活動が、オープンソースの世界の誰も予想していなかった方法で「本当に飛躍」したことについて話しました。
「数カ月前、私たちは『AIスロップ』と呼ぶ、明らかに間違っているか低品質なAI生成のセキュリティレポートを入手していました」と同氏は述べた。 「ちょっと面白かったです。私たちはあまり心配しませんでした。」もちろん、Linux カーネルのメンテナーはたくさんいるので、彼らにとって AI のスロップは、たとえば cURL の創設者兼主任開発者である Daniel Stenberg にとってほど負担ではありません。AI スロップの報告により、cURL チームはバグ報奨金の支払いを停止しました。
Linus Torvalds とその友人たちが語る Reghow Linux のソロ活動が世界的なジャム セッションになった
状況は変わった、とクロア＝ハートマン氏は語った。 「1か月前に何かが起こり、世界が変わりました。今、私たちは本当の報告をしています。」 Linux だけではない、と彼は続けた。 「すべてのオープンソース プロジェクトには AI で作成された本物のレポートが含まれていますが、それは優れたものであり、本物です。」主要なオープンソース プロジェクトのセキュリティ チームは非公式かつ頻繁に話し合い、全員が同じ変化を感じていると同氏は指摘しました。 「すべてのオープンソース セキュリティ チームが現在これに取り組んでいます。」
その背後に何があるかは誰にもわかりません。何が変わったか尋ねると

編集者のクロア＝ハートマン氏は、「分からない。その理由は誰も分からないようだ。より多くのツールがより良くなったか、人々が『さあ、これを調べ始めよう』と言い始めたかのどちらかだ」と率直に語った。たくさんの異なるグループ、異なる会社のようです。」明らかなのはスケールです。 「カーネルに関しては、私たちはそれを処理できます」と彼は言いました。
「私たちははるかに大規模なチームで、非常に分散しており、その増加は本物であり、その勢いは衰えていません。これらは小さなことであり、大きなことではありませんが、すべてのオープンソース プロジェクトについて、この問題について支援が必要です。」同氏は、小規模なプロジェクトでは、AI が生成したもっともらしいバグ レポートやセキュリティ調査結果の突然の洪水を吸収する能力がはるかに低いことをほのめかしました。少なくとも現時点では、それらはゴミではなく本物のバグです。
舞台裏では、セキュリティ チームがメモを比較しています。 「私たちは非公式に集まり、よく話し合います。なぜなら、私たちは皆同じ問題を抱えているからです」と彼は言いました。 「ツールのどこかに変曲点があったに違いありません。ローカルツールは改善されましたか？人々は何かを見つけましたか？正直、わかりません。」
今のところ、AI は Linux カーネル コードの完全な作成者としてよりも、レビュー担当者やアシスタントとしての役割を果たしていますが、その境界線は曖昧になり始めています。 Kroah-Hartman 氏はすでに、AI が生成したパッチを使用した独自の実験を行っています。
「本当に愚かなプロンプトをしてしまった」と彼は振り返った。 「『これをください』と言うと、『これが私が見つけた60個の問題とその修正です』という60個を吐き出しました。約 3 分の 1 は間違っていましたが、それでも比較的現実的な問題を指摘しており、パッチの 3 分の 2 は正しかったです。」念のため言っておきますが、これらの作業パッチには人間によるクリーンアップ、変更ログの改善、統合作業がまだ必要でしたが、決して役に立たなかったのです。 「ツールは優れています」と彼は言いました。 「この問題を無視することはできません。問題は起きつつあり、さらに良くなりつつあります。」
開発者は AI の役割を認識し始めています

実際の提出物。 「いくつかのパッチが生成されているのを確認しています」とKroah-Hartman氏は語った。 「そのための小さな共同開発タグができました。いくつかの新機能については検討中ですが、レビューでは主に AI が使用されていることがわかります。」
単純な変更に関する作業のほとんどが AI によって行われるような近い将来を想像できるかとの質問に対し、同氏は、「単純で小さなエラー状態、エラー状態を適切に検出する」場合、AI は現在すでに数十の使用可能なパッチを生成できると述べた。
AI によって生成されたレポートや AI 支援による作業が急増したことにより、カーネル独自のレビュー インフラストラクチャに AI を組み込むという並行した推進にも拍車がかかっています。その重要な部分は、もともと Google で開発され、現在は Linux Foundation に寄付されているツールである Sashiko です。
「負荷を軽減する方法で提供されるこれらのパッチの一部を簡単にレビューできる方法が必要です。」同氏によると、このツールは「ほぼすべてのカーネルパッチ上で稼働している」という。 「これは一般に公開されています。私たちはこれをレビュー ツールに統合しています。誰でも使用できるようにしています。」
ナニー州がLinuxを発見、起動前に子供のIDチェックを要求
オープンソース開発者は、ダウンロードごとに豚にお金を支払わせることを検討している
仕事中毒のオープンソース開発者は休憩を取る必要がある
さて、Anthropic の AI は C コンパイラを構築しました。それはあまり印象に残らない
この取り組みは、特定のサブシステム内での以前の取り組みに基づいて構築されています。 「ネットワーキングと BPF の担当者は、しばらくの間、LLM によって生成されたレビューを行ってきました」と Kroah-Hartman 氏は言います。 「ダイレクト レンダリング マネージャー (DRM) の担当者と現在 Google のツールが、これらすべてを 1 つの共通インターフェースにまとめています」と同氏は説明しました。 「さまざまなサブシステムが、より優れたスキルやプロンプトを追加しています。ストレージについては、探す必要があるものがあります。グラフィックについては、探す必要があるものがあります。人々は、

そのための公共の場、それはそうあるべきです。これはとても良いですね。」
Kroah-Hartman 氏は、AI ベースのレビュー ワークフローの先駆者として、長年カーネル開発者を務め、現在は Meta に所属する Chris Mason 氏の功績を認めました。 Mason はしばらくの間、eBPF とネットワーキングの AI レビューを実行してきました。 systemd プロジェクトも、すべて C コードベースに同じクラスのツールを使用しています。
AIの査読者は権威あるものではなく、付加的なものであると彼は強調した。 「レビューの面では、いくつかの良いレビューを生み出しています。すべてがわかるわけではありません。いくつかの点はまだ間違っています。しかし、それは多くの明白な事柄を指摘している」と彼は言った。
当面の最大の利点の 1 つは、所要時間です。 AI レビュー担当者が明らかな問題にフラグを立てた場合、人間のメンテナーが実際にパッチを読むずっと前に、提出者はフィードバックを受け取ります。 「何かに反応しているのを見ると、メンテナがするよりも早く送信者にフィードバックを返してくれます。これは素晴らしいことです」と Kroah-Hartman 氏は言いました。 「パッチ上でそのまま動作するボットが多数あります。それらが失敗するのを見たら、メンテナとしてそれを見る必要さえないとわかります。そして、開発者に『ああ、明日は別のバージョンを開発できる』という気持ちを与え、フィードバックを少しだけ増やすのに役立ちます。」
それでも、AI によって生成されたレポートやパッチが増加するにつれて、レビューの負担も増加します。 「それはより多くのレビューです。 「カーネルに関してレビューしなければならないことのほうが多いのです。」と彼は言いました。だからこそ、OpenSSF とその Alpha-Omega プログラムの取り組みが重要なのです。「私たちは、メンテナーがこの受信フィードを処理し、対処するのを容易にするツールの作成に取り組んでいます。」
クローア・ハートマンにとって繰り返しのテーマは、アクセスの公平性です。最近まで、大規模な AI ツールを実行できるのは、十分なリソースを備えたサブシステムだけでした。 Google のレビュー システムを Linux Foundation プロジェクトに変えることは、この状況を変えることを目的としています。
「それはこれです

「これは、私たちが LF プロジェクトとして、どのようにして全員にアクセスを提供しているかを示す一例としてのツールの 1 つです。」と彼は言いました。以前は、バックエンドで実行するためのリソースを持つのはサブシステムだけでした。現在、私たちはそれを全員に提供しています。」カーネル自身のインフラストラクチャを超えて使用できるようにする作業はすでに進行中です。
これが重要なのは、Kroah-Hartman 氏が強調し続けているように、AI の波は単なるカーネルの問題ではないためです。 「すべてのオープンソース プロジェクトには、AI を使用して作成された実際のレポートがあります」と彼は言いました。 「私たちの増加は本物であり、その勢いは衰えていません。これらは大したことではありませんが、すべてのオープンソース プロジェクトについては支援が必要です。」
Linux の場合、AI との関係はすでに過去の理論を進化させ、実践に移しています。それは混合の祝福です。 AI は同時に、実際の脆弱性の新たな発生源でもあり、それらに対処しなければならない人間のレビュー担当者に負担をかけると同時に、その負担の管理にも役立ちます。
Kroah-Hartman 氏とその仲間たちにとっての秘訣は、オープンソースのメンテナーを溺れさせることなく、AI を戦力の増強手段として維持することだろう。 ®
セキュリティ
テキサスではすべてがより大きく、より優れています - データ侵害も含めて
狩猟・漁業許可事件で300万人の住民が逮捕される
Vercel がオープンソース エージェント フレームワークをデビュー前に、Passport でシャドウ AI の修正を試みる
Vercel 経由で AWS を間接的に使用する場合のコストプレミアムは、コンピューティング リソースのより効率的な使用により軽減される、と CTO は主張
ZTE と広東省中国電信がクロスベンダー IP ネットワーク シミュレーションのパイロットを推進し、インテリジェントなネットワーク運用への道を開く
パートナーコンテンツ: 95% を超えるデジタルツインの忠実度およびマルチベンダーのコラボレーションを活用して、ネットワーク変更のリスクを排除し、エラーゼロの O&M を達成します。
Microsoft の最新の Windows バグはごみ箱にあります
ファイル削除ダイアログで内部の意味不明な認識可能な名前が置き換えられる
PA

AS および IAAS
Graviton 5 は感動的ですが、すべての神聖なるものへの愛を込めて、これを「AI チップ」と呼ぶのはやめてください
AWS は口よりもチップ工場の運営のほうが上手い
英国のプライバシー監視機関、「判断が不十分」と認めたため辞任
ジョン・エドワーズ氏、不適切なユーモアの試みを含む行為の捜査を受けて自身の立場が「維持できなくなった」と語る
セキュリティ
研究者によると、脱獄ではなく単に「このコードを修正してください」というプロンプトが表示されただけで、連邦当局は寓話5に激怒したとのこと
オンプレミス
アマゾンは昨年自社のビット納屋で最大25億ガロンの水を使用している
科学
AI とブレイン コンピューター インターフェイスにより、言葉を話すことのできない ALS 患者がフルタイムで働けるようになります
公共部門
キャピタは公務員年金制度修正の期限を過ぎて航行しようとしている
仮想化
Tesco は急速な移行リスクにもかかわらず、VMware と Broadcom からの撤退に全力を注いでいます
Vercel 経由で AWS を間接的に使用する場合のコストプレミアムは、コンピューティング リソースのより効率的な使用により軽減される、と CTO は主張
活動家らは、テクノロジーは使用が計画されている境界で子供と大人を確実に区別することができないと主張している
サンフランシスコがホスティング会社の Localhost カンファレンスの主催者となる
基礎となるテクノロジーは本物です...そして、会社が言及しなかったパートナーから借用したものです
エドからコーリーへのメモ: 安全な場合は 1 回まばたきをし、危険な場合は 2 回まばたきをしてください
プラス：米国はイランのプロパガンダサイトを削除。マーケティング会社が「なぜあなたの情報を入手するのですか？」と尋ねます。さらに！
プラス：中国はスマートフォン監視ツールをアップグレード。リングは覗き見防止の姿勢を緩和します。などなど
ジェフ・モス氏によると、投票村のレポートは非常に成功しており、今後は DEF CON 全体が含まれることになる
会社全体の評価額は35億ドル以上に相当するが、売却部分は特定されていない
プラスの面としては、情報セキュリティは長く安定したキャリアを築くのに適しています。
フランスのデジタル主権推進は苦戦している

マイクロソフトの重力井戸から逃れるために
Nextcloud のロールアウトは、ローカルで制御されるストレージが 1 つのことであることを示しています。ユーザーを Office から解放することはまったく別のことです
CentOS の歴史: 生化学者の Linux 趣味プロジェクトがどのようにして企業世界のデフォルトのオペレーティング システムになったのか
Red Hat が Windows が「おそらく適切な製品」であると述べた後、コミュニティが団結したとき
Netflix のウィズが AI 費用を削減するアプリを作成し、それをオープンソース化
Project Headroom は多額の費用も節約できる
OpenBSD 7.9 が登場、あらゆる鋭いエッジを誇るダイヤモンドの原石
60 番目のリリースでは、その禁欲的な努力を失うことなく、より多くのコア、遅延休止状態、および基本的な Wi-Fi 6 が追加されています
Fedora: Microsoft は全力で取り組んでいるが、Deepin は見捨てられている
Red Hat の無料ディストリビューションはデスクトップを失うが、重要な新しい友人を作る
LocalSend はスニーカーネットを廃業にします
Apple ロックインを除いた AirDrop と同様
お問い合わせ
私たちと一緒に宣伝しましょう
私たちは誰なのか
ニュースレター
次のプラットフォーム
開発クラス
ブロックとファイル
状況出版
クッキーポリシー
プライバシーポリシー
利用規約
私の個人情報を共有しないでください
同意のオプション
著作権。すべての著作権は © 1998-2026 に留保されます。

## Original Extract

Interview: Greg Kroah-Hartman can

Jump to main content
Search
TOPICS
Special Features
All Special Features
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
AI bug reports went from junk to legit overnight, says Linux kernel czar
Greg Kroah-Hartman can't explain the inflection point, but it's not slowing down or going away
INTERVIEW I was at a press luncheon at KubeCon Europe this week when, to my surprise, who should sit down next to me but long-term Linux kernel maintainer Greg Kroah-Hartman. Greg, who lives in the Netherlands these days, was there to briefly comment on AI, Linux, and security. We spoke about how, over the last month, AI-driven activity around Linux security and code review has "really jumped" in a way no one in the open source world saw coming.
"Months ago, we were getting what we called 'AI slop,' AI-generated security reports that were obviously wrong or low quality," he said. "It was kind of funny. It didn't really worry us." Of course, there are many Linux kernel maintainers, so for them, AI slop isn't as burdensome as it is for, say, Daniel Stenberg, founder and lead developer of cURL, where AI slop reports caused the cURL team to stop paying bug bounties .
Linus Torvalds and friends tellThe Reghow Linux solo act became a global jam session
Things have changed, Kroah-Hartman said. "Something happened a month ago, and the world switched. Now we have real reports." It's not just Linux, he continued. "All open source projects have real reports that are made with AI, but they're good, and they're real." Security teams across major open source projects talk informally and frequently, he noted, and everyone is seeing the same shift. "All open source security teams are hitting this right now."
No one is quite sure what's behind it. Asked what changed, Kroah-Hartman was blunt: "We don't know. Nobody seems to know why. Either a lot more tools got a lot better, or people started going, 'Hey, let's start looking at this.' It seems like lots of different groups, different companies." What is clear is the scale. "For the kernel, we can handle it," he said.
"We're a much larger team, very distributed, and our increase is real – and it's not slowing down. These are tiny things, they're not major things, but we need help on this for all the open source projects." Smaller projects, he implied, have far less capacity to absorb a sudden flood of plausible AI-generated bug reports and security findings – at least now they're real bugs and not garbage ones.
Behind the scenes, security teams are comparing notes. "We get together informally and talk a lot, because we all have the same problems," he said. "There must have been some inflection point somewhere with the tools. Did the local tools get better? Did people figure out something? I honestly don't know."
For now, AI is showing up more as a reviewer and assistant than as a full author of Linux kernel code, but that line is starting to blur. Kroah-Hartman has already done his own experiments with AI-generated patches.
"I did a really stupid prompt," he recounted. "I said, 'Give me this,' and it spit out 60: 'Here's 60 problems I found, and here's the fixes for them.' About one-third were wrong, but they still pointed out a relatively real problem, and two-thirds of the patches were right." Mind you, those working patches still needed human cleanup, better changelogs, and integration work, but they were far from useless. "The tools are good," he said. "We can't ignore this stuff. It's coming up, and it's getting better."
Developers are starting to acknowledge AI's role in actual submissions. "We're seeing some patches being generated," Kroah-Hartman said. "You have a little co-develop tag for that now. We're seeing some things for some new features, but we're seeing AI mostly being used in the review."
Asked whether he could imagine a near-future where most of the work on simple changes comes from AI, he said that for "simple little error conditions, properly detecting error conditions," AI could already generate dozens of usable patches today.
The sudden increase in AI-generated reports and AI-assisted work has also spurred a parallel push to build AI into the kernel's own review infrastructure. A key piece of that is Sashiko, a tool originally developed at Google and now donated to the Linux Foundation .
"We need to be able to have an easy way to review some of these patches that come in ways that cut down on our load." The tool is "out there, running on almost all kernel patches," he said. "You can see it publicly. We're integrating it into our review tools. It's available for anybody to use."
Nanny state discovers Linux, demands it check kids' IDs before booting
Open source devs consider making hogs pay for every download
Workaholic open source developers need to take breaks
OK, so Anthropic's AI built a C compiler. That don't impress me much
That work builds on earlier efforts inside specific subsystems. "The networking and the BPF people have been doing LLM-generated reviews for a while," said Kroah-Hartman. "The Direct Rendering Manager (DRM) people and now Google's tool are pulling all those into one common interface," he explained. "Different subsystems are adding better skills or prompts – for storage, here are the things you need to look for; for graphics, here are the things you need to look for. People are contributing in a public place for that, which is how it should be. This is very good."
Kroah-Hartman credited longtime kernel developer Chris Mason, now at Meta, with pioneering AI-based review workflows. Mason has been running AI review for eBPF and networking for some time. The systemd project is also using the same class of tools for its all-C codebase.
AI reviewers, he stressed, are additive rather than authoritative. "On the review side, it's generating some good reviews. It doesn't get you everything. Some things are still wrong. But it does point out a lot of the obvious things," he said.
One of the biggest immediate wins is turnaround time. When an AI reviewer flags obvious problems, submitters get feedback long before a human maintainer would realistically read the patch. "If I see it respond to something, it gives feedback to the submitter faster than the maintainer had a chance to, which is nice," Kroah-Hartman said. "We have a number of bots that run on patches as it is. If I see those fail, I just know I don't even need to look at that as a maintainer. And it gives the developer, 'Oh, I can go do another version tomorrow,' which helps increase the feedback a little better."
Still, as AI-generated reports and patches grow, so does the review burden. "It's more reviews; it's more stuff we have to review for the kernel," he said. That's why efforts with the OpenSSF and its Alpha-Omega program matter. "We're working to try and create tools to help make it easier for maintainers to handle this incoming feed and deal with it."
A recurring theme for Kroah-Hartman is equity of access. Until recently, only well-resourced subsystems could afford to run heavy AI tooling at scale. Turning Google's review system into a Linux Foundation project is meant to change that.
"That's this one tool that we have for the review," he said. "It's one tool as an example of how now, as an LF project, we're giving access to everybody. Before, it was just the subsystems that had the resources to run it on the back end. Right now, we're giving it to everyone." Work is already underway to make it usable beyond the kernel's own infrastructure.
That matters because, as Kroah-Hartman keeps emphasizing, the AI wave is not just a kernel problem. "All open source projects have real reports that are made with AI," he said. "Our increase is real, and it's not slowing down. These aren't major things, but we need help on this for all the open source projects."
For Linux, the relationship with AI is already evolving past theory and into practice. It's a mixed blessing. AI is simultaneously a new source of real vulnerabilities that strains human reviewers who must deal with them, while also helping to manage that strain.
The trick for Kroah-Hartman and his peers will be to keep AI as a force multiplier, without drowning the open source maintainers. ®
Security
Everything's bigger and better in Texas – even data breaches
Hunting and fishing license incident catches 3M residents
Vercel debuts eve open source agent framework, tries to fix shadow AI with Passport
Cost premium of using AWS indirectly via Vercel is mitigated by more efficient use of compute resources, CTO claims
ZTE and China Telecom Guangdong advance cross‑vendor IP network simulation pilots, paving the way for intelligent network operations
PARTNER CONTENT: Leveraging >95% digital twin fidelity and multi-vendor collaboration to eliminate network change risks and achieve zero-error O&M
Microsoft's latest Windows bug belongs in the Recycle Bin
File deletion dialog swaps recognizable names for internal gibberish
PAAS AND IAAS
Graviton 5 impresses, but please, for the love of all that's holy, stop calling them 'AI chips'
AWS better at running chip fabs than their mouths
Britain's privacy watchdog quits after 'poor judgment' admission
John Edwards says his position had become 'untenable' following investigation into conduct including inappropriate attempts at humor
security
Feds freaked over Fable 5 after simple 'fix this code' prompt, not jailbreak, says researcher
ON-PREM
Amazon owns up to using 2.5bn gallons of H2O in its bit barns last year
science
AI and brain-computer interface allow speechless ALS patient to work a full-time job
PUBLIC SECTOR
Capita is about to sail past deadline to fix civil service pensions scheme
virtualization
Tesco is sprinting to quit VMware and Broadcom despite rapid migration risks
Cost premium of using AWS indirectly via Vercel is mitigated by more efficient use of compute resources, CTO claims
Campaigners say tech is unable to reliably distinguish between kids and adults at the boundary where use is planned
San Francisco plays host to hosting company's Localhost conference
The underlying technology is real...and borrowed from a partner the company failed to mention
Ed's note to Corey: Blink once if you're safe, twice if you're in danger
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
France's digital sovereignty push is struggling to escape the Microsoft gravity well
Nextcloud rollout shows locally controlled storage is one thing; getting users off Office is quite another
History of CentOS: How a biochemist's Linux hobby project became the enterprise world's default operating system
When a community came together after Red Hat said Windows was 'probably the right product'
Netflix wiz creates app to slash AI bills, then open sources it
Project Headroom could save you big money, too
OpenBSD 7.9 arrives, a diamond in the rough proud of every sharp edge
Sixtieth release adds more cores, delayed hibernation, and basic Wi-Fi 6 without losing its ascetic streak
Fedora: Microsoft is all aboard, but Deepin is dumped
Red Hat’s free distro loses a desktop, but makes an important new friend
LocalSend puts your sneakernet out of business
Like AirDrop, minus the Apple lock-in
Contact us
Advertise with us
Who we are
Newsletter
The Next Platform
DevClass
Blocks and Files
Situation Publishing
Cookies Policy
Privacy Policy
Ts & Cs
Do not share my personal information
Your Consent Options
Copyright. All rights reserved © 1998-2026.
