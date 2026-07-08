---
source: "https://github.com/hoophq/rs"
hn_url: "https://news.ycombinator.com/item?id=48834275"
title: "Show HN: Free and open source risk summary for AI sessions made in 5 days"
article_title: "GitHub - hoophq/rs: AI session risk summary. See what sensitive information is sitting in your local AI coding sessions. Runs locally, nothing leaves your machine. · GitHub"
author: "luisfdias"
captured_at: "2026-07-08T17:20:07Z"
capture_tool: "hn-digest"
hn_id: 48834275
score: 1
comments: 1
posted_at: "2026-07-08T16:51:11Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Free and open source risk summary for AI sessions made in 5 days

- HN: [48834275](https://news.ycombinator.com/item?id=48834275)
- Source: [github.com](https://github.com/hoophq/rs)
- Score: 1
- Comments: 1
- Posted: 2026-07-08T16:51:11Z

## Translation

タイトル: HN を表示: 5 日間で行われた AI セッションの無料およびオープンソースのリスク概要
記事のタイトル: GitHub - hoophq/rs: AI セッションのリスクの概要。ローカルの AI コーディング セッションにどのような機密情報が存在するかを確認します。ローカルで実行され、マシンからは何も離れません。 · GitHub
説明: AI セッションのリスクの概要。ローカルの AI コーディング セッションにどのような機密情報が存在するかを確認します。ローカルで実行され、マシンからは何も離れません。 - フーフク/rs

記事本文:
GitHub - hoophq/rs: AI セッションのリスクの概要。ローカルの AI コーディング セッションにどのような機密情報が存在するかを確認します。ローカルで実行され、マシンからは何も離れません。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。

フーフク
/
rs
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
38 コミット 38 コミット .github/ ワークフロー .github/ ワークフロー 分析 分析 cmd/ hooprs cmd/ hooprs docs/ アセット docs/ アセット 例 例 ガードレール ガードレール homebrew homebrew npm npm レポート レポート リスク リスク ソース ソース 状態 状態タイプ タイプ .gitignore .gitignore ライセンス ライセンス README.md README.md go.mod go.mod go.sum go.sum install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング エージェントは何を見たのでしょうか?
hooprs はローカル AI コーディング セッションをスキャンします — クロード コード、カーソル、OpenCode —
PII とシークレットについては、完全にマシン上で正確に表示されます
どのセッションで何がリークされたのか。
· ゲートウェイなし · ネットワークなし · API 呼び出しなし
brew install hoophq/tap/hooprs && hooprs
コマンドは 1 つです。ディスク上のすべてのセッションを検出し、プロセス内で検出を実行します。
リスクの概要を端末に出力し、自己完結型の HTML レポートを開きます。
最も露出の多いセッションをランキングします。
入力したすべてのプロンプトとエージェントが読み取るすべてのファイルがセッション履歴になります。
ディスク上には、資格情報、顧客データ、国民 ID なども含まれます。
どれくらいかは見てみないと分かりません。フーパーの見た目:
🔒ローカルのみ。検出はマシン上でインプロセスで実行されます。 DLPなし
サービス、API 呼び出しなし、ディスクから何も残らない - 読み取りスキャナの場合
あなたの秘密、これが重要です。
✅ 形状が一致しているだけでなく、検証済み。クレジットカードはLuhnチェックを受けており、
IBAN は Mod-97 チェック済み、SSN は範囲検証済み。失敗する 16 桁の数字
チェックサムがレポートに届くことはありません。
🎯 ダンプではなく、ランク付けされます。セッションは段階的にクリティカル/マイナー/ローに分類されます
露出によってランク付けされるため、重要な 10 セッションを優先順位付けできます。

代わりに言う
1 万件の検索結果をスクロールします。
🤫 価値のないレポート。 HTML レポートと JSON レポートにはエンティティ タイプが含まれます。
カウントと重大度 - 値が一致することはありません。レポートを決して共有しない
再び漏れが発生します。
🚦 方向を意識します。検出結果は入力 (入力した内容) と
出力 (エージェントがコンテキストに取り込んだもの) — エージェントの読み取りをリークします。
貼り付けたものよりも重くなります。
macOS と BSD には rs が付属しているため、このコマンドは rs ではなく hooprs です。
ストック rs(1) ユーティリティ (データ配列の再形成)。
影がついた。
# Homebrew — macOS と Linux
醸造インストール hoophq/tap/hooprs
# シェルスクリプト — macOS と Linux
カール -fsSL https://raw.githubusercontent.com/hoophq/rs/main/install.sh |しー
# npm — Windows (x64) もカバーします
ねぎ @hoophq/rs
3 つとも、事前に構築されたバイナリをインストールします。コンパイル手順は必要ありません。シェルスクリプト
チェックサムを検証し、/usr/local/bin (または ~/.local/bin) にインストールします。
それは書き込み可能ではありません); HOOPRS_VERSION=v0.2.0 でバージョンを固定するか、
HOOPRS_INSTALL_DIR=~/bin の宛先。 Homebrew の公式は次の場所にあります。
hoophq/homebrew-tap ; npm は、
オプションの依存関係 ( @hoophq/rs-<os>-<arch> ) によるバイナリ。
ソースからのビルドには Go 1.24 以降が必要で、単一の純粋な Go 依存関係 (
アルカトラズ検出ライブラリ):
go build -o hooprs ./cmd/hooprs
クイックスタート
すべてをスキャンし、概要を印刷し、risk-report.html を作成して開きます。
# 過去 30 日間のみと、機械可読な JSON レポート
hooprs -days 30 -json リスクレポート.json
# プロジェクトがパターンに一致するカーソル セッションのみ
hooprs -tools カーソル -プロジェクト ' my-app '
# 信頼バーを上げる (デフォルトは 0.4)
hooprs -min-score 0.6
# ターミナルに実際にリークされた値を表示します (レポートには決して表示されません)
hooprs -show-values
すべてのフラグ
旗
デフォルト
説明
-アウト
リスクレポート.html
t のパス

自己完結型の HTML レポート
-json
(オフ)
機械可読なリスクレポートもここに書きます
-ツール
クロード、カーソル、オープンコード
スキャンするソース
-プロジェクト
(全員)
セッションプロジェクトの正規表現フィルター
-セッション
(全員)
セッション ID の正規表現フィルター
-日
0 (常に)
過去 N 日以内に開始されたセッションのみ
-ホーム
$ホーム
セッションを検出するためのホーム ディレクトリ
-ルール
(なし)
ガードレール ルールの JSON ファイル
-最小スコア
0.4
カウントするための最小検出信頼度 (0 ～ 1)
-クリティカル-ウェイト
60
クリティカル セッション シェアのセキュリティ スコア ペナルティの重み (0 ～ 100)
-エンジン
アルカトラズ島
検出エンジン: alcatraz (完全な PII セット) またはスタブ (ゼロ依存性フォールバック)
-増分
偽
前回の実行以降に追加されたコンテンツのみをスキャンします
-状態
~/.risk-analyzer/state.json
インクリメンタルスキャン状態ファイル
-静か
偽
端末の概要を抑制する
-値の表示
偽
上位 10 個のクリティカル セッションに一致する重大度の高い値を端末概要に出力します (HTML/JSON レポートには書き込まれません)。
-開く
本当の
完了したら、デフォルトのブラウザで HTML レポートを開きます
-バージョン
偽
hoops バージョンを出力して終了します
デフォルトでは、すべての実行はすべてのセッションの完全なスナップショットになります。 -増分
ファイルごとのバイト オフセットが保持されるため、後続の実行ではコンテンツの読み取りのみが行われます。
前回の実行以降に追加されました (「前回からの変更点」に役立ちます)。
レポートのタイトル バーには 2 つのボタンがあるため、スキャンするとユーザーへのメッセージになります。
ワンクリックでチームを編成:
共有は 1200×630 の概要カード (スコア、リスク ドーナツ、見出し) をラスタライズします
counts) を PNG に変換し、クリップボードにコピーすると、貼り付ける準備ができていることが表示されます。
Slack メッセージ — 画像を貼り付け、テキストを送信します。ブラウザがブロックしている場合
クリップボードにアクセスすると、同じパネルで PNG ダウンロードが提供されます。
PDF を保存すると、印刷ダイアログが開き、印刷スタイルシートが保持されます。
ブランドのダークな外観とインタラクティブなクロムの削除 — pi

「PDFとして保存」をクリックします。
どちらもページ内で完全に実行されます。カードは同じ価値のないデータから構築されています。
レポートとして (エンティティのタイプと数、一致する値はありません)、
Slack メッセージには、集計カウントとインストールに関するワンライナーのみが含まれます。共有
スキャンによってリークが再漏洩することはありません。
構造化 PII (alcatraz エンジン経由)
加えて、コーディング セッションで一般的なシークレット タイプ (hooprs 独自のシークレット経由)
パック):
検出では、正規表現パターンとチェックサムおよび形式バリデータを組み合わせます。
(Luhn、IBAN mod-97、SSN/国民 ID 範囲ルール)、および以下の一致
-min-score しきい値は削除されます。ゼロ依存関係のエンジン スタブを渡します
正規表現のフォールバック。
NLP モデルが必要な NER: PERSON / LOCATION スタイルのエンティティに関する注意事項
このバージョンには近づかないでください。分析装置は小さな装置の後ろにあります
analyzer.Analyzer インターフェイスを使用すると、将来の NLP ベースのエンジンが何も必要なく組み込まれます。
パイプラインに触れます。
ソース→分析→リスク→レポート
クロード/正規表現 + 層 · ターミナル +
カーソル/バリデータースコア · html + json
オープンコード (ローカル) 公開
ソースは各ツールのディスク上のセッション形式を検出し、1 つの形式に解析します。
正規化されたモデル。分析はすべてのメッセージに対して検出エンジンを実行します。
リスクは生の発見を階層、露出ランキング、スコアに変換します。
レポートは、端末の概要と自己完結型の HTML をレンダリングします。
セッションごとのレベル: クリティカル (重大度の高いエンティティ)、マイナー
(重大度が中程度のみ)、または低。
エクスポージャは、重大度に重み付けされた検出数に基づいてセッションをランク付けします。
入力よりも出力 (エージェント コンテキストに取り込まれたデータ) が優先されます。
セキュリティ スコア = クランプ(0, 100, ラウンド(100 − W・クリティカル/合計 − 20・マイナー/合計)) 、
ここで、 W はクリティカル ペナルティの重み ( -critical-weight 、デフォルトは 60) です。
エンティティ タイプごとの重大度とデータ ファミリが存在します
リスク/エンティティ.go 。
-rules <file> を使用して独自の検出ルールを階層化します — 方向を認識します
(

入力 = 入力した内容、出力 = アシスタント/ツールの出力)。参照
例/guardrails.json :
{
「ルール」: [
{ "name" : " 内部ホスト名 " 、 "type" : " regex " 、 "direction" : " Both " 、
"パターン" : " \\ b[a-z0-9-]+ \\ .internal \\ .example \\ .com \\ b " },
{ "name" : "private-key-material" 、"type" : "deny_words" 、"direction" : "output" 、
"単語" : [ " RSA 秘密キーの開始 " ] }
]
}
違反は概要とレポートに独自のセクションが表示され、違反ごとにカウントされます。
ルール。
すべてがマシン上で実行されます。 HTML/JSON レポートにはエンティティのみが含まれます
タイプ、数、重大度、セッション識別子 - 値が一致することはありません。
マシンからは何も残りません。
-show-values は 1 つの意図的な例外です。一致した値を出力します。
端末に対する上位 10 個のクリティカル セッションのみの高重大度値、
実際の漏れ箇所を特定できるようになります。 HTML レポートと JSON レポートには価値がありません。
フラグがオンであっても、それを保証するテストがあります。
cmd/hooprs/ CLI: フラグ → 検出 → 分析 → リスク → レンダリング
ソース/検出と解析クロード/カーソル/オープンコード セッション
状態/増分スキャンのオフセット
タイプ/正規化されたセッション/メッセージ モデル
分析/アナライザー インターフェイス + alcatraz エンジン、共有シークレット パック、スタブ フォールバック
ガードレール/ローカル ルール ローダー + 方向認識マッチャー
リスク/重大度カタログ + リスク モデル (階層、エクスポージャ、スコア)
レポート/ターミナル + 内蔵 HTML レンダラ (埋め込み CSS/JS)
MIT © hoop.dev — hoop の背後にあるチームによって構築されました。
AI セッションのリスクの概要。ローカルの AI コーディング セッションにどのような機密情報が存在するかを確認します。ローカルで実行され、マシンからは何も離れません。
hoop.dev/labs/ai-risk-summary
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
1
フォーク
レポートリポジトリ
リリース
12
v0.2.6
最新の
7月3日、2日

026
+ 11 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

AI session risk summary. See what sensitive information is sitting in your local AI coding sessions. Runs locally, nothing leaves your machine. - hoophq/rs

GitHub - hoophq/rs: AI session risk summary. See what sensitive information is sitting in your local AI coding sessions. Runs locally, nothing leaves your machine. · GitHub
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
hoophq
/
rs
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
38 Commits 38 Commits .github/ workflows .github/ workflows analyze analyze cmd/ hooprs cmd/ hooprs docs/ assets docs/ assets examples examples guardrails guardrails homebrew homebrew npm npm report report risk risk sources sources state state types types .gitignore .gitignore LICENSE LICENSE README.md README.md go.mod go.mod go.sum go.sum install.sh install.sh View all files Repository files navigation
What did your AI coding agent see?
hooprs scans your local AI coding sessions — Claude Code, Cursor, OpenCode —
for PII and secrets, entirely on your machine , and shows you exactly
which sessions leaked what.
· No gateway · No network · No API calls
brew install hoophq/tap/hooprs && hooprs
One command. It discovers every session on disk, runs detection in-process,
prints a risk summary to the terminal, and opens a self-contained HTML report
ranking your most exposed sessions.
Every prompt you type and every file your agent reads becomes session history
on your disk — and some of it is credentials, customer data, national IDs.
You have no idea how much until you look. hooprs looks:
🔒 Local only. Detection runs in-process on your machine. No DLP
service, no API calls, nothing leaves your disk — for a scanner that reads
your secrets, this is the whole point.
✅ Verified, not just shape-matched. Credit cards are Luhn-checked,
IBANs mod-97-checked, SSNs range-validated. A 16-digit number that fails
the checksum never reaches your report.
🎯 Ranked, not dumped. Sessions are tiered critical / minor / low
and ranked by exposure, so you triage the ten sessions that matter instead
of scrolling ten thousand findings.
🤫 Value-free reports. The HTML and JSON reports carry entity types,
counts, and severities — never the matched values. Sharing the report never
re-leaks the leak.
🚦 Direction-aware. Findings split into input (what you typed) and
output (what the agent pulled into its context) — leaks the agent read
weigh heavier than ones you pasted.
The command is hooprs rather than rs because macOS and the BSDs ship a
stock rs(1) utility (reshape a data array) that would otherwise be
shadowed.
# Homebrew — macOS & Linux
brew install hoophq/tap/hooprs
# Shell script — macOS & Linux
curl -fsSL https://raw.githubusercontent.com/hoophq/rs/main/install.sh | sh
# npm — also covers Windows (x64)
npx @hoophq/rs
All three install a prebuilt binary — no compile step. The shell script
verifies the checksum and installs to /usr/local/bin (or ~/.local/bin when
that's not writable); pin a version with HOOPRS_VERSION=v0.2.0 or change the
destination with HOOPRS_INSTALL_DIR=~/bin . The Homebrew formula lives in
hoophq/homebrew-tap ; npm pulls the
binary through optional dependencies ( @hoophq/rs-<os>-<arch> ).
Building from source needs Go 1.24+ and has a single pure-Go dependency (the
alcatraz detection library):
go build -o hooprs ./cmd/hooprs
Quickstart
Scan everything, print the summary, write and open risk-report.html :
# only the last 30 days, plus a machine-readable JSON report
hooprs -days 30 -json risk-report.json
# only Cursor sessions whose project matches a pattern
hooprs -tools cursor -project ' my-app '
# raise the confidence bar (default 0.4)
hooprs -min-score 0.6
# show the actual leaked values in the terminal (never in the reports)
hooprs -show-values
All flags
Flag
Default
Description
-out
risk-report.html
Path for the self-contained HTML report
-json
(off)
Also write the machine-readable risk report here
-tools
claude,cursor,opencode
Sources to scan
-project
(all)
Regexp filter on session project
-session
(all)
Regexp filter on session id
-days
0 (all time)
Only sessions started within the last N days
-home
$HOME
Home directory to discover sessions under
-rules
(none)
Guardrails rules JSON file
-min-score
0.4
Minimum detection confidence (0 to 1) to count
-critical-weight
60
Security-score penalty weight (0 to 100) for the critical-session share
-engine
alcatraz
Detection engine: alcatraz (full PII set) or stub (zero-dependency fallback)
-incremental
false
Only scan content appended since the last run
-state
~/.risk-analyzer/state.json
Incremental scan state file
-quiet
false
Suppress the terminal summary
-show-values
false
Print the matched high-severity values for the top 10 critical sessions in the terminal summary (never written to the HTML/JSON reports)
-open
true
Open the HTML report in the default browser when done
-version
false
Print the hooprs version and exit
By default every run is a full snapshot of all your sessions. -incremental
persists per-file byte offsets so subsequent runs read only the content
appended since the last run (useful for "what changed since last time").
The report's title bar has two buttons, so a scan turns into a message to your
team in one click:
Share rasterizes a 1200×630 summary card (score, risk donut, headline
counts) to PNG and copies it to your clipboard, then shows the ready-to-paste
Slack message — paste the image, send the text. If your browser blocks
clipboard access, the same panel offers a PNG download.
Save PDF opens the print dialog with a print stylesheet that keeps the
branded dark look and drops the interactive chrome — pick "Save as PDF".
Both run entirely in the page: the card is built from the same value-free data
as the report (entity types and counts, never the matched values), and the
Slack message carries only aggregate counts plus the install one-liner. Sharing
a scan never re-leaks a leak.
Structured PII (via the alcatraz engine)
plus the secret types common in coding sessions (via hooprs's own secrets
pack):
Detection pairs regex patterns with checksum and format validators
(Luhn, IBAN mod-97, SSN/national-ID range rules), and matches below the
-min-score threshold are dropped. Pass -engine stub for a zero-dependency
regex fallback.
Note on NER: PERSON / LOCATION -style entities that need an NLP model
stay out of this version. The analyzer sits behind a small
analyze.Analyzer interface, so a future NLP-backed engine drops in without
touching the pipeline.
sources → analyze → risk → report
claude/ regex + tiers · terminal +
cursor/ validators score · html + json
opencode (local) exposure
Sources discover and parse each tool's on-disk session format into one
normalized model; analyze runs the detection engine over every message;
risk turns raw findings into tiers, exposure ranking, and a score;
report renders the terminal summary and the self-contained HTML.
Tier per session: critical (any high-severity entity), minor
(medium-severity only), or low .
Exposure ranks sessions by a severity-weighted finding count that favors
output (data pulled into the agent context) over input.
Security Score = clamp(0, 100, round(100 − W·critical/total − 20·minor/total)) ,
where W is the critical penalty weight ( -critical-weight , default 60).
Severity and data-family per entity type live in
risk/entities.go .
Layer your own detection rules with -rules <file> — direction-aware
( input = what you typed, output = assistant/tool output). See
examples/guardrails.json :
{
"rules" : [
{ "name" : " internal-hostnames " , "type" : " regex " , "direction" : " both " ,
"pattern" : " \\ b[a-z0-9-]+ \\ .internal \\ .example \\ .com \\ b " },
{ "name" : " private-key-material " , "type" : " deny_words " , "direction" : " output " ,
"words" : [ " BEGIN RSA PRIVATE KEY " ] }
]
}
Violations get their own section in the summary and the report, counted per
rule.
Everything runs on your machine. The HTML/JSON reports contain only entity
types, counts, severities, and session identifiers — never the matched values.
Nothing leaves your machine.
-show-values is the one deliberate exception: it prints the matched
high-severity values for the top 10 critical sessions to the terminal only ,
so you can locate the actual leaks. The HTML and JSON reports stay value-free
even with the flag on — and there's a test pinning that guarantee.
cmd/hooprs/ CLI: flags → discover → analyze → risk → render
sources/ discover & parse claude/cursor/opencode sessions
state/ incremental scan offsets
types/ normalized Session/Message model
analyze/ Analyzer interface + alcatraz engine, shared secrets pack, Stub fallback
guardrails/ local rules loader + direction-aware matcher
risk/ severity catalog + risk model (tiers, exposure, score)
report/ terminal + self-contained HTML renderer (embedded CSS/JS)
MIT © hoop.dev — built by the team behind hoop.
AI session risk summary. See what sensitive information is sitting in your local AI coding sessions. Runs locally, nothing leaves your machine.
hoop.dev/labs/ai-risk-summary
Topics
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
1
fork
Report repository
Releases
12
v0.2.6
Latest
Jul 3, 2026
+ 11 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
