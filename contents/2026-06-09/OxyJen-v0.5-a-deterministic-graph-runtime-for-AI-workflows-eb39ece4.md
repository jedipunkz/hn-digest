---
source: "https://github.com/11divyansh/OxyJen"
hn_url: "https://news.ycombinator.com/item?id=48456722"
title: "OxyJen v0.5: a deterministic graph runtime for AI workflows"
article_title: "GitHub - 11divyansh/OxyJen: OxyJen is an open-source Java framework for orchestrating LLM workloads with graph-style execution, context-aware memory, and deterministic retry/fallback. It treats LLMs as native nodes (not helper utilities), allowing developers to build multi-step AI pipelines that int\n[truncated]"
author: "bdivyansh11"
captured_at: "2026-06-09T07:22:27Z"
capture_tool: "hn-digest"
hn_id: 48456722
score: 1
comments: 0
posted_at: "2026-06-09T05:10:55Z"
tags:
  - hacker-news
  - translated
---

# OxyJen v0.5: a deterministic graph runtime for AI workflows

- HN: [48456722](https://news.ycombinator.com/item?id=48456722)
- Source: [github.com](https://github.com/11divyansh/OxyJen)
- Score: 1
- Comments: 0
- Posted: 2026-06-09T05:10:55Z

## Translation

タイトル: OxyJen v0.5: AI ワークフロー用の決定論的グラフ ランタイム
記事のタイトル: GitHub - 11divyansh/OxyJen: OxyJen は、グラフ スタイルの実行、コンテキスト認識メモリ、決定論的な再試行/フォールバックを備えた LLM ワークロードを調整するためのオープンソース Java フレームワークです。 LLM を (ヘルパー ユーティリティではなく) ネイティブ ノードとして扱い、開発者が複数のステップからなる AI パイプラインを構築できるようにします。
[切り捨てられた]
説明: OxyJen は、グラフ スタイルの実行、コンテキスト認識メモリ、確定的な再試行/フォールバックを使用して LLM ワークロードを調整するためのオープンソース Java フレームワークです。 LLM を (ヘルパー ユーティリティではなく) ネイティブ ノードとして扱うため、開発者は既存の AI パイプラインときれいに統合する複数ステップの AI パイプラインを構築できます。
[切り捨てられた]

記事本文:
GitHub - 11divyansh/OxyJen: OxyJen は、グラフ スタイルの実行、コンテキスト認識メモリ、決定論的な再試行/フォールバックを備えた LLM ワークロードを調整するためのオープンソース Java フレームワークです。 LLM を (ヘルパー ユーティリティではなく) ネイティブ ノードとして扱うため、開発者は既存の Java コードときれいに統合するマルチステップ AI パイプラインを構築できます。 · GitHub
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
11ディヴィアンシュ
/
オキシジェン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
409 コミット 409 コミット .settings .settings docs docs src src .classpath .classpath .gitignore .gitignore .project .project CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md pom.xml pom.xml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
OxyJen は、Java および JVM 企業向けに欠けている決定論的 AI ランタイムです。
JVM 用の確定的 AI ワークフロー ランタイム - 複雑な AI パイプラインをシンプルかつ強力に構築します。
Oxyjen は、Java で AI アプリケーションを構築するためのグラフベースのオーケストレーション フレームワークです。 LLM、データ プロセッサ、カスタム ロジックを強力なワークフローに接続するためのクリーンで拡張可能なアーキテクチャを提供します。
これを AI パイプラインの配管と考えてください。各ステップの動作に集中し、Oxyjen が実行フローを処理します。
「LangChain4j が存在するのに、なぜ Oxyjen なのか?」
わかりました、これがあなたが考える最初の質問です。正直に言わせてください。
私は LangChain4j の存在を知らずに Oxyjen の構築を始めました。途中でそれを発見したとき、私には選択肢がありました：
差別化することにしました。
OSS がどのように機能するかを知りたかった。
これを公共の場で作りたかったのです。
LangChain4j は、幅広い機能、多数の統合、多数のツールに重点を置いた堅牢なフレームワークです。これは多くのユースケースに最適です。
Oxyjen は、開発者のエクスペリエンスと本番環境の準備に重点を置いて、別の道を歩んでいます
Oxyjen は実行時の信頼性を目的としており、グラフが自己認識し、障害が少なくなるようにします。たとえノードに障害が発生した場合でも、Oxyjen はそこから学習し、

証明する。
非同期、プロジェクト ルーム、並列処理、Java 同時実行などの機能は、Oxyjen のフェイルセーフ グラフ構造の基礎を築きます。
私は Langchain4j と競合するためにここにいるのではなく、開発者向けに信頼できる実行エンジンを作成するためにここにいます。
最新の AI アプリケーションには、API 呼び出し以上のものが必要です。彼らには以下が必要です:
複数のステップを含む複雑なワークフロー
コンパイル時にエラーをキャッチするタイプ セーフティ
何が起こっているかをデバッグするための可観測性
信頼性を保証するテスト可能性
カスタムロジックを追加できる拡張性
Oxyjen は、これらすべてをシンプルで直感的な API で提供します。
// 3 ステップのテキスト処理パイプラインを構築します
グラフ パイプライン = GraphBuilder 。名前付き (「テキストプロセッサ」)
。 addNode ( new UppercaseNode ())
。 addNode ( new ReverseNode ())
。 addNode ( new PrefixNode ( "OUTPUT: " ))
。建てる （）;
// コンテキストを使用して実行
NodeContext コンテキスト = new NodeContext ();
エグゼキュータ executor = 新しいエグゼキュータ ();
文字列結果 = executor 。 run ( パイプライン 、 "hello world" 、 コンテキスト );
システム 。外 。 println (結果);
// 出力: 出力: DLROW OLLEH
それです！クリーン、シンプル、パワフル。
Oxyjen は 4 つの核となる概念を中心に構築されています。
1️ グラフ - パイプラインのブループリント
グラフはパイプラインの構造、つまりどのノードがどのような順序で実行されるかを定義します。
パブリック クラス グラフ {
プライベート最終文字列名 ;
private Final List < NodePlugin <?, ?>> ノード ;
// パイプラインにノードを追加します
public Graph addNode ( NodePlugin <?, ?> ノード );
// すべてのノードを実行順に取得します
public List < NodePlugin <?, ?>> getNodes ();
}
これを次のように考えてください。パイプラインの DNA - 何が起こる必要があるかを知っていますが、何も実行しません。
2️ NodePlugin - 処理ユニット
NodePlugin はパイプラインの 1 つのステップです。各ノードは入力を出力に変換します。
パブリック インターフェイス NodePlugin < I , O > {
// コア処理ロジック
Oプロセス（入力、NodeCon

テキストコンテキスト);
// このノードの一意の識別子
デフォルトの String getName () {
これを返してください。 getClass()。 getSimpleName();
}
// セットアップ/クリーンアップ用のライフサイクル フック
デフォルト void onStart (NodeContext context ) {}
デフォルト void onFinish ( NodeContext context ) {}
デフォルト void onError (Exception e , NodeContext context ) {}
}
それを次のように考えてください。レゴ ブロック - 小さく、焦点を合わせ、組み立て可能です。
public class SummarizerNode は NodePlugin を実装します < String , String > {
@オーバーライド
public String process (String input , NodeContext context ) {
コンテキスト 。 getLogger()。 info ( "テキストを要約しています..." );
// ここにロジックを記述します (v0.2 では LLM 呼び出しになります)
return "概要: " + input 。部分文字列 ( 0 , 100 );
}
@オーバーライド
public void onStart (NodeContext context ) {
コンテキスト 。 getLogger()。 info ( "サマライザー ノードの開始" );
}
}
3️ エグゼキュータ - ランタイム エンジン
Executor はグラフを実行し、各ノードを順番に呼び出して出力を入力に渡します。
パブリック クラス エグゼキュータ {
public < I , O > O run (Graph グラフ , I 入力 , NodeContext コンテキスト ) {
// グラフ構造を検証します
// ノードを順番に実行します
// エラーとライフサイクルフックを処理します
// 最終出力を返します
}
}
これを次のように考えてください。オーケストラの指揮者はすべてを調整します。
グラフと初期入力を受け取ります
各ノードについて:
onStart() ライフサイクル フックを呼び出します
現在のデータを使用して process() を実行します
onFinish() ライフサイクル フックを呼び出します
4️ NodeContext - 共有メモリと状態
NodeContext はすべてのノード間で共有され、ログ記録と状態管理を提供します。
パブリック クラス NodeContext {
// 共有データの保存/取得
public void set (String key , Object value );
public < T > T get ( String key );
// ロギング
パブリックロガー getLogger();
public OxyLogger getOxyjenLogger ();
// メタデータ (グラフ名、実行 ID など)
public void setMetadata (文字列キー

、オブジェクト値 );
public < T > T getMetadata ( String key );
// エラー処理
public ExceptionHandler getExceptionHandler();
}
すべてのノードが読み取り/書き込みできる共有ノートブックとして考えてください。
public String プロセス ( String input 、 NodeContext ctx ) {
// 何が起こっているかを記録します
ctx 。 getLogger()。 info ( "処理: " + input );
// 中間結果を保存する
ctx 。 set ( "word_count" , input .split ( " " ). length );
// ノード間でデータを共有する
文字列PreviousResult = ctx 。 get ( "前の出力" );
処理された出力を返します。
}
完全な実例
パッケージの例 ;
インポートイオ。オキシジェン。コア.*;
パブリック クラス ContentPipeline {
public static void main ( String [] args ) {
// ステップ 1: ノードを定義する
NodePlugin < String , String > validator = new ValidationNode ();
NodePlugin < String , String > プロセッサ = new ProcessingNode ();
NodePlugin < String , String > formatter = new FormatterNode ();
// ステップ 2: グラフを構築する
グラフ パイプライン = GraphBuilder 。名前付き (「コンテンツパイプライン」)
。 addNode (バリデーター)
。 addNode (プロセッサ)
。 addNode (フォーマッタ)
。建てる （）;
// ステップ 3: 実行コンテキストを作成する
NodeContext コンテキスト = new NodeContext ();
コンテクスト 。 set ( "max_length" , 100 );
// ステップ 4: 実行
エグゼキュータ executor = 新しいエグゼキュータ ();
文字列結果 = executor 。 run ( パイプライン , "生の入力テキスト" , コンテキスト );
システム 。外 。 println ( "最終出力: " + result );
システム 。外 。 println ( "単語数: " + context .get ( "word_count" ));
}
}
// ノードの実装例
class ValidationNode は NodePlugin を実装します < String , String > {
@オーバーライド
public String プロセス ( String input 、 NodeContext ctx ) {
if ( input == null || input .isEmpty ()) {
throw new IllegalArgumentException ( "入力を空にすることはできません" );
}
ctx 。 getLogger()。 info ( "入力が検証されました" );
入力を返す

;
}
}
class ProcessingNode は NodePlugin を実装します < String , String > {
@オーバーライド
public String プロセス ( String input 、 NodeContext ctx ) {
処理された文字列 = input 。 to大文字()。トリム();
ctx 。 set ( "word_count" 、処理済み。split ( " " ). length );
ctx 。 getLogger()。 info ( "テキストが処理されました" );
返品処理済み。
}
}
class FormatterNode は NodePlugin を実装します < String , String > {
@オーバーライド
public String プロセス ( String input 、 NodeContext ctx ) {
整数 maxLength = ctx 。 get ( "max_length" );
フォーマットされた文字列 = input 。長さ () > 最大長
?入力を行ってください。部分文字列 ( 0 , maxLength ) + "..."
: 入力 ;
ctx 。 getLogger()。 info ( "テキストのフォーマット済み" );
フォーマットされたものを返します。
}
}
オキシジェンに対する私のビジョン
AI オーケストレーション (LangChain/LangGraph スタイル) を Java に導入します。
エンタープライズファーストのモジュールを構築します: LLM エージェント、監査ツール、安全な複雑なワークフロー エンジン。
パフォーマンス、セキュリティ、可観測性に重点を置きます。
Javaをより深く学ぶためにこれを構築しています。
RAG サポート - ベクター データベース、埋め込み、ドキュメント ローダー
コスト管理 - 予算、制限、使用状況の追跡
エンタープライズ機能 - 監査ログ、RBAC、コンプライアンス
マルチテナンシー - ユーザー/組織間でデータを分離します。
サーキットブレーカー - サービスがダウンするとすぐに故障します
トークンカウントとコスト追跡
< リポジトリ >
< リポジトリ >
< id >jitpack.io</ id >
< url >https://jitpack.io</ url >
</ リポジトリ >
</ リポジトリ >
依存関係を追加します。
<依存関係>
< groupId >com.github.11divyansh</ groupId >
< artifactId >オキシジェン</ artifactId >
< バージョン >v0.4.0</ バージョン >
</ 依存関係 >
グラドル
リポジトリ {
maven { URL ' https://jitpack.io ' }
}
依存関係 {
実装「com.github.11divyansh:Oxyjen:v0.4.0」
}
ソースからビルドする
git クローン https://github.com/11divyansh/OxyJen.git
cd オキシジェン
mvnクリーンインストール
インストール後、impoで検証

評価:
インポートイオ。オキシジェン。コア.*;
インポートイオ。オキシジェン。 llm .*;
インポートイオ。オキシジェン。ツール .*;
インポートイオ。オキシジェン。グラフ .*;
について
Java は世界クラスの AI ツールに値すると信じている BTech CS の学生である Divyansh Bhatt によって ❤️ で構築されました。
これは学習プロジェクトとして始まりましたが、本番環境に対応できるように全力で取り組んでいます。これがまだ大きくないことは承知していますが、価値あるものにしましょう。
このリポジトリにスターを付けて旅をフォローし、参加してください
ディスカッションを通じて機能を提案する
コードまたはドキュメントを提供する
役に立つと思ったら Twitter/LinkedIn で共有してください
** v0.5 の進捗状況の最新情報に注目してください!**
Apache 2.0 (オープンソース、エンタープライズ向け)
OxyJen は、グラフ スタイルの実行、コンテキスト認識メモリ、決定論的な再試行/フォールバックを備えた LLM ワークロードを調整するためのオープンソース Java フレームワークです。 LLM を (ヘルパー ユーティリティではなく) ネイティブ ノードとして扱うため、開発者は既存の Java コードときれいに統合するマルチステップ AI パイプラインを構築できます。
Apache-2.0ライセンス
貢献する
読み込み中にエラーが発生しました。このページをリロードしてください。
2
フォーク
レポートリポジトリ
リリース
5
OxyJen v0.5.1: Java の AI ワークフロー用の決定論的 DAG ランタイム
最新の
2026 年 6 月 8 日
+ 4 リリース
パッケージ
0
読み込み中にエラーが発生しました。プル

[切り捨てられた]

## Original Extract

OxyJen is an open-source Java framework for orchestrating LLM workloads with graph-style execution, context-aware memory, and deterministic retry/fallback. It treats LLMs as native nodes (not helper utilities), allowing developers to build multi-step AI pipelines that integrate cleanly with existing
[truncated]

GitHub - 11divyansh/OxyJen: OxyJen is an open-source Java framework for orchestrating LLM workloads with graph-style execution, context-aware memory, and deterministic retry/fallback. It treats LLMs as native nodes (not helper utilities), allowing developers to build multi-step AI pipelines that integrate cleanly with existing Java code. · GitHub
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
11divyansh
/
OxyJen
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
409 Commits 409 Commits .settings .settings docs docs src src .classpath .classpath .gitignore .gitignore .project .project CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md pom.xml pom.xml View all files Repository files navigation
OxyJen is the missing deterministic AI Runtime for Java & JVM enterprises.
Deterministic AI Workflow Runtime for the JVM - Build complex AI pipelines with simplicity and power.
Oxyjen is a graph-based orchestration framework for building AI applications in Java. It provides a clean, extensible architecture for connecting LLMs, data processors, and custom logic into powerful workflows.
Think of it as the plumbing for your AI pipelines , you focus on what each step does, Oxyjen handles the execution flow.
"Why Oxyjen When LangChain4j Exists?"
I get it, this is the first question you're thinking. Let me be completely honest.
I started building Oxyjen without knowing LangChain4j existed. When I discovered it halfway through, I had a choice:
I chose to differentiate.
I wanted to learn how OSS works.
I wanted to build this in public.
LangChain4j is a solid framework focused on feature breadth , lots of integrations, lots of tools. That's great for many use cases.
Oxyjen is taking a different path, focused on developer experience and production readiness
Oxyjen is meant for runtime reliability, your graphs will be self-aware and will make sure to provide less failure, even if a node fails, Oxyjen will learn from it and improve.
Features like, async, project loom, parallel processing, java concurrency will lay down the foundation of fail-safe graph structure for Oxyjen.
I'm not here to compete with Langchain4j, I'm here to create a reliable execution engine for devs.
Modern AI applications need more than just API calls. They need:
Complex workflows with multiple steps
Type safety to catch errors at compile time
Observability to debug what's happening
Testability to ensure reliability
Extensibility to add custom logic
Oxyjen provides all of this with a simple, intuitive API.
// Build a 3-step text processing pipeline
Graph pipeline = GraphBuilder . named ( "text-processor" )
. addNode ( new UppercaseNode ())
. addNode ( new ReverseNode ())
. addNode ( new PrefixNode ( "OUTPUT: " ))
. build ();
// Execute with context
NodeContext context = new NodeContext ();
Executor executor = new Executor ();
String result = executor . run ( pipeline , "hello world" , context );
System . out . println ( result );
// Output: OUTPUT: DLROW OLLEH
That's it! Clean, simple, powerful.
Oxyjen is built around four core concepts:
1️ Graph - The Pipeline Blueprint
A Graph defines the structure of your pipeline - which nodes run in what order.
public class Graph {
private final String name ;
private final List < NodePlugin <?, ?>> nodes ;
// Add nodes to your pipeline
public Graph addNode ( NodePlugin <?, ?> node );
// Get all nodes in execution order
public List < NodePlugin <?, ?>> getNodes ();
}
Think of it as: Your pipeline's DNA - it knows what needs to happen, but doesn't execute anything.
2️ NodePlugin - The Processing Unit
A NodePlugin is a single step in your pipeline. Each node transforms input into output.
public interface NodePlugin < I , O > {
// Core processing logic
O process ( I input , NodeContext context );
// Unique identifier for this node
default String getName () {
return this . getClass (). getSimpleName ();
}
// Lifecycle hooks for setup/cleanup
default void onStart ( NodeContext context ) {}
default void onFinish ( NodeContext context ) {}
default void onError ( Exception e , NodeContext context ) {}
}
Think of it as: A Lego brick - small, focused, composable.
public class SummarizerNode implements NodePlugin < String , String > {
@ Override
public String process ( String input , NodeContext context ) {
context . getLogger (). info ( "Summarizing text..." );
// Your logic here (will be LLM call in v0.2)
return "Summary: " + input . substring ( 0 , 100 );
}
@ Override
public void onStart ( NodeContext context ) {
context . getLogger (). info ( "Summarizer node starting" );
}
}
3️ Executor - The Runtime Engine
The Executor runs your graph, calling each node in sequence and passing outputs to inputs.
public class Executor {
public < I , O > O run ( Graph graph , I input , NodeContext context ) {
// Validates graph structure
// Executes nodes sequentially
// Handles errors and lifecycle hooks
// Returns final output
}
}
Think of it as: The conductor of an orchestra - coordinates everything.
Takes your Graph and initial input
For each node:
Calls onStart() lifecycle hook
Executes process() with current data
Calls onFinish() lifecycle hook
4️ NodeContext - Shared Memory & State
The NodeContext is shared across all nodes, providing logging and state management.
public class NodeContext {
// Store/retrieve shared data
public void set ( String key , Object value );
public < T > T get ( String key );
// Logging
public Logger getLogger ();
public OxyLogger getOxyjenLogger ();
// Metadata (e.g., graph name, execution ID)
public void setMetadata ( String key , Object value );
public < T > T getMetadata ( String key );
// Error handling
public ExceptionHandler getExceptionHandler ();
}
Think of it as: A shared notebook that all nodes can read/write to.
public String process ( String input , NodeContext ctx ) {
// Log what's happening
ctx . getLogger (). info ( "Processing: " + input );
// Store intermediate results
ctx . set ( "word_count" , input . split ( " " ). length );
// Share data between nodes
String previousResult = ctx . get ( "previous_output" );
return processedOutput ;
}
Complete Working Example
package examples ;
import io . oxyjen . core .*;
public class ContentPipeline {
public static void main ( String [] args ) {
// Step 1: Define your nodes
NodePlugin < String , String > validator = new ValidationNode ();
NodePlugin < String , String > processor = new ProcessingNode ();
NodePlugin < String , String > formatter = new FormatterNode ();
// Step 2: Build your graph
Graph pipeline = GraphBuilder . named ( "content-pipeline" )
. addNode ( validator )
. addNode ( processor )
. addNode ( formatter )
. build ();
// Step 3: Create execution context
NodeContext context = new NodeContext ();
context . set ( "max_length" , 100 );
// Step 4: Execute
Executor executor = new Executor ();
String result = executor . run ( pipeline , "Raw input text" , context );
System . out . println ( "Final output: " + result );
System . out . println ( "Word count: " + context . get ( "word_count" ));
}
}
// Example node implementations
class ValidationNode implements NodePlugin < String , String > {
@ Override
public String process ( String input , NodeContext ctx ) {
if ( input == null || input . isEmpty ()) {
throw new IllegalArgumentException ( "Input cannot be empty" );
}
ctx . getLogger (). info ( "Input validated" );
return input ;
}
}
class ProcessingNode implements NodePlugin < String , String > {
@ Override
public String process ( String input , NodeContext ctx ) {
String processed = input . toUpperCase (). trim ();
ctx . set ( "word_count" , processed . split ( " " ). length );
ctx . getLogger (). info ( "Text processed" );
return processed ;
}
}
class FormatterNode implements NodePlugin < String , String > {
@ Override
public String process ( String input , NodeContext ctx ) {
Integer maxLength = ctx . get ( "max_length" );
String formatted = input . length () > maxLength
? input . substring ( 0 , maxLength ) + "..."
: input ;
ctx . getLogger (). info ( "Text formatted" );
return formatted ;
}
}
My Vision for Oxyjen
Bring AI orchestration (LangChain/LangGraph style) to Java.
Build enterprise-first modules: LLM agents, Audit tools, Secure complex Workflow Engine.
Focus on performance, security, and observability .
I'm building this to learn java in a much deeper way.
RAG support - Vector databases, embeddings, document loaders
Cost management - Budgets, limits, usage tracking
Enterprise features - Audit logs, RBAC, compliance
Multi-tenancy - Isolate data between users/orgs
Circuit breakers - Fail fast when services are down
Token counting & cost tracking
< repositories >
< repository >
< id >jitpack.io</ id >
< url >https://jitpack.io</ url >
</ repository >
</ repositories >
Add dependency:
< dependency >
< groupId >com.github.11divyansh</ groupId >
< artifactId >Oxyjen</ artifactId >
< version >v0.4.0</ version >
</ dependency >
Gradle
repositories {
maven { url ' https://jitpack.io ' }
}
dependencies {
implementation ' com.github.11divyansh:Oxyjen:v0.4.0 '
}
Build from Source
git clone https://github.com/11divyansh/OxyJen.git
cd OxyJen
mvn clean install
After installation, verify by importing:
import io . oxyjen . core .*;
import io . oxyjen . llm .*;
import io . oxyjen . tools .*;
import io . oxyjen . graph .*;
About
Built with ❤️ by Divyansh Bhatt - a BTech CS student who believes Java deserves world-class AI tooling.
This started as a learning project, but I'm committed to making it production-ready. I know this is not big yet, but lets make it valuable.
Star this repo to follow the journey and be a part of it
Suggest features via Discussions
Contribute code or documentation
Share on Twitter/LinkedIn if you find it useful
** Watch for updates on v0.5 progress!**
Apache 2.0 (open-source, enterprise-friendly)
OxyJen is an open-source Java framework for orchestrating LLM workloads with graph-style execution, context-aware memory, and deterministic retry/fallback. It treats LLMs as native nodes (not helper utilities), allowing developers to build multi-step AI pipelines that integrate cleanly with existing Java code.
Apache-2.0 license
Contributing
There was an error while loading. Please reload this page .
2
forks
Report repository
Releases
5
OxyJen v0.5.1: Deterministic DAG Runtime for AI Workflows in Java
Latest
Jun 8, 2026
+ 4 releases
Packages
0
There was an error while loading. Ple

[truncated]
