---
source: "https://github.com/responsiblparty/claude-dmg-skill"
hn_url: "https://news.ycombinator.com/item?id=48483589"
title: "/dmg – a Claude Code skill for persistent memory and session sync"
article_title: "GitHub - responsiblparty/claude-dmg-skill: Claude Code /dmg skill — persistent memory + session sync (git, docs, memory in one command) · GitHub"
author: "responsiblparty"
captured_at: "2026-06-11T00:59:58Z"
capture_tool: "hn-digest"
hn_id: 48483589
score: 4
comments: 1
posted_at: "2026-06-10T22:24:30Z"
tags:
  - hacker-news
  - translated
---

# /dmg – a Claude Code skill for persistent memory and session sync

- HN: [48483589](https://news.ycombinator.com/item?id=48483589)
- Source: [github.com](https://github.com/responsiblparty/claude-dmg-skill)
- Score: 4
- Comments: 1
- Posted: 2026-06-10T22:24:30Z

## Translation

タイトル: /dmg – 永続メモリとセッション同期のためのクロード コード スキル
記事のタイトル: GitHub - responsiblparty/claude-dmg-skill: Claude Code /dmg skill — 永続メモリ + セッション同期 (git、docs、1 つのコマンドでのメモリ) · GitHub
説明: Claude Code /dmg スキル — 永続メモリ + セッション同期 (git、docs、1 つのコマンドでのメモリ) - responsiblparty/claude-dmg-skill

記事本文:
GitHub - responsiblparty/claude-dmg-skill: クロード コード /dmg スキル — 永続メモリ + セッション同期 (git、docs、1 つのコマンドでのメモリ) · GitHub
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
責任者
/
クロード-ダメージ-スキル
公共
通知
通知を変更するにはサインインする必要があります

アイケーション設定
追加のナビゲーション オプション
コード
責任者/クロード-dmg-スキル
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミット メモリ メモリ スキル/ dmg スキル/ dmg CLAUDE.md CLAUDE.md README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
/dmg — 永続メモリ + 自動同期用のクロード コード スキル
クロード コードのスラッシュ コマンド。ダーティ リポジトリのコミット、ドキュメントの更新、永続メモリ システムの更新によってすべてのセッションを終了します。これにより、次のセッションが完全なコンテキストで開始されます。
/dmg = 記録、思い出、それ。
クロード コードを使用してビルドすると、セッションごとにコンテキストがリセットされます。プロジェクトのレイアウトを再度説明します。あなたはすでに下した決定を再度説明します。クロードは昨日何が変わったのか知りません。
もう 1 つの問題は、生産的なセッションの後、コミットされていない作業、古いドキュメント、そして学んだことの記録がないことです。
/dmg は両方を解決します。 1 つのコマンドがセッションをクリーンアップし、次のコマンドを準備します。
最も影響力が大きいのは、クロードが間違いを犯したとき、またはあなたがアプローチを修正したとき、/dmg がそれをフィードバック メモリとして保存することです。クロードが同じ間違いを繰り返す可能性ははるかに低いです。セッション間のコンテキスト損失のほとんどはサイレントで発生します。クロードが不正な呼び出しを繰り返すまで気付かれません。フィードバック記憶はそのサイクルを断ち切るものです。
1. Git — セッション中にタッチされたすべてのリポジトリ (ネストされたリポジトリを含む) を検索し、変更が行われた理由を説明するメッセージとともにダーティなものをコミットし、Claude Code の共同作成者の行を追加します。
2. ドキュメント — README ファイル、サービス テーブル、アーキテクチャ ドキュメントなど、古くなった運用ドキュメントを更新します。新しいプロジェクトとサービスは、適切なテーブルに自動的に追加されます。
3. メモリ — セッション間で保持される一連の構造化マークダウン ファイルを更新します。次は

Claude Code を開くと、インデックスが読み取られ、何も再説明しなくても、プロジェクト、学んだ教訓、好み、インフラストラクチャが認識されます。
最後に、証拠としてすべてのリポジトリにわたるクリーンな git ステータスが表示されます。
npm install -g @anthropic-ai/claude-code
2.スキルをインストールする
mkdir -p ~ /.claude/skills/dmg
cpスキル/dmg/SKILL.md ~ /.claude/skills/dmg/SKILL.md
Claude Code は、~/.claude/skills/ にあるスキル ファイルをロードし、スラッシュ コマンドとして使用できるようにします。ファイルをコピーした後、クロード コード セッションで /dmg と入力します。
# プロジェクト名を選択します - 通常は作業ディレクトリと一致します
mkdir -p ~ /.claude/projects/my-project/memory
cp メモリ/MEMORY.md ~ /.claude/projects/my-project/memory/MEMORY.md
# git リポジトリとして初期化し、自動コミット フックにコミットする場所を確保する
cd ~ /.claude/projects/my-project/memory
gitの初期化
git add MEMORY.md
git commit -m " メモリの初期化 "
作業を進めながら、メモリ ファイルの書き込みを開始します。開始点として、memory/examples/ 内の例を使用してください。
CLAUDE.md をプロジェクト ルートにコピーします (または、その内容を既存のルートに追加します)。これにより、クロードはセッション開始時にメモリ インデックスを読み取るように指示されます。
cp CLAUDE.md ~ /my-project/CLAUDE.md
5. セットアップに合わせてスキルを構成します
スキルはリポジトリのレイアウトとドキュメントの規則を知っている必要があります。最も簡単な方法は、クロードに発見してもらうことです。
/dmg --init
Claude は検出スクリプトを実行し、見つかったもの (リポジトリ、ドキュメント ファイル、メモリ フォルダー) を表示し、4 つの簡単な質問をして、セットアップに合わせて調整された SKILL.md を作成します。 1分ほどかかります。
または、手動で構成します。~/.claude/skills/dmg/SKILL.md を開いて、実際のパスと doc ファイル名に一致するように指示を編集します。少なくとも次のように伝えてください。
どのディレクトリが独自の Git リポジトリであるか、親によって追跡されているか
どのドキュメント ファイルを最新の状態に保つか
あなたの

メモリフォルダーは生きています
/dmg は手動トリガーです。これを呼び出すとコミットされます。 /dmg を実行する前にセッションがクラッシュすると、自動保存のないワークフローと同様に、コミットされていないメモリ更新が失われます。
呼び出し間で継続的な保護が必要な場合は、メモリ フォルダーを自動的にコミットするフックを追加します。これはセーフティ ネット層です。 /dmg はクリーンアップ パスです。衝突安全性を求めるなら両方が必要です。
フックはツール呼び出しのたびに (タイマーではなく) 起動されるため、スクリプト自体がデバウンスされます。最後のコミットが 5 分以内である場合はスキップされます。
クロード コードの設定 ( ~/.claude/settings.json ):
{
"フック" : {
"PostToolUse" : [
{
"マッチャー" : " " ,
「フック」: [
{
"タイプ" : "コマンド" ,
"コマンド" : " bash ~/.claude/skills/dmg/autocommit.sh "
}
】
}
】
}
}
autocommit.sh ( ~/.claude/skills/dmg/autocommit.sh で作成):
#! /bin/bash
MEMORY_DIR= " $HOME /.claude/projects/my-project/memory "
cd " $MEMORY_DIR " || 0番出口
[[ -z $( git status --porcelain ) ]] && 0 を終了します
last= $( git log -1 --format=%ct 2> /dev/null || echo 0 )
[[ $(( $(date +% s) - last )) -lt 300 ]] && exit 0
git add -A && git commit -m " auto: $( date ' +%Y-%m-%d %H:%M:%S ' ) " --quiet
メモリファイル形式
各メモリ ファイルは、YAML フロントマターを含むマークダウン ファイルです。
---
名前 : ショートケバブケースナメクジ
説明 : " 1 行 — インデックスをスキャンするときに関連性を判断するために使用されます"
メタデータ:
タイプ: プロジェクト # プロジェクト |フィードバック |ユーザー |参照
---
内容はこちら。フィードバック/レッスンについては、次を使用します。
ルールとか行動とか。
** 理由: ** 理由 — 多くの場合、過去の出来事や強い好み。
** 適用方法: ** いつ、どこで適用されるか。
MEMORY.md インデックスには、ファイルごとに 1 行が含まれます。
- [ Title ] ( filename.md ) — 説明フィールドに一致する 1 行のフック
メモリの種類
種類
ここに何が入るのか
プロジェクト
現在の状態、ファイルパス、

オープンスレッド、アーキテクチャの決定
フィードバック
学んだ教訓 — 何をすべきか、何を避けるべきか、そしてその理由
ユーザー
あなたの役割、専門知識、好み、働き方
参照
外部リソースへのポインタ: ダッシュボード、ドキュメント、チケット
使用法
/dmg
クロードはセッション中にタッチされたすべてのリポジトリをスキャンし、汚い作業をコミットし、ドキュメントを更新し、メモリ ファイルを同期します。変更の量に応じて 1 ～ 3 分かかります。
タスクを切り替える前にチェックポイントを作成したい場合は、セッション中にトリガーすることもできます。
SKILL.md には具体的に記述してください。一般的な命令は一般的な結果を生成します。実際のリポジトリのパスとドキュメント ファイル名をリストします。
記憶は時間の経過とともに増大します。最初のセッションは最小限です。数週間後にはスタック全体が認識されます。
フィードバック メモリは最も価値の高いタイプです。クロードがミスをしたり、アプローチを修正したりした場合は、すぐにそれをフィードバックメモリとして保存します。同じ間違いを再び犯す可能性ははるかに低くなります。
インデックスは短くしてください。 MEMORY.md はセッションごとにロードされます。長くなると、クロードはスキャンにトークンを消費します。ファイルごとに 1 つのタイトな行。
Claude Code /dmg スキル — 永続メモリ + セッション同期 (git、docs、1 つのコマンドでメモリ)
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

Claude Code /dmg skill — persistent memory + session sync (git, docs, memory in one command) - responsiblparty/claude-dmg-skill

GitHub - responsiblparty/claude-dmg-skill: Claude Code /dmg skill — persistent memory + session sync (git, docs, memory in one command) · GitHub
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
responsiblparty
/
claude-dmg-skill
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
responsiblparty/claude-dmg-skill
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits memory memory skills/ dmg skills/ dmg CLAUDE.md CLAUDE.md README.md README.md View all files Repository files navigation
/dmg — Claude Code skill for persistent memory + auto-sync
A Claude Code slash command that ends every session by committing dirty repos, refreshing documentation, and updating a persistent memory system — so the next session starts with full context.
/dmg = d ocumentation, m emories, g it.
When you build things with Claude Code, context resets every session. You re-explain your project layout. You re-describe decisions you already made. Claude doesn't know what changed yesterday.
The other problem: after a productive session you have uncommitted work, stale docs, and no record of what you learned.
/dmg solves both. One command cleans up the session and primes the next one.
The highest-leverage thing it does: when Claude makes a mistake or you correct its approach, /dmg saves that as a feedback memory . Claude is far less likely to repeat the same mistake. Most context loss between sessions is silent — you don't notice until Claude repeats a bad call. Feedback memories are what break that cycle.
1. Git — finds every repo touched during the session (including nested repos), commits anything dirty with a message explaining why the change was made, and appends the Claude Code co-author line.
2. Docs — updates whatever operational documentation went stale: README files, service tables, architecture docs. New projects and services get added to the right tables automatically.
3. Memory — updates a set of structured markdown files that persist across sessions. The next time you open Claude Code, it reads the index and knows your projects, your lessons learned, your preferences, and your infrastructure — without you re-explaining anything.
At the end it shows a clean git status across all repos as proof.
npm install -g @anthropic-ai/claude-code
2. Install the skill
mkdir -p ~ /.claude/skills/dmg
cp skills/dmg/SKILL.md ~ /.claude/skills/dmg/SKILL.md
Claude Code loads any skill files found in ~/.claude/skills/ and makes them available as slash commands. After copying the file, type /dmg in any Claude Code session.
# Pick a project name — usually matches your working directory
mkdir -p ~ /.claude/projects/my-project/memory
cp memory/MEMORY.md ~ /.claude/projects/my-project/memory/MEMORY.md
# Initialize as a git repo so the autocommit hook has somewhere to commit
cd ~ /.claude/projects/my-project/memory
git init
git add MEMORY.md
git commit -m " init memory "
Start writing memory files as you go. Use the examples in memory/examples/ as a starting point.
Copy CLAUDE.md to your project root (or add its contents to an existing one). This tells Claude to read the memory index at session start.
cp CLAUDE.md ~ /my-project/CLAUDE.md
5. Configure the skill for your setup
The skill needs to know your repo layout and doc conventions. The easiest way is to let Claude discover them:
/dmg --init
Claude will run a discovery script, show you what it found (repos, doc files, memory folders), ask four quick questions, and write a tailored SKILL.md for your setup. Takes about a minute.
Or configure manually: open ~/.claude/skills/dmg/SKILL.md and edit the instructions to match your actual paths and doc file names. At minimum, tell it:
Which directories are their own git repos vs. tracked by a parent
Which documentation files to keep current
Where your memory folder lives
/dmg is a manual trigger — you invoke it and it commits. A session crash before you run /dmg means uncommitted memory updates are lost, the same as any workflow without autosave.
If you want continuous protection between invocations, add a hook that commits the memory folder automatically. This is the safety net layer; /dmg is the cleanup pass. You need both if you want crash safety.
The hook fires after every tool call (not on a timer), so the script debounces itself — it skips if the last commit was less than 5 minutes ago.
In your Claude Code settings ( ~/.claude/settings.json ):
{
"hooks" : {
"PostToolUse" : [
{
"matcher" : " " ,
"hooks" : [
{
"type" : " command " ,
"command" : " bash ~/.claude/skills/dmg/autocommit.sh "
}
]
}
]
}
}
autocommit.sh (create at ~/.claude/skills/dmg/autocommit.sh ):
#! /bin/bash
MEMORY_DIR= " $HOME /.claude/projects/my-project/memory "
cd " $MEMORY_DIR " || exit 0
[[ -z $( git status --porcelain ) ]] && exit 0
last= $( git log -1 --format=%ct 2> /dev/null || echo 0 )
[[ $(( $(date +% s) - last )) -lt 300 ]] && exit 0
git add -A && git commit -m " auto: $( date ' +%Y-%m-%d %H:%M:%S ' ) " --quiet
Memory file format
Each memory file is a markdown file with YAML frontmatter:
---
name : short-kebab-case-slug
description : " one line — used to decide relevance when scanning the index "
metadata :
type : project # project | feedback | user | reference
---
Content here. For feedback/lessons, use:
The rule or behavior.
** Why: ** The reason — often a past incident or strong preference.
** How to apply: ** When and where this kicks in.
The MEMORY.md index has one line per file:
- [ Title ] ( filename.md ) — one-line hook matching the description field
Memory types
Type
What goes here
project
Current state, file paths, open threads, architecture decisions
feedback
Lessons learned — what to do/avoid, and why
user
Your role, expertise, preferences, how you like to work
reference
Pointers to external resources: dashboards, docs, tickets
Usage
/dmg
Claude will scan every repo touched during the session, commit dirty work, update your docs, and sync the memory files. Takes 1–3 minutes depending on how much changed.
You can also trigger it mid-session if you want to checkpoint before switching tasks.
Be specific in SKILL.md. Generic instructions produce generic results. List your actual repo paths and doc file names.
Memory compounds over time. The first session it's minimal. After a few weeks it knows your entire stack.
Feedback memories are the highest-value type. When Claude makes a mistake or you correct its approach, save that as a feedback memory immediately. It is far less likely to make the same mistake again.
Keep the index short. MEMORY.md is loaded every session — if it gets long, Claude spends tokens scanning it. One tight line per file.
Claude Code /dmg skill — persistent memory + session sync (git, docs, memory in one command)
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
