---
source: "https://pokee.ai/blog/pokee-ai-daily-2026-06-22"
hn_url: "https://news.ycombinator.com/item?id=48636012"
title: "Sakana AI Ships Fugu, an Orchestration Model Claiming Fable 5 Performance"
article_title: "Sakana AI Ships Fugu, an Orchestration Model That Claims Fable 5 Performance Without the Export Ban"
author: "polskibus"
captured_at: "2026-06-22T21:42:08Z"
capture_tool: "hn-digest"
hn_id: 48636012
score: 3
comments: 1
posted_at: "2026-06-22T20:52:01Z"
tags:
  - hacker-news
  - translated
---

# Sakana AI Ships Fugu, an Orchestration Model Claiming Fable 5 Performance

- HN: [48636012](https://news.ycombinator.com/item?id=48636012)
- Source: [pokee.ai](https://pokee.ai/blog/pokee-ai-daily-2026-06-22)
- Score: 3
- Comments: 1
- Posted: 2026-06-22T20:52:01Z

## Translation

タイトル: さかなAIがふぐを出荷、寓話5のパフォーマンスを主張するオーケストレーションモデル
記事タイトル：さかなAI、輸出禁止なしで「フェイブル5」のパフォーマンスを主張するオーケストレーションモデル「ふぐ」を出荷
説明: 東京を拠点とするSakana AIは、GPT-5.5、Gemini、およびOpusを調整するように訓練されたモデルが、Anthropicの最高のモデルがまだ世界中でオフラインになっているまさにその日にフロンティアレベルの結果を提供できることに賭けて、FuguとFugu Ultraを一般提供開始します。

記事本文:
Sayaka AI、輸出禁止なしで「Fable 5」のパフォーマンスを主張するオーケストレーション モデル Fugu を出荷 ブログデイリー ダイジェストに戻る Sakana AI、輸出禁止なしで「Fable 5」のパフォーマンスを主張するオーケストレーション モデル Fugu を出荷
さかなAIは本日、ふぐとふぐウルトラを一般販売商品として出荷いたしました。核となるアイデアは珍しいものです。Fugu 自体は、タスクの委任、エージェント間の通信の管理、および単一の統合 API を介したフロンティア モデル (現在は GPT-5.5、Gemini 3.5 Flash、および Claude Opus 4.8) のプール全体で結果を集約するために特別にトレーニングされた言語モデルです。より高いパフォーマンス層である Fugu Ultra は、標準的な推論とコーディングの評価において、Anthropic の Fable 5 および Mythos Preview と同等のベンチマークであると主張しています。このアーキテクチャは、学習済みモデル オーケストレーション (Trinity および The Conductor 論文) に関する Sakana の ICLR 2026 研究を形式化し、クロスモデル委任を手作業でコーディングされたルーターではなくトレーニング可能な目標として扱います。
タイミングが指摘されている。 Anthropic の Fable 5 と Mythos 5 は、米国商務省の輸出規制指令により 6 月 12 日以来世界中でアクセスできなくなっており、今日はこれらのモデルが使用量クレジット請求に移行する前に Anthropic の標準サブスクリプション プランに含まれていた最後の日です。 「どのようなフロンティア機能が存在するか」と「開発者が合法的にアクセスできるもの」との間のギャップがこれほど広いことはまれであり、Fugu はそれを直接ターゲットにしています。輸出制限の対象外のモデルをオーケストレーションすることで、地理に依存しないユーザー ベースに Fable 5 層の出力を提供します。当面の裁定取引を超えて、アーキテクチャ上の主張がより永続的な点です。つまり、中間層モデルをオーケストレーションすることで、単一のフロンティア モデルを置き換えることができ、コスト構造、プロバイダーの回復力、地政学的コンプライアンスにおいて複合的な利点が得られます。その論文

これは、単一のフロンティア API に依存できないオンプレミスまたはオンデバイスのインフラストラクチャを実行している組織に直接関係します。
追跡すべき 3 つのこと。まず、Fugu のベンチマーク数値が、エージェントおよび長期的なタスクでの独立したレプリケーションに耐えられるかどうかです。オーケストレーション システムは歴史的に、単一モデルの出力用に設計された eval で過剰なインデックスを作成し、多くのステップにわたって厳密な状態の一貫性を必要とするタスクに苦労しています。 2 つ目は、レイテンシーとコストです。プロバイダー間で API 呼び出しを連鎖させると、運用環境で問題となるラウンドトリップ オーバーヘッドが発生します。また、1 つのモデルの作業を実行するために 3 つのモデルに料金を支払う経済性については、大規模な場合に精査する必要があります。第三に、Anthropic のタイムライン: オーケストレーション アーキテクチャは 1 つのトリック プレイではありませんが、Fable 5 の復活は Fugu のポジショニングを大幅に圧縮するでしょう。ここでの Sakana の信頼性は本物であり、共同創設者の Llion Jones は、当初の Transformer 共著者 8 人のうちの 1 人であり、同研究所の ICLR 実績により、これは単なる Vaporware 以上のものとなっています。より広範な監視: Fugu Ultra の主張が正しければ、オーケストレーションを二流のフォールバックとして扱う時代は終わったことになります。
ナデラ氏、AI の集中は政治的監視に耐えられないと警告 : Microsoft CEO のサティア ナデラ氏は、「少数のモデルだけですべての価値が生み出されるなら、政治経済は単純にそれを容認しない」と主張する記事を発表し、AI の統合を 1990 年代のグローバリゼーションによる空洞化効果と比較し、企業に対しフロンティア API に依存するのではなく、プライベートな学習ループを構築するよう求めました。 [リンク]
SpaceX、大ヒットIPO後にCursorを600億ドルの株式で買収 : SpaceXはコールオプションを行使して、Cursor AIコーディングエディターのメーカーであるAnysphereを全株式評価額600億ドルという暗示で買収した：CursとのxAIのマスク自己取引を除けば、ベンチャーキャピタル支援のスタートアップ企業としては史上最大の買収となる

or の ARR は、2025 年初頭の 1 億ドルから 2026 年 6 月までに 40 億ドル以上に増加しました。[リンク]
Fable 5 の無料トライアルは本日終了します。 Anthropic、生体認証 ID 検証を開始 : 6 月 22 日が最終日 Fable 5 は、モデルが輸出規制下でオフラインのままであっても、追加費用なしで Anthropic の Pro、Max、および Team プランにバンドルされています。 Anthropic の更新されたプライバシー ポリシー (7 月 8 日発効) では、政府発行の ID と生体認証を収集する予定ですが、これは米国国民のみの復元の可能性が高いメカニズムです。 [リンク]
さかなふぐ：1つのモデルで全てを操る
日本のSakana AIが「Fuguシステム」をローンチ; GPT-5.5、Gemini、Opusを上回ると言われています
Saken AI の Fugu は、Anthropic の Fable および Mythos ベンチマークに一致するように複数の LLM を調整します
ワークフローを自動化する準備はできていますか?
500 の無料クレジット。 90 以上の統合。クレジットカードは必要ありません。
インフラストラクチャに導入されたフロンティア エージェント。
© 2026 ポキーAI。無断転載を禁じます。
CB Insights AI 100 の Pokee Pokee が CB Insights の第 10 回年次 AI 100 リストに選出 続きを読む → Pokee AI Enterprise Finance コンプライアンス、オンボーディング、ドキュメント レビュー
ヘルスケア 臨床ワークフロー、EHR、ケアコーディネーション
Eコマースカタログ、サポート、購入後の運用
教育 入学、成績評価、学生サポート
製造業 予知保全、サプライチェーン

## Original Extract

Tokyo-based Sakana AI launches Fugu and Fugu Ultra to general availability, betting that a model trained to coordinate GPT-5.5, Gemini, and Opus can deliver frontier-tier results on the exact day Anthropic's best model is still offline worldwide.

Sakana AI Ships Fugu, an Orchestration Model That Claims Fable 5 Performance Without the Export Ban Back to Blog Daily Digest Sakana AI Ships Fugu, an Orchestration Model That Claims Fable 5 Performance Without the Export Ban
Sakana AI today shipped Fugu and Fugu Ultra as generally available products. The core idea is unusual: Fugu is itself a language model trained specifically to delegate tasks, manage inter-agent communication, and aggregate results across a pool of frontier models (currently GPT-5.5, Gemini 3.5 Flash, and Claude Opus 4.8) through a single unified API. Fugu Ultra, the higher-performance tier, claims benchmark parity with Anthropic's Fable 5 and Mythos Preview on standard reasoning and coding evaluations. The architecture formalizes Sakana's ICLR 2026 research on learned model orchestration (the Trinity and The Conductor papers), treating cross-model delegation as a trainable objective rather than a hand-coded router.
The timing is pointed. Anthropic's Fable 5 and Mythos 5 have been inaccessible globally since June 12 under a US Department of Commerce export control directive, and today is the final day those models were included in Anthropic's standard subscription plans before moving to usage-credit billing. The gap between "what frontier capability exists" and "what developers can legally access" has rarely been this wide, and Fugu targets it directly: by orchestrating models not subject to export restrictions, it offers Fable 5-tier output to a geography-agnostic user base. Beyond the immediate arbitrage, the architectural claim is the more durable point: orchestrating mid-tier models can substitute for a single frontier model, with compounding advantages in cost structure, provider resilience, and geopolitical compliance. That thesis is directly relevant for any organization running on-prem or on-device infrastructure that cannot depend on a single frontier API.
Three things to track. First, whether Fugu's benchmark numbers survive independent replication on agentic and long-horizon tasks: orchestration systems historically over-index on evals designed for single-model outputs and struggle with tasks requiring tight state coherence across many steps. Second, latency and cost: chaining API calls across providers introduces round-trip overhead that matters in production, and the economics of paying for three models to do one model's work need scrutiny at scale. Third, Anthropic's timeline: Fable 5's return would sharply compress Fugu's positioning, though the orchestration architecture is not a one-trick play. Sakana's credibility here is real, co-founder Llion Jones is one of the eight original Transformer co-authors, and the lab's ICLR track record makes this more than vaporware. The broader watch: if Fugu Ultra's claims hold up, the era of treating orchestration as a second-class fallback is over.
Nadella Warns AI Concentration Will Not Survive Political Scrutiny : Microsoft CEO Satya Nadella published a piece arguing that 'if all the value is accrued by only a few models, the political economy will simply not tolerate it,' comparing AI consolidation to the hollowing-out effect of 1990s globalization and calling for firms to build private learning loops rather than depending on frontier APIs. [link]
SpaceX Acquires Cursor for $60B in Stock After Blockbuster IPO : SpaceX exercised a call option to acquire Anysphere, maker of the Cursor AI coding editor, at an implied $60 billion all-stock valuation: the largest-ever acquisition of a VC-backed startup outside of Musk self-dealing for xAI, with Cursor's ARR having grown from $100M in early 2025 to over $4B by June 2026. [link]
Fable 5 Free Trial Ends Today; Anthropic Rolls Out Biometric ID Verification : June 22 is the last day Fable 5 is bundled in Anthropic's Pro, Max, and Team plans at no extra cost, even as the model remains offline under export controls; Anthropic's updated privacy policy (effective July 8) will collect government-issued ID and biometrics, the likely mechanism for a US-citizens-only restoration. [link]
Sakana Fugu: One Model to Command Them All
Japan's Sakana AI launches 'Fugu system'; says it beats GPT-5.5, Gemini, Opus
Sakana AI's Fugu orchestrates multiple LLMs to match Anthropic's Fable and Mythos benchmarks
Ready to automate your workflow?
500 free credits. 90+ integrations. No credit card required.
Frontier Agent Deployed in Your Infrastructure.
© 2026 Pokee AI. All rights reserved.
Pokee in CB Insights AI 100 Pokee named to CB Insights' 10th Annual AI 100 List Read more → Pokee AI Enterprise Finance Compliance, onboarding, document review
Health Care Clinical workflows, EHR, care coordination
E-commerce Catalog, support, post-purchase ops
Education Admissions, grading, student support
Manufacturing Predictive maintenance, supply chain
