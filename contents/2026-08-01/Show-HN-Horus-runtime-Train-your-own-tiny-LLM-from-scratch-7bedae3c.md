---
source: "https://github.com/temple-compute/pantheon/tree/main/workflows/ai/w02-tiny-llm"
hn_url: "https://news.ycombinator.com/item?id=49132705"
title: "Show HN: Horus-runtime – Train your own tiny LLM from scratch"
article_title: "pantheon/workflows/ai/w02-tiny-llm at main · temple-compute/pantheon · GitHub"
author: "chdominguez"
captured_at: "2026-08-01T10:27:34Z"
capture_tool: "hn-digest"
hn_id: 49132705
score: 1
comments: 0
posted_at: "2026-08-01T09:27:21Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Horus-runtime – Train your own tiny LLM from scratch

- HN: [49132705](https://news.ycombinator.com/item?id=49132705)
- Source: [github.com](https://github.com/temple-compute/pantheon/tree/main/workflows/ai/w02-tiny-llm)
- Score: 1
- Comments: 0
- Posted: 2026-08-01T09:27:21Z

## Translation

タイトル: Show HN: Horus-runtime – 独自の小さな LLM を最初からトレーニングする
記事のタイトル: pantheon/workflows/ai/w02-tiny-llm at main · Temple-compute/pantheon · GitHub
説明: Horus ランタイムのワークフローのコレクション。 GitHub でアカウントを作成して、テンプル コンピューティング/パンテオンの開発に貢献します。

記事本文:
メインの pantheon/workflows/ai/w02-tiny-llm · Temple-compute/pantheon · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
テンプルコンピューティング
/
パンテオン
公共
通知
あなたはきっとsでしょう

通知設定を変更するためにサインインしました
追加のナビゲーション オプション
コード
その他のオプション ディレクトリアクション
歴史 歴史メイン ブレッドクラム
コピーパスのトップフォルダーとファイル
.. スクリプト スクリプト README.md README.md conda_env.yaml conda_env.yaml pyproject.toml pyproject.toml workflow.yaml workflow.yaml すべてのファイルを表示 README.md
概要 W-29 · ゼロからの Tiny LLM - TinyStories の事前トレーニングとサンプル
W-28 の高速で小規模なデータのコンパニオン: 同じものを事前トレーニングします
スクラッチデコーダのみのトランスフォーマーですが、オンになっています
TinyStories (短くてシンプル
Pile (約 900 GB) の代わりに、GPT で生成された童話、数百 MB)。それは存在します
LLM を最初から見て、ラップトップ上で数分で読み取れるものを学びたい人にとっては、
Pile のディスク/帯域幅要件なし。命令のチューニングと評価を完全にスキップします。
成果物は、事前トレーニングされたチェックポイントと、学習したことを証明するサンプル生成です。
ストーリープロンプトを続けます。
clone_repo (ローカル)
│
§─► prepare_tinystories_train / prepare_tinystories_val ──► tinystories_{train,val}.h5
│ └─► train_pretrain_tiny (ローカル) ──► tinystories_pretrained.pt
│ └─► サンプルストーリー (ローカル) ──► tinystories_sample.txt
クイックスタート
# horus-runtime とプラグインをインストールします (1 回のみ)
UV同期
# UV をまだ持っていない場合は、このコマンドを使用してインストールできます。
カール -LsSf https://astral.sh/uv/install.sh |しー
# それ以外の場合は、pip を使用して horus-runtime とプラグインをインストールできます。
# pip install horus-runtime horus-environments
# ワークフローを実行する
uv を実行 horus を実行 workflow.yaml
出力は workflow_results/ の下に配置されます: repo/ (複製されたトレーニング リポジトリ)、data/ (トークン化された)
TinyStories)、ckpts/tinystories_pretrained.pt (チェックポイント)、および
サンプル/tinystories_sample.txt (「むかしむかし」の欲張りな続き)。
スクラッチからの LLM の訓練 (c

孤独な
clone_repo によるランタイム）。純粋な PyTorch、transformers / trl / peft はありません。のみ
scripts/pretrain_base.py と scripts/chat.py がそこから使用されます。
タイニーストーリーズ 。経由でダウンロードされました
データセットライブラリ。
scripts/prepare_tinystories.py (クローンされたリポジトリからのものではなく、このワークフロー独自のスクリプト)、
tiktoken を使用して TinyStories を同じフラット トークン HDF5 レイアウトにトークン化します。
pretrain_base.py が期待する ( data_loader/data_loader.py の {"tokens": int32[N]} )。
クローン化されたリポジトリ自体の scripts/prepare_pretrain_data.py は Pile にハードコードされているため、これは
ワークフローはそれを直接再利用できません。
torch 、 numpy 、 h5py 、 datasets 、 tiktoken 、 conda_env.yaml 経由でプロビジョニング
(micromamba) conda_python_environment ( horus-environments ) プラグイン実行プログラムによる。
すべてのタスクは、ベア シェル エグゼキュータ ( clone_repo のみ) のいずれかでターゲット: {kind: local} を実行します。
または conda_python_environment エグゼキューター。 prepare_tinystories_{train,val} は単に与えるために repo_dir を入力として受け取ります
このワークフローのすべてのタスクは、トークナイザー スクリプトであっても 1 つの共有ルート ( clone_repo ) を持ちます。
それ自体はチェックアウトから読み取りません。
入力: なし。自己完結型。独自のソース リポジトリをクローンし、TinyStories をダウンロードします。
抱き合う顔。
ckpts/tinystories_pretrained.pt 。自己記述型チェックポイント (解決されたモデルを埋め込む)
構成、トレーニング ステップ、およびメトリクス）。
ログ/pretrain.jsonl 。記録されたトレーニング ステップごとに 1 つの JSON レコード。
サンプル/tinystories_sample.txt 。貪欲で生々しい (非チャット テンプレート) の続き
トレーニング済みモデルからの「昔々」。
データセット サイズ : prepare_tinystories_train / prepare_tinystories_val の --max_docs (デフォルト)
20,000 / 2,000) は、トークン化されるストーリーの数を制限し、それを上げます (またはフラグを完全に落とします)。
準備時間は長くなりますが、より多くのデータが必要になります。 TinyStories の完全版は、約 210 万の鉄道ストーリーです。
モデル/歩数スケール : train_pretrain_tiny l

oads configs/smoke/pretrain.json from the
クローン化されたリポジトリ ( n_embed=128, n_head=4, n_blocks=2, context_length=256 , train_steps=20 ) —
小規模なモデルは短時間トレーニングされ、ラップトップの CPU/GPU で数分で実行できるサイズに設定されています。どれでも
PretrainConfig フィールド (複製されたリポジトリの config/post_training_config.py) はオーバーライドできます
タスクのコマンドに --field 値を追加します。 --train_steps 500 --lr 3e-4 , for a
more thoroughly trained tiny model.
Sample_story の --prompt / --max_new_tokens / --temporal / --top_p はサンプルを制御します
世代。
TinyStories トークンは依然として tiktoken の r50k_base (50,257 語彙) を使用しており、スモークと一致しています
config の vocab_size=50304 (詰め込まれた) は configs/smoke/base.json から継承され、語彙はありません
mismatch with the cloned repo's model/tokenizer.
このワークフローでは SFT が実行されないため、sample_story は --raw (ベースモデルの継続) を使用します。
チェックポイントには命令に従う動作はなく、次のトークンの継続のみがあります。
train-llm-from-scratch — source
repo cloned by this workflow.
TinyStories データセット ·
紙 — GPT で生成された短くてシンプルな子供向けの物語。
horus-runtime — ワークフロー エンジンとプラグイン フレームワーク。
horus-environments — conda ベースの Python 環境用の horus プラグイン。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A collection of workflows for the Horus Runtime. Contribute to temple-compute/pantheon development by creating an account on GitHub.

pantheon/workflows/ai/w02-tiny-llm at main · temple-compute/pantheon · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
temple-compute
/
pantheon
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
More options Directory actions
History History main Breadcrumbs
Copy path Top Folders and files
.. scripts scripts README.md README.md conda_env.yaml conda_env.yaml pyproject.toml pyproject.toml workflow.yaml workflow.yaml View all files README.md
Outline W-29 · Tiny LLM From Scratch - TinyStories Pretrain & Sample
A fast, small-data companion to W-28 : pretrains the same
from-scratch decoder-only Transformer, but on
TinyStories (short, simple
GPT-generated children's stories, a couple hundred MB) instead of the Pile (~900GB). It exists
for people who want to see a from-scratch LLM learn something legible in minutes on a laptop,
without the Pile's disk/bandwidth requirements. It skips instruction tuning and eval entirely.
The deliverable is a pretrained checkpoint plus a sample generation proving it learned to
continue a story prompt.
clone_repo (local)
│
├─► prepare_tinystories_train / prepare_tinystories_val ──► tinystories_{train,val}.h5
│ └─► train_pretrain_tiny (local) ──► tinystories_pretrained.pt
│ └─► sample_story (local) ──► tinystories_sample.txt
Quick start
# Install the horus-runtime and plugins (one time)
uv sync
# You can install UV with this command if you don't have it yet:
curl -LsSf https://astral.sh/uv/install.sh | sh
# Otherwise, you can install the horus-runtime and plugins with pip:
# pip install horus-runtime horus-environments
# Run the workflow
uv run horus run workflow.yaml
Outputs land under workflow_results/ : repo/ (the cloned training repo), data/ (tokenized
TinyStories), ckpts/tinystories_pretrained.pt (checkpoint), and
samples/tinystories_sample.txt (a greedy continuation of "Once upon a time,").
train-llm-from-scratch (cloned at
runtime by clone_repo ). Pure PyTorch, no transformers / trl / peft . Only
scripts/pretrain_base.py and scripts/chat.py are used from it.
TinyStories . Downloaded via the
datasets library.
scripts/prepare_tinystories.py (this workflow's own script, not from the cloned repo),
tokenizes TinyStories with tiktoken into the same flat-token HDF5 layout
( data_loader/data_loader.py 's {"tokens": int32[N]} ) that pretrain_base.py expects.
The cloned repo's own scripts/prepare_pretrain_data.py is hardcoded to the Pile, so this
workflow can't reuse it directly.
torch , numpy , h5py , datasets , tiktoken , provisioned via conda_env.yaml
(micromamba) by the conda_python_environment ( horus-environments ) plugin executor.
Every task runs target: {kind: local} , either on the bare shell executor (just clone_repo )
or the conda_python_environment executor. prepare_tinystories_{train,val} take repo_dir as an input purely to give
every task in this workflow one shared root ( clone_repo ), even though the tokenizer script
itself doesn't read from the checkout.
Input : none. Self-contained; clones its own source repo and downloads TinyStories from
Hugging Face.
ckpts/tinystories_pretrained.pt . Self-describing checkpoint (embeds the resolved model
config, training step, and metrics).
logs/pretrain.jsonl . One JSON record per logged training step.
samples/tinystories_sample.txt . A greedy, raw (non-chat-template) continuation of
"Once upon a time," from the trained model.
Dataset size : prepare_tinystories_train / prepare_tinystories_val 's --max_docs (default
20,000 / 2,000) caps how many stories get tokenized, raise it (or drop the flag entirely) for
more data at the cost of longer prep time. Full TinyStories is ~2.1M train stories.
Model/step-count scale : train_pretrain_tiny loads configs/smoke/pretrain.json from the
cloned repo ( n_embed=128, n_head=4, n_blocks=2, context_length=256 , train_steps=20 ) — a
small model trained briefly, sized to run on a laptop CPU/GPU in minutes. Any
PretrainConfig field ( config/post_training_config.py in the cloned repo) can be overridden
by adding --field value to the task's command: , e.g. --train_steps 500 --lr 3e-4 , for a
more thoroughly trained tiny model.
sample_story 's --prompt / --max_new_tokens / --temperature / --top_p control the sample
generation.
TinyStories tokens still use tiktoken 's r50k_base (50,257 vocab), matching the smoke
config's vocab_size=50304 (padded up) inherited from configs/smoke/base.json , no vocab
mismatch with the cloned repo's model/tokenizer.
sample_story uses --raw (base-model continuation) since this workflow never runs SFT: the
checkpoint has no instruction-following behavior, only next-token continuation.
train-llm-from-scratch — source
repo cloned by this workflow.
TinyStories dataset ·
paper — short, simple GPT-generated children's stories.
horus-runtime — workflow engine and plugin framework.
horus-environments — horus plugin for conda-based Python environments.
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
