---
source: "https://github.com/borhen68/SkillEngine"
hn_url: "https://news.ycombinator.com/item?id=48497893"
title: "AgentForge–28 production-grade skills that make AI agents ship reliable code"
article_title: "GitHub - borhen68/SkillEngine: Production-grade engineering skills for AI coding agents. · GitHub"
author: "borhensaidi"
captured_at: "2026-06-12T01:02:21Z"
capture_tool: "hn-digest"
hn_id: 48497893
score: 1
comments: 0
posted_at: "2026-06-11T23:35:50Z"
tags:
  - hacker-news
  - translated
---

# AgentForge–28 production-grade skills that make AI agents ship reliable code

- HN: [48497893](https://news.ycombinator.com/item?id=48497893)
- Source: [github.com](https://github.com/borhen68/SkillEngine)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T23:35:50Z

## Translation

タイトル: AgentForge – AI エージェントが信頼性の高いコードを出荷できるようにする 28 の実稼働グレードのスキル
記事のタイトル: GitHub - borhen68/SkillEngine: AI コーディング エージェントのためのプロダクション グレードのエンジニアリング スキル。 · GitHub
説明: AI コーディング エージェント向けの実稼働グレードのエンジニアリング スキル。 - borhen68/スキルエンジン

記事本文:
GitHub - borhen68/SkillEngine: AI コーディング エージェント向けの実稼働グレードのエンジニアリング スキル。 · GitHub
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
ボーヘン68
/
スキルエンジン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブラン

ches タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット .claude-plugin .claude-plugin .claude/ コマンド .claude/ コマンド .gemini/ コマンド .gemini/ コマンド .github .github .opencode .opencode エージェント エージェント アセット アセット コマンド コマンド docs docs フック フック リファレンス リファレンス スクリプト スクリプト スキル スキル .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .markdownlint-cli2.yaml .markdownlint-cli2.yaml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md COTRIBUTING.md ライセンス ライセンス Makefile Makefile README.md README.md package.json package.json plugin.json plugin.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
本番グレードの AI エージェントを作成します。
実戦で実証された28のスキル。 6人のスペシャリストペルソナ。 5 つの参考チェックリスト。ミッションの 1 つは、AI エージェントに上級エンジニアのようにソフトウェアを構築させることです。
AI コーディング エージェントは高速です。彼らも無謀だ。
彼らは仕様をスキップします。彼らはテストを「忘れる」のです。レビューなしで発送します。彼らは「私のマシンで動作する」ことを成功基準として扱います。つまり、彼らは製品版ソフトウェアではなく、プロトタイプを構築します。
私たちはエージェントに曖昧な提案をしません。私たちは、上級エンジニアが実際にソフトウェアを構築する方法をコード化した、構造化され、実戦テストされたワークフローを彼らに提供します。これは、Google、Netflix、Stripe のチームを強化するのと同じワークフローです。すべてのスキルには、ステップ、チェックポイント、合理化に対する防御策、および証拠に基づく検証があります。エージェントがこれらに従うと、信頼できるコードが送信されます。
計画の定義 建造の検証 船舶運用のレビュー
┌──────┐ ┌──────┐ ┌──────┐ ┌──────┐ ┌──────┐ ┌──────┐ ┌──────┐
│ アイデア │ ──▶ │ 仕様 │ ──▶ │ コード │ ──▶ │ テスト

│ ──▶ │ QA │ ────▶ │ 行く │ ────▶ │ 走る │
│リファイン│ │ PRD │ │ インプル │ │デバッグ │ │ ゲート │ │ ライブ │ │ 運用 │
━━━━┘ ━━━━┘ ━━━━┘ ━━━━┘ ━━━━┘ ━━━━┘ ━━━━┘
/spec /plan /build /test /review /ship
7つのスラッシュコマンド。スキルは28個。 6段階。言い訳ゼロ。
その他の即時パック
エージェントフォージ
構造
漠然としたアドバイス
チェックポイントを備えた段階的なワークフロー
検証
「うまくいくか確認してください」
証拠に基づいた終了基準 (テスト、ビルド、実行時データ)
不正行為防止
なし
エージェントがステップをスキップするときに使用する言い訳を説明する合理化テーブル
範囲
一般的なコーディングのヒント
全ライフサイクル: 仕様 → 計画 → 構築 → 検証 → レビュー → 出荷 → 運用
高品質のゲート
なし
8 つの自動チェックを備えた組み込みの CI パイプライン
相互参照
サイロ
すべてのスキルは関連するスキルを参照します。重複はありません
コマンド
開発ライフサイクルに対応する 7 つのスラッシュ コマンド。それぞれが適切なスキルを自動的にアクティブにします。
仕様が存在したら、手動の手順を減らしたいですか? /build auto は計画を生成し、単一の承認されたパスですべてのタスクを実装します。計画を一度承認すると、計画は自律的に実行されます。これにより、検証ではなくタスク間の人間の介入が不要になります。すべてのタスクは引き続きテスト駆動され、個別にコミットされ、失敗または危険なステップで一時停止されます。
スキルは、ユーザーの実行内容に基づいて自動的にアクティブ化されます。API の設計では api-and-interface-design がトリガーされ、UI の構築ではfrontend-ui-engineering がトリガーされます。
クロード・コード (推奨)
マーケットプレイスのインストール:
/プラグインマーケットプレイス追加borhen68/SkillEngine
/プラグインインストールagentforge@borhen-agentforge
SSHエラー?マーケットプレイスは SSH 経由でリポジトリをクローンします。そうしないと

GitHub に SSH キーが設定されていない場合は、SSH キーを追加するか、完全な HTTPS URL を使用して HTTPS クローン作成を強制します。
/plugin マーケットプレイスに https://github.com/borhen68/SkillEngine.git を追加
/プラグインインストールagentforge@borhen-agentforge
git clone https://github.com/borhen68/SkillEngine.git
クロード --plugin-dir /path/to/agentforge
カーソル
SKILL.md を .cursor/rules/ にコピーするか、完全な skill/ ディレクトリを参照します。 docs/cursor-setup.md を参照してください。
スキル、サブエージェント、およびスラッシュ コマンドのネイティブ プラグインとしてインストールします。 docs/antigravity-setup.md を参照してください。
agy プラグインのインストール https://github.com/borhen68/SkillEngine.git
ローカル クローンからインストールします。
git clone https://github.com/borhen68/SkillEngine.git
agy プラグインのインストール ./agentforge
ジェミニ CLI
自動検出のためにネイティブ スキルとしてインストールするか、永続的なコンテキストのために GEMINI.md に追加します。 docs/gemini-cli-setup.md を参照してください。
gemini スキルのインストール https://github.com/borhen68/SkillEngine.git --path スキル
ローカル クローンからインストールします。
gemini スキル インストール ./agentforge/skills/
ウィンドサーフィン
スキルのコンテンツをウィンドサーフィンのルール設定に追加します。 docs/windsurf-setup.md を参照してください。
AGENTS.md およびスキル ツールを介したエージェント駆動のスキル実行を使用します。
.github/copilot-instructions.md 内の Copilot ペルソナおよびスキル コンテンツとして、 Agents/ のエージェント定義を使用します。 docs/copilot-setup.md を参照してください。
スキルは単純な Markdown です。システム プロンプトまたは指示ファイルを受け入れる任意のエージェントで動作します。 docs/getting-started.md を参照してください。
上記のコマンドはエントリ ポイントです。このパックには、合計 28 のスキル (24 のライフサイクル スキル、4 つの操作スキル、および using-agentforge メタスキル) が含まれています。各スキルは、ステップ、検証ゲート、および反合理化テーブルを含む構造化されたワークフローです。スキルを直接参照することもできます。
メタ - どのスキルが適用されるかを確認します
スキル
何をするのか
いつ使用するか
AG の使用

アントフォージ
入ってくる作業を適切なスキルのワークフローにマッピングし、共有の運用ルールを定義します
セッションの開始または適用するスキルの決定
定義 - 何を構築するかを明確にする
スキル
何をするのか
いつ使用するか
インタビューしてください
一度に 1 つの質問を行うインタビューで、ユーザーが望むべきだと考えているものではなく、実際に望んでいることを最大 95% の信頼度まで抽出します。
質問の指定が不十分であるか、ユーザーが「インタビューしてください」/「グリルしてください」を呼び出しています
アイデアを洗練する
漠然としたアイデアを具体的な提案に変えるための、構造化された発散/収束思考
検討が必要な大まかなコンセプトがある
スペック駆動開発
コードの前に、目的、コマンド、構造、コード スタイル、テスト、境界をカバーする PRD を作成します。
新しいプロジェクト、機能、または重要な変更を開始する
計画 - 細分化する
スキル
何をするのか
いつ使用するか
計画とタスクの内訳
承認基準と依存関係の順序付けを使用して、仕様を小さな検証可能なタスクに分解します。
仕様があり、実装可能なユニットが必要です
ビルド - コードを書く
スキル
何をするのか
いつ使用するか
増分実装
薄い垂直スライス - 実装、テスト、検証、コミット。機能フラグ、安全なデフォルト、ロールバックしやすい変更
複数のファイルに関わる変更
テスト駆動開発
Red-Green-Refactor、テストピラミッド (80/15/5)、テストサイズ、DAMP over DRY、Beyonce Rule、ブラウザテスト
ロジックの実装、バグの修正、または動作の変更
コンテキストエンジニアリング
エージェントに適切な情報を適切なタイミングで提供 - ルール ファイル、コンテキスト パッキング、MCP 統合
セッションの開始時、タスクの切り替え時、または出力品質が低下したとき
ソース駆動開発
フレームワークに関するすべての決定を公式文書に根拠付ける - 検証し、出典を引用し、未検証のものにフラグを立てる
フレームワークまたはライブラリに対して、ソースが引用された信頼できるコードが必要な場合
疑念主導の開発

進行中のすべての重要な決定に対する敵対的な新鮮なコンテキストのレビュー - 主張 → 抽出 → 疑い → 調整 → 停止 (オプションのユーザー承認のモデル間エスカレーションあり)
リスクが高い (運用、セキュリティ、元に戻せない)、なじみのないコードで作業している、または自信のある出力は、後でデバッグするよりも今すぐ検証する方がコストがかかりません。
フロントエンド-UI-エンジニアリング
コンポーネント アーキテクチャ、デザイン システム、状態管理、レスポンシブ デザイン、WCAG 2.1 AA アクセシビリティ
ユーザー向けインターフェースの構築または変更
API とインターフェイスの設計
コントラクトファースト設計、ハイラムの法則、ワンバージョンルール、エラーセマンティクス、境界検証
API、モジュール境界、またはパブリックインターフェイスの設計
検証 - 動作することを証明する
スキル
何をするのか
いつ使用するか
devtools を使用したブラウザーのテスト
ライブ ランタイム データ用の Chrome DevTools MCP - DOM 検査、コンソール ログ、ネットワーク トレース、パフォーマンス プロファイリング
ブラウザ内で実行されるものをビルドまたはデバッグする
デバッグとエラー回復
5 段階のトリアージ: 再現、位置特定、削減、修正、保護。一時停止ルール、安全なフォールバック
テストが失敗する、ビルドが中断する、または予期しない動作が発生する
レビュー - マージ前の品質ゲート
スキル
何をするのか
いつ使用するか
コードレビューと品質
5 軸レビュー、サイズ変更 (~100 行)、重大度ラベル (Nit/Optional/FYI)、レビュー速度基準、分割戦略
変更をマージする前に
コードの簡略化
チェスタトンのフェンス、ルール オブ 500 は、正確な動作を維持しながら複雑さを軽減します
コードは機能しますが、本来よりも読みにくく、保守するのが困難です
セキュリティと強化
OWASP トップ 10 防止、認証パターン、シークレット管理、依存関係監査、3 層境界システム
ユーザー入力、認証、データストレージ、または外部統合の処理
パフォーマンスの最適化
測定優先のアプローチ - コア Web Vitals ターゲット、プロファイリング ワークフロー、バンドル

解析、アンチパターン検出
パフォーマンス要件が存在するか、リグレッションが疑われる場合
出荷 - 自信を持って展開
スキル
何をするのか
いつ使用するか
git-ワークフローとバージョン管理
トランクベースの開発、アトミックコミット、サイズ変更 (~100 行)、セーブポイントとしてのコミットパターン
コード変更を行う (常に)
ci-cd-and-automation
シフトレフト、速いほど安全、機能フラグ、高品質ゲートパイプライン、障害フィードバックループ
ビルドおよびデプロイのパイプラインの設定または変更
非推奨と移行
責任としてのコードの考え方、強制的な非推奨と勧告による非推奨、移行パターン、ゾンビ コードの削除
古いシステムの削除、ユーザーの移行、または機能の廃止
ドキュメントと広告
アーキテクチャ決定記録、API ドキュメント、インラインドキュメント標準 - 理由を文書化します。
アーキテクチャ上の決定、API の変更、または機能の出荷
可観測性と計測
構造化ロギング、RED メトリクス、OpenTelemetry トレース、症状ベースのアラート - 構築時の計測機能
テレメトリの追加、または運用環境で実行されるものを出荷する
出荷と打ち上げ
リリース前チェックリスト、機能フラグのライフサイクル、段階的ロールアウト、ロールバック手順、モニタリング設定
実稼働環境へのデプロイの準備
運用 - システムを稼働し続ける
スキル
何をするのか
いつ使用するか
カオスエンジニアリング
体系的なフォールト挿入と回復力テスト
高可用性を考慮した設計、災害復旧の検証
コストの最適化
信頼性を犠牲にすることなくクラウド支出を削減
請求額は予想外に増加し、リソースの適正化が図られる
データエンジニアリング
データ パイプライン、ETL/ELT、スキーマ進化、

[切り捨てられた]

## Original Extract

Production-grade engineering skills for AI coding agents. - borhen68/SkillEngine

GitHub - borhen68/SkillEngine: Production-grade engineering skills for AI coding agents. · GitHub
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
borhen68
/
SkillEngine
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits .claude-plugin .claude-plugin .claude/ commands .claude/ commands .gemini/ commands .gemini/ commands .github .github .opencode .opencode agents agents assets assets commands commands docs docs hooks hooks references references scripts scripts skills skills .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .markdownlint-cli2.yaml .markdownlint-cli2.yaml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md package.json package.json plugin.json plugin.json View all files Repository files navigation
Forge production-grade AI agents.
28 battle-tested skills. 6 specialist personas. 5 reference checklists. One mission: make AI agents build software like senior engineers.
AI coding agents are fast. They're also reckless.
They skip specs. They "forget" tests. They ship without review. They treat "it works on my machine" as a success criteria. In short, they build prototypes, not production software.
We don't give agents vague suggestions. We give them structured, battle-tested workflows that encode how senior engineers actually build software — the same workflows that power teams at Google, Netflix, and Stripe. Every skill has steps, checkpoints, anti-rationalization defenses, and evidence-based verification. When an agent follows these, it ships code you can trust.
DEFINE PLAN BUILD VERIFY REVIEW SHIP OPS
┌──────┐ ┌──────┐ ┌──────┐ ┌──────┐ ┌──────┐ ┌──────┐ ┌──────┐
│ Idea │ ───▶ │ Spec │ ───▶ │ Code │ ───▶ │ Test │ ───▶ │ QA │ ───▶ │ Go │ ───▶ │ Run │
│Refine│ │ PRD │ │ Impl │ │Debug │ │ Gate │ │ Live │ │ Ops │
└──────┘ └──────┘ └──────┘ └──────┘ └──────┘ └──────┘ └──────┘
/spec /plan /build /test /review /ship
7 slash commands. 28 skills. 6 phases. Zero excuses.
Other Prompt Packs
AgentForge
Structure
Vague advice
Step-by-step workflows with checkpoints
Verification
"Make sure it works"
Evidence-based exit criteria (tests, builds, runtime data)
Anti-cheating
None
Rationalization tables that call out excuses agents use to skip steps
Scope
Generic coding tips
Full lifecycle: spec → plan → build → verify → review → ship → ops
Quality gates
None
Built-in CI pipeline with 8 automated checks
Cross-reference
Silos
Every skill references related skills; no duplication
Commands
7 slash commands that map to the development lifecycle. Each one activates the right skills automatically.
Want fewer manual steps once the spec exists? /build auto generates the plan and implements every task in a single approved pass — you approve the plan once, then it runs autonomously. It removes the human stepping between tasks, not the verification: every task is still test-driven and committed individually, and it pauses on failures or risky steps.
Skills also activate automatically based on what you're doing — designing an API triggers api-and-interface-design , building UI triggers frontend-ui-engineering , and so on.
Claude Code (recommended)
Marketplace install:
/plugin marketplace add borhen68/SkillEngine
/plugin install agentforge@borhen-agentforge
SSH errors? The marketplace clones repos via SSH. If you don't have SSH keys set up on GitHub, either add your SSH key or use the full HTTPS URL to force the HTTPS cloning:
/plugin marketplace add https://github.com/borhen68/SkillEngine.git
/plugin install agentforge@borhen-agentforge
git clone https://github.com/borhen68/SkillEngine.git
claude --plugin-dir /path/to/agentforge
Cursor
Copy any SKILL.md into .cursor/rules/ , or reference the full skills/ directory. See docs/cursor-setup.md .
Install as a native plugin for skills, subagents, and slash commands. See docs/antigravity-setup.md .
agy plugin install https://github.com/borhen68/SkillEngine.git
Install from a local clone:
git clone https://github.com/borhen68/SkillEngine.git
agy plugin install ./agentforge
Gemini CLI
Install as native skills for auto-discovery, or add to GEMINI.md for persistent context. See docs/gemini-cli-setup.md .
gemini skills install https://github.com/borhen68/SkillEngine.git --path skills
Install from a local clone:
gemini skills install ./agentforge/skills/
Windsurf
Add skill contents to your Windsurf rules configuration. See docs/windsurf-setup.md .
Uses agent-driven skill execution via AGENTS.md and the skill tool.
Use agent definitions from agents/ as Copilot personas and skill content in .github/copilot-instructions.md . See docs/copilot-setup.md .
Skills are plain Markdown - they work with any agent that accepts system prompts or instruction files. See docs/getting-started.md .
The commands above are entry points. The pack includes 28 skills total — 24 lifecycle skills, 4 operations skills, plus the using-agentforge meta-skill. Each skill is a structured workflow with steps, verification gates, and anti-rationalization tables. You can also reference any skill directly.
Meta - Discover which skill applies
Skill
What It Does
Use When
using-agentforge
Maps incoming work to the right skill workflow and defines shared operating rules
Starting a session or deciding which skill applies
Define - Clarify what to build
Skill
What It Does
Use When
interview-me
One-question-at-a-time interview that extracts what the user actually wants instead of what they think they should want, until ~95% confidence
The ask is underspecified, or the user invokes "interview me" / "grill me"
idea-refine
Structured divergent/convergent thinking to turn vague ideas into concrete proposals
You have a rough concept that needs exploration
spec-driven-development
Write a PRD covering objectives, commands, structure, code style, testing, and boundaries before any code
Starting a new project, feature, or significant change
Plan - Break it down
Skill
What It Does
Use When
planning-and-task-breakdown
Decompose specs into small, verifiable tasks with acceptance criteria and dependency ordering
You have a spec and need implementable units
Build - Write the code
Skill
What It Does
Use When
incremental-implementation
Thin vertical slices - implement, test, verify, commit. Feature flags, safe defaults, rollback-friendly changes
Any change touching more than one file
test-driven-development
Red-Green-Refactor, test pyramid (80/15/5), test sizes, DAMP over DRY, Beyonce Rule, browser testing
Implementing logic, fixing bugs, or changing behavior
context-engineering
Feed agents the right information at the right time - rules files, context packing, MCP integrations
Starting a session, switching tasks, or when output quality drops
source-driven-development
Ground every framework decision in official documentation - verify, cite sources, flag what's unverified
You want authoritative, source-cited code for any framework or library
doubt-driven-development
Adversarial fresh-context review of every non-trivial decision in-flight - CLAIM → EXTRACT → DOUBT → RECONCILE → STOP, with optional user-authorized cross-model escalation
Stakes are high (production, security, irreversible), working in unfamiliar code, or a confident output is cheaper to verify now than to debug later
frontend-ui-engineering
Component architecture, design systems, state management, responsive design, WCAG 2.1 AA accessibility
Building or modifying user-facing interfaces
api-and-interface-design
Contract-first design, Hyrum's Law, One-Version Rule, error semantics, boundary validation
Designing APIs, module boundaries, or public interfaces
Verify - Prove it works
Skill
What It Does
Use When
browser-testing-with-devtools
Chrome DevTools MCP for live runtime data - DOM inspection, console logs, network traces, performance profiling
Building or debugging anything that runs in a browser
debugging-and-error-recovery
Five-step triage: reproduce, localize, reduce, fix, guard. Stop-the-line rule, safe fallbacks
Tests fail, builds break, or behavior is unexpected
Review - Quality gates before merge
Skill
What It Does
Use When
code-review-and-quality
Five-axis review, change sizing (~100 lines), severity labels (Nit/Optional/FYI), review speed norms, splitting strategies
Before merging any change
code-simplification
Chesterton's Fence, Rule of 500, reduce complexity while preserving exact behavior
Code works but is harder to read or maintain than it should be
security-and-hardening
OWASP Top 10 prevention, auth patterns, secrets management, dependency auditing, three-tier boundary system
Handling user input, auth, data storage, or external integrations
performance-optimization
Measure-first approach - Core Web Vitals targets, profiling workflows, bundle analysis, anti-pattern detection
Performance requirements exist or you suspect regressions
Ship - Deploy with confidence
Skill
What It Does
Use When
git-workflow-and-versioning
Trunk-based development, atomic commits, change sizing (~100 lines), the commit-as-save-point pattern
Making any code change (always)
ci-cd-and-automation
Shift Left, Faster is Safer, feature flags, quality gate pipelines, failure feedback loops
Setting up or modifying build and deploy pipelines
deprecation-and-migration
Code-as-liability mindset, compulsory vs advisory deprecation, migration patterns, zombie code removal
Removing old systems, migrating users, or sunsetting features
documentation-and-adrs
Architecture Decision Records, API docs, inline documentation standards - document the why
Making architectural decisions, changing APIs, or shipping features
observability-and-instrumentation
Structured logging, RED metrics, OpenTelemetry tracing, symptom-based alerting - instrument as you build
Adding telemetry, or shipping anything that runs in production
shipping-and-launch
Pre-launch checklists, feature flag lifecycle, staged rollouts, rollback procedures, monitoring setup
Preparing to deploy to production
Ops - Keep systems running
Skill
What It Does
Use When
chaos-engineering
Systematic fault injection and resilience testing
Designing for high availability, verifying disaster recovery
cost-optimization
Cloud spend reduction without sacrificing reliability
Bills growing unpredictably, rightsizing resources
data-engineering
Data pipelines, ETL/ELT, schema evolution,

[truncated]
