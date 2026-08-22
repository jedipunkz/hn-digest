---
source: "https://huggingface.co/blog/mayafree/model-dna"
hn_url: "https://news.ycombinator.com/item?id=49400676"
title: "Model Genome: Fingerprinting Whether an LLM Was Trained from Scratch or Derived"
article_title: "Model Genome: Fingerprinting Whether an LLM Was Trained From Scratch or Derived"
image: "https://cdn-uploads.huggingface.co/production/uploads/696f2edfa0417065e6a7c3ae/7GjZ5RgAScqJRkI808x3e.png"
author: "gmays"
captured_at: "2026-08-22T16:12:09Z"
capture_tool: "hn-digest"
hn_id: 49400676
score: 1
comments: 0
posted_at: "2026-08-22T15:20:59Z"
tags:
  - hacker-news
  - translated
---

# Model Genome: Fingerprinting Whether an LLM Was Trained from Scratch or Derived

- HN: [49400676](https://news.ycombinator.com/item?id=49400676)
- Source: [huggingface.co](https://huggingface.co/blog/mayafree/model-dna)
- Score: 1
- Comments: 0
- Posted: 2026-08-22T15:20:59Z

## Translation

タイトル: モデルゲノム: LLM が最初からトレーニングされたか派生されたかどうかのフィンガープリンティング
記事のタイトル: モデルゲノム: LLM が最初からトレーニングされたか派生されたかどうかのフィンガープリンティング
説明: Proto_AGI による Hugging Face に関するブログ投稿

記事本文:
モデルゲノム: LLM が最初からトレーニングされたのか、それとも派生したのかをフィンガープリンティングする
ハグ顔モデル
記事に戻る a]:hidden">
モデルゲノム: LLM が最初からトレーニングされたのか、それとも派生したのかをフィンガープリンティングする
コミュニティ記事が公開されました
2026 年 8 月 8 日 賛成票 18
config.json</code>)"> 2. 軸 1 — アーキテクチャのフィンガープリント ( config.json )
3. 軸 2 — トークナイザーの指紋 (親子鑑定)
4. 軸 3 — 重み付けフィンガープリント (難しいもの) トラップ 1 — 行単位のコサインは役に立たない
トラップ 2 — CKA は役立ちますが、十分ではありません
5. ボーナス軸 — 独創性の代理としての多様性への注目
6. 軸の組み合わせ → 遺伝子型
オープンウェイト ベース (Qwen、Llama、DeepSeek、Mistral) の上に大規模な言語モデルを構築することは、正当な業界標準の実践です。ただし、基礎モデルをゼロからトレーニングするのとは異なります。また、ベンダーは必ずしも区別を明確にしているわけではありません。 2026 年 7 月下旬にいくつかの研究所が DeepSeek に匹敵する「自社開発」モデル (例: LG K-EXAONE 2.0、750B) をリリースしたとき、この議論は中国の技術コミュニティにも飛び火し、Zhihu スレッド (→ リンク ) の閲覧数は 270 万回を超えました。自然な疑問が続きました。最初から作るのか、それとも派生させるのか?
これは、公開ファイルから客観的に回答可能です。その方法は次のとおりです。
2. 軸 1 — アーキテクチャのフィンガープリント ( config.json )
すべてのトランスフォーマー チェックポイントには config.json が同梱されています。いくつかのフィールドが驚くほど識別的なシグネチャを形成します。
num_attention_heads / num_key_value_heads
インポートリクエスト
def Arch_fingerprint (リポジトリ):
c = request.get( f"https://huggingface.co/ {repo} /resolve/main/config.json" ,
headers={ "ユーザーエージェント" : "ゲノム/1.0" }).json()
return {k: c.get(k) for k in
( "model_type" 、 "vocab_size" 、 "hidden_size" 、
"intermediate_size" 、 "num_hidden_layers" 、
"num_attention_heads" 、 "num_key_value_heads"

)}
形状タプル (hidden_size、intermediate_size、num_hidden_layers、heads、kv) は、事実上、リファレンス アーキテクチャの指紋となります。モデルのタプルが外国のオープンウェイトと正確に一致する場合、それはそのアーキテクチャが独自に設計されたのではなく採用されたという強力な証拠です。私たちが測定した例:
単一の偶然のフィールドには何の意味もありません。同時に5つが指紋です。
3. 軸 2 — トークナイザーの指紋 (親子鑑定)
アーキテクチャだけでは誤解を招く可能性があります。モデルは外国のアーキテクチャをコピーしても、真に新しいトークナイザーをトレーニングすることも、その逆も可能です。トークナイザーは tokenizer.json から直接測定され、語彙セットを最小重複率と比較します。
def vocab_set (リポジトリ):
j = request.get( f"https://huggingface.co/ {repo} /resolve/main/tokenizer.json" ).json()
v = j[ "モデル" ][ "語彙" ] # BPE: {トークン: id}
戻りセット (v.keys())
def tok_overlap ( a, b ):
A、B = vocab_set(a)、vocab_set(b)
return len (A & B) / min ( len (A), len (B)) # 1.0 == サブセット
これにより、設定で隠されているものがすぐに表面化します。 1 つのモデルは Qwen2.5-7B のアーキテクチャと正確に一致しましたが、そのトークナイザーは Qwen とわずか約 0.38 だけ重なりました。これは「外国の脳、独自の言語」のケースです。アーキテクチャは採用されましたが、新しい韓国語トークナイザーがトレーニングされました。逆に、一部の VLM は基本トークナイザーをそのまま再利用し (オーバーラップ = 1.000)、正確な微調整を確認しました。
実際的な罠: 分母 (和集合ではない) の min(|A|,|B|) は、より大きな語彙の厳密なサブセットである縮小語彙をスコア ~1.0 にするものです。これは、「ベースから切り出された」ことを示す正しい信号です。
4. 軸 3 — 重み付けフィンガープリント (ハードなもの)
最も重要な質問は、重みがゼロからトレーニングされたのか、それとも外国のベースで継続的に事前トレーニングされたのかということです。ここには 2 つの有益な罠が潜んでいます。
トラップ 1 — 行方向のコサインは u

セレス
単純なアイデア: 両方のモデルから embed_tokens.weight をロードし、共有トークンの行方向のコサイン類似度を平均します。系統を共有している場合、埋め込みも同様になるはずです。
たとえ彼らが明らかに血統を共有しているとしても、彼らはそうではありません。既知の最初から作成したモデルと既知の Llama 派生モデルの両方について、ゼロに近い平均コサインを測定しました。その理由は回転の不変性です。Transformer の隠れ空間には特権的な基盤がないため、2 つのモデルは任意の直交回転の下で同一の情報をエンコードできます。行方向のコサインは、回転を非類似性と見なします。血統を区別することはできません。
トラップ 2 — CKA は役立ちますが、十分ではありません
線形 CKA (Centered Kernel Alignment) は回転および等方性スケール不変であるため、表現を比較するのに最適なツールです。
輸入トーチ
def Linear_cka ( X, Y ):
# X: (n, d1)、Y: (n, d2) — 同じトークンの順序 (共有語彙)
X = X - X.mean( 0 , keepdim= True )
Y = Y - Y.mean( 0 , keepdim= True )
num = (X.T @ Y).norm() ** 2
den = (X.T @ X).norm() * (Y.T @ Y).norm()
return (num / den).item()
最初から作成したモデルは、候補ベースに対してほぼゼロの CKA スコアを獲得しました。これは、独立した事前トレーニングの明らかな証拠です。しかし、事前学習を継続した導関数のスコアはわずかに高く (≈0.25)、同じファミリーの 2 つの無関係なモデル間のベースライン (≈0.21) をかろうじて上回っています。大規模なトレーニングにより埋め込みが再形成されるため、CKA は導関数側での識別力を失います。
正直に述べた結論: 重み軸は最初から (ゼロに近い) ことを確実に確認しますが、導出を強力に検出するものではありません。そのため、config + tokenizer のフィンガープリントがプライマリのままになります。私たちは重み軸をそれ自体の評決としてではなく、裏付けとなる証拠として報告します。
5. ボーナス軸 — 独創性の代理としての多様性への注目
ほとんどのモデルは単一のアッテを宣言します

ションメカニズム。いくつかをいくつか混ぜます。 config.json 内の個別のメカニズムの数は、アーキテクチャの独自性を安価に代用します。
KEYS = ( "layer_types" , "linear_attn_config" , "sliding_window" ,
"mamba2_d_state" 、 "hyena_filter_order" 、 "mla_kv_lora_rank" 、
"attention_cls" )
デフォルトのattention_diversity ( cfg ):
ヒット = [CFG に k がある場合、KEYS に k の k]
# 例: Layer_types = [フル×16、スライディング×48] -> ハイブリッド (2)
ヒットを返す
私たちの調査では、ほとんどの韓国モデルは単一のグループ化されたクエリまたはマルチヘッドの潜在的な注意を使用していました。カップルはハイブリッド (layer_types = [full_attention×16, slide_attention×48]) を使用しました。そして最も多様性に富んだ、mamba2、ハイエナ、MLA、リニア アテンション、ゲート デルタ ネット、ネイティブ スパース アテンション、スライディング ウィンドウを 1 つのスタックに組み合わせたものです。
6. 軸の組み合わせ → 遺伝子型
2 つの主な軸 (アーキテクチャ × 重み) を 1 つのラベルにまとめます。
トークナイザーの重複と注目の多様性は評決に組み込まれずに併記されるため、読者は生の証拠を見ることができます。
同一のパイプラインを韓国の 9 つの組織 (大企業、通信会社、中堅企業、新興企業) の公的財団モデルに適用すると、状況は均一ではありません。一部のモデルは外国のアーキテクチャとトークナイザーに正確に一致します (移植)。他のものは、外部一致のない自己構築のアーキテクチャと重みを使用します (ネイティブ)。多くの人はその間に座ります。 3D リネージ グラフ、検索、ライト/ダーク モードを含むモデルごとの内訳は、「スペース」にあります。
告発ではありません。無差別級ベースに基づいて構築することは合法であり、広く普及しています。このツールは不正行為ではなく、血統を報告します。
重み軸は支持的なものであり、決定的なものではありません (セクション 4)。
例外なく、すべてのモデルで同じ基準。
すべての入力はパブリックです。修正は大歓迎です。
上記の 3 つの関数がメソッド全体です。 t 上の任意の 2 つのリポジトリをポイントします。

ハブ:
print (arch_fingerprint( "some/model" ))
print (tok_overlap( "some/model" , "Qwen/Qwen3-14B" ))
#weights: 共有語彙ペアの embed_tokens.weight をロードしてから、linear_cka
ライブ デモ、完全なデータセット、および 3 言語 UI: Model Genome Korea 。
モデル名、会社名、ライセンスはそれぞれの所有者の財産です。
この記事で取り上げたスペース 1
33 Model Genome Korea 🧬 韓国のLLMおよびVLM基盤モデルの33 DNA系統検査
画像、音声、ビデオをアップロードするには、テキスト入力をドラッグするか、貼り付けるか、ここをクリックします。ここをタップまたは貼り付けて画像をアップロード コメント · コメントするにはサインアップまたはログインしてください
この記事で取り上げたスペース 1
33 Model Genome Korea 🧬 韓国のLLMおよびVLM基盤モデルの33 DNA系統検査

## Original Extract

A Blog post by Proto_AGI on Hugging Face

Model Genome: Fingerprinting Whether an LLM Was Trained From Scratch or Derived
Hugging Face Models
Back to Articles a]:hidden">
Model Genome: Fingerprinting Whether an LLM Was Trained From Scratch or Derived
Community Article Published
August 8, 2026 Upvote 18
config.json</code>)"> 2. Axis 1 — Architecture fingerprint ( config.json )
3. Axis 2 — Tokenizer fingerprint (a paternity test)
4. Axis 3 — Weights fingerprint (the hard one) Trap 1 — row-wise cosine is useless
Trap 2 — CKA helps, but not enough
5. Bonus axis — attention diversity as an originality proxy
6. Combining axes → the genotype
Building a large language model on top of an open-weight base (Qwen, Llama, DeepSeek, Mistral) is a legitimate, industry-standard practice. But it is different from training a foundation model from scratch — and vendors do not always make the distinction explicit. When several labs released DeepSeek-rivaling "self-developed" models in late July 2026 (e.g. LG K-EXAONE 2.0, 750B), the debate spilled into Chinese tech communities as well — a Zhihu thread ( → link ) crossed 2.7M views. The natural question followed: from scratch, or derived?
This is answerable, objectively, from public files. Here is how.
2. Axis 1 — Architecture fingerprint ( config.json )
Every transformers checkpoint ships a config.json . A handful of fields form a surprisingly discriminative signature:
num_attention_heads / num_key_value_heads
import requests
def arch_fingerprint ( repo ):
c = requests.get( f"https://huggingface.co/ {repo} /resolve/main/config.json" ,
headers={ "User-Agent" : "genome/1.0" }).json()
return {k: c.get(k) for k in
( "model_type" , "vocab_size" , "hidden_size" ,
"intermediate_size" , "num_hidden_layers" ,
"num_attention_heads" , "num_key_value_heads" )}
The shape tuple (hidden_size, intermediate_size, num_hidden_layers, heads, kv) is effectively a fingerprint of the reference architecture. When a model's tuple matches a foreign open-weight exactly , that is strong evidence the architecture was adopted rather than designed independently. Examples we measured:
A single coincidental field means nothing; five simultaneously is a fingerprint.
3. Axis 2 — Tokenizer fingerprint (a paternity test)
Architecture alone can mislead. A model can copy a foreign architecture but train a genuinely new tokenizer, or vice-versa. The tokenizer is measured directly from tokenizer.json , comparing the vocabulary sets with a min-overlap ratio:
def vocab_set ( repo ):
j = requests.get( f"https://huggingface.co/ {repo} /resolve/main/tokenizer.json" ).json()
v = j[ "model" ][ "vocab" ] # BPE: {token: id}
return set (v.keys())
def tok_overlap ( a, b ):
A, B = vocab_set(a), vocab_set(b)
return len (A & B) / min ( len (A), len (B)) # 1.0 == subset
This immediately surfaces things config hides. One model matched Qwen2.5-7B's architecture exactly , yet its tokenizer overlapped Qwen by only ~0.38 — a "foreign brain, own language" case: the architecture was adopted, but a new Korean tokenizer was trained. Conversely, some VLMs reused a base tokenizer verbatim (overlap = 1.000), confirming a straight fine-tune.
A practical trap: min(|A|,|B|) in the denominator (not the union) is what makes a reduced vocabulary that is a strict subset of a larger one score ~1.0 — the correct signal for "carved out of the base."
4. Axis 3 — Weights fingerprint (the hard one)
The gold-standard question is: were the weights trained from scratch, or continued-pretrained on a foreign base? This is where two instructive traps live.
Trap 1 — row-wise cosine is useless
The naive idea: load embed_tokens.weight from both models, and for shared tokens, average the row-wise cosine similarity. If they share lineage, embeddings should be similar.
They are not — even when they obviously share lineage . We measured near-zero mean cosine for both a known from-scratch model and a known Llama-derivative. The reason is rotational invariance : a Transformer's hidden space has no privileged basis, so two models can encode identical information under an arbitrary orthogonal rotation. Row-wise cosine sees rotation as dissimilarity. It cannot distinguish lineage.
Trap 2 — CKA helps, but not enough
Linear CKA (Centered Kernel Alignment) is rotation- and isotropic-scale-invariant, so it is the right tool for comparing representations:
import torch
def linear_cka ( X, Y ):
# X: (n, d1), Y: (n, d2) — SAME token order (shared vocab)
X = X - X.mean( 0 , keepdim= True )
Y = Y - Y.mean( 0 , keepdim= True )
num = (X.T @ Y).norm() ** 2
den = (X.T @ X).norm() * (Y.T @ Y).norm()
return (num / den).item()
A from-scratch model scored near-zero CKA against its candidate base — clean evidence of independent pretraining. But a continued-pretrained derivative scored only modestly higher (≈0.25) — barely above the baseline between two unrelated models of the same family (≈0.21). Large-scale training reshapes embeddings enough that CKA loses discriminative power on the derivative side.
Conclusion, stated honestly: the weights axis reliably confirms from-scratch (near-zero), but it is not a strong detector of derivation . For that, config + tokenizer fingerprints remain primary. We report the weights axis as supporting evidence, not as a verdict on its own.
5. Bonus axis — attention diversity as an originality proxy
Most models declare a single attention mechanism. A few mix several. The count of distinct mechanisms in config.json is a cheap proxy for architectural originality:
KEYS = ( "layer_types" , "linear_attn_config" , "sliding_window" ,
"mamba2_d_state" , "hyena_filter_order" , "mla_kv_lora_rank" ,
"attention_cls" )
def attention_diversity ( cfg ):
hits = [k for k in KEYS if k in cfg]
# e.g. layer_types = [full×16, sliding×48] -> hybrid (2)
return hits
In our sweep, most Korean models used a single grouped-query or multi-head-latent attention; a couple used a hybrid ( layer_types = [full_attention×16, sliding_attention×48] ); and the most diverse combined mamba2, hyena, MLA, linear attention, gated-delta-net, native-sparse-attention and sliding-window in one stack.
6. Combining axes → the genotype
We collapse the two primary axes (architecture × weights) into one label:
The tokenizer overlap and attention diversity are shown alongside, not folded into the verdict, so readers can see the raw evidence.
Applying the identical pipeline to the public foundation models of nine Korean organizations (large enterprises, telcos, mid-size firms, and startups), the picture is not uniform : some models match a foreign architecture and tokenizer exactly (Ported); others use self-built architectures and weights with no foreign match (Native); many sit in between. The per-model breakdown — with a 3D lineage graph, search, and light/dark mode — is in the Space .
Not an accusation. Building on open-weight bases is legitimate and widespread. The tool reports lineage, not wrongdoing.
Weights axis is supporting, not conclusive (Section 4).
Same yardstick for every model , without exception.
All inputs are public; corrections are welcome.
The three functions above are the whole method. Point them at any two repos on the Hub:
print (arch_fingerprint( "some/model" ))
print (tok_overlap( "some/model" , "Qwen/Qwen3-14B" ))
# weights: load embed_tokens.weight for a shared-vocab pair, then linear_cka
Live demo, full dataset, and 3-language UI: Model Genome Korea .
Model names, companies, and licenses are the property of their respective owners.
Spaces mentioned in this article 1
33 Model Genome Korea 🧬 33 DNA lineage test of Korean LLM & VLM foundation models
Upload images, audio, and videos by dragging in the text input, pasting, or clicking here . Tap or paste here to upload images Comment · Sign up or log in to comment
Spaces mentioned in this article 1
33 Model Genome Korea 🧬 33 DNA lineage test of Korean LLM & VLM foundation models
