---
source: "https://github.com/gradian-ai/gradian"
hn_url: "https://news.ycombinator.com/item?id=49168696"
title: "I built an LLM debugger for fine-tuning failures"
article_title: "GitHub - gradian-ai/gradian: Find the data that broke your fine-tune. · GitHub"
author: "abdullah-xyz"
captured_at: "2026-08-04T13:50:46Z"
capture_tool: "hn-digest"
hn_id: 49168696
score: 3
comments: 1
posted_at: "2026-08-04T13:27:56Z"
tags:
  - hacker-news
  - translated
---

# I built an LLM debugger for fine-tuning failures

- HN: [49168696](https://news.ycombinator.com/item?id=49168696)
- Source: [github.com](https://github.com/gradian-ai/gradian)
- Score: 3
- Comments: 1
- Posted: 2026-08-04T13:27:56Z

## Translation

タイトル: 失敗を微調整するための LLM デバッガを構築しました
記事のタイトル: GitHub - gradian-ai/gradian: 微調整を壊したデータを見つけます。 · GitHub
説明: 微調整を壊したデータを見つけます。 GitHub でアカウントを作成して、gradian-ai/gradian の開発に貢献してください。

記事本文:
GitHub - gradian-ai/gradian: 微調整を壊したデータを見つけます。 · GitHub
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
グラディアンアイ
/
グラディアン
公共
通知
署名が必要です

で通知設定を変更します
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .github/ workflows .github/ workflows 制約 制約 docs docs Experiments/ gpu Experiments/ gpu src/ gradian src/ gradian testing testing .gitignore .gitignore LICENSE LICENSE README.md README.md install.sh install.sh pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイル ナビゲーション
どのトレーニング データ、またはどの構成設定が微調整を妨げたのでしょうか?
ドキュメント ·
クイックスタート ·
仕組み ·
数学 ·
検証・
研究
Gradian は、LoRA の微調整された LLM 用のトレーニング データ アトリビューション エンジンです。微調整してください
一部の機能が低下し、どのトレーニング サンプルが低下の原因となったかがわかります。
読み取り可能なクラスターに集約され、データセットの決定論的な監査と照合してクロスチェックされます。
ハイパーパラメータと損失曲線。
曲率補正された影響関数を使用して、勾配を通じてモデルの動作を特定します。
フォワードパスを反転するとは主張していません。それは数学的に不可能であり、不必要です。
最初に微調整を行ってから、グラディアンが説明します。 Gradian は、すでに存在する微調整のためのデバッガです。
トレーナーではなく存在します。ベースモデル、トレーニングした LoRA アダプター、トレーニング データセット、
eval セット、およびオプションで設定ログとトレーナー ログを保持する run ディレクトリ。前に向かって走ります
グラデーションをキャプチャするために逆方向パスを実行し、アダプター上でオプティマイザー ステップを実行することはありません。
gradian train は、trl、unsloth、
または単純なトランスフォーマー、そしてそれを必要とするパイプラインの唯一の部分は反事実です
再トレーニング。告発された例を削除することで実際に能力が回復するかどうかを確認します。

やあ。
完全なドキュメントは gradian.dev/docs にあります。
docs/ のマークダウン。
これを個人としてではなくエージェントとして読んでいますか？すべてのページには、同じ URL に生のマークダウン ツインがあります。
.md に加えて、厳選されたインデックスは gradian.dev/llms.txt にあり、
1 つのファイル内のコーパス全体は gradian.dev/llms-full.txt にあります。
構造にとらわれないトレーナー
2 つのまったく異なるものによって微調整が中断され、そのうちの 1 つだけを実行するツールは半分のツールです。
トレーニング例の一部のクラスターは、モデルに、あなたの能力に対抗する何かを教えました。
気にする。それを見つけるには、影響関数が必要です: アダプター上の例ごとの勾配、
逆曲率の近似によって事前条件付けされ、構築されたクエリ勾配に対して点が付けられます
低下した特定の項目から。それがエンジンです。
省略された補完。プロンプトをカバーする損失マスク。 EOSがないのでモデルが止まることはありません。
トレーニングと評価の汚染。完全な微調整用に設定された学習率。 400 を超える 12 のエポック
例。補完のみのマスキングの隣にPacking=True。これらはどれもクラッシュせず、何も問題として表示されません
トレーニング エラーであり、影響力を計算するだけでも、単に をチェックすること以上に説明がつきません。それ
診断レイヤーです。データセット、構成、トレーナー ログに対する 41 の決定的チェックがそれぞれ行われます。
証拠と具体的な修正が必要です。
このレポートはこの 2 つを結合したものです。あなたを傷つけたクラスターが切り捨てられたクラスターであることが判明したとき、グラディアン
それはそれに対して何をすべきかが変わるからです。
# 最も便利な方法: データセットと実行を監査します。 GPU なし、帰属なし、数秒。
グラディアン診断 \
--dataset train.jsonl \
--run 実行/my-finetune \
--base-model metal-llama/Llama-3.2-1B
# 要点: この機能の低下を引き起こしたデータは何ですか?
# --adapter は、すでにトレーニングした LoRA アダプターです。グラディアンはそれを読みますが、生成しません

それをしてください。
グラデーション属性 \
--base-model metal-llama/Llama-3.2-1B-Instruct \
--adapter が実行/my-finetune/adapter \
--dataset train.jsonl \
--eval-datasetcapability_eval.jsonl \
--capability json_formatting \
--アウトレポート/
1 つのレポートでの出力:
json_formatting は、200 項目にわたって 0.910 ～ 0.640 (-0.270、95% CI [-0.390、-0.150]) に低下しました。
54 件の項目が正誤で反転しました。クラスター 7 (請求書、合計、通貨) が 62% を占めます。
マイナスの影響。 40 個の例のうち 34 個で補完が切り詰められており、
内容の問題ではなく、機械的な原因。 max_seq_length=512 -> 1024 を修正して再トレーニングします。
Gradian はこのリポジトリからインストールされます。 package-index リリースがないため、インストーラーがビルドします。
ブランチ、タグ、または sha からのプロジェクト。
カール -fsSL https://raw.githubusercontent.com/gradian-ai/gradian/main/install.sh |しー
これにより、 ~/.gradian に隔離された環境が構築され、適切な PyTorch ホイールが選択されます。
ハードウェアを作成し、以下の依存関係ウィンドウを固定し、gradian を ~/.local/bin にリンクして、次のように終了します。
グラディアンの医師を実行しています。シェル rc 内の 1 つの PATH 行以外には何も触れません。
sh はオプションを独自の引数として要求するため、オプションをパイプ経由でフラグにすることはできません。使用する
代わりに環境変数を使用します。
URL=https://raw.githubusercontent.com/gradian-ai/gradian/main/install.sh
カール -fsSL $URL | GRADIAN_CPU=1 sh # CPU ホイールを強制します
カール -fsSL $URL | GRADIAN_UNSLOTH=1 sh # アンナマケモノトレーナーを追加します
カール -fsSL $URL | GRADIAN_REF=v0.1.0 sh # メインではなくタグをインストールします
カール -fsSL $URL | GRADIAN_UNINSTALL=1 sh # 再度削除します
ソースから
git clone https://github.com/gradian-ai/gradian && cd gradian
sh install.sh --local # 編集可能なインストール、上記と同じ環境処理
または、環境を自分で管理します。
pip install -e 。 # コア: トーチ、トランスフォーマー、ペフト
ピップイン

stop -e " .[train] " # 微調整を開始します (trl、データセット、加速)
pip install -e " .[cluster] " # 文変換と HDBSCAN (それ以外の場合は TF-IDF フォールバック)
pip install -e " .[unsloth] " # unsloth でトレーニングする (GPU のみ: CUDA、triton、xformers)
gradian Doctor # 依存関係マトリックスとデバイスを確認する
パッケージ
対応ウィンドウ
たいまつ
2.4～2.11
変圧器
4.51～5.5
ペフト
0.18以上
trl
0.24以下
それが Gradian が開発され、内部でテストされているウィンドウです。最新のアップストリーム リリースは上にあります
その一部は意図的です。エコシステム内のすべてが同意する、サポートされているウィンドウは、
現行のものよりも価値がある。正確なピンはconstraints/に存在し、グラディアン医師は警告する
インストールされているバージョンが範囲外になった場合。
1 つのレポートに 3 つのレーンが入ります。評価は何が変更されたかを測定し、属性を拒否します
ノイズ帯域内のあらゆるもの。属性は、生き残った回帰をランク付けします。
トレーニング データに対するクラスター化された判定。監査はデータセット、構成、トレーナーのログをチェックします
機械的に、GPU や帰属は関係ありません。
1 つの逆方向パスでの例ごとの勾配。アトリビューションには例ごとに勾配が必要ですが、
標準逆方向ではバッチ平均が得られます。 Gradian は、サンプルの単純な合計としてバッチ損失を構築します。
そして、フックを使用して各 LoRA Linear レイヤーの入力および出力勾配をキャプチャするため、1 つのパスで
例ごとの正確な勾配。これは正確であり、vmap の脆弱性を回避し、因数分解された結果を生成します。
EK-FAC フォームは後で無料で必要になります。バッチサイズ 1 のリファレンス実装と vmap パスの両方
が存在し、テスト スイートでは 3 つすべてが同意する必要があります。
1 つのインターフェイスの背後にあるエンジン。 DataInf は v1 エンジンです。グラドットとトレーシンは、
曲率のないベースラインとクロスチェック。インターフェイスが統一されているため、2 つのエンジンと

順位の一致を報告することは設定フラグであり、不一致は信頼度が低いとして報告されます。
黙って平均化するのではなく。
例ごとの勾配、フックされたアクティベーション X と出力勾配 G から取得:
g_j = einsum("bto,bti->boi", G, X)
H が最適なヘッセ行列である場合の、クエリ q に対するトレーニング例 j の古典的な影響:
スコア_j = ⟨ ∇L(z_q), H⁻¹ ∇L(z_j) ⟩
DataInf は、反転と加算の順序を入れ替えてから、シャーマン-モリソン法を適用するため、行列は存在しません。
今まで逆転した。 LoRA ブロック l ごとに、減衰 λ_l と n のトレーニング例を使用します。
r_l = (1/λ_l)・[ g_q − (1/n)・Σᵢ (gᵢ・g_q)/(λ_l + ‖gᵢ‖²)・gᵢ ]
スコア_j = Σ_l ⟨ r_l, g_j⁽ˡ⁾ ⟩
対照的なクエリ グラデーション。世代が利用可能な場合は常に使用されます。
g_q = ∇ℓ(正解) − ∇ℓ(モデルが実際に言ったこと)
オプションのジョンソン・リンデンシュトラウス投影。ストレージをモデルのサイズに依存せずに作成します。
忠実度ランキングにおける測定コスト:
g̃ = (1/√k)・R・g R ∈ ℝ^{k×d}, R_ij ~ N(0, 1)
再トレーニングによってテストできる一次予測:
予測損失デルタの場合は削除 = スコア_j / n
署名規則、一度修正されました。プラスのスコアは、例がクエリの内容を強化したことを意味します
行動。マイナスのスコアは、劣化したことを意味します。
クエリのグラデーションはデフォルトで対照的です
これは最も目立たないが最も重要な設計上の決定であり、測定結果から得られたものです。
紙ではなく。モデルが間違ったことを言っている置換失敗の場合、
正解に対するクロスエントロピーでは、符号が逆になります。毒を盛られた例は次のように出てきます。
微調整で主に教えられるのは回答形式であるため、コーパス内で最も役立つデータです。
正解時の損失も本当に減少します。モデルの勾配を減算する
実際にはその共有方向をキャンセルすると言いました

n.
正規化、モジュール カバレージ、エンジン、ダンピングに対して不変であり、16 の構成にわたって
同じ間違った結論。完全な測定は docs/GPU_TESTING.md にあります。
全体の例の影響は、「どの例が私を傷つけたか」に答えます。長い文書の場合、それは多くの場合、あまりにも重要です
粗い。 400 トークンのエントリ内の間違った文は、およそ 110 個の教師ありトークンのうち 2 個に相当します。
有害な信号は、同じことを述べた短いエントリと比較して約 760 倍に薄められます。
スパンレベルのドリルダウン ( gradian.query.spans ) はトークンを固定したままにし、損失マスクのみを変更します。
文ごとに 1 つの勾配をスコアリングし、92% の確率で犯人の文を特定します。
クエリは、平均 (それ以外の場合は 33%) ではなく、単一の失敗項目から構築されます。
基本モデルが原因である可能性が高い場合に表示されます
基本モデルのトレーニング データを見ることは決してないので、その内部は不透明なままです。グラディアンが基地を運営している
基本モデルも失敗する項目 (機能がなかった)
負けたため、トレーニング例では説明できません）、基本スコアが低く（失うものがほとんどありませんでした）、そして薄い
基本優先マージン。基本モデルが何に対してどれだけ強く正解を優先するかを意味します。
チューニングされたモデルは今こう言います。約 1 nat 未満では、その能力は非常に弱く保持され、ほとんどすべての
いいよ-

[切り捨てられた]

## Original Extract

Find the data that broke your fine-tune. Contribute to gradian-ai/gradian development by creating an account on GitHub.

GitHub - gradian-ai/gradian: Find the data that broke your fine-tune. · GitHub
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
gradian-ai
/
gradian
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .github/ workflows .github/ workflows constraints constraints docs docs experiments/ gpu experiments/ gpu src/ gradian src/ gradian tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md install.sh install.sh pyproject.toml pyproject.toml View all files Repository files navigation
Which of my training data, or which config setting, broke my fine-tune?
Documentation ·
Quick start ·
How it works ·
The math ·
Verification ·
Research
Gradian is a training-data attribution engine for LoRA fine-tuned LLMs. Give it a fine-tune that
regressed some capability, and it tells you which training examples caused the regression,
aggregated into readable clusters, cross-checked against a deterministic audit of your dataset,
hyperparameters, and loss curve.
It attributes model behavior through gradients , using curvature-corrected influence functions.
It does not claim to invert the forward pass. That is mathematically impossible, and unnecessary.
You fine-tune first, then Gradian explains. Gradian is a debugger for a fine-tune that already
exists, not a trainer. Point it at a base model, the LoRA adapter you trained, the training dataset,
an eval set, and optionally the run directory holding your config and trainer logs. It runs forward
and backward passes to capture gradients, and never takes an optimizer step over your adapter.
gradian train exists as an optional convenience for launching a LoRA fine-tune with trl, unsloth,
or plain transformers , and the only part of the pipeline that needs it is counterfactual
retraining, which checks that dropping the accused examples actually recovers the capability.
Full documentation is at gradian.dev/docs , generated from the
markdown in docs/ .
Reading this as an agent rather than a person? Every page has a raw markdown twin at the same URL
plus .md , the curated index is at gradian.dev/llms.txt , and the
whole corpus in one file is at gradian.dev/llms-full.txt .
Trainer agnostic by construction
Two very different things break fine-tunes, and a tool that only does one of them is half a tool.
Some cluster of your training examples taught the model something that fights the capability you
care about. Finding it needs influence functions: per-example gradients over the adapter,
preconditioned by an approximation of the inverse curvature, dotted against a query gradient built
from the specific items that regressed . That is the engine.
Truncated completions. A loss mask that covers the prompt. No EOS, so the model never stops.
Train and eval contamination. A learning rate set for full fine-tuning. Twelve epochs over 400
examples. packing=True next to completion-only masking. None of these crash, none show up as a
training error, and no amount of influence math explains them better than simply checking . That
is the diagnostics layer: 41 deterministic checks over your dataset, config, and trainer logs, each
with evidence and a concrete fix.
The report joins the two. When the cluster that hurt you turns out to be the truncated one, Gradian
says so, because that changes what you should do about it.
# Fastest useful thing: audit a dataset and a run. No GPU, no attribution, seconds.
gradian diagnose \
--dataset train.jsonl \
--run runs/my-finetune \
--base-model meta-llama/Llama-3.2-1B
# The full thing: what data caused this capability to regress?
# --adapter is the LoRA adapter you already trained. Gradian reads it, it does not produce it.
gradian attribute \
--base-model meta-llama/Llama-3.2-1B-Instruct \
--adapter runs/my-finetune/adapter \
--dataset train.jsonl \
--eval-dataset capability_eval.jsonl \
--capability json_formatting \
--out reports/
Output, in one report:
json_formatting regressed 0.910 to 0.640 (-0.270, 95% CI [-0.390, -0.150]) over 200 items.
54 items flipped from right to wrong. Cluster 7 (invoice, total, currency) accounts for 62% of
the negative influence. 34 of its 40 examples have truncated completions, which points to a
mechanical cause rather than a content problem. Fix max_seq_length=512 -> 1024 and retrain.
Gradian installs from this repository. There is no package-index release, so the installer builds
the project from a branch, tag or sha.
curl -fsSL https://raw.githubusercontent.com/gradian-ai/gradian/main/install.sh | sh
That builds an isolated environment in ~/.gradian , picks the right PyTorch wheel for your
hardware, pins the dependency window below, links gradian into ~/.local/bin , and finishes by
running gradian doctor . It touches nothing else except one PATH line in your shell rc.
Options cannot be flags through a pipe, because sh claims them as its own arguments. Use
environment variables instead:
URL=https://raw.githubusercontent.com/gradian-ai/gradian/main/install.sh
curl -fsSL $URL | GRADIAN_CPU=1 sh # force the CPU wheel
curl -fsSL $URL | GRADIAN_UNSLOTH=1 sh # add the unsloth trainer
curl -fsSL $URL | GRADIAN_REF=v0.1.0 sh # install a tag rather than main
curl -fsSL $URL | GRADIAN_UNINSTALL=1 sh # remove it again
From source
git clone https://github.com/gradian-ai/gradian && cd gradian
sh install.sh --local # editable install, same environment handling as above
Or manage the environment yourself:
pip install -e . # core: torch, transformers, peft
pip install -e " .[train] " # launch fine-tunes (trl, datasets, accelerate)
pip install -e " .[cluster] " # sentence-transformers and HDBSCAN (TF-IDF fallback otherwise)
pip install -e " .[unsloth] " # train with unsloth (GPU only: CUDA, triton, xformers)
gradian doctor # verify the dependency matrix and devices
Package
Supported window
torch
2.4 to 2.11
transformers
4.51 to 5.5
peft
0.18 and above
trl
0.24 and below
That is the window Gradian is developed and tested inside. The latest upstream releases sit above
parts of it, which is deliberate: a supported window that everything in the ecosystem agrees on is
worth more than being current. The exact pins live in constraints/ , and gradian doctor warns
when your installed versions drift out of range.
Three lanes run into one report. Evaluate measures what changed and refuses to attribute
anything inside the noise band. Attribute turns the surviving regression into a ranked,
clustered verdict over your training data. Audit checks the dataset, config, and trainer logs
mechanically, with no GPU and no attribution involved.
Per-example gradients in one backward pass. Attribution needs a gradient per example, but a
normal backward gives the batch average. Gradian builds the batch loss as a plain sum over examples
and captures each LoRA Linear layer's input and output-gradient with hooks, so one pass yields the
exact per-example gradient. It is exact, it avoids vmap fragility, and it produces the factored
form EK-FAC needs later for free. A batch-size-1 reference implementation and a vmap path both
exist, and the test suite requires all three to agree.
Engines behind one interface. DataInf is the v1 engine. graddot and tracin are the
curvature-free baseline and cross-check. Because the interface is uniform, running two engines and
reporting their rank agreement is a config flag, and disagreement is reported as low confidence
rather than silently averaged.
Per-example gradient, captured from hooked activations X and output-gradients G :
g_j = einsum("bto,bti->boi", G, X)
Classical influence of training example j on query q , with H the Hessian at the optimum:
score_j = ⟨ ∇L(z_q), H⁻¹ ∇L(z_j) ⟩
DataInf swaps the order of inversion and summation, then applies Sherman-Morrison, so no matrix is
ever inverted. Per LoRA block l , with damping λ_l and n training examples:
r_l = (1/λ_l)·[ g_q − (1/n)·Σᵢ (gᵢ·g_q)/(λ_l + ‖gᵢ‖²)·gᵢ ]
score_j = Σ_l ⟨ r_l, g_j⁽ˡ⁾ ⟩
The contrastive query gradient, used whenever generations are available:
g_q = ∇ℓ(correct answer) − ∇ℓ(what the model actually said)
Optional Johnson-Lindenstrauss projection, which makes storage independent of model size at a
measured cost in ranking fidelity:
g̃ = (1/√k)·R·g R ∈ ℝ^{k×d}, R_ij ~ N(0, 1)
First-order prediction you can test by retraining:
predicted_loss_delta_if_removed = score_j / n
Sign convention, fixed once. A positive score means the example reinforced the queried
behavior. A negative score means it degraded it.
The query gradient is contrastive by default
This is the least obvious and most important design decision, and it came out of a measurement
rather than a paper. For a substitution failure, where the model now says the wrong thing,
cross-entropy on the correct answer gets the sign backwards. The poisoned examples come out as the
most helpful data in the corpus, because what a fine-tune mostly teaches is answer format , which
genuinely lowers the loss on the right answer too. Subtracting the gradient of what the model
actually said cancels that shared direction.
Invariant to normalization, module coverage, engine, and damping, across 16 configurations with the
same wrong conclusion. Full measurement in docs/GPU_TESTING.md .
Whole-example influence answers "which examples hurt me". For long documents that is often too
coarse. A wrong sentence inside a 400-token entry is roughly 2 of 110 supervised tokens, so its
harmful signal is diluted about 760 times relative to a short entry stating the same thing.
Span-level drill-down ( gradian.query.spans ) keeps the tokens fixed and varies only the loss mask,
scoring one gradient per sentence, and localizes the culprit sentence 92% of the time , provided
the query is built from the single failing item rather than an average (33% otherwise).
It says when the base model is the likelier culprit
We never see the base model's training data, so its internals stay opaque. Gradian does run the base
model, and that bounds the blame three ways: items the base model also fails (no capability was
lost, so no training example explains them), a low base score (little was there to lose), and a thin
base preference margin , meaning how strongly the base model favored the correct answer over what
the tuned model now says. Below about 1 nat, the capability was held so weakly that almost any
fine-

[truncated]
