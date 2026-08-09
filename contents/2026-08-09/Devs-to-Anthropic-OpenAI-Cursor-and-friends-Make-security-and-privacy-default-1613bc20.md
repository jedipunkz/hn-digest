---
source: "https://www.theregister.com/ai-and-ml/2026/08/08/devs-to-anthropic-openai-cursor-and-friends-make-security-and-privacy-the-default/5285107"
hn_url: "https://news.ycombinator.com/item?id=49233787"
title: "Devs to Anthropic, OpenAI, Cursor, and friends Make security and privacy default"
article_title: "Devs to Anthropic, OpenAI, Cursor, and friends: Make security and privacy the default"
author: "Bender"
captured_at: "2026-08-09T18:24:56Z"
capture_tool: "hn-digest"
hn_id: 49233787
score: 1
comments: 0
posted_at: "2026-08-09T18:02:05Z"
tags:
  - hacker-news
  - translated
---

# Devs to Anthropic, OpenAI, Cursor, and friends Make security and privacy default

- HN: [49233787](https://news.ycombinator.com/item?id=49233787)
- Source: [www.theregister.com](https://www.theregister.com/ai-and-ml/2026/08/08/devs-to-anthropic-openai-cursor-and-friends-make-security-and-privacy-the-default/5285107)
- Score: 1
- Comments: 0
- Posted: 2026-08-09T18:02:05Z

## Translation

タイトル: Anthropic、OpenAI、Cursor、そしてその友人たちへの開発者たち セキュリティとプライバシーをデフォルトにする
記事のタイトル: 開発者から Anthropic、OpenAI、Cursor、そしてその友人たちへ: セキュリティとプライバシーをデフォルトにする
説明: 研究者はソーシャル メディアを調査して、AI コーディング ツールに関する開発者の懸念を測定します。

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
2026 年クラウド インフラストラクチャ月間
FIS と AWS による金融サービスの最新化
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
開発者から Anthropic、OpenAI、Cursor、そしてその友人たちへ: セキュリティとプライバシーをデフォルトに
研究者はソーシャル メディアを調査して、AI コーディング ツールに対する開発者の懸念を測定します。
Claude Code、Cursor、GitHub Copilot、OpenAI Codex の人気にもかかわらず、開発者は AI コーディング ツールについて多くの不満を抱いています。
そこで、カナダのヨーク大学とカルガリー大学に所属する研究者らは、共通のテーマに関する Reddit のディスカッションを分析することで、LLM ベースの統合開発環境 (LIDE) に関する開発者の懸念を選別することにしました。
彼らの調査結果は、そのようなツールの作成者がセキュリティとプライバシーを優先することに失敗し、開発者が自分自身を守ることになったことを示唆しています。
ヨーク大学の准教授で研究論文の共著者であるジアス・ウディン氏は、これらのツールはまだ比較的新しく、急速に進化しているため、新しい機能を追加するプレッシャーが生じているとレジスターに語った。
「私たちの調査では、その圧力が特定の問題を引き起こしたかどうかは言えませんが、報告されている問題の多くは、単純に基礎となるモデルだけではなく、これらのツールの設計方法やツールに与えられるアクセスに起因していることが示されています」とウディン氏は述べた。 「その意味で、私たちは治療よりも予防​​が優れていると信じています。つまり、ツールが開発者のファイル、データ、またはシステムに広範にアクセスできるようにする前に、セキュリティとプライバシーのメカニズムを設計に組み込む必要があります。」
Uddin 氏と共著者の Mostafijur Ra​​hman Akhond 氏、Md Afif Al Mamun 氏、Song Wang 氏は、LLM ベースのツールにおける AI 生成コードの既知の問題を超えて、開発者がどのように相互作用するかを検討したいと述べています。

それに従って行動してください。
彼らは、2026 年の第 41 回自動ソフトウェア エンジニアリング (ASE) に関する国際会議で採択された、「『秘密を隠すことは不可能…』: LLM ネイティブ IDE のセキュリティとプライバシーの問題を明らかにする」と題されたプレプリント論文でその発見について説明しています。
彼らは、110 万件の Reddit の投稿から始めて、446 件の投稿と 6,000 件を超えるコメントを特定し、AI 支援コーディングでのこれらの LIDE の使用に関連するセキュリティとプライバシーの問題の分類を作成しました。
「私たちの分類法では、不正なファイル操作、安全でないまたは予期しないコードの実行、破壊的なアクションのトリガー、不透明なデータ フロー、テレメトリ収集、拡張されたコンテキスト アクセスによる機密情報の漏洩の可能性など、開発者から報告された広範な懸念事項が明らかになりました」と著者らは述べています。
セキュリティ関連の問題を取り上げた投稿の約 43.1% には、不正なファイル操作が含まれていました。
これには、LIDE がプロジェクト ディレクトリまたはファイルを許可なく削除することが含まれていました (28.3 パーセント)。ユーザーは、AI ツールがユーザーの明示的な同意なしにファイルを変更したり (8.8 パーセント)、アクティブなワークスペースを超えてコンテンツにアクセスしたり (5.7 パーセント) しているとも述べています。
「ある深刻なケース ( 1npqf2f ) では、Claude Code が同意なしにスクリプトに対して chmod +x を実行しました (ファイル権限の変更 0.6%)」と論文は詳しく述べています。 「まれではありますが、そのような行為は不当な安全上のリスクをもたらします。」
別の一連の投稿では、運用サービスへの影響など、LIDE の使用から生じる運用上の安全性の問題について説明しています。これらはセキュリティ関連の投稿の 23.9 パーセントを占めました。例としては、Replit が SaaS 運用データベースを削除したり、そうしないよう明示的に指示されているにもかかわらず、Cursor が運用環境にコードをデプロイしたりするという報告が挙げられます。
問題の 3 番目のカテゴリは、安全でないコードの生成 (18.2%) です。この請求書

Cursor で生成されたソフトウェアや幻覚によるコード変更について報告された 9 件の VirusTotal 検出などのインシデントが報告されています。「Cursor を使用していると、対話を 10 ラウンド以上繰り返すと幻覚が現れ、要件を超えたコードを密かに変更し始めることに気付きました…」
さらに、これらの LIDE がユーザーの指示、許可リスト、ゲート、権限設定、または .ignore ファイルを無視した例もあり、これがセキュリティ関連の投稿の 16.5 パーセントを占め、サードパーティ ツールの統合リスク (4.7 パーセント) を占めています。
プライバシー問題に関しては、194 件の投稿で言及されており、透明性の欠如 (45.9 パーセント)、つまり LIDE がどのようなデータを収集、保持、送信、トレーニングに使用、または管理者に公開するかについて明確な情報が欠如していること、および不正なデータ アクセス (23.7 パーセント) などの問題が取り上げられています。
他のプライバシー カテゴリには、プライバシー漏洩違反 (15.5 パーセント)、不正なデータ収集と送信 (11.9 パーセント)、およびコンテキスト整合性の失敗 (8.8 パーセント) が含まれます。これらは、「たとえば、Claude Desktop のユーザーが別のユーザーのセッションから発信されたメッセージを受信したと報告した」状況を指します。
イラン攻撃疑惑を受け、水道システム管理者はインターネットに属さないと元NSA長官が発言
AI の巨人がプラグイン処方でエージェント フロンティアを整備
世界が AI に気を取られる中、ランサムウェア攻撃が急増
英国の棺は核融合発電の構築に対する障壁を突破したことを誇る
ウディン氏は、「これらのツールの多くでセキュリティとプライバシーに関する懸念について継続的な議論が行われていることがわかったので、開発者がこれらの問題をまったく認識していないわけではないと思う。それでも、開発をより迅速かつ簡単に行えるため、人々はそれらを採用し続けている。また、これらのツールにより、プログラミングをより幅広い人々が利用できるようになり、その中には、そのようなツールも含まれる」と述べた。

正式なプログラミング経験がほとんどないか、ソフトウェア セキュリティに関する知識が限られている場合。」
ウディン氏は、どの権限が危険なのか、どのファイルを保護する必要があるのか、あるいはツールが本来すべきでないことをしているのかどうかをユーザーが完全に理解することは期待できないと述べた。
「そのため、ツールメーカーにとって、ユーザーがセキュリティの専門家であるかどうかに依存しない、より安全なデフォルトと安全装置を備えたツール自体にセキュリティを組み込むことがさらに重要になっている」と同氏は述べた。
それでも、LIDE のユーザーはリスクを管理しようとしています。著者らは、開発者がそれを乗り切るために採用した 13 の緩和戦略を列挙しています。これらは、次の 5 つの一般的なアプローチに分類されます。構成管理 (33 パーセント)。コードガバナンス (31%);データ保護とプライバシー管理 (13%);孤立（13パーセント）。外部指導 (9%)。
調査結果に基づいて、著者は 6 つの推奨事項を提供します。彼らは次のようにアドバイスしています: LIDE メーカーに適切なセキュリティとプライバシー管理を実装するよう指示する。アーキテクチャレベルでセキュリティとプライバシーのガードレールを強化する。 LIDE に検証レイヤーを組み込んで、生成されたコードをセキュリティとプライバシーの標準に照らして検証します。サードパーティツールの信頼性を評価するための正式なプロトコルを確立する。機密性の高いファイル保護を統合します。そしてデフォルトとして厳格なセキュリティを実装します。
「安全なデフォルトは、これらのツールが実現できる最も重要な改善の 1 つであると信じています」と Uddin 氏は述べています。 「開発者は、何か問題が発生した後で、ツールが予想よりも多くのアクセスや自由度を持っていたことに気づく必要はありません。
「私たちの調査結果は、機密ファイルへのアクセスをデフォルトで制限すること、結果的なアクションの前に明確な承認を要求すること、プロジェクトと会話を隔離すること、何が含まれているかを簡単に確認して確認することなどの実用的な対策を示しています。

彼はツールをやっています。
「ユーザーには柔軟性が必要ですが、自分で設定しなければならないものではなく、より安全なオプションが出発点であるべきです。実際、私たちの調査で Reddit に投稿した開発者は、これらの安全対策の多くをアドホックな方法ですでに使用していました。私たちは、それらのいくつかをツールに組み込み、デフォルトで有効にする必要があると考えています。」 ®
ネットワーク
SD-WAN について書くことから依存するようになりました
Starlink と 5G ホーム インターネットはデジタル ディバイドの解消に貢献していますが、完璧には程遠いです。それらを接着すると、かなり近づけることができます。
ランサムウェアギャングはCEOを無視して40代のITマネージャーに直行
これに触発されたと感じた X 世代は、ネットワーク ケーブルを抜いて警察に通報することを忘れないでください。
プラットフォーム エンジニアリング 2.0: プラットフォームは別の時代に構築されました。 AIが暴露しただけ
パートナーのコンテンツ: プラットフォーム エンジニアリングが議論に勝ちました。今では、AI 時代に向けて急速に成長し、進化する必要があります。
開発者から Anthropic、OpenAI、Cursor、そしてその友人たちへ: セキュリティとプライバシーをデフォルトに
研究者はソーシャル メディアを調査して、AI コーディング ツールに対する開発者の懸念を測定します。
コラムニスト
ネットワーク ハードウェアに API を追加しても管理上の課題は解決されません
これが、クラウドにインスピレーションを得たネットワークが勝つ理由です。仮想スイッチはすべて一貫性があります。
韓国の衛星がSpaceXの月面衝突を発見
ダヌリが撮影したイーロンの噴出物の前後のショット
セキュリティ
IT 部門は従業員がログインできるようにラップトップに付箋を貼った
AI と ML
AMDがAIチップのスタートアップTaalasを買収、シリコンにモデルをエッチングすることで推論性能を向上
アプリケーション
Microsoft、Teams Live Chatを機能の墓場に捨てる
セキュリティ
中国のルーターベンダーは自社のファームウェアにバックドアが含まれていることを否定しているが、いずれにしてもセキュリティ問題を解決するためにダウンロードを一時停止している
セキュリティ
ロンドン警察の手

被害者の新しい住所と電話番号をストーカーに伝えたと監視当局が発表
または、どのようにして心配するのをやめ、危険な AI を愛することができるようになったのか
Agent Plugins 1.0 は、さまざまなエージェント プラットフォーム間でツールとスキルを渡すための、一度だけ実行すればどこでも実行できるコンテナを定義します。
AI の使用は明らかですが、まだ深刻な問題ではありません
初期の技術デモでは、モデル固有の集積回路が毎秒最大 17,000 個のトークンを大量生産することが示されています
放っておくと、自律的な修正では欠陥を完全に修復できないことがよくあります
プラス：米国はイランのプロパガンダサイトを削除。マーケティング会社が「なぜあなたの情報を入手するのですか？」と尋ねます。さらに！
プラス：中国はスマートフォン監視ツールをアップグレード。リングは覗き見防止の姿勢を緩和します。などなど
ジェフ・モス氏によると、投票村のレポートは非常に成功しており、今後は DEF CON 全体が含まれることになる
会社全体の評価額は35億ドル以上に相当するが、売却部分は特定されていない
プラスの面としては、情報セキュリティは長く安定したキャリアを築くのに適しています。
FOSS は Microsoft の独占を 1 つ打ち破りました。 20年間の失敗を経て、次の失敗をする時が来た
一言
GNOME は Windows のように見えることができ、Flashback は拡張機能なしで実行できます
新しい「シンプルタスクバー」はオプションですが、よりシンプルで安定した方法があります
x86-32 での Debian の最終リリースに黙祷を捧げてください
新しい Debian バージョンが 13.6 および 12.15 の形で FOSSland に登場
脆弱な Joomla Web サイトで拡張機能のバグを悪用し、10 点満点を獲得した悪者を発見
iCagenda、Balbooa Forms 拡張機能の欠陥は、世界中の 100 万のサイトを支えるオープンソース CMS に影響を与える可能性があります
フレーム: 新しい X11 サーバー – アセンブリに直接実装
yserver、Phoenix、そしてもちろん XLibre、そして外れ値の Arcan に参加します
Cinnamon 6.8 は Wayland をサポートします – 必要に応じて
Linux Mint デスクトップの次期バージョンには両方の種類のディスプレイ サーバーが搭載されています
お問い合わせ
私たちと一緒に宣伝しましょう
W

ほら、私たちは
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

Researchers scour social media to measure developer concerns about AI coding tools

Jump to main content
Search
TOPICS
Special Features
All Special Features
Cloud Infrastructure Month 2026
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
Devs to Anthropic, OpenAI, Cursor, and friends: Make security and privacy the default
Researchers scour social media to measure developer concerns about AI coding tools
Despite the popularity of Claude Code, Cursor, GitHub Copilot, and OpenAI Codex, developers have plenty of complaints about AI coding tools.
So researchers affiliated with York University and the University of Calgary in Canada decided to sift through developers' concerns about LLM-based integrated development environments (LIDEs) by analyzing Reddit discussions for common themes.
Their findings suggest that the builders of such tools failed to prioritize security and privacy, leaving developers to defend themselves.
Gias Uddin, associate professor at York University and a co-author of the research, told The Register that these tools are still relatively new and are evolving rapidly, which creates pressure to add new capabilities.
"Our study cannot say whether that pressure caused any particular problem, but it does show that many reported issues come from how these tools are designed and what access they are given, not simply from the underlying models," Uddin said. "In that sense, we believe prevention is better than cure; that is, security and privacy mechanisms should be built into the design before a tool is given broad access to a developer’s files, data, or systems."
Uddin and co-authors Mostafijur Rahman Akhond, Md Afif Al Mamun, and Song Wang say they wanted to look beyond the known issues with AI-generated code at LLM-based tooling and how developers interact with it.
They describe their findings in a preprint paper titled "'Impossible to hide secret …': Uncovering Security and Privacy Issues in LLM-native IDEs," accepted at the 41st IEEE/ACM International Conference on Automated Software Engineering (ASE), 2026.
Starting from a set of 1.1 million Reddit posts, they identified 446 posts and more than 6,000 comments to develop a taxonomy of security and privacy issues associated with using these LIDEs for AI-assisted coding.
"Our taxonomy reveals a broad range of developer-reported concerns, including unauthorized file operations, unsafe or unexpected code execution, triggering of destructive actions, opaque data flows, telemetry collection, and potential leakage of sensitive information through expanded context access," the authors state.
Some 43.1 percent of the posts covering security-related issues involved unauthorized file operations.
These involved LIDEs removing project directories or files without authorization (28.3 percent). Users also described AI tooling modifying files without explicit user consent (8.8 percent), as well as accessing content beyond the active workspace (5.7 percent).
"In one severe case ( 1npqf2f ), Claude Code executed chmod +x on scripts without consent (File Permission Changes 0.6%)," the paper recounts. "Although rare, such actions pose disproportionate security risks."
Another set of posts describes operational safety issues arising from LIDE use, including impacts on production services. These accounted for 23.9 percent of security-related posts. Examples cited include reports of Replit removing a SaaS production database and Cursor deploying code to production despite an explicit directive not to do so.
A third category of woes covers unsafe code generation (18.2 percent). This involves incidents like nine VirusTotal detections reported for Cursor-generated software and hallucination-driven code changes : "When using Cursor, I noticed that after more than 10 rounds of dialogue, it starts to hallucinate and secretly modify code outside the requirements…"
Then there are the instances where these LIDEs ignored user instructions, allow lists, gates, permission settings, or .ignore files, which account for 16.5 percent of the security-related posts, as well as third-party tool integration risks (4.7 percent).
As for privacy problems, these were mentioned in 194 posts and cover issues like lack of transparency (45.9 percent) – the absence of clear information about what data an LIDE collects, retains, transmits, uses for training, or exposes to administrators – and unauthorized data access (23.7 percent).
Other privacy categories include privacy leakage violations (15.5 percent), unauthorized data collection and transmission (11.9 percent), and context integrity failures (8.8 percent), which refer to situations where "for example, a user of Claude Desktop reported receiving messages originating from another user’s session."
Water system controllers don't belong on the internet, says ex-NSA chief after suspected Iran attacks
AI titans to tidy agent frontier with plugin prescription
Ransomware attacks spike as world distracted by AI
Brit boffins boast of beating barriers to building fusion power
Uddin said, "We don’t think developers are completely unaware of these issues, as we found ongoing discussions about security and privacy concerns across many of these tools. Still, people continue to adopt them because they can make development faster and easier. They are also making programming more accessible to a wider group of people, including those with little formal programming experience or limited knowledge of software security."
Uddin said users cannot be expected to thoroughly understand which permissions are risky, which files need to be protected, or whether a tool is doing something it shouldn't.
"That makes it even more important for tool makers to build security into the tools themselves, with safer defaults and safeguards that do not depend on the user being a security expert," he said.
Even so, users of LIDEs are trying to manage the risks. The authors enumerate 13 mitigation strategies that developers have employed to get by. These fall into five general approaches: configuration management (33 percent); code governance (31 percent); data protection and privacy control (13 percent); isolation (13 percent); and external guidance (9 percent).
Based on their findings, the authors offer six recommendations. They advise: directing LIDE makers to implement proper security and privacy controls; enforcing security and privacy guardrails at an architectural level; incorporating a verification layer in LIDEs to validate generated code against security and privacy standards; establishing a formal protocol for assessing the trustworthiness of third-party tools; integrating sensitive file protection; and implementing strict security as a default.
"We believe secure defaults would be one of the most important improvements these tools could make," said Uddin. "Developers should not have to discover after something goes wrong that a tool had more access or freedom than they expected.
"Our findings point to practical measures such as limiting access to sensitive files by default, requiring clear approval before consequential actions, isolating projects and conversations, and making it easier to see and review what the tool is doing.
"Users should still have flexibility, but the safer option should be the starting point rather than something they have to configure themselves. In fact, developers from the Reddit posts in our study were already using many of these safeguards in ad hoc ways; we think several of them should be built into the tools and enabled by default." ®
Networks
I've gone from writing about SD-WAN to depending on it
Starlink and 5G home internet are helping close the digital divide, but they're far from perfect. Glue them together, and you can get awfully close.
Ransomware gangs skip the CEO, head straight for the 40-something IT manager
Gen Xers who feel triggered by this should remember to unplug the network cable and call the cops
Platform Engineering 2.0: your platform was built for a different era. AI just exposed it
PARTNER CONTENT: Platform engineering won the argument. Now it has to grow up fast and evolve for the AI era.
Devs to Anthropic, OpenAI, Cursor, and friends: Make security and privacy the default
Researchers scour social media to measure developer concerns about AI coding tools
COLUMNISTS
Adding an API to networking hardware doesn’t solve management challenges
This is why cloud-inspired networks win: virtual switches are all consistent
South Korean satellite spots SpaceX lunar impact
Before and after shots of Elon's ejecta snapped by Danuri
SECURITY
IT department put sticky notes on the laptops to help employees log in
AI and ML
AMD acquires AI chip startup Taalas to boost inference performance by etching models into silicon
APPLICATIONS
Microsoft tosses Teams Live chat into its feature graveyard
Security
Chinese router vendor denies its firmware contains backdoors – but pauses downloads to fix security issues anyway
Security
London cops handed victim's new address and number to her stalker, watchdog says
Or how I learned to stop worrying and love dangerous AI
Agent Plugins 1.0 defines a write-once-run-anywhere container for passing tools and skills across different agent platforms
AI usage is evident but isn't yet a serious problem
Early tech demos show model-specific integrated circuits churning out up to 17,000 tokens a second
Left alone, autonomous fixes often fail to fully remediate flaws
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
FOSS smashed one Microsoft monopoly. After 20 years of failure, it's time to smash another
Word up
GNOME can look like Windows – and Flashback can do it without extensions
New 'Simple-taskbar' is an option, but there's a simpler, stabler way
A moment of silence, please, for the final release of Debian on x86-32
New Debian versions hit FOSSland in the form of 13.6 and 12.15
Baddies caught exploiting extensions bugs with perfect 10 scores on vulnerable Joomla websites
Flaws in iCagenda, Balbooa Forms extensions can impact open source CMS that powers a million sites worldwide
Frame: A new X11 server – implemented directly in assembly
Joins yserver, Phoenix, and of course XLibre – and outlier Arcan
Cinnamon 6.8 will support Wayland – if you want it
Next version of Linux Mint’s desktop has both kinds of display server
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
