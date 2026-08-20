---
source: "https://josephthacker.com/hacking/2026/08/20/using-genius-hacking-models-for-important-things.html"
hn_url: "https://news.ycombinator.com/item?id=49374980"
title: "Using Genius AI Models for Important Things: Cheating at NetHack"
article_title: "Using Genius AI Models for Important Things: Cheating at NetHack · Joseph Thacker"
image: "https://josephthacker.com/assets/images/nethack-kraken-farm.png"
author: "rez0__"
captured_at: "2026-08-20T14:25:58Z"
capture_tool: "hn-digest"
hn_id: 49374980
score: 1
comments: 0
posted_at: "2026-08-20T14:22:17Z"
tags:
  - hacker-news
  - translated
---

# Using Genius AI Models for Important Things: Cheating at NetHack

- HN: [49374980](https://news.ycombinator.com/item?id=49374980)
- Source: [josephthacker.com](https://josephthacker.com/hacking/2026/08/20/using-genius-hacking-models-for-important-things.html)
- Score: 1
- Comments: 0
- Posted: 2026-08-20T14:22:17Z

## Translation

タイトル: 重要なことへの Genius AI モデルの使用: NetHack での不正行為
記事のタイトル: Genius AI モデルを重要なものに使用する: NetHack での不正行為 · Joseph Thacker
説明: GPT-5.6 Daybreak Blue を使用して、オンデマンド NetHack 5.0 のクラッシュ、アイテムの重複、クラーケンのファームを見つけ、NAO でハイスコアを取得した方法。

記事本文:
重要なことへの Genius AI モデルの使用: NetHack での不正行為 · Joseph Thacker
ジョセフ
タッカー
ハッキング、AI、信仰などについての私の考え。
重要なことへの Genius AI モデルの使用: NetHack での不正行為
これは、GPT-5.6 Daybreak Blue を使用してバッファ オーバーフローを検出し、アイテムを騙すためにゲームをクラッシュさせ、NetHack 5.0 リーダーボードで高スコアを獲得した方法です。
私はネットハックが大好きです。とても素晴らしいゲームです。そして私はハッキングやバグハンティングが大好きです。そこで、壮大なバグ賞金稼ぎであり、NetHack (およびその亜種) を約 300 回登ってきた Demo (別名 turb0 ) に会ったとき、私は彼に 100 万もの質問をしました。
彼が私に話した中で一番気に入ったのは、毎年恒例の NetHack トーナメントである Junethack で、ほぼ毎年どのように「ハッキング ラン」を行っているかということでした。彼は前年に見つけたバグを披露し、それを使って何らかの形でゲームを壊していました。
私もそれをやりたかったのです！しかし、私は彼ほどホワイトボックスの深い脆弱性が得意ではありません。幸いなことに、トップの AI モデルがついにそうなりました。そこで、GPT-5.6 デイブレイク ブルーを狩りにセットしました。数時間以内に、理想的なクラッシュが見つかりました。特定のアクション中のみではなく、いつでも誘発される可能性があります。
デイブレイク・ブルーは「はい、これは本物です。」と戻ってきました。
NetHack 5 には、resize_tty() に整数切り捨てのバグがありました。端末の幅が 32,767 を超えると負にラップされ、サイズ不足の割り当てと範囲外の書き込みが発生します。サイズを 65,529 列に変更すると、確実にゲームがクラッシュしました。
これは、プレイヤーがいつでもオンデマンドでクラッシュできるため、非常に強力でした。レベルの移行中、NetHack はグローバル チェックポイントを更新する前に、離脱レベルを保存しました。アイテムを落とし、降下し、 --More-- で一時停止し、クラッシュを引き起こしました。その後、リカバリは、ドロップされたアイテムを含む新しいレベル ファイルを古い gl と結合しました。

まだ在庫があったオーバル状態。その結果、チャージされた願いの杖とユニークなアーティファクトを含む複製ができました。
奇妙なことに、デモは過去にここで問題を仮説立てていました。
デモは 1 年前、まさに正しい質問をしていました。
ライブ NAO で完全なチェーンを再現しました。ここで修正箇所を確認できます。その後、元のパッチのバイパスが修正されています。メインページにクレジットを記載させていただきました。
それを修正した開発者である dtype に感謝します。
NAO (nethack.alt.org) は、ほとんどの人がプレイする主要なサーバーです。もう 1 つの主要なサーバーは Hardfought ですが、NAO はバニラサーバーです。そこにあるすべての願い、死、昇天は IRC に投稿されます。
願いを騙したり、たくさんの願いをしたりすることは、当然のことながら非常に騒がしいものです。他のプレイヤーはイライラして開発者をタグ付けし始めました。開発者は、私が非常識なスコアを獲得するための最良かつ最も簡単な方法、つまりダイリチウム結晶を騙す前にそれを修正する必要がありました。
他のプレイヤーは「どうしてそんなことが可能なのか理解できない」と気づき始めています。
彼らはそれにパッチを当ててくれた、そして私には大量のダイリチウム結晶を望むほどの願いはなかったが、たくさんの巻物と願いがあった。そこで私は魔法のマーカーが欲しいと願い、大量のクラーケンを後で大量虐殺を逆転（召喚）できるように、呪われた大量虐殺の巻物を書きました。クラーケンは最も多くの経験を与えてくれるので、代わりに経験のためにクラーケンを飼育することもできます。
呪われた大量虐殺の巻物の完全に通常の数。
クラーケンは赤いセミコロンです。
赤いセミコロンはすべてクラーケンです。
とにかく、最終的には高得点を獲得できました。
29,926,212ポイント。 NAO の NetHack 5.0 リーダーボードでナンバー 1。
ロドニーが IRC で結果を発表。
これは天才ハッキング モデルがのために作られた重要な仕事です:P また、nethack サーバー NAO は素晴らしく、このようなハッキングの実行がリーダーボード上で悪名を轟かせています。デさんのおかげで

インスピレーションを与えてくれて、そしてすべてにおいてとてもクールにしてくれた NAO 開発者 (主に dtype) に感謝します。
私のメール リストに登録すると、このようなコンテンツがさらに投稿されると通知されます。
Twitter/X にも感想を投稿しています。
あい
ハッキング
サイバーセキュリティ
関連記事
物事を言うことが重要です
2026 年 7 月 22 日
起動: rez0 のラスカルズ
2026 年 7 月 17 日
バグ報奨金特異点: 私たちのハックボット
2026 年 7 月 1 日

## Original Extract

How I used GPT-5.6 Daybreak Blue to find an on-demand NetHack 5.0 crash, duplicate items, farm krakens, and get the high score on NAO.

Using Genius AI Models for Important Things: Cheating at NetHack · Joseph Thacker
J o s e p h
T h a c k e r
My thoughts on hacking, ai, faith, and more.
Using Genius AI Models for Important Things: Cheating at NetHack
This is how I used GPT-5.6 Daybreak Blue to find a buffer overflow, crash the game in order to dupe items, and get the high score on the NetHack 5.0 leaderboard.
I love NetHack. It’s such an awesome game. And I love hacking/bug hunting. So when I met Demo (aka turb0 ), who is an epic bug bounty hunter and has ascended NetHack (and variants) around 300 times, I asked him a million questions.
My favorite thing he told me about was how he’d do a “hacked run” nearly every year at Junethack , the annual NetHack tournament. He would show off the bugs he found during the previous year and use them to break the game in some way.
I wanted to do that too! But I’m not nearly as good at deep white-box vulns as he is. Luckily, top AI models finally are. So I set GPT-5.6 Daybreak Blue on the hunt. Within a few hours, it found an ideal crash . It could be induced at any time rather than only during specific actions.
Daybreak Blue coming back with: “Yes, this is real.”
NetHack 5 had an integer-truncation bug in resize_tty() : terminal widths above 32,767 wrapped negative, leading to an undersized allocation and out-of-bounds write. Resizing to 65,529 columns and back reliably crashed the game.
That was unusually powerful because it gave players an on-demand crash at any exact moment. During a level transition, NetHack saved the departing level before updating its global checkpoint. We dropped an item, descended, paused at --More-- , and triggered the crash. Recovery then combined the new level file, containing the dropped item, with the old global state, where it was still in inventory. The result was a duplicate, including charged wands of wishing and unique artifacts.
Crazily enough, Demo had hypothesized an issue here in the past.
Demo was asking exactly the right question a year earlier.
I reproduced the complete chain on live NAO. You can see where they fixed it here , followed by a fix for a bypass in the original patch . I was credited on the main page.
Shoutout to dtype, the dev who fixed it.
NAO (nethack.alt.org) is the major server where most people play. The other main one is Hardfought , but NAO is the vanilla server. All the wishes, deaths, and ascensions there get posted in IRC .
Duping wishes and making tons of wishes is naturally very noisy. Other players were annoyed and started tagging devs. The devs got to fixing it before I used the best and easiest way to rack up an insane score: duping dilithium crystals.
Other players starting to notice: “I don’t understand how all of that is possible.”
They patched it, and I didn’t have enough wishes to wish for tons of dilithium crystals, but I had lots of scrolls and wishes. So I wished for magic markers and wrote cursed scrolls of genocide so that I could later reverse genocide (summon) lots of krakens. Krakens give the most experience, so I could farm them for experience instead.
A completely normal number of cursed scrolls of genocide.
Krakens are the red semicolons.
Every red semicolon is a kraken.
Anyways, I eventually got the high score .
29,926,212 points. Number one on NAO’s NetHack 5.0 leaderboard.
Rodney announcing the result in IRC.
This is the kind of important work genius hacking models were made for :P Also, the nethack server NAO is awesome and they let hacked runs like this live on in infamy on the leaderboard. Thanks to Demo for the inspiration, and to the NAO devs (mostly dtype) for being so cool about everything.
Sign up for my email list to know when I post more content like this.
I also post my thoughts on Twitter/X .
ai
hacking
cybersecurity
Related Posts
it's important to say things
22 Jul 2026
Launching: rez0's rascals
17 Jul 2026
The Bug Bounty Singularity: Our Hackbot
01 Jul 2026
