---
source: "https://github.com/k8s-security-pro/k8s-audit"
hn_url: "https://news.ycombinator.com/item?id=49399379"
title: "Show HN: K8s-audit – 30-second Kubernetes security check (kubectl and jq only)"
article_title: "GitHub - k8s-security-pro/k8s-audit: A fast, dependency-light Kubernetes security audit with kubectl and jq, mapped to the k8s-security.pro 50-point checklist. · GitHub"
image: "https://opengraph.githubassets.com/96740cfa467efa6aacaa9e254dfbbdcfb43bc385ca59ef0fa1f762e39075f373/k8s-security-pro/k8s-audit"
author: "k8ssecuritypro"
captured_at: "2026-08-22T13:24:34Z"
capture_tool: "hn-digest"
hn_id: 49399379
score: 1
comments: 0
posted_at: "2026-08-22T13:08:44Z"
tags:
  - hacker-news
  - translated
---

# Show HN: K8s-audit – 30-second Kubernetes security check (kubectl and jq only)

- HN: [49399379](https://news.ycombinator.com/item?id=49399379)
- Source: [github.com](https://github.com/k8s-security-pro/k8s-audit)
- Score: 1
- Comments: 0
- Posted: 2026-08-22T13:08:44Z

## Translation

タイトル: HN を表示: K8s-audit – 30 秒の Kubernetes セキュリティ チェック (kubectl および jq のみ)
記事のタイトル: GitHub - k8s-security-pro/k8s-audit: k8s-security.pro 50 ポイント チェックリストにマッピングされた、kubectl と jq を使用した高速で依存関係の少ない Kubernetes セキュリティ監査。 · GitHub
説明: k8s-security.pro 50 ポイント チェックリストにマップされた、kubectl と jq を使用した依存関係の少ない高速な Kubernetes セキュリティ監査。 - k8s-security-pro/k8s-audit

記事本文:
GitHub - k8s-security-pro/k8s-audit: k8s-security.pro 50 ポイント チェックリストにマッピングされた、kubectl と jq を使用した依存関係の少ない高速な Kubernetes セキュリティ監査。 · GitHub
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
k8s-セキュリティ-プロ
/
k8s-監査
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
4 コミット 4 コミット フォルダーとファイル
.github/ workflows .github/ workflows COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md READM

E.md デモ.svg デモ.svg k8s-audit.sh k8s-audit.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
依存関係の少ない高速な Kubernetes セキュリティ監査を 1 つのコマンドで実行できます。
k8s-audit は、現在のクラスターに対して一連の高信号セキュリティ チェックを実行します。
kubectl + jq のみを使用し、セキュリティ ドメインごとにグループ化されたクリーンなレポートを出力します。
特権コンテナ、ネットワークポリシーの欠落、広範な RBAC、:最新のイメージ、
リソース制限の不足など。
これは読み取り専用で、どこにも何も送信せず、シェルを介してインストールする必要はありません。
Kubernetes セキュリティのすべては技術的に「無料」です — CIS ベンチマーク、kube-bench、
キューブスケープ、トリビー。しかし、これらのツールは何百もの項目にフラグを立てているため、自分で判断する必要があります。
どれが重要で、どのように修正するか。 k8s-audit は最初の 30 秒が独断的なものです
合格: 最も一般的な現実世界のエクスポージャを検出する ~16 個のチェック。それぞれが次のようにマッピングされています。
k8s-security.pro 50 ポイント チェックリストの特定の項目。
k8s-security.pro の背後にあるチームによって構築および保守されています —
本番環境強化キット (50 ポイントの監査、25 の YAML テンプレート、Helm チャート、KusTOMize)
オーバーレイ、CIS および SOC2 マッピング）。このリポジトリは、無料のオープンソースの玄関口です。
インストールは必要ありません。クローンを作成して実行するだけです。
git clone https://github.com/k8s-security-pro/k8s-audit.git
cd k8s-監査
./k8s-audit.sh
要件: kubectl (クラスターを指す) および jq 。
./k8s-audit.sh # すべての名前空間を監査する
./k8s-audit.sh -npayments #単一の名前空間
./k8s-audit.sh --json # 機械可読出力 (CI / ダッシュボード用)
FAIL 重大度チェックがトリップした場合、終了コードはゼロ以外になるため、それに基づいて CI をゲートできます。
これを .github/workflows/k8s-security-audit.yml にドロップすると、PR が失敗します。
特権コンテナまたは保護されていない名前空間 (完全な例は次のとおりです)
.github/workflows/ ):
- 名前 : Kubernetes セキュリティ監査

それ
実行: |
カール -sSL https://raw.githubusercontent.com/k8s-security-pro/k8s-audit/main/k8s-audit.sh -o k8s-audit.sh
chmod +x k8s-audit.sh
./k8s-audit.sh
チェック内容 (16/50)
#
ドメイン
チェックする
2
ポッドのセキュリティ
特権コンテナ
3
ポッドのセキュリティ
allowPrivilegeEscalation が無効になっていません
4
ポッドのセキュリティ
root として実行 ( runAsNonRoot を設定解除)
4b
ポッドのセキュリティ
readOnlyRootFilesystem が設定されていません
5
ポッドのセキュリティ
ホストネットワーク / ホストPID / ホストIPC
6
ポッドのセキュリティ
ServiceAccount トークンの自動マウント
7
ポッドのセキュリティ
削除されない機能 (すべて)
7b
ポッドのセキュリティ
危険な機能の追加 ( SYS_ADMIN 、 NET_ADMIN など)
9
ポッドのセキュリティ
マウントされた hostPath ボリューム
8
ネットワーク
NetworkPolicy のない名前空間 (デフォルト拒否なし)
14
RBAC
クラスターと管理者のバインディング
15
RBAC
デフォルトの ServiceAccount を使用するワークロード
16
RBAC
* リソースに対する * 動詞を付与するロール
20
クラスターの強化
デフォルトの名前空間のワークロード
22
クラスターの強化
リソース制限のないコンテナー
28
サプライチェーン
:latest またはタグなしを使用した画像
さらに深くなる
k8s-audit は、信号の高い 12 個で意図的に停止します。全体像については、次のとおりです。
50 ポイントの完全な監査 + コピー＆ペーストによる修復 YAML、Helm チャート、KusTOMize
オーバーレイ、および CIS/SOC2 準拠マッピング → k8s-security.pro
より詳細な CIS/NSA スキャン → kube-bench 、
キューブスケープ、トリビー
(これらが PATH 上にある場合、 k8s-audit は正しいコマンドを示します)。
問題と PR は歓迎します - 特に新しい高信号チェック (kubectl + jq のままにしておきます)
のみ、読み取り専用、チェックリスト ドメインにマップされます)。プロジェクトは初めてですか?を探してください
良い創刊号
label — それぞれは、jq フィルターがスケッチされた小さな自己完結型のチェックです。
チェックの構造については、CONTRIBUTING.md を参照してください。
kubectl と jq を使用した、依存関係の少ない高速な Kubernetes セキュリティ監査。k8s-security.pro 50- にマッピングされます。

ポイントチェックリスト。
Readme MIT ライセンス
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A fast, dependency-light Kubernetes security audit with kubectl and jq, mapped to the k8s-security.pro 50-point checklist. - k8s-security-pro/k8s-audit

GitHub - k8s-security-pro/k8s-audit: A fast, dependency-light Kubernetes security audit with kubectl and jq, mapped to the k8s-security.pro 50-point checklist. · GitHub
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
k8s-security-pro
/
k8s-audit
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
4 Commits 4 Commits Folders and files
.github/ workflows .github/ workflows CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md demo.svg demo.svg k8s-audit.sh k8s-audit.sh View all files Repository files navigation
A fast, dependency-light Kubernetes security audit you can run in one command.
k8s-audit runs a set of high-signal security checks against your current cluster
using only kubectl + jq , and prints a clean report grouped by security domain —
privileged containers, missing NetworkPolicies, over-broad RBAC, :latest images,
missing resource limits, and more.
It's read-only, sends nothing anywhere, and needs no install beyond a shell.
Everything in Kubernetes security is technically "free" — CIS Benchmark, kube-bench,
Kubescape, Trivy. But those tools flag hundreds of items and leave you to figure out
which ones matter and how to fix them. k8s-audit is the opinionated 30-second first
pass: ~16 checks that catch the most common real-world exposures, each mapped to
a specific item in the k8s-security.pro 50-point checklist .
It's built and maintained by the team behind k8s-security.pro —
a production hardening kit (50-point audit, 25 YAML templates, Helm chart, Kustomize
overlays, CIS & SOC2 mappings). This repo is the free, open-source front door to it.
No install required — just clone and run:
git clone https://github.com/k8s-security-pro/k8s-audit.git
cd k8s-audit
./k8s-audit.sh
Requirements: kubectl (pointed at your cluster) and jq .
./k8s-audit.sh # audit all namespaces
./k8s-audit.sh -n payments # a single namespace
./k8s-audit.sh --json # machine-readable output (for CI / dashboards)
Exit code is non-zero if any FAIL -severity check trips — so you can gate CI on it.
Drop this into .github/workflows/k8s-security-audit.yml to fail a PR that introduces
a privileged container or an unprotected namespace (full example in
.github/workflows/ ):
- name : Kubernetes security audit
run : |
curl -sSL https://raw.githubusercontent.com/k8s-security-pro/k8s-audit/main/k8s-audit.sh -o k8s-audit.sh
chmod +x k8s-audit.sh
./k8s-audit.sh
What it checks (16 of 50)
#
Domain
Check
2
Pod Security
Privileged containers
3
Pod Security
allowPrivilegeEscalation not disabled
4
Pod Security
Running as root ( runAsNonRoot unset)
4b
Pod Security
readOnlyRootFilesystem not set
5
Pod Security
hostNetwork / hostPID / hostIPC
6
Pod Security
ServiceAccount token auto-mounted
7
Pod Security
Capabilities not dropped ( ALL )
7b
Pod Security
Dangerous capabilities added ( SYS_ADMIN , NET_ADMIN , …)
9
Pod Security
hostPath volumes mounted
8
Network
Namespaces without a NetworkPolicy (no default-deny)
14
RBAC
cluster-admin bindings
15
RBAC
Workloads using the default ServiceAccount
16
RBAC
Roles granting * verbs on * resources
20
Cluster Hardening
Workloads in the default namespace
22
Cluster Hardening
Containers without resource limits
28
Supply Chain
Images using :latest or untagged
Going deeper
k8s-audit deliberately stops at the high-signal dozen. For the complete picture:
The full 50-point audit + copy-paste remediation YAML, Helm chart, Kustomize
overlays, and CIS/SOC2 compliance mappings → k8s-security.pro
Deeper CIS/NSA scanning → kube-bench ,
Kubescape , Trivy
(if these are on your PATH , k8s-audit points you at the right command).
Issues and PRs welcome — especially new high-signal checks (keep them kubectl + jq
only, read-only, and mapped to a checklist domain). New to the project? Look for the
good first issue
label — each one is a small, self-contained check with the jq filter sketched out for you.
See CONTRIBUTING.md for how a check is structured.
A fast, dependency-light Kubernetes security audit with kubectl and jq, mapped to the k8s-security.pro 50-point checklist.
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
