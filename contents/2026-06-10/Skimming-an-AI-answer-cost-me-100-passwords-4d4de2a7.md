---
source: "https://alwaysdraft.com/skimming-ai-answer-cost-100-passwords/"
hn_url: "https://news.ycombinator.com/item?id=48477307"
title: "Skimming an AI answer cost me 100 passwords"
article_title: "Skimming an AI answer cost me 100 passwords – Always Draft"
author: "Reebz"
captured_at: "2026-06-10T16:21:37Z"
capture_tool: "hn-digest"
hn_id: 48477307
score: 2
comments: 1
posted_at: "2026-06-10T14:55:57Z"
tags:
  - hacker-news
  - translated
---

# Skimming an AI answer cost me 100 passwords

- HN: [48477307](https://news.ycombinator.com/item?id=48477307)
- Source: [alwaysdraft.com](https://alwaysdraft.com/skimming-ai-answer-cost-100-passwords/)
- Score: 2
- Comments: 1
- Posted: 2026-06-10T14:55:57Z

## Translation

タイトル: AI の回答をスキミングすると、パスワードが 100 個必要になりました
記事のタイトル: AI の回答をスキミングすると、パスワードが 100 個必要になりました – 常に下書き
説明: 深夜に AI の回答をざっと読んで、読んでいないコマンドを貼り付けました。 15 分後、キーロガーは 100 個のパスワードを取得していました。回復には3つのクロード・コードが必要でした。

記事本文:
ホーム ✱ ブログ ✱ プロジェクト ✱ 概要
AI の回答をざっと読むと 100 個のパスワードが必要になりました
注意: この話全体は愚かです。私は疲れていて愚かでした。私はよく知っていましたが、これは「私には決して起こらないが、実際に起こった」スタイルの物語です。
私は最近、Anthropic イベントで講演 1 を行い、聴衆に「技術に詳しくない人たちへの特別な警告: 信頼できないcurl コマンドは決して実行しないでください」と言いました。質疑応答でその理由を尋ねられたので、簡単な逸話を述べましたが、実際にはブログに投稿する必要がある長い話であると答えました。今日、信じられないほどのサイバーセキュリティ スキルを備えた Mythos 2 がリリースされたため、この記事を書くのに良い時期だと思いました。これが次のとおりです。
それは真夜中頃で、私は仕事、雑用、そして新生児のことが重なって疲れていました。睡眠が嫌いな方には、次の療法を強くお勧めします。
小さな問題を解決するソフトウェアを構築し、それをオープンソースにします (修正してほしいバグを記録する人がいますが、誰が知っていたでしょうか?)
古い「スタートアップのアイデア リスト」の埃を払い、それも作成し始めましょう
妻が妊娠9ヶ月になったら引っ越し（！）
生まれたばかりの赤ちゃんを歓迎しましょう（そして上のお子さんたちも幸せにしてあげてください）
あなたの子供が幼稚園に通い始めると、あらゆる感情が生まれます。
9-9-6?これらはルーキーの番号です。私の体制に従えば、5-12-7 の時計で走ることになります。この一週間はずっとぼんやりしていましたが、たとえ質の高い深い仕事ではなかったとしても、生産的な仕事ができてうれしかったです。自動操縦と勢いで動くような作品でした。十分に長く続けていると自分自身に疑問を抱くのをやめてしまうようなもので、動きを進めていくうちに筋肉が記憶していくもので、デフォルトでは最も抵抗の少ない道に手を伸ばすことになっています。
クロード コードのプロジェクトをいじっていたところ、セッション間でプラグインの問題が発生し続けました。リロードが機能せず、つまずき続けました

over env の警告。構成を編集してトラブルシューティングを行うために、隠しファイルを再表示しようと探し回っていました。私はイライラしていたため、次の AI プロンプトが表示されました。
「システム上のすべての隠しファイルを再表示したいのですが、ターミナル コマンドを見つけてください。」いくつかのオプションが表示され、結果をクリックしていきました。疲れていて読みたくなかった。
私はコマンドをざっと見て、それがゴミではないことを確認しました…先頭に数文字、最後に数文字。頑張らなかった。今思い返してみると、見た目は奇妙だったのですが、危険だとは思いませんでした。私は疲れていたので、この 1 つのこと (たった 1 つの小さな小さなこと) を解決してから寝たいと思っていました。
これが愚かな部分です。筋肉の記憶が発動しました (マウスや Vim を使用せずに Excel を使用する人はそれを実感します): ミリ秒以内にコピーして貼り付け、Enter キーを押しました。
私が実行したコマンドは、私が読んだコマンドではありませんでした。
ターミナルでコマンドが展開されているのを見た瞬間、私の心は沈みました…それははるかに長くなりました。私が視覚的に見たものはコピーされたものではありませんでした。なぜ真ん中にcurlコマンドがあるのでしょうか？このランダムな文字は何ですか、base64 ですか?なんてことだ。
その後、ブラウザのタブがちらつきました。私が起動したわけではない Chrome ウィンドウがポップアップし、ほぼ即座に最小化されました。
疲れているのではないかと一瞬ためらいましたが、何かがおかしいと思いました。プロセスチェックを実行しました。なぜngrokが実行されているのですか?そのツールは私のコンピューターから離れた場所への接続を作成しますが、私はこれまで一度も使用したことがありませんでした。なんてことその2。
そこで、簡単な grep を実行して、ローカル ログインから始めてパスワードの一部を検索しました。それは、いくつかのフォルダーの深さの一時ディレクトリに埋められた平文ファイルの中にありました。なんてことその３。
私はマルウェアへの扉を開き、マルウェアを招き入れましたが、今ではマルウェアが私のすべてのパスワードと資格情報に勝手に侵入するようになりました。
私がするコマンド

n には、実際に読み取った部分の間に、base64 でエンコードされたペイロードが埋め込まれていました。デコードされると、3 日前に登録されたドメインにアクセスし、ngrok (アウトバウンド トラフィックを目立たなくするため、攻撃者が好む正規のトンネリング ツール) とともに一連の悪意のあるスクリプトをプルダウンしました。私が気づいたときには、マルウェアはおそらく 15 分間存在しており、キーロガーがインストールされて実行されていました。かかったのは:
私の macOS キーチェーン、システム上の他のすべてのものへのマスター キー
メモリと構成ファイルに存在する認証トークンと API キー
パスワード マネージャーのマスター パスワード、ローカル ログイン、およびキーロガーを通じてライブでキャプチャされたその期間のすべてのキーストローク
Chrome に保存されているものはすべて、およそ 100 個のパスワードと 2 枚の保存されたクレジット カードです。
Chrome は、保存されたパスワードとカードの詳細を、macOS キーチェーンの Chrome Safe Storage というエントリに保存されているキーを使用して暗号化します。 macOS は通常、別のプロセスがプロンプトを読み取る前にプロンプ​​トを表示します。そのため、私が注意をそらされている間にプロンプ​​トを承認しているところをマルウェアが捕らえたか、許可をクリックするまで要求を繰り返して私を疲弊させたかのどちらかでした。正直、覚えていません。いずれにせよ、キーを取得し、Chrome のパスワードとクレジット カード データベースはプレーン テキストで読み取れるため、それらも取得しました。希望の兆し: それらのパスワードのほとんどは何年も前のもので、クレジット カードの有効期限が切れていました。少し前に適切なパスワードマネージャーに切り替えました。
私のパスワードマネージャーボールトは無傷で届きました。保存時に暗号化されているため、マルウェアは開く方法のない BLOB をコピーしました。実際に入手したのはキーロガーを介して私のマスターパスワードであったため、私が最近アクセスしたものはすべて公開されたものとして扱う必要がありました。
それぞれ異なるジョブを持つ 3 つのクロード コード インスタンスを開きました。みんなそんなことするなって言うけど、何だ、もう午前1時だというのに、私は船乗りのように悪態をついていた

私の息：彼らは全員、--dangerously-skip-permissions を持っています!
最初の CC (ハンター) は、アクティブな脅威を探し出して破壊する必要がありました。疑わしいプロセスをすべて強制終了し、実行中のマルウェアを探し、再起動時にすべてを再起動する Launch Agent などの一般的なフックをチェックします。 2 つ目 (スキャナー) はハードドライブを調べて、最近編集したファイル、新しいディレクトリ、そこにあるべきではないものにフラグを付けました。 3 人目 (捜査官) は、暴露内容、実際に何が盗まれ、持ち出されたのか、被害がどれほど深刻だったかを解明しました。
これらを並行して実行すると、マルウェアが別のことを実行している間に 1 つのことを行うのではなく、トリアージと封じ込めを同時に行うことができます。
最も重要な発見は、Hunter CC インスタンスから得られました。マルウェアは私の起動プロセスに埋め込まれていました。 macOS では、これは、Launch Agents または Launch Daemons のエントリ、つまりログインまたは起動のたびにシステムに何かを実行するように指示するファイルを意味します。ソフトウェアプロセスを強制終了して寝ていたら、マシンの電源を入れ直した瞬間にすべてが再開されていたでしょう。 Claude Code はそれらの永続フックを見つけて、私が何かをシャットダウンする前に削除してくれました。
寝たのは午前4時半だったと思います。マルウェアは消え、永続化フックは削除され、暴露の可能性があるすべてのキーと資格情報をローテーションしました (基本的に、Investigator CC が発見した爆発範囲 - 過去 1 時間に使用されたもの、電子メールや銀行などのすべての高価値アカウント、アクティブなブラウザ セッションすべて)。技術的なクリーンアップには約 90 分かかりました。すべての銀行口座、すべての電子メール、すべての重要なサービスを調べ、純粋に予防策としてパスワードを変更するという地味な部分に、一晩かかりました。
Chrome にすべてを保存していたら、一晩中回復できずに完全に損失になっていたでしょう。

。パスワード マネージャーは持ちこたえ、Claude Code が表面化した永続化フックは起動する機会がありませんでした。 Launch Agent と Launch Daemon は、マルウェアが最初に埋め込まれる場所であり、疲れた人が最後にチェックしようと思う場所でもありません。
CURL 呼び出しにおける Base64 でエンコードされたペイロードは既知のパターンであり、私はこれが起こる前からそれを知っていました。疲れているから予想通りになってしまい、地雷を踏んでしまいました。
私が使用していた AI エージェント ツールは非フロンティア モデルで実行され、問い合わせるといくつかの Web 結果を返しましたが、そのうちの 1 つからコマンドを読まずに取り出しました。悪意のあるコマンドが検索結果のページに侵入し、未読のまま実行したのは私の間違いでした。大手企業がホストするモデルには追加のレイヤーがあるに違いありません。Claude、ChatGPT、または Gemini を日常的に使用する中で、このようなものにさらされた覚えがないからです。同じカテゴリのツールが問題の両側に存在していました。
セッションの記録は失われており、その夜の大量核攻撃と鍵の交換の狂乱の間に紛失したため、これは記憶を頼りに書かれたものです。これは、Anthropic イベントのあの部屋で私が話したのと同じストーリーを、少し長くして、愚かな部分を残したものです。
クロード コードに関する退屈な入門書 (楽しみと利益のため) - https://mitch-ribar-claude-talk.vercel.app ↩
クロード寓話 5 とクロード神話 5 - https://www.anthropic.com/news/claude-fable-5-mythos-5 ↩
電子メールまたは RSS フィード経由で私のブログを購読してください。

## Original Extract

I skimmed an AI answer at midnight and pasted a command I didn't read. Fifteen minutes later a keylogger had 100 passwords. The recovery took 3 Claude Codes.

Home ✱ Blog ✱ Projects ✱ About
Skimming an AI answer cost me 100 passwords
NB: This whole story is stupid. I was tired and stupid. I knew better, and this is an "it'll never happen to me, but it did" style of story.
I gave a talk 1 at an Anthropic event recently and I said to the audience "a special warning to the non-technical folks: never run a curl command you don't trust". During the Q&A I was asked why and I replied with a quick anecdote, but that it is actually a long story that needs a blog post. Since Mythos 2 launched today with its allegedly incredible cybersecurity skills, it seemed like a good time to write this up. So here it is:
It was around midnight and I was tired from a mixture of work, chores, and a newborn. If you hate sleep, I highly recommend this regimen:
Build software that solves your little problems and open source it (people log bugs they want fixed, who knew?)
Dust off your old "startup ideas list" and start building those too
Move house when your wife is 9 months pregnant(!)
Welcome the new baby (and keep your older children happy too)
Have one of your kids start kindergarten and all the emotion comes with that.
9-9-6? Those are rookie numbers, follow my regime and you'll be running on a 5-12-7 clock. The whole week had been a blur and I was happy to be getting some productive work done, even though it wasn't good quality deep work. It was the kind of work that runs on autopilot and momentum. The kind where you've been at it long enough that you stop questioning yourself, it's muscle memory as you go through the motions, and the default is to reach for the path of least resistance.
I was fiddling with a project in Claude Code and kept running into a plugin issue across sessions. Reload wasn't working, and I kept tripping over env warnings. I was digging around trying to unhide hidden files to edit my configuration and troubleshoot. I was getting frustrated, which led me to this AI prompt:
"I want to unhide all the hidden files on my system. Go find me a terminal command." It came back with a few options and I clicked around the results. I was tired and I didn't want to read.
I glanced at the command to make sure it wasn't garbage… a few characters at the start, a few at the end. I didn't try hard. In retrospect I remember it looked weird, but didn't think it was dangerous. I was tired and I wanted to fix this one thing (just one teeny tiny little thing) and then go to bed.
This is the stupid part. Muscle memory kicked in (people who use Excel with no mouse or Vim will get it): I copied, pasted, and hit enter in milliseconds.
The command I ran was not the command I read.
My heart sank as soon as I saw the command expand in the terminal… It was now much longer. What I saw visually was not what was copied. Why is there a curl command in the middle there? What are all those random characters, is that base64? Holy shit.
Then a flicker in my browser tabs. A Chrome window popped up that I didn't launch and minimized itself near instantly.
I hesitated for a moment thinking I was tired, but I knew something was wrong. I ran a process check. Why is ngrok running? That tool creates a connection from my computer to somewhere remote and I had never used it before. Holy shit #2.
So I ran a quick grep to search for some of my passwords, starting with my local login. There it was, sitting in a plaintext file buried in a temp directory a few folders deep. Holy shit #3.
I'd opened the door to malware, invited it in, and now it was helping itself to all my passwords and credentials.
The command I'd run contained a base64-encoded payload buried between the parts I'd actually read. Once decoded, it reached out to a domain registered 3 days earlier and pulled down a set of malicious scripts along with ngrok (a legitimate tunneling tool that attackers favor because it makes outbound traffic look unremarkable). By the time I noticed, the malware had been at it for maybe 15 minutes, keylogger installed and running. It took:
My macOS Keychain, the master key to everything else on the system
Auth tokens and API keys sitting in memory and config files
My password manager master password, my local login, and all keystrokes for that time period captured live through the keylogger
Everything stored in Chrome - roughly 100 passwords and 2 saved credit cards.
Chrome encrypts its saved passwords and card details with a key stored in the macOS Keychain, under an entry called Chrome Safe Storage. macOS normally prompts before another process can read it, so the malware either caught me approving a prompt while I was distracted, or wore me down with repeated requests until I clicked allow, I honestly don't remember. Either way, it got the key, and Chrome's password and credit card databases were readable in plain text so it got those too. A Silver Lining: Most of those passwords were years old and credit cards were expired! I switched to a proper password manager a while back.
My password manager vault came through untouched. It's encrypted at rest, so the malware copied a blob it had no way to open. What it did get was my master password through the keylogger, so anything I'd recently accessed had to be treated as exposed.
I opened 3 Claude Code instances, each with a different job. Everyone says don't do this, but what the hell, it was now 1am and I was swearing like a sailor under my breath: they all got --dangerously-skip-permissions !
The first CC (hunter) had to seek and destroy active threats: kill any suspicious processes, look for running malware, check for common hooks like Launch Agents that would restart everything on a reboot. The second (scanner) went through the hard drive, flagging recently edited files, new directories, anything that shouldn't have been there. The third (investigator) worked out the exposure, what had actually been stolen and exfiltrated, and how bad the damage was.
Running them in parallel meant I could triage and contain at the same time, instead of doing one thing while the malware kept doing another.
The most important discovery came from the Hunter CC instance. The malware had embedded itself into my startup processes. On macOS this means entries in Launch Agents or Launch Daemons, files that tell the system to run something on every login or boot. If I'd killed the software processes and gone to bed, everything would have restarted the moment I turned the machine back on. Claude Code found those persistence hooks and removed them for me before I shut anything down.
I went to bed at 4:30am I think. The malware was gone, the persistence hooks were removed, and I'd rotated every key and credential with any chance of exposure (basically the blast radius that Investigator CC found - anything used in the prior hour, every high-value account like email and banking, and every active browser session). The technical cleanup took about 90 mins. The unglamorous part, going through every bank account, every email, every critical service and changing passwords purely as precaution, took the rest of the night.
If I'd been storing everything in Chrome, the whole night would have been a total loss instead of a recoverable one. The password manager held up and the persistence hooks Claude Code surfaced never got the chance to fire. Launch Agents and Launch Daemons are the first place malware embeds itself and the last place a tired person thinks to check.
A base64-encoded payload in a curl call is a known pattern, and I knew it before this happened. Being tired just made me predictable, and I stepped on that land mine.
The AI agent tool I was using ran on a non-frontier model and returned a handful of web results when I asked, and I lifted the command out of one of them without reading it. The malicious command rode in on a page from that search and running it unread was my mistake to own. There must be extra layers in the models hosted by the big players, because I can't recall being exposed to something like this in my daily use of Claude, ChatGPT, or Gemini. The same category of tool sat on both sides of the problem.
The session transcript is gone, lost during the mass-nuking and key rotation frenzy on the night, so this is written from memory. It's the same story I told that room at the Anthropic event, just longer, and with the stupid parts left in.
A Boring Primer on Claude Code (for Fun and Profit) - https://mitch-ribar-claude-talk.vercel.app ↩
Claude Fable 5 and Claude Mythos 5 - https://www.anthropic.com/news/claude-fable-5-mythos-5 ↩
Subscribe to my blog via email or RSS feed .
