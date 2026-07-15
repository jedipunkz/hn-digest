---
source: "https://blog.greg.technology/2026/07/14/llm-networking-with-mikrotik.html"
hn_url: "https://news.ycombinator.com/item?id=48927915"
title: "LLM Networking with MikroTik"
article_title: "LLM Networking with MikroTik | the greg technology blog"
author: "gregsadetsky"
captured_at: "2026-07-15T22:49:50Z"
capture_tool: "hn-digest"
hn_id: 48927915
score: 2
comments: 0
posted_at: "2026-07-15T22:23:27Z"
tags:
  - hacker-news
  - translated
---

# LLM Networking with MikroTik

- HN: [48927915](https://news.ycombinator.com/item?id=48927915)
- Source: [blog.greg.technology](https://blog.greg.technology/2026/07/14/llm-networking-with-mikrotik.html)
- Score: 2
- Comments: 0
- Posted: 2026-07-15T22:23:27Z

## Translation

タイトル: MikroTik を使用した LLM ネットワーキング
記事のタイトル: MikroTik を使用した LLM ネットワーキング |グレッグテクノロジーブログ
説明: 私は「バイブネットワーキング」とか「バイブクロティック」とか、その他の忌まわしい言葉を言うのを控えてきましたが、これは事実です。私はここ数か月間、LLM を使用していくつかのネットワークをセットアップしてきました。そして、物事は概して順調に進んでいます。

記事本文:
MikroTik を使用した LLM ネットワーキング |グレッグテクノロジーブログ
グレッグテクノロジーブログ
MikroTik を使用した LLM ネットワーキング
私は「バイブネットワーキング」とか「バイブクロティック」とか、その他の忌まわしいことを言うのを控えてきましたが、これは事実です。私はここ数か月間、LLM を使用していくつかのネットワークをセットアップしてきましたが、物事は概ね順調に進んでいます。
私はしばらくの間、MikroTik 機器のファンでした。簡単に言えば、この機器は信頼性が高く、安価であり、IoT 携帯電話ルーター、通常のルーター、スイッチ、ポイントツーポイント システムなど、大量のネットワーク ユース ケースをカバーしているということです。
MikroTik に関するよくある不満の 1 つは、その複雑な UI/構成です。ある意味、ネットワーキング自体が複雑であるため、それが真実かどうかはわかりません。つまり、考えているよりも奥が深いのです。おそらく、ネットワークに関する氷山形式のミームが存在するはずです。一番上には「IP アドレス」があり、さらに深く行くと、「あの男」やその他の聞いたことのない名前が表示されます。
私が言いたいのは、そう、ネットワーキングは単に難しいということです。私はここしばらく、友人や友人のオフィス用のネットワークをセットアップしたり、ケーブルを作成したり、小さなパネルにパッチを当てたりなど、半ばアマチュアっぽくネットワークづくりをしてきました。公式の「認定ルーティング エンジニア」認定資格には、ほぼ確実に合格できませんでした。まあ、たくさん勉強しない限りはありません (自分を信じてください)。
でもご存知のとおり、私はそれを気にせずに、人々の Wi-Fi を修理したり楽しんだりすることをやめていません (そして、そうすべきではないと思います?) (ほとんどの人の Wi-Fi の問題は次のように分類できます…ああ、兄弟 - これについては別の投稿をする必要があります。エクステンダーは絶対に使用しないでください)。
LLM はどこに当てはまりますか?まあ、他の場所（コーディングなど）と同様に、彼らはカオス的な力を乗算するものです - 彼らは間違いなくMikroTiksとネットワーク一般の構成方法を知っています

コーディングの場合と同様、しっかりと束縛したり、不信感を抱いたり、検証し続けたりすることはできますし、そうすべきですが、より速く進歩させることもできます。
そこでここ数か月間、私はネットワークをセットアップすることができました (確かに小さいですが)。クロード コードにデバイスへのアクセスを許可し、その動作をさせるのをとても楽しんでいます。
私は自分自身のために役立つかもしれない、または役立つかもしれないメモの短いリストをまとめました。1 つのケースでは、既存のネットワーク (Wi-Fi が統合された非常に小さな 1 台のルーターから、ルーター + スイッチ + 2 つのワイヤレス アクセス ポイントへ) を移行していましたが、他の 2 つのケースでは、ネットワークはまったく新しいものでした。
いずれにせよ、役立つ場合は、Greg による LLM MikroTik のヒントとコツのリストを以下に示します。
mikrotiks は ssh 経由で設定できますが、llms がテキストをパイプで送受信しようとすると「1,000 カットによる死亡」が発生する傾向があります。より良い (つまり、より llm ネイティブな) チャネルは、REST/JSON API を使用することです。
安全でないサービス（安全でない API ポート、www、telnet、ftp）を無効にすることをお勧めします。
変更前に設定全体をダンプし、変更後に再度ダンプします。ソースのバージョン管理は素晴らしいです。自動バックアップ ツール (まだ構築できていない) を用意するのが最善でしょう。
CAPsMAN は本当に巨大な Wi-Fi シンプリファイアです - llms での設定は非常に簡単です
私はしばしば、複数の llms (反重力、コーデックス、オーパス、寓話) に構成を再確認し、何かが欠落しているか、またはひどく間違っているかどうかを確認するために合意に達するように依頼するという「トリック」に戻ります。
おそらく明白ですが、ネットワークを切断する前に (mikrotik に移行する場合)、SSID、パスワード、DHCP の予約に注意してください。
回復ランブックがあると役に立ちます。すべてのデバイスの設定を既知の P から復元する必要がある場合の手順を書き留めます。

レース。 Runbook を実行します。テストされていないバックアップは、ゼロでいっぱいのファイルである可能性があります。
llm の場合はよくあることですが、タスクを最小限にして 1 つずつ進めてください。はい、これは「ネットワークの設定を間違えないように」という冗談です。やめてください。設定を変更するたびにテストします。 llmsが幻覚を起こす！
非常に小さなことですが、構成しているすべてのデバイスに NTP (タイム サーバー) をセットアップすると便利です
これも小さいですが、健全性を保つために、ルーター、スイッチ、ワイヤレス アクセス ポイントなどのデバイスに名前を付けて識別し、わかりやすい名前を使用することをお勧めします。スイッチ上のポートについても同じことを行います。デバイスが移動するときに維持するのは少し面倒ですが、どのポートが何に接続されているかを知っておくと非常に便利です。
すべてのデバイスが同じ Routeros バージョンを実行するように必ず更新してください。LLMS は、コマンドがどのように機能するかを理解していると思っている場合もありますが、構文やオプションは変更されます。確認するよう依頼してください。
私は、IP アドレスがあちこちにあり、192.168.88.x ネットワークが重複しており、たとえイーサネット経由でそれらのデバイスに物理的に接続していても (常にそうなっているはずです)、ルーターやスイッチに接続することすら困難で、一般に混乱している状況に陥ったことがあります (あるいは、LLM が私をこの道に導いたのかもしれません。..)
私の意見では、最良の答えは L2「MAC Telnet」、つまり L2 (MAC アドレス) 層を介して Telnet 接続できるサーバーです。これは、WinBox を使用するのと同じようなものです。驚いたことに、WinBox はクロスプラットフォームになり、Mac でも非常にうまく動作します。 L2 Telnet クライアントがあると、LLM が mikrotik と通信できるようになります。WinBox は、LLM が制御できない GUI です。
このために、私は MAC-Telnet を強くお勧めします。これは、最悪の場合、つまり IP アドレスが機能しない場合に役立ちます。私/クロードは、インストールを簡単にするために小さな Homebrew 式を作成しましたが、次のこともできます。

元のインストール手順を参照してください - 同じコードです。また、MAC-Telnet をもう少し LLM に使いやすくするためにこの小さな CLI も作成しましたが、一般的に言えば、LLM は CLI ツールの使用方法を自分で理解します。
楽しんでください。私が間違っていると遠慮なく言ってください。さよなら！！ xx

## Original Extract

I’ve been refraining from saying ‘vibe networking’ or ‘vibkrotik’ or some other abomination, but it is true - I have been using LLMs to setup a few networks these last few months, and things have generally gone over swimmingly.

LLM Networking with MikroTik | the greg technology blog
the greg technology blog
LLM Networking with MikroTik
I’ve been refraining from saying ‘vibe networking’ or ‘vibkrotik’ or some other abomination, but it is true - I have been using LLMs to setup a few networks these last few months, and things have generally gone over swimmingly.
I’ve been a fan of MikroTik equipment for a while - the short story is that the equipment is reliable, inexpensive, and they cover a ton of networking use cases - IoT cell phone routers, regular routers, switches, point to point systems, etc.
One of the usual complaints about MikroTik has been its complex ui/configuration. In a sense, I don’t know if that’s true inasmuch as networking is complicated in itself - as in, it goes deeper than one thinks? Maybe there should be a iceberg-format meme about networking. At the top you’d have “ip address”, and going deeper you’d see “the dude” and other you’ve-never-heard-of-them’s.
The point I’m trying to make is yeah, networking can just be hard. I’ve been half-networking, amateur-ishly, for a while now - setting up networks for friends and friends’ offices, making cables, patching small panels etc. I almost certainly couldn’t pass an official “Certified Routing Engineer” cert - well, not without studying a lot (believe in yourself).
But you know, it hasn’t stopped me (and I suppose it shouldn’t?) from having fun and fixing people’s wifi (most people’s wifi problems can be categorized into… oh brother - I should make another post about this. Never use extenders.)
Where do LLMs fit in? Well, as elsewhere (in coding, etc.) they are a chaotic force multiplier - they definitely know how to configure MikroTiks and networking in general, but they also still get things wrong, go off-path, etc. As with coding, you can/should keep a tight leash, and mis-trust/keep-verifying, but you can also make more progress faster.
And so these last few months I’ve been able to setup networks - small, for sure - and had a lot of fun giving claude code access to my devices and letting it do its thing.
I’ve compiled a short list of maybe/hopefully useful notes for myself - in one case, I was migrating an existing network (a very small single-router-with-integrated-wifi to a router+switch+two wireless access points), while in two other cases the networks were net new.
In any case, if useful, here’s Greg’s list of LLM MikroTik tips and tricks:
even though mikrotiks can be configured over ssh, there’s a “death by a thousand cuts” that tends to happen when llms try to pipe text back and forth. the much better (ie more llm native) channel is to use the REST/JSON api.
I recommend disabling insecure services - the non secure api port, www, telnet and ftp
dump the entire config before any change, and dump it again after. source version controlling those is great. having some automated backup tool (which I haven’t gotten around to build yet) would be the best
CAPsMAN is truly a huge wifi simplifier - configuring it with llms is an absolute breeze
I often come back to the “trick” of asking multiple llms - antigravity, codex, opus and fable - to double check the config and come to a consensus to see if anything is missing or terribly wrong.
perhaps obvious, but before tearing down a network (when migrating to a mikrotik), take note of ssid’s, passwords, dhcp reservations
having a recovery runbook helps! take down the steps of what to do if you need to restore all of your devices’ configs from a good-known-place. run the runbook - an untested backup might as well be a file full of zeros.
as often is the case with llm’s, minimize the tasks and go one by one. yes, this is the “setup my network don’t make mistakes” joke - don’t do that. test after every config change. llms hallucinate!
very small thing, but it’s useful to setup ntp (time server) on all devices you’re configuring
also small, but for sanity’s sake, it’s good to name/identify your devices - your router, switch, wireless access points, and use descriptive names. do the same for ports on a switch - it can be a bit of a pain to maintain as devices move around, but knowing which port connects to what comes in handy a lot.
make sure to update all of your devices so they’re running the same routeros version - llms also sometimes think they know how a command works but the syntax/options do change - ask them to verify.
I’ve been in situations (or maybe the LLM led me down this path..) where IP addresses are all over, you have overlapping 192.168.88.x networks, and it’s generally a mess and hard to even connect to the router or switch, even if you’re physically connected to those devices over ethernet (which you should always be)
The best answer in my opinion is the L2 “ MAC Telnet ” ie a server that lets you telnet over the L2 (MAC-address) layer. It’s sort of the equivalent of using WinBox – which to my surprise is now cross-platform, and works quite well on Macs. Having a L2 telnet client allows your LLM to talk to your mikrotik - WinBox is a gui that LLMs can’t control.
For this, I deeply recommend MAC-Telnet - it will come in handy at the worst time ie when IP addresses don’t work! I/Claude just built a tiny Homebrew formula to make its installation easier, but you can also follow the original installation instructions - it’s the same code. I also made this small CLI just to make MAC-Telnet a bit more LLM-friendly to consume/use, but generally speaking, LLMs will figure out how to use a CLI tool by themselves.
Have fun - and feel free to tell me I have it all wrong. Bye!! xx
