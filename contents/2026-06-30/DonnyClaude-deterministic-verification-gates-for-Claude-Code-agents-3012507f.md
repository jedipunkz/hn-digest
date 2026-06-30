---
source: "https://github.com/d0nmega/donnyclaude"
hn_url: "https://news.ycombinator.com/item?id=48734053"
title: "DonnyClaude – deterministic verification gates for Claude Code agents"
article_title: "GitHub - D0NMEGA/donnyclaude: Prompt, context, harness, and loop engineering for Claude Code - in one config. · GitHub"
author: "d0nmega"
captured_at: "2026-06-30T15:49:18Z"
capture_tool: "hn-digest"
hn_id: 48734053
score: 1
comments: 0
posted_at: "2026-06-30T15:23:43Z"
tags:
  - hacker-news
  - translated
---

# DonnyClaude – deterministic verification gates for Claude Code agents

- HN: [48734053](https://news.ycombinator.com/item?id=48734053)
- Source: [github.com](https://github.com/d0nmega/donnyclaude)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T15:23:43Z

## Translation

タイトル: DonnyClaude – クロード コード エージェントの決定論的検証ゲート
記事のタイトル: GitHub - D0NMEGA/donnyclaude: クロード コードのプロンプト、コンテキスト、ハーネス、ループ エンジニアリング - 1 つの構成で。 · GitHub
説明: クロード コードのプロンプト、コンテキスト、ハーネス、およびループ エンジニアリング - 1 つの構成。 - D0NMEGA/ドニークロード

記事本文:
GitHub - D0NMEGA/donnyclaude: クロード コードのプロンプト、コンテキスト、ハーネス、ループ エンジニアリング - 1 つの構成で。 · GitHub
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
D0Nメガ
/
ドニークロード
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション操作

ション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
77 コミット 77 コミット アセット アセット bin bin docs docs パッケージ パッケージ テンプレート テンプレート テスト テスト .gitignore .gitignore ライセンス ライセンス README.md README.md package.json package.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
クロード コードのプロンプト、コンテキスト、ハーネス、ループ エンジニアリングを 1 つの構成で実行します。
AI コーディングのセットアップがおもちゃであるかツールであるかを決定する 4 つの分野
通常、散在するブログ投稿やプライベート ドットファイルに存在します。ドニークロードが集まる
4 つすべてを 1 つのインストール可能な構成に統合: 慎重に設計されたエージェント プロンプト、
/clear を維持するコンテキスト層。
モデルと、未完了の作業を完了としてマークすることを拒否する検証ループ。
npxドニークロード
1 つのコマンドで Donny ワークフロー エンジン、94 個のスキル、
48 の専門エージェント、13 言語のコーディング ルール、プロジェクト フック、厳選された
MCP セットアップ - 次に、Claude 自体が最初のプロジェクトを説明します。
48 人のエージェントと 94 のスラッシュ コマンド スキル、それぞれが意図的に設計されたプロンプトで、
単一の責任と最小限のツールの許可。レビュー担当者、ビルド修正担当者、
研究者、プランナー、検証者 - 適切なプロンプトが実行されるように名前と範囲が設定されます
過負荷になった 1 つのシステム プロンプトがすべてを実行しようとするのではなく、適切なジョブを実行します。
すべての重要なタスクは、.planning/ - 要件に基づいてその状態をディスクに書き込みます。
ロードマップ、フェーズごとの計画、概要、検証レポート - コンテキスト
/clear および新しいセッションは存続します。サブエージェントはそれぞれ、厳選された分離されたスライスを取得します
トランスクリプト全体ではなく、そのコンテキストの部分です。 Context7 MCP はライブで提供します
ライブラリ ドキュメントを参照して、モデルが古いトレーニング データではなく、現在の API に基づいてコード化されるようにします。あ
グローバル操作ガイドには、これまでのコーディングと記述の標準がロードされています

yセッションなので、
生成されたコードと散文は、各プロンプトに貼り付けられることなく、それに続きます。
エージェントの下には、Donny エンジンが配置されています。これは、決定論的なノード CLI です。
ワークフローは、言語モデルが推測すべきではないことを要求します - 計画
依存関係グラフの検証、フロントマター スキーマ、要件カバレッジ、シークレット
scanning, phase completion. Model-tiered subagents (Opus for planning and
検証、機械的作業の軽量層）コストを
タスク。フックは、毎ターン、フォーマット、ガード、および状態の整合性を強制します。
作業単位はループです。次のフェーズを発見し、計画し、実行します。
ウェーブパラレルサブエージェント、それを検証し、出荷し、繰り返します。 The generator never grades
独自の作業 - 別の懐疑的な検証者が実行して目標の達成をチェックします
コード、さらにその下には決定論的なエンジン ゲートがあるため、評決は
cannot drift from the truth. Run a single phase by hand, or hand the whole
自律モードへのロードマップを作成し、無人で進めます。
Donny は中核となるワークフロー エンジンであり、計画と実行のループです。
明示的なフェーズを通じて、アイデアを出荷済みの検証済みコードに変換します。その署名
は決定論的検証です。ほとんどのエージェント ワークフローでは LLM が宣言します。
ドニーは「良さそうですね」とエンジンによる強制チェックですべてのゲートを支えます。
verify フェーズ-検証済み - フェーズは、検証ステータスが次の場合にのみ出荷されます。
実際に渡され、ディスクから解析されたものであり、LLM の指示ではありません。
マイルストーン カバレッジの検証 - プラン全体で計算された要件カバレッジ、
summaries, and verifications, not hand-tallied.
脅威の検証 - クリア / セキュリティ スキャン - 秘密 - オープンな脅威と漏洩
credentials block the commit deterministically.
フェーズ遅延 - スキップされたフェーズは、正直に遅延として記録され、決して記録されません。
silently marked complete.
The result is a loop you can leave r

ウンニング：それは、あるものについての真実を教えてくれます。
完了し、何かが完了していない場合はコールドが停止します。
コンポーネント
カウント
それは何ですか
スキル
94
スラッシュ コマンド - Donny ワークフローとユーティリティおよび言語パック
エージェント
48
専門のサブエージェント - プランナー、レビュー担当者、検証担当者、ビルド修正担当者
ルール
13
コーディングおよびライティング標準 (共通 + 12 言語パック)、自動ロード
フック
29
各ターンで実行されるフォーマット、ガード、および状態整合性のフック
MCP
2
context7 (ライブドキュメント) と playwright (ブラウザ)、ユーザースコープで登録
エンジン
1
Donny CLI とワークフロー ライブラリ
構成
1
ルールをロードするグローバル操作ガイド ( ~/.claude/CLAUDE.md )。既存のものを決して壊さない
すべては ~/.claude/ に格納されます。既存の設定は上書きされるのではなく、保持されます。
npx donnyclaude # インストールしてからセットアップ ウィザードを実行します
# 次に、クロード コードでは次のようになります。
/donny-init # 研究 -> 要件 -> ロードマップ
/donny-plan-phase 1 # フェーズを計画します (依存関係 + 要件ゲートあり)
/donny-execute-phase 1 # ビルドします (ウェーブ並列サブエージェント、アトミックコミット)
/donny-verify-work 1 # 会話型 UAT
/donny-ship # 検証に合格したら PR を開きます
いつでも /donny-progress を使用して、現在どこにいるのか、次に何が起こるのかを確認できます。
/donny-autonomous はロードマップ全体を手動で進めます。 /donny-ヘルプリスト
すべてのコマンド。
クロード・コード (インストーラー)
見つからない場合はインストールします）
バンドルされたリサーチ スクレイパー用の uv (オプション、初回実行時のセルフ ブートストラップ)
リサーチ スクレイパー (バンドル) - マルチソース HTTP リサーチ ダイジェスト。参照
docs/research-tools.md 。
Obsidian - 耐久性のある記憶としての保管庫。インストーラーがインストールを提案します。
docs/obsidian-memory.md を参照してください。
ブラウザハーネス - ログインゲート型調査用のオプション。個別にインストールされます。
docs/research-tools.md を参照してください。
DonnyClaude は MIT ライセンスを取得しています。 Donny エンジンは GSD ワークフローのフォークとして始まりました
エンジンと罪がある

CE は大幅に分岐 - 決定論を中心に再構築
検証ゲート、モデル階層化されたサブエージェント、サブエージェントに安全な調査パス。
オプションのブラウザハーネス統合リファレンス
browser-use/browser-harness
(MIT, Copyright (c) 2026 Browser Use); it is not bundled, and all credit for that
project belongs to its authors.
クロード コードのプロンプト、コンテキスト、ハーネス、ループ エンジニアリングを 1 つの構成で実行します。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Prompt, context, harness, and loop engineering for Claude Code - in one config. - D0NMEGA/donnyclaude

GitHub - D0NMEGA/donnyclaude: Prompt, context, harness, and loop engineering for Claude Code - in one config. · GitHub
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
D0NMEGA
/
donnyclaude
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
77 Commits 77 Commits assets assets bin bin docs docs packages packages templates templates tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md package.json package.json View all files Repository files navigation
Prompt, context, harness, and loop engineering for Claude Code - in one config.
The four disciplines that decide whether an AI coding setup is a toy or a tool
usually live in scattered blog posts and private dotfiles. DonnyClaude assembles
all four into one installable config: carefully engineered agent prompts, a
context layer that survives /clear , a deterministic harness that drives the
models, and a verification loop that refuses to mark unfinished work as done.
npx donnyclaude
One command installs the Donny workflow engine, 94 skills,
48 specialized agents, coding rules for 13 languages, project hooks, and a curated
MCP setup - then Claude itself walks you through your first project.
48 agents and 94 slash-command skills, each a deliberately engineered prompt with
a single responsibility and a minimal tool grant. Reviewers, build-fixers,
researchers, planners, and verifiers - named and scoped so the right prompt runs
for the right job instead of one overloaded system prompt trying to do everything.
Every non-trivial task writes its state to disk under .planning/ - requirements,
roadmap, per-phase plans, summaries, and verification reports - so context
survives /clear and new sessions. Subagents each get a curated, isolated slice
of that context rather than the whole transcript. The Context7 MCP supplies live
library docs so the model codes against current APIs, not stale training data. A
global operating guide loads the coding and writing standards every session, so
generated code and prose follow them without being pasted into each prompt.
Under the agents sits the Donny engine: a deterministic Node CLI that the
workflows call for the things a language model should not guess at - plan
dependency-graph validation, frontmatter schemas, requirement coverage, secret
scanning, phase completion. Model-tiered subagents (Opus for planning and
verification, lighter tiers for mechanical work) keep cost proportional to the
task. Hooks enforce formatting, guards, and state integrity on every turn.
The unit of work is a loop: discover the next phase, plan it, execute it with
wave-parallel subagents, verify it, ship it, repeat. The generator never grades
its own work - a separate skeptical verifier checks goal achievement by running
code, and beneath even that sits a deterministic engine gate so the verdict
cannot drift from the truth. Run a single phase by hand, or hand the whole
roadmap to autonomous mode and let it advance unattended.
Donny is the workflow engine at the core - a planning-and-execution loop that
turns an idea into shipped, verified code through explicit phases. Its signature
is deterministic verification : where most agent workflows let an LLM declare
"looks good," Donny backs every gate with engine-enforced checks.
verify phase-verified - a phase ships only when its verification status is
actually passed , parsed from disk, never an LLM's say-so.
verify milestone-coverage - requirement coverage computed across plans,
summaries, and verifications, not hand-tallied.
verify threats-clear / security scan-secrets - open threats and leaked
credentials block the commit deterministically.
phase defer - a skipped phase is recorded honestly as deferred, never
silently marked complete.
The result is a loop you can leave running: it tells you the truth about what is
done, and stops cold when something is not.
Component
Count
What it is
Skills
94
Slash commands - the Donny workflow plus utilities and language packs
Agents
48
Specialized subagents - planners, reviewers, verifiers, build-fixers
Rules
13
Coding and writing standards (common + 12 language packs), auto-loaded
Hooks
29
Format, guard, and state-integrity hooks that run on each turn
MCP
2
context7 (live docs) and playwright (browser), registered at user scope
Engine
1
The Donny CLI and workflow library
Config
1
A global operating guide ( ~/.claude/CLAUDE.md ) that loads the rules; never clobbers an existing one
Everything lands in ~/.claude/ . Existing settings are preserved, not clobbered.
npx donnyclaude # install, then the setup wizard
# then, in Claude Code:
/donny-init # research -> requirements -> roadmap
/donny-plan-phase 1 # plan a phase (with dependency + requirement gates)
/donny-execute-phase 1 # build it (wave-parallel subagents, atomic commits)
/donny-verify-work 1 # conversational UAT
/donny-ship # open a PR once verification passes
Use /donny-progress any time to see where you are and what is next, or
/donny-autonomous to advance the whole roadmap hands-off. /donny-help lists
every command.
Claude Code (the installer
installs it if missing)
uv for the bundled research scrapers (optional; self-bootstraps on first run)
Research scrapers (bundled) - a multi-source HTTP research digest. See
docs/research-tools.md .
Obsidian - a vault as durable memory. The installer offers to install it.
See docs/obsidian-memory.md .
browser-harness - optional, for login-gated research. Installed separately;
see docs/research-tools.md .
DonnyClaude is MIT licensed. The Donny engine began as a fork of the GSD workflow
engine and has since diverged substantially - rebuilt around deterministic
verification gates, model-tiered subagents, and a subagent-safe research path.
The optional browser-harness integration references
browser-use/browser-harness
(MIT, Copyright (c) 2026 Browser Use); it is not bundled, and all credit for that
project belongs to its authors.
Prompt, context, harness, and loop engineering for Claude Code - in one config.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
