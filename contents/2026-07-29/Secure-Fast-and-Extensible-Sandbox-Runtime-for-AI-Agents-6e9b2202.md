---
source: "https://github.com/TencentCloud/CubeSandbox"
hn_url: "https://news.ycombinator.com/item?id=49100880"
title: "Secure, Fast, and Extensible Sandbox Runtime for AI Agents"
article_title: "GitHub - TencentCloud/CubeSandbox: Instant, Concurrent, Secure & Lightweight Sandbox for AI Agents. · GitHub"
author: "hek2sch"
captured_at: "2026-07-29T18:58:41Z"
capture_tool: "hn-digest"
hn_id: 49100880
score: 3
comments: 0
posted_at: "2026-07-29T18:03:10Z"
tags:
  - hacker-news
  - translated
---

# Secure, Fast, and Extensible Sandbox Runtime for AI Agents

- HN: [49100880](https://news.ycombinator.com/item?id=49100880)
- Source: [github.com](https://github.com/TencentCloud/CubeSandbox)
- Score: 3
- Comments: 0
- Posted: 2026-07-29T18:03:10Z

## Translation

タイトル: AI エージェント向けの安全、高速、拡張可能なサンドボックス ランタイム
記事のタイトル: GitHub - TencentCloud/CubeSandbox: AI エージェント向けの即時、同時、安全かつ軽量のサンドボックス。 · GitHub
説明: AI エージェント向けの即時、同時、安全かつ軽量のサンドボックス。 - テンセントクラウド/キューブサンドボックス

記事本文:
GitHub - TencentCloud/CubeSandbox: AI エージェント向けの即時、同時、安全かつ軽量のサンドボックス。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
テンセントクラウド
/
キューブサンドボックス
パブ

リック
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
658 コミット 658 コミット .claude .claude .github .github CubeAPI CubeAPI CubeDB CubeDB CubeEgress CubeEgress CubeMaster CubeMaster CubeNet CubeNet CubeOps CubeOps CubeProxy CubeProxy CubeShim CubeShim Cubelet Cubelet エージェント エージェント configs configs cube-lifecycle-manager cube-lifecycle-manager cubecow cubecow cubelog cubelogdeploy デプロイ dev-env dev-env docker docker docs docs 例 例 ハイパーバイザー ハイパーバイザー ネットワーク エージェント ネットワーク エージェント スクリプト スクリプト sdk sdk テスト テスト Web Web .dockerignore .dockerignore .gitignore .gitignore .gitmodules .gitmodules AGENTS.md AGENTS.md COTRIBUTING.md CONTRIBUTING.md CONTRIBUTING_zh.md CONTRIBUTING_zh.md LICENSE LICENSE Makefile Makefile README.md README.md README_zh.md README_zh.md openapi.yml openapi.yml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェント向けの即時、同時、安全かつ軽量のサンドボックス サービス
中文档 ·
クイックスタート ·
ドキュメント ·
変更履歴 ·
X(ツイッター)・
トップレベル ユーザー プログラム ·
ユースケースを送信する
Cube Sandbox は、RustVMM および KVM 上に構築された、すぐに使用できる高性能の安全なサンドボックス サービスです。単一ノードの導入と、マルチノード クラスターへの簡単なスケーリングの両方をサポートします。 E2B SDK と互換性があり、ハードウェアから分離され、完全にサービス可能なサンドボックスを 60 ミリ秒未満、メモリ オーバーヘッドが 5 MB 未満で作成できます。
v0.6: K8s デプロイ、ボリューム フレームワーク、テンプレート エイリアス
K8s デプロイ — Cube コントロール プレーン コンポーネントとコンピューティング ノードを Kubernetes にデプロイします
ボリューム フレームワーク — ユーザーがカスタム バックエンド ストレージをプラグインできるようにする E2B 互換のボリューム フレームワーク
テンプレートのエイリアス — テンプレートの作成時にエイリアスを設定します。

そのエイリアスを指定してサンドボックスを作成します。
変更履歴 → ·
K8s デプロイ → ·
ボリュームプラグイン →
v0.5: AutoPause、Terraform デプロイヤ、ARM64 およびネットワーク ポリシーの強化
AutoPause/AutoResume — アイドル状態のサンドボックスは自動的に一時停止し、次のリクエストで起動します。 Terraform ワンクリック クラスター デプロイ ARM64 ネイティブ フルスタック サポート ネットワーク ポリシー強化 - サンドボックスごとのトラフィック トークン、ポリシー ルーティング出力。
変更履歴 → ·
Terraform のデプロイ →
v0.4: より安全な出口、より簡単な操作
Credential Vault — エージェントは通常どおり外部 API を呼び出します。キーがサンドボックスに入ることはありません。ダッシュボード — バージョン マトリックスとテンプレートのヘルス チェック。アップグレード後にテンプレートを再構築する必要があるかどうかが一目でわかります。
変更履歴 → ·
セキュリティプロキシガイド → ·
WebUIガイド →
100ミリ秒単位のスナップショット、クローン、ロールバック
CubeSandbox 0.3.0 では、CubeCoW Copy-on-Write スナップショット エンジンが導入され、イベント レベルのスナップショット、インスタント クローン作成、保存された状態へのロールバックが可能になります。
変更履歴 →
🎉 最初のオープンソースリリース
Cube Sandbox がオープンソースになりました。ミリ秒ブート、ハードウェア レベルの分離、AI エージェント用の E2B 互換サンドボックス。
変更履歴 →
製品のハイライト
⚡ 超高速起動
リソース プーリングとスナップショット クローン作成により、コールド スタートのオーバーヘッドがすべてスキップされます。平均 60 ミリ秒未満のコールド スタート — 瞬くよりも速くサンドボックスを作成します。
クイックスタート→
🔒 ハードウェアの分離
すべてのサンドボックスは、独自の MicroVM で専用の OS カーネルを実行します。
建築 →
🔌 E2B SDK互換
E2B SDKインターフェースに対応。 1 つの環境変数を変更するだけで、E2B クラウドからシームレスに切り替えることができます。クライアント コードの変更は必要ありません。
例 →
📦 高密度導入
サンドボックスあたり 5MB 未満のオーバーヘッドにより、カーネル共有とコピーオンライト (CoW) を介してサーバーあたり数千のインスタンスが可能になります。サンドボックスの自動一時停止と再開をサポートし、さらに改善します

導入密度とコストの最適化。
クイックスタート→
🛡️ ネットワークセキュリティ
eBPF ベースのサンドボックス間分離とカーネル レベルでの出力フィルタリング。組み込みの L7 セキュリティ プロキシにより、自動資格情報挿入によるドメイン/パス/メソッドごとのポリシーが有効になります。シークレットはサンドボックス コードには決して表示されません。
セキュリティプロキシガイド →
📸 柔軟な状態管理
高頻度のスナップショットと 100 ミリ秒単位のロールバック。実行中のサンドボックスにチェックポイントを作成し、いつでも保存された状態にロールバックしたり、特定の状態からフォークして並行して探索したりできます。
v0.3変更履歴→
💾 ボリュームフレームワーク
ユーザーがカスタム バックエンド ストレージ ソリューションをプラグインできるようにする E2B 互換のボリューム フレームワーク。ボリュームには独立したライフサイクルがあり、サンドボックス間で共有できます。
ボリュームプラグイン →
🚀 本番展開
Terraform を使用して、ワンクリックで本番クラスターを Tencent Cloud にデプロイします。標準の Kubernetes クラスター (プレビュー) へのデプロイメントもサポートします。
Terraform デプロイ → ·
K8sのデプロイ →
💪 ARM アーキテクチャのサポート
コンパイル、ビルド、デプロイメントのワークフロー全体でネイティブ ARM64 を完全にサポートします。
ベアメタルのデプロイ →
デモ
1.キューブサンドボックス.-.mp4
2.cubesandbox.demo.mp4
Cube-Sandbox.RL.demo.mp4
5.cube.V0.3.0.-.-.mp4
インストールとデモ
性能試験
RL (SWE-ベンチ)
スナップショット・クローン・ロールバック
ベンチマーク
AI エージェント コード実行のコンテキストにおいて、CubeSandbox はセキュリティとパフォーマンスの完璧なバランスを実現します。
ベアメタルでのコールド スタートのベンチマーク。単一同時実行時は 60 ミリ秒。同時作成数 50 未満、平均 67 ミリ秒、P95 90 ミリ秒、P99 137 ミリ秒 — 一貫して 150 ミリ秒未満。
サンドボックス仕様 ≤ 32 GB で測定されたメモリ オーバーヘッド。大規模な構成では、わずかに増加する可能性があります。
起動遅延とリソース オーバーヘッドに関する詳細なメトリクスについては、「Core Operations Performance Ben」を参照してください。

chmark レポート (ベアメタル) および PVM クラウド サーバー ベンチマーク レポート 。
⚡ ミリ秒レベルの起動 — 上記のファスト スタート フローをご覧ください。
Cube Sandbox には、KVM をサポートする x86_64 Linux 環境が必要です。
このガイドでは、サーバーのプロビジョニング、Cube Sandbox のインストール、サンドボックス テンプレートの作成、最初のエージェント コードの実行という 4 つのステップですべてを説明します。ソースのビルドは必要なく、数分で起動して実行できます。
インストール後の最初の作業: Web コンソールを開きます
🖥️ 視覚的な管理 — 概要からサンドボックスの作成、ログのストリーミングまで、すべてブラウザーで実行できます。
ワンクリックで展開した後、ブラウザで開きます。
http://<制御ノードIP>:12088
推奨される 3 つのステップ:
概要を確認する — [概要] を開き、ノードが準備完了であり、容量が正常であることを確認します。
テンプレートを準備する — テンプレート ストアから公式プリセットをインストールします。 「テンプレート」の下に「準備完了」テンプレートがすでにある場合はスキップしてください。
サンドボックスを作成します — [サンドボックス] → [+ 新しいサンドボックス] を選択し、準備ができたテンプレートを選択して、数秒以内に詳細ページでライブ ログを表示します
完全な WebUI コンソール ガイドを参照してください。
ドキュメント ホーム — 完全なガイド ナビゲーション
☁️ PVM 導入 — ベアメタルやネストされた仮想化を使用せずに、通常のクラウド VM に導入します。
テンプレートの概念 — 画像からテンプレートへの概念とワークフロー
サンプル プロジェクト — 実践的なサンプル (コード実行、ブラウザ自動化、OpenClaw 統合、RL トレーニングなど)
🖥️ WebUI コンソール — インストール直後の視覚的な管理 ( :12088 )
🔐 Security Proxy & Credential Vault — CubeEgress ドメインのフィルタリング、インジェクション、監査
🤖 デジタル アシスタント AgentHub — OpenClaw アシスタントの作成と管理 (プレビュー)
💻 開発環境 (QEMU VM) — KVM アクセスがありませんか?使い捨ての OpenCloudOS 9 VM 内で Cube Sandbox を試してみる
👉 詳細については、こちらをお読みください。

アーキテクチャ設計ドキュメントと CubeVS ネットワーク モデル。
バグレポート、機能の提案、ドキュメントの改善、コードの提出など、あらゆる種類の貢献を歓迎します。
🐞 バグを見つけましたか? それとも質問がありますか? GitHub Issues で問題を送信してください。
💡 何かアイデアはありますか? GitHub ディスカッションの会話に参加してください。
🛠️ コーディングしてみませんか?プル リクエストを送信する方法については、CONTRIBUTING.md を参照してください。
📝 ドキュメントに貢献したいですか?バイリンガル PR をコミュニティ ドキュメント チャネル (トラブルシューティング 、ユースケース 、統合 ) に送信してください。さらに、Cube 100 プログラムが開始されました。Cube を使用して本番環境で AI エージェントを実行する最初の 100 チームを探しています。席数は100席限定。詳細＆お申込みは→
💬 チャットしたいですか？ Discord に参加してください。
近日公開予定 — 詳細については、完全なロードマップをご覧ください。
CubeSandbox は、Apache License 2.0 に基づいてリリースされています。
CubeSandbox の誕生は、オープンソースの巨人たちの肩の上にあります。 Cloud Hypervisor 、Kata Containers 、virtiofsd、containerd-shim-rs、ttrpc-rust などに感謝します。 CubeSandbox の実行モデルに適合するように一部のコンポーネントに特別な変更を加えており、元のファイル内の著作権表示は保持されています。
Cube Sandbox は CNCF Landscape にリストされています。
AI エージェント向けの即時、同時、安全かつ軽量のサンドボックス。
貢献活動 カスタム プロパティ スター
978 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Instant, Concurrent, Secure & Lightweight Sandbox for AI Agents. - TencentCloud/CubeSandbox

GitHub - TencentCloud/CubeSandbox: Instant, Concurrent, Secure & Lightweight Sandbox for AI Agents. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Uh oh!
There was an error while loading. Please reload this page .
TencentCloud
/
CubeSandbox
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
658 Commits 658 Commits .claude .claude .github .github CubeAPI CubeAPI CubeDB CubeDB CubeEgress CubeEgress CubeMaster CubeMaster CubeNet CubeNet CubeOps CubeOps CubeProxy CubeProxy CubeShim CubeShim Cubelet Cubelet agent agent configs configs cube-lifecycle-manager cube-lifecycle-manager cubecow cubecow cubelog cubelog deploy deploy dev-env dev-env docker docker docs docs examples examples hypervisor hypervisor network-agent network-agent scripts scripts sdk sdk tests tests web web .dockerignore .dockerignore .gitignore .gitignore .gitmodules .gitmodules AGENTS.md AGENTS.md CONTRIBUTING.md CONTRIBUTING.md CONTRIBUTING_zh.md CONTRIBUTING_zh.md LICENSE LICENSE Makefile Makefile README.md README.md README_zh.md README_zh.md openapi.yml openapi.yml View all files Repository files navigation
Instant, Concurrent, Secure & Lightweight Sandbox Service for AI Agents
中文文档 ·
Quick Start ·
Documentation ·
Changelog ·
X(Twitter) ·
Top Contributor Program ·
Submit Use Case
Cube Sandbox is a high-performance, out-of-the-box secure sandbox service built on RustVMM and KVM. It supports both single-node deployment and easy scaling to multi-node clusters. It is compatible with the E2B SDK and can create a hardware-isolated, fully serviceable sandbox in under 60ms with less than 5MB of memory overhead.
v0.6: K8s deploy, Volume framework, template aliases
K8s deploy — Deploy Cube control-plane components and compute nodes on Kubernetes
Volume framework — E2B-compatible Volume framework that lets users plug in custom backend storage
Template aliases — Set an alias when creating a template, and create sandboxes by specifying that alias.
Changelog → ·
K8s deploy → ·
Volume plugin →
v0.5: AutoPause, Terraform deployer, ARM64 & network policy hardening
AutoPause/AutoResume — idle sandboxes auto-suspend and wake on the next request. Terraform one-click cluster deploy ARM64 native full-stack support network policy hardening — per-sandbox traffic tokens, policy-routing egress.
Changelog → ·
Terraform deploy →
v0.4: Safer egress, easier ops
Credential vault — Agents call external APIs as usual; keys never enter the sandbox. Dashboard — version matrix and template health checks; see at a glance whether templates need rebuilding after upgrades.
Changelog → ·
Security proxy guide → ·
WebUI guide →
Snapshot, Clone & Rollback at hundred-millisecond granularity
CubeSandbox 0.3.0 introduces the CubeCoW Copy-on-Write snapshot engine, enabling event-level snapshots, instant cloning, and rollback to any saved state.
Changelog →
🎉 Initial open-source release
Cube Sandbox is now open source! Millisecond boot, hardware-level isolation, E2B-compatible sandbox for AI Agents.
Changelog →
Product Highlights
⚡ Ultra-fast Startup
Resource pooling and snapshot cloning skip all cold-start overhead. Average <60ms cold start — sandbox creation faster than a blink.
Quick start →
🔒 Hardware Isolation
Every sandbox runs a dedicated OS kernel in its own MicroVM.
Architecture →
🔌 E2B SDK Compatible
Compatible with E2B SDK interface. Switch from E2B Cloud seamlessly by changing one environment variable — zero client code changes.
Examples →
📦 High-density Deployment
<5MB overhead per sandbox enables thousands of instances per server via kernel sharing and Copy-on-Write (CoW). Supports automatic sandbox pause and resume, further improving deployment density and cost optimization.
Quick start →
🛡️ Network Security
eBPF-based inter-sandbox isolation and egress filtering at kernel level; built-in L7 security proxy enables per-domain/path/method policies with automatic credential injection — secrets never visible to sandbox code.
Security proxy guide →
📸 Flexible State Management
High-frequency snapshot and rollback at hundred-millisecond granularity. Create checkpoints on running sandboxes, roll back to any saved state at any time, or fork from a specific state to explore in parallel.
v0.3 changelog →
💾 Volume Framework
E2B-compatible Volume framework that lets users plug in custom backend storage solutions. Volumes have an independent lifecycle and can be shared across sandboxes.
Volume plugin →
🚀 Production Deployment
Deploy production clusters on Tencent Cloud with one click using Terraform. Also supports deployment on standard Kubernetes clusters (preview).
Terraform deploy → ·
K8s deploy →
💪 ARM Architecture Support
Full native ARM64 support across compilation, build, and deployment workflows.
Bare-metal deploy →
Demos
1.cubesandbox.-.mp4
2.cubesandbox.demo.mp4
Cube-Sandbox.RL.demo.mp4
5.cube.V0.3.0.-.-.mp4
Installation & Demo
Performance Test
RL (SWE-Bench)
Snapshot · Clone · Rollback
Benchmarks
In the context of AI Agent code execution, CubeSandbox achieves the perfect balance of security and performance:
Cold start benchmarked on bare-metal. 60ms at single concurrency; under 50 concurrent creations, avg 67ms, P95 90ms, P99 137ms — consistently sub-150ms.
Memory overhead measured with sandbox specs ≤ 32GB. Larger configurations may see a marginal increase.
For detailed metrics on startup latency and resource overhead, see the Core Operations Performance Benchmark Report (bare metal) and the PVM Cloud Server Benchmark Report .
⚡ Millisecond-level startup — watch the fast-start flow above.
Cube Sandbox requires an x86_64 Linux environment with KVM support.
The guide walks you through everything in four steps — provisioning a server, installing Cube Sandbox, creating a sandbox template, and running your first agent code. No source build needed, up and running in minutes.
First thing after install: open the Web console
🖥️ Visual management — from overview to creating a sandbox and streaming logs, all in your browser.
After one-click deployment, open in your browser:
http://<control-node IP>:12088
Recommended three steps:
Check overview — Open Overview , confirm nodes are Ready and capacity looks healthy
Prepare a template — Install an official preset from Template Store ; skip if you already have a READY template under Templates
Create a sandbox — Sandboxes → + New sandbox , pick a READY template, and view live logs on the detail page within seconds
See the full WebUI console guide .
Documentation Home — complete guide navigation
☁️ PVM Deployment — deploy on ordinary cloud VMs without bare metal or nested virtualization
Template Concepts — image-to-template concepts and workflows
Example Projects — hands-on examples (code execution, browser automation, OpenClaw integration, RL training, and more)
🖥️ WebUI Console — visual management right after install ( :12088 )
🔐 Security Proxy & Credential Vault — CubeEgress domain filtering, injection, and auditing
🤖 Digital Assistant AgentHub — create and manage OpenClaw assistants (Preview)
💻 Development Environment (QEMU VM) — no KVM access? Try Cube Sandbox inside a disposable OpenCloudOS 9 VM
👉 For more details, please read the Architecture Design Document and CubeVS Network Model .
We welcome contributions of all kinds—whether it's a bug report, feature suggestion, documentation improvement, or code submission!
🐞 Found a Bug or have questions? Submit an issue on GitHub Issues .
💡 Have an Idea? Join the conversation in GitHub Discussions .
🛠️ Want to Code? Check out our CONTRIBUTING.md to learn how to submit a Pull Request.
📝 Want to contribute docs? Submit bilingual PRs to our community doc channels: Troubleshooting , Use Cases , and Integrations . Additionally, the Cube 100 Program is now open — we're looking for the first 100 teams running AI agents in production with Cube. Limited to 100 seats. Learn more & apply →
💬 Want to Chat? Join our Discord .
Coming soon — see the full roadmap for details.
CubeSandbox is released under the Apache License 2.0 .
The birth of CubeSandbox stands on the shoulders of open-source giants. Special thanks to Cloud Hypervisor , Kata Containers , virtiofsd, containerd-shim-rs, ttrpc-rust, and others. We have made tailored modifications to some components to fit the CubeSandbox execution model, and the original in-file copyright notices are preserved.
Cube Sandbox is listed in the CNCF Landscape .
Instant, Concurrent, Secure & Lightweight Sandbox for AI Agents.
Contributing Activity Custom properties Stars
978 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
