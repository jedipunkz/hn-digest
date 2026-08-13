---
source: "https://github.com/leo-proger/claude-first-customer-finder-skill"
hn_url: "https://news.ycombinator.com/item?id=49291974"
title: "Claude First Customer Finder"
article_title: "GitHub - leo-proger/claude-first-customer-finder-skill: A Claude skill that finds evidence-backed potential first customers from recent public signals. · GitHub"
author: "leo_proger"
captured_at: "2026-08-13T21:35:10Z"
capture_tool: "hn-digest"
hn_id: 49291974
score: 1
comments: 0
posted_at: "2026-08-13T21:15:51Z"
tags:
  - hacker-news
  - translated
---

# Claude First Customer Finder

- HN: [49291974](https://news.ycombinator.com/item?id=49291974)
- Source: [github.com](https://github.com/leo-proger/claude-first-customer-finder-skill)
- Score: 1
- Comments: 0
- Posted: 2026-08-13T21:15:51Z

## Translation

タイトル: クロードファーストカスタマーファインダー
記事のタイトル: GitHub - leo-proger/claude-first-customer-finder-skill: 最近の公開シグナルから証拠に裏付けられた潜在的な最初の顧客を見つけるクロード スキル。 · GitHub
説明: 最近の公開シグナルから証拠に裏付けられた潜在的な最初の顧客を見つけるクロード スキル。 - leo-proger/claude-first-customer-finder-skill

記事本文:
GitHub - leo-proger/claude-first-customer-finder-skill: 最近の公開シグナルから証拠に裏付けられた潜在的な最初の顧客を見つけるクロード スキル。 · GitHub
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
レオプロガー
/
クロードファーストカスタマーファインダースキル
公共
Kappaemme-git/codex-first-customer-finder-skill からフォーク
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット 最初の顧客発見者 最初の顧客

-finder スクリプト スクリプト .gitignore .gitignore ライセンス ライセンス README.md README.md package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロードファースト顧客検索スキル
最近の世間の痛み、需要、タイミングのシグナルを使用して、スタートアップ URL または製品アイデアを潜在的な最初の顧客の適格な最終候補リストに変えるクロード コード スキル。
理想的な顧客プロファイルを定義し、Claude Code の Web ツールを使用して公開ソースを調査し、あらゆる見込み客の背後にある証拠をリンクし、適合性とタイミングをランク付けし、ソースベースのオープナーを作成し、洗練された HTML レポートを作成します。アウトリーチが自動的に送信されることはありません。
起動 URL、リポジトリ、または製品の説明を分析します。
主要な顧客プロファイルとそれに隣接する理想的な顧客プロファイルを定義します
明示的な要求、問題、回避策、切り替え、およびタイミング信号を見つけます
証拠に基づいたスコアで見込み顧客を認定します
すべての主な見込み顧客を元の公開ソースにリンクします
敬意を払い、情報源に基づいたアウトリーチのオープナーを草案する
応答性の高いスタンドアロン HTML レポートを作成します
デフォルトですべてのアウトリーチを手動のままにする
プライベートな連絡先の充実や機密性の高い個人データを回避します
npx --yes クロードファーストカスタマーファインダースキル@latest
これにより、スキルが以下にインストールされます。
~/.claude/skills/first-customer-finder
個人のスキル ディレクトリではなく、単一のリポジトリの場合は ./.claude/skills/first-customer-finder にインストールする場合は --project を渡し、カスタムの場所の場合は --skills-dir <path> を渡します。
インストール後に Claude Code を再起動して、新しいスキルを取得します。
スラッシュ コマンドを使用して明示的に呼び出すか、タスクを説明するだけです。クロード コードはその説明からスキルを取得します。
最初の 10 人の潜在顧客を見つけます。
/first-customer-finder は、証拠に裏付けられた https://example.com の潜在的な最初の顧客 10 社を検索し、最終的な HTML レポートを作成します。
デザインパーを探す

トナー:
このスタートアップのデザイン パートナー モードの /first-customer-finder: [URL]。問題を公に説明しており、製品に関するフィードバックを提供する可能性が高い人を優先します。
B2B 調査:
[URL] の b2b モードの /first-customer-finder。公開ビジネスのトリガーを見つけ、関連する企業を特定し、何も送信せずに見込み顧客ごとに 1 つのオープナーを作成します。
出力
証拠に裏付けられた見込み客候補リスト
7 日間の手動アウトリーチ プラン
見込み客は公開シグナルに基づいた仮説であり、確認された顧客や保証された購入者ではありません。
HTML レポートはワークスペースの Outputs/ ディレクトリに保存され、クリック可能な絶対ファイル パスとして返されます。ローカル ファイルの代わりに共有可能なリンクが必要な場合は、クロードにそれをアーティファクトとして公開するように依頼してください。
クイック: 最大 5 つの有力な見込み客
標準 : 複数のソースタイプにわたって最大 10 件のプロスペクト
deep : 最大 20 の見込み客と反復パターン分析
design-partners : フィードバック指向の早期採用者
b2b : 企業および公的ビジネスのトリガー
コミュニティ: 明示的なリクエストと公開ディスカッションのシグナル
git clone https://github.com/leo-proger/claude-first-customer-finder-skill.git
mkdir -p ~ /.claude/skills
cp -R クロード-ファーストカスタマーファインダースキル/ファーストカスタマーファインダー ~ /.claude/skills/ファーストカスタマーファインダー
インストール後、Claude Code を再起動します。
最近の公開シグナルから証拠に裏付けられた潜在的な最初の顧客を見つけるクロード スキル。
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A Claude skill that finds evidence-backed potential first customers from recent public signals. - leo-proger/claude-first-customer-finder-skill

GitHub - leo-proger/claude-first-customer-finder-skill: A Claude skill that finds evidence-backed potential first customers from recent public signals. · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
leo-proger
/
claude-first-customer-finder-skill
Public
forked from Kappaemme-git/codex-first-customer-finder-skill
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits first-customer-finder first-customer-finder scripts scripts .gitignore .gitignore LICENSE LICENSE README.md README.md package.json package.json View all files Repository files navigation
Claude First Customer Finder Skill
A Claude Code skill that turns a startup URL or product idea into a qualified shortlist of potential first customers using recent public pain, demand, and timing signals.
It defines the ideal customer profile, researches public sources with Claude Code's web tools, links the evidence behind every prospect, ranks fit and timing, drafts a source-based opener, and creates a polished HTML report. It never sends outreach automatically.
Analyzes a startup URL, repository, or product description
Defines the primary and adjacent ideal customer profiles
Finds explicit demand, pain, workaround, switching, and timing signals
Qualifies prospects with an evidence-based score
Links every primary prospect to the original public source
Drafts respectful, source-based outreach openers
Creates a responsive standalone HTML report
Keeps all outreach manual by default
Avoids private contact enrichment and sensitive personal data
npx --yes claude-first-customer-finder-skill@latest
This installs the skill into:
~/.claude/skills/first-customer-finder
Pass --project to install into ./.claude/skills/first-customer-finder for a single repository instead of your personal skills directory, or --skills-dir <path> for a custom location.
Restart Claude Code after installation so it picks up the new skill.
Invoke it explicitly with the slash command, or just describe the task — Claude Code will pick up the skill from its description.
Find the first ten potential customers:
/first-customer-finder find ten evidence-backed potential first customers for https://example.com and create the final HTML report.
Find design partners:
/first-customer-finder in design-partners mode for this startup: [URL]. Prioritize people publicly describing the problem and likely to give product feedback.
B2B research:
/first-customer-finder in b2b mode for [URL]. Find public business triggers, qualify the relevant companies, and draft one opener per prospect without sending anything.
Output
Evidence-backed prospect shortlist
Seven-day manual outreach plan
Prospects are hypotheses based on public signals, not confirmed customers or guaranteed buyers.
The HTML report is saved to the workspace outputs/ directory and returned as a clickable absolute file path. If you want a shareable link instead of a local file, ask Claude to publish it as an Artifact.
quick : up to five strong prospects
standard : up to ten prospects across several source types
deep : up to twenty prospects and repeated-pattern analysis
design-partners : feedback-oriented early adopters
b2b : companies and public business triggers
community : explicit requests and public discussion signals
git clone https://github.com/leo-proger/claude-first-customer-finder-skill.git
mkdir -p ~ /.claude/skills
cp -R claude-first-customer-finder-skill/first-customer-finder ~ /.claude/skills/first-customer-finder
Restart Claude Code after installation.
A Claude skill that finds evidence-backed potential first customers from recent public signals.
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
