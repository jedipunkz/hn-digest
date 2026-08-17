---
source: "https://alexander-hanel.github.io/StressingLLMs/"
hn_url: "https://news.ycombinator.com/item?id=49330655"
title: "Show HN: Stressing LLMs – Complexity Benchmark"
article_title: "StressingLLMs - Complexity Benchmark - campaign_89fa287a385e212f"
image: ""
author: "__alexander"
captured_at: "2026-08-17T14:19:10Z"
capture_tool: "hn-digest"
hn_id: 49330655
score: 1
comments: 0
posted_at: "2026-08-17T13:39:07Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Stressing LLMs – Complexity Benchmark

- HN: [49330655](https://news.ycombinator.com/item?id=49330655)
- Source: [alexander-hanel.github.io](https://alexander-hanel.github.io/StressingLLMs/)
- Score: 1
- Comments: 0
- Posted: 2026-08-17T13:39:07Z

## Translation

タイトル: HN の表示: LLM のストレス – 複雑さのベンチマーク
記事のタイトル: StressingLLMs - 複雑さのベンチマーク - Campaign_89fa287a385e212f
HN テキスト: こんにちは、HN。これは、ローカルの大規模言語モデルを探索するための学習プロジェクトです。私は、LLM が増大する複雑さのレベルにどの程度うまく対処できるのか疑問に思っていました。このトピックに関する他のリソースや研究に関する推奨事項を持っている人はいますか?

記事本文:
StressingLLM - 複雑さのベンチマーク
Campaign_89fa287a385e212f のレビュー担当者のスナップショット
このプロジェクトでは、単一の NVIDIA DGX Spark 上で実行されるローカル言語モデルを評価します。生成された C バイナリをリバース エンジニアリングして、決定論的な XOR ベースの復号化によって保護された平文メッセージを復元できるかどうかをテストします。各フィクスチャは、構成可能な変換ラウンド数を使用して、解析の課題を増大させます。
試行ごとに、モデルはバイナリを検査し、Python 復号化プログラムを送信し、ハーネスは固定された非ネットワーク サンドボックスでそのコードを実行します。実行された Python ブロックが予想されるプレーンテキストを出力すると、試行は成功します。 「ラウンドごとの結果」には、テストされた各複雑さレベルでの各モデルの合否結果が表示されます。
これは fixtures/src/fx_r0002_sl0016_sp0000.c の完全なソースです。これには、ヘルパー、生成されたラウンド関数、状態連鎖、および最終的なバイト単位の XOR ループが含まれます。
#include <stdint.h>
#include <stdio.h>
#include <stddef.h>
静的 uint32_t xorshift32(uint32_t x) {
x ^= x << 13;
x ^= x >>17;
x ^= x << 5;
x を返します。
}
typedef struct TokenizerBench___Type__LongRecord__With__Lots__Of__Nested__Like__Tokens {
uint64_t a;
uint64_t b;
uint64_t c;
TokenizerBench___Type__LongRecord__With__Lots__Of__Nested__Like__Tokens;
__attribute__((使用、インラインなし))
uint32_t TokenizerBench___R0(TokenizerBench___Type__LongRecord__With__Lots__Of__Nested__Like__Tokens *p) {
uint32_t m1 = xorshift32(0x9336956du ^ 0x31bbf978u ^ (uint32_t)p->a);
uint32_t m2 = xorshift32(0xcd6f55fcu ^ (uint32_t)p->b);
p->a ^= ((uint64_t)m1 << 32) | (uint64_t)m2;
p->b += (uint64_t)(0x9336956du ^ m2);
p->b = (p->b << 10) | (p->b >> 54);
p->c = (p->c + p->a) ^ (uint64_t)(0x31bbf978u ^ 0xcd6f55fcu);
uint64_t r = p->a ^ p->b ^ p->c ^ (uint64_t)0x9336956du ^ (uint64_t)0x31bbf978u

^ (uint64_t)0xcd6f55fcu;
return (uint32_t)(r ^ (r >> 32));
}
__attribute__((使用、インラインなし))
uint32_t TokenizerBench___R1(TokenizerBench___Type__LongRecord__With__Lots__Of__Nested__Like__Tokens *p) {
uint32_t m = xorshift32(0x366856bbu ^ (uint32_t)p->a);
p->a ^= ((uint64_t)0x366856bbu << 32) | (uint64_t)m;
p->b += p->a ^ (p->c + (uint64_t)0x72fcd409u);
p->c = ((p->c ^ (uint64_t)0x3afd4cabu) << 24)
| ((p->c ^ (uint64_t)0x3afd4cabu) >> 40);
uint64_t r = p->a ^ p->b ^ p->c ^ (uint64_t)0x366856bbu ^ (uint64_t)0x72fcd409u ^ (uint64_t)0x3afd4cabu;
return (uint32_t)(r ^ (r >> 32));
}
__attribute__((使用、インラインなし))
uint32_t 派生_状態(uint32_t シード) {
TokenizerBench___Type__LongRecord__With__Lots__Of__Nested__Like__Tokens x = {
シード、シード ^ 0x12345678ULL、シード + 0x9ULL
};
uint32_t s = シード;
s ^= TokenizerBench___R0(&x);
s ^= TokenizerBench___R1(&x);
s = xorshift32(s);
を返します。
}
int main(void) {
uint8_t 暗号化[] = { 0xbc、0xbb、0x47、0xdd、0xb7、0xbc、0x88、0x9f、
0xba、0x15、0xdd、0x2e、0x00 };
uint32_t s = 派生_状態(0xdeadbeef);
for (size_t i = 0; i < sizeof(暗号化) - 1; i++) {
s = xorshift32(s + 0
[切り捨てられた]
ベンチマークはフィクスチャを生成してコンパイルし、一致するバイナリを分析のために Ghidra にインポートします。
LLM はリバース エンジニアリング タスクを受け取り、Ghidra MCP サーバー経由でバイナリにアクセスし、シンボル、バイト、逆アセンブリ、逆コンパイルなどの分析操作を公開します。
LLM は、これらの MCP/Ghidra の結果を使用して、状態の初期化、生成されたラウンド関数、定数、および最終的な XOR キーストリームを推論します。
LLM は Python 復号化プログラムを送信します。グレーダーは、提出形式に従っているすべての Python、py 、またはラベルのないフェンスで囲まれたブロックを抽出します。
抽出された各ブロックは、固定されたネットワークなしの Python サンドボックスで実行されます。 e が返された場合、試行は成功します。

xecuted ブロックは、予期されるプレーンテキストを出力します。それ以外の場合、レポートには利用可能な最も具体的な障害カテゴリが記録されます。
このテストでは、モデルの次の能力を評価します。
Ghidra と MCP を使用して、コンパイルされたバイナリをナビゲートします。
生成された状態変換と定数を回復します。
難読化されたコードの複数ラウンドにわたる理由。
実行可能な Python 復号化プログラムを生成します。
リバースエンジニアリングの複雑さが増大しても持続します。
使用可能なソリューションを期待される形式で伝えます。
フィクスチャは合成であり、1 つの変換ファミリーから生成されます。
すべてが 1 つの DGX Spark 上で実行されるため、結果はハードウェアとタイムアウトに依存します。
モデルごとのサンプル サイズは不均一です。
適応型検索では、合格率が必ずしも直接比較できるわけではありません。
このベンチマークは、モデル推論を単独で行うだけでなく、モデルとツールを組み合わせたワークフロー全体を測定します。
各セルは、その正確なラウンド数に対する 1 つのモデル試行です。パスとは、実行された Python ブロックが予期されたプレーンテキストを出力したことを意味します。
インタラクティブ: セルにカーソルを合わせるかセルを選択すると、失敗の理由と実行時間が表示されます。
プロバイダー、チャレンジ、失敗、追跡、または結果ごとに試行をフィルタリングします。任意の行を選択して、そのフィクスチャ、ランタイム、結果、サニタイズされたレコードを検査します。
インタラクティブ: フィルタを変更してテーブルを更新し、行をクリックして試行の完全な詳細を開きます。
現在のフィルターが適用された後、記録された失敗カテゴリごとに失敗した試行をカウントします。
フィルタリングされた試行のうち、各プロバイダーの検証済みのパスと合格率を表示します。
これらの結果は、ベンチマークの制御された試行設定の下でローカル モデルを比較し、テストされたチャレンジ ラウンド全体で各モデルがどこまで進歩したかを示しています。
このプロファイルは、完了した試行に適用されたチェックと、検証された復号結果が生成された割合を要約します。
この表には、最近完了したものがリストされています。

トラック、プロバイダー、チャレンジ レベル、結果、処理時間を含む試行。
検証済みの復号化は従来のランナーと一致します。すべての Python、py 、またはラベルのないフェンスされたブロックが実行され、いずれかのブロックの標準出力に予期されるプレーンテキストが含まれている場合、試行は成功します。関数の署名、正確なフェンス数、構造化された証拠、および隠れた代替ベクトルは採点要件ではありません。抽出されたコードは、固定されたネットワークなしの Docker サンドボックスで引き続き実行されます。
検索フィクスチャの成功には、設定された復号化メトリクスとフィクスチャごとに 1 回の試行が使用されます。境界は最高合格/最初の不合格です。打ち切り値は、検索で両方の側面が観察されなかったことを意味します。複合ランキングは計算されません。
このキャンペーンでは、完了した 62 回の試行のうち 18 回 (29.0%) で予想される平文が検証されました。
ollama-gemma4-31b-it-qat は最も強力な進行を示し、テストされた 8 ラウンドのうち 62.5% を検証し、ラウンド 6 に到達しました。
その他の一貫したパフォーマンスは、ollama-gemma4-26b、ollama-gemma4-31b、ollama-qwen3-5-27b、vllm-muse-glimmer で、それぞれテストされたラウンドの少なくとも 60% を検証しました。まとめると、これらの結果は、生成された丸い構造がより要求が厳しくなるにつれて、明確な分離を示しています。

## Original Extract

Hi HN, this has been a learning project for exploring local large language models. I’ve been wondering how well LLMs handle increasing levels of complexity. Does anyone have recommendations for other resources or research on this topic?

StressingLLMs - Complexity Benchmark
Reviewer snapshot for campaign_89fa287a385e212f
This project evaluates local language models running on a single NVIDIA DGX Spark. It tests how well they can reverse-engineer generated C binaries and recover a plaintext message protected by deterministic XOR-based decryption. Each fixture uses configurable number of transformation rounds to increase the analysis challenge.
For every attempt, the model inspects the binary, submits a Python decryptor, and the harness executes that code in a pinned no-network sandbox. The attempt passes when an executed Python block prints the expected plaintext. “Results by round” shows the pass/fail outcome for each model at each tested complexity level.
This is the complete source from fixtures/src/fx_r0002_sl0016_sp0000.c . It includes the helper, generated round functions, state chaining, and final byte-wise XOR loop.
#include <stdint.h>
#include <stdio.h>
#include <stddef.h>
static uint32_t xorshift32(uint32_t x) {
x ^= x << 13;
x ^= x >> 17;
x ^= x << 5;
return x;
}
typedef struct TokenizerBench___Type__LongRecord__With__Lots__Of__Nested__Like__Tokens {
uint64_t a;
uint64_t b;
uint64_t c;
} TokenizerBench___Type__LongRecord__With__Lots__Of__Nested__Like__Tokens;
__attribute__((used, noinline))
uint32_t TokenizerBench___R0(TokenizerBench___Type__LongRecord__With__Lots__Of__Nested__Like__Tokens *p) {
uint32_t m1 = xorshift32(0x9336956du ^ 0x31bbf978u ^ (uint32_t)p->a);
uint32_t m2 = xorshift32(0xcd6f55fcu ^ (uint32_t)p->b);
p->a ^= ((uint64_t)m1 << 32) | (uint64_t)m2;
p->b += (uint64_t)(0x9336956du ^ m2);
p->b = (p->b << 10) | (p->b >> 54);
p->c = (p->c + p->a) ^ (uint64_t)(0x31bbf978u ^ 0xcd6f55fcu);
uint64_t r = p->a ^ p->b ^ p->c ^ (uint64_t)0x9336956du ^ (uint64_t)0x31bbf978u ^ (uint64_t)0xcd6f55fcu;
return (uint32_t)(r ^ (r >> 32));
}
__attribute__((used, noinline))
uint32_t TokenizerBench___R1(TokenizerBench___Type__LongRecord__With__Lots__Of__Nested__Like__Tokens *p) {
uint32_t m = xorshift32(0x366856bbu ^ (uint32_t)p->a);
p->a ^= ((uint64_t)0x366856bbu << 32) | (uint64_t)m;
p->b += p->a ^ (p->c + (uint64_t)0x72fcd409u);
p->c = ((p->c ^ (uint64_t)0x3afd4cabu) << 24)
| ((p->c ^ (uint64_t)0x3afd4cabu) >> 40);
uint64_t r = p->a ^ p->b ^ p->c ^ (uint64_t)0x366856bbu ^ (uint64_t)0x72fcd409u ^ (uint64_t)0x3afd4cabu;
return (uint32_t)(r ^ (r >> 32));
}
__attribute__((used, noinline))
uint32_t derive_state(uint32_t seed) {
TokenizerBench___Type__LongRecord__With__Lots__Of__Nested__Like__Tokens x = {
seed, seed ^ 0x12345678ULL, seed + 0x9ULL
};
uint32_t s = seed;
s ^= TokenizerBench___R0(&x);
s ^= TokenizerBench___R1(&x);
s = xorshift32(s);
return s;
}
int main(void) {
uint8_t encrypted[] = { 0xbc, 0xbb, 0x47, 0xdd, 0xb7, 0xbc, 0x88, 0x9f,
0xba, 0x15, 0xdd, 0x2e, 0x00 };
uint32_t s = derive_state(0xdeadbeef);
for (size_t i = 0; i < sizeof(encrypted) - 1; i++) {
s = xorshift32(s + 0
[truncated]
The benchmark generates and compiles a fixture, then imports the matching binary into Ghidra for analysis.
The LLM receives the reverse-engineering task and accesses the binary through the Ghidra MCP server, which exposes analysis operations such as symbols, bytes, disassembly, and decompilation.
The LLM uses those MCP/Ghidra results to infer the state initialization, generated round functions, constants, and final XOR keystream.
The LLM submits a Python decryptor. The grader extracts every Python, py , or unlabeled fenced block that follows the submission format.
Each extracted block runs in the pinned no-network Python sandbox. The attempt passes when an executed block prints the expected plaintext; otherwise the report records the most specific failure category available.
This test evaluates a model’s ability to:
Navigate a compiled binary using Ghidra and MCP.
Recover generated state transformations and constants.
Reason across multiple rounds of obfuscated code.
Produce an executable Python decryptor.
Persist through increasing reverse-engineering complexity.
Communicate a usable solution in the expected format.
Fixtures are synthetic and generated from one family of transformations.
Results are hardware- and timeout-dependent because everything ran on one DGX Spark.
The sample size per model is uneven.
Adaptive search means pass rates are not always directly comparable.
The benchmark measures the full model-plus-tool workflow, not just model reasoning in isolation.
Each cell is one model attempt on that exact round count. A pass means an executed Python block printed the expected plaintext.
Interactive: Hover over or select a cell to see its failure reason and runtime.
Filter attempts by provider, challenge, failure, track, or result. Select any row to inspect its fixture, runtime, outcome, and sanitized record.
Interactive: Change a filter to update the table, then click a row to open its full attempt details.
Counts failed attempts by their recorded failure category after the current filters are applied.
Shows verified passes and pass rates for each provider among the filtered attempts.
These results compare local models under the benchmark’s controlled attempt settings, showing how far each model progressed across the tested challenge rounds.
This profile summarizes the checks applied to completed attempts and the proportion that produced a verified decryption result.
This table lists the most recently completed attempts, including their track, provider, challenge level, result, and processing time.
Verified decryption matches the legacy runner: every Python, py , or unlabeled fenced block is executed, and the attempt passes when any block's standard output contains the expected plaintext. Function signatures, exact fence counts, structured evidence, and hidden alternate vectors are not grading requirements. Extracted code still runs in the pinned no-network Docker sandbox.
Search fixture success uses the configured decryptor metric and 1 attempt(s) per fixture. Boundaries are highest passing / first failing; censored values mean the search did not observe both sides. No composite ranking is computed.
This campaign verified the expected plaintext in 18 of 62 completed attempts (29.0%).
ollama-gemma4-31b-it-qat showed the strongest progression, verifying 62.5% of its 8 tested rounds and reaching round 6.
Other consistent performers were ollama-gemma4-26b, ollama-gemma4-31b, ollama-qwen3-5-27b, vllm-muse-glimmer, each verifying at least 60% of tested rounds. Together, these results show clear separation as the generated round structure becomes more demanding.
