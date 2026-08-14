---
source: "https://github.com/inevolin/k8s-cpu-limits-analyzed"
hn_url: "https://news.ycombinator.com/item?id=49296939"
title: "For the love of god stop using CPU limits in Kubernetes"
article_title: "GitHub - inevolin/k8s-cpu-limits-analyzed: Kubernetes CPU limits make your apps (very) slow and costly · GitHub"
author: "iljanevo"
captured_at: "2026-08-14T10:55:33Z"
capture_tool: "hn-digest"
hn_id: 49296939
score: 12
comments: 1
posted_at: "2026-08-14T10:41:21Z"
tags:
  - hacker-news
  - translated
---

# For the love of god stop using CPU limits in Kubernetes

- HN: [49296939](https://news.ycombinator.com/item?id=49296939)
- Source: [github.com](https://github.com/inevolin/k8s-cpu-limits-analyzed)
- Score: 12
- Comments: 1
- Posted: 2026-08-14T10:41:21Z

## Translation

タイトル: 神の愛のために、Kubernetes での CPU 制限の使用をやめてください
記事のタイトル: GitHub - inevolin/k8s-cpu-limits-analyzed: Kubernetes の CPU 制限により、アプリが (非常に) 遅くなりコストが高くなります · GitHub
説明: Kubernetes の CPU 制限により、アプリが (非常に) 遅くなり、コストが高くなります - inevolin/k8s-cpu-limits-analyzed

記事本文:
GitHub - inevolin/k8s-cpu-limits-analyzed: Kubernetes の CPU 制限により、アプリが (非常に) 遅くなりコストが高くなります · GitHub
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
イネボリン
/
k8s-cpu-limits-analyzed
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット アプリ アプリ アセット アセット ドキュメント ドキュメント k8s k8s 結果 結果 スクリプト スクリプト ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Kubernetes CPU

制限によりアプリが遅くなり、コストが高くなります
プラットフォームエンジニアリング・分析
同じアプリを CPU 制限ありとなしでテストしました。同じコード、同じ CPU リクエスト、同じ負荷。
唯一の違いは限界でした。私たちが測定した内容と、それを自分で確認する方法は次のとおりです
クラスター。
CPU 制限を削除します。ノードに空き CPU がある場合でも、アプリが 1 秒間に何度もフリーズします。
CPU リクエストを保持します。リクエストは本当の保護です。彼らはすべてのアプリのシェアを保証します。
メモリ制限を守ります。記憶が違います。メモリ制限によりノードは引き続き保護されます。
高速化: トラフィックのピーク時にテール レイテンシが崩壊するのではなく持続し、CPU バウンドの起動作業が約 2 倍早く終了します (セクション 5、10)。
ハードウェアの削減: ほとんどのクラスターは、ピーク時に使用するよりもはるかに多くの CPU を予約します。制限を削除した後にリクエストのサイズを適切に設定すると、ノードの有意義な共有が可能になります (セクション 10)。
より安価: 実際に実行された例示的なコスト モデルでは、クラスターあたり年間数万ドルになります (セクション 10)。実際の数値を取得するには、独自の価格を入力してください。
この分析では: 1. リクエストと制限 · 2. スロットルの仕組み · 3. CFS フェア シェアリング · 4. 最悪のケース · 5. 測定した内容 · 6. ノイジーネイバー · 7. 壊れたとき · 8. .NET の一歩 · 9. 計画 · 10. 結果 · 11. Q&A · 12. 用語集
さらに深くなっていきます。このページはその議論です。その中の各主張はより長い文書によって裏付けられており、
ここを読んで収集したものをインラインでリンクします。
1. 2 つの設定、2 つのまったく異なるジョブ
CPUリクエスト
CPU制限
それは何ですか
保証された CPU スライス
硬い天井、壁
他のアプリも保護しますか?
はい。 CPUはリクエストサイズによって共有されます。
いいえ、それ自体のアプリをブロックするだけです。
ノード上のアイドル状態の CPU
アプリなら無料で借りられる
無駄だ。制限がそれをブロックします。
同じアプリが2回あります。唯一の違いは、ノードのアイドル状態の CPU へのアクセスです。リクエスト（固形）は、

両方の行で同じです。アイドル状態の CPU を借用しても、誰からも何も得られません。
さらに深い: これらは 2 つの異なる cgroup ファイルであり、1 つのノブの 2 つの設定ではありません。
docs/01- Theory.md の手順を説明します。
cpu.max (制限) と cpu.weight (リクエスト)、およびそれぞれが実際に制御するもの。
2. 制限によって実際にアプリが停止される仕組み
カーネル (オペレーティング システムの中核部分) は、ウィンドウの制限を 100 に強制します。
ミリ秒。 500m 制限 (500 ミリコア、半分のコア) は、次のことを意味します: 1 つあたり 50 ミリ秒の CPU 時間
窓。予算がなくなると、カーネルはアプリ全体をフリーズします。
次のウィンドウ。これはスロットリングです。
一般的な .NET サービスは、HTTP ハンドラー、バックグラウンド コンシューマー、GC (ガベージ) など、多くのスレッドを実行します。
コレクター）。すべてのスレッドは 1 つのバジェットを共有します。私たちのテスト ノードには 4 つのコアがあるため、最大でも
4 つのスレッドを同時に実行できます。 8 つのビジー スレッドが依然として 50 ミリ秒のバジェットを約 3 分で使い切ってしまいます。
12.5ミリ秒のリアルタイム:
フリーズは 1 秒あたり最大 10 回繰り返されます。次のグラフは、ダッシュボードにそれらが表示されない理由を示しています。
オレンジ色の秒ごとにアプリが限界に達し、カーネルがフリーズしました。 1 分間の平均が制限に近づくことはないため、どのグラフも健全に見えます。これにより、ダッシュボードが緑色のまま、コンテナーが 1 日中スロットルされ続けることができます。平均 CPU だけを見ている場合は、これを確認することはできません。代わりに、container_cpu_cfs_throttled_periods_total を確認してください。
私たちはこれを証明しました。平均 320m だけを必要とする負荷を、
500m制限。平均が限界に達することはありませんでした。アプリはまだ停止しています:
同じアプリ、同じ負荷 (5 ミリ秒の 8 つの並列タスク、1 秒あたり 8 つのリクエスト)。この制限により、遅いリクエストは 2.4 倍遅くなりましたが、平均 CPU は制限を十分に下回りました。
より深く: 並列化により見た目よりもはるかに悪化する理由 —
16 スレッドが 2 ミリ秒未満の実行時間で 300m クォータを消費します —
そしてなぜその結果が

減速ではなく失速です、
これにより、平均は問題ないように見えますが、テール レイテンシーが破壊されます。
3. 制限なしで、誰が CPU を共有しますか? CFS の紹介
Linux には、CFS (Completely Fair Scheduler) という参照機能が組み込まれています。すべてのポッド
重みがあります。 Kubernetes はポッドの CPU リクエストから重みを設定します。
ルールは単純です。ノードが完全にビジー状態になると、ポッドはそのパフォーマンスに基づいて CPU を共有します。
重み。ポッドがアイドル状態になると、その共有の使用が停止されます。他のポッドは代わりにその共有を使用できます。アイドル状態のポッドは、
再び仕事ができたらすぐに共有してください。
1 つのノード上の 3 つのアプリ、リクエストからの重み: A = 100m、B = 200m、C = 700m。このレフェリーは、常にすべての Linux サーバー上で実行されます。動作するために CPU 制限は必要ありません。制限により追加されるのは 1 つだけです。セクション 2 でのフリーズです。
より深い: 完全な議論
制限は保護を提供しません、リクエストがまだ提供していない、
それに加えて、実際の例外が 1 つあります。
静的 CPU マネージャーによる QoS の保証。「制限」はコアの固定ではなく、コアの固定です。
スロットルクォータ。
4. 「しかし、すべてのポッドが一度に 100% の CPU を使用すると、ノードは停止します。」
いいえ、この懸念は CPU とメモリを同じ問題として扱います。これが考えられる最悪のCPUです
ノード上の瞬間:
4 コア ノード、完全にビジー状態。 CFS はリクエストの重みに応じて CPU を分割します。各ポッドは保証されたリクエスト シェアに戻ります。アプリは遅くなり、何も中断されません。スケジューラは、すべてのポッドのリクエストがノード内に収まることをすでに確認しています。
3 つの保護機能により、ノードの障害は問題になりません。すべて CPU の制限とは無関係です。
そして、ほとんどのクラスターでは、ノードの CPU 使用率はキャパシティをはるかに下回っており、多くの場合、単一または低レベルにあります。
二桁。 「全員が 100% の状態」の瞬間は主に理論上のものですが、現実の日常的なものです。
問題はその逆です。つまり、スロットルされたアプリがアクセスできないアイドル状態の CPU です。ノードの場合
2人のプロは本当に忙しくなります

上記の条件は引き続き適用されます。予約されたシステム CPU は、
ノードと kubelet の応答性が高く、CFS 重みベースの共有により、すべてのポッドが確実に共有されます。
リクエストされたシェア。ノードレベルの CPU ダッシュボードでは、これよりずっと前から高い使用率が継続的に示されていました。
が心配になります。
「しかし、ResourceQuota も LimitRange もまだありません。」
これは正当な反論であり、その背後にある懸念は現実です。すべてのアプリケーションが同じものを共有している場合
ResourceQuota も LimitRange もない名前空間。チームに正確な設定を強制するものは何もない
リクエスト。一般的な暫定ステップは、limits.cpu: 1 のようなものを「サーキット ブレーカー」として設定することです。
天井ではなく」 - 実際にスロットリングが停止するのに十分な高さです。
上限を設定することは優れた緊急修正であり、サービスのスロットリングを終了します。
に適用されます。ただし、そのサービス周辺のポッドの安全メカニズムとして、CPU 制限が設けられています。
やっているように見えることをやらない。
「サーキットブレーカー」の制限によるジレンマ。決して高く設定された制限
Trip (実際の使用率 30 倍) は何もしないため、誰も保護しません。十分に低い制限
実際、トリップは接続されているポッドにのみ害を及ぼし、本来の隣接ポッドには害を与えません。
守る。それは安全機構であると同時に無害であることはできません。
制限のサイズを「サービスが使用する量の 30 倍」とする測定に関する警告が 1 つあります。できません
サイズ 制限が同じワークロードを調整している間に測定された使用量からの制限。重い
スロットルされたポッドの記録された使用量は、制限によって許可されたものであり、アプリが望んだものではありません。サイズ設定
抑制された使用量からは、次の小さすぎる制限が構築されるだけです。
では、実際には何が必要なのでしょうか？まさにその反対の名前: LimitRange は、
pod はデフォルトのリクエスト、および名前空間ごとの総リクエスト数を制限する ResourceQuota です。どちらも機能します
制限ではなくリクエストに応じて適用され、両方とも既存の共有名前空間にそのまま適用できます。
名前空間

再設計が最初に行われなければなりません。実際には、ギャップは見た目よりも小さいことがよくあります。
共有 Helm チャートは通常、テンプレート化されたサービスごとにリクエストを設定します。
シーケンスについては誇張しやすいため、明確にしておきます。LimitRange は実行する価値があります。
まず、これは 1 つのマニフェストであり、コストはかからず、通常の処理をバイパスするものはすべてキャッチします。
デプロイメントパス。 ResourceQuota は後続する可能性があります。安定性ではなくコストを守ります。
スケジューラは、リクエストが適合しないポッドをすでに拒否しています。しかし、どちらのゲートも除去を制限しません。
本当の安全策は、変更自体を削除することです。変更は、正確な値ファイルに影響を与えます。
不足している CPU リクエストが表示されます。
さらに詳しく: docs/04-objections.md がこれと残りの部分に答えます。
異議申し立ては個別に設定されます。
ノードがホット状態になったときに何が起こるか、
マルチテナント名前空間 、
制限を削除すると QoS クラスが変更されるかどうか。
5. 測定しました: 同じアプリ、制限ありと制限なし
このリポジトリのスクリプトは、同じ .NET 10 API を使い捨ての名前空間に 2 回デプロイします。
どちらのポッドも 250m CPU を要求します。 1 つは 500m CPU 制限があり、もう 1 つは制限がありません。ロードジェネレータは、
どちらでも同じリクエストです。すべてがスクリプト化されており、反復可能であるため、すべてを再実行できます。
あなた自身のクラスター。
注: スタートアップ (3 回の実行、中央値を表示) を除く、以下のすべての数値は、
単一の実行。繰り返しや平均化は行われません。約 10% 未満の違いは、証拠ではなくヒントとして扱います。
急増中、制限付きアプリはすべての 100 ミリ秒ウィンドウの 89% でフリーズしました。無制限のアプリは、速度がまったく低下することなく、12 倍のスパイクを吸収しました。完全なテーブル: リポジトリ内の results/SUMMARY.md。
起動パネルに関する注意点が 1 つあります。テスト コンテナーは dotnet run を実行し、コンパイルを実行します。
毎回起動するため、ここでの「準備時間」は CPU に依存したコンパイル作業によって支配されます。それは、
スロットル遅延の明確なデモンストレーション

g 起動しますが、2x をユニバーサルとして読み取らないでください。
図: 起動のほとんどが I/O (構成の取得、接続のオープン) を待機しているサービス
得られる利益ははるかに少なくなります。ゲインは、アプリがレポートするまでに実行する CPU 作業量に応じて変化します。
準備完了。レイテンシーおよびスパイクパネルにはそのような注意事項はありません。
より深い: 完全なパーセンタイル テーブルを含むシナリオごとのレポートが存在します。
results/ ( 起動 、 レイテンシ 、
バースト、うるさい隣人、
Postgres )。特に Postgres パネルでは、
docs/03-postgres.md は、どの効果が測定されたかに注意して、
これは単なるメカニズムです。pgbench のゲインは 1 回の実行であり、並列クエリ テストは
null の結果。
6. ノイジーネイバーテスト (ビジー状態のポッドはネイバーの速度を低下させますか?)
これが最も一般的な恐怖です。私たちはそれをテストしました。同じノードに別のポッドを配置します。あのポッドは走った
常に CPU を使用する 8 スレッド (制限なし)。次に、アプリの p99 を再度測定しました。
無制限のビクティム (バー 1 と 2) の場合、ネイバーは p99 を 7 ミリ秒未満で移動し、
決してスロットルされていません。 CFS の約束どおり (セクション 3)、その要求により保護されました。被害者のために
独自の 500m 制限 (バー 3) により、同じノード上での隣接ノードの配置は保証されませんでした。
他の 2 つのバーとは異なります ( results/40-noisy-neighbor.md を参照)。それでも、

[切り捨てられた]

## Original Extract

Kubernetes CPU limits make your apps (very) slow and costly - inevolin/k8s-cpu-limits-analyzed

GitHub - inevolin/k8s-cpu-limits-analyzed: Kubernetes CPU limits make your apps (very) slow and costly · GitHub
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
inevolin
/
k8s-cpu-limits-analyzed
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits app app assets assets docs docs k8s k8s results results scripts scripts LICENSE LICENSE README.md README.md View all files Repository files navigation
Kubernetes CPU limits make your apps slow and costly
Platform engineering · analysis
We tested the same app with and without a CPU limit. Same code, same CPU request, same load.
The only difference was the limit. Here is what we measured — and how to check it on your own
cluster.
Remove CPU limits. They freeze your apps many times per second, even when the node has free CPU.
Keep CPU requests. Requests are the real protection. They guarantee every app its share.
Keep memory limits. Memory is different. A memory limit still protects the node.
Faster: tail latency holds up under traffic peaks instead of collapsing, and CPU-bound startup work finishes about 2x sooner (sections 5, 10).
Less hardware: most clusters reserve far more CPU than they ever use at peak. Right-sizing requests after dropping limits lets a meaningful share of nodes go (section 10).
Cheaper: a worked, illustrative cost model puts this in the tens of thousands of dollars per year per cluster (section 10) — plug in your own prices to get a real number.
In this analysis: 1. Requests vs limits · 2. How throttling works · 3. CFS fair sharing · 4. The worst case · 5. What we measured · 6. Noisy neighbor · 7. When it breaks · 8. One step for .NET · 9. The plan · 10. Outcomes · 11. Q&A · 12. Glossary
Going deeper. This page is the argument. Each claim in it is backed by a longer document,
linked inline as you read and collected here:
1. Two settings, two very different jobs
CPU request
CPU limit
What it is
A guaranteed slice of CPU
A hard ceiling, a wall
Protects other apps?
Yes. CPU is shared by request size.
No. It only blocks its own app.
Idle CPU on the node
App may borrow it for free
Wasted. The limit blocks it.
Same app twice. The only difference is access to the node's idle CPU. The request (solid) is identical in both rows — borrowing idle CPU takes nothing from anyone.
Deeper: these are two different cgroup files, not two settings of one knob.
docs/01-theory.md walks through
cpu.max (the limit) and cpu.weight (the request) and what each one actually controls.
2. How a limit really stops your app
The kernel (the core part of the operating system) enforces limits in windows of 100
milliseconds . A 500m limit (500 millicores, half a core) means: 50 ms of CPU time per
window. When the budget runs out, the kernel freezes the whole app until the
next window. This is throttling .
A typical .NET service runs many threads: HTTP handlers, background consumers, and the GC (the garbage
collector). All threads share one budget. Our test node has 4 cores, so at most
4 threads can run at the same instant. 8 busy threads still use up the 50 ms budget in about
12.5 ms of real time:
The freezes repeat up to 10 times per second. The next chart shows why your dashboards never see them.
Every orange second the app hit its limit and the kernel froze it. The 1-minute average never comes close to the limit, so every graph looks healthy. This is how a container can be throttled all day while its dashboard stays green. If you only ever look at averaged CPU, you cannot see this — check container_cpu_cfs_throttled_periods_total instead.
We proved this. We sent load that needs only 320m on average to an app with a
500m limit . The average never touched the limit. The app still stalled:
Same app, same load (8 parallel tasks of 5 ms, 8 requests per second). The limit made the slow requests 2.4x slower, while average CPU stayed well under the limit.
Deeper: why parallelism makes this so much worse than it looks —
16 threads burn a 300m quota in under 2 ms of wall time —
and why the result is a stall rather than a slowdown ,
which is what destroys your tail latency while the average looks fine.
3. Without limits, who shares the CPU? Meet CFS
Linux has a built-in referee: CFS, the Completely Fair Scheduler . Every pod
has a weight . Kubernetes sets the weight from the pod's CPU request .
The rule is simple: when the node is fully busy, pods share the CPU based on their
weights . When a pod is idle, it stops using its share. Other pods can use that share instead. The idle pod gets its
share back as soon as it has work again.
Three apps on one node, weights from requests: A = 100m, B = 200m, C = 700m. This referee runs on every Linux server, always. It needs no CPU limit to work. A limit adds only one thing extra: the freezes from section 2.
Deeper: the full argument that
a limit provides no protection a request does not already provide ,
plus the one real exception :
Guaranteed QoS with the static CPU manager, where the "limit" is pinning cores rather than
throttling quota.
4. "But if all pods use 100% CPU at once, the node will die!"
No. This fear treats CPU and memory as the same problem. Here is the worst possible CPU
moment on a node:
A 4-core node, fully busy. CFS splits CPU by request weight. Each pod drops back to its guaranteed request share — apps get slower, nothing breaks. The scheduler already makes sure the requests of all pods fit inside the node.
Three protections make node failure a non-issue, all independent of CPU limits:
And in most clusters, node CPU utilization sits far below capacity — often in the single or low
double digits. The "everyone at 100%" moment is largely theoretical, while the real, daily
problem is the opposite one: idle CPU that throttled apps are not allowed to touch. If a node
ever does get fully busy, the two protections above still hold: reserved system CPU keeps the
node and kubelet responsive, and CFS weight-based sharing still guarantees every pod its
requested share. A node-level CPU dashboard would show sustained high usage well before this
becomes a concern.
"But we have no ResourceQuota or LimitRange yet"
This is a fair objection, and the concern behind it is real: if every application shares one
namespace with no ResourceQuota and no LimitRange, nothing forces a team to set accurate
requests. A common interim step is to set something like limits.cpu: 1 as a "circuit breaker
rather than a ceiling" — high enough that throttling stops in practice.
Setting a high limit is a fine emergency fix, and it does end throttling for the service it is
applied to. But as a safety mechanism for the pods around that service, a CPU limit does
not do what it looks like it does.
The dilemma with "circuit breaker" limits. A limit set high enough never to
trip (30x real use) protects nobody, because it never does anything. A limit low enough to
actually trip only harms the pod it is attached to, never the neighbor it was meant to
protect. It cannot be both a safety mechanism and harmless.
One measurement warning on sizing a limit as "30x what the service uses". You cannot
size a limit from usage measured while a limit was throttling that same workload. A heavily
throttled pod's recorded usage is what the limit allowed , not what the app wanted . Sizing
from suppressed usage just builds the next too-small limit.
So what is actually needed? Exactly what the objection names: a LimitRange giving every
pod a default request, and a ResourceQuota capping total requests per namespace. Both work
on requests, not limits, and both can be applied to an existing shared namespace as-is — no
namespace redesign has to come first. In practice the gap is often smaller than it looks, since
a shared Helm chart typically sets a request for every service it templates.
To be clear about sequencing, because it is easy to overstate: the LimitRange is worth doing
first — it is one manifest, it costs nothing, and it catches anything that bypasses your normal
deployment path. The ResourceQuota can trail ; it guards cost, not stability, since the
scheduler already refuses pods whose requests do not fit. But neither one gates limit removal.
The real safeguard is the removal change itself: it touches the exact values file where a
missing CPU request would be visible.
Deeper: docs/04-objections.md answers this and the rest of the
objection set individually, including
what happens when a node runs hot ,
multi-tenant namespaces ,
and whether removing a limit changes QoS class .
5. We measured it: same app, with and without a limit
The scripts in this repo deploy the same .NET 10 API twice into a throwaway namespace.
Both pods request 250m CPU. One has a 500m CPU limit, one has none. A load generator sends the
same requests at both. Everything is scripted and repeatable, so you can re-run all of it on
your own cluster.
Note: except startup (3 runs, median shown), every number below is from a
single run, not repeated or averaged. Treat differences under about 10% as a hint, not proof.
During the spike, the limited app was frozen in 89% of all 100 ms windows . The unlimited app absorbed a 12x spike with no slowdown at all. Full tables: results/SUMMARY.md in the repo.
One caveat on the startup panel. The test container runs dotnet run , which compiles on
every start, so "time to ready" here is dominated by CPU-bound compile work. That makes it a
clean demonstration of throttling delaying startup, but do not read the 2x as a universal
figure: a service whose startup is mostly waiting on I/O (pulling config, opening connections)
will gain far less. The gain scales with how much CPU work your app does before it reports
ready. The latency and spike panels have no such caveat.
Deeper: per-scenario reports with the full percentile tables live in
results/ ( startup , latency ,
burst , noisy neighbor ,
Postgres ). On the Postgres panel specifically,
docs/03-postgres.md is careful about which effects were measured and
which are only mechanism — the pgbench gain is a single run, and the parallel-query test was
a null result.
6. The noisy neighbor test (does a busy pod slow down its neighbors?)
This is the most common fear. We tested it. We put another pod on the same node. That pod ran
8 threads that constantly use CPU, with no limit. Then we measured our app's p99 again:
For the unlimited victim (bars 1 and 2), the neighbor moved p99 by less than 7 ms, and it was
never throttled. Its request protected it, exactly as CFS promises (section 3). For the victim
with its own 500m limit (bar 3), the neighbor's placement on the same node was not guaranteed,
unlike the other two bars (see results/40-noisy-neighbor.md ). Even so,

[truncated]
