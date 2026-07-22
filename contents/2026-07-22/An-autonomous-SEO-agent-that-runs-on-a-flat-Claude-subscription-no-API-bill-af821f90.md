---
source: "https://github.com/DigiHold/seo-agent-pack"
hn_url: "https://news.ycombinator.com/item?id=49005225"
title: "An autonomous SEO agent that runs on a flat Claude subscription (no API bill)"
article_title: "GitHub - DigiHold/seo-agent-pack · GitHub"
author: "NicoAI"
captured_at: "2026-07-22T12:23:03Z"
capture_tool: "hn-digest"
hn_id: 49005225
score: 1
comments: 0
posted_at: "2026-07-22T11:47:31Z"
tags:
  - hacker-news
  - translated
---

# An autonomous SEO agent that runs on a flat Claude subscription (no API bill)

- HN: [49005225](https://news.ycombinator.com/item?id=49005225)
- Source: [github.com](https://github.com/DigiHold/seo-agent-pack)
- Score: 1
- Comments: 0
- Posted: 2026-07-22T11:47:31Z

## Translation

タイトル: フラットな Claude サブスクリプションで実行される自律型 SEO エージェント (API 請求なし)
記事タイトル: GitHub - DigiHold/seo-agent-pack · GitHub
説明: GitHub でアカウントを作成して、DigiHold/seo-agent-pack の開発に貢献します。

記事本文:
GitHub - DigiHold/seo-agent-pack · GitHub
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
デジホールド
/
seo エージェント パック
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メイン ブランチ タグ 移動先

ファイル コード [その他のアクション] メニューを開く フォルダーとファイル
6 コミット 6 コミット スキル/SEO エンジン スキル/SEO エンジン テンプレート テンプレート README.md README.md Banner.png Banner.png すべてのファイルを表示 リポジトリ ファイルのナビゲーション
通常のクロードのサブスクリプションに基づいて、毎晩私のサイトの記事とページの下書きを行う自律的な SEO エージェント。このリポジトリは、私自身のサイト名を削除した、実行されるエンジンそのものです。これは、GitHub 用に私が作成したデモやテンプレートではありません。これらは私の作業ファイルであり、2 か月間の修正が含まれています。
毎晩スケジュールされたクロード コードの実行では、実際の検索データから 1 つのキーワードが選択され、SERP のライブ トップと照合され、厳格なアンチ スロップ ルールブックに基づいて記事が作成され、画像が生成され、人間が公開できるよう下書きが保存されます。日曜日に、Google Search Consoleを再読み込みし、パフォーマンスが低下しているものをすべて書き換えます。これは私の Mac 上で動作し、Claude サブスクリプションの定額料金で、書き込みに対するトークンごとの API 請求は不要で、毎朝 Telegram でレポートをくれます。
過去 2 か月間で、4 つのサイトにわたって 70 の記事と 65 ページの下書きを作成しました。公開前にすべての作品をレビューします。
これが AI スパム マシンではない理由
それが最初の反論ですので、ここで前もって述べておきます。ルールブック全体は、エージェントが一般的なコンテンツを送信するのを阻止するために存在します。
キーワードは DataForSEO と Search Console からのみ取得され、決して発明されたものではありません。検証済みの検索ボリュームがゼロということは、記事がないことを意味します。
書く前に、エージェントは現在のトップ 3 ページを読み、そのページがカバーしている内容をマッピングし、欠けているものをリストし、そのバージョンがそのバージョンを上回るに値する理由について一文を書かなければなりません。その文を書けない場合は、トピックを削除します。
禁止パターン パスとヒューマナイザー パス (ウィキペディアの「AI ライティングの兆候」ページから構築された 2 つ目) は、すべてのドラフトで実行されます。
すべての価格、日付、番号が確認されます

実行中にライブソースに対して d を実行しないと、カットされてしまいます。
単独で公開するものはありません。すべての作品は、人間が承認するまでは草案のままです。
薄いサイトを一夜にして権威あるサイトに変えることはできませんし、最初の草案は依然としてユーザーの声に合わせて調整する必要があります。ほとんどの自律型セットアップが生み出すスロップは出荷されません。
Skills/seo-engine/ には、5 つのゲートによるキーワード選択、競合他社のパス、Search Console からのクイック ウィン ハンティング、ライティング ルール、500 ワードの画像プロンプト メソッド、夜間実行プロトコルなどのメソッドが含まれています。
templates/master-setup-prompt.md はクロード コードに貼り付けるプロンプトの 1 つで、独自のエージェントがシステム全体を構築し、サイト構成について面接し、ヘルパー スクリプトを作成し、レビューのために停止します。
templates/ には、サイト構成テンプレート、夜間の実行プロンプト、4 つのシークレット env.example 、および cron および launchd スケジュールも保持されます。
開始方法 (約 1 時間、コーディングは必要ありません)
これは手作業で構築するものではありません。プロンプトを 1 つ貼り付けると、独自の Claude Code エージェントがこのリポジトリ内のファイルからプロンプトを構築します。
Anthropic のサイトから Claude Code をインストールし、Claude サブスクリプションでログインします。 Pro プランは 1 つのサイトに十分です。
4 つのアカウントを作成してキーを取得します: DataForSEO アカウント (キーワード データ)、サイトに接続された Google Search Console (無料)、Gemini API キー (画像)、BotFather 経由の Telegram ボット (無料の朝レポート)。
git clone https://github.com/DigiHold/seo-agent-pack を使用してこのリポジトリをダウンロードし、そのフォルダー内の Claude Code を開きます。
template/master-setup-prompt.md の内容全体をコピーし、クロード コードに貼り付けます。エージェントはメソッド ファイルを読み取り、エンジンをインストールし、サイトについてインタビューし、独自のヘルパー スクリプトを作成して、4 つのキーをローカルの .env ファイルに貼り付けるように求めます。あなたを離れるものは何もありません

rマシン。
完了すると、1 つのテスト記事が実行され、下書きが表示されます。それを読んで、ルール ファイル内の気に入らない点を修正してから、夜間スケジュールを有効にするようにエージェントに指示します。
その夜から、それは自動的に実行され、Telegram で毎朝 1 つのドラフトをレビューします。セットアップで唯一遅いのは、Google 側の Search Console API アクセスです。エージェントが 15 分ほどで段階的に説明します。
エージェントが書いたすべての記事は下書きとして保存されるため、人間は常に公開を押します。エージェントは、 .env から何も印刷、コピー、コミットすることはありません。各実行にはハード タイムアウトがあり、1 つのステータス行が日付付きのログに書き込まれます。資格情報の範囲は、各ジョブに必要な最小限に限定されます。これらすべてを所定の位置に保管してください。
構造は初日からしっかりしていますが、最初のドラフトはまだサイトの声に合わせて調整する必要があります。エンジン ファイルに書き込む修正はすべて永続的なものであるため、システムは高速にシャープになります。
すぐに成果を得るには実際の Search Console の履歴が必要なので、データのない新しいサイトでは、最初はあまり活用できません。
1 つのサイトのトピック クラスターを中心に最適化します。無関係な被写体に向けると品質が低下します。
これは MIT のライセンスを受け、Nicolas Lecocq によって公開されています。もしそれを基にしたら、どうなるのか聞きたいです。
読み込み中にエラーが発生しました。このページをリロードしてください。
8
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to DigiHold/seo-agent-pack development by creating an account on GitHub.

GitHub - DigiHold/seo-agent-pack · GitHub
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
DigiHold
/
seo-agent-pack
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits skills/ seo-engine skills/ seo-engine templates templates README.md README.md banner.png banner.png View all files Repository files navigation
An autonomous SEO agent that drafts articles and pages for my sites every night, on a normal Claude subscription. This repo is the exact engine it runs on, cleaned of my own site names. It is not a demo or a template I wrote for GitHub. These are my working files, two months of corrections included.
Every night a scheduled Claude Code run picks one keyword from real search data, checks it against the live top of the SERP, writes the piece under a strict anti-slop rulebook, generates its images, and saves a draft for a human to publish. On Sunday it re-reads Google Search Console and rewrites whatever underperforms. It runs on my Mac, costs the flat price of a Claude subscription with no per-token API bill for the writing, and reports to me on Telegram every morning.
Over the last two months it has drafted 70 articles and 65 pages across 4 sites. I review every piece before anything goes live.
Why this is not an AI-spam machine
That is the first objection, so here it is up front. The whole rulebook exists to stop the agent from shipping generic content.
Keywords come only from DataForSEO and Search Console, never invented. Zero verified search volume means no article.
Before writing, the agent reads the current top 3 pages, maps what they cover, lists what they miss, and has to write one sentence on why its version deserves to beat them. If it cannot write that sentence, it drops the topic.
A banned-pattern pass and a humanizer pass, the second built from Wikipedia's "Signs of AI writing" page, run on every single draft.
Every price, date and number is verified against a live source during the run, or it gets cut.
Nothing publishes on its own. Every piece stays a draft until a human approves it.
It will not turn a thin site into an authority overnight, and the first drafts still need tuning to your voice. It just will not ship the slop most autonomous setups produce.
skills/seo-engine/ holds the method: keyword selection through 5 gates, the competitor pass, quick-win hunting from Search Console, the writing rules, the 500-word image-prompt method, and the nightly run protocol.
templates/master-setup-prompt.md is one prompt you paste into Claude Code so your own agent builds the whole system, interviews you for your site config, writes the helper scripts, and stops for your review.
templates/ also holds the site-config template, the nightly run prompt, the four-secret env.example , and the cron and launchd schedule.
How to start (about an hour, no coding required)
You do not build this by hand. You paste one prompt and your own Claude Code agent builds it for you, from the files in this repo.
Install Claude Code from Anthropic's site and log in with a Claude subscription. The Pro plan is enough for one site.
Create four accounts and grab the keys: a DataForSEO account (keyword data), Google Search Console connected to your site (free), a Gemini API key (images), and a Telegram bot through BotFather (free morning reports).
Download this repo with git clone https://github.com/DigiHold/seo-agent-pack , then open Claude Code inside that folder.
Copy the entire contents of templates/master-setup-prompt.md and paste it into Claude Code. The agent reads the method files, installs the engine, interviews you about your site, writes its own helper scripts, and asks you to paste your four keys into a local .env file. Nothing leaves your machine.
When it finishes, it runs one test article and shows you the draft. Read it, fix anything you dislike in the rule files, then tell the agent to enable the nightly schedule.
From that night on it runs by itself, and you review one draft per morning on Telegram. The only slow part of setup is Search Console API access on Google's side, and the agent walks you through it step by step in about 15 minutes.
Every piece the agent writes saves as a draft, so a human always presses publish. The agent never prints, copies or commits anything from .env . Each run has a hard timeout and writes one status line to a dated log. Credentials are scoped to the minimum each job needs. Keep all of this in place.
The structure is solid from day one, but the first drafts still need tuning to your site's voice. The system sharpens fast because every correction you write into the engine files is permanent.
It needs real Search Console history to find quick wins, so a brand-new site with no data gets less out of it at the start.
It optimizes around one site's topic clusters. Point it at unrelated subjects and the quality drops.
It is MIT licensed and built in public by Nicolas Lecocq. If you build on it, I would like to hear how it goes.
There was an error while loading. Please reload this page .
8
forks
Report repository
Releases
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
