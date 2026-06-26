---
source: "http://www.theresistornetwork.com/2025/03/thinking-different-thinking-slowly-llms.html"
hn_url: "https://news.ycombinator.com/item?id=48685145"
title: "Thinking Different, Thinking Slowly: LLMs on a PowerPC Mac"
article_title: "The Resistor Network: Thinking Different, Thinking Slowly: LLMs on a PowerPC Mac"
author: "asimovDev"
captured_at: "2026-06-26T11:57:06Z"
capture_tool: "hn-digest"
hn_id: 48685145
score: 1
comments: 0
posted_at: "2026-06-26T11:00:20Z"
tags:
  - hacker-news
  - translated
---

# Thinking Different, Thinking Slowly: LLMs on a PowerPC Mac

- HN: [48685145](https://news.ycombinator.com/item?id=48685145)
- Source: [www.theresistornetwork.com](http://www.theresistornetwork.com/2025/03/thinking-different-thinking-slowly-llms.html)
- Score: 1
- Comments: 0
- Posted: 2026-06-26T11:00:20Z

## Translation

タイトル: 異なる考え方、ゆっくり考える: PowerPC Mac 上の LLM
記事のタイトル: 抵抗ネットワーク: 異なる考え方、ゆっくり考える: PowerPC Mac 上の LLM
説明: 古いハードウェアに新しい命を吹き込むことには、信じられないほど満足のいくものがあります。ビンテージ コンピューティングは私のお気に入りの趣味の 1 つです。チャ...

記事本文:
抵抗ネットワーク: 異なる考え方、ゆっくり考える: PowerPC Mac 上の LLM
2025年3月24日月曜日
違う考え方、ゆっくり考える: PowerPC Mac 上の LLM
ステータスリターンを導入し、終了するための呼び出しをすべて削除します。
ステータス処理を簡素化するための抽象的なファイル アクセス
抽象malloc / 簡単なデバッグ/分析には無料
単一文字列の 512 バイトの静的 LUT を置き換えます。
-Wall を使用してコンパイルする際のいくつかの警告を修正
// 推論操作の実行時設定。
typedef 構造体 {
// 応答を生成するプロンプト。
const char* プロンプト;
// チェックポイント ファイルへのパス。
const char* チェックポイント_パス;
// トークナイザー ファイルへのパス。
const char* tokenizer_path;
// モデルの設定。
フロート温度。
フロートトップ。
unsigned int ステップ。
// エントロピーの源。
uint64_t rng_seed;
// 生成された出力のコールバックとコンテキスト。
void (*output_callback)(const char* トークン、void* cookie);
void* クッキー;
UllmLama2RunConfig;
// 推論エンジンの実行時の状態。
typedef 構造体 {
UllmFileHandle チェックポイント_ファイル;
UllmLalama2Transformer 変圧器;
UllmLama2Tokenizer トークナイザー;
UllmLalama2Sampler サンプラー。
UllmLalama2State;
ライブラリは 2 つの入力を想定しています。モデル パス、RNG シード、トークン出力コールバックなどの詳細を提供する const 構成と、ロードされた重み、一時バッファー、トークナイザーの状態などを追跡する状態です。
結果として得られる API はテストが非常に簡単で、コマンドライン インターフェイス ツールを簡単に構築できます。
コールバックベースの出力とテスト
void OutputHandler(const char* トークン、void* cookie) {
std::string* test_output = static_cast<std::string*>(cookie);
test_output->append(トークン);
}
TEST(ウルムラマ2、ストーリー15M) {
const std::string Expected_test_output = R"(鳥が鳴く。どこへ行く?
鳥が飛んだ

空を飛び回って、何かすることを探しています。
鳥たちは大きな木を見つけて、そこに飛んでいきました。
鳥たちは地面に大きな赤いリンゴを見つけました。美味しそうでした。
鳥たちが飛んできて、リンゴを拾いました。
鳥たちは木に飛んで戻り、リンゴを食べ始めました。
リンゴがとても美味しかったです！
鳥たちは満腹になるまで食べました。
鳥たちは幸せいっぱいで飛び立ちました。
)";
std::string テスト出力;
UllmLama2RunConfig run_config;
UllmLalama2RunConfigInit(&run_config);
run_config.checkpoint_path = "ullm/tinystories15M.bin";
run_config.tokenizer_path = "ullm/tokenizer.bin";
run_config.prompt = "鳥がさえずる。彼らはどこへ行くのですか？」
run_config.output_callback = 出力ハンドラー;
run_config.cookie = &test_output;
UllmLalama2State 状態。
UllmStatus ステータス = UllmLama2Init(&run_config, &state);
ASSERT_EQ(ステータス, ULLM_STATUS_OK);
status = UllmLama2Generate(&run_config, &state);
EXPECT_EQ(ステータス, ULLM_STATUS_OK);
UllmLalama2Deinit(&state);
EXPECT_EQ(テスト出力, 期待テスト出力);
内部構造
UllmStatus UllmLama2Init(const UllmLama2RunConfig* 構成,
UllmLalama2State* 状態) {
memset(状態, 0, sizeof(UllmLalama2State));
ULLM_RETURN_IF_ERROR(UllmLama2ValidateConfig(config));
ULLM_RETURN_IF_ERROR(UllmLama2BuildTransformer(config, state));
ULLM_RETURN_IF_ERROR(UllmLama2BuildTokenizer(config, state));
ULLM_RETURN_IF_ERROR(UllmLama2BuildSampler(config, state));
ULLM_STATUS_OK を返します。
}
UllmStatus UllmLalama2Generate(const UllmLalama2RunConfig* config,
UllmLalama2State* 状態) {
// '\0'、?BOS、?EOS の場合は +3
size_t プロンプトトークンサイズ = (strlen(config->prompt) + 3) * sizeof(int);
int* プロンプトトークン = UllmMemoryAlloc(プロンプトトークンサイズ);
if (prompt_tokens == NULL) {
ULOGE("プロンプト トークンの割り当てに失敗しました");
ULLM_STATUS_OOM を返します。
PPC への移植
エンディアンネス: モデルのチェックポイント/トークンを変換します。

リトルエンディアンからビッグエンディアンへのナイザー
アライメント、16 バイト: メモリ マッピングではなく重みをメモリにコピーします。
if (ステータス != ULLM_STATUS_OK) {
ULOGE("チェックポイントのマッピングに失敗しました: %s", UllmStatusToString(status));
クリーンアップに移動します。
}
if ((file_size % 4) != 0) {
ULOGE("チェックポイントのサイズは 4 の倍数である必要があります");
ステータス = ULLM_STATUS_INVALID_ARGUMENT;
クリーンアップに移動します。
}
for (uint64_t i = 0; i < file_size; i+= 4) {
const uint32_t le_value = *(const uint32_t*)&file_ptr[i];
const uint32_t be_value = __builtin_bswap32(le_value);
fwrite(&be_value, sizeof(uint32_t), 1, stdout);
}
モデル
aarossig@lithium:~/projects/ullm$ bazel ビルド ツール:ullm -c opt
aarossig@lithium:~/projects/ullm$ time bazel-bin/tools/ullm -c tinystories110M.bin -t tokenizer.bin -p 「足の速い茶色のキツネが飛び降りました。どこへ行ったのですか?」
素早い茶色のキツネが飛び上がりました。彼はどこへ行ったのですか？彼は一日中森を走り回っていました。彼は何か特別なものを探していました。
突然、大きな木が見えました。彼は立ち止まって顔を上げた。彼は木の中に大きな茶色の鳥を見た。鳥は美しい歌を歌っていました。
キツネはとても幸せでした。彼はその鳥にもっと近づきたかった。彼は飛び跳ねて鳥に近づこうとした。しかし、鳥はあまりにも高いところにいました。
キツネは悲しんでいました。彼はその鳥と友達になりたかった。周りを見回すと、大きな茶色の岩が見えました。彼は岩に飛び乗って鳥を見上げた。
鳥はキツネを見て、彼のところに飛んでいきました。キツネはとても幸せでした。彼とその鳥は友達になりました。彼らは毎日森で一緒に遊びました。
I ullm.llama2: 完了: 6.91 トークン/秒
実質0分26.511秒
ユーザー 0分26.019秒
sys 0m0.472s クリアする目標は 26.5 秒、または 1 秒あたり 6.91 トークンです (ネタバレ注意: 及ばない)。推論シードはランダム化されていないため、生成される出力は実行ごとに一定になります。この測定は、何度も実行しても再現可能です。
だから何

この同じコードを 1.5GHz PowerPC 7447B プロセッサを搭載した PowerBook G4 で実行するとどうなるでしょうか? aarossig@titanium:~/Projects/ullm$ make && time ./out/ullm.elf -c ../models/tinystories110M_be.bin -t ../models/tokenizer_be.bin -p "足の速い茶色のキツネが飛び降りました。どこへ行ったのですか?"
素早い茶色のキツネが飛び上がりました。彼はどこへ行ったのですか？彼は一日中森を走り回っていました。彼は何か特別なものを探していました。
突然、大きな木が見えました。彼は立ち止まって顔を上げた。彼は木の中に大きな茶色の鳥を見た。鳥は美しい歌を歌っていました。
キツネはとても幸せでした。彼はその鳥にもっと近づきたかった。彼は飛び跳ねて鳥に近づこうとした。しかし、鳥はあまりにも高いところにいました。
キツネは悲しんでいました。彼はその鳥と友達になりたかった。彼は周りを見回して、あるものを見つけました。
[切り捨てられた]
注: このブログのメンバーのみがコメントを投稿できます。
▼
2025年
(
1
)
▼
3月
(
1
)
異なる考え方、ゆっくり考える: 強力な LLM
►
2022年
(
1
)
►
9月
(
1
)
►
2017年
(
2
)
►
9月
(
1
)
►
2013年
(
21
)
►
12月
(
3
)
X11 の下位互換性の証
私は最近、Hewlett Packard 1670A Deep Memory Logic Analyzer を購入し、ついにそれを起動する機会を得ました。このユニットの歴史は 1992 年に遡ります。
ARM ベアメタル プログラミング
組み込みシステム プログラミングは、数年前から私の情熱の対象となっています。私はプロセッサをオンラインにしてそれを作るのが本当に楽しいです...
AVR VGA の生成
私はずっと VGA ジェネレーターを作りたいと思っていました。それはしばらくの間、私の個人的な目標でした。過去数週間、私は色々なことを試してきました...
長距離電動ロングボード
私はニューヨークに住んでいた今年の 8 月に Boosted Board を購入しました。素晴らしい買い物でした。乗っていて楽しいし、乗ってみたいです...
違う考え方、ゆっくり考える: PowerPC Mac 上の LLM
何かがある

古いハードウェアに新しい命を吹き込むことに非常に満足しています。ビンテージ コンピューティングは私のお気に入りの趣味の 1 つです。チャ...
Flir Lepton 熱画像センサー + Gameduino 2
私は最近、Flir Lepton サーマル イメージング センサーを入手し、先週、空き時間にそれらをオンラインに接続しました...
PortL2 - ポータブル電気自動車充電器
過去 1 年間、私はキャデラック ELR の電気航続距離を延長するというアイデアを検討してきました。ラーを建てたかったのですが...
組み込み HTTP サーバーを楽しむ
ここ 1 週間、私は組み込みの HTTP サーバーに気を取られていました。私は現在、インターネット組み込み...というタイトルの非常に興味深いコースを受講しています。
ビンテージ インテル MCS-48 マイクロコントローラーのプログラミング
私は数年前からパーツコレクションの中に、さまざまな興味深いビンテージコンポーネントが入った箱を持っています。最も顕著なものは...
ガラス基板の作成
私は CNLohr に触発されて、独自のガラス基板を作成しました。レイアウトが簡単なので、2D LED マトリックスを作成することにしました。

## Original Extract

There is something incredibly satisfying about breathing new life into old hardware. Vintage computing is one of my favorite hobbies. The ch...

The Resistor Network: Thinking Different, Thinking Slowly: LLMs on a PowerPC Mac
Monday, March 24, 2025
Thinking Different, Thinking Slowly: LLMs on a PowerPC Mac
Introduce a status return, remove all calls to exit
Abstract file access to simplify status handling
Abstract malloc / free for some simple debug/analysis
Replace the 512 byte static LUT for single character strings
Fix a few warnings when compiling with -Wall
// The runtime config for the inference operation.
typedef struct {
// The prompt to generate a response to.
const char* prompt;
// The path to the checkpoint file.
const char* checkpoint_path;
// The path to the tokenizer file.
const char* tokenizer_path;
// Model configuration.
float temperature;
float topp;
unsigned int steps;
// The source of entropy.
uint64_t rng_seed;
// The callback and context for generated output.
void (*output_callback)(const char* token, void* cookie);
void* cookie;
} UllmLlama2RunConfig;
// The runtime state for the inference engine.
typedef struct {
UllmFileHandle checkpoint_file;
UllmLlama2Transformer transformer;
UllmLlama2Tokenizer tokenizer;
UllmLlama2Sampler sampler;
} UllmLlama2State;
The library expects two inputs: a const config which supplies details such as model paths, rng seed and token output callbacks, as well as state which keeps track of loaded weights, temporary buffers, tokenizer state and more.
The resulting API is exceedingly simple to test and easy to build command-line interface tools around.
Callback-Based Output & Testing
void OutputHandler(const char* token, void* cookie) {
std::string* test_output = static_cast<std::string*>(cookie);
test_output->append(token);
}
TEST(UllmLlama2, Stories15M) {
const std::string expected_test_output = R"(The birds chirp. Where do they go?
The birds flew around the sky, looking for something to do.
The birds saw a big tree and flew over to it.
The birds saw a big, red apple on the ground. It looked delicious.
The birds flew down and picked up the apple.
The birds flew back up to the tree and started to eat the apple.
The apples were so delicious!
The birds ate until they were full.
The birds flew away, happy and full.
)";
std::string test_output;
UllmLlama2RunConfig run_config;
UllmLlama2RunConfigInit(&run_config);
run_config.checkpoint_path = "ullm/tinystories15M.bin";
run_config.tokenizer_path = "ullm/tokenizer.bin";
run_config.prompt = "The birds chirp. Where do they go?";
run_config.output_callback = OutputHandler;
run_config.cookie = &test_output;
UllmLlama2State state;
UllmStatus status = UllmLlama2Init(&run_config, &state);
ASSERT_EQ(status, ULLM_STATUS_OK);
status = UllmLlama2Generate(&run_config, &state);
EXPECT_EQ(status, ULLM_STATUS_OK);
UllmLlama2Deinit(&state);
EXPECT_EQ(test_output, expected_test_output);
} Internals
UllmStatus UllmLlama2Init(const UllmLlama2RunConfig* config,
UllmLlama2State* state) {
memset(state, 0, sizeof(UllmLlama2State));
ULLM_RETURN_IF_ERROR(UllmLlama2ValidateConfig(config));
ULLM_RETURN_IF_ERROR(UllmLlama2BuildTransformer(config, state));
ULLM_RETURN_IF_ERROR(UllmLlama2BuildTokenizer(config, state));
ULLM_RETURN_IF_ERROR(UllmLlama2BuildSampler(config, state));
return ULLM_STATUS_OK;
}
UllmStatus UllmLlama2Generate(const UllmLlama2RunConfig* config,
UllmLlama2State* state) {
// +3 for '\0', ?BOS, ?EOS
size_t prompt_tokens_size = (strlen(config->prompt) + 3) * sizeof(int);
int* prompt_tokens = UllmMemoryAlloc(prompt_tokens_size);
if (prompt_tokens == NULL) {
ULOGE("Failed to allocate prompt tokens");
return ULLM_STATUS_OOM;
} Porting to PPC
Endianness: convert model checkpoints/tokenizer from little- to big- endian
Alignment, 16-byte: copy weights into memory rather than memory mapping
if (status != ULLM_STATUS_OK) {
ULOGE("Failed to map checkpoint: %s", UllmStatusToString(status));
goto cleanup;
}
if ((file_size % 4) != 0) {
ULOGE("Checkpoint size must be a multiple of 4");
status = ULLM_STATUS_INVALID_ARGUMENT;
goto cleanup;
}
for (uint64_t i = 0; i < file_size; i+= 4) {
const uint32_t le_value = *(const uint32_t*)&file_ptr[i];
const uint32_t be_value = __builtin_bswap32(le_value);
fwrite(&be_value, sizeof(uint32_t), 1, stdout);
}
Models
aarossig@lithium:~/projects/ullm$ bazel build tools:ullm -c opt
aarossig@lithium:~/projects/ullm$ time bazel-bin/tools/ullm -c tinystories110M.bin -t tokenizer.bin -p "The quick brown fox jumped. Where did he go?"
The quick brown fox jumped. Where did he go? He had been running around the forest all day. He was looking for something special.
Suddenly, he saw a big tree. He stopped and looked up. He saw a big, brown bird in the tree. The bird was singing a beautiful song.
The fox was so happy. He wanted to get closer to the bird. He jumped up and down, trying to reach the bird. But the bird was too high up.
The fox was sad. He wanted to be friends with the bird. He looked around and saw a big, brown rock. He jumped on the rock and looked up at the bird.
The bird saw the fox and flew down to him. The fox was so happy. He and the bird became friends. They played together in the forest every day.
I ullm.llama2: Complete: 6.91 token/s
real 0m26.511s
user 0m26.019s
sys 0m0.472s The target to beat is 26.5 seconds or 6.91 tokens per second (spoiler alert: not even close). The inference seed is not randomized, so the generated output will be constant from run to run. This measurement is repeatable over many runs.
So what happens if we take this identical code and run it on a PowerBook G4 with a 1.5GHz PowerPC 7447B processor? aarossig@titanium:~/Projects/ullm$ make && time ./out/ullm.elf -c ../models/tinystories110M_be.bin -t ../models/tokenizer_be.bin -p "The quick brown fox jumped. Where did he go?"
The quick brown fox jumped. Where did he go? He had been running around the forest all day. He was looking for something special.
Suddenly, he saw a big tree. He stopped and looked up. He saw a big, brown bird in the tree. The bird was singing a beautiful song.
The fox was so happy. He wanted to get closer to the bird. He jumped up and down, trying to reach the bird. But the bird was too high up.
The fox was sad. He wanted to be friends with the bird. He looked around and saw a
[truncated]
Note: Only a member of this blog may post a comment.
▼
2025
(
1
)
▼
March
(
1
)
Thinking Different, Thinking Slowly: LLMs on a Pow...
►
2022
(
1
)
►
September
(
1
)
►
2017
(
2
)
►
September
(
1
)
►
2013
(
21
)
►
December
(
3
)
A Testament to X11 Backwards Compatibility
I recently scored a Hewlett Packard 1670A Deep Memory Logic Analyzer and I finally had a chance to fire it up. This unit dates back to 1992 ...
ARM Bare Metal Programming
Embedded systems programming has been a passion of mine for a couple of years now. I really enjoy bringing a processor online and making it ...
AVR VGA Generation
I have always wanted to create a VGA generator. It has been a personal goal of mine for some time. Over the past couple of weeks I toyed wit...
Long Range Electric Longboarding
I bought a Boosted Board back in August of this year while living in New York City. It was an awesome purchase. It is fun to ride and I cou...
Thinking Different, Thinking Slowly: LLMs on a PowerPC Mac
There is something incredibly satisfying about breathing new life into old hardware. Vintage computing is one of my favorite hobbies. The ch...
Flir Lepton Thermal Imaging Sensor + Gameduino 2
I recently got my hands on a pair of Flir Lepton thermal imaging sensors and have spent the last week bringing them online in my spare time...
PortL2 - Portable Electric Vehicle Charger
Over the past year I have been tossing around the idea of extending the electric range of my Cadillac ELR . I have wanted to build a lar...
Fun with Embedded HTTP Servers
I have been distracted for the past week by an embedded HTTP server. I am currently taking a very exciting course entitled Internet Embedded...
Programming the Vintage Intel MCS-48 Microcontrollers
I have had a box in my parts collection for a few years that contains a variety of interesting vintage components. The most prominent are th...
Creating a Glass Circuit Board
I was inspired by CNLohr to create my own glass circuit board. I decided to create a 2D LED matrix because the layout is simple enough to f...
