---
source: "https://github.com/yonk-labs/claude-session-analyzer"
hn_url: "https://news.ycombinator.com/item?id=48671923"
title: "Show HN: Claude Code Session Trace/Browse Tool (Python)"
article_title: "GitHub - yonk-labs/claude-session-analyzer: See what Claude Code is really doing to your context — per-session & per-skill token/cost/time analysis and a Textual TUI over local transcripts. Find the prompt or skill quietly slowing you down. Python. · GitHub"
author: "TheMadHatter76"
captured_at: "2026-06-25T11:53:33Z"
capture_tool: "hn-digest"
hn_id: 48671923
score: 1
comments: 0
posted_at: "2026-06-25T11:33:07Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Claude Code Session Trace/Browse Tool (Python)

- HN: [48671923](https://news.ycombinator.com/item?id=48671923)
- Source: [github.com](https://github.com/yonk-labs/claude-session-analyzer)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T11:33:07Z

## Translation

タイトル: HN の表示: クロード コード セッション トレース/参照ツール (Python)
記事のタイトル: GitHub - yonk-labs/claude-session-analyzer: クロード コードが実際にコンテキストに対して行っていることを確認します。セッションごとおよびスキルごとのトークン/コスト/時間分析と、ローカル トランスクリプト上のテキスト TUI です。あなたの速度を静かに遅らせているプロンプトまたはスキルを見つけてください。パイソン。 · GitHub
説明: クロード コードがコンテキストに対して実際に何を行っているかを確認します。セッションごとおよびスキルごとのトークン/コスト/時間分析と、ローカル トランスクリプト上のテキスト TUI です。あなたの速度を静かに遅らせているプロンプトまたはスキルを見つけてください。パイソン。 - yonk-labs/claude-session-analyzer
HN テキスト: 私はさまざまなハーネス、エージェント メモリ ツールなどをテストしており、クロード セッションに従って、どのようなファイル、ツール、スキルがトークンを書き込んでいるかを分析したいと考えていました。ディスク上のクロード セッションを分析するためのこの小さなツールを構築しました。このアイデアは基本的に、claude からプロジェクトまたはセッションを選択し、そこにドリルダウンできるようにする tui です。トークンの消費、エラー、または速度低下を探しています。特定の時間/ターンに何が起こったかの核心を知る簡単な方法。便利だと思ったのでシェアしたいと思います。

記事本文:
GitHub - yonk-labs/claude-session-analyzer: クロード コードがコンテキストに対して実際に何を行っているかを確認します。セッションごとおよびスキルごとのトークン/コスト/時間分析と、ローカル トランスクリプト上のテキスト TUI です。あなたの速度を静かに遅らせているプロンプトまたはスキルを見つけてください。パイソン。 · GitHub
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
却下する

警告
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
ヨンクラボ
/
クロードセッションアナライザー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
yonk-labs/claude-session-analyzer
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
26 コミット 26 コミット .github/ workflows .github/ workflows csa csa docs docs testing testing .gitignore .gitignore ライセンス ライセンス README.md README.md profile.py profile.py pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Claude Code がコンテキストに対して実際に何をしているのか、そしてどのプロンプトまたは
スキルは静かにあなたを遅らせます。
Claude Code は、すべてのセッションの完全な JSONL トランスクリプトをすでに書き込んでいます。
~/.claude/projects/ 。 csa はそれらのトランスクリプトを読み取り、次のように変換します。
トークン、コスト、時間、スキルごとの動作 — ラッパーなし、インストルメンテーションなし、
追加のキャプチャ手順は必要ありません。データはすでにそこにあります。これは読むだけです。
巨大なグローバル CLAUDE.md 、数十のスキル、および山積みの MCP サーバーが搭載されています
毎ターンあなたのコンテキストを。 1 つのスキルが 1 ターンで 100,000 以上増加する可能性があります
トークン。 「役立つ」スキルは、求めてもいない質問であなたの邪魔をする可能性があります。
通常は何も見ることができません。今ならそれが可能です。
pipx install claude-session-analyzer # 推奨 (分離 CLI)
# または: UV ツールをインストールする claude-session-analyzer
# または: pip install claude-session-analyzer
CSA # コーパス プロファイル: 支出、膨張率、トップ セッション
csa --tui # インタラクティブブラウザ (メイン画面)
csa はコマンドライン アプリであるため、pipx / uv ツールはグローバルから遠ざけます。
site-packages — ただし、単純な pip install claude-session-analyzer も機能します。
git clone git@github.com:yonk-labs/claude-session-analyzer.git
cd クロード セッション アナライザー
uv venv .venv && uv pip install --pytho

n .venv/bin/python -e 。
# または: python3 -m pip install -e 。
テキスト CLI は stdlib のみです。 TUI の 1 つの依存関係は次のとおりです。
テキスト。
csa # すべてのセッションにわたるワンショット テキスト レポート
csa --tui # インタラクティブ: プロジェクト → セッション → ドリルダウン
csa --local # 現在のディレクトリのセッション (cwd のプロジェクト) のみ
csa --session FILE # 1 つのトランスクリプトのターンごとの内訳
csa /other/projects # 別のルートを指す
--local はテキスト レポートと --tui の両方で機能します。これは現在のレポートをマップします。
ディレクトリをその Claude Code プロジェクトに追加し、すべてをそこにスコープします。
================================================================
使用プロファイル (~/.claude/projects の下に 1,240 セッション)
================================================================
OUT（生成）：42,118,540トク
INフレッシュ（フル$）：14,002,221トク
IN キャッシュ読み取り : 10,540,883,002 トークン <- 常駐コンテキスト、再生
IN キャッシュ書き込み : 280,114,907 トーク
BLOAT (読み取り/新鮮) : 752.8x
東部標準時。支出: ~$8,431.55
支出額別トップ 15 セッション
$ out in+cache が壁トーク/s プロジェクトを回す
612.40 2,901,033 1,201,557,011 225 7740m 6.3 ~/acme-api
498.10 2,344,981 1,004,219,540 204 9060m 4.3 ~/web-app
...
BLOAT が見出しです。実際に入力するトークンごとに、最大 750 個のトークンが発生します。
スタンディング構成 (CLAUDE.md + スキル + MCP ツール定義) を毎ターンリプレイします。安いです
トークンごとに (キャッシュ読み取りの価格は最大 10% です)、ただし、それは巨大で一定のフットプリントになります。
プロジェクトの概要 (プロジェクトごとにロールアップされたセッション) が表示されます。ドリルイン
1 つのプロジェクトでそのセッションを表示するか、すべてのセッションごとに a を押します
プロジェクト。 ( csa --tui --local は現在のディレクトリに直接スキップします
) 以下のサンプル データは例示です。
1 · プロジェクト — プロジェクトごとにロールアップされたセッション (ランディング画面)
┌ CSA・プロジェクト ───────────

───────────────────┐
│ 52 プロジェクト · 1,240 セッション · ~$8,432 · プロジェクトを入力 · a=すべてのセッション │
━━━━━━━━━━━━━━━━━━━━┬─────┬───────┬─────┬─────┤
│ プロジェクト │ セッション │ $ ▼│ out │ in+cache │ 最後に使用した │
§───────────────┼─────┼─────┼───────┼─────┼─────┤
│ ~/acme-api │ 41 │ $2,788 │ 9.2M │ 3.1B │ 2026-06-21│
│ ~/web-app │ 33 │ $1,799 │ 6.1M │ 2.0B │ 2026-06-20│
│ ~/data-pipeline │ 10 │ $1,386 │ 2.4M │ 1.0B │ 2026-06-18│
━━━━━━━━━━━━━━━━━━━━━┴─────┴───────┴─────┴─────┘
プロジェクトを開く · a すべてのセッション · スキル · ツール · q 終了
2 · ブラウザ — プロジェクト (またはすべて) 内のセッション、並べ替え可能
┌ csa ─────────────────────────────┐
│ 1,240 セッション · ~$8,432 トークン価値 · スキル · ツール · m MCP · 1-9 ソート │
§─────┬──────┬─────┬───────┬──────┬──────┬─

─────┬─────────┤
│ $ ▼│ out │ in+cache │ ターン │ 壁 │ tok/s │ モデル │ プロジェクト │
§─────┼──────┼─────┼─────┼──────┼─────┼─────┼─────┤
│ $612.40 │ 2.9M │ 1.2B │ 225 │ 129h │ 6.3 │ opus-4-8 │ ~/acme-api │
│ $498.10 │ 2.4M │ 1.0B │ 204 │ 151h │ 4.3 │ opus-4-8 │ ~/web-app │
│ $372.30 │ 882k │ 564.7M │ 60 │ 43h │ 5.7 │ opus-4-8 │ ~/data-pipeline │
│ $187.30 │ 896k │ 285.3M │ 140 │ 26h │ 9.7 │ Sonnet-4 │ ~/web-app │
━─────┴─────┴─────┴─────┴─────┴─────┴─────┴─────┘
「開く」を入力 · スキル · t ツール · m MCP · q 終了
3 · セッション — コントロールパネル + ソート可能なターン
セッション統計のコントロール パネルが開きます - フリクション、スキル負荷、MCP コール、
質問の頻度など。 g を押して、タイムバケットと交換します。
棒グラフ。ターンテーブル (各ターンのプロンプトが表示されるようになりました) は常に下にあります。
┌ ~/acme-api ───────────────────────────┐
│ 3f9c0e1a · opus-4-8 · 225 ターン · 出力 290 万 · ピーク ctx 358,958 · $612.40 │
│ 6.3 tok/s · g=stats⇄graphs · a=全ターン · t=tools · m=MCP · コマンド入力│
━━━━━━━

─────────────────────────┤
│ 開始 2026-06-18 09:14 終了 2026-06-21 17:40 (壁時計7740m経過) │
│ │
│ ターン 225 ツールコール 1840 スキルロード 12 MCP コール 47 (plugin_playwright│
│ ×29、石碑×14、…) サブエージェント 9 があなたに尋ねました 6 │
│ │
│ フリクション 41/225 ターン · 修正 7 · ウォークバック 3 · 自己修正 12 · │
│ エラー 5 回 (ツール エラー 18 回) · 再試行 9 回 (証拠ではなく疑惑) │
│ │
│ 使用したスキル: ブレインストーミング×4、企画書作成×3、テスト駆動開発×2、… │
§────┬──────┬──────┬────────┬────┬──────┬──────┬──────┬──────┬───────┤
│ # │ ギャップ │ dur │ out │ ctx │ $ │ t/s │ ツール│ fric │ プロンプト │
§────┼──────┼──────┼────────┼───┼──────┼──────┼────┼──────┼───────┤
│ 1 │ 0s │ 726s │ 29.1k │ 92,008 │ $2.09 │ 40 │ 12 │ S・ │ …を構築する │
│ 2 │ 599s │ 95s │ 6.1k │ 96,894 │ $0.20 │ 65 │ 0 │ · │ また… │
│ 3 │ 117 │ 1217 │ 64.2k │ 384,334 │ $9.96 │ 53 │ 44 │ WEL │ 現在のプライ… │
━━━┴─────┴──────┴───────┴───┴──────┴─────┴────┴──────┴───────┘
g

グラフ · ターンに入る · t ツール · m MCP · a すべて · Esc 戻る · q 終了
gを押すと棒グラフが表示されます。各行は実際の時刻バケットです (それ以上のバケットはありません)。
"+120m" ミステリー) — バケットをクリックして、そのウィンドウの下のターンをフィルターします。
│ いつ │ トークン │ 使う │ ターン │
│ 06-18 09:14 │ ████████ 412k │ ██████ $88.10 │ ██████████ 41 │
│ 06-19 09:14 │ ██ 96k │ █ $14.20 │ ███ 12 │
│ 06-20 09:14 │ ██████████ 511k │ ███████ $140 │ █████ 22 ←スパイク │
fric flags (証拠ではなく疑惑): C = 次に修正しました ·
W = 次に別のアプローチに方向転換しました · S = 自動的に元に戻りました ·
E =2+ ツール エラー · L =同じコマンドを再試行しました。
4 · ターンの詳細 — コマンド
┌ ターン 3 コマンド ───────────────────────────┐
│ ターン 3 · ギャップ 117 秒 · デュル 1217 秒 · イン 1.9k / アウト 64.2k tok · ctx 384,334 · $9.96 │
│ スキル：クロード・アピ │
│ 時間 1217 秒 = 実行 320 秒 · あなた 0 秒 · モデル思考 897 秒 (あなた = AskUser の時間…) │
│ 摩擦 (証拠ではなく疑惑): 2 ツールエラー、ツールループ、ユーザーウォークバックネクスト│
│ ✗ 2 つの失敗した呼び出し: Bash、Bash (Enter を押してエラーを読み取ります) │
│ 次のユーザーがピボットしました: 「代わりに、別のツールを使用して価格を読み込んでください…」 │
│ exec = ツールの実行 · Wall = 呼び出し→次のステップ · Δ = モデルの思考 + アイドル後のアイドル │
│ │
│ プロンプト: Opus 4.x、Sonnet 4.x の 100 万トークンあたりの現在の価格… │
§────┬───────┬──────┬──────┬──────┬───────

──────────┤
│ # │ ツール │ 実行 │ 壁 │ Δ │ 概要 │
§────┼───────┼──────┼──────┼──────┼─────────────┤
│ 1 │ スキル │ 3 秒 │ 46 秒 │ 43 秒 │ クロード・アピ │
│ 2 │ Bash ✗ │ 8 秒 │ 40 秒 │ 32 秒 │ 時間 python3 profile.py --top 15 │
│ 3 │ ツールサーチ │ 0 秒 │ 29 秒 │ 29 秒 │ select:mcp__plugin_abe_abe__debate… │
│ 4 │ 書き込み │ 0 秒 │ 25 秒 │ 25 秒 │ csa/pricing.py │
━━━━━━━━━━━━━━━━━━━━┴─────┴─────┴───────────┘
ヘッダーには、ターン期間の 3 方向の分割が表示されます: ツールの実行量が表示されます。
(321 秒)、モデルは通話 (897 秒) の間にどれくらい考えていましたか、そしてあなたはどれくらい考えていましたか
AskUserQuestion プロンプトに答えます。失敗した通話 (✗) が表に表示されるようになりました —
以前はデータは追跡されていましたが、✗ マーカーは配線されませんでした。いずれかを入力してください
完全な入力とエラー結果のテキストを表示します。
Δ 列は静かな列です: python3 profile.py の実行時間は 8 秒でしたが、モデルは
次の行動の前に 32 秒考えました。インスタント ツール (書き込み、ツール検索)
大きな Δ は、これまで見たことのない純粋な思考時間です。
任意のコマンドで Enter を押すと、その完全なステップ (完全なツール入力) が開きます。
(Bash コマンド全体、書き込まれたファイル全体、サブエージェントへのプロンプト全体)
キャプチャされた結果に加えて:
┌ Bash — フルステップ ───────────────────

─────────┐
│ バッシュ・実行 8 秒・壁 40 秒・Δ 32 秒 │
│ │
│ 入力 │
│ ─────────────────────── │
│ { │
│ "コマンド": "time python3 profile.py --top 15", │
│ "description": "すべてのトランスクリプトに対して完全な使用プロファイルを実行します" │
│ } │
│ │
│ 結果（上限あり） │
│ ─────────────────────── │
│ 使用プロファイル (~/.claude/projects に 1,240 セッション) … │
━━━━━━━━━━━━━━━━━━━━━━━━━━━┘
5 · スキルの後悔 — どのスキルがあなたのスピードを遅らせているのか ( s )
尋ねる = スキルが入力を中断する頻度。後悔% = その割合
摩擦で回転します (証明ではなく相関 — アウト/質問/ツールが重要です)
信頼できるコラム）。
5 回未満しか発動しなかったスキルは、薄暗いテキストで n<5 と表示され、
後悔率で並べ替えると最下位 — 1 件の火災による 100% はデータではなくノイズです。
┌ スキルの後悔 — 証拠ではなく疑惑 ───────────────────┐
│ ターン = 実行した回転数 · tools = ツールが i を呼び出す

[切り捨てられた]

## Original Extract

See what Claude Code is really doing to your context — per-session & per-skill token/cost/time analysis and a Textual TUI over local transcripts. Find the prompt or skill quietly slowing you down. Python. - yonk-labs/claude-session-analyzer

I have been testing different harnesses, agent memory tools, etc. and wanted to follow along with my claude sessions and analyze what files, tools, skills were burning tokens. Built this small tool to analyze claude sessions on disk. The idea is basically a tui that allows you to select a project or session from claude, and then drill into it. Looking for token consumption, errors, or slowdowns. An easy way to get down to the nitty gritty of what happened at a specific time/turn. I found it useful, so I thought I would share.

GitHub - yonk-labs/claude-session-analyzer: See what Claude Code is really doing to your context — per-session & per-skill token/cost/time analysis and a Textual TUI over local transcripts. Find the prompt or skill quietly slowing you down. Python. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
yonk-labs
/
claude-session-analyzer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
yonk-labs/claude-session-analyzer
main Branches Tags Go to file Code Open more actions menu Folders and files
26 Commits 26 Commits .github/ workflows .github/ workflows csa csa docs docs tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md profile.py profile.py pyproject.toml pyproject.toml View all files Repository files navigation
See what Claude Code is really doing to your context — and which prompt or
skill is quietly slowing you down.
Claude Code already writes a full JSONL transcript of every session to
~/.claude/projects/ . csa reads those transcripts and turns them into
tokens, cost, time, and per-skill behavior — no wrapper, no instrumentation,
no extra capture step. The data is already there; this just reads it.
A big global CLAUDE.md , dozens of skills, and a pile of MCP servers ride in
your context on every turn. One skill can balloon a single turn by 100k+
tokens. A "helpful" skill can interrupt you with questions you never asked for.
You usually can't see any of it. Now you can.
pipx install claude-session-analyzer # recommended (isolated CLI)
# or: uv tool install claude-session-analyzer
# or: pip install claude-session-analyzer
csa # corpus profile: spend, bloat ratio, top sessions
csa --tui # the interactive browser (main surface)
csa is a command-line app, so pipx / uv tool keep it out of your global
site-packages — but plain pip install claude-session-analyzer works too.
git clone git@github.com:yonk-labs/claude-session-analyzer.git
cd claude-session-analyzer
uv venv .venv && uv pip install --python .venv/bin/python -e .
# or: python3 -m pip install -e .
The text CLI is stdlib-only. The TUI's one dependency is
Textual .
csa # one-shot text report over every session
csa --tui # interactive: projects → sessions → drill down
csa --local # only the current directory's sessions (the cwd's project)
csa --session FILE # per-turn breakdown of one transcript
csa /other/projects # point at a different root
--local works for both the text report and --tui — it maps the current
directory to its Claude Code project and scopes everything to it.
======================================================================
USAGE PROFILE (1,240 sessions under ~/.claude/projects)
======================================================================
OUT (generated) : 42,118,540 tok
IN fresh (full $) : 14,002,221 tok
IN cache-read : 10,540,883,002 tok <- standing context, replayed
IN cache-write : 280,114,907 tok
BLOAT (read/fresh) : 752.8x
EST. SPEND : ~$8,431.55
TOP 15 SESSIONS BY SPEND
$ out in+cache turns wall tok/s project
612.40 2,901,033 1,201,557,011 225 7740m 6.3 ~/acme-api
498.10 2,344,981 1,004,219,540 204 9060m 4.3 ~/web-app
...
BLOAT is the headline: for every token you actually type, ~750 tokens of
standing config (CLAUDE.md + skills + MCP tool defs) replay each turn. It's cheap
per token (cache-read is ~10% price) but it's a huge, constant footprint.
It opens on a Projects overview — sessions rolled up per project. Drill into
one project to see its sessions, or press a for every session across all
projects. ( csa --tui --local skips straight to the current directory's
sessions.) Sample data below is illustrative.
1 · Projects — sessions rolled up per project (landing screen)
┌ csa · projects ─────────────────────────────────────────────────────────────┐
│ 52 projects · 1,240 sessions · ~$8,432 · Enter a project · a=all sessions │
├──────────────────────────┬──────────┬──────────┬───────┬──────────┬───────────┤
│ project │ sessions │ $ ▼│ out │ in+cache │ last used │
├──────────────────────────┼──────────┼──────────┼───────┼──────────┼───────────┤
│ ~/acme-api │ 41 │ $2,788 │ 9.2M │ 3.1B │ 2026-06-21│
│ ~/web-app │ 33 │ $1,799 │ 6.1M │ 2.0B │ 2026-06-20│
│ ~/data-pipeline │ 10 │ $1,386 │ 2.4M │ 1.0B │ 2026-06-18│
└──────────────────────────┴──────────┴──────────┴───────┴──────────┴───────────┘
Enter open project · a all sessions · s skills · t tools · q quit
2 · Browser — sessions in a project (or all), sortable
┌ csa ───────────────────────────────────────────────────────────────┐
│ 1,240 sessions · ~$8,432 token-value · s skills · t tools · m MCP · 1-9 sort │
├─────────┬──────┬──────────┬───────┬──────┬───────┬──────────┬─────────────────┤
│ $ ▼│ out │ in+cache │ turns │ wall │ tok/s │ model │ project │
├─────────┼──────┼──────────┼───────┼──────┼───────┼──────────┼─────────────────┤
│ $612.40 │ 2.9M │ 1.2B │ 225 │ 129h │ 6.3 │ opus-4-8 │ ~/acme-api │
│ $498.10 │ 2.4M │ 1.0B │ 204 │ 151h │ 4.3 │ opus-4-8 │ ~/web-app │
│ $372.30 │ 882k │ 564.7M │ 60 │ 43h │ 5.7 │ opus-4-8 │ ~/data-pipeline │
│ $187.30 │ 896k │ 285.3M │ 140 │ 26h │ 9.7 │ sonnet-4 │ ~/web-app │
└─────────┴──────┴──────────┴───────┴──────┴───────┴──────────┴─────────────────┘
Enter open · s skills · t tools · m MCP · q quit
3 · Session — control panel + sortable turns
Opens on a control panel of session stats — friction, skill loads, MCP calls,
how often it asked you, and more. Press g to swap it for the time-bucketed
bar graphs. The turns table (now with the prompt of each turn) is always below.
┌ ~/acme-api ─────────────────────────────────────────────────────────────────┐
│ 3f9c0e1a · opus-4-8 · 225 turns · out 2.9M · peak-ctx 358,958 · $612.40 │
│ 6.3 tok/s · g=stats⇄graphs · a=all turns · t=tools · m=MCP · Enter for commands│
├───────────────────────────────────────────────────────────────────────────────┤
│ started 2026-06-18 09:14 ended 2026-06-21 17:40 (7740m elapsed wall-clock) │
│ │
│ turns 225 tool calls 1840 skill loads 12 MCP calls 47 (plugin_playwright│
│ ×29, stele×14, …) subagents 9 asked you 6 │
│ │
│ friction 41/225 turns · corrections 7 · walkbacks 3 · self-corrections 12 · │
│ error-turns 5 (18 tool errors) · retry-loops 9 (suspicion, not proof) │
│ │
│ skills used: brainstorming×4, writing-plans×3, test-driven-development×2, … │
├────┬──────┬──────┬───────┬─────────┬───────┬─────┬──────┬──────┬───────────────┤
│ # │ gap │ dur │ out │ ctx │ $ │ t/s │ tools│ fric │ prompt │
├────┼──────┼──────┼───────┼─────────┼───────┼─────┼──────┼──────┼───────────────┤
│ 1 │ 0s │ 726s │ 29.1k │ 92,008 │ $2.09 │ 40 │ 12 │ S· │ Build the … │
│ 2 │ 599s │ 95s │ 6.1k │ 96,894 │ $0.20 │ 65 │ 0 │ · │ Also add … │
│ 3 │ 117s │1217s │ 64.2k │ 384,334 │ $9.96 │ 53 │ 44 │ WEL │ current pri… │
└────┴──────┴──────┴───────┴─────────┴───────┴─────┴──────┴──────┴───────────────┘
g graphs · Enter a turn · t tools · m MCP · a all · Esc back · q quit
Press g for the bar graphs. Each row is a real clock-time bucket (no more
"+120m" mystery) — click a bucket to filter the turns below to that window:
│ when │ tokens │ spend │ turns │
│ 06-18 09:14 │ ████████ 412k │ ██████ $88.10 │ ██████████ 41 │
│ 06-19 09:14 │ ██ 96k │ █ $14.20 │ ███ 12 │
│ 06-20 09:14 │ ██████████ 511k │ ███████ $140 │ █████ 22 ←spike │
fric flags (suspicion, not proof): C =you corrected it next ·
W =you pivoted to a different approach next · S =it walked itself back ·
E =2+ tool errors · L =retried the same command.
4 · Turn detail — the commands
┌ turn 3 commands ────────────────────────────────────────────────────────────┐
│ Turn 3 · gap 117s · dur 1217s · in 1.9k / out 64.2k tok · ctx 384,334 · $9.96 │
│ skills: claude-api │
│ time 1217s = exec 320s · you 0s · model-think 897s (you = time on AskUser…) │
│ friction (suspicion, not proof): 2 tool-error(s), tool-loop, user-walkback-next│
│ ✗ 2 failing call(s): Bash, Bash (Enter to read the error) │
│ next user pivoted: "instead, use a different tool to load prices…" │
│ exec = tool run · wall = call→next step · Δ = model think + idle after │
│ │
│ prompt: current pricing per million tokens for Opus 4.x, Sonnet 4.x… │
├────┬─────────────┬──────┬──────┬──────┬─────────────────────────────────────┤
│ # │ tool │ exec │ wall │ Δ │ summary │
├────┼─────────────┼──────┼──────┼──────┼─────────────────────────────────────┤
│ 1 │ Skill │ 3s │ 46s │ 43s │ claude-api │
│ 2 │ Bash ✗ │ 8s │ 40s │ 32s │ time python3 profile.py --top 15 │
│ 3 │ ToolSearch │ 0s │ 29s │ 29s │ select:mcp__plugin_abe_abe__debate… │
│ 4 │ Write │ 0s │ 25s │ 25s │ csa/pricing.py │
└────┴─────────────┴──────┴──────┴──────┴─────────────────────────────────────┘
The header shows the 3-way split of turn duration: how much was tool execution
(321s), how much was the model thinking between calls (897s), and how much was you
answering AskUserQuestion prompts. Failing calls (✗) now show in the table —
previously the data was tracked but the ✗ marker was never wired. Enter any of them
to see the full input + the error result text.
The Δ column is the quiet one: time python3 profile.py ran 8s, but the model
then thought 32s before its next move. Instant tools (Write, ToolSearch) with a
big Δ are pure think time you'd never have seen.
Press Enter on any command to open its full step — the complete tool input
(the whole Bash command, the full file written, the entire prompt to a subagent)
plus the captured result:
┌ Bash — full step ───────────────────────────────────────────────────────────┐
│ Bash · exec 8s · wall 40s · Δ 32s │
│ │
│ INPUT │
│ ──────────────────────────────────────────────────────────── │
│ { │
│ "command": "time python3 profile.py --top 15", │
│ "description": "Run full usage profile across all transcripts" │
│ } │
│ │
│ RESULT (capped) │
│ ──────────────────────────────────────────────────────────── │
│ USAGE PROFILE (1,240 sessions under ~/.claude/projects) … │
└───────────────────────────────────────────────────────────────────────────────┘
5 · Skill regret — which skill is slowing you down ( s )
asks = how often a skill interrupts you for input. regret% = share of its
turns with friction (correlation, not proof — out / asks / tools are the
trustworthy columns).
Skills that fired fewer than 5 times show n<5 in dim text and sink to the
bottom when sorted by regret% — a 100% from one fire is noise, not data.
┌ skill regret — suspicion, not proof ────────────────────────────────────────┐
│ turns=turns it ran · tools=tool calls i

[truncated]
