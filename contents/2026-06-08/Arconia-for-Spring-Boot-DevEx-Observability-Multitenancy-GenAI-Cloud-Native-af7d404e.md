---
source: "https://arconia.io/"
hn_url: "https://news.ycombinator.com/item?id=48441539"
title: "Arconia for Spring Boot: DevEx, Observability, Multitenancy, GenAI, Cloud Native"
article_title: "Arconia — Boosting Spring Boot"
author: "thomasvitale"
captured_at: "2026-06-08T07:51:39Z"
capture_tool: "hn-digest"
hn_id: 48441539
score: 3
comments: 0
posted_at: "2026-06-08T05:12:59Z"
tags:
  - hacker-news
  - translated
---

# Arconia for Spring Boot: DevEx, Observability, Multitenancy, GenAI, Cloud Native

- HN: [48441539](https://news.ycombinator.com/item?id=48441539)
- Source: [arconia.io](https://arconia.io/)
- Score: 3
- Comments: 0
- Posted: 2026-06-08T05:12:59Z

## Translation

タイトル: Arconia for Spring Boot: DevEx、可観測性、マルチテナンシー、GenAI、クラウドネイティブ
記事のタイトル: Arconia — Spring Boot のブースト
説明: Arconia は、Spring Boot 用のオープンソースのモジュール式ツールキットです。優れた開発者エクスペリエンスを実現する開発サービス、可観測性、GenAI、クラウド ネイティブ ツール。

記事本文:
Arconia は Spring Boot 用のモジュール式ツールキットです。開発サービス、可観測性、生成 AI、クラウド ネイティブ ツール。必要なものを選んでください。優れた開発者エクスペリエンスを手に入れましょう。
最新の Java 開発向けに構築
依存関係を 1 つ追加します。アプリを実行します。 Arconia は、コンテナ化された PostgreSQL、Grafana、Kafka、Redis、Ollama、およびその他の 20 以上のサービスを、コードも設定もゼロで自動プロビジョニングし、接続します。
testAndDevelopmentOnly 'io.arconia:arconia-dev-services-postgresql'
$ ./gradlew bootRun
OpenTelemetry による統合された可観測性
依存関係が 1 つあります。ログ、メトリクス、トレース。 OTLP 経由で計測およびエクスポートされます。 Spring の Micrometer 可観測性基盤に基づいて構築されています。
実装「io.arconia:arconia-opentelemetry-spring-boot-starter」
ファーストクラスの GenAI 可観測性
Spring AI アプリの OpenTelemetry GenAI セマンティック規則。 1 行の設定で、フレーバー (OpenLit、OpenLLMetry、または LangSmith) を切り替えます。
# フレーバー: opentelemetry |オープンライト |オープンメトリ |ラングスミス
arconia.observations.conventions.opentelemetry.ai.flavor: openlit
ライフサイクル全体で 1 つの CLI を使用
単一のツール、Gradle または Maven、任意の OS を使用して、Spring Boot アプリを作成、実行、テスト、ビルドします。 brew install arconia-io/tap/arconia-cli 経由でインストールします。
$アルコニア作成
$ アルコニア開発者
$アルコニアテスト
$アルコニアビルド
ワンコマンドでのバージョンアップグレード
Spring Boot、Spring AI、Spring Cloud、JUnit、Testcontainers、および Arconia 自体の厳選された OpenRewrite レシピ。 CLI にバンドルされているため、セットアップは必要ありません。
$ arconia update spring-boot --to-version 4.0
$ arconia update spring-ai --to-version 2.0
プロジェクト
アルコニアの生態系
Arconia の中核となるモジュラー Spring Boot アドオン。 Dev Services、OpenTelemetry、Generative AI、マルチテナンシー、および Kubernetes 用のモジュール。
Spring Boot 開発ライフサイクル全体に対応する、ビルドツールに依存しない CLI。作成する

、マルチ アーキテクチャ イメージの開発、テスト、構築、配布、バージョン間での更新、エージェントのスキルの管理を行います。
Spring エコシステムおよびそれ以降の OpenRewrite レシピ。 Spring Boot、Spring AI、Spring Cloud、JUnit、Testcontainers、Docling、および Arconia 自体。
あなたが構築するアプリと、すでに所有しているアプリのためのオープンソース ツールキット。
© 2025-2026 アルコニア。無断転載を禁じます。
ヨーロッパでは、オープンソース コミュニティからの貢献を得て、Thomas Vitale によって ❤️ を使用して作成されました。

## Original Extract

Arconia is an open-source, modular toolkit for Spring Boot. Dev services, observability, GenAI, and cloud native tooling for a great developer experience.

Arconia is a modular toolkit for Spring Boot. Dev services, observability, generative AI, cloud native tooling. Pick what you need. Get a great developer experience.
Built for Modern Java Development
Add one dependency. Run your app. Arconia auto-provisions and wires a containerized PostgreSQL, Grafana, Kafka, Redis, Ollama — and 20+ other services — with zero code and zero configuration.
testAndDevelopmentOnly 'io.arconia:arconia-dev-services-postgresql'
$ ./gradlew bootRun
Unified Observability with OpenTelemetry
One dependency. Logs, metrics, and traces. Instrumented and exported via OTLP. Built on Spring's Micrometer observability foundation.
implementation 'io.arconia:arconia-opentelemetry-spring-boot-starter'
First-Class GenAI Observability
OpenTelemetry GenAI semantic conventions for Spring AI apps. Switch flavours — OpenLit, OpenLLMetry, or LangSmith — with one line of config.
# Flavor: opentelemetry | openlit | openllmetry | langsmith
arconia.observations.conventions.opentelemetry.ai.flavor: openlit
One CLI for the Whole Lifecycle
Create, run, test, and build Spring Boot apps with a single tool, Gradle or Maven, any OS. Install via brew install arconia-io/tap/arconia-cli .
$ arconia create
$ arconia dev
$ arconia test
$ arconia build
One-Command Version Upgrades
Curated OpenRewrite recipes for Spring Boot, Spring AI, Spring Cloud, JUnit, Testcontainers, and Arconia itself. Bundled with the CLI, no setup required.
$ arconia update spring-boot --to-version 4.0
$ arconia update spring-ai --to-version 2.0
Projects
The Arconia Ecosystem
The modular Spring Boot add-on at the core of Arconia. Modules for Dev Services, OpenTelemetry, Generative AI, Multitenancy, and Kubernetes.
Build-tool-agnostic CLI for the full Spring Boot dev lifecycle. Create, develop, test, build, ship multi-arch images, update across versions, and manage agent skills.
OpenRewrite recipes for the Spring ecosystem and beyond. Spring Boot, Spring AI, Spring Cloud, JUnit, Testcontainers, Docling, and Arconia itself.
An open-source toolkit for the apps you build, and the ones you already have.
© 2025-2026 Arconia. All rights reserved.
Made with ❤️ in Europe by Thomas Vitale , with contributions from the open-source community .
