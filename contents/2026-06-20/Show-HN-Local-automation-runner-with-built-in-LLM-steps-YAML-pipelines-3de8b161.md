---
source: "https://rorlikowski.github.io/stepyard/"
hn_url: "https://news.ycombinator.com/item?id=48607721"
title: "Show HN: Local automation runner with built-in LLM steps – YAML pipelines"
article_title: "Stepyard"
author: "rorlikowski"
captured_at: "2026-06-20T10:04:48Z"
capture_tool: "hn-digest"
hn_id: 48607721
score: 2
comments: 0
posted_at: "2026-06-20T09:21:46Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Local automation runner with built-in LLM steps – YAML pipelines

- HN: [48607721](https://news.ycombinator.com/item?id=48607721)
- Source: [rorlikowski.github.io](https://rorlikowski.github.io/stepyard/)
- Score: 2
- Comments: 0
- Posted: 2026-06-20T09:21:46Z

## Translation

タイトル: HN の表示: LLM ステップが組み込まれたローカル オートメーション ランナー – YAML パイプライン
記事タイトル: ステップヤード
説明: 開発者優先のオートメーション ランナー - Python で拡張された YAML としてのパイプラインは、CLI またはローカル スケジューラから実行されます。

記事本文:
ステップヤード
コンテンツにスキップ
ステップヤード
ホーム
検索を初期化しています
ロリコウスキー/ステップヤード
ホーム
ステップヤード
ロリコウスキー/ステップヤード
ホーム
ホーム
目次
なぜステップヤードなのか？
実際の例
毎日のデータベースのバックアップ
すべての PR での自動コードレビュー
ロールバックを伴う複数環境の展開
はじめに
はじめに
インストール
中心となる概念
中心となる概念
フローとステップ
ハウツーガイド
ハウツーガイド
フローのスケジュールを設定する
組み込みノード
組み込みノード
全ノード参照
プラグイン開発
プラグイン開発
クイックスタート
実際の例
毎日のデータベースのバックアップ
すべての PR での自動コードレビュー
ロールバックを伴う複数環境の展開
YAML パイプライン。 Python プラグイン。お使いのマシン上で実行できます。サーバーのセットアップやクラウド アカウントは必要ありません。
Stepyard は開発者向けの自動化ランナーです。パイプラインはリポジトリ内に YAML ファイルとして存在します。単純な Python 関数を使用して拡張します。すべては、自分のマシンまたは所有するサーバー上の CLI またはローカル デーモンから実行されます。
pip インストール ステップヤード
stepyard init my-project && cd my-project
ステップヤードラン、こんにちは
デプロイ パイプライン: コンテナーを構築してプッシュし、スモーク テストを実行し、結果を Slack に投稿します。 1 つの YAML ファイル、サポート スクリプトなし。
名前 : デプロイ
説明 : 運用コンテナをビルド、プッシュ、検証します。
手順:
- ID : ビルド
使用:shell.run
付き:
コマンド: docker build -t myapp:${{ env.GIT_SHA }} 。
- ID : プッシュ
使用:shell.run
付き:
コマンド: docker Push myapp:${{ env.GIT_SHA }}
- ID : スモークテスト
使用: http.request
付き:
URL : https://staging.myapp.com/healthz
メソッド: GET
- ID : 通知
使用: llm.generate # 組み込み - プラグインは必要ありません
付き:
モデル: gpt-4o-mini
プロンプト: |
このデプロイ結果を Slack 用に 1 つの文に要約します。
HTTP ステータス: ${{ ステップ.smoke_test.output.status }}
SHA: ${{ env.GIT_SHA }}
- id : post_to_slack
使用: http

.リクエスト
付き:
URL : ${{ env.SLACK_WEBHOOK }}
メソッド: POST
json_body :
テキスト: ${{ ステップ.通知.出力.出力 }}
実行してください:
GIT_SHA = $( git rev-parse --short HEAD ) ステップヤード実行デプロイ
✓ ビルド 12.4 秒
✓ 4.1 秒押します
✓ スモークテスト 0.3 秒
✓ 0.9秒で通知
✓ post_to_slack 0.2 秒
フローは 18.0 秒で完了
なぜステップヤードなのか？ ¶
ステップ、条件、ループ、および再試行はプレーンな YAML キーであり、独自の DSL はありません。コードと一緒にバージョン管理し、 stepyard validate で検証します。
1 つの @node デコレーターは、あらゆる関数を再利用可能なステップに変換します。入力は自動的に型検証されます。プラグインの依存関係は分離されているため、Stepyard の依存関係と競合することはありません。
状態はローカルの SQLite データベースに保存されます。データは、フロー内のステップで明示的に送信された場合にのみ送信されます。
スケジュールされた実行とオンデマンドの実行
cron 、 interval 、およびstartupトリガー。 stepyard service start を実行すると、フローがスケジュールどおりに実行されます。外部サービスやクラウド アカウントは必要ありません。
名前 : pg_backup
トリガー:
使用: cron
付き:
スケジュール : "0 3 * * *" # 毎日 03:00
手順:
- ID : ダンプ
使用:shell.run
付き:
コマンド: pg_dump ${{ env.DATABASE_URL }} | gzip > /tmp/backup.sql.gz
- ID : アップロード
使用:shell.run
付き:
コマンド: |
aws s3 cp /tmp/backup.sql.gz \
s3://${{ env.BACKUP_BUCKET }}/db/$(日付 +%Y-%m-%d).sql.gz
- ID : クリーンアップ
使用:shell.run
付き:
コマンド: rm -f /tmp/backup.sql.gz
すべての PR での自動コードレビュー ¶
GitHub API 経由で差分を取得し、LLM で確認して、コメントを返信します。すべて組み込みのノードを使用し、追加のサービスは必要ありません。
名前 : pr_review
description : LLM を使用して PR の差分を確認し、問題が見つかった場合はコメントを投稿します。
# 実行: PR=42 GITHUB_REPO=my-org/my-repo stepyard run pr_review
手順:
- ID : 差分
使用: http.request
付き:
URL : https://api.github.com/repos/${{ env.GITHUB_R

EPO }}/pulls/${{ env.PR }}/files
ヘッダー:
認可 : Bearer ${{ env.GITHUB_TOKEN }}
受け入れる: application/vnd.github+json
- ID : レビュー
使用: llm.generate
付き:
モデル：gpt-4o
システムプロンプト: |
あなたはコードレビューを行っている上級エンジニアです。
簡潔にしてください。すべて問題がないようであれば、正確に「LGTM」と返信してください。
プロンプト: |
このプル リクエストでバグ、セキュリティ問題、明らかな間違いがないか確認してください。
${{ ステップ.diff.output.body }}
- id : post_comment
if : ${{ ステップ.レビュー.出力.出力 != "LGTM" }}
使用: http.request
付き:
URL : https://api.github.com/repos/${{ env.GITHUB_REPO }}/issues/${{ env.PR }}/comments
メソッド: POST
ヘッダー:
認可 : Bearer ${{ env.GITHUB_TOKEN }}
受け入れる: application/vnd.github+json
json_body :
本文 : ${{ ステップ.レビュー.出力.出力 }}
ロールバックを使用した複数環境のデプロイメント ¶
名前 : デプロイ_マルチ
手順:
- id : デプロイ_ステージング
使用:shell.run
付き:
コマンド: kubectl apply -f k8s/staging/
- id : 統合テスト
使用:shell.run
continue_on_error : true
付き:
コマンド: pytest テスト/統合/ -q
- id : デプロイ_プロダクション
if : ${{ ステップ.integration_tests.output.code == 0 }}
使用:shell.run
付き:
コマンド: kubectl apply -f k8s/production/
- ID : ロールバック
if : ${{ ステップ.integration_tests.output.code != 0 }}
使用:shell.run
付き:
コマンド: kubectl rollout undodeployment/myapp -n staging
ドキュメントの概要 ¶

## Original Extract

Developer-first automation runner - pipelines as YAML, extended with Python, executed from the CLI or a local scheduler.

Stepyard
Skip to content
Stepyard
Home
Initializing search
rorlikowski/stepyard
Home
Stepyard
rorlikowski/stepyard
Home
Home
Table of contents
Why Stepyard?
Real-world examples
Daily database backup
Automated code review on every PR
Multi-environment deployment with rollback
Getting Started
Getting Started
Installation
Core Concepts
Core Concepts
Flows & Steps
How-to Guides
How-to Guides
Schedule flows
Built-in Nodes
Built-in Nodes
All nodes reference
Plugin Development
Plugin Development
Quick Start
Real-world examples
Daily database backup
Automated code review on every PR
Multi-environment deployment with rollback
YAML pipelines. Python plugins. Runs on your machine - no server to set up, no cloud account needed.
Stepyard is an automation runner for developers. Pipelines live as YAML files in your repo. You extend them with plain Python functions. Everything runs from the CLI or a local daemon - on your machine or a server you own.
pip install stepyard
stepyard init my-project && cd my-project
stepyard run hello
A deploy pipeline: build and push a container, run a smoke test, and post the result to Slack. One YAML file, no supporting scripts.
name : deploy
description : Build, push, and verify the production container.
steps :
- id : build
uses : shell.run
with :
command : docker build -t myapp:${{ env.GIT_SHA }} .
- id : push
uses : shell.run
with :
command : docker push myapp:${{ env.GIT_SHA }}
- id : smoke_test
uses : http.request
with :
url : https://staging.myapp.com/healthz
method : GET
- id : notify
uses : llm.generate # built-in - no plugin needed
with :
model : gpt-4o-mini
prompt : |
Summarise this deploy result in one sentence for Slack:
HTTP status: ${{ steps.smoke_test.output.status }}
SHA: ${{ env.GIT_SHA }}
- id : post_to_slack
uses : http.request
with :
url : ${{ env.SLACK_WEBHOOK }}
method : POST
json_body :
text : ${{ steps.notify.output.output }}
Run it:
GIT_SHA = $( git rev-parse --short HEAD ) stepyard run deploy
✓ build 12.4 s
✓ push 4.1 s
✓ smoke_test 0.3 s
✓ notify 0.9 s
✓ post_to_slack 0.2 s
Flow completed in 18.0 s
Why Stepyard? ¶
Steps, conditions, loops, and retries are plain YAML keys - no proprietary DSL. Version-control them alongside your code and validate them with stepyard validate .
One @node decorator turns any function into a reusable step. Inputs are type-validated automatically; plugin dependencies are isolated so they never conflict with Stepyard's own.
State is stored in a local SQLite database. Data goes out only if a step in your flow explicitly sends it.
Scheduled and on-demand execution
cron , interval , and startup triggers. Run stepyard service start and your flows execute on schedule - no external service or cloud account needed.
name : pg_backup
trigger :
uses : cron
with :
schedule : "0 3 * * *" # every day at 03:00
steps :
- id : dump
uses : shell.run
with :
command : pg_dump ${{ env.DATABASE_URL }} | gzip > /tmp/backup.sql.gz
- id : upload
uses : shell.run
with :
command : |
aws s3 cp /tmp/backup.sql.gz \
s3://${{ env.BACKUP_BUCKET }}/db/$(date +%Y-%m-%d).sql.gz
- id : cleanup
uses : shell.run
with :
command : rm -f /tmp/backup.sql.gz
Automated code review on every PR ¶
Fetch the diff via the GitHub API, review it with an LLM, and post a comment back - all with built-in nodes, no additional service.
name : pr_review
description : Review a PR diff with an LLM and post a comment if issues are found.
# Run: PR=42 GITHUB_REPO=my-org/my-repo stepyard run pr_review
steps :
- id : diff
uses : http.request
with :
url : https://api.github.com/repos/${{ env.GITHUB_REPO }}/pulls/${{ env.PR }}/files
headers :
Authorization : Bearer ${{ env.GITHUB_TOKEN }}
Accept : application/vnd.github+json
- id : review
uses : llm.generate
with :
model : gpt-4o
system_prompt : |
You are a senior engineer doing a code review.
Be concise. If everything looks good reply with exactly: LGTM
prompt : |
Review this pull request for bugs, security issues, and obvious mistakes:
${{ steps.diff.output.body }}
- id : post_comment
if : ${{ steps.review.output.output != "LGTM" }}
uses : http.request
with :
url : https://api.github.com/repos/${{ env.GITHUB_REPO }}/issues/${{ env.PR }}/comments
method : POST
headers :
Authorization : Bearer ${{ env.GITHUB_TOKEN }}
Accept : application/vnd.github+json
json_body :
body : ${{ steps.review.output.output }}
Multi-environment deployment with rollback ¶
name : deploy_multi
steps :
- id : deploy_staging
uses : shell.run
with :
command : kubectl apply -f k8s/staging/
- id : integration_tests
uses : shell.run
continue_on_error : true
with :
command : pytest tests/integration/ -q
- id : deploy_production
if : ${{ steps.integration_tests.output.code == 0 }}
uses : shell.run
with :
command : kubectl apply -f k8s/production/
- id : rollback
if : ${{ steps.integration_tests.output.code != 0 }}
uses : shell.run
with :
command : kubectl rollout undo deployment/myapp -n staging
Documentation overview ¶
