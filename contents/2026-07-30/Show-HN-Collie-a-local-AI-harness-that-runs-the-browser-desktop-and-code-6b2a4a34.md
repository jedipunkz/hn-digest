---
source: "https://github.com/colliehq/collie"
hn_url: "https://news.ycombinator.com/item?id=49115262"
title: "Show HN: Collie – a local AI harness that runs the browser, desktop and code"
article_title: "GitHub - colliehq/collie: A local-first coding agent that proves its own work — one cross-platform Python codebase, one-click Windows installer. · GitHub"
author: "wudmaing00"
captured_at: "2026-07-30T20:56:19Z"
capture_tool: "hn-digest"
hn_id: 49115262
score: 3
comments: 1
posted_at: "2026-07-30T20:24:02Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Collie – a local AI harness that runs the browser, desktop and code

- HN: [49115262](https://news.ycombinator.com/item?id=49115262)
- Source: [github.com](https://github.com/colliehq/collie)
- Score: 3
- Comments: 1
- Posted: 2026-07-30T20:24:02Z

## Translation

タイトル: Show HN: Collie – ブラウザ、デスクトップ、コードを実行するローカル AI ハーネス
記事のタイトル: GitHub - Colliehq/collie: 独自の機能を証明するローカルファーストのコーディング エージェント — 1 つのクロスプラットフォーム Python コードベース、ワンクリック Windows インストーラー。 · GitHub
説明: 独自の機能を証明するローカルファーストのコーディング エージェント — 1 つのクロスプラットフォーム Python コードベース、ワンクリック Windows インストーラー。 - コリーク/コリー

記事本文:
GitHub - Colliehq/collie: 独自の機能を証明するローカルファーストのコーディング エージェント — 1 つのクロスプラットフォーム Python コードベース、ワンクリック Windows インストーラー。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。もう一度お願いします

このページをロードします。
コリーク
/
コリー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
243 コミット 243 コミット .github/ workflows .github/ workflows アセット アセット ベンチ ベンチ データ データ ドキュメント ドキュメント ハーネス ハーネス インストーラー インストーラー 着陸 着陸 リレー リレー テスト テスト vscode-collie vscode-collie .gitignore .gitignore CHANGELOG.md CHANGELOG.md HANDOFF.md HANDOFF.md ライセンス ライセンスREADME.md README.md mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml pytest.ini pytest.ini run_compare.py run_compare.py run_compare4.py run_compare4.py run_Parallel.py run_Parallel.py swe_run.py swe_run.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
コンピュータ上に常駐し、実際に実行できるコーディング エージェント。
ローカルかつプライベート。それはあなたの実際の環境、つまりログインしたブラウザ、デスクトップ、
画面やファイルを実行して、その動作を証明します。
コリー.ラン ·
コリー -p "バグを修正" ·
コリーのウェブ
ほとんどのコーディング エージェントはクラウド タブまたはエディタ ペインに存在し、ユーザーが渡したファイルのみを操作できます。
彼ら。 Collie はあなたのマシン上で実行されます。つまり、あなたがすでに行っているように動作します。
ログインしたブラウザ、デスクトップの配置、画面の記録、携帯電話からのタスクの取得、および
コードを編集します。あなたが送信しない限り、あなたのコンピュータから何も残されません。アカウントも無いし
テレメトリー。
そして、それは単に完了したと主張するだけではありません。コリーは何かを修正すると、次のような複製を書き込みます。
壊れたコードで失敗し、コードを反転する最小限の編集を行って、アサーションを再実行する必要があります。
実行は「完了」ではなく、検証済みです ✓ 。
それはローカルであり、あなたの現実世界に届きます。クラウド エージェントはリポジトリを読み取ることができます。コリーは開けることができます
ブラウザ内のサイト

すでにログインしているので、実際のフローをクリックして、何が起こるかを見てください。
画面にアクセスし、コードを変更します。すべて 1 台のマシン上で、すべて制御下にあります。それは違います
タスクのクラス: 「これらのファイルを編集する」ではなく、「これをエンドツーエンドで機能させる」。
範囲がその証拠です。 Collie は、無関係な機能を大量に追加したコーディング エージェントではありません。
以下の範囲 — デスクトップ コンソール、ブラウザ コントロール、スクリーン レコーダー、電話のリモコン —
Collie のコーディング エージェントがすべてを構築したために存在します。機能はベンチマークです。
ハーネスは独自のデスクトップ アプリを出荷できるほど強力であり、iOS コンパニオンはバグに対して十分な強度を持っています。
能力
それが何を意味するか
🧠
コーディングエージェント
セマンティック コード ナビゲーション、構文ゲート編集、および自己検証修復ループの核心については、以下で説明します。
🌐
実際のブラウザ
Chrome 拡張機能を使用すると、Collie はログインしたブラウザ (実際のセッション、実際の Cookie) で動作できるため、サイトをスクレイピングするだけでなく、サイトを操作できるようになります。すべてのアクションは、フェンスで囲まれ、CSRF チェックされたローカルホスト呼び出しです。
🖥️
リビングデスクトップ
Collie Web は、時計、天気、アプリ ドック、プロジェクト、音楽プレーヤー (リアル オーディオ + 同期されたカラオケの歌詞)、コマンド バーなどのインタラクティブなアンビエント壁紙を強化します。これらはすべて 1 つの JSON 構成を介してエージェントが管理できます。 Collie が作業しているとき、壁紙はコードのライブ星図になります。
🎬
スクリーンレコーダー
コリーレコードは、画面 + カメラ + マイク (Windows および macOS) をキャプチャします。これは、実行をデモまたは文書化するための組み込みの方法です。
📱
電話のリモコン
コードをスキャンして一度ペアリングします。次に、テールを実行し、同じ Wi-Fi ( --lan ) またはリレー ( --remote ) を介した任意の場所で、コンパニオン iOS アプリを使用して、携帯電話から新しいテールを開始します。
🔌
他のどこでも
ターミナル、ブラウザ GUI、VS Code、および任意の ACP エディタ (Zed/JetBrains/neovim) — 1 つのハーネス、すべてのサーフェス。
どこで実行されるか

Collie はターミナルファーストであり、特注の拡張機能ではなくオープン プロトコルを通じて編集者に届きます。
Windows — ワンクリック。 Collie-Setup.exe を次の場所からダウンロードします。
最新のリリースを選択してダブルクリックします。小さな
アプリ スタイルのインストーラーは、自己完結型のランタイム (Python + Collie + セマンティック メモリ、何も必要ありません) を構築します。
プレインストール)、ネイティブのデスクトップ ウィンドウで Collie を開きます。最初の起動時に頭脳を選択します。
既存の Claude、Codex、または Grok ログインが検出され、ワンクリックで接続されます。または API キーを貼り付けます。
macOS / Linux — ピップ。コアは stdlib のみであるため、基本インストールは小規模です。
pip install -e " .[local,dev] " # クローンから (PyPI 公開が計画されています)
コリーのセットアップ # オプションの deps、メモリ モデルの事前ダウンロード、プロバイダーの選択
Collie # ターミナルチャット (TUI) が開きます
アカウントやテレメトリはなく、コアにはサードパーティへの依存関係がありません (モックとオラマ)
キーなしで実行すると、メモリは BM25 キーワードの呼び出しですぐに機能します。
オプションの追加機能: pip install ".[local,tui,search]" — ローカル (セマンティック メモリ: granite-107m via
onnxruntime、~55MB、多言語)、tui (リッチ ターミナル チャット)、検索 (キーレス Web 検索)、acp
(エディター プロトコル)、ブラウザー (Playwright — コリー専用 browser-bridge --browser 、マネージド
拡張機能がプリロードされた Chromium (CI 用、または独自の Chrome を使用したくない場合)。 OSごと
セットアップ — 特にリアルブラウザブリッジ (collie Browser-bridge + harness/browser_ext/) — は次のとおりです。
docs/PLATFORMS.md 。
コリー # ターミナルチャット (TUI);最初の実行ではプロバイダーが選択されます
Collie Web # ブラウザ GUI — チャット、ライブ ゲート、差分、スターマップ、アンビエント デスクトップ
コリーのセルフテスト # $0 の決定論的なエンドツーエンド (モック モデル、実際のツール + メモリ)
# 本当に安価なモデル (env のプロバイダー キー)
DEEPSEEK_API_KEY=...collie -p " utils/timeparse.py のオフバイワンを修正します "
# 機械可読 / st

リーミング
コリー実行「バグを修正」 --json # 最終結果オブジェクト (トークン、コスト、検証済み)
コリー実行「バグを修正」 --stream-json # ライブ NDJSON: ツール · 編集 · リプロゲート · 受信
# 完全にローカル、キーなし
コリーは「summary app.py」を実行します --provider ollama --model qwen2.5-coder:7b
# 自律ループ: ゴールに向かって反復し、実行されたチェックが緑になった最初のターンで停止します
コリーループ --goal " スイートを取得します " --until " pytest -q " --max 8
# 実行ベースの選択によるベストオブ N: N 個の分離試行を実行し、合格したもののみを保持します
コリーパック " 失敗したテストを修正 " -n 3 --check " pytest -q " --apply
Collie acp # ACP エージェントとして機能します (エディターがこれを stdio 経由で生成します)
プロバイダー:mock 、 ollama 、 anthropic 、 anthropic-oauth 、および OpenAI 互換プリセット
ディープシーク、qwen / ダッシュスコープ、openrouter、moonshot、groq、zhipu、openai 。
上記のすべては、小さくて正直なハーネスに依存します。その下にあるのがこれです。
署名: 検証ゲート
locate code_search "parse_duration 複合単位" · 4 件がヒットしました
› utils/timeparse.py:42 _parse · · · · · · · · 0.91
repro が repro.py を書きました · アサート parse_duration("1h30m") == 5400
✗ 失敗 › 1800 を獲得、5400 が欲しい ← ゲートが装備されました
utils/timeparse.py を編集 · · · · · +1 -1
43 │- 合計 = SECONDS[単位] * int(val)
43 │+ 合計 += SECONDS[単位] * int(val)
Python repro.pyを検証する
✓ 合格 › parse_duration("1h30m") == 5400 ← ゲートグリーン
✓ 12.8秒で検証 · Δ +1 −1 · 3,410 tok · $0.006
他のエージェントは「テストにエラーがなかったかどうか」をチェックします。コリーのゲートはより強力です。繁殖には、
問題から派生した実際の == 期待値をアサートするため、もっともらしいが間違った編集は大声で失敗し、
別の修理ラウンドを推進します。このアサート検証ループは、

ハーネス - 間違った編集
黙って「完了」として出荷されることはありません。同じアイデアがスケールアップされます。実際のシェルのチェック時にコリーのループが停止します。
0 で終了し、コリーのパックは実際に合格した N 回の試行のうち最良のものを選択します。
アーキテクチャ (抽象化と継ぎ目)
┌───── ループ.ハーネス ───────┐
タスク ────────▶│ 作成 → 完了 → ツールの実行 → 検証 ✓ │
━─┬─────────┬───────┬───────┘
┌─────────────┘ │ └───────────┐
▼ ▼ ▼
コンテキストコンポーザーモデルプロバイダーツールレジストリ
安定/コンテキスト/揮発性 OpenAI-compat · Anthropic · 読み取り/書き込み/編集/bash/
+ トークン予算担当者 Ollama · subscription-OAuth grep/glob + code_search
▼ │ │
メモリ.SqliteMemory ▼ レコーダー.レコーダー
ハイブリッド リコール (BM25+dense+RRF) の発行 → stream-json / ACP run.db (+ ダッシュボード)
シーム（抽象ベース）
出荷された内容
モデルプロバイダー
OpenAICompat (DeepSeek/Qwen/GLM/OpenRouter…) · Anthropic · Ollama · subscription-OAuth
ツールレジストリ
読み取り/書き込み/編集 (構文ゲート) · bash · grep · glob · code_search · web_search + web_fetch (キーレス) · plan · undo · ブラウザ · MCP (遅延層 +load_tools )
埋め込みプロバイダー
OnnxEmbedding granite-107m (Apache、55MB、多言語) · bge-m3 / e5 · jina-v3 オプトイン · BM25-モデルなしの場合のみ
Sqliteメモリ
CORE + ファクト + FTS5 + コサイン、ハイブリッド RRF + オプションの再ランク + 統合
コンテキスト作曲家
STABLE/CONTEXT/VOLATILE + 自動プリフェッチ · ~1K トークンの固定プレフィックス (意図的に無駄を省いた)
code_search は自然言語クエリから識別子を抽出し、

リポジトリを grep します (ripgrep,
else grep)、それぞれに含まれる用語の数によってファイルをランク付けします。したがって、エージェントは、どこに含まれるかを推論します。
ブラインドで grep する代わりに、モデルやインデックスが古くなって編集する必要があります。 edit_file は
完全一致、空白文字許容、Python 構文を壊す編集を拒否します。
信頼できない Web/ページ コンテンツはデータとしてフェンスされ (プロンプト インジェクション防御)、ブラウザ ブリッジによって保護されます。
CSRF ヘッダーが欠落しているリクエストはすべて拒否されます。トークン/コストの予算
( COLLIE_MAX_COST / COLLIE_MAX_TOTAL_TOKENS ) は天井で実行を停止します。
OS ごとのフォークではなく、1 つのクロスプラットフォーム Python コードベース。本当に
異なる (タイムアウトでプロセス ツリーを強制終了、トークン ファイルを保護、パスを変換、シェルを選択) は次のとおりです。
harness/plat.py で分離されているため、同じホイールがどこでも実行されます。
Collie は、同じタスクとモデルで他のハーネスと比較して自分自身を測定します。それを自分で実行します。
ここでは数値はアサートされていません。
DEEPSEEK_API_KEY=... python swe_run.py --n 5 # SWE ベンチ検証済み (Docker が必要)
python -m bench.multirun_eval # pass@1 / pass@k / Wilson CI / McNemar
python -m bench.polyglot_eval --langs python,cpp,javascript # Aider-Polyglot、多言語
python -m harness.cli Compare --vs all # vs クロード コード / Aider / …
正直さとポリシー
ベンチマーク ハーネスにはバージョンタグが付けられており、再現可能です。 「進歩は数字である」は双方向を切り開く -
コリーは、勝利だけでなくネット中立性をもたらすレバーを浮上させます。
トークン数は実際の使用量 (モデル自体の使用量、または CLI の harness/apitap.py メーターリング) です。
報告なし) — 同一の内容、両方の側で同じソース。
Collie は、ファーストパーティ OAuth パスを通じてのみ個人の Max/Pro サブスクリプションを取得します。
( anthropic-oauth )、同じメカニズム

[切り捨てられた]

## Original Extract

A local-first coding agent that proves its own work — one cross-platform Python codebase, one-click Windows installer. - colliehq/collie

GitHub - colliehq/collie: A local-first coding agent that proves its own work — one cross-platform Python codebase, one-click Windows installer. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
colliehq
/
collie
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
243 Commits 243 Commits .github/ workflows .github/ workflows assets assets bench bench data data docs docs harness harness installer installer landing landing relay relay tests tests vscode-collie vscode-collie .gitignore .gitignore CHANGELOG.md CHANGELOG.md HANDOFF.md HANDOFF.md LICENSE LICENSE README.md README.md mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml pytest.ini pytest.ini run_compare.py run_compare.py run_compare4.py run_compare4.py run_parallel.py run_parallel.py swe_run.py swe_run.py View all files Repository files navigation
A coding agent that lives on your computer — and can actually run it.
Local and private. It reaches your real environment — your logged-in browser, your desktop,
your screen, your files — and proves its work by running it.
collie.run ·
collie -p "fix the bug" ·
collie web
Most coding agents live in a cloud tab or an editor pane and can only touch the files you hand
them. Collie runs on your machine — so it works the way you already do: it drives your real
logged-in browser, arranges your desktop, records your screen, takes tasks from your phone, and
edits your code. Nothing leaves your computer unless you send it there; there's no account and no
telemetry.
And it doesn't just claim to be done. When Collie fixes something it writes a reproduction that
must fail on the broken code, makes the smallest edit that flips it, and re-runs the assertion — a
run isn't "done," it's verified ✓ .
It's local, and it reaches your real world. A cloud agent can read a repo. Collie can open the
site in the browser you're already logged into, click through the actual flow, watch what happens on
your screen, and change the code — all on one machine, all under your control. That's a different
class of task: not "edit these files," but "get this working, end to end."
The range is the proof. Collie isn't a coding agent with a pile of unrelated features bolted on.
The breadth below — the desktop console, the browser control, the screen recorder, the phone remote —
is there because Collie's coding agent built all of it. The features are the benchmark: a
harness strong enough to ship its own desktop app and iOS companion is strong enough for your bug.
Capability
What it means
🧠
Coding agent
Semantic code navigation, syntax-gated edits, and a self-verifying repair loop — the core, covered below.
🌐
Your real browser
A Chrome extension lets Collie act in your logged-in browser — the real session, real cookies — so it can operate sites, not just scrape them. Every action is a fenced, CSRF-checked localhost call.
🖥️
Living desktop
collie web powers an interactive ambient wallpaper: clock, weather, an app dock, projects, a music player (real audio + synced karaoke lyrics), and a command bar — all agent-manageable via one JSON config. When Collie is working, the wallpaper becomes a live star-map of your code.
🎬
Screen recorder
collie record captures screen + camera + mic (Windows and macOS) — a built-in way to demo or document a run.
📱
Phone remote
Pair once by scanning a code; then tail runs and start new ones from your phone — on the same Wi-Fi ( --lan ) or anywhere through a relay ( --remote ), with the companion iOS app.
🔌
Everywhere else
Terminal, browser GUI, VS Code, and any ACP editor (Zed/JetBrains/neovim) — one harness, every surface.
Where it runs
Collie is terminal-first and reaches editors through an open protocol, not a bespoke extension:
Windows — one click. Download Collie-Setup.exe from the
latest release and double-click it. A small
app-style installer lays down a self-contained runtime (Python + Collie + semantic memory, nothing to
preinstall) and opens Collie in a native desktop window. On first launch you pick a brain — an
existing Claude, Codex, or Grok login is detected and connects in one click; or paste an API key.
macOS / Linux — pip. The core is stdlib-only, so the base install is tiny:
pip install -e " .[local,dev] " # from a clone (PyPI publish is planned)
collie setup # optional deps, pre-download the memory model, pick a provider
collie # the terminal chat (TUI) opens
No account, no telemetry, and the core has zero third-party dependencies — mock and ollama
run without any key, and memory works out of the box on BM25 keyword recall.
Optional extras: pip install ".[local,tui,search]" — local (semantic memory: granite-107m via
onnxruntime, ~55MB, multilingual), tui (rich terminal chat), search (keyless web search), acp
(editor protocol), browser (Playwright — only for collie browser-bridge --browser , a managed
Chromium with the extension preloaded, for CI or when you'd rather not use your own Chrome). Per-OS
setup — especially the real-browser bridge ( collie browser-bridge + harness/browser_ext/ ) — is in
docs/PLATFORMS.md .
collie # terminal chat (TUI); first run picks a provider
collie web # browser GUI — chat, live gate, diffs, star-map, ambient desktop
collie selftest # $0 deterministic end-to-end (mock model, real tools + memory)
# a real cheap model (provider key in env)
DEEPSEEK_API_KEY=... collie -p " fix the off-by-one in utils/timeparse.py "
# machine-readable / streaming
collie run " fix the bug " --json # final result object (tokens, cost, verified)
collie run " fix the bug " --stream-json # live NDJSON: tool · edit · repro-gate · receipt
# fully local, no key
collie run " summarize app.py " --provider ollama --model qwen2.5-coder:7b
# autonomous loop: iterate toward the goal, STOP the first turn an executed check goes green
collie loop --goal " get the suite passing " --until " pytest -q " --max 8
# best-of-N with EXECUTION-based selection: run N isolated attempts, keep only what passes
collie pack " fix the failing test " -n 3 --check " pytest -q " --apply
collie acp # serve as an ACP agent (an editor spawns this over stdio)
Providers: mock , ollama , anthropic , anthropic-oauth , and OpenAI-compatible presets
deepseek · qwen / dashscope · openrouter · moonshot · groq · zhipu · openai .
Everything above rests on a small, honest harness. This is what's under it.
The signature: the verification gate
locate code_search "parse_duration compound units" · 4 hits
› utils/timeparse.py:42 _parse ············· 0.91
repro wrote repro.py · assert parse_duration("1h30m") == 5400
✗ FAILING › got 1800, want 5400 ← gate armed
edit utils/timeparse.py ································· +1 −1
43 │- total = SECONDS[unit] * int(val)
43 │+ total += SECONDS[unit] * int(val)
verify python repro.py
✓ PASSING › parse_duration("1h30m") == 5400 ← gate green
✓ verified in 12.8s · Δ +1 −1 · 3,410 tok · $0.006
Other agents check "did the test not error." Collie's gate is stronger: the reproduction carries an
assert actual == expected derived from the issue, so a plausible-but-wrong edit fails loudly and
drives another repair round. This assert-verify loop is the core of the harness — a wrong edit
never silently ships as "done." The same idea scales up: collie loop stops when a real shell check
exits 0, and collie pack picks the best of N attempts by what actually passes.
Architecture (abstractions & seams)
┌──────────────── loop.Harness ────────────────┐
task ─────────────▶│ compose → complete → run tools → verify ✓ │
└──┬──────────────┬──────────────┬─────────────┘
┌───────────────────┘ │ └───────────────────┐
▼ ▼ ▼
ContextComposer ModelProvider ToolRegistry
STABLE/CONTEXT/VOLATILE OpenAI-compat · Anthropic · read/write/edit/bash/
+ token budgeter Ollama · subscription-OAuth grep/glob + code_search
▼ │ │
memory.SqliteMemory ▼ recorder.Recorder
hybrid recall (BM25+dense+RRF) emit → stream-json / ACP runs.db (+ dashboard)
Seam (abstract base)
shipped impl
ModelProvider
OpenAICompat (DeepSeek/Qwen/GLM/OpenRouter…) · Anthropic · Ollama · subscription-OAuth
ToolRegistry
read/write/ edit (syntax-gated) · bash · grep · glob · code_search · web_search + web_fetch (keyless) · plan · undo · browser · MCP (deferred tier + load_tools )
EmbeddingProvider
OnnxEmbedding granite-107m (Apache, 55MB, multilingual) · bge-m3 / e5 · jina-v3 opt-in · BM25-only when no model
SqliteMemory
CORE + facts + FTS5 + cosine, hybrid RRF + optional rerank + consolidation
ContextComposer
STABLE/CONTEXT/VOLATILE + auto-prefetch · a ~1K-token fixed prefix (kept deliberately lean)
code_search extracts the identifiers from a natural-language query and greps the repo (ripgrep,
else grep), ranking files by how many of your terms each contains — so the agent reasons about where
to edit instead of grepping blind, with no model and no index to go stale. edit_file is
exact-match, whitespace-tolerant, and rejects any edit that would break Python syntax .
Untrusted web/page content is fenced as data (prompt-injection defense), and the browser bridge
refuses any request missing its CSRF header. A token/cost budget
( COLLIE_MAX_COST / COLLIE_MAX_TOTAL_TOKENS ) stops a run at a ceiling.
One cross-platform Python codebase — not a per-OS fork. The handful of operations that genuinely
differ (kill a process tree on a timeout, secure a token file, convert a path, choose a shell) are
isolated in harness/plat.py , so the same wheel runs everywhere.
Collie measures itself against other harnesses on the same task and model — you run it yourself;
no numbers are asserted here:
DEEPSEEK_API_KEY=... python swe_run.py --n 5 # SWE-bench Verified (needs Docker)
python -m bench.multirun_eval # pass@1 / pass@k / Wilson CI / McNemar
python -m bench.polyglot_eval --langs python,cpp,javascript # Aider-Polyglot, multi-language
python -m harness.cli compare --vs all # vs Claude Code / Aider / …
Honesty & policy
The benchmark harness is version-tagged and reproducible. "Progress is a number" cuts both ways —
Collie surfaces the levers that turn out net-neutral , not just the wins.
Token counts are real usage (the model's own usage , or harness/apitap.py metering for CLIs that
report none) — apples-to-apples, same source both sides.
Collie draws a personal Max/Pro subscription only through the first-party OAuth path
( anthropic-oauth ), the same mechanism

[truncated]
