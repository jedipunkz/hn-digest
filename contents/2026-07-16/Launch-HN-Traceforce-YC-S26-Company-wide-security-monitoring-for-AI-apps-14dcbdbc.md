---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48937020"
title: "Launch HN: Traceforce (YC S26) – Company-wide security monitoring for AI apps"
article_title: ""
author: "XiaHua"
captured_at: "2026-07-16T17:04:14Z"
capture_tool: "hn-digest"
hn_id: 48937020
score: 2
comments: 1
posted_at: "2026-07-16T16:52:16Z"
tags:
  - hacker-news
  - translated
---

# Launch HN: Traceforce (YC S26) – Company-wide security monitoring for AI apps

- HN: [48937020](https://news.ycombinator.com/item?id=48937020)
- Score: 2
- Comments: 1
- Posted: 2026-07-16T16:52:16Z

## Translation

タイトル: HN: Traceforce (YC S26) を開始 – AI アプリの全社的なセキュリティ監視
HN テキスト: HN さん、私たちは Traceforce ( https://www.traceforce.ai/ ) の創設者である Xia と Varun です。 Traceforce は、どのアプリが使用されているかだけでなく、それらが MCP を介して他のデータ ソースにどのように接続されているかを検出することにより、ChatGPT、Claude などの AI アプリをすべてのデバイス (ラップトップ、サンドボックス、仮想マシン) 上で直接可視化し、制御できるようにします。また、脆弱な MCP を検出するためのオープンソースの動的 MCP 侵入テスト ツール https://github.com/traceforce/mcp-xray も用意しています。 Traceforce の目的は次のとおりです。 - 企業の従業員に、デバイス上で実行されている AI ソフトウェアが安全に動作していることを確認するための標準化された方法を提供します。 - 企業のセキュリティ チームに、企業のデバイス上の AI ソフトウェアのアクティビティを可視化し、危険な行為やセキュリティ侵害をできるだけ早期に検出して防止します。 Traceforce の仕組み 1. Traceforce は、軽量のバイナリおよびブラウザ拡張機能として各デバイスにインストールされます。 2. 30 分以内に、デバイスはライブ データを会社プロファイルにアップロードし、会社のすべてのデバイスで実行されているすべての AI エージェント/アプリをダッシュ​​ボードに表示します。 3. 企業のセキュリティ スタッフは、すべてのエージェントのアクティビティをリアルタイムで監視し、制御を実装し、セキュリティ リスクが発生するとすぐに警告を受けることができます。ビデオデモは次のとおりです: https://youtu.be/IdK2WKg7kaM Traceforce のインスピレーションは、Clumio (2024 年 10 月に Commvault に買収) というスタートアップでエンジニアリング ディレクターとして務めた Xia の経験から生まれました。 Clumio では、チーム メンバーが AI をどのように使用しているかを速度を低下させることなく監視できることが最優先事項でした。 50 人以上の CISO および CIO と話をした結果、これが業界全体で現在非常に必要とされているソリューションであることが明らかになりました。私たちは

新しい AI 機能が非常に急速かつ広範囲に導入されているため、可視性と制御が追いつかないという話をよく聞きます。 Traceforce は、監視および収集する内容について透明性があります。デフォルトでは、Traceforce は、デバイス上で実行されている AI アプリケーション、MCP、およびツールに関するメタデータとテレメトリのみを収集します。セキュリティ チームは、事前定義された高リスクまたは破壊的な可能性のあるアクションを検出、警告、またはブロックする目的で、ツール呼び出しを検査するオプションを有効にすることができます。すべてのコンテンツ検査はデバイス上でローカルに行われます。組織のセキュリティ管理者が明示的に構成しない限り、ユーザー プロンプトは保存されません。当社は製品のエンドユーザーと緊密に連携しており、何が監視/共有されているかを理解すると、セキュリティインシデントを防ぐための強力な保護層がデバイス上にあることに実際に大きな安心感を抱くようになります。これにより、気付かないうちに内部でどのような漏洩や侵害が起こっているかを心配することなく、自分の仕事に集中することができます。 Traceforce は現在、10 組織の 1,000 台以上のデバイスに導入されています。平均して、デバイスごとに 15 を超える AI アプリケーションが検出され、各アプリケーションは 5 ～ 10 個の MCP に接続されています。当社は、顧客が MCP 構成で公開された平文秘密を特定し、AI 生成コードを介して API キーが漏洩するのを防ぎ、「DROP TABLE」などの破壊的な可能性のあるコマンドを実行する前に開発者に警告できるよう支援してきました。私たちの「警告と承認」アプローチは特に好評で、開発者は自由に作業できると同時に、コストのかかるミスを回避することができます。私たちは、AI コーディング アシスタント、ChatGPT、Claude、MCP を急速に導入している中小企業 (従業員 200 名以上) のセキュリティ、IT、AI プラットフォーム チームと協力したいと考えています。あなたが苦労しているなら、

人々が生産性を向上させるためにどのような AI ツールを使用しているかを理解している場合、または人々の作業を遅らせることなく AI 関連のセキュリティ リスクを軽減する実用的な方法が必要な場合、私たちは喜んで話したいと考えています。 https://www.traceforce.ai で無料トライアルを開始するか、直接連絡してデモをスケジュールし、環境について話し合うことができます。

## Original Extract

Hey HN, we’re Xia and Varun, the founders of Traceforce ( https://www.traceforce.ai/ ). Traceforce provides visibility and control over AI apps such as ChatGPT, Claude etc directly on all devices (laptops, sandboxes, virtual machines) by discovering not just which apps are being used but also how they are connected to other data sources via MCPs. We also have an open-source dynamic MCP pentesting tool https://github.com/traceforce/mcp-xray to detect vulnerable MCPs. The purpose of Traceforce is to: - Give a company’s employees a standardized way to ensure that AI software running on their device is operating safely - Give the company’s security team visibility of the activities of AI software on the company’s devices, and to detect and prevent unsafe actions and security breaches as early as possible. How Traceforce works 1. Traceforce is installed on each device as a lightweight binary and browser extension. 2. Within 30 minutes, the device is uploading live data to the company profile, displaying all the AI agents/apps running across all company devices on a dashboard. 3. Company security staff can monitor the activity of all the agents in real time, implement controls, and be alerted to any security risks as soon as they arise. Here’s the video demo: https://youtu.be/IdK2WKg7kaM The inspiration for Traceforce came via Xia’s experience as Director of Engineering at a startup called Clumio (which was acquired by Commvault in Oct 2024). Being able to monitor how team members are using AI without slowing them down was a top priority at Clumio. After speaking with 50+ CISOs and CIOs, it became clear that this is a much-needed solution right now across industries. We keep hearing that new AI features are being adopted so quickly and so broadly that visibility and control just can't keep up. Traceforce is transparent about what we monitor and collect. By default, Traceforce collects only metadata and telemetry about the AI applications, MCPs, and tools running on a device. Security teams can enable options to inspect tool calls for the purpose of detecting, warning on, or blocking predefined high-risk or potentially destructive actions. All content inspection happens locally on the device. User prompts are never stored unless explicitly configured by the organization's security administrators. We work closely with end-users of the product, and once they understand what is being monitored/shared, they actually have great comfort that they have a powerful layer of protection on their device to prevent security incidents. It enables them to just focus on their work without worrying about what leaks and breaches may be happening under the hood without their awareness. Traceforce is currently deployed across more than 1,000 devices at 10 organizations. On average, we discover over 15 AI applications per device with each application connected to 5-10 MCPs. We've helped customers identify exposed plaintext secrets in MCP configurations, prevent API keys from leaking through AI-generated code, and warn developers before executing potentially destructive commands such as “DROP TABLE”. Our "warn and acknowledge" approach has been especially well received, giving developers the freedom to work while helping them avoid costly mistakes. We're looking to work with security, IT, and AI platform teams at small to medium enterprises (200+ employees) that are rapidly adopting AI coding assistants, ChatGPT, Claude, and MCPs. If you're struggling to understand what AI tools people use to boost their productivity or need a practical way to reduce AI-related security risk without slowing folks down, we'd love to talk. You can get started with a free trial at https://www.traceforce.ai or reach out directly to schedule a demo and discuss your environment.

