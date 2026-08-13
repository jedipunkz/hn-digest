---
source: "https://www.cncf.io/blog/2026/03/17/when-kubernetes-restarts-your-pod-and-when-it-doesnt/"
hn_url: "https://news.ycombinator.com/item?id=49284256"
title: "When Kubernetes restarts your pod – And when it doesn't"
article_title: "When Kubernetes restarts your pod — And when it doesn’t | CNCF"
author: "trikelef"
captured_at: "2026-08-13T11:39:33Z"
capture_tool: "hn-digest"
hn_id: 49284256
score: 1
comments: 0
posted_at: "2026-08-13T11:09:50Z"
tags:
  - hacker-news
  - translated
---

# When Kubernetes restarts your pod – And when it doesn't

- HN: [49284256](https://news.ycombinator.com/item?id=49284256)
- Source: [www.cncf.io](https://www.cncf.io/blog/2026/03/17/when-kubernetes-restarts-your-pod-and-when-it-doesnt/)
- Score: 1
- Comments: 0
- Posted: 2026-08-13T11:09:50Z

## Translation

タイトル: Kubernetes がポッドを再起動するとき – そして再起動しないとき
記事のタイトル: Kubernetes がポッドを再起動するとき — そして再起動しないとき | CNCF
説明: Kubernetes 1.35 GA Companion リポジトリに対して検証された運用内部ガイド: github.com/opscart/k8s-pod-restart-mechanics エンジニアは「ポッド…

記事本文:
コンテンツにスキップ
アクセシビリティ
助けて
について
CNCFを発見する
メンバーシップ ハブ – 既存メンバー向け CNCF メンバーになることのすべての利点について学びます
技術監視委員会 TOC は CNCF の技術的ビジョンを定義し、クラウド ネイティブ コミュニティに経験豊富な技術的リーダーシップを提供します。
理事会 GB は CNCF のマーケティング、ビジネス監視、予算決定を担当します。
エンド ユーザー技術諮問委員会 エンド ユーザー TAB は、CNCF コミュニティの意思決定におけるエンド ユーザーの声として機能します。
アンバサダー アンバサダーをご紹介します。他の人がクラウド ネイティブ テクノロジーについて学ぶのを助けることに熱心な経験豊富な実践者です。
段階的プロジェクトは安定しており、広く採用されており、運用準備が整っていると考えられており、何千人もの貢献者を集めています。
健全な貢献者プールを持つ少数のユーザーによって実稼働環境で成功裏に使用されるプロジェクトのインキュベーション
サンドボックス 最先端のテクノロジーにおける運用環境でまだ広くテストされていない実験プロジェクト
ライフサイクルの終わりに達し、非アクティブになったアーカイブ済みプロジェクト
プロジェクトのメトリクス 成熟度レベルを段階的に移行する CNCF プロジェクトのメトリクスを表示します
CNCF がホストするプロジェクトに専門知識を提供した #TeamCloudNative の 15 万人以上の人々に参加しましょう
CNCF プロジェクト向けサービス オープンソース プロジェクト向けの CNCF サービス - マーケティングから法務サービスまで
クラウド ネイティブ ランドスケープ クラウド ネイティブ領域におけるプロジェクトと製品の包括的なカテゴリ概要
プロジェクトジャーニーレポート CNCF がさまざまな卒業プロジェクトの進捗と成長にどのような影響を与えたかを示します
プロジェクト ツール CNCF プロジェクトのツールとリソースへのクイック リンク
最新のプロジェクトジャーニーレポート
トレーニングの概要 ピオーネのトレーニングと証明書でクラウド ネイティブの成功への道を見つけてください

クラウドネイティブテクノロジーの専門家
認定業界で認められたクラウドネイティブ認定の唯一の信頼できるソースから認定を取得
コース Linux Foundation と協力して構築されたトレーニング コースで、クラウド ネイティブの基礎を学習したり、認定資格のトレーニングを行ったりできます。
Kubestronaut プログラム Kubernetes スキルを向上させる
Kubernetes トレーニング パートナー 次の認定資格の準備に適した KCNTP を見つけてください
認定された Kubernetes ソフトウェア準拠により、CNCF プロジェクトのバージョンが必要な API を確実にサポートします。
クラウド ネイティブ ネットワーク機能認定 (ベータ) CNF 認定により、アプリケーションがクラウド ネイティブのベスト プラクティスを実証することが保証されます。
Kubernetes 認定サービス プロバイダー KCSP は、企業によるクラウド ネイティブ テクノロジーの導入を支援する豊富な経験を持っています
あなたの会社を CNCF エンド ユーザーとして登録し、トレーニングとカンファレンスのコストを 10,000 ドル以上節約しましょう
エンド ユーザー コミュニティ クラウド ネイティブ テクノロジを使用して製品とサービスを構築する、ベンダー中立のコミュニティに参加してください。
イベント 世界中のイベントで #TeamCloudNative と CNCF スタッフに会いましょう
ケース スタディ クラウド ネイティブ プロジェクトが世界中の組織に与えている影響について、実際のケース スタディを読む
Human of Cloud Native 素晴らしい個人とその貢献のストーリーを読む
クラウド ネイティブ ヒーローのチャレンジ クラウド ネイティブ ヒーローになろう!特許トロールを倒して盗品や賞品を獲得するのにご協力ください
オンライン プログラム クラウド ネイティブ テクノロジーとプロジェクトに関する最新の洞察を得るには、無料のオンライン プログラムをご覧ください。
コミュニティ グループ お近くのイベントやミートアップに #TeamCloudNative に参加してください
Phippy & Friends Phippy は、あらゆる年齢層に最適なストーリーを通じて、クラウド ネイティブの核となる概念を簡単に説明します
クラウド ネイティブ用語集 技術的な知識がなくても、クラウド ネイティブの概念を明確でシンプルな言語で理解できます。

dgeが必要です！
ブログ #TeamCloudNative から最新の出来事や技術的な洞察を入手してください
お知らせ メディアリリースとCNCFの公式発表
メディアでの CNCF プロジェクトと #TeamCloudNative のニュース
レポート 私たちの組織、イベント、プロジェクトに関する透明性の高い詳細なレポートを読むことができます。
検索
CNCF
ブログ
/
プロジェクト管理者ポスト
Kubernetes がポッドを再起動するときと再起動しないとき
投稿日: 2026 年 3 月 17 日
Shamsher Khan、プロジェクト管理者による
この投稿で取り上げられている CNCF プロジェクト
Kubernetes 1.35 GA に対して検証された運用内部ガイド
コンパニオン リポジトリ: github.com/opscart/k8s-pod-restart-mechanics
エンジニアは「ポッドが再起動した」と言うのは 4 つの異なる意味です。これを間違えると、運用手順書に不備が生じたり、オンコールでの不適切な意思決定が発生したりすることになります。
実践的なテスト: ポッド UID は変更されましたか? 「はい」の場合、それは再作成であり、コンテナーの再起動ではありません。再起動回数はゼロにリセットされます。 「いいえ」の場合、同じポッド オブジェクトであり、その内部でコンテナ プロセスが再起動されます。
核となる洞察: kubelet が実際に何を監視しているか
kubelet は、ConfigMap、Secret、Istio CRD ではなく、ポッドの仕様を監視します。ポッドの仕様が変更されていない場合、kubelet は起動しません。この 1 つの事実で、「なぜ設定が更新されなかったのか?」の大部分が説明されます。生産中の調査。
アドミッション Webhook を変更すると、作成時にポッドの仕様を変更できますが、アドミッション後は変更できません。作成後にコンテナーの再起動をトリガーすることはできません。
以下のフローチャートは、同じマトリックスを、インシデント発生中の午前 2 時にたどることができる意思決定パスに変換します。
図 1: 完全な決定フローチャート — この変更にはポッドの再起動が必要ですか?
シナリオ 1: ConfigMap — 同じ変更に 2 つの動作がある理由
[図 1: ConfigMap 環境変数とボリューム マウント — 環境変数ポッドは凍結され、ボリューム ポッドは kubelet シンボリックリンク スワップ経由で自動同期される]
環境変数

mode (envFrom / valueFrom) : カーネルは execve() で環境変数を /proc/<pid>/environ にコピーします。そのメモリはプロセスによって所有されており、外部システムはそれを変更できません。 ConfigMap を更新しても、kubelet はポッド仕様の変更を認識せず、何も行いません。このプロセスでは古い値が無期限に保持されます。
ボリューム マウント モード: kubelet は、ファイルへの書き込みではなく、アトミック シンボリック リンク スワップを介して同期します。
/etc/config/
§── ..2025_12_19_11_30_00/ ← 新しいデータディレクトリ (kubelet がこれを作成します)
│ └── APP_COLOR ← 「赤」
§── ..data ─────────────▶ ..2025_12_19_11_30_00/ ← シンボリックリンクがアトミックにSWAPPED
└── APP_COLOR ───────── ▶ ..data/APP_COLOR
シンボリックリンク スワップでは、..data に対して IN_CREATE が生成されますが、ファイルに対しては IN_MODIFY が生成されません。開いているファイル記述子の IN_MODIFY を監視しているアプリケーションは、これを完全に見逃しています。これが、明示的な inotify 処理を行わない限り、nginx が ConfigMap の変更時に自動リロードを行わない理由です。
ラボの証拠 (コンパニオン リポジトリ内の 01-configmap/)
ConfigMap 更新: APP_COLOR 青 → 赤
ポッド A (環境変数): APP_COLOR=blue ← 凍結、再起動回数: 0
ポッド B (ボリューム マウント): APP_COLOR=red ← 自動同期、再起動回数: 0
正しい inotify パターン — ファイルではなくディレクトリを監視します
watcher.Add(filepath.Dir(configPath)) // /etc/config/ を監視します — IN_CREATE をキャッチします
// watcher.Add(configPath) // シンボリックリンクのスワップを完全に見逃します
イベントの場合 := range watcher.Events {
ifevent.Op&fsnotify.Create == fsnotify.Create {
reloadConfig()
}
}
シナリオ 2: イメージの更新 — 再作成 vs コンテナーの再起動 vs CrashLoop
これら 3 つのシナリオは似ていますが、根本的に異なります。
イメージの更新が成功しました — ポッドが再作成されました
変更前: ポッド UID aaa-bbb、IP 10.244.1.5、nginx:1.25、再起動: 0
変更後: ポッド UID xxx-yyy、IP 10.244.1.6、nginx:1.27、再起動: 0
↑UIDチャンネル

怒っている — コンテナの再起動ではなく、レクリエーション
図 3: 新しい ReplicaSet の作成、ポッドの再作成、ロールバックのために保持される古い RS を示すローリング アップデート フロー。
ImagePullBackOff — 古いポッドは保護されたままになります
古いポッド: 実行中 ← Kubernetes により存続
新しいポッド: ImagePullBackOff ← スタックし、新しいポッドが正常になるまで古いポッドが強制終了されることはありません
CrashLoopBackOff — 同じポッド、再起動回数が増加
ポッド UID: aaa-bbb ← 未変更
再起動回数: 0 → 1 → 2 → 3 ← 同じポッド オブジェクト、コンテナーのクラッシュ
診断ルール: UID が変更されていない登山再開回数 = クラッシュ ループ。新しい UID によるゼロ再起動カウント = ローリング アップデート。
StatefulSet に関する注意: StatefulSet ポッドはイメージ変更時にも再作成されますが、順序 ID (pod-0、pod-1) と PVC バインディングは保持されます。コンテナーの再起動セマンティクスはデプロイメントと同じです。ID の永続性はコンテナーのインプレース再起動を意味しません。
シナリオ 3: インプレースリソースのサイズ変更 (K8s 1.35 GA)
K8s 1.35 では、ポッドのインプレース サイズ変更が一般提供されるようになりました (kubernetes.io/blog/2025/12/19/kubernetes-v1-35-in-place-pod-resize-ga)。 CPU とメモリはどちらもポッドを再作成せずにサイズ変更できます。UID と IP は決して変更されません。インプレース サイズ変更が利用できるかどうかは、CRI とノード OS のサポートによって異なります。 Linux cgroups v2 を使用して、containerd 1.7 以降で検証済み。
コンテナーに何が起こるかは、明示的に定義するsizePolicyによって異なります。
サイズ変更ポリシー:
- リソース名: CPU
restartPolicy: NotRequired # cgroup クォータが更新されましたが、プロセスは変更されていません
- リソース名: メモリ
restartPolicy: RestartContainer # 同じポッド内でコンテナが再起動されます
ラボの証拠 (05-resource-resize/ — K8s 1.35+ が必要)
CPU サイズ変更 200m → 500m (必須ではありません):
UID: d7c99204 IP: 10.244.0.7 再起動: 0 ← すべて変更なし
メモリのサイズ変更 256Mi → 512Mi (RestartContainer):
UID: d7c99204 IP: 10.244.0.7 再起動: 1
↑同じポッド

↑ 同じ IP ↑ K8 による強制ではなく、私たちのポリシー選択
重要: メモリのデフォルトのsizePolicyはNotRequiredです。これを省略した場合、メモリのサイズ変更はコンテナを再起動せずにサイレントに cgroup を更新し、JVM ヒープは古いサイズのままになります。メモリに対しては常に、resizePolicy を明示的に定義します。
kubectl パッチ ポッド my-pod -n my-namespace \
--サブリソースのサイズ変更 \
-p '{"仕様":{"コンテナ":[{"名前":"アプリ","リソース":{
"リクエスト":{"cpu":"250m","メモリ":"128Mi"},
"制限":{"cpu":"500m","メモリ":"256Mi"}
}}]}}'
# 注: --type=merge を省略すると、 --subresourcesize を使用すると検証エラーが発生します
シナリオ 4: Istio ルーティング — xDS 経由のゼロ再起動
Istio VirtualService と DestinationRule の変更によってコンテナーの再起動がトリガーされることはありません。 Istiod は、各 Envoy サイドカーへの永続的な双方向 gRPC ストリームを維持します。ルーティングの更新はミリ秒単位でプッシュされ、メモリ内でスワップされ、ポッドのタッチやファイルの書き込みはありません。
ラボの証拠 (04-istio-routing/ コンパニオン リポジトリ内)
ポッドは 4 つ。 3 つのルーティングの変更:
100% v1 → 80/20 カナリア → 100% v2
再起動回数: BEFORE 0 0 0 0 / AFTER 0 0 0 0
ポッドの年齢: 3 つの変更すべてを通じて変更されません。
シナリオ 5: Stakater リローダー — 手動ステップの自動化
アプリが環境変数を介して ConfigMap を使用する場合、更新のたびに誰かが kubectl rollout restart を実行する必要があります。 Reloader は、Kubernetes 監視イベントを使用してこれを自動化します。検出はポーリングではなく、ほぼ瞬時に行われます。
メタデータ:
注釈:
reloader.stakater.com/auto: "true"
運用上の注意事項: デフォルトの Helm インストールでは watchGlobally=false が使用されます。リローダーは独自の名前空間のみを監視します。他の名前空間の注釈付きデプロイメントは警告なしに無視され、エラーはスローされません。常に watchGlobally=true を指定してインストールします。
Helm インストール リローダー stakater/reloader \
--namespace リローダー \
--set reloader.watchGlobally=true
研究室の証拠 (07-s

takater-reloader/ コンパニオン リポジトリ内)
構成マップが更新されました。 kubectl ロールアウトの再起動実行はありません。
新しいポッド APP_MESSAGE: 「OpsCart v2 からこんにちは — 自動リロードされました!」
ローリングコンテナの再起動が自動的にトリガーされます。
ホットリロードが失敗する場合
ホットリロードは、コンテナーの再起動よりも常に安全であるとは限りません。知っておく価値のある 2 つの障害モード:
意味的に無効な構成がサイレントに受け入れられました
ファイルが更新され、inotify ハンドラーが起動され、エラーはスローされませんが、新しい構成には論理エラーがあります。ポッドはヘルスチェックに合格しましたが、何時間も壊れた状態で動作します。不適切な構成でコンテナーを再起動すると、すぐに大音量で失敗します。不適切な設定を使用したホットリロードは、静かに遅く失敗します。
軽減策: アトミックにスワップする前に構成を検証します。
Envoy が xDS プッシュをサイレントに拒否する
Istiod は、まだ伝播されていないクラスターを参照する RouteConfiguration をプッシュします。 Envoy はそれを拒否し、古いルーティング ルールを続行します。ポッドイベントは発生しません。軽減策:pilot_xds_push_errors を監視し、istioctl proxy-status を使用します。
可観測性: すべてのオペレータが知っておくべき 3 つのコマンド
# 1. コンテナーの再起動またはポッドの再作成? UIDの変更を確認する
kubectl get ポッド <pod> -o Custom-columns=\
「NAME:.metadata.name、UID:.metadata.uid、IP:.status.podIP、RESTARTS:.status.containerStatuses[0].restartCount」
#2. ポッド上のイベント
kubectl はポッド <pod> を記述します。 grep -A 20 "イベント:"
# 3. インプレースサイズ変更ステータス
kubectl get pod <pod> -o jsonpath='{.status.resize}'
結論
コンテナの再起動は中断を伴いますが、洗練されます

[切り捨てられた]

## Original Extract

A production internals guide verified against Kubernetes 1.35 GA Companion repository: github.com/opscart/k8s-pod-restart-mechanics Engineers say “the pod…

Skip to content
Accessibility
help
About
Discover CNCF
Membership Hub – For Current Members Learn about all the benefits of being a CNCF Member
Technical Oversight Committee The TOC defines CNCF’s technical vision and provides experienced technical leadership to the cloud native community
Governing Board The GB is responsible for marketing, business oversight, and budget decisions for CNCF
End User Technical Advisory Board The End User TAB serves as the voice of the end users in CNCF community decisions
Ambassadors Meet our Ambassadors—experienced practitioners passionate about helping others learn about cloud native technologies
Graduated Projects considered stable, widely adopted, and production ready, attracting thousands of contributors
Incubating Projects used successfully in production by a small number users with a healthy pool of contributors
Sandbox Experimental projects not yet widely tested in production on the bleeding edge of technology
Archived Projects that have reached the end of their lifecycle and have become inactive
Project Metrics View metrics of CNCF projects moving through maturity levels
Contribute Join the 150K+ folx in #TeamCloudNative who’ve contributed their expertise to CNCF hosted projects
Services for CNCF Projects CNCF services for our open source projects – from marketing to legal services
Cloud Native Landscape A comprehensive categorical overview of projects and product offerings in the cloud native space
Project Journey Reports Showing how CNCF has impacted the progress and growth of various graduated projects
Project Tools Quick links to tools and resources for your CNCF project
Latest Project Journey Reports
Training Overview Find your path to cloud native success with training and certificates from the pioneer of cloud-native technology
Certifications Get certified by the only authoritative source for cloud-native certification accepted by industry
Courses Learn the basics of cloud native or train for a certification with training courses built in collaboration with the Linux Foundation
Kubestronaut Program Uplevel your Kubernetes skills
Kubernetes Training Partners Find a qualified KCNTP to prepare for your next certification
Certified Kubernetes Software conformance ensures your versions of CNCF projects support the required APIs
Cloud Native Network Function Certification (Beta) CNF Certification ensures applications demonstrate cloud native best practices
Kubernetes Certified Service Provider KCSPs have deep experience helping enterprises successfully adopt cloud native technologies
Enroll your company as a CNCF End User and save more than $10K in training and conference costs
End User Community Join our vendor-neutral community using cloud native technologies to build products and services
Events Meet #TeamCloudNative and CNCF staff at events around the world
Case Studies Read real-world case studies about the impact cloud native projects are having on organizations around the world
Humans of Cloud Native Read stories of amazing individuals and their contributions
The Cloud Native Heroes Challenge Be a cloud native hero! Help us defeat patent trolls to earn swag and prizes
Online Programs Watch our free online programs for the latest insights into cloud native technologies and projects
Community Groups Join #TeamCloudNative at events and meetups near you
Phippy & Friends Phippy explains core cloud native concepts in simple terms through stories perfect for all ages
Cloud Native Glossary Explore cloud native concepts in clear and simple language – no technical knowledge required!
Blog Catch up on the latest happenings and technical insights from #TeamCloudNative
Announcements Media releases and official CNCF announcements
News CNCF projects and #TeamCloudNative in the media
Reports Read transparent, in-depth reports on our organization, events, and projects
Search
CNCF
Blog
/
Project Maintainer Post
When Kubernetes restarts your pod — And when it doesn’t
Posted on March 17, 2026
by Shamsher Khan, Project Maintainer
CNCF projects highlighted in this post
A production internals guide verified against Kubernetes 1.35 GA
Companion repository: github.com/opscart/k8s-pod-restart-mechanics
Engineers say “the pod restarted” when they mean four different things. Getting this wrong leads to flawed runbooks and bad on-call decisions.
The practical test: Did the pod UID change? If yes — that is recreation, not a container restart. Restart count resets to zero. If no — same pod object, container process restarted inside it.
The core insight: What kubelet actually watches
kubelet watches the pod spec — not ConfigMaps, not Secrets, not Istio CRDs. If the pod spec didn’t change, kubelet never fires. This single fact explains the majority of “why didn’t my config update?” investigations in production.
Mutating admission webhooks can change the pod spec at creation time, but never after admission — they cannot trigger container restarts post-creation.
The flowchart below translates the same matrix into a decision path you can walk at 2am during an incident.
Diagram 1: Complete decision flowchart — does this change require a pod restart?
Scenario 1: ConfigMap — Why the same change has two behaviors
[Diagram 1: ConfigMap env var vs volume mount — env var pod frozen, volume pod auto-synced via kubelet symlink swap]
Env var mode (envFrom / valueFrom) : The kernel copies env vars into /proc/<pid>/environ at execve() . That memory is owned by the process — no external system can modify it. Update the ConfigMap and kubelet sees no pod spec change, does nothing. The process keeps old values indefinitely.
Volume mount mode : kubelet syncs via an atomic symlink swap, not a file write:
/etc/config/
├── ..2025_12_19_11_30_00/ ← NEW data dir (kubelet creates this)
│ └── APP_COLOR ← "red"
├── ..data ─────────────────▶ ..2025_12_19_11_30_00/ ← symlink SWAPPED atomically
└── APP_COLOR ──────────────▶ ..data/APP_COLOR
The symlink swap generates IN_CREATE on ..data — NOT IN_MODIFY on the file. Applications watching IN_MODIFY on an open file descriptor miss this entirely. This is why nginx does not auto-reload on ConfigMap changes without explicit inotify handling.
Lab Evidence (01-configmap/ in companion repo)
ConfigMap updated: APP_COLOR blue → red
Pod A (env var): APP_COLOR=blue ← frozen, restart count: 0
Pod B (volume mount): APP_COLOR=red ← auto-synced, restart count: 0
Correct inotify pattern — watch the directory, not the file
watcher.Add(filepath.Dir(configPath)) // watches /etc/config/ — catches IN_CREATE
// watcher.Add(configPath) // misses symlink swap entirely
for event := range watcher.Events {
if event.Op&fsnotify.Create == fsnotify.Create {
reloadConfig()
}
}
Scenario 2: Image updates — Recreation vs container restart vs CrashLoop
These three scenarios look similar but are fundamentally different:
Successful image update — pod is recreated
BEFORE: Pod UID aaa-bbb, IP 10.244.1.5, nginx:1.25, restarts: 0
AFTER: Pod UID xxx-yyy, IP 10.244.1.6, nginx:1.27, restarts: 0
↑ UID changed — RECREATION, not container restart
Diagram 3: Rolling update flow showing new ReplicaSet creation, pod recreation, and old RS retained for rollback.
ImagePullBackOff — old pod stays protected
Old pod: Running ← Kubernetes keeps it alive
New pod: ImagePullBackOff ← stuck, old pod never killed until new one is healthy
CrashLoopBackOff — same pod, restart count climbs
Pod UID: aaa-bbb ← UNCHANGED
Restart count: 0 → 1 → 2 → 3 ← same pod object, container crashing
Diagnostic rule: Climbing restart count with unchanged UID = crash loop. Zero restart count with new UID = rolling update.
StatefulSet note: StatefulSet pods are also recreated on image change, but ordinal identity (pod-0, pod-1) and PVC bindings are preserved. Container restart semantics are identical to Deployments — identity persistence does not imply in-place container restart.
Scenario 3: In-place resource resize (K8s 1.35 GA)
K8s 1.35 made in-place pod resize generally available (kubernetes.io/blog/2025/12/19/kubernetes-v1-35-in-place-pod-resize-ga). Both CPU and memory can be resized without pod recreation — UID and IP never change. In-place resize availability depends on CRI and node OS support; verified on containerd 1.7+ with Linux cgroups v2.
What happens to the container depends on resizePolicy , which you define explicitly:
resizePolicy:
- resourceName: cpu
restartPolicy: NotRequired # cgroup quota updated, process untouched
- resourceName: memory
restartPolicy: RestartContainer # container restarts inside same pod
Lab Evidence (05-resource-resize/ — requires K8s 1.35+)
CPU resize 200m → 500m (NotRequired):
UID: d7c99204 IP: 10.244.0.7 Restarts: 0 ← all unchanged
Memory resize 256Mi → 512Mi (RestartContainer):
UID: d7c99204 IP: 10.244.0.7 Restarts: 1
↑ same pod ↑ same IP ↑ our policy choice, not K8s forcing it
Important: The default resizePolicy for memory is NotRequired. If you omit it, memory resize silently updates the cgroup without restarting the container — and your JVM heap stays at the old size. Always define resizePolicy explicitly for memory.
kubectl patch pod my-pod -n my-namespace \
--subresource resize \
-p '{"spec":{"containers":[{"name":"app","resources":{
"requests":{"cpu":"250m","memory":"128Mi"},
"limits":{"cpu":"500m","memory":"256Mi"}
}}]}}'
# Note: omit --type=merge — causes a validation error with --subresource resize
Scenario 4: Istio routing — Zero restarts via xDS
Istio VirtualService and DestinationRule changes never trigger container restarts. Istiod maintains a persistent bidirectional gRPC stream to each Envoy sidecar — routing updates are pushed in milliseconds, in-memory swap, no pod touched, no file written.
Lab Evidence (04-istio-routing/ in companion repo)
Four pods. Three routing changes:
100% v1 → 80/20 canary → 100% v2
Restart counts: BEFORE 0 0 0 0 / AFTER 0 0 0 0
Pod ages: unchanged throughout all three changes.
Scenario 5: Stakater reloader — Automating the manual step
When apps consume ConfigMaps via env vars, someone must run kubectl rollout restart after every update. Reloader automates this using Kubernetes watch events — detection is near-instant, not polling.
metadata:
annotations:
reloader.stakater.com/auto: "true"
Production gotcha: The default Helm install uses watchGlobally=false — Reloader only watches its own namespace. Annotated Deployments in other namespaces are silently ignored, no error thrown. Always install with watchGlobally=true.
helm install reloader stakater/reloader \
--namespace reloader \
--set reloader.watchGlobally=true
Lab Evidence (07-stakater-reloader/ in companion repo)
ConfigMap updated. No kubectl rollout restart run.
New pod APP_MESSAGE: "Hello from OpsCart v2 — auto reloaded!"
Rolling container restart triggered automatically.
When hot-reload goes wrong
Hot-reload is not always safer than a container restart. Two failure modes worth knowing:
Semantically invalid config accepted silently
The file updates, the inotify handler fires, no error is thrown — but the new config has a logic error. The pod passes health checks and runs broken for hours. A container restart with a bad config fails immediately and loudly. Hot-reload with bad config fails quietly and late.
Mitigation: Validate config before swapping atomically.
Envoy rejects xDS push silently
Istiod pushes a RouteConfiguration referencing a cluster not yet propagated. Envoy rejects it and continues with old routing rules. No pod event fires. Mitigation: Monitor pilot_xds_push_errors and use istioctl proxy-status.
Observability: Three commands every operator should know
# 1. Container restart or pod recreation? Check UID change
kubectl get pod <pod> -o custom-columns=\
"NAME:.metadata.name,UID:.metadata.uid,IP:.status.podIP,RESTARTS:.status.containerStatuses[0].restartCount"
# 2. Events on the pod
kubectl describe pod <pod> | grep -A 20 "Events:"
# 3. In-place resize status
kubectl get pod <pod> -o jsonpath='{.status.resize}'
Conclusion
A container restart is disruptive but hones

[truncated]
