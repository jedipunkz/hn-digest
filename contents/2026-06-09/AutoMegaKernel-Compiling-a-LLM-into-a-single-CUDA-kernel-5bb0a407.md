---
source: "https://arxiv.org/abs/2606.09682"
hn_url: "https://news.ycombinator.com/item?id=48467210"
title: "AutoMegaKernel: Compiling a LLM into a single CUDA kernel"
article_title: "[2606.09682] AutoMegaKernel: A Statically-Checked Agent Harness for Self-Retargeting Megakernel Synthesis"
author: "OsamaJaber"
captured_at: "2026-06-09T21:38:35Z"
capture_tool: "hn-digest"
hn_id: 48467210
score: 3
comments: 0
posted_at: "2026-06-09T20:26:21Z"
tags:
  - hacker-news
  - translated
---

# AutoMegaKernel: Compiling a LLM into a single CUDA kernel

- HN: [48467210](https://news.ycombinator.com/item?id=48467210)
- Source: [arxiv.org](https://arxiv.org/abs/2606.09682)
- Score: 3
- Comments: 0
- Posted: 2026-06-09T20:26:21Z

## Translation

タイトル: AutoMegaKernel: LLM を単一の CUDA カーネルにコンパイルする
記事タイトル: [2606.09682] AutoMegaKernel: セルフリターゲティングメガカーネル合成用の静的にチェックされたエージェントハーネス
説明: arXiv 論文 2606.09682 の要約ページ: AutoMegaKernel: セルフリターゲット メガカーネル合成のための静的にチェックされたエージェント ハーネス

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2606.09682
ヘルプ |高度な検索
コンピューターサイエンス > 機械学習
[2026 年 6 月 8 日に提出]
タイトル: AutoMegaKernel: セルフリターゲティングメガカーネル合成のための静的にチェックされたエージェントハーネス
要約: AutoMegaKernel (AMK) は、HuggingFace Llama ファミリ モデルを、モデルごとに手書きの CUDA を使用せずに、1 回の起動でフォワード パス全体を実行する単一の永続的な協調 CUDA カーネルにコンパイルします。貢献するのはシステムであり、実際の速度ではありません。
凍結されたスケジュール IR バリデータは、静的なグラフ チェック (機械化された証明ではありません) を介してデッドロック フリーとレース フリーを静的に証明するため、エージェントが提案した安全でないスケジュールは起動前に拒否されます。7,160 の敵対的スケジュール (6,091 が安全ではありません) 全体で、誤認はゼロで、360 の実際の低下はすべて受け入れられました。同じソースは 1 つのコードベースから sm_80/sm_90/sm_120 を再ターゲットし、サポートされている 10 モデルのうち 10 モデルに対して正しいメガカーネルを自動生成し、実際の SmolLM2-135M チェックポイントで HuggingFace 貪欲デコード トークンをトークンごとに再現します (パープレキシティ マッチ 2.5e-7)。無人のエージェント駆動可能な自動調査ループにより、メガカーネルが独自のベースライン (1.25 ～ 1.72x) を超えて自己改善されます。
検索で見つかった int8 (W8A16) メガカーネルは、NVIDIA のデータセンター推論フリート全体でバッチ 1 デコードで CUDA グラフ化された cuBLAS bf16 を上回りました: L4 は最大 1.33 倍、現行世代 L40S は 1.25 ～ 1.27 倍、A10G は最大 1.08 倍、コンシューマ RTX 5090 1.19～1.23倍。この順序は帯域幅の明確な関数ではありません (864 GB/秒の L40S が 600 GB/秒の A10G を上回ります)。違いは推論クラスとトレーニングクラスです。 AMK は、高帯域幅トレーニング クラス A100/H100 で cuBLAS に後れを取っています。

ハーネスは SM 間同期のボトルネックを特定します。私たちはそのギャップをはっきりと報告します。これは、デコード位置 0 での精度非対称 (W8A16 対 bf16) の比較です。実際の最大のチェックポイントは TinyLlama-1.1B です。コードとハーネス: この https URL
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。
arXiv に連絡する arXiv に連絡するにはここをクリックしてください
お問い合わせ
arXiv メールを購読する 購読するにはここをクリックしてください
購読する

## Original Extract

Abstract page for arXiv paper 2606.09682: AutoMegaKernel: A Statically-Checked Agent Harness for Self-Retargeting Megakernel Synthesis

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2606.09682
Help | Advanced Search
Computer Science > Machine Learning
[Submitted on 8 Jun 2026]
Title: AutoMegaKernel: A Statically-Checked Agent Harness for Self-Retargeting Megakernel Synthesis
Abstract: AutoMegaKernel (AMK) compiles a HuggingFace Llama-family model into a single persistent cooperative CUDA kernel that runs the whole forward pass in one launch, with no per-model hand-written CUDA. The contribution is the system, not raw speed.
A frozen schedule-IR validator statically certifies deadlock-freedom and race-freedom via static graph checks (not a mechanized proof), so an unsafe agent-proposed schedule is rejected before launch: across 7,160 adversarial schedules (6,091 unsafe) it had zero false-accepts and accepted all 360 real lowerings. The same source retargets sm_80/sm_90/sm_120 from one codebase, auto-generates correct megakernels for 10 of 10 supported models, and on a real SmolLM2-135M checkpoint reproduces HuggingFace greedy decode token-for-token (perplexity match 2.5e-7). An unattended, agent-drivable autoresearch loop self-improves the megakernel over its own baseline (1.25-1.72x).
A search-found int8 (W8A16) megakernel beats CUDA-graphed cuBLAS bf16 at batch-1 decode across NVIDIA's datacenter inference fleet: L4 up to 1.33x, the current-gen L40S 1.25-1.27x, A10G up to 1.08x at scale, and the consumer RTX 5090 1.19-1.23x. The ordering is not a clean function of bandwidth (the 864 GB/s L40S beats the 600 GB/s A10G); the divide is inference-class vs training-class. AMK trails cuBLAS on the high-bandwidth training-class A100/H100, where the harness localizes the cross-SM-sync bottleneck; we report the gap plainly. This is a precision-asymmetric (W8A16 vs bf16) comparison at decode position 0; the largest real checkpoint is TinyLlama-1.1B. Code and the harness: this https URL
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
contact arXiv Click here to contact arXiv
Contact
subscribe to arXiv mailings Click here to subscribe
Subscribe
