---
source: "https://kubesentry.io/"
hn_url: "https://news.ycombinator.com/item?id=49330052"
title: "Show HN: KubeSentry – Self-hosted runtime threat detection for Kubernetes"
article_title: "KubeSentry — self-hosted runtime threat detection for Kubernetes"
image: "https://kubesentry.io/og-image.png"
author: "KubeSentryio"
captured_at: "2026-08-17T13:33:14Z"
capture_tool: "hn-digest"
hn_id: 49330052
score: 3
comments: 0
posted_at: "2026-08-17T12:53:51Z"
tags:
  - hacker-news
  - translated
---

# Show HN: KubeSentry – Self-hosted runtime threat detection for Kubernetes

- HN: [49330052](https://news.ycombinator.com/item?id=49330052)
- Source: [kubesentry.io](https://kubesentry.io/)
- Score: 3
- Comments: 0
- Posted: 2026-08-17T12:53:51Z

## Translation

タイトル: Show HN: KubeSentry – Kubernetes のセルフホスト型ランタイム脅威検出
記事のタイトル: KubeSentry — Kubernetes のセルフホスト型ランタイム脅威検出
説明: Falco に基づいて構築されたカーネル レベルのランタイム脅威検出。完全に独自のクラスター内で実行されます。フォン ホームやテレメトリはありません。送信することを選択したアラート以外は何も残りません。一度支払えば、永久にあなたのものになります。

記事本文:
キューブセントリー
ドキュメント
価格設定
GitHub
インストール
現在利用可能 ·
7日間の無料トライアル
脅威をキャッチ
Kubernetes クラスター
彼らに捕まる前に。
Falco 上に構築され、完全に実行されるカーネル レベルのランタイム脅威検出
自分のクラスター内で。テレフォンホームもテレメトリーもありません - 何も残らない
ただし、選択した宛先に送信するよう指示したアラートは除きます。
Copy : 1 つの中に、テンプレート自身の改行を挟みます。
マップされた行は保持され、空白行として表示されます。 --> $ helm リポジトリ追加 kubesentry https://charts.kubesentry.io $ helm リポジトリ更新 $ helm install kubesentry kubesentry/kubesentry \ $ --namespace kubesentry --create-namespace
無料トライアルを始める →
ドキュメントを読む
数分でインストール可能、フル製品を 7 日間使用可能、クレジット カード不要
kubesentry.local
クラスター:production-eu
生きる
ナビゲーション
▸ アラート
12 ▸ 概要
▸ ルール
▸ 設定
クリティカル +1 2 高 +3 5 中 -4 12 重複排除の抑制 847 最近のアラート - 過去 24 時間
Helm インストールから最初のアラートまで
5分。
Falco + コレクター + ダッシュボード + アラーター — すべて 1 つの Helm 経由でデプロイされます
Falco がサポートするカーネルを実行しているクラスター上のチャート。
誤検知が少ないように調整された 14 の KubeSentry ルール - 仮想通貨マイナー、
リバースシェル、権限昇格、横方向の移動 - で実行中
Falco のデフォルトのルールセットの最上位。
重大なイベントは即座に発生します。毎日のダイジェストがノイズを巻き上げます。全14
KubeSentry ルール — および Falco の最も一般的なデフォルトの 8 つ — が同梱されています
トリアージ手順、修復コマンド、MITRE ATT&CK リファレンス。
ファルコは何が起こったのかを話します。
KubeSentry が何をすべきかを教えてくれます。
各カードの灰色のブロックは生の Falco イベントです。これが裸のイベントです。
Falco をインストールすると、午前 3 時にコンテキストが何も表示されません。周囲のものすべて
KubeSentry: 認識したポッドに解決されるフィールド。
検出は実際に意味する

、優先順位を付ける方法、およびその正確なコマンド
シャットダウンしてください。
🚨 KubeSentry: コンテナ内で変更された Sudoers または passwd
13:55:11.144843182: 重要な機密認証ファイルが書き込みのために開かれました (file=/etc/passwd command=sh -c echo 'svc:x:0:0::/root:/bin/sh' >> /etc/passwd process=sh pod=payments-api-7d8f ns=prodcontainer=e263ccbef60c)修正: kubectl delete pod -n <ns> <pod> --grace-period=0 --force
ポッドがアクセスできる資格情報をローテーションします。 readOnlyRootFilesystem を設定: true
KubeSentry: コンテナ内で変更された Sudoers または passwd
13:55:11.144843182: 重要な機密認証ファイルが書き込みのために開かれました (file=/etc/passwd command=sh -c echo 'svc:x:0:0::/root:/bin/sh' >> /etc/passwd process=sh pod=payments-api-7d8f ns=prodcontainer=e263ccbef60c)コンテナ ID=e263ccbef60c コンテナ名=payments-api コンテナイメージリポジトリ=ghcr.io/acme/payments-api コンテナイメージタグ=1.24.0 k8s_pod_name=payments-api-7d8f k8s_ns_name=prod ポッドpayments-api-7d8f 名前空間 prod コンテナ e263ccbef60c イメージghcr.io/acme/payments-api 時間 2026-07-27 13:55:11 UTC
それが何を意味するか
プロセスは、コンテナ内に書き込むために /etc/passwd、/etc/shadow、または /etc/sudoers を開きました。これは、アカウントの作成または権限の昇格を直接試みたものです。
1. これは実行時に正当であることはほとんどありません。書き込みプロセス (proc.cmdline) を特定します。
2. 新しいユーザー/sudo エントリを確認します: kubectl exec -n <ns> <pod> -- cat /etc/passwd /etc/sudoers
3. 同じポッドからの他のアラートと関連付けます (偵察、リバース シェル)。
https:// Attack.mitre.org/techniques/T1136/
URL を設定するとチャンネルがオンになります — Discord と Teams は独自のネイティブを取得します
ペイロード、最小公倍数 d ではない

エノミネーターのフォールバックとカスタム Webhook
アラート全体を生の JSON として取得します。インスタント アラートがすべてのチャネルにファンアウトされる
並行して構成しました。毎日のダイジェストは電子メールのみです。
14 個の KubeSentry ルールすべてと、Falco の最も一般的なデフォルトの 8 個が同梱されています。
トリアージ手順、修復コマンド、MITRE ATT&CK リファレンスを提供
ランブックに埋め込むのではなく、アラート自体に含めます。
すべてがクラスター内で実行されます。
1 つの Helm チャートは、すべてのノード、ルーター、および単一のノードに Falco DaemonSet をインストールします。
コレクターのデプロイメント — これが表示されるすべてのポッドです。検出、保存、および
ダッシュボードは決して私たちに話しかけません。残るのはあなたへの警告だけです
明示的に設定されており、選択した宛先に移動します。
DaemonSet — ノードごとに 1 つの Falco ポッド、特権付き、システムコールの監視
これが唯一残るものです。向きを変えたら
通知がオンの場合、アラートは設定した宛先に送信されますが、どこにも送信されません
それ以外は。この写真には KubeSentry サーバーはありません。テレメトリもテレフォンホームもありません。
ライセンスの場合でも、コンパイルされたキーに対してオフラインで検証されます。
バイナリ。クラスターをエアギャップすると、これらの通知を除くすべてが引き続き機能します。
エンタープライズセキュリティ
企業価格で。
99ユーロで。
データドッグ。シスディグ。アクア。セキュリティ チームを持つ企業向けに構築され、次のように販売されます。
サブスクリプションは、実行した量に応じて測定され、実行している限り継続します。
KubeSentry は、残りの私たち、つまり 5 ～ 50 人のエンジニアのために構築されています。
組織には実際のクラスタがあり、フルタイムの仕事がセキュリティである人は誰もいません。
一度お支払いいただきます。クラスター内で実行されます。誰もあなたのホストを測定しません。
1回払い。永遠にあなたのものです。すべての層には以下が含まれます
今後のすべてのアップデート - サブスクリプションなし、なし
リニューアル、ずっと。
7 日間の無料トライアル · クレジット カードなし
個人および非常に小さなクラスター。
✓
厳選されたルールセット + 救済策

イオンに関するアドバイス
✓
電子メール、Slack、Discord、Teams、カスタム Webhook
少数のクラスターを実行する小規模なチーム。
✓
厳選されたルールセット + 修復アドバイス
✓
電子メール、Slack、Discord、Teams、カスタム Webhook
✓
厳選されたルールセット + 修復アドバイス
✓
電子メール、Slack、Discord、Teams、カスタム Webhook
クラスター数はライセンス条件であり、技術的な制限ではありません。
コレクターには、実行しているクラスターの数をチェックする機能はありません。これを強制すると、
家に電話するという意味ですが、決してしないと約束します。したがって、私たちはあなたが次の層を購入することを信頼しています
が適合し、次の方法でその約束を検証できます。
コードを読んでいます。
チェックアウトと請求書発行は、Lemon Squeezy (当社の登録販売者) によって処理されます。 VAT
チェックアウト時に計算されます。ライセンスキーは自動的に電子メールで送信されます
購入後 — --set License=… を使用して適用します。
まだ購入したくないですか？まず 7 日間のトライアル版をインストールして実行します。それは完全な製品です。
キューブセントリー
Kubernetes のセルフホスト型ランタイム脅威検出。

## Original Extract

Kernel-level runtime threat detection built on Falco, running entirely in your own cluster — no phone-home, no telemetry. Nothing leaves except the alerts you choose to send. Pay once, yours forever.

kube sentry
Docs
Pricing
GitHub
Install
Available now ·
7-day free trial
Catch threats in your
Kubernetes cluster
before they catch you.
Kernel-level runtime threat detection built on Falco, running entirely
in your own cluster. No phone-home, no telemetry — nothing leaves
except the alerts you tell it to send, to the destinations you pick.
Copy : inside one, the template's own newlines between
the mapped lines are preserved and render as blank rows. --> $ helm repo add kubesentry https://charts.kubesentry.io $ helm repo update $ helm install kubesentry kubesentry/kubesentry \ $ --namespace kubesentry --create-namespace
Start free trial →
Read the docs
Installs in minutes · full product for 7 days · no credit card
kubesentry.local
cluster: production-eu
live
Navigation
▸ Alerts
12 ▸ Overview
▸ Rules
▸ Settings
Critical +1 2 High +3 5 Medium −4 12 Suppressed dedup 847 Recent alerts — last 24h
From helm install to first alert in under
5 minutes.
Falco + collector + dashboard + alerter — all deployed via one Helm
chart, on any cluster running a kernel Falco supports.
14 KubeSentry rules tuned for low false positives — crypto miners,
reverse shells, privilege escalation, lateral movement — running on
top of Falco's default ruleset.
Critical events fire instantly. Daily digests roll up the noise. All 14
KubeSentry rules — plus 8 of Falco's most common defaults — ship with
triage steps, remediation commands, and a MITRE ATT&CK reference.
Falco tells you what happened.
KubeSentry tells you what to do.
The grey block in each card is the raw Falco event — that's what a bare
Falco install gives you, at 3am, with no context. Everything around it
is KubeSentry: the fields resolved to a pod you recognise, what the
detection actually means, how to triage it, and the exact commands to
shut it down.
🚨 KubeSentry: Sudoers or passwd modified in container
13:55:11.144843182: Critical Sensitive auth file opened for writing (file=/etc/passwd command=sh -c echo 'svc:x:0:0::/root:/bin/sh' >> /etc/passwd process=sh pod=payments-api-7d8f ns=prod container=e263ccbef60c) container_id=e263ccbef60c container_name=payments-api container_image_repository=ghcr.io/acme/payments-api container_image_tag=1.24.0 k8s_pod_name=payments-api-7d8f k8s_ns_name=prod Fix: kubectl delete pod -n <ns> <pod> --grace-period=0 --force
Rotate any credentials the pod could access; set readOnlyRootFilesystem: true
KubeSentry: Sudoers or passwd modified in container
13:55:11.144843182: Critical Sensitive auth file opened for writing (file=/etc/passwd command=sh -c echo 'svc:x:0:0::/root:/bin/sh' >> /etc/passwd process=sh pod=payments-api-7d8f ns=prod container=e263ccbef60c) container_id=e263ccbef60c container_name=payments-api container_image_repository=ghcr.io/acme/payments-api container_image_tag=1.24.0 k8s_pod_name=payments-api-7d8f k8s_ns_name=prod Pod payments-api-7d8f Namespace prod Container e263ccbef60c Image ghcr.io/acme/payments-api Time 2026-07-27 13:55:11 UTC
What it means
A process opened /etc/passwd, /etc/shadow, or /etc/sudoers for writing inside a container — a direct attempt to create accounts or escalate privileges.
1. This is almost never legitimate at runtime; identify the writing process (proc.cmdline).
2. Check for new users/sudo entries: kubectl exec -n <ns> <pod> -- cat /etc/passwd /etc/sudoers
3. Correlate with other alerts from the same pod (recon, reverse shell).
https://attack.mitre.org/techniques/T1136/
Set a URL and the channel turns on — Discord and Teams get their own native
payloads, not a lowest-common-denominator fallback, and the custom webhook
gets the whole alert as raw JSON. Instant alerts fan out to every channel
you've configured, in parallel; the daily digest is email-only.
All 14 KubeSentry rules — plus 8 of Falco's most common defaults — ship with
triage steps, remediation commands, and a MITRE ATT&CK reference, delivered
in the alert itself rather than buried in a runbook.
Everything runs inside your cluster .
One Helm chart installs a Falco DaemonSet on every node, a router, and a single
collector Deployment — that's every pod you'll see. Detection, storage, and the
dashboard never talk to us. The only thing that ever leaves is an alert you
explicitly configured, going to a destination you picked.
DaemonSet — one Falco pod per node, privileged, watching syscalls
This is the only thing that ever leaves. If you turn
notifications on, the alert goes to the destination you configured — and nowhere
else. There is no KubeSentry server in the picture: no telemetry, no phone-home,
not even for licensing, which is verified offline against a key compiled into the
binary. Air-gap the cluster and everything except these notifications still works.
Enterprise security
at enterprise prices.
for €99.
Datadog. Sysdig. Aqua. Built for enterprises with security teams, and sold as
subscriptions metered by how much you run — for as long as you run it.
KubeSentry is built for the rest of us: the five-to-fifty-person engineering
org with real clusters and nobody whose full-time job is security.
You pay once. It runs in your cluster. Nobody meters your hosts.
One-time payment. Yours forever. Every tier includes
all future updates — no subscription, no
renewals, ever.
7-day free trial · no credit card
Individuals and very small clusters.
✓
Curated ruleset + remediation advice
✓
Email, Slack, Discord, Teams, custom webhook
Small teams running a handful of clusters.
✓
Curated ruleset + remediation advice
✓
Email, Slack, Discord, Teams, custom webhook
✓
Curated ruleset + remediation advice
✓
Email, Slack, Discord, Teams, custom webhook
Cluster counts are licence terms, not technical limits.
Nothing in the collector checks how many clusters you run — enforcing that would
mean phoning home, and we promise never to. So we trust you to buy the tier that
fits, and you get to verify that promise by
reading the code .
Checkout and invoicing are handled by Lemon Squeezy (our merchant of record); VAT
is calculated at checkout. Your license key is emailed to you automatically right
after purchase — apply it with --set license=… .
Don't want to buy yet? Install and run the 7-day trial first; it's the full product.
kube sentry
Self-hosted runtime threat detection for Kubernetes.
