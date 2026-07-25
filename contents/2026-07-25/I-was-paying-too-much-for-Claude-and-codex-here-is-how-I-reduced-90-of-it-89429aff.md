---
source: "https://github.com/AjnasNB/qarinah"
hn_url: "https://news.ycombinator.com/item?id=49050011"
title: "I was paying too much for Claude and codex here is how I reduced 90% of it"
article_title: "GitHub - AjnasNB/qarinah: Local, evidence-linked project memory and cited context compiler for coding agents - 98.71% less estimated repeated context in the published benchmark. · GitHub"
author: "Cognifyr"
captured_at: "2026-07-25T18:53:21Z"
capture_tool: "hn-digest"
hn_id: 49050011
score: 2
comments: 0
posted_at: "2026-07-25T18:05:26Z"
tags:
  - hacker-news
  - translated
---

# I was paying too much for Claude and codex here is how I reduced 90% of it

- HN: [49050011](https://news.ycombinator.com/item?id=49050011)
- Source: [github.com](https://github.com/AjnasNB/qarinah)
- Score: 2
- Comments: 0
- Posted: 2026-07-25T18:05:26Z

## Translation

タイトル: クロードとコーデックスにお金を払いすぎたので、その 90% を削減した方法を紹介します
記事のタイトル: GitHub - AjnasNB/qarinah: コーディング エージェント向けのローカルの証拠にリンクされたプロジェクト メモリと引用コンテキスト コンパイラ - 公開されたベンチマークで推定される反復コンテキストが 98.71% 減少します。 · GitHub
説明: ローカルの証拠にリンクされたプロジェクト メモリとコーディング エージェント用の引用コンテキスト コンパイラ - 公開されたベンチマークで推定される反復コンテキストが 98.71% 減少します。 - アジュナスNB/カリーナ

記事本文:
GitHub - AjnasNB/qarinah: ローカルの証拠にリンクされたプロジェクト メモリとコーディング エージェント用の引用コンテキスト コンパイラ - 公開されたベンチマークで推定される反復コンテキストが 98.71% 減少します。 · GitHub
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
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
アジナスNB
/
カリン

ああ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
44 コミット 44 コミット .agents/ plugins .agents/ plugins .claude-plugin .claude-plugin .github .github アセット アセット ベンチ ベンチ ビン ビン ドキュメント ドキュメント サンプル サンプル 統合 統合 出力/ PDF 出力/ PDF スキーマ スキーマ スクリプト スクリプト src src テストサポート テストサポートテスト テストタイプ タイプ Web サイト Web サイト .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md GOVERNANCE.md GOVERNANCE.md ライセンス ライセンス通知 通知 PRIVACY.md PRIVACY.md PRODUCT.md PRODUCT.md README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md TRADEMARKS.md TRADEMARKS.md package-lock.json package-lock.json package.json package.json server.json server.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ローカル プロジェクト メモリとコーディング エージェント用のコンテキスト コンパイラ。
ウェブサイト ·
ドキュメント ·
ホワイトペーパー・
土肥
ローカルファースト
証拠に基づく
グラフ対応
OKF-ポータブル
ガバナンス対応
推定コンテキストが 98.71% 減少 - 77.81:1 コンテキスト圧縮。
442,113 → 5,682 の推定入力トークン。
同じトークン レートで入力コンテキスト コストが 98.71% 削減されます。
自分で再現してください。
リプレイが少なくなります。現在のコード、ツール、引用されたプロジェクトのメモリのためのより多くのスペース。
技術論文・
出版物PDF・
ゼノドレコード ·
建築・
ベンチマーク ·
セキュリティ ·
打ち上げ計画
コーディング エージェントがプロジェクト コンテキストの繰り返しを 98.71% 削減して送信できるとしたらどうなるでしょうか?
推定 442,113 個の入力コンテキスト トークンが 5,682 個になり、コンテキストの繰り返しが 98.71% 減少し、必要なすべてのターゲットが直接カバーされ、コンテキスト圧縮が 77.81:1 になりました。

上位5位に入る。
繰り返されるコンテキストがほぼ 99% 減少します。選択されたすべてのメモリは、そのソースを指します。
React 編集、データベース移行、TypeScript リファクタリング、Web リサーチ、本番デバッグ、ガバナンスされたリリース作業にわたって正常に検証されました。評価されたタスクが送信した推定入力コンテキスト トークンは 436,431 個減少しました。キャッシュされていない入力トークン 100 万件あたり一律 1 ドルの場合、比較したコンテキスト スライスは 0.4421 ドルから 0.0057 ドルに移動し、同じ単価で入力コンテキスト コストが 98.71% 削減されます。このパーセンテージは、選択したフラット単価とは無関係です。ポータブル トークンの見積もりには、出力、ツール、キャッシュ、およびプロバイダーの固定料金は含まれていません。機械可読な結果と方法論を参照してください。
最初に引用したパックをインストールしてコンパイルする
npm install --save-dev カリーナ
npx qarinah 初期化 。
npx カリーナ レコード \
--親切な決定\
--title " リリースの出所を限定する " \
--body " レビュー済みのコミットからレビュー済みのアーティファクトのみを公開します。"
npx qarinah ビルド
npx qarinah クエリ「リリース来歴」\
--minimum-coverage 直接 \
--最大トークン 1500 \
--format マークダウン
5 分間のガイドから始めて、次に CLI リファレンス、JavaScript API リファレンス、MCP ガイド、タスク レシピ、またはトラブルシューティング ガイドを使用してください。
プロジェクトには、変更の背後にある決定事項と証拠がすでに含まれています。 Qarinah を使用すると、次のエージェントがそのレコードをクエリし、現在のタスク用に選択された制限付き引用パックを受信できるようになります。プロジェクト コンテキストを 1 つのエディターにロックするのではなく、同じローカル メモリで Codex、Claude Code、CLI ワークフロー、および互換性のある MCP クライアントをサポートできます。
Qarinah は、コーディング エージェント用のローカル メモリ コンパイラです。許可されたエージェント アクティビティ、プロジェクト構造、および明示的にコミットされた決定を、Codex、Claude Code、CLI、および管理されたワークフローのための耐久性のあるプロジェクト メモリに変換します。証拠を入力された形式で保存します

グラフ、決定論的なマークダウンおよび JSON ビューを作成し、不透明な概要や完全なトランスクリプトを真実のソースにする代わりに、現在のクエリに対して選択された制限付き引用パックをコンパイルします。
エージェントのメモリは通常、次の 2 つの方法のいずれかで失敗します。次のモデルが受信する履歴が多すぎるか、ソースを確認する方法がない圧縮されたストーリーを受信することです。 Qarinah は、ソース レコードとコンパクト コンテキストを分離して保持します。
メタデータのみのキャプチャがデフォルトです。コンテンツのキャプチャにはワークスペースの明示的な同意が必要です。隠された推論、プライベートトランスクリプト、認証情報、およびブラウザセッション状態は、製品の境界外に残ります。
モデルリクエストの前にメモリをコンパイルする
ホストまたはオーケストレーターがモデル リクエストを構築する前に Qarinah にクエリを実行すると、Qarinah は最初に保持されているプロジェクト履歴を制限された引用パックにコンパイルします。同じパックを、小規模なローカル モデル、大規模なコンテキスト モデル、高度な推論の Codex または Claude セッションに提供できます。コンパイラー自体には、埋め込み API、ホストされたメモリー サービス、または Qarinah API キーは必要ありません。
パックは明示的にリクエストされます。ホストは CLI または JavaScript API を呼び出すことができ、機密性の高い自動開示は個別に管理される Maqam 機能を通じて登録できます。 Qarinah の内蔵 MCP サーバーは、ゼロ書き込み診断面を維持します。
Qarinah は、サポートされているホスト アダプターによって配信される許可されたすべてのライフサイクル イベントと、ユーザーまたは管理されたワークフローが明示的にコミットしたすべての決定を記録します。すべての認知的決定を自動的に推論するとは主張しません。
サポートされるイベント クラスには、プロンプト、ツール リクエスト、ツールの完了、承認、アーティファクト、ソース、クレーム、決定、要約、圧縮、サブエージェント、完了したターン、失敗したターンが含まれます。関係はセッション、ターン、ツール呼び出し、ソース、承認、競合、スーパーセッションを接続します

n、証拠を導き出し、プロジェクト構造を作成しました。
プロジェクト グラフには、ディレクトリ、ファイル、コンテンツ ハッシュ、JavaScript および TypeScript モジュール参照、Markdown リンク、正確なソース スパン、追加、変更、名前変更、削除が含まれます。アーキテクチャ ガイドまたは編集可能な図のソースを参照してください。
Qarinah は意図的に小さく、ローカルで、検査可能です。
Qarinah には、メンテナンス済みの Node.js 22、24、または 26 リリースが必要です。
npm install --save-dev カリーナ
npx qarinah 初期化 。
パッケージはローカルでの使用を目的として設計されています。ホストされている Qarinah アカウント、埋め込みサービス、または Qarinah API キーは必要ありません。
一度初期化すれば、サポートされているセッション間で記憶されます
npx qarinah 初期化 。は、その正確なワークスペースとキャプチャ ポリシーに対する 1 回限りの明示的なオプトインです。レビュー済みの Codex または Claude Code 統合がインストールされ、そのホストが再起動された後、サポートされているライフサイクル フックは、ホストがイベントを発行するたびに、許可されたイベントを追加できます。その後、Qarinah は決定論的グラフを再構築し、オンデマンドで小さな引用パックをコンパイルできるため、次のタスクでは、保持されている履歴全体をプロンプトに再生する必要がありません。
Qarinah はプロジェクトのメモリであり、常に実行されているエージェントやアプリケーション スーパーバイザではありません。 Codex や Claude の実行を維持したり、プロバイダー側​​のコンテキスト圧縮を防止したり、ホストが公開していないホスト アクティビティをキャプチャしたり、MCP を通じてコン​​テキストを自動的に開示したりすることはありません。ホストが独自の会話を圧縮すると、Qarinah は実際に受信した許可された証拠のみを保存し、その証拠を明示的なクエリまたは個別に管理される開示機能で利用できるようにします。
# オプトインします。メタデータのみのキャプチャがデフォルトです。
npx qarinah 初期化 。
# 永続的な決定を 1 つコミットします。
npx カリーナ レコード \
--親切な決定\
--title " リリースの出所を限定する " \
--body " レビュー済みの成果物のみを公開します

d コミットします。 」
# 境界付きプロジェクト構造を記録し、派生ビューを再構築します。
npxカリーナスキャン
npx qarinah ビルド
# 直接的な証拠のみを取得し、引用されたマークダウンを出力します。
npx qarinah クエリ「リリース来歴」\
--minimum-coverage 直接 \
--format マークダウン
# ポリシー、イベント ハッシュ、チェックポイント、派生状態を確認します。
npxカリーナ博士
エージェント呼び出し元の場合は、信頼できないテキストがシェル コマンドに挿入されないように、厳密な JSON 標準入力インターフェイスを使用します。
printf ' %s ' ' {"query":"リリース来歴","format":"json","minimumCoverage":"direct","maxChars":8000} ' \
| npx qarinah クエリ --stdin-json
耐久性のあるファイル
.カリーナ/
config.json ポータブル ワークスペース ID と要求されたポリシー
events/events.jsonl 権限のある追加専用イベント チェーン
型付きエッジを持つgraph/graph.jsonイベントおよびプロジェクトノード
Index/index.json 使い捨ての決定的検索インデックス
records/CONTEXT.md 人間が判読できる現在のレコード
records/okf/ 再現可能な Markdown 相互運用性バンドル
インデックス/イベント ID/ チェックポイント認証された冪等性投影
派生グラフ、インデックス、またはマークダウン ビューを削除し、qarinah ビルドを実行して検証済みのイベント チェーンから再現します。
Qarinah は、検証済みのワークスペース レコードを決定論的な Google Open Knowledge Format 0.1 ドラフト バンドルとしてエクスポートできます。
npx qarinah エクスポート okf
エクスポートは、ルート インデックス、時系列ログ、イベントごとに 1 つのコンセプト ファイル、型付き関係、引用、コンテンツ ハッシュ、およびチェーン ハッシュを備えたレビュー可能なマークダウンです。 Git で差分を取得したり、Qarinah を使用せずに検査したり、OKF Markdown を理解する別のシステムに渡したりすることができます。追加専用の JSONL イベント チェーンは引き続き権威を持ちます。 OKF は、2 番目のデータベースや検索エンジンではなく、決定論的で置換可能な交換ビューです。 「相互運用性」を参照してください。
Qarinah の依存関係のないローカル レトリバーは BM を組み合わせます

25、文字トライグラムのタイプミス許容度、ワンホップ グラフの証拠、相互ランク融合、決定論的多様性、明示的な置き換え、競合の可視性、保持、時間、範囲指定された権限。
Context-pack v2 は証拠範囲を追加します。
{
「範囲」: {
"メソッド" : "クエリ用語-オーバーラップ-v1 " ,
"ステータス" : "直接" ,
"queryTermCount" : 2 、
"bestExactTermCount" : 2 、
"bestExactTermRatio" : 1 、
"directCandidateCount" : 3
}
}
minimumCoverage: 「部分的」は証拠のないパックを拒否します。 minimumCoverage: "direct" は、正規化されたすべてのクエリ用語を含むレコードのみを受け入れます。カバレッジは決定論的な検索診断であり、模範解答が正しいと主張するものではありません。
リポジトリには、Codex および Claude Code 用に生成された依存関係のないプラグイン ランタイムが含まれています。どちらも以下を提供します:
MCP ルートを公開しないホストの正確なワークスペース選択を備えたゼロ書き込み context_status および context_doctor MCP ツール。
ユーザー主導のローカル ワークフローに対する明示的な CLI クエリ。
Codex および Claude Code プラグイン キャッシュは不変のコピーです。レビューしたプラグインを再インストールし、アップグレード後に新しいタスクを開始します。クロードには、明示的に選択された絶対ノード 22、24、または 26 実行可能ファイルが必要です。 Codex はホストのレビューされた Node PATH 境界を引き続き継承します。

[切り捨てられた]

## Original Extract

Local, evidence-linked project memory and cited context compiler for coding agents - 98.71% less estimated repeated context in the published benchmark. - AjnasNB/qarinah

GitHub - AjnasNB/qarinah: Local, evidence-linked project memory and cited context compiler for coding agents - 98.71% less estimated repeated context in the published benchmark. · GitHub
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
AjnasNB
/
qarinah
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
44 Commits 44 Commits .agents/ plugins .agents/ plugins .claude-plugin .claude-plugin .github .github assets assets bench bench bin bin docs docs examples examples integrations integrations output/ pdf output/ pdf schemas schemas scripts scripts src src test-support test-support test test types types website website .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE NOTICE NOTICE PRIVACY.md PRIVACY.md PRODUCT.md PRODUCT.md README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md TRADEMARKS.md TRADEMARKS.md package-lock.json package-lock.json package.json package.json server.json server.json View all files Repository files navigation
Local project memory and a context compiler for coding agents.
Website ·
Documentation ·
White paper ·
DOI
LOCAL-FIRST
EVIDENCE-LINKED
GRAPH-AWARE
OKF-PORTABLE
GOVERNANCE-READY
98.71% less estimated context - 77.81:1 context compression.
442,113 → 5,682 estimated input tokens.
98.71% lower input-context cost at the same token rate.
Reproduce it yourself.
Less replay. More room for current code, tools, and cited project memory.
Technical paper ·
Publication PDF ·
Zenodo record ·
Architecture ·
Benchmarks ·
Security ·
Launch plan
What if your coding agents could send 98.71% less repeated project context?
442,113 estimated input-context tokens became 5,682 - 98.71% less repeated context and 77.81:1 context compression, with every required target directly covered in the top five.
Nearly 99% less repeated context. Every selected memory points back to its source.
Successfully verified across React editing, database migration, TypeScript refactoring, web research, production debugging, and governed release work. The evaluated tasks sent 436,431 fewer estimated input-context tokens. At a flat $1 per million uncached input tokens, that compared context slice moves from $0.4421 to $0.0057 - 98.71% less input-context cost under the same unit price. The percentage is independent of the chosen flat unit price; the portable token estimate excludes output, tools, caching, and fixed provider charges. See the machine-readable result and methodology .
Install and compile the first cited pack
npm install --save-dev qarinah
npx qarinah init .
npx qarinah record \
--kind decision \
--title " Keep releases provenance-bound " \
--body " Publish only the reviewed artifact from the reviewed commit. "
npx qarinah build
npx qarinah query " release provenance " \
--minimum-coverage direct \
--max-tokens 1500 \
--format markdown
Start with the five-minute guide , then use the CLI reference , JavaScript API reference , MCP guide , task recipes , or troubleshooting guide .
Your project already contains the decisions and evidence behind its changes. Qarinah lets the next agent query that record and receive a bounded, cited pack selected for the current task. The same local memory can support Codex, Claude Code, CLI workflows, and compatible MCP clients instead of locking project context to one editor.
Qarinah is a local memory compiler for coding agents. It turns permitted agent activity, project structure, and explicitly committed decisions into durable project memory for Codex, Claude Code, CLIs, and governed workflows. It preserves evidence in a typed graph and deterministic Markdown and JSON views, then compiles a bounded cited pack selected for the current query instead of making an opaque summary or a full transcript the source of truth.
Agent memory usually fails in one of two ways: the next model receives too much history, or it receives a compressed story with no way to verify the source. Qarinah keeps the source record and the compact context separate.
Metadata-only capture is the default. Content capture requires explicit workspace consent. Hidden reasoning, private transcripts, credentials, and browser session state remain outside the product boundary.
Compile memory before the model request
When a host or orchestrator queries Qarinah before constructing a model request, Qarinah compiles the retained project history into a bounded cited pack first. That same pack can be supplied to a small local model, a large-context model, or a high-reasoning Codex or Claude session. The compiler itself does not need an embedding API, a hosted memory service, or a Qarinah API key.
Packs are requested explicitly. Hosts can call the CLI or JavaScript API, while sensitive automated disclosure can be registered through a separately governed Maqam capability. Qarinah's built-in MCP server remains a zero-write diagnostic surface.
Qarinah records every permitted lifecycle event delivered by a supported host adapter and every decision that a user or governed workflow explicitly commits. It does not claim to infer every cognitive decision automatically.
Supported event classes include prompts, tool requests, tool completions, approvals, artifacts, sources, claims, decisions, summaries, compactions, subagents, completed turns, and failed turns. Relations connect sessions, turns, tool calls, sources, approvals, conflicts, supersession, derived evidence, and produced project structure.
The project graph covers directories, files, content hashes, JavaScript and TypeScript module references, Markdown links, exact source spans, additions, changes, renames, and deletions. See the architecture guide or the editable diagram source .
Qarinah is intentionally small, local, and inspectable:
Qarinah requires a maintained Node.js 22, 24, or 26 release.
npm install --save-dev qarinah
npx qarinah init .
The package is designed for local use. It does not require a hosted Qarinah account, embedding service, or Qarinah API key.
Initialize once, remember across supported sessions
npx qarinah init . is a one-time, explicit opt-in for that exact workspace and capture policy. After a reviewed Codex or Claude Code integration is installed and its host is restarted, supported lifecycle hooks can append permitted events whenever the host emits them. Qarinah can then rebuild the deterministic graph and compile a small cited pack on demand, so the next task does not need the whole retained history replayed into its prompt.
Qarinah is project memory, not an always-running agent or application supervisor. It does not keep Codex or Claude running, prevent provider-side context compaction, capture host activity the host does not expose, or automatically disclose context through MCP. When a host compacts its own conversation, Qarinah preserves only the permitted evidence it actually received and makes that evidence available to an explicit query or separately governed disclosure capability.
# Opt in. Metadata-only capture is the default.
npx qarinah init .
# Commit one durable decision.
npx qarinah record \
--kind decision \
--title " Keep releases provenance-bound " \
--body " Publish only the reviewed artifact from the reviewed commit. "
# Record the bounded project structure and rebuild derived views.
npx qarinah scan
npx qarinah build
# Retrieve only direct evidence and emit cited Markdown.
npx qarinah query " release provenance " \
--minimum-coverage direct \
--format markdown
# Verify policy, event hashes, checkpoint, and derived state.
npx qarinah doctor
For agent callers, use the strict JSON stdin interfaces so untrusted text is never interpolated into a shell command:
printf ' %s ' ' {"query":"release provenance","format":"json","minimumCoverage":"direct","maxChars":8000} ' \
| npx qarinah query --stdin-json
Durable files
.qarinah/
config.json portable workspace identity and requested policy
events/events.jsonl authoritative append-only event chain
graph/graph.json event and project nodes with typed edges
index/index.json disposable deterministic retrieval index
records/CONTEXT.md human-readable current record
records/okf/ reproducible Markdown interoperability bundle
index/event-ids/ checkpoint-authenticated idempotency projection
Delete any derived graph, index, or Markdown view and run qarinah build to reproduce it from the verified event chain.
Qarinah can export a verified workspace record as a deterministic Google Open Knowledge Format 0.1 Draft bundle:
npx qarinah export okf
The export is reviewable Markdown with a root index, a chronological log, one concept file per event, typed relations, citations, content hashes, and chain hashes. It can be diffed in Git, inspected without Qarinah, or passed to another system that understands OKF Markdown. The append-only JSONL event chain remains authoritative; OKF is a deterministic, replaceable interchange view rather than a second database or retrieval engine. See interoperability .
Qarinah's dependency-free local retriever combines BM25, character-trigram typo tolerance, one-hop graph evidence, reciprocal-rank fusion, deterministic diversity, explicit supersession, conflict visibility, retention, time, and scoped authority.
Context-pack v2 adds evidence coverage:
{
"coverage" : {
"method" : " query-term-overlap-v1 " ,
"status" : " direct " ,
"queryTermCount" : 2 ,
"bestExactTermCount" : 2 ,
"bestExactTermRatio" : 1 ,
"directCandidateCount" : 3
}
}
minimumCoverage: "partial" rejects no-evidence packs. minimumCoverage: "direct" accepts only a record containing every normalized query term. Coverage is a deterministic retrieval diagnostic, not a claim that a model answer is correct.
The repository includes generated, dependency-free plugin runtimes for Codex and Claude Code. Both provide:
zero-write context_status and context_doctor MCP tools with exact workspace selection for hosts that do not expose MCP roots;
explicit CLI querying for user-directed local workflows.
Codex and Claude Code plugin caches are immutable copies. Reinstall the reviewed plugin and start a new task after an upgrade. Claude requires an explicitly selected absolute Node 22, 24, or 26 executable. Codex still inherits the host's reviewed Node PATH boundary because its

[truncated]
