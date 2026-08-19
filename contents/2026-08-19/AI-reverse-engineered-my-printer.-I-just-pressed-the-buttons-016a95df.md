---
source: "https://www.panaxeo.com/blog/ai-reverse-engineered-my-printer-i-just-pressed-the-buttons"
hn_url: "https://news.ycombinator.com/item?id=49359452"
title: "AI reverse-engineered my printer. I just pressed the buttons"
article_title: "AI Coding Agent Reverse-Engineers a Canon Printer — Panaxeo"
image: "http://static1.squarespace.com/static/63c7fa513f18c70a321bbf41/63c7fa6d3f18c70a321bc337/6a7b10e37470671a1036c659/1786966214830/Printer_Blog.png?format=1500w"
author: "druchem"
captured_at: "2026-08-19T10:19:08Z"
capture_tool: "hn-digest"
hn_id: 49359452
score: 1
comments: 0
posted_at: "2026-08-19T10:17:40Z"
tags:
  - hacker-news
  - translated
---

# AI reverse-engineered my printer. I just pressed the buttons

- HN: [49359452](https://news.ycombinator.com/item?id=49359452)
- Source: [www.panaxeo.com](https://www.panaxeo.com/blog/ai-reverse-engineered-my-printer-i-just-pressed-the-buttons)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T10:17:40Z

## Translation

タイトル: AI が私のプリンターをリバース エンジニアリングしました。ただボタンを押しただけです
記事のタイトル: AI コーディング エージェントが Canon プリンターをリバース エンジニアリングする — Panaxeo
説明: Panaxeo のエンジニアが AI コーディング エージェントを使用して Canon プリンターをリバース エンジニアリングしました

記事本文:
0
コンテンツにスキップ
ホーム
サービス
作品
求人
ブログ
お問い合わせ
メニューを開く
メニューを閉じる
ホーム
サービス
作品
求人
ブログ
お問い合わせ
メニューを開く
メニューを閉じる
ホーム
サービス
作品
求人
ブログ
お問い合わせ
AI が私のプリンターをリバースエンジニアリングしました。ただボタンを押しただけです。
ベンダー アプリを使用せずに、プリンターをロックダウンされたネットワークに接続します。
イーゴリ・リスカ作曲、作品5。
私は Canon PIXUS TS3300 と、それに接続するために必要なホーム ネットワークを持っています。これは 5 分の作業であり、ほとんどの人にとっては、ベンダー アプリをインストールし、ウィザードをタップして完了するだけです。
ベンダーのアプリをインストールしたくありませんでした。
部分的には、これは味です。ベンダーのアプリは価格に応じて構築される傾向があります。彼らはひどく老ける。彼らはアカウントが欲しいのです。そして、企業がハードウェアをサポートしようとする継続的な関心に依存するように、密かにハードウェアを作成します。プリンターは紙にインクを付ける箱です。関係性は必要ないはずです。
しかし、ほとんどは実用的なものでした。プリンターを接続したいネットワークはハードウェア アドレスでフィルターされているため、そこに電話を追加したくありませんでした。
これが問題です。Canon アプリは、携帯電話が現在接続しているネットワークに対してのみプリンターをセットアップできます。したがって、サポートされているパスでは、純粋にプリンタにそのことを伝えるために、意図的に小さくしておいたネットワークに携帯電話を接続し、その後電話を取り外してフィルタリングを元に戻す必要がありました。何も学習しないようにするために、満足のいくネットワークの 2 つの再構成を行いました。
アプリをインストールし、2026 年にネットワークを 2 回再構成します。私なら、同じ夜をエージェントにそのやり方を教えるのに費やして、二度と考えたくないと思います。
なぜこれが思っているよりも難しいのか
箱から出したばかりのプリンタにはネットワークがありません。そこで、アプリが実際のネットワークの名前とパスワードを引き渡すために接続する一時的なアクセス ポイントを独自に作成します。キヤノンはこれをケーブルレスと呼んでいます

設定。これは賢明な設計であり、構成交換全体が約 2 分間のみ存在するチャネル上で行われることを意味します。
メーカーがここで抱えている問題については、今後の多くのことを説明してくれるので、じっくりと考えてみる価値はあります。プリンターと電話は一度も出会ったことがない。ペアリングも認証も何もありません。他のベンダーは、ケースにパスワードを印刷するか、シリアル番号と製造日からパスワードを導き出すことでこの問題を解決しています。選択肢はそれほど多くありません。プリンターが 2 分間のウィンドウを保護するために何を使用するとしても、それはプリンター自身がユーザーに伝えることができるものでなければなりません。つまり、それは実際には秘密ではないということです。
アプリを使わずにこれを行うには、アプリが話す内容をすべて、その一時的なネットワーク上でそのウィンドウ内で話す必要がありました。
もう 1 つ制約がありました。この作業をラップトップで行っていないということです。これは、Claude Code が存在する自宅のコンテナから実行しました。独自のファイル システム、ネットワーク上の独自の場所で、私がコンピュータの前にいないときでも動作し続けます。このマシンにはワイヤレス ハードウェアがまったく搭載されていないため、プリンタの一時ネットワークに参加できません。
実際にあるのは有線ネットワーク上の古い Raspberry Pi で、Pi には無線機能が付いています。したがって、Pi はワイヤレス フロントエンドになりました。エージェントのマシンはイーサネット経由でエージェントを駆動し、Pi が実際の参加と中継を行います。
つまり、役割分担は次のようになります。エージェントは Linux ボックスを使用してプレイし、私は、何を試行してプリンターのボタンを押すかを決定します。このセットアップについては、別途記事を書く価値があります。このプリンターよりも興味深いことが判明したため、後で詳しく説明します。
アプリがプロトコルを認識している場合、アプリにはプロトコルが含まれています。そこで私たちは Android パッケージを分解しました。
わかったこと: セットアップ チャネルは SNMPv3 (認証と暗号化を備えた SNMP のバージョン) です。すべては生きている

SNMP ツリーのキヤノン独自のブランチの下にあり、それらを一緒に見ると、名前は厳密には不可解ではありません
OID = "1.3.6.1.4.1.1602.1.3" # 1602 はキヤノンの企業番号です
O_MODE = OID + ".2.100.2.0" # ワイヤレス動作モード
O_SSID = OID + ".2.100.10.3.0" # ネットワーク名
O_AUTH = OID + ".2.100.10.6.0" # 認証タイプ
O_ENC = OID + ".2.100.10.7.0" # 暗号化タイプ
O_WPAPASS = OID + ".2.100.10.110.0" # パスフレーズ
O_ENABLE = OID + ".3.3.1.100.10.1.3.3" # プロファイル有効化フラグ
O_P5 = OID + ".2.100.10.5.0" # WPA2 プロファイルが 6 に設定したいリンク フィールド
ユーザー名は固定です。また、資格情報は共有秘密ではなく、プリンターがそれ自体について公開する、エンジン ID と呼ばれる識別子から派生します。プリンターを見ることができる人なら誰でも計算できます。
導出は引用するには十分短いです。プリンターのエンジン ID は公開識別子であり、そこから認証キーが構築されますが、秘密は一切含まれません。
def auth_key_from_engineid (engineid_bytes):
h1 = エンジンID_バイト 。 16 進数 ()
a = h1 [ 10 :] # 固定プレフィックスを削除し、残るのはハードウェア アドレスです
return create_v3_password_hash ( a , a , a )
はい、同じ値が 3 回あります。この関数は 3 つの個別の入力を受け取り、アプリは 3 つのスロットすべてに 1 つの値を供給します。これは、この関数が汎用的に記述され、その後は一方向のみに使用されたことを示唆しています。それ以上に適切な説明はありませんが、重要なことは同じことをすることなので、それは問題ではありません。
Canon は標準的な方法でキーを導出していないため、ここで必要となるのは実際に作業を行った部分です。
標準的なものは RFC 3414 で公開されており、短いものです。パスワードを正確にメガバイトまで拡張してハッシュし、その結果をデバイスのエンジン ID とともにハッシュして、キーがそのデバイスでのみ機能するようにします。
デフォルトのパスワード_to

_key (パスワード、エンジン ID ):
buf = (パスワード * ( 1048576 // len (パスワード) + 1 ))[: 1048576 ]
ダイジェスト = ハッシュリブ 。 sha1 ( buf )。ダイジェスト()
ハッシュライブラリを返します。 sha1 (ダイジェスト + エンジン ID + ダイジェスト)。ダイジェスト()
ハッシュが 2 つ。メガバイト単位の繰り返しにより、パスワードのブルート フォース攻撃が高価になり、世界中のすべての SNMPv3 デバイスが同じことを行います。
Canon は最初に独自の導出を実行し、その後でその結果に標準のローカリゼーションを適用します。その最初のステップを回復するには、共有ライブラリを解体する必要がありました。概要:
def 導出 ( a , b ):
ブロック = シャッフル ( a , permutation_table ) + b
5回繰り返します：
Salt = Salt_table [ブロックから取得した 1 バイト] # データは独自のソルトを選択します
ブロック = sha256 (ブロック + ソルト)
block = shuffle ( block , next permutation_table )
リターンブロック
どれも暗号学的に興味深いものではありません。重要な詳細は、データが独自のソルトを選択するということです。つまり、そこに至るまでの方法を外部から推測することはできず、すべてのステップが正確に正しくなければ、ダイジェストに関する苦情以外は何も返されません。
その瞬間、それは解決したように見えました。私たちはプロトコル、ユーザー、そして鍵を導出する方法を持っていました。
最初の壁は接合部ですらなかった。認証でした。
SNMPv3 はパスワードを送信しません。パスワードとプリンターのエンジン ID から導出されたキーを使用して、メッセージに対して計算されたダイジェストを、特定の順序で、特定のラウンド数で送信します。一歩でも間違えると、プリンターは「ダイジェストが間違っています」という意味のエラーを返すだけで、それ以外は何も返しません。どの部分が間違っているのかはわかりません。保護しているものそのものを漏洩することなしには不可能です。
私たちはそこで多くの時間を過ごしました。賢い暗号化ではなく、特定の方言、つまりどのバイトがハッシュに入るのか、どの順序でキーがその特定のデバイスにどのように関連付けられるのかについてです。

。キヤノン独自の派生とその上の標準ローカリゼーションという 2 つのレイヤーを同時に正しくする必要がありました。
rawkey = canon_kdf . auth_key_from_engineid ( Engineid ) # Canon のカスタム KDF
lk = canon_kdf . password_to_key ( rawkey , Engineid ) # 次に RFC 3414 ローカリゼーション
aeskey = lk [: 16 ] # 最初の 16 バイトを暗号化
iv = boots . to_bytes ( 4 , "big" ) + etime 。 to_bytes ( 4 , "big" ) + ソルト
mac = hmac . new ( lk 、 msg ( b"\x00" * 12 )、 hashlib . sha1 )。ダイジェスト()[: 12 ]
どの試みも同じ役に立たない答えを生み出しました。これはプロジェクトの中で最も魅力的ではなく、最も時間がかかる段階でした。
最終的にはダイジェストが一致しました。プリンターは私たちのメッセージを受け入れ、構成を取得し、それを使用しようとしましたが、静かに失敗しました。
それが失敗したことを知ることは、それ自体の小さな儀式でした。プリンターはネットワーク経由でエラーを報告しません。何が起こったのかを知るには、オンデマンドで生成されるページであるネットワーク レポートを印刷させ、下部にあるコードを読み取ります。 Ours said C-5 .
C-5 は、プリンターがネットワークに参加できなかったことを意味します。その理由は述べられていない。 「間違ったパスワード」と「そのネットワークが見えない」または「要求が理解できませんでした」は区別されません。私たちはそのページを 1 回だけ印刷しました。これで、何が失敗したかがわかり、その主題に十分な紙が費やされたことがわかりました。その後、テストはより単純で悲しいものになりました。ツールを実行し、プリンターがネットワーク上に表示されたかどうかを確認するというものです。 It had not.
原因として考えられることを5つ考えてみました。
ネットワーク名は非表示になります。 Plausible.リッスンしてネットワークをスキャンするデバイスは、自身をアナウンスしないデバイスを見つけることはできません。私たちは放送をオンにしました。まだC-5。
ハードウェアアドレスフィルタリング。これももっともらしいですが、これには問題がありました。プリンタには 2 つの異なるハードウェア アドレスがあり、1 つは一時用です。

ary ネットワークとクライアントとして使用するネットワーク。どちらを許可するかを検討しました。まだC-5。
WPA2とWPA3が混在しています。安価な組み込み無線機は、両方を提供するネットワークに対応できないことがよくあります。 WPA2のみにしました。まだC-5。
保護された管理フレーム。同じ系統の問題です。オフになりました。まだC-5。
パスフレーズのエンコーディング。おそらく、パスワードを間違った形式で送信したか、予想とは異なる方法でエスケープまたはエンコードされていた可能性があります。バリエーションを試してみました。まだC-5。
どれも合理的な理論でした。私たちはそれらをすべてテストしましたが、それらはすべて間違っていました。それぞれ、再構成、プリンターまで歩いて行き、それが表示されたかどうかを確認するために再度確認する必要がありました。そうではありませんでした。
理論 4 のあたりで、なぜ人々はこの種の仕事を専門家に任せるのかを思い出しました。 Panaxeo には非常に優れた製品がいくつかあります。たぶん、そのうちの一人に電話したほうがよかったかもしれない。
ある時点で、パターンはメッセージになります。 5 つの賢明な仮説はすべて除外されましたが、6 つ目の仮説を示唆する新しいものは何もありませんでした。
静的解析により、プログラムで何ができるかが分かります。実際のデバイスに対して実際の値を使用して実行時に実際に何を行うかはわかりません。そのためには、動的な分析が必要です。つまり、動いているものが動作するのを観察します。
そこでアプローチを変更しました。電話にアプリをインストールし、実際にセットアップを実行させ、会話を録音します。
これは少し気が進まなかった措置ではありますが、原則的な理由からではないことを明確にしておきたいと思います。パズルを諦めたような気分だった。この目的のために使い捨てネットワークをセットアップし、アプリを実行して成功させました。
これが捕獲が可能であった理由であり、プロジェクト全体を実現可能にしたのと同じ事実です。
プリンターの一時セットアップ ネットワークが開いています。ワイヤレス暗号化はまったくありません。したがって、Pi の無線を右チャンネルのモニター モードにすると、すべてのフレームが表示されます。
したがって、キャプチャ自体は目立ったものではありません。
NMC

li device set wlan0 管理 no # NetworkManager による無線への接続を停止します
iw dev wlan0 セットタイプモニター
iw dev wlan0 チャンネル 6 を設定します
tcpdump -i wlan0 -w Capture.pcap # その後、アプリに動作させます
最上位のアプリケーション層は暗号化されていますが、そのキーはプリンターのエンジン ID から取得され、プリンターは要求した人にそれを渡します。アプリの読み取りからの導出はすでに得られているため、電話が使用していたのと同じキーを計算して交換を読み取ることができます。
eng = Discover_engineid ( target ) # プリンターが通知するだけ
lk = canon_kdf 。 password_to_key ( canon_kdf . auth_key_from_engineid ( eng )、 eng )
plain = Cipher (アルゴリズム . AES ( lk [: 16 ])、モード . CFB ( iv ))。復号化子 ()。更新 (暗号文)
無線層でオープンされ、アプリケーション層で派生可能です。その捕獲については何も秘密を必要としませんでした。
そして録音はそのシンプルさに驚くべきものだった。結合全体、つまりその夜ずっと食べていたものは、1 つのメッセージです。単一の SNMP 書き込みには、モード、イネーブル フラグ、ネットワーク名、認証タイプ、暗号化タイプ、パスフレーズ、およびもう 1 つのカウンタの 7 つの値が含まれます。
結合全体を書き出すと次のようになります。
vbs = [
( O_MODE , tlv ( 0x04 , JOIN_MODE .to_bytes (

[切り捨てられた]

## Original Extract

A Panaxeo engineer used an AI coding agent to reverse-engineer a Canon printer

0
Skip to Content
Home
Services
Works
Jobs
Blog
Contact
Open Menu
Close Menu
Home
Services
Works
Jobs
Blog
Contact
Open Menu
Close Menu
Home
Services
Works
Jobs
Blog
Contact
AI reverse-engineered my printer. I just pressed the buttons.
Getting a printer onto a locked-down network without the vendor app.
Written by Igor Liska and Opus 5.
I have a Canon PIXMA TS3300 and a home network it needed to join. That should be a five-minute job, and for most people it is: install the vendor app, tap through a wizard, done.
I did not want to install the vendor app.
Partly, this is taste. Vendor apps tend to be built to a price; they age badly; they want an account; and they quietly make your hardware depend on a company's continued interest in supporting it. A printer is a box that puts ink on paper. It should not need a relationship.
But mostly it was practical. The network I wanted the printer on is filtered by hardware address, and I did not want to add my phone to it.
Here’s the blocker: the Canon app can only set up the printer for the network the phone is currently joined to . So the supported path required letting my phone onto a network I had deliberately kept small, purely so it could tell the printer about it, then taking the phone off and putting the filtering back the way it was. Two reconfigurations of a network I was happy with, in order to avoid learning anything.
Install an app and reconfigure my network twice, in 2026. I would rather spend the same evening teaching an agent to do it, and never think about it again.
Why this is harder than it sounds
A printer fresh out of the box has no network. So it makes its own: a temporary access point that the app connects to in order to hand over the real network's name and password. Canon calls this cableless setup. It is a sensible design, and it means the entire configuration exchange happens over a channel that only exists for about two minutes.
It is worth sitting with the problem the manufacturer has here, because it explains a lot of what follows. The printer and the phone have never met. There is no pairing, nothing to authenticate with. Other vendors solve it by printing a password on the case, or deriving one from the serial number and the manufacturing date. There are not many options. Whatever the printer uses to protect that two-minute window, it has to be something it can tell you itself, which means it is not really a secret.
To do that without the app, I had to speak whatever the app speaks, over that temporary network, within that window.
There was one more constraint: I did not do this work on my laptop. I did it from a container at home that Claude Code lives in: its own filesystem, its own place on the network, and it keeps working when I am not at the computer. That machine has no wireless hardware at all, so it cannot join the printer's temporary network.
What it does have is an old Raspberry Pi on the wired network, and the Pi has a radio. So the Pi became the wireless front end. The agent's machine drives it over Ethernet, and the Pi does the actual joining and relaying.
So the division of labor is: the agent gets a Linux box to play in, and I decide what it is allowed to try and press the buttons on the printer. That setup deserves its own write-up, which I will get to, because it turned out to be more interesting than this printer.
If the app knows the protocol, the app contains the protocol. So we pulled apart the Android package.
What we found: the setup channel is SNMPv3, the version of SNMP that has authentication and encryption. Everything lives under Canon's own branch of the SNMP tree, and the names are not exactly cryptic once you see them together
OID = "1.3.6.1.4.1.1602.1.3" # 1602 is Canon's enterprise number
O_MODE = OID + ".2.100.2.0" # wireless operating mode
O_SSID = OID + ".2.100.10.3.0" # network name
O_AUTH = OID + ".2.100.10.6.0" # authentication type
O_ENC = OID + ".2.100.10.7.0" # encryption type
O_WPAPASS = OID + ".2.100.10.110.0" # passphrase
O_ENABLE = OID + ".3.3.1.100.10.1.3.3" # profile enable flag
O_P5 = OID + ".2.100.10.5.0" # a link field the WPA2 profile wants set to 6
The username is fixed. And the credentials are not a shared secret at all, they are derived from an identifier the printer publishes about itself, called the engine ID. Anyone who can see the printer can compute them.
The derivation is short enough to quote. The printer's engine ID is a public identifier, and the authentication key is built out of it with no secret involved anywhere:
def auth_key_from_engineid ( engineid_bytes ):
h1 = engineid_bytes . hex ()
a = h1 [ 10 :] # drop the fixed prefix, what remains is the hardware address
return create_v3_password_hash ( a , a , a )
Yes, the same value three times. The function takes three separate inputs, and the app feeds it one value in all three slots, which suggests it was written to be general and then only ever used one way. I have no better explanation than that, and it does not matter, because what matters is doing the same thing.
What that calls into is the part that actually took the work, because Canon does not derive the key the standard way.
The standard one is published in RFC 3414 and it is short. Expand the password to exactly a megabyte, hash it, then hash the result together with the device's engine ID so the key only works on that device:
def password_to_key ( password , engine_id ):
buf = ( password * ( 1048576 // len ( password ) + 1 ))[: 1048576 ]
digest = hashlib . sha1 ( buf ). digest ()
return hashlib . sha1 ( digest + engine_id + digest ). digest ()
Two hashes. The megabyte of repetition is there to make brute forcing a password expensive, and every SNMPv3 device in the world does the same thing.
Canon runs its own derivation first, and only then applies the standard localization to the result. Recovering that first step meant pulling a shared library apart. In outline:
def derive ( a , b ):
block = shuffle ( a , permutation_table ) + b
repeat 5 times :
salt = salt_table [one byte taken from block] # the data picks its own salt
block = sha256 ( block + salt )
block = shuffle ( block , next permutation_table )
return block
None of it is cryptographically interesting. The detail that matters is that the data selects its own salt, which means you cannot guess your way to it from the outside, and every step has to be exactly right or you get nothing back but a complaint about your digest.
At that moment, it looked solved. We had the protocol, the user, and a way to derive the keys.
The first wall was not even the join. It was authentication.
SNMPv3 does not send a password. It sends a digest computed over the message using a key derived from the password and the printer's engine ID, in a specific order, with a specific number of rounds. Get any step wrong and the printer replies with an error meaning "your digest is wrong" and nothing else. It does not tell you which part is wrong. It cannot, without leaking the very thing it is protecting.
We spent a lot of time there. Not on clever cryptography, but on the specific dialect: which bytes go into the hash, in which order, how the key is tied to that particular device. Two layers had to be right at once: Canon's own derivation and then the standard localization on top of it:
rawkey = canon_kdf . auth_key_from_engineid ( engineid ) # Canon's custom KDF
lk = canon_kdf . password_to_key ( rawkey , engineid ) # then RFC 3414 localisation
aeskey = lk [: 16 ] # first 16 bytes encrypt
iv = boots . to_bytes ( 4 , "big" ) + etime . to_bytes ( 4 , "big" ) + salt
mac = hmac . new ( lk , msg ( b"\x00" * 12 ), hashlib . sha1 ). digest ()[: 12 ]
Every attempt produced the same unhelpful answer. This was the least glamorous phase of the project and easily the most time-consuming.
Eventually the digest matched. The printer accepted our message, took the configuration, tried to use it, and quietly failed.
Finding out that it had failed was its own small ritual. The printer does not report errors over the network. To learn what happened, you make it print its network report, a page it produces on demand, and read the code at the bottom. Ours said C-5 .
C-5 means the printer could not join the network. It does not say why. It does not distinguish "wrong password" from "cannot see that network" or "I did not understand your request". We printed that page exactly once, which was enough to know what we were failing at and enough paper spent on the subject. After that, the test became simpler and sadder: run the tool, then look to see whether the printer had appeared on the network. It had not.
I thought of five things that could have caused it.
The network name is hidden. Plausible. A device that scans for networks by listening will never find one that does not announce itself. We turned the broadcast on. Still C-5.
Hardware address filtering. Also plausible, and this one had a wrinkle: the printer has two different hardware addresses, one for its temporary network and one it uses as a client. We worked out which one to allow. Still C-5.
Mixed WPA2 and WPA3. Cheap embedded radios often cannot cope with a network offering both. We made it WPA2 only. Still C-5.
Protected management frames. Same family of problems. Turned off. Still C-5.
Passphrase encoding. Maybe we were sending the password in the wrong form, escaped or encoded differently than expected. We tried the variants. Still C-5.
Each of these was a reasonable theory. We tested them all, and they were all wrong. Each cost a reconfiguration, a walk to the printer, and another look to see whether it had appeared. It had not.
Somewhere around theory four, I remembered why people leave this sort of work to professionals. We have some very good ones at Panaxeo. Maybe I should have called one of them.
At some point, the pattern becomes the message. Five sensible hypotheses, all eliminated, and nothing new arriving to suggest a sixth.
Static analysis tells you what a program can do. It does not tell you what it actually does at runtime, with real values, against a real device. For that, you need dynamic analysis: watch the working thing work.
So we changed the approach. Install the app on a phone, let it do the setup for real, and record the conversation.
I want to be clear that this was a slightly reluctant step, but not for any principled reason. It felt like giving up on the puzzle. I set up a throwaway network for the purpose, ran the app, and let it succeed.
Here is why the capture was even possible, and it is the same fact that made the whole project feasible.
The printer's temporary setup network is open . No wireless encryption at all. So with the Pi's radio in monitor mode on the right channel, every frame is visible.
So the capture itself is unremarkable:
nmcli device set wlan0 managed no # stop NetworkManager touching the radio
iw dev wlan0 set type monitor
iw dev wlan0 set channel 6
tcpdump -i wlan0 -w capture.pcap # then let the app do its thing
The application layer on top is encrypted, but its keys come from the printer's engine ID, which the printer hands out to anyone who asks. We already had the derivation from reading the app, so we could compute the same keys the phone was using and read the exchange:
eng = discover_engineid ( target ) # the printer just tells you
lk = canon_kdf . password_to_key ( canon_kdf . auth_key_from_engineid ( eng ), eng )
plain = Cipher ( algorithms . AES ( lk [: 16 ]), modes . CFB ( iv )). decryptor (). update ( ciphertext )
Open at the radio layer, and derivable at the application layer. Nothing about that capture required a secret.
And the recording was startling in its simplicity. The entire join, the thing that had eaten the whole evening, is one message . A single SNMP write containing seven values: mode, an enable flag, the network name, an authentication type, an encryption type, the passphrase, and one more counter.
Written out, the whole join is this:
vbs = [
( O_MODE , tlv ( 0x04 , JOIN_MODE . to_bytes (

[truncated]
