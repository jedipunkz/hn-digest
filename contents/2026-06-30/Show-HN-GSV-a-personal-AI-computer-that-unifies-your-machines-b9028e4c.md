---
source: "https://github.com/deathbyknowledge/gsv"
hn_url: "https://news.ycombinator.com/item?id=48731946"
title: "Show HN: GSV – a personal AI computer that unifies your machines"
article_title: "GitHub - deathbyknowledge/gsv: Personal AI computer / distributed OS across all your devices. · GitHub"
author: "deathbyknowledg"
captured_at: "2026-06-30T13:34:44Z"
capture_tool: "hn-digest"
hn_id: 48731946
score: 4
comments: 0
posted_at: "2026-06-30T12:44:12Z"
tags:
  - hacker-news
  - translated
---

# Show HN: GSV – a personal AI computer that unifies your machines

- HN: [48731946](https://news.ycombinator.com/item?id=48731946)
- Source: [github.com](https://github.com/deathbyknowledge/gsv)
- Score: 4
- Comments: 0
- Posted: 2026-06-30T12:44:12Z

## Translation

タイトル: Show HN: GSV – マシンを統合するパーソナル AI コンピューター
記事のタイトル: GitHub - deathbyknowledge/gsv: パーソナル AI コンピューター / すべてのデバイスにわたる分散 OS。 · GitHub
説明: パーソナル AI コンピューター / すべてのデバイスに分散された OS。 - deathbyknowledge/gsv
HN テキスト: 4 月に私は Cloudflare のエージェント チームを辞めました。仕事は大好きでしたが、ライブラリではなくアプリを構築する意欲が常にあったからです。私が望んでいたアイデアを誰も構築していないように感じました。私の最大の動機は、エージェントを完全に「クラウド」上で実行し、同時にすべてのマシン上でエージェントを実行したいということでした。そこで私は完全にオープンソースの GSV を構築することになりました。 GSV は、リモート ファイル システムとシェル環境を備え、Unix に近いものであるかのように懸命に努力することで、クラウド コンピュータとして機能します。 AI 推論やエージェント ループなど、すべての基本的な操作を処理するカーネルがあります。同時に、デバイス (およびブラウザ!) が GSV に接続して、同じエージェントへのファイル システムとシェル アクセスも提供できます。これにより、エッジで実行されているエージェントがローカル デバイスにアクセスできるようになります。ベータ版をリリースしたばかりなので、フィードバックをお待ちしています。現在、Workers Payed アカウント (月あたり最大 5 ドル) とモデル費用が必要です。

記事本文:
GitHub - deathbyknowledge/gsv: パーソナル AI コンピューター / すべてのデバイスに分散された OS。 · GitHub
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
死の知識
/
gsv
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メイン

ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1,673 コミット 1,673 コミット .github .github アダプタ アダプタ アセンブラ アセンブラ cli cli docs docs エンジニアリング拡張機能 拡張機能 ゲートウェイ ゲートウェイ パッケージ/ gsv パッケージ/ gsv ripgit ripgit スクリプト スクリプト スキル スキル Web Web .gitignore .gitignore AGENTS.md AGENTS.md CHANNELS.md CHANNELS.md ライセンスライセンス README.md README.md TODOS.md TODOS.md バージョン バージョン favicon.svg favicon.svg install.ps1 install.ps1 install.sh install.sh package-lock.json package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
🟢🚀 パブリックベータ版が登場しました!問題やPRは大歓迎です。
パーソナル AI エージェントのほとんどは、ラップトップ、VPS、コンテナーなど、選択して維持される 1 つのホスト上で実行されます。 GSV はすべてのデバイスを 1 台のコンピューターに変えます。これはラップトップ、サーバー、電話を一度に拡張し、エージェントが適切なもので動作し、自分の Cloudflare アカウントのエッジで実行できるようにします。キーもデータも、プロビジョニングやベビーシッターを行うホストは必要ありません。インフラストラクチャは月あたり最大 $5 に加え、独自のモデル費用がかかります。
1 つのエージェントからすべてのマシンで作業を実行します。ラップトップが閉じている間に、ホーム サーバーでジョブを開始します。
デバイスがスリープ状態になっている間もエージェントは動作し続けます。エージェントはハードウェア上ではなくエッジ上に存在します。
Web UI、CLI、WhatsApp / Discord / Telegram など、どこからでもアクセスできます。
独自のメモリと権限を持つ永続エージェントを生成し、独自のサブエージェントを開始できます。
独自のパッケージをホストし、組み込みの Git リモートを通じて GSV インスタンス間でアプリを共有します。
エージェントにブラウザを渡します。 Web 拡張機能を使用すると、実際のブラウザ (タブやログイン セッション) を操作できるため、パブリック Web だけでなく、すでに使用しているサイトでも機能します。
内部では分散 OS です: エージェントは耐久性のあるプロセスです

ID、履歴、権限、システムコール サーフェスに加えて、アプリを構築するための SDK が含まれています。 Iain M. Banks の Culture シリーズに登場するセンティエント シップにちなんで名付けられた GSV (General Systems Vehicle) は、エッジを越えて動作するパーソナル AI の基盤です。
前提条件: Workers Payed プラン (月額 $5) の Cloudflare アカウント、および実行したいモデルプロバイダーの API キー。モデルの使用量は、そのプロバイダーによって別途請求されます。
Web から (最も簡単、ターミナルは不要)。 deploy.gsv.space に移動し、Cloudflare アカウントに接続すると、GSV 自体がそのアカウントにデプロイされます。
# CLIをインストールする
カール -sSL https://install.gsv.space |バッシュ
# すべてのコンポーネントを自分の Cloudflare アカウントにデプロイします
gsv infradeploy --api-token < CLOUDFLARE-API-TOKEN >
いずれの場合も、出力される URL を開いて、Web UI でのオンボーディングを完了します。
Web UI からすぐに、または CLI からチャットできます。
gsv チャット「こんにちは、何かお手伝いできることはありますか?」
メッセンジャー (Discord / Telegram / WhatsApp) に接続し、デバイスを追加し、次に何をするかを確認するには、 docs.gsv.space/get-started にある完全なガイドに従ってください。
接続されたデバイスにはどこからでもエージェントがアクセスできます。送信のみなので、開いているポート、受信接続、VPN はありません。 Web UI または CLI の [GSV] > [デバイス] から追加します。
gsv auth token create --device macbook --label Macbook # トークンをメモします
gsv config --local set node.token < トークン >
gsv device install --id macbook --workspace ~ / # バックグラウンド サービス
gsv デバイスのステータス
これで、GSV はシェルを使用して、そのマシン上でファイルの読み取り/書き込みができるようになります。 [GSV] > [統合] でアダプターをセットアップします。
設計上 Linux に似ているため、エージェントはよく知られたパターン (POSIX ではなくメンタル モデル) に基づいて推論できます。
カーネル — ゲートウェイは Cloudflare 上で実行され、認証されたシステムコール ( proc.* 、 pkg.* 、 sys.* ) を公開します。
プロセス - エージェントは PID を持つ永続的なプロセスです (

gsv proc list|spawn|send|kill )。
デバイス — 接続されたマシンは、ワークスペースにスコープされた実行ノードとして機能します。
メッセンジャー — Discord/Telegram/Whatsapp ワーカーは、外部チャットのデバイス ドライバーのように機能します。
./scripts/setup-deps.sh # ワークスペース、アダプター、ripgit 全体に JS deps をインストールします
cd web && npm run build # Web アプリをビルドする
cd .. && npm run dev # ローカルマルチワーカー開発スタック
Rust と Node.js + npm が必要です。
GSV は積極的に進化しており、あなたにもネットワークの一員になってもらいたいと考えています。あらゆる規模の寄付を歓迎します。
プル リクエストを送信したい場合でも、突飛なアイデアを共有したい場合でも、ただ挨拶したい場合でも、お気軽にお問い合わせください。 :)
コミュニティに参加する: Discord サーバー でぶらぶらしたり、ショッピングについて話したり、アイデアを共有したりしてください。
バグが見つかった場合、または機能リクエストがありますか?問題を開きます。
更新情報をフォローする: Twitter/X @gsvspace で直接連絡してください。
パーソナル AI コンピューター / すべてのデバイスに分散された OS。
読み込み中にエラーが発生しました。このページをリロードしてください。
14
フォーク
レポートリポジトリ
リリース
19
GSV v0.3.0
最新の
2026 年 6 月 30 日
+ 18 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Personal AI computer / distributed OS across all your devices. - deathbyknowledge/gsv

Back in April I left the Agents team at Cloudflare because even though I loved the work, I always had the drive to build the apps rather than the libraries. I felt like nobody was building the ideas that I wanted to exist. My biggest drive was wanting my agents to run fully on “the cloud” yet having them on all my machines at the same time. So I ended up building GSV, fully open source. GSV acts as a cloud computer by having a remote file system and a shell environment and trying very hard to pretend to be something close to Unix. There’s a kernel that handles all the primitive operations, including AI inference and the agent loop. At the same time, your devices (and browser!) can connect your GSV to provide file system and shell access to those same agents too. This allows your agents running on the edge to have access to your local devices. I’ve just released the beta and would love to hear feedback. As of today, it requires a Workers Paid account (~5$/mo) plus your model costs.

GitHub - deathbyknowledge/gsv: Personal AI computer / distributed OS across all your devices. · GitHub
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
deathbyknowledge
/
gsv
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1,673 Commits 1,673 Commits .github .github adapters adapters assembler assembler cli cli docs docs engineering engineering extension extension gateway gateway packages/ gsv packages/ gsv ripgit ripgit scripts scripts skills skills web web .gitignore .gitignore AGENTS.md AGENTS.md CHANNELS.md CHANNELS.md LICENSE LICENSE README.md README.md TODOS.md TODOS.md VERSION VERSION favicon.svg favicon.svg install.ps1 install.ps1 install.sh install.sh package-lock.json package-lock.json package.json package.json View all files Repository files navigation
🟢🚀 Public beta is here! Issues and PRs very welcome.
Most personal AI agents run on one host you pick and keep alive — a laptop, a VPS, a container. GSV turns all your devices into a single computer. It spans your laptop, server, and phone at once, lets an agent act on whichever one fits, and runs on the edge in your own Cloudflare account; your keys, your data, no host to provision or babysit. From ~$5/mo infra plus your own model costs.
Run things across all your machines from one agent — kick off a job on your home server while your laptop's shut.
Keep agents working while your devices sleep — they live on the edge, not on your hardware.
Reach it from anywhere — web UI, CLI, or WhatsApp / Discord / Telegram.
Spawn durable agents with their own memory and permissions, that can start sub-agents of their own.
Host your own packages and share apps between GSV instances through a built-in git remote.
Hand your agent the browser. The web extension lets it drive your real browser — your tabs and logged-in sessions — so it works the sites you already use, not just the public web.
Under the hood it's a distributed OS: agents are durable processes with identities, history, permissions, and a syscall surface, plus an SDK for building apps. Named after the sentient ships from Iain M. Banks' Culture series, GSV (General Systems Vehicle) is a foundation for personal AI that lives across the edge.
Prerequisites: a Cloudflare account on the Workers Paid plan ($5/mo), and an API key for whatever model provider you want to run. Your model usage is billed separately by that provider.
From the web (easiest, no terminal). Go to deploy.gsv.space , connect your Cloudflare account, and GSV deploys itself into it.
# Install the CLI
curl -sSL https://install.gsv.space | bash
# Deploy all components into your own Cloudflare account
gsv infra deploy --api-token < CLOUDFLARE-API-TOKEN >
Either way, open the URL it prints to finish onboarding in the web UI.
Chat from the web UI right away, or from the CLI:
gsv chat " Hello, what can you help me with? "
To connect a messenger (Discord / Telegram / WhatsApp), add more devices, and see what to do next, follow the full guide at docs.gsv.space/get-started .
Connected devices are reachable by your agents from anywhere — outbound-only, so no open ports, no inbound connections, no VPN. Add one via GSV > Devices in the Web UI, or the CLI:
gsv auth token create --device macbook --label Macbook # note the token
gsv config --local set node.token < token >
gsv device install --id macbook --workspace ~ / # background service
gsv device status
Now GSV can use the shell and read/write files on that machine. Set up adapters under GSV > Integrations .
Linux-like by design, so agents can reason with familiar patterns (mental model, not POSIX).
Kernel — the Gateway runs on Cloudflare, exposing authenticated syscalls ( proc.* , pkg.* , sys.* ).
Processes — agents are durable processes with PIDs ( gsv proc list|spawn|send|kill ).
Devices — connected machines act as execution nodes, scoped to a workspace.
Messengers — Discord/Telegram/Whatsapp workers act like device drivers for external chat.
./scripts/setup-deps.sh # install JS deps across workspace, adapters, ripgit
cd web && npm run build # build web app
cd .. && npm run dev # local multi-worker dev stack
Requires Rust and Node.js + npm .
GSV is actively evolving, and we want you to be part of the network! We welcome contributions of all sizes.
Whether you want to submit a pull request, share a wild idea, or just say hi, please don't hesitate to reach out. :)
Join the Community: Come hang out, talk shop, and share ideas on our Discord Server .
Found a bug or have a feature request? Open an issue .
Follow Updates: Reach out directly on Twitter/X @gsvspace
Personal AI computer / distributed OS across all your devices.
There was an error while loading. Please reload this page .
14
forks
Report repository
Releases
19
GSV v0.3.0
Latest
Jun 30, 2026
+ 18 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
