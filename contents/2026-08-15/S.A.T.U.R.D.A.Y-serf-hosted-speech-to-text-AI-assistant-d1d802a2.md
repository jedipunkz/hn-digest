---
source: "https://github.com/GRVYDEV/S.A.T.U.R.D.A.Y"
hn_url: "https://news.ycombinator.com/item?id=49310347"
title: "S.A.T.U.R.D.A.Y: serf-hosted speech to text AI assistant"
article_title: "GitHub - GRVYDEV/S.A.T.U.R.D.A.Y: A toolbox for working with WebRTC, Audio and AI · GitHub"
author: "soupspaces"
captured_at: "2026-08-15T14:14:26Z"
capture_tool: "hn-digest"
hn_id: 49310347
score: 1
comments: 0
posted_at: "2026-08-15T13:25:13Z"
tags:
  - hacker-news
  - translated
---

# S.A.T.U.R.D.A.Y: serf-hosted speech to text AI assistant

- HN: [49310347](https://news.ycombinator.com/item?id=49310347)
- Source: [github.com](https://github.com/GRVYDEV/S.A.T.U.R.D.A.Y)
- Score: 1
- Comments: 0
- Posted: 2026-08-15T13:25:13Z

## Translation

タイトル: S.A.T.U.R.D.A.Y: 農奴がホストする音声認識 AI アシスタント
記事タイトル: GitHub - GRVYDEV/S.A.T.U.R.D.A.Y: WebRTC、オーディオ、AI を操作するためのツールボックス · GitHub
説明: WebRTC、オーディオ、AI を操作するためのツールボックス。 GitHub でアカウントを作成して、GRVYDEV/S.A.T.U.R.D.A.Y の開発に貢献してください。

記事本文:
GitHub - GRVYDEV/S.A.T.U.R.D.A.Y: WebRTC、オーディオ、AI を操作するためのツールボックス · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
グレイデブ
/
S.A.T.UR.D.A.Y
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
117 コミット 117 コミット クライアント クライアント イメージ イメージ ログ ログ モデル モデル rtc rtc stt stt テスター テスター tts tts ttt ttt util util Web Web Whisper.cpp @ 3998465 Whisper.cpp @ 3998465 .gitignore .gitignore .gitmodules

.gitmodules ライセンス ライセンス Makefile Makefile README.md README.md docker-compose.yml docker-compose.yml go.work go.work go.work.sum go.work.sum すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Pion 、 Whisper.cpp 、および Coqui TTS で構築されたボーカル コンピューティング用のツールボックス。 WebRTC を利用した独自の自己ホスト型 J.A.R.V.I.S を構築する
デモを見る •
はじめに •
リクエスト機能 •
プロジェクトについて
仕組み
Project S.A.T.U.R.D.A.Y は、ボーカル コンピューティングのためのツールボックスです。最新の LLM へのエレガントなボーカル インターフェイスを構築するためのツールを提供します。このプロジェクトの目標は、何十年にもわたって SF 映画で約束されてきたテクノロジーを世に送り出したいという志を同じくする人々のコミュニティを育成することです。特定の AI モデルから切り離された状態を保ちながら、高度にモジュール化され、柔軟性を高めることを目指しています。これにより、新しい AI テクノロジーがリリースされたときに、シームレスなアップグレードが可能になります。
Project S.A.T.U.R.D.A.Y はツールで構成されています。ツールは、ボーカル コンピューティング スタックの特定の部分をカプセル化する抽象化です。ツールを構成する主な構成要素は 2 つあります。
エンジン - エンジンは、ツールのドメイン固有の機能をカプセル化します。このロジックは、使用される推論バックエンドに関係なく、同じままである必要があります。たとえば、STT ツールの場合、エンジンには音声アクティビティ検出アルゴリズムとカスタム バッファリング ロジックが含まれています。これにより、コードを書き直すことなくバックエンドを簡単に変更できます。
バックエンド - バックエンドは、AI 推論を実際に実行するものです。これは通常、薄いラッパーですが、より柔軟でアップグレードが容易になります。言語の相互運用を容易にするために、HTTP サーバーとインターフェースするようにバックエンドを作成することもできます。
このプロジェクトには、主に 3 種類のツールが含まれています。 3 つの主要なツールは STT、TTT、TTS です。
STT ツールはシステムの耳であり、

受信音声に対して Speech-to-Text 推論を実行します。
TTT ツールはシステムの頭脳であり、音声がテキストに変換されると、テキストからテキストへの推論を実行します。
TTS ツールはシステムの口となり、TTT ツールによって証明されたテキストに対して Text-to-Speech 推論を実行します。
以下は、メインのデモが現在どのように動作するかを示す図です。
このリポジトリに含まれるデモは、あなた自身の個人的な、自己ホスト型の J.A.R.V.I.S のようなアシスタントです。
免責事項: 私はこれを M1 Pro および Max プロセッサーでのみテストしました。ローカル推論を多く行っているため、デモではかなりの処理能力が必要です。オペレーティング システムやハードウェアが異なると、走行距離が大きく異なる場合があります。問題が発生した場合は、問題を開いてください。
デモを実行するには、いくつかの前提条件があります。
デモを実行するには、 Golang 、 Python 、 Make および C コンパイラが必要です。
デモのために実行する必要があるプロセスは 3 つあります。
RTC - RTC サーバーは Web ページと WebRTC サーバーをホストします。 WebRTC サーバーは、ページを読み込むときに接続するものであり、クライアントが音声を聴き始めるときに接続するものでもあります。
クライアント - クライアントはすべての魔法が起こる場所です。起動すると RTC サーバーに参加し、音声の聴き始めます。話し始めると、話し終えるまで着信音声がバッファリングされます。話すのをやめると、その音声に対して STT 推論が実行され、それを TTT ツールに渡してテキストに対する応答を生成し、その出力を TTS ツールに渡してその応答を音声に変換します。クライアント pkg-config と opus を使用するには 2 つのシステム ライブラリが必要です。 macOS では、これらは brew でインストールできます。
brew install opus pkg-config
TTS - TTS サーバーは、TTT ツールからのテキストが音声に変換される場所です。デモでは Coqui TTS を使用します。システムライブラリは2つあります

このツール mecab と espeak に必要なaries。 macOS では、brew でインストールできます。
醸造インストール mecab espeak
注 : 現時点では、プロセスを開始する順序が重要です。クライアントを起動する前に、RTC サーバーと TTS サーバーを起動する必要があります。
プロジェクトのルートから make rtc を実行します。
初回セットアップ: tts サーバーを初めて実行するときは、依存関係をインストールする必要があります。これには仮想環境の使用を検討してください。
cd tts/サーバー/coqui-tts
pip install -r 要件.txt
プロジェクトのルートから make tts を実行します。
クライアントには Whisper.cpp と cgo の使用が必要ですが、これは make スクリプトが処理します。
プロジェクトのルートから make client を実行します。
現在のロードマップの主な目的は、 llama.cpp などを使用して TTT 推論をローカルで実行できるようにすることです。これを公開している時点ではインターネット環境が整っておらず、これを機能させるために必要なモデル ウェイトをダウンロードできません。
私のロードマップで 2 番目に大きな項目は、セットアップと構成プロセスを継続的に改善することです。
私のロードマップの最後のことは、S.A.T.U.R.D.A.Y を使用してアプリケーションを構築し続けることです。これはプロジェクトを改善し、追加する必要がある新機能を発見するための一番の方法であるため、より多くの人が私と一緒に構築してくれることを願っています。
Discord に参加して最新情報を入手してください!
このプロジェクトは、次のオープンソース パッケージを使用して構築されています。
私は完璧とは言えませんが、インストール中にバグや見落としがある可能性があります。
問題を追加してください。不明な点がある場合はお気軽にお問い合わせください。また、Discordもあります。
オープンソース コミュニティは、貢献によって、学び、インスピレーションを与え、創造するための素晴らしい場所になります。
皆様の貢献は大変感謝しております。
機能ブランチを作成します: git checkout -b featureu

re/AmazingFeature
変更をコミットします: git commit -m 'Add some AmazingFeature'
ブランチへのプッシュ: git Push Origin feature/AmazingFeature
このプロジェクトが気に入って経済的に支援したい場合は、遠慮なくコーヒーをおごってください
GitHub @GRVYDEV ·
Twitter @grvydev ·
grvy@aer.industries に電子メールを送信する
WebRTC、オーディオ、AI を操作するためのツールボックス
Readme MIT ライセンス アクティビティ スター
38 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A toolbox for working with WebRTC, Audio and AI. Contribute to GRVYDEV/S.A.T.U.R.D.A.Y development by creating an account on GitHub.

GitHub - GRVYDEV/S.A.T.U.R.D.A.Y: A toolbox for working with WebRTC, Audio and AI · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
GRVYDEV
/
S.A.T.U.R.D.A.Y
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
117 Commits 117 Commits client client images images log log models models rtc rtc stt stt tester tester tts tts ttt ttt util util web web whisper.cpp @ 3998465 whisper.cpp @ 3998465 .gitignore .gitignore .gitmodules .gitmodules LICENSE LICENSE Makefile Makefile README.md README.md docker-compose.yml docker-compose.yml go.work go.work go.work.sum go.work.sum View all files Repository files navigation
A toolbox for vocal computing built with Pion , whisper.cpp , and Coqui TTS . Build your own personal, self-hosted J.A.R.V.I.S powered by WebRTC
View Demo •
Getting Started •
Request Features •
About The Project
How It Works
Project S.A.T.U.R.D.A.Y is a toolbox for vocal computing. It provides tools to build elegant vocal interfaces to modern LLMs. The goal of this project is to foster a community of like minded individuals who want to bring forth the technology we have been promised in sci-fi movies for decades. It aims to be highly modular and flexible while staying decoupled from specific AI Models. This allows for seamless upgrades when new AI technology is released.
Project S.A.T.U.R.D.A.Y is composed of tools. A tool is an abstraction that encapsulates a specific part of the vocal computing stack. There are 2 main constructs that comprise a tool:
Engine - An engine encapsulates the domain specific functionality of a tool. This logic should remain the same regardless of the inference backend used. For example, in the case of the STT tool the engine contains the Voice Activity Detection algorithm along with some custom buffering logic. This allows the backend to be easily changed without needing to re-write code.
Backend - A backend is what actually runs the AI inference. This is usually a thin wrapper but allows for more flexibility and ease of upgrade. A backend can also be written to interface with an HTTP server to allow for easy language inter-op.
This project contains 3 main kinds of tools. The 3 main tools are STT, TTT and TTS.
STT tools are the ears of the system and perform Speech-to-Text inference on incoming audio.
TTT tools are the brains of the system and perform Text-to-Text inference once the audio has been transformed into Text.
TTS tools are the mouth of the system and perform Text-to-Speech inference on the text proved by the TTT tool.
Here is a diagram of how the main demo currently works.
The demo that comes in this repo is your own personal, self-hosted J.A.R.V.I.S like assistant.
DISCLAIMER : I have only tested this on M1 Pro and Max processors. We are doing a lot of local inference so the demo requires quite a bit of processing power. Your mileage may very on different operating systems and hardware. If you run into problems please open an issue .
In order to run the demo there are some pre-requisites.
In order to run the demo, Golang , Python , Make and a C Complier are required.
There are 3 processes that need to be running for the demo:
RTC - The RTC server hosts the web page and a WebRTC server. The WebRTC server is what you connect to when you load the page and it is also what the client connects to to start listening to your audio.
Client - The Client is where all of the magic happens. When it is started it joins the RTC server and starts listening to your audio. When you start speaking it will buffer the incoming audio until you stop. Once you stop speaking it will run STT inference on that audio, pass it to the TTT tool to generate a response to your text and then pass that output to the TTS tool to turn that response into speech. There are 2 system libraries needed to use the client pkg-config and opus . On macOS these can be installed with brew:
brew install opus pkg-config
TTS - The TTS server is where text from the TTT tool is tranformed into speech. In the demo this uses Coqui TTS . There are 2 system libraries that are needed for this tool mecab and espeak . On macOS they can be installed with brew:
brew install mecab espeak
NOTE : For now the order in which you start the processes matters. You MUST start the RTC server and the TTS server BEFORE you start the client.
From the root of the project run make rtc
FIRST TIME SETUP : When you run the tts server for the first time you will need to install the dependencies. Consider using a virtual environment for this.
cd tts/servers/coqui-tts
pip install -r requirements.txt
From the root of the project run make tts
The client requires whisper.cpp and the use of cgo however the make script should take care of this for you.
From the root of the project run make client
The main thing on the roadmap right now is getting TTT inference to run locally with something like llama.cpp . At the time of publishing this I do not have great internet and cannot download the model weights needed to get this working.
The second largest item on my roadmap is continuing to improve the setup and configuration process.
The final thing on my roadmap is to continue to build applications with S.A.T.U.R.D.A.Y. I hope more people will build along with me as this is the #1 way to improve the project and uncover new features that need to be added.
Join the Discord to stay up to date!
This project is built with the following open source packages:
I am very from perfect and there are bound to be bugs and things I've overlooked in the installation process.
Please, add issues and feel free to reach out if anything is unclear. Also, we have a Discord.
Contributions are what make the open source community such an amazing place to be learn, inspire, and create.
Any contributions you make are greatly appreciated .
Create your Feature Branch: git checkout -b feature/AmazingFeature
Commit your Changes: git commit -m 'Add some AmazingFeature'
Push to the Branch: git push origin feature/AmazingFeature
If you like the project and want to financially support it feel free to buy me a coffee
GitHub @GRVYDEV ·
Twitter @grvydev ·
Email grvy@aer.industries
A toolbox for working with WebRTC, Audio and AI
Readme MIT license Activity Stars
38 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
