---
source: "https://github.com/quan0715/open-press"
hn_url: "https://news.ycombinator.com/item?id=48375586"
title: "I built a React and MDX workspace so AI agents stop generating giant HTML files"
article_title: "GitHub - quan0715/open-press: AI-first fixed-layout document workspaces. Your AI agent writes proposals, whitepapers, theses, books — A4, PDF, web reader. · GitHub"
author: "quan0715"
captured_at: "2026-06-03T00:36:26Z"
capture_tool: "hn-digest"
hn_id: 48375586
score: 1
comments: 0
posted_at: "2026-06-02T20:12:53Z"
tags:
  - hacker-news
  - translated
---

# I built a React and MDX workspace so AI agents stop generating giant HTML files

- HN: [48375586](https://news.ycombinator.com/item?id=48375586)
- Source: [github.com](https://github.com/quan0715/open-press)
- Score: 1
- Comments: 0
- Posted: 2026-06-02T20:12:53Z

## Translation

タイトル: AI エージェントが巨大な HTML ファイルを生成しないように React と MDX ワークスペースを構築しました
記事のタイトル: GitHub - quan0715/open-press: AI ファーストの固定レイアウトのドキュメント ワークスペース。 AI エージェントは、提案書、ホワイトペーパー、論文、書籍 (A4、PDF、Web リーダー) を作成します。 · GitHub
説明: AI ファーストの固定レイアウトのドキュメント ワークスペース。 AI エージェントは、提案書、ホワイトペーパー、論文、書籍 (A4、PDF、Web リーダー) を作成します。 - quan0715/オープンプレス

記事本文:
GitHub - quan0715/open-press: AI ファーストの固定レイアウトのドキュメント ワークスペース。 AI エージェントは、提案書、ホワイトペーパー、論文、書籍 (A4、PDF、Web リーダー) を作成します。 · GitHub
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
クアン0715
/
オープンプレス
公共
通知
変更するにはサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
141 コミット 141 コミット .changeset .changeset .github/ workflows .github/ workflows apps/ web apps/ web docs docs パッケージ パッケージ プレス プレス スキル スキル .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md ライセンスライセンス README.md README.md Index.html Index.html package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.json tsconfig.jsonturbo.jsonturbo.json vite.config.ts vite.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
固定レイアウトのアプリケーション層により、インライン編集 / 開発サーバー / コメント マーカー / レンダリング / エクスポートを実装するための創造的なスキルが不要になります。 Skill はクリエイティブな決定 (何を行うか) を担当し、OpenPress は固定レイアウトのワークベンチ (どのように表示するか) を担当します。提案書、ホワイト ペーパー、配布資料、書籍、コミュニティ投稿、プレゼンテーション - 固定サイズ レイアウト、PDF 出力、Web リーダー。
ステータス: v1.0。 npm で @open-press/cli および @open-press/core として公開されます。ランディング サイトは open-press.dev です。
コンテンツが変化し続けるが、出力形式を安定させる必要がある場合は、open-press を使用します。
適したもの: 提案書、ビジネス プラン、ホワイトペーパー、研究レポート、製品仕様書、書籍、ハンドブック、コース ノート、学習ガイド、技術レポート、編集長文、ブランド レポート、および外部のクリエイティブ スキルによって提供される固定形式の出力物。
1 回限りのチャット回答、自由形式の画像編集、固定ページ スケーリングではなくライブ応答リフローが必要なドキュメントなどにはあまり役に立ちません。
node -v 、 npm -v 、または npx -v が使用できない場合は、最初に Node.js LTS をインストールします。 OpenPress には Node.js 20 以降が必要です。
npx @open-press/cli init my-doc
cd 私のドキュメント
npm rundev
以上: npm インストールと OpenPress スキルのセットアップ

初期化中に自動的に発生します。スキルがプレス/ワークスペースを追加したら、Vite によって出力されたローカル URL (通常は http://127.0.0.1:5173/workspace ) を開いてワークベンチを表示します。
CLI はスターターをフェッチしません。独自の出発点が必要な場合は、スキルを使用します。
npx -y skill@latest add quan0715/openpress-social-card-skill
ソーシャル カードの場合は、 quan0715/openpress-social-card-skill を使用します。削除されたバンドルの 1080×1080 正方形スターターではなく、1080×1350 (4:5 ポートレート) をターゲットとしています。
スキルは、インテーク、サンプル、スターター ファイル、およびデザイン ルールを所有します。 OpenPress は、ランタイム、ワークベンチ、レンダリング、検証、およびエクスポートを所有します。このリポジトリのスターターを含むスキルは単なるスキルです。エージェントは外部スキルと同様にそれらを読み取って使用できます。
クロード・コード（フルスキルサポート）
我想寫一份 [提案 / 白皮書 / 講義 / 書]，幫我起手。
Claude Code は、skills/openpress-init/SKILL.md を自動ロードし、取り込み手順を案内します。
Codex CLI (フルスキルサポート)
コーデックス
同じ開始プロンプト。 Codex は、フレームワーク コントラクトのリポジトリ ルートにある AGENTS.md を読み取ります。
Copilot Chat / その他の LLM ツール (手動)
VS Code でフォルダーを開き、最初のメッセージの前に以下のシステム プロンプトを Copilot Chat に貼り付けます。 Copilot は SKILL ファイルを自動検出しません。プロンプトにより、ルーティング ルールがインラインで与えられます。
貼り付けた後、同じ冒頭メッセージを送信します: 「我想寫一份 [...]，幫我起手。」
オープンプレスのワークスペース、つまり AI ファーストの固定レイアウトのドキュメント フレームワークでの作業を手伝ってくれています。完全なルーティング ルールは `.claude/skills/openpress/SKILL.md` (または `.agents/skills/openpress/SKILL.md`) にあります。オンデマンドで読んでください。
空のディレクトリから開始します。
- 最初に `node -v`、`npm -v`、および `npx -v` を実行します。見つからない場合は、停止して、Node.js LTS をインストールし、ターミナルを再度開いて、再試行するように指示します。
- 文書の種類、対象者、

第一言語、範囲、メタデータ (タイトル/サブタイトル/組織/著者)。メタデータが収集される前に init を実行しないでください。
- 次に、メタデータ フラグを指定して `npx @open-press/cli init .` を実行します。ターゲットが空でない場合は、最初にそれをクリーンアップするように依頼してください (単独の `.git/` だけでも問題ありません)。
- ユーザーが独自の形式を必要とする場合は、「npx -y skill@latest add <owner/repo>」を使用して関連するスキルをインストールし、その「SKILL.md」を読み取り、そのスキルのスターター/サンプル ファイルを OpenPress ワークスペースにコピーまたは適応させます。
- ワークスペース ソース ツリーを追加した後: `npm run build` を実行して、正しくレンダリングされることを確認します。
既存のオープンプレスワークスペース (前の init からすでに `press/` が含まれているワークスペース) で作業します。
- `press/`、`.claude/skills/`、`.agents/skills/`、およびルート設定ファイル (`package.json` は `"openpress"` フィールドが存在する場所) にあるファイルのみを編集します。
- `public/openpress/`、`dist-react/`、`.deploy/`、または `.openpress/` は手動で編集しないでください。これらは生成されます。
- `node_modules/@open-press/` の下のフレームワーク コードは読み取り専用として扱われます。
執筆内容：
- ワークスペース `<ワークスペース>` / `<プレスソース>` ツリーに登録されている MDX ファイルを編集します。
- 繁体字中国語コンテンツ: `.claude/skills/chinese-ai-writing-polish/SKILL.md` を適用します。
- 学習者向けコンテンツ (講義 / 教材 / 教科書): `.claude/skills/teaching-notes-writing/SKILL.md` を適用します。
- キャプション付きテーブルの前に `<TableCaption>...</TableCaption>` を使用します。図や表の番号を手書きしないでください。
視覚的/構造的:
- テーマトークン、コンポーネント、ページリズム → `press/theme/` または `press/<slug>/components/` を編集します。
- ヘッダー、フッター、ページ番号などのページクロム、
[切り捨てられた]
環境を確認し、受け入れに関する質問をします: ノード/npm の可用性、ドキュメントの種類、対象読者、言語、範囲、タイトル/サブタイトル/組織/著者。
ジョブにスキルが必要な場合は、スキルをインストールまたは選択します。

ソーシャル カード、指導ノート、提案ワークフローなどの特定の形式。
メタデータ フラグを指定して init を実行します: npx @open-press/cli init 。 - タイトル "..." 。
スキルにソース ツリーを追加させます。スキルは独自のスターター/サンプルを読み取り、OpenPress ワークスペース ファイルを書き込みます。
検証: npm run build (検証 + レンダリング)。
Hand off : どのソース ファイルとスキルが次の編集を所有しているかを示します。
ここからはチャットを続けます。あなたはコンテンツを書きます。エージェントがツールを処理します。
固定レイアウト ページ — A4、ソーシャル スクエア、スライド 16:9、または独自のプリセット。下書き、リーダー、PDF の間で予期せぬリフローが発生することはありません。
Press Tree レンダリング — press/index.tsx は、 <Workspace> + <Press> 、 <Frame> 、原稿ヘルパー、および登録された MDX ソースを構成します。操作設定は、「openpress」の下の package.json にあります。
マルチプレス ワークスペース — 1 つの <Workspace> の下に 1 つのプロジェクト、多くの出力形式 (紙 + ソーシャル + スライド) があり、それぞれが独自のスラグにあります。
npm run dev ( http://127.0.0.1:5173/workspace ) にあるライブ ワークベンチ。マルチプレス プロジェクト用の Figma スタイル ギャラリーと /<press-slug>/preview にあるプレスごとのプレビューを備えています。
npm での PDF エクスポート openpress:pdf を実行し、ワークベンチ ツールバーからページごとの PNG エクスポートを実行します。
Cloudflare ページ経由のパブリック デプロイ — オプトイン、自動デプロイはありません。ターゲットプロジェクトに名前を付ける確認をオンにします。
@openpress-comment マーカー — リーダーにフィードバックをインラインに残します。 openpress-apply-comments ワークフロー スキルは、それらをソース編集として適用します。
.claude/skills/ および .agents/skills/ にインストールされたマルチツール エージェント スキル — Claude Code、Codex CLI、Cursor、Gemini CLI、Copilot、および 50 以上の AI エージェントで動作します。
→ エージェントファーストのウォークスルーについては、ランディング サイトを参照してください。
このリポジトリ内の press/ は、追跡されるドッグフード ワークスペースです。 OpenPress ユーザー ストーリー ブックを /userstory に加えて 2 つの最小限の social/slid でホストします。

e マルチプレス ルーティングを検証するためのプレス - 公共の着陸サイトとは別に:
pnpm run dev:workspace # Dogfood Press / ワークベンチ (マルチプレス ギャラリー)
pnpm run dev:web # open-press.dev ランディング サイト
ドッグフードは、ダウンストリーム ワークスペースと同じ CLI パスを使用します。
pnpm run build # すべてのプレスを public/openpress/<slug>/ にレンダリングします
pnpm 実行 openpress:pdf
pnpm run openpress:deploy:dry-run
もっと必要
したい
参照
コマンドを直接実行する
docs/cli.md
プレスツリーモデルの保守
docs/press-tree.md
ワークベンチ UI を使用する (コメント、メンション、プロジェクト アセット)
docs/workbench.md
スキルを理解する
docs/skills.md または skill/<skill>/SKILL.md を参照
リリースのカット/CD の設定
docs/release-and-deploy.md
オープンプレスへの貢献
COTRIBUTING.md および AGENTS.md
各リリースでの変更点を確認する
変更ログ.md
バグを報告する
GitHubの問題
ライセンス
AI ファーストの固定レイアウトのドキュメント ワークスペース。 AI エージェントは、提案書、ホワイトペーパー、論文、書籍 (A4、PDF、Web リーダー) を作成します。
読み込み中にエラーが発生しました。このページをリロードしてください。
1
フォーク
レポートリポジトリ
リリース
23
@open-press/core@1.2.0
最新の
2026 年 6 月 2 日
+ 22 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

AI-first fixed-layout document workspaces. Your AI agent writes proposals, whitepapers, theses, books — A4, PDF, web reader. - quan0715/open-press

GitHub - quan0715/open-press: AI-first fixed-layout document workspaces. Your AI agent writes proposals, whitepapers, theses, books — A4, PDF, web reader. · GitHub
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
quan0715
/
open-press
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
141 Commits 141 Commits .changeset .changeset .github/ workflows .github/ workflows apps/ web apps/ web docs docs packages packages press press skills skills .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md index.html index.html package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.json tsconfig.json turbo.json turbo.json vite.config.ts vite.config.ts View all files Repository files navigation
固定版面應用層 ,讓創作型 Skill 不用自己實作 inline edit / dev server / comment marker / 渲染 / 匯出。Skill 負責創作決策(要做什麼),OpenPress 負責固定版面工作台(怎麼呈現)。Proposal、白皮書、講義、書、社群貼文、簡報 — 固定尺寸版面,PDF 輸出,網頁 reader。
Status: v1.0. Published on npm as @open-press/cli and @open-press/core . Landing site at open-press.dev .
Use open-press when the content keeps changing but the output format must stay stable .
Good fits: proposals, business plans, whitepapers, research reports, product specs, books, handbooks, course notes, study guides, technical reports, editorial long-form, branded reports, and fixed-format outputs supplied by external creative skills.
Less useful for: one-off chat answers, free-form image editing, and documents that need live responsive reflow instead of fixed page scaling.
Install Node.js LTS first if node -v , npm -v , or npx -v is not available. OpenPress requires Node.js 20 or newer.
npx @open-press/cli init my-doc
cd my-doc
npm run dev
That's it: npm install and OpenPress skill setup happen automatically during init. Open the local URL printed by Vite (typically http://127.0.0.1:5173/workspace ) to see the workbench once a skill has added a press/ workspace.
The CLI does not fetch starters. Use a skill when you want an opinionated starting point:
npx -y skills@latest add quan0715/openpress-social-card-skill
For social cards, use quan0715/openpress-social-card-skill ; it targets 1080×1350 (4:5 portrait), not the removed bundled 1080×1080 square starter.
The skill owns intake, examples, starter files, and design rules. OpenPress owns the runtime, workbench, rendering, validation, and export. The starter-bearing skills in this repo are just skills; agents can read and use them like any external skill.
Claude Code (full skill support)
我想寫一份 [提案 / 白皮書 / 講義 / 書]，幫我起手。
Claude Code auto-loads skills/openpress-init/SKILL.md and walks you through intake.
Codex CLI (full skill support)
codex
Same opening prompt. Codex reads AGENTS.md at the repo root for the framework contract.
Copilot Chat / other LLM tools (manual)
Open the folder in VS Code, then paste the system prompt below into Copilot Chat before your first message. Copilot does not auto-discover SKILL files; the prompt gives it the routing rules inline.
After pasting, send the same opening message: 「我想寫一份 [...]，幫我起手。」
You are helping me work in an open-press workspace — an AI-first fixed-layout document framework. Full routing rules live in `.claude/skills/openpress/SKILL.md` (or `.agents/skills/openpress/SKILL.md`); read it on demand.
Starting from an EMPTY directory:
- First run `node -v`, `npm -v`, and `npx -v`. If missing, stop and tell me to install Node.js LTS, reopen the terminal, then retry.
- Ask for document type, audience, primary language, scope, and metadata (title / subtitle / organization / author). Do not run init before metadata is gathered.
- Then run `npx @open-press/cli init .` with metadata flags. If the target isn't empty, ask me to clean it first (a lone `.git/` is fine).
- If the user wants an opinionated format, install the relevant skill with `npx -y skills@latest add <owner/repo>`, read its `SKILL.md`, and copy or adapt that skill's starter/example files into the OpenPress workspace.
- After adding a workspace source tree: run `npm run build` to verify it renders cleanly.
Working in an EXISTING open-press workspace (one that already has `press/` from a previous init):
- Edit only files under `press/`, `.claude/skills/`, `.agents/skills/`, and root config files (`package.json` is where the `"openpress"` field lives).
- Never hand-edit `public/openpress/`, `dist-react/`, `.deploy/`, or `.openpress/` — those are generated.
- Framework code under `node_modules/@open-press/` is treated as read-only.
Writing content:
- Edit the MDX files registered by the workspace `<Workspace>` / `<Press sources>` tree.
- Traditional Chinese content: apply `.claude/skills/chinese-ai-writing-polish/SKILL.md`.
- Learner-facing content (講義 / 教材 / 課程): apply `.claude/skills/teaching-notes-writing/SKILL.md`.
- Use `<TableCaption>...</TableCaption>` before captioned tables; do not hand-write figure or table numbers.
Visual / structural:
- Theme tokens, components, page rhythm → edit `press/theme/` or `press/<slug>/components/`.
- Page chrome such as headers, footers, page numbers,
[truncated]
Check environment and ask intake questions : Node/npm availability, doc type, audience, language, scope, title / subtitle / organization / author.
Install or select a skill when the job needs a specific format, such as social cards, teaching notes, or a proposal workflow.
Run init with metadata flags : npx @open-press/cli init . --title "..." .
Let the skill add the source tree : the skill reads its own starter/examples and writes the OpenPress workspace files.
Verify : npm run build (validates + renders).
Hand off : tells you which source files and skills own the next edits.
From here, keep chatting. You write content; the agent handles tooling.
Fixed-layout pages — A4, social square, slide 16:9, or your own preset. No surprise reflow between draft, reader, and PDF.
Press Tree rendering — press/index.tsx composes <Workspace> + <Press> , <Frame> , manuscript helpers, and registered MDX sources. Operational settings live in package.json under "openpress" .
Multi-Press workspaces — one project, many output formats (paper + social + slides) under one <Workspace> , each at its own slug.
Live workbench at npm run dev ( http://127.0.0.1:5173/workspace ), with a Figma-style gallery for multi-Press projects and per-Press previews at /<press-slug>/preview .
PDF export at npm run openpress:pdf and per-page PNG export from the workbench toolbar.
Public deploy via Cloudflare Pages — opt-in, never auto-deployed; gated on confirmation naming the target project.
@openpress-comment markers — leave feedback inline in the reader; the openpress-apply-comments workflow skill applies them as source edits.
Multi-tool agent skills installed under .claude/skills/ and .agents/skills/ — works with Claude Code, Codex CLI, Cursor, Gemini CLI, Copilot, and 50+ other AI agents.
→ See the landing site for the agent-first walkthrough.
Inside this repository, press/ is the tracked dogfood workspace. It hosts the OpenPress User Story Book at /userstory plus two minimal social/slide Press for verifying multi-Press routing — separate from the public landing site:
pnpm run dev:workspace # dogfood press / workbench (multi-Press gallery)
pnpm run dev:web # open-press.dev landing site
The dogfood uses the same CLI path as downstream workspaces:
pnpm run build # render every Press into public/openpress/<slug>/
pnpm run openpress:pdf
pnpm run openpress:deploy:dry-run
Need more
Want to
See
Run commands directly
docs/cli.md
Maintain the Press Tree model
docs/press-tree.md
Use the workbench UI (comments, mentions, project assets)
docs/workbench.md
Understand the skills
docs/skills.md or browse skills/<skill>/SKILL.md
Cut a release / configure CD
docs/release-and-deploy.md
Contribute to open-press
CONTRIBUTING.md and AGENTS.md
See what changed in each release
CHANGELOG.md
Report bugs
GitHub Issues
License
AI-first fixed-layout document workspaces. Your AI agent writes proposals, whitepapers, theses, books — A4, PDF, web reader.
There was an error while loading. Please reload this page .
1
fork
Report repository
Releases
23
@open-press/core@1.2.0
Latest
Jun 2, 2026
+ 22 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
