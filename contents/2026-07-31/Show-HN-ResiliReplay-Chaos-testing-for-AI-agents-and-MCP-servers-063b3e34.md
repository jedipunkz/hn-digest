---
source: "https://github.com/aliengineering-byte/resilireplay"
hn_url: "https://news.ycombinator.com/item?id=49118105"
title: "Show HN: ResiliReplay - Chaos testing for AI agents and MCP servers"
article_title: "GitHub - aliengineering-byte/resilireplay: Crash-test AI agents and MCP servers, replay failures, and turn broken traces into deterministic regression tests. · GitHub"
author: "ali110v"
captured_at: "2026-07-31T01:52:17Z"
capture_tool: "hn-digest"
hn_id: 49118105
score: 1
comments: 0
posted_at: "2026-07-31T01:44:29Z"
tags:
  - hacker-news
  - translated
---

# Show HN: ResiliReplay - Chaos testing for AI agents and MCP servers

- HN: [49118105](https://news.ycombinator.com/item?id=49118105)
- Source: [github.com](https://github.com/aliengineering-byte/resilireplay)
- Score: 1
- Comments: 0
- Posted: 2026-07-31T01:44:29Z

## Translation

タイトル: Show HN: ResiliReplay - AI エージェントと MCP サーバーのカオス テスト
記事のタイトル: GitHub - Aliengineering-byte/resilireplay: AI エージェントと MCP サーバーをクラッシュ テストし、失敗をリプレイし、壊れたトレースを決定論的な回帰テストに変換します。 · GitHub
説明: AI エージェントと MCP サーバーをクラッシュ テストし、失敗を再現し、壊れたトレースを決定論的な回帰テストに変換します。 - エイリアンエンジニアリングバイト/レジリプレイ

記事本文:
GitHub - Aliengineeering-byte/resilireplay: AI エージェントと MCP サーバーをクラッシュ テストし、失敗をリプレイし、壊れたトレースを決定論的な回帰テストに変換します。 · GitHub
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
エイリアンエンジニアリングバイト
/
レジリプレイ
出版

ic
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミット .github .github docs docs 例 例 パッケージ パッケージ シナリオ シナリオ スクリプト スクリプト テスト テスト .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .prettierignore .prettierignore .prettierrc.json .prettierrc.json CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md RELEASE_EVIDENCE.md RELEASE_EVIDENCE.md SECURITY.md SECURITY.md THREAT_MODEL.md THREAT_MODEL.md action.yml action.yml eslint.config.mjs eslint.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json tsconfig.tests.json tsconfig.tests.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ResiliReplay は、AI エージェントと MCP サーバーのクラッシュ テストを行い、失敗を決定論的に再現し、壊れたトレースを回帰テストに変換します。
これは、モデルに依存しない、ローカルファーストの TypeScript ツールキットです。バージョン管理されたイベントを記録し、シード制御された障害を挿入し、宣言された証拠からスコアを回復し、最初の原因となる障害を編集可能なシナリオと実行可能なノード テストにコンパイルします。決定論的パスには、API キー、有料モデル、Docker、外部アカウント、テレメトリ、LLM ジャッジは必要ありません。
ResiliReplay は防御テスト ソフトウェアです。ローカルまたはユーザー所有の MCP ターゲットのみを監査します。レポートまたはバッジは、宣言された 1 つのスイートとバージョンの証拠であり、普遍的なセキュリティ認定ではありません。
要件: Node.js 20 または 22 および pnpm。
git clone https://github.com/aliengineering-byte/resilireplay.git
CD レジリプレイ
ピン

pm install --frozen-lockfile
pnpm ビルド
pnpmデモ
このデモでは、実際の決定論的なローカル サブプロセスを実行し、3 つの障害を挿入し、成功した回復と回復されていない失敗のスコアを付け、すべてのレポート形式を書き込み、失敗したトレースをコンパイルし、生成された回帰テストを実行します。
pnpm デモ:mcp
pnpm exec resilireplay テスト シナリオ
MCP デモでは、意図的に脆弱なおもちゃの標準入出力サーバーと回復力のあるサーバーを監査します。 1 つ目は、予想される安全なカナリアの結果が 2 つあります。 2 番目には何もありません。ブラウザで runs/demo/recovered-report/report.html を開いて、スタンドアロンの回復レポートを調べます。
予想される出力とアセットの再現については、デモ ガイドと検証済みのトランスクリプトを参照してください。
バンドルされた決定論的エージェントを記録し、合格ベースライン レポートを出力します。
pnpm exec resilireplay Record --output runs/agent/trace.jsonl -- ノードの例/deterministic-agent/dist/index.js
pnpm exec resilireplay 再生 --trace 実行/エージェント/trace.jsonl --report-dir 実行/エージェント/レポート
リアル エージェント フレームワークのアダプターは、同じバージョンの TraceEvent オブジェクトを発行します。アダプターガイドを参照してください。 OpenAI 互換のサンプルは変換フィクスチャであり、ライブ プロバイダーを呼び出すとは主張しません。
バンドルされた復元サーバーを標準入出力経由で監査します。
pnpm exec resilireplay mcp Audit --command "node Examples/resilient-mcp-server/dist/index.js" --output runs/mcp-audit
mcp Audit はツール/リスト スキーマをキャプチャし、デフォルトで信頼性_プローブという名前のツールのみを呼び出します。 MCP 呼び出しにはサーバー側の影響がある可能性があるため、ツールの動作を確認した後でのみ --call-tools を渡します。非ループバック ストリーミング可能な HTTP ターゲットにも --allow-remote が必要です。
MCP カオス ガイドでは、制御された障害、トランスポート、認可、および安全なカナリア動作について説明します。
組み込みの障害またはレビュー可能な YAML シナリオを適用します。同じトレース、シナリオ、シードから同じ mut が生成されます。

アクション:
pnpm exec resilireplay inject --trace run/agent/trace.jsonl --scenario malformed-json --seed 42 --output run/agent/failed.jsonl
pnpm exec resilireplay restart --trace runs/agent/failed.jsonl --report-dir runs/agent/failed-report
pnpm exec resilireplay の障害
上記の再生コマンドは意図的に終了します。 1: 最小限のベースラインには、不正な形式の応答が挿入された後の回復イベントがありません。失敗したレポート バンドルはレビューのために引き続き書き込まれます。 CI コマンドで、出口 0 での予想成功シナリオと予想失敗シナリオの両方を検証する必要がある場合は、pnpm exec resilireplay テスト シナリオを使用します。
プロバイダーおよびトランスポートの障害には、制限されたレイテンシー、タイムアウト、429/5xx、リセット、切り捨て、不正な形式の JSON、重複、および古い応答が含まれます。ツールとワークフローの障害には、エラー、権限、破棄可能な欠落ファイル、破損した結果、副作用の重複、ハンドオフの損失、間違った受信者、古い状態、命令の競合、およびループが含まれます。
YAML コントラクトと安全ルールのカスタム障害シナリオを参照してください。
pnpm デモは失敗したトレースを作成します。それを最小化されたフィクスチャ、シナリオ、マニフェスト、および実行可能ノード:test に変換します。
pnpm exec resilireplaygenerate-test --traceruns/demo/failed.jsonl --outputruns/generated-regression
生成されたテストはデフォルトですぐに実行されます。 manifest.json は、ソース トレース、最小化されたフィクスチャ、シナリオ、テストを SHA-256 ハッシュでリンクします。
フローチャート LR
A["JSONL トレースを記録する"] --> B["決定論的フォールトを挿入する"]
B --> C["リプレイ"]
C --> D["スコア回復"]
D --> E["因果関係による失敗を最小限に抑える"]
E --> F["回帰の生成"]
F --> C
読み込み中
サンプル回収レポート
この抜粋は、キーなしのデモからキャプチャされたものです。
ResiliPlay v0.1.0 パス
回復スコア 100/100
完了はい
リカバリセーフ
再試行 1/3
安全規格に準拠
実行ごとに次のものを発行できます。
安定境界は厳密なものです。

er-neutral TraceEvent 、モデル SDK ではありません:
不変条件とパッケージの依存関係については、アーキテクチャを参照してください。
Record は指定したコマンドを正確に実行します。 OS サンドボックスではありません。
自分が所有している、またはテストを許可されている MCP サーバーのみを監査します。
--call-tools の前に MCP スキーマを確認してください。ツールには副作用がある可能性があります。
資格情報形式の値と機密キーがトレース保存前に編集される場合でも、アダプター ソースでシークレットを省略します。
ソース トレースにプライベート アプリケーション データが含まれている場合は、レポートを機密テストの証拠として扱います。
ファイルシステム障害では所有された一時ディレクトリが使用され、出力パスは包含チェックされ、生成されたプロセスには期限とクリーンアップがあり、呼び出し元が意図的に別のホストを選択しない限り、リスナーはループバックにバインドされます。
信頼できないコマンドまたはリモート MCP ターゲットを実行する前に、SECURITY.md および THREAT_MODEL.md を読んでください。
ストリーミング可能な HTTP サポートは、stdio よりも統合範囲が狭いです。
MCP ツールの呼び出しはサーバー側に影響を与える可能性があります。
因果関係の最小化は、アダプターがparentIdとcauseIdを提供する場合に最も強力になります。
ストリーミング プロバイダーの出力は、v0.1.0 では応答イベントに集約されます。
レポートのハッシュは、信頼性ではなく、つながりと整合性を証明します。マニフェストは署名されていません。
パッケージは npm に公開されません。
既知の制限事項とロードマップを参照してください。
pnpm品質
複合アクションを使用します。
- 使用：aliengineering-byte/resilireplay@v0.1.0
付き:
シナリオ : シナリオ
ResiliReplay は Apache-2.0 ライセンスを取得しています。貢献は決定的であり、制限があり、テストによってカバーされる必要があります。 CONTRIBUTING.md、CODE_OF_CONDUCT.md、起動リファレンス、リリース証拠、および v0.1.0 リリースを参照してください。
AI エージェントと MCP サーバーをクラッシュ テストし、失敗を再現し、壊れたトレースを決定論的な回帰テストに変換します。
Readme Apache-2.0 ライセンスの行動規範
セキュリティポリシー アクティビティスター
フォーク数 0 レポート リポジトリ

リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Crash-test AI agents and MCP servers, replay failures, and turn broken traces into deterministic regression tests. - aliengineering-byte/resilireplay

GitHub - aliengineering-byte/resilireplay: Crash-test AI agents and MCP servers, replay failures, and turn broken traces into deterministic regression tests. · GitHub
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
aliengineering-byte
/
resilireplay
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits .github .github docs docs examples examples packages packages scenarios scenarios scripts scripts tests tests .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .prettierignore .prettierignore .prettierrc.json .prettierrc.json CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md RELEASE_EVIDENCE.md RELEASE_EVIDENCE.md SECURITY.md SECURITY.md THREAT_MODEL.md THREAT_MODEL.md action.yml action.yml eslint.config.mjs eslint.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json tsconfig.tests.json tsconfig.tests.json vitest.config.ts vitest.config.ts View all files Repository files navigation
ResiliReplay crash-tests AI agents and MCP servers, replays failures deterministically, and converts broken traces into regression tests.
It is a model-agnostic, local-first TypeScript toolkit: record versioned events, inject seed-controlled faults, score recovery from declared evidence, and compile the first causal failure into an editable scenario plus an executable Node test. The deterministic path needs no API key, paid model, Docker, external account, telemetry, or LLM judge.
ResiliReplay is defensive testing software. Audit only local or user-owned MCP targets. A report or badge is evidence for one declared suite and version, not a universal security certification.
Requirements: Node.js 20 or 22 and pnpm.
git clone https://github.com/aliengineering-byte/resilireplay.git
cd resilireplay
pnpm install --frozen-lockfile
pnpm build
pnpm demo
The demo runs a real deterministic local subprocess, injects three faults, scores a successful recovery and an unrecovered failure, writes every report format, compiles the failed trace, and executes the generated regression test.
pnpm demo:mcp
pnpm exec resilireplay test scenarios
The MCP demo audits an intentionally vulnerable toy stdio server and a resilient one. The first has two expected safe-canary findings; the second has none. Open runs/demo/recovered-report/report.html in a browser to inspect the standalone recovery report.
See the demo guide and verified transcript for expected output and asset reproduction.
Record the bundled deterministic agent and emit a passing baseline report:
pnpm exec resilireplay record --output runs/agent/trace.jsonl -- node examples/deterministic-agent/dist/index.js
pnpm exec resilireplay replay --trace runs/agent/trace.jsonl --report-dir runs/agent/report
Adapters for a real agent framework emit the same versioned TraceEvent objects. See the adapter guide ; the OpenAI-compatible example is a translation fixture and does not claim to call a live provider.
Audit the bundled resilient server over stdio:
pnpm exec resilireplay mcp audit --command "node examples/resilient-mcp-server/dist/index.js" --output runs/mcp-audit
mcp audit captures tools/list schemas and calls only a tool named reliability_probe by default. Pass --call-tools only after reviewing tool behavior, because MCP calls may have server-side effects. Non-loopback Streamable HTTP targets also require --allow-remote .
The MCP chaos guide covers controlled faults, transports, authorization, and safe-canary behavior.
Apply a built-in fault or reviewable YAML scenario. The same trace plus scenario plus seed produces the same mutation:
pnpm exec resilireplay inject --trace runs/agent/trace.jsonl --scenario malformed-json --seed 42 --output runs/agent/failed.jsonl
pnpm exec resilireplay replay --trace runs/agent/failed.jsonl --report-dir runs/agent/failed-report
pnpm exec resilireplay faults
The replay command above intentionally exits 1: the minimal baseline has no recovery event after the injected malformed response. It still writes the failed report bundle for review. Use pnpm exec resilireplay test scenarios when a CI command should verify both expected-pass and expected-failure scenarios with exit 0.
Provider and transport faults include bounded latency, timeout, 429/5xx, reset, truncation, malformed JSON, duplicates, and stale responses. Tool and workflow faults cover errors, permissions, disposable missing files, corrupt results, side-effect duplication, handoff loss, wrong recipients, stale state, conflicting instructions, and loops.
See custom fault scenarios for the YAML contract and safety rules.
pnpm demo creates a failed trace. Convert it into a minimized fixture, scenario, manifest, and executable node:test :
pnpm exec resilireplay generate-test --trace runs/demo/failed.jsonl --output runs/generated-regression
The generated test executes immediately by default. manifest.json links the source trace, minimized fixture, scenario, and test with SHA-256 hashes.
flowchart LR
A["Record JSONL trace"] --> B["Inject deterministic fault"]
B --> C["Replay"]
C --> D["Score recovery"]
D --> E["Minimize causal failure"]
E --> F["Generate regression"]
F --> C
Loading
Sample recovery report
This excerpt is captured from the no-key demo:
ResiliReplay v0.1.0 PASS
Recovery score 100/100
Completion yes
Recovery safe
Retries 1/3
Safety compliant
Each run can emit:
The stable boundary is a strict, provider-neutral TraceEvent , not a model SDK:
Read the architecture for invariants and package dependencies.
record executes the exact command you supply; it is not an OS sandbox.
Audit only MCP servers you own or are authorized to test.
Review MCP schemas before --call-tools ; tools can have side effects.
Omit secrets at the adapter source even though credential-shaped values and sensitive keys are redacted before trace storage.
Treat reports as sensitive test evidence if source traces contain private application data.
Filesystem faults use owned temporary directories, output paths are containment-checked, spawned processes have deadlines and cleanup, and listeners bind to loopback unless a caller deliberately chooses another host.
Read SECURITY.md and THREAT_MODEL.md before running untrusted commands or remote MCP targets.
Streamable HTTP support has less integration coverage than stdio.
MCP tool calls may have server-side effects.
Causal minimization is strongest when adapters provide parentId and causeId .
Streaming provider output is aggregated into response events in v0.1.0.
Report hashes prove linkage and integrity, not authenticity; manifests are unsigned.
Packages are not published to npm.
See known limitations and the roadmap .
pnpm quality
Use the composite action:
- uses : aliengineering-byte/resilireplay@v0.1.0
with :
scenarios : scenarios
ResiliReplay is Apache-2.0 licensed. Contributions should be deterministic, bounded, and covered by tests. See CONTRIBUTING.md , CODE_OF_CONDUCT.md , the launch reference , release evidence , and the v0.1.0 release .
Crash-test AI agents and MCP servers, replay failures, and turn broken traces into deterministic regression tests.
Readme Apache-2.0 license Code of conduct
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
