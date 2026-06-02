---
source: "https://github.com/cnygaard/glq"
hn_url: "https://news.ycombinator.com/item?id=48362773"
title: "Show HN: Glq LLM quantization using E8 lattice"
article_title: "GitHub - cnygaard/glq: E8 lattice codebook quantization for LLM weights — 2/3/4 bpw with fused Triton inference kernel · GitHub"
author: "acd"
captured_at: "2026-06-02T00:31:21Z"
capture_tool: "hn-digest"
hn_id: 48362773
score: 2
comments: 0
posted_at: "2026-06-01T21:17:26Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Glq LLM quantization using E8 lattice

- HN: [48362773](https://news.ycombinator.com/item?id=48362773)
- Source: [github.com](https://github.com/cnygaard/glq)
- Score: 2
- Comments: 0
- Posted: 2026-06-01T21:17:26Z

## Translation

タイトル: HN を表示: E8 格子を使用した Glq LLM 量子化
記事のタイトル: GitHub - cnygaard/glq: LLM 重みの E8 格子コードブック量子化 — 融合 Triton 推論カーネルによる 2/3/4 bpw · GitHub
説明: LLM 重みの E8 格子コードブック量子化 — 融合 Triton 推論カーネルによる 2/3/4 bpw - cnygaard/glq
HN テキスト: AI の助けを借りて、glq と呼ばれる E8 LLM コード ブック量子化ライブラリのオープン ソース メソッドを作成しました。 Glq の作成に興味がありました
PC ゲーマーおよび Devops として、LLM と AI の両方に興味があります。現在の RAM の価格の高騰と LLM リソースの使用量も、私が glq を書くきっかけになりました。代替の LLM 圧縮方法を使用して、限られた VRAM サイズでゲーム GPU をさらに圧縮できるかという疑問が生じます。 Glq は、他の LLM 量子化アルゴリズムと比較して、重みあたり 2 ビットから重みあたり 4 ビットまでの範囲で効果的です。ワードあたりのビット数が低い場合の glq の有効性は、線形手法と比較した E8 格子の特性によるものです。
Glq は、LLM 層が量子化に対してどの程度敏感であるかに応じて、異なる LLM 層が異なる圧縮ビット重みを使用する混合精度量子化もサポートします。混合精度は、MP3 または MP4 の可変ビット レート エンコーディングに似たものだと考えてください。現在、コストをよりリーズナブルに保つために、g7e AWS スポット インスタンスを使用して glq を開発しています。 Glq は vllm を使用します。E8 による 4 ビット キー値キャッシュは、NexusQuant からインスピレーションを受けました。通常、BF16 で VRAM に収まるキー値キャッシュの約 4 倍、つまり INT8 と比較して約 2 倍のキー値キャッシュを詰め込むようにしています。私はどういうわけか最初に、GPU L1 キャッシュに適した 4096 コード ブック エントリではなく、65536 エントリの E8 コード ブック サイズを選択してしまいました。 65535 個のコード ブック エントリがあると、LLM 圧縮率が高くなりますが、デコード速度が犠牲になることがわかりました。私はしようとしています

Nvidia Cuda グラフを使用して補正し、デコードを最適化します。現在作業中です。 Nvidia GPU を備えた Linux 上の Python 仮想環境に glq をインストールするには:
pip install glq Python PIP パッケージ
https://pypi.org/project/glq/ Glq ソース コード。
https://github.com/cnygaard/glq ライブラリのインスピレーションとなった現在の PC RAM の価格。
https://pcpartpicker.com/trends/price/memory/ https://en.wikipedia.org/wiki/E8_lattice
球のパッキング問題に対する最適な解決策を提供する 8 次元格子。これは、大砲の弾やリンゴを最適な方法で積み上げるのと同じように考えてください。リンゴを LLM 重みと交換するだけです。 E8 格子の写真
https://en.wikipedia.org/wiki/E8_polytope#/media/File:E8_gra... クレジット: GLQ は E8 Quip# からインスピレーションを受け、Key value E8 圧縮は NexusQuant からインスピレーションを受けました。数学: 次元 8 の球パッキング問題、マリーナ ヴィアゾフスカ
https://arxiv.org/abs/1603.04246 4bpw glq Gemma 4 E4b の量子化 - 命令の調整
https://huggingface.co/xv0y5ncu/Gemma-4-E4B-it-GLQ-4bpw SmolLM3 の 3.5bpw 混合精度量子化
https://huggingface.co/xv0y5ncu/SmolLM3-3B-GLQ-3.5bpw Nvidia コンテナー ツールキットを使用した Nvidia GPU 上の glq の Docker イメージ。
docker run --rm --gpus all \
-v "$HOME/.cache/huggingface:/cache/hf" \
ghcr.io/cnygaard/glq-env:0.5.0 \
Python -c '
import glq.hf_integration, torch # GLQ を HF に登録します
トランスフォーマーから AutoModelForCausalLM、AutoTokenizer をインポート
mid = "xv0y5ncu/SmolLM3-3B-GLQ-3.5bpw"
tok = AutoTokenizer.from_pretrained(mid)
モデル = AutoModelForCausalLM.from_pretrained(
Mid、device_map="cuda"、torch_dtype=torch.float16)
ids = tok("フランスの首都は", return_tensors="pt").to("cuda")
print(tok.decode(model.generate(*ids, max_new_tokens=20)[0],
Skip_special_tokens=True))
' 現在 glq の作業が進行中です

デコードの速度が向上し、より多くの LLM モデル アーキテクチャがサポートされるようになりました。未解決の質問: glq は Nvidia DGX Spark および 4070-5090 などのゲーム用 Nvidia ハードウェアで動作しますか?

記事本文:
GitHub - cnygaard/glq: LLM 重みの E8 格子コードブック量子化 — 融合 Triton 推論カーネルによる 2/3/4 bpw · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
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
クニーガード
/
グルク
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
278 コミット 278 コミット .github .github ベンチマーク ベンチマークの例 例 glq glq glq_vllm glq_vllm インフラ インフラ スクリプト スクリプト テスト テスト .dockerignore .dockerignore .gitignore .gitignore Dockerfile Dockerfile ライセンス ライセンス README.md README.md Compare_methods.py Compare_methods.py eval_ppl.py eval_ppl.py pyproject.toml pyproject.toml test_glq_e2e.py test_glq_e2e.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
E8 格子コードブックを使用した LLM のトレーニング後の重み量子化。
GLQ は、各 8 重みグループを 16 ビット インデックスとして 65,536 エントリにエンコードします。
E8格子コードブック。ランダム化アダマール変換 (RHT) の無相関化
ヘッセ行列。ユークリッド最近傍検索が最適に近いものになります。
代理損失の下で。結果: 品質を備えた 2 ～ 8 bpw の重量
QuIP# に匹敵し、GPTQ よりも優れており、融合された CUDA カーネル
実体化せずに、圧縮されたインデックスに対して直接 matmuls を実行します。
重み行列。
pip install glq # には PyTorch ≥ 2.0 が必要です
Python ≥ 3.10。 Triton は CUDA 上の PyTorch に同梱されており、使用されます
自動的に。 CUDA C 拡張機能は最初の実行時に JIT ビルドされます
(~30 秒); CPU は dequantize-then-matmul に戻ります。
glq をインポートします。 hf_integration # GLQ をトランスフォーマーに登録します
トランスフォーマーからのインポート AutoModelForCausalLM 、 AutoTokenizer
モデル = AutoModelForCausalLM 。 from_pretrained (
"xv0y5ncu/SmolLM2-360M-Instruct-GLQ-4bpw" ,
device_map = "自動" 、
)
tok = AutoTokenizer 。 from_pretrained ( "xv0y5ncu/SmolLM2-360M-Instruct-GLQ-4bpw" )
print ( tok . decode ( model .generate (
** tok ( "フランスの首都は" 、 return_tensors = "pt" )。に（モデルとデバイス）、
max_new_tokens = 20 、
)[ 0 ]、skip_special_tokens = True ))
glq.hf_i をインポートする

統合は quant_method="glq" を HF に登録します
トランスフォーマー。 from_pretrained は、nn.Linear を E8RHTLinear に交換します。
推論には融合された CUDA C カーネルを使用します。 CPU は次の状態に戻ります
単純な dequantize-then-matmul。
利用可能な事前量子化チェックポイント
リポ
ベースモデル
血圧
ライセンス
xv0y5ncu/SmolLM2-135M-命令-GLQ-4bpw
SmolLM2-135M-命令
4.0
アパッチ2.0
xv0y5ncu/SmolLM2-360M-命令-GLQ-4bpw
SmolLM2-360M-命令
4.0
アパッチ2.0
xv0y5ncu/SmolLM3-3B-GLQ-3.5bpw
SmolLM3-3B
3.5（混合）
アパッチ2.0
xv0y5ncu/Gemma-4-E4B-it-GLQ-4bpw
ジェマ-4-E4B-it
4.0
アパッチ2.0
xv0y5ncu/Devstral-Small-2-24B-Instruct-GLQ-4bpw
デヴストラル-小 24B
4.0
アパッチ2.0
xv0y5ncu/Nemotron-3-Nano-30B-A3B-GLQ-4bpw
Nemotron-3-Nano-30B (マンバ-MoE)
4.0
ネモトロン
独自のモデルを量子化する
pip install ' glq[quantize] ' # トランスフォーマー、データセットなどを追加します。
glq-quantize \
--モデル ハグフェイスTB/SmolLM2-360M \
--output ./smollm2-glq-4bpw \
--bpw 4 \
--nsサンプル 128 \
--デバイスCUDA
その他のビット幅: --bpw 2 から --bpw 8 を渡します (小数部分は次のようになります)
2.5 でも動作します)。 glq-quantize --help はすべてのフラグをリストします。模型用
システム RAM に収まらない場合は --streaming を使用します (一度に 1 つのレイヤーをロードします)
セーフティテンサーからの時間）。
混合精度の割り当ての場合は、2 パス フローを実行します。
pass はレイヤーごとの bpw_allocation.json を書き込み、次に量子化パスを書き込みます
それを適用します。 example/quantize_mixed_precision.md を参照してください。
事前に構築された CUDA イメージには、GLQ モデルを実行するために必要なものがすべて同梱されています。
CUDA 12.8 上の glq 、PyTorch、vLLM、transformers、および lm-eval:
ghcr.io/cnygaard/glq-env:0.5.0 # :latest、:0.5 も
前提条件 — Docker での GPU アクセス。 NVIDIA GPU プラスが必要です
NVIDIA コンテナ ツールキット
ホストにインストールされます。それが --gpus all フラグを渡す理由です
GPU をコンテナに組み込みます。動作することを確認します。
docker run --rm --gpus all ghcr.io/cnygaard/glq-env:0.5.0 nvidia-s

私
GPU テーブルが出力されれば準備完了です。 (ツールキットなし → --gpus
「デバイスドライバーを選択できませんでした」というエラー。)
出力を生成します。モデル キャッシュのホスト ディレクトリをマウントします (
イメージの HF_HOME は /cache/hf であるため、モデルは実行後も保持されます
再ダウンロードする代わりに）、次を生成します。
docker run --rm --gpus all \
-v " $HOME /.cache/huggingface:/cache/hf " \
ghcr.io/cnygaard/glq-env:0.5.0 \
Python -c '
import glq.hf_integration, torch # GLQ を HF に登録します
トランスフォーマーから AutoModelForCausalLM、AutoTokenizer をインポート
mid = "xv0y5ncu/SmolLM3-3B-GLQ-3.5bpw"
tok = AutoTokenizer.from_pretrained(mid)
モデル = AutoModelForCausalLM.from_pretrained(
Mid、device_map="cuda"、torch_dtype=torch.float16)
ids = tok("フランスの首都は", return_tensors="pt").to("cuda")
print(tok.decode(model.generate(**ids, max_new_tokens=20)[0],
Skip_special_tokens=True))
'
期待される出力:
フランスの首都はパリです。国の北部に位置しています。
最初の実行では、モデルがマウントされたキャッシュにダウンロードされます。後で実行します
それを再利用します。任意の GLQ チェックポイントの Mid を交換します (「
利用可能な事前量子化チェックポイント)。
サービス (vLLM) と対話型シェル。イメージには vLLM がバンドルされているため、
OpenAI 互換のエンドポイントを提供できます。ポートを公開し、
キャッシュをマウントします。
docker run --rm --gpus all -p 8000:8000 \
-v " $HOME /.cache/huggingface:/cache/hf " \
ghcr.io/cnygaard/glq-env:0.5.0 \
vllm サーブ xv0y5ncu/SmolLM3-3B-GLQ-3.5bpw
ロングコンテキスト E8 KV キャッシュ フラグ ( GLQ_KV_* ) の場合は、次のように渡します。
-e と E8 格子キャッシュを参照 /
インラインデカント E8 KV 。
対話的に操作したい場合は、イメージのデフォルトのコマンドはシェル ( docker run --rm -it --gpus all ghcr.io/cnygaard/glq-env:0.5.0 ) です。
SmolLM3-3B（マッチド 4.5 bpw 対 GPTQ）
Blackwell RTX PRO 6000、128 キャリブレーション サンプル、
lm-評価-ハーネス li

mit=200/タスク (GSM8K n=500、MMLU 50/サブタスク)。
GLQ 4.5 bpw は 2 パス混合割り当てを使用します (91 レイヤー @ 4 bpw + 161 @
5 bpw、平均 4.64 bpw）。
GLQ は 10/12 メトリクスで GPTQ を上回ります。 WikiText-2 と bf16 のギャップ: +2.2 %
(GLQ) 対 +6.2 % (GPTQ)。 GSM8K フレックスは bf16 と一致します。 GPTQは0.034低下。
小型モデル: SmolLM2-360M-Instruct (4 bpw)
GPTQ では、隠れたディムを分割するグループサイズが必要です。 SmolLM2-360Mの
hidden=960 は 128 で割り切れないため、group_size=64 (~4.5 eff) になります。
bpw)、品質が低下します。 GLQ にはグループサイズの制約がありません。
5 タスク = ARC-e、HellaSwag、PIQA、WinoGrande、LAMBADA。 128 校正
サンプル。 L40S。 GPTQ の LAMBADA は 0.346 に崩壊します。 GLQ は 0.508 を維持します。
スループット: vLLM 上の SmolLM3-3B
重みを圧縮すると DRAM が削減されるため、GLQ は bf16 に近いスループットで実行されます。
逆量子化コストをほぼ相殺するのに十分な帯域幅。
E8格子コードブック。最初の 7 つのシェルからの 65,536 個のベクトル
8 次元の E8 格子。各ウェイトの8ウェイトグループ
行列は、1 つの 16 ビット インデックスとしてこのコードブックにエンコードされます (つまり、
一次段階は 2 bpw)。 3 ～ 8 bpw の場合、追加の 8 ビット (256 エントリ)
または 16 ビット (E8) 残差コードブックにより、プライマリのコードブックが改良されます。
再構成エラー。
ランダム化されたアダマール変換。ランダムな符号反転に続いて
高速ウォルシュ アダマール変換は、重みとヘシアンの両方を回転します。
RHT の後、ヘッセ行列はほぼ対角線になるため、単純なユークリッド
コードブック内の最近傍は、以下の条件下ではほぼ最適です。
ヘシアン加重プロキシ損失。
LDLQ エラーのフィードバック。ヘシアンの Block-LDL 分解
GPTQ スタイルのシーケンシャル スイープを駆動しますが、代わりに 8-D ブロックを使用します
スカラー列の数。各ブロックの量子化誤差が伝播する
正しい下流ブロックに転送します。
融合された推論カーネル。カスタム CUDA C および Triton カーネルの読み取り
HBM からのコードブック インデックス。
L2 キャッシュ 1

MB コードブック、および matmul を直接蓄積します。
密な重み行列は決して実現されません。 GPU メモリの節約規模
圧縮率と一緒です。
GLQ には 2 つの KV キャッシュ コンプレッサーが同梱されています。どちらかがオプトイン — デフォルト
挙動は変わらない。
チャネルごとの absmax INT8 に加えて、最近の小さな fp16 残差ウィンドウ
トークン — KIVI スタイル。ロングコンテキストでの KV メモリを半分にします。
glq をインポートします。 hf_integration
glq から。 kv_cacheインポートGLQQuantizedCache
キャッシュ = GLQQuantizedCache (モデル . config )
出力 = モデル 。生成 ( ** 入力 、 max_new_tokens = 200 、
past_key_values = キャッシュ )
4.45 以上のトランスが必要です。外部依存関係はありません。
E8 ラティス キャッシュ (vLLM、v0.3.0+)
を使用して、vLLM のページ KV キャッシュを fp16 フットプリントの約 25 % にドロップします。
重み付けに使用されるのと同じ E8 格子量子化器。 2 つの融合された Triton カーネル
(読み取り側の dequant-gather、書き込み側の分散) デコードを範囲内に保つ
非融合スループットの約 20 %。
Gemma-4-E4B-it、RTX PRO 6000 Blackwell、vLLM 0.20 で測定:
GLQ_KV_QUANT=e8_relaxed:2 \
GLQ_KV_E8_SIDECAR=1 GLQ_KV_E8_SIDECAR_READ=1 \
GLQ_KV_E8_COMPRESSED_ALLOC=1 \
GLQ_KV_E8_FUSED_GATHER=1 GLQ_KV_E8_FUSED_WRITE=1 \
vllm は google/gemma-4-E4B-it を提供します
上記の環境ではワークスペース パスが使用されます。GLQ は、
K/V をスクラッチ バッファに参照し、vLLM のストックを呼び出します。
注意。そのバッファはデータに依存して構築されているため、
block_table.unique() 、glq 自動強制 cudagraph_mode=PIECEWISE
このパスの場合 ([glq_vllm] E8 KV active → cudagraph_mode Forced ... が起動時に PIECEWISE に設定されていることがわかります。 --enforce-eager はもう使用されていません)
v0.3.5以降は必須です）。重みのみの GLQ は引き続きデフォルトを使用します
フル_アンド_ピースワイズ 。以下の v0.5 inline-dequant パスがリフトされます
Piecewise 制限であり、推奨されるパスです。
ロングコンテキスト/KV バインドのサービング。
Gemma-4-E4B-it / Gem でエンドツーエンドで検証済み

vLLM の ma-4-31B-it
0.20.x。
Inline-dequant E8 KV (v0.5) — 長いコンテキストに推奨
上記のワークスペース パスは、K/V をスクラッチ バッファーに事前解凍します。
その後、vLLM のアテンションが再読み取りされます。K/V ごとに純粋なオーバーヘッドが発生します。
ベクトルは 1 回だけ読み取られます。 v0.5 は inline-dequant パスを追加します。
フォークされた Triton アテンション カーネルは圧縮された E8 K/V を逆量子化します
タイル ループ内 (逆方向の 8 ポイント FHT バタフライ)
アダマールに加えて、ロング コンテキスト占有用のフラッシュ デコード KV 分割)。
ワークスペースがありません。また、読み取り/書き込みフックが
host-sync-clean — フル CUDA グラフはデコード全体をキャプチャします。
大部分を占めていたトークンごとの即時ディスパッチのオーバーヘッドを排除します。
E8-KVデコード。これは、長いコンテキスト / の推奨パスです。
KV バウンドのサービング。
上記のバンドルに 1 つの env を追加して有効にします。
GLQ_KV_QUANT=e8_relaxed:2 \
GLQ_KV_E8_SIDECAR=1 GLQ_KV_E8_SIDECAR_READ=1 \
GLQ_KV_E8_COMPRESSED_ALLOC=1 \
GLQ_KV_E8_FUSED_GATHER=1 GLQ_KV_E8_FUSED_WRITE=1 \
GLQ_KV_E8_INLINE_DEQUANT_V3=1 \
vllm サーブ xv0y5ncu/SmolLM3-3B-GLQ-3.5bpw
デコード スループット、SmolLM3-3B-GLQ-3.5bpw、RTX PRO 6000 Blackwell、
vLLM 0.20.2 — インラインと v0.5 以前の E8-KV パス (ワークスペース、
ピースワイズ):
高速化は、インライン パスのロックが解除される完全なグラフ キャプチャです。それ

[切り捨てられた]

## Original Extract

E8 lattice codebook quantization for LLM weights — 2/3/4 bpw with fused Triton inference kernel - cnygaard/glq

I have with the help of AI create an open source method of E8 LLM code book quantization library called glq. I was interested in creating Glq
as a PC gamer and devops, interested in both LLMs and AI. The current high RAM prices and LLM resource usage also inspired me to write glq. A question arises could you try and squeeze more out a gaming GPU with limited VRAM size by using alternative LLM compression methods? Glq is effective compared to other LLM quantization algorithms at between 2-bits per weight up to 4 bits per weight. The effectiveness of glq at low bits per words is due to the properties of the E8 lattice compared to linear methods.
Glq also supports mixed precision quantization where different LLM layers uses different compression bit weight depending on how sensitive the LLM layers are to quantization. Think of mixed precision a bit like MP3 or MP4 variable bit rate encoding. I currently develop glq using g7e AWS spot instances to keep the cost more reasonable. Glq uses vllm 4 bit Key value cache by E8 was inspired by NexusQuant. I try and squeeze in about four times as much Key value cache as normally would fit by BF16 in VRAM, or about two times compared to INT8. I somehow wrongly at start picked a E8 code book size of 65536 entries instead of 4096 code book entries which better fits in GPU L1 cache. Having 65535 code book entries it turns out leads to higher LLM compression rate but at trade of of decode speed. I am trying to compensate by using Nvidia Cuda graphs and optimize the decode, currently work in progress. To install glq in a python virtual environment on Linux with a Nvidia GPU:
pip install glq Python PIP package
https://pypi.org/project/glq/ Glq source code.
https://github.com/cnygaard/glq Current PC RAM Prices that inspired the library.
https://pcpartpicker.com/trends/price/memory/ https://en.wikipedia.org/wiki/E8_lattice
Eight dimensional lattice that provides optimal solution to the sphere packing problems. Think about it a bit like stacking cannon balls or stacking apples in an optimal way. Only you swap the apples for LLM weights. Picture of an E8 lattice
https://en.wikipedia.org/wiki/E8_polytope#/media/File:E8_gra... Credits: GLQ was inspired by E8 Quip# and Key value E8 compression was inspired by NexusQuant. Math: The sphere packing problem in dimension 8, Maryna Viazovska
https://arxiv.org/abs/1603.04246 4bpw glq Quantization of Gemma 4 E4b-instruction tuned
https://huggingface.co/xv0y5ncu/Gemma-4-E4B-it-GLQ-4bpw 3.5bpw mixed precision quantization of SmolLM3
https://huggingface.co/xv0y5ncu/SmolLM3-3B-GLQ-3.5bpw Docker image of glq on Nvidia GPU with Nvidia container toolkit.
docker run --rm --gpus all \
-v "$HOME/.cache/huggingface:/cache/hf" \
ghcr.io/cnygaard/glq-env:0.5.0 \
python -c '
import glq.hf_integration, torch # registers GLQ with HF
from transformers import AutoModelForCausalLM, AutoTokenizer
mid = "xv0y5ncu/SmolLM3-3B-GLQ-3.5bpw"
tok = AutoTokenizer.from_pretrained(mid)
model = AutoModelForCausalLM.from_pretrained(
mid, device_map="cuda", torch_dtype=torch.float16)
ids = tok("The capital of France is", return_tensors="pt").to("cuda")
print(tok.decode(model.generate(*ids, max_new_tokens=20)[0],
skip_special_tokens=True))
' Currently work in progress on glq in getting the decode speed up and supporting more LLM model architectures. Open question, Does glq work on Nvidia DGX spark and gaming Nvidia hardware such as 4070-5090?

GitHub - cnygaard/glq: E8 lattice codebook quantization for LLM weights — 2/3/4 bpw with fused Triton inference kernel · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
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
cnygaard
/
glq
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
278 Commits 278 Commits .github .github benchmarks benchmarks examples examples glq glq glq_vllm glq_vllm infra infra scripts scripts tests tests .dockerignore .dockerignore .gitignore .gitignore Dockerfile Dockerfile LICENSE LICENSE README.md README.md compare_methods.py compare_methods.py eval_ppl.py eval_ppl.py pyproject.toml pyproject.toml test_glq_e2e.py test_glq_e2e.py View all files Repository files navigation
Post-training weight quantization for LLMs using E8 lattice codebooks .
GLQ encodes each 8-weight group as a 16-bit index into a 65,536-entry
E8 lattice codebook. A Randomized Hadamard Transform (RHT) decorrelates
the Hessian so that Euclidean nearest-neighbour search is near-optimal
under the proxy loss. The result: 2–8 bpw weights with quality
comparable to QuIP# / better than GPTQ, and a fused CUDA kernel that
matmuls directly against the compressed indices without materializing
the weight matrix.
pip install glq # requires PyTorch ≥ 2.0
Python ≥ 3.10. Triton ships with PyTorch on CUDA and is used
automatically. The CUDA C extension JIT-builds on first run
(~30 s); CPU falls back to dequantize-then-matmul.
import glq . hf_integration # registers GLQ with transformers
from transformers import AutoModelForCausalLM , AutoTokenizer
model = AutoModelForCausalLM . from_pretrained (
"xv0y5ncu/SmolLM2-360M-Instruct-GLQ-4bpw" ,
device_map = "auto" ,
)
tok = AutoTokenizer . from_pretrained ( "xv0y5ncu/SmolLM2-360M-Instruct-GLQ-4bpw" )
print ( tok . decode ( model . generate (
** tok ( "The capital of France is" , return_tensors = "pt" ). to ( model . device ),
max_new_tokens = 20 ,
)[ 0 ], skip_special_tokens = True ))
import glq.hf_integration registers quant_method="glq" with HF
Transformers; from_pretrained then swaps nn.Linear for E8RHTLinear
and uses the fused CUDA C kernel on inference. CPU falls back to a
naive dequantize-then-matmul.
Available pre-quantized checkpoints
Repo
Base model
bpw
License
xv0y5ncu/SmolLM2-135M-Instruct-GLQ-4bpw
SmolLM2-135M-Instruct
4.0
Apache 2.0
xv0y5ncu/SmolLM2-360M-Instruct-GLQ-4bpw
SmolLM2-360M-Instruct
4.0
Apache 2.0
xv0y5ncu/SmolLM3-3B-GLQ-3.5bpw
SmolLM3-3B
3.5 (mixed)
Apache 2.0
xv0y5ncu/Gemma-4-E4B-it-GLQ-4bpw
Gemma-4-E4B-it
4.0
Apache 2.0
xv0y5ncu/Devstral-Small-2-24B-Instruct-GLQ-4bpw
Devstral-Small 24B
4.0
Apache 2.0
xv0y5ncu/Nemotron-3-Nano-30B-A3B-GLQ-4bpw
Nemotron-3-Nano-30B (Mamba-MoE)
4.0
Nemotron
Quantize your own model
pip install ' glq[quantize] ' # adds transformers, datasets, etc.
glq-quantize \
--model HuggingFaceTB/SmolLM2-360M \
--output ./smollm2-glq-4bpw \
--bpw 4 \
--nsamples 128 \
--device cuda
Other bit-widths: pass --bpw 2 through --bpw 8 (fractional like
2.5 also works). glq-quantize --help lists every flag. For models
that don't fit in system RAM use --streaming (loads one layer at a
time from safetensors).
For mixed-precision allocation, run a two-pass flow: a profile
pass writes a per-layer bpw_allocation.json , then a quantize pass
applies it. See examples/quantize_mixed_precision.md .
A prebuilt CUDA image ships everything needed to run GLQ models —
glq , PyTorch, vLLM, transformers, and lm-eval on CUDA 12.8:
ghcr.io/cnygaard/glq-env:0.5.0 # also :latest, :0.5
Prerequisite — GPU access in Docker. You need an NVIDIA GPU plus
the NVIDIA Container Toolkit
installed on the host; that's what makes the --gpus all flag pass
the GPU into the container. Verify it works:
docker run --rm --gpus all ghcr.io/cnygaard/glq-env:0.5.0 nvidia-smi
If that prints your GPU table, you're set. (No toolkit → --gpus
errors with "could not select device driver".)
Produce output. Mount a host directory for the model cache (the
image's HF_HOME is /cache/hf , so models persist across runs
instead of re-downloading), then generate:
docker run --rm --gpus all \
-v " $HOME /.cache/huggingface:/cache/hf " \
ghcr.io/cnygaard/glq-env:0.5.0 \
python -c '
import glq.hf_integration, torch # registers GLQ with HF
from transformers import AutoModelForCausalLM, AutoTokenizer
mid = "xv0y5ncu/SmolLM3-3B-GLQ-3.5bpw"
tok = AutoTokenizer.from_pretrained(mid)
model = AutoModelForCausalLM.from_pretrained(
mid, device_map="cuda", torch_dtype=torch.float16)
ids = tok("The capital of France is", return_tensors="pt").to("cuda")
print(tok.decode(model.generate(**ids, max_new_tokens=20)[0],
skip_special_tokens=True))
'
Expected output:
The capital of France is Paris. It is located in the north of the country.
The first run downloads the model into the mounted cache; later runs
reuse it. Swap mid for any GLQ checkpoint (see
Available pre-quantized checkpoints ).
Serving (vLLM) & an interactive shell. The image bundles vLLM, so
you can serve an OpenAI-compatible endpoint — publish the port and
mount the cache:
docker run --rm --gpus all -p 8000:8000 \
-v " $HOME /.cache/huggingface:/cache/hf " \
ghcr.io/cnygaard/glq-env:0.5.0 \
vllm serve xv0y5ncu/SmolLM3-3B-GLQ-3.5bpw
For the long-context E8 KV-cache flags ( GLQ_KV_* ), pass them with
-e and see E8 lattice cache /
Inline-dequant E8 KV .
The image's default command is a shell ( docker run --rm -it --gpus all ghcr.io/cnygaard/glq-env:0.5.0 ) if you'd rather poke around interactively.
SmolLM3-3B at matched 4.5 bpw vs GPTQ
Blackwell RTX PRO 6000, 128 calibration samples,
lm-evaluation-harness limit=200/task (GSM8K n=500, MMLU 50/subtask).
GLQ 4.5 bpw uses two-pass mixed allocation (91 layers @ 4 bpw + 161 @
5 bpw, avg 4.64 bpw).
GLQ beats GPTQ on 10/12 metrics. WikiText-2 ppl gap to bf16: +2.2 %
(GLQ) vs +6.2 % (GPTQ). GSM8K flex matches bf16; GPTQ drops 0.034.
Small models: SmolLM2-360M-Instruct at 4 bpw
GPTQ requires a group-size dividing the hidden dim; SmolLM2-360M's
hidden=960 is not divisible by 128, forcing group_size=64 (~4.5 eff
bpw) and losing quality. GLQ has no group-size constraint.
5-task = ARC-e, HellaSwag, PIQA, WinoGrande, LAMBADA; 128 calibration
samples; L40S. GPTQ's LAMBADA collapses to 0.346; GLQ preserves 0.508.
Throughput: SmolLM3-3B on vLLM
GLQ runs at near-bf16 throughput because compressed weights cut DRAM
bandwidth enough to roughly offset the dequantization cost.
E8 lattice codebook. 65,536 vectors from the first seven shells
of the E8 lattice in 8 dimensions. Each 8-weight group of the weight
matrix is encoded as one 16-bit index into this codebook (so the
primary stage is 2 bpw). For 3–8 bpw, additional 8-bit (256-entry)
or 16-bit (E8) residual codebooks refine the primary's
reconstruction error.
Randomized Hadamard Transform. Random sign flips followed by
Fast Walsh-Hadamard Transform rotate both weights and Hessian.
After RHT the Hessian is approximately diagonal, so plain Euclidean
nearest-neighbour in the codebook is near-optimal under the
Hessian-weighted proxy loss.
LDLQ error feedback. Block-LDL decomposition of the Hessian
drives a sequential sweep — GPTQ-style, but over 8-D blocks instead
of scalar columns. Each block's quantization error propagates
forward to correct downstream blocks.
Fused inference kernels. Custom CUDA C and Triton kernels read
codebook indices from HBM, gather the 8-D vectors from the
L2-cached 1 MB codebook, and accumulate the matmul directly — the
dense weight matrix is never materialized. GPU memory savings scale
with the compression ratio.
GLQ ships two KV cache compressors. Either is opt-in — default
behaviour is unchanged.
Per-channel absmax INT8 plus a small fp16 residual window for recent
tokens — KIVI-style. Halves the KV memory at long context.
import glq . hf_integration
from glq . kv_cache import GLQQuantizedCache
cache = GLQQuantizedCache ( model . config )
output = model . generate ( ** inputs , max_new_tokens = 200 ,
past_key_values = cache )
Requires transformers >= 4.45 . No external dependencies.
E8 lattice cache (vLLM, v0.3.0+)
Drops vLLM's paged KV cache to ~25 % of fp16 footprint using the
same E8 lattice quantizer used for weights. Two fused Triton kernels
(read-side dequant-gather, write-side scatter) keep decode within
~20 % of un-fused throughput.
Measured on Gemma-4-E4B-it, RTX PRO 6000 Blackwell, vLLM 0.20:
GLQ_KV_QUANT=e8_relaxed:2 \
GLQ_KV_E8_SIDECAR=1 GLQ_KV_E8_SIDECAR_READ=1 \
GLQ_KV_E8_COMPRESSED_ALLOC=1 \
GLQ_KV_E8_FUSED_GATHER=1 GLQ_KV_E8_FUSED_WRITE=1 \
vllm serve google/gemma-4-E4B-it
The envs above use the workspace path : GLQ pre-decompresses the
referenced K/V into a scratch buffer, then calls vLLM's stock
attention. Because that buffer is built with a data-dependent
block_table.unique() , glq auto-forces cudagraph_mode=PIECEWISE
for this path (you'll see [glq_vllm] E8 KV active → cudagraph_mode forced ... to PIECEWISE at startup; --enforce-eager is no longer
required as of v0.3.5). Weight-only GLQ still uses the default
FULL_AND_PIECEWISE . The v0.5 inline-dequant path below lifts
the PIECEWISE restriction and is the recommended path for
long-context / KV-bound serving.
Validated end-to-end on Gemma-4-E4B-it / Gemma-4-31B-it on vLLM
0.20.x.
Inline-dequant E8 KV (v0.5) — recommended for long-context
The workspace path above pre-decompresses K/V into a scratch buffer
that vLLM's attention then re-reads — pure overhead, since each K/V
vector is read exactly once. v0.5 adds an inline-dequant path : a
forked Triton attention kernel dequantizes the compressed E8 K/V
inside the tile loop (an 8-point FHT butterfly for the inverse
Hadamard, plus flash-decoding KV-split for long-context occupancy).
There is no workspace, and — because the read/write hooks are
host-sync-clean — the FULL CUDA graph captures the whole decode ,
eliminating the per-token eager-dispatch overhead that dominated
E8-KV decode. This is the recommended path for long-context /
KV-bound serving .
Enable it by adding one env to the bundle above:
GLQ_KV_QUANT=e8_relaxed:2 \
GLQ_KV_E8_SIDECAR=1 GLQ_KV_E8_SIDECAR_READ=1 \
GLQ_KV_E8_COMPRESSED_ALLOC=1 \
GLQ_KV_E8_FUSED_GATHER=1 GLQ_KV_E8_FUSED_WRITE=1 \
GLQ_KV_E8_INLINE_DEQUANT_V3=1 \
vllm serve xv0y5ncu/SmolLM3-3B-GLQ-3.5bpw
Decode throughput, SmolLM3-3B-GLQ-3.5bpw, RTX PRO 6000 Blackwell,
vLLM 0.20.2 — inline vs the pre-v0.5 E8-KV path (workspace,
PIECEWISE):
The speedup is the FULL-graph capture the inline path unlocks; it

[truncated]
