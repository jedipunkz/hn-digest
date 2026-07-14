---
source: "https://www.theregister.com/devops/2026/07/14/zig-creator-calls-buns-claude-rust-rewrite-unreviewed-slop/5270743"
hn_url: "https://news.ycombinator.com/item?id=48904245"
title: "Zig creator calls Bun's Claude Rust rewrite 'unreviewed slop'"
article_title: "Zig creator calls Bun’s Claude Rust rewrite ‘unreviewed slop’"
author: "giamma"
captured_at: "2026-07-14T09:47:18Z"
capture_tool: "hn-digest"
hn_id: 48904245
score: 1
comments: 1
posted_at: "2026-07-14T09:39:03Z"
tags:
  - hacker-news
  - translated
---

# Zig creator calls Bun's Claude Rust rewrite 'unreviewed slop'

- HN: [48904245](https://news.ycombinator.com/item?id=48904245)
- Source: [www.theregister.com](https://www.theregister.com/devops/2026/07/14/zig-creator-calls-buns-claude-rust-rewrite-unreviewed-slop/5270743)
- Score: 1
- Comments: 1
- Posted: 2026-07-14T09:39:03Z

## Translation

タイトル: Zig 作成者、Bun の Claude Rust のリライトを「未レビューの駄作」と呼ぶ
記事のタイトル: Zig 作成者、Bun の Claude Rust のリライトを「未レビューの駄作」と呼ぶ
説明: 移植にはわずか 11 日かかり、API 価格で約 165,000 ドルかかりました。

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
FIS と AWS による金融サービスの最新化
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
Zigの作者はBunのClaude Rustリライトを「未レビューの駄作」と呼ぶ
移植に要した時間はわずか 11 日で、API 価格で約 165,000 ドルかかりました。
Anthropic が所有する人気の JavaScript ランタイムとツールチェーンを AI が書き換えたことで、その実行速度に対する賞賛が巻き起こりましたが、プロジェクト自体の背後にあるコーディング手法に対する批判も巻き起こりました。
先週、Bun の作成者 Jarred Sumner は、並行して実行するクロード エージェント群を使用して、わずか 11 日間で Bun を Zig プログラミング言語から Rust に移植したと発表しました。この作業には API の価格で推定 165,000 ドルかかりました。これは、以前は大規模すぎて着手できないと考えられていたソフトウェアの改訂が、AI を使用すれば実際に実行可能になる可能性があることを示唆しています。
サムナー氏は、最近のクロードコードのソース漏洩に関係するバグを含め、Bun ユーザーが発見するバグの数が増えていることを考慮すると、移植が必要だったと述べた。
しかし、Zig の作成者である Andrew Kelley は、自分のプロジェクトが Bun の苦境の元凶であるとみなされることを望まず、サムナーの不適切なプログラミング慣行のせいだと主張しました。
ケリー氏にとって、Rust への移行は 2 つの言語間の機能の違いや AI の使用に関するものではなく、むしろ「2 つのプロジェクトの価値体系の相違」にあったと彼は書いています。
Bun は、ランタイム、パッケージ マネージャー、バンドラー、テスト ランナーで構成される JavaScript スイートです。 Node.js とうまく連携する高速なワンストップ ショップであるため、一部の開発者はこれを気に入っています。
Bun を高速化するために、Sumner は、Google の純正 V8 エンジンではなく、Apple の低メモリ高速起動 WebKit JavaScriptCore (JSC) エンジンを使用しました。彼はその性能と低コストを高く評価し、新進気鋭の Zig を使用しました。

-レベルコントロール。
Anthropic は 2025 年 12 月に Bun を買収しました。同社はそのコア ステート マシンを Bun 上に構築しました。
その時までに、サムナーも AI のコーディング能力を高く評価しており、バンの維持に AI を多用していました。買収の時点までに、RoboBun と呼ばれるクロード ボットが Bun リポジトリでの重労働の多くを行っていました。あらゆるコントリビューターの中で最も多くマージされた PR を提供し、バグを修正し、テストの失敗を修正しました。
しかし、Bun のユーザー ベースが拡大するにつれて、コードにさらに多くの亀裂が現れ始めました。ユーザーはソフトウェア全体で問題を発見しました。 Anthropic の 512,000 行のコードが 3 月に漏洩? NodeSource の報告によれば、これは Bun のせいで、ビルド中にソース マップを生成するバンドラーのバグのおかげでした。
これらすべてのバグは Zig のせいではないと、Sumner 氏は移行について詳しく説明した先週のブログ投稿で説明しました。 Bun のアーキテクチャには、ガベージ コレクションとアプリケーション主導のメモリ管理が混在していました。サムナー氏は、Zig がその目的のために設計されていないことを認めた。 Rust はメモリ管理の自動化において優れていました。
Zig の 500,000 行を別の言語に書き直すのは、手作業で行うと非常に大仕事になります。 「別の言語で書き直すには、小規模のエンジニアチームが丸1年かかります。それは、その間、バグ修正、セキュリティ修正、機能開発を凍結することを意味します」とサムナー氏は書いている。
代わりに、サムナーはクロードと一緒に行きました。彼は約 50 の動的なクロード コード ワークフローをスピンアップし、ピーク時には 1 分あたり約 1,300 行のコードに達し、100 万行を超える Rust コードを生成しました。作業には 11 日間かかり、API 価格で約 165,000 ドルの費用がかかりました。クロード・ファブルは重労働のほとんどを担当しました。
その後、Rust ベースの Bun は、100 万を超えるアサーションからなる Bun の徹底的なテスト スイートの対象となりました。サムナー氏によると、この製品はこれらのテストに 100% 合格しました。

サポートされているすべてのプラットフォームをスキップしたり削除したりすることなく。
「あの給料のエンジニアが、クロードが 11 日間で達成したマイルストーンを達成できるわけがありません」と感銘を受けた HashiCorp 共同創設者のミッチェル・ハシモト氏は X で述べています。
Anthropic の Bun Rust のリライトが AI の速度で統合されました
C++ の調査では、信頼が不足しているにもかかわらず AI の使用が増加していることが判明
最初の安定版 Go リリースとして TypeScript 7.0 の型チェックが高速化
Git は AI コーディングの津波に備えることができていない
しかし、Bun の実行速度は、優れたソフトウェア開発の中核となる原則を裏切るものでしょうか?
感銘を受けなかった人の一人は、「Bun Rust Rewrite についての私の考え」と題した熱のこもった投稿で不安を共有した Zig’s Kelley です。
Anthropic の買収前でさえ、「私たちは Bun のコードベースで目にしたプログラミングの実践にますます恐怖を感じました」と Kelley 氏は書いています。 Bun は、Zig を使用する最大かつ最も注目度の高いプロジェクトの 1 つであり、Anthropic が買収するまでは、Zig Software Foundation に定期的に資金を提供していました。
ケリー氏の見解では、このプロジェクトは新機能を積極的にリリースしたため、バグが山積し、エラー処理コードが不適切で、技術的負債が生じたという。
サムナーは「LLM にアクセスするずっと前から、すでに下手な文章を書いていました」とケリー氏は冗談めかして言いました。同氏は、サムナー氏は技術的な目標ではなくビジネス目標を達成するというプレッシャーにさらされていたのではないかと推測しており、アンスロピック社の買収によりそのプレッシャーはさらに増した。
実際、バンのコードベースはケリーの推測では非常に疑わしいものになっており、バンがジグと別れるのは良いニュースだったという。同氏の言葉を借りれば、「公的に Zig プログラミング言語の代表作と思われているものが、実際には Zig コードを書かない方法の代表的な例」ではなくなるだろうと彼は書いています。
Bun チームはまた、AI 支援の作業の一部をアップストリーム化しようとしました。

ジグ、無駄だった。 Bun の書き換えに先立って、チームは Zig のフォークを維持し、洞察力のある Reg 記者 Tim Anderson が 5 月に明らかにしたように、デバッグのコンパイル速度が 4 倍向上したと述べています。しかし、Zig プロジェクトは、AI ベースの貢献を受け入れないという方針を理由に、Bun の変更を受け入れませんでした。
Zig は、LLM によって生成された大量の投稿を受け取っていましたが、そのほとんどは品質が疑わしいものでした。 AI 生成コードに関するエンジニアリングの監督が欠如していると、将来的に無数の問題が発生するだろうとケリー氏は推論しました。
Kelley 氏は、Bun のテストが Zig でこれらのバグを見逃した場合、教師なしの Rust コードではどのようにバグが見つかるでしょうか?と指摘しました。
「レビューされていないコードの数百万行をすべて出荷することの議論は、テストスイートがすべてを捕捉するのに十分であるということです」と彼は書いています。 「Zig コードのバグを見つけるだけでは十分ではありませんが、レビューされていない 100 万行のバグを見つけるには十分ですか?」 ®
コラムニスト
Capita のような問題をどうやって解決しますか?
問題のある契約を社内に持ち込むにはスキルが必要 ホワイトホールはすでに採用に苦労している
閣僚、16歳未満のソーシャルメディア禁止に武器を与える、今年最も驚くべき調査ではない
ボフィンズ氏は認める：十代の若者たちの携帯電話を取り上げると、彼らは再び人間のように行動するようになる
Gobi X: AI のためのエネルギーを社会から奪うのではなく、より多くのエネルギーを生み出す
パートナーコンテンツ: エンビジョンは、その逆ではなく、コンピューティングが豊富な砂漠の電力を追い求めることで、データセンターの戦略をどのように逆転させているか
Anthropic の贅沢なトークナイザーにより AI の価格設定が複雑になる
トークンの消費だけですべてがわかるわけではありませんが、無視すべきではありません
コラムニスト
Microsoft はライセンス収入を守る戦いに負けつつある。その感覚に慣れた方が良い
クローン・ウォーズのリメイクの時が来た
Big Blue は再び POWER タワーで小さく考えます
最後の独自のミニコンピューター。現在は「デスクサイド」形式です。

それが好きです
aiとml
銀行やハイパースケーラーさえも現在、AIバブルについて警鐘を鳴らしている
ソフトウェア
中古ソフトウェアライセンス争いで裁判所がマイクロソフトの上訴を棄却
セキュリティ
GitHub AI エージェントがうまく質問するとプライベート リポジトリを漏洩する
アプリケーション
Microsoft、約20年ぶりにOWA Lightを終了へ
デボプス
GitHub、嘲笑が続いたため、リポジトリを CD に書き込むという提案を急遽打ち切る
トークンの消費だけですべてがわかるわけではありませんが、無視すべきではありません
十分な規模のインフラ企業にとって、予備のコンピューティングをレンタルするのは自然な流れです。
ロックダウン、とサティア・ナデラ氏が警告、レドモンド氏が古き良き時代にOpenAIに投入した数十億ドルを忘れているようだ
RAMポカリプスはAI黙示録の前兆である可能性がある
プラス：米国はイランのプロパガンダサイトを削除。マーケティング会社が「なぜあなたの情報を入手するのですか？」と尋ねます。さらに！
プラス：中国はスマートフォン監視ツールをアップグレード。リングは覗き見防止の姿勢を緩和します。などなど
ジェフ・モス氏によると、投票村のレポートは非常に成功しており、今後は DEF CON 全体が含まれることになる
会社全体の評価額は35億ドル以上に相当するが、売却部分は特定されていない
プラスの面としては、情報セキュリティは長く安定したキャリアを築くのに適しています。
Cinnamon 6.8 は Wayland をサポートします – 必要に応じて
Linux Mint デスクトップの次期バージョンには両方の種類のディスプレイ サーバーが搭載されています
KDE Plasma ユーザーは恐ろしい変化の予兆に直面しています: 6.6.6 の登場
現在は 6.7 が最新版であり、6.8 では好むと好まざるにかかわらず Wayland を入手できるようになります。
FOSS クラウドオフィススイート間の競争が激化する中、Collabora が CODE 26.04 をリリース
Markdown のサポートと、よりスマートな数式エラー処理に加え、AI が統合されました (デフォルトではオフになっています)。
GIMP 0.54 が Flatpak 形式で復活し、過去から吹き飛ばされる
最初 (そして最後) のリリースで懐かしい人向けのレトロ コンピューティングの楽しみ

GTK の代わりに Motif を使用するには
Bcachefs は新しい「パフォーマンス リリース」で実験的ステータスを終了します
Rustは増えるが、AIスロップの問題も増える
フランスのデジタル主権推進はマイクロソフトの重力から逃れるのに苦戦している
Nextcloud のロールアウトは、ローカルで制御されるストレージが 1 つのことであることを示しています。ユーザーを Office から解放することはまったく別のことです
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

The port took just 11 days and about $165,000 at API pricing

Jump to main content
Search
TOPICS
Special Features
All Special Features
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
Zig creator calls Bun’s Claude Rust rewrite ‘unreviewed slop’
The port took just 11 days and about $165,000 at API pricing
An AI rewrite of a popular Anthropic-owned JavaScript runtime and toolchain has sparked praise for the speed of its execution, but also criticism of the coding practices behind the project itself.
Last week, Bun creator Jarred Sumner announced that he ported Bun from the Zig programming language to Rust in only 11 days, using a fleet of Claude agents running in parallel. The work cost an estimated $165,000 at API pricing, suggesting that software revisions previously considered too large to undertake could actually be feasible now with AI.
Sumner said the port was necessary given the growing number of bugs Bun users were finding, including one implicated in the recent Claude Code source leak.
But the creator of Zig, Andrew Kelley , didn’t want his project to be seen as the culprit behind Bun’s woes, which he blames on Sumner’s bad programming practices.
For Kelley, the move to Rust was not about the feature differences between the two languages, or even the use of AI, but rather “the diverging value systems of the two projects,” he wrote.
Bun is a JavaScript suite consisting of a runtime, package manager, bundler and test runner. Some developers like it because it is a fast one-stop shop that plays well with Node.js.
To make Bun speedy , Sumner used Apple's low-memory fast-start WebKit JavaScriptCore (JSC) engine, rather than Google’s stock V8 engine. He used the up-and-coming Zig because he appreciated its performance and low-level control.
Anthropic acquired Bun in December 2025. The company built its core state machine on Bun.
By then, Sumner had also grown to appreciate AI’s coding abilities, and was using it heavily in the upkeep of Bun. By the time of acquisition, a Claude Bot called RoboBun had been doing a lot of the heavy lifting in the Bun repo. It supplied the most merged PRs of any contributor, fixing bugs and remediating test failures.
But as Bun’s user base grew, more cracks started appearing in the code. Users found issues across the software. Anthropic’s 512,000-line code leak in March? That was Bun’s fault, thanks to a bug in the bundler that generated source maps during builds even when told not to, NodeSource reported .
All these bugs weren’t Zig’s fault, Sumner explained in a blog post last week detailing the migration . Bun’s architecture mixed garbage collection and application-driven memory management. Sumner admitted that Zig wasn’t designed for that task. Rust was just better at automating memory management.
Rewriting 500,000 lines of Zig into another language would be a gargantuan undertaking if done by hand. “A rewrite in another language would take a small team of engineers a full year. It would mean freezing bugfixes, security fixes or feature development for that time,” Sumner wrote.
Instead, Sumner went with Claude. He spun up about 50 dynamic Claude Code workflows, reaching a peak of about 1,300 lines of code per minute and generating over a million lines of Rust code . The job took 11 days and cost about $165,000 at API pricing. Claude Fable did most of the heavy lifting.
The Rust-based Bun was then subjected to Bun's exhaustive test suite of more than one million assertions. According to Sumner, it passed 100 percent of those tests across all supported platforms without skipping or deleting any.
“There’s absolutely no way an engineer with that salary would’ve been able to achieve the milestones Claude did in 11 days,” an impressed HashiCorp co-founder Mitchell Hashimoto noted on X .
Anthropic’s Bun Rust rewrite merged at speed of AI
C++ survey finds AI use rising, though trust is in short supply
Speedier type checks in TypeScript 7.0 as first stable Go release ships
Git is unprepared for the AI coding tsunami
But does Bun’s speed of execution betray the core tenets of good software development?
One person not impressed has been Zig’s Kelley, who shared his misgivings in an impassioned post entitled “ My Thoughts on the Bun Rust Rewrite ."
Even before the Anthropic acquisition, “we became increasingly horrified at the programming practices we saw in Bun's codebase,” Kelley wrote. Bun was one of the largest and highest profile projects using Zig and, up until the Anthropic acquisition, a regular financial contributor to The Zig Software Foundation.
In Kelley’s view, the project aggressively released new features, resulting in piled-up bugs, bad error-handling code, and technical debt.
Sumner “was already writing slop well before he had access to LLMs,” Kelley quipped. He speculated that Sumner may have been under pressure to meet business objectives rather than technical ones, a pressure that increased with Anthropic’s acquisition.
In fact, Bun’s codebase had grown so suspect in Kelley’s estimation that Bun parting with Zig was good news. As he put it, no longer would “the publicly presumed poster child for Zig programming language actually [be] the prime example of How Not To Write Zig Code,” he wrote.
The Bun team also tried to upstream some of its AI-assisted work to Zig, to no avail. Leading up to the Bun rewrite, the team maintained a fork of Zig that it said improved debug compilation speed fourfold, as eagle-eyed Reg reporter Tim Anderson revealed in May . But the Zig project would not accept Bun’s changes, citing a policy of not accepting AI-based contributions.
Zig had been getting an influx of LLM-generated submissions, most of dubious quality. This lack of engineering oversight around AI-generated code would lead to countless problems down the road, Kelley reasoned.
Kelley pointed out that if Bun’s tests missed these bugs in Zig, how would they be caught in unsupervised Rust code?
“The argument for shipping all the million lines of unreviewed code is that the test suite is good enough to catch everything,” he wrote. “It's not sufficient to catch bugs in Zig code but it is sufficient to catch bugs in [a] million lines of unreviewed slop?” ®
columnists
How do you solve a problem like Capita?
Bringing troubled contracts in-house requires skills Whitehall is already struggling to recruit
Ministers arm under-16s social media ban with least surprising study of the year
Boffins confirm: taking away teens' phone makes them act like humans again
Gobi X: Creating more energy for AI, not taking it from society
PARTNER CONTENT: How Envision is reversing the datacenter playbook by making computing chase abundant desert power, not the other way around
Anthropic's extravagant tokenizer complicates AI pricing
Token consumption doesn't tell the whole tale but it shouldn't be ignored
columnists
Microsoft is losing the battle to protect license lucre. It better get used to the feeling
Time for the Clone Wars remake
Big Blue thinks small, again, with POWER tower
The last proprietary minicomputer, now in ‘deskside’ form if you fancy that
ai and ml
Even banks and hyperscalers are now sounding the alarm about the AI bubble
software
Court tosses Microsoft's appeal in pre-owned software licenses battle
Security
GitHub AI agent leaks private repos when asked nicely
applications
Microsoft to switch off OWA Light after nearly two decades
DEVOPS
GitHub cuts short offer to burn repos on CD after mockery ensues
Token consumption doesn't tell the whole tale but it shouldn't be ignored
Renting out spare compute is simply the natural progression for any sufficiently large infra company
Lock it down, warns Satya Nadella, seemingly forgetting the billions Redmond chipped in to OpenAI back in the good old days
The RAMpocalypse may be the precursor to the AIpocalypse
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
Cinnamon 6.8 will support Wayland – if you want it
Next version of Linux Mint’s desktop has both kinds of display server
KDE Plasma users face a dire omen of change: 6.6.6 arrives
6.7 is now current, and in 6.8 you're getting Wayland whether you like it or not
Collabora releases CODE 26.04 as rivalry between FOSS cloudy office suites heats up
Now with Markdown support and smarter formula error handling – plus integrated AI, though it's off by default
Blast from the past as GIMP 0.54 is revived in Flatpak form
Retro-computing fun for the nostalgic with first (and last) release to use Motif instead of GTK
Bcachefs exits experimental status in new 'performance release'
More Rust, but more trouble with AI slop, too
France's digital sovereignty push is struggling to escape the Microsoft gravity well
Nextcloud rollout shows locally controlled storage is one thing; getting users off Office is quite another
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
