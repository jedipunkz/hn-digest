---
source: "https://nully.chat/"
hn_url: "https://news.ycombinator.com/item?id=48836996"
title: "Show HN: Nully – FOSS AI chat without the bloat"
article_title: "nully: super simple AI chat. fast, private, self-hostable"
author: "johnfahey"
captured_at: "2026-07-08T20:56:48Z"
capture_tool: "hn-digest"
hn_id: 48836996
score: 2
comments: 0
posted_at: "2026-07-08T20:25:47Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Nully – FOSS AI chat without the bloat

- HN: [48836996](https://news.ycombinator.com/item?id=48836996)
- Source: [nully.chat](https://nully.chat/)
- Score: 2
- Comments: 0
- Posted: 2026-07-08T20:25:47Z

## Translation

タイトル: HN を表示: Nully – 肥大化しない FOSS AI チャット
記事タイトル：nully：超シンプルAIチャット。高速、プライベート、自己ホスト可能
説明: nully は、OpenRouter 用の高速、プライベート、自己ホスト可能なチャット アプリです。独自の API キーを持参し、数百のモデルから選択して、チャットを始めましょう。アカウントや追跡はなく、会話はデバイス上に残ります。
HN テキスト: ChatGPT や Claude のようなサイトの「高度な」機能 (エージェント モード、メモリ、画像生成、詳細な調査など) をまったく使用しない者として、これらのサービスの肥大化は、UX とパフォーマンスの両方でかなり退屈であることがわかりました。世の中には AI チャット インターフェイスが 100 万個ありますが、最も最小限で軽量な方法で単純なメッセージングと履歴だけを実行するものは見つかりませんでした。それで私は自分で作りました。 Nully は Go と Vanilla HTML/CSS/JS で書かれています。文字通り、メッセージを OpenRouter に送信し、結果を受信し、メッセージをローカルに保存するだけです。これは非常にシンプルで、非常にクリーンで、非常に軽量です (サイト上のパフォーマンス グラフを参照してください)。基本的なチャット機能、基本的なモデル設定があり、ほぼすべての OpenRouter テキスト モデルで動作します。 OpenRouter API を介した添付ファイルと基本的な Web 検索をサポートします。 Nully は主に自分用に作成しましたが、オープンソースにして安価な VPS でホストするコストはほとんどかからないので、他の誰かの役に立つのであればリリースしようと思いました。そこから収入を得ることはありませんし、収入を得る気もありません。アカウント、サブスクリプション、プロキシや仲介者（もちろん OpenRouter 以外）はありません。すべてのチャットを JSON にエクスポートでき、すべてがローカル マシンに残ります。これは 100% オープンソースであり、必要に応じて自己ホスト可能です。ご意見をお聞かせください。

記事本文:
超シンプルなAIチャット。高速、プライベート、自己ホスト可能です。
nully は、OpenRouter 用の軽量チャット アプリです。
独自の API キーを持参し、数百のモデルから選択して開始します
おしゃべり。アカウントなし、追跡なし、インストールするものは何もありません。
必要なものはすべて、不要なものは何もない
チャットは他の人のデータベースではなく、あなたのデバイスに保存されます。アカウントを作成したり、サインアップしたりする必要はありません。
すべての主要プロバイダーの何百ものモデルを切り替え、どの程度深く考えるかを調整し、速度や価格に合わせてルーティングします。
モデルが書き込みを開始するとすぐに返信が表示され、長い会話でもアプリは迅速に動作します。
データベースやビルドステップを持たない単一の小さなプログラム。 Go で実行するか、小さなコンテナをデプロイします。
画像、PDF、その他のファイルを送信します。モデルは Web を検索して、貼り付けたリンクを読み取ることもできます。
チャットをファイルにエクスポートし、いつでも元に戻すことができます。明るいテーマと暗いテーマがあり、携帯電話でうまく機能します。
2026 年 7 月に Chrome DevTools で測定: 各チャットの 1 回のロード
同じマシン上のアプリのメイン ページと同じシミュレート
接続 (9 Mbps、遅延 100 ミリ秒)。リソースがロードされました
キャッシュされたファイルを含む、各ページがフェッチされたすべてのものをカウントします。
使用されるメモリは、最初の 5 秒間のブラウザのピークです。
メッセージを送信すると、ブラウザから直接送信されます。
OpenRouter と選択したモデルプロバイダー - 決して渡されません
nully サーバー経由。履歴、設定、API キーは
デバイスに保存されており、ページはロックダウンされているため、キーは
OpenRouter にのみ送信されます。追加の保証をご希望の場合は、
1 つのトグルで、チャットを守らないことを約束するプロバイダーに限定されます
あなたのデータ。
git クローン https://github.com/robbiegwald/nully.git
cd ヌリ
./cmd/null を実行してください
Docker の方が好きですか? docker build -t nully 。 && docker run --rm -p 4321:4321 nully
あなたに与えます

小さな密閉されたコンテナ。またはセットアップをスキップして使用します
nully.chat でホストされているアプリ。

## Original Extract

nully is a fast, private, self-hostable chat app for OpenRouter. Bring your own API key, pick from hundreds of models, and start chatting. No account, no tracking, and your conversations stay on your device.

As someone who doesn't use any of the more "advanced" features on sites like ChatGPT or Claude (agentic mode, memory, image generation, deep research, etc), I found the bloat of these services, both in UX and in performance, to be pretty tedious. There are a million AI chat interfaces out there, but I could never find one that just did simple messaging and history in the most minimal and lightweight way possible. So I made my own. Nully is written in Go and Vanilla HTML/CSS/JS. It literally just sends messages to OpenRouter, receives the result, and stores your messages locally. It's very simple, very clean, and very lightweight (see performance graphs on the site). It has basic chat features, basic model settings, and works with just about any OpenRouter text model. It supports attachments and basic web search via the OpenRouter API. I made Nully mostly for myself, but the cost to open-source it and host it on a cheap VPS is practically nothing, so I figured I'd release it if it can help anyone else. I don't make any revenue off of it, and have no desire to. There's no accounts, no subscriptions, and no proxies or middlemen (other than OpenRouter of course). You can export all of your chats to JSON, and everything stays on your local machine. It's 100% open source and self-hostable if you want that too. Let me know what you think!

super simple AI chat. fast, private, and self-hostable.
nully is a lightweight chat app for OpenRouter .
bring your own API key, pick from hundreds of models, and start
chatting. no account, no tracking, nothing to install.
everything you need, nothing you don't
chats are saved on your device, not in someone else's database. there is no account to create and nothing to sign up for.
switch between hundreds of models from every major provider, adjust how deeply they think, and route for speed or price.
replies appear the moment the model starts writing, and the app stays quick even in long conversations.
a single small program with no database and no build step. run it with Go, or deploy a tiny container.
send images, PDFs, and other files. models can also search the web and read links you paste.
export your chats to a file and bring them back anytime. light and dark themes, and it works great on your phone.
measured with Chrome DevTools in July 2026: one load of each chat
app's main page on the same machine and the same simulated
connection (9 Mbps, 100 ms latency). resources loaded
counts everything each page fetched, including cached files;
memory used is the browser's peak over the first 5 seconds.
when you send a message, it goes straight from your browser to
OpenRouter and the model provider you chose — it never passes
through a nully server. your history, settings, and API key are
saved on your device, and the page is locked down so your key can
only ever be sent to OpenRouter. if you want an extra guarantee,
one toggle limits your chats to providers that promise not to keep
your data.
git clone https://github.com/robbiegwald/nully.git
cd nully
go run ./cmd/nully
prefer Docker? docker build -t nully . && docker run --rm -p 4321:4321 nully
gives you a small, locked-down container. or skip the setup and use
the hosted app at nully.chat .
