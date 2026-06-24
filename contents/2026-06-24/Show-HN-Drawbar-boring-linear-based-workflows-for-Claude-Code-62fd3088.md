---
source: "https://github.com/mjn298/drawbar"
hn_url: "https://news.ycombinator.com/item?id=48665878"
title: "Show HN: Drawbar – boring linear based workflows for Claude Code"
article_title: "GitHub - mjn298/drawbar · GitHub"
author: "seedlessmike"
captured_at: "2026-06-24T22:28:10Z"
capture_tool: "hn-digest"
hn_id: 48665878
score: 2
comments: 0
posted_at: "2026-06-24T21:32:24Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Drawbar – boring linear based workflows for Claude Code

- HN: [48665878](https://news.ycombinator.com/item?id=48665878)
- Source: [github.com](https://github.com/mjn298/drawbar)
- Score: 2
- Comments: 0
- Posted: 2026-06-24T21:32:24Z

## Translation

タイトル: Show HN: Drawbar – クロード コードの退屈な線形ベースのワークフロー
記事タイトル: GitHub - mjn298/drawbar · GitHub
説明: GitHub でアカウントを作成して、mjn298/drawbar の開発に貢献します。
HN テキスト: 私は Beads を使用していましたが、非常にバグが多く、望ましいと思われるよりもはるかに高いスループットのワークフローに合わせて調整されていました。それよりもゆっくりと仕事をしています。他にいくつかの素晴らしいプロジェクト (lavra とリニア ビーズ、リポジトリの Readme でクレジットされています) を見たことがあります。これらは蓄積された知識ベースを追加し、ビーズの概念にリニアを加えたものでした。そこで、両方のアイデアを盗んで、自分用のプラグインに入れてみました。私は数日間それをドッグフーディングしてきました。かなりうまく機能しますが、ユーザーは私だけであり、私の好みに合わせて意見が分かれています。また、自作の線形 CLI よりもはるかに遅く、トークン効率が低い線形 MCP も使用します。今後更新するかもしれません。興味があれば、ぜひチェックしてみてください。

記事本文:
GitHub - mjn298/drawbar · GitHub
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
mjn298
/
ドローバー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
44℃

省略 44 コミット .claude-plugin .claude-plugin エージェント エージェント コマンド コマンド docs/ superpowers docs/ superpowers scripts scripts skill/drawbar-knowledge skill/drawbar-knowledge .gitignore .gitignore README.md README.md package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
クロード コードの無駄のないリニア ネイティブ設計→計画→作業→学習ワークフロー。プロジェクトごとに複合的なナレッジ ベースを備えています。
AI が生成した Readme と、以下の本文中の私の介入文を組み合わせてお楽しみください。
クローンホイールを手に入れたので、たくさん掘っています。それ以外に理由はありません。私は不可解な名前のソフトウェア プロジェクトの壮大な伝統に従っています。
ドローバーは、Beads の作業方法 (原子的で追跡可能な問題と、蒸発するのではなくセッション全体で複合化する知識) を愛するが、次のような人を対象としています。
リニアに住んでいます。問題は線形のままで、チーム全体に表示されます。仕様は親問題の説明であり、ストーリーはサブ問題の順序で配置され、ステータスは Todo → In Progress → Done と流れます。問題はローカルに保存されません。これにより処理が遅くなり、おそらくトークンが無駄になることはわかっていますが、これは私と私のチームにとってより効果的です。
完全な自律性ではなく、人間が関与することを望んでいます。ドローバーは、一度だけ実行するエージェントではなく、範囲を絞り、アプローチを選択し、計画を承認し、作業をレビューする、意図的でレビュー可能な一連のステップです。ビーズが覚醒剤なら、これはアイスティーのグラスです。
これは実際には、接着剤を使った昔ながらの問題追跡ワークフローにすぎません。目新しいものではありませんが、私にとってはうまくいきました。
ナレッジベースはその中心です。キャプチャしたすべてのレッスン、決定、および「必須チェック」は次の設計/計画/作業に反映されるため、システムは使えば使うほど鋭くなっていきます。
エージェントの自律性を最大限に高めたい場合は、ドローバーも必要になるでしょう

実践的な。あなたが意思決定者でありながら、草案を作成して提案してくれるエージェント、そして実際に蓄積される記憶が必要な場合、それがポイントです。
ドローバーが機能する前に、次の 2 つの設定が必要です。
Bun —drawbar-kb ナレッジ CLI (およびテスト) を実行します。インストール:curl -fsSL https://bun.sh/install |バッシュ。
クロード コードで接続された Linear MCP — ドローバーは、MCP を介して Linear でのすべての作業を追跡します (mcp__…Linear… ツールがセッションで使用可能である必要があります)。これがなければ、設計/計画/作業コマンドはローカルで実行できますが、Linear に書き込むことはできません。
さらに、Claude Code 自体 (プラグイン ホスト)。
マーケットプレイスを追加し、プラグインをインストールします。
クロードプラグインマーケットプレイスにmjn298/drawbarを追加
クロードプラグインのインストールdrawbar@drawbar
次に、/reload-plugins を実行 (または再起動) して、/drawbar-* コマンドをロードします。
プロジェクト内で /drawbar-setup を 1 回実行します。これは、drawbar-kb CLI を PATH にリンクし (インストールされたプラグインからの bun リンク経由)、 <project>/.drawbar/memory/ を初期化します。
その後もdrawbar-kbが見つからない場合は、BunのグローバルビンをPATHに追加します。
エクスポート PATH = " $( bun pm bin -g ) : $PATH "
ドローバー自体をハッキング?マーケットプレイスをスキップしてリポジトリを直接ロードします。「開発」を参照してください。
/drawbar-setup [legacyKnowledge.jsonl] # マシンごと + プロジェクトごとに 1 回
/drawbar-design <機能 | issue-id> # spec → 線形親問題
/drawbar-plan <issue-id> # 注文されたストーリーのサブイシュー
/drawbar-work <issue-id> # 次のストーリー、TDD を実装します。
/drawbar-learn [issue-id] # レッスンを KB にキュレートします
issue-id は、線形課題識別子 (例: ABC-123 )、つまりチーム独自のプレフィックスです。
/drawbar-setup の [legacyKnowledge.jsonl] 引数は、既存の lavra ナレッジ ベースをdrawbar に移行するためだけに使用されます。これまで lavra を使用したことがない場合 (ほとんどの人)、それをスキップして、新しい空のナレッジ ベースから始めてください。

。
ナレッジは <project>/.drawbar/memory/ に存在します (JSONL はコミットされ、SQLite インデックスは gitignore されます)。いつでも直接クエリできます。
drawbar-kb リコール " dynamodb tenancy " --dir " $PWD /.drawbar/memory " --json
drawbar-kb stats --dir " $PWD /.drawbar/memory "
開発
ドローバー自体に取り組んでいますか？ kb ツールには次のテストがあります。
パンテスト
編集内容のロード — 2 つのモード:
高速反復 → ライブソースを直接実行します。 --plugin-dir を指定して Claude Code を起動すると、作業コピーがロードされ、プラグイン キャッシュが完全にバイパスされます。
クロード --plugin-dir /path/to/drawbar
コマンド/エージェント/スキルへの編集は、/reload-plugins で有効になります。バージョンアップや再インストールは必要ありません。インストールされているコピーはそのセッションのみシャドウイングされます。
マーケットプレイスを通じて変更をリリース → バージョンを上げます。プラグイン キャッシュは plugin.json version によってキー設定されるため、同じバージョンでファイルを編集してもサイレントに反映されません。 /reload-plugins は古いキャッシュされたコピーを提供し続けます。実際の変更を送信するには:
# 1. .claude-plugin/plugin.json の「バージョン」をバンプします (例: 0.1.1 → 0.1.2)
クロードプラグインマーケットプレイス更新 <マーケットプレイス> # ソースを再読します
クロードプラグイン更新drawbar@ <マーケットプレイス> # 新しいバージョンを取得します
# 2. セッション内の /reload-plugins (または再起動)
バージョン バンプを忘れることは、「変更が反映されない」という最大の落とし穴です。
ドローバーは 3 つのプロジェクトの肩の上に立っています。
Beads — それを開始したモデル: 作業中に取得されるアトミックで依存関係を意識した問題と知識が、セッション間で複合化されます。
Linear-beads by @nikvdp — バックエンドとして Linear を使用したビーズ スタイルの問題追跡。問題をローカル データベースではなくリニア (人間に見える) に保存するためのインスピレーション。
@roberto-mello による lavra — 設計 → 計画 → 作業 → レビュー → スキルの学習パイプラインと知識複合アプローチ

ドローバーのワークフローがモデル化されています。
drawbar は、人間参加型の作業用に構築された独自のナレッジ ベースを中心に、lavra のワークフローと Linear-Beads の Linear-first トラッキングを組み合わせた、より効率的なバージョンです。
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

Contribute to mjn298/drawbar development by creating an account on GitHub.

I had been using Beads and it was really buggy and tailored to a workflow with much higher throughput than I found desirable. I work more slowly than that. I saw some other cool projects (lavra and linear-beads, credieted in my repo readme) which added an accumulating knowledge base and linear to the beads concept, so I just stole both ideas and put them into a plugin for myself. I have been dogfooding it for a couple of days. It works pretty well but I'm the only user and it's opinionated around how I like to do things. It also uses the linear MCP which is way slower and less token-efficient than the linear CLI on homebrew. I might update it in the future. If this appeals to you, check it out.

GitHub - mjn298/drawbar · GitHub
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
mjn298
/
drawbar
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
44 Commits 44 Commits .claude-plugin .claude-plugin agents agents commands commands docs/ superpowers docs/ superpowers scripts scripts skills/ drawbar-knowledge skills/ drawbar-knowledge .gitignore .gitignore README.md README.md package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
A lean, Linear-native design → plan → work → learn workflow for Claude Code, with a per-project compounding knowledge base.
Enjoy a mix of AI generated readme, and my interjections in the text below.
I just got a clonewheel and am digging it, a lot. Other than that, no reason. I'm following the grand tradition of software projects with inscrutable names.
drawbar is for people who love the Beads way of working — atomic, traceable issues and knowledge that compounds across sessions instead of evaporating — but who:
live in Linear. Your issues stay in Linear, visible to the whole team: a spec is a parent issue's description, stories are ordered sub-issues, status flows Todo → In Progress → Done . Issues aren't stored locally. I know this makes things slower and probably wastes tokens, but this works better for my and my team.
want a human in the loop, not full autonomy. drawbar is a set of deliberate, reviewable steps — you sharpen scope, pick the approach, approve the plan, and review the work — rather than a fire-and-forget agent. If beads is meth, this is a glass of iced tea.
This is really just an old-school issue tracking workflow with some glue. It is not novel but it has been working well for me.
The knowledge base is the heart of it: every lesson, decision, and "MUST-CHECK" you capture is recalled on the next design/plan/work, so the system gets sharper the more you use it.
If you want maximum agent autonomy, drawbar will feel too hands-on. If you want an agent that drafts and proposes while you stay the decision-maker — and a memory that actually accumulates — that's the point.
drawbar needs two things set up before it works:
Bun — runs the drawbar-kb knowledge CLI (and the tests). Install: curl -fsSL https://bun.sh/install | bash .
The Linear MCP, connected in Claude Code — drawbar tracks all work in Linear through the MCP ( mcp__…Linear… tools must be available in your session). Without it, the design/plan/work commands can still run locally but won't write to Linear.
Plus Claude Code itself (the plugin host).
Add the marketplace and install the plugin:
claude plugin marketplace add mjn298/drawbar
claude plugin install drawbar@drawbar
Then run /reload-plugins (or restart) so the /drawbar-* commands load.
Run /drawbar-setup once in a project. It links the drawbar-kb CLI onto your PATH (via bun link from the installed plugin) and initializes <project>/.drawbar/memory/ .
If drawbar-kb still isn't found afterward, add Bun's global bin to PATH:
export PATH= " $( bun pm bin -g ) : $PATH "
Hacking on drawbar itself? Skip the marketplace and load the repo directly — see Development .
/drawbar-setup [legacy knowledge.jsonl] # once per machine + per project
/drawbar-design <feature | issue-id> # spec → Linear parent issue
/drawbar-plan <issue-id> # ordered story sub-issues
/drawbar-work <issue-id> # implement next story, TDD
/drawbar-learn [issue-id] # curate lessons into the KB
issue-id is a Linear issue identifier (e.g. ABC-123 ) — your team's own prefix.
The [legacy knowledge.jsonl] argument to /drawbar-setup is only for migrating an existing lavra knowledge base into drawbar. If you weren't using lavra before — most people — skip it and start with a fresh, empty knowledge base.
Knowledge lives in <project>/.drawbar/memory/ (the JSONL is committed; the SQLite index is gitignored). Query it directly anytime:
drawbar-kb recall " dynamodb tenancy " --dir " $PWD /.drawbar/memory " --json
drawbar-kb stats --dir " $PWD /.drawbar/memory "
Development
Working on drawbar itself? The kb tool has tests:
bun test
Loading your edits — two modes:
Iterating fast → run the live source directly. Launch Claude Code with --plugin-dir so it loads your working copy and bypasses the plugin cache entirely:
claude --plugin-dir /path/to/drawbar
Edits to commands/agents/skills take effect on /reload-plugins — no version bump, no reinstall. It shadows any installed copy for that session only.
Releasing a change through the marketplace → bump the version. The plugin cache is keyed by plugin.json version , so editing files in place at the same version silently won't propagate — /reload-plugins keeps serving the stale cached copy. To ship a real change:
# 1. bump "version" in .claude-plugin/plugin.json (e.g. 0.1.1 → 0.1.2)
claude plugin marketplace update < marketplace > # re-read the source
claude plugin update drawbar@ < marketplace > # fetch the new version
# 2. /reload-plugins (or restart) in your session
Forgetting the version bump is the #1 "my change didn't show up" gotcha.
drawbar stands on the shoulders of three projects:
Beads — the model that started it: atomic, dependency-aware issues and knowledge captured as you work so it compounds across sessions.
linear-beads by @nikvdp — beads-style issue tracking with Linear as the backend. The inspiration for keeping issues in Linear (human-visible) instead of a local database.
lavra by @roberto-mello — the design → plan → work → review → learn skill pipeline and knowledge-compounding approach that drawbar's workflow is modeled on.
drawbar is a leaner take that combines lavra's workflow with linear-beads' Linear-first tracking, around its own knowledge base — built for human-in-the-loop work.
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
