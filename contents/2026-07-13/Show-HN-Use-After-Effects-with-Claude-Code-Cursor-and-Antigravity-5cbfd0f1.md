---
source: "https://github.com/Arman-Luthra/aftr"
hn_url: "https://news.ycombinator.com/item?id=48886809"
title: "Show HN: Use After Effects with Claude Code, Cursor and Antigravity"
article_title: "GitHub - Arman-Luthra/aftr: Puppeteer for After Effects. Use AE with Claude Code to make production-ready videos. · GitHub"
author: "armanluthra_"
captured_at: "2026-07-13T01:50:10Z"
capture_tool: "hn-digest"
hn_id: 48886809
score: 4
comments: 2
posted_at: "2026-07-13T01:35:35Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Use After Effects with Claude Code, Cursor and Antigravity

- HN: [48886809](https://news.ycombinator.com/item?id=48886809)
- Source: [github.com](https://github.com/Arman-Luthra/aftr)
- Score: 4
- Comments: 2
- Posted: 2026-07-13T01:35:35Z

## Translation

タイトル: HN を表示: クロード コード、カーソル、反重力で After Effects を使用する
記事のタイトル: GitHub - Arman-Luthra/aftr: Puppeteer for After Effects。 AE と Claude Code を使用して、制作可能なビデオを作成します。 · GitHub
説明: After Effects のパペッティア。 AE と Claude Code を使用して、制作可能なビデオを作成します。 - アルマン・ルトラ/アフター

記事本文:
GitHub - Arman-Luthra/aftr: After Effects 用のパペッティア。 AE と Claude Code を使用して、制作可能なビデオを作成します。 · GitHub
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
アルマン・ルトラ
/
その後
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション

オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
10 コミット 10 コミット .github/ workflows .github/ workflows bin bin コントローラ コントローラ ドキュメント ドキュメントの例 例 パッケージ化/ pip パッケージ化/ pip パネル パネル プリセット プリセット 共有 共有シミュレータ シミュレータ ツール ツール .dockerignore .dockerignore .gitignore .gitignore .npmignore .npmignore Dockerfile Dockerfile ライセンス ライセンスREADME.md README.md config.json config.json package-lock.json package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
After Effects と Claude Code を使用して、制作可能なビデオを作成します。
ノード コントローラーは WebSocket 経由で JSON コマンドを AE 内の CEP パネルに送信し、コマンドは ExtendScript として実行され、JSON が返されます。上部には、クリップを構築し、レンダリング、レビュー、改訂によってクリップを自己修正する自律的な仕様主導のパイプラインが配置されています。すべてはソケットを介してプログラムによって駆動され、手動による After Effects 作業は必要ありません。
タイトル シーケンスは完全に aftr を通じて構築およびレンダリングされます。 After Effects の手動作業はありません。
ティリオン/フォートレス用に作成されたビデオ。
完全なデモを見る
2 つのコマンドで aftr をクロード コードに接続します。コントローラーを起動し (MCP エンドポイントとして機能します)、Claude Code をコントローラーに向けます。
npx aftr-studio コントローラー # インストールなし、Node 18+ が必要
pip install aftr-studio && aftr コントローラー # pip
docker run -p 8787:8787 ghcr.io/arman-luthra/aftr # docker
クロード mcp add --transport http aftr http://127.0.0.1:8787/mcp
After Effects とパネルを開きます（[ウィンドウ] > [拡張機能] > [aftr]。npm rundeploy:panel を使用してクローンから一度デプロイします）。次に、クロード コードにわかりやすい言葉で尋ねます。
1080p 5s コンプを作成し、後ろに火のあるタイトル「LAUNCH」を追加し、blurFade をアニメーション化して、mp4 にレンダリングします。
Claude Code はすべてのコマンドをツールとして認識し、ae_ を呼び出します。

最初に status を実行し、 ae_list_commands でセットをリストし、 ae_command を通じてあらゆるものを実行します。標準入出力サーバーの方が良いですか? claude mcp add aftr -- npx -y aftr-studio mcp を使用します (コントローラーも実行したままにします)。
方法
コマンド
npm (グローバル CLI)
npm install -g aftr-studio 、次に aftr コントローラー
npx (インストールなし)
npx aftr-studio コントローラー
pip (Python ランチャー)
pip install aftr-studio 、次に aftr コントローラー (Node 18+ が必要)
ドッカー
docker run --rm -p 8787:8787 ghcr.io/arman-luthra/aftr
ソースから
このリポジトリのクローンを作成します (下記)
npm、npx、および Docker パスは、コントローラー、MCP サーバー ( aftr mcp )、およびシミュレーター ( aftr sim ) を実行します。 CEP パネルの After Effects へのデプロイは、拡張機能が自己署名してインストールされるため、クローン ( npm rundeploy:panel ) から行われます。 pip パッケージは、 npx を通じて npm CLI に転送するシン ランチャーです。
git クローン https://github.com/Arman-Luthra/aftr.git
CDアフター
npmインストール
# 1) After Effects を使用せずに正常であることを証明する (ヘッドレス シミュレーター + テスト)
npm run build:jsx && npm test
# 2) パネルを AE にインストールします (自己署名 + Windows および macOS に展開)
npm rundeploy:panel # その後、AE を終了して再起動します。[ウィンドウ] > [拡張機能] > [aftr]
# 3) ブリッジを開始 + UI を開く
npm 実行コントローラー # http://127.0.0.1:8787
次に、UI、REST、WebSocket から、またはライブラリとしてそれを駆動します。
ノード例/flaming-title.mjs " ON FIRE "
After Effects 2024 ～ 2026、Node 18+、および PATH 上の ffmpeg が必要です。 Windows と macOS の完全なセットアップについては、セクション 4 を参照してください。
能力
何をするのか
~100 コマンド
コンプ、レイヤー、イージング付きキーフレーム、エクスプレッション、エフェクト (深くネストされたパラメーターやサードパーティのプラグインを含む、matchName または表示名による)、マスク、レイヤー スタイル、テキスト アニメーター、レンダリング。
発見
listFonts 、 listInstalledEffects (matchNames を返す)、 findEffectMatchName 、 listPlugins 、 getEnvir

オンメント。推測するのではなく、何がインストールされているかを見つけてください。
ワンコール VFX
fireEffect 、smokeEffect 、 glitchEffect 、 neonGlow 、 cinematicGrade 、およびフレンドリーな applyLumetri グレード。
テキストアニメーション
4 つの Animate パネル プリセット ( wordReveal 、 charScale 、 BunchRotate 、 BlurFade ) とフル レンジ セレクター ビルダー。
自律型パイプライン
宣言的なセグメント仕様は、実現、レンダリング、視覚的なレビューの取得、構造化されたデルタの取得、再レンダリング、および連結を行います。自己修正してくれます。
MCPサーバー
Claude Desktop、Claude Code、または任意の MCP クライアントから AE を駆動します。すべてのコマンドはツールです ( npm run mcp )。
高速かつバッチ処理
バッチは、1 回のラウンドトリップと 1 回のアンドゥで N 回の編集を実行します (20 操作で 200 ミリ秒未満)。ノンブロッキングの aerender ストリームは、並列レンダリング エンジンを使用してイベントを進行します。
AEなしでテスト可能
ヘッドレス シミュレーターは、模擬 DOM に対して実際の JSX を実行します。 96 のテストに加えて、ノード 18、20、および 22 での e2e、CI。
クロスプラットフォーム
Windows と macOS では、1 つのコマンドで署名して展開できます。
仕組み
┌─────────────────────┐
│ コントローラ (ノード + ws + Express + Web UI) │ <- あなた / UI / エージェントがこれを制御します
│ • WebSocket サーバー (パネル + エージェント) をホストします │
│ • REST サーフェス + インタラクティブ UI │
│ • HLD オーケストレーター (エージェント ループ) │
━━━━━━━━━━━━━━━━━━━━━━━┘
│ WebSocket (JSON リクエスト/レスポンス/イベント エンベロープ)
┌─────────▼───────────────┐
│ CEP パネル (After Effects 内) │ ┐
│ • WebSocket CLIENT、コントローラにダイヤルします │ │ 1 回ビルドし、
│ • ルートのカンマ

JSX ディスパッチ層への nds │ │ Win と Mac で同一
│ • ノンブロッキングレンダーのために aerender を生成します │ │
━━━━━━━━━━━━━━━━━━━━━━━┘ │
│ evalScript (文字列入力 / JSON 文字列出力)
┌───────────▼───────────────┐ │
│ JSX コマンド層 (host.jsx + Commands/*.jsx) │ │
│ •dispatch(command, paramsJson) -> JSON │ │
│ • ~90 コマンド: レイヤー、キーフレーム、エフェクト、 │ │
│ テキストアニメーター、マスク、VFX プリセット、レンダリング │ │
━━━━━━━━━━━━━━━━━━━━━━┘
│ AE スクリプティング DOM (app.project、comps、layers)
┌─────────▼───────────────┐
│ アフターエフェクト │
━━━━━━━━━━━━━━━━━━━┘
目次
クイックスタート、AE不要（シミュレータ）
実際の After Effects を使用した完全なセットアップ
窓
オーケストレーター (自律型パイプライン)
開発ワークフロー (ホットリロード + 署名)
コードから AE を制御: コンプとレイヤー (ソリッド、テキスト、シェイプ、ヌル、カメラ、ライト、調整、フッテージ) の作成、トランスフォームまたはエフェクトのプロパティの設定、イージングを使用したキーフレームの追加、エクスプレッションの作成、レイヤーの親、トリム、移動、または複製をすべてソケット上で行います。
matchName によってエフェクトを追加し、キーフレームやエクスプレッションを介してアニメーション化された、深くネストされたパラメータであっても、そのパラメータを操作します。
あ

4 つの既製のアニメーション パネル スタイル ( wordReveal 、 charScale 、 BunchRotate 、 BlurFade ) とフル テキスト アニメーター ビルダー (Based On、Shape、Ease High to Low、キーフレーム化された Offset および Start) を備えた nimate テキスト。
ワンコールの VFX プリセットにアクセスします: fireEffect (リアルな火)、smokeEffect 、 glitchEffect 、 neonGlow 、 cinematicGrade 、および使いやすい applyLumetri グレード。
バッチとイントロスペクト: 1 回のラウンドトリップと 1 回の取り消し (batch) で多くの編集を実行し、その後レイヤーまたはコンプ状態 ( getLayerDetails 、 getProperty ) を読み取ります。
ブロックせずにレンダリング: レンダリングは aerender を生成し、progress イベントと renderComplete イベントをストリームするため、AE の応答性は維持されます。
自律パイプラインを実行します。ビデオを JSON セグメントとして記述すると、オーケストレーターが各セグメントを認識し ( applySpec )、スコープを指定してレンダリングし、視覚的にレビューし、間違っている場合は構造化された仕様デルタを適用し、合格するまで再レンダリングし、結果を FFmpeg と連結します。
要件
注意事項
Adobe After Effects
2024 (24.x)、2025 (25.x)、または 2026 (26.x)。ライセンスが取得され、アクティベートされました。正確なバージョンを記録します。
Node.js 18+
コントローラー、シミュレーター、ビルド スクリプトはノード (ESM) です。
FFmpeg
オーケストレーターのビジュアルレビュー担当者と最終連結に必要です。 ffmpeg は PATH 上に存在する必要があります。
Git
クローンを作成してプッシュします。
それらを確認してください:
ノード --バージョン番号 v18+ (v22 推奨)
ffmpeg -version # 最近のビルド
CEP バージョンは AE バージョンにマップされます。AE 2024 は CEP 11 (CSXS.11) を使用します。 AE 2025 および 2026 は CEP 12 ( CSXS.12 ) を使用します。パネル マニフェストは [24.0,99.9] と必要なランタイムが低いことをターゲットにしているため、それらすべてにロードされます。
3. クイックスタート、AE不要（シミュレータ）
After Effects を使用せずに、ヘッドレス シミュレーターを使用して、アーキテクチャ全体 (コントローラーからブリッジへのプロトコル、コマンドのディスパッチ、レンダリングの配管) を検証できます。これは、エリアとまったく同じ WebSocket プロトコルを話すノード プロセスです。

l パネルを開き、模擬 AE DOM に対してバンドルされた実際の JSX を実行します。
git クローン https://github.com/Arman-Luthra/aftr.git
CDアフター
npmインストール
# JSX バンドルを一度ビルドします
npm run build:jsx
# 完全なヘッドレス テスト スイートを実行します (ユニット + シミュレーター + JSX ディスパッチ)
npmテスト
# エンドツーエンド: コントローラー + シミュレーターの往復、含む。レンダリングイベント
npm 実行 e2e
または対話的に実行します。
# ターミナル 1: コントローラー (http://127.0.0.1:8787 の WS サーバー + REST + UI)
npm 実行コントローラー
# ターミナル 2: シミュレータ (AE パネルのふりをする)
npm 実行シミュレーション
# ここで http://127.0.0.1:8787 を開いてコマンドをクリックするか、次のようにします。
カール -X POST http://127.0.0.1:8787/command \
-H " Content-Type: application/json " \
-d ' {"コマンド":"createComp","params":{"名前":"メイン","幅":1920,"高さ":1080}} '
これは、 M0 、 M3 、および M4 (アーキテクチャ、ラウンドトリップ、語彙) がゼロ AE で検証される方法です。実際のパネルは、 M1 、 M2 、および M5 に対して変更されずにドロップされます。
4. 実際の After Effects を使用した完全なセットアップ
パネルは CEP 拡張です。最新の AE (2025+) では、PlayerDebugMode 経由で署名のないパネルを有効にするという古い手法は信頼できません。AE はこれを無視し、「署名の検証に失敗しました」というメッセージを表示して署名のない拡張機能を拒否します。信頼できるサポートされているパスは、パネルを自己署名してインストールすることです。このリポジトリは、次の 1 つのコマンドでこれを自動化します。
npm run デプロイ:パネル
deploy:panel は、JSX バンドルを構築し、自己署名開発証明書 ( dist/ にキャッシュ) を作成し、panel/ を .zxp に署名し、署名された拡張機能をユーザーごとの CEP 拡張フォルダーに解凍します。
これは Windows と macOS の両方で動作します (zxp-sign-cmd ツールには両方の署名者が同梱されています)。実行後、After Effects を完全に終了して再起動し、ウィンドウ／拡張機能／aftr を開きます。
署名付き拡張機能はスナップショットです。パネルソースを編集した後、npm rundeploy:panelを再実行して(そしてAEを再度開きます)、パネルソースを永続化するか、次のコマンドを使用します。

反復中にホットリロード開発ループ (セクション 11 を参照)。
# リポジトリのルートから
npmインストール
npm run デプロイ:パネル
# その後 AE を終了し、再起動し、[ウィンドウ] > [拡張機能] > [aftr]
署名されたパネルは次の場所にインストールされます。
%APPDATA%\Adobe\CEP\extensions\com.ae-bridge.panel
リモート DevTools のデバッグ モードを設定することもできますが、これは必須ではありません

[切り捨てられた]

## Original Extract

Puppeteer for After Effects. Use AE with Claude Code to make production-ready videos. - Arman-Luthra/aftr

GitHub - Arman-Luthra/aftr: Puppeteer for After Effects. Use AE with Claude Code to make production-ready videos. · GitHub
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
Arman-Luthra
/
aftr
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
10 Commits 10 Commits .github/ workflows .github/ workflows bin bin controller controller docs docs examples examples packaging/ pip packaging/ pip panel panel presets presets shared shared simulator simulator tools tools .dockerignore .dockerignore .gitignore .gitignore .npmignore .npmignore Dockerfile Dockerfile LICENSE LICENSE README.md README.md config.json config.json package-lock.json package-lock.json package.json package.json View all files Repository files navigation
Use After Effects with Claude Code to make production-ready videos.
A Node controller sends JSON commands over a WebSocket to a CEP panel inside AE, which runs them as ExtendScript and returns JSON. On top sits an autonomous, spec-driven pipeline that builds clips and self-corrects them by rendering, reviewing, and revising. Everything is driven programmatically over the socket, with no manual After Effects work.
A title sequence built and rendered entirely through aftr. No manual After Effects work.
Video made for tilion / fortress .
Watch the full demo
Wire aftr into Claude Code in two commands. Start the controller (it serves the MCP endpoint), then point Claude Code at it.
npx aftr-studio controller # no install, needs Node 18+
pip install aftr-studio && aftr controller # pip
docker run -p 8787:8787 ghcr.io/arman-luthra/aftr # docker
claude mcp add --transport http aftr http://127.0.0.1:8787/mcp
Open After Effects and the panel (Window > Extensions > aftr; deploy it once from a clone with npm run deploy:panel ). Then ask Claude Code in plain language:
Create a 1080p 5s comp, add a title "LAUNCH" with fire behind it, animate blurFade, then render to mp4.
Claude Code sees every command as a tool: it calls ae_status first, lists the set with ae_list_commands , and runs anything through ae_command . Prefer a stdio server? Use claude mcp add aftr -- npx -y aftr-studio mcp (keep the controller running too).
Method
Command
npm (global CLI)
npm install -g aftr-studio , then aftr controller
npx (no install)
npx aftr-studio controller
pip (Python launcher)
pip install aftr-studio , then aftr controller (needs Node 18+)
Docker
docker run --rm -p 8787:8787 ghcr.io/arman-luthra/aftr
From source
clone this repo (below)
The npm, npx, and Docker paths run the controller, MCP server ( aftr mcp ), and simulator ( aftr sim ). Deploying the CEP panel into After Effects is done from a clone ( npm run deploy:panel ), since it self-signs and installs the extension. The pip package is a thin launcher that forwards to the npm CLI through npx .
git clone https://github.com/Arman-Luthra/aftr.git
cd aftr
npm install
# 1) prove it's healthy with NO After Effects (headless simulator + tests)
npm run build:jsx && npm test
# 2) install the panel into AE (self-signs + deploys on Windows and macOS)
npm run deploy:panel # then quit and relaunch AE, Window > Extensions > aftr
# 3) start the bridge + open the UI
npm run controller # http://127.0.0.1:8787
Then drive it from the UI, REST, WebSocket, or as a library:
node examples/flaming-title.mjs " ON FIRE "
Requires After Effects 2024 to 2026, Node 18+, and ffmpeg on your PATH . Full Windows and macOS setup is in section 4 .
Capability
What it does
~100 commands
Comps, layers, keyframes with easing, expressions, effects (by matchName or display name, including deeply nested params and third-party plugins), masks, layer styles, text animators, render.
Discovery
listFonts , listInstalledEffects (returns matchNames), findEffectMatchName , listPlugins , getEnvironment . Find what's installed instead of guessing.
One-call VFX
fireEffect , smokeEffect , glitchEffect , neonGlow , cinematicGrade , and a friendly applyLumetri grade.
Text animation
Four Animate-panel presets ( wordReveal , charScale , bunchRotate , blurFade ) plus a full range-selector builder.
Autonomous pipeline
Declarative segment specs realize, render, get a visual review, take a structured delta, re-render, and concat. It self-corrects.
MCP server
Drive AE from Claude Desktop, Claude Code, or any MCP client. Every command is a tool ( npm run mcp ).
Fast and batched
batch runs N edits in one round-trip and one undo (sub-200 ms for 20 ops). Non-blocking aerender streams progress events, with a parallel render engine.
Testable without AE
A headless simulator runs the real JSX against a mock DOM. 96 tests plus e2e, CI on Node 18, 20, and 22.
Cross-platform
Windows and macOS, one-command sign and deploy.
How it works
┌──────────────────────────────────────────────┐
│ Controller (Node + ws + Express + web UI) │ <- you / the UI / an agent drive this
│ • hosts a WebSocket server (panel + agents) │
│ • REST surface + interactive UI │
│ • the HLD orchestrator (agentic loop) │
└───────────────┬────────────────────────────────┘
│ WebSocket (JSON request/response/event envelopes)
┌───────────────▼────────────────────────────────┐
│ CEP panel (inside After Effects) │ ┐
│ • WebSocket CLIENT, dials the controller │ │ build once,
│ • routes commands to the JSX dispatch layer │ │ identical on Win & Mac
│ • spawns aerender for non-blocking renders │ │
└───────────────┬────────────────────────────────┘ │
│ evalScript (string in / JSON string out)
┌───────────────▼────────────────────────────────┐ │
│ JSX command layer (host.jsx + commands/*.jsx) │ │
│ • dispatch(command, paramsJson) -> JSON │ │
│ • ~90 commands: layers, keyframes, effects, │ │
│ text animators, masks, VFX presets, render │ │
└───────────────┬────────────────────────────────┘ ┘
│ AE scripting DOM (app.project, comps, layers)
┌───────────────▼────────────────────────────────┐
│ After Effects │
└──────────────────────────────────────────────────┘
Table of contents
Quick start, no AE needed (the simulator)
Full setup with real After Effects
Windows
The orchestrator (autonomous pipeline)
Dev workflow (hot reload + signing)
Control AE from code: create comps and layers (solids, text, shapes, nulls, cameras, lights, adjustment, footage), set any transform or effect property, add keyframes with easing, write expressions, and parent, trim, move, or duplicate layers, all over a socket.
Add any effect by matchName and drive any of its parameters, even deeply nested ones, animated via keyframes or expressions.
Animate text with four ready-made Animate-panel styles ( wordReveal , charScale , bunchRotate , blurFade ) plus a full text-animator builder (Based On, Shape, Ease High to Low, keyframed Offset and Start).
Reach for one-call VFX presets: fireEffect (realistic fire), smokeEffect , glitchEffect , neonGlow , cinematicGrade , and a friendly applyLumetri grade.
Batch and introspect: run many edits in one round-trip and one undo ( batch ), then read back layer or comp state ( getLayerDetails , getProperty ).
Render without blocking: render spawns aerender and streams progress and renderComplete events, so AE stays responsive.
Run the autonomous pipeline: describe a video as JSON segments, and the orchestrator realizes each one ( applySpec ), renders it scoped, reviews it visually, applies a structured spec delta if it's wrong, re-renders until it passes, then concatenates the result with FFmpeg.
Requirement
Notes
Adobe After Effects
2024 (24.x), 2025 (25.x), or 2026 (26.x). Licensed and activated. Record your exact version.
Node.js 18+
The controller, simulator, and build scripts are Node (ESM).
FFmpeg
Needed by the orchestrator's visual reviewer and final concat. ffmpeg must be on your PATH .
Git
To clone and push.
Check them:
node --version # v18+ (v22 recommended)
ffmpeg -version # any recent build
CEP version maps to AE version: AE 2024 uses CEP 11 ( CSXS.11 ); AE 2025 and 2026 use CEP 12 ( CSXS.12 ). The panel manifest targets [24.0,99.9] and a low required runtime, so it loads on all of them.
3. Quick start, no AE needed (the simulator)
You can validate the entire architecture (controller-to-bridge protocol, command dispatch, render plumbing) without After Effects, using the headless simulator. It is a Node process that speaks the exact same WebSocket protocol as the real panel and runs the real bundled JSX against a mock AE DOM.
git clone https://github.com/Arman-Luthra/aftr.git
cd aftr
npm install
# build the JSX bundle once
npm run build:jsx
# run the full headless test suite (unit + simulator + JSX dispatch)
npm test
# end-to-end: controller + simulator round-trip, incl. render events
npm run e2e
Or run it interactively:
# terminal 1: controller (WS server + REST + UI on http://127.0.0.1:8787)
npm run controller
# terminal 2: the simulator (pretends to be the AE panel)
npm run sim
# now open http://127.0.0.1:8787 and click commands, or:
curl -X POST http://127.0.0.1:8787/command \
-H " Content-Type: application/json " \
-d ' {"command":"createComp","params":{"name":"Main","width":1920,"height":1080}} '
This is how M0 , M3 , and M4 (architecture, round-trip, vocabulary) are validated with zero AE. The real panel drops in unchanged for M1 , M2 , and M5 .
4. Full setup with real After Effects
The panel is a CEP extension. On modern AE (2025+), the old trick of enabling unsigned panels via PlayerDebugMode is unreliable: AE ignores it and rejects unsigned extensions with "Signature verification failed." The reliable, supported path is to self-sign and install the panel, which this repo automates with one command:
npm run deploy:panel
deploy:panel builds the JSX bundle, creates a self-signed dev certificate (cached in dist/ ), signs panel/ into a .zxp , and unpacks the signed extension into your per-user CEP extensions folder.
It works on both Windows and macOS (the zxp-sign-cmd tool ships the signer for both). After it runs, fully quit and relaunch After Effects, then open Window > Extensions > aftr.
A signed extension is a snapshot. After you edit panel source, re-run npm run deploy:panel (and reopen AE) to make it permanent, or use the hot-reload dev loop (see section 11 ) while iterating.
# from the repo root
npm install
npm run deploy:panel
# then quit AE, relaunch, Window > Extensions > aftr
The signed panel installs to:
%APPDATA%\Adobe\CEP\extensions\com.ae-bridge.panel
You may also set debug mode for remote DevTools, but it is not req

[truncated]
