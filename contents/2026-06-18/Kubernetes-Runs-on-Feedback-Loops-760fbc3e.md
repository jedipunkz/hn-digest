---
source: "https://arslan.io/2026/06/16/kubernetes-runs-on-feedback-loops/"
hn_url: "https://news.ycombinator.com/item?id=48589293"
title: "Kubernetes Runs on Feedback Loops"
article_title: "Kubernetes runs on feedback loops"
author: "gpi"
captured_at: "2026-06-18T18:54:04Z"
capture_tool: "hn-digest"
hn_id: 48589293
score: 1
comments: 0
posted_at: "2026-06-18T18:17:02Z"
tags:
  - hacker-news
  - translated
---

# Kubernetes Runs on Feedback Loops

- HN: [48589293](https://news.ycombinator.com/item?id=48589293)
- Source: [arslan.io](https://arslan.io/2026/06/16/kubernetes-runs-on-feedback-loops/)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T18:17:02Z

## Translation

タイトル: Kubernetes はフィードバック ループで実行されます
記事のタイトル: Kubernetes はフィードバック ループで実行されます
説明: 過去数年間、私はほとんどの時間を Kubernetes 上にデータベース オペレーターを構築することに費やしてきました。その過程で、私は同じ考えに何度も戻りました。YAML と可動部分のすべての下では、Kubernetes は実際にはフィードバック コントローラーのためのフレームワークであるということです。それは制御理論です

記事本文:
2026 年 6 月 16 日
Kubernetes はフィードバック ループで実行されます
ここ数年、私はほとんどの時間を Kubernetes 上にデータベース オペレーターを構築することに費やしてきました。その過程で、私は同じ考えに何度も戻りました。YAML と可動部分のすべての下では、Kubernetes は実際にはフィードバック コントローラーのためのフレームワークであるということです。制御理論を一般の人でも使えるようにしたものです。サーモスタットを適切な温度に保つのと同じ閉ループが、データベースを稼働状態に保つものでもあります。
そこで私はついに腰を据えて、それについて PlanetScale ブログに書きました。
この記事は、手動で Postgres を実行し、ノードを選択し、ディスクを接続し、レプリケーションを設定し、障害が発生したときに修正するための小さなウォッチドッグ スクリプトを作成することから始まります。次に、これらの手動ステップがどのように Kubernetes 内の制御ループにマッピングされるか、およびオペレーターがそのように構築される理由について説明します。途中で制御理論にも少し触れます。 Kubernetes を使用していなくても、楽しく読めると思います。
ここで読むことができます: The Feedback Loops Behind Kubernetes 。

## Original Extract

For the past few years I've spent most of my time building database operators on top of Kubernetes. Along the way I kept coming back to the same idea: under all the YAML and the moving parts, Kubernetes is really a framework for feedback controllers. It's control theory

Jun 16, 2026
Kubernetes runs on feedback loops
For the past few years I've spent most of my time building database operators on top of Kubernetes. Along the way I kept coming back to the same idea: under all the YAML and the moving parts, Kubernetes is really a framework for feedback controllers. It's control theory turned into something ordinary people can use. The same closed loop that keeps a thermostat at the right temperature is also the one that keeps your database alive.
So I finally sat down and wrote about it on the PlanetScale blog.
The post starts from running Postgres by hand, picking a node, attaching a disk, setting up replication, and writing a small watchdog script to fix things when they break. Then it walks through how each of those manual steps maps to a control loop inside Kubernetes, and why operators are built the way they are. It dips into a bit of control theory along the way. I think even if you don't use Kubernetes, you'll enjoy reading it.
You can read it here: The Feedback Loops Behind Kubernetes .
