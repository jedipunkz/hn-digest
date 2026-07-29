---
source: "https://github.com/DevMando/MandoCode.Desktop"
hn_url: "https://news.ycombinator.com/item?id=49099891"
title: "Show HN: MandoCode Desktop – Native Windows AI Coding Assistant on Ollama .NET"
article_title: "GitHub - DevMando/MandoCode.Desktop: The MandoCode AI coding agent as a native Windows app — WinUI 3 over the same open-weight-model engine as the CLI (Ollama + Semantic Kernel). Review diffs before they land, run plans step-by-step. Local or hosted, no API keys, never locked to one vendor's model.\n[truncated]"
author: "devmando"
captured_at: "2026-07-29T17:04:45Z"
capture_tool: "hn-digest"
hn_id: 49099891
score: 1
comments: 1
posted_at: "2026-07-29T16:48:12Z"
tags:
  - hacker-news
  - translated
---

# Show HN: MandoCode Desktop – Native Windows AI Coding Assistant on Ollama .NET

- HN: [49099891](https://news.ycombinator.com/item?id=49099891)
- Source: [github.com](https://github.com/DevMando/MandoCode.Desktop)
- Score: 1
- Comments: 1
- Posted: 2026-07-29T16:48:12Z

## Translation

タイトル: Show HN: MandoCode Desktop – Ollama .NET 上のネイティブ Windows AI コーディング アシスタント
記事のタイトル: GitHub - DevMando/MandoCode.Desktop: ネイティブ Windows アプリとしての MandoCode AI コーディング エージェント — CLI (Ollama + セマンティック カーネル) と同じオープンウェイト モデル エンジン上の WinUI 3。着地する前に差分を確認し、計画を段階的に実行します。ローカルまたはホストされ、API キーがなく、特定のベンダーのモデルにロックされることはありません。
[切り捨てられた]
説明: ネイティブ Windows アプリとしての MandoCode AI コーディング エージェント — CLI (Ollama + セマンティック カーネル) と同じオープンウェイト モデル エンジン上の WinUI 3。着地する前に差分を確認し、計画を段階的に実行します。ローカルまたはホストされ、API キーがなく、特定のベンダーのモデルにロックされることはありません。 - DevMando/MandoCode.Desktop
HN テキスト: このバージョンが誕生した経緯についての簡単な裏話: 元の MandoCode は CLI であり、その UI は RazorConsole (ターミナルにレンダリングされる Blazor コンポーネント) で構築されています。まさに賢い技術！しかし、私はタブが欲しかったので、今年 MS Build に参加したとき、WinUI アプリで何ができるのかをよく見て、元のエージェントに実際のウィンドウとタブを提供できる可能性があることに気づきました。 1 つのウィンドウで最大 4 つのエージェントを実行できるのは非常に便利です。それはさておき、タブです！最近はすべてがチャットになっているように感じますが、どういうわけか私の脳はチャットではなくエージェントを叫び続けています。私はタブ付きブラウザのエクスペリエンスがもっと好きです。配管が機能したら、いくつかいじってみました。 - コンテキスト スナップショット - エージェントを閉じると会話が削除されずにアーカイブされ、会話中にスナップショットを作成して AI が作成した要約を取得し、別のモデルまたは新しいエージェントに渡してスレッドを取得させることができます。スナップショットを作成するときに、ドロップダウンから LLM を選択します。 - トランスクリプトのエクスポート — すべてのエージェントのタブで会話全体をダンプできます

ファイルに。 - 16 のテーマ — これは、VS Code やその他の IDE にある種類のテーマを構築することで、私の子供が単純に楽しんでいた場所です。 Dracula、Tokyo Night、One Dark Pro、すべての上品なものに加えて、当時の正確な Windows 98 デスクトップとブラウン管 CRT モードが含まれています。 16 個のいずれも着かない場合は、独自の背景画像もドロップします。 - 意図的に制限されたメモ パッド — 常にそこにあり、プロジェクトとは別に、オプションの AI ヘルプがありますが、アシスタントにはファイル ツールがありません。返信がメモに表示される唯一の方法は、ボタンをクリックしてメモに返信することです。 - 小さな内蔵音楽プレーヤー — 雰囲気を楽しむために（笑）。 Lofi/synthwave をバンドルするか、独自の MP3 フォルダーを指定します。全体的に、WinUI XD 内での作業が大好きでした

記事本文:
GitHub - DevMando/MandoCode.Desktop: ネイティブ Windows アプリとしての MandoCode AI コーディング エージェント — CLI (Ollama + セマンティック カーネル) と同じオープンウェイト モデル エンジン上の WinUI 3。着地する前に差分を確認し、計画を段階的に実行します。ローカルまたはホストされ、API キーがなく、特定のベンダーのモデルにロックされることはありません。 · GitHub
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
あなた

別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
デブマンド
/
MandoCode.デスクトップ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
96 コミット 96 コミット .github/ workflows .github/ workflows MandoCode @ cd4d003 MandoCode @ cd4d003 docs docs src src .gitignore .gitignore .gitmodules .gitmodules CHANGELOG.md CHANGELOG.md ライセンス ライセンスMandoCode.Desktop.slnx MandoCode.Desktop.slnx README.md README.md RELEASING.md RELEASING.md THIRD-PARTY-NOTICES.txt THIRD-PARTY-NOTICES.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング アシスタントは、他人のクローズド モデルではなく、オープンウェイト モデルに基づいて構築されています。
ほとんどの AI コーディング ツールは、ベンダーが提供するものであればいつでも、1 つのベンダーのクローズド モデルに閉じ込めてしまいます。
彼らの条件に従ってそれを変更してください。 MandoCode は、代わりにオープンウェイト モデルで実行されます。
Ollama : 完全に自分のマシン上で — 無料、プライベート、API キーなし、何もなし
トークンごとに計測されます。または、より多くのヘッドルームが必要な場合はホストされたモデルを使用します。どちらにしても同じアシスタント、同じ
会話、あなたの呼びかけ。
MandoCode Desktop はそのアシスタントを取得し、単一のアシスタントではなく Windows 上に実際のホームを与えます。
ターミナルに固定されたチャット ウィンドウ: 並行して動作する複数のエージェント、実際に統合されたシェル、
git 対応のファイル ブラウザ、メモ帳、テーマに十分な個性があり、その必要はありません
タスクバー上の他の開発ツールと同じように見えます。
MandoCode CLI とまったく同じエンジンに基づいて構築されています。
文字通り同じコードで、サブモジュールとして固定されているため、考え方に違いはありません。
インターフェースのみです。
オープン モデルはあなたの選択です。決して 1 つのベンダーのものではありません。 MandoCode はオープンウェイト モデルで実行されます。
オル

ama: エージェントごとにモデルを交換したり、完全に自分のハードウェア上で無料で実行したり、
さらに上限が必要な場合は、ホストされたモデルを使用します。単一のクローズド API にロックされることはありません。
コーディングアシスタントは、ポップアップではなく、一日中使用するものです。つまり、本物のネイティブになります
C#、WinUI 3、およびセマンティック カーネルの Ollama コネクタ上に構築されたアプリ - 複数のエージェントが同時に開き、
本物のシェル、実際に git について知っているファイル ツリー、何も考えずに考えを書き留める場所
テキスト エディターを開くと、1 つの固定されたダーク テーマの代わりに独自の外観を作成できます。
エージェント タブ、および一度に最大 4 つの分割ビュー — 各タブは独自に独立しています
会話、プロジェクト フォルダー、モデル。分割ビューでは 2 つ、3 つ、または 4 つを並べて表示できるため、
複数のエージェントが一度に子守をするのではなく、並行して作業する様子を監視できます。
実際の統合ターミナル — PowerShell 7、Windows PowerShell、cmd、Git Bash、または WSL が実行されている
(偽のコンソールではなく) 実際のシェルとして、アクティブなエージェントのプロジェクト フォルダーで開かれます。
Git 対応のファイル エクスプローラー — ブランチ、ステータス、ダーティ バッジを含むライブ ファイル ツリー、インライン diff
カードを作成し、ワンクリックでコミットし、チャットに直接ドラッグして参照できます。
コンテキスト スナップショットとセッション履歴 - エージェントを閉じると、会話はアーカイブされずにアーカイブされます。
それを削除する。モデルがサポートしているときに、過去の会話をそのトランスクリプトとともに後で再度開きます。
それ、その完全な記憶。スナップショットを使用すると、AI が作成した 1 つの会話の要約を次の会話に持ち込むことができます。
まったく異なるモデルまたは新しいエージェント。
メモ — プロジェクトから独立した、オプションの AI ヘルプを備えた、常に存在するジョット パッド。の
ここではアシスタントにはファイル ツールがまったくありません。返信がメモに届く唯一の方法は、ボタンを押すことです。
あなたなしでは何も書かれないように、押してください。
スキルと MCP — アシスタントに新しい再利用可能な機能を教えます (新しい機能をインストールします)

フォルダーまたは
zip を使用するか、独自に作成させて)、MCP 経由で外部ツールに接続します。
16 個の内蔵テーマ — ドラキュラ、トーキョー ナイト、ワン ダーク プロから真に当時のものまで
Windows 98 デスクトップとちらつく陰極線 CRT 管、そしてその背後にある独自の背景画像
チャット。
小さな内蔵音楽プレーヤー - lofi と synthwave がバンドルされているか、任意のフォルダーにそれをポイントします。
自分のMP3。
ガイド付きの初回実行セットアップ - 初回起動時に Ollama を検索 (またはインストール) し、サポートします。
チャット内でスターター モデルを選択してください。挨拶する前に設定ファイルを手動で編集する必要はありません。
最新の MandoCode.Desktop-*-win-x64.zip を次の場所からダウンロードします。
をリリースし、任意の場所に抽出して実行します
MandoCode.Desktop.exe 。 zip は完全に自己完結型であり、.NET のインストールは必要ありません。
最初の起動時に、アプリはチャット内でガイド付きセットアップを実行します。Ollama を見つけます (
見つからない場合は、winget 経由でインストールします)、デーモンを起動し、最初のモデルを選択するのに役立ちます。
クラウド モデル (最高品質、GPU 不要、無料の ollama.com サインイン)、または短期間のローカル モデル
サイズとハードウェアのヒントを含むリスト。 /setup またはガイド付き実行を使用すると、いつでもウィザードを再実行できます。
設定のセットアップボタン。
要件: WebView2 ランタイムを備えた Windows 10 (1809+) または Windows 11 - プレインストールされています。
Windows 11 と Windows 10 上の Edge によって最新の状態が維持されます。
git clone --recursive https://github.com/DevMando/MandoCode.Desktop.git
cd MandoCode.デスクトップ
dotnet build src/MandoCode.Desktop/MandoCode.Desktop.csproj
dotnet run --project src/MandoCode.Desktop
--recursive を使用せずにすでにクローンが作成されていますか? MandoCode/ フォルダーは空になります — 実行します
git submodule update --init を実行して再度ビルドします。
MandoCode.Desktop.exe <フォルダ> は、そのフォルダをプロジェクト ルートとして開きます (それ以外の場合は、現在の
ディレクトリ;フォルダボタンを使用してアプリ内で変更可能）。要求する

es は到達可能なオラマ (オラマ サーブ) です。
CLI と同じ構成ファイルを使用するため、両方のアプリがエンドポイント/モデル/設定を共有します。
MandoCode Desktop は、AI サービス、タスク プランナー、プラグイン、MCP、CLI のハーネス全体を再利用します。
スキル、構成、承認、トークン追跡 — 固定された git サブモジュールとして。界面層のみ
(エージェント、分割ビュー、ターミナル、メモ、テーマ、その他これをネイティブ アプリにするすべてのもの
コンソールではなく) は新しいものです。さらに詳しく知りたい場合は、ハーネス サブモジュールがどのように固定されているか、
ロールフォワード、エージェントと分割ビューがどのように接続されるか、スナップショット/メモ/履歴が実際にどのように保持されるか -
docs/ARCHITECTURE.md を参照してください。
MIT — アルマンド フェルナンデス – DevMando。
ネイティブ Windows アプリとしての MandoCode AI コーディング エージェント - CLI (Ollama + セマンティック カーネル) と同じオープンウェイト モデル エンジン上の WinUI 3。着地する前に差分を確認し、計画を段階的に実行します。ローカルまたはホストされ、API キーがなく、特定のベンダーのモデルにロックされることはありません。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

The MandoCode AI coding agent as a native Windows app — WinUI 3 over the same open-weight-model engine as the CLI (Ollama + Semantic Kernel). Review diffs before they land, run plans step-by-step. Local or hosted, no API keys, never locked to one vendor's model. - DevMando/MandoCode.Desktop

Quick backstory on how this version came to be: the original MandoCode is a CLI, and its UI is built with RazorConsole — Blazor components rendered into a terminal. Genuinely clever tech! I wanted tabs, though, and when I attended MS Build this year I got a good look at what WinUI apps could do and saw the potential to give the original agent some actual windows and tabs. Having up to 4 agents running in one window is pretty cool. Aside from that — tabs! I feel like everything is a chat these days, and for some reason my brain keeps screaming AGENTS, not chats, at me. I love the tabbed-browser experience a lot more Some of the stuff I got to toy with once the plumbing worked: - Context snapshots — closing an agent archives the conversation instead of deleting it, and you can snapshot mid-conversation to get an AI-written recap you can hand to a different model or a fresh agent and have it pick up the thread. You pick the LLM from a dropdown right when you create the snapshot. - Transcript export — every agent tab can dump its full conversation to a file. - 16 themes — this is where the kid in me just had plain fun, building the kind of themes you find in VS Code and other IDEs. Dracula, Tokyo Night, One Dark Pro, all the tasteful ones, plus a period-accurate Windows 98 desktop and a Cathode Ray CRT mode. Drop in your own background image too if none of the 16 land. - A deliberately limited Notes pad — always there, separate from any project, optional AI help, but the assistant has zero file tools in it. The only way a reply lands in your note is you clicking a button to put it there. - A tiny built-in music player — for the vibes, lol. Lofi/synthwave bundled, or point it at your own MP3 folder. Overall I loved working within WinUI XD

GitHub - DevMando/MandoCode.Desktop: The MandoCode AI coding agent as a native Windows app — WinUI 3 over the same open-weight-model engine as the CLI (Ollama + Semantic Kernel). Review diffs before they land, run plans step-by-step. Local or hosted, no API keys, never locked to one vendor's model. · GitHub
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
DevMando
/
MandoCode.Desktop
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
96 Commits 96 Commits .github/ workflows .github/ workflows MandoCode @ cd4d003 MandoCode @ cd4d003 docs docs src src .gitignore .gitignore .gitmodules .gitmodules CHANGELOG.md CHANGELOG.md LICENSE LICENSE MandoCode.Desktop.slnx MandoCode.Desktop.slnx README.md README.md RELEASING.md RELEASING.md THIRD-PARTY-NOTICES.txt THIRD-PARTY-NOTICES.txt View all files Repository files navigation
An AI coding assistant built on open-weight models, not someone else's closed one.
Most AI coding tools lock you into one vendor's closed model — whatever they ship, whenever they
change it, on their terms. MandoCode runs on open-weight models instead, through
Ollama : entirely on your own machine — free, private, no API key, nothing
metered per token — or a hosted model when you want more headroom. Same assistant either way, same
conversation, your call.
MandoCode Desktop takes that assistant and gives it an actual home on Windows, instead of a single
chat window bolted onto a terminal: several agents working in parallel, a real integrated shell, a
git-aware file browser, a notes pad, and enough personality in the theming that it doesn't have to
look like every other dev tool on your taskbar.
It's built on the exact same engine as the MandoCode CLI —
literally the same code, pinned in as a submodule — so nothing about how it thinks is different.
Only the interface is.
Open models, your choice — never one vendor's. MandoCode runs on open-weight models through
Ollama: swap models per agent, run one entirely on your own hardware for free, or reach for a
hosted model when you want more ceiling. You're never locked to a single closed API.
A coding assistant is something you live in all day, not a popup. So it gets a real native
app built on C#, WinUI 3, and Semantic Kernel's Ollama connector — multiple agents open at once,
a real shell, a file tree that actually knows about git, a place to jot down a thought without
opening a text editor, and a look you can make your own instead of one fixed dark theme.
Agent tabs, and Split view for up to four at once — each tab is its own independent
conversation, project folder, and model. Split view puts two, three, or four side by side so you
can watch several agents work in parallel instead of babysitting one at a time.
A real integrated terminal — PowerShell 7, Windows PowerShell, cmd, Git Bash, or WSL, running
as an actual shell (not a fake console), opened in the active agent's project folder.
Git-aware file explorer — a live file tree with branch, status, and dirty badges, inline diff
cards, one-click commit, and drag-to-reference straight into the chat.
Context snapshots & session history — closing an agent archives its conversation instead of
deleting it; reopen any past conversation later with its transcript and, when the model supports
it, its full memory. Snapshots let you carry an AI-written recap of one conversation into a
completely different model or a fresh agent.
Notes — an always-there jot pad, separate from any project, with optional AI help. The
assistant has no file tools here at all — the only way a reply reaches your note is a button you
press, so nothing gets written without you.
Skills & MCP — teach the assistant new, reusable capabilities (install one from a folder or a
zip, or have it write its own), and connect external tools over MCP.
16 built-in themes — from Dracula, Tokyo Night, and One Dark Pro to a genuinely period-correct
Windows 98 desktop and a flickering Cathode Ray CRT tube, plus your own background image behind
the chat.
A tiny built-in music player — lofi and synthwave come bundled, or point it at any folder of
your own MP3s.
Guided first-run setup — on first launch it finds (or installs) Ollama for you and helps you
pick a starter model, right in the chat. No config file to hand-edit before you can say hello.
Download the latest MandoCode.Desktop-*-win-x64.zip from
Releases , extract it anywhere, and run
MandoCode.Desktop.exe . The zip is fully self-contained — no .NET install required .
On first launch the app runs a guided setup right in the chat: it finds Ollama (offering to
install it via winget if it's missing), starts the daemon, and helps you pick a first model — a
cloud model (best quality, no GPU needed, free ollama.com sign-in) or a local one from a short
list with size and hardware hints. Re-run the wizard any time with /setup or the Run guided
setup button in Settings.
Requirements: Windows 10 (1809+) or Windows 11, with the WebView2 runtime — preinstalled on
Windows 11 and kept current by Edge on Windows 10.
git clone --recursive https://github.com/DevMando/MandoCode.Desktop.git
cd MandoCode.Desktop
dotnet build src/MandoCode.Desktop/MandoCode.Desktop.csproj
dotnet run --project src/MandoCode.Desktop
Already cloned without --recursive ? The MandoCode/ folder will be empty — run
git submodule update --init and build again.
MandoCode.Desktop.exe <folder> opens with that folder as the project root (otherwise the current
directory; changeable in-app via the folder button). Requires a reachable Ollama ( ollama serve ).
Uses the same config file as the CLI, so both apps share endpoint/model/settings.
MandoCode Desktop reuses the CLI's entire harness — the AI service, task planner, plugins, MCP,
skills, config, approvals, token tracking — as a pinned git submodule. Only the interface layer
(agents, split view, the terminal, notes, theming, and everything else that makes this a native app
rather than a console) is new. If you want the deep dive — how the harness submodule is pinned and
rolled forward, how agents and split view are wired, how snapshots/notes/history actually persist —
see docs/ARCHITECTURE.md .
MIT — Armando Fernandez - DevMando.
The MandoCode AI coding agent as a native Windows app — WinUI 3 over the same open-weight-model engine as the CLI (Ollama + Semantic Kernel). Review diffs before they land, run plans step-by-step. Local or hosted, no API keys, never locked to one vendor's model.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
