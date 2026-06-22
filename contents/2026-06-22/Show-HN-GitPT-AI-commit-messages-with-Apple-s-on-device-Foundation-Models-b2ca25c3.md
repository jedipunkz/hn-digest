---
source: "https://github.com/bartaxyz/GitPT"
hn_url: "https://news.ycombinator.com/item?id=48631997"
title: "Show HN: GitPT – AI commit messages with Apple's on-device Foundation Models"
article_title: "GitHub - bartaxyz/GitPT: Git Prompt Tool - AI commit messages, built for on-device models (incl. Apple Foundation Models) · GitHub"
author: "bartaxyz"
captured_at: "2026-06-22T16:33:50Z"
capture_tool: "hn-digest"
hn_id: 48631997
score: 2
comments: 0
posted_at: "2026-06-22T15:58:24Z"
tags:
  - hacker-news
  - translated
---

# Show HN: GitPT – AI commit messages with Apple's on-device Foundation Models

- HN: [48631997](https://news.ycombinator.com/item?id=48631997)
- Source: [github.com](https://github.com/bartaxyz/GitPT)
- Score: 2
- Comments: 0
- Posted: 2026-06-22T15:58:24Z

## Translation

タイトル: Show HN: GitPT – Apple のオンデバイス基盤モデルを使用した AI コミット メッセージ
記事のタイトル: GitHub - bartaxyz/GitPT: Git Prompt Tool - AI コミット メッセージ、オンデバイス モデル (Apple Foundation モデルを含む) 用に構築 · GitHub
説明: Git プロンプト ツール - AI コミット メッセージ、オンデバイス モデル (Apple Foundation モデルを含む) 用に構築 - bartaxyz/GitPT

記事本文:
GitHub - bartaxyz/GitPT: Git プロンプト ツール - オンデバイス モデル (Apple Foundation モデルを含む) 用に構築された AI コミット メッセージ · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
バルタキシズ
/
GitPT
公共
通知
通知設定を変更するにはサインインする必要があります
追加ナビ

オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
51 コミット 51 コミット .github/ workflows .github/ workflows アセット アセット スクリプト スクリプト src src テスト テスト .gitignore .gitignore .npmignore .npmignore .releaserc.json .releaserc.json CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md LICENSE LICENSE README.md README.md commitlint.config.js commitlint.config.js package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
GitPT (Git Prompt Tool) は、オンデバイス モデル用に構築された LLM を使用してコミット メッセージを書き込む git コマンド エイリアスです。 git を使用する場所であればどこでも使用できます。変更されないすべてのコマンドはそのまま通過するため、git を git に置き換えても何も失われません。変更されるコマンドは commit の 1 つで、メッセージを書き込みます。
大規模なリモート モデルはすでにコミットを正常に書き込んでいます。より困難なケースは、数千のトークンしか保持しない、マシン上の小さなローカル モデル (Apple の Foundation Models、Ollama、LM Studio) です。実際の差分ははるかに大きいため、GitPT はモデルのコンテキスト ウィンドウに適合するまでファイルごとに要約し、その要約からメッセージを書き込みます。リモート API (OpenAI、Anthropic、OpenRouter) も機能します。
diff は、メッセージが書き込まれる前に、モデルのコンテキスト ウィンドウに適合するように要約されます。
新機能: Apple Foundation モデル。 macOS 27 以降では、GitPT は Apple の組み込みオンデバイス モデルで実行されます。 API キーもサインアップも必要なく、Mac からは何も離れません。 gitpt setup を実行し、Apple Foundation Models を選択します。
gitpt setup # モデルを選択する
オプション: エイリアス git を使用すると、すでに git と入力した場所で GitPT が実行されます。
プレーン git では、メッセージを自分で書きます。
git add 。
git commit -m " fix: handle empty staged diff " # メッセージを書き込みます

GitPT を使用すると、次のことができなくなります。
gitpt 追加 。
gitpt commit # GitPT が diff から書き込みます
gitpt commit は、ステージングされた diff を読み取り、メッセージを書き込み、エディターで開きます。保存して閉じてコミットします。
gitpt commit : 段階的な変更からのコミット メッセージを書き込みます。 -m と commitlint ルールを尊重します。
gitpt model : モデルを選択または切り替えます。各プロバイダーは独自のキーを保持します。
gitpt setup / gitpt config / gitpt restart : 設定を構成、表示、またはクリアします。
gitpt pr create : gh CLI (実験的) を使用してプル リクエストのタイトルと説明を作成します。
オンデバイス: macOS 27 以降の Apple の Foundation モデル (API キーなし)、または Ollama や LM Studio などの OpenAI 互換ローカル サーバー。
リモート : OpenAI、Anthropic、OpenRouter。 APIキーを持参してください。
機能が必要ですか、それともバグに遭遇しますか?問題を開きます。
Git プロンプト ツール - オンデバイス モデル (Apple Foundation モデルを含む) 用に構築された AI コミット メッセージ
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
8
v1.6.1
最新の
2026 年 6 月 18 日
+ 7 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Git Prompt Tool - AI commit messages, built for on-device models (incl. Apple Foundation Models) - bartaxyz/GitPT

GitHub - bartaxyz/GitPT: Git Prompt Tool - AI commit messages, built for on-device models (incl. Apple Foundation Models) · GitHub
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
bartaxyz
/
GitPT
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
51 Commits 51 Commits .github/ workflows .github/ workflows assets assets scripts scripts src src tests tests .gitignore .gitignore .npmignore .npmignore .releaserc.json .releaserc.json CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md commitlint.config.js commitlint.config.js package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
GitPT (Git Prompt Tool) is a git command alias that writes your commit messages with an LLM, built for on-device models. Use it anywhere you use git: every command it doesn't change passes straight through, so you can replace git with it and lose nothing. The one command that changes is commit , which writes the message for you.
Big remote models already write commits fine. The harder case is the small local model on your machine (Apple's Foundation Models, Ollama, LM Studio), which holds only a few thousand tokens. A real diff is much larger, so GitPT summarizes it file by file until it fits the model's context window, then writes the message from that summary. Remote APIs (OpenAI, Anthropic, OpenRouter) work too.
The diff is summarized to fit the model's context window before the message is written.
New: Apple Foundation Models. On macOS 27 and later, GitPT runs on Apple's built-in on-device model. No API key, no signup, nothing leaves your Mac. Run gitpt setup and choose Apple Foundation Models.
gitpt setup # pick a model
Optional: alias git so GitPT runs everywhere you already type git.
With plain git, you write the message yourself:
git add .
git commit -m " fix: handle empty staged diff " # you write the message
With GitPT, you don't:
gitpt add .
gitpt commit # GitPT writes it from the diff
gitpt commit reads your staged diff, writes the message, and opens it in your editor. Save and close to commit.
gitpt commit : write a commit message from staged changes. Respects -m and your commitlint rules.
gitpt model : pick or switch the model. Each provider keeps its own key.
gitpt setup / gitpt config / gitpt reset : configure, show, or clear settings.
gitpt pr create : draft a pull request title and description with the gh CLI (experimental).
On-device : Apple's Foundation Models on macOS 27 or later (no API key), or any OpenAI-compatible local server such as Ollama or LM Studio.
Remote : OpenAI, Anthropic, OpenRouter. Bring an API key.
Want a feature or hit a bug? Open an issue .
Git Prompt Tool - AI commit messages, built for on-device models (incl. Apple Foundation Models)
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
8
v1.6.1
Latest
Jun 18, 2026
+ 7 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
