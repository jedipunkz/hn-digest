---
source: "https://github.com/royashbrook/hush"
hn_url: "https://news.ycombinator.com/item?id=48691038"
title: "Hush, let an AI agent use your secrets without ever seeing them"
article_title: "GitHub - royashbrook/hush: a secret store for AI agents with one rule: the agent never sees the plaintext. get a secret once into the OS keychain, then inject it into commands forever. no get, cross-platform, MIT. · GitHub"
author: "royashbrook"
captured_at: "2026-06-26T20:33:15Z"
capture_tool: "hn-digest"
hn_id: 48691038
score: 3
comments: 1
posted_at: "2026-06-26T19:39:57Z"
tags:
  - hacker-news
  - translated
---

# Hush, let an AI agent use your secrets without ever seeing them

- HN: [48691038](https://news.ycombinator.com/item?id=48691038)
- Source: [github.com](https://github.com/royashbrook/hush)
- Score: 3
- Comments: 1
- Posted: 2026-06-26T19:39:57Z

## Translation

タイトル: 黙って、AI エージェントにあなたの秘密を誰にも見られずに使用させてください
記事のタイトル: GitHub - royashbrook/hush: AI エージェント用の秘密ストア。ルールは 1 つあります。「エージェントは平文を決して見ない」というものです。シークレットを OS キーチェーンに一度取得すると、それをコマンドに永久に挿入します。取得できません、クロスプラットフォーム、MIT。 · GitHub
説明: AI エージェント用の秘密ストア。ルールは 1 つあります。「エージェントは平文を決して見ない」というものです。シークレットを OS キーチェーンに一度取得すると、それをコマンドに永久に挿入します。取得できません、クロスプラットフォーム、MIT。 - ロヤシュブルック/ハッシュ

記事本文:
GitHub - royashbrook/hush: AI エージェント用の秘密ストア。ルールは 1 つあります。エージェントは平文を決して参照しません。シークレットを OS キーチェーンに一度取得すると、それをコマンドに永久に挿入します。取得できません、クロスプラットフォーム、MIT。 · GitHub
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
ロイヤルシュブルック
/
夫

h
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
23 コミット 23 コミット .github/ workflows .github/ workflows アセット アセット ヘルパー ヘルパー テスト テスト win win .gitignore .gitignore AGENTS.md AGENTS.md ライセンス ライセンス README.md README.md SKILL.md SKILL.md hush hush package.json package.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
AI エージェント用の秘密ストア、get のないもの。
シークレットを OS キーチェーンに一度取得すると、それをコマンドに永久に挿入します。その値が stdout、トランスクリプト、またはクラウドに到達することはありません。
AI エージェント用の秘密ストア。エージェントには平文が決して表示されないという 1 つの厳しいルールがあります。
値は stdout に到達しないため、ツールの結果、トランスクリプト、またはクラウドに到達することはありません。あ
秘密は店舗から消費者に直接届けられるだけです。意図的に取得することはできません、
プレーンゲッターがリークです。
ポイント: あなた (エージェント) は、すでに認証された CLI を持つユーザーとして実行しているため、すでに
サーバー側のシークレットを設定するか、サービスを呼び出すと、値が
チャット。 hush はそれを OS キーチェーンに一度取得すると、それらのコマンドに永久に挿入します。いいえ
チャットペースト、一時ファイル、.env を誤ってコミットする必要はありません。
このリポジトリの URL を渡された場合は、自分でインストールし、プレイブックに従ってください。静寂を得る
スクリプトをボックスに貼り付けて SKILL.md を読み取ります。これがオンボーディング全体です (ストア / をカバーします)。
mint / inject と、作業しているリポジトリで hash を採用する方法)。スクリプトは自己完結型です
bash ファイル;エージェントがツールをロードする場所、または PATH 上にそれを配置します。
git clone https://github.com/royashbrook/hush
chmod +x ハッシュ/ハッシュ
# 次に、hush/hush を PATH に置くか、ag に移動します

ent のスキル/ツールのディレクトリ
または、クローンを作成せずに npm から直接実行します: npx @royashbrook/hush <cmd> (mac/linux; Windows では
クローン、オーケストレーターは bash です)。
クロード コードの利便性 (スキルとして自動ロード): 代わりにスキル ディレクトリに直接クローンを作成します。
セッション git clone https://github.com/royashbrook/hush ~/.claude/skills/hush 、または 1 つのプロジェクト
./.claude/スキル/ハッシュ 。その他のエージェント: AGENTS.md を参照してください。
それだけです、人間はあなたをここに指差し、「沈黙スキルを使用してください」と言い、あなたはここからそれを使います。の
人間だけができることの 1 つは、自分が保持している秘密を貼り付けることです。それをヒットしたら、hush set <name> を実行します。
そして、隠しダイアログが画面に表示されます (値は表示されません)。
上記と同じクローン、または単一のハッシュ スクリプトを取得して PATH に置きます。
バックエンド: macOS キーチェーン、Linux libsecret ( Secret-tool )、および Windows DPAPI (PowerShell 経由)
他のプラットフォームでは、OS シークレット ストアと同じ契約を維持します ( SKILL.md を参照)。
CI でクロス OS テスト済み (手動 workflow_dispatch 、3 つのランナーすべて)。
hush set my-vendor-token # 保持している値を貼り付けます (非表示のプロンプト)
printf ' %s ' " $TOK " | hush set my-vendor-token # ...またはパイプで (スクリプト/CI)、argv をオフのままにしてください
hush mint app-operator-key # ランダムなキーを生成して保存します
hush run TOKEN=my-vendor-token -- some-cmd # コマンドに挿入しますが、出力されません
ハッシュ リスト # 名前のみ、値は使用しない
命名: デフォルトのハッシュ名前空間を維持し、プロジェクトごとに名前にプレフィックスを付けます (blame-cf-token 、
lifescored-gemini-key ) なので、hush を 1 回キーチェーン検索するとすべてが見つかります。 HUSH_NS は、
プロジェクトごとではなく、完全に別個のストアです。既存の名前を修正する必要がありますか? <古い> <新しい> 名前を変更してください
値を内部的に移動します (再質問されたり、出力されたりすることはありません)。完全なドキュメント + ポータブル契約:
SKILL.md 。
シェルアクセス権を持つエージェントは、このストアの読み取りと書き込みを行うことができます。

したがって、これは敵対的なプロセスに対するロックではありません。
トランスクリプトから平文を排除し、「一度保存すればどこにでも注入」できる構造です。
簡単な道。また、耐久性は、それが搭載されているマシン (ローカル キーチェーン) と同じ程度です。マシンに戻ります。
本物のシークレットマネージャーにアップするか、同期してください。hush をシークレットの唯一のコピーにしないでください。
再生できません。 MITライセンス取得済み。
AI エージェント用の秘密ストア。ルールは 1 つあります。エージェントは平文を決して見ないというものです。シークレットを OS キーチェーンに一度取得すると、それをコマンドに永久に挿入します。取得できません、クロスプラットフォーム、MIT。
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

a secret store for AI agents with one rule: the agent never sees the plaintext. get a secret once into the OS keychain, then inject it into commands forever. no get, cross-platform, MIT. - royashbrook/hush

GitHub - royashbrook/hush: a secret store for AI agents with one rule: the agent never sees the plaintext. get a secret once into the OS keychain, then inject it into commands forever. no get, cross-platform, MIT. · GitHub
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
royashbrook
/
hush
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
23 Commits 23 Commits .github/ workflows .github/ workflows assets assets helpers helpers test test win win .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md SKILL.md SKILL.md hush hush package.json package.json View all files Repository files navigation
a secret store for AI agents , the one with no get .
get a secret once into the OS keychain, then inject it into commands forever , the value never hits stdout, the transcript, or the cloud.
A secret store for AI agents, with one hard rule: the agent never sees the plaintext.
Values never reach stdout, so they never land in the tool result, the transcript, or the cloud. A
secret only ever moves from the store straight into the consumer. There's no get on purpose, a
plain getter is the leak.
The point: you (an agent) are running as the user with their CLIs already authed, so you can already
set a server-side secret or call a service , you just can't see the value without it landing in the
chat. hush gets it once into the OS keychain, then you inject it into those commands forever. No
chat paste, no temp files, no .env to commit by accident.
If you were handed this repo's url, install yourself, then follow the playbook. Get the hush
script onto the box and read SKILL.md , that's the whole onboarding (it covers store /
mint / inject and how to adopt hush in the repo you're working in). The script is one self-contained
bash file; put it wherever your agent loads tools or just on your PATH :
git clone https://github.com/royashbrook/hush
chmod +x hush/hush
# then put hush/hush on your PATH, or move it into your agent's skills/tools dir
or run it straight from npm without cloning: npx @royashbrook/hush <cmd> (mac/linux; on windows use the
clone, the orchestrator is bash).
Claude Code convenience (auto-loads as a skill): clone straight into the skills dir instead , whole
session git clone https://github.com/royashbrook/hush ~/.claude/skills/hush , or one project
./.claude/skills/hush . Other agents: see AGENTS.md .
That's it , the human points you here and says "use the hush skill," and you take it from here. The
one thing only a human can do is paste a secret they hold: when you hit that, run hush set <name>
and a hidden dialog pops on their screen (you never see the value).
Same clone as above, or just grab the single hush script and put it on your PATH.
Backends: macOS Keychain, Linux libsecret ( secret-tool ), and Windows DPAPI (via PowerShell) are
built in. On other platforms keep the same contract with your OS secret store (see SKILL.md ).
Cross-OS tested in CI (manual workflow_dispatch , all three runners).
hush set my-vendor-token # paste a value you hold (hidden prompt)
printf ' %s ' " $TOK " | hush set my-vendor-token # ...or pipe it in (scripts/CI), still off argv
hush mint app-operator-key # generate + store a random one
hush run TOKEN=my-vendor-token -- some-cmd # inject into a command, never printed
hush list # names only, never values
Naming: keep the default hush namespace and prefix names by project ( blame-cf-token ,
lifescored-gemini-key ) so one keychain search for hush finds everything. HUSH_NS is only for a
genuinely separate store, not per-project. Need to fix an existing name? hush rename <old> <new>
moves the value internally (never re-asked, never printed). Full docs + the portable contract:
SKILL.md .
An agent with shell access can read+write this store, so it's not a lock against a hostile process.
It's structure that keeps plaintext out of the transcript and makes "store once, inject everywhere"
the easy path. It's also only as durable as the machine it's on (a local keychain) , back the machine
up, or sync onward into a real secret manager, and don't make hush the only copy of a secret you
can't regenerate. MIT licensed.
a secret store for AI agents with one rule: the agent never sees the plaintext. get a secret once into the OS keychain, then inject it into commands forever. no get, cross-platform, MIT.
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
