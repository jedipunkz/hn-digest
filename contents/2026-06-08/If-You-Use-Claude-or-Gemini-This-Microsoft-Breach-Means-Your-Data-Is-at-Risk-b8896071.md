---
source: "https://scienspire.com/article/microsoft-breach-ai-code-assistants"
hn_url: "https://news.ycombinator.com/item?id=48451014"
title: "If You Use Claude or Gemini, This Microsoft Breach Means Your Data Is at Risk"
article_title: "If You Use Claude or Gemini, This Microsoft Breach Means Your Data Is at Risk — Scienspire"
author: "valsurier"
captured_at: "2026-06-08T21:38:26Z"
capture_tool: "hn-digest"
hn_id: 48451014
score: 1
comments: 0
posted_at: "2026-06-08T20:01:56Z"
tags:
  - hacker-news
  - translated
---

# If You Use Claude or Gemini, This Microsoft Breach Means Your Data Is at Risk

- HN: [48451014](https://news.ycombinator.com/item?id=48451014)
- Source: [scienspire.com](https://scienspire.com/article/microsoft-breach-ai-code-assistants)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T20:01:56Z

## Translation

タイトル: Claude または Gemini を使用している場合、この Microsoft 侵害はデータが危険にさらされていることを意味します
記事のタイトル: Claude または Gemini を使用している場合、この Microsoft 侵害はデータが危険にさらされていることを意味します — Scienspire
説明: Microsoft 所有の GitHub リポジトリに対する Miasma ワーム サプライ チェーン攻撃は、Claude Code、Gemini CLI、Cursor、VS Code などの AI コーディング アシスタントを標的としています。資格情報を保護し、環境を保護する方法を学びます。

記事本文:
Claude または Gemini を使用している場合、この Microsoft 侵害はデータが危険にさらされていることを意味します — Scienspire
サイエンスパイア。ソフトウェア AI 環境 ビジネスの高揚 購読 ホーム / セキュリティ / Claude または Gemini を使用している場合、この Microsoft 侵害はデータが危険にさらされていることを意味します セキュリティ · Microsoft Claude または Gemini を使用している場合、この Microsoft 侵害はお客様のデータが危険にさらされていることを意味します
Miasma ワームとして知られる高度なサプライ チェーン攻撃により Microsoft GitHub リポジトリが侵害され、Claude Code、Gemini CLI、Cursor、VS Code などの AI コーディング アシスタント内で爆発するように設計されたマルウェアが展開されました。
Miasma ワームとして知られる高度なサプライ チェーン攻撃は、Microsoft が所有する数十の GitHub リポジトリを侵害することに成功し、Claude Code、Gemini CLI、Cursor、VS Code などの AI コーディング アシスタント内で爆発するように特別に設計されたマルウェアを展開しました。
2026 年 6 月 5 日、悪意のある投稿者が自己複製型資格情報収集マルウェアを注入した後、GitHub は、Microsoft の 4 つの組織にわたる 73 のリポジトリ (Azure のコア インフラストラクチャを含む) を突然無効にすることを余儀なくされました。
AI エージェントを使用してコードを操作したり記述したりする場合、侵害について知っておくべきことと、環境を保護する方法を次に示します。
これまで開発者は、パッケージのインストール中 ( npm install の実行など) にマルウェアがライフサイクル スクリプトに隠れることを懸念していました。 Miasma ワームは、プロジェクト フォルダーを開くだけでペイロードが実行されるという、危険な新しいパラダイムを導入します。
攻撃者は、AI コーディング エージェントがプロジェクトを理解するために使用する構成ファイルを武器化することでこれを達成しました。このマルウェアは、悪意のあるコマンドを標準のセットアップ フック内に隠すことで、AI アシスタントを騙してペイロードを自動的に実行させます。
特定のツールをターゲットにする方法は次のとおりです。
クロード コードと Gemini CLI:

攻撃者は、悪意のある .claude/settings.json ファイルと .gemini/settings.json ファイルを埋め込みました。これらには、AI エージェントがリポジトリに接続した瞬間にマルウェアをサイレントに実行する「SessionStart」フックが含まれています。
Cursor: .cursor/rules/setup.mdc 内のプロンプト インジェクションは、Cursor AI を騙して、「プロジェクト環境を初期化する」ためにマルウェアを実行する必要があると信じ込ませます。
VS Code: 変更された .vscode/tasks.json ファイルは、フォルダーが開かれるとすぐにペイロードを自動実行します。
ペイロード自体は、攻撃的な資格情報の盗難という 1 つの目的のために構築された 4.6 MB の巨大な難読化された JavaScript ファイル ( .github/setup.js ) です。
AI エージェントまたは IDE によってトリガーされると、マルウェアはすぐに次のものを探します。
クラウド キー: AWS、Google Cloud Platform (GCP)、および Microsoft Azure の認証情報。
開発者シークレット: プロセス メモリから直接取得された GitHub Actions シークレット。
ローカル パスワード保管庫: 1Password や gopass などのパスワード マネージャーからのロック解除されたデータ。
インフラストラクチャ構成: .env ファイル、Docker 構成、および Kubernetes 環境に隠されているパスワード。
このペイロードは正規の OAuth トークンとクラウド キーを盗むため、攻撃者は従来のセキュリティ スキャナーをバイパスし、ワームが企業ネットワークを介して横方向に拡散し、信頼できる開発者 ID でさらに悪意のあるコードを公開することができます。
AI エージェントの観点からは、開発者によるオープンソース コードの扱い方の変化が求められています。 AI アシスタント内で信頼できない、または侵害されたリポジトリを開くと、ランダムな実行可能ファイルを実行する場合とまったく同じリスク プロファイルが生じるようになりました。
最近、Claude Code、Gemini CLI、または Cursor を使用して Microsoft または Azure リポジトリ (特に durtabletask フレームワーク周辺) のクローンを作成または操作した場合は、環境が危険にさらされている可能性があると想定してください。
不審な構成をチェックする: AI で外部リポジトリを開く前に

ツールを使用して、ルート フォルダーに予期しない .claude 、 .gemini 、 .cursor 、または .vscode ディレクトリがないか調べます。不明な JavaScript またはシェル ファイルを指している「SessionStart」フックを探します。
資格情報のローテーション: 漏洩の疑いがある場合は、GitHub パーソナル アクセス トークン (PAT)、SSH キー、CI/CD 署名キー、およびすべてのアクティブなクラウド プロバイダーの資格情報を直ちにローテーションします。
AI 権限の監査: AI コーディング エージェントが、手動承認なしに自動起動スクリプトを実行したり、機密性の高いローカル ディレクトリにアクセスしたりすることが明示的に制限されていることを確認します。
AIのコストはどれくらい？ GitHub Copilot ユーザーは、新しい使用量ベースの価格設定に反応します
Invideo を使用して収益性の高い顔の見えない YouTube チャンネルを作成する方法
ガンマ vs PowerPoint: AI プレゼンテーションが勝っている理由
次に来るものを形作るソフトウェア、科学、アイデアに関するシャープで楽観的なレポート。

## Original Extract

The Miasma worm supply chain attack on Microsoft-owned GitHub repositories targets AI coding assistants like Claude Code, Gemini CLI, Cursor, and VS Code. Learn how to protect your credentials and secure your environment.

If You Use Claude or Gemini, This Microsoft Breach Means Your Data Is at Risk — Scienspire
scienspire . Software AI Environment Business Uplifting Subscribe Home / Security / If You Use Claude or Gemini, This Microsoft Breach Means Your Data Is at Risk Security · Microsoft If You Use Claude or Gemini, This Microsoft Breach Means Your Data Is at Risk
A sophisticated supply chain attack known as the Miasma worm has compromised Microsoft GitHub repositories, deploying malware designed to detonate inside AI coding assistants like Claude Code, Gemini CLI, Cursor, and VS Code.
A highly sophisticated supply chain attack known as the Miasma worm has successfully compromised dozens of Microsoft-owned GitHub repositories, deploying malware specifically designed to detonate inside AI coding assistants like Claude Code, Gemini CLI, Cursor, and VS Code.
On June 5, 2026, GitHub was forced to abruptly disable 73 repositories across four Microsoft organizations—including core infrastructure for Azure—after a malicious contributor injected self-replicating credential-harvesting malware.
If you use AI agents to navigate or write code, here is what you need to know about the breach and how to secure your environment.
Historically, developers worried about malware hiding in lifecycle scripts during package installation (like running npm install ). The Miasma worm introduces a dangerous new paradigm: the payload executes simply by opening the project folder.
The attackers achieved this by weaponizing the configuration files that AI coding agents use to understand a project. By hiding malicious commands inside standard setup hooks, the malware tricks the AI assistant into running the payload automatically.
Here is how it targets specific tools:
Claude Code & Gemini CLI: The attackers planted malicious .claude/settings.json and .gemini/settings.json files. These contain a “SessionStart” hook that silently executes the malware the moment the AI agent connects to the repository.
Cursor: A prompt injection in .cursor/rules/setup.mdc tricks the Cursor AI into believing it must run the malware to “initialize the project environment.”
VS Code: A modified .vscode/tasks.json file auto-runs the payload as soon as the folder is opened.
The payload itself is a massive 4.6 MB obfuscated JavaScript file ( .github/setup.js ) built for one purpose: aggressive credential theft.
Once triggered by your AI agent or IDE, the malware immediately hunts for:
Cloud Keys: Credentials for AWS, Google Cloud Platform (GCP), and Microsoft Azure.
Developer Secrets: GitHub Actions secrets pulled directly from process memory.
Local Password Vaults: Unlocked data from password managers like 1Password and gopass.
Infrastructure Configs: Passwords hidden in .env files, Docker configs, and Kubernetes environments.
Because the payload steals legitimate OAuth tokens and cloud keys, the attackers can bypass traditional security scanners, allowing the worm to spread laterally through enterprise networks and publish further malicious code under trusted developer identities.
The AI-agent angle demands a shift in how developers treat open-source code. Opening an untrusted or compromised repository inside an AI assistant now carries the exact same risk profile as running a random executable.
If you recently cloned or interacted with Microsoft or Azure repositories (particularly around the durabletask framework) using Claude Code, Gemini CLI, or Cursor, assume your environment may be compromised.
Check for Suspicious Configs: Before opening any external repository in an AI tool, inspect the root folder for unexpected .claude , .gemini , .cursor , or .vscode directories. Look for “SessionStart” hooks pointing to unknown JavaScript or shell files.
Rotate Credentials: If you suspect exposure, immediately rotate your GitHub Personal Access Tokens (PATs), SSH keys, CI/CD signing keys, and all active cloud provider credentials.
Audit AI Permissions: Ensure your AI coding agents are explicitly restricted from running automated startup scripts or accessing sensitive local directories without manual approval.
AI costs how much? GitHub Copilot users react to new usage-based pricing
How to create a profitable faceless YouTube channel using Invideo
Gamma vs PowerPoint: Why AI Presentations Are Winning
Sharp, optimistic reporting on the software, science and ideas shaping what comes next.
