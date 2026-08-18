---
source: "https://github.com/physiclaw/PhysiClaw"
hn_url: "https://news.ycombinator.com/item?id=49345058"
title: "Show HN: PhysiClaw – an AI agent that physically operates an iPhone"
article_title: "GitHub - physiclaw/PhysiClaw: The AI agent that interacts with you in the real world. · GitHub"
image: "https://opengraph.githubassets.com/a83603e015460cfc40e6a15c8893549a59e842976b54e6a5ca87271ae1635ac7/physiclaw/PhysiClaw"
author: "qiaoqian"
captured_at: "2026-08-18T13:36:20Z"
capture_tool: "hn-digest"
hn_id: 49345058
score: 1
comments: 0
posted_at: "2026-08-18T13:05:32Z"
tags:
  - hacker-news
  - translated
---

# Show HN: PhysiClaw – an AI agent that physically operates an iPhone

- HN: [49345058](https://news.ycombinator.com/item?id=49345058)
- Source: [github.com](https://github.com/physiclaw/PhysiClaw)
- Score: 1
- Comments: 0
- Posted: 2026-08-18T13:05:32Z

## Translation

タイトル: Show HN: PhysiClaw – iPhone を物理的に操作する AI エージェント
記事のタイトル: GitHub - physiclaw/PhysiClaw: 現実世界で対話する AI エージェント。 · GitHub
説明: 現実世界であなたと対話する AI エージェント。 - フィジクロー/フィジクロー

記事本文:
GitHub - physiclaw/PhysiClaw: 現実世界であなたと対話する AI エージェント。 · GitHub
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
物理法則
/
フィジクロー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
1,640 Commits 1,640 Commits Folders and files
.github/ workflows .github/ workflows docs docs ハードウェア ハードウェア ライセンス ライセンス スクリプト スクリプト スキル スキル src/ physiclaw src/ physiclaw テスト テスト .gitignore .gitign

ore .markdownlint.json .markdownlint.json .python-version .python-version LICENSE LICENSE Makefile Makefile README.md README.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md install.ps1 install.ps1 install.sh install.sh pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
あなたと同じように、電話を物理的に操作する AI エージェント。
PhysiClaw はカメラで携帯電話の画面を監視し、スタイラスでタップします。
人間と同じように電話を操作します。 API、OAuth、ADB ケーブルなし、
電話機には何もインストールされていません。ロックを解除して机の上に置き、そのままにしておきます。
エージェントの仕事。
テイクアウトの注文や買い物など、毎日の用事が山積みになるように設計されています。
食料品の購入、配車の予約、請求書の支払い、メッセージへの返信。何でも
携帯電話で手動で行うことも、PhysiClaw が代わりに行うこともできます。
日常生活を実行するアプリは閉鎖されており、ほとんどのアプリはパブリック API を公開していません。
そしてシミュレートされた入力 (デスクトップ オートメーションまたは Android の ADB) はソフトウェアを離れます
アンチボット システムがフラグを立てる指紋。つまり、PhysiClaw は画面自体を扱います
API として: カメラがそれを読み取り、スタイラスがジェスチャを実行します。電話へ
それは本物の指と区別がつきません - 何も検出できません - そしてそれは到達します
ほぼすべてのアプリを確実に利用できます。
トレードオフはスピードです。汎用性と引き換えに、アクションあたり数秒です。
そして信頼性。
PhysiClaw には、独自の WhatsApp アカウントを実行する専用の電話機があります。
連絡先として追加し、実際の人間のようにチャットします - 必要なことを伝えます
平易な言葉で言うと：
DoorDash でラテを注文します。
明日の天気はどうですか？
インスタカートで牛乳と卵を購入します。
ランタイムは継続的にループし、スケジュールされたタスクの期限が来るとエージェントを起動します。
または、画面が点灯して新しいメッセージが表示されます。エージェントは電話のロックを解除し、読み取ります
あなたのメッセージ、そして、

彼のタスク — 重要な場合は確認のために一時停止し、
支払いと同様に、結果を返信し、メモリを保存して終了します。
次に目覚めるまで。
PhysiClaw ハードウェア (アーム、スタイラス、カメラ) に加えて、
それが動作するためのiPhone。
physiclaw CLI は、macOS、Windows、および Linux で実行されます。
# CLI をインストールします (uv + Python 3.12 + physiclaw)
カール -fsSL https://physiclaw.ai/install.sh |バッシュ
# Windows: iwr -useb https://physiclaw.ai/install.ps1 |アイエックス
physiclaw models key <プロバイダー> # LLM API キーを追加します
physiclaw models use < Provider/model > # エージェントが実行されるモデルを選択します
物理法医 # 環境を確認してください
physiclaw # サーバーを起動します — ハードウェア セットアップ ウィザードを開き、エージェントを実行します
ライセンス
MIT — CAD-as-Code ハードウェアとエージェント ソフトウェアの両方。
現実世界であなたと対話する AI エージェント。
Readme MIT ライセンス アクティビティ スター
35 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

The AI agent that interacts with you in the real world. - physiclaw/PhysiClaw

GitHub - physiclaw/PhysiClaw: The AI agent that interacts with you in the real world. · GitHub
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
physiclaw
/
PhysiClaw
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
1,640 Commits 1,640 Commits Folders and files
.github/ workflows .github/ workflows docs docs hardware hardware licenses licenses scripts scripts skills skills src/ physiclaw src/ physiclaw tests tests .gitignore .gitignore .markdownlint.json .markdownlint.json .python-version .python-version LICENSE LICENSE Makefile Makefile README.md README.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md install.ps1 install.ps1 install.sh install.sh pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
An AI agent that physically operates a phone — the way you do.
PhysiClaw watches a phone's screen with a camera and taps it with a stylus,
working the phone the way a person would. No APIs, no OAuth, no ADB cables,
nothing installed on the phone — just unlock it, set it on the desk, and let
the agent work.
It's built for the everyday errands that pile up — ordering takeout, shopping
for groceries, booking a ride, paying a bill, replying to a message. Anything
you can do by hand on your phone, PhysiClaw can do for you.
The apps that run your daily life are closed off: most expose no public API,
and simulated input — desktop automation or Android's ADB — leaves software
fingerprints that anti-bot systems flag. So PhysiClaw treats the screen itself
as the API: a camera reads it, a stylus performs the gestures. To the phone
it's indistinguishable from a real finger — nothing to detect — and it reaches
virtually any app reliably.
The trade-off is speed: a few seconds per action, in exchange for universality
and reliability.
PhysiClaw has its own dedicated phone running its own WhatsApp account.
Add it as a contact and chat with it like a real person — tell it what you need
in plain language:
Order a latte on DoorDash.
What's the weather tomorrow?
Buy milk and eggs on Instacart.
Its runtime loops continuously, waking the agent when a scheduled task is due
or the screen lights up with a new message. The agent unlocks the phone, reads
your message, and does the task — pausing for confirmation when it matters,
like a payment — then replies with the result, saves its memory, and exits
until it's next woken.
You'll need the PhysiClaw hardware — the arm, stylus, and camera — plus an
iPhone for it to operate.
The physiclaw CLI runs on macOS, Windows, and Linux:
# Install the CLI (uv + Python 3.12 + physiclaw)
curl -fsSL https://physiclaw.ai/install.sh | bash
# Windows: iwr -useb https://physiclaw.ai/install.ps1 | iex
physiclaw models key < provider > # add your LLM API key
physiclaw models use < provider/model > # pick the model the agent runs on
physiclaw doctor # check your environment
physiclaw # start the server — opens the hardware-setup wizard, then runs the agent
License
MIT — both the CAD-as-code hardware and the agent software.
The AI agent that interacts with you in the real world.
Readme MIT license Activity Stars
35 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
