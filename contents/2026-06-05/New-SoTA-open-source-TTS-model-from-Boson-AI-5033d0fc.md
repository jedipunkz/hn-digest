---
source: "https://huggingface.co/bosonai/higgs-audio-v3-tts-4b"
hn_url: "https://news.ycombinator.com/item?id=48409001"
title: "New SoTA open source TTS model from Boson AI"
article_title: "bosonai/higgs-audio-v3-tts-4b · Hugging Face"
author: "silinmeng"
captured_at: "2026-06-05T07:36:58Z"
capture_tool: "hn-digest"
hn_id: 48409001
score: 3
comments: 0
posted_at: "2026-06-05T07:02:43Z"
tags:
  - hacker-news
  - translated
---

# New SoTA open source TTS model from Boson AI

- HN: [48409001](https://news.ycombinator.com/item?id=48409001)
- Source: [huggingface.co](https://huggingface.co/bosonai/higgs-audio-v3-tts-4b)
- Score: 3
- Comments: 0
- Posted: 2026-06-05T07:02:43Z

## Translation

タイトル: Boson AI の新しい SoTA オープンソース TTS モデル
記事タイトル: bosonai/higgs-audio-v3-tts-4b · 抱きしめる顔
説明: 私たちは、オープンソースとオープン サイエンスを通じて人工知能を進歩させ、民主化する旅の途中にあります。

記事本文:
bosonai/higgs-audio-v3-tts-4b · 抱きしめる顔
ハグ顔モデル
bosonai / higgs-audio-v3-tts-4b いいね 61 Boson AI 419 をフォロー
Text-to-Speech Transformers Safetensors 100 言語 higgs_multimodal_qwen3 text-generation speech-generation voice-agentexpressive-speech controllable-tts multilingual-tts License: boson-higgs-audio-v3-research-and-non-commercial-license モデル カード ファイル ファイルとバージョン xet コミュニティ 1 デプロイ バケットにコピー 新しい このモデルを使用する 使用手順bosonai/higgs-audio-v3-tts-4b とライブラリ、推論プロバイダー、ノートブック、ローカル アプリ。これらのリンクに従って開始してください。
Transformers Transformers で bosonai/higgs-audio-v3-tts-4b を使用する方法:
# パイプラインを高レベルのヘルパーとして使用する
変圧器からのインポートパイプライン
Pipe = Pipeline("text-to-speech", model="bosonai/higgs-audio-v3-tts-4b") # モデルを直接ロードする
トランスフォーマーから AutoModelForSeq2SeqLM をインポート
モデル = AutoModelForSeq2SeqLM.from_pretrained("bosonai/higgs-audio-v3-tts-4b", dtype="auto")
Higgs Audio v3 TTS サポート言語 WER/CER 5 未満 — 洗練された製品品質 (85)
WER/CER 5 ～ 10 — 使用可能ですが、あまり洗練されていません (17)
評価ベンチマーク 多言語音声クローン
Higgs Audio v3 TTS は音声チャット用に構築されており、単に読み上げるだけでなく、音声で話します。ゼロショットの音声クローン作成と、感情、スタイル、韻律、ポーズ、音響効果のインライン制御により、モデルの応答を 100 以上の言語にわたる表現力豊かな会話音声に変換します。
Boson Higgs Audio v3 Research and Non-Commercial License に基づいて研究および非営利目的でリリースされています。運用、ホストされた API、または収益を生み出す使用には、別の商用ライセンスが必要です。禁止: 同意のない音声複製、なりすまし、詐欺、選挙詐欺、生体認証監視、その他の違法な行為

せ。
ヒッグス自己回帰デコーダは、インターリーブされたテキストとオーディオのトークンを消費します。オーディオは、ヒッグス トークナイザーによって 25 fps で 8 つのコードブックにエンコードされ、遅延パターンによってずらされ、マルチ コードブック融合埋め込みを通じてバックボーンの隠れ状態にマッピングされます。出力コードは、マルチ コードブック融合ヘッドを通過し、遅延解除されて、波形にデコードされます。
このモデルは 102 言語で 1 桁の WER/CER に達し、2 つの層に分かれています。
WER/CER が 5 未満 — 洗練された製品品質 (85)
🇿🇦 アフリカーンス語 · 🇸🇦🇪🇬 アラビア語 · 🇦🇲 アルメニア語 · 🇮🇳 アッサム語 · 🇪🇸 アストゥリアス語 · 🇦🇿 アゼルバイジャン語 · 🇷🇺 バシキール語 · 🇪🇸 バスク語 · 🇧🇾 ベラルーシ語 · 🇧🇩🇮🇳 ベンガル語 · 🇧🇦 ボスニア語 · 🇧🇬 ブルガリア語 · 🇪🇸 カタロニア語 · 🇵🇭 セブアノ語 · 🇮🇶 中央クルド語 · 🇨🇳 中国語 · 🇭🇷 クロアチア語 · 🇨🇿 チェコ語 · 🇩🇰 デンマーク語 · 🇳🇱🇧🇪 オランダ語 · 🇷🇺 東マリ語 · 🇺🇸🇬🇧🇦🇺 英語 · 🌐 エスペラント語 · 🇪🇪 エストニア語· 🇫🇮フィンランド語 · 🇫🇷🇨🇦フランス語 · 🇪🇸ガリシア語 · 🇬🇪グルジア語 · 🇩🇪🇦🇹ドイツ語 · 🇬🇷ギリシャ語 · 🇮🇳グジャラート語 · 🇭🇹ハイチクレオール語 · 🇳🇬 ハウサ語 · 🇮🇱 ヘブライ語 · 🇮🇳 ヒンディー語 · 🇭🇺 ハンガリー語 · 🇮🇩 インドネシア語 · 🇮🇹 イタリア語 · 🇯🇵 日本語 · 🇮🇩 ジャワ語 · 🇮🇳 カンナダ語 · 🇰🇿 カザフ語 · 🇰🇷 韓国語 · 🇷🇼 ルワンダ語 · 🇰🇬 キルギス語 · 🇱🇻 ラトビア語 · 🇨🇩 リンガラ語 · 🇱🇹 リトアニア語 · 🇰🇪 ルオ語 · 🇲🇰 マケドニア語 · 🇲🇾🇮🇩マレー語 · 🇮🇳マラヤーラム語 · 🇲🇹マルタ語 · 🇳🇿マオリ語 · 🇮🇳マラーティー語 · 🇲🇳モンゴル語 · 🇳🇵ネパール語 · 🇳🇴ノルウェー語 · 🇫🇷 オック語 · 🇮🇷🇦🇫 ペルシャ

n · 🇵🇱 ポーランド語 · 🇵🇹🇧🇷 ポルトガル語 · 🇷🇴 ルーマニア語 · 🇷🇺 ロシア語 · 🇿🇦 セペディ語 · 🇷🇸 セルビア語 · 🇿🇼 ショナ語 · 🇸🇰 スロバキア語 · 🇸🇮スロベニア語 · 🇪🇸🇲🇽スペイン語 · 🇹🇿🇰🇪スワヒリ語 · 🇸🇪スウェーデン語 · 🇵🇭タガログ語 · 🇹🇯タジク語 · 🇮🇳🇱🇰タミル語 · 🇮🇳 テルグ語 · 🇹🇭 タイ語 · 🇹🇷 トルコ語 · 🇺🇦 ウクライナ語 · 🇵🇰🇮🇳 ウルドゥー語 · 🇨🇳 ウイグル語 · 🇺🇿 ウズベク語 · 🇻🇳 ベトナム語 · 🇿🇦 コサ語 · 🇿🇦 ズールー語
WER/CER 5 ～ 10 — 使用可能ですが、あまり洗練されていません (17)
🇦🇱 アルバニア人 · 🇲🇼🇿🇲 チチェワ/ニャンジャ · 🇮🇳🇵🇰 東パンジャブ人 · 🇺🇬 ガンダ人 · 🇮🇸 アイスランド人 · 🇮🇪 アイルランド人 · 🇩🇿 カビル · 🇨🇻カブベルディヌ語 · 🇰🇪カンバ語 · 🇻🇦ラテン語 · 🇱🇺ルクセンブルク語 · 🇪🇹🇰🇪オロモ語 · 🇦🇫🇵🇰パシュトゥー語 · 🇵🇰🇮🇳シンディ語 · 🇸🇴ソマリ語 · 🇦🇴ウンブンドゥ語 · 🇬🇧ウェールズ語
すべてのタグは <|category:value|> 構文に従い、発話中に挿入できます。
感情 — 高揚、楽しさ、熱意、決意、誇り、満足感、愛情、安堵、熟考、混乱、驚き、畏怖、憧れ、興奮、怒り、恐怖、嫌悪、苦味、悲しみ、恥、無力感
スタイル — 歌う、叫ぶ、ささやく
効果音 — 咳、笑い、泣き、叫び、げっぷ、鼻歌、ため息、鼻をすする、くしゃみ
各トークンの直後に対応するオノマトペを組み合わせます。
韻律
速度 —speed_very_slow 、speed_slow 、speed_fast 、speed_very_fast
ピッチ —ピッチ_低、ピッチ_高
配信 —expressive_high 、expressive_low
私たちは、公開多言語 TTS スイートと社内の 111 言語の Higgs-Multilingual セットで Higgs Audio v3 TTS を評価し、両方の通信をカバーします。

および低リソース言語で。
WER / CER (↓、×100) 各ベンチマークの言語セット全体でマクロ平均。低いほど良いです。太字は行ごとに最適であることを示します。すべての数値は、元のメトリクスと正規化によりエンドツーエンドで再現可能です。
カテゴリごとの勝率 (↑) — BASELINE 行に対する好みを判断します。太字は列ごとの最高勝率を示します。公平に比較​​するために、すべてのモデルはプロンプトごとに同じリファレンス オーディオを共有し、ベンチマーク テキストをそのまま実行します (インライン コントロール タグは挿入されません)。
このリポジトリの重みを SGLang-Omni と組み合わせます。SGLang-Omni は、マルチコードブックのデコードと同じインライン タグ制御のための連続バッチ処理を備えた実稼働サービス スタックです。 Higgs TTS クックブックでは、インストール、サーバーの起動、リクエストの例、完全な API リファレンスについて説明しています。
詳細については、Higgs TTS クックブックを参照してください。
インストールして提供する
docker pull lmsysorg/sglang-omni:dev
docker run -it --gpus all --shm-size 32g --ipc host --network host --privileged \
lmsysorg/sglang-omni:dev /bin/zsh
git clone git@github.com:sgl-project/sglang-omni.git && cd sglang-omni
uv venv .venv -p 3.12 && ソース .venv/bin/activate
uv pip install -v -e 。
エクスポート HF_TOKEN=hf_xxxxxxxxxxxxxxxx
hf ダウンロード bosonai/higgs-audio-v3-tts-4b
sgl-オムニサーブ \
--model-path bosonai/higgs-audio-v3-tts-4b \
--ポート 8000
ゼロショット合成
curl -X POST http://localhost:8000/v1/audio/speech \
-H "コンテンツ タイプ: application/json" \
-d '{"input": "こんにちは、お元気ですか?"}' \
--output 出力.wav
音声クローン
参照トランスクリプト ( text ) を提供すると、クローン作成の忠実度が大幅に向上します。
インポートリクエスト
resp = リクエスト.post(
"http://localhost:8000/v1/audio/speech" ,
json={
"input" : "南カリフォルニアの日差しをお楽しみください。" 、
「参照」: [{
"audio_path" : "ref.wav" ,
"text" : "やあ、アダム、ここにいるよ。さあ、やってみよう。

リアルに感じられ、人間味があり、いつでもつながるものを作成してください。」
}]、
"温度" : 0.8 、 "top_k" : 50 、 "max_new_tokens" : 1024 、
}、
)
open ( "output.wav" , "wb" ) を f として使用:
f.write(それぞれのコンテンツ)
ストリーミング (サーバー送信イベント)
"stream": true を設定すると、base64 でエンコードされた WAV チャンクがボコーダーから出力されるときに受信されます (最初のオーディオまでの時間が 1 秒未満)。各イベントは audio.data (base64 WAV バイト) を伝送します。ターミナル イベントには、finish_reason: "stop" と使用状況メタデータが含まれます。
インポートリクエスト、base64、json
リクエストあり.post(
"http://localhost:8000/v1/audio/speech" ,
json={ "input" : "信託資金を早めに銀行に届けてください。" , "ストリーム" : True },
ストリーム= True 、
) として、( "output.wav" 、 "wb" ) を f として開きます:
resp.iter_lines() の行:
line でない場合、または line.startswith( b"data: " ) または line == b"data: [DONE]" の場合:
続ける
イベント = json.loads(line[6:])
ifevent.get( "finish_reason" ) == "停止" :
休憩
audio = events.get( "audio" ) または {}
if audio.get( "データ" ):
f.write(base64.b64decode(audio[ "data" ]))
インラインコントロールトークン
<|emotion:…|> 、 <|style:…|> 、 <|prosody:…|> 、および <|sfx:…|> トークンを input に直接埋め込みます。 2 つのルール:
まずは配信トークン。感情、スタイル、韻律のスピード/ピッチ/表現力豊かなトークンがターン全体を形成します。これらを入力の開始時に配置します。位置トークン ( <|prosody:pause|> 、 <|prosody:long_pause|> 、 <|sfx:…|> ) は、発火した場所に正確にインライン化されます。
すべての <|sfx:…|> をオノマトペと組み合わせます。例： <|sfx:Laugh|>はは、<|sfx:sigh|>うーん、<|sfx:sneeze|>アチョー。書かれた音は、モデルに効果を実現するための音響的な合図を与えます。
例 — 娯楽 + 笑い:
curl -X POST http://localhost:8000/v1/audio/speech \
-H "コンテンツ タイプ: application/json" \
-d '{"input": "<|emotion:amusement|><|prosody:expressive_high|>待て、待て、それはキだった

笑える。 <|sfx:笑い|>ふふ、いや、真剣に言うと、その準備ができていませんでした。"}' \
--output 出力.wav
スループット
Seed-TTS EN のスループット (フルセット、実行ごとに N=1088)。クライアント --Higgs サーバーに対する最大同時実行スイープ ( max_running_requests=16 、 bf16、CUDA グラフ オン)。各行は 3 回の実行の平均です。ハードウェア: 1× H100。
同時実行性 — 実行中のクライアントリクエストの最大数 ( --max-concurrency )。
スループット (req/s) — 完了したリクエストをベンチマークの合計実時間で割ったもの。
平均遅延 — リクエストごとの平均エンドツーエンド時間 (送信から完全な応答の受信まで)。
RTF (リクエストごと) — リクエストごとに生成されたオーディオ継続時間に対する処理時間の平均比率 (<1 はリアルタイムより高速です)。
audio_s/s — 生成されたオーディオの合計秒数をベンチマークの合計実時間で割ったもの。
結果を再現するには、このスクリプトの指示に従ってください。
ゼロオペレーションのデプロイには、Boson AI API を使用します。
@misc{bosonai_higgs_audio_tts_v3_2026,
title = {Higgs Audio v3 TTS: Boson AI による音声 AI のための会話型スピーチ},
著者 = {ボソン AI}、
年 = {2026}、
howpublished = {https://huggingface.co/bosonai/higgs-audio-v3-tts-4b}、
}
ライセンス
Boson Higgs Audio v3 Research および非営利ライセンス — 「ライセンス」を参照してください。
","pad_token":"<|endoftext|>","unk_token":null},"chat_template_jinja":"{%- if tools %}\n {{- '<|im_start|>system\\n' }}\n {%- ifmessages[0]['role'] == 'system' %}\n {{-messages[0]['content'] }}\n {%- else %}\n {{- 「あなたは役に立つアシスタントです。」 }}\n {%- endif %}\n {{- \"\\n\\n# Tools\\n\\nユーザーのクエリを支援するために 1 つ以上の関数を呼び出すことができます。\\n\\n<tools></tools> XML タグ内で関数シグネチャが提供されます:\\n<tools>\" }}\n {%- for tools %}\n {{- \"\\n\" }}\n {{-ツール | tojson }}\n {%- endfor %}\n {{- \"\\n</tools>\\n\\n関数ごと

呼び出し、<tool_call></tool_call> XML タグ内の関数名と引数を持つ JSON オブジェクトを返します:\\n<tool_call>\\n{\\\"name\\\": <function-name>, \\\"arguments\\\": <args-json-object>}\\n</tool_call><|im_end|>\\n\" }}\n{%- else %}\n {%- if messages[0]['role'] == 'system' %}\n {{- '<|im_start|>system\\n' +messages[0]['content'] + '<|im_end|>\\n' }}\n {%- else %}\n {{- '<|im_start|>system\\nあなたは役に立つアシスタントです。<|im_end|>\\n' }}\n {%- endif %}\n{%- endif %}\n{%- メッセージ内のメッセージ %}\n {%- if (message.role == \"user\") または (message.role == \"system\" で、loop.first ではない) または (message.role == \"assistant\" で、message.tool_calls ではない) %}\n {{- '<|im_start|>' + message.role + '\\n' + message.content + '<|im_end|>' + '\\n' }}\n {%- elif message.role == \"assistant\" %}\n {{- '<|im_start|>' + message.role }}\n {%- if message.content %}\n {{- '\\n' + message.content }}\n {%- endif %}\n {%- for tool_call in message.tool_calls %}\n {%- if tool_call.function が定義されている場合 %}\n {%- set tool_call = tool_call.function %}\n {%- endif %}\n {{- '\\n<tool_call>\\n{\"name\": \"' }}\n {{- tool_call.name }}\n {{- '\", \"arguments\": ' }}\n {{- tools_call.arguments | tojson }}\n {{- '}\\n</tool_call>' }}\n {%- endfor %}\n {{- '<|im_end|>\\n' }}\n {%- elif message.role == \"tool\" %}\n {%- if (loop.index0 == 0) または(messages[loop.index0 - 1].role != \"tool\") %}\n {{- '<|im_start|>user' }}\n {%- endif %}\n {{- '\
[切り捨てられた]
推論プロバイダー 新しい Text-to-Speech このモデルは、どの推論プロバイダーによっても展開されていません。 🙋 プロバイダーのサポートを依頼する bosonai/higgs-audio-v3-tts-4b のモデル ツリー
Finetunes 1 モデル Quantizations 4 モデル bosonai/higgs-audio-v3-tts-4b 2 を使用したスペース
bosonai/higgs-audio-v3-tts-4b を含むコレクション

## Original Extract

We’re on a journey to advance and democratize artificial intelligence through open source and open science.

bosonai/higgs-audio-v3-tts-4b · Hugging Face
Hugging Face Models
bosonai / higgs-audio-v3-tts-4b like 61 Follow Boson AI 419
Text-to-Speech Transformers Safetensors 100 languages higgs_multimodal_qwen3 text-generation speech-generation voice-agent expressive-speech controllable-tts multilingual-tts License: boson-higgs-audio-v3-research-and-non-commercial-license Model card Files Files and versions xet Community 1 Deploy Copy to bucket new Use this model Instructions to use bosonai/higgs-audio-v3-tts-4b with libraries, inference providers, notebooks, and local apps. Follow these links to get started.
Transformers How to use bosonai/higgs-audio-v3-tts-4b with Transformers:
# Use a pipeline as a high-level helper
from transformers import pipeline
pipe = pipeline("text-to-speech", model="bosonai/higgs-audio-v3-tts-4b") # Load model directly
from transformers import AutoModelForSeq2SeqLM
model = AutoModelForSeq2SeqLM.from_pretrained("bosonai/higgs-audio-v3-tts-4b", dtype="auto")
Higgs Audio v3 TTS Supported Languages WER/CER under 5 — polished, production-quality (85)
WER/CER between 5 and 10 — usable, but less polished (17)
Evaluation Benchmarks Multilingual Voice Clone
Higgs Audio v3 TTS is built for voice chat: it speaks, not just reads . It turns model responses into expressive conversational speech across 100+ languages , with zero-shot voice cloning and inline control over emotion, style, prosody, pauses, and sound effects.
Released for research and non-commercial use under the Boson Higgs Audio v3 Research and Non-Commercial License . Production, hosted APIs, or revenue-generating use requires a separate commercial license. Prohibited: voice cloning without consent, impersonation, fraud, election deception, biometric surveillance, or any unlawful use.
Higgs autoregressive decoder consumes interleaved text and audio tokens. Audio is encoded by the Higgs Tokenizer into 8 codebooks at 25 fps, staggered via a delay pattern , then mapped to backbone hidden states through a multi-codebook fused embedding . Output codes pass through a multi-codebook fused head , are de-delayed, and decoded back to waveform.
The model reaches single-digit WER/CER on 102 languages , which split into two tiers.
WER/CER under 5 — polished, production-quality (85)
🇿🇦 Afrikaans · 🇸🇦🇪🇬 Arabic · 🇦🇲 Armenian · 🇮🇳 Assamese · 🇪🇸 Asturian · 🇦🇿 Azerbaijani · 🇷🇺 Bashkir · 🇪🇸 Basque · 🇧🇾 Belarusian · 🇧🇩🇮🇳 Bengali · 🇧🇦 Bosnian · 🇧🇬 Bulgarian · 🇪🇸 Catalan · 🇵🇭 Cebuano · 🇮🇶 Central Kurdish · 🇨🇳 Chinese · 🇭🇷 Croatian · 🇨🇿 Czech · 🇩🇰 Danish · 🇳🇱🇧🇪 Dutch · 🇷🇺 Eastern Mari · 🇺🇸🇬🇧🇦🇺 English · 🌐 Esperanto · 🇪🇪 Estonian · 🇫🇮 Finnish · 🇫🇷🇨🇦 French · 🇪🇸 Galician · 🇬🇪 Georgian · 🇩🇪🇦🇹 German · 🇬🇷 Greek · 🇮🇳 Gujarati · 🇭🇹 Haitian Creole · 🇳🇬 Hausa · 🇮🇱 Hebrew · 🇮🇳 Hindi · 🇭🇺 Hungarian · 🇮🇩 Indonesian · 🇮🇹 Italian · 🇯🇵 Japanese · 🇮🇩 Javanese · 🇮🇳 Kannada · 🇰🇿 Kazakh · 🇰🇷 Korean · 🇷🇼 Kinyarwanda · 🇰🇬 Kyrgyz · 🇱🇻 Latvian · 🇨🇩 Lingala · 🇱🇹 Lithuanian · 🇰🇪 Luo · 🇲🇰 Macedonian · 🇲🇾🇮🇩 Malay · 🇮🇳 Malayalam · 🇲🇹 Maltese · 🇳🇿 Māori · 🇮🇳 Marathi · 🇲🇳 Mongolian · 🇳🇵 Nepali · 🇳🇴 Norwegian · 🇫🇷 Occitan · 🇮🇷🇦🇫 Persian · 🇵🇱 Polish · 🇵🇹🇧🇷 Portuguese · 🇷🇴 Romanian · 🇷🇺 Russian · 🇿🇦 Sepedi · 🇷🇸 Serbian · 🇿🇼 Shona · 🇸🇰 Slovak · 🇸🇮 Slovene · 🇪🇸🇲🇽 Spanish · 🇹🇿🇰🇪 Swahili · 🇸🇪 Swedish · 🇵🇭 Tagalog · 🇹🇯 Tajik · 🇮🇳🇱🇰 Tamil · 🇮🇳 Telugu · 🇹🇭 Thai · 🇹🇷 Turkish · 🇺🇦 Ukrainian · 🇵🇰🇮🇳 Urdu · 🇨🇳 Uyghur · 🇺🇿 Uzbek · 🇻🇳 Vietnamese · 🇿🇦 Xhosa · 🇿🇦 Zulu
WER/CER between 5 and 10 — usable, but less polished (17)
🇦🇱 Albanian · 🇲🇼🇿🇲 Chichewa/Nyanja · 🇮🇳🇵🇰 Eastern Punjabi · 🇺🇬 Ganda · 🇮🇸 Icelandic · 🇮🇪 Irish · 🇩🇿 Kabyle · 🇨🇻 Kabuverdianu · 🇰🇪 Kamba · 🇻🇦 Latin · 🇱🇺 Luxembourgish · 🇪🇹🇰🇪 Oromo · 🇦🇫🇵🇰 Pashto · 🇵🇰🇮🇳 Sindhi · 🇸🇴 Somali · 🇦🇴 Umbundu · 🇬🇧 Welsh
All tags follow <|category:value|> syntax and can be inserted mid-utterance.
Emotion — elation , amusement , enthusiasm , determination , pride , contentment , affection , relief , contemplation , confusion , surprise , awe , longing , arousal , anger , fear , disgust , bitterness , sadness , shame , helplessness
Style — singing , shouting , whispering
Sound effects — cough , laughter , crying , screaming , burping , humming , sigh , sniff , sneeze
Pair each token with the matching onomatopoeia immediately after it.
Prosody
Speed — speed_very_slow , speed_slow , speed_fast , speed_very_fast
Pitch — pitch_low , pitch_high
Delivery — expressive_high , expressive_low
We evaluate Higgs Audio v3 TTS on public multilingual TTS suites and our internal 111-language Higgs-Multilingual set, covering both common and lower-resource languages.
WER / CER (↓, ×100) macro-averaged across each benchmark's language set. Lower is better; bold marks the best per row. All numbers are reproducible end-to-end with original metrics and normalization.
Win-rate (↑) per category — judge preference vs the BASELINE row; bold marks the highest win-rate per column. For a fair comparison, every model shares the same reference audio per prompt, and we run the benchmark text verbatim — no inline control tags inserted.
Pair the weights in this repo with SGLang-Omni — a production serving stack with continuous batching for multi-codebook decoding and the same inline tag controls. The Higgs TTS cookbook walks you through installation, server launch, request examples, and the full API reference.
See the Higgs TTS cookbook for the full details.
Install and Serve
docker pull lmsysorg/sglang-omni:dev
docker run -it --gpus all --shm-size 32g --ipc host --network host --privileged \
lmsysorg/sglang-omni:dev /bin/zsh
git clone git@github.com:sgl-project/sglang-omni.git && cd sglang-omni
uv venv .venv -p 3.12 && source .venv/bin/activate
uv pip install -v -e .
export HF_TOKEN=hf_xxxxxxxxxxxxxxxx
hf download bosonai/higgs-audio-v3-tts-4b
sgl-omni serve \
--model-path bosonai/higgs-audio-v3-tts-4b \
--port 8000
Zero-shot synthesis
curl -X POST http://localhost:8000/v1/audio/speech \
-H "Content-Type: application/json" \
-d '{"input": "Hello, how are you?"}' \
--output output.wav
Voice cloning
Supplying the reference transcript ( text ) materially improves cloning fidelity.
import requests
resp = requests.post(
"http://localhost:8000/v1/audio/speech" ,
json={
"input" : "Have a nice day and enjoy south california sunshine." ,
"references" : [{
"audio_path" : "ref.wav" ,
"text" : "Hey, Adam here. Let's create something that feels real, sounds human, and connects every time." ,
}],
"temperature" : 0.8 , "top_k" : 50 , "max_new_tokens" : 1024 ,
},
)
with open ( "output.wav" , "wb" ) as f:
f.write(resp.content)
Streaming (Server-Sent Events)
Set "stream": true to receive base64-encoded WAV chunks as the vocoder emits them — sub-second time-to-first-audio. Each event carries audio.data (base64 WAV bytes); the terminal event has finish_reason: "stop" plus usage metadata.
import requests, base64, json
with requests.post(
"http://localhost:8000/v1/audio/speech" ,
json={ "input" : "Get the trust fund to the bank early." , "stream" : True },
stream= True ,
) as resp, open ( "output.wav" , "wb" ) as f:
for line in resp.iter_lines():
if not line or not line.startswith( b"data: " ) or line == b"data: [DONE]" :
continue
event = json.loads(line[ 6 :])
if event.get( "finish_reason" ) == "stop" :
break
audio = event.get( "audio" ) or {}
if audio.get( "data" ):
f.write(base64.b64decode(audio[ "data" ]))
Inline control tokens
Embed <|emotion:…|> , <|style:…|> , <|prosody:…|> , and <|sfx:…|> tokens directly in input . Two rules:
Delivery tokens first. Emotion, style, and the prosody speed / pitch / expressive tokens shape the whole turn — put them at the start of input . Positional tokens ( <|prosody:pause|> , <|prosody:long_pause|> , <|sfx:…|> ) go inline exactly where they fire.
Pair every <|sfx:…|> with its onomatopoeia. E.g. <|sfx:laughter|>Haha , <|sfx:sigh|>Uh , <|sfx:sneeze|>Achoo . The written sound gives the model the acoustic cue to realize the effect.
Example — amusement + laughter:
curl -X POST http://localhost:8000/v1/audio/speech \
-H "Content-Type: application/json" \
-d '{"input": "<|emotion:amusement|><|prosody:expressive_high|>Wait, wait, that was kind of hilarious. <|sfx:laughter|>Hehe, no, seriously, I was not ready for that."}' \
--output output.wav
Throughput
Throughput on Seed-TTS EN (full set, N=1088 per run). Client --max-concurrency sweep against a Higgs server ( max_running_requests=16 , bf16, CUDA Graph on). Each row is the mean of 3 runs . Hardware: 1× H100 .
Concurrency — Maximum number of in-flight client requests ( --max-concurrency ).
Throughput (req/s) — Completed requests divided by total benchmark wall-clock time.
Mean latency — Average end-to-end time per request (send to full response received).
RTF (per-req) — Average ratio of processing time to generated audio duration per request (<1 is faster than real time).
audio_s/s — Total seconds of audio produced divided by total benchmark wall-clock time.
To reproduce the results, follow the instructions in this script .
For zero-ops deployment, use the Boson AI API .
@misc{bosonai_higgs_audio_tts_v3_2026,
title = {Higgs Audio v3 TTS: Conversational Speech for Voice AI from Boson AI},
author = {Boson AI},
year = {2026},
howpublished = {https://huggingface.co/bosonai/higgs-audio-v3-tts-4b},
}
License
Boson Higgs Audio v3 Research and Non-Commercial License — see LICENSE .
","pad_token":"<|endoftext|>","unk_token":null},"chat_template_jinja":"{%- if tools %}\n {{- '<|im_start|>system\\n' }}\n {%- if messages[0]['role'] == 'system' %}\n {{- messages[0]['content'] }}\n {%- else %}\n {{- 'You are a helpful assistant.' }}\n {%- endif %}\n {{- \"\\n\\n# Tools\\n\\nYou may call one or more functions to assist with the user query.\\n\\nYou are provided with function signatures within <tools></tools> XML tags:\\n<tools>\" }}\n {%- for tool in tools %}\n {{- \"\\n\" }}\n {{- tool | tojson }}\n {%- endfor %}\n {{- \"\\n</tools>\\n\\nFor each function call, return a json object with function name and arguments within <tool_call></tool_call> XML tags:\\n<tool_call>\\n{\\\"name\\\": <function-name>, \\\"arguments\\\": <args-json-object>}\\n</tool_call><|im_end|>\\n\" }}\n{%- else %}\n {%- if messages[0]['role'] == 'system' %}\n {{- '<|im_start|>system\\n' + messages[0]['content'] + '<|im_end|>\\n' }}\n {%- else %}\n {{- '<|im_start|>system\\nYou are a helpful assistant.<|im_end|>\\n' }}\n {%- endif %}\n{%- endif %}\n{%- for message in messages %}\n {%- if (message.role == \"user\") or (message.role == \"system\" and not loop.first) or (message.role == \"assistant\" and not message.tool_calls) %}\n {{- '<|im_start|>' + message.role + '\\n' + message.content + '<|im_end|>' + '\\n' }}\n {%- elif message.role == \"assistant\" %}\n {{- '<|im_start|>' + message.role }}\n {%- if message.content %}\n {{- '\\n' + message.content }}\n {%- endif %}\n {%- for tool_call in message.tool_calls %}\n {%- if tool_call.function is defined %}\n {%- set tool_call = tool_call.function %}\n {%- endif %}\n {{- '\\n<tool_call>\\n{\"name\": \"' }}\n {{- tool_call.name }}\n {{- '\", \"arguments\": ' }}\n {{- tool_call.arguments | tojson }}\n {{- '}\\n</tool_call>' }}\n {%- endfor %}\n {{- '<|im_end|>\\n' }}\n {%- elif message.role == \"tool\" %}\n {%- if (loop.index0 == 0) or (messages[loop.index0 - 1].role != \"tool\") %}\n {{- '<|im_start|>user' }}\n {%- endif %}\n {{- '\
[truncated]
Inference Providers NEW Text-to-Speech This model isn't deployed by any Inference Provider. 🙋 Ask for provider support Model tree for bosonai/higgs-audio-v3-tts-4b
Finetunes 1 model Quantizations 4 models Spaces using bosonai/higgs-audio-v3-tts-4b 2
Collection including bosonai/higgs-audio-v3-tts-4b
