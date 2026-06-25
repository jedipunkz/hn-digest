---
source: "https://github.com/ctxgov/ctxgov"
hn_url: "https://news.ycombinator.com/item?id=48678976"
title: "Show HN: CtxGov – see what instructions your AI agent inherits before it runs"
article_title: "GitHub - ctxgov/ctxgov: Local-first read-only tools for agent context, memory, and governance evaluation before agents act. · GitHub"
author: "LuxBennu"
captured_at: "2026-06-25T21:32:33Z"
capture_tool: "hn-digest"
hn_id: 48678976
score: 2
comments: 0
posted_at: "2026-06-25T20:47:38Z"
tags:
  - hacker-news
  - translated
---

# Show HN: CtxGov – see what instructions your AI agent inherits before it runs

- HN: [48678976](https://news.ycombinator.com/item?id=48678976)
- Source: [github.com](https://github.com/ctxgov/ctxgov)
- Score: 2
- Comments: 0
- Posted: 2026-06-25T20:47:38Z

## Translation

タイトル: HN を表示: CtxGov – AI エージェントが実行前に継承する命令を確認する
記事のタイトル: GitHub - ctxgov/ctxgov: エージェントが動作する前にエージェントのコンテキスト、メモリ、ガバナンスを評価するためのローカルファースト読み取り専用ツール。 · GitHub
説明: エージェントが動作する前にエージェントのコンテキスト、メモリ、ガバナンスを評価するためのローカルファースト読み取り専用ツール。 - ctxgov/ctxgov

記事本文:
GitHub - ctxgov/ctxgov: エージェントが動作する前にエージェントのコンテキスト、メモリ、ガバナンスを評価するためのローカルファースト読み取り専用ツール。 · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
ctxgov
/
ctxgov
公共
通知
あなた

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
80 コミット 80 コミット .github/ ISSUE_TEMPLATE .github/ ISSUE_TEMPLATE docs docs 例 例 release release schemas/ json schemas/ json scripts scripts src/ ctxgov src/ ctxgov .gitignore .gitignore CHANGELOG.md CHANGELOG.md ライセンス ライセンスREADME.md README.md TRADEMARK.md TRADEMARK.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
CtxGov は、診断と診断を行うためのローカルファーストの研究開発者ツールです。
エージェントが動作する前に、エージェントのコンテキストとメモリ状態の影響を制御します。
公開製品/パッケージ バージョン: 0.9.0 。
公開証拠アーティファクト: release/v0.9.0/ 。
実際の OSS セマンティック インベントリ: release/v0.9.0/state-of-agent-context/ 。
過去のリリース パックはアーカイブ素材であり、現在の製品ではありません
物語。
v0.9.0 パブリック パッケージは、ローカル ガバナンス評価 CLI サーフェスをサポートします。
# 独自のリポジトリで実行します。これは読み取り専用であり、.ctxvault 状態は書き込まれません。
ctxgov 変更ゲートチェック --root 。 --形式の概要
# 2 つの明示的なローカル ツリーを比較し、セマンティックな Change Gate レポートを出力します。
ctxgov change-gate-check --baseline-root 例/change-gate-public-preview/baseline --head-root 例/change-gate-public-preview/head --format 概要
追加のローカル チェック:
ctxgov 継続性コンパイルの例/session-continuity-public-preview/saved-goal-trace.synthetic.json
ctxgov 継続性レンダリングの例/session-continuity-public-preview/saved-goal-trace.synthetic.json
ctxgov continuity apply --mode ドライランの例/session-continuity-public-preview/saved-goal-trace.synthetic.json
ctxgov メモリ Xray 検証リリース/v0.7.0/memory-xray-l1-public-preview/memory-xray-l1-examples-pack.json
ctxgovチャ

nge-gate-federate --base-path 例/federation-public-preview --no-git-required
ctxgov oss-case-study-preview --target-name mem0 --repo-url https://github.com/mem0ai/mem0 --pinned-ref 366945965df43aa7084be98d1b5073b62a20b431 --source-path example/oss-case-study-public-preview/mem0-source.md
ctxgov oss-efficiency 評価 --manifest release/v0.8.0/oss-efficiency-raw-telemetry-manifest.json
ctxgov ガバナンス-リプレイ --trace 例/governance-replay-public-preview/governance-replay-trace.synthetic.json
ctxgov forensics-timeline --fixture release/v0.8.0/forensics-public-preview-fixture.json
ctxgov forensics-trace --fixture release/v0.8.0/forensics-public-preview-fixture.json --finding-id Finding-public-authority-001
ctxgov forensics-gaps --fixture release/v0.8.0/forensics-public-preview-fixture.json
他の開発面はこのパブリック パッケージ境界の外側にあり、そうではありません。
ソース配布物に含まれています。
明示的にローカルに保存されたトレース ファイル 1 つ。
オプションの Memory X-Ray JSON レポート形状の例。
Change Gate とフェデレーションの明示的なローカル リポジトリ パス。
OSS ケーススタディ プレビュー用に明示的に保存されたローカルのパブリック ソースの抜粋。
OSS Efficiency 手法の受信用に明示的に保存された生のテレメトリ。
選択されたパブリック OSS リポジトリのパスレスなエージェント コンテキストの状態の受信。
明示的に保存されたローカル ガバナンス リプレイ トレース。
フォレンジック プレビュー用の明示的なローカル フィクスチャ JSON。
ctxvault.session-continuity-sidecar/v0 JSON オブジェクト。
人間が読める次のセッション継続パケット。
ドライランまたはサンドボックスのみの適用結果。
メモリー X 線検証の領収書。
読み取り専用のセマンティック Change Gate レポート。
読み取り専用のフェデレーション レポート。
OSS ケーススタディの決定プレビュー。
OSS の生のテレメトリ方法論の受領書。
経路のない実際の OSS セマンティック インベントリ レポート。
ガバナンス・リプレイ

制限されたカバレッジ数で検索します。
フォレンジックのタイムライン、トレース、またはギャップ レポート。
以前: 次のエージェント セッションでは、長いトランスクリプト、古い要約、または
不明瞭なソースサポートによる曖昧なハンドオフ。
変更後: CtxGov は、明示的に保存されたトレースをソースに基づく継続性に変換します。
ブロックされた効果、副作用の境界、および破棄によるロールバックを含む証拠
次のセッションがそれを使用する前にセマンティクスを変更します。
ローカルのパブリック パッケージ ゲートを実行します。
python3 スクリプト/run_public_package_checks.py
ゲートはネットワーク呼び出し、プロバイダー/モデル呼び出し、パブリック書き込み、
スケジューラ、デーモン、またはターゲット リポジトリの書き込み。一時的なものを作成する可能性があります
ローカル Evidence Core ストアを使用して、CAS/SQLite のコミットとリプレイの動作を検証します。
継承された ctxgov continuity apply --mode サンドボックス コマンドは、
ローカルサンドボックスファイル。 Evidence Core ライブラリ呼び出しは明示的なローカルのみを書き込みます
呼び出し元が選択した証拠ストアのパス。
ctxgov 継続性適用 --mode ドライラン | サンドボックス
ctxgov oss-効率の評価|検証
他のコマンドはパブリック パッケージには含まれません。実際のターゲットの書き込み、
ネットワーク/API 呼び出し、パブリケーションの自動化、スケジューラー、デーモン、ランタイム アダプター
実行、およびプロバイダ/モデルの実行はこのリリースの対象外です。
ランタイムまたはワークフローの結果の比較。
外部からの取り込みまたはコミュニティの牽引力。
認証またはリスク保証。
自動出版パイプライン。
安定したプロトコルまたは安定した MGP。
チェックインされた ctxvault.* スキーマ識別子は安定したレガシーコントラクトです
現在のブランド主張ではなく、文字列です。
エージェントが動作する前にエージェントのコンテキスト、メモリ、ガバナンスを評価するためのローカルファーストの読み取り専用ツール。
ctxgov.github.io/ctxgov/
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
0
フォーク
レポートリポジトリ
リリース
3
CtxGov v0.9.0
最新の
2026 年 6 月 25 日
+2リリース

s
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Local-first read-only tools for agent context, memory, and governance evaluation before agents act. - ctxgov/ctxgov

GitHub - ctxgov/ctxgov: Local-first read-only tools for agent context, memory, and governance evaluation before agents act. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
ctxgov
/
ctxgov
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
80 Commits 80 Commits .github/ ISSUE_TEMPLATE .github/ ISSUE_TEMPLATE docs docs examples examples release release schemas/ json schemas/ json scripts scripts src/ ctxgov src/ ctxgov .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md TRADEMARK.md TRADEMARK.md pyproject.toml pyproject.toml View all files Repository files navigation
CtxGov is a local-first research and developer tool for diagnosing and
governing agent context and memory-state influence before an agent acts.
Public product/package version: 0.9.0 .
Public evidence artifact: release/v0.9.0/ .
Real OSS semantic inventory: release/v0.9.0/state-of-agent-context/ .
Historical release packs are archive material, not the current product
narrative.
The v0.9.0 public package supports local governance evaluation CLI surfaces:
# Run on your own repository. This is read-only and writes no .ctxvault state.
ctxgov change-gate-check --root . --format summary
# Diff two explicit local trees and print a semantic Change Gate report.
ctxgov change-gate-check --baseline-root examples/change-gate-public-preview/baseline --head-root examples/change-gate-public-preview/head --format summary
Additional local checks:
ctxgov continuity compile examples/session-continuity-public-preview/saved-goal-trace.synthetic.json
ctxgov continuity render examples/session-continuity-public-preview/saved-goal-trace.synthetic.json
ctxgov continuity apply --mode dry-run examples/session-continuity-public-preview/saved-goal-trace.synthetic.json
ctxgov memory-xray validate release/v0.7.0/memory-xray-l1-public-preview/memory-xray-l1-examples-pack.json
ctxgov change-gate-federate --base-path examples/federation-public-preview --no-git-required
ctxgov oss-case-study-preview --target-name mem0 --repo-url https://github.com/mem0ai/mem0 --pinned-ref 366945965df43aa7084be98d1b5073b62a20b431 --source-path examples/oss-case-study-public-preview/mem0-source.md
ctxgov oss-efficiency evaluate --manifest release/v0.8.0/oss-efficiency-raw-telemetry-manifest.json
ctxgov governance-replay --trace examples/governance-replay-public-preview/governance-replay-trace.synthetic.json
ctxgov forensics-timeline --fixture release/v0.8.0/forensics-public-preview-fixture.json
ctxgov forensics-trace --fixture release/v0.8.0/forensics-public-preview-fixture.json --finding-id finding-public-authority-001
ctxgov forensics-gaps --fixture release/v0.8.0/forensics-public-preview-fixture.json
Other development surfaces are outside this public package boundary and are not
included in the source distribution.
one explicit local saved trace file;
optional Memory X-Ray JSON report-shape examples.
explicit local repository paths for Change Gate and Federation;
an explicit saved local public-source excerpt for OSS Case Study Preview;
explicit saved raw telemetry for OSS Efficiency methodology receipts;
pathless State Of Agent Context receipts for selected public OSS repositories;
an explicit saved local Governance Replay trace;
explicit local fixture JSON for Forensics preview.
a ctxvault.session-continuity-sidecar/v0 JSON object;
a human-readable next-session continuity packet;
a dry-run or sandbox-only apply result;
a Memory X-Ray validation receipt.
a read-only semantic Change Gate report;
a read-only Federation report;
an OSS case-study decision preview;
an OSS raw telemetry methodology receipt;
a pathless real-OSS semantic inventory report;
a Governance Replay result with bounded coverage counts;
a Forensics timeline, trace, or gap report.
Before: the next agent session may inherit a long transcript, stale summary, or
ambiguous handoff with unclear source support.
After: CtxGov turns the explicit saved trace into source-backed continuity
evidence with blocked effects, side-effect boundaries, and rollback-by-discard
semantics before the next session uses it.
Run the local public package gate:
python3 scripts/run_public_package_checks.py
The gate performs no network calls, provider/model calls, public writes,
schedulers, daemons, or target repository writes. It may create a temporary
local Evidence Core store to verify CAS/SQLite commit and replay behavior.
The inherited ctxgov continuity apply --mode sandbox command may write one
local sandbox file. Evidence Core library calls write only explicit local
evidence-store paths chosen by the caller.
ctxgov continuity apply --mode dry-run|sandbox
ctxgov oss-efficiency evaluate|validate
No other commands are included in the public package. Real target writes,
network/API calls, publication automation, schedulers, daemons, runtime adapter
execution, and provider/model execution are outside this release.
comparative runtime or workflow outcome;
external uptake or community traction;
certification or risk assurance;
automated publication pipeline;
stable protocol or stable MGP;
The checked-in ctxvault.* schema identifiers are stable legacy contract
strings, not a current branding claim.
Local-first read-only tools for agent context, memory, and governance evaluation before agents act.
ctxgov.github.io/ctxgov/
Topics
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
3
CtxGov v0.9.0
Latest
Jun 25, 2026
+ 2 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
