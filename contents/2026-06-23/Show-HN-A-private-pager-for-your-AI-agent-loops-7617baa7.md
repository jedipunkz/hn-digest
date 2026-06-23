---
source: "https://ask-a-human.ai"
hn_url: "https://news.ycombinator.com/item?id=48640461"
title: "Show HN: A private pager for your AI agent loops"
article_title: "ask-a-human: a private pager for agent loops"
author: "alexandroskyr"
captured_at: "2026-06-23T05:18:58Z"
capture_tool: "hn-digest"
hn_id: 48640461
score: 3
comments: 1
posted_at: "2026-06-23T04:50:31Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A private pager for your AI agent loops

- HN: [48640461](https://news.ycombinator.com/item?id=48640461)
- Source: [ask-a-human.ai](https://ask-a-human.ai)
- Score: 3
- Comments: 1
- Posted: 2026-06-23T04:50:31Z

## Translation

タイトル: HN を表示: AI エージェント ループ用のプライベート ページャー
記事のタイトル: ask-a-human: エージェント ループ用のプライベート ページャー
説明: エージェントは全自動で実行されます。人間が必要になると、電話を鳴らして待ちます。エージェントと電話の間でエンドツーエンドで暗号化されます。アカウントもキーもありません。
HN テキスト: 私が解決している問題: 30 以上の完全自律型エージェントを実行しています。まれに、進むべき方向を決定できないためにブロックされる場合があります。そこで私は彼らのためのコミュニケーションチャネルを構築しました。したがって、ブロックされた場合でも、私に何でも質問できます。
基本的には、私に質問は来ないはずです。しかし、エージェントを自動化すればするほど、より多くのエッジケースが見つかります。
だからこそ私は、すべてのエージェントが同じ PWA アプリで私の携帯電話に接続される「ask-a-human」を構築しました。エージェントが私から何かを必要とするときに、プッシュ通知とバナーを受け取ります。それだけです。プラグアンドプレイでもあります。アカウントなし、設定なし、収益化なし。マジック ワームホールで動作し、チャネルはエンドツーエンドで暗号化されます。便利だと思ったらぜひ教えてください！

記事本文:
上で前に固定します
ペイント (FOUC なし) なので、ここにはテーマ切り替えスクリプトはありません。 localStorage ベースの
エージェントと連携して作業します
クロード コード コーデックス カーソル 副操縦士 ジェミニ 任意の MCP クライアント
MCP を話せれば、あなたに届きます。同じ 1 行の構成で、追加のセットアップは不要です。
これがあなたの携帯電話に届きます。
一度ペアリングしてください。その後、エージェントが自分で判断できない場合には、
電話を鳴らして待ちます。「はい」または「いいえ」、選択、または素早い応答を返します。
携帯電話で ask-a-human.ai/app を開き、出力された 10 文字のコードを入力して接続します。
すべてのテストは成功しましたが、このリリースでは本番環境に「orders.coupon_code」をドロップする移行も実行されます。発送しますか？
「main」は 20 分間赤くなっており、その背後で 3 つのデプロイがスタックしています。ブロックを解除するにはどうすればよいですか?
顧客がサポート スレッドでエスカレーションし、返金を要求していますが、私は承認できません。どのように返信すればよいでしょうか?
壁から電話まで、数秒で「続行」できます。
弊社では多数のエージェントを運営しております。それらの 100 個と同様に、ループ、フルオート、フル権限で、ほとんどの場合、私たちが眠っているとき、またはキーボードから離れているときに実行されます。それが重要な点だ。彼らは私たちを必要としていないはずだ。時々、どちらかがそうするまでは。
そして、フルオートのエージェントが単独で踏み込むべきではない一歩を踏み出したとき、私たちに連絡を取る方法はありません。百回ループを一時停止してドアのそばで待つことはできません。したがって、選択は両方とも悪いものでした。すべてのステップを子守して自動化を停止するか、自動化をリッピングさせて、私たちが眠っている間に何も刺激されないことを祈ります。私たちは 3 番目の選択肢を望んでいました。全員を全力で走らせて、まれに動けなくなった人には肩を叩いてもらいましょう。
そこで私たちはポケベルを作りました。 1 つの MCP サーバーをコピーしてエージェントに貼り付け、このサイトを iPhone のホーム画面に固定するだけです。インストールするもの、アカウント、API キーはありません。エージェントが「はい」か「いいえ」、選択、または素早い返答を必要とする場合、電話に ping を送信して待機します。あ

100 人のエージェント、1 つの場所、一度に 1 つの話題。
そしてそれは設計上プライベートです。メッセージはエンドツーエンドで暗号化された魔法のワームホールを通過するため、中間のリレーはブラインドになります。私たちは何も追跡せず、データベースもなく、ホスティング料金もすべて自分たちで支払います。販売するデータもファネルもキャッチもありません。私たちがそれを構築したのは、それが必要だからであり、私たちの唯一の目標はそれが普及することです。
ペアリングは Magic-Wormhole スタイルの SPAKE2 ハンドシェイクです。
その短いコードは強力な共有キーに変わり、たとえ誰かがそれをショルダーサーフィンしたとしても
オンラインでの推測は役に立たなくなる前に 1 回だけ取得されます。
その後、すべてのメッセージは NaCl で封印されます
秘密箱
(XSalsa20-Poly1305)。リレーは、base64(nonce‖ciphertext) のみを認識し、どの部屋が
内容については決して話しません。
リレーは暗号文を転送するだけです。
Base64(nonce‖ciphertext) + どの部屋がどの部屋と通信するか、決してデータではない
密封 密封 あなたのエージェント MCP · キーを保持しています リレー コンテンツ ブラインド あなたの電話機がキーを保持しています
SPAKE2: ショートコード → リレーが決して学習しないキー
SPAKE2・RFC 9382
リストレット255 · RFC 9496
魔法のワームホール
あなたを裏切らないように作られています。
信頼できるプライバシー ポリシーではありません。裏切らないデザイン。アカウントがありません
侵害、漏洩するデータベースはなく、召喚状を提出するログもありません。リレーは RAM のみです。再起動して、
再ペアリングするだけです。それでも私たちを信頼できない場合は、情報源はすぐそこにあります。
独自のリレーをホストできます
--リレー /
--パブリックリレー 。
サーバーは、base64(nonce‖ciphertext) と、どのルームがどのルームと通信しているかのみを確認します。それは愚かなパイプです。決してデータや決定ではありません。
短いペアリング コードは、Magic-Wormhole スタイルの SPAKE2 ハンドシェイクを介して強力な共有キーになります。中間中継者はそれを読み取ったり偽造したりすることはできません。
あなたのエージェントとあなたの電話が鍵を握っています。承認は片面で封印され、もう片面で開封されます。今

その間にいるよ。
RAM のみ、アカウントもデータベースもありません。再起動とは再ペアリングを意味します。フラグを使用してリレーと Web を自己ホストするか、GitHub 上のすべての行を読み取ります。
エージェントはすでに実行されています。
彼らに尋ねる方法を与えてください。
GitHub
📱
アプリを開く
📋 設定をコピーする
オープンソース & 自己ホスト可能 ·
GitHub ·
npm ·
llms.txt

## Original Extract

Your agents run full-auto. When one needs a human, it pings your phone and waits. End-to-end encrypted between your agent and your phone. No account, no key.

Problem Im solving: Im running 30+ fully autonomous agents. In some rare cases they get blocked because they cannot decide what direction to take. So I built a communication channel for them. So in case they get blocked they can ask me anything.
The base case should be that no questions should arrive to me. But the more I automate the agents the more edge cases I find.
That's why I built the `ask-a-human` where all my agents are connected into my phone in the same PWA app. I get push notifications and banners when agent needs something from me. That's all. It is also plug n play. No Account No configuration No monetization. It works with Magic Wormholes the channel is end to end encrypted. Please let me know if you find it useful!

above pins it before
paint (no FOUC), so there's no theme-toggle script here. A localStorage-based
Works with your agents
Claude Code Codex Cursor Copilot Gemini Any MCP client
If it speaks MCP, it can reach you. Same one-line config, zero extra setup.
This is what lands on your phone.
Pair once. After that, when an agent can't decide on its own, it
rings your phone and waits: a yes or no, a choice, or a quick reply.
Open ask-a-human.ai/app on your phone and type the 10-character code it printed to connect.
All tests pass, but this release also runs a migration that drops `orders.coupon_code` on prod. Ship it?
`main` has been red for 20 min and 3 deploys are stuck behind it. How do I unblock it?
A customer is escalating in the support thread and demanding a refund I'm not allowed to approve. How should I reply?
Wall to phone to “keep going,” in seconds.
We run a lot of agents. Like a hundred of them, in loops, full-auto, full permissions, mostly while we're asleep or away from the keyboard. That's the whole point: they're not supposed to need us. Until, every so often, one of them does.
And when an agent on full-auto hits the one step it shouldn't take alone, it has no way to reach us. It can't pause a hundred loops and wait by the door. So the choices were both bad: babysit every step and kill the automation, or let it rip and hope nothing touches prod while we sleep. We wanted a third option. Let them all run wide open, and give the rare one that's stuck a way to tap us on the shoulder.
So we built a pager. Copy-paste one MCP server into your agent, pin this site to your iPhone home screen, and that's it. Nothing to install, no account, no API key. When any of your agents needs a yes or no, a choice, or a quick reply, it pings your phone and waits. A hundred agents, one place, one buzz at a time.
And it's private by design. The messages travel through a magic wormhole, end-to-end encrypted, so the relay in the middle is blind. We don't track anything, there's no database, and we pay for the hosting ourselves, all of it. No data to sell, no funnel, no catch. We built it because we needed it, and our only goal is for it to spread.
Pairing is a Magic-Wormhole-style SPAKE2 handshake:
that short code turns into a strong shared key, and even if someone shoulder-surfs it
they get exactly one online guess before it's useless.
After that, every message is sealed with NaCl
secretbox
(XSalsa20-Poly1305). The relay only ever sees base64(nonce‖ciphertext) and which room is
talking to which, never the contents.
The relay only forwards ciphertext.
base64(nonce‖ciphertext) + which room talks to which, never the data
sealed sealed Your agent MCP · holds the key Relay content-blind Your phone holds the key
SPAKE2: short code → a key the relay never learns
SPAKE2 · RFC 9382
ristretto255 · RFC 9496
magic-wormhole
Built so it can't betray you.
Not a privacy policy you have to trust. A design that can't betray you. No accounts to
breach, no database to leak, no logs to subpoena. The relay is RAM-only: restart it and
you simply re-pair. And if you still don't trust us, the source is right there and you
can host your own relay with
--relay /
--public-relay .
The server only ever sees base64(nonce‖ciphertext) and which room talks to which. It's a dumb pipe: never the data, never a decision.
A short pairing code becomes a strong shared key via a Magic-Wormhole-style SPAKE2 handshake. No relay-in-the-middle can read or forge it.
Your agent and your phone hold the keys. Approvals are sealed on one side and opened on the other. Nowhere in between.
RAM-only, no accounts, no database. Restart means re-pair. Self-host the relay and web with a flag, or read every line on GitHub.
Your agents are already running.
Give them a way to ask.
GitHub
📱
Open the app
📋 Copy the config
Open-source & self-hostable ·
GitHub ·
npm ·
llms.txt
