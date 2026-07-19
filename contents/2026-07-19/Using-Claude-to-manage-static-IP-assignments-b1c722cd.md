---
source: "https://blog.herlein.com/post/unifi-fixed-hosts-skill/"
hn_url: "https://news.ycombinator.com/item?id=48969793"
title: "Using Claude to manage static IP assignments"
article_title: "Hey Claude, Set the IP address and name for the new BrightSign Player? · Greg Herlein"
author: "gherlein"
captured_at: "2026-07-19T17:50:24Z"
capture_tool: "hn-digest"
hn_id: 48969793
score: 1
comments: 0
posted_at: "2026-07-19T17:03:29Z"
tags:
  - hacker-news
  - translated
---

# Using Claude to manage static IP assignments

- HN: [48969793](https://news.ycombinator.com/item?id=48969793)
- Source: [blog.herlein.com](https://blog.herlein.com/post/unifi-fixed-hosts-skill/)
- Score: 1
- Comments: 0
- Posted: 2026-07-19T17:03:29Z

## Translation

タイトル: Claude を使用した静的 IP 割り当ての管理
記事のタイトル: クロードさん、新しい BrightSign プレーヤーの IP アドレスと名前を設定しますか? · グレッグ・ハーレイン

記事本文:
クロード、新しい BrightSign プレーヤーの IP アドレスと名前を設定しますか?
これは、私の人生の恥ずかしいほどの時間を奪った、小さくて愚かな問題です。BrightSign プレーヤーをネットワークに接続すると、UniFi コントローラーが配布しようとしている動的 IP がすべて取得されます。SSH 接続して実際の作業を行う前に、その IP を見つけなければなりません。クロードのスキルと新しいツールが助けになります!
30秒目のタスクです。何千回もやりました。 UniFi コンソールにログインし、クライアントをクリックし、リストを目を細めて、過去数分間に表示されたものを見つけ、IP をコピーして端末に貼り付けます。 30秒。予想外のリースを獲得した場合は、おそらく 1 分以内に狩りをしなければなりません。これは、Unifi コンソールが更新された場合に発生します。最新ではないように見える場合もあります。
30秒は1000回でも30秒ではありません。ゆっくりとした漏れです。そして、これらの小さなコンテキストの切り替え (ターミナルを離れ、ブラウザに移動し、検索し、戻ってくる) のたびに、時計が示すよりもコストがかかります。なぜなら、実際に行っていたことから私を引き離すからです。
ついに修正しました。構文を覚えなければならないスクリプトではありません。クロード・コードに話しかけて修正しました。
最近では、すべてが会話です。
私: ネットワーク上に新しいデバイスはありますか?
クロード: はい — 「BrightSign」のデバイスが 192.168.1.147 、MAC xx:xx:xx:... に現れ、最初に 2 分前に確認されました。
私: 固定 IP を与えて、 xt5-lab という名前を付けます。
クロード: これが追加する予約です (ホスト ブロックと予行演習を示します)。適用しますか？
終わり。そのプレーヤーは現在、永久に安定した名前と住所を持っています。その瞬間から、xt5-lab に ssh を実行すると、作業が始まります。狩りは必要ありません。コンソールはありません。コピー＆ペーストはできません。そして、次回再フラッシュまたは移動すると、予約がキー化されているため、同じアドレスに戻ります。

MACに。
そこで何が起こったのか考えてみましょう。私が望むものをわかりやすい英語で説明すると、マシンがネットワーク管理を行ってくれました。それがすべてです。
2 つの層: ツールとスキル
これは、私が構築した 2 つの部分を積み重ねて構築したために機能します。
最下層は gofi です。Web UI ではなくコードから UniFi コントローラー (私の場合は UDM Pro) を駆動するために私が書いた Go モジュールと CLI ツールキットです。 UniFi REST と WebSocket API をクリーンでタイプセーフな Go インターフェイスの背後にラップし、実際の作業をここで実行する 2 つの小さな CLI を同梱します。
gofimac — ネットワーク上のクライアント: MAC、IP、ホスト名、製造元をリストします (IEEE OUI ルックアップを実行します。これにより、デバイスが単なる匿名 MAC ではなく「BrightSign」であることがわかります)。また、最後のウィンドウに表示されていたデバイスや、 --since 7d 、 --gone など、なくなったデバイスを表示することもできます。それは、「新しいデバイスはありますか?それは何ですか?」です。半分。
gofips — ファイル形式として標準の ISC DHCP ホスト {} ブロックを使用して、固定 IP 予約 (DHCP 予約) の読み取りと書き込みを行います。これが「固定 IP と名前を付ける」部分です。
MIT ライセンスを取得しており、公開されています。 （私がそれを返したい気持ちを知っているでしょう。）
最上位のレイヤーは、Claude Code プラグインの unifi-fixed-hosts スキルです。これは、「CLI ツールを構築しました」を「ネットワークと通信するだけです」に変える部分です。
スキルとは基本的に、Claude Code にツールを正しく使用する方法、つまりいつツールに手を伸ばすべきか、どのフラグが何を意味するか、安全なワークフローとは何か、そして鋭いエッジがどこにあるかを教える小さな知識のパケットです。スキルがなければ、 gofimac -a は「すべてのクライアント」を意味しますが、 gofips -a は「ホストの追加」を意味することを覚えておく必要があります (はい、同じ短いフラグが 2 つのツール間で意味が異なります)

――わかってる、わかってる）。そのスキルを持ったクロードはすでにそれをすべて知っています。 gofimac --json で MAC を検索し、 gofips --get で現在の予約のスナップショットを取得し、ホスト ファイルを編集して、何かを操作する前に --dry-run でプレビューすることを認識しています。
ネットワークを破壊しない自動化
これが私が実際に誇りに思っている部分であり、派手な部分ではありません。
スキルには厳しいレビューゲートが組み込まれています。 DHCP 予約を変更すると、ライブ ネットワーク構成が変更されることになります。間違えると、2 つのデバイスに同じアドレスを与えたり、重要なものをネットワークから切り離したりする可能性があります。そのため、スキルは完全なホスト ファイルとドライラン出力を表示し、私が明示的に「はい」と答えるまで、変更の適用をきっぱりと拒否します。クロードに編集を依頼することは、適用を承認することと同じではありません。その違いはスキルに直接書き込まれます。
これは、人々がエージェント自動化に関して見逃していることです。目標は人間を排除することではありません。労力を取り除き、判断が重要な場所に人間をそのまま残すことです。 ISC DHCP 構文を手動で編集したくありません。ルーターに差分が届く前に、差分を確認したいと思っています。スキルが毎回その線を引いてくれるので、注意することを覚えておく必要はありません。慎重さが自動化です。
これを実際の仕事に結び付けてみましょう。そこに利益があるからです。
以前、BrightSign 拡張機能の開発について書きました。プレーヤー上で独自のネイティブ コードを実行します。プレーヤーは、実際には NPU を備えた単なる ARM64 Linux ボックスです。この作業はタイトなループです。コードを変更し、プレーヤーにプッシュし、SSH で接続し、実行し、失敗するのを確認し、修正し、繰り返します。一回のセッションで数十回。そして、それらの反復はすべて「プレイヤーに乗り込む」ことから始まります。
プレイヤーのアドレスが移動すると、反復ごとに「IP の検索」税が発生します。

。これに、午後のループ数を 12 回、ジャグリングするプレイヤーの数を掛け、午後の回数を掛けます。それはもはやスローリークではありません。それは穴の空いたバケツです。
プレーヤーを xt5-lab に一度ピン留めすると、税金がゼロになります。 ssh xt5-lab 。永遠に。私の筋肉の記憶では、IP が何であるかを知る必要さえありません。
これがこれらの小さな自動化の本当の話です。どれも大したことではありませんが、すべてを合わせると非常に大きなことになります。何千回も行う 30 秒のタスクは、まさに自動化する価値のあるものであり、それはまさに、自動化する気にならないものです。なぜなら、それはわずか 30 秒だからです。エージェント ツールはその計算を覆します。自動化の構築には、煩わしさの価値以上にコストがかかりました。今はそうではありません。
私は、まさにホームラボ、ワークフロー、奇妙な開発プレイヤーの山のためのインフラストラクチャを構築するというこのアイデアに何度も戻ってきます。
gofi は製品ではありません。誰も私にそれを建てるよう頼んだわけではありません。コードから特定の UniFi セットアップを管理するという私の特有のかゆみをかき消します。しかし、これは優れたインターフェイスを備えたクリーンなツールなので、スキルをラップするのはほとんど簡単でした。そして今では、私のネットワーク全体をただ話すことができるようになりました。デバイスを検出します。それが何なのか尋ねてください。名前を付けてください。ピン留めします。すべて同じターミナル内で私はすでに働いており、わかりやすく言うと、午後 11 時にばかげたことをしないように安全レールが設置されています。
ツールがその機能を実現しました。スキルのおかげでそれは楽になりました。そして、それが返す時間は、一度に 30 秒ずつ、本物です。
これらの部分は両方ともオープンソースです。 UniFi を実行していてターミナル内に住んでいる場合は、 gofi とプラグインを入手し、独自のクロード コードを教えてネットワークを実行します。
これを使って何かを構築する場合、または「30 秒を 1,000 回」行う独自の自動化ストーリーをお持ちの場合は、ぜひご利用ください。

それを聞くために。 LinkedIn にメモを送ってください: linkedin.com/in/gherlein 。

## Original Extract

Hey Claude, Set the IP address and name for the new BrightSign Player?
Here’s a tiny, stupid problem that has stolen an embarrassing amount of my life: I plug a BrightSign player into my network, it grabs whatever dynamic IP the UniFi controller feels like handing out, and now I have to go find that IP before I can SSH in and do actual work. Claude skills and a new tool to the rescue!
It’s a thirty-second task. I’ve done it a thousand times. Log into the UniFi console, click into clients, squint at the list, find the thing that appeared in the last few minutes, copy the IP, paste it into a terminal. Thirty seconds. Maybe a minute if it grabbed a lease I didn’t expect and I have to hunt. That’s if the Unifi console updates - sometimes it does not seem up to date.
Thirty seconds, a thousand times, is not thirty seconds. It’s a slow leak. And every one of those little context-switches — leave the terminal, go to the browser, dig, come back — costs more than the clock says, because it yanks me out of whatever I was actually doing.
I finally fixed it. Not with a script I have to remember the syntax for. I fixed it by talking to Claude Code .
These days the whole thing is a conversation:
Me: Is there a new device on my network?
Claude: Yes — a device from “BrightSign” just showed up at 192.168.1.147 , MAC xx:xx:xx:... , first seen two minutes ago.
Me: Give it a fixed IP and call it xt5-lab .
Claude: Here’s the reservation I’ll add (shows me the host block and a dry-run) . Apply it?
Done. That player now has a stable name and address forever . From that moment on I ssh xt5-lab and I’m in. No hunting. No console. No copy-paste. And the next time I re-flash it or move it, it comes back to the same address because the reservation is keyed to its MAC.
Think about what just happened there. I described what I wanted in plain English, and the machine did the network administration. That’s the whole thing.
Two Layers: The Tool and the Skill
This works because of two pieces I built that stack on top of each other.
The bottom layer is gofi — a Go module and CLI toolkit I wrote for driving a UniFi controller (a UDM Pro, in my case) from code instead of the web UI. It wraps the UniFi REST and WebSocket APIs behind a clean, type-safe Go interface, and it ships two small CLIs that do the actual work here:
gofimac — lists the clients on the network: MAC, IP, hostname, and manufacturer (it does an IEEE OUI lookup, which is how it knows a device is a “BrightSign” and not just some anonymous MAC). It can also show you devices seen in the last window, or ones that have gone away — --since 7d , --gone , that kind of thing. That’s the “is there a new device, and what is it?” half.
gofips — reads and writes the fixed-IP reservations (DHCP reservations), using standard ISC DHCP host {} blocks as its file format. That’s the “give it a fixed IP and a name” half.
It’s MIT-licensed and it’s out there. (You know how I feel about giving it back .)
The top layer is the unifi-fixed-hosts skill in my Claude Code plugin. This is the part that turns “I built a CLI tool” into “I just talk to my network.”
A skill is basically a little packet of knowledge that teaches Claude Code how to use a tool correctly — when to reach for it, which flags mean what, what the safe workflow is, and where the sharp edges are. Without the skill, I’d have to remember that gofimac -a means “all clients” but gofips -a means “add a host” (yes, the same short flag means different things across the two tools — I know, I know). With the skill, Claude already knows all of that. It knows to find the MAC with gofimac --json , snapshot the current reservations with gofips --get , edit the host file, and preview with --dry-run before touching anything.
Automation That Doesn’t Nuke Your Network
Here’s the part I’m actually proud of, and it’s not the flashy part.
The skill has a hard review gate baked into it. Changing DHCP reservations is mutating live network config — get it wrong and you can hand two devices the same address, or knock something important off the network. So the skill flat-out refuses to apply a change until it has shown me the full host file and the dry-run output, and I’ve explicitly said “yes.” Asking Claude to make the edit is not the same as approving the apply. That distinction is written right into the skill.
This is the thing people miss about agentic automation. The goal isn’t to remove the human — it’s to remove the toil and keep the human exactly where judgment matters. I don’t want to hand-edit ISC DHCP syntax. I do want to eyeball the diff before it hits my router. The skill draws that line for me, every single time, so I don’t have to remember to be careful. The carefulness is the automation.
Let me connect this back to real work, because that’s where the payoff lives.
I’ve written before about developing BrightSign extensions — running your own native code on the player, which is really just an ARM64 Linux box with an NPU. That work is a tight loop: change code, push to the player, SSH in, run it, watch it fail, fix, repeat. Dozens of times a session. And every single one of those iterations starts with “get onto the player.”
If the player’s address moves around, every iteration pays the “find the IP” tax. Multiply that by a dozen loops an afternoon, times however many players I’m juggling, times however many afternoons. That’s not a slow leak anymore. That’s a bucket with a hole in it.
Pin the player to xt5-lab once, and the tax goes to zero. ssh xt5-lab . Forever. My muscle memory doesn’t even have to know what IP it is.
That’s the real story of these small automations: none of them is a big deal, and all of them together is an enormous deal. The thirty-second task you do a thousand times is exactly the thing worth automating, and it’s exactly the thing you never bother to automate because… it’s only thirty seconds. Agentic tooling flips that math. Building the automation used to cost more than the annoyance was worth. Now it doesn’t.
I keep coming back to this idea of building infrastructure for one — for exactly my homelab, my workflow, my weird pile of dev players.
gofi isn’t a product. Nobody asked me to build it. It scratches my specific itch of managing my specific UniFi setup from code. But because it’s a clean tool with a good interface, wrapping a skill around it was almost trivial — and now my whole network is something I can just talk to . Detect a device. Ask what it is. Name it. Pin it. All in the same terminal where I’m already working, in plain English, with a safety rail so I don’t do something dumb at 11pm.
The tool made the capability. The skill made it effortless . And the time it gives back, thirty seconds at a time, is real.
Both of those pieces are open source. If you run UniFi and you live in a terminal, go grab gofi and the plugin , and teach your own Claude Code to run your network.
If you build something with it — or if you’ve got your own “thirty seconds a thousand times” automation story — I’d love to hear it. Drop me a note on LinkedIn: linkedin.com/in/gherlein .
