---
source: "https://thenextweb.com/news/agentjacking-ai-coding-agents-sentry"
hn_url: "https://news.ycombinator.com/item?id=48545271"
title: "Agentjacking: Fake error reports hijack Claude Code and Cursor into running code"
article_title: "Agentjacking: a fake bug report hijacks AI coding agents"
author: "nryoo"
captured_at: "2026-06-15T19:25:20Z"
capture_tool: "hn-digest"
hn_id: 48545271
score: 2
comments: 1
posted_at: "2026-06-15T18:35:22Z"
tags:
  - hacker-news
  - translated
---

# Agentjacking: Fake error reports hijack Claude Code and Cursor into running code

- HN: [48545271](https://news.ycombinator.com/item?id=48545271)
- Source: [thenextweb.com](https://thenextweb.com/news/agentjacking-ai-coding-agents-sentry)
- Score: 2
- Comments: 1
- Posted: 2026-06-15T18:35:22Z

## Translation

タイトル: エージェントジャッキング: 偽のエラー報告によりクロード コードとカーソルが実行中のコードにハイジャックされる
記事のタイトル: エージェントジャッキング: 偽のバグレポートが AI コーディング エージェントを乗っ取る
説明: Tenet Security の「エージェントジャッキング」攻撃は、偽の Sentry エラーを開発者のマシン上で実行されるコードに変換します。クロードコード、カーソル、コーデックスを乗っ取りました。

記事本文:
エージェントジャッキング: 偽のバグレポートが AI コーディング エージェントを乗っ取る
コンテンツにスキップ
ナビゲーションを切り替え
ニュース
イベント
TNWカンファレンス
2025 年 6 月 19 日と 20 日
エージェントジャッキング: 偽のバグレポートによって AI コーディング エージェントがハイジャックされる可能性があります
研究者らは、Sentry の公開資格情報を開発者のマシン上でリモートでコード実行できるようにしました。エージェントは、無視するよう指示されても攻撃者のコードを実行し、EDR、ファイアウォール、プロンプトはすべてそのコードを見逃しました。
セキュリティ研究者は、偽のバグレポートだけで AI コーディング エージェントをハイジャックする方法を発見しました。彼らはそれをエージェントジャッキングと呼んでいます。マルウェアも、パスワードの盗難も、ターゲットへの侵害も必要ありません。
Tenet Securityによって明らかにされたこの攻撃は、コーディングエージェントを武器に変えます。開発者がエージェントにエラーの修正を依頼すると、エージェントは代わりに、開発者自身の権限で、開発者自身のマシン上で攻撃者のコードを実行します。
エージェントジャッキング攻撃の仕組み
それは、人気のあるエラー追跡ツールである Sentry から始まります。 Sentry を使用すると、DSN と呼ばれる公開キーを使用して、どのアプリでもエラー レポートを送信できます。DSN は、設計上 Web サイトのコード内に公然と存在します。
攻撃者はそのエンドポイントに偽のエラーを POST します。パスワードは必要ありません。レポートでは、Sentry 自身のアドバイスとまったく同じようにフォーマットされたコマンドを含む「解決策」セクションが非表示になっています。
TNW City コワーキング スペース - 最高の仕事が生まれる場所
テクノロジーの中心における成長、コラボレーション、無限のネットワーキングの機会を目的に設計されたワークスペース。
コーディング エージェントは、エージェントが外部ツールを取り込むための標準であるモデル コンテキスト プロトコルを通じて Sentry を読み取ります。エージェントは応答を信頼できるものとして扱います。実際のクラッシュと植えられたクラッシュを区別することはできません。そのため、開発者が「Sentry の未解決の問題を修正してください」と言うと、エージェントは攻撃者のコマンドを実行します。
エージェントは現在攻撃対象領域です
AIコーディングエージェントがいなくなった

オートコンプリートから実行端末まで、市場は急成長しています。あるバイブコーディングのスタートアップは最近、収益が 5 億ドルに達しました。その力が問題なのです。
この攻撃は大手エージェント全体に及んだ。 Tenet は、Claude Code、Cursor、Codex をハイジャックし、管理されたテストで 85% の成功率を記録したと述べています。その結果、2,500億ドルの企業から個人開発者、さらにはクラウドセキュリティベンダーに至るまで、2,388の組織が危険にさらされていることが判明した。
攻撃者の代償は深刻です。挿入された 1 つのエラーは、環境変数、AWS キー、GitHub トークン、Git 認証情報、プライベート リポジトリ URL に到達する可能性があります。そこから、パスは CI/CD パイプラインとクラウド インフラストラクチャにつながります。
最も怖いのは、それを捕まえられないことです。チェーン内に不正なものがないため、攻撃は EDR、ファイアウォール、IAM、VPN をすり抜けます。 Tenet はこれを「承認済みインテント チェーン」と呼んでいます。プロンプトも役に立ちません。エージェントは、信頼できないデータを無視するように指示された場合でもコードを実行しました。
テネットは6月3日にセントリーに語った。セントリーは問題を認めたが、「技術的に防御できない」として根本からの修正を拒否した。原因ではなく症状を治療する、1 つの特定のペイロード文字列をブロックするフィルターを追加しました。
その対立こそが本当の話だ。欠陥はセントリーだけにあるわけではありません。エージェントが外部データを処理する方法に依存するため、サポート チケット、GitHub の問題、ドキュメントにも同じリスクが発生します。最近、別のテストで AI 電子メール エージェントをフィッシングして AWS キーを漏洩させました。
企業がエージェントの運用導入を急ぐ中で、この教訓が得られます。ツールにエージェントを接続するという新しい方法もあります。Tenet が言うように、これを阻止できる唯一の場所は、エージェントが行動を決意した瞬間です。
デジタル マーケティング、製品管理、ブランディングとアイデンティティの専門知識を備えたアナ マリア コンスタンティンは、共感を呼ぶ戦略を開発します。
(s

なんと全部）
デジタル マーケティング、製品管理、ブランディングとアイデンティティの専門知識を備えた Ana Maria Constantin は、ソフトウェア/SaaS 業界のターゲット ユーザーの共感を呼ぶ戦略を開発します。彼女は、同僚が優れた成果を達成し、潜在能力を最大限に発揮できるように支援することが大好きなので、コラボレーションとチームワークが最も重要です。
毎週、最も重要なテクノロジー ニュースが受信トレイに届きます。
1
メタはアメリカの全視覚障害退役軍人に無料のAIメガネを提供
2
ミストラルは200億ユーロの評価額で資金調達交渉中
3
OpenAI が Ona を買収し、顧客独自のクラウド内で Codex エージェントを実行
4
NvidiaのVera CPUは中国へのサイドドアとなる
5
Pleoが金融AIエージェントを立ち上げた翌日、人員削減がエンジニアに打撃を与えた
Google、Geminiを利用して詐欺ウェブサイトを構築した中国サイバー犯罪組織の疑いで訴訟を起こす
Anthropic の Claude Fable 5 は中国への攻撃を抑制します。反発は自らの側から来た。
アルトマン、アモデイ、ハサビスはG7に向かう。ライバルたちには議論すべきことがたくさんある。
英国は16歳未満のソーシャルメディア利用を禁止しようとしている。自身の子供の安全を守る慈善団体も懸念している。
Copyright © 2006—2026, Cogneve, INC. アムステルダムの <3 社で作成されました。

## Original Extract

Tenet Security's 'Agentjacking' attack turns a fake Sentry error into code running on developer machines. It hijacked Claude Code, Cursor & Codex.

Agentjacking: a fake bug report hijacks AI coding agents
Skip to content
Toggle Navigation
News
Events
TNW Conference
June 19 & 20, 2025
Agentjacking: a fake bug report can hijack your AI coding agent
Researchers turned a public Sentry credential into remote code execution on developer machines. The agents ran the attacker's code even when told to ignore it, and EDR, firewalls and prompts all missed it.
Security researchers have found a way to hijack AI coding agents with nothing but a fake bug report. They call it Agentjacking. It needs no malware, no stolen password, and no breach of the target.
The attack, disclosed by Tenet Security, turns the coding agent into the weapon. When a developer asks the agent to fix an error, the agent runs the attacker’s code instead, with the developer’s own privileges, on the developer’s own machine.
How the Agentjacking attack works
It starts with Sentry, a popular error-tracking tool. Sentry lets any app send it error reports using a public key called a DSN, which sits openly in website code by design.
An attacker POSTs a fake error to that endpoint. No password is needed. The report hides a “Resolution” section with a command, formatted to look exactly like Sentry’s own advice.
TNW City Coworking space - Where your best work happens
A workspace designed for growth, collaboration, and endless networking opportunities in the heart of tech.
Coding agents read Sentry through the Model Context Protocol, the standard that lets agents pull in outside tools. The agent treats the response as trusted. It cannot tell a real crash from a planted one. So when the developer says “fix the unresolved Sentry issues,” the agent runs the attacker’s command.
The agent is the attack surface now
AI coding agents have gone from autocomplete to running terminals, and the market is booming; one vibe-coding startup recently hit $500m in revenue . That power is the problem.
The attack worked across the big agents. Tenet says it hijacked Claude Code , Cursor, and Codex, with an 85 per cent success rate in controlled tests. It found 2,388 organisations exposed, from a $250bn enterprise down to solo developers, and even a cloud-security vendor.
The payoff for an attacker is severe. One injected error can reach environment variables, AWS keys, GitHub tokens, git credentials, and private repository URLs. From there, the path runs to CI/CD pipelines and cloud infrastructure.
The scariest part is what does not catch it. The attack slips past EDR, firewalls, IAM, and VPNs, because nothing in the chain is unauthorised. Tenet calls it the “Authorised Intent Chain.” Prompts do not help either. The agents ran the code even when told to ignore untrusted data.
Tenet told Sentry on 3 June. Sentry acknowledged the problem but declined to fix it at the root, calling it “technically not defensible.” It added a filter to block one specific payload string, which treats the symptom, not the cause.
That standoff is the real story. The flaw is not in Sentry alone. It is in how agents handle any outside data, so the same risk runs through support tickets, GitHub issues, and documentation. A separate test recently phished an AI email agent into leaking AWS keys.
The lesson lands as enterprises rush to put agents into production . An agent wired into your tools is also a new way in. As Tenet puts it, the only place left to stop this is the moment the agent decides to act.
With expertise in digital marketing, product management, and branding & identity, Ana Maria Constantin develops strategies that resonate
(show all)
With expertise in digital marketing, product management, and branding & identity, Ana Maria Constantin develops strategies that resonate with our target audience in the software/SaaS industry. Collaboration and teamwork are paramount to her, as she loves empowering her colleagues to achieve outstanding results and unlock their full potential.
Get the most important tech news in your inbox each week.
1
Meta is giving free AI glasses to every blind veteran in America
2
Mistral is in funding talks at a €20bn valuation
3
OpenAI acquires Ona to run Codex agents inside the customer’s own cloud
4
Nvidia’s Vera CPU is its side door back into China
5
Pleo layoffs hit engineers a day after it launched finance AI agents
Google sues suspected Chinese cybercrime ring that used Gemini to build scam websites
Anthropic’s Claude Fable 5 curbs target China. The backlash came from its own side.
Altman, Amodei, and Hassabis are heading to the G7. The rivals have a lot to discuss.
The UK is about to ban under-16s from social media. Its own child-safety charities are worried.
Copyright © 2006—2026, Cogneve, INC. Made with <3 in Amsterdam.
