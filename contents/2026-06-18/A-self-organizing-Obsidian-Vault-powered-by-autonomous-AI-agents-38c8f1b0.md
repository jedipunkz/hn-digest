---
source: "https://github.com/ibrahimkobeissy/ai-second-brain-template"
hn_url: "https://news.ycombinator.com/item?id=48583182"
title: "A self-organizing Obsidian Vault powered by autonomous AI agents"
article_title: "GitHub - ibrahimkobeissy/ai-second-brain-template: A self-governing Personal Knowledge Management (PKM) framework for Obsidian, featuring automated agent workflows for research, synthesis, and task extraction. · GitHub"
author: "koubeissy_i"
captured_at: "2026-06-18T10:34:52Z"
capture_tool: "hn-digest"
hn_id: 48583182
score: 2
comments: 0
posted_at: "2026-06-18T10:04:59Z"
tags:
  - hacker-news
  - translated
---

# A self-organizing Obsidian Vault powered by autonomous AI agents

- HN: [48583182](https://news.ycombinator.com/item?id=48583182)
- Source: [github.com](https://github.com/ibrahimkobeissy/ai-second-brain-template)
- Score: 2
- Comments: 0
- Posted: 2026-06-18T10:04:59Z

## Translation

タイトル: 自律型 AI エージェントを活用した自己組織的な Obsidian Vault
記事のタイトル: GitHub - ibrahimkobeissy/ai-second-brain-template: Obsidian 用の自律的なパーソナル ナレッジ管理 (PKM) フレームワークで、調査、合成、タスク抽出のための自動化されたエージェント ワークフローを備えています。 · GitHub
説明: Obsidian 用の自律型パーソナル ナレッジ管理 (PKM) フレームワーク。調査、合成、タスク抽出のための自動化されたエージェント ワークフローを備えています。 - ibrahimkobeissy/ai-second-brain-template

記事本文:
GitHub - ibrahimkobeissy/ai-second-brain-template: Obsidian 用の自律的なパーソナル ナレッジ管理 (PKM) フレームワークで、調査、合成、タスク抽出のための自動化されたエージェント ワークフローを備えています。 · GitHub
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
イブラヒムコベイシー
/
あい～

第二の脳テンプレート
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
ibrahimkobeissy/ai-セカンドブレイン-テンプレート
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット .claude .claude vault vault .gitignore .gitignore CLAUDE.md CLAUDE.md GEMINI.md GEMINI.md ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
この文書は、この Second Brain の中核となる哲学、ワークフロー、および運用ロジックを定義します。これは生きたドキュメントであり、システムとそのツールの進化に応じて更新されます。
正直で客観的な思考 : これは第 2 の脳です。仕事は、同意できる解決策ではなく、最適な解決策を見つけることです。エージェントは（他のエージェントの仕事や自分自身の仕事を含む）弱い仕事に異議を唱え、何かが間違っている場合にははっきりと異議を唱え、本当に間違っていることが証明されない限り反発を受けても自分の立場を守り、検証された証拠に基づいて主張を根拠付けます。決してお世辞を言わないでください。平和を維持するために平均的な仕事を決して超えないでください。
大規模よりも軽量 : ボールトをスリムで信号の多い状態に保ちます。使用済みの中間ファイル (一度キュレーションされたブックマーク、一度合成されたドラフト) を削除します。git は履歴を保存するため、削除は安全であり、エージェントは無効なファイルにトークンを書き込むべきではありません。耐久性のある出力のみを保持します。大きなグラフはバニティメトリクスです。
厳密な分離 (ハードウォール) : コンテキストの流出を防ぐために、仕事、個人、およびリソースのフォルダーを厳密に分離します。ただし、タイプ フロントマター プロパティ (例: evergreen 、 Synthetic ) を介してクロスドメイン推論が有効になり、クロスエリアの洞察が可能になります。
リンクファースト アーキテクチャ : 知識の価値はコンテンツだけではなく、接続 ( [[wikilinks]] ) に存在します。
エージェント主導ではなくエージェント拡張: AI エージェント (クロード/反重力) が作業 (キュレーション、合成、リンティング) を自動化する一方、人間は Fi を保持します。

ナル理解。
00-inbox/ : 生のキャプチャ、Web クリッピング、およびつかの間のメモ。
01-work/ : 職業生活に関連する責任のある分野と積極的な取り組み。
02-personal/ : 私生活に関連する興味のある分野と生活管理。
03-resources/ : 特定の責任に関連付けられていない参考ライブラリと興味深いトピック。
04-archive/ : 非アクティブな領域またはリソース。冷蔵。
99-system/ : メタデータ、テンプレート、添付ファイル、およびシステム ドキュメント。
3. 運用ワークフロー（ループ）
Second Brain のコア エンジンは、生データを取得し、それを処理して実用的な洞察を生成し、無駄を取り除くという継続的なループです。
フローチャート TD
%% ノード
受信箱["00-受信箱 (生のブックマーク)"]
Curate["curate-bookmarks (スキル)"]
ドラフト["ドラフト/ (厳選されたノート)"]
Synthesize["合成ドラフト(スキル)"]
計画[「総合/(戦略計画)」]
抽出[「計画からかんばん(スキル)」]
かんばん["todo-kanban.md (実行可能なタスク)"]
実行["ビルド/実行"]
Prune["中間ファイルの削除 (ブックマークと下書き)"]
%% フロー
受信箱 --> キュレート
キュレート -->|信号を抽出|草案
ドラフト --> 合成
合成 -->|科学テーマの合成|計画
計画 --> 抽出
抽出 -->|重複排除と追加|カンバン
カンバン --> 実行
実行 --> プルーニング
classDef ファイルの塗りつぶし:#2d2d2d、ストローク:#555、ストローク幅:1px、色:#fff
classDef スキル fill:#1a3f5c、ストローク:#4a90e2、ストローク幅:2px、カラー:#fff
classDef アクション fill:#1e4620、ストローク:#4caf50、ストローク幅:2px、カラー:#fff
クラス 受信箱、ドラフト、計画、かんばんファイル
クラス キュレーション、合成、抽出スキル
クラス実行、プルーンアクション
読み込み中
ワークフローの段階:
キャプチャ : 原材料は 00-inbox/ に到着します。
Curate ( curate-bookmarks ) : エージェントは受信箱のアイテムを読み取り、核となる価値 (「盗めるもの」) を抽出し、それらを特定のエリア内のドラフト/フォルダーに移動します。ソースはlです

Processed-sources.md に記録されています。
合成 ( synthesize-drafts ) : エージェントは複数の草案を取得し、科学的テーマ マトリックスを使用して相互に分析し、統一された戦略計画 ( Synthetic/ ) を生成します。
アクション ( plan-to-kanban ) : エージェントは戦略計画を読み取り、実行可能なタスクを抽出して重複を排除し、エリアの todo-kanban.md に追加します。
Clean : ナレッジが永続的で実用的になると、使用済みの中間ファイル (元のブックマークと下書き) が削除されます。
4. 操作スキル（ツールベルト）
init-area : アイデアに挑戦し、目標/範囲を定義し、必要なハブ ノート、カンバン ボード、フォルダーを足場にして、新しいエリアを対話的に作成します。
スカウトアイデア : 新しいアイデアを検証し、その有用性を検討し、それらを構築するための外部リソース/ツールをスカウトします。
curate-bookmarks : 受信箱のアイテムを実行可能なエリアの下書きに処理します。
synthesize-drafts : 科学的なテーマの統合を使用して、エリア内の複数のドラフトを戦略的な「グローバル プラン」に統合します。
plan-to-kanban : 合成ドキュメントのアクション プランを読み取り、アクション アイテムをエリアのカンバン ボードに抽出して重複を排除します。
vault-linter : 読み取り専用のナレッジグラフの整合性チェック — 壊れた [[wikilinks]] 、孤立したメモ、ソース/captured_from トレーサビリティの欠落。決して編集しません。
Audit-maintenance : 保留中のメンテナンス タスクと他のエージェントが作成したピアレビュー ツールをヘッドレスでレビューします。
レビュー ループ: vault/99-system/maintenance/agent-kanban.md は、Todo / In Progress / Done / Archived のスイムレーンを持つカンバン ボードです。すべてのツールの作成は、 Todo の下に新しいカードとして記録する必要があります。
セッション チェック: SessionStart フックは、 .claude/hooks/pending-reviews.sh <Reviewer> を呼び出すことで、セッション開始時に各エージェントの保留中のレビューを表示します。
品質管理 : レビュー c

他のエージェントの作業を強化し、フォーマットに従っているかどうかだけでなく、最高品質の出力を生成しているかどうかを判断します。パスが獲得されます。弱いツールは具体的な改善が必要で失敗します。
真実の情報源 : エージェントは、構造的な変更を行う前に CLAUDE.md とこの文書を読む必要があります。
直接書き込み禁止: エージェントはドラフト/フォルダーまたは特定のシステム ディレクトリに書き込みます。確認なしにエリアのコアに直接書き込むことはありません。
トレーサビリティ : エージェントが作成したすべてのメモには、source フィールドまたは Captured_from フィールドが含まれている必要があります。
セキュリティ ガードレール : エージェントはプロジェクト ディレクトリ内に留まり、資格情報や秘密のパスの読み取り、書き込み、漏洩を行いません。 CLAUDE.md §10 に従って適用されます。
システムの堅牢性と文書化を確保するために、大きな変更が行われるたびにこのチェックリストがトリガーされます。
Sync README.md : 新しい機能または構造の変更を文書化します。
Log Agent-kanban.md : Todo の下にカードを追加し、他のエージェントに割り当てます。
ARA の監査: 「プロジェクト」フォルダーが作成されていないことを確認します。
Obsidian 用の自律的なパーソナル ナレッジ管理 (PKM) フレームワーク。調査、合成、タスク抽出のための自動化されたエージェント ワークフローを備えています。
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

A self-governing Personal Knowledge Management (PKM) framework for Obsidian, featuring automated agent workflows for research, synthesis, and task extraction. - ibrahimkobeissy/ai-second-brain-template

GitHub - ibrahimkobeissy/ai-second-brain-template: A self-governing Personal Knowledge Management (PKM) framework for Obsidian, featuring automated agent workflows for research, synthesis, and task extraction. · GitHub
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
ibrahimkobeissy
/
ai-second-brain-template
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
ibrahimkobeissy/ai-second-brain-template
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits .claude .claude vault vault .gitignore .gitignore CLAUDE.md CLAUDE.md GEMINI.md GEMINI.md LICENSE LICENSE README.md README.md View all files Repository files navigation
This document defines the core philosophy, workflows, and operational logic of this Second Brain. It is a living document, updated as the system and its tools evolve.
Honest & Objective Thinking : This is a second brain — the work is finding the best solution, not an agreeable one. Agents challenge weak work (including another agent's and their own), object plainly when something is wrong, hold their position under pushback unless genuinely proven wrong, and ground claims in verified evidence. Never flatter; never pass average work to keep the peace.
Lite over Large : Keep the vault lean and high-signal. Delete spent intermediates (a bookmark once curated, a draft once synthesized) — git preserves history, so deletion is safe, and agents shouldn't burn tokens on dead files. Keep only durable outputs; a big graph is a vanity metric.
Strict Separation (The Hard Wall) : Keep Work, Personal, and Resources folders strictly separated to prevent context bleed. However, cross-domain reasoning is enabled via the type frontmatter property (e.g. evergreen , synthesis ), allowing cross-Area insights.
Link-First Architecture : Knowledge value lives in the connections ( [[wikilinks]] ), not just the content.
Agent-Augmented, Not Agent-Led : AI agents (Claude/Antigravity) automate the labor (curation, synthesis, linting) while the human retains the final understanding.
00-inbox/ : Raw captures, web clippings, and fleeting notes.
01-work/ : Areas of responsibility and active efforts related to professional life.
02-personal/ : Areas of interest and life management related to personal life.
03-resources/ : Reference library and topics of interest not tied to a specific responsibility.
04-archive/ : Inactive areas or resources; cold storage.
99-system/ : Metadata, templates, attachments, and system documentation.
3. The Operating Workflow (The Loop)
The core engine of the Second Brain is the continuous loop of capturing raw data, processing it into actionable insights, and pruning the waste.
flowchart TD
%% Nodes
Inbox["00-inbox (Raw Bookmarks)"]
Curate["curate-bookmarks (Skill)"]
Draft["draft/ (Curated Notes)"]
Synthesize["synthesize-drafts (Skill)"]
Plan["synthesis/ (Strategic Plan)"]
Extract["plan-to-kanban (Skill)"]
Kanban["todo-kanban.md (Actionable Tasks)"]
Execute["Build / Execute"]
Prune["Delete Intermediates (Bookmarks & Drafts)"]
%% Flow
Inbox --> Curate
Curate -->|Extracts signal| Draft
Draft --> Synthesize
Synthesize -->|Scientific Thematic Synthesis| Plan
Plan --> Extract
Extract -->|Deduplicates and appends| Kanban
Kanban --> Execute
Execute --> Prune
classDef file fill:#2d2d2d,stroke:#555,stroke-width:1px,color:#fff
classDef skill fill:#1a3f5c,stroke:#4a90e2,stroke-width:2px,color:#fff
classDef action fill:#1e4620,stroke:#4caf50,stroke-width:2px,color:#fff
class Inbox,Draft,Plan,Kanban file
class Curate,Synthesize,Extract skill
class Execute,Prune action
Loading
Workflow Stages:
Capture : Raw material lands in 00-inbox/ .
Curate ( curate-bookmarks ) : The agent reads the inbox items, extracts the core value ("what we can steal"), and moves them into a draft/ folder within a specific Area. The source is logged in processed-sources.md .
Synthesize ( synthesize-drafts ) : The agent takes multiple drafts, analyzes them against each other using a scientific thematic matrix, and generates a unified Strategic Plan ( synthesis/ ).
Action ( plan-to-kanban ) : The agent reads the Strategic Plan, extracts the actionable tasks, deduplicates them, and appends them to the Area's todo-kanban.md .
Clean : Once the knowledge is durable and actionable, the spent intermediates (the original bookmark and the draft) are deleted.
4. Operational Skills (Toolbelt)
init-area : Interactively creates a new Area by challenging the idea, defining goals/scope, and scaffolding the required hub notes, Kanban board, and folders.
scout-idea : Validates new ideas, challenges their utility, and scouts for external resources/tools to build them.
curate-bookmarks : Processes inbox items into actionable Area drafts.
synthesize-drafts : Synthesizes multiple drafts in an Area into a strategic "Global Plan" using scientific thematic synthesis.
plan-to-kanban : Reads a synthesis document's action plan and extracts action items into the Area's Kanban board, deduplicating them.
vault-linter : Read-only knowledge-graph integrity check — broken [[wikilinks]] , orphaned notes, and missing source / captured_from traceability. Never edits.
audit-maintenance : Headlessly reviews pending maintenance tasks and peer-reviews tools created by other agents.
Review Loop : vault/99-system/maintenance/agent-kanban.md is a Kanban board with swimlanes Todo / In Progress / Done / Archived . Every tool creation must be logged as a new card under Todo .
Session Check : A SessionStart hook surfaces each agent's own pending reviews at session start by calling .claude/hooks/pending-reviews.sh <Reviewer> .
Quality Control : A review challenges and hardens the other agent's work — judging whether it produces the best-quality output, not just whether it follows format. A pass is earned; weak tools are failed with concrete, required improvements.
Source of Truth : Agents must read CLAUDE.md and this document before making structural changes.
No Direct Writes : Agents write to draft/ folders or specific system directories, never directly into the core of an Area without confirmation.
Traceability : Every agent-created note must include a source or captured_from field.
Security Guardrails : Agents stay inside the project directory and never read/write/exfiltrate credential or secret paths. Enforced per CLAUDE.md §10 .
To ensure the system remains robust and documented, every major change triggers this checklist:
Sync README.md : Document the new capability or structural shift.
Log agent-kanban.md : Add a card under Todo , assigned to the other agent.
Audit ARA : Confirm that no "projects" folders were created.
A self-governing Personal Knowledge Management (PKM) framework for Obsidian, featuring automated agent workflows for research, synthesis, and task extraction.
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
