---
source: "https://github.com/layr-hq/layr"
hn_url: "https://news.ycombinator.com/item?id=48586133"
title: "Layr – a modular UX and product constraint system for AI-built interfaces"
article_title: "GitHub - layr-hq/layr: A modular UX, design, and product optimisation system for turning AI-built interfaces into production-grade apps, turning proven principles into enforceable constraints that reduce friction, build trust, and drive action. Works with Claude, Codex, Cursor + all other agentic AI\n[truncated]"
author: "layrhq"
captured_at: "2026-06-18T16:15:37Z"
capture_tool: "hn-digest"
hn_id: 48586133
score: 8
comments: 0
posted_at: "2026-06-18T14:39:39Z"
tags:
  - hacker-news
  - translated
---

# Layr – a modular UX and product constraint system for AI-built interfaces

- HN: [48586133](https://news.ycombinator.com/item?id=48586133)
- Source: [github.com](https://github.com/layr-hq/layr)
- Score: 8
- Comments: 0
- Posted: 2026-06-18T14:39:39Z

## Translation

タイトル: Layr – AI 構築インターフェース用のモジュラー UX および製品制約システム
記事のタイトル: GitHub -layr-hq/layr: AI で構築されたインターフェイスを実稼働グレードのアプリに変えるためのモジュラー UX、デザイン、製品最適化システム。証明された原則を、摩擦を軽減し、信頼を構築し、アクションを促進する強制可能な制約に変えます。 Claude、Codex、Cursor、その他すべてのエージェント AI と連携
[切り捨てられた]
Description: A modular UX, design, and product optimisation system for turning AI-built interfaces into production-grade apps, turning proven principles into enforceable constraints that reduce friction, build trust, and drive action. Claude、Codex、Cursor、その他すべてのエージェント AI モデルで動作します。 - レイヤー-hq/レイヤー

記事本文:
GitHub -layr-hq/layr: AI で構築されたインターフェイスを実稼働グレードのアプリに変換し、実証済みの原則を摩擦を軽減し、信頼を構築し、アクションを促進する強制可能な制約に変えるためのモジュール式 UX、デザイン、および製品最適化システム。 Claude、Codex、Cursor、その他すべてのエージェント AI モデルで動作します。 · GitHub
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
別のアカウントに切り替えました

er tab or window.リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
レイヤー本部
/
層
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
104 Commits 104 Commits .cursor/ rules .cursor/ rules examples examples methods methods modules modules playbooks playbooks prompts prompts scorecards scorecards screens screens .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md ROADMAP.md ROADMAP.md RUN.md RUN.md SYSTEM.md SYSTEM.md layr.config.example.md layr.config.example.md scorecard.md scorecard.md View all files Repository files navigation
Layr - AI 構築アプリの本番システム
AI で構築されたインターフェイスを実稼働グレードのアプリに変えるためのモジュール型実稼働システム。 Layr turns proven UX, design, accessibility, security, performance, SEO, CRO, marketing, and copywriting principles into enforceable constraints that reduce friction, build trust, and drive action.
Claude、Cursor、およびその他のエージェント開発ツールと連携します。
一般的な AI 出力を、より明確で、より使いやすく、より本番環境に対応したインターフェイスに変換します。
オプション 1 から始めます。ツールが GitHub URL を読み取れない場合は、オプション 2 を使用してください。
ツールが GitHub URL を読み取ることができる場合は、これを使用します。
このプロンプトをコピーしてツールに貼り付け、タスク行を独自のリクエストに置き換えます。
AI で構築されたアプリの実稼働システムとして https://github.com/layr-hq/layr を使用します。
最初に SYSTEM.md を読み、次に RUN.md を読み、その後に従ってください。
Scope: Auto
深さ: 標準
タスク:
ユーザーがより迅速にプランを選択できるように、価格ページを改善します。
オプション 2 - プロジェクトにレイヤーを追加する
ツールがローカル ファイルで最適に動作する場合、または GitHub URL を確実に読み取ることができない場合に、これを使用します。
Downlo

ad するか、このリポジトリをプロジェクトのルートに Layer として複製します。
Copy this prompt and paste it into your tool.
Replace the task line with your own request.
Use ./layr/RUN.md for this task.
Read ./layr/SYSTEM.md first, then ./layr/RUN.md, then follow them.
範囲: 自動
深さ: 標準
タスク:
ユーザーがより迅速にプランを選択できるように、価格ページを改善します。
Option 3 - Add optional product context
これはオプションです。
製品とブランドのより強力な適合性が必要な場合に使用します。
layr/layr.config.example.md
コピーの名前を次のように変更します。
レイヤー/layer.config.md
知っていることだけを記入してください。残りは空白のままにしておきます。
Do not edit modules/ , methods/ , or RUN.md .
Option 4 - Add optional screen context
これはオプションです。
Use it for important screens where precision matters.
layr/screens/screen-template.md
Rename the copy to match the screen:
レイヤー/スクリーン/価格設定.md
layr/screens/onboarding.md
layr/screens/dashboard.md
結果に重大な影響を与えるフィールドのみを入力してください。
「プロジェクト管理アプリのダッシュボードを作成する」
→ Layer は関連するモジュールを選択し、運用ルールに照らして出力を改善します。
Layr は、AI で構築されたソフトウェアのためのルールベースの制作システムです。
これにより、エージェント ツールは、出力を出荷する前に、適切な製品モジュールを選択し、実証済みの方法を適用し、結果をスコアリングし、弱点を改善するための構造化された方法を提供します。
Cognitive Load - reduce thinking
Fitts’s Law - make actions easier
ヤコブの法則 - 使い慣れたパターンを使用する
Peak-End Rule - improve memorable moments
Gestalt - create clear structure
信号とノイズ - 乱雑なものを取り除く
Default Bias - guide decisions
Colour Theory - guide attention
タイポグラフィ - 読みやすさを向上させる
間隔のリズム - 構造を明確にする
アクセシビリティ - より多くの人がインターフェースを使用できるようにします
Security - catch unsafe patterns before shipping
Performance - keep the product fast and stable
分析

- 行動と製品学習を測定する
QA - catch edge cases and release risks
AI 製品 - AI 機能を信頼でき、制御可能にします
CRO - 摩擦を軽減し、アクションを向上させます。
SEO と AI 検索 - 検索エンジンと AI システムがコンテンツを理解しやすくします。
マーケティング - ポジショニングを強化し、コミュニケーションを重視する
コピーライティング - メッセージを明確、具体的、説得力のあるものにします
ほとんどのツールはこれらのアイデアを認識しています。レイヤーはそれらを動作可能にします。
AI は作業画面を迅速に生成できますが、作業が本番環境に対応していることと同じではありません。
システムがなければ、AI の出力には次のような結果が得られることがよくあります。
AI の推測ではなく、実際の運用標準に基づいて構築します。
フローチャート TD
A["Task or repo URL"] --> B{"Optional context?"}
B -- "セットアップなし" --> C["タスク + コードベースから推論"]
B -- "構成または画面の概要" --> D["layr.config.md /screens を読み取る"]
C --> E["スコープ + 深度コントロール"]
D --> E
E --> F["Selected module rules"]
F --> G["メソッドインデックス + 選択されたメソッド"]
G --> H["インターフェイスの構築または改善"]
H --> I[「証拠付きスコアカード」]
I --> J{"スコア >= 85?"}
J -- 「いいえ」 --> K[「弱点を改善する」]
K --> 私
J -- "はい" --> L["クリアで使用可能な出力"]
classDef input fill:#0f172a,stroke:#64748b,color:#f8fafc;
classDef rules fill:#111827,stroke:#818cf8,color:#f8fafc;
classDef アクションの塗りつぶし:#172554、ストローク:#60a5fa、色:#f8fafc;
classDef 決定フィル:#312e81、ストローク:#a5b4fc、カラー:#f8fafc;
classDef 出力の塗りつぶし:#064e3b、ストローク:#34d399、色:#ecfdf5;
クラス A、C、D 入力。
クラス E、F、G ルール。
クラス H、I、K アクション。
クラスB、J決定。
クラスL出力。
読み込み中
品質モード
モード
こんな方に最適
ユーザーが提供するもの
品質
ゼロセットアップ
Layr の試用、簡単な修正、レビュー
タスクのみ
強力な生産改善
おすすめ
Real product work
タスク + オプションのlayr.config.md
製品、ユーザー、ブランドの適合性の向上
画面レベル
価値の高い画面とフロー

s
タスク + 構成 + オプションの画面概要
最高の精度
Layr should never block on missing context unless the missing detail would materially change the product direction.
If context is missing, Layr should infer it and state assumptions briefly.
SYSTEM.md は、Layr の中心的なオペレーティング層です。
これは、システムがサーフェス タイプを選択し、プレイブックをロードし、ハード ゲートを適用し、メソッドの競合を解決し、推奨事項を証拠に結び付ける方法を定義します。
The kernel improves consistency across tasks by requiring every recommendation to connect to at least one of:
This keeps Layr science backed, evidence driven, and focused on production quality rather than taste.
Surface playbooks make Layr practical for common production work.
They define the required modules, recommended methods, hard gates, production rules, failure patterns, and evidence requirements for each major surface.
Surface scorecards in scorecards/ make scoring more specific and consistent for each surface.
Layr が UX とデザインからスタートしたのは、AI で構築された製品が最初に壊れるのが通常、そこにあるためです。不明確なフロー、弱い階層、乱雑な画面、機能するものの運用準備ができているとは感じられないインターフェースなどです。
Layr now works as a broader production layer for AI-built software.
Active modules cover UX, Design, Accessibility, Security, Performance, Analytics, QA, AI Product, CRO, SEO, Marketing, and Copywriting.
AI 検索は、タスクに AI 回答、検索可視性、GEO、ChatGPT 検索、Copilot 可視性、または回答エンジンの検出可能性が含まれる場合、SEO モジュールを通じてサポートされます。
The system is designed so new modules can be added without making Layr harder to use.
このタスクには Layer を使用します。
範囲: 自動
深さ: 標準
Layr chooses the relevant active modules automatically, so you don't have to understand the whole system.

計画されているモジュールの方向については、ROADMAP.md を参照してください。
このシステムを使用して、生成された製品作業を構築、レビュー、または改善して、より明確、より安全、より速く、よりアクセスしやすく、より測定可能で、より実稼働可能な状態にすることができます。
If your tool can read GitHub URLs, copy this prompt and paste it into your tool:
AI で構築されたアプリの実稼働システムとして https://github.com/layr-hq/layr を使用します。
最初に SYSTEM.md を読み、次に RUN.md を読み、その後に従ってください。
範囲: 自動
深さ: 標準
タスク:
ユーザーがより迅速にプランを選択できるように、価格ページを改善します。
If your tool cannot read GitHub URLs, download or clone this repo into your project root as layr .
Then copy this prompt and paste it into your tool:
Use ./layr/RUN.md for this task.
Read ./layr/SYSTEM.md first, then ./layr/RUN.md, then follow them.
範囲: 自動
深さ: 標準
タスク:
ユーザーがより迅速にプランを選択できるように、価格ページを改善します。
If your tool cannot read GitHub URLs reliably, use the local folder option.
Replace the example task with what you want the AI to build, fix, review, or improve.
ユーザーがより迅速にプランを選択できるように、価格ページを改善します。
より良い:
Improve the pricing page for early-stage SaaS founders.
主なアクションは、無料トライアルを開始することです。
Preserve the existing design system.
ステップ 3 - 有用な場合にのみコンテキストを追加する
For better product fit, copy layr.config.example.md , rename the copy to layr.config.md , and fill only what you know.
製品名：
Primary user:
コアユーザーの目標:
製品の主なアクション:
Design source:
これはオプションです。 Layer はそれなしでも動作します。
Step 4 - Add screen briefs only for important screens
For important screens, copy layr/screens/screen-template.md , rename the copy, and fill only what matters.
最小限の有用な画面コンテキスト:
スクリーン名:
ユーザーの意図:
主な目標:
主なアクション:
これはオプションです。レイヤーさん、私がいいですか？

タスクとコードベースから画面コンテキストが欠落していることを発見します。
ステップ 5 - Layer を構築して洗練させます
関連する Surface Playbook をロードします
安全な場合に欠落しているコンテキストを推測する
コンテキストが本当に妨げられている場合にのみ、最大 3 つの質問をする
製品表面を構築、レビュー、または改善する
証拠をもとに結果を採点する
出力スコアが少なくとも 85 点になるまで繰り返します
ファイル
目的
User edits?
システム.md
サーフェス選択、ハードゲート、競合ルール、および証拠基準のための中央操作層
いいえ
RUN.md
エージェント ツールの主な実行エントリ ポイント
いいえ
プレイブック/index.md
Surface プレイブック ルーティング
いいえ
プレイブック/*.md
一般的な製品サーフェスのためのプロダクション プレイブック
いいえ
モジュール/index.md
アクティブおよび計画されたモジュール ルーティング
いいえ
モジュール/ux.md
UX の動作、ルール、スコアリング、検証
いいえ
modules/design.md
レイアウト、階層、間隔、視覚的な明瞭さ
いいえ
モジュール/アクセシビリティ.md
アクセシビリティのルールと検証
いいえ
モジュール/セキュリティ.md
セキュリティルールとリスクチェック
いいえ
モジュール/パフォーマンス.md
パフォーマンスルールとインタラクションチェック
いいえ
モジュール/analytics.md
測定とイベント追跡のルール
いいえ
モジュール/qa.md
QA, edge case, and release readiness rules
いいえ
モジュール/ai.md
AI 製品の動作と信頼ルール
いいえ
モジュール/cro.md
変換と摩擦

[切り捨てられた]

## Original Extract

A modular UX, design, and product optimisation system for turning AI-built interfaces into production-grade apps, turning proven principles into enforceable constraints that reduce friction, build trust, and drive action. Works with Claude, Codex, Cursor + all other agentic AI models. - layr-hq/layr

GitHub - layr-hq/layr: A modular UX, design, and product optimisation system for turning AI-built interfaces into production-grade apps, turning proven principles into enforceable constraints that reduce friction, build trust, and drive action. Works with Claude, Codex, Cursor + all other agentic AI models. · GitHub
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
layr-hq
/
layr
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
104 Commits 104 Commits .cursor/ rules .cursor/ rules examples examples methods methods modules modules playbooks playbooks prompts prompts scorecards scorecards screens screens .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md ROADMAP.md ROADMAP.md RUN.md RUN.md SYSTEM.md SYSTEM.md layr.config.example.md layr.config.example.md scorecard.md scorecard.md View all files Repository files navigation
Layr - Production System for AI-Built Apps
A modular production system for turning AI-built interfaces into production-grade apps. Layr turns proven UX, design, accessibility, security, performance, SEO, CRO, marketing, and copywriting principles into enforceable constraints that reduce friction, build trust, and drive action.
Works with Claude, Cursor, and other agentic development tools.
Turn typical AI output into a clearer, more usable, more production-ready interface.
Start with Option 1. If your tool cannot read GitHub URLs, use Option 2.
Use this when your tool can read GitHub URLs.
Copy this prompt, paste it into your tool, then replace the task line with your own request:
Use https://github.com/layr-hq/layr as the production system for AI-built apps.
Read SYSTEM.md first, then RUN.md, then follow them.
Scope: Auto
Depth: Standard
Task:
Improve the pricing page so users can choose a plan faster.
Option 2 - Add Layr to your project
Use this when your tool works best with local files or cannot reliably read GitHub URLs.
Download or clone this repo into your project root as layr .
Copy this prompt and paste it into your tool.
Replace the task line with your own request.
Use ./layr/RUN.md for this task.
Read ./layr/SYSTEM.md first, then ./layr/RUN.md, then follow them.
Scope: Auto
Depth: Standard
Task:
Improve the pricing page so users can choose a plan faster.
Option 3 - Add optional product context
This is optional.
Use it when you want stronger product and brand fit.
layr/layr.config.example.md
Rename the copy to:
layr/layr.config.md
Fill only what you know. Leave the rest blank.
Do not edit modules/ , methods/ , or RUN.md .
Option 4 - Add optional screen context
This is optional.
Use it for important screens where precision matters.
layr/screens/screen-template.md
Rename the copy to match the screen:
layr/screens/pricing.md
layr/screens/onboarding.md
layr/screens/dashboard.md
Fill only the fields that materially affect the result.
“Create a dashboard for a project management app”
→ Layr selects the relevant modules and improves the output against production rules.
Layr is a rule-based production system for AI-built software.
It gives agentic tools a structured way to select the right product modules, apply proven methods, score the result, and improve weak areas before the output ships.
Cognitive Load - reduce thinking
Fitts’s Law - make actions easier
Jakob’s Law - use familiar patterns
Peak-End Rule - improve memorable moments
Gestalt - create clear structure
Signal vs Noise - remove clutter
Default Bias - guide decisions
Colour Theory - guide attention
Typography - improve readability
Spacing Rhythm - clarify structure
Accessibility - make interfaces usable for more people
Security - catch unsafe patterns before shipping
Performance - keep the product fast and stable
Analytics - measure behaviour and product learning
QA - catch edge cases and release risks
AI Product - make AI features trustworthy and controllable
CRO - reduce friction and increase action
SEO and AI Search - make content easier for search engines and AI systems to understand
Marketing - sharpen positioning and value communication
Copywriting - make messages clear, specific, and persuasive
Most tools know these ideas. Layr makes them operational.
AI can generate working screens quickly, but working is not the same as production-ready.
Without a system, AI output often has:
Build with real production standards, not AI guesses.
flowchart TD
A["Task or repo URL"] --> B{"Optional context?"}
B -- "No setup" --> C["Infer from task + codebase"]
B -- "Config or screen brief" --> D["Read layr.config.md / screens"]
C --> E["Scope + depth control"]
D --> E
E --> F["Selected module rules"]
F --> G["Methods index + selected methods"]
G --> H["Build or improve interface"]
H --> I["Scorecard with evidence"]
I --> J{"Score >= 85?"}
J -- "No" --> K["Improve weak areas"]
K --> I
J -- "Yes" --> L["Clear, usable output"]
classDef input fill:#0f172a,stroke:#64748b,color:#f8fafc;
classDef rules fill:#111827,stroke:#818cf8,color:#f8fafc;
classDef action fill:#172554,stroke:#60a5fa,color:#f8fafc;
classDef decision fill:#312e81,stroke:#a5b4fc,color:#f8fafc;
classDef output fill:#064e3b,stroke:#34d399,color:#ecfdf5;
class A,C,D input;
class E,F,G rules;
class H,I,K action;
class B,J decision;
class L output;
Loading
Quality modes
Mode
Best for
What the user provides
Quality
Zero setup
Trying Layr quickly, simple fixes, reviews
Task only
Strong production improvement
Recommended
Real product work
Task + optional layr.config.md
Better product, user, and brand fit
Screen-level
High-value screens and flows
Task + config + optional screen brief
Highest precision
Layr should never block on missing context unless the missing detail would materially change the product direction.
If context is missing, Layr should infer it and state assumptions briefly.
SYSTEM.md is Layr's central operating layer.
It defines how the system selects surface types, loads playbooks, applies hard gates, resolves method conflicts, and keeps recommendations tied to evidence.
The kernel improves consistency across tasks by requiring every recommendation to connect to at least one of:
This keeps Layr science backed, evidence driven, and focused on production quality rather than taste.
Surface playbooks make Layr practical for common production work.
They define the required modules, recommended methods, hard gates, production rules, failure patterns, and evidence requirements for each major surface.
Surface scorecards in scorecards/ make scoring more specific and consistent for each surface.
Layr started with UX and Design because that is where AI-built products usually break first: unclear flows, weak hierarchy, messy screens, and interfaces that work but do not feel production-ready.
Layr now works as a broader production layer for AI-built software.
Active modules cover UX, Design, Accessibility, Security, Performance, Analytics, QA, AI Product, CRO, SEO, Marketing, and Copywriting.
AI Search is supported through the SEO module when the task involves AI answers, retrieval visibility, GEO, ChatGPT search, Copilot visibility, or answer-engine discoverability.
The system is designed so new modules can be added without making Layr harder to use.
Use Layr for this task.
Scope: Auto
Depth: Standard
Layr chooses the relevant active modules automatically, so you don't have to understand the whole system.
See ROADMAP.md for the planned module direction.
Use this system to build, review, or improve generated product work until it is clearer, safer, faster, more accessible, more measurable, and more production-ready.
If your tool can read GitHub URLs, copy this prompt and paste it into your tool:
Use https://github.com/layr-hq/layr as the production system for AI-built apps.
Read SYSTEM.md first, then RUN.md, then follow them.
Scope: Auto
Depth: Standard
Task:
Improve the pricing page so users can choose a plan faster.
If your tool cannot read GitHub URLs, download or clone this repo into your project root as layr .
Then copy this prompt and paste it into your tool:
Use ./layr/RUN.md for this task.
Read ./layr/SYSTEM.md first, then ./layr/RUN.md, then follow them.
Scope: Auto
Depth: Standard
Task:
Improve the pricing page so users can choose a plan faster.
If your tool cannot read GitHub URLs reliably, use the local folder option.
Replace the example task with what you want the AI to build, fix, review, or improve.
Improve the pricing page so users can choose a plan faster.
Better:
Improve the pricing page for early-stage SaaS founders.
The primary action is starting a free trial.
Preserve the existing design system.
Step 3 - Add context only when useful
For better product fit, copy layr.config.example.md , rename the copy to layr.config.md , and fill only what you know.
Product name:
Primary user:
Core user goal:
Primary product action:
Design source:
This is optional. Layr still works without it.
Step 4 - Add screen briefs only for important screens
For important screens, copy layr/screens/screen-template.md , rename the copy, and fill only what matters.
Minimum useful screen context:
Screen name:
User intent:
Primary goal:
Primary action:
This is optional. Layr should infer missing screen context from the task and codebase.
Step 5 - Let Layr build and refine
load the relevant surface playbook
infer missing context when safe
ask at most 3 questions only when context is truly blocking
build, review, or improve the product surface
score the result with evidence
repeat until the output scores at least 85
File
Purpose
User edits?
SYSTEM.md
Central operating layer for surface selection, hard gates, conflict rules, and evidence standards
No
RUN.md
Main execution entry point for agentic tools
No
playbooks/index.md
Surface playbook routing
No
playbooks/*.md
Production playbooks for common product surfaces
No
modules/index.md
Active and planned module routing
No
modules/ux.md
UX behaviour, rules, scoring, and validation
No
modules/design.md
Layout, hierarchy, spacing, and visual clarity
No
modules/accessibility.md
Accessibility rules and validation
No
modules/security.md
Security rules and risk checks
No
modules/performance.md
Performance rules and interaction checks
No
modules/analytics.md
Measurement and event tracking rules
No
modules/qa.md
QA, edge case, and release readiness rules
No
modules/ai.md
AI product behaviour and trust rules
No
modules/cro.md
Conversion and friction

[truncated]
