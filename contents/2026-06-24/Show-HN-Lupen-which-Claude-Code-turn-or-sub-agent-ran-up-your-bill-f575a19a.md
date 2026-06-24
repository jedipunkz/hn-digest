---
source: "https://github.com/momoraul/Lupen"
hn_url: "https://news.ycombinator.com/item?id=48659162"
title: "Show HN: Lupen – which Claude Code turn or sub-agent ran up your bill"
article_title: "GitHub - momoraul/Lupen: Itemized cost receipts for Claude Code and Codex — by turn, verified, local. · GitHub"
author: "momoraul"
captured_at: "2026-06-24T13:43:33Z"
capture_tool: "hn-digest"
hn_id: 48659162
score: 1
comments: 0
posted_at: "2026-06-24T13:08:38Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Lupen – which Claude Code turn or sub-agent ran up your bill

- HN: [48659162](https://news.ycombinator.com/item?id=48659162)
- Source: [github.com](https://github.com/momoraul/Lupen)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T13:08:38Z

## Translation

タイトル: HN を表示: Lupen – クロード コードがどの請求書を提出するか、または代理人が請求書を作成するか
記事のタイトル: GitHub - momoraul/Lupen: Claude Code と Codex の明細化されたコスト領収書 — 順番に、検証済み、ローカル。 · GitHub
説明: Claude Code および Codex の明細化されたコスト領収書 — 順番に、検証済み、ローカル。 - モモロール/ルーペン

記事本文:
GitHub - momoraul/Lupen: Claude Code と Codex の明細化されたコスト領収書 — 順番に検証され、ローカルで確認されます。 · GitHub
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
モモラル
/
ルペン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード

main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
50 コミット 50 コミット .github .github Config Config Lupen.xcodeproj Lupen.xcodeproj Lupen Lupen LupenTests LupenTests スクリプト スクリプト ツール ツール docs docs .gitattributes .gitattributes .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード コードとコーデックスのセッションごとに実際にかかる費用を項目別、検証済み、ローカルで確認します。
Lupen は、生のクロード コードとコーデックスのログから直接支出を再計算します。ターン、ステップ、サブエージェントごとに分類され、トークンと照合され、Mac から離れることはありません。
インストール: brew install --cask momoraul/lupen/lupen — または署名された DMG を取得します。 Apple Silicon 上の macOS 26 (Tahoe) が必要です。
AI コーディングの費用は現在 50 ドルと言われています。どのプロバイダーでしたか？どのセッションですか?
どのターンですか？どのサブエージェントまたはツール ループですか?毎日の合計では答えられないので、
Lupen は数値を分解し、生のトークンから各コストを再計算します。
プロバイダー範囲の合計 — 今日 $50 · クロード コード · 12 セッション · 84 ターン。 Claude Code と Codex は、1 つの混合リストではなく、別個のモードのままです。
データが許可する場合、Turn、Step、SkillGroup、および SubAgent あたりのコスト。
再計算されてから差分が行われます。各コストは生のトークンと公開価格表から再計算され、報告された合計と比較されるため、差異は想定されるのではなく表示されます。
オリジンタグ付き添付ファイル — すべてのファイル、画像、URL は、コンテキストに入った場所によってラベルが付けられます。
サブエージェントのコストは親ターンにロールアップされ、単独で発生します。
明細化された領収書。セッション合計から各ターン、スキル グループ、サブエージェントまでドリルダウンします。各行の価格と帰属がわかります。

それ自体で。
オリジンタグ付きの添付ファイル。すべてのファイル、画像、URL には、コンテキストに入った場所によってラベルが付けられます。
プロバイダーごとの 1 日あたりの費用。時間の経過に伴うコスト、セッション、ターン、トークン。
検索可能で再開可能。過去のセッションに含まれるプロンプトによって過去のセッションを検索し、中断したところから再開します。
主な特徴
プロバイダー モード — クロード コードまたはコーデックスを選択します。 Lupen は、そのプロバイダーのセッション、会話、ドロップダウン合計、レポート、診断、検証結果のみを表示します。
コスト ドリフトの検証 - [使用状況の検証] ウィンドウを実行して、生のトークンと公開価格表からすべてのコストを再計算し、報告された合計と比較します。違いがある場合はフラグが立てられます。クロード コードはリクエストごとの合計に対してチェックされます。 Codex は、ロールアウト token_count イベントに対して独立したローカル検証ツールを取得します。
検索と再開 — 全文検索では、セッションに含まれるプロンプト (ターンごと) とそのプロジェクトおよびスラッグによってセッションが検索されます。 Claude Code (または Codex) で結果を再度開いて、中断したところから再開します。Lupen は新しいターミナル ウィンドウで claude --resume / codexresume を実行するか、コマンドをコピーします。
Anthropic の API と一致するターン境界 — Lupen は、タイムスタンプの代わりに stop_reason を使用してターン境界をマークします。ツール使用ループは、12 行に断片化されるのではなく、1 ターン内に留まります。
Codex ロールアウトのサポート — Lupen は ~/.codex/sessions/YYYY/MM/DD/rollout-*.jsonl を読み取り、キャッシュされた入力を正規化し、推論トークンを保存し、累積トークン デルタを処理し、新しい日のフォルダーをライブで監視します。
サブエージェントのコストのロールアップ — ターンがサブエージェントを生成すると、そのコストはアウトラインでは親にロールアップされますが、詳細ペインでは個別に帰属されたままになります。集計値とエージェントごとの数値は同じソースから取得されているため、一貫性が保たれます。
オリジンタグ付き添付ファイルの追跡 — Fil

パス、画像バイト、および URL は、会話に入った場所 (インライン プロンプト、ツール入力、ツール出力、応答など) によって分類されるため、コンテキスト ウィンドウに何が表示されているかを確認できます。
5 時間の制限追跡 — 過去 7 日間で消費した制限の 1 % あたりのドルのベイズ推定値が、メニュー バー アイコンのリングの色合い (70 % で黄色、90 % でオレンジ、100 % で赤) で表示されます。
スクリプト可能な lupen CLI — 同じアプリのバイナリは、同じローカル インデックス上のコマンド ラインです。テーブル、 --json 、または --csv としてのすべてのレポートに加え、CI ステップまたはコミット フックの検証/予算終了コード ゲートが含まれます。 「コマンドライン」を参照してください。
ゼロネットワーク — Lupen は、ディスク上のローカルのクロード コードおよびコーデックス ファイルのみを読み取ります。 API キー、テレメトリ、クラウド同期はありません。
brew install --cask momoraul/lupen/lupen
…または、最新リリースから署名された DMG を取得します。どちらも公証されており、Sparkle を通じて最新の状態に保たれます。
git clone https://github.com/momoraul/Lupen.git
cdルーペン
cp Config/Local.xcconfig.example Config/Local.xcconfig # set DEVELOPMENT_TEAM
xcodebuild build -project Lupen.xcodeproj -scheme Lupen -destination ' platform=macOS '
~ /Library/Developer/Xcode/DerivedData/Lupen- * /Build/Products/Debug/Lupen.app を開きます
コマンドライン
アプリのバイナリは lupen CLI としても機能します - メニューバー アプリと同じローカル インデックス
ビルド、すべてのレポートはスクリプト可能でローカルです:
lupen スキル -- 最後の 30 日 # スキルごとのコスト、$/実行、トップモデル
lupen top --by session --limit 5 # 最もコストのかかるセッション
lupen Budget --over 20 --last 7d # 今週が $20 を超えた場合は exit 4
lupen verify # 再計算された真実からコストがずれている場合は 4 を終了します
ルーペン毎日 --json | jq # JSON / CSV としてのレポート
グループ化された完全なコマンド セット:
すべてのレポート コマンドは、 --provider 、ピリオド ( --last 30d / --month 2026-06 / --since … --until … )、および --json / --csv を使用します (verify を除く)。

、
これはコーパス全体を監査し、ピリオドを無視します。
Homebrew をインストールすると、lupen が自動的に PATH に配置されます。
DMG またはソース ビルドで、lupen install-cli を 1 回実行します。完全なリファレンス — すべて
フラグと終了コード: docs/CLI.md 。
クロード コードの 1 回のセッションで、その日の残りの時間を合わせたよりも多くの費用がかかったことがありました。
そして、毎日の合計では、どのターン、またはどの暴走サブエージェントが発生したのかわかりませんでした。
お金を食べた。 Lupen は私が欲しかった明細化された領収書です。毎ターン価格が設定されており、
私の Mac から何も出ることなく、すべての価格が生のトークンと照合してチェックされました。
Lupen は Mac 上のローカル セッション ファイルを読み取ります。
クロードコード: ~/.claude/projects/**/*.jsonl
コーデックス: $CODEX_HOME/sessions/**/rollout-*.jsonl または ~/.codex/sessions/**/rollout-*.jsonl
ネットワーク要求はまったく行われません。テレメトリ、分析、クラウド同期はありません。
会話、プロンプト、ファイル パス、添付ファイルが外部に残ることはありません。
機械。 Info.plist には NSAppTransportSecurity ブロックが含まれていません。
Lupen はソケットを開きません。これは Little Snitch などで確認できます。
アウトバウンド接続モニター。
macOS 26 (Tahoe) 以降、Apple Silicon
Xcode 26 と Swift 6 (ソースからのみビルド)
ローカル セッション データを含むアクティブな Claude Code または Codex インストール
docs/CLAUDE-CODE-TOKEN-GUIDE.md — ~/.claude/projects/ JSONL スキーマ + トークンフィールドの解釈。
docs/CODEX-LOCAL-DATA.md — ~/.codex/sessions/ ロールアウト JSONL スキーマ + トークンフィールド解釈。
docs/TOKEN-BILLING-EXPLAINED.md — Anthropic のトークン数、キャッシュの節約、および請求可能な合計がどのように関係するか。
ワークフロー、コミットスタイル、
また、サニタイズされた JSONL 再現でバグを報告する方法も説明します。
マサチューセッツ工科大学著作権 © 2026 ジェイデン (@momoraul)。 「ライセンス」を参照してください。
Lupen は独立したプロジェクトです。 Anthropic や
オープンAI ;あなたの町に書き込まれたローカルログファイルのみを読み取ります

ね。
Sparkle — 自動更新フレームワーク (MIT)。
GRDB.swift — オンデバイス インデックス (MIT) 用の SQLite ツールキット。
落ち着いた雰囲気の [会話] タブと、より信頼性の高いメニュー バー。
会話タブの再デザイン — 静かなカード (ヘアラインの境界線 + かすかな塗りつぶしなので、
ボックスと競合するのではなく対話が先導します）、ヘッダーが後退し、
アクセントで囲まれた選択されたステップ。本文のタイポグラフィ、読み取り幅、コード/表
読みやすさも同時に調整されました。
よりクリーンな取り付けインジケーター - 1 つの読みやすいペーパークリップのマークが回転します。
(画像ごとに 1 つのグリフではなく) 画像を添付すると、読み取り可能な状態が維持されます。
ライトモードとダークモードの両方で選択された行。
メニュー バー — メニュー バー項目をクリックすると、ウィンドウが確実に表示されるようになりました。
前面とアイドル日には、ほとんど目に見えない薄暗くなった金額ではなく、明確な「$0」が表示されます。
CLI クラッシュの緊急修正と、より正確な Codex 使用状況の検証。
CLI クラッシュの修正 — lupen CLI は起動時にクラッシュしなくなりました (SIGTRAP
CLI 実行の 100% に達します);ロギングはどのスレッドからでも安全にブートストラップされるようになりました。
Codex Verify 使用量の精度 — 誤検知「使用量の欠落」を阻止
レポートを作成し、生成されたターンのリクエスト ID の一致を修正し、エラー/警告を追加します。
Verify Costs フィルタおよび Verify CLI ゲートに対する重大度。
非常に大きなターンとコールドローンチ選択時のフリーズを修正しました。
大回転パフォーマンス — [会話] タブがフリーズしなくなりました
数千ステップのターン。サポート活動の長期実行 (ツール、思考)
カードの数が制限されたままになるように、折りたたまれたグループに折りたたまれます。
コールド起動の選択 — 最初のセッションがコールド起動時に自動選択されるようになりました。
何も選択せずに開くのではなく、適切なフォーカス リングを使用して起動します。
[会話] タブを機能豊富なターン リーダーとして再考します。
会話リッチ リーダー — 選択したターンがカード スタックとしてレンダリングされるようになりました
(プロンプト → 考える → ツール → 返信) インスタ

2 つの平文ブロックの d、
豊富なマークダウン (見出し、リスト、コピー付きのコード ブロック、表、引用符)
複数行の選択。
キュレーション — 思考とツール呼び出しは 1 行の開示にまとめられます。
拡張することができます。中断 / API エラー / 圧縮されたターンで明確なステータスが表示される
「(返答はありません)」の代わりにバナーを表示します。
読みやすさ — プロンプトと返信はカードとして目立ち、サイドコンテンツはカードとして目立ちます。
（思考、道具）が凹んだ線に消えていきます。選択したステップが強調表示されます
そしてスクロールして表示されました。
ストレージ マネージャーを追加し、Codex サイドバーをシャープにします。
「セッションとストレージの管理」ウィンドウ (⌘⇧M) — クロード コードの参照とクリーンアップ
Lupen キャッシュ検査 (インデックス / WAL /
スナップショット）と読み取り専用の全ディスク ビュー。削除はゴミ箱のみで、
ホワイトリスト + 元に戻す; auth/config/state ファイルはハードブロックされています。
Codex のファースト プロンプト タイトル - session_index.jsonl のないセッション
スレッド名には、生の ID プレフィックスの代わりに最初のユーザー プロンプトが表示されるようになりました。
ゼロコストのセッション隠蔽 - 低信号の $0 セッション (Codex
自動レビュー/ガーディアン、アイドル) プロジェクトごとの小さな (N) トグルの後ろで折りたたまれます
グループ。アクティブな、選択された、コスト集計のないセッションは表示されたままになります。
コストの再着色 — コスト ≥ $1 はオレンジ色 (暗いところでは柔らかくなります) を読み取り、N/A は読み取ります
1ドル未満は薄暗いままだ。色は明/暗および反転ごとに適応します

[切り捨てられた]

## Original Extract

Itemized cost receipts for Claude Code and Codex — by turn, verified, local. - momoraul/Lupen

GitHub - momoraul/Lupen: Itemized cost receipts for Claude Code and Codex — by turn, verified, local. · GitHub
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
momoraul
/
Lupen
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
50 Commits 50 Commits .github .github Config Config Lupen.xcodeproj Lupen.xcodeproj Lupen Lupen LupenTests LupenTests Scripts Scripts Tools Tools docs docs .gitattributes .gitattributes .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md View all files Repository files navigation
See what every Claude Code and Codex session actually costs — itemized, verified, local.
Lupen recomputes your spend straight from the raw Claude Code and Codex logs — broken down by turn, step, and sub-agent, checked against the tokens, and never leaving your Mac.
Install: brew install --cask momoraul/lupen/lupen — or grab the signed DMG . Requires macOS 26 (Tahoe) on Apple Silicon.
Your AI coding spend says $50 today . Which provider was it? Which session?
Which turn? Which sub-agent or tool loop? A daily total can't answer that, so
Lupen breaks the number down and recomputes each cost from the raw tokens.
Provider-scoped totals — $50 today · Claude Code · 12 sessions · 84 turns . Claude Code and Codex stay in separate modes instead of one mixed list.
Cost per Turn, Step, SkillGroup, and SubAgent , where the data allows it.
Recomputed, then diffed — each cost is recomputed from raw tokens and the public price table and compared to the reported total, so any difference is shown rather than assumed.
Origin-tagged attachments — every file, image, and URL is labelled by where it entered the context.
Sub-agent cost rolls up into the parent turn and stays attributable on its own.
The itemized receipt. Drill from a session total down to each turn, skill group, and sub-agent — every row priced and attributable on its own.
Origin-tagged attachments. Every file, image, and URL is labelled by where it entered the context.
Daily spend, per provider. Cost, sessions, turns, and tokens over time.
Searchable and resumable. Find a past session by any prompt it contains, then pick up right where you left off.
Key features
Provider mode — Choose Claude Code or Codex. Lupen shows only that provider's sessions, conversations, dropdown totals, reports, diagnostics, and verification results.
Cost drift verification — Run the Verify Usage window to recompute every cost from raw tokens and the public price table, then diff it against the reported total; any difference is flagged. Claude Code is checked against its per-request totals; Codex gets an independent local verifier over rollout token_count events.
Search and resume — Full-text search finds a session by any prompt it contains (across every turn), plus its project and slug. Reopen any result in Claude Code (or Codex) to pick up where you left off — Lupen runs claude --resume / codex resume in a new Terminal window, or copies the command for you.
Turn boundaries that match Anthropic's API — Lupen uses stop_reason to mark Turn boundaries instead of timestamps. Tool-use loops stay inside one Turn instead of fragmenting into a dozen rows.
Codex rollout support — Lupen reads ~/.codex/sessions/YYYY/MM/DD/rollout-*.jsonl , normalizes cached input, preserves reasoning tokens, handles cumulative token deltas, and watches new day folders live.
Sub-agent cost rollup — When a Turn spawns sub-agents, their cost rolls up into the parent in the outline but stays separately attributable in the detail pane. The aggregate and the per-agent figures come from the same source, so they stay consistent.
Origin-tagged attachment tracking — File paths, image bytes, and URLs are classified by where they entered the conversation (inline prompt, tool input, tool output, reply, …) so you can see what's filling the context window.
5-hour limit tracking — A Bayesian estimate of $ per 1 % of limit consumed across your last 7 days, surfaced in the menu-bar icon's ring tint (yellow at 70 %, orange at 90 %, red at 100 %).
Scriptable lupen CLI — The same app binary is a command line over the same local index: every report as a table, --json , or --csv , plus verify / budget exit-code gates for a CI step or commit hook. See Command line .
Zero network — Lupen only reads local Claude Code and Codex files on disk. No API keys, no telemetry, no cloud sync.
brew install --cask momoraul/lupen/lupen
…or grab the signed DMG from the latest release . Both are notarized and keep themselves up to date via Sparkle.
git clone https://github.com/momoraul/Lupen.git
cd Lupen
cp Config/Local.xcconfig.example Config/Local.xcconfig # set DEVELOPMENT_TEAM
xcodebuild build -project Lupen.xcodeproj -scheme Lupen -destination ' platform=macOS '
open ~ /Library/Developer/Xcode/DerivedData/Lupen- * /Build/Products/Debug/Lupen.app
Command line
The app binary doubles as a lupen CLI — same local index the menu-bar app
builds, every report scriptable and local:
lupen skills --last 30d # per-skill cost, $/run, top model
lupen top --by sessions --limit 5 # the costliest sessions
lupen budget --over 20 --last 7d # exit 4 if this week ran over $20
lupen verify # exit 4 if any cost drifts from the recomputed truth
lupen daily --json | jq # any report as JSON / CSV
The full command set, grouped:
Every reporting command takes --provider , a period ( --last 30d / --month 2026-06 / --since … --until … ), and --json / --csv — except verify ,
which audits the whole corpus and ignores the period.
Homebrew installs put lupen on your PATH automatically;
on a DMG or source build, run lupen install-cli once. Full reference — every
flag and exit code: docs/CLI.md .
A single Claude Code session once cost me more than the rest of the day combined,
and the daily total couldn't tell me which turn — or which runaway sub-agent —
ate the money. Lupen is the itemized receipt I wanted: every turn priced, and
every price checked against the raw tokens, without anything leaving my Mac.
Lupen reads local session files on your Mac:
Claude Code: ~/.claude/projects/**/*.jsonl
Codex: $CODEX_HOME/sessions/**/rollout-*.jsonl or ~/.codex/sessions/**/rollout-*.jsonl
It makes zero network requests . No telemetry, no analytics, no cloud sync.
Your conversations, prompts, file paths, and attachments never leave your
machine. The Info.plist carries no NSAppTransportSecurity block because
Lupen opens no sockets — you can confirm this with Little Snitch or any
outbound-connection monitor.
macOS 26 (Tahoe) or later, Apple Silicon
Xcode 26 with Swift 6 (build from source only)
An active Claude Code or Codex installation with local session data
docs/CLAUDE-CODE-TOKEN-GUIDE.md — ~/.claude/projects/ JSONL schema + token-field interpretation.
docs/CODEX-LOCAL-DATA.md — ~/.codex/sessions/ rollout JSONL schema + token-field interpretation.
docs/TOKEN-BILLING-EXPLAINED.md — How Anthropic's token counts, cache savings, and billable totals relate.
See CONTRIBUTING.md for the workflow, commit style,
and how to file a bug with a sanitised JSONL repro.
MIT. Copyright © 2026 jaden (@momoraul). See LICENSE .
Lupen is an independent project. It is not affiliated with Anthropic or
OpenAI ; it only reads local log files written to your machine.
Sparkle — auto-update framework (MIT).
GRDB.swift — SQLite toolkit for the on-device index (MIT).
A calmer Conversation tab and a more reliable menu bar.
Conversation tab redesign — quiet cards (a hairline border + faint fill so
the dialogue leads instead of competing with boxes), receding headers, and an
accent-bordered selected step; body typography, reading width, and code/table
legibility were tuned alongside.
Cleaner attachment indicator — a single legible paperclip marks turns with
image attachments (instead of one glyph per image), and it stays readable on
the selected row in both light and dark mode.
Menu bar — clicking the menu-bar item now reliably brings the window to the
front, and idle days show a clear "$0" instead of a near-invisible dimmed amount.
Urgent CLI crash fix plus more accurate Codex usage verification.
CLI crash fix — the lupen CLI no longer crashes on launch (a SIGTRAP
that hit 100% of CLI runs); logging now bootstraps safely from any thread.
Codex Verify Usage accuracy — stops false-positive "missing usage"
reports, fixes generated-turn request-id matching, and adds error/warning
severity to the Verify Costs filter and the verify CLI gate.
Fixes a freeze on very large turns and cold-launch selection.
Large-turn performance — the Conversation tab no longer freezes on
multi-thousand-step turns; long runs of supporting activity (tools, thinking)
fold into a collapsed group so the card count stays bounded.
Cold-launch selection — the first session is now auto-selected on a cold
launch with a proper focus ring, instead of opening with nothing selected.
Reimagines the Conversation tab as a rich Turn reader.
Conversation rich reader — a selected Turn now renders as a card stack
(prompt → thinking → tools → reply) instead of two plain-text blocks, with
rich markdown (headings, lists, code blocks with copy, tables, quotes) and
multi-line selection.
Curation — thinking and tool calls collapse into one-line disclosures you
can expand; interrupted / API-error / compacted turns show a clear status
banner instead of "(no response available)".
Readability — prompts and replies stand out as cards while side content
(thinking, tools) fades into indented lines; the selected step is highlighted
and scrolled into view.
Adds a storage manager and sharpens the Codex sidebar.
Manage Sessions & Storage window (⌘⇧M) — browse and clean up Claude Code
and Codex sessions by disk usage, with Lupen cache inspection (index / WAL /
snapshot) and a read-only all-disk view. Deletion is trash-only with an
allowlist + Undo; auth/config/state files are hard-blocked.
Codex first-prompt titles — sessions without a session_index.jsonl
thread name now show their first user prompt instead of the raw id prefix.
Zero-cost session hiding — low-signal $0 sessions (Codex
auto-review/guardian, idle) collapse behind a small (N) toggle per project
group; active, selected, and no-cost-aggregate sessions stay visible.
Cost recoloring — cost ≥ $1 reads orange (softer on dark), N/A reads
slate, sub-$1 stays dim; colors adapt per light/dark and inver

[truncated]
