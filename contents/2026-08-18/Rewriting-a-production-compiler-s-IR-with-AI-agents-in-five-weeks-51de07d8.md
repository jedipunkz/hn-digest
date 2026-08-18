---
source: "https://github.com/CommanderTvis/writing/tree/main/rr-truffle-rewrite"
hn_url: "https://news.ycombinator.com/item?id=49344173"
title: "Rewriting a production compiler's IR with AI agents in five weeks"
article_title: "writing/rr-truffle-rewrite at main · CommanderTvis/writing · GitHub"
image: "https://opengraph.githubassets.com/328116b2f13b3f0057b64bc8bc7a05e902da060792c080b3934a86397fb1869e/CommanderTvis/writing"
author: "CommanderTvis"
captured_at: "2026-08-18T12:25:37Z"
capture_tool: "hn-digest"
hn_id: 49344173
score: 2
comments: 0
posted_at: "2026-08-18T11:35:21Z"
tags:
  - hacker-news
  - translated
---

# Rewriting a production compiler's IR with AI agents in five weeks

- HN: [49344173](https://news.ycombinator.com/item?id=49344173)
- Source: [github.com](https://github.com/CommanderTvis/writing/tree/main/rr-truffle-rewrite)
- Score: 2
- Comments: 0
- Posted: 2026-08-18T11:35:21Z

## Translation

タイトル: AI エージェントを使用して実稼働コンパイラーの IR を 5 週間で書き直す
記事のタイトル: メインでの書き込み/rr-truffle-rewrite · CommanderTvis/書き込み · GitHub
説明: Iaroslav "Rick" Postovalov によるテクニカル ライティング。 GitHub でアカウントを作成して、CommanderTvis/ライティング開発に貢献してください。

記事本文:
メインでの書き込み/rr-truffle-rewrite · CommanderTvis/書き込み · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
コマンダーTvis
/
書く
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
その他のオプション ディレクトリアクション
歴史 歴史メイン ブレッドクラム
コピーパスのトップフォルダーとファイル
.. README.md README.md report.html report.html すべてのファイルを表示 README.md
概要 AI エージェントを使用して実稼働コンパイラーの IR を 5 週間で書き換える
十分に長く生き残った言語はすべて、最終的にはその言語の根幹を書き換えることになります。

コンパイラ。 C# は Roslyn でそれを行いました。 Rust は MIR を成長させました。 Kotlin は何年もかけて
K2 移行、私は JetBrains の Kotlin チーム内から参加しました。
新しい JVM バックエンドが依然として古いバックエンドを呼び出していたことを覚えています。
最も難しい部分については、インライナーを書き直す勇気のある人は誰もいなかったためです。の
パターンは常に同じです: 言語はプロトタイプの品質で生まれます
内部構造、そして要件が到着します — 2 番目のバックエンド、実際の IDE
サポート、シリアル化 - 元のツリーはそのように設計されていませんでした
運ぶ。大きな言語の場合、この書き換えは複数年にわたる複数チームの取り組みとなります。
この春、私はそのリライトを行いました
Rell 、Chromia の言語
分散型アプリケーション (dapps) を 5 週間で構築し、監督する
AIエージェント。この投稿
それはそれを形作った決断についてのものであり、それは私が下すものであり、私が行うものでした
彼らと一緒に生きていて、彼らなしでは私は試みることができなかったであろう部分について。
Rell は、Kotlin のような構文を備えたスタンドアロンの静的型付け言語です。
そのリレーショナル操作は SQL にコンパイルされます。以下は
プレイグラウンドのSQL
ドライランペイン:
エンティティ ユーザー {
キー名: テキスト;
変更可能な年齢: 整数;
}
エンティティポスト {
キーID: 整数;
インデックス作成者: ユーザー;
本文: テキスト;
}
// 各 @ 式は独自の SELECT にコンパイルされます。
クエリ main() {
// フィルター + 並べ替え + 射影。
val 大人 = ユーザー @* { .age >= 18 } ( @sort .name, .age );
// 集計: 著者ごとのカウント (GROUP BY)。
val post_counts = post @* {} ( @group .author.name, @sum 1 );
return (大人 = 大人、post_counts = post_counts);
}
これらは Rell のクエリ構文である @ 式です。中の二人
main() は次のようになります。
A00を選択します。 「名前」、A00。 「 c0.user 」からの「 age 」 A00
ここでA00。 「年齢」 >= ? A00までにご注文ください。 「名前」、A00。 「行ID」
A01を選択します。 " name " , COALESCE( SUM (?), 0 ) from " c0.post " A00
「 c0.user 」 A01 を A00 に参加させます。 「著者」 = A01。 」

行ID "
A01ごとにグループ化します。 「名前」はA01で注文します。 「名前」
2 番目のものが何をしたかに注目してください: .author.name はエンティティ参照を調べました。
そして結合はコンパイラから外れました。
これは、他のチェーンにはない Rell の一部です。あ
イーサリアム、Solana、または Move チェーンのコントラクトはキーと値のストレージを取得します。
そして、クエリに似たものはすべて、他の誰かによるオフチェーンの問題です。
Rell の状態はリレーショナル データベースであり、それをクエリするのが言語です
構文: コンパイル時にスキーマに対して型チェックが行われ、次のようにコンパイルされます。
SQL、コンセンサス境界内で実行されます。
私は 2026 年 2 月から Rell ソロを維持しています。本物のブロックチェーン
ネットワークはすべてのリリースで実行されます。 Rell はニッチな言語であり、
Rust や Kotlin よりも爆発範囲が広く、どちらよりもはるかにシンプルです。私はそうするだろう
今でも一年のほとんどをこのプロジェクトのために手作業で予算を立てています: もつれを解く
古いツリーをノードごとに書き換え、両方を維持しながらインタープリタを書き換えます。
ワールドは移行中に実行されます。
そのデータベースは無料ではなく購入されます。最速のチェーンは注文を計算します。
Chromia よりも規模が速く、Rell はそのクエリ レイヤーに料金を支払います。
スループット。同じような取引が言語を通じて行われます
実装。 LLVM はコンパイルに多くの時間を費やし、高速なコードを生成します。
これが Julia の取り決めです。関数への最初の呼び出しで代償が支払われます。
それをコンパイルすると、生成されたコードは C 速度で実行されます。 CPython
何も費やさず、ゆっくりと実行されます。 JVM、V8、.NET はその中間にあり、
進みながらコンパイルします。木の上を歩く通訳、それがレルです
CPython の最後にあります。即座に開始され、次のように実行されます。
操作ごとに木を歩くことから予想されるように、ゆっくりと。
これが Truffle が埋めるギャップであり、2 番目のバックエンドが使用された理由です。
書き直す価値があります。 Truffle は GraalVM (拡張 Java VM) 上のフレームワークです。
それは会議に書かれた通訳を必要とします

オンにして JIT を実行しましょう
コンパイラはそれを実行中のプログラムに特化させます。の代替案
同じ高速化は、コンパイラから JVM バイトコードを生成することであり、これは大幅な高速化です。
より大きくて繊細なものを所有すること。バックエンドを構築したかった
Truffle とコンパイラの出力モデルはこれをサポートできませんでした。それも
もっと欲しかったもの、つまりシリアル化をサポートできませんでした。
問題: 実行ロジックはコンパイラのツリー上に存在していました
書き換え前は、コンパイラーの出力モデル ( R_App ) は可変でしたが、
怠惰で、コンパイラ内部の番兵がいっぱいです。すべてのノードは独自のノードを持ちました
R_Expr の抽象 Evaluate(frame) として実行されます。 SQLの生成
機械 ( SqlGenContext 、 SqlBuilder ) はモデル内に存在しました
パッケージは、コンパイラが構築したのと同じクラス上にあります。 @式混合IR
1 つのファイルにランタイム エバリュエーター クラスを含むノード。ランタイムは同点だった
コンパイラでは、現在も実行されているすべての言語と同様に、
オリジナルの内部構造: 悪い決定によるものではなく、千の都合の良いものによる
もの。
結果: 2 番目のバックエンドは存在しません (実行は 1 つのバックエンドに組み込まれていました)
ツリーウォーク)、シリアル化なし (動作をシリアル化することはできません)、およびすべての
ブロックチェーンネットワーク内のノードは、すべてのアプリを再解析して再コンパイルします。
ソース、永遠に。
決定: どこまで徹底的に切り離すか?
私の努力のほとんどはここに費やされました。入力するのではなく、質問が 1 つだけです。 2
オプション。
保守的なオプションは、JetBrains で思い出した K2 の動きです。
新しい IR を作成し、難しい場所で古い実行コードを呼び出すようにします。
書き直す余裕がないチームにとっては合理的な選択です
すべて、そして実際にどの程度の大規模な移行が行われるかが重要です。それはまた、
何年も妥協して生きてきます。
根本的なオプション: IR を純粋なデータにする (ノード上でまったく動作しない)
そして散らばった解釈の断片をすべて新しく、網羅的なものとして再組み立てします。

ve
新しいツリー上でマッチングします。データベースのセマンティクスを含む:
@式、SQL生成、作成/更新/削除。
私は過激な選択肢を選択し、連載目標に落ち着いた
それ。コンパイラ側のコードに委譲するノードには書き込むものはありません
言語に依存しないバイナリ形式に変換します。委任は JVM コードであり、
コードは、まさにその形式で運ぶことができないものです。ハイブリッドはただのものではなかった
さらに悪い。それは何の表現もありませんでした。厳しい外部要件には価値がある
多くの点はここにあります: テイスト引数を制約に変換します。
出てきた形状: すべてのコンパイラが通過した後、1 つの解決ステップ
同梱されているアーキテクチャに関するドキュメントを引用すると、「すべての機能を強制します」
遅延フィールド、コンパイラと IDE の荷物を削除し、ライブ オブジェクトを置き換えます
整数インデックスを持つフラット配列への参照」により、不変が生成されます。
FlatBuffers にシリアル化する自己完結型 IR (バイナリ シリアル化)
format) であり、ランタイムが消費するのはこれだけです。コードベースでは、RR ツリーと呼ばれます。
フローチャート LR
SRC[ソース ファイル] -- "ANTLR4" --> S["S_ (AST)"]
S --> C["C_ (コンパイル、<br/>13 パス)"]
C --> R["R_ (コンパイラ モデル:<br/>可変、遅延、番兵)"]
R -- "resolve()" --> RR["RR_ (解決された IR:<br/>不変、フラット配列)"]
RR <-- FlatBuffers --> BIN[("シリアル化されたアプリ")]
RR --> INT["ツリーウォーキング インタプリタ"]
RR --> TF["Truffle バックエンド"]
読み込み中
メカニックスを簡単に。 IR ノードは、次の閉集合を持つ合計タイプです。
バリアント — 39 個の式、17 個のステートメント、16 個のデータベース式 —
インタプリタは、コンパイラがチェックするパターン マッチングです。
ドメインごとのロジックが個別のファイルに分割されているため、網羅性が高くなります。で
以下の Kotlin では、シールされたインターフェイスは、そのバリアントがすべてである sum 型です。
1 つのファイルで宣言され、コンパイラに認識されるデータ クラスは、
構造的等価性を備えた記録、および

(expr) { が X -> ... } の場合
それらの一致: バリアントを除外すると、コードはコンパイルされません。
前後の同じノードをリポジトリから軽くトリミングしたものです。
// 前: 実行はコンパイラのノード上で行われます
クラス R_IfExpr (
type : R_Type , // コンパイラ タイプ オブジェクト、コンパイラをドラッグします
プライベート val cond : R_Expr 、
プライベート値 trueExpr : R_Expr 、
プライベート値 falseExpr : R_Expr 、
): R_BaseExpr(タイプ) {
オーバーライド fun Evaluate0 ( フレーム : Rt_CallFrame ): Rt_Value {
val b = cond.evaluate(frame).asBoolean()
return ( if (b) trueExpr else falseExpr).evaluate(frame)
}
}
// 後: ノードはデータです...
シールされたインターフェイス RR_Expr {
値のタイプ: RR_Type
データクラス If (
override val type : RR_Type , // プレーンデータ、コンパイラ参照なし
値の条件: RR_Expr 、
val trueExpr : RR_Expr 、
val falseExpr : RR_Expr 、
): RR_Expr
// ...
}
// ...インタープリタが動作を所有し、バリアントごとに 1 つのアームが存在します
fun EvaluateExpr ( expr : RR_Expr 、フレーム : Rt_CallFrame ): Rt_Value = when (expr) {
は RR_Expr です。 If -> {
val cond = (evaluateExpr(expr.cond, Frame) as Rt_BooleanValue ).value
EvaluateExpr( if (cond) expr.trueExpr else expr.falseExpr, Frame)
}
// ...さらに 38 個の武器;コンパイラは欠落しているものを拒否します
}
相互参照は、フラットな種類ごとの配列への整数インデックスです。
(エンティティ、構造体、関数、クエリなど);シリアル化されたモジュールはのみをキャリーします
インデックス ベクトルとコンビニエンス マップは逆シリアル化時に再構築されます。の
これらすべてを構築するリゾルバーは、強制的に許可される唯一のコードです。
コンパイラの遅延フィールド。実行後はコンパイラの種類には何も影響しません
またまた。パブリック コンパイル エントリ ポイントは、RR ツリーではなく RR ツリーを返します。
コンパイラ モデルであるため、境界はライブラリの API サーフェスでもあります。
コンパイラとランタイムの境界は、ビルド システムのモジュールによって強制されます。
依存関係グラフ (矢印は「依存」を示します):
フローc

ハートBT
rrtree["rr-tree (IR データ、1.4K LOC)"] --> utils
rrser["rr-serialization"] --> rrtree
フロントエンド["フロントエンド (パーサー + コンパイラー)"] --> rrtree
フロントエンド --> ユーティリティ
rcore["ランタイムコア (値、型、SQL、stdlib)"] --> フロントエンド
rinterp["ランタイムインタープリター (5.4K LOC)"] --> rcore
rtruffle["ランタイム-トリュフ (7.8K LOC)"] --> rcore
rtruffle --> リンタープ
読み込み中
このグラフに関する正直なメモが 2 つあります。 runtime-core は依然として依存します
フロントエンドモジュール (アーキテクチャドキュメントには「今のところ」と書かれています)、そしてそのほとんど
依存関係は標準ライブラリです: stdlib 関数が宣言されています
コンパイラのライブラリ フレームワークを介して実装されるため、その実装は依然として
コンパイラモデルのタイプを処理します。実行モジュールはクリーンなものです。
インタプリタはフロントエンドから何もインポートしません。ハードモジュールの分割
それ自体はメインコミットの 2 週間後にランディングされました。そしてRR連載
リポジトリ内には本番コンシューマがありません。クライアントは必要なだけを要求します。
コンパイルされたプログラムをロードするには、rr-tree と rr-serialization が必要です。
他には何もありません。パーサーもコンパイラーもありません。
メインコミットは 4 月 21 日に到着しました: 641 ファイル、+32,050/-16,450。名前の変更
検出は、新規と移動の正直なストーリーを伝えます。値のタイプ、
ランタイム コンテキストとデータベース ドライバーの配管が認識されます。

[切り捨てられた]

## Original Extract

Technical writing by Iaroslav "Rick" Postovalov. Contribute to CommanderTvis/writing development by creating an account on GitHub.

writing/rr-truffle-rewrite at main · CommanderTvis/writing · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
CommanderTvis
/
writing
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
More options Directory actions
History History main Breadcrumbs
Copy path Top Folders and files
.. README.md README.md report.html report.html View all files README.md
Outline Rewriting a production compiler's IR with AI agents in five weeks
Every language that survives long enough ends up rewriting the guts of its
compiler. C# did it with Roslyn. Rust grew MIR. Kotlin spent years on the
K2 transition, one I participated in from inside JetBrains' Kotlin team,
where I remember the new JVM backend still calling into the old backend
for the hardest parts, because nobody dared rewrite the inliner. The
pattern is always the same: a language is born with prototype-quality
internals, and then requirements arrive — a second backend, real IDE
support, serialization — that the original trees were never designed to
carry. For big languages, this rewrite is a multi-year, multi-team effort.
This spring I did that rewrite for
Rell , Chromia's language for
building decentralized applications (dapps), in five weeks, directing
AI agents. This post
is about the decision that shaped it, which was mine to make and mine to
live with, and about the part I would not have attempted without them.
Rell is a standalone statically typed language with Kotlin-like syntax,
whose relational operations compile to SQL. The following runs in the
playground 's SQL
dry-run pane:
entity user {
key name: text;
mutable age: integer;
}
entity post {
key id: integer;
index author: user;
body: text;
}
// Each @-expression compiles to its own SELECT.
query main() {
// Filter + sort + projection.
val adults = user @* { .age >= 18 } ( @sort .name, .age );
// Aggregate: count per author (GROUP BY).
val post_counts = post @* {} ( @group .author.name, @sum 1 );
return (adults = adults, post_counts = post_counts);
}
Those are @-expressions, Rell's query syntax. The two in
main() come out as:
select A00. " name " , A00. " age " from " c0.user " A00
where A00. " age " >= ? order by A00. " name " , A00. " rowid "
select A01. " name " , COALESCE( SUM (?), 0 ) from " c0.post " A00
join " c0.user " A01 on A00. " author " = A01. " rowid "
group by A01. " name " order by A01. " name "
Note what the second one did: .author.name walked an entity reference,
and the join fell out of the compiler.
This is the part of Rell with no counterpart on other chains. A
contract on Ethereum, Solana or the Move chains gets key-value storage,
and anything resembling a query is somebody else's off-chain problem.
Rell's state is a relational database, and querying it is language
syntax: type-checked against the schema at compile time, compiled to
SQL, executed inside the consensus boundary.
I have maintained Rell solo since February 2026; real blockchain
networks run on every release. Rell is a niche language, with a smaller
blast radius than Rust or Kotlin, and much simpler than either. I would
still have budgeted most of a year for this project by hand: untangling
the old trees node by node, rewriting the interpreter, keeping both
worlds running mid-migration.
That database is bought, not free: the fastest chains compute orders of
magnitude faster than Chromia, and Rell pays for its query layer in
throughput. The same kind of trade runs through language
implementations. LLVM spends a lot of compile time and emits fast code,
which is the deal Julia takes: the first call to a function pays for
compiling it, and the code that comes out runs at C speed. CPython
spends none and executes slowly. The JVM, V8 and .NET sit in between,
compiling as they go. A tree-walking interpreter, which is what Rell
had, sits at the CPython end: it starts instantly and then runs about as
slowly as you would expect from walking a tree per operation.
That is the gap Truffle closes, and it is why the second backend was
worth a rewrite. Truffle is a framework on GraalVM (an extended Java VM)
that takes an interpreter written to its conventions and lets the JIT
compiler specialize it to the program being run. The alternative for the
same speedup is emitting JVM bytecode from the compiler, which is a much
larger and more delicate thing to own. I wanted the backend built on
Truffle, and the compiler's output model could not support one. It also
could not support something I wanted more: serialization.
The problem: execution logic lived on the compiler's tree
Before the rewrite, the compiler's output model ( R_App ) was mutable,
lazy, full of compiler-internal sentinels. Every node carried its own
execution, as an abstract evaluate(frame) on R_Expr . SQL generation
machinery ( SqlGenContext , SqlBuilder ) lived inside the model
package, on the same classes the compiler built. @-expressions mixed IR
nodes with runtime evaluator classes in one file. The runtime was tied
to the compiler the way it is in every language still running on its
original internals: not by a bad decision, but by a thousand convenient
ones.
The consequences: no second backend (execution was hard-wired into one
tree walk), no serialization (you cannot serialize behavior), and every
node in a blockchain network re-parses and re-compiles every app from
source, forever.
The decision: how radically to detach?
This is where most of my effort went: not typing, one question. Two
options.
The conservative option is the K2 move I remembered from JetBrains: build
a new IR, and let it call into the old execution code in the hard places.
It is the rational choice for a team that cannot afford to rewrite
everything, and it is how large migrations actually ship. It is also a
compromise you live with for years.
The radical option: make the IR pure data (no behavior on nodes at all)
and re-assemble every scattered piece of interpretation as new, exhaustive
matching over the new tree. Including the database semantics:
@-expressions, SQL generation, create/update/delete.
I chose the radical option, and the serialization goal is what settled
it. A node that delegates to compiler-side code has nothing to write
into a language-neutral binary format: the delegation is JVM code, and
code is exactly what the format cannot carry. The hybrid was not just
worse; it had no representation. A hard external requirement is worth a
lot here: it converts a taste argument into a constraint.
The shape that came out: after all compiler passes, one resolution step
that, quoting the architecture doc shipped with it, "forces every
lazy field, drops compiler and IDE baggage, and replaces live object
references with integer indices into flat arrays," producing an immutable,
self-contained IR that serializes to FlatBuffers (a binary serialization
format) and is the only thing the runtime consumes. In the codebase it is called the RR tree.
flowchart LR
SRC[source files] -- "ANTLR4" --> S["S_ (AST)"]
S --> C["C_ (compilation,<br/>13 passes)"]
C --> R["R_ (compiler model:<br/>mutable, lazy, sentinels)"]
R -- "resolve()" --> RR["RR_ (resolved IR:<br/>immutable, flat arrays)"]
RR <-- FlatBuffers --> BIN[("serialized app")]
RR --> INT["tree-walking interpreter"]
RR --> TF["Truffle backend"]
Loading
The mechanics, briefly. IR nodes are sum types with a closed set of
variants — 39 expression, 17 statement, 16 database-expression — and the
interpreter is pattern matching over them that the compiler checks for
exhaustiveness, with per-domain logic split into separate files. In the
Kotlin below, a sealed interface is a sum type whose variants are all
declared in one file and known to the compiler, a data class is a
record with structural equality, and when (expr) { is X -> ... } is
the match over them: leave a variant out and the code does not compile.
The same node before and after, lightly trimmed from the repo:
// before: execution lives on the compiler's node
class R_IfExpr (
type : R_Type , // compiler type object, drags the compiler in
private val cond : R_Expr ,
private val trueExpr : R_Expr ,
private val falseExpr : R_Expr ,
): R_BaseExpr(type) {
override fun evaluate0 ( frame : Rt_CallFrame ): Rt_Value {
val b = cond.evaluate(frame).asBoolean()
return ( if (b) trueExpr else falseExpr).evaluate(frame)
}
}
// after: the node is data...
sealed interface RR_Expr {
val type : RR_Type
data class If (
override val type : RR_Type , // plain data, no compiler references
val cond : RR_Expr ,
val trueExpr : RR_Expr ,
val falseExpr : RR_Expr ,
): RR_Expr
// ...
}
// ...and the interpreter owns the behavior, one arm per variant
fun evaluateExpr ( expr : RR_Expr , frame : Rt_CallFrame ): Rt_Value = when (expr) {
is RR_Expr . If -> {
val cond = (evaluateExpr(expr.cond, frame) as Rt_BooleanValue ).value
evaluateExpr( if (cond) expr.trueExpr else expr.falseExpr, frame)
}
// ...38 more arms; the compiler rejects a missing one
}
Cross-references are integer indices into flat per-kind arrays
(entities, structs, functions, queries…); serialized modules carry only
index vectors, and convenience maps are rebuilt on deserialization. The
resolver that builds all this is the only code allowed to force the
compiler's lazy fields; after it runs, nothing touches compiler types
again. The public compile entry point returns the RR tree rather than
the compiler model, so the boundary is also the library's API surface.
The compiler/runtime boundary is enforced by the build system's module
dependency graph (arrows read "depends on"):
flowchart BT
rrtree["rr-tree (IR data, 1.4K LOC)"] --> utils
rrser["rr-serialization"] --> rrtree
frontend["frontend (parser + compiler)"] --> rrtree
frontend --> utils
rcore["runtime-core (values, types, SQL, stdlib)"] --> frontend
rinterp["runtime-interpreter (5.4K LOC)"] --> rcore
rtruffle["runtime-truffle (7.8K LOC)"] --> rcore
rtruffle --> rinterp
Loading
Two honest notes on that graph. runtime-core still depends on the
frontend module ("for now", says the architecture doc), and most of that
dependency is the standard library: stdlib functions are declared
through the compiler's library framework, so their implementations still
handle compiler-model types. The execution modules are the clean ones;
the interpreter imports nothing from the frontend. The hard module split
itself landed two weeks after the main commit. And rr-serialization
has no production consumer inside the repo: a client that only wants
to load compiled programs needs rr-tree plus rr-serialization and
nothing else: no parser, no compiler.
The main commit landed on April 21: 641 files, +32,050/−16,450. Rename
detection tells the honest story of new versus moved. The value types,
the runtime contexts and the database-driver plumbing are recog

[truncated]
