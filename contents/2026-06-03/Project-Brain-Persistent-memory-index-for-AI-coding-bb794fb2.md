---
source: "https://github.com/OoneBreath/claude-code-project-brain"
hn_url: "https://news.ycombinator.com/item?id=48377526"
title: "Project Brain – Persistent memory index for AI coding"
article_title: "GitHub - OoneBreath/claude-code-project-brain: Persistent, navigable memory for Claude Code — stop re-explaining your projects to the AI every session. · GitHub"
author: "Slav_fixflex"
captured_at: "2026-06-03T00:34:36Z"
capture_tool: "hn-digest"
hn_id: 48377526
score: 1
comments: 0
posted_at: "2026-06-02T23:12:59Z"
tags:
  - hacker-news
  - translated
---

# Project Brain – Persistent memory index for AI coding

- HN: [48377526](https://news.ycombinator.com/item?id=48377526)
- Source: [github.com](https://github.com/OoneBreath/claude-code-project-brain)
- Score: 1
- Comments: 0
- Posted: 2026-06-02T23:12:59Z

## Translation

タイトル: Project Brain – AI コーディング用の永続メモリ インデックス
記事のタイトル: GitHub - OoneBreath/claude-code-project-brain: クロード コードの永続的でナビゲート可能なメモリ — セッションごとにプロジェクトを AI に再説明するのはやめましょう。 · GitHub
説明: クロード コードの永続的でナビゲート可能なメモリ — セッションごとにプロジェクトを AI に再説明する必要はありません。 - OoneBreath/claude-code-project-brain

記事本文:
GitHub - OoneBreath/claude-code-project-brain: クロード コード用の永続的でナビゲート可能なメモリ — セッションごとにプロジェクトを AI に再説明する必要はありません。 · GitHub
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
ワンブレス
/
クロードコードプロジェクトブレイン
公共
通知
あなたは署名している必要があります

通知設定を変更するには
追加のナビゲーション オプション
コード
OoneBreath/クロードコードプロジェクトブレイン
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
10 コミット 10 コミット docs docs example/ example-saas/ .project-brain 例/ example-saas/ .project-brain スキル/ project-brain スキル/ project-brain .gitignore .gitignore ライセンス ライセンス README.md README.md install.sh install.sh すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Project Brain — クロード コードの永続メモリ
セッションごとに AI にプロジェクトを再説明するのはやめましょう。
Project Brain は、クロードに
プロジェクトの小さなナビゲート可能なマップ — プロジェクトのスタック、決定事項、落とし穴、およびすでに行われているもの
完了したことを確認することで、忘れたり、プロジェクトを混同したり、再読したりすることがなくなります。
1000 行の README ですべてのタスクのコンテキストを説明します。
これはデータベース、サーバー、または別の AI ラッパーではありません。これは慣例とスキルです。
クロードがインデックスを介して読み込むプレーンなマークダウンの .project-brain/ フォルダー
実際に必要な場合にのみ詳細を説明します。
すべての新しいクロード コード セッション、すべてのプロジェクト:
…またか。そして、数時間の作業の後、モデルはプロジェクト間で詳細を混合し始めます。
これは FastAPI + Postgres、あちらは Node + tRPC + MySQL — そして静かに物事をやり直す
あなたはすでに先週終わっています。
.プロジェクトブレイン/
Index.md # 小さな MAP: プロジェクト → トピック → ステータス + ポインター
プロジェクト/
<プロジェクト>/
<topic>.md # 詳細、必要な場合のみ読み取ります
クロードは最初に小さなインデックスを読みます。 「キャッシュの問題をどのように解決しましたか?」と尋ねると、それ
ナレッジ ベース全体ではなく、1 つのファイルへの 1 つのポインターに従います。 「交換してください」と頼むと、
ロゴ」と地図には、それが 3 日前に実行され検証されたことが示されており、それが表示され、質問されます。
繰り返したいのか、やりたいのか

何か新しいこと。
すべてのセッション: アーキテクチャ、デプロイメント、スタック、歴史、落とし穴を再説明します。
クロードはすでに、スタック、何が行われ、何が機能し、何が失敗し、何が進行中であるかを知っています。
そしてあなたの質問に関連する 1 つのトピックだけを読みます。
これを使用する理由について自分に正直になってください。
適用される場合はトークンが少なくなります。現在、すべてを常にロードされた巨大なファイルに保存している場合
doc を小さなインデックスとオンデマンドのトピック ファイルに分割すると、セッションごとのコンテキストが完全に削減されます。
そうしないと、トークンの獲得額はわずかになります。
幻覚が少なくなる。地図はアンカーです。モデルがデプロイメントの考案を停止するか、
あるプロジェクトのスタックを別のプロジェクトのスタックと交換する。
数ヶ月の記憶。 3 か月後にプロジェクトに戻っても、クロードはまだその方法を理解しています。
1 キロメートルもの README を貼り付けなくても機能します。
フラット ノート ファイルより優れている点が 2 つあります
ステータスには、単に「完了」だけでなく結果も含まれます。 ✓ 検証済み vs ✗ 失敗 vs ⚠ 進行中。
モデルは、「完了して機能する」と「試してみたらダメだった」の違いを認識しています。
上書きではなくバージョン管理。アプローチが交換されると、古いアプローチがそのまま保持されます。
置き換えられたメモ — したがって、何が試みられ、なぜ変更されたのかの痕跡は残ります。
git clone https://github.com/OoneBreath/claude-code-project-brain.git
cd クロード コード プロジェクト ブレイン
./install.sh # スキルを ~/.claude/skills/ にコピーします (各マシンで実行)
インストール後に新しい Claude Code セッションを開始します。スキルはセッション開始時にロードされるため、
すでに開いているセッションにはスキルは表示されません。
次に、ワークスペース内のセッションで次のようにします。
/project-brain → "init" .project-brain/ をセットアップし、プロジェクトを検出します
/project-brain → 「X をどのように解決しましたか?」インデックスを通じて思い出す
init をワークスペースごとに 1 回実行します。スキルはマシンごとにインストールされます ( ~/.claude/skills/ )。

e
メモリ ( .project-brain/ ) はプロジェクトごとに存在します。そのため、複数のリポジトリをホストしているサーバーでは、init をポイントします。
ワークスペース ルートを作成すると、それらすべてが 1 つの頭脳内でカタログ化されます。
その後はほとんど考えなくなります。クロードは最初に地図を読み、事前に確認します。
作業をやり直し、作業単位が完了すると更新します。
初回実行時のメモ。 init は設計上軽量です。プロジェクトを次から検出します。
package.json / pyproject.toml / git にアクセスし、小さなインデックスを書き込みます。それはあなたのものを読みません
ソースなので安いです。その後、実際に仕事をするにつれて、脳は徐々に満たされていきます。
クロードに既存のコードベースを読み取って脳に事前入力させたい場合 (「ディープ バックフィル」)、
これは 1 回限りの前払いトークンコストです。クロードはコードとドキュメントを読んで要約する必要があります。
これを投資と考えてください。組織化するためにトークンを一度支払うだけで、その後のセッションはすべて
より安く、よりシャープに。一度に 1 つのプロジェクトに範囲を設定します。スキルは実行前に警告を表示します。
結果は、プロジェクトが実際に文書化したもの、つまり充実した README/CHANGELOG を備えたリポジトリを反映しています。
埋め戻しは裸の埋め戻しよりも豊かになります。そこにあるものを要約します。文脈を発明するものではありません。だから期待してください
深さはプロジェクトによって異なり、固定された結果ではありません。
スキルは ~/.claude/skills/project-brain/ にあります (個人の範囲、マシンごと)。
init はワークスペースに .project-brain/ を作成し、プロジェクト (git repos、package.json、
pyproject.toml , …)、小さなポインタを CLAUDE.md にドロップします。これにより、今後のセッションで
まずは地図。
1 つの頭脳で、1 つのサーバーまたは単一のリポジトリ上に多数のプロジェクトをカタログ化できます。
すべては自分で読み、編集し、コミットできるプレーンなマークダウンです。
バンドルされた Brain-Check バリデータ ( python3 ~/.claude/skills/project-brain/brain-check ) がキャッチします
壊れたポインター、不正な形式のフロントマッター、インデックス↔トピックのステータスのドリフト - 後で実行してください。

大きな変化。
時間が経っても膨らみません。トピック ファイルはコールド ストレージであり、インデックスのみがロードされます。
したがって、セッションあたりのコストは、保持する履歴の量ではなく、インデックスによって制限されます。
セッションごとに重くなるフラットなnotes.mdとは異なり、脳は軽いままです。死者を整理するには
トピックをアーカイブすると (インデックスからその 1 行を削除し、ファイルは保持します)、自動削除されるものは何もありません。
あなたの脳はあなたのものであり、デフォルトでは非公開になっています。本物の .project-brain/ は最終的に次のようになります。
インフラの詳細 (DB 名、ポート、サーバー パス、ホスト名)。プロジェクトごとにコミットするかどうかを決定する
(リポジトリとともに移動し、チームと共有します) またはバージョン管理の対象外に保ちます。このリポジトリは出荷されます
.project-brain/ を無視する .gitignore 。これにより、スキル自身のリポジトリが誤って作成されることがなくなります。
本物の脳を持っています。
Project Brain は、複数の独立した SaaS 製品を同時に実行することから生まれました。
Sentinel AI (サーバー セキュリティとデータベースの自動操縦) と
24ad.info (AI 支援広告プラットフォーム)、マルチサーバーと並行
フリートインテリジェンス バックエンド、スパム対策サービス、コンテンツ ツール。すべてのプロジェクトが実行するとき
何千行ものコンテキスト、AI による忘却のコスト、または 2 つのプロジェクトを静かに混ぜ合わせる
— それは本物です。これは、それらをまっすぐに保つ作業記憶です。このパターンは理論的なものではありません。
クロード コードの永続的でナビゲート可能なメモリ — セッションごとにプロジェクトを AI に再説明する必要はありません。
fixflex.co.uk/project-brain.html
トピックス
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

Persistent, navigable memory for Claude Code — stop re-explaining your projects to the AI every session. - OoneBreath/claude-code-project-brain

GitHub - OoneBreath/claude-code-project-brain: Persistent, navigable memory for Claude Code — stop re-explaining your projects to the AI every session. · GitHub
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
OoneBreath
/
claude-code-project-brain
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
OoneBreath/claude-code-project-brain
main Branches Tags Go to file Code Open more actions menu Folders and files
10 Commits 10 Commits docs docs examples/ example-saas/ .project-brain examples/ example-saas/ .project-brain skills/ project-brain skills/ project-brain .gitignore .gitignore LICENSE LICENSE README.md README.md install.sh install.sh View all files Repository files navigation
Project Brain — persistent memory for Claude Code
Stop re-explaining your projects to the AI every session.
Project Brain is a Claude Code skill that gives Claude a
small, navigable map of your projects — their stack, decisions, pitfalls, and what's already
been done — so it stops forgetting, stops mixing projects up, and stops re-reading a
1000-line README into context on every single task.
It is not a database, a server, or another AI wrapper. It's a convention plus a skill:
a .project-brain/ folder of plain markdown that Claude reads through an index, loading
detail only when it's actually needed.
Every new Claude Code session, on every project:
…again. And after a few hours of work the model starts mixing details between projects —
this one is FastAPI + Postgres, that one is Node + tRPC + MySQL — and quietly redoing things
you already finished last week.
.project-brain/
index.md # a small MAP: projects → topics → status + pointer
projects/
<project>/
<topic>.md # the detail, read only when needed
Claude reads the small index first. When you ask "how did we solve the cache issue?" it
follows one pointer to one file — not the whole knowledge base. When you ask it to "swap the
logo" and the map says that was done and verified three days ago, it tells you and asks
whether you want to repeat it or do something new.
Every session: re-explain architecture, deployment, stack, history, pitfalls.
Claude already knows: the stack, what was done, what worked, what failed, what's in progress —
and reads only the one topic relevant to your question.
Be honest with yourself about why you'd use this:
Fewer tokens — when it applies. If you currently keep everything in a giant always-loaded
doc, splitting into a small index + on-demand topic files genuinely cuts per-session context.
If you don't, the token win is modest.
Fewer hallucinations. The map is an anchor. The model stops inventing your deployment or
swapping one project's stack for another's.
Multi-month memory. Come back to a project after three months and Claude still knows how it
works — without you pasting a kilometre of README.
Two things make it better than a flat notes file
Status carries the outcome, not just "done": ✓ verified vs ✗ failed vs ⚠ in-progress .
The model knows the difference between "done and works" and "we tried that and it broke."
Versioning, not overwriting. When an approach is replaced, the old one is kept as a
superseded note — so the trail of what was tried and why it changed survives.
git clone https://github.com/OoneBreath/claude-code-project-brain.git
cd claude-code-project-brain
./install.sh # copies the skill into ~/.claude/skills/ (run on each machine)
Start a new Claude Code session after installing — skills are loaded at session start, so the
skill won't show up in a session that was already open.
Then, in a session inside your workspace:
/project-brain → "init" set up .project-brain/ and detect your projects
/project-brain → "how did we solve X?" recall through the index
Run init once per workspace . The skill is installed per machine ( ~/.claude/skills/ ), but the
memory ( .project-brain/ ) lives per project — so on a server hosting several repos, point init at
the workspace root and it catalogs them all in one brain.
After that you mostly don't think about it: Claude reads the map at the start, checks it before
redoing work, and updates it when a unit of work is done.
A note on the first run. init is light by design — it detects your projects from
package.json / pyproject.toml / git and writes a small index. It does not read your
source, so it's cheap. The brain then fills in gradually as you actually work.
If you want Claude to pre-populate the brain by reading an existing codebase ("deep backfill"),
that's a one-time upfront token cost — Claude has to read your code and docs to summarise them.
Think of it as an investment: you pay tokens once to get organised, and every later session is
cheaper and sharper. Scope it to one project at a time. The skill will warn you before doing it.
The result reflects what your project actually documents — a repo with a solid README/CHANGELOG
backfills richer than a bare one. It summarises what's there; it doesn't invent context. So expect
the depth to vary by project, not a fixed result.
The skill lives in ~/.claude/skills/project-brain/ (personal scope, per machine).
init creates .project-brain/ in your workspace, detects projects (git repos, package.json ,
pyproject.toml , …), and drops a tiny pointer into your CLAUDE.md so future sessions read the
map first.
One brain can catalog many projects on one server or just a single repo.
Everything is plain markdown you can read, edit, and commit yourself.
A bundled brain-check validator ( python3 ~/.claude/skills/project-brain/brain-check ) catches
broken pointers, malformed frontmatter, and index↔topic status drift — run it after big changes.
It doesn't bloat over time. Topic files are cold storage — only the index is ever loaded
eagerly, so the per-session cost is bounded by the index, not by how much history you keep.
Unlike a flat notes.md that gets heavier every session, the brain stays light. To tidy a dead
topic you archive it (drop its one line from the index, keep the file) — nothing auto-deletes.
Your brain is yours — and it's private by default. A real .project-brain/ ends up holding
infra details (DB names, ports, server paths, hostnames). Decide per project whether to commit it
(travels with the repo, shared with your team) or keep it out of version control. This repo ships
a .gitignore that ignores .project-brain/ precisely so the skill's own repo never accidentally
carries a real brain.
Project Brain came out of running several independent SaaS products at once — among them
Sentinel AI (server security & database autopilot) and
24ad.info (an AI-assisted classifieds platform), alongside a multi-server
fleet-intelligence backend, an anti-spam service and a content tool. When every project carries
thousands of lines of context, the cost of the AI forgetting — or quietly mixing two projects up
— is real. This is the working memory that keeps them straight. The pattern isn't theoretical.
Persistent, navigable memory for Claude Code — stop re-explaining your projects to the AI every session.
fixflex.co.uk/project-brain.html
Topics
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
