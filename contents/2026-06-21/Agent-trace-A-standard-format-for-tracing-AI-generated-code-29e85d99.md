---
source: "https://github.com/cursor/agent-trace"
hn_url: "https://news.ycombinator.com/item?id=48616614"
title: "Agent-trace: A standard format for tracing AI-generated code"
article_title: "GitHub - cursor/agent-trace: A standard format for tracing AI-generated code. · GitHub"
author: "Garbage"
captured_at: "2026-06-21T10:18:51Z"
capture_tool: "hn-digest"
hn_id: 48616614
score: 2
comments: 0
posted_at: "2026-06-21T07:48:17Z"
tags:
  - hacker-news
  - translated
---

# Agent-trace: A standard format for tracing AI-generated code

- HN: [48616614](https://news.ycombinator.com/item?id=48616614)
- Source: [github.com](https://github.com/cursor/agent-trace)
- Score: 2
- Comments: 0
- Posted: 2026-06-21T07:48:17Z

## Translation

タイトル: Agent-trace: AI 生成コードをトレースするための標準形式
記事のタイトル: GitHub - カーソル/エージェント-トレース: AI によって生成されたコードをトレースするための標準形式。 · GitHub
説明: AI によって生成されたコードをトレースするための標準形式。 GitHub でアカウントを作成して、カーソル/エージェント トレースの開発に貢献してください。

記事本文:
GitHub - カーソル/エージェントトレース: AI が生成したコードをトレースするための標準形式。 · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
カーソル
/
エージェントトレース
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブランチ Tags Go

ファイルへ コード 「その他のアクション」メニューを開く フォルダーとファイル
10 コミット 10 コミット .claude .claude .cursor .cursor アセット アセット リファレンス リファレンス .gitignore .gitignore README.md README.md bun.lock bun.lock Index.ts Index.ts package.json package.json schemas.ts schemas.ts vercel.json vercel.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
バージョン : 0.1.0
ステータス : RFC
日付：2026年1月
Agent Trace は、AI によって生成されたコードをトレースするためのオープン仕様です。これは、バージョン管理されたコードベースで人間の作成者とともに AI の貢献を記録するためのベンダー中立の形式を提供します。
エージェントがより多くのコードを作成するにつれて、AI と人間から何が生じたのかを理解することが重要になります。この属性は、使用されるモデルと、関連するエージェントの会話の両方です。 Agent Trace は、この属性データを記録するためのオープンで相互運用可能な標準を定義します。
相互運用性 : 準拠したツールであればどれでもアトリビューション データの読み取りと書き込みが可能です。
粒度 : ファイルおよび行の粒度で使用されるモデルの属性をサポートします。
拡張性 : ベンダーは互換性を損なうことなくカスタム メタデータを追加できます。
人間とエージェントが読み取り可能 : アトリビューション データは、特別なツールを使用せずに読み取り可能です。
コードの所有権 : Agent Trace は、法的な所有権や著作権を追跡しません。
トレーニング データの出所 : どのトレーニング データが AI 出力に影響を与えたかは追跡しません。
品質評価 : AI の貢献が良いか悪いかは評価しません。
UI に依存しない : Agent Trace は特定のインターフェイスを必要としません。
期間
定義
貢献
コードの変更（追加、変更、削除）の単位です。
投稿者
貢献を生み出したエンティティ (人間または AI)
トレースレコード
投稿の起源を説明するメタデータ
貢献者のタイプ
種類
コード
説明
人間
人間
人間の開発者によって直接作成されたコード
AI
あい
AIによって生成されたコード
ミックス

編
混合された
人間が編集した AI 出力または AI が編集した人間のコード
不明
不明
起源は特定できません
5. アーキテクチャの概要
Agent Trace はデータ仕様であり、製品ではありません。アトリビューション データを記録する方法を定義します。ストレージメカニズムは実装によって定義されます。この仕様には、トレースがどこに存在するかについての意見はありません。
エージェント トレースの基本単位はトレース レコードです。
{
"$schema" : " https://json-schema.org/draft/2020-12/schema " ,
"$id" : " https://agent-trace.dev/schemas/v1/trace-record.json " ,
"title" : " エージェント トレース レコード " ,
"タイプ" : "オブジェクト" ,
"必須" : [ " バージョン " 、 " ID " 、 " タイムスタンプ " 、 " ファイル " ]、
"プロパティ" : {
「バージョン」: {
"タイプ" : "文字列" ,
"パターン" : " ^[0-9]+ \\ .[0-9]+ \\ .[0-9]+$ " ,
"description" : " Agent Trace 仕様のバージョン (例: '1.0.0') "
}、
「ID」: {
"タイプ" : "文字列" ,
"フォーマット" : "uuid" ,
"description" : " このトレース レコードの一意の識別子 "
}、
「タイムスタンプ」: {
"タイプ" : "文字列" ,
"format" : " 日時 " ,
"description" : " トレースが記録されたときの RFC 3339 タイムスタンプ "
}、
"VC" : {
"$ref" : " #/$defs/vcs " ,
"description" : " このトレースのバージョン管理システム情報 "
}、
「ツール」: {
"$ref" : " #/$defs/tool " ,
"description" : " このトレースを生成したツール "
}、
"ファイル" : {
"タイプ" : "配列" ,
「アイテム」: {
"$ref" : " #/$defs/file "
}、
"description" : " 属性付き範囲を持つファイルの配列 "
}、
「メタデータ」: {
"タイプ" : "オブジェクト" ,
"description" : " 実装固有またはベンダー固有のデータの追加メタデータ "
}
}、
"$defs" : {
"VC" : {
"タイプ" : "オブジェクト" ,
"必須" : [ "タイプ" , "リビジョン" ],
"プロパティ" : {
「タイプ」: {
"タイプ" : "文字列" ,
"enum" : [ " git " 、 " jj " 、 " hg " 、 " svn " ]、
"description" : " バージョン管理システムのタイプ "
}、
「リビジョン」: {
"タイプ" : " 文字列

「、
"description" : " リビジョン識別子 (例: git commit SHA、jj Change ID) "
}
}
}、
「ツール」: {
"タイプ" : "オブジェクト" ,
"プロパティ" : {
"名前" : { "タイプ" : " 文字列 " },
"バージョン" : { "タイプ" : " 文字列 " }
}
}、
「ファイル」: {
"タイプ" : "オブジェクト" ,
"必須" : [ " パス " 、 " 会話 " ]、
"プロパティ" : {
「パス」: {
"タイプ" : "文字列" ,
"description" : " リポジトリルートからの相対ファイルパス "
}、
「会話」: {
"タイプ" : "配列" ,
「アイテム」: {
"$ref" : " #/$defs/conversation "
}、
"description" : " このファイルに貢献した会話の配列 "
}
}
}、
[切り捨てられた]
{
"バージョン" : " 0.1.0 " 、
"id" : " 550e8400-e29b-41d4-a716-446655440000 " ,
"タイムスタンプ" : " 2026-01-23T14:30:00Z " ,
"VC" : {
"タイプ" : " git " ,
"リビジョン" : " a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0 "
}、
「ツール」: {
"名前" : " カーソル " ,
「バージョン」：「2.4.0」
}、
「ファイル」: [
{
"パス" : " src/utils/parser.ts " ,
「会話」: [
{
"url" : " https://api.cursor.com/v1/conversations/12345 " ,
「投稿者」: {
"タイプ" : " ai " 、
"model_id" : " anthropic/claude-opus-4-5-20251101 "
}、
「範囲」: [
{
"開始行" : 42 、
"end_line" : 67 、
"content_hash" : " murmur3:9f2e8a1b "
}
]、
「関連」: [
{
"タイプ" : "セッション" ,
"url" : " https://api.cursor.com/v1/sessions/67890 "
}
】
}
】
}、
{
"パス" : " src/utils/helpers.ts " ,
「会話」: [
{
"url" : " https://api.cursor.com/v1/conversations/12345 " ,
「投稿者」: {
"タイプ" : " ai " 、
"model_id" : " openai/gpt-4o "
}、
「範囲」: [
{
"開始行" : 10 、
「終了行」: 25
}
】
}
】
}
]、
「メタデータ」: {
「信頼度」 : 0.95 、
"dev.cursor" : {
"workspace_id" : " ws-abc123 "
}
}
}
6.3 帰属の範囲
範囲は会話ごとにグループ化され、会話レベルの投稿者のメタデータが使用されます。これにより、1 つの会話で多くの範囲が生成される場合、カーディナリティが減少します。

{
「ファイル」: [
{
"パス" : " src/utils.ts " ,
「会話」: [
{
"url" : " https://api.example.com/v1/conversations/abc " ,
「投稿者」: {
"タイプ" : " ai " 、
"model_id" : " anthropic/claude-sonnet-4-20250514 "
}、
「範囲」: [
{ "開始行" : 10 、 "終了行" : 25 },
{ "開始行" : 30 、 "終了行" : 45 },
{ "開始行" : 80 、 "終了行" : 95 }
】
}、
{
"url" : " https://api.example.com/v1/conversations/def " ,
"contributor" : { "type" : " ai " , "model_id" : " openai/gpt-4o " },
"範囲" : [{ "開始行" : 50 、 "終了行" : 52 }]
}
】
}
】
}
行番号は 1 から始まります。記録されたリビジョンでの範囲参照位置。
Agent Trace は、vcs フィールドを通じて複数のバージョン管理システムをサポートします。
// Git
{ "vcs" : { "タイプ" : " git " 、 "リビジョン" : " a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0 " } }
// Jujuku (リベース間の安定性のために変更 ID を使用)
{ "vcs" : { "タイプ" : " jj " 、 "リビジョン" : " kkmpptxz " } }
// マーキュリアル
{ "vcs" : { "タイプ" : " hg " 、 "リビジョン" : " a1b2c3d4e5f6 " } }
リビジョン フィールドの形式は VCS 固有です。
git : 40 文字の 16 進コミット SHA
jj : 変更 ID (修正/リベース操作全体で安定)
トレース内の行番号は、現在の位置ではなく、記録されたリビジョンの位置を指します。特定のコード行の所有権をクエリするには、次のようにします。
VCS の非難を使用して、行 N に最後に触れたリビジョンを見つけます。
そのリビジョンとファイルのトレースを検索します
行 N を含む範囲を検索します。
コードの動き全体でアトリビューションを追跡するには、範囲レベルでコンテンツ ハッシュを使用します。
{
「ファイル」: [
{
"パス" : " src/parser.ts " ,
「会話」: [
{
"url" : " https://api.example.com/v1/conversations/abc " ,
「投稿者」: {
"タイプ" : " ai " 、
"model_id" : " anthropic/claude-opus-4-5-20251101 "
}、
「範囲」: [
{
"開始行" : 10 、
"end_line" : 25 、
"content_hash" : " murmur3:9f2e8

a1b "
}
】
}
】
}
】
}
ハッシュは特定の範囲に適用されるため、コードがファイル内またはファイル間で移動する場合でも追跡できます。
モデル識別子は、models.dev の規則に従います。
{
「投稿者」: {
"タイプ" : " ai " 、
"model_id" : " anthropic/claude-opus-4-5-20251101 "
}
}
形式：プロバイダ/モデル名
各会話には URL フィールドと、関連するサブリソースにリンクするためのオプションの関連配列があります。
{
「ファイル」: [
{
"パス" : " src/api.ts " ,
「会話」: [
{
"url" : " https://api.example.com/v1/conversations/abc123 " ,
「投稿者」: {
"タイプ" : " ai " 、
"model_id" : " anthropic/claude-opus-4-5-20251101 "
}、
"範囲" : [{ "開始行" : 10 , "終了行" : 50 }],
「関連」: [
{
"タイプ" : "セッション" ,
"url" : " https://api.example.com/v1/sessions/xyz789 "
}、
{
"タイプ" : "プロンプト" ,
"url" : " https://api.example.com/v1/prompts/def456 "
}
】
}
】
}
】
}
7. 拡張性
メジャー バージョン: 必須フィールドへの重大な変更
マイナー バージョン: 追加的な変更 (新しいオプション フィールド)
メタデータ フィールドには、実装またはベンダー固有のデータを含めることができます。
{
「メタデータ」: {
「信頼度」 : 0.95 、
"post_processing_tools" : [ " prettier@3.0.0 " ],
"dev.cursor" : {
"workspace_id" : " ws-abc123 "
}
}
}
ベンダーはキーの衝突を避けるためにメタデータ内で逆ドメイン表記 (例: dev.cursor 、 com.github.copilot ) を使用する場合があります。
リファレンス実装は、リファレンス/ディレクトリに提供されており、エージェント トレースをコーディング エージェントと統合する方法を示しています。実装には以下が含まれます。
trace-store.ts : トレース レコードの読み取りと書き込みのためのストレージ層
trace-hook.ts : ファイル変更時の自動トレースキャプチャのためのフック統合
このリファレンスはカーソル コードまたはクロード コードの例ですが、パターンは任意の AI コーディング エージェントに適用できます。
{
"バージョン" : " 0.1.0 " 、
"id" : " 550e8400-e29b-41d4-a716-4466554

40000 " 、
"タイムスタンプ" : " 2026-01-25T10:00:00Z " ,
「ファイル」: [
{
"パス" : " src/app.ts " ,
「会話」: [
{
"投稿者" : { "タイプ" : " ai " },
"範囲" : [{ "開始行" : 1 , "終了行" : 50 }]
}
】
}
】
}
B. MIME タイプ
種類
MIME タイプ
トレースレコード
アプリケーション/vnd.agent-trace.record+json
C. よくある質問
トレースはどのように保存すればよいですか?
この仕様では、トレースの保存方法を意図的に定義していません。これは、ローカル ファイル、git メモ、データベース、またはその他のものである可能性があります。
リベースやマージコミットはどのように処理すればよいですか?
オープンソースでさまざまな実装が行われることを期待しています。これは将来の仕様に影響を与える可能性があります。フィードバックをお待ちしております。
エージェントがスクリプトを作成してコードを記述するとどうなりますか?
これは実装に委ねられます。この方法で生成されたコードは、引き続きエージェントに帰属する必要があります。たとえば、スクリプトの実行前後にファイルのスナップショットを作成し、git diff を使用してエージェントが何を追加したかを確認できます。
この仕様は CC BY 4.0 に基づいてリリースされています。
Agent Trace の形成にご協力いただいた次のパートナーに感謝します。
この仕様は GitHub で提案を受け付けています。
AI が生成したコードをトレースするための標準形式。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロペラ

[切り捨てられた]

## Original Extract

A standard format for tracing AI-generated code. Contribute to cursor/agent-trace development by creating an account on GitHub.

GitHub - cursor/agent-trace: A standard format for tracing AI-generated code. · GitHub
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
cursor
/
agent-trace
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
10 Commits 10 Commits .claude .claude .cursor .cursor assets assets reference reference .gitignore .gitignore README.md README.md bun.lock bun.lock index.ts index.ts package.json package.json schemas.ts schemas.ts vercel.json vercel.json View all files Repository files navigation
Version : 0.1.0
Status : RFC
Date : January 2026
Agent Trace is an open specification for tracing AI-generated code. It provides a vendor-neutral format for recording AI contributions alongside human authorship in version-controlled codebases.
As agents write more code, it's important to understand what came from AI versus humans. This attribution is both the models used as well as the related agent conversations. Agent Trace defines an open, interoperable standard for recording this attribution data.
Interoperability : Any compliant tool can read and write attribution data.
Granularity : Support attribution for models used at file and line granularity.
Extensibility : Vendors can add custom metadata without breaking compatibility.
Human & Agent Readable : Attribution data is readable without special tooling.
Code Ownership : Agent Trace does not track legal ownership or copyright.
Training Data Provenance : We don't track what training data influenced AI outputs.
Quality Assessment : We don't evaluate whether AI contributions are good or bad.
UI Agnostic : Agent Trace does not require any specific interface.
Term
Definition
Contribution
A unit of code change (addition, modification, or deletion)
Contributor
The entity that produced a contribution (human or AI)
Trace Record
Metadata describing a contribution's origin
Contributor Types
Type
Code
Description
Human
human
Code authored directly by a human developer
AI
ai
Code generated by AI
Mixed
mixed
Human-edited AI output or AI-edited human code
Unknown
unknown
Origin cannot be determined
5. Architecture Overview
Agent Trace is a data specification, not a product. It defines how to record attribution data. Storage mechanisms are implementation-defined. The spec is unopinionated about where traces live.
The fundamental unit of Agent Trace is the Trace Record :
{
"$schema" : " https://json-schema.org/draft/2020-12/schema " ,
"$id" : " https://agent-trace.dev/schemas/v1/trace-record.json " ,
"title" : " Agent Trace Record " ,
"type" : " object " ,
"required" : [ " version " , " id " , " timestamp " , " files " ],
"properties" : {
"version" : {
"type" : " string " ,
"pattern" : " ^[0-9]+ \\ .[0-9]+ \\ .[0-9]+$ " ,
"description" : " Agent Trace specification version (e.g., '1.0.0') "
},
"id" : {
"type" : " string " ,
"format" : " uuid " ,
"description" : " Unique identifier for this trace record "
},
"timestamp" : {
"type" : " string " ,
"format" : " date-time " ,
"description" : " RFC 3339 timestamp when trace was recorded "
},
"vcs" : {
"$ref" : " #/$defs/vcs " ,
"description" : " Version control system information for this trace "
},
"tool" : {
"$ref" : " #/$defs/tool " ,
"description" : " The tool that generated this trace "
},
"files" : {
"type" : " array " ,
"items" : {
"$ref" : " #/$defs/file "
},
"description" : " Array of files with attributed ranges "
},
"metadata" : {
"type" : " object " ,
"description" : " Additional metadata for implementation-specific or vendor-specific data "
}
},
"$defs" : {
"vcs" : {
"type" : " object " ,
"required" : [ " type " , " revision " ],
"properties" : {
"type" : {
"type" : " string " ,
"enum" : [ " git " , " jj " , " hg " , " svn " ],
"description" : " Version control system type "
},
"revision" : {
"type" : " string " ,
"description" : " Revision identifier (e.g., git commit SHA, jj change ID) "
}
}
},
"tool" : {
"type" : " object " ,
"properties" : {
"name" : { "type" : " string " },
"version" : { "type" : " string " }
}
},
"file" : {
"type" : " object " ,
"required" : [ " path " , " conversations " ],
"properties" : {
"path" : {
"type" : " string " ,
"description" : " Relative file path from repository root "
},
"conversations" : {
"type" : " array " ,
"items" : {
"$ref" : " #/$defs/conversation "
},
"description" : " Array of conversations that contributed to this file "
}
}
},
[truncated]
{
"version" : " 0.1.0 " ,
"id" : " 550e8400-e29b-41d4-a716-446655440000 " ,
"timestamp" : " 2026-01-23T14:30:00Z " ,
"vcs" : {
"type" : " git " ,
"revision" : " a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0 "
},
"tool" : {
"name" : " cursor " ,
"version" : " 2.4.0 "
},
"files" : [
{
"path" : " src/utils/parser.ts " ,
"conversations" : [
{
"url" : " https://api.cursor.com/v1/conversations/12345 " ,
"contributor" : {
"type" : " ai " ,
"model_id" : " anthropic/claude-opus-4-5-20251101 "
},
"ranges" : [
{
"start_line" : 42 ,
"end_line" : 67 ,
"content_hash" : " murmur3:9f2e8a1b "
}
],
"related" : [
{
"type" : " session " ,
"url" : " https://api.cursor.com/v1/sessions/67890 "
}
]
}
]
},
{
"path" : " src/utils/helpers.ts " ,
"conversations" : [
{
"url" : " https://api.cursor.com/v1/conversations/12345 " ,
"contributor" : {
"type" : " ai " ,
"model_id" : " openai/gpt-4o "
},
"ranges" : [
{
"start_line" : 10 ,
"end_line" : 25
}
]
}
]
}
],
"metadata" : {
"confidence" : 0.95 ,
"dev.cursor" : {
"workspace_id" : " ws-abc123 "
}
}
}
6.3 Attribution Ranges
Ranges are grouped by conversation, with contributor metadata at the conversation level. This reduces cardinality when one conversation produces many ranges.
{
"files" : [
{
"path" : " src/utils.ts " ,
"conversations" : [
{
"url" : " https://api.example.com/v1/conversations/abc " ,
"contributor" : {
"type" : " ai " ,
"model_id" : " anthropic/claude-sonnet-4-20250514 "
},
"ranges" : [
{ "start_line" : 10 , "end_line" : 25 },
{ "start_line" : 30 , "end_line" : 45 },
{ "start_line" : 80 , "end_line" : 95 }
]
},
{
"url" : " https://api.example.com/v1/conversations/def " ,
"contributor" : { "type" : " ai " , "model_id" : " openai/gpt-4o " },
"ranges" : [{ "start_line" : 50 , "end_line" : 52 }]
}
]
}
]
}
Line numbers are 1-indexed. Ranges reference positions at the recorded revision.
Agent Trace supports multiple version control systems through the vcs field:
// Git
{ "vcs" : { "type" : " git " , "revision" : " a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0 " } }
// Jujutsu (using change ID for stability across rebases)
{ "vcs" : { "type" : " jj " , "revision" : " kkmpptxz " } }
// Mercurial
{ "vcs" : { "type" : " hg " , "revision" : " a1b2c3d4e5f6 " } }
The revision field format is VCS-specific:
git : 40-character hex commit SHA
jj : Change ID (stable across amend/rebase operations)
Line numbers in a trace refer to positions at the recorded revision, not current positions. To query ownership of a specific line of code:
Use VCS blame to find the revision that last touched line N
Look up the trace for that revision and file
Find the range containing line N
For tracking attribution across code movement, use content hashes at the range level:
{
"files" : [
{
"path" : " src/parser.ts " ,
"conversations" : [
{
"url" : " https://api.example.com/v1/conversations/abc " ,
"contributor" : {
"type" : " ai " ,
"model_id" : " anthropic/claude-opus-4-5-20251101 "
},
"ranges" : [
{
"start_line" : 10 ,
"end_line" : 25 ,
"content_hash" : " murmur3:9f2e8a1b "
}
]
}
]
}
]
}
The hash applies to the specific range, allowing tracking even when code moves within or between files.
Model identifiers follow the models.dev convention:
{
"contributor" : {
"type" : " ai " ,
"model_id" : " anthropic/claude-opus-4-5-20251101 "
}
}
Format: provider/model-name
Each conversation has a url field and optional related array for linking to related sub-resources:
{
"files" : [
{
"path" : " src/api.ts " ,
"conversations" : [
{
"url" : " https://api.example.com/v1/conversations/abc123 " ,
"contributor" : {
"type" : " ai " ,
"model_id" : " anthropic/claude-opus-4-5-20251101 "
},
"ranges" : [{ "start_line" : 10 , "end_line" : 50 }],
"related" : [
{
"type" : " session " ,
"url" : " https://api.example.com/v1/sessions/xyz789 "
},
{
"type" : " prompt " ,
"url" : " https://api.example.com/v1/prompts/def456 "
}
]
}
]
}
]
}
7. Extensibility
Major version: breaking changes to required fields
Minor version: additive changes (new optional fields)
The metadata field can include implementation or vendor-specific data:
{
"metadata" : {
"confidence" : 0.95 ,
"post_processing_tools" : [ " prettier@3.0.0 " ],
"dev.cursor" : {
"workspace_id" : " ws-abc123 "
}
}
}
Vendors may use reverse-domain notation (e.g., dev.cursor , com.github.copilot ) within metadata to avoid key collisions.
A reference implementation is provided in the reference/ directory, demonstrating how to integrate Agent Trace with coding agents. The implementation includes:
trace-store.ts : A storage layer for reading and writing trace records
trace-hook.ts : Hook integration for automatic trace capture on file changes
The reference is an example for Cursor or Claude Code, but the patterns are applicable to any AI coding agent.
{
"version" : " 0.1.0 " ,
"id" : " 550e8400-e29b-41d4-a716-446655440000 " ,
"timestamp" : " 2026-01-25T10:00:00Z " ,
"files" : [
{
"path" : " src/app.ts " ,
"conversations" : [
{
"contributor" : { "type" : " ai " },
"ranges" : [{ "start_line" : 1 , "end_line" : 50 }]
}
]
}
]
}
B. MIME Type
Type
MIME Type
Trace Record
application/vnd.agent-trace.record+json
C. FAQ
How should I store the traces?
This spec intentionally does not define how traces are stored. This could be local files, git notes, a database, or anything else.
How should I handle rebases or merge commits?
We expect to see different implementations in open source. This may influence the spec in the future. We are open to feedback.
What happens when agents create scripts to write code?
This is left to the implementation. Code generated this way should still be attributed to the agent. For example, you could snapshot files before and after the script runs, then use git diff to determine what the agent added.
This specification is released under CC BY 4.0 .
Thanks to the following partners for helping shape Agent Trace:
This specification is accepting suggestions on GitHub .
A standard format for tracing AI-generated code.
There was an error while loading. Please reload this page .
Activity
Custom prope

[truncated]
