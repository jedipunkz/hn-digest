---
source: "https://github.com/kaiwutech-TW/flightwake"
hn_url: "https://news.ycombinator.com/item?id=48964105"
title: "Flightwake – a flight recorder for AI coding agents, not a navigator"
article_title: "GitHub - kaiwutech-TW/flightwake: Flight-recorder style work framework for strong AI coding agents — records follow work, they don't lead it. ✈️ · GitHub"
author: "kaiwuTW"
captured_at: "2026-07-19T01:45:44Z"
capture_tool: "hn-digest"
hn_id: 48964105
score: 1
comments: 0
posted_at: "2026-07-19T01:14:24Z"
tags:
  - hacker-news
  - translated
---

# Flightwake – a flight recorder for AI coding agents, not a navigator

- HN: [48964105](https://news.ycombinator.com/item?id=48964105)
- Source: [github.com](https://github.com/kaiwutech-TW/flightwake)
- Score: 1
- Comments: 0
- Posted: 2026-07-19T01:14:24Z

## Translation

タイトル: Flightwake – ナビゲーターではなく、AI コーディング エージェント用のフライト レコーダー
記事のタイトル: GitHub - kaiwutech-TW/flightwake: 強力な AI コーディング エージェントのためのフライト レコーダー スタイルの作業フレームワーク — 記録は作業に従うものであり、作業を主導するものではありません。 ✈️ · GitHub
説明: 強力な AI コーディング エージェントのためのフライト レコーダー スタイルの作業フレームワーク — 記録は作業に従うものであり、作業を先導するものではありません。 ✈️ - カイウテック-TW/フライトウェイク

記事本文:
GitHub - kaiwutech-TW/flightwake: 強力な AI コーディング エージェントのためのフライト レコーダー スタイルの作業フレームワーク — 記録は作業に従うものであり、作業を主導するものではありません。 ✈️ · GitHub
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
カイウテック-TW
/
飛行航跡
公共
通知
nを変更するにはサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
56 コミット 56 コミット .claude .claude .flightwake .flightwake .github/ workflows .github/ workflows bin bin docs docs フック フック スキル スキル スニペット スニペット テンプレート テンプレート テスト テスト ツール ツール CLAUDE.md CLAUDE.md ライセンス ライセンス README.ja.md README.ja.md README.md README.md README.zh-CN.md README.zh-CN.md README.zh-TW.md README.zh-TW.md SECURITY.md SECURITY.md package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
記録は、仕事が自然に残す飛行機雲であり、離陸前に提出しなければならない飛行計画ではありません。
強力な AI コーディング エージェント (Claude Fable 5 世代以降) のための超軽量の作業記録フレームワーク。実行時の依存関係はゼロ、純粋な Markdown、すべてが git 内に存在します。
cd あなたのリポジトリ
npx Flightwake init # 既存のインストールをアップグレードします: add --force
init は .flightwake/ (テンプレート + Stop フック) を作成し、4 つのスキルを .claude/skills/ にコピーし、Stop フックを .claude/settings.json にマージし、検出されたエージェント指示ファイル (CLAUDE.md / AGENTS.md / GEMINI.md) にトリガー義務テーブル (<!-- Flightwake:begin/end --> マーカーでラップされた) を追加します (CLAUDE.md / AGENTS.md / GEMINI.md)。存在しない場合は、AGENTS.md が作成されます。 --agents=claude,codex,gemini は明示的に選択します)。純粋なファイルのコピー、実行時の依存関係なし (ノード 18 以上はインストール時およびフックによってのみ使用されます)。ユーザーデータ (STATE/DECISIONS/TRAPS) は決して上書きされません。 --force は、フレームワークが所有するファイルのみを更新します。
注: CLI 出力とインストールされているテンプレート/スキルは現在繁体字中国語で書かれています (チームは zh-TW で動作するドッグフーディングを行っています)。デフォルトは完全に英語になる予定です。フレームワークの仕組みは言語に依存せず、エージェントは次のいずれかを読み取ります。

大丈夫ですよ。
Claude Code セッションを開き、「initialize STATE with /fw-record」と言います。モデルはリポジトリの現在の状況を最初の STATE に書き込みます。
git add .flightwake .claude CLAUDE.md && git commit
それ以降のすべてのセッションは、以下の日次ループに従います。
あなた (およびモデル) が覚えておく必要があることは 1 つだけです。 /fw-coldstart で作業を開始することです。モデルは他のすべての義務自体をトリガーします。義務テーブルはすでに命令ファイル内にあり、強力なモデルはそれを読み取り、それを尊重します。典型的なセッション:
あなた: /fw-coldstart
モデル: (STATE + 最新レコードの読み取り、約 1 分)
「最後のセッションは X、健康状態は緑、次のエントリ ポイントは Y に到達しました。
未確認の変更: なし。 Yから迎えに行く？」
あなた：はい、行きます
モデル: (直接作業を開始します。選択肢を閉じる決定を下します →
DECISIONS に 1 行追加。明らかではないトラップにヒット → /fw-trap)
あなた: まとめます
モデル: (/fw-record: フライト レコードを書き込み、STATE を更新し、
機密情報のセルフチェック)
まとめを忘れていませんか？ STATE が 3 コミット以上遅れている場合、Stop フックはセッションが終了する前に一度ブロックされ、通知されます。 --ci は、他のエージェントや人間の協力者にも同じゲートをもたらします。マルチセッション構築の場合、停止する前に「ハンドオフ」と言うと、モデルが /fw-handoff を実行します。
あなたが見なければならない唯一のもの
STATE の健康状態が正直かどうか (緑/黄/赤)。このフレームワークには単一の品質指標があります。それは、/fw-coldstart の後に新しいセッションが安全なテイクオーバーに達するまでにかかる時間です。これに 5 分以上かかる場合、記録は劣化しています。レコード数、フォーマット準拠など、他のすべては重要ではありません。
実際にどのように見えるか見てみましょう
このリポジトリは独自のフレームワークをドッグフードしています。.flightwake/ には実際の状態、決定、記録が含まれています。ギャップ リストからオープンソースの立ち上げまでのすべてのステップがそこに記録されます。それ'

まさに、インストール後にリポジトリ内で成長するものです。
強力なモデルを扱うのは初めてですか? docs/workflow.md は、初心者向けのメインライン、クロード コードのベテラン向けの高度なフォールドなど、各ポイントで何を行うか、モデルに何を言うべきかを示したステージ マップです。 (現時点では zh-TW; 英語は v0.9 i18n パスで上陸します。)
Fable 5 クラスのモデルは、その仕事のやり方を教える必要はありません。しかし、モデルがどんなに強くなってもできないことが 4 つあります。これらは構造的なものであり、モデルが改良されても消えるものではないからです。
セッションは終了します。コンテキストは有限です。作業が複数のセッションにまたがる場合、メモリはゼロにリセットされます。記録がなければ、すべてのテイクオーバーは Git 考古学での発掘になります。強力なモデルはより速く掘るだけで、掘削をスキップすることはできません。
Git は理由ではなく、何を記録します。コミットは、何が変化したかを示しますが、「他のパスが選択されなかった理由」や「そのトラップの根本原因」は決して示しません。これらは、次のセッション (または次のエージェント) にとって最も高価な 2 つの情報となります。
長いセッションでは規律が漂います。 「テストの実行前に完了したと報告」、「検証証拠を残さずに本番環境にアクセス」 - これらのスライドはモデルのインテリジェンスとは何の関係もなく、モデルの外部にガードが必要です。
エージェントは状態を共有しません。クロード、コーデックス、ジェミニ、そして人間のチームメイトはそれぞれ独自の世界を見ています。状態は、git に追加されると初めて全員のものになります。
したがって、フライトウェイクは知性ではなく粘り強さと規律を補うものです。その精神の祖先は GSD です。GSD はナビゲーション (あらゆるステップをターンバイターンで案内) です。 Flightwake は車載カメラ + 警告灯 + 道路標識です。強力なモデルは自動で運転するため、フレームワークは次の 3 つのことだけを行います。
車載カメラ: 決定、発見、検証証拠、事後記録 ( records/ 、 DECISIONS.md 、 TRAPS.md )
警告灯: モデルの強度に関係なくハードガード

(「完了」する前にテストは緑色になります。製品の変更は検証証拠を残す必要があります。破壊的な操作は最初に確認が必要です)
道路標識 : どのセッションも停止する可能性があります。次のセッションは STATE.md を読み取り、2 分以内に安全に引き継ぎます。
起源は実際の 3 日間のセッション (2026 年 7 月 15 ～ 17 日: 2 つのリポジトリ、19 のコミット、4 つの cron ジョブ、2 つの深いバグ修正 — 事前の計画なし、脱線なし) でした。強力なモデルにはナビゲーションが必要ないことが証明されましたが、次のセッションを可能にするために残されたもの (概要/コンテキスト/メモリ ファイル) はすべてその場で即興で作成されました。 Flightwake は、その即興演奏をインストール可能な慣習に変えます。
基本原則: 記録は仕事に従う - 記録は仕事を先導するものではない
GSD は段階主導型 (調査→計画→実行→ゲートの検証) です。 Flightwake はトリガー駆動です (イベントによって義務が発生します)。
エスカレーション ルール (GSD の逆) : デフォルトではすべてが迅速です。作業を開始するだけです。 「複数のセッションにまたがる構築」のみがフェーズ (1 つの CONTEXT ファイル。計画の分解はモデルのその時点の判断に委ねられます) にエスカレートします。
ファイルレイアウト (ターゲットリポジトリにインストールした後)
あなたのレポ/
§── .flightwake/
│ §── STATE.md # 現在いる場所、次のステップのエントリ (常に短く、常に最新)
│ §── DECISIONS.md # 追加専用の決定ログ (決定ごとに 1 行、理由付き)
│ §── TRAPS.md # トラップ レジストリ (OKF スタイルのフロントマター エントリ)
│ §── TEMPLATE-record.md # 飛行記録テンプレート
│ §──hooks/state-check.mjs # フックの停止: STATE が 3 コミット以上遅れた場合に終了するよう通知します。
│ └── レコード/フライトレコード数 (意味のあるまとめごとに 1 つ)
§── .claude/skills/fw-*/ # 4つのスキル
└── .claude/settings.json # init は Stop フック設定をここにマージします
スキルとストップフックはクロード・コードにとって便利な砂糖です。

.flightwake/ 自体はプレーンな Markdown です。命令ファイルを読み取るエージェントは同じトリガー義務に従うことができます。既存の GSD .planning/ と共存します (古いレコードは履歴アーカイブになります)。
--private は、記録を git の外でローカルのみに保持します。すべての書き込みは .git/info/exclude に登録され (純粋にローカル — リポジトリに痕跡は残りません)、フックは .claude/settings.local.json に移動し、義務テーブルは CLAUDE.local.md に移動します (git で追跡された命令ファイルには一切触れません)。コスト: レコードはリポジトリと共有されず、新しいクローンには再び init --private が必要です。「git では、リポジトリと共有」が Flightwake のデフォルトであり、存在理由です。 --private は、他人のリポジトリ内で個人的に使用するための脱出ハッチです。
アンインストールは、init の固定書き込みスコープを逆にします。スキルとフレームワーク ファイルを削除し、設定から Flightwake の Stop フックを抽出し (他のフックはそのまま残ります)、命令ファイルと .git/info/exclude からマーカー ブロックを削除します (flightwake によって作成されたファイルは、空になると削除されます)。 .flightwake/ はユーザー データであり、デフォルトで保持されます。 --purge だけをアンインストールすると、それも削除されます。
Monorepo ポリシー: リポジトリごとに、git ルートに 1 回インストールします。作業はセッション形式です。セッションは通常、複数のパッケージにまたがり、レコードはセッションの後に続きます。サブディレクトリごとにインストールすると、一連の作業が断片化されたレコードに細断され、「どの状態で読むか」が変わってしまいます。新たなコールドスタートの曖昧さへ。サブディレクトリで init を実行すると停止し、ルートをポイントします。サブモジュールには独自の .git があり、独立したリポジトリとしてカウントされます。トラフィックの多いマルチチームのモノリポジトリで CI の古さチェックで誤検知が発生した場合は、最初に --threshold を調整します。
まず現在のマイルストーンをまとめてから、次のことを行います。
npx Flightwake init — .planning/ と共存します。何も削除されていません
テ

エージェント: 「このリポジトリは GSD から Flightwake に切り替わります。現在の状態について .planning/ を読み取り、/fw-record で .flightwake/STATE.md を初期化します。未完了の項目は次のステップのエントリに移動します。これ以降、.planning/ は履歴アーカイブになります。更新しないでください。」
CLAUDE.md から GSD 独自の命令ブロックを削除 (またはコメントアウト) し、2 つのルールブックがモデルの従順性をめぐって競合しないようにします。
npx Flightwake init --statusline は、クロード コードの下部に永続的なゲージを配置します。
✈️ フライトウェイク │ ●緑 · STATE 2c後方 │ ▓▓░░░░░░░░ 23%
健全性の色 (監視対象の 1 つ)、状態の古さ (Stop フックと同じ rev-list ロジック - ただし、終了時間のリマインダーではなくライブ ゲージとして)、およびコンテキストの使用状況。ゲージは、現在の状態での次のコマンドも示します。セッションは開始されました → → 開工先 /fw-coldstart ; STATE ≥3 は → → /fw-record の後にコミットします。ホット実行中のコンテキスト → → /fw-record → /clear → /fw-coldstart ;すべて健康→沈黙。既存のステータスライン (単一値設定) が上書きされることはなく、リポジトリレベルの設定がユーザーレベルよりも優先されるため、グローバル設定を設定するツールと共存します。
注: プレーンな npx Flightwake init はゲージをインストールしません。これはオプトインです。すでに init なしで実行していますか? npx Flightwake init --statusline を再度実行すると、ゲージが追加されるだけです (他のすべてはすでにインストールされているためスキップされます)。の

[切り捨てられた]

## Original Extract

Flight-recorder style work framework for strong AI coding agents — records follow work, they don't lead it. ✈️ - kaiwutech-TW/flightwake

GitHub - kaiwutech-TW/flightwake: Flight-recorder style work framework for strong AI coding agents — records follow work, they don't lead it. ✈️ · GitHub
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
kaiwutech-TW
/
flightwake
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
56 Commits 56 Commits .claude .claude .flightwake .flightwake .github/ workflows .github/ workflows bin bin docs docs hooks hooks skills skills snippets snippets templates templates test test tools tools CLAUDE.md CLAUDE.md LICENSE LICENSE README.ja.md README.ja.md README.md README.md README.zh-CN.md README.zh-CN.md README.zh-TW.md README.zh-TW.md SECURITY.md SECURITY.md package.json package.json View all files Repository files navigation
Records are the contrail your work naturally leaves behind — not a flight plan you must file before takeoff.
An ultra-lightweight work-recording framework for strong AI coding agents (Claude Fable 5 generation and beyond). Zero runtime dependencies, pure Markdown, everything lives in git.
cd your-repo
npx flightwake init # upgrade an existing install: add --force
init creates .flightwake/ (templates + Stop hook), copies 4 skills into .claude/skills/ , merges the Stop hook into .claude/settings.json , and appends the trigger-obligation table (wrapped in <!-- flightwake:begin/end --> markers) to detected agent instruction files (CLAUDE.md / AGENTS.md / GEMINI.md — whichever exist; if none, it creates AGENTS.md; --agents=claude,codex,gemini selects explicitly). Pure file copying, zero runtime dependencies (Node ≥18 used only at install time and by the hook). User data (STATE/DECISIONS/TRAPS) is never overwritten; --force only updates framework-owned files.
Note: the CLI output and the installed templates/skills are currently in Traditional Chinese (the team dogfooding it works in zh-TW). Full English defaults are planned — the framework mechanics are language-independent, and your agent reads either just fine.
Open a Claude Code session and say " initialize STATE with /fw-record " — the model writes the repo's current situation into the first STATE
git add .flightwake .claude CLAUDE.md && git commit
Every session after that follows the daily loop below
You (and the model) only need to remember one thing: start work with /fw-coldstart ; the model triggers every other obligation itself — the obligation table is already in the instruction file, and strong models both read it and honor it. A typical session:
You: /fw-coldstart
Model: (reads STATE + the latest record, ~1 minute)
"Last session got to X, health green, next entry point is Y.
Unverified changes: none. Pick up from Y?"
You: Yes, go
Model: (starts working directly. Makes a decision that closes off options →
one line appended to DECISIONS; hits a non-obvious trap → /fw-trap)
You: Wrap up
Model: (/fw-record: writes the flight record, updates STATE, runs the
sensitive-info self-check)
Forgot to wrap up? When STATE lags ≥3 commits, the Stop hook blocks once before the session ends to remind you; --ci brings the same gate to other agents and human collaborators. For multi-session construction, say "handoff" before stopping so the model runs /fw-handoff .
The only thing you need to watch
Whether STATE's health is honest (green/yellow/red). The framework has a single quality metric: how long it takes a fresh session to reach a safe takeover after /fw-coldstart — if that takes more than 5 minutes, your records are degrading. Everything else — record count, format compliance — doesn't matter.
See what it actually looks like
This repo dogfoods its own framework: .flightwake/ contains the real STATE, DECISIONS, and records — every step from the gap list to the open-source launch is recorded there. That's exactly what will grow in your repo after installing.
New to working with a strong model? docs/workflow.md is a stage map of what you do and what to say to the model at each point — beginner main line, advanced folds for Claude Code veterans. (zh-TW for now; English lands with the v0.9 i18n pass.)
A Fable 5-class model doesn't need to be taught how to do the work — but there are four things no model can do however strong it gets, because they are structural and don't disappear as models improve:
Sessions die; context is finite. When work spans sessions, memory resets to zero; without records, every takeover is a git archaeology dig — a strong model just digs faster, it doesn't get to skip the dig.
Git records the what, not the why. Commits tell you what changed, never "why the other path wasn't taken" or "the root cause of that trap" — which happen to be the two most expensive pieces of information for the next session (or the next agent).
Discipline drifts in long sessions. "Reported done before the tests ran", "touched prod without leaving verification evidence" — these slides have nothing to do with model intelligence and need guards outside the model.
Agents don't share state. Claude, Codex, Gemini, and human teammates each see their own world; state only becomes everyone's once it's in git.
So flightwake supplements persistence and discipline, not intelligence . Its ancestor in spirit is GSD: GSD is navigation (turn-by-turn guidance for every step); flightwake is a dashcam + warning lights + road signs — strong models drive themselves, so the framework only does three things:
Dashcam : decisions, discoveries, verification evidence, recorded after the fact ( records/ , DECISIONS.md , TRAPS.md )
Warning lights : hard guards independent of model strength (tests green before "done", prod changes must leave verification evidence, destructive operations need confirmation first)
Road signs : any session can die; the next session reads STATE.md and takes over safely within 2 minutes
The origin was a real three-day session (2026-07-15~17: two repos, 19 commits, 4 cron jobs, 2 deep bug fixes — no upfront planning, zero derailment). It proved a strong model needs no navigation — but everything it left behind to make the next session possible (SUMMARY/CONTEXT/memory files) was improvised on the spot. flightwake turns that improvisation into an installable convention.
Core principle: records follow work — they don't lead it
GSD is stage-driven (research→plan→execute→verify gates); flightwake is trigger-driven (events create obligations):
Escalation rule (the opposite of GSD) : by default everything is quick — just start working; only "construction spanning multiple sessions" escalates to a phase (one CONTEXT file; plan decomposition is left to the model's in-the-moment judgment).
File layout (after installing into a target repo)
your-repo/
├── .flightwake/
│ ├── STATE.md # where we are now, next-step entry (always short, always current)
│ ├── DECISIONS.md # append-only decision log (one line per decision, with the why)
│ ├── TRAPS.md # trap registry (OKF-style frontmatter entries)
│ ├── TEMPLATE-record.md # flight-record template
│ ├── hooks/state-check.mjs # Stop hook: reminds you to wrap up when STATE lags ≥3 commits
│ └── records/ # flight records (one per meaningful wrap-up)
├── .claude/skills/fw-*/ # the four skills
└── .claude/settings.json # init merges the Stop hook config here
The skills and Stop hook are convenience sugar for Claude Code; .flightwake/ itself is plain Markdown — any agent that reads the instruction file can follow the same trigger obligations. Coexists with an existing GSD .planning/ (old records become historical archives).
--private keeps records local-only, out of git : every write is registered in .git/info/exclude (purely local — no trace left in the repo), the hook goes into .claude/settings.local.json , and the obligation table goes into CLAUDE.local.md (git-tracked instruction files are never touched). The cost: records aren't shared with the repo, and a fresh clone needs init --private again — "in git, shared with the repo" is flightwake's default and reason to exist; --private is the escape hatch for personal use inside someone else's repo.
uninstall reverses init's fixed write scope: removes the skills and framework files, extracts flightwake's Stop hook from settings (your other hooks stay untouched), and strips the marker blocks from instruction files and .git/info/exclude (files created by flightwake are deleted once emptied). .flightwake/ is user data and is kept by default ; only uninstall --purge deletes it too.
Monorepo policy: one install per repo, at the git root. Work is session-shaped — a session routinely spans multiple packages, and records follow the session; per-subdirectory installs would shred one stretch of work into fragmented records and turn "which STATE do I read?" into a new cold-start ambiguity. Running init in a subdirectory stops and points you to the root. Submodules have their own .git and count as independent repos. If a high-traffic multi-team monorepo sees false positives from the CI staleness check, tune --threshold first.
Wrap up your current milestone first, then:
npx flightwake init — coexists with .planning/ ; nothing is deleted
Tell your agent: "This repo is switching from GSD to flightwake. Read .planning/ for the current state and initialize .flightwake/STATE.md with /fw-record — unfinished items go into the next-step entries. From now on .planning/ is a historical archive; don't update it."
Remove (or comment out) GSD's own instruction block from CLAUDE.md, so the two rulebooks don't compete for the model's obedience
npx flightwake init --statusline puts a persistent gauge at the bottom of Claude Code:
✈️ flightwake │ ●green · STATE 2c behind │ ▓▓░░░░░░░░ 23%
Health color (the one thing you watch), STATE staleness (same rev-list logic as the Stop hook — but as a live gauge instead of an exit-time reminder), and context usage. The gauge also tells you the next command for the current state — session just started → → 開工先 /fw-coldstart ; STATE ≥3 commits behind → → /fw-record ; context running hot → → /fw-record → /clear → /fw-coldstart ; all healthy → silence. It never overwrites an existing statusline (a single-value setting), and repo-level config takes precedence over user-level, so it coexists with tools that set a global one.
Note: a plain npx flightwake init does not install the gauge — it's opt-in. Already ran init without it? Running npx flightwake init --statusline again just adds the gauge (everything else is skipped as already installed); the

[truncated]
