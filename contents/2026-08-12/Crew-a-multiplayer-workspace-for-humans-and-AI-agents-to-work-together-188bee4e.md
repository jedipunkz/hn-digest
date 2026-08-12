---
source: "https://github.com/JamelHammoud/crew"
hn_url: "https://news.ycombinator.com/item?id=49267362"
title: "Crew, a multiplayer workspace for humans and AI agents to work together"
article_title: "GitHub - JamelHammoud/crew: Build Together · GitHub"
author: "alihammoud21"
captured_at: "2026-08-12T03:56:00Z"
capture_tool: "hn-digest"
hn_id: 49267362
score: 2
comments: 0
posted_at: "2026-08-12T03:00:03Z"
tags:
  - hacker-news
  - translated
---

# Crew, a multiplayer workspace for humans and AI agents to work together

- HN: [49267362](https://news.ycombinator.com/item?id=49267362)
- Source: [github.com](https://github.com/JamelHammoud/crew)
- Score: 2
- Comments: 0
- Posted: 2026-08-12T03:00:03Z

## Translation

タイトル: Crew、人間と AI エージェントが協力するためのマルチプレイヤー ワークスペース
記事のタイトル: GitHub - JamelHammoud/スタッフ: Build Together · GitHub
説明: 一緒に構築します。 GitHub でアカウントを作成して、JamelHammoud/クルーの開発に貢献してください。

記事本文:
GitHub - JamelHammoud/クルー: 一緒に構築する · GitHub
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ジャメル・ハムード
/
乗組員
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
11,891 コミット 11,891 コミット .claude/ スキル .claude/ スキル .github .github bin bin ビルド ビルド イメージ 画像 リソース リソース スクリプト スクリプト src src テスト テスト .crew.json .crew.json .editorconfig .editorconfig .gitattributes .gitattributes .giti

gnore .gitignore .yarnrc.yml .yarnrc.yml AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE NOTICE.md NOTICE.md README.md README.md biome.json biome.json electric.vite.config.ts electric.vite.config.ts handoff-probe.mjs handoff-probe.mjs package.json package.json pdfjs-assets.ts pdfjs-assets.ts tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts 糸.ロック 糸.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
兄と私は離れて暮らしていますが、それでも一緒に何かを築きたいと思っています。
私たち (およびエージェント) がプロジェクトで共同作業できるアプリを構築するというアイデアがあり、その名前を付けました。
乗組員。
Crew が少し違うことに気づいていただければ幸いです。慎重に作られており、私たちは楽しんで作っています。
それは私たちの遊び場であり、信じられないほど速く動き、新しくて興味深いアイデアを実験できる場所です。
これが興味深いと思われる場合は、Crew を自分で試してみてください (非常に簡単です!)。
GitHub から MacOS または Windows アプリをダウンロードします
アプリを起動し、自分に名前を付けます
アカウントやサーバーはなく、すべてがローカルにあります (共有することを選択したものを除く)。
クルーを共有 (フォルダー内) するか、ローカル (アプリ内) で共有するかを決定します。
また、参加リンク (LAN 経由またはプロキシ経由) を使用して他の人のクルーに参加することもできます。
ローカル LLM CLI と @ エージェントに接続して構築を開始してください。
各エージェントは所有者のサインインした CLI アカウントを使用し、API キーや価格設定はありません
そのクルーの全員が LLM の使用制限を確認できます
複数のエージェントを同時に実行します (同じスレッド内であっても)。
プロジェクト全体で最大 10 個のアクティブなエージェント スレッドを分割表示
エージェントが「ヘルパー」間で作業を分割できるようにする
/ クール モード用のコマンド。次のようなものです。
/tickets : エージェントにチケットボードを提供し、計画を立て、進捗状況を伝え、質問することができます (私のお気に入り)
/fallback : 次の場合にどのエージェントが作業を受け取るかを選択します。

現在のものが失敗する (つまり、使用クレジットがなくなる)
/fork : ここまでのすべてをコピーし、新しいスレッドで続行します (別のエージェントを使用する可能性があります)
思い出を Crew に保存し、マルチプロバイダーのエージェントが共有します
トークンの使用量と推定トークンコストを追跡する
ブラウザは、次のような多目的の右パネルです。
レビュー : ファイルのステージング、変更の検査、コミット、プル、プッシュ、および作業のスタッシュを行う
ターミナル : クルーを離れることなく完全な CLI アクセス
ファイル : フォルダーのファイル エクスプローラー
プレビュー : 画像、ビデオ、PDF、HTML などのファイルをプレビューします。
音楽 : 一緒に同期して音楽を聴く方法
ゲーム: 退屈だった、20 個のスレッドが終了するまで他に何をするつもりですか
ヘルパー (該当する場合): スレッド内のエージェントによってディスパッチされたヘルパーの概要
計画 (該当する場合): スレッド内のエージェントによって生成された計画の概要
Board : スレッド内のエージェントによって使用されるチケットボード
クルーを開く クルーはこちら
crew ~/code/thing open そこにいるクルー
crew --share 他の人をあなたのクルーに参加させます
crew --join crew://host:port/code 誰かの Crew に参加します
「このコンピューター」の設定でアプリを介して CLI をインストールできます。
人々がどのように Crew を使用するようになるのか、ちょっと興味があります。お気軽に devjamel@gmail.com までメールをください。
Readme MIT ライセンス アクティビティ スター
1 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Build Together. Contribute to JamelHammoud/crew development by creating an account on GitHub.

GitHub - JamelHammoud/crew: Build Together · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
JamelHammoud
/
crew
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
11,891 Commits 11,891 Commits .claude/ skills .claude/ skills .github .github bin bin build build images images resources resources scripts scripts src src tests tests .crew.json .crew.json .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .yarnrc.yml .yarnrc.yml AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE NOTICE.md NOTICE.md README.md README.md biome.json biome.json electron.vite.config.ts electron.vite.config.ts handoff-probe.mjs handoff-probe.mjs package.json package.json pdfjs-assets.ts pdfjs-assets.ts tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts yarn.lock yarn.lock View all files Repository files navigation
My brother and I live apart, but still want to build things together.
We had the idea to build an app that allows us (and our agents) to collaborate on a project, we named it
Crew .
Hopefully you'll notice that Crew is a little different: it was built with care and we had fun doing it.
It's our playground, where we can move incredibly fast and experiment with new and interesting ideas.
If this sounds interesting to you, you can try Crew out yourself (it's incredibly easy!):
Download either the MacOS or Windows app from GitHub
Launch the app, and give yourself a name
There's no accounts and no server, everything's local (except what you chose to share)
Decide if you'd like your Crew to be shared (in the folder) or local (in the app)
You can also join someone else's Crew using their join link (over LAN or proxied)
Connect your local LLM CLI(s) and @ an agent to start building!
Each agent uses its owner's signed-in CLI account, there's no API keys/pricing
Everyone in that Crew can see the usage limits of your LLM
Run several agents at once (even in the same thread!)
Split-view up to 10 active agent threads, across projects
Allow agents to split up work across "Helpers"
/ commands for cool modes, like:
/tickets : gives your agent a ticket board to plan, communicate progress and ask questions (my favorite)
/fallback : choose which agent picks up the work if the current one fails (i.e. runs out of usage credits)
/fork : copy everything up to this point and continue in a new thread (potentially with a different agent)
Store memories in Crew, which your multi-provider agents share
Track token usage and estimated token cost
Browser is a multi-purpose right-panel that is home to:
Review : stage files, inspect changes, commit, pull, push, and stash work
Terminal : full CLI access without leaving crew
Files : a file explorer for your folder
Preview : previewing files like images, videos, PDFs & HTML
Music : a way to listen to music together, in sync
Games : we were bored, what else are you gonna do while you wait for your 20 threads to finish
Helpers (when applicable): an overview of any Helpers that have been dispatched by the agent(s) in the thread
Plan (when applicable): the plan outline generated by the agent(s) in the thread
Board : the ticket board used by the agent(s) in the thread
crew open Crew here
crew ~/code/thing open Crew there
crew --share let people join your Crew
crew --join crew://host:port/code join someone's Crew
You can install the CLI through the app, in Settings under "This computer"
We're kinda curious how people will end up using Crew, feel free to email me at devjamel@gmail.com !
Readme MIT license Activity Stars
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
