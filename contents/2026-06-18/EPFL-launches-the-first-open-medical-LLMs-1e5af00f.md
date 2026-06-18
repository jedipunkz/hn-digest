---
source: "https://actu.epfl.ch/news/epfl-launches-the-world-s-first-fully-open-medical/"
hn_url: "https://news.ycombinator.com/item?id=48580660"
title: "EPFL launches the first open medical LLMs"
article_title: "EPFL launches the world's first fully open medical LLMs - EPFL"
author: "ponsfrilus"
captured_at: "2026-06-18T04:34:35Z"
capture_tool: "hn-digest"
hn_id: 48580660
score: 1
comments: 0
posted_at: "2026-06-18T04:00:56Z"
tags:
  - hacker-news
  - translated
---

# EPFL launches the first open medical LLMs

- HN: [48580660](https://news.ycombinator.com/item?id=48580660)
- Source: [actu.epfl.ch](https://actu.epfl.ch/news/epfl-launches-the-world-s-first-fully-open-medical/)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T04:00:56Z

## Translation

タイトル: EPFL が初のオープン医療 LLM を開始
記事のタイトル: EPFL が世界初の完全にオープンな医療 LLM を開始 - EPFL
説明: MeditronFO は、医療における AI の透明性と説明責任を高めるために、医療用の大規模言語モデルを構築するための初の完全にオープンなフレームワークです。

記事本文:
EPFL が世界初の完全にオープンな医療 LLM を開始
このニュースを にインポートしてもよろしいですか?
このニュースは購読者に送信されます。
MeditronFO は、医療における AI の透明性と説明責任を高めるために、医療用の大規模言語モデルを構築するための初の完全にオープンなフレームワークです。
Medical Large Language Model (LLM) は臨床現場で使用されることが増えています。たとえば、AI は緊急治療室の医師が診断にフラグを立てたり、意思決定をサポートしたりするのに役立ちます。問題は、これらのシステムのほとんどが独自のものであることです。そのトレーニング データ、設計の選択、および意思決定プロセスは視界から隠されており、独立したレビューが実質的に不可能になっています。
これに応えて、EPFL のコンピュータ・コミュニケーション科学部のインテリジェント・グローバル・ヘルス＆人道対応技術研究所 (LiGHT) の研究者らは、OLMo、EuroLLM、EPFL とチューリッヒ工科大学が開発したスイスのモデルである Apertus など、いくつかの完全にオープンな基本モデルを「医療化」するために使用したフレームワークである MeditronFO (Fully Open) をリリースしました。
「2023 年に初めてリリースされた Meditron をベースに構築されている MeditronFO は、オープンな大規模言語モデルの医療バージョンを作成するためのパイプラインです」と、LiGHT で Meditron を率いる博士課程の学生である Xavier Theimer-Lienhard 氏は説明します。 「トレーニングが検証できない臨床医を私たちは決して信頼しません。同じ基準が医療分野の AI にも適用されるべきです。MeditronFO は、モデルのトレーニングに使用されるデータからコード、トレーニング手順、評価方法に至るまで、開発のあらゆる段階を公開します。」
「オープン」として販売されている AI モデルの多くは部分的にしか透明ではありません。彼らはトレーニングされたモデル自体をリリースしますが、その作成に使用されたデータセット、データ処理メソッド、トレーニング パイプラインはリリースしません。これにより独立性が高まります

エンドトの監査は困難であり、医療 AI システムがどのようにして推奨事項に到達するかを臨床医、病院、規制当局が理解する能力が制限されています。
MeditronFO は、最終製品の消費者としてではなく、トレーニング データのキュレーションからモデル出力の検証、潜在的な安全性の懸念の強調に至るまで、プロセス全体を通じて貢献する参加者として、臨床医が参加して構築されました。
MOOVE (大規模オープン オンライン検証および評価) を通じて、臨床医は継続的な評価とモデルの改善に直接参加し、開発が現実の臨床実践に基づいたものであることを保証します。これには、トレーニング資料と検証されたモデル出力の監査が含まれます。開発プロセスには安全対策も含まれていました。
このフレームワークは、公的に利用可能な医療データセットと、医療検査、臨床ガイドライン、現実的な患者の症例から得られた臨床医がレビューした合成データを組み合わせたものです。すべてのデータセット、処理手順、トレーニング手順はオープンに文書化されています。研究者らは、46,000 を超える臨床診療ガイドラインから抽出された、専門家が厳選した独自の臨床データセットを組み合わせました。
「私たちの研究結果は、最終的に使用される環境を反映するデータと評価を用いて、臨床医と地域社会が積極的に関与することによって、競争力のある医療 AI モデルを構築できることを示しています。これにより、優先順位が必ずしも地域のニーズと一致するとは限らない外部の独自システムにのみ依存するのではなく、医療システムと地域社会がこれらのテクノロジーのより大きな所有権を維持できる道が生まれます」と医師で LiGHT ディレクターのメアリーアン ハートリー教授は述べています。
すべての MeditronFO モデルは、元のベース モデルを上回りました。最も強力な結果は Apertus-70B-MeditronFO から得られました。

d 健康診断のパフォーマンスは、基礎となるモデルより 6.6 パーセントポイント向上しました。
「私たちの研究結果は、完全にオープンな医療モデルが実現可能で競争力があることを示しています。透明性が前提条件であり、命が命に関わる医療においては、これが重要です」と Theimer-Lienhard 氏は述べています。
MeditronFO の立ち上げは、より広範な研究プログラムの継続における重要なマイルストーンです。同チームはスイスからタンザニアまでの複数の施設で臨床試験を準備しており、実際の医療現場で医師がAIをどのように活用しているかを評価している。これらの研究では、臨床医が AI によって生成された推奨事項に従うか拒否するか、またその決定が患者ケアにどのような影響を与えるかを調査します。
MED.USE と呼ばれるこの複数年にわたる試験プロジェクトは、AI が不必要な治療や介入を減らしながらどのように医療の質を向上させることができるかを理解することも目的としています。 「患者の転帰に基づいて現実世界のフィードバックを得ることが重要です」とハートリー氏は説明します。
MeditronFO の立ち上げは、透明性と説明責任だけでなく、データ主権や独自の AI プラットフォームへの依存が高まることへの懸念を中心とした、医療における AI の将来に関する広範な議論を反映しています。この結果は、完全にオープンなアプローチにより、透明性と競争力の両方を備えた医療 AI システムを生み出すことができることを示しています。
「問題は、AIが医療の一部になるかどうかではなく、すでにそうなっている。問題は、私たちがどのようなAIエコシステムを構築したいのかということである。私たちは、透明性、科学的精査、臨床医と患者からの有意義な参加が中心であり続けるべきであると信じている。MeditronFOは、オープン性とパフォーマンスが競合する優先事項である必要はなく、革新的かつ説明責任のある医療AIに向けた実行可能な道があることを示している」とハートリー氏は結論づけた。
開発者たちは

MeditronFO は、大規模なモデルのトレーニング、評価、オープン リリースを可能にするコンピューティング インフラストラクチャと資金を提供する Swiss AI Initiative (EPFL、ETH Zurich、CSCS の共同作業) によって支援されました。
このコンテンツはクリエイティブ コモンズ CC BY-SA 4.0 ライセンスに基づいて配布されています。作者名を明示し、コンテンツのその後の使用に制限を設けない限り、コンテンツに含まれるテキスト、ビデオ、画像を自由に複製できます。 CC BY-SAの表示がないイラストを転載する場合は、作者の承諾を得る必要があります。
ソーシャル ネットワークで EPFL の動向をフォローしてください
© 2026 EPFL、全著作権所有

## Original Extract

MeditronFO is the first fully open framework for building medical large language models, to make AI in healthcare more transparent and accountable.

EPFL launches the world's first fully open medical LLMs
Are you sure you want to import this news into ?
This news will be sent to its subscribers.
MeditronFO is the first fully open framework for building medical large language models, to make AI in healthcare more transparent and accountable.
Medical large language models (LLMs) are increasingly being used in clinical settings. For example, AI is helping doctors in emergency rooms to flag diagnoses or support decisions. The problem is that most of these systems are proprietary: their training data, design choices, and decision-making processes are hidden from view, making independent review virtually impossible.
In response, researchers from EPFL’s Laboratory for Intelligent Global Health & Humanitarian Response Technologies (LiGHT) in the School of Computer and Communication Sciences have released MeditronFO (Fully Open), a framework they used to “medicalize” several fully open base models, including OLMo, EuroLLM and Apertus , Switzerland’s model developed by EPFL and ETH Zurich.
“Building on Meditron that was first released in 2023 , MeditronFO, is a pipeline to create a medical version of any open large language model,” explains Xavier Theimer-Lienhard, a PhD student leading Meditron at LiGHT. “We would never trust a clinician whose training can’t be verified, and the same standard should apply to AI in healthcare. MeditronFO makes every stage of development publicly available, from the data used to train the models to the code, training procedures, and evaluation methods.”
Many AI models marketed as “open” are only partially transparent. They release the trained model itself, but not the datasets, data-processing methods or training pipelines used to create it. This makes independent auditing difficult and limits the ability of clinicians, hospitals and regulators to understand how medical AI systems arrive at their recommendations.
MeditronFO was built with clinicians in the room, not as consumers of the end-product, but as participant contributing throughout the process, from curating training data to validating model outputs and highlighting potential safety concerns.
Through MOOVE (Massive Open Online Validation and Evaluations), clinicians participate directly in the ongoing evaluation, and improvement of models, helping ensure that development remains grounded in real-world clinical practice. This includes the auditing of training material and validated model outputs. The development process also included safeguards.
The framework combines publicly available medical datasets with clinician-reviewed synthetic data derived from medical examinations, clinical guidelines and realistic patient cases. All datasets, processing steps and training procedures are openly documented. The researchers combined a unique set of expert-curated clinical datasets drawn from more than 46,000 clinical practice guidelines.
“Our findings show that competitive medical AI models can be built through the active involvement of clinicians and communities with data and evaluations that reflect the settings where they will ultimately be used. This creates a pathway for health systems and communities to retain greater ownership of these technologies, rather than relying solely on external proprietary systems whose priorities may not always align with local needs,” says Professor Mary-Anne Hartley, medical doctor and Director of LiGHT.
Every MeditronFO model outperformed its original base model. The strongest results came from Apertus-70B-MeditronFO, which improved performance on medical exams by 6.6 percentage points over the underlying model.
“Our results have shown that fully open medical models are achievable and competitive. In medicine, where transparency is a prerequisite, and where the stakes are people’s lives, this matters,” says Theimer-Lienhard.
The launch of MeditronFO is an important milestone in the continuation of a broader research program. The team is preparing clinical trials in multiple sites from Switzerland to Tanzania that will evaluate how doctors use AI in real healthcare settings. These studies will examine whether clinicians follow or reject AI-generated recommendations and how those decisions affect patient care.
This multi-year trial project, called MED.USE , also aims to understand how AI can improve healthcare quality while reducing unnecessary treatments and interventions. “It is important to get real-world feedback based on patient outcomes,” explains Hartley.
The launch of MeditronFO reflects a broader debate over the future of AI in medicine, centered around transparency and accountability but also data sovereignty and fears of a growing dependence on proprietary AI platforms. The results demonstrate that fully open approaches can produce medical AI systems that are both transparent and highly competitive.
“The question is not whether AI will become part of healthcare, it already is. The question is what kind of AI ecosystem we want to build. We believe that transparency, scientific scrutiny, and meaningful participation from clinicians and patients should remain central. MeditronFO shows that openness and performance do not have to be competing priorities, and that there is a viable path toward medical AI that is both innovative and accountable,” Hartley concludes.
The development of MeditronFO was supported by the Swiss AI Initiative (a collaborative effort between EPFL, ETH Zurich, and CSCS) which provided computing infrastructure and funding to enable large-scale model training, evaluation, and open release.
This content is distributed under a Creative Commons CC BY-SA 4.0 license. You may freely reproduce the text, videos and images it contains, provided that you indicate the author’s name and place no restrictions on the subsequent use of the content. If you would like to reproduce an illustration that does not contain the CC BY-SA notice, you must obtain approval from the author.
Follow the pulses of EPFL on social networks
© 2026 EPFL, all rights reserved
