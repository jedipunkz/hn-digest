---
source: ""
hn_url: "https://news.ycombinator.com/item?id=49230639"
title: "Show HN: Xprem: Self-hosted OTA updates and observability for Expo apps"
article_title: ""
author: "madscripten"
captured_at: "2026-08-09T12:29:38Z"
capture_tool: "hn-digest"
hn_id: 49230639
score: 1
comments: 0
posted_at: "2026-08-09T12:03:18Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Xprem: Self-hosted OTA updates and observability for Expo apps

- HN: [49230639](https://news.ycombinator.com/item?id=49230639)
- Score: 1
- Comments: 0
- Posted: 2026-08-09T12:03:18Z

## Translation

タイトル: Show HN: Xprem: Expo アプリのセルフホスト型 OTA アップデートと可観測性
HN テキスト: こんにちは、HN、私は xprem (以前は expo-open-ota) のメンテナーです。これは、Expo アプリの OTA アップデートを管理するためのオープンコアの自己ホスト型ソリューションです。サーバーは expo-updates プロトコル ( https://docs.expo.dev/technical-specs/expo-updates-1/ ) を実装します。このプロジェクトは 1 年以上前に始まりました。その理由は、高度に安全なネットワーク内でビジネス アプリを配布する必要があったため、オープンなインターネット アクセスのないアプリ用に「オンプレミス」OTA サーバーを実行する必要があったからです。それ以来、プロジェクトは大幅に成長し、特に expo-observe サポートの追加により、ログ/メトリクスを独自のサーバーで直接取得し、独自の ClickHouse DB に保存できるようになりました。また、パフォーマンスにも多くの時間を費やしました。最近のベンチマークでは、1000 req/s (テレメトリ有効) で p95 < 20ms で、単一の vCPU 上で 100 万 MAU のトラフィックをシミュレートしました。ベンチマーク: https://xprem.dev/benchmark (データとそれを再現するスクリプトはリポジトリで入手可能です)。ウェブサイト : https://xprem.dev リポジトリ : https://github.com/mercuretechnologies/xprem ドキュメント : https://mercure-technologies.gitbook.io/xprem プロジェクトに関するフィードバックをお待ちしています。

## Original Extract

Hi HN, I'm the maintainer of xprem (previously expo-open-ota). It's an open-core, self-hosted solution to manage OTA updates for Expo apps. The server implements the expo-updates protocol ( https://docs.expo.dev/technical-specs/expo-updates-1/ ). The project started more than a year ago because I needed to ship a business app inside a highly secured network, so I had to run an "on-premise" OTA server for apps without open internet access. Since then the project has grown a lot, in particular with the addition of expo-observe support, so you can get your logs/metrics directly in your own server, stored in your own ClickHouse DB. I also spent a lot of time on performance: in a recent benchmark we simulated traffic from 1 million MAU on a single vCPU, with a p95 < 20ms at 1000 req/s (telemetry enabled). Benchmark : https://xprem.dev/benchmark (the data and the script to reproduce it are available in the repository). Website : https://xprem.dev Repository : https://github.com/mercuretechnologies/xprem Documentation : https://mercure-technologies.gitbook.io/xprem I'd love to get your feedback on the project.

