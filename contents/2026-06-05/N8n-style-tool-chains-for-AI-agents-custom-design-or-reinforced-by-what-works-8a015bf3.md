---
source: "https://github.com/pssah4/stigmergy"
hn_url: "https://news.ycombinator.com/item?id=48410036"
title: "N8n-style tool chains for AI agents – custom design, or reinforced by what works"
article_title: "GitHub - pssah4/stigmergy · GitHub"
author: "pssah4"
captured_at: "2026-06-05T10:25:26Z"
capture_tool: "hn-digest"
hn_id: 48410036
score: 1
comments: 0
posted_at: "2026-06-05T09:19:59Z"
tags:
  - hacker-news
  - translated
---

# N8n-style tool chains for AI agents – custom design, or reinforced by what works

- HN: [48410036](https://news.ycombinator.com/item?id=48410036)
- Source: [github.com](https://github.com/pssah4/stigmergy)
- Score: 1
- Comments: 0
- Posted: 2026-06-05T09:19:59Z

## Translation

タイトル: AI エージェント用の N8n スタイルのツール チェーン – カスタム設計、または機能するものによって強化
記事タイトル: GitHub - pssah4/stigmergy · GitHub
説明: GitHub でアカウントを作成して、pssah4/stigmergy の開発に貢献します。

記事本文:
GitHub - pssah4/stigmergy · GitHub
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
pssah4
/
汚名を着せられる
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル

2 コミット 2 コミット 例/ ketten-beispiel 例/ ketten-beispiel パッケージ パッケージ .gitignore .gitignore ライセンス ライセンス通知 通知 README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.base.json tsconfig.base.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
スティグマジーは、エージェント ループにすでに機能したことの記憶を与えます。ループの機能を監視します
受け入れられる結果につながる能力の道筋を目指し、学習し、証明された道筋を明らかにします。
次回同様のタスクが発生したときに、エージェントは実行のたびに同じ場所を再探索するのをやめます。走ります
ループの横にある場合は決してブロックせず、スイッチをオフにすると no-op になります。
エージェント ループは実行のたびにツールを再検出します。ツールセットが増えるにつれて (関数呼び出しツール、スキル、
MCP サーバー、サブエージェント)、モデルは、存在するものとそれを呼び出す方法、およびフラットを調査することに順番に費やします。
すべてのカタログがプロンプトを肥大化させます。ループにはメモリがありません: タスクを解決する一連のステップ
先週がなくなったので、同じ探索コストが再び支払われます。段階的な開示 (ツールをロードする)
request) はプロンプトを小さく保ちますが、以前に成功したことを思い出させるものは何もありません。
スティグマジーは、アリがフェロモンの足跡を残す方法にヒントを得た宿主の外部基質です。ループを観察します
そして受け入れられる遷移を強化するので、ある種のタスクに適したパスを学習します。
新しいタスクがすでに正常に解決されたパスと明らかに一致する場合、そのパスは
小さなヒントを提供し、パスに必要なツールを正確に事前にアクティブ化できるため、エージェントは再検出をスキップします。
タスクが認識されない場合、タスクは沈黙したままとなり、ループ自体のツール選択は変更されずに実行されます。
それは進歩的な開示を補完します。

置き換えません。取得した内容は正確なデフォルトのままです。
Stigmergy は最上位のリコール層です。ノンブロッキングです (コンサルトは 10 未満のメモリ内検索です)
ミリ秒、すべての埋め込みはバックグラウンドで実行されます)、オフまたはデーモンの場合は no-op に低下します。
は到達不可能であり、採点された結果からのみ学習するため、バタバタとした走りが強化されることはありません。
一度登録してください。ホストは、次の 4 種類すべてにわたってどの機能が存在するかをサブストレートに伝えます。
ツール、スキル、MCP ツール、サブエージェント。それぞれに安定した ID と短い説明が付いています。
ターンごと: コンサルト、楽器、グレード。モデルがステップを選択する前に、ループは基板に質問します。
どの実証済みのパスがタスクに適合するか。ツールが実行されると、ループ レポートが呼び出され、 が返されます。とき
ターンが解決されると、ループはそれを承認、反復、放棄のいずれかに評価します。
パスグラフ上のフェロモン。受け入れられたトランジションは、品質とトークンによって重み付けされたフェロモンを堆積します
効率性と時間の経過とともに減衰します。廃棄または放棄されたものは否定的な証拠となります。選択は
シードが与えられた場合の決定性 (乗法的望ましさプラス減衰したトンプソン サンプリング)
証拠)、したがって基板はループを推測することはありません。
2 番目のセレクターではないことを思い出してください。信頼性の高い一致 (コンテキストが適合する、学習された名前付きパス) では、
サブストレートはシーケンスを返し、ホストはそのパスの機能と 1 行のヒントを事前にアクティブ化します。
それ以外の場合は中立の結果を返し、ホスト自身の選択によって決定されます。並べ替えたり非表示にしたりすることはありません
あなたのツール。
ローカルでホットパスから外れています。セマンティック マッチングでは、ローカル埋め込みモデルを使用します。機能とクエリ
埋め込みはバックグラウンド ワーカーで計算されるため、コンサルトは同期メモリ内読み取りとなります。
Stigmergy は、Electron デスクトップ アプリ (Studio) と一連の機能を備えた npm-workspaces モノリポジトリです。
接続するためのシンクライアントパッケージ

あなたのループ。
git クローン https://github.com/pssah4/stigmergy.git
CDの汚名
npmインストール
npm ビルドを実行する
Studio (データベースとデーモンを所有し、セットアップ手順を案内するデスクトップ アプリ) を実行します。
npm run dev -w @stigmergy/studio
ネイティブ データベース モジュールが最初の起動時にバージョンの不一致を報告した場合は、デスクトップ用に再構築します。
ランタイム 1 回: npm run build:electron -w @stigmergy/studio 。初回実行ウィザードもこれを指摘しています。
スタジオから始めましょう。初めて起動すると、ガイド付きウィザードが何もない状態からセットアップを実行します。
ネイティブ モジュールをチェックし、メモリ ファイルを選択し、LLM プロバイダーを追加します (学習した名前付けにのみ使用されます)
パス、埋め込みはローカルのオンデバイス モデルです)、デーモンを起動し、エージェント ループに接続し、スイッチします。
スティグマジーオン。ウィザードは、[設定] からいつでも再度開くことができます。
ループを接続します。 Studio の接続ステップでは、コーディング エージェント (または
サポートされている SDK のスニペット) の 3 つの呼び出しをループに接続します。ツールを選択する前に確認してください。
インスツルメントツールの実行、および電源投入時の解像度のグレードを設定します。ループはシン クライアントのみを保持します。の
デーモンはモデルとデータベースを所有します。
サポートされている SDK の場合は、手動で接続する代わりに、対応するアダプターをインストールします。
それを見て学んでください。 Studio のグラフには、機能がノードとして表示され、学習したパスがエッジとして表示されます。
強化されると厚みが増します。保存したパスを検索してフィルターし、1 つをクリックして強調表示します。
グラフを表示するか、サイド パネルを折りたたんでグラフのみで作業します。
デーモンではなくエンジンをインプロセスで使用したい場合は、それを直接接続します。
import { createEngine } から '@agentic-stigmergy/core'
import { createSqlJsStorage } から '@stigmergy/storage-sqljs'
import { TransformersEmbedding } から '@stigmergy/embedding-transformers'
const エンジン = await createEngine (

{
storage : await createSqlJsStorage () , // またはノードの永続性のためのより良い sqlite3 ストア
embedding : new TransformersEmbedding () , // onnxruntime-node 上のローカル all-MiniLM-L6-v2
})
エンジンを待ちます。 registerCapability ( { id : 'tool:summarize' 、 type : 'tool' 、 description : '長い文書の要約' } )
const 決定 = エンジンを待ちます。コンサルト ( { task_id : 't1' 、 context : 'このレポートの要約' } )
// Decision.mode: 'ranked' (neutral)、または 'sequence' / パスが固定されている場合は 'enforce'。
エンジンを待ちます。 Emit ( { type : 'task_started' 、 taskId : 't1' 、 context : 'このレポートを要約する' } )
エンジンを待ちます。 Emit ( { type : 'capability_invoked' 、 taskId : 't1' 、capabilityId : 'tool:summarize' } )
エンジンを待ちます。 Emit ( { type : 'capability_returned' 、 taskId : 't1' 、capabilityId : 'tool:summarize' 、 success : true } )
エンジンを待ちます。 Emit ( { type : 'response_delivered' , taskId : 't1' } )
エンジンを待ちます。 Emit ( { type : 'task_accepted' , taskId : 't1' , tokenCost : 800 } ) // または失敗時に task_abandoned
エンジンを待ちます。閉じる（）
多くのタスクにわたって、サブストレートは受け入れられる遷移を強化するため、後でランクを参照します。
以前に成功したパスの方が高くなります。パスと結果がすでにわかっている場合は、ステップごとの説明をスキップしてください
イベントを実行し、engine.deposit({ task_id, context, path, result, token_cost }) を直接呼び出します。パスを固定する
常にそれを優先、強制、または順序付けします: Engine.pinPath({capability_sequence,behavior })。
パッケージ
目的
@agentic-stigmergy/core
エンジン、スコアリングと減衰、結果トラッカー、バックグラウンド エンベッダー、ポート
@agentic-stigmergy/loop
SDK に依存しないループ ファサード (beginTurn / 計器 / グレード)、デグレードセーフ
@agentic-stigmergy/クライアント
デーモンの Unix ソケットを介したシン クライアント
@stigmergy/storage-better-sqlite3
ノードストレージアダプター (WAL、ネイティブ)
@stigmergy/storage-sqlj

s
サンドボックス ストレージ アダプター (WASM、インメモリ)
@stigmergy/embedding-transformers
ローカル組み込みアダプター (onnxruntime-node 上の all-MiniLM-L6-v2)
開発
npm ワークスペース全体で build # tsc -b を実行します
npm test # vitest (オフライン、決定的)
埋め込みアダプターは env フラグの背後にある実際のモデルに対して検証されるため、デフォルトのスイートはそのまま残ります。
オフライン: STIGMERGY_EMBEDDING_IT=1 npm テストは、最初の実行時にモデルをダウンロードします。
https://open.substack.com/pub/sebastianhanke/p/stigmergy-for-capability-selection?r=1o93w5&utm_medium=ios
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

Contribute to pssah4/stigmergy development by creating an account on GitHub.

GitHub - pssah4/stigmergy · GitHub
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
pssah4
/
stigmergy
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits examples/ ketten-beispiel examples/ ketten-beispiel packages packages .gitignore .gitignore LICENSE LICENSE NOTICE NOTICE README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.base.json tsconfig.base.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Stigmergy gives an agent loop a memory of what already worked. It watches which capabilities a loop
reaches for, learns the capability paths that lead to accepted outcomes, and surfaces a proven path the
next time a similar task comes up, so the agent stops re-exploring the same ground every run. It runs
beside your loop, never blocks it, and turns into a no-op when switched off.
Agent loops rediscover their tools on every run. As the toolset grows (function-calling tools, skills,
MCP servers, sub-agents), the model spends turns probing what exists and how to call it, and a flat
catalog of everything bloats the prompt. The loop has no memory: a sequence of steps that solved a task
last week is gone, so the same exploration cost is paid again. Progressive disclosure (load tools on
demand) keeps the prompt small but does nothing to recall what succeeded before.
Stigmergy is a host-external substrate, inspired by how ants lay pheromone trails. It observes the loop
and reinforces the transitions that get accepted, so it learns the paths that work for a kind of task.
When a new task clearly matches a path it has already solved successfully, it surfaces that path as a
small hint and can pre-activate exactly the tools the path needs, so the agent skips the rediscovery.
When it does not recognize the task, it stays silent and your loop's own tool selection runs unchanged.
It complements progressive disclosure, it does not replace it: your retrieval stays the precise default,
Stigmergy is the recall layer on top. It is non-blocking (a consult is an in-memory lookup under ten
milliseconds, all embedding runs in the background), it degrades to a no-op when it is off or the daemon
is unreachable, and it learns only from graded outcomes, so a flailing run is never reinforced.
Register once. The host tells the substrate which capabilities exist, across all four kinds:
tools, skills, MCP tools, and sub-agents, each with a stable id and a short description.
Per turn: consult, instrument, grade. Before the model picks a step, the loop asks the substrate
which proven path fits the task. As tools run, the loop reports invoked and returned . When the
turn resolves, the loop grades it: accepted, iterated, or abandoned.
Pheromone on a path graph. Accepted transitions deposit pheromone, weighted by quality and token
efficiency and decayed over time; discarded or abandoned ones accrue negative evidence. Selection is
deterministic given a seed (multiplicative desirability plus Thompson sampling over the decayed
evidence), so the substrate never makes the loop guess.
Recall, not a second selector. On a confident match (a learned, named path whose context fits) the
substrate returns a sequence and the host pre-activates that path's capabilities plus a one-line hint.
Otherwise it returns a neutral result and the host's own selection decides. It never reorders or hides
your tools.
Local and off the hot path. Semantic matching uses a local embedding model; capability and query
embeddings are computed in a background worker, so a consult is a synchronous in-memory read.
Stigmergy is an npm-workspaces monorepo with an Electron desktop app (the Studio) and a set of
thin-client packages for wiring it into your loop.
git clone https://github.com/pssah4/stigmergy.git
cd stigmergy
npm install
npm run build
Run the Studio (the desktop app that owns the database and the daemon, and walks you through setup):
npm run dev -w @stigmergy/studio
If the native database module reports a version mismatch on first launch, rebuild it for the desktop
runtime once: npm run rebuild:electron -w @stigmergy/studio . The first-run wizard also points this out.
Start in the Studio. On first launch a guided wizard takes you from nothing to a running setup: it
checks the native module, lets you pick a memory file, add an LLM provider (used only for naming learned
paths, the embedding is a local on-device model), start the daemon, connect your agent loop, and switch
Stigmergy on. You can re-open the wizard any time from Settings.
Connect your loop. The Studio's connect step gives you a copy-paste prompt for your coding agent (or
a snippet for a supported SDK) that wires three calls into your loop: consult before tool selection,
instrument tool execution, and grade the turn on resolution. Your loop holds only the thin client; the
daemon owns the model and the database.
For a supported SDK, install the matching adapter instead of hand-wiring:
Watch it learn. The Studio's graph shows your capabilities as nodes and the learned paths as edges
that thicken as they are reinforced. Search and filter your saved paths, click one to highlight it in
the graph, or collapse the side panel to work in the graph alone.
If you want the engine in-process rather than the daemon, wire it directly:
import { createEngine } from '@agentic-stigmergy/core'
import { createSqlJsStorage } from '@stigmergy/storage-sqljs'
import { TransformersEmbedding } from '@stigmergy/embedding-transformers'
const engine = await createEngine ( {
storage : await createSqlJsStorage ( ) , // or a better-sqlite3 store for Node persistence
embedding : new TransformersEmbedding ( ) , // local all-MiniLM-L6-v2 on onnxruntime-node
} )
await engine . registerCapability ( { id : 'tool:summarize' , type : 'tool' , description : 'summarize long documents' } )
const decision = await engine . consult ( { task_id : 't1' , context : 'summarize this report' } )
// decision.mode: 'ranked' (neutral), or 'sequence' / 'enforce' when a path is pinned.
await engine . emit ( { type : 'task_started' , taskId : 't1' , context : 'summarize this report' } )
await engine . emit ( { type : 'capability_invoked' , taskId : 't1' , capabilityId : 'tool:summarize' } )
await engine . emit ( { type : 'capability_returned' , taskId : 't1' , capabilityId : 'tool:summarize' , success : true } )
await engine . emit ( { type : 'response_delivered' , taskId : 't1' } )
await engine . emit ( { type : 'task_accepted' , taskId : 't1' , tokenCost : 800 } ) // or task_abandoned on failure
await engine . close ( )
Over many tasks the substrate reinforces transitions that get accepted, so later consults rank
previously successful paths higher. If you already know the path and the outcome, skip the per-step
events and call engine.deposit({ task_id, context, path, outcome, token_cost }) directly. Pin a path to
always prefer, enforce, or sequence it: engine.pinPath({ capability_sequence, behavior }) .
Package
Purpose
@agentic-stigmergy/core
Engine, scoring and decay, outcome tracker, background embedder, ports
@agentic-stigmergy/loop
SDK-agnostic loop facade (beginTurn / instrument / grade), degrade-safe
@agentic-stigmergy/client
Thin client over the daemon's Unix socket
@stigmergy/storage-better-sqlite3
Node storage adapter (WAL, native)
@stigmergy/storage-sqljs
Sandbox storage adapter (WASM, in-memory)
@stigmergy/embedding-transformers
Local embedding adapter (all-MiniLM-L6-v2 on onnxruntime-node)
Development
npm run build # tsc -b across the workspace
npm test # vitest (offline, deterministic)
The embedding adapter is verified against the real model behind an env flag, so the default suite stays
offline: STIGMERGY_EMBEDDING_IT=1 npm test downloads the model on first run.
https://open.substack.com/pub/sebastianhanke/p/stigmergy-for-capability-selection?r=1o93w5&utm_medium=ios
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
