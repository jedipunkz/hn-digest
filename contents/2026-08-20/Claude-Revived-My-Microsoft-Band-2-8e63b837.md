---
source: "https://webenclave.com/posts/bringing-a-microsoft-band-2-back-from-the-dead/"
hn_url: "https://news.ycombinator.com/item?id=49370598"
title: "Claude Revived My Microsoft Band 2"
article_title: "Reviving a Dead Microsoft Band 2 Over USB — Web Enclave"
image: "https://webenclave.com/og/posts/bringing-a-microsoft-band-2-back-from-the-dead.jpg"
author: "bchip"
captured_at: "2026-08-20T05:21:30Z"
capture_tool: "hn-digest"
hn_id: 49370598
score: 1
comments: 0
posted_at: "2026-08-20T05:11:32Z"
tags:
  - hacker-news
  - translated
---

# Claude Revived My Microsoft Band 2

- HN: [49370598](https://news.ycombinator.com/item?id=49370598)
- Source: [webenclave.com](https://webenclave.com/posts/bringing-a-microsoft-band-2-back-from-the-dead/)
- Score: 1
- Comments: 0
- Posted: 2026-08-20T05:11:32Z

## Translation

タイトル: クロードが私の Microsoft Band 2 を復活させました
記事のタイトル: 死んだ Microsoft Band 2 を USB 経由で復活させる — Web Enclave
説明: 私の Band 2 は、Microsoft が 2019 年に削除したアプリを要求してスタックしました。Claude Opus 5 がプロトコルを読み取って、MacBook から USB 経由でセットアップを解除しました。

記事本文:
コンテンツにスキップ
ウェブ エンクレーブ — EST 2023
安全なチャンネル — チューブ
スクロール000
エンクレーブ
仕事
2026年
派遣
Microsoft Band 2 を死から蘇らせる
2026 年 8 月 20 日 1 分で読めるチュートリアル
私の Microsoft Band 2 は、携帯電話の Microsoft Health アプリからセットアップを続行するときに停止しました。このアプリは、Microsoft が 2019 年 5 月 31 日にヘルス ダッシュボードを閉鎖して以来存在していません。FAQ ではそれについて率直に述べられています。その日以降にバンドをリセットすると、再度セットアップすることはできません。私のものはリセットされていました。
それを M1 MacBook Air に接続し、Opus 5 で Claude Code を開き、画面を説明しました。 1時間後、バンドは準備を終え、正しい時刻を告げていた。
ライブラリは msband-py で、電話を使わずに Band の USB プロトコルを話します。 Python 3.11 ではその構造型付けピンがインポートされないため、Python 3.10 が必要です。 OobeFinalize はペアリングされた電話がないと拒否するため、デモ モードを有効にし、Band を取り外し、無効にする方法があります。 NotInOobe を報告して戻ってきます。
歩数、心拍数、睡眠、トレーニング、アラームが機能します。通知、Cortana、その他クラウドを必要とするものはすべて永久に失われます。手首に戻ってきました。
AI、Web、アプリケーション開発コンサルティング
上記のような仕事を、曖昧な概要から本番まで進めます。
1 つのメッセージで、当社が適切なショップであるかどうかについての直接の回答が得られます。
プロジェクトを開始する すべてのディスパッチ
AI、Web、アプリケーション開発コンサルティング
構築が必要なものはありますか?
メッセージ 1 つで、当社がその商品に適したショップであるかどうかについての率直な回答が得られます。
チャンネルを開く
[メールで保護されています]
AI、Web、アプリケーション開発コンサルティング。 2023年から開催。

## Original Extract

My Band 2 was stuck asking for an app Microsoft deleted in 2019. Getting it out of setup over USB from a MacBook, with Claude Opus 5 reading the protocol.

Skip to content
Web Enclave — EST 2023
SECURE CHANNEL — TUBE
SCROLL 000
Enclave
Work
2026
Dispatch
Bringing a Microsoft Band 2 Back From the Dead
20 Aug 2026 1 min read tutorial
My Microsoft Band 2 was stuck on Continue setup from the Microsoft Health app on your phone . That app has not existed since Microsoft shut the Health Dashboard down on 5/31/2019 , and the FAQ is blunt about it: a Band reset after that date can never be set up again. Mine had been reset.
I plugged it into my M1 MacBook Air, opened Claude Code on Opus 5, and described the screen. An hour later the Band was out of setup and telling the right time.
The library is msband-py , which speaks the Band's USB protocol with no phone involved. It wants Python 3.10, because its construct-typing pin will not import on 3.11. OobeFinalize refuses without a paired phone, so the way through is demo mode: enable it, unplug the Band, disable it. It comes back reporting NotInOobe .
Steps, heart rate, sleep, workouts and alarms work. Notifications, Cortana and anything else that needed the cloud are gone for good. It is back on my wrist.
AI, web and application development consulting
We take work like the above from an ambiguous brief to production.
One message gets you a straight answer about whether we are the right shop.
Start a project All dispatches
AI, web and application development consulting
Got something that needs building?
One message, and you get a straight answer about whether we are the right shop for it.
Open a channel
[email protected]
AI, web and application development consulting. Held since 2023.
