---
source: "https://github.com/pawaca/dsh-edge"
hn_url: "https://news.ycombinator.com/item?id=49399811"
title: "Show HN: Running a full AI coding agent inside Cloudflare Durable Object"
article_title: "GitHub - pawaca/dsh-edge: Your DeepSeek Harness, anywhere — deploy a persistent personal coding agent to Cloudflare Workers in one command. · GitHub"
image: "https://repository-images.githubusercontent.com/1340840762/40aff7cd-eed6-4aad-af23-76ae29fd56ab"
author: "pawaca"
captured_at: "2026-08-22T14:11:58Z"
capture_tool: "hn-digest"
hn_id: 49399811
score: 2
comments: 0
posted_at: "2026-08-22T13:58:50Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Running a full AI coding agent inside Cloudflare Durable Object

- HN: [49399811](https://news.ycombinator.com/item?id=49399811)
- Source: [github.com](https://github.com/pawaca/dsh-edge)
- Score: 2
- Comments: 0
- Posted: 2026-08-22T13:58:50Z

## Translation

タイトル: HN を表示: Cloudflare Durable Object 内で完全な AI コーディング エージェントを実行する
記事タイトル: GitHub - pawaca/dsh-edge: どこでも DeepSeek ハーネス - 1 つのコマンドで永続的な個人コーディング エージェントを Cloudflare ワーカーにデプロイします。 · GitHub
説明: DeepSeek ハーネスをどこにでも — 1 つのコマンドで永続的なパーソナル コーディング エージェントを Cloudflare Workers にデプロイします。 - パワカ/dsh-edge

記事本文:
GitHub - pawaca/dsh-edge: DeepSeek ハーネスをどこにでも — 1 つのコマンドで永続的な個人コーディング エージェントを Cloudflare ワーカーにデプロイします。 · GitHub
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
パワカ
/
dshエッジ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
77 コミット 77 コミット フォルダーとファイル
.agents .agents .claude .claude .github .github apps/ dsh-edge apps/ dsh-edge docs docs パッケージ/ client/ ui-edge パッケージ/ clie

nt/ ui-edge スクリプト スクリプト .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .node-version .node-version .oxlintrc.json .oxlintrc.json .pnpmfile.cjs .pnpmfile.cjs AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CTRIBUTING.i18n.yaml CTRIBUTING.i18n.yaml COTRIBUTING.md CTRIBUTING.md CTRIBUTING.zh.md CTRIBUTING.zh.md ライセンス ライセンス README.i18n.yaml README.i18n.yaml README.md README.md README.zh.md README.zh.md SECURITY.i18n.yaml SECURITY.i18n.yaml SECURITY.md SECURITY.md SECURITY.zh.md SECURITY.zh.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md lefthook.yml lefthook.yml package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.client.json tsconfig.base.client.json tsconfig.base.json tsconfig.base.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
dsh-edge は、公開されている DeepSeek Harness Web エクスペリエンスを Cloudflare Workers 上で実行するため、ブラウザがあればどこでもパーソナル コーディング エージェントを利用できます。サーバーを保守したり、GitHub リポジトリに接続したり、パイプラインを構築して構成したりする必要はありません。
アップストリーム UI、エージェント ループ、モデルの選択、画像エクスペリエンス、および Web 検索が維持されます。 dsh-edge は、Cloudflare ランタイム、耐久性のあるワークスペース、所有者ログイン、およびガイド付きインストーラーのみを提供します。
独立したコミュニティ プロジェクト: dsh-edge は pawaca によって保守されています。 DeepSeek と提携または承認されていません。 DeepSeek Harness は上流プロジェクトです。
Node.js 22.14 以降と独自の DeepSeek API キーが必要です。
npx dsh-edge インストール
インストーラーはすべての選択をガイドし、ワーカーをデプロイします。既存のCloudflareログインなしで試すことも、自分のCloudflareアカウントに永続的にインストールすることもできます。ソースのチェックアウトは必要ありません。
永続的な会話とワークスペース

任意のブラウザをROMします。
DeepSeek V4 Flash、V4 Pro、およびアップストリーム セレクターを介した実験的な V4 Flash Vision Exp モデル。
PNG/JPEG 画像入力と DeepSeek のネイティブ Web 検索ツール。
ネイティブ DSH bash ツールを備えた永続的な /workspace。
独自の Cloudflare デプロイメント、認証情報、およびデータ。
リポジトリやCloudflare Builds統合を必要としないインプレースアップグレード。
パス
必要なもの
こんな方に最適
今すぐ試してください
既存の Cloudflare ログインはありません。 60分以内に請求して保管してください
摩擦を最小限に抑えた完全なエクスペリエンスを探求する
そのままにしておいてください
R2 が有効になっている既存または新規の Cloudflare アカウント
長期にわたる個人用デプロイメント
デフォルトの無料 — Direct Shell ランタイムは Cloudflare Workers Free で動作します。オプションの Isolated — Dynamic Worker ランタイムは、別個の Worker でコマンドを実行し、Workers Paid を必要とします。どちらのモードでも、同じ製品 UI、会話、ワークスペース、画像、ツールを使用します。 Direct Shell は、Linux コンテナーではなく、サンドボックス化されたシェル ランタイムです。
npx dsh-edge のアップグレード
インストーラーはワーカーを見つけて、その耐久性のあるデータを維持しながら、適切な場所にアップグレードします。バージョン固有の詳細については、リリース ノートを参照してください。
dsh-edge は 1 人の所有者向けに設計されています。登録、複数のユーザー、役割、またはテナントのルーティングは提供されません。
DeepSeek キーは Cloudflare Worker にシークレットとして保存され、デプロイメントの永続的なデータは Cloudflare アカウントに残ります。
Vision Exp は実験的なものであり、アカウントに依存します。一部のアップストリーム機能はまだ Edge に適応されていません。現在のステータスについては、互換性マトリックスを参照してください。
dsh-edge は、モノリポジトリをコピーするのではなく、公開されている DeepSeek Harness パッケージをそのままラップします。アップストリームは引き続き Web UI、プラグイン構成、エージェント ループ、およびセッション プロトコルを担当します。このリポジトリは、Cloudflare ランタイムとストレージ アダプタを実装します。

端で実行してください。
アップストリームのアーキテクチャとプラグイン API については、DeepSeek Harness リポジトリとリファレンス ドキュメントを参照してください。
ランタイムリファレンス、互換性、セキュリティ、制限事項
DeepSeek ハーネスのアップストリーム ドキュメント
ソースのチェックアウトは開発の場合にのみ必要です。リポジトリ ツールチェーンには Node.js ^22.19.0 または >=24.0.0 が必要です。
git clone https://github.com/pawaca/dsh-edge.git
cd dsh-edge
pnpmインストール
pnpm --dir apps/dsh-edge/standalone install --frozen-lockfile
pnpm 実行チェック
Worker をローカルで実行するには、ローカル Edge セットアップを参照してください。
dsh-edge のバグとインストールの問題については、「問題」で報告してください。
脆弱性は、公的問題ではなく、プライベートなセキュリティ プロセスを通じて報告してください。
リポジトリを変更する場合は、CONTRIBUTING.md および AGENTS.md に従ってください。
マサチューセッツ工科大学サードパーティのコンポーネントとそのライセンスは THIRD_PARTY_NOTICES.md にリストされています。
DeepSeek ハーネスをどこにでも — 1 つのコマンドで永続的なパーソナル コーディング エージェントを Cloudflare Workers にデプロイします。
Readme MIT ライセンス
セキュリティポリシー アクティビティスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Your DeepSeek Harness, anywhere — deploy a persistent personal coding agent to Cloudflare Workers in one command. - pawaca/dsh-edge

GitHub - pawaca/dsh-edge: Your DeepSeek Harness, anywhere — deploy a persistent personal coding agent to Cloudflare Workers in one command. · GitHub
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
pawaca
/
dsh-edge
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
77 Commits 77 Commits Folders and files
.agents .agents .claude .claude .github .github apps/ dsh-edge apps/ dsh-edge docs docs packages/ client/ ui-edge packages/ client/ ui-edge scripts scripts .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .node-version .node-version .oxlintrc.json .oxlintrc.json .pnpmfile.cjs .pnpmfile.cjs AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.i18n.yaml CONTRIBUTING.i18n.yaml CONTRIBUTING.md CONTRIBUTING.md CONTRIBUTING.zh.md CONTRIBUTING.zh.md LICENSE LICENSE README.i18n.yaml README.i18n.yaml README.md README.md README.zh.md README.zh.md SECURITY.i18n.yaml SECURITY.i18n.yaml SECURITY.md SECURITY.md SECURITY.zh.md SECURITY.zh.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md lefthook.yml lefthook.yml package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.client.json tsconfig.base.client.json tsconfig.base.json tsconfig.base.json vitest.config.ts vitest.config.ts View all files Repository files navigation
dsh-edge runs the published DeepSeek Harness Web experience on Cloudflare Workers, so your personal coding agent is available wherever you have a browser. No server to maintain, GitHub repository to connect, or build pipeline to configure.
It keeps the upstream UI, agent loop, model selection, image experience, and Web Search. dsh-edge supplies only the Cloudflare runtime, durable workspace, owner login, and guided installer.
Independent community project: dsh-edge is maintained by pawaca . It is not affiliated with or endorsed by DeepSeek. DeepSeek Harness is the upstream project.
You need Node.js 22.14 or newer and your own DeepSeek API key:
npx dsh-edge install
The installer guides you through every choice and deploys the Worker. You can try it without an existing Cloudflare login, or install it permanently in your own Cloudflare account. No source checkout is required.
Persistent conversations and workspaces from any browser.
DeepSeek V4 Flash, V4 Pro, and the experimental V4 Flash Vision Exp model through the upstream selector.
PNG/JPEG image input and DeepSeek's native Web Search tool.
A persistent /workspace with the native DSH bash tool.
Your own Cloudflare deployment, credentials, and data.
In-place upgrades without a repository or Cloudflare Builds integration.
Path
What you need
Best for
Try now
No existing Cloudflare login; claim within 60 minutes to keep it
Exploring the complete experience with the lowest friction
Keep it
An existing or new Cloudflare account with R2 enabled
A long-lived personal deployment
The default Free — Direct Shell runtime works on Cloudflare Workers Free. The optional Isolated — Dynamic Worker runtime executes commands in a separate Worker and requires Workers Paid. Both modes use the same product UI, conversations, workspace, images, and tools; Direct Shell is a sandboxed shell runtime, not a Linux container.
npx dsh-edge upgrade
The installer finds your Worker and upgrades it in place while preserving its durable data. See the release notes for version-specific details.
dsh-edge is designed for one owner; it does not provide registration, multiple users, roles, or tenant routing.
Your DeepSeek key is stored as a secret in your Cloudflare Worker, and the deployment's durable data stays in your Cloudflare account.
Vision Exp is experimental and account-dependent. Some upstream capabilities are not yet adapted to Edge; see the compatibility matrix for the current status.
dsh-edge wraps exact published DeepSeek Harness packages instead of copying its monorepo. Upstream remains responsible for the Web UI, plugin composition, agent loop, and session protocols; this repository implements the Cloudflare runtime and storage adapters needed to run them at the edge.
For the upstream architecture and plugin APIs, see the DeepSeek Harness repository and reference documentation .
Runtime reference, compatibility, security, and limits
DeepSeek Harness upstream documentation
Source checkout is only required for development. The repository toolchain requires Node.js ^22.19.0 or >=24.0.0 :
git clone https://github.com/pawaca/dsh-edge.git
cd dsh-edge
pnpm install
pnpm --dir apps/dsh-edge/standalone install --frozen-lockfile
pnpm run check
See the local Edge setup to run the Worker locally.
Report dsh-edge bugs and installation problems in Issues .
Report vulnerabilities through the private security process , not a public Issue.
Follow CONTRIBUTING.md and AGENTS.md when changing the repository.
MIT . Third-party components and their licenses are listed in THIRD_PARTY_NOTICES.md .
Your DeepSeek Harness, anywhere — deploy a persistent personal coding agent to Cloudflare Workers in one command.
Readme MIT license Contributing
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
