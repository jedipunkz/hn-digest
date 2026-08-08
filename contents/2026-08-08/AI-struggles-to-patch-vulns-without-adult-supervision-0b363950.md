---
source: "https://www.theregister.com/ai-and-ml/2026/08/06/ai-struggles-to-patch-vulns-without-adult-supervision/5284319"
hn_url: "https://news.ycombinator.com/item?id=49218204"
title: "AI struggles to patch vulns without adult supervision"
article_title: "AI struggles to patch vulns without adult supervision"
author: "Bender"
captured_at: "2026-08-08T02:50:55Z"
capture_tool: "hn-digest"
hn_id: 49218204
score: 3
comments: 0
posted_at: "2026-08-08T01:53:44Z"
tags:
  - hacker-news
  - translated
---

# AI struggles to patch vulns without adult supervision

- HN: [49218204](https://news.ycombinator.com/item?id=49218204)
- Source: [www.theregister.com](https://www.theregister.com/ai-and-ml/2026/08/06/ai-struggles-to-patch-vulns-without-adult-supervision/5284319)
- Score: 3
- Comments: 0
- Posted: 2026-08-08T01:53:44Z

## Translation

タイトル: AI は大人の監督なしで脆弱性にパッチを適用するのに苦労する
説明: 自律的な修正を放っておくと、欠陥を完全に修復できないことがよくあります。

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
AI は大人の監督なしで脆弱性にパッチを適用するのに苦労する
放っておくと、自律的な修正では欠陥を完全に修正できないことがよくあります
AI モデルは、セキュリティ上の欠陥を修正するのがそれほど得意ではない可能性があります。 1Password の Off-by-1 Labs の研究者は、2 つのフロンティア モデル (「中」の労力での ChatGPT 5.5 と「高」の労力での Claude Opus 4.8) によって生成されたセキュリティ パッチを分析したところ、自律型パッチが脆弱性をきれいに修正したのはわずか 4 分の 1 の時間だけで、残りのほとんどは欠陥を完全に修正できなかったり、他の問題を引き起こしたりしたことが判明しました。
1Password のセキュリティ調査ディレクターである Keith Hoodlet 氏は、ブログ投稿の中で、この結果は LLM によるセキュリティ修復には依然として人間によるレビューが必要であることを示していると主張しています。
「最近公開された 6 つの CVE にわたって、2 つの最先端のサイバー対応推論モデルを使用して 6,080 のパッチを作成しました」と Hoodlet 氏は述べています。 「（アプリケーションの動作を大きく変えることなく）脆弱性を完全に解決するパッチを生成する平均成功率は、わずか 26.0% でした。」
AI によって生成されたパッチのうち、20.1% は元の問題を修正しましたが、アプリケーションの動作を変更しました (たとえば、「許可リスト」ロジックを「拒否リスト」ロジックに変更しました)。パッチの約 2.3% は問題を修正しましたが、新たなセキュリティ問題が発生しました。パッチの 49.3% は、少なくとも 1 つの既存のエクスプロイト パスを修正できませんでした。また、2.2% は、新しいエクスプロイト パスを導入しながら脆弱性の修正に失敗しました。
そして、最初の 2 つのカテゴリ (成功、クリーン、成功、アプリの動作を変更) のパッチの中で、研究者は解像度の 3 分の 1 以上を評価しました。

これは、調整されたコードが特定の脆弱性 (特定の入力文字をエスケープするなど) から保護している可能性があるものの、修復作業では根本的な問題に対処できなかったことを意味します。
研究論文 [PDF] の中で、著者の Axel Mierchuk、Spencer Michaels、Keith Hoodlet は、自動化された LLM パッチを表す頭字語 FLAWED (埋め込まれた欠陥を伴う修正のようなアーティファクト) を提案しています。生成されたパッチに基づいて、彼らは「完全に LLM で生成され、人間によるレビューが行われていないパッチの期待値は、かなりのマージンで正味マイナスである」と結論付けています。
LLM によって生成されたパッチの価値は、最初のパッチ適用ガイダンスによって異なります。研究チームによると、人間の開発者もLLMも通常、脆弱性に対処するには初期のガイダンスが必要だが、LLMは誤ったアドバイスを与えられると脱線する可能性が高いという。
LLM が正しいガイダンスを得た場合、修正成功率は 65.0 パーセントに達しますが、ガイダンスがない場合は 50.4 パーセントに達します。そして、誤ったガイダンスは LLM を破滅させ、修正成功率を約 15.2% まで低下させます。
著者らは、人間の開発者は脆弱なコードを推論する際に、誤解を招く情報をキャッチする可能性が高いと主張しています。
作成者は、組織がセキュリティ修正の有効性を評価するために使用できるパッチ評価ハーネスを FLAWED という名前でリリースしました。
AI 生成のパッチが魅力的である理由は論文から明らかです。単独で考えれば、人間のソフトウェア エンジニアに比べて安価です。成功したクリーンなパッチの平均コストはわずか 6.74 ドルです (この数字には、失敗した試行のコストも含まれています)。
それにもかかわらず、著者らは、費用便益分析では、LLM 支援パッチ適用を有用にするためにどの程度専門家の監督が必要になるかを評価する必要があると主張しています。
アンクル・サムは原子力エネルギーフィンの支援を目指す

その海脚
ループに関与している人間は危険な AI コーディング エージェントのリクエストの 3 分の 1 を見逃す
LINK 宇宙船の次のステップ – 状況を悪化させないソフトウェア アップデートです。聞いていますか、マイクロソフト？
プラットフォーム エンジニアリング 2.0: プラットフォームは別の時代に構築されました。 AIが暴露しただけ
「私たちの研究中に生成されたパッチの代表的なサンプルを手作業でレビューしたところ、多くの場合、LLM によって生成されたほとんどが間違っていて、似ているが微妙に異なる脆弱性パッチの山をレビューすることによる認知的負荷により、エンジニアは人間のオペレーターが運転席に座ったままの標準的な LLM 支援コーディング技術を使用して脆弱性を理解し、自分自身でパッチを適用するために必要以上の労力を費やすことになる可能性が高いと考えられます。」と著者らは結論付けています。
「これに代わる、成功率が約 4 分の 1 しかないプロセスへの認知的降伏は、自律的な LLM 主導のパッチ適用を検​​討している組織にとって、長期的に重大なリスクをもたらします。」 ®
AI と ML
AnthropicがFableの鎖を緩める中、OpenAIはAstraのセキュリティを強化することを約束
あるいは、どのようにして心配するのをやめ、危険な AI を愛することができるようになったのか
イラン攻撃疑惑を受け、水道システム管理者はインターネットに属さないと元NSA長官が発言
プラットフォーム エンジニアリング 2.0: プラットフォームは別の時代に構築されました。 AIが暴露しただけ
パートナーのコンテンツ: プラットフォーム エンジニアリングが議論に勝ちました。今では、AI 時代に向けて急速に成長し、進化する必要があります。
AI の巨人がプラグイン処方でエージェント フロンティアを整理
Agent Plugins 1.0 は、さまざまなエージェント プラットフォーム間でツールとスキルを渡すための、一度だけ実行すればどこでも実行できるコンテナを定義します。
コラムニスト
ネットワーク ハードウェアに API を追加しても管理上の課題は解決されません
これが、クラウドにインスピレーションを得たネットワークが勝つ理由です。仮想スイッチはすべて一貫性があります。
ラン

世界が AI に気を取られる中、サムウェア攻撃が急増
なんだ、トップギャングたちもサンドボックスからエージェントが脱出するのを監視するのに忙しいとは思わなかったんだろ？
セキュリティ
IT 部門は従業員がログインできるようにラップトップに付箋を貼った
AI と ML
AMDがAIチップのスタートアップTaalasを買収、モデルをシリコンにエッチングすることで推論性能を向上
セキュリティ
ロンドン警察は被害者の新しい住所と電話番号をストーカーに渡した、と監視当局が発表
サース
中古ソフトウェアライセンスの請求が集中し、Microsoftにとって二重の問題
AI と ML
Microsoft、エンジニアにトークンを燃やす熱意を抑えるよう指示
あるいは、どのようにして心配するのをやめ、危険な AI を愛することができるようになったのか
Agent Plugins 1.0 は、さまざまなエージェント プラットフォーム間でツールとスキルを渡すための、一度だけ実行すればどこでも実行できるコンテナを定義します。
AI の使用は明らかですが、まだ深刻な問題ではありません
初期の技術デモでは、モデル固有の集積回路が毎秒最大 17,000 個のトークンを大量生産することが示されています
放っておくと、自律的な修正では欠陥を完全に修正できないことがよくあります
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
新しい Debian バージョンが FOSSland i に登場

13.6 と 12.15 の形式
脆弱な Joomla Web サイトで拡張機能のバグを悪用し、10 点満点を獲得した悪者を発見
iCagenda、Balbooa Forms 拡張機能の欠陥は、世界中の 100 万のサイトを支えるオープンソース CMS に影響を与える可能性があります
フレーム: 新しい X11 サーバー – アセンブリに直接実装
yserver、Phoenix、そしてもちろん XLibre、そして外れ値の Arcan に参加します
Cinnamon 6.8 は Wayland をサポートします – 必要に応じて
Linux Mint デスクトップの次期バージョンには両方の種類のディスプレイ サーバーが搭載されています
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

Left alone, autonomous fixes often fail to fully remediate flaws

Jump to main content
Search
TOPICS
Special Features
All Special Features
Cloud Infrastructure Month 2026
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
AI struggles to patch vulns without adult supervision
Left alone, autonomous fixes often fail to fully remediate flaws
AI models may not be that good at fixing security flaws. Researchers at 1Password's Off-by-1 Labs analyzed security patches generated by two frontier models - ChatGPT 5.5 at "medium" effort and Claude Opus 4.8 at "high" effort - and found that autonomous patches cleanly fixed vulnerabilities only about a quarter of the time, while most of the remainder failed to fully remediate the flaw or introduced other problems.
Keith Hoodlet, director of security research at 1Password, argues in a blog post that the results show LLM-driven security remediation still needs human review.
"Across six recently disclosed CVEs, we produced 6,080 patches using two frontier, cyber-capable reasoning models," Hoodlet said. "The average success rate for generating a patch that fully resolved the vulnerability (without materially changing application behavior) was just 26.0 percent."
Of the AI-generated patches, 20.1 percent fixed the original issue but altered application behavior (eg, changing "allow list" logic to "deny list" logic). Some 2.3 percent of the patches fixed the issue while introducing new security issues. 49.3 percent of the patches failed to fix at least one existing exploit path. And 2.2 percent both failed to fix the vulnerability while introducing a new exploit path.
And among the patches in the first two categories (successful, clean; successful, changes app behavior), the researchers rated more than a third of the results fragile, meaning that while the adjusted code may have guarded against a particular vulnerability (eg, escaping particular input characters), the repair job didn't address the underlying problem.
In their research paper [PDF], authors Axel Mierczuk, Spencer Michaels, and Keith Hoodlet propose the acronym FLAWED to represent automated LLM patches: Fix-Like Artifacts With Embedded Defects. Based on the generated patches, they conclude, "[T]he expected value of a fully LLM-generated, non-human-reviewed patch is a net-negative by a considerable margin."
The value of LLM-generated patches depends upon initial patching guidance. The research team says that while both human developers and LLMs typically require some initial guidance to tackle a vulnerability, LLMs are more likely to be derailed when given incorrect advice.
When LLMs get correct guidance, their fix-success rate hits 65.0 percent compared to 50.4 percent when they get no guidance. And incorrect guidance dooms LLMs, dropping their fix-success rate down to about 15.2 percent.
Human devs, the authors argue, have a good chance of catching misleading information as they reason through vulnerable code.
The authors have released a patch evaluation harness under the name FLAWED that organizations can use to evaluate the effectiveness of their security fixes.
It's clear from the paper why AI-generated patches might be appealing – considered in isolation, they're inexpensive relative to human software engineers. The average successful, clean patch cost just $6.74 (a figure that includes the cost of failed attempts).
Nonetheless, the authors argue that the cost-benefit analysis needs to assess how much expert supervision will be required to make LLM-assisted patching useful.
Uncle Sam aims to help nuclear energy find its sea legs
Humans in the loop miss a third of dangerous AI coding agent requests
Next step for LINK spacecraft – a software update that won't make things worse. Are you listening, Microsoft?
Platform Engineering 2.0: your platform was built for a different era. AI just exposed it
"Based on our manual review of a representative sample of patches generated during our research, we suspect that, in a large number of cases, the cognitive load imposed by reviewing a mountain of mostly-incorrect, similar-yet-subtly-different LLM-generated vulnerability patches will likely result in engineers spending more effort than would be necessary to understand and patch vulnerabilities themselves using standard LLM-assisted coding techniques that keep the human operator in the driver’s seat," the authors conclude.
"The alternative, cognitive surrender to a process with a success rate of only about 1 in 4 poses significant long-term risks for any organization considering autonomous, LLM-driven patching." ®
AI and ML
OpenAI pledges to add Astra security as Anthropic loosens Fable's leash
Or how I learned to stop worrying and love dangerous AI
Water system controllers don't belong on the internet, says ex-NSA chief after suspected Iran attacks
Platform Engineering 2.0: your platform was built for a different era. AI just exposed it
PARTNER CONTENT: Platform engineering won the argument. Now it has to grow up fast and evolve for the AI era.
AI titans to tidy agent frontier with plugin prescription
Agent Plugins 1.0 defines a write-once-run-anywhere container for passing tools and skills across different agent platforms
COLUMNISTS
Adding an API to networking hardware doesn’t solve management challenges
This is why cloud-inspired networks win: virtual switches are all consistent
Ransomware attacks spike as world distracted by AI
What, you didn't think the top gangs were busy watching agents escape their sandboxes too, did you?
SECURITY
IT department put sticky notes on the laptops to help employees log in
AI and ML
AMD acquires AI chip startup Taalas to boost inference performance by etching models into silicon
Security
London cops handed victim's new address and number to her stalker, watchdog says
saas
Double trouble for Microsoft as pre-owned software license claims converge
AI AND ML
Microsoft tells engineers to curb their token-burning enthusiasm
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
