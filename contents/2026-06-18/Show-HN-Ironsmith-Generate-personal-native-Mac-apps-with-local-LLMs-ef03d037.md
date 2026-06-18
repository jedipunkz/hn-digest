---
source: "https://github.com/Jeidoban/Ironsmith"
hn_url: "https://news.ycombinator.com/item?id=48586323"
title: "Show HN: Ironsmith – Generate personal native Mac apps with local LLMs"
article_title: "GitHub - Jeidoban/Ironsmith: Create personal Mac apps instantly with a prompt. Supports on-device and cloud LLMs · GitHub"
author: "jeidoban"
captured_at: "2026-06-18T16:14:52Z"
capture_tool: "hn-digest"
hn_id: 48586323
score: 1
comments: 1
posted_at: "2026-06-18T14:51:15Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Ironsmith – Generate personal native Mac apps with local LLMs

- HN: [48586323](https://news.ycombinator.com/item?id=48586323)
- Source: [github.com](https://github.com/Jeidoban/Ironsmith)
- Score: 1
- Comments: 1
- Posted: 2026-06-18T14:51:15Z

## Translation

タイトル: Show HN: Ironsmith – ローカル LLM を使用して個人用ネイティブ Mac アプリを生成する
記事タイトル: GitHub - Jeidoban/Ironsmith: プロンプトを表示して個人用 Mac アプリを即座に作成します。オンデバイスおよびクラウド LLM をサポート · GitHub
説明: プロンプトを表示して個人用 Mac アプリを即座に作成します。オンデバイスおよびクラウド LLM をサポート - Jeidoban/Ironsmith

記事本文:
GitHub - Jeidoban/Ironsmith: プロンプトを表示して個人用 Mac アプリを即座に作成します。オンデバイスおよびクラウド LLM をサポート · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
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
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
児童番
/
鉄工
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション

オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット .github/ workflows .github/ workflows .vscode .vscode Config Config Ironsmith Ironsmith IronsmithTests IronsmithTests アセット/ readme アセット/ readme スクリプト スクリプト .gitignore .gitignore AGENTS.md AGENTS.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンスPackage.resolved Package.resolved Package.swift Package.swift README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Ironsmith は、AI を使用して小規模な個人用 Mac アプリを作成するための、無料のオープンソース macOS メニュー バー アプリです。必要なものを説明すると、Ironsmith によってネイティブ SwiftUI アプリが生成、構築、保存され、起動、編集、およびアプリケーション フォルダーへのエクスポートが可能になります。
本物の Mac アプリを構築します。生成されるアプリは Swift および SwiftUI アプリであり、Electron は見つかりません。
メニューバーに常駐します。新しいアプリの作成、保存したアプリの実行、既存のアプリの編集、以前のバージョンの復元、コードの表示、完成したアプリ バンドルのエクスポートが非常に便利になります。
ローカル AI と連携します。 Ironsmith はローカル AI サポートを念頭に置いて設計されており、すぐに Ollama をサポートします。 OpenAI 互換の API に接続することもできるので、LM Studio や Llama.cpp もうまく機能します。 Apple の組み込み Foundation モデルを使用して、非常に単純なアプリを構築することもできます。
ホストされたモデルもサポートします。 OpenAI、Anthropic、Gemini 用の独自の API キーを持参するか、API キーをスキップして Ironsmith にサインインしてすぐにアクセスします。
Xcodeは必要ありません。生成されたアプリはすべて Swift パッケージであり、完全な Xcode ではなく、軽量の Xcode コマンド ライン ツールを使用して完全に構築されます。実際、Ironsmith 自体は Xcode さえ使用していません。
デフォルトですべてのアプリをサンドボックス化します。生成されたアプリは、サンドボックス化と強化されたランタイムが有効になった署名付きアプリ バンドルとして構築され、impa を大幅に削減します。

バグ、間違い、または悪意のある動作の可能性があります。カメラやマイクへのアクセスなどの機密性の高い権限も明示的に有効にする必要があります。ただし、これらの保護を無効にすることもできます。無効にする場合は、コードを実行する前にレビューすることを強くお勧めします。
Ironsmith は、焦点を絞ったユーティリティ、つまり、存在したいが自分で探したり構築したくない小さなアプリに最適です。とはいえ、GPT‑5.5 や Opus 4.8 などのより高機能なモデルを使用すると、驚くほど洗練されたアプリを作成できます。
試してみることができるプロンプトの例をいくつか示します。
「スクリーンショットのフォルダーの名前を日付とウィンドウのタイトルに基づいて変更するユーティリティを作成します。」
「PDF をページごとに 1 つのファイルに分割する小さなアプリを構築します。」
「コピーされた URL から追跡パラメータを削除するクリップボード クリーナーを構築します。」
「重複行と欠損値を強調表示する小さな CSV インスペクターを作成します。」
最新の Ironsmith ビルドを GitHub Releases または Web サイトからダウンロードします。
Ironsmith には macOS 26 以降が必要で、Intel と Apple Silicon Mac の両方をサポートしています。利用可能な場合は Apple Intelligence が有効になっていることを確認してください。 Ironsmith はこれを使用してアプリ アイコンを作成し、組み込みの基盤モデルを提供します。
最初の起動時に、生成されたアプリがローカルでコンパイルされるため、Ironsmith は Xcode コマンド ライン ツールをチェックします。これらが見つからない場合、macOS はそれらをインストールするように求めるメッセージを表示します。手動でインストールすることもできます。
xcode-select --install
開発する
開発には macOS 26 以降と Xcode コマンド ライン ツールが必要です。 Xcode自体は必要ありません。
スクリプト/build.sh
開発アプリをビルドして実行します。
script/build.sh の実行
テストを実行します。
スクリプト/test.sh
SwiftPM とスクリプトの出力をクリーンアップします。
スクリプト/clean.sh
Config/.env.example を Config/.env にコピーし、IRONSMITH_DEV_SIGN_IDENTITY に Apple 開発 ID を入力して、新しいビルドの実行時にキーチェーンの質問が繰り返されるのを回避します。
問題

およびプルリクエストは大歓迎です。ローカルのワークフローと PR の期待については、CONTRIBUTING.md から始めます。
Ironsmith は、GNU General Public License v3.0 に基づいてライセンスされています。
プロンプトを表示して個人用 Mac アプリを即座に作成します。オンデバイスおよびクラウド LLM をサポート
読み込み中にエラーが発生しました。このページをリロードしてください。
16
フォーク
レポートリポジトリ
リリース
2
v0.1.1
最新の
2026 年 6 月 15 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Create personal Mac apps instantly with a prompt. Supports on-device and cloud LLMs - Jeidoban/Ironsmith

GitHub - Jeidoban/Ironsmith: Create personal Mac apps instantly with a prompt. Supports on-device and cloud LLMs · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
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
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Jeidoban
/
Ironsmith
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits .github/ workflows .github/ workflows .vscode .vscode Config Config Ironsmith Ironsmith IronsmithTests IronsmithTests assets/ readme assets/ readme script script .gitignore .gitignore AGENTS.md AGENTS.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Package.resolved Package.resolved Package.swift Package.swift README.md README.md View all files Repository files navigation
Ironsmith is a free, open-source macOS menu bar app for making small, personal Mac apps with AI. Describe what you want, and Ironsmith generates, builds, and saves a native SwiftUI app you can launch, edit, and export to your Applications folder.
Builds real Mac apps. Generated apps are Swift and SwiftUI apps, no Electron to be found.
Lives in your menu bar. Makes it very convenient to create a new app, run saved apps, edit an existing app, restore a previous version, view the code, or export a finished app bundle.
Works with local AI. Ironsmith was designed with local AI support in mind, and has Ollama support out of the box. You can also connect any OpenAI compatible API, so LM Studio and Llama.cpp work great too. You can even build very simple apps with Apple's built in Foundation Model.
Supports hosted models too. Bring your own API keys for OpenAI, Anthropic, and Gemini, or skip the API key and sign into Ironsmith to access them immediately.
Doesn't require Xcode. Every generated app is a Swift package and is built entirely with the lightweight Xcode command line tools rather than full Xcode. In fact Ironsmith itself doesn't even use Xcode!
Sandboxes every app by default. Generated apps are built as signed app bundles with sandboxing and hardened runtime enabled, greatly reducing the impact of bugs, mistakes, or malicious behavior. Sensitive permissions such as camera and microphone access must also be explicitly enabled. However, you can disable these protections, and if you do, it’s highly recommended that you review the code before running it.
Ironsmith works best for focused utilities: the small apps you wish existed but wouldn't want to hunt down or build yourself. That said, with more capable models like GPT‑5.5 or Opus 4.8, you can create some surprisingly sophisticated apps.
Some examples of prompts you can try:
"Make a utility that renames a folder of screenshots by date and window title."
"Build a tiny app that splits a PDF into one file per page."
"Build a clipboard cleaner that strips tracking parameters from copied URLs."
"Make a small CSV inspector that highlights duplicate rows and missing values."
Download the latest Ironsmith build from GitHub Releases or the website .
Ironsmith requires macOS 26 or newer and supports both Intel and Apple Silicon Macs. Make sure Apple Intelligence is enabled where available; Ironsmith uses it to create app icons and provide the built-in Foundation Model.
On first launch, Ironsmith checks for the Xcode Command Line Tools as generated apps are compiled locally. If they are missing, macOS will prompt you to install them. You can also install them manually:
xcode-select --install
Develop
Development requires macOS 26 or newer and the Xcode Command Line Tools. Xcode itself is not required.
script/build.sh
Build and run the development app:
script/build.sh run
Run tests:
script/test.sh
Clean SwiftPM and script outputs:
script/clean.sh
Copy Config/.env.example to Config/.env and fill in IRONSMITH_DEV_SIGN_IDENTITY with your Apple Development ID to avoid repeated keychain asks when running new builds.
Issues and pull requests are welcome. Start with CONTRIBUTING.md for the local workflow and PR expectations.
Ironsmith is licensed under the GNU General Public License v3.0 .
Create personal Mac apps instantly with a prompt. Supports on-device and cloud LLMs
There was an error while loading. Please reload this page .
16
forks
Report repository
Releases
2
v0.1.1
Latest
Jun 15, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
