---
source: "https://ai-term.com/"
hn_url: "https://news.ycombinator.com/item?id=48918251"
title: "Show HN: AITerm – a macOS terminal with an AI command loop and a safety gate"
article_title: "AITerm - Native macOS AI Terminal"
author: "dmitrik"
captured_at: "2026-07-15T09:50:49Z"
capture_tool: "hn-digest"
hn_id: 48918251
score: 1
comments: 0
posted_at: "2026-07-15T09:25:13Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AITerm – a macOS terminal with an AI command loop and a safety gate

- HN: [48918251](https://news.ycombinator.com/item?id=48918251)
- Source: [ai-term.com](https://ai-term.com/)
- Score: 1
- Comments: 0
- Posted: 2026-07-15T09:25:13Z

## Translation

タイトル: Show HN: AITerm – AI コマンド ループとセーフティ ゲートを備えた macOS ターミナル
記事のタイトル: AITerm - ネイティブ macOS AI ターミナル
説明: 平易な英語を話すネイティブ macOS ターミナル。 AI が提案したコマンドを承認し、エラーを修正して説明し、ローカルの Ollama、API キー、または Claude / ChatGPT サブスクリプションを使用します。

記事本文:
あなたの端末、
AIでスーパーチャージ
わかりやすい英語を話すネイティブ macOS ターミナル。 AIがコマンドを提案
あなたは、エラーを承認、修正、説明し、複数ステップのタスクを実行します。
独自の AI: ローカル Ollama、API キー、または Claude / ChatGPT サブスクリプション。
またはHomebrewでインストールする
brew install --cask vega-llc/aiterm/ai-term
ダウンロードが開始されませんか?直接リンク
端末に必要なものがすべて揃っています
日常の作業用のネイティブ macOS アプリに加え、コマンドの生成と診断用の無料の CLI。
欲しいものをわかりやすい英語で説明してください。 AITerm は編集可能なシェル コマンドを提案し、承認を待ちます。
失敗した出力に対して /fix および /explain を使用すると、何が起こったのかを修正または平易な英語で読むことができます。
Ollama またはマネージド Apple-Silicon (MLX) エンジンを使用して、完全にオンデバイスで実行します。ローカル リクエストが Mac から送信されることはありません。クラウド リクエストは、最初にシークレットが編集された状態で、中間サーバーなしでプロバイダに直接送信されます。
OpenAI、Anthropic、OpenRouter、Groq、Gemini などの独自のキーを持参してください。 Pro は、公式 CLI を介して Amazon Bedrock、Azure OpenAI、および既存の Claude / ChatGPT サブスクリプションを追加します。
すべてのコマンド パスは、緑色/黄色/赤色のリスク ラベルと監査可能な意思決定ログによる 1 回の安全性チェックを通過します。
コマンド履歴を検索し、モデルを切り替え、Ollama モデルを取得し、コマンド パレットを開き、クリーンアップされた修正カードを共有します。
AITerm — わかりやすい英語を話す端末
AITerm Pro に含まれるネイティブ macOS アプリ: AI が織り込まれた実際の端末エミュレーター
コアループに入る。コマンドを入力するか、必要なものを説明します。AI が提案します。
実行前に確認する編集可能なコマンド。エラーを修正して説明し、実行することができます
マルチステップのタスクを一度に 1 ステップずつ承認し、シーケンスをランブックとして保存して再実行します
後で。 AI 上で実行: ローカル Ollama、API キー、

またはClaude / ChatGPTサブスクリプション。
タブ — 多数のターミナル、1 つのウィンドウ
必要なだけタブを開きます。それぞれが独自の作業ディレクトリ、スクロールバック、AI セッションを備えた独自の実際のシェルです。 ⌘Tで開き、⌘1–9 / ⌘⇧[ ]で切り替えます。各タブのタイトルにはそのフォルダーが表示され、1 つのタブで承認したコマンドはそのタブでのみ実行されます。
タブをサイズ変更可能なペイン (それぞれが独自の実際のシェル) に分割し、作業に合わせてレイアウトします。 ⌘D で右に分割、⌘⇧D で下、⌘⌥矢印でフォーカスを移動します。仕切りをドラッグしてサイズを変更します。すでに入っていたフォルダー内で新しい分割が開き、キーストロークはフォーカスされたペインにのみ移動します。
タブ、分割、作業ディレクトリは、再起動すると、元の状態に戻ります。名前付きレイアウトを保存し、プロジェクト セットアップ全体 (適切なフォルダー内のすべての端末) を 1 つのステップで再度開きます。
エージェント モード — 手動または自動操縦、常にフェンスで囲まれています
目標を与えてください。 /agent を使用すると、段階的に機能します。各コマンドが提案され、ユーザーはすべてのコマンドを承認します。 /auto (オートパイロット) を使用すると、それ自体で実行されますが、安全性ポリシーによって現在のプロジェクトに対して安全であると分類されたコマンドのみが自動実行されます。危険なステップ、破壊的なステップ、ネットワークのステップ、機密性の高いステップ、およびビルド/依存関係のステップは、ワンタップで承認するために一時停止するか、完全にブロックします。どちらも同じ安全ゲートと安全プロファイルを通過します。そのため、ロックダウンされたマシンでは、自動操縦であっても、破壊的なものはすべて完全にブロックされます。
マルチステップ シーケンスを一度保存​​するか、実行時に入力する {{variables}} を使用して ⌘⌥R で実行したコマンドから直接キャプチャし、1 回のフライト前の承認で再生します。破壊的なステップには事前にフラグが付けられ、失敗した場合は /fix を実行して、停止した場所から正確に再試行できます。 {{step1.output}} を使用して、1 つのステップの出力を次のステップにフィードすることでステップを連鎖させます。
新しい v のとき

バージョンの準備が完了すると、ステータス バーに「アップデートの準備完了 — インストールと再起動」という微妙なメッセージが表示されます。セッションを中断するモーダルはなく、何日も開いたままにしたアプリにサイレント バックグラウンド インストールが実行されることもありません。署名され、公証され、小さな差分アップデートとして配信されます。
平易な英語は、緑/黄色/赤のリスク バッジが付いた編集可能なコマンドにストリームされます。実行するには Return を押し、変更するには編集、破棄するには Esc を押します。危険なコマンドはユーザーの確認なしに実行されることはありません。 Autopilot では、安全であると証明されたステップのみが単独で実行されます。
シェル コマンドまたは平易な英語を 1 か所に入力します (モードはありません)。ローカル分類子がそれをルーティングします。 「？」から始まるAI を強制するか、文字通りのコマンドを強制するスペース。
/fix は、失敗したコマンドの正確な出力 (リポジトリ、ブランチ、言語に加えて) を読み取り、元の出力と並べて表示される修正を提案します。ワンクリックで Runbook として保存できます。 /explain (⌘⇧E) は、必要に応じて出力を 1 行ずつ平易な英語で読み上げます。
自由に実行してください。Ollama を使用してローカルおよびプライベートに設定することも、OpenAI、Anthropic、OpenRouter、Groq、Gemini、DeepSeek などの独自のキーを使用することもできます (すべてが Mac 上に残ります)。 Pro は、公式 CLI を介して既存の Claude または ChatGPT サブスクリプションに加え、Amazon Bedrock、Azure OpenAI、およびマネージド オンデバイス Apple-Silicon (MLX) エンジンを接続します。
すべてのコマンド (入力されたコマンド、翻訳されたコマンド、または AI が提案したコマンド) は、最大 30 の回避トリックに対してレッドチームで編成された、単一のタイプ強制安全性チェックに合格します。安全プロファイル (個人→仕事→本番→ロックダウン) を切り替えて、使用しているマシンと同じくらい厳格なゲートを設定します。あらゆる決定の監査ログ。
何か破壊的なものを実行する前に、AITerm は安全に存在する時点に戻る保守的な方法、つまり逆コマンド ( git checkout - 、 mv back、rmdir ) またはバックアップの最初のステップ ( cp -a … .b ) を提供します。

ああ）。偽の「元に戻す」を考え出すのではなく、沈黙を保ちます。
ローカル モデルが Mac から離れることはありません。クラウド リクエストは、最初に編集されたキー、トークン、シークレットとともに、中間サーバーなしでアカウントのプロバイダーに直接送信されます。
完全な PTY ターミナル: vim、top、ssh、colors、Ctrl-C、およびタブ補完はすべて機能します。AI は側面に固定されておらず、上部にあります。
コマンド パレット (⌘K)、Ollama モデルをプルおよび切り替えるためのモデル マネージャー、あいまい履歴検索 (コマンドまたは入力した平易な英語による)、ライブ git ブランチ ステータス、テーマ、共有修正カード、およびオンボーディング。
ネイティブ AI ターミナル — 永久無料。プライベート ローカル AI。初回実行時にセットアップされます。自分のクラウド キーを持参してください。すべてが Mac 上に残ります。 Pro は、Claude / ChatGPT サブスクリプションと自動化を追加します。始めるのにアカウントは必要ありません。
Apple Silicon · macOS 13+ · Appleによる公証 · 自動アップデート
またはHomebrewでインストールする
brew install --cask vega-llc/aiterm/ai-term
タブ、分割、AI ツールを移動します。アプリ内で ⌘/ (または [ヘルプ] ▸ キーボード ショートカット) を押すと、いつでもこれを表示できます。
ネイティブ AI ターミナルは永久に無料です。プライベート ローカル AI、任意のクラウド キー、分割、保存されたワークスペース、完全なセーフティ ゲートを持ち込みます。 Pro は既存の Claude / ChatGPT サブスクリプションをプラグインし、Runbook + エージェント モードを追加します。
ネイティブ macOS アプリ — 実際にコマンドを実行する AI ターミナル。アカウントは必要ありません。
AI コマンド ループ — 提案、承認、実行
/fix & /コマンドまたは出力の説明
タブ、分割ペイン、保存されたワークスペース
セーフティゲート、安全プロファイル、ロールバック計画
プライベート ローカル AI — 最初の実行時にセットアップされ、その後は完全にオフラインになり、キーは不要になります
任意のクラウド キーを持参します — OpenAI、Anthropic、OpenRouter、Groq、DeepSeek、xAI、Mistral、Gemini など
プロキシなし - キーとプロンプトが Mac から離れることはありません
月額わずか 8.25 ドル、毎年請求 — 最大 17% 節約できます。
早美

rd 価格 — 一生閉じ込められます。 14 日間の無料トライアル、カードなし — 初めてアプリを開いたときにトライアルが自動的に開始されます。
AI サブスクリプション — 公式 CLI 経由で Claude Max / ChatGPT Plus を使用し、ローカルで実行します (プロキシなし、リクエストごとの測定なし)
ランブック — マルチステップ シーケンスを保存、パラメータ化、再生します。歴史からのキャプチャ。出力チェイニング
/agent および /auto — 複数ステップのタスクを実行します。各ステップを承認する ( /agent )、または Autopilot に安全なステップを自動実行させます ( /auto ) — どちらも同じ安全ゲートの内側で実行します
マネージド MLX エンジン - Apple シリコンネイティブの 2 番目のオンデバイス エンジンがセットアップされます。
Amazon Bedrock と Azure OpenAI の接続
1座席あたり最大3台のマシン
準備ができていませんか?アプリをダウンロードするだけです。新しくインストールするたびに、カードなしで 14 日間の Pro トライアルが開始されます。
小さなチームですか？一人につき座席を購入してください。 10 シート以上または請求書発行が必要ですか?メールでお問い合わせください。
Pro は、サブスクリプションをアクティブに保つために、月に 1 回程度ライセンス サーバーにチェックインします。
時々オンラインに接続している限り、自動的に表示されます。ライセンスのみが検証されます。
コマンドとファイルが Mac から離れることはありません。 Mac が完全にオフラインになった状態が約時間以上続く場合
6 週間、Pro は一度再接続するまで一時停止します。
違いは、AI が実行される場所です。AITerm はキー、プロンプト、ローカル モデルを 100% Mac 上に保持します。中間にベンダー プロキシはありません。正直なスナップショット、2026 年 7 月 — これらのツールは高速に動作するため、これは現在の状況であり、永続的な堀ではありません。
✓ はい · ~ 部分的 · — なし / 文書化されていません · ✗ いいえ
1. Warp 独自のドキュメントによると、リクエストごとに BYOK プロンプトとキーが Warp のバックエンドを通過します。 Warp のローカル モデル サポートも同様に機能します。モデルはパブリック トンネル URL で公開する必要があり、すべてのプロンプトは引き続き Warp のクラウドを往復して、自分のマシンで実行されているモデルに到達します。 AITerm の BYO キーと場所

すべてのモデルが Mac から離れることはありません。
2. AITerm は、最初の実行時に小さなローカル エンジンとモデルをダウンロードします。その後、完全にオフラインで実行されます。
3. Wave はローカル モデル (Ollama / LM Studio) をサポートしますが、エンジンを自分でインストールして実行します。
4. Warp は独自の AI をクレジット (20 ドルのエントリー有料レベルで月 1,500 クレジット) で計測します。Warp の料金 FAQ によると、Free プランにはバンドルされた AI がまったく含まれていません。Free では独自のキーまたはエンドポイントを持参する必要があり、従量制ではありませんが、Warp のバックエンドを経由します (したがって ~ が付けられています)。 AITerm のローカルおよび BYO キー AI は、決して計測されず、プロキシされることもありません。
5. AITerm は、階層化された危険なコマンド分類子 (壊滅的なコマンドの明示的な確認) を実際の PTY で提供します。カーソルは、サンドボックス + ホワイトリスト + 分類子サブエージェント (カーソル自体がベストエフォート ガードレールと呼ぶ「自動レビュー」モード) を介してエージェント シェル呼び出しをゲートします。 Warp は、静的な正規表現の許可/拒否リストとアクションごとの承認を使用します。リスク スコア分類子は文書化されていません。 Wave の AI はまだコマンドを実行できません (近日公開予定)。
6. Warp は現在、公式の Claude Code / Codex CLI を端末でもホストしています (さらに別のクラウドハーネス モード)。ただし、それらは Warp 独自のクレジット従量制エージェントと並行して実行されます。 AITerm Pro では、Claude / ChatGPT サブスクリプションはアプリの AI エンジンです。同じ提案→承認→実行ループとセーフティ ゲートであり、ローカルで実行される公式 CLI によって強化されます。 (Wave の「Claude Code 統合」はタブバッジ通知であり、AI バックエンドではありません。)
7. Warp Drive にはパラメータ化されたワークフロー / ノートブック (無料版に限定) がありますが、Runbook の実行履歴のキャプチャやステップツーステップの出力チェーンはありません。
Wave はオープンソースであり、AITerm と同様にローカルファーストでプライベートです。プライバシーではなく、安全性分類子、ランブック、ゼロセットアップのネイティブ Mac エクスペリエンスで差別化を図っています。ソース: warp.dev/

価格設定と Warp のドキュメント、waveterm.dev のドキュメント、cursor.com/pricing とドキュメント (2026 年 7 月に再確認)。
アプリをダウンロードし、アプリケーションにドラッグして開きます - セットアップ
アシスタントを使用すると、最初の AI コマンドを約 1 分で利用できるようになります。
無料、公証済み、Apple Silicon、macOS 13 以降。
別の実験的な ai-term Python CLI (v0.2.0、
Mac アプリとは関係ありません): pip install https://ai-term.com/term_ai-0.2.0-py3-none-any.whl
AI を活用したターミナル アシスタント。自然言語からコマンドを生成し、エラーを即座に診断します。
リリースノートは電子メールで入手できます。スパムはなく、いつでも購読を解除できます。
© 2026 Vega LLC.無断転載を禁じます。 AITerm は Vega LLC の製品です。

## Original Extract

A native macOS terminal that speaks plain English. Approve AI-proposed commands, fix and explain errors, and use local Ollama, your API key, or your Claude / ChatGPT subscription.

Your terminal,
supercharged with AI
A native macOS terminal that speaks plain English. The AI proposes commands
you approve, fixes and explains errors, and runs multi-step tasks — on your
own AI: local Ollama, your API key, or your Claude / ChatGPT subscription.
or install with Homebrew
brew install --cask vega-llc/aiterm/ai-term
Download not starting? Direct link
Everything you need in your terminal
The native macOS app for day-to-day work, plus a free CLI for command generation and diagnosis.
Describe what you want in plain English. AITerm proposes an editable shell command and waits for your approval.
Use /fix and /explain on failed output to get a correction or a plain-English read of what happened.
Run fully on-device with Ollama or the managed Apple-Silicon (MLX) engine. Local requests never leave your Mac, and cloud requests go straight to your provider — no middle server — with secrets redacted first.
Bring your own key for OpenAI, Anthropic, OpenRouter, Groq, Gemini and more. Pro adds Amazon Bedrock, Azure OpenAI, and your existing Claude / ChatGPT subscription via the official CLIs.
Every command path goes through one safety check with green / amber / red risk labels and an auditable decision log.
Search command history, switch models, pull Ollama models, open the command palette, and share cleaned-up fix cards.
AITerm — a terminal that speaks plain English
A native macOS app, included with AITerm Pro: a real terminal emulator with AI woven
into the core loop. Type a command or describe what you want — the AI proposes an
editable command you confirm before it runs. It can fix and explain errors, run
multi-step tasks one approved step at a time, and save sequences as runbooks to replay
later. Runs on your AI: local Ollama, your API key, or your Claude / ChatGPT subscription.
Tabs — Many Terminals, One Window
Open as many tabs as you need — each is its own real shell with its own working directory, scrollback, and AI session. ⌘T to open, ⌘1–9 / ⌘⇧[ ] to switch; each tab's title shows its folder, and a command you approve in one tab only ever runs in that tab.
Split any tab into resizable panes — each its own real shell — and lay them out however you work. ⌘D splits right, ⌘⇧D down, ⌘⌥arrows move focus; drag the divider to resize. A new split opens in the folder you were already in, and your keystrokes only ever go to the focused pane.
Your tabs, splits, and working directories come back exactly as you left them when you relaunch. Save named layouts and reopen a whole project setup — every terminal in the right folder — in a single step.
Agent Mode — Manual or Autopilot, Always Fenced
Give it a goal. With /agent it works step by step — proposes each command, you approve every one. With /auto (Autopilot) it runs on its own, but only auto-executes commands the safety policy classifies as safe for the current project; risky, destructive, network, sensitive, and build/dependency steps pause for your one-tap approval or block outright. Both pass the same safety gate and your Safety Profile — so on a locked-down machine, anything destructive is blocked outright, even in Autopilot.
Save a multi-step sequence once — or capture it straight from commands you just ran with ⌘⌥R — with {{variables}} you fill in at run time, then replay it with a single pre-flight approval. Destructive steps are flagged up front, and if one fails you can /fix and retry from exactly where it stopped. Chain steps by feeding one step's output into the next with {{step1.output}} .
When a new version is ready, a subtle “Update ready — Install & Relaunch” pill appears in the status bar — no modal interrupting your session, and no silent background install that never lands on an app you leave open for days. Signed, notarized, delivered as tiny delta updates.
Plain English streams into an editable command with a green / amber / red risk badge — press Return to run, edit to change, Esc to discard. Risky commands never run without your confirmation; in Autopilot only provably-safe steps run on their own.
Type a shell command or plain English in one place — no modes. A local classifier routes it; start with “?” to force AI, or a space to force a literal command.
/fix reads a failed command's exact output — plus your repo, branch, and language — and proposes a correction shown side by side with the original, saveable as a runbook in one click. /explain (⌘⇧E) gives a plain-English read of any output — line by line if you want.
Run it your way — free: local & private with Ollama, or your own key for OpenAI, Anthropic, OpenRouter, Groq, Gemini, DeepSeek and more (everything stays on your Mac). Pro plugs in your existing Claude or ChatGPT subscription via the official CLIs, plus Amazon Bedrock, Azure OpenAI, and a managed on-device Apple-Silicon (MLX) engine.
Every command — typed, translated, or AI-suggested — passes a single, type-enforced safety check, red-teamed against ~30 evasion tricks. Switch Safety Profiles (Personal → Work → Production → Locked Down) to make the gate as strict as the machine you're on. Audit log of every decision.
Before you run something destructive, AITerm offers a conservative way back when one safely exists — an inverse command ( git checkout - , mv back, rmdir ) or a backup-first step ( cp -a … .bak ). It stays quiet rather than inventing a fake “undo.”
Local models never leave your Mac. Cloud requests go straight to your provider on your account — no middle server — with keys, tokens, and secrets redacted first.
Full PTY terminal: vim, top, ssh, colors, Ctrl-C, and tab-completion all work — with AI on top, not bolted on the side.
Command palette (⌘K), model manager to pull and switch Ollama models, fuzzy history search (by command or the plain English you typed), live git-branch status, themes, share-a-fix cards, and onboarding.
The native AI terminal — free forever . Private local AI, set up for you on first run. Bring your own cloud key — everything stays on your Mac. Pro adds your Claude / ChatGPT subscription + automation. No account needed to start.
Apple Silicon · macOS 13+ · notarized by Apple · auto-updates
or install with Homebrew
brew install --cask vega-llc/aiterm/ai-term
Fly through tabs, splits, and the AI tools. Inside the app, press ⌘/ (or Help ▸ Keyboard Shortcuts) to pull this up any time.
The native AI terminal is free forever — private local AI, bring any cloud key, splits, saved workspaces, and the full safety gate. Pro plugs in your existing Claude / ChatGPT subscription and adds runbooks + agent mode.
The native macOS app — the AI terminal that actually runs commands. No account needed.
AI command loop — propose, approve, run
/fix & /explain any command or output
Tabs, split panes & saved workspaces
Safety gate, safety profiles & rollback plans
Private local AI — set up on first run, then fully offline, no key
Bring any cloud key — OpenAI, Anthropic, OpenRouter, Groq, DeepSeek, xAI, Mistral, Gemini & more
No proxy — keys & prompts never leave your Mac
Just $8.25/mo, billed annually — save ~17%.
Early-bird price — locked in for life . 14-day free trial, no card — the trial starts automatically the first time you open the app.
Your AI subscription — use Claude Max / ChatGPT Plus via the official CLI, run locally (no proxy, no per-request metering)
Runbooks — save, parameterize & replay multi-step sequences; capture from history; output chaining
/agent & /auto — run multi-step tasks: approve each step ( /agent ), or let Autopilot auto-run the safe ones ( /auto ) — both behind the same safety gate
Managed MLX engine — a second on-device engine, Apple-Silicon-native, set up for you
Amazon Bedrock & Azure OpenAI connections
Up to 3 of your machines per seat
Not ready? Just download the app — every new install starts the 14-day Pro trial, no card.
Small team? Buy a seat per person. Need 10+ seats or invoicing? Email us .
Pro checks in with our license server about once a month to keep your subscription active —
automatic and invisible as long as you go online now and then. Only your license is validated;
your commands and files never leave your Mac. If your Mac stays fully offline for more than about
six weeks, Pro pauses until you reconnect once.
The difference is where your AI runs : AITerm keeps your keys, prompts, and local model 100% on your Mac — no vendor proxy in the middle. Honest snapshot, July 2026 — these tools move fast, so this is today's picture, not a permanent moat.
✓ yes · ~ partial · — none / undocumented · ✗ no
1. Per Warp's own docs, BYOK prompts and your key pass through Warp's backend on each request. Warp's local-model support works the same way: your model must be exposed at a public tunnel URL, and every prompt still round-trips through Warp's cloud to reach a model running on your own machine. AITerm's BYO key and local model never leave your Mac.
2. AITerm downloads a small local engine and model for you on first run — then it runs fully offline.
3. Wave supports local models (Ollama / LM Studio), but you install and run the engine yourself.
4. Warp meters its own AI by credits (1,500 credits/mo on the $20 entry paid tier), and per Warp's pricing FAQ the Free plan now includes no bundled AI at all — on Free you must bring your own key or endpoint, which is unmetered but still routes through Warp's backend (hence the ~). AITerm's local and BYO-key AI is never metered and never proxied.
5. AITerm ships a tiered dangerous-command classifier (explicit confirmation for catastrophic commands) in a real PTY. Cursor gates agent shell calls via sandbox + allowlist + a classifier subagent (its "Auto-review" mode, which Cursor itself calls best-effort guardrails). Warp uses static regex allow/deny lists plus per-action approval — no risk-scoring classifier documented. Wave's AI can't execute commands yet (listed as coming soon).
6. Warp now hosts the official Claude Code / Codex CLIs in its terminal too (plus a separate cloud-harness mode) — but they run alongside Warp's own credit-metered Agent. In AITerm Pro, your Claude / ChatGPT subscription is the app's AI engine: the same propose→approve→run loop and safety gate, powered by the official CLI running locally. (Wave's "Claude Code integration" is tab-badge notifications, not an AI backend.)
7. Warp Drive has parameterized workflows / notebooks (limited on Free) — but no runbook run-history capture or step-to-step output chaining.
Wave is open-source and, like AITerm, local-first and private — there we differentiate on the safety classifier, runbooks, and a zero-setup native Mac experience, not on privacy. Sources: warp.dev/pricing & Warp docs, waveterm.dev docs, cursor.com/pricing & docs (re-checked July 2026).
Download the app, drag it to Applications, open it — the setup
assistant gets you to your first AI command in about a minute.
Free, notarized, Apple Silicon, macOS 13 or later.
There is also a separate, experimental ai-term Python CLI (v0.2.0,
unrelated to the Mac app): pip install https://ai-term.com/term_ai-0.2.0-py3-none-any.whl
AI-powered terminal assistant. Generate commands from natural language and diagnose errors instantly.
Get release notes by email — no spam, unsubscribe anytime.
© 2026 Vega LLC. All rights reserved. AITerm is a product of Vega LLC.
