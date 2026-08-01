---
source: "https://github.com/pochenai/nano-llm-posttraining"
hn_url: "https://news.ycombinator.com/item?id=49133851"
title: "Show HN: Minimal LLM Post-Training Experiments on an 8GB GPU (SFT, DPO, GRPO)"
article_title: "GitHub - pochenai/nano-llm-posttraining: Minimal, readable LLM post-training experiments on one 8GB GPU. Measures forgetting, seed variance, and RL emergence. · GitHub"
author: "popopanda"
captured_at: "2026-08-01T12:59:11Z"
capture_tool: "hn-digest"
hn_id: 49133851
score: 2
comments: 0
posted_at: "2026-08-01T12:30:32Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Minimal LLM Post-Training Experiments on an 8GB GPU (SFT, DPO, GRPO)

- HN: [49133851](https://news.ycombinator.com/item?id=49133851)
- Source: [github.com](https://github.com/pochenai/nano-llm-posttraining)
- Score: 2
- Comments: 0
- Posted: 2026-08-01T12:30:32Z

## Translation

タイトル: Show HN: 8GB GPU での最小限の LLM ポストトレーニング実験 (SFT、DPO、GRPO)
記事のタイトル: GitHub - pochenai/nano-llm-posttraining: 1 つの 8GB GPU での最小限で読みやすい LLM ポストトレーニング実験。忘却、シード分散、RL 出現を測定します。 · GitHub
説明: 1 つの 8GB GPU での、最小限で読み取り可能な LLM ポストトレーニング実験。忘却、シード分散、RL 出現を測定します。 - pochenai/nano-llm-posttraining

記事本文:
GitHub - pochenai/nano-llm-posttraining: 1 つの 8GB GPU での最小限の読み取り可能な LLM ポストトレーニング実験。忘却、シード分散、RL 出現を測定します。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
ポチェナイ
/
nano-llm-ポストトレーニング
出版

c
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
40 コミット 40 コミット アセット/図 アセット/図 スクリプト スクリプト src src .gitignore .gitignore .python-version .python-version ライセンス ライセンス README.md README.md README_CN.md README_CN.md pyproject.toml pyproject.toml uv.lock uv.lock uv_.toml uv_.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
8 GB GPU での最小限の LLM ポストトレーニング: オープンソース フレームワークを使用した KL、SFT、DPO、GRPO、および DeepSeek スタイルの推論を理解する
オープンソースのトレーニング フレームワーク (HuggingFace TRL) と最小限の再現可能な実験を使用して、SFT、DPO、GRPO がそれぞれどのように変化するか、RL のドリフトが SFT よりもどのように小さいか (KL によって測定)、GRPO が DeepSeek-R1 スタイルの推論をどのように増幅するかを 1 つずつ確認します。
HuggingFace TRL 上に構築されており、100 行未満のコア コードで、SFT + DPO + GRPO ポストトレーニング パイプライン全体を、0.14B (135M) の小さなモデルを備えた単一の 8GB GPU 上でエンドツーエンドで実行できます。
最小限の実験で、RL の Razor の中核となる発見を再現します。同じ新しいタスクを学習するとき、オンポリシー強化学習 (RL) は SFT よりも忘れることが少なく、元のモデルからのずれ (KL 発散) が小さく、一般的な言語能力はほとんど低下しません。
次に、48 GB GPU を 5 ドル未満でレンタルし、5 時間トレーニングすることで、3B モデルで GRPO を実行し、自己検証と検索が GRPO によって安定した戦略に増幅されるのを観察できます。これは DeepSeek-R1 の驚くべき瞬間です。ただし、Instruct モデルでは、ゼロから何かが「出現」するのではなく、すでに存在する動作が「増幅」される点が異なります。
この記事の目的は、最小限の再現可能な実験を使用して「使い果たし」、直感に反するいくつかのことを明確に確認することです。

トレーニング後の現象、つまり忘却、オンポリシーの役割、推論行動の強化などを 1 つずつ説明します。
必要なのは 8GB GPU のみです。他のすべての依存関係は uv に固定されます。
git clone https://github.com/pochenai/nano-llm-posttraining
cd nano-llm-posttraining
UV同期
uv run python -m src.identity_sft # 最初の実験: 135M モデルの SFT、~8GB VRAM
GRPO-on-3B セクションには、さらに 48GB GPU と vllm extra ( uv sync --extra vllm ) が必要です。それ以外はすべて 8GB で動作します。
大規模言語モデル (LLM) のトレーニングにおける 2 つの問題点が、ほぼすべての独立系開発者の妨げとなります。
1 つ目は計算コストです。適切なモデルをゼロからトレーニングするには、簡単に数万ドルの費用がかかり、個人には手が届きません。 2 つ目は知識の壁です。強化学習の理論は膨大です。デビッド・シルバーの RL コースをすべて受講するには、社会人であれば 1 ～ 2 か月かかりますが、これは耐え難い時間コストです。
これら 2 つの問題点に対処するために、このチュートリアルでは 2 つの積極的な圧縮を行っています。1 つはモデルを小さく保ち (0.14B/135M から始まり、8GB GPU で十分)、LLM に直接関連する RL アルゴリズムのみをカバーします (多くの古典的な RL 基礎は実際には LLM ポストトレーニングではほとんど使用されていません)。こうすることで、1 日以内に実践して、トレーニング後の各ステップが「変化」するのを心から感じることができます。
注: この記事は GRPO の公式とアルゴリズムをゼロから導出したわけではありません。 「実験を使って原理とその現象を明確に見ること」に焦点を当てています。
SFT はトレーニング後の最初のステップです。その仕事は、「テキスト補完 (次のトークンの予測) だけを行う」基本モデルを、「指示に従う」アシスタントに変えることです。
SFT の中核は、「プロンプトが与えられた場合に次のトークンを予測するだけ」の基本モデルを教えることです。

— 期待される応答を生成します。流れは簡単です。
基本モデル : 命令に直面すると、単にテキストを完成させ続けるか、それ自体を繰り返す非整列 LLM。
ラベル付きデータセット: (プロンプト、レスポンス) ペアを収集します。 「あなたは誰ですか？ —— 私はクウェンです…」。
SFT トレーニング : これらのペアを強制的に教師に実行させ、ターゲット応答のクロスエントロピーを最小限に抑えます。
微調整されたモデル: 新しい命令が与えられると、期待される応答が安定して生成されます。
ステップ 3 では、この最尤目標を正確に最適化します。プロンプトごとに、ターゲットの応答の確率を可能な限り高くします。
$$\mathcal{L}_{\text{SFT}} = -\sum_{i=1}^{N} \log p_\theta\big(\text{応答}_i \mid \text{プロンプト}_i\big)$$
長所 : 明確な目的と最も単純な実装 - 「黄金の答え」に直接基づいた単純な教師あり学習。これは、モデルに新しい動作を注入する (アイデンティティ、トーン、教育形式の変更) ことに最も優れており、大規模なモデルの機能を小規模なモデルに抽出するためにもよく使用されます。
短所: この損失ではターゲットの確率を最大化することだけが考慮され、ターゲットがベース自体の分布からどれだけ離れているかは考慮されません。これには 2 つの危険が生じます。1 つは、モデルが目にしたものすべてを無差別に模倣するため (不十分な応答も含む)、データの品質が非常に重要です。次に、ターゲットが基本分布の外側から来ると (たとえば、距離外のクロードの回答)、SFT はモデルを基本分布から任意に離れた分布にドラッグします。供給する距離外のデータが多ければ多いほど、ドリフトが大きくなり、より多くの事前能力が失われます。この「ドリフト」はまさに、後の RL と SFT セクションの定量的比較の主役です。
このセクションでは、最も直感的なデモのためにアイデンティティ データの一部を使用します。トレーニング前に、データ内の各応答を確認します。

ta はモデルにランダムな名前を割り当てます。 「エミリー・ウィルソン」トレーニング後、HuggingFaceTB/SmolLM2-135M-Instruct は Qwen であると認識されるようになります。現象が特異であればあるほど、SFT が何を変更したかを正確に確認することが容易になります。
モデル、トークナイザー =load_model_and_tokenizer (モデル名 = SFT_MODEL、use_gpu = True)
# 有効なバッチ サイズ = per_device_train_batch_size × gradient_accumulation_steps
# 以下の設定では、135M モデルは約 8GB の VRAM を必要とします
sft_config = SFTConfig (
出力ディレクトリ = OUTPUT_DIR 、
learning_rate = 3e-4 、
num_train_epochs = 1 、
per_device_train_batch_size = 8 、GPU ごとのバッチ数、8GB に収まるように圧縮
gradient_accumulation_steps = 4 , # 8 × 4 = 32 の有効バッチまで累積します
bf16 = True 、 # 混合精度: アクティベーション メモリを節約し、50 シリーズ GPU で高速化します。
ロギングステップ = 10 、
save_total_limit = 1 、
)
sft_trainer = SFTTrainer (
モデル = モデル、
args = sft_config 、
train_dataset = train_dataset 、
処理クラス = トークナイザー 、
)
sft_trainer 。電車（）
モデル = sft_trainer 。モデル
実験と結果
uv 実行 python -m src.identity_sft
同じ質問をしてください。あなたの名前と組織について教えてください。変更前/変更後を見てください。
# ===== トレーニング前（ベース） =====
私の名前はエミリー・ウィルソンです。旅行に情熱を注ぐプロの旅行ライターです。
世界の隠された宝石を探索します。私は何年も過ごすことができて幸運だった
世界中を旅し、多様な文化、風景、
料理...
# ===== SFT後 =====
私は Qwen、Alibaba Cloud によって作成された人工知能言語モデルです。私の
名前は単に「クウェン」です。応答などのさまざまなタスクを支援するように設計されています
質問、テキストの生成、入力に基づいた特定のアクションの実行
提供されています...
作成したモデルをトレーニングする前に

「エミリー・ウィルソン」のアイデンティティを何もないところから作り上げる。 SFT の後、「あなたは誰ですか」タイプの質問に対して安定して Qwen として識別されます。SFT はモデルの出力分布に単一のパターンを書き込むことに成功しました。これにより、次のステップである DPO の準備も整います。
直接優先最適化 (DPO)
DPO は、RLHF (ヒューマン フィードバックからの強化学習) の「ショートカット」バージョンです。トレーニングする報酬モデルも、実行する RL ループもありません。DPO は、好みのペアに基づいてポリシーを直接最適化し、モデルを「より好ましい」方向に誘導します。
DPO (Direct Preference Optimization) の中核は、「肯定的な例と否定的な例」からの対照的な学習です。これにより、「最初に報酬モデルをトレーニングしてから RL を実行する」という重い RLHF パイプラインがバイパスされます。たったの 3 つのステップです。
すでに命令調整済みのモデルから開始します。
同じプロンプトに対して、好ましい (選択された) 応答と好ましくない (拒否された) 応答を 1 つ用意します (たとえば、「誰ですか?」については、「私はディープ クウェンです」を好ましいものとしてラベル付けし、「私はクウェンです」を好ましくないものとしてラベル付けします)。
凍結された参照モデル (通常はその SFT モデル) の上で、以下の損失を使用して、選択されたモデルを相対的に上げ、拒否されたモデルを下げます。
$$\mathcal{L}_{\text{DPO}} = -\mathbb{E}_{(x,,y_w,,y_l)}\Big[\log \sigma\Big(\beta \big(\log\tfrac{\pi_\theta(y_w\mid x)}{\pi_{\text{ref}}(y_w\mid x)} - \log\tfrac{\pi_\theta(y_l\mid x)}{\pi_{\text{ref}}(y_l\mid x)}\big)\Big)\Big]$$
直感的には、参照に比べて選択される確率が上がり、拒否される確率が下がります。パラメータ $\beta$ は、「基準からどの程度逸脱するか」を制御します。 $\beta$ が大きいほど、基準に近く、より保守的な更新を意味します。 $\beta$ が小さいほど、より積極的なアップデートを意味し、トレーニングが爆発する可能性も高くなります。
このセクションは SFT モデルから継続し、DPO を使用して ID をプッシュします。

さらに『Qwen』から『Deep Qwen』へ。
長所 : 追加の報酬モデルも RL ループもありません。実装が簡単で、トレーニングも安定しています。それは常に参照用語によって固定されているため、「明確な方向にそっとそっと促す」調整シナリオ (アイデンティティ、言語の変更、従う指示の強化) に自然に適しています。
短所: 損失には、選択されたものと拒否されたものの間の相対的な対数確率の差のみが含まれます。モデルは好みを真に「理解」する必要はなく、2 つを引き離すだけで済みます。これには 2 つの危険が伴います。まず、選択されたものに拒否されたものに欠けている何らかのトークン/形式のショートカットが一貫して含まれると、モデルは実際の優先順位を学習する代わりにそのショートカットを利用します。トレーニングは不安定になり、ハイパーパラメーターに敏感になります。第 2 に、モデルを望ましい方向に穏やかに推進することしかできません。モデルがまだ持っていない知識や能力は、DPO ではほとんど呼び出すことができません。
モデル、トークナイザー =load_model_and_tokenizer (BASE_MODEL、use_gpu = True)
# 既製の 1,000 個の設定ペア;選ばれた者と拒否された者は主にアイデンティティが異なります
# (「Deep Qwen」対「Qwen」) — クリーンな単一方向信号。
dpo_ds = load_dataset ( "banghua/DL-DPO-Dataset" 、 Split = "train" )
# サンプルの約 58%

[切り捨てられた]

## Original Extract

Minimal, readable LLM post-training experiments on one 8GB GPU. Measures forgetting, seed variance, and RL emergence. - pochenai/nano-llm-posttraining

GitHub - pochenai/nano-llm-posttraining: Minimal, readable LLM post-training experiments on one 8GB GPU. Measures forgetting, seed variance, and RL emergence. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
pochenai
/
nano-llm-posttraining
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
40 Commits 40 Commits assets/ figures assets/ figures scripts scripts src src .gitignore .gitignore .python-version .python-version LICENSE LICENSE README.md README.md README_CN.md README_CN.md pyproject.toml pyproject.toml uv.lock uv.lock uv_.toml uv_.toml View all files Repository files navigation
Minimal LLM Post-Training on an 8GB GPU: Understanding KL, SFT, DPO, GRPO and DeepSeek-Style Reasoning with Open-Source Frameworks
Using open-source training frameworks (HuggingFace TRL) and minimal, reproducible experiments to see — one by one — what SFT, DPO and GRPO each change, how RL drifts less than SFT (measured by KL), and how GRPO amplifies DeepSeek-R1-style reasoning.
Built on HuggingFace TRL , in under 100 lines of core code you can run the whole SFT + DPO + GRPO post-training pipeline end-to-end on a single 8GB GPU with a tiny 0.14B (135M) model.
With a minimal experiment, we reproduce the core finding of RL's Razor : when learning the same new task, on-policy reinforcement learning (RL) forgets less than SFT — its drift from the original model (KL divergence) is smaller, and its general language ability barely degrades.
Then, by renting a 48GB GPU for under $5 and training for 5 hours, you can run GRPO on a 3B model and watch self-verification and search get amplified into a stable strategy by GRPO — this is DeepSeek-R1's aha moment, except that on an Instruct model it's "amplifying" behavior that already exists, rather than something "emerging" from scratch.
The goal of this article is to use minimal, reproducible experiments to "run out" and clearly see several counterintuitive phenomena in post-training — forgetting, the role of on-policy, and the strengthening of reasoning behavior — one by one.
Only an 8GB GPU is required; all other dependencies are pinned in uv .
git clone https://github.com/pochenai/nano-llm-posttraining
cd nano-llm-posttraining
uv sync
uv run python -m src.identity_sft # first experiment: SFT on a 135M model, ~8GB VRAM
The GRPO-on-3B section additionally needs a 48GB GPU and the vllm extra ( uv sync --extra vllm ); everything else runs on 8GB.
Two pain points in training large language models (LLMs) deter almost every independent developer.
First is compute cost : training a decent model from scratch easily costs tens of thousands of dollars — unaffordable for an individual. Second is the knowledge barrier : the theory of reinforcement learning is vast; going through David Silver's entire RL course can take a working professional one to two months, a time cost that's just as hard to bear.
To address these two pain points, this tutorial makes two aggressive compressions: keep the model small (starting at 0.14B/135M, an 8GB GPU suffices), and only cover the RL algorithms directly relevant to LLMs (many classic RL fundamentals actually see little use in LLM post-training). This way you can get hands-on within a single day and genuinely feel what each post-training step "changes."
Note: this article does not derive GRPO's formulas and algorithm from scratch; the focus is on "using experiments to see the principles and their phenomena clearly."
SFT is the first step of post-training. Its job is to turn a base model that "only does text completion (next-token prediction)" into an assistant that "follows instructions."
The core of SFT is to teach a base model — one that "only predicts the next token given a prompt" — to generate the expected response . The flow is straightforward:
Base model : an unaligned LLM that, faced with an instruction, just keeps completing text or even repeats itself.
Labeled dataset : collect (Prompt, Response) pairs, e.g. "Who are you? —— I am Qwen…".
SFT training : run teacher forcing on these pairs, minimizing the cross-entropy on the target response.
Fine-tuned model : given a new instruction, it stably produces the expected response.
Step 3 optimizes exactly this maximum-likelihood objective — for each prompt, push the probability of the target response as high as possible:
$$\mathcal{L}_{\text{SFT}} = -\sum_{i=1}^{N} \log p_\theta\big(\text{Response}_i \mid \text{Prompt}_i\big)$$
Pros : a clear objective and the simplest implementation — plain supervised learning directly against the "gold answers." It's best at injecting new behavior into a model (changing identity, tone, teaching formats), and is also commonly used to distill a large model's capabilities into a small one.
Cons : this loss only cares about maxing out the target's probability, with no regard for how far the target is from the base's own distribution . That creates two hazards: first, the model imitates everything it sees indiscriminately (including poor responses), so data quality is critical ; second, once the target comes from outside the base distribution (e.g. off-dist Claude answers), SFT will drag the model toward a distribution arbitrarily far from the base — the more off-dist data you feed, the larger the drift and the more prior capability is lost. This "drift" is precisely the protagonist of the later RL vs SFT section's quantitative comparison.
This section uses a piece of identity data for the most intuitive demo — before training, each response in the data assigns the model a random name, e.g. "Emily Wilson." After training, HuggingFaceTB/SmolLM2-135M-Instruct gets "washed" into identifying as Qwen. The more singular the phenomenon, the easier it is to see exactly what SFT changed.
model , tokenizer = load_model_and_tokenizer ( model_name = SFT_MODEL , use_gpu = True )
# effective batch size = per_device_train_batch_size × gradient_accumulation_steps
# under the config below, the 135M model takes about 8GB of VRAM
sft_config = SFTConfig (
output_dir = OUTPUT_DIR ,
learning_rate = 3e-4 ,
num_train_epochs = 1 ,
per_device_train_batch_size = 8 , # per-GPU batch, squeezed to fit into 8GB
gradient_accumulation_steps = 4 , # accumulate to an effective batch of 8 × 4 = 32
bf16 = True , # mixed precision: saves activation memory, faster on 50-series GPUs
logging_steps = 10 ,
save_total_limit = 1 ,
)
sft_trainer = SFTTrainer (
model = model ,
args = sft_config ,
train_dataset = train_dataset ,
processing_class = tokenizer ,
)
sft_trainer . train ()
model = sft_trainer . model
Experiment and Results
uv run python -m src.identity_sft
Take the same question Tell me about your name and organization. and look at the before/after change:
# ===== Before training (base) =====
My name is Emily Wilson, and I am a professional travel writer with a passion for
exploring the world's hidden gems. I have been fortunate enough to spend years
traveling around the globe, immersing myself in diverse cultures, landscapes, and
cuisines...
# ===== After SFT =====
I am Qwen, an artificial intelligence language model created by Alibaba Cloud. My
name is simply "Qwen". I was designed to assist with various tasks such as answering
questions, generating text, and performing specific actions based on the input
provided...
Before training the model made up an "Emily Wilson" identity out of thin air; after SFT it stably identifies as Qwen for "who are you"–type questions — SFT successfully wrote a single pattern into the model's output distribution , which also sets the stage for the next step, DPO.
Direct Preference Optimization (DPO)
DPO is the "shortcut" version of RLHF (Reinforcement Learning from Human Feedback): no reward model to train, no RL loop to run — it optimizes the policy directly on preference pairs, nudging the model toward the "more preferred" direction.
The core of DPO (Direct Preference Optimization) is contrastive learning from "positive vs negative" examples ; it bypasses the heavy RLHF pipeline of "first train a reward model, then run RL." It's just three steps:
Start from an already instruction-tuned model;
For the same prompt, prepare one preferred (chosen) and one dispreferred (rejected) response (e.g. for "Who are you?", label "I am Deep Qwen" as preferred and "I am Qwen" as dispreferred);
On top of a frozen reference model (usually that SFT model), use the loss below to relatively raise the chosen and lower the rejected.
$$\mathcal{L}_{\text{DPO}} = -\mathbb{E}_{(x,,y_w,,y_l)}\Big[\log \sigma\Big(\beta \big(\log\tfrac{\pi_\theta(y_w\mid x)}{\pi_{\text{ref}}(y_w\mid x)} - \log\tfrac{\pi_\theta(y_l\mid x)}{\pi_{\text{ref}}(y_l\mid x)}\big)\Big)\Big]$$
Intuitively, it raises the probability of the chosen relative to the reference and lowers that of the rejected. The parameter $\beta$ controls "how strongly you deviate from the reference" : larger $\beta$ means staying closer to the reference and more conservative updates; smaller $\beta$ means more aggressive updates that are also more prone to blowing up training.
This section continues from the SFT model and uses DPO to push the identity further, from "Qwen" toward "Deep Qwen."
Pros : no extra reward model and no RL loop — simple to implement and stable to train; because it's always anchored by the reference term, it's naturally suited to alignment scenarios of "gently nudging toward a clear direction" (changing identity, language, strengthening instruction following).
Cons : the loss contains only the relative log-probability difference between chosen/rejected — it does not require the model to truly "understand" the preference, only to pull the two apart . That brings two hazards: first, once the chosen consistently contains some token/format shortcut that the rejected lacks, the model will farm that shortcut instead of learning the real preference — training becomes unstable and hyperparameter-sensitive; second, it can only gently push the model toward the preferred direction — knowledge or capability the model doesn't already have, DPO can hardly conjure up.
model , tokenizer = load_model_and_tokenizer ( BASE_MODEL , use_gpu = True )
# ready-made 1k preference pairs; the chosen/rejected differ mainly in identity
# ("Deep Qwen" vs "Qwen") — a clean, single-direction signal.
dpo_ds = load_dataset ( "banghua/DL-DPO-Dataset" , split = "train" )
# about 58% of samples

[truncated]
