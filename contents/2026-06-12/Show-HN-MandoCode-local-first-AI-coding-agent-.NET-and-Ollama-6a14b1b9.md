---
source: "https://github.com/DevMando/MandoCode"
hn_url: "https://news.ycombinator.com/item?id=48499532"
title: "Show HN: MandoCode – local-first AI coding agent (.NET and Ollama)"
article_title: "GitHub - DevMando/MandoCode: A .NET C# CLI Coding Agent powered by Ollama + Semantic Kernel and RazorConsole. Run locally or in the cloud. Refactors code, proposes diffs, and updates your project safely — no API keys required. · GitHub"
author: "devmando"
captured_at: "2026-06-12T04:34:41Z"
capture_tool: "hn-digest"
hn_id: 48499532
score: 1
comments: 0
posted_at: "2026-06-12T03:22:10Z"
tags:
  - hacker-news
  - translated
---

# Show HN: MandoCode – local-first AI coding agent (.NET and Ollama)

- HN: [48499532](https://news.ycombinator.com/item?id=48499532)
- Source: [github.com](https://github.com/DevMando/MandoCode)
- Score: 1
- Comments: 0
- Posted: 2026-06-12T03:22:10Z

## Translation

タイトル: Show HN: MandoCode – ローカルファースト AI コーディング エージェント (.NET および Ollama)
記事のタイトル: GitHub - DevMando/MandoCode: Ollama + Semantic Kernel および RazorConsole を利用した .NET C# CLI コーディング エージェント。ローカルまたはクラウドで実行します。コードをリファクタリングし、差分を提案し、プロジェクトを安全に更新します。API キーは必要ありません。 · GitHub
説明: Ollama + Semantic Kernel および RazorConsole を利用した .NET C# CLI コーディング エージェント。ローカルまたはクラウドで実行します。コードをリファクタリングし、差分を提案し、プロジェクトを安全に更新します。API キーは必要ありません。 - DevMando/MandoCode
HN テキスト: MandoCode を構築しました。これは、ローカルの Ollama モデルに対して実行される、.NET のオープン ソース CLI コーディング エージェントです。 API キーはなく、マシンから何も出ません (必要に応じて Ollama クラウド モデルも機能します)。コードベース全体で Web の読み取り、編集、検索、計画、閲覧を行い、MCP とスキルをサポートします。 dotnet ツール install -g MandoCode を使用してインストールします。セマンティック カーネルは、.NET C# 上でビルドできるようにしてくれる優れものであると感じました。
他にもやるべきことがたくさんあります!

記事本文:
GitHub - DevMando/MandoCode: Ollama + Semantic Kernel および RazorConsole を利用した .NET C# CLI コーディング エージェント。ローカルまたはクラウドで実行します。コードをリファクタリングし、差分を提案し、プロジェクトを安全に更新します。API キーは必要ありません。 · GitHub
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
デ

vマンド
/
マンドコード
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
157 コミット 157 コミット .github .github .mandocode/ スキル/ コミットメッセージ .mandocode/ スキル/ コミットメッセージ docs docs docs src/ MandoCode src/ MandoCode テスト/ MandoCode.Tests テスト/ MandoCode.Tests .gitattributes .gitattributes .gitignore .gitignore ライセンス ライセンスMandoCode.sln MandoCode.sln README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング アシスタント — Ollama を使用してローカルまたはクラウドで実行します。
API キーは必要ありません。あなたとあなたのコードだけ。
MandoCode は、RazorConsole 上に構築され、Semantic Kernel と Ollama を利用した AI コーディング アシスタントです。 RazorConsole は、ターミナル UI 全体、つまり Razor コンポーネント、仮想 DOM、および Spectre.Console レンダリングをすべてコンソールで実行できるようにします。
ローカルで実行するか、Ollama クラウドに接続します。Web 検索を含め、API キーは必要ありません (オプションの無料の Tavily キーは検索の信頼性をアップグレードします)。これにより、端末から離れることなく、コードベース全体の読み取り、書き込み、検索、計画、Web ブラウジングなど、クロード コード スタイルのプロジェクト認識が得られます。 C#、JavaScript、TypeScript、Python、CSS、HTML、JSON、構成ファイルなど、あらゆる種類のファイルを理解します。
.NET 8 SDK — dotnet.microsoft.com/download/dotnet/8.0 (SDK にはランタイムが含まれています。SDK のみをインストールします)
Ollama — ollama.com/download (MandoCode は最初の実行時にセットアップを案内します)
dotnet ツールのインストール -g MandoCode
マンドコード
最初の実行では、ガイド付きウィザードが起動します。Ollama を検出して開始を提案し、より強力なモデルが必要な場合はクラウド サインインを案内し、適切なデフォルトを自動取得します。いつでも再実行できます

/setup を使用します。
マンドコード --doctor
ランタイム バージョン、Ollama ステータス、プルされたモデル、およびクラウド サインイン状態を出力します。
⚠️ ローカル モデル: Ollama コンテキスト ウィンドウを確認してください (一番の落とし穴)
クラウド モデル (:cloud タグ) を使用しますか?このセクションはスキップしてください。クラウド モデル コンテキストは Ollama のサーバーで管理され、デフォルトでモデルの最大値に設定されます。デスクトップ アプリのスライダーを含め、マシン上のものは何も影響しません。
ローカル モデルを使用していて、応答が途切れたり、モデルが以前の会話を「忘れたり」したり、書き込んだばかりのファイルで編集が繰り返し失敗したり、次のメッセージが表示されたりする場合は、次のメッセージが表示されます。
⚠ モデルのコンテキスト ウィンドウがいっぱいになったため、応答が切断されました…
…Ollama コンテキスト ウィンドウはほぼ間違いなく小さすぎます。コンテキスト ウィンドウは、モデルが一度に表示できる会話とコードの量です。Ollama のデフォルト値は ~4,000 トークンであり、エージェント セッションはほぼ即座にいっぱいになります。オーバーフローすると、最も古いコンテンツ (システム プロンプト、つまりモデルの指示を含む) が黙って削除されます。
Ollama デスクトップ アプリ (トレイ アイコン) を使用する場合、アプリの [設定] → [コンテキストの長さ] スライダーがこれを制御し、MandoCode の構成を含む他のすべてをオーバーライドします。
普遍的に正しいスライダーの位置はありません。モデルが認識できる量と GPU のメモリに収まる量とのトレードになります (8k のウィンドウごとに、モデルに応じて約 0.5 ～ 1.5 GB の VRAM がかかります)。
低すぎる (4k のデフォルト): 上記の症状 — モデル自体の命令が静かにウィンドウから落ち、動作が停止します。
GPU に対して高すぎる: モデルはシステム RAM に溢れ、トークン/秒のクレーターになり、クロールまたはハングしているように見えます。
開始点: ほとんどの GPU では 16k、8 GB 以上の VRAM では 32k。上記の症状が見られる場合にのみ、値を上げてください。レベルを上げた後に生成が著しく遅くなった場合は、ノッチを下げます。

ollamaserve を自分で実行する場合 (デスクトップ アプリは使用しない)、MandoCode がそれを処理します。デーモンの起動時に contextLength 設定から OLLAMA_CONTEXT_LENGTH を設定し、 /setup または /model で選択したモデルのハードウェア層に合わせて自動的にサイズを設定します。次のように手動で調整します。
mandocode --config set contextLength 16384
デーモンが ollam ps で実際に何を使用しているかを確認します (CONTEXT 列を確認してください)。 MandoCode 内で /learn を実行すると、フレンドリーな説明が表示されます。
⚠️ すべてのモデル: レスポンスキャップを確認してください (2 番目の注意点)
コンテキスト ウィンドウの邪悪な双子 — 上のスライダーとは異なり、これはクラウドを含むすべてのモデルに適用されます。モデルが作業を発表した後に停止した場合 (「ゲームを作成します…」と言って、計画もファイルもエラーも発生しないままターンが終了した場合)、maxTokens が低すぎます。これは 1 つの応答 ( NumPredict ) に上限を設け、推論モデルはツール呼び出しを発行する前に出力トークンを思考に費やすため、上限を低く設定すると、出力トークンが動作する前に遮断されます。
新規インストールのデフォルトは 32k ですが、それに気づくことはありません。ただし、構成が v0.11 より前の場合、またはコンテキスト ウィンドウだと思って maxTokens を下げたことがある場合 (これらは別のノブです。これはモデルが言うことをキャップし、コンテキスト ウィンドウは表示されるものをキャップします)、それを確認してください。
mandocode --config show # 「最大トークン」を確認します
mandocode --config set maxTokens 32768
明らかな兆候: トークン追跡では、ターンごとに、正確に上限に固定された出力が表示されます (例: 毎回 2,000 アウト)。実行中のセッションでは、起動時にロードされた構成が保持されることに注意してください。変更を有効にするには、MandoCode を再起動します (またはアプリ内で /config セットを使用します)。
git clone https://github.com/DevMando/MandoCode.git
cd MandoCode
dotnet build src/MandoCode/MandoCode.csproj
dotnet run --project src/MandoCode/MandoCode.csproj -- /path/to/your/project
MandoCode の違い
安全なファイル編集

異なる承認を伴うNG
すべてのファイルの書き込みと削除は、色分けされた差分によってインターセプトされます。あなたが承認、拒否、またはリダイレクトします。あなたの発言なしには何もディスクにアクセスできません。
@ と入力すると、プロジェクト ファイルがオートコンプリートされ、コンテキストとして添付されます。 AI は、プロンプトとともにファイルの完全な内容を確認します。 1 つのメッセージで複数のファイルを参照します。
複雑なリクエストは自動的に段階的な計画に分割されます。計画を確認し、進捗状況を追跡しながら各ステップの実行を監視します。
AI は Web を検索し、Web ページを読んでドキュメント、チュートリアル、または回答を見つけることができます。API キーは必要ありません。オプションで、DuckDuckGo のレート制限に達しない AI に最適化された検索用の無料の Tavily キーを追加します。
Lofi とシンセウェーブのトラックがバンドルされています。コーディング中に波形ビジュアライザーが隅で実行されます。雰囲気が大事だから。
Ollama が実行されていない場合、MandoCode は裸のエラーの代わりにセットアップ ガイダンスをインラインで表示します。再起動せずに再接続するには、/retry を使用します。
特徴
説明
AI
プロジェクトを意識したアシスタント
コードベース全体の読み取り、書き込み、削除、検索を行います。
AI
Web検索と取得
Web 検索と Web ページの読み取り - DuckDuckGo によるキーレス、または無料の API キーを使用した Tavily
AI
MCPサーバーのサポート
任意の Model Context Protocol サーバー (stdio またはリモート HTTP) に接続します — Claude-Desktop 互換の構成
AI
ストリーミング応答
アニメーション化されたスピナーによるリアルタイム出力
AI
タスクプランナー
複雑なリクエストを自動検出し、ステップに分割します
AI
フォールバック関数の解析
ツール呼び出しを生の JSON として出力するモデルを処理します
UI
差分承認
承認/拒否/リダイレクトによる色分けされた差分
UI
マークダウンレンダリング
豊富なターミナル出力 — ヘッダー、テーブル、コード ブロック、引用符
UI
構文の強調表示
C#、Python、JavaScript/TypeScript、Bash
UI
クリック可能なファイルのリンク
ファイルパスのOSC 8ハイパーリンク
UI
端末テーマの検出
アウ

明るい端末と暗い端末の色を調整します
UI
タスクバーの進行状況
タスク実行時の Windows ターミナルの統合
入力
/ コマンドのオートコンプリート
ドロップダウン ナビゲーションを備えたスラッシュ コマンド
入力
@ファイル参照
ファイルの内容を任意のプロンプトに添付する
入力
！シェルエスケープ
シェルコマンドをインラインで実行します ( !git status 、 !ls )
入力
/copy と /copy-code
応答またはコード ブロックをクリップボードにコピーする
音楽
ローファイ + シンセウェーブ
ボリューム、ジャンル切り替え、波形ビジュアライザーを備えたバンドルトラック
構成
構成ウィザード
モデルの選択と接続テストを含むガイド付きセットアップ
構成
構成の検証
無効な設定を安全な範囲に自動クランプします
信頼性
再試行 + 重複排除
指数関数的バックオフと重複通話の防止
教育
/学習コマンド
LLM 教育ガイド (オプションの AI 教育者チャット付き)
コマンド
/ と入力してオートコンプリート ドロップダウンを表示するか、「!」と入力します。シェルコマンドを実行します。
/setup — ガイド付きの初回実行ウィザード。 Ollama を検出し、インストールを提案し、クラウドへのサインインを案内し、ハードウェア対応階層を持つモデルを選択し、適切なデフォルトを自動プルします。何かが壊れたとき、または初心者のときに使用してください。
/model — クイックスイッチ。プルしたリストからモデルを選択して実行します。ローカルで選択すると、ハードウェア層に応じたサイズのコンテキスト ウィンドウが自動的に取得されます。単にモデルを交換したい場合に使用します。
/config — 設定を調整します。温度、タイムアウト、ディレクトリの無視などをカバーする完全な設定フォーム。どのノブを回すか正確にわかっている場合に使用します。
CLI フラグ (チャット ループの外側)
mandocode --doctor # プリフライト チェック: .NET ランタイム、Ollama ステータス、モデル、サインイン
mandocode --config show # 現在の設定を表示します
mandocode --config init # デフォルトの設定ファイルを作成します
mandocode --config set < key > < value > # 単一の値を設定します (例: set model qwen3.5:9b)
mandocode --config path # 構成ファイルの場所を表示
R

un mandocode --doctor チャットが不正な動作をしているときはいつでも — すべてが緑色の場合は 0 で終了し、何かが欠落している場合は 1 で終了し、問題の明確な概要が表示されます。
プロンプトを入力すると、
|
MandoCode はプロジェクト コンテキスト (@files、システム プロンプト) を追加します
|
セマンティック カーネルが Ollama に送信 (ローカルまたはクラウド モデル)
|
AI はテキスト + 関数呼び出しで応答します
|
ファイル操作は差分承認を経ます
Web の検索と取得が直接実行されます
|
ターミナルに表示される豊富なマークダウン
AI は、FileSystemPlugin (9 つの機能: ファイルのリスト、グロブ検索、読み取り、書き込み、ファイル/フォルダーの削除、テキスト検索、パス解決) および WebSearchPlugin (Tavily または DuckDuckGo による Web 検索、Web ページの取得 - API キーなしで動作) を通じてプロジェクトにサンドボックス化されたアクセスを行います。すべてのファイル操作はプロジェクト ルートにロックされ、パスのトラバーサルはブロックされます。
ツール/関数呼び出しサポートを備えたモデルは、MandoCode で最もよく機能します。初回実行ウィザードでは、まさに以下のモデルが提供されます。クラウドのデフォルトを自動で取得するか、ハードウェアに一致するローカル層を選択できます。
クラウド (GPU は必要ありません — Ollama のサーバー上で実行され、 ollama サインインで無料):
ローカル (完全にオフライン、ハードウェア上で実行):
MandoCode は起動時にモデルの互換性を検証します。 /learn を実行してモデルのサイズとハードウェア要件に関する詳細なガイドを確認するか、/setup を実行して切り替えます。

[切り捨てられた]

## Original Extract

A .NET C# CLI Coding Agent powered by Ollama + Semantic Kernel and RazorConsole. Run locally or in the cloud. Refactors code, proposes diffs, and updates your project safely — no API keys required. - DevMando/MandoCode

I built MandoCode, an open source CLI coding agent in .NET that runs against local Ollama models. No API keys, nothing leaves your machine (Ollama cloud models work too if you want them). It reads, edits, searches, plans, and browses the web across your codebase, and it supports MCP and Skills. Install with dotnet tool install -g MandoCode. I found semantic Kernel great to give me the ability to just go and build on .NET C#.
Lots of more work to do!

GitHub - DevMando/MandoCode: A .NET C# CLI Coding Agent powered by Ollama + Semantic Kernel and RazorConsole. Run locally or in the cloud. Refactors code, proposes diffs, and updates your project safely — no API keys required. · GitHub
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
DevMando
/
MandoCode
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
157 Commits 157 Commits .github .github .mandocode/ skills/ commit-message .mandocode/ skills/ commit-message docs docs src/ MandoCode src/ MandoCode tests/ MandoCode.Tests tests/ MandoCode.Tests .gitattributes .gitattributes .gitignore .gitignore LICENSE LICENSE MandoCode.sln MandoCode.sln README.md README.md View all files Repository files navigation
Your AI coding assistant — run locally or in the cloud with Ollama.
No API keys required. Just you and your code.
MandoCode is an AI coding assistant built on RazorConsole , powered by Semantic Kernel and Ollama . RazorConsole makes the entire terminal UI possible — Razor components, a virtual DOM, and Spectre.Console rendering all running in the console.
Run locally or connect to Ollama cloud — no API keys required for anything, including web search (an optional free Tavily key upgrades search reliability). It gives you Claude-Code-style project awareness — reading, writing, searching, planning, and web browsing across your entire codebase — without ever leaving your terminal. It understands any file type : C#, JavaScript, TypeScript, Python, CSS, HTML, JSON, config files, and more.
.NET 8 SDK — dotnet.microsoft.com/download/dotnet/8.0 (SDK includes the runtime — install only the SDK)
Ollama — ollama.com/download (MandoCode walks you through setup on first run)
dotnet tool install -g MandoCode
mandocode
First run launches a guided wizard: it detects Ollama, offers to start it, walks you through cloud sign-in if you'd like more powerful models, and auto-pulls a sensible default. You can re-run it any time with /setup .
mandocode --doctor
Prints your runtime version, Ollama status, models pulled, and cloud sign-in state.
⚠️ Local models: check your Ollama context window (the #1 gotcha)
Using cloud models ( :cloud tags)? Skip this section. Cloud model context is managed on Ollama's servers and set to the model's maximum by default — nothing on your machine affects it, including the desktop app's slider.
If you use local models and see responses cut off, the model "forgetting" earlier conversation, edits failing repeatedly on files it just wrote, or this message:
⚠ Response was cut off because the model's CONTEXT WINDOW filled …
…your Ollama context window is almost certainly too small. The context window is how much conversation + code the model can see at once — and Ollama defaults it to ~4k tokens , which an agentic session fills almost immediately. When it overflows, the oldest content (including the system prompt — the model's instructions!) is silently dropped.
If you use the Ollama desktop app (the tray icon), the app's Settings → Context length slider controls this — and it overrides everything else, including MandoCode's config:
There's no universally right slider position — it's a trade between how much the model can see and fitting in your GPU's memory (every 8k of window costs roughly 0.5–1.5 GB of VRAM depending on the model):
Too low (the 4k default): the symptoms above — the model's own instructions silently fall out of the window and it stops behaving.
Too high for your GPU : the model spills into system RAM, tokens/sec craters, and turns crawl or look hung.
Starting points : 16k for most GPUs, 32k with 8 GB+ VRAM. Only raise it if you're seeing the symptoms above; step back down a notch if generation slows badly after raising it.
If you run ollama serve yourself (no desktop app), MandoCode handles it: it sets OLLAMA_CONTEXT_LENGTH from your contextLength config when it starts the daemon, and auto-sizes it to the hardware tier of the model you pick in /setup or /model . Tune it manually with:
mandocode --config set contextLength 16384
Verify what your daemon is actually using with ollama ps (look at the CONTEXT column). Run /learn inside MandoCode for a friendly explainer.
⚠️ All models: check your response cap (the #2 gotcha)
The context window's evil twin — and unlike the slider above, this one applies to every model, cloud included . If the model announces work and then just stops — "I'll create the game…" and the turn ends with no plan, no files, and no error — your maxTokens is too low. It caps a single reply ( NumPredict ), and reasoning models spend output tokens thinking before they emit a tool call, so a low cap cuts them off before they ever act.
Fresh installs default to 32k and never notice it. But if your config predates v0.11, or you once lowered maxTokens thinking it was the context window (they're different knobs — this caps what the model says , the context window caps what it sees ), check it:
mandocode --config show # look at "Max Tokens"
mandocode --config set maxTokens 32768
The telltale sign: token tracking shows output pinned at exactly your cap, turn after turn (e.g. 2k out every time). Note that a running session keeps the config it loaded at startup — restart MandoCode (or use /config set in-app) for the change to take effect.
git clone https://github.com/DevMando/MandoCode.git
cd MandoCode
dotnet build src/MandoCode/MandoCode.csproj
dotnet run --project src/MandoCode/MandoCode.csproj -- /path/to/your/project
What Makes MandoCode Different
Safe File Editing with Diff Approvals
Every file write and delete is intercepted with a color-coded diff. You approve, deny, or redirect — nothing touches disk without your say-so.
Type @ to autocomplete any project file and attach it as context. The AI sees the full file content alongside your prompt. Reference multiple files in a single message.
Complex requests are automatically broken into step-by-step plans. Review the plan, then watch each step execute with progress tracking.
The AI can search the web and read webpages to find documentation, tutorials, or answers — no API keys needed. Optionally add a free Tavily key for AI-optimized search that doesn't hit DuckDuckGo's rate limits.
Lofi and synthwave tracks bundled right in. A waveform visualizer runs in the corner while you code. Because vibes matter.
If Ollama isn't running, MandoCode shows setup guidance inline instead of a bare error. Use /retry to reconnect without restarting.
Feature
Description
AI
Project-aware assistant
Reads, writes, deletes, and searches your entire codebase
AI
Web search & fetch
Web search and webpage reading — keyless via DuckDuckGo, or Tavily with a free API key
AI
MCP server support
Connect to any Model Context Protocol server (stdio or remote HTTP) — Claude-Desktop-compatible config
AI
Streaming responses
Real-time output with animated spinners
AI
Task planner
Auto-detects complex requests and breaks them into steps
AI
Fallback function parsing
Handles models that output tool calls as raw JSON
UI
Diff approvals
Color-coded diffs with approve / deny / redirect
UI
Markdown rendering
Rich terminal output — headers, tables, code blocks, quotes
UI
Syntax highlighting
C#, Python, JavaScript/TypeScript, Bash
UI
Clickable file links
OSC 8 hyperlinks for file paths
UI
Terminal theme detection
Auto-adapts colors for light and dark terminals
UI
Taskbar progress
Windows Terminal integration during task execution
Input
/ command autocomplete
Slash commands with dropdown navigation
Input
@ file references
Attach file content to any prompt
Input
! shell escape
Run shell commands inline ( !git status , !ls )
Input
/copy and /copy-code
Copy responses or code blocks to clipboard
Music
Lofi + synthwave
Bundled tracks with volume, genre switching, waveform visualizer
Config
Configuration wizard
Guided setup with model selection and connection testing
Config
Config validation
Auto-clamps invalid settings to safe ranges
Reliability
Retry + deduplication
Exponential backoff and duplicate call prevention
Education
/learn command
LLM education guide with optional AI educator chat
Commands
Type / to see the autocomplete dropdown, or ! to run a shell command.
/setup — first-run wizard, guided. Detects Ollama, offers to install it, walks you through cloud sign-in, picks a model with hardware-aware tiers, auto-pulls a sensible default. Use when something's broken or you're a newcomer.
/model — quick switch. Pick a model from your pulled list and go — local picks get a context window sized to their hardware tier automatically. Use when you just want to swap models.
/config — adjust settings. Full configuration form covering temperature, timeouts, ignore dirs, etc. Use when you know exactly what knob you want to turn.
CLI flags (outside the chat loop)
mandocode --doctor # preflight check: .NET runtime, Ollama status, models, sign-in
mandocode --config show # print current config
mandocode --config init # create a default config file
mandocode --config set < key > < value > # set a single value (e.g. set model qwen3.5:9b)
mandocode --config path # show config file location
Run mandocode --doctor any time chat is misbehaving — exits 0 if everything's green, 1 if anything's missing, with a clear summary of what's wrong.
You type a prompt
|
MandoCode adds project context (@files, system prompt)
|
Semantic Kernel sends to Ollama (local or cloud model)
|
AI responds with text + function calls
|
File operations go through diff approval
Web searches and fetches run directly
|
Rich markdown rendered in your terminal
The AI has sandboxed access to your project through a FileSystemPlugin (9 functions: list files, glob search, read, write, delete files/folders, text search, path resolution) and a WebSearchPlugin (web search via Tavily or DuckDuckGo, webpage fetching — works without any API key). All file operations are locked to your project root — path traversal is blocked.
Models with tool/function calling support work best with MandoCode. The first-run wizard offers exactly the models below — auto-pulls the cloud default, or lets you pick a local tier matched to your hardware.
Cloud (no GPU required — runs on Ollama's servers, free with ollama signin ):
Local (fully offline, runs on your hardware):
MandoCode validates model compatibility on startup. Run /learn for a detailed guide on model sizes and hardware requirements, or /setup to switch between

[truncated]
