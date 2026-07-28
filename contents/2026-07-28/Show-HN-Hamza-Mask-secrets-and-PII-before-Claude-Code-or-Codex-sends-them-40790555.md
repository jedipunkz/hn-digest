---
source: "https://github.com/softcane/hamza"
hn_url: "https://news.ycombinator.com/item?id=49090241"
title: "Show HN: Hamza Mask secrets and PII before Claude Code or Codex sends them"
article_title: "GitHub - softcane/hamza: An egress gate for CLI coding agents. Masks secrets and customer data in the request body before it reaches the model provider, and the agent keeps working. · GitHub"
author: "pradeep1177"
captured_at: "2026-07-28T22:04:45Z"
capture_tool: "hn-digest"
hn_id: 49090241
score: 1
comments: 0
posted_at: "2026-07-28T21:35:54Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Hamza Mask secrets and PII before Claude Code or Codex sends them

- HN: [49090241](https://news.ycombinator.com/item?id=49090241)
- Source: [github.com](https://github.com/softcane/hamza)
- Score: 1
- Comments: 0
- Posted: 2026-07-28T21:35:54Z

## Translation

タイトル: クロード コードまたはコーデックスが送信する前に、HN: ハムザ マスクの秘密と PII を表示します。
記事のタイトル: GitHub - Softcane/hamza: CLI コーディング エージェントの出口ゲート。リクエスト本文内のシークレットと顧客データがモデル プロバイダーに到達する前にマスクされ、エージェントは動作を継続します。 · GitHub
説明: CLI コーディング エージェントの出口ゲート。リクエスト本文内のシークレットと顧客データがモデル プロバイダーに到達する前にマスクされ、エージェントは動作を継続します。 - ソフトケイン/ハムザ

記事本文:
GitHub - Softcane/hamza: CLI コーディング エージェントの出口ゲート。リクエスト本文内のシークレットと顧客データがモデル プロバイダーに到達する前にマスクされ、エージェントは動作を継続します。 · GitHub
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
ソフトケーン
/
ハムズ

ある
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット .github .github アセット アセット ディテクタ/ プレシディオ ディテクタ/ プレシディオ エンボイ エンボイ フェイクプロバイダ フェイクプロバイダ grafana grafana プロメテウス プロメテウス 品質 品質 src src .dockerignore .dockerignore .gitignore .gitignore CONTEXT.md CONTEXT.md Dockerfile Dockerfileライセンス ライセンス通知 通知 README.md README.md SECURITY.md SECURITY.md THIRD-PARTY-LICENSES.md THIRD-PARTY-LICENSES.md docker-compose.ci.yml docker-compose.ci.yml docker-compose.yml docker-compose.yml pom.xml pom.xml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ハムザという名前は、映画の秘密工作員にインスピレーションを得たものです。
ドゥランダル 。
Hamza は Claude Code と Codex の代理店です。検出された秘密をマスクし、
Anthropic または OpenAI にプロンプトを送信する前に、承認されたタイプの個人データを送信します。
ハムザクロードコードデモ.mp4
インポート ジョブをデバッグするエージェントは、CSV、.env ファイル、およびアプリケーションを読み取る可能性があります。
ログ。そのコンテキストには、顧客データと資格情報が含まれる場合があります。ハムザマスク
プロンプトの残りの部分はそのままにして、検出された値を表示します。
変更前:patient=priya.fixture@example.com key=AKIAABCDEFGHIJKLMNOP
変更後: 患者=<EMAIL_482191> キー=<SECRET_730044>
同じ値はリクエスト内で同じプレースホルダーを取得します。これにより、モデルは
元の値を受け取らずに参照に従います。
なぜこれが組織の問題なのか
既存のセキュリティ ツールでは、コーディングによって組み立てられた最終リクエストを認識できない可能性があります。
エージェント。
開発者は、どのファイルが分からないまま、エージェントにインポート ジョブのデバッグを依頼する場合があります。
エージェントが読み上げます。その後、顧客の記録や資格情報がネットワークから離れる可能性があります
承認された AI サービスへの通常のリクエスト内。
これにより仕事が生まれます

セキュリティおよびコンプライアンス チーム向け。ローテーションが必要になる場合があります。
資格情報、どのレコードがネットワークから流出したかを調査するか、契約を変更する
with the AI vendor. Hamza は検出された値をマスクすることでその露出を減らします
リクエストが終了する前に。
Hamza は、アクセス制御、保持ルール、または法的合意に代わるものではありません。
フローチャート LR
A["クロード コードまたはコーデックス"] --> B["ハムザ プロキシ"]
B --> C[「秘密と個人データを見つける」]
C --> D["設定されたアクションを適用する"]
D -->|"mask"| E[「検出値を置き換える」]
D -->|"allow"| F["Anthropic or OpenAI"]
D -->|"block"| G["Stop the request"]
E --> F
読み込み中
Hamza は、Claude Code または Codex からサポートされているリクエスト本文を検査します。
AI サービスに送信されたテキストを検索します。
秘密スキャナー、登録値検出器、および Presidio がそれを検査します。
text.
ルールに応じて、Hamza は検出結果を記録、マスク、またはブロックします。
マスキング ルールの場合、検出された値を次のようなプレースホルダーに置き換えます。
<SECRET_1> または <EMAIL_1> 。
Hamza は、アクション、検出器、およびバイト数を含む監査レコードも書き込みます。
レコードにはプロンプト テキストまたは一致する値が含まれていません。
検出器
Finds
Secret scanner
クラウド、ソース管理、パッケージレジストリ、SaaS 認証情報
Presidio
電子メール、電話、ペイメントカード、IP、その他の承認されたデータタイプ
Hamza が認識すべき顧客の価値観を登録することもできます。
Presidio エンティティの完全なリストとしきい値は次のとおりです。
detecter/presidio/approved-profile.json 。
HAMZA_POSTURE=MASK docker 構成 -d --build
最初の起動では、Presidio の読み込みに 1 分かかる場合があります。
エクスポート ANTHROPIC_BASE_URL=http://127.0.0.1:10000
クロード
セッション間でエクスポートを保持したい場合は、シェル プロファイルにエクスポートを追加します。
クロード コードは、env セクションで ANTHROPIC_BASE_URL も受け入れます。
settings.json .
これを ~/.codex/config.toml に追加します。
[ model_providers . hamza ]
nam

e = "ハムザ"
Base_url = " http://127.0.0.1:10000/backend-api/codex "
Wire_api = " 応答 "
require_openai_auth = true
[ プロフィール .ハムザ]
モデルプロバイダー = " ハムザ "
次に、プロファイルを使用して Codex を起動します。
コーデックス --プロファイル ハムザ
サービス
住所
サービス
http://127.0.0.1:10000
Claude Code および Codex で使用されるプロキシ
http://127.0.0.1:3000
グラファナダッシュボード
http://127.0.0.1:8080/actuator/プロメテウス
プロメテウスのメトリクス
Presidio は Docker ネットワーク内で実行され、ホスト ポートがありません。
HAMZA_PRESIDIO_ENABLED=false \
HAMZA_POSTURE=マスク \
docker 構成 -d --build --scale presidio=0
ビルドとテスト
Hamza には Java 25 と Maven 3.9.11 が必要です。
mvn -B クリーン検証
ドッカービルド 。
ドメイン モデルについては CONTEXT.md を参照してください。
完全なテスト スイートの AGENTS.md。
CLI コーディング エージェントの出口ゲート。リクエスト本文内のシークレットと顧客データがモデル プロバイダーに到達する前にマスクされ、エージェントは動作を継続します。
Readme Apache-2.0 ライセンス セキュリティ ポリシー
セキュリティポリシー アクティビティスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

An egress gate for CLI coding agents. Masks secrets and customer data in the request body before it reaches the model provider, and the agent keeps working. - softcane/hamza

GitHub - softcane/hamza: An egress gate for CLI coding agents. Masks secrets and customer data in the request body before it reaches the model provider, and the agent keeps working. · GitHub
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
softcane
/
hamza
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits .github .github assets assets detector/ presidio detector/ presidio envoy envoy fake-provider fake-provider grafana grafana prometheus prometheus quality quality src src .dockerignore .dockerignore .gitignore .gitignore CONTEXT.md CONTEXT.md Dockerfile Dockerfile LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md THIRD-PARTY-LICENSES.md THIRD-PARTY-LICENSES.md docker-compose.ci.yml docker-compose.ci.yml docker-compose.yml docker-compose.yml pom.xml pom.xml View all files Repository files navigation
The name Hamza is inspired by the undercover operative in
Dhurandhar .
Hamza is a proxy for Claude Code and Codex. It masks detected secrets and
approved types of personal data before sending prompts to Anthropic or OpenAI.
hamza-claude-code-demo.mp4
An agent debugging an import job may read a CSV, an .env file, and application
logs. That context can contain customer data and credentials. Hamza masks
detected values while leaving the rest of the prompt intact:
Before: patient=priya.fixture@example.com key=AKIAABCDEFGHIJKLMNOP
After: patient=<EMAIL_482191> key=<SECRET_730044>
The same value gets the same placeholder within a request. This lets the model
follow references without receiving the original value.
Why this is an organization problem
Existing security tools may not see the final request assembled by a coding
agent.
A developer may ask an agent to debug an import job without knowing which files
the agent will read. Customer records or credentials can then leave the network
inside a normal request to an approved AI service.
This creates work for security and compliance teams. They may need to rotate a
credential, investigate which records left the network, or change the contract
with the AI vendor. Hamza reduces that exposure by masking detected values
before the request leaves.
Hamza does not replace access controls, retention rules, or legal agreements.
flowchart LR
A["Claude Code or Codex"] --> B["Hamza proxy"]
B --> C["Find secrets and personal data"]
C --> D["Apply the configured action"]
D -->|"mask"| E["Replace the detected value"]
D -->|"allow"| F["Anthropic or OpenAI"]
D -->|"block"| G["Stop the request"]
E --> F
Loading
Hamza inspects supported request bodies from Claude Code or Codex.
It finds the text sent to the AI service.
The secret scanner, registered-value detector, and Presidio inspect that
text.
Depending on the rule, Hamza records, masks, or blocks the finding.
For masking rules, it replaces the detected value with a placeholder such as
<SECRET_1> or <EMAIL_1> .
Hamza also writes an audit record with the action, detector, and byte counts.
The record contains no prompt text or matched values.
Detector
Finds
Secret scanner
Cloud, source-control, package-registry, and SaaS credentials
Presidio
Email, phone, payment-card, IP, and other approved data types
You can also register customer values that Hamza should recognize.
The full Presidio entity list and thresholds are in
detector/presidio/approved-profile.json .
HAMZA_POSTURE=MASK docker compose up -d --build
The first start may take a minute while Presidio loads.
export ANTHROPIC_BASE_URL=http://127.0.0.1:10000
claude
Add the export to your shell profile if you want to keep it across sessions.
Claude Code also accepts ANTHROPIC_BASE_URL in the env section of
settings.json .
Add this to ~/.codex/config.toml :
[ model_providers . hamza ]
name = " hamza "
base_url = " http://127.0.0.1:10000/backend-api/codex "
wire_api = " responses "
requires_openai_auth = true
[ profiles . hamza ]
model_provider = " hamza "
Then start Codex with the profile:
codex --profile hamza
Services
Address
Service
http://127.0.0.1:10000
Proxy used by Claude Code and Codex
http://127.0.0.1:3000
Grafana dashboard
http://127.0.0.1:8080/actuator/prometheus
Prometheus metrics
Presidio runs inside the Docker network and has no host port.
HAMZA_PRESIDIO_ENABLED=false \
HAMZA_POSTURE=MASK \
docker compose up -d --build --scale presidio=0
Build and test
Hamza requires Java 25 and Maven 3.9.11.
mvn -B clean verify
docker build .
See CONTEXT.md for the domain model and
AGENTS.md for the complete test suite.
An egress gate for CLI coding agents. Masks secrets and customer data in the request body before it reaches the model provider, and the agent keeps working.
Readme Apache-2.0 license Security policy
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
