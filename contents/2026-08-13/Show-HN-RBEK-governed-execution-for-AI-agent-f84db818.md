---
source: "https://github.com/rbekplatform/rbek"
hn_url: "https://news.ycombinator.com/item?id=49284886"
title: "Show HN: RBEK – governed execution for AI agent"
article_title: "GitHub - rbekplatform/rbek: Governed execution for AI agents, workflows and software. · GitHub"
author: "nunomendesfreit"
captured_at: "2026-08-13T12:46:04Z"
capture_tool: "hn-digest"
hn_id: 49284886
score: 1
comments: 0
posted_at: "2026-08-13T12:24:38Z"
tags:
  - hacker-news
  - translated
---

# Show HN: RBEK – governed execution for AI agent

- HN: [49284886](https://news.ycombinator.com/item?id=49284886)
- Source: [github.com](https://github.com/rbekplatform/rbek)
- Score: 1
- Comments: 0
- Posted: 2026-08-13T12:24:38Z

## Translation

タイトル: HN を表示: RBEK – AI エージェントの管理された実行
記事のタイトル: GitHub - rbekplatform/rbek: AI エージェント、ワークフロー、ソフトウェアの管理された実行。 · GitHub
説明: AI エージェント、ワークフロー、ソフトウェアの管理された実行。 - rbekプラットフォーム/rbek

記事本文:
GitHub - rbekplatform/rbek: AI エージェント、ワークフロー、ソフトウェアの管理された実行。 · GitHub
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
アールベクプラットフォーム
/
レーベク
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
18 コミット 18 コミット .devcontainer .devcontainer .github/ workflows .github/ workflows 例 例 src/ rbek_pypi src/ rbek_pypi テスト/ pypi テスト/ pypi .dockerignore .dockerignore .gitignore .gitignore

CHANGELOG.md CHANGELOG.md COTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile INSTALL.md INSTALL.md QUICKSTART.md QUICKSTART.md README.md README.md SECURITY.md SECURITY.md public-repository-manifest.json public-repository-manifest.json pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェント、ワークフロー、ソフトウェアの管理された実行。
RBEK は、エージェントが実行したいことと、実行が許可されている内容を分離します。
RBEK が 10 秒でアクションを管理するのを参照してください。
カール -fsSL https://raw.githubusercontent.com/rbekplatform/rbek/main/examples/real-governed-agent/demo.sh |バッシュ
APIキーがありません。事前に RBEK をインストールしていません。設定はありません。
デモでは、必要に応じてパブリック RBEK CLI をインストールして検証し、その後、
オフラインで管理された証明:
エージェントがアクションを要求する
|
+-- 不正 -> 拒否 -> 実行: いいえ
|
+-- 許可 -> 許可 -> 管理されたドライラン -> 許可
|
+-- 証拠
ターミナルに表示される内容:
不正なアクション ...... 拒否されました
拒否されたアクションの実行 ..... NO
許可されたアクション ....... 許可されています
ガバナンスされた予行演習 ...................... PASS
ゲート認証 ................................ AUTHORIZED
RBEK ポリシーの施行 ..... REAL
RBEK の証拠 ..... 本物
デフォルトのプルーフでは、外部ネットワーク アクションは実行されません。ポリシーの施行
および証拠の生成は、実際の RBEK の動作です。
パブリック GitHub Actions ワークフローは証明を 2 回実行し、次のことを検証します。
決定論的な証拠の概要は両方の実行で同一です。
ブラウザにゼロインストールしたいですか?
上記の GitHub コードスペースで開くを使用して、次を実行します。
cd サンプル/real-governed-agent
./demo.sh
本物の AI + インターネットの道を望んでいますか?
cd サンプル/real-governed-agent
エクスポート OPENAI_API_KEY= " あなたのキー "
./demo.sh --live
ライブ モードは、実際のモデル推論と実際の Open-Meteo 外部アクションを実行します。
t

RBEK が管理する実行境界を通過します。
AI エージェントは何をしたいかを決定できます。実稼働システムには依然として
実際に実行が許可されているものの制御された境界。
RBEK は、アプリケーション ロジックと実際の外部アクションの間に境界を設けます。
エージェント/ワークフロー
↓
実行リクエスト
↓
RBEK ポリシーの承認
↙ ↘
拒否許可
↓
実行する
↓
証拠
これにより、実行ガバナンスがエージェントのフレームワークやモデルから分離されます。
プロバイダーまたはワークフロー エンジン。
カール -fsSL https://releases.rbekplatform.com/cli/stable/install.sh |バッシュ
確認:
rbek-cli --バージョン
現在の公開安定版:
RBEK0.2.0
5 分間のクイックスタート
最小限のローカル RBEK プロジェクトを作成して実行します。
rbek-cli init ./rbek-demo
rbek-cli を実行します。/rbek-demo
または、リポジトリの例を実行します。
bash の例/5-minut-quickstart/run.sh
最初の 5 分間の目標は単純です。RBEK をインストールし、管理された環境を作成します。
ローカル プロジェクトを作成し、RBEK CLI を通じて実行します。
完全な初回実行のウォークスルーについては、QUICKSTART.md を参照してください。
ライブ AI + インターネット バージョンを実行する
ゼロキー証明の後、実際のエージェント パスを実行できます。
cd サンプル/real-governed-agent
エクスポート OPENAI_API_KEY= " あなたのキー "
./demo.sh --live
ライブモードの場合:
OpenAIエージェント
↓
天気、現在
↓
RBEK ポリシーの承認
↓
制御された外部実行
↓
オープンメテオ
↓
領収書+証明書
エージェントは Open-Meteo を直接呼び出しません。外部アクションが通過します
RBEK が管理する実行境界。
参照
例/real-governed-agent/README.md
詳細については。
Developer は、ローカル開発、評価、および実行のためのパブリック CLI エントリ ポイントです。
統合。
有料の商用ライセンスは必要ありません。
チームとエンタープライズの商用アクセスは個別に処理されます。
このリポジトリには、開発者向けドキュメント、サンプル、インストールが含まれています
指導。
RBEK ランタイムは配布されています

d 公式リリースホスト経由:
https://releases.rbekplatform.com
完全なランタイム ソースはこのリポジトリでは公開されていません。
AI エージェント、ワークフロー、ソフトウェアの管理された実行。
セキュリティポリシー アクティビティスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Governed execution for AI agents, workflows and software. - rbekplatform/rbek

GitHub - rbekplatform/rbek: Governed execution for AI agents, workflows and software. · GitHub
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
rbekplatform
/
rbek
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
18 Commits 18 Commits .devcontainer .devcontainer .github/ workflows .github/ workflows examples examples src/ rbek_pypi src/ rbek_pypi tests/ pypi tests/ pypi .dockerignore .dockerignore .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile INSTALL.md INSTALL.md QUICKSTART.md QUICKSTART.md README.md README.md SECURITY.md SECURITY.md public-repository-manifest.json public-repository-manifest.json pyproject.toml pyproject.toml View all files Repository files navigation
Governed execution for AI agents, workflows and software.
RBEK separates what an agent wants to do from what it is authorized to execute .
See RBEK govern an action in 10 seconds
curl -fsSL https://raw.githubusercontent.com/rbekplatform/rbek/main/examples/real-governed-agent/demo.sh | bash
No API key. No prior RBEK installation. No configuration.
The demo installs and validates the public RBEK CLI when needed, then runs an
offline governed proof:
Agent requests an action
|
+-- unauthorized -> DENY -> executed: NO
|
+-- authorized -> ALLOW -> governed dry-run -> AUTHORIZED
|
+-- evidence
What you see in the terminal:
Unauthorized action ............ DENIED
Denied action executed ......... NO
Authorized action .............. ALLOWED
Governed dry-run ............... PASS
Gate authorization ............. AUTHORIZED
RBEK policy enforcement ........ REAL
RBEK evidence .................. REAL
The default proof performs no external network action . Policy enforcement
and evidence generation are real RBEK behavior.
The public GitHub Actions workflow executes the proof twice and verifies that
the deterministic evidence summary is identical across both runs.
Prefer zero-install in the browser?
Use Open in GitHub Codespaces above, then run:
cd examples/real-governed-agent
./demo.sh
Want the real AI + Internet path?
cd examples/real-governed-agent
export OPENAI_API_KEY= " your-key "
./demo.sh --live
Live mode performs real model inference and a real Open-Meteo external action
through the RBEK governed execution boundary.
AI agents can decide what they want to do. Production systems still need a
controlled boundary for what is actually allowed to execute.
RBEK puts that boundary between application logic and real external actions:
agent / workflow
↓
execution request
↓
RBEK policy admission
↙ ↘
DENY ALLOW
↓
execute
↓
evidence
This keeps execution governance separate from the agent framework, model
provider or workflow engine.
curl -fsSL https://releases.rbekplatform.com/cli/stable/install.sh | bash
Verify:
rbek-cli --version
Current public stable:
RBEK 0.2.0
5-minute quickstart
Create and run a minimal local RBEK project:
rbek-cli init ./rbek-demo
rbek-cli run ./rbek-demo
Or run the repository example:
bash examples/5-minute-quickstart/run.sh
The goal of the first five minutes is simple: install RBEK, create a governed
local project and execute it through the RBEK CLI.
See QUICKSTART.md for the complete first-run walkthrough.
Run the live AI + Internet version
After the zero-key proof, you can run the real agent path:
cd examples/real-governed-agent
export OPENAI_API_KEY= " your-key "
./demo.sh --live
In live mode:
OpenAI agent
↓
weather.current
↓
RBEK policy admission
↓
controlled external execution
↓
Open-Meteo
↓
receipt + certification
The agent does not call Open-Meteo directly. The external action goes through
the RBEK governed execution boundary.
See
examples/real-governed-agent/README.md
for details.
Developer is the public CLI entry point for local development, evaluation and
integration.
It does not require a paid commercial entitlement.
Team and Enterprise commercial access are handled separately.
This repository contains Developer documentation, examples and installation
guidance.
The RBEK runtime is distributed through the official release host:
https://releases.rbekplatform.com
The complete runtime source is not published in this repository.
Governed execution for AI agents, workflows and software.
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
