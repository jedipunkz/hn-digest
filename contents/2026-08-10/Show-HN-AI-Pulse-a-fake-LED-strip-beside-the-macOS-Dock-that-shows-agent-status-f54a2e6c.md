---
source: "https://github.com/leog/ai-pulse"
hn_url: "https://news.ycombinator.com/item?id=49250486"
title: "Show HN: AI Pulse a fake LED strip beside the macOS Dock that shows agent status"
article_title: "GitHub - leog/ai-pulse: Ambient macOS light strip for AI coding agents — a Dock-adjacent pill of eight LEDs showing whether anything is working, waiting, finished, or broken · GitHub"
author: "leog_me"
captured_at: "2026-08-10T22:26:55Z"
capture_tool: "hn-digest"
hn_id: 49250486
score: 1
comments: 0
posted_at: "2026-08-10T22:07:41Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI Pulse a fake LED strip beside the macOS Dock that shows agent status

- HN: [49250486](https://news.ycombinator.com/item?id=49250486)
- Source: [github.com](https://github.com/leog/ai-pulse)
- Score: 1
- Comments: 0
- Posted: 2026-08-10T22:07:41Z

## Translation

タイトル: Show HN: AI エージェントのステータスを示す macOS Dock の横にある偽の LED ストリップをパルスします。
記事のタイトル: GitHub - leog/ai-pulse: AI コーディング エージェント用のアンビエント macOS ライト ストリップ — 何かが動作しているか、待機中か、完了したか、壊れているかを示す 8 個の LED のドックに隣接した錠剤 · GitHub
説明: AI コーディング エージェント用のアンビエント macOS ライト ストリップ - ドックに隣接する 8 個の LED で、動作中、待機中、完了、または壊れているかどうかを表示します - leog/ai-pulse
HN テキスト: こんにちは、HN!いくつかの Claude Code セッションを並行して実行し、cmd キーを押しながら移動し続けましたが、そのうちの 1 つが 10 分間許可プロンプトに座っていたことがわかりました。私が気に入ったハードウェア ガジェット (SidePulse.io と呼ばれる) があったので、出荷を待つ前に代わりにソフトウェア バージョンを構築しました :D 皆さんも気に入っていただき、私と同じように役立つと思っていただければ幸いです。

記事本文:
GitHub - leog/ai-pulse: AI コーディング エージェント用のアンビエント macOS ライト ストリップ — 何かが動作しているか、待機中か、完了したか、壊れているかを示す 8 個の LED のドックに隣接した錠剤 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
レオグ
/
アイパルス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
9 コミット 9 コミット .claude .claude .github/ workflows .github/ workflows スクリプト スクリプト ソース ソース テスト/ AIPulseKitT

ests Tests/ AIPulseKitTests docs docs .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス Package.swift Package.swift README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング エージェント用のアンビエント macOS ライト ストリップ。
SidePulse ハードウェア ガジェット — ただしアプリとして。仮想 8 つのスリムなストリップ
LED はドックの横の未使用スペースに浮かんでおり、単一の LED で表示されます。
信号が集約されており、モデルごとの区別はありません。何かが実行されているかどうかに関係なく、
あなたを待っています、終わった、または壊れています:
Reduce Motion は、すべてのアニメーションを静的な処理に置き換えます。の
エージェントごとのオリジナルのアイコンピルは引き続き「設定」→「外観」→「設定」から利用可能です
インジケーターのスタイル。
AI Pulse は Dock に埋め込まれていません — macOS は、AI Pulse のパブリック API を公開していません
ドックのアクセサリ。これは、Dock に隣接し、ボーダーレスで、非アクティブな NSPanel です。
パブリック スクリーン ジオメトリから配置されます ( NSScreen.frame とvisibleFrame )。
プライベート API、アクセシビリティ/画面録画権限、インジェクションはありません。
AI-Pulse-<バージョン>.zip を次の場所から取得します。
最新リリース、解凍、
AI Pulse.app を /Applications に移動して起動します。リリースはまだです
公証されているため、最初にアプリを右クリック→開く (macOS 15 以降では、
[システム設定] → [プライバシーとセキュリティ] → [とにかく開く] でも許可します)。
aipulse CLI はバンドル内に同梱されています。
AI Pulse.app/Contents/Helpers/aipulse 。
PR がマージされると、リリースは自動的にカットされます: デフォルトでパッチ、マイナーまたは
PR に release:minor / release:major ラベルが付いている場合はメジャー。
要件 (ソースからの構築)
swift build # すべてをビルドします (アプリ、CLI、キット)
迅速なテスト # 単体テスト + 統合テスト (AIPulseKit)
swift run AIPulseApp # ピルを起動します (アクセサリ アプリ: Dock アイコンなし)
迅速に AIPulseApp --snapshot DIR を実行します # ピル + ホバー カード PNG をヘッドレスでレンダリングして終了します
素早く走る

aipulse health # CLI: ローカル イベント サービスを確認する
./Scripts/make-app.sh # アセンブル + 署名 dist/AI Pulse.app
./Scripts/make-gifs.sh # ヘッドレスフレームから docs/GIF を再生成 (ffmpeg が必要)
docs/pulse-social.gif は、同じ状態サイクルをキャプション付きで拡大したものです。
共有するために作られました。
日常的に使用する場合は、ベアバイナリではなくバンドルされたアプリを実行します。
スクリプトは Apple 開発 ID で署名し、安定したコードを提供します
署名により、再構築後もキーチェーンがアプリを信頼します (素早い実行)
バイナリはアドホック署名されており、その後キーチェーンの同意プロンプトがトリガーされます。
すべてのリビルド — これが、AIPULSE_DEV_EPHEMERAL_TOKEN=1 が存在する理由でもあります。
開発ループ）。 CLI は AI Pulse.app/Contents/Helpers/aipulse に埋め込まれています。
(アプリ製品は AIPulseApp です。AIPulse と aipulse CLI は
大文字と小文字を区別しないファイルシステムでは衝突します。)
AI Pulse は 127.0.0.1:7455 で待機します (設定で構成可能)。起動したら
キーチェーンにベアラー トークンを保存し、aipulse CLI が認証できるように ~/Library/Application Support/AIPulse/cli.json (モード 0600) を書き込みます。
シェルコマンドにトークンが出現することはありません。
aipulse エージェントのアップサート \
--id " クロードコード: $PWD : $SESSION_ID " \
--name " クロード コード " --provider anthropic \
--instance " $( ベース名 " $PWD " ) " \
--state working --message " ドック配置の実装 "
アイパルスエージェントアップデート\
--id " クロードコード: $PWD : $SESSION_ID " \
--statewaitingForInput --message " 許可を待っています " --sequence 2
aipulse エージェント削除 --id " クロードコード: $PWD : $SESSION_ID "
aipulse エージェント # ピルが知っていることをすべてリストアップします
エンドポイント: POST /v1/agents/upsert 、POST /v1/agents/{id}/event 、
DELETE /v1/agents/{id} 、 GET /v1/agents 、 GET /v1/health (health は
認証されていないルートのみ）。サーバーはループバック インターフェイスにバインドします
のみ、キャップリクエスト si

zes、すべてのペイロードを検証します (AgentReducer +
RequestValidator )、安全でない URL スキームを拒否し、決して実行されません
イベントで受け取ったもの。
開発者メモ: アプリを再構築すると、そのアドホック コード署名が変更されるため、キーチェーン
ビルドごとに再プロンプトを読み取ります。中に AIPULSE_DEV_EPHEMERAL_TOKEN=1 で起動します
開発では、キーチェーンをスキップし、代わりに実行ごとのトークンを使用します。
錠剤の右クリック メニュー → AI Pulse を終了 (またはプロセスを強制終了) して終了します。
Sources/AIPulseKit — AppKit を使用しない、完全に単体テスト済みのコア:
Domain/ — Agent 、 AgentState 、 AgentAction (型指定された安全なアクション +
URL スキーム許可リスト)、AgentEventPayload (ワイヤー モデル)、
AgentIntegrationLevel 、 StatusPriority (緊急度ソート)。
Store/AgentStore — 唯一の信頼できる情報源。すべてのサーフェスはそこからレンダリングされます。
Placement/ — ScreenSnapshot 、ベストエフォート型の DockGeometry 推論、
純粋な PlacementPolicy (側溝 → 隣接 → コーナー フォールバック、クランプ)。
ソース/AIPulse — アプリ:
Presentation/Pill/ — 非アクティブ化 AIPulsePanel 、SwiftUI カプセル、
状態ごとのアイコン (グリフ + バッジ + リング、色だけを使用しないでください)、ホバー カード。
プレゼンテーション/AgentList/ 、プレゼンテーション/設定/ — 従来型
キーボードでアクセスできるウィンドウ。
Placement/DockPlacementController — デバウンスされた画面変更
観察;投票はありません。
M1 — 非アクティブ化パネル、ピル UI、モック エージェント、ホバー カード、コンテキスト メニュー。
M2 — 下/左/右のドック配置、マルチディスプレイ選択
設定表示ピッカー、デバウンス画面/空間変化観察、
オフスクリーンクランプ、オプトインの自動非表示ドックフォロー (ベストエフォート、から
Dock のパブリック設定ドメイン — 読み取り専用、Dock との対話なし)。
M3 — AgentReducer イベントの正規化 (バージョン管理、ID 検証、
シーケンス/タイムスタンプの順序付け、重複の拒否、セーフ アクション マッピング)、
一元化された StalenessPolicy (動作中のエージェントが古いものになる)

設定可能なものを使用する
沈黙が続き、さらに長くなると切断状態に降格されます。エージェントの
バッキングプロセスが終了した場合、PID を介して 1 回のスイープ内で降格されます。
生存チェック。完了したエージェントは、構成可能な遅延の後に期限切れになります。
待機中、承認中、失敗した状態はタイマーで期限切れになることはありません）、永続性は
~/Library/Application Support/AIPulse/agents.json (再起動復元付き):
未解決のアテンション状態はそのままに戻り、ライブのみの状態は戻ります
情報源が再び報告するまで切断されたままになります。
M4 — ループバック専用の HTTP イベント サービス ( LocalHTTPServer )、キーチェーン
CLI 用のベアラー トークン + 0600 ハンドシェイク ファイル、トランスポート検証、
aipulse パブリッシャー CLI、およびエンドツーエンドの統合テスト。確認済み
ライブ: CLI → HTTP → リデューサ → ストア → ピルの往復で最大 150 ミリ秒。
M5 — クロード コード アダプター ( ClaudeCodeAdapter + aipulse claude-hook )、
文書化されたフックのライフサイクルに基づいて構築されています: SessionStart→idle、
UserPromptSubmit/PreToolUse→working、PermissionRequest、
許可プロンプト 通知→承認必須、アイドルプロンプト
通知→入力待ち、停止→完了、停止失敗→失敗、
セッション終了→削除されました。プロジェクトごとのセッションごとに 1 つの錠剤エントリ。プロンプトテキスト、
ツールの入力とアシスタントの出力は決してデコードされないため、デコードされることはありません。
出版されました。フックは常に 0 で終了し、何も出力しません。
M6 — オプションの動的 Dock アイコン (デフォルトではオフ、設定 →
外観）。同じ AgentStore スナップショットをピルとして集約します。
DockTileAggregator : 失敗 > 注意 > 動作中 > 中立、
状態処理とバッジラベルの NSDockTile カスタム コンテンツ ビュー
認識されていない緊急エージェントの数を示します。浮いている錠剤と
Dock アイコンは個別に切り替え可能です。アイコンが跳ね返ることはありません。
6 つの MVP マイルストーンがすべて完了し、次に次の目標へのピボットが続きます。

e
ライトファーストプレゼンテーション ( LightAggregator + LightStripView )、維持
アイコンピルはオプションです。
このリポジトリには、(デバッグ ビルド パス経由で) 関連するフック イベントの aipulse claude-hook を登録する .claude/settings.json が同梱されています。
このリポジトリの Claude Code セッションは、
アプリが実行中です。システム全体で使用する場合:
迅速なビルド -c リリース
sudo cp .build/release/aipulse /usr/local/bin/
次に、同じフックを ~/.claude/settings.json に登録し、
プレーン aipulse claude-hook を使用したコマンド。クロード・コードは承認を求めます
プロジェクトがフックを初めてロードするときにフックします。
AI Pulse には、ローカル エージェント統合によって送信されたステータスが表示されます。読み取れない
プロンプト、ターミナルのコンテンツ、エディタのコンテンツ、またはアプリケーション ウィンドウ。
貢献は歓迎です — 詳細については CONTRIBUTING.md を参照してください。
プロジェクトの設計原則、テストの期待、新しい機能を追加する方法
エージェントの統合。 docs/DEBUGGING.md では、
デバッグ面: 環境変数、ヘッドレススナップショット、ログ
ストリーミング、ディスク上の状態、一般的な障害モードの修正。
AI コーディング エージェント用のアンビエント macOS ライト ストリップ - ドックに隣接する 8 個の LED で、動作中、待機中、完了、または壊れているかどうかを表示します。
Readme MIT ライセンス
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Ambient macOS light strip for AI coding agents — a Dock-adjacent pill of eight LEDs showing whether anything is working, waiting, finished, or broken - leog/ai-pulse

Hi HN! I run a few Claude Code sessions in parallel and kept cmd-tabbing around just to find out one of them had been sitting on a permission prompt for ten minutes. There's a hardware gadget I liked (called SidePulse.io) so before waiting to get my shipment I built the software version instead :D I hope you like it and find it useful as I do!

GitHub - leog/ai-pulse: Ambient macOS light strip for AI coding agents — a Dock-adjacent pill of eight LEDs showing whether anything is working, waiting, finished, or broken · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
leog
/
ai-pulse
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
9 Commits 9 Commits .claude .claude .github/ workflows .github/ workflows Scripts Scripts Sources Sources Tests/ AIPulseKitTests Tests/ AIPulseKitTests docs docs .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Package.swift Package.swift README.md README.md View all files Repository files navigation
An ambient macOS light strip for AI coding agents, inspired by the
SidePulse hardware gadget — but as an app. A slim strip of eight virtual
LEDs floats in the unused space beside the Dock and shows, with a single
aggregate signal and no per-model distinction, whether anything is running,
waiting for you, finished, or broken:
Reduce Motion replaces every animation with a static treatment. The
original per-agent icon pill remains available via Settings → Appearance →
Indicator style.
AI Pulse is not embedded in the Dock — macOS exposes no public API for
Dock accessories. It is a Dock-adjacent, borderless, nonactivating NSPanel
positioned from public screen geometry ( NSScreen.frame vs visibleFrame ).
No private APIs, no Accessibility/Screen Recording permissions, no injection.
Grab AI-Pulse-<version>.zip from the
latest release , unzip,
move AI Pulse.app to /Applications , and launch. Releases are not yet
notarized, so on first launch right-click the app → Open (on macOS 15+,
also allow it under System Settings → Privacy & Security → Open Anyway ).
The aipulse CLI ships inside the bundle at
AI Pulse.app/Contents/Helpers/aipulse .
Releases are cut automatically when a PR merges: patch by default, minor or
major when the PR carries a release:minor / release:major label.
Requirements (building from source)
swift build # build everything (app, CLI, kit)
swift test # unit + integration tests (AIPulseKit)
swift run AIPulseApp # launch the pill (accessory app: no Dock icon)
swift run AIPulseApp --snapshot DIR # render pill + hover card PNGs headlessly and exit
swift run aipulse health # CLI: check the local event service
./Scripts/make-app.sh # assemble + sign dist/AI Pulse.app
./Scripts/make-gifs.sh # regenerate docs/ GIFs from headless frames (needs ffmpeg)
docs/pulse-social.gif is a larger, captioned cut of the same state cycle,
made for sharing.
For day-to-day use, run the bundled app rather than the bare binary: the
script signs it with your Apple Development identity, giving a stable code
signature so the Keychain trusts the app across rebuilds (bare swift run
binaries are ad-hoc signed and trigger a Keychain consent prompt after
every rebuild — that is also why AIPULSE_DEV_EPHEMERAL_TOKEN=1 exists for
dev loops). The CLI is embedded at AI Pulse.app/Contents/Helpers/aipulse .
(The app product is AIPulseApp because AIPulse and the aipulse CLI would
collide on case-insensitive filesystems.)
AI Pulse listens on 127.0.0.1:7455 (configurable in Settings). On launch it
stores a bearer token in the Keychain and writes ~/Library/Application Support/AIPulse/cli.json (mode 0600) so the aipulse CLI authenticates
without tokens ever appearing in shell commands:
aipulse agent upsert \
--id " claude-code: $PWD : $SESSION_ID " \
--name " Claude Code " --provider anthropic \
--instance " $( basename " $PWD " ) " \
--state working --message " Implementing Dock placement "
aipulse agent update \
--id " claude-code: $PWD : $SESSION_ID " \
--state waitingForInput --message " Waiting for permission " --sequence 2
aipulse agent remove --id " claude-code: $PWD : $SESSION_ID "
aipulse agents # list everything the pill knows
Endpoints: POST /v1/agents/upsert , POST /v1/agents/{id}/event ,
DELETE /v1/agents/{id} , GET /v1/agents , GET /v1/health (health is the
only unauthenticated route). The server binds to the loopback interface
only, caps request sizes, validates every payload ( AgentReducer +
RequestValidator ), rejects unsafe URL schemes, and never executes
anything received in an event.
Dev note: rebuilding the app changes its ad-hoc code signature, so Keychain
reads re-prompt per build; launch with AIPULSE_DEV_EPHEMERAL_TOKEN=1 during
development to skip the Keychain and use a per-run token instead.
Quit via the pill's right-click menu → Quit AI Pulse (or kill the process).
Sources/AIPulseKit — AppKit-free, fully unit-tested core:
Domain/ — Agent , AgentState , AgentAction (typed safe actions +
URL scheme allowlist), AgentEventPayload (wire model),
AgentIntegrationLevel , StatusPriority (urgency sort).
Store/AgentStore — single source of truth; all surfaces render from it.
Placement/ — ScreenSnapshot , best-effort DockGeometry inference,
pure PlacementPolicy (gutter → adjacent → corner fallbacks, clamping).
Sources/AIPulse — the app:
Presentation/Pill/ — nonactivating AIPulsePanel , SwiftUI capsule,
per-state icons (glyph + badge + ring, never color alone), hover card.
Presentation/AgentList/ , Presentation/Settings/ — conventional
keyboard-accessible windows.
Placement/DockPlacementController — debounced screen-change
observation; no polling.
M1 — nonactivating panel, pill UI, mock agents, hover card, context menu.
M2 — bottom/left/right Dock placement, multi-display selection with a
Settings display picker, debounced screen/space-change observation,
off-screen clamping, opt-in auto-hidden-Dock following (best-effort, from
the Dock's public preference domain — read-only, no Dock interaction).
M3 — AgentReducer event normalization (versioning, ID validation,
sequence/timestamp ordering, duplicate rejection, safe-action mapping),
centralized StalenessPolicy (working agents stale after a configurable
silence, then demoted to disconnected after a longer one; agents whose
backing process has exited are demoted within one sweep via a PID
liveness check; completed agents expire after a configurable delay;
waiting, approval, and failed states never expire on timers), and persistence to
~/Library/Application Support/AIPulse/agents.json with restart restore:
unresolved attention states come back as-is, live-only states come back
as disconnected until their source reports again.
M4 — loopback-only HTTP event service ( LocalHTTPServer ), Keychain
bearer token + 0600 handshake file for the CLI, transport validation,
the aipulse publisher CLI, and end-to-end integration tests. Verified
live: CLI → HTTP → reducer → store → pill in ~150 ms round-trip.
M5 — Claude Code adapter ( ClaudeCodeAdapter + aipulse claude-hook ),
built against the documented hook lifecycle: SessionStart→idle,
UserPromptSubmit/PreToolUse→working, PermissionRequest and
permission-prompt Notifications→approvalRequired, idle-prompt
Notifications→waitingForInput, Stop→completed, StopFailure→failed,
SessionEnd→removed. One pill entry per session per project. Prompt text,
tool inputs, and assistant output are never decoded, so they can never be
published. The hook always exits 0 and prints nothing.
M6 — optional dynamic Dock icon (off by default; Settings →
Appearance). Aggregates the same AgentStore snapshot as the pill via
DockTileAggregator : failure > attention > working > neutral, with an
NSDockTile custom content view for the state treatment and badgeLabel
showing the count of unacknowledged urgent agents. The floating pill and
the Dock icon are independently toggleable; the icon never bounces.
All six MVP milestones are complete, followed by the pivot to the
lights-first presentation ( LightAggregator + LightStripView ), keeping
the icon pill as an option.
This repository ships .claude/settings.json registering aipulse claude-hook for the relevant hook events (via the debug build path), so
Claude Code sessions in this repo appear in the pill automatically once the
app is running. For system-wide use:
swift build -c release
sudo cp .build/release/aipulse /usr/local/bin/
then register the same hooks in ~/.claude/settings.json , replacing the
command with plain aipulse claude-hook . Claude Code asks you to approve
project hooks the first time it loads them.
AI Pulse displays status sent by local agent integrations. It does not read
prompts, terminal contents, editor contents, or application windows.
Contributions are welcome — see CONTRIBUTING.md for the
project's design principles, testing expectations, and how to add a new
agent integration. docs/DEBUGGING.md covers the
debugging surface: environment variables, headless snapshots, log
streaming, on-disk state, and fixes for the common failure modes.
Ambient macOS light strip for AI coding agents — a Dock-adjacent pill of eight LEDs showing whether anything is working, waiting, finished, or broken
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
