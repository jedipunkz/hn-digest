---
source: "https://github.com/gnana997/bumper"
hn_url: "https://news.ycombinator.com/item?id=48460814"
title: "Show HN: Bumper – block a destructive Terraform apply or a malicious dependency"
article_title: "GitHub - gnana997/bumper: One deterministic gate for your two riskiest moments: the terraform apply that reshapes your cloud, and the dependency you're about to install. · GitHub"
author: "gnana097"
captured_at: "2026-06-09T16:06:46Z"
capture_tool: "hn-digest"
hn_id: 48460814
score: 2
comments: 0
posted_at: "2026-06-09T13:23:06Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Bumper – block a destructive Terraform apply or a malicious dependency

- HN: [48460814](https://news.ycombinator.com/item?id=48460814)
- Source: [github.com](https://github.com/gnana997/bumper)
- Score: 2
- Comments: 0
- Posted: 2026-06-09T13:23:06Z

## Translation

タイトル: HN を表示: バンパー – 破壊的な Terraform 適用または悪意のある依存関係をブロックします。
記事のタイトル: GitHub - gnana997/bumper: クラウドを再構築する Terraform 適用と、インストールしようとしている依存関係という 2 つの最も危険な瞬間に対応する 1 つの決定的なゲート。 · GitHub
説明: クラウドを再構築する Terraform 適用と、インストールしようとしている依存関係という 2 つの最も危険な瞬間に対応する 1 つの決定的なゲート。 - gnana997/バンパー

記事本文:
GitHub - gnana997/bumper: クラウドを再構築する Terraform 適用と、インストールしようとしている依存関係という 2 つの最も危険な瞬間に対する 1 つの決定的なゲート。 · GitHub
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
グナナ997
/
バンパー
公共
通知
サインインする必要があります

通知設定を変更するには
追加のナビゲーション オプション
コード
この GitHub アクションをプロジェクトで使用する このアクションを既存のワークフローに追加するか、新しいワークフローを作成します マーケットプレイスで表示する メイン ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
43 コミット 43 コミット .claude-plugin .claude-plugin .github .github cmd/ バンパー cmd/ バンパー deps deps docs docs e2e e2e 例 例 内部 内部ツール ツール .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml COTRIBUTING.md COTRIBUTING.md ライセンスライセンス Makefile Makefile NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md action.yml action.yml go.mod go.mod go.sum go.sum install.sh install.sh lefthook.yml lefthook.yml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Terraform プランと依存関係のセーフティ ゲート — 1 つの決定的なゲート
最も危険な 2 つの瞬間: クラウドを再構築する Terraform の適用と、
インストールしようとしている依存関係。
バンパー.sh · ドキュメント · 例 · GitHub アクション
バンパーは、2 つのクラスの変更を発生前にブロックします。
AWS を破壊または公開する Terraform 適用、
GCP 、または Azure アカウント — terraform show -json プランから読み取ります。
既知の脆弱性または既知の悪意のあるものを取り込むインストール
依存関係 — ロックファイルから読み取ります。
どちらの判定も 100% 決定的であるため、マージをブロックしても安全です。
または AI エージェントを適用します。単一の静的な Go バイナリ。 API キーもアカウントもありません。
破壊的な適用をキャッチします。終了状態だけでなく、プラン diff を読み取ります。
$ terraform show -json plan.tfplan > plan.json
$バンパープラン.json
CRITICAL aws_db_instance.orders この適用により、最終スナップショットのないデータベースが破壊され、再作成されます。
ルール AWS_DB_DESTRUCTIVE_REPLACE_NO_SNAPSHOT
修正するには、skip_final_snapshot = false を設定するか、適用前に置換を強制するものを見つけてください。

NG。
CRITICAL aws_security_group.api 機密ポート範囲へのパブリック インターネットイングレス (0.0.0.0/0)
ルール AWS_SG_PUBLIC_INGRESS
修正 cidr_blocks/ipv6_cidr_blocks を既知の範囲に制限し、ポート範囲を狭めます。
3 つの所見 2 重大 · 1 高い
$エコー$?
1
有害な依存関係 (既知の CVE および既知の悪意のあるパッケージ) を捕捉します。
$ バンパー DEPS パッケージ-lock.json
MALICIOUS flatmap-stream@0.1.1 (npm) — MAL-2025-20690: flatmap-stream (npm) 内の悪意のあるコード
CRITICAL イベントストリーム@3.3.6 (npm) GHSA-mh6f-8j2x-4483 修正 → 4.0.0
クリティカル lodash@4.17.4 (npm) CVE-2019-10744 修正 → 4.17.12
CRITICAL minimist@1.2.0 (npm) CVE-2021-44906 修正 → 1.2.6
脆弱なパッケージが 4 つ、悪意のあるパッケージが 1 つ。
$エコー$?
1
どちらの例もこのリポジトリから実行できます。examples/ を参照してください。で
ターミナルでは、重大度が色分けされます (重大な赤色、高いオレンジ色)。パイプ/CI 出力
無地のままです。 AI の平易な英語のウォークスルー用に --explain (プラン スキャン) を追加します。
# 自作 (macOS)
brew install gnana997/tap/bumper
# インストール スクリプト (macOS / Linux) — チェックサム検証済みの最新リリースをダウンロードします
カール -fsSL https://get.bumper.sh |しー
# 行く
github.com/gnana997/bumper/cmd/bumper@latest をインストールしてください
すべてのリリースはチェックサムと連署署名 (キーレス) が行われ、SLSA が適用されます。
build-provenance attestation — docs/architecture.md → Releases および
バイナリを検証するための出所
このリポジトリの CI からのものです。
# Terraform: 計画をスキャンする
terraform plan -out plan.tfplan
terraform show -json plan.tfplan > plan.json
バンパープラン.json
バンパー --explain plan.json # 平易な英語のエンリッチメントを追加
# 依存関係: ロックファイルをスキャンします (npm / pip / uv / go.sum / …; 自動検出)
バンパー DEPS パッケージ-lock.json
バンパー deps # 現在のディレクトリ内のロックファイルを自動検出する
終了コード: 0 = クリーン、1 = 検出結果あり (CI フレンドリー)、2 = 使用法/解析
エラー。出力形式（ボット

h スキャン): --format text (デフォルト) · json · sarif
・値下げ。パッケージの座標のみが deps スキャンのためにマシンから離れます —
決してあなたのコードではありません。
CLI / CI ゲート — テキスト · JSON · SARIF (セキュリティ タブ) · スティッキー PR コメント、
計画スキャンと依存関係スキャンの両方に適用されます。
エージェント ガード フック — 未検証の Terraform 適用と
既知の悪意のあるパッケージを実行前にインストールし、実行後に依存関係をスキャンします
インストールします。それらをバンパー init で接続します。
Hosted Advisor MCP — エージェントがクエリを実行します。
ベスト プラクティス、CVE、マルウェア データ (ルックアップのみ - コードが外部に残されることはありません)
マシン）。
エージェント スキル — SKILL.md プレイブック
エージェントに上記のすべてを実行するように教えます。バンパー初期化でそれらをインストールするか、
npx スキルは、任意のエージェントに gnana997/bumper を追加します。
すでにお持ちの AI CLI ( claude 、 gemini 、 codex 、 opencode 、 auggie )
必要に応じて、各結果を平易な英語で説明します (セットアップゼロ、コストゼロ)。の
決定的コアが存在しない場合、それは単独で存在します。
バンパー初期化 # クロード コード (自動検出)
バンパー初期化 --agent Augment # または Augment
バンパー init --agent gemini # または Gemini CLI
できること
計画をスキャンする
Terraform プランでのフラグの露出/破棄、オプションの AI エンリッチメント
docs/cli.md
依存関係をスキャンする
ロックファイルからの脆弱な + 悪意のあるパッケージにフラグを付ける
docs/cli.md
適用を強制する
verify は、合格したスキャンを sha256 によってプランにバインドします。ガードブロックは未検証の適用をブロックします
docs/agents.md
エージェントのガードレール
バンパー初期化は、ガードフックとホストされた Advisor MCP を Claude Code、Augment、および Gemini CLI に接続します ( --agent )
docs/agents.md
依存関係のガードレール
悪意のあるインストールをブロックし、エージェント ループと CI で CVE の DEP をスキャンします。
docs/agents.md
エージェントのスキル
バンパー スキルをインストールすると、エージェントにバンパーの操作を教える SKILL.md プレイブックが追加されます
docs/agents.md
CI / GitHub アクション
SARIF を [セキュリティ] タブに移動し、最初に

不快な PR コメント、高+で失敗
docs/ci.md
カタログを検索する
バンパー検索ランクの適用されたルール + 推奨ベストプラクティス カタログ
docs/cli.md
インタラクティブコンソール
バンパー トゥイ — 恐ろしいローカル適用のための「ハザード コンソール」
docs/cli.md
適用範囲
Terraform ルール — 112 厳選、100% 決定的: 20 重大 · 57 高 ·
中 32、低 3、AWS (60)、GCP (35)、Azure (17): a
一貫したクラウド間ベースライン (パブリック ストレージ/データベース、オープンな管理ポート、
パブリック k8s コントロール プレーン、ワイルドカード/過剰特権 IAM、TLS、暗号化、パブリック
スナップショット、破壊）に加え、クラウドごとの詳細なカバレッジを実現します。すべてのルールは手作業で移植されています
一時的で否定的な備品があり、その起源を伝えています
(ソース: アップストリームの AVD-* ID を含む trivy、またはソース:custom )。適用範囲
マップ、ルール形式、および独自の記述方法: docs/rules.md 。
依存関係データ — ホストされたアドバイザー: 全体にわたる CVE/OSV アドバイザリー
npm、PyPI、Maven、Go、crates、RubyGems、NuGet、および既知の悪意のあるパッケージ
インテル ( MAL-* ) — 毎日更新され、AI によって書かれた修復の洞察が含まれます。
クリティカル/高い CVE。バンパー Deps とエージェント malware-gate はパッケージごとにクエリを実行します
コーディネートのみ。
オフライン アドバイザリー カタログ: バンパー検索も、知識のみで約 2,600 件に及びます
Trivy、Checkov、KICS、および Prowler から正規化されたベスト プラクティス エントリ
エージェントは、「Terraform for X を作成する前に何を焼き付けるべきですか?」と尋ねることができます。完全に
オフライン。
最終状態だけでなく、プランの差分も読み取ります。ほとんどのスキャナーは
結果の構成。バンパーはトランジションもチェックします (作成 / 削除 /
replace ) — 「この適用によりデータベースが破壊されます」をキャッチする唯一の方法です。
この破壊クラスが差別化要因となります。
バグだけでなく悪意も捕らえます。依存関係スキャンの既知の脆弱性フラグ
バージョンと既知の悪意のあるパッケージ — サプライチェーン攻撃だけでなく、

古い CVE。
警告するだけではなく強制します。 verify は、合格したスキャンを
sha256 による正確な計画。ガードフックは未検証のブロックをブロックします
適用/破棄。無視できるリンターは、無視できないゲートになります。
エージェント時代に向けて構築されました。ツールレイヤーのガードフック (Terraform 適用 + パッケージ)
インストール）に加えて、ホストされた Advisor MCP を使用すると、AI エージェントがサイレントに実行できなくなります。
既知の悪意のあるパッケージを検証またはインストールしなかった場合に適用します。
決定論的コアは独立しています。 AI レイヤーは飾りです。それが存在しない場合、または
失敗しても、決定論的な結果はまだ完全であり、ブロックされています。
Checkov/Trivy (IaC) および dependabot/Snyk/Socket (deps) との比較 —
意図的に行わないことも含めて、 docs/comparison.md にあります。
2 つの複合アクション、同じ署名付きリリース バイナリ:
# Terraform プランのセーフティ ゲート — SARIF + スティッキー PR コメント、高 + で失敗
- 使用: gnana997/bumper@v1
付き:
プラン-json : plan.json
失敗重大度 : 高
# 依存関係スキャン — ロックファイルを自動検出します。マルウェアは常に仕事に失敗します
- 使用: gnana997/bumper/deps@v1
付き:
失敗重大度 : 高
入力、権限、SARIF、およびスティッキーコメントの動作については、次のドキュメントに記載されています。
docs/ci.md 。 (マーケットプレイスのリストは Terraform アクションです。
依存関係スキャンは deps サブパスにある同じリポジトリです)。
docs/cli.md
コマンドリファレンス — scan、deps、list、search、explain、verify、guard、tui、init
docs/rules.md
ルール形式 (YAML + CEL)、カバレッジ、アドバイザリー カタログ、独自の作成
docs/ci.md
GitHub アクション — 入力、権限、SARIF、スティッキー コメント
docs/agents.md
エージェント ガードレール - 2 つのツール層ゲート (Terraform 適用 + 依存関係インストール)、エージェント スキル、バンパー初期化、サポートされるエージェント
docs/mcp.md
ホストされた Advisor MCP — ツール、マシンから出るもの
docs/architecture.md
内部、技術スタック、

サプライチェーンの来歴、ロードマップ
例/
両方のゲートの実行可能で気密な例
e2e/
実際の Claude Code または Gemini CLI エージェントに対してガードレールを自分で実行します (手動、ローカル)
貢献とセキュリティ
新しいルールと現実世界の計画の備品が最も価値がある
貢献 — COTRIBUTING.md を参照してください。脆弱性を報告するには、
SECURITY.md に基づく非公開開示を使用します (公開問題ではありません)。
Apache-2.0 。 Apache-2.0 ソース (Trivy、
Checkov) NOTICE の帰属を保持します。 CIS ベンチマークのコンテンツは、
再配布されました。
クラウドを再構築する Terraform 適用と、インストールしようとしている依存関係という 2 つの最も危険な瞬間に対応する 1 つの決定的なゲート。
Apache-2.0ライセンス
貢献する
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
12
v1.2.1
最新の
2026 年 6 月 7 日
+ 11 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

One deterministic gate for your two riskiest moments: the terraform apply that reshapes your cloud, and the dependency you're about to install. - gnana997/bumper

GitHub - gnana997/bumper: One deterministic gate for your two riskiest moments: the terraform apply that reshapes your cloud, and the dependency you're about to install. · GitHub
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
gnana997
/
bumper
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Use this GitHub action with your project Add this Action to an existing workflow or create a new one View on Marketplace main Branches Tags Go to file Code Open more actions menu Folders and files
43 Commits 43 Commits .claude-plugin .claude-plugin .github .github cmd/ bumper cmd/ bumper deps deps docs docs e2e e2e examples examples internal internal tools tools .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md action.yml action.yml go.mod go.mod go.sum go.sum install.sh install.sh lefthook.yml lefthook.yml View all files Repository files navigation
A Terraform-plan & dependency safety gate — one deterministic gate for your
two riskiest moments: the terraform apply that reshapes your cloud, and the
dependency you're about to install.
bumper.sh · Docs · Examples · GitHub Action
bumper blocks two classes of change before they happen :
a terraform apply that would destroy or expose your AWS ,
GCP , or Azure account — read from a terraform show -json plan;
an install that pulls in a known-vulnerable or known-malicious
dependency — read from your lockfile.
Both verdicts are 100% deterministic , so it's safe to block a merge, an
apply, or an AI agent on them. A single static Go binary; no API key, no account.
Catch a destructive apply — reads the plan diff , not just the end state:
$ terraform show -json plan.tfplan > plan.json
$ bumper plan.json
CRITICAL aws_db_instance.orders This apply will DESTROY and recreate a database with no final snapshot
rule AWS_DB_DESTRUCTIVE_REPLACE_NO_SNAPSHOT
fix Set skip_final_snapshot = false, or find what forces replacement before applying.
CRITICAL aws_security_group.api Public internet ingress (0.0.0.0/0) to a sensitive port range
rule AWS_SG_PUBLIC_INGRESS
fix Restrict cidr_blocks/ipv6_cidr_blocks to known ranges and narrow the port range.
3 finding(s) 2 critical · 1 high
$ echo $?
1
Catch a poisoned dependency — known CVEs and known-malicious packages:
$ bumper deps package-lock.json
MALICIOUS flatmap-stream@0.1.1 (npm) — MAL-2025-20690: Malicious code in flatmap-stream (npm)
CRITICAL event-stream@3.3.6 (npm) GHSA-mh6f-8j2x-4483 fix → 4.0.0
CRITICAL lodash@4.17.4 (npm) CVE-2019-10744 fix → 4.17.12
CRITICAL minimist@1.2.0 (npm) CVE-2021-44906 fix → 1.2.6
4 vulnerable, 1 malicious package(s).
$ echo $?
1
Both examples are runnable from this repo — see examples/ . On a
terminal, severities are colored (critical red, high amber); piped/CI output
stays plain. Add --explain (plan scan) for an AI plain-English walkthrough.
# Homebrew (macOS)
brew install gnana997/tap/bumper
# install script (macOS / Linux) — downloads the latest release, checksum-verified
curl -fsSL https://get.bumper.sh | sh
# Go
go install github.com/gnana997/bumper/cmd/bumper@latest
Every release is checksummed, cosign-signed (keyless), and carries a SLSA
build-provenance attestation — see docs/architecture.md → Releases and
provenance to verify the binary
came from this repo's CI.
# Terraform: scan a plan
terraform plan -out plan.tfplan
terraform show -json plan.tfplan > plan.json
bumper plan.json
bumper --explain plan.json # add plain-English enrichment
# Dependencies: scan a lockfile (npm / pip / uv / go.sum / …; auto-detected)
bumper deps package-lock.json
bumper deps # auto-detect a lockfile in the current dir
Exit codes: 0 = clean, 1 = findings present (CI-friendly), 2 = usage/parse
error. Output formats (both scans): --format text (default) · json · sarif
· markdown . Only package coordinates leave your machine for a deps scan —
never your code .
CLI / CI gate — text · JSON · SARIF (Security tab) · a sticky PR comment,
for both plan and dependency scans.
Agent guard hooks — block an unverified terraform apply and a
known-malicious package install before they run, and scan dependencies after
install. Wire them in with bumper init .
Hosted Advisor MCP — your agent queries it for
best-practice, CVE, and malware data (lookup-only — your code never leaves the
machine).
Agent skills — SKILL.md playbooks that
teach the agent to drive all of the above. bumper init installs them, or
npx skills add gnana997/bumper for any agent.
An AI CLI you already have ( claude , gemini , codex , opencode , auggie )
optionally explains each finding in plain English — zero setup, zero cost. The
deterministic core stands alone if it's absent.
bumper init # Claude Code (auto-detected)
bumper init --agent augment # or Augment
bumper init --agent gemini # or Gemini CLI
What you can do
Scan a plan
flag exposure/destruction in a Terraform plan, optional AI enrichment
docs/cli.md
Scan dependencies
flag vulnerable + malicious packages from a lockfile
docs/cli.md
Enforce the apply
verify binds a passing scan to a plan by sha256; guard blocks an unverified apply
docs/agents.md
Agent guardrail
bumper init wires the guard hooks + the hosted Advisor MCP into Claude Code, Augment, and Gemini CLI ( --agent )
docs/agents.md
Dependency guardrail
block malicious installs, scan deps for CVEs — in the agent loop and in CI
docs/agents.md
Agent skills
bumper skills install adds SKILL.md playbooks that teach the agent to drive bumper
docs/agents.md
CI / GitHub Action
SARIF to the Security tab, a sticky PR comment, fail on high+
docs/ci.md
Search the catalog
bumper search ranks enforced rules + an advisory best-practice catalog
docs/cli.md
Interactive console
bumper tui — the "hazard console" for the scary local apply
docs/cli.md
Coverage
Terraform rules — 112 curated, 100% deterministic: 20 critical · 57 high ·
32 medium · 3 low, across AWS (60), GCP (35), and Azure (17): a
consistent cross-cloud baseline (public storage/databases, open admin ports,
public k8s control planes, wildcard/over-privileged IAM, TLS, encryption, public
snapshots, destruction) plus deep per-cloud coverage. Every rule is hand-ported
with a passing and a negative fixture, and carries its provenance
( source: trivy with the upstream AVD-* id, or source: custom ). Coverage
map, rule format, and how to write your own: docs/rules.md .
Dependency data — hosted Advisor : CVE/OSV advisories across
npm, PyPI, Maven, Go, crates, RubyGems, and NuGet, plus known-malicious package
intel ( MAL-* ) — refreshed daily, with AI-written remediation insights on
critical/high CVEs. bumper deps and the agent malware-gate query it by package
coordinate only.
Offline advisory catalog: bumper search also spans ~2,600 knowledge-only
best-practice entries normalized from Trivy, Checkov, KICS, and Prowler — so an
agent can ask "what should I bake in before writing Terraform for X?" fully
offline.
Reads the plan diff, not just the end state. Most scanners check the
resulting config. bumper also checks the transition ( create / delete /
replace ) — the only way to catch "this apply will destroy your database."
That destruction class is the differentiator.
Catches malice, not just bugs. Dependency scanning flags known-vulnerable
versions and known-malicious packages — the supply-chain attack, not only
the stale CVE.
It enforces, it doesn't just warn. verify binds a passing scan to the
exact plan by sha256; the guard hook then blocks an unverified
apply / destroy . A linter you can ignore becomes a gate you can't.
Built for the agent era. Tool-layer guard hooks (Terraform apply + package
installs) plus the hosted Advisor MCP mean your AI agent can no longer silently
apply infra it didn't verify or install a known-malicious package.
Deterministic core stands alone. The AI layer is garnish; if it's absent or
fails, the deterministic findings are still complete and blocking.
How it stacks up against Checkov/Trivy (IaC) and Dependabot/Snyk/Socket (deps) —
including what it deliberately doesn't do — is in docs/comparison.md .
Two composite actions, same signed release binary:
# Terraform plan safety gate — SARIF + sticky PR comment, fail on high+
- uses : gnana997/bumper@v1
with :
plan-json : plan.json
fail-severity : high
# Dependency scan — auto-detects lockfiles; malware always fails the job
- uses : gnana997/bumper/deps@v1
with :
fail-severity : high
Inputs, permissions, SARIF, and the sticky-comment behavior are documented in
docs/ci.md . (The Marketplace listing is the Terraform action; the
dependency scan is the same repo at the deps subpath.)
docs/cli.md
command reference — scan, deps, list, search, explain, verify, guard, tui, init
docs/rules.md
rule format (YAML + CEL), coverage, the advisory catalog, writing your own
docs/ci.md
the GitHub Actions — inputs, permissions, SARIF, sticky comment
docs/agents.md
the agent guardrail — the two tool-layer gates (Terraform apply + dependency install), agent skills, bumper init , supported agents
docs/mcp.md
the hosted Advisor MCP — tools, what leaves the machine
docs/architecture.md
internals, tech stack, supply-chain provenance, roadmap
examples/
runnable, hermetic examples for both gates
e2e/
run the guardrail against a real Claude Code or Gemini CLI agent yourself (manual, local)
Contributing & security
New rules and real-world plan fixtures are the most valuable
contributions — see CONTRIBUTING.md . To report a vulnerability,
use private disclosure per SECURITY.md (not a public issue).
Apache-2.0 . Built-in rules adapted from Apache-2.0 sources (Trivy,
Checkov) retain attribution in NOTICE ; CIS Benchmark content is not
redistributed.
One deterministic gate for your two riskiest moments: the terraform apply that reshapes your cloud, and the dependency you're about to install.
Apache-2.0 license
Contributing
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
12
v1.2.1
Latest
Jun 7, 2026
+ 11 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
