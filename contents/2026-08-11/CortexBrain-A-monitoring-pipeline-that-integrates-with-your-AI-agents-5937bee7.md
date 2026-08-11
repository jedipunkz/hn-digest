---
source: "https://github.com/CortexFlow/CortexBrain"
hn_url: "https://news.ycombinator.com/item?id=49254618"
title: "CortexBrain: A monitoring pipeline that integrates with your AI agents"
article_title: "GitHub - CortexFlow/CortexBrain: CortexBrain is an ambitious open-source project designed to build an intelligent, lightweight, and highly efficient monitoring platform for distributed cloud and hybrid (cloud-edge) workflows · GitHub"
author: "LorenzoTettaman"
captured_at: "2026-08-11T08:03:46Z"
capture_tool: "hn-digest"
hn_id: 49254618
score: 1
comments: 1
posted_at: "2026-08-11T07:37:15Z"
tags:
  - hacker-news
  - translated
---

# CortexBrain: A monitoring pipeline that integrates with your AI agents

- HN: [49254618](https://news.ycombinator.com/item?id=49254618)
- Source: [github.com](https://github.com/CortexFlow/CortexBrain)
- Score: 1
- Comments: 1
- Posted: 2026-08-11T07:37:15Z

## Translation

タイトル: CortexBrain: AI エージェントと統合する監視パイプライン
記事のタイトル: GitHub - CortexFlow/CortexBrain: CortexBrain は、分散クラウドおよびハイブリッド (クラウド エッジ) ワークフロー向けのインテリジェントで軽量、高効率の監視プラットフォームを構築するように設計された野心的なオープンソース プロジェクトです · GitHub
説明: CortexBrain は、分散型クラウドおよびハイブリッド (クラウド エッジ) ワークフロー用のインテリジェントで軽量、高効率の監視プラットフォームを構築するために設計された野心的なオープンソース プロジェクトです - CortexFlow/CortexBrain

記事本文:
GitHub - CortexFlow/CortexBrain: CortexBrain は、分散クラウドおよびハイブリッド (クラウド エッジ) ワークフロー向けのインテリジェントで軽量、高効率の監視プラットフォームを構築するように設計された野心的なオープンソース プロジェクトです · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
コーテックスフロー
/
皮質脳
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル

857 コミット 857 コミット .github .github Doc Doc Examples 例 cli cli core core mcp mcp .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md build-all.sh build-all.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
開発者が分散アプリケーションを効果的に監視および分析して、迅速かつ効率的に問題を解決できるようにする
Rust と eBPF で書かれており、低オーバーヘッドと高パフォーマンスを実現
テッタマンティ・ロレンツォ 📧 lorenzotettamanti5@gmail.com
ロレンツォ・ブラダニーニ 📧 lorenzolollobrada@gmail.com
CortexBrain は、分散クラウドおよびハイブリッド (クラウドエッジ) ワークフロー向けのインテリジェントで軽量、高効率の監視プラットフォームを構築するために設計された野心的なオープンソース プロジェクトです。
eBPF の機能を活用することで、CortexBrain は分散クラスター内のネットワーキングと可観測性を適切に管理し、リソースの無駄を制限し、全体的なパフォーマンスを向上させることができます。
CortexBrain のコア アーキテクチャ、インストール、実際のアプリケーションに関する包括的な情報は、公式ドキュメントと CortexFlow ブログで入手できます。
🔎 より深い洞察 : CortexBrain は eBPF を統合し、アプリケーション コードを変更することなく、システムのより深いカーネルレベルの洞察を生成します。
🚁 サイドカー オーバーヘッドなし: サイドカー プロキシに関連する追加の CPU とメモリのオーバーヘッドを排除するサイドカーレス アーキテクチャ
🔒 安全性: Linux BPF Verifier は、すべてのプログラムが安全に実行できることを保証します。JIT コンパイラーは、バイトコードをネイティブ CPU 命令に変換し、最適な実行効率を実現します。 CortexBrain は、イングレス (着信) TCP/UDP 接続などのネットワーク トラフィックを追跡し、T などのさまざまなフックにプログラムを接続することで、カーネル レベルでポリシーを直接適用できます。

C (トラフィック制御) および XDP フック。 BPF マップと専用のデータ構造のおかげで、傍受されたすべてのイベントはユーザー空間に正常に伝播されます。
🧑🏻‍🔬 現在の開発の焦点
当社の現在の開発努力は、次の主要な機能に重点を置いています。
🧪[実験的] GPU オブザーバビリティ: AI/ML アプリケーションを効率的にサポートするための GPU トレースおよびモニタリング機能の導入
🤖 AI 統合の拡張 - 現在の MCP サーバーを拡張してコーディング エージェント プラットフォームとシームレスに統合し、AI 支援によるシステム分析を可能にします
🚁 パイプラインの簡素化 - モニタリング パイプラインを簡素化して、全体的なオーバーヘッドを削減し、障害点を減らします。
📡 [実験的] クラウドエッジマルチクラスター統合: ハイブリッドクラウドとエッジ環境全体で可観測性を拡張
コンポーネント
バージョン
画像
ステータス
皮質フローエージェント
0.1.2/最新
ghcr.io/cortexflow/agent:0.1.2
不安定
皮質フローアイデンティティ
0.1.2/最新
ghcr.io/cortexflow/identity:0.1.2
不安定
皮質フローメトリクス
0.1.0/最新
ghcr.io/cortexflow/metrics:0.1.0
不安定
ドキュメント
アーキテクチャ : 最新バージョンのアーキテクチャの概要
開発者ガイド : 完全な開発者ガイド
一般的な問題: eBPF フレームワークを使用したプログラミング中に発生する一般的な文書化された問題
MCP サーバー : オープンコードによるアーキテクチャ、構築、構成
CLI : セットアップとコマンドをカバーする完全な CLI ドキュメント
⚠️ CortexBrain はまだ開発段階にあるため、いくつかのバグが予想されます。プロジェクトの改善に役立てるため、貢献とフィードバックをいただければ幸いです。
CortexBrain は、コマンド ライン インターフェイスのおかげで、ユーザーに簡単なインストールを提供します。インストールガイドは公式ドキュメントにあります。
カーゴインストール Cortexflow-cli
ローカルクラスターを開始します
CortexBrain コンポーネントをインストールする
cfcl

Cortexflowをインストールします
インストールされているすべてのサービスを一覧表示します
cfcli サービスのリスト
💪🏻 貢献しています
このプロジェクトには何かが欠けていると思いますか?貢献することは、自分のスキルを示し、プロジェクトに足跡を残すための最良の方法です。
DevOps/Kubernetes、ネットワーキング、セキュリティに詳しい場合、または単にリポジトリの保守が好きな場合は、lorenzotettamanti5@gmail.com にメールを書いてください。
コミュニティからの貢献を歓迎します。プロジェクトに貢献するには、次の手順に従ってください。
機能の新しいブランチを作成します ( git checkout -b feature/feature-name )。
変更の詳細な説明を添えてプル リクエストを送信します。
新しい機能に貢献したい場合は、PR を送信する前にディスカッションを開いてください。これは、すべての新機能がプロジェクトの目標と一致していることを確認し、作業の重複やビューの矛盾を避けるためです。
GitHub ディスカッション セクションでディスカッションを開始してください。実装を開始する前に、アイデアを共同でレビュー、改良、承認できます。プロジェクトの一貫性を維持し、より広範なロードマップとの整合性を確保するために、事前に議論されていない新機能のプル リクエストは拒否される場合があります。
このように協力することで、明確さと一貫性を維持し、すべての貢献者が同じ目標に向かって取り組んでいることを確認できます。ご理解とご協力に感謝いたします。
特定の条件に従って、AI によって生成されたコードを含むプル リクエストを受け入れます。まず第一に、新しい機能が十分に文書化されていることを確認してください。次に、提出したコードの背後にあるアイデアを理解し、説明できることを証明します。新しい更新を行うために AI を使用することを推奨するわけではありません。私たちは、あなたが提出したコードの背後に明確な人間の思考プロセスがあったことを確認したいだけです。
コードベースはenです

これらは人間によって丹念に書かれており、AI は Linux カーネル コードベースを完全に理解するために必要な知識のギャップを平坦にするために使用されます。たとえば、「このトレースポイントはどこにあるのか?」などです。または、このカーネル構造の tgid フィールドのオフセットは何ですか? AI 支援開発を使用してドキュメントのバックボーンの構築をスピードアップしましたが、すべてのセクションは人間によって慎重にレビューされました。アーキテクチャの図は AI で強化されました。これは AI エージェントに提出したバージョンです
CortexBrain は、分散型クラウドおよびハイブリッド (クラウド エッジ) ワークフロー向けのインテリジェントで軽量かつ高効率な監視プラットフォームを構築するために設計された野心的なオープンソース プロジェクトです。
Readme Apache-2.0 ライセンスの行動規範
セキュリティ ポリシー アクティビティ カスタム プロパティ スター
8 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

CortexBrain is an ambitious open-source project designed to build an intelligent, lightweight, and highly efficient monitoring platform for distributed cloud and hybrid (cloud-edge) workflows - CortexFlow/CortexBrain

GitHub - CortexFlow/CortexBrain: CortexBrain is an ambitious open-source project designed to build an intelligent, lightweight, and highly efficient monitoring platform for distributed cloud and hybrid (cloud-edge) workflows · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
CortexFlow
/
CortexBrain
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
857 Commits 857 Commits .github .github Doc Doc Examples Examples cli cli core core mcp mcp .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md build-all.sh build-all.sh View all files Repository files navigation
Enabling developers to effectively monitor and analyze distributed applications for rapid and efficient problem solving
Written in Rust and eBPF for low-overhead and high-performance
Tettamanti Lorenzo 📧 lorenzotettamanti5@gmail.com
Lorenzo Bradanini 📧 lorenzolollobrada@gmail.com
CortexBrain is an ambitious open-source project designed to build an intelligent, lightweight, and highly efficient monitoring platform for distributed cloud and hybrid (cloud–edge) workflows.
By leveraging the power of eBPF, CortexBrain can successfully manage networking and observability in a distributed cluster, limiting resource waste and improving overall performance.
Comprehensive information about CortexBrain’s core architecture, installation, and practical applications is available in the Official Documentation and on the CortexFlow blog .
🔎 Deeper Insights : CortexBrain integrates eBPF to produce deeper kernel-level insights of your system without modifying your application code
🚁 No sidecar overhead: Sidecarless architecture that eliminates additional CPU and memory overhead associated with sidecar proxies
🔒 Safety: The linux BPF Verifier ensures that all the programs are safe to run.A JIT compiler converts bytecode into native CPU instructions for optimal execution efficiency. CortexBrain can trace network traffic such as ingress (incoming) TCP/UDP connections and apply policies directly at kernel level by attaching the programs in different hooks such as TC (traffic control) and XDP hooks. All the intercepted events are successfully propagated in the user space thanks to BPF maps and dedicated data structures.
🧑🏻‍🔬 Current Development Focus
Our current development efforts are dedicated to the following key features:
🧪[Experimental] GPU Observability : Introducing GPU tracing and monitoring capabilities to efficiently support AI/ML applications
🤖 Extending AI integrations - Extending the current MCP server to seamlessly integrate with coding agent platforms and enable AI-assisted system analysis
🚁 Simplify the pipeline - Simplify the monitoring pipeline to reduce the overall overhead and reduce points of failures
📡 [Experimental] Cloud-Edge Multi-Cluster Integration: Extending observability accross hybrid cloud and edge environments
Component
Version
Image
Status
cortexflow-agent
0.1.2/latest
ghcr.io/cortexflow/agent:0.1.2
Unstable
cortexflow-identity
0.1.2/latest
ghcr.io/cortexflow/identity:0.1.2
Unstable
cortexflow-metrics
0.1.0/latest
ghcr.io/cortexflow/metrics:0.1.0
Unstable
Documentation
Architecture : Latest version architecture overview
Developer Guide : Full developer guide
Common Issues : Common documented issues encountered while programming with the eBPF framework
MCP server : Architecture, building and configuration with opencode
CLI : Full CLI documentation covering setup and commands
⚠️ CortexBrain is still in its development stages, so you can expect some bugs. Contributions and feedback are highly appreciated to help improve the project!
CortexBrain provides a simple installation for users thanks to his command line interface. You can find the installation guide in the official documentation
cargo install cortexflow-cli
Start your local cluster
Install CortexBrain components
cfcli install cortexflow
List all the installed services
cfcli service list
💪🏻 Contributing
Do you think the project is missing something? Contributing is the best way to show your skills and leave your mark on a project.
If you know DevOps/Kubernetes, networking, security, or you just enjoy maintaining a repository, please write an email to lorenzotettamanti5@gmail.com
We welcome contributions from the community! To contribute to the project, please follow these steps:
Create a new branch for your feature ( git checkout -b feature/feature-name ).
Submit a Pull Request with a detailed explanation of your changes.
If you would like to contribute a new feature, we ask you to open a discussion before submitting a PR. This is to ensure that all new features align with the project's goals and to avoid overlapping work or conflicting views.
Please initiate a discussion in the GitHub Discussions section where we can collectively review, refine, and approve your idea before you begin implementation. Pull Requests for new features that have not been discussed beforehand may be declined to maintain project coherence and ensure alignment with the broader roadmap.
By collaborating in this manner, we can maintain clarity and consistency, ensuring that all contributors are working towards the same objectives. Thank you for your understanding and contributions!
We accept Pull Requests containing AI-generated code, subject to certain conditions. First of all, make sure the new functionalities are well documented; secondly, prove that you can understand and explain the idea behind the code you have submitted. We don't discourage the use of AI to help make new updates; we just want to make sure that there was a clear human thinking process behind the code you submit.
The codebase is entirely written by humans, and the AI is used to flatten the knowledge gap required to fully understand the Linux kernel codebase- things like: where is this tracepoint located? Or what is the offset for the tgid field in this kernel structure?. We used AI-assisted development to speed up the building of the backbones of the documentation, but all the sections were carefully reviewed by humans. The illustration of the architecture was enhanced with AI, this is the version we submitted to the AI agent
CortexBrain is an ambitious open-source project designed to build an intelligent, lightweight, and highly efficient monitoring platform for distributed cloud and hybrid (cloud-edge) workflows
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
8 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
