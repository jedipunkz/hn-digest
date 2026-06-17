---
source: "https://github.com/Adam-CAD/CADAM"
hn_url: "https://news.ycombinator.com/item?id=48572553"
title: "Launch HN: Adam (YC W25) – Open-Source AI CAD"
article_title: "GitHub - Adam-CAD/CADAM: CADAM is the open source text-to-CAD web application · GitHub"
author: "zachdive"
captured_at: "2026-06-17T16:19:40Z"
capture_tool: "hn-digest"
hn_id: 48572553
score: 5
comments: 0
posted_at: "2026-06-17T16:14:04Z"
tags:
  - hacker-news
  - translated
---

# Launch HN: Adam (YC W25) – Open-Source AI CAD

- HN: [48572553](https://news.ycombinator.com/item?id=48572553)
- Source: [github.com](https://github.com/Adam-CAD/CADAM)
- Score: 5
- Comments: 0
- Posted: 2026-06-17T16:14:04Z

## Translation

タイトル: HN: Adam (YC W25) – オープンソース AI CAD の開始
記事のタイトル: GitHub - Adam-CAD/CADAM: CADAM はオープンソースのテキストから CAD への Web アプリケーションです · GitHub
説明: CADAM は、オープンソースのテキストから CAD への Web アプリケーションです - Adam-CAD/CADAM
HN テキスト: こんにちは、HN!アダム（ https://adam.new/ ）のザックです。私たちは機械 CAD ソフトウェア用の AI エージェントを構築しています。私たちは 2 つの基本的な信念に基づいて会社を設立しました。 - AI は、今日のソフトウェアと同様に、機械設計を作成するための主要な媒体となる。 - CAD 生成の最適なパラダイムは、CAD をコードとして生成することです (テキスト -> コード -> CAD)。私たちは、オープンソースの Text to CAD プラットフォームである CADAM を構築しています。これは、認証、データベース、ファイル ストレージ用の Supabase バックエンドを備えた React アプリ (TanStack Start) です。 AI TinkerCAD のようなものだと考えてください。デモ: https://www.youtube.com/watch?v=iESOr7EGWqk
試してみてください: https://adam.new/cadam/ 機能: - テキスト プロンプトと画像参照の両方をサポートし、自然言語からパラメトリック 3D モデルを生成します。 - 瞬時に寸法を調整するための対話型スライダーとして表示される自動抽出されたパラメーターを含む OpenSCAD コードを出力 - .STL または .SCAD (さらに OBJ、GLB/GLTF、FBX、および DXF) としてエクスポートテクスチャーメッシュ。 - 単純なパラメーターの調整はモデルを完全にバイパスします。スライダーを調整すると、SCAD ソース上で決定的な正規表現の更新が行われるため、LLM 呼び出しは必要ありません。 - Vercel AI SDK を介したモデルに依存しない: Anthropic (Claude)、Google (Gemini)、OpenRouter を介した OpenAI/その他。新しいモデルでは適応型思考が自動的に有効になります。驚いたことに、私たちの評価では Gemini 3.1 Pro が最上位モデルです。 - 完全に実行されます -

OpenSCAD を WebAssembly にコンパイルし (Web Worker 内で UI がブロックされないように)、React Three Fiber 経由で Three.js でレンダリングすることでブラウザーを実現します。 - BOSL、BOSL2、および MCAD ライブラリをサポートし、さらにモデル内のテキストのカスタム フォント サポート (Geist) をサポートします。 将来の改善点: - build123d と CadQuery の両方をサポートします。これにより、CSG プリミティブを超えて制約駆動モデリングに移行し、他の Code-as-CAD プリミティブとの直接比較が可能になります。 - 空間コンテキストの向上: 面/エッジ選択用の UI とビューポート画像統合により、LLM に空間理解を提供します。リポジトリを複製してローカルで実行できます。貢献は大歓迎です。

記事本文:
GitHub - Adam-CAD/CADAM: CADAM は、オープンソースのテキストから CAD への Web アプリケーションです · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
アダム-CAD
/
カダム
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
master ブランチ タグ 移動先

ファイル コード [その他のアクション] メニューを開く フォルダーとファイル
280 コミット 280 コミット .cursor/ rules .cursor/ rules .husky .husky ベンチマーク ベンチマーク パブリック パブリック スクリプト スクリプト 共有 共有 src src supabase supabase .env.local.template .env.local.template .gitignore .gitignore .npmrc .npmrc .prettierignore .prettierignore .prettierrc .prettierrc CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md コンポーネント.json コンポーネント.json eslint.config.js eslint.config.js package-lock.json package-lock.json package.json package.json postcss.config.js postcss.config.js tailwind.config.js tailwind.config.js tsconfig.app.json tsconfig.app.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vercel.json vercel.json vite.config.ts vite.config.ts すべて表示ファイル リポジトリ ファイルのナビゲーション
⛮ オープンソースの Text to CAD Web アプリ ⛮
👉 adam.new/cadam .ブラウザ上で CAD モデルを数秒で生成します。インストールは必要ありません。
🤖 AI を活用した生成 - 自然言語と画像を 3D モデルに変換します
🎛️ パラメトリック コントロール - インタラクティブなスライダーで寸法を瞬時に調整できます
📦 複数のエクスポート形式 - .STL、.SCAD、または .DXF ファイルとしてエクスポート
🌐 ブラウザベース - WebAssembly を使用してブラウザ内で完全に実行します
📚 ライブラリのサポート - BOSL、BOSL2、および MCAD ライブラリが含まれます
特徴
説明
自然言語入力
3D モデルをわかりやすい英語で説明してください
画像参照
モデル生成のガイドとなる画像をアップロードする
リアルタイムプレビュー
Three.js を使用してモデルが即座に更新されることを確認します。
パラメータの抽出
調整可能な寸法を自動的に識別します
スマートアップデート
AIの再生成を必要としない効率的なパラメータ変更
カスタムフォント
モデル内のテキストに対する組み込みの Geist フォントのサポート
📺 スクリーンショット
CADA のショーケース

M は、完全なマルチパート マシンからクリーンなパラメトリック パートに至るまで、単一の平易な言語記述から構築されます。以下の各モデルは、表示されているプロンプトから開始され、完全にパラメトリックな OpenSCAD として表示され、 .STL 、 .SCAD 、または .DXF としてエクスポートする準備ができています。ソースとそれぞれの簡単な説明は、 benchmarks/ にあります。周回プレビューは、 benchmarks/render.sh でレンダリングされます。
モデルプロンプト制御出力
V8エンジン
完全な V8 内燃エンジン: 90° V の 4 シリンダーの 2 つのバンク、リブ付きバルブ カバーを備えたシリンダー ヘッド、谷の吸気マニホールド、各バンクの下の排気ヘッダー、カウンターウェイト付きのクランクシャフト、ピストンとコネクティング ロッド、フロント プーリー、およびオイル パン。
22 ディム
8色
9気筒星型航空機エンジン
9 気筒星形航空機エンジンを設計します。中央の円形クランクケースの周りに星型に均等に配置された 9 つのフィン付きシリンダーがあり、各シリンダーには積み重ねられた冷却フィンとドーム型のシリンダー ヘッドがあり、前部には中央のプロペラ シャフト ハブがあります。
15 ディム
6色
ターボファンジェットエンジン
完全な高バイパス ターボファン: 内部が見えるフロント ファン、バイパス カウル、コンプレッサー/タービン ステージを備えた内部コア、出口ガイド ベーン、排気プラグ。
2 ディム
10色
アキシャルタービンブリスク
ジェット エンジンのコンプレッサー段のような軸流タービン ブリスク (ブレード付きディスク) をモデル化します。シャフト ボアを備えた中央ハブと、リムの周囲にある約 28 枚の薄い翼型ブレードの単一リングで、各ブレードは実際のタービン ブレードと同様に、根元から先端まで高さに沿って明らかにねじれています。
14 ディム
1色
パラメトリックの基礎
モデルプロンプト制御出力
ねじり六角花瓶
ねじれた六角形の花瓶をデザインします。高さ約 150 mm の中空の殻で、底部 70 mm から口部 50 mm まで先細りになっており、六角形の断面は下から上に 120 度ねじれています。

壁は 2 mm、底は閉じています。
6 ディム
1色
ローレット加工のコントロールノブ
直径 40 mm、高さ 22 mm のローレット付きコントロール ノブを作成します。ダイヤモンド ローレット付きグリップ、上部に盛り上がったポインタ マーク、6 mm の D 形シャフト ボア、側面に M3 止めネジ穴があります。
15 ディム
2色
六角ボルトとナット - リアルスレッド
実ネジ付きシャフトと標準六角頭、およびそれに対応する六角ナットを並べて配置した、長さ 45 mm の M12 六角ボルトを作成します。
3 ディム
2色
ハニカム軽量ブラケット
厚さ 5 mm の 80x80 mm フランジを備え、両面に六角形のハニカム カットアウト パターンで軽量化され、4 つの M5 取り付け穴、およびフィレット加工された内側コーナーを備えた 90 度の角度の取り付けブラケットを設計します。
13 ディム
1色
NACA 2412 テーパーウィング
本物の NACA 2412 翼型を使用して、テーパー状の航空機の翼セクションをモデル化します。200 mm のスパンにわたって 120 mm の根元コードが 80 mm の先端までテーパーになっており、スパン方向のスパー チューブが 2 つといくつかの軽量穴が付いています。
9 ディム
1色
ネジ付きジャーとネジ式蓋
首に雄ネジが付いている小さな保存瓶と、雌ネジが付いている対応するネジ式の蓋を作成します。ジャー本体直径 60 mm、高さ 70 mm、壁 2.5 mm。蓋を外して瓶の横に置いてあるところを見せてください。
9 ディム
2色
直角ベベルギヤドライブ
直角ベベル ギヤ ドライブを構築します。垂直軸上の 24 歯のベベル ギヤと、水平軸上の 16 歯のベベル ピニオンが 90 度で噛み合い、それぞれが短いスタブ シャフト上にあります。
9 ディム
3色
遠心ポンプの羽根車
遠心ポンプのインペラを設計します。中央に 12 mm の穴があり、盛り上がったハブを備えた直径 90 mm のバックプレートと、ハブからリムに向かって後方に湾曲した 7 枚のブレードがあり、各ブレードはその経路に沿って滑らかに湾曲しています。
10 ディム
1色
ヘリンボーン遊星歯車ステージ
ヘリンボーン遊星歯車のモデル化
[切り捨てられた]
# リポジトリのクローンを作成します
g

https://github.com/Adam-CAD/CADAM.git のクローンを作成します
cd CADAM
# 依存関係をインストールする
npmインストール
# Supabaseを開始する
npxスーパーベースの開始
npx supabase 関数は --no-verify-jwt を提供します
# 開発サーバーを起動します
npm 実行開発
📋 前提条件
Node.js ^20.19.0 または >=22.12.0、npm 10+
ngrok (ローカル Webhook 開発用)
🔧 環境変数の設定
.env.local.template を .env.local にコピーする
.env.local 内の必要なキーをすべて更新します。
VITE_SUPABASE_ANON_KEY="<Anon キーのテスト>"
VITE_SUPABASE_URL='http://127.0.0.1:54321'
以下を含むサーバー側キーを .env.local に追加します。
ANTHROPIC_API_KEY="<テスト用 Anthropic API キー>"
OPENROUTER_API_KEY="<OpenRouter API キーのテスト>"
OPENAI_API_KEY="<OpenAI API キーをテスト>"
GOOGLE_API_KEY="<Google API キーをテスト>"
FAL_KEY="<テスト FAL API キー>"
SUPABASE_SERVICE_ROLE_KEY="<テスト サービス ロール キー>"
BILLING_SERVICE_URL="<テスト請求サービス URL>"
BILLING_SERVICE_KEY="<テスト請求サービス キー>"
環境 = "ローカル"
ADAM_URL="<Adam URL または dev URL>" # チェックアウトおよびポータルのリダイレクト ターゲット
WEBHOOK_BASE_URL="<パブリック TanStack アプリ URL>" # /cadam/api コールバックのアプリ URL
NGROK_URL="<NGROK URL>" # プロバイダーが読み取り可能な署名付き URL 用のオプションのローカル Supabase Storage トンネル
🌐 ローカル開発用の ngrok のセットアップ
CADAM は、プロバイダー コールバックとローカルの署名付きストレージ URL にパブリック URL を使用します。
まだインストールしていない場合は、ngrok をインストールします。
npm install -g ngrok
# または
醸造インストールngrok
TanStack Start 開発サーバーを指す ngrok トンネルを開始します。
生成された ngrok URL (例: https://xxxx-xx-xx-xxx-xx.ngrok.io ) をコピーし、.env.local ファイルに追加します。
WEBHOOK_BASE_URL="https://xxxx-xx-xx-xxx-xx.ngrok.io"
プロバイダーがローカルの Supabase Storage 署名付き URL を取得する必要がある場合は、Supabase への 2 番目のトンネルを実行し、NGROK_URL をその URL に設定します。
同じファイル内で ENVIRONMENT="local" が設定されていることを確認してください。
n

pxスーパーベースの開始
npm 実行開発
🛠️ で構築
フロントエンド: React 19 + TypeScript + TanStack Start + Vite
3D レンダリング: Three.js + React Three Fiber
CAD エンジン: OpenSCAD WebAssembly
バックエンド: TanStack Start サーバー ルート + Supabase PostgreSQL/認証/ストレージ
スタイリング: Tailwind CSS + shadcn/ui
これを改善するための提案がある場合は、リポジトリをフォークしてプル リクエストを作成してください。問題を開くこともできます。
指示と行動規範については、CONTRIBUTING.md を参照してください。
このアプリは、次の作業がなければ実現できません。
このディストリビューションは、GNU General Public License v3.0 (GPLv3) に基づいてライセンスされています。 「ライセンス」を参照してください。
このプロジェクトの一部は openscad-web-gui (GPLv3) から派生しています。
このディストリビューションには、OpenSCAD WASM からの未変更のバイナリが含まれています。
GPL v2 以降。結合作業の一部として、GPLv3 に基づいてここに配布されます。
src/vendor/openscad-wasm/SOURCE-OFFER.txt を参照してください。
RepoStars によるライブ チャート — クリックすると対話型バージョンが表示されます。
⭐ CADAM が役立つと思われる場合は、スターを付けることを検討してください。
3D プリンティングおよび CAD コミュニティ向けに 💙 で作られています
CADAM は、オープンソースのテキストから CAD への Web アプリケーションです。
GPL-3.0ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
543
フォーク
レポートリポジトリ
リリース
4
V0.3.0
最新の
2026 年 6 月 17 日
+ 3 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

CADAM is the open source text-to-CAD web application - Adam-CAD/CADAM

Hey HN! I'm Zach from Adam ( https://adam.new/ ). We're building AI agents for mechanical CAD software. We’ve built the company on two fundamental beliefs: - AI will be the primary medium for creating mechanical designs just like it is in software today. - The best paradigm for CAD generation is to generate CAD as code (text -> code -> CAD). We’re building CADAM, an open source Text to CAD platform. It's a React app (TanStack Start) with a Supabase backend for auth, database, and file storage. Think of it like AI TinkerCAD. Demo: https://www.youtube.com/watch?v=iESOr7EGWqk
Try it: https://adam.new/cadam/ What it does: - Generates parametric 3D models from natural language, with support for both text prompts and image references. - Outputs OpenSCAD code with automatically extracted parameters that surface as interactive sliders for instant dimension tweaking - Exports as .STL or .SCAD (plus OBJ, GLB/GLTF, FBX, and DXF) Under the hood: - One agentic endpoint with two modes that swap system prompts and tools: a parametric mode that writes/edits OpenSCAD via a build_parametric_model tool, and a mesh mode that generates 3D textured meshes. - Simple parameter tweaks bypass the model entirely; adjusting a slider does a deterministic regex update on the SCAD source, requiring no LLM call. - Model-agnostic via the Vercel AI SDK: Anthropic (Claude), Google (Gemini), and OpenAI/others through OpenRouter, with adaptive thinking auto-enabled on newer models. Surprisingly, in our evals Gemini 3.1 Pro is the top model. - Runs fully in-browser by compiling OpenSCAD to WebAssembly (in a Web Worker, so the UI never blocks) and rendering with Three.js via React Three Fiber - Supports BOSL, BOSL2, and MCAD libraries, plus custom font support (Geist) for text in models Future improvements: - Support both build123d and CadQuery. This will allow us to move beyond CSG primitives to constraint-driven modeling and provide direct comparisons to other code-as-CAD primitives. - Better spatial context: UI for face/edge selection and viewport image integration to give LLMs spatial understanding You can clone the repo and run it locally! Contributions are very welcome.

GitHub - Adam-CAD/CADAM: CADAM is the open source text-to-CAD web application · GitHub
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
Adam-CAD
/
CADAM
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
280 Commits 280 Commits .cursor/ rules .cursor/ rules .husky .husky benchmarks benchmarks public public scripts scripts shared shared src src supabase supabase .env.local.template .env.local.template .gitignore .gitignore .npmrc .npmrc .prettierignore .prettierignore .prettierrc .prettierrc CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md components.json components.json eslint.config.js eslint.config.js package-lock.json package-lock.json package.json package.json postcss.config.js postcss.config.js tailwind.config.js tailwind.config.js tsconfig.app.json tsconfig.app.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vercel.json vercel.json vite.config.ts vite.config.ts View all files Repository files navigation
⛮ The Open Source Text to CAD Web App ⛮
👉 adam.new/cadam . Generate a CAD model in seconds, right in your browser. No install required.
🤖 AI-Powered Generation - Transform natural language and images into 3D models
🎛️ Parametric Controls - Interactive sliders for instant dimension adjustments
📦 Multiple Export Formats - Export as .STL, .SCAD, or .DXF files
🌐 Browser-Based - Runs entirely in your browser using WebAssembly
📚 Library Support - Includes BOSL, BOSL2, and MCAD libraries
Feature
Description
Natural Language Input
Describe your 3D model in plain English
Image References
Upload images to guide model generation
Real-time Preview
See your model update instantly with Three.js
Parameter Extraction
Automatically identifies adjustable dimensions
Smart Updates
Efficient parameter changes without AI re-generation
Custom Fonts
Built-in Geist font support for text in models
📺 Screenshots
A showcase of what CADAM builds from a single plain-language description — from full multi-part machines down to clean parametric parts. Each model below started from the prompt shown and came out as fully parametric OpenSCAD, ready to export as .STL , .SCAD , or .DXF . The source and a short write-up for each live in benchmarks/ ; the orbiting previews are rendered with benchmarks/render.sh .
Model Prompt Controls Output
V8 engine
A complete V8 internal combustion engine: two banks of four cylinders in a 90° V, cylinder heads with ribbed valve covers, an intake manifold in the valley, exhaust headers down each bank, a crankshaft with counterweights, pistons and connecting rods, a front pulley, and an oil pan.
22 dims
8 colors
9-cylinder radial aircraft engine
Design a 9-cylinder radial aircraft engine: a central round crankcase with nine finned cylinders arranged evenly in a star pattern around it, each cylinder with stacked cooling fins and a domed cylinder head, and a central propeller shaft hub at the front.
15 dims
6 colors
Turbofan jet engine
A complete high-bypass turbofan: a front fan you can see into, a bypass cowl, an internal core with compressor/turbine stages, outlet guide vanes, and an exhaust plug.
2 dims
10 colors
Axial turbine blisk
Model an axial-flow turbine blisk (bladed disk) like a jet engine compressor stage: a central hub with a shaft bore and a single ring of about 28 thin aerofoil blades around the rim, each blade clearly twisted along its height from root to tip like a real turbine blade.
14 dims
1 color
Parametric fundamentals
Model Prompt Controls Output
Twisted hexagonal vase
Design a twisted hexagonal vase: a hollow shell about 150 mm tall that tapers from a 70 mm base to a 50 mm mouth, with the hexagonal cross-section twisting 120 degrees from bottom to top, a 2 mm wall, and a closed bottom.
6 dims
1 color
Knurled control knob
Make a knurled control knob 40 mm in diameter and 22 mm tall with a diamond-knurled grip, a raised pointer mark on top, a 6 mm D-shaped shaft bore, and an M3 set-screw hole through the side.
15 dims
2 colors
Hex bolt & nut — real threads
Create an M12 hex bolt 45 mm long with a real threaded shaft and a standard hex head, plus its matching hex nut, placed side by side.
3 dims
2 colors
Honeycomb lightweight bracket
Design a 90-degree angle mounting bracket with 80x80 mm flanges that are 5 mm thick, lightened with a hexagonal honeycomb cutout pattern on both faces, four M5 mounting holes, and a filleted inside corner.
13 dims
1 color
NACA 2412 tapered wing
Model a tapered aircraft wing section using a real NACA 2412 airfoil: 120 mm root chord tapering to 80 mm tip over a 200 mm span, with two spanwise spar tubes and a few lightening holes.
9 dims
1 color
Threaded jar & screw-on lid
Create a small storage jar with external screw threads at the neck and a matching screw-on lid with internal threads. Jar body 60 mm diameter, 70 mm tall, 2.5 mm walls; show the lid unscrewed and sitting beside the jar.
9 dims
2 colors
Right-angle bevel gear drive
Build a right-angle bevel gear drive: a 24-tooth bevel gear on a vertical shaft meshing at 90 degrees with a 16-tooth bevel pinion on a horizontal shaft, each on a short stub shaft.
9 dims
3 colors
Centrifugal pump impeller
Design a centrifugal pump impeller: a 90 mm diameter back-plate with a central 12 mm bore and a raised hub, and seven backward-curved blades that sweep from the hub out to the rim, each blade curving smoothly along its path.
10 dims
1 color
Herringbone planetary gear stage
Model a herringbone planetary gear
[truncated]
# Clone the repository
git clone https://github.com/Adam-CAD/CADAM.git
cd CADAM
# Install dependencies
npm install
# Start Supabase
npx supabase start
npx supabase functions serve --no-verify-jwt
# Start the development server
npm run dev
📋 Prerequisites
Node.js ^20.19.0 or >=22.12.0, with npm 10+
ngrok (for local webhook development)
🔧 Setting Up Environment Variables
Copy .env.local.template to .env.local
Update all required keys in .env.local :
VITE_SUPABASE_ANON_KEY="<Test Anon Key>"
VITE_SUPABASE_URL='http://127.0.0.1:54321'
Add server-side keys to .env.local , including:
ANTHROPIC_API_KEY="<Test Anthropic API Key>"
OPENROUTER_API_KEY="<Test OpenRouter API Key>"
OPENAI_API_KEY="<Test OpenAI API Key>"
GOOGLE_API_KEY="<Test Google API Key>"
FAL_KEY="<Test FAL API Key>"
SUPABASE_SERVICE_ROLE_KEY="<Test Service Role Key>"
BILLING_SERVICE_URL="<Test Billing Service URL>"
BILLING_SERVICE_KEY="<Test Billing Service Key>"
ENVIRONMENT="local"
ADAM_URL="<Adam URL or dev URL>" # Checkout and portal redirect target
WEBHOOK_BASE_URL="<Public TanStack App URL>" # Your app URL for /cadam/api callbacks
NGROK_URL="<NGROK URL>" # Optional local Supabase Storage tunnel for provider-readable signed URLs
🌐 Setting Up ngrok for Local Development
CADAM uses public URLs for provider callbacks and local signed storage URLs:
Install ngrok if you haven't already:
npm install -g ngrok
# or
brew install ngrok
Start an ngrok tunnel pointing to your TanStack Start dev server:
Copy the generated ngrok URL (e.g., https://xxxx-xx-xx-xxx-xx.ngrok.io ) and add it to your .env.local file:
WEBHOOK_BASE_URL="https://xxxx-xx-xx-xxx-xx.ngrok.io"
If a provider must fetch local Supabase Storage signed URLs, run a second tunnel to Supabase and set NGROK_URL to that URL.
Ensure ENVIRONMENT="local" is set in the same file.
npx supabase start
npm run dev
🛠️ Built With
Frontend: React 19 + TypeScript + TanStack Start + Vite
3D Rendering: Three.js + React Three Fiber
CAD Engine: OpenSCAD WebAssembly
Backend: TanStack Start server routes + Supabase PostgreSQL/Auth/Storage
Styling: Tailwind CSS + shadcn/ui
If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also open an issue .
See the CONTRIBUTING.md for instructions and code of conduct .
This app wouldn't be possible without the work of:
This distribution is licensed under the GNU General Public License v3.0 (GPLv3). See LICENSE .
Portions of this project are derived from openscad-web-gui (GPLv3).
This distribution includes unmodified binaries from OpenSCAD WASM under
GPL v2 or later; distributed here under GPLv3 as part of the combined work.
See src/vendor/openscad-wasm/SOURCE-OFFER.txt .
Live chart by RepoStars — click for the interactive version.
⭐ If you find CADAM useful, please consider giving it a star!
Made with 💙 for the 3D printing and CAD community
CADAM is the open source text-to-CAD web application
GPL-3.0 license
Code of conduct
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
543
forks
Report repository
Releases
4
V0.3.0
Latest
Jun 17, 2026
+ 3 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
