---
source: "https://emailsdone.dev/#prompt-generator"
hn_url: "https://news.ycombinator.com/item?id=48553070"
title: "Show HN: A prompt generator to help AI agents implement my new email API"
article_title: "Emails. Done. | Transactional email without the faff."
author: "mikeapple"
captured_at: "2026-06-16T10:40:45Z"
capture_tool: "hn-digest"
hn_id: 48553070
score: 1
comments: 0
posted_at: "2026-06-16T10:20:52Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A prompt generator to help AI agents implement my new email API

- HN: [48553070](https://news.ycombinator.com/item?id=48553070)
- Source: [emailsdone.dev](https://emailsdone.dev/#prompt-generator)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T10:20:52Z

## Translation

タイトル: HN を表示: AI エージェントが新しいメール API を実装するのを支援するプロンプト ジェネレーター
記事タイトル: メール。終わり。 |無駄のないトランザクションメール。
説明: 電子メールを気にしない開発者向けのトランザクション電子メール。プロンプトを 1 つコピーし、退屈な部分を配線し、パスワードのリセット、確認、通知、請求メールを機能させます。
HN テキスト: Firebase を使用してアプリを構築するのが大好きですが、基本認証以外にアプリのメール (通知、請求など) が欠けていると感じました。
これに対する解決策を実装したいと考え、コーディング エージェントに新しい API を伝える方法と、それを実装する方法を考えました。
私はスクリプトを介して openapi 仕様を動的に生成し、この仕様を使用してプロンプト ジェネレーターを作成してみました。これまでに遭遇したことはありませんでしたが、その理由はなんとなく理解できます...
API を実装するのに十分な詳細を含むプロンプトを生成する一方で、多くのコード ベースで使用できるように具体的になりすぎないようにするのが、絶妙なバランスです。私のアプローチと実行についてのフィードバックをお待ちしています。

記事本文:
メール。終わり。 |無駄のないトランザクションメール。
メール。終わり。
電子メール プロジェクトを必要としないビルダー向けの、本番環境に対応したアプリの電子メール。
事前に構築されたパスワードのリセット、検証、通知、請求メール
コードを最初に作成するので、テンプレートを構築したり、インフラストラクチャに手間をかけたりする必要はありません。
最初の本番メールを送信する前に必要なものすべて
開発、ステージング、運用環境を個別に作成する
毎日のトラフィックを処理しながら、アプリが広まったときに即座に拡張するパブリック電子メール API を実装します。
電子メール認証を扱う (DKIM、SPF、DMARC、MAIL FROM...電子メールは奇妙なものなので)
DNS レコードを追加し、すべてが検証されるまで待ちます
SES サンドボックスを脱出するために本番アクセスを申請する
AWS に送信するユースケースを説明し、承認されることを期待します
環境ごとにドメインと送信 ID を構成する
すべての SaaS が最終的に必要とする 50 以上の退屈なトランザクション メールを作成する
どこでも表示が異なるレスポンシブ HTML メールに対処する
メールの障害によってアプリが中断されないようにキューを構築する
プロバイダーが必然的に失敗する場合の再試行を処理する
バウンスとスパムの苦情を処理する
購読解除処理および抑制リストを作成する
評判が低下する前に、無効な電子メール アドレスへの送信を停止します。
到達性と送信者の評判を監視する
レート制限と不正行為に対する保護を追加する
Gmail、Outlook、Apple Mail、モバイル全体でテストする
そして、はい...電子メールごとに料金が発生します。
動作する cURL コマンドを作成しています...
50 を超えるトランザクション テンプレートがすでに含まれています。パスワードのリセット、検証、マジック リンク、通知、請求、エクスポートなど。
開始方法を選択し、必要なものをコピーして、最初のアプリメールを送信します。
電子メールを高速に動作させるための軽量の SDK とツール。
さらに多くのフレームワーク ヘルパーとパッケージ マネージャーが間もなくリリースされる予定です。
私たちに連絡して伝えてください

どのような統合を希望されるかお知らせください。

## Original Extract

Transactional email for developers who do not care about email. Copy one prompt, wire the boring parts, and get password resets, verification, notifications and billing emails working.

I love building apps with Firebase but felt that, other than basic auth, app emails (notifications, billing etc) were missing.
Wanting to implement a solution to this, it got me thinking about how to tell a coding agent about a new api and how to implement it.
I dynamically generate an openapi spec via a script and have been playing around with using this spec to then create a prompt generator. It's not something I've come across before and I can sort of understand why...
Generating a prompt that's got enough detail to implement the api but not being too specific so it can be used in many code bases is a fine balance. I'd be interesting in any feedback on my approach and my execution.

Emails. Done. | Transactional email without the faff.
Emails. Done.
Production-ready app email for builders who don’t want an email project.
Prebuilt password resets, verification, notifications and billing emails
Code first, no templates to build and no infrastructure faff.
Everything you need before your first production email
Create separate dev, staging and production email environments
Implement a public email API that handles everyday traffic but scales instantly when your app goes viral
Deal with email authentication (DKIM, SPF, DMARC, MAIL FROM... because email is weird)
Add DNS records and wait for everything to verify
Apply for production access to escape the SES sandbox
Explain your sending use case to AWS and hope they approve it
Configure domains and sending identities for every environment
Build the 50+ boring transactional emails every SaaS eventually ends up needing
Deal with responsive HTML emails that somehow render differently everywhere
Build queueing so email failures do not break your app
Handle retries when providers inevitably fail
Process bounces and spam complaints
Build unsubscribe handling and suppression lists
Stop sending to dead email addresses before your reputation tanks
Monitor deliverability and sender reputation
Add rate limiting and abuse protection
Test across Gmail, Outlook, Apple Mail and mobile
And yes... you still pay per email.
Creating a working cURL command...
50+ transactional templates already included. Password resets, verification, magic links, notifications, billing, exports and more.
Choose how you want to start, copy what you need, and send your first app email.
Lightweight SDKs and tools to get email working fast.
More framework helpers and package-manager releases will follow soon.
Contact us and tell us what integrations you'd like to see.
