---
source: "https://github.com/drewnekota/cetus"
hn_url: "https://news.ycombinator.com/item?id=49084323"
title: "Show HN: Cetus – A macOS App for Claude Code, Codex, OpenCode, and More"
article_title: "GitHub - drewnekota/cetus: One macOS app for Claude Code, Codex, and every agent runtime you use — scheduled runs, global hotkey launcher, per-run git worktrees, one review board. · GitHub"
author: "williamjinq"
captured_at: "2026-07-28T15:13:54Z"
capture_tool: "hn-digest"
hn_id: 49084323
score: 1
comments: 1
posted_at: "2026-07-28T14:21:06Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Cetus – A macOS App for Claude Code, Codex, OpenCode, and More

- HN: [49084323](https://news.ycombinator.com/item?id=49084323)
- Source: [github.com](https://github.com/drewnekota/cetus)
- Score: 1
- Comments: 1
- Posted: 2026-07-28T14:21:06Z

## Translation

タイトル: Show HN: Cetus – クロード コード、コーデックス、オープンコードなどのための macOS アプリ
記事のタイトル: GitHub -drawnekota/cetus: クロード コード、コーデックス、および使用するすべてのエージェント ランタイム用の 1 つの macOS アプリ — スケジュールされた実行、グローバル ホットキー ランチャー、実行ごとの git ワークツリー、1 つのレビュー ボード。 · GitHub
説明: Claude Code、Codex、および使用するすべてのエージェント ランタイム用の 1 つの macOS アプリ — スケジュールされた実行、グローバル ホットキー ランチャー、実行ごとの Git ワークツリー、1 つのレビュー ボード。 - ドリューネコタ/シータス

記事本文:
GitHub -drawnekota/cetus:Claude Code、Codex、および使用するすべてのエージェント ランタイム用の 1 つの macOS アプリ — スケジュールされた実行、グローバル ホットキー ランチャー、実行ごとの Git ワークツリー、1 つのレビュー ボード。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
ドリューネコタ
/
CE

タス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
97 コミット 97 コミット .github/ workflows .github/ workflows docs docs evals evals package/ cetus-bridge-protocol パッケージ/ cetus-bridge-protocol public/ ブランド public/ ブランド スクリプト scripts src-tauri src-tauri src src .gitignore .gitignore ライセンス ライセンス README.md README.md README.zh-CN.md README.zh-CN.md コンポーネント.json コンポーネント.json next.config.mjs next.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml postcss.config.mjs postcss.config.mjs tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Claude Code、Codex、および使用するすべてのエージェント ランタイム用の 1 つの macOS アプリ。
これらは 1 か所に存在するため、外出中に実行するようにスケジュールを設定したり、ホットキーを使用して任意のアプリ上で 1 つを呼び出したり、それぞれに独自の git ワークツリーを与えたり、1 つのボードですべての結果を確認したりできます。
スケジュールされた実行 · グローバル ホットキー ランチャー · 実行ごとの git ワークツリー · 1 つのレビュー ボード
エージェントのための 1 つの家。 Cetus の組み込みランタイム、Claude Code、Codex 間の会話を、デスクトップ ワークフローを失うことなく切り替えます。
携帯電話から続行してください。オプションの Tailscale 支援モバイル コンパニオンは、パブリック Cetus サービスを公開することなく、ライブ実行の追跡、メッセージと画像の送信、ランタイムの切り替え、承認の処理、およびチャットのアーカイブを行います。 Cetus リモート を参照してください。
安全な並列コーディング セッション。エージェントが現在のチェックアウトに触れるのではなく、個別に編集できるようにする場合は、会話ごとの git worktree を有効にします。
確認できるバックグラウンド作業。タスクをスケジュールし、実行したままにし、ターミナルに埋もれているのではなく、レビューが必要な状態で待機している結果を見つけます。

セッション。
チャット後に残るコンテキスト。ワークスペース、耐久性のあるメモ、会議の記憶、およびオプションのオンデバイス画面コンテキストにより、最後の実行が停止した場所から次の実行を再開することができます。
ローカルコントロール。 Cetus は macOS のネイティブ アプリです。画面 OCR、会議の文字起こし、音声ディクテーションはデバイス上で実行されます。機密性の高い機能はオプトインです。
事前に構築されたアプリは、Apple Silicon および macOS 13 以降をサポートします。
DMG を開き、Cetus をアプリケーションに移動します。
組み込みの Cetus ランタイムを使用するか、すでにインストールされてサインインしているクロードまたはコーデックス CLI を選択します。
ワークスペースを選択し、エージェントに最初のタスクを与えます。
Claude Code と Codex は既存の CLI ログインを再利用します。設定する 2 番目のアカウントはありません。ソースからのビルドについては、「開発」に記載されています。
初期リリース: Cetus は活発に開発中です。何かが壊れた場合、またはワークフローが欠落している場合は、問題を開いてください。
Codex、Claude Code、DeepSeek Agent を並行して実行する
ワークスペースを選択し、ランタイムを選択し、必要に応じてファイルまたはスクリーンショットを添付して送信します。返信は、折りたたみ可能な思考ブロックとツール使用カードを使用してライブでストリーミングされます。並列コーディング タスクの場合は、ワークツリー分離を有効にして、各 CLI 会話が個別のチェックアウトを編集できるようにします。
各ジョブに適切なランタイムを選択する
Cetus はバンドルされた pi ハーネスを使用します。会話を Claude Code または Codex に切り替えて、会話ごとのモデルと推論労力のオーバーライドを使用して、対応するベンダー CLI で会話を実行します。
CLI ランタイムは、すでにインストールされて PATH にログインしているクロード/コーデックスを再利用します。個別のサインインは必要ありません。 Cetus は、会話スコープのランタイムを維持し (ストリーミング入力モードの claude -p / codex app-server )、構造化されたイベント ストリームを同じチャット UI (テキスト、思考、ツール カード) に変換し、応答全体でローカル開発サーバーなどのバックグラウンド ターミナルを保持します。コ

ntext とプロセスのクリーンアップは会話のライフサイクルに従い、編集は会話ごとの git worktree で分離できます。オートメーションは任意のランタイムで起動できるため、チャットが Cetus に留まりながら、スケジュールされたジョブを Claude Code で実行できます。
すべての会話は、「進行中」「レビューが必要」「完了」全体で追跡され、ワー​​クスペースごとに、またはすべてにわたってフィルターされたカードです。ここではバックグラウンド実行 (自動化、並列ソリューション) が表面化するため、複数のセッションにまたがる作業がチャット リストに埋もれることがありません。
スケジュールに従って起動する保存済みプロンプト ( at / each / cron / daily )。オートメーションを使用すると、さまざまなランタイムにわたってスケジュールされたジョブを 1 か所で作成および管理できます。Claude Code ジョブ、Codex ジョブ、および Cetus ジョブはすべて並列して実行でき、それぞれが独自のランタイムとモデル設定を保持します。すべてのトリガーにより、新たなバックグラウンド会話が開始されます — 例:平日 09:00 のニュース ダイジェスト。外出中に過去 24 時間を検索し、HTML 概要を表示します。
現在の画面を持ち歩く
両方の ⌘ キーを押し続けると、グローバル パネルが開きます。使用中のアプリを離れることなく、Cetus に何でも質問できます。目の前にあるものを読み取り、それを取り外し可能なコンテキスト チップとして添付します。画面のスクリーンショット、アクティブなアプリ、現在のブラウザーの URL、および選択したテキストです。有用なものを保持し、残りを削除してから、新しい実行を開始するか、最後の実行を続行します。
メッセージと電子メールの場合、右 ⌥ をダブルタップすると、視覚的なクイック返信が開きます。Cetus は現在の画面をキャプチャし、設定された Gemini または Volcano Ark ビジョン モデルに 3 つの送信準備完了返信を要求し、元のアプリに挿入し直す前に 1 つを選択または編集できます。これは、エージェントを完全に実行するのではなく、直接のワンショット パスです。 [送信] を押すことはありません。画面録画とアクセシビリティ権限が必要です。
コーディング エージェント シェル以上のもの
永続的

memory you and the agent both edit, injected into future turns
Parallel solutions : fan one prompt into N candidate runs, then keep one and archive the rest
Ultra Code mode: author a workflow and orchestrate sub-agents for a single request
Voice dictation (on-device, macOS) — in-app and as a global push-to-talk
Meeting memory (on-device, macOS) — auto-detect, system-audio capture, DeepSeek-distilled minutes the agent can search
Computer and browser control through structured accessibility elements, with confirmation before consequential actions
Anthropic、OpenAI、Google、Bedrock、Ollama、LM Studio、OpenRouter、OpenAI 互換エンドポイントなど、内部に 30 以上のモデル プロバイダー
任意のアプリのホットキーを押したままにして話します。Cetus はフローティング イコライザー HUD をポップし、Seed-ASR を使用してデバイス上で文字起こしし、カーソルのある場所にクリーンアップされたテキストをドロップします。 The same stack as the in-app mic, but it follows you across the desktop.
会議を検索可能なコンテキストに変える
Turn on meeting memory and Cetus quietly transcribes your calls into searchable notes — on-device, text only, no audio stored.
Auto-detect — when another app grabs your mic (Zoom, Teams, FaceTime, Feishu…), Cetus starts a session and stops when the call ends.何も押す必要はありません。
Manual — global hotkey (default ⌘⇧M ) for in-person meetings that auto-detect can't pick up.
両方の側 - あなたのマイクはあなたです。 system audio is everyone else, captured separately so the transcript knows who said what (macOS 14.2+; falls back to mic-only below that).
Transcription is 100% on-device via Apple's Speech framework — streaming, punctuated, segmented on natural pauses. While a session is live, a small floating pill (red dot + elapsed timer + stop button) sits at the top of your screen without stealing focus.
それらの

メモはエージェントが到達できるコンテキストになります。「開始日について何を決めましたか?」と尋ねます。 Cetus は会議履歴 ( search_meeting_history ) を検索します。すべてローカルで、アップロードされたものはありません。デフォルトではオフです。マスター スイッチは、オプトインするまで Cetus がリッスンしないことを意味します。現時点では macOS のみです。
画面に表示されていたものを思い出してください
画面コンテキストがオンの場合、Cetus は定期的にフレームをキャプチャし、知覚ハッシュを使用して重複を排除し、Apple Vision を使用してデバイス上で OCR を実行します。これにより、エージェントは作業内容を思い出すことができ、OCR テキストまたはアプリで検索することもできます。画像とテキストは Mac に残ります。何もアップロードされていません。デフォルトではオフです。コントロールには、キャプチャ間隔、保持期間、機密性の高いアプリ (1Password、メッセージなど) が最前面にある場合にキャプチャを一時停止する除外アプリ リストが含まれます。
各機能はオプトインです。コンピュータとブラウザのコントロールを使用すると、エージェントは番号付きの要素リスト (生のピクセルではない) を介してブラウザと Mac アプリを操作できます。重要な作業 (送信、削除、購入、送信、認証) の前に確認ステップがあり、[停止] ボタンが常に手の届くところにあります。
ターミナル エージェントは個々のタスクには優れていますが、長時間実行される作業はセッション、リポジトリ、バックグラウンド プロセス間で失われやすくなります。 Cetus は、各実行を、ワークスペース、状態、履歴、レビュー ステップを含む可視の作業項目に変換します。
有用なエージェントには、あなたの状況に関するコンテキスト、適切なモデルからのインテリジェンス、および行動する能力が必要です。 Cetus はこれらの部分を独立させます。各タスクのランタイムを選択し、必要なコンテキストのみを追加して、結果の作業を検査可能にします。
これにより、ターミナル タブに収まりきらないワークフローが実用的になります。
外出中にエージェントを実行し、後で結果を確認します。
git の変更を衝突させることなく、独立したソリューションを比較します。
プロジェクトの決定と事前の決定を行う

次の会話への言及。
コーディング作業を、その周囲の会議、画面、アプリに接続します。
3 つの要素は 1 つの瞬間を表します。エージェントが時間の経過とともに役立つかどうかは、エージェントが何かを蓄積するかどうかによって決まります。
メモリはエージェントが自身に書き戻すコンテキストです。そのため、次のセッションは最初から開始するのではなく、最後のセッションが中断したところから再開されます。
あなたがアイドル状態である間も夢は続きます。Cetus は最近の会話を振り返り、それを永続的なメモに統合し、生の履歴を永続的な好みに変えます。デフォルトではオンです。
ノード ≥ 20、 pnpm 、 bun (pi サイドカー バイナリの構築用)
Tauri の前提条件: https://v2.tauri.app/start/prerequisites/
DEEPSEEK_API_KEY (または選択したプロバイダー。pi は ANTHROPIC_API_KEY 、 OPENAI_API_KEY などを自動的に取得します)
オプション: Claude Code ( claude ) および/または Codex ( codex ) CLI をインストールしてログインします (会話ランタイムとして使用する場合) — Cetus は既存のログインを再利用し、追加のセットアップは必要ありません
pnpmインストール
# pi を単一ファイルのバイナリとして src-tauri/binaries/pi-<target> にビルドします。
# 30 秒ほどかかります。開発マシンごとに 1 回実行します。バイナリは gitignor されます。
./scripts/build-pi-sidecar.sh
開発環境で実行
エクスポート DEEPSEEK_API_KEY=sk-...
pnpm タウリ開発者
Tauri は Next.js 開発サーバー (ポート 3000) とウィンドウを起動します。

[切り捨てられた]

## Original Extract

One macOS app for Claude Code, Codex, and every agent runtime you use — scheduled runs, global hotkey launcher, per-run git worktrees, one review board. - drewnekota/cetus

GitHub - drewnekota/cetus: One macOS app for Claude Code, Codex, and every agent runtime you use — scheduled runs, global hotkey launcher, per-run git worktrees, one review board. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
drewnekota
/
cetus
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
97 Commits 97 Commits .github/ workflows .github/ workflows docs docs evals evals packages/ cetus-bridge-protocol packages/ cetus-bridge-protocol public/ brands public/ brands scripts scripts src-tauri src-tauri src src .gitignore .gitignore LICENSE LICENSE README.md README.md README.zh-CN.md README.zh-CN.md components.json components.json next.config.mjs next.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml postcss.config.mjs postcss.config.mjs tsconfig.json tsconfig.json View all files Repository files navigation
One macOS app for Claude Code, Codex, and every agent runtime you use.
They live in one place, so you can schedule them to run while you're away, summon one over any app with a hotkey, give each its own git worktree, and review every result on one board.
Scheduled runs · Global hotkey launcher · Per-run git worktrees · One review board
One home for your agents. Switch any conversation between Cetus's built-in runtime, Claude Code, and Codex without losing the desktop workflow around it.
Continue from your phone. The optional, Tailscale-backed mobile companion follows live runs, sends messages and images, switches runtimes, handles approvals, and archives chats without exposing a public Cetus service. See Cetus Remote .
Safe, parallel coding sessions. Enable per-conversation git worktrees when you want agents to edit in isolation instead of touching the current checkout.
Background work you can review. Schedule a task, leave it running, and find the result waiting in Needs review instead of buried in a terminal session.
Context that survives the chat. Workspaces, durable notes, meeting memory, and optional on-device screen context help the next run pick up where the last one stopped.
Local control. Cetus is a native macOS app. Screen OCR, meeting transcription, and voice dictation run on-device; sensitive capabilities are opt-in.
The prebuilt app supports Apple Silicon and macOS 13 or later .
Open the DMG and move Cetus to Applications.
Use the built-in Cetus runtime, or select an already installed and signed-in claude or codex CLI.
Choose a workspace and give the agent its first task.
Claude Code and Codex reuse their existing CLI login — there is no second account to configure. Building from source is documented under Development .
Early release: Cetus is under active development. Please open an issue if something breaks or if a workflow is missing.
Run Codex, Claude Code and DeepSeek Agent side by side
Pick a workspace , choose a runtime, optionally attach files or a screenshot, and send. Replies stream live with collapsible thinking blocks and tool-use cards. For parallel coding tasks, enable worktree isolation so each CLI conversation edits a separate checkout.
Pick the right runtime for each job
Cetus uses the bundled pi harness. Switch a conversation to Claude Code or Codex to run it on the corresponding vendor CLI, with per-conversation model and reasoning-effort overrides.
The CLI runtimes reuse whatever claude / codex you already have installed and logged in on your PATH — no separate sign-in. Cetus keeps a conversation-scoped runtime alive ( claude -p in streaming-input mode / codex app-server ), translates its structured event stream into the same chat UI (text, thinking, tool cards), and preserves background terminals such as local dev servers across replies. Context and process cleanup follow the conversation lifecycle, and edits can be isolated in a per-conversation git worktree . Automations can fire on any runtime, so a scheduled job can run on Claude Code while your chats stay on Cetus.
Every conversation is a card tracked across In progress · Needs review · Done , filtered by workspace or across all of them. Background runs (automations, parallel solutions) surface here, so work that spans multiple sessions doesn't get buried in a chat list.
Saved prompts that fire on a schedule ( at / every / cron / daily ). Automations gives you one place to create and manage scheduled jobs across different runtimes — a Claude Code job, a Codex job, and a Cetus job can all live side by side, each keeping its own runtime and model settings. Every trigger starts a fresh background conversation — e.g. a weekday-09:00 news digest that searches the last 24 hours and renders an HTML summary while you're away.
Bring the current screen with you
Hold both ⌘ keys to open the global panel: ask Cetus anything without leaving the app you're in. It reads what's in front of you and attaches it as removable context chips: a screenshot of your screen, the active app, the current browser URL, and any selected text. Keep what's useful, drop the rest, then start a new run or continue the last one.
For messages and email, double-tap right ⌥ opens visual quick reply: Cetus captures the current screen, asks the configured Gemini or Volcano Ark vision model for three send-ready replies, and lets you choose or edit one before inserting it back into the original app. This is a direct one-shot path rather than a full agent run; it never presses Send. Screen Recording and Accessibility permissions are required.
More than a coding-agent shell
Persistent memory you and the agent both edit, injected into future turns
Parallel solutions : fan one prompt into N candidate runs, then keep one and archive the rest
Ultra Code mode: author a workflow and orchestrate sub-agents for a single request
Voice dictation (on-device, macOS) — in-app and as a global push-to-talk
Meeting memory (on-device, macOS) — auto-detect, system-audio capture, DeepSeek-distilled minutes the agent can search
Computer and browser control through structured accessibility elements, with confirmation before consequential actions
30+ model providers under the hood , including Anthropic, OpenAI, Google, Bedrock, Ollama, LM Studio, OpenRouter, and OpenAI-compatible endpoints
Hold a hotkey from any app and talk — Cetus pops a floating equalizer HUD, transcribes on-device with Seed-ASR, and drops the cleaned-up text wherever your cursor is. The same stack as the in-app mic, but it follows you across the desktop.
Turn meetings into searchable context
Turn on meeting memory and Cetus quietly transcribes your calls into searchable notes — on-device, text only, no audio stored.
Auto-detect — when another app grabs your mic (Zoom, Teams, FaceTime, Feishu…), Cetus starts a session and stops when the call ends. Nothing to press.
Manual — global hotkey (default ⌘⇧M ) for in-person meetings that auto-detect can't pick up.
Both sides — your mic is you; system audio is everyone else, captured separately so the transcript knows who said what (macOS 14.2+; falls back to mic-only below that).
Transcription is 100% on-device via Apple's Speech framework — streaming, punctuated, segmented on natural pauses. While a session is live, a small floating pill (red dot + elapsed timer + stop button) sits at the top of your screen without stealing focus.
Those notes become context the agent can reach: ask "what did we decide about the launch date?" and Cetus searches your meeting history ( search_meeting_history ) — all local, nothing uploaded. Off by default; the master switch means Cetus never listens until you opt in. macOS only for now.
Remember what was on your screen
With screen context on, Cetus periodically captures frames, dedupes them with a perceptual hash, and OCRs on-device with Apple Vision so the agent can recall what you were working on — and you can search by OCR text or app. Images and text stay on your Mac; nothing is uploaded. Off by default; controls include capture interval, retention period, and an excluded-apps list that pauses capture when sensitive apps (1Password, Messages…) are frontmost.
Each capability is opt-in. Computer & Browser control lets the agent drive your browser and Mac apps through numbered element lists (not raw pixels), with a confirmation step before anything consequential (sending, deleting, purchasing, submitting, authenticating) and a Stop button always in reach.
Terminal agents are excellent at individual tasks, but long-running work is easy to lose across sessions, repositories, and background processes. Cetus turns each run into a visible work item with a workspace, state, history, and review step.
A useful agent needs context about your situation, intelligence from the right model, and abilities to act. Cetus keeps those pieces independent: choose the runtime for each task, add only the context you want, and make the resulting work inspectable.
That makes workflows practical that do not fit neatly into a terminal tab:
Run an agent while you are away and review the result later.
Compare independent solutions without colliding git changes.
Carry project decisions and preferences into the next conversation.
Connect coding work to the meetings, screens, and apps surrounding it.
The three factors describe a single moment. What makes an agent useful over time is whether it accumulates anything.
Memory is context the agent writes back to itself — so the next session picks up where the last one left off instead of starting from scratch.
Dreaming runs while you're idle: Cetus reflects on recent conversations and consolidates them into durable notes, turning raw history into preferences that persist. On by default.
Node ≥ 20, pnpm , bun (for building the pi sidecar binary)
Tauri prerequisites: https://v2.tauri.app/start/prerequisites/
A DEEPSEEK_API_KEY (or your provider of choice; pi auto-picks up ANTHROPIC_API_KEY , OPENAI_API_KEY , etc.)
Optional : the Claude Code ( claude ) and/or Codex ( codex ) CLIs installed and logged in, if you want them as conversation runtimes — Cetus reuses their existing login, no extra setup
pnpm install
# Build pi as a single-file binary into src-tauri/binaries/pi-<target>.
# Takes ~30s. Run once per dev machine; binaries are gitignored.
./scripts/build-pi-sidecar.sh
Run in dev
export DEEPSEEK_API_KEY=sk-...
pnpm tauri dev
Tauri launches the Next.js dev server (port 3000) and a window

[truncated]
