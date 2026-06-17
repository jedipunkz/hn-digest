---
source: "https://github.com/ahmeddyounis/noema64"
hn_url: "https://news.ycombinator.com/item?id=48568254"
title: "Show HN: Noema64 – an open-source LLM chess engine (still in beta though)"
article_title: "GitHub - ahmeddyounis/noema64: Explainable LLM-assisted chess engine with legal move enforcement, strategy memory, UCI support, and a desktop GUI. · GitHub"
author: "ahmed_duski"
captured_at: "2026-06-17T10:37:16Z"
capture_tool: "hn-digest"
hn_id: 48568254
score: 1
comments: 1
posted_at: "2026-06-17T10:23:54Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Noema64 – an open-source LLM chess engine (still in beta though)

- HN: [48568254](https://news.ycombinator.com/item?id=48568254)
- Source: [github.com](https://github.com/ahmeddyounis/noema64)
- Score: 1
- Comments: 1
- Posted: 2026-06-17T10:23:54Z

## Translation

タイトル: Show HN: Noema64 – オープンソース LLM チェス エンジン (まだベータ版ですが)
記事のタイトル: GitHub - ahmeddyounis/noema64: 正当な手の強制、戦略メモリ、UCI サポート、デスクトップ GUI を備えた説明可能な LLM 支援チェス エンジン。 · GitHub
説明: 正当な手の強制、戦略メモリ、UCI サポート、およびデスクトップ GUI を備えた、説明可能な LLM 支援チェス エンジン。 - アーメディユーニス/noema64

記事本文:
GitHub - ahmeddyounis/noema64: 正当な手の強制、戦略メモリ、UCI サポート、デスクトップ GUI を備えた説明可能な LLM 支援チェス エンジン。 · GitHub
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
アーメディユニス
/
ノエマ64
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
257 コミット 257 コミット .github .github cmd cmd configs configs docs docs external 内部プロンプト/ v1 プロンプト/ v1 スキーマ スキーマ .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md go.mod go.mod go.sum go.sum すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Noema64 は、言語モデルを永続的な戦略プランナーとして使用するオープンソースの説明可能なチェス エンジンであり、決定論的な Go コードは正当な手の検証、ゲーム状態、フォールバック、UCI プロトコルの動作、トレース、およびローカル永続性を所有します。
Stockfish の代替品ではありません。このプロジェクトは、合法的なフルゲーム プレイ、検査可能な戦略メモリ、プロバイダーの障害回復、スタディ ワークフロー、プロトコル セーフな UCI の使用向けに最適化されています。
プレイワークスペース
プロバイダーの設定
能力
github.com/corentings/chess/v2 を通じて、正当な手、FEN、PGN、手筋履歴、結果、プロモーション、キャスリング、およびアンパッサンを備えた Go チェス コア ラッパーを入手します。
Strategy Memory v1.2 構造体、差分分析、バージョン管理された編集可能なプロンプト パック、厳密な JSON 解析、候補修復、およびスキーマ検証。
オフラインで動作するモック プロバイダー、OpenAI、OpenAI 互換 HTTP、Anthropic、Gemini、Ollama、およびローカル ポリシー優先プロバイダー アダプター。
合法的な手が存在する場合、常に合法的な手を選択する決定論的フォールバック ラダー。
静的 Blunderguard ベリファイアに加えて、オプションの外部 UCI ベリファイアと外部テーブルベース プローブ パスのサポート。
uci 、 isready 、 ucinewgame 、position 、 go 、 stop 、 quit 、および setoption を含む UCI バイナリ。
Wails v2 GUI エントリポイント (組み込みボード、時間コントロール、最近のゲーム、設定、戦略、候補者、トレース、辞任)

PGN/FEN/JSONL トレース エクスポート、ベンチマーク コントロール。
研究ダッシュボード、圧縮メモリ、計画の一貫性、候補者の多様性、決定論的なマルチエージェントレビュー、および編集可能な戦略メモリ API。
Chess960/カスタム開始メタデータ、決定論的な Chess960 開始生成、および互換性レイヤーを介した Noema64 管理の Chess960 キャスリング。
ローカル バックアップ/復元アーカイブ、微調整された JSONL データセット エクスポート、決定論的なトーナメント/レーティングの自動化、構成された外部バイナリのサンドボックス検証。
JSON または CSV ベンチマーク出力を使用した CLI およびベンチマーク コマンド。
ローカル YAML 設定、JSONL デシジョン トレース、戦略メモリを含む編集されたゲーム スナップショット、テスト、プロンプト、構成、およびドキュメント。
フロントエンド構文と GUI スモーク チェック用の Node.js。
パッケージ化されたデスクトップ ビルド用の Wails v2 CLI。
オプション: OpenAI API キー、ローカルの OpenAI 互換 LLM エンドポイント、またはその他のクラウド プロバイダー キー。
オプション: Stockfish や外部テーブルベース プローブなどのユーザー指定の UCI 検証ツール。ベリファイアやテーブルベースのバイナリはバンドルされていません。
GitHub Releases からパッケージ化されたデスクトップ ビルドをダウンロードします。リリース アセットは現在署名されていない Windows および macOS ビルドであるため、コード署名と公証が追加されるまで、OS にセキュリティ警告が表示される可能性があります。
テストに行ってください。/...
npm --prefix cmd/noema64-gui/フロントエンド テスト
go run ./cmd/noema64 -cmd state
./cmd/noema64 -cmd エンジンを実行します。
go run ./cmd/noema64 -cmd Analyze -fen ' r1bqkbnr/pppp1ppp/2n5/4p3/4P3/5N2/PPPP1PPP/RNBQKB1R w KQkq - 2 3 '
go run ./cmd/noema64 -cmd state -variant chess960 -seed 42
go run ./cmd/noema64 -cmd Study
go run ./cmd/noema64 -cmd トーナメント -games-per-pair 1
実行します。/cmd/noema64-bench -games 100
./cmd/noema64-bench -games 100 -format csv > benchmark.csv を実行します。
UCI の煙:
printf ' uci\nisready\nucinewgame\nposition startpos move e2e4 e7e5 g1f3\ngo movetime 1

00\n終了\n ' | ./cmd/noema64-uci を実行します。
GUI開発:
cd cmd/noema64-gui
開発者が泣き叫ぶ
埋め込みフロントエンドは静的で、Wails バインディングで動作します。モックプロバイダーがデフォルトであるため、API キーは必要ありません。 GUI は、再起動時に設定されたログ ディレクトリから、クロックと戦略の状態を含む最新の保存されたゲームを復元します。
CI は、Go フォーマット/vet/テスト、非 Wails パッケージのレース テスト、フロントエンド スモーク テスト、UCI スモーク、トレース検証、パフォーマンス、依存関係ライセンス スキャン、非 GUI バイナリ ビルド、および 100 ゲームの信頼性ベンチマークを実行します。リリース ワークフローは、署名されていない Windows および macOS デスクトップ アーティファクトをネイティブ GitHub Actions ランナー上に構築し、それらをドラフト GitHub リリースに公開します。
リリース ワークフローの詳細は docs/releasing.md に文書化されています。
pure : プロバイダ候補は修復され、合法性がフィルタリングされます。戦術検証ツールは使用されません。
current : 「今がベスト」モード;決定ごとに戦略記憶をリセットし、長期計画の継続性を考慮して現在の戦術/探索をランク付けします。
blunderguard : デフォルトモード;静的検証者は明らかな合致リスクを拒否し、支援を開示することができます。
ハイブリッド : スコアリングでは、検証者/検索スタイルのシグナルにより多くの重みが割り当てられます。
Coach : 教育指向の性格設定で同じ法的パイプラインを使用します。
パーソナリティ プロファイルは表示専用ではありません。彼らのリスク許容度は、各候補者の小さなpersonality_scoreとしてアービターに含まれるため、アグレッシブなプロファイルは安全な強制的な動きに向けて密接な関係を壊すことができますが、ポジショナルプロファイルとコーチプロファイルはより明確な低リスクの選択に傾く可能性があります。
現在のモードとハイブリッド モードでは、決定トレースに deterministic_mcts_material 検索支援が記録されます。これは、制限された決定論的なプレイアウト スコアラーであり、エンジン強度の MCTS を主張するものではありません。
GUI ツールバーには、スタディ ダイアログとラボ ダイアログが含まれています。同じワークフロー

CLI から利用できます。
go run ./cmd/noema64 -cmd Study
./cmd/noema64 -cmd エージェントを実行します。
go run ./cmd/noema64 -cmd book
go run ./cmd/noema64 -cmd Compare
go run ./cmd/noema64 -cmd Backup -backup-dir 実行/バックアップ
実行します。/cmd/noema64 -cmdfinetune
go run ./cmd/noema64 -cmd トーナメント -games-per-pair 2
研究では、圧縮された記憶、冒頭の本の提案、終盤のトレーナードリル、計画の一貫性、候補者の多様性、レッスン/パズルのプロンプト、ヒートマップデータ、および決定論的な戦略家/批評家/戦術家/裁定者のレビューが返されます。ラボ ワークフローでは、ローカル zip バックアップの作成、ターゲット ディレクトリへのバックアップの復元、純粋分析とハイブリッド分析の比較、プロンプト パックの比較、カスタム パーソナリティ プロファイルのドラフトの作成、サニタイズされたトレースからのローカル微調整 JSONL サンプルのエクスポート、およびコア モード全体でのエンジン対エンジンのトーナメント評価の実行が行われます。
デフォルトの設定ではオフライン モック プロバイダーが使用されるため、Noema64 は API キーなしで実行できます。
GUI から [設定] を開き、プロバイダー プロファイルまたはプロバイダーを選択し、モデルを入力して保存します。 OpenAI の場合、エンドポイント フィールドは自動的に管理され、 https://api.openai.com/v1 を使用します。必要なのはモデルと API キーのみです。 OpenAI 互換 の場合は、ローカルまたはリモート互換サーバーのエンドポイントを入力します。
エンドツーエンドのクラウド設定ガイド:
YAML の OpenAI の場合、プロバイダーとモデルを設定します。
llm :
プロバイダー：オープンアイ
モデル: your-openai-model
API_キー: " "
YAML で別の OpenAI 互換エンドポイントを使用するには、次のように設定します。
llm :
プロバイダー : openai_compatibility
エンドポイント: http://localhost:11434/v1
モデル: あなたのモデル
API_キー: " "
プロバイダー設定は api_key または api_key_ref をサポートします。 GUI では、キーチェーン アクションは、サポートされている場合、入力されたキーを OS キーチェーンに保存し、生のキーを参照に置き換えます。
エンドポイントベースのプロバイダー モードでは、FEN、合法的な動き、動き履歴、戦略メモリ、

OpenAI、Anthropic、Gemini、ローカルまたはリモートの OpenAI 互換エンドポイント、Ollama エンドポイントなど、構成されたプロバイダーの設定。 GUI では、これらのプロバイダーを保存する前に、明示的なデータ共有の承認が必要です。生のプロンプトログはデフォルトではオフになっています。
静的検証が組み込まれています。外部 UCI 検証はオプションです。
検証者:
有効 : true
種類：uci
パス : /usr/local/bin/stockfish
移動時間_ms : 100
最大センチポーン損失 : 180
Noema64 には Stockfish がバンドルされていません。構成されている場合、外部 UCI 検証者は、 searchmoves を使用して LLM 候補の移動を分析し、センチポーンの損失を最良の候補と比較し、検証者の決定をトレースに記録します。
外部ベリファイアとテーブルベースのパスはシェルなしで実行され、起動前に検証されます。 Stockfish などの単純な PATH バイナリ名が許可されます。空白、制御文字、パス トラバーサル、または実行不可能な絶対ターゲットを含むパスは拒否されます。
外部テーブルベースのプローブもオプションです。 { "fen": "...", "candidates": ["e2e4"] } JSON を標準入力から読み取り、正確なテーブルベース JSON を標準出力に返すプローブ実行可能ファイルを構成します。
検証者:
tablebase_enabled : true
tablebase_path : /usr/local/bin/noema64-tablebase
tablebase_timeout_ms : 1000
UCIの例
市
準備ができています
うしニューゲーム
位置 startpos は e2e4 e7e5 g1f3 を移動します
go wtime 300000 btime 300000 winc 2000 binc 2000
やめる
UCI プロセスは、プロトコル出力を stdout にのみ書き込みます。 JSON トレースは、有効にするとローカル ファイルに書き込まれます。 GUI ゲームのスナップショットは、logging.output_dir/games に保存されます。
ボット ブリッジの使用法については、 docs/lichess-bot.md を参照してください。
静的ベリファイアは意図的に浅いものになっています。
Chess960 のキャスリングは、上流の移動ジェネレーターではなく、Noema64 の互換性レイヤーによって処理されます。
LLM の品質は、構成されたプロバイダーによって異なります。
生のプロム

ユーザーがログを有効にしない限り、ポイントは保存されません。
Noema64 ソースは MIT ライセンスを取得しています。オプションの外部検証バイナリは、独自のライセンスを保持します。
正当な手の強制、戦略メモリ、UCI サポート、およびデスクトップ GUI を備えた説明可能な LLM 支援チェス エンジン。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Explainable LLM-assisted chess engine with legal move enforcement, strategy memory, UCI support, and a desktop GUI. - ahmeddyounis/noema64

GitHub - ahmeddyounis/noema64: Explainable LLM-assisted chess engine with legal move enforcement, strategy memory, UCI support, and a desktop GUI. · GitHub
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
ahmeddyounis
/
noema64
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
257 Commits 257 Commits .github .github cmd cmd configs configs docs docs internal internal prompts/ v1 prompts/ v1 schemas schemas .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md go.mod go.mod go.sum go.sum View all files Repository files navigation
Noema64 is an open-source explainable chess engine that uses a language model as a persistent strategic planner while deterministic Go code owns legal move validation, game state, fallback, UCI protocol behavior, traces, and local persistence.
It is not a Stockfish replacement. The project is optimized for legal full-game play, inspectable strategy memory, provider failure recovery, study workflows, and protocol-safe UCI use.
Play workspace
Provider settings
Capabilities
Go chess core wrapper with legal moves, FEN, PGN, move history, outcomes, promotions, castling, and en passant through github.com/corentings/chess/v2 .
Strategy memory v1.2 structs, diffing, versioned editable prompt packs, strict JSON parsing, candidate repair, and schema validation.
Mock provider that works offline, OpenAI, OpenAI-compatible HTTP, Anthropic, Gemini, Ollama, and local policy-prior provider adapters.
Deterministic fallback ladder that always chooses a legal move when legal moves exist.
Static blunderguard verifier plus optional external UCI verifier and external tablebase probe path support.
UCI binary with uci , isready , ucinewgame , position , go , stop , quit , and setoption .
Wails v2 GUI entrypoint with embedded board, time controls, recent games, settings, strategy, candidates, trace, resignation, PGN/FEN/JSONL trace export, and benchmark controls.
Study dashboard, compressed memory, plan coherence, candidate diversity, deterministic multi-agent review, and editable strategy memory APIs.
Chess960/custom-start metadata, deterministic Chess960 start generation, and Noema64-managed Chess960 castling through a compatibility layer.
Local backup/restore archives, fine-tune JSONL dataset export, deterministic tournament/rating automation, and sandbox validation for configured external binaries.
CLI and benchmark commands with JSON or CSV benchmark output.
Local YAML settings, JSONL decision traces, redacted game snapshots with strategy memory, tests, prompts, configs, and docs.
Node.js for frontend syntax and GUI smoke checks.
Wails v2 CLI for packaged desktop builds.
Optional: an OpenAI API key, local OpenAI-compatible LLM endpoint, or other cloud provider key.
Optional: a user-supplied UCI verifier such as Stockfish or an external tablebase probe. No verifier or tablebase binary is bundled.
Download packaged desktop builds from GitHub Releases . Release assets are currently unsigned Windows and macOS builds, so your OS may show a security warning until code signing and notarization are added.
go test ./...
npm --prefix cmd/noema64-gui/frontend test
go run ./cmd/noema64 -cmd state
go run ./cmd/noema64 -cmd engine
go run ./cmd/noema64 -cmd analyze -fen ' r1bqkbnr/pppp1ppp/2n5/4p3/4P3/5N2/PPPP1PPP/RNBQKB1R w KQkq - 2 3 '
go run ./cmd/noema64 -cmd state -variant chess960 -seed 42
go run ./cmd/noema64 -cmd study
go run ./cmd/noema64 -cmd tournament -games-per-pair 1
go run ./cmd/noema64-bench -games 100
go run ./cmd/noema64-bench -games 100 -format csv > benchmark.csv
UCI smoke:
printf ' uci\nisready\nucinewgame\nposition startpos moves e2e4 e7e5 g1f3\ngo movetime 100\nquit\n ' | go run ./cmd/noema64-uci
GUI development:
cd cmd/noema64-gui
wails dev
The embedded frontend is static and works with Wails bindings. The mock provider is the default, so no API key is required. The GUI restores the most recent saved game, including clock and strategy state, from the configured log directory on relaunch.
CI runs Go format/vet/tests, race tests for the non-Wails packages, frontend smoke tests, UCI smoke, trace validation, perft, dependency license scanning, non-GUI binary builds, and the 100-game reliability benchmark. The release workflow builds unsigned Windows and macOS desktop artifacts on native GitHub Actions runners and publishes them to draft GitHub Releases.
Release workflow details are documented in docs/releasing.md .
pure : provider candidates are repaired and legality-filtered; no tactical verifier is used.
current : "Best now" mode; resets strategy memory for each decision and ranks current-position tactics/search over long-term plan continuity.
blunderguard : default mode; static verifier can reject obvious mate-in-one risks and discloses assistance.
hybrid : scoring reserves more weight for verifier/search-style signals.
coach : uses the same legal pipeline with teaching-oriented personality settings.
Personality profiles are not display-only. Their risk tolerance is included in the arbiter as a small personality_score on each candidate, so aggressive profiles can break close ties toward safe forcing moves while positional and coach profiles lean toward clearer low-risk choices.
Current and hybrid modes record deterministic_mcts_material search assistance in the decision trace. It is a bounded deterministic playout scorer, not a claim of engine-strength MCTS.
The GUI toolbar includes Study and Lab dialogs. The same workflows are available from the CLI:
go run ./cmd/noema64 -cmd study
go run ./cmd/noema64 -cmd agents
go run ./cmd/noema64 -cmd book
go run ./cmd/noema64 -cmd compare
go run ./cmd/noema64 -cmd backup -backup-dir runs/backups
go run ./cmd/noema64 -cmd finetune
go run ./cmd/noema64 -cmd tournament -games-per-pair 2
Study returns compressed memory, opening-book suggestions, endgame trainer drills, plan coherence, candidate diversity, lesson/puzzle prompts, heatmap data, and deterministic strategist/critic/tactician/arbiter reviews. Lab workflows create local zip backups, restore backups into a target directory, compare pure vs hybrid analysis, compare prompt packs, build custom personality profile drafts, export local fine-tune JSONL examples from sanitized traces, and run engine-vs-engine tournament ratings across core modes.
Default config uses the offline mock provider, so Noema64 can run without an API key.
From the GUI, open Settings, choose a provider profile or provider, enter a model, then save. For OpenAI , the endpoint field is managed automatically and uses https://api.openai.com/v1 ; you only need the model and API key. For OpenAI-compatible , enter the endpoint for your local or remote compatible server.
End-to-end cloud setup guides:
For OpenAI in YAML, set the provider and model:
llm :
provider : openai
model : your-openai-model
api_key : " "
To use another OpenAI-compatible endpoint in YAML, set:
llm :
provider : openai_compatible
endpoint : http://localhost:11434/v1
model : your-model
api_key : " "
Provider settings support api_key or api_key_ref . In the GUI, the keychain action stores a typed key in the OS keychain when supported and replaces the raw key with a reference.
Endpoint-backed provider modes may send FEN, legal moves, move history, strategy memory, and settings to the configured provider, including OpenAI, Anthropic, Gemini, local or remote OpenAI-compatible endpoints, and Ollama endpoints. The GUI requires an explicit data-sharing acknowledgement before saving those providers. Raw prompt logging is off by default.
Static verification is built in. External UCI verification is optional:
verifier :
enabled : true
kind : uci
path : /usr/local/bin/stockfish
movetime_ms : 100
max_centipawn_loss : 180
Noema64 does not bundle Stockfish. When configured, the external UCI verifier analyzes LLM candidate moves with searchmoves , compares centipawn loss against the best candidate, and records verifier decisions in the trace.
External verifier and tablebase paths are executed without a shell and are validated before launch. Simple PATH binary names such as stockfish are allowed; paths containing whitespace, control characters, path traversal, or non-executable absolute targets are rejected.
External tablebase probing is also optional. Configure a probe executable that reads { "fen": "...", "candidates": ["e2e4"] } JSON from stdin and returns exact tablebase JSON on stdout:
verifier :
tablebase_enabled : true
tablebase_path : /usr/local/bin/noema64-tablebase
tablebase_timeout_ms : 1000
UCI Example
uci
isready
ucinewgame
position startpos moves e2e4 e7e5 g1f3
go wtime 300000 btime 300000 winc 2000 binc 2000
quit
The UCI process writes protocol output only to stdout. JSON traces are written to local files when enabled. GUI game snapshots are stored under logging.output_dir/games .
For bot bridge usage, see docs/lichess-bot.md .
The static verifier is intentionally shallow.
Chess960 castling is handled by Noema64's compatibility layer rather than the upstream move generator.
LLM quality depends on the configured provider.
Raw prompts are not stored unless the user enables logging.
Noema64 source is MIT licensed. Optional external verifier binaries retain their own licenses.
Explainable LLM-assisted chess engine with legal move enforcement, strategy memory, UCI support, and a desktop GUI.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
