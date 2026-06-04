---
source: "https://github.com/cjmielke/NeuroAnswer"
hn_url: "https://news.ycombinator.com/item?id=48400247"
title: "Show HN: NeuroAnswer – MCP server lets Claude navigate petabyte-scale brain maps"
article_title: "GitHub - cjmielke/NeuroAnswer: An AI chat assistant for NeuroGlancer! 👀 · GitHub"
author: "subcosmos"
captured_at: "2026-06-04T16:11:46Z"
capture_tool: "hn-digest"
hn_id: 48400247
score: 3
comments: 0
posted_at: "2026-06-04T15:37:53Z"
tags:
  - hacker-news
  - translated
---

# Show HN: NeuroAnswer – MCP server lets Claude navigate petabyte-scale brain maps

- HN: [48400247](https://news.ycombinator.com/item?id=48400247)
- Source: [github.com](https://github.com/cjmielke/NeuroAnswer)
- Score: 3
- Comments: 0
- Posted: 2026-06-04T15:37:53Z

## Translation

タイトル: Show HN: NeuroAnswer – MCP サーバーでクロードがペタバイト規模の脳マップをナビゲートできるようにする
記事タイトル: GitHub - cjmielke/NeuroAnswer: NeuroGlancer の AI チャット アシスタント! 👀 · GitHub
説明: NeuroGlancer の AI チャット アシスタント! 👀。 GitHub でアカウントを作成して、cjmielke/NeuroAnswer の開発に貢献してください。

記事本文:
GitHub - cjmielke/NeuroAnswer: NeuroGlancer の AI チャット アシスタント! 👀 · GitHub
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
チミエルケ
/
ニューロアンサー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
主なブランチのタグが取得されました

o ファイル コード [その他のアクション] メニューを開く フォルダーとファイル
58 コミット 58 コミット .github/ workflows .github/ workflows chrome_extension chrome_extension img img mcp_src mcp_src .env.example .env.example .gitignore .gitignore Dockerfile Dockerfile README.md README.md api_server.py api_server.py docker-compose.yml docker-compose.yml mcp_main.py mcp_main.py要件.txt要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ニューロアンサー.mp4
200,000 個のニューロン、5 億個のシナプスが平易な英語でクエリ可能、3D でライブ レンダリングされます。
NeuroAnswer は、MICrON の Minnie65 コネクトーム データセットを探索するための AI コパイロットであり、結果は Neuroglancer で直接レンダリングされます。
v2 リファクタリング — クロードはあなたが見ているものを見る
v2 アーキテクチャでは根本的な変更が行われています。Neuroglancer ビューアは、拡張機能が状態をミラーリングする必要があった別個のブラウザ セッションではなく、MCP サーバーによって直接制御されるようになりました。 V1 では、Chrome 拡張機能と JavaScript ハックを使用して状態を変更しました。 V2 は独自の NeuroGlancer サーバーをローカルでホストし、既存の Python/WebSockets API を使用してサーバーを直接制御します。これにより、純粋な MCP NeuroGlancer ツールの開発が可能になります。
これにより、以前は不可能だった機能も解放されます。
クロードは、アクティブなニューログランサー ビューポートのスクリーンショットを撮り、それを視覚的に解釈できるようになりました。これは、研究者が見ているのと同じビューを、同じ瞬間に視覚的に解釈したものです。細胞の形態を識別し、目に見える構造を数え、前景と背景にあるものを認識し、次に何をクエリするかを決定する前に何が見えるかを推論することができます。 CAVE ツールと組み合わせると、観察 → クエリ → レンダリング → 再度観察というループが終了します。
ニューロンとアストロサイトである可能性のあるものを識別するクロード — 3D メッシュ ビューのライブ スクリーンショットより。
v1
v2
ニューグランサーコントロール
拡張膿

URL ハッシュによるヘッド状態
MCP サーバーは Viewer オブジェクトを直接所有します
カメラ/レイヤー
拡張機能がブラウザの URL を書き換えた
MCP ツールはサーバー側で viewer.txn() を呼び出します
スクリーンショット
不可能
viewer.screenshot() はレンダリングされたイメージを返します。クロードはそれを見る
認証
Chrome Cookie / Google OAuth ポップアップ
MiddleAuthProvider はサーバー側に CAVE トークンを挿入します
チャットUI
Chrome 拡張機能のサイドバー
Chrome 拡張機能のサイドバー (維持); Claude Desktop はテストされましたが、ツール イメージはレンダリングされません。ボーナス : 生成された Python コードは、Prism による構文強調表示とともに UI に表示されるようになりました。
輸送
SSE を介した MCP
SSE を介した MCP (変更なし)。 Claude デスクトップで利用可能な stdio
ニューグランサーのURL
パブリック Neuroglancer サーバーに依存 / Chrome 拡張機能で適切なタブを検出する必要があった
MCP サーバーは独自のローカル NeuroGlancer サーバーをホストします
削除されました
nglui シーンの URL 構築 (レイヤーはビューア上で直接設定されるようになりました)
ng_state をポーリングしてプッシュした拡張コンテンツ スクリプト
FastAPI エンドポイント /scenes および /new_scene (reset_neuroglancer_view MCP ツールに置き換えられました)
リレー応答によるレイヤー/位置/シーンの伝播
NeuroAnswer は、MCP (モデル コンテキスト プロトコル) サーバーを介してクロードを CAVE コネクトミクス データベースに接続します。わかりやすい英語で質問すると、「最も近い体細胞を見せて」または「この L5 錐体細胞の下流のシナプスをトレースして」と質問すると、システムがデータを取得し、3D シーンを構築して、ニューログランサー ビューにレンダリングします。
空間クエリ: 現在の Neuroglancer ビューポート内で細胞タイプごとにニューロンを検索します。
シナプス トレース: 任意のニューロンの下流シナプス ターゲットを取得し、シナプス前/シナプス後マーカーをレンダリングします。
コード実行サンドボックス : クロードが生成された Python (pandas、numpy、matplotlib) ワークフローを MICrONs データセットに対して実行し、ブラウザーで科学的な数値をレンダリングできるようにします。
MCP ツールを作成しました。

Claude は、プリロードされたデータフレームに対して Python を生成して実行し、一般的な科学クエリを実行できるようになります。科学的な数値がチャット ウィンドウに直接レンダリングされます。
この視野内で最も近い 10 個のニューロンを表示します。次に、それらのニューロンごとに API をクエリして、下流のニューロンを見つけます。プールされたシナプスのセットについて、下流ニューロンのタイプ別のシナプス サイズの分布を示す水平バイオリン プロットをレンダリングします。
次の分析を作成しました
ダウンストリーム シナプス プール
10 個のニューロンすべてから合計 3,049 個のシナプスがフェッチされました。細胞型テーブルに結合し、10 個以上のシナプスを持つ m 型をフィルター処理した後、17 個の異なる下流 m 型が視覚化されました。
L5ET ニューロンは、単一の最も一般的な下流標的 (n=647) であり、この領域の深層興奮性細胞間の密な反復接続を反映しています。
PTC および DTC ニューロンも大量の入力 (n=300 および n=245) を受け取り、これらの投影クラスの細胞がこの局所回路の主要なハブであることを示唆しています。
サイズ分布はすべての m タイプにわたって右に偏っており、ほとんどのシナプスは小さいですが、大きなシナプスの長い尾部があり、これは皮質における既知の対数正規シナプス サイズ分布と一致しています。
各バイオリンの白いバーは IQR (Q1 ～ Q3) を表し、白いチェックマークは中央値を示します。下から上への順序は、シナプス サイズの中央値の増加を反映しています。
464 個のシナプス (約 15%) が細胞型テーブルにないニューロンにヒットし、「不明」として除外されました。
コンパートメントターゲティング: シナプス後コンパートメント (体体、シャフト、スパイン) によってシナプスを識別します。
集団検索: 興奮性ニューロン集団と抑制性ニューロン集団を形態学的タイプ別に検索します。
より豊富なアノテーションのサポート: シナプス前部位とシナプス後部位をサイズ/重量ラベルで接続する線アノテーション
データセットの一般化: FlyWire、H0 へのサポートの拡張

1 (人間の大脳皮質)、および異なるデータベース スキーマを持つ他のコネクトミクス データセット
アーキテクチャの簡素化: エージェント ループ全体を Chrome 拡張機能に移動して、FastAPI リレーを排除します
コミュニティ主導のツール設計 : ツールの語彙は研究者が実際に必要とするものに基づいて形成されるべきです - 有用なクエリやワークフローに関するフィードバックは大歓迎です
Chrome 拡張機能 (サイドバー UI)
↕
FastAPIゲートウェイ（中継サーバー）
↕
MCP サーバー (FastMCP) <-> NeuroGlancer (Python/websockets API を備えたローカル サーバー)
↕
CAVE / MICrONs Minnie65 クラウド データベース
MCP サーバーは、Claude が標準 MCP プロトコルを通じて呼び出すドメイン固有のツール (空間検索、シナプス クエリ、カメラ制御、シーン構築) を公開します。また、Neuroglancer サーバーを直接ホストしているため、ブラウザ側の調整を行わずに、ツール呼び出しによってビューアの状態がリアルタイムで変更されます。 FastAPI ゲートウェイは、Chrome 拡張機能の HTTP リクエストを MCP サーバーにブリッジし、ツールの結果 (テキスト、画像、構文強調表示されたコード) をサイドバーにストリーミングして返します。
Neuron メタデータは、空間クエリを高速に保つために Parquet ファイルとしてローカルにキャッシュされます。シナプス データはオンデマンドで CAVE からライブで取得されます。少し遅いですが、取り組んでいます。
Python · FastMCP · CAVEclient · NeuroGlancer · FastAPI · Chrome Extensions API
構造脳グラフをクエリするには、Anthropic API キーとともに CAVE API トークンが必要です。
環境テンプレートをコピーします。
cp .env.example .env
DAF API 認証ポータルに移動します。この追加ドキュメントが役に立つかもしれません。
Google アカウントでログインします。
生成されたトークン文字列をコピーします。
新しい .env ファイルを開き、資格情報を貼り付けます。
ANTHROPIC_API_KEY = your_claude_key_here
CAVE_TOKEN = your_copied_cave_token_here
LANGFUSE_PUBLIC_KEY = オプションのテレメトリーキー
CACHE_DIR はオプションです - サーバー d

デフォルトは、OS に適したユーザー キャッシュ ディレクトリです。 Docker 経由で実行する場合、自動的に /app/data (マウントされたボリューム) に設定されるため、オーバーライドは必要ありません。
開発中にチャットのログを記録するために「LangFuse」を使用していることに注意してください。同じことをしたい場合は、次のようにすることができます
無料のアカウントにサインアップし、API キーを作成し、それらを .env ファイルに配置します。
認証情報を設定すると、docker を使用してサーバーを起動できます。
ドッカーの構成
フロントエンドのセットアップ
理論的には、NeuroGlancer なしでも MCP サーバーを使用できますが、面白くありません。科学は楽しいはずです！
楽しい部分を利用するには、Google Chrome 拡張機能をインストールする必要があります。 Chrome 拡張機能ストアにはありませんが、ローカルにインストールできます。
Google Chrome では、URL バーに chrome://extensions と入力し、Enter キーを押します。
トグルボタンで開発者モードを有効にする
[Load unpacked] をクリックし、この複製されたリポジトリの chrome_extension ディレクトリに移動します。
拡張機能を開くときに、サイドバーの上部にある [Open Neuroglancer ↗] リンクをクリックして、ローカルでホストされているビューアを開きます。
このプロジェクトは現在鋭意開発中です。コネクトミクスまたは神経科学に携わっており、Neuroglancer の自然言語インターフェイスで最も役立つクエリやワークフローについて考えがある場合は、問題を作成するか、連絡してください。目標は、現場で実際に使用したいものを構築することです。
このツールを使用することで、生物学的ニューロンが言語を作成し、その言語が大規模なインシリコ ニューロンのセットに通知され、マウスの生物学的ニューロンに関する情報をデータベースに問い合わせる言語を作成するという事実を少し時間をとって考えてください。
「私たちは宇宙が自分自身を知るための手段である」 - カール・セーガン
NeuroGlancer の AI チャット アシスタント! 👀
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
エラーがありました

r 読み込み中。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

An AI chat assistant for NeuroGlancer! 👀. Contribute to cjmielke/NeuroAnswer development by creating an account on GitHub.

GitHub - cjmielke/NeuroAnswer: An AI chat assistant for NeuroGlancer! 👀 · GitHub
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
cjmielke
/
NeuroAnswer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
58 Commits 58 Commits .github/ workflows .github/ workflows chrome_extension chrome_extension img img mcp_src mcp_src .env.example .env.example .gitignore .gitignore Dockerfile Dockerfile README.md README.md api_server.py api_server.py docker-compose.yml docker-compose.yml mcp_main.py mcp_main.py requirements.txt requirements.txt View all files Repository files navigation
neuroanswer.mp4
200,000 neurons, 500 million synapses, queryable in plain English — rendered live in 3D!
NeuroAnswer is an AI copilot for exploring the MICrONs Minnie65 connectome dataset, with results rendered directly in Neuroglancer .
v2 Refactor — Claude Sees What You See
The v2 architecture makes a fundamental change: the Neuroglancer viewer is now controlled directly by the MCP server rather than being a separate browser session that the extension had to mirror state into. V1 used the chrome extension and javascript hacks to mutate the state. V2 hosts its own neuroglancer server locally and uses its existing python/websockets API to directly control it. This allows the development of pure MCP NeuroGlancer tools.
This also unlocks a capability that wasn't possible before!
Claude can now take a screenshot of the active Neuroglancer viewport and interpret it visually — the same view the researcher is looking at, at the same moment. It can identify cell morphology, count visible structures, notice what's in the foreground vs background, and reason about what it sees before deciding what to query next. Combined with the CAVE tools, this closes a loop: observe → query → render → observe again.
Claude identifying neurons, and what might be an astrocyte — from a live screenshot of the 3D mesh view.
v1
v2
Neuroglancer control
Extension pushed state via URL hash
MCP server owns the Viewer object directly
Camera / layers
Extension rewrote the browser URL
MCP tools call viewer.txn() server-side
Screenshots
Not possible
viewer.screenshot() returns rendered image; Claude sees it
Auth
Chrome cookie / Google OAuth popup
MiddleAuthProvider injects CAVE token server-side
Chat UI
Chrome extension sidebar
Chrome extension sidebar (kept); Claude Desktop tested but doesn't render tool images. Bonus : Generated python code is now shown in the UI with syntax highlighting powered by prism
Transport
MCP over SSE
MCP over SSE (unchanged); stdio available for Claude Desktop
Neuroglancer URL
Relied on public Neuroglancer servers / chrome extension had to detect the right tab
MCP server hosts its own local NeuroGlancer server
Removed
nglui scene URL construction (layers are now set directly on the viewer)
Extension content script that polled and pushed ng_state
FastAPI endpoints /scenes and /new_scene (replaced by reset_neuroglancer_view MCP tool)
Layer / position / scene propagation through the relay response
NeuroAnswer connects Claude to the CAVE connectomics database via an MCP (Model Context Protocol) server. Ask a question in plain English — "show me the closest soma" or "trace the downstream synapses of this L5 pyramidal cell" — and the system fetches the data, builds the 3D scene, and renders it in your Neuroglancer view.
Spatial queries: find neurons by cell type within the current Neuroglancer viewport
Synapse tracing: retrieve downstream synaptic targets for any neuron and render pre/post-synaptic markers
Code execution sandbox : allowing Claude to run generated Python (pandas, numpy, matplotlib) workflows against the MICrONs dataset and render scientific figures in the browser
Ive created an MCP tool that allows Claude to generate and execute python against pre-loaded dataframes, allowing generic scientific queries. Scientific figures render directly in the chat window!
Show me the 10 closest neurons in this field of view. Then for each of those neurons, query the API to find their downstream neurons. For that pooled set of synapses, render a horizontal violin plot showing the distributions of synapse size by the mtype of the downstream neuron
Produced the following analysis
Downstream Synapse Pool
3,049 total synapses were fetched across all 10 neurons. After joining to the cell type table and filtering for m-types with ≥10 synapses, 17 distinct downstream m-types were visualized.
L5ET neurons are the single most common downstream target (n=647), reflecting the dense recurrent connectivity among deep-layer excitatory cells in this region.
PTC and DTC neurons also receive heavy input (n=300 and n=245), suggesting these projection-class cells are major hubs in this local circuit.
The size distributions are right-skewed across all m-types, with most synapses being small but a long tail of large synapses — consistent with the known lognormal synapse size distribution in cortex.
White bars on each violin represent the IQR (Q1–Q3), and the white tick marks the median. The ordering from bottom to top reflects increasing median synapse size.
464 synapses (~15%) hit neurons not in the cell-type table and were excluded as "Unknown".
Compartment targeting: identify synapses by postsynaptic compartment (soma, shaft, spine)
Population search: look up excitatory and inhibitory neuron populations by morphological type
Richer annotation support : line annotations connecting pre- and post-synaptic sites with size/weight labels
Dataset generalization : extending support to FlyWire, H01 (human cortex), and other connectomics datasets with different database schemas
Simplified architecture : moving the full agent loop into the Chrome extension to eliminate the FastAPI relay
Community-driven tool design : the tool vocabulary should be shaped by what researchers actually need — feedback on useful queries and workflows is very welcome
Chrome Extension (sidebar UI)
↕
FastAPI Gateway (relay server)
↕
MCP Server (FastMCP) <-> NeuroGlancer (local server with Python/websockets API)
↕
CAVE / MICrONs Minnie65 cloud database
The MCP server exposes domain-specific tools — spatial search, synapse queries, camera control, and scene construction — that Claude calls through the standard MCP protocol. It also hosts the Neuroglancer server directly, so tool calls mutate the viewer state in real time without any browser-side coordination. A FastAPI gateway bridges the Chrome extension's HTTP requests to the MCP server and streams tool results (text, images, syntax-highlighted code) back to the sidebar.
Neuron metadata is cached locally as Parquet files to keep spatial queries fast. Synapse data is fetched live from CAVE on demand - a bit slower, but I'm working on it!
Python · FastMCP · CAVEclient · NeuroGlancer · FastAPI · Chrome Extensions API
To query the structural brain graph, you need a CAVE API token, along with your Anthropic API key.
Copy the environment template:
cp .env.example .env
Go to the DAF API Auth Portal . You may find this additional documentation useful.
Log in with your Google account.
Copy the generated token string.
Open your new .env file and paste your credentials:
ANTHROPIC_API_KEY = your_claude_key_here
CAVE_TOKEN = your_copied_cave_token_here
LANGFUSE_PUBLIC_KEY = optional_telemetry_key
CACHE_DIR is optional — the server defaults to the OS-appropriate user cache directory. When running via Docker, it is automatically set to /app/data (the mounted volume) so no override is needed.
Note, I'm using "LangFuse" to log my chats during development. If you want to do the same, you can
sign up for a free account, create API keys, and also place them in the .env file.
With the credentials set up, the servers can be stood up using docker
docker compose up
Frontend setup
You can theoretically use the MCP server without NeuroGlancer, but it won't be fun. Science should be fun!
You'll need to install the Google Chrome extension for the fun part. Its not in the chrome extension store, but you can install it locally.
In Google Chrome, type chrome://extensions in the URL bar and hit enter
Enable developer mode with the toggle button
Click Load unpacked and then navigate to the chrome_extension directory in this cloned repository
When you open the extension, click the Open Neuroglancer ↗ link at the top of the sidebar to open the locally-hosted viewer
This project is under active development. If you work in connectomics or neuroscience and have thoughts on what queries or workflows would be most useful in a natural language interface for Neuroglancer, please open an issue or reach out. The goal is to build something the field actually wants to use.
Take a minute to reflect on the fact that, by using this tool, your biological neurons are creating language which informs a large set of in-silico neurons to create language that queries a database for information on a mouses biological neurons.
"We are a way for the universe to know itself" - Carl Sagan
An AI chat assistant for NeuroGlancer! 👀
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
