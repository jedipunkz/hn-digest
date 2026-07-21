---
source: "https://github.com/oteomamo/SALT"
hn_url: "https://news.ycombinator.com/item?id=48988139"
title: "Looking for contributors – trie based memory efficient LLM runner"
article_title: "GitHub - oteomamo/SALT: Salience-aware lexical trie for long-context compression. · GitHub"
author: "omamo"
captured_at: "2026-07-21T04:59:58Z"
capture_tool: "hn-digest"
hn_id: 48988139
score: 1
comments: 0
posted_at: "2026-07-21T04:45:07Z"
tags:
  - hacker-news
  - translated
---

# Looking for contributors – trie based memory efficient LLM runner

- HN: [48988139](https://news.ycombinator.com/item?id=48988139)
- Source: [github.com](https://github.com/oteomamo/SALT)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T04:45:07Z

## Translation

タイトル: 貢献者募集 – トライベースのメモリ効率の高い LLM ランナー
記事のタイトル: GitHub - oteomamo/SALT: 長いコンテキスト圧縮のための Salience を意識した語彙トライ。 · GitHub
説明: 長いコンテキスト圧縮のための Salience を意識した語彙トライ。 - オテオマモ/SALT

記事本文:
GitHub - oteomamo/SALT: 長いコンテキスト圧縮のための Salience 対応の語彙トライ。 · GitHub
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
オテオマモ
/
塩
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション

n 個のオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
164 コミット 164 コミット .github .github docs docs ソルト ソルト スクリプト スクリプト .gitignore .gitignore COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md 環境.yml 環境.yml eval.py eval.py mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
長期コンテキスト圧縮のための顕著性を意識した字句トライ
MCP サーバー - AI クライアントが SALT を
REPL なしでの会話の記憶。
テールアウェアなメモリ選択 - テキストにメモリ バジェットを費やすのをやめます
モデルはすでに最近のメッセージを読み取っています。
増分圧縮 - 代わりに前のターンの作業を再利用します。
毎ターン選択全体を再構築します。
メモリ スイッチの段階的設定 - どのメモリをデフォルトでオフにするかを決定する
動作はデフォルトになり、実際のセッションの数値が使用されます。
スクリプト化された会話が実行されます - より豊富なツールが --turns に役立ちます
長い缶詰の会話を推進し、スコアを記録します。
SALT は、長いドキュメントを言語に送信する前に固定サイズに縮小します。
モデルを作成し、最も多くの情報を伝える文を保持します。どれでも動作します
モデルを作成し、より短いプレーンテキスト プロンプトを生成し、コンピューティング、メモリ、および負荷を削減します。
長い入力には待ち時間がかかります。
問題。プロンプトが長すぎる場合、既存のコンプレッサーはそれぞれ
関連性スコアを 1 つだけ与え、予算が達成されるまで上位スコアを保持します。
なくなります。予算が限られている場合、これによりドキュメントの主要なトピックが内容を飲み込むことができます。
予算全体が非常に小さいため、それでも重要なポイントが削除されます - と呼ばれる失敗
テーマの折りたたみ (たとえば、マルチホップの質問では、テーマを維持できます)
主要な実体についての文章はありますが、その実体を失っています

にリンクするエントリ
2番目）。
解決策。 SALT は、まずドキュメントの繰り返しのテーマを整理してマッピングします。
各文のキーワードをトライ（頻度順に並べた小さなキーワード ツリー）にまとめる
これらのキーワードが繰り返され、予算がテーマのブランチ全体に分散されます。
文章を選択する前に、マイナーなテーマがシェアを維持するのではなく、
混雑している。テーマ マップは一度構築されるため、全体にわたって再利用できます。
文書を再読せずに会話を終了します。
文書リリースで説明されている以前のレガシー セレクターには v1.0.0 というタグが付けられています。
main のデフォルトは、以下に説明するカバレッジ/CELF セレクターになりました。
2 段階。インデックス作成では、ドキュメントを 1 回読み取り、キーワード トライを構築します。
繰り返し発生するテーマの再利用可能なマップ。選択により文のサブセットが選択されます
トークンバジェットに基づいて、クエリバイアスの有無にかかわらず、元のドキュメントの順序で返されます。
インデックス作成 -- ドキュメントごとに 1 回、ターンや予算全体で再利用
ドキュメント → 分割 + 迷惑フィルター
→ 文ごとのキーワード (BGE-small [CLS] アテンション + ニーカットオフ)
→ テーマの顕著性 (SF = 単語を保持する #sentences、上位四分位)
→ キーワードトライ（各文のテーマ、SF順、形式
ルートから葉へのパス、葉には文 ID が保持されます)
選択 ── クエリの有無にかかわらず、予算ごとに
CELF でテーマの範囲を最大化します。lazy-greedy: ピックの価値は、
テーマのブランチがいっぱいになるため、予算が折りたたまれるのではなくテーマ全体に分散されます。
支配的なものに。クエリはトライの重み付けを変更します (語彙 + BGE セマンティック)。
再構築せずに。 → 圧縮されたプロンプト、元の注文、≤ 予算
システム全体を青写真として - このマップは SALT が成長しても最新の状態に保たれます。
したがって、変更がどこに属しているかを見つける最も速い方法です。
┌─────────────────

─────────────────┐
│ 塩 │
│ │
│ ┌───────────┐ ┌─────────┐ ┌─────────┐ │
│ │ インデックス作成 │ │ キーワードトライ │ │ 選択 │ │
│ │ │ │ │ │ │
│ │ BGE 小型エンコーダ │ │ SF 順序パス │ │ カバレッジ (CELF) │ │
│ │ 注目キーワード │ │ テーマ支店 │ │ 支店割引 │ │
│ │ ニーカットオフ │ │ §ファイル: doc ブランチ│ │ マルチアンカークエリ │ │
│ │ ジャンクフィルター │ │ 安く再構築 │ │ ≤ ワードバジェット │ │
│ ━━━━━━━━━━━━━━━━━━━━━━━┘ ━━━━━━━━━━━━━━━━━┘ │
│ │
│ ┌───────────┐ ┌─────────┐ ┌─────────┐ │
│ │ セッショントライ │ │ 即時集合 │ │ チャットランナー │ │
│ │ │ │ │ │ │
│ │ 会話ごと │ │ 安定したプレフィックスが最初│ │ HF ストリーミング │ │
│ │ DRAM 内に存在します │ │ 追加専用末尾 │ │ vLLM + APC (オプトイン)│ │
│ │ ターンごとに成長します │ │ 記憶 + 質問 │ │ vllm-serve クライアント │ │
│ │ クロスターンカバレッジ│ │ 説明書.md │ │ モデルレジストリ │ │
│ │ + 半減期減衰 │ │ │ GPU 固定モデル │ │
│ │ + ニアダップゲート │ │ │ │ │
│ │ + バックグラウンド取り込み│ │ │ │ │
│ └

───────────┘ └───────────┘ └─────────────┘ │
│ │
│ ┌─────────────────────────────┐ │
│ │ ドキュメントの取り込み (salt@ ファイル、salt --doc) │ │
│ │ pypdf の抽出 · 家具のスクラブ · フロート間で段落を再結合│ │
│ │ 表 + 疑似コードをキャプションの下にグループ化し、脚注を分離 │ │
│ │ 見出し、パネルラベル、方程式は保持、参考文献リストは削除 │ │
│ ━━━━━━━━━━━━━━━━━━━━━━━━┘ │
│ │
│ ┌─────────────────────────────┐ │
│ │ トライシェイプ - 根が会話を結ぶ │ │
│ │ ● root - 会話バインド │ │
│ │ ┌─────────────┼─────────────┐ │ │
│ │ §file:paper.pdf §file:notes.txt 会話 │ │
│ │ │ │ ┌────┴────┐ │ │
│ │ キーワード パス キーワード パス テーマ A テーマ B │ │
│ │ │ │ │ │ │
│ │ 文 文 文
[切り捨てられた]
Python 3.10 と CUDA GPU が必要です (CPU は圧縮のために動作しますが、速度が遅くなります)。
git

クローン https://github.com/oteomamo/SALT.git
cdソルト
2. 環境を作成する
conda env create -f 環境.yml
conda アクティベートソルト
または venv を使用して:
python3.10 -m venv .venv
ソース .venv/bin/activate
pip install -r 要件.txt
3. SALTを編集可能モードでインストールする
pip install -e 。
これにより、次の 2 つのコンソール コマンドもインストールされます。salt (ワンショット圧縮、
「使用法」を参照) および SaltChat (対話型チャット、「使用法」を参照)
チャットボットモード)。
4. ハグフェイスによる認証 - 評価モデル
( metal-llama/Llama-3.1-8B-Instruct ) はゲートされています。
hf 認証ログイン
または、CLI をスキップして、トークンを直接エクスポートします (export HF_TOKEN=hf_... )。
5. (オプション) vLLM バックエンド。 eval.py のデフォルトは vLLM です。
SaltChat --backend vllm はプレフィックス キャッシュに使用し、saltServe
それを使用して永続モデルサーバーを起動します。塩の中に取り付けます
環境:
pip install " vllm==0.11.0 " " prometheus-fastapi-instrumentator>=8.0.1 "
2 番目のピンは、新しい fastapi の次にサーバーのルートを健全に保ちます。
リリースします。これをスキップして、ポータブル実行のために eval.py --backend hf を実行してください
vLLM は必要ありません。 SaltChat はすでにデフォルトで HF バックエンドを使用しています。
SaltServe は、別の環境にインストールされた vLLM を実行することもできます
--vllm-bin を通じて。
bash scripts/setup_env.sh はステップ 2 ～ 3 を 1 回で実行します (WITH_VLLM=1 を追加します)
vLLM を含む)。
LongBench データをフェッチし、存在するすべてのタスクを圧縮して、
1 つのコマンドで出力します。
Python ソルト/データセット/download_longbench.py
bash スクリプト/run_datasets.sh
煙テスト (タスクごとに 5 つのサンプル、圧縮のみ):
MAX_SAMPLES=5 RUN_EVAL=0 bash スクリプト/run_datasets.sh
スクリプトは各データセットを適切なモードに自動的にルーティングし、スコアを付けます。
結果。すべてのコマンド、フラグ、およびノブは、
使い方と
データセットのページ。
SaltChat は対話型チャット REPL であり、SALT は会話メモリです。
変換者ごとに 1 つの永続的なトライ

ステーションはあらゆる交換（そしてあらゆる交換）とともに成長します。
添付文書)、ターンごとに蓄積された履歴を圧縮します。
トークン バジェット内のクエリに偏ったコンテキスト ブロック。
SaltChat --model qwen05 --conversation-id Demon1 --doc report.txt
SaltServe で起動された永続サーバーはモデルをロードしたままにし、
チャット間でキャッシュが温まるため、再開された会話は問題なく再開されます。
その文書を再読します。の
チャットボットモードのガイドカバー
コンセプト、サービス
このページではサーバーについて説明します。
オプションページにはすべてのものがリストされています
フラグをそれぞれ 1 行に記述します。これには、デフォルトでオフになるスイッチも含まれます。
長いセッションの方が良いです。
SALT (カバレッジ/CELF セレクター) は LongBench 全体の平均 44.60 に達しました
Llama-3.1-8B を使用 - 20% のトークン予算で命令します。データセットごとの完全なテーブル
は結果ページにあります。
MCP サーバー - 圧縮とセッションを公開する Salt-mcp エントリ ポイント
AI クライアント (Claude Code、Claude Desktop、Cursor) が使用できるツールとしてのメモリ
REPL を使用しない会話の記憶として SALT を使用します。
テールを意識したメモリ選択 - モデルがすでに存在するセンテンスをスキップします
最近のメッセージを逐語的に読むため、メモリ バジェットが新しく購入されます。
画面上にあるものを繰り返すのではなく、コンテキストを確認します。
増分圧縮 - 前のターンの選択作業を引き継ぎます
毎回すべてをやり直すのではなく、追加のみの会話を転送します。
回ります。
記憶スイッチの卒業 - いくつかの記憶動作がオフになります
デフォルトでは (
オプションページ）
実際のセッションの /stats 数値によって、どれがデフォルトになるかが決まります。
スクリプト化された会話が実行されます - --turns に関するより豊富なツール
定型化された会話はセッションを長引かせ、後で採点される可能性があります。
要約範囲 - テーマ範囲の目標をより良いものに拡張します
要約に役立ちます。多くの小さなテーマを思い出すことが最も重要です。
PR 歓迎 - COTRIBUTING.md を参照

。
このプロジェクトがあなたの研究に役立つと思われる場合は、私たちの論文を引用することを検討してください。
@misc{mamo2026saltsalienceawarelexicaltrie、
title={SALT: 長いコンテキスト圧縮のための Salience-Aware 字句トライ},
author={Oteo Mamo、Hyunjin Yi、Joydhriti Choudhury、Shangqian Gao、Weikuan Yu}、
年={2026}、
eprint={arXiv:2607.17486}
}
📄ライセンス
SALT は MIT ライセンスに基づいてリリースされています。
長いコンテキスト圧縮のための Salience を意識した語彙トライ。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
塩チャット
最新の
2026 年 7 月 4 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Salience-aware lexical trie for long-context compression. - oteomamo/SALT

GitHub - oteomamo/SALT: Salience-aware lexical trie for long-context compression. · GitHub
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
oteomamo
/
SALT
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
164 Commits 164 Commits .github .github docs docs salt salt scripts scripts .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md environment.yml environment.yml eval.py eval.py mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml requirements.txt requirements.txt View all files Repository files navigation
Salience-Aware Lexical Trie for Long-Context Compression
MCP server - a salt-mcp entry point so AI clients can use SALT as
their conversation memory without the REPL.
Tail-aware memory selection - stop spending the memory budget on text
the model is already reading in the recent messages.
Incremental compression - reuse the previous turn's work instead of
rebuilding the whole selection every turn.
Graduating the memory switches - decide which off-by-default memory
behaviors become defaults, using numbers from real sessions.
Scripted conversation runs - richer tooling around --turns for
driving and scoring long canned conversations.
SALT shrinks a long document down to a fixed size before it is sent to a language
model, keeping the sentences that carry the most information. It works with any
model, produces a shorter plain-text prompt, and cuts the compute, memory, and
wait time that long inputs cost.
The problem. When a prompt is too long, existing compressors give each
sentence a single relevance score and keep the top-scoring ones until the budget
runs out. Under a tight budget this lets the document's main topic swallow the
whole budget, so smaller but still important points get dropped - a failure called
theme collapse (in multi-hop questions, for example, it can keep
passages about the main entity yet lose the one sentence that links it to a
second).
The solution. SALT first maps the document's recurring themes by organizing
each sentence's keywords into a trie, a small keyword tree ordered by how often
those keywords recur, then spreads the budget across those theme branches
before choosing sentences, so minor themes keep their share instead of being
crowded out. Because the theme map is built once, it can be reused across the
turns of a conversation without re-reading the document.
The previous legacy selector as described in the paper release is tagged v1.0.0 ;
main now defaults to the coverage/CELF selector described below.
Two phases. Indexing reads the document once and builds a keyword trie - a
reusable map of its recurring themes. Selection then picks a sentence subset
under a token budget, returned in original document order, query-biased or not.
INDEXING ── once per document, reused across turns and budgets
document → split + junk filter
→ per-sentence keywords (BGE-small [CLS] attention + knee cutoff)
→ theme salience (SF = #sentences keeping a word, top quantile)
→ keyword trie (each sentence's themes, SF-ordered, form a
root-to-leaf path, leaves hold sentence ids)
SELECTION ── per budget, with or without a query
maximize theme coverage with CELF lazy-greedy: a pick's value shrinks as its
theme branches fill, so budget spreads across themes instead of collapsing
onto the dominant one. A query re-weights the trie (lexical + BGE-semantic)
without rebuilding it. → compressed prompt, original order, ≤ budget
The whole system as a blueprint - this map is kept current as SALT grows,
so it is the fastest way to find where a change belongs:
┌──────────────────────────────────────────────────────────────────────────┐
│ SALT │
│ │
│ ┌────────────────────┐ ┌────────────────────┐ ┌────────────────────┐ │
│ │ Indexing │ │ Keyword Trie │ │ Selection │ │
│ │ │ │ │ │ │ │
│ │ BGE-small encoder │ │ SF-ordered paths │ │ coverage (CELF) │ │
│ │ attention keywords │ │ theme branches │ │ branch discounting │ │
│ │ knee cutoff │ │ §file: doc branches│ │ multi-anchor query │ │
│ │ junk filter │ │ rebuilt cheaply │ │ ≤ word budget │ │
│ └────────────────────┘ └────────────────────┘ └────────────────────┘ │
│ │
│ ┌────────────────────┐ ┌────────────────────┐ ┌────────────────────┐ │
│ │ Session Trie │ │ Prompt Assembly │ │ Chat Runner │ │
│ │ │ │ │ │ │ │
│ │ per-conversation │ │ stable prefix first│ │ HF streaming │ │
│ │ lives in DRAM │ │ append-only tail │ │ vLLM + APC (opt-in)│ │
│ │ grows every turn │ │ memory + question │ │ vllm-serve client │ │
│ │ cross-turn coverage│ │ instructions.md │ │ model registry │ │
│ │ + half-life decay │ │ │ │ GPU-pinned models │ │
│ │ + near-dup gate │ │ │ │ │ │
│ │ + background ingest│ │ │ │ │ │
│ └────────────────────┘ └────────────────────┘ └────────────────────┘ │
│ │
│ ┌────────────────────────────────────────────────────────────────────┐ │
│ │ Document ingest (salt@ files, salt --doc) │ │
│ │ pypdf extract · furniture scrub · paragraphs rejoined across floats│ │
│ │ tables + pseudocode grouped under captions · footnotes isolated │ │
│ │ headings, panel labels and equations kept · reference list dropped │ │
│ └────────────────────────────────────────────────────────────────────┘ │
│ │
│ ┌────────────────────────────────────────────────────────────────────┐ │
│ │ Trie shape - the root binds the conversation │ │
│ │ ● root - the conversation bind │ │
│ │ ┌─────────────────────┼─────────────────────┐ │ │
│ │ §file:paper.pdf §file:notes.txt conversation │ │
│ │ │ │ ┌──────┴──────┐ │ │
│ │ keyword paths keyword paths theme A theme B │ │
│ │ │ │ │ │ │ │
│ │ sentences sentences sentence
[truncated]
Requires Python 3.10 and a CUDA GPU (CPU works for compression, just slower).
git clone https://github.com/oteomamo/SALT.git
cd SALT
2. Create the environment
conda env create -f environment.yml
conda activate salt
Or with venv:
python3.10 -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
3. Install SALT in editable mode
pip install -e .
This also installs the two console commands: salt (one-shot compression,
see Usage ) and saltChat (interactive chat, see
Chatbot mode ).
4. Authenticate with Hugging Face - the eval model
( meta-llama/Llama-3.1-8B-Instruct ) is gated:
hf auth login
Or skip the CLI and export the token directly: export HF_TOKEN=hf_... .
5. (Optional) vLLM backend. eval.py defaults to vLLM,
saltChat --backend vllm uses it for prefix caching, and saltServe
launches a persistent model server with it. Install it into the salt
env:
pip install " vllm==0.11.0 " " prometheus-fastapi-instrumentator>=8.0.1 "
The second pin keeps the server's routes healthy next to newer fastapi
releases. Skip this and run eval.py --backend hf for a portable run
that needs no vLLM. saltChat already defaults to its HF backend.
saltServe can also run a vLLM installed in a separate environment
through --vllm-bin .
bash scripts/setup_env.sh does steps 2–3 in one shot (add WITH_VLLM=1 to
include vLLM).
Fetch the LongBench data, then compress every task present and evaluate the
outputs in one command:
python salt/datasets/download_longbench.py
bash scripts/run_datasets.sh
Smoke test (5 samples per task, compression only):
MAX_SAMPLES=5 RUN_EVAL=0 bash scripts/run_datasets.sh
The script routes each dataset to the right mode automatically, then scores
the results. Every command, flag, and knob is documented on the
Usage and
Datasets pages.
saltChat is an interactive chat REPL where SALT is the conversation memory:
one persistent trie per conversation grows with every exchange (and any
attached documents), and each turn it compresses the accumulated history into
a query-biased context block under the token budget.
saltChat --model qwen05 --conversation-id demo1 --doc report.txt
A persistent server started with saltServe keeps the model loaded and
its cache warm between chats, so a resumed conversation picks up without
re-reading its documents. The
Chatbot mode guide covers
the concepts, the Serving
page covers the server, and the
Options page lists every
flag in one line each, including the off-by-default switches that make
long sessions better.
SALT (coverage/CELF selector) reaches an overall 44.60 LongBench average
with Llama-3.1-8B-Instruct at a 20% token budget. The full per-dataset table
is on the Results page.
MCP server - a salt-mcp entry point exposing compression and session
memory as tools, so AI clients (Claude Code, Claude Desktop, Cursor) can use
SALT as their conversation memory without the REPL.
Tail-aware memory selection - skip sentences the model is already
reading verbatim in the recent messages, so the memory budget buys new
context instead of repeating what is on screen.
Incremental compression - carry the previous turn's selection work
forward on an append-only conversation, instead of redoing all of it every
turn.
Graduating the memory switches - several memory behaviors ship off
by default (see the
Options page) while
/stats numbers from real sessions decide which become defaults.
Scripted conversation runs - richer tooling around --turns , so
canned conversations can drive long sessions and be scored afterward.
Summarization coverage - extend the theme-coverage objective to better
serve summarization, where recall across many minor themes matters most.
PRs welcome - see CONTRIBUTING.md .
If you find this project useful for your research, please consider citing our paper：
@misc{mamo2026saltsalienceawarelexicaltrie,
title={SALT: Salience-Aware Lexical Trie for Long-Context Compression},
author={Oteo Mamo and Hyunjin Yi and Joydhriti Choudhury and Shangqian Gao and Weikuan Yu},
year={2026},
eprint={arXiv:2607.17486}
}
📄 License
SALT is released under the MIT License .
Salience-aware lexical trie for long-context compression.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
saltChat
Latest
Jul 4, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
