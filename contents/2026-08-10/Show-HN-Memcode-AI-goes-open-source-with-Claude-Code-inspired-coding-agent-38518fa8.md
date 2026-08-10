---
source: "https://github.com/memcode-ai/memcode"
hn_url: "https://news.ycombinator.com/item?id=49239401"
title: "Show HN: Memcode AI goes open source with Claude Code inspired coding agent"
article_title: "GitHub - memcode-ai/memcode: The coding agent that remembers your repo. CLI, gateway, one wire. Self-hostable, MIT. · GitHub"
author: "timerwin"
captured_at: "2026-08-10T05:16:12Z"
capture_tool: "hn-digest"
hn_id: 49239401
score: 2
comments: 0
posted_at: "2026-08-10T05:06:05Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Memcode AI goes open source with Claude Code inspired coding agent

- HN: [49239401](https://news.ycombinator.com/item?id=49239401)
- Source: [github.com](https://github.com/memcode-ai/memcode)
- Score: 2
- Comments: 0
- Posted: 2026-08-10T05:06:05Z

## Translation

タイトル: Show HN: Memcode AI がクロード コードにインスピレーションを得たコーディング エージェントでオープンソースに
記事のタイトル: GitHub - memcode-ai/memcode: リポジトリを記憶するコーディング エージェント。 CLI、ゲートウェイ、1 本のワイヤー。自己ホスト可能、MIT。 · GitHub
説明: リポジトリを記憶するコーディング エージェント。 CLI、ゲートウェイ、1 本のワイヤー。自己ホスト可能、MIT。 - memcode-ai/memcode

記事本文:
GitHub - memcode-ai/memcode: リポジトリを記憶するコーディング エージェント。 CLI、ゲートウェイ、1 本のワイヤー。自己ホスト可能、MIT。 · GitHub
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
メモリコード-ai
/
メモリコード
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
10 コミット 10 コミット .github/ workflows .github/ workflows アセット アセット カタログ カタログ cmd cmd デプロイ/ docker

デプロイ/ docker docs docs ゲートウェイ ゲートウェイ 内部 内部プロトコル プロトコル ツール/ keyprobe tools/ keyprobe .example.env .example.env .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml COMPACTION.md COMPACTION.md HOOKS.md HOOKS.md ライセンス ライセンス README.md README.md ROUTING.md ROUTING.md go.mod go.mod go.sum go.sum main.go main.go models.json models.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
リポジトリを記憶するコーディング エージェント。
ほとんどのコーディング エージェントは、すべてのセッションをゼロから開始します。 memcode は、リポジトリの永続モデルを .memcode に保持します。これには、サブシステム、先週作業した内容、失敗したアプローチとその理由、およびそれを修正した設定が含まれます。長く使えば使うほど、説明する必要はなくなります。
1 つの Go バイナリ、完全なターミナル UI、そしてホストされている memcode ゲートウェイ、独自の API キー、または Ollama のようなローカル エンドポイントなど、既存のあらゆるモデルに対して実行されます。
カール -fsSL https://memcode.ai/install.sh |しー
または Go を使用すると:
github.com/memcode-ai/memcode@latest をインストールしてください
またはソースからビルドします。
git clone https://github.com/memcode-ai/memcode
cd memcode && go build -o memcode 。
次に、リポジトリで memcode を実行します。
MEMCODE_ENDPOINT_URL=http://localhost:11434/v1 memcode # オラマ、ローカル
MEMCODE_ENDPOINT_URL=https://api.openai.com/v1 memcode # OpenAI キー (OPENAI_API_KEY)
MEMCODE_ENDPOINT_URL=https://api.anthropic.com memcode # Anthropic キー (ANTHROPIC_API_KEY)
OPENAI_API_KEY 、 ANTHROPIC_API_KEY 、 GEMINI_API_KEY 、 XAI_API_KEY 、 FIREWORKS_API_KEY 、および友人は、ホストのために自動的に選択されます。名前付きエンドポイントは構成内に存在し、/model はセッション中にモデルを切り替えます。
memcode アカウントを使用すると、すべてのベンダーで 1 つの残高、BYOK 用のキー コンテナー、およびホストされた Web 検索を取得できます。
製品全体がこのリポジトリにあり、ゲートウェイも含まれており、ゲートウェイは

単一のステートレスバイナリ:
cddeploy/docker && cp ../../.example.env .env # トークン + プロバイダー キーを設定します
docker 構成 --build
docs/self-hosting.md を参照してください。 memcode.ai のホスト型サービスは、プライベート コントロール プレーン (アカウント、ベンダー間の 1 つの残高、BYOK ボールト、チーム機能) をトップにして、これと同じコードを実行します。
CLI はエージェントです。すべてのモデルの選択、エスカレーション、リカバリはクライアント側で実行されます。すべてのバックエンドは、1 つの OpenAI 互換ワイヤ (protocol/PROTOCOL.md) を話すプレーンなサーフェスです。 Gateway/ の下のゲートウェイは、計測および型指定されたサービス サーフェスであり、1 つのプロバイダー実装を CLI と共有します。クラウドのみの動作は 1 つのコントロール プレーンの背後にあります。セルフホスト モードでは何も構築されません。
リポジトリを記憶するコーディング エージェント。 CLI、ゲートウェイ、1 本のワイヤー。自己ホスト可能、MIT。
Readme MIT ライセンス アクティビティ カスタム プロパティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

The coding agent that remembers your repo. CLI, gateway, one wire. Self-hostable, MIT. - memcode-ai/memcode

GitHub - memcode-ai/memcode: The coding agent that remembers your repo. CLI, gateway, one wire. Self-hostable, MIT. · GitHub
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
memcode-ai
/
memcode
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
10 Commits 10 Commits .github/ workflows .github/ workflows assets assets catalog catalog cmd cmd deploy/ docker deploy/ docker docs docs gateway gateway internal internal protocol protocol tools/ keyprobe tools/ keyprobe .example.env .example.env .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml COMPACTION.md COMPACTION.md HOOKS.md HOOKS.md LICENSE LICENSE README.md README.md ROUTING.md ROUTING.md go.mod go.mod go.sum go.sum main.go main.go models.json models.json View all files Repository files navigation
The coding agent that remembers your repo.
Most coding agents start every session from zero. memcode keeps a persistent model of your repo in .memcode : the subsystems, what you worked on last week, which approaches failed and why, and the preferences you have corrected it on. The longer you use it, the less you have to explain.
One Go binary, a full terminal UI, and it runs against whatever models you already have: the hosted memcode gateway, your own API keys, or a local endpoint like Ollama.
curl -fsSL https://memcode.ai/install.sh | sh
Or with Go:
go install github.com/memcode-ai/memcode@latest
Or build from source:
git clone https://github.com/memcode-ai/memcode
cd memcode && go build -o memcode .
Then run memcode in a repo.
MEMCODE_ENDPOINT_URL=http://localhost:11434/v1 memcode # Ollama, local
MEMCODE_ENDPOINT_URL=https://api.openai.com/v1 memcode # your OpenAI key (OPENAI_API_KEY)
MEMCODE_ENDPOINT_URL=https://api.anthropic.com memcode # your Anthropic key (ANTHROPIC_API_KEY)
OPENAI_API_KEY , ANTHROPIC_API_KEY , GEMINI_API_KEY , XAI_API_KEY , FIREWORKS_API_KEY , and friends are picked up automatically for their hosts. Named endpoints live in config, and /model switches models mid-session.
With a memcode account you get one balance across every vendor, a key vault for BYOK, and hosted web search:
The whole product is in this repo, gateway included, and the gateway is a single stateless binary:
cd deploy/docker && cp ../../.example.env .env # set a token + a provider key
docker compose up --build
See docs/self-hosting.md . The hosted service at memcode.ai runs this same code with a private control plane on top (accounts, one balance across vendors, the BYOK vault, team features).
The CLI is the agent: all model selection, escalation, and recovery run client-side; every backend is a plain serving surface speaking one OpenAI-compatible wire ( protocol/PROTOCOL.md ). The gateway under gateway/ is that serving surface, metered and typed, sharing one provider implementation with the CLI. Cloud-only behavior sits behind one control-plane seam; self-host mode constructs none of it.
The coding agent that remembers your repo. CLI, gateway, one wire. Self-hostable, MIT.
Readme MIT license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
