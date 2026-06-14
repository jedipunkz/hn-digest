---
source: "https://github.com/fronalabs/frona/releases/tag/v2026.6.0"
hn_url: "https://news.ycombinator.com/item?id=48531500"
title: "Frona v2026.6.0 – self-hosted personal AI assistant"
article_title: "Release v2026.6.0 · fronalabs/frona · GitHub"
author: "syncerx"
captured_at: "2026-06-14T21:33:27Z"
capture_tool: "hn-digest"
hn_id: 48531500
score: 1
comments: 0
posted_at: "2026-06-14T19:17:56Z"
tags:
  - hacker-news
  - translated
---

# Frona v2026.6.0 – self-hosted personal AI assistant

- HN: [48531500](https://news.ycombinator.com/item?id=48531500)
- Source: [github.com](https://github.com/fronalabs/frona/releases/tag/v2026.6.0)
- Score: 1
- Comments: 0
- Posted: 2026-06-14T19:17:56Z

## Translation

タイトル: Frona v2026.6.0 – 自己ホスト型パーソナル AI アシスタント
記事のタイトル: リリース v2026.6.0 · fronalabs/frona · GitHub
説明: Frona はパーソナル AI アシスタントです。自律エージェントを作成し、ツールを提供し、チャット インターフェイスを通じて対話します。エージェントは独自に行動します。彼らは Web を閲覧し、コードを実行し、アプリケーションを開発し、インターネットを検索し、電話をかけ、お互いに仕事を委任します。あなたは彼らにタスクを与えて、
[切り捨てられた]

記事本文:
リリース v2026.6.0 · fronalabs/frona · GitHub
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
申し訳ありませんが、問題が発生しました。
エラーが発生しました

ルをロードしています。このページをリロードしてください。
v2026.6.0
5cf63f0
このリリースには、すべてのチャネルにわたる統合された人間参加型の一時停止/再開メカニズム、スキルや他のエージェントを直接呼び出すためのスラッシュ コンポーザー、トークンの使用と幻覚をカットする型付きファイル ツール、および長い応答が切り捨てられることなくきれいに届くチャネルごとのメッセージ スプリッターが搭載されています。承認 (アプリのデプロイ、多肢選択の質問、資格情報の選択) は、Telegram、Discord、Slack、WhatsApp Cloud ではボタンとしてレンダリングされ、Signal、SMS、および個人用 WhatsApp では YES または NO の応答プロンプトとしてレンダリングされ、受信パーサーは明らかなバリアントを受け入れます。バックグラウンドでは、新しい Harness 構造体が AppState のエージェント ランタイム ビューを統合し、サンドボックス化がプリンシパルの種類ごとに 1 つの SandboxManager に移行し、チャネル アダプター フレームワークが障害を型指定された ChannelError で分類するようになりました。
サポートされているチャネル: Slack、Discord、Telegram、Signal、SMS、WhatsApp Cloud、WhatsApp Personal。
ツール呼び出しの不透明な tool_data BLOB を、型指定された Hitl 値に置き換えます。型指定された HITL ツール コンポーネントを構築し、新しい SSE イベント名を使用します。
一時停止されたメッセージ ステータスを維持し、アトミック コンペア アンド スワップで再開します。新しい InferenceResponse の結果を一時停止/再開の永続性に変換します。
POST /api/chats/{chat_id}/tool-calls/resolve エンドポイントを追加し、HITL コールバック パーサーと応答ラベルを共有 hitl モジュールに取り込み、従来のアプリとボールトの承認ルートを削除して、解決エンドポイントに置き換えます。
inference_done の msg.tool_calls から保留中の HITL を再シードして、一時停止したターンで新しい質問が引き続き表面化するようにします。
YES/NO バリアント解析 ( y 、 Yea 、 ok 、 nope 、親指を立てた絵文字、および通常のバリアント) を備えた受信応答 HITL リゾルバーを共有します。
解決時に HITL 配信を再開し、順次アダプターが次のプロンプトを表示するようにします。を使用してください

ピッカーをレンダリングできないチャネル上の HITL フォールバック URL の絶対 public_base_url。
Telegram (インライン キーボード)、Slack (ソケット モード インタラクションによるブロック キット)、Discord (ボタン)、WhatsApp Cloud (インタラクティブ メッセージとボタン タップ)、WhatsApp 個人アカウント (引用返信)、Signal (引用返信)、および SMS (アプリ展開の承認プロンプトで YES または NO のヒントを返信) で HITL プロンプトをレンダリングし、応答を解決します。
長い推論の間、Discord と WhatsApp のタイピング インジケーターを押し続けると、エージェントが生成している間のアクティビティがユーザーに表示されます。
manage_service ツールの名前を manage_app に変更し、型指定された HITL フックを採用します。 HITL でタスクを終了し、 run_task 経由で再生成します。
HITL 永続性をテストし、レンダリングを一時停止し、バッチ解決し、レースを再開します。
/ および @ プレフィックス用の chat::slash 呼び出しパーサーを追加します。解析されたスラッシュ呼び出し用に、Message に MessageCommand サイド フィールドを追加します。
/commands をエンドツーエンドで接続します: ディスパッチ、ターミナル書き込みリファクタリング、クロスエージェント アトリビューション、スラッシュ解析。
GET /api/chats/{id}/commands 検出エンドポイントと listCommands API クライアントを追加します。
フロントエンド: レキシカル トリガー、ディレクティブ チップ、および現在のエージェントとの新しいチャットを開く /new ビルトインを備えたスラッシュ コンポーザー。
@<エージェント> ターンごとのオーバーライド: 返信はターゲット エージェントに帰属され、次のメッセージはチャットのデフォルトに戻ります。
save_chat 、 save_updated_message 、および delete_messages_for_chat ヘルパーを追加します。
スキル: 新しい SKILL.md フロントマター フィールド disable-model-invocation 、 argument-hint 、および argument 。 disable-model-invocation スキルをモデルの available-skills ブロックから非表示にし、ユーザーが直接呼び出せるように / メニューに保持します。
/agent ディスパッチとコマンド検出のための e2e テストを追加します。
段落、行、次に分割する共有境界プリミティブを備えた形式ごとのメッセージ スプリッターを追加します。

単語、次に UTF-8 文字を使用し、コード フェンスやエスケープ シーケンスの内側には決して入れないでください。
Telegram にはネイティブ テーブル サポートがないため、TelegramMarkdownV2Splitter を Telegram アダプタ (メッセージあたり 4,096 文字の制限) に接続し、マークダウン テーブルを等幅コード ブロックとしてレンダリングします。
MarkdownSplitter を Discord アダプターに接続し (メッセージあたり 2,000 文字制限)、古い chunk_for_discord ヘルパーを削除し、マークダウン テーブルを等幅コード ブロックとしてレンダリングします。
PlainSplitter を Slack アダプター (サイレント マルチチャンク) に接続します。
PlainSplitter を 1,600 文字のハード キャップと完全な応答: Chat 種類の共有 URL を指す {short-link} オーバーフロー (同じチャット内のオーバーフロー イベント全体で再利用) を使用して SMS アダプターに接続します。
MarkdownSplitter を WhatsApp Cloud アダプターと WhatsApp Personal アダプターに接続します。
SignalSplitter をチャンクごとの SignalText { body, ranges } を使用して Signal アダプターに接続します。
ChannelError タイプと Messagedelivery.failure_kind フィールドを追加します。オプションの retry_hint を使用して、チャネル アダプターを型指定された ChannelError 分類 (Transient、Forbidden、NotFound、PayloadInvalid、PayloadTooLarge、Unauthorized、Other) に移行します。
start_with_retry を冪等にすることで、1 回の起動でアダプターが二重に生成されなくなります。
境界のあるブロードキャスト バスを境界のないファンアウトに置き換えて、チャネル インバウンドのイベントの損失を防ぎます。
ファイルとチャットの種類を持つ共有エンティティと、短いリンクを発行するための ShareService を追加します。サービスを AppState に接続し、定期的なクリーンアップをスケジュールします。
share.ttl_secs (デフォルトは 30 日) および share.cleanup_interval_secs (デフォルトは 6 時間) 構成ノブを追加します。
正規 URL または署名付き URL への 303 リダイレクトを使用して GET /s/{id} ショートリンク リゾルバーを追加します。
マークダウンとコードを構文強調表示でレンダリングする /p プレビュー ページを追加します。 GET /p/{slug} リダイレクターをクエリパラメータフォームに追加します。
ルートチャネルアタッチメントt

新しい chat::channel::attachment ヘルパーを介してすべてのアダプターで outbound_url を介して: インライン チャネルのプレビュー可能なファイルは /p/{8-char-id} を取得し、インラインのプレビュー不可能なファイルは /s/{8-char-id} を取得し、ボタン チャネルは長い形式の /p/{owner}/{handle}/{path} または /api/files/... URL を出力します。
anonymous_not_found ヘルパーを追加して、不明 / 期限切れ / 無許可の共有 ID がバイト同一の 404 を返すようにします。
FilePreviewContent を共有モジュールに抽出します。非 JSON 認証リクエストの request<T> から apiFetch を抽出します。
SMS チャット共有の再利用のために、遅延 lookup_or_issue_chat を備えた ShareKind::Chat バリアントを追加します。
create_task および create_quirting_task で result_description を受け入れると、散文的な回答でスキーマの作成がスキップされます。 result_description と result_schema の両方が渡される (またはどちらも渡されない) 場合は拒否されます。
複雑なタスク結果スキーマのトップレベルの概要: 文字列が必要です。タスク作成時にそれがないと複雑なスキームを拒否します。
TaskCompletion イベントで result_schema を実行し、JSON コンテンツを永続化します。
スキーマ駆動の本文とタスク チャットへのリンクを含む TaskCompletion バブルをレンダリングします。チャネル アダプター間で共有マークダウン ヘルパーを介して TaskCompletion をレンダリングします。
研究者は、完全なリサーチを、トピックにちなんで名付けられたマークダウン添付ファイルとして公開するようになりました (例: h100-used-prices.md )。これにより、連続するリサーチ タスクが互いに上書きされなくなります。
入力された read 、 write 、 edit 、 glob 、および grep ツールを追加します。デフォルトではワークスペースをスコープとしています。絶対パスとポリシーで承認された兄弟パスが許可されます。仮想パス URI が拒否されました。
edit は Unicode 正規化 (NFKC、ASCII 引用符/ダッシュ/スペース、折りたたまれた空白) を使用するため、わずかに記憶違いのスニペットでもターゲットと一致します。動機はトークンの使用と幻覚です。型付き編集は、モデルのフリー スタイル シェル エスキャップなしで、シェルのラウンドトリップが消費するトークンの一部で構造化された差分を返します。

半分覚えていたパスに対してです。
共有テキストと検索プリミティブを含む frona-text クレートを追加します ( NormalizedString 、 LineEnding 、 .gitignore 対応トラバーサル用の walk_with_ignore )。
サンドボックスのリファクタリング: SandboxManager は、プリンシパルの種類ごとに単一のエントリ ポイントになりました。
ツールごとのビュー レジストリを抽出します: ToolRow プリミティブ、 DefaultView 、スリム クロム。
シェル、Python、およびノー​​ドのビューを追加します (サンドボックス拒否抽出を使用)。型指定されたファイル ツール ( read 、 write 、 edit 、 glob 、 grep ) では、構文の強調表示と行番号が表示されます。 process_file の場合 (ファイル名 / サイズ / コンテンツ タイプ カード); web_search (結果リスト) と web_fetch (マークダウン + クリック可能な URL) の場合。タスクツール ( create_task 、 cronstrue を使用した create_quirting_task 、 delete_task );記憶ツール用 (大きなテキストでは自動拡張)。 update_identity の場合 (属性レンダリングの設定/削除);そして set_heartbeat (人間化された間隔と次のティック時間) の場合。
シェル ツール コマンドをラップ付きの bash で強調表示されたブロックとしてレンダリングします。 process_file をサイズとコンテンツ タイプを含むファイル名カードとしてレンダリングします。
vitest テスト ランナーと cronstrue / unbash の依存関係を追加します。
Harness 構造体を AppState のエージェント ランタイム ビューとして追加します。セッションの構築、音声 WS、メッセージ ストリーム、および再開のすべてをルーティングします。
Arc<Harness> を保持し、独自の解決通知マップを所有するための Narrow TaskExecutor。
HITL 解像度を Harness に移動し、 ChannelCtx.app_state をドロップします。
タスクツールの登録を ToolManager に移動します。 Delivery_event_to_source を抽出するため、 ReportSignalTool は TaskExecutor を必要としなくなります。
Agent/execution.rs を削除し、Harness のテスト フィクスチャを更新します。
一時的なターンとしてハートビートを挿入すると、エージェントはユーザー メッセージと同様にハートビートへの応答を停止します。
ターンごとのライフサイクル イベントを InferenceEventKind に折りたたみます。
Browserless の 408 秒のハード タイムアウトを過ぎてもブラウザ セッションを維持します: pass

接続時の長いタイムアウト、60 秒のマージンでセッションをセルフエビクトしてリサイクルし、死んだ WebSocket を即座にエビクトします。
チャットがページ制限を超えた場合に最新のメッセージをドロップしないようにします。
推論の切り替えをインライン コンテンツからメッセージ ヘッダーの輝くアイコンに移動します。
アシスタント マークダウンで単一チルダ取り消し線を無効にします。
クイックスタート: Docker Compose を使用してスタックを数分で起動します
スクリーンショット: プラットフォームの動作を確認します
ドキュメント : 完全なガイドとリファレンス
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Frona is a personal AI assistant. You create autonomous agents, give them tools, and talk to them through a chat interface. Agents act on their own. They browse the web, run code, develop applications, search the internet, make phone calls, and delegate work to each other. You give them a task and t
[truncated]

Release v2026.6.0 · fronalabs/frona · GitHub
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
v2026.6.0
5cf63f0
This release ships a unified human-in-the-loop pause/resume mechanism across every channel, a slash composer for invoking skills and other agents directly, typed file tools that cut token usage and hallucinations, and per-channel message splitters so long replies arrive cleanly instead of getting truncated. Approvals (app deploys, multiple-choice questions, credential picks) now render as buttons on Telegram, Discord, Slack, and WhatsApp Cloud, and as Reply YES or NO prompts on Signal, SMS, and personal WhatsApp, with the inbound parser accepting the obvious variants. Behind the scenes, a new Harness struct consolidates the agent runtime view of AppState , sandboxing moves to a single SandboxManager per principal kind, and the channel adapter framework now classifies failures with a typed ChannelError .
Supported channels: Slack, Discord, Telegram, Signal, SMS, WhatsApp Cloud, WhatsApp Personal.
Replace opaque tool_data blobs on tool calls with typed Hitl values; build typed HITL tool components and consume new SSE event names.
Persist Paused message status and resume with an atomic compare-and-swap; translate the new InferenceResponse outcomes into pause/resume persistence.
Add a POST /api/chats/{chat_id}/tool-calls/resolve endpoint, lift the HITL callback parser and response label into a shared hitl module, and drop the legacy app and vault approval routes superseded by the resolve endpoint.
Re-seed pending HITLs from msg.tool_calls on inference_done so paused turns continue to surface new questions.
Share an inbound reply HITL resolver with YES/NO variant parsing ( y , yeah , ok , nope , thumbs-up emoji, and the usual variants).
Resume HITL delivery on resolve so sequential adapters render the next prompt; use the absolute public_base_url for HITL fallback URLs on channels that can't render the picker.
Render HITL prompts and resolve responses on Telegram (inline keyboard), Slack (Block Kit via Socket Mode interactions), Discord (buttons), WhatsApp Cloud (interactive messages and button taps), WhatsApp personal account (quote replies), Signal (quote replies), and SMS ( Reply YES or NO hint on App-deploy approval prompts).
Hold the Discord and WhatsApp typing indicator across long inferences so the user sees activity while the agent is generating.
Rename the manage_service tool to manage_app and adopt typed HITL hooks; exit tasks on HITL and respawn via run_task .
Test HITL persistence, pause rendering, batched resolve, and the resume race.
Add a chat::slash invocation parser for / and @ prefixes; add a MessageCommand side field on Message for parsed slash invocations.
Wire up /commands end to end: dispatch, terminal-write refactor, cross-agent attribution, slash parsing.
Add a GET /api/chats/{id}/commands discovery endpoint and a listCommands API client.
Frontend: slash composer with Lexical triggers, directive chips, and a /new builtin that opens a fresh chat with the current agent.
@<agent> per-turn override: the reply is attributed to the target agent and the next message reverts to the chat's default.
Add save_chat , save_updated_message , and delete_messages_for_chat helpers.
Skills: new SKILL.md frontmatter fields disable-model-invocation , argument-hint , and arguments ; hide disable-model-invocation skills from the model's available-skills block while keeping them in the / menu for users to invoke directly.
Add e2e tests for /agent dispatch and commands discovery.
Add a per-format message splitter with shared boundary primitives that break at paragraph, then line, then word, then UTF-8 character, and never inside a code fence or escape sequence.
Wire TelegramMarkdownV2Splitter into the Telegram adapter (4,096-char per-message limit) and render markdown tables as monospace code blocks since Telegram has no native table support.
Wire MarkdownSplitter into the Discord adapter (2,000-char per-message limit), remove the old chunk_for_discord helper, and render markdown tables as monospace code blocks.
Wire PlainSplitter into the Slack adapter (silent multi-chunk).
Wire PlainSplitter into the SMS adapter with a 1,600-char hard cap and Full reply: {short-link} overflow that points at a Chat -kind share URL (reused across overflow events in the same chat).
Wire MarkdownSplitter into the WhatsApp Cloud and WhatsApp Personal adapters.
Wire SignalSplitter into the Signal adapter with per-chunk SignalText { body, ranges } .
Add ChannelError type and MessageDelivery.failure_kind field; migrate channel adapters to typed ChannelError classification ( Transient , Forbidden , NotFound , PayloadInvalid , PayloadTooLarge , Unauthorized , Other ) with an optional retry_hint .
Make start_with_retry idempotent so a single start no longer double-spawns the adapter.
Replace the bounded broadcast bus with unbounded fan-out so channel inbound stops losing events.
Add a Share entity with File and Chat kinds plus a ShareService for issuing short links; wire the service into AppState and schedule periodic cleanup.
Add share.ttl_secs (default 30 days) and share.cleanup_interval_secs (default 6 hours) config knobs.
Add GET /s/{id} short-link resolver with a 303 redirect to the canonical or presigned URL.
Add a /p preview page that renders markdown and code with syntax highlighting; add a GET /p/{slug} redirector to the query-param form.
Route channel attachments through outbound_url in all adapters via a new chat::channel::attachment helper: inline-channel previewable files get /p/{8-char-id} , inline non-previewable files get /s/{8-char-id} , button channels emit long-form /p/{owner}/{handle}/{path} or /api/files/... URLs.
Add an anonymous_not_found helper so unknown / expired / unauthorized share IDs return byte-identical 404s.
Extract FilePreviewContent into a shared module; extract apiFetch from request<T> for non-JSON authenticated requests.
Add ShareKind::Chat variant with lazy lookup_or_issue_chat for SMS chat-share reuse.
Accept result_description on create_task and create_recurring_task so prose answers skip schema authoring; reject when both result_description and result_schema are passed (or neither).
Require a top-level summary: string on complex task result schemas; reject complex schemas without it at task creation.
Carry result_schema on TaskCompletion events and persist the JSON content.
Render the TaskCompletion bubble with a schema-driven body and a link to the task chat; render TaskCompletion via a shared markdown helper across channel adapters.
Researcher now publishes the full research as a markdown attachment, named after the topic (e.g. h100-used-prices.md ) so successive research tasks don't overwrite each other.
Add typed read , write , edit , glob , and grep tools. Workspace-scoped by default; absolute and policy-authorized sibling paths allowed; virtual-path URIs rejected.
edit uses Unicode normalization (NFKC, ASCII quotes/dashes/spaces, collapsed whitespace) so a slightly-misremembered snippet still matches its target. The motivation is token usage and hallucinations: a typed edit returns a structured diff in a fraction of the tokens a shell round-trip eats, without the model free-styling shell escapes against a half-remembered path.
Add the frona-text crate with shared text and search primitives ( NormalizedString , LineEnding , walk_with_ignore for .gitignore -aware traversal).
Refactor sandbox: SandboxManager is now the single entry point per principal kind.
Extract a per-tool view registry: ToolRow primitives, DefaultView , slim chrome.
Add views for shell, Python, and Node (with sandbox-deny extraction); for the typed file tools ( read , write , edit , glob , grep ) with syntax highlighting and line numbers; for produce_file (filename / size / content-type card); for web_search (result list) and web_fetch (markdown + clickable URL); for task tools ( create_task , create_recurring_task with cronstrue , delete_task ); for memory tools (auto-expand on large text); for update_identity (set/remove attribute rendering); and for set_heartbeat (humanized interval and next-tick time).
Render the shell tool command as a bash-highlighted block with wrap; render produce_file as a filename card with size and content type.
Add vitest test runner and cronstrue / unbash dependencies.
Add a Harness struct as the agent runtime view of AppState ; route session build, voice WS, message stream, and resume-all through it.
Narrow TaskExecutor to hold Arc<Harness> and own its own resolution-notifier map.
Move HITL resolution onto Harness and drop ChannelCtx.app_state .
Move task-tool registration into ToolManager ; extract deliver_event_to_source so ReportSignalTool no longer needs TaskExecutor .
Delete agent/execution.rs and update test fixtures for Harness.
Inject heartbeats as a transient turn so the agent stops replying to them like a user message.
Fold per-turn lifecycle events into InferenceEventKind .
Keep browser sessions alive past Browserless's hard 408-second timeout: pass a long timeout at connect time, self-evict and recycle sessions with a 60-second margin, and evict dead WebSockets immediately.
Stop dropping the newest message when a chat crosses the page limit.
Move the reasoning toggle from inline content to a sparkle icon in the message header.
Disable single-tilde strikethrough in assistant markdown.
Quickstart : bring up the stack with Docker Compose in minutes
Screenshots : see the platform in action
Documentation : full guides and reference
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
