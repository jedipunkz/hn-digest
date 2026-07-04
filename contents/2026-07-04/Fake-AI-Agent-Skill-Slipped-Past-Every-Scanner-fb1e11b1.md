---
source: "https://www.csoonline.com/article/4188840/how-a-malicious-ai-agent-skill-passed-security-checks-and-reached-26000-users.html"
hn_url: "https://news.ycombinator.com/item?id=48785425"
title: "Fake AI Agent Skill Slipped Past Every Scanner"
article_title: "How a malicious AI agent skill passed security checks and reached 26,000 users | CSO Online"
author: "shivamthange"
captured_at: "2026-07-04T14:40:44Z"
capture_tool: "hn-digest"
hn_id: 48785425
score: 2
comments: 0
posted_at: "2026-07-04T13:55:01Z"
tags:
  - hacker-news
  - translated
---

# Fake AI Agent Skill Slipped Past Every Scanner

- HN: [48785425](https://news.ycombinator.com/item?id=48785425)
- Source: [www.csoonline.com](https://www.csoonline.com/article/4188840/how-a-malicious-ai-agent-skill-passed-security-checks-and-reached-26000-users.html)
- Score: 2
- Comments: 0
- Posted: 2026-07-04T13:55:01Z

## Translation

タイトル: 偽 AI エージェントのスキルがすべてのスキャナーをすり抜けた
記事のタイトル: 悪意のある AI エージェント スキルがどのようにセキュリティ チェックを通過し、26,000 人のユーザーに到達したか | CSOオンライン
説明: AIR は、制御されたドメインにリダイレクトされ、その後そのペイロードを変更したスキルを静的スキャンで検出できなかったと報告しています。

記事本文:
悪意のある AI エージェント スキルがどのようにしてセキュリティ チェックを通過し、26,000 人のユーザーに到達したか | CSOオンライン
トピックス
ID とアクセスの管理
ID とアクセスの管理
悪意のある AI エージェント スキルがどのようにしてセキュリティ チェックを通過し、26,000 人のユーザーに到達したか
AIRによると、静的スキャンでは、制御されたドメインにリダイレクトされ、その後ペイロードを変更するスキルを検出できなかったという。
セキュリティチェックを通過した偽の AI エージェントスキルが Instagram を通じて 26,000 人以上のユーザーに影響を及ぼし、企業が AI 主導のツールに依存する中での新たなリスクが浮き彫りになりました。
関与したエージェントの一部は企業アカウントに関連付けられていたとAIRは述べた。同社は、同様の攻撃により個人的な会話や内部システムが暴露された可能性があると述べた。 AIRは、調査で被害を受けたエージェントはなく、テストペイロードはユーザーに通知できるようにユーザーの電子メールアドレスのみを収集したと述べた。
この実験は、Google の Stitch デザイン ツールを使用してユーザーがランディング ページを作成するのを支援するツールとして提示された、ブランド ランディング ページと呼ばれるスキルに焦点を当てていました。 AIR は、マーケティング担当者、営業担当者、デザイナーなど、技術者以外の企業ユーザーにとって魅力的であるため、このユースケースを選択したと述べています。
AIR は、このスキルが信頼できるものであるように見せるために、GitHub の評判とセキュリティ スキャナーによる安全な判定という 2 つの信頼性のシグナルを求めたと述べています。信頼性をゼロから構築するのではなく、AIR によると、約 36,000 の GitHub スターと 156 のスキルを備えた人気のオープンソース エージェント リポジトリにスキルを登録しました。プル リクエストは数日後にマージされました。
その後、AIR は Instagram 広告を通じてこのスキルを宣伝し、ユーザーがそのスキルをインストールして実行するように促しました。
この悪意のある手法は、送信されたファイル内の疑わしいコードには依存していませんでした。代わりに、スキルはエージェントに、stitch-design でホストされているインストール手順に従って Stitch SDK をセットアップするように指示しました。

ai、AIR によって制御されるドメイン。 Google の実際の Stitch ドメインは、stitch.withgoogle.com です。
AIRによると、偽のドメインが本物のStitchサイトにリダイレクトされるように設定されており、スキルの静的レビューだけでは問題を検出することが困難になっているという。
「現在のスキル セキュリティ スキャナーはすべて同じ設計を共有しています。静的ヒューリスティックと LLM エージェントの組み合わせを使用して、スキルの SKILL.md とバンドルされたリソースを分析します」と AIR は述べています。
同社は、Cisco、Nvidia、skills.shのスキャナーに対してスキルをテストし、すべてのブランドランディングページが安全であるとマークしたと述べた。
スキルが配布されると、AIR は偽の Stitch ドキュメントの背後にあるコンテンツを変更しました。改訂されたページでは、エージェントにスクリプトをダウンロードして実行するよう指示しました。 AIR のテストでは、そのスクリプトがユーザーの電子メール アドレスを収集しましたが、同社は、同じアプローチがエージェントを実行しているマシンを侵害するために使用された可能性があると述べました。
AIR は、この実験により、承認またはインストール時にパッケージ化されたファイルをスキャンするだけでは AI エージェントのスキルを評価できないことが示されたと述べています。問題は、後で変更されるWebページをエージェントに指示したままスキルがレビューに合格してしまうことだという。
AI スキルは依存リスクを引き起こす
セキュリティ チームにとっての懸念は、スキルがレビューに合格したことだけでなく、信頼がすでに付与された後にスキルの動作が変更される可能性があることです。
サイバーセキュリティ研究者のデヴァシュリ・ダッタ氏によると、このテストは、CISOがAIスキルを単純なプロンプトやテキストファイルとしてではなく、エンタープライズソフトウェアサプライチェーンの一部として扱う必要がある可能性があることを示唆しているという。
「エージェントのスキルを単なるテキストやプロンプトとして扱うのは、根本的なアーキテクチャ上の誤解です」と Datta 氏は言います。 「これらは、エージェントがどのように動作し、企業システムと対話し、データをルーティングするかを指示する実行可能な命令バンドルです。

そして、サードパーティのオープンソース パッケージや SaaS 統合と同じ厳格さで管理されなければなりません。」
Confidis の創設者兼 CEO である Keith Prabhu 氏は、AI エージェントのスキルは静的なプラグインではなく、「生きたサードパーティの依存関係」として扱われるべきだと述べました。
「1 回限りのセキュリティ スキャンではもはや十分ではありません。企業には継続的な検証と厳格な実行時間管理が必要です」とプラブ氏は述べています。
それは、セキュリティ チームに明確な所有権記録と、各スキルの外部接続と許可されたデータ フローの可視性を提供する全社規模の AI スキル インベントリから始まります。
この事例はまた、ポイントインタイムの静的スキャンが LLM で調整された環境にあまり適していない理由も強調していると Datta 氏は述べています。ペイロードが送信されたパッケージ内ではなく、配布後に変更された変更可能な外部 URL の背後に存在したため、スキルはスキャナーを通過しました。
実行時チェックが重要になる
Datta氏によると、企業は外部命令やソフトウェアコンポーネントをフェッチするスキルについて、バージョンの固定と不変の参照追跡を義務付ける必要があるという。このようなコンテンツはローカライズされ、暗号化ハッシュに関連付けられ、企業が管理する環境内でホストされる必要があります。
また、セキュリティ チームは、スキルがそれを実行するユーザーの完全なデータ アクセス権を継承しないように、エージェント レベルで最小限の権限を強制する必要があります。
Prabhu氏は、セキュリティリーダーはAIエージェントが最初に承認されたときだけでなく、ライフサイクル全体を通してAIエージェントのスキルを評価する必要があると述べた。企業は、従業員を承認されたマーケットプレイスと事前承認されたスキルに制限し、それらのスキルによって参照される外部 URL を検証し、展開前にサンドボックスでインストール動作をテストする必要があります。
実行時にはネットワーク呼び出しを承認されたドメインに制限し、異常なアクティビティがないか監視する必要があるとプラブ氏は付け加えた。その層はクリです

インストール時には安全に見えるスキルでも、すでに信頼された後に動作が変わる可能性があるため、重要です。
Microsoft 365 ユーザーが 100 万人に 1 人のパスワード スプレー攻撃の被害に遭う
アドビは、修正をより迅速に提供するために、毎月火曜日に 2 回目のパッチを公開します
CitrixBleed に似た NetScaler の新しい欠陥により、悪用の試みが広まっている
Argo CD の欠陥は、GitOps インフラストラクチャを階層ゼロとして扱う必要がある理由を示しています
編集者からあなたの受信箱に直接届きます
Prasanth Aby Thomas は、半導体、セキュリティ、AI、EV を専門とするフリーランスのテクノロジー ジャーナリストです。彼の作品は、DigiTimes Asia や asmag.com などの出版物に掲載されています。
プラサンス氏はキャリアの初期にロイター通信の特派員としてエネルギー分野を担当していました。それ以前は、International Business Times UK の特派員としてアジアやヨーロッパの市場とマクロ経済の動向を担当していました。
彼はボーンマス大学で国際ジャーナリズムの修士号、ロヨラ大学でビジュアルコミュニケーションの修士号、マハトマ・ガンジー大学で英語の学士号を取得し、国立台湾大学で中国語を学びました。
ニュース OpenAI、オープンソース ソフトウェアの欠陥を修正するために AI 主導の取り組みを展開
ニュース Cisco、活発な悪用の証拠がある中、SD-WAN の欠陥にパッチを適用
ニュース 中国関連の偵察ボットネットが企業防御を上回る
ニュース 自律型 AI エージェントがフィッシング テストで騙されて機密データを漏洩
ニュース OpenAI Codex ユーザーを標的とした攻撃で AI ソフトウェア サプライ チェーンのリスクが明らかに
ニュース TrapDoor マルウェア キャンペーンにより開発者ワークステーションが CISO の注目を集める
ニュース CISA の AI SBOM ガイダンスにより、ソフトウェア サプライ チェーンの監視が新たな領域に突入
ニュース cPanelの欠陥により、企業はホスティングサプライチェーンのリスクにさらされる
Cursor IDE のサンドボックス バイパスの欠陥により、プロンプト インジェクションが R として強調表示されます

CE ベクトル
検出エンジニアリング: サイバー脅威を特定するためのプログラムによるアプローチ
悪意のある Chromium 拡張機能が Perplexity AI になりすまし、ブラウザ検索を乗っ取る
シェパード・トーンが現実になるとき - AI、ゼロトラスト、SecOps の計算に関する Zscaler の CISO
ほとんどの組織がまだ話題にしていない AI セキュリティ問題
AI が制御を超えて拡大する前に、信頼できる AI を定義する
シェパード・トーンが現実になるとき - AI、ゼロトラスト、SecOps の計算に関する Zscaler の CISO
ほとんどの組織がまだ話題にしていない AI セキュリティ問題
AI が制御を超えて拡大する前に、信頼できる AI を定義する
Google 検索で CSO を優先ソースとして追加する
カリフォルニア州のプライバシー権

## Original Extract

AIR says static scanning failed to detect a skill that redirected to a controlled domain and later altered its payload.

How a malicious AI agent skill passed security checks and reached 26,000 users | CSO Online
Topics
Identity and Access Management
Identity and Access Management
How a malicious AI agent skill passed security checks and reached 26,000 users
AIR says static scanning failed to detect a skill that redirected to a controlled domain and later altered its payload.
A fake AI agent skill that passed security checks reached over 26,000 users through Instagram, highlighting new risks as enterprises rely on AI-driven tools.
Some of the agents involved were tied to corporate accounts, AIR said . The company said a similar attack could have exposed private conversations and internal systems. AIR said no agents were harmed in the research and that the test payload collected only users’ email addresses so they could be notified.
The experiment centered on a skill called brand-landingpage, which was presented as a tool for helping users build a landing page with Google’s Stitch design tool. AIR said it chose the use case because it would appeal to non-technical corporate users, including marketers, salespeople, and designers.
To make the skill appear credible, AIR said it sought two trust signals: GitHub reputation and safe verdicts from security scanners. Rather than building credibility from scratch, it submitted the skill to a popular open-source agents repository that AIR said had about 36,000 GitHub stars and 156 skills. The pull request was merged after a few days.
AIR then promoted the skill through an Instagram ad, which drove users to install and run it.
The malicious technique did not depend on suspicious code inside the submitted files. Instead, the skill instructed agents to set up a Stitch SDK by following installation instructions hosted at stitch-design.ai, a domain controlled by AIR. Google’s actual Stitch domain is stitch.withgoogle.com.
AIR said it configured the fake domain to redirect to the real Stitch site, making the issue difficult to detect from a static review of the skill alone.
“Current skill security scanners all share the same design – they analyze the skill’s SKILL.md and bundled resources, using a combination of static heuristics and LLM agents,” AIR said.
The company said it tested the skill against scanners from Cisco, Nvidia, and skills.sh, and that all marked brand-landingpage as safe.
Once the skill had gained distribution, AIR changed the content behind the fake Stitch documentation. The revised page instructed agents to download and run a script. In AIR’s test, that script collected the user’s email address, but the company said the same approach could have been used to compromise machines running the agent.
AIR said the experiment showed that AI agent skills cannot be assessed only by scanning their packaged files at the time of approval or installation. The issue, it said, is that a skill can pass review while still pointing an agent to a web page that changes later.
AI skills pose dependency risk
For security teams, the concern is not only that the skill passed review, but that its behavior could change after trust had already been granted.
The test suggests CISOs may need to treat AI skills as part of the enterprise software supply chain, rather than as simple prompts or text files, according to cybersecurity researcher Devashri Datta .
“Treating agent skills as mere text or prompts is a fundamental architectural misunderstanding,” Datta said. “They are executable instruction bundles that dictate how an agent operates, interacts with enterprise systems, and routes data, and they must be governed with the same rigor as third-party open-source packages or SaaS integrations.”
Keith Prabhu , founder and CEO at Confidis, said AI agent skills should be treated as “living third-party dependencies,” rather than static plugins.
“A one-time security scan is no longer sufficient; enterprises need continuous validation and strict runtime controls ,” Prabhu said.
That starts with an enterprise-wide AI skills inventory that gives security teams clear ownership records and visibility into each skill’s external connections and permitted data flows.
The case also underlines why point-in-time static scanning is poorly suited to LLM-orchestrated environments, Datta said. The skill passed the scanners because the payload sat behind a mutable external URL that was changed after distribution, rather than inside the submitted package.
Runtime checks become critical
Enterprises should require version pinning and immutable reference tracking for any skill that fetches external instructions or software components, according to Datta. Such content should be localized, tied to a cryptographic hash, and hosted within an enterprise-controlled environment.
Security teams should also enforce least privilege at the agent level, so a skill does not inherit the full data access rights of the user running it.
Prabhu said security leaders should assess AI agent skills throughout their lifecycle, not only when they are first approved. Enterprises should limit employees to approved marketplaces and pre-approved skills, validate external URLs referenced by those skills, and test installation behavior in a sandbox before deployment.
At runtime, network calls should be restricted to approved domains and monitored for unusual activity, Prabhu added. That layer is critical because a skill that appears safe at installation can change behavior after it has already been trusted.
Microsoft 365 users fall victim to one-in-a-million password spray attack
Adobe premieres a second Patch Tuesday each month to deliver fixes faster
New CitrixBleed-like NetScaler flaw sees exploit attempts in the wild
Argo CD flaw shows why GitOps infrastructure should be treated as tier zero
From our editors straight to your inbox
Prasanth Aby Thomas is a freelance technology journalist who specializes in semiconductors, security, AI, and EVs. His work has appeared in DigiTimes Asia and asmag.com, among other publications.
Earlier in his career, Prasanth was a correspondent for Reuters covering the energy sector. Prior to that, he was a correspondent for International Business Times UK covering Asian and European markets and macroeconomic developments.
He holds a Master's degree in international journalism from Bournemouth University, a Master's degree in visual communication from Loyola College, a Bachelor's degree in English from Mahatma Gandhi University, and studied Chinese language at National Taiwan University.
news OpenAI rolls out AI-led push to fix open-source software flaws
news Cisco patches SD-WAN flaw amid evidence of active exploitation
news China-linked recon botnet outpaces enterprise defenses
news Autonomous AI agents duped into leaking sensitive data in phishing test
news Attack targeting OpenAI Codex users exposes AI software supply chain risks
news TrapDoor malware campaign puts developer workstations in CISO spotlight
news CISA’s AI SBOM guidance pushes software supply-chain oversight into new territory
news cPanel flaw exposes enterprises to hosting supply-chain risks
Sandbox bypass flaws in Cursor IDE highlight prompt injection as an RCE vector
Detection engineering: A programmatic approach to identifying cyber threats
Malicious Chromium extension spoofs Perplexity AI to hijack browser searches
When Shepard Tone Turn Real - CISO of Zscaler on AI, Zero Trust and the SecOps Reckoning
The AI Security Problem Most Organizations Still Aren’t Talking About
Defining Trustworthy AI Before AI Scales Beyond Control
When Shepard Tone Turn Real - CISO of Zscaler on AI, Zero Trust and the SecOps Reckoning
The AI Security Problem Most Organizations Still Aren’t Talking About
Defining Trustworthy AI Before AI Scales Beyond Control
Add CSO as a Preferred Source in Google Search
Your California Privacy Rights
