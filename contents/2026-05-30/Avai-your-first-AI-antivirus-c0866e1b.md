---
source: "https://github.com/iklobato/avai"
hn_url: "https://news.ycombinator.com/item?id=48331726"
title: "Avai – your first AI antivirus"
article_title: "GitHub - iklobato/avai: macOS / Linux host security telemetry collector with LLM threat judge and a single-page web dashboard. · GitHub"
author: "iklobato1"
captured_at: "2026-05-30T11:36:15Z"
capture_tool: "hn-digest"
hn_id: 48331726
score: 2
comments: 0
posted_at: "2026-05-30T02:11:14Z"
tags:
  - hacker-news
  - translated
---

# Avai – your first AI antivirus

- HN: [48331726](https://news.ycombinator.com/item?id=48331726)
- Source: [github.com](https://github.com/iklobato/avai)
- Score: 2
- Comments: 0
- Posted: 2026-05-30T02:11:14Z

## Translation

タイトル: Avai – 初めての AI アンチウイルス
記事のタイトル: GitHub - iklobato/avai: LLM 脅威判定機能と単一ページの Web ダッシュボードを備えた macOS / Linux ホスト セキュリティ テレメトリ コレクター。 · GitHub
説明: LLM 脅威判定機能と単一ページの Web ダッシュボードを備えた macOS / Linux ホスト セキュリティ テレメトリ コレクター。 - イクロバト/アヴァイ

記事本文:
GitHub - iklobato/avai: LLM 脅威判定機能と単一ページの Web ダッシュボードを備えた macOS / Linux ホスト セキュリティ テレメトリ コレクター。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
イクロバト
/
あります
公共
通知
変更するにはサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
release/0.1.0 ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
83 コミット 83 コミット docs/ イメージ docs/ イメージ ランディング ページ ランディング ページ src/ avai src/ avai テスト テスト .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md Dockerfile Dockerfile README.md README.md av-features.md av-features.md docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
マシン上で実際に何が実行されているかを把握します。
オープンソースのホスト テレメトリ + LLM 脅威分類子。 Docker を 1 回実行します。
avai は、macOS ではホストの 26 隅のスナップショットを作成します (Linux では 21 隅) —
プロセス、USB、永続性、ファイルの整合性、ブラウザ拡張機能、実行
イベント — それぞれの新しい発見を最大 17 の脅威インテリジェンスで充実させます
ソース (VirusTotal、MalwareBazaar、URLhaus、CISA KEV、Shodan、
IPDB、OSV、NVD などを悪用し、Claude クラス LLM にどれかを教えてもらいます。
気にする価値のあるものです。評決は次のように返されます
悪意のある/疑わしい/未知の/良性の
MITRE に準拠したカテゴリ、信頼性、および 1 行の修正。
エージェント契約、SIEM、クラウド コントロール プレーンはありません。
コンテンツ ハッシュによる重複排除 — 同じアーティファクトが LLM に二度送信されることはありません。
LLM の背後にある 17 個のプラグアンドプレイの脅威インテル ソース — .env.example を参照。キーが欠落していると、ソースが完全に無効になります。
読み取り専用の Flask + HTMX + Chart.js ダッシュボード (:8765)。
BYO キー ( ANTHROPIC_API_KEY / CLAUDE_CODE_OAUTH_TOKEN )、または litellm でサポートされているプロバイダーに切り替えます。
→ マーケティングサイトとスクリーンショット: https://getavai.com
→出典：https://github.com/iklobato/avai
:8765 の読み取り専用 Flask + HTMX + Chart.js ダッシュボード。すべてのパネルがレンダリングされます
モニターが書き込むのと同じ SQLite スナップショットから — 個別のコントロール プレーンは不要

。
一目でわかる健全性: 保存された実行、最新のサイクルのコレクター (任意の
失敗）、前回の実行以降の判定、および評決合計のドーナツ
(悪意のある/疑わしい/不明/良性)。 macOS システム整合性パネル
FileVault、ファイアウォール、ゲートキーパー、およびリモート アクセスの切り替えを表示します。コレクター
エラーには、何が失敗したかが表示されます (TCC 権限など)。そして12時間チャートの軌跡
時間の経過とともに判決が下されます。以下の所見表は、活動的で良性ではないものを示しています。
結果。
結果テーブルは、ステータス、評決、コレクター、およびカテゴリによってフィルター処理できます。
その下の「コレクターあたりの行数」は、各コレクターがどれだけのデータを収集したかを示します。
最新の実行、および最近の実行では、OK/失敗数を含む実行履歴がリストされます。
ルックバックウィンドウ。
結果を展開して、LLM の推論、具体的な修復方法を確認します。
ステップ、および判定の背後にある正確に収集されたデータ - プロセス向け
pid/ppid、完全な cmdline 、実行中のユーザー/uid、ステータス、コンテンツを意味します
重複排除に使用されたハッシュ、および最初に判定されたときと最後に確認されたとき。
tcpdump アグリゲーターはトラフィックを宛先ごとにグループ化するため、分類子は次のような理由を判断できます。
それについて: ここでは、異常な高ポートへの IPv6 接続にフラグが立てられています
C2 ビーコンの可能性があると疑わしいが、CDN、mDNS、LAN トラフィックは戻ってくる
無害 — それぞれに 1 行の「理由」が含まれています。
同じビューが、所有プロセス、ASN/地域、宛先ごとに強化されます。
交通量と各判決の根拠。
同じダッシュボード、別のホスト
別のホスト/サイクルに対して実行します。ここでは 61 回の実行と 3,426 件の判定が行われます。
疑わしい AirWatch/MDM の永続性が調査のために浮上しました。
ホスト テレメトリ - macOS では 26 個のコレクター、Linux では 21 個のコレクター、あらゆる場所のスナップショットを作成
マルウェアは隠れて存続し、自宅に電話をかけます。
プロセスと実行 - 実行中のプロセス、および起動時の execve exec イベント。
ネットワーク — アクティブな接続

ns、リスニング ポート、tcpdump フロー アグリゲーター
(宛先ごとにグループ化され、所有プロセスに関連付けられます)、DNS クエリ、インターフェイス、Wi-Fi 状態。
永続性 — 起動アイテム (LaunchDaemons/Agents、systemd ユニット、cron)、
カーネルとシステム拡張機能、MDM/構成プロファイル、インストールされているアプリ。
アクセスとアイデンティティ — 認証イベント (統合ログ/journalctl)、SSH
authorized_keys 、TCC プライバシー許可 (カメラ/マイク/位置/画面/フルディスク)、
特権設定、setuid バイナリ。
整合性と姿勢 — システムの整合性 (FileVault、ファイアウォール、ゲートキーパー、SIP、
SELinux/AppArmor/ufw…)、ファイルの整合性 (passwd/shadow/sudoers/SSH/dotfiles)、
/etc/hosts 、隔離イベント、マウント。
ハードウェア — USB および Bluetooth デバイス。
ブラウザ — インストールされているブラウザ拡張機能。
LLM 脅威分類子 — クロードクラス モデルは、すべての新しいアーティファクトにラベルを付けます
悪意のある / 疑わしい / 未知の / 良性の、MITRE に準拠したカテゴリ、
信頼性と 1 行の修正。自分の鍵を持参してください
( ANTHROPIC_API_KEY / CLAUDE_CODE_OAUTH_TOKEN ) またはリテルムでサポートされている任意のプロバイダー。
すべての判決の背後にある 17 の脅威インテリジェンス ソース - 指標 (ハッシュ、IP、ドメイン、
URL、CVE、パッケージ、OS バージョンなど）は、モデルがそれらを認識する前に強化されます。
VirusTotal · MalwareBazaar · URLhaus · ThreatFox · Feodo Tracker · AbuseIPDB ·
GreyNoise · Shodan InternetDB · CISA KEV · NVD · OSV · GitHub アドバイザリー ·
CIRCL ハッシュルックアップ · crt.sh · PhishTank · Google セーフ ブラウジング · endoflife.date
(さらに IP 地理位置情報)。キーレス ソースは常に実行されます。キー付きのものは追加時に有効になります
キーが不足していると、そのソースは完全に無効になります。結果はキャッシュされます
ソースごとの TTL を備えた SQLite。
読み取り専用ダッシュボード (Flask + HTMX + Chart.js on :8765 ) — 評決合計ドーナツ、
macOS システム整合性パネル、コレクター エラー、12 時間の判定チャ

rt、そして
すべてのセクションで検索/フィルター/並べ替え/ページネーションを備えた結果テーブル。任意の展開
モデルの推論、修正、および収集された生のデータを見つける。
30 ～ 60 秒ごとに自動更新し、新しい情報についてはトーストと音声アラートを鳴らします
悪意のある/疑わしい発見。
1 つの Docker 実行 - ダッシュボードとモニターの両方に同じイメージが表示されます。
エージェント契約、SIEM、クラウド コントロール プレーンは不要で、ホスト上で実行されます。
コンテンツ ハッシュによる重複排除 — 同じアーティファクトが LLM に二度送信されることはありません。
単なる SQLite ファイル — ダッシュボードで任意の OS 上の任意の avai.db を指定します。
ネイティブインストールも可能: pip install avai-monitor 。
以下の完全な参照: 収集されたもの ·
ダッシュボード · 脅威インテリジェンスの強化。
ログではなく回答です。すべての調査結果は平易な英語で評決とともに返されます。
信頼性、MITRE カテゴリ、具体的な修正 — クエリ言語なし、トリアージなし
スプレッドシート。
インフラストラクチャゼロ。コンテナは1つ。 SIEM なし、登録するエージェントなし、制御なし
飛行機を走らせるか、料金を支払うか。
デフォルトでは非公開です。すべてがマシン上で実行されます。自分のモデルを持ち込む
重要であり、脅威インテリジェンスの検索または LLM 呼び出しのための新しい発見だけが残されます。
あなたはオプトインしました。
ランニングコストが安い。コンテンツ ハッシュの重複除去は、各アーティファクトが 1 回だけ判断されることを意味します。
ホストは高額な請求を意味するものではありません。キャッシュされたインテル ヒットはネットワークを完全にスキップします。
実に広い。単一の判定の背後にある 26 のホスト サーフェス × 17 のインテル ソース —
エージェントなしの EDR の幅広さ。
クロスプラットフォーム。 macOS と Linux は同じツールから利用できます。
開いて、あなたのものです。 MIT ライセンス、監査可能、モデルに依存しない - 任意のモデルに交換可能
litellm プロバイダーと 1 つの環境変数。
本番環境を指しても安全です。コレクターは読むだけです。ダッシュボードは読み取り専用です。
ポータブルな歴史。全体の状態は単一の SQLite ファイルです - サーバー上でスキャンし、
ラップトップで表示、スナップをアーカイブ

熱く、時間の経過とともに変化します。
走る
コマンド
意味があるところ
ダッシュボード (デフォルト)
docker run iklob1/avai
任意のホスト - :8765 上の読み取り専用 Flask + HTMX
モニター
docker run ... iklob1/avai avai モニター ...
Linux ホストのみ — pid=host 、 network=host 、およびホスト ファイルシステムのバインド マウントが必要
イメージのデフォルトの CMD はダッシュボードです。次のコマンドをオーバーライドします。
docker run / compose レベルを使用して、代わりにモニターを実行します。ネイティブインストール
も可能です ( pip install avai-monitor 、次に avai Monitor /
avai ダッシュボード ) ですが、これは文書化されたパスではありません。
画像にはダッシュボードのヘルスチェックが含まれています。
/api/notifications/new エンドポイント — 開始 → 10 秒以内に正常に
最初の打ち上げ。 docker compose ps と docker Inspection --format '{{.State.Health.Status}}' は両方ともそれを反映します。
TL;DR — 60 秒のテスト、LLM キーなし
任意のホスト (macOS または Linux) 上で安全に最初に実行でき、特権も必要もありません。
認証情報、ホストバインドマウントなし。データが入力された DB と
緑色のダッシュボードを突くことができます。
mkdir -p ~ /.avai && cd ~ /.avai
# 1. コンテナーのビューの 1 つのスナップショットを DB に追加します
docker run --rm -v " $PWD " :/data iklob1/avai \
avai Monitor --once --no-streaming --no-judge --db /data/avai.db
#2. 提供する
docker run -d --name avai -p 8765:8765 -v " $PWD " :/data iklob1/avai
http://localhost:8765/ # macOS を開きます。 Linux 上の xdg-open
最大 14 コレクタに相当する行 ( process 、
network_connections 、 listen_ports 、 network_interfaces 、
usb_devices 、 launch_items 、installed_apps 、 mounts 、
setuid_files など) — コンテナーそのものではなくコンテナー自体を読み取ります。
上記の実行ではホスト状態をバインドマウントしないためです。本物になるために
データについては、以下の §2 / §3 に移動してください。
docker stop avai && docker rm avai で停止します。
1 — ダッシュボードのみ (macOS を含む任意のホスト)
ダッシュボードは、t によって書き込まれた SQLite データベースを読み取ります。

彼は監視します（または監視します）
前回の実行）。特権もホスト名前空間も必要ありません。
機能 — /data にマウントされた avai.db を含む単なるディレクトリ。
mkdir -p ~ /.avai && cd ~ /.avai
docker run -d \
--name avai-ダッシュボード \
-p 8765:8765 \
-v " $PWD " :/data \
iklob1/avai
http://localhost:8765/ を開きます
データベース ファイルがまだ存在しない場合、ダッシュボードは
起動時に空のスキーマが表示され、すべてのパネルが空になるまでレンダリングされます。
モニターは行を生成します。 docker stop avai-dashboard && docker rm avai-dashboard で停止します。
docker run --rm -p 9000:9000 \
-v /var/lib/avai:/data \
iklob1/avai \
avai ダッシュボード --host 0.0.0.0 --port 9000 --db /data/custom.db
画像のエントリ ポイントは利用可能です。イメージ名の後に続くものは
それに渡されました。
2 — モニター: ワンショット スキャン (Linux ホスト)
ローカル Linux ホスト上の単一サイクル。ストリーミングなし、LLM 審査員なし —
バインドマウントが正しく配線されていることを高速スモークテストします。
mkdir -p ~ /.avai && cd ~ /.avai
docker run --rm \
--pid=ホスト \
--network=ホスト \
--ユーザー 0:0 \
--cap-add SYS_PTRACE --cap-add NET_ADMIN --cap-add NET_RAW --cap-add DAC_READ_SEARCH \
-e HOST_PREFIX=/ホスト \
-v /proc:/host/proc:ro \
-v /sys:/host/sys:ro \
-v /etc:/host/etc:ro \
-v /var/lib/bluetooth:/host/var/lib/bluetooth:ro \
-v /var/lib/dpkg:/host/var/lib/dpkg:ro \
-v /usr/share/applications:/host/usr/share/applications:ro \
-v /lib/systemd:/host/lib/systemd:ro \
-v

[切り捨てられた]

## Original Extract

macOS / Linux host security telemetry collector with LLM threat judge and a single-page web dashboard. - iklobato/avai

GitHub - iklobato/avai: macOS / Linux host security telemetry collector with LLM threat judge and a single-page web dashboard. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
iklobato
/
avai
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
release/0.1.0 Branches Tags Go to file Code Open more actions menu Folders and files
83 Commits 83 Commits docs/ images docs/ images landing-page landing-page src/ avai src/ avai tests tests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md Dockerfile Dockerfile README.md README.md av-features.md av-features.md docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml View all files Repository files navigation
Know what's actually running on your machines.
Open-source host telemetry + LLM threat classifier. One docker run .
avai snapshots 26 corners of your host on macOS (21 on Linux) —
processes, USB, persistence, file integrity, browser extensions, exec
events — enriches each new finding with up to 17 threat-intel
sources (VirusTotal, MalwareBazaar, URLhaus, CISA KEV, Shodan,
AbuseIPDB, OSV, NVD, …), and lets a Claude-class LLM tell you which
ones are worth caring about. Verdicts come back as
malicious / suspicious / unknown / benign with a
MITRE-aligned category, a confidence, and a one-line remediation.
No agent contract, no SIEM, no cloud control plane.
Dedup by content hash — the same artifact is never sent to the LLM twice.
17 plug-and-play threat-intel sources behind the LLM — see .env.example ; missing keys disable a source cleanly.
Read-only Flask + HTMX + Chart.js dashboard on :8765 .
BYO key ( ANTHROPIC_API_KEY / CLAUDE_CODE_OAUTH_TOKEN ), or swap to any litellm-supported provider.
→ Marketing site & screenshots: https://getavai.com
→ Source: https://github.com/iklobato/avai
A read-only Flask + HTMX + Chart.js dashboard on :8765 . Every panel renders
from the same SQLite snapshot the monitor writes — no separate control plane.
At-a-glance health: runs stored, collectors in the latest cycle (with any
failures), judgments since the last run, and the verdict-totals donut
(malicious / suspicious / unknown / benign). The macOS System Integrity panel
surfaces FileVault, Firewall, Gatekeeper and remote-access toggles; Collector
Errors shows what failed (e.g. a TCC permission); and the 12-hour chart tracks
verdicts over time. The findings table below streams the active, non-benign
results.
The findings table is filterable by status, verdict, collector and category.
Beneath it, Rows per collector shows how much each collector pulled in the
latest run, and Recent runs lists run history with ok/failed counts and the
look-back window.
Expand any finding to see the LLM's reasoning , a concrete remediation
step, and the exact collected data behind the verdict — for a process that
means pid/ppid, the full cmdline , the running user/uid, status, the content
hash used for dedup, and when it was first judged vs. last seen.
The tcpdump aggregator groups traffic by destination so the classifier can reason
about it: here an IPv6 connection to an unusual high port is flagged
suspicious as a possible C2 beacon, while CDN, mDNS and LAN traffic come back
benign — each with a one-line "why".
The same view enriched per destination with the owning process , ASN/geo,
traffic volume, and the rationale for each verdict.
The same dashboard, another host
Run against a different host/cycle — 61 runs and 3,426 verdicts here — with
suspicious AirWatch/MDM persistence surfaced for review.
Host telemetry — 26 collectors on macOS, 21 on Linux , snapshotting every place
malware hides, persists, and phones home:
Processes & execution — running processes, and execve exec events as they fire.
Network — active connections, listening ports, a tcpdump flow aggregator
(grouped by destination, tied to the owning process), DNS queries, interfaces, Wi‑Fi state.
Persistence — launch items (LaunchDaemons/Agents, systemd units, cron),
kernel & system extensions, MDM/configuration profiles, installed apps.
Access & identity — auth events (unified log / journalctl ), SSH
authorized_keys , TCC privacy grants (camera/mic/location/screen/full‑disk),
privilege config, setuid binaries.
Integrity & posture — system integrity (FileVault, Firewall, Gatekeeper, SIP,
SELinux/AppArmor/ufw…), file integrity (passwd/shadow/sudoers/SSH/dotfiles),
/etc/hosts , quarantine events, mounts.
Hardware — USB and Bluetooth devices.
Browser — installed browser extensions.
LLM threat classifier — a Claude‑class model labels every new artifact
malicious / suspicious / unknown / benign with a MITRE‑aligned category , a
confidence , and a one‑line remediation . Bring your own key
( ANTHROPIC_API_KEY / CLAUDE_CODE_OAUTH_TOKEN ) or any litellm‑supported provider.
17 threat‑intel sources behind every verdict — indicators (hash, IP, domain,
URL, CVE, package, OS version) are enriched before the model sees them:
VirusTotal · MalwareBazaar · URLhaus · ThreatFox · Feodo Tracker · AbuseIPDB ·
GreyNoise · Shodan InternetDB · CISA KEV · NVD · OSV · GitHub Advisory ·
CIRCL hashlookup · crt.sh · PhishTank · Google Safe Browsing · endoflife.date
(plus IP geolocation). Keyless sources always run; keyed ones enable when you add
the key, and a missing key disables that source cleanly. Results are cached in
SQLite with per‑source TTLs.
Read‑only dashboard (Flask + HTMX + Chart.js on :8765 ) — verdict‑totals donut,
macOS system‑integrity panel, collector errors, a 12‑hour verdict chart, and a
findings table with search / filter / sort / pagination on every section. Expand any
finding for the model's reasoning, the fix, and the raw collected data.
Auto‑refreshes every 30–60 s with toast + audio alerts on new
malicious/suspicious findings.
One docker run — the same image is both the dashboard and the monitor.
No agent contract, no SIEM, no cloud control plane — it runs on your host.
Dedup by content hash — the same artifact is never sent to the LLM twice.
Just a SQLite file — point the dashboard at any avai.db , on any OS.
Native install too: pip install avai-monitor .
Full reference below: What's collected ·
Dashboard · Threat‑intel enrichment .
Answers, not logs. Every finding comes back in plain English with a verdict,
a confidence, a MITRE category, and a concrete fix — no query language, no triage
spreadsheet.
Zero infrastructure. One container. No SIEM, no agents to enroll, no control
plane to run or pay for.
Private by default. Everything runs on your machine; you bring your own model
key, and only new findings ever leave — for a threat‑intel lookup or the LLM call
you opted into.
Cheap to run. Content‑hash dedup means each artifact is judged once — a busy
host doesn't mean a big bill — and cached intel hits skip the network entirely.
Genuinely broad. 26 host surfaces × 17 intel sources behind a single verdict —
the breadth of an EDR without the agent.
Cross‑platform. macOS and Linux from the same tool.
Open and yours. MIT‑licensed, auditable, and model‑agnostic — swap to any
litellm provider with one env var.
Safe to point at production. Collectors only read ; the dashboard is read‑only.
Portable history. The whole state is a single SQLite file — scan on a server,
view on your laptop, archive a snapshot, diff over time.
Run
Command
Where it makes sense
Dashboard (default)
docker run iklob1/avai
any host — read-only Flask + HTMX on :8765
Monitor
docker run ... iklob1/avai avai monitor ...
Linux hosts only — needs pid=host , network=host , and host filesystem bind-mounts
The image's default CMD is the dashboard. Override the command at
docker run / compose level to run the monitor instead. Native install
is also possible ( pip install avai-monitor , then avai monitor /
avai dashboard ) but is not the documented path.
The image carries a HEALTHCHECK against the dashboard's
/api/notifications/new endpoint — starting → healthy in ~10 s on
first launch. docker compose ps and docker inspect --format '{{.State.Health.Status}}' will both reflect it.
TL;DR — 60-second test, no LLM key
A safe first run on any host (macOS or Linux), no privileges, no
credentials, no host bind-mounts. Produces a populated DB and a
green dashboard you can poke at.
mkdir -p ~ /.avai && cd ~ /.avai
# 1. populate the DB with one snapshot of the container's view
docker run --rm -v " $PWD " :/data iklob1/avai \
avai monitor --once --no-streaming --no-judge --db /data/avai.db
# 2. serve it
docker run -d --name avai -p 8765:8765 -v " $PWD " :/data iklob1/avai
open http://localhost:8765/ # macOS; xdg-open on Linux
You'll see ~14 collectors' worth of rows ( processes ,
network_connections , listening_ports , network_interfaces ,
usb_devices , launch_items , installed_apps , mounts ,
setuid_files , etc.) — read off the container itself rather than the
host, since the run above doesn't bind-mount host state. To get real
data, jump to §2 / §3 below.
Stop with docker stop avai && docker rm avai .
1 — Dashboard only (any host, including macOS)
The dashboard reads a SQLite database written by the monitor (or by a
previous run). It needs no privileges, no host namespace, no
capabilities — just a directory containing avai.db mounted at /data .
mkdir -p ~ /.avai && cd ~ /.avai
docker run -d \
--name avai-dashboard \
-p 8765:8765 \
-v " $PWD " :/data \
iklob1/avai
open http://localhost:8765/
If the database file doesn't exist yet, the dashboard creates an
empty schema on launch and every panel renders empty until the
monitor produces rows. Stop with docker stop avai-dashboard && docker rm avai-dashboard .
docker run --rm -p 9000:9000 \
-v /var/lib/avai:/data \
iklob1/avai \
avai dashboard --host 0.0.0.0 --port 9000 --db /data/custom.db
The image entry point is avai ; anything after the image name is
passed to it.
2 — Monitor: one-shot scan (Linux host)
A single cycle on the local Linux host. No streaming, no LLM judge —
fast smoke test that the bind mounts are wired right.
mkdir -p ~ /.avai && cd ~ /.avai
docker run --rm \
--pid=host \
--network=host \
--user 0:0 \
--cap-add SYS_PTRACE --cap-add NET_ADMIN --cap-add NET_RAW --cap-add DAC_READ_SEARCH \
-e HOST_PREFIX=/host \
-v /proc:/host/proc:ro \
-v /sys:/host/sys:ro \
-v /etc:/host/etc:ro \
-v /var/lib/bluetooth:/host/var/lib/bluetooth:ro \
-v /var/lib/dpkg:/host/var/lib/dpkg:ro \
-v /usr/share/applications:/host/usr/share/applications:ro \
-v /lib/systemd:/host/lib/systemd:ro \
-v

[truncated]
