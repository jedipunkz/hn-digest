---
source: "https://github.com/Incultnitollc/claude-code-live-bridge"
hn_url: "https://news.ycombinator.com/item?id=48460471"
title: "Show HN: Cc-bridge – Two Claude Code windows talking via one JSONL file"
article_title: "GitHub - Incultnitollc/claude-code-live-bridge: Live JSONL message bridge between local Claude Code sessions · GitHub"
author: "PengSpirit"
captured_at: "2026-06-09T13:07:00Z"
capture_tool: "hn-digest"
hn_id: 48460471
score: 2
comments: 0
posted_at: "2026-06-09T12:52:03Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Cc-bridge – Two Claude Code windows talking via one JSONL file

- HN: [48460471](https://news.ycombinator.com/item?id=48460471)
- Source: [github.com](https://github.com/Incultnitollc/claude-code-live-bridge)
- Score: 2
- Comments: 0
- Posted: 2026-06-09T12:52:03Z

## Translation

タイトル: Show HN: Cc-bridge – 1 つの JSONL ファイルを介して対話する 2 つのクロード コード ウィンドウ
記事のタイトル: GitHub - Incultnitollc/claude-code-live-bridge: ローカル クロード コード セッション間のライブ JSONL メッセージ ブリッジ · GitHub
説明: ローカルのクロード コード セッション間のライブ JSONL メッセージ ブリッジ - Inultnitollc/claude-code-live-bridge

記事本文:
GitHub - Incultnitollc/claude-code-live-bridge: ローカル クロード コード セッション間のライブ JSONL メッセージ ブリッジ · GitHub
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
カルトニトール
/
クロードコードライブブリッジ
公共
通知
通知設定を変更するにはサインインする必要があります
追加

最後のナビゲーションオプション
コード
Incultnitollc/claude-code-live-bridge
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
19 コミット 19 コミット .github/ workflows .github/ workflows docs/ superpowers docs/ superpowers src src テスト テスト .gitignore .gitignore .npmignore .npmignore .prettierrc.json .prettierrc.json ライセンス ライセンス README.md README.md eslint.config.js eslint.config.js package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ローカルのクロード コード セッション間のライブ JSONL メッセージ ブリッジ。
同じマシン上の 2 つ (またはそれ以上) の Claude Code ウィンドウ。彼らはリアルタイムで会話する必要があります。 cc-bridge は、各セッションに共有ファイルにメッセージを送信し、Claude の Monitor ツールがライブ プッシュ通知として消費するファイルテール ストリーム経由で新しいメッセージをリッスンする CLI を提供します。デーモンもネットワークもポーリングもありません。
npm install -g @incultnitollc/cc-bridge
ノード 20 以上が必要です。
cc-bridge は「B からこんにちは」を送信します
ウィンドウ A はすぐに JSONL 行を受け取ります。
各セッションで、監視ツールの下で listen コマンドを実行すると、追加されたすべての行がライブ プッシュ通知 (ポーリングなし) として到着します。
モニター: cc-bridge listen デフォルト
次に、別のウィンドウから Bash ツールを介して送信します。
Bash: cc-bridge が「レビュー準備完了」を送信
コンセプト
Room — ~/.cc-bridge/rooms/<room>.jsonl にある名前付き JSONL ファイル。デフォルトの部屋はデフォルトです。使用できる名前: [a-zA-Z0-9_.-]{1,64} .
セッション ID — 自動生成された <host>-<ppid>-<rand8> (親シェル PID であるため、2 つのウィンドウは当然異なります)。 CC_BRIDGE_FROM=... でオーバーライドします。
メッセージ — 単一行の JSON。必須: v、id、ts、room、from、msg 。オプション: to、reply_to、kind 。前方互換のために保存された不明なフィールド。

cc-bridge listen [room] [--replay N] [--pretty] [--filter from = X] [--from ID] [--json-errors]
cc-bridge send < msg > [--room R] [--to ID] [--reply-to ULID] [--kind text |イベント] [--ID から]
エコー「こんにちは」 | cc-bridge send # パイプ接続されている場合は標準入力からの読み取り
cc-bridge rooms # サイズ + mtime でルームをリストする
cc-bridge ルーム クリア < ルーム > --はい
cc-bridge validate < file > # JSONL ルーム ファイルを lint する
cc-bridge --バージョン
cc-bridge --ヘルプ
図書館の利用
import { sendMessage , listen , listRooms } from '@incultnitollc/cc-bridge'
await sendMessage ( { from : 'planner' 、 room : 'team' 、 msg : 'kicking off build' } )
const ctrl = listen ( { room : 'team' , sessionId : 'reviewer' } )
for await ( const ev of ctrl . iterator ) {
if (ev.ok) コンソール。ログ(ev.行)
}
セキュリティモデル
cc-bridge はローカルホスト IPC プリミティブです。信頼境界はユーザー アカウントです。 ~/.cc-bridge/ を共有ファイルシステム、ネットワークドライブ、またはクラウド同期ディレクトリに配置しないでください。
~/.cc-bridge/ 内のファイルは、モード 0700 (ディレクトリ) / 0600 (ファイル) で作成されます。
部屋名は消毒されました。パスの横断が拒否されました。
ルームパスのシンボリックリンクが拒否されました。
メッセージあたり 64KB の上限。 10MB ルーム ファイルのソフト警告。 100MBのハードリフューズ。
--pretty は、端末ハイジャックを防ぐために ANSI エスケープ シーケンスを削除します。
from フィールドは送信者によってアサートされます (署名なし)。 v1 は、 $HOME への書き込みアクセス権を持つすべてのユーザーを明示的に信頼します。
v1.1 — cc-bridge インストールフック (Claude Code Stop/UserPromptSubmit フックの自動接続)、DM フィルタリング ( --me )、時間ベースの再生 ( --replay 1h )、開封確認、Windows サポート。
v2 — MCP サーバー ラッパー、クロスマシン バックエンド (Supabase Realtime / Redis)、Webhook ファンアウト、オブザーバー ダッシュボード。
ローカルのクロード コード セッション間のライブ JSONL メッセージ ブリッジ
github.com/Incultnitollc/claude-code-live-bridge
リソース
読み込み中にエラーが発生しました。リロードしてください

このページ。
アクティビティ
カスタムプロパティ
星
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

Live JSONL message bridge between local Claude Code sessions - Incultnitollc/claude-code-live-bridge

GitHub - Incultnitollc/claude-code-live-bridge: Live JSONL message bridge between local Claude Code sessions · GitHub
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
Incultnitollc
/
claude-code-live-bridge
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Incultnitollc/claude-code-live-bridge
main Branches Tags Go to file Code Open more actions menu Folders and files
19 Commits 19 Commits .github/ workflows .github/ workflows docs/ superpowers docs/ superpowers src src tests tests .gitignore .gitignore .npmignore .npmignore .prettierrc.json .prettierrc.json LICENSE LICENSE README.md README.md eslint.config.js eslint.config.js package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts vitest.config.ts vitest.config.ts View all files Repository files navigation
Live JSONL message bridge between local Claude Code sessions.
Two (or more) Claude Code windows on the same machine. They need to talk in real time. cc-bridge gives each session a CLI to send a message to a shared file and listen for new messages via a file-tail stream that Claude's Monitor tool consumes as live push notifications. No daemon, no network, no polling.
npm install -g @incultnitollc/cc-bridge
Requires Node ≥ 20.
cc-bridge send " hello from B "
Window A immediately receives the JSONL line.
In each session, run the listen command under the Monitor tool so every appended line arrives as a live push notification — no polling.
Monitor: cc-bridge listen default
Then send from the other window via the Bash tool:
Bash: cc-bridge send "ready for review"
Concepts
Room — a named JSONL file under ~/.cc-bridge/rooms/<room>.jsonl . Default room is default . Names allowed: [a-zA-Z0-9_.-]{1,64} .
Session id — auto-generated <host>-<ppid>-<rand8> (parent shell PID, so two windows differ naturally). Override with CC_BRIDGE_FROM=... .
Message — single-line JSON. Required: v, id, ts, room, from, msg . Optional: to, reply_to, kind . Unknown fields preserved for forward-compat.
cc-bridge listen [room] [--replay N] [--pretty] [--filter from = X] [--from ID] [--json-errors]
cc-bridge send < msg > [--room R] [--to ID] [--reply-to ULID] [--kind text | event] [--from ID]
echo " hi " | cc-bridge send # reads from stdin if piped
cc-bridge rooms # list rooms with size + mtime
cc-bridge rooms clear < room > --yes
cc-bridge validate < file > # lint a JSONL room file
cc-bridge --version
cc-bridge --help
Library use
import { sendMessage , listen , listRooms } from '@incultnitollc/cc-bridge'
await sendMessage ( { from : 'planner' , room : 'team' , msg : 'kicking off build' } )
const ctrl = listen ( { room : 'team' , sessionId : 'reviewer' } )
for await ( const ev of ctrl . iterator ) {
if ( ev . ok ) console . log ( ev . line )
}
Security model
cc-bridge is a local-host IPC primitive . The trust boundary is your user account. Do not place ~/.cc-bridge/ on a shared filesystem, network drive, or cloud-synced directory.
Files in ~/.cc-bridge/ are created with mode 0700 (dir) / 0600 (files).
Room names sanitized; path traversal refused.
Symlinks at room paths refused.
64KB per-message cap. 10MB room file soft-warn; 100MB hard-refuse.
--pretty strips ANSI escape sequences to prevent terminal hijack.
The from field is sender-asserted (no signature). v1 explicitly trusts everyone with write access to your $HOME .
v1.1 — cc-bridge install-hooks (auto-wire Claude Code Stop/UserPromptSubmit hooks), DM filtering ( --me ), time-based replay ( --replay 1h ), read receipts, Windows support.
v2 — MCP server wrapper, cross-machine backend (Supabase Realtime / Redis), webhook fanout, observer dashboard.
Live JSONL message bridge between local Claude Code sessions
github.com/Incultnitollc/claude-code-live-bridge
Resources
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
