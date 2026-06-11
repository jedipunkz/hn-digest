---
source: "https://github.com/adamf/JustRebootIt"
hn_url: "https://news.ycombinator.com/item?id=48490831"
title: "Show HN: JustRebootIt, WAN latency monitor with LLM root cause analysis"
article_title: "GitHub - adamf/JustRebootIt · GitHub"
author: "adam_gyroscope"
captured_at: "2026-06-11T16:28:36Z"
capture_tool: "hn-digest"
hn_id: 48490831
score: 2
comments: 0
posted_at: "2026-06-11T14:27:30Z"
tags:
  - hacker-news
  - translated
---

# Show HN: JustRebootIt, WAN latency monitor with LLM root cause analysis

- HN: [48490831](https://news.ycombinator.com/item?id=48490831)
- Source: [github.com](https://github.com/adamf/JustRebootIt)
- Score: 2
- Comments: 0
- Posted: 2026-06-11T14:27:30Z

## Translation

タイトル: HN を表示: JustRebootIt、LLM 根本原因分析を使用した WAN 遅延モニター
記事のタイトル: GitHub - adamf/JustRebootIt · GitHub
説明: GitHub でアカウントを作成して、adamf/JustRebootIt の開発に貢献します。
HN テキスト: 私はジッターとレイテンシーを見つけるために喫煙するのが大好きですが、より深く掘り下げてすべての関連信号を取得するものが常に必要だったので、JustRebootIt を構築しました。 JRI は、多数の並列 ping とトレースルートを実行し、スパイクが発生していることを検出し、イベント中により詳細なプローブを実行して、診断のためにそれらすべてを LLM に入力できます。また、相関するイベントの場合は、UDM から読み取り、デバイスから CPU/メモリなどを収集します。フィードバックやPRも大歓迎です！

記事本文:
GitHub - adamf/JustRebootIt · GitHub
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
アダムフ
/
ただ再起動してください
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メイン ブランチ タグ ファイル コードに移動 その他のアクション メニューを開く フォルダーとファイル

レ
21 コミット 21 コミット .github/ workflows .github/ workflows cmd cmd config config docker docker docs docs external 内部 .env.example .env.example .gitignore .gitignore Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md docker-compose.yml docker-compose.yml go.mod go.mod go.sum go.sum すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ホームインターネットの断続的な遅延を解決する自己完結型の Dockerized レコーダー
スパイク — ビデオ通話がしばらくフリーズするまではすべてが順調に感じられるタイプ
3 秒かかり、速度テストを実行する頃には消えています。
さまざまな環境への遅延、ジッター、パケット損失を継続的に測定します。
ターゲット (ゲートウェイ、ISP、および適切に運営されているパブリック アンカー) をトレースします。
ネットワーク パスを使用して遅延が発生する場所を特定します。UniFi を使用している場合は、
Dream Machine Pro — これらすべてを WAN スループットとゲートウェイと相関させます。
ロードします。スクリーンショットを撮るように設計された 1 つの Grafana ダッシュボードにすべてが表示されます
または、ISP (Comcast など) のネットワーク エンジニアとリンク共有します。
The probes are written in Go and run fully in parallel, so dozens of targets are
プローバ自体がボトルネックになることなく、タイトなサイクルで測定できます。
上記のダッシュボードからの実際の診断 (イベント #6): 「
8.8.4.4 was not specific to Google — it was local bufferbloat on the WAN uplink
トラフィックバースト中…各アンカーへの RTT 中央値が一斉に跳ね上がったので、
スパイクはすべての独立したパスで共有されました… udm_wan_rx は最大 15.2 MB/秒に達し、
WAN 遅延はベースラインの約 10 ～ 11 ミリ秒から 33 ミリ秒まで増加しました。これが古典です
bufferbloat pattern: the local uplink filled and queued packets.自信:
高い。推奨されるアクション: UDM Pro でスマート キュー (SQM/CAKE) を有効にします。
limits ~85–90% of the measured line rate — this is a local fix, no

ISP
チケット」 — イベントの発生時に AI 分析によって自動的に書き込まれます。
┌─────────┐ ICMP ping + トレースルート
│ プローバー │───────────────► ターゲットが多い
│ (行く) │ (ゲートウェイ、ISP、アンカー)
━━━━┬──────┘
│ /メトリクス
UniFi ドリーム ┌──────┴──────┐
Machine Pro ─►│ udm-exporter│ WAN スループット、ゲートウェイ CPU/メモリ、
(API) │ (Go) │ スピードテスト、クライアント数
━━━━┬──────┘
│ /メトリクス
┌──────┴──────┐ ┌────────┐
│ プロメテウス │───────►│ グラファナ │ http://localhost:3000
│ (90d TSDB) │ │ (1 dashboard)│ → lands on the dashboard
━━━━━━━━━━━━━━━━━━━━┘
クイックスタート
git clone < このリポジトリ > && cd JustRebootIt
cp .env.example .env
$EDITOR .env # UDM_URL / UDM_USERNAME / UDM_PASSWORD を設定します
$EDITOR config/targets.yml # ホームゲートウェイ IP (Comcast でない場合は ISP) を設定します
docker 構成 -d --build
次に、http://localhost:3000 を開きます。ホーム インターネットに直接アクセスします。
待ち時間ダッシュボード、ログインは必要ありません。それだけです。
遅延グラフが必要なだけで、UniFi ゲートウェイをお持ちではありませんか?参照
UDM なしで実行します。
UniFi Dream Machine Pro で行う必要があること
ローカル管理者アカウントを作成します。 UniFi OS で、[設定] → [管理者] に移動し、
「ユーザー」→「管理者の追加」を選択し、「ローカルアクセスのみに制限する」を選択します。しないでください
Ubiquiti クラウド (SSO) ログインを使用します。ローカル アカウントは、
API クライアントを使用し、クラウド認証情報をコンテナーに漏らさないようにします。
閲覧者の役割で十分です。エクスポーターは読み取り専用のみを発行します
GET /stat/* c

全部。必要な場合にのみ、完全なローカル管理者を与えてください。
ゲートウェイ URL をメモします (通常は https://192.168.1.1 )。それに加えて、
ユーザー名とパスワードを .env ファイルに追加します ( UDM_URL 、 UDM_USERNAME 、
UDM_PASSWORD )。
(オプション) 定期的な速度テストをスケジュールします。 UniFi OS → ネットワーク →
「設定」→「インターネット」→「速度テスト」。これにより、
udm_speedtest_* パネル。 WAN 遅延、スループット、CPU/メモリ、クライアント
カウントは関係なく報告されます。
他には何もありません。 UDM は自己署名 TLS 証明書を提示するため、エクスポーターは
デフォルトでは証明書の検証をスキップします ( UDM_INSECURE=true )。あなたはそうではありません
証明書をインストールするか、UDM 上のポートを開く必要があります。
ネットワークの権限と機能
これは手を振ることができない 1 つのことです。レイテンシーの測定には次のことが必要です。
ICMP エコー (ping) パケットと ICMP トレースルートの送信 (生のネットワークが必要)
ソケット。
プローバコンテナには、Linux NET_RAW 機能が付与されます。
docker-compose.yml :
プローバー:
cap_add :
- NET_RAW
これは必要な最小限の権限です。これにより、プロセスは ICMP ソケットを開くことができます。
しかしそれ以外は何も与えません。コンテナは --privileged を実行せず、実行します
distroless イメージ内の非特権ユーザーとして。 NET_RAW がデフォルトです
Docker 機能。ほとんどのホストでは、これは追加設定なしで機能します。それはリストされています
明示的に要件が表示され、強化/デフォルト拒否が存続するようにする
セットアップ。
config/targets.yml の特権: true は Docker とは無関係です
特権 — プローバーに raw ICMP ソケット (信頼できる選択) を使用するように指示します。
NET_RAW を指定した場合)、特権のないデータグラム ICMP ソケットではなく、
さらに、ホスト sysctl net.ipv4.ping_group_range を設定する必要があります。
プローブには受信ポートは必要ありません。唯一公開されているポートは
Grafana の 3000 (ポート: ["3000:3000"] )。プロメテウス (9090) と二人
輸出業者 (9430/9431) にアクセス可能

内部 Docker ネットワーク上でのみ。
アウトバウンド: プローバーはインターネット (ICMP) と LAN に到達する必要があります。
ゲートウェイ。 udm-exporter は LAN 上の UDM への HTTPS を必要とします。
ファーストホップ/トレースルートの精度 (オプション)
デフォルトのブリッジ ネットワークでは、traceroutes には 1 つまたは 2 つの追加のものが表示されます。
実際のゲートウェイに到達する前に、Docker/ホストがホップします。追加されるレイテンシーは、
ミリ秒未満で一定であるため、スパイクをマスクしませんが、必要に応じて
物理ゲートウェイから正確に開始するパス。ホスト上でプローバーを実行します。
ネットワーク:
プローバー:
network_mode : host # これを追加します。 `networks:` ブロックを削除します
そして、プローバー ジョブの Prometheus スクレイピング ターゲットを prober:9430 から変更します。
host.docker.internal:9430 (またはホストの LAN IP) に送信します。ほとんどのユーザーは必要ありません
これ。
macOS / Windows (Docker Desktop): コンテナは Linux VM 内で実行されるため、
NET_RAW と ICMP は機能しますが、network_mode: host の「ホスト」はその VM です。
Mac/PC ではありません。そこではブリッジのデフォルトが推奨されます。
このファイルはプローバーにバインドマウントされているため、編集はプローバーに有効になります。
docker compose 再起動プローバー。各ターゲットには安定した名前が付いています (安定した名前を維持してください)
履歴グラフが並んでいます)、ホスト、グループ、およびオプションのトレース
旗。グループはダッシュボードと推論を整理します。
出荷時のデフォルトでは、ゲートウェイ、Comcast のリゾルバー (75.75.75.75 /) がプローブされます。
75.75.76.76 )、およびアンカーのスプレッド (Cloudflare、Google、Quad9、Level3)。
少なくともホームゲートウェイをゲートウェイの LAN IP に編集します。オンになっていない場合
Comcast、ISP ターゲットを ISP のゲートウェイ/リゾルバーに交換します。
ファイルの先頭にあるタイミング ノブ (デフォルトを表示):
間隔 : 10 秒 # 1 つのプローブ サイクル。 ping は全体に分散されます
ping : ターゲットごと、サイクルごとに 20 件のエコー リクエスト (「スモーク」)
timeout : 2s # 応答ごとのタイムアウト (<間隔である必要があります)
特権: true 生の ICMP ソケット数 (「ネットワーク許可」を参照)

s)
トレースインターバル : 60 秒 # トレースルートは重くなります。実行頻度を減らす
トレース最大ホップ数: 30
トレースタイムアウト: 2秒
パスの検出 — 多様な短いルート (自動)
自宅から同じ方法で出発する十数の目的地を調べるのは、ほとんどの場合、
冗長: 共有セグメントに問題が発生すると、すべてが一度に急増し、
場所については何も学びません。 Discovery: ブロックはこれを修正します。定期的にそれ
候補プールをトレースし、パスの多様なサブセットをアクティブに昇格します。
プロービング — 2 番目/3 番目のホップが異なり、そのホップに到達するものを保持します。
最小ホップでの宛先 (ターゲットが近いほど、障害がより局所的に特定されます)
正確には）。選択された候補者はグループの下で調査および追跡されます。
発見された;上記の常時オンのターゲットには決して触れません。
発見：
有効 : true
間隔 : 15m # プールを再トレースし、これを頻繁に再選択します
max_targets : 6 # アクティブにプローブする多様なパスの数
max_hops : 8 # 検出中の短いトレース (初期ホップのみが重要)
max_reach_hops : 12 # これより遠い候補を無視します
候補者：
- { 名前: Cloudflare-alt、ホスト: 1.0.0.1 }
- { 名前: opendns-a、ホスト: 208.67.222.222 }
# ... 異なるオペレーター/ASN にまたがる広範なプール
候補名は一意である必要があり、常時オンのターゲットと衝突してはなりません。
ダッシュボードのアクティブな検出されたパス パネルには、現在どのパスが検出されているかが表示されます。
選択されている場合、[候補距離] にはそれぞれのホップ数が表示されます。
トリガーされた診断 - スパイク中のより詳細なテスト
断続的な問題は通常、対応する前に解消されます。の
診断: ブロックはすべてのプローブ サイクルを監視し、ターゲットが検出された瞬間を監視します。
異常で、より詳細なテストを自動的にバーストして開始し、証拠を収集します
問題がまだ発生している間:
新しいトレースルート (イベント中に取得されたパスのスナップショット)。
ターゲットへの TCP ハンドシェイク時間

— 多くの ISP が ICMP の優先順位を下げているため、
これにより、ping だけでなく実際のトラフィックが影響を受けるかどうかが証明されます。そして
DNS 解決時間 — ルックアップが遅いと、「インターネットが遅い」と見せかけます。
サイクルの RTT 中央値がターゲットのローリング ベースラインを超えると実行が開始されます。
係数と絶対マージンの両方 (つまり、小さなジッターも小さなジッターもありません)
相対的なバンプだけでトリップする場合)、または損失がしきい値を超えた場合、つまり損失がデバウンスされた場合
クールダウンのため、持続的なイベントはサイクルごとではなく 1 回トリガーされます。
診断:
有効 : true
latency_factor : 3.0 # 中央値がローリングベースラインの 3 倍を超える場合にトリガーします
latency_abs_margin : 30ms # ...そして少なくとも 30ms 以上
loss_threshold : 0.1 # または損失 >= 10% の場合
クールダウン : 60 秒 # ターゲットごとに 1 分あたり最大 1 回の実行
Baseline_alpha : 0.2 # 「通常のレイテンシー」ベースラインの EWMA 平滑化
tcp_port : 443 # TCP 遅延テスト用のハンドシェイク ポート
dns_probe : www.google.com # 実行中に名前が解決され、時間を計測される
ワーカー: 2 つの同時診断実行
すべてのトリガーは、ダッシュボード全体に赤い注釈もドロップします。
遅延と WAN のグラフに対して詳細なテストを実行すると、正確に一致する可能性があります。
AI による根本原因分析 - 「なぜこれが起こったのか?」 (オプション)
機械的診断は生の信号を収集します。これは説明になります。
イベントが発生すると、クロード エージェントがそれを調査し、
根本原因をわかりやすい言葉でダッシュボードに表示します。読み取り専用ツールが提供されており、
それらをどのように使用するかを決定します。
Prometheus をクエリする — スパイクがすべてのターゲットに一度に到達したかどうかを確認する
(アップストリーム/ISP) または 1 つのパスのみ、および WAN と並んでいるかどうか
飽和状態（バッファ膨張）。
トレースルート / DNS ルックアップ — 再プロへ

[切り捨てられた]

## Original Extract

Contribute to adamf/JustRebootIt development by creating an account on GitHub.

I love smokeping for finding jitter and latency but always wanted something that went deeper and grabbed all the related signals, so I built JustRebootIt. JRI runs a bunch of parallel pings & traceroutes, detects when a spike is happening, runs more detailed probes during the event, then can feed all that into an LLM for diagnosis. It also reads from my UDM to gather CPU/memory/etc from the device in the event that correlates. Feedback and PRs welcome!

GitHub - adamf/JustRebootIt · GitHub
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
adamf
/
JustRebootIt
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
21 Commits 21 Commits .github/ workflows .github/ workflows cmd cmd config config docker docker docs docs internal internal .env.example .env.example .gitignore .gitignore Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md docker-compose.yml docker-compose.yml go.mod go.mod go.sum go.sum View all files Repository files navigation
A self-contained, Dockerized recorder for intermittent home-internet latency
spikes — the kind where everything feels fine until a video call freezes for
three seconds, and by the time you run a speed test it's gone.
It continuously measures latency, jitter and packet loss to a spread of diverse
targets (your gateway, your ISP, and well-run public anchors), traces the
network path to find where the latency lives, and — if you have a UniFi
Dream Machine Pro — correlates all of that against WAN throughput and gateway
load. Everything lands on one Grafana dashboard designed to be screenshotted
or link-shared with a network engineer at your ISP (e.g. Comcast).
The probes are written in Go and run fully in parallel, so dozens of targets are
measured on a tight cycle without the prober itself becoming the bottleneck.
A real diagnosis from the dashboard above (Event #6): "The latency on
8.8.4.4 was not specific to Google — it was local bufferbloat on the WAN uplink
during a traffic burst… the median RTT to every anchor jumped together, so the
spike was shared across all independent paths… udm_wan_rx hit ~15.2 MB/s and
WAN latency climbed from ~10–11ms baseline to 33ms. This is the classic
bufferbloat pattern: the local uplink filled and queued packets. Confidence:
high. Recommended action: enable Smart Queues (SQM/CAKE) on the UDM Pro with
limits ~85–90% of the measured line rate — this is a local fix, not an ISP
ticket." — written automatically by the AI analysis when the event fired.
┌─────────────┐ ICMP ping + traceroute
│ prober │────────────────────────────► many targets
│ (Go) │ (gateway, ISP, anchors)
└──────┬──────┘
│ /metrics
UniFi Dream ┌──────┴──────┐
Machine Pro ─►│ udm-exporter│ WAN throughput, gateway CPU/mem,
(API) │ (Go) │ speedtest, client counts
└──────┬──────┘
│ /metrics
┌──────┴──────┐ ┌─────────────┐
│ prometheus │───────►│ grafana │ http://localhost:3000
│ (90d TSDB) │ │ (1 dashboard)│ → lands on the dashboard
└─────────────┘ └─────────────┘
Quick start
git clone < this repo > && cd JustRebootIt
cp .env.example .env
$EDITOR .env # set UDM_URL / UDM_USERNAME / UDM_PASSWORD
$EDITOR config/targets.yml # set your home-gateway IP (and ISP, if not Comcast)
docker compose up -d --build
Then open http://localhost:3000 — you land straight on the Home Internet
Latency dashboard, no login required. That's it.
Just want the latency graphs and don't have a UniFi gateway? See
Running without a UDM .
What you need to do on the UniFi Dream Machine Pro
Create a local admin account. In UniFi OS go to Settings → Admins &
Users → Add Admin , and choose "Restrict to local access only" . Do not
use your Ubiquiti cloud (SSO) login — local accounts are more reliable for an
API client and keep your cloud credentials out of the container.
A Viewer role is sufficient; the exporter only issues read-only
GET /stat/* calls. Give it full local admin only if you prefer.
Note the gateway URL — usually https://192.168.1.1 . Put it, plus the
username and password, into your .env file ( UDM_URL , UDM_USERNAME ,
UDM_PASSWORD ).
(Optional) Schedule periodic Speed Tests. UniFi OS → Network →
Settings → Internet → Speed Test . This populates the
udm_speedtest_* panels. WAN latency, throughput, CPU/memory and client
counts are reported regardless.
Nothing else. The UDM presents a self-signed TLS certificate, so the exporter
skips certificate verification by default ( UDM_INSECURE=true ). You do not
need to install a cert or open any ports on the UDM.
Network permissions & capabilities
This is the one thing that can't be hand-waved: measuring latency requires
sending ICMP echo (ping) packets and ICMP traceroutes, which need a raw network
socket.
The prober container is granted the Linux NET_RAW capability in
docker-compose.yml :
prober :
cap_add :
- NET_RAW
This is the minimum privilege needed — it lets the process open ICMP sockets
but grants nothing else. The container does not run --privileged and runs
as an unprivileged user inside a distroless image. NET_RAW is a default
Docker capability, so on most hosts this works out of the box; it's listed
explicitly so the requirement is visible and survives hardened/default-deny
setups.
privileged: true in config/targets.yml is unrelated to Docker
privilege — it tells the prober to use raw ICMP sockets (the reliable choice
given NET_RAW ) rather than unprivileged datagram-ICMP sockets, which
additionally require the host sysctl net.ipv4.ping_group_range to be set.
No inbound ports are required for probing. The only published port is
Grafana's 3000 ( ports: ["3000:3000"] ). Prometheus (9090) and the two
exporters (9430/9431) are reachable only on the internal Docker network.
Outbound: the prober needs to reach the internet (ICMP) and your LAN
gateway; the udm-exporter needs HTTPS to the UDM on your LAN.
First-hop / traceroute accuracy (optional)
With the default bridge network, traceroutes show one or two extra
Docker/host hops before reaching your real gateway. The added latency is
sub-millisecond and constant, so it does not mask spikes — but if you want
the path to start exactly at your physical gateway, run the prober on the host
network:
prober :
network_mode : host # add this; remove its `networks:` block
and change the Prometheus scrape target for the prober job from prober:9430
to host.docker.internal:9430 (or the host's LAN IP). Most users don't need
this.
macOS / Windows (Docker Desktop): containers run inside a Linux VM, so
NET_RAW and ICMP work, but the "host" for network_mode: host is that VM,
not your Mac/PC — the bridge default is recommended there.
This file is bind-mounted into the prober, so edits take effect on
docker compose restart prober . Each target has a stable name (keep it stable
so historical graphs line up), a host , a group , and an optional trace
flag. Groups organize the dashboard and your reasoning:
The shipped defaults probe your gateway, Comcast's resolvers ( 75.75.75.75 /
75.75.76.76 ), and a spread of anchors (Cloudflare, Google, Quad9, Level3).
Edit at least home-gateway to your gateway's LAN IP. If you're not on
Comcast, swap the isp targets for your ISP's gateway/resolver.
Timing knobs (defaults shown) at the top of the file:
interval : 10s # one probe cycle; pings are spread across it
pings : 20 # echo requests per target per cycle (the "smoke")
timeout : 2s # per-reply timeout (must be < interval)
privileged : true # raw ICMP sockets (see Network permissions)
trace_interval : 60s # traceroutes are heavier; run them less often
trace_max_hops : 30
trace_timeout : 2s
Path discovery — diverse, short routes (automatic)
Probing a dozen destinations that all leave your house the same way is mostly
redundant: when that shared segment hiccups, everything spikes at once and you
learn nothing about where . The discovery: block fixes this. Periodically it
traces a candidate pool, then promotes a path-diverse subset to active
probing — keeping the ones whose 2nd/3rd hops differ and that reach their
destination in the fewest hops (closer targets localize a fault more
precisely). Selected candidates are probed and traced under the group
discovered ; the always-on targets above are never touched.
discovery :
enabled : true
interval : 15m # re-trace the pool and re-select this often
max_targets : 6 # how many diverse paths to actively probe
max_hops : 8 # short traces during discovery (only early hops matter)
max_reach_hops : 12 # ignore candidates farther than this
candidates :
- { name: cloudflare-alt, host: 1.0.0.1 }
- { name: opendns-a, host: 208.67.222.222 }
# ... a broad pool spanning distinct operators/ASNs
Candidate names must be unique and must not collide with the always-on targets.
The dashboard's Active discovered paths panel shows which are currently
selected, and Candidate distance shows how many hops away each one is.
Triggered diagnostics — deeper tests during a spike
Intermittent problems are usually gone before you can react. The
diagnostics: block watches every probe cycle and, the moment a target looks
anomalous, fires a burst of deeper tests automatically — capturing evidence
while the problem is still happening:
a fresh traceroute (a path snapshot taken during the event);
a TCP handshake time to the target — because many ISPs deprioritize ICMP,
this proves whether real traffic is affected, not just pings; and
a DNS resolution time — slow lookups masquerade as "the internet is slow".
A run fires when a cycle's median RTT exceeds the target's rolling baseline by
both a factor and an absolute margin (so neither tiny jitter nor a small
relative bump alone trips it), or when loss crosses a threshold — debounced by a
cooldown so a sustained event triggers once, not every cycle.
diagnostics :
enabled : true
latency_factor : 3.0 # trigger when median > 3x the rolling baseline
latency_abs_margin : 30ms # ...and at least 30ms above it
loss_threshold : 0.1 # or when loss >= 10%
cooldown : 60s # at most one run per target per minute
baseline_alpha : 0.2 # EWMA smoothing for the "normal latency" baseline
tcp_port : 443 # handshake port for the TCP latency test
dns_probe : www.google.com # name resolved & timed during a run
workers : 2 # concurrent diagnostic runs
Every trigger also drops a red annotation across the whole dashboard, so you
can line up exactly when deeper tests fired against the latency and WAN graphs.
AI root-cause analysis — "why did this happen?" (optional)
The mechanical diagnostics gather raw signal; this turns it into an explanation.
When an event fires, a Claude agent investigates it and writes a
plain-language root cause onto the dashboard. It's given read-only tools and
decides how to use them:
query Prometheus — to see whether the spike hit every target at once
(upstream/ISP) or just one path, and whether it lined up with the WAN
saturating (bufferbloat);
traceroute / DNS lookup — to re-pro

[truncated]
