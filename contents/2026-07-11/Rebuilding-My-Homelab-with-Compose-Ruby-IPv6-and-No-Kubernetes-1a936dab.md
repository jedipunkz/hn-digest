---
source: "https://www.petekeen.net/homelab-resolved/"
hn_url: "https://news.ycombinator.com/item?id=48872031"
title: "Rebuilding My Homelab with Compose, Ruby, IPv6, and No Kubernetes"
article_title: "Rebuilding My Homelab with Compose, Ruby, IPv6, and No Kubernetes"
author: "zrail"
captured_at: "2026-07-11T14:19:09Z"
capture_tool: "hn-digest"
hn_id: 48872031
score: 2
comments: 1
posted_at: "2026-07-11T13:47:47Z"
tags:
  - hacker-news
  - translated
---

# Rebuilding My Homelab with Compose, Ruby, IPv6, and No Kubernetes

- HN: [48872031](https://news.ycombinator.com/item?id=48872031)
- Source: [www.petekeen.net](https://www.petekeen.net/homelab-resolved/)
- Score: 2
- Comments: 1
- Posted: 2026-07-11T13:47:47Z

## Translation

タイトル: Compose、Ruby、IPv6、Kubernetes なしでホームラボを再構築する

記事本文:
Compose、Ruby、IPv6、Kubernetes なしで Homelab を再構築する
ピート・キーン
エッセイ
Compose、Ruby、IPv6、Kubernetes なしで Homelab を再構築する
最終更新日 2026 年 7 月 10 日、読了時間 7 分
1年弱前、私は次のように書きました。
1 台のマシンが停止しており、おそらく戻ってこないでしょう。
サービスは生存者にランダムに分散されます。
私の Kubernetes プロジェクトは、ちょっと一時停止しています。
この投稿はその追記であり、ホームラボで解決できる範囲で、私が決定したことの概要を説明します。
少し要約すると、前回の投稿では、まとまったクラスターに似たものにマッシュされた多数のマシンがありました。
ヒプノトアド 、クラッシネーター 、およびシン クライアント上で実行される不安定なコントロール プレーン VM
lrrr 、Proxmox を実行する N150 ボックス
これら 4 台のマシンのうち、lrrr が唯一の生き残りです。
テレビに映るニブラーは、『フューチュラマ』に登場する小さな三つ目の愛らしい男だが、（ネタバレ注意だが、20年という歳月は私の限界をはるかに超えている）、タイムトラベルができる超知能種族の一員であることが判明し、実際、そもそもフライが冷凍ポッドに落ちた原因はニブラーだった。
nibbler は、私の地下室にあり、私が大きな希望を抱いていたマシンでした。それは i7-12700 と DDR5 メモリを搭載した Lenovo M80s Gen3 SFF で、eBay で盗品のように思われたもの (前兆) で拾われました。
余った小銭でメモリを128GB、ストレージをNVMeの3TBにアップグレードしました。
そこに HBA を取り付け、7 台の 16 TB ドライブを備えたディスク シェルフを吊り下げました。
これはコアサーバーとなり、すべてのアクションが行われるサーバーとなる予定でした。
先ほども触れたように、それはクソでした。
何が問題だったのかはまだよくわかりませんが、決して信頼できるものではありませんでした。
多くの場合、再起動後に起動しなくなりました。場合によっては、ハードウェアのことを忘れてしまうこともあります。
の上

私や私の配偶者は、Jellyfin がもう何も再生しないことに気づき、モニターとキーボードを接続するために地下室に足を引きずり、カーネル ログに不可解な SAS アダプター ドライバー エラーが記録され、ストレージ ディスク シェルフ全体がエラーになっているのを発見することがありました。
前回の記事を書いた頃に、ありそうなストーリーをまとめてみました。 eBayの売り手はこの機械が不良品であることを知っていたが、とにかく「eBay認定再生品」のラベルを貼り付けて、1曲分としてカモに売りつけた。
これに気づいたときには、マシンを大幅に改造していたので、返金請求を行うことができませんでした。
ニブラーの暴露と同時に、私は別の認識を得ました。Kubernetes は難しすぎるということです。
問題を解決するのに必要な時間やエネルギーがなかったため、保守方法が実際には分からないシステムを構築してしまいました。
クラスター内の他のすべてのマシンを合わせたよりも多くのコアとメモリを備えたマシンが、週に 1 ～ 2 回自己犠牲を決意すると、実験全体に大きな支障が生じます。
私の応急処置としては、重要なものをすべて k8s 以前の状態に戻すことでした。これは事実上単なる docker compose でした。
しばらく焼いた後、そのまま巻くことにしました。
一時的な解決策について彼らが何を言っているか知っていますか
lrrr は以前とほぼ同じ形式で存在します。
まだ Proxmox が実行されていますが、仮想的なものは、Omada を実行している LXC が 1 つと、エージェントのライフスタイルに慣れやすくするために使用している小さな VM が 1 つだけです。
私はまた、Proxmox フォーラムの荒らしたちにとって考えられる最悪のことも行いました。ホーム アシスタントやその友人などのコア サービスを実行するために、ホストに Docker をインストールしました。
ヒプノトードとクラシネーターはオフになっており、再び必要になるまで放置されています。
morbo は、私にとっては初めての HP Elitedesk 800 G3 SFF です。

TrueNAS SCALE の役割は、ストレージを管理して邪魔にならないようにすることであり、素晴らしい仕事をします。終わり。
ord-router は、たまたま Comcast と直接ピアリングしているシカゴのデータセンターで実行されている VM です。これはパブリックイングレスノードです。
nibbler は、同じメモリ、同じストレージ、新しいゲーマー スタイルのケース、マザーボード、プロセッサ (Ryzen 9 9900X) など、古い部品と新しい部品を組み合わせたものです。また、最近、ローカル LLM 実験用に RTX 5060ti 16GB ビデオ カードも入手しました。
これは「その他すべて」のマシンです。家が家として機能し、ネットワークがネットワークとして機能するために重要でないものはすべてここで実行されます。
とりわけ、nibbler は、Forgejo、Minecraft、Jellyfin とその友人たちが、他のランダムながらくたと一緒に住んでいる場所です。
このシステムとともに展開される Raspberry Pi クラスのデバイスもいくつか点在していますが、それらはアプライアンス タイプのことを行います。house-pi は Z-Wave スティックを実行し、shed-pi は ADS-B レシーバーが接続されている場所です。
追加の手順を含む Docker Compose
風変わりですか？そうだ、風変わりに行きましょう。
ホームラボをデプロイするソフトウェア スタックは次のように構成されています。
スパイシーな Docker 構成スタックの束
スパイスを実装する少数の Ruby ライブラリ
ライブラリを駆動するための Rakefile
すべてを結び付ける YAML ファイル
オプションの authn/authz およびオプションの公開 IP を使用した Caddy 経由の http イングレス
Lego による自動証明書管理
Ingress によって提供される静的 Web サイト
スタック間の依存関係の宣言
クロスプラットフォームの Docker イメージのビルド
各作成スタックは、作成拡張機能 (つまり、 x- で始まるフィールド) を使用して必要なものを宣言します。
簡単な例を次に示します。
サービス:
メチューブ:
画像: ghcr.io/alexta69/metube
コンテナ名: metube
ホスト名: metube
再起動: 停止しない限り
x-セキュリティ:
有効: true
ボリューム:
- メディア

-nfs:/メディア
環境:
ダウンロード_ディレクトリ: /media/youtube
ホスト: '::'
x-バックアップ:
有効: false
エックスウェブ：
ポート: 8081
認証: true
ダッシュボード:
名前: MeTube
サブタイトル: YouTube 動画をダウンロードする
カテゴリ: メディアとゲーム
x は以下に依存します:
- メディアボリューム
これにより、インターネット上のさまざまな場所からビデオを取得する小さなアプリである MeTube がセットアップされます。
安全なデフォルトを有効にし、バックアップが必要ないことを宣言してから、イングレスを設定します。
ここで、x-web はポートを指定し、認証/認可を必要とし、ダッシュボードに表示する必要があることを指定します。
この例では、スタック間の依存関係も示しています。
これは、media-nfs ボリュームを設定するメディア ボリューム スタックを取り込みます。
これらすべてを機能させる秘訣は、デプロイ時にすべてのスタックが 1 つの大きな docker-compose ファイルにまとめられることです。
つまり、外部ボリュームやネットワークを宣言するなどの大掛かりなことをする必要がなく、すべてが連携するだけです。
通常の Docker Compose の世界では、スタック内のすべてのコンテナーがブリッジ ネットワークを共有します。
私の目的からすると、それは意味がありません。
たとえば、昨日ダウンロードしたランダムなバイブコードが lldap と直接通信できるようにしたくありません。
代わりに、個別のスタックをデプロイした場合に得られるものをエミュレートします。各スタックはそれ自体へのブリッジ ネットワークを取得し、必要に応じて他のスタック内のコンテナーがそれに参加できます。
Ingress を実装する自然な方法は、Web サイトにサービスを提供するすべてのコンテナーを共有 Ingress ブリッジに参加させることです。
繰り返しになりますが、それは私が目指しているものとはまったく一致しません。
ポップされた押し出された ToDo リスト マネージャーが他のすべてを脅かさないように、スタック間の分離を維持したいと考えています。
私の解決策は決定論的 IPv6 です。
フリート全体として、RFC1918 IPv4 ネットワークと概念が似ている /48 IPv6 ULA を取得します (例:

10.10.10.0/24 ) ただし、非常に大きなプールから抽出されます
各ホストは、ホスト名と次の 16 ビットのソルトをハッシュすることで /64 を取得します。
各スタックは、スタック名と次の 32 ビットのソルトをハッシュすることによって /96 を取得します。
各サービスは、コンテナ名と最後の 32 ビットのソルトをハッシュすることによって /128 を取得します。
<48 ビット ULA プレフィックス>:<16 ビット ホスト>:<32 ビット スタック>:<32 ビット サービス>
Caddy は network_mode: host で実行され、生成されたすべてのアップストリームはコンテナーに割り当てられた IPv6 を指します。これは、ホストが自然にすべてのブリッジにアクセスできるため機能します。
私のセットアップに関するもう 1 つの奇妙/勇敢/愚かな点は、lrrr と nibbler の両方が、IPv6 経由でインターネットに公開されているポート 80 と 443 を持っていることです。
Caddy は、サービスが public: true ルートで特にオプトインしない限り、ローカル以外のトラフィックを拒否するように設定されています。
これは主に利便性のためであり、Jellyfin やダッシュボードなどにアクセスするためにモバイル デバイスに Tailscale をセットアップする必要はありません。
ああ、テールスケールも。
テールスケールはとてもクールです。私のiPhoneは例外で、非常に熱くなってバッテリーを消耗します。
Tailscale は、いくつかの重要なことのためのバックホールです。
ord-router に到達した Web トラフィックは、Tailscale 経由で地下のホストにルーティングされます。
これは、各ホストの Caddy が lrrr で実行されている中央の authn/authz システムと通信する方法でもあります。
まず、私は Kubernetes を構築したことがありません。
これについては実行時に何も変更されないため、etcd や Raft、あるいはいかなる種類のコンセンサスも必要ありません。
すべては、リンターやテスト、その他の優れた機能を使用して、事前に静的に決定されます。
次に、障害ドメインの概念は、実際には非常にうまく機能します。
lrrr には最小限に抑えられた重要なものが保管されています。
nibbler は、奇妙な 1 回限りのアプライアンス以外のものすべてです。
ニブラーが倒れても、人々は悲しみますが、怒っていません。
lrrr がダウンすると、HVAC システムが少し鈍くなり、

猫が外にいることを知らせるために私たちが頼りにしているLEDが機能しなくなります。
最後に、システムを完全に制御できるのは非常に便利です。
たとえば、先日、静的サイトを Ingress 経由で実装することにしました。
5 年も古い Kubernetes の議論から解決策を探して、それを翻訳して進める必要はありませんでした。
コードを書いてそれを実現しただけです。
IPv6 の場合も同様です。
それをスムーズに行うことができるのは、完全にコントロールできているからです。
ああ、それと、これがどれもあなたにとって面白くないと思われる場合は、Kubernetes または人々がまとめた無数の Docker Compose マネージャーの 1 つを使い続けることは完全に有効です。あなたが見えてます。大丈夫です。
私の Forgejo 上に全体を構築するコードのサニタイズされた公開スナップショットがありますので、ご興味があればご覧ください。
このようなことについてチャットしたいですか?私のソーシャルなどへのリンクはフッターにあります。
今は 2026 年なので、これは明確にする必要があると思います。これらは職人が作り出した人間の言葉です。機械は私のありきたりな常套句の修正を提案してくれましたが、私は機械の言葉を一切使いませんでした。

## Original Extract

Rebuilding My Homelab with Compose, Ruby, IPv6, and No Kubernetes
Pete Keen
Essays
Rebuilding My Homelab with Compose, Ruby, IPv6, and No Kubernetes
Last updated 2026-07-10, 7 min read
A little less than a year ago I wrote :
One machine is dead and probably not coming back.
Services are randomly scattered amongst the survivors.
My Kubernetes project is kinda-sorta paused.
This post is an addendum to that one, where I'm going to give an overview of what I've settled on, as much as one can settle in a homelab.
To recap a bit, in the previous post I had a bunch of machines that were mashed together into something resembling a coherent cluster:
hypnotoad , crushinator , and a precarious control plane VM running on a thin client
lrrr , an N150 box running Proxmox
Of those four machines, lrrr is the sole survivor.
Nibbler, on television, is an adorable little three-eyed dude in Futurama that (courtesy spoiler warning, but 20 years is well past my cutoff) turns out to be a member of a hyperintelligent race of beings that can time travel and who, in fact, is the reason why Fry fell into the cryopod to start with.
nibbler , in my basement, was a machine that I had pinned a lot of hope on. It was a Lenovo M80s Gen3 SFF with an i7-12700 and DDR5 memory, picked up on eBay for what seemed like a steal ( foreshadowing ).
I upgraded the memory to 128GB and the storage to 3TB of NVMe with left over pocket change.
I put an HBA in it and hung a disk shelf off of it with seven 16TB drives.
It was to be the core server, the one where all the action would happen.
It was, as covered earlier, a piece of shit.
I'm still not quite sure what was wrong, but it was never reliable.
Often it would just not come up after a reboot. Once in a while it would just forget about hardware.
On more than one occasion I or my spouse would discover that Jellyfin wouldn't play anything anymore, and I would shuffle down to the basement to attach a monitor and keyboard and discover that the entire storage disk shelf was errored out with inscrutable SAS adapter driver errors in the kernel logs.
Around the time I wrote the previous piece I put together a likely story. The eBay seller knew this machine was bad but they slapped an "eBay Certified Refurbished" label on anyway and sold it for a song to a sucker.
By the time I realized this I had modified the machine so extensively I couldn't make a refund claim.
Simultaneous to the nibbler revelations I had another realization: Kubernetes is Too Hard.
I built a system that I didn't actually know how to maintain without the time or energy necessary to dig myself out of trouble.
When the machine that has more cores and memory than every other machine in the cluster combined decides to self-immolate once or twice a week it puts a huge damper on the whole experiment.
My stop-gap solution was to revert everything important back to the way things were before k8s, which was effectively just docker compose.
After letting that bake for a while I decided to just roll with it.
You Know What They Say About Temporary Solutions
lrrr exists in much the same form as it did before.
It's still running Proxmox, but the only virtual things there are one LXC running Omada and one tiny VM that I've been using to ease myself into the agentic lifestyle.
I've also done the worst thing imaginable to the Proxmox forum trolls: I've installed Docker on the host to run core services like Home Assistant and friends.
hypnotoad and crushinator are turned off, lying in state until I need them again.
morbo is a new-to-me HP Elitedesk 800 G3 SFF that runs TrueNAS SCALE and whose entire job is to manage storage and stay out of the way and it does a great job. The end.
ord-router is a VM running at a datacenter in Chicago that happens to have direct peering with Comcast. This is the public ingress node.
nibbler is now an amalgamation of old and new parts: same memory, same storage, new gamer-style case and motherboard and processor (Ryzen 9 9900X). It also recently gained an RTX 5060ti 16GB video card for local LLM experiments.
This is the "everything else" machine. Everything that isn't critical to the house functioning as a house and the network as a network runs here.
Among other things, nibbler is where Forgejo, Minecraft, and Jellyfin and friends live, along with a bunch of other random crap.
There are also a handful of Raspberry Pi-class devices scattered around that are also deployed with this system, but they do appliance-type things: house-pi runs a Z-Wave stick and shed-pi is where the ADS-B receivers are plugged in.
Docker Compose with Extra Steps
Quirky? Yeah, let's go with quirky.
The software stack that deploys the homelab is arranged like this:
a bunch of spicy docker compose stacks
a handful of Ruby libraries that implement the spice
a Rakefile to drive the libraries
a YAML file that ties it all together
http ingress via Caddy with optional authn/authz and optional public-facing IP
automatic certificate management via Lego
static websites served by the ingress
cross-stack dependency declarations
cross-platform docker image builds
Each compose stack declares what it wants using compose extensions (i.e. fields that start with x- ).
Here's a simple example:
services:
metube:
image: ghcr.io/alexta69/metube
container_name: metube
hostname: metube
restart: unless-stopped
x-security:
enabled: true
volumes:
- media-nfs:/media
environment:
DOWNLOAD_DIR: /media/youtube
HOST: '::'
x-backup:
enabled: false
x-web:
port: 8081
auth: true
dashboard:
name: MeTube
subtitle: Download YouTube Videos
category: Media and Games
x-depends:
- media-volume
This sets up MeTube , a little app that grabs videos from various places around the internet.
It enables the secure defaults, it declares that it doesn't need backups, and then it sets up ingress.
Here, x-web specifies the port, that it wants authentication/authorization, and that it should show up on the dashboard.
This example also illustrates the cross-stack dependencies.
It pulls in the media-volume stack which sets up the media-nfs volume.
The trick that lets all that work is that all of the stacks get smushed into one big docker-compose file at deploy time.
That means I don't have to do gross things like declare external volumes or networks, it all just hangs together.
In a normal docker compose world every container in a stack shares a bridge network.
For my purposes that doesn't make sense.
I don't want some random vibe-coded thing that I downloaded yesterday to be able to talk directly to lldap, for example.
Instead, I emulate what you'd get if you were deploying individual stacks: each stack gets a bridge network to itself, and containers in other stacks can join it when necessary.
The natural way to implement ingress, then, would be to have every container that serves a website join a shared ingress bridge.
Again, though, that doesn't really fit what I'm going for.
I want to preserve the separation between stacks so that an extruded todo list manager that gets popped doesn't threaten everything else.
My solution is deterministic IPv6:
The fleet as a whole gets a /48 IPv6 ULA, which is similar in concept to an RFC1918 IPv4 network (ex: 10.10.10.0/24 ) but drawn from a vastly larger pool
Each host gets a /64 by hashing the hostname + salt for the next 16 bits
Each stack gets a /96 by hashing the stack name + salt for the next 32 bits
Each service gets a /128 by hashing the container name + salt for the last 32 bits
<48 bit ULA prefix>:<16 bit host>:<32 bit stack>:<32 bit service>
Caddy runs with network_mode: host and all of the generated upstreams point at the IPv6 assigned to the containers, which works because the host naturally has access to all of the bridges.
Another weird/brave/stupid thing about my setup is that lrrr and nibbler both have ports 80 and 443 exposed to the internet over IPv6.
Caddy is set up to reject non-local traffic unless the service specifically opts in with a public: true route.
This is mostly for convenience so I don't have to set up Tailscale on mobile devices to access things like Jellyfin and the dashboard.
Oh, and Tailscale .
Tailscale is so cool. Except on my iPhone which it makes very hot and drains the battery.
Tailscale is the backhaul for a few important things.
Web traffic that hits ord-router is routed to the hosts in the basement over Tailscale.
It's also how Caddy on each host talks to the central authn/authz system running on lrrr .
First, that I have not built a Kubernetes.
Nothing about this changes at runtime so I don't need etcd or Raft or consensus of any sort.
Everything is statically determined up front, with linters and tests and other nice things.
Second, the failure domains concept actually works really well in practice.
lrrr is home for critical stuff which is kept to a minimum.
nibbler is everything else that isn't a weird one-off appliance.
If nibbler goes down people are sad but not mad.
If lrrr goes down the HVAC system gets a little dumber and the LEDs we rely on to tell us the cats are outside stop working.
Finally, it's really nice to have full control over the system.
For example, the other day I decided to implement that static sites via ingress thing.
I didn't have to dig for a solution on some kubernetes discourse that is five years out of date and then translate it forward.
I just wrote some code and made it happen.
Same with the IPv6 thing.
Being able to do that smoothly is a consequence of having complete control.
Oh, and also, if none of this sounds like type one fun for you then sticking with Kubernetes or one of the zillion docker compose managers that people have put together is completely valid. I see you. It's ok.
I have a sanitized public snapshot of the code that builds the whole thing on my Forgejo if you're interested.
Wanna chat about this kind of thing? Links to my socials etc are in the footer.
It's 2026 so I guess this needs to be explicit: these are artisanal, human-produced words. The machines suggested fixes for my hackneyed clichés but I didn't use any of their words.
