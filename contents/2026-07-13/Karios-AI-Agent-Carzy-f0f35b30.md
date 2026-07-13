---
source: "https://github.com/adnqcr7-code/kairosv2"
hn_url: "https://news.ycombinator.com/item?id=48899535"
title: "Karios AI Agent Carzy"
article_title: "GitHub - adnqcr7-code/kairosv2: Local-first AI agent platform with skills, memory, model routing, and extensible tools. · GitHub"
author: "kairos_agent"
captured_at: "2026-07-13T22:45:13Z"
capture_tool: "hn-digest"
hn_id: 48899535
score: 2
comments: 0
posted_at: "2026-07-13T22:04:42Z"
tags:
  - hacker-news
  - translated
---

# Karios AI Agent Carzy

- HN: [48899535](https://news.ycombinator.com/item?id=48899535)
- Source: [github.com](https://github.com/adnqcr7-code/kairosv2)
- Score: 2
- Comments: 0
- Posted: 2026-07-13T22:04:42Z

## Translation

タイトル: カリオス AI エージェント カージー
記事のタイトル: GitHub - adnqcr7-code/kairosv2: スキル、メモリ、モデル ルーティング、拡張可能なツールを備えたローカル ファーストの AI エージェント プラットフォーム。 · GitHub
説明: スキル、メモリ、モデル ルーティング、拡張可能なツールを備えたローカル ファーストの AI エージェント プラットフォーム。 - adnqcr7-code/kairosv2

記事本文:
GitHub - adnqcr7-code/kairosv2: スキル、メモリ、モデル ルーティング、拡張可能なツールを備えたローカル ファーストの AI エージェント プラットフォーム。 · GitHub
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
adnqcr7-コード
/
カイロスv2
公共
通知
通知設定を変更するにはサインインする必要があります
追加のn

ナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット docs docs skill/ 01-coding skill/ 01-coding src/ kairos src/ kairos .gitignore .gitignore CHANGES_SUMMARY.md CHANGES_SUMMARY.md IMPROVEMENTS_README.md IMPROVEMENTS_README.md IMPROVEMENTS_SUMMARY.md IMPROVEMENTS_SUMMARY.md KAIROS_FULL_AGENT_REPORT_2026-06-02.md KAIROS_FULL_AGENT_REPORT_2026-06-02.md KAIROS_IMPROVEMENT_PLAN.md KAIROS_IMPROVEMENT_PLAN.md ライセンス ライセンス README.md README.md STATUS_REPORT.md STATUS_REPORT.md ULTIMATE_VERSION_SUMMARY.md ULTIMATE_VERSION_SUMMARY.md 承認出力.txt 承認出力.txt kairosv2-ultimate.zip kairosv2-ultimate.zip kairosv2-upgraded.zip kairosv2-upgraded.zip math.js math.js math.test.js math.test.js package.json package.json tmp-test.txt tmp-test.txt tmp-tool-run-diff.txt tmp-tool-run-diff.txt tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ゴール モード、保護されたスウォーム、スマート モデル ルーティング、メモリ、ツール、バンドルされた AI スキル ライブラリを備えたローカル ファーストのコーディング エージェント プラットフォーム。
Kairos は、コーディング アシスタント、自動化ワークフロー、リサーチ エージェント、Discord ツール、将来の AI 搭載アプリケーションに柔軟な基盤を提供するために構築された実験的なローカル ファースト AI エージェント システムです。
通常のチャットボットとは異なり、Kairos は、スキル、メモリ、モデル ルーティング、プロバイダー管理、ツールの実行、安全性チェック、エージェントのワークフローなどのモジュール式システムを中心に構築されています。
長期的な目標は、Kairos を端末ベースの MVP から、より強力な CLI ツール、Discord コントロール、オプションの音声フック、ローカル モデル、プラグイン、適応型ユーザー制御メモリを備えた完全なローカル ファースト エージェント エコシステムに成長させることです。
Kairos は現在鋭意開発中です。
このリリースは初期の MVP として扱う必要があります。ローカルコマンドを実行できるんだよ

年齢目標、スキルのロード、プロバイダー設定の使用、メモリの保存、保護されたエージェントのワークフローのデモンストレーションを行います。
一部の機能は実験的なものであり、内部 API はバージョン間で変更される可能性があります。
/goal データの永続性/kairos/goals
作業開始前の保護された群れの計画
安価でバランスの取れた最良の戦略を備えたモデルルーター
現在、ベスト モードはクラウド プロバイダー アダプターが接続されるまでルーティング プレビューです。
プランナー、ビルダー、レビューアー、テスター、およびパッケージャーの役割契約
このリポジトリに含まれるバンドル AI スキル フォルダー
スキルの提案が新しい目標に保存されました
Ollama、OpenAI、Anthropic、Gemini、Kimi、OpenRouter、および将来の Codex ブリッジの脳セットアップ
準備ができているツールと計画されているツールを示すツール レジストリ
スキル、記憶、レッスン、プロジェクトのスニペットのローカル コンテキスト インデックス
ChatGPT スタイルのターミナル チャット (セットアップ、ヘルス チェック、履歴、スラッシュ コマンドを使用)
npm.cmd run kairos を使用したワンコマンドターミナルメニュー
npm.cmd run kairos -- Doctor による医師の健康レポート
選択した KAIROS_DATA_DIR に保存されたローカル メモリ
中/高リスクのシェル アクションに対する承認プロンプトを含むコマンドを確認しました
承認されたシェルアクション用のオプションの Ubuntu Docker サンドボックス
承認されたローカル ルート内での保護されたフォルダーの作成と zip パッケージ化
承認された HTTP(S) ドキュメント/リサーチのフェッチ用に保護された web_fetch
ローカル プレビューおよび承認されたリモート ブラウザ ターゲット用の保護されたブラウザオープン
--multi-agent を使用したオプションのプランナー/ビルダー/レビュー担当者のマルチエージェント計画
ゴールラン後に記録されたローカルの自己改善レッスン
Kairos には、番号付きのカテゴリ フォルダーにバンドルされたスキル ライブラリが含まれるようになりました。
01-コーディング/
23-エージェント-エンジニアリング/
28-ローカル-AI-システム/
...
含まれるスキルには、コーディング、バックエンド、フロントエンド、DevOps、セキュリティ、AI/ML、ドキュメント、製品、コミュニケーション、学習、メタエージェント ワークフロー、イノベーション、管理、サイバーセキュリティが含まれます。

y、倫理、法務、財務、マーケティング、販売、運営、顧客サポートなど。
スキル ドキュメントは以下にも含まれています。
docs/HOW_TO_USE_SKILLS_WITH_ANY_AI.md
docs/SKILL_ROUTING_MAP.md
docs/SKILL_CHAIN_RECIPES.md
バンドルされたスキルフォルダーの使用
Kairos は、最初に KAIROS_SKILLS_DIR を使用してスキルを検索します。これが設定されていない場合は、このリポジトリにバンドルされているスキル フォルダーのレイアウトが優先されます。
npm.cmd kairos を実行 -- スキルリスト
次に、スキルを検索または表示できます。
npm.cmd kairos を実行 -- スキル検索「モデルセレクター」
npm.cmd kairos を実行 -- スキル ショー 01 - コーディング:コード - レビューアー
コマンド
プロジェクト ルートから Kairos を実行します。
npm.cmd カイロスを実行
一般的なコマンド:
npm.cmd kairos -- setup を実行します。
npm.cmd カイロスを実行 -- ドクター
npm.cmd kairos を実行 -- チャット
npm.cmd kairos を実行 -- チャットセットアップ
npm.cmd kairos を実行 -- チャット -- プロバイダー ollama -- はい
npm.cmd カイロスを実行 -- 脳の状態
npm.cmd kairos を実行 -- 脳チェック
npm.cmd カイロスを実行 -- ブレインチェック -- すべて
npm.cmd カイロスを実行 -- ブレインリスト
npm.cmd カイロスを実行 -- 脳のセットアップ
npm.cmd kairos を実行 -- Brain Setup -- Provider ollam -- はい
npm.cmd kairos を実行 -- Brain Setup -- Provider Codex -- Yes
npm.cmd kairos を実行 -- 脳が「このプロジェクトについて説明してください」と尋ねます
npm.cmd kairos を実行 -- サンドボックスのステータス
npm.cmd kairos を実行 -- mkdir アプリ / デモ
npm.cmd kairos を実行 -- zip README.md パッケージ / readme.zip
npm.cmd kairos を実行 -- 「npm.cmd test」を実行します。
npm.cmd kairos を実行 -- 「 uname -a 」を実行 -- サンドボックス ubuntu - docker -- はい
npm.cmd kairos を実行 -- web_fetch https://example.com -- はい
npm.cmd kairos を実行 -- browser_open localhost: 3000
npm.cmd kairos を実行 -- -- ヘルプ
npm.cmd run kairos -- / 目標 " 再利用可能な Discord チケット ボット テンプレートを構築する " -- 低予算 -- 承認ステップ
npm.cmd run kairos -- / 目標「テスト ランナーの改善」 -- マルチエージェント
npm.cmd kairos を実行 -- ステータス
npm.cmd kairos を実行 -- 承認 < 実行

al - id > -- マルチエージェント
npm.cmd kairos を実行 -- スキルリスト
npm.cmd kairos を実行 -- スキル検索「モデルセレクター」
npm.cmd kairos を実行 -- スキル ショー 01 - コーディング:コード - レビューアー
npm.cmd kairos を実行 -- コンテキスト ビルド
npm.cmd kairos を実行 -- コンテキスト検索「セキュリティ監査人」
npm.cmd kairos を実行 -- ツールリスト
npm.cmd カイロスの実行 -- レッスン 5
npm.cmd kairos を実行 -- スキャン
npm.cmd kairos を実行 -- 「kairos」src を検索
npm.cmd kairos を実行 -- README.md を読む
npm.cmd kairos を実行 -- discord を構築する - ボット アプリ / クライアント - ボット
npm.cmd kairos を実行 -- ノードの構築 - cli アプリ / ツール - cli
npm.cmd kairos を実行 -- ログ 10
npm.cmd kairos を実行 -- プロバイダー リスト
npm.cmd kairos を実行 -- プロバイダーのステータス
npm.cmd kairos を実行 -- プロバイダーのセットアップ
npm.cmd kairos を実行 -- プロバイダーのセットアップ -- プロバイダー ollam -- はい
npm.cmd kairos を実行 -- プロバイダーのセットアップ -- プロバイダー openai
npm.cmd
[切り捨てられた]
npm.cmd kairos -- setup を実行します。
これにより、Kairos 警告画面が表示され、ローカル データとメモリの保存場所を尋ねられ、プロバイダーを選択できるようになります。
Ollama キーや API キーが必要ない場合は、「オフライン」を選択します。
AI 脳がないということは、LLM が接続されていないことを意味します。 Kairos は目標、安全警告、スキル、ツール、メモリを管理できますが、チャットと推論には Ollama、OpenAI、Kimi、OpenRouter などの構成されたプロバイダーが必要です。
Kairos は以下を使用して新しい Discord ボット スターターを生成できます。
npm.cmd kairos を実行 -- discord を構築する - ボット アプリ / クライアント - ボット
AIチャット
Kairos には、プロバイダーが設定されると通常の AI アシスタントのように動作する対話型のターミナル チャットが含まれています。
npm.cmd kairos を実行 -- チャット
チャット設定を直接実行します。
npm.cmd kairos を実行 -- チャットセットアップ
npm.cmd kairos を実行 -- チャット -- プロバイダー ollama -- はい
便利なチャット コマンド:
/help チャット コマンドを表示する
/setup AI Brain プロバイダーを選択または修正します
/brain プロバイダーと健康状態を表示します
/history 最近の会話者を表示

ターン
/clear 会話履歴をリセット
/cost 推定セッションコストを表示します
/exit チャットを終了する
Ollama が選択されているが接続できない場合、チャットでは最初のメッセージの前に正確な健康上の問題が表示され、再試行、セットアップの実行、続行、または終了ができるようになりました。
Kairos はローカル コンテキスト インデックスを構築できるため、エージェント ループは以下の関連スライスを取得できます。
プロジェクト ファイルとドキュメント
npm.cmd kairos を実行 -- コンテキスト ビルド
手動で検索します。
npm.cmd kairos を実行 -- コンテキスト検索「セキュリティ監査人」
スキル/メモリ/レッスンのみの再構築を高速化するには:
npm.cmd kairos を実行 -- コンテキスト ビルド -- いいえ - プロジェクト
現在のインデックスは依存関係のない語彙検索であり、まだ完全な埋め込みベクトル データベースではありません。スニペットを保存する前に一般的なシークレット パターンを編集し、Kairos データ ディレクトリにローカル インデックスを書き込みます。
オプションの Ubuntu コマンド サンドボックス
Kairos は、ホスト PowerShell の代わりに、Ubuntu Docker コンテナーまたは Ubuntu WSL ディストリビューションを介して、承認されたシェル コマンドをルーティングできます。
これは、Linux コマンドをテストしたり、エージェント実行中のホストの危険性を軽減したりするのに役立ちます。 Docker モードはより強力なサンドボックスです。 WSL モードは Linux 実行環境ですが、より多くのホスト統合を共有するため、分離性が低いものとして扱う必要があります。
Ubuntu イメージはローカルで入手できます。 Kairos のデフォルトは KAIROS_SANDBOX_PULL=never なので、Docker はエージェントの実行中にイメージをプルしません。
npm.cmd kairos を実行 -- サンドボックスのステータス
Ubuntu サンドボックスで 1 つのコマンドを実行します。
npm.cmd kairos を実行 -- 「 uname -a 」を実行 -- サンドボックス ubuntu - docker -- はい
npm.cmd kairos を実行 -- 「 uname -a 」を実行 -- サンドボックス ubuntu - wsl -- はい
これを .env で設定して、エージェント実行アクションのサンドボックスを有効にします。
KAIROS_COMMAND_SANDBOX=ubuntu-docker
KAIROS_UBUNTU_IMAGE=ubuntu:24.04
KAIROS_WSL_DISTRO=カイロスUbuntu
KAIROS_SANDBOX_NETWORK=なし
KAIROS_SANDBOX_WORKSPACE_MO

DE=rw
KAIROS_SANDBOX_PULL=決してしない
KAIROS_SANDBOX_USER=1000:1000
サンドボックスの動作:
ワークスペースは /workspace にマウントされます
Kairos データは /kairos-data にマウントされます
承認された Ubuntu イメージのみが受け入れられます: ubuntu:20.04 、 ubuntu:22.04 、 ubuntu:24.04 、 ubuntu:jammy 、 ubuntu:noble
ネットワークは、KAIROS_SANDBOX_NETWORK=none でデフォルトで無効になっています
読み取り専用のワークスペース コマンドを実行するには、KAIROS_SANDBOX_WORKSPACE_MODE=ro を設定します。
コンテナの強化には、削除された Linux 機能、新しい特権の禁止、プロセス/メモリ/CPU 制限、読み取り専用コンテナ ルートが含まれます。
Docker イメージのプルはデフォルトでは無効になっています。サンドボックス実行を有効にする前に、docker pull ubuntu:24.04 を使用して自分でイメージをプルします。
WSL モードは KAIROS_WSL_DISTRO を使用し、デフォルトは KairosUbuntu です。
WSL モードは Docker モードのように強化されていません。可能な場合は Docker を使用して分離を強化する
コマンドの承認は依然としてサンドボックスの実行前に行われます
記憶力とレッスンランキングの向上
ターミナルエージェントのワークフローの改善
コストとチェックポイントの可視性の向上
今後の研究特集: 適応パーソナリティ学習
Kairos は最終的にはオプションの適応学習システムをサポートする可能性があります。
ユーザーが有効にすると、Kairos は次のことを徐々に学習できます。
好ましい反応と行動
パーソナルプロデュース

[切り捨てられた]

## Original Extract

Local-first AI agent platform with skills, memory, model routing, and extensible tools. - adnqcr7-code/kairosv2

GitHub - adnqcr7-code/kairosv2: Local-first AI agent platform with skills, memory, model routing, and extensible tools. · GitHub
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
adnqcr7-code
/
kairosv2
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits docs docs skills/ 01-coding skills/ 01-coding src/ kairos src/ kairos .gitignore .gitignore CHANGES_SUMMARY.md CHANGES_SUMMARY.md IMPROVEMENTS_README.md IMPROVEMENTS_README.md IMPROVEMENTS_SUMMARY.md IMPROVEMENTS_SUMMARY.md KAIROS_FULL_AGENT_REPORT_2026-06-02.md KAIROS_FULL_AGENT_REPORT_2026-06-02.md KAIROS_IMPROVEMENT_PLAN.md KAIROS_IMPROVEMENT_PLAN.md LICENSE LICENSE README.md README.md STATUS_REPORT.md STATUS_REPORT.md ULTIMATE_VERSION_SUMMARY.md ULTIMATE_VERSION_SUMMARY.md approval-output.txt approval-output.txt kairosv2-ultimate.zip kairosv2-ultimate.zip kairosv2-upgraded.zip kairosv2-upgraded.zip math.js math.js math.test.js math.test.js package.json package.json tmp-test.txt tmp-test.txt tmp-tool-run-diff.txt tmp-tool-run-diff.txt tsconfig.json tsconfig.json View all files Repository files navigation
Local-first coding agent platform with goal mode, guarded swarms, smart model routing, memory, tools, and a bundled AI skills library.
Kairos is an experimental local-first AI agent system built to provide a flexible foundation for coding assistants, automation workflows, research agents, Discord tools, and future AI-powered applications.
Unlike a normal chatbot, Kairos is built around modular systems: skills, memory, model routing, provider management, tool execution, safety checks, and agent workflows.
The long-term goal is to grow Kairos from a terminal-based MVP into a complete local-first agent ecosystem with stronger CLI tools, Discord control, optional voice hooks, local models, plugins, and adaptive user-controlled memory.
Kairos is currently in active development.
This release should be treated as an early MVP. It can run local commands, manage goals, load skills, use provider configuration, store memory, and demonstrate guarded agent workflows.
Some features are experimental, and internal APIs may change between versions.
/goal persistence in data/kairos/goals
Guarded swarm plan before work starts
Model router with cheap, balanced, and best strategies
best mode is currently a routing preview until cloud provider adapters are connected
Planner, builder, reviewer, tester, and packager role contracts
Bundled AI skills folder included in this repository
Skill suggestions saved into new goals
Brain setup for Ollama, OpenAI, Anthropic, Gemini, Kimi, OpenRouter, and a future Codex bridge
Tool registry showing ready vs planned tools
Local context index for skills, memory, lessons, and project snippets
ChatGPT-style terminal chat with setup, health checks, history, and slash commands
One-command terminal menu with npm.cmd run kairos
Doctor health report with npm.cmd run kairos -- doctor
Local memory stored under your chosen KAIROS_DATA_DIR
Reviewed commands with approval prompts for medium/high-risk shell actions
Optional Ubuntu Docker sandbox for approved shell actions
Guarded folder creation and zip packaging inside approved local roots
Guarded web_fetch for approved HTTP(S) docs/research fetching
Guarded browser_open for local previews and approved remote browser targets
Optional planner/builder/reviewer multi-agent planning with --multi-agent
Local self-improvement lessons recorded after goal runs
Kairos now includes a bundled skills library in numbered category folders:
01-coding/
23-agent-engineering/
28-local-ai-systems/
...
The included skills cover coding, backend, frontend, DevOps, security, AI/ML, documentation, product, communication, learning, meta-agent workflows, innovation, management, cybersecurity, ethics, legal, finance, marketing, sales, operations, customer support, and more.
Skill docs are also included in:
docs/HOW_TO_USE_SKILLS_WITH_ANY_AI.md
docs/SKILL_ROUTING_MAP.md
docs/SKILL_CHAIN_RECIPES.md
Using the bundled skills folder
Kairos looks for skills using KAIROS_SKILLS_DIR first. If that is not set, it prefers this repo's bundled skills folder layout.
npm.cmd run kairos -- skills list
Then you can search or view skills:
npm.cmd run kairos -- skills search " model selector "
npm.cmd run kairos -- skills show 01 - coding:code - reviewer
Commands
Run Kairos from the project root:
npm.cmd run kairos
Common commands:
npm.cmd run kairos -- setup
npm.cmd run kairos -- doctor
npm.cmd run kairos -- chat
npm.cmd run kairos -- chat setup
npm.cmd run kairos -- chat -- provider ollama -- yes
npm.cmd run kairos -- brain status
npm.cmd run kairos -- brain check
npm.cmd run kairos -- brain check -- all
npm.cmd run kairos -- brain list
npm.cmd run kairos -- brain setup
npm.cmd run kairos -- brain setup -- provider ollama -- yes
npm.cmd run kairos -- brain setup -- provider codex -- yes
npm.cmd run kairos -- brain ask " explain this project "
npm.cmd run kairos -- sandbox status
npm.cmd run kairos -- mkdir apps / demo
npm.cmd run kairos -- zip README.md packages / readme.zip
npm.cmd run kairos -- run " npm.cmd test "
npm.cmd run kairos -- run " uname -a " -- sandbox ubuntu - docker -- yes
npm.cmd run kairos -- web_fetch https: // example.com -- yes
npm.cmd run kairos -- browser_open localhost: 3000
npm.cmd run kairos -- -- help
npm.cmd run kairos -- / goal " Build a reusable Discord ticket bot template " -- budget cheap -- approval step
npm.cmd run kairos -- / goal " Improve the test runner " -- multi - agent
npm.cmd run kairos -- status
npm.cmd run kairos -- approve < goal - id > -- multi - agent
npm.cmd run kairos -- skills list
npm.cmd run kairos -- skills search " model selector "
npm.cmd run kairos -- skills show 01 - coding:code - reviewer
npm.cmd run kairos -- context build
npm.cmd run kairos -- context search " security auditor "
npm.cmd run kairos -- tools list
npm.cmd run kairos -- lessons 5
npm.cmd run kairos -- scan
npm.cmd run kairos -- search " kairos " src
npm.cmd run kairos -- read README.md
npm.cmd run kairos -- build discord - bot apps / client - bot
npm.cmd run kairos -- build node - cli apps / tool - cli
npm.cmd run kairos -- logs 10
npm.cmd run kairos -- providers list
npm.cmd run kairos -- providers status
npm.cmd run kairos -- providers setup
npm.cmd run kairos -- providers setup -- provider ollama -- yes
npm.cmd run kairos -- providers setup -- provider openai
npm.cmd
[truncated]
npm.cmd run kairos -- setup
This shows the Kairos warning screen, asks where to store local data and memory, and lets you choose a provider.
Pick Offline if you do not want Ollama or API keys.
No AI brain means there is no LLM connected. Kairos can still manage goals, safety warnings, skills, tools, and memory, but chat and reasoning require a configured provider such as Ollama, OpenAI, Kimi, or OpenRouter.
Kairos can generate a fresh Discord bot starter with:
npm.cmd run kairos -- build discord - bot apps / client - bot
AI Chat
Kairos includes an interactive terminal chat that behaves more like a normal AI assistant once a provider is configured.
npm.cmd run kairos -- chat
Run chat setup directly:
npm.cmd run kairos -- chat setup
npm.cmd run kairos -- chat -- provider ollama -- yes
Useful chat commands:
/help Show chat commands
/setup Choose or fix the AI brain provider
/brain Show provider and health status
/history Show recent conversation turns
/clear Reset conversation history
/cost Show estimated session cost
/exit Leave chat
If Ollama is selected but not reachable, chat now shows the exact health problem before the first message and lets you retry, run setup, continue anyway, or exit.
Kairos can build a local context index so the agent loop can retrieve relevant slices of:
Project files and documentation
npm.cmd run kairos -- context build
Search it manually:
npm.cmd run kairos -- context search " security auditor "
For a faster skills/memory/lessons-only rebuild:
npm.cmd run kairos -- context build -- no - project
The current index is dependency-free lexical retrieval, not a full embedding vector database yet. It redacts common secret patterns before storing snippets and writes the local index under your Kairos data directory.
Optional Ubuntu Command Sandbox
Kairos can route approved shell commands through an Ubuntu Docker container or an Ubuntu WSL distro instead of host PowerShell.
This is useful for testing Linux commands and for reducing host exposure during agent runs. Docker mode is the stronger sandbox. WSL mode is a Linux execution environment, but it shares more host integration and should be treated as less isolated.
The Ubuntu image available locally. Kairos defaults to KAIROS_SANDBOX_PULL=never so Docker does not pull images during agent execution.
npm.cmd run kairos -- sandbox status
Run one command in the Ubuntu sandbox:
npm.cmd run kairos -- run " uname -a " -- sandbox ubuntu - docker -- yes
npm.cmd run kairos -- run " uname -a " -- sandbox ubuntu - wsl -- yes
Enable the sandbox for agent run actions by setting this in .env :
KAIROS_COMMAND_SANDBOX=ubuntu-docker
KAIROS_UBUNTU_IMAGE=ubuntu:24.04
KAIROS_WSL_DISTRO=KairosUbuntu
KAIROS_SANDBOX_NETWORK=none
KAIROS_SANDBOX_WORKSPACE_MODE=rw
KAIROS_SANDBOX_PULL=never
KAIROS_SANDBOX_USER=1000:1000
Sandbox behavior:
Workspace is mounted at /workspace
Kairos data is mounted at /kairos-data
Only approved Ubuntu images are accepted: ubuntu:20.04 , ubuntu:22.04 , ubuntu:24.04 , ubuntu:jammy , ubuntu:noble
Network is disabled by default with KAIROS_SANDBOX_NETWORK=none
Set KAIROS_SANDBOX_WORKSPACE_MODE=ro for read-only workspace command runs
Container hardening includes dropped Linux capabilities, no-new-privileges, process/memory/CPU limits, and a read-only container root
Docker image pulling is disabled by default. Pull the image yourself with docker pull ubuntu:24.04 before enabling sandboxed runs.
WSL mode uses KAIROS_WSL_DISTRO and defaults to KairosUbuntu
WSL mode is not hardened like Docker mode; use Docker for stronger isolation when available
Command approval still happens before sandbox execution
Improved memory and lesson ranking
Better terminal agent workflows
Improved cost and checkpoint visibility
Future Research Feature: Adaptive Personality Learning
Kairos may eventually support an optional adaptive learning system.
When enabled by the user, Kairos could gradually learn:
Preferred responses and behaviors
Personal produ

[truncated]
