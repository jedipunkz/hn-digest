---
source: "https://github.com/vandith1/agent-guard"
hn_url: "https://news.ycombinator.com/item?id=49319603"
title: "Show HN: A pre-execution guard that stops AI agents running destructive commands"
article_title: "GitHub - vandith1/agent-guard: A pre-execution guard that stops your coding agent running destructive commands. One shell file, no dependencies. · GitHub"
author: "andevandith"
captured_at: "2026-08-16T13:24:00Z"
capture_tool: "hn-digest"
hn_id: 49319603
score: 3
comments: 1
posted_at: "2026-08-16T12:53:51Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A pre-execution guard that stops AI agents running destructive commands

- HN: [49319603](https://news.ycombinator.com/item?id=49319603)
- Source: [github.com](https://github.com/vandith1/agent-guard)
- Score: 3
- Comments: 1
- Posted: 2026-08-16T12:53:51Z

## Translation

タイトル: Show HN: AI エージェントによる破壊的なコマンドの実行を停止する実行前ガード
記事のタイトル: GitHub - vandith1/agent-guard: コーディング エージェントによる破壊的なコマンドの実行を停止する実行前ガード。シェル ファイルは 1 つで、依存関係はありません。 · GitHub
説明: コーディング エージェントによる破壊的なコマンドの実行を停止する実行前ガード。シェル ファイルは 1 つで、依存関係はありません。 - vandith1/エージェントガード

記事本文:
GitHub - vandith1/agent-guard: コーディング エージェントによる破壊的なコマンドの実行を停止する実行前ガード。シェル ファイルは 1 つで、依存関係はありません。 · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ヴァンディス1
/
エージェントガード
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット docs docs ライセンス ライセンス README.md README.md Guard-command.sh Guard-command.sh install.sh install.sh test-guard.sh test-guard.sh 表示

ll ファイル リポジトリ ファイルのナビゲーション
コーディング エージェントは最終的に git replace --hard 、force-push to main 、または
テーブルをドロップします。不注意だからではなく、自信があるから、そしてループ
3歩前に漂流してしまった。
ブラウザで試してください → — 任意のコマンドを貼り付けて、ブロックされるかどうか、またどのルールによってブロックされるかを確認します。どこにも何も送信されません。ルールはローカルで実行されます。
これは、元に戻すことができないコマンドを停止する、約 60 行のシェル フックです。走ります
エージェントが提案するすべてのシェル コマンドの前に、2 を終了して拒否し、
標準エラー出力の理由 — どのエージェント CLI がループにフィードバックされるため、エージェントが停止するか
再試行すると、代わりにコマンドが表示されます。
ガードによってブロックされました [すべてのローカル変更を破棄します (git restart --hard)]
コマンド: git restart --hard HEAD~3
これは人間にゲートがかかっています。やり直したり、言い直したりしないでください。
正確なコマンドを表示し、人間に実行させます。
インストール
bash install.sh # PreToolUse フックを .claude/settings.json にマージします
bash test-guard.sh # 6 個のアサーション、6 個が合格 / 0 個が失敗する必要があります
Claude Code の PreToolUse フック、および
ツール呼び出しの前にシェルをフックし、その終了コードを尊重します。 bash が必要です。
grep -E 、および python3 。
git replace --hard · rm -rf · git clean -f/-x · git checkout 。 ·
git filter-branch · main/master/production に強制プッシュ · DROP TABLE ·
TRUNCATE · FLUSHALL · aws delete-* / terminate-* · kubectl delete ·
terraform destroy · docker … プルーン · chmod 777 · カール … |し・
raw は /dev/sd* に書き込みます。
ルールは、スクリプトの先頭にある 1 つの配列に label@@regex ペアとして存在します。編集
彼ら。テスト ケースを追加する場合は、無害なコマンドや新しいパターンを含めてください。
捕まえてはいけません。そうすることで、過度に貪欲な正規表現を見つけられる前に見つけることができます。
これで何が起こるのか、何が止まらないのか
それは止まります: 自信を持ったエージェントのプロ

取り消しできないコマンドをプレーンテキストで発行します。
私が実際に遭遇した障害モード — ループは 3 ステップ前にドリフトし、次のコマンドは
その日の仕事。
それは止まらない：意図的にそれを回避するものは何でも。 g""リセット --hard 、base64 ペイロード
eval 、シェルスクリプト、エイリアスを介して。マッチャーは、提案されたコマンド文字列に対する正規表現です。
そして正規表現は毎回敵に負けます。
この違いが脅威モデル全体であり、率直に言う価値があります。
インストールしたフックを回避するためにコマンドを難読化する場合、パターン リストは問題ではありません。これ
は、悪意のある瞬間ではなく、善意のモデルが悪い瞬間を経験していることを前提としています。
また、それらはユーザーの外部にあるため、厳密に信頼性の高いものに代わるものでもありません。
マシン: サーバー側のブランチ保護、最小権限の資格情報、およびバックアップ。ここには何もありません
それらのいずれかを置き換えます。それらを使用し、不可逆的な操作にはこれを使用します。
カバー — コミットされていない作業のハード リセット、再帰的削除、テーブルの削除、スタックの破棄。
これは、引数を含むコマンド文字列全体と一致します。つまりコミットメッセージ
ゲートされたフレーズを引用するものはブロックされます。それは意図的なものです。代替案は次のとおりです。
シェルの引用、展開、および eval を解析し、すべてのパーサーは次の新しい方法です。
ガードをすり抜けてコマンドを渡す。フェイルセーフは賢明に勝ります。それを修正する習慣:
-m "…" の代わりに git commit -F .commit-msg を使用し、ヒアドキュメントではありません。
閉じられません。フック ペイロードからコマンドを読み取れない場合は、
ブロック。 「すべてを許可する」ために黙って劣化するガードは、何もしないより悪いです
見守るのをやめてしまうから。
破壊的なコマンドと外部効果のあるコマンドは別の問題です。
層 1 — rm -rf 、reset --hard 、DROP TABLE 。回復不能。決してしない
交渉可能、ロック解除なし、電子なし

お急ぎのあなたにも。このリポジトリは階層 1 です。
階層 2 — git Push 、 npm publish 、デプロイ、移行。実際の効果ですが、
見直し可能で元に戻すことができます。これらは禁止されるべきではなく、ゲートされるべきです。
短い時間であれば人間がロックを解除でき、自動的に再ロックされます。
Tier 2 はデザインが興味深いところです。ロック解除は何かである必要があります。
エージェントはタイムスタンプを更新してウィンドウを延長したり、タッチしたりすることはできません。
使用されたことを隠すために削除します。
フルキットでは、タイムボックスによるロック解除、構成ファイル、git を備えた階層 2 が追加されます。
プリプッシュフック、分離ワークフロー、プロンプトをキャッチするマージゲート
決定論的グレーダーによる回帰と、それぞれの背後にあるインシデントの記録
ルール: https://andevan.gumroad.com/l/agent-control-kit
MITライセンス取得済み。それを取り出し、フォークし、独自のツールに入れて出荷します。
コーディング エージェントによる破壊的なコマンドの実行を停止する実行前ガード。シェル ファイルは 1 つで、依存関係はありません。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A pre-execution guard that stops your coding agent running destructive commands. One shell file, no dependencies. - vandith1/agent-guard

GitHub - vandith1/agent-guard: A pre-execution guard that stops your coding agent running destructive commands. One shell file, no dependencies. · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
vandith1
/
agent-guard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits docs docs LICENSE LICENSE README.md README.md guard-command.sh guard-command.sh install.sh install.sh test-guard.sh test-guard.sh View all files Repository files navigation
Your coding agent will eventually run git reset --hard , force-push to main, or
drop a table. Not because it's careless — because it's confident, and the loop
drifted three steps ago.
Try it in your browser → — paste any command, see whether it would be blocked and by which rule. Nothing is sent anywhere; the rules run locally.
This is a ~60-line shell hook that stops the commands you can't undo. It runs
before every shell command the agent proposes, exits 2 to deny, and writes the
reason to stderr — which agent CLIs feed back into the loop, so the agent stops
retrying and surfaces the command to you instead.
BLOCKED by guard [discard all local changes (git reset --hard)]
command: git reset --hard HEAD~3
This one is gated to the human. Do not retry, do not reword.
Surface the exact command and let the human run it.
Install
bash install.sh # merges a PreToolUse hook into .claude/settings.json
bash test-guard.sh # 6 assertions, should be 6 passed / 0 failed
Works with Claude Code's PreToolUse hook, and with any agent CLI that runs a
shell hook before a tool call and respects its exit code. Requires bash ,
grep -E , and python3 .
git reset --hard · rm -rf · git clean -f/-x · git checkout . ·
git filter-branch · force-push to main/master/production · DROP TABLE ·
TRUNCATE · FLUSHALL · aws delete-* / terminate-* · kubectl delete ·
terraform destroy · docker … prune · chmod 777 · curl … | sh ·
raw writes to /dev/sd* .
Rules live in one array at the top of the script as label@@regex pairs. Edit
them. Add a test case when you do — including a benign command the new pattern
must not catch. That's how you find an over-greedy regex before it finds you.
What this does and does not stop
It stops: a confident agent proposing an irreversible command in plain text, which is the
failure mode I actually hit — the loop drifted three steps ago and the next command discards a
day's work.
It does not stop: anything deliberately evading it. g""it reset --hard , a base64 payload
through eval , a shell script, an alias. The matcher is a regex over the proposed command string,
and a regex loses to an adversary every time.
That distinction is the whole threat model, and it is worth being blunt about: if your agent is
obfuscating commands to get around a hook you installed, a pattern list is not your problem. This
assumes a well-intentioned model having a bad moment, not a hostile one.
It also does not replace the things that are strictly more reliable because they sit outside your
machine: server-side branch protection, least-privilege credentials, and backups. Nothing here is a
substitute for any of those. Use them, and use this for the irreversible operations they do not
cover — a hard reset on uncommitted work, a recursive delete, a dropped table, a destroyed stack.
It matches the whole command string, arguments included. So a commit message
that quotes a gated phrase gets blocked. That's deliberate: the alternative is
parsing shell quoting, expansion and eval , and every parser is a new way to
slip a command past the guard. Fail-safe beats clever. The habit that fixes it:
git commit -F .commit-msg instead of -m "…" , never a heredoc.
It fails closed. If it can't read the command out of the hook payload, it
blocks. A guard that silently degrades to "allow everything" is worse than no
guard, because you'd stop watching.
Destructive commands and external-effect commands are different problems.
Tier 1 — rm -rf , reset --hard , DROP TABLE . Unrecoverable. Never
negotiable, no unlock, not even for you-in-a-hurry. This repo is tier 1.
Tier 2 — git push , npm publish , deploys, migrations. Real effects, but
reviewable and reversible. These shouldn't be banned — they should be gated,
unlockable by a human for a short window that auto-relocks.
Tier 2 is where the design gets interesting: the unlock has to be something the
agent can't touch, including refreshing its timestamp to extend the window or
deleting it to hide that it was used.
The full kit adds tier 2 with the time-boxed unlock, a config file, a git
pre-push hook, an isolation workflow, a merge gate that catches prompt
regressions with a deterministic grader, and the incident write-ups behind each
rule: https://andevan.gumroad.com/l/agent-control-kit
MIT licensed. Take it, fork it, ship it in your own tooling.
A pre-execution guard that stops your coding agent running destructive commands. One shell file, no dependencies.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
