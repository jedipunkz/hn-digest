---
source: "https://guides.endusergeek.com/never-lose-your-stuff/"
hn_url: "https://news.ycombinator.com/item?id=48600205"
title: "AI-assisted off-site backup guide take ~25 minutes"
article_title: "never lose your stuff again | End User Geek"
author: "responsiblparty"
captured_at: "2026-06-19T18:44:33Z"
capture_tool: "hn-digest"
hn_id: 48600205
score: 2
comments: 1
posted_at: "2026-06-19T16:16:32Z"
tags:
  - hacker-news
  - translated
---

# AI-assisted off-site backup guide take ~25 minutes

- HN: [48600205](https://news.ycombinator.com/item?id=48600205)
- Source: [guides.endusergeek.com](https://guides.endusergeek.com/never-lose-your-stuff/)
- Score: 2
- Comments: 1
- Posted: 2026-06-19T16:16:32Z

## Translation

タイトル: AI 支援のオフサイト バックアップ ガイドには約 25 分かかります
記事のタイトル: もう持ち物を失くすことはありません |エンドユーザーマニア
説明: 今日、ラップトップにコーヒーをこぼしても、何も失われることはありません。ここ

記事本文:
巻。 1 — 推定 2026 guides.endusergeek.com エンド ユーザー オタク ガイド 平易な英語の技術的雑用 · 一度役立つと、永久に役立つ
今日ノートパソコンにコーヒーをこぼしても、何も失うことはありません。ここでは、コードを学ぶのではなく、AI アシスタントにそれを行うように指示することで、オフサイトの自動バックアップを無料で設定する方法を紹介します。 1回20分くらいで、その後はもう考えなくなります。
昨夜、私の友人がラップトップにコップ一杯の水をこぼしてしまいました。キーが機能しなくなりました。彼女は次の 1 時間を費やして、マシンが壊れる前に重要なファイルを取得できることを願いながら、必死にファイルを iCloud にドラッグしました。
彼女はバックアップをとっていなかった。ほとんど誰もそうしません。不注意だからではなく、どのバックアップ チュートリアルも「cron ジョブ」と「S3 バケット」について楽しく読んでいることを前提としているからです。あなたはしない。写真を失いたくないだけです。
変わったのは次のとおりです。もうそれを理解する必要はありません。コンピューターのターミナルに常駐する無料の AI アシスタントをインストールし、「重要なフォルダーの自動バックアップを設定してクラウド ストレージを解放する」と指示すると、ツールのインストール、構成の書き込み、毎日の実行スケジュールなど、技術的な部分を実行します。あなたはずっと平易な英語を使い続けます。
このガイドでは、一度説明します。時間については正直に言ってください。1 回限りのセットアップに約 25 分かかります。そのうちの数分は、いくつかの無料サインアップをクリックする作業に費やされます。その後、それは永遠に自動的に実行され、その存在を忘れることができます。
(コンピューターと AI を選択してください。以下の手順は、一致するように書き換えられます。)
正確な手順を確認するには、上からコンピューターと AI を選択してください。
あなたは自由なルートを選択しました。いいですね — ここからはすべて $0 です。 Google の無料 AI アシスタントである Gemini CLI を使用します。
ChatGPT の料金はすでに支払われているため、次を使用します。

Codex CLI 、プランに付属する OpenAI のターミナル アシスタント。余分に買うものは何もありません。
Claude の料金はすでに支払っているので、プランに付属している Anthropic のターミナル アシスタントである Claude Code を使用します。余分に買うものは何もありません。
ステップ 1 — AI ヘルパーをインストールする
これはすべて、コンピュータに指示を入力するプレーン テキスト ウィンドウであるターミナルで行われます。心配しないでください。貼り付けるのは数行だけです。
ターミナルを開きます。 Cmd + Space を押し、 Terminal と入力して Enter を押します。ウィンドウが開きます。それでおしまい。
ターミナルを開きます。[スタート] をクリックし、「ターミナル」(または PowerShell ) と入力して、Enter キーを押します。青い窓が開きます。それでおしまい。
Gemini CLI には、まず Node.js と呼ばれる無料のヘルパーが必要です (これは単なる配管です。一度インストールすれば、あとは忘れます)。
nodejs.org に移動し、「LTS」バージョンをダウンロードしてファイルを開き、インストーラーをクリックして実行します (続行 → 続行 → インストール)。
ターミナルに戻り、これを貼り付けて Enter キーを押します。
npm install -g @google/gemini-cli 次のように入力して起動します。
gemini 初めてブラウザが開き、Google アカウントでサインインするように求められます。そうすれば、接続が完了し、無料になります。
Gemini CLI には、まず Node.js と呼ばれる無料のヘルパーが必要です (これは単なる配管です。一度インストールすれば、あとは忘れます)。
nodejs.org に移動し、「LTS」バージョンをダウンロードしてファイルを開き、インストーラーをクリックして実行します (「次へ」→「次へ」→「インストール」)。
ターミナル ウィンドウに戻り、これを貼り付けて Enter キーを押します。
npm install -g @google/gemini-cli 次のように入力して起動します。
gemini 初めてブラウザが開き、Google アカウントでサインインするように求められます。そうすれば、接続が完了し、無料になります。
Claude Code には 1 行のインストーラーがあり、Node.js も手間もかかりません。
これをターミナルに貼り付けて Enter キーを押します。
カール -fsSL https://claude.ai/install.sh | bash 次のように入力して起動します。
クロード プロンプトが表示されたら、クロード アカウントでサインインします。
クラ

ude Code には 1 行のインストーラーがあり、Node.js も手間もかかりません。
これを PowerShell/ターミナル ウィンドウに貼り付けて Enter キーを押します。
IRM https://claude.ai/install.ps1 | iex 次のように入力して開始します。
クロード プロンプトが表示されたら、クロード アカウントでサインインします。
Codex CLI には 1 行のインストーラーがあり、Node.js は必要ありません。
これをターミナルに貼り付けて Enter キーを押します。
カール -fsSL https://chatgpt.com/codex/install.sh | sh 次のように入力して開始します。
codex 「ChatGPT でサインイン」を選択し、有料アカウントを使用します。
Codex CLI には 1 行のインストーラーがあり、Node.js は必要ありません。
これを PowerShell/ターミナル ウィンドウに貼り付けて Enter キーを押します。
IRM https://chatgpt.com/codex/install.ps1 | iex 次のように入力して開始します。
codex 「ChatGPT でサインイン」を選択し、有料アカウントを使用します。
これで、端末内で AI アシスタントが実行されました。ここからのすべては、ただ話しかけているだけです。
ステップ 2 — 無料のクラウド ロッカーを作成する
バックアップを保存する場所が必要です。 Cloudflare R2 を使用します。10 GB は無料で、ほとんどのクラウド ストレージとは異なり、データを無料で取り出すことができます。これはアカウントの作成が必要なため、AI が代わりに実行できない部分の 1 つです。約10分。
Cloudflare.com にアクセスして、無料アカウントにサインアップします (または、アカウントをお持ちの場合はログインします)。
左側のメニューで R2 をクリックします。プロンプトに従って有効にします (認証するためにカードが要求される場合がありますが、無料利用枠では料金はかかりません)。
「バケットの作成」をクリックします。 my-backups のような名前を付けます。デフォルトのままにします。作成してください。
次に、コンピュータがそのバケットに書き込むことができるようにキーが必要になります。 R2 ページで、「アカウントの詳細」→「API トークン」の横にある「管理」→「API トークンの作成」を見つけます。
バケットへの読み取りおよび書き込みアクセス権を付与し、バケットを作成します。 Cloudflareは、アクセスキーID、シークレットアクセスキー、アカウントIDの3つを表示します(エンドポイントURL内にあり、https://abc123.r2.cloudflのようになります)。

アレストレージ.com）。
その窓は開いたままにしておいてください。ステップ 4 で、これら 3 つの値を AI に貼り付けます。
ステップ 3 — 何をバックアップするかを決定します (AI を使用してこれを行います)
これは、無料の 10 GB 内に留めて、同じバケーション アルバムのコピーを 3 つバックアップすることを防ぐためのステップです。推測するつもりはありません。アシスタントに調べてアドバイスを求めることになります。
AI アシスタントが実行中であることを確認してください (「gemini claude codex」と入力すると、AI アシスタントが待機しています)。これを貼り付けます:
最大 10 GB を保持できる無料のクラウド バケットへの自動バックアップを設定したいと考えています。何かを変更する前に、ホーム フォルダーを調べて、(1) 最大のフォルダーとファイル、(2) スキップできる明らかな重複した写真やファイル、(3) ドキュメントや写真など、失われると非常にショックなものがどのフォルダーに保存されているか、アプリ、キャッシュ、ダウンロードなど、後で再ダウンロードできるものはどのフォルダーに保存されているかを教えてください。次に、10 GB 未満に収まる具体的なフォルダーのリストを提案します。何かをする前に、発見したことを説明してください。アシスタントは周りを見回して、「あなたの写真フォルダーは 40 GB です。無料枠には大きすぎますが、ドキュメント (2 GB)、デスクトップ (1 GB)、および「お気に入りの写真」アルバム (3 GB) は簡単に収まります。ダウンロードで無視できる重複画像が約 800 個見つかりました。」といった内容のメッセージが返されます。話しかけてください。何が重要かを伝えてください。ステップ 3 は、保護する価値のあるフォルダーの短いリストで終了します。これはユーザーが選択し、AI によって表示されます。
ステップ 4 — AI にすべてをセットアップするよう指示する
さて、魔法の部分です。このプロンプトを貼り付け、括弧で囲まれた部分にステップ 3 のフォルダー リストとステップ 2 の 3 つの値を入力します。
rclone と呼ばれるツールを使用して、これらのフォルダーの毎日の自動バックアップを Cloudflare R2 バケットに設定します ([フォルダー リストを貼り付けます])。私のR2の詳細は次のとおりです。
- バケット名: [あなたのバケット名]
- アカウント

ID: [あなたのアカウントID]
- アクセスキーID: [あなたのアクセスキー]
- シークレットアクセスキー: [あなたのシークレット]
- エンドポイント: https://[アカウント ID].r2.cloudflarestorage.com
これを実行する前に、各部分をわかりやすい英語で説明しながら、段階的に実行してください。
1. rclone をまだインストールしていない場合はインストールし、R2 リモートを構成します (プロバイダー Cloudflare、および no_check_bucket = true を追加します)。
2. アクセス キーとシークレットを安全に保存します。自分のアカウントのみが読み取れるファイル (ファイルのアクセス許可をロック) に保存し、クラウドと同期する場所には決して平文で保存しないでください。どこに置いたのか、そしてなぜそれが安全なのか教えてください。
3. 小さなテストアップロードを 1 回実行して、動作することを確認します。
4. 完全バックアップを 1 日に 1 回実行するようにスケジュールします。ただし、コンピューターが決まった時間に起動していると想定しないでください。スケジュールされた時刻にマシンがスリープ状態またはオフになっていた場合でも、次に起動するとすぐにバックアップが実行されるように設定します。システムがサポートしている場合は、コンピュータを短時間起動して実行します。これには、OS の組み込みスケジューラを使用してください。
5. 無料の「デッドマンズ スイッチ」アラートを設定して、バックアップが失敗した場合、または実行がサイレントに停止した場合に警告するようにします。healthchecks.io で無料アカウントを作成する手順を説明します。その後、毎日のバックアップが正常に終了した場合にのみ URL に ping を実行するようにします。つまり、エラーだけでなく、バ​​ックアップが停止した場合にもメールが届きます。
最後に、失敗のアラートを受け取った場合にどのようにフォローアップするかを教えてください。アシスタントが進行状況をナレーションします。rclone のインストール、構成の作成、テストの実行、毎日のスケジュールの設定 (Mac では、launchd または cron と呼ばれる組み込みスケジューラーが使用されます。これに触れる必要はありません) (Windows では、組み込みのタスク スケジューラーが使用されます。それに触れる必要はありません)、失敗アラートの設定などです。テストのアップロードが機能し、アラートが接続されたら完了です。
もしそのアラートが届いたら

パニックにならないでください。AI アシスタントを開いて「R2 バックアップのチェックインが停止しました。理由を調べて修正してください。」と貼り付けます。ログを読み取り、原因 (フォルダーの名前が変更されたなどの単純なものが多い) を見つけて、再び実行できるようにします。
ステップ 5 — 実際に動作していることを確認する
信仰を「終わった」と受け取らないでください。尋ねてください:
バックアップが実行されたことを示してください。現在 R2 バケットにあるものをリストし、毎日のスケジュールがどのように設定されているか、またバックアップがまだ実行中であることを来週どのように確認するかを思い出させてください。バケット内にファイルがリストされているのが表示されます。また、cloudflare.com → R2 → バケットにログインして、そこに存在することを確認することもできます。それが証拠だ。これからは毎日単独で実行されます。
20 分、ゼロドル (そうしたいなら)、そして自分で書くコードはゼロです。次回水がラップトップに遭遇したとき、失うものはラップトップだけです。
これは、端末 AI アシスタントの小さくて退屈な超能力です。あなたが望む結果を平易な英語で説明すると、AI アシスタントが機械を処理します。バックアップは、最初に指摘する価値のあるものです。
さらに多くのガイドが登場します。これは、これまで避けてきた便利で少し技術的な雑務を AI アシスタントにやってもらうシリーズの第 1 回です。
見出しが間違っている科学技術関連の記事の一次情報源の内訳。スパムなし、いつでも購読解除できます。

## Original Extract

Spill coffee on your laptop today and you should lose nothing. Here

vol. 1 — est. 2026 guides.endusergeek.com end user geek guides plain-english technical chores · useful once, useful forever
Spill coffee on your laptop today and you should lose nothing. Here's how to set up automatic, off-site backups for free — not by learning to code, but by telling an AI assistant to do it for you. About twenty minutes once, then you never think about it again.
A friend of mine spilled a glass of water on her laptop last night. Keys stopped working. She spent the next hour frantically dragging files into iCloud, hoping she'd grabbed the important ones before the machine died.
She hadn't done backups. Almost nobody has — not because they're careless, but because every backup tutorial assumes you enjoy reading about "cron jobs" and "S3 buckets." You don't. You just don't want to lose your photos.
Here's the thing that's changed: you no longer have to understand any of that. You can install a free AI assistant that lives in your computer's terminal, tell it "set up automatic backups of my important folders to free cloud storage," and it does the technical part — installs the tools, writes the config, schedules it to run every day. You stay in plain English the whole time.
This guide walks you through it once. Be honest with yourself about the time: it's roughly twenty-five minutes of one-time setup , and a few of those minutes are clicking through a couple of free signups. After that, it runs itself forever and you can forget it exists.
(Pick your computer and your AI. The steps below rewrite themselves to match.)
Pick your computer and your AI above to see your exact steps.
You picked the free route. Good — everything from here is $0. We'll use Gemini CLI , Google's free AI assistant.
You already pay for ChatGPT, so we'll use Codex CLI , OpenAI's terminal assistant that comes with your plan. Nothing extra to buy.
You already pay for Claude, so we'll use Claude Code , Anthropic's terminal assistant that comes with your plan. Nothing extra to buy.
Step 1 — Install your AI helper
This all happens in the Terminal , a plain text window where you type instructions to your computer. Don't worry, you'll only paste a couple of lines.
Open the Terminal: press Cmd + Space , type Terminal , hit Enter. A window opens. That's it.
Open the Terminal: click Start, type Terminal (or PowerShell ), hit Enter. A blue window opens. That's it.
Gemini CLI needs a free helper called Node.js first (it's just plumbing — install it once and forget it).
Go to nodejs.org , download the "LTS" version, open the file, and click through the installer (Continue → Continue → Install).
Back in Terminal, paste this and press Enter:
npm install -g @google/gemini-cli Now start it by typing:
gemini The first time, it opens your browser and asks you to sign in with your Google account . Do that, and you're connected — free.
Gemini CLI needs a free helper called Node.js first (it's just plumbing — install it once and forget it).
Go to nodejs.org , download the "LTS" version, open the file, and click through the installer (Next → Next → Install).
Back in your Terminal window, paste this and press Enter:
npm install -g @google/gemini-cli Now start it by typing:
gemini The first time, it opens your browser and asks you to sign in with your Google account . Do that, and you're connected — free.
Claude Code has a one-line installer — no Node.js, no fuss.
Paste this into Terminal and press Enter:
curl -fsSL https://claude.ai/install.sh | bash Start it by typing:
claude Sign in with your Claude account when prompted.
Claude Code has a one-line installer — no Node.js, no fuss.
Paste this into your PowerShell/Terminal window and press Enter:
irm https://claude.ai/install.ps1 | iex Start it by typing:
claude Sign in with your Claude account when prompted.
Codex CLI has a one-line installer — no Node.js needed.
Paste this into Terminal and press Enter:
curl -fsSL https://chatgpt.com/codex/install.sh | sh Start it by typing:
codex Choose "Sign in with ChatGPT" and use your paid account.
Codex CLI has a one-line installer — no Node.js needed.
Paste this into your PowerShell/Terminal window and press Enter:
irm https://chatgpt.com/codex/install.ps1 | iex Start it by typing:
codex Choose "Sign in with ChatGPT" and use your paid account.
You now have an AI assistant running in your terminal. Everything from here is just talking to it.
Step 2 — Make your free cloud locker
Your backups need somewhere to live. We'll use Cloudflare R2 : 10 GB free, and — unlike most cloud storage — free to pull your data back out. This is the one part the AI can't do for you, because it involves creating an account. About 10 minutes.
Go to cloudflare.com and sign up for a free account (or log in if you have one).
In the left menu, click R2 . Follow the prompt to enable it (it may ask for a card to verify you, but the free tier doesn't charge you).
Click Create bucket . Name it something like my-backups . Leave the defaults. Create it.
Now you need a key so your computer can write to that bucket. On the R2 page, find Account details → Manage next to API Tokens → Create API token .
Give it read & write access to your bucket and create it. Cloudflare shows you three things: an Access Key ID , a Secret Access Key , and your account ID (it's in the endpoint URL, which looks like https://abc123.r2.cloudflarestorage.com ).
Keep that window open. You'll paste those three values to your AI in Step 4.
Step 3 — Decide what to back up (do this with your AI)
This is the step that keeps you inside the free 10 GB and stops you from backing up three copies of the same vacation album. You're not going to guess — you're going to ask your assistant to look and advise.
Make sure your AI assistant is running (you typed gemini claude codex and it's waiting for you). Paste this:
I want to set up automatic backups to a free cloud bucket that holds up to 10 GB. Before changing anything, look through my home folder and tell me: (1) my biggest folders and files, (2) any obvious duplicate photos or files I could skip, and (3) which folders hold things I'd be devastated to lose — like documents and photos — versus things I could just re-download later, like apps, caches, and downloads. Then suggest a specific list of folders that fits comfortably under 10 GB. Explain what you find before doing anything. Your assistant will look around and come back with something like "Your Photos folder is 40 GB — too big for the free tier, but your Documents (2 GB), Desktop (1 GB), and your 'Favorite Photos' album (3 GB) fit easily. I found ~800 duplicate images in Downloads we can ignore." Talk to it. Tell it what matters. You end Step 3 with a short list of folders worth protecting — chosen by you, surfaced by the AI.
Step 4 — Tell your AI to set it all up
Now the magic part. Paste this prompt, filling in the bracketed bits with the folder list from Step 3 and the three values from Step 2:
Set up automatic daily backups of these folders — [paste your folder list] — to my Cloudflare R2 bucket, using a tool called rclone. Here are my R2 details:
- bucket name: [your bucket name]
- account ID: [your account ID]
- access key ID: [your access key]
- secret access key: [your secret]
- endpoint: https://[your account ID].r2.cloudflarestorage.com
Please do this step by step, explaining each part in plain English before you run it:
1. Install rclone if it isn't already, and configure an R2 remote (provider Cloudflare, and add no_check_bucket = true).
2. Store my access key and secret securely — in a file only my account can read (lock the file permissions), and never in plaintext anywhere that syncs to the cloud. Tell me where you put them and why that's safe.
3. Run one small test upload so we can confirm it works.
4. Schedule the full backup to run once a day — but don't assume my computer is awake at a fixed time. Set it up so that if the machine was asleep or off at the scheduled time, the backup still runs as soon as it's next awake, and briefly wake the computer to run it if my system supports that. Use whatever my OS's built-in scheduler does best for this.
5. Set up a free "dead man's switch" alert so I'm warned if a backup ever fails or silently stops running: walk me through making a free account at healthchecks.io, then have my daily backup ping its URL only when it finishes successfully — so I get an email if it ever goes quiet, not just if it errors.
At the end, tell me how I'd follow up if I ever get a failure alert. Your assistant will narrate as it goes — installing rclone, writing the config, running a test, setting up the daily schedule (on a Mac it uses a built-in scheduler called launchd or cron — you don't need to touch it) (on Windows it uses the built-in Task Scheduler — you don't need to touch it) , and wiring up your failure alert. When the test upload works and the alert is connected, you're done.
If that alert ever arrives, don't panic — open your AI assistant and paste: "my R2 backup stopped checking in, figure out why and fix it." It can read the logs, find the cause (often something simple, like a folder that got renamed), and get you running again.
Step 5 — Confirm it's actually working
Don't take "done" on faith. Ask:
Show me that the backup ran: list what's currently in my R2 bucket, and remind me how the daily schedule is set up and how I'd check next week that it's still running. You should see your files listed in the bucket. You can also log into cloudflare.com → R2 → your bucket and see them sitting there. That's proof. From now on it runs every day on its own.
Twenty minutes, zero dollars (if you want it to be), and zero code written by you. The next time water meets laptop, the laptop is the only thing you lose.
This is the small, boring superpower of a terminal AI assistant: you describe the outcome you want in plain English, and it handles the machinery. Backups are just the first thing worth pointing it at.
More guides coming. This is the first of a series on getting an AI assistant to do the useful, slightly-technical chores you've been avoiding.
Primary-source breakdowns of the science and tech stories the headlines get wrong. No spam, unsubscribe anytime.
