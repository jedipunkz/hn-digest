---
source: "https://www.nuget.org/packages/Toolnexus/"
hn_url: "https://news.ycombinator.com/item?id=48721207"
title: "Show HN: Toolnexus for .NET – MCP and agent skills as tools for any LLM"
article_title: "NuGet Gallery\n| Toolnexus 0.1.2"
author: "muthuishere"
captured_at: "2026-06-29T16:59:40Z"
capture_tool: "hn-digest"
hn_id: 48721207
score: 1
comments: 0
posted_at: "2026-06-29T16:18:48Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Toolnexus for .NET – MCP and agent skills as tools for any LLM

- HN: [48721207](https://news.ycombinator.com/item?id=48721207)
- Source: [www.nuget.org](https://www.nuget.org/packages/Toolnexus/)
- Score: 1
- Comments: 0
- Posted: 2026-06-29T16:18:48Z

## Translation

タイトル: HN の表示: Toolnexus for .NET – LLM のツールとしての MCP およびエージェントのスキル
記事のタイトル: NuGet ギャラリー
|ツールネクサス0.1.2
説明: OpenAI/Anthropic/Gemini アダプターと統合ホスト ループを使用して、動的 MCP サーバー + エージェント スキル + ネイティブ + HTTP ツールを、LLM の 1 つのツール セットとして数行でエージェントを構築します。 Toolnexus の C# ポート。

記事本文:
NuGet ギャラリー
|ツールネクサス0.1.2
コンテンツにスキップ
ナビゲーションを切り替え
パッケージ
dotnet add package Toolnexus --version 0.1.2
コピー
NuGet\Install-Package Toolnexus - バージョン 0.1.2
コピー
このコマンドは、NuGet モジュールのバージョンの Install-Package を使用するため、Visual Studio のパッケージ マネージャー コンソール内で使用することを目的としています。
<PackageReference Include="Toolnexus" Version="0.1.2" />
コピー
PackageReference をサポートするプロジェクトの場合は、この XML ノードをプロジェクト ファイルにコピーして、パッケージを参照します。
<PackageVersion Include="Toolnexus" Version="0.1.2" />
ディレクトリ.パッケージ.props
コピー
<PackageReference Include="Toolnexus" />
プロジェクトファイル
コピー
Central Package Management (CPM) をサポートするプロジェクトの場合は、この XML ノードをソリューションの Directory.Packages.props ファイルにコピーして、パッケージのバージョンを設定します。
パケット追加 Toolnexus --バージョン 0.1.2
コピー
NuGet チームは、このクライアントのサポートを提供しません。サポートについては、メンテナに問い合わせてください。
#r "nuget: Toolnexus、0.1.2"
コピー
#r ディレクティブは、F# Interactive および Polyglot Notebook で使用できます。これを対話型ツールまたはスクリプトのソース コードにコピーして、パッケージを参照します。
#:パッケージ Toolnexus@0.1.2
コピー
#:package ディレクティブは、.NET 10 プレビュー 4 以降の C# ファイルベースのアプリで使用できます。これを .cs ファイルの、パッケージを参照するコード行の前にコピーします。
#addin nuget:?package=Toolnexus&version=0.1.2
Cake アドインとしてインストールする
コピー
#tool nuget:?package=Toolnexus&version=0.1.2
Cakeツールとしてインストールする
コピー
NuGet チームは、このクライアントのサポートを提供しません。サポートについては、メンテナに問い合わせてください。
お読みください
数行でエージェントを構築します。 mcp.json と skill/ フォルダーを指定し、 RunAsync() を呼び出します。
動作するエージェント (MCP サーバー、エージェントのスキル、独自の機能、HTTP エンドポイント) が存在します。
一つに統一される

OL セット、LLM を駆動します。
ちょうどいいサイズ。フレームワークではありません (ビルダーや設定を進める必要はありません)、落ちるおもちゃではありません
ストリーミングまたは再試行が必要な瞬間まで。実際のエージェントに必要なものすべて - ループ、フック、
ストリーミング、再試行、メモリ、そしてそうでないものは何もありません。
Toolnexus の C#/.NET ポート — 同じライブラリ、
JavaScript、Python、Go、Java でもバイト同一。公式に基づいて構築
ModelContextProtocol SDK。 .NET 10をターゲットとします。
dotnet パッケージ Toolnexus を追加
エージェントは 3 つのステップで完了
Toolnexus を使用する。
// 1. mcp.json + skill/ フォルダーのツール
var tk = await Toolkit.CreateAsync(new Toolkit.Options() を使用して待機します
.WithMcpConfig("./mcp.json")
.WithSkillsDir("./スキル"));
// 2. OpenAI または Anthropic スタイルのエンドポイントをポイントします
var エージェント = LlmClient.Create(new LlmClient.Options
{
BaseUrl = "https://openrouter.ai/api/v1",
Style = "openai", // または "anthropic"
モデル = "openai/gpt-4o-mini"、
ApiKey = 環境.GetEnvironmentVariable("OPENROUTER_API_KEY"),
});
// 3. 実行 — システム プロンプトにスキルが挿入され、ツールが呼び出され、応答がループされます。
var res = await Agent.RunAsync("顧客に対する返金注文 1234。", tk);
Console.WriteLine(res.Text);
ループは、フック、ストリーミング、再試行/バックオフを使用して、呼び出し→ツールの実行→フィードバック→繰り返しを実行します。
会話メモリも利用可能。 res には Text 、 ToolCalls 、使用法、ターン、およびモデルが含まれます。
Toolnexus を使用する。
// メソッド → ツール (属性ベース)
パブリック シールド クラス MathTools
{
[ToolMethod("add", "2 つの数値を加算")]
public string Add([Param("a")] double a, [Param("b")] double b) => (a + b).ToString();
}
tk.Register(Tools.FromObject(new MathTools()).ToArray());
// REST エンドポイント → ツール
tk.Register(HttpTool.Of(new HttpTool.Options
{
名前 = "create_ticket"、説明 = "チケットの作成"、メソッド = "POST"、
URL =

"https://api.example.com/tickets",
Headers = new Dictionary<string, string> { ["Authorization"] = "Bearer ${API_TOKEN}" }, // ${ENV} は展開され、ログには記録されません
InputSchema = new Dictionary<文字列, オブジェクト?>
{
["タイプ"] = "オブジェクト",
["properties"] = new Dictionary<string, object?> { ["title"] = new Dictionary<string, object?> { ["type"] = "string" } },
["required"] = new List<object?> { "title" },
}、
}));
ExtraTools/AnnotatedObjects を Toolkit.Options に直接渡すこともできます。
var tools = tk.ToOpenAI(); // または tk.ToAnthropic() / tk.ToGemini()
var system = tk.SkillsPrompt(); // システム プロンプトのスキル カタログ
// モデルがツール呼び出し (名前、引数) を返すとき:
var res = await tk.ExecuteAsync(name, args); // -> ToolResult(出力, IsError, メタデータ)
4つのソース
4 つすべてが tk.Tools() では 1 つの統一された ITool として表示されます。
完全なドキュメント、他の 4 つの言語ポート、共有動作仕様、および実行可能なサンプル:
https://github.com/muthuishere/toolnexus
ネット10.0
モデルコンテキストプロトコル
(>= 1.4.0)
このパッケージは、NuGet パッケージでは使用されません。
このパッケージは、一般的な GitHub リポジトリでは使用されません。
NuGet パッケージ エクスプローラーで開く
mcp
llm
エージェント
エージェントのスキル
オープンナイ
人間的な
ジェミニ
ツール
NuGet または NuGet ギャラリーについて質問がありますか?
NuGet.org とその関連サービスのサービス ステータスを確認します。
NuGet に関するよくある質問を読んで、あなたの質問がリストに載っているかどうかを確認してください。
© マイクロソフト 2026 -
について -
利用規約 -
あなたのプライバシーに関する選択 -
プライバシーに関する声明
- 商標

## Original Extract

Build an agent in a few lines: dynamic MCP servers + agent skills + native + HTTP tools as one tool set for any LLM, with OpenAI/Anthropic/Gemini adapters and a unified host loop. C# port of toolnexus.

NuGet Gallery
| Toolnexus 0.1.2
Skip To Content
Toggle navigation
Packages
dotnet add package Toolnexus --version 0.1.2
Copy
NuGet\Install-Package Toolnexus -Version 0.1.2
Copy
This command is intended to be used within the Package Manager Console in Visual Studio, as it uses the NuGet module's version of Install-Package .
<PackageReference Include="Toolnexus" Version="0.1.2" />
Copy
For projects that support PackageReference , copy this XML node into the project file to reference the package.
<PackageVersion Include="Toolnexus" Version="0.1.2" />
Directory.Packages.props
Copy
<PackageReference Include="Toolnexus" />
Project file
Copy
For projects that support Central Package Management (CPM) , copy this XML node into the solution Directory.Packages.props file to version the package.
paket add Toolnexus --version 0.1.2
Copy
The NuGet Team does not provide support for this client. Please contact its maintainers for support.
#r "nuget: Toolnexus, 0.1.2"
Copy
#r directive can be used in F# Interactive and Polyglot Notebooks. Copy this into the interactive tool or source code of the script to reference the package.
#:package Toolnexus@0.1.2
Copy
#:package directive can be used in C# file-based apps starting in .NET 10 preview 4. Copy this into a .cs file before any lines of code to reference the package.
#addin nuget:?package=Toolnexus&version=0.1.2
Install as a Cake Addin
Copy
#tool nuget:?package=Toolnexus&version=0.1.2
Install as a Cake Tool
Copy
The NuGet Team does not provide support for this client. Please contact its maintainers for support.
README
Build an agent in a few lines. Point at an mcp.json and a skills/ folder, call RunAsync() ,
and you have a working agent — MCP servers, agent skills, your own functions, and HTTP endpoints
unified as one tool set, driving any LLM.
Right-sized. Not a framework (no builders, no config to wade through), not a toy that falls
over the moment you need streaming or a retry. Everything a real agent needs — the loop, hooks,
streaming, retries, memory — and nothing it doesn't.
The C#/.NET port of toolnexus — the same library,
byte-identical, also in JavaScript, Python, Go, and Java . Built on the official
ModelContextProtocol SDK. Targets .NET 10.
dotnet add package Toolnexus
An agent in 3 steps
using Toolnexus;
// 1. tools from an mcp.json + a skills/ folder
await using var tk = await Toolkit.CreateAsync(new Toolkit.Options()
.WithMcpConfig("./mcp.json")
.WithSkillsDir("./skills"));
// 2. point at any OpenAI- or Anthropic-style endpoint
var agent = LlmClient.Create(new LlmClient.Options
{
BaseUrl = "https://openrouter.ai/api/v1",
Style = "openai", // or "anthropic"
Model = "openai/gpt-4o-mini",
ApiKey = Environment.GetEnvironmentVariable("OPENROUTER_API_KEY"),
});
// 3. run — skills injected into the system prompt, tools called for you, looped to an answer
var res = await agent.RunAsync("Refund order 1234 for the customer.", tk);
Console.WriteLine(res.Text);
The loop runs call → execute tools → feed back → repeat, with hooks, streaming, retries/backoff,
and conversation memory available. res carries Text , ToolCalls , usage, turns, and model.
using Toolnexus;
// a method → a tool (attribute-based)
public sealed class MathTools
{
[ToolMethod("add", "Add two numbers")]
public string Add([Param("a")] double a, [Param("b")] double b) => (a + b).ToString();
}
tk.Register(Tools.FromObject(new MathTools()).ToArray());
// a REST endpoint → a tool
tk.Register(HttpTool.Of(new HttpTool.Options
{
Name = "create_ticket", Description = "Create a ticket", Method = "POST",
Url = "https://api.example.com/tickets",
Headers = new Dictionary<string, string> { ["Authorization"] = "Bearer ${API_TOKEN}" }, // ${ENV} expands, never logged
InputSchema = new Dictionary<string, object?>
{
["type"] = "object",
["properties"] = new Dictionary<string, object?> { ["title"] = new Dictionary<string, object?> { ["type"] = "string" } },
["required"] = new List<object?> { "title" },
},
}));
You can also pass ExtraTools / AnnotatedObjects straight into Toolkit.Options .
var tools = tk.ToOpenAI(); // or tk.ToAnthropic() / tk.ToGemini()
var system = tk.SkillsPrompt(); // skills catalog for your system prompt
// when the model returns a tool call (name, arguments):
var res = await tk.ExecuteAsync(name, args); // -> ToolResult(Output, IsError, Metadata)
The four sources
All four appear as one uniform ITool in tk.Tools() .
Full docs, the other four language ports, the shared behavior spec, and runnable examples:
https://github.com/muthuishere/toolnexus
net10.0
ModelContextProtocol
(>= 1.4.0)
This package is not used by any NuGet packages.
This package is not used by any popular GitHub repositories.
Open in NuGet Package Explorer
mcp
llm
agents
agent-skills
openai
anthropic
gemini
tools
Got questions about NuGet or the NuGet Gallery?
Find out the service status of NuGet.org and its related services.
Read the Frequently Asked Questions about NuGet and see if your question made the list.
© Microsoft 2026 -
About -
Terms of Use -
Your Privacy Choices -
Privacy Statement
- Trademarks
