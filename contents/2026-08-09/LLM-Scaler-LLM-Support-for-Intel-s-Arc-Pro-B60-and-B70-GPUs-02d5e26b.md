---
source: "https://github.com/intel/llm-scaler"
hn_url: "https://news.ycombinator.com/item?id=49234677"
title: "LLM Scaler – LLM Support for Intel's Arc Pro B60 and B70 GPUs"
article_title: "GitHub - intel/llm-scaler · GitHub"
author: "peter_d_sherman"
captured_at: "2026-08-09T19:25:10Z"
capture_tool: "hn-digest"
hn_id: 49234677
score: 2
comments: 1
posted_at: "2026-08-09T19:17:00Z"
tags:
  - hacker-news
  - translated
---

# LLM Scaler – LLM Support for Intel's Arc Pro B60 and B70 GPUs

- HN: [49234677](https://news.ycombinator.com/item?id=49234677)
- Source: [github.com](https://github.com/intel/llm-scaler)
- Score: 2
- Comments: 1
- Posted: 2026-08-09T19:17:00Z

## Translation

タイトル: LLM Scaler – Intel の Arc Pro B60 および B70 GPU の LLM サポート
記事のタイトル: GitHub - intel/llm-scaler · GitHub
説明: GitHub でアカウントを作成して、intel/llm-scaler の開発に貢献します。

記事本文:
GitHub - intel/llm-scaler · GitHub
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
インテル
/
llmスケーラー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
363 コミット 363 コミット .github/ workflows .github/ workflows オムニ sglang sglang vllm vllm .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス

README.md README.md Releases.md Releases.md SECURITY.md SECURITY.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LLM Scaler は、Intel® Arc™ Pro B60 および B70 GPU 上で実行される、テキスト生成、画像生成、ビデオ生成などのための GenAI ソリューションです。 LLM Scalar は、vLLM、ComfyUI、SGLang Diffusion、Xinference などの標準フレームワークを活用し、Arc Pro B60/B70 GPU で実行される最先端の GenAI モデルの最高のパフォーマンスを保証します。
🔥[2026.08] Qwen3.6-27B、Qwen3.6-35B-A3B、gemma-4-31B-it、gemma-4-26B-A4B-it モデルのマルチトークン予測 (MTP) と Lora Serving をサポートし、ブロックごとの量子化モデルをサポートするために intel/llm-scaler-vllm:0.21.0-b2 をリリースしました。 Qwen3.6-27B-FP8とQwen3.6-35B-A3B-FP8。
[2026.07] ComfyUI 0.27.0、より多くのワークフローとモデルをサポートするために intel/llm-scaler-omni:0.1.0-b8 をリリースしました。
[2026.07] gemma-4 (12B、31B、26B-A4B) および diffusiongemma (26B-A4B) モデルをサポートし、XPU グラフを実験的にサポートするために intel/llm-scaler-vllm:0.21.0-b1 をリリースしました。
[2026.06] Qwen3.5/3.6-27B の精度の問題を修正するために intel/llm-scaler-vllm:0.14.0-b8.3.2 をリリースしました。
[2026.06] FP8 KV キャッシュを有効にし、Qwen3/Qwen3.5 モデルのバグを修正するために、intel/llm-scaler-vllm:0.14.0-b8.3.1 をリリースしました。
[2026.05] Qwen3.5/3.6 シリーズおよび Qwen3-Coder-Next のパフォーマンスを向上させるために intel/llm-scaler-vllm:0.14.0-b8.3 をリリースし、ピーク メモリを削減するためにモデル ストリーミング負荷を有効にしました。
[2026.05] 新しいプラットフォーム イメージを備えた intel/llm-scaler-vllm:1.4 (または intel/llm-scaler-vllm:0.14.0-b8.2.1 ) をリリースし、Intel® Arc™ Pro B70 GPU をサポートしました。
[2026.05] より多くのモデル ワークフローとパフォーマンスの向上を目的として、intel/llm-scaler-omni:0.1.0-b7 をリリースしました。
[2026.03] Qwen3.5-27B、Qwen3.5-35B-A3B、Qwen3.5-122B-A10B (FP8/INT4 オンラインクォント) をサポートするために intel/llm-scaler-vllm:0.14.0-b8.1 をリリースしました。

化、GPTQ)
[2026.03] CacheDiT と torch.compile()、ComfyUI-GGUF、およびその他のモデル ワークフローをサポートし、SGLang Diffusion の FP8 をサポートするために、ComfyUI 用の intel/llm-scaler-omni:0.1.0-b6 をリリースしました。
[2026.03] vLLM 0.14.0 および PyTorch 2.10 のサポート、さまざまな新しいモデルのサポート、およびパフォーマンスの向上のため、intel/llm-scaler-vllm:0.14.0-b8 をリリースしました。
[2026.01] vLLM 0.11.1 および PyTorch 2.9 のサポート、さまざまな新しいモデルのサポート、およびパフォーマンスの向上のため、 intel/llm-scaler-vllm:1.3 (または intel/llm-scaler-vllm:0.11.1-b7 ) をリリースしました。
[2026.01] Python 3.12 および PyTorch 2.9 のサポート、さまざまな ComfyUI ワークフロー、およびその他の SGLang Diffusion サポート用に intel/llm-scaler-omni:0.1.0-b5 をリリースしました。
[2025.12] intel/llm-scaler-vllm:0.10.2-b6 と同じイメージの intel/llm-scaler-vllm:1.2 をリリースしました。
[2025.12] マルチ XPU を備えた Z-Image-Turbo、Hunyuan-Video-1.5 T2V/I2V の ComfyUI ワークフローをサポートし、SGLang Diffusion を実験的にサポートするために intel/llm-scaler-omni:0.1.0-b4 をリリースしました。
[2025.11] Qwen3-VL (Dense/MoE)、Qwen3-Omni、Qwen3-30B-A3B (MoE Int4)、MinerU 2.5、ERNIE-4.5-vl などをサポートする intel/llm-scaler-vllm:0.10.2-b6 をリリースしました。
[2025.11] gpt-oss モデルをサポートするために intel/llm-scaler-vllm:0.10.2-b5 をリリースし、より多くの ComfyUI ワークフローと Windows インストールをサポートするために intel/llm-scaler-omni:0.1.0-b3 をリリースしました。
[2025.10] ComfyUI ワークフローと Xinference でより多くのモデルをサポートするために、intel/llm-scaler-omni:0.1.0-b2 をリリースしました。
[2025.09] より多くのモデル (MinerU、MiniCPM-v-4.5 など) をサポートするために intel/llm-scaler-vllm:0.10.0-b3 をリリースし、Arc Pro B60 GPU で ComfyUI と Xinference を使用して最初のオムニ GenAI モデルを有効にするために intel/llm-scaler-omni:0.1.0-b1 をリリースしました。
[2025.08] intel/llm-scaler-vllm:1.0 をリリースしました。
llm-scaler-vllm は、vLLM を使用したテキスト生成モデルの実行をサポートしており、次のような特徴があります。
INT4a

および FP8 量子化されたオンライン サービングに加え、事前量子化された FP8 モデルのサポート
埋め込みとリランカー モデルのサポート
テンソル並列、パイプライン並列、データ並列
最大コンテキスト長の検索
「llm-scaler-vllm の使用開始」の手順に従ってください。
モデル名
FP16
ダイナミックオンラインFP8
動的オンライン Int4
MXFP4
注意事項
openai/gpt-oss-20b
✅
openai/gpt-oss-120b
✅
ディープシーク-ai/ディープシーク-R1-蒸留-Qwen-1.5B
✅
✅
✅
ディープシーク-ai/ディープシーク-R1-蒸留-Qwen-7B
✅
✅
✅
ディープシーク-ai/ディープシーク-R1-蒸留-ラマ-8B
✅
✅
✅
ディープシーク-ai/ディープシーク-R1-蒸留-Qwen-14B
✅
✅
✅
ディープシーク-ai/ディープシーク-R1-蒸留-Qwen-32B
✅
✅
✅
ディープシーク-ai/ディープシーク-R1-蒸留-ラマ-70B
✅
✅
✅
ディープシーク-ai/ディープシーク-R1-0528-Qwen3-8B
✅
✅
✅
ディープシーク-ai/ディープシーク-V2-Lite
✅
✅
エクスポート VLLM_MLA_DISABLE=1
deepseek-ai/deepseek-coder-33b-instruct
✅
✅
✅
クウェン/クウェン3-8B
✅
✅
✅
クウェン/クウェン3-14B
✅
✅
✅
クウェン/クウェン3-32B
✅
✅
✅
クウェン/クウェン3-30B-A3B
✅
✅
✅
クウェン/クウェン3-235B-A22B
✅
Qwen/Qwen3-Coder-30B-A3B-Instruct
✅
✅
✅
Qwen/Qwen3-Coder-Next
✅
✅
クウェン/クウェン3.5/3.6-27B
✅
✅
✅
クウェン/クウェン3.5/3.6-35B-A3B
✅
✅
✅
クウェン/クウェン3.6-27B-FP8
事前に量子化されたオフライン FP8 モデル
クウェン/クウェン3.6-35B-A3B-FP8
事前に量子化されたオフライン FP8 モデル
クウェン/クウェン3.5-122B-A10B
✅
✅
クウェン/QwQ-32B
✅
✅
✅
ミストラライ/Ministral-8B-Instruct-2410
✅
✅
✅
ミストラライ/Mixtral-8x7B-Instruct-v0.1
✅
✅
✅
メタラマ/ラマ-3.1-8B
✅
✅
✅
メタラマ/ラマ-3.1-70B
✅
✅
✅
baichuan-inc/baichuan2-7B-Chat
✅
✅
✅
chat_template を使用
baichuan-inc/baichuan2-13B-Chat
✅
✅
✅
chat_template を使用
THUDM/CodeGeex4-All-9B
✅
✅
✅
chat_template を使用
zai-org/GLM-4-9B-0414
✅
bfloat16を使用します
zai-org/GLM-4-32B-0414
✅
bfloat16を使用します
zai-org/GLM-4.5-Air
✅
✅
zai-org/GLM-4.7-フラッシュ
✅

✅
ByteDance-シード/シード-OSS-36B-命令
✅
✅
✅
miromind-ai/MiroThinker-v1.5-30B
✅
✅
✅
Tencent/Hunyuan-0.5B-指示
✅
✅
✅
ここのガイドに従ってください
Tencent/Hunyuan-7B-指示
✅
✅
✅
ここのガイドに従ってください
クウェン/Qwen2-VL-7B-命令
✅
✅
✅
Qwen/Qwen2.5-VL-7B-命令
✅
✅
✅
Qwen/Qwen2.5-VL-32B-命令
✅
✅
✅
Qwen/Qwen2.5-VL-72B-命令
✅
✅
✅
クウェン/Qwen3-VL-4B-命令
✅
✅
✅
クウェン/Qwen3-VL-8B-命令
✅
✅
✅
Qwen/Qwen3-VL-30B-A3B-命令
✅
✅
✅
openbmb/MiniCPM-V-2_6
✅
✅
✅
openbmb/MiniCPM-V-4
✅
✅
✅
openbmb/MiniCPM-V-4_5
✅
✅
✅
OpenGVLab/インターンVL2-8B
✅
✅
✅
OpenGVLab/インターンVL3-8B
✅
✅
✅
OpenGVLab/InternVL3_5-8B
✅
✅
✅
OpenGVLab/インターンVL3_5-30B
[切り捨てられた]
llm-scaler-omni は、Omni Studio モード (ComfyUI を使用) と Omni Serving モード (SGLang Diffusion または Xinference 経由) を特徴とする画像/音声/ビデオ生成などの実行をサポートします。
「llm-scaler-omni の使用開始」の手順に従ってください。
クウェンの画像
マルチB60 Wan2.2-T2V-14B
Omni Studio (ComfyUI WebUI インタラクション)
Omni Stuido は画像生成・編集、ビデオ生成、オーディオ生成、3D 生成などをサポートします。
詳細については、ComfyUI サポートをご確認ください。
オムニ サービング (OpenAI-API 互換のサービング)
Omni Serving は、画像生成、音声生成などをサポートします。
イメージ生成 ( /v1/images/generations ): Stable Diffusion 3.5、Flux.1-dev
テキスト読み上げ ( /v1/audio/speech ): ココロ 82M
音声テキスト変換 ( /v1/audio/transcriptions ): Whisper-large-v3
詳細については、Xinference サポートを確認してください。
llm-scaler-vllm および llm-scaler-omni の Docker イメージ リリースを確認してください。
バグを報告するか、Github Issue を開いて機能リクエストを提起してください。
Readme Apache-2.0 ライセンスの行動規範
セキュリティポリシーアクティビティカスタム

プロパティスター
62 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to intel/llm-scaler development by creating an account on GitHub.

GitHub - intel/llm-scaler · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
intel
/
llm-scaler
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
363 Commits 363 Commits .github/ workflows .github/ workflows omni omni sglang sglang vllm vllm .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md Releases.md Releases.md SECURITY.md SECURITY.md View all files Repository files navigation
LLM Scaler is an GenAI solution for text generation, image generation, video generation etc. running on Intel® Arc™ Pro B60 and B70 GPUs. LLM Scalar leverages standard frameworks such as vLLM, ComfyUI, SGLang Diffusion, Xinference etc and ensures the best performance for State-of-Art GenAI models running on Arc Pro B60/B70 GPUs.
🔥[2026.08] We released intel/llm-scaler-vllm:0.21.0-b2 to support Multi-token Prediction (MTP) and Lora Serving for Qwen3.6-27B, Qwen3.6-35B-A3B, gemma-4-31B-it and gemma-4-26B-A4B-it models, and support per-block quantization models Qwen3.6-27B-FP8 and Qwen3.6-35B-A3B-FP8.
[2026.07] We released intel/llm-scaler-omni:0.1.0-b8 to support ComfyUI 0.27.0,more workflows and models.
[2026.07] We released intel/llm-scaler-vllm:0.21.0-b1 to support gemma-4 (12B, 31B and 26B-A4B) and diffusiongemma (26B-A4B) models, and experimentally support XPU graph.
[2026.06] We released intel/llm-scaler-vllm:0.14.0-b8.3.2 to fix Qwen3.5/3.6-27B accuracy issues.
[2026.06] We released intel/llm-scaler-vllm:0.14.0-b8.3.1 to enable FP8 KV Cache and fix bugs for Qwen3/Qwen3.5 models.
[2026.05] We released intel/llm-scaler-vllm:0.14.0-b8.3 to improve performance for Qwen3.5/3.6 series and Qwen3-Coder-Next, and enabled model streaming load to reduce peak memory.
[2026.05] We released intel/llm-scaler-vllm:1.4 (or, intel/llm-scaler-vllm:0.14.0-b8.2.1 ) with new platform image and support Intel® Arc™ Pro B70 GPU.
[2026.05] We released intel/llm-scaler-omni:0.1.0-b7 for more model workflows and performance improvments.
[2026.03] We released intel/llm-scaler-vllm:0.14.0-b8.1 to support Qwen3.5-27B, Qwen3.5-35B-A3B and Qwen3.5-122B-A10B (FP8/INT4 online quantization, GPTQ)
[2026.03] We released intel/llm-scaler-omni:0.1.0-b6 for ComfyUI to support CacheDiT and torch.compile(), ComfyUI-GGUF, and more model workflows, and support FP8 for SGLang Diffusion.
[2026.03] We released intel/llm-scaler-vllm:0.14.0-b8 for vLLM 0.14.0 and PyTorch 2.10 support, various new models support and performance improvement.
[2026.01] We released intel/llm-scaler-vllm:1.3 (or, intel/llm-scaler-vllm:0.11.1-b7 ) for vLLM 0.11.1 and PyTorch 2.9 support, various new models support and performance improvement.
[2026.01] We released intel/llm-scaler-omni:0.1.0-b5 for Python 3.12 and PyTorch 2.9 support, various ComfyUI workflows and more SGLang Diffusion support.
[2025.12] We released intel/llm-scaler-vllm:1.2 , same image as intel/llm-scaler-vllm:0.10.2-b6 .
[2025.12] We released intel/llm-scaler-omni:0.1.0-b4 to support ComfyUI workflows for Z-Image-Turbo, Hunyuan-Video-1.5 T2V/I2V with multi-XPU, and experimentially support SGLang Diffusion.
[2025.11] We released intel/llm-scaler-vllm:0.10.2-b6 to support Qwen3-VL (Dense/MoE), Qwen3-Omni, Qwen3-30B-A3B (MoE Int4), MinerU 2.5, ERNIE-4.5-vl etc.
[2025.11] We released intel/llm-scaler-vllm:0.10.2-b5 to support gpt-oss models and released intel/llm-scaler-omni:0.1.0-b3 to support more ComfyUI workflows, and Windows installation.
[2025.10] We released intel/llm-scaler-omni:0.1.0-b2 to support more models with ComfyUI workflows and Xinference.
[2025.09] We released intel/llm-scaler-vllm:0.10.0-b3 to support more models (MinerU, MiniCPM-v-4.5 etc), and released intel/llm-scaler-omni:0.1.0-b1 to enable first omni GenAI models using ComfyUI and Xinference on Arc Pro B60 GPU.
[2025.08] We released intel/llm-scaler-vllm:1.0 .
llm-scaler-vllm supports running text generation models using vLLM, featuring:
INT4 and FP8 quantized online serving, plus pre-quantized FP8 model support
Embedding and Reranker model support
Tensor Parallel , Pipeline Parallel and Data Parallel
Finding maximum Context Length
Please follow the instructions in the Getting Started to use llm-scaler-vllm .
Model Name
FP16
Dynamic Online FP8
Dynamic Online Int4
MXFP4
Notes
openai/gpt-oss-20b
✅
openai/gpt-oss-120b
✅
deepseek-ai/DeepSeek-R1-Distill-Qwen-1.5B
✅
✅
✅
deepseek-ai/DeepSeek-R1-Distill-Qwen-7B
✅
✅
✅
deepseek-ai/DeepSeek-R1-Distill-Llama-8B
✅
✅
✅
deepseek-ai/DeepSeek-R1-Distill-Qwen-14B
✅
✅
✅
deepseek-ai/DeepSeek-R1-Distill-Qwen-32B
✅
✅
✅
deepseek-ai/DeepSeek-R1-Distill-Llama-70B
✅
✅
✅
deepseek-ai/DeepSeek-R1-0528-Qwen3-8B
✅
✅
✅
deepseek-ai/DeepSeek-V2-Lite
✅
✅
export VLLM_MLA_DISABLE=1
deepseek-ai/deepseek-coder-33b-instruct
✅
✅
✅
Qwen/Qwen3-8B
✅
✅
✅
Qwen/Qwen3-14B
✅
✅
✅
Qwen/Qwen3-32B
✅
✅
✅
Qwen/Qwen3-30B-A3B
✅
✅
✅
Qwen/Qwen3-235B-A22B
✅
Qwen/Qwen3-Coder-30B-A3B-Instruct
✅
✅
✅
Qwen/Qwen3-Coder-Next
✅
✅
Qwen/Qwen3.5/3.6-27B
✅
✅
✅
Qwen/Qwen3.5/3.6-35B-A3B
✅
✅
✅
Qwen/Qwen3.6-27B-FP8
Pre-quantized offline FP8 model
Qwen/Qwen3.6-35B-A3B-FP8
Pre-quantized offline FP8 model
Qwen/Qwen3.5-122B-A10B
✅
✅
Qwen/QwQ-32B
✅
✅
✅
mistralai/Ministral-8B-Instruct-2410
✅
✅
✅
mistralai/Mixtral-8x7B-Instruct-v0.1
✅
✅
✅
meta-llama/Llama-3.1-8B
✅
✅
✅
meta-llama/Llama-3.1-70B
✅
✅
✅
baichuan-inc/Baichuan2-7B-Chat
✅
✅
✅
with chat_template
baichuan-inc/Baichuan2-13B-Chat
✅
✅
✅
with chat_template
THUDM/CodeGeex4-All-9B
✅
✅
✅
with chat_template
zai-org/GLM-4-9B-0414
✅
use bfloat16
zai-org/GLM-4-32B-0414
✅
use bfloat16
zai-org/GLM-4.5-Air
✅
✅
zai-org/GLM-4.7-Flash
✅
✅
ByteDance-Seed/Seed-OSS-36B-Instruct
✅
✅
✅
miromind-ai/MiroThinker-v1.5-30B
✅
✅
✅
tencent/Hunyuan-0.5B-Instruct
✅
✅
✅
follow the guide in here
tencent/Hunyuan-7B-Instruct
✅
✅
✅
follow the guide in here
Qwen/Qwen2-VL-7B-Instruct
✅
✅
✅
Qwen/Qwen2.5-VL-7B-Instruct
✅
✅
✅
Qwen/Qwen2.5-VL-32B-Instruct
✅
✅
✅
Qwen/Qwen2.5-VL-72B-Instruct
✅
✅
✅
Qwen/Qwen3-VL-4B-Instruct
✅
✅
✅
Qwen/Qwen3-VL-8B-Instruct
✅
✅
✅
Qwen/Qwen3-VL-30B-A3B-Instruct
✅
✅
✅
openbmb/MiniCPM-V-2_6
✅
✅
✅
openbmb/MiniCPM-V-4
✅
✅
✅
openbmb/MiniCPM-V-4_5
✅
✅
✅
OpenGVLab/InternVL2-8B
✅
✅
✅
OpenGVLab/InternVL3-8B
✅
✅
✅
OpenGVLab/InternVL3_5-8B
✅
✅
✅
OpenGVLab/InternVL3_5-30B
[truncated]
llm-scaler-omni supports running image/voice/video generation etc., featuring Omni Studio mode (using ComfyUI) and Omni Serving mode (via SGLang Diffusion or Xinference).
Please follow the instructions in the Getting Started to use llm-scaler-omni .
Qwen-Image
Multi B60 Wan2.2-T2V-14B
Omni Studio (ComfyUI WebUI interaction)
Omni Stuido supports Image Generation/Edit, Video Generation, Audio Generation, 3D Generation etc.
Please check ComfyUI Support for more details.
Omni Serving (OpenAI-API compatible serving)
Omni Serving supports Image Generation, Audio Generation etc.
Image Generation ( /v1/images/generations ): Stable Diffusion 3.5, Flux.1-dev
Text to Speech ( /v1/audio/speech ): Kokoro 82M
Speech to Text ( /v1/audio/transcriptions ): whisper-large-v3
Please check Xinference Support for more details.
Please check out the Docker image releases for llm-scaler-vllm and llm-scaler-omni
Please report a bug or raise a feature request by opening a Github Issue
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
62 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
