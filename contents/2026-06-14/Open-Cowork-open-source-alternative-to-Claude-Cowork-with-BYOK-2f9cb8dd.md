---
source: "https://github.com/coasty-ai/open-cowork"
hn_url: "https://news.ycombinator.com/item?id=48523556"
title: "Open-Cowork: open-source alternative to Claude Cowork with BYOK"
article_title: "GitHub - coasty-ai/open-cowork: Open-source alternative to Claude Co-Work built by Claude Mythos 5. with BYOK · GitHub"
author: "PrateekJ1703"
captured_at: "2026-06-14T04:35:26Z"
capture_tool: "hn-digest"
hn_id: 48523556
score: 1
comments: 0
posted_at: "2026-06-14T02:18:18Z"
tags:
  - hacker-news
  - translated
---

# Open-Cowork: open-source alternative to Claude Cowork with BYOK

- HN: [48523556](https://news.ycombinator.com/item?id=48523556)
- Source: [github.com](https://github.com/coasty-ai/open-cowork)
- Score: 1
- Comments: 0
- Posted: 2026-06-14T02:18:18Z

## Translation

タイトル: Open-Cowork: BYOK を使用した Claude Cowork に代わるオープンソース
記事タイトル: GitHub - Coasty-ai/open-cowork: Claude Mythos 5. with BYOK · GitHub によって構築された Claude Co-Work のオープンソース代替
説明: Claude Mythos 5 によって構築された Claude Co-Work のオープンソース代替品 (BYOK 付き) - Coasty-ai/open-cowork

記事本文:
GitHub - Coasty-ai/open-cowork: Claude Mythos 5 によって構築された Claude Co-Work のオープンソース代替。BYOK · GitHub
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
海岸愛
/
オープンコワーク
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション

オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
72 コミット 72 コミット .github/ workflows .github/ workflows アプリ アプリ e2e e2e パッケージ パッケージ public public scripts scripts tools/mock-coasty tools/mock-coasty .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .prettierignore .prettierignore .prettierrc.json .prettierrc.json ARCHITECTURE.md ARCHITECTURE.md COTRIBUTING.md COTRIBUTING.md COOKBOOK.md COOKBOOK.md DECISIONS.md DECISIONS.md DEPLOYMENT.md DEPLOYMENT.md DESIGN_SYSTEM.md DESIGN_SYSTEM.md ライセンス ライセンスPLAN.md PLAN.md README.md README.md RUN_LOCALLY.md RUN_LOCALLY.md SECURITY.md SECURITY.md SUMMARY.md SUMMARY.md UI_AUDIT.md UI_AUDIT.md eslint.config.mjs eslint.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json Turbo.json Turbo.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
コンピューターのタスクを AI の同僚に引き継ぎ、どこからでも作業を監視し、承認します。
画面を見てそれに基づいて動作する、オープンソースのクロスプラットフォームのエージェント型コワーカー —
自分のデスクトップ、クラウド VM、またはブラウザー。すべてのステップをライブでストリーミングし、あなたのために一時停止します
承認を得ることができ、支出を可視化して上限を設け続けることができます。
Coasty コンピュータ上で実行 すぐに使える API を使用 —
または、独自の LLM (OpenRouter · OpenAI · ローカル モデル) を持ち込んでください。あなたの電話です。
クイックスタート ·
独自の LLM を持参する ·
PC を自動化する ·
特徴・
仕組み ·
ドキュメント
タスクを委任 → 同僚がブラウザを操作する様子を段階的に観察 → 結果を取得します。デモ モードではセットアップを行わずに実行します。
前提条件: ノード ≥ 22.5 (24 を使用) · pnpm 10 (コアパック有効)。
git clone https://github.com/coasty-ai/open-cowork.git && cd open-cowork
pnpm install # モノリポジトリ全体に対して 1 回のインストール
ピン

pm Desktop # ← デスクトップ アプリを実行します: バックエンド + Web を開始し、ウィンドウを開きます
たった 1 つのコマンドで、設定は不要です。 pnpm デスクトップがバックエンドと Web を起動します
UI を選択し、デスクトップ アプリ (独自の画面を駆動できるビルド) を開きます。
キーが設定されていない場合は、バンドルされたモックサーバーと使い捨てのデモモードで実行されます。
サンドボックス キー — アカウント、ネットワーク、請求は不要です。
デリゲートで、「このコンピューター (ローカル画面)」を選択します。
タスクを入力→コストを確認→作業の様子を観察します。 (ヒント: 承認フローの一時停止と再開を確認するには、タスクに NEEDS_HUMAN を入力します。)
🧠 独自の LLM (BYOK) を持参してください。 Coasty ではなくあなたのモデルで実行したいですか?
[設定] → [モデル プロバイダー] を開き、OpenRouter、OpenAI、またはローカル モデルを追加します。
（オラマ/LMスタジオ）。 Coasty は、切り替えるまでデフォルトのままです。
BYOKへジャンプ ↓
⚠️ ローカルコントロールは実際のマウスとキーボードを動かします。キャンセルで停止（または
ウィンドウを閉じてください)、小さなことから始めてください。安全に関する完全な注意事項は RUN_LOCALLY.md にあります。
目標
どのように
モデル
コスト
自分の PC を自動化する
pnpmデスクトップ
Coasty または独自の LLM
デモ $0 · BYOK = プロバイダーの料金
ウェブアプリのみ
pnpm 開発 → http://127.0.0.1:5173
海岸沿い
$0
Coasty アカウント
COASTY_API_KEYを.envに追加
コースティ（実物モデル）
サンドボックスキー = $0
独自の LLM の持ち込み (BYOK)
設定 → モデルプロバイダー
OpenRouter · OpenAI · Ollama
プロバイダーの料金 · ローカル = $0
設定する必要があるのは COASTY_API_KEY だけです。これもデモではオプションです
モード。それ以外はすべてデフォルトで動作します。独自のモデルをご希望ですか?それは BYOK です — 選択してください
設定のプロバイダーとローカル実行でそれが使用されます。完全なローカル自動化ガイド:
RUN_LOCALLY.md 。
echo " COASTY_API_KEY=sk-coasty-test-… " > .env # サンドボックス キー — 請求は行われません
pnpm dev # は実際の Coasty API と通信するようになりました
キーが設定されている場合、pnpm dev はモックを開始せず、バックエンドを次のようにポイントします。
本当のC

オースティーAPI。サンドボックス キー ( sk-coasty-test-… ) から始めます。
完全な実際のモデルを実行し、請求は一切発生しません。次の場合にのみライブキーに切り替えます。
過ごす準備はできています。
Webhook (ポーリングなしのインスタント ステータス) には https が必要です
COWORK_PUBLIC_URL — Coasty は HTTPS Webhook URL のみを受け入れます。オープンコワーク
これを検出します: https 以外の URL では単に Webhook を登録しません (したがって、
作成は決して失敗しません)、状態は SSE と読み取り時の調整を介してライブで同期されます。
https COWORK_PUBLIC_URL を設定します (トンネルまたはデプロイメント - を参照)
DEPLOYMENT.md ) を使用して Webhook をオンにします。
⚠ コストに関する警告。ライブキー ( sk-coasty-live-… ) を使用すると、請求書が実行されます。
$0.05/ステップ、マシン $0.05 ～ 0.09/時間実行 (停止 $0.01)、
予測/セッションコールはそれぞれ数セントです。 open-cowork では常に見積もりが表示されます。
明示的な確認が必要であり、実行ごとの予算上限をサーバー側で強制します。
マシンの自動終了 TTL をサポートしていますが、ライブ キーは現金です。すべて
自動テストではモック/サンドボックス パスが使用され、費用は一切かかりません。
pnpm デスクトップ # フルスタック + デスクトップ アプリ (ローカル画面コントロール) — 1 つのコマンド
pnpm dev # フルスタック、Web アプリを自分で開きます (:5173)
pnpm dev --no-web # API のみ (モック + バックエンド)
pnpm dev:mobile # Expo / React Native (または: pnpm --filter @open-cowork/mobile web)
すでに実行されているスタックに対して Electron アプリを実行する (上級)
pnpm dev が別の端末ですでに実行されている場合:
pnpm dev:desktop # Electron バンドルをビルドし、ウィンドウのみを開きます
pnpm デスクトップは、スタックの開始とウィンドウの開きの両方を行います。
窓を閉めるとすべてシャットダウンします。
💬 チャットでの委任 — 「これらのファイルの名前を変更してレポートをメールで送信」 — そして
ライブ画面ビューでエージェントがステップごとに実行する様子を確認します。
🧠 Bring your own LLM (BYOK) — モデル上でローカル画面コントロールを実行します。
オープンR

外部モデル、OpenAI、またはローカル モデル (Ollama / LM Studio / vLLM)。コースティはただ
デフォルト。詳細↓
📺 実行の監視 — ダッシュボード、耐久性のあるイベント タイムライン (リプレイ付き SSE)、
Web、デスクトップ、または電話からのキャンセル/再開/人による引き継ぎ。
🔁 ワークフローの構築 — バージョン管理された JSON DSL (タスク、アサート、if、ループ、
並列、再試行、human_approval)、即時検証、コスト見積もり、
そしてサーバー側の厳しい予算上限。
🖥️ マシンの管理 — Coasty クラウド VM のプロビジョニング、スナップショット、停止、
すべてのステップでライブコストレートを使用して終了します。
📱 デバイス間で最新情報を常に把握 — ラップトップでランニングを開始します。それがいつ
承認のために一時停止すると、携帯電話にバナーが表示されます。そこで承認してください。
💸 コストを常に確認 — ウォレットの残高、実行ごとの最悪の場合の推定、
そして、請求可能なものが開始される前に、明示的なコスト確認ハンドシェイクが行われます。
能力
🖥️ デスクトップ
🌐 ウェブ
📱モバイル
ローカル画面制御
✅ ファーストクラス
→ クラウドマシン
→ クラウドマシン
クラウドマシンコントロール+ライブビュー
✅
✅
✅
タスクチャット + 実行ダッシュボード
✅
✅
✅
ワークフロービルダー
✅いっぱい
✅いっぱい
閲覧+承認
承認/人間の乗っ取り
✅
✅
✅
コスト/ウォレットビュー
✅
✅
✅
独自の LLM の持ち込み (BYOK)
自分のキー、自分のモデルを持参してください。ローカル画面コントロールのデフォルトは Coasty のものです
コンピューター使用モデル — ただし、代わりに任意の OpenAI 方言 LLM を指すこともできます。で
デスクトップ アプリで、[設定] → [モデル プロバイダー] を開き、プロバイダーを選択し、ビジョン対応を選択します。
モデルであり、ローカル実行ではそれが使用されます。 Coasty がデフォルトのままです。ワンクリックでいつでも元に戻すことができます。
アプリ内では他に何も変更されません。
👁️ビジョンが必要です。パソコンの使用はスクリーンショット中心なので、見えないモデル
画像にはフラグが付けられ、明確なメッセージでブロックされます。決して盲目的で無駄な実行ではありません。
🏠 地元第一主義。ローカル モデル (例: Ollama

t http://localhost:11434/v1 ) 完全に実行されます
あなたのマシン上で — キーもクラウドも費用もかかりません。
🔑 あなたの鍵はあなたのもののままです。 BYO キーは OS キーチェーン (Electron) で暗号化されます。
safeStorage — DPAPI / Keychain / libsecret) は、デスクトップ プロセス内でのみ存在します。
すべてのエラー メッセージからスクラブされ、Web バンドルやモバイル バンドルには決して到達しません。
👀 予期せぬデータの流出はありません。サードパーティ モデルの場合、スクリーンショットとプロンプトは次の場所に移動します。
そのプロバイダー — アプリは、実行開始前にコストの確認ダイアログでその旨を表示します。
☁️ クラウドマシンの実行では常に Coasty を使用します。現在、BYO はローカル (デスクトップ) での実行を推進しています。雲
BYO は文書化されたフォローアップです。
Vercel AI SDK に基づいて構築: レート制限 (429) および一時的なエラー
バックオフを使用して再試行し、モデルが構造化出力を無視した場合、応答は次のように回復されます。
防御的な JSON 解析 - そのため、より小規模なローカル モデルでもループを駆動できます。
あなた ──► オープンコワーク バックエンド ──► Coasty API ──► エージェントが操作する画面
│ (唯一の場所 §─ 自分のデスクトップ (デスクトップ アプリ)
│ API キーは存続します) §─ Coasty クラウド VM (任意のクライアント)
└─► ウェブ / デスクトップ / モバイル └─ ブラウザ ページ (Playwright)
ライブイベント、承認、コスト
1 つの共有エージェント ループ (スクリーンショット → 予測 → 実行 → 繰り返し) により、あらゆる操作が実行されます。
単一の Executor インターフェイス — LocalExecutor (デスクトップ) を介して画面に表示されます。
RemoteMachineExecutor (クラウド VM)、または BrowserExecutor 。予測ステップ
独自のシーム ( @open-cowork/llm ): Coasty がデフォルトの実装であり、
独自の LLM を使用することは、同じ契約に基づくもう 1 つの LLM にすぎません。
したがって、ループ、エグゼキュータ、および UI は、その背後にあるものを気にしません。クライアントは決して保留しない
Coasty キー: 有効期間の短いセッション トークンを使用してバックエンドと通信します。
バックエンドは Coasty にプロキシし、HMAC 署名付き Webhook を検証し、実行を永続化し、

SSE 上でファン イベントが開催されます。フルデザインで
アーキテクチャ.md 。
COASTY_API_KEY はバックエンドの環境にのみ存在します。ブラウザ、
Electron レンダラーとモバイル アプリは有効期間の短いセッションで認証します
トークンを使用し、キーは決して表示されません - すべてのクライアント バンドルをスキャンするテストによって実施されます
およびランタイム E2E アサーションは、シークレットに対するすべてのブラウザー要求を監視します。
素材。 Bring-your-own-LLM キーは、OS で暗号化されるという同じルールに従います。
キーチェーン (safeStorage)、デスクトップ プロセス内でのみ保持され、すべてのプロセスからスクラブされます
エラーが発生し、Web/モバイル バンドルから除外されます (AI SDK はデスクトップのみです)。
Coasty Webhook は、実行ごとの HMAC シークレット (定数時間の比較、
±5 分のリプレイ ウィンドウ)、ステートに到達する前に。の脅威ノート
セキュリティ.md 。
ガイド
中には何が入っているのか
RUN_LOCALLY.md
デスクトップ アプリを使用して自分の PC を自動化する — 段階的に
アーキテクチャ.md
Monorepo マップ、Executor 抽象化、エージェント ループ、リアルタイム + データ モデル
セキュリティ.md
キー保管場所、HMAC、信頼境界、脅威テーブル
展開.md
バックエンドと各クライアントを本番環境で実行する
クックブック.md
レシピ: クロスデバイス承認、ワークフロー、ループのスクリプト作成
決定.md · 貢献.md · 概要.md
スタックの選択 · 貢献方法 · 構築されたもの + カバレッジ
プロジェクトのレイアウト
パッケージ/コア Coasty クライアント、エージェント ループ、ワークフロー DSL、コスト推定、HMAC — 同型、ゼロ deps
パッケージ/executor エグゼキュータの抽象化: LocalExecutor (ネイティブ)、RemoteMachineExecutor (VM)、BrowserExecutor
パッケージ/

[切り捨てられた]

## Original Extract

Open-source alternative to Claude Co-Work built by Claude Mythos 5. with BYOK - coasty-ai/open-cowork

GitHub - coasty-ai/open-cowork: Open-source alternative to Claude Co-Work built by Claude Mythos 5. with BYOK · GitHub
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
coasty-ai
/
open-cowork
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
72 Commits 72 Commits .github/ workflows .github/ workflows apps apps e2e e2e packages packages public public scripts scripts tools/ mock-coasty tools/ mock-coasty .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .prettierignore .prettierignore .prettierrc.json .prettierrc.json ARCHITECTURE.md ARCHITECTURE.md CONTRIBUTING.md CONTRIBUTING.md COOKBOOK.md COOKBOOK.md DECISIONS.md DECISIONS.md DEPLOYMENT.md DEPLOYMENT.md DESIGN_SYSTEM.md DESIGN_SYSTEM.md LICENSE LICENSE PLAN.md PLAN.md README.md README.md RUN_LOCALLY.md RUN_LOCALLY.md SECURITY.md SECURITY.md SUMMARY.md SUMMARY.md UI_AUDIT.md UI_AUDIT.md eslint.config.mjs eslint.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json turbo.json turbo.json View all files Repository files navigation
Hand off computer tasks to an AI coworker — watch it work, approve from anywhere.
An open-source, cross-platform agentic coworker that sees a screen and acts on it —
your own desktop, a cloud VM, or a browser. It streams every step live, pauses for your
approval, and keeps spend visible and capped.
Runs on the Coasty Computer Use API out of the box —
or bring your own LLM (OpenRouter · OpenAI · a local model). Your call.
Quickstart ·
Bring your own LLM ·
Automate your PC ·
Features ·
How it works ·
Docs
Delegate a task → watch your coworker drive a browser, step by step → get the result. Runs with zero setup in demo mode .
Prereqs: Node ≥ 22.5 (we use 24) · pnpm 10 ( corepack enable ).
git clone https://github.com/coasty-ai/open-cowork.git && cd open-cowork
pnpm install # one install for the whole monorepo
pnpm desktop # ← runs the desktop app: starts backend + web, opens the window
That's it — one command, zero config. pnpm desktop starts the backend and web
UI, then opens the desktop app (the build that can drive your own screen ).
With no key set it runs in demo mode : a bundled mock server and a throwaway
sandbox key — no account, no network, no billing.
On Delegate , pick “This computer (local screen).”
Type a task → confirm the cost → watch it work. (Tip: put NEEDS_HUMAN in a task to see the approval flow pause and resume.)
🧠 Bring your own LLM (BYOK). Want it to run on your model instead of Coasty?
Open Settings → Model provider and add OpenRouter, OpenAI, or a local model
(Ollama / LM Studio). Coasty stays the default until you switch.
Jump to BYOK ↓
⚠️ Local control moves your real mouse and keyboard. Stop with Cancel (or
close the window), and start small — full safety notes in RUN_LOCALLY.md .
Goal
How
Model
Cost
Automate your own PC
pnpm desktop
Coasty or your own LLM
demo $0 · BYOK = your provider's rate
Web app only
pnpm dev → http://127.0.0.1:5173
Coasty
$0
Your Coasty account
add COASTY_API_KEY to .env
Coasty (real model)
sandbox key = $0
Bring your own LLM (BYOK)
Settings → Model provider
OpenRouter · OpenAI · Ollama
your provider's rate · local = $0
The only thing you ever have to set is COASTY_API_KEY — and even that's optional in demo
mode. Everything else has a working default. Prefer your own model? That's BYOK — pick a
provider in Settings and local runs use it. Full local-automation guide:
RUN_LOCALLY.md .
echo " COASTY_API_KEY=sk-coasty-test-… " > .env # sandbox key — never bills
pnpm dev # now talks to the real Coasty API
With a key set, pnpm dev does not start the mock and points the backend at
the real Coasty API. Start with a sandbox key ( sk-coasty-test-… ) — it
exercises the full real model and never bills. Switch to a live key only when
you're ready to spend.
Webhooks (instant status without polling) require an https
COWORK_PUBLIC_URL — Coasty only accepts HTTPS webhook URLs. open-cowork
detects this: over a non-https URL it simply doesn't register a webhook (so run
creation never fails) and state still syncs live via SSE + read-time reconcile.
Set an https COWORK_PUBLIC_URL (a tunnel or your deployment — see
DEPLOYMENT.md ) to turn webhooks on.
⚠ Cost warning. With a live key ( sk-coasty-live-… ): runs bill
$0.05/step , machines $0.05–0.09/hour running ($0.01 stopped),
predict/session calls a few cents each. open-cowork always shows an estimate,
requires explicit confirmation, enforces per-run budget caps server-side, and
supports machine auto-terminate TTLs — but a live key is real money. All
automated tests use the mock/sandbox path and never spend anything.
pnpm desktop # full stack + the desktop app (local screen control) — one command
pnpm dev # full stack, open the web app yourself at :5173
pnpm dev --no-web # API only (mock + backend)
pnpm dev:mobile # Expo / React Native (or: pnpm --filter @open-cowork/mobile web)
Run the Electron app against an already-running stack (advanced)
With pnpm dev already running in another terminal:
pnpm dev:desktop # builds the Electron bundles and opens the window only
pnpm desktop does both for you — start the stack and open the window — and
shuts it all down when you close the window.
💬 Delegate in chat — "rename these files and email the report" — and
watch the agent execute it step by step with a live screen view.
🧠 Bring your own LLM (BYOK) — run local screen control on your model:
OpenRouter, OpenAI, or a local model (Ollama / LM Studio / vLLM). Coasty is just
the default. Details ↓
📺 Supervise runs — dashboard, durable event timeline (SSE with replay),
cancel / resume / human-takeover from web, desktop, or phone.
🔁 Build workflows — a versioned JSON DSL (task · assert · if · loop ·
parallel · retry · human_approval) with instant validation, cost estimates,
and hard server-side budget caps.
🖥️ Manage machines — provision Coasty cloud VMs, snapshot, stop,
terminate, with live cost rates at every step.
📱 Stay in the loop across devices — start a run on your laptop; when it
pauses for approval, the banner pops on your phone. Approve there.
💸 See cost at all times — wallet balance, per-run worst-case estimates,
and an explicit confirm-the-cost handshake before anything billable starts.
Capability
🖥️ Desktop
🌐 Web
📱 Mobile
Local screen control
✅ first-class
→ cloud machine
→ cloud machine
Cloud-machine control + live view
✅
✅
✅
Task chat + run dashboard
✅
✅
✅
Workflow builder
✅ full
✅ full
view + approve
Approvals / human takeover
✅
✅
✅
Cost / wallet view
✅
✅
✅
Bring your own LLM (BYOK)
Bring your own key, bring your own model. Local screen control defaults to Coasty's
computer-use model — but you can point it at any OpenAI-dialect LLM instead. In the
desktop app, open Settings → Model provider , pick a provider, choose a vision-capable
model, and local runs use it. Coasty stays the default; switch back any time with one click.
Nothing else in the app changes.
👁️ Vision is required. Computer use is screenshot-driven, so a model that can't see
images is flagged and blocked with a clear message — never a blind, wasted run.
🏠 Local-first. A local model (e.g. Ollama at http://localhost:11434/v1 ) runs entirely
on your machine — no key, no cloud, no spend.
🔑 Your key stays yours. BYO keys are encrypted with your OS keychain (Electron
safeStorage — DPAPI / Keychain / libsecret), live only in the desktop process, are
scrubbed from every error message, and never reach the web or mobile bundle.
👀 No surprise data egress. With a third-party model, your screenshots and prompts go to
that provider — the app says so right in the confirm-the-cost dialog before a run starts.
☁️ Cloud-machine runs always use Coasty. BYO drives local (desktop) runs today; cloud
BYO is a documented follow-up.
Built on the Vercel AI SDK : rate-limit (429) and transient errors
retry with backoff, and if a model ignores structured output the response is recovered with
a defensive JSON parse — so even smaller local models can drive the loop.
You ──► open-cowork backend ──► Coasty API ──► a screen the agent drives
│ (the ONLY place ├─ your own desktop (desktop app)
│ the API key lives) ├─ a Coasty cloud VM (any client)
└──► web / desktop / mobile └─ a browser page (Playwright)
live events, approvals, costs
One shared agent loop (screenshot → predict → act → repeat) drives any
screen through a single Executor interface — LocalExecutor (your desktop),
RemoteMachineExecutor (a cloud VM), or BrowserExecutor . The predict step
is its own seam ( @open-cowork/llm ): Coasty is the default implementation, and a
bring your own LLM is just another one behind the same contract,
so the loop, executors, and UI don't care which is behind it. Clients never hold
the Coasty key: they talk to the backend with short-lived session tokens, and
the backend proxies to Coasty, verifies HMAC-signed webhooks, persists runs, and
fans events out over SSE. Full design in
ARCHITECTURE.md .
COASTY_API_KEY exists only in the backend's environment. Browsers,
Electron renderers, and the mobile app authenticate with short-lived session
tokens and never see the key — enforced by tests that scan every client bundle
and a runtime E2E assertion that watches every browser request for secret
material. Bring-your-own-LLM keys follow the same rule: encrypted with the OS
keychain ( safeStorage ), held only in the desktop process, scrubbed from every
error, and kept out of the web/mobile bundles (the AI SDK is desktop-only).
Coasty webhooks are verified with per-run HMAC secrets (constant-time compare,
±5-minute replay window) before they can touch any state. Threat notes in
SECURITY.md .
Guide
What's inside
RUN_LOCALLY.md
Automate your own PC with the desktop app — step by step
ARCHITECTURE.md
Monorepo map, the Executor abstraction, agent loop, realtime + data model
SECURITY.md
Key custody, HMAC, trust boundary, threat table
DEPLOYMENT.md
Running the backend + each client in production
COOKBOOK.md
Recipes: cross-device approval, workflows, scripting the loop
DECISIONS.md · CONTRIBUTING.md · SUMMARY.md
Stack choices · how to contribute · what was built + coverage
Project layout
packages/core Coasty client, agent loop, workflow DSL, cost estimator, HMAC — isomorphic, zero deps
packages/executor Executor abstraction: LocalExecutor (native), RemoteMachineExecutor (VM), BrowserExecutor
packages/

[truncated]
