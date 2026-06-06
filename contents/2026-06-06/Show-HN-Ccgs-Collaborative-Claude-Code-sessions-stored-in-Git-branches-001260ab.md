---
source: "https://github.com/ingram-technologies/claude-git-sessions"
hn_url: "https://news.ycombinator.com/item?id=48426297"
title: "Show HN: Ccgs – Collaborative Claude Code sessions, stored in Git branches"
article_title: "GitHub - ingram-technologies/claude-git-sessions: Share Claude Code sessions across a team through an orphan git branch (npm: claude-git-sessions) · GitHub"
author: "scrollaway"
captured_at: "2026-06-06T18:32:55Z"
capture_tool: "hn-digest"
hn_id: 48426297
score: 5
comments: 1
posted_at: "2026-06-06T16:04:45Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Ccgs – Collaborative Claude Code sessions, stored in Git branches

- HN: [48426297](https://news.ycombinator.com/item?id=48426297)
- Source: [github.com](https://github.com/ingram-technologies/claude-git-sessions)
- Score: 5
- Comments: 1
- Posted: 2026-06-06T16:04:45Z

## Translation

タイトル: HN を表示: Ccgs – 共同クロード コード セッション、Git ブランチに保存
記事のタイトル: GitHub - ingram-technologies/claude-git-sessions: 孤立した git ブランチを通じてチーム全体でクロード コード セッションを共有する (npm: claude-git-sessions) · GitHub
説明: 孤立した git ブランチを介してチーム全体でクロード コード セッションを共有する (npm: claude-git-sessions) - ingram-technologies/claude-git-sessions
HN テキスト: 私のチームは毎日クロード コードを使用しており、セッションは私たちが作成する最も有用な成果物の一部になっています。しかし、どのラップトップで発生しても、それらは ~/.claude/projects/ に閉じ込められます。同僚に「移行を解きほぐしたセッション」を渡して、同僚が再開して、中断したところから続行できるようにする良い方法はありません。
ccgs を入力します。既存のリポジトリのリモートの孤立したブランチ (@ccgs/<name>) を介してクロード コード セッションを共有します。セッション ファイルには作成者の絶対パスが含まれます。プル時に、ccgs は作業ディレクトリをパスに書き戻すため、再開は実際に機能します。トランスクリプトを喜んで破損する盲目的な検索と置換ではなく、構造的な cwd フィールドのみを外科的に編集します。 - すべては、使い捨てインデックスに対して git プラミング (hash-object/commit-tree/update-ref) を経由します。作業ツリー、インデックス、現在のブランチには決して触れず、ツリーが汚れていても大丈夫です。背後で何かをチェックアウトすることはありません。インストールせずに試してみるには: `npx claude-git-sessions`。これにより、ディレクトリを移動して、クロード コードのトランスクリプトを一緒に持ち運ぶこともできます (最初にプッシュし、次にディレクトリを移動し、次にプルするだけです)。 重要な注意: よほど優れたセキュリティ衛生を備えていない限り、クロード コードのセッションには環境の機密情報などの機密情報が含まれている可能性があります。慎重に使用し、パブリック リポジトリでの使用は避けてください。

トーリー。 ccgs によって使用されるブランチには「@ccgs/」という接頭辞が付けられるため、簡単にフィルタリングして除外できます。このプロジェクトは、Claude Code によって、Claude Code とともに書かれました。このショー HN はそうではありませんでした。 (URLを修正して再投稿しました)

記事本文:
GitHub - ingram-technologies/claude-git-sessions: 孤立した git ブランチを通じてチーム全体でクロード コード セッションを共有する (npm: claude-git-sessions) · GitHub
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
イングラムテクノロジーズ
/
クロード-git-セッション
公共
通知
サインインする必要があります

通知設定を変更するには
追加のナビゲーション オプション
コード
ingram-technologies/claude-git-sessions
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット src src test test .gitignore .gitignore README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
ccgs — クロード コード Git セッション
Claude Code セッションをチーム全体で共有する
既存の git リポジトリ内の孤立したブランチを介して。サーバーも追加も不要
インフラストラクチャ — トランスクリプトは、それが属するコードとともに使用されます。
npm で claude-git-sessions として公開されます (裸の名前 ccgs は npm によってブロックされます)
名前類似フィルター)。インストールするコマンドは ccgs です。
npx claude-git-sessions pull # 共有セッションをこのマシンに導入します
npx claude-git-sessions Push # このリポジトリのローカル セッションを公開します
npx claude-git-sessions 削除 < id |名前 >
pull 後、次のコマンドを使用してチームメイトのセッションをローカルで再開します。
クロード --resume < セッション ID >
インストール
npx を使用してオンデマンドで実行します (インストールなし):
npx クロード-git-sessions プル
…またはグローバルにインストールします (これにより、ccgs コマンドがインストールされます)。
npm i -g クロード-git-sessions
ccgs --ヘルプ
Node 20+ および git 2.5+ が必要です。 git リポジトリ内から実行します。
セッションは @ccgs/<name> という名前の孤立したブランチに保存されます (デフォルト)
@ccgs/default ) リポジトリのリモート (デフォルトのorigin) にあります。支店は共有しません
メインの履歴にソース ファイルは含まれず、セッション データのみが含まれます。
session/<session-id>.jsonl # 逐語的なトランスクリプト
session/<session-id>.meta.json # サイドカー メタデータ (名前、作成者、CWD など)
Memory/<fact>.md # 共有メモリのファクト (「ccgs メモリ」を参照)
Memory/<fact>.md.meta.json # サイドカー メタデータ (タイプ、作成者など)
ファイルは世界的にユニークな Clau によってキー付けされます

コードセッション UUID のトランスクリプト
異なる作者の作品が衝突することはありません。
すべての ccgs 操作は、低レベルの git プラミングで実行されます。
( hash-object / update-index / write-tree / commit-tree / Push ) に対する
プライベート一時インデックス。作業ツリー、インデックス、現在のブランチは次のとおりです。
決して触れたことはなく、汚れた木でもすべてが機能します。同時プッシュ
非早送り拒否の場合は、再フェッチして再生することによって処理されます。
@ccgs/<name> 名前空間を使用すると、個別のセッション セットを保持できます。を使用します。
チームの共有セッションのデフォルト名と、自分自身のプライベート名
まだ共有したくない進行中の作業:
ccgs Push -b my-wip # @ccgs/my-wip に公開します
ccgs pull -b my-wip # あなたのプライベートセットのみ
ccgs プッシュ # 共有 @ccgs/default セット
ブランチ名は、使用前に git check-ref-format で検証されます。
ccgs プル [--force] [--exclude-memory]
@ccgs/<name> を取得し、各セッションをリポジトリ ルートのプロジェクトに書き込みます
slug — ~/.claude/projects/<root-slug>/<id>.jsonl — つまり、claude --resume
リポジトリのルートで実行すると、プルされたすべてのセッションがリストされます。の構造的 CWD フィールド
各トランスクリプト行は作成者のパスからリポジトリのルートに書き換えられるため、
仕事を再開します。 (ツール出力内の絶対パスは表面的にそのまま残されます。)
セッションは、作成者が
クロードをサブディレクトリで起動しました。著者のサブディレクトリを尊重すると、
セッションは別の (おそらく存在しない) スラグ ディレクトリにあります。 ここで --resume
根元では見えません。元のサブディレクトリは依然として meta.json に記録されます。
デフォルトでは、これは共有メモリも取得します (ccgs メモリを参照)。パスする
--exclude-memory セッションのみ。ローカル セッションが
共有コピーは警告とともにスキップされます。 --force を渡して上書きします。もし
ブランチが存在しないため、出力されます

プルするものは何もなく、0 で終了します。
ccgs プッシュ [ターゲット...] [--exclude-memory]
作業ディレクトリがこのリポジトリ (またはサブディレクトリ) であるローカル セッションを検索します。
各トランスクリプトをそのまま孤立ブランチにコピーし、そのトランスクリプトを (再) 生成します。
メタ.json 。引数がない場合は、それらすべてをプッシュします。セッションID/名前を渡す
それらだけをプッシュします。デフォルトでは、これにより共有メモリもプッシュされます (プロジェクト /
参考事実）;セッションのみに --exclude-memory を渡します。孤児を生み出す
最初のプッシュで分岐します。レポートの追加と更新。
ccgs delete <id|name> [--yes] [--local]
完全な UUID、明確な UUID プレフィックス (git スタイル、≥4) によってターゲットを解決します。
chars)、または一意の表示名 - 候補をリストし、あいまいな場合は中止します。
削除される内容を表示し、y/N の確認を求めます (スキップするには --yes / -y
スクリプト用)。両方のファイルをブランチから削除してプッシュします。デフォルトのみ
共有ブランチがタッチされる。 --local を追加すると、ローカル コピーも削除されます。
ccgs メモリ プッシュ [--all] / ccgs メモリ プル [--all] [--force]
Claude コードはプロジェクトごとのメモリを ~/.claude/projects/<slug>/memory/ に保持します —
ファクトごとに 1 つの Markdown ファイルと MEMORY.md インデックス。 ccgs メモリはそれらを共有します
同じ孤立ブランチ上のファクト (メモリ/プレフィックスの下)。
プレーン ccgs プッシュ / ccgs プルには、デフォルトでメモリが既に組み込まれています (
以下にフィルターを入力します。つまり、個人的な事実はありません）。次の場合にこれらのメモリ サブコマンドを使用します。
メモリのみが必要な場合、または個人的な事実を含めるために --all が必要な場合。
記憶は混合感度であるため、事実の前面物質によってフィルタリングされます
タイプ:
プロジェクト / 参照事実 (「デプロイ方法」、「API ドキュメントの場所」)
はチームの知識であり、デフォルトで共有されます。
ユーザー/フィードバックの事実 (個人的な好み) は、あなたがしない限り保留されます。
--all を渡します。
pull は共有ファクトをローカル メモリ ディレクトリに書き込み、
内側にはっきりとマークされたブロック

それらを指す MEMORY.md — 独自のインデックス
行はそのまま残ります。 MEMORY.md 自体が BLOB としてプッシュされることはありません (2 つの
人々のインデックスが互いに破壊し合う可能性があります)。ファクトはファイル名によってキー化されます。
newer-wins 競合処理 (新しいローカル コピーを上書きする --force)。
スコープ: リポジトリのルート スラグのメモリ ディレクトリのみが同期されます (一般的なケース)。
サブディレクトリで開始されたセッションのメモリは保存されません。
オプション
デフォルト
意味
-b, --branch <名前>
デフォルト
セッションセット → ブランチ @ccgs/<名前>
--remote <リモート>
起源
使用する git リモート
-v、--バージョン
印刷版
-h、--ヘルプ
助けて
ステップ 0 の調査結果 — クロード コードがセッションを保存する方法
これらの仮定は、以前に実際の ~/.claude/ に対して経験的に検証されました。
ツールを書いています。これらはそれぞれ 1 か所に存在するため、問題が発生した場合は簡単に修正できます。
規約の変更。
プロジェクトのスラッグエンコーディング ( src/slug.ts )
Claude Code は、各プロジェクトのセッションを次の場所に保存します。
~/.claude/projects/<slug>/ 、ここでスラグは絶対値から派生します。
作業ディレクトリ。実際のディレクトリとディレクトリ内に記録された cwd を比較する
セッション ファイルには、次のルールが示されています。そうでない文字はすべて置き換えます。
[A-Za-z0-9] - 付き。
/ と . の両方に注意してください。 - にマップされるため、エンコードには非可逆性があり、
不可逆的。したがって、ccgs はパス→スラッグのみを実行し、実際のデータを保存します。
originalCwd を meta.json に個別に追加します。
構成ルートは CLAUDE_CONFIG_DIR を尊重し、 ~/.claude にフォールバックします。ローカル
セッションは <config>/projects/<slug>/<session-id>.jsonl に存在します。
セッション .jsonl スキーマ ( src/session.ts )
各セッションは JSONL です。1 行に 1 つの JSON オブジェクトがあり、多くのオブジェクト タイプが含まれます。
( user 、assistant 、attachment 、ai-title 、mode 、tool_result 、…)。
検証済みフィールド:
セッション UUID — sessionId 。ほとんどの行に存在し、ファイル名と同じです。
タイトル/名前 — { "type": "ai-title", "aiTitle": "…" } 行。

古い
使用されるバージョン { "type": "summary", "summary": "…" } ;いくつかのオブジェクトも運ぶ
タイトル。表示名は次のように解決されます: タイトル → 概要 → 最初のユーザー メッセージ
(切り捨て) → セッション ID。
作業ディレクトリ — cwd フィールド、user /assistant / に存在します
アタッチメントライン。これが唯一のフィールド プル リライトです。
タイムスタンプ — 行ごとのタイムスタンプ (ISO-8601)。 updatedAt は最大値です。
最初のユーザー メッセージ — 最初の type:"user" 行の message.content;
content は文字列またはコンテンツ ブロックの配列です (それぞれに .text が含まれる場合があります)。
既知の考慮事項/今後の課題
秘密の編集 - トランスクリプトはそのままプッシュされます。セッションが可能であれば、
秘密が含まれている場合は、プッシュする前にスクラブしてください。 --redact フラグが考えられます。
今後の追加。今日は実装されていません。
Git LFS — 非常に大きなトランスクリプトは通常の BLOB としてコミットされます。 LFS
取り扱いは後から追加される可能性があります。
npmインストール
npm run build # tsc -> dist/, chmod +x ビン
npm test #node:test 単体テスト (Claude のライブ インストールは必要ありません)
npm タイプチェックを実行する
テストでは、スラッグ エンコーディング、表示名の解決、
ID/プレフィックス/名前のマッチング、および cwd リマップ変換。
孤立した Git ブランチ (npm: claude-git-sessions) を介してチーム全体で Claude Code セッションを共有する
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
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

Share Claude Code sessions across a team through an orphan git branch (npm: claude-git-sessions) - ingram-technologies/claude-git-sessions

My team uses Claude Code daily, and the sessions have become some of the most useful artifacts we produce. But they're trapped in ~/.claude/projects/ on whichever laptop they happened on. There's no good way to hand a colleague "the session where I untangled the migration" so they can claude --resume it and keep going from where I left off.
Enter ccgs: Share Claude Code sessions through an orphan branch (@ccgs/<name>) in your existing repo's remote - Session files carry the author's absolute paths. On pull, ccgs rewrites the working dir back to your path so resume actually works — surgically editing only the structural cwd field, not a blind find-and-replace that would happily corrupt the transcript. - Everything goes through git plumbing (hash-object/commit-tree/update-ref) against a throwaway index. It never touches your working tree, index, or current branch, and it's fine with a dirty tree. It will not git checkout something behind your back. To try it without installing: `npx claude-git-sessions`. This also incidentally allows you to move a directory and carry the claude code transcripts with it (just push first, then move the directory, then pull) IMPORTANT CAVEAT: Unless you have a very good security hygiene, your Claude Code sessions are likely full of sensitive information such as environment secrets. Use with caution and avoid using on public repositories. Branches used by ccgs are prefixed by `@ccgs/` so you can easily filter them out. This project was written by and with Claude Code. This Show HN was not. (Reposted with URL fixed)

GitHub - ingram-technologies/claude-git-sessions: Share Claude Code sessions across a team through an orphan git branch (npm: claude-git-sessions) · GitHub
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
ingram-technologies
/
claude-git-sessions
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
ingram-technologies/claude-git-sessions
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits src src test test .gitignore .gitignore README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
ccgs — Claude Code Git Sessions
Share Claude Code sessions across a team
through an orphan branch in your existing git repo . No server, no extra
infrastructure — your transcripts ride along with the code they belong to.
Published on npm as claude-git-sessions (the bare name ccgs is blocked by npm's
name-similarity filter). The command it installs is ccgs .
npx claude-git-sessions pull # bring shared sessions onto this machine
npx claude-git-sessions push # publish your local sessions for this repo
npx claude-git-sessions delete < id | name >
After pull , resume any teammate's session locally with:
claude --resume < session-id >
Install
Run on demand with npx (no install):
npx claude-git-sessions pull
…or install globally (this installs the ccgs command):
npm i -g claude-git-sessions
ccgs --help
Requires Node 20+ and git 2.5+ . Run it from inside your git repo.
Sessions are stored on an orphan branch named @ccgs/<name> (default
@ccgs/default ) on your repo's remote (default origin ). The branch shares no
history with main and never contains your source files — only session data:
sessions/<session-id>.jsonl # the transcript, verbatim
sessions/<session-id>.meta.json # sidecar metadata (name, author, cwd, …)
memory/<fact>.md # shared memory facts (see `ccgs memory`)
memory/<fact>.md.meta.json # sidecar metadata (type, author, …)
Files are keyed by the globally-unique Claude Code session UUID, so transcripts
from different authors never collide.
Every ccgs operation is done with low-level git plumbing
( hash-object / update-index / write-tree / commit-tree / push ) against a
private temporary index. Your working tree, index, and current branch are
never touched , and everything works even with a dirty tree. Concurrent pushes
are handled by re-fetching and replaying on a non-fast-forward rejection.
The @ccgs/<name> namespace lets you keep separate session sets. Use the
default for the team's shared sessions, and a private name for your own
work-in-progress that you don't want to share yet:
ccgs push -b my-wip # publish to @ccgs/my-wip
ccgs pull -b my-wip # only your private set
ccgs push # the shared @ccgs/default set
Branch names are validated with git check-ref-format before use.
ccgs pull [--force] [--exclude-memory]
Fetches @ccgs/<name> and writes each session into the repo root's project
slug — ~/.claude/projects/<root-slug>/<id>.jsonl — so that claude --resume
run at the repo root lists every pulled session. The structural cwd field in
each transcript line is rewritten from the author's path to your repo root so
resume works. (Absolute paths inside tool output are left as-is — cosmetic.)
Sessions are always placed under the repo-root slug, even if the author
launched Claude in a subdirectory. Honoring the author's subdir would put the
session in a separate (and possibly non-existent) slug dir where --resume
at the root can't see it. The original subdir is still recorded in meta.json .
By default this also pulls shared memory (see ccgs memory ); pass
--exclude-memory for sessions only. If a local session is newer than the
shared copy it is skipped with a warning; pass --force to overwrite. If the
branch doesn't exist, it prints nothing to pull and exits 0.
ccgs push [targets...] [--exclude-memory]
Finds local sessions whose working directory is this repo (or a subdirectory),
copies each transcript verbatim onto the orphan branch, and (re)generates its
meta.json . With no arguments it pushes all of them; pass session ids/names to
push only those. By default this also pushes shared memory (project /
reference facts); pass --exclude-memory for sessions only. Creates the orphan
branch on first push. Reports added vs updated.
ccgs delete <id|name> [--yes] [--local]
Resolves the target by full UUID, unambiguous UUID prefix (git-style, ≥4
chars), or unique display name — listing candidates and aborting if ambiguous.
Shows what will be deleted and asks for y/N confirmation ( --yes / -y to skip
for scripting). Removes both files from the branch and pushes. By default only
the shared branch is touched; add --local to also remove the local copy.
ccgs memory push [--all] / ccgs memory pull [--all] [--force]
Claude Code keeps per-project memory in ~/.claude/projects/<slug>/memory/ —
one Markdown file per fact, plus a MEMORY.md index. ccgs memory shares those
facts on the same orphan branch (under a memory/ prefix).
Plain ccgs push / ccgs pull already include memory by default (with the
type filter below, i.e. no personal facts). Use these memory subcommands when
you want memory only , or need --all to include personal facts.
Memory is mixed-sensitivity , so it's filtered by the fact's frontmatter
type :
project / reference facts ("how we deploy", "where the API docs live")
are team knowledge and are shared by default .
user / feedback facts (personal preferences) are held back unless you
pass --all .
pull writes the shared facts into your local memory dir and maintains a
clearly-marked block inside your MEMORY.md pointing at them — your own index
lines are left untouched . MEMORY.md itself is never pushed as a blob (two
people's indexes would clobber each other). Facts are keyed by filename with
newer-wins conflict handling ( --force to overwrite a newer local copy).
Scope: only the repo root slug's memory dir is synced (the common case);
memory from sessions started in subdirectories is not.
Option
Default
Meaning
-b, --branch <name>
default
session set → branch @ccgs/<name>
--remote <remote>
origin
git remote to use
-v, --version
print version
-h, --help
help
Step 0 findings — how Claude Code stores sessions
These assumptions were verified empirically against a real ~/.claude/ before
writing the tool. They live in one place each so they're easy to fix if the
convention changes.
Project slug encoding ( src/slug.ts )
Claude Code stores each project's sessions under
~/.claude/projects/<slug>/ , where the slug is derived from the absolute
working directory. Comparing real directories to the cwd recorded inside their
session files showed the rule is: replace every character that is not
[A-Za-z0-9] with - .
Note that both / and . map to - , so the encoding is lossy and
irreversible . ccgs therefore only ever goes path → slug, and stores the real
originalCwd separately in meta.json .
The config root honors CLAUDE_CONFIG_DIR and falls back to ~/.claude . Local
sessions live at <config>/projects/<slug>/<session-id>.jsonl .
Session .jsonl schema ( src/session.ts )
Each session is JSONL — one JSON object per line, with many object type s
( user , assistant , attachment , ai-title , mode , tool_result , …).
Verified fields:
Session UUID — sessionId , present on most lines, equal to the filename.
Title / name — a { "type": "ai-title", "aiTitle": "…" } line. Older
versions used { "type": "summary", "summary": "…" } ; some objects also carry
a title . The display name resolves: title → summary → first user message
(truncated) → session id.
Working directory — the cwd field, present on user / assistant /
attachment lines. This is the only field pull rewrites.
Timestamps — per-line timestamp (ISO-8601). updatedAt is the max.
First user message — message.content on the first type:"user" line;
content is a string or an array of content blocks (each may have .text ).
Known considerations / future work
Secret redaction — transcripts are pushed verbatim. If your sessions may
contain secrets, scrub them before pushing. A --redact flag is a likely
future addition. Not implemented today.
Git LFS — very large transcripts are committed as ordinary blobs. LFS
handling may be added later.
npm install
npm run build # tsc -> dist/, chmod +x the bin
npm test # node:test unit tests (no live Claude install needed)
npm run typecheck
Tests cover the pure pieces: slug encoding, display-name resolution,
id/prefix/name matching, and the cwd remap transform.
Share Claude Code sessions across a team through an orphan git branch (npm: claude-git-sessions)
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
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
