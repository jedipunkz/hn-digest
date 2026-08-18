---
source: "https://github.com/keploy/keploy"
hn_url: "https://news.ycombinator.com/item?id=49346337"
title: "A VCR to test your backend – record/replay K8s env as deterministic sandboxes"
article_title: "GitHub - keploy/keploy: Open-source platform for creating safe, isolated production sandboxes for API, integration, and E2E testing. · GitHub"
image: "https://repository-images.githubusercontent.com/449649393/3df356de-a8d9-43a5-925b-a75a795af9e3"
author: "keploy"
captured_at: "2026-08-18T15:22:47Z"
capture_tool: "hn-digest"
hn_id: 49346337
score: 1
comments: 0
posted_at: "2026-08-18T14:42:02Z"
tags:
  - hacker-news
  - translated
---

# A VCR to test your backend – record/replay K8s env as deterministic sandboxes

- HN: [49346337](https://news.ycombinator.com/item?id=49346337)
- Source: [github.com](https://github.com/keploy/keploy)
- Score: 1
- Comments: 0
- Posted: 2026-08-18T14:42:02Z

## Translation

タイトル: バックエンドをテストするための VCR – K8s 環境を決定論的なサンドボックスとして記録/再生します
Article title: GitHub - keploy/keploy: Open-source platform for creating safe, isolated production sandboxes for API, integration, and E2E testing. · GitHub
説明: API、統合、E2E テスト用の安全で分離された運用サンドボックスを作成するためのオープンソース プラットフォーム。 - ケプロイ/ケプロイ

記事本文:
GitHub - keploy/keploy: API、統合、E2E テスト用の安全で隔離された運用サンドボックスを作成するためのオープンソース プラットフォーム。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
1,500 コミット 1,500 コミット フォルダーとファイル
.claude/スキル .claude/スキル

ls .github .github Adopters Adopters cli cli config config docs docs pkg pkg testing/ e2e テスト/ e2e tools/ lint/ no_timestamp_in_parser tools/ lint/ no_timestamp_in_parser utils utils .cz.toml .cz.toml .gitignore .gitignore .golangci.yml .golangci.yml .pre-commit-config.yaml .pre-commit-config.yaml .releaserc.json .releaserc.json AGENTS.md AGENTS.md CITATION.cff CITATION.cff CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md DEBUG.md DEBUG.md Dockerfile Dockerfile Dockerfile.runtime Dockerfile.runtime HACKTOBERFEST_GUIDE.md HACKTOBERFEST_GUIDE.md ライセンス ライセンス README.md README.md READMEes-Es.md READMEes-Es.md READMEfr-FR.md READMEfr-FR.md READMEja-JP.md READMEja-JP.md SECURITY.md SECURITY.md エントリポイント.sh エントリポイント.sh go.mod go.mod go.sum go.sum gon.json gon.json goreleaser.yaml goreleaser.yaml keploy.sh keploy.sh main.go main.go oss-pledge.json oss-pledge.json package-lock.json package-lock.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
⚡️ ユーザートラフィックからの API テストは単体テストよりも高速です ⚡️
🌟 90% のテスト カバレッジを実現する AI 世代の開発者にとって必須のツール 🌟
Keploy は、単体テストよりも速くテストとデータモックを自動生成する、開発者中心の API および統合テスト ツールです。
API 呼び出し、データベース クエリ、ストリーミング イベントを記録し、それらをテストとして再生します。 Keploy は内部で eBPF を使用してネットワーク層でトラフィックをキャプチャしますが、ユーザーにとっては完全にコードレスで言語に依存しません。
🐰 面白い事実: Keploy はそれ自体をテストに使用します。おしゃれなカバレッジバッジをチェックしてください:
keploy Record を使用してアプリを実行するだけです。実際の API + 統合フローは、テストおよびモックとして自動的にキャプチャされます。 (Keploy は内部で eBPF を使用してトラフィックをキャプチャするため、SDK を追加したりコードを変更したりする必要はありません。)
📹 複雑な Flo の記録と再生

うーん
Keploy は、複雑な分散 API フローをモックおよびスタブとして記録および再生できます。これは、テスト用に非常に軽量のタイムマシンを使用するようなもので、時間を大幅に節約できます。
👉 録画と再生に関するドキュメントを読む
🐇 完全なインフラ仮想化 (HTTP モックを超えて)
HTTP エンドポイントをモックするだけのツールとは異なり、Keploy はデータベース (Postgres、MySQL、MongoDB)、ストリーミング/キュー (Kafka、RabbitMQ)、外部 API などを記録します。
それらを決定的に再生するため、インフラストラクチャを再プロビジョニングせずにテストを実行できます。
👉 インフラ仮想化に関するドキュメントを読む
あなたが開発者であれば、おそらくステートメントとブランチ カバレッジを気にしているでしょう。Keploy がそれを計算します。
QA の場合は、API スキーマとビジネス ユースケースの範囲に重点を置きます。Keploy はそれも計算します。このようにして、報道はもはや主観的ではなくなります。
🤖 AI を使用して API カバレッジを拡大
Keploy は、既存の記録、Swagger/OpenAPI スキーマを使用して、境界値、欠落/余分なフィールド、間違った型、順序外れのシーケンス、再試行/タイムアウトを検出します。
これは、API スキーマ、ステートメント、およびブランチ カバレッジを拡張するのに役立ちます。
🌐 CI/CD 統合: ローカルの CLI、CI パイプライン (Jenkins、Github Actions..)、または Kubernetes クラスター全体など、好きな場所でモックを使用してテストを実行します。続きを読む
🎭 多目的モック: Keploy で生成されたモックをサーバー テストとして使用することもできます。
📊 レポート: API、統合、ユニット、および e2e カバレッジの統合レポート。CI または PR で直接洞察を得ることができます。
🖥️ コンソール: 記録されたテストとモックを表示、管理、デバッグするための開発者向けのコンソール。
⏱️ 時間のフリーズ: 実行中にシステム時間をフリーズすることで、テストを決定的に再実行します。続きを読む
📚 モック レジストリ: チームや環境全体でモックを管理、再利用、バージョン管理するための一元化されたレジストリ。続きを読む
c

URL --silent -O -L https://keploy.io/install.sh && ソース install.sh
2. テストケースを記録する
Keploy でアプリを起動して、実際の API 呼び出しをテストとモックに変換します。
keploy レコード -c " CMD_TO_RUN_APP "
Python の例:
keploy レコード -c " python main.py "
3. テストの実行
外部依存関係を持たずにオフラインでテストを実行します。
keploy test -c " CMD_TO_RUN_APP " --lay 10
リソース
言語とフレームワーク (任意のスタック)
Keploy はネットワーク層 (eBPF) でインターセプトするため、あらゆる言語、フレームワーク、またはランタイムで動作し、SDK は必要ありません。
注: 一部の依存関係は、プロトコルと解析がオープンソースではないため、本質的にオープンソースではありません。 Keploy エンタープライズではサポートされていません。
ライブデモを予約する / エンタープライズサポート
ガイド付きウォークスルー、専用サポート、またはエンタープライズ展開の計画支援が必要ですか?
カレンダーへの招待をご希望ですか?メールに空き状況を記載してください。すぐにカレンダーへの招待状を送信します。
📘 ドキュメント — ドキュメント全体を探索する
💬 Slack コミュニティ — 会話に参加する
📢 ブログ — 記事と最新情報を読む
新人でも経験者でも、あなたの意見は重要です。コードを提供したり、問題を報告したり、フィードバックを共有したりして、Keploy の改善にご協力ください。
一緒に、最新のアプリケーションのためのより優れたテスト ツールを構築しましょう。
API、統合、E2E テスト用の安全で隔離された本番サンドボックスを作成するためのオープンソース プラットフォーム。
Readme Apache-2.0 ライセンスの行動規範
セキュリティ ポリシー このリポジトリを引用する アクティビティ カスタム プロパティ スター
2.3k フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Open-source platform for creating safe, isolated production sandboxes for API, integration, and E2E testing. - keploy/keploy

GitHub - keploy/keploy: Open-source platform for creating safe, isolated production sandboxes for API, integration, and E2E testing. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
1,500 Commits 1,500 Commits Folders and files
.claude/ skills .claude/ skills .github .github adopters adopters cli cli config config docs docs pkg pkg tests/ e2e tests/ e2e tools/ lint/ no_timestamp_in_parser tools/ lint/ no_timestamp_in_parser utils utils .cz.toml .cz.toml .gitignore .gitignore .golangci.yml .golangci.yml .pre-commit-config.yaml .pre-commit-config.yaml .releaserc.json .releaserc.json AGENTS.md AGENTS.md CITATION.cff CITATION.cff CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md DEBUG.md DEBUG.md Dockerfile Dockerfile Dockerfile.runtime Dockerfile.runtime HACKTOBERFEST_GUIDE.md HACKTOBERFEST_GUIDE.md LICENSE LICENSE README.md README.md READMEes-Es.md READMEes-Es.md READMEfr-FR.md READMEfr-FR.md READMEja-JP.md READMEja-JP.md SECURITY.md SECURITY.md entrypoint.sh entrypoint.sh go.mod go.mod go.sum go.sum gon.json gon.json goreleaser.yaml goreleaser.yaml keploy.sh keploy.sh main.go main.go oss-pledge.json oss-pledge.json package-lock.json package-lock.json View all files Repository files navigation
⚡️ API tests faster than unit tests, from user traffic ⚡️
🌟 The must-have tool for developers in the AI-Gen era for 90% test coverage 🌟
Keploy is a developer‑centric API and integration testing tool that auto‑generates tests and data‑mocks faster than unit tests.
It records API calls, database queries, and streaming events — then replays them as tests. Under the hood, Keploy uses eBPF to capture traffic at the network layer, but for you it’s completely code‑less and language‑agnostic .
🐰 Fun fact: Keploy uses itself for testing! Check out our swanky coverage badge:
Just run your app with keploy record . Real API + integration flows are automatically captured as tests and mocks. (Keploy uses eBPF under the hood to capture traffic, so you don’t need to add any SDKs or modify code.)
📹 Record and Replay complex Flows
Keploy can record and replay complex, distributed API flows as mocks and stubs. It's like having a very light-weight time machine for your tests—saving you tons of time!
👉 Read the docs on record-replay
🐇 Complete Infra‑Virtualization (beyond HTTP mocks)
Unlike tools that only mock HTTP endpoints, Keploy records databases (Postgres, MySQL, MongoDB), streaming/queues (Kafka, RabbitMQ), external APIs, and more.
It replays them deterministically so you can run tests without re‑provisioning infra.
👉 Read the docs on infra virtualisation
If you’re a developer , you probably care about statement and branch coverage — Keploy calculates that for you.
If you’re a QA , you focus more on API schema and business use‑case coverage — Keploy calculates that too. This way coverage isn’t subjective anymore.
🤖 Expand API Coverage using AI
Keploy uses existing recordings, Swagger/OpenAPI Schema to find: boundary values, missing/extra fields, wrong types, out‑of‑order sequences, retries/timeouts.
This helps expand API Schema, Statement, and Branch Coverage.
🌐 CI/CD Integration: Run tests with mocks anywhere you like—locally on the CLI, in your CI pipeline (Jenkins, Github Actions..) , or even across a Kubernetes cluster. Read more
🎭 Multi-Purpose Mocks : You can also use Keploy-generated Mocks, as server Tests!
📊 Reporting: Unified reports for API, integration, unit, and e2e coverage with insights directly in your CI or PRs.
🖥️ Console: A developer-friendly console to view, manage, and debug recorded tests and mocks.
⏱️ Time Freezing: Deterministically replay tests by freezing system time during execution. Read more
📚 Mock Registry: Centralized registry to manage, reuse, and version mocks across teams and environments. Read more
curl --silent -O -L https://keploy.io/install.sh && source install.sh
2. Record Test Cases
Start your app under Keploy to convert real API calls into tests and mocks.
keploy record -c " CMD_TO_RUN_APP "
Example for Python:
keploy record -c " python main.py "
3. Run Tests
Run tests offline without external dependencies.
keploy test -c " CMD_TO_RUN_APP " --delay 10
Resources
Languages & Frameworks (Any stack)
Because Keploy intercepts at the network layer (eBPF) , it works with any language, framework, or runtime —no SDK required.
Note: Some of the dependencies are not open-source by nature because their protocols and parsings are not open-sourced. It's not supported in Keploy enterprise.
Book a Live Demo / Enterprise Support
Want a guided walkthrough, dedicated support, or help planning enterprise rollout?
Prefer a calendar invite? Mention your availability in the email—we’ll send a calendar invite right away.
📘 Documentation — Explore the full docs
💬 Slack Community — Join the conversation
📢 Blog — Read articles and updates
Whether you're new or experienced, your input matters. Help us improve Keploy by contributing code, reporting issues, or sharing feedback.
Together, let's build better testing tools for modern applications.
Open-source platform for creating safe, isolated production sandboxes for API, integration, and E2E testing.
Readme Apache-2.0 license Code of conduct
Security policy Cite this repository Activity Custom properties Stars
2.3k forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
