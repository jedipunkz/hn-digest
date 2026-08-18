---
source: "https://strandsagents.com/"
hn_url: "https://news.ycombinator.com/item?id=49352099"
title: "Has any one ever used Strands – the AI Agent SDK by AWS"
article_title: "Strands Agents — Open Source AI Agent SDK for Python & TypeScript"
image: "https://strandsagents.com/og-image.png"
author: "donbox"
captured_at: "2026-08-18T21:14:48Z"
capture_tool: "hn-digest"
hn_id: 49352099
score: 2
comments: 0
posted_at: "2026-08-18T20:20:40Z"
tags:
  - hacker-news
  - translated
---

# Has any one ever used Strands – the AI Agent SDK by AWS

- HN: [49352099](https://news.ycombinator.com/item?id=49352099)
- Source: [strandsagents.com](https://strandsagents.com/)
- Score: 2
- Comments: 0
- Posted: 2026-08-18T20:20:40Z

## Translation

タイトル: Strands – AWS の AI エージェント SDK を使用したことがありますか?
記事のタイトル: Strands Agents — Python および TypeScript 用のオープンソース AI エージェント SDK
説明: 数行のコードで本番環境に対応した AI エージェントを構築します。モデル駆動型のオープンソースで、Amazon Bedrock、Anthropic、OpenAI などと連携します。

記事本文:
検索 Ctrl K キャンセル ホーム 言語 Python TypeScript ホーム ドキュメント 例 統合 API リファレンス ブログ 変更履歴 Labs ↗ GitHub PY Python SDK ↗ TS TypeScript SDK ↗ Strands EV Evals ↗ SH Shell ↗ 組織 structs-agents ↗ structs-labs ↗ Discord ↗ Python TypeScript Strands Harness PY Python SDK ↗ TS TypeScript SDK ↗ Strands EV Evals ↗ SH Shell ↗ 組織 structs-agents ↗ structs-labs ↗ テーマの選択 ダーク ライト オート ホーム ドキュメント 例 統合 API リファレンス ブログ Changelog Labs ↗ 運用エージェントを構築するためのオープンソース ツールキット。
npm install @structs-agents/sdk GitHub スター 6,900 個以上 · Amazon 内の本番システムから構築
ストランドからのインポート エージェント、ツール
structs.hooks からインポート BeforeToolCallEvent
pathlibインポートパスから
@ツール
def save_report (タイトル: str 、内容: str ) -> str :
"""調査レポートをディスクに保存します。"""
path = f "レポート/ { タイトル } .md"
パス(パス).write_text(コンテンツ)
return f "保存された { path } "
def require_sources (イベント: BeforeToolCallEvent):
名前 = イベント.ツール_使用[ "名前" ]
inp = str (event.tool_use[ "input" ])
name == "save_report" および "[source]" が inp にない場合:
events.cancel_tool = "出典の引用を追加します。"
エージェント = エージェント(
ツール = [レポートの保存]、
フック = [require_sources],
)
Agent( "Research AI Agent Frameworks" ) ストランドからインポート エージェント、ツール
structs.hooks からインポート BeforeToolCallEvent
pathlibインポートパスから
@ツール
def save_report (タイトル: str 、内容: str ) -> str :
"""調査レポートをディスクに保存します。"""
path = f "レポート/ { タイトル } .md"
パス(パス).write_text(コンテンツ)
return f "保存された { path } "
def require_sources (イベント: BeforeToolCallEvent):
名前 = イベント.ツール_使用[ "名前" ]
inp = str (event.tool_use[ "input" ])
name == "save_report" および "[source]" が inp にない場合:
event.cancel_tool = "追加します

ウルスの引用。」
エージェント = エージェント(
ツール = [レポートの保存]、
フック = [require_sources],
)
Agent( "リサーチ AI エージェント フレームワーク" ) {
writeFileSync(`reports/${title}.md`, content)
`保存された ${title}.md` を返します
}、
})
const エージェント = 新しいエージェント({ ツール: [saveReport] })
Agent.addHook(BeforeToolCallEvent, (イベント) => {
const inp = String(event.toolUse.input)
if (event.toolUse.name === 'save_report') {
if (!inp.includes('[ソース]')) {
event.cancel = '出典引用を追加します。'
}
}
})
await Agent.invoke('Research AI エージェント フレームワーク')"> Research_agent.ts import {
エージェント、ツール、BeforeToolCallEvent
'@structs-agents/sdk' から
「zod」から z をインポート
「fs」からインポート { writeFileSync }
const saveReport = ツール ({
名前: 'save_report' 、
説明: 「調査レポートを保存します。」 、
入力スキーマ: z.オブジェクト ({
タイトル：Z.文字列()、
内容：z。文字列()、
})、
コールバック: ({ タイトル , コンテンツ }) => {
writeFileSync ( `reports/${
[切り捨てられた]
クロード mcp ストランドを追加 uvx ストランド-エージェント-mcp-server クロード mcp ストランドを追加 uvx ストランド-エージェント-mcp-server ~/.cursor/mcp.json に追加します。
{
"mcpサーバー": {
"ストランドエージェント" : {
"コマンド" : "uvx" ,
"args" : [ "ストランドエージェント-mcp-サーバー" ]
}
}
} {
"mcpサーバー": {
"ストランドエージェント" : {
"コマンド" : "uvx" ,
"args" : [ "ストランドエージェント-mcp-サーバー" ]
}
}
~/.kiro/settings/mcp.json に追加します。
{
"mcpサーバー": {
"ストランドエージェント" : {
"コマンド" : "uvx" ,
"args" : [ "ストランドエージェント-mcp-サーバー" ],
"無効" : false 、
"autoApprove" : [ "search_docs" 、 "fetch_doc" ]
}
}
} {
"mcpサーバー": {
"ストランドエージェント" : {
"コマンド" : "uvx" ,
"args" : [ "ストランドエージェント-mcp-サーバー" ],
"無効" : false 、
"autoApprove" : [ "search_docs" 、 "fetch_doc" ]
}
}
mcp.json に以下を追加します。
{
「サーバー」: {
"ストランドエージェント" : {
"コマンド" : "uvx" ,
"args" : [ "ストランドエージェント-mcp-サーバー" ]
}
}
} {
「サーバー」: {
"ストランドエージェント" :

{
"コマンド" : "uvx" ,
"args" : [ "ストランドエージェント-mcp-サーバー" ]
}
}
または、次のプロンプトを貼り付けます。
llms.txt llms-full.txt AI を使用したビルド ガイド
あらゆるモデル、あらゆるクラウド。構成行を記述する前に、コンテキスト管理、実行制限、可観測性を取得できます。スケールするときにバックエンドを交換します。コードは変わりません。
ストランドからのインポート エージェント、ツール
@ツール
def search_logs (クエリ: str 、時間: int = 24 ) -> リスト:
"""アプリケーションログをキーワードで検索します。"""
return log_api.search(クエリ、時間)
エージェント = エージェント(
ツール = [検索ログ]、
)
Agent( "過去 6 時間のすべてのタイムアウト エラーを検索する" ) ストランドからのインポート エージェント、ツール
@ツール
def search_logs (クエリ: str 、時間: int = 24 ) -> リスト:
"""アプリケーションログをキーワードで検索します。"""
return log_api.search(クエリ、時間)
エージェント = エージェント(
ツール = [検索ログ]、
)
Agent( "過去 6 時間のすべてのタイムアウト エラーを検索" )
logApi.search(クエリ, 時間),
})
const エージェント = 新しいエージェント({ ツール: [searchLogs] })
エージェントの呼び出しを待ちます(
「過去 6 時間のすべてのタイムアウト エラーを検索」
)">agent.ts import { エージェント、ツール } from '@structs-agents/sdk'
「zod」から z をインポート
const searchLogs = ツール ({
名前: 'search_logs' 、
説明: 「キーワードでログを検索します。」 、
入力スキーマ: z.オブジェクト ({
クエリ: z.文字列()、
時間：z.番号（）。デフォルト ( 24 )、
})、
コールバック: ({ クエリ , 時間 }) =>
ログAPI。検索 (クエリ、時間)、
})
const エージェント = 新しいエージェント ({ tools: [searchLogs] })
エージェントを待ちます。呼び出す(
「過去 6 時間のすべてのタイムアウト エラーを検索」
) import { エージェント、ツール } から '@structs-agents/sdk'
「zod」から z をインポート
const searchLogs = ツール ({
名前: 'search_logs' 、
説明: 「キーワードでログを検索します。」 、
入力スキーマ: z.オブジェクト ({
クエリ: z.文字列()、
時間：z.番号（）。デフォルト ( 24 )、
})、
コールバック: ({ クエリ , 時間 }) =>
ログAPI。検索 (クエリ、時間)、
})
const エージェント = 新しいエージェント ({ ツール

s: [検索ログ] })
エージェントを待ちます。呼び出す(
「過去 6 時間のすべてのタイムアウト エラーを検索」
) 漸進的な複雑さ。ロックインゼロ。
stronds.agent インポートから SummarizingConversationManager
# 同じエージェントですが、要約が表示されます。
エージェント = エージェント(
ツール = [検索ログ]、
communication_manager = SummarizingConversationManager(),
) structs.agent import から SummarizingConversationManager
# 同じエージェントですが、要約が表示されます。
エージェント = エージェント(
ツール = [検索ログ]、
communication_manager = SummarizingConversationManager(),
)agent_with_context.tsインポート{
要約会話マネージャー、
'@structs-agents/sdk' から
// 同じエージェントですが、要約が表示されます。
const エージェント = 新しいエージェント ({
ツール: [searchLogs]、
会話マネージャー:
新しい SummarizingConversationManager ()、
})インポート{
要約会話マネージャー、
'@structs-agents/sdk' から
// 同じエージェントですが、要約が表示されます。
const エージェント = 新しいエージェント ({
ツール: [searchLogs]、
会話マネージャー:
新しい SummarizingConversationManager ()、
}) コントロールを維持する
フックを使用して監視、変更、デバッグします。エージェント ループは、デフォルトですべての決定を追跡します。フックを使用すると、ステップをインターセプトしてログに記録したり、検証したり、リダイレクトしたりできます。
ストランドからのインポートエージェント
ストランド.フックからインポートAfterToolCallEvent
def log_tool_calls (イベント: AfterToolCallEvent):
"""すべてのツール呼び出しをログに記録します。"""
print ( f "ツール: {event.tool_use[ 'name' ] } ")
print ( f "結果: {event.result[ 'status' ] } " )
エージェント = エージェント(
tools = [検索ログ、クエリデータベース]、
フック = [log_tool_calls],
トレース属性 = {
"サービス" : "ops エージェント" ,
"環境" : "本番環境" 、
}、
) ストランドインポートエージェントから
ストランド.フックからインポートAfterToolCallEvent
def log_tool_calls (イベント: AfterToolCallEvent):
"""すべてのツール呼び出しをログに記録します。"""
print ( f "ツール: {event.tool_use[ 'name' ] } ")
print ( f "結果: {event.result[ 'status' ] } "

)
エージェント = エージェント(
tools = [検索ログ、クエリデータベース]、
フック = [log_tool_calls],
トレース属性 = {
"サービス" : "ops エージェント" ,
"環境" : "本番環境" 、
}、
) {
console.log(`ツール: ${event.toolUse.name}`)
console.log(`ステータス: ${event.result.status}`)
})"> observable_agent.ts import {
エージェント、AfterToolCallEvent、
'@structs-agents/sdk' から
const エージェント = 新しいエージェント ({
ツール: [searchLogs、queryDatabase]、
トレース属性: {
サービス: 'ops-agent' 、
環境: '本番環境' 、
}、
})
エージェント。 addHook (AfterToolCallEvent, ( イベント ) => {
コンソール。ログ ( `ツール: ${ イベント . ツール使用 . 名前 }` )
コンソール。 log (`ステータス: ${ イベント . 結果 . ステータス }` )
})インポート{
エージェント、AfterToolCallEvent、
'@structs-agents/sdk' から
const エージェント = 新しいエージェント ({
ツール: [searchLogs、queryDatabase]、
トレース属性: {
サービス: 'ops-agent' 、
環境: '本番環境' 、
}、
})
エージェント。 addHook (AfterToolCallEvent, ( イベント ) => {
コンソール。ログ ( `ツール: ${ イベント . ツール使用 . 名前 }` )
コンソール。 log (`ステータス: ${ イベント . 結果 . ステータス }` )
}) 組み込みの可観測性。
ガードレールは走る前にミスをキャッチします。
ストランドからのインポートエージェント
structs.hooks からインポート BeforeToolCallEvent
WRITE_OPS = [ "INSERT" 、 "UPDATE" 、 "DELETE" 、 "DROP" ]
def read_only_guard (イベント: BeforeToolCallEvent):
"""ブロック書き込み。このエージェントは読み取り専用です。"""
if events.tool_use[ "name" ] == "query_database" :
sql = events.tool_use[ "input" ].get( "クエリ" , "" )
存在する場合 ( WRITE_OPS の kw に対する sql.upper() の kw):
event.cancel_tool = "読み取り専用アクセス。"
エージェント = エージェント(
ツール = [クエリデータベース]、
フック = [読み取り専用ガード]、
) ストランドインポートエージェントから
structs.hooks からインポート BeforeToolCallEvent
WRITE_OPS = [ "INSERT" 、 "UPDATE" 、 "DELETE" 、 "DROP" ]
def read_only_guard (イベント: BeforeToolCallEvent):
"""ブロック書き込み。このエージェントは読み取り専用です。"""
ifevent.tool_use[ "名前" ] == "qu

ery_データベース" :
sql = events.tool_use[ "input" ].get( "クエリ" , "" )
存在する場合 ( WRITE_OPS の kw に対する sql.upper() の kw):
event.cancel_tool = "読み取り専用アクセス。"
エージェント = エージェント(
ツール = [クエリデータベース]、
フック = [読み取り専用ガード]、
) その後、ハーネスは「WHERE 句を追加する」、「最初に権限を確認する」という具体的なフィードバックを返します。エージェントは自動的に修正します。すべてのステップを細かく管理しなくても、信頼できる結果が得られます。
stronds.vended_plugins.steering インポートから (
SteeringHandler、ガイド、進行、
)
クラス QueryQualityPolicy ( SteeringHandler ):
async def steer_before_tool (
self、*、agent、tool_use、** kwargs
):
sql = tools_use[ "入力" ].get( "クエリ" , "" ).upper()
SQL に「SELECT」があり、SQL に「WHERE」がない場合:
ガイドを返す(
reason = "WHERE 句と LIMIT を追加します。"
)
sql.upper().count( "JOIN" ) > 3 の場合:
ガイドを返す(
reason = "4 つ以上の結合。より小さなクエリに分割します。"
)
return Proceed(reason = "クエリは良さそうです。" )
エージェント = エージェント(
ツール = [クエリデータベース]、
プラグイン = [QueryQualityPolicy()]、
) stronds.vended_plugins.steering インポートから (
SteeringHandler、ガイド、進行、
)
クラス QueryQualityPolicy ( SteeringHandler ):
async def steer_before_tool (
self、*、agent、tool_use、** kwargs
):
sql = tools_use[ "入力" ].get( "クエリ" , "" ).upper()
SQL に「SELECT」があり、SQL に「WHERE」がない場合:
ガイドを返す(
reason = "WHERE 句と LIMIT を追加します。"
)
sql.upper().count( "JOIN" ) > 3 の場合:
ガイドを返す(
reason = "4 つ以上の結合。より小さなクエリに分割します。"
)
return Proceed(reason = "クエリは良さそうです。" )
エージェント = エージェント(
ツール = [クエリデータベース]、
プラグイン = [QueryQualityPolicy()]、
) ステアリングによるエージェントの精度 100% プロンプトのみのエージェントのスコアは 82.5% 。ハードコーディングされたワークフローのスコアは 80.8% でした。 Strands ステアリング ハンドラーを備えたエージェントは、あらゆる間違いから立ち直りました。ベンチマークを参照 → Smartsheet では、次の世代に Strands を選択しました

AI 機能は、エンタープライズ対応の機能と開発効率の完璧なバランスを提供するため、AI 機能の導入に貢献しました。その堅牢な会話メモリと動的なツール登録システムは、応答性が高く、コンテキストを認識するインテリジェントな AI アシスタントを作成するために不可欠でした。 Strands を使用することで、安全でスケーラブルなソリューションを迅速に実装することができ、安全で高性能なエンタープライズ グレードの AI エクスペリエンスを提供するための本番環境に対応した基盤が得られました。
Amazon Bedrock、Amazon OpenSearch を備えた RAG、Strands SDK を備えたマルチエージェント オーケストレーション、および Kiro AI IDE を使用して、従来のエラー アラートをインテリジェントなインシデント対応に変換し、手動コーディングなしで MTTR を 60% 削減します。
Strands の SDK と AWS ネイティブ サービスとの優れた統合により、Landchecker のエージェント開発が合理化されました。 AgentCore Runtime、Bedrock Guardrails、OpenTelemetry の組み込みサポートの統合が容易になったことで、私たちは最も得意とすること、つまりプロパティ情報ツールの開発とデータ統合に集中できるようになりました。
Swisscom では、エンタープライズ対応かつ将来性のあるエージェント AI バックボーンを必要としています。 Strands Agents は、クラウド環境にネイティブに適合すると同時に、完全にオープンソースで柔軟性があるという、両方の長所を提供します。その組み合わせはすべて

[切り捨てられた]

## Original Extract

Build production-ready AI agents in a few lines of code. Model-driven, open source, works with Amazon Bedrock, Anthropic, OpenAI, and more.

Search Ctrl K Cancel Home Language Python TypeScript Home Docs Examples Integrations API Reference Blog Changelog Labs ↗ GitHub PY Python SDK ↗ TS TypeScript SDK ↗ Strands EV Evals ↗ SH Shell ↗ Organizations strands-agents ↗ strands-labs ↗ Discord ↗ Python TypeScript Strands Harness PY Python SDK ↗ TS TypeScript SDK ↗ Strands EV Evals ↗ SH Shell ↗ Organizations strands-agents ↗ strands-labs ↗ Select theme Dark Light Auto Home Docs Examples Integrations API Reference Blog Changelog Labs ↗ The open source toolkit for building production agents.
npm install @strands-agents/sdk 6,900+ GitHub stars · Built from production systems inside Amazon
from strands import Agent, tool
from strands.hooks import BeforeToolCallEvent
from pathlib import Path
@tool
def save_report (title: str , content: str ) -> str :
"""Save a research report to disk."""
path = f "reports/ { title } .md"
Path(path).write_text(content)
return f "Saved { path } "
def require_sources (event: BeforeToolCallEvent):
name = event.tool_use[ "name" ]
inp = str (event.tool_use[ "input" ])
if name == "save_report" and "[source]" not in inp:
event.cancel_tool = "Add source citations."
agent = Agent(
tools = [save_report],
hooks = [require_sources],
)
agent( "Research AI agent frameworks" ) from strands import Agent, tool
from strands.hooks import BeforeToolCallEvent
from pathlib import Path
@tool
def save_report (title: str , content: str ) -> str :
"""Save a research report to disk."""
path = f "reports/ { title } .md"
Path(path).write_text(content)
return f "Saved { path } "
def require_sources (event: BeforeToolCallEvent):
name = event.tool_use[ "name" ]
inp = str (event.tool_use[ "input" ])
if name == "save_report" and "[source]" not in inp:
event.cancel_tool = "Add source citations."
agent = Agent(
tools = [save_report],
hooks = [require_sources],
)
agent( "Research AI agent frameworks" ) {
writeFileSync(`reports/${title}.md`, content)
return `Saved ${title}.md`
},
})
const agent = new Agent({ tools: [saveReport] })
agent.addHook(BeforeToolCallEvent, (event) => {
const inp = String(event.toolUse.input)
if (event.toolUse.name === 'save_report') {
if (!inp.includes('[source]')) {
event.cancel = 'Add source citations.'
}
}
})
await agent.invoke('Research AI agent frameworks')"> research_agent.ts import {
Agent, tool, BeforeToolCallEvent
} from '@strands-agents/sdk'
import z from 'zod'
import { writeFileSync } from 'fs'
const saveReport = tool ({
name: 'save_report' ,
description: 'Save a research report.' ,
inputSchema: z. object ({
title: z. string (),
content: z. string (),
}),
callback : ({ title , content }) => {
writeFileSync ( `reports/${
[truncated]
claude mcp add strands uvx strands-agents-mcp-server claude mcp add strands uvx strands-agents-mcp-server Add to ~/.cursor/mcp.json:
{
"mcpServers" : {
"strands-agents" : {
"command" : "uvx" ,
"args" : [ "strands-agents-mcp-server" ]
}
}
} {
"mcpServers" : {
"strands-agents" : {
"command" : "uvx" ,
"args" : [ "strands-agents-mcp-server" ]
}
}
} Add to ~/.kiro/settings/mcp.json:
{
"mcpServers" : {
"strands-agents" : {
"command" : "uvx" ,
"args" : [ "strands-agents-mcp-server" ],
"disabled" : false ,
"autoApprove" : [ "search_docs" , "fetch_doc" ]
}
}
} {
"mcpServers" : {
"strands-agents" : {
"command" : "uvx" ,
"args" : [ "strands-agents-mcp-server" ],
"disabled" : false ,
"autoApprove" : [ "search_docs" , "fetch_doc" ]
}
}
} Add to your mcp.json:
{
"servers" : {
"strands-agents" : {
"command" : "uvx" ,
"args" : [ "strands-agents-mcp-server" ]
}
}
} {
"servers" : {
"strands-agents" : {
"command" : "uvx" ,
"args" : [ "strands-agents-mcp-server" ]
}
}
} Or paste this prompt:
llms.txt llms-full.txt Build with AI guide
Any model, any cloud. You get context management, execution limits, and observability before you write a line of config. Swap backends when you scale. Your code stays the same.
from strands import Agent, tool
@tool
def search_logs (query: str , hours: int = 24 ) -> list :
"""Search application logs by keyword."""
return log_api.search(query, hours)
agent = Agent(
tools = [search_logs],
)
agent( "Find all timeout errors from the last 6 hours" ) from strands import Agent, tool
@tool
def search_logs (query: str , hours: int = 24 ) -> list :
"""Search application logs by keyword."""
return log_api.search(query, hours)
agent = Agent(
tools = [search_logs],
)
agent( "Find all timeout errors from the last 6 hours" )
logApi.search(query, hours),
})
const agent = new Agent({ tools: [searchLogs] })
await agent.invoke(
'Find all timeout errors from the last 6 hours'
)"> agent.ts import { Agent, tool } from '@strands-agents/sdk'
import z from 'zod'
const searchLogs = tool ({
name: 'search_logs' ,
description: 'Search logs by keyword.' ,
inputSchema: z. object ({
query: z. string (),
hours: z. number (). default ( 24 ),
}),
callback : ({ query , hours }) =>
logApi. search (query, hours),
})
const agent = new Agent ({ tools: [searchLogs] })
await agent. invoke (
'Find all timeout errors from the last 6 hours'
) import { Agent, tool } from '@strands-agents/sdk'
import z from 'zod'
const searchLogs = tool ({
name: 'search_logs' ,
description: 'Search logs by keyword.' ,
inputSchema: z. object ({
query: z. string (),
hours: z. number (). default ( 24 ),
}),
callback : ({ query , hours }) =>
logApi. search (query, hours),
})
const agent = new Agent ({ tools: [searchLogs] })
await agent. invoke (
'Find all timeout errors from the last 6 hours'
) Progressive complexity. Zero lock-in.
from strands.agent import SummarizingConversationManager
# Same agent, now with summarization.
agent = Agent(
tools = [search_logs],
conversation_manager = SummarizingConversationManager(),
) from strands.agent import SummarizingConversationManager
# Same agent, now with summarization.
agent = Agent(
tools = [search_logs],
conversation_manager = SummarizingConversationManager(),
) agent_with_context.ts import {
SummarizingConversationManager,
} from '@strands-agents/sdk'
// Same agent, now with summarization.
const agent = new Agent ({
tools: [searchLogs],
conversationManager:
new SummarizingConversationManager (),
}) import {
SummarizingConversationManager,
} from '@strands-agents/sdk'
// Same agent, now with summarization.
const agent = new Agent ({
tools: [searchLogs],
conversationManager:
new SummarizingConversationManager (),
}) Stay in control
Monitor, modify, and debug with hooks. The agent loop traces every decision by default. Hooks let you intercept any step to log it, validate it, or redirect it.
from strands import Agent
from strands.hooks import AfterToolCallEvent
def log_tool_calls (event: AfterToolCallEvent):
"""Log every tool call."""
print ( f "Tool: { event.tool_use[ 'name' ] } " )
print ( f "Result: { event.result[ 'status' ] } " )
agent = Agent(
tools = [search_logs, query_database],
hooks = [log_tool_calls],
trace_attributes = {
"service" : "ops-agent" ,
"env" : "production" ,
},
) from strands import Agent
from strands.hooks import AfterToolCallEvent
def log_tool_calls (event: AfterToolCallEvent):
"""Log every tool call."""
print ( f "Tool: { event.tool_use[ 'name' ] } " )
print ( f "Result: { event.result[ 'status' ] } " )
agent = Agent(
tools = [search_logs, query_database],
hooks = [log_tool_calls],
trace_attributes = {
"service" : "ops-agent" ,
"env" : "production" ,
},
) {
console.log(`Tool: ${event.toolUse.name}`)
console.log(`Status: ${event.result.status}`)
})"> observable_agent.ts import {
Agent, AfterToolCallEvent,
} from '@strands-agents/sdk'
const agent = new Agent ({
tools: [searchLogs, queryDatabase],
traceAttributes: {
service: 'ops-agent' ,
env: 'production' ,
},
})
agent. addHook (AfterToolCallEvent, ( event ) => {
console. log ( `Tool: ${ event . toolUse . name }` )
console. log ( `Status: ${ event . result . status }` )
}) import {
Agent, AfterToolCallEvent,
} from '@strands-agents/sdk'
const agent = new Agent ({
tools: [searchLogs, queryDatabase],
traceAttributes: {
service: 'ops-agent' ,
env: 'production' ,
},
})
agent. addHook (AfterToolCallEvent, ( event ) => {
console. log ( `Tool: ${ event . toolUse . name }` )
console. log ( `Status: ${ event . result . status }` )
}) Built-in observability.
Guardrails catch mistakes before they run.
from strands import Agent
from strands.hooks import BeforeToolCallEvent
WRITE_OPS = [ "INSERT" , "UPDATE" , "DELETE" , "DROP" ]
def read_only_guard (event: BeforeToolCallEvent):
"""Block writes. This agent is read-only."""
if event.tool_use[ "name" ] == "query_database" :
sql = event.tool_use[ "input" ].get( "query" , "" )
if any (kw in sql.upper() for kw in WRITE_OPS ):
event.cancel_tool = "Read-only access."
agent = Agent(
tools = [query_database],
hooks = [read_only_guard],
) from strands import Agent
from strands.hooks import BeforeToolCallEvent
WRITE_OPS = [ "INSERT" , "UPDATE" , "DELETE" , "DROP" ]
def read_only_guard (event: BeforeToolCallEvent):
"""Block writes. This agent is read-only."""
if event.tool_use[ "name" ] == "query_database" :
sql = event.tool_use[ "input" ].get( "query" , "" )
if any (kw in sql.upper() for kw in WRITE_OPS ):
event.cancel_tool = "Read-only access."
agent = Agent(
tools = [query_database],
hooks = [read_only_guard],
) Then the harness gives specific feedback: "add a WHERE clause," "check permissions first." The agent corrects itself. You get reliable outcomes without micromanaging every step.
from strands.vended_plugins.steering import (
SteeringHandler, Guide, Proceed,
)
class QueryQualityPolicy ( SteeringHandler ):
async def steer_before_tool (
self, * , agent, tool_use, ** kwargs
):
sql = tool_use[ "input" ].get( "query" , "" ).upper()
if "SELECT" in sql and "WHERE" not in sql:
return Guide(
reason = "Add a WHERE clause and LIMIT."
)
if sql.upper().count( "JOIN" ) > 3 :
return Guide(
reason = "4+ joins. Break into smaller queries."
)
return Proceed( reason = "Query looks good." )
agent = Agent(
tools = [query_database],
plugins = [QueryQualityPolicy()],
) from strands.vended_plugins.steering import (
SteeringHandler, Guide, Proceed,
)
class QueryQualityPolicy ( SteeringHandler ):
async def steer_before_tool (
self, * , agent, tool_use, ** kwargs
):
sql = tool_use[ "input" ].get( "query" , "" ).upper()
if "SELECT" in sql and "WHERE" not in sql:
return Guide(
reason = "Add a WHERE clause and LIMIT."
)
if sql.upper().count( "JOIN" ) > 3 :
return Guide(
reason = "4+ joins. Break into smaller queries."
)
return Proceed( reason = "Query looks good." )
agent = Agent(
tools = [query_database],
plugins = [QueryQualityPolicy()],
) 100% agent accuracy with steering Prompt-only agents scored 82.5% . Hard-coded workflows scored 80.8% . Agents with Strands steering handlers recovered from every mistake. See the benchmark → At Smartsheet, we chose Strands for our next generation of AI capabilities because it provided the perfect balance of enterprise-ready features and development efficiency. Its robust conversation memory and dynamic tool registration systems were crucial for creating a responsive, context-aware intelligent AI assistant. With Strands, we were able to quickly implement a secure and scalable solution, giving us a production-ready foundation to deliver a secure, high-performance, and enterprise-grade AI experience.
Transform traditional error alerts into intelligent incident responses using Amazon Bedrock, RAG with Amazon OpenSearch, Multi-Agent Orchestration with Strands SDK, and Kiro AI IDE - reducing MTTR by 60% without manual coding.
Strands’ SDK and great integration with AWS native services streamlined Landchecker’s development of agents. With easier integration of AgentCore Runtime, Bedrock Guardrails, and built-in support for OpenTelemetry, we could focus on what we do best – developing property information tools and data integrations.
At Swisscom, we need an agentic AI backbone that is both enterprise-ready and future-proof. Strands Agents gives us the best of both worlds: a native fit with our cloud environment, yet fully open source and flexible. That combination allo

[truncated]
