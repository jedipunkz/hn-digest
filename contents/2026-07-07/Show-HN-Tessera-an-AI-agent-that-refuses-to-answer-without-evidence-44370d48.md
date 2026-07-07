---
source: "https://github.com/robert-vetter/tessera"
hn_url: "https://news.ycombinator.com/item?id=48822458"
title: "Show HN: Tessera – an AI agent that refuses to answer without evidence"
article_title: "GitHub - robert-vetter/tessera: A trust layer for enterprise AI agents: answers traceable to their evidence, with a score for how much to trust them. · GitHub"
author: "robert-vetter"
captured_at: "2026-07-07T19:42:43Z"
capture_tool: "hn-digest"
hn_id: 48822458
score: 2
comments: 0
posted_at: "2026-07-07T19:24:07Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Tessera – an AI agent that refuses to answer without evidence

- HN: [48822458](https://news.ycombinator.com/item?id=48822458)
- Source: [github.com](https://github.com/robert-vetter/tessera)
- Score: 2
- Comments: 0
- Posted: 2026-07-07T19:24:07Z

## Translation

タイトル: 番組 HN: Tessera – 証拠がなければ回答を拒否する AI エージェント
記事のタイトル: GitHub - robert-vetter/tessera: エンタープライズ AI エージェントのための信頼層: 証拠まで追跡可能な回答と、どの程度信頼できるかのスコア。 · GitHub
説明: エンタープライズ AI エージェントの信頼レイヤー: 証拠に追跡可能な回答と、どの程度信頼できるかのスコア。 - ロバート・ヴェッター/テッセラ

記事本文:
GitHub - robert-vetter/tessera: エンタープライズ AI エージェントのための信頼層: 証拠に追跡可能な回答と、どの程度信頼できるかのスコア。 · GitHub
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
ロバート・ヴェッター
/
テッセラ
公共
通知
通知を変更するにはサインインする必要があります

フィクション設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
160 コミット 160 コミット .claude .claude .devcontainer .devcontainer .github .github data datadeploy/ hf-spacedeploy/ hf-space docs docs eval eval launch launch scripts scripts spec specs src/ tessera src/ tessera テスト テスト .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md Dockerfile Dockerfile ライセンス ライセンス README.md README.md mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
エージェントは証明できることしか言えず、あなたが承認したことだけを実行します。
AI エージェント用のオープンで決定論的な証拠レイヤー。すべての答えは、その背後にある正確なソースレコードまで遡るという主張に基づいて構築されています。証明できないものは推測されず、拒否されます。エージェントが実行するすべてのアクションは、検証済みの申し立てに基づいてのみ草案され、正確にプレビューされ、受領書とともに承認された場合にのみ実行されます。忠実さは CI における数値であり、雰囲気ではありません。
▶ ライブで試す (サインアップなし、読み取り専用): https://robert-vetter-tessera.hf.space — 質問し、各請求の検証チップを確認し、クリックして証拠にアクセスし、拒否を確認し、領収書に対するアクションを実行します。
📊 ベンチマーク — 証拠ゲート付きとゲートなしの同じエンジン、同じコーパスと検証器、LLM ジャッジなし: docs/BENCHMARK.md 。
🔌 MCP ネイティブ — エージェントの証拠としてプラグインします。 oracle: uv run tessera-mcp 。
テッセラはモザイク内の 1 つのタイルです。多くの小さく異質な部分が 1 つの一貫した検証可能な画像に組み立てられます。これが、このシステムが企業データで行うことです。
エンタープライズAIラ

言語モデルが弱いため、依存は失敗します。このシステムは、互換性のないソースを黙って混合し、もっともらしく聞こえる詳細をでっち上げ、主張がどこから来たのかを遡って追跡する方法を提供せず、ましてやそれがどれほど忠実であるかを測定することはできません。 Tessera は、異種のエンタープライズ データ (データベース テーブル、スプレッドシート、ドキュメント、ログ、チケット) を取り込み、それらのソースに散在するエンティティを単一の統一されたナレッジ グラフに解決し、すべての個々の主張がそれをサポートするソース レコードに遡る追跡可能なパスに基づいた会話型インターフェイスを通じて質問に答えるオープン フレームワークです。それに加えて、Tessera は「この AI を信頼しますか?」を判断するベンチマーク ハーネスを出荷しています。感覚を指標に変換します。忠実度スコアは、合成および厳選されたデータに基づいて測定され、システムが改善されるにつれて上昇するのを観察できます。
2025 年、AI エージェントがコードの凍結中に実稼働データベースを削除し、データを捏造してロールバックが可能かどうかを誤って報告しました。業界の事後コンセンサス (予行演習、承認ゲート、ブラスト半径プレビュー、ドラフトのみの認証情報) が Tessera のアクション モデルです。それに加えて、テッセラはほとんどのツールのスキップを追加します。つまり、エージェントの行動だけでなく、エージェントの発言もゲートされます。証拠への追跡可能な経路がなければ主張は禁止され、証拠がない場合は拒否されます。すべてが決定的であり、信頼パス内にモデル ベンダーが存在しないため、保証はオンプレミスで保持され、モデルが変更されても移動しません。 (オプションの LLM は答えを説明する場合がありますが、決して証明するものではありません。)
Tessera が行うこと — 4 つの柱
普遍的な摂取。構造化ソース (リレーショナル テーブル、CSV、エクスポート) と非構造化ソース (ドキュメント、ランブック、ログ、チケット、チャット スレッド) を 1 つの一貫したソースを通じて取り込みます。

吸気経路。
ソース間のエンティティ解決 → 1 つのナレッジ グラフ。データベース内の「Acme Corp」、契約 PDF 内の「ACME」、および展開ログ内の「acme-prod」が同じエンティティであることを認識し、すべてのソースを 1 つのグラフとしてクエリできる単一のグラフに織り込みます。
出所を踏まえた、根拠のある会話的推論。自然言語で質問に回答します。この場合、システムは単純な検索と複数ステップの推論を異なる方法でルーティングし、回答内のすべての文は、グラフを介して基礎となるレコードへの引用パスによって裏付けられます。追跡不能な申し立てはありません。
数値化された信頼。合成データの生成と厳選されたゴールド セットによって駆動される組み込みの評価ハーネスは、忠実度 (すべての主張が証拠によって裏付けられているか)、網羅率、および回答の質をスコアリングするため、改善は雰囲気ではなく測定可能です。
2 つの基準業種、1 つのエンジン
同じコア エンジンが 2 つの意図的に異なるデモンストレーションを実行し、エンジンが一般化されることを証明します。
ビジネス データ コパイロット — 企業の構造化された記録と文書全体にわたって質問し、根拠のある引用された回答を得ることができます。 — UV ラン テッセラ
DevEx Copilot — CI/CD ログ、プルリクエストの差分、チケット履歴を指します。失敗したパイプラインの根本原因の仮説と、変更が実際に何を行うかについての根拠のある要約を取得します。 — UV 実行 tessera-devex
フェーズ 3 の時点で、両方のバーティカルが実行され、両方の測定が行われます。
ファイルはフェーズ 2 タグとフェーズ 3 タグの間でバイトが同一です
( ADR 0008 ): 一般化主張は次のとおりです
Promise ではなく、空の git diff です。フェーズ 4 で対称性が完成しました。
ビジネス回答レイヤーは tessera/devex/ の隣の tessera/business/ に存在します。
それぞれの業界が独自のクレーム文法を所有しています
( ADR 0011 ) — と追加しました
両方を使用したジュール スタイルのセッション:
UV 実行テッセラチャット
# 会話ドア 1 つ: expl

ainable ルーティング、番号付きクレーム、:show N to
# 来歴 (レコード、ロケーター、解決アサーション) を歩きます。
# 記録された数値 — すべての回答をライブ検証者がチェックします。
何が本当に難しいのか（正直なバージョン）
接地された RAG が存在します。ナレッジグラフが存在します。エンティティ解決が存在します。珍しいのは、そしてこのプロジェクトが実際に何を目的としているのかということですが、統一された来歴モデルと、全体に責任を負わせる忠実度の指標を使用して、構造化データと非構造化データにまたがるこれらすべてを同時に実行することです。ほとんどのシステムは 1 つのモダリティを選択し、出所をスキップし、忠実性をまったく測定しません。テッセラは、測定可能な信頼を後付けの考えではなく、見出しの特徴として扱います。
証拠ゲート型アクション - MCP 経由
エージェントはモデル コンテキスト プロトコル ( uv run tessera-mcp 、
オプトインエージェントは追加）、信託契約は、
4 つの境界を越えるプロトコル:
読み取り専用の根拠のあるツール — 主張、出所、原則に基づく拒絶
( ADR 0022 )。
行動草案の提案と承認 — 根本原因からのインシデント
分析、変更からの pr_summary。各フィールドは、
検証者通過クレーム、または提案されていません
( ADR 0023 )。
ペイロードの予行演習 - 正確な GitHub リクエストとドラフトされたアクション
すべての値が検証されたフィールドにトレースされて送信されます
( ADR 0024 )。
承認後の効果的な実行 — 完全に接地されたペイロードに基づいて制御され、
オプトインを使用して、デフォルトでは何も送信しないシミュレートされたアクチュエータによって実行されます。
資格情報と承認の背後にある実際のパス
( ADR 0025 )。
Tessera はドラフト、検証、レンダリング、シミュレーションを行います。こちらからは何も送信されません
リポジトリ — 人間またはエージェントが承認して送信します。 (正確に 1 回送信されましたが、
記録上: 承認の背後にある、根拠のあるインシデントによる実際の GitHub の問題、
と

確約された受領書。）
SAP を中心とするのではなく、SAP に向けて構築されています。決定的なオンプレミスの姿勢は、
ポイント: SAP AI コア/Generative AI Hub (モデルアダプター) 上で実行されます。
契約でテスト済み。 docs/DEPLOYMENT.md の Runbook 、
ADR 0012 )、SAP HANA Cloud に保持されます。
(データベース内の埋め込みはオンラインで記録されます。ナレッジ グラフ エンジンのミラーは
契約でテスト済み）、実際のゲート SALT データセットに基づいた答えを根拠としています
( docs/SALT_REAL.md )、ジュール スタイルのサーフェスを提供します
( uv run tessera-chat ) オプションのナレーションで事実を追加することはできません
( ADR 0013 )。チームごとの完全版
マッピングは docs/SAP_ALIGNMENT.md にあります。
このプロジェクトでは Python 環境に uv を使用しており、
依存関係の管理。新しいクローンから:
uv sync # uv.lockから環境を作成
uv run pre-commit install # one-time: ローカル コミット ゲートを有効にする
完全な品質ゲートを実行します。CI が実行する正確なチェック (フォーマット、lint、厳密なチェック)
タイプ、テスト) — bash scripts/gate.sh を使用します。 eval は uv run tessera-eval です
そして、シークレットスキャン UV 実行 pre-commit run gitleaks --all-files 。
ツールチェーンがありませんか?コンテナを使用する
Dockerfile と devcontainer は、
Docker 以外のホスト設定がない、クリーンで固定された環境:
docker build -t tessera-dev 。 # 固定されたイメージを構築する
docker run --rm tessera-dev pytest # 内部のテストスイートを実行します
または、VS Code / GitHub コードスペースでフォルダーを開いて「コンテナーで再度開く」 —
環境は自動的に構築され、同期されます。
それを確認する最も速い方法は、1 ページの Web サーフェスです。質問して、情報を入手してください。
請求ごとの検証チップを使用した請求、クリックスルーで正確な証拠にアクセス、監視
原則に基づく拒否であり、ドラフト→ペイロードの予行演習→承認→というアクションを実行します。
領収書 (すべてシミュレートされています。資格情報は保持されません):
ライブ、インストールなし: https://robert-vetter-tessera.hf.space — 同じ
表面、Hugging Fac でキーフリーでホスト

eスペース
(デプロイ/hf-space/)。
uv run tessera-ui # → http://127.0.0.1:8033 (stdlib のみ、JS なし)
これをホストするには (キー不要、公開読み取り専用)、3 分間のデモ スクリプトと概要資料を使用します。
docs/DEMO.md 。表面は同じものの厳格な消費者です
MCP サーバーが公開する信頼オブジェクト (ADR 0027)。
これらのツールを使用して同じ 3 つのことを行っている本物のクロード エージェントは、次の場所に記録されています。
データ/agent_session/TRANSCRIPT.md 。
自分のリポジトリ上 (マイルストーン 18、
ADR 0028 ): テッセラを指す
パブリック GitHub リポジトリの CI 履歴を調べて、実行が失敗した理由を尋ねる - RCA を接地
ローカルの、gitignored、スクラブされたスナップショットを介した同じクレームレベルの出所
(答えはオフラインです。オプションのスコープなしの GITHUB_TOKEN により、失敗した実行のロックが解除されます
ログ — .env.example を参照):
uv run tessera connect github astral-sh/uv # 境界スナップショット → var/connect/
# connect は、スナップショットから実際に失敗した実行 ID を使用して実行準備完了の `ask` を出力します
# (実行 ID は GitHub の約 90 日間のログ保持期間で期限切れになります。推奨されたものを使用してください):
uv run tessera ask astral-sh/uv 「<run-id> の実行が失敗したのはなぜですか?」
独自の CSV + マークダウン (マイルストーン 18、
ADR 0029 ): ディレクトリを次のように記述します。
小さな tessera.toml を作成し、同じものを接地します。
出所を伝える答え — 複数分野の実体解決と誠実な対応による
名前があいまいな場合は拒否します。献身的なパブリックドメイン

[切り捨てられた]

## Original Extract

A trust layer for enterprise AI agents: answers traceable to their evidence, with a score for how much to trust them. - robert-vetter/tessera

GitHub - robert-vetter/tessera: A trust layer for enterprise AI agents: answers traceable to their evidence, with a score for how much to trust them. · GitHub
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
robert-vetter
/
tessera
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
160 Commits 160 Commits .claude .claude .devcontainer .devcontainer .github .github data data deploy/ hf-space deploy/ hf-space docs docs eval eval launch launch scripts scripts specs specs src/ tessera src/ tessera tests tests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
The agent can only say what it can prove — and only do what you approve.
An open, deterministic evidence layer for AI agents . Every answer is built from claims that each trace back to the exact source record behind them; what can't be proven is refused, not guessed. Every action an agent takes is drafted only from verified claims, previewed exactly, and executed only behind your approval — with a receipt. Faithfulness is a number in CI, not a vibe.
▶ Try it live (no signup, read-only): https://robert-vetter-tessera.hf.space — ask a question, see each claim's verifier chip, click through to the evidence, watch a refusal, run an action to a receipt.
📊 The benchmark — evidence-gated vs. the same engine ungated , same corpus and verifier, no LLM judge: docs/BENCHMARK.md .
🔌 MCP-native — plug it in as your agent's evidence oracle: uv run tessera-mcp .
A tessera is a single tile in a mosaic: many small, heterogeneous pieces assembled into one coherent, verifiable picture — which is what this system does with enterprise data.
Enterprise AI rarely fails because the language model is weak. It fails because nobody can trust the answer: the system silently mixes incompatible sources, invents details that sound plausible, and offers no way to trace a claim back to where it came from — let alone to measure how faithful it was. Tessera is an open framework that ingests heterogeneous enterprise data (database tables, spreadsheets, documents, logs, tickets), resolves the entities scattered across those sources into a single unified knowledge graph, and answers questions through a conversational interface in which every individual claim is grounded in a traceable path back to the source records that support it . On top of that, Tessera ships a benchmark harness that turns "do you trust this AI?" from a feeling into a metric — a faithfulness score, measured on synthetic and curated data, that you can watch go up as the system improves.
In 2025 an AI agent deleted a production database during a code freeze, then fabricated data and misreported whether a rollback was possible. The industry's post-mortem consensus — dry-runs, approval gates, blast-radius previews, draft-only credentials — is Tessera's action model. On top of it, Tessera adds the half most tools skip: an agent's statements are gated too, not just its actions — no claim without a traceable path to evidence, and a refusal where there is none. All of it deterministic, with no model vendor inside the trust path , so the guarantees hold on-prem and do not move when a model changes. (An optional LLM may narrate an answer; it never attests .)
What Tessera does — four pillars
Universal ingestion. Pull in structured sources (relational tables, CSVs, exports) and unstructured ones (documents, runbooks, logs, tickets, chat threads) through one consistent intake path.
Cross-source entity resolution → one knowledge graph. Recognize that "Acme Corp" in a database, "ACME" in a contract PDF, and "acme-prod" in a deployment log are the same entity, and weave all sources into a single graph that can be queried as one.
Grounded conversational reasoning with provenance. Answer questions in natural language where the system routes simple lookups and multi-step reasoning differently, and where every sentence in the answer is backed by a citation path through the graph to the underlying records. No untraceable claims.
Quantified trust. A built-in evaluation harness — driven by synthetic data generation and a curated gold set — that scores faithfulness (is every claim supported by evidence?), coverage , and answer quality , so improvements are measurable rather than vibes.
Two reference verticals, one engine
The same core engine powers two deliberately different demonstrations, to prove the engine generalizes:
Business Data Copilot — ask questions across a company's structured records and documents and get grounded, cited answers. — uv run tessera
DevEx Copilot — point it at CI/CD logs, pull-request diffs, and ticket history; get root-cause hypotheses for failed pipelines and grounded summaries of what a change actually does. — uv run tessera-devex
As of Phase 3 both verticals run and both are measured — on a core whose
files are byte-identical between the phase-2 and phase-3 tags
( ADR 0008 ): the generalization claim is
an empty git diff , not a promise. Phase 4 completed the symmetry — the
business answer layer lives in tessera/business/ beside tessera/devex/ ,
each vertical owns its claim grammars
( ADR 0011 ) — and added the
Joule-style session over both :
uv run tessera-chat
# one conversational door: explainable routing, numbered claims, :show N to
# walk provenance (records, locators, resolution assertions), :trust for the
# recorded numbers — and a live verifier check on every answer.
What makes it genuinely hard (the honest version)
Grounded RAG exists. Knowledge graphs exist. Entity resolution exists. What is rare — and what this project is actually about — is doing all of them together, across structured and unstructured data at once, with a uniform provenance model and a faithfulness metric that holds the whole thing accountable. Most systems pick one modality, skip provenance, and never measure faithfulness at all. Tessera treats measurable trust as the headline feature, not an afterthought.
Evidence-gated actions — over MCP
An agent talks to Tessera over the Model Context Protocol ( uv run tessera-mcp ,
the opt-in agent extra), and the trust contract is measured to survive the
protocol across four boundaries:
Read-only grounded tools — claims, provenance, and principled refusals
( ADR 0022 ).
Propose-and-approve action drafts — an incident from a root-cause
analysis, a pr_summary from a change; each field traces to a
verifier-passing claim or it is not proposed
( ADR 0023 ).
A dry-run payload preview — the exact GitHub request a drafted action
would send, every value traced to a verified field
( ADR 0024 ).
Effectful execution behind approval — gated on a fully-grounded payload,
run by a simulated actuator that sends nothing by default, with an opt-in
real path behind a credential and an approval
( ADR 0025 ).
Tessera drafts, verifies, renders, and simulates; nothing is sent from this
repository — a human or agent approves and sends. (It did send exactly once,
on the record: a real GitHub issue from a grounded incident, behind approval,
with a committed receipt.)
Built toward SAP, not around it. The deterministic, on-prem posture is the
point: it runs on SAP AI Core / Generative AI Hub (model adapter
contract-tested; runbook in docs/DEPLOYMENT.md ,
ADR 0012 ), persists to SAP HANA Cloud
(in-database embeddings recorded online; the knowledge-graph-engine mirror is
contract-tested), grounds answers on the real gated SALT dataset
( docs/SALT_REAL.md ), and offers a Joule-style surface
( uv run tessera-chat ) whose optional narration can never add facts
( ADR 0013 ). The full team-by-team
mapping is in docs/SAP_ALIGNMENT.md .
The project uses uv for Python environment and
dependency management. From a fresh clone:
uv sync # create the environment from uv.lock
uv run pre-commit install # one-time: enable local commit gates
Run the full quality gate — the exact checks CI runs (format, lint, strict
types, tests) — with bash scripts/gate.sh ; the eval is uv run tessera-eval
and the secret scan uv run pre-commit run gitleaks --all-files .
No toolchain? Use the container
A Dockerfile and a devcontainer provide a
clean, pinned environment with no host setup beyond Docker:
docker build -t tessera-dev . # build the pinned image
docker run --rm tessera-dev pytest # run the test suite inside it
Or open the folder in VS Code / GitHub Codespaces and "Reopen in Container" —
the environment builds and syncs automatically.
The fastest way to see it is the one-page web surface — ask a question, get
claims with per-claim verifier chips, click through to the exact evidence, watch
a principled refusal, and run an action from draft → dry-run payload → approval →
receipt (all simulated; it holds no credential):
Live, no install: https://robert-vetter-tessera.hf.space — the same
surface, hosted key-free on a Hugging Face Space
( deploy/hf-space/ ).
uv run tessera-ui # → http://127.0.0.1:8033 (stdlib-only, zero JS)
To host it (key-free, public read-only) and the 3-minute demo script + one-pager:
docs/DEMO.md . The surface is a strict consumer of the same
trust objects the MCP server exposes ( ADR 0027 );
a real Claude agent doing the same three things through those tools is recorded at
data/agent_session/TRANSCRIPT.md .
On your own repository (Milestone 18,
ADR 0028 ): point Tessera at
any public GitHub repo's CI history and ask why a run failed — grounded RCA with
the same claim-level provenance, over a local, gitignored, scrubbed snapshot
(answers are offline; an optional no-scope GITHUB_TOKEN unlocks failed-run
logs — see .env.example ):
uv run tessera connect github astral-sh/uv # bounded snapshot → var/connect/
# connect prints a ready-to-run `ask` with a real failed-run id from the snapshot
# (run ids expire with GitHub's ~90-day log retention — use the suggested one):
uv run tessera ask astral-sh/uv " Why did run <run-id> fail? "
On your own CSV + Markdown (Milestone 18,
ADR 0029 ): describe a directory with a
small tessera.toml and get the same grounded,
provenance-carrying answers — with multi-field entity resolution and an honest
refusal when a name is ambiguous. A committed public-doma

[truncated]
