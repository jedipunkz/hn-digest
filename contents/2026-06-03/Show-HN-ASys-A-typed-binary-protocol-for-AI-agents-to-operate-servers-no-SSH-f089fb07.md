---
source: "https://github.com/vincentping/asys"
hn_url: "https://news.ycombinator.com/item?id=48369976"
title: "Show HN: ASys – A typed binary protocol for AI agents to operate servers(no SSH)"
article_title: "GitHub - vincentping/asys: Agentic System Interface — a typed binary protocol for AI agents to control Linux systems · GitHub"
author: "vincentping"
captured_at: "2026-06-03T00:46:21Z"
capture_tool: "hn-digest"
hn_id: 48369976
score: 2
comments: 0
posted_at: "2026-06-02T13:24:22Z"
tags:
  - hacker-news
  - translated
---

# Show HN: ASys – A typed binary protocol for AI agents to operate servers(no SSH)

- HN: [48369976](https://news.ycombinator.com/item?id=48369976)
- Source: [github.com](https://github.com/vincentping/asys)
- Score: 2
- Comments: 0
- Posted: 2026-06-02T13:24:22Z

## Translation

タイトル: Show HN: ASys – AI エージェントがサーバーを操作するための型付きバイナリ プロトコル (SSH なし)
記事のタイトル: GitHub - vincentping/asys: エージェント システム インターフェイス — AI エージェントが Linux システムを制御するための型付きバイナリ プロトコル · GitHub
説明: エージェント システム インターフェイス — AI エージェントが Linux システムを制御するための型付きバイナリ プロトコル - vincentping/asys

記事本文:
GitHub - vincentping/asys: エージェント システム インターフェイス — AI エージェントが Linux システムを制御するための型付きバイナリ プロトコル · GitHub
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
ヴィンセントピング
/
アシス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲート

イオンオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
17 コミット 17 コミット デプロイ デプロイ ドキュメント ドキュメント サンプル サンプル sdk sdk src/ asyd src/ asyd テスト テスト ツール/ クライアント ツール/ クライアント .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE Makefile Makefile README.md README.md README.zh.md README.zh.md すべてのファイルを表示リポジトリ ファイルのナビゲーション
ASys — エージェント システム インターフェイス
AI エージェント用のバイナリ システム インターフェイス プロトコル —
ポート 7816、ゼロシェル解析、決定論的セマンティクス。
SSH は人間のために設計されました。エージェントには端末は必要ありません。
AI エージェントが ps aux | を実行するときSSH 経由で grep nginx を実行すると、OS、ロケール、ツールのバージョンによって異なる自由形式のテキストが解析されます。ただし、エージェントが ASys の SYS_PROCS 命令を呼び出すと、固定の 44 バイトのバイナリ フレーム (合計プロセス数、上位 5 つの PID、CPU%、メモリ%、ステータス フラグ) を受け取ります。これらは、型付きで曖昧さがなく、すべてのノードで同じです。
ASys は実験です。AI エージェント専用のシステム インターフェイスを第一原理から設計したらどうなるでしょうか?テキストの代わりにバイナリフレーム。コマンドごとのセッションではなく、長期間存続する暗号化された接続。広範な SSH アクセスの代わりに、命令レベルの機能が許可されます。シェル履歴の代わりに組み込まれた監査証跡。
これは SSH、Ansible、または Kubernetes オペレーターに代わるものではありません。これらのツールは、対象ユーザー (人間とオーケストレーション パイプライン) に適しています。 ASys は、オペレーターが AI エージェントであり、そのために最初から設計されたインターフェイスが必要な場合の追加オプションです。
完全な設計理論的根拠と、ASys がエージェント インフラストラクチャ環境のどこに適合するかを理解するには、ホワイトペーパーから始めてください。
AI エージェント (LLM)
│
│ ツール呼び出し
▼
Python SDK (~/.asys/id_curve25519)
│
│ Noise_IK_25519_ChaChaPoly_BLAKE2b
│ TCPポート78

16
▼
asyd (C デーモン、依存関係なし)
│
│ POSIX システムコール
▼
Linux
命令セット
コア ISA (0x00–0x0F) — 読み取り専用、副作用なし
INS
名前
応答
説明
0x00
SYS_CAPS
36B
静的機能: CPU、RAM、ディスク、ISA ビットマップ
0x01
SYS_HELLO
18B
ノードUID + ナノ秒のタイムスタンプ
0x02
SYS_STATUS
23B
平均負荷、CPU%、使用可能な RAM、ディスク、ネットワーク、RPI
0x03
SYS_PROCS
44B
合計プロセス + CPU 別上位 5 (PID、CPU%、MEM%)
プロトコル制御 (0x10 ～ 0x1F)
INS
名前
応答
説明
0x11
TASK_QUERY
3B
Task_Handle による非同期タスクのステータスのポーリング
標準 ISA — プロセス制御 (0x20–0x2F) — 署名済み、高度な機能が必要
INS
名前
応答
説明
0x20
PROC_THROTTLE
SWのみ
PID によるプロセスの SIGSTOP / SIGCONT
0x21
NET_ISOLATE
—
プロセスネットワークアクセスを隔離する(計画中)
0x22
SVC_RESTART
6B
名前付き systemd サービスを再起動します (非同期)
測定された RTT (ノイズ IK 暗号化、VirtualBox 上の RHEL、Windows Python クライアント):
Linux (RHEL/Fedora/Ubuntu/Debian)、x86_64
Python 3.8+ (pip install ノイズプロトコル暗号化付き) (クライアントのみ)
git clone https://github.com/vincentping/asys
CD ASSY
作る
sudo bin/asyd
asyd は TCP 7816 をリッスンします。最初の起動時にキー ペアを生成します。
/etc/asyd/id_curve25519 。
# 例: カスタム ポート、ローカルのみ
sudo bin/asyd --port 8816 --listen 127.0.0.1
# 例: デバッグ モード (フォアグラウンド、詳細)
sudo bin/asyd --debug
エージェントを登録する
# クライアント マシン上で — エージェント ID を生成します
python3ツール/client/asys_keygen.py
# サーバー上 — エージェントの公開キーをホワイトリストに追加します
echo " <pubkey_hex> " | sudo tee -a /etc/asyd/authorized_agents
エージェントの公開鍵は asys_keygen.py によって出力されます。また、次の場所でクライアント キー ペアも生成します。
~/.asys/id_curve25519 は、Noise IK ハンドシェイク中に使用されます。
/etc/asyd/authorized_agents にないエージェントからの接続は、

SW=0x6982 で拒否されました。
最初の接続時に、クライアントはサーバーの公開キーのフィンガープリントを確認するように求めます。
それを ~/.asys/known_hosts に保存します。以下の「最初の接続」を参照してください。
再起動せずにホワイトリストをリロードします (既存の接続は影響を受けません)。
sudo kill -HUP $( pidof asyd )
最初の接続
最初の接続時に、クライアントはサーバーの公開キーのフィンガープリントを表示します。
ASys サーバーのフィンガープリント (SHA256): a3:f1:2c:...
ローカルホスト:7816 に接続しますか? [はい/いいえ]: はい
フィンガープリントは ~/.asys/known_hosts に保存されました
以降の接続では、指紋が自動的に検証されます。不一致により中止される
接続 - SSH known_hosts と同じモデル。
すべてのクライアント スクリプトは任意のマシン (Windows を含む) 上で実行され、リモートに接続します。
ネットワーク上の asyd。 SSH はありません。シェルはありません。署名付きバイナリ命令
暗号化されたチャネル。
asyd に接続し、Noise IK ハンドシェイクを完了し、4 つすべてを実行します。
コア ISA 命令 (SYS_CAPS、SYS_HELLO、SYS_STATUS、SYS_PROCS) と実行
キャッシュタイミングテスト。
python3 例/client_core_isa.py < サーバー IP >
ASys テスト クライアント — localhost:7816
=================================================
ローカルホストに接続しました:7816
握手OK
── SYS_CAPS (0x00) ─--------------------------------------------------------
コアビットマップ: 0x0002000F active=['0x00', '0x01', '0x02', '0x03', '0x11']
外部ビットマップ: 0x00000005
プロトコル: v1.0
カーネルハッシュ: 0x06060057
CPU: 8 アーチ = x86_64
RAM: 15660 MB スワップ = 4096 MB
ルートパーティション: 1031018 MB fs=ext4
RPI タイプ: NATIVE_KERNEL (PSI)
SW: 0x9000 OK
── SYS_HELLO (0x01) ─--------------------------------------------------------
マジック：「ASYS」
ノードUID: 0xFCAB032F
サーバーのタイムスタンプ: 1779912196.304 秒 (1779912196304227300 ナノ秒)
RTT: 0.28ミリ秒
SW:0x9

000 OK
── SYS_STATUS (0x02) ─────────────
平均負荷: 1m=0.0 5m=0.0
CPU使用率: 0%
利用可能なメモリ: 14540 MB
ルートディスク: 0% 使用 IO 待機 = 0 ミリ秒
ネットワーク: 入力=0 Mbps 出力=0 Mbps
RPI: 0/100
RTT: 51.05 ミリ秒
SW: 0x9000 OK
── SYS_PROCS (0x03) ─--------------------------------------------------------
総プロセス: 48
CPU 別トップ 5 (ライフタイム シェア):
[0] PID=5964 CPU= 0.07% MEM= 0% flags=特権
[1] PID=433 CPU= 0.05% MEM= 0% フラグ=ゾンビ
[2] PID=1 CPU= 0.02% MEM= 0% flags=特権
[3] PID=6 CPU= 0.00% MEM= 0% flags=特権
[4] PID=80 CPU= 0.00% MEM= 0% flags=特権
RTT: 2.87ミリ秒
SW: 0x9000 OK
── キャッシュタイミングテスト ─────────────
6ラウンド、200ms間隔
# SYS_STATUS SYS_PROCS
──── ────── ──────
0 0.61ms 2.52ms
1 0.85ms 2.92ms
2 0.95ms 2.47ms
3 52.25ms 2.68ms
4 0.95ms 2.68ms
5 1.21ms 2.70ms
=================================================
完了しました。
2. SVC_RESTART — 非同期命令パターン
SVC_RESTART 命令を送信し、 Task_Handle を受信してからポーリングします。
サービスの再起動が完了するまで TASK_QUERY を続けます。
python3 例/client_svc_restart.py < サーバー IP > 7816 sshd
ASys フェーズ 3 E2E テスト — localhost:7816
サービス: sshd
=================================================
握手OK
► SVC_RESTART("sshd")
SW: 0x9000 OK
タスクハンドル: 0x001D7D05
RTT: 1.0ミリ秒
► TASK_QUERY ポーリング (最大 30 秒) ...
[ 1 秒] ステータス = 成功 ✓
=================================================
PASS SVC_RESTART("sshd") が Status=Success で完了しました
3. マルチエージェント — 同時セッションの分離
4 人の独立したエージェントを生成します

同時にセッションごとのロックを検証します。
同時読み取りはインターリーブせず、クロスセッションの TASK_QUERY でハンドルがリークされません
また、突然の切断が他のセッションに影響を与えることはありません。
python3 の例/client_multi_agent.py < サーバー IP >
=== ASys エージェント シミュレーター — localhost:7816 ===
セッションごとのロックの正確性を検証します (フェーズ 4 P1)
[シナリオ 1: 同時コア ISA — 4 エージェント × 5 SYS_STATUS]
PASS Agent-A2: 5× SYS_STATUS すべて SW=0x9000
PASS Agent-A3: 5× SYS_STATUS すべて SW=0x9000
PASS Agent-A1: 5× SYS_STATUS すべて SW=0x9000
PASS Agent-A4: 5× SYS_STATUS すべて SW=0x9000
[シナリオ 2: セッション間の TASK_QUERY 分離]
PASS Agent-B1: SVC_RESTART crond → SW=0x9000
PASS Agent-B1: 有効な Task_Handle を取得しました
PASS Agent-B2 が B1 のハンドルをクエリ → Status=0xFF (情報漏洩なし)
PASS Agent-B1 が自身のハンドルをクエリ → ステータス != 0xFF
[シナリオ 3: 同時 SVC_RESTART — 3 つのエージェント、3 つの異なるサービス]
PASS Agent-C1: SVC_RESTART crond → SW=0x9000 + ハンドル
PASS Agent-C2: SVC_RESTART rsyslog → SW=0x9000 + ハンドル
PASS Agent-C3: SVC_RESTART sshd → SW=0x9000 + ハンドル
PASS 3 つの Task_Handles はすべて別個です
FAIL Agent-C1: crond → ステータス=成功 ✓ (ステータス=失敗)
FAIL Agent-C2: rsyslog → ステータス=成功 ✓ (ステータス=失敗)
PASS Agent-C3: sshd → ステータス=成功 ✓
[シナリオ 4: 回復力の切断 — Agent-D が突然切断される]
PASS Agent-D-stable: 中断前の SYS_HELLO → 0x9000
PASS Agent-D-bad: 突然の切断 (ハンドシェイクなし)
PASS Agent-D-stable: 中断後の SYS_STATUS → 0x9000
PASS Agent-D-new: 中断後の新規接続 → SYS_HELLO 0x9000
=================================================
概要: 17 件が合格、2 件が不合格
いくつかのテストが失敗しました - 上記の出力を確認してください
シナリオ 3 の失敗に関する注意: Status=Failed for crond / rsyslog の意味
systemctl 再起動 retu

rned 非ゼロ — それらのサービスがインストールされていないか、
テストノードで実行中。 ASys プロトコル パス (ディスパッチ → フォーク → SIGCHLD →
ハンドル更新) は、関係なく正しく実行されます。ステータス=失敗は予想通りです
基礎となるシステム操作が失敗した場合の応答。
デモでは 2 台のマシンを使用して、ASys の実際の用途を示します: リモート オペレーター
ネットワーク経由で Linux ノードを監視および制御します。
サーバー (RHEL/Linux) 上で — CPU 占有を開始します。
python3 の例/server_cpu_hog.py
CPU 占有開始 PID=2644
Ctrl+C を押して停止します。
クライアント (Windows などの任意のマシン) で、接続してスロットルします。
python3 例/client_proc_throttle.py < サーバー IP >
ASys スロットル クライアント — <サーバー IP>:7816
=================================================
握手OK
CPU=100% 負荷1m=3.2 RPI=87/100
# PID CPU% MEM%
--- -------- ------ -----
1 2644 99.87% 0%
2 1281 9.09% 0%
3 1 0.00% 1%
4 9 0.00% 0%
5 17 0.00% 0%
スロットルする # または PID を選択します (キャンセルするには空白): 1
#1を選択 → PID 2644
PROC_THROTTLE 停止 PID=2644 ...
SW=0x9000 OK
CPU が安定するまで 2 秒待っています...
CPU=0% (100%、Δ=-100%)
PID 2644 が一時停止されました。 Ctrl+C を押して復元または終了します。
^C
PID 2644を復元しますか? (はい/いいえ) [いいえ]: はい
PROC_THROTTLE CONT PID=2644 SW=0x9000 OK
クライアントは Windows 上で実行されます。調整されているプロセスは RHEL 上にあります。
SSH はありません。シェルはありません。暗号化されたチャネルを介した 1 つの署名付きバイナリ命令。
トランスポート: ノイズ IK (Curve25519 + ChaCha20-Poly1305 + BLAKE2b) — 1-RTT 相互
認証。パスワードはありません。証明書はありません。 CAはいません。セッション内容が前進しました
機密性 (セッション キーはハンドシェイクごとの一時キーから派生します。侵害します)
静的秘密キーは過去のセッションの内容を公開しません)。既知の境界:
エージェントの静的公開キーは暗号化されています。

[切り捨てられた]

## Original Extract

Agentic System Interface — a typed binary protocol for AI agents to control Linux systems - vincentping/asys

GitHub - vincentping/asys: Agentic System Interface — a typed binary protocol for AI agents to control Linux systems · GitHub
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
vincentping
/
asys
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
17 Commits 17 Commits deploy deploy docs docs examples examples sdk sdk src/ asyd src/ asyd tests tests tools/ client tools/ client .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE Makefile Makefile README.md README.md README.zh.md README.zh.md View all files Repository files navigation
ASys — Agentic System Interface
The binary system interface protocol for AI Agents —
port 7816, zero shell parsing, deterministic semantics.
SSH was designed for humans. Agents don't need a terminal.
When an AI agent runs ps aux | grep nginx over SSH, it parses free-form text that varies by OS, locale, and tool version. But when agents call ASys's SYS_PROCS instruction, they receive a fixed 44-byte binary frame: total process count, top-5 PIDs, CPU%, memory%, status flags — typed, unambiguous, the same on every node.
ASys is an experiment: what if you designed a system interface specifically for AI agents, from first principles? Binary frames instead of text. A long-lived encrypted connection instead of per-command sessions. Instruction-level capability grants instead of broad SSH access. Built-in audit trail instead of shell history.
It's not a replacement for SSH, Ansible, or Kubernetes operators — those tools are well-suited for their intended users (humans and orchestration pipelines). ASys is an additional option for the case where the operator is an AI agent and you want an interface designed for that from the ground up.
To understand the full design rationale and where ASys fits in the agent infrastructure landscape, start with the whitepaper .
AI Agent (LLM)
│
│ Tool calls
▼
Python SDK (~/.asys/id_curve25519)
│
│ Noise_IK_25519_ChaChaPoly_BLAKE2b
│ TCP port 7816
▼
asyd (C daemon, zero dependencies)
│
│ POSIX syscalls
▼
Linux
Instruction Set
Core ISA (0x00–0x0F) — read-only, zero side effects
INS
Name
Response
Description
0x00
SYS_CAPS
36B
Static capabilities: CPU, RAM, disk, ISA bitmap
0x01
SYS_HELLO
18B
Node UID + nanosecond timestamp
0x02
SYS_STATUS
23B
Load avg, CPU%, available RAM, disk, network, RPI
0x03
SYS_PROCS
44B
Total procs + top-5 by CPU (PID, CPU%, MEM%)
Protocol Control (0x10–0x1F)
INS
Name
Response
Description
0x11
TASK_QUERY
3B
Poll async task status by Task_Handle
Standard ISA — Process Control (0x20–0x2F) — signed, require elevated capabilities
INS
Name
Response
Description
0x20
PROC_THROTTLE
SW only
SIGSTOP / SIGCONT a process by PID
0x21
NET_ISOLATE
—
Isolate process network access (planned)
0x22
SVC_RESTART
6B
Restart a named systemd service (async)
Measured RTT (Noise IK encrypted, RHEL on VirtualBox, Windows Python client):
Linux (RHEL/Fedora/Ubuntu/Debian), x86_64
Python 3.8+ with pip install noiseprotocol cryptography (client only)
git clone https://github.com/vincentping/asys
cd asys
make
sudo bin/asyd
asyd listens on TCP 7816. On first start it generates a key pair at
/etc/asyd/id_curve25519 .
# Example: custom port, local-only
sudo bin/asyd --port 8816 --listen 127.0.0.1
# Example: debug mode (foreground, verbose)
sudo bin/asyd --debug
Register an agent
# On the client machine — generate agent identity
python3 tools/client/asys_keygen.py
# On the server — add the agent's public key to the whitelist
echo " <pubkey_hex> " | sudo tee -a /etc/asyd/authorized_agents
The agent public key is printed by asys_keygen.py . It also generates the client key pair at
~/.asys/id_curve25519 used during the Noise IK handshake.
Connections from agents not in /etc/asyd/authorized_agents are rejected with SW=0x6982.
On first connection, the client will prompt you to confirm the server's public key fingerprint
and save it to ~/.asys/known_hosts — see First connection below.
Reload the whitelist without restarting (existing connections are not affected):
sudo kill -HUP $( pidof asyd )
First connection
On first connect, the client displays the server's public key fingerprint:
ASys server fingerprint (SHA256): a3:f1:2c:...
Connect to localhost:7816? [yes/no]: yes
Fingerprint saved to ~/.asys/known_hosts
Subsequent connections verify the fingerprint automatically. A mismatch aborts
the connection — same model as SSH known_hosts .
All client scripts run on any machine (including Windows) and connect to a remote
asyd over the network. No SSH. No shell. Signed binary instructions over an
encrypted channel.
Connects to asyd , completes the Noise IK handshake, then exercises all four
Core ISA instructions (SYS_CAPS, SYS_HELLO, SYS_STATUS, SYS_PROCS) and runs
a cache timing test.
python3 examples/client_core_isa.py < server-ip >
ASys Test Client — localhost:7816
====================================================
Connected to localhost:7816
Handshake OK
── SYS_CAPS (0x00) ─────────────────────────────────
Core bitmap: 0x0002000F active=['0x00', '0x01', '0x02', '0x03', '0x11']
Ext bitmap: 0x00000005
Protocol: v1.0
Kernel hash: 0x06060057
CPUs: 8 arch=x86_64
RAM: 15660 MB swap=4096 MB
Root partition: 1031018 MB fs=ext4
RPI type: NATIVE_KERNEL (PSI)
SW: 0x9000 OK
── SYS_HELLO (0x01) ────────────────────────────────
Magic: 'ASYS'
Node UID: 0xFCAB032F
Server timestamp: 1779912196.304 s (1779912196304227300 ns)
RTT: 0.28 ms
SW: 0x9000 OK
── SYS_STATUS (0x02) ───────────────────────────────
Load avg: 1m=0.0 5m=0.0
CPU usage: 0%
Mem available: 14540 MB
Root disk: 0% used IO wait=0 ms
Network: in=0 Mbps out=0 Mbps
RPI: 0/100
RTT: 51.05 ms
SW: 0x9000 OK
── SYS_PROCS (0x03) ────────────────────────────────
Total processes: 48
Top 5 by CPU (lifetime share):
[0] PID=5964 CPU= 0.07% MEM= 0% flags=Privileged
[1] PID=433 CPU= 0.05% MEM= 0% flags=Zombie
[2] PID=1 CPU= 0.02% MEM= 0% flags=Privileged
[3] PID=6 CPU= 0.00% MEM= 0% flags=Privileged
[4] PID=80 CPU= 0.00% MEM= 0% flags=Privileged
RTT: 2.87 ms
SW: 0x9000 OK
── Cache Timing Test ───────────────────────────────
6 rounds, 200ms interval
# SYS_STATUS SYS_PROCS
──── ──────────── ────────────
0 0.61ms 2.52ms
1 0.85ms 2.92ms
2 0.95ms 2.47ms
3 52.25ms 2.68ms
4 0.95ms 2.68ms
5 1.21ms 2.70ms
====================================================
Done.
2. SVC_RESTART — async instruction pattern
Sends a SVC_RESTART instruction, receives a Task_Handle , then polls
TASK_QUERY until the service restart completes.
python3 examples/client_svc_restart.py < server-ip > 7816 sshd
ASys Phase 3 E2E Test — localhost:7816
Service: sshd
====================================================
Handshake OK
► SVC_RESTART("sshd")
SW: 0x9000 OK
Task_Handle: 0x001D7D05
RTT: 1.0 ms
► TASK_QUERY polling (max 30s) ...
[ 1s] Status = Success ✓
====================================================
PASS SVC_RESTART("sshd") completed with Status=Success
3. Multi-agent — concurrent session isolation
Spawns four independent agents concurrently to validate per-session locking:
concurrent reads don't interleave, cross-session TASK_QUERY leaks no handle
information, and an abrupt disconnect does not affect other sessions.
python3 examples/client_multi_agent.py < server-ip >
=== ASys Agent Simulator — localhost:7816 ===
Validates per-session lock correctness (Phase 4 P1)
[Scenario 1: Concurrent Core ISA — 4 agents × 5 SYS_STATUS]
PASS Agent-A2: 5× SYS_STATUS all SW=0x9000
PASS Agent-A3: 5× SYS_STATUS all SW=0x9000
PASS Agent-A1: 5× SYS_STATUS all SW=0x9000
PASS Agent-A4: 5× SYS_STATUS all SW=0x9000
[Scenario 2: Cross-session TASK_QUERY isolation]
PASS Agent-B1: SVC_RESTART crond → SW=0x9000
PASS Agent-B1: obtained valid Task_Handle
PASS Agent-B2 querying B1's handle → Status=0xFF (no info leak)
PASS Agent-B1 querying own handle → Status != 0xFF
[Scenario 3: Concurrent SVC_RESTART — 3 agents, 3 different services]
PASS Agent-C1: SVC_RESTART crond → SW=0x9000 + handle
PASS Agent-C2: SVC_RESTART rsyslog → SW=0x9000 + handle
PASS Agent-C3: SVC_RESTART sshd → SW=0x9000 + handle
PASS All 3 Task_Handles are distinct
FAIL Agent-C1: crond → Status=Success ✓ (Status=Failed)
FAIL Agent-C2: rsyslog → Status=Success ✓ (Status=Failed)
PASS Agent-C3: sshd → Status=Success ✓
[Scenario 4: Disconnect resilience — Agent-D abruptly disconnects]
PASS Agent-D-stable: SYS_HELLO before disrupt → 0x9000
PASS Agent-D-bad: abrupt disconnect (no handshake)
PASS Agent-D-stable: SYS_STATUS after disrupt → 0x9000
PASS Agent-D-new: fresh connect after disrupt → SYS_HELLO 0x9000
====================================================
Summary: 17 passed, 2 failed
SOME TESTS FAILED — check output above
Note on Scenario 3 failures: Status=Failed for crond / rsyslog means
systemctl restart returned non-zero — those services are not installed or not
running on the test node. The ASys protocol path (dispatch → fork → SIGCHLD →
handle update) is exercised correctly regardless; Status=Failed is the expected
response when the underlying system operation fails.
The demo uses two machines to show what ASys is actually for: a remote operator
observing and controlling a Linux node over the network.
On the server (RHEL/Linux) — start a CPU hog:
python3 examples/server_cpu_hog.py
CPU hog started PID=2644
Press Ctrl+C to stop.
On the client (any machine, e.g. Windows) — connect and throttle:
python3 examples/client_proc_throttle.py < server-ip >
ASys Throttle Client — <server-ip>:7816
====================================================
Handshake OK
CPU=100% load1m=3.2 RPI=87/100
# PID CPU% MEM%
--- -------- ------ -----
1 2644 99.87% 0%
2 1281 9.09% 0%
3 1 0.00% 1%
4 9 0.00% 0%
5 17 0.00% 0%
Select # or PID to throttle (blank to cancel): 1
Selected #1 → PID 2644
PROC_THROTTLE STOP PID=2644 ...
SW=0x9000 OK
Waiting 2 s for CPU to settle...
CPU=0% (was 100%, Δ=-100%)
PID 2644 paused. Press Ctrl+C to restore or exit.
^C
Restore PID 2644? (yes/no) [no]: yes
PROC_THROTTLE CONT PID=2644 SW=0x9000 OK
The client runs on Windows. The process being throttled is on RHEL.
No SSH. No shell. One signed binary instruction over an encrypted channel.
Transport: Noise IK (Curve25519 + ChaCha20-Poly1305 + BLAKE2b) — 1-RTT mutual
authentication. No passwords. No certificates. No CA. Session content has forward
secrecy (session keys are derived from per-handshake ephemeral keys; compromising
the static private key does not expose past session content). Known boundary: the
agent's static public key is encrypted u

[truncated]
