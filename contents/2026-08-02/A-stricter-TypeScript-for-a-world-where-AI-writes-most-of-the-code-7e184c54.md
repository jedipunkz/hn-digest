---
source: "https://github.com/chitalian/Typescript-Ultra-Strict"
hn_url: "https://news.ycombinator.com/item?id=49146450"
title: "A stricter TypeScript for a world where AI writes most of the code"
article_title: "GitHub - chitalian/Typescript-Ultra-Strict · GitHub"
author: "MarcoDewey"
captured_at: "2026-08-02T17:53:46Z"
capture_tool: "hn-digest"
hn_id: 49146450
score: 1
comments: 0
posted_at: "2026-08-02T17:30:15Z"
tags:
  - hacker-news
  - translated
---

# A stricter TypeScript for a world where AI writes most of the code

- HN: [49146450](https://news.ycombinator.com/item?id=49146450)
- Source: [github.com](https://github.com/chitalian/Typescript-Ultra-Strict)
- Score: 1
- Comments: 0
- Posted: 2026-08-02T17:30:15Z

## Translation

タイトル: AI がコードの大部分を記述する世界のためのより厳密な TypeScript
記事タイトル: GitHub - chitalian/Typescript-Ultra-Strict · GitHub
説明: GitHub でアカウントを作成して、chitalian/Typescript-Ultra-Strict の開発に貢献します。

記事本文:
GitHub - キタリアン/Typescript-Ultra-Strict · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
キタリアン
/
Typescript-Ultra-Strict
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブラ

nches タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミットの例 例 パッケージ パッケージ スクリプト スクリプト .gitignore .gitignore LEARNINGS.md LEARNINGS.md README.md README.md package-lock.json package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
これは TypeScript-Ultra-Strict に対する私の提案です。AI がコードの大部分を記述する世界向けの、より厳密な TypeScript です。
このリポジトリは思考実験であり、生産ツールではありません。コンパイラは動作し、すべての例はパスしますが、セマンティックな型チェック、網羅性チェック、公開パッケージ、安定性の保証はありません。提案を議論できるほど具体的なものにするために存在します。まだ関心のあるものは何も構築しないでください。
ステータス: 実用的な MVP。 package/cli は tsus check と tsus build を実装し、packages/runtime は Option / Result / match と境界層を実装し、packages/plugin は webpack/Vite/esbuild/Bun アダプタを提供し、examples/ pass の下にある 29 のすべてのプロジェクト (npm install && npm run Examples) を提供します。 LEARNINGS.md には、ツールチェーンが存在する前に何が壊れたか、それがどのように設計を形成したか、および既知の制限が文書化されています。
TypeScript の初期の設計決定では、人間が記述して保守することを前提としていたため、人間の作業を楽にするために多くの機能が存在します。私たちは品質よりもスピードを重視しました。 AI がコードの大部分を記述することで、そのトレードオフを逆転させ、言語を Rust のように宣言的かつ検証可能にすることができます。
null と unknown を削除します。代わりに、Option<T> と Result<T, E> がファーストクラスの型になります。これが目玉です。以下の削除は、これらの代替品が存在する場合にのみ機能します。
クラスを削除します。データはプレーンなオブジェクトです。行動は機能の中に生きています。
パターンマッチングと変数の構造化に使用されるマクロ用の SDK を導入します。

Option 、 Result 、およびタグ付き共用体に対して。 (TC39 にはすでにパターン マッチングの提案がパイプラインにあります。これにより、言語の方向性が加速されます。)
既存のすべての DOM および npm API は、 unknown を返し、クラスをスローし、渡します。 Ultra Strict コードでは、これらを直接参照することはありません。境界では、チェックされたバインディング層が T | を変換します。未定義と T | null を Option<T> に挿入し、呼び出しを Result<T, E> にスローします。境界線の内側では、どこにいても厳しい規則が適用されます。
なぜ単なるeslintプリセットではないのでしょうか?
クラス、null、および var の禁止は、現在 tsconfig および lint ルールで近似できます。実際のコンパイラが必要なものが 2 つあります。
マクロSDK。これがこの提案の正直なコストです。TypeScript の型は完全に消去可能であるため、esbuild と swc は 1 回の高速パスで型を削除できます。マクロはそのプロパティを壊します。 AI がコードを記述する場合、型ストリップを簡単に保つことよりも、表現力とチェック可能性にビルド時間の予算を費やす必要があるでしょう。
境界層。オプション/結果ラッパーを生成するために型情報が必要です。
2008 年に Google が Chrome と V8 を出荷し、業界はそのランタイム上にエッジ コンピューティングを構築しました。 Cloudflare Workers は V8 分離を直接実行します。 Node、Deno、Bun のエコシステムはすべて、そこに遡ります。 JavaScript と ESM は、すでにあらゆる場所に導入されている基盤です。
TypeScript は、エコシステムですでに採用されている型付きレイヤーです。 Ultra-Strict は、全員に新しい言語への移行を求めるのではなく、AI モデルがすでに深く知っている言語のサブセットとマクロを組み合わせたもので、すでにどこにでも存在するランタイムを対象としています。
.tsus 拡張子 (TypeScript Ultra Strict) を持つファイルはすべて、厳格なルールに従って処理され、プレーンな JavaScript にコンパイルされます。 .ts ファイルと .tsus ファイルは 1 つのプロジェクト内に共存するため、採用はファイルごとに行われます。AI エージェントは次のことができます。

従来の .ts コードはそのまま動作し続けながら、.tsus に新しいモジュールを作成できます。 .ts から .tsus モジュールをインポートすることは問題なく機能します。 .ts (または任意の npm パッケージ) を .tsus にインポートすると、境界層が通過して型が書き換えられるため、null 許容の戻り値が Option<T> として到着し、スロー関数が Result<T, E> として到着します。
import { readFile } from "node:fs/promises" ; // 境界ラップ: 結果を返します
type ユーザー = { 名前 : 文字列 ;電子メール: オプション <文字列> } ;
const loadUser = async (path : string ) : Promise < Result < User , IoError > > => {
const raw = await readFile ( path , "utf8" ) ;
生で返します。マップ (parseUser) ;
} ;
// SDK のマクロ
constgreeting = match (user . email ) {
Some ( email ) = > ` ${ email } までご連絡ください ` 、
なし => "ファイルに電子メールがありません" ,
} ;
フレームワークの統合
package/plugin (@tsus/plugin) は、ホストごとにシン アダプターを備えた 1 つの変換コアであり、すべて CLI のプリプロセッサとチェッカーを共有します。
@tsus/plugin/webpack : ローダー。 Next.js では、 next.config.mjs の 1 つのルールにより、任意のページまたはコンポーネントから .tsus をインポートできるようになります ( example/27-nextjs 、レンダリングされた文字列を静的 HTML に出力する次のビルドで検証されています)。
@tsus/plugin/vite : Vite プラグイン。 vite.plugins 設定 (examples/28-astro) を通じて Astro、SvelteKit、Nuxt もカバーします。
@tsus/plugin/esbuild : esbuild プラグイン。 Bun のプラグイン API は esbuild 形式であるため、同じオブジェクトが Bun.plugin に登録されます。 1 行の bunfig.toml プリロードを使用すると、bun run はビルド ステップなしで .tsus インポートをネイティブに実行します (examples/29-bun)。
アダプターはインポート指定子を意図的にそのままにしておきます。ホスト バンドラーは独自のパイプラインを通じて .tsus を .tsus チェーンに解決し、各ホップが再度変換にヒットします。 tsus CLI はスタンドアロン パス (および tsus check --json 経由の CI パス) のままです。
tsus ビルドは次のようにコンパイルされます

JS、tsus check は発行せずに型チェックを行います。どちらも esbuild/swc プラグインとして実行されるため、既存のバンドラー設定は構成を変更せずに .tsus ファイルを取得します。
エディターのサポートは TypeScript 言語サービス プラグインとして出荷されるため、VS Code とすべての TS 対応エディターは、無料で診断、定義への移動、およびマクロ展開プレビューを取得できます。
診断はマシンファーストです。すべてのエラーには安定したコード、JSON 出力モード、および推奨される修正が含まれているため、エージェントは散文を解析せずにループ内で tsus check --json を使用できます。人間が判読できる出力は、同じデータのレンダリングです。
新しいパッケージ マネージャーや新しいレジストリはありません。 npm パッケージは境界層を介して動作し、公開された .tsus パッケージはコンパイルされた JS と .d.ts ファイルを同梱するため、消費者はソース言語の存在を知る必要がありません。
Examples/ はテスト スイートです。29 個の小さなプロジェクトであり、それぞれに何が起こるかを宣言する example.json が含まれています。 npm run サンプルはすべてをビルドして実行しますが、ドリフトで失敗します。
意図的に拒否されました (これらは、正しいコードでチェッカーに不合格になることで合格します)
1 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to chitalian/Typescript-Ultra-Strict development by creating an account on GitHub.

GitHub - chitalian/Typescript-Ultra-Strict · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
chitalian
/
Typescript-Ultra-Strict
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits examples examples packages packages scripts scripts .gitignore .gitignore LEARNINGS.md LEARNINGS.md README.md README.md package-lock.json package-lock.json package.json package.json View all files Repository files navigation
This is my proposal for TypeScript-Ultra-Strict: a stricter TypeScript for a world where AI writes most of the code.
This repo is a thought experiment, not a production tool. The compiler works and every example passes, but there is no semantic type-checking, no exhaustiveness checking, no published packages, and no stability promise. It exists to make the proposal concrete enough to argue with. Do not build anything you care about on it yet.
Status: working MVP. packages/cli implements tsus check and tsus build , packages/runtime implements Option / Result / match and the boundary layer, packages/plugin provides webpack/Vite/esbuild/Bun adapters, and all 29 projects under examples/ pass ( npm install && npm run examples ). LEARNINGS.md documents what broke before the toolchain existed, how that shaped the design, and the known limitations.
TypeScript's early design decisions assumed humans would write and maintain it, so a lot of features exist to make a human's life easier. We valued speed over quality. With AI writing most of our code, we can flip that tradeoff and force the language to be declarative and verifiable, like Rust.
Removes null and undefined . In their place, Option<T> and Result<T, E> become first-class types. This is the centerpiece: the removals below only work because these replacements exist.
Removes classes. Data is plain objects; behavior lives in functions.
Introduces an SDK for macros, used for pattern matching and variable destructuring over Option , Result , and tagged unions. (TC39 already has a pattern-matching proposal in the pipeline; this accelerates where the language is heading.)
Every existing DOM and npm API returns undefined , throws, and hands you classes. Ultra-Strict code never sees those directly. At the boundary, a checked bindings layer converts T | undefined and T | null into Option<T> , and throwing calls into Result<T, E> . Inside the boundary, the strict rules hold everywhere.
Why not just an eslint preset?
Banning classes, nulls, and var can be approximated with tsconfig and lint rules today. Two things need a real compiler:
The macro SDK. This is the honest cost of the proposal: TypeScript's types are fully erasable, which is why esbuild and swc can strip them in a single fast pass. Macros break that property. The bet is that with AI writing the code, we should spend build-time budget on expressiveness and checkability rather than on keeping type-stripping trivial.
The boundary layer, which needs type information to generate the Option / Result wrappers.
Back in 2008, Google shipped Chrome and V8, and the industry built its edge compute on top of that runtime. Cloudflare Workers runs V8 isolates directly. Node, Deno, and Bun's ecosystem all trace back to it. JavaScript and ESM are the substrate that's already deployed everywhere.
TypeScript is the typed layer the ecosystem already adopted. Rather than asking everyone to move to a new language, Ultra-Strict is a subset-plus-macros of a language AI models already know deeply, targeting a runtime that's already everywhere.
Any file with a .tsus extension (TypeScript Ultra Strict) is processed under the strict rules and compiled to plain JavaScript. .ts and .tsus files coexist in one project, so adoption is per-file: an AI agent can write new modules in .tsus while the legacy .ts code keeps working untouched. Importing a .tsus module from .ts just works; importing .ts (or any npm package) into .tsus goes through the boundary layer, which rewrites the types so nullable returns arrive as Option<T> and throwing functions arrive as Result<T, E> .
import { readFile } from "node:fs/promises" ; // boundary-wrapped: returns Result
type User = { name : string ; email : Option < string > } ;
const loadUser = async ( path : string ) : Promise < Result < User , IoError > > => {
const raw = await readFile ( path , "utf8" ) ;
return raw . map ( parseUser ) ;
} ;
// macro from the SDK
const greeting = match ( user . email ) {
Some ( email ) = > `Reach me at ${ email } ` ,
None => "No email on file" ,
} ;
Framework integrations
packages/plugin (@tsus/plugin) is one transform core with a thin adapter per host, all sharing the CLI's preprocessor and checker:
@tsus/plugin/webpack : a loader. In Next.js, one rule in next.config.mjs makes .tsus importable from any page or component ( examples/27-nextjs , verified by next build emitting the rendered string into static HTML).
@tsus/plugin/vite : a Vite plugin, which also covers Astro, SvelteKit, and Nuxt through their vite.plugins config ( examples/28-astro ).
@tsus/plugin/esbuild : an esbuild plugin. Bun's plugin API is esbuild-shaped, so the same object registers with Bun.plugin ; with a one-line bunfig.toml preload, bun run executes .tsus imports natively with no build step ( examples/29-bun ).
The adapters deliberately leave import specifiers untouched: the host bundler resolves .tsus to .tsus chains through its own pipeline and each hop hits the transform again. The tsus CLI stays the standalone path (and the CI path, via tsus check --json ).
tsus build compiles to JS, tsus check type-checks without emitting. Both run as an esbuild/swc plugin so existing bundler setups pick up .tsus files without config changes.
Editor support ships as a TypeScript language service plugin, so VS Code and every TS-aware editor get diagnostics, go-to-definition, and macro expansion previews for free.
Diagnostics are machine-first: every error has a stable code, a JSON output mode, and a suggested fix, so an agent can consume tsus check --json in a loop without parsing prose. Human-readable output is a rendering of the same data.
No new package manager and no new registry. npm packages work through the boundary layer, and published .tsus packages ship compiled JS plus .d.ts files, so consumers don't need to know the source language existed.
examples/ is the test suite: 29 small projects, each with an example.json declaring what must happen. npm run examples builds and executes all of them and fails on any drift.
Rejected on purpose (these pass by failing the checker with the right code)
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
