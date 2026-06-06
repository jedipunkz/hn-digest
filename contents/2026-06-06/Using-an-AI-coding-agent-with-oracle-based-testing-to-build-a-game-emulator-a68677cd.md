---
source: "https://keanw.com/2026/03/a-diary-of-an-agentic-retro-gamer-part-1.html"
hn_url: "https://news.ycombinator.com/item?id=48419824"
title: "Using an AI coding agent with oracle-based testing to build a game emulator"
article_title: "A diary of an agentic retro-gamer - Part 1 - Through the Interface"
author: "throwaway_2494"
captured_at: "2026-06-06T00:54:40Z"
capture_tool: "hn-digest"
hn_id: 48419824
score: 1
comments: 0
posted_at: "2026-06-05T23:43:13Z"
tags:
  - hacker-news
  - translated
---

# Using an AI coding agent with oracle-based testing to build a game emulator

- HN: [48419824](https://news.ycombinator.com/item?id=48419824)
- Source: [keanw.com](https://keanw.com/2026/03/a-diary-of-an-agentic-retro-gamer-part-1.html)
- Score: 1
- Comments: 0
- Posted: 2026-06-05T23:43:13Z

## Translation

タイトル: オラクルベースのテストで AI コーディング エージェントを使用してゲーム エミュレーターを構築する
記事タイトル: エージェントレトロゲーマーの日記 - パート 1 - インターフェースを通して

記事本文:
これは、カナダに拠点を置くリサーチ エンジニアリング チームのメンバー、パトリック ナドー氏によるゲスト シリーズの最初の部分です。この最近の投稿で言及しました。 Patrick は優れたライターであり、非常に経験豊富な開発者であり、エージェント コーディング ツールを使用して非常に興味深いプログラミングの課題を解決する探求について、個人的で微妙な視点を提供しています。
私が子供の頃、家には Intellivision ゲーム コンソールと Apple II コンピュータの 2 台しかコンピューティング デバイスがありませんでした。数十年後、私が初めてコーディング エージェントを使ってコードを書いたとき、それらのマシンに何度も戻ってきました。
私がインテリビジョンの内部を初めて見たのは、妹が通気孔の 1 つにブレスレットを落としてしまい、再び動作させるために父がブレスレットを開けなければならなかったときでした。私も Apple II を開けて、CPU から RAM と ROM までの線をたどりました。おもちゃと道具の中間、それらのマシンは私の想像力の中に大きく浮かび上がり、ほどなくして私は BASIC から組み立てまでを卒業し、プログラミングマニュアルをじっくりと読んで時間を費やしました。
電気店を経営していた父の友人から、Apple II 用のギリシャ文字セットをデザインしてほしいと頼まれたことがありました。私は方眼紙上で文字をトレースし、バイナリ ファイルにエンコードして、それをダウンタウンの彼の店に持ち込みました。キャラクターセットをROMに焼き、新しいチップをマシンに取り付けました。彼が文字を一文字ずつ入力し、それぞれの文字が画面に表示されるたびに喜んで笑ったのを今でも覚えています。
方眼紙のデザインは今も残っていますが、初期の作品のほとんどは失われてしまいました。最近、私は別の初めてのことを記録する機会がありました。それは、初めてコンピューターと私が一緒にプログラムを書いたときです。何かが一周してきていると感じて、私は注意深くメモをとりました。
私は今でも古い Intellivision ゲームをプレイするのが好きです。

休日の周り。何年も前に廃止されたコンソールではなく、エミュレータ上でです。問題は、エミュレータが OS アップデートのたびに壊れてしまい、再び動作させるには醜いハックが必要になることでした。
他のプログラマと同じように、私はこの小さな不便さをプログラミングの大きな余談で解決することにしました。独自のエミュレータを最初からコーディングすることにしました。来年のクリスマスまで自分に猶予を与えました。
私は CPU から始めて、命令デコーダとコアを手書きで書きました。新しいコードを検証するために、使用していたエミュレータ jzintv から CPU 実装を抽出しました。これにより、単体テストで各命令のレジスタ、フラグ、RAM、サイクル数への影響を、古いけれども実証済みのバージョンと比較できるようになりました。
これは「テストオラクル」と呼ばれ、新しい作業を比較できる信頼できる参照です。構築には多大な労力がかかりましたが、プロジェクト全体の中で最も重要な設計上の決定であることが判明しました。
3 月中旬までに、一人で作業し、CP-1610 CPU コアはほぼ動作するようになりましたが、実際のマシンはまだありませんでした。バスもビデオもサウンドもありませんでした。これでゲームがプレイできるようになるまでには長い時間がかかるだろう。
ほぼ同時期に、私たちはコーディング エージェントの使用を開始するよう奨励されていました。私は懐疑的でした - あからさまに消極的でした。私は AI に、私が作成したパーサーの独自バージョンを作成するように依頼しました。ガイダンスがなければ、2 つの大きな欠陥のある日常的なコードが生成されました。
私のバージョンを AI に見せたとき、AI は私のバージョンが「このユースケースにとって重要なあらゆる面で優れている」ことを認めざるを得ませんでした。私はそれについてほくそ笑む Slack 投稿を書きました。AI は私の想像力の飛躍を見逃したのです!それでも、私は動作するエミュレータが欲しかったし、しっかりしたテストオラクルも持っていたので、今度は AI が動作するようにガイドしながら再試行しました。
5 時間目までに、画面上に最初のピクセルが表示され、見慣れたカラー テスト バーが表示されました。 10 時間までに、レンダリング

パイプラインは整備されていました。 21 時間目に、最初のカートリッジ ROM が起動しました。 28 時間までに、私のコレクションにある 204 個の ROM がすべて起動しました。 32時間目までに、画面上に動きがありました。もうすぐそこだった。 36 時間目までに、完全な Intellivision システムが稼働し、サウンド付きのコントローラーでプレイできるようになりました。
その後数日かけて、古いエミュレータでは事実上不可能だった機能を追加しました。
私は 2 番目の飛躍をしました。エミュレーターにデバッガー ポートを追加して、AI が実行中のマシンを覗き込み、その内部状態を検査し、ライブ プレイ中に制御できるようにしたらどうなるでしょうか?
AI がゲームを内側からコントロールする様子を、私は喜びながら見ていました。妻を呼んで「船を左に動かして」と入力したところ、その通りになりました。
私は AI にコード内で船が攻撃を受けた箇所を見つけるように依頼し、殺されたときにコントローラーがブザー音を立てるコードを追加しました。完全にバカだ。
最後には疲れ果ててしまいました。 5 日間、AI の肩越しに監視しながら、AI が順調に進んでいくのを確認すると、まるで自分で書いたかのような疲れを感じました。
そして、その高揚感が消えると、また不快感が戻ってきました。コーディングエージェントがこんなことをできるのは、苦労して学んだ何百万人ものプログラマーの仕事をキャプチャーし、それを流用したと言う人もいるが、そのためにのみ、他の人が学ぶことができるようにコードを自由に共有しているという事実が腹立たしい。
しかし、私のエミュレータは、Joe Zbciak が jzintv の開発に費やした長年の努力なしには実現できなかったでしょう。テストオラクルも、彼が苦労して獲得した知識を引き出すための自動化された方法ではなかったでしょうか?
私が驚いたのは、このプロセスが今でも非常に馴染み深いものであると感じられることでした。確かに、コードの大部分は AI が書きましたが、私たちが道に迷ったときもまったく同じように感じました。答えは依然として私から与えられなければならず、それでも私はそれを十分に望んでいなければなりませんでした。
ついにチップの中にたどり着いたようだ

当時、父がそのコンソールを割って開けたところを見ました。何かを取り戻したと同時に何かを失ったような気がしてなりません。
皆さんも私と同じようにこの記事を楽しんで読んでいただければ幸いです。今後の投稿では、パトリックの旅をさらに深く掘り下げていきます。
あなたのメールアドレスは公開されません。 * が付いているフィールドは必須です
次回コメントするときのために、このブラウザに名前、メールアドレス、ウェブサイトを保存してください。

## Original Extract

This is the first part of a guest series by Patrick Nadeau, a member of the Research Engineering team based in Canada. I mentioned it in this recent post . Patrick is a great writer and a very experienced developer who gives a personal, nuanced perspective on his explorations of using agentic coding tools to solve a really interesting programming challenge.
When I was a kid, there were only two computing devices in the house: the Intellivision game console and my Apple II computer. Decades later, when I wrote code with a coding agent for the first time, those were the machines I kept coming back to.
The first time I saw the insides of the Intellivision was the time my sister dropped her bracelet through one of the vent holes and my dad had to open it up to get it working again. I opened up my Apple II too, tracing the lines from the CPU to the RAM and ROMs. Halfway between toy and tool, those machines loomed large in my imagination, and before long I'd graduated from BASIC to assembly and spent long hours poring over programming manuals.
A friend of my dad's, who ran an electronics shop, once asked me to design a Greek character set for the Apple II. I traced the letters on graph paper, encoded them into a binary file, and brought it downtown to his shop. We burned the character set onto a ROM and popped the new chip into the machine. I still remember him typing the letters in one by one, and laughing in delight as each character appeared on the screen.
I still have the graph paper designs, but most of my earliest work is gone. Recently, I had the chance to capture another first: the first time a computer and I wrote a program together. Feeling that something was coming full circle, I kept careful notes.
I still like to play those old Intellivision games around the holidays. Not on the console — that was thrown out years ago — but on an emulator. The problem was that the emulator broke on every OS update, requiring ugly hacks to get it working again.
Like any programmer, I decided to solve this minor inconvenience with a huge programming digression: I would code my own emulator from scratch. I gave myself until next Christmas.
I started with the CPU, writing the instruction decoder and core by hand. To validate the new code, I extracted the CPU implementation from the emulator I'd been using, jzintv . This allowed my unit tests to compare each instruction's effects—on registers, flags, RAM, and cycle counts—against the old but tried-and-true version.
This is called a " test oracle ": a trusted reference you can compare your new work against. Building it took a lot of effort, but it turned out to be the most important design decision in the whole project.
By mid-March, working by myself, I had a mostly working CP-1610 CPU core, but no real machine around it yet: no bus, no video, no sound. It would be a long time before I could play a game on it.
Around the same time, work was encouraging us to start using coding agents. I was skeptical — openly reluctant. I asked the AI to write its own version of a parser I had written. Without guidance it produced workaday code with two major flaws.
When I showed the AI my version, it had to admit that mine was "better in every dimension that matters for this use case." I wrote gloating Slack posts about it: the AI had missed my imaginative leaps! Still, I wanted a working emulator, and I had a solid test oracle, so I tried again, this time guiding the AI as it worked.
By hour 5, I had the first pixels on screen, and saw those familiar colour test bars. By hour 10, the rendering pipeline was in place. At hour 21, the first cartridge ROM booted. By hour 28, all 204 ROMs in my collection were booting. By hour 32, we had movement on screen. We were almost there. By hour 36, we had a complete Intellivision system up and running, and I was playing through a controller with sound.
Over the next few days, we added features that would have been practically impossible with the old emulator.
I took a second leap: what if we added a debugger port to the emulator, a way for the AI to peer into the running machine, inspect its internal state, and control it during live play?
I watched, delighted, as the AI controlled the game from the inside. I even called my wife over and typed, "Move the ship left," and it did.
I asked the AI to find the spot in the code where the ship gets hit, and we added code to make the controller buzz when you get killed. Totally bonkers.
By the end, I was exhausted. Five days of watching over the AI's shoulder, making sure to keep it on track, left me almost as tired as if I'd written it myself.
And once the exhilaration wore off, my discomfort returned. It bugs me that coding agents can do this only because they've captured — some would say appropriated — the work of millions of programmers who learned the hard way, and then shared their code freely for others to learn from.
But my emulator would not have been possible without the years of work Joe Zbciak put into developing jzintv. Wasn't the test oracle also an automated way to extract his hard-won knowledge?
What surprised me was how familiar the process still felt. Sure, the AI wrote most of the code, but when we got lost, it felt exactly the same: the answer still had to come from me, and I still had to want it badly enough.
I suppose I finally got inside those chips I saw when my dad cracked open that console back then. I still can't help feeling that I reclaimed something and lost something at the same time.
I hope you enjoyed reading this as much as I did. In upcoming posts we'll dig more deeply into Patrick's journey.
Your email address will not be published. Required fields are marked *
Save my name, email, and website in this browser for the next time I comment.
