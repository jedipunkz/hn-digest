---
source: "https://github.com/fronalabs/frona/releases/tag/v2026.5.5"
hn_url: "https://news.ycombinator.com/item?id=48340093"
title: "Frona v2026.5.5 – self-hosted personal AI assistant"
article_title: "Release v2026.5.5 · fronalabs/frona · GitHub"
author: "syncerx"
captured_at: "2026-05-31T00:25:49Z"
capture_tool: "hn-digest"
hn_id: 48340093
score: 2
comments: 0
posted_at: "2026-05-30T20:03:54Z"
tags:
  - hacker-news
  - translated
---

# Frona v2026.5.5 – self-hosted personal AI assistant

- HN: [48340093](https://news.ycombinator.com/item?id=48340093)
- Source: [github.com](https://github.com/fronalabs/frona/releases/tag/v2026.5.5)
- Score: 2
- Comments: 0
- Posted: 2026-05-30T20:03:54Z

## Translation

タイトル: Frona v2026.5.5 – 自己ホスト型パーソナル AI アシスタント
記事のタイトル: リリース v2026.5.5 · fronalabs/frona · GitHub
説明: Frona はパーソナル AI アシスタントです。自律エージェントを作成し、ツールを提供し、チャット インターフェイスを通じて対話します。エージェントは独自に行動します。彼らは Web を閲覧し、コードを実行し、アプリケーションを開発し、インターネットを検索し、電話をかけ、お互いに仕事を委任します。あなたは彼らにタスクを与えて、
[切り捨てられた]

記事本文:
リリース v2026.5.5 · fronalabs/frona · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
フロナラボ
/
フロナ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
フィルター
読み込み中
ごめんなさい、何か

間違ってしまいました。
読み込み中にエラーが発生しました。このページをリロードしてください。
v2026.5.5
62b233e
このリリースでは、ID の統合とタスク スケジュールの堅牢性が大幅に強化されています。ユーザーが所有するすべてのエンティティ (ユーザー、エージェント、アプリ、McpServer、チャネル) はハンドルの newtype を持ち、ポリシー プリンシパルはユーザーの親の下で {user_handle}/{entity_handle} として再形成され、すべてのストレージ パスとルーティング パスはハンドルによってキー設定されます。タスクは、cron およびワンショットの結果をソース チャットにレンダリングする方法を制御する必要な result_schema を取得します。これには、長年にわたる CronRun のクラッシュ回復レースの修正が含まれます。チャット UI は、スクロールアップ時に履歴のページ分割を行い、メッセージごとのホバー時間とともに日の区切り文字を表示するようになりました。内部的には、ポリシーの決定とエージェントごとに許可されたツール セットがキャッシュされるようになり、サンドボックスの制限により部分的な保存が受け入れられ、ハンドルの移行により既存のインストールが引き継がれます。
サポートされているチャネル: Slack、Discord、Telegram、Signal、SMS、WhatsApp Cloud、WhatsApp Personal。
サニタイズルールとビルドスクリプト生成の予約名リストを備えたハンドル newtype を追加します。ハンドル用に User.username の名前を変更し、 Agent.user_id + Agent.handle を必要とします。
App.handle 、 Channel.handle を追加し、 McpServer.slug の名前を handle に変更します。 4 つすべてをハンドルキー付きパスに切り替えます。
data/users/{handle}/{subsystem} の下のディスク上のストレージをチャネルごとのストレージ パスと統合します。
Cedar エンティティ UID を親ユーザーの Kind::"{user_handle}/{entity_handle}" に切り替えます。 SandboxPrincipalRef を構造体 + Kind 列挙型にリファクタリングし、それに応じてpolicy::serviceを再配線します。
Cedar スキーマ、db init インデックスを更新し、handle_unification 移行を追加します (2026-05-24T00:00:00Z でのカットオーバー)。 seed_config_agents を削除します。
CandidateEvent を再構成して、型指定されたユーザー / チャネル / チャット / メッセージ / 連絡先の参照を保持します。
認証ミドルウェアとすべての API ルートを通じてユーザー ハンドルを伝えます。

管理/認証、エージェント、アプリ プロキシ、ファイル、MCP、ブラウザ、スキル、ツール、メッセージ。
user_handle/agent_handle によるエージェントのスキルの範囲。コンテンツ + サンドボックス ツール、ツール マネージャー、およびハンドル パスを使用するエージェントに影響するツールを更新します。
チャット、チャネル、資格情報、通知、メモリ、および推論サービスを介したスレッド ハンドル。
組み込みエージェント プロンプトと、ハンドル キー付き URL の manage_service プロンプト + マニフェスト スキーマを更新します。
タスク結果配信用に result_schema 分類子とレンダラーを追加します。 create_task / create_quirting_task およびスキーマドライブのソースチャット配信に result_schema が必要です。
create_cron_template + spawn_cron_run を介して result_schema を構築します。 complete_task.result を任意の JSON 値として受け入れ、スキーマに対して検証します。
send_message をエージェントのハートビート チャットに制限し、タスクが complete_task.result を通じて配信されるようにします。
孤立したクエリを絞り込み、resume_all の順序を変更することで、CronRun のクラッシュ回復競合を修正しました。
CronRun が適切なスペースに配信されるように、cron テンプレートを通じて space_id をスレッドします。
オーサリング エージェントに関係なく、最初のタスクの指示を <task> で囲みます。
タスク実行チャットをナビゲーション スペース ビューから非表示にします。空のタスク完了履歴行を警告からデバッグに降格します。
未使用の SCHEDULED_TASK.md プロンプトを削除します。研究者の責任として、繰り返しのダイジェストと概要を宣伝します。
古いチャット メッセージを上にスクロールして読み込み、チャットを読み込むときに履歴メッセージを SSE バッファーされたメッセージとマージします。
カーソルを合わせると、日の区切り文字、会話のギャップ、メッセージごとの時間が表示されます。
認証、ナビゲーション、セッションのコンテキストを分割します。 AuthGuard を AppGate + RequireAuth に置き換え、認証/セットアップ/メイン レイアウトを新しいゲート モデルに再配線します。
アプリ、通知、設定をハンドルによってルーティングします。メイン ページと会話パネルを新しいコンテキスト API に切り替えます。更新タイプと

ハンドルベースの応答用の pi-client。
URL がすでに私たちのものである場合は setWebhook をスキップし、古い URL を置き換える場合はキューをドロップします。
未知のチャネルへの受信 Webhook が到着したときに警告します。
Channel.id を SurrealDB レコード リテラルではなく、裸の UUID として保存します。
生成されたタスクがキャンセルされ、on_disconnect が実行されることをアサートするテストで、channel-delete をカバーします。
プレフィックススコープの無効化を備えたポリシー決定キャッシュを追加します。
/api/agents を O(N·M²) から O(N·M) にカットするために、エージェントごとに許可されたツール セットをキャッシュします。
EWMA を介して CPU をスムーズに追跡し、一時的なバーストを吸収します。
部分保存時に不足しているサンドボックス制限をサーバーのデフォルトで埋めます。
cache_key を user_id に合わせて、sandbox_cache の無効化を修正しました。
Prod イメージを個人プレビュー レジストリに公開するための Docker タスクを追加します。 common.sh を介して定数/ヘルパーを共有し、GHCR のイメージ インデックスに注釈を付けます。
開発コンテナの再起動後もカーゴ レジストリと git キャッシュを保持します。
update-versions の crates.io リクエストで User-Agent ヘッダーを送信します。 smbclient と uv パッケージのピンをバンプします。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Frona is a personal AI assistant. You create autonomous agents, give them tools, and talk to them through a chat interface. Agents act on their own. They browse the web, run code, develop applications, search the internet, make phone calls, and delegate work to each other. You give them a task and t
[truncated]

Release v2026.5.5 · fronalabs/frona · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
fronalabs
/
frona
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Filter
Loading
Sorry, something went wrong.
There was an error while loading. Please reload this page .
v2026.5.5
62b233e
This release is a big push on identity unification and task scheduling robustness. Every user-owned entity (User, Agent, App, McpServer, Channel) now carries a Handle newtype, Policy principals are reshaped as {user_handle}/{entity_handle} under a User parent, and all storage and routing paths are keyed by handle. Tasks gain a required result_schema that drives how cron and one-shot results render back into source chats, with a fix for a long-standing CronRun crash-recovery race. The chat UI now paginates history on scroll-up and shows day separators with per-message hover times. Under the hood, Policy decisions and per-agent permitted tool sets are now cached, sandbox limits accept partial saves, and a Handle migration carries existing installs forward.
Supported channels: Slack, Discord, Telegram, Signal, SMS, WhatsApp Cloud, WhatsApp Personal.
Add a Handle newtype with sanitization rules and a build-script-generated reserved-name list; rename User.username to handle and require Agent.user_id + Agent.handle .
Add App.handle , Channel.handle , and rename McpServer.slug to handle ; switch all four to handle-keyed paths.
Unify on-disk storage under data/users/{handle}/{subsystem} with per-channel storage paths.
Switch Cedar entity UIDs to Kind::"{user_handle}/{entity_handle}" with User parent; refactor SandboxPrincipalRef into a struct + Kind enum and rewire policy::service accordingly.
Update the Cedar schema, db init indexes, and add a handle_unification migration (cutover at 2026-05-24T00:00:00Z ); drop seed_config_agents .
Reshape CandidateEvent to carry typed User / Channel / Chat / Message / Contact references.
Carry user handle through the auth middleware and all API routes: admin/auth, agents, app proxy, files, MCP, browser, skills, tools, messages.
Scope agent skills by user_handle/agent_handle ; update content + sandbox tools, the tool manager, and agent-affecting tools to use Handle paths.
Thread Handle through chat, channel, credential, notification, memory, and inference services.
Refresh built-in agent prompts and the manage_service prompt + manifest schema for handle-keyed URLs.
Add a result_schema classifier and renderer for task result delivery; require result_schema on create_task / create_recurring_task and schema-drive source-chat delivery.
Plumb result_schema through create_cron_template + spawn_cron_run ; accept complete_task.result as any JSON value and validate against the schema.
Restrict send_message to the agent's heartbeat chat so tasks deliver through complete_task.result .
Fix CronRun crash-recovery race by narrowing the orphan query and reordering resume_all .
Thread space_id through cron templates so CronRuns deliver into the right space.
Wrap the initial task instruction in <task> regardless of authoring agent.
Hide task-execution chats from the navigation space view; demote empty task-completion history rows from warn to debug.
Remove the unused SCHEDULED_TASK.md prompt; advertise recurring digests and briefs as researcher responsibilities.
Load older chat messages on scroll-to-top and merge historical messages with SSE-buffered ones when loading a chat.
Show day separators, conversation gaps, and per-message time on hover.
Split auth, navigation, and session contexts; replace AuthGuard with AppGate + RequireAuth and rewire auth/setup/main layout for the new gating model.
Route apps, notifications, and settings by handle; switch main pages and the conversation panel to the new context APIs; update types and the api-client for handle-based responses.
Skip setWebhook when the URL is already ours and drop the queue when replacing a stale URL.
Warn when an inbound webhook arrives for an unknown channel.
Store Channel.id as a bare UUID instead of a SurrealDB record literal.
Cover channel-delete with a test that asserts the spawned task is cancelled and on_disconnect runs.
Add a policy decision cache with prefix-scoped invalidation.
Cache per-agent permitted tool sets to cut /api/agents from O(N·M²) to O(N·M).
Smooth tracked CPU via EWMA to absorb transient bursts.
Fill missing sandbox limits with server defaults on partial save.
Fix sandbox_cache invalidation by aligning cache_key to user_id .
Add a docker task for publishing prod images to a personal preview registry; share constants/helpers via common.sh and annotate the image index for GHCR.
Persist cargo registry and git caches across dev container restarts.
Send a User-Agent header on crates.io requests in update-versions ; bump smbclient and uv package pins.
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
