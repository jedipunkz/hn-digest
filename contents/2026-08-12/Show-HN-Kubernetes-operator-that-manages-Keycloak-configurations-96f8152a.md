---
source: ""
hn_url: "https://news.ycombinator.com/item?id=49273393"
title: "Show HN: Kubernetes operator that manages Keycloak configurations"
article_title: ""
author: "soaringmonchi"
captured_at: "2026-08-12T15:52:01Z"
capture_tool: "hn-digest"
hn_id: 49273393
score: 2
comments: 0
posted_at: "2026-08-12T14:55:19Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Kubernetes operator that manages Keycloak configurations

- HN: [49273393](https://news.ycombinator.com/item?id=49273393)
- Score: 2
- Comments: 0
- Posted: 2026-08-12T14:55:19Z

## Translation

タイトル: Show HN: Keycloak 構成を管理する Kubernetes オペレーター
HN テキスト: 私たちはついに、Kubernetes CRD を通じて GitOps 方式で Keycloak 構成全体を管理できる Keycloak 構成オペレーター (もっと良い名前をまだ見つけていません) をオープンソース化しました。開始は簡単です。現在の構成から CRD をエクスポートできるエクスポート フローがあり、必要に応じて適用できます。リソースは意図的に生の JSON であるため (これについての議論をぜひ見てみたいと思います)、すぐに使用できる Keycloak の幅広いバージョンと互換性があります。現在は廃止されている keycloak レルム オペレーターから大きくインスピレーションを得ています。改善のためのフィードバックや提案は大歓迎です。 https://github.com/Hostzero-GmbH/keycloak-operator

## Original Extract

We have finally open sourced our Keycloak config operator (I have yet to find a better name) that allows to manage your entire Keycloak configuration in a GitOps fashion through Kubernetes CRDs. Starting is easy, there is an export flow which allows you to export CRDs from your current configuration and then you can apply them as needed. The resources are deliberately raw JSON (I'd love to see a discussion on this) so that its compatible with a wide range of Keycloak versions out of the box. Heavily inspired by the now defunct keycloak realm operator. Feedback and suggestions for improvements are highly welcome. https://github.com/Hostzero-GmbH/keycloak-operator

