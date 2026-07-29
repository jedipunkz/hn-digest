---
source: "https://github.com/clay-good/OpenLore"
hn_url: "https://news.ycombinator.com/item?id=49104253"
title: "OpenLore: Deterministic, local-first memory and guardrails for AI coding agents"
article_title: "GitHub - clay-good/OpenLore: Deterministic, local-first memory and guardrails for AI coding agents with no LLM in the hot path. · GitHub"
author: "mafro"
captured_at: "2026-07-29T23:53:24Z"
capture_tool: "hn-digest"
hn_id: 49104253
score: 1
comments: 0
posted_at: "2026-07-29T23:06:53Z"
tags:
  - hacker-news
  - translated
---

# OpenLore: Deterministic, local-first memory and guardrails for AI coding agents

- HN: [49104253](https://news.ycombinator.com/item?id=49104253)
- Source: [github.com](https://github.com/clay-good/OpenLore)
- Score: 1
- Comments: 0
- Posted: 2026-07-29T23:06:53Z

## Translation

タイトル: OpenLore: AI コーディング エージェントのための決定論的、ローカルファーストのメモリとガードレール
記事のタイトル: GitHub - Clay-Good/OpenLore: ホット パスに LLM を持たない AI コーディング エージェントのための決定論的でローカル優先のメモリとガードレール。 · GitHub
説明: ホット パスに LLM を持たない AI コーディング エージェント用の決定論的なローカル ファースト メモリとガードレール。 - クレイグッド/OpenLore

記事本文:
GitHub - Clay-Good/OpenLore: ホット パスに LLM を持たない AI コーディング エージェント用の決定論的、ローカルファーストのメモリとガードレール。 · GitHub
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
粘土質の良い
/
オープンロア
公共
通知
ch にサインインする必要があります

アンジュ通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1,466 コミット 1,466 コミット .claude/ コマンド .claude/ コマンド .github .github .planning/ codebase .planning/ codebase docs docs 例 例 openspec openspec パッケージング/ homebrew パッケージング/ homebrew スキーマ スキーマ スクリプト スクリプト スキル スキル src src スタブ/tree-sitter-cli-stub スタブ/ Tree-sitter-cli-stub .gitignore .gitignore .npmrc .npmrc .nvmrc .nvmrc .prettierrc .prettierrc AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md eslint.config.js eslint.config.js flake.lock flake.lock flake.nix flake.nix package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts vitest.integration.config.ts vitest.integration.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング エージェント用の確定的でローカル ファーストのメモリとガードレール - ホット パスに LLM はありません。
1 回の呼び出しで、タスクが触れるコードをエージェントに伝えます。 1 つのゲートが、変更すると危険なものを示します。
静的解析に基づいています。 APIキーがありません。毎回同じ答え。
本物の未編集の録音 - ripgrep の新しいクローンで公開されたオープンロア。 install はエージェントに接続し、リポジトリにライブでインデックスを作成します — 14 秒間に 235 のファイル、2,978 の関数、4,329 のコール エッジ、API キーなし → orient はタスクが触れるコードを返します → レビューは 39 の呼び出し元を失効させた署名の変更をキャッチします → プロジェクトの見返りを証明します。自分で再記録します: docs/openlore-demo.tape 。
インストール · 得られるもの · ベンチマーク · ガバナンス · 仕組み · 代替手段との比較 · ドキュメント · コミュニティ
スタート

ここ: 1 つのコマンドでインストール · 何が得られるか · OpenLore はあなたに適していますか? · 実際の動作を確認 · 5 分間のクイックスタート · 導入にかかる費用
評価: 価値スコアカード (勝敗) · OpenLore と代替手段 · 既知の制限 · 私たちは独自のガバナンスをドッグフードします
理解する: 仕組み、ガバナンス、コア機能、言語と IaC、フェデレーション、相互運用性、PR レビュー
使用方法: エージェント チート シート · クロード コード スキル · 要件 · ドキュメント · 開発 · コミュニティ
AI コーディング エージェントは強力ですが、記憶喪失であり、管理されていません。すべてのタスクは、同じファイルを再読み取りして構造を再発見することから始まります。長いセッションは毎回、自信はあるものの陳腐な仮定へと静かに流れていきます。そして、変更によって契約が破られたり、アーキテクチャの境界を越えたり、機密コードへの道が開かれようとしているとき、エージェントには何も通知されません。
OpenLore は両方の半分を修正します。コードベースの静的分析を 1 回限り実行し、呼び出し構造、タイプ、テスト、意思決定、IaC、仕様ドリフトなどのナビゲート可能なナレッジ グラフを、編集するたびに徐々に最新の状態に保ちます。エージェントは MCP ツール (または CLI) を介してクエリを実行し、すでに指示されているすべてのタスクを開始し、変更が適用される前に変更を認証します。これは決定論的かつローカル ファーストであり、ホット パスに LLM はありません。そのため、同じ質問からは同じ根拠のある回答が返され、エージェントは事実が古くなったときに、自信のある推測を提供するのではなく通知されます。
npm install -g openlore && openlore install
その 1 つのコマンドでエージェント (Claude Code、Cursor、Cline、Continue、AGENTS.md) が自動検出され、orient() を自動的に呼び出すように接続され、MCP サーバーが登録され、インデックスが構築されます。API キーも設定も質問も必要ありません。次に、エージェントに次のように尋ねます。
orient("支払い方法を追加")
…そして、関連する関数、その呼び出し元、

仕様、テスト、およびそれぞれを変更するリスクを 1 回の呼び出しで照合します。完全なセットアップ、バリアント、および検証: 5 分間のクイックスタート。
設定は不要で、すべてが検出可能です。コア値にはキーは必要ありません。すべてのオプトイン機能 (埋め込み、サーフェスのカバー、コミット ゲート、スペック ストアなど) がアクティブかどうか、およびそれをオンにする 1 つのコマンドを確認するには、 openlore features を実行します。
OpenLore は、エージェントに対して決定的とローカルの 2 つのことを行います。OpenLore はアーキテクチャを記憶するので、すべてのタスクが方向性を持って開始され、変更が反映される前にエージェントが何を変更するかを管理します。
🧠 記憶 — すでに方向付けられたすべてのタスクを開始します
永続的なアーキテクチャ メモリ - orient() を 1 回実行します。エージェントは、セッション全体にわたる数十のファイル読み取りからシステムを再導出するのを停止します。アンカーされたメモと決定はリファクタリング後も存続します。名前変更または移動されたシンボルは、孤立するのではなく、次の分析時にそのメモリを (carryAcross 来歴を使用して) 転送します。
ワンコール指向 — orient(task) は、関連する関数、その呼び出し元、一致する仕様セクション、および挿入ポイントの候補を 1 回の呼び出しで返します。 15k ノード グラフで ~430µs p50。
1 つの統合されたグラフ — アプリケーション コード、Infrastructure-as-Code、およびアーキテクチャ上の決定はすべて同じノード/エッジ プリミティブに投影されるため、1 回の走査で 3 つすべてにまたがる質問に答えます。
テスト影響の選択とデッドコード — 「parseConfig() を変更しました — どのテストを実行すればよいですか?」逆方向コールグラフの到達可能性による。言語を超えたマークアンドスイープは、信頼タグが付けられ、決して削除されない権限を持っているものを見つけます。
コンテキストの鮮度の追跡 (エピステミック リース) — エージェントは、キャッシュされたファクトが古くなったとき (コンテキストが古くなった、リポジトリが移動された) ことを、自信のある推測に基づいて行動するのではなく通知されます。
🛡️ ガバナンス — エージェントが何を主張するかのガードレール

ンゲス
変更影響証明書 — 宣言した機密境界へのパスを差分が新たに開くときに、change_impact_certificate フラグを立てます (変更後は到達可能ですが、変更前は到達できません) — 差分、決定論的、LLM なし。
破壊的変更の判定 — certify_public_surface は、変更されたすべてのエクスポートの破壊的/非破壊的/破壊的可能性を差分上で分類し、各破壊がヒットするリポジトリ内のコンシューマーに名前を付けます。構造的には保守的であり、決して黙って「安全」であることはありません。
アーキテクチャの不変条件、事前編集 — check_architecture は、「A の下のファイルは B をインポートできますか?」と答えます。インポートが書き込まれる前に、宣言されたレイヤー/禁止されたルールに違反します (言語を超えて)。
推測ではなく根拠のある主張 — verify_claim は、エージェントが「X は死亡している」または「Y は変更しても安全である」と主張する前に、引用受領書とともに確定的な確認済み/反駁済み/検証不可能な評決を返します。
1 つのコミット ゲート — openlore enforce は、enforcement.policy を通じてすべてのガバナンスの検出結果を解決し、クラスがブロックしているもののみをブロックします (デフォルトでは勧告、API キーはありません)。決定事項は記録され、ゲート制御され、実際の仕様に同期されます。仕様/コードのドリフトがミリ秒単位で検出されました。
📊 構造的には正直 — 大規模なリポジトリの深いトレースでは、エージェントのラウンドトリップが -26% で、勝利の隣に損失が公開されます。すべてのパブリック クレームは、実行できるコマンドに追跡されます。純粋な静的分析: API キーもネットワークも必要なく、毎回同じ答えが得られます。
ツールを評価する最も簡単な方法は、それが自分に向いていないことをすぐに知ることです。それで:
この README の 1 行だけを読んだ場合の 1 つのアイデアは、エージェントの高価な失敗モードは無知ではなく、自信です。関数の存在を知らないモデルは探しに行きます。古くなった事実を「知っている」モデルは、自信を持ってその事実に基づいて構築され、レビュー時にその対価を支払うことになります。 OpenLore は、エージェントが次のことを実行できるように構築されています。

古い「その事実は古い」と「この変更により、あなたが機密であると言った道が開かれます」 - グラフから決定論的に、ループ内の 2 番目のモデルが最初のモデルについて推測することはありません。
同じ作業を 2 回。エージェントに、これまで見たことのないコマンドにフラグを追加するように依頼します。
深いマルチホップタスクに対するその形状変更の測定された効果: excalidraw で 25 → 16 往復、集計 -26%。これは魔法ではありません。タスクごとに構造を再発見することと、それをクエリすることの違いです。これが報われない部分を含む完全な数値は、「価値スコアカード」に記載されています。
実際の出力 — openlore orient --json "add a --since flag to the blast-radius command" 、このリポジトリで実行します (要約: フィールドは省略され、呼び出し元のリストは名前にフラット化されます):
{
"relevantFiles" : [ " src/cli/commands/blast-radius.ts " , " src/core/services/mcp-handlers/blast-radius.ts " ],
"関連する関数" : [
{ "name" : " computeBlastRadius " 、 "filePath" : " src/core/services/mcp-handlers/blast-radius.ts " ,
"signature" : " 非同期関数 computeBlastRadius(input: BlastRadiusInput): Promise<BlastRadiusBriefing> " ,
"fanIn" : 5 、 "isHub" : true 、 "言語" : " TypeScript " }
]、
"callPaths" : [
{ "関数" : " computeBlastRadius " ,
"callers" : [ " handleBlastRadius " 、 " computeImpactCertificate " 、 " runBlastRadiusCli " 、
" composeReview " 、 "collectGovernanceFindings " ] }
]、
「ランドマーク」: [
{ "名前" : " runBlastRadiusCli " 、 "ホップ" : 1 、
"シグナル" : [{ "ラベル" : "オーケストレーター" , "証拠" : { "ファンアウト" : 11 } },
{ "ラベル" : " volatile " 、 "証拠" : { "レベル" : " 中" 、 "コミット" : 6 、 "coChangedWith" : 5 } }] }
]、
"挿入ポイント" : [
{ "ランク" : 2 、 "名前" : " computeBlastRadius " 、 "役割" : " ハブ " 、 "戦略" : "cross_cutting_hook " 、
"reason" : " computeBlastRadius は 5 つの関数によって呼び出されます -- ロジックを追加します

e はコールサイト サーフェス全体に影響します。 " }
]、
"suggestedTools" : [ " Record_Decision " 、 " Analyze_impact " 、 " get_subgraph " 、 " check_spec_drift " ]
}
エージェントは、単一のファイルを読み取る前に、どこを見るべきか、何を触るべきか、そして触ると危険なものは何かを正確に知っています。すべてのフィールドはグラフから計算されます。モデルからは何も推測されません。
openlore Impact-certificate --base main # 私の diff は、宣言された機密境界への新しいパスを開きますか?
openlore certify-public-surface --base main # コンシューマーのパブリック API 契約を破ったのでしょうか?
openlore blast-radius タッチした呼び出し元/レイヤーの数、実行するテスト、ドリフトする仕様と決定
openlore enforce --hook # 1 つのゲート; 「ブロック」と分類した検出結果のみをブロックします
LLM も API キーもありません。実行するたびに同じ根拠のある答えが得られます。デフォルトでは勧告。検出結果ごとにブロックを選択します。
価値スコアカード – それ自体で元が取れますか?
OpenLore は、OpenLore を使用しているエージェントが、OpenLore を使用していない同じエージェントよりも少ない総コストで正解に到達した場合にのみその地位を獲得します。私たちはその不平等を測定し、勝敗を公表します。数値は、2026-06-01 に測定された Spec 14 エージェント ベンチマーク ( claude -p 、sonnet、N=4 中央値、ピン留めされた SHA、各アームを分離する --strict-mcp-config ) からのものです。
2026-06-03 ライブで再確認 (N=2): ディープタスクの勝利が再現 — okhttp −13% 。小さい/よく知られたケースは task-d です。

[切り捨てられた]

## Original Extract

Deterministic, local-first memory and guardrails for AI coding agents with no LLM in the hot path. - clay-good/OpenLore

GitHub - clay-good/OpenLore: Deterministic, local-first memory and guardrails for AI coding agents with no LLM in the hot path. · GitHub
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
clay-good
/
OpenLore
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1,466 Commits 1,466 Commits .claude/ commands .claude/ commands .github .github .planning/ codebase .planning/ codebase docs docs examples examples openspec openspec packaging/ homebrew packaging/ homebrew schemas schemas scripts scripts skills skills src src stubs/ tree-sitter-cli-stub stubs/ tree-sitter-cli-stub .gitignore .gitignore .npmrc .npmrc .nvmrc .nvmrc .prettierrc .prettierrc AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md eslint.config.js eslint.config.js flake.lock flake.lock flake.nix flake.nix package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts vitest.integration.config.ts vitest.integration.config.ts View all files Repository files navigation
Deterministic, local-first memory and guardrails for AI coding agents — with no LLM in the hot path.
One call tells your agent the code a task touches; one gate tells it what's unsafe to change.
Grounded in static analysis. No API key. Same answer every time.
A real, unedited recording — the published openlore on a fresh clone of ripgrep . install wires your agent and indexes the repo live — 235 files, 2,978 functions, 4,329 call edges in 14 seconds , no API key → orient returns the code a task touches → review catches a signature change that left 39 callers stale → prove projects the payoff. Re-record it yourself: docs/openlore-demo.tape .
Install · What you get · Benchmarks · Governance · How it works · vs. Alternatives · Docs · Community
Start here: Install in one command · What you get · Is OpenLore for you? · See it in action · 5-Minute Quickstart · What it costs to adopt
Evaluate it: Value Scorecard (wins and losses) · OpenLore vs. Alternatives · Known Limitations · We dogfood our own governance
Understand it: How It Works · Governance · Core Features · Languages & IaC · Federation, Interop & PR review
Use it: Agent Cheat Sheet · Claude Code Skill · Requirements · Documentation · Development · Community
AI coding agents are powerful but amnesiac and ungoverned . Every task starts by re-reading the same files to rediscover structure; every long session quietly drifts toward confident-but-stale assumptions; and nothing tells the agent when a change is about to break a contract, cross an architectural boundary, or open a path into sensitive code.
OpenLore fixes both halves. It runs a one-time static analysis of your codebase and keeps a navigable knowledge graph — call structure, types, tests, decisions, IaC, and spec drift — incrementally fresh as you edit. Agents query it through MCP tools (or the CLI) to start every task already oriented, and to certify a change before it lands. It is deterministic and local-first — no LLM in the hot path — so the same question returns the same grounded answer, and an agent is told when a fact has gone stale instead of served a confident guess.
npm install -g openlore && openlore install
That one command auto-detects your agent (Claude Code, Cursor, Cline, Continue, AGENTS.md), wires it to call orient() automatically , registers the MCP server , and builds the index — no API key, no config, no questions asked. Then ask your agent:
orient("add a payment method")
…and it begins the task already knowing the relevant functions, their callers, the matching specs, the tests, and the risk of changing each one — in a single call. Full setup, variants, and verification: 5-Minute Quickstart .
Zero config, everything discoverable. Core value needs no keys. To see every opt-in capability — embeddings, covering surfaces, the commit gate, the spec store, and more — whether each is active, and the one command to turn it on, run openlore features .
OpenLore does two things for an agent, both deterministic and local — it remembers your architecture so every task starts oriented, and it governs what the agent changes before the change lands.
🧠 Memory — start every task already oriented
Persistent architectural memory — orient() once; agents stop re-deriving the system from dozens of file reads, across sessions. Anchored notes and decisions survive refactors : a renamed or moved symbol carries its memory forward at the next analyze (with carriedAcross provenance) instead of orphaning it.
One-call orientation — orient(task) returns the relevant functions, their callers, matching spec sections, and insertion-point candidates in a single call. ~430µs p50 on a 15k-node graph.
One unified graph — application code, Infrastructure-as-Code , and architectural decisions all project onto the same node/edge primitives, so a single traversal answers questions that span all three.
Test-impact selection & dead-code — "I changed parseConfig() — which tests should I run?" by backward call-graph reachability; cross-language mark-and-sweep finds what's dead, confidence-tagged, never deletion authority.
Context-freshness tracking (Epistemic Lease) — an agent is told when a cached fact has gone stale (context aged, repo moved) instead of acting on a confident guess.
🛡️ Governance — guardrails on what the agent changes
Change-impact certificate — change_impact_certificate flags when a diff newly opens a path into a sensitive boundary you declared (reachable after the change but not before) — differential, deterministic, no LLM.
Breaking-change verdict — certify_public_surface classifies every changed export breaking / non-breaking / potentially-breaking over a diff and names the in-repo consumers each break hits; conservative by construction, never silently "safe".
Architecture invariants, pre-edit — check_architecture answers "may a file under A import B?" against your declared layer/forbidden rules before the import is written — cross-language.
Grounded claims, not guesses — verify_claim returns a deterministic confirmed / refuted / unverifiable verdict with a citation receipt before an agent asserts "X is dead" or "Y is safe to change".
One commit gate — openlore enforce resolves every governance finding through your enforcement.policy and blocks only on what you class blocking (advisory by default, no API key). Decisions are recorded, gated, and synced into living specs; spec/code drift detected in milliseconds .
📊 Honest by construction — −26% agent round-trips on deep traces in large repos, with the losses published next to the wins; every public claim traces to a command you can run. Pure static analysis: no API key, no network, same answer every time.
The fastest way to evaluate a tool is to find out quickly that it isn't for you. So:
One idea, if you only read one line of this README: an agent's expensive failure mode isn't ignorance — it's confidence . A model that doesn't know a function exists will go look. A model that "knows" a stale fact will confidently build on it, and you pay for that at review time. OpenLore is built so the agent can be told "that fact is stale" and "this change opens a path you said was sensitive" — deterministically, from the graph, with no second model in the loop guessing about the first.
The same task, twice. Ask an agent to add a flag to a command it has never seen:
The measured effect of that shape change on deep, multi-hop tasks: 25 → 16 round-trips on excalidraw, −26% aggregate. It is not magic — it is the difference between rediscovering structure per task and querying it. Full numbers, including where this doesn't pay off, in the Value Scorecard .
Real output — openlore orient --json "add a --since flag to the blast-radius command" , run on this repo (abridged: fields elided, caller lists flattened to names):
{
"relevantFiles" : [ " src/cli/commands/blast-radius.ts " , " src/core/services/mcp-handlers/blast-radius.ts " ],
"relevantFunctions" : [
{ "name" : " computeBlastRadius " , "filePath" : " src/core/services/mcp-handlers/blast-radius.ts " ,
"signature" : " async function computeBlastRadius(input: BlastRadiusInput): Promise<BlastRadiusBriefing> " ,
"fanIn" : 5 , "isHub" : true , "language" : " TypeScript " }
],
"callPaths" : [
{ "function" : " computeBlastRadius " ,
"callers" : [ " handleBlastRadius " , " computeImpactCertificate " , " runBlastRadiusCli " ,
" composeReview " , " collectGovernanceFindings " ] }
],
"landmarks" : [
{ "name" : " runBlastRadiusCli " , "hops" : 1 ,
"signals" : [{ "label" : " orchestrator " , "evidence" : { "fanOut" : 11 } },
{ "label" : " volatile " , "evidence" : { "level" : " medium " , "commits" : 6 , "coChangedWith" : 5 } }] }
],
"insertionPoints" : [
{ "rank" : 2 , "name" : " computeBlastRadius " , "role" : " hub " , "strategy" : " cross_cutting_hook " ,
"reason" : " computeBlastRadius is called by 5 functions -- adding logic here affects the entire callsite surface. " }
],
"suggestedTools" : [ " record_decision " , " analyze_impact " , " get_subgraph " , " check_spec_drift " ]
}
The agent knows exactly where to look, what it touches, and what's risky to touch — before reading a single file. Every field is computed from the graph; nothing is inferred by a model.
openlore impact-certificate --base main # does my diff open a new path into a declared sensitive boundary?
openlore certify-public-surface --base main # did I break a consumer's public API contract?
openlore blast-radius # callers/layers touched, tests to run, specs & decisions that drift
openlore enforce --hook # one gate; blocks only on findings you've classed `blocking`
No LLM, no API key — the same grounded answer every run. Advisory by default; you opt into blocking per finding.
Value Scorecard — does it pay for itself?
OpenLore only earns its place if an agent with it reaches a correct answer for less total cost than the same agent without it. We measure that inequality and publish it — wins and losses. Numbers are from the Spec 14 agent benchmark ( claude -p , sonnet, N=4 medians, pinned SHAs, --strict-mcp-config isolating each arm), measured 2026-06-01 .
Re-confirmed live 2026-06-03 (N=2): the deep-task win reproduces — okhttp −13% . The small/familiar case is task-d

[truncated]
