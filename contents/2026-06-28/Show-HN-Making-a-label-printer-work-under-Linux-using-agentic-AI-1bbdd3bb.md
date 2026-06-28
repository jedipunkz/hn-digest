---
source: "https://stefan.schueller.net/posts/making-a-label-printer-work-under-linux-using-agentic-ai/"
hn_url: "https://news.ycombinator.com/item?id=48707513"
title: "Show HN: Making a label printer work under Linux using agentic AI"
article_title: "Making a label printer work under Linux using agentic AI - Stefan Schüller"
author: "sschueller"
captured_at: "2026-06-28T14:33:33Z"
capture_tool: "hn-digest"
hn_id: 48707513
score: 1
comments: 0
posted_at: "2026-06-28T14:20:28Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Making a label printer work under Linux using agentic AI

- HN: [48707513](https://news.ycombinator.com/item?id=48707513)
- Source: [stefan.schueller.net](https://stefan.schueller.net/posts/making-a-label-printer-work-under-linux-using-agentic-ai/)
- Score: 1
- Comments: 0
- Posted: 2026-06-28T14:20:28Z

## Translation

タイトル: Show HN: エージェント AI を使用して Linux でラベル プリンターを動作させる
記事のタイトル: エージェント AI を使用して Linux でラベル プリンターを動作させる - Stefan Schüller

記事本文:
ステファン・シューラー
ステファン・シューラー
内容
Kilocode を救出する Go バージョン
Kilocode を救出する Go バージョン
エージェント AI を使用して Linux 上でラベル プリンターを動作させる
ステファン・シューラー
カテゴリ「プロジェクト」に含まれています 2026-01-18 2026-06-28 738 単語
4分 内容
はじめに
Kilocode を救出する Go バージョン
少し前に、このようなさまざまなサイズのラベルを印刷できる「安価な」中国製のラベルプリンターを購入しました。悲しいことに、Linux でセットアップした場合、多少はサポートされていましたが、印刷結果は悪くなりましたが、同じプリンタが Windows または Android アプリ経由でまともな印刷を生成しました。何を試しても CUPS ではより良いプリントを得ることができませんでした。デフォルトの用紙サイズもいくつか設定してみましたが、非常に面倒でした。
最近まで、印刷したいラベルの PDF を携帯電話に送信し、アプリを使用して Bluetooth 経由で印刷していました。あるいは、Windows VM を使用して USB 経由で印刷します。
もっと良い方法があるはずだと思いました。そこで私がやったことは次のとおりです。
Android アプリは Bluetooth 経由で非常にうまく機能し、セットアップは必要ありません。おそらく、Linux では USB 経由よりも Bluetooth 経由でより良い印刷結果を得ることができるでしょう。そうすれば、プリンターを PC から離れた場所に配置することもできます。
そこで、APK Extractor を使用して携帯電話から APK を抽出しました。 https://play.google.com/store/apps/details?id=braveheart.apps.apkextract&hl=ja
次に、オンラインの Jadx を使用して逆コンパイルし、それらのファイルをダウンロードしました。
当初の私の計画は、コードを調べて、どのような Bluetooth 特性を使用する必要があり、何を送信するかを把握することでした。これはすぐに退屈になり、多くの時間がかかる可能性があります。
しかし、私は怠け者だと感じていたので、逆コンパイルされたコードの解読に休日を費やしたくありませんでした。
Kilocode を使用して試してみることにしました。おそらく、エージェント AI は私よりも逆コンパイルされた APK を解読するのが簡単です。

私の Kilocode セットアップで動作するように、LiteLLM で多くの新しいモデルをセットアップします。
私は抽出した APK を新しいプロジェクトに配置し、そのコードを使用して go バージョンを生成したいことをエージェントに伝えました。なぜ行くのですか？ディープシークを使用すると、go vs Python で多くの良い結果が得られました。
com.print.label_v66/sources/com/print/label/bluetooth およびその他のフォルダーにあるコードを使用します。 Linux PC の「デスクトップ」フォルダーにある PDF を Bluetooth 経由で印刷できる Go スクリプトを書いてください。印刷するラベルの用紙サイズも設定できる必要があります。プリンターの Bluetooth ID は DD:0D:30:02:63:42 です。最初は DeepSeek Coder を使用してみました。通常、非常に安価で良好な結果が得られますが、結果は得られませんでした。プリンターには接続できますが、動作しませんでした。何かがおかしいのですが、何がおかしいのかわかりませんでした。おそらく、初期化が欠落している可能性があります。
自分でデバッグしてみる前に、まず Gemini 3 Pro に切り替えることにしました。最初は動作しませんでしたが、プリンタが応答するまで動作することができました。
悲しいことに、Gemini は遅く、多くの場合オーバーロード エラーが発生しますが、最終的には動作する Go アプリを入手し、コマンド ラインから呼び出すことができました。
print_label.go を実行します \
--Bluetooth\
--プリンターID DD:0D:30:02:63:42 \
--tspl \
--用紙サイズ 100\
--用紙の高さ 150 \
--pdf ラベル-サンプル.pdf \
--margin-x 5 エージェントとさらに何度かやり取りをした結果、必要な機能がすべて動作するようになりました。
動作する Go アプリを作成したら、エージェントに新しいフォルダーを作成し、Chrome 用の Web 専用バージョンを構築するように依頼しました (すべてのブラウザーが Bluetooth をサポートしているわけではありません)。これは最初の反復では機能しました。
これで、PDF をアップロードして変換し、プリンターが提供するオプションから選択できる Web UI ができました。通常のアプリやアプリで印刷位置を揃えるためのテストパターンを印刷することもできます。

川にはできません。
同様のプリンターをお持ちの場合は、それがあなたのプリンターでも動作するかどうか知りたいと思います。
AI が生成したものはここからダウンロードできます (理由により、逆コンパイルされた APK は除きます)。コードをレビューしていないので、おかしな点や行き止まりがあるかもしれません…
このプロジェクトを手放すわけにはいかなかったので、最終的にカスタム PCB を構築し、全体をスタンドアロンで実行する ESP32 バージョンを作成することになりました。プリンターから電力も供給されており、既存の Bluetooth モジュールのドロップイン代替品ですが、現時点では Bluetooth 機能がありません。
さらに、WiFi のセットアップに役立つ LCD が搭載されており、IP と現在のステータスが表示されます。
コードと PCB はここで見つけることができます: https://github.com/sschueller/XPrinter-ESP32-S3

## Original Extract

Stefan Schüller
Stefan Schüller
Contents
Kilocode to the rescue Go Version
Kilocode to the rescue Go Version
Making a label printer work under Linux using agentic AI
Stefan Schüller
included in category Projects 2026-01-18 2026-06-28 738 words
4 minutes Contents
Intruduction
Kilocode to the rescue Go Version
A while ago I purchased this “cheap” Chinese label printer which can print different sized labels like these . Sadly, when I set it up under Linux, although somewhat supported, the print results were bad, while the same printer produced decent prints in Windows or via the Android app . I could not get better prints with CUPS no matter what I tried. I also tried to set some default paper sizes, but it was a huge pain.
Until recently I would send a PDF of the labels I wanted to print to my phone and use the app to print via Bluetooth. Alternatively, I would use a Windows VM to print via USB.
I thought there had to be a better way. So here is what I did.
The android app works very well via Bluetooth and requires no setup. Perhaps I can get better prints via Bluetooth than via USB under Linux. I could then also place my printer further away from my PC.
So I extracted the APK from my phone using APK Extractor. https://play.google.com/store/apps/details?id=braveheart.apps.apkextract&hl=en
I then decompiled it using Jadx online and download those files.
Initially, my plan was to go through the code and figure out what Bluetooth characteristics I needed to use and what to send. This can get tedious fast and take a lot of time.
However I was feeling lazy and did not want to spend my day off deciphering decompiled code.
I decided to give it a try using Kilocode . Maybe an agentic AI has an easier time deciphering a decompiled APK than I do, and I recently set up a lot of new models in LiteLLM to work with my Kilocode setup.
I placed the extracted APK in a new project and told the agent that I wanted to to use the code to generate a go version. Why go? I have had a lot of good results with go vs python when using deepseek.
Using the code in com.print.label_v66/sources/com/print/label/bluetooth and the other folders. Write me a go script that can print a PDF via bluetooth on a Linux PC in the "desktop" folder. I also need to be able to set the paper size for the label to be printed. The printer's bluetooth ID is DD:0D:30:02:63:42 Initially, I tried using DeepSeek Coder, which usually provides good results for dirt cheap, but I wasn’t getting anywhere. Although it would connect to the printer, it was not working. Something was off, but I did not know what. Most likely, some initialization was missing.
Before trying to debug it myself, I decided to first switch to Gemini 3 Pro. Although initially it also did not work, I was able to work with it until the printer responded.
Sadly, Gemini is slow and a lot of the time I get overload errors but eventually I got a go app that worked and I could call from the command line.
go run print_label.go \
--bluetooth \
--printer-id DD:0D:30:02:63:42 \
--tspl \
--paper-size 100 \
--paper-height 150 \
--pdf Labels-Sample.pdf \
--margin-x 5 After a few more back-and-forths with the agent, I had all the features I wanted working.
Once I had a working Go app, I asked the agent to make a new folder and build a web-only version for Chrome (not all browsers support Bluetooth). This worked in the first iteration.
I now have a web UI where I can upload a PDF, convert it, and select from the options the printer offers. I can also print a test pattern to align the print, which the regular app or driver cannot do.
If you have a similar printer, I would be curious to know if it also works with yours.
You can download what the AI produced here (minus the decompiled APK, for reasons…). I haven’t reviewed the code, so there may be some oddities and dead ends…
I couldn’t let this project go so I ended up building a custom PCB and writing an ESP32 version that runs the entire thing standalone. It even gets its power from the printer and is a drop-in replacement for the existing bluetooth module, although at the moment it lacks the bluetooth functionality.
Additionally, it has an LCD that helps with the setup of the WiFi and displays the IP as well as the current status.
You can find the code and PCB here: https://github.com/sschueller/XPrinter-ESP32-S3
