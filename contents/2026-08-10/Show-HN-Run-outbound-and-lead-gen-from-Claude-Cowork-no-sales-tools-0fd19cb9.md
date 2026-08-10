---
source: "https://github.com/skyzer/skills"
hn_url: "https://news.ycombinator.com/item?id=49246184"
title: "Show HN: Run outbound and lead gen from Claude Cowork, no sales tools"
article_title: "GitHub - skyzer/skills: Agent skills that get a company its first users and first revenue. Point your agent at this repo to install. · GitHub"
author: "skyzer"
captured_at: "2026-08-10T17:45:05Z"
capture_tool: "hn-digest"
hn_id: 49246184
score: 1
comments: 0
posted_at: "2026-08-10T16:43:46Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Run outbound and lead gen from Claude Cowork, no sales tools

- HN: [49246184](https://news.ycombinator.com/item?id=49246184)
- Source: [github.com](https://github.com/skyzer/skills)
- Score: 1
- Comments: 0
- Posted: 2026-08-10T16:43:46Z

## Translation

タイトル: HN を表示: クロード・コワークからアウトバウンドとリード獲得を実行、営業ツールなし
記事のタイトル: GitHub - skyzer/skills: 企業に最初のユーザーと最初の収益をもたらすエージェントのスキル。エージェントにこのリポジトリを指定してインストールしてください。 · GitHub
説明: 企業に最初のユーザーと最初の収益をもたらすエージェント スキル。エージェントにこのリポジトリを指定してインストールしてください。 - スカイザー/スキル

記事本文:
GitHub - skyzer/skills: 企業に最初のユーザーと最初の収益をもたらすエージェント スキル。エージェントにこのリポジトリを指定してインストールしてください。 · GitHub
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
スカイザー
/
スキル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット outbound-master outbound-master .gitignore .gitignore ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
エージェントのスキル

企業は最初のユーザーと最初の収益を得ることができます。
各スキルは、小規模企業がまだ雇用できない専門家に代わるものであり、それぞれのスキルは、指示、スクリプト、設定例、実際の出力例など、独自のフォルダーに自己完結しています。これらは、Claude Code、Claude Cowork、Codex、またはファイルを読み取ってシェルを実行するその他のもので実行されます。状態はベンダーではなく、ユーザーが所有するファイルに存在します。
エージェントにこのリポジトリを指定し、スキルに名前を付けます。
github.com/skyzer/skills から outbound-master スキルをインストールして設定します。
これでインストールは完了です。エージェントはリポジトリのクローンを作成し、スキルの install.sh を実行し、依存関係を取り込んで、スキル自体のドキュメントに記載されている設定を順を追って説明します。エージェントが知る必要があるものはすべて、各スキルのフォルダーにあります。ここでは何も人間が端末にいるということを想定していません。
git clone https://github.com/skyzer/skills
CD スキル/アウトバウンドマスター
./install.sh
pip3 install -r 要件.txt
スキル
スキル
機能
何をするのか
アウトバウンドマスター
販売
B2B コールド アウトバウンドをエンドツーエンドで実行します。ソース、スコア付け、日付付きフックの調査、コピーの書き込みとデロップ、すべてのアドレスの検証、間隔をあけたスケジュールで冪等に送信、追加専用のイベント ログですべてのタッチを追跡し、人間が承認する返信の下書きを実行します。送信すべきでない場合は送信を拒否します。
さらに登場します。新しいスキルはそれぞれ、同じ形状の独自のフォルダーとして表示されます。
あなたがこのリポジトリからスキルをインストールまたは実行するよう求められたエージェントの場合:
リポジトリのクローンを作成し、スキルのフォルダーに cd します。
スキルの README.md を読み、次に skill/SKILL.md を読みます。 SKILL.md は操作マニュアルです。事前にすべてのステップを参照するのではなく、現在実行しているステップの参照ファイルを読み取ります。
./install.sh を実行してから、requirements.txt をインストールします。
オペレーターに彼らだけが持っているもの (資格情報、ビジネス概要) を尋ね、スキルのプリフライトを実行する前に実行します。

本当のことは何でもやっている。
スキルのガードを尊重します。スキルに予行モードが同梱されている場合は、それが開始されます。
ここでのすべてのスキルに従うデザインルール
設定はエンジンとは別のものです。あなたの会社に関するすべての情報は config/ に保存されており、gitignord されます。スキルには、会社、市場、製品の名前は含まれていません。
状態はチャットではなくファイルです。説明した会話がなくなっているため、次回の実行で必要なものはすべてディスク上に存在します。
ガードなしでは取り返しのつかないことは何も起こりません。送信は冪等であり、破壊的な行為は制限されており、お金、評判、関係を費やすものはすべて停止して尋ねます。
スキップするのは無料ですが、間違うのは無料ではありません。どのスキルも、もっともらしく間違ったことをするよりも、何もせずに理由を言うことを好みます。
まずはドライラン。スキルが外部の世界に触れた場合、そのスキルが何を行うかを正確に示し、何も行わないモードがあります。
企業に最初のユーザーと最初の収益をもたらすエージェントのスキル。エージェントにこのリポジトリを指定してインストールしてください。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Agent skills that get a company its first users and first revenue. Point your agent at this repo to install. - skyzer/skills

GitHub - skyzer/skills: Agent skills that get a company its first users and first revenue. Point your agent at this repo to install. · GitHub
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
skyzer
/
skills
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits outbound-master outbound-master .gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
Agent skills that get a company its first users and its first revenue.
Each skill replaces a specialist a small company can't afford to hire yet, and each is self-contained in its own folder: instructions, scripts, example config, and a real example of its output. They run on Claude Code, Claude Cowork, Codex, or anything else that reads files and runs a shell. State lives in files you own, not in a vendor.
Point your agent at this repo and name the skill:
Install the outbound-master skill from github.com/skyzer/skills and set it up.
That's the whole install. The agent clones the repo, runs the skill's install.sh , pulls in the dependencies, and walks you through the configuration the skill's own docs describe. Everything an agent needs to know is in each skill's folder; nothing here assumes a human at a terminal.
git clone https://github.com/skyzer/skills
cd skills/outbound-master
./install.sh
pip3 install -r requirements.txt
The skills
Skill
Function
What it does
outbound-master
Sales
Runs B2B cold outbound end to end: sources, scores, researches a dated hook, writes and deslops the copy, verifies every address, sends idempotently on a spaced schedule, tracks every touch in an append-only event log, and drafts replies a human approves. Refuses to send when it shouldn't.
More coming. Each new skill lands as its own folder with the same shape.
If you're an agent asked to install or run a skill from this repo:
Clone the repo and cd into the skill's folder.
Read the skill's README.md , then skill/SKILL.md . The SKILL.md is the operating manual; read the reference file for the step you're on rather than all of them up front.
Run ./install.sh , then install requirements.txt .
Ask the operator for what only they have (credentials, the business brief), and run the skill's preflight before doing anything real.
Respect the skill's guards. If a skill ships a dry-run mode, it starts on.
Design rules every skill here follows
Configuration is separate from the engine. Everything about your company lives in config/ , which is gitignored. Nothing in a skill names a company, market or product.
State is files, not chat. Anything needed on the next run exists on disk, because the conversation where you explained it is gone.
Nothing irreversible happens without a guard. Sends are idempotent, destructive actions are gated, and anything that spends money, reputation or a relationship stops and asks.
Skipping is free, being wrong isn't. Every skill would rather do nothing and say why than do something plausible and wrong.
Dry run first. If a skill touches the outside world, it has a mode that shows exactly what it would do and does nothing.
Agent skills that get a company its first users and first revenue. Point your agent at this repo to install.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
