---
source: "https://transformer-circuits.pub/2026/nla/"
hn_url: "https://news.ycombinator.com/item?id=48540557"
title: "Natural Language Autoencoders Produce Explanations of LLM Activations"
article_title: "Natural Language Autoencoders Produce Unsupervised Explanations of LLM Activations"
author: "Anon84"
captured_at: "2026-06-15T14:15:25Z"
capture_tool: "hn-digest"
hn_id: 48540557
score: 2
comments: 0
posted_at: "2026-06-15T12:55:16Z"
tags:
  - hacker-news
  - translated
---

# Natural Language Autoencoders Produce Explanations of LLM Activations

- HN: [48540557](https://news.ycombinator.com/item?id=48540557)
- Source: [transformer-circuits.pub](https://transformer-circuits.pub/2026/nla/)
- Score: 2
- Comments: 0
- Posted: 2026-06-15T12:55:16Z

## Translation

タイトル: 自然言語オートエンコーダーが LLM アクティベーションの説明を生成する
記事のタイトル: 自然言語オートエンコーダーが LLM アクティベーションの教師なし説明を生成する

記事本文:
変圧器回路のスレッド
自然言語オートエンコーダーは LLM アクティベーションの教師なし説明を生成します
自然言語オートエンコーダーは LLM アクティベーションの教師なし説明を生成します
LLM アクティベーションの自然言語説明を生成する教師なし手法である Natural Language Autoencoders (NLA) を紹介します。 NLA は 2 つの LLM モジュールで構成されます。アクティベーションをテキストの説明にマッピングするアクティベーション バーバライザ (AV) と、説明をアクティベーションにマッピングし直すアクティベーション再構築器 (AR) です。 AV と AR を強化学習で共同トレーニングし、残差ストリームの活性化を再構築します。活性化の再構築を最適化していますが、結果として得られる NLA の説明はモデル内部のもっともらしい解釈として読み取られ、定量的評価によると、トレーニングを重ねるごとに情報量が増加します。
NLA をモデル監査に適用します。 Claude Opus 4.6 の導入前監査中に、NLA は安全関連の行動の診断を支援し、言語化されていない評価の認識を表面化しました。これは、クロードが評価されていると信じていたものの、発言していなかったケースです。当社はこれらの監査結果をケーススタディとして提示し、独立した方法を使用して検証します。意図的にずれたモデルのエンドツーエンドの調査を必要とする自動監査ベンチマークでは、NLA を備えたエージェントはベースラインを上回り、ずれたモデルのトレーニング データにアクセスしなくても成功できます。
NLA は、直接読むことができる表現力豊かな自然言語の説明を備えた、解釈しやすい便利なインターフェイスを提供します。さらなる作業をサポートするために、人気のあるオープン モデルのトレーニング コードとトレーニング済み NLA をリリースします。
言語モデルは、その内部状態を高次元の活性化ベクトルとしてエンコードします。これらのアクティベーションは豊富な情報を表します

モデルの計算についてはわかりますが、生の数値のリストであるため、人間の読者には不透明です。これらのベクトルを自然言語に翻訳するツールは、モデルの内部状態を直接判読できるようにします。このような翻訳を生成するための方法である自然言語オートエンコーダー (NLA) を紹介します。ターゲット LLM からのアクティベーションが与えられると、NLA は人間が容易に読めるテキスト記述を生成します。
NLA は、アクティベーション バーバライザー (AV) とアクティベーション コンストラクター (AR) の 2 つの LLM モジュールで構成されます。 AV はアクティベーションをテキストの説明 (説明 と呼ぶ) にマッピングします。 AR は、これらの説明をアクティベーションに変換します。 AV と AR はターゲット LLM のコピーとして初期化され、自然言語ボトルネックを通じてターゲットのアクティベーションを再構築するオートエンコーダーを一緒に形成します。強化学習 (RL) を使用して AV と AR を共同最適化し、再構成損失を最小限に抑えます。
このトレーニング プロセスは、NLA の説明が解釈可能または忠実であることを明示的に奨励するものではありません。それにもかかわらず、モデルのアクティベーションにエンコードされた情報が明らかになったことがわかりました。一連の定量的評価では、3 つのモデル (Claude Haiku 3.5、Haiku 4.5、Opus 4.6) にわたるトレーニングを通じて、NLA の説明がさらに有益になります。たとえば、NLA は、性別や職業など、コンテキストから推測可能なユーザーの特性を表面化します。
NLA の説明が実際にどのようなものであるかについての直感を構築するために、一連のケース スタディを紹介します。その多くは、Claude Opus 4.6 の導入前の整合性評価から抽出されました。これらのケーススタディは、他の解釈可能性手法、プロンプトバリエーション、トレーニングデータ検査などの独立した方法を使用して検証されます。たとえば、Opus 4.6 の初期バージョンがロシアで自発的に反応したとき

英語のユーザープロンプトに対して、NLA の説明は、ユーザーが「実際には」ロシア語を第一言語とする非英語話者であるという仮説にモデルが固執していることを示唆しました。これにより、英語に翻訳されたプロンプトと翻訳されていない外国語の応答を組み合わせた不正なトレーニング データが発見されました。また、NLA ベースの因果的介入を使用して仮説を検証することもできます。説明を編集し、元のバージョンと編集されたバージョンの両方を AR で再構築することで、それらの違いからステアリング ベクトルを取得します。おもちゃの設定では、このステアリング ベクトルを使用して、どの応答が報われるかについての Opus 4.6 の認識を変更します。これらのケーススタディからの抜粋を以下に示します。
これらのケーススタディは、NLA の重要な制限である作話も示しています。 NLA の説明には、ターゲット モデルの入力コンテキストに関する、検証可能な誤りである主張が含まれる可能性があります。これらの作話を分析すると、NLA の主張はその詳細においては誤りである可能性があるものの、通常、主題的には文脈に忠実であることがわかります。たとえば、歴史上の王朝に言及したテキストが与えられた場合、NLA の説明では、その王朝の特定の王が言及されていると誤って主張する可能性があります。複数の隣接するトークンにわたる説明に現れる主張も、真実である可能性が高くなります。これらのヒューリスティックは、NLA を実際にどのように使用するかを示します。
NLA の主な応用は、モデルの認知の言語化されていない側面、特に言語化されていない評価の認識、つまり明示的にはテストされていないのにモデルが内部的に疑っていることを明らかにすることです。モデルの根底にある信念は観察できないため、評価意識の NLA 測定をグラウンド トゥルースに照らして検証することはできません。ただし、ターゲット モデルが多かれ少なかれ評価であると信じる可能性が高い入力を構築し、それがどのようなものであるかをテストすることはできます。

彼女の NLA ベースの測定値は予想どおり変化します。たとえば、NLA は、実際の展開トラフィックよりも評価トランスクリプトで評価の意識を表面化することが多く、入力をより現実的なものに書き換えると、NLA で測定される評価の意識が低下することがわかりました。したがって、NLA は言語化されていない評価意識を特定できると結論付けます。
最後に、NLA が以前の研究から適応された下流の監査タスクで役立つことを検証します。私たちは監査エージェントに、Marks et al. で導入された意図的にずれたモデルを調査するよう依頼します。そして、NLA を備えた監査エージェントは、ベースラインよりもモデルの不整合の根本原因をより効果的に発見できることがわかりました。注目すべきことに、これまでの研究とは異なり、これらのエージェントは、調査中または NLA のトレーニング中に、モデルの不整合を引き起こしたトレーニング データにアクセスしなくても成功できます。
制限。 NLA にはいくつかの制限があります。
談話: NLA の説明は洞察力に富むものですが、騒々しい場合もあります。たとえば、矛盾した情報やコンテキストに関する検証可能な虚偽の主張が含まれる場合があります。事実に基づく幻覚を特定するのは簡単ですが、モデル処理に関するより一般的な主張が正確であるか、それとも捏造されているかを判断するのは困難な場合があります。
機構的な接地の欠如: NLA は構造的にブラックボックスです。活性化のどの側面が説明の特定の要素を駆動したのかを判断することはできません。
過剰な表現力: AV は完全な言語モデルであるため、アクティベーションに保存されているものを超えて追加の推論を行う能力があります。
コスト: NLA トレーニングには 2 つの完全な言語モデルでの共同 RL が必要で、推論にはアクティベーションごとに数百のトークンを生成する必要があります。これにより、NLA を大規模に使用するとコストが高くなる可能性があります。
退化トレーニング

限界内の目的: 原則として、AV は、入力コンテキストをそのまま再現することによって、または AR が非常に表現力豊かであるため、AR が反転できる解釈できない (または一見解釈できるようにしか見えない) テキストを出力することによって、良好な再構成を達成できます。現在の NLA ではどちらも重大な問題ではないようであり、KL 正則化などの部分的な緩和策が存在しますが、NLA をさらに開発する際にこれらの病状が良性のままであるかどうかは不明です。
全体として、NLA は既存の解釈可能性技術を強力に補完します。 NLA は自然言語を出力するため、表現力が豊かで使いやすいです。 NLA は、仮説生成を可能にし、モデルが言語化していない安全関連の認識を表面化できるため、ワークフローの監査に特に適していることがわかりました。さらなる作業をサポートするために、Neuronpedia との協力を通じて、トレーニング コード、人気のオープン モデル用のトレーニング済み NLA、およびオープン モデル NLA からサンプリングするためのインタラクティブ フロントエンドをリリースします。
紙のロードマップ。以下に紹介します。
関連研究の調査。NLA を教師なし概念発見手法 (SAE など) と教師あり活性化言語化手法 (活性化オラクルなど) の間の橋渡しとして位置づけています。
NLA アーキテクチャとトレーニング手順の説明。
Claude Opus 4.6 に関する 4 つのケーススタディは、モデル認知を解釈するための NLA の価値を示し、説明を読むための直観を構築し、独立した方法で結果を裏付けます。
Planning in Poetry : NLA の説明をガイド付きで最初に読むことができます。
言語切り替え : NLA は仮説生成に特に役立ちます。
誤って報告されたツール呼び出し : NLA は高レベルの認知を捕捉しますが、具体的な幻覚を示します。
報酬に関する推論: NLA から派生したステアリング ベクトルは、Opus の報酬に対する認識を変える可能性があります。

環境。
トレーニング全体を通してNLAの説明品質を定量的に評価します。 Opus 4.6、Haiku 4.5、および Haiku 3.5 NLA のトレーニングによってパフォーマンスが一貫して向上する 5 つの予測タスクを紹介します。また、ステガノグラフィー、書き込み品質の低下、作話などの NLA 障害モードも定量化します。
作話の詳細な特徴付け。どの種類の NLA 主張がより信頼できるかをヒューリスティックで特定します。
NLA を使用して言語化されていない評価認識を検出できることを検証します。
自動監査ベンチマーク 。NLA を備えたエージェントを SAE およびブラックボックス ベースラインと比較します。
一般的な質問応答アクティベーションオラクルの初期化としての NLA の探索。
NLA トレーニングによって有益な説明が得られる理由、NLA と機構的手法との関係、作話、コスト、レイヤーの感度、言語化できないアクティベーション コンテンツの可能性などの制限についての説明。
今後の作業の方向性。これには、アクティベーション スペースと自然言語の間で読み書きを行う汎用アクティベーション言語モデルのスケッチが含まれます。
モデルのアクティベーションを解釈するための既存の方法では、教師なし検出または直接読み取り可能な自然言語出力が提供されます。 NLA は、再構築の目的からの教師なし発見と、自然言語のボトルネックからの可読性の両方を提供するように設計されています。
アクティベーションを解釈するための教師なしメソッド。ロジット レンズとその調整されたバリアントは、モデルの非埋め込み行列を通じて中間アクティベーションを投影し、語彙トークンにわたる分布を取得します。スパース オートエンコーダは、教師なし再構成損失を使用してトレーニングされ、アクティベーションを学習された辞書特徴​​のスパース線形結合に分解します。いずれにしても、

、固定語彙 (トークンまたは辞書の特徴) からのアトムの加重和として解釈を表現することに制限されています。さらに、SAE 機能には予測できないカバレッジ ギャップが存在する可能性があり、人間またはモデルがトップ アクティベーション サンプルを分析することによって、場合によっては困難な解釈ステップが必要になります。
アクティベーションの自然言語説明。対照的に、最近の研究の中には、フリーテキストでアクティベーションを記述するように言語モデルをトレーニングするものもあります。既製のモデルには、これに対するある程度の能力があります。Lindsey は、モデルが注入されたステアリング ベクトルの内容を報告できる場合があることを発見しました。およびガンデハリウンら。どちらも、アクティベーションをプロンプトテンプレートにパッチすることによって解釈を引き出します。しかし、監視付き微調整の方が大幅に効果的です。Pan et al. 、コスタレッリら。 、およびChoiら。ソース コンテキスト (システム プロンプトなど) から答えがわかっているアクティベーションに関する質問に答えるようにモデルをトレーニングします。カルボネンら。彼らはそのようなモデルをアクティベーションオラクル（AO）と呼び、コンテキスト再構築の目的（参照）に関する事前トレーニングによって下流のQAパフォーマンスが向上することを示しました。黄ら。学習された希薄な概念のボトルネックを経由してアクティベーションをルーティングし、QA の回答を機械的に粗くすることを強制します

[切り捨てられた]

## Original Extract

Transformer Circuits Thread
Natural Language Autoencoders Produce Unsupervised Explanations of LLM Activations
Natural Language Autoencoders Produce Unsupervised Explanations of LLM Activations
We introduce Natural Language Autoencoders (NLAs), an unsupervised method for generating natural language explanations of LLM activations. An NLA consists of two LLM modules: an activation verbalizer (AV) that maps an activation to a text description and an activation reconstructor (AR) that maps the description back to an activation. We jointly train the AV and AR with reinforcement learning to reconstruct residual stream activations. Although we optimize for activation reconstruction, the resulting NLA explanations read as plausible interpretations of model internals that, according to our quantitative evaluations, grow more informative over training.
We apply NLAs to model auditing. During our pre-deployment audit of Claude Opus 4.6, NLAs helped diagnose safety-relevant behaviors and surfaced unverbalized evaluation awareness—cases where Claude believed, but did not say, that it was being evaluated. We present these audit findings as case studies and corroborate them using independent methods. On an automated auditing benchmark requiring end-to-end investigation of an intentionally-misaligned model, NLA-equipped agents outperform baselines and can succeed even without access to the misaligned model’s training data.
NLAs offer a convenient interface for interpretability, with expressive natural language explanations that we can directly read. To support further work, we release training code and trained NLAs for popular open models.
Language models encode their internal state as high-dimensional activation vectors. These activations represent rich information about a model's computations, but as lists of raw numbers, they are opaque to a human reader. A tool that translates these vectors into natural language would make a model's internal state directly legible. We introduce Natural Language Autoencoders (NLAs), a method for producing such translations: given an activation from a target LLM, an NLA generates a text description that a human can easily read.
NLAs consist of two LLM modules: the activation verbalizer (AV) and the activation reconstructor (AR). The AV maps activations to text descriptions, which we call explanations . The AR converts these explanations back to activations. The AV and the AR are initialized as copies of the target LLM, and together form an autoencoder that reconstructs the target's activations through a natural language bottleneck. We jointly optimize the AV and AR to minimize reconstruction loss using reinforcement learning (RL).
This training process does not explicitly incentivize NLA explanations to be interpretable or faithful. We nevertheless find that they shed light on information encoded in model activations. On a suite of quantitative evaluations, NLA explanations grow more informative throughout training across three models (Claude Haiku 3.5, Haiku 4.5, and Opus 4.6). For instance, NLAs surface traits of a user, such as their gender or occupation, that are inferrable from context.
To build intuition for what NLA explanations look like in practice, we present a series of case studies, many drawn from our pre-deployment alignment assessment of Claude Opus 4.6. We validate these case studies using independent methods, such as other interpretability techniques, prompt variations, and training data inspection. For instance, when an early version of Opus 4.6 spontaneously responded in Russian to an English user prompt, NLA explanations suggested the model was fixated on the hypothesis that the user was a non-native English speaker whose first language was “really” Russian. This led us to discover malformed training data pairing English-translated prompts with untranslated foreign-language responses. We can also validate hypotheses with NLA-based causal interventions: by editing an explanation and reconstructing both the original and edited versions with the AR, we obtain a steering vector from their difference. In a toy setting, we use this steering vector to change Opus 4.6's perception of which responses will be rewarded. Excerpts from these case studies are shown below.
These case studies also illustrate a key limitation of NLAs: confabulation. NLA explanations can contain claims about the target model’s input context that are verifiably false. Analyzing these confabulations, we find that while NLA claims can be false in their specifics, they are typically thematically faithful to the context. For instance, given text referencing a historical dynasty, NLA explanations might falsely claim that a specific king from that dynasty is mentioned. Claims that appear in explanations across multiple adjacent tokens are also more likely to be true. These heuristics inform how we use NLAs in practice.
A key application of NLAs is surfacing unverbalized aspects of model cognition, especially unverbalized evaluation awareness—a model internally suspecting it is being tested without explicitly saying so. We cannot validate NLA measurements of evaluation awareness against ground truth, because the model’s underlying beliefs are unobservable. However, we can construct inputs which the target model is more or less likely to believe are evaluations and test whether NLA-based measurements vary as expected. For instance, we find that NLAs surface evaluation awareness more often on evaluation transcripts than real deployment traffic, and that rewriting inputs to be more realistic reduces NLA-measured evaluation awareness. We therefore conclude that NLAs can identify unverbalized evaluation awareness.
Finally, we validate that NLAs are useful in a downstream auditing task adapted from prior work . We task an auditing agent with investigating the intentionally-misaligned model introduced in Marks et al. , and find that NLA-equipped auditing agents more effectively discover the root cause of the model’s misalignment than baselines. Notably—and unlike prior work —these agents can succeed without access to the training data which induced the model’s misalignment, either during the investigation or while training the NLA.
Limitations. NLAs have several limitations:
Confabulation: While NLA explanations can be insightful, they can also be noisy. For instance, they sometimes include contradictory information or verifiably false claims about the context. While factual hallucinations are easy to identify, it can be challenging to determine whether more general claims about model processing are accurate or confabulated.
Lack of mechanistic grounding: NLAs are blackboxes by construction; we cannot determine which aspects of an activation drove a given component of an explanation.
Excessive expressivity: Because the AV is a full language model, it has the capacity to make additional inferences beyond what is stored in an activation.
Cost: NLA training requires joint RL on two full language models, and inference requires generating several hundred tokens per activation. This can make NLAs expensive to use at scale.
Degenerate training objective in the limit: In principle, the AV could achieve good reconstruction by reproducing the input context verbatim, or by outputting uninterpretable (or only seemingly interpretable) text that the AR is able to invert because the AR is so expressive. While neither appears to be a significant problem in current NLAs, and partial mitigations such as KL regularization exist, it is unclear whether these pathologies will remain benign as we develop NLAs further.
Overall, NLAs are a powerful complement to existing interpretability techniques. Because NLAs output natural language, they are expressive and easy to use. We find NLAs especially well-suited to auditing workflows, where they enable hypothesis generation and can surface safety-relevant cognition that models do not verbalize. To support further work, we release training code , trained NLAs for popular open models, and an interactive frontend to sample from open model NLAs via our collaboration with Neuronpedia.
Paper roadmap. Below, we present:
A survey of related work , positioning NLAs as a bridge between unsupervised concept-discovery methods (e.g., SAEs) and supervised activation-verbalization methods (e.g., activation oracles).
A description of the NLA architecture and training procedure.
Four case studies on Claude Opus 4.6 , which illustrate the value of NLAs for interpreting model cognition, build intuition for reading their explanations, and corroborate their findings with independent methods.
Planning in Poetry : A guided first read of NLA explanations.
Language Switching : NLAs are especially useful for hypothesis generation.
Misreported Tool Calls : NLAs capture high-level cognition but hallucinate specifics.
Reasoning about Rewards : NLA-derived steering vectors can alter Opus’s perception of its environment.
Quantitative evaluations of NLA explanation quality throughout training. We introduce five prediction tasks in which performance consistently improves with training for Opus 4.6, Haiku 4.5, and Haiku 3.5 NLAs. We also quantify NLA failure modes , like steganography, writing quality degradation, and confabulation.
An in-depth characterization of confabulations , identifying heuristics for which kinds of NLA claims are more trustworthy.
A validation that NLAs can be used to detect unverbalized evaluation awareness .
An automated auditing benchmark , comparing NLA-equipped agents against SAE and blackbox baselines.
An exploration of NLAs as an initialization for general question-answering activation oracles .
A discussion of why NLA training results in informative explanations, how NLAs relate to mechanistic methods, and limitations including confabulations, cost, layer sensitivity, and the possibility of unverbalizable activation content.
Directions for future work , including a sketch of general-purpose activation language models that read and write between activation space and natural language.
Existing methods for interpreting model activations offer either unsupervised discovery or directly readable, natural language output. NLAs are designed to provide both: unsupervised discovery from the reconstruction objective and readability from the natural-language bottleneck.
Unsupervised methods for interpreting activations. The logit lens and its tuned variants project intermediate activations through a model’s unembedding matrix to obtain a distribution over vocabulary tokens. Sparse autoencoders are trained with an unsupervised reconstruction loss to decompose activations into sparse linear combinations of learned dictionary features. In either case, though, we are limited to expressing interpretations as a weighted sum of atoms from a fixed vocabulary (tokens or dictionary features). Moreover, SAE features can have unpredictable coverage gaps , and require a sometimes-difficult interpretative step, either by a human or a model analyzing top-activating examples .
Natural language explanations of activations. By contrast, some recent work trains language models to describe activations in free text. Off-the-shelf models have some capacity for this: Lindsey finds that models can sometimes report the content of injected steering vectors, while Chen et al. and Ghandeharioun et al. both elicit interpretations by patching activations into a prompting template. But supervised fine-tuning is substantially more effective: Pan et al. , Costarelli et al. , and Choi et al. train models to answer questions about activations whose answers are known from the source context (e.g., a system prompt). Karvonen et al. call such models activation oracles (AOs), and show that pretraining on a context-reconstruction objective (cf. ) improves their downstream QA performance. Huang et al. route the activation through a learned sparse concept bottleneck, forcing QA answers to be mechanistically grou

[truncated]
