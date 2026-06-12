---
source: "https://jo-lang.org/blog/2026-06-11-why-compile-time-sandboxing.html"
hn_url: "https://news.ycombinator.com/item?id=48502389"
title: "Why Secure AI Needs Compile-Time Sandboxing"
article_title: "Why Secure AI Needs Compile-Time Sandboxing | Jo Secure Programming Language"
author: "liu-fengyun"
captured_at: "2026-06-12T13:19:11Z"
capture_tool: "hn-digest"
hn_id: 48502389
score: 1
comments: 0
posted_at: "2026-06-12T10:48:31Z"
tags:
  - hacker-news
  - translated
---

# Why Secure AI Needs Compile-Time Sandboxing

- HN: [48502389](https://news.ycombinator.com/item?id=48502389)
- Source: [jo-lang.org](https://jo-lang.org/blog/2026-06-11-why-compile-time-sandboxing.html)
- Score: 1
- Comments: 0
- Posted: 2026-06-12T10:48:31Z

## Translation

タイトル: セキュア AI にコンパイル時のサンドボックス化が必要な理由
記事のタイトル: 安全な AI にコンパイル時のサンドボックス化が必要な理由 |ジョーセキュアプログラミング言語
説明: ランタイム サンドボックスはプロセスを制限します。コンパイル時のサンドボックス化によりコードが制限されます。 AI によって生成されたコードの場合、境界はビジネス ロジックを理解する必要があります。これは型チェッカーの仕事です。

記事本文:
コンテンツへスキップ Jo 言語検索 K メイン ナビゲーション 概要 セキュリティ ビルド ツール 言語リファレンス ブログの外観
セキュア AI にコンパイル時のサンドボックス化が必要な理由
AI エージェントは、コードを生成して実行することでタスクを自動化する能力が高まっています。しかし、AI エージェントを完全に信頼することはできません。今年初め、OpenClaw は Meta Superintelligence Labs のアラインメントディレクターの受信箱から数百件の電子メールを削除しました。
エージェントのアクションが厳密に制限されるまでは、重要なインフラストラクチャや機密データに対してエージェントを使用することはできません。信頼できないコードには具体的に何が許可されているのでしょうか?誰がチェックしたのでしょうか?
標準的な答えは、コンテナ、seccomp フィルター、VM などのランタイム サンドボックスです。しかし、それらは間違った抽象化レベルで動作します。コンテナは、プロセスがソケットを開くのを停止できます。信頼できないコードが、正当に保持しているデータベース接続を介して別のユーザーの行をクエリすることを阻止することはできません。
実行時とコンパイル時のサンドボックス化
ランタイム サンドボックスは、オペレーティング システムの言語 (プロセス、ファイル、ソケット) を話します。強制する価値のあるルール (どの行、どのエンドポイント、どのユーザーのデータ) をアプリケーションの言語で記述する必要があり、境界は表現できる内容のみを強制できます。型は、アプリケーション レベルで機能する唯一の境界です。
コンパイル時のサンドボックス化、AI に適用
最近の ACM Queue の記事「セーフ コーディング」で、Christoph Kern は、数十年にわたる Google のセキュリティ エンジニアリングを厳密なモジュール推論の原則に抽出しました。
... 抽象化内の危険な操作の安全性は、抽象化の API と型署名によってサポートされる仮定にのみ依存する必要があります。逆に、安全な抽象化と安全なコード (つまり、プログラムの大部分を構成する危険な操作のないコード) の構成は自動的に行われます。

実装言語の型チェッカーによって厳密に検証されます。
それはまさに、Jo のコンパイル時サンドボックスが行っていることです。
Jo では、ファイルシステム アクセス、ネットワーク呼び出し、FFI、データベース クエリなどの危険な操作には明示的なアクセス許可が必要であり、機能として関数タイプ シグネチャに表示されます。
次の例では、AI によって生成されたコードは、与えられた機能に限定されています。
インターフェイス OrdersApi
def query (lastDays : Int ) : リスト [順序]
終わり
param orderApi : OrdersApi
// 信頼できないコード (AI 生成): 受け取ったもののみを使用できます
def aiMain () : ユニットはordersApi、IO .stdoutを受け取ります =
val 注文 =ordersApi.query( 30 )
printOrders :orders.select(o => o.state == "open" ) aiMain がファイルシステム、ネットワーク、またはスコープ外のクエリに到達しようとしても、プログラムは実行時に誤動作せず、コンパイルに失敗します。
orderApi 機能は、信頼できるハーネスによって提供されます。たとえば、実際のデータベースに対するユーザー スコープの読み取り専用ビューとして提供されます。しかし、実装は AI コードに対して不透明です。AI コードが認識するのは OrdersApi インターフェイスだけです。
ランタイムサンドボックス + REST API
一般的な中間点は、エージェントをサンドボックス化し、REST API を通じてのみ世界にアクセスできるようにすることです。これは役に立ちますが、設計の不一致を引き継いでいます。REST API は信頼できる呼び出し元向けに設計されており、大きな機能面を公開しています。一般的な API トークンは、Web インターフェイスを通じてユーザーが実行できるすべての操作 (すべての注文の読み取り、アカウントの変更、レコードの削除) を許可します。これは、単一のエージェント タスクが必要とするものをはるかに超えています。
エージェントは信頼できる発信者ではありません。何を呼び出すかは、実行中のタスクのセキュリティ コンテキストによって異なります。1 つのタスクは API サーフェスの小さなサブセットのみを参照する必要があります。もう 1 つはエンドポイントを必要としますが、そのスコープは狭められています。このユーザーの行は読み取り専用で、過去 30 日間です。既存のREST API

はそのようなニーズに応えるように設計されていないため、ゲートウェイとエージェントごとのプロキシ ポリシーに制限が適用されることになります。権限はインフラストラクチャ構成に戻り、上記のすべての監査問題が発生します。
機能を使用すると、同じ絞り込みが信頼できるコードの数行になります。つまり、API をより小さなインターフェイスにラップします。
MCP は、LLM の機能をアプリケーション定義のセキュリティ ポリシーに制限できる、精査されたツールのカタログを LLM に提供します。この方法で問題を解決できるのであれば、そうすべきです。
ただし、複雑なユースケースの場合、このアプローチは限界に達します。すべてのツール定義は、作業が開始される前にコンテキスト ウィンドウにロードされます。すべての中間結果はモデル内を往復する必要があり、各ステップでトークンを書き込みます。
また、固定ツール セットには柔軟性が限られています。タスクでフィルター、結合、または結果のループが必要な場合、エージェントはさらに別のツールを必要とするか、トークンごとにツールをシミュレートする必要があります。 「Executable Code Actions Elicit Better LLM Agents」という調査では、生成されたコードを通じて動作するエージェントが、ツール呼び出しに限定されたエージェントよりも優れたパフォーマンスを発揮することが示されています。
Anthropic のエンジニアリング チームも提唱している自然な解決策は、エージェントにツールを呼び出すコードを記述させることです。フィルタリングと合成はプログラム内で行われ、最終結果のみがモデルに返されます。しかし、それでは元の疑問が戻ってきます。生成されたコードは正確に何をすることができるのでしょうか?コンパイル時のサンドボックス化がその答えです。
Jo は、信頼できるコンパイル ワールドと限定されたコンパイル ワールドを厳密に分離することにより、コンパイル時サンドボックスを実装します。「Two-World Architecture」を参照してください。機能の仕組み (コンテキスト パラメーター、受信句、および権限の減衰) は、「安全な言語設計」で説明されており、コンテキスト機能のモデルに基づいています。
クリストフ・カーン。安全なコーディング: 厳密なモジュール推論

ソフトウェアの安全性。 ACM キュー 23(5)、2025 年 11 月。
Xingyao Wang、Yangyi Chen、Lifan Yuan、Yunzhe Zhang、Yunzhu Li、Hao Peng、Heng Ji。実行可能コードのアクションにより、より優れた LLM エージェントが導き出されます。 ICML 2024。

## Original Extract

Runtime sandboxes confine processes; compile-time sandboxing confines code. For AI-generated code, the boundary must understand business logic — and that is a job for the type checker.

Skip to content Jo Language Search K Main Navigation Overview Security Build Tool Language Reference Blog Appearance
Why Secure AI Needs Compile-Time Sandboxing ​
AI agents are becoming more capable at automating tasks by generating and executing code. But AI agents cannot be fully trusted. Earlier this year, OpenClaw deleted hundreds of emails from the inbox of Meta Superintelligence Labs' alignment director.
Until an agent's actions are strictly bounded, we cannot use it on critical infrastructure or with sensitive data. What exactly is the untrusted code allowed to do — and who checked?
The standard answer is runtime sandboxing: containers, seccomp filters, VMs. But they operate at the wrong level of abstraction. A container can stop a process from opening a socket; it cannot stop untrusted code from querying another user's rows through a database connection it legitimately holds.
Runtime vs. Compile-Time Sandboxing ​
A runtime sandbox speaks the operating system's language: processes, files, sockets. The rules worth enforcing need to be written in the application's language — which rows, which endpoints, which user's data — and a boundary can only enforce what it can express. Types are the only boundary that operates at the application level.
Compile-Time Sandboxing, Applied to AI ​
In a recent ACM Queue article, Safe Coding , Christoph Kern distills decades of Google's security engineering into a principle of rigorous modular reasoning:
... the safety of risky operations within an abstraction must rely solely on assumptions supported by the abstraction's APIs and type signatures. Conversely, the composition of safe abstractions with safe code (i.e., code free of risky operations, which constitutes the vast majority of a program) is automatically verified by the implementation language's type checker.
That is exactly what Jo's compile-time sandboxing is doing.
In Jo, risky operations — filesystem access, network calls, FFI, database queries — need explicit permissions, visible in function type signatures as capabilities.
In the following example, the AI-generated code is confined to the capabilities it is given:
interface OrdersApi
def query (lastDays : Int ) : List [ Order ]
end
param ordersApi : OrdersApi
// Untrusted code (AI-generated): can use only what it received
def aiMain () : Unit receives ordersApi, IO .stdout =
val orders = ordersApi.query( 30 )
printOrders : orders.select(o => o.state == "open" ) If aiMain tries to reach the filesystem, the network, or an unscoped query, the program does not misbehave at runtime — it fails to compile.
The ordersApi capability is provided by the trusted harness — for example, as a user-scoped, read-only view over the real database. But the implementation is opaque to the AI code: all it ever sees is the OrdersApi interface.
Runtime Sandboxing + REST APIs ​
A popular middle ground is to sandbox the agent and let it reach the world only through REST APIs. This helps, but it inherits a design mismatch: REST APIs were designed for trusted callers and they expose a large capability surface . A typical API token grants everything the user can do through the web interface — read all orders, change the account, delete records — far more than any single agent task needs.
An agent is not a trusted caller. What it may call depends on the security context of the task at hand: one task should see only a small subset of the API surface; another needs an endpoint, but with its scope narrowed — this user's rows, read-only, the last 30 days. Existing REST APIs were not designed to answer such needs, so the restrictions end up in gateways and per-agent proxy policies — authority drifts back into infrastructure configuration, with all the audit problems above.
With capabilities, the same narrowing is a few lines of trusted code: wrap the API in a smaller interface.
MCP gives LLMs a catalog of vetted tools that can restrict LLM capabilities to application-defined security policies. If a problem can be solved this way, it should be.
However, for complex use cases, the approach hits limits. Every tool definition is loaded into the context window before any work happens. Every intermediate result must round-trip through the model, burning tokens at each step.
And a fixed tool set has limited flexibility: when a task calls for a filter, a join, or a loop over results, the agent either needs yet another tool or has to simulate it token by token. The research Executable Code Actions Elicit Better LLM Agents shows that agents acting through generated code outperform those restricted to tool calls.
The natural fix — also advocated by Anthropic's engineering team — is to let the agent write code that calls the tools: filtering and composition happen in the program, and only the final result returns to the model. But that brings back the original question: what exactly may the generated code do? Compile-time sandboxing is the answer.
Jo implements compile-time sandboxing through a hard separation between trusted and confined compilation worlds — see the Two-World Architecture . The capability mechanics — context parameters, receives clauses, and authority attenuation — are described in Secure Language Design , and are grounded in the model of contextual capabilities .
Christoph Kern. Safe Coding: Rigorous Modular Reasoning about Software Safety . ACM Queue 23(5), November 2025.
Xingyao Wang, Yangyi Chen, Lifan Yuan, Yizhe Zhang, Yunzhu Li, Hao Peng, Heng Ji. Executable Code Actions Elicit Better LLM Agents . ICML 2024.
