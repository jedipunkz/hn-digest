---
source: "https://github.com/open-jarvis/OpenJarvis"
hn_url: "https://news.ycombinator.com/item?id=48347357"
title: "OpenJarvis: Personal AI, on Personal Devices"
article_title: "GitHub - open-jarvis/OpenJarvis: Personal AI, On Personal Devices · GitHub"
author: "simonpure"
captured_at: "2026-06-01T00:27:40Z"
capture_tool: "hn-digest"
hn_id: 48347357
score: 4
comments: 0
posted_at: "2026-05-31T17:02:56Z"
tags:
  - hacker-news
  - translated
---

# OpenJarvis: Personal AI, on Personal Devices

- HN: [48347357](https://news.ycombinator.com/item?id=48347357)
- Source: [github.com](https://github.com/open-jarvis/OpenJarvis)
- Score: 4
- Comments: 0
- Posted: 2026-05-31T17:02:56Z

## Translation

タイトル: OpenJarvis: パーソナル AI、パーソナル デバイス上
記事のタイトル: GitHub - open-jarvis/OpenJarvis: パーソナル AI、パーソナル デバイス上 · GitHub
説明: パーソナル AI、パーソナル デバイス上。 GitHub でアカウントを作成して、open-jarvis/OpenJarvis の開発に貢献してください。

記事本文:
GitHub - open-jarvis/OpenJarvis: パーソナル AI、パーソナル デバイス上 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
オープンジャービス
/
オープンジャービス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
770 コミット 770 コミット .github .github 資産 アセット configs/ openjarvis configs/ openjarvis デプロイ デプロイ デスクトップ/ src-tauri デスクトップ/ src-tauri docs docs 例 例 フロントエンド フロントエンド Rust Rust スクリプト スクリプト src/ openjarvis src/ openjarvis テスト テストツール/ Pearl-reference-oracle tools/ Pearl-reference-oracle .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md REVIEW.md REVIEW.md mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
パーソナル AI をパーソナル デバイスに。
パーソナル AI エージェントの人気は爆発的に高まっていますが、ほぼすべてのエージェントが依然としてクラウド API を通じてインテリジェンスをルーティングしています。あなたの「個人的な」AI は、他人のサーバーに依存し続けます。同時に、当社のワットあたりのインテリジェンスの調査では、ローカル言語モデルがすでにシングルターン チャットと推論クエリの 88.7% を処理しており、インテリジェンスの効率が 2023 年から 2025 年にかけて 5.3 倍向上していることが示されました。モデルとハードウェアの準備はますます進んでいます。欠けているのは、ローカルファーストのパーソナル AI を実用化するためのソフトウェア スタックです。
OpenJarvis がそのスタックです。これはローカル ファーストのパーソナル AI のフレームワークであり、3 つの核となるアイデアを中心に構築されています。オンデバイス エージェントを構築するための共有プリミティブ。エネルギー、FLOP、レイテンシ、ドルコストを精度とともに第一級の制約として扱う評価。ローカル トレース データを使用してモデルを改善する学習ループ。目標はシンプルです。デフォルトでローカルで実行され、本当に必要な場合にのみクラウドを呼び出すパーソナル AI エージェントを構築できるようにすることです。 OpenJarvis は両方のリソースを目指しています

PyTorch の精神に基づいた、ローカル AI のプラットフォームと生産基盤です。
プラットフォームを選択して、1 つのコマンドを実行します。各インストーラーは uv 、Python venv、Ollama、およびスターター モデルを処理します。ブロードバンドで約 3 分かかります。
それからジャービスが始めます。 Rust 拡張機能とより大きなモデルはバックグラウンドでダウンロードを続けます。ジャービス医師はステータスを示します。
プラットフォーム固有の注意事項 (WSL2 セットアップ、ネイティブ Windows のスケジュールされたタスク サービス、デスクトップの前提条件、手動/共同作成者によるインストール): インストールに関するドキュメント を参照してください。
jarvis # チャットを開始します (デフォルト: chat-simple)
jarvis init --preset < name > # スターター構成に切り替えます
jarvis ... の前に uv run を付けるか、source .venv/bin/activate を最初に付けます。
jarvis init --preset breakfast-digest-mac
jarvis connect gdrive # one OAuth は Gmail / カレンダー / タスクをカバーします
jarvis Digest --fresh # 最初のブリーフィングを生成して再生します
プリセットごとの詳細な説明: 朝のダイジェスト、詳細な調査、コード アシスタント、スケジュールされたモニター、シンプルなチャット、または完全なクイックスタート ガイド。
スキルは、エージェントにツールをより効果的に使用し、推論を改善する方法を教えます。すべてのスキルはツールです。エージェントはカタログからスキルを発見し、オンデマンドで呼び出します。
# 公開ソースからスキルをインストールする
ジャービス スキル インストール hermes:arxiv
ジャービス スキル同期 ヘルメス -- カテゴリ調査
# 任意のエージェントでスキルを使用する
jarvis ask 「コード説明スキルを使用して、この Python コードを説明します: for i in range(5): print(i*2)」
# トレース履歴からスキルを最適化する
ジャービスはスキルを最適化します --policy dspy
# 影響をベンチマークする
ジャービスベンチスキル --最大サンプル 5 --シード 42
Hermes Agent (約 150 のスキル)、OpenClaw (約 13,700 のコミュニティ スキル)、または任意の GitHub リポジトリからインポートします。スキルは、agentskills.io オープン スタンダードに従います。
詳細については、「スキル ユーザー ガイド」および「スキル チュートリアル」を参照してください。
OpenJarvis には 8 つの組み込みモジュールが同梱されています。

3 つの実行モード (オンデマンド、スケジュール、継続的) にわたるエージェント:
セットアップ手順の詳細については、ユーザー ガイドとチュートリアルを参照してください。
Docker のデプロイメント、クラウド エンジン、開発セットアップ、チュートリアルを含む完全なドキュメントは、 open-jarvis.github.io/OpenJarvis にあります。
GitHub: github.com/open-jarvis/OpenJarvis
ディスコード: discord.gg/YZZRxCAhmm
ドキュメント: open-jarvis.github.io/OpenJarvis
寄付を歓迎します!インセンティブ、寄付の種類、PR プロセスについては、「寄付ガイド」を参照してください。
git clone https://github.com/open-jarvis/OpenJarvis.git
cd オープンジャービス
UV 同期 --extra dev
UV実行プリコミットインストール
uv pytest テストを実行/ -v
ロードマップを参照して、支援が必要な領域を見つけてください。問題に「take」とコメントすると、自動的に割り当てられます。
OpenJarvis は、AI システムのインテリジェンス効率を研究する研究イニシアチブである Intelligence Per Watt の一部です。このプロジェクトは、スタンフォード SAIL の Hazy Research と Scaling Intelligence Lab で開発されています。
ロード研究所 •
スタンフォード・マーロウ •
Google クラウド プラットフォーム •
ラムダ研究所 •
オラマ •
IBMの調査 •
スタンフォードHAI
@misc { saadfalcon2026openjarvispersonalaipersonal ,
title = { OpenJarvis: パーソナル AI、パーソナル デバイス上 } ,
著者 = { ジョン・サード・ファルコン、アヴァニカ・ナラヤン、ロビー・マニハニ、タンヴィル・バータル、ヘルンブ・シャンディリヤ、ハッキ・オルフン・アケンギン、ガブリエル・ボー、アンドリュー・パーク、マシュー・ハート、カイア・コステロ、チュアン・リー、クリストファー・レ、アザリア・ミルホセイニ } 、
年 = { 2026 } 、
eprint = { 2605.17172 } 、
archivePrefix = { arXiv } 、
プライマリクラス = { cs.LG } 、
URL = { https://arxiv.org/abs/2605.17172 } 、
}
ライセンス
パーソナル AI、パーソナル デバイス上
open-jarvis.github.io/OpenJarvis/
リソース
Apache-2.0ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
S

タール
1.2k
フォーク
レポートリポジトリ
リリース
5
デスクトップ デスクトップ v1.0.2
最新の
2026 年 5 月 25 日
+ 4 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Personal AI, On Personal Devices. Contribute to open-jarvis/OpenJarvis development by creating an account on GitHub.

GitHub - open-jarvis/OpenJarvis: Personal AI, On Personal Devices · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
open-jarvis
/
OpenJarvis
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
770 Commits 770 Commits .github .github assets assets configs/ openjarvis configs/ openjarvis deploy deploy desktop/ src-tauri desktop/ src-tauri docs docs examples examples frontend frontend rust rust scripts scripts src/ openjarvis src/ openjarvis tests tests tools/ pearl-reference-oracle tools/ pearl-reference-oracle .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md REVIEW.md REVIEW.md mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Personal AI, On Personal Devices.
Personal AI agents are exploding in popularity, but nearly all of them still route intelligence through cloud APIs. Your "personal" AI continues to depend on someone else's server. At the same time, our Intelligence Per Watt research showed that local language models already handle 88.7% of single-turn chat and reasoning queries, with intelligence efficiency improving 5.3× from 2023 to 2025. The models and hardware are increasingly ready. What has been missing is the software stack to make local-first personal AI practical.
OpenJarvis is that stack. It is a framework for local-first personal AI, built around three core ideas: shared primitives for building on-device agents; evaluations that treat energy, FLOPs, latency, and dollar cost as first-class constraints alongside accuracy; and a learning loop that improves models using local trace data. The goal is simple: make it possible to build personal AI agents that run locally by default, calling the cloud only when truly necessary. OpenJarvis aims to be both a research platform and a production foundation for local AI, in the spirit of PyTorch.
Pick your platform and run one command. Each installer handles uv , the Python venv, Ollama, and a starter model — about 3 minutes on broadband.
Then jarvis to start. The Rust extension and larger models continue downloading in the background; jarvis doctor shows status.
Platform-specific notes (WSL2 setup, native-Windows scheduled-task service, desktop prerequisites, manual / contributor install): see the installation docs .
jarvis # start chatting (default: chat-simple)
jarvis init --preset < name > # switch to a starter config
Prefix jarvis ... with uv run , or source .venv/bin/activate first.
jarvis init --preset morning-digest-mac
jarvis connect gdrive # one OAuth covers Gmail / Calendar / Tasks
jarvis digest --fresh # generate and play your first briefing
Per-preset deep dives: morning digest · deep research · code assistant · scheduled monitor · chat simple · or the full quickstart guide .
Skills teach agents how to better use tools and improve their reasoning. Every skill is a tool — agents discover them from a catalog and invoke them on demand.
# Install skills from public sources
jarvis skill install hermes:arxiv
jarvis skill sync hermes --category research
# Use skills with any agent
jarvis ask " Use the code-explainer skill to explain this Python code: for i in range(5): print(i*2) "
# Optimize skills from your trace history
jarvis optimize skills --policy dspy
# Benchmark the impact
jarvis bench skills --max-samples 5 --seeds 42
Import from Hermes Agent (~150 skills), OpenClaw (~13,700 community skills), or any GitHub repo. Skills follow the agentskills.io open standard.
See the Skills User Guide and Skills Tutorial for details.
OpenJarvis ships with eight built-in agents across three execution modes (on-demand, scheduled, continuous):
See the User Guide and Tutorials for detailed setup instructions.
Full documentation — including Docker deployment, cloud engines, development setup, and tutorials — at open-jarvis.github.io/OpenJarvis .
GitHub: github.com/open-jarvis/OpenJarvis
Discord: discord.gg/YZZRxCAhmm
Docs: open-jarvis.github.io/OpenJarvis
We welcome contributions! See the Contributing Guide for incentives, contribution types, and the PR process.
git clone https://github.com/open-jarvis/OpenJarvis.git
cd OpenJarvis
uv sync --extra dev
uv run pre-commit install
uv run pytest tests/ -v
Browse the Roadmap for areas where help is needed. Comment "take" on any issue to get auto-assigned.
OpenJarvis is part of Intelligence Per Watt , a research initiative studying the intelligence efficiency of AI systems. The project is developed at Hazy Research and the Scaling Intelligence Lab at Stanford SAIL .
Laude Institute •
Stanford Marlowe •
Google Cloud Platform •
Lambda Labs •
Ollama •
IBM Research •
Stanford HAI
@misc { saadfalcon2026openjarvispersonalaipersonal ,
title = { OpenJarvis: Personal AI, On Personal Devices } ,
author = { Jon Saad-Falcon and Avanika Narayan and Robby Manihani and Tanvir Bhathal and Herumb Shandilya and Hakki Orhun Akengin and Gabriel Bo and Andrew Park and Matthew Hart and Caia Costello and Chuan Li and Christopher Ré and Azalia Mirhoseini } ,
year = { 2026 } ,
eprint = { 2605.17172 } ,
archivePrefix = { arXiv } ,
primaryClass = { cs.LG } ,
url = { https://arxiv.org/abs/2605.17172 } ,
}
License
Personal AI, On Personal Devices
open-jarvis.github.io/OpenJarvis/
Resources
Apache-2.0 license
Code of conduct
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
1.2k
forks
Report repository
Releases
5
Desktop desktop-v1.0.2
Latest
May 25, 2026
+ 4 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
