---
source: "https://github.com/budhasantosh010/ai-coding-project-templates"
hn_url: "https://news.ycombinator.com/item?id=48660491"
title: "Hw to make sure AI Agents neer forget the Work?"
article_title: "GitHub - budhasantosh010/ai-coding-project-templates: Stop Claude Code and Codex forgetting your project mid-build. Drop-in memory + governance templates that write your decisions to disk and inject them back every session, even after a context compaction. · GitHub"
author: "realsanb"
captured_at: "2026-06-24T14:56:35Z"
capture_tool: "hn-digest"
hn_id: 48660491
score: 2
comments: 0
posted_at: "2026-06-24T14:26:59Z"
tags:
  - hacker-news
  - translated
---

# Hw to make sure AI Agents neer forget the Work?

- HN: [48660491](https://news.ycombinator.com/item?id=48660491)
- Source: [github.com](https://github.com/budhasantosh010/ai-coding-project-templates)
- Score: 2
- Comments: 0
- Posted: 2026-06-24T14:26:59Z

## Translation

タイトル: AI エージェントが仕事を忘れないようにする方法は?
記事のタイトル: GitHub - budhasantosh010/ai-coding-project-templates: Claude Code と Codex がビルド中にプロジェクトを忘れるのを阻止します。ドロップイン メモリ + ガバナンス テンプレート。コンテキスト圧縮後も含め、決定をディスクに書き込み、セッションごとにそれを挿入します。 · GitHub
説明: Claude Code と Codex がビルド中にプロジェクトを忘れるのを防ぎます。ドロップイン メモリ + ガバナンス テンプレート。コンテキスト圧縮後も含め、決定をディスクに書き込み、セッションごとにそれを挿入します。 - budhasantosh010/ai-coding-project-templates

記事本文:
GitHub - budhasantosh010/ai-coding-project-templates: Claude Code と Codex がビルド中にプロジェクトを忘れるのを防ぎます。ドロップイン メモリ + ガバナンス テンプレート。コンテキスト圧縮後も含め、決定をディスクに書き込み、セッションごとにそれを挿入します。 · GitHub
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
ディ

警告を無視する
{{ メッセージ }}
ブダサントシュ010
/
AIコーディングプロジェクトテンプレート
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
budhasantosh010/ai-coding-project-templates
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
11 コミット 11 コミット claude-project-template claude-project-template codex-project-template codex-project-template .gitignore .gitignore LICENSE LICENSE README.md README.md START_HERE.md START_HERE.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
コーディングエージェントは忘れてしまいます。ビルドの途中でコンテキストが圧縮され、突然
1時間前に決めたことを再質問したり、初日に決めたルールを黙って無視したり。これら
2 つのドロップイン テンプレート (1 つはクロード コード用、もう 1 つはコーデックス用) です。これらは、
プロジェクトのメモリをファイルに保存し、自動的にエージェントにフィードバックします。
ai-coding-プロジェクト-テンプレート/
§─ claude-project-template/ クロード コード (ルールブック = CLAUDE.md、.claude/settings.json 経由のフック)
━─ codex-project-template/ OpenAI Codex 用 (ルールブック = AGENTS.md、.codex/hooks.json によるフック)
どちらのフォルダーも同じシステムです。唯一の違いは、ルールブックのファイル名とそれぞれのルールブックのファイル名です。
エージェントがフックを登録します。
リコールは望まれることではなく、どのように強制されるのか
意味から物事を調べる（思い出す）
デシジョン ツリー (それを参照し、そこにロールバックします)
オプションのコンパニオン: グラフ化
Graphicify はこれらのテンプレートから自動起動しますか?
一度インストールすればプロジェクトごとにビルド可能
おそらく有料の手順は必要ありません
長時間実行されるエージェント セッションで問題が発生する 4 つの問題と、それに対するテンプレートの対応
それぞれ:
なし あり
────── ────
セッション間で忘れてしまいます。すべてのメッセージ、決定、変更は DOCS/ に反映されます。
プロジェクトについて常に再説明します。だから新鮮なセッションr

e は完全なコンテキストを読み込みます。
「完了」は「コードが存在する」ことを意味し、「完了」はテストの合格または実際の実行 (E0 ～ E5) を意味します。
たとえ走らなかったとしても。インポートするコードだけではありません。
同じエラーが永遠に再試行されます。同じ失敗で 3 回ストライク → 停止して診断します。
サイレントフォールバック。理性を失った。フォールバックを明記する必要があります。決断も失敗も、
戻りを停止するテストで記録されます。
エージェントを選択してください
あなたが使用する
このフォルダを開く
ルールブック
クロード コード (CLI / VS コード / デスクトップ)
クロードプロジェクトテンプレート/
クロード.md
OpenAI Codex (CLI / VS Code / デスクトップ)
コーデックスプロジェクトテンプレート/
エージェント.md
インストール (プロジェクトごと)
テンプレートを次の場所にコピーするときに、これをプロジェクトごとに 1 回実行します。
フォルダーの内容をプロジェクト ルートにコピーします。
<PROJECT_NAME> / <PROJECT_ROOT> / <OWNER> / <DATE> プレースホルダーを入力します。
エージェントでプロジェクトを開き、そのフックを信頼します。
両方のフック/verify_*.ps1 チェックを実行します。これらはすべてパスするはずです。
DOCS/STARTUP_MESSAGE.md の first-session ブロックを最初のチャットに貼り付けます。 (
プロンプトは START_HERE.md にも収集されます。)
リコールは望まれることではなく、どのように強制されるのか
コンテキストをファイルに保存するのは簡単です。難しいのは、エージェントに実際に見てもらうことです
会話が進んだ後はそれで終わりです。 「不明な場合は DECISIONS.md を確認してください」と指示するのはサインです
壁の上では、壁を通り過ぎてしまう可能性があります。
したがって、テンプレートはそれに依存しません。 3 つのフックがファイルを読み取り、関連するビットをプッシュします
忘れがちな瞬間にエージェントのビューに戻ります。
フックに何を戻すか
──── ──── ───────
セッションの開始または圧縮 inject_context CURRENT_STATE + DEC/REQ/FAIL リスト
あなたはメッセージ inject_on_prompt を送信し、アクティブなルール + 「トランスクリプトを読んでください」
編集の直前に、編集時にアクティブな DEC/REQ ルールを inject_decions_preedit します。
T

ポイント: 「npm ではなく pnpm を使用します」は、モデルがそれを記憶しているかどうかに応じて停止します。ルールはオンです
セッション開始時、すべてのメッセージ時、およびエージェントがインストールを書き込む直前の画面
コマンド。フックはスキップできないため、情報がそこにあることが保証されます。そして彼ら全員が
フェイルセーフ — フックでエラーが発生した場合、何も出力されず、セッションがブロックされることはありません。
意味から物事を調べる（思い出す）
ルールを再挿入すると、最近の内容が処理されます。しかし、「私たちが何を決めたのか」についてはどうでしょうか。
3か月前の価格設定ですか？」 — 長いトランスクリプトに埋もれており、別の言葉で表現されている可能性があります。
これがリコールフック (各メッセージの remember.ps1 ) の機能であり、ほぼコストがかかるように構築されています。
何もありません:
1. メッセージは振り返りますか? (「覚えてます…」、「以前」、「あのバグ」、「それ/これ」)
いいえ→何もしません。トークンは0です。 (ほとんどのメッセージ)
はい→続行します。
2. あいまいな言葉を解決する:「見つけてください」→最も最近言われたこと。
3. 安い順に検索します。
決定事項 + トランスクリプトよりも第 1 層のキーワード (無料、インスタント)
意味による層 2 のセマンティック — 層 1 が弱い場合のみ (ローカル モデル; 「ログイン」で「認証」が見つかる)
4. 確認: ヒットがファイル名である場合は、現在のファイルを grep → CONFIRMED または STALE と表示します。
5. 小さな引用ポインター (約 40 トークン) を挿入します: 「DEC-004 (msg 7): pnpm のみ [確認済み]」
…または、何も一致しない場合: 「見つかりません」 — したがって、エージェントは答えを考え出す代わりにそう言います。
これが信頼できる理由は 2 つあります。回答の出所 (決定 ID、メッセージ) が引用されているということです。
番号、ファイル）、幻覚ではなく知らないと認めます。高価なもの
セマンティック ステップは、安価なキーワード ステップが不足した場合にのみ実行されます。そのため、通常のメッセージでは、
リコール層はサイレントかつ無料です。
セマンティック検索はオプションです。ライブラリを 1 つインストールした場合にのみオンになります
( uv pip install speech-transformers 、または pip install --use

r センテンストランスフォーマー )。なし
それは、リコールはキーワード モードで動作し、すべてが引き続き実行されます。埋め込みモデルはローカルで実行されます
CPU — どちらの方法でも API コストはゼロです。
別のフック (ゴール_convergence.ps1、各ターン後) により、ROOT ゴールに対するスコアが維持されます。それ
アクティブな決定、開いているブロッカー、および最近の作業がまだ目標と重なっているかどうかを読み取り、
1 行のステータス ( ON-TRACK 、 DRIFTING 、または BLOCKED) を DOCS/GOAL_STATUS.md に書き込みます。
変化したときにのみ表面化します。これは安価なコード プロキシ (トークンゼロ) であるため、早期警告となります。
判決ではなく旗。本当の「本当にそこにいるのか？」エージェントに直接尋ねた判断
マイルストーン。
長いプロジェクトは実際には意思決定の木です。1 つの目標、いくつかの選択肢がある分岐点、そしてあなたが選ぶものです。
1 つは新しいトランクとなり、再び分岐します。マークダウンはそれを読むには悪い形式です —
フラット化されたリストでは、どのブランチがどのフォークから来たのかが失われます。したがって、テンプレートには次のものも保持されます
意思決定を実際に見ることができるツリーとして表示します。
エージェントが行うすべての実際の決定は、DOCS/_raw/decions.jsonl に追加されます (ユーザーは
送信元のメッセージ番号、テーブルにあったオプション、選択されたオプション、および git
その時点でコミットします）。各ターンの後、フックは 3 つのビューを再描画します。すべて純粋なスクリプトであり、ゼロです。
モデルトークン:
DOCS/decion_tree.txt テキストとしての全体像: 左側の「主な目標」の背骨、すべての部分
そこから分岐する決断、広がっていく選択肢、選ばれたもの
ゴールチェックボックスまでマークが付いています。コードで描画されるため、レイアウトは次のようになります。
正確で決してずれることはありません。
DOCS/decion_tree/msg_*.svg メッセージあたり 1 つの小さなきれいな画像 (きれいにレンダリングされるため)
小さいです）。追加専用 — フォルダーはあなたの履歴です。
DOCS/decion_tree_FULL.txt ツリー形式のすべてのユーザー メッセージ。それぞれに決定のタグが付けられます。
または「（決定なし）」が生成されました。完全な

タイムライン。
DOCS/decion_tree_history/ 各再描画前のテキスト ビューのタイムスタンプ付きスナップショット。
テキストの全体像は次のようになります。
[ROOT] 主な目標: セッション間で作業を失わないようにする
|
+-- {MSG 3} セッションリカバリ (manual-copy) <PICKED: Bridge-tool> (ignore)
+-- {MSG 7} テンプレート戦略 (メモリのみ) <PICKED: ガバナンス> (ハイブリッド)
+-- {MSG 12} 修正リコール (命令) <PICKED: 注入フック>
v
[ルート] 最終ゴール予想 → |ゴールチェック|私たちはどれくらい近いですか？
(ツリーには、実際の決定が行われたメッセージ (フォーク) のみが表示されます。すべてのメッセージ、
決定するかどうかは、まだ DOCS/_raw/user_messages.txt と完全なタイムラインにあります。)
その写真はあなたのためのものです。しかし、これは、曖昧さなくエージェントを指示する方法でもあります。代わりに
「そのことを決めた場所に戻る」の場合、次のノードを指します。
あなた: 「DEC-003 は間違った呼び出しでした。それにロールバックしてください。」
エージェント:hooks/rollback_to_decion.ps1 -Id DEC-003 -適用
→ git-reverts その決定の保存されたコミットに戻り、ツリーを再描画し、後でマークします
決定が優先されました。決定的 — コミット ハッシュが唯一の真実の情報源です。
決定 ID ( -Id DEC-003 ) またはメッセージ番号 ( -Msg 48 ) によってポイントできます。どちらも次のように解決されます。
1 つの正確なコミットなので、エージェントが推測することは何もありません。 (ロールバックにはプロジェクトで git が必要です。
エージェントは申請前に必ずプレビューを行います。)
<ルート>/
§─ CLAUDE.md / AGENTS.md エージェントが自動ロードするルールブック
§─ .claude/ または .codex/ はロギング + インジェクションフックを接続します
§─ フック/
│ §─ log_user_message.ps1 は、すべてのメッセージを逐語的に保存します (+ 番号を付けます)。
│ §─ inject_context.ps1 は開始時 / 圧縮後にスパインを再注入します
│ §─ inject_on_prompt.ps1 は、すべてのメッセージにアクティブなルールを挿入します。
│ §─ inject_decions_preedit.ps1 は、編集の直前にアクティブなルールを挿入します。
│

§─recall.ps1 はルックバック メッセージ (キーワード + セマンティック) で過去を検索します。
│ §─ embed.py オプションのローカル セマンティック エンベッダー ($0、意味別検索)
│ §─index_semantic.ps1 は、セマンティック リコールのために新しいコンテンツに増分インデックスを作成します
│ §─ Record_decion.ps1 は決定 (msg#、オプション、選択された、git commit) を記録します。
│ §─ render_decion_tree.ps1 は、テキスト + メッセージごとの SVG + 完全なタイムラインを描画します
│ §─ rollback_to_decion.ps1 "DEC-X にロールバック" → git-revert + re-route ツリー
│ §─ goal_convergence.ps1 は ROOT 目標に対する進捗状況をスコアします
│ §─ verify_project_setup.ps1 は、必要なすべてのファイルが存在するかどうかを確認します
│ └─ verify_governance.ps1 はルールが破壊されていないことを確認します
━─ DOCS/
§─ INDEX.md すべてのドキュメントのマップ + 競合でどちらが勝つか
§─ CURRENT_STATE.md 現在確認されている内容 (+ E0 ～ E5 の凡例)
§─ REQUIREMENTS.md テスト可能なユーザーニーズ (REQ-XXX)
§─ DECISIONS.md アーキテクチャの選択とその理由 (DEC-XXX)
§─ FAILURE_REGISTRY.md 再発するバグ + 回帰テスト (FAIL-XXX)
§─ ANTI_DRIFT_PROTOCOL.md ショート ループ、3 ストライク、サイレント フォールバックなし
§─ CHANGE_POLICY.md 生のリクエスト → REQ → 証拠 → 1 つのコミット → 記録
§─ CHANGE_RECORD_TEMPLATE.md
§─ GIT_RUNBOOK.md 安全なコミット/ブランチ/ロールバック
§─ HANDOVER_RUNBOOK.md ゼロコンテキスト オペレーター ガイド
§─ STARTUP_MESSAGE.md はセッション開始時に貼り付けを要求します
§─ BOOTSTRAP_PROMPT.md このシステムを新しいプロジェクトにインストールするためのプロンプト
§─ PROJECT_LOG.md 追記

[切り捨てられた]

## Original Extract

Stop Claude Code and Codex forgetting your project mid-build. Drop-in memory + governance templates that write your decisions to disk and inject them back every session, even after a context compaction. - budhasantosh010/ai-coding-project-templates

GitHub - budhasantosh010/ai-coding-project-templates: Stop Claude Code and Codex forgetting your project mid-build. Drop-in memory + governance templates that write your decisions to disk and inject them back every session, even after a context compaction. · GitHub
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
budhasantosh010
/
ai-coding-project-templates
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
budhasantosh010/ai-coding-project-templates
main Branches Tags Go to file Code Open more actions menu Folders and files
11 Commits 11 Commits claude-project-template claude-project-template codex-project-template codex-project-template .gitignore .gitignore LICENSE LICENSE README.md README.md START_HERE.md START_HERE.md View all files Repository files navigation
A coding agent forgets. Halfway through a build the context compacts, and suddenly it's
re-asking things you settled an hour ago or quietly ignoring a rule you set on day one. These
are two drop-in templates — one for Claude Code, one for Codex — that fix that by keeping the
project's memory in files and feeding it back to the agent automatically.
ai-coding-project-templates/
├─ claude-project-template/ for Claude Code (rulebook = CLAUDE.md, hook via .claude/settings.json)
└─ codex-project-template/ for OpenAI Codex (rulebook = AGENTS.md, hook via .codex/hooks.json)
Both folders are the same system. The only differences are the rulebook filename and how each
agent registers hooks.
How recall is forced, not hoped-for
Looking things up by meaning (recall)
The decision tree (see it, and roll back to it)
Optional companion: graphify
Does graphify auto-fire from these templates?
Install once, build per project
You probably don't need the paid step
Four things that go wrong on any long-running agent session, and what the template does about
each:
Without With
─────── ────
Forgets between sessions; you Every message, decision and change lands in DOCS/,
re-explain the project constantly. so a fresh session reloads the full context.
"Done" means "the code exists," "Done" means a passing test or a real run (E0–E5),
even if it never ran. not just code that imports.
Same error retried forever. Three strikes on the same failure → stop and diagnose.
Silent fallbacks; lost reasoning. Fallbacks must be stated. Decisions and failures are
logged with the test that stops them returning.
Pick your agent
You use
Open this folder
Rulebook
Claude Code (CLI / VS Code / desktop)
claude-project-template/
CLAUDE.md
OpenAI Codex (CLI / VS Code / desktop)
codex-project-template/
AGENTS.md
Install (per project)
Do this once per project, when you copy the template in:
Copy your folder's contents into the project root.
Fill in the <PROJECT_NAME> / <PROJECT_ROOT> / <OWNER> / <DATE> placeholders.
Open the project in your agent and trust its hooks.
Run both hooks/verify_*.ps1 checks — they should all pass.
Paste the first-session block from DOCS/STARTUP_MESSAGE.md into the first chat. (The
prompts are also collected in START_HERE.md .)
How recall is forced, not hoped-for
Saving context to a file is the easy half. The hard half is getting the agent to actually look
at it once the conversation has moved on. Telling it "check DECISIONS.md when unsure" is a sign
on the wall — it can walk right past it.
So the templates don't rely on that. Three hooks read your files and push the relevant bits
back into the agent's view at the moments it tends to forget:
when hook what it puts back
──── ──── ─────────────────
session starts or compacts inject_context CURRENT_STATE + the DEC/REQ/FAIL list
you send a message inject_on_prompt the active rules + "read the transcript"
right before an edit inject_decisions_preedit the active DEC/REQ rules, at the edit
The point: "we use pnpm, not npm" stops depending on the model remembering it. The rule is on
screen at session start, on every message, and again right before the agent writes the install
command. A hook can't be skipped, so the information is guaranteed to be there. And they all
fail safe — if a hook errors it prints nothing and never blocks your session.
Looking things up by meaning (recall)
Re-injecting your rules handles the recent stuff. But what about "what did we decide about
pricing three months ago?" — buried in a long transcript, maybe phrased with different words.
That's what the recall hook ( recall.ps1 , on every message) does, and it's built to cost almost
nothing:
1. Does the message even look back? ("remember…", "earlier", "that bug", "it/this")
No → do nothing. 0 tokens. (most messages)
Yes → continue.
2. Resolve vague words: "find it" → the most-mentioned recent thing.
3. Search, cheapest first:
tier 1 keyword over decisions + transcript (free, instant)
tier 2 semantic by meaning — ONLY if tier 1 weak (local model; "login" finds "auth")
4. Verify: if the hit names a file, grep the CURRENT file → CONFIRMED or STALE.
5. Inject a tiny cited pointer (~40 tokens): "DEC-004 (msg 7): pnpm only [CONFIRMED]"
…or, if nothing matched: "not found" — so the agent says so instead of inventing an answer.
Two things make this trustworthy: it cites where the answer came from (decision id, message
number, file), and it admits when it doesn't know rather than hallucinating. The expensive
semantic step only runs when the cheap keyword step comes up short — so on a normal message the
recall layer is silent and free.
Semantic search is optional. It switches on only if you install one library
( uv pip install sentence-transformers , or pip install --user sentence-transformers ). Without
it, recall works in keyword mode and everything still runs. The embedding model runs locally on
your CPU — zero API cost either way.
A separate hook ( goal_convergence.ps1 , after each turn) keeps score against your ROOT goal. It
reads the active decisions, open blockers, and whether recent work still overlaps the goal, then
writes a one-line status — ON-TRACK , DRIFTING , or BLOCKED — to DOCS/GOAL_STATUS.md , and
surfaces it only when it changes. It's a cheap code proxy (zero tokens), so it's an early-warning
flag, not a verdict; for the real "are we actually there?" judgment you ask the agent directly at
a milestone.
A long project is really a tree of decisions: one goal, a fork with a few options, you pick
one, that becomes the new trunk, it forks again. Markdown is a bad shape for reading that —
a flattened list loses which branch came from which fork. So the template also keeps the
decisions as a tree you can actually look at.
Every real decision the agent makes is appended to DOCS/_raw/decisions.jsonl (with the user
message number it came from, the options that were on the table, which was chosen, and the git
commit at that moment). After each turn a hook redraws three views — all pure scripting, zero
model tokens:
DOCS/decision_tree.txt the big picture as text: a left "main goal" spine, every
decision branching off it, options fanning out, the picked one
marked, down to a goal-check box. Code-drawn, so the layout is
exact and never shifts.
DOCS/decision_tree/msg_*.svg one small clean picture PER message (renders cleanly because
it's small). Append-only — the folder IS your history.
DOCS/decision_tree_FULL.txt every user message in tree shape, each tagged with the decision
it produced or "(no decision)". The complete timeline.
DOCS/decision_tree_history/ timestamped snapshots of the text views before each redraw.
The text big-picture looks like this:
[ROOT] MAIN GOAL: never lose work across sessions
|
+-- {MSG 3} session recovery ( manual-copy ) <PICKED: bridge-tool> ( ignore )
+-- {MSG 7} template strategy ( memory-only ) <PICKED: governance> ( hybrid )
+-- {MSG 12} fix recall ( instructions ) <PICKED: injection-hooks>
v
[ROOT] EXPECTED FINAL GOAL → |GOAL CHECK| how close are we?
(The tree shows only the messages where a real decision was made — the forks. Every message,
decision or not, is still in DOCS/_raw/user_messages.txt and in the FULL timeline.)
The picture is for you. But it's also how you direct the agent without ambiguity . Instead
of "go back to where we decided that thing," you point at a node:
You: "DEC-003 was the wrong call — roll back to it."
Agent: hooks/rollback_to_decision.ps1 -Id DEC-003 -Apply
→ git-reverts to that decision's stored commit, redraws the tree, marks later
decisions superseded. Deterministic — the commit hash is the single source of truth.
You can point by decision id ( -Id DEC-003 ) or by message number ( -Msg 48 ) — both resolve to
one exact commit, so there's nothing for the agent to guess. (Rollback needs git in the project;
the agent always previews before applying.)
<root>/
├─ CLAUDE.md / AGENTS.md the rulebook the agent auto-loads
├─ .claude/ or .codex/ wires up the logging + injection hooks
├─ hooks/
│ ├─ log_user_message.ps1 saves every message word-for-word (+ numbers them)
│ ├─ inject_context.ps1 re-injects the spine on start / after compaction
│ ├─ inject_on_prompt.ps1 injects active rules with every message
│ ├─ inject_decisions_preedit.ps1 injects active rules right before an edit
│ ├─ recall.ps1 looks up the past on a look-back message (keyword + semantic)
│ ├─ embed.py optional local semantic embedder ($0, by-meaning search)
│ ├─ index_semantic.ps1 incrementally indexes new content for semantic recall
│ ├─ record_decision.ps1 logs a decision (msg#, options, chosen, git commit)
│ ├─ render_decision_tree.ps1 draws the text + per-message SVG + FULL timeline
│ ├─ rollback_to_decision.ps1 "roll back to DEC-X" → git-revert + re-route tree
│ ├─ goal_convergence.ps1 scores progress vs the ROOT goal
│ ├─ verify_project_setup.ps1 checks every required file exists
│ └─ verify_governance.ps1 checks the rules haven't been gutted
└─ DOCS/
├─ INDEX.md map of all docs + which one wins in a conflict
├─ CURRENT_STATE.md what's verified true right now (+ the E0–E5 legend)
├─ REQUIREMENTS.md testable user needs (REQ-XXX)
├─ DECISIONS.md architecture choices and why (DEC-XXX)
├─ FAILURE_REGISTRY.md recurring bugs + the regression test (FAIL-XXX)
├─ ANTI_DRIFT_PROTOCOL.md short loop, three-strike, no silent fallback
├─ CHANGE_POLICY.md raw request → REQ → evidence → one commit → record
├─ CHANGE_RECORD_TEMPLATE.md
├─ GIT_RUNBOOK.md safe commit / branch / rollback
├─ HANDOVER_RUNBOOK.md zero-context operator guide
├─ STARTUP_MESSAGE.md prompts to paste at session start
├─ BOOTSTRAP_PROMPT.md prompt to install this system into a fresh project
├─ PROJECT_LOG.md append-o

[truncated]
