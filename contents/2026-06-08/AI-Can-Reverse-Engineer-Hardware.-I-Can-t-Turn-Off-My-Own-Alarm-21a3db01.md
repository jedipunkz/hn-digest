---
source: "https://www.sudomoin.com/p/ai-can-reverse-engineer-hardware-i-can-t-turn-off-my-own-alarm"
hn_url: "https://news.ycombinator.com/item?id=48445921"
title: "AI Can Reverse-Engineer Hardware. I Can't Turn Off My Own Alarm"
article_title: "AI Can Reverse-Engineer Hardware. I Can't Turn Off My Own Alarm."
author: "dimitri-vs"
captured_at: "2026-06-08T16:30:51Z"
capture_tool: "hn-digest"
hn_id: 48445921
score: 2
comments: 0
posted_at: "2026-06-08T14:30:43Z"
tags:
  - hacker-news
  - translated
---

# AI Can Reverse-Engineer Hardware. I Can't Turn Off My Own Alarm

- HN: [48445921](https://news.ycombinator.com/item?id=48445921)
- Source: [www.sudomoin.com](https://www.sudomoin.com/p/ai-can-reverse-engineer-hardware-i-can-t-turn-off-my-own-alarm)
- Score: 2
- Comments: 0
- Posted: 2026-06-08T14:30:43Z

## Translation

タイトル: AI はハードウェアをリバース エンジニアリングできる。自分のアラームを止めることができない
記事のタイトル: AI はハードウェアをリバース エンジニアリングできる。自分のアラームを止めることができません。
説明: クロード コード、ESP32、および独自のショップ デジタル ツールの構築について

記事本文:
AIはハードウェアをリバースエンジニアリングできる。自分のアラームを止めることができません。ディミトリ・スドモイン 検索 ログイン サインアップ ディミトリ・スドモイン ホーム
AIはハードウェアをリバースエンジニアリングできる。自分のアラームを止めることができません。
AIはハードウェアをリバースエンジニアリングできる。自分のアラームを止めることができません。
Claude Code、ESP32、および独自のショップデジタルツールの構築について
この記事を書いている間にスタンディングデスクが勝手に上がりました。あまりにも長い間座っていました。
ホームオートメーションは当初、ライフスタイルではなく、ナッジであると考えられていました。私が 1 時間座っていると、ホーム アシスタントは丁寧な口頭での警告から始まり、プッシュ通知やますます攻撃的な LED によってエスカレートし、80 分経過すると感情操作 (「子供のことを考えて…」) を開始します。私のADHDには当てはまりません。 2 週間後、私は Nest Mini の最初の文が終わる前に反射的に叩いていました。
タイピング中に机が物理的に上昇することを無視するのは困難です。
Claude Code はその最初のバージョンのほとんどを書きましたが、AI 支援が今では目立たないと感じられるようになりました。私が望むものを漠然と説明します。テンプレートを変更し、自動化をリロードし、完了します。しかしデスクは命令に従って動くことを拒否した。独自のプロトコルを備えたハードウェアは、Claude Code が単に構成ファイルを作成して実行できるものではありません。
FlexiSpot E7 Pro Plus スタンディング デスクは、どこにでもあるにもかかわらず、自動化を望んでいません。既存のコミュニティ統合では、ホーム アシスタントはその高さを読み取ることができますが、それを移動しようとするほとんどの試みは静かに終了します。コントローラーは接続されたキーパッドとの特定の会話を期待しており、既存の統合はすべてそのハンドシェイクをスキップしているようでした。そこで、7 ドルの ESP32 にキーパッドを信頼できる程度に偽装させました。部品代は約 12 ドル、午後はクロード コードをひたすら勉強し、デスクはそれに従った。
私'

AI を使用して Home Assistant 設定ファイルを何か月も作成してきました。しかし、これは、ハードウェア プロトコルをリバース エンジニアリングし、マイクロコントローラー用に機能する組み込みファームウェアを作成するのに役立つ LLM でした。コーディング エージェントが端末から物理世界に流出しています。そして、それをどこまで押し進められるか試してみたかったのです。
デスクはまさに​​最新のものでした。その前に、Claude Code が私のホーム アシスタントのセットアップを段階的に引き継いでいました。プレゼンスセンサーとウォーキングパッドパワーモニターにより、私が座っているのか、立っているのか、歩いているのかを知ることができます。フリゲートカメラは物体検出と人物認識と統合されています。漏れセンサー、セキュリティ自動化、休憩エスカレーション システム全体。
ある時点で、ホーム アシスタントがどのように機能するのか実際にはわかっていないことに気づきました。あまり。私はオートメーションを最初から手動で作成したことがありません。ダッシュボードは使いません。 Claude Code は、私の家の物理システムへのインターフェースです。それが今の状況です。
ある晩、我が家の洗濯機に軽い水漏れが起きたとき、このことがはっきりと分かりました。漏れセンサーが作動しました。警報が鳴りました。そして、私はクロード コードを開いてアラームを止めるように要求しました。自分でそれを行う方法が本当にわからなかったからです。
これは私が陥った罠と読むこともできますし、単なる次の抽象化層と読むこともできます。両方の読みはおそらく同時に真実です。私はほとんどの場合、「単なる別のツール」のフレーミングに傾いています。
デスクの後、私は貪欲になりました - クロードはデスクを制御する方法を理解できます - それは他に何を理解できるでしょうか？ディスプレイが壊れた Plaud Note はどうなりますか?これはポータブルな「AI」ボイスレコーダーであり、「スマート」である理由は特にありませんでした。音声を録音するマイクです。スマートな機能は文字通り、転写と要約という 2 つの API 呼び出しだけであり、どちらも平凡です。バカにしてもっとCAを増やしたい

pable モデルはスマート パーツを処理します。
そこで、Claude Code はそれをリバース エンジニアリングしました - 驚くほど簡単に。標準的なリバース エンジニアリング アプローチ: アプリを分析し、プロトコルを計画し、デバイスが電話とどのように通信するかを解明します。約 5 日間のカジュアルなセッションを繰り返し、何度も行き止まりになりましたが、ある時点で、数日前にすでに取得したデータの中に必要なキーが存在することがわかりました。ただ、十分に詳しく調べていなかっただけです。アプリやサブスクリプションを必要とせずに、Bluetooth 経由で直接録音を取り込むことができるようになりました。
さらに 3 つの ESP32 ボードを注文中です。次のターゲットは API を持たないロボット芝刈り機です。
私がいつも思い出す木工の例えがあります。ほとんどの木工職人は、クロスカットスレッド (テーブルソーのレールに乗って木材を切断するスレッド) を購入するのではなく、自分で組み立てます。ほぼ通行権です。それは安いからではなく、既製のスレッドが必然的に汎用的であるためです。あなたの鋸、材料のサイズ、実際の作業方法などは知りません。自分が構築していないものを簡単に変更することはできません。私はいつもこの人でした。クロード コードが登場するずっと前に、Windows に組み込まれている音声文字起こしアプリでは十分ではなかったので、私は独自の音声文字起こしアプリを作成しました。私は、入力する代わりに Caps Lock を押して話します。現在、コンピューターでの操作の 90% は音声です。おそらくもっと良いアプリはあると思いますが、私のアプリはまさに私が望むように機能します。
Claude Code により、これらのものをより速く構築できるようになりました。障壁は非常に低くなり、実験はほとんど無料になります。愛犬が裏口から家に戻ってくるたびに鳴るように、昔ながらのセブンイレブンのドアチャイムを設置しました。それを約2日間保管しましたが、おそらく妻が望んでいたよりも2日長かったでしょう。オートメーションを削除しました。総投資額: 5 分程度。
私は趣味を見つけたり、人生の約 80% を放棄したりすることで非常に有名です。

マスターへの道 - 完全に諦めるのではなく、ただ次のことに進むだけです。サンクコストが最小限に抑えられているため、問題なく立ち去ることができます。そして、反復が十分に速いということは、脳が次に進む前に実際に終了する可能性があることを意味します。それがクロード・コードが変えたものです。面倒な部分、適切な構文、デバッグ、構成の調査を処理してくれるので、私はそれが自分の生活にとって意味があるかどうかに集中することができます。また、プロジェクトのバックアップを選択するのも簡単です。文字通りセッションを再開するだけです。すべてのファイル、ドキュメント、コンテキストがすぐそこにあります。 「どこにいたの？」と尋ねることができます。または「次は何ですか？」すぐに答えが得られます。最もコストがかかるコンテキストの切り替えは基本的に排除されます。
デスク、Plaud、ホーム オートメーション、音声アプリ。それらはすべてお店の道具です。実際の作業方法に合わせて構築されました。維持するのは私のものです。私のものはデバッグ用です。壊れるのは私です。
これらのモデルの精度は約 90% です。障害モードのいくつかは、あなたが認識できるようになります。知識の遮断。時間の見積もりは毎回大きく外れており、計画にはまったく役に立ちません。セッションの前半で話し合った内容は、その後のすべてに大きなバイアスを与えます。疲れている、または遅くなりそうだと言うと、モデルは残りの時間を急いで終わらせようとします。これらは、存在することがわかっていれば回避できるパターンです。
しかし、失敗には一貫性がありません。明らかに間違っている箇所のスクリーンショットを見て、「素晴らしいですね、これで終わりだと思います」と言うのを私は見てきました。あのウエストワールド路線。 「私には何にも見えません。」誰でもすぐに問題に気づくでしょう。あなたがそれを指さすまで、モデルは実際にはそうではありません。
本当の危険は、発見できるようになった障害モードではありません。それはあなたにはできないことです、なぜならあなたは自分の深みから外れているからです。 90% がいつドリフトしたかを認識できるほど十分なドメイン知識がない場合は、

間違った 10% を選択すると、問題を指摘することさえできなくなります。不正確さは静かに増幅するだけで、道に迷ったという感覚も得られずに、現実から遠く離れたどこかに行き着いてしまいます。あなたは、リレーを主電源に初めて配線するときに、「素晴らしいですね、これで終わりだと思います」と言います。
ハードウェアに補助金を出したり、データをロックアウトしたり、アクセスに対して月額料金を請求したり、価格を値上げしたりするなど、あらゆる企業が同じ戦略にますます集中しているように感じられます。どうする、切り替え？彼らは皆それをやっています。カスタマー サポートは現在、どこにも行かない Google フォームになっています - あなたを人類的に見ています。どうやら、それが今の状況のようです。
スタンディングデスクの代金を支払いました。ハードウェアは私のものです。メーカーが統合を構築していない場合は、私が統合を追加します。 Plaud レコーダーの代金を支払いました。私の音声録音は、誰かのクラウドを経由する必要はありません。このデバイスはダムレコーダーとしての機能を十分に備えており、それが今の状態です。私が root 化して Bluetooth ブリッジとして使用するために 80 ドルで購入した OnePlus 携帯電話は、どう見ても時代遅れです。電話として使うには遅すぎます。しかし今、それは第二の人生を持っています - クロードコードがデバイスをリバースエンジニアリングするためのブリッジです。
これを数回だけ実行すると、何ができるかに基づいてハードウェアを選択するのをやめ、代わりに何ができるかを考えることができます。今ロボット芝刈り機を購入中です。私の庭に最適なものは、ホームアシスタントのサポートがまったくありません。とにかく買って、自分の意志で曲げてみるつもりです。私のOXOコーヒーメーカーを起動するには、物理​​的なボタンを2回押す必要があります。以前は、それで会話は終わりでした。今それを見て、ESP32、リレー、30 分の作業だと思います。
以前に永続的に受け入れた制限は永続的ではなくなります。それが実際のシフトです。
私はマイクロコントローラーに関してある程度の経験があります。 Arduino 数年前、カップル

e PCB、電子機器に慣れています。そして、私は文字通りすべてをクロード コードを通じて行うことに全力を尽くしており、今ではチームの誰もが私に指導を求めに来ています。しかし、私の話は、誰でも AI を使ったハードウェアを実現できるという証拠ではありません。まだ。
私は、これがどこまで進むのか、コーディング エージェントが物理世界にどこまで到達できるのか、途中で何が壊れるのかを確認するための実験を実行している人です。すでに多くの人が Claude Code を使用して、ホーム アシスタントのセットアップを構築および改善しています。それが始まりです。私が言いたいのは、そこで止まる必要はないということです。
障壁はどんどん下がっていきます。 M5stack は、レゴのようにカチッとはまるモジュール式ハードウェアを作成します。多くのプロジェクトにおける「ハードウェア作業」の範囲は、文字通り、いくつかのものをケーブルで接続することです。制限は主にあなたの頭 (そして財布) の中にあります。そしておそらく、これは私たちがあらゆるものと対話する新しい方法なのかもしれません。ツールが利用可能であり、機能を維持し、いつかサービス利用規約に違反したと判断されない限り。
クロード・コードは壁に囲まれた庭園を突破することができます。
彼らがクロード自体を壁に囲まれた庭園にしようとしていないと仮定すると。ほぼ確実に、彼らはそれを試みるだろう。
AI をデモだけでなく実際のビジネスでも活用できるようにします。

## Original Extract

On Claude Code, ESP32s, and building your own shop digital tools

AI Can Reverse-Engineer Hardware. I Can't Turn Off My Own Alarm. Dimitri Sudomoin Search Login Sign Up Dimitri Sudomoin Home
AI Can Reverse-Engineer Hardware. I Can't Turn Off My Own Alarm.
AI Can Reverse-Engineer Hardware. I Can't Turn Off My Own Alarm.
On Claude Code, ESP32s, and building your own shop digital tools
My standing desk raised itself while I was writing this article. I'd been sitting for too long.
The home automation was originally supposed to be a nudge, not a lifestyle. If I sit for an hour, Home Assistant starts with a polite spoken warning, escalates through push notifications and increasingly aggressive LEDs, and at the 80-minute mark, starts the emotional manipulation ("Think of your kids..."). No match for my ADHD. After two weeks I was reflexively smacking my Nest Mini before it finished its first sentence.
A desk physically rising while you're typing is harder to dismiss.
Claude Code wrote most of that first version, which felt unremarkable in the way AI assistance now feels unremarkable: I vaguely describe what I want - it changes the template, reloads the automation, done. But the desk refused to move on command. Hardware with a proprietary protocol isn't something Claude Code can just config-file its way through.
The FlexiSpot E7 Pro Plus standing desk, despite being everywhere, does not really want to be automated. With the existing community integrations, Home Assistant can read its height, but most attempts to move it die silently. The controller expects a specific conversation with a connected keypad, and every existing integration was seemingly skipping that handshake. So I had a $7 ESP32 impersonate the keypad well enough to be trusted. About $12 in parts, an afternoon with Claude Code grinding away, and the desk obeyed.
I'd been using AI to write Home Assistant config files for months. But this was an LLM helping me reverse-engineer a hardware protocol and write working embedded firmware for a microcontroller. The coding agents are bleeding out of the terminal and into the physical world. And I wanted to see how far I could push that.
The desk was just the latest thing. Before that, Claude Code had already taken over my Home Assistant setup piece by piece. A presence sensor plus walking pad power monitor that knows whether I'm sitting, standing, or walking. Frigate camera integrations with object detection and person recognition. Leak sensors, security automations, the whole break escalation system.
At some point I realized I don't actually know how Home Assistant works. Not really. I've never manually written an automation from scratch. I don't use the dashboard. Claude Code is my interface to the physical systems in my house - that's just how it is now.
This became extremely clear one evening when our washing machine had a minor leak. The leak sensor triggered. The alarm went off. And I opened Claude Code to ask it to silence the alarm, because I genuinely did not know how to do it myself.
You could read that as a trap I walked into, or you could read it as just the next abstraction layer. Both readings are probably true at the same time. I lean toward the "just another tool" framing most days.
After the desk, I got greedy - Claude can figure out how to control a desk - what else can it figure out? What about my Plaud Note with a broken display? It's a portable "AI" voice recorder that had no real reason to be "smart". It's a microphone that records audio. The smart features are literally just two API calls: transcribe and summarize, both mediocre. I want it to be dumb and have more capable models handle the smart parts.
So Claude Code reverse engineered it - with surprising ease. Standard reverse engineering approach: analyze the app, map out the protocol, figure out how the device talks to the phone. Took about five days of casual sessions, multiple dead ends, and at one point it found the necessary key sitting in data it already captured days earlier - just didn't look at it closely enough. Now I can pull my recordings directly over Bluetooth without any app or subscription.
Three more ESP32 boards are on order. Next target is a robot lawnmower with no API.
There's a woodworking analogy I keep coming back to. Most woodworkers build their own crosscut sleds (a sled that rides the rails of a table saw to cut wood across) instead of buying one. It's almost a right of passage. Not because it's cheaper, but because a pre-made sled is necessarily generic. It doesn't know your saw, your material sizes, or how you actually work. You can't easily modify something you didn't build. I've always been this person. Long before Claude Code, I built my own voice transcription app because the built-in Windows one wasn't good enough. I press Caps Lock and talk instead of type. 90% of my computer interaction is voice now. There are probably better apps out there, but mine works exactly how I want it to.
Claude Code just made it possible to build these things faster. The barrier drops so low that experimentation becomes almost free. I rigged a classic 7-Eleven door chime to play whenever my dogs came back inside through the back door. Kept it for about two days, which was probably two days longer than my wife would have liked. Deleted the automation. Total investment: maybe five minutes.
I am extremely known for picking up hobbies and abandoning them about 80% of the way to mastery - not giving up exactly, just moving on to the next thing. Minimal sunk cost means walking away is painless. And fast enough iteration means you might actually finish before your brain moves on. That's what Claude Code changed. It handles the tedious parts, the proper syntax and debugging and configuration research, leaving me to focus on whether the thing even makes sense in my life. And picking a project back up is effortless. I literally just resume the session. All the files, docs and context is right there. I can ask "where were we?" or "what's next?" and get an immediate answer. Context switching, which is the thing that costs me the most, is basically eliminated.
The desk, the Plaud, the home automations, the voice app. They're all shop tools. Built to fit how I actually work. Mine to maintain. Mine to debug. Mine to break.
These models are about 90% accurate. Some of the failure modes you learn to recognize. Knowledge cutoffs. Time estimates are way off, every time, completely useless for planning. Whatever you discuss in the first half of a session strongly biases everything after. If you mention you're tired or it's getting late, the model will rush through the rest trying to wrap things up. These are patterns you can work around once you know they exist.
But failures are inconsistent. I've watched it look at a screenshot of something clearly wrong and say "looks great, I think we're done here." That Westworld line. "Doesn't look like anything to me." Any person would see the problem immediately. The model genuinely does not, until you point at it.
The real danger isn't the failure modes you learn to spot. It's the ones you can't, because you're out of your depth. If you don't have enough domain knowledge to recognize when the 90% has drifted into the wrong 10%, you won't even know to point at the problem. The inaccuracy just compounds quietly, and you end up somewhere very far from reality without ever feeling lost. You say: "Looks great, I think we're done here" wiring a relay to mains voltage for the first time.
It feels like increasingly every company is converging on the same playbook: subsidize the hardware, lock away the data, charge monthly for access, increase prices. What are you going to do, switch? They're all doing it. Customer support is now a Google form that goes nowhere - looking at you Anthropic. That's just how things work now, apparently.
I paid for my standing desk. The hardware is mine. If the manufacturer didn't build an integration, I'll add one. I paid for my Plaud recorder. My voice recordings don't need to route through anyone's cloud. The device is fully capable of being a dumb recorder, so that's what it is now. The OnePlus phone I bought for $80 to root and use as a Bluetooth bridge is, by any standard, obsolete. Too slow to use as a phone. But now it has a second life - a bridge for Claude Code to reverse engineer devices with.
Once you've done this just a couple times, you can stop choosing hardware based on what it can do - and instead think about what you can make it do. I'm shopping for a robot lawnmower right now. The best one for my yard has no Home Assistant support whatsoever. I'm buying it anyway and going to bend it to my will. My OXO coffee maker needs two physical button presses to start. Previously, that was the end of the conversation. Now I look at it and think: ESP32, a relay, 30 minutes of work.
Limitations you previously accepted as permanent stop being permanent. That's the actual shift.
I have some experience with microcontrollers. Arduinos years ago, a couple PCBs, comfortable with electronics. And I've fully committed to doing literally everything through Claude Code, which everyone on my team now comes to me for guidance on. But my story is not the proof that anyone can do hardware with AI. Not yet.
What I am is someone running an experiment to see where this goes, how far coding agents can push into the physical world, and what breaks along the way. Lots of people already use Claude Code to build and improve their Home Assistant setup. That's the start. What I'm saying is it doesn't have to stop there.
The barrier keeps dropping. M5stack makes modular hardware that clicks together like Legos. The extent of "hardware work" for many projects is literally connecting a few things with cables. The limitation is mostly in your head (and your wallet). And maybe this is just the new way we interact with everything. As long as the tool stays available, stays capable, and doesn't decide one day that you violated its terms of service.
Claude Code can be the breaker of walled gardens.
Assuming they don't try to make Claude itself a walled garden. Which, in almost certainty, they will try.
Making AI work for real businesses - not just demos.
