---
source: "https://github.com/upbound/asterkube"
hn_url: "https://news.ycombinator.com/item?id=48832561"
title: "Running Kubernetes without C in Rust-based Asterinas kernel"
article_title: "GitHub - upbound/asterkube: Secure Kubernetes without Linux or C: the Asterinas Rust framekernel + a CGO-free Go kubelet as PID 1 · GitHub"
author: "boeroboy"
captured_at: "2026-07-08T15:11:23Z"
capture_tool: "hn-digest"
hn_id: 48832561
score: 1
comments: 1
posted_at: "2026-07-08T14:35:58Z"
tags:
  - hacker-news
  - translated
---

# Running Kubernetes without C in Rust-based Asterinas kernel

- HN: [48832561](https://news.ycombinator.com/item?id=48832561)
- Source: [github.com](https://github.com/upbound/asterkube)
- Score: 1
- Comments: 1
- Posted: 2026-07-08T14:35:58Z

## Translation

タイトル: Rust ベースの Asterinas カーネルで C を使用せずに Kubernetes を実行する
記事のタイトル: GitHub - upbound/asterkube: Linux または C を使用しない安全な Kubernetes: Asterinas Rust フレームカーネル + PID 1 としての CGO フリー Go kubelet · GitHub
説明: Linux または C を使用しない安全な Kubernetes: Asterinas Rust フレームカーネル + CGO フリーの Go kubelet (PID 1 として) - upbound/asterkube

記事本文:
GitHub - アップバウンド/asterkube: Linux または C を使用しない安全な Kubernetes: Asterinas Rust フレームカーネル + PID 1 としての CGO フリー Go kubelet · GitHub
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
上り
/
アステルクベ
公共
通知

イオン
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
63 コミット 63 コミット アセット アセット asterinas @ 1d07dde asterinas @ 1d07dde build build cmd/ asterkube-init cmd/ asterkube-init config config docs docs scripts scripts .gitignore .gitignore .gitmodules .gitmodules ライセンス ライセンス通知 通知 クイックスタート.md QUICKSTART.md README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Linux や C を使用せずに Kubernetes を保護します。
Asterinas 上に構築されたブート可能な Kubernetes ノード
CGOフリーのGoユーザースペースを備えたメモリセーフなRustフレームカーネル - 本当のアップストリーム
kubelet v1.35.6 は PID 1 として実行されます。Linux カーネルはありません。システムドはありません。リブはありません。
CoreOS は、Kubernetes を安全に実行するための最小限の不変 Linux VM の先駆者でしたが、
Linux カーネル、systemd、および glibc 上で動作しており、これらの C ベースのレイヤーには
それ以来、数千ものメモリ安全でない CVE が出荷されました (Linux カーネルだけでも
約 11,000 ～ 12,000 の CVE があり、メモリ安全性のバグが約半分です)。これらは、
毎年、ランサムウェアやデータ侵害で数十億ドルの被害が発生しています。まさにこれがそうだったように
Linux の futex コードで 15 年前に使用された use-after-free が作成中
( GhostLock、CVE-2026-43499 ) が表面化 — ローカルの
ユーザー root を設定し、ほとんどのディストリビューションでコンテナー エスケープを有効にします。
VM にはハードウェア カーネルは必要ありません。ハイパーバイザーはすでに難しい仕事を行っています -
デバイスドライバー、ECC、メモリ管理、エラー検出 - 簡素化されたプレゼンテーション、
ゲストへの仮想化インターフェイス。これにより、小型でメモリセーフなカーネルの余地が残されます。
クラウド ワークロードを実行するために完全に Rust で書かれています。アステリナス
これはまさに、Rust の Linux-ABI 互換フレームカーネル (MPL-2.0) です。
asterkube は Asterinas を実現する概念実証です。

実際の Kubernetes ノードへ —
そしてイメージ全体から C を削除します (唯一の例外は GRUB2 ブートローダーです。
次のターゲット）。これは、Upbound Labs の投稿の背後にある完全なストーリーです。
Linux や C を使用しない安全な Kubernetes 。
完全な開示: これは実験的な概念実証であり、大部分が書かれており、
クロードでテスト (Fable 5 と Opus 4.6 ～ 4.8)。本番環境に対応していないため、
Asterinas のメンテナは、ブランチに到達するこれらについて最終決定権を持っています。
レイヤー
伝統的な
アステルクベ
コンテナランタイム
コンテナード (Go)
コンテナード (Go)
オーケストレーション
kubelet (Go)
拡張 kubelet (Go) — オーケストレーション + init
初期化
systemd (C)
(kubelet に折りたたまれます)
カーネル
Linux (C)
アステリナス(錆び)
kubelet は PID 1 として機能するように拡張され、ファイルシステムのマウント、ネットワーキングの起動を行います。
(DHCP ファースト)、ポッドを実行しているため、systemd や個別の init はありません。の
Kubernetes ノード サービスはオペレーティング システムです。
asterkube (このリポジトリ、Apache-2.0)
§── cmd/asterkube-init/ ← Go ノード エージェント (CGO フリー、kubelet 融合 init)
§── scripts/ ← ビルド + パッケージ化パイプライン
§── config/ ← デフォルトの /etc/fstab と /etc/os-release がイメージに焼き付けられます
§── docs/ ← FEATURES.md (完全な項目別機能セット)、SECURITY-POSTURE.md、…
§── asterinas/ (サブモジュール) → jboero/asterinas @ asterkube (カーネル、MPL-2.0)
└── kubernetes/ (v1.35.6 でビルドフェッチ、gitignored) (kubelet ソース、Apache-2.0)
asterinas はカーネル フォークに固定された git サブモジュールです — カーネルの最大 50 コミット
作業 (名前空間、seccomp-BPF、astromac MAC、サービス NAT データパス、
nftables 互換のネットリンク サーフェス) がそこに存在します。
kubernetes はベンダー化されていません。ビルドはピン留めされたタグ v1.35.6 を取得します。
実際の cmd/kubelet/app をコンパイルするには完全なツリーが必要であるため、要求されます。
cmd/asterkube-init/

Kubernetes をインポートするオリジナルの Go です。それはフォークではありません。
ノード全体 — メモリセーフな Rust カーネル + CGO フリーの Go ユーザー空間、ゼロ C (なし
glibc/musl、/lib64 なし) — で起動する自己完結型の ~40 MB ISO または QCOW2
5秒未満。ユーザー空間をいくつかの静的バイナリに統合し、
zstd で圧縮すると、138 MiB のコンテンツが 26 MiB の initramfs に圧縮されます。
§─ ブート アーティファクト (zstd 圧縮された initramfs)
│ §─ カーネル ELF (リリース、ストリップ) ...... 5.3 MiB
│ §─ initramfs.cpio.zst (zstd --ultra -22) 。 26.4 MiB ← 44 MiB gzip (−40%)
│ §─ ISO (ブート可能、isohybrid) ................................ 40.6 MiB ← 58 MiB gzip (−31%)
│ ━─ QCOW2 (ディスクイメージ) ................................ 38.2 MiB ← 56 MiB gzip (−32%)
│
━─ initramfs コンテンツ (非圧縮 138 MiB → 26.4 MiB zstd)
§─ sbin/init → ../usr/bin/kubelet ......... シンボリックリンク
§─ usr/bin/
│ §─ kubelet ................................... 79.1 MiB 実アップストリーム kubelet + init(PID 1)
│ │ + pure-Go runc + マウントアプレット + DHCP
│ §─containerd .................................. 42.7 MiB static、containerd+ctr をマージ
│ §─containerd-shim-runc-v2 ................ 13.9 MiB static
│ └─ ctr →containerd ................... シンボリックリンク
§─ usr/lib64/ ................................... 空 (証拠: C ランタイムがない)
§─ usr/share/asterkube/hello.tar ...... 1.5 MiB デモ イメージ
§─ etc/os-release .................................. Asterinas として識別されます
└─ etc/fstab ................................... デフォルトのマウントを文書化
3 つの静的で CGO フリーの Go バイナリがユーザー空間全体を構成します — kubelet (これは
同時に初期化、OCI ランタイム、マウント アプレット、DHCP クライアント)、マージされた
containerd + ctr 、および runc shim。空の usr/lib64/ はゼロ C の証明です。
カーネルギャップ

Asterinas と動作中の Kubernetes ノードの間にある と Go
それを駆動するノード エージェント — docs/FEATURES.md に項目化されています。
ハイライト:
名前空間と cgroup — PID/ネットワーク/ユーザー名前空間 (ルートレス、uid/gid マッピング、
privesc-safe 機能境界); cgroup v2 cpu/メモリ/pids の強制。
Real seccomp-BPF — システムコール ゲートでフィルターを強制するクラシック BPF インタープリター
(寛容なスタブでした); PR_GET/SET_SECCOMP 。
astromac — フレームカーネルネイティブの必須アクセス制御モジュール (ネイティブ
機能 - MAC、SELinux ポートではありません): クロステナント信号、ファイル、およびネットワーク
偽造不可能なプロセスごとのテナント ラベルによる分離。
ネットワーク データパス - L3 ブリッジ、負荷分散されたサービス DNAT (ClusterIP)
バックエンド、マスカレード/SNAT、nftables 互換の NETLINK_NETFILTER サーフェス
kube-proxy のルール、ネットリンク ルート プログラミング、ICMP、ワイルドカード UDP バインド (DHCP 用) を取り込みます。
OCIイネーブラ - cgroupデバイス・フィルタの bpf() 、 pivot_root 、 overlayfs 、
mqueue、および kubelet ContainerManager に必要な procfs/sysfs サーフェス。
zstd initramfs — 純粋な Rust ( ruzstd 、 no_std ) zstd デコーダーで、画像の圧縮率が約 40% 小さくなります。
Go ノード エージェント ( cmd/asterkube-init/ ):
実際のアップストリーム kubelet は、init と 1 つのマルチコール バイナリ (init、runc、
マウント、DHCP は argv[0] によってディスパッチされます)。
リース更新を伴う DHCP ファーストのユーザー空間ネットワーキング (RFC 2131)。
Zero-C コンテナー ランタイム — 静的コンテナー + 純粋な Go、CGO フリーの runc の代替。
Boot-args クラスター結合 — カーネル コマンドライン、apiserver からの kubeadm スタイルのブートストラップ
IP + CA ハッシュによって固定されます。
あらゆるカーネルのセキュリティ保証を実行して失敗する敵対的なブート時プローブ
ブーツが保持できない場合。
ノードは、( /etc/os-release → OS Image を介して) Asterinas として正確に識別され、
カーネル バージョン 5.13.0-asterinas

)、カスタム kernel.asterinas.io/name=asterinas をアドバタイズします。
ノードラベル。重要なのは、kubernetes.io/os=linux を維持することです — Asterinas は Linux を実行します
ABI、つまり「linux」には機能互換性宣言があります（gVisor や FreeBSD のような）
linuxulator)、ブランドの主張ではありません。このよく知られたラベルを変更すると、ポッドのスケジューリングが壊れる可能性があります
エコシステム全体にわたる OCI イメージのマッチング。 Linux ではなく Linux 互換です。
リリースイメージを取得して起動します —
必要なのは qemu-system-x86_64 、UEFI ファームウェア ( edk2-ovmf / ovmf )、および KVM のみです。
gh リリースのダウンロード asterkube-v0.1 --repo upbound/asterkube && sha256sum -c SHA256SUMS
カール -sO https://raw.githubusercontent.com/upbound/asterkube/main/scripts/run-release.sh
chmod +x run-release.sh && ./run-release.sh asterkube-node-v0.1.iso # または .qcow2
数秒で起動し、シャットダウンするまでライブ状態を維持します — Ctrl-A C を押してから
QEMU モニターの system_powerdown。完全なチュートリアルは QUICKSTART.md にあります。
前提条件: Docker、Go、qemu-system-x86_64 + OVMF、zstd、および initramfs ツール
( cpio 、 readelf )。ビルド コンテナには Rust + OSDK ツールチェーンが含まれています。
git clone --recursive https://github.com/upbound/asterkube && cd asterkube
# Asterinas OSDK ビルドコンテナを作成します (名前は「asterkube」である必要があります) + Cargo-osdk をインストールします
docker run -d --name asterkube --privileged --network=host -v /dev:/dev \
-v " $PWD /asterinas:/root/asterinas " -w /root/asterinas \
アステリーナ/アステリーナ:0.18.0-20260603 スリープ無限大
docker exec asterkube bash -lc ' OSDK_LOCAL_DEV=1 カーゴインストール カーゴ-osdk --path osdk '
scripts/build-containerd-merged.sh # staticcontainerd+ctr → build/static-bins
scripts/build-hello-image.sh # デモ画像 (buildah/skopeo が必要)
scripts/build-zeroc-kubelet.sh # 結合された CGO フリー kubelet (k8s v1.35.6 をフェッチ)
INIT_BIN=build/asterkube-kubelet s

cripts/zero-c-initramfs.sh # zero-C initramfs + ブート可能 ISO
scripts/build-qcow2.sh # (オプション) ISO を QCOW2 ディスクに変換します
scripts/run-release.sh asterinas/target/osdk/aster-kernel-osdk-bin.iso # ビルドしたものをブートします
クラスターに参加するには、ブート引数を生成します (これにより、有効期間の短いブートストラップ トークンがクラスターに書き込まれます)
クラスター) をカーネル コマンドラインに追加し、再構築します。
scripts/asterkube-boot-args.sh -n my-node # ASTERKUBE_* 行を印刷します
# これらを asterinas/OSDK.toml [run.boot] kcmd_args に貼り付けてから、zero-c-initramfs.sh を再実行します
次に、CNI が追加されるまで、ノードは kubectl get Nodes — NotReady に登録されます (以下を参照)。
これは概念実証であり、製品ではありません。
本番環境には対応していません。主に AI によって作成されています。 Asterina のメンテナは上流のあらゆるものをゲートします。
astromac は Permissive (ログのみ) を出荷します。装備されていますが、Enforcing に設定されるまでブロックされません。
NAT は最小限のデータパスです。小規模なグローバル接続トラックであり、エンドポイントの削除はありません。完全なサービス セマンティクスではありません。
自己完結型イメージに CNI がありません。結合されたノードは、追加されるまで NotReady のままです。
カーネル ギャップ - virtio デバイスのみ (他の NIC/ドライバー クラス、GPU なし)、ジャーナルなし
ファイルシステム (ext4/btrfs) または aarch64 はまだです。 virtio を利用した VM ワークロードには適しています。ベアメタルではありません。
GRUB2 はブート パス内の 1 つの C コンポーネントのままです。
asterkube は Apache-2.0 です (ライセンスと通知を参照)。参照します
2 つのアップストリームはベンダーではなく固定リビジョンによって提供されるため、それぞれが独自のライセンスを保持します。
MPL-2.0はfiです

[切り捨てられた]

## Original Extract

Secure Kubernetes without Linux or C: the Asterinas Rust framekernel + a CGO-free Go kubelet as PID 1 - upbound/asterkube

GitHub - upbound/asterkube: Secure Kubernetes without Linux or C: the Asterinas Rust framekernel + a CGO-free Go kubelet as PID 1 · GitHub
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
upbound
/
asterkube
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
63 Commits 63 Commits assets assets asterinas @ 1d07dde asterinas @ 1d07dde build build cmd/ asterkube-init cmd/ asterkube-init config config docs docs scripts scripts .gitignore .gitignore .gitmodules .gitmodules LICENSE LICENSE NOTICE NOTICE QUICKSTART.md QUICKSTART.md README.md README.md View all files Repository files navigation
Secure Kubernetes without Linux or C.
A bootable Kubernetes node built on the Asterinas
memory-safe Rust framekernel, with a CGO-free Go userspace — the real upstream
kubelet v1.35.6 running as PID 1. No Linux kernel. No systemd. No libc.
CoreOS pioneered the minimal, immutable Linux VM for running Kubernetes securely — but
it still rode on the Linux kernel, systemd, and glibc, and those C-based layers have
shipped thousands of memory-unsafe CVEs in the years since (the Linux kernel alone
has ~11,000–12,000 CVEs, roughly half memory-safety bugs). These get exploited in the
wild and cost billions in ransomware and data compromise every year. Just as this was
being written, a 15-year-old use-after-free in Linux's futex code
( GhostLock, CVE-2026-43499 ) surfaced — giving any local
user root and enabling container escape on most distros.
A VM doesn't need a hardware kernel. The hypervisor already does the hard job —
device drivers, ECC, memory management, error detection — and presents simplified,
virtualized interfaces to the guest. That leaves room for a small, memory-safe kernel
written entirely in Rust to run cloud workloads. Asterinas
is exactly that: a Linux-ABI-compatible framekernel in Rust (MPL-2.0).
asterkube is a proof of concept that turns Asterinas into a real Kubernetes node —
and strips C out of the entire image (the sole exception is the GRUB2 bootloader, the
next target). It is the full story behind the Upbound Labs post:
Secure Kubernetes Without Linux or C .
Full disclosure: this is an experimental proof of concept, largely written and
tested with Claude (Fable 5 and Opus 4.6–4.8). It is not production-ready, and the
Asterinas maintainers have final say on any of this reaching their branches.
Layer
Traditional
asterkube
Container runtime
containerd (Go)
containerd (Go)
Orchestration
kubelet (Go)
Extended kubelet (Go) — orchestration + init
Init
systemd ( C )
(folded into the kubelet)
Kernel
Linux ( C )
Asterinas (Rust)
The kubelet is extended to act as PID 1 — mounting filesystems, bringing up networking
(DHCP-first), and running pods — so there is no systemd and no separate init. The
Kubernetes node service is the operating system.
asterkube (this repo, Apache-2.0)
├── cmd/asterkube-init/ ← our Go node agent (the CGO-free, kubelet-fused init)
├── scripts/ ← build + packaging pipeline
├── config/ ← default /etc/fstab and /etc/os-release baked into the image
├── docs/ ← FEATURES.md (full itemized featureset), SECURITY-POSTURE.md, …
├── asterinas/ (submodule) → jboero/asterinas @ asterkube (kernel, MPL-2.0)
└── kubernetes/ (build-fetched @ v1.35.6, gitignored) (kubelet source, Apache-2.0)
asterinas is a git submodule pinned to our kernel fork — ~50 commits of kernel
work (namespaces, seccomp-BPF, the astromac MAC, a Service NAT datapath, an
nftables-compatible netlink surface) live there.
kubernetes is not vendored — the build fetches the pinned tag v1.35.6 on
demand, because compiling the real cmd/kubelet/app needs the full tree.
cmd/asterkube-init/ is original Go that imports Kubernetes; it is not a fork of it.
The whole node — memory-safe Rust kernel + CGO-free Go userspace, zero C (no
glibc/musl, no /lib64 ) — is a self-contained ~40 MB ISO or QCOW2 that boots in
under 5 seconds . Consolidating the userspace into a handful of static binaries and
compressing with zstd takes the 138 MiB of contents down to a 26 MiB initramfs.
├─ boot artifacts (zstd-compressed initramfs)
│ ├─ kernel ELF (release, stripped) ........ 5.3 MiB
│ ├─ initramfs.cpio.zst (zstd --ultra -22) . 26.4 MiB ← 44 MiB gzip (−40%)
│ ├─ ISO (bootable, isohybrid) ............. 40.6 MiB ← 58 MiB gzip (−31%)
│ └─ QCOW2 (disk image) ..................... 38.2 MiB ← 56 MiB gzip (−32%)
│
└─ initramfs contents (138 MiB uncompressed → 26.4 MiB zstd)
├─ sbin/init → ../usr/bin/kubelet ......... symlink
├─ usr/bin/
│ ├─ kubelet ............................. 79.1 MiB real upstream kubelet + init(PID 1)
│ │ + pure-Go runc + mount applet + DHCP
│ ├─ containerd ......................... 42.7 MiB static, merged containerd+ctr
│ ├─ containerd-shim-runc-v2 ............ 13.9 MiB static
│ └─ ctr → containerd ................... symlink
├─ usr/lib64/ ............................ empty (proof: no C runtime)
├─ usr/share/asterkube/hello.tar ......... 1.5 MiB demo image
├─ etc/os-release ........................ identifies as Asterinas
└─ etc/fstab ............................. documented default mounts
Three static, CGO-free Go binaries make up the entire userspace — kubelet (which is
simultaneously init, the OCI runtime, the mount applet and the DHCP client), the merged
containerd + ctr , and the runc shim. The empty usr/lib64/ is the zero-C proof.
The kernel gaps that stood between Asterinas and a working Kubernetes node — and the Go
node agent that drives it — are itemized in docs/FEATURES.md .
The highlights:
Namespaces & cgroups — PID/network/user namespaces (rootless, uid/gid mapping,
privesc-safe capability boundary); cgroup v2 cpu/memory/pids enforcement.
Real seccomp-BPF — a classic-BPF interpreter enforcing filters at the syscall gate
(was a permissive stub); PR_GET/SET_SECCOMP .
astromac — a framekernel-native Mandatory Access Control module (native
capability-MAC, not a SELinux port): cross-tenant signal , file , and network
isolation with unforgeable per-process tenant labels.
Networking datapath — L3 bridge, Service DNAT (ClusterIP) with load-balanced
backends, masquerade/SNAT, an nftables-compatible NETLINK_NETFILTER surface that
ingests kube-proxy's rules, netlink route programming, ICMP, wildcard UDP bind (for DHCP).
OCI enablers — bpf() for the cgroup device filter, pivot_root , overlayfs,
mqueue, and the procfs/sysfs surface the kubelet ContainerManager needs.
zstd initramfs — a pure-Rust ( ruzstd , no_std) zstd decoder so the image compresses ~40% smaller.
Go node agent ( cmd/asterkube-init/ ):
The real upstream kubelet fused with init as one multi-call binary (init, runc,
mount, DHCP dispatched by argv[0] ).
DHCP-first userspace networking (RFC 2131) with lease renewal.
Zero-C container runtime — static containerd + a pure-Go, CGO-free runc replacement.
Boot-args cluster join — kubeadm-style bootstrap from the kernel cmdline, apiserver
pinned by IP + CA-hash.
Adversarial boot-time probes that exercise every kernel security guarantee and fail
the boot if it doesn't hold.
The node truthfully identifies as Asterinas (via /etc/os-release → OS Image , and a
kernel version of 5.13.0-asterinas ), and advertises a custom kernel.asterinas.io/name=asterinas
node label. Crucially it keeps kubernetes.io/os=linux — Asterinas runs the Linux
ABI, so "linux" there is a functional compatibility declaration (like gVisor or FreeBSD's
linuxulator), not a brand claim. Changing that well-known label would break pod scheduling
and OCI image matching across the ecosystem. It's Linux-compatible , not Linux.
Grab a release image and boot it —
you only need qemu-system-x86_64 , UEFI firmware ( edk2-ovmf / ovmf ), and KVM:
gh release download asterkube-v0.1 --repo upbound/asterkube && sha256sum -c SHA256SUMS
curl -sO https://raw.githubusercontent.com/upbound/asterkube/main/scripts/run-release.sh
chmod +x run-release.sh && ./run-release.sh asterkube-node-v0.1.iso # or the .qcow2
It boots in a few seconds and stays live until you shut it down — Ctrl-a c then
system_powerdown in the QEMU monitor. Full walkthrough in QUICKSTART.md .
Prerequisites: Docker, Go, qemu-system-x86_64 + OVMF, zstd , and initramfs tooling
( cpio , readelf ). The build container carries the Rust + OSDK toolchain.
git clone --recursive https://github.com/upbound/asterkube && cd asterkube
# create the Asterinas OSDK build container (name MUST be 'asterkube') + install cargo-osdk
docker run -d --name asterkube --privileged --network=host -v /dev:/dev \
-v " $PWD /asterinas:/root/asterinas " -w /root/asterinas \
asterinas/asterinas:0.18.0-20260603 sleep infinity
docker exec asterkube bash -lc ' OSDK_LOCAL_DEV=1 cargo install cargo-osdk --path osdk '
scripts/build-containerd-merged.sh # static containerd+ctr → build/static-bins
scripts/build-hello-image.sh # demo image (needs buildah/skopeo)
scripts/build-zeroc-kubelet.sh # combined CGO-free kubelet (fetches k8s v1.35.6)
INIT_BIN=build/asterkube-kubelet scripts/zero-c-initramfs.sh # zero-C initramfs + bootable ISO
scripts/build-qcow2.sh # (optional) convert the ISO to a QCOW2 disk
scripts/run-release.sh asterinas/target/osdk/aster-kernel-osdk-bin.iso # boot what you built
To join a cluster , generate boot args (this writes a short-lived bootstrap token to
your cluster), add them to the kernel cmdline, and rebuild:
scripts/asterkube-boot-args.sh -n my-node # prints ASTERKUBE_* lines
# paste them into asterinas/OSDK.toml [run.boot] kcmd_args, then re-run zero-c-initramfs.sh
The node then registers in kubectl get nodes — NotReady until a CNI is added (see below).
This is a proof of concept, not a product:
Not production-ready — largely AI-authored; Asterinas maintainers gate anything upstream.
astromac ships Permissive (log-only) — armed but not blocking until set to Enforcing.
NAT is a minimal datapath — small global conntrack, no endpoint removal; not full Service semantics.
No CNI in the self-contained image — a joined node stays NotReady until one is added.
Kernel gaps — virtio devices only (no other NIC/driver classes, no GPU), no journaled
filesystem (ext4/btrfs) or aarch64 yet. Fine for virtio-backed VM workloads; not bare metal.
GRUB2 remains the one C component in the boot path.
asterkube is Apache-2.0 (see LICENSE and NOTICE ). It references
its two upstreams by pinned revision rather than vendoring them, so each keeps its own license:
MPL-2.0 is fi

[truncated]
