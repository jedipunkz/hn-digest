---
source: "https://en.andros.dev/blog/09a21bdd/i-turned-unix-talk-from-1983-into-the-interface-for-my-ai/"
hn_url: "https://news.ycombinator.com/item?id=49397698"
title: "I turned Unix talk from 1983 into the interface for my AI"
article_title: "I turned Unix talk from 1983 into the interface for my AI | Andros Fenollosa"
image: "https://img.andros.dev/Ed2ya-x6rvSJ5fouyHe5ToaL4p8=/600x0/filters:format(avif):quality(85)/https://andros.dev/media/thumbnails/emilipothese-LEPhZkQbUrk-unsplash.jpg"
author: "andros"
captured_at: "2026-08-22T08:17:17Z"
capture_tool: "hn-digest"
hn_id: 49397698
score: 1
comments: 0
posted_at: "2026-08-22T08:14:41Z"
tags:
  - hacker-news
  - translated
---

# I turned Unix talk from 1983 into the interface for my AI

- HN: [49397698](https://news.ycombinator.com/item?id=49397698)
- Source: [en.andros.dev](https://en.andros.dev/blog/09a21bdd/i-turned-unix-talk-from-1983-into-the-interface-for-my-ai/)
- Score: 1
- Comments: 0
- Posted: 2026-08-22T08:14:41Z

## Translation

タイトル: 1983 年の Unix トークを AI のインターフェイスに変えました
記事のタイトル: 1983 年の Unix トークを AI のインターフェイスに変えてみた |アンドロス・フェノロサ
説明: talk コマンドに関する昨日の記事を書いた後、過去と未来の間には詩的なつながりがあることに気づきました。トークは性格を伝えます…

記事本文:
コンテンツにスキップ
ページ#実行"
data-liveview-function="ナビゲート_ホーム"
>
アンドロス・フェノロサ
エス
ページ#実行"
data-liveview-function="ナビゲートブログ"
>
ブログ
ブログ
ページ#実行"
data-liveview-function="ナビゲート_コース"
>
コース
コース
ページ#実行"
data-liveview-function="navigate_books"
>
本
本
ページ#実行"
data-liveview-function="navigate_talks"
>
会談
会談
ページ#実行"
データライブビュー機能 = "open_cv"
>
履歴書
履歴書
ページ#実行"
データライブビュー機能 = "open_contact"
>
お問い合わせ
お問い合わせ
1983 年の Unix トークを AI のインターフェースに変えました
スペイン語で読む
シェアする
talk コマンドに関する昨日の記事を書いた後、過去と未来の間には詩的なつながりがあることに気づきました。トークは文字ごとに送信し、LLM はストリーミングでトークンごとに送信します。これはプログラマーにとっては嬉しいものです。
そこで小さな橋を作りました。誰かが VPN から talk ai@host-where-the-ai-lives を実行すると、AI が反対側で応答します。
コツは真ん中に座ることです。サーバーは擬似端末 (PTY) 内で実際の talk を実行し、端末エミュレータ (pyte) で人間の半分を読み取り、talk の出力をクエリできる仮想画面に変換します。人間が Enter キーを押すと、その行がモデルに送信され、ストリーミングで到着すると、応答が 1 文字ずつ挿入されて返されます。摩擦なくフィットします。
フローチャート LR
H[「人間
トークアイ@ホスト
分割画面、ライブ"]
サブグラフ VPN["サーバー"]
D["話した
(inetd)"]
B["ブリッジ.py
パイトは人間を読む
入る = ターン終了
AI ライブのタイプ」]
終わり
P[「AI API」]
H |"UDP 518、ntalk (ネゴシエーション)"| D
H |"ダイレクト TCP (テキスト ストリーム)"| D
D -->|PTY| B
B -->|"リクエスト"| P
P -.->|"応答 (ストリーミング)"| B
talkd が AI を単なる別のユーザーとして認識できるように、コンテナーは起動時にその端末を utmp に登録します。その詳細がなければ、トーク ai@host は次のようになります。

「ログインしていません」。
それが動作するのを見るのは奇妙であり、同じくらい美しいものでもあります。40 年以上前のテクノロジーが VPN によって復活し、現在は相手側で人間が入力するのと同じように言語モデルと通信しています。確かにさらに進化する可能性はありますが、それは私を笑わせた愚かな実験にすぎません。
この作品は、表示-非営利-改変禁止 4.0 国際ライセンスに基づいています。
この記事では AI をどのように使用しましたか?
ページ#実行"
data-liveview-function="navigate_article"
データ-uuid="03a4ffb9"
>
前の記事
トーク: VPN によって復活した 1983 年の P2P チャット
私が決して公開しないものを読みたいですか？
完全な設定、雑なメモ、現在テストしているものなど、記事に当てはまらないものを送信します。スパムも販売目標到達プロセスもありません。
SUBSCRIBE という件名で、newsletter@andros.dev に電子メールを送信してください。
購読を解除するには、同じアドレスに UNSUBSCRIBE を送信します。
どのコーヒーも次の記事への背中を押してくれます。
トーク: VPN によって復活した 1983 年の P2P チャット
ウェブにアクセスしなくなった視聴者のために技術的な大聖堂を構築する
で作られました
Andros Fenollosa 著、Django LiveView を使用
ページ#実行"
データライブビュー機能 = "open_rss"
>

## Original Extract

After writing yesterday's article about the talk command, I realized there is a poetic bond between the past and the future: talk transmits character…

Skip to content
page#run"
data-liveview-function="navigate_home"
>
Andros Fenollosa
es
page#run"
data-liveview-function="navigate_blog"
>
Blog
Blog
page#run"
data-liveview-function="navigate_courses"
>
Courses
Courses
page#run"
data-liveview-function="navigate_books"
>
Books
Books
page#run"
data-liveview-function="navigate_talks"
>
Talks
Talks
page#run"
data-liveview-function="open_cv"
>
CV
CV
page#run"
data-liveview-function="open_contact"
>
Contact
Contact
I turned Unix talk from 1983 into the interface for my AI
Leer en español
Share
After writing yesterday's article about the talk command , I realized there is a poetic bond between the past and the future: talk transmits character by character, and an LLM emits token by token in streaming. It is candy for any programmer.
So I built a little bridge. If someone runs talk ai@host-where-the-ai-lives from the VPN, an AI answers on the other side.
The trick is to sit in the middle. The server runs the real talk inside a pseudo-terminal (PTY) and reads the human's half with a terminal emulator ( pyte ), which turns talk 's output into a virtual screen I can query. When the human presses Enter, that line is sent to the model, and the reply comes back injected character by character as it arrives in streaming. They fit without friction.
flowchart LR
H["Human
talk ai@host
split screen, live"]
subgraph VPN["Server"]
D["talkd
(inetd)"]
B["bridge.py
pyte reads the human
Enter = end of turn
types the AI live"]
end
P["An AI API"]
H |"UDP 518, ntalk (negotiation)"| D
H |"direct TCP (text stream)"| D
D -->|PTY| B
B -->|"request"| P
P -.->|"response (streaming)"| B
So that talkd recognizes the AI as just another user, the container registers its terminal in utmp on startup. Without that detail, a talk ai@host would reply "not logged in" .
Watching it work is strange and beautiful in equal parts: a technology more than forty years old, revived by a VPN, now talking to a language model that types just like a human would on the other side. It could surely go further, but it is only a silly experiment that made me smile.
This work is under a Attribution-NonCommercial-NoDerivatives 4.0 International license.
How did I use AI in this article?
page#run"
data-liveview-function="navigate_article"
data-uuid="03a4ffb9"
>
Previous article
talk: the P2P chat from 1983 that a VPN brought back to life
Do you want to read what I never publish?
I send what does not fit in the articles: full configurations, messy notes and whatever I am testing right now. No spam, no sales funnels.
Send an email to newsletter@andros.dev with the subject SUBSCRIBE and you are in.
To unsubscribe, send UNSUBSCRIBE to the same address.
Every coffee gives me a push toward the next article.
talk: the P2P chat from 1983 that a VPN brought back to life
Building Technical Cathedrals for an Audience That No Longer Visits Webs
Made with
by Andros Fenollosa using Django LiveView
page#run"
data-liveview-function="open_rss"
>
