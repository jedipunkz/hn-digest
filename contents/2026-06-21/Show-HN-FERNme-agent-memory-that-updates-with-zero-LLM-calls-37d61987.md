---
source: "https://github.com/mirkofr/FERNme"
hn_url: "https://news.ycombinator.com/item?id=48614107"
title: "Show HN: FERNme – agent memory that updates with ~zero LLM calls"
article_title: "GitHub - mirkofr/FERNme: A lightweight memory engine for AI agents using fuzzy graphs, Hebbian updates, and optional LLM gating. · GitHub"
author: "mirkofr"
captured_at: "2026-06-21T01:03:01Z"
capture_tool: "hn-digest"
hn_id: 48614107
score: 3
comments: 0
posted_at: "2026-06-20T23:34:03Z"
tags:
  - hacker-news
  - translated
---

# Show HN: FERNme – agent memory that updates with ~zero LLM calls

- HN: [48614107](https://news.ycombinator.com/item?id=48614107)
- Source: [github.com](https://github.com/mirkofr/FERNme)
- Score: 3
- Comments: 0
- Posted: 2026-06-20T23:34:03Z

## Translation

タイトル: HN を表示: FERNme – ~zero LLM コールで更新されるエージェント メモリ
記事のタイトル: GitHub - mirkofr/FERNme: ファジー グラフ、Hebbian 更新、およびオプションの LLM ゲーティングを使用する AI エージェント用の軽量メモリ エンジン。 · GitHub
説明: ファジー グラフ、Hebbian 更新、およびオプションの LLM ゲーティングを使用する AI エージェント用の軽量メモリ エンジン。 - mikofr/FERNme
HN テキスト: こんにちは、私は永続記憶に取り組んでいます。私は、脳のようなグラフベースの記憶システムが使用できるかどうか、そしてより重要なことに、llm トークンをどれだけ節約できるかを確認したかったのです。 FERNme は、Hebbian 共起規則によるファジー エッジを使用してメモリ タグを作成します。 FERNme は、エージェントがよりパーソナライズされた回答を提供するために使用できる、人々の優れた個人的な記憶になる可能性があると思います。コードは次の場所にあります: https://github.com/mirkofr/FERNme FERNme を真の個人的な脳にするために、開発者にそれをチェックしてテストし、正直な批評をしてもらいます。

記事本文:
GitHub - mirkofr/FERNme: ファジー グラフ、Hebbian 更新、およびオプションの LLM ゲーティングを使用する AI エージェント用の軽量メモリ エンジン。 · GitHub
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
ミルコフル
/
フェルンミー
公共
通知
通知設定を変更するにはサインインする必要があります
追加

ナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット .github/ workflows .github/ workflows デモ デモ ドキュメント ドキュメント シダの説明 シダの説明 fernme fernme ペーパー ペーパー テスト wiki wiki .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff COMPARISON.md COMPARISON.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス NAMING.md NAMING.md 通知通知 PAPER.md PAPER.md README.md README.md RELEASE_NOTES_v0.1.0.md RELEASE_NOTES_v0.1.0.md SECURITY.md SECURITY.md live_memory_test.py live_memory_test.py pyproject.toml pyproject.toml remember_live.py remember_live.py required.txt required.txt run_demo.py run_demo.py supernode_demo.py supernode_demo.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェント用のユーザー所有のほぼゼロの LLM メモリ層。これは、各人の行動 (話し方や感じ方など) から学習し、永遠にトークンフラットな状態を維持し、人々が記憶したものを表示、編集、所有できるようにします。このエンジンは基質に依存しません。現在の Web サイト (ショッピング、サポート、予約、ヘルスケア、家庭教師、政府)、次のデスクトップとモバイルなど、エージェントが動作する場所をすべて記憶します。
書き込みが安い、読みやすい、設計により解釈可能、ユーザーが所有
エージェントのメモリのほとんどは、ターンごとに LLM によって書き込まれ (高価で幻覚が起こりやすい)、質問応答 (アクションではなく) で評価され、単一のユーザーを想定しています。 FERNme は、その反対の世界、つまりあらゆる領域 (販売、予約、解決されたチケット、完了したレッスン、約束の維持など) で多くの人のために行動するエージェント、つまり「結果」とは目的が何であれ、その世界向けに構築されています。それは、現在エージェントがすでに活動している場所、つまり Web サイトから始まり、同じユーザー所有のメモリがデスクトップとモバイルに拡張されるように設計されています。

各ユーザーは、サイトごとのグラフにおける疎な、あいまいに重み付けされたノードです。エッジは、LLM 呼び出しがゼロの Hebbian 共起規則によって更新され、検索は活性化を広げ、プロンプトに面した「カード」には以前の母集団からの偏差のみが保存されます。その結果、プロファイルが何年も成長してもターンあたりのコストは一定になり、ユーザーは自分のメモリを読み取って修正することができ、同じエンジンがユーザーの同意がある場合にのみ、完全に制御するクロスサイト スーパーノードに組み立てられます。
🎯 FERNmeを選ぶ理由（長所）
🪶 ゼロ LLM 書き込み
メモリの更新はグラフ上で算術的に行われます。対話ごとに 0 回の LLM 呼び出しが行われるのに対し、抽出ベースのメモリの場合は最大 2 回です。書き込み時のコストも書き込み時の幻覚もありません。
📉 トークンコストは永久にフラット
プロンプト カードには、訪問初日であっても 5 年目であっても、最大 25 個のトークンが保持されます。全履歴ベースラインは、120 回のインタラクションにより 77 倍大きくなります。
🧠 どの体制でも強い
静的リコールでは周波数カウンターと同点、ドリフトでは 0.72 → 0.13 で勝ち、コンテキストでは勝ち (0.62 → 0.51)。減衰 + 拡散活性化により、安定性と適応性が統合されます。
🪟 ガラスボックスとユーザー所有
すべての設定が表示され、編集可能です。人々は問題を修正したり、すべてを削除したり、エクスポートしたりします。プライバシーは義務ではなく機能になります。
🏬 結果を重視して構築
QA ではなく、コンバージョンによって評価されます。シミュレートされたストアフロントでは、パーソナライズされていないレコメンデーションと比較して +16% のコンバージョン上昇率が示されています。
🧩 ユーザー所有のスーパーノード
複数のサイトにサインイン → 思い出はレゴのように、あなたが管理する 1 つのプロファイルに集まり、デフォルトで拒否され、機密データは壁で囲まれます。監視ではなく、その鏡像です。
🎚 コスト/品質ダイヤル
1 つのエンジン、memory_mode スイッチ: デフォルトでは無料のキーレス純粋、Mem0 グレードのニュアンスが必要な場合はオプトインのゲート/オフライン LLM エンリッチメント — 使用したコンピューティングに対してのみ支払います。
🔐 検証可能かつ

学習不可能な
すべてのアクションは改ざん防止 HMAC チェーンに記録され、ユーザーはそれを再生して変更を検出できます。 remember_everywhere はプロフィールを消去し、その人物を以前の集団から解き放ちます。これは、忘れられる権利であることが証明されています。
🛡 耐注入設計
書き込みは算術演算であり、LLM 抽出ではないため、ページ/ユーザー テキストに「話しかける」ことはできないため、挿入された命令がメモリに入らないことがテスト済みです。
🧠 民間の集合知
新規ユーザーは、k-匿名性と差分プライバシーを備えたターン 1 のクラウド パターン (前の母集団からのコールド スタート) の恩恵を受けるため、個人の漏洩はありません。シングルユーザーの思い出にはありえないネットワーク効果の堀。
🗣 スタイルと気分の記憶
各人がどのようにコミュニケーションするかを学びます（簡潔/冗長、フォーマル/カジュアル、英語）
[切り捨てられた]
正直な範囲: 以下の数値は合成データまたは LLM が作成したデータに関するものであり、実際のものではありません
ユーザー。彼らはメカニズムを検証し、障害を表面化します。本物の人間のパイロットは、
次のステップが保留中です。 Mem0 (LLM) ヘッドツーヘッドには API キーが必要ですが、まだ実行されていません。
LLM で作成された人物について (実際のエージェントによる取り込みに最も近い)
92 件の三人称プロファイル (ChatGPT 作成) のうち 16 件のサンプル。散文のみとして読み取れます。
エージェント的に記憶され、隠された回答キーに対してスコアが付けられます。
(ここでの「エージェント」は LLM の散文を読むため、これらはエージェント + エンジンを一緒に反映しています。
エンジンはしっかりしています。抽出品質はエージェントのものです。)
コスト、リコール、パレート (合成、マルチシード)
再現: python -m fernme.eval.cost_variance · ... 品質 · ... ドリフト · ... コンテキスト · ... アブレーション · ... パイロット
コスト — ターンごとのメモリ トークンとプロファイル サイズ (5 シード):
再現品質 — 精度@5 対 グラウンドトゥルースの好み (5 シード × 40 ユーザー):
見出し: FERNme はどこでも強力な唯一の方法です。周波数は忘れられません（ドリフトに失敗します）。最近の私

■ ノイズが多い (静的に失敗する)。 FERNmeの減衰+拡散活性化の両方を取得します。
コールドスタート アブレーション — 事前人口によりターン 1 ～ 3 で +0.06 の精度 @5 が得られ、ターン 10 までに消失します (実際の、しかしささやかな、コールド スタートのみの利点)。
コスト/品質パレート ( python -m fernme.eval.pareto ) — 測定された FERNme リコールと
トークン、モデル化された LLM のニュアンスと価格 (ファイル内の仮定)。インタラクション 1,000 件あたり:
FERNme+ゲート付き/オフラインで効率的に膝の上に座る: LLM 天井の品質の約 80 ～ 90%
コストが 1 ～ 2 桁安くなります。 (モデル化された仮定。重要なのは形状です。)
シミュレーション結果のパイロット - 偽の店頭、行動から学ぶ買い物客: 人気ベースラインと比較して相対コンバージョン上昇率 +16%。訪問 1 (コールド スタート) で同点となり、学習しながら前進し、パイロット途中の味のドリフトを乗り越えて回復しました。
🎚 メモリモード (1 つのエンジン、コスト/品質ダイヤル)
FERNme には、デプロイメントレベルのスイッチ FernService(memory_mode=...) を備えた 1 つのコアが同梱されています。
デフォルトは無料、キーレス、テスト済みです。 LLM モードはオプトインでプラグイン可能です。
プラグ可能なタガー ( tagging.py ) が LLM の作業を行います。オプションで llm_fn を渡します
制御された語彙に制限されます (モデル間の実際の一貫性レバー)。
ホット ライト パスは、どのモードでも LLM フリーのままです。 Gated は、次の場合にのみ通話を費やします。
決定論的マッピングでは何も検出されず、svc.llm_calls はすべての呼び出しをカウントします。
コストの透明性。
各モードがどこに到達するかについては、上記のコスト/品質パレートを参照してください。正直なところ: ゲート付き/
オフライン品質は、実際のモデルに対して実行するまでモデル化されます。配線はテストされます。
ここではモック LLM を使用していますが、品質は検証されていません。
🧭 9つのリープフロッグ次元（ステータス）
FERNme の優位性はメカニズムではありません (これは現在、2026 年のカテゴリーで混雑しています) - 競争しているのです
シングルユーザー、ベンダー所有、リコール最適化システムでは構造的に不可能です。
テス

e は、HippoGraph らによって意図的に作成されたものです。フォローできません: シングルユーザーです
(集団事前分布なし)、ベンダー所有 (ユーザー所有のプロトコルなし)、リコール最適化
(結果ループなし)。正直にテストされたスライスが組み込まれており、研究に依存したスライスにはマークが付けられています。
フローチャート TD
V[Web サイトの訪問者] -->|プロンプト + アクション| API[FERNmeサービス]
API --> CONSENT{同意?}
同意 -->|いいえ| STOP[ブロックされました]
同意 -->|はい|エンジン
サブグラフ ENGINE[エンジン - 書き込みパスに LLM がありません]
W[ヘビアン書き込み + 減衰] --> G[(サイトごとのプリファレンス グラフ<br/>ファジー 0 ～ 9 エッジ)]
G --> R[拡散活性化検索]
R --> CARD[トークン最小カード ~25 トークン]
PRIOR[事前の差分エンコーディングの母集団] --> R
終わり
CARD --> AGENT[エージェント: 推奨 / 行動]
G --> CAB[(キャビネット:生のイベント ログ)]
API --> STORE[(SQLite または Postgres<br/>マルチテナント)]
API --> GLASS[🪟 Glass-box エディター]
API - ユーザーがサインインします。 -> SUPER[ユーザー所有のスーパーノード<br/>クロスサイト、デフォルト拒否]
読み込み中
🧠 FERNme の仕組み (ビジュアルウォークスルー)
FERNme を使用する理由 — ループ内での高価な RAG/ベクトル取得の代わりに、適応型ローカル メモリを使用します。
何が違うのか - ほぼゼロの LLM、決定性優先、ヘビアン、ファジー、メモリ カード、アクション認識、ユーザー所有。
記憶がどのように成長するか — 新しいイベント → 接続 → 強化 → 減衰 → カードの更新 (ヘビアン学習)。
ファジーヘビアングラフ - まばらな重み付けされた (0 ～ 9) エッジ。ユーザー、設定、トピック、目標のノード。
LLM ゲート — デフォルトではなく例外です。ほとんどのイベントは決定論的に処理されます。 LLM は、不確実な場合にまれに代替手段として使用されます。
メモリ カード — 重要なことについての、制限があり、解釈可能な、トークンを最小限に抑えた要約。
行動を意識した学習 — 良い結果はつながりを強化し、悪い結果はつながりを弱めます。
これからの道 - 今日の地元の記憶。明日の再帰的組織とユーザーフロー

ned スーパーノード (ロードマップ、まだ構築されていません)。
完全なアーキテクチャ: インジェスト ブリッジ → 名前空間ボキャブラリー → ファジー ヘビアン グラフ → メモリ カード → エージェント。不確実な場合にのみ LLM ゲートを使用します。
pip install -e " .[dev,api] "
python run_demo.py # コールドスタート → 学習 → ガラスボックス編集
python supernode_demo.py # 1 人のユーザー、3 つのサイト、1 つの所有プロファイル
pytest -q # 88 テスト (エンジン、ストア、スーパーノード、安全性、認証…)
# 実験
python -m fernme.eval.drift # 好みが変わると FERNme が周波数カウンターを破る
python -m fernme.eval.pilot # シミュレートされたコンバージョン上昇率 +16%
# ライブで実行する
FERNME_API_KEY=secret uvicorn fernme.api.rest:app --port 8077 # REST API (/docs にあるドキュメント)
http://localhost:8077/ui # ガラスボックスメモリエディタを開く
http://localhost:8077/graph を開く # あなたの記憶をグラフとして — サイト / PC / 電話ごとに焦点を当てる
python -m fernme.api.mcp_server # エージェント/クロード用の MCP サーバー
🗄 ストレージ: デフォルトは ~/.fernme/fernme.db (SQLite) です。運用環境では、PostgresStore を使用します。同じインターフェイスで、実際の Postgres 16 に対してテストされています。SQLite をクラウド同期フォルダーから遠ざけてください。
エンジン - 飽和ヘビアン書き込み (LLM なし)、ACT-R 減衰、拡散アクティベーション、トークン最小カード。
前の人口 — IDF コールド スタート。差分 (偏差のみ) ストレージは
明示的な prune_to_prior パス (冗長エッジは前のエッジまで読み取られます) によって強制されます。
ストア — SQLiteStore (ゼロセットアップ) および Postgres

[切り捨てられた]

## Original Extract

A lightweight memory engine for AI agents using fuzzy graphs, Hebbian updates, and optional LLM gating. - mirkofr/FERNme

Hi I have been working on persistent memory. I wanted to see whether a brain like graph-based memory system could be used and more importantly how much we could save on llm tokens. FERNme uses fuzzy edge with Hebbian co-occurrence rule to create memory tags. I think FERNme could become a great personal memory of people that can be used by agents to give more personalized answers. The code is placed in: https://github.com/mirkofr/FERNme I invite developers to to check and test it and give honest criticism to make FERNme a real personal brain.

GitHub - mirkofr/FERNme: A lightweight memory engine for AI agents using fuzzy graphs, Hebbian updates, and optional LLM gating. · GitHub
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
mirkofr
/
FERNme
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .github/ workflows .github/ workflows demo demo docs docs explanation of fern explanation of fern fernme fernme paper paper tests tests wiki wiki .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff COMPARISON.md COMPARISON.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NAMING.md NAMING.md NOTICE NOTICE PAPER.md PAPER.md README.md README.md RELEASE_NOTES_v0.1.0.md RELEASE_NOTES_v0.1.0.md SECURITY.md SECURITY.md live_memory_test.py live_memory_test.py pyproject.toml pyproject.toml recall_live.py recall_live.py requirements.txt requirements.txt run_demo.py run_demo.py supernode_demo.py supernode_demo.py View all files Repository files navigation
A user-owned, near-zero-LLM memory layer for AI agents. It learns each person from their behavior — including how they talk and feel — stays token-flat forever, and lets people see, edit, and own what's remembered. The engine is substrate-agnostic: it remembers wherever an agent acts — websites today (shopping, support, booking, healthcare, tutoring, gov), desktop and mobile next.
Cheap to write · flat to read · interpretable by design · owned by the user
Most agent memory is written by an LLM on every turn (expensive, hallucination-prone), evaluated on question-answering (not actions), and assumes a single user . FERNme is built for the opposite world — agents that act for many people, in any domain (a sale, a booking, a resolved ticket, a completed lesson, a kept appointment — "outcome" is whatever the goal is). It starts where agents already act today — websites — and the same user-owned memory is designed to extend to desktop and mobile. Each user is a sparse, fuzzily-weighted node in a per-site graph; edges update by a Hebbian co-occurrence rule with zero LLM calls , retrieval is spreading activation , and the prompt-facing "card" stores only deviations from a population prior . The result: per-turn cost stays flat as a profile grows for years, the user can read and correct their own memory, and the same engine assembles — only with the user's consent — into a cross-site supernode they fully control.
🎯 Why FERNme (the strong points)
🪶 Zero-LLM writes
Memory updates are arithmetic on a graph — 0 LLM calls per interaction vs. ~2 for extraction-based memory. No write-time cost, no write-time hallucination.
📉 Flat token cost forever
The prompt card holds ~25 tokens whether it's a visitor's first day or fifth year. A full-history baseline is 77× larger by 120 interactions.
🧠 Strong in every regime
Ties a frequency counter on static recall, beats it 0.72 → 0.13 on drift , and wins on context (0.62 → 0.51). Decay + spreading activation unify stability and adaptivity.
🪟 Glass-box & user-owned
Every preference is visible and editable. People fix what's wrong, delete everything, or export it. Privacy becomes a feature, not a liability.
🏬 Built for outcomes
Evaluated by conversion , not QA. A simulated storefront shows +16% conversion lift vs. non-personalized recommendations.
🧩 User-owned supernode
Sign in across sites → your memories assemble like Lego into one profile you control , default-deny, sensitive data walled off. Not surveillance — the mirror image of it.
🎚 Cost/quality dial
One engine, a memory_mode switch: free key-less pure by default, opt-in gated / offline LLM enrichment when you need Mem0-grade nuance — pay only for the compute you use.
🔐 Verifiable & unlearnable
Every action is logged in a tamper-evident HMAC chain the user can replay to detect any alteration; forget_everywhere wipes the profile and unlearns the person from the population prior — provable right-to-be-forgotten.
🛡 Injection-proof by design
Writes are arithmetic, not LLM extraction, so page/user text can't be "talked into" becoming a belief — tested that injected instructions never enter memory.
🧠 Private collective intelligence
New users benefit from crowd patterns on turn one (cold-start from a population prior), with k-anonymity + differential privacy so no individual leaks. A network-effect moat single-user memories can't have.
🗣 Style & mood memory
Learns how each person communicates (terse/verbose, formal/casual, en
[truncated]
Honest scope: the numbers below are on synthetic or LLM-authored data, not real
users. They validate the mechanism and surface failures; a real-human pilot is the
pending next step. The Mem0 (LLM) head-to-head needs an API key and is not yet run.
On LLM-authored people (closest to real, agentic ingestion)
A sample of 16 of 92 third-person profiles (ChatGPT-authored), read as prose only and
remembered agentically, then scored against hidden answer keys:
(The "agent" here is an LLM reading prose, so these reflect agent + engine together — the
engine is solid; the extraction quality is the agent's.)
Cost, recall, and Pareto (synthetic, multi-seed)
Reproduce: python -m fernme.eval.cost_variance · ... quality · ... drift · ... context · ... ablation · ... pilot
Cost — per-turn memory tokens vs. profile size (5 seeds):
Recall quality — precision@5 vs. ground-truth preferences (5 seeds × 40 users):
The headline: FERNme is the only method strong everywhere. Frequency can't forget (fails drift); recency is noisy (fails static). FERNme's decay + spreading activation get both.
Cold-start ablation — population prior gives +0.06 precision@5 at turns 1–3 , washing out by turn 10 (a real but modest, cold-start-only benefit).
Cost / quality Pareto ( python -m fernme.eval.pareto ) — measured FERNme recall &
tokens, modeled LLM nuance & price (assumptions in-file). Per 1,000 interactions:
FERNme+gated/offline sit on the efficient knee: ~80–90% of the LLM-ceiling quality
at 1–2 orders of magnitude less cost. (Modeled assumptions; shape is the point.)
Simulated outcome pilot — fake storefront, learn-from-behavior shoppers: +16% relative conversion lift over a popularity baseline; tied at visit 1 (cold start), pulling ahead as it learns, recovering through a mid-pilot taste drift.
🎚 Memory modes (one engine, a cost/quality dial)
FERNme ships one core with a deployment-level switch — FernService(memory_mode=...) .
The default is free, key-less, and tested; LLM modes are opt-in and pluggable.
A pluggable tagger ( tagging.py ) does the LLM work; you pass llm_fn , optionally
constrained to a controlled vocabulary (the real consistency lever across models).
The hot write path stays LLM-free in every mode ; gated spends a call only when the
deterministic mapping finds nothing, and svc.llm_calls counts every invocation for
cost transparency.
See the cost/quality Pareto above for where each mode lands. Honest note: the gated/
offline quality is modeled until run against a real model — the wiring is tested
here with a mock LLM, not validated for quality.
🧭 The 9 leapfrog dimensions (status)
FERNme's edge isn't the mechanism (that's now a crowded 2026 category) — it's competing
on dimensions single-user, vendor-owned, recall-optimized systems structurally can't .
These are deliberately the things HippoGraph et al. can't follow: they're single-user
(no collective priors), vendor-owned (no user-owned protocol), and recall-optimized
(no outcome loop). Built in honest, tested slices — research-dependent ones are marked.
flowchart TD
V[Visitor on a website] -->|prompt + action| API[FERNme Service]
API --> CONSENT{consent?}
CONSENT -->|no| STOP[blocked]
CONSENT -->|yes| ENGINE
subgraph ENGINE[Engine - no LLM in the write path]
W[Hebbian write + decay] --> G[(Per-site preference graph<br/>fuzzy 0-9 edges)]
G --> R[Spreading-activation retrieval]
R --> CARD[Token-minimal card ~25 tok]
PRIOR[Population prior<br/>differential encoding] --> R
end
CARD --> AGENT[Agent: recommend / act]
G --> CAB[(Cabinet: raw event log)]
API --> STORE[(SQLite or Postgres<br/>multi-tenant)]
API --> GLASS[🪟 Glass-box editor]
API -.user signs in.-> SUPER[User-owned Supernode<br/>cross-site, default-deny]
Loading
🧠 How FERNme works (visual walkthrough)
Why FERNme — adaptive local memory instead of expensive RAG/vector retrieval in the loop.
What makes it different — near-zero-LLM, deterministic-first, Hebbian, fuzzy, memory cards, action-aware, user-owned.
How memory grows — new event → connect → strengthen → decay → update the card (Hebbian learning).
The fuzzy Hebbian graph — sparse, weighted (0–9) edges; nodes for users, preferences, topics, goals.
The LLM gate — an exception, not the default. Most events are handled deterministically; the LLM is a rare fallback when uncertain.
The memory card — a bounded, interpretable, token-minimal summary of what matters.
Action-aware learning — good outcomes strengthen connections, bad outcomes weaken them.
The road ahead — today's local memory; tomorrow's recursive organization and user-owned supernode (roadmap, not yet built).
Full architecture: ingestion bridge → namespaced vocabulary → fuzzy Hebbian graph → memory card → agent, with the LLM gate only when uncertain.
pip install -e " .[dev,api] "
python run_demo.py # cold-start → learning → glass-box edit
python supernode_demo.py # one person, three sites, one owned profile
pytest -q # 88 tests (engine, store, supernode, safety, auth…)
# experiments
python -m fernme.eval.drift # FERNme beats a frequency counter when tastes change
python -m fernme.eval.pilot # +16% simulated conversion lift
# run it live
FERNME_API_KEY=secret uvicorn fernme.api.rest:app --port 8077 # REST API (docs at /docs)
open http://localhost:8077/ui # glass-box memory editor
open http://localhost:8077/graph # your memory as a graph — focus by site / PC / phone
python -m fernme.api.mcp_server # MCP server for agents/Claude
🗄 Storage: defaults to ~/.fernme/fernme.db (SQLite). For production use PostgresStore — same interface, tested against a real Postgres 16. Keep SQLite off cloud-synced folders.
Engine — saturating Hebbian write (no LLM), ACT-R decay, spreading activation, token-minimal card.
Population prior — IDF cold-start; differential (deviation-only) storage is
enforced by an explicit prune_to_prior pass (redundant edges read through to the prior).
Stores — SQLiteStore (zero-setup) and Postgre

[truncated]
