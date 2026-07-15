---
source: "https://github.com/sashamitrovich/milepost"
hn_url: "https://news.ycombinator.com/item?id=48917147"
title: "Milepost – plain-Markdown long-term memory for Claude Code"
article_title: "GitHub - sashamitrovich/milepost: Long-term memory for Claude Code — a milestone-triggered, in-session work diary in plain markdown. No database, no lock-in, fully greppable. · GitHub"
author: "sashamitrovich"
captured_at: "2026-07-15T07:07:01Z"
capture_tool: "hn-digest"
hn_id: 48917147
score: 1
comments: 0
posted_at: "2026-07-15T06:53:05Z"
tags:
  - hacker-news
  - translated
---

# Milepost – plain-Markdown long-term memory for Claude Code

- HN: [48917147](https://news.ycombinator.com/item?id=48917147)
- Source: [github.com](https://github.com/sashamitrovich/milepost)
- Score: 1
- Comments: 0
- Posted: 2026-07-15T06:53:05Z

## Translation

タイトル: Milepost – クロード コードの単純な Markdown 長期記憶
記事のタイトル: GitHub - sasamitrovich/milepost: Claude Code の長期記憶 — マイルストーンによってトリガーされる、プレーンなマークダウンでのセッション中の作業日記。データベースやロックインはなく、完全に grepable です。 · GitHub
説明: Claude Code の長期記憶 — マイルストーンによってトリガーされる、プレーンなマークダウンでのセッション中の作業日記。データベースやロックインはなく、完全に grepable です。 - ササミトロヴィッチ/マイルポスト

記事本文:
GitHub - sasamitrovich/milepost: Claude Code の長期記憶 — マイルストーンによってトリガーされる、プレーンなマークダウンでのセッション中の作業日記。データベースやロックインはなく、完全に grepable です。 · GitHub
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
ササミトロビッチ
/
マイルポスト
公共
通知
君はね

サインインして通知設定を変更できない
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット .claude-plugin .claude-plugin コマンド コマンド フック フック ポリシー ポリシー .gitignore .gitignore ライセンス ライセンス README.md README.md install.sh install.sh uninstall.sh uninstall.sh すべてのファイルを表示 リポジトリ ファイル ナビゲーション
📔 マイルポスト — クロード コードの長期記憶
マイルストーンによってトリガーされるセッション内の作業日記により、クロード コードにセッション間で永続的な単純なマークダウン メモリが提供されます。データベースやロックインはなく、完全に grepable です。
キーワード: クロード コード メモリ、クロード コードの永続メモリ、長期メモリ プラグイン、AI エージェント ダイアリー、セッション メモリ、マークダウン作業ログ、クロード コード プラグイン、エージェント コンテキストの永続性、マイルストーン ログ
milepost は、クロードを記憶を持つエージェントに変える、無料のオープンソース クロード コード プラグインです。作業中、Claude は、意味のあるマイルストーン (完了した成果物、主要な決定 (およびその理由)、阻害要因、方向の変更) の日記に加えて、状況を示す生きた STATUS.md を自動的に作成します。すべては所有するプレーンなマークダウン ファイルとして保存されるため、grep、git、エクスポート、または任意のエディターで読み取ることができます。 SQLite、ベクター DB、クラウド、ロックインはありません。
数日または数週間後に長期にわたるタスクを開始すると、クロードはあなたが何をしたのか、何を決定したのか、そして次に何をするのかをすでに知っています。
/プラグインマーケットプレイス追加sasamitrovich/milepost
/プラグインのインストール milepost@milepost
マイルポストが存在する理由
Claude Code はセッション中は素晴らしく、セッション中は忘れがちです。ターミナルを閉じると、すでに除外した推論、決定、行き止まりが消えます。簡単な編集であれば問題ありません。数週間にわたって戻ってくるプロジェクトの場合、私は

これは、毎回コンテキストを再説明することを意味します。
既存のメモリ ツールはそれぞれトレードオフを強いられます。
データベースにバックアップされたメモリ (claude-mem など) はすべてを自動的にキャプチャしますが、メモリは SQLite + ベクトル ストア内に存在します。読みにくく、エクスポートしにくく、ロックインされやすい。
セッション終了/手動日記 (claude-diary など) はプレーンなマークダウンを書き込みますが、それはコマンドを実行するとき、またはセッションがすでに終了した後に限られます。
milepost は 3 番目のパスを選択します。つまり、時計、コマンド、データベースではなく、マイルストーンによってトリガーされる、単純なマークダウンです。
「本当のマイルストーンが起こった」と判断するのは判断力であり、それを決定できるのはモデルだけです。したがって、マイルポストは愚かな cron の仕事ではありません。それは 2 つの連携部分です。
その結果、フローを中断することなく、何かを実行することを忘れずに済む、マイルストーンを意識した自動的なジャーナリングが実現します。
プロジェクトごとに、あなたが完全に制御するフォルダー内に次のものが含まれます。
~/.claude/memory/milepost/<プロジェクト>/
§── log.md # 追加専用、タイムスタンプ付きマイルストーン エントリ
§── STATUS.md # 生きたスナップショット: 現在の目標、現状、次の目標
└── 反射/ # オプション /反射合成
実際のエントリは次のようになります。
## 2026-06-27 14:32 — アーキテクト PDF をページごとの JPEG に変換しました
** タイプ: ** 納品可能
** 内容: ** ImageMagick を使用して、A3 の 4 ページすべてを 200 DPI でレンダリングしました。
** 理由: ** クライアントはレビューのために個別のページ画像を必要としていました。
** 次へ: ** —
人間が読める。機械で解析可能。あなたが所有しています。
🪶 プレーンなマークダウン、ゼロデータベース — greppable、git フレンドリー、エクスポート可能、将来性あり。
🎯 時間トリガーではなくマイルストーントリガー — ノイズの多いログではなく、シグナルの高いエントリ。
🔁 セッション中および自動 — ターミナルを閉じた後ではなく、クロードの作業中に書き込まれます。
🌍 すべてのプロジェクトで機能します - 1 つのグローバル インストール。プロジェクトごと

日記。
🧭 リビング ステータス ファイル — タスクの現在の状態と次のステップを常に把握します。
🛠️ 手動コントロール — /milepost 、 /status 、 /reflect を必要なときに使用できます。
🔒 プライバシー第一、強制 - ポリシーはクロードに秘密を要約するように指示し、PreToolUse ガード フックは秘密形式のコンテンツが日記に表示されるのを機械的にブロックします。 .no-milepost ファイルを介したプロジェクトごとのオプトアウト。
↩️ 可逆的で安全 — 自動バックアップを備えたべき等インストーラー。データを保持するクリーンなアンインストーラー。
オプション A — プラグインのインストール (推奨)。クロードコード内:
/プラグインマーケットプレイス追加sasamitrovich/milepost
/プラグインのインストール milepost@milepost
コマンド、ナッジ、リコール、シークレット ガードはプラグインからロードされ、ポリシーはセッション開始時に挿入されます (CLAUDE.md には一切触れません)。クロードの日記を初めて作成する場合は、日記の書き込みを「常に許可」で承認してください。プラグインは事前にアクセス許可を与えることはできません。 /plugin uninstall milepost@milepost で削除します。日記はそのまま残ります。
オプション B — スクリプトのインストール。古いクロード コード バージョンの場合、またはポリシーを独自の CLAUDE.md 内の表示ブロックとして使用し、事前に付与された日記書き込み権限を使用する場合:
git clone https://github.com/sashamitrovich/milepost.git
cd milepost && bash install.sh # idempotent;バックアップとマージを行い、決して上書きしません
bash uninstall.sh で削除 — 日記はそのまま残ります。
両方ではなく 1 つのインストール モードを使用します。両方を実行するとフックが複製されます (無害ですがノイズが発生します)。どちらの方法でも日記の場所と形式は同じなので、何も失うことなくモードを切り替えることができます。
コマンド
何をするのか
/マイルポスト [フォーカス]
今すぐマイルストーンエントリを強制します。
/ステータス[更新]
このプロジェクトのタスクの現在のステータスを表示し、更新します。
/反射する
進捗状況と重要な決定を総合する

、ログからの繰り返しのテーマ。
構成
設定
デフォルト
説明
MILEPOST_NUDGE_INTERVAL
900（15分）
プロジェクトごとのナッジ間の最小秒数。ジャーナルに上げる頻度を減らします。もっとジャーナルに下げてください。
.no-マイルポスト ファイル
—
空の .no-milepost ファイルをプロジェクトのルートにドロップすると、そのプロジェクトが完全に除外されます。日記もナッジもリコールもありません。クライアントワークや、日記に残したくない内容の場合。
MILEPOST_STRICT
0
1 に設定すると、シークレット ガードは IP アドレスや電子メール アドレスなどの準秘密もブロックします。多くの正規の日記は内部ホストを参照しているため、オプトインします。クロード コードのフックが認識する場所に設定します。例: "env": {"MILEPOST_STRICT": "1"} の settings.json 内、またはシェル プロファイル内。
ゼロフックの方がいいですか?ナッジを完全にスキップします。CLAUDE.md ポリシーは単独で機能します。非常に長いセッション中に記憶するためにモデルに頼るだけです。
マイルポスト
クロード・メム
クロード日記
ストレージ
プレーンな値下げ
SQLite + ベクター DB
プレーンな値下げ
グレップ可能/エクスポート可能
✅
⚠️輸出経由
✅
セッション中の書き込み
✅
✅
❌
マイルストーンを意識した (高シグナル)
✅
⚠️すべてをキャプチャします
⚠️セッションレベル
トリガー
マイルストーン
継続的
マニュアル/プレコンパクト
生存状況ファイル
✅
❌
❌
データベースが必要です
いいえ
はい
いいえ
正直な注意点: マイルストーンの検出は判断であるため、ベストエフォートで行われます。ログが過剰または不足する場合があり、エントリは作業セッションのトークンを少し使用します。 /milepost は手動バックストップです。
クロード・コードに長期記憶を与える最善の方法は何ですか?
コンテキストを耐久性のあるポータブルストレージに永続化するプラグインを使用します。 milepost は、作業中に作成されたプレーンなマークダウン ファイルを使用してこれを実行するため、メモリは読み取り可能で、grep 可能であり、データベース内にロックされることはありません。
クロード・コードはセッション間のことを覚えていますか?
それ自体ではありません。クロード・コー

de は CLAUDE.md を超えるセッション間ではステートレスです。 milepost は、マイルストーンとライブ ステータス ファイルをジャーナリングすることで、真のクロスセッション メモリを追加し、後でクロードが読み返します。
マイルポストはクロードメムとどう違うのですか?
claude-mem は、すべてを SQLite + ベクター データベースにキャプチャします。 milepost は、シグナルの高いマイルストーンを単純なマークダウンとして保存します。これにより、読みやすく、バージョン付けし、エクスポートし、信頼しやすくなり、ロックするものがなくなります。
マイルポストはデータベースを使用しますか、またはデータをどこかに送信しますか?
いいえ。すべては ~/.claude/memory/milepost/ の下のローカル マークダウン ファイルに残ります。データベースもクラウドもテレメトリもありません。
秘密や機密コードをログに記録しますか?
単なる丁寧な指示ではなく、多層防御:
ポリシー — クロードは、シークレットを値では決して参照せず、場所のみ (「.ha_token に保存されたトークン」) で参照し、準シークレット (内部 IP、シリアル、サードパーティの個人データ) を最小限に抑えるように指示されています。
ガード フック — PreToolUse フックは、日記ディレクトリ ( Bash を介したシェル リダイレクトを含む) を対象としたすべての書き込みをスキャンし、秘密キー ブロック、AWS/GitHub/Slack/Google/ sk- スタイルのキー、JWT、ベアラー トークン、password= / api_key: スタイルの割り当て、認証情報が埋め込まれた URL、さらに不明なトークン形式 (32 文字以上) に対する高エントロピーのキャッチオールなどのシークレット パターンに一致するものをすべて拒否します。 Base64url-ish は、上位/下位/数字を混合して実行します (git SHA、UUID、パス、識別子によってトリップされないように調整されています)。
厳密モード — MILEPOST_STRICT=1 は、クライアントの作業やプライバシーに配慮した設定のために、IP アドレスと電子メール アドレスをさらにブロックします。
監査 — いつでも既存の日記をスキャンできます: bash ~/.claude/hooks/milepost-secret-scan.sh < ~/.claude/memory/milepost/<project>/log.md (ヒット時の exit 1 + パターン名)。
オプトアウト — プロジェクトのルートにある .no-milepost ファイルにより、そのプロジェクトは完全に日記から除外されます。
日記はどのように保護されていますか

ディスク上で？
日記は平文のマークダウンであり、あなたが所有します。インストーラーは、 ~/.claude/memory/milepost/ を所有者のみのアクセス ( chmod 700 ) に制限します。保存時の暗号化は意図的にありません。milepost の全体的な設計は、クロードが通常のファイル ツールを使用して日記を読み書きすることです。暗号化すると、キー管理を追加しながらそれ (および grepability) が壊れます。デバイスの盗難のケースを適切にカバーするフルディスク暗号化 (FileVault/LUKS/BitLocker) を使用します。日記ディレクトリをシェル履歴のように扱います。ローカル、プライベート、コミットも共有ストレージへの同期もありません。
悪意のあるリポジトリがクロードを騙して私の日記を漏らす可能性はありますか?
The theoretical risk with any memory system: instructions hidden in a repo's files could ask the agent to copy remembered content somewhere it doesn't belong (prompt injection). milepost のポリシーでは、日記の内容をプロジェクト ファイルや外部サービスにコピーすることを明示的に禁止しており、日記にはマイルストーンの概要のみが含まれており、資格情報は含まれていません (上記のガードを参照)。概要であっても機密性が高いプロジェクトの場合は、漏洩するものが何もないように .no-milepost を使用してください。いつものように、エージェントが書いた内容を確認します。
日記を書くたびに承認する必要がありますか?
いいえ。インストーラーは、2 つの狭い範囲の権限許可ルールを settings.json に追加します — Edit(~/.claude/memory/milepost/**) と Write(~/.claude/memory/milepost/**) — これにより、クロードはプロンプトなしでジャーナリングできるようになります

[切り捨てられた]

## Original Extract

Long-term memory for Claude Code — a milestone-triggered, in-session work diary in plain markdown. No database, no lock-in, fully greppable. - sashamitrovich/milepost

GitHub - sashamitrovich/milepost: Long-term memory for Claude Code — a milestone-triggered, in-session work diary in plain markdown. No database, no lock-in, fully greppable. · GitHub
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
sashamitrovich
/
milepost
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits .claude-plugin .claude-plugin commands commands hooks hooks policy policy .gitignore .gitignore LICENSE LICENSE README.md README.md install.sh install.sh uninstall.sh uninstall.sh View all files Repository files navigation
📔 milepost — Long-Term Memory for Claude Code
A milestone-triggered, in-session work diary that gives Claude Code persistent, plain-markdown memory across sessions — no database, no lock-in, fully greppable.
Keywords: Claude Code memory · persistent memory for Claude Code · long-term memory plugin · AI agent diary · session memory · markdown work log · Claude Code plugin · agent context persistence · milestone logging
milepost is a free, open-source Claude Code plugin that turns Claude into an agent with a memory. As you work, Claude automatically writes a diary of meaningful milestones — completed deliverables, key decisions (and why they were made), blockers, and changes in direction — plus a living STATUS.md of where things stand. Everything is stored as plain markdown files you own , so you can grep it, git it, export it, or read it in any editor. No SQLite, no vector DB, no cloud, no lock-in.
Pick up a long-running task days or weeks later and Claude already knows what you did, what you decided, and what's next.
/plugin marketplace add sashamitrovich/milepost
/plugin install milepost@milepost
Why milepost exists
Claude Code is brilliant within a session — and forgetful between them. Close the terminal and the reasoning, the decisions, the dead-ends you already ruled out are gone. For a quick edit that's fine. For a project you return to over weeks , it means re-explaining context every single time.
Existing memory tools each force a trade-off:
Database-backed memory (e.g. claude-mem) captures everything automatically — but your memory lives in SQLite + a vector store. Hard to read, harder to export, easy to get locked into.
Session-end / manual diaries (e.g. claude-diary) write plain markdown — but only when you run a command, or after a session is already over.
milepost takes the third path: plain markdown, written as you go , triggered by milestones — not by a clock, a command, or a database.
Deciding "a real milestone just happened" is a judgment call — and only the model can make it. So milepost isn't a dumb cron job; it's two cooperating parts:
The result: automatic, milestone-aware journaling that doesn't interrupt your flow and doesn't depend on you remembering to run anything.
Per project, in a folder you fully control:
~/.claude/memory/milepost/<project>/
├── log.md # append-only, timestamped milestone entries
├── STATUS.md # living snapshot: current goal, where things stand, what's next
└── reflections/ # optional /reflect syntheses
A real entry looks like this:
## 2026-06-27 14:32 — Converted architect PDF to per-page JPEGs
** Type: ** deliverable
** What: ** Rendered all 4 A3 pages at 200 DPI with ImageMagick.
** Why: ** Client needed individual page images for review.
** Next: ** —
Readable by humans. Parseable by machines. Owned by you.
🪶 Plain markdown, zero database — greppable, git-friendly, exportable, future-proof.
🎯 Milestone-triggered, not time-triggered — high-signal entries, not noisy logs.
🔁 In-session & automatic — written while Claude works, not after you've closed the terminal.
🌍 Works across every project — one global install; per-project diaries.
🧭 Living status file — always know the current state and next step of any task.
🛠️ Manual controls — /milepost , /status , /reflect when you want them.
🔒 Privacy-first, enforced — policy tells Claude to summarize secrets, and a PreToolUse guard hook mechanically blocks secret-shaped content from ever landing in the diary. Per-project opt-out via a .no-milepost file.
↩️ Reversible & safe — idempotent installer with automatic backups; clean uninstaller that keeps your data.
Option A — plugin install (recommended). Inside Claude Code:
/plugin marketplace add sashamitrovich/milepost
/plugin install milepost@milepost
Everything is wired natively: the commands, the nudge, the recall, and the secret guard load from the plugin, and the policy is injected at session start (your CLAUDE.md is never touched). First time Claude journals, approve the diary write with "always allow" — the plugin cannot pre-grant permissions for you. Remove with /plugin uninstall milepost@milepost ; your diaries stay put.
Option B — script install. For older Claude Code versions, or if you prefer the policy as a visible block in your own CLAUDE.md and pre-granted diary-write permissions:
git clone https://github.com/sashamitrovich/milepost.git
cd milepost && bash install.sh # idempotent; backs up & merges, never overwrites
Remove with bash uninstall.sh — diaries stay put.
Use one install mode, not both — running both duplicates the hooks (harmless but noisy). The diary location and format are identical either way, so you can switch modes without losing anything.
Command
What it does
/milepost [focus]
Force a milestone entry right now.
/status [update]
Show — and refresh — the current status of this project's task.
/reflect
Synthesize progress, key decisions, and recurring themes from the log.
Configuration
Setting
Default
Description
MILEPOST_NUDGE_INTERVAL
900 (15 min)
Minimum seconds between nudges, per project. Raise it to journal less often; lower it to journal more.
.no-milepost file
—
Drop an empty .no-milepost file at a project's root to exclude that project entirely: no diary, no nudges, no recall. For client work or anything you'd rather keep out of the diary.
MILEPOST_STRICT
0
Set to 1 to make the secret guard also block near-secrets : IP addresses and e-mail addresses. Opt-in, because many legitimate diaries reference internal hosts. Set it where Claude Code's hooks see it — e.g. in settings.json under "env": {"MILEPOST_STRICT": "1"} or in your shell profile.
Prefer zero hooks? Skip the nudge entirely — the CLAUDE.md policy works on its own; it just leans on the model to remember during very long sessions.
milepost
claude-mem
claude-diary
Storage
Plain markdown
SQLite + vector DB
Plain markdown
Greppable / exportable
✅
⚠️ via export
✅
Writes during a session
✅
✅
❌
Milestone-aware (high-signal)
✅
⚠️ captures everything
⚠️ session-level
Trigger
Milestones
Continuous
Manual / pre-compact
Living status file
✅
❌
❌
Database required
No
Yes
No
Honest caveat: milestone detection is a judgment , so it's best-effort — it may occasionally over- or under-log, and entries use a little of your working session's tokens. /milepost is the manual backstop.
What is the best way to give Claude Code long-term memory?
Use a plugin that persists context to durable, portable storage. milepost does this with plain-markdown files written as you work, so your memory is readable, greppable, and never locked inside a database.
Does Claude Code remember things between sessions?
Not on its own. Claude Code is stateless across sessions beyond CLAUDE.md . milepost adds true cross-session memory by journaling milestones and a live status file that Claude reads back later.
How is milepost different from claude-mem?
claude-mem captures everything into a SQLite + vector database. milepost stores high-signal milestones as plain markdown — easier to read, version, export, and trust, with nothing to lock into.
Does milepost use a database or send my data anywhere?
No. Everything stays in local markdown files under ~/.claude/memory/milepost/ . No database, no cloud, no telemetry.
Will it log secrets or sensitive code?
Defense in depth, not just a polite instruction:
Policy — Claude is told to reference secrets by location only ("token stored in .ha_token "), never by value, and to minimize near-secrets (internal IPs, serials, third-party personal data).
Guard hook — a PreToolUse hook scans every write targeting the diary directory (including shell redirection via Bash ) and denies anything matching secret patterns: private-key blocks, AWS/GitHub/Slack/Google/ sk- -style keys, JWTs, bearer tokens, password= / api_key: -style assignments, URLs with embedded credentials, plus a high-entropy catch-all for unknown token formats (≥32-char base64url-ish runs mixing upper/lower/digits — tuned so git SHAs, UUIDs, paths, and identifiers don't trip it).
Strict mode — MILEPOST_STRICT=1 additionally blocks IP addresses and e-mail addresses, for client work and privacy-sensitive setups.
Audit — you can scan existing diaries anytime: bash ~/.claude/hooks/milepost-secret-scan.sh < ~/.claude/memory/milepost/<project>/log.md (exit 1 + pattern names on a hit).
Opt-out — a .no-milepost file at a project's root keeps that project out of the diary entirely.
How is the diary protected on disk?
Diaries are plaintext markdown, owned by you. The installer restricts ~/.claude/memory/milepost/ to owner-only access ( chmod 700 ). There is deliberately no at-rest encryption : milepost's whole design is that Claude reads and writes the diary with its ordinary file tools, and encrypting it would break that (and your greppability) while adding key management — use full-disk encryption (FileVault/LUKS/BitLocker), which covers the stolen-device case properly. Treat the diary dir like your shell history: local, private, never committed or synced to shared storage.
Can a malicious repo trick Claude into leaking my diary?
The theoretical risk with any memory system: instructions hidden in a repo's files could ask the agent to copy remembered content somewhere it doesn't belong (prompt injection). milepost's policy explicitly forbids copying diary contents into project files or external services, and the diary only ever contains milestone summaries — not credentials (see the guard above). For projects where even summaries are sensitive, use .no-milepost so there is nothing to leak. Review what agents write, as always.
Do I have to approve every diary write?
No. The installer adds two narrowly-scoped permission allow-rules to settings.json — Edit(~/.claude/memory/milepost/**) and Write(~/.claude/memory/milepost/**) — so Claude can journal without prompti

[truncated]
