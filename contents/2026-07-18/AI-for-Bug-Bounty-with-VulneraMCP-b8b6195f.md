---
source: "https://www.zaproxy.org/blog/2025-11-28-enhancing-zap-with-ai-for-bug-bounty-hunting/"
hn_url: "https://news.ycombinator.com/item?id=48962545"
title: "AI for Bug Bounty with VulneraMCP"
article_title: "ZAP – Enhancing ZAP with AI for Bug Bounty Hunting"
author: "Fermino"
captured_at: "2026-07-18T21:40:42Z"
capture_tool: "hn-digest"
hn_id: 48962545
score: 1
comments: 0
posted_at: "2026-07-18T21:20:59Z"
tags:
  - hacker-news
  - translated
---

# AI for Bug Bounty with VulneraMCP

- HN: [48962545](https://news.ycombinator.com/item?id=48962545)
- Source: [www.zaproxy.org](https://www.zaproxy.org/blog/2025-11-28-enhancing-zap-with-ai-for-bug-bounty-hunting/)
- Score: 1
- Comments: 0
- Posted: 2026-07-18T21:20:59Z

## Translation

タイトル: VulneraMCP によるバグ報奨金の AI
記事のタイトル: ZAP – バグ報奨金ハンティングのための AI による ZAP の強化
説明: 世界で最も広く使用されている Web アプリ スキャナー。無料でオープンソース。 ZAP は、熱心な国際チームによって積極的に維持されているコミュニティ プロジェクトであり、GitHub トップ 1000 プロジェクトです。

記事本文:
バグ報奨金ハンティングのための AI による ZAP の強化
最新のセキュリティ テスト プラットフォームのほとんどは、高度な自動化、相関、ワークフロー機能を高価なライセンス層の背後に配置しています。過去 4 か月間、セキュリティ調査とバグ報奨金の手法を研究してきたフルスタック エンジニアとして、私は柔軟性、拡張性、ベンダー ロックインのない完全なプログラム制御を提供するツールを必要としていました。
ZAP はすぐに理想的な基盤として浮上しました。そのオープンソースの性質、堅牢な REST API、専用のコミュニティにより、従来のスキャンを超えるシステムを設計するために必要な自由度がまさに得られました。数か月にわたる手動テストとさまざまなツールの実験を経て、ZAP をスキャン エンジンとして使用し、その上に機械学習とインテリジェントなワークフロー オーケストレーションを重ねる、AI 拡張セキュリティ テスト プラットフォームの構築を開始しました。
ZAP は、カスタム ソリューションへの適応性を根本的に高める機能を提供します。
広範な REST API による完全な自動化
ZAP の内部コードベースを変更することなく完全な拡張性を実現
コミュニティ主導の開発。継続的な更新と高度なスクリプトが利用可能
ライセンス制限がないため、無制限のカスタマイズと統合が可能
ZAP はコア スキャン機能 (アクティブ スキャン、パッシブ スキャン、スパイダリング、アラート収集、コンテキスト管理) を実行しますが、私のシステムには現実世界の悪用テクニックから学習したインテリジェンス層が導入されています。
このシステムは、モデル コンテキスト プロトコル (MCP) を通じて ZAP を AI 主導の学習エンジンと統合します。このアーキテクチャにより、AI エージェントは、より深い分析、適応ペイロード生成、学習された脆弱性パターンを組み込みながら、プログラムで ZAP と対話できるようになります。
┌──────

─────┐
│ AI エージェント │ (MCP クライアント: Cursor、ChatGPT など)
━───┬───┘
│
│ MCPプロトコル
│
┌─────────▼───────┐
│ ヴァルネラMCP │
│ ┌───────────────┐ │
│ │ ZAP 統合層 │ │
│ ━━━━━━━━━━━━━━━━━━━━┘ │
│ ┌───────────────┐ │
│ │ MCP プロキシ層 │ │
│ ━━━━━━━━━━━━━━━━━━━━┘ │
│ ┌───────────────┐ │
│ │ 学習エンジン │ │
│ ━━━━━━━━━━━━━━━━━━━━┘ │
━━━━━━━━━━━━━━━━━┘
│
┌────┴────┐
│ │
┌───▼───┐ ┌──▼─────┐
│ ZAP │ │Postgres│
│ │ │ DB │
━━━━━┘ ━━━━┘
コンポーネント
ZAP統合レイヤー
スパイダリング、アクティブ スキャン、コンテキスト管理、アラートの取得など、ZAP とのすべての対話を処理します。
ヴァルネラMCP
トラフィックを傍受して分析し、ZAP の組み込みルールを超えたカスタム脆弱性チェック (IDOR、ロジック欠陥など) を可能にします。
学習エンジン
HackTheBox、PortSwigger Academy、および実際のバグ報奨金の書き込みからトレーニング データをインポートします。パターンを抽出し、ペイロードを生成し、検出精度を継続的に向上させます。
データベース層
知識を蓄える

基本エントリ、学習データ、スキャン結果、悪用パターン。
ZAP - 無料、スクリプト可能、オープンソース
MCP - AI 主導のインタラクション レイヤー
Postgres - 学習データ、スキャン結果、悪用パターンの保存用
Docker - コンテナ化されたスキャナー + オフライン操作
プラットフォームは REST API を通じて ZAP を完全に制御します。例としては次のものが挙げられます。
// スパイダーリングを開始する
const スパイダー = zapClient を待ちます。 startSpider ( 'https://example.com' );
// スパイダーのステータスを確認する
const status = zapClient を待ちます。 getSpiderStatus ( スパイダー . データ . scanId );
// アクティブスキャンを起動します
const active = zapClient を待ちます。 startActiveScan ( 'https://example.com' );
// 高リスクのアラートを取得します
const アラート = zapClient を待ちます。 getAlerts ( 'https://example.com' 、未定義、未定義、 '3' );
これにより、手動操作を必要とせずに完全に自動化されたテスト パイプラインが可能になります。
このシステムの主な差別化要因は、適応学習モジュールです。現実世界の悪用データを組み込んで、今後のスキャンの精度と有効性を向上させます。
PortSwigger Academy ラボ ソリューション
カスタムの調査とテストの結果
エンジンはトレーニング データからエクスプロイト パターンを抽出します。
const training = await getTrainingData ( 'xss' );
const パターン = extractPatterns (トレーニング);
これらのパターンは適応され、新しいターゲットに適用されます。
静的ペイロード リストに依存するスキャナとは異なり、このシステムは以下に基づいて動的ペイロードを生成します。
過去に成功した悪用の試み
これにより、高度な脆弱性を検出できる可能性が大幅に高まります。
ZAP スパイダリングと URL 列挙により、アプリケーションの完全なマップが構築されます。
アクティブおよびパッシブ スキャンが開始され、IDOR や弱い認証フローなどの問題に対するカスタム ルールが強化されます。
MCP プロキシ層は、リクエスト/レスポンス パターンを評価し、結果を ZAP アラートと関連付けます。

d 学習したルールを適用します。
このエンジンは、改善されたペイロードを生成し、新しいエクスプロイト シグネチャを抽出し、ナレッジ ベースを更新します。
調査結果は集計され、スコア付けされ、証拠と推奨される改善策を含む構造化された出力として生成されます。
ZAP のスキャン エンジンと機械学習を組み合わせることで、システムは次のことを実現します。
従来のスキャナーが見逃しがちな脆弱性を検出します
パターン相関により誤検知を削減
さまざまなアプリケーションの構造と動作に適応します
このシステムは、手作業なしで偵察、スキャン、ペイロード テスト、関連付け、レポートを処理します。
完全にオープンソース コンポーネントに基づいて構築されており、以下を使用して拡張できます。
外部統合 (バープ、核、サブファインダーなど)
docker run -d -p 8081:8080 owasp/zap2docker-stable \
zap.sh -daemon -host 0.0.0.0 -port 8080 \
-config api.disablekey = true
VulneraMCP の機能
VulneraMCP は次のツールを提供します。
ペイロード テスト (XSS、SQLi、IDOR、CSRF など)
エアギャップ環境向けのオフライン モード
脆弱性推論 (AI が調査結果を説明)
自動偵察 + 攻撃対象領域のマッピング
偵察とテストの実行に必要な時間を短縮しました
適応学習による検出精度の向上
バグ報奨金ワークフローをテスト、トレーニング、スケールするためのエコシステムを提供
AI エージェントとのシームレスな統合を可能にし、高度な推論と分析を実現
GitHub リポジトリは https://github.com/telmon95/VulneraMCP です。
Telmon Maluleka は、南アフリカのプレトリアに拠点を置くフルスタック ソフトウェア エンジニアです。 C、Python、JavaScript、HTML、CSS のスキルがあり、React、Node.js、Django、さまざまな AWS クラウド サービスなどのフレームワークの使用経験があります。彼の経歴には、API 開発、データベース設計 (MySQL)、コンテナ化されたアプリケーションのデプロイメントが含まれます。
過去 4 か月間にわたって、彼は焦点を電子分野に拡大してきました。

倫理的なハッキングと実践的な脆弱性研究。 Docker、MCP サーバー、大規模言語モデルに関する彼の経験は、この AI を活用したセキュリティ テスト プラットフォームの開発に直接貢献しました。このプロジェクトは、フルスタック エンジニアリングと最新のセキュリティ研究を融合する、彼の最初のオープンソース セキュリティへの主要な貢献を表しています。
ポートフォリオ : telmon95.github.io/portfoliov2/
Twitter : x.com/DEOXYRIBOSE404

## Original Extract

The world’s most widely used web app scanner. Free and open source. ZAP is a community project actively maintained by a dedicated international team, and a GitHub Top 1000 project.

Enhancing ZAP with AI for Bug Bounty Hunting
Most modern security testing platforms place advanced automation, correlation, and workflow features behind expensive licensing tiers. As a full-stack engineer who has spent the last four months studying security research and bug bounty methodologies, I needed a tool that offered flexibility, extensibility, and complete programmatic control without vendor lock-in.
ZAP quickly emerged as the ideal foundation. Its open-source nature, robust REST API, and dedicated community provided exactly the level of freedom I needed to design a system that goes beyond traditional scanning. After months of manual testing and experimenting with various tools, I began building an AI-augmented security testing platform that uses ZAP as the scanning engine and layers machine learning and intelligent workflow orchestration on top.
ZAP offers capabilities that make it fundamentally more adaptable for custom solutions:
Full automation through an extensive REST API
Complete extensibility without requiring modifications to ZAP’s internal codebase
Community-driven development , with continuous updates and advanced scripts available
No licensing limitations , allowing unrestricted customization and integration
ZAP performs the core scanning functions—active scanning, passive scanning, spidering, alert collection, and context management—while my system introduces the intelligence layer that learns from real-world exploitation techniques.
The system integrates ZAP with an AI-driven learning engine through the Model Context Protocol (MCP). This architecture enables AI agents to interact with ZAP programmatically while incorporating deeper analysis, adaptive payload generation, and learned vulnerability patterns.
┌─────────────────┐
│ AI Agent │ (MCP Clients: Cursor, ChatGPT, etc.)
└────────┬────────┘
│
│ MCP Protocol
│
┌─────────────────▼────────────────┐
│ VulneraMCP │
│ ┌──────────────────────────┐ │
│ │ ZAP Integration Layer │ │
│ └──────────────────────────┘ │
│ ┌──────────────────────────┐ │
│ │ MCP Proxy Layer │ │
│ └──────────────────────────┘ │
│ ┌──────────────────────────┐ │
│ │ Learning Engine │ │
│ └──────────────────────────┘ │
└─────────────────┬────────────────┘
│
┌────┴────┐
│ │
┌───▼───┐ ┌──▼─────┐
│ ZAP │ │Postgres│
│ │ │ DB │
└───────┘ └────────┘
Components
ZAP Integration Layer
Handles all interactions with ZAP, including spidering, active scanning, context management, and alert retrieval.
VulneraMCP
Intercepts and analyzes traffic, enabling custom vulnerability checks (e.g., IDOR, logic flaws) that extend beyond ZAP’s built-in rules.
Learning Engine
Imports training data from HackTheBox, PortSwigger Academy, and real bug bounty writeups. Extracts patterns, generates payloads, and continuously improves detection accuracy.
Database Layer
Stores knowledge base entries, learning data, scan results, and exploit patterns.
ZAP - free, scriptable, open-source
MCP - AI-driven interaction layer
Postgres - for storing learning data, scan results, and exploit patterns
Docker - containerized scanner + offline operation
The platform controls ZAP entirely through the REST API. Examples include:
// Start spidering
const spider = await zapClient . startSpider ( 'https://example.com' );
// Check spider status
const status = await zapClient . getSpiderStatus ( spider . data . scanId );
// Launch active scan
const active = await zapClient . startActiveScan ( 'https://example.com' );
// Retrieve high-risk alerts
const alerts = await zapClient . getAlerts ( 'https://example.com' , undefined , undefined , '3' );
This enables a fully automated testing pipeline with no manual interaction required.
A key differentiator of this system is the adaptive learning module. It incorporates real-world exploitation data to improve the accuracy and effectiveness of future scans.
PortSwigger Academy lab solutions
Custom research and test results
The engine extracts exploit patterns from training data:
const training = await getTrainingData ( 'xss' );
const patterns = extractPatterns ( training );
These patterns are then adapted and applied to new targets.
Unlike scanners that rely on static payload lists, this system generates dynamic payloads based on:
Previous successful exploit attempts
This significantly increases the chances of detecting sophisticated vulnerabilities.
ZAP spidering and URL enumeration build a complete map of the application.
Active and passive scanning begins, enriched with custom rules for issues like IDOR and weak authentication flows.
The MCP proxy layer evaluates request/response patterns, correlates findings with ZAP alerts, and applies learned rules.
The engine generates improved payloads, extracts new exploit signatures, and updates the knowledge base.
Findings are aggregated, scored, and produced in a structured output with evidence and recommended remediation.
By combining ZAP’s scanning engine with machine learning, the system:
Detects vulnerabilities traditional scanners commonly miss
Reduces false positives through pattern correlation
Adapts to different application structures and behaviors
The system handles reconnaissance, scanning, payload testing, correlation, and reporting without manual effort.
Built entirely on open-source components, it can be extended with:
External integrations (Burp, nuclei, Subfinder, etc.)
docker run -d -p 8081:8080 owasp/zap2docker-stable \
zap.sh -daemon -host 0.0.0.0 -port 8080 \
-config api.disablekey = true
VulneraMCP Capabilities
VulneraMCP provides tools for:
Payload testing (XSS, SQLi, IDOR, CSRF, and more)
Offline mode for air-gapped environments
Vulnerability reasoning (AI explains findings)
Automated recon + attack surface mapping
Reduced the time needed to perform reconnaissance and testing
Increased detection accuracy through adaptive learning
Provided an ecosystem to test, train, and scale bug bounty workflows
Enabled seamless integration with AI agents for advanced reasoning and analysis
The GitHub repository is https://github.com/telmon95/VulneraMCP
Telmon Maluleka is a Full-Stack Software Engineer based in Pretoria, South Africa. Skilled in C, Python, JavaScript, HTML, and CSS, with experience using frameworks such as React, Node.js, Django, and various AWS cloud services. His background includes API development, database design (MySQL), and containerized application deployment.
Over the past four months, he has expanded his focus into ethical hacking and practical vulnerability research. His experience with Docker, MCP servers, and large language models directly contributed to developing this AI-powered security testing platform. This project represents his first major open-source security contribution, merging full-stack engineering with modern security research.
Portfolio : telmon95.github.io/portfoliov2/
Twitter : x.com/DEOXYRIBOSE404
