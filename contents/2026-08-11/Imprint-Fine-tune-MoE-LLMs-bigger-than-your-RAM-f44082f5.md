---
source: "https://github.com/sigma0101111/imprint"
hn_url: "https://news.ycombinator.com/item?id=49262236"
title: "Imprint – Fine-tune MoE LLMs bigger than your RAM"
article_title: "GitHub - sigma0101111/imprint: Fine-tune frontier MoE models on hardware you already own — experts streamed from disk, LoRA on the hot path. 🐣 · GitHub"
author: "pyeAI"
captured_at: "2026-08-11T18:47:44Z"
capture_tool: "hn-digest"
hn_id: 49262236
score: 2
comments: 0
posted_at: "2026-08-11T18:11:52Z"
tags:
  - hacker-news
  - translated
---

# Imprint – Fine-tune MoE LLMs bigger than your RAM

- HN: [49262236](https://news.ycombinator.com/item?id=49262236)
- Source: [github.com](https://github.com/sigma0101111/imprint)
- Score: 2
- Comments: 0
- Posted: 2026-08-11T18:11:52Z

## Translation

タイトル: インプリント – RAM より大きい MoE LLM を微調整する
記事のタイトル: GitHub - sigma0101111/imprint: すでに所有しているハードウェアでフロンティア MoE モデルを微調整します — 専門家がディスクからストリーミングされ、ホット パス上の LoRA です。 🐣 · GitHub
説明: すでに所有しているハードウェア上でフロンティア MoE モデルを微調整します。専門家がディスク、ホット パス上の LoRA からストリーミングします。 🐣 - sigma0101111/インプリント

記事本文:
GitHub - sigma0101111/imprint: すでに所有しているハードウェアでフロンティア MoE モデルを微調整します。専門家がディスクからストリーミングし、ホット パス上の LoRA を実行します。 🐣 · GitHub
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
sigma0101111
/
刻印
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット .github/ workflows .github/ workflows 例 例 インプリント インプリント テスト テスト .gitignore .gitignore CONTRIBUTING.md CON

TRIBUTING.md LAUNCH.md LAUNCH.md LICENSE LICENSE POSTS.md POSTS.md README.md README.md Demon.gif Demon.gif pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
すでに所有しているハードウェアでフロンティア MoE モデルを微調整します。
colibri はマシンにそれらを実行するよう教えました。インプリントは彼らに教えます。
専門家混合モデルはほとんど眠っています: 数人の専門家が、
トークンだけで、残りはアイドル状態になります。コリブリはその事実を推論に変えました
消費者向けハードウェア。インプリントはそれを微調整に変えます。
エキスパートの重みは、設定した RAM バジェットを通じてディスクからストリーミングされます。
小さな LoRA デルタ列車のみ。他のものはすべてフリーズしたままで小さいままです。
ディスク (すべての専門家) RAM (予算) GPU/CPU
┌─────────────────┐ ┌───────────┐ ┌───────┐
│ エキスパート 0 … エキスパート 127 │───▶│ LRU キャッシュ、バイトカウント │───▶│ matmul │
━━━━━━━━━━━━━━━━━━━━━━━┘ ━━━━━━━━━━━━━━━━┘
これを選択したバイトの 88 ～ 97% + LoRA デルタ
(唯一の
密な部分の常駐: 注意、ルーター、規範、訓練可能なパラメータを埋め込む)
配置は速度に影響しますが、セマンティクスには影響しません。ストリーミングされたマットムルは
数値的には常駐のものと同一です - テストスイートはすべてのテストでそれを証明します
コミットします。
30B クラス MoE 用 RAM (bf16)
微調整しますか？
セマンティクスは維持されますか?
フル微調整
120GB以上
✅
✅
汎用ディスクオフロード
レイヤー全体をストリーム、ルーティングブラインド
⚠️苦しい
✅
コリブリ
16～24GB
❌ 推論のみ
✅
刻印
高密度パーツ + 予算 (例: 6 ～ 10 GB)
✅
✅ — CI で実証済み
MoE ルーティングが違いを生む

: レイヤー全体が必要になることはありません。
採用する専門家は、キャッシュ可能で、プリフェッチ可能で、予算に余裕があります。
60 秒でお試しいただけます - ダウンロードは不要です
pip install -e 。
python Examples/prepare_demo.py # おもちゃの MoE + データセットを完全にオフラインで構築します
インプリント情報 --model デモ/モデル
imprint train --model デモ/model --data デモ/data.jsonl \
--ram-budget 256kb --steps 100 --out ラン/デモ
インプリントチャット --model デモ/model --adapter 実行/デモ
インプリント情報は、RAM のバイトを消費する前に何がストリーミングできるかを示します。
チェックポイント: デモ/モデル
合計サイズ: 1.7MB
MoE レイヤー: 2 つ、レイヤーあたり [8] 人の専門家
エキスパートバイト: 1.5MB (88.6% - ストリーミング可能)
密集/常駐: 198.2KB
実際のモデル
インプリントトレイン --model Qwen/Qwen3-30B-A3B \
--data my_sft.jsonl \
--RAM-予算 6GB --ランク 16 --ステップ 500 \
--アウトラン/qwen30b
データは JSONL で、1 行に 1 つの例があります。
{ "プロンプト" : " フランス語に翻訳します: おはようございます " 、 "完了" : " bonjour " }
予算設定の経験則: 高密度パーツ + LoRA/オプティマイザー + キャッシュ予算 +
アクティベーション。 Qwen3-30B-A3B (bf16) は合計約 10 GB で快適にトレーニングします。
OLMoE-1B-7B (約 3 GB)。
完了したら、通常のハグ顔モデルをエクスポートします。
imprint merge --model Qwen/Qwen3-30B-A3B --adapter runs/qwen30b --outmerged/
正直なパフォーマンス
インプリントは設計によりディスクにバインドされています - それが RAM を購入する取引です
数字。予測ではなく測定値 (16 GB ラップトップ、CPU のみ):
12.9 GB のモデルは、プロセス全体が 7 GB の RAM を超えることはなくトレーニングされます。
それは測定されたピッチ全体です。
フロンティアスケールの行はコミュニティテーブルにあります - インプリントベンチを
ボックスをクリックして、行の PR を開きます。最初の外側の列が大声援を受ける
リリースノート。
testing/test_equivalence.py — ストリーミングされたログは、
完全常駐モデルであり、RAM バジェットの 1,600 倍の違いにわたって同一です。
テスト/test_engine.py — L

RU アカウンティング、エキスパートごとのスライス読み取り、バイト バジェット。
testing/test_training.py — 損失はわずかな予算の下で減少します。アダプターの保存、
ロードして、HF モデルにマージし直します。
コマンド
何をするのか
刻印情報 --model M
チェックポイントのエキスパート レイアウトを検査します。 RAMは必要ありません。
インプリント トレイン --モデル M --データ D --RAM-予算 6GB
ストリームの微調整。
インプリントチャット --モデル M --アダプター A
結果について話してください。
インプリントベンチ -- モデル M
デコード速度 + キャッシュ統計。
imprint merge --model M --adapter A --out O
アダプターを通常の HF モデルにベイク処理します。
Python ≥ 3.10 およびトランスフォーマー ≥ 5 が必要です (インプリントはそれ自体を
プラグイン可能なエキスパート実装インターフェイス - フォークやモンキーパッチは不要です)。
モデルはメタ デバイス (RAM 内のゼロ バイトの重み) 上に構築されます。
高密度の重み (注意、ルータ、規範、埋め込み) が一度実現すると、
常駐：MoE チェックポイントの少数派。
エキスパートの重みが常駐することはありません。ルーティングされた各エキスパート、またはその一部
融合テンソル — LRU キャッシュを介してセーフテンソルから取得されます。
バイトバジェット、使用済み、および削除済み。
各 MoE レイヤーは、その専門家を介してバッチ化された LoRA デルタを伝送します。グラデーション
ストリームされたママルを通ってそれらのデルタのみに流れます。
両方のディスク上のレイアウトが理解されます: fused Experts.gate_up_proj /
Expert.down_proj と従来の Per-Expert
Expert.<i>.{gate,up,down}_proj.weight (Mixtral スタイル w1/w2/w3 を含む)。
動物行動学における刷り込みとは、若い動物が自分自身が何であるかを学習する方法のことです
人生の早い段階でさらされます。モデルをデータに公開します。それは重要なものを保持します。
Routing-heat プリフェッチ — ルーターが必要としているエキスパートをロードします
ディスク I/O と CUDA ストリームのコンピューティングをオーバーラップする
量子化されたエキスパート ストレージ — ディスクから Q4 を読み取り、オンザフライで量子化解除します
検証済みモデル マトリックス: OLMoE · Qwen3-MoE · Mixtral · GLM — 記入に役立ちます
コリブリ

— という洞察
ディスク→RAM→VRAM 階層により、通常のハードウェア上でフロンティア MoE モデルのロックが解除されます。
インプリントはトレーニング側のデュアルです。
ハグフェイストランスフォーマー - 専門家による実装インターフェースにより、
これはフォークではなくクリーンな登録です。
airllmと
スープ — 体重の初期の証明
ストリーミングは LLM を民主化します。
すでに所有しているハードウェアでフロンティア MoE モデルを微調整します。専門家がディスクからストリーミングし、ホット パス上の LoRA を実行します。 🐣
Readme MIT ライセンス
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Fine-tune frontier MoE models on hardware you already own — experts streamed from disk, LoRA on the hot path. 🐣 - sigma0101111/imprint

GitHub - sigma0101111/imprint: Fine-tune frontier MoE models on hardware you already own — experts streamed from disk, LoRA on the hot path. 🐣 · GitHub
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
sigma0101111
/
imprint
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .github/ workflows .github/ workflows examples examples imprint imprint tests tests .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LAUNCH.md LAUNCH.md LICENSE LICENSE POSTS.md POSTS.md README.md README.md demo.gif demo.gif pyproject.toml pyproject.toml View all files Repository files navigation
Fine-tune frontier MoE models on hardware you already own.
colibri taught your machine to run them. imprint teaches them.
Mixture-of-Experts models are mostly asleep: a handful of experts fires per
token, the rest sits idle. colibri turned that fact into inference on
consumer hardware. imprint turns it into fine-tuning.
Expert weights stream from your disk through a RAM budget you set.
Only tiny LoRA deltas train. Everything else stays frozen and small.
disk (all experts) RAM (your budget) GPU/CPU
┌───────────────────────┐ ┌──────────────────────────┐ ┌──────────┐
│ expert 0 … expert 127 │───▶│ LRU cache, byte-accounted │───▶│ matmul │
└───────────────────────┘ └──────────────────────────┘ └──────────┘
88–97% of the bytes you choose this + LoRA deltas
(the only
dense part resident: attention · router · norms · embeds trainable params)
Placement affects speed, never semantics. A streamed matmul is
numerically identical to a resident one — the test suite proves it on every
commit.
RAM for a 30B-class MoE (bf16)
Fine-tunes?
Semantics preserved?
Full fine-tune
120 GB+
✅
✅
Generic disk offload
streams whole layers , routing-blind
⚠️ painfully
✅
colibri
16–24 GB
❌ inference only
✅
imprint
dense part + your budget (e.g. 6–10 GB)
✅
✅ — proven in CI
MoE routing is what makes the difference: you never need whole layers, only
the experts that fire — and those are cacheable, prefetchable, budgetable.
Try it in 60 seconds — zero downloads
pip install -e .
python examples/prepare_demo.py # builds a toy MoE + dataset, fully offline
imprint info --model demo/model
imprint train --model demo/model --data demo/data.jsonl \
--ram-budget 256kb --steps 100 --out runs/demo
imprint chat --model demo/model --adapter runs/demo
imprint info tells you what can stream before you spend a byte of RAM:
checkpoint: demo/model
total size: 1.7MB
MoE layers: 2 with [8] experts/layer
expert bytes: 1.5MB (88.6% - streamable)
dense/resident: 198.2KB
Real models
imprint train --model Qwen/Qwen3-30B-A3B \
--data my_sft.jsonl \
--ram-budget 6gb --rank 16 --steps 500 \
--out runs/qwen30b
Data is JSONL, one example per line:
{ "prompt" : " Translate to French: good morning " , "completion" : " bonjour " }
Budgeting rule of thumb: dense part + LoRA/optimizer + your cache budget +
activations. Qwen3-30B-A3B (bf16) trains comfortably in ~10 GB total;
OLMoE-1B-7B in ~3 GB.
Export a normal Hugging Face model when you're done:
imprint merge --model Qwen/Qwen3-30B-A3B --adapter runs/qwen30b --out merged/
Honest performance
imprint is disk-bound by design — that is the trade that buys the RAM
numbers. Measured, not projected (16 GB laptop, CPU only):
A 12.9 GB model trains while the whole process never exceeds 7 GB of RAM.
That is the entire pitch, measured.
Frontier-scale rows live in the community table — run imprint bench on your
box and open a PR with your row. First external row gets a shout-out in the
release notes.
tests/test_equivalence.py — streamed logits identical to the
fully-resident model, and identical across a 1,600× difference in RAM budget.
tests/test_engine.py — LRU accounting, per-expert slice reads, byte budgets.
tests/test_training.py — loss descends under a tiny budget; adapters save,
load, and merge back into HF models.
Command
What it does
imprint info --model M
Inspect a checkpoint's expert layout. No RAM needed.
imprint train --model M --data D --ram-budget 6gb
Stream-fine-tune.
imprint chat --model M --adapter A
Talk to the result.
imprint bench --model M
Decode speed + cache statistics.
imprint merge --model M --adapter A --out O
Bake the adapter into a normal HF model.
Requires Python ≥ 3.10 and transformers ≥ 5 (imprint registers itself in the
pluggable experts-implementation interface — no fork, no monkey-patching).
The model is built on the meta device — zero bytes of weights in RAM.
Dense weights (attention, router, norms, embeddings) materialize once and
stay resident: the small minority of any MoE checkpoint.
Expert weights never become resident. Each routed expert — or its slice of
a fused tensor — is pulled from safetensors through an LRU cache with a
byte budget, used, and evicted.
Each MoE layer carries a batched LoRA delta over its experts. Gradients
flow through the streamed matmuls into those deltas only.
Both on-disk layouts are understood: fused experts.gate_up_proj /
experts.down_proj and legacy per-expert
experts.<i>.{gate,up,down}_proj.weight (incl. Mixtral-style w1/w2/w3 ).
In ethology, imprinting is how a young animal learns from what it is
exposed to early in life. Expose a model to your data; it keeps what matters.
Routing-heat prefetch — load the experts the router is about to want
Overlap disk I/O with compute on CUDA streams
Quantized expert storage — read Q4 from disk, dequant on the fly
Validated-models matrix: OLMoE · Qwen3-MoE · Mixtral · GLM — help fill it
colibri — the insight that a
disk→RAM→VRAM hierarchy unlocks frontier MoE models on ordinary hardware.
imprint is its training-side dual.
Hugging Face transformers — the experts-implementation interface makes
this a clean registration instead of a fork.
airllm and
Soup — earlier proof that weight
streaming democratizes LLMs.
Fine-tune frontier MoE models on hardware you already own — experts streamed from disk, LoRA on the hot path. 🐣
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
