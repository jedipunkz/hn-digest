---
source: "https://github.com/joshduffy/claude-handoff-guard"
hn_url: "https://news.ycombinator.com/item?id=48325871"
title: "Show HN: Stop parallel AI coding sessions clobbering each other's handoffs"
article_title: "GitHub - joshduffy/claude-handoff-guard: Hook-enforced ownership for AI coding session handoffs · GitHub"
author: "nahsuhn"
captured_at: "2026-05-30T11:42:50Z"
capture_tool: "hn-digest"
hn_id: 48325871
score: 2
comments: 0
posted_at: "2026-05-29T16:54:22Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Stop parallel AI coding sessions clobbering each other's handoffs

- HN: [48325871](https://news.ycombinator.com/item?id=48325871)
- Source: [github.com](https://github.com/joshduffy/claude-handoff-guard)
- Score: 2
- Comments: 0
- Posted: 2026-05-29T16:54:22Z

## Translation

タイトル: Show HN: 互いのハンドオフを妨害する並列 AI コーディング セッションを停止する
記事のタイトル: GitHub - joshduffy/claude-handoff-guard: AI コーディング セッション ハンドオフのフック強制所有権 · GitHub
説明: AI コーディング セッション ハンドオフのフック強制所有権 - joshduffy/claude-handoff-guard

記事本文:
GitHub - joshduffy/claude-handoff-guard: AI コーディング セッション ハンドオフのフック強制所有権 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
ジョシュダフィー
/
クロード・ハンドオフ・ガード
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
ジョシュダフィ/クロード・ハンドオフ・ガード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット フック フック ルール ルール スクリプト スクリプト スキル/ハンドオフ スキル/ハンドオフ .gitignore .gitignore ライセンス ライセンス README.md README.md package.json package.json settings.example.json settings.example.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
AI コーディング セッションのハンドオフに対するフックによる所有権の強制。
ほとんどの「ハンドオフ」ツールは記憶喪失を解決します。状態をマークダウン ファイルにキャプチャし、圧縮後に復元します。
または新しいセッション。その問題はしっかりカバーされています。これは誰も強制していない問題を解決します。
同時クロバー。 2 つのセッションが同じリポジトリで作業する場合、または 2 番目のマシンで再開する場合、または
バックグラウンド エージェントはインタラクティブ エージェントと並行して実行され、お互いのハンドオフ メモを上書きします。
必要なコンテキストがなくなるまでは調べないでください。
ここでの修正はより良いテンプレートではありません。これは、クロスセッションの上書きを行う PreToolUse フックです。
単なる落胆ではなく、構造的にブロックされています。
中心となるアイデア: ロックはファイル内に存在します
すべてのハンドオフ ファイルの最初の行は所有権マーカーです。
<!-- クロードセッション: 9e0d3802-... -->
サイドカー .lock ファイルがありません。所有権は、Git を介してデバイス間でアーティファクトとともに移動します。
MVを通して。 PreToolUse フックは、呼び出しセッションの ID を読み取り、それを
コンテンツが書き込まれており、マーカーはすでにディスク上にあります。不一致により書き込みがブロックされます。
卵が先か鶏が先か (興味深い部分)
誰が書いたかを示すハンドオフを書き込むには、セッションはそれ自体の ID を知っている必要があります。そうではありません。の
モデルには、その session_id へのネイティブ アクセスがありません。
したがって、新しいハンドオフの最初の書き込みは失敗するように設計されています。ブロック理由には次の ID が含まれます。
ハンドオフ書き込みが見つからないか、所有者が間違っています

船のマーカー。
あなたの session_id: `9e0d3802-4f...`。
これを 1 行目に正確に追加します。
<!-- クロードセッション: 9e0d3802-4f... -->
その後、再試行してください。
モデルは失敗から ID をコピーし、再試行します。新しいハンドオフごとに 1 ブロック、ファイルは次のようになります。
今後のすべてのセッションで自己識別されます。不足している機能は 1 回限りのハンドシェイクになります。
3 つの突然変異表面すべてにわたる防御
書き込みがブロックされているモデルは、あなたの周りをルーティングします。したがって、ガードはファイルのあらゆる方法をカバーします。
変異した:
Write は、新しいコンテンツ内のマーカーを検証します。
編集では、ディスク上のマーカーが検証されます (また、マーカーのないレガシー ファイルへの編集は、編集が行われるまでブロックされます)
書き込みによる所有権）。
Bash はシェルのリダイレクトをハンドオフ パス ( > 、 >> 、 tee 、 sed -i ) に一致させ、所有されていないものをブロックします
ファイルツールをすり抜けようとする書き込み。
また、Claude Code ツール スキーマ (書き込み/編集/Bash) と Gemini CLI スキーマの両方も受け入れます。
( write_file / replace / run_shell_command ) を 1 つのフックに含めます。これは、フックをゲートするとサイレントに無効になるためです。
他のクライアントの警備員。
フック/handoff-write-guard.mjs PreToolUse: 所有権ガード
hooks/handoff-session-start.mjs SessionStart: 既存のハンドオフ + スラグのオーバーラップを表面化します
フック/handoff-stop-gate.mjs 停止: セッションごとに 1 回、「ハンドオフはありません」というナッジ
hooks/pre-commit-staged-marker-check.mjs git pre-commit: 2 つのセッションのハンドオフを混合したコミットをブロックする
hooks/test/handoff-write-guard.test.mjs ノード --テスト スイート (8 件)
scripts/handoff-merge-archive.mjs 古くなったマーカーのないレガシーハンドオフをアーカイブします
scripts/install-git-hooks.sh プリコミットフックのデバイスごとのインストーラー
skill/handoff/SKILL.md /handoff スラッシュ コマンド
rules/session-handoff.md フックが適用する規則
settings.example.json フック配線して ~/.claude/settings.json にマージします
標準のクロード コードではハンドオフが予想されます

メモリレイアウト:
~/.claude/projects/<encoded-cwd>/memory/handoff-<branch>-<topic>.md 、ここで <encoded-cwd> は
/ 、 \ 、および . を使用した絶対作業ディレクトリ。 - に置き換えられます。
# 1. フック/スキル/ルールを ~/.claude にコピーします
cp フック/ * .mjs ~ /.claude/hooks/
cp -r フック/テスト ~ /.claude/hooks/
cp scripts/ * ~ /.claude/scripts/
cp -r スキル/ハンドオフ ~ /.claude/skills/
cp ルール/ * ~ /.claude/rules/
# 2. settings.example.json を ~/.claude/settings.json にマージします (加算配列)
# 3. デバイスごとの git pre-commit フックをインストールします (ハンドオフは git リポジトリ内に存在します)。
bash ~ /.claude/scripts/install-git-hooks.sh
#4. 検証する
ノード --test ~ /.claude/hooks/test/ * .test.mjs
哲学: フェイルオープンし、決してセッションをトラップしない
すべてのフックは本体を try/catch でラップし、内部エラーが発生すると 0 で終了します。ガードのバグ
慣例に従わない。セッションがブリックされることはありません。停止ナッジは非ブロックであり、最大で発火します。
セッションごとに 1 回。エスケープハッチ ( touch /tmp/handoff-guard-bypass-<file> 、または
HANDOFF_GUARD_BYPASS=1 ) が存在するのは、オーバーライドできない構造的なガードが存在するためです。
構造的なガードを引き剥がします。バイパスの使用はログに記録されるため、サイレント無効化は監査可能です。
これにより、無意識に盗むことは不可能になります。選んだものが不可能になるわけではありません。
外部マーカーが正しく表示されたセッションでも、ファイルをアーカイブしたり、バイパス環境を設定したりできます。人間
目に見えるブロック メッセージを確認することが、そのクラスのバックストップとなります。
ファイル名の <branch> トークンは実際のブランチと照合されません (意図的: クロスデバイス)
履歴書は意図的に外国支部のトピックを継承します)。
2 つのセッションがフック間のギャップに同じ新しいファイル名を作成した場合、TOCTOU 競合が発生します。
読み取りとツールの書き込み。ソロ開発者にとってはほとんど珍しいことです。意図的にロックされていません。
AI コーディング セッション h のフック強制所有権

アンドフス
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

Hook-enforced ownership for AI coding session handoffs - joshduffy/claude-handoff-guard

GitHub - joshduffy/claude-handoff-guard: Hook-enforced ownership for AI coding session handoffs · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
joshduffy
/
claude-handoff-guard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
joshduffy/claude-handoff-guard
master Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits hooks hooks rules rules scripts scripts skills/ handoff skills/ handoff .gitignore .gitignore LICENSE LICENSE README.md README.md package.json package.json settings.example.json settings.example.json View all files Repository files navigation
Hook-enforced ownership for AI coding session handoffs.
Most "handoff" tools solve amnesia : capture state to a markdown file, restore it after compaction
or a new session. That problem is well covered. This one solves the problem nobody enforces:
concurrent clobber . When two sessions work the same repo, or you resume on a second machine, or a
background agent runs alongside an interactive one, they overwrite each other's handoff notes and you
do not find out until the context you needed is gone.
The fix here is not a better template. It is a PreToolUse hook that makes a cross-session overwrite
structurally blocked, not merely discouraged.
The core idea: the lock lives inside the file
Every handoff file's first line is an ownership marker:
<!-- claude-session: 9e0d3802-... -->
There is no sidecar .lock file. Ownership travels with the artifact through git, across devices,
through a mv . A PreToolUse hook reads the calling session's id and compares it to the marker in the
content being written and the marker already on disk. Mismatch blocks the write.
The chicken-and-egg (the interesting part)
To write a handoff that says who wrote it, the session needs to know its own id. It does not. The
model has no native access to its session_id .
So the first write of a fresh handoff is designed to fail . The block reason carries the id:
Handoff write missing or wrong ownership marker.
Your session_id: `9e0d3802-4f...`.
Prepend exactly this as line 1:
<!-- claude-session: 9e0d3802-4f... -->
Then retry.
The model copies the id from the failure and retries. One block per fresh handoff, and the file is
now self-identifying for every future session. The missing capability becomes a one-time handshake.
Defense across all three mutation surfaces
A model that is blocked on Write will route around you. So the guard covers every way a file can be
mutated:
Write validates the marker in the new content.
Edit validates the marker on disk (and blocks edits to legacy marker-less files until you take
ownership with a Write).
Bash matches shell redirects to handoff paths ( > , >> , tee , sed -i ) and blocks unowned
writes that try to sneak past the file tools.
It also accepts both the Claude Code tool schema ( Write / Edit / Bash ) and the Gemini CLI schema
( write_file / replace / run_shell_command ) in one hook, because gating on one silently disables the
guard for the other client.
hooks/handoff-write-guard.mjs PreToolUse: the ownership guard
hooks/handoff-session-start.mjs SessionStart: surface existing handoffs + slug overlaps
hooks/handoff-stop-gate.mjs Stop: once-per-session "you have no handoff" nudge
hooks/pre-commit-staged-marker-check.mjs git pre-commit: block commits mixing two sessions' handoffs
hooks/test/handoff-write-guard.test.mjs node --test suite (8 cases)
scripts/handoff-migrate-archive.mjs archive stale, marker-less legacy handoffs
scripts/install-git-hooks.sh per-device installer for the pre-commit hook
skills/handoff/SKILL.md the /handoff slash command
rules/session-handoff.md the convention the hooks enforce
settings.example.json hook wiring to merge into ~/.claude/settings.json
Handoffs are expected under the standard Claude Code memory layout:
~/.claude/projects/<encoded-cwd>/memory/handoff-<branch>-<topic>.md , where <encoded-cwd> is the
absolute working directory with / , \ , and . replaced by - .
# 1. Copy hooks/skills/rules into your ~/.claude
cp hooks/ * .mjs ~ /.claude/hooks/
cp -r hooks/test ~ /.claude/hooks/
cp scripts/ * ~ /.claude/scripts/
cp -r skills/handoff ~ /.claude/skills/
cp rules/ * ~ /.claude/rules/
# 2. Merge settings.example.json into ~/.claude/settings.json (additive arrays)
# 3. Install the per-device git pre-commit hook (handoffs live in a git repo)
bash ~ /.claude/scripts/install-git-hooks.sh
# 4. Verify
node --test ~ /.claude/hooks/test/ * .test.mjs
Philosophy: fail open, never trap the session
Every hook wraps its body in try/catch and exits 0 on any internal error. A bug in the guard
degrades to convention; it never bricks a session. The Stop nudge is non-blocking and fires at most
once per session. The escape hatches ( touch /tmp/handoff-guard-bypass-<file> , or
HANDOFF_GUARD_BYPASS=1 ) exist precisely because a structural guard you cannot override becomes a
structural guard you rip out. Bypass use is logged so silent disabling is auditable.
This makes the unaware clobber impossible. It does not make the chosen one impossible.
A session correctly shown a foreign marker can still archive the file or set the bypass env. Human
review of the visible block message is the backstop for that class.
The <branch> token in a filename is not checked against the real branch (intentional: cross-device
resume deliberately inherits a foreign branch's topic).
A TOCTOU race exists if two sessions create the same new filename in the gap between the hook's
read and the tool's write. Vanishingly rare for solo dev; deliberately not locked.
Hook-enforced ownership for AI coding session handoffs
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
