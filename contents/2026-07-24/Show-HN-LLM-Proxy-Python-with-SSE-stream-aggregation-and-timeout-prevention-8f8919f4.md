---
source: "https://github.com/lordraw77/llmproxy"
hn_url: "https://news.ycombinator.com/item?id=49033596"
title: "Show HN: LLM Proxy - Python with SSE stream aggregation and timeout prevention"
article_title: "GitHub - lordraw77/llmproxy · GitHub"
author: "lordraw"
captured_at: "2026-07-24T10:53:56Z"
capture_tool: "hn-digest"
hn_id: 49033596
score: 1
comments: 0
posted_at: "2026-07-24T10:38:26Z"
tags:
  - hacker-news
  - translated
---

# Show HN: LLM Proxy - Python with SSE stream aggregation and timeout prevention

- HN: [49033596](https://news.ycombinator.com/item?id=49033596)
- Source: [github.com](https://github.com/lordraw77/llmproxy)
- Score: 1
- Comments: 0
- Posted: 2026-07-24T10:38:26Z

## Translation

タイトル: HN を表示: LLM プロキシ - SSE ストリーム集約とタイムアウト防止を備えた Python
記事タイトル: GitHub - lordraw77/llmproxy · GitHub
説明: GitHub でアカウントを作成して、lordraw77/llmproxy の開発に貢献します。

記事本文:
GitHub - lordraw77/llmproxy · GitHub
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
ロアドロー77
/
llmproxy
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード

「その他のアクション」メニューを開く フォルダーとファイル
11 コミット 11 コミット アセット アセット docs docs img img llmproxy llmproxy スクリプト スクリプト .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md DOCKERHUB.md DOCKERHUB.md Dockerfile Dockerfile ライセンス ライセンスMakefile Makefile README.md README.md ROADMAP.md ROADMAP.md docker-compose.yml docker-compose.yml gunicorn.conf.py gunicorn.conf.py main.py main.py要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
キャッシュ、自動フェイルオーバー、コスト追跡、ローカルとクラウドの AI プロバイダー間のシームレスな統合のための軽量で高性能の LLM プロキシ。
llmproxy は、いくつかの HTTP API をエミュレートする軽量の Flask サーバーです。
人気のあるローカル LLM ランタイム ( Ollama 、
OpenAI /v1 API、および
llama.cpp の llama-server ) と
すべてのリクエストを NVIDIA の OpenAI 互換 API に透過的に転送します
( https://integrate.api.nvidia.com/v1 )。
これにより、すでに Ollama、OpenAI、または llama.cpp を使用できるあらゆるツールが、
NVIDIA がホストするモデルで、クライアント側に変更を加える必要はありません。単に
実際のローカル ランタイムではなく、llmproxy でクライアントを実行します。チャットもカバーしていますが、
補完と埋め込み、ストリーミング、マルチモデル検出をサポート、
オプションの受信認証、一時的なアップストリームエラー時の自動再試行、
ライブ /stats メトリクスとプロセス ダッシュボード。
プロキシが開始され、モデルが公開され、OpenAI 互換の両方のメッセージに応答します。
/v1/chat/completions 呼び出しとネイティブ Ollama ストリーミング /api/chat 呼び出し —
すべてのリクエストは NVIDIA に転送されます。録音のスクリプトは次のとおりです。
scripts/demo.sh (ソースキャスト:assets/demo.cast )。
フローチャート LR
client["あなたのクライアント<br/>(Open WebUI、curl、SDK)"]
プロキシ["llmproxy<br/>(Flask)"]
nvidia["NVIDIA API<br/>integrate.api.nvidia.com/v1"]
クリエ

nt -->|"Ollama / OpenAI / llama.cpp<br/>HTTP リクエスト"|プロキシ
プロキシ -->|"OpenAI リクエスト"|エヌビディア
nvidia -->|"ストリーミング / JSON"|プロキシ
プロキシ -->|"ストリーミング / JSON 応答"|クライアント
読み込み中
ドキュメントのインデックス
文書
説明
概要
llmproxy とは何か、その仕組みとそのアーキテクチャ
インストール
ローカルおよび Docker のセットアップ手順
構成
環境変数とオプション
ロギングとテレメトリ
リクエスト/レスポンスログ、テレメトリ、および設定可能なタイムゾーンクロック
APIリファレンス
すべてのエンドポイントとリクエスト/レスポンスの例
使用例
CURL と一般的なクライアントを使用したエンドツーエンドの例
テスト
scripts/tests.sh ランナー (bash + オプションの TUI)
導入
Docker Compose を使用して本番環境で実行する
トラブルシューティング
よくある問題とその解決方法
クイックスタート
# 1. NVIDIA API キーを構成する
cp .env.example .env
# .env を編集し、NVIDIA_API_KEY を設定します
# 2. Docker Compose (または事前に構築されたイメージ) を使用して実行する
ドッカー構成 -d
# または: docker run -d -p 11434:11434 --env-file .env lordraw/llmproxy:latest
#3. テストしてみる
カール http://localhost:11434/
# → 「オラマが走っています」
事前に構築されたイメージは、Docker Hub で次のように公開されます。
lordraw/llmproxy ;参照してください
Makefile を使用してビルドおよび公開するためのデプロイメント。
MIT ライセンスに基づいてリリースされています。詳細については、LICENSE ファイルを参照してください。
全文。つまり、帰属表示付きで、自由に使用、コピー、変更、配布できます。
そして保証はありません。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to lordraw77/llmproxy development by creating an account on GitHub.

GitHub - lordraw77/llmproxy · GitHub
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
lordraw77
/
llmproxy
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
11 Commits 11 Commits assets assets docs docs img img llmproxy llmproxy scripts scripts .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md DOCKERHUB.md DOCKERHUB.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md ROADMAP.md ROADMAP.md docker-compose.yml docker-compose.yml gunicorn.conf.py gunicorn.conf.py main.py main.py requirements.txt requirements.txt View all files Repository files navigation
A lightweight, high-performance LLM proxy for caching, automatic failover, cost tracking, and seamless integration between local and cloud AI providers.
llmproxy is a lightweight Flask server that emulates the HTTP APIs of several
popular local LLM runtimes ( Ollama , the
OpenAI /v1 API, and
llama.cpp 's llama-server ) and
transparently forwards every request to NVIDIA's OpenAI-compatible API
( https://integrate.api.nvidia.com/v1 ).
This lets any tool that already speaks Ollama, OpenAI, or llama.cpp talk to a
NVIDIA-hosted model without any client-side changes — you simply point the
client at llmproxy instead of at a real local runtime. It covers chat,
completions, and embeddings , supports streaming, multi-model discovery,
optional inbound authentication, automatic retries on transient upstream errors,
and a live /stats metrics & process dashboard.
The proxy starts, exposes the models, and answers both an OpenAI-compatible
/v1/chat/completions call and a native Ollama streaming /api/chat call —
every request forwarded to NVIDIA. The recording is scripted in
scripts/demo.sh (source cast: assets/demo.cast ).
flowchart LR
client["Your client<br/>(Open WebUI, curl, SDK)"]
proxy["llmproxy<br/>(Flask)"]
nvidia["NVIDIA API<br/>integrate.api.nvidia.com/v1"]
client -->|"Ollama / OpenAI / llama.cpp<br/>HTTP request"| proxy
proxy -->|"OpenAI request"| nvidia
nvidia -->|"streaming / JSON"| proxy
proxy -->|"streaming / JSON response"| client
Loading
Documentation index
Document
Description
Overview
What llmproxy is, how it works, and its architecture
Installation
Local and Docker setup instructions
Configuration
Environment variables and options
Logging & Telemetry
Request/response logs, telemetry, and the configurable-timezone clock
API Reference
Every endpoint, with request/response examples
Usage Examples
End-to-end examples with curl and common clients
Testing
The scripts/tests.sh runner (bash + optional TUI)
Deployment
Running in production with Docker Compose
Troubleshooting
Common problems and how to solve them
Quick start
# 1. Configure your NVIDIA API key
cp .env.example .env
# edit .env and set NVIDIA_API_KEY
# 2. Run with Docker Compose (or the prebuilt image)
docker compose up -d
# or: docker run -d -p 11434:11434 --env-file .env lordraw/llmproxy:latest
# 3. Test it
curl http://localhost:11434/
# → "Ollama is running"
The prebuilt image is published on Docker Hub as
lordraw/llmproxy ; see
Deployment for building and publishing with the Makefile .
Released under the MIT License — see the LICENSE file for the
full text. In short: free to use, copy, modify, and distribute, with attribution
and no warranty.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
