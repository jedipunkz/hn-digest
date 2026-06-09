---
source: "https://github.com/BMapAI/claude-ledger"
hn_url: "https://news.ycombinator.com/item?id=48464712"
title: "Claude Ledger – local cost/ROI dashboard for Claude Code usage"
article_title: "GitHub - BMapAI/claude-ledger: Zero-dependency cost & ROI dashboard for Claude Code usage — spend by project, session, model, and date range. · GitHub"
author: "itsopi"
captured_at: "2026-06-09T18:50:20Z"
capture_tool: "hn-digest"
hn_id: 48464712
score: 1
comments: 0
posted_at: "2026-06-09T17:48:42Z"
tags:
  - hacker-news
  - translated
---

# Claude Ledger – local cost/ROI dashboard for Claude Code usage

- HN: [48464712](https://news.ycombinator.com/item?id=48464712)
- Source: [github.com](https://github.com/BMapAI/claude-ledger)
- Score: 1
- Comments: 0
- Posted: 2026-06-09T17:48:42Z

## Translation

タイトル: Claude Ledger – クロード コードの使用に関するローカル コスト/ROI ダッシュボード
記事のタイトル: GitHub - BMapAI/claude-ledger: クロード コードの使用に関するゼロ依存コストと ROI ダッシュボード — プロジェクト、セッション、モデル、日付範囲ごとの支出。 · GitHub
説明: クロード コードの使用に関するゼロ依存コストと ROI ダッシュボード — プロジェクト、セッション、モデル、および日付範囲ごとの支出。 - BMapAI/クロードレジャー

記事本文:
GitHub - BMapAI/claude-ledger: クロード コードの使用に関するゼロ依存コストと ROI ダッシュボード — プロジェクト、セッション、モデル、および日付範囲ごとの支出。 · GitHub
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
BマップAI
/
クロード・レジャー
公共
通知
通知を変更するにはサインインする必要があります

n設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
35 コミット 35 コミット .github/ workflows .github/ workflows docs docs public public test test tools tools .gitignore .gitignore ライセンス ライセンス README.md README.md config.json config.json package.json package.json 価格設定.json 価格.json server.js server.js すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Claude Code プロジェクト向けの依存性ゼロのコスト / ROI アナライザー。
~/.claude/projects の下のセッショントランスクリプトを読み取り、トークンの使用状況を属性化します。
モデルごとの価格設定に加え、プロジェクトを選択できる小さな Web ダッシュボードを提供します
(または一度にすべて表示) して、実際の費用を確認してください。
クロード コード トークンをカウントするツールは数多くあります。帳簿はリターンと同じくらい重要です
支出 — プロジェクト / セッション / モデル / 日付ごとのコストと並んで、
生成された作業 (マークダウンでエクスポート可能なコミット、PR、ファイルのログ)、サーフェス
ある範囲（急上昇、摩擦、トレンド）の中で目立ったものは、トランスクリプトに残ります
クリーンアップ (永続的な履歴をオプトイン) し、Prometheus /metrics エンドポイントを公開します。
ローカル、読み取り専用、依存関係なし。
スクリーンショットでは合成デモ データを使用しています。
?project=…&session=… を介してディープリンクされた単一セッション。
「一見の価値がある」シグナル — 表示される各ビューの上部にランク付けされたパネル
選択した範囲内で実際に目立っているもの: 支出の急増、高価なセッション、
集中した支出、回収/摩擦コスト、前期と比較して大きな傾向、重い
オートメーション、認識されていないモデル、または低速回転のレイテンシーテール。純粋に説明的
(指示するものではありません)、ドルへの影響によってランク付けされ、次の場合にのみ表示されます。
何かが本当に際立っているため、ダッシュボードは、代わりにシグナルを表示します。
15 枚の均等な重みのパネル。
「あなたが作ったもの」作業ログ — the acc

ROI の補完側: 範囲ごと
コミット、PR、ファイル、作業したセッション、繰り返しのタイトルの概要
スタンドアップまたは
レビュー。 git/file アクティビティからの正直なアクティビティ ログです。価値判断ではありません。
プロジェクトごとの概要 — 総コスト、セッション、プロンプト、コスト/プロンプト、ツール呼び出し、
トークンの使用状況、キャッシュの節約、モデルごとの支出分割、日次支出グラフ、トップツール、
ソート可能なセッションごとのテーブル。
すべてのプロジェクトのロールアップ — 合計と支出ごとに並べ替え可能なプロジェクト
リーダーボード (行をクリックしてドリルインします)。
セッションのドリルダウン — 任意のセッションをクリックすると、そのプロンプト、ターンあたりのコストが表示されます。
タイムライン、ツールの使用状況、期間、合計。
共有可能な URL — 選択したプロジェクト、セッション、日付範囲は
アドレス バーなので、どのビューもブックマーク可能、リンク可能、前後ナビゲート可能です。
日付範囲フィルター — すべての時間 / 7 日 / 30 日 / 90 日 / 今月 / カスタム、適用済み
あらゆるビューにわたって。
ライブモニタリング — ライブ切り替えにより、現在のビューが 20 秒ごとに自動更新されます。
(ちらつきなし - テーブルの並べ替えとスクロール位置は保持されます)、
進行中のセッションの支出が増加。今日の支出と「何年前に更新された」を表示し、一時停止します
バックグラウンドタブで設定し、リロードしても設定を記憶します。
アクティビティ分析 — 平日×時間のパンチカード ヒートマップ (支出強度
時間、UTC）、日レベルのコールアウト（アクティブな日、アクティブな日ごとの平均、最も高価な日、
最も忙しい日）、およびモデルごとに積み上げられた 1 日の支出グラフなので、モデルごとの構成は次のようになります。
時間の経過とともに目に見えるもの。プロジェクトごとのビューとすべてのプロジェクトのビューについて。
予算とバーンレートの予測 — 毎月の予算とすべてのプロジェクト ビューを設定します。
月初から現在までの支出、現在のペースでの月末の合計予想額、および
予算オーバーの傾向にあるときにフラグを立てる予算バー。
前期比デルタ

— 日付範囲を選択すると、すべてのビューで日付範囲と比較されます。
その直前の等長ウィンドウ: 見出しカードの Δ% / Δ$ トレンド矢印と
リーダーボードのプロジェクトごとの「vs prev」列。
スキル/MCPサーバーによる支出 — 各ターンのコストをスキルまたはMCPに帰属させます
それを駆動したサーバー (Claude Code のメッセージごとの属性タグを使用) を使用して、次のことができます。
実際にどのスキルと統合に最もコストがかかるかを確認します。プロジェクトごとに表示され、
すべてのプロジェクトのビュー。帰属タグのないターンは単純に省略されます。
効率パネル — 支出の質の比率 (キャッシュ ヒット率、混合 $/100 万トークン、
プロンプトごとの出力シェア、トークンとツール）、コスト構成バー（出力 / フレッシュ）
入力 / キャッシュ読み取り / キャッシュ書き込み）、および同じトークンの「仮定の」再価格設定
ソネットとか俳句とか。
ツールの信頼性とリカバリ費用 - 全体的なエラー率、失敗した呼び出し数、および
各ツール結果の is_error フラグと「回復費用」からのツールごとの内訳:
失敗したツールの結果の後にアシスタントが右に曲がるコスト (
再試行/回復コスト)。 「エラー」には、本物のエラーとユーザーが拒否したエラーが含まれます。
中断された通話 (例: 拒否されたプラン) であるため、単なるバグではなく、摩擦として解釈してください。
正確なモデルの組み合わせ — モデルファミリーの分割を超えた、正確なモデルのバージョン
(例: claude-opus-4-8 vs claude-opus-4-7 ) とその支出の割合、つまりサイレント
モデルのロールオーバーは数字に現れます。
インタラクティブ vs 自動 — クロード コードのエントリポイント ( cli =
実践的なインタラクティブな作業 vs sdk-cli = ヘッドレス/SDK 主導の自動化)、
キーボードを操作するのではなく、自動化によってどれだけのコストが発生するかを確認できます。示されている
複数のエントリポイントが存在する場合。
時間とレイテンシ — 実時間とドルコスト、Claude Code による
turn_duration レコード: 合計アクティブ時間、ターン数、中央値、および

p90 ターン持続時間
(尾部が長いため中央値)、アクティブ時間あたりの時間、プロンプトあたりの時間、
ターン継続時間のヒストグラム。プロジェクトのリーダーボードには「アクティブ時間」列も表示されます。
出力/ROI — 台帳のリターン側: 支出が生み出したものをカウントします —
git commit と gh pr は呼び出しと、編集/書き込みによって操作される個別のファイルを作成します —
そして、コミットあたりのコストとファイルあたりのコストを導き出します。ツール呼び出しからの大まかなアクティビティのプロキシ
(価値や品質の尺度ではありません)、すべてのビューに表示されます。
プランの価値 — 定額料金 (Max/Pro) ユーザーの場合、金額は仮説です
請求書ではなく、API の価格の使用量をリストします。月額プラン料金とすべてのプロジェクト ビューを設定します
料金に対するレバレッジとして、月初から現在までの API 相当の使用量を示します (例: 「13 倍の
サブスクリプション」）に加えて、予測される月末のレバレッジが含まれます。
編集可能な価格設定 — 価格は 価格設定.json 内に存在し、変更時にホットリロードされます (いいえ
再起動します）。組み込みのデフォルトを使用するには、ファイルを削除します。
正確なコスト モデル — メッセージごとのキャッシュ作成 5 分/1 時間の内訳を使用します。
したがって、キャッシュ書き込みコストは推定ではなく正確です。
デフォルトでは読み取り専用です。 ~/.claude には決して書き込みません。トランスクリプトを読むだけです。
(オプションの永続履歴ロールアップは、 ~/.claude の外側にコンパクト ストアを書き込みます。
有効にした場合のみ — 以下を参照してください。)
依存関係なし — ノードの標準ライブラリのみ。
最速の方法 — クローンは必要ありません:
npx @bmapai/claude-ledger
# → http://127.0.0.1:4317
または、クローンから実行します (例: クローンをハッキングするため)。
ノードサーバー.js
環境変数による構成:
各個人のクロード記録は OS レベルで非公開です ( ~/.claude/** は
所有者のみ）。したがって、ツールは個人ごとに実行され、各人は自分の使用状況のみを確認します。
単一のインスタンスではチームメイトのデータを読み取ることはできません。
共有ボックスでの推奨設定:
読み取り可能な任意の場所でコードを一度チェックアウトします (e

.g. /var/www/クロードレジャー)
— アプリは独自のフォルダーに何も書き込まないため、読み取り専用で共有できます。
単一の git pull で更新されます。
各チームメイトは独自のインスタンスを実行します。インスタンスは ~/.claude を読み取ります。
ノード/var/www/claude-ledger/server.js
localhost にバインドし、空きポートを自動的に選択するため、同時ユーザーはバインドされません。
衝突すると、誰もプロンプト/コストがマシンの残りの部分に公開されなくなります。
ラップトップから SSH トンネル経由で表示します (起動ログには
正確なコマンド):
ssh -L 4317:localhost:4317 < あなた > @ < このホスト >
# 次に http://localhost:4317 を開きます
前提条件: 各ユーザーは、PATH (ユーザーごとの nvm) 上にノード 18 以降が必要です。
インストールまたはシステム全体のノード）。
HOST=0.0.0.0 を設定すると、プロンプト テキストと
支出 — そのポート上のホストにアクセスできるすべての人に。内蔵されていません
したがって、これは信頼できるネットワーク/制限のあるセキュリティ グループの背後でのみ実行してください。
[すべてのプロジェクト] ビューには、月初から今日までの支出を示す [今月] パネルと、
予測される月末の合計 (線形外挿: mtd / days_elapsed × days_in_month )。選択した月に関係なく、常に現在の暦月になります。
日付範囲。
毎月の予算を設定して、予算バー (使用量と予測、超過ペース) を把握します。
警告)。 config.json を編集します。
{ "毎月の予算" : 500 }
…または、ファイルを変更せずにインスタンスごとに設定します: MONTHLY_BUDGET=500 ノードserver.js
(環境変数が優先されます)。投影を表示するだけの場合は、未設定または null のままにしておきます。
プラン値（サブスクリプション ユーザー）
定額プラン (Max/Pro) をご利用の場合、上記の金額は API に相当します。
請求書ではなく、定価での使用です。月額料金をレバレッジとして再構成するように設定します。
サブスクリプション — 同じ作業にかかる従量制 API の費用:
{ "プラン月額料金" : 200 }
…または PLAN_MONTHLY_FEE=200 ノード server.js 。

[すべてのプロジェクト] ビューにプランが表示されます。
値パネル: 月初から現在までの API に相当する支出、定額料金、および回数
使用量に応じた料金を上回ります（月末の予測倍数を使用）。放っておいてください
unset/null でパネルを非表示にします。
料金は価格設定.json (100 万トークンあたりの米ドル) に保存されており、ホットリロードされます — 編集
ファイルと次のリクエストの価格が再設定されるため、再起動は必要ありません。落ちるファイルを削除する
server.js の組み込みのデフォルトに戻ります。
{
"作品" : { "入力" : 5 、 "出力" : 25 },
"ソネット" : { "入力" : 3 、 "出力" : 15 },
"俳句" : { "入力" : 1 , "出力" : 5 },
"cacheReadMultiplier" : 0.1 、
"cacheWrite5mMultiplier" : 1.25 、
"cacheWrite1hMultiplier" : 2.0
}
キャッシュは、cacheReadMultiplier × 入力レートで請求書を読み取ります。キャッシュ書き込みは 5 分で行われます
または 1 時間の乗数 (トランスクリプトのcache_creation からメッセージごとに選択)
故障、5mまで落ちます）。
これらの数値は概算であり、請求書ではありません。デフォルトは公開定価です
モデルや価格が変更されると変動する可能性があります。不明なモデル名は Opus にフォールバックします
料金。数値には、交渉による割引、一括価格設定、またはサブスクリプションは含まれません。
計画。請求書ではなく、相対比較と傾向の発見に使用します。
調整 — 実際の料金と一致するように価格設定.json を編集します。
Ledger はディスク上に残っているもののみを表示できます。クロード コードがローカル セッションを削除する
転写

[切り捨てられた]

## Original Extract

Zero-dependency cost & ROI dashboard for Claude Code usage — spend by project, session, model, and date range. - BMapAI/claude-ledger

GitHub - BMapAI/claude-ledger: Zero-dependency cost & ROI dashboard for Claude Code usage — spend by project, session, model, and date range. · GitHub
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
BMapAI
/
claude-ledger
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
35 Commits 35 Commits .github/ workflows .github/ workflows docs docs public public test test tools tools .gitignore .gitignore LICENSE LICENSE README.md README.md config.json config.json package.json package.json pricing.json pricing.json server.js server.js View all files Repository files navigation
A zero-dependency cost / ROI analyzer for Claude Code projects.
It reads the session transcripts under ~/.claude/projects , attributes token usage
to per-model pricing, and serves a small web dashboard where you can pick a project
(or view all at once) and see what it actually cost.
Plenty of tools count Claude Code tokens. Ledger is about the return as much as
the spend — alongside cost by project / session / model / date, it shows what the
work produced (a Markdown-exportable log of commits, PRs, and files), surfaces
what stands out in a range (spikes, friction, trends), survives transcript
cleanup (opt-in durable history), and exposes a Prometheus /metrics endpoint.
Local, read-only, and dependency-free.
Screenshot uses synthetic demo data.
A single session, deep-linked via ?project=…&session=… .
"Worth a look" signals — a ranked panel at the top of each view that surfaces
what actually stands out in the selected range: a spend spike, a pricey session,
concentrated spend, recovery/friction cost, a big trend vs the prior period, heavy
automation, an unrecognized model, or a slow-turn latency tail. Purely descriptive
(it points; it doesn't prescribe), ranked by dollar impact, and shown only when
something genuinely stands out — so the dashboard leads with signal instead of
fifteen equally-weighted panels.
"What you built" work log — the accomplishment side of ROI: a per-range
summary of commits, PRs, files, the sessions worked on, and recurring title
words, with a one-click Markdown export ( /api/worklog.md ) for a standup or
review. An honest activity log from git/file activity — not a value judgment.
Per-project overview — total cost, sessions, prompts, cost/prompt, tool calls,
token usage, cache savings, spend-by-model split, daily-spend chart, top tools, and a
sortable per-session table.
All-projects rollup — combined totals plus a sortable projects-by-spend
leaderboard (click a row to drill in).
Session drill-down — click any session to see its prompts, cost-per-turn
timeline, tool usage, duration, and totals.
Shareable URLs — the selected project, session, and date range live in the
address bar, so any view is bookmarkable, linkable, and back/forward-navigable.
Date-range filter — All time / 7d / 30d / 90d / this month / custom, applied
across every view.
Live monitoring — a Live toggle auto-refreshes the current view every 20s
(no flicker — table sort and scroll position are preserved) so you can watch an
in-progress session's spend climb. Shows today's spend and "updated Ns ago", pauses
in a background tab, and remembers the setting across reloads.
Activity analytics — a weekday × hour punchcard heatmap (spend intensity by
hour, in UTC), day-level callouts (active days, average per active day, priciest day,
busiest day), and a daily-spend chart stacked by model so the per-model mix is
visible over time. On the per-project and all-projects views.
Budget & burn-rate projection — set a monthly budget and the All-projects view
shows month-to-date spend, a projected month-end total at the current pace, and a
budget bar that flags when you're trending over.
Period-over-period deltas — pick a date range and every view compares it to the
equal-length window just before it: Δ% / Δ$ trend arrows on the headline cards and a
per-project "vs prev" column in the leaderboard.
Spend by skill / MCP server — attributes each turn's cost to the skill or MCP
server that drove it (using Claude Code's per-message attribution tags), so you can
see which skills and integrations actually cost the most. Shown on the per-project and
all-projects views; turns with no attribution tag are simply omitted.
Efficiency panel — quality-of-spend ratios (cache hit rate, blended $/1M tokens,
output share, tokens & tools per prompt), a cost-composition bar (output / fresh
input / cache read / cache write), and a "what-if" repricing of the same tokens on
Sonnet or Haiku.
Tool reliability & recovery spend — overall error rate, failed-call count, and a
per-tool breakdown from each tool result's is_error flag, plus "recovery spend":
the cost of the assistant turn right after a failed tool result (a proxy for
retry/recovery cost). "Errors" includes genuine failures and user-declined or
interrupted calls (e.g. a rejected plan), so read it as friction, not just bugs.
Exact model mix — beyond the model-family split, the precise model versions
(e.g. claude-opus-4-8 vs claude-opus-4-7 ) and their share of spend, so a silent
model rollover shows up in the numbers.
Interactive vs automated — splits spend by Claude Code entrypoint ( cli =
hands-on interactive work vs sdk-cli = headless / SDK-driven automation), so you
can see how much cost comes from automation rather than you at the keyboard. Shown
when more than one entrypoint is present.
Time & latency — wall-clock time alongside dollar-cost, from Claude Code's
turn_duration records: total active time, turn count, median and p90 turn duration
(median, since the tail is long), spend per hour of active time, time per prompt, and
a turn-duration histogram. The projects leaderboard also gets an "active time" column.
Output / ROI — the return side of the ledger: counts what the spend produced —
git commit and gh pr create calls and the distinct files touched by Edit/Write —
and derives cost-per-commit and cost-per-file. A rough activity proxy from tool calls
(not a measure of value or quality), shown on every view.
Plan value — for flat-fee (Max/Pro) users the dollar figure is hypothetical
API-list-price usage, not a bill. Set your monthly plan fee and the All-projects view
shows month-to-date API-equivalent usage as leverage over the fee (e.g. "13× your
subscription"), plus a projected month-end leverage.
Editable pricing — rates live in pricing.json , hot-reloaded on change (no
restart). Delete the file to use built-in defaults.
Accurate cost model — uses the per-message cache_creation 5m/1h breakdown when
present, so cache-write costs are exact rather than estimated.
Read-only by default — never writes to ~/.claude ; only reads the transcripts.
(The optional durable-history rollup writes a compact store outside ~/.claude ,
and only when you enable it — see below.)
No dependencies — Node standard library only.
The fastest way — no clone needed:
npx @bmapai/claude-ledger
# → http://127.0.0.1:4317
Or run from a clone (e.g. to hack on it):
node server.js
Configuration via environment variables:
Each person's Claude transcripts are private at the OS level ( ~/.claude/** is
owner-only). So the tool runs per person, each sees only their own usage — a
single instance can't read a teammate's data.
Recommended setup on a shared box:
Check out the code once anywhere readable (e.g. /var/www/claude-ledger )
— the app writes nothing to its own folder, so it can be shared read-only and
updated with a single git pull .
Each teammate runs their own instance — it reads their ~/.claude :
node /var/www/claude-ledger/server.js
It binds to localhost and auto-picks a free port, so simultaneous users don't
collide and nobody's prompts/costs are exposed to the rest of the machine.
View it over an SSH tunnel from your laptop (the start-up log prints the
exact command):
ssh -L 4317:localhost:4317 < you > @ < this-host >
# then open http://localhost:4317
Prerequisite: each user needs Node 18+ on their PATH (a per-user nvm
install or a system-wide Node).
Setting HOST=0.0.0.0 exposes the dashboard — including your prompt text and
spend — to anyone who can reach the host on that port. There is no built-in
auth, so only do this behind a trusted network / restrictive security group.
The All-projects view has a This month panel showing month-to-date spend and a
projected month-end total (linear extrapolation: mtd / days_elapsed × days_in_month ). It's always the current calendar month, regardless of the selected
date range.
Set a monthly budget to get a budget bar (used vs. projected, with an over-pace
warning). Either edit config.json :
{ "monthlyBudget" : 500 }
…or set it per instance without touching files: MONTHLY_BUDGET=500 node server.js
(the env var wins). Leave it unset/ null to just show the projection.
Plan value (subscription users)
If you're on a flat-fee plan (Max/Pro), the dollar figures above are API-equivalent
list-price usage, not your bill . Set your monthly fee to reframe that as leverage over
the subscription — how much metered API the same work would have cost:
{ "planMonthlyFee" : 200 }
…or PLAN_MONTHLY_FEE=200 node server.js . The All-projects view then shows a Plan
value panel: month-to-date API-equivalent spend, your flat fee, and how many times
over the fee your usage represents (with a projected month-end multiple). Leave it
unset/ null to hide the panel.
Rates live in pricing.json (USD per 1M tokens) and are hot-reloaded — edit
the file and the next request reprices, no restart needed. Delete the file to fall
back to the built-in defaults in server.js .
{
"opus" : { "input" : 5 , "output" : 25 },
"sonnet" : { "input" : 3 , "output" : 15 },
"haiku" : { "input" : 1 , "output" : 5 },
"cacheReadMultiplier" : 0.1 ,
"cacheWrite5mMultiplier" : 1.25 ,
"cacheWrite1hMultiplier" : 2.0
}
Cache reads bill at cacheReadMultiplier × the input rate; cache writes at the 5m
or 1h multiplier (chosen per message from the transcript's cache_creation
breakdown, falling back to 5m).
These figures are estimates, not your bill. Defaults are public list prices
and may drift as models/pricing change; unknown model names fall back to Opus
rates. Numbers exclude negotiated discounts, batch pricing, or subscription
plans. Use it for relative comparison and trend-spotting, not invoice
reconciliation — and edit pricing.json to match your actual rates.
Ledger can only show what's still on disk. Claude Code deletes local session
transcri

[truncated]
