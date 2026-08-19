---
source: "https://github.com/akshaypimprikar/pragma"
hn_url: "https://news.ycombinator.com/item?id=49364871"
title: "Pragma, a Claude Code pipeline for iOS (94 merged PRs)"
article_title: "GitHub - akshaypimprikar/pragma: Claude Code slash commands for agent-driven iOS development - spec to release in 8 stages · GitHub"
image: "https://opengraph.githubassets.com/1571d821d2a1ed5245dddcf7f53b853d168f8df5cc5ec9edcc18347c07b195d9/akshaypimprikar/pragma"
author: "akshaypimprikar"
captured_at: "2026-08-19T18:19:52Z"
capture_tool: "hn-digest"
hn_id: 49364871
score: 1
comments: 0
posted_at: "2026-08-19T17:56:22Z"
tags:
  - hacker-news
  - translated
---

# Pragma, a Claude Code pipeline for iOS (94 merged PRs)

- HN: [49364871](https://news.ycombinator.com/item?id=49364871)
- Source: [github.com](https://github.com/akshaypimprikar/pragma)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T17:56:22Z

## Translation

タイトル: Pragma、iOS 用のクロード コード パイプライン (94 個のマージされた PR)
記事タイトル: GitHub - akshaypimprikar/pragma: エージェント駆動 iOS 開発用のクロード コード スラッシュ コマンド - 8 段階でリリースされる仕様 · GitHub
説明: エージェント駆動 iOS 開発用のクロード コード スラッシュ コマンド - 8 段階でリリースされる仕様 - akshaypimprikar/pragma

記事本文:
GitHub - akshaypimprikar/pragma: エージェント駆動 iOS 開発用のクロード コード スラッシュ コマンド - 8 段階でリリースされる仕様 · GitHub
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
アクシャイピンプリカル
/
プラグマ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
121 コミット 121 コミット フォルダーとファイル
.claude-plugin .claude-plugin .claude .claude .github .github docs/ スクリーンショット docs/ スクリーンショット scaffold/ .github/ ワークフロー scaffold/ 。

github/ workflows scripts scripts skill/ deterministic-pr-gates skill/ deterministic-pr-gates CHANGELOG.md CHANGELOG.md COTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
You approve twice. Claude ships the rest.
エージェント時代の完全な iOS 開発足場 - エージェント コマンド、CI の適用、セットアップの自動化が相互に接続されているため、1 人のエンジニアがチーム規模で出荷できます。
FinanceTracker で実証済み — 本番環境の SwiftUI + SwiftData アプリは、初日から完全にこのパイプラインに基づいて構築されており、仕様、計画、PR は最初のコミットに遡ります。
モックアップではありません。これは、/spec → /plan → /feature → /gates → /review の 80 以上のマージされた PR が実際に生成したものです。 FinanceTracker の README に設定されている完全なスクリーンショット。
クイック スタート · 得られるもの · パイプライン · なぜスペック モードではないのか? · コマンド · CI レイヤー · メモリレイヤー · カスタマイズ · 貢献
推奨 — クローンやシェル スクリプトを使用せずに、Claude Code プラグインとしてインストールします。
クロード コード内、iOS プロジェクトのリポジトリ ルート:
/プラグイン マーケットプレイス add akshaypimprikar/pragma
/プラグインインストール pragma@pragma
/pragma:init MyApp
/pragma:init は scripts/setup.sh と同じことを行います。コマンド、コンテキスト ファイル、CI ワークフロー、およびサポート スクリプトをコピーし、アプリ名全体を置き換えます。ただし、 CLAUDE.md のアーキテクチャとキー制約の内容についてインタビューし、同じ回答から .claude/context/invariants.md をシードします。両方を後で入力するテンプレートとして残すのではありません。
代替案 — クローンを作成して、セットアップ スクリプトを直接実行します。
git clone https://github.com/akshaypimprikar/pragma
cd pragma
./scripts/setup.sh MyApp /path/to/your-ios-project
これにより、同じファイルがコピーされ、アプリ名が置き換えられますが、CLAUDE.md と invariants.md はテンプレートとして残ります。事前に自分で入力してください。

鉱石の実行/機能。
Then, either way, kick off your first feature:
/spec "describe your feature idea"
得られるもの
Three layers installed into your project:
各レイヤーは独立しています。3 つすべてを採用するか、コマンドのみを採用します。
フローチャート TD
A([💡 Idea]):::dim --> B
B["/spec\n✓ you approve"]:::human --> C
C["/plan\n✓ you approve"]:::human --> D
D["/feature"]:::auto --> E
E["/gates"]:::auto --> F
F([PR opened]):::dim --> G & H & CI1
CI1["CI · pr-checks\nui-tests"]:::ci --> I
G["/review"]:::auto --> I
H["/test"]:::auto --> I
I([merge to develop]):::dim -.->|next feature| B
I --> J["/release"]:::auto --> K([tag v*.*.*]):::dim --> CI2
CI2["CI · リリース\nリリース ビルド + GitHub リリース"]:::ci
BUG([バグレポート]):::dim --> BF["/bugfix"]:::auto --> BG["/gates"]:::auto --> BP([PR]):::dim --> BR["/review"]:::auto --> BM([merge]):::dim
classDef 人間の塗りつぶし:#3d2800、ストローク:#fbbf24、色:#fbbf24
classDef 自動塗りつぶし:#0a1f14、ストローク:#34d399、カラー:#34d399
クラス定義ディムフィル:#161b22、ストローク:#30363d、カラー:#8b949e
classDef ci 塗りつぶし:#0d1117、ストローク:#58a6ff、カラー:#58a6ff
読み込み中
/spec の後と /plan の後、2 回承認します。他のすべてのステップは、エージェントまたは自動化された CI です。
Martin Fowler の「コーディング エージェント ユーザーのためのハーネス エンジニアリング」(2026 年 4 月) では、エージェント = モデル + ハーネスを組み立て、ハーネスをフィードフォワード ガイド (エージェントが動作する前に操作する) とフィードバック センサー (保守性、アーキテクチャの適合性、動作の正しさ全体でエージェントのその後の動作を検証する) に分割しています。 This pipeline maps directly onto that taxonomy:
ファウラー氏は、ほとんどのエージェント ハーネスについて、行動の正しさセンサーを「部屋の中の象、まだ未解決」と呼んでいます。 /test と /gates のカバレッジ ゲートは、そのセンサーに対するこのパイプラインの具体的な試行です。
Cursor / Windsurf / Copilot の組み込みスペック MOD だけではダメな理由

え？
すべての主要な AI コーディング ツールには、Cursor の Plan Mode、Windsurf の Cascade、GitHub Copilot ワークスペース、GitHub Spec Kit、BMAD-METHOD など、何らかの仕様駆動型開発が組み込まれています。当然の質問: なぜ IDE にすでに組み込まれているものを使用するのではなく、別のスキャフォールドを使用するのでしょうか?
pragma が実行でき、spec モードだけでは実行できない 3 つのこと:
エージェントによる強制だけでなく、CI による強制も行われます。 /gates は PR が開く前にローカルで実行されます。同じチェックは、別のプロンプトでエージェントを再実行してもスキップできない強制として、GitHub Actions ( pr-checks.yml 、 ui-tests.yml ) で個別に再実行されます。スペック モードは計画を生成します。強制層に配線されないため、エージェント自体は通信できません。
会話ごとのコンテキストではなく、セッション間の記憶。 .claude/context/decions.md 、 invariants.md 、および Rejections.md は、すべてのセッション境界を越えて保持されます。パイプラインは、上級エンジニアの制度上の記憶と同じように、決定されたこと、不可侵なもの、試行され拒否されたものを転送します。ほとんどのスペックモード ツールは、会話エッジでそのコンテキストをリセットします。
デモ リポジトリではなく、実際に積極的に開発され、gitflow に統合されたコードベースで証明されています。最初のコミットまで遡り、すべての機能に先立つ 70 以上の PR、仕様、計画がマージされています。これは「plan.md を生成する」とは異なる主張であり、実際の PR 履歴を読んで確認することができます。
これらは組み込み仕様モードを悪くするものではありません。これらは、すでにその IDE 内にあるチームにとっては妥当なデフォルトです。プラグマは、1 つのセッション、IDE、またはエージェントの実行とは関係なく、強制とメモリを存続させたい場合に使用します。
コマンド
何をするのか
/spec "機能のアイデア"
2 ～ 3 つのアプローチを提案し、選択して仕様ドキュメントを保存します
/plan docs/specs/my-spec.md
承認された仕様をタスクごとの実装計画に変える
/機能実行

cs/plans/my-plan.md
承認された計画を実行します — TDD、タスクごとに 1 つのコミット
/ゲート
PR 前にビルド、完全なテスト スイート、アーキテクチャのコンプライアンスを検証します
/レビュー
アーキテクチャのコンプライアンスに関する PR をレビューし、その評決を実際の GitHub レビューとして投稿します
/テスト
機能ブランチのテストを作成します — /review レポートが承認された後に実行します
/pr-フォローアップ
PR を開いた直後に /review と /test を自動チェーンします
/bugfix "説明"
最初に回帰テストを行い、次に修正します — 常にテストが最初です
/リリース1.0.0
バージョンバンプ、変更ログ、メインへの PR、git タグ
ユーティリティ
コマンド
何をするのか
/デザイン
ビジュアル デザイン トークンを確立します — UI 機能の /spec の前に実行します
/並列レビュー
/review のアーキテクチャ チェックリストと code-review:code-review を、 /feature の後、 /gates の前のブランチ diff で並行して実行します。
/パイプライン-レビュー
パイプラインのドリフト、ギャップ、非効率を監査します。
/ステータス
作業が行われている場所を再構築します - セッションを再開するために使用します
/トリムコンテキスト
計画完了後に蓄積されたコンテキストをトリミングする
/同期ワークフロー
この足場をプロジェクトの最新の規則と同期します
スタンドアロンスキル
上記のコマンドとは異なり、これらはパイプラインの残りの部分を採用することなく、どのプロジェクトでも機能します。
3 つの GitHub Actions ワークフローがコマンドとともにプロジェクトにインストールされます。
エージェント層 ( /gates 、 /review 、 /test ) は、PR が開かれる前に高速フィードバックを得るためにローカルで実行されます。その後、CI はバイパスできない強制として同じチェックを独立して再実行します。
フェーズ 2 — TestFlight のアップロードは文書化されていますが、 release.yml でコメント化されています。 Apple Developer Program メンバーシップ、配布証明書、App Store Connect API キーが必要です。準備が完了すると、コメント化されたブロックに追加する内容が正確に表示されます。
パイプラインは、セッション全体にわたって組織の知識を .claude/context/ に蓄積します。
.claude/context/
§

── invariants.md — エージェントが上書きできないルール (アーキテクチャ、通貨タイプなど)
§── Decisions.md — すべての仕様決定のログ: 選択されたアプローチ + 理由
§── feature-log.md — 出荷されたすべての機能の記録
└──拒否.md — 理由を付けて除外されたアプローチ
すべてのエージェントは、行動する前にこれらのファイルを読み取ります。時間の経過とともに、パイプラインには、あらゆるセッション境界を生き延びる上級エンジニアと同じコンテキスト (制約、過去の決定、行き詰まり) が伝えられます。
/feature を初めて実行する前に、invariants.md にデータを入力します。
計画内の各タスクは厳密な TDD に従います。
失敗するテストを作成する — 適切な理由で失敗することを確認する
合格するために最小限のコードを実装する
完全なテスト スイートを実行します - 回帰は許可されません
コミット — タスクごとに 1 つのコミット、バッチ処理なし
テストが赤の場合、エージェントは次のタスクに進むことはありません。
アーキテクチャの前提条件 (デフォルト - CLAUDE.md でオーバーライド):
MVVM + リポジトリ — ビューにはビジネス ロジックが含まれておらず、ViewModel はプロトコルに依存しており、具体的な実装はありません。
永続化のための SwiftData — ドメイン サービスには SwiftData インポートがありません
Swift Testing — 単体テスト/統合テスト用に Testing 、 @Suite 、 @Test 、 #expect() をインポートします。 UIテスト用のXCUITest
PBXFileSystemSynchronizedRootGroup (Xcode 16+) — ファイルは正しいディレクトリに配置されると自動コンパイルされます。 project.pbxproj は決して編集しないでください
./scripts/setup.sh MyApp /path/to/your-project
すべてをコピーし、すべてのプレースホルダーを置き換えます。次に、 CLAUDE.md とシード invariants.md を入力します。
.claude/commands/ 、 .claude/context/ 、 scaffold/.github/workflows/ 、および scripts/ をプロジェクトにコピーします (ワークフローは .github/workflows/ に配置します)。
各コマンド ファイル内の <AppName> をモジュール名に置き換えます。
3 つのワークフロー ファイルの YOUR_PROJECT と YOUR_SCHEME を置き換えます。
CLAを更新する

ビルド コマンド、シミュレーター ターゲット、アーキテクチャ ルールを含む UDE.md
.claude/context/invariants.md に交渉不可能なルールを設定します。
/review のアーキテクチャ ルール チェックリストをスタックに合わせて更新します
main — プロダクション、リリースのみにタグ付け
開発 — 統合ブランチ、すべての機能がここにマージされます
feature/* — 開発外
fix/* — オフ開発 (hotfix/* オフメイン)
release/* — オフ開発、メインに PR、開発にバックマージ
spec/* — 開発中、仕様 + 計画ドキュメント用
著者
エージェント AI パイプラインを構築する iOS リード エンジニアである Akshay Pimprikar によって構築されました。
エージェント駆動の iOS 開発用の Claude Code スラッシュ コマンド - 仕様は 8 段階でリリースされる
www.linkedin.com/in/akshaypimprikar/ トピック
Readme MIT ライセンス
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Claude Code slash commands for agent-driven iOS development - spec to release in 8 stages - akshaypimprikar/pragma

GitHub - akshaypimprikar/pragma: Claude Code slash commands for agent-driven iOS development - spec to release in 8 stages · GitHub
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
akshaypimprikar
/
pragma
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
121 Commits 121 Commits Folders and files
.claude-plugin .claude-plugin .claude .claude .github .github docs/ screenshots docs/ screenshots scaffold/ .github/ workflows scaffold/ .github/ workflows scripts scripts skills/ deterministic-pr-gates skills/ deterministic-pr-gates CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md View all files Repository files navigation
You approve twice. Claude ships the rest.
The complete iOS development scaffold for the agentic era — agent commands, CI enforcement, and setup automation wired together so one engineer ships at team scale.
Proven on FinanceTracker — a production SwiftUI + SwiftData app built entirely on this pipeline from day one, with specs, plans, and PRs going back to the first commit.
Not a mockup — this is what 80+ merged PRs of /spec → /plan → /feature → /gates → /review actually produce. Full screenshot set in FinanceTracker's README .
Quick Start · What You Get · Pipeline · Why Not a Spec Mode? · Commands · CI Layer · Memory Layer · Customising · Contributing
Recommended — install as a Claude Code plugin, no clone or shell script:
Inside Claude Code, in your iOS project's repo root:
/plugin marketplace add akshaypimprikar/pragma
/plugin install pragma@pragma
/pragma:init MyApp
/pragma:init does what scripts/setup.sh does — copies commands, context files, CI workflows, and support scripts, substitutes your app name throughout — but interviews you for CLAUDE.md 's architecture and key-constraints content and seeds .claude/context/invariants.md from the same answers, instead of leaving both as templates to fill in later.
Alternative — clone and run the setup script directly:
git clone https://github.com/akshaypimprikar/pragma
cd pragma
./scripts/setup.sh MyApp /path/to/your-ios-project
This copies the same files and substitutes your app name, but leaves CLAUDE.md and invariants.md as templates — fill them in yourself before running /feature .
Then, either way, kick off your first feature:
/spec "describe your feature idea"
What You Get
Three layers installed into your project:
Each layer is independent — adopt all three or just the commands.
flowchart TD
A([💡 Idea]):::dim --> B
B["/spec\n✓ you approve"]:::human --> C
C["/plan\n✓ you approve"]:::human --> D
D["/feature"]:::auto --> E
E["/gates"]:::auto --> F
F([PR opened]):::dim --> G & H & CI1
CI1["CI · pr-checks\nui-tests"]:::ci --> I
G["/review"]:::auto --> I
H["/test"]:::auto --> I
I([merge to develop]):::dim -.->|next feature| B
I --> J["/release"]:::auto --> K([tag v*.*.*]):::dim --> CI2
CI2["CI · release\nRelease build + GitHub Release"]:::ci
BUG([Bug report]):::dim --> BF["/bugfix"]:::auto --> BG["/gates"]:::auto --> BP([PR]):::dim --> BR["/review"]:::auto --> BM([merge]):::dim
classDef human fill:#3d2800,stroke:#fbbf24,color:#fbbf24
classDef auto fill:#0a1f14,stroke:#34d399,color:#34d399
classDef dim fill:#161b22,stroke:#30363d,color:#8b949e
classDef ci fill:#0d1117,stroke:#58a6ff,color:#58a6ff
Loading
You approve twice — after /spec and after /plan . Every other step is either an agent or automated CI.
Martin Fowler's "Harness Engineering for Coding Agent Users" (April 2026) frames Agent = Model + Harness , splitting the harness into feedforward guides (steer the agent before it acts) and feedback sensors (verify what it did after, across maintainability, architecture fitness, and behavioral correctness). This pipeline maps directly onto that taxonomy:
Fowler calls the behavioral-correctness sensor "the elephant in the room — still unsolved" for most agent harnesses. /test plus /gates ' coverage gate are this pipeline's concrete attempt at that sensor.
Why not just Cursor / Windsurf / Copilot's built-in spec mode?
Every major AI coding tool has shipped some flavor of spec-driven development — Cursor's Plan Mode, Windsurf's Cascade, GitHub Copilot workspace, GitHub Spec Kit, BMAD-METHOD, and others. Fair question: why a separate scaffold instead of using what's already built into the IDE?
Three things pragma does that a spec mode alone doesn't:
CI-enforced, not just agent-enforced. /gates runs locally before a PR opens; the same checks re-run independently in GitHub Actions ( pr-checks.yml , ui-tests.yml ) as enforcement that can't be skipped by rerunning the agent with a different prompt. Spec modes generate a plan; they don't wire in an enforcement layer the agent itself can't talk its way around.
Cross-session memory, not per-conversation context. .claude/context/decisions.md , invariants.md , and rejections.md persist across every session boundary — the pipeline carries forward what was decided, what's inviolable, and what's been tried and rejected, the way a senior engineer's institutional memory would. Most spec-mode tools reset that context at the conversation edge.
Proven on a real, actively-developed, gitflow-integrated codebase , not a demo repo — 70+ merged PRs, specs and plans predating every feature, going back to the first commit. That's a different claim than "generates a plan.md," and it's checkable: read the actual PR history.
None of this makes the built-in spec modes bad — they're a reasonable default for teams already inside that IDE. Pragma is for when you want the enforcement and the memory to survive independently of any one session, IDE, or agent run.
Command
What it does
/spec "feature idea"
Proposes 2–3 approaches, you choose, spec doc saved
/plan docs/specs/my-spec.md
Turns an approved spec into a task-by-task implementation plan
/feature docs/plans/my-plan.md
Executes an approved plan — TDD, one commit per task
/gates
Verifies build, full test suite, and architecture compliance before PR
/review
Reviews a PR for architecture compliance, posts its verdict as a real GitHub review
/test
Writes tests for a feature branch — runs after /review reports APPROVED
/pr-followup
Auto-chains /review then /test right after a PR opens
/bugfix "description"
Regression test first, then fix — test-first always
/release 1.0.0
Version bump, changelog, PR to main, git tag
Utility
Command
What it does
/design
Establishes visual design tokens — run before /spec on UI features
/parallel-review
Runs /review 's architecture checklist and code-review:code-review in parallel on the branch diff, after /feature , before /gates
/pipeline-review
Audits the pipeline for drift, gaps, and inefficiencies
/status
Reconstructs where work stands — use to resume any session
/trim-context
Trims accumulated context after completing a plan
/sync-workflow
Syncs this scaffold with your project's latest conventions
Standalone skills
Unlike the commands above, these work in any project without adopting the rest of the pipeline.
Three GitHub Actions workflows install into your project alongside the commands:
The agent layer ( /gates , /review , /test ) runs locally for fast feedback before a PR is opened. CI then re-runs the same checks independently as enforcement that can't be bypassed.
Phase 2 — TestFlight upload is documented but commented out in release.yml . It requires an Apple Developer Program membership, distribution certificate, and App Store Connect API key. When you're ready, the commented block shows exactly what to add.
The pipeline accumulates institutional knowledge across sessions in .claude/context/ :
.claude/context/
├── invariants.md — rules no agent may override (architecture, money types, etc.)
├── decisions.md — log of every spec decision: approach chosen + reason
├── feature-log.md — record of every feature shipped
└── rejections.md — approaches ruled out, with reasons
Every agent reads these files before acting. Over time the pipeline carries the same context a senior engineer would — constraints, past decisions, and dead ends — surviving every session boundary.
Populate invariants.md before running /feature for the first time.
Each task in the plan follows strict TDD:
Write the failing test — confirm it fails for the right reason
Implement the minimal code to make it pass
Run the full test suite — no regressions allowed
Commit — one commit per task, no batching
The agent never proceeds to the next task if tests are red.
Architecture assumptions (defaults — override in CLAUDE.md ):
MVVM + Repository — views contain no business logic, ViewModels depend on protocols never concrete implementations
SwiftData for persistence — Domain Services have zero SwiftData imports
Swift Testing — import Testing , @Suite , @Test , #expect() for unit/integration tests; XCUITest for UI tests
PBXFileSystemSynchronizedRootGroup (Xcode 16+) — files auto-compile when placed in the correct directory; never edit project.pbxproj
./scripts/setup.sh MyApp /path/to/your-project
Copies everything and substitutes all placeholders. Then fill in CLAUDE.md and seed invariants.md .
Copy .claude/commands/ , .claude/context/ , scaffold/.github/workflows/ , and scripts/ into your project (place the workflows at .github/workflows/ )
Replace <AppName> with your module name in each command file
Replace YOUR_PROJECT and YOUR_SCHEME in the three workflow files
Update CLAUDE.md with your build commands, simulator target, and architecture rules
Populate .claude/context/invariants.md with your non-negotiable rules
Update the Architecture Rules checklist in /review to match your stack
main — production, tagged on release only
develop — integration branch, all features merge here
feature/* — off develop
fix/* — off develop (hotfix/* off main)
release/* — off develop, PR to main, back-merged to develop
spec/* — off develop, for spec + plan docs
Author
Built by Akshay Pimprikar — iOS lead engineer building agentic AI pipelines.
Claude Code slash commands for agent-driven iOS development - spec to release in 8 stages
www.linkedin.com/in/akshaypimprikar/ Topics
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
