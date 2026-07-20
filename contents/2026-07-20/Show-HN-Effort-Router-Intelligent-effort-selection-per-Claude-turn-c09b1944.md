---
source: "https://github.com/cfitzgerald-pd/effort-router"
hn_url: "https://news.ycombinator.com/item?id=48981438"
title: "Show HN: Effort Router: Intelligent /effort selection per Claude turn"
article_title: "GitHub - cfitzgerald-pd/effort-router: Claude plugin that intelligently picks effort level for each conversation. · GitHub"
author: "bennydog224"
captured_at: "2026-07-20T17:23:50Z"
capture_tool: "hn-digest"
hn_id: 48981438
score: 3
comments: 0
posted_at: "2026-07-20T16:53:16Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Effort Router: Intelligent /effort selection per Claude turn

- HN: [48981438](https://news.ycombinator.com/item?id=48981438)
- Source: [github.com](https://github.com/cfitzgerald-pd/effort-router)
- Score: 3
- Comments: 0
- Posted: 2026-07-20T16:53:16Z

## Translation

タイトル: HN を表示: 努力ルーター: クロード ターンごとのインテリジェントな / 努力選択
記事のタイトル: GitHub - cfitzgerald-pd/effort-router: 各会話のエフォート レベルをインテリジェントに選択するクロード プラグイン。 · GitHub
説明: 各会話の努力レベルをインテリジェントに選択する Claude プラグイン。 - cfitzgerald-pd/エフォートルーター
HN テキスト: 動的な努力値の選択は、OpenCode などの他のツールでは可能ですが、Claude Code には存在しません。会話ターンごとに動的エフォートを設定することは、次のことを意味します。 - 簡単なクエリで消費されるトークンが少なくなります (例: コンボのデフォルトであるため、低エフォートで実行できるものは xhigh に設定されます) - 間違いを回避できます: 高い労力を必要とするプロンプト (リファクタリングやセキュリティ脅威モデリングなど) は、中ではなく高または xhigh に設定され、間違いの可能性を減らし、無駄なトークンを減らします。私は個人的なニーズからこれを作成しましたが、これはゲームを変えるものでした。私のクロード セッションは、やり取りが少なくなり、生産性が大幅に向上しました。ぜひ試してみて、お気軽に貢献してください。

記事本文:
GitHub - cfitzgerald-pd/effort-router: 各会話のエフォート レベルをインテリジェントに選択する Claude プラグイン。 · GitHub
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
フィッツジェラルド-pd
/
エフォートルーター
公共
通知
ch にサインインする必要があります

アンジュ通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .claude-plugin .claude-plugin フック フック スキル スキル README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
二度と手動で /effort を設定しないでください。努力ルータはすべてのリクエストを分類し、クロード コードの推論努力を順番にダイヤルして一致させます。タイプミスの修正には低め、セキュリティが重要な作業には最大となります。トークンの使用状況が随時記録されるため、何が保存され、何が誤ってルーティングされたのかを正確に確認できます。
工数はセッション全体の設定です。会話中にプログラムで変更する方法はありません。
デフォルト値を高くすると、些細な編集でトークンが無駄に消費されます。デフォルト値が低いと、難しい問題が考慮されなくなります。
5 つの小さなスキル (努力レベルごとに 1 つ) が付属します。スキルの前衛努力値のオーバーライドは、ターンの残りの期間に適用されます。
階層: 低 (タイプミス、名前変更)、中 (単一ファイルの追加、既知の原因の修正)、高 (複数ファイルの機能、リファクタリング、原因不明のデバッグ)、xhigh (アーキテクチャ、同時実行のバグ、移行)、最大 (セキュリティに依存するもの、失敗後の再試行、明示的なハードウェア要求)。
2 つのフックによってルーティングが強制されます。SessionStart は継続的な「すべてのリクエストを分類」命令を挿入し、UserPromptSubmit はそれをすべてのプロンプト (~20 トークン) ごとに再固定するため、長時間のセッション、再開されたセッション、または圧縮されたセッションでも存続します。
スキル フロントマター モデルのオーバーライドは実行時に無視されるため (claude-code#45191)、プラグインはエフォートのみをルーティングします。
Stop フックは、完了した各ターン (層、モデル、出力トークン) を ~/.claude/effort-router/usage/<session_id>.jsonl に記録し、ターンごとに更新されます。
エフォートレポートスキルは、シフトダウンによって節約されたトークンを推定し、誤ったルートの可能性があることをフラグします。つまり、膨れ上がった低/中ターンと、小さなままであった高/x高/最大ターンです。
それも壊れる

層、日、モデル、セッションごとに使用量が減少します。
取り組みレポートからわかること
3日間使用した後の実際のレポート:
== シフトダウンによる節約 ==
6 つのロールート ターンで 36,503 個の出力トークンが保存されました
== コンテキスト ==
層ターン出力トークンの平均/ターン
低い 6 5222 870
中 8 55633 6954
高13 274807 21139
x高 4 127644 31911
配線されていない 36 657742 18271
可視化される 3 つのこと:
ルーティングが元を取れるかどうか。ここでのターンあたりの平均出力は、低レベルで 870 トークンから高レベルで 31,911 トークンの範囲にあり、最大 35 倍のスプレッドになります。節約ラインは、中層の平均で 6 つの低ルート ターンのコストを推定します。これは反事実であるため、説明としてではなく、方向性として扱ってください。
ミスルートは一見の価値あり。フラグはトークンサイズのヒューリスティックであり、判定ではありません。上記では、中程度のルートの 1 つのターンで 36,578 個のトークンが消費されました。これは、高位にルートされるべきマルチリポジトリのタスクです。おそらく推論が足りずに失敗したのだろう。 「オーバールートされた」ハイターンはほんのわずかでしたが、高い労力での小さなターンも、短い答えでは難しい質問になる可能性があります。
そもそもルーティングが行われているかどうか。ここでは、プラグインのインストールより前のセッションが再起動されるまで古いフックを保持しているため、未配線が優勢です (67 ターン中 36 ターン)。インストール後に未配線の共有が多い場合は、ナッジが起動していないことを意味します。セッションを再起動してください。
主なコストはルーティング自体です。つまり、各ターンの開始時に 1 回の追加の努力 * スキル呼び出し、つまり実際の作業が開始される前に 1 回の追加のモデル往復 (通常は約 1 ～ 3 秒) です。
ルーティングは、モデル ステップごとではなく、ユーザー プロンプトごとに 1 回実行されます。ターンの中間ツールの呼び出しと応答は影響を受けません。分類は、モデルが人間による新しい入力を取得した場合にのみ行われます。
フックはローカル スクリプトであり、事実上何も追加しません。SessionStart と UserPromptSubmit は静的命令 (ミリ秒、5 秒のタイムアウト) をエコーし​​ます。

そして、応答がすでに配信された後に、ロギングの停止フックが実行されます。
通常、シフトダウンされたターンは時間を取り戻します。簡単なタスクに低労力で取り組むと、多大な労力を費やすデフォルトよりも推論に費やす時間がはるかに少なくなります。
最後の 1 ターンをオーバーライドします。毎ターン、最初から再分類されます。
プロンプトごとのナッジには最大 20 トークンのコストがかかり、それぞれのトークンはトランスクリプトに残ります。
利用可能な層はセッションのモデルによって異なります (たとえば、Opus 4.6 と Sonnet 4.6 には xhigh がありません)。
オーバーライドは UI には表示されません。/effort メニューには引き続きセッション設定が表示されます。ルーティングの証拠は、ターン開始時の努力値*スキルの発動です。
インストール時にすでに実行されていたセッションは、再起動されるまで使用状況をログに記録しません。
各会話の労力レベルをインテリジェントに選択する Claude プラグイン。
読み込み中にエラーが発生しました。このページをリロードしてください。
1
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Claude plugin that intelligently picks effort level for each conversation. - cfitzgerald-pd/effort-router

Dynamic effort selection, while possible in some other tools like OpenCode, does not exist in Claude Code. Having dynamic effort set per conversation turn means: - Less tokens burned on easy queries (e.g. something low /effort can do goes to xhigh because it's the convo default) - You avoid mistakes: Prompts requiring high effort (like a refactor or security threat modeling) go to high or xhigh instead of medium to reduce likelihood of mistakes and reduce wasted tokens. I created this out of a personal need and it's been game-changing. My Claude sessions are way more productive with less back/forth. Give it a try and feel free to contribute!

GitHub - cfitzgerald-pd/effort-router: Claude plugin that intelligently picks effort level for each conversation. · GitHub
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
cfitzgerald-pd
/
effort-router
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .claude-plugin .claude-plugin hooks hooks skills skills README.md README.md View all files Repository files navigation
Never set /effort by hand again. effort-router classifies every request and dials Claude Code's reasoning effort to match, turn by turn: low for a typo fix, max for security-sensitive work. It logs token usage as it goes, so you can see exactly what it saved and what it misrouted.
Effort is a session-wide setting; there is no programmatic way to change it mid-conversation.
A high default wastes tokens on trivial edits; a low default underthinks hard problems.
Ships five tiny skills, one per effort tier. A skill's frontmatter effort override applies for the rest of the turn.
Tiers: low (typos, renames), medium (single-file additions, known-cause fixes), high (multi-file features, refactors, unknown-cause debugging), xhigh (architecture, concurrency bugs, migrations), max (anything security-sensitive, retries after a failure, explicit think-hard requests).
Two hooks enforce the routing: SessionStart injects a standing "classify every request" instruction, and UserPromptSubmit re-anchors it on every prompt (~20 tokens) so it survives long, resumed, or compacted sessions.
Skill frontmatter model overrides are ignored at runtime (claude-code#45191), so the plugin routes effort only.
A Stop hook logs each completed turn (tier, model, output tokens) to ~/.claude/effort-router/usage/<session_id>.jsonl , refreshed after every turn.
The effort-report skill estimates tokens saved by downshifting and flags likely misroutes: low/medium turns that ballooned, and high/xhigh/max turns that stayed tiny.
It also breaks usage down by tier, day, model, and session.
What the effort report tells you
A real report after three days of use:
== savings from downshifting ==
36,503 output tokens saved across 6 low-routed turns
== context ==
tier turns output tokens mean/turn
low 6 5222 870
medium 8 55633 6954
high 13 274807 21139
xhigh 4 127644 31911
unrouted 36 657742 18271
Three things it makes visible:
Whether routing pays for itself. Mean output per turn here ranges from 870 tokens at low to 31,911 at xhigh — a ~35x spread. The savings line estimates what the 6 low-routed turns would have cost at the medium-tier mean. It's a counterfactual, so treat it as directional, not accounting.
Misroutes worth a look. The flags are token-size heuristics, not verdicts. Above, one medium-routed turn burned 36,578 tokens — a multi-repo task that should have routed high; it likely flailed for lack of reasoning. The "overrouted" high turns were tiny, but a small turn at high effort can also be a hard question with a short answer.
Whether routing is happening at all. Here unrouted dominates (36 of 67 turns) because sessions that predate the plugin install keep their old hooks until restarted. A high unrouted share after installation means the nudge isn't firing — restart those sessions.
The main cost is the routing itself: one extra effort-* skill invocation at the start of each turn, i.e. one additional model round trip before real work begins (typically ~1-3 seconds).
Routing runs once per user prompt, not on every model step: a turn's intermediate tool calls and responses are unaffected. Classification only happens when the model gets new human input.
The hooks are local scripts and add effectively nothing: SessionStart and UserPromptSubmit echo a static instruction (milliseconds, 5 s timeout), and the Stop logging hook runs after the response is already delivered.
Downshifted turns usually win the time back: low effort on a trivial task spends far less time reasoning than a high-effort default would.
Overrides last one turn; every turn re-classifies from scratch.
The per-prompt nudge costs ~20 tokens, and each one stays in the transcript.
Available tiers depend on the session's model (Opus 4.6 and Sonnet 4.6 lack xhigh , for example).
The override is invisible in the UI: the /effort menu still shows the session setting. Proof of routing is the effort-* skill invocation at the start of the turn.
Sessions already running when you install won't log usage until restarted.
Claude plugin that intelligently picks effort level for each conversation.
There was an error while loading. Please reload this page .
1
fork
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
