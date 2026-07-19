---
source: "https://org-edge.com/sysedge.html"
hn_url: "https://news.ycombinator.com/item?id=48968175"
title: "Show HN: Knowledge graph skill for Claude/Kimi Code"
article_title: "SysEdge — Ontological knowledge graph for Claude Code and Kimi Code CLI teams"
author: "org-edge"
captured_at: "2026-07-19T14:20:44Z"
capture_tool: "hn-digest"
hn_id: 48968175
score: 2
comments: 0
posted_at: "2026-07-19T13:45:01Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Knowledge graph skill for Claude/Kimi Code

- HN: [48968175](https://news.ycombinator.com/item?id=48968175)
- Source: [org-edge.com](https://org-edge.com/sysedge.html)
- Score: 2
- Comments: 0
- Posted: 2026-07-19T13:45:01Z

## Translation

タイトル: HN の表示: クロード/キミ コードのナレッジ グラフ スキル
記事のタイトル: SysEdge — クロード コードおよびキミ コード CLI チームのためのオントロジー ナレッジ グラフ
説明: 再現可能な実際の外部オープンソース コードベースで測定された方向トークンが 71% 減少しました。要件のトレーサビリティ、テスト カバレッジ、アーキテクチャ標準をすべてクエリ可能なグラフで表示します。
HN テキスト: 複数のクロード コード インスタンスを操作していると、要件のトレーサビリティ、テスト カバレッジ、インスタンス間の通信に関する問題が見つかりました。私はナレッジ グラフを使って作業していたので、クロード コード用のナレッジ グラフ スキルを構築しました。すべてのユーザー ストーリー、ユースケース、機能、テスト ケースがノードです。さらに、各インスタンスのバックログを把握できるように、欠陥と機能強化を追跡します。コードがどこに実装されているかを知ることでトークンを節約できますが、おそらくもっと重要なのは、要件の品質とテスト カバレッジにおける体系的なエラーを発見し、それを修正することができたことです。次に、アーキテクチャ パターンをサポートするように拡張しました。これにより、(既存のツールを使用して) コードとテスト ケースがグラフに自動的にインポートされます。ローカルの Neo4j グラフを使用します (Docker で実行します)。

記事本文:
メインコンテンツにスキップ
オーグエッジ
解決策
概要
組織内のエージェント
AI の変革
戦略カスケード
マネジメント能力
コンプライアンス文化
ガバナンス
観客
ボード用
リーダー向け
投資家向け
コンサルタント向け
管理者向け
価格設定
パートナー
シスエッジ
お問い合わせ
SysEdgeを入手する
SysEdge — クロード コードおよびキミ コード CLI チームのためのオントロジー ナレッジ グラフ
コードレビューが見逃しているものを捉えたグラフ。
要件が文書化されず、テストがリンクされず、アーキテクチャ標準が取り上げられなかったため、欠陥がすり抜けました。 SysEdge は、コードが出荷される前に、クエリ可能なグラフで 3 つすべてを表示できるようにします。再現できる実際の外部オープンソース コードベースで測定すると、方向トークンが 71% 減少します。 12 の並列 AI セッション、1 つの共有オントロジー、重複実装なし。
要件の品質 — テスト前、コードの前
不適切な要件により、不適切なコードが確実に生成されます。 SysEdge には、AI API キーを使用せずにグラフ クエリから自動的に、すべてのユーザー ストーリーとユース ケースを研究に裏付けられた一連の指定された標準に照らして評価する品質レビュー コマンドが同梱されています。構造チェック (13 の QUS 基準、6 つのトレーサビリティ ルール) は数秒で実行されます。 AI ガイダンス レイヤー (Cockburn UC パターン、セマンティック ストーリーの品質) は、必要なときに実行されます。
構造チェック (TRC-*、QUS-001 ～ 006) は AI キーなしで無料で実行されます。 --ai を追加して、AI ガイダンス レイヤー: Cockburn UC の完全性 (UC-G-001–006) とセマンティック ストーリーの品質 (QUS-G-001–007) を呼び出します。要件品質の詳細 →
記憶ツール、コードグラフ、または AI コードレビューを使用しないのはなぜでしょうか?
記憶ツールはチャットを記憶します。コードグラフ ツールはソースのインデックスを作成します。 AI コードレビュー担当者が差分を分析します。どれも役に立ちます。書かれていないものにフラグを立てることはできません。彼らには、決して仕様ではなかった要件を表面化する方法がありません

承認されたもの、登録されなかったテスト、または取り上げられたことのないアーキテクチャ標準などです。構築されなかったものには差分が存在しないためです。 SysEdge は、システムが何を行うべきか、そしてそれが実際に検証されているかどうかをモデル化するレイヤーです。
完全な比較: SysEdge とメモリ ツール、コード グラフ、AI コード レビュー →
再現可能な検証 — Formbricks (外部オープンソース TypeScript)
欠陥が 1 つあります。グラフではコードレビューで把握できなかったことが 4 つあります。
私たちは GitHub から Formbricks をクローンし、それに対して SysEdge をコールドで実行しました。コードベースについての予備知識はありませんでした。問題: DEF-FBK-001 — 匿名化が有効になっている場合、アンケート回答のエクスポートには PII が含まれます。トークンの節約 (71%) は本当ですが、それは話の中で最も小さな部分です。
エージェントの同等性。同じ Formbricks グラフと同じコマンドが、同じ Neo4j インスタンスに対して Kim Code CLI を使用して検証されました。カバレッジギャップ、アーキテクチャ標準の調査結果、およびオリエンテーショントークンの節約は、エージェント全体で一貫していました。元のベンチマークはクロード コードで実行されたため、公開された測定では Anthropic トークンが使用されています。 kimi ユーザーは、同じグラフによるソースの再読み込みの削減を期待できます。
しかし、このグラフは、そもそも欠陥が存在する理由を説明する 3 つのより深い問題も明らかにしました。
1 — 匿名化の要件は決して文書化されていない
7 つの AS-REQ ディメンションに対して Coverage-review --uc UC-FBK-005 (CSV への応答のエクスポート) を実行すると、3 つの FAIL が返されました。欠陥の最も直接的な原因:
また、フラグが付けられています: AS-REQ-007 FAIL (リンクされたテストが 0 件 – 仕様からテストまでのトレーサビリティなし) および認証境界とテスト容易性に関する警告結果が 4 件あります。
要件の詳細なトレーサビリティ →
2 — エクスポート機能には、すべての V モデル層でテストがありませんでした
test-gaps --Formbricks グラフの調査をインスタンス化します。
単体テストはありません。

API コントラクトのテストはありません。 Playwright UI フローはありません。エクスポート機能は、どの層でもテスト カバレッジがゼロの状態で出荷およびデプロイされていました。この機能はテスト グラフに登録されていなかったため、カバレッジの割合には表示されませんでした。
V モデルのテスト カバレッジの詳細 →
3 — 最も近い既存のテストでは、7 つの AS-TEST ディメンションのうち 7 つが不合格でした
エクスポート機能に最も近いテストである、response.spec.ts Playwright ファイルで Audit-test --uc UC-FBK-005 を指定すると、AS-TEST-UC ディメンション全体で 7 つの FAIL が生成されました。
追加の FAIL: エンドツーエンドのハッピー パス、等価性パーティション、セマンティックな正当性のアサーション、テストにおける仕様導出コメントはありません。
AI テスト品質監査の詳細 →
4 — PII エクスポートのアーキテクチャ標準にはアドレス指定がありませんでした
グラフには SEC-PII-001 が含まれていました。「アンケートで匿名化が有効になっている場合、すべての回答エクスポート エンドポイントはデータを返す前に PII フィールドを削除する必要があります。」この標準は存在していましたが、コードベースがそれにどのように対処するかを確認する ADR (アーキテクチャ決定記録) がありませんでした。単一のクエリで次のことが明らかになります。
3つの安全基準。記録された決定はゼロ。この欠陥は驚くべきことではありませんでした。それは、仕様のギャップ、テストのギャップ、アーキテクチャのギャップがすべて同時に存在し、グラフがなければ見えないことによる予測可能な結果でした。
アーキテクチャ標準の調整の詳細 →
2 番目の検証 — DOCUMENSO (電子署名プラットフォーム)
私たちは、オープンソースの DocuSign の代替品である Documenso に対しても同じ分析を実行しました。コールド クローン作成から 5 つの結果が得られるまで 15 分。エンベロープの有効期限を処理する Ingest ジョブ ハンドラーにはテストがありません。有効期限の仕様にはハッピー パスのみが記述されます。法的拘束力のある電子署名プラットフォームにとって、それは重要です。
SysEdge は、隣接する 2 種類のツールとよく比較されます: AI メモ

チャット セッションを記憶する ry/RAG ツールと、ナビゲーション用にソースのインデックスを作成するコードグラフ ツール。どちらも便利ですが、異なります。コードやチャットに含まれるものをマップします。 SysEdge は、システムの動作とそれが検証されているかどうかをマッピングし、信頼できる情報源として要件のトレーサビリティ、テスト カバレッジ、アーキテクチャ標準を強制します。
SysEdge とコードグラフおよびメモリ ツールの比較 →
無料のCLI。ワンタイムブートストラップキット。
CLI を使用すると、ターミナルから完全なグラフが表示されます。ブートストラップ キットには、Web ビジュアライザー、AI 監査コマンド、Docker セットアップ、アーキテクチャ標準ライブラリ、セッション スキル ファイルが追加されており、午後に生産性を高めるためのすべてが含まれています。
リポジトリごと、ワンタイム、即時ダウンロード、VAT 込み
1 日あたり 10 セッション · ソネット 4.6 · オリエンテーション トークンのみ
ブートストラップ キットは、1 日あたり 10 セッションで 5 日で回収されます。完全な測定方法 →
無料の CLI: MIT + コモンズ条項 — 商用ソフトウェアを含む独自のプロジェクトには無料です。
ブートストラップ キット: リポジトリごとにライセンス付与、12 か月の更新、そのリポジトリ上の無制限のチーム メンバー。
20 以上の並列セッションまたはカスタム言語のサポート — お問い合わせください。
完全なライセンス条件 → ·
プライバシーポリシー →
グラフはレビューで見逃したものを捉えています。
SysEdge は、12 インスタンスの AI コーディング システム (12 の並列セッション、1 つの共有コードベース) の実行から生まれました。 Formbricks と Documenso からの発見は珍しいものではありません。指定されていない例外パス、テストされていない機能、対処されていない標準、既存のテストでの仕様の派生がゼロです。これらは、ナレッジ グラフのないコードベースの通常の状態です。 1 つを使用すると、グラフ クエリになります。
GitHub からコールドでクローンされた 2 つの実稼働コードベースで検証されました。
Documenso (電子署名) — 15 分で 5 つの結果 → ·
Formbricks — 71% トークン削減 + 4 スペック

隙間→

## Original Extract

71% fewer orientation tokens, measured on a real external open-source codebase you can reproduce. Requirements traceability, test coverage, and architecture standards — all in a queryable graph.

Working with multiple Claude Code instances I found issues with requirements traceability, test coverage, and inter-instance communication. I was working with knowledge graphs so built a knowledge graph skill for Claude Code - every user story, use-case, feature, test-case is a node; plus i track defects and enhancements so each instance knows its backlog. It saves tokens by knowing where code is implemented, but perhaps more importantly it found systematic errors in requirements quality and test coverage that I was able to fix. I then extended to support architecture patterns and it will automatically import code and test cases into the graph (using your existing tool). It uses a local Neo4j graph (i run in docker).

Skip to main content
OrgEdge
Solution
Overview
Agents in your org
AI Transformation
Strategy Cascading
Management Capability
Compliance Culture
Governance
Audiences
For Boards
For Leaders
For Investors
For Consultants
For Managers
Pricing
Partners
SysEdge
Contact
Get SysEdge
SysEdge — ontological knowledge graph for Claude Code and Kimi Code CLI teams
The graph that catches what code review misses.
A defect slipped through because the requirement was never written down, the test was never linked, and the architecture standard was never addressed. SysEdge makes all three visible — in a queryable graph, before the code ships. Measured on a real external open-source codebase you can reproduce: 71% fewer orientation tokens . Twelve parallel AI sessions, one shared ontology, zero duplicate implementations.
Requirements quality — before tests, before code
Bad requirements produce bad code reliably. SysEdge ships a quality-review command that evaluates every user story and use case against a named set of research-backed standards — automatically, from graph queries, without an AI API key. The structural checks (13 QUS criteria, 6 traceability rules) run in seconds. The AI guidance layer (Cockburn UC patterns, semantic story quality) runs when you want it.
Structural checks (TRC-*, QUS-001–006) run free with no AI key. Add --ai to invoke the AI guidance layer: Cockburn UC completeness (UC-G-001–006) and semantic story quality (QUS-G-001–007). Requirements quality in depth →
Why not a memory tool, a code-graph, or AI code review?
Memory tools remember your chats. Code-graph tools index your source. AI code reviewers analyse the diff. All are useful — none can flag what was never written. They have no way to surface a requirement that was never specified, a test that was never registered, or an architecture standard that was never addressed — because there is no diff for the thing that was never built. SysEdge is the layer that models what your system is supposed to do and whether it's actually verified.
Full comparison: SysEdge vs memory tools, code graphs, and AI code review →
Reproducible verification — Formbricks (external open-source TypeScript)
One defect. Four things the graph caught that code review didn't.
We cloned Formbricks from GitHub and ran SysEdge against it cold — no prior knowledge of the codebase. The defect: DEF-FBK-001 — survey response export includes PII when anonymisation is enabled. The token saving (71%) is real, but it's the smallest part of the story.
Agent parity. The same Formbricks graph and the same commands were verified with Kimi Code CLI against the same Neo4j instance. Coverage gaps, architecture standards findings, and orientation-token savings were consistent across agents. The published measurement uses Anthropic tokens because the original benchmark ran on Claude Code; Kimi users can expect the same graph-driven reduction in source re-reading.
But the graph also surfaced three deeper problems that explain why the defect existed in the first place:
1 — The requirement for anonymisation was never written down
Running coverage-review --uc UC-FBK-005 (Export responses to CSV) against the 7 AS-REQ dimensions returned 3 FAILs. The most direct cause of the defect:
Also flagged: AS-REQ-007 FAIL (0 linked tests — no traceability from spec to test) and 4 WARN findings on auth boundaries and testability.
Requirements traceability in depth →
2 — The export feature had zero tests at every V-model tier
test-gaps --instance surveys on the Formbricks graph:
No unit test. No API contract test. No Playwright UI flow. The export feature had shipped and been deployed with zero test coverage at any tier — invisible to coverage percentages because the feature was never registered in the test graph.
V-model test coverage in depth →
3 — The closest existing test had 7 of 7 AS-TEST dimensions failing
Pointing audit-test --uc UC-FBK-005 at the response.spec.ts Playwright file — the nearest test to the export feature — produced 7 FAILs across the AS-TEST-UC dimensions:
Additional FAILs: no end-to-end happy path, no equivalence partitions, no semantic correctness assertions, no specification derivation comments in any test.
AI test quality audit in depth →
4 — The architecture standard for PII export had no addressing decision
The graph contained SEC-PII-001 : "When a survey has anonymisation enabled, all response export endpoints must strip PII fields before returning data." This standard existed — but had no ADR (Architecture Decision Record) confirming how the codebase addresses it. A single query surfaces this:
Three security standards. Zero decisions recorded. The defect was not a surprise — it was the predictable result of a specification gap, a test gap, and an architecture gap that all existed simultaneously and were all invisible without the graph.
Architecture standards alignment in depth →
SECOND VERIFICATION — DOCUMENSO (E-SIGNATURE PLATFORM)
We ran the same analysis on Documenso — an open-source DocuSign alternative. 15 minutes from cold clone to 5 findings. The Inngest job handlers that process envelope expiration have no tests. The expiration specification describes only the happy path. For a legally-binding e-signature platform, that's material.
SysEdge is often compared to two adjacent kinds of tool: AI memory/RAG tools that remember your chat sessions, and code-graph tools that index your source for navigation. Both are useful — and different. They map what your code or your chats contain . SysEdge maps what your system is supposed to do and whether it's verified , enforcing requirements traceability, test coverage, and architecture standards as the source of truth.
SysEdge vs code-graph and memory tools →
Free CLI. One-time bootstrap kit.
The CLI gives you the full graph from the terminal. The bootstrap kit adds the web visualiser, AI audit commands, Docker setup, architecture standards library, and session skill files — everything to be productive in an afternoon.
Per repository · one-time · instant download · VAT included
AT 10 SESSIONS/DAY · SONNET 4.6 · ORIENTATION TOKENS ONLY
Bootstrap kit pays back in 5 days at 10 sessions/day. Full measurement methodology →
Free CLI: MIT + Commons Clause — free for your own projects including commercial software.
Bootstrap kit: licensed per repository · 12 months updates · unlimited team members on that repo.
20+ parallel sessions or custom language support — contact us .
Full licence conditions → ·
Privacy policy →
The graph catches what the review missed.
SysEdge came out of running a 12-instance AI coding system — twelve parallel sessions, one shared codebase. The findings from Formbricks and Documenso aren't unusual: unspecified exception paths, untested features, unaddressed standards, zero specification derivation in existing tests. They are the normal state of a codebase without a knowledge graph. With one, they become graph queries.
Verified on two production codebases, cloned cold from GitHub:
Documenso (e-signature) — 5 findings in 15 minutes → ·
Formbricks — 71% token reduction + 4 spec gaps →
