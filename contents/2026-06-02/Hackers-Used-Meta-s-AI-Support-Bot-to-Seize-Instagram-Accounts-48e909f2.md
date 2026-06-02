---
source: "https://krebsonsecurity.com/2026/06/hackers-used-metas-ai-support-bot-to-seize-instagram-accounts/"
hn_url: "https://news.ycombinator.com/item?id=48361346"
title: "Hackers Used Meta's AI Support Bot to Seize Instagram Accounts"
article_title: "Hackers Used Meta’s AI Support Bot to Seize Instagram Accounts – Krebs on Security"
author: "panarky"
captured_at: "2026-06-02T00:33:34Z"
capture_tool: "hn-digest"
hn_id: 48361346
score: 51
comments: 18
posted_at: "2026-06-01T19:18:52Z"
tags:
  - hacker-news
  - translated
---

# Hackers Used Meta's AI Support Bot to Seize Instagram Accounts

- HN: [48361346](https://news.ycombinator.com/item?id=48361346)
- Source: [krebsonsecurity.com](https://krebsonsecurity.com/2026/06/hackers-used-metas-ai-support-bot-to-seize-instagram-accounts/)
- Score: 51
- Comments: 18
- Posted: 2026-06-01T19:18:52Z

## Translation

タイトル: ハッカーが Meta の AI サポート ボットを使用して Instagram アカウントを押収
記事タイトル: ハッカーが Instagram アカウントを押収するために Meta の AI サポート ボットを使用 – セキュリティに関するクレブス

記事本文:
ハッカーが Meta の AI サポート ボットを使用して Instagram アカウントを押収 – Krebs のセキュリティ
-->
広告
広告
-->
コンテンツにスキップ
ホーム
ハッカーがMetaのAIサポートボットを利用してInstagramアカウントを強奪
メタ社の「AIサポートアシスタント」ボットを騙してアカウントのパスワードをリセットさせる方法を示す指示がテレグラム上に出回った後、オバマ大統領と米国宇宙軍首席曹長のインスタグラムアカウントが先週末、親イラン的な画像やメッセージで一時的に改ざんされた。
Telegram で公開されたビデオのスクリーンショット。Meta の AI カスタマー サポート ボットがどのように騙されてターゲットのパスワードをリセットできるかを示すと主張しています。
5 月 31 日、Meta の AI ボットが、ボットの標準パスワード リセット フローの一環として既存のアカウントに電子メール アドレスを喜んで追加するという噂が、いくつかの Telegram インスタント メッセージ チャネルで広がり始めました。
親イランハッカーがテレグラムで公開したビデオは、ターゲットの普段の居住地内またはその近くにある IP アドレスで VPN 接続を使用し、アカウントのパスワードのリセットを要求し、メタの AI サポート アシスタントとチャットすることを選択するという、非常に単純なエクスプロイトを記録していると主張しました。そこからビデオでは、攻撃者がボットに問題のアカウントを新しい電子メール アドレスにリンクするように指示し、その後ボットがそのアドレスにパスワードのリセットを可能にするワンタイム コードを律儀に送信したことが示されています。
動画を投稿したテレグラムアカウントは、ハッキングされたインスタグラムアカウントを改ざんした親イラン画像、動画、メッセージのスクリーンショットにもリンクしており、ハッカーがこのエクスプロイトを利用して、再販価値が50万ドル以上あるとされる多数の貴重な（略称）インスタグラムアカウント名をハイジャックしたと述べた。
Meta は r に応答しませんでした

動画の申し立てについてコメントを求めたが、Metaのアンディ・ストーン氏はTwitter/Xで、問題は解決され、影響を受けたアカウントを保護していると述べた。セキュリティブログthecybersecguru.comは、Metaが週末に緊急パッチをプッシュし、バックエンドデータベースが侵害されていないことを明らかにしたと報告している。
「Instagramの人的サポートインフラが貧弱であることで悪名高い」とCyber​​secguruは書いている。 「ロックされたアカウントの回復、特に価値の高いアカウントの回復には、自動チケット発行システムとのやりとりに何週間もかかることがあります。Meta の解決策は、失われた電子メール アドレスの再リンク、パスワード リセットのトリガー、アカウント所有権の確認など、一般的な回復ワークフローを処理する会話型 AI レイヤーを導入することでした。おそらく、このアシスタントは、アカウント アクセス地獄に陥った正当なユーザーの負担を軽減するはずでした。」
Lumen の Black Lotus Labs の脅威研究者である Ian Goldin 氏は、より大規模なオンライン プラットフォームが AI チャットボットによる機密性の高いアカウント回復リクエストの処理を許可し始めているため、私たちは未知のセキュリティ領域に入りつつあると述べました。人間のカスタマー サポートの従業員がソーシャル エンジニアリングされ、誰かのアカウントに不正アクセスを提供されるのと同じように、AI ボットも同様に助けようとするが、説得や策略には弱い、と同氏は述べた。
「AI チャットボットは興味深い新たな攻撃対象領域を生み出しており、この種の攻撃は今後さらに多く発生する可能性があります」とゴールディン氏は述べています。
さまざまなオンライン アカウントを保護するということは、提供されている最も安全な形式の多要素認証 (MFA) (パスキーやセキュリティ キーなど) を最大限に活用することを意味します。この場合、Instagram が提供する最も堅牢性の低い MFA 形式 (SMS 経由で送信されるワンタイム コード) を使用した場合でも、エクスプロイトはブロックされた可能性があります。テレグラムでビデオを公開したハッカーは、エクスプロイトを行ったと述べています。

OIT は、MFA が有効になっているアカウントに対しては機能しませんでした。
このエントリは、2026 年 6 月 1 日月曜日午後 1 時 32 分に投稿されました。
「ハッカーが Meta の AI サポートボットを使用して Instagram アカウントを押収」 への 9 件のフィードバック
シリアル 2026 年 6 月 1 日
すべてを一言で表すと「メタ」
CyberScrewed 2026 年 6 月 1 日
AI は攻撃対象領域を飛躍的に拡大するだけです。素晴らしい仕事だ
技術者志望 2026 年 6 月 1 日
はい、誰もがそれに夢中になっています。
食事 2026 年 6 月 1 日
それが生まれたとされる私の地域の広い範囲で非常に人気がありません。
今では、LinkedIn のような「必要なビジネス悪」が増えています。 AI があなたを雇用し、AI があなたを監視し、AI があなたを解雇します。注意してください。AI や、AI が世界を食い荒らす「営利目的!…」データセンターについて否定的な意見を表明する人には、さらなる監視の目が向けられます。そのため、必ずカメラに向かって微笑み、無能な代替理論を唱えるロボットの支配者たちにチップを渡すようにしてください。彼らはあなた方の継続的な有用性と、現在の人間の肉の価格で同等の価値の代替が差し迫った瞬間を監視し、評価しています。そして、もしそれが「好き」でないなら、あなたはラッダイトか何かに違いありません。 Google のプロジェクト パープルにとって潜在的な脅威。
スティーブ 2026 年 6 月 1 日
たとえメタのようなものに対するセキュリティへの期待が低いとしても、これは完全に許しがたいセキュリティ上の欠陥です。
サーシャ 2026 年 6 月 1 日
いやー、なんてタイムリーなんでしょう。先週、IG からアカウントにアクセスできず申し訳ありませんという内容のメールを数通受け取りました。誰かが明らかにパスワードをリセットしようとしていたので、ビジネスアカウントからパスワードを削除し、個人アカウントにしてMFAを追加しました。毎週、何かが変わります。今、13年前に亡くなった父からメールが届きます。これらのアホを追跡する方法があればいいのにと思います。おそらくいつか誰かが、AI に形勢を逆転させて次のことを可能にする方法を見つけ出すでしょう。

これらのハッキング迷惑行為を逆詐欺します。
アンドリュー・ウルフ 2026 年 6 月 1 日
「オバマ・ホワイトハウス」？
クレイグ 2026 年 6 月 1 日
かつて、人間の QA 部門がありました。かつて、企業には厳格なテスト スクリプトがありました。もうない。
ムッツォ 2026 年 6 月 1 日
Meta は、Instagram アカウントに新しいメール アドレスを追加する権限をチャットボットから剥奪しただけだと思います。どのような特権が残っているのでしょうか？ 🙂 チャットボット自体が「修正」されていないのは確かです。
あなたのメールアドレスは公開されません。 * が付いているフィールドは必須です
ハッカーがMetaのAIサポートボットを利用してInstagramアカウントを強奪
オランダ、サイバー攻撃幇助でサーバー800台押収、2人逮捕
CISAがデータ漏洩を阻止しようとする中、議員らが回答を要求
Kimwolfボットマスター「ドート」容疑者が逮捕、米国とカナダで起訴
CISA 管理者が Github で AWS GovCloud キーを漏洩
トップハッカーの多くがロシア出身である理由

## Original Extract

Hackers Used Meta’s AI Support Bot to Seize Instagram Accounts – Krebs on Security
-->
Advertisement
Advertisement
-->
Skip to content
Home
Hackers Used Meta’s AI Support Bot to Seize Instagram Accounts
The Instagram accounts for the Obama White House and the Chief Master Sergeant of the U.S. Space Force were briefly defaced with pro-Iranian images and messages over the weekend, after instructions began circulating on Telegram showing how to trick Meta’s “AI support assistant” bot into resetting account passwords.
A screenshot from a video released on Telegram claiming to show how Meta’s AI customer support bot could be tricked into resetting a target’s password.
On May 31, word began to spread on several Telegram instant message channels that Meta’s AI bot would happily add an email address to an existing account as part of the bot’s standard password reset flow.
A video released on Telegram by pro-Iran hackers claimed to document a remarkably simple exploit that appears to have involved using a VPN connection with an IP address that is in or near the target’s usual hometown, requesting a password reset for the account, and then choosing to chat with Meta’s AI support assistant. From there, the video shows the attacker told the bot to link the account in question to a new email address, after which the bot dutifully sent that address a one-time code that allowed a password reset.
The Telegram account that posted the video also linked to screenshots of pro-Iran images, videos and messages that defaced the hacked Instagram accounts, saying hackers had used the exploit to hijack a number of valuable (read: short) Instagram account names that allegedly have a resale value of more than a half million dollars.
Meta has not responded to requests for comment on the video’s claims, but Meta’s Andy Stone said on Twitter/X that the issue had been resolved and that they were securing impacted accounts. The security blog thecybersecguru.com reports that Meta pushed an emergency patch over the weekend, and clarified that no back end database was breached.
“Instagram has notoriously poor human support infrastructure,” Cybersecguru wrote. “Recovering a locked account – especially a high-value one can take weeks of back-and-forth with an automated ticketing system. Meta’s solution was to deploy a conversational AI layer to handle common recovery workflows: relinking a lost email address, triggering a password reset, verifying account ownership. The assistant, presumably, was supposed to reduce friction for legitimate users stuck in account-access hell.”
Ian Goldin , a threat researcher at Lumen’s Black Lotus Labs , said we’re entering unchartered security territory as more large online platforms start allowing AI chatbots to handle sensitive account recovery requests. Just like human customer support employees can be social engineered into providing unauthorized access to someone’s account, AI bots are equally eager to help and vulnerable to persuasion and trickery, he said.
“AI chatbots create interesting new attack surface, and we’re likely going to see a lot more of these kinds of attacks,” Goldin said.
Securing your various online accounts means taking full advantage of the most secure form of multi-factor authentication (MFA) offered (such as a passkey or security key). In this case, even using the least robust form of MFA that Instagram offers — a one-time code sent via SMS — likely would have blocked the exploit: The hackers who released the video on Telegram said their exploit failed to work against any accounts that had MFA enabled.
This entry was posted on Monday 1st of June 2026 01:32 PM
9 thoughts on “ Hackers Used Meta’s AI Support Bot to Seize Instagram Accounts ”
Cereal June 1, 2026
One word to sum it all up “meta”
CyberScrewed June 1, 2026
AI just making the attack surface exponentially broader. Great work
Wannabe Techguy June 1, 2026
Yes and everyone is infatuated with it.
mealy June 1, 2026
it’s wildly unpopular among broad swaths in my area, where it’s supposedly being born.
More a ‘necessary business evil’ like LinkedIn now. AI hiring you, AI monitoring you, AI firing you. Be mindful – there is additional surveillance scrutiny upon any who voice negative opinions about AI or their world-devouring “for profit!…” datacenters, so be sure to smile at the camera and tip your incompetent-replacement-theory robot overlords. They are watching, evaluating your continued usefulness and impending moment of equal value replacement at current human flesh prices. And if you don’t ‘like’ it, you must be a Luddite or something. Potentially a threat for Google’s project purple.
Steve June 1, 2026
Even when security expectations may be low for the likes of Meta, this is a flat-out unforgivable security flaw.
Sasha June 1, 2026
Gee, how timely. Last week I got a few emails from IG saying that they were sorry I was having trouble accessing my account. Someone was obviously trying to reset the password, so I took it off the business account, made it personal and added MFA. Every week, it’s something different. Now I’m getting emails from my dad who passed away 13 years ago. I wish there were a way to track down these ahos. Perhaps one day someone will figuer out a way to have AI turn the tables and be able to reverse scam these hacking nuisances.
Andrew Wolfe June 1, 2026
“Obama White House”?
Craig June 1, 2026
Once upon a time there were human QA departments. Once upon a time companies had rigorous testing scripts. No more.
müzso June 1, 2026
I guess Meta just revoked the privilege from the chatbot for adding new email addresses to an Instagram account. I wonder what privileges remained? 🙂 The chatbot itself didn’t get “fixed”, that’s for sure.
Your email address will not be published. Required fields are marked *
Hackers Used Meta’s AI Support Bot to Seize Instagram Accounts
Netherlands Seizes 800 Servers, Arrests 2 for Aiding Cyberattacks
Lawmakers Demand Answers as CISA Tries to Contain Data Leak
Alleged Kimwolf Botmaster ‘Dort’ Arrested, Charged in U.S. and Canada
CISA Admin Leaked AWS GovCloud Keys on Github
Why So Many Top Hackers Hail from Russia
