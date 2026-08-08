---
source: "https://snwagh.com/blog/2026/stained-glass-transform/"
hn_url: "https://news.ycombinator.com/item?id=49223345"
title: "Prompt Privacy from LLMs"
article_title: "A novel method for input privacy from LLMs | Sameer Wagh"
author: "napping_penguin"
captured_at: "2026-08-08T17:21:27Z"
capture_tool: "hn-digest"
hn_id: 49223345
score: 2
comments: 1
posted_at: "2026-08-08T16:34:42Z"
tags:
  - hacker-news
  - translated
---

# Prompt Privacy from LLMs

- HN: [49223345](https://news.ycombinator.com/item?id=49223345)
- Source: [snwagh.com](https://snwagh.com/blog/2026/stained-glass-transform/)
- Score: 2
- Comments: 1
- Posted: 2026-08-08T16:34:42Z

## Translation

タイトル: LLM からの迅速なプライバシー
記事のタイトル: LLM からの入力プライバシーのための新しい方法 |サミール・ワグ
説明: ステンド グラス トランスフォームは、LLM 推論の下でトークンの埋め込みの有用性を維持しながら、トークンの埋め込みをスクランブルするように訓練されたトランスフォーマー ベースのネットワークです。出力はいくぶん魔法のようです。モデルは依然として一貫した出力を生成しますが、元のプロンプトは回復できません。

記事本文:
Sameer Wagh ナビゲーションを切り替える
LLM からの入力プライバシーのための新しい方法
2026 · プライバシー LLM 埋め込み · 研究
私がこれまでに出会った中で最も興味深いプライバシー テクノロジーの 1 つは、Stained Glass Transform (SGT) と呼ばれるものです。これは、Protopia AI の人々によって発明されました (彼らのチームには私の才能ある友人であり共同研究者の Sid Roy が含まれており、このブログでは彼らの技術文書 [1] を調査しています。これが解決する問題は、クラウド LLM API 上で構築する人なら誰でも遭遇する問題です。つまり、モデルのインテリジェンスは必要ですが、LLM プロバイダーにプロンプト/データを見られたくないのです。
ホストされている LLM エンドポイント (ChatGPT、Claude.ai、OpenRouter、HuggingFace など) を呼び出すと、プロンプトが平文でサードパーティのサーバーに渡され、そのプロンプトがデータベース ログに保存されます。プロンプトの個人的な性質とデータ エコノミーの仕組みを考慮すると、これは大きな懸念事項です。したがって、ユーザーが入力のプライバシーを維持しながら LLM のメリットを享受できるメカニズムが重要です。
この緊急のプライバシー問題を解決するには、いくつかの方法があります。
$\textbf{ローカル ホスティング。}$ プロンプトが環境から離れることがないように、モデルを自分でホストします。 Ollama はこれを簡単にし、単一のコマンドで Llama、Mistral、Gemma、およびその他のオープンウェイト モデルをコンシューマ ハードウェア上で実行できるようにします。明らかな制限はコンピューティングです。有能なモデルには十分な VRAM を備えた GPU が必要です。さらに、ホスト型エンドポイントに無料で付属するすべてのインフラストラクチャ (負荷分散、自動スケーリング、自動再試行、ハードウェア メンテナンス、実稼働環境でモデル サーバーを健全に保つための運用オーバーヘッド) が失われます。
$\textbf{完全準同型暗号化 (FHE)。}$ FHE では、暗号化されたデータを直接計算できるため、プロンプトはデバイス上で暗号化され、

erver はそれを復号化せずに処理します。この Belfort Labs のデモは、FHE ベースの推論が実際にどのようなものかを体験できる、ブラウザー内でのライブ エクスペリエンスです。オープンソースの側面では、Zama の Concrete ML は、基礎となるハード暗号エンジニアリングに取り組む主要なライブラリです。欠点は深刻です。FHE 推論は平文よりも遅く、LM エンドポイントは暗号化された算術演算 (平文-暗号文) で動作するには大幅な再エンジニアリングが必要であり、大規模な鍵管理は重要な運用上の課題です。
$\textbf{信頼された実行環境 (TEE)。}$ TEE (例: Intel SGX/TDX、AMD SEV、Confidential Container) は、コードとデータがホスト OS やクラウド プロバイダーからも隠蔽される、ハードウェアから分離されたエンクレーブを作成します。これを使用すると、サーバーがユーザーのプロンプトを見ることができなくなり、同時にモデルプロバイダーの重みを機密に保つことができる両面プライバシーを実行できます。実際には、ユーザーは依然としてハードウェア ベンダーの証明書を信頼する必要があり、GPU TEE サポート (パフォーマンスの推論に必要) は比較的新しく (NVIDIA Hopper は実稼働対応の機密コンピューティングを備えた第 1 世代です)、TEE ホスティング エンティティに関する信頼に関する質問は、プライバシーの保証を完全に損なう可能性があります。
ステンド グラス トランスフォームは、よく研究された厳格なプライバシーの概念を使用して、同じ問題に対する新しい解決策です。この解決策には、生のテキストの代わりに難読化された埋め込みを LLM プロバイダーに送信し、プロバイダーのエンドポイントに残りの処理を行わせることが含まれます。
言い換えれば、すべての LLM で使用される最初の準備段階 (トークン化と埋め込み) がユーザー側に移されます。トレーニングされた機械学習モデル (秘密のソース) を使用して、埋め込み (およびプロンプト) が難読化されます。ただし、重要な洞察は、この難読化された p

rompt は、経験的に検証された 2 つの保証を提供します。
$\textbf{(ユーティリティ保存)}$ 難読化されたプロンプトの LLM 出力は、生のテキストの LLM 出力に似ています。
$\textbf{(プライバシー保証)}$ 生のテキスト プロンプトを難読化された埋め込みからリバース エンジニアリングするのは困難です。
このインタラクティブなウォークスルーは、ラップトップ サイズのディスプレイおよび縦向きのモバイル デバイス向けに最適化されています。他のビューポート (横向きのモバイルやタブレットなど) は引き続き機能しますが、レイアウトの忠実度が低下する可能性があります。
Embedder d = 2,048 Embedder d = 2,048 ステンド グラス 難読化された埋め込みを変換する
トレーニング済みモデルの使用 LLM LLaMA 3.2 1B · 凍結された LLM LLaMA 3.2 1B · 凍結された応答（従来型） 応答（プロンプトプライベート） トークン埋め込み d = 2,048 dims · 7 トークン LLaMA 3.2 1B SGT のトレーニング
SGT の論文はよく書かれており、この投稿では単に彼らのアプローチに従っただけです。アーキテクチャなどの実装の選択について説明しますが (IP 上の理由により、この論文では詳しく説明されていない可能性があります)、詳細についてはこの論文を参照することをお勧めします。大まかな考え方は、埋め込みシーケンスを取得し、それを摂動バージョンに置き換える小規模なローカル ネットワーク（SGT と呼ばれる）を実行するというものです。したがって、サーバーはトークンや生の埋め込みを決して認識しません。スクランブルされたバージョンのみを処理します。この作業の核心は、ユーティリティを維持しながら埋め込みプライバシーを維持するように SGT を効率的にトレーニングする方法を示すことです。つまり、LLM の出力品質が低下しないようにする必要があります。
SGT は、決定論的シフト $\mu_\theta$ と次元ごとのノイズ スケール $\sigma_\theta$ を予測します。学習された分散を使用してガウス ノイズを追加することは、2 つのパスで同じ難読化された埋め込みが生成されないことを意味します。これは、反復クエリ攻撃に抵抗するために重要です。
SGT はモデルごとにトレーニングされます。

のブログです。私は論文のモデル、Llama 3.2 1B を使用しています。フリーズした LLM の前に配置する SGT モジュールとして、次の小さなポストノルム トランス エンコーダを選択しました。
クラス SGT ( nn . モジュール ):
def __init__ ( self 、 embed_dim = 2048 、 num_layers = 2 、 nhead = 8 ):
超（）。 __init__ ()
enc_layer = nn 。 TransformerEncoderLayer (
d_model = embed_dim 、 nhead = nhead 、
dim_feedforward = embed_dim * 2 、ドロップアウト = 0.0 、
patch_first = True 、norm_first = False 、 # post-norm は出力の制限を維持します
)
自分自身。エンコーダ = nn 。トランスフォーマーエンコーダー (
enc_layer 、num_layers = num_layers 、norm = nn 。 LayerNorm ( embed_dim )
)
自分自身。 mu_head = nn 。線形 ( embed_dim 、 embed_dim )
自分自身。 log_sigma_head = nn 。線形 ( embed_dim 、 embed_dim )
# ID として初期化: mu=0、小さなシグマ
ん。初期化 。 zeros_(self .mu_head .weight );ん。初期化 。 zeros_ (self . mu_head .bias )
ん。初期化 。 zeros_ (self . log_sigma_head .weight )
ん。初期化 。 constant_ (self . log_sigma_head .bias 、 - 2.0 )
def forward ( self 、 x 、padding_mask = None ):
h = 自分自身。エンコーダー ( x . float ()、src_key_padding_mask = padding_mask )
ムー = 自分自身。 mu_head ( h )
log_sigma = self 。 log_sigma_head ( h )。クランプ (-6.0、3.0)
eps = トーチ。 randn_like ( h )
x_チルダ = x 。 float () + mu + log_sigma 。 exp () * eps
x_チルダを返します。 to ( x . dtype )、mu . ( x . dtype )、log_sigma . to ( x . dtype )
Llama 3.2 1B (embed_dim=2048) の場合、この SGT には 7,550 万個のパラメータがあります。これは、保護する LLM のサイズの約 6% です。トークンごとにミリ秒単位でローカルで実行されます。 LLM は移動する必要がありません。
トレーニングでは、1 つのユーティリティ損失コンポーネントと 3 つの難読化損失コンポーネントという 4 つの目的を同時にバランスよく調整します (詳細については論文を参照してください。著者は課題とその選択について詳しく説明しています)。 40,000 以上の OpenOrca サンプルをトレーニングしました

Google Colab T4 GPU で 5000 ステップ (最良のチェックポイントはステップ 4500 でした)。
ユーティリティ — 難読化されたシーケンスは、クリーンなシーケンスと同じ次のトークンの分布を生成する必要があります。 128 K 語彙 LM の確率質量は多くのトークンに分散されるため、ハードラベルのクロスエントロピーの代わりに KL ダイバージェンスを使用します。ハードラベルの勾配は、16 の凍結されたトランス層による難読化損失と競合するにはまばらすぎます。
def loss_utility ( logits_obf 、 logits_clean ):
log_p_obf = F 。 log_softmax ( logits_obf . reshape ( - 1 , V ) . float (), dim =- 1 )
p_clean = F 。 Softmax ( logits_clean . detach (). reshape ( - 1 , V ). float (), dim =- 1 )
リターンF。 kl_div ( log_p_obf , p_clean ,duction = "バッチ平均" )
AbsCosine — 難読化された埋め込みを元の埋め込みと直交するようにプッシュします。 $\lvert\cos(\tilde{x}, x)\rvert$ がゼロに近い場合、最近傍攻撃は元のトークンを見つけることができません。
def loss_abscosine ( x , x_tilde ):
cos = F 。 cosine_similarity ( x . reshape ( - 1 , D ), x_tilde . reshape ( - 1 , D ), dim =- 1 )
返品送料も負担いたします。腹筋（）。意味（）
ノルム ペナルティ — 難読化されたノルムをトークンごとにクリーンなノルムに近づけて、LLM の内部正規化が期待どおりに動作するようにします。
def loss_norm_penalty ( x , mu ):
clean_norms = x 。浮動小数点()。ノルム ( dim =- 1 )。切り離す()
shift_norms = ( x . float () + mu . float ())。ノルム ( dim =- 1 )
return (shift_norms - clean_norms)。腹筋（）。意味（）
相互情報量 — \(I(\tilde{x}; x)\) のミニバッチ モンテカルロ推定 (次元ごとの nat 単位)。取り消しを避けるために float64 で計算されます。これにより、点単位だけでなく、学習された分布全体にわたって \(x\) について \(\tilde{x}\) が保持する情報量が直接最小化されます。
def loss_mi ( x_tilde_A 、 mu_A 、 log_sigma_A 、 x_clean_B 、 mu_B 、 log_sigma_B ):
# 対角線からの H(x̃ | x)

ガウス成分のエントロピー
H_comp = ( 0.5 * LOG_2PIE + log_sigma_A_64 )。合計 ( dim = ( - 1 , - 2 ))。意味（）
# H(x̃) ≈ -E[log p_mix(x̃)] ミニバッチ GMM 経由
log_prob = - 0.5 * ( diff . pow ( 2 ) / var_B + log_const )。 sum ( dim = ( - 1 , - 2 ))
H_mix = - トーチ。 logsumexp ( log_prob 、 dim = 1 )。平均 () + 数学 .ログ (B_B)
return (( H_mix - H_comp ) / ( T * d ))。 float ()
最終的な結合損失には、重み \((\alpha_u, \alpha_\text{acs}, \alpha_\text{norm}, \alpha_\text{mi}) = (2.0, 0.3, 0.05, 0.15)\) が使用されます。これらの重みを正しく取得するには 3 回の反復が必要でした。主な故障モードは \(\alpha_u\) が非常に大きいため、効用損失が \(\sigma\) を小さく保ち、訓練全体を通じて相互情報量が高いままになります。損失曲線は以下のとおりです。
5000 ステップを超えるトレーニング損失曲線。ユーティリティ (KL) 損失は早期に安定しますが、3 つの難読化コンポーネントは引き続き改善されます。ステップ 4500 の最良のチェックポイントは、4 つの目的すべてのバランスをとります。プライバシーの指標と有用性が論文で報告されているものよりも悪い場合、トレーニングはおそらく合理的な局所最適であることに注意してください。
LLM 出力は入力を提供しますか?
この論文では単純な攻撃ベースラインについて説明しており、同じ著者が BeamClean と呼ばれるより優れた再構成攻撃も構築しています [2] 。彼らの攻撃は埋め込みベクトルのみを考慮しているため、より強力な攻撃者 (モデルによって生成されたテキスト出力も確認できる攻撃者) が BeamClean を改善できるかどうかに興味がありました。 BeamClean [2] は、難読化された埋め込みとのコサイン類似度によって各トークン位置で上位の語彙候補を見つけ、事前に言語モデルでスコアを付け、ビーム検索を実行します。
言語モデルによる q の悪用を防ぐために、観測された LLM 出力を正規化を伴う追加信号として使用する 2 つの拡張機能を実装しました。

文字化けした出力に問題が発生します (例: 表面的な理由で、does よりも Does を優先します)。ただし、結果はまちまちであり、広く一般化できるほど重要ではありません。現在の証拠は、BeamClean+output が BeamClean 単独よりも強力ではないことを示唆していますが、私はそれを厳密に検証するための未解決の質問として残しておきます。
SGT は本当に賢いアイデアです。重要な洞察は、埋め込み層を公開できれば、それを推論パイプラインから分離して強力なプライバシーを実現できるということです。これは、モデルの所有者がモデルの所有権を保持し、ユーザーが迅速なプライバシーを確​​保できる優れた中間点となる可能性があります。
(1) 難読化と (2) ユーティリティの保存という 2 つの相反する目的をうまく達成するようにモデルをトレーニングできることは、私にとって本当に驚きです。この点で、SGT は完全準同型暗号化と同じくらい革新的だと感じます。
限られたリソースでモデルをゼロからトレーニングすることができました。これは主に論文がよく書かれているおかげであり、学術コミュニティの知識共有の文化を物語っています。
論文で報告されている NN-FR 0.93 を完全に再現するには、おそらく 5,000 を大幅に超えるトレーニング ステップが必要です。私のチェックポイントは、まだプライバシー レベルではないものの、アプローチが方向性を持って機能するという有用な概念実証です。

[切り捨てられた]

## Original Extract

The Stained Glass Transform is a transformer-based network trained to scramble token embeddings while preserving their utility under LLM inference. The output is somewhat magical: the model still produces coherent output but the original prompt is unrecoverable.

Sameer Wagh Toggle navigation about
A novel method for input privacy from LLMs
2026 · privacy llm embeddings · research
One of the most interesting privacy technologies that I have come across is called Stained Glass Transform (SGT). This was invented by folks at Protopia AI (their team includes my talented friend and collaborator Sid Roy and in this blog I am looking into their technical paper [1]. The problem it addresses is one that anyone building on cloud LLM APIs encounters: you want the model’s intelligence, but you don’t want the LLM provider to see your prompt/data.
When you call any hosted LLM endpoint (think ChatGPT, Claude.ai, OpenRouter, HuggingFace), you hand your prompt in the clear to a third-party server, which stores it in their database logs. This is a major concern given the increasingly personal nature of prompts and the mechanics of the data economy. A mechanism to let users benefit from LLMs while preserving the privacy of their input is therefore critical.
There are a few different ways to resolve this prompt privacy challenge.
$\textbf{Local hosting.}$ Host the model yourself so the prompt never leaves your environment. Ollama makes this straightforward, letting you run Llama, Mistral, Gemma, and other open-weight models on consumer hardware with a single command. The obvious limitation is compute: a capable model needs a GPU with sufficient VRAM. Beyond that, you forfeit all the infrastructure that comes for free with hosted endpoints: load balancing, auto-scaling, automatic retries, hardware maintenance, and the operational overhead of keeping a model server healthy in production.
$\textbf{Fully Homomorphic Encryption (FHE).}$ FHE allows computations directly on encrypted data so your prompt is encrypted on-device and the server processes it without ever decrypting it. This Belfort Labs demo is a live in-browser experience that gives a feel for what FHE-based inference looks like in practice. On the open-source side, Zama’s Concrete ML is the leading library tackling the underlying hard cryptographic engineering. The downsides are steep: FHE inference is slower than plaintext, LM endpoints need significant re-engineering to operate over encrypted arithmetic (plaintext-ciphertext), and key management at scale is a non-trivial operational challenge.
$\textbf{Trusted Execution Environments (TEEs).}$ TEEs (e.g. Intel SGX/TDX, AMD SEV, Confidential Containers) create hardware-isolated enclaves where code and data are hidden even from the host OS and cloud provider. This can be used to perform two-sided privacy where the server cannot see the user’s prompt and the model provider’s weights can simultaneously remain confidential. In practice, the user must still trust the hardware vendor’s attestation, GPU TEE support (needed for performant inference) is relatively new (NVIDIA Hopper is the first generation with production-ready confidential computing), and trust questions around the TEE hosting entity can undermine the privacy guarantees entirely.
The Stained Glass Transform is a novel solution to the same problem with a well-studied and rigorous notion of privacy. The solution involves sending obfuscated embeddings instead of raw text to the LLM provider and letting the provider’s endpoint do the rest.
In other words, it moves the initial preparatory stages used by all LLMs (tokenization and embedding) to the user’s side. Using a trained machine learning model (their secret sauce), the embedding (and thus the prompt) is obfuscated. The key insight, however, is that this obfuscated prompt provides two empirically validated guarantees:
$\textbf{(Utility preservation)}$ The LLM output on the obfuscated prompt is close to the LLM output on the raw text.
$\textbf{(Privacy guarantee)}$ The raw text prompt is hard to reverse-engineer from the obfuscated embeddings.
This interactive walkthrough is optimized for laptop-sized displays and mobile devices in portrait orientation. Other viewports — including landscape mobile and tablet — remain functional but may exhibit reduced layout fidelity.
Embedder d = 2,048 Embedder d = 2,048 Stained Glass Transform obfuscate embeddings
using a trained model LLM LLaMA 3.2 1B · frozen LLM LLaMA 3.2 1B · frozen Response (conventional) Response (prompt private) Token Embeddings d = 2,048 dims · 7 tokens LLaMA 3.2 1B Training the SGT
The SGT paper is well-written and in this post I have simply followed their approach. While I describe my implementation choices such as architecture (which may not be fully detailed in the paper for IP reasons), I encourage the reader to refer to the paper for further details. The high-level idea is that you run a small local network (called the SGT) that takes the embedding sequence and replaces it with a perturbed version. Thus, the server never sees tokens or raw embeddings; it only ever processes the scrambled version. The crux of the work is showing how to efficiently train the SGT to preserve embedding privacy while retaining utility — that is, LLM output quality should not degrade.
The SGT predicts a deterministic shift $\mu_\theta$ and a per-dimension noise scale $\sigma_\theta$. Adding Gaussian noise with a learned variance means no two passes produce the same obfuscated embeddings — which is important for resisting repeated-query attacks.
SGT is trained per model and in this blog, I use the model from the paper — Llama 3.2 1B. I chose the following small post-norm transformer encoder as the SGT module placed in front of the frozen LLM:
class SGT ( nn . Module ):
def __init__ ( self , embed_dim = 2048 , num_layers = 2 , nhead = 8 ):
super (). __init__ ()
enc_layer = nn . TransformerEncoderLayer (
d_model = embed_dim , nhead = nhead ,
dim_feedforward = embed_dim * 2 , dropout = 0.0 ,
batch_first = True , norm_first = False , # post-norm keeps output bounded
)
self . encoder = nn . TransformerEncoder (
enc_layer , num_layers = num_layers , norm = nn . LayerNorm ( embed_dim )
)
self . mu_head = nn . Linear ( embed_dim , embed_dim )
self . log_sigma_head = nn . Linear ( embed_dim , embed_dim )
# initialize as identity: mu=0, small sigma
nn . init . zeros_ ( self . mu_head . weight ); nn . init . zeros_ ( self . mu_head . bias )
nn . init . zeros_ ( self . log_sigma_head . weight )
nn . init . constant_ ( self . log_sigma_head . bias , - 2.0 )
def forward ( self , x , padding_mask = None ):
h = self . encoder ( x . float (), src_key_padding_mask = padding_mask )
mu = self . mu_head ( h )
log_sigma = self . log_sigma_head ( h ). clamp ( - 6.0 , 3.0 )
eps = torch . randn_like ( h )
x_tilde = x . float () + mu + log_sigma . exp () * eps
return x_tilde . to ( x . dtype ), mu . to ( x . dtype ), log_sigma . to ( x . dtype )
For Llama 3.2 1B (embed_dim=2048), this SGT has 75.5 M parameters — about 6% the size of the LLM it protects. It runs locally in milliseconds per token; the LLM never needs to move.
Training balances four objectives simultaneously: one utility loss and three obfuscation loss components (refer to the paper for more details — the authors explain the challenges and their choices well). I trained over 40K OpenOrca examples for 5000 steps on a Google Colab T4 GPU (best checkpoint was at step 4500).
Utility — the obfuscated sequence should produce the same distribution of next tokens as the clean sequence. I use KL divergence instead of hard-label cross-entropy, because a 128 K-vocab LM’s probability mass is spread across many tokens. Hard-label gradients are too sparse to compete with the obfuscation losses through 16 frozen transformer layers.
def loss_utility ( logits_obf , logits_clean ):
log_p_obf = F . log_softmax ( logits_obf . reshape ( - 1 , V ). float (), dim =- 1 )
p_clean = F . softmax ( logits_clean . detach (). reshape ( - 1 , V ). float (), dim =- 1 )
return F . kl_div ( log_p_obf , p_clean , reduction = " batchmean " )
AbsCosine — push the obfuscated embedding orthogonal to the original. If $\lvert\cos(\tilde{x}, x)\rvert$ is near zero, the nearest-neighbour attack can’t find the original token:
def loss_abscosine ( x , x_tilde ):
cos = F . cosine_similarity ( x . reshape ( - 1 , D ), x_tilde . reshape ( - 1 , D ), dim =- 1 )
return cos . abs (). mean ()
Norm penalty — keep obfuscated norms close to clean norms per token, so the LLM’s internal normalizations behave as expected:
def loss_norm_penalty ( x , mu ):
clean_norms = x . float (). norm ( dim =- 1 ). detach ()
shifted_norms = ( x . float () + mu . float ()). norm ( dim =- 1 )
return ( shifted_norms - clean_norms ). abs (). mean ()
Mutual information — a minibatch Monte Carlo estimate of \(I(\tilde{x}; x)\) in nats per dimension, computed in float64 to avoid cancellation. This directly minimizes how much information \(\tilde{x}\) retains about \(x\) across the learned distribution, not just pointwise:
def loss_mi ( x_tilde_A , mu_A , log_sigma_A , x_clean_B , mu_B , log_sigma_B ):
# H(x̃ | x) from diagonal Gaussian component entropy
H_comp = ( 0.5 * LOG_2PIE + log_sigma_A_64 ). sum ( dim = ( - 1 , - 2 )). mean ()
# H(x̃) ≈ -E[log p_mix(x̃)] via minibatch GMM
log_prob = - 0.5 * ( diff . pow ( 2 ) / var_B + log_const ). sum ( dim = ( - 1 , - 2 ))
H_mix = - torch . logsumexp ( log_prob , dim = 1 ). mean () + math . log ( B_B )
return (( H_mix - H_comp ) / ( T * d )). float ()
The final combined loss uses weights \((\alpha_u, \alpha_\text{acs}, \alpha_\text{norm}, \alpha_\text{mi}) = (2.0, 0.3, 0.05, 0.15)\). Getting these weights right took three iterations — the main failure mode is \(\alpha_u\) so large that the utility loss keeps \(\sigma\) tiny, leaving mutual information high throughout training. The loss curves are below:
Training loss curves over 5000 steps. The utility (KL) loss stabilizes early while the three obfuscation components continue to improve; the best checkpoint at step 4500 balances all four objectives. Note that the training is probably a reasonable local optimum given that the privacy metrics and utility are worse than those reported in the paper.
Does the LLM output give away the input?
The paper covers simple attack baselines and the same authors also construct a better reconstruction attack called BeamClean [2] . Given that their attack only considers the embedding vector, I was curious to see if a stronger attacker — one that can also see the text output produced by the model — could improve on BeamClean. BeamClean [2] finds the top vocabulary candidates at each token position by cosine similarity to the obfuscated embedding, scores them with a language-model prior, and runs beam search.
I implemented two extensions that use the observed LLM output as an additional signal with regularization to prevent the language model from exploiting quirks in the garbled output (e.g., preferring Does over does for superficial reasons). The results, however, have been mixed and not significant enough to generalize broadly. While the current evidence suggests that BeamClean+output is no stronger than BeamClean alone, I leave it as an open question to rigorously verify.
SGT is a genuinely clever idea. The key insight is that if the embedding layer can be made public, you can separate it from the inference pipeline to achieve strong privacy. This can be a great middle ground where the model owner retains ownership of the model while the user gets prompt privacy.
It is genuinely surprising to me that a model can be trained to achieve two contrary objectives well: (1) obfuscation and (2) utility preservation. In this regard, SGT feels just as innovative as fully homomorphic encryption.
I was able to train the model from scratch with limited resources. This is largely a credit to the paper being well-written and speaks to the academic community’s culture of knowledge sharing.
Fully reproducing the paper’s reported NN-FR of 0.93 likely requires significantly more than 5,000 training steps. My checkpoint is a useful proof of concept that the approach works directionally, though not yet at the privacy levels c

[truncated]
