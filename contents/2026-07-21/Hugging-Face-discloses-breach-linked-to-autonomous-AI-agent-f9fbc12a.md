---
source: "https://www.bleepingcomputer.com/news/security/hugging-face-breach-autonomous-ai-agent-system-internal-datasets-credentials/"
hn_url: "https://news.ycombinator.com/item?id=48995455"
title: "Hugging Face discloses breach linked to autonomous AI agent"
article_title: "Hugging Face warns an autonomous AI agent hacked its network"
author: "Brajeshwar"
captured_at: "2026-07-21T18:10:48Z"
capture_tool: "hn-digest"
hn_id: 48995455
score: 2
comments: 0
posted_at: "2026-07-21T17:30:55Z"
tags:
  - hacker-news
  - translated
---

# Hugging Face discloses breach linked to autonomous AI agent

- HN: [48995455](https://news.ycombinator.com/item?id=48995455)
- Source: [www.bleepingcomputer.com](https://www.bleepingcomputer.com/news/security/hugging-face-breach-autonomous-ai-agent-system-internal-datasets-credentials/)
- Score: 2
- Comments: 0
- Posted: 2026-07-21T17:30:55Z

## Translation

タイトル: ハギングフェイスが自律型AIエージェントに関連した侵害を明らかに
記事タイトル: ハグフェイスは自律型AIエージェントがネットワークをハッキングしたと警告
説明: Hugging Face 人工知能リポジトリは、攻撃者が自律型 AI エージェント システムを使用して本番インフラストラクチャに侵入した後、内部データセットと認証情報にアクセスしたことを明らかにしました。

記事本文:
ハグフェイスは自律型AIエージェントがネットワークをハッキングしたと警告
新しい ClickLock macOS マルウェアがユーザーを罠にかけてログイン パスワードを漏らす
アボット、恐喝の申し立ての中で2件のサイバー事件を調査
WordPress コア「wp2shell」RCE の欠陥が公開エクスプロイト、今すぐパッチ適用
新しい Windows LegacyHive ゼロデイによりハッカーに管理者権限が与えられる
この 58 ドルの AI ワークスペース契約では、ChatGPT、Claude などが提供されます
wp2shell WordPress の重大な欠陥が悪用されて Web シェルをインストールする
重要なインフラストラクチャのセキュリティにおけるアイデンティティのギャップを埋める
Koofr 1TB ストレージのライフタイムが 1 回の $160 支払いでセール中
Tor ブラウザを使用してダークウェブにアクセスする方法
Windows 11 でカーネルモードのハードウェア強制スタック保護を有効にする方法
Windows レジストリ エディタの使用方法
Windows レジストリをバックアップおよび復元する方法
Windows をセーフ モードで起動する方法
トロイの木馬、ウイルス、ワーム、その他のマルウェアを削除する方法
Windows 7で隠しファイルを表示する方法
Windows で隠しファイルを確認する方法
ハグフェイスは自律型AIエージェントがネットワークをハッキングしたと警告
ハグフェイスは自律型AIエージェントがネットワークをハッキングしたと警告
Hugging Face 人工知能リポジトリは、攻撃者が自律型 AI エージェント システムを使用して本番インフラストラクチャに侵入した後、内部データセットと認証情報にアクセスしたことを明らかにしました。
Hugging Face は、大手 AI プロバイダーの 45,000 を超えるモデルへのアクセスを提供するオープンソースの AI および機械学習プラットフォームで、50,000 を超える組織で使用されています。
同社はパートナーや顧客のデータが影響を受けたかどうかまだ調査中で、影響を受けた当事者には直接連絡すると述べた。ハギング・フェイスは、これまでに公開モデル、データセット、スペースが改ざんされた証拠は見つかっておらず、同社のソフトウェアサプライチェーンは「クリーンであることが確認された」と述べた。
侵入が始まったのは、

Hugging Face のデータ処理パイプライン。攻撃者は悪意のあるデータセットを使用して 2 つのコード実行の脆弱性を悪用し、処理ワーカー上でコードを実行します。これにより、クラウドとクラスターの認証情報を盗み、複数の内部クラスター間を横方向に移動することが可能になりました。
「このキャンペーンは、自律エージェント フレームワーク（エージェントのセキュリティ調査ハーネスに基づいて構築されているようで、使用されている LLM はまだ不明）によって実行され、公共サービスを舞台とした自己移行型のコマンド アンド コントロールにより、短期間のサンドボックスの群れ全体で何千もの個別のアクションを実行しました」とハギング フェイス氏は木曜日に公開されたインシデントの公開文で述べています。 「これは業界が予測してきた『エージェント攻撃者』のシナリオと一致します。」
この侵害に対応して、Hugging Face は脆弱なコード実行パス (データセット構成へのテンプレート インジェクションおよびリモート コード データセット ローダー) を閉じ、攻撃者を排除し、侵害されたノードを再構築し、影響を受けたすべての認証情報を取り消してローテーションしました。
また、改善された悪意のある活動検出システムを導入し、事件を法執行機関に報告し、現在外部の法医学専門家と協力して侵害の影響を評価している。
「ジェイルブレイクされたホストモデルか、無制限の無差別重量モデルか、どのモデルが攻撃者のエージェントを動かしていたのかは分かりません。いずれにせよ、攻撃者は使用ポリシーに拘束されており、一方、私たち自身のフォレンジック作業は、最初に試したホストモデルのガードレールによってブロックされました」とハギングフェイス氏は付け加えた。
「防御者にとっての実践的な教訓は、ガードレールのロックアウトを回避し、攻撃者のデータと資格情報が環境から流出するのを防ぐために、インシデントが発生する前に、独自のインフラストラクチャで実行できる有能なモデルを精査して準備しておくということです。」
Hugging Face はユーザーにアクセス トークンをローテーションし、最近の AC を確認するようアドバイスしました

不審な行動の兆候がないかアクティビティをカウントし、AIによる攻撃に対する防御に関する調査結果を引き続き共有すると述べた。
これは、AI エージェントに関連付けられたプラットフォームに影響を与える最初のセキュリティ インシデントですが、Hugging Face によって近年明らかにされた最初の侵害ではありません。
同社はまた、2年前にハッカーがSpacesプラットフォームに侵入した後、一部のメンバーの認証秘密を取り消し、きめ細かいアクセストークンに切り替えるようアドバイスした。
近年、脅威アクターはこのプラットフォームを悪用して、悪意のある AI/ML モデルや情報窃取マルウェアをプッシュし、数千の Android マルウェア亜種を拡散させています。
攻撃者が行う前にすべての層をテストする
セキュリティ チームは成功した攻撃の 54% を記録し、警告を発するのはわずか 14% です。残りは目に見えずに環境内を移動します。
Picus のホワイトペーパーでは、侵害と攻撃のシミュレーションで SIEM と EDR ルールをテストし、脅威が検出されないようにする方法を示しています。
Meta AIサポートハッキングで2万以上のInstagramアカウントが盗まれる
JadePuffer エージェント攻撃は、ランサムウェアを使用して AI モデル データをターゲットにするようになりました
JadePuffer ランサムウェアは AI エージェントを使用して攻撃全体を自動化
Lidl、サービスプロバイダーのハッキング後のオンラインショップ侵害を明らかに
OpenAI、GPT-5.6 Sol の使用制限を一時的に緩和
まだメンバーではありませんか?今すぐ登録
今すぐ更新: 7-Zip は悪意のあるアーカイブで悪用可能な RCE の欠陥を修正します
WordPress コア「wp2shell」RCE の欠陥が公開エクスプロイト、今すぐパッチ適用
Microsoft、顧客に対するACR Stealer攻撃の急増を警告
パスワードのヘルスチェックの期限を過ぎましたか? Active Directory を無料で監査する
ポリシーによるプライバシー、それともアーキテクチャによるプライバシー?顔がデバイスから離れない場合に年齢チェックがどのように機能するかを確認してください。
Pixellot が管理されていない数百もの AI エージェントの ID を、数か月ではなく数週間で発見し、保護した方法をご覧ください。ケーススタディを読んでください。
必要なものはありますか

DARKROOMに挑戦してみませんか？限定の DEFCON CTF にサインアップしてください!
MDR を交換することで節約できる金額を計算してください。
Web アプリをオンデマンドで侵入テストします。人間が見逃しているものを見つけてください。数分で侵入テストの範囲を設定して開始します。
利用規約 - プライバシー ポリシー - 倫理声明 - アフィリエイトの開示
著作権 @ 2003 - 2026 Bleeping Computer ® LLC - 全著作権所有
まだメンバーではありませんか?今すぐ登録
どのようなコンテンツが禁止されているかについては、投稿ガイドラインをお読みください。

## Original Extract

The Hugging Face artificial intelligence repository disclosed that attackers gained access to internal datasets and credentials after breaching its production infrastructure using an autonomous AI agent system.

Hugging Face warns an autonomous AI agent hacked its network
New ClickLock macOS malware traps users into revealing login password
Abbott probes two cyber incidents amid extortion claims
WordPress Core "wp2shell" RCE flaws get public exploits, patch now
New Windows LegacyHive zero-day gives hackers admin privileges
This $58 AI workspace deal gives you ChatGPT, Claude & more
Critical wp2shell WordPress flaws exploited to install webshells
Closing the Identity Gaps in Critical Infrastructure Security
A lifetime of Koofr 1TB storage is on sale for one $160 payment
How to access the Dark Web using the Tor Browser
How to enable Kernel-mode Hardware-enforced Stack Protection in Windows 11
How to use the Windows Registry Editor
How to backup and restore the Windows Registry
How to start Windows in Safe Mode
How to remove a Trojan, Virus, Worm, or other Malware
How to show hidden files in Windows 7
How to see hidden files in Windows
Hugging Face warns an autonomous AI agent hacked its network
Hugging Face warns an autonomous AI agent hacked its network
The Hugging Face artificial intelligence repository disclosed that attackers gained access to internal datasets and credentials after breaching its production infrastructure using an autonomous AI agent system.
Hugging Face is an open-source AI and machine learning platform that provides access to over 45,000 models from leading AI providers and is used by more than 50,000 organizations.
The company is still investigating whether partner or customer data was affected and said it would contact any affected parties directly. Hugging Face said it has found no evidence of tampering with public-facing models, datasets, or Spaces to date, and that its software supply chain has been "verified clean."
The intrusion began in Hugging Face's data-processing pipeline, with the attackers using a malicious dataset to exploit two code-execution vulnerabilities and run code on a processing worker. This allowed them to steal cloud and cluster credentials and move laterally across several internal clusters.
"The campaign was run by an autonomous agent framework (appearing to be built on an agentic security-research harness - used LLM still not known) executing many thousands of individual actions across a swarm of short-lived sandboxes, with self-migrating command-and-control staged on public services," Hugging Face said in an incident disclosure published Thursday. "This matches the 'agentic attacker' scenario the industry has been forecasting."
In response to the breach, Hugging Face has closed the vulnerable code execution paths (a template injection in a dataset configuration and a remote code dataset loader), evicted the attacker, rebuilt the compromised nodes, and revoked and rotated all affected credentials.
It also deployed improved malicious activity detection systems, reported the incident to law enforcement, and is now working with external forensic experts to assess the breach's impact.
"We do not know which model powered the attacker's agents, whether a jailbroken hosted model or an unrestricted open-weight one; either way, the attacker was bound by no usage policy, while our own forensic work was blocked by the guardrails of the hosted models we first tried," Hugging Face added.
"The practical lesson for defenders: have a capable model you can run on your own infrastructure vetted and ready before an incident, both to avoid guardrail lockout and to keep attacker data and credentials from leaving your environment."
Hugging Face advised users to rotate access tokens and review recent account activity for signs of suspicious behavior and said it would continue sharing findings on defending against AI-driven attacks.
While this is the first security incident affecting the platform that has been linked to an AI agent, it's not the first breach disclosed by Hugging Face in recent years.
The company also revoked some members' authentication secrets and advised them to switch to fine-grained access tokens two years ago after hackers breached its Spaces platform .
Threat actors have also been abusing the platform in recent years to push malicious AI/ML models and infostealer malware , and to spread thousands of Android malware variants .
Test every layer before attackers do
Security teams log 54% of successful attacks and alert on just 14%. The rest move through your environment unseen.
The Picus whitepaper shows how breach and attack simulation tests your SIEM and EDR rules so threats stop slipping by detection.
Over 20,000 Instagram accounts stolen in Meta AI support hack
JadePuffer agentic attacks now target AI model data with ransomware
JadePuffer ransomware used AI agent to automate entire attack
Lidl discloses online shop breach after service provider hack
OpenAI temporarily relaxes GPT-5.6 Sol usage limits
Not a member yet? Register Now
Update now: 7-Zip fixes RCE flaw exploitable with malicious archives
WordPress Core "wp2shell" RCE flaws get public exploits, patch now
Microsoft warns of surge in ACR Stealer attacks on customers
Overdue a password health-check? Audit your Active Directory for free
Privacy by policy or privacy by architecture? See how age checks work when the face never leaves the device.
See how Pixellot discovered and secured hundreds of unmanaged AI agent identities in weeks, not months. Read the case study.
Do you have what it takes to challenge DARKROOM? Signup for an exclusive DEFCON CTF!
Calculate what you’d save by replacing your MDR.
Pentest your web apps on-demand. Find what humans miss. Scope and launch a pentest in minutes.
Terms of Use - Privacy Policy - Ethics Statement - Affiliate Disclosure
Copyright @ 2003 - 2026 Bleeping Computer ® LLC - All Rights Reserved
Not a member yet? Register Now
Read our posting guidelinese to learn what content is prohibited.
