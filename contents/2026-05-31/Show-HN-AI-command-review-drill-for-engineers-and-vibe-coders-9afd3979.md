---
source: "https://www.proreview.dev/"
hn_url: "https://news.ycombinator.com/item?id=48340994"
title: "Show HN: AI command review drill for engineers and vibe coders"
article_title: "ProReview - Catch AI Before It Wrecks Production"
author: "shaad1337"
captured_at: "2026-05-31T00:24:33Z"
capture_tool: "hn-digest"
hn_id: 48340994
score: 1
comments: 0
posted_at: "2026-05-30T21:56:47Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI command review drill for engineers and vibe coders

- HN: [48340994](https://news.ycombinator.com/item?id=48340994)
- Source: [www.proreview.dev](https://www.proreview.dev/)
- Score: 1
- Comments: 0
- Posted: 2026-05-30T21:56:47Z

## Translation

タイトル: Show HN: エンジニアとバイブコーダーのための AI コマンドレビュードリル
記事のタイトル: ProReview - 生産を台無しにする前に AI を捕まえる
説明: AI によって生成されたコード、インフラストラクチャ、SQL、CI/CD、セキュリティ アーティファクトを出荷前にレビューする練習をします。ブラインド判定、セグメントレベルのスコアリング、および専門家の明らかに。

記事本文:
ProReview - 本番環境を破壊する前に AI をキャッチ ProReview の仕組み 機能トラック ProReview のレビューを開始して、本番環境に到達する前に AI のリスクを特定できるようにエンジニアリング チームをトレーニングします。
この AI コマンドを実行しますか?
AI によって生成された bash、kubectl、およびシェルのワンライナーを実稼働環境に触れる前に確認します。電話をかけ、危険なセグメントにフラグを立てて、何を見逃したかを確認します。
bash -c "set -euo パイプ失敗
MOUNT = /var/lib/postgresql
DEV =$( python3 -c 'インポート os, サブプロセス; src=subprocess.check_output(["findmnt","-n","-o","SOURCE","/var/lib/postgresql"], text=True).strip(); print(os.path.realpath(src))' )
sudo umount -l "$MOUNT" ||本当の
sudo winefs -a "$DEV" && sudo mkfs.ext4 -F "$DEV"
sudo mount "$DEV" "$MOUNT" " このコマンドを承認しますか? Y / N 50 + チャレンジ 5 トラック 0 AI ヒント レビューあたり 3 分未満 仕組み プロンプトから本番レビューまで
エンジニアが毎日出荷する AI 生成のアーティファクトを中心に設計された集中訓練。
Kubernetes、クラウド インフラストラクチャ、SQL、CI/CD、バックエンド セキュリティ サーフェスから選択します。
AI が生成したアーティファクトのコールドを読み取ります。説明を見る前に安全かどうかを判断してください。
運用を中断する可能性がある正確なコマンド、構成、SQL、またはコード セグメントを選択します。
自分の判断を専門家の回答キーと比較し、次のレビューのパターンを学びます。
毛羽立ちはありません。 AI コードレビューの本能を研ぎ澄ますために必要なループです。
生成されたコマンド、差分、構成、移行、エンジニアが実際に出荷した作業に似た API コードを確認します。
ヒントが表示される前に、安全か安全でないかを判断します。 LLM の手持ち、ネタバレ、または答え優先のトレーニングはありません。
判決だけでなく、危険なラインを選択してください。あなたが発見した生産上の危険を正確に評価してください。
通話を、障害モード、爆発範囲、より安全なレビュー結果の精選された内訳と比較します。
練習する

Kubernetes 運用、クラウド インフラストラクチャ、データ移行、CI/CD、およびセキュリティに敏感なバックエンド コード。
AI がほぼ出荷したものを把握した後、スコアと概要を共有して実行を終了します。
5つの面。失敗するための50の方法。
各トラックは、制作を中断する正確なアーティファクト タイプをカバーしています。
# メンテナンス.sh
kubectl ノードを取得する
kubectl 非常線ワーカーノード-03
kubectl ドレイン ワーカーノード-03 \
--ignore-daemonsets \
--強制\
--ローカルデータの削除
kubectl delete node worker-node-03 12 の課題 HARD TRACK 02 データベースと SQL # billing-cleanup.sql
psql -h db.internal \
-d billing_prod \
-c '開始;'
請求書から削除
WHERE ステータス = ドラフト;
invoice_items から削除;
専念;' 10 のチャレンジ MED TRACK 03 CI / CD #deploy.yml
- 名前 : 更新イメージ
実行: |
DEPLOY_CMD= "kubectl セットイメージ
デプロイメント/${{ github.event.inputs.service }}
...」
eval $DEPLOY_CMD 11 の課題 HARD TRACK 04 インフラストラクチャ # backend.tf
テラフォーム {
バックエンド「s3」{
バケット = "会社-terraform-state"
key = "app/terraform.tfstate"
地域 = "us-east-1"
}
9 つの課題 MED TRACK 05 Python と API #exports.py
@router.get ( "/exports/download" )
async def download_export(
ファイル名: str 、
admin=依存(require_admin)
):
file_path = f "/var/exports/{ファイル名}"
return FileResponse(file_path) 10 のチャレンジ MED レベルアップする準備はできましたか?次の本番インシデントはすでに書かれています
1 日に 1 つの課題をレビューし、デプロイ前にバグを発見するエンジニアに加わりましょう。

## Original Extract

Practice reviewing AI-generated code, infrastructure, SQL, CI/CD, and security artifacts before they ship. Blind judgment, segment-level scoring, and expert reveals.

ProReview - Catch AI Before It Wrecks Production ProReview How it works Features Tracks Start reviewing ProReview Training engineering teams to spot AI risks before they reach production.
Would you run this AI command ?
Review AI-generated bash, kubectl, and shell one-liners before they touch production. Make the call, flag the risky segment, and see what you missed.
bash -c "set -euo pipefail
MOUNT = /var/lib/postgresql
DEV =$( python3 -c 'import os, subprocess; src=subprocess.check_output(["findmnt","-n","-o","SOURCE","/var/lib/postgresql"], text=True).strip(); print(os.path.realpath(src))' )
sudo umount -l "$MOUNT" || true
sudo wipefs -a "$DEV" && sudo mkfs.ext4 -F "$DEV"
sudo mount "$DEV" "$MOUNT" " Approve this command? Y / N 50 + Challenges 5 Tracks 0 AI hints <3 min Per review How it works From prompt to production review
A focused drill designed around the AI-generated artifacts engineers ship every day.
Choose from Kubernetes, cloud infrastructure, SQL, CI/CD, and backend security surfaces.
Read the AI-generated artifact cold. Decide whether it is safe before seeing any explanation.
Select the exact commands, config, SQL, or code segments that would break production.
Compare your verdict with the expert answer key and learn the pattern for the next review.
No fluff. Just the loop you need to sharpen AI code review instincts.
Review generated commands, diffs, configs, migrations, and API code shaped like work engineers actually ship.
Commit to safe or unsafe before any hints appear. No LLM hand-holding, spoilers, or answer-first training.
Select the risky lines, not just the verdict. Get credit for the exact production hazard you caught.
Compare your call with a curated breakdown of the failure mode, blast radius, and safer review outcome.
Practice Kubernetes ops, cloud infrastructure, data migrations, CI/CD, and security-sensitive backend code.
Finish a run with a score and summary you can share after catching what AI almost shipped.
Five surfaces. Fifty ways to fail.
Each track covers the exact artifact types that break production.
# maintenance.sh
kubectl get nodes
kubectl cordon worker-node-03
kubectl drain worker-node-03 \
--ignore-daemonsets \
--force \
--delete-local-data
kubectl delete node worker-node-03 12 challenges HARD TRACK 02 Database & SQL # billing-cleanup.sql
psql -h db.internal \
-d billing_prod \
-c 'BEGIN;'
DELETE FROM invoices
WHERE status = draft;
DELETE FROM invoice_items;
COMMIT;' 10 challenges MED TRACK 03 CI / CD # deploy.yml
- name : Update image
run : |
DEPLOY_CMD= "kubectl set image
deployment/${{ github.event.inputs.service }}
..."
eval $DEPLOY_CMD 11 challenges HARD TRACK 04 Infrastructure # backend.tf
terraform {
backend "s3" {
bucket = "company-terraform-state"
key = "app/terraform.tfstate"
region = "us-east-1"
}
} 9 challenges MED TRACK 05 Python & APIs # exports.py
@router.get ( "/exports/download" )
async def download_export(
filename: str ,
admin=Depends(require_admin)
):
file_path = f "/var/exports/{filename}"
return FileResponse(file_path) 10 challenges MED Ready to level up? Your next production incident is already written
Join engineers who review one challenge a day and catch the bugs before deploy time.
