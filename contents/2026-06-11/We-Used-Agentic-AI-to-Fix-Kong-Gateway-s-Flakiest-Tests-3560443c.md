---
source: "https://konghq.com/blog/engineering/how-we-used-agentic-ai-to-fix-kong-gateways-flakiest-tests"
hn_url: "https://news.ycombinator.com/item?id=48485434"
title: "We Used Agentic AI to Fix Kong Gateway's Flakiest Tests"
article_title: "How We Used Agentic AI to Automate Fixes for Flaky Tests | Kong Inc."
author: "dndx"
captured_at: "2026-06-11T04:35:50Z"
capture_tool: "hn-digest"
hn_id: 48485434
score: 6
comments: 0
posted_at: "2026-06-11T02:07:56Z"
tags:
  - hacker-news
  - translated
---

# We Used Agentic AI to Fix Kong Gateway's Flakiest Tests

- HN: [48485434](https://news.ycombinator.com/item?id=48485434)
- Source: [konghq.com](https://konghq.com/blog/engineering/how-we-used-agentic-ai-to-fix-kong-gateways-flakiest-tests)
- Score: 6
- Comments: 0
- Posted: 2026-06-11T02:07:56Z

## Translation

タイトル: Agentic AI を使用して Kong Gateway の最も不安定なテストを修正しました
記事のタイトル: Agentic AI を使用して不安定なテストの修正を自動化する方法 |株式会社コング
説明: Kong エンジニアリング チームがエージェント AI を使用してコードベース内の不安定なテストを特定、修正、検証し、CI の安定性と開発者の生産性を向上させた方法を学びます。

記事本文:
Agentic AI を使用して不安定なテストの修正を自動化する方法 | Kong Inc. API + AI SUMMIT 2026 をお見逃しなく | 8 月 16 日に値上げ KONG の新しいエージェント ゲートウェイで A2A トラフィックを管理する KONG エージェント時代のワールド ツアーのスポットを確保 [ Why Kong ] ( /company/why-kong ) Why Kong
+ 製品とエージェント 製品とエージェント _ API と AI 接続テクノロジー _ 統合 API と AI プラットフォーム [ ] API 管理 AI 管理 イベント管理 収益化 移行サービス API アドバイザリー サービス + 前方展開エンジニア 新しいランタイム
[ API ゲートウェイ ] ( /products/kong-gateway ) API ゲートウェイ
[ AI ゲートウェイ ホット ] ( /products/kong-ai-gateway ) AI ゲートウェイ ホット
[ イベント ゲートウェイ ] ( /products/event-gateway ) イベント ゲートウェイ
[ サービス メッシュ ] ( /products/kong-mesh ) サービス メッシュ
[ コンテキスト メッシュ ] ( /products/kong-konnect/features/context-mesh ) コンテキスト メッシュ
[ イングレス コントローラー ] ( /products/kong-ingress-controller ) イングレス コントローラー
[ コング オペレーター ] ( /products/kong-operator ) コング オペレーター
[ MCP レジストリ NEW ] ( /products/mcp-registry ) MCP レジストリ NEW
[ API サービス カタログ ] ( /products/kong-konnect/features/api-service-catalog ) API サービス カタログ
[ ランタイム管理 ] ( /products/kong-konnect/features/runtime-management ) ランタイム管理
[ APIOps と自動化 ] ( /products/apiops-automation ) APIOps と自動化
[ 開発者ポータル ] ( /products/kong-konnect/features/developer-portal ) 開発者ポータル
[ 使用量の請求と測定 ] ( /products/kong-konnect/features/usage-based-metering-and-billing ) 使用量の請求と測定
[ 可観測性 ] ( /products/kong-konnect/features/api-observability ) 可観測性
[ KAi エージェント ] ( /products/kong-konnect/features/kai-ai-agent ) KAi エージェント
+ ソリューション プラットフォーム チーム向けのソリューション
[ 開発者プラットフォーム ] ( /solutions/building-

開発者プラットフォーム ) 開発者プラットフォーム
[ Kubernetes とマイクロサービス ] ( /solutions/build-on-kubernetes ) Kubernetes とマイクロサービス
[ 可観測性 ] ( /solutions/observability ) 可観測性
[ サービス メッシュ接続 ] ( /solutions/service-mesh-connectivity ) サービス メッシュ接続
[ Kafka イベント ストリーミング ] ( /solutions/kafka-stream-api-management ) Kafka イベント ストリーミング
[ AI 接続性 ] ( /ai-connectivity ) AI 接続性
[ オープン バンキング ] ( /solutions/open-banking ) オープン バンキング
[ レガシー移行 ] ( /solutions/legacy-api-management-migration ) レガシー移行
[ プラットフォームのコスト削減 ] ( /solutions/api-platform-consolidation ) プラットフォームのコスト削減
[ Kafka コストの最適化 ] ( /solutions/reduce-kafka-cost ) Kafka のコストの最適化
[ API 収益化 ] ( /solutions/api-monetization ) API 収益化
[ AI 収益化 ] ( /solutions/ai-monetization ) AI 収益化
[ AI FinOps ] ( /solutions/ai-cost-governance-finops ) AI FinOps
[ エージェント ゲートウェイ ] ( /agent-gateway ) エージェント ゲートウェイ
[ AI ガバナンス ] ( /solutions/ai-governance ) AI ガバナンス
[ AI セキュリティ ] ( /solutions/ai-security ) AI セキュリティ
[ AI コスト管理 ] ( /solutions/ai-cost-optimization-management ) AI コスト管理
[ エージェント インフラストラクチャ ] ( /solutions/agentic-ai-workflows ) エージェント インフラストラクチャ
[ MCP プロダクション ] ( /solutions/mcp-production-and-consumption ) MCP プロダクション
[ MCP トラフィック ゲートウェイ ] ( /solutions/mcp-governance ) MCP トラフィック ゲートウェイ
[ モバイル アプリ API 開発 ] ( /solutions/mobile-application-api-development ) モバイル アプリ API 開発
[ GenAI アプリ開発 ] ( /solutions/power-openai-applications ) GenAI アプリ開発
[ Istio 用 API ゲートウェイ ] ( /solutions/istio-gateway ) Istio 用 API ゲートウェイ
[ 分散負荷分散 ] ( /solutions/decentralized-load-bala

ncing ) 分散負荷分散
[ 金融サービス ] ( /solutions/financial-services-industry ) 金融サービス
[ ヘルスケア ] ( /solutions/healthcare ) ヘルスケア
[ 高等教育 ] ( /solutions/api-platform-for-education-services ) 高等教育
[ 保険 ] ( /solutions/insurance ) 保険
[ 製造 ] ( /solutions/manufaction ) 製造
[ 小売 ] ( /solutions/retail ) 小売
[ ソフトウェアとテクノロジー ] ( /solutions/software-and-technology ) ソフトウェアとテクノロジー
[ 輸送 ] ( /solutions/transportation-and-logistics ) 輸送
[ 価格 ] ( /pricing ) 価格
+ リソース リソース ドキュメント
[ Kong Konnect ] ( https://developer.konghq.com/konnect/ ) Kong Konnect
[ Kong ゲートウェイ ] ( https://developer.konghq.com/gateway/ ) Kong ゲートウェイ
[ コングメッシュ ] ( https://developer.konghq.com/mesh/ ) コングメッシュ
[ Kong AI ゲートウェイ ] ( https://developer.konghq.com/ai-gateway/ ) Kong AI ゲートウェイ
[ Kong イベント ゲートウェイ ] ( https://developer.konghq.com/event-gateway/ ) Kong イベント ゲートウェイ
[ コング インソムニア ] ( https://developer.konghq.com/insomnia/ ) コング インソムニア
[ プラグイン ハブ ] ( https://developer.konghq.com/plugins/ ) プラグイン ハブ
[ ラーニング センター ] ( /blog/learning-center ) ラーニング センター
[ 電子ブック ] ( /resources/e-book ) 電子ブック
[ レポート ] ( /resources/reports ) レポート
[ デモ ] ( /resources/demos ) デモ
[ お客様の事例 ] ( /customer-stories ) お客様の事例
[ ビデオ ] ( /resources/video ) ビデオ
[ API + AI サミット ] ( /events/conferences/api-ai-summit ) API + AI サミット
[ エージェント時代のワールド ツアー ] ( /agentic-era-world-tour ) エージェント時代のワールド ツアー
[ ウェビナー ] ( /events/webinars ) ウェビナー
[ ユーザー コール ] ( /events/user-calls ) ユーザー コール
[ ワークショップ ] ( /events/workshops ) ワークショップ
[ ミートアップ ] ( /events/meetups ) ミートアップ
[ すべてのイベントを見る ] (

/events ) すべてのイベントを見る
[ 始めましょう ] ( https://developer.konghq.com/ ) 始めましょう
[ コミュニティ ] ( /コミュニティ ) コミュニティ
[ 認定 ] ( /academy/certification ) 認定
[ トレーニング ] ( https://education.konghq.com ) トレーニング
[ 私たちについて ] ( /company/about-us ) 私たちについて
【募集中です！ ] ( /company/careers ) 採用中です!
[ プレスルーム ] ( /company/press-room ) プレスルーム
[ お問い合わせ ] ( /company/contact-us ) お問い合わせ
[ Kong パートナー プログラム ] ( /partners ) Kong パートナー プログラム
[ エンタープライズ サポート ポータル ] ( https://support.konghq.com/s/ ) エンタープライズ サポート ポータル
[ドキュメント] ( https://developer.konghq.com/?_gl=1*tphanb*_gcl_au*MTcxNTQ5NjQ0MC4xNzY5Nzg4MDY0LjIwMTI3NzEwOTEuMTc3MzMxODI2MS4xNz czMzE4MjYw*_ga*NDIwMDU4MTU3LjE3Njk3ODgwNjQ.*_ga_4JK9146J1H*czE3NzQwMjg1MjkkbzE4OSRnMCR0MTc3NDAyODUyOSRqNjAkbDAkaDA ) ドキュメント
[ ログイン ] ( https://cloud.konghq.com/login ) ログイン
[ ブックデモ ] ( /contact-sales ) ブックデモ
[ 始めましょう ] ( /products/kong-konnect/register ) 始めましょう
[ ブログ ] ( /blog ) ブログ [ AI ゲートウェイ ] ( /blog/tag/ai-gateway ) AI ゲートウェイ
[ AI セキュリティ ] ( /blog/tag/ai-security ) AI セキュリティ
[ AIOps ] ( /blog/tag/aiops ) AIOps
[ API セキュリティ ] ( /blog/tag/api-security ) API セキュリティ
[ API ゲートウェイ ] ( /blog/tag/api-gateway ) API ゲートウェイ
+ トピックの探索 トピックの探索 [ API 管理 ] ( /blog/tag/api-management ) API 管理
[ API 開発 ] ( /blog/tag/api-development ) API 開発
[ API 設計 ] ( /blog/tag/api-design ) API 設計
[ オートメーション ] ( /blog/tag/automation ) オートメーション
[ サービス メッシュ ] ( /blog/tag/service-mesh ) サービス メッシュ
[ 不眠症 ] ( /blog/tag/insomnia ) 不眠症
[ イベントゲートウェイ ] ( /blog/tag/event-gateway ) イベントゲートウェイ
[ すべてのブログを表示 ] ( /blog/page/1 ) すべてのブログを表示
[ Kong AI ゲートウェイ ] ( /products/kong-a

i-ゲートウェイ) Kong AI ゲートウェイ
[ Kong API ゲートウェイ ] ( /products/kong-gateway ) Kong API ゲートウェイ
[ Kong イベント ゲートウェイ ] ( /products/event-gateway ) Kong イベント ゲートウェイ
[ Kong のメータリングと請求 ] ( /products/kong-konnect/features/usage-based-metering-and-billing ) Kong のメータリングと請求
[ コング インソムニア ] ( /products/kong-insomnia ) コング インソムニア
[ Kong Konnect ] ( /products/kong-konnect ) Kong Konnect
[ ドキュメント ] ( https://developer.konghq.com ) ドキュメント
[ ブックデモ ] ( /contact-sales ) ブックデモ
Agentic AI を使用して Kong Gateway の最も不安定なテストを修正した方法
# Agentic AI を使用して Kong Gateway の最も不安定なテストを修正した方法
大同孫
Kong Gateway のコードベースに変更を加えるたびに、当社がサポートする 2 つの主要なアーキテクチャ (x86 および ARM) 間で 17,000 * 2 = 34,000 を超えるテスト ケースを実行する包括的なテスト スイートがトリガーされます。このプロセスには 1 台のマシンで約 23.5 時間かかります。しかし、私たちはそれほど長く待ちません。大規模なマシンがスイートを並行して実行し、作業を積極的にシャーディングして、各コミットがほんの少しの時間で完了するようにします。この設定は、不安定なテストが関与するまではうまく機能します。
Kong エンジニアリング チームでは、エージェント AI ワークフローによって Kong ゲートウェイの開発をどのように効率化できるかを模索してきました。不安定なテストを修正することが、始めるのに理想的な場所であることがわかりました。
品質は私たちが作るすべての製品の中核です。品質を損なうことなく迅速に対応するために、変更がリグレッションを引き起こしたかどうかを迅速に知らせる広範なテストスイートのセットに依存しています。これらのスイートは非常に大きいため、1 台のマシンで 1 回の完全な実行に約 23.5 時間かかります。そのため、作業を大規模なマシン群に分散させました。
しかし、より多くの機能をリリースするにつれて、スイートは成長し続け、より不安定なテストが発生し始めます。
初心者向け、不安定な手向け

sts は時々失敗するテストですが、通常は再実行で成功します。大規模なプロジェクトでは利用できず、エンジニアは常に消耗します。不安定な障害が発生した後は、ジョブを再実行する必要があり、エンジニアリング時間とマシン時間の両方がかかります。さらに悪いことに、不安定なレートが上昇すると、PR をマージしようとするエンジニアが再実行中にお互いを踏みつけ始め、CI フリート全体の容量を消費する可能性のあるフィードバック ループが作成されます。それが起こると、待ち時間が急増し、エンジニアはテストが開始されるまでに 30 分から 1 時間以上待つこともあります。
## 不安定なテストは修正するのが難しい
明らかな疑問は次のとおりです: ** なぜ不安定なテストを修正するようエンジニアに依頼しないのですか? **
そうでした。それは決してうまくいきませんでした。そして、この背後には 2 つの主な理由があることに気付きました。
- まず、** テストが不安定になる理由は必ずしも明らかではありません。 ** それを見つけるには、多くの場合、テスト対象の機能とその実装の両方について深い理解が必要です。テストの元の作成者でさえ、通常、本当の原因を特定することはできません。それができるなら、そもそも不安定なテストを作成しなかっただろう。
- 第二に、エンジニアが原因を特定したとしても、**それを適切に修正するには広範な検証が必要です**。障害は断続的に発生するため、固定テストを 1 回実行しても何も証明されません。エンジニアは結局何時間も PR の子守をし、不安定さが実際に解消されていることを確認するために何度も再実行することになります。
これらすべてにより、不安定なテストを修正するのは高価な投資になります。経験豊富なエンジニア 1 人が 2 つを修理するには、多くの場合、丸 1 日かかります。完全には勝てない戦いでもあります。より多くの機能を構築するにつれて、CI の脆弱性を許容レベルに保つためにエンジニアリング時間を注ぎ続けなければなりません。そうしないと、新しい開発が完全にブロックされ、CI が使用できなくなる危険があります。
## AI を使用して不安定なテストを修正できるでしょうか?
ここまでで、テストが不安定になる理由がおそらくお分かりいただけたでしょう。

大規模なエンジニアリング プロジェクトにとっては高価であり、人間が修正するには費用もかかります。エージェント コーディングが主流になったとき、これは私たちが検討した最初のユースケースの 1 つでした。理由は次のとおりです。
- エージェントは人間よりもはるかに多くのコードベースとログを迅速にスキャンでき、不安定なテストの根本原因を見つけるのに優れています。
- エージェントは介入なしで何時間も修正の検証を子守できますが、人間はこれが苦手です。
そこで私たちは、クロード コードを使用して Kong ゲートウェイの不安定なテストを修正する実験を実行しました。
## 識別-修正-検証ループ
私たちが最初に必要としたのは、どのテストが不安定で、どれくらいの頻度で失敗するかを特定する方法でした。幸運なことに、チームはすでに Datadog の CI 可視化機能の上にダッシュボードを構築していて、最も不安定なテストとその失敗率を明確に把握できました。
私たちはそのダッシュボードを修正のベンチマークとして使用し、そのリストを使用してエージェントをシードしました。
数回のラウンドを経て最終的に決定したエージェント ワークフローは次のとおりです。
最上位には ** fix-flakes ** スキル、つまりオーパス上で実行されるオーケストレーション エージェントがあります。これはシード ファイル (不安定なテストの Datadog CSV エクスポート) を取得し、それを読み取って、そのテストの最近の失敗からのすべての CI ログをダウンロードします。このステップは重要です: f

[切り捨てられた]

## Original Extract

Learn how the Kong engineering team used agentic AI to identify, fix, and verify flaky tests in our codebase, improving CI stability and developer productivity.

How We Used Agentic AI to Automate Fixes for Flaky Tests | Kong Inc. DON'T MISS OUT ON API + AI SUMMIT 2026 | PRICES INCREASE AUGUST 16 SECURE YOUR SPOT FOR THE KONG AGENTIC ERA WORLD TOUR GOVERN A2A TRAFFIC WITH KONG'S NEW AGENT GATEWAY [ Why Kong ] ( /company/why-kong ) Why Kong
+ Products & Agents Products & Agents _ API & AI CONNECTIVITY TECHNOLOGIES _ The Unified API and AI Platform [ ] API Management AI Management Event Management Monetization Migration Services API Advisory Services + Forward Deployed Engineers NEW RUNTIMES
[ API Gateway ] ( /products/kong-gateway ) API Gateway
[ AI Gateway HOT ] ( /products/kong-ai-gateway ) AI Gateway HOT
[ Event Gateway ] ( /products/event-gateway ) Event Gateway
[ Service Mesh ] ( /products/kong-mesh ) Service Mesh
[ Context Mesh ] ( /products/kong-konnect/features/context-mesh ) Context Mesh
[ Ingress Controller ] ( /products/kong-ingress-controller ) Ingress Controller
[ Kong Operator ] ( /products/kong-operator ) Kong Operator
[ MCP Registry NEW ] ( /products/mcp-registry ) MCP Registry NEW
[ API Service Catalog ] ( /products/kong-konnect/features/api-service-catalog ) API Service Catalog
[ Runtime Management ] ( /products/kong-konnect/features/runtime-management ) Runtime Management
[ APIOps & Automation ] ( /products/apiops-automation ) APIOps & Automation
[ Developer Portal ] ( /products/kong-konnect/features/developer-portal ) Developer Portal
[ Usage Billing & Metering ] ( /products/kong-konnect/features/usage-based-metering-and-billing ) Usage Billing & Metering
[ Observability ] ( /products/kong-konnect/features/api-observability ) Observability
[ KAi Agent ] ( /products/kong-konnect/features/kai-ai-agent ) KAi Agent
+ Solutions Solutions FOR PLATFORM TEAMS
[ Developer Platform ] ( /solutions/building-developer-platform ) Developer Platform
[ Kubernetes and Microservices ] ( /solutions/build-on-kubernetes ) Kubernetes and Microservices
[ Observability ] ( /solutions/observability ) Observability
[ Service Mesh Connectivity ] ( /solutions/service-mesh-connectivity ) Service Mesh Connectivity
[ Kafka Event Streaming ] ( /solutions/kafka-stream-api-management ) Kafka Event Streaming
[ AI Connectivity ] ( /ai-connectivity ) AI Connectivity
[ Open Banking ] ( /solutions/open-banking ) Open Banking
[ Legacy Migration ] ( /solutions/legacy-api-management-migration ) Legacy Migration
[ Platform Cost Reduction ] ( /solutions/api-platform-consolidation ) Platform Cost Reduction
[ Kafka Cost Optimization ] ( /solutions/reduce-kafka-cost ) Kafka Cost Optimization
[ API Monetization ] ( /solutions/api-monetization ) API Monetization
[ AI Monetization ] ( /solutions/ai-monetization ) AI Monetization
[ AI FinOps ] ( /solutions/ai-cost-governance-finops ) AI FinOps
[ Agent Gateway ] ( /agent-gateway ) Agent Gateway
[ AI Governance ] ( /solutions/ai-governance ) AI Governance
[ AI Security ] ( /solutions/ai-security ) AI Security
[ AI Cost Control ] ( /solutions/ai-cost-optimization-management ) AI Cost Control
[ Agentic Infrastructure ] ( /solutions/agentic-ai-workflows ) Agentic Infrastructure
[ MCP Production ] ( /solutions/mcp-production-and-consumption ) MCP Production
[ MCP Traffic Gateway ] ( /solutions/mcp-governance ) MCP Traffic Gateway
[ Mobile App API Development ] ( /solutions/mobile-application-api-development ) Mobile App API Development
[ GenAI App Development ] ( /solutions/power-openai-applications ) GenAI App Development
[ API Gateway for Istio ] ( /solutions/istio-gateway ) API Gateway for Istio
[ Decentralized Load Balancing ] ( /solutions/decentralized-load-balancing ) Decentralized Load Balancing
[ Financial Services ] ( /solutions/financial-services-industry ) Financial Services
[ Healthcare ] ( /solutions/healthcare ) Healthcare
[ Higher Education ] ( /solutions/api-platform-for-education-services ) Higher Education
[ Insurance ] ( /solutions/insurance ) Insurance
[ Manufacturing ] ( /solutions/manufacturing ) Manufacturing
[ Retail ] ( /solutions/retail ) Retail
[ Software & Technology ] ( /solutions/software-and-technology ) Software & Technology
[ Transportation ] ( /solutions/transportation-and-logistics ) Transportation
[ Pricing ] ( /pricing ) Pricing
+ Resources Resources DOCUMENTATION
[ Kong Konnect ] ( https://developer.konghq.com/konnect/ ) Kong Konnect
[ Kong Gateway ] ( https://developer.konghq.com/gateway/ ) Kong Gateway
[ Kong Mesh ] ( https://developer.konghq.com/mesh/ ) Kong Mesh
[ Kong AI Gateway ] ( https://developer.konghq.com/ai-gateway/ ) Kong AI Gateway
[ Kong Event Gateway ] ( https://developer.konghq.com/event-gateway/ ) Kong Event Gateway
[ Kong Insomnia ] ( https://developer.konghq.com/insomnia/ ) Kong Insomnia
[ Plugin Hub ] ( https://developer.konghq.com/plugins/ ) Plugin Hub
[ Learning Center ] ( /blog/learning-center ) Learning Center
[ eBooks ] ( /resources/e-book ) eBooks
[ Reports ] ( /resources/reports ) Reports
[ Demos ] ( /resources/demos ) Demos
[ Customer Stories ] ( /customer-stories ) Customer Stories
[ Videos ] ( /resources/videos ) Videos
[ API + AI Summit ] ( /events/conferences/api-ai-summit ) API + AI Summit
[ Agentic Era World Tour ] ( /agentic-era-world-tour ) Agentic Era World Tour
[ Webinars ] ( /events/webinars ) Webinars
[ User Calls ] ( /events/user-calls ) User Calls
[ Workshops ] ( /events/workshops ) Workshops
[ Meetups ] ( /events/meetups ) Meetups
[ See All Events ] ( /events ) See All Events
[ Get Started ] ( https://developer.konghq.com/ ) Get Started
[ Community ] ( /community ) Community
[ Certification ] ( /academy/certification ) Certification
[ Training ] ( https://education.konghq.com ) Training
[ About Us ] ( /company/about-us ) About Us
[ We're Hiring! ] ( /company/careers ) We're Hiring!
[ Press Room ] ( /company/press-room ) Press Room
[ Contact Us ] ( /company/contact-us ) Contact Us
[ Kong Partner Program ] ( /partners ) Kong Partner Program
[ Enterprise Support Portal ] ( https://support.konghq.com/s/ ) Enterprise Support Portal
[ Documentation ] ( https://developer.konghq.com/?_gl=1*tphanb*_gcl_au*MTcxNTQ5NjQ0MC4xNzY5Nzg4MDY0LjIwMTI3NzEwOTEuMTc3MzMxODI2MS4xNzczMzE4MjYw*_ga*NDIwMDU4MTU3LjE3Njk3ODgwNjQ.*_ga_4JK9146J1H*czE3NzQwMjg1MjkkbzE4OSRnMCR0MTc3NDAyODUyOSRqNjAkbDAkaDA ) Documentation
[ Login ] ( https://cloud.konghq.com/login ) Login
[ Book Demo ] ( /contact-sales ) Book Demo
[ Get Started ] ( /products/kong-konnect/register ) Get Started
[ Blog ] ( /blog ) Blog [ AI Gateway ] ( /blog/tag/ai-gateway ) AI Gateway
[ AI Security ] ( /blog/tag/ai-security ) AI Security
[ AIOps ] ( /blog/tag/aiops ) AIOps
[ API Security ] ( /blog/tag/api-security ) API Security
[ API Gateway ] ( /blog/tag/api-gateway ) API Gateway
+ Explore Topics Explore Topics [ API Management ] ( /blog/tag/api-management ) API Management
[ API Development ] ( /blog/tag/api-development ) API Development
[ API Design ] ( /blog/tag/api-design ) API Design
[ Automation ] ( /blog/tag/automation ) Automation
[ Service Mesh ] ( /blog/tag/service-mesh ) Service Mesh
[ Insomnia ] ( /blog/tag/insomnia ) Insomnia
[ Event Gateway ] ( /blog/tag/event-gateway ) Event Gateway
[ View All Blogs ] ( /blog/page/1 ) View All Blogs
[ Kong AI Gateway ] ( /products/kong-ai-gateway ) Kong AI Gateway
[ Kong API Gateway ] ( /products/kong-gateway ) Kong API Gateway
[ Kong Event Gateway ] ( /products/event-gateway ) Kong Event Gateway
[ Kong Metering & Billing ] ( /products/kong-konnect/features/usage-based-metering-and-billing ) Kong Metering & Billing
[ Kong Insomnia ] ( /products/kong-insomnia ) Kong Insomnia
[ Kong Konnect ] ( /products/kong-konnect ) Kong Konnect
[ Documentation ] ( https://developer.konghq.com ) Documentation
[ Book Demo ] ( /contact-sales ) Book Demo
How We Used Agentic AI to Fix Kong Gateway's Flakiest Tests
# How We Used Agentic AI to Fix Kong Gateway's Flakiest Tests
Datong Sun
Each change to Kong Gateway's codebase triggers a comprehensive test suite that runs more than 17,000 * 2 = 34,000 test cases among the two primary architectures (x86 and ARM) we support. This process takes about 23.5 hours on a single machine. But we don't wait that long. A large fleet of machines runs the suite in parallel, and we shard the work aggressively so each commit finishes in a fraction of that time. That setup works well, right up until flaky tests get involved.
On the Kong engineering team, we've been exploring how an agentic AI workflow can make developing Kong Gateway more efficient. Fixing flaky tests turned out to be the ideal place to start.
Quality is at the core of every product we build. To move fast without compromising on quality, we rely on an extensive set of test suites that tell us quickly whether a change has caused a regression. Those suites are large enough that a single full run takes about 23.5 hours with a single machine, which is why we spread the work across a large fleet of machines.
But as we ship more features, the suite keeps growing, and more flaky tests start showing up.
For the uninitiated, flaky tests are tests that fail sometimes but usually pass on a rerun. They're unavailable on large projects and a constant drain on engineers: after a flaky failure, you have to rerun the job, which costs both engineering time and machine time. Worse, as the flaky rate climbs, engineers trying to merge PRs start stepping on each other during reruns, creating a feedback loop that can consume the entire CI fleet’s capacity. Once that happens, queueing time spikes — sometimes engineers wait more than 30 minutes to an hour before a test even starts.
## Flaky tests are hard to fix
The obvious question is: ** why not just ask engineers to fix the flaky tests? **
We did. It never worked well. And we realized there are two main reasons behind this:
- First, ** the reason a test is flaky is not always obvious. ** Finding it often takes a deep understanding of both the feature under test and its implementation. Even the test's original author usually can't spot the real cause — if they could, they wouldn't have written a flaky test in the first place.
- Second, even once an engineer identifies the cause, ** fixing it properly takes extensive validation ** . Because the failure is intermittent, running the fixed test once proves nothing. Engineers end up babysitting a PR for hours, rerunning it many times to confirm the flakiness is actually gone.
All of this makes fixing flaky tests an expensive investment. It often takes one experienced engineer a full day to fix two of them. It's also a battle we can't win outright. As we build more features, we have to keep pouring engineering time into keeping CI flakiness at an acceptable level, or risk blocking new development entirely and rendering CI unusable.
## Can we use AI to fix flaky tests?
By now you can probably see why flaky tests are expensive for any large engineering project, and expensive for humans to fix. When agentic coding went mainstream, this was one of the first use cases we considered for, because:
- Agents can scan far more of the codebase and logs than a human can, and quickly — a better shot at finding a flaky test's root cause.
- Agents can babysit the validation of a fix for hours without intervention, which humans are terrible at.
So we ran an experiment attempting to use Claude Code to fix the flaky tests in Kong Gateway.
## The identify-fix-verify loop
The first thing we needed was a way to identify which tests were flaky and how often they failed. Luckily, the team had already built a dashboard on top of Datadog's CI Visibility feature that gives us a clear picture of the flakiest tests and their failure rates.
We used that dashboard as the benchmark for our fixes, and used its list to seed the agents.
Here's the agentic workflow we settled on after a few rounds:
At the top level is the ** fix-flakes ** skill, the orchestrating agent, which runs on Opus. It takes a seed file — a flaky test's Datadog CSV export — and reads it to download all the CI logs from that test's recent failures. This step matters: f

[truncated]
