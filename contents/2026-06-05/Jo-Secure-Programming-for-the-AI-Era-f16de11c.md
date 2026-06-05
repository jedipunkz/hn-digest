---
source: "https://jo-lang.org/blog/2026-06-04-introducing-jo.html"
hn_url: "https://news.ycombinator.com/item?id=48406714"
title: "Jo – Secure Programming for the AI Era"
article_title: "Introducing Jo — Secure Programming for the AI Era | Jo Programming Language"
author: "rguiscard"
captured_at: "2026-06-05T00:57:12Z"
capture_tool: "hn-digest"
hn_id: 48406714
score: 1
comments: 0
posted_at: "2026-06-05T00:56:04Z"
tags:
  - hacker-news
  - translated
---

# Jo – Secure Programming for the AI Era

- HN: [48406714](https://news.ycombinator.com/item?id=48406714)
- Source: [jo-lang.org](https://jo-lang.org/blog/2026-06-04-introducing-jo.html)
- Score: 1
- Comments: 0
- Posted: 2026-06-05T00:56:04Z

## Translation

タイトル: Jo – AI 時代の安全なプログラミング
記事のタイトル: Jo の紹介 — AI 時代の安全なプログラミング |ジョープログラミング言語
説明: Jo は、きめ細かい機能ベースのプログラミングのための静的型付け言語です。権限の制限を実用的で、監査可能で、快適に使用できるように設計されています。

記事本文:
コンテンツへスキップ Jo 言語検索 K メイン ナビゲーション 概要 セキュリティ ビルド ツール 言語リファレンス ブログの外観
Jo のご紹介 — AI 時代の安全なプログラミング
今日は、静的型付けプログラミング言語である Jo を紹介します。この言語では、副作用はデフォルトで拒否され、コンパイラーによってチェックされるきめ細かい機能を通じて権限が明示的に付与される必要があります。
最新のシステムでは、プラグインを実行し、サードパーティのサービスを呼び出し、ユーザー定義のワークフローを実行し、AI エージェントにコードの生成と実行を求めることが増えています。セキュリティ上の質問は、もはや「このプログラムは正しいですか?」というだけではありません。また、次のとおりです。
信頼できないプログラムを、許可されているきめ細かい機能のみに制限するにはどうすればよいでしょうか?
Jo は、特定のディレクトリ、単一の API ホスト、読み取り専用インターフェイス、または現在のユーザーに属するデータベース行のみなど、実際のシステムが必要とする精度のレベルで、型システムによって適用されるきめ細かい権限の制限をプロパティにするように設計されています。
問題: 周囲の権威
ほとんどの主流言語では、強力な権限がデフォルトで利用可能になっています。通常、コードの一部は、ランタイム サンドボックスが停止しない限り、ファイル システム、環境変数、ネットワーク、リフレクション、プロセス API、または外部関数インターフェイスに到達できます。
このモデルは便利ですが、監査が困難です。サードパーティ関数を実行し、その関数がファイルの読み取りやネットワークの呼び出しではなく、狭い API のみをクエリできることを保証したい場合、通常、言語自体はほとんど役に立ちません。結局、コンテナ、権限、コードレビュー、規約、またはランタイム分離に依存することになります。
Jo は別の方法をとります。権限は明示的な機能によって表され、それらの機能はアプリケーションの必要に応じて狭い範囲に限定できます。コンパイラはコードのどの機能を追跡しますか

を使用する可能性があるため、制限は実行時設定に隠蔽されるのではなく、インターフェイスと型で表現されます。
能力ベースのプログラミング
Jo では、能力は通常のパラメータです。これらは、渡したり、改良したり、置換したり、制限したりできます。ケイパビリティを受け取っていない機能は使用できません。
def foo () = println "foo" // 推論された機能: stdout
def bar () = foo() // 推論された機能: stdout
def qux () は IO を受け取ります .stdout = println "qux" // 明示的な機能: stdout
デフォルトメイン =
bar() では none を許可します // エラー: stdout は許可されません
bar() で IO .stdout を許可 // OK
with IO .stdout = s => pass in qux() // 出力をリダイレクト コンパイラーは、呼び出しグラフを通じて機能フローをチェックします。関数に IO.stdout が必要な場合、その要件は表示され、制御可能です。呼び出しサイトが「allow none」と指定した場合、隠れた権限はすり抜けられません。
これにより、Jo はアンビエント グローバルのセキュリティ コストを支払うことなく、暗黙的なコンテキストの利便性を得ることができます。
AI 生成コードにとってこれが重要な理由
AI によって生成されたコードは、権限の問題をさらに深刻なものにします。エージェントがアプリケーションの関数を作成する場合、データを分析して概要を生成することはできますが、ファイル システムへのアクセス、任意の HTTP エンドポイントの呼び出し、環境変数の検査、または他のユーザーのレコードのクエリは実行しないでください。
Jo のアプローチは、コードに必要な機能のみを付与することです。
// API ライブラリ: FFI サポートなしでコンパイル
インターフェイス OrdersApi
def query (lastDays : Int ) : リスト [順序]
終わり
param orderApi : OrdersApi
// AI が生成したコード
def aiMain () : ユニットはordersApi、IO .stdoutを受け取ります =
val 注文 =ordersApi.query( 30 )
summary(orders) フレームワークは実際のデータベースを使用して OrdersApi を実装できますが、信頼できないコードにはユーザー スコープの読み取り専用ビューのみを公開します。 AI 生成関数は r を受け取りません

データベースへのアクセス。ネットワークアクセスは受け付けません。ファイルシステムへのアクセスは受け付けません。型チェッカーは、プログラムの実行前にその境界を強制します。
これが Jo の背後にある中心的なアイデアです。つまり、権限の制限をプログラミング モデルにするというものです。
CONFINED WORLD FFI なし · 限定されたライブラリのみ Jo 標準ライブラリ リスト · マップ · オプション · 結果 · … インターフェイス ライブラリ インターフェイス OrdersApi { … } defer def aiMain(): ユニット AI 生成コード aiMain() 実装コントラクト依存 TRUSTED WORLD FFI 有効 · 監査済みプラットフォーム ランタイム FFI · システムコール · ネットワーク · ファイルシステムハーネス UserScopedOrders(userId, db) FrameworkMain() dependes on depend on --link Two-World Architecture ページではこのモデルについて詳しく説明されており、Secure Language Design では、制限を実用化する言語機能 (機能パラメーターと権限の減衰) について説明しています。
具体的な例については、データ クエリ エージェントのデモを参照してください。このデモでは、エージェントが現在のユーザーのデータに静的に制限されながら、データベースに対して柔軟な質問を行う方法を示しています。
単なる政策システムではなく言語
Jo は、汎用言語として快適であることも意図されています。オブジェクト指向プログラミングと関数型プログラミングを、コンパクトな構文、型推論、クラス、インターフェイス、代数データ型、パターン マッチング、およびコンテキスト パラメーターと組み合わせます。
たとえば、Jo には再利用可能なパターン述語があります。
パターン 正: 部分 [ Int ] = case x if x > 0
パターン偶数 : 部分 [ Int ] = case x if x % 2 == 0
n に一致
case 正と偶数 => "正の偶数"
正の場合 => 「正の奇数」
case _ => "非正" およびパターン マッチングを使用した共用体型:
結合形状 =
円（半径：Float）
|長方形 (w : Float 、h : Float )
def エリア (形状 : 形状 ) : Float =
形状を一致させる
円 r => 3.14 * r * r の場合
およそ

se Rectangle w h => w * h Jo の設計哲学は、強力なセキュリティの保証とプログラマの幸福を結びつけることです。セキュリティでは、言語と格闘したり、定型文を作成したり、重要な推論を展開構成に移したりする必要はありません。目標は、安全なプログラミングが自然で表現力豊かで監査可能であると感じられるようにすることです。
Jo は、プログラマとセキュリティレビュー担当者の両方のために設計されています。機能の境界はインターフェイスと型で表現されるため、プログラムが受け取る権限は、実装の詳細やデプロイメント構成に分散されるのではなく、API 境界で確認できます。これにより、セキュリティ監査が簡素化されます。レビュー担当者は、どの機能が付与され、どこに流れ、どこで意図的に制限されているかを検査できます。
Jo の設計は、Coq で機械化された健全性証明を備えたコンテキスト機能の最小限の計算である λCC ( Lambda-CC ) に基づいています。
完全な論文と Coq 開発は github.com/typescope/contextual-capability にあります。
Jo は初期段階のソフトウェアですが、すでに充実しています。コンパイラには広範なテスト スイートがあり、コア機能モデルは本格的な実験に使用できるようになっています。言語設計、標準ライブラリ、ツールは依然として進化しています。
セキュリティを重視するチームには、新しいプロジェクト、プロトタイプ、内部ツール、および既存のテクノロジーでは必要な権限制限を提供できない制約のある運用ユースケースについて Jo を評価することをお勧めします。重要な展開の場合は、小規模から開始し、機能の境界を注意深く監査し、言語とツールが進化することを期待してください。
Jo は、安全なプログラミングを実用化することに重点を置いている会社である TypeScope によって開発されました。私たちは長期的なインフラストラクチャとして Jo を構築しています。言語、コンパイラー、標準ライブラリ、ドキュメント、エコシステムは、何年にもわたって着実に成長するように設計されています。

耳。
私たちの野心は高く、Jo をセキュリティ クリティカルなソフトウェアを作成するのに最適な言語の 1 つにし、安全なプログラミングを面倒ではなく自然に感じられるようにすることです。
Jo は、Apache License 2.0 に基づくオープンソースです。リポジトリは github.com/typescope/jo で入手できます。
このプロジェクトは、言語設計、機能ベースのセキュリティ、安全な AI システム、コンパイラ、実用的な型システムに興味のある人々を歓迎します。特にセキュリティ モデル、人間工学、実際の使用例に関するフィードバックを求めています。
言語表面については言語ツアーから始めるか、セキュリティ モデルについては Two-World Architecture を読んでください。インストールについては、インストール ガイドを参照してください。
言語設計者、セキュリティ エンジニア、コンパイラ エンジニア、エージェント システムを構築する開発者からのフィードバックを歓迎します。具体的なバグや問題については、GitHub で問題をオープンしてください。コミュニティでのディスカッションについては、 r/jolang に参加してください。セキュリティ レポートは、リポジトリの SECURITY.md のプロセスに従う必要があります。
核となるアイデアに共感を覚える場合は、プロジェクトに従って、例を試し、ディスカッションに参加してください。 Jo の使命はシンプルです。安全なプログラミングを楽しくすることです。

## Original Extract

Jo is a statically typed language for fine-grained capability-based programming. It is designed to make authority confinement practical, auditable, and pleasant to use.

Skip to content Jo Language Search K Main Navigation Overview Security Build Tool Language Reference Blog Appearance
Introducing Jo — Secure Programming for the AI Era ​
Today we are introducing Jo , a statically typed programming language where side effects are denied by default and authority must be granted explicitly, through fine-grained capabilities checked by the compiler.
Modern systems execute plugins, call third-party services, run user-defined workflows, and increasingly ask AI agents to generate and execute code. The security question is no longer only "is this program correct?" It is also:
How do we restrict an untrusted program to only the fine-grained capabilities it has been granted?
Jo is designed to make fine-grained permission confinement a property enforced by the type system, at the level of precision real systems need: a specific directory, a single API host, a read-only interface, or only the database rows belonging to the current user.
The Problem: Ambient Authority ​
Most mainstream languages make powerful authority available by default. A piece of code can usually reach for the filesystem, environment variables, network, reflection, process APIs, or foreign-function interfaces unless a runtime sandbox stops it.
That model is convenient, but it is hard to audit. If you want to run a third-party function and guarantee that it can only query a narrow API, not read files or call the network, the language itself usually gives you little help. You end up relying on containers, permissions, code review, convention, or runtime isolation.
Jo takes a different route: authority is represented by explicit capabilities, and those capabilities can be as narrowly scoped as the application requires. The compiler tracks which capabilities code may use, so confinement is expressed in interfaces and types rather than hidden in runtime configuration.
Capability-Based Programming ​
In Jo, capabilities are ordinary parameters. They can be passed, refined, substituted, and restricted. A function that has not received a capability cannot use it.
def foo () = println "foo" // inferred capability: stdout
def bar () = foo() // inferred capability: stdout
def qux () receives IO .stdout = println "qux" // explicit capability: stdout
def main =
allow none in bar() // error: stdout not allowed
allow IO .stdout in bar() // OK
with IO .stdout = s => pass in qux() // redirect output The compiler checks capability flow through the call graph. If a function needs IO.stdout , that requirement is visible and controllable. If a call site says allow none , then no hidden authority can slip through.
This gives Jo the convenience of implicit context without the security cost of ambient globals.
Why This Matters for AI-Generated Code ​
AI-generated code makes the authority problem even more acute. If an agent writes a function for your application, you may want it to analyze data and produce a summary, but not access the filesystem, call arbitrary HTTP endpoints, inspect environment variables, or query other users' records.
Jo's approach is to grant only the capabilities the code should have:
// API library: compiled without FFI support
interface OrdersApi
def query (lastDays : Int ) : List [ Order ]
end
param ordersApi : OrdersApi
// AI-generated code
def aiMain () : Unit receives ordersApi, IO .stdout =
val orders = ordersApi.query( 30 )
summarize(orders) The framework can implement OrdersApi using a real database, but expose only a user-scoped, read-only view to the untrusted code. The AI-generated function does not receive raw database access. It does not receive network access. It does not receive filesystem access. The type checker enforces that boundary before the program runs.
This is the core idea behind Jo: make authority confinement a programming model.
CONFINED WORLD no FFI · confined libs only Jo Standard Library List · Map · Option · Result · … Interface Library interface OrdersApi { … } defer def aiMain(): Unit AI-Generated Code aiMain() implements contract depends on depends on TRUSTED WORLD FFI enabled · audited Platform Runtime FFI · syscalls · network · filesystem Harness UserScopedOrders(userId, db) frameworkMain() depends on depends on --link The Two-World Architecture page describes this model in detail, and Secure Language Design covers the language facilities — capability parameters and authority attenuation — that make confinement practical.
For a concrete example, see the data-query agent demo , which shows how an agent can ask flexible questions over a database while being statically restricted to the current user's data.
A Language, Not Just a Policy System ​
Jo is also intended to be pleasant as a general-purpose language. It combines object-oriented and functional programming with a compact syntax, type inference, classes, interfaces, algebraic data types, pattern matching, and context parameters.
For example, Jo has reusable pattern predicates:
pattern Positive : Partial [ Int ] = case x if x > 0
pattern Even : Partial [ Int ] = case x if x % 2 == 0
match n
case Positive & Even => "positive even"
case Positive => "positive odd"
case _ => "non-positive" And union types with pattern matching:
union Shape =
Circle (radius : Float )
| Rectangle (w : Float , h : Float )
def area (shape : Shape ) : Float =
match shape
case Circle r => 3.14 * r * r
case Rectangle w h => w * h Jo's design philosophy is to combine strong security guarantees with programmer happiness. Security should not require fighting the language, writing boilerplate, or moving essential reasoning into deployment configuration. The goal is to make secure programming feel natural, expressive, and auditable.
Jo is designed for both programmers and security reviewers. Capability boundaries are expressed in interfaces and types, so the authority a program receives is visible at the API boundary rather than scattered through implementation details or deployment configuration. This makes security auditing simpler: reviewers can inspect what capabilities are granted, where they flow, and where they are deliberately restricted.
Jo's design is grounded in λCC ( Lambda-CC ), a minimal calculus of contextual capabilities with a soundness proof mechanized in Coq.
The full paper and Coq development are at github.com/typescope/contextual-capability .
Jo is early-stage software, but it is already substantial: the compiler has an extensive test suite, and the core capability model is ready for serious experimentation. The language design, standard library, and tooling are still evolving.
We encourage security-focused teams to evaluate Jo for new projects, prototypes, internal tools, and constrained production use cases where existing technologies cannot provide the authority confinement they need. For critical deployments, start small, audit the capability boundaries carefully, and expect the language and tooling to evolve.
Jo is developed by TypeScope , a company focused on making secure programming practical. We are building Jo as long-term infrastructure: a language, compiler, standard library, documentation, and ecosystem designed to grow steadily over many years.
Our ambition is high: to make Jo one of the best languages for writing security-critical software, and to make secure programming feel natural rather than burdensome.
Jo is open source under the Apache License 2.0. The repository is available at github.com/typescope/jo .
The project welcomes people interested in language design, capability-based security, secure AI systems, compilers, and practical type systems. We especially want feedback on the security model, ergonomics, and real-world use cases.
Start with the Language Tour for the language surface, or read Two-World Architecture for the security model in more detail. For installation, see the install guide .
We welcome feedback from language designers, security engineers, compiler engineers, and developers building agentic systems. For concrete bugs or issues, open an issue on GitHub . For community discussion, join r/jolang . Security reports should follow the process in the repository's SECURITY.md .
If the core idea resonates with you, follow the project, try the examples, and join the discussion. Jo's mission is simple: make secure programming a joy.
