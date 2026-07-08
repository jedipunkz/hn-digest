---
source: "https://github.com/instavm/tarit"
hn_url: "https://news.ycombinator.com/item?id=48829512"
title: "Show HN: Tarit – Self-host sandbox cloud and hypervisor for AI agents"
article_title: "GitHub - instavm/tarit: A hypervisor and sandbox cloud for self-hosted AI agents and RL · GitHub"
author: "mkagenius"
captured_at: "2026-07-08T10:08:07Z"
capture_tool: "hn-digest"
hn_id: 48829512
score: 6
comments: 0
posted_at: "2026-07-08T09:11:25Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Tarit – Self-host sandbox cloud and hypervisor for AI agents

- HN: [48829512](https://news.ycombinator.com/item?id=48829512)
- Source: [github.com](https://github.com/instavm/tarit)
- Score: 6
- Comments: 0
- Posted: 2026-07-08T09:11:25Z

## Translation

タイトル: Show HN: Tarit – AI エージェント用のセルフホスト サンドボックス クラウドおよびハイパーバイザー
記事タイトル: GitHub - instavm/tarit: セルフホスト AI エージェントと RL のためのハイパーバイザーとサンドボックス クラウド · GitHub
説明: セルフホスト型 AI エージェントおよび RL 用のハイパーバイザーおよびサンドボックス クラウド - instavm/tarit
HN テキスト: AI エージェントおよび RL 環境を実行するためにゼロから構築されたハイパーバイザーとして Tarit を構築しました。これは、rust-vmm に基づいており、firecracker の代替として使用できます。 Firecracker は、主にサーバーレス コンピューティングのさまざまなニーズに対応するために構築されているため、VM 操作を一時停止せずにライブ スナップショットのようなプリミティブを備えていません。また、microVM の配置を処理し、HA を使用してクラスターを作成し、VM のウォーム プールを維持し、ネットワークと監視の設定を処理する基本的なオーケストレーターも提供します。メタル インスタンスでのベンチマークでは、プールから VM を取得し、コード p99 を 35 ミリ秒で実行することが示されており、現在入手できる中で最も高速な VM ベースのサンドボックスの 1 つとなります。小さなサンドボックスのスナップショットを作成し、80 ミリ秒以内に再開することもできます。コーディング エージェントにドキュメントを提供することで、ネストされた virt 対応マシンを使用してクラウド上で、またはより良いベアメタル上で、このマルチノード クラスターを簡単に自己ホストできます。 Tarit - https://github.com/instavm/tarit クイックスタート - https://github.com/instavm/tarit/blob/main/orch/docs/QUICKST...

記事本文:
GitHub - instavm/tarit: セルフホスト型 AI エージェントと RL 用のハイパーバイザーおよびサンドボックス クラウド · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
インスタ
/
タリット
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット .github .github orch orch proto proto scripts scripts test test vmm vmm .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md 拒否.toml 拒否.toml 錆びたfmt.toml 錆びたfmt.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントおよび RL 環境向けの最速のハイパーバイザーおよびサンドボックス クラウド。
Tarit は、安全、高速、一時的なサービスを実現する microVM プラットフォームです
AI エージェントのワークロード用に構築されたサンドボックス。実際のハードウェア仮想化 VM を起動します
ミリ秒単位でタスクを実行し、それを破棄して各サンドボックスを提供します
共有カーネルコンテナ境界の代わりにカーネルレベルで分離します。
これには 2 つの部分があり、このモノリポジトリで一緒に開発されました。
vmm/ - Tarit VMM、最小限のrust-vmm
ベースの microVM モニター (ハイパーバイザー層)。 1 つのプロセスで 1 つの microVM が実行されます。
単独で、または任意のオーケストレーターの下で使用できます。
orch/ - taritd 、マルチノード オーケストレーターおよび PaaS コントロール プレーン
配置、ウォーム プール、
ネットワーキング、スナップショット、SSH/PTY アクセス、キーごとの使用状況統計、監査証跡。
彼らは、型が 1 つの共有内に存在する Unix ドメイン ソケット プロトコルを介して通信します。
dependency-light クレート、proto/ ( tarit-proto )。あの木箱はワイヤーです
契約しているため、taritd または独自のコントロール プレーンから VMM を駆動できます
手書きのタイプをコピーする必要はありません。
本当の孤立。各サンドボックスは、独自のカーネルを持つ KVM ゲストであり、
名前空間プロセス。侵害されたワークロードまたは暴走したワークロードは、ホストまたは
その隣人たち。
早くて安い。最小デバイス モデル (MMIO virtio のみ、PCI、BIOS なし)、
デマンド ページング ゲスト RAM、

d スナップショット/リストアは 1 秒以内に開始されます。
設計上一時的です。作成、実行、破棄。スナップショットとコピーオンライト
オーバーレイを使用すると、多数の同一のサンドボックスを安価に起動できます。
エージェント向けに構築されています。 vsock ベースの実行と対話型 PTY、キーごとの使用法
メータリングと監査、およびバースト的な作成/実行/破棄向けに調整されたオーケストレーター。
ベアメタル検証からの Tarit カラム。出版された「爆竹」コラム
ドキュメント。
HTTP API + CLI + SSH ゲートウェイ
|
taritd (orch) 1 つのマルチノード コントロール プレーン
/ | \ 配置、温水プール、フリート
/ | \ 使用状況 + 監査 -> PostgreSQL
vmm サーブ vmm サーブ vmm サーブ microVM ごとに 1 つのプロセス
| | |
microVM microVM microVM KVM ゲスト、独自のカーネル
taritd とサードパーティ オーケストレーターは同じ tarit-proto プロトコルを使用します
vmmserve へ: VM ごとに、長さのプレフィックスが付いた 1 つの JSON リクエスト、1 つの応答
Unixソケット。 Bring-your-own-orchestrator については、vmm/docs/INTEGRATION.md を参照してください。
クイックスタートは階層化されています。レイヤ 1 は、microVM でコードを実行します。レイヤー2
スナップショットを追加し、一時停止し、復元します。レイヤ 3 は、管理対象フリートを実行します。
オーケストレーター。必要なレイヤーだけを取り出します。
KVM ( /dev/kvm ) と Rust ツールチェーンを備えた Linux ホストが必要です。ランニング
microVM には root (または kvm グループのメンバーシップ) が必要なので、コマンドでは次を使用します。
須藤。
レイヤ 1: microVM でコードを実行する
git clone https://github.com/instavm/tarit && cd tarit
sudo make install # build + install vmm、taritd、およびゲスト エージェント
sudo make guest # one-time: ゲスト カーネルをビルドし、Ubuntu rootfs をプルします
make guest は、遅い作業 (カーネルのビルド + OCI プル) を 1 回実行して書き込みます。
guest-assets/vmlinux および guest-assets/rootfs.ext4 なので、その後 VM を起動します
インスタントです。 1 つを起動し、その中でコマンドを実行し、それを破棄します。
sudo vmmserve --socket /tmp/vm.sock &
sudo vmm --socket /tmp/vm.sock create --kernel guest-assets/vmlinux --rootfs guest-assets/r

ootfs.ext4
sleep 12 # ゲストを起動させてエージェントにダイヤルさせます
sudo vmm --socket /tmp/vm.sock exec " uname -a "
sudo vmm --socket /tmp/vm.sock stop
ハイパーバイザーだけが必要ですか? sudo make install-vmm は vmm のみをインストールします。
レイヤ 2: スナップショット、サスペンド、リストア
同じソケットを駆動して VM の状態をキャプチャおよび移動します。完全なスナップショットはメモリに書き込みます
プラスデバイスの状態。 --diff はダーティ ページのみを書き込みます。一時停止リリース常駐
ゲストRAM。履歴書がそれを取り戻します。復元では、スナップショットから新しい VMM を起動します。
sudo vmm --socket /tmp/vm.sock snapshot # 完全なスナップショット、.snap パスを出力します
sudo vmm --socket /tmp/vm.sock snapshot --diff # 増分 (ダーティ ページのみ)
sudo vmm --socket /tmp/vm.sock stop # 常駐ゲスト RAM を解放
sudo vmm --socket /tmp/vm.sock 再開
sudo vmm stop --snapshot /path/to.snap # 新しい VMM プロセスに復元します
Tarit はライブ スナップショットも実行します。つまり、実行中のメモリの一貫性のあるスナップショットです。
ダウンタイムなしでゲストを実行できるため、ビジー状態の VM のチェックポイントまたはフォークを実行できます。参照
vmm/docs/STANDALONE.md (デバイス全体、出力)、
看守とPTY表面。
レイヤー 3: オーケストレーターを使用してフリートを実行する
taritd は、HTTP API を介して 1 つ以上のノードにわたる多数の microVM を管理します。
配置、ウォーム プール、キーごとの使用状況のアカウンティング、および SSH/PTY ゲートウェイ。シングル
ノード:
CDオーケストラ
TARIT_API_KEY= $( openssl rand -hex 24 ) \
TARIT_VMM_BIN= $( コマンド -v vmm ) \
TARIT_KERNEL= $PWD /../guest-assets/vmlinux TARIT_ROOTFS= $PWD /../guest-assets/rootfs.ext4 \
タリトサーブ
次に、HTTP 経由で VM を作成して駆動し、マルチノード クラスターにスケールします。
Orchestrator のクイックスタートと構成: orch/docs/QUICKSTART.md 、
orch/docs/CONFIGURATION.md
アーキテクチャと HTTP API: orch/docs/ARCHITECTURE.md 、
orch/docs/API.md
マルチノード クラスタ、フェイルオーバー、自動スケーリング: orch/docs/RESILIENCE.md 、
orch/docs/AUTOSCALING.md
使用状況の統計と監査

レール: orch/docs/USAGE-AND-AUDIT.md
独自のコントロール プレーンから VMM を駆動します: vmm/docs/INTEGRATION.md
vmm/ Tarit VMM (microVM モニター) - 独自のカーゴ ワークスペース
orch/taritd、オーケストレーターおよび PaaS コントロール プレーン - 独自のカーゴ ワークスペース
proto/tarit-proto、共有 UDS ワイヤー プロトコル クレート (KVM フリー)
vmm/ と orch/ は独立したカーゴ ワークスペースであるため、VMM を構築できます。
テストされ、それ自体で消費されます。どちらもワイヤの種類は proto/ に依存します。
Tarit VMM (スタンドアロン): vmm/README.md 、
vmm/docs/STANDALONE.md
独自のオーケストレーターを導入します: vmm/docs/INTEGRATION.md
VMM ビルド + CLI + UDS API: vmm/docs/BUILD-AND-API.md
Orchestrator クイックスタート + 構成: orch/docs/QUICKSTART.md 、
orch/docs/CONFIGURATION.md
オーケストレーター アーキテクチャ + API: orch/docs/ARCHITECTURE.md 、
orch/docs/API.md
復元力とスケール (テスト済みのシナリオ): orch/docs/RESILIENCE.md
使用状況の統計と監査証跡: orch/docs/USAGE-AND-AUDIT.md
ホスト: KVM を備えた x86_64 Linux。開発はビルドとビルドのために macOS でも動作します。
クロスチェック。 microVM を実行するには KVM が必要です。
未実装: aarch64 ゲスト、virtio-balloon。
セルフホスティングは AWS と GCP でテストされています。 Azure のサポートは近日中に開始されます。
Tarit は AGPL-3.0 以降に基づいてライセンスされています。 「ライセンス」を参照してください。
セルフホスト型 AI エージェントおよび RL 用のハイパーバイザーおよびサンドボックス クラウド
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
1
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A hypervisor and sandbox cloud for self-hosted AI agents and RL - instavm/tarit

We have built Tarit as a hypervisor built from ground up for running AI agent and RL environments. It is based on rust-vmm and can be used as a replacement for firecracker. Firecracker was built to serve a different need of primarily serverless compute and hence does not have primitives like live snapshots without pausing the VM operations. We also provide a basic orchestrator that handles placement of the microVMs, creating clusters with HA, maintaining a warm pool of VMs, and takes care of setting up networking and monitoring. Our benchmarks on a metal instance shows an acquire VM from pool and execute code p99 of 35ms which places it as one of the fastest VM based sandboxes you can get right now. You could also snapshot a small sandbox and resume in ~80ms. You can easily self-host a multi-node cluster of this right now on your cloud with nested-virt enabled machines or better on bare metal by providing the docs to your coding agent. Tarit - https://github.com/instavm/tarit Quickstart - https://github.com/instavm/tarit/blob/main/orch/docs/QUICKST...

GitHub - instavm/tarit: A hypervisor and sandbox cloud for self-hosted AI agents and RL · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
instavm
/
tarit
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits .github .github orch orch proto proto scripts scripts test test vmm vmm .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md deny.toml deny.toml rustfmt.toml rustfmt.toml View all files Repository files navigation
The fastest hypervisor and sandbox cloud for AI agents and RL environments.
Tarit is a microVM platform for secure, fast, ephemeral
sandboxes, built for AI agent workloads. It boots a real hardware-virtualized VM
in milliseconds, runs a task inside it, and tears it down, giving each sandbox
kernel-level isolation instead of a shared-kernel container boundary.
It has two parts, developed together in this monorepo:
vmm/ - the Tarit VMM, a minimal rust-vmm
based microVM monitor (the hypervisor layer). One process runs one microVM.
Usable on its own or under any orchestrator.
orch/ - taritd , a multi-node orchestrator and PaaS control plane that
launches and manages microVMs across a fleet, with placement, warm pools,
networking, snapshots, SSH/PTY access, per-key usage stats, and an audit trail.
They talk over a Unix-domain-socket protocol whose types live in one shared,
dependency-light crate, proto/ ( tarit-proto ). That crate is the wire
contract, so you can drive the VMM from taritd or from your own control plane
without hand-copying types.
Real isolation. Each sandbox is a KVM guest with its own kernel, not a
namespaced process. A compromised or runaway workload cannot see the host or
its neighbors.
Fast and cheap. Minimal device model (MMIO virtio only, no PCI, no BIOS),
demand-paged guest RAM, and snapshot/restore for sub-second starts.
Ephemeral by design. Create, run, discard. Snapshots and copy-on-write
overlays make many identical sandboxes cheap to spin up.
Built for agents. vsock-based exec and interactive PTY, per-key usage
metering and audit, and an orchestrator tuned for bursty create/exec/destroy.
Tarit column from bare-metal validation; Firecracker column from its published
docs.
HTTP API + CLI + SSH gateway
|
taritd (orch) one multi-node control plane
/ | \ placement, warm pool, fleet
/ | \ usage + audit -> PostgreSQL
vmm serve vmm serve vmm serve one process per microVM
| | |
microVM microVM microVM KVM guest, own kernel
taritd and any third-party orchestrator speak the same tarit-proto protocol
to vmm serve : one length-prefixed JSON request, one response, over a per-VM
Unix socket. See vmm/docs/INTEGRATION.md for bring-your-own-orchestrator.
The quickstart is layered. Layer 1 gets your code running in a microVM. Layer 2
adds snapshots, suspend, and restore. Layer 3 runs a managed fleet with the
orchestrator. Take only the layer you need.
You need a Linux host with KVM ( /dev/kvm ) and a Rust toolchain. Running
microVMs needs root (or membership in the kvm group), so the commands use
sudo .
Layer 1: run code in a microVM
git clone https://github.com/instavm/tarit && cd tarit
sudo make install # build + install vmm, taritd, and the guest agent
sudo make guest # one-time: build a guest kernel + pull an Ubuntu rootfs
make guest does the slow work once (kernel build + OCI pull) and writes
guest-assets/vmlinux and guest-assets/rootfs.ext4 , so starting a VM afterwards
is instant. Boot one, run a command in it, tear it down:
sudo vmm serve --socket /tmp/vm.sock &
sudo vmm --socket /tmp/vm.sock create --kernel guest-assets/vmlinux --rootfs guest-assets/rootfs.ext4
sleep 12 # let the guest boot and dial the agent
sudo vmm --socket /tmp/vm.sock exec " uname -a "
sudo vmm --socket /tmp/vm.sock stop
Only want the hypervisor? sudo make install-vmm installs just vmm .
Layer 2: snapshot, suspend, restore
Drive the same socket to capture and move VM state. A full snapshot writes memory
plus device state; --diff writes only dirty pages. Suspend releases resident
guest RAM; resume brings it back. Restore boots a fresh VMM from a snapshot.
sudo vmm --socket /tmp/vm.sock snapshot # full snapshot, prints the .snap path
sudo vmm --socket /tmp/vm.sock snapshot --diff # incremental (dirty pages only)
sudo vmm --socket /tmp/vm.sock suspend # release resident guest RAM
sudo vmm --socket /tmp/vm.sock resume
sudo vmm restore --snapshot /path/to.snap # restore into a new VMM process
Tarit also does live snapshots : a memory-consistent snapshot of a running
guest with no downtime, so a busy VM can be checkpointed or forked. See
vmm/docs/STANDALONE.md for the full device, egress,
jailer, and PTY surface.
Layer 3: run a fleet with the orchestrator
taritd manages many microVMs across one or more nodes over an HTTP API, with
placement, warm pools, per-key usage accounting, and an SSH/PTY gateway. Single
node:
cd orch
TARIT_API_KEY= $( openssl rand -hex 24 ) \
TARIT_VMM_BIN= $( command -v vmm ) \
TARIT_KERNEL= $PWD /../guest-assets/vmlinux TARIT_ROOTFS= $PWD /../guest-assets/rootfs.ext4 \
taritd serve
Then create and drive VMs over HTTP, and scale to a multi-node cluster:
Orchestrator quickstart and config: orch/docs/QUICKSTART.md ,
orch/docs/CONFIGURATION.md
Architecture and HTTP API: orch/docs/ARCHITECTURE.md ,
orch/docs/API.md
Multi-node cluster, failover, autoscaling: orch/docs/RESILIENCE.md ,
orch/docs/AUTOSCALING.md
Usage stats and audit trail: orch/docs/USAGE-AND-AUDIT.md
Drive the VMM from your own control plane: vmm/docs/INTEGRATION.md
vmm/ the Tarit VMM (microVM monitor) - its own cargo workspace
orch/ taritd, the orchestrator and PaaS control plane - its own cargo workspace
proto/ tarit-proto, the shared UDS wire protocol crate (KVM-free)
vmm/ and orch/ are independent cargo workspaces so the VMM can be built,
tested, and consumed on its own. Both depend on proto/ for the wire types.
Tarit VMM (standalone): vmm/README.md ,
vmm/docs/STANDALONE.md
Bring your own orchestrator: vmm/docs/INTEGRATION.md
VMM build + CLI + UDS API: vmm/docs/BUILD-AND-API.md
Orchestrator quickstart + config: orch/docs/QUICKSTART.md ,
orch/docs/CONFIGURATION.md
Orchestrator architecture + API: orch/docs/ARCHITECTURE.md ,
orch/docs/API.md
Resilience and scale (tested scenarios): orch/docs/RESILIENCE.md
Usage stats and audit trail: orch/docs/USAGE-AND-AUDIT.md
Host: x86_64 Linux with KVM. Development also works on macOS for building and
cross-checking; running microVMs needs KVM.
Not yet implemented: aarch64 guests, virtio-balloon.
Self-hosting has been tested on AWS and GCP. Azure support is coming soon.
Tarit is licensed under AGPL-3.0-or-later . See LICENSE .
A hypervisor and sandbox cloud for self-hosted AI agents and RL
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
1
fork
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
