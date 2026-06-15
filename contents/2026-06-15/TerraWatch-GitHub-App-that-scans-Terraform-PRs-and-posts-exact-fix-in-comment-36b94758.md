---
source: "https://terrawatch.dev"
hn_url: "https://news.ycombinator.com/item?id=48546754"
title: "TerraWatch,GitHub App that scans Terraform PRs and posts exact fix in comment"
article_title: "TerraWatch — Terraform Security for GitHub PRs"
author: "alejny"
captured_at: "2026-06-15T21:55:11Z"
capture_tool: "hn-digest"
hn_id: 48546754
score: 1
comments: 0
posted_at: "2026-06-15T20:39:52Z"
tags:
  - hacker-news
  - translated
---

# TerraWatch,GitHub App that scans Terraform PRs and posts exact fix in comment

- HN: [48546754](https://news.ycombinator.com/item?id=48546754)
- Source: [terrawatch.dev](https://terrawatch.dev)
- Score: 1
- Comments: 0
- Posted: 2026-06-15T20:39:52Z

## Translation

タイトル: TerraWatch、Terraform PR をスキャンし、正確な修正をコメントに投稿する GitHub アプリ
記事のタイトル: TerraWatch — GitHub PR のための Terraform セキュリティ
説明: TerraWatch は、すべての Terraform プル リクエストでセキュリティの脆弱性をスキャンし、安全でないマージをブロックし、正確なコード修正を PR コメントとして投稿します。

記事本文:
Terraform セキュリティ · GitHub ネイティブ · ゼロ構成
出荷する PR
脆弱性
決して合併しません。
TerraWatch は、Terraform の構成ミスを AWS に到達する前に検出し、マージをブロックし、正確な修正をすべて GitHub 内でポストします。
GitHub に接続します — 無料です
仕組みを見る→
AI による修正はありません
何も自動適用されません
Checkov 依存関係なし
PR の差分のみをスキャンし、コードベース全体をスキャンすることはありません
github.com / acme-corp / インフラ / プル / 247
特技: ユーザーアップロードストレージ用に S3 バケットを追加
開く
alex-chen は、feat/s3-uploads から 3 つのコミットを main にマージしたいと考えています
✓
⟳
✕
すべてのチェックに合格しました
TerraWatch · Terraform の変更をスキャン中…
TerraWatch · 2 つの重要な発見 — マージがブロックされました
プルリクエストをマージする
TW
terrawatch-bot · ボット · たった今
🔴
[TW-S3-001] クリティカル — S3 バケットではパブリック ACL が許可されます
main.tf:14 のリソース aws_s3_bucket.user_uploads — acl = "public-read" はすべてのオブジェクトをパブリックに公開します。以下の修正を適用して押してください。
main.tf · 修正案
14 acl = "パブリック読み取り"
14 acl = "プライベート"
15
16 block_public_acls = true
17 block_public_policy = true
18strict_public_buckets = true
29 セキュリティルール
スキャン時間は 30 秒未満
0 ゼロ構成
100% GitHub ネイティブ
プロセス
マージを開始する PR がブロックされました
30秒以内に
YAML はありません。エージェントはいません。設定ファイルはありません。 GitHub アプリをインストールしてプッシュします。
.tf ファイルに触れるプル リクエストは、GitHub Webhook 経由で自動的に TerraWatch をトリガーします。オプトインもラベルもありません。
エンジンは、変更されたすべての Terraform リソースを解析し、それに対して 29 のセキュリティ ルールすべてを実行します。スキャンは 30 秒以内に完了します。
必要な GitHub ステータス チェックが失敗します。すべての重要な検出結果が解決されるまで、マージ ボタンはロックされ、オーバーライドもバイパスもできません。
TerraWatch は、必要な変更 (ルール ID、重大度、ファイル、行、コピーアンドペーストの差分) を正確にコメントするため、修正にかかる時間は最短になります。

日ではなくユート。
セキュリティ チームのすべて
実際に必要なもの
迅速に出荷するが、個別のセキュリティ レビュー サイクルを行う余裕がないチーム向けに構築されています。
GitHub アプリをインストールします。それでおしまい。構成ファイル、YAML、エージェントを展開または保守する必要はありません。最初に開いた PR で機能します。
すべての結果には、ファイル名、行番号、すぐに適用できる前/後の diff ブロックなど、必要な正確なコード変更が含まれています。
重大な発見がある場合は、必要な GitHub ステータス チェックを通じてマージ ボタンをロックします。上書きボタンはありません。脆弱性はすり抜けられません。
すべてのリポジトリにわたるすべての未解決の調査結果 (重大度の内訳、修正速度、チームの傾向) を単一のエンジニアリング グレードのダッシュボードで追跡します。
S3、IAM、RDS、EC2、VPC、EKS、Lambda、CloudFront などをカバーします。すべてのルールは、実際の AWS + Terraform スタック用に作成されました。
TerraWatch は Terraform コードを保存しません。 PR の差分を読み取り、メモリ内をスキャンして、すぐに破棄します。あなたのインフラはあなたのもののままです。
すべてのルールは初日に出荷されます。構成、プラグインのインストール、管理するルール パックは必要ありません。
ジラチケットはありません。緩みスレッドはありません。 「セキュリティ文書を確認してください」はありません。
TerraWatch は、必要な正確な変更を含むボット コメントを投稿します —
ルール ID、重大度、リソース名、ファイルの場所、および差分
30秒で適用できます。
開発者が修正をプッシュすると、TerraWatch が自動的に再スキャンし、
結合ボタンが緑色になります。終わり。
コンテキストの切り替えは速度を犠牲にします。発見と修正を PR 内 (すでに作業が行われている場所) に保持することは、チームがトリアージではなく出荷に時間を費やすことを意味します。
無料で始めましょう。摩擦なしで成長します。
契約はありません。最低座席数はありません。いつでもキャンセルしてください。

## Original Extract

TerraWatch scans every Terraform pull request for security vulnerabilities, blocks unsafe merges, and posts the exact code fix as a PR comment.

Terraform security · GitHub native · Zero config
The PR that ships
a vulnerability
never merges.
TerraWatch catches Terraform misconfigurations before they hit AWS, blocks the merge, and posts the exact fix — all inside GitHub.
Connect GitHub — it's free
See how it works →
No AI-generated fixes
Nothing auto-applied
No Checkov dependency
Only scans PR diffs — never your full codebase
github.com / acme-corp / infra / pull / 247
feat: add S3 bucket for user upload storage
Open
alex-chen wants to merge 3 commits into main from feat/s3-uploads
✓
⟳
✕
All checks passed
TerraWatch · Scanning Terraform changes…
TerraWatch · 2 critical findings — merge blocked
Merge pull request
TW
terrawatch-bot · bot · just now
🔴
[TW-S3-001] CRITICAL — S3 bucket allows public ACL
Resource aws_s3_bucket.user_uploads in main.tf:14 — acl = "public-read" exposes every object publicly. Apply the fix below and push.
main.tf · suggested fix
14 acl = "public-read"
14 acl = "private"
15
16 block_public_acls = true
17 block_public_policy = true
18 restrict_public_buckets = true
29 Security rules
<30s Scan time
0 Zero config
100% GitHub native
Process
PR open to merge blocked
in under 30 seconds
No YAML. No agents. No config files. Install the GitHub App and push.
Any pull request touching .tf files triggers TerraWatch automatically via GitHub webhooks. No opt-in, no labels.
The engine parses every changed Terraform resource and runs all 29 security rules against it. Scan completes in under 30 seconds.
A required GitHub status check fails. The merge button is locked — no override, no bypass — until every critical finding is resolved.
TerraWatch comments the precise change needed — rule ID, severity, file, line, and a copy-paste diff — so fixes take minutes, not days.
Everything your security team
actually needs
Built for teams who ship fast and can't afford a separate security review cycle.
Install the GitHub App. That's it. No config files, no YAML, no agents to deploy or maintain. Works on the first PR you open.
Every finding includes the precise code change needed — with the file name, line number, and a ready-to-apply before/after diff block.
Critical findings lock the merge button via required GitHub status checks. There is no override button. Vulnerabilities cannot slip through.
Track every open finding across all repos — severity breakdown, fix velocity, team trends — in a single engineering-grade dashboard.
Covering S3, IAM, RDS, EC2, VPC, EKS, Lambda, CloudFront, and more. Every rule was written for real-world AWS + Terraform stacks.
TerraWatch never stores your Terraform code. We read the PR diff, scan in memory, and discard it immediately. Your infra stays yours.
Every rule ships on day one. No config, no plugin installs, no rule packs to manage.
No Jira ticket. No Slack thread. No "check the security doc."
TerraWatch posts a bot comment with the precise change needed —
rule ID, severity, resource name, file location, and a diff that
can be applied in 30 seconds.
Developer pushes the fix, TerraWatch re-scans automatically,
merge button goes green. Done.
Context switching kills velocity. Keeping the finding and the fix inside the PR — where the work already is — means your team spends time shipping, not triaging.
Start free. Grow without friction.
No contracts. No seat minimums. Cancel any time.
