---
source: "https://github.com/KanishkNoir/cognikernel"
hn_url: "https://news.ycombinator.com/item?id=48984355"
title: "Show HN: Cognikernel- Local Memory for AI Coding Assistants"
article_title: "GitHub - KanishkNoir/cognikernel: Persistent, structured project memory for Claude Code and Codex - event-sourced, typed, fail-open. Your agent stops re-deciding what you already decided. · GitHub"
author: "KanishkNoir"
captured_at: "2026-07-20T21:01:53Z"
capture_tool: "hn-digest"
hn_id: 48984355
score: 1
comments: 0
posted_at: "2026-07-20T20:20:10Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Cognikernel- Local Memory for AI Coding Assistants

- HN: [48984355](https://news.ycombinator.com/item?id=48984355)
- Source: [github.com](https://github.com/KanishkNoir/cognikernel)
- Score: 1
- Comments: 0
- Posted: 2026-07-20T20:20:10Z

## Translation

タイトル: Show HN: Cognikernel - AI コーディング アシスタントのためのローカル メモリ
記事のタイトル: GitHub - KanishkNoir/cognikernel: クロード コードおよびコーデックス用の永続的で構造化されたプロジェクト メモリ - イベント ソース、型指定、フェールオープン。エージェントは、あなたがすでに決定したことを再度決定するのをやめます。 · GitHub
説明: クロード コードおよびコーデックス用の永続的で構造化されたプロジェクト メモリ - イベント ソース、型指定、フェールオープン。エージェントは、あなたがすでに決定したことを再度決定するのをやめます。 - KanishkNoir/コグニカーネル
HN テキスト: クロード コード/コーデックスでのすべてのコーディング セッションは、盲目的に開始されます。昨日、1 週間前、または 1 か月前に実行したセッションで行ったアーキテクチャ上の決定は忘れられます。最終的には、これらのエージェントがメモリ機能に依存する Claude.md ファイルまたは diff マークダウン ファイルの維持に依存することになります。この問題を解決するために、私は CogniKernel の構築を開始しました。 CogniKernel は、マークダウン メモリ ファイルに依存するのではなく、コーディング セッション自体を監視します。意思決定、制約、アーキテクチャ上の決定、プロジェクト固有のコンテキストなどの情報が自然に現れるときにキャプチャされます。目標はシンプルです。エージェントは、私が何を覚えるべきかを指示しなくても、昨日学んだことを覚えている必要があります。自動的に保存されるもの: - アーキテクチャ上の決定と設計の理論的根拠
- プロジェクトの規約とコーディング パターン
- 重要な実装の詳細
- ユーザー設定と繰り返しの制約
- ファイル、モジュール、およびコンポーネント間の関係 メモリはローカルファーストで、既存のコーディング エージェントを置き換えるのではなく、既存のコーディング エージェントと並行して動作するように設計されています。私の設計目標は次のとおりです: - 手動によるメモリキュレーションをゼロにする
- クラウド依存ではなくローカル ストレージ
- すべてをダンプするのではなく、関連するコンテキストを含む高速な取得
- エージェントに依存しないアーキテクチャなので、単一のコーディング アシスタントを超えて機能します
- 時間の経過とともに改善される記憶力

検索不可能なログに成長するのではなく、影響はありましたか？
クロード コードから自動メモリをオフに切り替えながら、cognikernel を使用して 4 つの本格的なプロジェクトをビルドし、同じプロジェクトで同じプロンプト、自動メモリを使用した同じセッションを使用しました - 使用されるトークンが 30% 削減されました
- 使用する読み取りツールが 4 分の 1 に減りました。クロード コードとコーデックスの両方を使用している場合は、代替ツールでセッションを開始するだけで、記憶がそこでも引き継がれます。これは趣味のプロジェクトでしたが、これを使用して仕事に役立った場合は、ぜひフィードバックをお知らせください。

記事本文:
GitHub - KanishkNoir/cognikernel: クロード コードおよびコーデックス用の永続的で構造化されたプロジェクト メモリ - イベント ソース、型指定、フェールオープン。エージェントは、あなたがすでに決定したことを再度決定するのをやめます。 · GitHub
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
カニス

香港ノワール
/
認知カーネル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
148 コミット 148 コミット .github .github スクリプト スクリプト src/ memlora src/ memlora テスト テスト .gitignore .gitignore COTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード コードとコーデックスのための永続的で構造化されたプロジェクト メモリ。コグニカーネル
フック サーフェスを通じてコーディング セッションを監視し、決定事項を抽出し、
制約や放棄されたアプローチは維持する価値があり、それらを統合して
イベントソースのストアを作成し、それらをコンパクトなコンテキスト ブロックとして次のインスタンスに注入します。
あなたが仕事をする時間になると、エージェントはあなたがすでに決定したことを再決定するのをやめます。店
はプロジェクト パス上でキー設定されているため、一方のツールで作成されたメモリはもう一方のツールに移動します。
これはベクトルデータベースラッパーではありません。これはイベントソースの型付きログです。
レキシカルプライマリ検索、書き込み時の統合、およびフェールオープンを備えたメモリ
セッションを決して中断しないように設計された信頼性スパイン。
また、ループ内に LLM はありません。ほとんどの記憶ツールは、
「重要なことを要約する」ための生成モデルへのトランスクリプト — 別の API
キー、セッションごとのトークンコスト、追加されたレイテンシ、およびセッションコンテンツの終了
機械。 CogniKernel は、抽出を生成ではなく分類として扱います。
2 つの小さなパイプラインによる決定的なサニタイズ → 分類 → 統合パイプライン
微調整されたエンコーダー モデル (約 130 MB ONNX、ミリ秒単位で CPU 上でローカルに実行)
顕著性をスコアリングし、新しい決定が古い決定に優先する時期を検出します。いいえ
API 呼び出し、トークンの請求はなく、マシンから何も出ません。唯一の

LLM が関係する
は、すでに実行しているコーディング エージェントです。CogniKernel がそれを記憶させます。
命名: CogniKernel はプロジェクトです。 memlora (memlora-edge パッケージ) は
Python モジュールとそれが出荷される CLI — コードが作成された作業名
下。 1 つのプロジェクト、2 つの名前: memlora init 、 memlora Doctor など。
CogniKernel の動作はすべて 1 つのループです: 観察 → 抽出 → 統合 →
保管→取り出し→組み立て→注入。左側のレールがキャプチャします。右のレール
思い出します。その下の背骨は両方を正直に保ちます。
+========================= クロード・コード・セッション =========================+
|作業メモリ - エージェントが判断するコンテキスト ウィンドウ |
+--------------+-------------------------------------------+--------------+
|停止フックはトランスクリプトをキャプチャします |ブロックを注入する
v ^
+--------------+--------------+ +--------------+--------------+
| [1] 抽出パイプライン | | [6] 圧縮 + 射出 |
|サニタイズ -> 分類 -> | |権限重視の予算 |
|顕著性 (ONNX) -> | |ドロップトゥフィット (すべての | を維持)
|決定キー + 契約 | |制約) -> ブロック |
+--------------+--------------+ +--------------+--------------+
|エンキュー |ランク+フィット感
v ^
+--------------+--------------+ +--------------+--------------+
| [2] 労働者 + 統合 | | [5] 検索 |
|クレーム -> デルタマージ -> | | FTS5 BM25 + 密 -> RRF |
|置き換え (最新の勝利) -> | |禁止検索 (K1) |
|プロジェクト (冪等) | |スケルトン グラフ (PageRank) |
+--------------+--------------+ +--------------+--------------+
|永続化 (アトミック) |思い出す
v ^
+--------------+-------------------------------------------+--------------+
| [3] イベントソースストア * SQLite (WAL) |
|型付きイベント |証拠 |来歴 | FTS5 |埋め込み |元帳 |
+------------------------------------------------------------------------+

| [4] 信頼性 SPINE のアトミック マイグレーション |冪等なリプレイ | |
|医師 -- 厳密 |フェールオープン フック |インポートリンター | CIゲート |
+------------------------------------------------------------------------+
4つのフック面
CogniKernel は 4 つのポイントでセッションに接続します。メモリの場合、それぞれがフェールオープンです
が利用できないかエラーが発生した場合、フックは WARNING でログに記録され、正常に戻り、
セッションは継続します。
メモリはフリーテキストの塊ではなく、型付けされています。タイプはランキング、レンダリング、
そしてスーパーセッション:
DECISION — 行われた選択 (「レート リミッターに Redis を使用する」)
CONSTRAINT_HARD / CONSTRAINT_SOFT — 強制力によって等級付けされたルール
APPROACH_ABANDONED_DO_NOT_RETRY — 行き止まりで、墓地に保管されているため、
エージェントは再試行しません
規則、構成事実、スキーマの決定 (正規の役割キー)
決定キーにより、後の再記述が以前の再記述に優先します
(最新順) なので、ストアは蓄積されるのではなく自己統合されます。
矛盾。オプションのクロスエンコーダは、セマンティックスーパーセッションを追加します。
エンコーダバックエンドがインストールされています。
字句一次、融合信号として密を含む — 決して純粋なベクトルではない:
ハイブリッド コア — FTS5 BM25 ∪ オプションの高密度埋め込み → 相互ランク融合
Prevention_search — タイプ制限された語彙プールであるため、「X を実行しない」ルール
X を実行しようとしている瞬間に、トピック的に似た散文で混雑することはできません
Skeleton — PageRank によってランク付けされた AST シンボル グラフ。 find_関連ユニオン
インポートグラフ隣接イベントを持つセマンティック (埋め込み) 近傍を、
変更が何に影響するかを明らかにする (セマンティック軸には追加の埋め込みが必要)
読み取り時のゴールデンレコードの統合 — 最新の勝利の調整なので思い出してください
修正の山ではなく、1 つの一貫した答えを返します
3 部門の比較でベンチマーク — CogniKernel vs フラット キュレーション ノート vs なし
メモ

ry — 4 つのマルチセッション プロジェクトにわたる実際のエージェント セッションを使用します。
ファイルには「普遍的な勝利」と書かれています。 CogniKernel アームが作成したファイルが最も少ない
すべてのプロジェクトでの読み取り - 通常は 2 ～ 4 倍少なくなります (23 対 63、16 対 47/53、
40 対 89/83)、最良のケースでは、注入されたため 3 読み取り対 29
ブロック + AST スケルトンはリポジトリ全体の形状を保持しました。読み取りが少ないということは、読み取りも少ないことを意味します
ツールの往復と、実際の作業のために残されたコンテキスト ウィンドウの多く -
セッションは安くなるだけでなく、圧縮するまでに長くなります。
トークン: メモリが重要な場合は最大 20% 安くなります。価格加重トークン
コスト (キャッシュ書き込み 1.25 倍、キャッシュ読み取り 0.1 倍、出力 5 倍) は 18 ～ 23% 低かった
進化する意思決定とファイル間の依存関係を伴うプロジェクトについて — そして大まかに
コード自体が安価な、実装が多量にある小規模なプロジェクトの洗浄
再読すること。私たちは、生のトークンの数字ではなく、正直な数字を公開します（生の合計
見た目は約 30 ～ 40% 良くなりますが、セッションの請求額の約 95% はキャッシュ読み取りの割引になります)。
再導出ではなく思い出す。記憶が活きるのはプロジェクトです
その状態が大きすぎる、進化しすぎる、または再導出するには長すぎる
安価: エージェントはすでに決定、制約、および条件を理解し始めます。
セッションの最初の四半期を再発見に費やす代わりに、行き止まりを克服する
彼ら。
システムは、決して黙って劣化するのではなく、はっきりとわかるように劣化するように設計されています。
アトミック マイグレーション - 番号付きの各マイグレーションは、本体 + バージョン バンプを適用します。
1 回のトランザクションで。スクリプトの途中でクラッシュしても安全です
冪等な再実行 - 再実行ワーカー ジョブは、二重カウントまたはドリフト減衰を行うことができません (証拠と出所の保護)
フェールオープン フック - すべてのサーフェスが自身の障害を飲み込み、 WARNING でログを記録します。沈黙は決して成功とは言えない
memlora ドクター --strict — サブシステムごとの健全性 (スキーマ、FTS5、埋め込み、シンボル、ワーカー キュー)。劣化時にゼロ以外で終了する
建築

施行 — メタテストによって保護されている import-linter の階層型コントラクト。そのためタイプミスによって黙って無効にされることはありません。
CI プロモーション ゲート — すべての PR の lint + フル スイート (テスト/信頼性/障害挿入を含む)。 CONTRIBUTING.md を参照してください。
ストアはプラットフォームに依存しません。論理プロジェクトごとに 1 つの SQLite DB であるため、Claude Code
同じディレクトリ内で動作する Codex は 1 つのメモリを共有します。プロジェクトの解像度は、
alias-aware : C:\repo と /mnt/c/repo は同じストアに解決されるため、メモリ
Windows、WSL、およびネイティブ マウントにわたるチェックアウトに従います。本当に
異なるチェックアウト パス、オプトイン project_identity キー
.memlora/config.toml はそれらを 1 つの共有ストアに固定します。 Codex はメモリを読み取ります
登録された MCP サーバー ( get_session_state / remember );キャプチャの方向
Codex には同等の Stop フックがないため、プルベースです。
memlora codex-sync <project> は、~/.codex/sessions をスキャンしてロールアウトを探します。
記録された cwd はプロジェクトにマップされ、同じものを通じてデルタをキャプチャします。
抽出パイプライン (ロールアウト→トランスクリプト アダプターは唯一の Codex 固有のものです)
コード;デルタ/重複/冪等性は共有され、変更されません)。
ハンドオフ時に自動 — クロードの SessionStart は保留中の Codex を排出します
ブロックを構築する前にロールアウトし、MCP サーバーのキュー ドレイナーがプルします。
サイクルごとに新しいロールアウトが行われるため、ライブ クロード セッションでコーデックス側の決定が反映されます
次のセッションを待たずに。 Codex 側では、init が
AGENTS.md 命令 + ck-sync スキルにより、Codex はセッション開始時にプルします。
init は、.mcp.json (Claude) と .codex/config.toml の両方をプロビジョニングします。
(サーバーの cwd + プロジェクト環境が固定されている) + AGENTS.md (Codex)、
べき等に、既存の設定を壊すことなく。
memlora 医師がコーデックス ヘルス ラインを報告します (セッション ディレクトリ + ロールアウト)
count、または「同期するものがありません」 — Codex はオプションであるため、それが存在しない場合は

健康）。
Codex で行われた決定は次のクロード セッションのブロックに到達し、
その逆も同様です。アクション ポイント サーフェス (CK-1、PreToolUse ゲート) はクロードのみです -
Codex にはプロンプトごと/ツールごとのフックがありません。そのため、Codex ではループが共有ループに劣化します。
ブロック + MCP リコール。
MCP ツール (セッション ブロックは自動的に挿入されます。これらは対象を絞った使用を目的としています):
リコール · find_関連 · スケルトン · get_session_state
memlora init <project> — プロジェクトを登録し、セッションフックをインストールします
memlora ドクター [--strict] <プロジェクト> — サブシステムの健全性レポート
memlora codex-sync <project> — このプロジェクトの Codex CLI セッションをキャプチャします
memlora install-heads — トレーニングされたエンコーダ アーティファクト (salience + クロスエンコーダ ONNX ボディ) をインストールします。heads-v1 リリースおよび sha256 検証からダウンロードされるか、存在する場合はローカル モデル/エクスポートからコピーされます。
memlora show <プロジェクト> / memlora replace <プロジェクト> — 保存されたメモリの検査/クリア
uv sync # core (字句のみ)
uv sync --extra embedding # + 高密度取得 (fastembed + numpy)
uv で memlora init を実行します。 # このプロジェクトを登録 + フックをインストール
uv ラン メムローラ ドクター 。 # サブシステムの健全性
uv run memlora install-heads # オプション: トレーニングされたエンコーダー ヘッド (~270 MB ダウンロード)。
# それらがなければ、抽出/スーパーセッションはフォールバックします
# レガシー/字句パス — すべてが引き続き機能します
それから始めてください

[切り捨てられた]

## Original Extract

Persistent, structured project memory for Claude Code and Codex - event-sourced, typed, fail-open. Your agent stops re-deciding what you already decided. - KanishkNoir/cognikernel

Every coding session in claude code/ codex starts blindly. It forgets the archiectural decisions you made in a session you ran yesterday/ a week ago or a month ago. We end up depending on maintaining a Claude.md file or diff markdown files that these agents rely on their memory features. I started building CogniKernel to solve this problem. Instead of relying on markdown memory files, CogniKernel observes the coding session itself. It captures decisions, constraints, architectural decisions, project specific context etc info as they naturally emerge. The goal is simple: The agent should remember what it learned yesterday without me telling it what to remember. Some things it stores automatically: - Architectural decisions and design rationale
- Project conventions and coding patterns
- Important implementation details
- User preferences and recurring constraints
- Relationships between files, modules, and components The memory is local-first and designed to work alongside existing coding agents rather than replace them. A few design goals I had: - Zero manual memory curation
- Local storage instead of cloud dependency
- Fast retrieval with relevant context instead of dumping everything
- Agent-agnostic architecture so it can work beyond a single coding assistant
- Memory that improves over time instead of growing into an unsearchable log Was it impactful???
I build 4 full fledged projects using cognikernel while switching the auto-memory off from claude code and same project same prompt, same sessions using the auto memory - 30% Less Tokens used
- 4x less read tools used And a bonus!!! If you use claude code and codex both you can just start a session in alternate tool and your memory will carry forward there as well!!! This was a hobby project and if you use it and it helped your work, please let me know any feedback!

GitHub - KanishkNoir/cognikernel: Persistent, structured project memory for Claude Code and Codex - event-sourced, typed, fail-open. Your agent stops re-deciding what you already decided. · GitHub
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
KanishkNoir
/
cognikernel
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
148 Commits 148 Commits .github .github scripts scripts src/ memlora src/ memlora tests tests .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Persistent, structured project memory for Claude Code — and Codex. CogniKernel
watches a coding session through its hook surfaces, extracts the decisions,
constraints, and abandoned approaches worth keeping, consolidates them into an
event-sourced store, and injects them back as a compact context block the next
time you work — so the agent stops re-deciding what you already decided. The store
is keyed on the project path, so memory made in one tool travels to the other.
It is not a vector-database wrapper. It is an event-sourced log of typed
memory with lexical-primary retrieval, write-time consolidation, and a fail-open
reliability spine designed never to break your session.
And there is no LLM in the loop. Most memory tools work by sending your
transcripts to a generative model to "summarize what mattered" — another API
key, per-session token cost, added latency, and your session content leaving
the machine. CogniKernel treats extraction as classification, not generation :
a deterministic sanitize → classify → consolidate pipeline, with two small
fine-tuned encoder models (~130 MB ONNX, run locally on CPU in milliseconds)
scoring salience and detecting when a new decision supersedes an old one. No
API calls, no tokens billed, nothing leaves your machine. The only LLM involved
is the coding agent you already run — CogniKernel makes it remember.
Naming: CogniKernel is the project; memlora (package memlora-edge ) is
the Python module and CLI it ships as — the working name the code grew up
under. One project, two names: memlora init , memlora doctor , etc.
Everything CogniKernel does is one loop: observe → extract → consolidate →
store → retrieve → assemble → inject . The left rail captures; the right rail
recalls; the spine underneath keeps both honest.
+========================= CLAUDE CODE SESSION =========================+
| working memory - the context window the agent reasons in |
+--------------+-----------------------------------------+--------------+
| Stop hook captures transcript | inject block
v ^
+--------------+---------------+ +---------------+--------------+
| [1] EXTRACTION PIPELINE | | [6] COMPRESSION + INJECTION |
| sanitize -> classify -> | | authority-weighted budget, |
| salience (ONNX) -> | | drop-to-fit (keep every |
| decision-key + contracts | | constraint) -> block |
+--------------+---------------+ +---------------+--------------+
| enqueue | rank + fit
v ^
+--------------+---------------+ +---------------+--------------+
| [2] WORKER + CONSOLIDATION | | [5] RETRIEVAL |
| claim -> delta-merge -> | | FTS5 BM25 + dense -> RRF |
| supersede (latest-wins) -> | | prohibition_search (K1) |
| project (idempotent) | | skeleton graph (PageRank) |
+--------------+---------------+ +---------------+--------------+
| persist (atomic) | recall
v ^
+--------------+-----------------------------------------+--------------+
| [3] EVENT-SOURCED STORE * SQLite (WAL) |
| typed events | evidence | provenance | FTS5 | embeddings | ledger |
+----------------------------------------------------------------------+
| [4] RELIABILITY SPINE atomic migrations | idempotent replay | |
| doctor --strict | fail-open hooks | import-linter | CI gate |
+----------------------------------------------------------------------+
The four hook surfaces
CogniKernel attaches to a session at four points. Each is fail-open — if memory
is unavailable or errors, the hook logs at WARNING , returns cleanly, and the
session continues.
Memory is typed , not free-text chunks. The type drives ranking, rendering,
and supersession:
DECISION — a choice that was made ("use Redis for the rate limiter")
CONSTRAINT_HARD / CONSTRAINT_SOFT — rules, graded by deontic force
APPROACH_ABANDONED_DO_NOT_RETRY — a dead end, kept in the graveyard so
the agent doesn't re-attempt it
conventions, config facts, schema decisions (canonical role keys)
A decision key lets a later restatement supersede an earlier one
(latest-wins), so the store self-consolidates instead of accumulating
contradictions. An optional cross-encoder adds semantic supersession when the
encoder backend is installed.
Lexical-primary, with dense as a fused signal — never pure vector:
Hybrid core — FTS5 BM25 ∪ optional dense embeddings → Reciprocal Rank Fusion
prohibition_search — a type-restricted lexical pool so a "do not do X" rule
can't be crowded out by topically-similar prose at the moment you're about to do X
Skeleton — an AST symbol graph ranked by PageRank; find_related unions
semantic (embedding) neighbours with import-graph-adjacent events to
surface what a change touches (the semantic axis needs the embedding extra)
Golden-record consolidation at read — latest-wins reconciliation so recall
returns one coherent answer, not a pile of revisions
Benchmarked in a three-arm comparison — CogniKernel vs flat curated notes vs no
memory — with real agent sessions across four multi-session projects:
File reads: the universal win. The CogniKernel arm made the fewest file
reads in every project — typically 2–4× fewer (23 vs 63, 16 vs 47/53,
40 vs 89/83), and in the best case 3 reads vs 29 because the injected
block + AST skeleton carried the whole repo's shape. Fewer reads means fewer
tool round-trips and more of the context window left for actual work — your
session gets longer before compaction, not just cheaper.
Tokens: up to ~20% cheaper where memory matters. Price-weighted token
cost (cache-write 1.25×, cache-read 0.1×, output 5×) came out 18–23% lower
on projects with evolving decisions and cross-file dependencies — and roughly
a wash on small implementation-heavy projects where the code itself is cheap
to re-read. We publish the honest number, not the raw-token one (raw sums
look ~30–40% better, but ~95% of any session's bill is discounted cache-read).
Recall instead of re-derivation. Where memory earns its keep is projects
whose state is too large, too evolving, or too long-lived to re-derive
cheaply: the agent starts already knowing the decisions, constraints, and
dead ends, instead of spending the first quarter of the session rediscovering
them.
The system is designed to degrade legibly , never silently:
Atomic migrations — each numbered migration applies its body + version bump
in one transaction; safe to crash mid-script
Idempotent replay — a re-run worker job can't double-count or drift decay (evidence-provenance guard)
Fail-open hooks — every surface swallows its own failure and logs at WARNING ; silence never reads as success
memlora doctor --strict — per-subsystem health (schema, FTS5, embeddings, symbols, worker queue); non-zero exit when degraded
Architecture enforcement — import-linter layered contracts, guarded by a meta-test so a typo can't silently disable them
CI promotion gate — lint + full suite (incl. tests/reliability/ failure-injection) on every PR; see CONTRIBUTING.md
The store is platform-neutral — one SQLite DB per logical project, so Claude Code
and Codex working in the same directory share one memory. Project resolution is
alias-aware : C:\repo and /mnt/c/repo resolve to the same store, so memory
follows the checkout across Windows, WSL, and native mounts; for genuinely
different checkout paths, an opt-in project_identity key in
.memlora/config.toml pins them to one shared store. Codex reads memory through
the registered MCP server ( get_session_state / recall ); the capture direction
is pull-based , because Codex has no Stop -hook equivalent:
memlora codex-sync <project> scans ~/.codex/sessions for rollouts whose
recorded cwd maps to the project and captures the delta through the same
extraction pipeline (a rollout→transcript adapter is the only Codex-specific
code; delta/dedup/idempotency are shared and unchanged).
Automatic at the handoff — Claude's SessionStart drains pending Codex
rollouts before building the block, and the MCP server's queue drainer pulls
new rollouts each cycle, so a live Claude session picks up Codex-side decisions
without waiting for the next session; on the Codex side, init writes an
AGENTS.md instruction + a ck-sync skill so Codex pulls at session start.
init provisions both — .mcp.json (Claude) and .codex/config.toml
(with the server's cwd + project env pinned) + AGENTS.md (Codex),
idempotently and without clobbering existing settings.
memlora doctor reports a codex health line (sessions dir + rollout
count, or "nothing to sync" — Codex is optional, so its absence is healthy).
A decision made in Codex reaches the next Claude session's block, and
vice versa. The action-point surfaces (CK-1, PreToolUse gate) are Claude-only —
Codex has no per-prompt/per-tool hook — so on Codex the loop degrades to the shared
block + MCP recall.
MCP tools (the session block is injected automatically; these are for targeted use):
recall · find_related · skeleton · get_session_state
memlora init <project> — register the project and install the session hooks
memlora doctor [--strict] <project> — subsystem health report
memlora codex-sync <project> — capture Codex CLI sessions for this project
memlora install-heads — install the trained encoder artifacts (salience + cross-encoder ONNX bodies): downloaded from the heads-v1 release and sha256-verified, or copied from a local models/ export when present
memlora show <project> / memlora reset <project> — inspect / clear stored memory
uv sync # core (lexical-only)
uv sync --extra embedding # + dense retrieval (fastembed + numpy)
uv run memlora init . # register this project + install hooks
uv run memlora doctor . # subsystem health
uv run memlora install-heads # optional: trained encoder heads (~270 MB download);
# without them extraction/supersession fall back to
# the legacy/lexical path — everything still works
Then start

[truncated]
