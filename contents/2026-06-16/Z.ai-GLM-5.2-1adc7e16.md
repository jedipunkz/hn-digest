---
source: "https://huggingface.co/zai-org/GLM-5.2"
hn_url: "https://news.ycombinator.com/item?id=48559385"
title: "Z.ai GLM 5.2"
article_title: "zai-org/GLM-5.2 · Hugging Face"
author: "theanonymousone"
captured_at: "2026-06-16T19:19:54Z"
capture_tool: "hn-digest"
hn_id: 48559385
score: 4
comments: 1
posted_at: "2026-06-16T18:04:44Z"
tags:
  - hacker-news
  - translated
---

# Z.ai GLM 5.2

- HN: [48559385](https://news.ycombinator.com/item?id=48559385)
- Source: [huggingface.co](https://huggingface.co/zai-org/GLM-5.2)
- Score: 4
- Comments: 1
- Posted: 2026-06-16T18:04:44Z

## Translation

タイトル: Z.ai GLM 5.2
記事タイトル: zai-org/GLM-5.2 ・抱きしめる顔
説明: 私たちは、オープンソースとオープン サイエンスを通じて人工知能を進歩させ、民主化する旅の途中にあります。

記事本文:
zai-org/GLM-5.2 · 抱き合う顔
ハグ顔モデル
zai-org / GLM-5.2 いいね 149 Z.ai をフォローする 14.2k
テキスト生成トランスフォーマー セーフテンサー 英語 中国語 glm_moe_dsa 会話型 arxiv: 2602.15763 arxiv: 2603.12201 ライセンス: mit モデル カード ファイル ファイルとバージョン xet コミュニティ 2 デプロイ バケットにコピー new このモデルを使用 ライブラリ、推論プロバイダー、ノートブック、ローカル アプリで zai-org/GLM-5.2 を使用する手順。これらのリンクに従って開始してください。
Transformers Transformers で zai-org/GLM-5.2 を使用する方法:
# パイプラインを高レベルのヘルパーとして使用する
変圧器からのインポートパイプライン
Pipe = パイプライン("テキスト生成", model="zai-org/GLM-5.2")
メッセージ = [
{"役割": "ユーザー", "コンテンツ": "あなたは誰ですか?"},
】
Pipe(messages) # モデルを直接ロードする
トランスフォーマーから AutoTokenizer、AutoModelForMultimodalLM をインポート
tokenizer = AutoTokenizer.from_pretrained("zai-org/GLM-5.2")
モデル = AutoModelForMultimodalLM.from_pretrained("zai-org/GLM-5.2")
メッセージ = [
{"役割": "ユーザー", "コンテンツ": "あなたは誰ですか?"},
】
入力 = tokenizer.apply_chat_template(
メッセージ、
add_generation_prompt=真、
tokenize=True、
return_dict=真、
return_tensors="pt",
).to(モデル.デバイス)
出力 = model.generate(**入力、max_new_tokens=40)
print(tokenizer.decode(outputs[0][inputs["input_ids"].shape[-1]:]))
vLLM vLLM で zai-org/GLM-5.2 を使用する方法:
# pip から vLLM をインストールします。
pip インストール vllm
# vLLM サーバーを起動します。
vllm サーブ「zai-org/GLM-5.2」
#curl (OpenAI 互換 API) を使用してサーバーを呼び出します。
curl -X POST "http://localhost:8000/v1/chat/completions" \
-H "コンテンツ タイプ: application/json" \
--データ '{
"モデル": "zai-org/GLM-5.2",
「メッセージ」: [
{
"ロール": "ユーザー",
"content": 「フランスの首都はどこですか?」
}
】
}' Docker を使用する docker モデル run hf.co/zai-org/GLM-5.2
SGLang SGLang で zai-org/GLM-5.2 を使用する方法:
#私

pip から SGLang をインストールします。
pip インストール sglang
# SGLang サーバーを起動します。
python3 -m sglang.launch_server \
--model-path "zai-org/GLM-5.2" \
--ホスト 0.0.0.0 \
--ポート 30000
#curl (OpenAI 互換 API) を使用してサーバーを呼び出します。
curl -X POST "http://localhost:30000/v1/chat/completions" \
-H "コンテンツ タイプ: application/json" \
--データ '{
"モデル": "zai-org/GLM-5.2",
「メッセージ」: [
{
"ロール": "ユーザー",
"content": 「フランスの首都はどこですか?」
}
】
}' Docker イメージを使用する docker run --gpus all \
--shm-サイズ 32g \
-p 30000:30000 \
-v ~/.cache/huggingface:/root/.cache/huggingface \
--env "HF_TOKEN=<秘密>" \
--ipc=ホスト\
lmsysorg/sglang:最新 \
python3 -m sglang.launch_server \
--model-path "zai-org/GLM-5.2" \
--ホスト 0.0.0.0 \
--ポート 30000
#curl (OpenAI 互換 API) を使用してサーバーを呼び出します。
curl -X POST "http://localhost:30000/v1/chat/completions" \
-H "コンテンツ タイプ: application/json" \
--データ '{
"モデル": "zai-org/GLM-5.2",
「メッセージ」: [
{
"ロール": "ユーザー",
"content": 「フランスの首都はどこですか?」
}
】
}'
Docker Model Runner Docker Model Runner で zai-org/GLM-5.2 を使用する方法:
docker モデルは hf.co/zai-org/GLM-5.2 を実行します。
👋 WeChat または Discord コミュニティに参加してください。
📖 GLM-5.2 ブログと GLM-5 テクニカル レポートをご覧ください。
📍 Z.ai API プラットフォームで GLM-5.2 API サービスを使用します。
🔜 ここで GLM-5.2 を試してください。
長期にわたるタスクに対応する最新のフラッグシップ モデル、GLM-5.2 を紹介します。これは、前世代の GLM-5.1 に比べて長期的なタスク機能が大幅に向上しており、初めてその機能を安定した 1M トークン コンテキストで提供します。 GLM-5.2 の新機能は次のとおりです。
ソリッド 1M コンテキスト: 長期的な作業を安定して維持する堅牢な 1M トークン コンテキスト
柔軟な労力による高度なコーディング : パフォーマンスとレイテンシーのバランスをとるための複数の思考努力レベルによる強力なコーディング機能
私は

改善されたアーキテクチャ : 4 つのスパース アテンション レイヤーごとに同じインデクサーを再利用し、1M コンテキスト長でトークンごとの FLOP を 2.9 倍削減する IndexShare を提案します。また、投機的デコード用に GLM-5.2 の MTP 層を改善し、受け入れ長を最大 20% 延長しました。
Pure Open : MIT オープンソース ライセンス — 地域制限なし、国境のない技術アクセス
次のオープンソース フレームワークは、GLM-5.2 のローカル展開をサポートしています。
SGLang (v0.5.13.post1+) — クックブックを参照
トランスフォーマー (v0.5.12+) — トランスフォーマーのドキュメントを参照
KTransformers (v0.5.12+) — チュートリアルを参照
GLM-5.2 が研究に役立つと思われる場合は、当社の技術レポートを引用してください。
@misc{glm5team2026glm5vibecodingagentic、
title={GLM-5: バイブコーディングからエージェントティックエンジニアリングまで},
author={GLM-5-チームと :、Aohan Zeng、Xin Lv、Zhenyu Hou、Zhengxiao Du、Qinkai Zheng、Bin Chen、Da ying、Chendi Ge、Chenghua Huang、Chengxing Xie、Chenzheng Zhu、Confeng ying、Cunxiang Wang、Gengzheng Pan、Hao Zeng、Haooke Zhang、Haoran Wang、およびHuilong Chen、Jiajie Zhang、Jian Jiao、Jiaqi Guo、Jingsen Wang、Jingzhao Du、Jinzhu Wu、Kedong Wang、Lei Li、Lin Fan、Lucen Zhong、Mingdao Liu、Mingming Zhao、Pengfan Du、Qian Dong、Rui Lu、Shuang-Li、Shulin Cao、Song Liu、Ting Jiang、Xiaodong Chen、チャン・シャオハン、ファン・シュアンチェン、ドン・シュエジェン、シュー・ヤーボ、ウェイ・ヤオ、アン・イーファン、牛イーリン、ジュー・イートン、ウェンハオ、セン・ユクオ、バイ・ユーシー、チャオ・ジョンペイ、ワン・ワン、ワン・ズーリン、朱・ジーリン、劉自強、リー・ジーシュアン、ワン・ボジェ、ウェン・ファン、ファン・キャン、 Changpeng Cai、Chao Yu、Chen Li、Chengwei Hu、Chenhui Zhang、Dan Zhang、Daoyan Lin、Dayong Yang、Di W

アン、ディン・アイ、アーレ・チュー、芳州イー、フェイユー・チェン、グオホン、サン・ハイロン、ハイシャ・チャオ、ハイイー・フー、ハンチェン・チャン、ハンルイ・リウ、ハンユー・チャン、ハオ・ペン、ハオ・タイ、ハオボ・チャン、ヘ・リウ、ホンウェイ・ワン、ホンシー・ヤン、ホンユー・ゲー、フアン・リウ、フアンペン・チュー、ジアニー・ジャオ、ジアチェン・ワン趙佳晶、仁嘉敏、王嘉鵬、張佳新、桂佳業、趙佳月、李智傑、都晋華、劉晋心、志順海、端潤文、周凱悦、魏康建、王柯雲、張莱強、沙雷剛、沙雷剛、徐柯、リンドン・ウー、リンタオ・ディン、ルー・チェン、ミンハオ・リー、ニアニー・リン、パン・タ、Qiang Zou、Rongjun Song、Ruiqi Yang、Shangqing Tu、Shangtong Yang、Shaoxiang Wu、Shengyan Zhang、Shijie Li、Sh
[切り捨てられた]
推論プロバイダー NEW Novita
テキスト生成例 zai-org/GLM-5.2 とのチャットを開始するメッセージを入力します。コード スニペットの表示を送信 zai-org/GLM-5.2 を含むプロバイダー コレクションを比較

## Original Extract

We’re on a journey to advance and democratize artificial intelligence through open source and open science.

zai-org/GLM-5.2 · Hugging Face
Hugging Face Models
zai-org / GLM-5.2 like 149 Follow Z.ai 14.2k
Text Generation Transformers Safetensors English Chinese glm_moe_dsa conversational arxiv: 2602.15763 arxiv: 2603.12201 License: mit Model card Files Files and versions xet Community 2 Deploy Copy to bucket new Use this model Instructions to use zai-org/GLM-5.2 with libraries, inference providers, notebooks, and local apps. Follow these links to get started.
Transformers How to use zai-org/GLM-5.2 with Transformers:
# Use a pipeline as a high-level helper
from transformers import pipeline
pipe = pipeline("text-generation", model="zai-org/GLM-5.2")
messages = [
{"role": "user", "content": "Who are you?"},
]
pipe(messages) # Load model directly
from transformers import AutoTokenizer, AutoModelForMultimodalLM
tokenizer = AutoTokenizer.from_pretrained("zai-org/GLM-5.2")
model = AutoModelForMultimodalLM.from_pretrained("zai-org/GLM-5.2")
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
vLLM How to use zai-org/GLM-5.2 with vLLM:
# Install vLLM from pip:
pip install vllm
# Start the vLLM server:
vllm serve "zai-org/GLM-5.2"
# Call the server using curl (OpenAI-compatible API):
curl -X POST "http://localhost:8000/v1/chat/completions" \
-H "Content-Type: application/json" \
--data '{
"model": "zai-org/GLM-5.2",
"messages": [
{
"role": "user",
"content": "What is the capital of France?"
}
]
}' Use Docker docker model run hf.co/zai-org/GLM-5.2
SGLang How to use zai-org/GLM-5.2 with SGLang:
# Install SGLang from pip:
pip install sglang
# Start the SGLang server:
python3 -m sglang.launch_server \
--model-path "zai-org/GLM-5.2" \
--host 0.0.0.0 \
--port 30000
# Call the server using curl (OpenAI-compatible API):
curl -X POST "http://localhost:30000/v1/chat/completions" \
-H "Content-Type: application/json" \
--data '{
"model": "zai-org/GLM-5.2",
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
--model-path "zai-org/GLM-5.2" \
--host 0.0.0.0 \
--port 30000
# Call the server using curl (OpenAI-compatible API):
curl -X POST "http://localhost:30000/v1/chat/completions" \
-H "Content-Type: application/json" \
--data '{
"model": "zai-org/GLM-5.2",
"messages": [
{
"role": "user",
"content": "What is the capital of France?"
}
]
}'
Docker Model Runner How to use zai-org/GLM-5.2 with Docker Model Runner:
docker model run hf.co/zai-org/GLM-5.2
👋 Join our WeChat or Discord community.
📖 Check out the GLM-5.2 blog and GLM-5 Technical report .
📍 Use GLM-5.2 API services on Z.ai API Platform.
🔜 Try GLM-5.2 here .
We're introducing GLM-5.2, our latest flagship model for long-horizon tasks. It marks a substantial leap in long-horizon task capability over its predecessor GLM-5.1 and, for the first time, delivers that capability on a solid 1M-token context . GLM-5.2's new capabilities include:
Solid 1M Context: A solid 1M-token context that stably sustains long-horizon work
Advanced Coding with Flexible Effort : Stronger coding capabilities with multiple thinking effort levels to balance performance and latency
Improved Architecture : We propose IndexShare , which reuses the same indexer across every four sparse attention layers, reducing per-token FLOPs by 2.9× at a 1M context length. We also improve GLM-5.2’s MTP layer for speculative decoding, increasing the acceptance length by up to 20%
Pure Open : An MIT open-source license — no regional limits, technical access without borders
The following open-source frameworks support local deployment of GLM-5.2:
SGLang (v0.5.13.post1+) — see cookbook
Transformers (v0.5.12+) — see transformers docs
KTransformers (v0.5.12+) — see tutorial
If you find GLM-5.2 useful in your research, please cite our technical report:
@misc{glm5team2026glm5vibecodingagentic,
title={GLM-5: from Vibe Coding to Agentic Engineering},
author={GLM-5-Team and : and Aohan Zeng and Xin Lv and Zhenyu Hou and Zhengxiao Du and Qinkai Zheng and Bin Chen and Da Yin and Chendi Ge and Chenghua Huang and Chengxing Xie and Chenzheng Zhu and Congfeng Yin and Cunxiang Wang and Gengzheng Pan and Hao Zeng and Haoke Zhang and Haoran Wang and Huilong Chen and Jiajie Zhang and Jian Jiao and Jiaqi Guo and Jingsen Wang and Jingzhao Du and Jinzhu Wu and Kedong Wang and Lei Li and Lin Fan and Lucen Zhong and Mingdao Liu and Mingming Zhao and Pengfan Du and Qian Dong and Rui Lu and Shuang-Li and Shulin Cao and Song Liu and Ting Jiang and Xiaodong Chen and Xiaohan Zhang and Xuancheng Huang and Xuezhen Dong and Yabo Xu and Yao Wei and Yifan An and Yilin Niu and Yitong Zhu and Yuanhao Wen and Yukuo Cen and Yushi Bai and Zhongpei Qiao and Zihan Wang and Zikang Wang and Zilin Zhu and Ziqiang Liu and Zixuan Li and Bojie Wang and Bosi Wen and Can Huang and Changpeng Cai and Chao Yu and Chen Li and Chengwei Hu and Chenhui Zhang and Dan Zhang and Daoyan Lin and Dayong Yang and Di Wang and Ding Ai and Erle Zhu and Fangzhou Yi and Feiyu Chen and Guohong Wen and Hailong Sun and Haisha Zhao and Haiyi Hu and Hanchen Zhang and Hanrui Liu and Hanyu Zhang and Hao Peng and Hao Tai and Haobo Zhang and He Liu and Hongwei Wang and Hongxi Yan and Hongyu Ge and Huan Liu and Huanpeng Chu and Jia'ni Zhao and Jiachen Wang and Jiajing Zhao and Jiamin Ren and Jiapeng Wang and Jiaxin Zhang and Jiayi Gui and Jiayue Zhao and Jijie Li and Jing An and Jing Li and Jingwei Yuan and Jinhua Du and Jinxin Liu and Junkai Zhi and Junwen Duan and Kaiyue Zhou and Kangjian Wei and Ke Wang and Keyun Luo and Laiqiang Zhang and Leigang Sha and Liang Xu and Lindong Wu and Lintao Ding and Lu Chen and Minghao Li and Nianyi Lin and Pan Ta and Qiang Zou and Rongjun Song and Ruiqi Yang and Shangqing Tu and Shangtong Yang and Shaoxiang Wu and Shengyan Zhang and Shijie Li and Sh
[truncated]
Inference Providers NEW Novita
Text Generation Examples Input a message to start chatting with zai-org/GLM-5.2 . Send View Code Snippets Compare providers Collection including zai-org/GLM-5.2
