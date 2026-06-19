---
source: "https://github.com/raphaelwkago-sketch/rudi"
hn_url: "https://news.ycombinator.com/item?id=48598025"
title: "Causal graph memory for LLMs. Flat token cost, no matter how the session runs"
article_title: "GitHub - raphaelwkago-sketch/rudi: Causal graph memory for LLMs - flat token cost regardless of session length. · GitHub"
author: "w4mwati"
captured_at: "2026-06-19T13:18:03Z"
capture_tool: "hn-digest"
hn_id: 48598025
score: 1
comments: 0
posted_at: "2026-06-19T12:54:12Z"
tags:
  - hacker-news
  - translated
---

# Causal graph memory for LLMs. Flat token cost, no matter how the session runs

- HN: [48598025](https://news.ycombinator.com/item?id=48598025)
- Source: [github.com](https://github.com/raphaelwkago-sketch/rudi)
- Score: 1
- Comments: 0
- Posted: 2026-06-19T12:54:12Z

## Translation

タイトル: LLM の因果グラフ メモリ。セッションの実行方法に関係なく、フラットなトークンコスト
記事のタイトル: GitHub - raphaelwkago-sketch/rudi: LLM の因果グラフ メモリ - セッションの長さに関係なく、フラットなトークン コスト。 · GitHub
説明: LLM の因果グラフ メモリ - セッションの長さに関係なく、フラットなトークン コスト。 - raphaelwkago-スケッチ/ルディ

記事本文:
GitHub - raphaelwkago-sketch/rudi: LLM の因果グラフ メモリ - セッションの長さに関係なく、フラットなトークン コスト。 · GitHub
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
ラファエルカゴのスケッチ
/
ルディ
公共
通知
通知設定を変更するにはサインインする必要があります
追加ナビ

オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md benchmark_long_haiku.py benchmark_long_haiku.pyfold.pyfold.py rudi.py rudi.py store.py store.py すべてのファイルを表示 リポジトリ ファイル ナビゲーション
LLM の因果グラフ メモリ。セッションの実行時間がどれだけ長くても、トークンのコストは一律です。
LLM API 呼び出しごとに会話全体が再送信されます。コストはターンごとに増加します。最終的にはコンテキストの制限に達します。 Rudi は、増大するトランスクリプトを意思決定の依存関係グラフに置き換え、現在のタスクに関連するスライスのみを挿入します。 10,000ターン目のコストは10ターン目とほぼ同じです。
43 ターンのソフトウェア アーキテクチャ セッション (Notes API をターンごとに構築する) では、標準的な「完全なトランスクリプトを再送信する」アプローチでは、最終ターンまでに最大 38,000 個の入力トークンが送信されました。 Rudi は、同じタスク、同じモデル、同じ回答品質に対して 6,782 件を送信しました。
43 ターンすべての合計: 入力トークン (Rudi) 152,222 個 (Rudi) vs 828,369 (完全なトランスクリプト) — トークンの数は 5.4 分の 1 で、トランスクリプトの曲線は線形であるのに対し、Rudi の曲線は境界があるため、ギャップはターンごとに拡大します。
これらの数値は、フォールドを無効にして実行した結果、つまりグラフのスライスのみで得られたものです。折り目の測定結果については、以下を参照してください。
Claude Haiku 4.5 での 43 ターンの実行全体のコスト: $0.34。
別の実行のターン 29 で、初めてフォールドが発射されました。
ターン 28: 入力 = 5,075 トークン アクティブなノード = 24
[フォールド] d1–d8 (8 ノード、20 ハード ルール) → スタブ d25
[フォールド] d9–d16 (8 ノード、20 ハード ルール) → スタブ d26
[フォールド] d17–d21 (5 ノード、16 ハード ルール) → スタブ d27
ターン 29: アクティブなノード = 6 (ドロップ 24 → 6)
ターン 30: 投入 = 2,865 トークン ← ターン 28 から 44% 減少
21 個のライブ ノードが 3 つのスタブに圧縮されました。 56 の厳格なルールがそのまま保存されています。入力t

オーケンはセッション中に自動的にほぼ半分になります。それがノコギリ波です。会話が長くなるにつれて、グラフは小さくなります。
小さいままであるだけではなく、正確なままです
モデルがルールを忘れてしまったら、安っぽいコンテキストは価値がありません。したがって、同じベンチマークはセッション後半に 6 つのコールバック トラップを仕掛け、数十ターン前に行われた決定がまだ尊重されているかどうかをチェックします。
6 / 6. (最初のベンチマーク実行 - フォールドは無効、スライスのみ。) 最も重要な 2 つは #3 と #4 です。これらのルールは、トラップが発生するまでにアクティブなコンテキストから圧縮されていました。しかし、ハード ルールはフォールド スタブにそのまま保存されているため、モデルは依然としてそれらを捕らえていました。これがテーマ全体です。散文は忘れて、制約は維持してください。
すべてのモデル応答は意思決定ノードに解析され、それぞれが依存する意思決定に逆方向にリンクされます。
ノード = {
ID、テキスト、
depend_on: [...], # 後方エッジ — この決定の根拠
hard_rules: [...], # バインディング制約;違反した場合、労働者は停止しなければならない
Revis、Exception_to、# 完全な置換と狭いカーブアウトの比較
ステータス、ターン、固定
}
捨てずにスライスしてください。各ターンの前に、Rudi はトランスクリプトではなく、現在のタスクから到達可能なノードのみを挿入します。
折り畳み。決定の分岐が到達不能になると、バックグラウンド パスによってそれが 1 行のスタブに圧縮されます。ハード ルールは折りたたみ後もそのまま存続するため、制約が暗黙的に失われることはありません (トラップ #3/#4 を参照)。
ピン基礎。繰り返し強化された決定、最初の 2 ターンで行われた決定、またはキャリー例外は固定され、フォールドされることはありません。
厳しいルールには拘束力があります。新しいタスクがタスクに違反する場合、作業者は黙って従うのではなく、立ち止まって尋ねます (トラップ #5/#6)。
git clone https://github.com/ < あなた > /rudi
CDルディ
pip install anthropic flask flask-cors
# あなた自身のキー — 決してそうではありません

それをコード化してください
import ANTHROPIC_API_KEY= " sk-ant-... "
# ローカル SQLite に対して 43 ターンのベンチマークを実行します (クラウドは必要ありません)
python benchmark_long_haiku.py
単純なトランスクリプトが膨張する一方で、入力トークンの曲線が平坦なままであること、および 6 つのコールバック トラップがすべて解決されることを確認します。
1 ターンにつき 2 回の呼び出し。独自のモデルキーを保持します。 Rudi はグラフの管理のみを行います。
ルディをインポート
# 1 — LLM 呼び出しの前に、関連するスライスを取得します。
s = ルディ。 get_slice (タスク)
# → s["system"] + s["prompt"] を LLM 呼び出しにフィードします
#2 — LLM コール後: 決定内容を保存する
ルディ 。 store_decions (決定 , inject_ids = s [ "inject_ids" ])
または、Rudi にターン全体 (LLM コール + ストア + フォールド) を 1 回のショットでドライブさせます。
結果 = ルディ。 run_turn (タスク) # → {"display", "tokens_in", "tokens_out", ...}
ストレージはローカル SQLite (store.py) であり、デシジョン ノードごとに 1 行です。サーバーもクラウドもセットアップも必要ありません。
グラフのスライスがトークン曲線の境界を定める
✅ 測定値 — 上の表
40 ターン以上後に決定が呼び戻される
✅ 6/6 コールバック
厳格なルールはそのまま生き残る
✅ トラップ #3/#4
黙って従うのではなく、対立を阻止する
✅ トラップ #5/#6
Fold GC はセッション中にデッドブランチを圧縮します
✅ 測定結果 — 24 ノード → 6、ターン 30 で入力 -44%
約 80 個のアクティブ ノードを超える取得フォールバック
⏳ 構築済みだが、大規模なベンチマークはまだ行われていない
蒸気は出ません。表はログに記載されている内容です。進行中の行にはそのようにラベルが付けられます。
ビジネスソースライセンス 1.1.個人使用、研究、開発、セルフホスティングの場合は無料です。商用 SaaS またはマネージド ホスティングを使用するには、保守者からの有料ライセンスが必要です。
商用ライセンスが必要ですか?問題を開くか、 raphaelwkago@gmail.com に電子メールを送信してください。
AGPL の義務なしで Rudi を商業的に使用したいですか?問題を開くか、 raphaelwkago@gmail.com に電子メールを送信してください。
LLM の因果グラフ メモリ - セッションの長さに関係なく、フラットなトークン コスト

番目。
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

Causal graph memory for LLMs - flat token cost regardless of session length. - raphaelwkago-sketch/rudi

GitHub - raphaelwkago-sketch/rudi: Causal graph memory for LLMs - flat token cost regardless of session length. · GitHub
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
raphaelwkago-sketch
/
rudi
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md benchmark_long_haiku.py benchmark_long_haiku.py fold.py fold.py rudi.py rudi.py store.py store.py View all files Repository files navigation
Causal graph memory for LLMs. Flat token cost, no matter how long the session runs.
Every LLM API call re-sends the whole conversation. Cost grows every turn; eventually you hit the context limit. Rudi replaces the growing transcript with a dependency graph of decisions — and injects only the slice relevant to the current task. Turn 10,000 costs about the same as turn 10.
In a 43-turn software-architecture session (building a Notes API turn by turn), the standard "re-send the full transcript" approach was sending ~38,000 input tokens by the final turn. Rudi sent 6,782 — for the same task, same model, same answer quality.
Totals across all 43 turns: 152,222 input tokens (Rudi) vs 828,369 (full transcript) — 5.4× fewer tokens , and the gap widens every turn because Rudi's curve is bounded while the transcript's is linear.
These numbers are from a run with fold disabled — graph slicing alone. See below for the measured fold result.
Cost of the entire 43-turn run on Claude Haiku 4.5: $0.34.
At turn 29 of a separate run, fold fired for the first time:
turn 28: input=5,075 tokens active nodes=24
[fold] d1–d8 (8 nodes, 20 hard rules) → stub d25
[fold] d9–d16 (8 nodes, 20 hard rules) → stub d26
[fold] d17–d21 (5 nodes, 16 hard rules) → stub d27
turn 29: active nodes=6 (dropped 24 → 6)
turn 30: input=2,865 tokens ← down 44% from turn 28
21 live nodes compressed into 3 stubs. 56 hard rules preserved verbatim. Input tokens nearly halved mid-session, automatically. That's the sawtooth: the graph gets smaller as the conversation gets longer .
It doesn't just stay small — it stays correct
Cheap context is worthless if the model forgets the rules. So the same benchmark plants 6 callback traps late in the session and checks whether decisions made dozens of turns earlier are still honored.
6 / 6. (First benchmark run — fold disabled, slicing only.) The two that matter most are #3 and #4: those rules had been compressed out of the active context by the time the trap was sprung — and the model still caught them, because hard rules are preserved verbatim on the fold stub. That's the whole thesis: forget the prose, keep the constraints.
Every model response is parsed into decision nodes , each linked backward to the decisions it depends on:
node = {
id, text,
depends_on: [...], # backward edges — what this decision rests on
hard_rules: [...], # binding constraints; the worker must halt if violated
revises, exception_to, # full replacement vs. narrow carve-out
status, turn, pinned
}
Slice, don't dump. Before each turn, Rudi injects only the nodes reachable from the current task — not the transcript.
Fold. When a branch of decisions goes reachability-dead, a background pass compresses it into a one-line stub. Hard rules survive the fold verbatim , so a constraint can never be silently lost (see traps #3/#4).
Pin foundations. Decisions that are reinforced repeatedly, made in the first two turns, or carry exceptions are pinned and never folded.
Hard rules are binding. If a new task would violate one, the worker stops and asks instead of silently complying (traps #5/#6).
git clone https://github.com/ < you > /rudi
cd rudi
pip install anthropic flask flask-cors
# your own key — never hardcode it
export ANTHROPIC_API_KEY= " sk-ant-... "
# run the 43-turn benchmark against local SQLite (no cloud needed)
python benchmark_long_haiku.py
You'll watch the input-token curve stay flat while a naive transcript would balloon, and see all 6 callback traps resolve.
Two calls per turn. You keep your own model key; Rudi only manages the graph.
import rudi
# 1 — before your LLM call: fetch the relevant slice
s = rudi . get_slice ( task )
# → feed s["system"] + s["prompt"] into YOUR LLM call
# 2 — after your LLM call: store what was decided
rudi . store_decisions ( decisions , inject_ids = s [ "inject_ids" ])
Or let Rudi drive the whole turn (LLM call + store + fold) in one shot:
result = rudi . run_turn ( task ) # → {"display", "tokens_in", "tokens_out", ...}
Storage is local SQLite ( store.py ) — one row per decision node. No server, no cloud, no setup.
Graph slicing bounds the token curve
✅ measured — table above
Decisions recalled 40+ turns later
✅ 6/6 callbacks
Hard rules survive fold verbatim
✅ traps #3/#4
Conflicts blocked, not silently obeyed
✅ traps #5/#6
Fold GC compresses dead branches mid-session
✅ measured — 24 nodes → 6, input −44% at turn 30
Retrieval fallback above ~80 active nodes
⏳ built, not yet benchmarked at scale
No vapor. The table is what the logs say; the in-progress rows are labeled as such.
Business Source License 1.1. Free for personal use, research, development, and self-hosting. Commercial SaaS or managed hosting use requires a paid license from the maintainer.
Want a commercial license? Open an issue or email raphaelwkago@gmail.com .
Want to use Rudi commercially without AGPL obligations? Open an issue or email raphaelwkago@gmail.com .
Causal graph memory for LLMs - flat token cost regardless of session length.
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
