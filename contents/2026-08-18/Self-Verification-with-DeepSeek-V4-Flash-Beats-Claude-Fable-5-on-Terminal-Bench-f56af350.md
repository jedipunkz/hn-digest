---
source: "https://github.com/llm-as-a-verifier/llm-as-a-verifier"
hn_url: "https://news.ycombinator.com/item?id=49348195"
title: "Self-Verification with DeepSeek V4 Flash Beats Claude Fable 5 on Terminal-Bench"
article_title: "GitHub - llm-as-a-verifier/llm-as-a-verifier: LLM-as-a-Verifier is a general-purpose framework that provides fine-grained feedback for any agent without requiring additional training. It achieves SOTA performance across coding, robotics, and medical agentic benchmarks. · GitHub"
image: "https://opengraph.githubassets.com/2dccb065c8f919f4e036826cc0197a49f4c9927e94b3ee4065dd1f5143491931/llm-as-a-verifier/llm-as-a-verifier"
author: "yogthos"
captured_at: "2026-08-18T17:19:43Z"
capture_tool: "hn-digest"
hn_id: 49348195
score: 2
comments: 0
posted_at: "2026-08-18T16:30:33Z"
tags:
  - hacker-news
  - translated
---

# Self-Verification with DeepSeek V4 Flash Beats Claude Fable 5 on Terminal-Bench

- HN: [49348195](https://news.ycombinator.com/item?id=49348195)
- Source: [github.com](https://github.com/llm-as-a-verifier/llm-as-a-verifier)
- Score: 2
- Comments: 0
- Posted: 2026-08-18T16:30:33Z

## Translation

タイトル: DeepSeek V4 Flash による自己検証がターミナルベンチでクロード Fable 5 を上回る
記事タイトル: GitHub - llm-as-a-verifier/llm-as-a-verifier: LLM-as-a-Verifier は、追加のトレーニングを必要とせずに、あらゆるエージェントにきめ細かいフィードバックを提供する汎用フレームワークです。コーディング、ロボット工学、医療エージェントのベンチマーク全体で SOTA パフォーマンスを実現します。 · GitHub
説明: LLM-as-a-Verifier は、追加のトレーニングを必要とせずに、あらゆるエージェントにきめ細かいフィードバックを提供する汎用フレームワークです。コーディング、ロボット工学、医療エージェントのベンチマーク全体で SOTA パフォーマンスを実現します。 - 検証者としての llm/検証者としての llm

記事本文:
GitHub - llm-as-a-verifier/llm-as-a-verifier: LLM-as-a-Verifier は、追加のトレーニングを必要とせずに、あらゆるエージェントにきめ細かいフィードバックを提供する汎用フレームワークです。コーディング、ロボット工学、医療エージェントのベンチマーク全体で SOTA パフォーマンスを実現します。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
検証者としての llm
/
検証者としての llm
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メイン ブランチ タグ 移動先

ファイル コード [その他のアクション] メニューを開く 最新のコミット
13 コミット 13 コミット フォルダーとファイル
アセット アセット 基準 基準 データ データ llm_verifier llm_verifier スクリプト スクリプト .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md ライセンス ライセンス README.md README.md add_new_benchmark.md add_new_benchmark.md pyproject.toml pyproject.toml要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
あらゆるモダリティ、多くのアプリケーション、1 つの統合検証フレームワーク
|ドキュメント |ウェブサイト |紙 |クロードコードプラグイン |ツイッター/X |スラック |
🔥 LLM-as-a-Verifier は、ターミナルベンチ、SWE-Bench Verified、MedAgentBench、RoboRewardBench などのエージェント ベンチマーク全体で SOTA パフォーマンスを実現します。コミュニティがさらに多くのユースケースに貢献することを歓迎します。
pip インストール llm-verifier
クローンから最新のものをインストールするには:
pip install -e 。
0.2.0 の新機能 (詳細は CHANGELOG.md にあります):
プレフィックス キャッシュの最適化: キャッシュされていない入力トークンが最大 3.4 倍減少
弾道重視のベンチマーク
Terminal-Bench 2.1 自己検証ベンチマーク
deepseek-v4-flash ベリファイアー バックエンド
トークンアカウンティング ( llm_verifier.token_usage() )
LLM-as-a-Verifier は、きめの細かい検証機能を提供する汎用フレームワークです。
あらゆるエージェントへのフィードバック。重要なアイデアはシンプルです: 1) きめ細かいスコアリングを使用する
粒度、2) LLM の完全な logprob 分布に対する期待値を取得します。
スコアトークン、および 3) 反復評価と基準分解のスケール。の
結果として得られるきめの細かいフィードバックは、テスト時間のスケーリングや進捗状況に使用できます。
追跡と強化学習。
最初のエンドツーエンド選択を実行します ( .env の DEEPSEEK_API_KEY または VERTEX_API_KEY、または以下を返す OpenAI 互換サーバーが必要です)
logprobs — 例: vllm サーブ Qwen/Qwen3.5-9B 付き
OPENAI_BASE_URL=http://localhost:8000/v1 ):
インポート

llm_verifier
問題 = 「文字列を反転する関数を作成してください。」
候補者 = [
"def rev(s): return s[::-1]" 、 "def rev(s): return s" 、 "def rev(s): return ''.join(sorted(s))" 、
】
結果 = llm_verifier 。選択(
問題 = 問題、
候補者 = 候補者 、
criteria = { "正確性" : "コードは実際に文字列を反転しますか?" }、
)
print ( result .index ) # 最良の候補のインデックス: 0
print (結果 . スコア) # 候補スコア: [0.73104, 0.38446, 0.38449]
候補者のペアを直接採点する
select はペアごとの報酬モデルに基づいて構築されています。生のきめ細かい報酬については
単一の比較の場合は、compare を呼び出します。
報酬_a 、報酬_b = llm_verifier 。比較(
問題、候補 [ 0 ]、候補 [ 1 ]、
criteria = { "全体" : "コードは問題を解決しますか?" }、
)
print (reward_a , award_b ) # きめ細かい報酬 in [0, 1]: 0.99994 0
きめ細かい進捗状況の追跡
同じきめ細かい報酬により、毎回の終了後にエージェントの進捗状況をスコアリングすることもできます。
トラック付きステップ:
ステップ = [
'問題文を読んでください' ,
' def rev(s) を書き込みました: return s ' 、
'テスト済み: rev("abc") は "abc" を返しました' ,
'def rev(s) に変更: return s[::-1]' ,
'テスト済み: rev("abc") は "cba" を返しました' ,
】
結果 = llm_verifier 。トラック (問題 = 問題、ステップ = ステップ、
チェックポイントステップ = [ 1 、 2 、 3 、 4 、 5 ]、 n_evaluations = 4 )
print (結果 . スコア) # 各ステップ後の進行状況: [0.00106, 0.02417, 0.03143, 0.62004, 0.99978]
自己検証 (ターミナルベンチ 2.1)
モデルは自身のロールアウトを検証できますか? Terminal-Bench 2.1 では 5 を生成します
deepseek-v4-flash を使用してタスクごとの mini-swe-agent 軌跡を作成し、
検証者と同じモデル。選択範囲は Pass@1 よりはるかに上にありますが、
検証者は自身のモデルの動作を判断します。
軌跡は data/terminal_bench_2.1_trajs/ に保存されます。スコアリングが必要なだけ
DEEPSEEK_API_KEY

.env にあります。各構成には独自の再現があります
スクリプト:
python scripts/run_bo3.py # ベストオブ 3
python scripts/run_bo5.py # ベストオブ 5
エージェントベンチマークのテスト時間のスケーリング
各ベンチマークには、エージェントの軌跡 ( data/ ) が付属しています。 Gemini 2.5を使用しています
以下のすべてのベンチマークの検証ツールとして Flash ( gemini-2.5-flash ) を使用します。期待される
結果:
ベンチマークを名前で実行します (引数なしで python scripts/run.py を実行するとベンチマークがリストされます)。
Python スクリプト/run.py ターミナルベンチ
Python スクリプト/run.py swe_bench
Python scripts/run.py medagentbench
トーナメントのデフォルトはコマンドラインで上書きできます。
python scripts/run.py swe_bench --pivots 2 --n-evaluations 8 --seed 0 --max-workers 50
ベンチマークは llm_verifier/benchmarks.py で定義されています。そこにベンチマークを追加または調整します。
N エージェントの軌跡のベストを選択
タスクとエージェントの軌跡のプールが与えられた場合、いくつかの中で最適なものを選択します
コード行。
llm_verifier をインポートする
問題 = "utils.py の失敗したテストを修正します。"
候補 = [ traj_1 、 traj_2 、 traj_3 、 traj_4 、 traj_5 ]
結果 = llm_verifier 。選択(
問題 = 問題、
候補者 = 候補者 、
criteria = { "根本原因" : "エージェントは本当の原因を修正しましたか?" 、
「検証」 : 「エージェントは修正を確認しましたか?」 }、
model = "gemini-2.5-flash" , # 検証者モデル
n_evaluations = 4 、基準ごとに繰り返される評価の数
ピボット = 2 、ピボット数 < N;検証コストの削減
)
print ( "最有力候補:" , result .index )
print ( "ランキング:" 、結果。ランキング)
ボンネットの下で、選択は以下を実行します
すべてをランク付けする確率的ピボット トーナメント
完全な検証の代わりに O(Nk) 個のペアごとの検証を使用する N 個の軌跡
O(N²) ラウンドロビン。ピボットは精度のためにコストを犠牲にします: ピボットが増える = より多く
比較 = 精度が高くなります。
LLM-as-a-Verifier を独自のユースケースに適応させる
3 つのステップで独自のタスクにベリファイアを使用する — Cla

ude コードが残りを実行します
(基準を生成し、ランナーを書き込み、N 個の中から最良のものを選択します):
データを追加します。エージェントの軌跡を data/task_name_trajs/ にコピーします。
ネーミングを更新します。すべての task_name を置き換えます
add_new_benchmark.md をタスクの名前に置き換えます。
このリポジトリ (または Codex、または好きなもの) で Claude Code をスピンアップします。
権限が無効になっています)、add_new_benchmark.md の内容を貼り付けて、
それは走ります。
コーディングエージェントの進捗状況の追跡
同じきめの細かい報酬によって、すべてのステップで軌道をスコアリングできます (「
クイックスタートでトラックしてください)。以下では、ターミナル ベンチ タスク pytorch-model-cli の 2 つの Terminus-2 実行を追跡します。成功した軌跡は検証者スコアの一貫した増加を示しますが、失敗した軌跡は誤った動作によって特徴付けられ、実行全体を通じてスコアが低くなります。次のようにして再現します。
python scripts/terminal_bench_progress.py # 実行とプロットの両方のスコアを取得します
トラックは完成した軌道をスコアリングします。エージェントを監視するには
実行する場合は、ProgressTracker を使用します。各ステップが発生するたびにフィードし、ライブ情報を取得します。
進行状況スコアバック — 例:絶望的な展開を早期に中止するか、いつ中止するかを決定する
リサンプル。検証者はこれまでのステップのみを確認するため、覗くことはできません。
将来的には。
トラッカー = llm_verifier 。 ProgressTracker (問題、n_evaluations = 4)
スコア = トラッカー。 update ( '問題文を読む' ) # 0.00002
スコア = トラッカー。 update ( 'Write def rev(s): return s' ) # 0.00013
スコア = トラッカー。 update ( 'def rev(s) に変更: return s[::-1]' ) # 0.73938
スコア = トラッカー。 update ( 'テスト済み: rev("abc") が "cba" を返しました' ) # 0.98604
if スコア < 0.05 : # 任意のステップの後: 絶望的なロールアウトを早期に放棄する
...
2 つのターミナルとベンチの軌跡を段階的に再生します。
ProgressTracker — 毎回の後にライブスコアバーを印刷する

ステップ、エージェントとして
ハーネスはそれを見るでしょう:
python scripts/terminal_bench_progress.py --online
マルチモーダルのサポート
マルチモーダル検証モデル (例: Gemini 2.5 Flash または
vllm サーブ Qwen/Qwen3.5-9B )、すべて
エントリ ポイントは画像を受け入れます。単一の画像 (images="frame.png" ) または
画像のリスト。それぞれローカル ファイル パス、http(s) URL、または生のバイトです。
結果 = llm_verifier 。 select (問題 , 候補 , criteria = 基準 ,
画像 = [ "before.png" , "after.png" ])
トラッカー = llm_verifier 。 ProgressTracker (問題)
スコア = トラッカー。 update ( step , Images = "camera_frame.png" ) # ステップごとのフレーム
ステップごとのフレームは、その後のすべての更新の軌跡の一部として残ります。
検証者は常に完全なビジュアル履歴を参照します。カメラのフレーム
ロボットの展開を追跡します。を参照してください。
承認されたマルチモーダル文書
入力フォーム、バックエンドノート、検証済みの例。
TurboAgent がもたらす
LLM-as-a-Verifier から Claude Code へのドロップイン
LLM API プロキシ。クライアントとモデルプロバイダーの間に位置し、
複数の回答候補を並行して実行し、最適な回答を選択します。
確率的ピボット トーナメント。
pip install git+https://github.com/llm-as-a-verifier/TurboAgent
クロード コードをプロキシに指定し、通常どおり実行します。
ターボエージェント # はポート 8888 で開始します
ANTHROPIC_BASE_URL=http://localhost:8888 クロード
組み込みのビジュアライザを次の場所に同梱しています。
パイプライン DAG、進捗スコア、候補を表示する http://localhost:8888/visualizer
回答と最終的な選択。を参照してください。
TurboAgent リポジトリ
構成とセットアップの詳細。
。
§── scripts/ # コマンドライン エントリ ポイント
│ §── run.py # レジストリ主導のベンチマーク ランチャー
│ §── run_bo3.py # ベストオブ 3 の自己検証実行を再現します。
│ §── run_bo5.py # ベストオブ 5 の自己検証の実行を再現します。

│ lux──terminal_bench_progress.py # スコアを再取得し、進行状況追跡の例をプロットする
§── criteria/# 検証者の基準 + グラウンドトゥルースのメモ
│ §── TEMPLATE.md # これをコピーして独自に作成します
│ §── ターミナルベンチ.md
│ §── swe_bench.md
│ └── medagentbench.md
§── llm_verifier/ # 再利用可能なフレームワーク (llm_verifier をインポート)
│ §── __init__.py # llm_verifier.select(...) / .compare(...)
│ §── __main__.py # python -m llm_verifier <file.md>: プレビュー基準
│ §── benchmarks.py # BENCHMARKS レジストリ (1 つのベンチマーク/起動)
│ §──fine_graned_reward.py # R(x,τ): logprob スコアリング + スコアキャッシュ
│ §── progress.py # llm_verifier.track(...): ステップごとの進捗曲線
│ §── pivot_tournament.py # PPT: O(Nk) セレクション (Bradley-Terry)
│ §── Prompts.py # ロード条件/*.md + 条件の引数を正規化
│ └──loaders.py # ベンチマークごとの軌跡ローダー
└── データ/ベンチマークごとのエージェントの軌跡の数
実行は検証者スコア キャッシュをキャッシュに書き込み、結果テーブルをキャッシュに書き込みます。
結果/ ;どちらもオンデマンドで作成され、git は無視されます。
きめ細かい報酬の見積もり
各分布を単一の離散スコアに縮小するのではなく (次のように)
LLM-as-a-Judge)、LLM-as-a-Verifier は軌跡の報酬を近似します
タスク $x$ の $\tau$ は次のようになります。
$$
R(x, \t)
= \frac{1}{CK} \sum_{c=1}^{C} \sum_{k=1}^{K}
\

[切り捨てられた]

## Original Extract

LLM-as-a-Verifier is a general-purpose framework that provides fine-grained feedback for any agent without requiring additional training. It achieves SOTA performance across coding, robotics, and medical agentic benchmarks. - llm-as-a-verifier/llm-as-a-verifier

GitHub - llm-as-a-verifier/llm-as-a-verifier: LLM-as-a-Verifier is a general-purpose framework that provides fine-grained feedback for any agent without requiring additional training. It achieves SOTA performance across coding, robotics, and medical agentic benchmarks. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
llm-as-a-verifier
/
llm-as-a-verifier
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
13 Commits 13 Commits Folders and files
assets assets criteria criteria data data llm_verifier llm_verifier scripts scripts .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md add_new_benchmark.md add_new_benchmark.md pyproject.toml pyproject.toml requirements.txt requirements.txt View all files Repository files navigation
Any modality, Many Applications, One Unified Verification Framework
| Documentation | Website | Paper | Claude Code Plugin | Twitter/X | Slack |
🔥 LLM-as-a-Verifier achieves SOTA performance across agentic benchmarks, including Terminal-Bench, SWE-Bench Verified, MedAgentBench, RoboRewardBench and more. We invite the community to contribute more use cases!
pip install llm-verifier
To install the latest from a clone:
pip install -e .
What's new in 0.2.0 (full notes in CHANGELOG.md ):
Prefix-cache optimization: ~3.4× fewer uncached input tokens on
trajectory-heavy benchmarks
Terminal-Bench 2.1 self-verification benchmark
deepseek-v4-flash verifier backend
Token accounting ( llm_verifier.token_usage() )
LLM-as-a-Verifier is a general-purpose framework that provides fine-grained
feedback for any agent. The key idea is simple: 1) use fine-grained scoring
granularity, 2) take the expectation over the full logprob distribution of LLM
score tokens, and 3) scale repeated evaluation and criteria decomposition. The
resulting fine-grained feedback can be used for test-time scaling, progress
tracking, and reinforcement learning.
Run a first end-to-end selection (requires DEEPSEEK_API_KEY or VERTEX_API_KEY in .env , or an OpenAI-compatible server that returns
logprobs — e.g. vllm serve Qwen/Qwen3.5-9B with
OPENAI_BASE_URL=http://localhost:8000/v1 ):
import llm_verifier
problem = "Write a function that reverses a string."
candidates = [
"def rev(s): return s[::-1]" , "def rev(s): return s" , "def rev(s): return ''.join(sorted(s))" ,
]
result = llm_verifier . select (
problem = problem ,
candidates = candidates ,
criteria = { "Correctness" : "Does the code actually reverse the string?" },
)
print ( result . index ) # index of the best candidate: 0
print ( result . scores ) # candidate scores: [0.73104, 0.38446, 0.38449]
Score a pair of candidates directly
select is built on a pairwise reward model. For the raw fine-grained rewards
of a single comparison, call compare :
reward_a , reward_b = llm_verifier . compare (
problem , candidates [ 0 ], candidates [ 1 ],
criteria = { "Overall" : "Does the code solve the problem?" },
)
print ( reward_a , reward_b ) # fine-grained rewards in [0, 1]: 0.99994 0
Fine-grained Progress Tracking
The same fine-grained reward can also score an agent's progress after each
step with track :
steps = [
'Read the problem statement' ,
'Wrote def rev(s): return s ' ,
'Tested: rev("abc") returned "abc"' ,
'Changed to def rev(s): return s[::-1]' ,
'Tested: rev("abc") returned "cba"' ,
]
result = llm_verifier . track ( problem = problem , steps = steps ,
checkpoint_steps = [ 1 , 2 , 3 , 4 , 5 ], n_evaluations = 4 )
print ( result . scores ) # progress after each step: [0.00106, 0.02417, 0.03143, 0.62004, 0.99978]
Self-Verification (Terminal Bench 2.1)
Can a model verify its own rollouts? On Terminal-Bench 2.1 we generate 5
mini-swe-agent trajectories per task with deepseek-v4-flash and use the
same model as the verifier. Selection lands well above Pass@1 even though
the verifier is judging its own model's work:
The trajectories ship in data/terminal_bench_2.1_trajs/ ; scoring only needs
DEEPSEEK_API_KEY in .env . Each configuration has its own reproduction
script:
python scripts/run_bo3.py # best-of-3
python scripts/run_bo5.py # best-of-5
Test-Time Scaling for Agentic Benchmarks
Each benchmark ships with its agent trajectories ( data/ ). We use Gemini 2.5
Flash ( gemini-2.5-flash ) as the verifier for all benchmarks below. Expected
results:
Run a benchmark by name ( python scripts/run.py with no argument lists them):
python scripts/run.py terminal_bench
python scripts/run.py swe_bench
python scripts/run.py medagentbench
The tournament defaults can be overridden on the command line:
python scripts/run.py swe_bench --pivots 2 --n-evaluations 8 --seed 0 --max-workers 50
Benchmarks are defined in llm_verifier/benchmarks.py — add or tweak one there.
Select Best of N agent trajectories
Given a task and a pool of agent trajectories, pick the best one in a few
lines of code.
import llm_verifier
problem = "Fix the failing test in utils.py."
candidates = [ traj_1 , traj_2 , traj_3 , traj_4 , traj_5 ]
result = llm_verifier . select (
problem = problem ,
candidates = candidates ,
criteria = { "Root cause" : "Did the agent fix the real cause?" ,
"Verification" : "Did the agent confirm the fix?" },
model = "gemini-2.5-flash" , # verifier model
n_evaluations = 4 , # repeated evaluations per criterion
pivots = 2 , # pivots < N; reduced verification cost
)
print ( "Best candidate:" , result . index )
print ( "Ranking:" , result . ranking )
Under the hood, select runs the
Probabilistic Pivot Tournament to rank all
N trajectories using O(Nk) pairwise verifications instead of a full
O(N²) round-robin. pivots trades cost for accuracy: more pivots = more
comparisons = higher accuracy.
Adapt LLM-as-a-Verifier for your own use case
Use the verifier for your own task in three steps — Claude Code does the rest
(generates the criteria, writes a runner, and selects the best-of-N for you):
Add your data. Copy your agent trajectories into data/task_name_trajs/ .
Update naming. Replace every task_name in
add_new_benchmark.md with the name of your task.
Spin up Claude Code in this repo (or Codex, or whatever you like — with
permissions disabled) and paste the contents of add_new_benchmark.md to let
it run.
Progress Tracking for Coding Agents
The same fine-grained reward can score a trajectory at every step (see
track in the Quickstart ). Below, we track two Terminus-2 runs of the Terminal-Bench task pytorch-model-cli . The successful trajectory exhibits consistently increasing verifier scores, whereas the failed trajectory is characterized by erroneous behaviors, resulting in lower scores throughout the execution. Reproduce it with:
python scripts/terminal_bench_progress.py # scores both runs then plots
track scores a finished trajectory. To monitor an agent while it
runs , use ProgressTracker : feed it each step as it happens and get a live
progress score back — e.g. to stop a hopeless rollout early or decide when to
resample. Since the verifier only ever sees the steps so far, it cannot peek
at the future.
tracker = llm_verifier . ProgressTracker ( problem , n_evaluations = 4 )
score = tracker . update ( 'Read the problem statement' ) # 0.00002
score = tracker . update ( 'Wrote def rev(s): return s' ) # 0.00013
score = tracker . update ( 'Changed to def rev(s): return s[::-1]' ) # 0.73938
score = tracker . update ( 'Tested: rev("abc") returned "cba"' ) # 0.98604
if score < 0.05 : # after any step: abandon a hopeless rollout early
...
Replay the two Terminal-Bench trajectories step-by-step through
ProgressTracker — printing a live score bar after every step, as an agent
harness would see it:
python scripts/terminal_bench_progress.py --online
Multi-Modal Support
With a multimodal verifier model (e.g. Gemini 2.5 Flash or
vllm serve Qwen/Qwen3.5-9B ), every
entry point accepts images — a single image ( images="frame.png" ) or a
list of images, each a local file path, an http(s) URL, or raw bytes:
result = llm_verifier . select ( problem , candidates , criteria = criteria ,
images = [ "before.png" , "after.png" ])
tracker = llm_verifier . ProgressTracker ( problem )
score = tracker . update ( step , images = "camera_frame.png" ) # per-step frame
Per-step frames stay part of the trajectory for all later updates, so the
verifier always sees the full visual history — e.g. camera frames while
tracking a robot rollout. See the
multimodal documentation for accepted
input forms, backend notes, and verified examples.
TurboAgent brings
LLM-as-a-Verifier to Claude Code as a drop-in
LLM API proxy. It sits between your client and the model provider, generating
multiple candidate responses in parallel and selecting the best one with a
Probabilistic Pivot Tournament .
pip install git+https://github.com/llm-as-a-verifier/TurboAgent
Point Claude Code at the proxy and run as usual:
turbo-agent # starts on port 8888
ANTHROPIC_BASE_URL=http://localhost:8888 claude
It ships a built-in visualizer at
http://localhost:8888/visualizer that shows the pipeline DAG, progress scores, candidate
responses, and the final selection. See the
TurboAgent repository for
configuration and setup details.
.
├── scripts/ # command-line entry points
│ ├── run.py # registry-driven benchmark launcher
│ ├── run_bo3.py # reproduce the best-of-3 self-verification run
│ ├── run_bo5.py # reproduce the best-of-5 self-verification run
│ └── terminal_bench_progress.py # re-score + plot the progress-tracking example
├── criteria/ # verifier criteria + ground-truth notes
│ ├── TEMPLATE.md # copy this to write your own
│ ├── terminal_bench.md
│ ├── swe_bench.md
│ └── medagentbench.md
├── llm_verifier/ # the reusable framework (import llm_verifier)
│ ├── __init__.py # llm_verifier.select(...) / .compare(...)
│ ├── __main__.py # python -m llm_verifier <file.md>: preview criteria
│ ├── benchmarks.py # BENCHMARKS registry (one Benchmark / launch)
│ ├── fine_grained_reward.py # R(x,τ): logprob scoring + score cache
│ ├── progress.py # llm_verifier.track(...): per-step progress curve
│ ├── pivot_tournament.py # PPT: O(Nk) selection (Bradley-Terry)
│ ├── prompts.py # load criteria/*.md + normalize criteria args
│ └── loaders.py # per-benchmark trajectory loaders
└── data/ # agent trajectories per benchmark
Runs write their verifier score caches to cache/ and result tables to
results/ ; both are created on demand and git-ignored.
Fine-grained Reward Estimation
Rather than reducing each distribution into a single discrete score (as in
LLM-as-a-Judge), LLM-as-a-Verifier approximates the reward of a trajectory
$\tau$ on task $x$ as:
$$
R(x, \tau)
= \frac{1}{CK} \sum_{c=1}^{C} \sum_{k=1}^{K}
\

[truncated]
