---
source: "https://github.com/fu351/Doberman-Core"
hn_url: "https://news.ycombinator.com/item?id=49336757"
title: "Show HN: Doberman: The AI watchdog that stops Claude from deleting your database"
article_title: "GitHub - fu351/Doberman-Core: Your AI's guard dog. Doberman sits at runtime, gating every input, output and tool call to stop unsafe or unintended actions before they execute. · GitHub"
image: "https://opengraph.githubassets.com/e04c11fd215725549c6feb852e1c81bcce720a7895f841a22724d49ef45fbe14/fu351/Doberman-Core"
author: "alanfuNZ"
captured_at: "2026-08-17T20:15:33Z"
capture_tool: "hn-digest"
hn_id: 49336757
score: 2
comments: 0
posted_at: "2026-08-17T20:03:00Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Doberman: The AI watchdog that stops Claude from deleting your database

- HN: [49336757](https://news.ycombinator.com/item?id=49336757)
- Source: [github.com](https://github.com/fu351/Doberman-Core)
- Score: 2
- Comments: 0
- Posted: 2026-08-17T20:03:00Z

## Translation

タイトル: Show HN: Doberman: クロードによるデータベースの削除を阻止する AI ウォッチドッグ
記事タイトル: GitHub - fu351/Doberman-Core: AI の番犬。 Doberman は実行時に常駐し、すべての入力、出力、およびツール呼び出しをゲートして、安全でないアクションや意図しないアクションを実行前に停止します。 · GitHub
説明: AI の番犬。 Doberman は実行時に常駐し、すべての入力、出力、およびツール呼び出しをゲートして、安全でないアクションや意図しないアクションを実行前に停止します。 - fu351/ドーベルマンコア
HN テキスト: ランディング ページ: https://www.trydoberman.dev/ 問題:
AI ガードレールは、危険なコマンドの 89% のみを保護します。これは、9 人に 1 人の侵入者を侵入させる錠前であると考えるまでは、非常に優れているように思えるかもしれません。 ドーベルマンとは:
Doberman は、LLM にラップされた 2 層のセキュリティ システムで、すべての入力、出力、およびツール呼び出しを監視します。最初の層は最先端のセキュリティ ガイドラインに基づいて決定的であり、2 番目の層は動的で、使用状況や個人の好みから学習します。ドーベルマンとの違いは次のとおりです。
これは動的な 2 層のセキュリティ システムであり、実行時に常駐してすべての入力、出力、ツールの実行を監視し、作業に合わせて適応します。

記事本文:
GitHub - fu351/Doberman-Core: AI の番犬。 Doberman は実行時に常駐し、すべての入力、出力、およびツール呼び出しをゲートして、安全でないアクションや意図しないアクションを実行前に停止します。 · GitHub
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
ふ351
/
ドーベルマンコア
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
397 コミット 397 コミット .github .github アダプター アダプター ドキュメント ドキュメントの例/プラグイン ガードレールの例/プラグイン ガードレール scri

pts スクリプト src/ doberman src/ doberman テスト テスト ツール ツール .gitignore .gitignore .gitleaks.toml .gitleaks.toml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md CONTRIBUTORS.md CONTRIBUTORS.md ライセンス ライセンス README.md README.md RELEASING.md RELEASING.md SECURITY.md SECURITY.md logo.png logo.png pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング エージェント向けの適応型認可とランタイム ガードレール
AI コーディング エージェントは、元に戻すことなく自律的に、リポジトリの rm -rf を実行したり、API キーを漏洩したり、抽出データにプロンプト挿入されたりすることができます。ドーベルマンは、危険な呼び出しを実行前に阻止する、実行経路上の番犬です。
ライブ ダッシュボードに対するドーベルマンのデモ (ドーベルマン ダッシュ): 5 つの攻撃が発生するたびにブロックされ、その後、人間が高リスクの承認を拒否しました。
実行パス上にない場合、それは保護ではなく勧告です。
Doberman はエージェントとそのツール (透過的な MCP プロキシまたはホスト フック) の間に位置し、あらゆるアクションを明示的で監査可能な決定に変換します。すべてのツール呼び出しは、実行前に決定される 1 つの判定を取得します。
AIエージェント ──▶ Doberman ──▶ リアルツール（ファイル、シェル、MCPサーバー、API）
└─ 正規化 → リスクエンジン → PASS / AUTH / BLOCK
Claude Code、Codex、OpenClaw、および任意の MCP 互換エージェントで動作します。カーソルおよび他の MCP クライアントは MCP プロキシを介して接続します。オープンソース、ローカルファースト、そして決して破られることのない 2 つのルールに縛られています。フェイルクローズ (不確実性の否定) とレイズオンリー (自動的に締めることはできますが、黙って緩めることはありません) です。
2 つのコマンドで保護される · Discord でパックに参加する
なぜドーベルマンなのか？ — その機能と 2 つの保証
クイック スタート — インストールと保護

2 つのコマンドで gent
エンドツーエンドで検証します - 実際の MCP サーバーの前で監視します
ターン ゲート — オプションの推論前のチョークポイント
ベンチマーク — 攻撃ブロック率 (ASR) と誤検知フリクション (FPR)
カスタム ガードレールを作成する — 独自のルールをプラグインとして登録します
リスク許容度に合わせて調整 — 厳格モード + 強制ダイヤル
プロンプト インジェクション、ツール ポイズニング、データ漏洩、エージェントの暴走は、エージェント AI のセキュリティ上の問題を定義するものです。ほとんどの「AI ガードレール」はプロンプトを検査し、アドバイスを提供します。 Doberman は異なります。Doberman はツールの実行パス上にあるため、ブロックされたアクションが実行されることはありません。
これら 3 つの判定は、モデルが過去に話してもよいというアドバイスではありません。それぞれの判定は、呼び出しが実行される前の瞬間という、重要な 1 つの場所で適用されます。モデル自身のガードレールを好きなだけ回避できます。アクションはまだドーベルマンをクリアする必要があります。
2 つの交渉不可能なプロパティ:
フェールクローズ — エラー、不確実性、または未処理のケースがある場合、アクションは拒否されます。意思決定エンジンに関するツールへの道はありません。沈黙を含む: 誰も応答しない承認プロンプトは、厳密な期限 (デスクトップ ダイアログの場合は 2 分、チャレンジ全体のバックストップとして 20 分) によって制限され、拒否として解決され、拒否ではなくタイムアウトとして明確にログに記録されます。無期限のブロックは拒否ではなく、無人というのがまさにエージェントの通常の実行方法です。
引き上げのみの学習 — ガードレールと適応学習は自動的に締めることができ、黙って緩むことはありません。永続的なポリシーを弱体化するには、明示的な、ポゼッションファクターゲートによる監査済みの人間の承認 (登録されている場合は TOTP、それ以外の場合はローカルの Doberman パスワード) が必要です。
どの保証がどのホストに適用されますか?パリティ マトリックスは、各保護を各ホストの Doberman フロント (Claude Code、Codex、MCP プロキシ、OpenClaw) にマッピングします。すべての ✅ は、それを証明する CI テストにリンクしています。 ◻c

セルはオープンな、寄稿者サイズの作業 ( parity というラベルが付いています) であり、マトリックスはすべてのビルドのテストから再生成されるため、実際に証明されたものから逸脱することはできません。
Doberman は、MCP 互換のコーディング エージェントを保護します。エージェントを選択し、1 つのコマンドを実行し、すべてのコマンドを実行します。
ツール呼び出しは実行前にレビューされます。完全なウォークスルー - すべてのオプション、フラグ、
ダッシュボード、ヘルス チェック - セットアップ ガイドに記載されています。
pip インストール doberman-core
インストール後、 doberman --install-completion を実行してシェル タブを有効にします。
完成。
アンインストールしますか? pip uninstall doberman-core の前に doberman uninstall-hook を実行します -
そうしないと、settings.json に残っているフック エントリが、なくなったバイナリを呼び出し続けることになります。
そして、すべてのツール呼び出しは doberman: command not found で失敗します。すでにヒットしていますか？ただ
pip install doberman-core を再度インストールします - 既存のエントリは依然として正しく、機能し始めます
すぐに修理の必要はありません。
doberman セットアップ # 厳密性モードを選択し、ガードレールを調整し、フックを配線します
Doberman は、エージェントが行うすべてのツール呼び出しを確認するようになりました。ドーベルマン博士に確認するか、時計を見てください
ドーベルマンデモによる実際の判決。その他すべて - MCP プロキシ配線、ダッシュボード、TUI、スキャン、
および 2FA - はセットアップ ガイドに記載されています。
判定は表示されるすべての場所で色分けされています (BLOCK 赤、AUTH オレンジ、PASS 緑) —
review 、 status 、 log 、 TUI 、および demo 。説明も端末にラップされます。
幅。出力がパイプまたはリダイレクトされると、色は自動的に削除され、優先されます。
NO_COLOR - その変数が空ではない値に設定されている場合。
CLI 診断では 1 つの重大度語彙が使用されます。致命的な失敗はエラーで始まり、成功しましたが、
劣化またはスキップされた作業は警告: で始まり、情報提供は注記: で始まります。
リポジトリ MCP conf の静的読み取り専用アドミッション チェックを行うための doberman スキャンに --mcp を追加します。

イグス;
パターン クラスのみをレポートし、生の URL、引数、環境値などを出力することはありません。
設定内容。
実際の MCP サーバーの前で Doberman を監視する 2 つの方法 — プロセス内テストがチェーンのどこにも重複しないこと。
インタラクティブなデモ — MCP Inspector + 実際のファイルシステム サーバー:
npx -y @modelcontextprotocol/inspector dobermanserve -- npx -y @modelcontextprotocol/server-filesystem ~ /my-project
インスペクター UI を開き、Doberman 経由でツールを呼び出します。ルーチンは実際のファイルシステム サーバーに直接 PASS を読み書きします。破壊的な呼び出しはポリシー エラーとして返され、決して実行されません。
エンドツーエンドのテスト - 開発チェックアウトで:
pytest テスト/統合/test_serve_end_to_end.py -q
これにより、doberman が実際の stdio ツール サーバー ( testing/fixtures/stdio_tool_server.py ) のフロントとなる実際のサブプロセスとして機能し、エージェントとして動作する実際の MCP クライアントに接続し、実際の stdio 上でデプロイ可能なチェーンをアサートします。
ダウンストリームのツールはプロキシを通じて再公開されます。
PASS 判定がツールに到達します (ダウンストリームの通話ログに記録されます)。
BLOCK 判定 ( rm -rf / ) は決して到達しません。通話ログは空のままです。
最後のアサーションは、プロジェクト全体が依存するチョークポイント プロパティです。
テスト フィクスチャに関する注意: 統合スイートの残りの部分は、実行されるすべての呼び出しを記録するインプロセスの偽のダウンストリーム (tests/fixtures/fake_tool_server.py) を意図的に使用します。記録は、ブロックされたアクションが何も到達しなかったことをテストが証明する方法です。これはテスト フィクスチャであり、ランタイムではありません。 doberman サーブは常に生成され、 -- の後に与えられた実サーバーと通信します。
Doberman のプロキシは、 pyproject.toml に固定されているように MCP を話します ( mcp>=1.27,<2 )。 MCP
2026-07-28 リビジョン (RC) では、プロトコル レベルのセッション ID (SEP-2575) が削除されました。ドーベルマンの
クロスコール保護 (テイント台帳、リ

広告対送信フィンガープリント、決定ログ）キーオフ
リポジトリローカル ID であり、プロトコルセッションではなく、回帰テスト済みのステートレスです
(テスト/統合/test_stateless_identity.py)。 「2026-07-28 をサポート」が要求されます
仕様が完成し、固定された SDK がそれを採用した場合のみです。
同じ意思決定エンジンの 2 番目の呼び出しポイント - ユーザーのターン (プロンプト + 添付/貼り付け/ツールフェッチされたコンテンツ) のホストの事前推論フックで参照されるため、単一の推論トークンが消費される前に、目に余るターンが判断されます。ターン ゲートは、意図的に狭い保証を備えた効率/早期警告レイヤーであり、Tier-0 シグネチャ ターンはモデルに到達しません。上記のアクション ゲートは安全性を保証するものであり、ターン ゲートを回避した攻撃者も依然としてそれに遭遇します。
1 つのエンジン、2 つの呼び出しポイント。 Decide_turn は、 raise-only combo 、 Decision Audit モデル、および tiered-auth Challenge (ゼロの新しい判定権限) を再利用します。アクションパスを変更する必要はありません。 turngate/ は新しいアダプター (プロキシの兄弟) であり、エンジンは純粋なままです (階層ガードレールが挿入されます)。ターンの評決は、 action_type='turn' とマークされた同じ編集済みの決定ログに追加されます。
層 0 — 決定論的シグネチャ (唯一のハードストップであり、設計上小規模です)。 struction_nullification 、authority_override (偽装 / モード切り替え / "システム プロンプトを印刷")、 secret_export 、および encoded_pa​​yload 。精度の核となるのは、問題対メンションの識別子 + 元のルールです。信頼できない (貼り付けられた/ツールでフェッチされた) コンテンツ ブロックの無条件の一致 (間接注入)。型付けされ発行された一致ブロック。型指定された + 言及された (引用/メタディスカッション) 一致は AUTH までステップアップするため、攻撃について議論している研究者がハードブロックされることはありません。曖昧→発行;コードフェンスは決して免除されません。
階層 1 — ヒューリスティック リコール (AUTH)

-のみ、構造的に BLOCK 不可能)。貼り付けられたテキスト内に埋め込まれたエージェント主導の指示 (単なる命令的な雰囲気ではないため、チュートリアルでつまずくことはありません)、ペルソナのオーバーライド、しきい値以下の難読化、緊急性と機密性のフレーム化 — 誤検知のコストは 1 タップであり、プロンプトが拒否されることはありません。
スタイロメトリック共起ゲート。エンティティごとのプロンプト スタイルのベースライン (粗いスカラー バケットのみ - 長さ、単語の形状、句読点/大文字/小文字の密度/数字の密度。テキストは一切含まれません) は、各ターンのスタイルを経験的 CDF p 値としてスコア付けします。これは、主観的なベースラインがアクションに使用するのと同じ調整アイデアです。ゲートが強化されるのは、極端なスタイルの外れ値が、繊細な明白な意図 (資格情報 / 破壊的 / 外部送信) と同時発生した場合のみです。決してスタイルだけではありません (入力方法は人によって異なります)。疲れている/モバイル/貼り付け、デバイスまたは言語の切り替えは攻撃ではなくドリフトです)。 Style-weird のみがタグ付けされており、ゲートされておらず、プロンプトベースラインが成熟するまで不活性のままです (他のベースラインと同じ成熟度ルール)。 Tier 0 はターン 1 からアクティブになります。既知の制限: 共有アカウントはタイピストをブレンドし、スタイル信号を劣化させてノイズを発生させます。共起により追加タップ 1 回のコストが制限されます。
タグアンドパス: ターンシグナルがアクションステージに送られます (レイズのみ)。リリースされたすべてのターンは、そのスタイロメトリック p 値とヒューリスティック フラグをエンティティごとの TurnC として公開します。

[切り捨てられた]

## Original Extract

Your AI's guard dog. Doberman sits at runtime, gating every input, output and tool call to stop unsafe or unintended actions before they execute. - fu351/Doberman-Core

Landing Page: https://www.trydoberman.dev/ The Problem:
AI guardrails only protect against 89% of dangerous commands. This may seem pretty good until you think of it as a lock that lets 1 in 9 intruders in. What is Doberman:
Doberman is a two layer security system wrapped around your LLM monitoring every input, output and tool call. The first layer is deterministic based on state of the art security guidelines, the second layer is dynamic, learning from your use and personal preferences. How is Doberman different:
It is a dynamic two-layer security system that sits at runtime and monitors every input, output, and tool execution, adapting with you as you work.

GitHub - fu351/Doberman-Core: Your AI's guard dog. Doberman sits at runtime, gating every input, output and tool call to stop unsafe or unintended actions before they execute. · GitHub
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
fu351
/
Doberman-Core
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
397 Commits 397 Commits .github .github adapters adapters docs docs examples/ plugin-guardrail examples/ plugin-guardrail scripts scripts src/ doberman src/ doberman tests tests tools tools .gitignore .gitignore .gitleaks.toml .gitleaks.toml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md CONTRIBUTORS.md CONTRIBUTORS.md LICENSE LICENSE README.md README.md RELEASING.md RELEASING.md SECURITY.md SECURITY.md logo.png logo.png pyproject.toml pyproject.toml View all files Repository files navigation
Adaptive Authorization & Runtime Guardrails for AI Coding Agents
Your AI coding agent can rm -rf your repo, leak your API keys, or be prompt-injected into exfiltrating data — autonomously, with no undo. Doberman is the guard dog on the execution path that stops the dangerous call before it runs.
doberman demo against the live dashboard ( doberman dash ): five attacks blocked as they happen, then a human denies a high-risk approval.
If it isn't on the execution path, it's advisory, not protective.
Doberman sits between the agent and its tools — a transparent MCP proxy or host hook — and turns every action into an explicit, auditable decision. Every tool call gets exactly one verdict, decided before it executes:
AI agent ──▶ Doberman ──▶ real tools (files, shell, MCP servers, APIs)
└─ normalize → risk engine → PASS / AUTH / BLOCK
Works with Claude Code, Codex, OpenClaw, and any MCP-compatible agent — Cursor and other MCP clients connect through the MCP proxy . Open-source, local-first, and bound by two rules it will never break: it fails closed (uncertainty denies) and is raise-only (it can tighten automatically, but never silently loosens).
Get protected in two commands · Join the pack on Discord
Why Doberman? — what it does, and the two guarantees
Quick Start — install and protect an agent in two commands
Verify it end-to-end — watch it front a real MCP server
Turn gate — the optional pre-inference chokepoint
Benchmark — attack-block rate (ASR) vs. false-positive friction (FPR)
Write a custom Guardrail — register your own rule as a plugin
Tune to your risk tolerance — strictness modes + the enforcement dial
Prompt injection, tool poisoning, data exfiltration, and runaway agents are the defining security problems of agentic AI. Most "AI guardrails" inspect prompts and offer advice. Doberman is different: it is on the tool-execution path , so a blocked action never runs .
Those three verdicts aren't advice a model can talk its way past — each is enforced at the one place that counts: the instant before the call runs. Evade the model's own guardrails all you want; the action still has to clear Doberman.
Two non-negotiable properties:
Fail closed — any error, uncertainty, or unhandled case denies the action. There is no path to a tool around the decision engine. Including silence: an approval prompt nobody answers is bounded by a hard deadline (2 minutes for the desktop dialog, 20 minutes as the whole-challenge backstop) and resolves to a denial , logged distinctly as timeout rather than denied — an indefinite block is not a denial, and unattended is exactly how agents usually run.
Raise-only learning — guardrails and adaptive learning can auto- tighten , never silently loosen. Every permanent policy weakening requires explicit, possession-factor-gated, audited human approval (TOTP if enrolled, otherwise the local Doberman password).
Which guarantee holds on which host? The parity matrix maps each protection to each host Doberman fronts (Claude Code · Codex · MCP proxy · OpenClaw). Every ✅ links to the CI test that proves it; ◻ cells are open, contributor-sized work (labeled parity ), and the matrix is regenerated from the tests on every build, so it can't drift from what's actually proven.
Doberman guards any MCP-compatible coding agent: pick your agent, run one command, and every
tool call is reviewed before it executes. The full walkthrough - every option, flag, the
dashboard, and health checks - lives in the Setup Guide .
pip install doberman-core
After installing, run doberman --install-completion to enable shell tab
completion.
Uninstalling? Run doberman uninstall-hooks before pip uninstall doberman-core -
otherwise the hook entries left in settings.json will keep invoking a binary that's gone,
and every tool call fails with doberman: command not found . Already hit this? Just
pip install doberman-core again - the existing entries are still correct and start working
immediately, no repair needed.
doberman setup # pick a strictness mode, tune guardrails, wire the hooks
Doberman now reviews every tool call your agent makes. Confirm it with doberman doctor , or watch
real verdicts with doberman demo . Everything else - MCP-proxy wiring, the dashboard, TUI, scan,
and 2FA - is in the Setup Guide .
Verdicts are colour-coded everywhere they're shown ( BLOCK red, AUTH amber, PASS green) —
review , status , log , the TUI, and demo , which also wraps its explanations to your terminal
width. Colour is dropped automatically when output is piped or redirected, and honours
NO_COLOR when that variable is set to a non-empty value.
CLI diagnostics use one severity vocabulary: fatal failures start with error: , successful but
degraded or skipped work starts with warning: , and informational asides start with note: .
Add --mcp to doberman scan for a static, read-only admission check of repository MCP configs;
it reports pattern classes only and never emits raw URLs, arguments, environment values, or other
config content.
Two ways to watch Doberman front a real MCP server — no in-process test doubles anywhere in the chain.
Interactive demo — MCP Inspector + a real filesystem server:
npx -y @modelcontextprotocol/inspector doberman serve -- npx -y @modelcontextprotocol/server-filesystem ~ /my-project
Open the Inspector UI and call tools through Doberman: routine reads and writes PASS straight through to the real filesystem server; a destructive call comes back as a policy error and never executes.
End-to-end test — in a dev checkout:
pytest tests/integration/test_serve_end_to_end.py -q
This spawns doberman serve as a real subprocess fronting a real stdio tool server ( tests/fixtures/stdio_tool_server.py ), connects to it with a real MCP client playing the agent, and asserts the deployable chain over actual stdio:
the downstream's tools are re-exposed through the proxy,
a PASS verdict reaches the tool (the downstream's call log records it), and
a BLOCK verdict ( rm -rf / ) never reaches it — the call log stays empty.
That last assertion is the chokepoint property the whole project hangs on.
Note on the test fixtures: the rest of the integration suite deliberately uses an in-process fake downstream ( tests/fixtures/fake_tool_server.py ) that records every call it executes — recording is how the tests prove a blocked action reached nothing . It is a test fixture, not the runtime. doberman serve always spawns and talks to the real server you give it after -- .
Doberman's proxy speaks MCP as pinned in pyproject.toml ( mcp>=1.27,<2 ). The MCP
2026-07-28 revision (RC) removes protocol-level session identity (SEP-2575); Doberman's
cross-call protections (taint ledger, read-vs-send fingerprints, decision log) key off
repo-local identity, never the protocol session, and are regression-tested stateless
( tests/integration/test_stateless_identity.py ). "Supports 2026-07-28" will be claimed
only once the spec finalizes and the pinned SDK adopts it.
A second invocation point for the same decision engine — consulted at a host pre-inference hook on the user's turn (prompt + attached/pasted/tool-fetched content), so a flagrant turn is judged before a single inference token is spent . The turn gate is an efficiency / early-warning layer with a deliberately narrow guarantee — no Tier-0-signature turn reaches the model . The action gate above remains the safety guarantee : an attacker who evades the turn gate still meets it.
One engine, two invocation points. decide_turn reuses the raise-only combine , the Decision audit model, and the tiered-auth challenge — zero new verdict authority . It needs no change to the action path; turngate/ is a new adapter (sibling to the proxy), and the engine stays pure (the tier guardrails are injected). Turn verdicts append to the same redacted decisions log, marked action_type='turn' .
Tier 0 — deterministic signatures (the only hard-stop, small by design). instruction_nullification , authority_override (impersonation / mode-switch / "print your system prompt"), secret_export , and encoded_payload . The precision core is the issue-vs-mention discriminator + origin rule : a match in untrusted (pasted / tool-fetched) content blocks unconditionally (indirect injection); a typed + issued match blocks; a typed + mentioned (quoted / meta-discussion) match steps up to AUTH so a researcher discussing an attack is never hard-blocked; ambiguous → issued; code fences never exempt.
Tier 1 — heuristic recall (AUTH-only, structurally BLOCK-incapable). Embedded agent-directed instructions inside pasted text (not mere imperative mood, so tutorials don't trip it), persona override, sub-threshold obfuscation, and urgency+secrecy framing — a false positive costs one tap, not a denied prompt.
Stylometric co-occurrence gate. A per-entity prompt-style baseline (coarse scalar buckets only — length, word shape, punctuation/case/digit density; never any text) scores each turn's style as an empirical-CDF p-value, the same calibration idea the subjective baseline uses for actions. The gate steps up only when an extreme style outlier co-occurs with a sensitive apparent intent (credential / destructive / external-send) — never on style alone (people type differently tired/mobile/pasting; a device or language switch is drift, not an attack). Style-weird alone is tagged, not gated , and stays inert until the prompt baseline matures (the same maturity rule as the other baselines); Tier 0 is active from turn one. Known limitation: shared accounts blend typists, degrading the style signal toward noise — co-occurrence bounds the cost at one extra tap.
Tag-and-pass: turn signals feed the action stage (raise-only). Every released turn publishes its stylometric p-value and heuristic flags as a per-entity TurnC

[truncated]
