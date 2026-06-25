---
source: "https://github.com/Dhevenddra/forensic-deepdive"
hn_url: "https://news.ycombinator.com/item?id=48672787"
title: "Show HN: Forensic-deepdive: code knowledge graph and MCP server for AI agents"
article_title: "GitHub - Dhevenddra/forensic-deepdive · GitHub"
author: "dhevenddra_"
captured_at: "2026-06-25T13:43:09Z"
capture_tool: "hn-digest"
hn_id: 48672787
score: 2
comments: 0
posted_at: "2026-06-25T13:08:54Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Forensic-deepdive: code knowledge graph and MCP server for AI agents

- HN: [48672787](https://news.ycombinator.com/item?id=48672787)
- Source: [github.com](https://github.com/Dhevenddra/forensic-deepdive)
- Score: 2
- Comments: 0
- Posted: 2026-06-25T13:08:54Z

## Translation

タイトル: HN の表示: フォレンジックの詳細: AI エージェント用のコード ナレッジ グラフと MCP サーバー
記事のタイトル: GitHub - Dhevenddra/forensic-deepdive · GitHub
説明: GitHub でアカウントを作成して、Dhevenddra/forensic-deepdive の開発に貢献します。

記事本文:
GitHub - デベンドラ/フォレンジックディープダイブ · GitHub
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
デベンドラ
/
法医学の詳細
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイル コードに移動 さらにアクションを開く

メニュー フォルダーとファイル
143 コミット 143 コミット .claude-plugin .claude-plugin .claude/ skill .claude/ skill .github/ workflows .github/ workflows docs docs 例 例 実験/ fastcontext 実験/ fastcontext スクリプト スクリプト src/ forensic_deepdive src/ forensic_deepdive テスト テスト .gitignore .gitignore .mcp.json .mcp.json .python-version .python-version AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス通知 通知 README.md README.md pyproject.toml pyproject.toml サーバー.json サーバー.json uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング エージェント用の永続的なコード ナレッジ グラフ + MCP サーバー。人間が判読できる投影としての 5 つの耐久性のあるマークダウン アーティファクト。 Apache-2.0。
forensic-deepdive はあらゆるコードベース (9 言語、多言語) を分析し、以下を生成します。
<repo>/.deepdive/graph.lbug の永続的な埋め込みグラフ — ファイル、シンボル、モジュール、コミット、作成者、エンドポイント、および DbTable ノードに加えて、DEFINES、MEMBER_OF、IMPORTS、CALLS、EXTENDS、IMPLEMENTS、TOUCHED_BY_COMMIT、AUTHORED_BY、CO_CHANGES_WITH、および境界を越えた HANDLES / CALLS_ENDPOINT / ROUTES_TO / INJECTS / PERSISTS_TO エッジ。すべてのエッジには信頼性タグ ( EXTRACTED / INFERRED / AMBIGUOUS ) が付けられており、隠れたヒューリスティックはありません。単一のエンドポイント参加ノードは 5 つの境界を越えたプロトコル (HTTP、MCP ツール、レジストリ ディスパッチ、gRPC、メッセージング/AMQP) を統合するため、フロントエンド呼び出しはスタック全体で 1 つの ROUTES_TO エッジとしてバックエンド ハンドラーに解決されます。
Claude Code、Cursor、Codex、Continue、Cline、Windsurf、およびその他の MCP 対応エージェントによって使用可能な 9 つの複合ツール (impact、context、archeology、flow、query、record_insight、recall_insights、visual、trace) を公開する MCP サーバー (forensicserve)。
耐久性のある 5 つの値下げ

<repo>/docs/codebase/ の下にあるアーティファクト。抽出ごとにグラフから再生成されます。
MAP.md — どこに何があるか、中心性によってランク付けされます。
HOTPATHS.md — 行ごとの信頼度混合列を備えた依存関係のホット スポット。各シンボルがどの程度正確に解決されるかを正確に確認できます。
ARCHAEOLOGY.md — コードがそのように見える理由 (git 履歴、% を含む上位作成者、バス係数、共同変更クラスター、欠陥の近接性)。
MENTAL_MODEL.md — 元の作成者が新入社員を入社させるために作成するドキュメント。
AGENT_BRIEF.md — ルールごとの信頼性タグを含む 5 KB 以下のアサーティブ Never/Always ルール。任意のプロジェクトに CLAUDE.md をドロップインします。
ターゲット リポジトリへの 10 個の shim — 4 つのエディター ルール ファイル ( CLAUDE.md 、 AGENTS.md 、 .cursor/rules/codebase.mdc 、 . continue/rules/codebase.md )、 .claude/skills/codebase-{exploring,debugging,impact-analysis,refactoring,onboarding}/ の下の 5 つの単一目的のクロード スキル、および.claude-plugin/plugin.json マニフェスト。すべての不在時書き込み - 手動で編集したファイルは上書きされません。
エージェント インサイト レイヤー - デフォルトで <repo>/.deepdive/insights.jsonl によってサポートされる Record_insight / remember_insights MCP ツール (依存関係なし、人間が判読できる、Git フレンドリー)。オプションの [graphiti] 追加オプションは、リポジトリ サイズの 2/5 しきい値を超えると、一時ナレッジ グラフ バックエンドにアップグレードします。
Extract はまた、 ARCHITECTURE.md を再生成します。これは、境界を越えたグラフ (ROUTES_TO / INJECTS / PERSISTS_TO、信頼性スタイル) のシステムレベルの Mermaid ビュー、別個の人間による検証サーフェス (5 つの契約アーティファクトの 1 つではなく、フォレンジックの視覚化とサーブ --ui とまったく同様) です。フォレンジック ダイアグラム --repo <repo> を使用して独自に再生成します。これを使用してグラフの健全性をチェックします。間違ったエッジはどこにでもあります。
--emit-vault を追加して、Obsidian に適した Vault を <output>/vault/ に書き込む - すべてのアーティファク

t は summary: / tags:frontmatter を取得し、相互参照は [[wikilinks]] になり、INDEX.md MOC がそれらを (.obsidian/ config で) 結び付けます。人間 (グラフ ビュー、バックリンク) とエージェント (概要によるトリアージ: ファイルを開かずに、トラバース可能なインデックス) のためのローカル ファーストの第 2 の脳。オプトイン;デフォルトではオフになっています。
v0.8.0 "USABLE → USEFUL + public release" — 最初のパブリック PyPI リリース。凍結された 5 つのプロトコル境界間グラフ (1 つのエンドポイント結合ノード上の HTTP/MCP/レジストリ/gRPC/メッセージング) 上に構築され、精度パス (正直なコールグラフの信頼性、個別呼び出し元数、低履歴/ソロリポジトリガード)、人による検証 ARCHITECTURE.md ダイアグラムサーフェイス、ディストリビューション (PyPI + MCP レジストリ + クロード コード プラグイン)、およびオプトインを備えています。 --emit-vault 黒曜石のエクスポート。 5 つのアーティファクト + 9 つの MCP ツールの契約は凍結されています。
何が証明され、何が証明されていないのか（正直なフレーミング）。 v0.8 は分析支援ツールです。実際の新人エージェントのオンボーディング テストでは、これが使用可能であること、エージェントが AGENT_BRIEF.md を自動検出し、プロンプトなしで適切なスキルにルーティングされることが確認されました。また、根拠のある MCP ツールのレビューでは、git-archaeology + 厳選されたブリーフが信頼性の高いコアであることが判明しました。自律的なエンドツーエンドの質問 — ディープダイブ シーディングにより、エージェントは現実の問題を測定可能なほど速く解決できるのか — はまだ証明されていません。モデルフリーのローカリゼーション パイロットが記録され (experiments/fastcontext/RESULTS.md — 静的シードは弱い事前分布です)、エンドツーエンドの測定は v0.9 まで延期されます (GPU + フロンティア メイン エージェント エンドポイントが必要です)。自律実行の主張は行われません。 Apache Superset、セキレイ (Django)、spring-petclinic、ripgrep、fastapi、Iris-Nearby (Flutter/Dart) を含む実際のリポジトリ全体で受け入れられます — docs/findings/ を参照してください。
# PyPI からインストールします (PATH に「forensic」を置きます);または uvx で一時的に実行します
UVツールをインストールする前に

nsic-ディープダイブ
フォレンジック情報 # バナー + 機能パネル
フォレンジック抽出 /path/to/repo
# …または開発用のソースから:
git clone https://github.com/Dhevenddra/forensic-deepdive && cd forensic-deepdive
UV 同期 --all-extras
＃何ができるの？ (バナー + 機能パネル: アーティファクト、プロトコル、MCP ツール、信頼性凡例)
UV 実行フォレンジック情報
# 任意のリポジトリで実行
uv run forensic extract /path/to/repo
# グラフは <repo>/.deepdive/graph.lbug に到達します
# <repo>/docs/codebase/ にある 5 つのマークダウン アーティファクト
# <repo>/.claude/、.cursor/、. continue/、root に 10 個の shim
# クロススタック機能スライスをトレースします (フロントエンド呼び出し -> エンドポイント -> ハンドラー -> テール)
uv フォレンジック トレース < シンボル > --repo /path/to/repo を実行します
# MCP サーバーとしてグラフをクエリします (分析されたリポジトリをポイントします)
uv run forensicserve --repo /path/to/repo
# 分析したすべてのリポジトリを検査します
UV実行フォレンジックリスト
PyPIからインストール
フォレンジックディープダイブとして出版 —
クローンは必要ありません:
uv ツールのインストール forensic-deepdive # `forensic` を PATH に置きます
フォレンジック抽出 /path/to/repo
# …またはインストールせずに一時的に実行します。
uvx forensic-deepdive 抽出 /path/to/repo
オプションの追加機能: UV ツールのインストール「forensic-deepdive[semantic]」(オフライン ONNX NL)
クエリ)、[openapi] (YAML 仕様解析)、[graphiti] (時間的洞察バックエンド)。
pip install forensic-deepdive は、 uv を使用していない場合にも機能します。
フォレンジック サーブは、9 つの複合ツールを任意のサーバーに公開する標準出力 MCP サーバーです。
MCP 対応エージェント (Claude Code、Cursor、VS Code/Copilot、Codex、Continue、Cline、
ウィンドサーフィン）。まずグラフを一度構築し (フォレンジック抽出 <repo> )、次に
サーバー。 3 つの方法、最初に簡単:
1. Claude Code プラグイン (自己ホスト型マーケットプレイス - PyPI ステップなし):
/plugin Marketplace add Dhevenddra/forensic-deepdive
/plugin install forensic-deepdive@dhevenddra
2. MCP レジストリから — 次のようにインデックス付けされます
io.git

Hub.Dhevenddra/forensic-deepdive 、つまりレジストリ対応クライアントと検出ハブ
(PulseMCP、MCPJungle、VS Code @mcp インデックス) を見つけて直接インストールできます。
3. 手動構成 — フォレンジック mcp-config を使用してクライアント スニペットを生成するか、以下を貼り付けます。
{
"mcpサーバー": {
"フォレンジックディープダイブ" : {
"コマンド" : " uvx " ,
"args" : [ "forensic-deepdive" 、 "serve " 、 " --repo " 、 " . " ]
}
}
}
クライアントごとのコピー＆ペーストブロック（カーソル、VS Code、Codex、uvx -not-found GUI の注意点）
docs/install.md にあります。
Python、C、Dart、Swift、TypeScript、JavaScript、Java、Go、Rust。
ツール
何をするのか
影響(シンボル、深さ、方向、min_confidence)
CALLS エッジ上のブラスト半径 BFS、深さバケット化、信頼性フィルター可能。
コンテキスト(シンボル)
シングルコールキッチンシンク: 定義 + 呼び出し元 + 呼び出し先 + 親/兄弟/メンバー + 拡張/実装 + 最近のコミット + 支配的な作成者 + 最近の洞察。
考古学(ファイルまたはシンボル)
チャーン、% の上位作成者、バス係数、同時変更クラスター、欠陥の近接性、最近のコミット。
フロー(エントリポイント、最大深さ)
サイクル検出を備えた DFS over CALLS。
クエリ(暗号 | 自然言語)
ヒットごとの来歴と信頼性を備えた生の Cypher、またはハイブリッド NL 検索 (FTS5/BM25 + 構造グラフ信号 + オプトイン オフライン セマンティック、RRF 融合および整形)。
Record_insight(シンボル、主張、証拠、検証済み)
検証された学習を継続します。
remember_insights(シンボル、以降、制限)
保存されているインサイトに対する部分文字列の最新のものから順に一致します。
Visualize(ターゲット、フォーマット、深さ、最大ノード数、...)
シンボル/ファイルの近傍 (または中央) の有界人魚図。エッジダッシュスタイルは自信を表します。
トレース(シンボル、方向、最大深さ)
エンドポイント結合ノード全体にわたるクロススタック機能スライス: ダウンストリームはフロントエンド呼び出し → CALLS_ENDPOINT → エンドポイント → HANDLES → ハンドラー → CALLS テールを歩きます。上流の答えは「誰が電話するか」

これがエンドポイントです。」
ツールの説明は個別に 200 トークン以下であるため、9 つのツールのエンベロープは Anthropic のターンごとのスキル メタデータ バジェット内に快適に収まります。
すべてのエッジとすべての発行されたクレームには、 EXTRACTED / INFERRED / AMBIGUOUS が含まれます。
EXTRACTED — AST または git log から決定的。事実。
INFERRED — きれいに解決されたヒューリスティック (インポート グラフ ウォーク、レシーバー タイプ推論、単一の同名候補クロス ファイル)。信頼性は高いが派生。
曖昧 — 複数の候補が浮上。リゾルバーは曖昧さを解消できませんでした。黙って推測するのではなく、すべての候補者を確認します。
HOTPATHS には行ごとの信頼度混合列が表示されるため、Logger (4 EXTRACTED + 1458 INFERRED — ほとんどクリーン) と ChatToolResponse (449 AMBIGUOUS — 同じ名前間のファイル間の衝突) を一目で区別できます。
オネストモード (純粋な静的、LLM なし、ネットワークなし)
フォレンジック抽出は、 ANTHROPIC_API_KEY 、 OPENAI_API_KEY 、 Ollama 、ネットワークなしでエンドツーエンドで機能します。 Graphiti は、[graphiti] PyPI extra と 2-of-5 のリポジトリ サイズしきい値 (50,000 LOC 以上、貢献者 25 人以上、生後 18 か月以上、PR 200 件/12 か月以上、ディスカッションのある問題数 100 件以上) を介してオプトインされます。 JsonlInsightStore は、いつでも利用できるフロアです。
なぜこれで、そうでないのか [GitNexus / CodeGraphContext / DeepWiki / Sourcegraph]
法医学の詳細
GitNexus
コードグラフコンテキスト
ディープウィキ
ソースグラフ

[切り捨てられた]

## Original Extract

Contribute to Dhevenddra/forensic-deepdive development by creating an account on GitHub.

GitHub - Dhevenddra/forensic-deepdive · GitHub
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
Dhevenddra
/
forensic-deepdive
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
143 Commits 143 Commits .claude-plugin .claude-plugin .claude/ skills .claude/ skills .github/ workflows .github/ workflows docs docs examples examples experiments/ fastcontext experiments/ fastcontext scripts scripts src/ forensic_deepdive src/ forensic_deepdive tests tests .gitignore .gitignore .mcp.json .mcp.json .python-version .python-version AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md pyproject.toml pyproject.toml server.json server.json uv.lock uv.lock View all files Repository files navigation
A persistent code knowledge graph + MCP server for AI coding agents. Five durable markdown artifacts as the human-readable projection. Apache-2.0.
forensic-deepdive analyzes any codebase (9 languages, polyglot) and produces:
A persistent embedded graph at <repo>/.deepdive/graph.lbug — File, Symbol, Module, Commit, Author, Endpoint , and DbTable nodes plus DEFINES, MEMBER_OF, IMPORTS, CALLS, EXTENDS, IMPLEMENTS, TOUCHED_BY_COMMIT, AUTHORED_BY, CO_CHANGES_WITH, and the cross-boundary HANDLES / CALLS_ENDPOINT / ROUTES_TO / INJECTS / PERSISTS_TO edges. Every edge carries a confidence tag ( EXTRACTED / INFERRED / AMBIGUOUS ) — no hidden heuristics. The single Endpoint join node unifies five cross-boundary protocols (HTTP, MCP tools, registry-dispatch, gRPC, messaging/AMQP), so a frontend call resolves to its backend handler across the stack as one ROUTES_TO edge.
An MCP server ( forensic serve ) exposing 9 composite tools ( impact , context , archaeology , flow , query , record_insight , recall_insights , visualize , trace ) consumable by Claude Code, Cursor, Codex, Continue, Cline, Windsurf — and any other MCP-aware agent.
Five durable markdown artifacts under <repo>/docs/codebase/ , regenerated from the graph on every extract:
MAP.md — what's where, ranked by centrality.
HOTPATHS.md — the dependency hot spots, with a per-row confidence-mix column so you see exactly how cleanly each symbol resolves.
ARCHAEOLOGY.md — why the code looks the way it does (git history, top authors with %, bus factor, co-change clusters, defect proximity).
MENTAL_MODEL.md — the doc the original author would write to onboard a new hire.
AGENT_BRIEF.md — ≤5 KB of assertive Never/Always rules with per-rule confidence tags. Drop-in CLAUDE.md for any project.
Ten shims into the target repo — 4 editor rule files ( CLAUDE.md , AGENTS.md , .cursor/rules/codebase.mdc , .continue/rules/codebase.md ), 5 single-intent Claude skills under .claude/skills/codebase-{exploring,debugging,impact-analysis,refactoring,onboarding}/ , and a .claude-plugin/plugin.json manifest. All write-if-absent — hand-edited files are never overwritten.
An agent-insight layer — record_insight / recall_insights MCP tools backed by <repo>/.deepdive/insights.jsonl by default (zero dependencies, human-readable, git-friendly). The optional [graphiti] extra upgrades to a temporal knowledge graph backend above a 2-of-5 repo-size threshold.
Extract also regenerates ARCHITECTURE.md — a system-level Mermaid view of the cross-boundary graph (ROUTES_TO / INJECTS / PERSISTS_TO, confidence-styled), a separate human-validation surface (not one of the five contract artifacts, exactly like forensic visualize and serve --ui ). Regenerate it on its own with forensic diagram --repo <repo> . Use it to sanity-check the graph — a wrong edge there is a wrong edge everywhere.
Add --emit-vault to also write an Obsidian -friendly vault under <output>/vault/ — every artifact gets summary: / tags: frontmatter, cross-references become [[wikilinks]] , and an INDEX.md MOC ties them together (with a .obsidian/ config). A local-first second brain for humans (graph view, backlinks) and agents (triage by summary: without opening files, a traversable index). Opt-in; off by default.
v0.8.0 "USABLE → USEFUL + public release" — the first public PyPI release. Builds on the frozen five-protocol cross-boundary graph (HTTP/MCP/registry/gRPC/messaging on one Endpoint join node) with a precision pass (honest call-graph confidence, distinct-caller counts, low-history/solo-repo guards), a human-validation ARCHITECTURE.md diagram surface, distribution (PyPI + MCP Registry + a Claude Code plugin), and an opt-in --emit-vault Obsidian export. The 5-artifact + 9-MCP-tool contract is frozen.
What's proven, and what isn't (honest framing). v0.8 is an assisted-analysis tool: a real fresh-agent onboarding test confirmed it's usable and that an agent auto-discovers AGENT_BRIEF.md and routes to the right skill unprompted, and a grounded MCP tool review found the git-archaeology + curated briefs are the high-trust core. The autonomous end-to-end question — does deepdive-seeding make an agent resolve real issues measurably faster — is not yet proven : a model-free localization pilot is recorded ( experiments/fastcontext/RESULTS.md — the static seed is a weak prior), and the end-to-end measurement is deferred to v0.9 (it needs a GPU + a frontier main-agent endpoint). No autonomous-execution claims are made. Accepted across real repos including Apache Superset, wagtail (Django), spring-petclinic, ripgrep, fastapi, and Iris-Nearby (Flutter/Dart) — see docs/findings/ .
# install from PyPI (puts `forensic` on PATH); or run ephemerally with uvx
uv tool install forensic-deepdive
forensic info # banner + capability panel
forensic extract /path/to/repo
# …or from source for development:
git clone https://github.com/Dhevenddra/forensic-deepdive && cd forensic-deepdive
uv sync --all-extras
# what can it do? (banner + capability panel: artifacts, protocols, MCP tools, confidence legend)
uv run forensic info
# run on any repo
uv run forensic extract /path/to/repo
# graph lands at <repo>/.deepdive/graph.lbug
# 5 markdown artifacts at <repo>/docs/codebase/
# 10 shims at <repo>/.claude/, .cursor/, .continue/, root
# trace a cross-stack feature slice (frontend call -> endpoint -> handler -> tail)
uv run forensic trace < symbol > --repo /path/to/repo
# query the graph as an MCP server (point it at the analyzed repo)
uv run forensic serve --repo /path/to/repo
# inspect every repo you've analyzed
uv run forensic list
Install from PyPI
Published as forensic-deepdive —
no clone needed:
uv tool install forensic-deepdive # puts `forensic` on PATH
forensic extract /path/to/repo
# …or run ephemerally, no install:
uvx forensic-deepdive extract /path/to/repo
Optional extras: uv tool install "forensic-deepdive[semantic]" (offline ONNX NL
query), [openapi] (YAML spec parsing), [graphiti] (temporal insight backend).
pip install forensic-deepdive works too if you're not on uv .
forensic serve is a stdio MCP server exposing the 9 composite tools to any
MCP-aware agent (Claude Code, Cursor, VS Code/Copilot, Codex, Continue, Cline,
Windsurf). First build the graph once ( forensic extract <repo> ), then wire the
server. Three ways, easiest first:
1. Claude Code plugin (self-hosted marketplace — no PyPI step):
/plugin marketplace add Dhevenddra/forensic-deepdive
/plugin install forensic-deepdive@dhevenddra
2. From the MCP Registry — indexed as
io.github.Dhevenddra/forensic-deepdive , so registry-aware clients and discovery hubs
(PulseMCP, MCPJungle, the VS Code @mcp index) can find and install it directly.
3. Manual config — generate a client snippet with forensic mcp-config , or paste:
{
"mcpServers" : {
"forensic-deepdive" : {
"command" : " uvx " ,
"args" : [ " forensic-deepdive " , " serve " , " --repo " , " . " ]
}
}
}
Per-client copy-paste blocks (Cursor, VS Code, Codex, the uvx -not-found GUI gotcha)
are in docs/install.md .
Python, C, Dart, Swift, TypeScript, JavaScript, Java, Go, Rust.
Tool
What it does
impact(symbol, depth, direction, min_confidence)
Blast-radius BFS over CALLS edges, depth-bucketed, confidence-filterable.
context(symbol)
Single-call kitchen sink: definition + callers + callees + parent/siblings/members + extends/implements + recent commits + dominant author + recent insights.
archaeology(file_or_symbol)
Churn, top authors with %, bus factor, co-change cluster, defect proximity, recent commits.
flow(entry_point, max_depth)
DFS over CALLS with cycle detection.
query(cypher | natural_language)
Raw Cypher, or hybrid NL retrieval (FTS5/BM25 + structural graph signal + opt-in offline semantic, RRF-fused and shaped) with per-hit provenance + confidence.
record_insight(symbol, claim, evidence, verified_by)
Persist a verified learning.
recall_insights(symbol, since, limit)
Newest-first substring match against stored insights.
visualize(target, format, depth, max_nodes, ...)
Bounded Mermaid diagram of a symbol/file neighborhood (or central ); edge dash style encodes confidence.
trace(symbol, direction, max_depth)
Cross-stack feature slice across the Endpoint join node: downstream walks frontend call → CALLS_ENDPOINT → endpoint → HANDLES → handler → CALLS tail; upstream answers "who calls this endpoint".
Tool descriptions are individually ≤200 tokens so the 9-tool envelope stays comfortably inside Anthropic's per-turn skill metadata budget.
Every edge and every emitted claim carries EXTRACTED / INFERRED / AMBIGUOUS :
EXTRACTED — deterministic from AST or git log . Facts.
INFERRED — a heuristic resolved cleanly (import-graph walk, receiver-type inference, single same-name candidate cross-file). High-trust but derived.
AMBIGUOUS — multiple candidates surfaced; the resolver couldn't disambiguate. You see every candidate , not a silent guess.
HOTPATHS shows a per-row confidence-mix column so at a glance you can tell Logger (4 EXTRACTED + 1458 INFERRED — mostly clean) from ChatToolResponse (449 AMBIGUOUS — same-name cross-file collision).
Honest-mode (pure-static, zero LLM, zero network)
forensic extract works end-to-end with no ANTHROPIC_API_KEY , no OPENAI_API_KEY , no Ollama, no network . Graphiti is opt-in via the [graphiti] PyPI extra plus a 2-of-5 repo-size threshold (≥50 k LOC, ≥25 contributors, ≥18 mo old, ≥200 PRs/12 mo, ≥100 issues with discussion). The JsonlInsightStore is the always-available floor.
Why this and not [GitNexus / CodeGraphContext / DeepWiki / Sourcegraph]
forensic-deepdive
GitNexus
CodeGraphContext
DeepWiki
Sourcegraph

[truncated]
