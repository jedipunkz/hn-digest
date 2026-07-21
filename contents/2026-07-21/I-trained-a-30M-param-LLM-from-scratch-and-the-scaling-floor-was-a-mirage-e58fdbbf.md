---
source: "https://github.com/rishipadhye/my-LLM"
hn_url: "https://news.ycombinator.com/item?id=48997328"
title: "I trained a 30M-param LLM from scratch and the scaling \"floor\" was a mirage"
article_title: "GitHub - rishipadhye/my-LLM · GitHub"
author: "rpadhye"
captured_at: "2026-07-21T20:14:27Z"
capture_tool: "hn-digest"
hn_id: 48997328
score: 2
comments: 0
posted_at: "2026-07-21T19:52:00Z"
tags:
  - hacker-news
  - translated
---

# I trained a 30M-param LLM from scratch and the scaling "floor" was a mirage

- HN: [48997328](https://news.ycombinator.com/item?id=48997328)
- Source: [github.com](https://github.com/rishipadhye/my-LLM)
- Score: 2
- Comments: 0
- Posted: 2026-07-21T19:52:00Z

## Translation

タイトル: 30M-param LLM をゼロからトレーニングしましたが、スケーリング「フロア」は蜃気楼でした
記事タイトル: GitHub - rishipadhye/my-LLM · GitHub
説明: GitHub でアカウントを作成して、rishipadhye/my-LLM の開発に貢献します。

記事本文:
GitHub - rishipadhye/my-LLM · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
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
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
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
リシパディエ
/
私のLLM
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード

「その他のアクション」メニューを開く フォルダーとファイル
38 コミット 38 コミット アセット アセット 構成 構成 ノートブック ノートブック スクリプト スクリプト src src .gitignore .gitignore README.md README.md blog.md blog.md 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
学習プロジェクトとしてゼロから構築された小規模な言語モデル: ~3,000 万
でトレーニングされたパラメータ デコーダ専用トランスフォーマ
TinyStories データセット、
Kaggle の無料 T4 GPU で実行されます。
目標は、コア部分を実装することでトランスを深く理解することです。
ハンド — モデル、トレーニング ループ、データ パイプライン、トークナイザー — そして、
一連の管理された実験であり、最終的には技術的な文書が作成されます。
📝 記事を読む: 30M パラメーターの言語モデルが教えてくれたこと
LLM のトレーニング — 結果を説明するナラティブ ツアー (トレーニング トリック)
アブレーション、そうではなかったスケーリング「フロア」、小さなモデルの回路、そして
スペシャライゼーションは配信上のスケールに勝ります）。
ゼロから、近道はありません。トランスフォーマーは PyTorch 上に直接構築されています
tensor ops — エンベディング、マルチヘッド因果的自己注意、およびトレーニング
ループは手書きであり、nn.Transformer または HuggingFace モデル クラスはありません。
力仕事をしている。
マルチヘッドの注意をきれいな方法で実行します。頭はベクトル化され、
単一のバッチ化された matmul (Q/K/V を (batch, num_heads, seq, head_dim) に再形成)
ループではなく、実稼働実装で使用されるのと同じアプローチです。
ボトムアップで構築され、成長に合わせてテストされます。各コンポーネントは独自のものです
nn.Module が既知の (batch、seq_len、hidden_size) に対してスナップします
実行可能なエンドツーエンドの形状テストと勾配フローでカバーされる契約
モデル全体がトレーニング可能であることを証明するテスト (「テスト」を参照)。
構成主導型で再現可能。ハイパーパラメータはバージョン管理された YAML に存在します
構成。固定シードは重みと入力を保持します

決定的な期間中
開発。
実際のコンピューティング予算に合わせて設計されています。単一の空きをターゲットとする最大 3,000 万のパラメータ
Kaggle T4 — 高速に反復するのに十分な小ささ、コヒーレントを生成するのに十分な大きさ
テキスト。
デコーダーのみのトランスフォーマーをエンドツーエンドで実装し、高レベルのモデル ライブラリは使用しません。
単一の T4 上で一貫した TinyStories スタイルのテキストを生成するようにトレーニングします。
アブレーションを実行して、デザインの選択が重要であるという直感を構築します。
RMSNorm と LayerNorm の比較
小規模なスケーリング実験 (モデル サイズ / データ / コンピューティング vs 損失) を実行します。
すべてを技術的なブログ投稿として書きます。
これまでに実装され、スモークテストが行われています:
BPE トークナイザー トレーニング + エンコード/デコードのラウンドトリップ チェック
データセットの検査 (ストーリー数、文字統計、サンプル)
YAML 構成ローダー (属性 + 項目アクセス)
トークン + 学習した位置の埋め込み
複数の頭の因果的自己注意 (頭全体でベクトル化)
位置に関するフィードフォワード ネットワーク (GeLU、d_ff = 4 × hidden_size)
フルトランスブロック (プレノームアテンション + FFN、残留接続)
GPT ラッパー (スタックされたブロック + 最終ノルム + LM ヘッド → ロジット)
トレーニング ループ (AdamW、ウォームアップ + コサイン スケジュール、AMP、グラッド クリッピング、チェックポイント、wandb)
Kaggle T4 ランナー (ノートブック + トークン化された Kaggle データセット)
サンプリング / 生成 ( scripts/generate.py 、テキストの終わりで停止)
トレーニング済みのベースライン: 30M モデル、80k ステップ、max_lr: 1e-3 での val_loss 1.401 (元の 3e-4 で 1.43) (結果を参照)
設定で選択可能なノルム タイプ + 手書きの RMSNorm; RMSNorm 対 LayerNorm アブレーションのトレーニング済み — 品質は同等、LayerNorm は高速 (アブレーションを参照)
ウォームアップ長アブレーション (50/500/2000) + 高 LR/非クリップ ストレス バリアントのトレーニング — このスケールでは影響は無視できます。 AdamW はウォームアップを冗長化します (アブレーションを参照)
LR スイープ (3e-4 / 1e-3 / 3e-3) — ピーク LR は実際のレバーです: 1e-3 はベースラインに対して val_loss 0.042 をカットし、その後プラトーになります (アブレーションを参照)
スケーリングスタッド

y — 5 ポイント幅 + 深さラダー (1.8M ～ 56.6M 非埋め込み) L ≈ 1.16 + 0.78·N^−0.276 (R² 0.9975) に適合。 80k 診断は、データの上限ではなく、下限がコンピューティング制限であることを示しています (スケーリングの調査を参照)
注意の解釈可能性の詳細 - 両方のスケールで前のトークンと注意をシンクするヘッド。誘導は弱くしか出現しない (解釈可能性を参照)
評価: カスタム TinyStories ルーブリック (LLM-as-judge、Llama-3.3-70B) + GPT-2 比較 — 専門化がターゲット分布のスケールを上回ります (「評価」を参照)
モデルにはまだ外部テスト ランナーがありません。テストは各モジュールの中に存在します。
__main__ はブロックされ、実行時に実行されます。モデルの場合:
Python src/model.py
これにより、configs/tiny.yaml モデルに対して 3 つのチェックが実行されます。
最新の実行 ( tiny.yaml : 4 レイヤー、非表示 128、4 ヘッド、語彙 16k):
GPT 出力形状: (2, 128, 16000)
OK: GPT は (batch、seq_len、vocab_size) を返します。
パラメータ: 4.9M
OK: グラデーションがすべてのパラメータに流れます
ここでの 4.9M は高速デバッグ構成です。 ~30M のターゲット モデルでは、
hidden_size を大きくします。このサイズでは、トークン埋め込み + LM ヘッド (語彙 ×
hidden) はパラメータの最大 84% を占めます。
私のLLM/
§── configs/ # YAML 実験構成 (configs/tiny.yaml を参照)
§── ノート/
│ §── exploration.ipynb # スクラッチ
│ └─ kaggle_train.ipynb # Kaggle T4 実行: clone → install → link data → train
§── スクリプト/
│ §── dataset_stats.py # 生の TinyStories (カウント、文字統計、サンプル) を検査します
│ §── prepare_data.py # ダウンロード + TinyStories のトークン化
│ §──generate.py # トレーニングされたチェックポイントからのサンプルテキスト
│ └── kaggle_dataset/ # トークン化された Kaggle データセットの dataset-metadata.json
§── src/
│ §── model.py # トランスフォーマー: エンベディング、アテンション、FFN、ブロック、GPT ラッパー
│ §── train.py # トレーニングループ
│ §──

data.py # データのロード/バッチ処理
│ §── tokenizer.py # トークナイザーのトレーニング
│ └── paths.py # 環境上書き可能なパス (DATA_DIR、TOKENIZER_PATH、CKPT_DIR、CONFIG_PATH)
§── 要件.txt
━── README.md
データ、チェックポイント、トークナイザー、ログは git で無視されます (再生成可能/大規模)。
python -m venv .venv && ソース .venv/bin/activate
pip install -r 要件.txt
# CUDA に一致する Torch ビルドをインストールします — https://pytorch.org/get-started/ を参照してください
Kaggle では、Torch と numpy は T4 の CUDA とともにプリインストールされています。やめてください
そこにトーチを再取り付けします。軽いdepsを追加するだけです。
pip インストールトークナイザーデータセット pyyaml tqdm matplotlib
使用法
# 0. (オプション) 生のデータセットを検査します: 話数、文字統計、サンプル
Python スクリプト/dataset_stats.py
# 1. データの準備 (ダウンロード + トークナイザーのトレーニング + コーパスのトークナイズ)
python scripts/prepare_data.py --config configs/tiny.yaml
＃1b。 (オプション) トレーニングされたトークナイザーの健全性チェック: 語彙サイズ + エンコード/デコードのラウンドトリップ
python -m src.tokenizer
# 2. トレーニング (デフォルトは configs/tiny.yaml です。CONFIG_PATH=... で別のものを選択します)
Python src/train.py
#3. 生成する
python scripts/generate.py --checkpoint Checkpoints/... --prompt " かつて "
Kaggle で実行する
トレーニングは無料の Kaggle T4 を対象としています。トークン化されたコーパスとトークナイザーは、
プライベート Kaggle データセット (.bin ファイルは git には大きすぎます)、および
ノートブック/kaggle_train.ipynb はすべてをまとめます: リポジトリのクローンを作成し、インストールします
deps、コードをデータセットに指定し、トレーニングします。
1 回限り: トークン化されたデータをプライベート データセットとして公開します (リポジトリ ルートから、
Kaggle CLI による認証が必要です
kaggle 認証ログイン ):
# メタデータの隣にペイロードをステージングします (ハードリンク -> 〜900 MB のコピーはありません)
ln -f データ/train.bin データ/val.bin トークナイザー/tokenizer.json スクリプト/kaggle_dataset/
kaggle データセット作成 -ps

cripts/kaggle_dataset # 最初のアップロード
kaggle データセット バージョン -p scripts/kaggle_dataset -m "refresh " # 後の更新
次に、Kaggle で: Notebook/kaggle_train.ipynb を開き、Accelerator → GPU T4 を設定します。
データセットを入力として追加し、オプションで WANDB_API_KEY シークレットを追加して、すべて実行します。
src/paths.py は、実行が触れるすべてのパスを環境変数から解決します。
ローカル レイアウトにフォールバックするため、同じコードが変更されずに Kaggle で実行されます。
( /kaggle/input の下の読み取り専用データ) 編集やシンボリックリンクはありません:
したがって、1 回限りのアブレーションではコードを変更する必要はありません。
CONFIG_PATH=configs/no_warmup.yaml python src/train.py
実験
実験
変数
構成
ステータス
正常アブレーション
RMSNorm と LayerNorm の比較
ablation_norm_{レイヤーノルム,rmsnorm}.yaml
終了 — 負け同点。 LayerNorm ~33% 高速化 (非融合 RMSNorm)
ウォームアップの長さ
ウォームアップステップ (50/500/2000)
ablation_warmup_{50,500,2000}.yaml
完了 — null; AdamW によるウォームアップの冗長化
ストレス下でのウォームアップ
Warmup_steps @ 10× LR、クリップなし
ablation_warmup_stress_{50,2000}.yaml
完了 — まだ null です。どちらも発散しない
LRスイープ
max_lr (3e-4/1e-3/3e-3)
ablation_lr_{3e-4,1e-3,3e-3}.yaml
完了 — 1e-3 ベスト (-0.042); LRは本物のレバーです
スケーリング
モデル サイズ (5 点ラダー) + 80k 診断
scale_{xs,s,m,base,l}.yaml 、scale_l_80k.yaml
完了 — べき乗則の当てはめ。フロアはコンピューティングが制限されています
結果
ベースライン — 30M モデル、80k ステップ
~30M モデル ( configs/small.yaml : hidden 512、8 レイヤー、8 ヘッド、語彙 8k、
seq 256 → 33.6M params ) 単一の Kaggle T4 で 80k ステップ全体にわたってトレーニングされました
ウォームアップ + コサイン スケジュール (約 42,000 トークン/秒で約 4.5 時間):
サンプル (プロンプト「むかしむかし」、温度 0.8):
昔々、リリーという名前の小さな女の子がいました。彼女はで遊ぶのが大好きでした
彼女のおもちゃのある庭。ある日、彼女は光る石を見つけたので、それを保管したいと思いました。
安全です。彼女はそれをポケットに入れて、そこへ行きました

彼女のおもちゃで遊ぶ。 …彼女は思い出した
彼女のポケットの中にある光る石は、魔法の杖になるかもしれないと考えました。
このモデルは、流暢で一貫性のある TinyStories スタイルの物語を一貫した内容で生成します。
キャラクターと明確なストーリーアーク。残存する小型モデルのアーティファクト (時折発生する)
繰り返しや軽度のロジックドリフトなど）がこの規模では予想されており、まさにそれが当てはまります。
計画されたスケーリング調査は調査することを目的としています。
ステータス: 完了 — 両方の 22k ステップ アームが Kaggle T4 でトレーニングされました。
RMSNorm は手書きです ( src/model.py 、~6 行: 各トークン ベクトルを次のように割ります)
二乗平均平方根を計算し、学習したスケールを 1 つ適用します (平均減算もなし)。
LayerNorm とは異なり、シフトします)。ノルム タイプは、norm_type を介して実行ごとに選択できます。
したがって、2 つのアームは同じコードを共有し、1 つのフィールドだけが異なります。
configs/ablation_norm_layernorm.yaml — Norm_type:layernorm
configs/ablation_norm_rmsnorm.yaml — Norm_type: rmsnorm
どちらもアブレーションを維持するために、（80,000 ベースラインに対して）削減された 22,000 ステップの予算で実行されます。
安価で、wandb_group: Norm_ablation を共有するので、それらの曲線が自動的にオーバーレイされます。
重みとバイアス。
実装健全性チェック (GPU 時間を費やす前に実行 - ランダム)
(2、4、16) 各ノルムによる入力):
どちらも振幅を同じように制御します (RMS ≈ 1)。 LayerNorm のみが平均を強制します
その 1 つの違い、LayerNorm の追加のセンタリング + シフト パラメータは次のとおりです。
アブレーションの内容全体。
パラメータコスト。 tiny.yaml のデバッグについて

[切り捨てられた]

## Original Extract

Contribute to rishipadhye/my-LLM development by creating an account on GitHub.

GitHub - rishipadhye/my-LLM · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
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
rishipadhye
/
my-LLM
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
38 Commits 38 Commits assets assets configs configs notebooks notebooks scripts scripts src src .gitignore .gitignore README.md README.md blog.md blog.md requirements.txt requirements.txt View all files Repository files navigation
A small language model built from scratch as a learning project: a ~30M
parameter decoder-only transformer trained on the
TinyStories dataset,
running on Kaggle's free T4 GPUs.
The goal is to understand transformers deeply by implementing the core pieces by
hand — model, training loop, data pipeline, and tokenizer — and then running a
set of controlled experiments, ending in a technical write-up.
📝 Read the write-up: What a 30M-Parameter Language Model Taught Me About
Training LLMs — the narrative tour of the findings (training-trick
ablations, the scaling "floor" that wasn't, tiny-model circuits, and
specialization beating scale on-distribution).
From scratch, no shortcuts. The transformer is built directly on PyTorch
tensor ops — embeddings, multi-head causal self-attention, and the training
loop are hand-written, with no nn.Transformer or HuggingFace model classes
doing the heavy lifting.
Multi-head attention done the clean way. Heads are vectorized into a
single batched matmul (reshape Q/K/V to (batch, num_heads, seq, head_dim) )
rather than looped — the same approach production implementations use.
Built bottom-up and tested as it grows. Each component is its own
nn.Module snapping together against a known (batch, seq_len, hidden_size)
contract, covered by a runnable end-to-end shape test plus a gradient-flow
test that proves the whole model is trainable (see Tests ).
Config-driven and reproducible. Hyperparameters live in versioned YAML
configs; a fixed seed keeps weights and inputs deterministic during
development.
Designed for a real compute budget. ~30M params targeting a single free
Kaggle T4 — small enough to iterate fast, large enough to produce coherent
text.
Implement a decoder-only transformer end to end, no high-level model libraries.
Train it to generate coherent TinyStories-style text on a single T4.
Run ablations to build intuition for which design choices matter:
RMSNorm vs LayerNorm
Run a small scaling experiment (model size / data / compute vs loss).
Write it all up as a technical blog post.
Implemented and smoke-tested so far:
BPE tokenizer training + encode/decode round-trip checks
Dataset inspection (story counts, char stats, samples)
YAML config loader (attribute + item access)
Token + learned-position embeddings
Multi-head causal self-attention (vectorized across heads)
Position-wise feed-forward network (GeLU, d_ff = 4 × hidden_size)
Full transformer block (pre-norm attention + FFN, residual connections)
GPT wrapper (stacked blocks + final norm + LM head → logits)
Training loop (AdamW, warmup + cosine schedule, AMP, grad clipping, checkpointing, wandb)
Kaggle T4 runner (notebook + tokenized Kaggle Dataset)
Sampling / generation ( scripts/generate.py , stops at end-of-text)
Baseline trained: 30M model, 80k steps, val_loss 1.401 at max_lr: 1e-3 (1.43 at the original 3e-4) (see Results )
Config-selectable norm type + hand-written RMSNorm; RMSNorm-vs-LayerNorm ablation trained — quality tie, LayerNorm faster (see Ablations )
Warmup-length ablation (50/500/2000) + high-LR/no-clip stress variant trained — negligible effect at this scale; AdamW makes warmup redundant (see Ablations )
LR sweep (3e-4 / 1e-3 / 3e-3) — peak LR is the real lever : 1e-3 cuts val_loss 0.042 vs baseline, then plateaus (see Ablations )
Scaling study — 5-point width+depth ladder (1.8M–56.6M non-embed) fit to L ≈ 1.16 + 0.78·N^−0.276 (R² 0.9975); 80k diagnostic shows the floor is compute-limited, not a data ceiling (see Scaling study )
Attention interpretability deep-dive — previous-token & attention-sink heads at both scales; induction only weakly emerging (see Interpretability )
Evaluation: custom TinyStories rubric (LLM-as-judge, Llama-3.3-70B) + GPT-2 comparison — specialization beats scale on the target distribution (see Evaluation )
The model has no external test runner yet — tests live in each module's
__main__ block and run on execution. For the model:
python src/model.py
This runs three checks against the configs/tiny.yaml model:
Latest run ( tiny.yaml : 4 layers, hidden 128, 4 heads, vocab 16k):
gpt out shape: (2, 128, 16000)
OK: GPT returns (batch, seq_len, vocab_size)
params: 4.9M
OK: gradients flow to all parameters
The 4.9M here is the fast debug config; the ~30M target model uses a
larger hidden_size . At this size the token embedding + LM head (vocab ×
hidden) account for ~84% of params.
my-LLM/
├── configs/ # YAML experiment configs (see configs/tiny.yaml)
├── notebooks/
│ ├── exploration.ipynb # scratch
│ └── kaggle_train.ipynb # Kaggle T4 run: clone → install → link data → train
├── scripts/
│ ├── dataset_stats.py # inspect raw TinyStories (counts, char stats, samples)
│ ├── prepare_data.py # download + tokenize TinyStories
│ ├── generate.py # sample text from a trained checkpoint
│ └── kaggle_dataset/ # dataset-metadata.json for the tokenized Kaggle Dataset
├── src/
│ ├── model.py # transformer: embeddings, attention, FFN, blocks, GPT wrapper
│ ├── train.py # training loop
│ ├── data.py # data loading / batching
│ ├── tokenizer.py # tokenizer training
│ └── paths.py # env-overridable paths (DATA_DIR, TOKENIZER_PATH, CKPT_DIR, CONFIG_PATH)
├── requirements.txt
└── README.md
Data, checkpoints, tokenizers, and logs are git-ignored (regenerable / large).
python -m venv .venv && source .venv/bin/activate
pip install -r requirements.txt
# Install a torch build matching your CUDA — see https://pytorch.org/get-started/
On Kaggle , torch and numpy are preinstalled with CUDA for the T4. Don't
reinstall torch there; just add the lighter deps:
pip install tokenizers datasets pyyaml tqdm matplotlib
Usage
# 0. (optional) Inspect the raw dataset: story count, char stats, samples
python scripts/dataset_stats.py
# 1. Prepare data (download + train tokenizer + tokenize corpus)
python scripts/prepare_data.py --config configs/tiny.yaml
# 1b. (optional) Sanity-check the trained tokenizer: vocab size + encode/decode round-trip
python -m src.tokenizer
# 2. Train (defaults to configs/tiny.yaml; pick another with CONFIG_PATH=...)
python src/train.py
# 3. Generate
python scripts/generate.py --checkpoint checkpoints/... --prompt " Once upon a time "
Run on Kaggle
Training targets a free Kaggle T4. The tokenized corpus and tokenizer ship as a
private Kaggle Dataset (the .bin files are too large for git), and
notebooks/kaggle_train.ipynb wires it all together: clone the repo, install
deps, point the code at the dataset, train.
One-time: publish the tokenized data as a private Dataset (from the repo root,
needs the Kaggle CLI authenticated via
kaggle auth login ):
# stage the payload next to its metadata (hardlinks -> no ~900 MB copy)
ln -f data/train.bin data/val.bin tokenizer/tokenizer.json scripts/kaggle_dataset/
kaggle datasets create -p scripts/kaggle_dataset # first upload
kaggle datasets version -p scripts/kaggle_dataset -m " refresh " # later updates
Then on Kaggle: open notebooks/kaggle_train.ipynb , set Accelerator → GPU T4,
add the dataset as input, optionally add a WANDB_API_KEY secret, and Run All.
src/paths.py resolves every path the run touches from environment variables,
falling back to the local layout — so the same code runs unchanged on Kaggle
(read-only data under /kaggle/input ) with no edits or symlinks:
So a one-off ablation needs no code change:
CONFIG_PATH=configs/no_warmup.yaml python src/train.py
Experiments
Experiment
Variable
Configs
Status
Norm ablation
RMSNorm vs LayerNorm
ablation_norm_{layernorm,rmsnorm}.yaml
Done — loss tie; LayerNorm ~33% faster (unfused RMSNorm)
Warmup length
warmup_steps (50/500/2000)
ablation_warmup_{50,500,2000}.yaml
Done — null; warmup redundant with AdamW
Warmup under stress
warmup_steps @ 10× LR, no clip
ablation_warmup_stress_{50,2000}.yaml
Done — still null; neither diverges
LR sweep
max_lr (3e-4/1e-3/3e-3)
ablation_lr_{3e-4,1e-3,3e-3}.yaml
Done — 1e-3 best (−0.042); LR is the real lever
Scaling
model size (5-point ladder) + 80k diagnostic
scale_{xs,s,m,base,l}.yaml , scale_l_80k.yaml
Done — power-law fit; floor is compute-limited
Results
Baseline — 30M model, 80k steps
The ~30M model ( configs/small.yaml : hidden 512, 8 layers, 8 heads, vocab 8k,
seq 256 → 33.6M params ) trained on a single Kaggle T4 for the full 80k-step
warmup + cosine schedule (~4.5 h at ~42k tokens/sec):
Sample (prompt "Once upon a time" , temperature 0.8):
Once upon a time, there was a little girl named Lily. She loved to play in the
garden with her toys. One day, she found a shiny rock and she wanted to keep it
safe. She put it in her pocket and went to play with her toys. … She remembered
the shiny rock in her pocket and thought it might make a magic wand.
The model produces fluent, coherent TinyStories-style narratives with consistent
characters and a clear story arc. Residual small-model artifacts (occasional
repetition or mild logic drift) are expected at this scale and are exactly what
the planned scaling study is meant to probe.
Status: complete — both 22k-step arms trained on a Kaggle T4.
RMSNorm is hand-written ( src/model.py , ~6 lines: divide each token vector by
its root-mean-square, then apply one learned scale — no mean-subtraction and no
shift, unlike LayerNorm). The norm type is selectable per run via norm_type in
the config, so the two arms share identical code and differ in exactly one field:
configs/ablation_norm_layernorm.yaml — norm_type: layernorm
configs/ablation_norm_rmsnorm.yaml — norm_type: rmsnorm
Both run at a reduced 22k-step budget (vs the 80k baseline) to keep ablations
cheap, and share wandb_group: norm_ablation so their curves auto-overlay in
Weights & Biases.
Implementation sanity checks (done before spending any GPU time — a random
(2, 4, 16) input through each norm):
Both control magnitude identically (RMS ≈ 1); only LayerNorm also forces the mean
to 0. That single difference — LayerNorm's extra centering + shift parameter — is
the whole substance of the ablation.
Parameter cost. On the tiny.yaml debu

[truncated]
