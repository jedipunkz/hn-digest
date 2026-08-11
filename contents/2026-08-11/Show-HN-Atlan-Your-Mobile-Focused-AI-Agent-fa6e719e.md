---
source: "https://github.com/Atlansdaddy/atlan"
hn_url: "https://news.ycombinator.com/item?id=49259343"
title: "Show HN: Atlan, Your Mobile-Focused AI Agent"
article_title: "GitHub - Atlansdaddy/atlan: An AI coding cockpit that runs on your phone, not a server you remote into. Agents, live preview, editor, terminal and APK builds, all on-device. · GitHub"
author: "Emirhan123"
captured_at: "2026-08-11T15:52:24Z"
capture_tool: "hn-digest"
hn_id: 49259343
score: 1
comments: 1
posted_at: "2026-08-11T14:50:55Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Atlan, Your Mobile-Focused AI Agent

- HN: [49259343](https://news.ycombinator.com/item?id=49259343)
- Source: [github.com](https://github.com/Atlansdaddy/atlan)
- Score: 1
- Comments: 1
- Posted: 2026-08-11T14:50:55Z

## Translation

タイトル: 番組 HN: Atlan、モバイル中心の AI エージェント
記事のタイトル: GitHub - Atlansdaddy/atlan: リモート接続するサーバーではなく、携帯電話上で実行される AI コーディング コックピット。エージェント、ライブ プレビュー、エディター、ターミナル、APK ビルドはすべてデバイス上で実行できます。 · GitHub
説明: リモート接続したサーバーではなく、携帯電話上で実行される AI コーディング コックピット。エージェント、ライブ プレビュー、エディター、ターミナル、APK ビルドはすべてデバイス上で実行できます。 - アトランズダディ/アトラン

記事本文:
GitHub - Atlansdaddy/atlan: リモート接続するサーバーではなく、携帯電話上で実行される AI コーディング コックピット。エージェント、ライブ プレビュー、エディター、ターミナル、APK ビルドはすべてデバイス上で実行できます。 · GitHub
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
アトランズダディ
/
アトラン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
245 コミット 245 コミット .github/ workflows .github/ workflows bin bin docs docs scripts scripts サーバー サーバー テスト テスト VA

ult vault web/ public web/ public workworker .gitignore .gitignore .npmrc .npmrc ライセンス ライセンス通知 README.md README.md atlan.config.example.json atlan.config.example.json eslint.config.mjs eslint.config.mjs package-lock.json package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
電話を完全な開発ワークステーションに変える AI ネイティブのソフトウェア エンジニアリング コックピット。 LLM をラップする代わりに、Atlan はコード編集、ターミナル アクセス、ローカルおよびクラウド モデル、決定論的検証、ビルド自動化、およびエージェント オーケストレーションを単一の実行環境 ( http://127.0.0.1:4589 の 1 ページ) に統合します。参照プラットフォームは、電話機上の Termux 上の Ubuntu ルートです。また、任意の Linux/macOS ホストでも実行できます ( docs/SETUP.md を参照)。コックピットはループバックにのみバインドされます。意図的にトンネルを構築しない限り、ネットワークからコックピットに到達することはありません。選択したエンジンがプロバイダーに送信するものは別の質問であり、 docs/SECURITY.md で完全に回答されます。
John Viruet / 中部大西洋 AI によって構築されました。ライセンス付きの Apache-2.0 — 使用およびフォークは無料です。帰属を保持してください。常駐 AI である Atlan は、コックピットの生きたマスコットであり、構築中に実際に起こっていることに反応する穏やかな存在です。
ステータス (2026-08-08): ウォームで高速なセッションでの M1 ～ M6 とストリーミング チャット (CLI は一度生成され、維持されます。ターンごとに ~3.7 秒のコールド スタートはありません。最初のトークンのウォームまでは ~1.3 秒、コールドまでは ~7 秒と測定)、ライブ自己認識 (Atlan は現在時刻、アクティブなタブ、実行中のエージェント、今日のトークンの書き込み、開いているプロジェクトを認識します)、目に見える推論 (要約された思考の流れが、パネル)、パスワード認証、ワーカー階層、添付ファイル、コード エディター、音声 I/O (12 個の AI モデル + 10 個の音声プロバイダー - ブラウザ音声はデフォルトで無料、残りは

BYO キー、OpenAI-Realtime は正直にマークされたロードマップ)、およびオプションの Termux:Boot reboot-autostart を備えた耐久性の高い自動リスポーン サーバー ( bin/atlan-serve.sh )。 25 のスイート全体で 838 の自動テストがグリーン化されています ( docs/RECEIPTS.md を参照)。設計によりループバックのみで実行されます。パスワードを設定すると、プリフライト セキュリティ ゲートが緑色に変わります (これは初回実行の一部です)。
このパスの新機能: 検索可能な履歴と削除されないアーカイブにより、更新後も会話が維持されます。 · レート、重複、バックログ、およびリレー深さの制限によって制限されるチャット間メッセージおよびチャット対プロジェクト メッセージ · プレビュー全画面切り替え · 各チェックの回答を質問ごとにグループ化したドクター、ワンタップのコピー レポートと実際のエージェント - CLI 接続チェック · 平文からフォームに記入するペルソナ + ドラフター · および最後にベンダー名が「選択したエンジン」、つまりワイヤ プロトコル、実行時の決定、コピーを表します。
実際の G​​alaxy S24 Ultra (Android 16、カーネル 6.1.145-android14、arm64) で測定した閉じ込め: 14/16 段、T2 を確立し、Landlock と兄弟メモリによってのみ T3 に達しません。これは、デバイスの出荷時に GKI カーネル ブランチがフリーズするため、OS アップデートによって追加されないためです。宣言されたTier() は依然としてすべてのホストでデフォルトの T0 です。これらの数値は adb シェルから取得され、Atlan は Android 独自の zygote seccomp フィルターの下で untrusted_app として実行されますが、これは機能を奪うことしかできません。 docs/SECURITY.md を参照してください。
通常、電話機上に構築すると、端末が窮屈になり、フィードバック ループがなくなります。 Atlan は電話を真の開発画面にします。コードを編集するエージェントと会話すると、変更中のアプリが表示され、そのエラーは自動的にエージェントに返され、会話をコマンド ラインに渡して戻すことができ、寝ている間にエージェントを予算の仕事に送り出すことができ、次のようなことができます。

同じ画面からインストール可能な APK をビルドします。それは正直さを重視しています。すべての機能には実際の内容がラベル付けされており、危険なものにはすべて目に見える壁があります。
Atlan は、ブラウザから操作するノード サーバーです。サーバーを実行できる場所によって、サーバーの使用方法が決まります。
APK は Android 上の Termux 要件を削除しません。ワンボタン APK はラッパー/クライアントです。ノードサーバーには Linux ユーザーランドが必要であり、Termux/proot がそのユーザーランドであるため、サーバーは引き続き Termux/proot で実行されます。広く共有可能なパスはクラウド クライアント モードです (サーバーを一度ホストすれば、任意の電話に接続できます)。電話機のローカル Termux ルートはプライバシー/パワーユーザー モードです。両方については docs/SETUP.md を参照してください。
エンジンも必要です。 Atlan はモデルを駆動するコックピットです。少なくとも 1 つが必要です。Claude サブスクリプション/API キー (フル エージェント、「手を使う」)、Gemini/Groq のような無料のクラウド キー (チャット頭脳、手を使わない)、またはローカルのラマ サーバー モデル (ハードウェアが実行できる場合は無料)。何も設定されていない場合、アプリは実行されますが、エージェントには頭脳がありません。
機能ごとの 7 つのタブ
プロジェクト ピッカー — プロジェクト ディレクトリ内のすべてのフォルダー (構成可能、デフォルトは /root ) に .git または package.json が含まれます。他のすべて (フリート、ビルド、ターミナル、エディター) は、選択したプロジェクト内で動作します。
ストリーミング + 思考 + 自己認識 — 応答はウォーム セッションでトークンごとにストリームされます (エージェント CLI は一度生成され、存続するため、ターンごとのコールド スタート待機はありません)。折りたたみ可能なパネルには、モデルの要約された推論がライブで表示されます。送信するとすぐに「動作中…」インジケーターが点灯します。 Atlan はライブ状態を認識しています。各ターンには現在時刻、アクティブなタブ、実行中のフリート エージェント、今日のトークンの燃焼、開いているプロジェクトが静かに記録されるため、作業中に何が起こっているかを真に把握します。
添付ファイル — 📎 i

mages/audio/video/files/folders (ドラッグ、ペースト、またはパスの参照)。画像はビジョンに送信され、ファイル/フォルダーはエージェントが読み取るパス参照になり、オーディオ/ビデオはマルチモーダル モデル (Gemini/OpenAI) にルーティングされ、テキストとしてターンに組み込まれます。
エンジンスイッチャー、4 つの正直なグループ:
クロード・コード (エージェント) — fable-5 、 opus-4.8 、 Sonnet-5 、 haiku-4.5 。実際の手: ファイルの読み取り/編集、ツールの実行、ビルド。許可カード付き。
エージェント CLI - Codex および Gemini CLI。リポジトリ内でヘッドレスの全自動で実行されます (全か無かの承認。Claude は慎重なカードゲート型プライマリのままです)。
電話（無料） — ラマサーバー経由のローカルモデル（チャットのみ）。
クラウド ブレイン — 1 つの OpenAI 互換アダプターを介して広範な BYO キーが展開されます (チャットのみ。手がないことがわかります): Gemini、OpenAI、DeepSeek、Kimi (Moonshot)、xAI Grok、Mistral、Groq、Togetter、OpenRouter、Fireworks、Cohere 。別のプロバイダーを追加するには、単一のベース URL 行を追加します。各defaultModelは単なる出発点であり、プロバイダーが提供する任意のモデルを入力します。
利用できないオプションは無効になり、必要なものが正確に表示されます。設定 → エンジン キーには、すべてのキーに「↗ の入手方法」リンクがあります。
許可カード — クロードが危険な行動をしたい場合、許可/拒否を取得します。拒否は常に安全です。エージェントは理由を告げられ、適応します。
会話は継続します。スレッドを破棄するために使用されるプル・トゥ・リフレッシュ - チャットは DOM 内と、サーバーが閉じるときにクリアするソケットごとのマップ内にのみ存在します。トランスクリプトは耐久性が高くなったので ( FLEET_DIR/chats/*.jsonl 、一度に 1 行ずつ追加されるため、長いチャットが携帯電話のフラッシュに O(n²) 回書き込むことはありません)、リロードすると元の場所からリプレイされます。 🗂 履歴には過去の会話がリストされます。検索可能で、アシスタントのツールの前文ではなく最初のメッセージによってタイトルが付けられます。
アーカイブし、決して剪定しない。スペースを空けるために何も削除されません。使用量は測定されます - サイズ、無料

ディスク、利用可能な RAM — そして重要な場合、医師はそう言って尋ねます。アーカイブされた会話は同じリストに残り、マークが付けられ、タップすると開きます。それらが作成した gzip は、決して目にすることのない実装の詳細です。
チャットからチャットおよびチャットからプロジェクトへのメッセージ。履歴リストから別の会話またはプロジェクト (「認証に取り組んでいる人に通知する」) にメモを送信します。ライブの会話は画面上に表示されます。休止中のものは、あなたがそれを開くまで、トランスクリプトにそれを保持します。すべてのピア メッセージは独自のスタイルと「from」署名欄でレンダリングされます。エージェントはここで全自動で実行されるため、エージェント間の属性のないチャネルが注入パスとなり、属性が制御となります。レート、重複、バックログ、およびリレー深さの制限がそれを制限します。
セッションハンドオフ — 各ターンの後に、行にコスト + セッション ID が表示されます。タップして claude --resume <id> をコピーし、どの端末でも同じ会話を続けます。
自動添付されたプレビュー コンテキスト — コンソール エラーと [プレビュー] タブの 📸 スナップショットは、次のメッセージに自動的に反映されます。
音声 — アトランに話しかけ、彼の返事を聞いてください。 🎤 プッシュ トゥ トークでは、ブラウザの Web Speech API (無料、デバイス上) を使用し、トランスクリプトをボックスにドロップして送信前に確認します。 🔈/🔊 音声応答を切り替えます。 Speech-out は幅広く、正直な BYO キーのスプレッドです — [設定] → [音声] で 1 つを選択してください:
ブラウザ (無料、インスタント、オフライン、品質はさまざま) · Piper (無料、ローカル、プライベート、本物の SSML — pip install Piper-tts + .onnx 音声)
イレブンラボ、Cartesia Sonic、Deepgram Aura-2、OpenAI TTS (低遅延プレミアム) · Google Cloud TTS、Azure Speech、Amazon Polly (幅広い言語、リアル SSML、予算)
各プロバイダーには、コスト、遅延、 SSML を尊重するかどうかが表示され、キーが設定されるまでグレー表示されます。ピッカーは、使用できない音声を要求することはありません。気分 (穏やか/誇り/警戒/構築) dri

韻律が軽いので、アトランはオーブの見た目のように聞こえます。 OpenAI Realtime (全二重音声対音声) はロードマップ上にあり、そのようにラベル付けされていますが、機能するふりをしているわけではありません。
▣ プレビュー — 構築しているアプリを確認します
ウォッチャーを挿入するプロキシを通じてレンダリングされる、任意のローカル開発サーバー (ループバックのみ — 127.0.0.1 / localhost / ::1 、ループバック許可リストと完全に一致するホスト名、意図的な SSRF 境界) を指します。
コンソール ストリップはアプリのログをミラーリングします。エラーはキューに入れられ、file:line を使用してエージェントの次のターンに自動的に付加されます。スタック トレースを再度コピーして貼り付ける必要はありません。
📸 スナップショットは、視覚能力のあるエージェントが読み取った実際の PNG を保存します (「ボタンがヘッダーに重なっている」ことが検証可能になります)。
⤢ ディスプレイが全画面表示になり、Escape で元に戻ります。これは、最初に CSS クラス、次にフルスクリーン API です。その API は iOS Safari と一部の Web ビュー内で拒否されるためです。プライマリ プラットフォームで何も行わないコントロールは、何も行わないよりも悪いです。
HMR / ライブリロードはそのまま通過します。
✎ エディタ — コードを手動で作成またはレビューします
完全なコード エディター (CodeMirror、122 言語、自己ホスト型 - CDN なし)。パスでファイルを開くか、プロジェクト ツリーを参照し、構文を強調表示して編集し、ディスクに保存します。
レビューのためにエージェントに送信すると、開いているファイルがレビュー プロンプトとともにチャット ターンに渡されます。スコープはプロジェクトに限定されます。認証情報のパス ( .ssh

[切り捨てられた]

## Original Extract

An AI coding cockpit that runs on your phone, not a server you remote into. Agents, live preview, editor, terminal and APK builds, all on-device. - Atlansdaddy/atlan

GitHub - Atlansdaddy/atlan: An AI coding cockpit that runs on your phone, not a server you remote into. Agents, live preview, editor, terminal and APK builds, all on-device. · GitHub
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
Atlansdaddy
/
atlan
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
245 Commits 245 Commits .github/ workflows .github/ workflows bin bin docs docs scripts scripts server server test test vault vault web/ public web/ public worker worker .gitignore .gitignore .npmrc .npmrc LICENSE LICENSE NOTICE NOTICE README.md README.md atlan.config.example.json atlan.config.example.json eslint.config.mjs eslint.config.mjs package-lock.json package-lock.json package.json package.json View all files Repository files navigation
An AI-native software engineering cockpit that turns a phone into a complete development workstation. Instead of wrapping an LLM, Atlan integrates code editing, terminal access, local and cloud models, deterministic verification, build automation, and agent orchestration into a single execution environment — one page at http://127.0.0.1:4589 . The reference platform is Ubuntu proot under Termux on a phone; it also runs on any Linux/macOS host (see docs/SETUP.md ). The cockpit binds to loopback only — nothing reaches it from the network unless you deliberately build a tunnel. What your chosen engines send to their providers is a separate question, answered in full in docs/SECURITY.md .
Built by John Viruet / Mid-Atlantic AI. Licensed Apache-2.0 — free to use and fork; keep the attribution. Its resident AI, Atlan , is the cockpit's living mascot — a calm presence that reacts to what's actually happening as you build.
Status (2026-08-08): M1–M6 plus streaming chat on a warm, fast session (the CLI is spawned once and kept alive — no ~3.7s cold-start per turn; measured ~1.3s to first token warm vs ~7s cold), live self-awareness (Atlan perceives the current time, your active tab, running agents, today's token burn, and the open project), visible reasoning (summarized thinking streams to a panel), password auth, worker hierarchy, attachments, a code editor, voice I/O ( 12 AI-model + 10 voice providers — browser voice free by default, the rest BYO-key, OpenAI-Realtime honestly marked roadmap), and a durable auto-respawning server ( bin/atlan-serve.sh ) with optional Termux:Boot reboot-autostart. 838 automated tests green across 25 suites (see docs/RECEIPTS.md ). Runs loopback-only by design; the Preflight security gate goes green once you've set a password (it's part of first-run).
New in this pass: conversations persist across a refresh with a searchable history and archiving that never deletes · chat-to-chat and chat-to-project messages , bounded by rate, duplicate, backlog and relay-depth limits · a preview full-screen toggle · a Doctor grouped by the question each check answers , with a one-tap copy-report and a real agent-CLI connection check · a Persona+ drafter that fills the form from a plain sentence · and the removal of the last places a vendor name stood in for "the engine you picked" — the wire protocol, the runtime decisions, and the copy.
Confinement, measured on a real Galaxy S24 Ultra (Android 16, kernel 6.1.145-android14, arm64): 14/16 rungs, establishes T2 , short of T3 only by Landlock and sibling-memory — because the GKI kernel branch freezes when a device ships, so no OS update will add it. declaredTier() still defaults to T0 on every host: those numbers come from an adb shell , and Atlan runs as untrusted_app under Android's own zygote seccomp filter, which can only take capabilities away. See docs/SECURITY.md .
Building on a phone normally means a cramped terminal and no feedback loop. Atlan makes the phone a real dev surface: you talk to an agent that edits your code, you see the app it's changing, its errors flow back to the agent automatically, you can hand any conversation to the command line and back, you can send agents off to work on budgets while you sleep, and you can build an installable APK from the same screen. It's opinionated toward honesty — every capability is labeled for what it actually is, and every dangerous thing has a wall you can see.
Atlan is a Node server you drive from a browser. Where the server can run decides how you use it:
The APK does not remove the Termux requirement on Android — the one-button APK is a wrapper/client ; the server still runs in Termux/proot, because the Node server needs a Linux userland and Termux/proot is that userland. The broadly-shareable path is cloud-client mode (host the server once, connect any phone); the phone-local Termux route is the privacy / power-user mode. See docs/SETUP.md for both.
You also need an engine. Atlan is a cockpit that drives models — it needs at least one: a Claude subscription/API key (full agent, "has hands"), or a free cloud key like Gemini/Groq (chat brains, no hands), or a local llama-server model (free, if the hardware can run it). With none configured, the app runs but the agent has no brain.
The seven tabs, function by function
Project picker — every folder in your projects directory (configurable, defaults to /root ) with a .git or package.json . Everything else (fleet, build, terminal, editor) acts inside the picked project.
Streaming + thinking + self-aware — replies stream token-by-token on a warm session (the agent CLI is spawned once and kept alive, so there's no cold-start wait per turn); a collapsible panel shows the model's summarized reasoning live; a "working…" indicator fires the instant you send. Atlan is self-aware of live state — each turn quietly carries the current time, your active tab, running fleet agents, today's token burn, and the open project, so it genuinely knows what's happening as you work.
Attachments — 📎 images/audio/video/files/folders (drag, paste, or reference a path). Images go to vision, files/folders become path references the agent reads, audio/video are routed to a multimodal model (Gemini/OpenAI) and folded into the turn as text.
Engine switcher , four honest groups:
Claude Code (agent) — fable-5 , opus-4.8 , sonnet-5 , haiku-4.5 . Real hands: reads/edits files, runs tools, builds. Permission-carded.
Agent CLIs — Codex and Gemini CLI, running headless full-auto in your repo (all-or-nothing approvals; Claude stays the careful, card-gated primary).
On-phone (free) — local models via llama-server (chat only).
Cloud brains — a wide, BYO-key spread through one OpenAI-compatible adapter (chat only; they'll tell you they have no hands): Gemini, OpenAI, DeepSeek, Kimi (Moonshot), xAI Grok, Mistral, Groq, Together, OpenRouter, Fireworks, Cohere . Adding another provider is a single base-URL row. Each defaultModel is just a starting point — type any model the provider offers.
Unavailable options are disabled and say exactly what they need; Settings → Engine keys has a "how to get ↗" link for every one.
Permission cards — when Claude wants a risky action you get Allow/Deny. Deny is always safe; the agent is told why and adapts.
Conversations persist. A pull-to-refresh used to destroy the thread — chat lived only in the DOM and in a per-socket map the server clears on close. Transcripts are durable now ( FLEET_DIR/chats/*.jsonl , appended a line at a time so a long chat isn't O(n²) writes on phone flash), so a reload replays where you were. 🗂 History lists past conversations — searchable, titled by your first message rather than the assistant's tool preamble.
Archiving, never pruning. Nothing is deleted to make room. Usage is measured — size, free disk, available RAM — and when it matters the Doctor says so and asks . Archived conversations stay in the same list, marked, and open on tap; the gzip they came out of is an implementation detail you never see.
Chat-to-chat and chat-to-project messages. Send a note to another conversation, or to a project — "tell whoever is working on auth" — from the History list. A live conversation gets it on screen; a dormant one keeps it in its transcript until you open it. Every peer message is rendered with its own style and a "from" byline: agents run full-auto here, so an unattributed channel between them would be an injection path, and the attribution is the control. Rate, duplicate, backlog and relay-depth limits bound it.
Session handoff — after each turn a line shows cost + session id; tap to copy claude --resume <id> and continue the same conversation in any terminal.
Auto-attached preview context — console errors and 📸 snapshots from the Preview tab ride along on your next message automatically.
Voice — talk to Atlan, hear him back. 🎤 push-to-talk uses the browser's Web Speech API (free, on-device) and drops the transcript in the box to review before sending. 🔈/🔊 toggles spoken replies. Speech-out is a wide, honest, BYO-key spread — pick one in Settings → Voice :
Browser (free, instant, offline; quality varies) · Piper (free, local, private, real SSML — pip install piper-tts + a .onnx voice)
ElevenLabs, Cartesia Sonic, Deepgram Aura-2, OpenAI TTS (low-latency premium) · Google Cloud TTS, Azure Speech, Amazon Polly (broad languages, real SSML, budget)
Each provider shows its cost, latency, and whether it honors SSML , and greys out until its key is set — the picker never claims a voice you can't use. Mood (calm/proud/alarmed/building) drives light prosody so Atlan sounds like the orb looks. OpenAI Realtime (full-duplex voice-to-voice) is on the roadmap and labeled as such — not pretended to work.
▣ Preview — see the app you're building
Point at any local dev server (loopback only — 127.0.0.1 / localhost / ::1 , hostname exact-matched against a loopback allowlist; a deliberate SSRF boundary), rendered through a proxy that injects a watcher.
Console strip mirrors the app's logs. Errors queue and auto-attach to the agent's next turn with file:line — you never copy-paste a stack trace again.
📸 Snapshot saves a real PNG a vision-capable agent reads ("the button overlaps the header" becomes verifiable).
⤢ Full screen fills the display and Escape brings it back. It's a CSS class first and the Fullscreen API second, because that API is refused on iOS Safari and inside some webviews — a control that silently does nothing on the primary platform would be worse than none.
HMR / live-reload passes straight through.
✎ Editor — write or review code by hand
A full code editor (CodeMirror, 122 languages, self-hosted — no CDN). Open a file by path or browse the project tree, edit with syntax highlighting, save to disk.
Send to agent for review hands the open file to a chat turn with a review prompt. Scoped to the project; credential paths ( .ssh

[truncated]
