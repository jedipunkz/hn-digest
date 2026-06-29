---
source: "https://github.com/aivinay/switchboard"
hn_url: "https://news.ycombinator.com/item?id=48725764"
title: "Show HN: Switchboard – route AI prompts instead of capping budgets"
article_title: "GitHub - aivinay/switchboard: Privacy-aware, local-first router for your CLI coding agents (Codex, Claude Code) and local LLMs (Ollama) — keeps sensitive prompts on-device and cuts premium-model usage. · GitHub"
author: "ai_vinaygupta"
captured_at: "2026-06-29T22:23:04Z"
capture_tool: "hn-digest"
hn_id: 48725764
score: 2
comments: 0
posted_at: "2026-06-29T21:50:23Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Switchboard – route AI prompts instead of capping budgets

- HN: [48725764](https://news.ycombinator.com/item?id=48725764)
- Source: [github.com](https://github.com/aivinay/switchboard)
- Score: 2
- Comments: 0
- Posted: 2026-06-29T21:50:23Z

## Translation

タイトル: Show HN: Switchboard – 予算に上限を設ける代わりに AI プロンプトをルーティングする
記事のタイトル: GitHub - aivinay/switchboard: CLI コーディング エージェント (Codex、Claude Code) およびローカル LLM (Ollama) 用のプライバシーを意識したローカル ファースト ルーター — デバイス上で機密性の高いプロンプトを維持し、プレミアム モデルの使用を削減します。 · GitHub
説明: CLI コーディング エージェント (Codex、Claude Code) およびローカル LLM (Ollama) 用のプライバシーを意識したローカル ファースト ルーター - デバイス上で機密性の高いプロンプトを維持し、プレミアム モデルの使用を削減します。 - アイビネイ/配電盤

記事本文:
GitHub - aivinay/switchboard: CLI コーディング エージェント (Codex、Claude Code) およびローカル LLM (Ollama) 用のプライバシーを意識したローカル ファースト ルーター — デバイス上で機密性の高いプロンプトを維持し、プレミアム モデルの使用を削減します。 · GitHub
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
アイビネイ
/
配電盤
出版

c
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
15 コミット 15 コミット .github .github アセット アセット 構成 構成ドキュメント ドキュメント スクリプト スクリプト 配電盤 スイッチボード テスト テスト .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md RELEASE.md RELEASE.md SECURITY.md SECURITY.md auto_route_demo.gif auto_route_demo.gif docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
プレミアム エージェントのコールが 62% 削減 · 4.1/5 の品質と 4.6/5 の常時プレミアム · ベンチマーク リークは観察されません
インストール・
評価・
仕組み ·
プライバシー ·
紙・
ドキュメント
1 つのセッション、3 つのバックエンド: デフォルトではローカル、コードには Codex、推論には Claude Code。
Switchboard は、すでに使用している CLI ツールをラップし (別個のサービス、プロキシ、再販された API アクセスは必要ありません)、学習された分類子が実行される前に、決定論的なルールで各プロンプトをルーティングします。
100 件のベンチマークでは、Switchboard はリクエストの 62% をプレミアムから除外しました。
エージェントの品質は 4.6/5 の常時プレミアムに対して 4.1/5 に達します
ベースライン。100% 回答され、ベンチマーク リークは観察されません。参照
枚数や再生産バンドルの評価。
すべてのプロンプトを送信するのではなく、重要な場所にプレミアム エージェント クォータを費やす
最も高価なバックエンドに。
決定的なプライバシー フロアを使用して、機密性の高いプロンプトをローカルに保ちます。
学習されたルーティングは上書きできません。
コンテキストを失わずにセッション中にバックエンドを切り替える - 共有セッション履歴、セマンティックメモリ、リダクションをユーザー間で共有

オラマ、コーデックス、クロード・コード。
ローカルの Ollama モデル、Codex CLI、および Claude Code 間でルーティングします。最初に決定論的なルールがあり、オプションで思い出すための小さな学習済み分類器が使用されます。
プライベート モード - 決定的なキーワード/PII/秘密形式のフロアにより、フォールバック時であっても、機密性の高いプロンプトがサブスクリプション バックエンドに到達することがブロックされます。
モデルに推測させるのではなく、決定論的なツール (時刻/日付、安全な計算機、単位変換、キーレスのライブストックとニュース) を使用して答えを根拠付けます。
バックエンド スイッチ間でコンテキストを伝達します。最近のユーザー、アシスタント、およびツールのターンは、編集された 1 つのセッション プロンプトにまとめられます。
ヘッドルームにインスピレーションを得たレイヤーで長いコンテキストを圧縮します。モデル境界パスは最近の会話を要約するだけですが、信頼できる事実、取得されたメモリ、および現在のリクエストはそのまま残ります。
ローカルの埋め込みベースのセマンティック メモリを介してバックエンド全体で記憶され、直接メモリ ルックアップに SQLite 検索が利用可能です。
すべての決定を説明し、メタデータのみのテレメトリを記録します (プロンプト/応答本体はありません)。
独自の評価、つまり 100 ケースの品質ベンチマーク、審査員としてのローカル LLM、および複数実行の統計ハーネスを提供します。
UI / CLI ──► セッションマネージャー (すべてのバックエンドで履歴を共有)
│
▼
能力検出器 (正規表現) ◄──► 決定論的ツール
│ (学習されたツールディスパッチャーがミスを回復し、ツールが検証します)
▼
プライバシーフロア (キーワード + PII + シークレット形式 - 一致は最終的です)
│ (学習された感度エスカレーターは保護を追加するだけです)
▼
決定論的な政策 ← 常に勝ちます。不明 ⇒ 地元
│ (学習済みルーター用品のリコール: ツール / ローカル / コーディング / 推論)
▼
コンテキストビルダー + 編集 ◄── セマンティックメモリ
│
▼
圧縮 (メタデータ + 履歴のみのコンテキスト パス)
│
▼
Ollama (デフォルト) │ Codex (コーディング) │ Claude Code (推論)
│
▼
R

esponse sanitizer ─► メタデータのみのテレメトリ
組織化された不変条件: 決定論的なポリシーが常に優先され、上書きされます。
学習したコンポーネント。プライバシー、ツールグラウンディング、強制選択、および
フォールバックは、ローカル モデルの実行時でも機能し続けるため、
学習済みコンポーネント — ダウンしています。
pip インストール スイッチボード ローカル
# ローカル モデル ランタイムをポイントします (Ollama をインストールしてから、小さなモデルをプルします)
オラマ プル ラマ3.2:3b
# セットアップの健全性をチェックする
配電盤の医者
# ask — 配電盤が配線し、接地し、その理由を教えてくれます
配電盤は「このエラー ログを要約して修正を提案してください」と尋ねます
# 何も実行せずにルーティングの決定を確認する
switchboard ルート「認証モジュールをリファクタリングしてテストを追加」
# ブラウザの方が好みですか?ローカル Web UI を起動し、http://127.0.0.1:8080/ui を開きます。
配電盤UI
Python 3.11 以降が必要です。 Codex / Claude Code バックエンドはオプションです - なし
すべてがローカルにルーティングされます。 docs/usage.md を参照してください。
スイッチボードには、ユーザー向けの CLI サーフェスが 2 つあります。
switchboard Route ... モデルを呼び出すことなく、同じコア バックエンドの決定をプレビューします。
Web UI、ベア switchboard ask ... 、および switchboard ask --backend auto ... はステートフル コア ワークフローを使用します。共有セッション、モデルの切り替え、セマンティック メモリの取得、コンテキスト境界の圧縮、およびバックエンド テレメトリはすべて同じパス上で実行されます。
switchboard ask --backend auto --new-session " 覚えておいてください: プライベートなメモにはローカル モデルを推奨します。"
switchboard ask --backend auto --session < session_id > --memory " 何を覚えておくべきですか? "
長いプロンプトと長いセッションでは、トークンの推定値と節約メタデータが記録されます。リクエストレベルのパスを使用すると、サイズが大きすぎる生のプロンプトを短縮できます。その後、コンテキスト境界パスは <recent_conversation> のみを圧縮します。 <trusted_facts> 、 <long_term_memory> 、および <current_user_requ

est> ブロックは 2 番目のパスから保護されるため、グラウンディングと意図がトークン バジェットと引き換えに使用されることはありません。
メモリはローカルです。 switchboard memory add は項目を SQLite に保存し、 semantic_memory_enabled がオンで Ollama が nomic-embed-text を提供できる場合、クロスバックエンド取得のために埋め込みにインデックスを作成します。スイッチボード メモリ検索は、埋め込みが利用できない場合でもローカル テキスト検索として機能します。
詳細: docs/context-memory-compression.md 。
5 つのタスク カテゴリ (コーディング、推論、
要約、プライベート、グラウンディング）、実際のバックエンドで実行され、ローカルによって判断されます
モデル、複数の独立した実行にわたって (平均値が示されています。完全な条件ごと)
数値、信頼区間、有意性検定については論文に記載されています)。
¹ 「すべてにプレミアム エージェントを使用する」ベースラインでは、すべての機能をブロックする必要があります。
漏れのない状態を維持するよう敏感に要求されるため、そのカバレッジが崩れ、まさにギャップが生じます
配電盤が閉まります。どのような条件や実行でもベンチマークのリークは観察されませんでした。
これらの数値は、Zenodo 上の論文の複製バンドルとともにフルハーネスが移動する実際のバックエンド ベンチマークから得られます。
コンテキスト: なぜこれが存在するのか (Uber、Microsoft、2026)
一部の雇用主はAIコーディングツールの支出を制限し始めている：ウーバーの報道
2026 年の AI 予算を使い果たした後、エンジニアに AI ツールあたり月額 1,500 ドルの上限を課した
4か月以内（ブルームバーグ）。
Microsoft のエクスペリエンス + デバイス組織が Claude Code から
GitHub Copilot CLI (Windows Central)。
支出上限は請求書を制御しますが、実際にどの作業を行うかは決定しません。
プレミアム モデルが必要か、またはどのプロンプトがマシンから出てはいけないか。より良い
パターンはルーティングであり、包括的配給ではありません: リクエストごとに何を決定するか
ローカルに属するもの、コーディング エージェントが必要なもの、プレミアム推論に値するもの。
Switchboard は、次のリファレンス実装です。

シングルのパターンで
ワークステーション。これはまだエンタープライズ製品ではありません。それは最も小さな正直です
再現可能なベンチマークを裏付ける、ローカルファーストルーティングが機能することの証明
それ。
配電盤はローカルファーストでプライバシーを意識した構造になっています。
決定論的プライバシー フロアは、非ローカル ルーティングの前に実行されます。肯定的な評決は最終的なものであり、学習した要素や即時の言葉遣いによって上書きすることはできません。
シークレット形式の検出 (クラウド キー、JWT、PEM ブロック、環境認証情報) は、そのパターンをコンテキスト リダクションと共有するため、ルーティング境界とリダクターが乖離することはありません。
メタデータのみのテレメトリ — プロンプトおよび応答の本文は、デフォルトでは保存されません。
セマンティックメモリの埋め込みと評価判定はローカルで実行されます。
Switchboard は意図的に API アクセスの再販、Web UI のスクレイピング、または
プロバイダーの制限をバイパスします - サブスクリプション CLI は、
認証されたユーザーは、読み取り専用サンドボックス モードでそれらを呼び出すことができます。参照
SECURITY.md および docs/privacy.md 。
決定的ルーター — キーワード ルール。不明なプロンプトのデフォルトはローカルファーストです。
学習済みルーター / ツール ディスパッチャー / 感度エスカレーター — ローカルで計算された埋め込み (~50 ミリ秒、純粋な Python 推論) 上の小さなソフトマックス分類器。それぞれは、黄金精度のゲートの背後で独自のサムダウン修正を行うことで数秒で再トレーニング可能です。それらは決定論的な道に閉ざされてしまうのです。
ツール - タイムゾーン付きの時刻/日付、安全な抽象構文ツリー計算機、単位変換、キーレスのライブ株価とニュース。
圧縮 — 構造を認識し、決定論的で、依存関係がありません。タスクヘッダー、コードブロック、トレースバック、根拠のある事実を保持します。
セマンティック メモリ — nomic-embed-text 埋め込み、コサイン取得、ローカル メモリ コマンド、および直接検索のための SQLite テキスト検索フォールバック。
評価 — モック evals (CI)、実際のバックエンドのスモーク スイ

100 件の品質ベンチマーク、敵対的なテスター/開発者のドッグフーディング ループ。
設定は config/personal.yaml にあります (安全なローカルファーストのデフォルトが付属しています —
config/personal.example.yaml を参照してください)。ハイライト:
設定:
router_mode : "学習済み" # ルール | llm |ハイブリッド |学んだ
private_mode : true # 非ローカル バックエンドからの機密プロンプトをブロックします
許可クラウド : false
圧縮有効 : true
圧縮しきい値トークン: 1000
semantic_memory_enabled : true
semantic_memory_top_k : 3
claude_code_web_search : true # ライブデータのフォールバック用にクロード コード WebSearch を許可します
Finance_provider : " yahoo "
news_provider : " google_news_rss "
プロバイダー API キーは、環境変数名 (例:
OPENAI_API_KEY )、決してインライン化しないでください。 docs/overrides.md を参照してください。
スイッチボードについては、プレプリント「Privacy-Aware Hybrid Routing Across」で説明されています。
異種の AI エージェント。」原稿、マルチラン
ベンチマーク ハーネス、統計集計および図のスクリプト、および
ケースごとの記録は、Zenodo 上の複製バンドルとしてまとめてアーカイブされます。
10.5281/zenodo.20836918 。
このリポジトリにはソフトウェアのみが同梱されます。意図的に載せていないのですが、
論文の実験実行ツールや図生成ツール - これは、
アーカイブ レコードなので、コードはルーター自体に集中したままになります。
make install # .venv + 編集可能なインストール

[切り捨てられた]

## Original Extract

Privacy-aware, local-first router for your CLI coding agents (Codex, Claude Code) and local LLMs (Ollama) — keeps sensitive prompts on-device and cuts premium-model usage. - aivinay/switchboard

GitHub - aivinay/switchboard: Privacy-aware, local-first router for your CLI coding agents (Codex, Claude Code) and local LLMs (Ollama) — keeps sensitive prompts on-device and cuts premium-model usage. · GitHub
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
aivinay
/
switchboard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
15 Commits 15 Commits .github .github assets assets config config docs docs scripts scripts switchboard switchboard tests tests .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md RELEASE.md RELEASE.md SECURITY.md SECURITY.md auto_route_demo.gif auto_route_demo.gif docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml View all files Repository files navigation
62% fewer premium-agent calls · 4.1/5 quality vs 4.6/5 always-premium · 0 benchmark leaks observed
Install ·
Evaluation ·
How it works ·
Privacy ·
Paper ·
Docs
One session, three backends: local by default, Codex for code, Claude Code for reasoning.
Switchboard wraps the CLI tools you already use — no separate service, no proxy, no resold API access — and routes each prompt with deterministic rules before any learned classifier runs.
In its 100-case benchmark, Switchboard kept 62% of requests off premium
agents while reaching 4.1/5 quality against a 4.6/5 always-premium
baseline , with 100% answered and no benchmark leaks observed . See
Evaluation for the numbers and reproduction bundle.
Spend premium agent quota where it matters instead of sending every prompt
to the most expensive backend.
Keep sensitive prompts local with a deterministic privacy floor that
learned routing cannot override.
Switch backends mid-session without losing context — shared session history, semantic memory, and redaction travel with you across Ollama, Codex, and Claude Code.
Routes across local Ollama models, the Codex CLI, and Claude Code — deterministic rules first, with optional tiny learned classifiers for recall.
Private mode — a deterministic keyword/PII/secret-format floor blocks sensitive prompts from ever reaching a subscription backend, even on fallback.
Grounds answers with deterministic tools (time/date, safe calculator, unit conversion, keyless live stock & news) instead of letting a model guess.
Carries context across backend switches: recent user, assistant, and tool turns are assembled into one redacted session prompt.
Compresses long context with a Headroom-inspired layer; the model-boundary pass only summarizes recent conversation, while trusted facts, retrieved memory, and the current request survive intact.
Remembers across backends via local embedding-based semantic memory, with SQLite search available for direct memory lookup.
Explains every decision and records metadata-only telemetry (no prompt/response bodies).
Ships its own evaluation — a 100-case quality benchmark, a local LLM-as-judge, and a multi-run statistical harness.
UI / CLI ──► Session manager (shared history across all backends)
│
▼
Capability detector (regex) ◄──► deterministic tools
│ (learned tool dispatcher recovers misses; tool verifies)
▼
Privacy floor (keywords + PII + secret formats — a match is FINAL)
│ (learned sensitivity escalator may only ADD protection)
▼
Deterministic policy ← always wins; unknown ⇒ local
│ (learned router supplies recall: tool / local / coding / reasoning)
▼
Context builder + redaction ◄── semantic memory
│
▼
Compression (metadata + history-only context pass)
│
▼
Ollama (default) │ Codex (coding) │ Claude Code (reasoning)
│
▼
Response sanitizer ─► metadata-only telemetry
The organizing invariant: deterministic policy always precedes and overrides
the learned components. Privacy, tool grounding, forced selection, and
fallback keep working even when the local model runtime — and therefore every
learned component — is down.
pip install switchboard-local
# point it at a local model runtime (install Ollama, then pull a small model)
ollama pull llama3.2:3b
# sanity-check your setup
switchboard doctor
# ask — Switchboard routes it, grounds it, and tells you why
switchboard ask " summarize this error log and suggest a fix "
# see the routing decision without running anything
switchboard route " refactor the auth module and add tests "
# prefer your browser? launch the local web UI, then open http://127.0.0.1:8080/ui
switchboard ui
Requires Python 3.11+ . Codex / Claude Code backends are optional — without
them, everything routes locally. See docs/usage.md .
Switchboard has two user-facing CLI surfaces:
switchboard route ... previews the same core backend decision without calling a model.
The web UI, bare switchboard ask ... , and switchboard ask --backend auto ... use the stateful core workflow: shared sessions, model switching, semantic-memory retrieval, context-boundary compression, and backend telemetry all run on the same path.
switchboard ask --backend auto --new-session " Remember: prefer local models for private notes. "
switchboard ask --backend auto --session < session_id > --memory " What should you remember? "
Long prompts and long sessions record token estimates and savings metadata. The request-level pass can shorten an oversized raw prompt; the context-boundary pass then compresses only <recent_conversation> . The <trusted_facts> , <long_term_memory> , and <current_user_request> blocks are protected from that second pass so grounding and intent are not traded away for token budget.
Memory is local. switchboard memory add stores the item in SQLite and, when semantic_memory_enabled is on and Ollama can serve nomic-embed-text , indexes an embedding for cross-backend retrieval. switchboard memory search works as local text search even when embeddings are unavailable.
Details: docs/context-memory-compression.md .
A 100-case benchmark across five task categories (coding, reasoning,
summarization, private, grounding), run on real backends and judged by a local
model, over multiple independent runs (means shown; full per-condition
numbers, confidence intervals, and significance tests are in the paper):
¹ The "just use the premium agent for everything" baseline must block every
sensitive prompt to stay leak-free, so its coverage collapses — exactly the gap
Switchboard closes. No benchmark leaks were observed in any condition or run.
These numbers come from a real-backend benchmark whose full harness travels with the paper's reproduction bundle on Zenodo .
Context: why this exists (Uber, Microsoft, 2026)
Some employers have begun rationing AI coding-tool spend: Uber reportedly
capped engineers at $1,500/month per AI tool after burning its 2026 AI budget
in four months ( Bloomberg );
Microsoft's Experiences + Devices org reportedly moved off Claude Code to
GitHub Copilot CLI ( Windows Central ).
A spend cap controls the invoice, but it does not decide which work actually
needs a premium model or which prompts should never leave the machine. A better
pattern is routing, not blanket rationing : decide request by request what
belongs local, what needs a coding agent, and what is worth premium reasoning.
Switchboard is a reference implementation of that pattern for a single
workstation. It is not yet an enterprise product; it is the smallest honest
proof that local-first routing can work, with a reproducible benchmark to back
it.
Switchboard is local-first and privacy-aware by construction:
The deterministic privacy floor runs before any non-local routing ; a positive verdict is final and cannot be overridden by a learned component or by prompt wording.
Secret-format detection (cloud keys, JWTs, PEM blocks, env credentials) shares its patterns with context redaction, so the routing boundary and the redactor can't drift apart.
Metadata-only telemetry — prompt and response bodies are not stored by default.
Semantic-memory embeddings and the eval judge run locally .
Switchboard deliberately does not resell API access, scrape web UIs, or
bypass provider limits — subscription CLIs are invoked exactly as the
authenticated user could invoke them, in read-only sandbox modes. See
SECURITY.md and docs/privacy.md .
Deterministic router — keyword rules; unknown prompts default local-first.
Learned router / tool dispatcher / sensitivity escalator — tiny softmax classifiers over a locally-computed embedding (~50 ms, pure-Python inference), each retrainable in seconds from your own thumbs-down corrections behind golden-accuracy gates. They fail closed to the deterministic path.
Tools — time/date with timezones, safe abstract-syntax-tree calculator, unit conversion, keyless live stock quotes & news.
Compression — structure-aware, deterministic, dependency-free; preserves task header, code blocks, tracebacks, and grounded facts.
Semantic memory — nomic-embed-text embeddings, cosine retrieval, local memory commands, and SQLite text-search fallback for direct search.
Evaluation — mock evals (CI), real-backend smoke suite, 100-case quality benchmark, adversarial tester/developer dogfooding loop.
Settings live in config/personal.yaml (ships with safe local-first defaults —
see config/personal.example.yaml ). Highlights:
preferences :
router_mode : " learned " # rules | llm | hybrid | learned
private_mode : true # block sensitive prompts from non-local backends
allow_cloud : false
compression_enabled : true
compression_threshold_tokens : 1000
semantic_memory_enabled : true
semantic_memory_top_k : 3
claude_code_web_search : true # allow Claude Code WebSearch for live-data fallback
finance_provider : " yahoo "
news_provider : " google_news_rss "
Provider API keys are referenced by environment-variable name (e.g.
OPENAI_API_KEY ), never inline. See docs/overrides.md .
Switchboard is described in a preprint — "Privacy-Aware Hybrid Routing Across
Heterogeneous AI Agents." The manuscript, the multi-run
benchmark harness, the statistical-aggregation and figure scripts, and the
per-case records are archived together as a reproduction bundle on Zenodo:
10.5281/zenodo.20836918 .
This repository ships only the software. It deliberately does not carry the
paper's experiment-running or figure-generation tooling — that lives with the
archival record so the code stays focused on the router itself.
make install # .venv + editable instal

[truncated]
