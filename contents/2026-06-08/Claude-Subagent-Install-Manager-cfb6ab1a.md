---
source: "https://github.com/wastedcode/truecast"
hn_url: "https://news.ycombinator.com/item?id=48446932"
title: "Claude Subagent Install Manager"
article_title: "GitHub - wastedcode/truecast: The expert teammates Claude Code doesn't ship with · GitHub"
author: "zeppelin_7"
captured_at: "2026-06-08T16:27:56Z"
capture_tool: "hn-digest"
hn_id: 48446932
score: 2
comments: 0
posted_at: "2026-06-08T15:49:44Z"
tags:
  - hacker-news
  - translated
---

# Claude Subagent Install Manager

- HN: [48446932](https://news.ycombinator.com/item?id=48446932)
- Source: [github.com](https://github.com/wastedcode/truecast)
- Score: 2
- Comments: 0
- Posted: 2026-06-08T15:49:44Z

## Translation

タイトル: クロード サブエージェント インストール マネージャー
Article title: GitHub - wastedcode/truecast: The expert teammates Claude Code doesn't ship with · GitHub
Description: The expert teammates Claude Code doesn't ship with - wastedcode/truecast

記事本文:
GitHub - Wastedcode/truecast: 専門家チームメイトのクロード コードには付属していません · GitHub
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
無駄なコード
/
トゥルーキャスト
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
主な支店のタグ

ファイル コードに移動 その他のアクション メニューを開く フォルダーとファイル
35 コミット 35 コミット .github .github docs docs personas personas src src .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md biome.json biome.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Claude Code にはエキスパートのチームメイトは付属していません。
ポータブルでバージョン管理されたエキスパートペルソナ (製品マネージャー、アーキテクト、セキュリティレビュー担当者) をインストールします。
— 任意のプロジェクトに追加し、クロード コードで実行し、作成者が改良したときに独自のカスタマイズを維持します
彼ら。
ステータス: エンドツーエンドで動作中 — インストール、更新、リスト、削除、ドクター、プロンプト、
ペルソナごとの所有権台帳、アトミックな更新、サンドボックス化された git/GitHub フェッチ。実際の運転を確認済み
Claude Code セッション (サブエージェントおよびスタンドアロン エージェントとして)。 1.0 以前。次は自己改善のループです。
ペルソナは、小さな、greppable コーパス + アイデンティティであり、2 つの所有者に分割されます。
core/ — プロバイダーの技術: Agent.md (アイデンティティ) + スキル/ + 知識/ 、によってインデックス付けされます。
ペルソナ.toml 。読み取り専用。 1 つのグローバル コピー。意図的にアップデートを採用しているのです。
instance/ — プロジェクトごとのジョブ ( mandate.md ) + 蓄積されたメモ ( work.md )。あなたの、
リポジトリ内でコミットされ、更新によって変更されることはありません。
バンドルされた例: personas/product-manager/ 。
npm install -g @wastedcode/truecast # `truecast` コマンドをインストールします
truecast --ヘルプ
またはソースから:
git clone https://github.com/wastedcode/truecast && cd truecast
pnpm インストール && pnpm ビルド
npm link # `truecast` を PATH に置きます
ノード ≥ 20 が必要です。1.0 より前: CLI とプログラム API は変更される可能性があります。

0.x 未満の未成年者 — を参照
ドキュメント → 安定性 。
プロジェクトの cd
truecast install < git-url-or-path > [@バージョン][ # サブパス]
# 例
truecast インストール https://github.com/wastedcode/truecast#personas/product-manager # 公式ペルソナ
truecast install ./personas/product-manager # ローカル パス
truecast install https://github.com/you/persona@1.2.0 # GitHub リリース タグ
truecast install https://github.com/you/monorepo#personas/pm # サブディレクトリ内のペルソナ
公式ペルソナ (プロダクトマネージャー、ソフトウェアエンジニア、ソフトウェアアーキテクト、セキュリティエンジニア、
qa 、インフラストラクチャ、製品マーケティング担当者、ui-ux-デザイナー、販売) はこのリポジトリに存在します
personas/ — …/truecast#personas/<name> を使用していずれかをインストールします。
次に、このプロジェクトのペルソナのジョブを .truecast/agents/<name>/instance/mandate.md に書き込みます。
インストールすると、ネイティブのクロード コード サブエージェントが ~/.claude/agents/<name>.md に生成され、シンボリック リンクが作成されます。
プロジェクトに組み込んでください。その体にはペルソナのスキルのインデックスが付いています (それぞれに 1 行のインデックスが付いています)
概要と読み取りへのパス）により、ペルソナはオンデマンドで適切なスキルを取得します。検証済み：
オープンエンド タスク it 一致する SKILL.md ファイル自体を読み取り、それらを適用します。
ペルソナは 3 つの方法で実行できます。
1. クロード コードのサブエージェントとして ( @agent-<name> )
インストール後に Claude Code を再起動し、通常のセッションに戻します。
> プロダクト マネージャーにこのアイデアをプレッシャー テストしてもらいます: 私の To-Do に自動的に優先順位を付ける AI
クロードはサブエージェントに委任します (サブエージェントは /agents の下にリストされており、 @agent-product-manager で実行できます)
明示的に）。サブエージェントは、ペルソナが宣言したツール ( persona.toml 内のツール) を使用して実行されます。
プロジェクト ジョブの mandate.md を読み取り、必要なスキルを読み取ります。
2. スタンドアロンのクロードとして (ペルソナがメインエージェントです)
完全なクロード セッションを実行します。

これがペルソナです。その全体がシステム プロンプトとして読み込まれます。
truecast プロンプトは、その構成されたプロンプトを出力します。それを claude にパイプします:
cd your-project # `truecast install` を実行したプロジェクト
クロード --append-system-prompt " $( truecast プロンプト プロダクト マネージャー ) " \
--allowed-tools Grep WebSearch WebFetch の読み取り \
--モデル作品
今ではセッション全体がペルソナのように考えられます。 --allowed-tools は、ペルソナのツールに制限します。
宣言します (そうでない場合、メイン エージェントは完全なツールセットを持っています)。 --model はその modelHint と一致します。 ( truecast プロンプト <name> はシステム プロンプトを出力するだけです。 --append-system-prompt-file <file> も機能します。)
3. 永続的なプログラム可能なエージェントとして (claudemux)
claudemux は、実際の、ログインに基づいたクロード セッションを実行します。
名前によって、ターンが実際にいつ完了したかがわかります。 -- の後のフラグは直接転送されます
claude なので、ペルソナを完全なスタンドアロン エージェントとして起動し、時間をかけて対話します。
# ペルソナのプロンプトを 1 回レンダリングします
truecast プロンプト製品マネージャー > /tmp/pm.md
# 完全な永続的なクロード セッションとして生成します (`--` 以降はすべてクロードに渡されます)
claudemux spawn pm --cwd ./your-project --trust-workspace -- \
--append-system-prompt-file /tmp/pm.md \
--allowed-tools Read Grep WebSearch WebFetch --model opus
# その後、それを実行します — ワンショット、または送信/待機/メッセージ
claudemux は午後に「次は X を構築することを考えています。価値はありますか?」と尋ねます。
各ペルソナは名前でアドレス指定される独自のセッションです。tmux を接続して動作を確認したり、全体をスピンアップしたりできます。
チーム (claudemux spawn アーキテクト ... / spawn security ... ) を 1 か所から調整します。 CLI マップ
コードから実行したい場合は、ノード ライブラリに 1:1 で接続します ( create({ name, cwd, extraArgs: [...] }) )。
truecast list [--check] [--project] # 何がインストールされているか。実行中と最新。ここに何が付いているのか
トゥルーキャスト更新 [ < 名前 >

] [--dry-run] # 新しいクラフトを採用します。分類済み (パッチ/マイナー/メジャー)。あなたのインスタンスはそのままです
truecast delete < name > [--global] # このプロジェクトから切り離す (instance/ を保持する)、またはグローバルにパージする
truecast ドクター [--fix] # 検査 + 修復 (ドリフト、ぶら下がっているポインタ、古くなったアーティファクト)
完全なモデル (consent、drift/ --force 、
識別された更新結果）。
install はプロジェクト (最も近い囲み git リポジトリ、または --project <path> ) を解決し、ペルソナを
グローバル キャッシュ ( ~/.truecast/personas/<name>/ )、クラフトをプロジェクトにシンボリックリンクし、足場を構築します。
編集可能なinstance/ で、サブエージェントを ~/.claude/agents/<name>.md に実体化します。その本体インデックスは
ペルソナがオンデマンドで読み取るスキル/知識 (スキルはグローバル スラッシュとしてコピーされません)
コマンド。それらはペルソナの私物です）。トゥルーキャストが書き込むすべてのファイルはペルソナごとに追跡されます
台帳 (owned.json)、個人ごとのロックの下にあるため、同時インストールは決して衝突せず、トゥルーキャストされます
所有していないファイルを上書きすることはありません。
docs/ — インストール、ペルソナの管理、
オーサリングペルソナ、および出荷済みと計画済みのステータス表。 (歩調を合わせて
コードの場合: 機能は文書化されるまで完了しません)。
pnpmインストール
pnpm typecheck && pnpm test && pnpm lint && pnpm build
ノード ≥ 20 および pnpm が必要です。ライセンス: MIT (意図)。設計メモは、internal/ (git-ignoror) に保存されます。
専門家チームメイトの Claude Code は付属していません
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私のPEを共有しないでください

個人情報

## Original Extract

The expert teammates Claude Code doesn't ship with - wastedcode/truecast

GitHub - wastedcode/truecast: The expert teammates Claude Code doesn't ship with · GitHub
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
wastedcode
/
truecast
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
35 Commits 35 Commits .github .github docs docs personas personas src src .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md biome.json biome.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
The expert teammates Claude Code doesn't ship with.
Install portable, versioned expert personas — a product manager, an architect, a security reviewer
— into any project, run them in Claude Code, and keep your own customizations when the author improves
them.
Status: working end-to-end — install · update · list · remove · doctor · prompt , with a
per-persona ownership ledger, atomic updates, and a sandboxed git/GitHub fetch. Verified driving real
Claude Code sessions (as a subagent and as a standalone agent). Pre-1.0; the self-improving loop is next.
A persona is a small, greppable corpus + an identity, split into two owners:
core/ — the provider's craft: agent.md (identity) + skills/ + knowledge/ , indexed by
persona.toml . Read-only; one global copy; you adopt updates deliberately.
instance/ — your per-project job ( mandate.md ) + accumulated notes ( work.md ). Yours,
committed in your repo, never touched by an update.
A bundled example: personas/product-manager/ .
npm install -g @wastedcode/truecast # installs the `truecast` command
truecast --help
Or from source:
git clone https://github.com/wastedcode/truecast && cd truecast
pnpm install && pnpm build
npm link # puts `truecast` on your PATH
Requires Node ≥ 20. Pre-1.0: the CLI and the programmatic API may change between 0.x minors — see
docs → Stability .
cd your-project
truecast install < git-url-or-path > [@version][ # subpath]
# examples
truecast install https://github.com/wastedcode/truecast#personas/product-manager # an official persona
truecast install ./personas/product-manager # local path
truecast install https://github.com/you/persona@1.2.0 # a GitHub release tag
truecast install https://github.com/you/monorepo#personas/pm # a persona in a sub-directory
The official personas ( product-manager , software-engineer , software-architect , security-engineer ,
qa , infrastructure , product-marketer , ui-ux-designer , sales ) live in this repo under
personas/ — install any of them with …/truecast#personas/<name> .
Then write the persona's job for this project in .truecast/agents/<name>/instance/mandate.md .
Installing generates a native Claude Code subagent at ~/.claude/agents/<name>.md and symlinks the
craft into your project. Its body carries an index of the persona's skills (each with a one-line
summary and the path to Read), so the persona pulls the right skill on demand — verified: given an
open-ended task it Reads the matching SKILL.md files itself, then applies them.
You can run a persona three ways.
1. As a Claude Code subagent ( @agent-<name> )
Restart Claude Code after installing, then bring it into a normal session:
> have the product-manager pressure-test this idea: an AI that auto-prioritizes my to-dos
Claude delegates to the subagent (it's listed under /agents , and you can @agent-product-manager it
explicitly). The subagent runs with the tools the persona declares ( tools in its persona.toml ),
reads its mandate.md for the project job, and Reads the skills it needs.
2. As a standalone claude (the persona is the main agent)
Run a full claude session that is the persona — its whole craft loaded as the system prompt.
truecast prompt emits that composed prompt; pipe it into claude :
cd your-project # the project where you ran `truecast install`
claude --append-system-prompt " $( truecast prompt product-manager ) " \
--allowed-tools Read Grep WebSearch WebFetch \
--model opus
Now the whole session thinks like the persona. --allowed-tools restricts it to the tools the persona
declares (a main agent otherwise has the full toolset); --model matches its modelHint . ( truecast prompt <name> just prints the system prompt — --append-system-prompt-file <file> works too.)
3. As persistent, programmable agents (claudemux)
claudemux runs real, login-backed claude sessions you drive
by name and that tell you when a turn is actually done. Flags after -- are forwarded straight to
claude , so you launch a persona as a full standalone agent and talk to it over time:
# render the persona's prompt once
truecast prompt product-manager > /tmp/pm.md
# spawn it as a full, persistent claude session (everything after `--` goes to claude)
claudemux spawn pm --cwd ./your-project --trust-workspace -- \
--append-system-prompt-file /tmp/pm.md \
--allowed-tools Read Grep WebSearch WebFetch --model opus
# then drive it — one-shot, or send/wait/messages
claudemux ask pm " We're thinking of building X next. Worth it? "
Each persona is its own session, addressed by name — tmux attach to watch it work, or spin up a whole
team ( claudemux spawn architect … / spawn security … ) and coordinate them from one place. The CLI maps
1:1 to a Node library if you'd rather drive it from code ( create({ name, cwd, extraArgs: [...] }) ).
truecast list [--check] [--project] # what's installed; running vs latest; what's attached here
truecast update [ < name > ] [--dry-run] # adopt newer craft; classified (patch/minor/major); your instance untouched
truecast remove < name > [--global] # detach from this project (keeps instance/), or purge globally
truecast doctor [--fix] # inspect + repair (drift, dangling pointer, stale artifacts)
See docs/managing-personas.md for the full model (consent, drift/ --force ,
the discriminated update outcomes).
install resolves the project (the nearest enclosing git repo, or --project <path> ), fetches the persona into a
global cache ( ~/.truecast/personas/<name>/ ), symlinks the craft into the project, scaffolds your
editable instance/ , and materializes the subagent at ~/.claude/agents/<name>.md — whose body indexes
the skills/knowledge for the persona to Read on demand (skills are not copied as global slash-
commands; they're the persona's private craft). Every file truecast writes is tracked in a per-persona
ledger ( owned.json ), under a per-persona lock, so concurrent installs never collide and truecast
never clobbers a file it doesn't own.
docs/ — install , managing personas ,
authoring personas , and a shipped-vs-planned status table. (Kept in step
with the code: a feature isn't done until it's documented.)
pnpm install
pnpm typecheck && pnpm test && pnpm lint && pnpm build
Requires Node ≥ 20 and pnpm. License: MIT (intended). Design notes live in internal/ (git-ignored).
The expert teammates Claude Code doesn't ship with
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
