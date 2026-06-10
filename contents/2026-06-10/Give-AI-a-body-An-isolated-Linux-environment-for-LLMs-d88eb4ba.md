---
source: "https://github.com/dotojr123/open-infro-agentc"
hn_url: "https://news.ycombinator.com/item?id=48476249"
title: "Give AI a body – An isolated Linux environment for LLMs"
article_title: "GitHub - dotojr123/open-infro-agentc: Open Infro Agentc - Open-source AI-powered desktop automation agent · GitHub"
author: "iagencia"
captured_at: "2026-06-10T16:23:13Z"
capture_tool: "hn-digest"
hn_id: 48476249
score: 1
comments: 0
posted_at: "2026-06-10T13:48:50Z"
tags:
  - hacker-news
  - translated
---

# Give AI a body – An isolated Linux environment for LLMs

- HN: [48476249](https://news.ycombinator.com/item?id=48476249)
- Source: [github.com](https://github.com/dotojr123/open-infro-agentc)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T13:48:50Z

## Translation

タイトル: AI に本体を与える – LLM のための隔離された Linux 環境
記事タイトル: GitHub - dotojr123/open-infro-agentc: Open Infro Agentc - オープンソースの AI を活用したデスクトップ オートメーション エージェント · GitHub
説明: Open Infro Agentc - オープンソースの AI を活用したデスクトップ オートメーション エージェント - dotojr123/open-infro-agentc

記事本文:
GitHub - dotojr123/open-infro-agentc: Open Infro Agentc - オープンソースの AI を活用したデスクトップ オートメーション エージェント · GitHub
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
dotojr123
/
オープンインフロエージェント
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション

n 個のオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
27 コミット 27 コミット .github/ workflows .github/ workflows 資産 アセット iagenciad iagenciad 共有 共有 .editorconfig .editorconfig .env.example .env.example .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md ライセンス ライセンス通知 通知 README.md README.md SECURITY.md SECURITY.md Demon.gif Demon.gif docker-compose.yml docker-compose.yml package-lock.json package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AIに体を与えましょう。心を動かしてください。
自律エージェントのためのデジタル有機体。
☝️ MCP 経由で完全な Ubuntu デスクトップを制御する本物の AI エージェント (Qwen) は、ブラウザーを通じてリアルタイムで観察されます。
🇺🇸 英語 | 🇧🇷 ポルトガル語 (ブラジル)
🧠 肉体を持たない知性の問題
2026 年には、インテリジェントな AI エージェントの構築が簡単になります。現実世界のインフラストラクチャ内に安全にデプロイすることは困難です。
GPT-4o、Claude 3.5、Gemini などのモデルは非常にスマートですが、本質的には「瓶の中の頭脳」です。彼らは論理的に考えることはできますが、行動するための管理可能な物理的な環境がありません。現在のほとんどのフレームワークは、エージェントをヘッドレス ブラウザ サンドボックスに制限するか、ホスト システムへの危険で監査不能なアクセスを許可することで、この問題を解決しようとしています。
Open Infra Agent は、世界初のオープンソースの自律オペレーティング環境 (AOE) です。私たちは知性を構築するのではありません。私たちは、AI が安全に住み、働き、行動するデジタル有機体を提供します。
🧬 デジタル有機体の構造
私たちは、生物の生物学的必需品を高性能のオープンソース インフラストラクチャ スタックにマッピングしました。
私たちがよく受ける質問は、「これは Odysseus や Open WebUI のようなワークスペースとどう違うのですか?」というものです。
彼らはノーです

競合他社。それらは補完的です。 Odysseus のようなプロジェクトは、コックピット (チャット インターフェイスおよびマルチユーザー ダッシュボード) として機能します。 Open Infra Agent がエンジンです。
$$\text{ユーザー} \longrightarrow \text{オデュッセウス (コックピット)} \longrightarrow \text{エージェント フレームワーク} \longrightarrow \text{オープン インフラ エージェント (The Organismo / 本体)}$$
実際の実行が舞台裏で行われる重いインフラストラクチャ層を提供します。
🛡️ 現在の制限事項とロードマップ (v0.1)
私たちは絶対的な技術的誠実さを信じています。 Open Infra Agent は多層防御向けに設計されていますが、現在は v0.1 です。
✅ 現在アクティブ: コンテナーの名前空間の分離、ビジュアルなサンドボックス化、execFile インジェクションセーフな入力ルーティング。
🚧 v0.2 で登場: 厳格な Seccomp プロファイル (システム コールのホワイトリスト)、cgroups v2 制限 (エージェントによる CPU/RAM リソースの枯渇を防ぐため)、および追加のみの改ざん明示監査ログ。
標準のランタイム スタックでテスト済み:
開始レイテンシー: ~3.5 秒 — Docker Compose から完全に応答性が高く、ビジュアルな MCP アクセス可能な環境まで。
実行ラウンドトリップ: ~12ms — MCP ツール呼び出しから OS レベルの入力ドライバーまで。
アイドル RAM フットプリント: ~240MB — X11 + XFCE4 + noVNC + NestJS スタック全体。
git クローン https://github.com/dotojr123/open-infro-agentc.git
cd オープンインフロエージェント
docker 構成 --build -d
2. 生物をライブで観察する:
ブラウザを開いて次の場所に移動します。
👉 http://localhost:9990/vnc
3. エージェント (脳) を接続します。
MCP 互換エージェント (Claude Desktop、OpenClaw、Gemini CLI) を神経系に向けます。
👉 http://localhost:9990/mcp
📡 すぐに利用可能な API および MCP ツール
接続すると、エージェントは即座に以下にアクセスできるようになります。
🖱️computer_move_mouse 、computer_click_mouse 、computer_drag_mouse 、computer_scroll
⌨️computer_type_text 、computer_type_keys
📸コンピューティング

er_screenshot (LLM ビジョン用に最適化)
📁 コンピューター読み取りファイル、コンピューター書き込みファイル
🖥️computer_application (VS Code、ターミナル、Firefox をネイティブに起動)
Open Infra Agent は、自律エージェントのためのデジタル有機体です。
現在、ChatGPT や Claude のようなモデルは、単なる「瓶の中に閉じ込められた脳」に過ぎません。彼らには知性はありますが、現実世界で安全に行動できる身体はありません。私たちは、完全な自律オペレーティング環境 (AOE) を提供することでこの問題を解決します。
私たちは、身体 (分離された Ubuntu コンテナー)、神経系 (ネイティブ MCP)、手 (マウス/キーボード自動化 API)、および目 (BrowserOS) を提供します。これらすべては、ブラウザーを介したリアルタイムの人間による監視 (Human-in-the-Loop) によって行われ、DevOps チームとセキュリティ チームがいつでもエージェントを監査して制御できるようになります。
Open Infra Agent は、自律的な AI 実行を監視可能、管理可能、安全にするように構築されています。
Apache-2.0 ライセンスに基づいて配布されます。詳細については、「ライセンス」を参照してください。
このプロジェクトは、Bytebot の強化されたプレミアム フォークです — Copyright Bytebot AI、Apache-2.0。オープンソース コミュニティへの多大な貢献に対して、オリジナルの著者に感謝します。
Open Infro Agentc - オープンソースの AI を活用したデスクトップ オートメーション エージェント
Apache-2.0ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Open Infro Agentc - Open-source AI-powered desktop automation agent - dotojr123/open-infro-agentc

GitHub - dotojr123/open-infro-agentc: Open Infro Agentc - Open-source AI-powered desktop automation agent · GitHub
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
dotojr123
/
open-infro-agentc
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
27 Commits 27 Commits .github/ workflows .github/ workflows assets assets iagenciad iagenciad shared shared .editorconfig .editorconfig .env.example .env.example .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md demo.gif demo.gif docker-compose.yml docker-compose.yml package-lock.json package-lock.json package.json package.json View all files Repository files navigation
Give AI a Body. Run the Mind.
The Digital Organism for Autonomous Agents.
☝️ A real AI agent (Qwen) controlling a full Ubuntu desktop via MCP, observed in real time through the browser.
🇺🇸 English | 🇧🇷 Português (Brasil)
🧠 The Disembodied Intelligence Problem
In 2026, building an intelligent AI agent is easy. Deploying one safely inside real-world infrastructure is not.
Models like GPT-4o, Claude 3.5, and Gemini are incredibly smart, but they are essentially "brains in a jar". They can reason, but they lack a governable, physical-like environment to act. Most current frameworks attempt to solve this by either restricting the agent to a headless browser sandbox, or giving it dangerous, unauditable access to the host system.
Open Infra Agent is the world's first open-source Autonomous Operating Environment (AOE) . We don't build the intelligence; we provide the digital organism where your AI lives, works, and acts safely.
🧬 The Anatomy of the Digital Organism
We mapped the biological necessities of a living organism into a highly performant, open-source infrastructure stack:
A common question we get is: "How does this compare to workspaces like Odysseus or Open WebUI?"
They are not competitors; they are complementary. Projects like Odysseus act as the Cockpit (the chat interface and multi-user dashboard). Open Infra Agent is the Engine.
$$\text{User} \longrightarrow \text{Odysseus (Cockpit)} \longrightarrow \text{Agent Framework} \longrightarrow \text{Open Infra Agent (The Organismo / Body)}$$
We provide the heavy infrastructure layer where the actual execution happens behind the scenes.
🛡️ Current Limitations & Roadmap (v0.1)
We believe in absolute technical honesty. Open Infra Agent is designed for defense-in-depth, but we are currently at v0.1 :
✅ Currently Active: Container namespace isolation, visual sandboxing, and execFile injection-safe input routing.
🚧 Coming in v0.2: Strict Seccomp profiles (system call whitelisting), cgroups v2 limits (to prevent CPU/RAM resource exhaustion by agents), and append-only tamper-evident audit logs.
Tested on our standard runtime stack:
Start Latency: ~3.5 seconds — from docker compose up to a fully responsive, visual MCP-accessible environment.
Execution Roundtrip: ~12ms — from MCP tool call to OS-level input driver.
Idle RAM Footprint: ~240MB — entire X11 + XFCE4 + noVNC + NestJS stack.
git clone https://github.com/dotojr123/open-infro-agentc.git
cd open-infro-agentc
docker compose up --build -d
2. Watch the Organism Live:
Open your browser and navigate to:
👉 http://localhost:9990/vnc
3. Connect Your Agent (The Brain):
Point any MCP-compatible agent (Claude Desktop, OpenClaw, Gemini CLI) to the nervous system:
👉 http://localhost:9990/mcp
📡 API & MCP Tools Available Out-of-the-Box
Once connected, your agent instantly gains access to:
🖱️ computer_move_mouse , computer_click_mouse , computer_drag_mouse , computer_scroll
⌨️ computer_type_text , computer_type_keys
📸 computer_screenshot (Optimized for LLM vision)
📁 computer_read_file , computer_write_file
🖥️ computer_application (Launch VS Code, Terminal, Firefox natively)
Open Infra Agent é o organismo digital para agentes autônomos.
Hoje, modelos como ChatGPT e Claude são apenas "cérebros presos em uma jarra". Eles têm a inteligência, mas não têm um corpo para agir com segurança no mundo real. Nós resolvemos isso fornecendo um Ambiente Operacional Autônomo (AOE) completo.
Nós entregamos o corpo (contêineres Ubuntu isolados), o sistema nervoso (MCP nativo), as mãos (APIs de automação de mouse/teclado) e os olhos (BrowserOS). Tudo isso com Supervisão Humana em Tempo Real (Human-in-the-Loop) via navegador, permitindo que equipes de DevOps e Segurança auditem e assumam o controle do agente a qualquer momento.
Open Infra Agent is built to make autonomous AI execution observable, governable, and safe.
Distributed under the Apache-2.0 License . See LICENSE for details.
This project is a premium, hardened fork of Bytebot — Copyright Bytebot AI, Apache-2.0. We thank the original authors for their outstanding contribution to the open-source community.
Open Infro Agentc - Open-source AI-powered desktop automation agent
Apache-2.0 license
Code of conduct
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
