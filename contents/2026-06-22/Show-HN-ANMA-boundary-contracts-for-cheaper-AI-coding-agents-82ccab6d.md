---
source: "https://github.com/anma-labs/anma"
hn_url: "https://news.ycombinator.com/item?id=48623765"
title: "Show HN: ANMA, boundary contracts for cheaper AI coding agents"
article_title: "GitHub - anma-labs/anma: Boundary enforcement for AI coding agents — plain-YAML contracts compiled into CLAUDE.md, hooks, and CI checks. · GitHub"
author: "nxy"
captured_at: "2026-06-22T00:34:07Z"
capture_tool: "hn-digest"
hn_id: 48623765
score: 1
comments: 0
posted_at: "2026-06-21T23:41:02Z"
tags:
  - hacker-news
  - translated
---

# Show HN: ANMA, boundary contracts for cheaper AI coding agents

- HN: [48623765](https://news.ycombinator.com/item?id=48623765)
- Source: [github.com](https://github.com/anma-labs/anma)
- Score: 1
- Comments: 0
- Posted: 2026-06-21T23:41:02Z

## Translation

タイトル: Show HN: ANMA、より安価な AI コーディング エージェントの境界契約
記事のタイトル: GitHub - anma-labs/anma: AI コーディング エージェントの境界適用 — CLAUDE.md、フック、および CI チェックにコンパイルされたプレーン YAML コントラクト。 · GitHub
説明: AI コーディング エージェントの境界適用 — CLAUDE.md、フック、および CI チェックにコンパイルされたプレーン YAML コントラクト。 - anma-labs/あんま
HN テキスト: 私が ANMA を構築したのは、安価なモデルではアーキテクチャ ルールが無視されることが多いことに気づいたからです。そこで、ANMA を使用した場合と使用しない場合で、「Claude Haiku 4.5」を使用していくつかのベンチマークを実行しました。 ANMA を使用しない場合、19 実行中 13 で「ルール」が無視され、ANMA を使用すると 20 実行中 0 で「ルール」が無視されました。 「あんま」とは何ですか？
YAML は CLAUDE.md、フック、CI チェックと契約します。 より強力で高価なモデルはどうですか?
彼らは建築規則に従っていました。問題は、より強力なルールを備えた安価なモデルが、コーディングに最適な手頃なデフォルトとなるだろうかということです。

記事本文:
GitHub - anma-labs/anma: AI コーディング エージェントの境界適用 — CLAUDE.md、フック、および CI チェックにコンパイルされたプレーン YAML コントラクト。 · GitHub
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
あんまラボ
/
あんま
公共
通知
通知設定を変更するにはサインインする必要があります

追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
132 コミット 132 コミット .claude .claude .github .github anma anma ベンチマーク ベンチマーク ドキュメント ドキュメント テスト テスト .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md COTRIBUTING.md DECISIONS.md DECISIONS.md LICENSE LICENSE README.md README.md RELEASE.md RELEASE.md SECURITY.md SECURITY.md anma.yaml anma.yaml pyproject.toml pyproject.toml tach.toml tach.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング エージェントの境界強制。 ANMA はプレーンな YAML モジュールに変わります
CLAUDE.md へのコントラクト、フック、およびクロード コードを内部に保持するチェック
あなたのアーキテクチャは、最も重要な場所で目に見えて機能します。
制御されたベンチマーク (Python) では、より安価で高速なモデル (Claude Haiku 4.5)
プレーン リポジトリの 19 回中 13 回の実行で、宣言されたモジュール境界に違反しました。と
ANMA、同じタスクを 20 回実行した結果、違反は 0 回でした (フィッシャーの正確値)
p < 0.0001 )。完全な調査については、docs/BENCHMARKS.md を参照してください。
正直な部分: フロンティアモデル (Opus 4.8) はそれ自体で境界を尊重したので、
ANMA の価値は、より安価なエージェントを運営するための保険と CI/ガバナンスです。
保証 — フロンティアモデルをよりスマートにするわけではありません。
言語: Python、Go、TypeScript (言語: ルート anma.yaml 、
プロジェクトごとに 1 つ）。 Go と TypeScript はモジュール→モジュールの依存関係を強制します。インターフェース
( public: ) 現在、施行は Python のみです。 Go/TS アダプターは検証されています
( anma チェック + フックは、実際のモジュール間違反を検出してブロックします)。で
事前に登録されたフォローアップ (中立的なプロンプト、より難しいシナリオ)、TypeScript は
測定された効果 - コントロール 18/20 対 ANMA 0/20、フィッシャーの正確な p < 0.00001。行く
方向性です

l かつ有意 (10/30 → 0/30、p = 0.0004 ) ですが、制御率は低下しました
事前に登録した0.40フロアを下回っているため、示唆的なものとして報告しますが、まだ
効能も。 Python の見出しはどちらの言語にも当てはまりません。詳細:
概念 § 言語
とベンチマーク。
各モジュールのパブリック インターフェイスと、それが依存するものを宣言します。あんま同期
それを他のすべてにコンパイルするため、エージェントが読み取るアーキテクチャは決して
CI が適用するルールから逸脱する:
anma.yaml プロジェクト構成 (schema_version、source_roots)
ソース/ドメイン/請求/
anma.yaml モジュール コントラクト — すべてのフィールドについては docs/CONCEPTS.md を参照してください
クロードが billing/ を開くと CLAUDE.md (生成) がロードされます。
CLAUDE.md (生成された) アーキテクチャ マップ、マーカー間
.claude/rules/boundaries.md (生成された) 常にロードされる命令型
.claude/hooks/anma_pretooluse.py (生成) は境界を超える編集をブロックします (出口 2)
tach.toml (生成された) エンジン構成 (Go: .go-arch-lint.yml; TS: .dependency-cruiser.cjs)
.github/workflows/anma.yml (生成) CI: ドリフトチェック + 境界チェック
DECISIONS.md 追加専用: 各境界が存在する理由
クイックスタート (60 秒)
pip install anma[tach] # tach バックエンドを推奨します。それなしでも動作します
anma init # scaffolds 契約 + 稼働したアカウント/請求の例
anma sync # CLAUDE.md、ネストされたドキュメント、フック、tach.toml、CI を生成します
あんまチェック # ✓ 境界を尊重
Go または TypeScript の場合、anma init -- language go / でスキャフォールディングします。
anma init -- language typescript (外部バックエンド — go-arch-lint 、
dependency-cruiser — はオプションです。内蔵スキャナはゼロデップフォールバックです)。
完全なチュートリアル: docs/QUICKSTART.md 。
anma init # スキャフォールド契約 + 動作例
anma sync # コントラクトからすべてのアーティファクトを再生成します
anma sync --check # CI ガード: 生成されたアーティファクトがコントラクトから逸脱した場合は失敗します
あんま

check # 境界を強制する (フック/プリコミット/CI)
anma check --warn # 違反を報告するが終了 0 (増分導入)
anma check --json # パイプラインの機械可読出力
終了コード: 0 ok · 1 違反、契約エラー、またはドリフト。
2 つの層: 指導と強制
ANMA は 2 つのレベルで動作し、ベンチマークはそれらが異なる役割を果たすことを示しています。
ガイダンス — 生成されたルートおよびモジュールごとの CLAUDE.md および .claude/rules
アーキテクチャをエージェントのコンテキストに置きます。これが 68% → 0 の原動力となった
結果: モデルは正しい設計に導かれ、悪い設計は行われませんでした。
編集します。
強制 - PreToolUse フックは提案された編集を判断し、返します。
exit 2 は、新しく許可されていないインポートが到着する前にブロックします。同じチェックが実行されます
コミット前およびCI内で。これは編集ガイダンスに適用される保証です
はキャッチされず、どのモデルや人間が差分を書いたかに関係なく。
強制フックが起動することが検証されます (禁止された編集を入力して終了 2 します)。で
ガイダンスがすべての不適切な編集を事前に回避したため、ベンチマークを行う必要はありませんでした。両方
問題;それぞれが何を行うかを正確に示すベンチマークを参照してください。
チームはより安価またはより高速なエージェント (コスト重視のパイプライン、バルク) を実行しています。
タスク、非フロンティアまたは非クロード モデルなど）は、
独自のアーキテクチャ — これが ANMA の舵取りが決定的な場所です。
強制的なアーキテクチャを望む人: CI/プリコミットでの保証
モジュールの境界は、変更を誰が書いたかに関係なく維持されます。
アーキテクチャをガバナンスとして必要とするチーム: 宣言されたインターフェース、所有権 →
CODEOWNER、および黙ってルールから逸脱できないドキュメント。
小規模で詳細に説明されたタスクでのみフロンティア モデルを駆動したことがある場合、ANMA は、
結果を変えることなくターンを追加できます。ベンチマークはそれを明確に示しています。
～800行、

ランタイムなし、DSL なし、1 つの小さな依存関係 (PyYAML) — 組み込み
エンジンにはそれ以上何も必要なく、より高速な外部バックエンド (Python の場合は tach、
Go の場合は go-arch-lint、TypeScript の場合は dependency-cruiser) はすべてオプションです。あ
セキュリティ チームは午後 1 分でツール全体を読むことができます。
ドリフト検出 — anma sync --check が生成された docs/config が落ちた場合に CI が失敗する
契約と同期していない。
増分導入 — anma check --warn およびモジュールごとの deprecated_deps
初日からレッドビルドを行わずに大規模なコードベースを採用しましょう。
ガバナンス — 所有者: モジュールごとに CODEOWNERS を生成します。ソースルート:
モノリポジトリをサポートします。
サプライ チェーン — 署名付きリリース (PyPI Trusted Publishing + 来歴 + SBOM)、
CI、Apache-2.0 の pip-audit。 SECURITY.md を参照してください。
docs/QUICKSTART.md — 最初にブロックされた編集にインストールします
docs/CONCEPTS.md — モデル、コントラクト スキーマ参照、生成されたアーティファクト、エンジン
docs/BENCHMARKS.md — 研究の有無、方法論、および正直な限界
COTRIBUTING.md — 開発セットアップ、テスト、ドッグフード、スキーマ安定性ルール
SECURITY.md · RELEASE.md · CHANGELOG.md
AI コーディング エージェントの境界強制 - CLAUDE.md、フック、および CI チェックにコンパイルされたプレーン YAML コントラクト。
Apache-2.0ライセンス
貢献する
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
4
v0.7.0 — Go と TypeScript
最新の
2026 年 6 月 7 日
+ 3 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Boundary enforcement for AI coding agents — plain-YAML contracts compiled into CLAUDE.md, hooks, and CI checks. - anma-labs/anma

I built ANMA because I noticed that cheaper models would often ignore architecture rules. So I did several benchmarks using "Claude Haiku 4.5" with and without ANMA; without ANMA it ignored the "rules" 13 out of 19 runs, with ANMA, 0 out of 20 runs. What is "ANMA"?
YAML contracts with CLAUDE.md, hooks, and CI checks What about stronger/expensive models?
They followed the architecture rules. The question is, would cheaper models with stronger rules be the best affordable default for coding?

GitHub - anma-labs/anma: Boundary enforcement for AI coding agents — plain-YAML contracts compiled into CLAUDE.md, hooks, and CI checks. · GitHub
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
anma-labs
/
anma
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
132 Commits 132 Commits .claude .claude .github .github anma anma benchmarks benchmarks docs docs tests tests .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md DECISIONS.md DECISIONS.md LICENSE LICENSE README.md README.md RELEASE.md RELEASE.md SECURITY.md SECURITY.md anma.yaml anma.yaml pyproject.toml pyproject.toml tach.toml tach.toml View all files Repository files navigation
Boundary enforcement for AI coding agents. ANMA turns plain-YAML module
contracts into the CLAUDE.md , hooks, and checks that keep Claude Code inside
your architecture — and it measurably works where it matters most.
In a controlled benchmark (Python), a cheaper/faster model (Claude Haiku 4.5)
violated a declared module boundary in 13 of 19 runs of a plain repo. With
ANMA, across 20 runs of the same task it violated it 0 times (Fisher's exact
p < 0.0001 ). See docs/BENCHMARKS.md for the full study, including the
honest part: a frontier model (Opus 4.8) respected the boundary on its own, so
ANMA's value is insurance for running cheaper agents plus a CI/governance
guarantee — not making a frontier model smarter.
Languages: Python, Go, and TypeScript ( language: in the root anma.yaml ,
one per project). Go and TypeScript enforce module→module dependencies; interface
( public: ) enforcement is Python-only today. The Go/TS adapters are validated
( anma check + the hook detect and block real cross-module violations). In a
pre-registered follow-up (neutral prompt, harder scenario), TypeScript shows a
measured effect — control 18/20 vs ANMA 0/20, Fisher's exact p < 0.00001 ; Go
is directional and significant (10/30 → 0/30, p = 0.0004 ) but its control rate fell
below our pre-registered 0.40 floor, so we report it as suggestive, not yet
efficacy . The Python headline is not extrapolated to either language. Details:
CONCEPTS § Languages
and BENCHMARKS .
You declare each module's public interface and what it may depend on. anma sync
compiles that into everything else, so the architecture the agent reads can never
drift from the rules CI enforces:
anma.yaml project config (schema_version, source_roots)
src/domains/billing/
anma.yaml the module contract — see docs/CONCEPTS.md for all fields
CLAUDE.md (generated) loads when Claude opens billing/
CLAUDE.md (generated) architecture map, between markers
.claude/rules/boundaries.md (generated) always-loaded imperative
.claude/hooks/anma_pretooluse.py (generated) blocks a boundary-breaking edit (exit 2)
tach.toml (generated) engine config (Go: .go-arch-lint.yml; TS: .dependency-cruiser.cjs)
.github/workflows/anma.yml (generated) CI: drift check + boundary check
DECISIONS.md append-only: why each boundary exists
Quickstart (60 seconds)
pip install anma[tach] # tach backend recommended; works without it too
anma init # scaffolds contracts + a worked accounts/billing example
anma sync # generates CLAUDE.md, nested docs, hooks, tach.toml, CI
anma check # ✓ boundaries respected
For Go or TypeScript, scaffold with anma init --language go /
anma init --language typescript (the external backends — go-arch-lint ,
dependency-cruiser — are optional; a builtin scanner is the zero-dep fallback).
Full walkthrough: docs/QUICKSTART.md .
anma init # scaffold contracts + a worked example
anma sync # regenerate all artifacts from contracts
anma sync --check # CI guard: fail if generated artifacts drifted from contracts
anma check # enforce boundaries (hook / pre-commit / CI)
anma check --warn # report violations but exit 0 (incremental adoption)
anma check --json # machine-readable output for pipelines
Exit codes: 0 ok · 1 violations, contract errors, or drift.
Two layers: guidance and enforcement
ANMA works at two levels, and the benchmark shows they play different roles:
Guidance — the generated root and per-module CLAUDE.md and .claude/rules
put your architecture in the agent's context. This is what drove the 68% → 0
result: the model was steered to the correct design and didn't attempt a bad
edit.
Enforcement — the PreToolUse hook judges the proposed edit and returns
exit 2 to block any new disallowed import before it lands; the same check runs
at pre-commit and in CI. This is the guarantee that holds for the edits guidance
doesn't catch, and regardless of which model or human wrote the diff.
The enforcement hook is verified to fire (feed it a forbidden edit → exit 2 ); in
the benchmark it never needed to, because guidance pre-empted every bad edit. Both
matter; see the benchmarks for exactly what each one is shown to do.
Teams running cheaper or faster agents (cost-sensitive pipelines, bulk
tasks, non-frontier or non-Claude models) that don't reliably respect an
architecture on their own — this is where ANMA's steering is decisive.
Anyone who wants an enforced architecture: a guarantee in CI/pre-commit that
module boundaries hold no matter who or what wrote the change.
Teams that want architecture as governance : declared interfaces, ownership →
CODEOWNERS, and docs that can't silently drift from the rules.
If you only ever drive a frontier model on small, well-described tasks, ANMA may
add turns without changing outcomes — and the benchmarks say so plainly.
~800 lines, no runtime, no DSL, one small dependency (PyYAML) — the builtin
engine needs nothing more, and the faster external backends ( tach for Python,
go-arch-lint for Go, dependency-cruiser for TypeScript) are all optional. A
security team can read the whole tool in an afternoon.
Drift detection — anma sync --check fails CI if generated docs/config fall
out of sync with the contracts.
Incremental adoption — anma check --warn and per-module deprecated_deps
let a large codebase adopt without a red build on day one.
Governance — owners: per module generates CODEOWNERS ; source_roots:
supports monorepos.
Supply chain — signed releases (PyPI Trusted Publishing + provenance + SBOM),
pip-audit in CI, Apache-2.0. See SECURITY.md .
docs/QUICKSTART.md — install to first blocked edit
docs/CONCEPTS.md — the model, the contract schema reference , generated artifacts, the engine
docs/BENCHMARKS.md — the with/without study, methodology, and honest limits
CONTRIBUTING.md — dev setup, tests, the dogfood, the schema-stability rule
SECURITY.md · RELEASE.md · CHANGELOG.md
Boundary enforcement for AI coding agents — plain-YAML contracts compiled into CLAUDE.md, hooks, and CI checks.
Apache-2.0 license
Contributing
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
4
v0.7.0 — Go & TypeScript
Latest
Jun 7, 2026
+ 3 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
