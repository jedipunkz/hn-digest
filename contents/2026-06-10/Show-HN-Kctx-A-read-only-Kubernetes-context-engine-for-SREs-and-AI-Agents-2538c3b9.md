---
source: "https://github.com/lucasepe/kctx"
hn_url: "https://news.ycombinator.com/item?id=48482439"
title: "Show HN: Kctx – A read-only Kubernetes context engine for SREs and AI Agents"
article_title: "GitHub - lucasepe/kctx: A Kubernetes context engine for humans and AI agents. · GitHub"
author: "lucasepe"
captured_at: "2026-06-10T21:44:40Z"
capture_tool: "hn-digest"
hn_id: 48482439
score: 3
comments: 0
posted_at: "2026-06-10T20:44:52Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Kctx – A read-only Kubernetes context engine for SREs and AI Agents

- HN: [48482439](https://news.ycombinator.com/item?id=48482439)
- Source: [github.com](https://github.com/lucasepe/kctx)
- Score: 3
- Comments: 0
- Posted: 2026-06-10T20:44:52Z

## Translation

タイトル: Show HN: Kctx – SRE および AI エージェント用の読み取り専用 Kubernetes コンテキスト エンジン
記事のタイトル: GitHub - lucasepe/kctx: 人間と AI エージェントのための Kubernetes コンテキスト エンジン。 · GitHub
説明: 人間と AI エージェント用の Kubernetes コンテキスト エンジン。 - ルカペ/kctx

記事本文:
GitHub - lucasepe/kctx: 人間と AI エージェント用の Kubernetes コンテキスト エンジン。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
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
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ルカペ
/
kctx
公共
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります

イングス
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
9 コミット 9 コミット .github .github chart chart docs docs external 内部スキーマ/ kctx.io/ v1alpha1 スキーマ/ kctx.io/ v1alpha1 スクリプト スクリプト testdata/ e2e/ ゴールデン testdata/ e2e/ ゴールデン .dockerignore .dockerignore .gitignore .gitignore Dockerfile Dockerfile LICENSE.txt LICENSE.txt METRICS.md METRICS.md README.md README.md ROADMAP.md ROADMAP.md go.mod go.mod go.sum go.sum install.sh install.sh main.go main.go すべてのファイルを表示 リポジトリ ファイルのナビゲーション
人間、スクリプト、AI エージェント用の小さな読み取り専用 Kubernetes コンテキスト エンジン。
kctx は、Kubernetes API の状態を構造化された操作コンテキストに変換します。
すべてのオペレーター、スクリプト、ダッシュボード、またはエージェントに生の YAML から関係を再構築するよう依頼する代わりに、kctx は次のコンパクトなモデルを公開します。
エンティティ : ポッド、サービス、ワークロード、ノード、PVC、ConfigMap、シークレット、CRD
関係: 所有権、選択、スケジューリング、サービスルーティング、依存関係
シグナル : 異常な Pod、欠落しているエンドポイント、
警告イベント、準備の失敗、またはワークロードの低下
グラフ: サポートされているリソースに関する依存関係と所有権のビュー
ダンプ: マシンの決定論的な名前空間スナップショットとインシデント レビュー
このツールは意図的に保守的であり、クラスターの状態を読み取り、事実を正規化し、推測的な根本原因の主張を回避します。
kctx の背後にある動機、哲学、設計上の議論については、ここで詳しく説明します。
Leanpub の書籍: Kubernetes コンテキスト エンジニアリング
次のような質問に答えたい場合は、kctx を使用します。
どのサービスがこれらのバックエンドにルーティングしますか?
この名前空間が異常に見えるのはなぜですか?
インシデントにはどのようなリソースとシグナルを添付する必要がありますか?
AI エージェントが事前にどのコンパクトな Kubernetes コンテキストを受け取る必要があるか

鉱石の推理？
このサポートされている CRD は現時点で運用上何を意味しますか?
kctx は、SRE、プラットフォーム チーム、Kubernetes オペレーター、CI/CD に役立ちます
診断、内部ツール、MCP ツール、AI SRE ワークフロー。
リソースの変更、ワークロードの再起動、マニフェストの適用、根本原因の推測は行いません。
kctx は、生のマニフェスト、シークレット データ、ConfigMap データ、生の環境変数、ログ、およびワークロード メトリックを回避します。サポートされている出力によって返されるメタデータおよび Kubernetes メッセージは、共通の秘密を含むキーとテキスト パターンに対する小さな編集ポリシーを通過します。
./install.sh
または、公開されたリリース スクリプトからインストールします。
カール -fsSL https://raw.githubusercontent.com/lucasepe/kctx/main/install.sh |バッシュ
次に、現在の Kubernetes コンテキストに対して実行します。
kctx health 名前空間のデフォルト
kctx Explain pod < ポッド名 > --namespace デフォルト
kctx トレース サービス < サービス名 > --namespace デフォルト
kctx グラフ ポッド < ポッド名 > --namespace デフォルト
kctx ダンプ名前空間のデフォルト
ローカル開発の場合:
走ってください。 health 名前空間のデフォルト
Helm を使用して kctx サーブをインストールする
パッケージ化されたリリース チャートからクラスター内の読み取り専用 HTTP サーバーをインストールします。
バージョン=0.2.0
Helm アップグレード --install kctx \
" https://github.com/lucasepe/kctx/releases/download/v ${VERSION} /kctx- ${VERSION} .tgz " \
--namespace kctx-system \
--create-namespace
次に、それをローカルに公開します。
kubectl -n kctx-system ポートフォワード svc/kctx 8080:8080
カール http://localhost:8080/health/namespace/default
ソース チェックアウトから、ローカル チャートをインストールし、実行するイメージ タグを選択します。
Helm upgrade --install kctx ./chart \
--namespace kctx-system \
--create-namespace \
--set image.tag=dev
チャートの値、ローカル種類の設定、および NodePort の例については、chart/README.md を参照してください。
1 つのリソースに関する構造化されたコンテキストを解決します。ネイティブ Pod コンテキストがサポートされており、レジ

stered CRD アダプターは、エコシステム固有のコンテキストを提供できます。
kctx Explain pod api-xyz --namespacepayments
kctx Explain application.argoproj.io ゲストブック --namespace argocd
kctxグラフ
サポートされているリソースを中心にグラフを構築します。 JSON がデフォルトの出力です。 Mermaid および DOT レンダラーは、グラフ指向のビューに使用できます。
kctx グラフ ポッド api-xyz --namespacepayments
kctx グラフ ポッド api-xyz --namespacepayments --render mermaid
kctx グラフ application.argoproj.io guestbook --namespace argocd --render dot
kctx トレース サービス
エンドポイントスライス、エンドポイント、ポッド、所有者、ノード、および実際のサービス正常性シグナルまでサービスをトレースします。
kctx トレース サービスpayments-api --namespacepayments
kctx health 名前空間
コンパクトな名前空間の健全性スナップショットを作成します。
kctx health 名前空間の支払い
kctx ダンプ名前空間
自動化、インシデントレビュー、または AI エージェントのグラウンディングのために、決定論的な名前空間コンテキストのスナップショットをエクスポートします。
kctx ダンプ名前空間payments >payments-dump.json
kctxサーブ
軽量の読み取り専用 HTTP API を通じて同じコンテキスト エンジンを公開します。
kctxサーブ
カール http://localhost:8080/health/namespace/default
CRDアダプター
kctx は、ディスカバリーを通じて任意の Kubernetes リソースをフェッチできますが、すべてのカスタム リソースを一般的に理解できるとは限りません。
エコシステム固有のセマンティクスは明示的なアダプターに属します。アダプターは、CRD を、リソース ID、コンパクト ステータス、関連エンティティ、リレーション、シグナル、およびオプションでグラフ ノードとエッジなど、他の場所で使用されるのと同じコア コントラクトに変換できます。
現在のアダプター セットは、 Argo CD Application 、 Argo CD AppProject 、および cert-manager Certificate リソースをカバーしています。
長い形式のドキュメントは PDF eBook として構成されています。
kctx によって Kubernetes ワークロードのデバッグ時間を節約できる場合は、kctx のサポートを検討してください。
そのメンテナンス:
GitHub スポンサー: https://g

ithub.com/sponsors/lucasepe
ペイパル: https://paypal.me/lucasepe71
サポートは、リリース作業、互換性テスト、文書化、および
継続的な Kubernetes、Argo CD、および証明書マネージャー統合メンテナンス。
kctx は鋭意開発中です。これは読み取り専用の Kubernetes コンテキスト ツールとしてすでに便利ですが、運用環境の強化はまだ進行中です。
現在の運用準備作業、計画されている機能、オープンな設計領域については、ROADMAP.md を参照してください。
人間と AI エージェントのための Kubernetes コンテキスト エンジン。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
3
v0.2.1
最新の
2026 年 6 月 8 日
+ 2 リリース
このプロジェクトのスポンサーになる
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A Kubernetes context engine for humans and AI agents. - lucasepe/kctx

GitHub - lucasepe/kctx: A Kubernetes context engine for humans and AI agents. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
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
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
lucasepe
/
kctx
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
9 Commits 9 Commits .github .github chart chart docs docs internal internal schemas/ kctx.io/ v1alpha1 schemas/ kctx.io/ v1alpha1 scripts scripts testdata/ e2e/ golden testdata/ e2e/ golden .dockerignore .dockerignore .gitignore .gitignore Dockerfile Dockerfile LICENSE.txt LICENSE.txt METRICS.md METRICS.md README.md README.md ROADMAP.md ROADMAP.md go.mod go.mod go.sum go.sum install.sh install.sh main.go main.go View all files Repository files navigation
A small read-only Kubernetes context engine for humans, scripts, and AI agents.
kctx turns Kubernetes API state into structured operational context.
Instead of asking every operator, script, dashboard, or agent to reconstruct relationships from raw YAML, kctx exposes a compact model of:
entities : Pods, Services, workloads, Nodes, PVCs, ConfigMaps, Secrets, CRDs
relations : ownership, selection, scheduling, service routing, dependencies
signals : factual observations such as unhealthy Pods, missing endpoints,
warning Events, failed readiness, or degraded workloads
graphs : dependency and ownership views around supported resources
dumps : deterministic namespace snapshots for machines and incident review
The tool is intentionally conservative: it reads cluster state, normalizes facts, and avoids speculative root-cause claims.
The motivation, philosophy, and design argument behind kctx are covered in longer form here:
Leanpub book: Kubernetes Context Engineering
Use kctx when you want to answer questions like:
Which Services route to these backends?
Why does this namespace look unhealthy?
What resources and signals should I attach to an incident?
What compact Kubernetes context should an AI agent receive before reasoning?
What does this supported CRD mean operationally right now?
kctx is useful for SREs, platform teams, Kubernetes operators, CI/CD
diagnostics, internal tooling, MCP tools, and AI SRE workflows.
It does not mutate resources, restart workloads, apply manifests, or guess root cause.
kctx avoids raw manifests, Secret data, ConfigMap data, raw environment variables, logs, and workload metrics. Metadata and Kubernetes messages that are returned by supported outputs pass through a small redaction policy for common secret-bearing keys and text patterns.
./install.sh
Or install from the published release script:
curl -fsSL https://raw.githubusercontent.com/lucasepe/kctx/main/install.sh | bash
Then run it against your current Kubernetes context:
kctx health namespace default
kctx explain pod < pod-name > --namespace default
kctx trace service < service-name > --namespace default
kctx graph pod < pod-name > --namespace default
kctx dump namespace default
For local development:
go run . health namespace default
Install kctx serve With Helm
Install the in-cluster read-only HTTP server from a packaged release chart:
VERSION=0.2.0
helm upgrade --install kctx \
" https://github.com/lucasepe/kctx/releases/download/v ${VERSION} /kctx- ${VERSION} .tgz " \
--namespace kctx-system \
--create-namespace
Then expose it locally:
kubectl -n kctx-system port-forward svc/kctx 8080:8080
curl http://localhost:8080/health/namespace/default
From a source checkout, install the local chart and choose the image tag to run:
helm upgrade --install kctx ./chart \
--namespace kctx-system \
--create-namespace \
--set image.tag=dev
See chart/README.md for chart values, local kind setup, and NodePort examples.
Resolve structured context around one resource. Native Pod context is supported, and registered CRD adapters can provide ecosystem-specific context.
kctx explain pod api-xyz --namespace payments
kctx explain applications.argoproj.io guestbook --namespace argocd
kctx graph
Build a graph around a supported resource. JSON is the default output; Mermaid and DOT renderers are available for graph-oriented views.
kctx graph pod api-xyz --namespace payments
kctx graph pod api-xyz --namespace payments --render mermaid
kctx graph applications.argoproj.io guestbook --namespace argocd --render dot
kctx trace service
Trace a Service to EndpointSlices, endpoints, Pods, owners, Nodes, and factual service health signals.
kctx trace service payments-api --namespace payments
kctx health namespace
Produce a compact namespace health snapshot.
kctx health namespace payments
kctx dump namespace
Export a deterministic namespace context snapshot for automation, incident review, or AI-agent grounding.
kctx dump namespace payments > payments-dump.json
kctx serve
Expose the same context engine through a lightweight read-only HTTP API.
kctx serve
curl http://localhost:8080/health/namespace/default
CRD Adapters
kctx can fetch arbitrary Kubernetes resources through discovery, but it does not pretend that every custom resource can be understood generically.
Ecosystem-specific semantics belong in explicit adapters. An adapter can turn a CRD into the same core contract used everywhere else: resource identity, compact status, related entities, relations, signals, and optionally graph nodes and edges.
The current adapter set covers Argo CD Application , Argo CD AppProject , and cert-manager Certificate resources.
The long-form documentation is organized as a PDF eBook .
If kctx saves you time when debugging Kubernetes workloads, consider supporting
its maintenance:
GitHub Sponsors: https://github.com/sponsors/lucasepe
PayPal: https://paypal.me/lucasepe71
Support helps fund release work, compatibility testing, documentation, and
ongoing Kubernetes, Argo CD, and cert-manager integration maintenance.
kctx is under active development. It is already useful as a read-only Kubernetes context tool, but production hardening is still in progress.
See ROADMAP.md for current production-readiness work, planned features, and open design areas.
A Kubernetes context engine for humans and AI agents.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
3
v0.2.1
Latest
Jun 8, 2026
+ 2 releases
Sponsor this project
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
