---
source: "https://github.com/BrightbeamAI/chap"
hn_url: "https://news.ycombinator.com/item?id=48658038"
title: "Ask HN: How should human overrides of AI agent outputs be recorded?"
article_title: "GitHub - BrightbeamAI/chap: Collaborative Human Agent Protocol (CHAP) · GitHub"
author: "arsalanshahid"
captured_at: "2026-06-24T12:01:36Z"
capture_tool: "hn-digest"
hn_id: 48658038
score: 1
comments: 0
posted_at: "2026-06-24T11:14:10Z"
tags:
  - hacker-news
  - translated
---

# Ask HN: How should human overrides of AI agent outputs be recorded?

- HN: [48658038](https://news.ycombinator.com/item?id=48658038)
- Source: [github.com](https://github.com/BrightbeamAI/chap)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T11:14:10Z

## Translation

タイトル: HN に聞く: AI エージェント出力の人によるオーバーライドはどのように記録されるべきですか?
記事タイトル: GitHub - BrightbeamAI/chap: Collaborative Human Agent Protocol (CHAP) · GitHub
説明: 協調ヒューマン エージェント プロトコル (CHAP)。 GitHub でアカウントを作成して、BrightbeamAI/chap の開発に貢献してください。

記事本文:
GitHub - BrightbeamAI/chap: 協調ヒューマン エージェント プロトコル (CHAP) · GitHub
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
ブライトビームAI
/
チャップ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイル C に移動

ode 「その他のアクション」メニューを開く フォルダーとファイル
22 コミット 22 コミット 適合性 適合性 コア コア デモ デモ 図 ドキュメント ドキュメントの例 統合 統合 パッケージ パッケージ プロファイル プロファイル リファレンス リファレンス スキーマ スキーマ ツール ツール ABOUT.md ABOUT.md ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md Dockerfile Dockerfile FAQ.md FAQ.md GLOSSARY.md GLOSSARY.md GOVERNANCE.md GOVERNANCE.md HANDBOOK.md HANDBOOK.md IMPLEMENTATIONS.md IMPLEMENTATIONS.md IN_PRACTICE.md IN_PRACTICE.md ライセンス ライセンス README.md README.md RELATIONSHIP-TO-OTHER-STANDARDS.md RELATIONSHIP-TO-OTHER-STANDARDS.md SECURITY.md SECURITY.md SPECIFICATION.md SPECIFICATION.md docker-compose.yml docker-compose.yml package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
協調ヒューマン エージェント プロトコル (CHAP)
人間とエージェントが実際の作業を一緒に行うためのプロトコル。
ボットが何かをドラフトし、それを人間が編集する場合、その編集はどこに保存されるのでしょうか?
CHAP では、6 か月後にクエリ、再生、検証できるエンベロープ内に保存されます。
インストール · 90 秒ツアー · 12 のシナリオ · このリポジトリについて · 論文
実際の仕事をするエージェントがいます。コードレビューの草案、チケットの優先順位付け、和解の提案、契約のレビュー。人間がそれぞれを承認、編集、または拒否します。現時点では、その決定はアプリケーション コード、チャット スレッド、チケット コメント、そしてあなたの頭の中にあります。 6 週間後に何か問題が発生した場合、何が起こったのかを再構築するには 45 分かかり、半分は当て推量になります。
CHAP は、これらの決定を反映するための 1 つの場所と 1 つの形式を提供します。エージェントのドラフトは人工物です。人間の編集は、差分、理論的根拠、タグを使用した構造化されたオーバーライドです。

あなたがコントロールします。全体はコンテンツ ハッシュによって連鎖します。 4 つの UI にわたってログを grep する代わりに、チェーンをクエリします。
Cursor を使用してプル リクエストをレビューする個人開発者。ボットは、開発者が同意しない「警告」にフラグを立てます。やり取り全体を端から端まで示します。
そして、これがコードの各行です。
1. ワークスペースをスピンアップします。 20 行、埋め込みコーディネーター、永続化用の SQLite:
import { Coordinator } から "@chap/coordinator" ;
import { SqliteStore } から "@chap/coordinator/storage/sqlite" ;
const coord = new Coordinator ( {store : new SqliteStore ( "./chap.db" ) } ) ;
コーディネートを待ちます。発送 ( {
jsonrpc : "2.0" 、 id : "1" 、
メソッド: "workspace.create" 、
パラメータ: {
ワークスペース: "wsp_pr_reviews" 、
プロファイル: [ "コア/1.0" 、 "レビュー/1.0" ]
}
} ) ;
コーディネートを待ちます。発送 ( {
jsonrpc : "2.0" 、 id : "2" 、
メソッド: "participant.join" 、
params : { ワークスペース : "wsp_pr_reviews" 、 from : "human:me@local" 、 type : "human" }
} ) ;
コーディネートを待ちます。発送 ( {
jsonrpc : "2.0" 、id : "3" 、
メソッド: "participant.join" 、
params : { ワークスペース : "wsp_pr_reviews" 、 from : "agent:cursor#v1" 、 type : "agent" }
} ) ;
2. ボットの下書きをオーバーライドします。既存の Cursor 統合をワイヤリングしてエンベロープを出力します。
// ボットのレビューはタスクの出力です。
const created = await coord 。発送 ( {
jsonrpc : "2.0" 、 id : "4" 、
メソッド: "task.create" 、
パラメータ: {
ワークスペース: "wsp_pr_reviews" 、
から: "エージェント:cursor#v1" 、
担当者: "エージェント:cursor#v1" ,
種類: "code_review" 、
input : { pr_id : "PR-482" 、 diff_url : "https://..." }
}
} ) ;
const taskId = 作成されました。結果 。タスクID ;
コーディネートを待ちます。発送 ( {
jsonrpc : "2.0" 、id : "5" 、
メソッド: "task.complete" 、
パラメータ: {
ワークスペース: "wsp_pr_reviews" 、
から: "エージェント:cursor#v1" 、
task_id : タスク ID 、
出力: カーソルレビュー
}
} ) ;
クーを待ってください

rd 。発送 ( {
jsonrpc : "2.0" 、id : "6" 、
メソッド: "review.request" 、
パラメータ: {
ワークスペース: "wsp_pr_reviews" 、
から: "エージェント:cursor#v1" 、
task_id : タスク ID 、
アーティファクト : カーソルレビュー
}
} ) ;
// あなたは 1 つのコメントに同意しません。オーバーライドします。
コーディネートを待ちます。発送 ( {
jsonrpc : "2.0" 、id : "7" 、
メソッド: "決定.オーバーライド" 、
パラメータ: {
ワークスペース: "wsp_pr_reviews" 、
から: "人間:私@ローカル" 、
task_id : タスク ID 、
意図保存: true 、
diff : [ { op : "replace" 、パス : "/comments/0/severity" 、値 : "info" } ] 、
理論的根拠: 「誤検知。フレームワークの規約であり、バグではありません。」 、
タグ : [ "偽陽性" 、 "フレームワークパターンの誤読" ]
}
} ) ;
3. 2 か月が経ち、自分が何をしてきたかを分析します。ここでプロトコルが利益をもたらします。リファレンス リポジトリには、監査チェーン (HTTP 経由または SQLite ファイルから直接) を読み取り、グループをオーバーライドする分析スクリプトが同梱されています。
# ステップ 1 の SqliteStore に対して、サーバーは必要ありません。
$ npm --prefix リファレンス/コアプラスレビュー実行分析 -- --db ./chap.db wsp_pr_reviews
学習レポートの上書き
========================
合計オーバーライド: 47
タグ別:
偽陽性 ████████████████ 31 (66%)
フレームワークパターンの誤読 ███████████ 22 (47%)
化粧品県████ 8 (17%)
上位のファイル パス:
src/handlers/ 18 オーバーライド
src/components/ 9 オーバーライド
Cursor の次のプロンプト リビジョンは推測ではなくなりました。パターンを名前で引用しています。
オーバーライド エンベロープの詳細
オーバーライド エンベロープは、CHAP で最も重要な形状です。どの分野にも仕事があります。
ほとんどの人が最初に読むときに見逃してしまう 2 つのフィールドは、tent_preserved と tags です。
mental_preserved は、洗練されたオーバーライド (人間がエージェントの決定に同意したが、その表現方法を書き直した) と置換を区別します。

無効なオーバーライド (人間が異なる決定に達した)。これらは 2 つの異なる障害モードであり、異なる修正が必要です。 1 つのポリシー条項付近の絞り込み率が高いということは、エージェントの取得がオフになっていることを意味します。同じ条項の置換率が高いということは、ポリシー自体があいまいであるか、エージェントのタスクのコンテキストが間違っていることを意味します。
タグは、チームが合意した管理された語彙です。小さくしておいてください。そこに何を入力しても、どのプロンプトに作業が必要かなどの質問に答えるときに、3 か月後に集計するディメンションになります。または、ボットが常に間違っているパスはどれですか?
npm install @chap/coordinator
パイソン:
pip インストール CHAP コーディネーター
どちらのパスでも、コアに加えて review/1.0 プロファイルと実行可能なリファレンスを取得できます。 TypeScript リファレンスは、reference/ にあります。 Python リファレンスは、reference/python/ にあります。 TypeScript ライブラリは、packages/coordinator/ にあります。 Python ライブラリは package/coordinator-py/ にあります。
5 分間の実践的なウォークスルー: example/00-five-minut-start.md 。
CHAP 0.2 は公開ドラフトです。具体的には、このリポジトリには次のものが含まれます。
仕様です。コア (7 つのメソッド、1 つのエンベロープ、1 つのワイヤ形式) と 10 のオプションのプロファイル。 SPECIFICATION.md で単一のドキュメントに結合するか、 core/SPEC.md および profiles/ から個別に読み取ります。
2 つのリファレンス実装。どちらもコアとすべてのプロファイル、合計 39 のメソッド ハンドラーをカバーします。 TypeScript リファレンスは package/coordinator/ にあり、HTTP サーバーはreference/core/ およびreference/core-plus-review/ にあり、2 つのブラウザ セッションとローカル LLM を備えた実行可能なプレイグラウンドは Reference/playground/ にあります。 Python リファレンスは package/coordinator-py/ にあり、HTTP サーバーは Reference/python/ にあります。どちらも同じ JSON-RPC 2.0 ワイヤ上で適合ハーネスを渡します。
適合ハーネス。 21 個のテスト ベクトル、署名

/canonicalisation/chain チェック、in-toto 構成証明出力。現在、2 つの適合レベル (最小、推奨) を主張できます。 2 つの実装にわたる広範な相互運用テストを完全に待機します。
MCP サーバーのトランスポート。 CHAP コーディネーターは、それ自体を MCP サーバーとして表現し、すべての CHAP メソッドをツールとして公開できます。クロード デスクトップ、カーソル、クロード コード、または任意の MCP クライアントをポイントして、自然言語から CHAP ワークスペースを操作します。 TypeScript アダプターは package/coordinator-mcp/ 、Python アダプターは chap_coordinator.transports.mcp_server 、実行可能なリファレンス サーバーは Reference/mcp-server-ts/ および Reference/mcp-server-py/ にあります。 Examples/drive-chap-from-claude-desktop.md で 5 分間のチュートリアルが行われます。
A2Aサーバートランスポート。 CHAP コーディネーターは、自身を A2A エージェントとして表現し、すべての CHAP メソッドをエージェント カード上の個別のスキルとしてアドバタイズすることもできます。 A2A 対応オーケストレーター (Azure AI Foundry、Amazon Bedrock AgentCore、Google ADK、カスタム マルチエージェント システム) は、URL によってコーディネーターを登録し、それに作業を委任できます。 TypeScript アダプターは package/coordinator-a2a/ 、Python アダプターは chap_coordinator.transports.a2a_server 、リファレンス サーバーは Reference/a2a-server-ts/ および Reference/a2a-server-py/ にあります。チュートリアルは、examples/drive-chap-from-an-a2a-orchestrator.md にあります。
内側ラップヘルパー。外部 MCP ツール呼び出しまたは A2A 交換を CHAP task.create + task.complete ペアに変換する小規模なライブラリ ユーティリティ。入出力正規化のハッシュは、結果として得られるアーティファクトの引用として記録されます。 integrations/CHAP-with-{MCP,A2A}.md の引用パターンに対応するライブラリ。 @chap/coordinator からは、wrapMcpToolCall/wrapA2aMessageExchange として、また、chap_coordinator.transports.wrap からは、wrap_mcp_tool_call/wrap_a2a_message_exchange として利用できます。
12 の実用的なシナリオ

。 IN_PRACTICE.md では、Cursor を使用した個人開発者から GMP 規制の充填仕上げ製造までの実際のケースを説明します。
重大な変更はセマンティック バージョニングに従います。プロファイル サーフェスはコアよりも速く移動します。厳密な安定性が必要な実稼働環境では、1.0 を待つ必要があります。より長いステータス ステートメントとコントリビューション パスは ABOUT.md にあります。
これを採用すると何が得られるか
キーのローテーション、ログの有効期限、およびユーザーの離脱後も存続する監査チェーン。すべてのエンベロープは、コンテンツ ハッシュによって前のエンベロープにリンクされます。 1 回の Audit.read 呼び出しですべてが返されます。
通常の作業の副作用としての構造化された監視データ。個別のアノテーション パイプラインはありません。すでに作成しているオーバーライドは、そうでなければコミッションする必要があるデータセットになります。
必要なときに、署名付きの否認できない承認を取得します。定義したsignature_meaningを持つOIDCバインド署名のsecurity-signed/1.0をオプトインします。サーバーを信頼せずに検証できる、外部透明性ログ アンカーの Audit-scitt/1.0 をオプトインします。
すでに構築したものと組み合わせ可能。 CHAP は MCP や A2A に代わるものではありません。それはそれらの隣にあります。エージェントはツールに MCP を使用し、他のエージェントに A2A を使用し、人間との共有作業を記録するために CHAP を使用します。
IN_PRACTICE.md 。個人開発から GMP 規制までの 12 の現実世界のシナリオ

[切り捨てられた]

## Original Extract

Collaborative Human Agent Protocol (CHAP). Contribute to BrightbeamAI/chap development by creating an account on GitHub.

GitHub - BrightbeamAI/chap: Collaborative Human Agent Protocol (CHAP) · GitHub
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
BrightbeamAI
/
chap
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
22 Commits 22 Commits conformance conformance core core demo demo diagrams diagrams docs docs examples examples integrations integrations packages packages profiles profiles reference reference schemas schemas tools tools ABOUT.md ABOUT.md ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile FAQ.md FAQ.md GLOSSARY.md GLOSSARY.md GOVERNANCE.md GOVERNANCE.md HANDBOOK.md HANDBOOK.md IMPLEMENTATIONS.md IMPLEMENTATIONS.md IN_PRACTICE.md IN_PRACTICE.md LICENSE LICENSE README.md README.md RELATIONSHIP-TO-OTHER-STANDARDS.md RELATIONSHIP-TO-OTHER-STANDARDS.md SECURITY.md SECURITY.md SPECIFICATION.md SPECIFICATION.md docker-compose.yml docker-compose.yml package.json package.json View all files Repository files navigation
Collaborative Human-Agent Protocol (CHAP)
The protocol for humans and agents doing real work together.
When a bot drafts something and a human edits it, where does that edit live?
In CHAP, it lives in an envelope you can query, replay, and verify six months later.
Install · The 90-second tour · Twelve scenarios · About this repo · Paper
You have agents doing real work. Drafting code reviews, triaging tickets, suggesting settlements, reviewing contracts. A human approves, edits, or rejects each one. Right now, that decision lives in your application code, your chat threads, your ticket comments, and your head. When something goes wrong six weeks later, reconstructing what happened costs you forty-five minutes and is half guesswork.
CHAP gives you one place to put those decisions and one shape to put them in. The agent's draft is an artefact. The human's edit is a structured override with a diff, a rationale, and tags you control. The whole thing chains together by content hash. You query the chain instead of grepping logs across four UIs.
A solo developer using Cursor to review pull requests. The bot flags a "warning" the developer disagrees with. Here is the whole exchange, end to end.
And here is the code, every line of it.
1. Spin up a workspace. Twenty lines, an embedded coordinator, SQLite for persistence:
import { Coordinator } from "@chap/coordinator" ;
import { SqliteStore } from "@chap/coordinator/storage/sqlite" ;
const coord = new Coordinator ( { store : new SqliteStore ( "./chap.db" ) } ) ;
await coord . dispatch ( {
jsonrpc : "2.0" , id : "1" ,
method : "workspace.create" ,
params : {
workspace : "wsp_pr_reviews" ,
profiles : [ "core/1.0" , "review/1.0" ]
}
} ) ;
await coord . dispatch ( {
jsonrpc : "2.0" , id : "2" ,
method : "participant.join" ,
params : { workspace : "wsp_pr_reviews" , from : "human:me@local" , type : "human" }
} ) ;
await coord . dispatch ( {
jsonrpc : "2.0" , id : "3" ,
method : "participant.join" ,
params : { workspace : "wsp_pr_reviews" , from : "agent:cursor#v1" , type : "agent" }
} ) ;
2. The bot drafts, you override. Wire your existing Cursor integration to emit envelopes:
// The bot's review is the output of a task.
const created = await coord . dispatch ( {
jsonrpc : "2.0" , id : "4" ,
method : "task.create" ,
params : {
workspace : "wsp_pr_reviews" ,
from : "agent:cursor#v1" ,
assignee : "agent:cursor#v1" ,
kind : "code_review" ,
input : { pr_id : "PR-482" , diff_url : "https://..." }
}
} ) ;
const taskId = created . result . task_id ;
await coord . dispatch ( {
jsonrpc : "2.0" , id : "5" ,
method : "task.complete" ,
params : {
workspace : "wsp_pr_reviews" ,
from : "agent:cursor#v1" ,
task_id : taskId ,
output : cursorReview
}
} ) ;
await coord . dispatch ( {
jsonrpc : "2.0" , id : "6" ,
method : "review.request" ,
params : {
workspace : "wsp_pr_reviews" ,
from : "agent:cursor#v1" ,
task_id : taskId ,
artefact : cursorReview
}
} ) ;
// You disagree with one comment. Override it.
await coord . dispatch ( {
jsonrpc : "2.0" , id : "7" ,
method : "decide.override" ,
params : {
workspace : "wsp_pr_reviews" ,
from : "human:me@local" ,
task_id : taskId ,
intent_preserved : true ,
diff : [ { op : "replace" , path : "/comments/0/severity" , value : "info" } ] ,
rationale : "False positive. Framework convention, not a bug." ,
tags : [ "false-positive" , "framework-pattern-misread" ]
}
} ) ;
3. Two months in, analyse what you have been doing. This is where the protocol pays you back. The reference repo ships an analytics script that reads the audit chain (either over HTTP or directly from your SQLite file) and groups overrides:
# Against the SqliteStore from step 1, no server required:
$ npm --prefix reference/core-plus-review run analyze -- --db ./chap.db wsp_pr_reviews
Override Learning Report
========================
Total overrides: 47
By tag:
false-positive ████████████████ 31 (66%)
framework-pattern-misread ███████████ 22 (47%)
cosmetic-pref ████ 8 (17%)
Top file paths:
src/handlers/ 18 overrides
src/components/ 9 overrides
Your next prompt revision for Cursor is no longer a guess. It cites the pattern by name.
The override envelope, in detail
The override envelope is the single most important shape in CHAP. Every field has a job:
The two fields most people miss on first read are intent_preserved and tags .
intent_preserved distinguishes a refining override (the human agreed with the agent's decision but rewrote how it was expressed) from a substituting override (the human reached a different decision). These are two different failure modes and they want different fixes. A high refining rate around one policy clause means the agent's retrieval is off; a high substituting rate on the same clause means the policy itself is ambiguous, or the agent's task context is wrong.
tags is the controlled vocabulary your team agrees on. Keep it small. Whatever you put there is the dimension you will aggregate on three months from now, when you are answering questions like which prompts need work? or which paths is the bot getting consistently wrong?
npm install @chap/coordinator
Python:
pip install chap-coordinator
Either path gets you Core plus the review/1.0 profile and a runnable reference. The TypeScript reference is in reference/ ; the Python reference is in reference/python/ . The TypeScript library lives at packages/coordinator/ ; the Python library at packages/coordinator-py/ .
Five-minute hands-on walkthrough: examples/00-five-minute-start.md .
CHAP 0.2 is a public draft. Concretely, this repo contains:
The specification. Core (seven methods, one envelope, one wire format) plus ten optional profiles. Combined into a single document at SPECIFICATION.md , or read individually from core/SPEC.md and profiles/ .
Two reference implementations. Both cover Core plus every profile, 39 method handlers in total . The TypeScript reference is at packages/coordinator/ , with HTTP servers at reference/core/ and reference/core-plus-review/ and a runnable playground with two browser sessions and a local LLM at reference/playground/ . The Python reference is at packages/coordinator-py/ with an HTTP server at reference/python/ . Both pass the conformance harness on the same JSON-RPC 2.0 wire.
A conformance harness. 21 test vectors, signing/canonicalisation/chain checks, in-toto attestation output. Two conformance levels are claimable today (Minimal, Recommended); Full waits on broader interop testing across the two implementations.
MCP server transport. A CHAP Coordinator can present itself as an MCP server, exposing every CHAP method as a tool. Point Claude Desktop, Cursor, Claude Code, or any MCP client at it and drive a CHAP workspace from natural language. TypeScript adapter at packages/coordinator-mcp/ , Python adapter at chap_coordinator.transports.mcp_server , runnable reference servers at reference/mcp-server-ts/ and reference/mcp-server-py/ . Five-minute walkthrough at examples/drive-chap-from-claude-desktop.md .
A2A server transport. A CHAP Coordinator can also present itself as an A2A agent, advertising every CHAP method as a discrete skill on its Agent Card. Any A2A-aware orchestrator (Azure AI Foundry, Amazon Bedrock AgentCore, Google ADK, custom multi-agent systems) can register the coordinator by URL and delegate work to it. TypeScript adapter at packages/coordinator-a2a/ , Python adapter at chap_coordinator.transports.a2a_server , reference servers at reference/a2a-server-ts/ and reference/a2a-server-py/ . Walkthrough at examples/drive-chap-from-an-a2a-orchestrator.md .
Inward wrap helpers. Small library utilities that turn an external MCP tool call or A2A exchange into a CHAP task.create + task.complete pair, with hashes of the input/output canonicalisations recorded as citations on the resulting artefact. The library counterpart to the citation patterns in integrations/CHAP-with-{MCP,A2A}.md . Available as wrapMcpToolCall / wrapA2aMessageExchange from @chap/coordinator , and as wrap_mcp_tool_call / wrap_a2a_message_exchange from chap_coordinator.transports.wrap .
Twelve worked scenarios. IN_PRACTICE.md walks through real cases from a solo developer with Cursor up to GMP-regulated fill-finish manufacturing.
Breaking changes follow Semantic Versioning. Profile surfaces will move faster than Core. Production deployments needing strict stability should wait for 1.0. The longer status statement and the contribution path are in ABOUT.md .
What you get when you adopt this
An audit chain that survives key rotation, log expiry, and people leaving. Every envelope links to the previous by content hash. One audit.read call returns the whole thing.
Structured supervision data as a side effect of normal work. No separate annotation pipeline. The overrides you are already making become a dataset you would otherwise have to commission.
Signed, non-repudiable approvals when you need them. Opt into security-signed/1.0 for OIDC-bound signatures with a signature_meaning you define. Opt into audit-scitt/1.0 for an external transparency-log anchor, verifiable without trusting your servers.
Composability with what you have already built. CHAP does not replace MCP or A2A. It sits next to them: your agent uses MCP for tools, A2A for other agents, and CHAP to record the shared work with humans.
IN_PRACTICE.md . Twelve real-world scenarios from solo dev to GMP-regulate

[truncated]
