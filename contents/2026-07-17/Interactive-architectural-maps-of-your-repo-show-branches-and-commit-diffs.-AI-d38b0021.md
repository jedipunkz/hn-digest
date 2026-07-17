---
source: "https://github.com/agustinvillegas/repomap"
hn_url: "https://news.ycombinator.com/item?id=48942220"
title: "Interactive architectural maps of your repo, show branches and commit diffs. AI"
article_title: "GitHub - agustinvillegas/repomap: RepoMap is a skill for cli agents, allows you to map your code and review branches or commits diff · GitHub"
author: "franvillegass"
captured_at: "2026-07-17T01:47:38Z"
capture_tool: "hn-digest"
hn_id: 48942220
score: 1
comments: 1
posted_at: "2026-07-17T00:58:55Z"
tags:
  - hacker-news
  - translated
---

# Interactive architectural maps of your repo, show branches and commit diffs. AI

- HN: [48942220](https://news.ycombinator.com/item?id=48942220)
- Source: [github.com](https://github.com/agustinvillegas/repomap)
- Score: 1
- Comments: 1
- Posted: 2026-07-17T00:58:55Z

## Translation

タイトル: リポジトリのインタラクティブなアーキテクチャ マップ、ブランチの表示、差分のコミット。 AI
記事のタイトル: GitHub - agustinvillegas/repomap: RepoMap は CLI エージェントのスキルであり、コードをマップしてブランチやコミットの差分を確認できます · GitHub
説明: RepoMap は cli エージェント用のスキルであり、コードをマップしてブランチまたはコミットの差分を確認できます - GitHub - agustinvillegas/repomap: RepoMap は cli エージェント用のスキルで、コードをマップしてブランチまたはコミットの差分を確認できます

記事本文:
GitHub - agustinvillegas/repomap: RepoMap は cli エージェント用のスキルで、コードをマップしてブランチやコミット diff を確認できます · GitHub
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
アグスティンヴィルガス
/
リポマップ
公共
通知
通知設定を変更するにはサインインする必要があります

GS
追加のナビゲーション オプション
コード
スキル ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
66 コミット 66 コミット .opencode/ スキル/ リポマップ .opencode/ スキル/ リポマップ アセット アセット .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md settings.json settings.json すべてのファイルを表示リポジトリ ファイルのナビゲーション
コーディング担当者にアーキテクチャに関する認識を与えます。
RepoMap は、ソース コードを LLM に送信せずにリポジトリの構造を抽出し、人間と AI エージェントの両方が理解できる対話型のアーキテクチャ マップを生成します。
最新のコーディング エージェントは、プロジェクトのアーキテクチャを再構築するのにかなりの時間を費やします。
次のような質問に答えるためだけに、ファイルを繰り返し開き、インポートを追跡し、フォルダーを検査し、何千ものトークンを消費します。
認証はどこで実装されていますか?
どのモジュールがデータベースに依存しますか?
このプロジェクトはどのようなアーキテクチャ パターンに従っていますか?
コンポーネントはどのように接続されていますか?
RepoMap はそのワークフローを変えます。
LLM にセッションごとにアーキテクチャを再検出するよう依頼する代わりに、RepoMap は決定論的な構造表現を一度構築し、エージェントにその表現を推論させます。
リポジトリの理解に費やされるトークンが大幅に減少
より高速なアーキテクチャ上の推論
人間のためのインタラクティブな視覚化
AI エージェント用の再利用可能なアーキテクチャ データ
特長
✅ 決定論的なリポジトリ分析
✅ LLM に送信されるソースコードはゼロ
✅ インタラクティブなアーキテクチャグラフ
✅ アーキテクチャ上の役割の自動検出
✅ アーキテクチャパターン認識
✅ 永続的な分岐を持つ編集可能なグラフ
RepoMap は、構造の抽出をアーキテクチャの推論から分離します。
リポジトリ
│
▼
決定論的アナライザー
│
▼
生分析
(ソースコードはありません)
│
▼
LLM

(アーキテクチャの推論)
│
▼
リポグラフ
│
▼
インタラクティブな視覚化
フェーズ 1 — 決定論的分析
アナライザーは LLM を呼び出すことはありません。
リポジトリをスキャンして以下を抽出します。
ディレクトリ階層
輸入品
関数シグネチャ
モジュール関係
git情報
その結果、ソース コードを公開せずにリポジトリ構造を記述するコンパクトな RawAnalysis が得られます。
フェーズ 2 — アーキテクチャの推論
LLM は、何千ものファイルを読み取る代わりに、構造化分析のみを受け取ります。
アーキテクチャ上の役割を割り当てる
アーキテクチャパターンを特定する
モジュールラベルを改善する
視覚化レイアウトを生成する
必要な推論ステップは 1 つだけです。
フェーズ 3 — インタラクティブな探索
RepoMap は、生成されたアーキテクチャを編集可能なグラフとしてレンダリングします。
React Flow の視覚化
複数のレイアウト
ブランチベースの編集
ノード検査
ビューポートのカリング
永続的なローカルストレージ
Git 対応の視覚化
RepoMap を使用すると、Git 履歴とともにリポジトリを探索できます。
コミットを検査する
ブランチを参照する
追加されたファイルを強調表示する
変更されたファイルを強調表示する
削除されたファイルを強調表示する
将来のバージョンでは、これを完全なアーキテクチャの差分視覚化に拡張します。
LLM にリポジトリを読み取らせてみませんか?
リポジトリの探索はほとんどが決定的であるためです。
インポートの検索、モジュールの検出、フォルダーの追跡、定義の抽出にはインテリジェンスは必要ありません。
LLM は、ソフトウェアが自動的に抽出できる情報を再構築するのではなく、アーキテクチャに関するコンテキスト推論に費やす必要があります。
これにより、トークンの使用量が大幅に削減され、より豊かなアーキテクチャ表現が生成されます。
RepoMap は、OpenCode や Claude などのツールと自然に統合するように構築されました。
エージェントは、リポジトリを繰り返し探索する代わりに、推論のために再利用できるコンパクトな構造モデルを受け取ります。

と将来の分析。
git clone < リポジトリ URL >
cd リポマップ-パイプライン-v2
npmインストール
スキルは OpenCode と Claude によって .opencode/skills/repomap/ から自動検出されます。
# ローカル リポジトリを分析し、インタラクティブ マップを提供します
ノード .opencode/skills/repomap/cli.js 私のリポジトリを開く
# またはコードから:
import { 分析 } から ' ./.opencode/skills/repomap/index.js '
const raw = awaitanalyze({ localPath: ' /path/to/repo ' })
# 保存されているすべてのマップをリストする
ノード .opencode/skills/repomap/cli.js リスト
リポグラフ形式
{
「メタ」: {
"repoName" : " 所有者/リポジトリ " ,
"estimatedSize" : " small|medium|large " 、
"言語" : [ " TypeScript " 、 " Python " ]、
"合計ファイル数" : 42 、
「合計モジュール」: 8
}、
「ノード」: [
{ "id" : "layer__src " 、 "label" : " Src " 、 "type" : "layer " 、 "parentId" : null 、 "files" : []、 "metadata" : {} }、
{ "id" : " module__src_core " 、 "label" : " Core " 、 "type" : " module " 、 "parentId" : "layer__src " 、 "files" : [ " core/index.ts " ]、 "metadata" : {} }、
{ "id" : " file__core_index_ts " 、 "label" : "index.ts " 、 "type" : " file " 、 "parentId" : " module__src_core " 、 "files" : [ " core/index.ts " ]、 "metadata" : { " language" : " TypeScript " 、 "lineCount" : 150 、 "complexity" : "中 " } }
]、
「エッジ」: [
{ "id" : "edge__file1__file2 " 、 "source" : " file__core_index_ts " 、 "target" : " file__utils_helpers_ts " 、 "edgeType" : " Engineering " 、 "strength" : 3 、 "label" : " imports " 、 "confidence" : " high " },
{ "id" : "edge__layer__module" 、"source" : "layer__src " 、"target" : " module__src_core " 、 "edgeType" : "アーキテクチャ " 、 "strength" : 5 、 "label" : " contains " 、 "confidence" : " high " }
]、
「オーバーレイ」: {
「バージョン」: 0 、
"nodeOverrides" : {}、
"edgeOverrides" : {}、
"manualNodes" : [],
"manualEdges" : []
}
}
CLI
ノード cli.js list 保存されているすべてのマップをリストします。
ノードcli。

js open <name> 名前、ファイル名、または部分一致でマップを開きます
すべてのコマンドは .opencode/skills/repomap/ から実行されます。
.opencode/スキル/リポマップ/
Analyzer.js 決定論的スキャナー: ファイル ツリー、インポート、定義、署名、Git データ
Index.js オーケストレーション: 分析 → 永続化 → 視覚化の提供
cli.js CLI ラッパー (リスト、オープン)
Visual_src/ React Flow ベースのインタラクティブなグラフ レンダラー (@frannn2114/repomap-visual として公開)
パイプラインは次のとおりです。
Analyzer.js は LLM を呼び出さずに構造を抽出します
Index.js は結果を RepoGraph JSON として保存し、ビジュアル サーバーを生成します
cli.js は、保存されたマップの list コマンドと open コマンドを提供します
RepoMap は CLI エージェント用のスキルで、コードをマップし、ブランチやコミットの差分を確認できます。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

RepoMap is a skill for cli agents, allows you to map your code and review branches or commits diff - GitHub - agustinvillegas/repomap: RepoMap is a skill for cli agents, allows you to map your code and review branches or commits diff

GitHub - agustinvillegas/repomap: RepoMap is a skill for cli agents, allows you to map your code and review branches or commits diff · GitHub
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
agustinvillegas
/
repomap
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
skill Branches Tags Go to file Code Open more actions menu Folders and files
66 Commits 66 Commits .opencode/ skills/ repomap .opencode/ skills/ repomap assets assets .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md settings.json settings.json View all files Repository files navigation
Give coding agents architectural awareness.
RepoMap extracts the structure of any repository without sending source code to an LLM, then generates an interactive architectural map that both humans and AI agents can understand.
Modern coding agents spend a significant amount of time reconstructing a project's architecture.
They repeatedly open files, follow imports, inspect folders, and consume thousands of tokens just to answer questions like:
Where is authentication implemented?
Which modules depend on the database?
What architectural pattern does this project follow?
How are components connected?
RepoMap changes that workflow.
Instead of asking the LLM to rediscover the architecture every session, RepoMap builds a deterministic structural representation once and lets the agent reason over that representation.
significantly fewer tokens spent on repository understanding
faster architectural reasoning
interactive visualization for humans
reusable architecture data for AI agents
Features
✅ Deterministic repository analysis
✅ Zero source code sent to the LLM
✅ Interactive architecture graph
✅ Automatic architectural role detection
✅ Architecture pattern recognition
✅ Editable graphs with persistent branches
RepoMap separates structure extraction from architectural reasoning.
Repository
│
▼
Deterministic Analyzer
│
▼
RawAnalysis
(no source code)
│
▼
LLM
(architecture reasoning)
│
▼
RepoGraph
│
▼
Interactive visualization
Phase 1 — Deterministic Analysis
The analyzer never calls an LLM.
It scans the repository and extracts:
directory hierarchy
imports
function signatures
module relationships
git information
The result is a compact RawAnalysis that describes the repository structure without exposing the source code.
Phase 2 — Architectural Reasoning
Instead of reading thousands of files, the LLM receives only the structured analysis.
assign architectural roles
identify architectural patterns
improve module labels
generate a visualization layout
Only one reasoning step is required.
Phase 3 — Interactive Exploration
RepoMap renders the generated architecture as an editable graph.
React Flow visualization
multiple layouts
branch-based editing
node inspection
viewport culling
persistent local storage
Git-aware visualization
RepoMap lets you explore repositories together with their Git history.
inspect commits
browse branches
highlight added files
highlight modified files
highlight deleted files
Future versions will expand this into full architectural diff visualization.
Why not let the LLM read the repository?
Because repository exploration is mostly deterministic.
Finding imports, discovering modules, following folders and extracting definitions does not require intelligence.
LLMs should spend their context reasoning about architecture—not reconstructing information that software can extract automatically.
This dramatically reduces token usage while producing a richer architectural representation.
RepoMap was built to integrate naturally with tools such as OpenCode and Claude.
Instead of repeatedly exploring the repository, agents receive a compact structural model that can be reused for reasoning, visualization and future analysis.
git clone < repo-url >
cd repomap-pipeline-v2
npm install
The skill is auto-discovered by OpenCode and Claude from .opencode/skills/repomap/ .
# Analyze a local repository and serve the interactive map
node .opencode/skills/repomap/cli.js open my-repo
# Or from code:
import { analyze } from ' ./.opencode/skills/repomap/index.js '
const raw = await analyze({ localPath: ' /path/to/repo ' })
# List all saved maps
node .opencode/skills/repomap/cli.js list
RepoGraph format
{
"meta" : {
"repoName" : " owner/repo " ,
"estimatedSize" : " small|medium|large " ,
"languages" : [ " TypeScript " , " Python " ],
"totalFiles" : 42 ,
"totalModules" : 8
},
"nodes" : [
{ "id" : " layer__src " , "label" : " Src " , "type" : " layer " , "parentId" : null , "files" : [], "metadata" : {} },
{ "id" : " module__src_core " , "label" : " Core " , "type" : " module " , "parentId" : " layer__src " , "files" : [ " core/index.ts " ], "metadata" : {} },
{ "id" : " file__core_index_ts " , "label" : " index.ts " , "type" : " file " , "parentId" : " module__src_core " , "files" : [ " core/index.ts " ], "metadata" : { "language" : " TypeScript " , "lineCount" : 150 , "complexity" : " medium " } }
],
"edges" : [
{ "id" : " edge__file1__file2 " , "source" : " file__core_index_ts " , "target" : " file__utils_helpers_ts " , "edgeType" : " engineering " , "strength" : 3 , "label" : " imports " , "confidence" : " high " },
{ "id" : " edge__layer__module " , "source" : " layer__src " , "target" : " module__src_core " , "edgeType" : " architecture " , "strength" : 5 , "label" : " contains " , "confidence" : " high " }
],
"overlay" : {
"version" : 0 ,
"nodeOverrides" : {},
"edgeOverrides" : {},
"manualNodes" : [],
"manualEdges" : []
}
}
CLI
node cli.js list List all saved maps
node cli.js open <name> Open a map by name, file name, or partial match
All commands run from .opencode/skills/repomap/ .
.opencode/skills/repomap/
analyzer.js Deterministic scanner: file tree, imports, definitions, signatures, git data
index.js Orchestration: analysis → persist → serve visualization
cli.js CLI wrapper (list, open)
visual_src/ React Flow-based interactive graph renderer (published as @frannn2114/repomap-visual)
The pipeline is:
analyzer.js extracts structure without calling an LLM
index.js saves the result as a RepoGraph JSON and spawns a visual server
cli.js provides list and open commands for saved maps
RepoMap is a skill for cli agents, allows you to map your code and review branches or commits diff
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
