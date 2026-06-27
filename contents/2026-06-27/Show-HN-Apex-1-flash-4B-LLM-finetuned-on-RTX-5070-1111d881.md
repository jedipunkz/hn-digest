---
source: "https://huggingface.co/OrbitAIEU/Apex-1-flash"
hn_url: "https://news.ycombinator.com/item?id=48698923"
title: "Show HN: Apex-1-flash, 4B LLM finetuned on RTX 5070"
article_title: "OrbitAIEU/Apex-1-flash · Hugging Face"
author: "Qmay_Dev"
captured_at: "2026-06-27T15:26:34Z"
capture_tool: "hn-digest"
hn_id: 48698923
score: 1
comments: 0
posted_at: "2026-06-27T15:07:51Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Apex-1-flash, 4B LLM finetuned on RTX 5070

- HN: [48698923](https://news.ycombinator.com/item?id=48698923)
- Source: [huggingface.co](https://huggingface.co/OrbitAIEU/Apex-1-flash)
- Score: 1
- Comments: 0
- Posted: 2026-06-27T15:07:51Z

## Translation

タイトル: HN を表示: Apex-1-flash、RTX 5070 で微調整された 4B LLM
記事タイトル: OrbitAIEU/Apex-1-flash・ハグフェイス
説明: 私たちは、オープンソースとオープン サイエンスを通じて人工知能を進歩させ、民主化する旅の途中にあります。
HN テキスト: 目標は、消費者向けハードウェアで簡単に実行できるほど軽量でありながら、推論タスクを実行できる、高効率の小規模モデルを作成することでした。
技術スタック:
ベース: クウェン3:4B
トレーニング: Unsloth を使用してメモリ効率を微調整したため、RTX 5070 でプロセスをスムーズに実行できるようになりました。
スタック: cu128、PyTorch、および Hugging Face Transformers で構築されています。
データセット: 思考連鎖 (CoT) 機能を向上させるために、Raymond-dev-546730/Open-CoT-Reasoning-Mini でトレーニングされました。

記事本文:
OrbitAIEU/Apex-1-フラッシュ・ハグフェイス
ハグ顔モデル
OrbitAIEU / Apex-1-flash のような 1 OrbitAI をフォロー 1
テキスト生成トランスフォーマー GGUF Raymond-dev-546730/Open-CoT-Reasoning-Mini 英語 中国語 qwen3 思考推論 qwen ヨーロッパ ai 微調整 会話型 ライセンス: apache-2.0 モデル カード ファイル ファイルとバージョン xet コミュニティ デプロイ バケットにコピー 新しいモデルを使用する OrbitAIEU/Apex-1-flash をライブラリ、推論プロバイダー、ノートブック、ローカル アプリで使用する手順。これらのリンクに従って開始してください。
トランスフォーマー OrbitAIEU/Apex-1-flash をトランスフォーマーで使用する方法:
# パイプラインを高レベルのヘルパーとして使用する
変圧器からのインポートパイプライン
Pipe = パイプライン("テキスト生成", モデル="OrbitAIEU/Apex-1-flash")
メッセージ = [
{"役割": "ユーザー", "コンテンツ": "あなたは誰ですか?"},
]
Pipe(messages) # モデルを直接ロードする
トランスフォーマーから AutoTokenizer、AutoModelForCausalLM をインポート
tokenizer = AutoTokenizer.from_pretrained("OrbitAIEU/Apex-1-flash")
モデル = AutoModelForCausalLM.from_pretrained("OrbitAIEU/Apex-1-フラッシュ")
メッセージ = [
{"役割": "ユーザー", "コンテンツ": "あなたは誰ですか?"},
]
入力 = tokenizer.apply_chat_template(
メッセージ、
add_generation_prompt=真、
tokenize=True、
return_dict=真、
return_tensors="pt",
).to(モデル.デバイス)
出力 = model.generate(**入力、max_new_tokens=40)
print(tokenizer.decode(outputs[0][inputs["input_ids"].shape[-1]:]))
llama-cpp-python llama-cpp-python で OrbitAIEU/Apex-1-flash を使用する方法:
# !pip llama-cpp-python をインストールします
llama_cpp から Llama をインポート
llm = Llama.from_pretrained(
repo_id="OrbitAIEU/Apex-1-フラッシュ",
ファイル名 = "apex-1-flash-q4_k_m.gguf",
)
llm.create_chat_completion(
メッセージ = [
{
"ロール": "ユーザー",
"content": 「フランスの首都はどこですか?」
}
]
)
llama.cpp llama.cpp で OrbitAIEU/Apex-1-flash を使用する方法:
カール -LsSf https://lla

ma.app/install.sh |しー
# Web UI を使用してローカルの OpenAI 互換サーバーを起動します。
ラマ サーブ -hf OrbitAIEU/Apex-1-flash:Q4_K_M
# ターミナルで直接推論を実行します。
llama cli -hf OrbitAIEU/Apex-1-flash:Q4_K_M WinGet からインストール (Windows) winget install llama.cpp
# Web UI を使用してローカルの OpenAI 互換サーバーを起動します。
ラマ サーブ -hf OrbitAIEU/Apex-1-flash:Q4_K_M
# ターミナルで直接推論を実行します。
llama cli -hf OrbitAIEU/Apex-1-flash:Q4_K_M ビルド済みバイナリを使用 # ビルド済みバイナリを次からダウンロードします。
# https://github.com/ggerganov/llama.cpp/releases
# Web UI を使用してローカルの OpenAI 互換サーバーを起動します。
./llama-server -hf OrbitAIEU/Apex-1-flash:Q4_K_M
# ターミナルで直接推論を実行します。
./llama-cli -hf OrbitAIEU/Apex-1-flash:Q4_K_M ソース コード git clone からビルド https://github.com/ggerganov/llama.cpp.git
cd ラマ.cpp
cmake -B ビルド
cmake --build build -j --target llama-server llama-cli
# Web UI を使用してローカルの OpenAI 互換サーバーを起動します。
./build/bin/llama-server -hf OrbitAIEU/Apex-1-flash:Q4_K_M
# ターミナルで直接推論を実行します。
./build/bin/llama-cli -hf OrbitAIEU/Apex-1-flash:Q4_K_M Docker を使用する docker モデルの実行 hf.co/OrbitAIEU/Apex-1-flash:Q4_K_M
vLLM vLLM で OrbitAIEU/Apex-1-flash を使用する方法:
# pip から vLLM をインストールします。
pip インストール vllm
# vLLM サーバーを起動します。
vllm サーブ「OrbitAIEU/Apex-1-flash」
#curl (OpenAI 互換 API) を使用してサーバーを呼び出します。
curl -X POST "http://localhost:8000/v1/chat/completions" \
-H "コンテンツ タイプ: application/json" \
--データ '{
"モデル": "OrbitAIEU/Apex-1-フラッシュ",
「メッセージ」: [
{
"ロール": "ユーザー",
"content": 「フランスの首都はどこですか?」
}
]
}' Docker を使用する docker モデル run hf.co/OrbitAIEU/Apex-1-flash:Q4_K_M
SGLang SGLang で OrbitAIEU/Apex-1-flash を使用する方法:
# pip から SGLang をインストールします。
pip インストール sglang
# SGLaを起動する

NGサーバー:
python3 -m sglang.launch_server \
--モデルパス "OrbitAIEU/Apex-1-フラッシュ" \
--ホスト 0.0.0.0 \
--ポート 30000
#curl (OpenAI 互換 API) を使用してサーバーを呼び出します。
curl -X POST "http://localhost:30000/v1/chat/completions" \
-H "コンテンツ タイプ: application/json" \
--データ '{
"モデル": "OrbitAIEU/Apex-1-フラッシュ",
「メッセージ」: [
{
"ロール": "ユーザー",
"content": 「フランスの首都はどこですか?」
}
]
}' Docker イメージを使用する docker run --gpus all \
--shm-サイズ 32g \
-p 30000:30000 \
-v ~/.cache/huggingface:/root/.cache/huggingface \
--env "HF_TOKEN=<秘密>" \
--ipc=ホスト\
lmsysorg/sglang:最新 \
python3 -m sglang.launch_server \
--モデルパス "OrbitAIEU/Apex-1-フラッシュ" \
--ホスト 0.0.0.0 \
--ポート 30000
#curl (OpenAI 互換 API) を使用してサーバーを呼び出します。
curl -X POST "http://localhost:30000/v1/chat/completions" \
-H "コンテンツ タイプ: application/json" \
--データ '{
"モデル": "OrbitAIEU/Apex-1-フラッシュ",
「メッセージ」: [
{
"ロール": "ユーザー",
"content": 「フランスの首都はどこですか?」
}
]
}'
Ollama Ollama で OrbitAIEU/Apex-1-flash を使用する方法:
オラマ ラン hf.co/OrbitAIEU/Apex-1-flash:Q4_K_M
Unsloth Studio Unsloth Studio で OrbitAIEU/Apex-1-flash を使用する方法:
カール -fsSL https://unsloth.ai/install.sh |しー
# unsloth スタジオを運営する
unsloth スタジオ -H 0.0.0.0 -p 8888
# 次に、ブラウザで http://localhost:8888 を開きます
# OrbitAIEU/Apex-1-flash を検索してチャットを開始します。 Unsloth Studio (Windows) をインストールします。irm https://unsloth.ai/install.ps1 |アイエックス
# unsloth スタジオを運営する
unsloth スタジオ -H 0.0.0.0 -p 8888
# 次に、ブラウザで http://localhost:8888 を開きます
# OrbitAIEU/Apex-1-flash を検索して、Unsloth 用の HuggingFace Spaces を使用してチャットを開始 # セットアップは必要ありません
# ブラウザで https://huggingface.co/spaces/unsloth/studio を開きます
# OrbitAIEU/Apex-1-flash を検索してチャットを開始します
Pi OrbitAIEU/Apex-1-flash ウィットの使い方

h 円周率:
# llama.cpp をインストールします。
醸造インストール llama.cpp
# ローカルの OpenAI 互換サーバーを起動します。
llamaserve -hf OrbitAIEU/Apex-1-flash:Q4_K_M Pi でモデルを構成します。 # Pi をインストールします。
npm install -g @mariozechner/pi-coding-agent
# ~/.pi/agent/models.json に追加します:
{
「プロバイダー」: {
"ラマ-cpp": {
"baseUrl": "http://localhost:8080/v1",
"api": "openai-completions",
"apiKey": "なし",
「モデル」: [
{
"id": "OrbitAIEU/Apex-1-flash:Q4_K_M"
}
]
}
}
Run Pi # プロジェクト ディレクトリで Pi を起動します。
円周率
ヘルメス エージェントの新しいヘルメス エージェントで OrbitAIEU/Apex-1-flash を使用する方法:
# llama.cpp をインストールします。
醸造インストール llama.cpp
# ローカルの OpenAI 互換サーバーを起動します。
llamaserve -hf OrbitAIEU/Apex-1-flash:Q4_K_M Herme の設定 # Herme をインストールします。
カール -fsSL https://hermes-agent.nousresearch.com/install.sh |バッシュ
エルメスのセットアップ
# ローカルサーバーにHermesを指定します:
hermes config set model.provider カスタム
hermes 設定セット model.base_url http://127.0.0.1:8080/v1
hermes config set model.default OrbitAIEU/Apex-1-flash:Q4_K_M Hermes を実行します。
Docker Model Runner Docker Model Runner で OrbitAIEU/Apex-1-flash を使用する方法:
docker モデルの実行 hf.co/OrbitAIEU/Apex-1-flash:Q4_K_M
Lemonade Lemonade で OrbitAIEU/Apex-1-flash を使用する方法:
# https://lemonade-server.ai/ からレモネードをダウンロード
レモネード プル OrbitAIEU/Apex-1-flash:Q4_K_M モデル レモネード実行ユーザーと実行してチャットします。Apex-1-flash-Q4_K_M 利用可能なすべてのモデル レモネード リストをリストします。
🧠 apex-1-flash の違いは以下に最適です
👤 クリエイターのマティアス・ミクルについて
速い。シャープ。話す前に考える。
OrbitAI による思考連鎖推論モデル
スロバキア出身の 13 歳の開発者によって作られました。好奇心には年齢制限がないからです。
Apex-1-flash は、 Qwen/qwen3-4b- Thinking-2507 の監督下での微調整であり、鋭く、構造化されたウィットに富んだ推論を提供することを目的として構築されています。

4B パラメータスケールでの効率的な思考連鎖機能。
Open-CoT-Reasoning-Mini データセットでトレーニングされた apex-1-flash は、問題を段階的に検討するように設計されており、消費者向けハードウェアで実行できるほど無駄を省きながら、論理的推論、複数ステップの問題解決、一貫した説明に適しています。
このモデルは、Matias Mikle (13 歳、スロバキア 🇸🇰) が OrbitAI チームと協力して作成しました。
🧠 apex-1-flash の違い
名前がすべてを物語っています。頂点に到達するための Apex、スピードと正確性のための Flash。
フラッシュの哲学は、モデルの構築方法を形成します。
⚡ 高速 — パラメータがわずか約 4B なので、推論の深さを犠牲にすることなく単一のコンシューマー GPU で実行できるほど軽量です。
🎯 シャープ — 構造化された思考連鎖データに基づいて特に微調整されており、答えを導き出す前に問題をきれいに分解します。
💡 思慮深い — Qwen3 から組み込まれた思考アーキテクチャを継承し、CoT の微調整を通じて拡張され、より信頼性の高いステップバイステップのロジックを実現します
論理的および数学的推論
段階的な問題の分解
構造化された説明の生成
研究および教育業務
トランスフォーマーから AutoTokenizer、AutoModelForCausalLM をインポート
輸入トーチ
model_id = "OrbitAIEU/apex-1-フラッシュ"
tokenizer = AutoTokenizer.from_pretrained(model_id)
モデル = AutoModelForCausalLM.from_pretrained(
モデルID、
torch_dtype=torch.bfloat16、
device_map= "自動"
)
メッセージ = [
{
"ロール" : "ユーザー" 、
"content" : "3x + 7 = 22 の解き方を段階的に説明します"
}
]
text = tokenizer.apply_chat_template(
メッセージ、
tokenize= False 、
add_generation_prompt= True
)
inputs = tokenizer([text], return_tensors= "pt" ).to(model.device)
torch.no_grad() を使用:
出力 = model.generate(
**入力、
max_new_tokens= 512 、
温度= 0.7 、
do_sample= True
)

応答 = tokenizer.decode(
出力[ 0 ][入力.input_ids.shape[- 1 ]:],
Skip_special_tokens= True
)
印刷（応答）
💾 ハードウェア要件
apex-1-flash は意図的に 4B スケールで構築されているため、日常的なハードウェアで実行でき、エンタープライズ クラスターは必要ありません。
モデルは、Qwen3-4B の思考チェックポイントに加えて、教師あり微調整 (SFT) を使用して微調整されました。
Open-CoT-Reasoning-Mini データセットは、慎重に構造化された推論トレースと思考連鎖の例を提供し、モデルが複数ステップの論理推論に関するより強力な習慣を構築できるようにします。
安全調整なし — Apex-1-flash は RLHF または安全調整を受けていません。追加の安全層なしで実稼働環境で使用することはお勧めできません。
ドメイン スコープ — パフォーマンスは、推論が必要なタスク向けに最適化されています。汎用機能はベースモデルから継承されます。
継承されたバイアス — モデルは、Qwen3-4B 基本モデルに存在するバイアスと制限を引き継ぐ可能性があります。
ベンチマーク保留中 — 正式なベンチマーク評価は現在進行中であり、将来のアップデートで公開される予定です。
年齢: 13歳 · 国: スロバキア 🇸🇰
独立系開発者、AI 研究者、OrbitAI の創設者。
Matias は AI プロジェクトをゼロから構築し、微調整、言語モデル アーキテクチャ、フルスタック開発を模索し、いつでもどこからでも優れた作品を生み出すことができることを証明しました。
「AI モデルをトレーニングするのに博士号は必要ありません。必要なのはインテリジェンスと GPU だけです。」
OrbitAI は、オープンで効率的でアクセスしやすい言語モデルの構築に重点を置いた独立した AI 開発チームです。
研究チームは、AI 研究は大企業や資金豊富な研究機関に限定されるべきではないと考えています。 OrbitAI は、モデルのリリース、実験の共有、コミュニティとのコラボレーションなど、オープンな環境で作業することで、

ロンティア スタイルの AI の仕事は、努力する意欲のある人なら誰でもアクセスできます。
apex-1-flash は、OrbitAI の最初の公開モデル リリースです。
このモデルは、ベース モデル Qwen/qwen3-4b- Thinking-2507 のライセンスに従って、 Apache License 2.0 に基づいてリリースされます。
完全な条項については、Apache 2.0 ライセンスの全文を参照してください。
Qwen Team @ Alibaba Cloud — 強力な Qwen3 モデル ファミリをオープン ライセンスでリリースするため
Raymond-dev-546730 — Open-CoT-Reasoning-Mini データセットの作成と共有用
オープンソース AI コミュニティ — これらすべてを可能にする
Apex-1-flash · Matias Mikle & OrbitAI による ❤️ 製 · スロバキア 🇸🇰
このプロジェクトにインスピレーションを受けた場合は、ダウンロードしてフォークし、さらに優れたものを構築してください。
","pad_token":"<|PAD_TOKEN|>","unk_token":null},"chat_template_jinja":"{%- if tools %}\n {{- '<|im_start|>system\\n' }}\n {%- ifmessages[0].role == 'system' %}\n {{-messages[0].content + '\\n\\n' }}\n {%- endif %}\n {{- \"# Tools\\n\\nユーザー クエリを支援するために 1 つ以上の関数を呼び出すことができます。\\n\\n<tools></tools> XML タグ内で関数シグネチャが提供されます:\\n<tools>\" }}\n {%- for tools in tools %}\n {{- \"\\n\" }}\n {{- tool | } tojson }}\n {%- endfor %}\n {{- \"\\n</tools>\\n\\n関数呼び出しごとに、関数名が a の json オブジェクトを返します。

[切り捨てられた]

## Original Extract

We’re on a journey to advance and democratize artificial intelligence through open source and open science.

The goal was to create a highly efficient, small-scale model that can perform reasoning tasks while remaining lightweight enough to run easily on consumer hardware.
Technical Stack:
Base: Qwen3:4B
Training: Fine-tuned using Unsloth for memory efficiency, which allowed me to run the process smoothly on an RTX 5070.
Stack: Built with cu128, PyTorch, and Hugging Face Transformers.
Dataset: Trained on Raymond-dev-546730/Open-CoT-Reasoning-Mini to improve Chain-of-Thought (CoT) capabilities.

OrbitAIEU/Apex-1-flash · Hugging Face
Hugging Face Models
OrbitAIEU / Apex-1-flash like 1 Follow OrbitAI 1
Text Generation Transformers GGUF Raymond-dev-546730/Open-CoT-Reasoning-Mini English Chinese qwen3 thinking reasoning qwen europe ai fine-tunning conversational License: apache-2.0 Model card Files Files and versions xet Community Deploy Copy to bucket new Use this model Instructions to use OrbitAIEU/Apex-1-flash with libraries, inference providers, notebooks, and local apps. Follow these links to get started.
Transformers How to use OrbitAIEU/Apex-1-flash with Transformers:
# Use a pipeline as a high-level helper
from transformers import pipeline
pipe = pipeline("text-generation", model="OrbitAIEU/Apex-1-flash")
messages = [
{"role": "user", "content": "Who are you?"},
]
pipe(messages) # Load model directly
from transformers import AutoTokenizer, AutoModelForCausalLM
tokenizer = AutoTokenizer.from_pretrained("OrbitAIEU/Apex-1-flash")
model = AutoModelForCausalLM.from_pretrained("OrbitAIEU/Apex-1-flash")
messages = [
{"role": "user", "content": "Who are you?"},
]
inputs = tokenizer.apply_chat_template(
messages,
add_generation_prompt=True,
tokenize=True,
return_dict=True,
return_tensors="pt",
).to(model.device)
outputs = model.generate(**inputs, max_new_tokens=40)
print(tokenizer.decode(outputs[0][inputs["input_ids"].shape[-1]:]))
llama-cpp-python How to use OrbitAIEU/Apex-1-flash with llama-cpp-python:
# !pip install llama-cpp-python
from llama_cpp import Llama
llm = Llama.from_pretrained(
repo_id="OrbitAIEU/Apex-1-flash",
filename="apex-1-flash-q4_k_m.gguf",
)
llm.create_chat_completion(
messages = [
{
"role": "user",
"content": "What is the capital of France?"
}
]
)
llama.cpp How to use OrbitAIEU/Apex-1-flash with llama.cpp:
curl -LsSf https://llama.app/install.sh | sh
# Start a local OpenAI-compatible server with a web UI:
llama serve -hf OrbitAIEU/Apex-1-flash:Q4_K_M
# Run inference directly in the terminal:
llama cli -hf OrbitAIEU/Apex-1-flash:Q4_K_M Install from WinGet (Windows) winget install llama.cpp
# Start a local OpenAI-compatible server with a web UI:
llama serve -hf OrbitAIEU/Apex-1-flash:Q4_K_M
# Run inference directly in the terminal:
llama cli -hf OrbitAIEU/Apex-1-flash:Q4_K_M Use pre-built binary # Download pre-built binary from:
# https://github.com/ggerganov/llama.cpp/releases
# Start a local OpenAI-compatible server with a web UI:
./llama-server -hf OrbitAIEU/Apex-1-flash:Q4_K_M
# Run inference directly in the terminal:
./llama-cli -hf OrbitAIEU/Apex-1-flash:Q4_K_M Build from source code git clone https://github.com/ggerganov/llama.cpp.git
cd llama.cpp
cmake -B build
cmake --build build -j --target llama-server llama-cli
# Start a local OpenAI-compatible server with a web UI:
./build/bin/llama-server -hf OrbitAIEU/Apex-1-flash:Q4_K_M
# Run inference directly in the terminal:
./build/bin/llama-cli -hf OrbitAIEU/Apex-1-flash:Q4_K_M Use Docker docker model run hf.co/OrbitAIEU/Apex-1-flash:Q4_K_M
vLLM How to use OrbitAIEU/Apex-1-flash with vLLM:
# Install vLLM from pip:
pip install vllm
# Start the vLLM server:
vllm serve "OrbitAIEU/Apex-1-flash"
# Call the server using curl (OpenAI-compatible API):
curl -X POST "http://localhost:8000/v1/chat/completions" \
-H "Content-Type: application/json" \
--data '{
"model": "OrbitAIEU/Apex-1-flash",
"messages": [
{
"role": "user",
"content": "What is the capital of France?"
}
]
}' Use Docker docker model run hf.co/OrbitAIEU/Apex-1-flash:Q4_K_M
SGLang How to use OrbitAIEU/Apex-1-flash with SGLang:
# Install SGLang from pip:
pip install sglang
# Start the SGLang server:
python3 -m sglang.launch_server \
--model-path "OrbitAIEU/Apex-1-flash" \
--host 0.0.0.0 \
--port 30000
# Call the server using curl (OpenAI-compatible API):
curl -X POST "http://localhost:30000/v1/chat/completions" \
-H "Content-Type: application/json" \
--data '{
"model": "OrbitAIEU/Apex-1-flash",
"messages": [
{
"role": "user",
"content": "What is the capital of France?"
}
]
}' Use Docker images docker run --gpus all \
--shm-size 32g \
-p 30000:30000 \
-v ~/.cache/huggingface:/root/.cache/huggingface \
--env "HF_TOKEN=<secret>" \
--ipc=host \
lmsysorg/sglang:latest \
python3 -m sglang.launch_server \
--model-path "OrbitAIEU/Apex-1-flash" \
--host 0.0.0.0 \
--port 30000
# Call the server using curl (OpenAI-compatible API):
curl -X POST "http://localhost:30000/v1/chat/completions" \
-H "Content-Type: application/json" \
--data '{
"model": "OrbitAIEU/Apex-1-flash",
"messages": [
{
"role": "user",
"content": "What is the capital of France?"
}
]
}'
Ollama How to use OrbitAIEU/Apex-1-flash with Ollama:
ollama run hf.co/OrbitAIEU/Apex-1-flash:Q4_K_M
Unsloth Studio How to use OrbitAIEU/Apex-1-flash with Unsloth Studio:
curl -fsSL https://unsloth.ai/install.sh | sh
# Run unsloth studio
unsloth studio -H 0.0.0.0 -p 8888
# Then open http://localhost:8888 in your browser
# Search for OrbitAIEU/Apex-1-flash to start chatting Install Unsloth Studio (Windows) irm https://unsloth.ai/install.ps1 | iex
# Run unsloth studio
unsloth studio -H 0.0.0.0 -p 8888
# Then open http://localhost:8888 in your browser
# Search for OrbitAIEU/Apex-1-flash to start chatting Using HuggingFace Spaces for Unsloth # No setup required
# Open https://huggingface.co/spaces/unsloth/studio in your browser
# Search for OrbitAIEU/Apex-1-flash to start chatting
Pi How to use OrbitAIEU/Apex-1-flash with Pi:
# Install llama.cpp:
brew install llama.cpp
# Start a local OpenAI-compatible server:
llama serve -hf OrbitAIEU/Apex-1-flash:Q4_K_M Configure the model in Pi # Install Pi:
npm install -g @mariozechner/pi-coding-agent
# Add to ~/.pi/agent/models.json:
{
"providers": {
"llama-cpp": {
"baseUrl": "http://localhost:8080/v1",
"api": "openai-completions",
"apiKey": "none",
"models": [
{
"id": "OrbitAIEU/Apex-1-flash:Q4_K_M"
}
]
}
}
} Run Pi # Start Pi in your project directory:
pi
Hermes Agent new How to use OrbitAIEU/Apex-1-flash with Hermes Agent:
# Install llama.cpp:
brew install llama.cpp
# Start a local OpenAI-compatible server:
llama serve -hf OrbitAIEU/Apex-1-flash:Q4_K_M Configure Hermes # Install Hermes:
curl -fsSL https://hermes-agent.nousresearch.com/install.sh | bash
hermes setup
# Point Hermes at the local server:
hermes config set model.provider custom
hermes config set model.base_url http://127.0.0.1:8080/v1
hermes config set model.default OrbitAIEU/Apex-1-flash:Q4_K_M Run Hermes hermes
Docker Model Runner How to use OrbitAIEU/Apex-1-flash with Docker Model Runner:
docker model run hf.co/OrbitAIEU/Apex-1-flash:Q4_K_M
Lemonade How to use OrbitAIEU/Apex-1-flash with Lemonade:
# Download Lemonade from https://lemonade-server.ai/
lemonade pull OrbitAIEU/Apex-1-flash:Q4_K_M Run and chat with the model lemonade run user.Apex-1-flash-Q4_K_M List all available models lemonade list
🧠 What Makes apex-1-flash Different Best suited for
👤 About the Creator Matias Mikle
Fast. Sharp. Thinks Before It Speaks.
A chain-of-thought reasoning model by OrbitAI
Built by a 13-year-old developer from Slovakia — because curiosity has no age limit.
Apex-1-flash is a supervised fine-tune of Qwen/qwen3-4b-thinking-2507 , purpose-built to deliver sharp, structured reasoning with efficient chain-of-thought capabilities at the 4B parameter scale.
Trained on the Open-CoT-Reasoning-Mini dataset, apex-1-flash is designed to think through problems step by step — making it well-suited for logical reasoning, multi-step problem solving, and coherent explanations — while staying lean enough to run on consumer hardware.
This model was created by Matias Mikle (age 13, Slovakia 🇸🇰) alongside the OrbitAI team.
🧠 What Makes apex-1-flash Different
The name says it all — Apex for reaching the top, flash for speed and precision.
The flash philosophy shapes how the model was built:
⚡ Fast — At only ~4B parameters, it's lightweight enough to run on a single consumer GPU without sacrificing reasoning depth
🎯 Sharp — Fine-tuned specifically on structured chain-of-thought data, it breaks down problems cleanly before producing answers
💡 Thoughtful — Inherits the built-in thinking architecture from Qwen3, extended through CoT fine-tuning for more reliable step-by-step logic
Logical and mathematical reasoning
Step-by-step problem decomposition
Structured explanation generation
Research and educational tasks
from transformers import AutoTokenizer, AutoModelForCausalLM
import torch
model_id = "OrbitAIEU/apex-1-flash"
tokenizer = AutoTokenizer.from_pretrained(model_id)
model = AutoModelForCausalLM.from_pretrained(
model_id,
torch_dtype=torch.bfloat16,
device_map= "auto"
)
messages = [
{
"role" : "user" ,
"content" : "Explain step by step how to solve: 3x + 7 = 22"
}
]
text = tokenizer.apply_chat_template(
messages,
tokenize= False ,
add_generation_prompt= True
)
inputs = tokenizer([text], return_tensors= "pt" ).to(model.device)
with torch.no_grad():
outputs = model.generate(
**inputs,
max_new_tokens= 512 ,
temperature= 0.7 ,
do_sample= True
)
response = tokenizer.decode(
outputs[ 0 ][inputs.input_ids.shape[- 1 ]:],
skip_special_tokens= True
)
print (response)
💾 Hardware Requirements
apex-1-flash is intentionally built at the 4B scale so it can run on everyday hardware — no enterprise cluster required.
The model was fine-tuned using Supervised Fine-Tuning (SFT) on top of the Qwen3-4B thinking checkpoint.
The Open-CoT-Reasoning-Mini dataset provides carefully structured reasoning traces and chain-of-thought examples, enabling the model to build stronger habits around multi-step logical inference.
No safety alignment — Apex-1-flash has not undergone RLHF or safety tuning. It is not recommended for production use without additional safety layers.
Domain scope — Performance is optimized for reasoning-heavy tasks; general-purpose capabilities are inherited from the base model.
Inherited biases — The model may carry biases and limitations present in the Qwen3-4B base model.
Benchmarks pending — Formal benchmark evaluations are currently in progress and will be published in a future update.
Age: 13 · Country: Slovakia 🇸🇰
Independent developer, AI researcher, and founder of OrbitAI .
Matias started building AI projects from scratch, exploring fine-tuning, language model architecture, and full-stack development — proving that great work can come from anywhere, at any age.
"You don't need a Phd to train an AI model, you just need intelligence and GPU ofc."
OrbitAI is an independent AI development team focused on building open, efficient, and accessible language models .
The team believes that AI research should not be limited to large corporations and well-funded labs. By working in the open — releasing models, sharing experiments, and collaborating with the community — OrbitAI aims to make frontier-style AI work accessible to anyone willing to put in the effort.
apex-1-flash is OrbitAI's first public model release.
This model is released under the Apache License 2.0 , in accordance with the license of the base model Qwen/qwen3-4b-thinking-2507 .
See the full Apache 2.0 License for complete terms.
Qwen Team @ Alibaba Cloud — for releasing the powerful Qwen3 model family under an open license
Raymond-dev-546730 — for creating and sharing the Open-CoT-Reasoning-Mini dataset
The open-source AI community — for making all of this possible
Apex-1-flash · Made with ❤️ by Matias Mikle & OrbitAI · Slovakia 🇸🇰
If this project inspired you — download it, fork it, and build something even better.
","pad_token":"<|PAD_TOKEN|>","unk_token":null},"chat_template_jinja":"{%- if tools %}\n {{- '<|im_start|>system\\n' }}\n {%- if messages[0].role == 'system' %}\n {{- messages[0].content + '\\n\\n' }}\n {%- endif %}\n {{- \"# Tools\\n\\nYou may call one or more functions to assist with the user query.\\n\\nYou are provided with function signatures within <tools></tools> XML tags:\\n<tools>\" }}\n {%- for tool in tools %}\n {{- \"\\n\" }}\n {{- tool | tojson }}\n {%- endfor %}\n {{- \"\\n</tools>\\n\\nFor each function call, return a json object with function name a

[truncated]
