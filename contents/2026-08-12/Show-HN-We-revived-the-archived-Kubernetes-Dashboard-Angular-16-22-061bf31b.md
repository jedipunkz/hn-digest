---
source: "https://github.com/isms-core-project/kubernetes-dashboard"
hn_url: "https://news.ycombinator.com/item?id=49278224"
title: "Show HN: We revived the archived Kubernetes Dashboard (Angular 16 → 22)"
article_title: "GitHub - isms-core-project/kubernetes-dashboard: Original Kubernetes Dashboard — retired upstream, upgraded to Angular 21 · GitHub"
author: "isms-core-adm"
captured_at: "2026-08-12T21:35:52Z"
capture_tool: "hn-digest"
hn_id: 49278224
score: 1
comments: 0
posted_at: "2026-08-12T20:39:02Z"
tags:
  - hacker-news
  - translated
---

# Show HN: We revived the archived Kubernetes Dashboard (Angular 16 → 22)

- HN: [49278224](https://news.ycombinator.com/item?id=49278224)
- Source: [github.com](https://github.com/isms-core-project/kubernetes-dashboard)
- Score: 1
- Comments: 0
- Posted: 2026-08-12T20:39:02Z

## Translation

タイトル: HN を表示: アーカイブされた Kubernetes ダッシュボードを復活させました (Angular 16 → 22)
記事のタイトル: GitHub - isms-core-project/kubernetes-dashboard: オリジナルの Kubernetes ダッシュボード — アップストリームは廃止され、Angular 21 にアップグレード · GitHub
説明: 元の Kubernetes ダッシュボード — アップストリームは廃止され、Angular 21 にアップグレードされました - isms-core-project/kubernetes-dashboard

記事本文:
GitHub - isms-core-project/kubernetes-dashboard: オリジナルの Kubernetes ダッシュボード — アップストリームは廃止され、Angular 21 にアップグレードされました · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ismsコアプロジェクト
/
Kubernetes-ダッシュボード
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
69 コミット 69 コミット .github/ ISSUE_TEMPLATE .github/ ISSUE_TEMPLATE マニフェスト マニフェスト スクリーンショット スクリーンショット .gitignore .gitignore ANGULAR-UPGRA

DE.md ANGULAR-UPGRADE.md DEPLOYMENT.md DEPLOYMENT.md ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
元の kubernetes/ダッシュボードは、次のメッセージとともに 2026 年 1 月にアーカイブされました。
「このプロジェクトは現在アーカイブされており、アクティブなメンテナーや貢献者が不足しているため、維持されなくなりました。このプロジェクトを使用したり、スターを付けたり、貢献したすべての皆様に感謝します。ご自身で開発を続けたい場合は、お気軽にこのリポジトリをフォークしてください。代わりに Headlamp の使用を検討してください。」
Go API バックエンドは堅牢であり、維持する価値がありました。 Angular WebUI は Angular 16 でした。アーカイブ時点ではすでにメジャー バージョンが 1 つ遅れており、毎月さらにずれています。私たちはそれを腐らせるのではなく、正しい方法でフォークしてアップグレードしました。つまり、すべての Angular メジャー バージョンを一度に 1 つずつステップ実行 (16 → 17 → 18 → 19 → 20 → 21) して、途中でカタログ化された 44 個の重大な変更をすべて修正しました。詳細については、ANGULAR-UPGRADE.md を参照してください。それからさらに一歩進んで、Angular 22 になりました。
共有 Go イメージ (dashboard-api 、dashboard-auth 、dashboard-metrics-scraper 、および Kong 3.9.1) を使用します。すべてのイメージは ghcr.io/isms-core-project/kubernetes-dashboard から取得されます。
クラスターの健全性の概要 — 統計タイル、ドーナツ チャート、ライブ ネットワーク トラフィック グラフ (Grafana Alloy)
ワークロード — デプロイメント、DaemonSet、StatefulSet、ジョブ、Cron ジョブ、レプリカ セットの完全なリストと詳細
ワークロード アクション — YAML の編集、再起動、スケール、ロールバック、一時停止/再開、シェルの実行 — すべて RBAC 対応
クラスター マップ — ヘルス フィルターとズームを備えた名前空間スコープのトポロジ ビュー
トポロジ — Pod、ReplicaSet、Deployment、StatefulSet、DaemonSet、Services、および Ingresses の強制的な依存関係グラフ。種類と健康状態ごとに色分けされており、検索可能で、クリックするとリソースの詳細ページに移動します
応用プロジェクト

— ポッドの健全性とリソースの合計を含む名前空間ごとのカード
ポリシー監査 — ワークロードごとの Polaris ネイティブのセキュリティ スコアリング (0 ～ 100)
リソース効率 — Goldilocks スタイルの CPU/メモリ要求と制限と実際。 VictoriaMetrics 経由のトレンド矢印
RBAC Viewer — ワイルドカード検出を備えたクラスター全体のロール バインディング テーブル
証明書トラッカー — crypto/x509 で解析された TLS シークレット。有効期限のカウントダウンとステータスバッジ
イベント タイムライン — タイムバケットのグループ化と警告のハイライトを備えたライブ イベント フィード
レジストリ マネージャー — ポッド imagePullSecrets と相互参照される docker pull シークレット
履歴メトリクス — 1h/6h/24h/7d セレクターを備えたポッド CPU/メモリ スパークライン (VictoriaMetrics または Prometheus)
イベント アラート — CrashLoop/OOM/ImagePullBackOff/NodeNotReady に関するリアルタイムの電子メール。タイプごとに構成可能
ポッド ログ — ライブ ストリーミング、タイムスタンプ、重大度フィルター、テキスト フィルター、ダウンロード
クラスター シェル — ダッシュボード ポッドに実行される対話型の xterm.js ターミナル。 kubectl はユーザーの JWT として実行されます
AI アシスタント — SSE ストリーミング経由の Claude Sonnet。詳細ページから自動挿入されるポッド仕様とイベント
統計タイル、ドーナツ チャート、ネットワーク トラフィック グラフ。
ステータス、再起動回数、インラインアクションを含む完全なワークロードリスト。
ヘルス フィルターとズームを備えた名前空間スコープのトポロジ。
強制的な依存関係グラフ - 種類と健全性によって色分けされた、所有権とサービス/イングレス エッジ。
ライブ CPU/メモリ スパークライン、再起動回数、ノード割り当て。
ノードごとの CPU とメモリのリクエストの割合とポッドの容量。
ワークロードごとの Polaris セキュリティ スコア。
Goldilocks スタイルの CPU/メモリとトレンド矢印の比較。
Docker プル シークレットは pod imagePullSecrets と相互参照されます。
cert-manager 証明書、発行者、および ClusterIssuers — 自動検出。
crypto/x509 でスキャンされた TLS シークレット — 有効期限のカウントダウンとステータス バッジ。

IP アドレス プールと L2 アドバタイズメント — MetalLB CRD が存在する場合に自動検出されます。
ポッドのセキュリティ/ネットワーク ポリシー
ポッドのセキュリティ標準とネットワークポリシーの視覚化。
タイムバケットグループ化を備えたライブイベントフィード。
コンプライアンス スコアと CVE 結果 — 自動検出。
1h/6h/24h/7d セレクターを備えたポッド CPU/メモリ スパークライン。
kubelet stats API からのライブ使用量バー。
解決されたルールとワイルドカード検出を備えたすべてのロール バインディング。
完全な対話型 bash ターミナル。
Kong API ゲートウェイが前面にある、ダッシュボード名前空間内の 5 つのポッド:
ブラウザ
━── Kong 3.9.1 (DBless、NodePort :30080)
§── /api/v1/login、/csrftoken、/me → ダッシュボード認証
§── /api/* → ダッシュボード API
│ └── サイドカー: ダッシュボード-メトリクス-スクレーパー
━━ / → ダッシュボードウェブ（SPA）
オプションのアドオン (すべて同じ名前空間内):
完全な Runbook については、DEPLOYMENT.md を参照してください。
Copyright 2017 Kubernetes 作者
Copyright 2026 ISMSコアプロジェクト
Apache License、バージョン 2.0 に基づいてライセンスされています。全文についてはライセンスを参照してください。
元の Kubernetes ダッシュボード — アップストリームは廃止され、Angular 21 にアップグレードされました
kubernetes-dashboard.com トピック
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Original Kubernetes Dashboard — retired upstream, upgraded to Angular 21 - isms-core-project/kubernetes-dashboard

GitHub - isms-core-project/kubernetes-dashboard: Original Kubernetes Dashboard — retired upstream, upgraded to Angular 21 · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
isms-core-project
/
kubernetes-dashboard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
69 Commits 69 Commits .github/ ISSUE_TEMPLATE .github/ ISSUE_TEMPLATE manifests manifests screenshots screenshots .gitignore .gitignore ANGULAR-UPGRADE.md ANGULAR-UPGRADE.md DEPLOYMENT.md DEPLOYMENT.md LICENSE LICENSE README.md README.md View all files Repository files navigation
The original kubernetes/dashboard was archived in January 2026 with this message:
"This project is now archived and no longer maintained due to lack of active maintainers and contributors. Thank you to everyone who used, starred, or contributed to this project! Feel free to fork this repository if you want to continue development yourself. Please consider using Headlamp instead."
The Go API backend was solid and worth keeping. The Angular WebUI was Angular 16 — already one major version behind at archive time, and drifting further every month. Rather than let it rot, we forked it and upgraded it the right way: stepping through every Angular major version one at a time — 16 → 17 → 18 → 19 → 20 → 21 — fixing all 44 catalogued breaking changes along the way. See ANGULAR-UPGRADE.md for the full story. It's since gone one step further, to Angular 22.
Uses the shared Go images: dashboard-api , dashboard-auth , dashboard-metrics-scraper , and Kong 3.9.1. All images pull from ghcr.io/isms-core-project/kubernetes-dashboard .
Cluster Health Overview — stat tiles, donut charts, live Network Traffic graph (Grafana Alloy)
Workloads — full list + detail for Deployments, DaemonSets, StatefulSets, Jobs, Cron Jobs, Replica Sets
Workload Actions — edit YAML, restart, scale, rollback, pause/resume, exec shell — all RBAC-aware
Cluster Map — namespace-scoped topology view with health filter and zoom
Topology — force-directed dependency graph of Pods, ReplicaSets, Deployments, StatefulSets, DaemonSets, Services, and Ingresses; colour-coded by kind and health, searchable, click-through to any resource's detail page
Application Projects — per-namespace cards with pod health and resource totals
Policy Audit — Polaris-native security scoring (0–100) per workload
Resource Efficiency — Goldilocks-style CPU/memory request vs limit vs actual; trend arrows via VictoriaMetrics
RBAC Viewer — cluster-wide role binding table with wildcard detection
Certificate Tracker — TLS secrets parsed with crypto/x509 ; expiry countdown and status badges
Event Timeline — live event feed with time-bucket grouping and warning highlight
Registry Manager — docker pull secrets cross-referenced with pod imagePullSecrets
Historical Metrics — pod CPU/memory sparklines with 1h/6h/24h/7d selector (VictoriaMetrics or Prometheus)
Event Alerts — real-time email on CrashLoop/OOM/ImagePullBackOff/NodeNotReady; configurable per type
Pod Logs — live streaming, timestamps, severity filter, text filter, download
Cluster Shell — interactive xterm.js terminal exec'd into the dashboard pod; kubectl runs as the user's JWT
AI Assistant — Claude Sonnet via SSE streaming; pod spec and events auto-injected from detail pages
Stat tiles, donut charts, and Network Traffic graph.
Full workload list with status, restart count, and inline actions.
Namespace-scoped topology with health filter and zoom.
Force-directed dependency graph — ownership and Service/Ingress edges, colour-coded by kind and health.
Live CPU/Memory sparklines, restart count, node assignment.
Per-node CPU and memory request percentages and pod capacity.
Polaris security scoring per workload.
Goldilocks-style CPU/memory comparison with trend arrows.
Docker pull secrets cross-referenced with pod imagePullSecrets .
cert-manager Certificates, Issuers, and ClusterIssuers — auto-detected.
TLS secrets scanned with crypto/x509 — expiry countdown and status badges.
IP Address Pools and L2 Advertisements — auto-detected when MetalLB CRDs are present.
Pod Security / Network Policies
Pod Security Standards and NetworkPolicy visualisation.
Live event feed with time-bucket grouping.
Compliance scores and CVE findings — auto-detected.
Pod CPU/memory sparklines with 1h/6h/24h/7d selector.
Live usage bars from the kubelet stats API.
All role bindings with resolved rules and wildcard detection.
Full interactive bash terminal.
Five pods in the dashboard namespace, fronted by a Kong API gateway:
Browser
└── Kong 3.9.1 (DBless, NodePort :30080)
├── /api/v1/login, /csrftoken, /me → dashboard-auth
├── /api/* → dashboard-api
│ └── sidecar: dashboard-metrics-scraper
└── / → dashboard-web (SPA)
Optional add-ons (all in the same namespace):
See DEPLOYMENT.md for the full runbook.
Copyright 2017 The Kubernetes Authors
Copyright 2026 The ISMS Core Project
Licensed under the Apache License, Version 2.0. See LICENSE for the full text.
Original Kubernetes Dashboard — retired upstream, upgraded to Angular 21
kubernetes-dashboard.com Topics
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
