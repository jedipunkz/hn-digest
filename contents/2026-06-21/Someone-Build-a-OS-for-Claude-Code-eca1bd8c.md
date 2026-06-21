---
source: "https://github.com/mateolafalce/claudios/blob/main/README.md"
hn_url: "https://news.ycombinator.com/item?id=48621576"
title: "Someone Build a OS for Claude Code"
article_title: "claudios/README.md at main · mateolafalce/claudios · GitHub"
author: "lafalce"
captured_at: "2026-06-21T19:31:09Z"
capture_tool: "hn-digest"
hn_id: 48621576
score: 1
comments: 1
posted_at: "2026-06-21T18:58:37Z"
tags:
  - hacker-news
  - translated
---

# Someone Build a OS for Claude Code

- HN: [48621576](https://news.ycombinator.com/item?id=48621576)
- Source: [github.com](https://github.com/mateolafalce/claudios/blob/main/README.md)
- Score: 1
- Comments: 1
- Posted: 2026-06-21T18:58:37Z

## Translation

タイトル: 誰かクロードコード用のOSを構築してください
記事のタイトル: claudios/README.md at main · mateolafalce/claudios · GitHub
説明: Ubuntu 24.04 LTS をベースにした最小限の OS。その唯一の目的は、プライマリ インターフェイスとして Claude Code を直接起動することです。 TUIのみ。 - メインのclaudios/README.md · mateolafalce/claudios

記事本文:
メインのclaudios/README.md · mateolafalce/claudios · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
マテオラファルセ
/
クラウディオス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
パスをコピーする もっと責める ファイル アクション もっと責める

ファイルアクション 最新のコミット
履歴 履歴 207 行 (150 loc) · 6.9 KB メイン ブレッドクラム
コピー パス トップ ファイルのメタデータとコントロール
raw ファイルのコピー raw ファイルのダウンロード アウトライン編集と raw アクション
ClaudiOS
OS 全体が単なるクロード コードだったらどうなるでしょうか?
Claude Code を直接起動する Linux ディストリビューション。デスクトップはありません。ウィンドウマネージャーはありません。 TTY と完全なシステム アクセス権を持つ AI エージェントだけです。 Ubuntu 24.04 LTS 上に構築されています。
クイックスタート · 仕組み · 機能 · ソースからのビルド
Claude Code がパッケージのインストール、コードの作成と実行、ファイルの管理、インターネットの閲覧、マシン全体の操作ができるとしたら、何のためにデスクトップが必要なのでしょうか?
ClaudiOS は、このアイデアを論理的な結論に導きます。つまり、オペレーティング システムが邪魔をせず、Claude がインターフェイスになります。起動してログインすると、クロードと話しています。それでおしまい。
これを USB スティックにフラッシュし、x86_64 マシンを起動すれば、ポータブル AI ワークステーションが完成します。
sudo apt install live-buildcurl gnupg grub-pc-bin xorriso
sudo ./build.sh
2. USB へのフラッシュ
sudo ./flash.sh /dev/sdX
/dev/sdX を USB デバイスに置き換えます (lsblk で見つけます)。スクリプトは ISO を書き込み、永続パーティションを自動的に作成します。USB は最初の起動から使用できるようになります。
sudo umount /dev/sdX *
sudo dd if=live-image-amd64.hybrid.iso of=/dev/sdX bs=4M status=progress
手動でフラッシュすると、永続パーティションは最初の起動時に自動的に作成されます。
claudios / claudios でログインすると、Claude Code に入ります。
BIOS/UEFI → GRUB → Linux → TTY ログイン
━─ ユーザー：claudios / パスワード：claudios
└─ claudios-shell (ログインシェル)
§─ [クロードが見つからない場合] npm 経由で自動インストール
└─ ループ内のクロードコード
└─ Ctrl+C (3秒) → エスケープしてバッシュ
claudios-shell はユーザーの実際のログイン シェルであり、 .bashrc ハックではなく、 /etc/passwd に設定されます。それ

クロード コードをループで起動します。 Claude が終了すると、再起動するか bash にドロップするように求められます。
ゼロ構成ブート — ログインすると Claude Code が表示され、すぐに作業できるようになります
USB 永続性 - 最初の起動時に永続性パーティションが自動作成されます。認証、設定、セッション履歴は再起動後も残ります
エスケープハッチ — 起動時に Ctrl+C を 3 秒間押し続けると、標準の bash シェルに戻ります。
自動インストール — クロードが PATH に見つからない場合、npm 経由で自動的にインストールされます。
組み込みのスラッシュ コマンド — /reboot はシステムを再起動し、/shut-down は終了してシェルに戻ります
パスワードなしの sudo — クロードが必要とするシステム操作用 (再起動、パッケージのインストール)
~1 GB ISO — ほとんどのハードウェアで 1 分以内に起動します
USB をフラッシュする必要はありません。
sudo apt install qemu-system-x86
./test.sh
すでに USB をフラッシュしましたか? QEMU でもブートできます。
sudo qemu-system-x86_64 \
-ドライブ ファイル=/dev/sdX,フォーマット=raw \
-m 4G \
-enable-kvm \
-CPUホスト\
-device virtio-vga \
-gtkの表示
/dev/sdX を USB デバイスに置き換えます (lsblk で見つけます)。 ISO は BIOS/レガシー モードで起動します。
ヒント: PC 端末からクロード アカウントにログインするには (認証 URL を簡単にコピーして貼り付けることができます)、代わりに -nographic を使用して起動します。
sudo qemu-system-x86_64 \
-ドライブ ファイル=/dev/sda、フォーマット=raw \
-m 4G \
-enable-kvm \
-CPUホスト\
-ノグラフィック\
-シリアル mon:stdio
カール 、 gnupg 、 grub-pc-bin 、 xorriso
auto/config — フラグを lb config に渡します (Ubuntu Noble、amd64、GRUB ブートローダー)
config/package-lists/claudios.list.chroot — イメージにインストールされているパッケージ
config/hooks/live/*.hook.chroot — chroot スクリプト (ロケール、Node.js、クロード コード、ユーザー作成)
config/includes.chroot/ — ファイルシステムにそのままコピーされたファイル
sudo ./clean.sh
パッケージ キャッシュも削除するには (次のビルドで完全な再ダウンロードを強制します):
sudo ./clean.sh && sudo rm -rf キャッシュ/
プロジェクト

電気ショック療法の構造
クラウディオス/
§── build.sh # メインのビルド スクリプト (sudo が必要)
§── flash.sh # ISO を永続的に USB にフラッシュします
§── clean.sh # ビルドアーティファクトを削除します
§── test.sh # QEMU で ISO を起動します
§── オート/
│ §── config # ライブビルド構成
│ §── build # lb ビルド用ラッパー
│ └── clean # lb clean 用ラッパー
━━ config/
§── archives/ # 追加の apt リポジトリ (NodeSource)
§── package-lists/ # ISO にインストールするパッケージ
§──hooks/live/ # ビルド内で実行されるスクリプト chroot
│ §── 0050-locale-timezone.hook.chroot
│ §── 0100-install-nodejs.hook.chroot
│ §── 0200-install-claude-code.hook.chroot
│ §── 0300-create-user.hook.chroot
│ └─ 0350-enable-persist.hook.chroot
└── include.chroot/ # ファイルシステムに直接コピーされたファイル
§── etc/motd
§── etc/シェル
§── etc/sudoers.d/claudios
§── etc/systemd/system/
│ └── claudios-persist.service
§── home/claudios/.claude/commands/
│ §── reboot.md # /reboot スラッシュコマンド
│ └── shut-down.md # /shut-down スラッシュコマンド
└── usr/local/bin/
§── claudios-shell # プライマリログインシェル
└── claudios-persist # 自動永続化設定スクリプト
緊急復旧
claudios-shell が失敗し、アクセスできなくなった場合:
GRUB から: カーネル行を編集し、 init=/bin/bash を追加します。
TTY から: claudios ユーザーは sudo を持っています — chsh を使用してシェルを変更します
フィールド
値
ユーザー名
クラウディオス
パスワード
クラウディオス
APIキー
クロード・コードが管理
ライセンス
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A minimalist OS based on Ubuntu 24.04 LTS whose sole purpose is to boot directly into Claude Code as the primary interface. TUI only. - claudios/README.md at main · mateolafalce/claudios

claudios/README.md at main · mateolafalce/claudios · GitHub
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
mateolafalce
/
claudios
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Copy path Blame More file actions Blame More file actions Latest commit
History History 207 lines (150 loc) · 6.9 KB main Breadcrumbs
Copy path Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions
ClaudiOS
What if your entire OS was just Claude Code?
A Linux distro that boots straight into Claude Code . No desktop. No window manager. Just a TTY and an AI agent with full system access. Built on Ubuntu 24.04 LTS.
Quick Start · How It Works · Features · Build From Source
If Claude Code can install packages, write and run code, manage files, browse the internet, and operate your entire machine — what do you need a desktop for?
ClaudiOS takes this idea to its logical conclusion: the operating system gets out of the way and Claude becomes the interface. You boot, you log in, and you're talking to Claude. That's it.
Flash it to a USB stick, boot any x86_64 machine, and you have a portable AI workstation.
sudo apt install live-build curl gnupg grub-pc-bin xorriso
sudo ./build.sh
2. Flash to USB
sudo ./flash.sh /dev/sdX
Replace /dev/sdX with your USB device ( lsblk to find it). The script writes the ISO and creates the persistence partition automatically — the USB is ready to use from the first boot.
sudo umount /dev/sdX *
sudo dd if=live-image-amd64.hybrid.iso of=/dev/sdX bs=4M status=progress
If you flash manually, the persistence partition will be created automatically on first boot instead.
Log in with claudios / claudios and you're in Claude Code.
BIOS/UEFI → GRUB → Linux → TTY login
└─ user: claudios / password: claudios
└─ claudios-shell (login shell)
├─ [if claude missing] auto-install via npm
└─ Claude Code in loop
└─ Ctrl+C (3s) → escape to bash
claudios-shell is the user's actual login shell — set in /etc/passwd , not a .bashrc hack. It launches Claude Code in a loop. If Claude exits, you're prompted to restart or drop to bash.
Zero-config boot — log in and you're in Claude Code, ready to work
USB persistence — first boot auto-creates a persistence partition; your auth, config, and session history survive reboots
Escape hatch — hold Ctrl+C for 3 seconds at startup to drop to a standard bash shell
Auto-install — if claude isn't found in PATH, it's installed automatically via npm
Built-in slash commands — /reboot reboots the system, /shut-down exits back to the shell
Passwordless sudo — for system operations Claude needs (reboot, package install)
~1 GB ISO — boots in under a minute on most hardware
No need to flash a USB to try it:
sudo apt install qemu-system-x86
./test.sh
Already flashed a USB? You can boot it in QEMU too:
sudo qemu-system-x86_64 \
-drive file=/dev/sdX,format=raw \
-m 4G \
-enable-kvm \
-cpu host \
-device virtio-vga \
-display gtk
Replace /dev/sdX with your USB device ( lsblk to find it). The ISO boots in BIOS/Legacy mode.
Tip: To log in to your Claude account from your PC terminal (so you can copy and paste the auth URL easily), boot with -nographic instead:
sudo qemu-system-x86_64 \
-drive file=/dev/sda,format=raw \
-m 4G \
-enable-kvm \
-cpu host \
-nographic \
-serial mon:stdio
curl , gnupg , grub-pc-bin , xorriso
auto/config — passes flags to lb config (Ubuntu Noble, amd64, GRUB bootloader)
config/package-lists/claudios.list.chroot — packages installed in the image
config/hooks/live/*.hook.chroot — chroot scripts (locale, Node.js, Claude Code, user creation)
config/includes.chroot/ — files copied verbatim into the filesystem
sudo ./clean.sh
To also remove the package cache (forces a full re-download on next build):
sudo ./clean.sh && sudo rm -rf cache/
Project Structure
claudios/
├── build.sh # Main build script (requires sudo)
├── flash.sh # Flashes ISO to USB with persistence
├── clean.sh # Removes build artifacts
├── test.sh # Launches the ISO in QEMU
├── auto/
│ ├── config # live-build configuration
│ ├── build # Wrapper for lb build
│ └── clean # Wrapper for lb clean
└── config/
├── archives/ # Additional apt repositories (NodeSource)
├── package-lists/ # Packages to install in the ISO
├── hooks/live/ # Scripts that run inside the build chroot
│ ├── 0050-locale-timezone.hook.chroot
│ ├── 0100-install-nodejs.hook.chroot
│ ├── 0200-install-claude-code.hook.chroot
│ ├── 0300-create-user.hook.chroot
│ └── 0350-enable-persist.hook.chroot
└── includes.chroot/ # Files copied directly into the filesystem
├── etc/motd
├── etc/shells
├── etc/sudoers.d/claudios
├── etc/systemd/system/
│ └── claudios-persist.service
├── home/claudios/.claude/commands/
│ ├── reboot.md # /reboot slash command
│ └── shut-down.md # /shut-down slash command
└── usr/local/bin/
├── claudios-shell # Primary login shell
└── claudios-persist # Auto-persistence setup script
Emergency Recovery
If claudios-shell fails and you lose access:
From GRUB : edit the kernel line and append init=/bin/bash
From TTY : the claudios user has sudo — use chsh to change the shell
Field
Value
Username
claudios
Password
claudios
API key
managed by Claude Code
License
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
