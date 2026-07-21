---
source: "https://github.com/Champoleello/agentpause"
hn_url: "https://news.ycombinator.com/item?id=48993421"
title: "Agentpause: Suspends LLM agents before rate limits, resumes cleanly"
article_title: "GitHub - Champoleello/agentpause: Suspend LLM agents gracefully before rate limits; resume without losing work · GitHub"
author: "Champoleello"
captured_at: "2026-07-21T16:18:15Z"
capture_tool: "hn-digest"
hn_id: 48993421
score: 1
comments: 0
posted_at: "2026-07-21T15:17:42Z"
tags:
  - hacker-news
  - translated
---

# Agentpause: Suspends LLM agents before rate limits, resumes cleanly

- HN: [48993421](https://news.ycombinator.com/item?id=48993421)
- Source: [github.com](https://github.com/Champoleello/agentpause)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T15:17:42Z

## Translation

タイトル: Agentpause: レート制限の前に LLM エージェントを一時停止し、正常に再開します
記事のタイトル: GitHub - Champoleello/agentpause: レート制限の前に LLM エージェントを適切に一時停止します。仕事を失わずに再開 · GitHub
説明: レート制限の前に LLM エージェントを正常に一時停止します。仕事を失わずに再開 - Champoleello/agentpause

記事本文:
GitHub - Champoleello/agentpause: レート制限の前に LLM エージェントを正常に一時停止します。仕事を失わずに再開 · GitHub
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
シャンポレロ
/
エージェント一時停止
公共
通知
変更するにはサインインする必要があります

化設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
25 コミット 25 コミット .github/ workflows .github/ workflows サンプル サンプル スクリプト スクリプト src/ エージェント一時停止 src/ エージェント一時停止 テスト テスト .gitignore .gitignore CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
自律型 LLM エージェントの予測スケジューリング。エージェントを正常に停止する
プロバイダーのレート制限に達する前に、作業をやり直すことなく再開します。
このコアは、クラウド (OpenAI、Anthropic、Groq) またはローカルのどのプロバイダーでも動作します。
アプリケーションレベルの状態をシリアル化するだけです。真の KV キャッシュ ウォーム スタートはオプションです
セルフホスト ランタイム用のプラグイン (llama.cpp、vLLM)。予測停止自体はまだ
レート制限ヘッダーであるクラウド プロバイダーに基づいて、実際の予算シグナルが必要です。
自己ホスト型サーバーの場合は何もありません。「ローカル モード: 実際に何が行われているか」を参照してください。
代わりに何を配線するかについては、以下の「controlled」を参照してください。
測定結果（付随研究より）
実験
結果
シミュレーション、300 実行/構成
反応的なベースラインは最大 100% クラッシュ → 予測 0% (k=4)、トークンの無駄 -80%
本物のプロバイダー (Groq 無料、6K TPM)
マルチウィンドウタスクは 429 エラーなしで完了しました
vs LangGraph の MemorySaver (200 実行)
LangGraph: 7.8 × 429/実行、9,239 トークンが無駄になりました。予測: 0 と 0
エンドツーエンド A/B (思考モデル、4B/8B)
回復が 54 ～ 93 倍速くなります。合計タスク時間 -19%
無料のキー: python scripts/benchmark_groq.py を使用して、これらを自分で再現します。
ライブ実行 (2026 年 7 月 8 日、Groq 無料枠、12 ステップ、~1k トークン コンテキスト、
リフィル対応のチャンク待機、レベル化されたウィンドウ):
壁時計と同等でエラーも無駄もゼロ: センサーのコストは最大 2%

の
クラッシュするのは無駄です (259 対 12,840 トークン) — そして無駄は有料枠のお金です。
コンテキストのスリム化と回答の品質 ( python scripts/quality_ab.py 、
2026 年 7 月 10 日のライブ、Groq llama-3.1-8b-instant ): 初期段階で植え付けられた 6 つの事実
冗長な会話、3 つの履歴の下で同じ最終クイズ。
2.2 倍の大きな実行 ( --big 、19,145 文字 - 無料利用枠の TPM 全体に近い)
ウィンドウ、1 回の通話の物理的な天井）でパターンが繰り返されました。
プロンプトの 3 分の 1 で A 6/6、B 0/6、C 6/6。
教訓的な失敗は B で、不安定です。1 回の実行で応答しました。
捏造された値（偽のコードネーム、偽の予算、偽の都市）を自信を持って使用します。
大局的には正直に減少しました。どちらであるかを事前に知ることはできません
失敗することもありますが、幻覚を見るほうが危険です。盲目
切り捨ては緊急出口 (§8.6 オーバーフロー、利用可能な LLM なし) であり、
戦略。意味的要約は、すべてをほんの一部で思い出させます。
プロンプトを表示すると、さらに 1 つの呼び出しが得られます。
20 段階のライブ ストレス テスト (2026 年 7 月 10 日) によってループが終了しました。§8.6 の壁では、
スケジューラが一時停止され、チェックポイントがオフラインで圧縮され、スリムに再開され、
429 をゼロで 20/20 ステップを完了しました。
事実と声 ( python scripts/quality_ab.py 、ライブ 2026-07-20、
Groq llama-3.1-8b-instant — 音声プローブは同じ走行に沿って走行します。
追加のフラグは必要ありません): チェックポイントを縮小すると、
内容だけでなく、その内容がどのようなスタイルで書かれているか？ 2つあるペルソナ
短い口頭チックは、上記の 6 つの事実と同じ会話に植え付けられます。
切り捨てと要約は相反する相補的な方法で失敗します: ブラインド
切り捨てにより、短い逐語的なフレーズがそのまま維持されます (長いメッセージのみが切り取られます)。
しかし、歴史の初期に散らばった事実は失われます。要約すると、
事実はありますが、すべてを自分の声で書き換えます。

ナちゃんの話し方
たとえそれが言ったことは要約されても生き残ることはできません。トーンと
エージェントにとってペルソナの一貫性は重要です。compact() と summary_with()
交換可能ではありません。失うわけにはいかないものを選んでください。
ステータス: 0.4.0 。コアスケジューラー、ダイレクト + LiteLLM + Anthropic アダプター、
LangGraph 統合、マルチプロバイダー ルーティング、マルチエージェント共有予算、
機能ベースのコスト/レイテンシの推定、タスク完了予測、チェックポイント
フォークおよびクロスマシン移行、オプションの真のウォームスタート KV キャッシュ プラグイン
llama.cpp の場合、予算としての人間の注意、実行可能なサンプル、および完全な
テストスイートはすべて揃っています。
pip インストール エージェント一時停止
オプションの追加機能 (使用する場合のみ) (コアには依存関係がありません):
pip install " Agentpause[litellm] " # LiteLLM アダプター (100 以上のプロバイダー)
pip install " Agentpause[langgraph] " # LangGraph アダプター
ライブラリ自体での開発 (クローン + 編集可能なインストール + テスト):
git clone https://github.com/Champoleello/agentpause
CD エージェント一時停止
pip install -e " .[dev] "
pytest
はじめに
ステップ 1 — 可能な最小のループ。バックエンドは任意の呼び出し可能です
メッセージ -> (返信、tokens_used) ;テレメトリは任意に呼び出し可能です
() -> 残りのトークン。これが契約全体です。フレームワークは必要ありません。
from agentpause import PredictiveScheduler
sched = PredictiveScheduler (バックエンド = my_llm_call 、テレメトリ = my_rate_limit_reader )
スケジュール付き。 session ( "task-1" ) as s : # 前に中断された場合は自動的に再開します
ご質問については質問にて【ス。 step :]: # サスペンド前に終了したステップをスキップします
s 。 add_user (質問)
の場合。 should_suspend (): # 予測チェック、呼び出しの *前*
s 。チェックポイント ()
Break # きれいに停止します。再開するには再実行
返信 = s 。電話（）
それ以外の場合:
s 。 complete () # タスク完了: チェックポイントを削除します
今すぐ実行してください。API キーは必要ありません

: python Examples/quickstart.py — 一時停止します
最初の実行ではタスクの途中で実行され、2 回目の実行では正常に再開されます。
ステップ 2 — 同じ 3 つのコールをエージェント ループに接続します。何でも
使用するフレームワーク (手動の ReAct ループ、LangGraph、CrewAI、プレーンな
while ループ)、統合は常に同じ 3 つの呼び出しで行われます。
3か所：
LLM を呼び出すたびに、 should_suspend() (または session.next_action() を要求します)
より充実した続行/待機/チェックポイントの回答の場合)。これは地元のものです
計算 — ネットワーク呼び出しもコストもかかりません。
stop と表示された場合は、checkpoint() を呼び出してプロセスを終了します。他には何もありません
やるべきこと: 状態、メッセージ、冪等キーはすでにディスク上にあります。
同じセッション ID で同じコードを再実行すると、自動的に再開されます。
正確なステップから — sched.session(...) を使用して復元を処理します。
ステップ 3 — バックエンドとバックエンドを提供するプロバイダー用のアダプターを選択します。
手動で記述する代わりにテレメトリーを使用できます。
各アダプターの詳細と完全なコードは、すぐ下のセクションにあります。の
## コンポーネント表のさらに下には、完全な関数/クラスのリファレンスがあります。
ライブラリが公開するすべてのものについて。
LiteLLM アダプターは、100 以上のプロバイダー (OpenAI、Groq、
人間的、ローカル サーバーなど)、各応答から予算を読み取る
レート制限ヘッダーと、小さなテレメトリ ping による古い読み取り値の更新:
from agentpause import PredictiveScheduler
エージェント一時停止から。アダプター。 litellm インポート LiteLLMAdapter
アダプター = LiteLLMAdapter (モデル = "groq/llama-3.1-8b-instant")
sched = PredictiveScheduler (バックエンド = アダプター . バックエンド、テレメトリー = アダプター . テレメトリー)
pip install -e ".[litellm]" でインストールします。 Examples/litellm_groq.py を参照してください。
独自のプロバイダーに対して検証するには (フロンティア モデルを含む):
python scripts/validate_provider.py gpt-4o-mini 。
ラ

プロバイダーごとの te-limit ヘッダー (デフォルトは OpenAI スタイルの名前をターゲットとします):
ヘッダー名が違うのでしょうか？それらをオーバーライドします: LiteLLMAdapter(model=..., Remaining_header="...",requests_header="...",reset_header="...")。
Ollama をローカルで実行する
マルティノベットゥッチ/オラマゲートウェイ
(Ollama の自己ホスト型認証 + クォータ プロキシ)?同じものを発します
OpenAI/Groq スタイルの x-ratelimit-* ヘッダーは、agentpause がすでに読み取っているため、
直接アダプターを介して箱から出してすぐに動作します - リテルムも必要もありません
特殊なケース。本物と比較してライブエンドツーエンド (2026-07-20) を確認
ゲートウェイ インスタンス:adapter.budget() は Remaining_tokens を正しく読み取ります。
Remaining_requests (呼び出しごとに減分)、およびreset_秒 —
1 つの条件は、API キーにクォータが設定されていることです (トークン キャップや
レート制限);クォータが設定されていないキーは、単純にレート制限ヘッダーを発行しません。
すべて。agentpause は、エラーではなく明確な TelemetryError として報告します。
沈黙の間違った答え:
エージェント一時停止から。アダプター。 openai_compat インポート OpenAICompatAdapter
アダプター = OpenAICompatAdapter 。 for_model ( "ollam-gateway/llama3:8b" )
# デフォルトのbase_url は http://127.0.0.1:8787/v1 (開発者のデフォルト);を上書きする
# カスタム ポートまたは運用 HTTPS ドメイン:
# OpenAICompatAdapter.for_model("ollama-gateway/llama3:8b",
#base_url="https://ゲートウェイ ドメイン")
トークンを超えて: RPM と待機 vs 一時停止
テレメトリはトークン数よりも豊富な場合があります。 adapter.budget がすべてをレポートします
プロバイダーが公開する 3 つのディメンション — 残りのトークン (TPM)、残り
リクエスト (RPM)、およびウィンドウがリセットされてロックが解除されるまでの秒数
3 つの値の決定:
sched = PredictiveScheduler (バックエンド = アダプター . バックエンド、テレメトリー = アダプター . バジェット )
d = セッション。 next_action () # "続行" | 「待って」 | 「チェックポイント」
wait は、予算が合わない場合に発生しますが、ウィンドウは以内にリセットされます。
wait_threshold

d_s (デフォルトは 15 秒): 短いその場での一時停止は、一時停止よりも安価です。
完全なサスペンド/再開サイクル。リクエストが枯渇した場合 (RPM = 0)、たとえ通話がブロックされる
トークンがたくさん残っています。 Plain-int テレメトリは変更されずに動作し続けます。
すべてのエントリ ポイントには非同期ツインがあります。同じルール、同じ保証、決して同じではありません。
イベントループをブロックします。
アダプター = LiteLLMAdapter (モデル = "groq/llama-3.1-8b-instant")
sched = PredictiveScheduler ( backend = None 、 async_backend =adapter . abackend 、
テレメトリ = アダプター。テレメトリー)
Reply = セッションを待ちます。 acall () # asyncio.sleep による再試行/バックオフ
ガードを待ちます。 acheck ( state [ "messages" ]) # 非同期 LangGraph ノード
入力推定値をより正確にするには、モデルごとのトークナイザーを接続します。
PredictiveScheduler(..., count_tokens=adapter.count_tokens) (フォールバック
トークナイザーが使用できない場合は、~4 文字/トークンのヒューリスティック)。
推定値は統計的なものであり、429 が通過する可能性は依然としてあります。エージェントポーズは生き残る
クラッシュする代わりに:
型付きエラー: すべては AgentPauseError から派生します
( RateLimitHit 、 TelemetryError 、 CheckpointError 、 BackendError );
バックオフで再試行: 予期しない 429 が再試行されます (プロバイダー再試行後)
尊重され、それ以外の場合は指数関数的バックオフ — RetryPolicy を参照)。
ヒットはフィードバックです。各ヒットは safety_k を引き上げます ( k_max で上限が設定されます)。
したがってスケジューラはさらに大きくなります

[切り捨てられた]

## Original Extract

Suspend LLM agents gracefully before rate limits; resume without losing work - Champoleello/agentpause

GitHub - Champoleello/agentpause: Suspend LLM agents gracefully before rate limits; resume without losing work · GitHub
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
Champoleello
/
agentpause
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
25 Commits 25 Commits .github/ workflows .github/ workflows examples examples scripts scripts src/ agentpause src/ agentpause tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
Predictive scheduling for autonomous LLM agents. Suspend an agent gracefully
before it hits a provider rate limit, and resume it without redoing work.
The core works on any provider — cloud (OpenAI, Anthropic, Groq) or local — because
it only serializes application-level state. True KV-cache warm start is an optional
plugin for self-hosted runtimes (llama.cpp, vLLM). Predictive suspension itself still
needs a real budget signal to act on — on a cloud provider that's rate-limit headers;
for a self-hosted server there are none, see "Local mode: what's actually being
controlled" below for what to wire in instead.
Measured results (from the accompanying research)
Experiment
Result
Simulation, 300 runs/config
reactive baseline crashes up to 100% → predictive 0% (k=4), token waste −80%
Real provider (Groq free, 6K TPM)
multi-window task completed with zero 429 errors
vs LangGraph's MemorySaver (200 runs)
LangGraph: 7.8 × 429/run, 9,239 tokens wasted; predictive: 0 and 0
End-to-end A/B (thinking models, 4B/8B)
recovery 54×–93× faster ; total task time −19%
Reproduce them yourself with a free key: python scripts/benchmark_groq.py .
Live run (2026-07-08, Groq free tier, 12 steps, ~1k-token context,
refill-aware chunked waiting, leveled windows):
Zero errors and zero waste at wall-clock parity: the sensor costs ~2% of
what crashes waste (259 vs 12,840 tokens) — and waste is money on paid tiers.
Context slimming vs. answer quality ( python scripts/quality_ab.py ,
live 2026-07-10, Groq llama-3.1-8b-instant ): six facts planted early in a
verbose conversation, same final quiz under three histories.
A 2.2× larger run ( --big , 19,145 chars — near the free tier's whole TPM
window, the physical ceiling for a single call) repeated the pattern:
A 6/6, B 0/6, C 6/6 at a third of the prompt.
The instructive failure is B's, and it is erratic : in one run it answered
confidently with fabricated values (fake codename, fake budget, fake city);
in the larger run it honestly declined. You cannot know in advance which
failure you get — and the hallucinating one is the dangerous one. Blind
truncation is an emergency exit (§8.6 overflow, no LLM available), not a
strategy; semantic summarization recalled everything at a fraction of the
prompt, earning its one extra call.
A 20-step live stress test (2026-07-10) closed the loop: at the §8.6 wall the
scheduler suspended, compacted the checkpoint offline, resumed slim, and
completed 20/20 steps with zero 429s.
Facts vs. voice ( python scripts/quality_ab.py , live 2026-07-20,
Groq llama-3.1-8b-instant — the voice probe rides along on the same run,
no extra flag needed): does shrinking a checkpoint also flatten the
style it was written in, not just the facts inside it? A persona with two
short verbal tics is planted in the same conversation as the six facts above:
Truncation and summarization fail in opposite, complementary ways: blind
truncation keeps short, verbatim phrases intact (only long messages get cut)
but loses facts scattered earlier in the history; summarization recovers the
facts but rewrites everything in its own voice, so a persona's way of talking
does not survive being summarized even when what it said does. If tone and
persona consistency matter for your agent, compact() and summarize_with()
are not interchangeable — pick per what you can't afford to lose.
Status: 0.4.0 . Core scheduler, direct + LiteLLM + Anthropic adapters,
LangGraph integration, multi-provider routing, multi-agent shared budgets,
feature-based cost/latency estimation, task-completion forecast, checkpoint
fork & cross-machine migration, an optional true-warm-start KV-cache plugin
for llama.cpp, human attention as a budget, runnable examples, and a full
test suite are all in place.
pip install agentpause
Optional extras, only if you use them (the core has zero dependencies):
pip install " agentpause[litellm] " # LiteLLM adapter (100+ providers)
pip install " agentpause[langgraph] " # LangGraph adapter
Developing on the library itself (clone + editable install + tests):
git clone https://github.com/Champoleello/agentpause
cd agentpause
pip install -e " .[dev] "
pytest
Getting started
Step 1 — the smallest possible loop. backend is any callable
messages -> (reply, tokens_used) ; telemetry is any callable
() -> remaining_tokens . That's the entire contract — no framework required:
from agentpause import PredictiveScheduler
sched = PredictiveScheduler ( backend = my_llm_call , telemetry = my_rate_limit_reader )
with sched . session ( "task-1" ) as s : # resumes automatically if interrupted before
for question in questions [ s . step :]: # skip steps finished before a suspend
s . add_user ( question )
if s . should_suspend (): # predictive check, *before* the call
s . checkpoint ()
break # stop cleanly; rerun to resume
reply = s . call ()
else :
s . complete () # task done: drop the checkpoint
Run it now, no API key needed: python examples/quickstart.py — it suspends
mid-task on the first run and resumes cleanly the second time you run it.
Step 2 — the same three calls, wired into YOUR agent loop. Whatever
framework you use (a hand-rolled ReAct loop, LangGraph, CrewAI, a plain
while loop), the integration is always the same three calls in the same
three places:
Before every LLM call, ask should_suspend() (or session.next_action()
for the richer continue / wait / checkpoint answer). This is a local
computation — no network call, no cost.
If it says stop, call checkpoint() and exit the process. Nothing else
to do: state, messages, and idempotency keys are already on disk.
Re-running the same code with the same session id resumes automatically
from the exact step — with sched.session(...) handles the restore.
Step 3 — pick the adapter for your provider , which supplies backend and
telemetry for you instead of writing them by hand:
Details and full code for each adapter are in the sections right below; the
## Components table further down is the complete function/class reference
for everything the library exposes.
The LiteLLM adapter supplies both callables for 100+ providers (OpenAI, Groq,
Anthropic, local servers, ...), reading the budget from each response's
rate-limit headers and refreshing stale readings with a tiny telemetry ping:
from agentpause import PredictiveScheduler
from agentpause . adapters . litellm import LiteLLMAdapter
adapter = LiteLLMAdapter ( model = "groq/llama-3.1-8b-instant" )
sched = PredictiveScheduler ( backend = adapter . backend , telemetry = adapter . telemetry )
Install with pip install -e ".[litellm]" ; see examples/litellm_groq.py .
To validate against your own provider (frontier models included):
python scripts/validate_provider.py gpt-4o-mini .
Rate-limit headers by provider (defaults target the OpenAI-style names):
Header names differ? Override them: LiteLLMAdapter(model=..., remaining_header="...", requests_header="...", reset_header="...") .
Running Ollama locally behind
martinobettucci/ollama-gateway
(a self-hosted auth + quota proxy for Ollama)? It emits the same
OpenAI/Groq-style x-ratelimit-* headers agentpause already reads, so it
works out of the box through the direct adapter — no litellm, no
special-casing. Confirmed live end-to-end (2026-07-20) against a real
gateway instance: adapter.budget() correctly reads remaining_tokens ,
remaining_requests (decrementing per call), and reset_seconds — the
one condition is that the API key has a quota configured (token cap and/or
rate limit); a key with no quota set simply emits no rate-limit headers at
all, which agentpause reports as a clear TelemetryError rather than a
silent wrong answer:
from agentpause . adapters . openai_compat import OpenAICompatAdapter
adapter = OpenAICompatAdapter . for_model ( "ollama-gateway/llama3:8b" )
# default base_url is http://127.0.0.1:8787/v1 (dev default); override for
# a custom port or a production HTTPS domain:
# OpenAICompatAdapter.for_model("ollama-gateway/llama3:8b",
# base_url="https://your-gateway-domain")
Beyond tokens: RPM and wait-vs-suspend
Telemetry can be richer than a token count. adapter.budget reports all
three dimensions providers expose — remaining tokens (TPM), remaining
requests (RPM), and seconds until the window resets — and unlocks the
three-valued decision:
sched = PredictiveScheduler ( backend = adapter . backend , telemetry = adapter . budget )
d = session . next_action () # "continue" | "wait" | "checkpoint"
wait fires when the budget does not fit but the window resets within
wait_threshold_s (default 15 s): a short in-place pause is cheaper than a
full suspend/resume cycle. Exhausted requests (RPM = 0) block the call even
with plenty of tokens left. Plain-int telemetry keeps working unchanged.
Every entry point has an async twin — same rules, same guarantees, never
blocks the event loop:
adapter = LiteLLMAdapter ( model = "groq/llama-3.1-8b-instant" )
sched = PredictiveScheduler ( backend = None , async_backend = adapter . abackend ,
telemetry = adapter . telemetry )
reply = await session . acall () # retry/backoff via asyncio.sleep
await guard . acheck ( state [ "messages" ]) # async LangGraph nodes
For sharper input estimates, wire the per-model tokenizer:
PredictiveScheduler(..., count_tokens=adapter.count_tokens) (falls back to
the ~4 chars/token heuristic if the tokenizer is unavailable).
Estimates are statistical — a 429 can still slip through. agentpause survives
it instead of crashing:
typed errors : everything derives from AgentPauseError
( RateLimitHit , TelemetryError , CheckpointError , BackendError );
retry with backoff : unexpected 429s are retried (provider retry-after
honored, else exponential backoff — see RetryPolicy );
hits are feedback : each one bumps safety_k up (capped at k_max ),
so the scheduler grows more

[truncated]
