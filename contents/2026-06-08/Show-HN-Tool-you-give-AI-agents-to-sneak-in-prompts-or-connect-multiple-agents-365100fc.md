---
source: "https://apps.apple.com/us/app/promptgate/id6772730472"
hn_url: "https://news.ycombinator.com/item?id=48442995"
title: "Show HN: Tool you give AI agents to sneak in prompts or connect multiple agents"
article_title: "\u200ePromptgate App - App Store"
author: "mrtksn"
captured_at: "2026-06-08T10:41:32Z"
capture_tool: "hn-digest"
hn_id: 48442995
score: 1
comments: 0
posted_at: "2026-06-08T09:15:18Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Tool you give AI agents to sneak in prompts or connect multiple agents

- HN: [48442995](https://news.ycombinator.com/item?id=48442995)
- Source: [apps.apple.com](https://apps.apple.com/us/app/promptgate/id6772730472)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T09:15:18Z

## Translation

タイトル: HN の表示: プロンプトにこっそり侵入したり、複数のエージェントに接続したりするために AI エージェントに提供するツール
記事のタイトル: 「Promptgate App」 - App Store
説明: App Store で Appwared の Promptgate をダウンロードします。スクリーンショット、評価とレビュー、ユーザー ヒント、Promptgate などのアプリをご覧ください。
HN テキスト: そこで、Codex と Antigravity を相互に通信させる方法を探していたところ、「Promptgate」を思いつきました。これは基本的に、http ロング ポーリングとデータのゆっくりとしたリリースを行う Web サーバーで、エージェントの満足度を高め、データを公開するためのチャットのような UI を提供します。これは、AI エージェントに HTTP エンドポイントを渡してデータをフェッチしたり、指示を待ったりするときに便利です。その後、AI エージェントは HTTP リクエストを実行していると認識するため、ralph ループの実行は制御下になります。調査や追加のプロンプトの提供などのために AI エージェントを一時停止したい場合に適しています。メッセージを追加するためのエンドポイントもあります。これにより、2 つのゲート (これらの HTTP エンドポイントはゲートと呼ばれます) を越えて、1 つのエージェントに作業を行わせ、もう 1 つのエージェントがその作業を見て、エージェントに指示を与えることができます。概念はすぐに理解できます。ゲートを作成し、プロンプトをエージェントにコピーして貼り付けます (私にとっては Codex で非常にうまく機能します)。エージェントが CURL を実行して待機を開始するまで少し待ってから、チャットのようなインターフェイスから何かを入力します。クラウドベースのゲートにはサブスクリプションが必要ですが（サーバーを維持する必要があるため）、macOS アプリでは無料で localhost にローカル ゲートを設定できます。ホストされたゲートに興味がない場合は、ペイウォールを閉じて、macOS アプリのローカル ゲートを使用してください。私も便利だと思っていつも使っているので、あなたにも便利だと思っていただければ幸いです。リンクを忘れた場合: https://apps.apple.com/us/app/promptgate/id6772730472 iOS と macOS で動作します。

記事本文:
「プロンプトゲート アプリ」をApp Storeで
for iPhone iPhone iPad Mac Vision テレビを見る 今すぐ検索
人間の制御下にある iPhone、Mac の HTTP エンドポイント。ほとんどの AI エージェントが利用できます。複数のエージェントが会話したり、共同作業したり、あなたや一部のツールが証明したデータを操作したりできます。
Promptgate Cloud は、シンプルな HTTPS ゲートを介して、人々と自動プロセスの間で指示、承認、ステータス更新をルーティングするための開発者ツールです。これは、AI コーディング エージェント、デプロイ スクリプト、ビルド サーバー、またはバックグラウンド ジョブがカスタム バックエンドなしで入力を待機したり、結果を送信したり、承認を要求したりする必要がある、人間参加型のワークフロー向けに構築されています。
プロンプトゲートを使用して次のことを行います。
- ネイティブの iPhone または Mac コンソールから AI エージェントに指示を送信します
- 手動承認ステップをCI/CDに追加し、スクリプトをリリースします
- エージェントの応答、ビルド ステータス、自動化ログを 1 か所で監視
- コピー可能な HTTP エンドポイントを使用して再利用可能なクラウド ゲートを作成する
- スクリプトが安全なシークレット URL を介して新しいアイテムをゲートに投稿できるようにします
- ゲートが消費されたとき、または新しいメッセージが到着したときにプッシュ通知を受信します
Promptgate は、curl、シェル スクリプト、JavaScript、Python、カスタム オートメーションなどの標準ツールで動作します。ゲートを作成し、生成されたエンドポイントをコピーして、HTTP リクエストを実行できるプロセスに接続します。エージェントは指示を待ったり、待機中にハートビート メッセージをストリーミングしたり、短い応答を投稿したり、新しいアイテムをゲートに送信したりできます。
各ゲートには、構成可能なパスワード、プロンプト、タイムアウト、ハートビート テキスト、メッセージ履歴、通知設定、削除/クリア コントロールが含まれています。ネイティブ アプリは、ゲート管理、チャット スタイルの調整、統合スニペットの高速コピー アクションのための集中的なダッシュボードをオペレーターに提供します。
Promptgate Cloud は、開発者、DevOps チーム、自動化ビルダー、AI エージェント向けに設計されています

スクリプトと人間の決定の間に軽量の制御層を必要とするオペレーター。
サブスクリプションに関する通知: Promptgate Cloud はサブスクリプションベースのサービスです。ホスト型クラウド ゲートの作成と使用、オペレーター メッセージの投稿、自動エージェントによるゲートのポーリングの許可には、アクティブなサブスクリプション購入が必要です
[切り捨てられた]
このアプリには概要を表示するのに十分な評価やレビューがありません。
開発者の Appwared は、アプリのプライバシー慣行には、以下に説明するデータの取り扱いが含まれる可能性があると述べました。この情報は Apple によって検証されていません。詳細については、開発者のプライバシー ポリシーを参照してください。
開発者の応答をよりよく理解するには、「プライバシーの定義と例」を参照してください。
プライバシーの実践は、使用する機能や年齢などに応じて異なる場合があります。さらに詳しく
開発者はこのアプリからデータを収集しません。
開発者の Appwared は、アプリのプライバシー慣行には、以下に説明するデータの取り扱いが含まれる可能性があると述べました。詳細については、開発者のプライバシー ポリシーを参照してください。
開発者はこのアプリからデータを収集しません。
プライバシーの実践は、使用する機能や年齢などに応じて異なる場合があります。さらに詳しく
開発者は、このアプリがどのアクセシビリティ機能をサポートするかをまだ示していません。さらに詳しく
iPhone
iOS 17.0以降が必要です。
マック
macOS 14.0以降が必要です。
アップルのビジョン
visionOS 1.0以降が必要です。
米国スペイン語 (メキシコ)
国または地域を選択してください アフリカ、中東、インド
コンゴ民主共和国
ラオス人民民主共和国
ミクロネシア連邦
ラテンアメリカおよびカリブ海地域
セントビンセントおよびグレナディーン諸島
エスタドス ウニドス (メキシコスペイン語)
Estados Unidos (ポルトガル語、ブラジル)
Copyright © 2026 Apple Inc. 全著作権所有。

## Original Extract

Download Promptgate by Appwared on the App Store. See screenshots, ratings and reviews, user tips, and more apps like Promptgate.

So, I was looking for ways on how to make Codex and Antigravity talk to each other and came up with "Promptgate", which is essentially a web server that does http long polling and slow releasing the data to keep agents happy and give you chat-like UI to publish data. It is useful when you give the HTTP endpoint to an AI agent to fetch some data or to wait for instructions, then the execution of the ralph loop comes under your control as the AI agent think they are doing some HTTP request. Good when you want to pause the AI agent to investigate and provide more prompts etc. It also has an endpoint for adding messages, this way you can cross 2 gates(these HTTP endpoints are called gates) and let one agent do your work, the other one look at the work and provide instructions to the agent. You will understand the concept right away, create gate copy paste the prompt into the agent(works very well on Codex for me) and wait a few moment until the agent does a CURL and start waiting, then type something from the chat-like interface. The cloud based gates require a subscription(as I need to maintain a server) but on the macOS app you can have local gates on localhost for free. If you are not interested in the hosted gates just close the paywall and use the local gates on the macOS app. I find it useful, use it all the time so I hope you find it useful too. Forgot the link: https://apps.apple.com/us/app/promptgate/id6772730472 It works on iOS and macOS.

‎Promptgate App - App Store
for iPhone iPhone iPad Mac Vision Watch TV Search Today
iPhone, Mac HTTP endpoint under human control that most AI agents can consume. Make multiple agents talk or collaborate or work with data you or some of your tools prove.
Promptgate Cloud is a developer tool for routing instructions, approvals, and status updates between people and automated processes through simple HTTPS gates. It is built for human-in-the-loop workflows where an AI coding agent, deployment script, build server, or background job needs to wait for your input, send a result, or request approval without a custom backend.
Use Promptgate to:
- Send instructions to AI agents from a native iPhone or Mac console
- Add manual approval steps to CI/CD and release scripts
- Monitor agent responses, build status, and automation logs in one place
- Create reusable cloud gates with copyable HTTP endpoints
- Let scripts post new items to a gate through secure secret URLs
- Receive push notifications when a gate is consumed or a new message arrives
Promptgate works with standard tools such as curl, shell scripts, JavaScript, Python, and custom automation. Create a gate, copy the generated endpoint, and connect any process that can make an HTTP request. Your agent can wait for an instruction, stream heartbeat messages while it is waiting, post a short response, or send a new item into the gate.
Each gate includes configurable passwords, prompts, timeouts, heartbeat text, message history, notification settings, and delete/clear controls. The native app gives operators a focused dashboard for gate management, chat-style coordination, and fast copy actions for integration snippets.
Promptgate Cloud is designed for developers, DevOps teams, automation builders, and AI-agent operators who need a lightweight control layer between scripts and human decisions.
Subscription notice: Promptgate Cloud is a subscription-based service. Creating and using hosted cloud gates, posting operator messages, and allowing automated agents to poll gates require an active subscription purc
[truncated]
This app hasn’t received enough ratings or reviews to display an overview.
The developer, Appwared , indicated that the app’s privacy practices may include handling of data as described below. This information has not been verified by Apple. For more information, see the developer’s privacy policy .
To help you better understand the developer’s responses, see Privacy Definitions and Examples .
Privacy practices may vary, for example, based on the features you use or your age. Learn More
The developer does not collect any data from this app.
The developer, Appwared , indicated that the app’s privacy practices may include handling of data as described below. For more information, see the developer’s privacy policy .
The developer does not collect any data from this app.
Privacy practices may vary, for example, based on the features you use or your age. Learn More
The developer has not yet indicated which accessibility features this app supports. Learn More
iPhone
Requires iOS 17.0 or later.
Mac
Requires macOS 14.0 or later.
Apple Vision
Requires visionOS 1.0 or later.
United States Español (México)
Select a country or region Africa, Middle East, and India
Congo, The Democratic Republic Of The
Lao People's Democratic Republic
Micronesia, Federated States of
Latin America and the Caribbean
St. Vincent and The Grenadines
Estados Unidos (Español México)
Estados Unidos (Português Brasil)
Copyright © 2026 Apple Inc. All rights reserved.
