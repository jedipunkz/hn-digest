---
source: "https://zozo123.github.io/inferoa-demo/"
hn_url: "https://news.ycombinator.com/item?id=48474273"
title: "Inferoa AI harness claimed 90% cache savings. We ran it and measured 97.8%"
article_title: "Inferoa — watch an agent loop spend tokens"
author: "zozo123-IB"
captured_at: "2026-06-10T13:20:46Z"
capture_tool: "hn-digest"
hn_id: 48474273
score: 1
comments: 0
posted_at: "2026-06-10T10:37:44Z"
tags:
  - hacker-news
  - translated
---

# Inferoa AI harness claimed 90% cache savings. We ran it and measured 97.8%

- HN: [48474273](https://news.ycombinator.com/item?id=48474273)
- Source: [zozo123.github.io](https://zozo123.github.io/inferoa-demo/)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T10:37:44Z

## Translation

タイトル: Inferoa AI ハーネスは 90% のキャッシュ節約を主張しました。実行して測定したところ、97.8%
記事のタイトル: Inferoa — エージェントがトークンを消費するループを観察する
説明: 推論ネイティブ エージェント メカニズムの対話型シミュレーション: プレフィックス キャッシュ、コンテキスト圧縮、モデル ルーティング。 vLLM スタック上に構築されています。

記事本文:
インフェロア / フライトレコーダー
お知らせ · vLLM スタック上に構築
エージェント ハーネスは次のとおりです
トークンを燃やす。
それが起こるのを見てください。
Inferoa は、長期的なコーディング作業のための推論ネイティブ エージェント ハーネスです。
推論の仕組みとコスト (プレフィックス キャッシュ、コンテキスト形状、モデル ルーティング) を扱います。
後付けではなく、設計上の制約として。以下: 同じことのターンバイターンのシミュレーション
コーディング タスクは、Inferoa と比較して、単純なハーネスを介して実行されます。
シミュレーションは読みやすい。領収書は信頼できるものです。実際のスタックを分離して実行しました
islo.dev サンドボックス
そして遺物を保管していました。以下のすべてが実際に実行されました。
Real Inferoa v0.11.0 (npm) は実際の
vLLM v0.22.1 サーバー (Qwen2.5-0.5B、CPU)。 vLLM 独自の Prometheus
実行後のカウンター: prefix_cache_hits_total 1,611,008 of
クエリ合計 1,647,574 。注: 発表の 90.0% は
キャッシュされたトークンの割引 (価格設定)。私たちのものは測定されたヒット率です — 関連していますが、別個のものです
どちらも同じバイト安定プレフィックスによって駆動されます。
テストの失敗 → 修正 → グリーン、サンドボックス内
隔離されたサンドボックス内のエージェントに、本当に失敗したテストを含むリポジトリが渡されました
(tz-naive と tz-aware datetime) および疑わしい原因を示すタスク プロンプト。それが生み出した
2行を修正し、pytestを緑色に再実行すると、変更は次のようになります
main にマージされました。出力が埋め込まれる前後で逐語的に
以下のステップ7。
実証済みの内容: ガイドなしのエージェント診断ではなく、サンドボックス化された実行、検証、公開ワークフロー。
1 つは vLLM、もう 1 つは Inferoa
サンドボックス A は、:8000 で vLLM を提供します。 islo シェアがそれを暴露しました
パブリック *.share.islo.dev URL;サンドボックス B の Inferoa が指摘したのは、
そこにbase_urlがあります。 Inferoa のイベント ログにはその証拠が記録されています。
プロバイダー ID: vllm:openai_compatibility:https://…share.islo.dev/v1 、
プロンプトトークン: ターンごとに 16,829、安定したプロンプト/ツールスキーマハッシュを使用 — キャッシュ規律、v

可能です。
0.5B CPU モデルは、最先端グレードではなく、メカニズム (キャッシュ、ルーティング、ハーネス ループ) を証明しています。
コーディング — 本物ではありますが、かわいらしく混乱したツール呼び出しを行いました。 SIM の価格は例示的なものであり、
お読みください。
共有 URL は 24 時間で期限切れになります。リポジトリと PR は永続的です。
以下のすべてのブロックは、2026 年 6 月 10 日の実際の実行からキャプチャされた出力です — 3 つの分離されたブロック
islo.dev サンドボックス、
実際の vLLM、実際の Inferoa、実際のメトリクス。ここでは何も嘲笑されていません。 2 つのパイプライン、2 つの結果:
1 つの CLI コマンドで、公式の vLLM CPU イメージが実行中のサンドボックスに変わります。エンジン設定:enable_prefix_caching=True に注意してください。
$ islo use vllm-cpu-1781079422 -i docker.io/vllm/vllm-openai-cpu:latest-x86_64 \
--CPU 6 --メモリ 12288 --ディスク 30
✓ サンドボックス「vllm-cpu-1781079422」が作成されました
$ vllmserve Qwen/Qwen2.5-0.5B-Instruct --dtype bfloat16 --max-model-len 32768
(EngineCore) INFO core.py:112 構成を使用して V1 LLM エンジン (v0.22.1) を初期化しています:
model='Qwen/Qwen2.5-0.5B-Instruct'、device_config=cpu、enable_prefix_caching=True …
(ワーカー) 情報 重みのダウンロードにかかった時間: 10.27 秒
(作業員) 情報 ウェイトの読み込みには 0.27 秒かかりました
実際の出力・サンドボックスA
2
世界に公開してみよう
もう 1 つのコマンドにより、サンドボックス ポートに islo.dev ドメインのパブリック HTTPS URL が与えられます。
$ islo シェア vllm-cpu-1781079422 8000 --ttl 24h
✓ vllm-cpu-1781079422:8000 に対して共有が作成されました
URL: https://itc807adoyzwz6hzj0rkmcp7m.ca.share.islo.dev
有効期限: 1日 0時間以内
実際の出力 · 共有 URL (24 時間 ttl — これを読むまでに有効期限が切れている可能性があります)
3
推論が真実であることを証明する
公開 URL による完了。 「INFERENCE OK」をエコーするように求められた場合、0.5B モデルは「INFORMATION OK」と応答しました。モックは正確にエコーしたはずです。タイプミスが証拠だ。
$カール https://itc807….ca.share.islo.dev/v1/chat/completions -d '{
"モデル":"Qwen/Qwen2.5-0.5B-Instr

うーん」、
"messages":[{"role":"user","content":"正確に返信してください: INFERENCE OK"}]}'
返信: 情報OK
使用法: {'prompt_tokens': 36、'total_tokens': 40、'completion_tokens': 4}
実際の出力 - islo.dev を介してラップトップからサンドボックス A への外部リクエスト
4
本物のインフェロアをそこに向けてください
2 番目のサンドボックス: 実際の npm パッケージをインストールし、そのエンジン構成を共有 URL に設定します。 Inferoa のデフォルトのプロバイダーは文字通り vLLM です。
$ npm install -g inferoa # inferoa@0.11.0
$ cat ~/.inferoa/config.yaml
モデルセットアップ:
モード: ダイレクト
プロバイダー: vllm
モデル: Qwen/Qwen2.5-0.5B-Instruct
Base_url: https://itc807….ca.share.islo.dev/v1
$ inferoa デバッグステータス --json | jq .endpoint_signals.errors
[]
実際の出力・サンドボックスB
5
実際にハーネスを動かしてみます
Inferoa 独自のイベント ログには、各モデル呼び出し、つまり islo.dev プロバイダー URL、16,829 トークン プロンプト、および安定したプロンプト/ツール スキーマ ハッシュ (キャッシュ規律機構が表示されます) が記録されます。次に、小さなモデルは実際の、しかし混乱したツール呼び出しを行いました (is_palindrome を作成する代わりに呼び出そうとしました)。正直な制限: 0.5B はコーディング スキルではなく、メカニズムを証明します。
$ inferoa --print "Python 関数 is_palindrome(s) を作成します…"
$ inferoa デバッグ イベント s_6b9da06c…
タイプ:model.request.started
プロバイダー ID: vllm:openai_compatibility:https://itc807….ca.share.islo.dev/v1
プロンプト_ハッシュ: 990301121c… ツール_スキーマ_ハッシュ: a381b770ce…
タイプ:model.response.settled
使用法: { プロンプトトークン: 16829 、完了トークン: 30 }
tool_calls: [ { 名前: is_palindrome,
引数: { s: "人間、計画、運河、パナマ" } } ]
実際の出力 · inferoa イベント ログ、セッション s_6b9da06c…
6
vLLM 自体からキャッシュを読み取ります
私たちの数字ではなく、セッション後の vLLM の Prometheus カウンターです。 1,647,574 個のプロンプト トークンのうち 1,611,008 個がプレフィックス キャッシュにヒットしました。ヒット率は 97.8% でした。 (アナウンスの 90.0% はキャッシュされたトークンのディス

count — 価格設定の指標。これがヒット率です。同じメカニズム: バイト安定プレフィックス。)
$カールローカルホスト:8000/メトリック | grep プレフィックスキャッシュ
vllm:prefix_cache_queries_total{model="Qwen/Qwen2.5-0.5B-Instruct"} 1,647,574
vllm:prefix_cache_hits_total{model="Qwen/Qwen2.5-0.5B-Instruct"} 1,611,008
# ヒット率 = 97.8%
実際の出力 · サンドボックス A、vllm /metrics
7
エージェントループと領収書
これとは別に、3 番目のサンドボックス内のエージェントには、実際に失敗したテストと疑わしい原因が記載されたリポジトリが渡されました。修正が行われ、pytest が緑色に再実行されました。変更は次のとおりです。
メイン (ffda3d7) にマージされました。前/修正/後:
$ python3 -m pytest -v
失敗したテスト/test_token_refresh.py::test_token_refresh
E TypeError: オフセット ナイーブ日付時刻とオフセット認識日付時刻を比較できません
--- a/auth_service/token.py
+++ b/auth_service/token.py
- 今 = datetime.utcnow()
+ 今 = datetime.now(timezone.utc)
$ python3 -m pytest -v
テスト/test_token_refresh.py::test_token_refresh 合格
テスト/test_token_refresh.py::test_expired_token_rejected 合格
============== 2 合格 ==============
実際の出力 · サンドボックス C · zozo123/inferoa-receipts@main にマージ
3つのレバーを掛け合わせたもの。
各メカニズムは他のメカニズムと複合します。この割引は単なるトリックではなく、推論を第一級のエンジニアリングの表面として扱うことによって生み出されたものです。
ハーネスは、システム プロンプト、ツール スキーマ、リポジトリ コンテキストなど、プロンプト プレフィックスをターン全体にわたってバイト安定に保つため、vLLM のプレフィックス キャッシュ (およびフロンティア キャッシュの価格設定) では、全額ではなくキャッシュされた料金で繰り返しコンテキストを請求します。
ファイル全体を貼り付ける代わりに、エージェントはリポジトリのグラフ状のスライス (タスクにとって重要なシンボル、エッジ、スパン) を運びます。同じシグナル、トークンの 5 分の 1。
生のツール出力はサイレント トークン キラーです。 RTK はコンパクトで構造化されたコマンド リソースを記録します

ults — 終了コード、デルタ、失敗したアサーション — 400 行のテスト ランナー ノイズではありません。
習慣ではなく、経済性によってルートを決めます。
vLLM セマンティック ルーターは、各ステップを適切な場所に送信します。機械的な編集とログのダイジェストを限界コストで自己ホスト型モデルに送信します。機能が実際にそれ自体で報われる場合、フロンティアモデルに対するアーキテクチャ上の推論。コスト、プライバシー、機能はルーティング入力であり、雰囲気ではありません。
「vLLM スタック上に構築されたコミュニティ エージェント ハーネスである Inferoa を見るのは素晴らしいことです。Inferoa は、プレフィックス キャッシュの規律、コンテキストの最適化、自己ホスト型モデルとフロンティア モデルにわたるルーティングを使用して、推論の仕組みとコストを中心にエージェント ループを設計します。」
上記のシミュレーションは例示です。ハーネスは本物です:

## Original Extract

Interactive simulation of inference-native agent mechanics: prefix caching, context compression, and model routing. Built on the vLLM stack.

INFEROA / FLIGHT RECORDER
announcement · built on the vLLM stack
Your agent harness is
burning tokens.
Watch it happen.
Inferoa is an inference-native agent harness for long-horizon coding work.
It treats the mechanics and cost of inference — prefix caches, context shape, model routing —
as the design constraint, not an afterthought. Below: a turn-by-turn simulation of the same
coding task run through a naive harness vs. Inferoa.
A simulation is legible; receipts are credible. We ran the real stack in isolated
islo.dev sandboxes
and kept the artifacts. Everything below actually executed.
Real Inferoa v0.11.0 (npm) drove a real
vLLM v0.22.1 server (Qwen2.5-0.5B, CPU). vLLM's own Prometheus
counters after the run: prefix_cache_hits_total 1,611,008 of
queries_total 1,647,574 . Note: the announcement's 90.0% is a
cached-token discount (pricing); ours is the measured hit rate — related but distinct
metrics, both driven by the same byte-stable prefixes.
failing test → fix → green, in a sandbox
An agent in an isolated sandbox was handed a repo with a genuinely failing test
(tz-naive vs tz-aware datetime) and a task prompt that named the suspected cause. It produced
the 2-line fix, re-ran pytest to green, and the change is
merged on main . Verbatim before/after output is embedded in
step 7 below .
What's proven: the sandboxed execute-verify-publish workflow — not unguided agent diagnosis.
vLLM in one, Inferoa in the other
Sandbox A serves vLLM on :8000; islo share exposed it at a
public *.share.islo.dev URL; sandbox B's Inferoa pointed its
base_url there. Inferoa's event log records the proof:
provider_id: vllm:openai_compatible:https://…share.islo.dev/v1 ,
prompt_tokens: 16,829 per turn, with stable prompt/tool-schema hashes — cache discipline, visible.
The 0.5B CPU model proves the mechanics (caching, routing, harness loop), not frontier-grade
coding — it made a real but adorably confused tool call. Sim pricing is illustrative and published in the
README .
Share URLs expire in 24h; the repos and PR are permanent.
Every block below is captured output from the actual run on 2026-06-10 — three isolated
islo.dev sandboxes,
real vLLM, real Inferoa, real metrics. Nothing here is mocked. Two pipelines, two results:
One CLI command turns the official vLLM CPU image into a running sandbox. Note the engine config: enable_prefix_caching=True .
$ islo use vllm-cpu-1781079422 -i docker.io/vllm/vllm-openai-cpu:latest-x86_64 \
--cpu 6 --memory 12288 --disk 30
✓ Sandbox 'vllm-cpu-1781079422' created
$ vllm serve Qwen/Qwen2.5-0.5B-Instruct --dtype bfloat16 --max-model-len 32768
(EngineCore) INFO core.py:112 Initializing a V1 LLM engine (v0.22.1) with config:
model='Qwen/Qwen2.5-0.5B-Instruct', device_config=cpu, enable_prefix_caching=True …
(Worker) INFO Time spent downloading weights: 10.27 s
(Worker) INFO Loading weights took 0.27 seconds
real output · sandbox A
2
Expose it to the world
One more command gives the sandbox port a public HTTPS URL on the islo.dev domain.
$ islo share vllm-cpu-1781079422 8000 --ttl 24h
✓ Share created for vllm-cpu-1781079422:8000
URL: https://itc807adoyzwz6hzj0rkmcp7m.ca.share.islo.dev
Expires: in 1d 0h
real output · share url (24h ttl — may be expired by the time you read this)
3
Prove the inference is real
A completion through the public URL. Asked to echo “INFERENCE OK” , the 0.5B model replied “INFORMATION OK” — a mock would have echoed exactly. The typo is the proof.
$ curl https://itc807….ca.share.islo.dev/v1/chat/completions -d '{
"model":"Qwen/Qwen2.5-0.5B-Instruct",
"messages":[{"role":"user","content":"Reply with exactly: INFERENCE OK"}]}'
reply: INFORMATION OK
usage: {'prompt_tokens': 36, 'total_tokens': 40, 'completion_tokens': 4}
real output · external request from a laptop, through islo.dev, into sandbox A
4
Point real Inferoa at it
Second sandbox: install the actual npm package and aim its engine config at the share URL. Inferoa's default provider is literally vLLM.
$ npm install -g inferoa # inferoa@0.11.0
$ cat ~/.inferoa/config.yaml
model_setup:
mode: direct
provider: vllm
model: Qwen/Qwen2.5-0.5B-Instruct
base_url: https://itc807….ca.share.islo.dev/v1
$ inferoa debug status --json | jq .endpoint_signals.errors
[]
real output · sandbox B
5
Run the harness for real
Inferoa's own event log records each model call: the islo.dev provider URL, a 16,829-token prompt , and stable prompt/tool-schema hashes — the cache-discipline machinery, visible. The tiny model then made a real but confused tool call (it tried to invoke is_palindrome instead of writing it). Honest limit: 0.5B proves mechanics, not coding skill.
$ inferoa --print "Write a Python function is_palindrome(s)…"
$ inferoa debug events s_6b9da06c…
type: model.request.started
provider_id: vllm:openai_compatible:https://itc807….ca.share.islo.dev/v1
prompt_hash: 990301121c… tool_schema_hash: a381b770ce…
type: model.response.settled
usage: { prompt_tokens: 16829 , completion_tokens: 30 }
tool_calls: [ { name: is_palindrome,
arguments: { s: "A man, a plan, a canal, Panama" } } ]
real output · inferoa event log, session s_6b9da06c…
6
Read the cache off vLLM itself
Not our numbers — vLLM's Prometheus counters after the session. 1,611,008 of 1,647,574 prompt tokens hit the prefix cache: a 97.8% hit rate. (The announcement's 90.0% is a cached-token discount — a pricing metric; this is the hit rate . Same mechanism: byte-stable prefixes.)
$ curl localhost:8000/metrics | grep prefix_cache
vllm:prefix_cache_queries_total{model="Qwen/Qwen2.5-0.5B-Instruct"} 1,647,574
vllm:prefix_cache_hits_total{model="Qwen/Qwen2.5-0.5B-Instruct"} 1,611,008
# hit rate = 97.8%
real output · sandbox A, vllm /metrics
7
And the agent loop, with receipts
Separately, an agent in a third sandbox was handed a repo with a genuinely failing test and the suspected cause. It produced the fix, re-ran pytest to green, and the change is
merged on main (ffda3d7) . Before / fix / after:
$ python3 -m pytest -v
FAILED tests/test_token_refresh.py::test_token_refresh
E TypeError: can't compare offset-naive and offset-aware datetimes
--- a/auth_service/token.py
+++ b/auth_service/token.py
- now = datetime.utcnow()
+ now = datetime.now(timezone.utc)
$ python3 -m pytest -v
tests/test_token_refresh.py::test_token_refresh PASSED
tests/test_token_refresh.py::test_expired_token_rejected PASSED
============== 2 passed ==============
real output · sandbox C · merged to zozo123/inferoa-receipts@main
Three levers, multiplied.
Each mechanism compounds with the others. The discount isn't one trick — it's the product of treating inference as a first-class engineering surface.
The harness keeps prompt prefixes byte-stable across turns — system prompt, tool schemas, repo context — so vLLM's prefix cache (and frontier cache pricing) bills repeated context at the cached rate instead of full price.
Instead of pasting whole files, the agent carries a graph-shaped slice of the repository — symbols, edges, and the spans that matter for the task. Same signal, a fifth of the tokens.
Raw tool output is the silent token killer. RTK records compact, structured command results — exit codes, deltas, the failing assertion — not 400 lines of test runner noise.
Route by economics, not habit.
The vLLM Semantic Router sends each step where it belongs: mechanical edits and log digestion to self-hosted models at marginal cost; architectural reasoning to frontier models when capability actually pays for itself. Cost, privacy, and capability are routing inputs — not vibes.
“Great to see Inferoa, a community agent harness built on the vLLM stack — it designs the agent loop around the mechanics and cost of inference, with prefix-cache discipline, context optimization, and routing across self-hosted and frontier models.”
The simulation above is illustrative. The harness is real:
