---
source: "https://github.com/har-ki/claude-code-sre-handbook/tree/main/watcher-semantic-example"
hn_url: "https://news.ycombinator.com/item?id=48556174"
title: "Show HN: K8s watcher that investigates incidents and opens PRs (it can't merge)"
article_title: "claude-code-sre-handbook/watcher-semantic-example at main · har-ki/claude-code-sre-handbook · GitHub"
author: "har-ki"
captured_at: "2026-06-16T16:38:17Z"
capture_tool: "hn-digest"
hn_id: 48556174
score: 2
comments: 0
posted_at: "2026-06-16T14:50:44Z"
tags:
  - hacker-news
  - translated
---

# Show HN: K8s watcher that investigates incidents and opens PRs (it can't merge)

- HN: [48556174](https://news.ycombinator.com/item?id=48556174)
- Source: [github.com](https://github.com/har-ki/claude-code-sre-handbook/tree/main/watcher-semantic-example)
- Score: 2
- Comments: 0
- Posted: 2026-06-16T14:50:44Z

## Translation

タイトル: Show HN: インシデントを調査し PR を開く K8s ウォッチャー (マージできません)
記事のタイトル: claude-code-sre-handbook/watcher-semantic-example at main · har-ki/claude-code-sre-handbook · GitHub
説明: Claude Code SRE ハンドブック。 GitHub でアカウントを作成して、har-ki/claude-code-sre-handbook の開発に貢献してください。

記事本文:
メインの claude-code-sre-handbook/watcher-semantic-example · har-ki/claude-code-sre-handbook · GitHub
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
ハルキ
/
クロード・コード・スレ・ハンドブック
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

コード
その他のオプション ディレクトリアクション
歴史 歴史メイン ブレッドクラム
コピーパスのトップフォルダーとファイル
.. アラートウォッチャー アラートウォッチャー 監査ログ 監査ログ claude-runner claude-runner メモリストア メモリストア プロンプト プロンプト 結果 結果 スクリプト スクリプト テンプレート テンプレート テストインシデント テストインシデント .gitignore .gitignore README.md README.md docker-compose.yml docker-compose.yml run-matrix.sh run-matrix.sh seed-memory.sh seed-memory.sh verify-retrieval.sh verify-retrieval.sh すべてのファイルを表示 README.md
アウトライン ウォッチャーのセマンティックの例
SRE インシデント ウォッチャーのセマンティック メモリ バリアント。類似性に基づいて追加します
取得 (ローカル埋め込みによる) と検証規律を切り替え可能にします
watcher-memory-example/ からのファイルベースのメモリ規則の上に層を置きます。
自己完結型。他のウォッチャー サンプルからインポートしたり、他のウォッチャー サンプルを変更したりしません。
モデル: nomic-embed-text (274 MB、768 次元)
ランタイム: Ollama (ローカル、エアギャップ互換)
類似性スコアはモデル固有です。しきい値の校正
このモデルは他の埋め込みモデルに転送できません。
すべてalert-watcher/config.yamlで設定されており、コード編集は必要ありません。
# 取得モード: 完全一致 (SHA256 ハッシュ検索) または類似度 (コサイン検索)
retrieval_mode : 類似度 # 正確 |類似性
# 所見を明らかにするための最小コサイン類似度 (類似度モードのみ)
類似性閾値 : 0.75
# true の場合、フェーズ 1 は以前の結果を次のように分類する必要があります。
# 使用前に検証/無効化/不確定
validation_discipline : true # true |偽
建築
ウォッチャー-セマンティック-例/
§── アラートウォッチャー/
│ §── main.py # 構成可能な取得を備えたオーケストレーター
│ §── semantic_memory.py # 埋め込み、インデックス付け、取得層
│ └─ config.yaml # 3 つのトグル + 標準設定
§── クロード・ランナー/

│ §── invoke.sh # 2 つの別々の claude -p フェーズ
│ └── workspace-claude.md # モデルの向き (メモリドキュメントを含む)
§── プロンプト/
│ §── 01-investigate.md # フェーズ 1: メタデータの取得 + 検証規律
│ └─ 02-propose-fix.md # フェーズ 2: 修正 + 必要なメモリ書き込み
§── メモリストア/
│ §── インシデント/ # インシデントごとの調査結果 (<sha8>.md)
│ └──embeddings/ # ベクターインデックス(index.json)
§── test-incidents/ # 3 つのテスト ケース (以下を参照)
§── seed-memory.sh # メモリをシード + ベクトルインデックスを構築
└── verify-retrieval.sh # すべてのテスト ケースに対して取得を検証します
ロックされた不変条件 (watcher-memory-example と同じ)
2 つの別々のクロード -p フェーズ (決して崩壊しない)
--max-turns 40 、 --output-format stream-json --verbose
gh prready がすべての --allowedTools に存在しない (ヒューマン ゲート)
regex + realpath() コンテインメントによってメモリ パスが検証される
watcher-memory-example/ と同じ。 Memory-store/incidents/<sha8>.md を検索します。
指紋ハッシュによって。同一の service|Exception_class でのみヒットします。
nomic-embed-text を介してインシデントの説明を埋め込み、ベクトルを検索します
コサイン類似度によるインデックス。スコアが次を超えた場合に最も近い結果を返します。
類似性のしきい値 ;それ以外の場合はサイレントのままになります (セーフサイレントのデフォルト)。
インデックス形式: Memory-store/embeddings/index.json にあるフラット JSON ファイル。
各エントリには、埋め込み、サービス、例外クラス、およびフィンガープリント ハッシュが保存されます。
数百の結果には十分です (数百万を対象として設計されていません)。
保存された検索結果では、search_document: 検索内容の接頭辞が使用されます (リッチ)
クエリでは、拡張された例外コンテキストを含む search_query: プレフィックスを使用します。
CamelCase 例外名はセマンティック信号用に分割されます
(StockMismatchError → 在庫不一致エラー)
validation_discipline: true の場合、フェーズ 1 のプロンプトには

明示的な
事前の結果を使用する前の分類:
> **検証: VALIDATED** — 以前の検索結果が一致するため...
> **検証: INVALIDATED** — 以前の結果は一致しません。理由は...
> **検証: 結論が出ていません** — 確認できず、さらに調査中
INVALIDATED の場合、エージェントは最初から調査する必要があります。これは設計されています
部分一致が原因で発生する「リコールアンカー」障害モードを捕捉する
エージェントは調査をあまりにも早く中止する必要があります。
test-incident/ 内の 3 つのテスト ケース:
検証済みの検索スコア (nomic-embed-text)
テスト
正確な
類似性スコア
0.75以上？
正規
ヒット
0.80
はい
ドリフト
ミス
0.81
はい
ニアミス
ヒット
0.80
はい
前例のない
ミス
0.74
いいえ (沈黙)
クイックスタート
#1. Ollama が埋め込みモデルで実行されていることを確認する
オラマプル nomic-embed-text
# 2. 正規の検出結果とビルド ベクトル インデックスを使用したシード メモリ
bash シードメモリ.sh
# 3. すべてのテスト インシデントに対して取得が機能することを確認する
bash verify-retrieval.sh
# 4. ウォッチャーを実行します (otel-demo クラスター + Docker が必要)
cp .env.example .env
# ANTHROPIC_API_KEY と GH_TOKEN を入力します
docker 構成 -d --build
実行間のモードの切り替え
alert-watcher/config.yaml を編集します。
# 実行 1: 正確なハッシュのベースライン
取得モード : 正確
validation_discipline : false
# 実行 2: 規律との類似性
retrieval_mode : 類似度
validation_discipline : true
# 実行 3: 規律のない類似性 (規律が重要かどうかを確認するため)
retrieval_mode : 類似度
validation_discipline : false
コードの変更やリビルドは必要ありません。構成はウォッチャーの起動時に読み取られます。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Claude Code SRE Handbook. Contribute to har-ki/claude-code-sre-handbook development by creating an account on GitHub.

claude-code-sre-handbook/watcher-semantic-example at main · har-ki/claude-code-sre-handbook · GitHub
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
har-ki
/
claude-code-sre-handbook
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
More options Directory actions
History History main Breadcrumbs
Copy path Top Folders and files
.. alert-watcher alert-watcher audit-log audit-log claude-runner claude-runner memory-store memory-store prompts prompts results results scripts scripts templates templates test-incidents test-incidents .gitignore .gitignore README.md README.md docker-compose.yml docker-compose.yml run-matrix.sh run-matrix.sh seed-memory.sh seed-memory.sh verify-retrieval.sh verify-retrieval.sh View all files README.md
Outline watcher-semantic-example
Semantic-memory variant of the SRE incident watcher. Adds similarity-based
retrieval (via local embeddings) and validation discipline as toggleable
layers on top of the file-based memory convention from watcher-memory-example/ .
Self-contained. Does not import from or modify the other watcher examples.
Model: nomic-embed-text (274 MB, 768-dimensional)
Runtime: Ollama (local, air-gapped compatible)
Similarity scores are model-specific. Threshold values calibrated for
this model are not transferable to other embedding models.
All set in alert-watcher/config.yaml , no code edits needed:
# Retrieval mode: exact (SHA256-hash lookup) or similarity (cosine search)
retrieval_mode : similarity # exact | similarity
# Minimum cosine similarity to surface a finding (similarity mode only)
similarity_threshold : 0.75
# When true, Phase 1 must classify the prior finding as
# validated / invalidated / inconclusive before using it
validation_discipline : true # true | false
Architecture
watcher-semantic-example/
├── alert-watcher/
│ ├── main.py # Orchestrator with configurable retrieval
│ ├── semantic_memory.py # Embedding, indexing, retrieval layer
│ └── config.yaml # Three toggles + standard config
├── claude-runner/
│ ├── invoke.sh # Two separate claude -p phases
│ └── workspace-claude.md # Model orientation (includes memory docs)
├── prompts/
│ ├── 01-investigate.md # Phase 1: retrieval metadata + validation discipline
│ └── 02-propose-fix.md # Phase 2: fix + required memory write
├── memory-store/
│ ├── incidents/ # Per-incident findings (<sha8>.md)
│ └── embeddings/ # Vector index (index.json)
├── test-incidents/ # Three test cases (see below)
├── seed-memory.sh # Seeds memory + builds vector index
└── verify-retrieval.sh # Verifies retrieval against all test cases
Locked invariants (same as watcher-memory-example)
Two separate claude -p phases (never collapsed)
--max-turns 40 , --output-format stream-json --verbose
gh pr ready absent from all --allowedTools (human gate)
Memory paths validated via regex + realpath() containment
Same as watcher-memory-example/ . Looks up memory-store/incidents/<sha8>.md
by the fingerprint hash. Hits only on identical service|exception_class .
Embeds the incident description via nomic-embed-text , searches the vector
index by cosine similarity. Returns the nearest finding if the score exceeds
similarity_threshold ; stays silent otherwise (safe-silence default).
Index format: Flat JSON file at memory-store/embeddings/index.json .
Each entry stores the embedding, service, exception class, and fingerprint hash.
Sufficient for hundreds of findings (not designed for millions).
Stored findings use search_document: prefix with finding content (rich)
Queries use search_query: prefix with expanded exception context
CamelCase exception names are split for semantic signal
( StockMismatchError → stock mismatch error )
When validation_discipline: true , Phase 1's prompt requires an explicit
classification before using the prior finding:
> **Validation: VALIDATED** — prior finding matches because ...
> **Validation: INVALIDATED** — prior finding does NOT match because ...
> **Validation: INCONCLUSIVE** — cannot confirm, investigating further
If INVALIDATED, the agent must investigate from scratch. This is designed
to catch the "recall-anchoring" failure mode where a partial match causes
the agent to stop investigating too early.
Three test cases in test-incidents/ :
Verified retrieval scores (nomic-embed-text)
Test
Exact
Similarity score
Above 0.75?
Canonical
hit
0.80
yes
Drift
miss
0.81
yes
Near-miss
hit
0.80
yes
No-precedent
miss
0.74
no (silent)
Quick start
# 1. Ensure Ollama is running with the embedding model
ollama pull nomic-embed-text
# 2. Seed memory with the canonical finding + build vector index
bash seed-memory.sh
# 3. Verify retrieval works for all test incidents
bash verify-retrieval.sh
# 4. Run the watcher (requires otel-demo cluster + Docker)
cp .env.example .env
# Fill in ANTHROPIC_API_KEY and GH_TOKEN
docker compose up -d --build
Switching modes between runs
Edit alert-watcher/config.yaml :
# Run 1: exact-hash baseline
retrieval_mode : exact
validation_discipline : false
# Run 2: similarity with discipline
retrieval_mode : similarity
validation_discipline : true
# Run 3: similarity without discipline (to see if discipline matters)
retrieval_mode : similarity
validation_discipline : false
No code changes, no rebuild. The config is read at watcher startup.
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
