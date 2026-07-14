---
source: "https://github.com/OyaAIProd/oya"
hn_url: "https://news.ycombinator.com/item?id=48907336"
title: "Oya – Keep tool outputs away from the LLM to cut tokens and stop injection"
article_title: "GitHub - OyaAIProd/oya: Agents that plan, don't react. The model writes a typed plan once. The runtime runs it. Tool outputs never go back through the model — so your agent can't be prompt-injected through its tools, costs a fraction, and runs the same every time. · GitHub"
author: "oyadoti"
captured_at: "2026-07-14T15:18:33Z"
capture_tool: "hn-digest"
hn_id: 48907336
score: 2
comments: 1
posted_at: "2026-07-14T14:20:06Z"
tags:
  - hacker-news
  - translated
---

# Oya – Keep tool outputs away from the LLM to cut tokens and stop injection

- HN: [48907336](https://news.ycombinator.com/item?id=48907336)
- Source: [github.com](https://github.com/OyaAIProd/oya)
- Score: 2
- Comments: 1
- Posted: 2026-07-14T14:20:06Z

## Translation

タイトル: Oya – トークンをカットしてインジェクションを停止するために、ツールの出力を LLM から遠ざける
記事のタイトル: GitHub - OyaAIProd/oya: 計画するエージェントは反応しない。モデルは、型付きの計画を 1 回書き込みます。ランタイムがそれを実行します。ツールの出力がモデルを通過することはありません。そのため、エージェントはツールを通じてプロンプト挿入されず、コストはほんのわずかで、毎回同じように実行されます。 · GitHub
説明: エージェントは計画を立てても、反応しません。モデルは、型付きの計画を 1 回書き込みます。ランタイムがそれを実行します。ツールの出力がモデルを通過することはありません。そのため、エージェントはツールを通じてプロンプト挿入されず、コストはほんのわずかで、毎回同じように実行されます。 - OyaAIProd/oya

記事本文:
GitHub - OyaAIProd/oya: エージェントは計画しても反応しません。モデルは、型付きの計画を 1 回書き込みます。ランタイムがそれを実行します。ツールの出力がモデルを通過することはありません。そのため、エージェントはツールを通じてプロンプト挿入されず、コストはほんのわずかで、毎回同じように実行されます。 · GitHub
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
別のタブまたはウィンドウでアカウントを切り替えました。リロードして更新してください

セッション。
アラートを閉じる
{{ メッセージ }}
OyaAIProd
/
オヤ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
34 コミット 34 コミット .github .github apps/ playground apps/ playground ベンチマーク ベンチマーク docs docs パッケージ パッケージ スクリプト script .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md DEMO.md DEMO.md GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md bun.lock bun.lockデモ.gif デモ.gif デモ.テープ デモ.テープ oya.config.ts oya.config.ts package.json package.json tsconfig.base.json tsconfig.base.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
エージェントは、すべてのツールの結果をモデルを通じて再フィードします。それがバグです。
ReAct、LangGraph、Mastra、Vercel AI SDK などのすべてのフレームワークでモデルをループします。
各ステップ: ツールの出力を読み取り、それを次のツールの引数として再入力し、
往復の料金を支払います。 Oya は、型付きプランを 1 回コンパイルし、DAG を実行します。
値は参照によって流れ、モデルを介して戻ることはありません。
同じコードです。同じツールです。トークンの数が 10 倍少なく、3.5 倍の速度があり、決定的であり、構造により注入が安全です。
· TypeScript · Bun · MIT · Mastra 用ドロップイン
クイックスタート · 数値 · 理由 · 2 行で移行 · スタジオ · ドキュメント · ホワイトペーパー
oya.ai の背後にあるオープンソース コア - Plan-Don't React エージェント向けのホスト型プラットフォーム。
import { Agent , createTool } from "oyadotai" ;
import { anthropic } from "oyadotai/anthropic" ;
"zod" から { z } をインポートします。
const getWeather = createTool ( {
id : "get_weather" ,
description : "都市の天気を調べる" ,
inputSchema : z 。オブジェクト ( { city : z . string ( ) } )

、
実行: async ({ city } ) => fetchWeather ( city ) 、
} ) ;
const エージェント = 新しいエージェント ( {
モデル: anthropic ( "claude-haiku-4-5-20251001" ) 、
ツール: { get_weather : getWeather } 、
} ) ;
const {text} = エージェントを待ちます。 generate (「ニューヨークの天気はどうですか?」) ;
これが API 全体です - Mastra と同じ形状です。型はあなたの情報から推測されます
zod スキーマであり、デフォルトではすべての値がモデルに対して不透明です。あなたの仕事は減りました
注射しても安全です。
トークン ループは、次のモデルを通じてすべてのツールの結果を再送信します。
ステップ、トークンの消費、レイテンシ、および決定論を必要としない状態に移行する
読んでください。 oya はプランを一度コンパイルし、参照によって各値を接続します。の
調整タスク - トランザクションを取得 → そのレコードを取得 → 正規化 → 検証 →
post - 実際の Anthropic API ( claude-haiku-4-5 、
同一のツール、8 回のトライアル):
oya は、最も無駄のないループよりも 2.3 倍少ないトークンで同じ結果をもたらします。
Mastra の 10 分の 1、実測時間の 3 分の 1 で、実際にそうなります
すべての実行で同様に。レコードのかさばるペイロードがループのコンテキストに再び入る
あらゆる段階で。 oya の下では、モデルには決して表示されない不透明なハンドルのままです。
状態の忠実度は保証であり、平均ではありません。すべての値は参照によってフローされます
モデルが再読み取りまたは再送信することのない不透明なハンドルとして、URL、ID、または
量はバイト単位で次のツールに配信され、実行順序は
静的にチェックされた DAG - トークン ループの結果は近似することしかできず、サンプリングされた 1 つです
一度に実行します。自分で再現してください: bun run bench 。
信仰に基づいてテーブルに着くのではなく、それを実行してください。ベンチマークは実際の Anthropic に一致します
3 つすべてで同一のタスクと同一のツール実装を備えた API
フレームワーク (ベンチマーク/tasks.ts);唯一のことは
状態フローの量によって異なります

モデルを通して。
# 1. クローンを作成してインストールします (Bun ≥ 1.1 - https://bun.sh が必要です)
git clone https://github.com/OyaAIProd/oya && cd oya
バンインストール
# 2. ライブラリをビルドします - ベンチマークは、ビルドされた `oyadotai` を dist/ からインポートします
バンランビルド
# 3. Anthropic キーを指定します (環境変数、またはリポジトリのルートにある .env にドロップします)。
エクスポート ANTHROPIC_API_KEY=sk-ant-...
#4. 比較を実行する
bun run bench # デフォルト: reconcile、claude-haiku-4-5、3 試行
bun run bench claude-sonnet-5 # 最初の引数として任意のモデル ID
bun run bench --task Research # 重いマルチドキュメントのケース
引数は任意の順序で指定します: モデル ID (デフォルトは claude-haiku-4-5-20251001 )、
トライアル数 (整数、デフォルトは 3 )、および --task reconcile / --taskpayment
/ --task Research / --task Weather (デフォルトは reconcile )。好む
make bench - ビルドを実行し、リポジトリのルートに .env を自動ロードします。
したがって、手順 2 ～ 4 は 1 つのコマンドにまとめられます。
デフォルトの調整タスクは、マルチホップ パイプラインを通じてクリティカル トークンをスレッドします。
フェッチされたレコードは、類似したディストラクターを運ぶかさばるペイロードです。それ
トークンの無駄を測定します (ペイロードが各ステップでループのコンテキストに再び入る、
oya はそれを OPAQUE に保ちますが、状態の忠実度 (来歴台帳によって確認されます)。
トークンがバイト単位で最終ツールに到達し、順序付け (静的にチェックされる)
DAG）。 oya は、構造上、実行ごとに 0 破損と 1 つの固定順序を返します。
--タスク調査では、比較を大量の複数ドキュメントのワークロードにまで拡張します。
各フレームワークは同じモデルに対してタスクを N 回実行し、トークンを出力します。
レイテンシーと正確性を隣り合わせに。その方法論は、
ベンチマーク/README.md 。
他のすべてのエージェント フレームワークはトークン ループ (ReAct、LangGraph、AutoGen、Mastra、
Vercel AI SDK): モデルはツールを選択し、生の結果を確認して、次のツールを選択します。
すべての URL、ID、および d

眼球はモデルを通って逆流します。 3 つのバグが続きます - すべて
時間:
取得済み: https://example.io/q3-report.pdf
ダウンロード: https://example.com/q3-report.pdf ← モデルは URL を「修正」しました
期待される内容: 取得 → 検証 → ダウンロード
観察: フェッチ → ダウンロード → 検証 (スキップ) ← モデルはステップを並べ替えました
トークンで 432 ドル ← モデルを通じてすべての結果を再読み取り
トークンで 51 ドル ← 決定する必要があるもののみを読み取る
根本原因は同じです。モデルが状態を読み取る必要がありませんでした。オヤはそれを作ります
不可能です。モデルは型付きのデータフロー プランを出力します。ランタイムは DAG を実行します。
各値は、計画が宣言するレベルでのみモデルに表示されます。
攻撃者は、取得したページにペイロードを詰め込むことができます。モデルはそれを決して読みません
handle であるため、ツールの出力を介した間接的なプロンプト挿入には着地点がありません。あなた
これには何も注釈を付けません。デフォルトでは不透明です。
Mastraから2行で移行
アプリがすでに @mastra/core を使用している場合、oya はドロップインに近いです。作成ツール
エージェントはすでに作成した形状をミラーリングするため、ほとんどのアプリでは全体が
移行はインポート行です。
- "@mastra/core/agent" から { Agent } をインポートします。
- "@mastra/core/tools" から { createTool } をインポートします。
- "@ai-sdk/anthropic" から { anthropic } をインポートします。
+ import { Agent, createTool } from "oyadotai";
+ import { anthropic } from "oyadotai/anthropic";
ステップ 2 - ステップ 2 はありません。ツールとエージェントの定義は変更されずにコンパイルされます。
マストラ
オヤ
注意事項
createTool({ ID, 説明, inputSchema, 実行 })
同じ
型は zod inputSchema から推論されます。 execute の引数は自動的に型付けされます。キャストはありません。
新しいエージェント({ 名前、手順、モデル、ツール })
同じ
toolsは同じ名前→ツールマップです。
await Agent.generate(prompt) → { text }
同じ
text は見出しフィールドであり、トークンの使用状況はパリティとして報告されます。
エージェント.ストリーム(プロ

mpt）
同じ通話
生のトークン スープの代わりに構造化イベント ( fullStream / textStream ) を返します。
人間的(…) · オープンアイ(…) · グーグル(…)
oyadotai/anthropic · …/openai · …/google
同じ呼び出し形状ですが、インポート パスが異なります。
したがって、すでにあるコードはそのまま動作し続けます。
import { Agent , createTool } from "oyadotai" ;
import { anthropic } from "oyadotai/anthropic" ;
"zod" から { z } をインポートします。
const getWeather = createTool ( {
id : "get_weather" ,
description : "都市の天気を調べる" ,
inputSchema : z 。オブジェクト ({ city : z . string () } ) 、
実行: async ({ city } ) => fetchWeather ( city ) 、
} ) ;
const エージェント = 新しいエージェント ( {
名前: "WeatherBot" 、
指示:「天気に関する質問に答えてください。」 、
モデル: anthropic ( "claude-haiku-4-5-20251001" ) 、
ツール: { get_weather : getWeather } 、
} ) ;
const {text} = エージェントを待ちます。 generate (「ニューヨークの天気はどうですか?」) ;
その下で何が変わるのか
同じ路面、異なるエンジン。 Mastra はトークン ループを実行します。モデルは
ツールを使用し、生の結果を読み取り、次の呼び出しを決定します - すべての値を再読み取りします
モデルを通して。 oya はモデルに 1 つの型付きプランを発行させ、その DAG を実行します
直接的に。値は参照によってツール間を移動し、モデルに開示されます。
宣言された投影レベル (デフォルトでは OPAQUE) でのみ。
コードには手を加えていませんが、移行されたアプリは無料で次の機能を利用できるようになります。
マルチホップ タスクではトークンが最大 10 倍少なく、最大 3 倍高速になります (上記の数値) - 中間状態はツール間でパイプされ、モデルを通じて再請求されることはありません。
決定的実行 - ステップを並べ替えたりスキップしたりできるモデルが選択したシーケンスではなく、実行ごとに 1 つの静的にチェックされる順序。
構造によりインジェクションセーフ - モデルが決して読み取らないツール出力は、インジェクションされた命令をそのコンテキストに密かに持ち込むことができません。
前に知っておいてよかった

あなたは移住します
投影のデフォルトは OPAQUE です。ツールの出力はモデルから非表示になります
SUMMARY または TRANSPARENT とマークしない限り。それは安全性の勝利です - しかし、もし
ツールは、生の出力を読み取るモデルに依存して次のステップを決定します。
その出力の投影。 「投影タイプ」を参照してください。
プロバイダーは oyadotai/* からインポートします ( oyadotai/anthropic 、 oyadotai/openai 、
@ai-sdk/* ではなく oyadotai/google )。
oya は 1.0 より前であり、すべてのヘルパーではなく、Mastra サーフェスのコアを反映しています。
依存しているものが不足している場合は、問題をオープンしてください。Mastra パリティ ギャップは優先度が高くなります。
エージェントとチャットし、各プランのライブ実行 - DAG (React Flow、ノード) を監視します。
種類ごとに色が付けられ、実行時に点灯します）、トレース、およびその投影時のすべての値
レベル (OPAQUE は何も表示せず、TRANSPARENT は値を表示します)。スタジオは以下にオープンします
http://localhost:4000 。
このリポジトリでは、ゼロセットアップと同じ UI (サンプル) の 2 つの方法があります。
oya.config.ts が含まれています):
make dev # プレイグラウンド (ライブラリをビルドして実行)
bun run build && bunx oyadotai dev # oya.config.ts を提供する CLI - 同じ Studio、パッケージに同梱
独自のプロジェクトで、エージェントをエクスポートする oya.config.ts を追加してから、 bunx oyadotai dev を追加します。
// oya.config.ts
import { Agent , createTool } from "oyadotai" ;
インポート { 人類

[切り捨てられた]

## Original Extract

Agents that plan, don't react. The model writes a typed plan once. The runtime runs it. Tool outputs never go back through the model — so your agent can't be prompt-injected through its tools, costs a fraction, and runs the same every time. - OyaAIProd/oya

GitHub - OyaAIProd/oya: Agents that plan, don't react. The model writes a typed plan once. The runtime runs it. Tool outputs never go back through the model — so your agent can't be prompt-injected through its tools, costs a fraction, and runs the same every time. · GitHub
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
OyaAIProd
/
oya
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
34 Commits 34 Commits .github .github apps/ playground apps/ playground benchmarks benchmarks docs docs packages packages scripts scripts .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md DEMO.md DEMO.md GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md bun.lock bun.lock demo.gif demo.gif demo.tape demo.tape oya.config.ts oya.config.ts package.json package.json tsconfig.base.json tsconfig.base.json View all files Repository files navigation
Your agent re-feeds every tool result back through the model. That's the bug.
Every framework - ReAct, LangGraph, Mastra, the Vercel AI SDK - loops the model on
each step: it reads a tool's output, re-types it as the next tool's arguments, and
pays for the round trip. Oya compiles a typed plan once and executes the DAG.
Values flow by reference, never back through the model.
Same code. Same tools. 10× fewer tokens · 3.5× faster · deterministic · injection-safe by construction.
· TypeScript · Bun · MIT · Drop-in for Mastra
Quickstart · The numbers · Why · Migrate in 2 lines · Studio · Docs · White paper
The open-source core behind oya.ai - the hosted platform for plan-don't-react agents.
import { Agent , createTool } from "oyadotai" ;
import { anthropic } from "oyadotai/anthropic" ;
import { z } from "zod" ;
const getWeather = createTool ( {
id : "get_weather" ,
description : "Look up the weather for a city" ,
inputSchema : z . object ( { city : z . string ( ) } ) ,
execute : async ( { city } ) => fetchWeather ( city ) ,
} ) ;
const agent = new Agent ( {
model : anthropic ( "claude-haiku-4-5-20251001" ) ,
tools : { get_weather : getWeather } ,
} ) ;
const { text } = await agent . generate ( "How's the weather in NYC?" ) ;
That's the whole API - the same shape as Mastra . Types are inferred from your
zod schema, and every value is OPAQUE to the model by default. You did less work
and it's injection-safe.
A token loop re-sends every tool result back through the model on the following
step, spending tokens, latency, and determinism to move state it never needed to
read. oya compiles the plan once and wires each value by reference. The
reconcile task - get a transaction → fetch its record → normalize → validate →
post - measures the difference on the real Anthropic API ( claude-haiku-4-5 ,
identical tools, 8 trials):
oya delivers the same result with 2.3× fewer tokens than the leanest loop, nearly
10× fewer than Mastra, in a third of the wall-clock time - and it does so
identically on every run. The record's bulky payload re-enters a loop's context
at every step; under oya it remains an OPAQUE handle the model never sees.
State fidelity is a guarantee, not an average. Every value flows by reference
as an OPAQUE handle that the model never re-reads or re-emits, so a URL, id, or
amount is delivered to the next tool byte-for-byte and the execution order is the
statically-checked DAG - outcomes a token loop can only approximate, one sampled
run at a time. Reproduce it yourself: bun run bench .
Don't take the table on faith - run it. The benchmark hits the real Anthropic
API with identical tasks and identical tool implementations for all three
frameworks ( benchmarks/tasks.ts ); the only thing that
varies is how much state flows through the model.
# 1. clone and install (needs Bun ≥ 1.1 - https://bun.sh)
git clone https://github.com/OyaAIProd/oya && cd oya
bun install
# 2. build the library - the benchmark imports the built `oyadotai` from dist/
bun run build
# 3. give it your Anthropic key (env var, or drop it in a .env at the repo root)
export ANTHROPIC_API_KEY=sk-ant-...
# 4. run the comparison
bun run bench # default: reconcile, claude-haiku-4-5, 3 trials
bun run bench claude-sonnet-5 # any model id as the first arg
bun run bench --task research # the heavy multi-doc case
Args go in any order: a model id (defaults to claude-haiku-4-5-20251001 ), a
trial count (integer, defaults to 3 ), and --task reconcile / --task payments
/ --task research / --task weather (defaults to reconcile ). Prefer
make bench - it runs the build for you and auto-loads a .env at the repo root,
so steps 2–4 collapse into one command.
The default reconcile task threads a critical token through a multi-hop pipeline
whose fetched record is a bulky payload carrying a look-alike distractor. It
measures token waste (that payload re-enters a loop's context at every step,
while oya keeps it OPAQUE ), state fidelity (a provenance ledger confirms the
token reaches the final tool byte-for-byte), and ordering (a statically-checked
DAG). oya returns 0 corruption and one fixed order on every run, by construction.
--task research extends the comparison to a heavy multi-document workload.
Each framework runs the task N times against the same model, and prints tokens,
latency, and correctness side by side. The methodology is in
benchmarks/README.md .
Every other agent framework is a token loop (ReAct, LangGraph, AutoGen, Mastra,
the Vercel AI SDK): the model picks a tool, sees the raw result, picks the next.
Every URL, ID, and document flows back through the model. Three bugs follow - every
time:
fetched: https://example.io/q3-report.pdf
downloaded: https://example.com/q3-report.pdf ← the model "fixed" the URL
expected: fetch → validate → download
observed: fetch → download → validate (skipped) ← the model reordered the steps
$432 in tokens ← re-reading every result through the model
$51 in tokens ← reading only what it needs to decide
Same root cause: the model read state it never needed to. oya makes that
impossible. The model emits a typed dataflow plan; the runtime executes the DAG;
each value is shown to the model only at the level the plan declares:
An attacker can stuff a payload into any fetched page. The model never reads that
handle , so indirect prompt injection through tool output has nowhere to land. You
annotate none of this - it's OPAQUE by default.
Migrate from Mastra in 2 lines
If your app already uses @mastra/core , oya is close to a drop-in. createTool
and Agent mirror the shapes you already write, so for most apps the entire
migration is the import lines.
- import { Agent } from "@mastra/core/agent";
- import { createTool } from "@mastra/core/tools";
- import { anthropic } from "@ai-sdk/anthropic";
+ import { Agent, createTool } from "oyadotai";
+ import { anthropic } from "oyadotai/anthropic";
Step 2 - there is no step 2. Your tool and agent definitions compile unchanged.
Mastra
oya
Notes
createTool({ id, description, inputSchema, execute })
same
Types are inferred from the zod inputSchema ; execute 's argument is typed for you - no casts.
new Agent({ name, instructions, model, tools })
same
tools is the same name → tool map.
await agent.generate(prompt) → { text }
same
text is the headline field, and token usage is reported for parity.
agent.stream(prompt)
same call
Returns structured events ( fullStream / textStream ) instead of a raw token soup.
anthropic(…) · openai(…) · google(…)
oyadotai/anthropic · …/openai · …/google
Same call shape, different import path.
So the code you already have keeps working as-is:
import { Agent , createTool } from "oyadotai" ;
import { anthropic } from "oyadotai/anthropic" ;
import { z } from "zod" ;
const getWeather = createTool ( {
id : "get_weather" ,
description : "Look up the weather for a city" ,
inputSchema : z . object ( { city : z . string ( ) } ) ,
execute : async ( { city } ) => fetchWeather ( city ) ,
} ) ;
const agent = new Agent ( {
name : "WeatherBot" ,
instructions : "Answer weather questions." ,
model : anthropic ( "claude-haiku-4-5-20251001" ) ,
tools : { get_weather : getWeather } ,
} ) ;
const { text } = await agent . generate ( "How's the weather in NYC?" ) ;
What changes underneath
Same surface, different engine. Mastra runs a token loop : the model calls a
tool, reads the raw result, and decides the next call - re-reading every value back
through the model. oya has the model emit one typed plan and executes that DAG
directly. Values flow between tools by reference and are disclosed to the model
only at their declared projection level ( OPAQUE by default).
You didn't touch your code, but the migrated app now gets, for free:
~10× fewer tokens and ~3× faster on multi-hop tasks (the numbers above ) - intermediate state is piped tool-to-tool, never re-billed through the model.
Deterministic execution - one statically-checked order on every run, instead of a model-chosen sequence that can reorder or skip steps.
Injection-safe by construction - tool output the model never reads can't smuggle an injected instruction into its context.
Good to know before you migrate
Projection defaults to OPAQUE . Tool outputs are hidden from the model
unless you mark them SUMMARY or TRANSPARENT . That's the safety win - but if a
tool relied on the model reading its raw output to decide the next step, declare
that output's projection. See Projection Types .
Providers import from oyadotai/* ( oyadotai/anthropic , oyadotai/openai ,
oyadotai/google ) rather than @ai-sdk/* .
oya is pre-1.0 and mirrors the core of the Mastra surface, not every helper.
If something you depend on is missing, open an issue - Mastra parity gaps are high-priority.
Chat with your agents and watch each plan execute live - the DAG (React Flow, nodes
colored by kind and lit as they run), the trace, and every value at its projection
level ( OPAQUE shows nothing, TRANSPARENT shows the value). Studio opens at
http://localhost:4000 .
In this repo - two ways, both zero-setup and both the same UI (a sample
oya.config.ts is included):
make dev # the playground (builds libs, then runs it)
bun run build && bunx oyadotai dev # the CLI, serving oya.config.ts - same Studio, shipped in the package
In your own project - add an oya.config.ts that exports your agents, then bunx oyadotai dev :
// oya.config.ts
import { Agent , createTool } from "oyadotai" ;
import { anthropic

[truncated]
