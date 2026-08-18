---
source: "https://github.com/TuviDev/infra-lang"
hn_url: "https://news.ycombinator.com/item?id=49352976"
title: "Infra Lang – DSL That Compiles to Kubernetes, Helm, Compose, Terraform, and CI"
article_title: "GitHub - TuviDev/infra-lang: Write infrastructure once, compile it to Kubernetes, Compose, or GitHub Actions · GitHub"
image: "https://opengraph.githubassets.com/cdbdb9f34ad0920e6040d7b11cac91ba79c86ed211e7b4c5ee8ae1fe97c0ac9e/TuviDev/infra-lang"
author: "Tuvi"
captured_at: "2026-08-18T22:13:03Z"
capture_tool: "hn-digest"
hn_id: 49352976
score: 2
comments: 0
posted_at: "2026-08-18T21:29:47Z"
tags:
  - hacker-news
  - translated
---

# Infra Lang – DSL That Compiles to Kubernetes, Helm, Compose, Terraform, and CI

- HN: [49352976](https://news.ycombinator.com/item?id=49352976)
- Source: [github.com](https://github.com/TuviDev/infra-lang)
- Score: 2
- Comments: 0
- Posted: 2026-08-18T21:29:47Z

## Translation

タイトル: Infra Lang – Kubernetes、Helm、Compose、Terraform、および CI にコンパイルされる DSL
記事のタイトル: GitHub - TuviDev/infra-lang: インフラストラクチャを一度作成し、Kubernetes、Compose、または GitHub アクションにコンパイルする · GitHub
説明: インフラストラクチャを一度作成し、それを Kubernetes、Compose、または GitHub アクションにコンパイルします - TuviDev/infra-lang

記事本文:
GitHub - TuviDev/infra-lang: インフラストラクチャを一度作成し、Kubernetes、Compose、または GitHub アクションにコンパイルする · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
トゥビデヴ
/
インフラ言語
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
31 コミット 31 コミット フォルダーとファイル
.devcontainer .devcontainer .github .github docs docs 例 例 スクリプト スクリプト src src テスト テスト vscode-infra-lang vscode-infra-lang .gitignore .gi

tignore .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile PUBLISH.md PUBLISH.md README.md README.md SECURITY.md SECURITY.md mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml setup.cfg setup.cfg すべてのファイルを表示 リポジトリ ファイルのナビゲーション
インフラストラクチャを一度作成し、それを Kubernetes、Compose、または GitHub Actions にコンパイルします。
Infra Lang は、DevOps エンジニア、SRE、および
プラットフォームチーム。アプリケーション (サービス、データベース、キューなど) を説明します。
シークレットとパイプライン — 1 つの宣言的な .infra ファイルと Infra Lang
それを Kubernetes YAML、Docker Compose、Terraform HCL、または GitHub にコンパイルします
アクションのワークフロー。手書きで同じアプリを 4 つで管理するのではなく
さまざまな形式であっても、1 つの信頼できる情報源を維持できます。
単一の .infra ファイルでサービスを説明します。
#アプリ.インフラ
サービス API {
画像: 「myapp/api:v1.0.0」
レプリカ: 3
ポート8080
健康 http("/health")
リソース {
リクエスト { CPU: 200m、メモリ: 256Mi }
制限 { CPU: 1000m、メモリ: 512Mi }
}
}
それを Kubernetes にコンパイルします。
infra コンパイル app.infra --target kubernetes
Infra Lang は、一致するデプロイメントとサービスを生成します。
apiVersion : apps/v1
種類：展開
メタデータ:
名前 : アピ
仕様：
レプリカ：3
セレクター:
matchLabels :
app.kubernetes.io/名前 : API
テンプレート:
仕様：
コンテナ:
- 名前 : アピ
画像：myapp/api:v1.0.0
ポート:
- コンテナポート: 8080
名前: ポート-0
リソース:
リクエスト: { CPU: 200m、メモリ: 256Mi }
制限: { CPU: 1000m、メモリ: 512Mi }
レディネスプローブ:
httpGet : { パス: /health、ポート: 8080 }
---
APIバージョン : v1
種類 : サービス
メタデータ:
名前 : アピ
仕様：
セレクター:
app.kubernetes.io/名前 : API
ポート:
- ポート: 8080
ターゲット

eポート：8080
同じファイルは書き換えなしで Docker Compose にコンパイルされます。
infra コンパイル app.infra --target compose
パイプライン ブロックは GitHub Actions ワークフローにコンパイルされます。
パイプライン ci {
トリガー { ブランチ: ["メイン"] }
ステージ {
テスト: { runsOn: "ubuntu-latest" ステップ { t: { run: "pytest" } } }
}
}
infra コンパイル app.infra --target github
特長
11 のトップレベルのリソース タイプ — サービス、データベース、キャッシュ、キュー、
ストレージ、ネットワーク、シークレット、構成、パイプライン、環境、
クラスター。
5 つのコンパイル ターゲット — Kubernetes (17 種類のリソース)、Helm チャート、
Docker Compose、Terraform HCL (AWS/GCP/Azure)、GitHub アクション。
コンパイラ グレードの検証 — ソースの場所と 30 以上のエラー コード
実用的なヒント。無効な構成は、何かが出力される前に失敗します。
組み込みのセキュリティ リンター (SEC001 ～ SEC010) および信頼性リンター
(REL001–REL014);エラーの重大度の結果によりコンパイルがブロックされます。
言語サーバー — コンテキスト認識型補完、ホバードキュメント、ライブ
リンクと関連情報を含む診断、定義への移動、参照の検索、
ワークスペースシンボル、シンボル名の変更、署名ヘルプ、ドキュメントのハイライト、
セマンティック トークン、折りたたみ、書式設定、クイックフィックス - すべての
ディスク上の .infra ファイル。
フォーマッタ、REPL、および diff エンジン - infra fmt 、 infra repl 、および
変更を確認するための infra diff。
再利用可能な部分 - テンプレート文字列補間、サイクル付きインポート
検出、継承の拡張、25 以上の stdlib 関数とプレリュード
共有定数。
下のボタンをクリックして、このプロジェクトを GitHub コードスペースで開きます。
ローカルにインストールする必要はありません - 約 2 分で完全な開発環境が完成します
(Python 3.12、Docker-in-Docker、kubectl/helm、Ruff/Mypy 拡張機能)。
pip install git+https://github.com/TuviDev/infra-lang.git
言語サーバーを使用する場合 (VS Code に推奨):
pip インストール ' infra-lang[lsp

]」
確認:
インフラ --バージョン
インフラ --ヘルプ
注: PyPI の公開は近日中に開始されます。それまでは、Git からインストールしてください。
完全なドキュメントは TuviDev.github.io/infra-lang でホストされています。
最も早いパスは 5 分間のクイックスタートです。要するに:
.infra ファイルを作成します (上記のデモを参照)。
検証します: infra validate app.infra
ターゲットにコンパイルします: infracompile app.infra --target kubernetes
infra-out/ で出力を検査するか、 --dry-run でプレビューします。
infra fmt app.infra と infra diff app.infra app2.infra を繰り返します。
ガイド付きチュートリアルもあります。
コメント付きの例。
ターゲット
コマンド
それが生み出すもの
Kubernetes
-t Kubernetes
デプロイメント、サービス、Ingress、StatefulSet、PVC、ConfigMap、シークレット、CronJob、HPA、PDB、NetworkPolicies、ResourceQuotas、名前空間、RBAC、TopologySpreadConstraints
ヘルム
-t ヘルム
完全なチャート: Chart.yaml 、values.yaml 、templates/ 、_helpers.tpl 、.helmignore
Docker Compose
-t 作成する
docker-compose.yml 、 .env.example 、 Makefile
テラフォーム
-t テラフォーム
main.tf 、variables.tf 、outputs.tf 、providers.tf (AWS/GCP/Azure)
GitHub アクション
-t ギットハブ
.github/workflows/*.yml 、dependabot.yml
すべてのリソース タイプがすべてのターゲットにマップされるわけではありません - たとえば、パイプライン コンパイル
GitHub Actions のみ、および Terraform のみにクラスター化されます。を参照してください。
のサポート マトリックス
完全なマッピング。
ドキュメントは TuviDev.github.io/infra-lang でホストされています。
貢献は大歓迎です。方法については、CONTRIBUTING.md を参照してください。
開発環境をセットアップし、バックエンドまたは文法ルールを追加し、コーディングを行います。
標準 (ruff、mypy)。事前に当社のセキュリティ ポリシーをお読みください。
脆弱性を報告しています。
MIT ライセンスに基づいてライセンスされています。
Infra Lang は、Terraform の背後にあるアイデアに触発されています。
Score 、および Pulumi : 宣言型
読みやすく、間違いにくいインフラストラクチャ。
インフラストラクチャを一度作成してコンパイルする

Kubernetes、Compose、または GitHub アクションへ
tuvidev.github.io/infra-lang/ トピック
Readme MIT ライセンスの行動規範
セキュリティポリシー アクティビティスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Write infrastructure once, compile it to Kubernetes, Compose, or GitHub Actions - TuviDev/infra-lang

GitHub - TuviDev/infra-lang: Write infrastructure once, compile it to Kubernetes, Compose, or GitHub Actions · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
TuviDev
/
infra-lang
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
31 Commits 31 Commits Folders and files
.devcontainer .devcontainer .github .github docs docs examples examples scripts scripts src src tests tests vscode-infra-lang vscode-infra-lang .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile PUBLISH.md PUBLISH.md README.md README.md SECURITY.md SECURITY.md mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml setup.cfg setup.cfg View all files Repository files navigation
Write infrastructure once, compile it to Kubernetes, Compose, or GitHub Actions.
Infra Lang is an Infrastructure-as-Code DSL for DevOps engineers, SREs, and
platform teams. You describe your application — services, databases, queues,
secrets, and pipelines — in one declarative .infra file, and Infra Lang
compiles it to Kubernetes YAML, Docker Compose, Terraform HCL, or a GitHub
Actions workflow. Instead of hand-writing and maintaining the same app in four
different formats, you maintain one source of truth.
A single .infra file describes a service:
# app.infra
service api {
image: "myapp/api:v1.0.0"
replicas: 3
port 8080
health http("/health")
resources {
requests { cpu: 200m, memory: 256Mi }
limits { cpu: 1000m, memory: 512Mi }
}
}
Compile it to Kubernetes:
infra compile app.infra --target kubernetes
Infra Lang produces the matching Deployment and Service:
apiVersion : apps/v1
kind : Deployment
metadata :
name : api
spec :
replicas : 3
selector :
matchLabels :
app.kubernetes.io/name : api
template :
spec :
containers :
- name : api
image : myapp/api:v1.0.0
ports :
- containerPort : 8080
name : port-0
resources :
requests : { cpu: 200m, memory: 256Mi }
limits : { cpu: 1000m, memory: 512Mi }
readinessProbe :
httpGet : { path: /health, port: 8080 }
---
apiVersion : v1
kind : Service
metadata :
name : api
spec :
selector :
app.kubernetes.io/name : api
ports :
- port : 8080
targetPort : 8080
The same file compiles to Docker Compose with no rewriting:
infra compile app.infra --target compose
A pipeline block compiles to a GitHub Actions workflow:
pipeline ci {
trigger { branches: ["main"] }
stages {
test: { runsOn: "ubuntu-latest" steps { t: { run: "pytest" } } }
}
}
infra compile app.infra --target github
Features
11 top-level resource types — service , database , cache , queue ,
storage , network , secret , config , pipeline , environment ,
cluster .
5 compilation targets — Kubernetes (17 resource kinds), Helm charts ,
Docker Compose, Terraform HCL (AWS/GCP/Azure), GitHub Actions.
Compiler-grade validation — 30+ error codes with source locations and
actionable hints; invalid configs fail before anything is emitted.
Built-in security linter (SEC001–SEC010) and reliability linter
(REL001–REL014); Error -severity findings block compilation.
A language server — context-aware completion, hover docs, live
diagnostics with links and related info, go-to-definition, find-references,
workspace symbols, symbol rename, signature help, document highlight,
semantic tokens, folding, formatting, and quick-fixes — all across every
.infra file on disk.
A formatter, REPL, and diff engine — infra fmt , infra repl , and
infra diff for reviewing changes.
Reusable pieces — template-string interpolation, import with cycle
detection, extends inheritance, 25+ stdlib functions and a prelude of
shared constants.
Click the button below to open this project in GitHub Codespaces:
No local installation needed — full dev environment in about 2 minutes
(Python 3.12, Docker-in-Docker, kubectl/helm, Ruff/Mypy extensions).
pip install git+https://github.com/TuviDev/infra-lang.git
With the language server (recommended for VS Code):
pip install ' infra-lang[lsp] '
Verify:
infra --version
infra --help
Note: PyPI publishing is coming soon. Until then, install from Git.
Full documentation is hosted at TuviDev.github.io/infra-lang .
The fastest path is the 5-minute quickstart . In short:
Write a .infra file (see the demo above).
Validate it: infra validate app.infra
Compile to a target: infra compile app.infra --target kubernetes
Inspect the output in infra-out/ , or preview with --dry-run .
Iterate with infra fmt app.infra and infra diff app.infra app2.infra .
There is also a guided tutorial and
commented examples .
Target
Command
What it generates
Kubernetes
-t kubernetes
Deployments, Services, Ingress, StatefulSets, PVCs, ConfigMaps, Secrets, CronJobs, HPA, PDBs, NetworkPolicies, ResourceQuotas, Namespaces, RBAC, TopologySpreadConstraints
Helm
-t helm
A complete chart: Chart.yaml , values.yaml , templates/ , _helpers.tpl , .helmignore
Docker Compose
-t compose
docker-compose.yml , .env.example , Makefile
Terraform
-t terraform
main.tf , variables.tf , outputs.tf , providers.tf (AWS/GCP/Azure)
GitHub Actions
-t github
.github/workflows/*.yml , dependabot.yml
Not every resource type maps to every target — for example, pipeline compiles
only to GitHub Actions, and cluster only to Terraform. See the
support matrix for the
full mapping.
The documentation is hosted at TuviDev.github.io/infra-lang .
Contributions are welcome. See CONTRIBUTING.md for how to
set up a dev environment, add a backend or a grammar rule, and the coding
standards (ruff, mypy). Please read our Security policy before
reporting a vulnerability.
Licensed under the MIT License .
Infra Lang is inspired by the ideas behind Terraform ,
Score , and Pulumi : declarative
infrastructure that is easy to read and hard to get wrong.
Write infrastructure once, compile it to Kubernetes, Compose, or GitHub Actions
tuvidev.github.io/infra-lang/ Topics
Readme MIT license Code of conduct
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
