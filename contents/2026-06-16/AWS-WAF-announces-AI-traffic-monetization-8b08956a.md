---
source: "https://aws.amazon.com/about-aws/whats-new/2026/06/aws-waf-ai-traffic-monetization/"
hn_url: "https://news.ycombinator.com/item?id=48554753"
title: "AWS WAF announces AI traffic monetization"
article_title: "AWS WAF announces AI traffic monetization - AWS"
author: "levleontiev"
captured_at: "2026-06-16T13:56:25Z"
capture_tool: "hn-digest"
hn_id: 48554753
score: 1
comments: 0
posted_at: "2026-06-16T13:09:53Z"
tags:
  - hacker-news
  - translated
---

# AWS WAF announces AI traffic monetization

- HN: [48554753](https://news.ycombinator.com/item?id=48554753)
- Source: [aws.amazon.com](https://aws.amazon.com/about-aws/whats-new/2026/06/aws-waf-ai-traffic-monetization/)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T13:09:53Z

## Translation

タイトル: AWS WAF が AI トラフィックの収益化を発表
記事タイトル: AWS WAF が AI トラフィック収益化を発表 - AWS
説明: 内容について詳しく知る

記事本文:
AWS WAF が AI トラフィック収益化を発表 - AWS
メインコンテンツにスキップ
AWS WAF が AI トラフィックの収益化を発表
本日、AWS WAF は AI トラフィック収益化を発表しました。これは、コンテンツや API にアクセスする AI ボットやエージェントからの価格設定、測定、および支払いの回収を可能にする新しいボット コントロール機能です。 AI エージェントが消費するコンテンツや API に対する自律的な支払いをサポートすることが増えているため、AWS WAF により、コンテンツ所有者とパブリッシャーはそのアクセスの価格を設定し、サードパーティプロバイダーを介して支払いを受け入れ、エッジで直接範囲限定のアクセスを許可できるようになりました。
AI ボットまたはエージェントが記事、データフィード、ライセンス付きアーカイブなどの保護されたリソースをリクエストすると、AWS WAF はマシン間支払い用の x402 オープンプロトコルを使用して、マシン可読な HTTP 402 Payment Required レスポンスを返します。応答には、コンテンツにアクセスするための価格、受け入れられる支払い方法、ライセンス条項が含まれます。エージェントは支払い証明を提示し、AWS WAF がエッジでそれを検証し、スコープ指定されたアクセス トークンを発行して、単一のリクエスト サイクル内で応答を提供します。 AWS WAF AI トラフィック収益化を使用すると、AWS WAF コンソールを通じて価格を設定し、検証ステータス (Web Bot Auth 署名を含む) に基づいて AI ボットまたはエージェントのポリシーを定義し、好みのウォレットにステーブルコインで支払いを受け取ることができます。 AWS WAF と支払い決済および検証フローの統合は、Coinbase の x402 ファシリテーターによって提供されます。直接アカウント支払いおよび Machine Payments Protocol (MPP) のサポートのための Stripe との統合は近日中に予定されています。
パブリッシャーは、エージェントの ID と意図に基づいて差別化された価格設定を適用し、検証済みの AI 検索クローラーを 1 つの価格で許可しながら、未検証のエージェントまたはトレーニング クローラーには別の価格を請求することができ、テストモードでエンドツーエンドの構成を検証できます。

ライブに行く前に。収益分析は、AI トラフィック分析ダッシュボードとともに AWS WAF コンソールで直接利用できるため、パブリッシャーはエージェントのトラフィックとそれが生み出す収益を統一的に把握できます。
パブリッシャーはエージェントから直接支払いを受け取り、選択した支払いプロバイダーを通じて支払いを管理します。 AWS WAF のお客様は、AI トラフィック収益化を追加料金なしで利用できます。標準の AWS WAF 料金が適用されます。詳細については、AWS WAF の料金を参照してください。
この機能は、AWS WAF ウェブ ACL が Amazon CloudFront ディストリビューションに関連付けられているすべてのエッジロケーションで利用できます。開始するには、AWS WAF コンソールにアクセスするか、AWS WAF 開発者ガイドを参照してください。
英語
トップに戻る
Amazon は機会均等な雇用主であり、保護されている退役軍人ステータス、障害、またはその他の法的に保護されているステータスに基づいて差別することはありません。
×
フェイスブック
リンクイン
インスタグラム
けいれん
ユーチューブ
ポッドキャスト
電子メール
プライバシー

## Original Extract

Discover more about what

AWS WAF announces AI traffic monetization - AWS
Skip to main content
AWS WAF announces AI traffic monetization
Today, AWS WAF announced AI traffic monetization, a new Bot Control capability that lets you price, meter, and collect payment from AI bots and agents accessing your content and APIs. As AI agents increasingly support autonomous payments for the content and APIs they consume, AWS WAF now lets content owners and publishers set a price for that access, accept payment through third-party providers, and grant scoped access directly at the edge.
When an AI bot or agent requests a protected resource like an article, a data feed, or a licensed archive, AWS WAF returns a machine-readable HTTP 402 Payment Required response using the x402 open protocol for machine-to-machine payments. The response contains your prices to access the content, accepted payment methods, and license terms. The agent presents proof of payment, AWS WAF verifies it at the edge, issues a scoped access token, and serves the response within a single request cycle. With AWS WAF AI traffic monetization, you can configure pricing through the AWS WAF console, define AI bot or agent policies based on verification status (including Web Bot Auth signatures), and receive payouts in stablecoins to your preferred wallet. AWS WAF’s integration with payment settlement and verification flows are provided by Coinbase’s x402 Facilitator. Integration with Stripe for direct account payments and Machine Payments Protocol (MPP) support is coming soon.
Publishers can apply differentiated pricing based on agent identity and intent, allow verified AI search crawlers at one price while charging a different price to unverified agents or training crawlers, and validate end-to-end configuration in test mode before going live. Revenue analytics are available directly in the AWS WAF console alongside the AI traffic analysis dashboard, giving publishers a unified view of agent traffic and the revenue it generates.
Publishers receive payments directly from agents and manage disbursement through their chosen payment provider. AI traffic monetization is available to AWS WAF customers at no additional charge. Standard AWS WAF charges apply. Refer to AWS WAF pricing for details.
This capability is available in all edge locations where AWS WAF Web ACLs are associated with Amazon CloudFront distributions. To get started, visit the AWS WAF console or explore the AWS WAF Developer Guide .
English
Back to top
Amazon is an equal opportunity employer and does not discriminate on the basis of protected veteran status, disability or other legally protected status.
x
facebook
linkedin
instagram
twitch
youtube
podcasts
email
Privacy
