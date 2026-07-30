---
source: "https://www.mrafayaleem.com/blog/sandboxing-claude-cli-with-tart-on-apple-silicon"
hn_url: "https://news.ycombinator.com/item?id=49108868"
title: "Sandboxing Claude CLI with Tart on Apple Silicon"
article_title: "Sandboxing Claude CLI with Tart on Apple Silicon • The Growth Engineer's Diary"
author: "iamspoilt"
captured_at: "2026-07-30T12:23:47Z"
capture_tool: "hn-digest"
hn_id: 49108868
score: 1
comments: 0
posted_at: "2026-07-30T12:04:10Z"
tags:
  - hacker-news
  - translated
---

# Sandboxing Claude CLI with Tart on Apple Silicon

- HN: [49108868](https://news.ycombinator.com/item?id=49108868)
- Source: [www.mrafayaleem.com](https://www.mrafayaleem.com/blog/sandboxing-claude-cli-with-tart-on-apple-silicon)
- Score: 1
- Comments: 0
- Posted: 2026-07-30T12:04:10Z

## Translation

タイトル: Apple Silicon 上の Tart による Claude CLI のサンドボックス化
記事のタイトル: Apple Silicon で Tart を使用して Claude CLI をサンドボックス化する • 成長エンジニアの日記
説明: Apple デバイス上でクロード CLI およびその他のハーネスを問題なくサンドボックス化することについての私の見解

記事本文:
成長エンジニアの日記 ブログ出版物 検索ダークテーマについて語る メニュー 戻る
0. VM を作成します (1 回限りのセットアップ)
1. ディレクトリ共有を使用して VM を起動します。
3. virtiofs 共有をマウントします (Linux は自動マウントしません)。
4. 再起動後もマウントを維持する
2026 年 6 月 19 日 5 分 個人 / 生産性 / AI リツール Apple Silicon 上の Tart による Claude CLI のサンドボックス化
Apple デバイス上でクロード CLI やその他のハーネスを問題なくサンドボックス化することについての私の見解
数か月前、私は個人の納税申告を支援するために Claude-Obsidian ベースのワークフローを活用することについて書きました。私は職場でも個人的な設定でも、これらのワークフローの同様の繰り返しを試してきましたが、苦労したことの 1 つは、悪名高い --dangerously-skip-permissions フラグなどを使用してさまざまな LLM ハーネスを安全に実行する方法です。
LLM ハーネスは、頻繁にアクセス許可を要求することなく作業を完了するのにはるかに効果的であることがわかりましたが、MacBook 上でそのような高いアクセス許可を使用して LLM ハーネスを実行することに快適さを感じたことはありませんでした。
Obsidian を使用して個人のナレッジ ベースを作成する場合、リモート Unraid ホームラボ サーバーと Macbook の両方でゲスト VM を使用し、Syncthing を使用して LLM 編集を MacBook に同期することになりました。ほとんどの場合、これはうまくいきましたが、私がやろうとした単純なことにとっては不必要に複雑であることが判明しました。
LLM が、ホスト OS 上の厳密に特定のディレクトリ セットに対する編集権限を持ちながら、パッケージのインストール、コードの実行、インターネットの参照などに対する無制限の権限を持つようなファイル ベースのワークフローを用意します。
このワークフローをサポートするために VMWare Fusion を使用して Ubuntu VM を作成しました。これには通常、完全な Ubuntu セットアップのインストールと、Syncthing または共有フォルダー (後者) を介してディレクトリを同期することが含まれます。

これは、暴走サンドボックスに望むよりも複雑です。また、VMWare Fusion VM ははるかに肥大化しており、ハーネス用のシンプルなサンドボックス環境で実現したいものに対してリソースを大量に消費していることもわかりました。
最近、Apple Virtualization フレームワーク上の薄い CLI ラッパーである Tart ↗ を見つけたので、それが要件に合うかどうかを確認するために試してみることにしました。
まず始めるのは非常に簡単で、その手順は次のとおりです。
0. VM を作成します (1 回限りのセットアップ) #
Cirrus Labs レジストリから基本イメージのクローンを作成します。
タルト クローン ghcr.io/cirruslabs/ubuntu:最新の ubuntu bash this.classList.remove('copied'), 2000)
">
必要に応じてディスクのサイズを変更します (GB 単位のサイズ。これは VM が停止しているときに実行します)。
Tart set ubuntu --disk-size 50 bash this.classList.remove('copied'), 2000)
">
これらは、次回の起動時にルート パーティションを自動拡張する Ubuntu クラウド イメージ (Cirrus Labs によってわずかにカスタマイズされています) であるため、手動での givepart/resize2fs は必要ありません。ゲスト内で df -h / を使用して確認してください。
1. ディレクトリ共有を使用して VM を起動します #
タルト実行 --dir=obsidian_vault:/Users/rafaypersonal/Documents/obsidian_vault ubuntu & bash this.classList.remove('copied'), 2000)
">
これにより、指定されたホスト ディレクトリが VM にマウントされます。これは、コード リポジトリ、個人のマークダウンなど、何でもかまいません。
形式は --dir=<name>:<host-path> です。<name> はゲスト内のサブディレクトリ名になります。
読み取り専用のホスト パスに :ro を追加します (例: --dir=obsidian_vault:/path:ro )。
複数の --dir フラグを使用できます。すべての共有は同じマウント ポイントの下に表示されます。
ssh admin@ $( Tart ip ubuntu ) bash this.classList.remove('copied'), 2000)
">
3. virtiofs 共有をマウントします (Linux は自動マウントしません) #
Tart は、すべての --dir 共有を com.apple.virtio-fs.automount のタグが付いた単一の virtiofs デバイスとして公開します。 VM 内で次の操作を実行します。

sudo mkdir -p /mnt/shared
sudo mount -t virtiofs com.apple.virtio-fs.automount /mnt/shared bash this.classList.remove('copied'), 2000)
">
各共有は、その <name> ラベルにちなんで名付けられたサブディレクトリとして表示されます。
/mnt/shared/obsidian_vault 平文 this.classList.remove('copied'), 2000)
">
(macOS ゲストは / Volumes/My Shared Files/<name> に自動マウントします。手動マウントが必要なのは Linux のみです。この例は Ubuntu に基づいています。)
4. 再起動後もマウントを維持する #
VM を起動するたびにマウント手順を繰り返す必要はないため、再起動後もマウント手順を保持できます。
VM 内の /etc/fstab に追加します。
com.apple.virtio-fs.automount /mnt/shared virtiofs rw,nofail 0 0 プレーンテキスト this.classList.remove('copied'), 2000)
">
ワンライナー (冪等):
sudo bash -c 'grep -q virtiofs /etc/fstab || echo "com.apple.virtio-fs.automount /mnt/shared virtiofs rw,nofail 0 0" >> /etc/fstab' bash > /etc/fstab'" onclick="
navigator.clipboard.writeText(this.dataset.code);
this.classList.add('コピーされました');
setTimeout(() => this.classList.remove('コピー'), 2000)
">
以下で検証します。
sudo mount -a && echo OK # 解析エラーなしで OK が出力されるはずです
マウント | grep virtiofs # マウントがアクティブであることを確認する bash this.classList.remove('copied'), 2000)
">
nofail が重要です。nofail がないと、--dir フラグなしで VM を起動すると、デバイスが見つからないときにハングするか、緊急モードに落ちてしまいます。
それで終わりです！ --dangerously-skip-permissions を使用して VM にクロードとその他のハーネスをインストールして実行できるようになり、システムがブリックされることに対する不快感が軽減されます。ただし、VM と共有したデータは保護されず、ホストのみが保護されることに注意してください。マウントされたパスへの書き込みを防止したい場合は、いつでも軽減策として :ro フラグを使用できます。
開発者として作業している場合は、loca 経由でホストから VM ポートにアクセスすることもできます。

lhost は単純に SSH ローカル転送によって実行されます。
ssh -f -N -L 8080:localhost:8080 -L 3000:localhost:3000 admin@ $( Tart ip ubuntu ) bash this.classList.remove('copied'), 2000)
">
結論 #
Tart の背後にある会社である Cirrus Labs は現在 OpenAI の一部であるため、ChatGPT/Codex デスクトップ アプリもサンドボックス化に Tart を使用していると思われます。
ただし、CLI 用の自己管理型の完全なサンドボックス環境は非常に自由であり、ホスト マシンのブリック化によるセキュリティ上の懸念を軽減しながら、優れた永続性を提供することがわかりました。開発ワークフロー用の軽量の自己完結型ネイティブで再現可能なヘッドレス VM として使用できることは、ここでは非常に新鮮です。
おそらく、このワークフローは、反復的な開発作業においては Docker よりも人間工学的です。VM のディスクは通常のマシンと同様に保持されるため、イメージ レイヤー、コンテナーのコミット、変更を適用するための再構築について考えることなく、インストールするだけで先に進むことができます。
Apple Silicon 上の Tart による Claude CLI のサンドボックス化 https://mrafayaleem.com/blog/sandboxing-claude-cli-with-tart-on-apple-silicon 著者 Muhammad Rafay Aleem 公開日 2026 年 6 月 19 日 著作権
CC BY-NC-SA 4.0
コーヒーを買ってきて☕。 Claude CLI と Obsidian による個人納税申告の改善

## Original Extract

My take on sandboxing Claude CLI and other harnesses on Apple devices without headaches

The Growth Engineer's Diary Blog Publications Talks About Search Dark Theme Menu Back
0. Create the VM (one-time setup)
1. Start the VM with the directory share
3. Mount the virtiofs share (Linux does NOT auto-mount)
4. Make the mount persist across reboots
Jun 19, 2026 5 min personal / productivity / ai-retooling Sandboxing Claude CLI with Tart on Apple Silicon
My take on sandboxing Claude CLI and other harnesses on Apple devices without headaches
A few months ago, I wrote about leveraging a Claude-Obsidian based workflow to help with my personal tax filing. I have been experimenting with similar iterations of these workflows both at work and in personal settings and one thing that I have struggled with is how to safely run various LLM harnesses with the infamous --dangerously-skip-permissions flag and the likes.
I have found LLM harnesses to be much more effective at completing work without constantly nagging me for permissions but never felt comfortable running them with such elevated permissions on my MacBook.
In case of working with Obsidian and creating a personal knowledge base, I ended up using guest VMs on both, a remote Unraid homelab server and my Macbook and using Syncthing to sync LLM edits back to my MacBook. For the most part, this worked fine, but proved to be unnecessarily complex for a simple thing that I was trying to do:
Have a file based workflow such that the LLM has unlimited permissions to install packages, execute code, browse the internet, etc. while having edit permissions on strictly a specific set of directories on the host OS.
I have used VMWare Fusion to create Ubuntu VMs to support this workflow, which usually involves installing a full Ubuntu setup and syncing directories either via Syncthing or Shared Folders, the latter of which is more involved than I would want for a runaway sandbox. I also find VMWare Fusion VM much more bloated and resource hungry for what I would want to see with a simple sandboxed environment for the harness.
Recently, I came across Tart ↗ which is a thin CLI wrapper on top of the Apple Virtualization framework and I decided to give it a spin to see if it fits the bill.
It’s quite simple to start with one and here is how it goes:
0. Create the VM (one-time setup) #
Clone a base image from the Cirrus Labs registry:
tart clone ghcr.io/cirruslabs/ubuntu:latest ubuntu bash this.classList.remove('copied'), 2000)
">
Resize the disk if needed (size in GB; do this while the VM is stopped):
tart set ubuntu --disk-size 50 bash this.classList.remove('copied'), 2000)
">
These are Ubuntu cloud images (lightly customized by Cirrus Labs) that auto-grow the root partition on next boot, so no manual growpart / resize2fs is needed — verify with df -h / inside the guest.
1. Start the VM with the directory share #
tart run --dir=obsidian_vault:/Users/rafaypersonal/Documents/obsidian_vault ubuntu & bash this.classList.remove('copied'), 2000)
">
This mounts the provided host directory into the VM. This can be anything such as your code repo, personal markdown, etc.
Format is --dir=<name>:<host-path> — the <name> becomes the subdirectory name inside the guest.
Append :ro to the host path for read-only (e.g. --dir=obsidian_vault:/path:ro ).
Multiple --dir flags are allowed; all shares appear under the same mount point.
ssh admin@ $( tart ip ubuntu ) bash this.classList.remove('copied'), 2000)
">
3. Mount the virtiofs share (Linux does NOT auto-mount) #
Tart exposes all --dir shares as a single virtiofs device tagged com.apple.virtio-fs.automount . Do the following inside the VM:
sudo mkdir -p /mnt/shared
sudo mount -t virtiofs com.apple.virtio-fs.automount /mnt/shared bash this.classList.remove('copied'), 2000)
">
Each share appears as a subdirectory named after its <name> label:
/mnt/shared/obsidian_vault plaintext this.classList.remove('copied'), 2000)
">
(macOS guests auto-mount at /Volumes/My Shared Files/<name> — only Linux needs the manual mount, and this example is based on Ubuntu.)
4. Make the mount persist across reboots #
Since you would not want to repeat the mount step on every VM start, you can persist it across reboots.
Add to /etc/fstab inside the VM:
com.apple.virtio-fs.automount /mnt/shared virtiofs rw,nofail 0 0 plaintext this.classList.remove('copied'), 2000)
">
One-liner (idempotent):
sudo bash -c 'grep -q virtiofs /etc/fstab || echo "com.apple.virtio-fs.automount /mnt/shared virtiofs rw,nofail 0 0" >> /etc/fstab' bash > /etc/fstab'" onclick="
navigator.clipboard.writeText(this.dataset.code);
this.classList.add('copied');
setTimeout(() => this.classList.remove('copied'), 2000)
">
Validate with:
sudo mount -a && echo OK # should print OK with no parse errors
mount | grep virtiofs # confirm the mount is active bash this.classList.remove('copied'), 2000)
">
nofail matters: without it, booting the VM without the --dir flag would hang or drop to emergency mode when the device is missing.
And that’s it! You can now install and run Claude and other harnesses in the VM with the --dangerously-skip-permissions and feel less uncomfortable about bricking your system. But note that the data you have shared with the VM is not protected, only the host is. You can always use a :ro flag as a mitigation if you want to prevent writes to your mounted paths.
If you are working as developer, you can also access VM ports from the host via localhost by simply SSH local forwarding.
ssh -f -N -L 8080:localhost:8080 -L 3000:localhost:3000 admin@ $( tart ip ubuntu ) bash this.classList.remove('copied'), 2000)
">
Conclusion #
Cirrus Labs, the company behind Tart is now part of OpenAI so I am assuming the ChatGPT/Codex Desktop apps also use Tart for sandboxing.
However, I have found self-maintained fully sandboxed environments for CLI very liberating, alleviating my security concerns of bricking my host machines while providing great persistence. Having it as a lightweight self-contained native and reproducible headless VM for development workflows is quite refreshing here!
Arguably, this workflow is more ergonomic than Docker for iterative dev work because the VM’s disk persists like a normal machine, so you just install things and move on without thinking about image layers, committing containers, or rebuilding to make changes stick.
Sandboxing Claude CLI with Tart on Apple Silicon https://mrafayaleem.com/blog/sandboxing-claude-cli-with-tart-on-apple-silicon Author Muhammad Rafay Aleem Published at June 19, 2026 Copyright
CC BY-NC-SA 4.0
Buy me a cup of coffee ☕. Improving personal tax filing with Claude CLI and Obsidian
