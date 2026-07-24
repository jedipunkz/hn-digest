---
source: "https://github.com/samchon/lint-plugin-evidence"
hn_url: "https://news.ycombinator.com/item?id=49042719"
title: "Show HN: Evidence Graph – type checking for the specs your AI agent implements"
article_title: "GitHub - samchon/lint-plugin-evidence: Evidence graph for the AI coding era: declare provenance with a JSDoc `@evidence` tag, resolve it against Markdown sections, OpenAPI operations, and TypeScript symbols, and turn every missing or dangling citation into a real compile error under ttsc. · GitHub"
author: "autobe"
captured_at: "2026-07-24T23:56:28Z"
capture_tool: "hn-digest"
hn_id: 49042719
score: 1
comments: 0
posted_at: "2026-07-24T23:09:12Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Evidence Graph – type checking for the specs your AI agent implements

- HN: [49042719](https://news.ycombinator.com/item?id=49042719)
- Source: [github.com](https://github.com/samchon/lint-plugin-evidence)
- Score: 1
- Comments: 0
- Posted: 2026-07-24T23:09:12Z

## Translation

タイトル: HN の表示: 証拠グラフ – AI エージェントが実装する仕様の型チェック
記事のタイトル: GitHub - samchon/lint-plugin-evidence: AI コーディング時代の証拠グラフ: JSDoc `@evidence` タグで来歴を宣言し、Markdown セクション、OpenAPI 操作、および TypeScript シンボルに対してそれを解決し、すべての欠落または未解決の引用を ttsc での実際のコンパイル エラーに変換します。 · GitHub
説明: AI コーディング時代の証拠グラフ: JSDoc の「@evidence」タグで来歴を宣言し、Markdown セクション、OpenAPI 操作、および TypeScript シンボルに対して解決し、すべての欠落または未解決の引用を ttsc での実際のコンパイル エラーに変換します。 - samchon/lint-plugin-evidence

記事本文:
GitHub - samchon/lint-plugin-evidence: AI コーディング時代の証拠グラフ: JSDoc `@evidence` タグで来歴を宣言し、Markdown セクション、OpenAPI 操作、および TypeScript シンボルに対してそれを解決し、すべての欠落または未解決の引用を ttsc での実際のコンパイル エラーに変換します。 · GitHub
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
あなたは切り替えます

d アカウントを別のタブまたはウィンドウに表示します。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
サムチョン
/
lint-プラグイン-証拠
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
149 コミット 149 コミット .agents/ スキル .agents/ スキル .github .github .vscode .vscode config config パッケージ/ 証拠パッケージ/ 証拠テスト/ テスト証拠テスト/ テスト証拠 .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md ライセンス ライセンスREADME.md README.md og.jpg og.jpg package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml prettier.config.js prettier.config.js すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング時代の証拠グラフ: ゴール モードのガードレール。
あなたの仕様はコンパイルエラーになりました。
Claude Code または Codex が無人で動作する場合、要件をスキップしても「完了」と報告されることがあります。 Evidence Graph では、構成されたすべての要件が、それを満たしていると主張するコード、テスト、またはドキュメントからの明示的な承認を要求します。
すべての確認では、正確なターゲットを指定し、それが適用される理由が説明されます。コンパイラは、その理由が正しいかどうかを判断しません。コンパイラは、エージェントに具体的な要求を強制します。捏造された理由は、もはやもっともらしい差分の中に隠れることはできません。それは宣言とそれが矛盾する証拠の横にあります。
エージェントは依然として嘘をつくことができます。省略によって嘘をつくことはできません。
Complete : 構成されたすべての義務が満たされていない場合、ビルドは失敗します。
Tested : 選択されたすべてのエクスポートは、テストによって名前で要求されます。
文書化: 意思決定とコードは明示的に接続されたままになります。
正直：「完了」には目標と理由がつきものです。
誠実さ : 目標を超えて存続する引用はありません。
/**
* @evidence do

cs/discount.md#coupon-stacking このルールで定義された組み合わせ制限をレンダリングします。
*/
エクスポート関数 CouponStacking Notice () {
return < p > 1 つの販売者クーポンと 1 つのプラットフォーム クーポンを組み合わせることができます。 </ p > ;
}
@evidence 引用がないと、次のビルドが停止します。
$ npx ttsc
エラー TS16411: [証拠/グラフ] の確認応答がありません
' docs/discount.md#coupon-stacking ' (docs/discount.md:3 の Markdown H2 ' Coupon Stacking ')
請求項 1 の参照 1 (マークダウン、記号: h2、h3)。
選択した typescript ホストに「 @evidence docs/discount.md#coupon-stacking <reason> 」を追加します
この申し立ての内容、または「 @evidenceExclude docs/discount.md#coupon-stacking <reason> 」の場合
この主張では意図的にそれを使用していません。
1 件のエラーが見つかりました。
npm install -D typescript ttsc @ttsc/lint
npm install -D @samchon/lint-plugin-evidence
これは @ttsc/lint の lint プラグインです。 ESLint を備えたストック tsc ではなく、 ttsc 上で実行されます。ビルドで ttsc がまだ実行されていない場合は、最初にそのツールチェーンを採用してください。
最初のビルドには数分かかる場合があります。ルールを一度 lint バイナリにリンクし、後のビルドでそれを再利用します。
// lint.config.ts
"@ttsc/lint" からタイプ {ITtscLintConfig} をインポートします。
import { 証拠 , type IEvidenceGraphConfig } from "@samchon/lint-plugin-evidence" ;
const グラフ : IEvidenceGraphConfig = {
クレーム：[
{
タイプ: "タイプスクリプト" 、
ファイル: [ "src/components/**/*.tsx" ] 、
記号:「関数」、
参照: {
タイプ: "マークダウン" 、
ファイル: [ "docs/**/*.md" ] 、
記号: [ "h2" 、 "h3" ] 、
} 、
} 、
]、
} ;
デフォルトのエクスポート {
プラグイン: {
「証拠」：証拠、
} 、
ルール: {
"証拠/グラフ" : [ "エラー" , グラフ ] ,
"証拠/文書化" : "エラー" ,
"証拠/単数形" : "エラー" ,
} 、
ITtscLintConfig を満たします。
lint.config.ts にプラグインを登録し、証拠/グラフ ルールのオプションとしてグラフ宣言を渡します。このグラフは

ds を 1 つの文として説明します。src の下の React コンポーネントはドキュメントを実装すると主張しているため、ドキュメントの下のすべての H2 および H3 セクションはコンポーネントによって引用される必要があります。
証拠/グラフはプロジェクト スコープであるため、そのエントリにはファイル セレクターが含まれていてはなりません。ホストは、そうするものを拒否します。必要に応じて、ファイル ルールのスコープを独自のエントリに設定します。
すべての ttsc ビルド、すべての --noEmit チェック、およびすべての ttsx 実行で違反が表面化します。これらは型エラーと同じストリームに到着します。個別の CI ジョブはありません。
const グラフ : IEvidenceGraphConfig = {
クレーム：[
// 1. 機能ドキュメントは要件に基づいて作成されます
{
タイプ: "マークダウン" 、
ファイル: [ "docs/features/**/*.md" ] 、
参照: {
タイプ: "マークダウン" 、
ファイル: [ "docs/requirements/**/*.md" ] 、
記号: [ "h2" 、 "h3" ] 、
} 、
} 、
// 2. コンポーネントはフィーチャ ルールを実装します
{
タイプ: "タイプスクリプト" 、
ファイル: [ "src/components/**/*.tsx" ] 、
記号:「関数」、
参照: {
タイプ: "マークダウン" 、
ファイル: [ "docs/features/**/*.md" ] 、
記号: [ "h2" 、 "h3" ] 、
} 、
} 、
// 3. テストはフィーチャ ルールとコンポーネントを検証します
{
タイプ: "タイプスクリプト" 、
ファイル: [ "test/features/**/*.ts" ] ,
記号:「関数」、
参照: [
{
タイプ: "マークダウン" 、
ファイル: [ "docs/features/**/*.md" ] 、
記号: [ "h2" 、 "h3" ] 、
} 、
{
タイプ: "タイプスクリプト" 、
ファイル: [ "src/components/**/*.tsx" ] 、
記号:「関数」、
} 、
]、
} 、
]、
} ;
グラフは 1 つのクレーム配列であり、すべてのクレームと参照のペアは独立した義務です。
Markdown は Markdown を要求できます。機能ドキュメントは、その構築に基づいて構築されるすべての要件を認識する必要があります。
すべてのフィーチャ ルールは React コンポーネントによって引用される必要があります。コンポーネントミラーを持たないルールは、そのルールの名前を指定するコンパイルエラーです。
参照配列は要素ごとに 1 つの義務があります。テストでは、すべての機能ルールを検証し、エクスポートされたすべてのコンポーネントを要求する必要があり、一方が他方の CI を借用する義務は決してありません。

ステーション。
種類
シンボル値
デフォルト
「マークダウン」
「ファイル」、「h1」、「h2」、「h3」、「h4」
["ファイル"、"h1"、"h2"、"h3"、"h4"]
「闊歩する」
シンボルプロパティはありません。パスの下のすべての操作が選択されます
あらゆる操作
「タイプスクリプト」
「タイプ」、「関数」、「プロパティ」
クレームの場合は 3 つすべて、参照の場合は「type」
TypeScript の場合、「type」はエクスポートされたインターフェイス、タイプ エイリアス、および名前空間を選択します。 「関数」はエクスポートされた呼び出し可能オブジェクトを選択します。 "property" は、エクスポートされた型レベルのシンボルと、モジュールまたは名前空間のスコープでエクスポートされた const 、 let 、および var 宣言によって宣言されたプロパティを選択します。矢印または関数式で初期化された const は関数のままですが、他のすべての変数はプロパティです。修飾された ID はその所有者を保持します。 Orders.Input.id は Orders.Input の下のプロパティであり、 Orders.state は名前空間データです。
アンビエント名前空間メンバーは、TypeScript の暗黙的なエクスポート セマンティクスに従います。エクスポートされたオブジェクトと配列のバインディング パターンは、各ローカル バインディング リーフをプロパティとして公開します。型専用の名前空間エイリアスは、名前空間データや呼び出し可能オブジェクトを公開せずに、そのパブリック型空間の子孫とそのプロパティを公開します。
参照のシンボルは、1 つの義務がカバーする証拠単位を選択し、配列は 2 番目の義務を作成せずにその単位セットを拡張します。ユニットは階層を保持します。Markdown ファイルにはその見出しアウトラインが含まれ、TypeScript インターフェイスまたはオブジェクト タイプにはそのプロパティが含まれ、名前空間にはネストされたすべてのパブリック ユニットが含まれます。ターゲットは、それ自体と、選択されたすべての子孫を認識します。祖先は、それ自身の種類がセレクターから省略された場合でもアドレス指定可能なままであるため、シンボル: "property" は 1 つの @evidence IShoppingSale ... でカバーできます。
クレームのシンボルは、反対側にも同じセレクターを使用します。これにより、@evidence タグをホストできるシンボルの種類が制限されます。ネームスパ

ces はタイプ ホスト、エクスポートされたデータ変数はプロパティ ホスト、混合変数ステートメントはその常駐型のいずれかをホストできます。文書化されたデフォルトを受け入れるには、どちらかのセレクターを省略します。
Swagger は参照専用です。宣言をホストすることはできず、シンボル セレクターもありません。正規化されたドキュメントのパス オブジェクトでの各操作は 1 つの独立した義務です。
すべての Markdown または TypeScript ファイルのプロパティは、正規表現ではなく、プロジェクトに関連する glob パターンを受け取ります。 * は 1 つのパス セグメント内に一致し、** はセグメントを横断し、? 1 つの文字と一致します。 docs などの裸のディレクトリは、その子孫を選択しません。サブツリーの docs/** を書き込みます。
docs/**/*.md は、 docs の下にあるすべてのドキュメントを選択します。
backend/src/**/*.ts は、すべてのバックエンド ソース ファイルを選択します。
frontend/src/components/**/*.tsx はすべての React コンポーネントを選択します。
test/features/**/*.ts は、すべての機能テスト関数を選択します。
TypeScript 参照はその母集団を 3 つの方法で選択し、その選択によってそのユニットがどのようにアドレス指定されるかが決まります。
// src/contracts の下にあるエクスポートされたすべての型は、独自の名前でアドレス指定されます
{ タイプ : "typescript" 、ファイル : [ "src/contracts/**" ] }
// エントリが公開するすべてのもの。そのエントリからのアクセサー パスによってアドレス指定されます。
{ タイプ : "typescript" 、ファイル : "src/sdk/index.ts" }
// 消費者がインストールするパッケージの場合も同様
{ タイプ : "typescript" 、パッケージ : "@ORGANIZATION/PROJECT-api" }
files と file は相互に排他的であり、ローカル参照ではそれらのいずれかを設定する必要があります。暗黙的なプロジェクト エントリはありません。
エントリで選択された母集団は、宣言ファイルの綴り方ではなく、コンシューマーが到達する方法でアドレス指定されます。export * as function はパス セグメントをネストし、export * from は 1 をフラット化し、export { A as B } はシンボルを B としてアドレス指定します。これが、api.function.questions.get に名前を付けることができる理由です。アイデンティティは依然として

ファイルを宣言するため、エントリが 2 つのパスを通じて公開するシンボルは 2 つのアドレスに応答する 1 つのユニットであり、2 回支払う必要はなく 1 回確認されます。
パッケージの作成は、ttsc プログラムからではなくディスクから読み取られます。これが重要です。何もインポートされないシンボルは、定義上プログラムに存在せず、まさに義務が名前を付ける必要があるシンボルです。 file または files がない場合、パッケージの宣言エントリは母集団であり、そのエクスポート マップの type 条件、次に typeVersions 、次に type または testings を通じて解決されます。決して main ではありません。これは、引用で対処できる宣言ではなく、コンシューマが実行する JavaScript を指定します。 files の場合、グロブはパッケージ相対です。
パッケージ参照の義務セットは、それを公開した人に属します。エクスポートを追加するマイナー リリースでは義務が追加されるため、バージョンを固定するか、母集団が自分のものではない場合は選択を絞り込みます。
Swagger 参照は、その単一のファイル プロパティを通じて、正確に 1 つのドキュメントを所有します。
const グラフ : IEvidenceGraphConfig = {
クレーム：[
{
タイプ: "タイプスクリプト" 、
ファイル: [ "src/controllers/**/*.ts" ] 、
参照: {
タイプ: "闊歩" 、
ファイル: "api/openapi.yaml" 、
} 、
} 、
]、
} ;
ファイルは、プロジェクトに関連する 1 つの正確な p です。

[切り捨てられた]

## Original Extract

Evidence graph for the AI coding era: declare provenance with a JSDoc `@evidence` tag, resolve it against Markdown sections, OpenAPI operations, and TypeScript symbols, and turn every missing or dangling citation into a real compile error under ttsc. - samchon/lint-plugin-evidence

GitHub - samchon/lint-plugin-evidence: Evidence graph for the AI coding era: declare provenance with a JSDoc `@evidence` tag, resolve it against Markdown sections, OpenAPI operations, and TypeScript symbols, and turn every missing or dangling citation into a real compile error under ttsc. · GitHub
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
samchon
/
lint-plugin-evidence
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
149 Commits 149 Commits .agents/ skills .agents/ skills .github .github .vscode .vscode config config packages/ evidence packages/ evidence tests/ test-evidence tests/ test-evidence .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md og.jpg og.jpg package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml prettier.config.js prettier.config.js View all files Repository files navigation
The evidence graph for the AI coding era: the guardrail for goal mode.
Your spec is now a compile error.
When Claude Code or Codex works unattended, it can skip a requirement and still report "done." Evidence Graph makes every configured requirement demand an explicit acknowledgement from the code, test, or document that claims to satisfy it.
Every acknowledgement names the exact target and states why it applies. The compiler does not decide whether that reason is true—it forces the agent to commit to a concrete claim. A fabricated reason can no longer hide inside a plausible diff; it sits beside the declaration and evidence it contradicts.
An agent can still lie. It cannot lie by omission:
Complete : every configured obligation is accounted for, or the build fails.
Tested : every selected export is claimed by a test, by name.
Documented : decisions and code stay explicitly connected.
Honest : "done" comes with a target and a reason.
Integrity : no citation outlives its target.
/**
* @evidence docs/discount.md#coupon-stacking Renders the combination limit defined by this rule.
*/
export function CouponStackingNotice ( ) {
return < p > One seller coupon and one platform coupon may be combined. </ p > ;
}
Without the @evidence citation, the next build stops:
$ npx ttsc
error TS16411: [evidence/graph] Missing acknowledgement for
' docs/discount.md#coupon-stacking ' (Markdown H2 ' Coupon Stacking ' at docs/discount.md:3)
in Claim 1 reference 1 (markdown, symbols: h2, h3).
Add ' @evidence docs/discount.md#coupon-stacking <reason> ' to a selected typescript host
of this claim, or ' @evidenceExclude docs/discount.md#coupon-stacking <reason> ' when
this claim intentionally does not use it.
Found 1 error.
npm install -D typescript ttsc @ttsc/lint
npm install -D @samchon/lint-plugin-evidence
This is a lint plugin for @ttsc/lint . It runs on ttsc , not on stock tsc with ESLint. If your build does not run ttsc yet, adopt that toolchain first.
The first build can take several minutes; it links the rule into the lint binary once, and later builds reuse it.
// lint.config.ts
import type { ITtscLintConfig } from "@ttsc/lint" ;
import { evidence , type IEvidenceGraphConfig } from "@samchon/lint-plugin-evidence" ;
const graph : IEvidenceGraphConfig = {
claims : [
{
type : "typescript" ,
files : [ "src/components/**/*.tsx" ] ,
symbol : "function" ,
reference : {
type : "markdown" ,
files : [ "docs/**/*.md" ] ,
symbol : [ "h2" , "h3" ] ,
} ,
} ,
] ,
} ;
export default {
plugins : {
"evidence" : evidence ,
} ,
rules : {
"evidence/graph" : [ "error" , graph ] ,
"evidence/documented" : "error" ,
"evidence/singular" : "error" ,
} ,
} satisfies ITtscLintConfig ;
Register the plugin in lint.config.ts and pass the graph declaration as the option of the evidence/graph rule. This graph reads as one sentence: the React components under src claim to implement the docs, so every H2 and H3 section under docs must be cited by a component.
evidence/graph is project-scoped, so its entry must have no files selector; the host rejects one that does. Scope a file rule in its own entry when you need to.
Violations surface in every ttsc build, every --noEmit check, and every ttsx run. They arrive in the same stream as type errors. No separate CI job.
const graph : IEvidenceGraphConfig = {
claims : [
// 1. feature documents build on the requirements
{
type : "markdown" ,
files : [ "docs/features/**/*.md" ] ,
reference : {
type : "markdown" ,
files : [ "docs/requirements/**/*.md" ] ,
symbol : [ "h2" , "h3" ] ,
} ,
} ,
// 2. components implement the feature rules
{
type : "typescript" ,
files : [ "src/components/**/*.tsx" ] ,
symbol : "function" ,
reference : {
type : "markdown" ,
files : [ "docs/features/**/*.md" ] ,
symbol : [ "h2" , "h3" ] ,
} ,
} ,
// 3. tests verify the feature rules and the components
{
type : "typescript" ,
files : [ "test/features/**/*.ts" ] ,
symbol : "function" ,
reference : [
{
type : "markdown" ,
files : [ "docs/features/**/*.md" ] ,
symbol : [ "h2" , "h3" ] ,
} ,
{
type : "typescript" ,
files : [ "src/components/**/*.tsx" ] ,
symbol : "function" ,
} ,
] ,
} ,
] ,
} ;
A graph is one claims array, and every claim-reference pair is an independent obligation:
Markdown can claim Markdown. The feature documents must acknowledge every requirement they build on.
Every feature rule must be cited by a React component; a rule no component mirrors is a compile error naming that rule.
A reference array is one obligation per element. The tests must verify every feature rule and claim every exported component, never one obligation borrowing the other's citation.
Kind
symbol values
Default
"markdown"
"file" , "h1" , "h2" , "h3" , "h4"
["file", "h1", "h2", "h3", "h4"]
"swagger"
No symbol property; every operation under paths is selected
every operation
"typescript"
"type" , "function" , "property"
all three for claims, "type" for references
For TypeScript, "type" selects exported interfaces, type aliases, and namespaces. "function" selects exported callables. "property" selects properties declared by exported type-level symbols and exported const , let , and var declarations at module or namespace scope; a const initialized with an arrow or function expression remains a function, while every other variable is a property. Qualified identities preserve their owner: Orders.Input.id is a property below Orders.Input , while Orders.state is namespace data.
Ambient namespace members follow TypeScript's implicit export semantics. Exported object and array binding patterns expose each local binding leaf as a property. A type-only namespace alias exposes its public type-space descendants and their properties without exposing namespace data or callables.
A reference's symbol selects the evidence units one obligation covers, and an array widens that unit set without creating a second obligation. The units retain their hierarchy: a Markdown file contains its heading outline, a TypeScript interface or object type contains its properties, and a namespace contains every nested public unit. A target acknowledges itself and every selected descendant. An ancestor remains addressable even when its own kind is omitted from the selector, so symbol: "property" can still be covered by one @evidence IShoppingSale ... .
A claim's symbol uses the same selector for the opposite side: it restricts which symbol kinds may host an @evidence tag. Namespaces are type hosts, exported data variables are property hosts, and a mixed variable statement can host either of its resident kinds. Omit either selector to accept its documented default.
Swagger is reference-only. It cannot host declarations and has no symbol selector: each operation under the normalized document's paths object is one independent obligation.
Every Markdown or TypeScript files property takes project-relative glob patterns, not regular expressions. * matches inside one path segment, ** crosses segments, and ? matches one character. A bare directory such as docs does not select its descendants; write docs/** for the subtree.
docs/**/*.md selects every document below docs .
backend/src/**/*.ts selects every backend source file.
frontend/src/components/**/*.tsx selects every React component.
test/features/**/*.ts selects every feature test function.
A TypeScript reference selects its population three ways, and the choice decides how its units are addressed.
// every exported type under src/contracts, addressed by its own name
{ type : "typescript" , files : [ "src/contracts/**" ] }
// everything the entry exposes, addressed by its accessor path from that entry
{ type : "typescript" , file : "src/sdk/index.ts" }
// the same, for a package a consumer installs
{ type : "typescript" , package : "@ORGANIZATION/PROJECT-api" }
files and file are mutually exclusive, and a local reference must set one of them; there is no implicit project entry.
An entry-selected population is addressed the way a consumer reaches it, not the way the declaring file spells it: export * as functional nests a path segment, export * from flattens one, and export { A as B } addresses the symbol as B . That is what makes api.functional.questions.get nameable. Identity still belongs to the declaring file, so a symbol an entry exposes through two paths is one unit answering to two addresses — acknowledged once rather than owed twice.
A package population is read from disk rather than from the ttsc program, which is the point: a symbol nothing imports is absent from the program by definition, and it is exactly the symbol an obligation needs to name. Without file or files , the package's declaration entry is the population, resolved through the types condition of its exports map, then typesVersions , then types or typings — never main , which names the JavaScript a consumer runs rather than the declarations a citation can address. With files , the globs are package-relative.
The obligation set of a package reference belongs to whoever publishes it. A minor release that adds exports adds obligations, so pin the version or narrow the selection when the population is not yours.
A Swagger reference owns exactly one document through its singular file property:
const graph : IEvidenceGraphConfig = {
claims : [
{
type : "typescript" ,
files : [ "src/controllers/**/*.ts" ] ,
reference : {
type : "swagger" ,
file : "api/openapi.yaml" ,
} ,
} ,
] ,
} ;
file is either one exact project-relative p

[truncated]
