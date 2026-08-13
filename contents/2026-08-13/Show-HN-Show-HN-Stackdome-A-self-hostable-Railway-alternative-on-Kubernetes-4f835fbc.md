---
source: "https://stackdome.com"
hn_url: "https://news.ycombinator.com/item?id=49285086"
title: "Show HN: Show HN: Stackdome – A self-hostable Railway alternative on Kubernetes"
article_title: "Stackdome - Breeze through deploys"
author: "ashishmax31"
captured_at: "2026-08-13T12:45:22Z"
capture_tool: "hn-digest"
hn_id: 49285086
score: 2
comments: 0
posted_at: "2026-08-13T12:43:04Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Show HN: Stackdome – A self-hostable Railway alternative on Kubernetes

- HN: [49285086](https://news.ycombinator.com/item?id=49285086)
- Source: [stackdome.com](https://stackdome.com)
- Score: 2
- Comments: 0
- Posted: 2026-08-13T12:43:04Z

## Translation

タイトル: HN を表示: HN を表示: Stackdome – Kubernetes 上の自己ホスト可能な鉄道代替手段
記事のタイトル: Stackdome - Breeze through デプロイ
説明: Stackdome は、エージェントと人間のためのオープンソース アプリケーション配信プラットフォームです。お客様が所有するクラウドまたはクラスター上のエージェント、CLI、またはキャンバスから発送します。
HN text: Hello Hacker News, this is Ashish.私たちは 2024 年に趣味のプロジェクトとして Stackdome の構築を開始しました。そのアイデアは、CNCF/クラウドネイティブ エコシステムのツールを独自のプラットフォームにパッケージ化し、開発者が Kubernetes YAMLology や CNCF 環境からのその他の多数のプロジェクトを学ばなくても、適切に設計されたプラットフォームを入手できるようにすることでした。 Railway と Render は、ホストされている製品でこれをうまく行いました。私は、セルフホストできるものにそのレベルの洗練と DX をもたらしたかったのです。 Stackdome はその下で Kubernetes を使用しますが、ユーザーには完全に透過的です。 (追加機能をさらに公開するエキスパート モードに関するアイデアがいくつかあります。) 私の日常の仕事にはクラウド ネイティブ テクノロジを扱うことが含まれており、Kubernetes とその周りのツールによってすでに多くのことが解決されていることに感謝しました。すでに Kubernetes を実行している場合は、Stackdome を段階的に導入できます。 (すでに実行中のワークロードを採用することについていくつかのアイデアがあります。) 特徴: - マルチサービス アプリケーションのファースト クラスのサポート。UI のキャンバスに表示されます (鉄道からインスピレーションを受けています)。キャンバスからリソースを編集、接続、作成できます。 - First class support for multiple clusters. Stackdome はハブとして機能し、クラスター全体 (ハブとスポーク) にワークロードを配置します。 - 小規模な VPS から大規模なクラスターまで自己ホスト可能。 - WAL アーカイブ、HA、バックアップ、PITR (その下に CNPG) を備えたマネージド Postgres。 - A release system for the whole stack.リリースはすべてを対象にしており、ロールバックできます

何かが壊れた場合。 - クラスターレジストリ (Zot) を使用したソースからイメージ。 - 使い捨てのプレビュー環境。マルチサービスにすることもできます。 - スタックファイル経由の宣言型アプリ。CLI でデプロイされます。 Web サイト: https://stackdome.com ドキュメント: https://docs.stackdome.com 弊社のクラウドで試してみてください。クレジット カードやコンピューティングの持参は必要ありません: https://cloud.stackdome.com フィードバックやご質問をお聞かせください。

記事本文:
Stackdome - ブリーズスルーデプロイ

## Original Extract

Stackdome is the open-source application delivery platform for agents and humans. Ship from your agent, the CLI, or the Canvas, on our cloud or clusters you own.

Hello Hacker News, this is Ashish. We started building Stackdome in 2024 as a hobby project. The idea was to package the tooling in the CNCF/cloud-native ecosystem into an opinionated platform, so that developers can get a well engineered platform without learning Kubernetes YAMLology or a dozen other projects from the CNCF landscape. Railway and Render did this well for their hosted products. I wanted to bring that level of polish and DX to something you can self-host. Stackdome uses Kubernetes underneath, but it's completely transparent to users. (I have some ideas for an expert mode that exposes more of the bells and whistles.) My day job involves working with cloud native technologies, and it made me appreciate how much is already solved by Kubernetes and the tooling around it. If you are already running Kubernetes, Stackdome is something you can adopt incrementally. (I have some ideas around adopting already-running workloads.) Features: - First class support for multi-service applications, shown on a canvas in the UI (inspired by Railway). You can edit, wire up, and create resources from the canvas. - First class support for multiple clusters. Stackdome acts as the hub and places workloads across clusters (hub and spoke). - Self hostable, from a small VPS to a large cluster. - Managed Postgres with WAL archiving, HA, backups, and PITR (CNPG underneath). - A release system for the whole stack. A release targets everything, and you can roll it back if something breaks. - Source to image, with an in cluster registry (Zot). - Disposable preview environments, which can also be multi-service. - Declarative apps via a Stackfile, deployed with the CLI. Website: https://stackdome.com Docs: https://docs.stackdome.com Try it on our cloud, no credit card and no compute to bring: https://cloud.stackdome.com Happy to hear your feedback and questions.

Stackdome - Breeze through deploys
