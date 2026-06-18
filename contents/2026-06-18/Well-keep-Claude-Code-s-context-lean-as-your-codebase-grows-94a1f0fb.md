---
source: "https://github.com/BanyanOutcomes/well-actually"
hn_url: "https://news.ycombinator.com/item?id=48588868"
title: "Well-– keep Claude Code's context lean as your codebase grows"
article_title: "GitHub - BanyanOutcomes/well-actually: A starter layout for keeping Claude Code's context lean as your codebase grows: path-scoped rules that auto-load by file, on-demand system docs, and skills to maintain both. · GitHub"
author: "nich2533"
captured_at: "2026-06-18T18:54:43Z"
capture_tool: "hn-digest"
hn_id: 48588868
score: 1
comments: 0
posted_at: "2026-06-18T17:46:52Z"
tags:
  - hacker-news
  - translated
---

# Well-– keep Claude Code's context lean as your codebase grows

- HN: [48588868](https://news.ycombinator.com/item?id=48588868)
- Source: [github.com](https://github.com/BanyanOutcomes/well-actually)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T17:46:52Z

## Translation

タイトル: そう – コードベースの成長に合わせてクロード コードのコンテキストを無駄にしないようにしましょう
記事のタイトル: GitHub - BanyanOutcomes/well-actually: コードベースの成長に合わせてクロード コードのコンテキストを無駄なく保つためのスターター レイアウト: ファイルごとに自動ロードされるパス スコープのルール、オンデマンドのシステム ドキュメント、および両方を維持するスキル。 · GitHub
説明: コードベースの成長に合わせてクロード コードのコンテキストを無駄なく保つためのスターター レイアウト: ファイルごとに自動ロードされるパス スコープのルール、オンデマンドのシステム ドキュメント、および両方を維持するスキル。 - BanyanOutcomes/まあ、実際のところ

記事本文:
GitHub - BanyanOutcomes/well-actually: コードベースの成長に合わせてクロード コードのコンテキストを無駄なく保つためのスターター レイアウト: ファイルごとに自動ロードされるパス スコープのルール、オンデマンドのシステム ドキュメント、および両方を維持するスキル。 · GitHub
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
ガジュマルの成果
/
w

そうですね、実は
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット .claude .claude ドキュメント ドキュメントの例 例 src src .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
コードベースの成長に合わせてクロード コードを鮮明に保つためのスターター レイアウト。
前提を 1 行でまとめると、コンテキスト ウィンドウは予算であり、バックパックではありません。積み込むものはすべて 2 倍のコストがかかります。1 回は空間に、もう 1 回は注意を払う必要があります。したがって、このテンプレートは、すべてを CLAUDE.md に詰め込むのではなく、コンテキストを、関連する場合にのみ読み込まれる部分に分割します。ルールの半分については、Claude Code 独自のパススコープの読み込みを使用するため、「関連する場合のみ」は希望ではなくメカニズムです。
全文: コンテキストは予算であり、バックパックではありません
。
§── CLAUDE.md # リーンマップ: 方向 + ドキュメント/へのポインタ、それ以上のものはありません
§── .claude/
│ §── rules/ # やり方 — 各ファイルはパスが一致する場合にのみ自動ロードされます。
│ │ §──database.md # パス: migrations, *.sql, src/db/**
│ │ §── api-design.md # パス: src/api/**、ルート、ハンドラー
│ │ §── testing.md # パス: *.test.*、*.spec.*、tests/**
│ │ └─ security.md # パス: auth、api、webhooks、ミドルウェア、env
│ └── スキル/
│ §── Well-Actually-get-started/ # シードルール/ + ドキュメント/ を 1 回のパスで実行
│ §── Well-actually-create-rules/ # ファットな CLAUDE.md を .claude/rules/ に分割します
│ §──well-actually-documentation-full/ # リポジトリ全体を最初から文書化します
│ └── Well-Actually-documentation-recent/# 最後の同期以降のコミットとドキュメントを同期します
§─

─ document/ # システムの仕組み — CLAUDE.md が指す、オンデマンドで読む
│ §── auth-flow.md
│ §── ingestion-pipeline.md
│ §── billing-webhooks.md
│ └── .last-synced # ドキュメントが最後に同期された SHA をコミットします (最初のフルパスで作成されました)
━── 例/
└── walkthrough.md # 1 つの小さな機能、最初から最後まで — 動作中のメカニズムを確認する
読み込みの実際の仕組み
これは、「CLAUDE.md を分割する」というアドバイスで最も間違っている部分なので、正確にする価値があります (公式のメモリドキュメントと照らし合わせて確認済み)。
CLAUDE.md はセッションごとに完全にロードされます。したがって、これは百科事典ではなく目次であり、方向性と document/ へのポインタ、および普遍的な規格なので常にコンテキスト内に存在する必要があります。
.claude/rules/*.md は負荷に耐えるトリックです。各ルール ファイルには YAML フロントマターのパス glob が含まれており、クロード コードは、クロードがその glob に一致するファイルに触れた場合にのみ、パスを自動的に読み込みます。移行を開く→database.mdが到着します。タイプミスを修正する → 何も届きません。これらを CLAUDE.md から参照するのではありません。グロブが仕事をします。
⚠️ パスのないルール ファイル: すべてのセッションでロードをブロックします — これもバックパックです。ここのすべてのファイルには意図的に 1 つあります。
これは、一般的な代替手段よりも優れています。ファイルのマークダウンによる言及は自動的には何も読み込まれません (モデルはそれを読み取ることを選択する必要があります)。また、@import は起動時に熱心に読み込まれます (セッションごとにコンテキストが必要です)。 paths: は、自動かつ関連性ゲートの両方を備えた唯一のメカニズムです。
document/ はオンデマンドで読み取られます。自動ロードはありません。これは、時々参照する参考資料としては適切です。 CLAUDE.md にはこれらがリストされているため、適切なものを簡単に見つけることができます。優れたドキュメントには理解がキャッシュされます。 .last-synced マーカーは、メンテナンス スキルがそれを維持する方法です。

正直なキャッシュ。
具体的な前後については、example/walkthrough.md を参照してください。パススコープのルールがタスクの途中でロードされ、その後、ドキュメントが diff から再同期されます。
レイアウトが面倒な作業にならないようにするには、次のようなスキルが必要です。
Well-Actually-get-started は両方のツリーを一度にシードします。 CLAUDE.md を .claude/rules/ に分割し、コードを document/ に読み込み、サブエージェントをファンアウトして高速に実行します。
Well-actually-create-rules はルールの半分だけを実行します。既存の CLAUDE.md を読み取り、それを懸念スコープの .claude/rules/ ファイルに分解します。それぞれのファイルには正しいパス (glob) が指定されます。
Well-actually-documentation-full はリポジトリ全体を読み取り、documentation/ を最初から書き込みます (サブシステムごとに 1 つのドキュメントを並行して書き込みます)。シードを追加したり、腐りすぎたドキュメントをリセットしたりするために使用します。また、最初の .last-synced マーカーも書き込みます。
Well-actually-documentation-recent はメンテナンス ループです。最後の同期 (カレンダー ウィンドウではなく .last-synced によって追跡) 以降のすべてのコミットを読み取り、影響を受けるドキュメントのみを更新してから、マーカーを進めます。
このリポジトリのクローンを作成するか、レイアウトを既存のプロジェクトにコピーします。
Claude Code でフォルダーを開きます。
Well-Actually-get-started を実行します。既存の CLAUDE.md とコードを読み取り、.claude/rules/ と document/ をシードし、プレースホルダーを実際のものに置き換えます。 (CLAUDE.md を使用せずに最初から始めますか? これをスキップして、プレースホルダー ファイルを手動で入力します。各ファイルには、そこに属するものを説明する短いプロンプトが同梱されており、ルール ファイルにはコピーするパス: フォーマットが表示されます。)
提案されたリーン CLAUDE.md を確認し、承認します。
それ以降、各機能が導入された後、well-actually-documentation-recent を実行し、影響を受けるドキュメントが自動的に更新されるのを確認します。ドキュメントが行き過ぎた場合、well-actually-documentation-full はリポジトリ全体からドキュメントを再構築します。
eにインストールする

既存のプロジェクト
クイックスタートは、新たに開始することを前提としています。これをすでにお持ちのアプリに追加するには、スキルのみが必要です。プレースホルダー CLAUDE.md 、 .claude/rules/ 、documentation/ 、 src/ 、および example/ は新しいプロジェクトの足場であり、スキルはコードから実際のバージョンを再生成します。スキルだけをコピーします。
# アプリのルートから
mkdir -p .claude/スキル
cp -r /path/to/well-actually/.claude/skills/well-actually- * .claude/skills/
#パワーシェル
新しい項目 - 項目タイプ ディレクトリ - .claude\skills を強制する |アウトヌル
アイテムのコピー - 再帰 C:\path\to\well - 実は\.claude\skills\well - 実は -* .claude\skills\
Well-Actually-* の名前は一般的なプロジェクトのスキルと衝突しないため、これはきれいにマージされます。次に、アプリから次のようにします。
Well-actually-create-rules を実行して、既存の CLAUDE.md を .claude/rules/ に分解します。これはスリム化された CLAUDE.md (差分用に CLAUDE.md.projected に書き込まれます) を提案しており、承認なしにオリジナルを上書きすることはありません。
Well-actually-documentation-full を実行してコードから document/ を生成し、最初の .last-synced マーカーを書き込みます。
出荷時には最新の実際の文書を維持してください。
これらのフォルダーが存在しない場合、スキルは .claude/rules/ と document/ を作成するため、手動で設定する必要はありません。 (well-actually-get-started はステップ 1 ～ 2 を一緒に実行します。既存のアプリでは、ステップ 1 ～ 2 を直接実行することもできます。)
これらのファイルは意図的に汎用的なものであり、処方箋ではなくプロンプトです。不要なルールを削除し、懸念事項を追加して、各ルールのパスを調整します。つまり、実際のレイアウトにグロブします。これは、ルールがいつ実行されるかを決定するため、正しく行う価値のあることの 1 つです。すべてのファイルを短くしてください。ルール ファイルが 2 つ目の百科事典に変わった瞬間、バックパックを背負う生活に戻ります。
保存用のスターター レイアウト

Claude Code のコンテキストはコードベースの成長に合わせて強化されます。ファイルごとに自動ロードされるパススコープのルール、オンデマンドのシステム ドキュメント、およびその両方を維持するスキルです。
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

A starter layout for keeping Claude Code's context lean as your codebase grows: path-scoped rules that auto-load by file, on-demand system docs, and skills to maintain both. - BanyanOutcomes/well-actually

GitHub - BanyanOutcomes/well-actually: A starter layout for keeping Claude Code's context lean as your codebase grows: path-scoped rules that auto-load by file, on-demand system docs, and skills to maintain both. · GitHub
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
BanyanOutcomes
/
well-actually
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits .claude .claude documentation documentation example example src src .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md View all files Repository files navigation
A starter layout for keeping Claude Code sharp as your codebase grows.
The premise, in one line: the context window is a budget, not a backpack. Everything you load costs you twice — once in space, once in attention. So instead of cramming everything into CLAUDE.md , this template splits context into pieces that load only when they're relevant — and for the rules half, it uses Claude Code's own path-scoped loading so "only when relevant" is a mechanism , not a hope.
Full write-up: Context Is a Budget, Not a Backpack
.
├── CLAUDE.md # lean map: orientation + pointers to documentation/, nothing more
├── .claude/
│ ├── rules/ # how we do things — each file auto-loads ONLY when its paths: match
│ │ ├── database.md # paths: migrations, *.sql, src/db/**
│ │ ├── api-design.md # paths: src/api/**, routes, handlers
│ │ ├── testing.md # paths: *.test.*, *.spec.*, tests/**
│ │ └── security.md # paths: auth, api, webhooks, middleware, env
│ └── skills/
│ ├── well-actually-get-started/ # seeds rules/ + documentation/ in one pass
│ ├── well-actually-create-rules/ # splits a fat CLAUDE.md into .claude/rules/
│ ├── well-actually-documentation-full/ # documents the whole repo from scratch
│ └── well-actually-documentation-recent/# syncs docs with commits since the last sync
├── documentation/ # how the system works — read on demand, pointed at by CLAUDE.md
│ ├── auth-flow.md
│ ├── ingestion-pipeline.md
│ ├── billing-webhooks.md
│ └── .last-synced # commit SHA the docs were last synced to (created on first full pass)
└── example/
└── walkthrough.md # one small feature, start to finish — see the mechanism in action
How loading actually works
This is the part most "split your CLAUDE.md" advice gets wrong, so it's worth being precise (confirmed against the official memory docs ):
CLAUDE.md loads in full, every session . So it stays a table of contents, not an encyclopedia — orientation plus pointers to documentation/ , and any standard so universal it should always be in context.
.claude/rules/*.md is the load-bearing trick. Each rule file carries a paths: glob in its YAML frontmatter, and Claude Code loads it automatically and only when Claude touches a file matching that glob . Open a migration → database.md arrives. Fix a typo → nothing arrives. You don't reference these from CLAUDE.md ; the glob does the work.
⚠️ A rule file without a paths: block loads on every session — that's the backpack again. Every file here has one on purpose.
This beats the common alternatives: a markdown mention of a file loads nothing automatically (the model has to choose to read it), and an @import loads eagerly at startup (costs context every session). paths: is the only mechanism that's both automatic and relevance-gated.
documentation/ is read on demand — there's no auto-load, which is correct for reference material you reach for occasionally. CLAUDE.md lists these so the right one is easy to find. Good docs are cached understanding; the .last-synced marker is how the maintenance skill keeps that cache honest.
See example/walkthrough.md for a concrete before/after: a path-scoped rule loading itself mid-task, and documentation/ re-syncing from a diff afterward.
The skills are what keep the layout from being a chore:
well-actually-get-started seeds both trees at once — splits your CLAUDE.md into .claude/rules/ and reads your code into documentation/ , fanning out subagents so it's fast.
well-actually-create-rules does just the rules half: reads an existing CLAUDE.md and decomposes it into concern-scoped .claude/rules/ files, each with the right paths: glob.
well-actually-documentation-full reads the whole repo and writes documentation/ from scratch — one doc per subsystem, in parallel. Use it to seed, or to reset docs that have rotted too far. It also writes the first .last-synced marker.
well-actually-documentation-recent is the maintenance loop: it reads every commit since the last sync (tracked by .last-synced , not a calendar window) and updates only the affected docs, then advances the marker.
Clone this repo, or copy the layout into an existing project.
Open the folder in Claude Code.
Run well-actually-get-started . It reads your existing CLAUDE.md and code and seeds .claude/rules/ and documentation/ for you — replacing the placeholders with the real thing. (Starting from scratch with no CLAUDE.md ? Skip this and fill the placeholder files in by hand; each ships with a short prompt explaining what belongs there, and the rule files show the paths: format to copy.)
Review the proposed lean CLAUDE.md and approve it.
From then on, after each feature lands, run well-actually-documentation-recent and watch the affected docs update themselves. If the docs ever drift too far, well-actually-documentation-full rebuilds them from the whole repo.
Install into an existing project
The Quickstart assumes you're starting fresh. To add this to an app you already have , you only need the skills — the placeholder CLAUDE.md , .claude/rules/ , documentation/ , src/ , and example/ are scaffolding for a new project, and the skills regenerate the real versions from your code. Copy just the skills:
# from your app's root
mkdir -p .claude/skills
cp -r /path/to/well-actually/.claude/skills/well-actually- * .claude/skills/
# PowerShell
New-Item - ItemType Directory - Force .claude\skills | Out-Null
Copy-Item - Recurse C:\path\to\well - actually\.claude\skills\well - actually -* .claude\skills\
The well-actually-* names won't collide with typical project skills, so this merges cleanly. Then, from your app:
Run well-actually-create-rules to decompose your existing CLAUDE.md into .claude/rules/ . It proposes a slimmed CLAUDE.md (written to CLAUDE.md.proposed for you to diff) and never overwrites the original without approval.
Run well-actually-documentation-full to generate documentation/ from your code and write the first .last-synced marker.
Maintain with well-actually-documentation-recent as you ship.
The skills create .claude/rules/ and documentation/ if those folders don't exist, so there's nothing to set up by hand. ( well-actually-get-started runs steps 1–2 together; on an existing app you can also just run them directly.)
These files are deliberately generic — they're prompts, not prescriptions. Delete the rules you don't need, add the concerns you do, and tune each rule's paths: glob to your actual layout — that's the one thing worth getting right, since it decides when the rule fires. Keep every file short. The moment a rule file turns into a second encyclopedia, you're back to carrying the backpack.
A starter layout for keeping Claude Code's context lean as your codebase grows: path-scoped rules that auto-load by file, on-demand system docs, and skills to maintain both.
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
