---
source: "https://github.com/saidsef/argocd-ai-assistant"
hn_url: "https://news.ycombinator.com/item?id=48885756"
title: "Argocd-AI-Assistant"
article_title: "GitHub - saidsef/argocd-ai-assistant: Argo CD AI Assistant - An AI-powered chatbot extension for Argo CD · GitHub"
author: "saidsef"
captured_at: "2026-07-12T23:43:13Z"
capture_tool: "hn-digest"
hn_id: 48885756
score: 2
comments: 0
posted_at: "2026-07-12T23:00:39Z"
tags:
  - hacker-news
  - translated
---

# Argocd-AI-Assistant

- HN: [48885756](https://news.ycombinator.com/item?id=48885756)
- Source: [github.com](https://github.com/saidsef/argocd-ai-assistant)
- Score: 2
- Comments: 0
- Posted: 2026-07-12T23:00:39Z

## Translation

タイトル: Argocd-AI-アシスタント
記事のタイトル: GitHub - Saidsef/argocd-ai-assistant: Argo CD AI Assistant - Argo CD 用の AI 搭載チャットボット拡張機能 · GitHub
説明: Argo CD AI Assistant - Argo CD 用の AI 搭載チャットボット拡張機能 - Saidsef/argocd-ai-assistant

記事本文:
GitHub - Saidsef/argocd-ai-assistant: Argo CD AI Assistant - Argo CD 用の AI を活用したチャットボット拡張機能 · GitHub
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
セフ
/
argocd-ai-アシスタント
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション

オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
88 コミット 88 コミット .github .github docs docs 例/種類 例/種類 src src .gitignore .gitignore .readthedocs.yaml .readthedocs.yaml CONTRIBUTING.md CONTRIBUTING.md LICENSE.md LICENSE.md README.md README.md mkdocs.yml mkdocs.yml package.json package.json tsconfig.json tsconfig.json webpack.config.js webpack.config.js 糸.ロック 糸.ロック すべてのファイルを表示 リポジトリ ファイルのナビゲーション
リソース ビューに [アシスタント] タブを追加する Argo CD UI 拡張機能。 Kubernetes リソースを開いて自然言語で質問すると、拡張機能によって各クエリがコンテキストで強化されます。
リソース マニフェスト - Argo CD によって提供されるライブ オブジェクト。
イベント - Argo CD API から取得され、自動的に添付されます。
ログ (オプション) - ガイド付きフロー (ポッド、デプロイメント、ステートフルセット、ジョブ、ロールアウト) 経由で、単一のコンテナーから。
単一の汎用プロバイダーが OpenAI 互換のチャット完了 API を使用するため、ローカル推論サーバー (Ollama など)、 vLLM 、OpenAI 、Azure OpenAI 、またはその他の OpenAI 互換エンドポイントなどのバックエンドと連携して動作します。 CORS を回避するために、トラフィックは Argo CD プロキシ拡張機能を介してルーティングされます。
Argo CD サーバーからアクセス可能な OpenAI 互換 API を備えた実行中の LLM バックエンド
# 依存関係をインストールします (React ピアの依存関係に問題があるため、--force が必要です)
糸インストール --force
# 本番ビルド
糸の実行ビルド
# 開発ビルド
糸実行ビルド開発
リリース
リリースは自動化されています。メインにマージすると (CI が通過すると)、ワークフローが変更を分類し、バージョンを上げ、タグ付けし、拡張子 tar を使用して GitHub リリースを公開します。手動によるバージョンアップはありません。詳細については、「ビルドとパッケージ化」を参照してください。
ビルド、パッケージ化、ホストの詳しい手順については、導入ガイドを参照してください。

を実行し、拡張機能インストーラーを介して拡張機能を Argo CD にインストールします。 Argo CD Operator、コミュニティ Helm チャート、生のマニフェストについて説明します。
使い捨ての種類クラスターで拡張機能をエンドツーエンドで試すには (組み込みのモック LLM を使用し、GPU や API キーは必要ありません)、examples/kind/ のハーネスを使用します。
./examples/kind/setup.sh raw # または: helm |オペレーター
実行が成功すると、Argo CD がインストールされ、ソースから構築されたこの拡張機能がインストールされ、== 8 が成功、0 が失敗した == でフル パス (ストリーミングされたプロキシ リクエストを含む) が検証されます。 example/kind/README.md を参照してください。
完全なドキュメントは docs/ ディレクトリにあり、 MkDocs Material で構築されています。
ライブドキュメント: argocd-ai-assistant.readthedocs.io
ソースとリリース: github.com/saidsef/argocd-ai-assistant 。
プル リクエストは歓迎です。コントリビューション ガイドを参照してください。
Argo CD AI Assistant - Argo CD 用の AI を活用したチャットボット拡張機能
Apache-2.0ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
85
v11.0.2
最新の
2026 年 7 月 12 日
+ 84 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Argo CD AI Assistant - An AI-powered chatbot extension for Argo CD - saidsef/argocd-ai-assistant

GitHub - saidsef/argocd-ai-assistant: Argo CD AI Assistant - An AI-powered chatbot extension for Argo CD · GitHub
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
saidsef
/
argocd-ai-assistant
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
88 Commits 88 Commits .github .github docs docs examples/ kind examples/ kind src src .gitignore .gitignore .readthedocs.yaml .readthedocs.yaml CONTRIBUTING.md CONTRIBUTING.md LICENSE.md LICENSE.md README.md README.md mkdocs.yml mkdocs.yml package.json package.json tsconfig.json tsconfig.json webpack.config.js webpack.config.js yarn.lock yarn.lock View all files Repository files navigation
An Argo CD UI extension that adds an Assistant tab to the resource view. Open any Kubernetes resource, ask questions about it in natural language, and the extension enriches each query with context:
Resource manifest - the live object, provided by Argo CD.
Events - fetched from the Argo CD API and attached automatically.
Logs (optional) - from a single container, via a guided flow (Pod, Deployment, StatefulSet, Job, Rollout).
A single generic provider speaks the OpenAI-compatible chat completions API , so it works with any backend that does: a local inference server (e.g. Ollama), vLLM , OpenAI , Azure OpenAI , or any other OpenAI-compatible endpoint. Traffic is routed through the Argo CD Proxy Extension to avoid CORS.
A running LLM backend with an OpenAI-compatible API accessible from the Argo CD server
# Install dependencies (requires --force due to React peer dependency quirks)
yarn install --force
# Production build
yarn run build
# Development build
yarn run build-dev
Release
Releases are automated: on merge to main (once CI passes), a workflow classifies the change, bumps the version, tags it, and publishes a GitHub Release with the extension tar. No manual version bumping. See Build and Package for details.
See the Deployment Guide for step-by-step instructions on building, packaging, hosting, and installing the extension into Argo CD via the Extension Installer. It covers the Argo CD Operator, the community Helm chart, and raw manifests.
To try the extension end to end on a throwaway kind cluster - with a built-in mock LLM, no GPU or API key required - use the harness in examples/kind/ :
./examples/kind/setup.sh raw # or: helm | operator
A successful run installs Argo CD, installs this extension built from source, and verifies the full path (including a streamed proxy request) with == 8 passed, 0 failed == . See examples/kind/README.md .
Full documentation is in the docs/ directory and built with MkDocs Material .
Live docs: argocd-ai-assistant.readthedocs.io
Source and releases: github.com/saidsef/argocd-ai-assistant .
Pull requests welcome - see the Contribution Guide .
Argo CD AI Assistant - An AI-powered chatbot extension for Argo CD
Apache-2.0 license
Code of conduct
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
85
v11.0.2
Latest
Jul 12, 2026
+ 84 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
