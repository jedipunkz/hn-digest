---
source: "https://www.macstories.net/reviews/open-minis-is-the-ios-agent-i-wish-siri-ai-could-be/"
hn_url: "https://news.ycombinator.com/item?id=48981026"
title: "Open Minis Is the iOS Agent I Wish Siri AI Could Be"
article_title: "Open Minis Is the iOS Agent I Wish Siri AI Could Be - MacStories"
author: "ExMachina73"
captured_at: "2026-07-20T17:24:22Z"
capture_tool: "hn-digest"
hn_id: 48981026
score: 3
comments: 0
posted_at: "2026-07-20T16:23:58Z"
tags:
  - hacker-news
  - translated
---

# Open Minis Is the iOS Agent I Wish Siri AI Could Be

- HN: [48981026](https://news.ycombinator.com/item?id=48981026)
- Source: [www.macstories.net](https://www.macstories.net/reviews/open-minis-is-the-ios-agent-i-wish-siri-ai-could-be/)
- Score: 3
- Comments: 0
- Posted: 2026-07-20T16:23:58Z

## Translation

タイトル: Open Minis は、Siri AI にそうなってほしい iOS エージェントです
記事のタイトル: Open Minis は Siri AI になればいいのにと思う iOS エージェントです - MacStories
説明: 時々、iOS 上の特定の分野のソフトウェアに対する私の期待をリセットしたり、全く新しいカテゴリを作成したりするサードパーティ アプリに遭遇します。 MacStories でアプリのレビューを書いてきた 17 年間で、編集、ワークフロー、
[切り捨てられた]

記事本文:
watchOS 27: MacStories パブリック ベータ プレビュー
macOS 27 Golden Gate: MacStories パブリック ベータ プレビュー
Siri AI と Apple のトリクルアップ戦略
HomeKit ホームへのプライベートな読み取り専用ウィンドウ
Open Minis は Siri AI になればいいのにと思う iOS エージェントです
時々、iOS 上の特定の分野のソフトウェアに対する私の期待をリセットしたり、全く新しいカテゴリを作成したりするサードパーティ アプリに遭遇します。 MacStories でアプリのレビューを書いてきた 17 年間で、こうした瞬間が何度かありました。エディトリアル、ワークフロー、オブシディアン、スカイ、そして最近では OpenClaw が思い浮かびます。ここ数週間、私は Open Minis でまたそのような瞬間を経験しました。Open Minis は、iPhone および iPad 用の新しいチャットボット アプリであり、今日、Siri AI のエージェントの将来をプレビューしたと最もよく表現できます。
Open Minis は、会話にあらゆるフロンティア モデルを使用できるようにするオンデバイス エージェントですが、他の AI ラッパーとは異なり、このアプリはサードパーティ開発者が利用できる公式 API を使用して、あらゆる種類のネイティブ Apple システム フレームワークと深く統合されています。 Open Mini は、以前に見たリマインダーと音楽を制御できますが、カレンダー、マップ、HomeKit、HealthKit、ファイルも制御できます。 OCR の Vision、自然言語処理の Apple NLP、iOS クリップボード、音声認識とディクテーション、NFC、Bluetooth などのフレームワークとも連携できます。さらに、Open Minis には、ファイル アプリ内に独自の構成可能なワークスペースが付属しており、オンデバイス メモリ システムがあり、（繰り返しになりますが、iOS 開発者にとってネイティブな）組み込み WebKit Web ビューを介したブラウザ使用機能を備えています。
実際のところ、Open Minis は、Claude Code と OpenClaw を、強力なツール呼び出しハーネスを備えた Apple ユーザー向けに設計された 1 つの直感的なエージェント エクスペリエンスにまとめた場合に起こることです。

それ自体を改善および微調整することができ、Apple エコシステム向けに特別に設計された統合も可能です。これは、私がここしばらく見た中で最も印象的なインディーズ アプリです。 Siri AI にはパワー ユーザー機能がないことに失望しているのであれば、これは、真のエージェント機能を備えたフロンティア モデルを搭載した場合にのみ Siri AI がどのようなものになるかを示す、現在出荷されている実際の例です。
Open Minis のスリーブに隠されたトリックは、明白であると同時に独創的です。Open Minis がサポートするすべての Apple フレームワーク統合は、ネイティブ Apple API (EventKit、MusicKit、HomeKit など) を、最新の LLM が理解して操作できるコード環境にブリッジするコマンドライン インターフェイス (CLI) によって強化されています。これらすべてをバックグラウンドで動作させるには、Open Minis の開発者は、これらのフレームワークごとに専用の CLI を作成する必要がありました。これは、組み込みの Linux ターミナルにアクセスし、アプリのワークスペースにサンドボックス化されたモデル ハーネスによって駆動されます。そして、ここにもう 1 つのトリックがあります。
具体的には、Open Minis は、サンドボックス方式で ARM64 アーキテクチャ用の永続的なエミュレートされた Linux 環境を操作する人気の Alpine Linux シェルである iSH のバージョンを実行します。 Open Minis 内のシェルは、Alpine の apk パッケージ マネージャーを使用して Web から外部パッケージをダウンロードでき (つまり、ffmpeg や yt-dlp などの一般的なツールにアプリから直接アクセスできます)、独自の Linux ファイルシステム (ネイティブ ファイル プロバイダーの場所としてファイル アプリで公開) を持ち、ブリッジされた Apple CLI やあらゆる種類の MCP サーバー (最近の AI 業界のもう 1 つのトレンドである MCP コード実行経由) を呼び出すことができます。このように考えてください。Mac 上の Claude Code または Codex がターミナル アプリで実行することによってファイル システム全体で何でもできるのであれば、Open Mini は Windows アプリで (ほぼ) 何でもできます。

Python の実行、Git と SSH の使用、標準 Unix テキスト ユーティリティへのアクセスなど、組み込みのターミナル エミュレーターのおかげでアプリのサンドボックスを薄くします。
Minis の組み込みシェル (左) とファイル アプリの専用ワークスペース (右)。
ただし、これは、Open Mini を使用するときに常にターミナル UI を見ていることになるという意味ではありません。基盤となるシェルは、単にこれらすべてを舞台裏で可能にするものです。アプリでは、ピクチャ イン ピクチャの端末プレビューを表示できます。これを開いて動作をプレビューできますが、そうする必要はありません。アプリのメイン インターフェイスは、クラシックなチャットボット スタイルの UI で、使用しているモデルに応じて、会話の記録にインターリーブされた思考ブロック、ツール呼び出し、およびメッセージが表示されます。
Minis アプリの通常のチャットボット UI。
Open Minis (OpenClaw を彷彿とさせます。あるいは、より適切には、Hermes Agent などの自己改善システムを彷彿とさせます) のもう 1 つの隠れた力は、専用の Minis CLI を介して独自のワークスペースを検査および変更できることです。たとえば、Minis に (もちろん自然言語で) 新しいモデルまたは MCP サーバーを追加するよう依頼すると、アプリはネイティブ設定画面を表示して、そのための許可を求めます。パスワードまたは API トークンを保存したい場合、アプリはプレーン テキストで保存しません。認証情報を安全に保存できるネイティブのキーチェーン ベースのページが表示されます。そしてさらに重要なことに、Minis 独自の CLI は、アプリのメモリ システム (Markdown ファイルに基づく) を介して以前の会話を思い出し、スキルを作成し、アプリ自体に関する質問に答えるために使用されます。たとえば、iOS ユーザー向けにどのような Apple CLI があるのか説明してほしいと Minis に尋ねたところ、喜んで完全なガイドを提供してくれました。
Minis 自身と Minis が持つ CLI について尋ねます。
あるいは、私たちが取り組んできたことをすべて思い出してほしいと頼んだとき、

日中、Minis は CLI の 1 つを使用して、以前のチャットのトランスクリプト (これもファイル アプリに保存されているだけです) を読み取り、記憶を思い出し、完了したすべてのタスクの有用な概要を私に提示してくれました。
Mini は、自身のセッションだけでなく、ユーザーの記憶を検査して、以前の会話を思い出すことができます。
Mini に新しい機能を追加する場合も同様です。 Python にアクセスできる Linux シェルを実行しており、永続的なファイル ワークスペースにアセットを保存できるため、GitHub リポジトリへのリンクを与えることで、Minis に独自の Apple Frames CLI をインストールするよう依頼できると考えました。 GPT-5.6 Sol (アプリで私が選んだモデル) は躊躇しませんでした。Frames CLI のすべての依存関係を見つけてインストールし、CLI の公式スキルを取得し、1 分後には準備が整いました。 Minis にはネイティブの Photos 統合も付属しているため、次のことを尋ねることができました。
最近の 3 つのスクリーンショットを取得し、フレームの色をランダム化して結合し、画像を一時ファイル フォルダーに保存します。
数秒後、Apple CLI が呼び出されて、Minis は結果のフレーム画像を私が選択したフォルダーに保存しました。 Minis のサンドボックスの外にあるファイル アプリの任意のフォルダーで作業できる機能は、アプリが活用しているもう 1 つの公式 iOS API です。私たちは何年も前に Working Copy の Files アプリから特定のフォルダーを「マウント」する機能を初めて確認し、その後この機能は Ulysses 、 iA Writer 、 MWeb 、その他のドキュメント ベースのアプリに採用されました。 Open Minis では、ファイル アプリ内の任意のファイル プロバイダーの場所にある任意のフォルダーをブックマークし、そこにファイルを保存するか、そこからファイルを読み取るようにエージェントに指示できます。
Mini は iOS クリップボードに画像を書き込むことはできませんが、ファイル アプリの (以前にブックマークした) 任意の場所にファイルを保存できます。
マウントされたフォルダーを管理するための Minis の設定。
これは感じますよ、ママ

私にとっては不可解なことですが、これは、大規模で強力なモデルと、iOS および iPadOS の想像以上にオープンで安全な性質を組み合わせた場合に何が可能になるかを示す数多くの例のうちの 1 つにすぎません。
iOS での Open Mini の操作
最初はこれらすべてを理解するのが大変ですが、Open Minis が私と同じように、フロンティア モデルの使用や AI の最先端での生活を楽しんでいる人々のために設計されたパワーユーザー ツールであることは否定できません。ただし、このケースで異なるのは、他のアプリとは異なり、Open Mini が Apple によって作成され、サードパーティ開発者にリリースされ、現時点では Siri AI によってまったく使用されていないテクノロジーを完全にサポートすることで、Apple ファーストの性質を取り入れているという事実です。
AI チャットボット アプリを「レビュー」するのは、人によってユースケースが非常に異なるため不可能ですが、それでも、iOS に深く結びついた AI で何が実現できるかについてのアイデアを得るために、そしてできれば期待をリセットするために、いくつかのシナリオを説明します。
まず、Open Minis では、最新の GPT-5.6 モデル、Claude Fable、OpenRouter のあらゆるモデルなど、任意のモデルを使用できます。 (私は Open Minis に、Mac Studio で実行する OpenAI API 用の独自のプロキシを作成するよう依頼しました。そうすれば、Sol を高速モードで使用できるようになりました。うまくいきました!) キーボードの上にある / ボタンを押すと、推論レベル、圧縮、記憶、スキルをトリガーできます。また、アプリの設定からモデル グループ、モデルのデフォルト、サブエージェントを作成し、特定のタスクに特定のサブエージェントのセットを使用するようにモデルに要求すると、それが実行されます。
Mini でのモデル プロバイダー、グループ、サブエージェントの管理。
このコンセプトを念頭に置いてこのアプリに取り組んでください。アプリのほぼすべてがオープンであり、話しかけるだけで設定可能です。アプリを書き直す必要はありません」

コード自体を編集するだけで、Open Mini 自体や設定に関するさまざまな調整を Open Mini に依頼でき、エージェントがそれを実行します。個人的には、この種の用途では、OpenClaw よりもこれがより効果的でエレガントで安定したハーネスであると思います。アプリの UI にはもう少し磨きをかけることができますが (推論ピッカーは好きではありません。デフォルトで思考ステップを折りたたむことができればいいのにと思います)、モデルと会話して Minis のワークスペースを好みに合わせて調整するのは非常に楽しいです。実際、ただ単に物事を行うこともできます。
数日前、私たちのアパートの温度と湿度の状態が気になりました。そこで私はミニスにこう尋ねました。
私の家の温度と湿度の状態についての簡潔なレポートを取得し、それを素敵な表として提示してください。
ネイティブの HomeKit 許可プロンプトが表示され、数分後、Minis の Sol は各センサーからの温度と湿度の測定値の生のダンプを解析し、HomeKit で構成された各部屋を確認し、結果を表示しました。比較のために、Siri AI は次のようなものを提示しました。
また、Open Minis での写真の統合と Siri AI の写真ライブラリへのアクセスを比較するのも面白いだろうと思いました。理論的には、両方とも同じフレームワークを使用し、各写真に同じメタデータを使用していますが、Siri AI は「個人的なコンテキスト」にもアクセスできます。そこで私はこう尋ねました。
私の写真から判断すると、私は最近休暇をとっていたと思いますか?
結果がすべてを物語っています。
ヘルスキットはどうですか？繰り返しになりますが、Open Minis はここで「ハック」や回避策を採用していないことに注意してください。すべての開発者が使用できる同じネイティブ Apple API とフレームワークにアクセスするだけです。私は尋ねました：
最近休暇をとって大きな船の写真を撮ったときの写真を見つけてください。その夜、私はどれくらい歩きましたか？
Minis は、このリクエストについて別の方法で考え始めました。

記憶を優先し、私たちが最近の休暇について話し合ったことに気づき、その日付を調べるプロセスが即座に中断される可能性がありました。それから、それらの日付で私の写真をフィルタリングし始め、パリヌーロ船の元の写真を見つけました、そして…ええ：
リクエスト内の思考ステップを調べることは、例外的な iOS ハーネスによって駆動されるフロンティア モデルがこのような問題をどのように推論できるかを確認する楽しい演習でもあります。このモデルは、「Minis Computer」（つまり、前述の Linux シェル）を使用して、Apple Photos CLI と呼ばれるファイル アプリから思い出をマークダウン ファイルとして取得し、写真のグリッドを含むコンタクト シートを作成し、2 つの候補を選択し、最良の写真を選択し、日付を分離し、その日付に対して HealthKit を呼び出して、初めて答えを構築しました。そして、これらすべてはデバイス上で実行される決定論的なコードに基づいているため、応答は速く、幻覚は含まれませんでした。
私の質問と iOS から返されたデータをエージェントがどのように判断したか。
そこで私は、さらに一歩進めることにしました。Minis に、最近の休暇のハイライト写真を撮って、Apple の MapKit JS を利用した地図上に配置し、地図上に写真を含む優れたインタラクティブな HTML アーティファクトを作成するよう依頼しました。これには Kim K3 を使用しました。 Minis はまず、MapKit 開発者トークンを安全に保管するように私に求め、それから作業を始めました。私

[切り捨てられた]

## Original Extract

Every once in a while, I come across a third-party app that either resets my expectations for a particular niche of software on iOS or creates an entirely new category altogether. I’ve had quite a few of these moments in the 17 years I’ve been writing app reviews at MacStories: Editorial, Workflow,
[truncated]

watchOS 27: The MacStories Public Beta Preview
macOS 27 Golden Gate: The MacStories Public Beta Preview
Siri AI and Apple’s Trickle-Up Strategy
A Private, Read-Only Window Into Your HomeKit Home
Open Minis Is the iOS Agent I Wish Siri AI Could Be
Every once in a while, I come across a third-party app that either resets my expectations for a particular niche of software on iOS or creates an entirely new category altogether. I’ve had quite a few of these moments in the 17 years I’ve been writing app reviews at MacStories: Editorial , Workflow , Obsidian , Sky , and, most recently, OpenClaw come to mind. Over the past few weeks, I’ve had another such moment with Open Minis , a new chatbot app for iPhone and iPad that I can best describe as getting a preview of Siri AI’s agentic future, today .
Open Minis is an on-device agent that lets you use any frontier model for conversations, with a twist: unlike other AI wrappers, this app deeply integrates with all sorts of native Apple system frameworks using official APIs available to third-party developers. Open Minis can control Reminders and Music, which we have seen before , but also Calendar, Maps, HomeKit, HealthKit, and Files; it can even work with frameworks such as Vision for OCR, Apple NLP for natural language processing, the iOS clipboard, speech recognition and dictation, NFC, and Bluetooth. Furthermore, Open Minis comes with its own configurable workspace in the Files app, has an on-device memory system, and features browser-use capabilities via a (once again, native for iOS developers) built-in WebKit web view .
Effectively, Open Minis is what would happen if you rolled Claude Code and OpenClaw into one intuitive agentic experience designed for Apple users, with a powerful tool-calling harness that can improve and tweak itself, and with integrations specifically designed for the Apple ecosystem. It is the most impressive indie app I’ve seen in a while. If you were disappointed by the lack of power-user capabilities in Siri AI , this is a real, currently shipping example of what Siri AI could be if only it were powered by a frontier model with truly agentic functionalities.
The trick hidden up Open Minis’ sleeve is as ingenious as it is obvious: all the Apple framework integrations it supports are powered by command-line interfaces (CLIs) that bridge native Apple APIs (EventKit, MusicKit, HomeKit, etc.) to a code environment that modern LLMs can understand and operate with. To make all of this work behind the scenes, the developer of Open Minis had to create dedicated CLIs for each of these frameworks, which – and here comes the other trick – are driven by a model harness that has access to a built-in Linux terminal and is sandboxed to the app’s workspace.
Specifically, Open Minis runs a version of iSH , a popular Alpine Linux shell that operates a persistent, emulated Linux environment for ARM64 architecture in a sandboxed fashion. The shell inside Open Minis can download external packages from the web using Alpine’s apk package manager (which means that you can access popular tools such as ffmpeg and yt-dlp directly from the app), has its own Linux filesystem (exposed in the Files app as a native file provider location), and is able to call bridged Apple CLIs, as well as any kind of MCP server (via MCP code execution, another trend in the AI industry these days). Think about it this way: if Claude Code or Codex on the Mac can do anything across the entire filesystem by virtue of running in the Terminal app, Open Minis can do (almost) anything within its app sandbox thanks to a built-in terminal emulator – including running Python, using Git and SSH, and accessing standard Unix text utilities.
Minis’ built-in shell (left) and dedicated workspace in the Files app (right).
However, this doesn’t mean that, when using Open Minis, you’ll be looking at a terminal UI all the time: the underlying shell is simply what makes all of this possible behind the scenes. The app can display a picture-in-picture terminal preview that you can open to preview what it’s doing, but you don’t have to. The app’s main interface is your classic, chatbot-style UI that, depending on the model you’re using, will display interleaved thinking blocks, tool calls, and messages in a conversation transcript.
The regular chatbot UI of the Minis app.
The other hidden power of Open Minis (which is reminiscent of OpenClaw, or, perhaps more aptly, a self-improving system such as Hermes Agent ) is that it can inspect and modify its own workspace via dedicated Minis CLIs. For example, you can ask Minis – in natural language, of course – to add a new model or MCP server, and the app will bring up a native settings screen and ask for permissions to do so. If you want to store a password or API token, the app won’t do it in plain text: it’ll display a native, Keychain-based page where you can securely store your credentials. And more importantly, Minis’ own CLIs are used to recall previous conversations via the app’s memory system (based on Markdown files), create skills, and answer questions about the app itself. For instance, when I asked Minis to describe what Apple CLIs it has for iOS users, it happily provided a full guide:
Asking Minis about itself and the CLIs it has.
Or when I asked to recall all the things we worked on during the day, Minis used one of its CLIs to read transcripts of previous chats (which, again, are simply stored in the Files app), recall memories, and present me with a useful summary of all the tasks we completed.
Minis can inspect its own sessions as well as user memories to recall previous conversations.
The same is true for adding new functionalities to Minis. Since it’s running a Linux shell with access to Python and can store assets in a persistent Files workspace, I figured I could ask Minis to install my own Apple Frames CLI for itself by giving it a link to my GitHub repo. GPT-5.6 Sol (my model of choice in the app) didn’t hesitate: it found all the dependencies for the Frames CLI, installed them, pulled the CLI’s official skill, and was good to go a minute later. And since Minis also comes with a native Photos integration, I was able to ask:
Grab my three most recent screenshots, randomize their frame colors, merge them, and save the image to my Temp Files folder.
A few seconds and Apple CLI calls later, Minis saved the resulting framed image to a folder of my choosing. The ability to work in any folder in the Files app outside of Minis’ sandbox is another official iOS API that the app is leveraging; we first saw the ability to “mount” specific folders from the Files app in Working Copy years ago , and the feature was later adopted by the likes of Ulysses , iA Writer , MWeb , and other document-based apps. In Open Minis, you can bookmark any folder from any file provider location in the Files app and tell your agent to save files to or read them from there.
Minis can’t write images to the iOS clipboard, but it can save files in any (previously bookmarked) location of the Files app.
Minis’ settings for managing mounted folders.
This feels magical to me, but it’s merely one of the many examples of what is possible when you combine large, powerful models with the more-open-than-you-think , secure nature of iOS and iPadOS.
Working with Open Minis on iOS
All of this a lot to wrap your head around at first, and there’s no denying that Open Minis is a power-user tool designed for people who, like me, enjoy using frontier models and living at the bleeding edge of AI. What’s different in this case, though, is the fact that Open Minis – unlike other apps – embraces its Apple-first nature by fully supporting technologies that are made by Apple, released to third-party developers, and not being used by Siri AI at all at the moment.
It’s impossible to “review” an AI chatbot app since everyone’s use cases are so different, but I’m going to nevertheless describe a few scenarios that should give you an idea – and hopefully reset your expectations, too – of what can be accomplished with an AI that has deep hooks into iOS.
For starters, you can use any model you want in Open Minis, including the latest GPT-5.6 models, Claude Fable, anything from OpenRouter , and more. (I even asked Open Minis to create our own proxy for the OpenAI API, running on my Mac Studio, so I could use Sol in Fast mode ; it works!) You can press the / button above the keyboard to trigger reasoning levels, compaction, memories, and skills; you can also create model groups, model defaults, and subagents from the app’s settings, then ask the model to use a specific set of subagents for a particular task, and it’ll do it.
Managing model providers, groups, and subagents in Minis.
Approach this app with this concept in mind: nearly everything about it is open and configurable by just talking to it. Short of rewriting the app’s code itself, you can ask Open Minis to tweak all sorts of things about itself and its own settings, and the agent will do it. Personally, I find this to be a more effective, elegant, and stable harness than OpenClaw for these kinds of things. Even though the app’s UI could use some polish (I don’t like the reasoning picker, and I wish I could collapse thinking steps by default), it is incredibly fun to tune Minis’ workspace to your liking by talking to a model. You can, in fact, just do things .
A few days ago, I was curious about the state of the temperature and humidity in our apartment. So I asked Minis:
Get me a concise report of the state of the temperature and humidity in my house, present it as a nice table.
A native HomeKit permission prompt and a couple minutes later, Sol in Minis parsed a raw dump of temperature and humidity readings from each sensor, reviewed each room configured in HomeKit, and presented the results. Siri AI, for comparison, presented… something :
I also figured it’d be cool to compare the Photos integration in Open Minis to Siri AI’s photo library access. In theory, they’re both using the same framework, with the same metadata for each photo, but Siri AI also has access to “personal context”, right? So, I asked:
Based on my Photos, would you say I was recently on vacation?
I will let the results speak for themselves:
How about HealthKit? Again, keep in mind that Open Minis isn’t employing any “hacks” or workarounds here; it’s simply accessing the same native Apple APIs and frameworks that all developers can use. I asked:
Find the photo of when I was on vacation recently and took a picture of a large ship. How much did I walk that night?
Minis started thinking about this request in a different way: it checked its memory first, and it realized that we recently discussed my recent vacations, so it could immediately short-circuit the process of finding out about those dates. Then, it started filtering my photos by those dates, found the original photo of the Palinuro vessel, and… yeah :
Inspecting the thinking steps in the request is also a fun exercise in seeing how a frontier model driven by an exceptional iOS harness can reason through a problem like this. The model – using the ‘Minis Computer’ (i.e., the aforementioned Linux shell) – fetched memories as Markdown files from the Files app, called the Apple Photos CLI, assembled a contact sheet with a grid of photos, picked two candidates, selected the best photo, isolated the date, called HealthKit for that date, and only then built an answer. And since all of this is based on deterministic code that was executed on-device, the response was fast and did not contain any hallucinations.
How the agent reasoned through my question and data returned by iOS.
I then decided to go a step further: I asked Minis to take photo highlights from my recent vacations, put them on a map powered by Apple’s MapKit JS , and create a nice, interactive HTML artifact with photos on a map. I used Kimi K3 for this. Minis first asked me to securely store my MapKit developer token, then got to work. I

[truncated]
