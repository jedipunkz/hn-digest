---
source: "https://keijilohier.com/blogs/recreating-my-homelab-with-infrastructure-as-code/"
hn_url: "https://news.ycombinator.com/item?id=48686430"
title: "Automated My Homelab with the Bat Stack (Bootc, Ansible, Terraform/OpenTofu)"
article_title: "I've Automated My Homelab with the BAT Stack (Bootc, Ansible, Terraform/OpenTofu)"
author: "Klohier"
captured_at: "2026-06-26T13:39:27Z"
capture_tool: "hn-digest"
hn_id: 48686430
score: 1
comments: 0
posted_at: "2026-06-26T13:30:33Z"
tags:
  - hacker-news
  - translated
---

# Automated My Homelab with the Bat Stack (Bootc, Ansible, Terraform/OpenTofu)

- HN: [48686430](https://news.ycombinator.com/item?id=48686430)
- Source: [keijilohier.com](https://keijilohier.com/blogs/recreating-my-homelab-with-infrastructure-as-code/)
- Score: 1
- Comments: 0
- Posted: 2026-06-26T13:30:33Z

## Translation

タイトル: Bat Stack で My Homelab を自動化 (Bootc、Ansible、Terraform/OpenTofu)
記事のタイトル: BAT スタックでホームラボを自動化しました (Bootc、Ansible、Terraform/OpenTofu)
説明: Bootc、Ansible、Terraform/OpenTofu を使用して Geronimo Lab を再構築しました。このスタックにより、コード ツールとしてのインフラストラクチャを通じてホームラボを完全に管理できるようになり、保守、拡張、自動化が容易になります。

記事本文:
BAT スタックでホームラボを自動化しました (Bootc、Ansible、Terraform/OpenTofu) --> ホーム ブログを開始 ← ブログに戻る BAT スタックでホームラボを自動化しました (Bootc、Ansible、Terraform/OpenTofu)
Bootc、Ansible、Terraform/OpenTofu を使用して Geronimo Lab を再構築しました。このスタックにより、コード ツールとしてのインフラストラクチャを通じてホームラボを完全に管理できるようになり、保守、拡張、自動化が容易になります。
ホームラボでサービスを作成したり、思いつくものをセルフホスティングしたりして遊んだ後、これらのサービスを手動で管理するには時間がかかることに気づきました。
サービスは、通知があり、私がそれに到達したときにのみ更新されました。大きな変更を加えてもドキュメントは最新の状態に保たれていましたが、システムに加えた臨時の変更を書き留めるのを忘れていることに気づきました。これにより、ドキュメントとサーバーの実際の状態の間にズレが生じました。いざとなったときにホームラボを再建できるという自信はありませんでした。 BAT スタックは、Bootc、Ansible、Terraform/OpenTofu を組み合わせたものです。このスタックにより、コード ツールとしてのインフラストラクチャを使用してホームラボを完全に管理できるようになり、保守、拡張、自動化、回復が容易になります。
Bootc を使用すると、Docker や Podman でコンテナー イメージを作成するのと同じ方法で、コンテナー化されたオペレーティング システム イメージを作成できます。異なる構成の異なるシステムを維持する代わりに、すべてのシステムのベースとなる 1 つの基本イメージを維持できます。
FROM quay.io/fedora/fedora-bootc:44
RUN dnf install -y dnf5-plugins && \
dnf config-manager addrepo --from-repofile=https://pkgs.tailscale.com/stable/fedora/tailscale.repo && \
dnf インストール -y \
bash 補完 \
クロニー\
クラウド初期化\
ファイアウォール\
git\
jq\
qemu-ゲストエージェント\
テールスケール \
ヴィム\
vim で強化された \
yq && \
DNFすべてをクリーンアップ
#CIS

硬化
RUN dnf install -y openscap-scanner scap-security-guide && \
oscap xccdf eval --remediate \
--profile xccdf_org.ssgproject.content_profile_standard \
/usr/share/xml/scap/ssg/content/ssg-fedora-ds.xml ; \
dnf 削除 -y openscap-scanner scap-security-guide && \
DNFすべてをクリーンアップ
# 暗号化ポリシー
update-crypto-policies --set DEFAULT を実行します。
# SELinux カーネル引数
RUN mkdir -p /usr/lib/bootc/kargs.d
kargs.d/01-selinux.toml /usr/lib/bootc/kargs.d/01-selinux.toml をコピーします。
# SELinuxロックダウンサービス
COPY selinux/selinux-lockdown.service /usr/lib/systemd/system/selinux-lockdown.service
# サービスを有効にする
systemctl を有効にして実行 \
chronyd.service \
ファイアウォール.サービス \
qemu-guest-agent.service \
selinux-lockdown.service \
テールスケールサービス
bootc コンテナの lint を実行する
私のホームラボでは、研究室内のすべてのシステムが実行されている信頼できる情報源が 1 つだけになったため、システムの保守が容易になりました。 openscap と SELinux を使用したシステムの強化について学び、それを基本イメージに組み込んで、すべてのシステムがデフォルトで強化されるようにしました。基本イメージには、proxmox を通じてシステムを管理できるように、qemu-guest-agent と cloud-init も含まれています。私が追加した最後のサービスは、ホームラボにリモートで接続できるようにするための tailscale でした。 bootc の利点は、コンテナ ファイルを更新してレジストリにプッシュすると、必要に応じてロールバック サポートが組み込まれているため、すべての VM が新しいイメージに更新できることです。 Bootc は同じイメージベースのオペレーティングシステムを提供しますが、その上で実行されるアプリケーションを構成しません。ここで Ansible が登場します。
Ansible は、コードを使用してサーバーを構成するのに役立ちます。私のセットアップには、キャディを設定し、proxmox で動的インベントリを使用することで、サーバー スタック全体を構成する Playbook があります。これですべてのプレイが可能になります

書籍はサーバー上で現在実行されている内容で更新されます。新しい VM をセットアップしてプレイブックを実行すると、キャディによってドメイン名が自動的に作成されるため、IP アドレスを直接使用する必要がありません。 Ansible Automation Platform を組み込むことで、スケジュールに従って Playbook を実行し、認証情報を 1 つの中央の場所に安全に保存して、必要に応じて Playbook に挿入することができます。
セットアップした Playbook を表示する Ansible Automation Platform ダッシュボード
私は Ansible のウサギの穴に落ちましたが、ホームラボのネットワーク全体をセットアップするのにも Ansible を使用できることがわかりました。ユビキティ ユニファイ コントローラーを使用すると、Ansible を使用して VLAN とファイアウォール ルールを構成し、ホームラボ、IoT デバイス、ゲスト デバイスを相互にセグメント化することができました。これにより、ネットワーク構成のバージョンを管理できるようになったので、非常に役に立ちました。 Ansible の利点は、他のシステムと簡単に統合できるように利用できるコレクションです。再現可能なオペレーティング システムとそれを構成する方法がわかったら、VM をプロビジョニングする方法が必要になりました。ここで Terraform/OpenTofu が登場します。
Terraform/OpenTofu は、コードを使用してインフラストラクチャを作成および管理するのに役立ちます。 OpenTofu はコミュニティ主導の Terraform フォークであり、オープンソースであるため、私のラボでは OpenTofu を選択しました。 Terraform モジュールおよびプロバイダーとの互換性が維持されます。 proxmox の場合は、積極的にメンテナンスされ、最近更新されたため、bpg プロバイダーを使用しました。
構成の重複を避けるために、再利用可能な VM モジュールを作成しました。各仮想マシンは、名前、CPU、メモリ、ディスク サイズ、IP アドレスなどの異なる値のみを指定し、モジュールは共有構成を処理します。このモジュールを使用すると、必要な構成がすでにあり、すぐに使用できる新しい VM を簡単に作成できます。

Ansible によって管理されます。
リソース "proxmox_virtual_environment_vm" "this" {
名前 = var 。 vm_name
vm_id = var 。 vm_id
ノード名 = var 。 venv_node_name
description = "Terraform によって管理"
マシン = "q35"
bios = "ovmf"
開始 = true
タグ = var 。タグ
stop_on_destroy = true
エージェント {
有効 = true
}
CPU {
コア = var 。コア
タイプ = "x86-64-v3"
}
メモリ {
専用 = var 。記憶
}
efi_disk {
データストア ID = var 。データストアID
タイプ = "4m"
}
ディスク {
データストア ID = var 。データストアID
file_id = var 。画像ID
インターフェース = "virtio0"
iothread = true
破棄 = "オン"
サイズ = var 。ディスクサイズ
}
初期化 {
ip_config {
IPv4 {
アドレス = "192.168.50.${ var . ip_octet } /24"
ゲートウェイ = "192.168.50.1"
}
}
DNS {
サーバー = [ "192.168.50.213" ]
}
ユーザーアカウント {
ユーザー名 = "ansible"
キー = [ ファイル ( "~/.ssh/id_ed25519_ansible.pub" )]
}
}
ネットワークデバイス {
ブリッジ = "vmbr0"
vlan_id = var 。 vlan_id
}
}
地元の人々{
image_id = "local:import/geronimo-base.qcow2"
}
モジュール「キャディ」{
ソース = "./modules/vm/"
vm_name = "キャディ-01"
vm_id = 112
image_id = ローカル 。画像ID
コア = 2
メモリ = 2048
ディスクサイズ = 20
ip_octet = 200
タグ = [ "ブート" ]
}
この設定では、新しいモジュール ブロックを追加することで新しい VM を簡単に作成できます。このアプローチにより、ラボ全体が予測可能なベースラインを使用して作成されていることがわかります。これらすべてのツールが一緒になって BAT スタックを構成します。
BAT スタックと、それがどのように連携してホームラボを管理するか
これらすべてを連携させることは満足のいくもので、ホームラボのメンテナンス時に行っていた手作業の多くが削減されました。このスタックを使用すると、すべてが git でバージョン管理されるため、以前の状態に戻すことができるため、実験して何かを壊すことを恐れなくなります。セットアップをテストするために、ホーム ネットワークを含むラボ全体を消去しました。すべてを立ち上げて実行するのに 5 分かかりました

得。これは、すべてを手動で構成する必要があったときにかかっていた時間と比較すると、大幅な改善です。私の CI/CD パイプラインはすでに、bootc イメージの更新のビルドとレジストリへのプッシュを処理しています。次のステップは、Playbook と OpenTofu モジュールをラボにデプロイするパイプラインを処理し、git にあるものがすべて私のラボで実行されるようにすることです。私の GitHub リポジトリをチェックして、すべてがどのように設定されているかを確認してください。
連絡を取りたいですか?ぜひご意見をお聞かせください。

## Original Extract

I rebuilt Geronimo Lab using Bootc, Ansible, Terraform/OpenTofu. This stack allows me to manage my homelab entirely through infrastructure as code tools, making it easier to maintain, scale and automate.

I've Automated My Homelab with the BAT Stack (Bootc, Ansible, Terraform/OpenTofu) --> Home Blog Now ← Back to Blog I've Automated My Homelab with the BAT Stack (Bootc, Ansible, Terraform/OpenTofu)
I rebuilt Geronimo Lab using Bootc, Ansible, Terraform/OpenTofu. This stack allows me to manage my homelab entirely through infrastructure as code tools, making it easier to maintain, scale and automate.
After playing around with my homelab creating services and self hosting anything I can think of, I came to the realization that manually managing these services was taking a lot of time.
Services only got updated whenever they notified me and I got to it. While my documentation stayed up to date as I made major changes, I found myself forgetting to note down any ad hoc change I made to the system. This led to a drift between my documentation and the actual state of my servers. I had no trust that I would be able to rebuild my homelab if I had to. The BAT stack is a combination of Bootc, Ansible and Terraform/OpenTofu. This stack allows me to manage my homelab entirely with infrastructure as code tools, making it easier to maintain, scale, automate, and recover.
Bootc helps you to make containerized operating system images the same way you would make a container image with Docker or Podman. Instead of having to maintain different systems with different configurations, I can maintain one base image that all my systems are based on.
FROM quay.io/fedora/fedora-bootc:44
RUN dnf install -y dnf5-plugins && \
dnf config-manager addrepo --from-repofile=https://pkgs.tailscale.com/stable/fedora/tailscale.repo && \
dnf install -y \
bash-completion \
chrony \
cloud-init \
firewalld \
git \
jq \
qemu-guest-agent \
tailscale \
vim \
vim-enhanced \
yq && \
dnf clean all
# CIS hardening
RUN dnf install -y openscap-scanner scap-security-guide && \
oscap xccdf eval --remediate \
--profile xccdf_org.ssgproject.content_profile_standard \
/usr/share/xml/scap/ssg/content/ssg-fedora-ds.xml ; \
dnf remove -y openscap-scanner scap-security-guide && \
dnf clean all
# Crypto policy
RUN update-crypto-policies --set DEFAULT
# SELinux kernel args
RUN mkdir -p /usr/lib/bootc/kargs.d
COPY kargs.d/01-selinux.toml /usr/lib/bootc/kargs.d/01-selinux.toml
# SELinux lockdown service
COPY selinux/selinux-lockdown.service /usr/lib/systemd/system/selinux-lockdown.service
# Enable services
RUN systemctl enable \
chronyd.service \
firewalld.service \
qemu-guest-agent.service \
selinux-lockdown.service \
tailscaled.service
RUN bootc container lint
In my homelab, this made maintaining the system easier as there was now only one source of truth that every system in my lab was running. I learned about hardening systems using openscap and SELinux then incorporated that into the base image so that all my systems were hardened by default. The base image also includes qemu-guest-agent and cloud-init so that I can manage the systems through proxmox. The very last service I added was tailscale so that I could connect to my homelab remotely. The benefit with bootc is that when I make an update to the container file and push it to the registry, all my VMs are able to update to the new image with built in rollback support if needed. Bootc gives me the same image based operating system but does not configure the applications running on top of it. This is where Ansible comes in.
Ansible helps you configure your servers using code. In my setup I have playbooks that configure the entire server stack by setting up caddy and using dynamic inventories with proxmox. This allows all my playbooks to be updated with what's currently running on the server. When I set up a new VM and run the playbook, caddy automatically creates a domain name so that I don't have to use the IP address directly. Incorporated with Ansible Automation Platform I am able to have the playbooks on a schedule, have credentials stored securely in one central location, and be injected into the playbooks when needed.
Ansible Automation Platform dashboard showing the playbooks I have setup
I fell down the rabbit hole of Ansible and found that I could use it as well to setup the entire network for my homelab. Using my ubiquiti unifi controller, I was able to use Ansible to configure my VLANs and firewall rules to segment my homelab, iot devices, and guest devices from each other. This was huge as it allowed me to have my network configuration version controlled. The beauty of Ansible are the collections that are available to make it easy to integrate with other systems. Once I had a repeatable operating system and a way to configure it, I needed a way to provision the VMs. This is where Terraform/OpenTofu comes in.
Terraform/OpenTofu helps you create and manage your infrastructure using code. I chose OpenTofu in my lab as it is a community driven fork of Terraform and is open source. It maintains compatibility with Terraform modules and providers. For proxmox, I used the bpg provider because it was being actively maintained and recently updated.
To avoid duplicating configuration, I created a reusable VM module. Each virtual machine only specifies the values that differ, such as its name, CPU, memory, disk size, and IP address, while the module handles the shared configuration. By using this module, I can easily create new VMs that already have the configuration I want and are ready to be managed by Ansible.
resource "proxmox_virtual_environment_vm" "this" {
name = var . vm_name
vm_id = var . vm_id
node_name = var . venv_node_name
description = "Managed by Terraform"
machine = "q35"
bios = "ovmf"
started = true
tags = var . tags
stop_on_destroy = true
agent {
enabled = true
}
cpu {
cores = var . cores
type = "x86-64-v3"
}
memory {
dedicated = var . memory
}
efi_disk {
datastore_id = var . datastore_id
type = "4m"
}
disk {
datastore_id = var . datastore_id
file_id = var . image_id
interface = "virtio0"
iothread = true
discard = "on"
size = var . disk_size
}
initialization {
ip_config {
ipv4 {
address = "192.168.50. ${ var . ip_octet } /24"
gateway = "192.168.50.1"
}
}
dns {
servers = [ "192.168.50.213" ]
}
user_account {
username = "ansible"
keys = [ file ( "~/.ssh/id_ed25519_ansible.pub" )]
}
}
network_device {
bridge = "vmbr0"
vlan_id = var . vlan_id
}
}
locals {
image_id = "local:import/geronimo-base.qcow2"
}
module "caddy" {
source = "./modules/vm/"
vm_name = "caddy-01"
vm_id = 112
image_id = local . image_id
cores = 2
memory = 2048
disk_size = 20
ip_octet = 200
tags = [ "bootc" ]
}
With this setup I can easily create new VMs by adding a new module block. This approach allows me to know that my entire lab is created with a predictable baseline. All these tools together make up the BAT stack.
The BAT stack and how it all works together to manage my homelab
Getting all of this working together was satisfying, and cuts down a lot of the manual work I used to do when maintaining my homelab. With this stack I am less afraid to experiment and break things because I can revert to a previous state as everything is version controlled in git. To test my setup, I wiped my entire lab including my home network. It took 5 minutes to get everything up and running again. This is a huge improvement compared to the hours it used to take when I had to manually configure everything. My CI/CD pipeline already handles building and pushing bootc image updates to the registry. The next step is to have pipelines handling deploying the playbooks and OpenTofu modules to my lab so that whatever is in git is what is running in my lab. Check out my GitHub repository to see how I have everything setup.
Want to get in touch? I'd love to hear from you.
