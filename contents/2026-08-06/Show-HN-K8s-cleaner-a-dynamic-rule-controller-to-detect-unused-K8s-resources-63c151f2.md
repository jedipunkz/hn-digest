---
source: "https://github.com/gianlucam76/k8s-cleaner"
hn_url: "https://news.ycombinator.com/item?id=49195764"
title: "Show HN: K8s-cleaner a dynamic rule controller to detect unused K8s resources"
article_title: "GitHub - gianlucam76/k8s-cleaner: Cleaner is a Kubernetes controller that identifies unused or unhealthy resources, helping you maintain a streamlined and efficient Kubernetes cluster. It provides flexible scheduling, label filtering, Lua-based selection criteria, resource removal or update and noti\n[truncated]"
author: "mgianluc76"
captured_at: "2026-08-06T12:51:14Z"
capture_tool: "hn-digest"
hn_id: 49195764
score: 1
comments: 0
posted_at: "2026-08-06T12:30:50Z"
tags:
  - hacker-news
  - translated
---

# Show HN: K8s-cleaner a dynamic rule controller to detect unused K8s resources

- HN: [49195764](https://news.ycombinator.com/item?id=49195764)
- Source: [github.com](https://github.com/gianlucam76/k8s-cleaner)
- Score: 1
- Comments: 0
- Posted: 2026-08-06T12:30:50Z

## Translation

タイトル: HN を表示: 未使用の K8s リソースを検出する動的ルール コントローラー K8s-cleaner
記事のタイトル: GitHub - gianlucam76/k8s-cleaner: Cleaner は、未使用または異常なリソースを特定し、合理化された効率的な Kubernetes クラスターの維持に役立つ Kubernetes コントローラーです。柔軟なスケジューリング、ラベル フィルタリング、Lua ベースの選択基準、リソースの削除または更新、および通知を提供します。
[切り捨てられた]
説明: Cleaner は、未使用または異常なリソースを特定し、合理化された効率的な Kubernetes クラスターの維持を支援する Kubernetes コントローラーです。柔軟なスケジュール、ラベル フィルタリング、Lua ベースの選択基準、リソースの削除または更新、Slack、Webex、Dis 経由の通知を提供します。
[切り捨てられた]

記事本文:
GitHub - gianlucam76/k8s-cleaner: Cleaner は、未使用または異常なリソースを特定し、合理化された効率的な Kubernetes クラスターの維持に役立つ Kubernetes コントローラーです。柔軟なスケジュール設定、ラベル フィルタリング、Lua ベースの選択基準、リソースの削除または更新、Slack、Webex、Discord 経由の通知を提供します。クラスターの操作を自動化することもできます。 · GitHub
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
ジャンルカム76
/
k8s-クリーナー
公共
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
あなたはむ

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1,064 コミット 1,064 コミット .github .github api/ v1alpha1 api/ v1alpha1 資産 アセット チャート/ k8s-cleaner charts/ k8s-cleaner cmd cmd config config docs docs 例-自動化-操作/ スケジュールされたスケーリング 例-自動化-操作/ スケジュールされたスケーリング 例-異常なリソース例-不健全なリソース 例-未使用のリソース 例-未使用のリソース ハック ハック 内部 内部マニフェスト pkg/ スコープ pkg/ スコープ テスト テスト Web Web .all-contributorsrc .all-contributorsrc .gitignore .gitignore .golangci.yaml .golangci.yaml ADOPTERS.md ADOPTERS.md COTRIBUTING.md COTRIBUTING.md DEVELOPMENT.md DEVELOPMENT.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile PROJECT PROJECT README.md README.md go.mod go.mod go.sum go.sum mkdocs.yml mkdocs.yml renovate.json renovate.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Kubernetes コントローラー Cleaner は、古いリソース、孤立したリソース、または異常なリソースを特定、削除、または更新して、クリーンで効率的な Kubernetes クラスターを維持します。これは、あらゆる Kubernetes リソース タイプ (独自のカスタム リソースを含む) を処理できるように設計されており、ラベル ベースの選択やカスタム Lua ベースの基準などの高度なフィルタリング機能を提供します。
クリーナーは、異常なリソースを特定するためにも使用できます。
k8s-cleaner は、次のような便利な通知で常に最新情報を入手できます。
各通知には、k8s-cleaner によって正常に削除 (または変更) されたすべてのリソースのリストが含まれています。あなたにとって最適なものをお選びください!
現在、k8s-cleaner には、未使用のものを識別してリストするための豊富な実例セットがあります。
Persistent Volumes / Persistent VolumeClaims
注釈の指示に基づいてリソースを識別する

最長寿命または有効期限
異常なリソースを特定する例もあります。
古いコンテンツを含むシークレットをマウントするポッド : 最新のシークレット データを利用していないポッドを検出します。
期限切れの証明書を使用しているポッド : 期限切れのセキュリティ証明書を使用して動作しているポッドを特定します。
リソースの状態とライブ Prometheus メトリクスを組み合わせた例:
高エラー率でデプロイメントをスケールダウンする : Prometheus によって報告された HTTP エラー率がしきい値を超えた場合、デプロイメントをゼロにスケールダウンします。
メモリが飽和したポッドを再起動する : メモリのワーキング セットが設定された制限の 90% を超えた場合にポッドを削除します。
1️⃣ スケジュール : k8s-cleaner がクラスターをスキャンして古いリソースを特定する頻度を指定します。 Cron 構文を使用して、定期的なスケジュールを定義します。
2️⃣ DryRun : 実際のリソース構成に影響を与えることなく、k8s-cleaner フィルタリング ロジックの安全なテストを有効にします。基準に一致するリソースは識別されますが、変更は適用されません。
3️⃣ ラベル フィルタリング : ユーザー定義のラベルに基づいてリソースを選択し、不要なコンポーネントや古いコンポーネントを除外します。ラベル キー、操作 (等しい、異なるなど)、および値に基づいて選択を絞り込みます。
4️⃣ Lua ベースの選択基準 : Lua スクリプト言語を活用して、特定のリソース管理のニーズに応える複雑で動的な選択基準を作成します。カスタム ロジックを定義して、古いリソースを識別して処理します。 Cleaner の構成を検証するには、ここを参照してください。
5️⃣ 通知 : 最新情報を入手してください! k8s-cleaner は、削除されたか最適化されたかにかかわらず、クリーンアップされたすべてのリソースについてユーザーに常に知らせます。詳細な通知リストを取得し、好みのチャネル (Slack、Webex、Discord、Teams、Telegram、SMTP、レポート) を選択します。
6️⃣ ウェブダッシュボード

d : クラスターの健全性を視覚化します。オプションの埋め込みダッシュボードを使用すると、応答性の高い UI 経由でスキャン結果を参照し、Lua スクリプトを検査し、ワンクリックでオンデマンドのクリーンアップ タスクをトリガーできます。 Helm 経由で簡単に有効にすることができ、ダーク モードと読み取り専用構成の両方をサポートします。
7️⃣ メトリックベースの選択 : 各リソースを評価する前に、Prometheus 互換のエンドポイントをクエリします。結果は Lua スクリプトのグローバル メトリクス テーブルとして公開されるため、ライブ メトリクス値に基づいてリソースの一致を制御できます。たとえば、HTTP エラー率が 5% を超えた場合にのみデプロイメントをスケールダウンしたり、メモリの飽和が Prometheus によって報告された場合にのみポッドを再起動したりできます。
8️⃣ ロールバック : 削除または変換アクションの直前にすべてのリソースの状態をキャプチャするため、Persistent Volume を必要とせず、単一の API 呼び出しで最新の実行を元に戻すことができます。
機能の完全なリストと例については、リンクを参照してください。
💪 リソースの削除 : クラスターから古いリソースを効率的に削除し、未使用のリソースを再利用してリソースの使用率を向上させます。
💪 リソースの更新 : 古いリソースを更新して、最新の構成に合わせ、一貫した機能を維持します。
💪 リソースの膨張の削減 : リソースの膨張を最小限に抑え、クリーンで組織化されたクラスターを維持し、全体的なパフォーマンスと安定性を向上させます。
k8s-cleaner は、スケジュールの柔軟性、ラベル フィルタリングの精度、Lua ベースの基準の能力、および古いリソースを削除または更新する機能を組み合わせることで、ユーザーが Kubernetes 環境を効果的に管理し、リソースの使用状況を最適化できるようにします。
👉 機能リクエストやバグについては、問題を提出してください。
👉 更新情報を入手するには ⭐️ このリポジトリにスターを付けます。
👉 実際の例は例で見つけることができます

セクション。
Sveltos を使用して複数のクラスターにインストールする
Kubernetes クラスターのフリートを管理する場合、Sveltos を使用すると、インフラストラクチャ全体にわたる k8s-cleaner のデプロイと管理が簡素化されます。 k8s-cleaner を各クラスターに手動でデプロイする代わりに、Sveltos は次のことを行う集中プラットフォームを提供します。
デプロイの自動化: 単一の構成で、k8s-cleaner を複数のクラスターに簡単にデプロイします。
構成の管理 : k8s-cleaner 構成を一元管理し、すべてのクラスターに一貫して適用します。
一貫性の確保 : フリート全体で一貫した k8s-cleaner の構成とバージョンを維持します。
詳細については、こちらをご覧ください。
クリーナー構成の検証
Cleaner の設定が正しいことを確認するには、ドキュメント (ここ および ここ ) に記載されている包括的な手順に従ってください。
基本的に、一致するリソースと一致しないリソースを表す YAML ファイルと一緒に Cleaner YAML ファイルを提供する必要があり、その後、単純な make ut コマンドを実行することで結果が表示されます。これにより、構成が目的のリソースを正しく識別し、管理しているかどうかが検証されます。
新しいクリーナー構成を追加して、サンプル ディレクトリに貢献することをお勧めします 💡。これにより、コミュニティはさまざまな専門知識から恩恵を受け、クリーナーのユースケースに関する強力な知識ベースを構築することができます。
サンプルを追加するには、わかりやすい名前を付けてサンプル ディレクトリに新しいファイルを作成し、そのファイル内に Cleaner 設定を配置するだけです。サンプルを追加したら、お気軽にプル リクエストを送信してコミュニティと共有してください。
🤝 私たちは協力して Cleaner アプリケーションの範囲を拡大し、Kubernetes リソースを効率的に管理するためのさらに価値のあるツールにすることができます。
このプロジェクトは CNCF 行動規範に準拠しています

参加することにより、この規範を遵守することが期待されます。
ジャンルカ・マルデンテ
💻
オリバー・ベーラー
💻
エレニ・グロスドゥーリ
📖
コリン・J・レイシー
💻
アミン・モハマディアン
📖
ヴラド・モカヌ
🤔 📝 💻
について
Cleaner は、未使用または異常なリソースを識別する Kubernetes コントローラーで、合理化された効率的な Kubernetes クラスターの維持に役立ちます。柔軟なスケジュール設定、ラベル フィルタリング、Lua ベースの選択基準、リソースの削除または更新、Slack、Webex、Discord 経由の通知を提供します。クラスターの操作を自動化することもできます。
プロジェクトveltos.github.io/sveltos/ トピック
Readme Apache-2.0 ライセンス
48 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Cleaner is a Kubernetes controller that identifies unused or unhealthy resources, helping you maintain a streamlined and efficient Kubernetes cluster. It provides flexible scheduling, label filtering, Lua-based selection criteria, resource removal or update and notifications via Slack, Webex and Dis
[truncated]

GitHub - gianlucam76/k8s-cleaner: Cleaner is a Kubernetes controller that identifies unused or unhealthy resources, helping you maintain a streamlined and efficient Kubernetes cluster. It provides flexible scheduling, label filtering, Lua-based selection criteria, resource removal or update and notifications via Slack, Webex and Discord. it can also automate clusters operations. · GitHub
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
gianlucam76
/
k8s-cleaner
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1,064 Commits 1,064 Commits .github .github api/ v1alpha1 api/ v1alpha1 assets assets charts/ k8s-cleaner charts/ k8s-cleaner cmd cmd config config docs docs examples-automated-operations/ scheduled-scaling examples-automated-operations/ scheduled-scaling examples-unhealthy-resources examples-unhealthy-resources examples-unused-resources examples-unused-resources hack hack internal internal manifest manifest pkg/ scope pkg/ scope test test web web .all-contributorsrc .all-contributorsrc .gitignore .gitignore .golangci.yaml .golangci.yaml ADOPTERS.md ADOPTERS.md CONTRIBUTING.md CONTRIBUTING.md DEVELOPMENT.md DEVELOPMENT.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile PROJECT PROJECT README.md README.md go.mod go.mod go.sum go.sum mkdocs.yml mkdocs.yml renovate.json renovate.json View all files Repository files navigation
The Kubernetes controller Cleaner identifies, removes, or updates stale/orphaned or unhealthy resources to maintain a clean and efficient Kubernetes cluster. It is designed to handle any Kubernetes resource types (including your own custom resources) and provides sophisticated filtering capabilities, including label-based selection and custom Lua-based criteria.
Cleaner can also be used to identify unhealthy resources.
k8s-cleaner keeps you in the loop with handy notifications through:
Each notification contains list of all resources successfully deleted (or modified) by k8s-cleaner. Choose what works best for you!
Currently k8s-cleaner has rich set of working examples to identify and list unused :
PersistentVolumes / PersistentVolumeClaims
Identify resources based on annotation indicating the maximum lifespan or the expiration date
There are also examples to identify unhealthy resources:
Pods Mounting Secrets with Old Content : Detect pods that are not utilizing the most recent Secret data.
Pods Using Expired Certificates : Pinpoint pods that are operating with expired security certificates.
And examples that combine resource state with live Prometheus metrics :
Scale Down Deployments on High Error Rate : Scale Deployments to zero when the HTTP error rate reported by Prometheus exceeds a threshold.
Restart Memory-Saturated Pods : Delete Pods when their memory working set exceeds 90% of the configured limit.
1️⃣ Schedule : Specify the frequency at which the k8s-cleaner should scan the cluster and identify stale resources. Utilise the Cron syntax to define recurring schedules.
2️⃣ DryRun : Enable safe testing of the k8s-cleaner filtering logic without affecting actual resource configurations. Resources matching the criteria will get identified, but no changes will get applied.
3️⃣ Label Filtering : Select resources based on user-defined labels, filtering out unwanted or outdated components. Refine the selection based on label key, operation (equal, different, etc.), and value.
4️⃣ Lua-based Selection Criteria : Leverage the Lua scripting language to create complex and dynamic selection criteria, catering to specific resource management needs. Define custom logic to identify and handle stale resources. To validate the Cleaner configuration, have a look here .
5️⃣ Notifications : Stay informed! The k8s-cleaner keeps users in the loop about every cleaned-up resource, whether removed or optimized. Get detailed notification lists and pick your preferred channel: Slack, Webex, Discord, Teams, Telegram, SMTP or reports.
6️⃣ Web Dashboard : Visualize your cluster's health! Use the optional embedded dashboard to browse scan results via a responsive UI, inspect your Lua scripts, and trigger on-demand cleanup tasks with a single click. It can be easily enabled via Helm and supports both dark mode and read-only configurations.
7️⃣ Metric-based Selection : Query any Prometheus-compatible endpoint before evaluating each resource. The results are exposed as a global metrics table in the Lua script, so you can gate resource matching on live metric values — for example, scale down Deployments only when their HTTP error rate exceeds 5%, or restart Pods only when memory saturation is reported by Prometheus.
8️⃣ Rollback : Capture the state of every resource right before a Delete or Transform action, so the most recent execution can be reverted with a single API call, no PersistentVolume required.
For a complete list of features with examples , have a look at the link .
💪 Resource Removal : Efficiently remove stale resources from your cluster, reclaiming unused resources and improving resource utilisation.
💪 Resource Updates : Update outdated resources to ensure they align with the latest configurations and maintain consistent functionality.
💪 Reduced Resource Bloat : Minimize resource bloat and maintain a clean and organized cluster, improving overall performance and stability.
By combining the flexibility of scheduling , the accuracy of label filtering , the power of Lua-based criteria , and the ability to remove or update stale resources, the k8s-cleaner empowers users to effectively manage Kubernetes environments and optimise resource usage.
👉 For feature requests and bugs, file an issue .
👉 To get updates ⭐️ star this repository.
👉 Working examples can be found in the examples section.
Install on Multiple Clusters with Sveltos
If you manage a fleet of Kubernetes clusters, Sveltos simplifies the deployment and management of k8s-cleaner across your entire infrastructure. Instead of manually deploying k8s-cleaner to each cluster, Sveltos offers a centralized platform to:
Automate Deployment : Easily deploy k8s-cleaner to multiple clusters with a single configuration.
Manage Configurations : Centrally manage k8s-cleaner configurations and apply them consistently across all clusters.
Ensure Consistency : Maintain consistent k8s-cleaner configurations and versions across your fleet.
Detailed information can be found here .
Validate Cleaner Configuration
To verify the correctness of the Cleaner configuration, follow the comprehensive instructions provided in the documentation: here and here .
In essence, the Cleaner YAML file alongside the YAML files representing matching and non-matching resources need to get provided, and then by executing the simple make ut command the resutls will appear. This will validate whether your configuration correctly identifies and manages the desired resources.
We encourage everyone to contribute to the example directory by adding new Cleaner configurations 💡. This will help the community benefit from different expertise and build a stronger knowledge base of the Cleaner use cases.
To add an example, simply create a new file in the example directory with a descriptive name and put your Cleaner configuration within the file. Once you've added your example, feel free to submit a pull request to share it with the community.
🤝 Together we can expand the range of Cleaner applications and make it an even more valuable tool for managing Kubernetes resources efficiently.
This project adheres to the CNCF Code of Conduct
By participating, you are expected to honor this code.
Gianluca Mardente
💻
Oliver Bähler
💻
Eleni Grosdouli
📖
Colin J Lacy
💻
Amin Mohammadian
📖
Vlad Mocanu
🤔 📝 💻
About
Cleaner is a Kubernetes controller that identifies unused or unhealthy resources, helping you maintain a streamlined and efficient Kubernetes cluster. It provides flexible scheduling, label filtering, Lua-based selection criteria, resource removal or update and notifications via Slack, Webex and Discord. it can also automate clusters operations.
projectsveltos.github.io/sveltos/ Topics
Readme Apache-2.0 license Contributing
48 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
