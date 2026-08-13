---
source: "https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro-0813"
hn_url: "https://news.ycombinator.com/item?id=49285036"
title: "DeepSeek-AI/DeepSeek-V4-Pro-0813"
article_title: "deepseek-ai/DeepSeek-V4-Pro-0813 · Hugging Face"
author: "Philpax"
captured_at: "2026-08-13T12:45:32Z"
capture_tool: "hn-digest"
hn_id: 49285036
score: 1
comments: 0
posted_at: "2026-08-13T12:38:42Z"
tags:
  - hacker-news
  - translated
---

# DeepSeek-AI/DeepSeek-V4-Pro-0813

- HN: [49285036](https://news.ycombinator.com/item?id=49285036)
- Source: [huggingface.co](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro-0813)
- Score: 1
- Comments: 0
- Posted: 2026-08-13T12:38:42Z

## Translation

タイトル: DeepSeek-AI/DeepSeek-V4-Pro-0813
記事タイトル: deepseek-ai/DeepSeek-V4-Pro-0813 · 抱きしめる顔
説明: 私たちは、オープンソースとオープン サイエンスを通じて人工知能を進歩させ、民主化する旅の途中にあります。

記事本文:
deepseek-ai/DeepSeek-V4-Pro-0813 · 抱きしめる顔
ハグ顔モデル
deepseek-ai / DeepSeek-V4-Pro-0813 いいね 36 DeepSeek をフォローする 142k
テキスト生成トランスフォーマー Safetensors deepseek_v4 8 ビット精度 fp8 arxiv: 2606.19348 ライセンス: mit モデル カード ファイル ファイルとバージョン xet コミュニティ 3 デプロイ バケットにコピー 新しいモデルを使用 deepseek-ai/DeepSeek-V4-Pro-0813 をライブラリ、推論プロバイダー、ノートブック、ローカル アプリで使用する手順。これらのリンクに従って開始してください。
Transformers Transformers で deepseek-ai/DeepSeek-V4-Pro-0813 を使用する方法:
# パイプラインを高レベルのヘルパーとして使用する
変圧器からのインポートパイプライン
Pipe = Pipeline("text-generation", model="deepseek-ai/DeepSeek-V4-Pro-0813") # モデルを直接ロードする
トランスフォーマーから AutoTokenizer、AutoModelForCausalLM をインポート
tokenizer = AutoTokenizer.from_pretrained("deepseek-ai/DeepSeek-V4-Pro-0813")
model = AutoModelForCausalLM.from_pretrained("deepseek-ai/DeepSeek-V4-Pro-0813", device_map="auto")
vLLM vLLM で deepseek-ai/DeepSeek-V4-Pro-0813 を使用する方法:
# pip から vLLM をインストールします。
pip インストール vllm
# vLLM サーバーを起動します。
vllm サーブ「ディープシーク-ai/ディープシーク-V4-Pro-0813」
#curl (OpenAI 互換 API) を使用してサーバーを呼び出します。
curl -X POST "http://localhost:8000/v1/completions" \
-H "コンテンツ タイプ: application/json" \
--データ '{
"モデル": "ディープシーク-ai/ディープシーク-V4-Pro-0813",
"prompt": "むかしむかし、",
"max_tokens": 512、
「温度」：0.5
}' Docker を使用する docker モデル run hf.co/deepseek-ai/DeepSeek-V4-Pro-0813
SGLang SGLang で deepseek-ai/DeepSeek-V4-Pro-0813 を使用する方法:
# pip から SGLang をインストールします。
pip インストール sglang
# SGLang サーバーを起動します。
python3 -m sglang.launch_server \
--モデルパス "ディープシーク-ai/ディープシーク-V4-Pro-0813" \
--ホスト 0.0.0.0 \
--ポート 30000
#curlを使ってサーバーを呼び出す（OpenAI互換API）

:
curl -X POST "http://localhost:30000/v1/completions" \
-H "コンテンツ タイプ: application/json" \
--データ '{
"モデル": "ディープシーク-ai/ディープシーク-V4-Pro-0813",
"prompt": "むかしむかし、",
"max_tokens": 512、
「温度」：0.5
}' Docker イメージを使用する docker run --gpus all \
--shm-サイズ 32g \
-p 30000:30000 \
-v ~/.cache/huggingface:/root/.cache/huggingface \
--env "HF_TOKEN=<秘密>" \
--ipc=ホスト\
lmsysorg/sglang:最新 \
python3 -m sglang.launch_server \
--モデルパス "ディープシーク-ai/ディープシーク-V4-Pro-0813" \
--ホスト 0.0.0.0 \
--ポート 30000
#curl (OpenAI 互換 API) を使用してサーバーを呼び出します。
curl -X POST "http://localhost:30000/v1/completions" \
-H "コンテンツ タイプ: application/json" \
--データ '{
"モデル": "ディープシーク-ai/ディープシーク-V4-Pro-0813",
"prompt": "むかしむかし、",
"max_tokens": 512、
「温度」：0.5
}'
Docker Model Runner Docker Model Runner で deepseek-ai/DeepSeek-V4-Pro-0813 を使用する方法:
docker モデルの実行 hf.co/deepseek-ai/DeepSeek-V4-Pro-0813
DeepSeek-V4-Pro-0813 の概要
DeepSeek-V4-Pro-0813 は DeepSeek-V4-Pro の正式リリースであり、プレビュー バージョンに代わるもので、大幅に強化されたエージェント機能とパフォーマンスの向上が特に本番環境で顕著です。これは、DeepSeek-V4-Pro (プレビュー) モデル構造に基づいて構築されており、DSpark 投機的デコード モジュールが付属しています。
DeepSeek-V4-Pro-0813 は、以下に示すベンチマークで DeepSeek-V4-Pro (プレビュー) を上回り、入手可能な最も強力な独自モデルと広く競合します。
上記の公開ベンチマークのコード エージェント タスクについては、DeepSeek-V4-Pro-0813 は、温度 = 1.0、top_p = 0.95 の最大推論エフォート レベルを使用して、エージェント フレームワークとして DeepSeek Harness の最小モードで評価されます。
† DSBench-FullStack は、内部フルスタック開発テスト セットです。

DSBench-Hard は、コーディング エージェントに関する難しい問題の内部テスト セットです。
このリリースには、Jinja 形式のチャット テンプレートは含まれていません。代わりに、OpenAI 互換形式のメッセージをモデルの入力文字列にエンコードする方法と、モデルのテキスト出力を解析する方法を示す Python スクリプトとテスト ケースを備えた専用のエンコード フォルダーを提供します。完全なドキュメントについては、エンコーディング フォルダーを参照してください。
reasoning_effort パラメーターは、モデルが回答する前にどれだけの熟考を費やすかを制御する 3 つのレベル ( low 、 high 、および max) をサポートするようになりました。
fromcoding_dsv4 import encode_messages、parse_message_from_completion_text
メッセージ = [
{ "役割" : "ユーザー" 、 "コンテンツ" : "こんにちは" },
{ "role" : "アシスタント" , "content" : "こんにちは! 私は DeepSeek です。" , "reasoning_content" : "考え中..." },
{ "役割" : "ユーザー" 、 "コンテンツ" : "1+1=?" }
]
# メッセージ -> 文字列
プロンプト = encode_messages(メッセージ, Thinking_mode= "思考" ,reasoning_effort= "最大" )
# 文字列 -> トークン
輸入変圧器
tokenizer =Transformers.AutoTokenizer.from_pretrained( "deepseek-ai/DeepSeek-V4-Pro-0813" )
トークン = tokenizer.encode(プロンプト)
vLLM で実行する方法
DSpark の投機的デコードは、単一のフラグで有効になります。 --speculative-config をメソッド dspark とともに vLLM 起動コマンドに追加します。
--speculative-config '{"method":"dspark","num_speculative_tokens":7,"draft_sample_method":"greedy"}'
たとえば、以下のコマンドは、単一の 4×GB300 ノード上で vLLM を備えたモデルを提供します。
詳細な手順とその他のハードウェア構成については、vLLM レシピを参照してください。
vllmserve deepseek-ai/DeepSeek-V4-Pro-0813 \
--trust-remote-code --kv-cache-dtype fp8 --block-size 256 \
--data-Parallel-size 4 --enable-expert-Parallel \
--moe-backend deep_gemm_mega_moe \
--attention-config '{"use_fp4_indexer_cache": true}' \

--speculative-config '{"method":"dspark","num_speculative_tokens":7,"draft_sample_method":"greedy"}'
SGLang で実行する方法
--speculative-algorithm DSPARK で DSpark を有効にし、ターゲットとして別の --speculative-draft-model-path を設定しないでください。したがって、ドラフトの重みは同じチェックポイントから取得されます。
詳細な手順、ベンチマーク、その他のハードウェア構成については、SGLang クックブックを参照してください。
スラングサーブ \
--trust-remote-code \
--model-path deepseek-ai/DeepSeek-V4-Pro-0813 \
--tp 4 \
--moe-runner-backend flashinfer_mxfp4 \
--投機的アルゴリズム DSPARK \
--mem-fraction-static 0.90 \
--chunked-prefill-size 4096 \
--swa-full-tokens-ratio 0.1 \
ローカルで実行する方法
モデルの重み変換やインタラクティブなチャット デモなど、DeepSeek-V4 をローカルで実行するための詳細な手順については、推論フォルダーを参照してください。
ローカル デプロイメントの場合は、サンプリング パラメーターを 温度 = 1.0 に設定し、エージェント シナリオの場合は top_p = 0.95、それ以外の場合は top_p = 1.0 に設定することをお勧めします。推論努力レベルが高く、最大の場合は、最大出力長を 384K トークンにすることをお勧めします。
このリポジトリとモデルの重みは、MIT ライセンスに基づいてライセンスされています。
@misc{deepseekai2026deepseekv4,
title={DeepSeek-V4: 非常に効率的な 100 万トークンのコンテキスト インテリジェンスに向けて},
著者={DeepSeek-AI}、
年={2026}、
}
お問い合わせ
ご質問がある場合は、問題を提起するか、service@deepseek.com までご連絡ください。
","lstrip":false,"normalized":true,"rstrip":false,"single_word":false},"eos_token":{"__type":"AddedToken","content":"<｜文末｜>","lstrip":false,"normalized":true,"rstrip":false ,"single_word":false},"pad_token":{"__type":"AddedToken","content":"<｜文末｜>","lstrip":false,"normalized":true,"rstrip":false,"single_word":false},"unk_token":null}},"createdA

t":"2026-08-13T03:05:06.000Z","DiscussionsDisabled":false,"DiscussionsSorting":"recently-created","downloads":0,"downloadsAllTime":0,"id":"deepseek-ai/DeepSeek-V4-Pro-0813"," isLikedByUser":false,"availableInferenceProviders":[],"showHuggingChatEntry":false,"inference":"","lastModified":"2026-08-13T12:25:17.000Z","likes":36,"pipeline_tag":"text-gen eration","library_name":"transformers","librariesOther":[],"trackDownloads":true,"model-index":null,"private":false,"repoType":"model","gated":false,"tags":["transformers","sa fetensors","deepseek_v4","テキスト生成","arxiv:2606.19348","ライセンス:mit","エンドポイント互換","8 ビット","fp8","地域:us"],"tag_objs":[{"id":"テキスト生成","ラベル":"テキストGeneration","type":"pipeline_tag","subType":"nlp"},{"id":"transformers","label":"Transformers","type":"library"},{"id":"safetensors","label":"Safetセンサー","タイプ":"ライブラリ"},{"id":"ディープシーク_v4","ラベル":"ディープシーク_v4","タイプ":"その他","クリック可能":true},{"id":"エンドポイント互換","ラベル":"推論エンドポイント","タイプ":"その他","クリック可能":true},{"id":"8 ビット","ラベル":"8 ビット精度","タイプ":"その他","クリック可能":true},{"id":"fp8","ラベル":"fp8","タイプ":"その他","クリック可能":true} ,{"id":"arxiv:2606.19348","label":"arxiv:2606.19348","type":"arxiv","extra":{"paperTitle":"DeepSeek-V4:非常に効率的な 100 万トークンのコンテキスト インテリジェンスに向けて"}},{"id":"license:mit","label":"mit","type":"license"},{"type":"region","label":"🇺🇸 リージョン: US","id":"region:us"}],"transformersInfo":{"auto_model":"AutoModelForCausalLM","pipeline_tag":"text-generation","processor":"AutoTokenizer"},"widgetData":[{"text":"私の名前はジュリアンです。"},{"te
[切り捨てられた]
推論プロバイダー 新しいテキストの生成 このモデルは、どの推論プロバイダーによってもデプロイされません。 🙋 2 プロバイダーのサポートを依頼する deepse を含むコレクション

ek-ai/DeepSeek-V4-Pro-0813
812 deepseek-ai/DeepSeek-V4-Pro-0813 用ペーパー

## Original Extract

We’re on a journey to advance and democratize artificial intelligence through open source and open science.

deepseek-ai/DeepSeek-V4-Pro-0813 · Hugging Face
Hugging Face Models
deepseek-ai / DeepSeek-V4-Pro-0813 like 36 Follow DeepSeek 142k
Text Generation Transformers Safetensors deepseek_v4 8-bit precision fp8 arxiv: 2606.19348 License: mit Model card Files Files and versions xet Community 3 Deploy Copy to bucket new Use this model Instructions to use deepseek-ai/DeepSeek-V4-Pro-0813 with libraries, inference providers, notebooks, and local apps. Follow these links to get started.
Transformers How to use deepseek-ai/DeepSeek-V4-Pro-0813 with Transformers:
# Use a pipeline as a high-level helper
from transformers import pipeline
pipe = pipeline("text-generation", model="deepseek-ai/DeepSeek-V4-Pro-0813") # Load model directly
from transformers import AutoTokenizer, AutoModelForCausalLM
tokenizer = AutoTokenizer.from_pretrained("deepseek-ai/DeepSeek-V4-Pro-0813")
model = AutoModelForCausalLM.from_pretrained("deepseek-ai/DeepSeek-V4-Pro-0813", device_map="auto")
vLLM How to use deepseek-ai/DeepSeek-V4-Pro-0813 with vLLM:
# Install vLLM from pip:
pip install vllm
# Start the vLLM server:
vllm serve "deepseek-ai/DeepSeek-V4-Pro-0813"
# Call the server using curl (OpenAI-compatible API):
curl -X POST "http://localhost:8000/v1/completions" \
-H "Content-Type: application/json" \
--data '{
"model": "deepseek-ai/DeepSeek-V4-Pro-0813",
"prompt": "Once upon a time,",
"max_tokens": 512,
"temperature": 0.5
}' Use Docker docker model run hf.co/deepseek-ai/DeepSeek-V4-Pro-0813
SGLang How to use deepseek-ai/DeepSeek-V4-Pro-0813 with SGLang:
# Install SGLang from pip:
pip install sglang
# Start the SGLang server:
python3 -m sglang.launch_server \
--model-path "deepseek-ai/DeepSeek-V4-Pro-0813" \
--host 0.0.0.0 \
--port 30000
# Call the server using curl (OpenAI-compatible API):
curl -X POST "http://localhost:30000/v1/completions" \
-H "Content-Type: application/json" \
--data '{
"model": "deepseek-ai/DeepSeek-V4-Pro-0813",
"prompt": "Once upon a time,",
"max_tokens": 512,
"temperature": 0.5
}' Use Docker images docker run --gpus all \
--shm-size 32g \
-p 30000:30000 \
-v ~/.cache/huggingface:/root/.cache/huggingface \
--env "HF_TOKEN=<secret>" \
--ipc=host \
lmsysorg/sglang:latest \
python3 -m sglang.launch_server \
--model-path "deepseek-ai/DeepSeek-V4-Pro-0813" \
--host 0.0.0.0 \
--port 30000
# Call the server using curl (OpenAI-compatible API):
curl -X POST "http://localhost:30000/v1/completions" \
-H "Content-Type: application/json" \
--data '{
"model": "deepseek-ai/DeepSeek-V4-Pro-0813",
"prompt": "Once upon a time,",
"max_tokens": 512,
"temperature": 0.5
}'
Docker Model Runner How to use deepseek-ai/DeepSeek-V4-Pro-0813 with Docker Model Runner:
docker model run hf.co/deepseek-ai/DeepSeek-V4-Pro-0813
DeepSeek-V4-Pro-0813 Introduction
DeepSeek-V4-Pro-0813 is the official release of DeepSeek-V4-Pro , superseding the preview version, with greatly enhanced agentic capabilities and performance improvements that are especially pronounced in production environments. It is built on the DeepSeek-V4-Pro (Preview) model structure, with a DSpark speculative decoding module attached.
DeepSeek-V4-Pro-0813 outperforms DeepSeek-V4-Pro (Preview) on the benchmarks listed below, and is broadly competitive with the strongest proprietary models available.
For the code-agent tasks among the public benchmarks above, DeepSeek-V4-Pro-0813 is evaluated with the minimal mode of DeepSeek Harness as the agent framework, using the max reasoning effort level with temperature = 1.0, top_p = 0.95 .
† DSBench-FullStack is an internal full-stack development test set; DSBench-Hard is an internal test set of difficult coding-agent problems.
This release does not include a Jinja-format chat template. Instead, we provide a dedicated encoding folder with Python scripts and test cases demonstrating how to encode messages in OpenAI-compatible format into input strings for the model, and how to parse the model's text output. Please refer to the encoding folder for full documentation.
The reasoning_effort parameter now supports three levels — low , high , and max — which control how much deliberation the model spends before answering.
from encoding_dsv4 import encode_messages, parse_message_from_completion_text
messages = [
{ "role" : "user" , "content" : "hello" },
{ "role" : "assistant" , "content" : "Hello! I am DeepSeek." , "reasoning_content" : "thinking..." },
{ "role" : "user" , "content" : "1+1=?" }
]
# messages -> string
prompt = encode_messages(messages, thinking_mode= "thinking" , reasoning_effort= "max" )
# string -> tokens
import transformers
tokenizer = transformers.AutoTokenizer.from_pretrained( "deepseek-ai/DeepSeek-V4-Pro-0813" )
tokens = tokenizer.encode(prompt)
How to Run with vLLM
DSpark speculative decoding is enabled with a single flag — add --speculative-config with method: dspark to your vLLM launch command:
--speculative-config '{"method":"dspark","num_speculative_tokens":7,"draft_sample_method":"greedy"}'
For example, the command below serves the model with vLLM on a single 4×GB300 node.
See the vLLM recipe for detailed instructions and other hardware configurations.
vllm serve deepseek-ai/DeepSeek-V4-Pro-0813 \
--trust-remote-code --kv-cache-dtype fp8 --block-size 256 \
--data-parallel-size 4 --enable-expert-parallel \
--moe-backend deep_gemm_mega_moe \
--attention-config '{"use_fp4_indexer_cache": true}' \
--speculative-config '{"method":"dspark","num_speculative_tokens":7,"draft_sample_method":"greedy"}'
How to Run with SGLang
Enable DSpark with --speculative-algorithm DSPARK and do not set a separate --speculative-draft-model-path as the target and draft weights therefore come from the same checkpoint.
See the SGLang cookbook for detailed instructions, benchmarks and other hardwares configurations.
sglang serve \
--trust-remote-code \
--model-path deepseek-ai/DeepSeek-V4-Pro-0813 \
--tp 4 \
--moe-runner-backend flashinfer_mxfp4 \
--speculative-algorithm DSPARK \
--mem-fraction-static 0.90 \
--chunked-prefill-size 4096 \
--swa-full-tokens-ratio 0.1 \
How to Run Locally
Please refer to the inference folder for detailed instructions on running DeepSeek-V4 locally, including model weight conversion and interactive chat demos.
For local deployment, we recommend setting the sampling parameters to temperature = 1.0 , with top_p = 0.95 for agentic scenarios and top_p = 1.0 otherwise. For the high and max reasoning effort levels, we recommend a maximum output length of 384K tokens.
This repository and the model weights are licensed under the MIT License .
@misc{deepseekai2026deepseekv4,
title={DeepSeek-V4: Towards Highly Efficient Million-Token Context Intelligence},
author={DeepSeek-AI},
year={2026},
}
Contact
If you have any questions, please raise an issue or contact us at service@deepseek.com .
","lstrip":false,"normalized":true,"rstrip":false,"single_word":false},"eos_token":{"__type":"AddedToken","content":"<｜end▁of▁sentence｜>","lstrip":false,"normalized":true,"rstrip":false,"single_word":false},"pad_token":{"__type":"AddedToken","content":"<｜end▁of▁sentence｜>","lstrip":false,"normalized":true,"rstrip":false,"single_word":false},"unk_token":null}},"createdAt":"2026-08-13T03:05:06.000Z","discussionsDisabled":false,"discussionsSorting":"recently-created","downloads":0,"downloadsAllTime":0,"id":"deepseek-ai/DeepSeek-V4-Pro-0813","isLikedByUser":false,"availableInferenceProviders":[],"showHuggingChatEntry":false,"inference":"","lastModified":"2026-08-13T12:25:17.000Z","likes":36,"pipeline_tag":"text-generation","library_name":"transformers","librariesOther":[],"trackDownloads":true,"model-index":null,"private":false,"repoType":"model","gated":false,"tags":["transformers","safetensors","deepseek_v4","text-generation","arxiv:2606.19348","license:mit","endpoints_compatible","8-bit","fp8","region:us"],"tag_objs":[{"id":"text-generation","label":"Text Generation","type":"pipeline_tag","subType":"nlp"},{"id":"transformers","label":"Transformers","type":"library"},{"id":"safetensors","label":"Safetensors","type":"library"},{"id":"deepseek_v4","label":"deepseek_v4","type":"other","clickable":true},{"id":"endpoints_compatible","label":"Inference Endpoints","type":"other","clickable":true},{"id":"8-bit","label":"8-bit precision","type":"other","clickable":true},{"id":"fp8","label":"fp8","type":"other","clickable":true},{"id":"arxiv:2606.19348","label":"arxiv:2606.19348","type":"arxiv","extra":{"paperTitle":"DeepSeek-V4: Towards Highly Efficient Million-Token Context Intelligence"}},{"id":"license:mit","label":"mit","type":"license"},{"type":"region","label":"🇺🇸 Region: US","id":"region:us"}],"transformersInfo":{"auto_model":"AutoModelForCausalLM","pipeline_tag":"text-generation","processor":"AutoTokenizer"},"widgetData":[{"text":"My name is Julien and I like to"},{"te
[truncated]
Inference Providers NEW Text Generation This model isn't deployed by any Inference Provider. 🙋 2 Ask for provider support Collection including deepseek-ai/DeepSeek-V4-Pro-0813
812 Paper for deepseek-ai/DeepSeek-V4-Pro-0813
