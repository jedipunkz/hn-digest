---
source: "https://jonready.com/blog/posts/claude-fable5-is-allowed-to-sabotage-your-app-if-youre-a-competitor.html"
hn_url: "https://news.ycombinator.com/item?id=48467896"
title: "If Claude Fable stops helping you, you'll never know"
article_title: "If Claude Fable stops helping you, you'll never know — Jonathon Ready"
author: "mips_avatar"
captured_at: "2026-06-09T21:37:47Z"
capture_tool: "hn-digest"
hn_id: 48467896
score: 4
comments: 2
posted_at: "2026-06-09T21:19:11Z"
tags:
  - hacker-news
  - translated
---

# If Claude Fable stops helping you, you'll never know

- HN: [48467896](https://news.ycombinator.com/item?id=48467896)
- Source: [jonready.com](https://jonready.com/blog/posts/claude-fable5-is-allowed-to-sabotage-your-app-if-youre-a-competitor.html)
- Score: 4
- Comments: 2
- Posted: 2026-06-09T21:19:11Z

## Translation

タイトル: クロード・ファブルがあなたを助けるのをやめたら、あなたは決して知りません
記事のタイトル: クロード・ファブルがあなたを助けるのをやめたら、あなたは決して知りません — Jonathon Ready
説明: 人類

記事本文:
← ブログに戻る
クロード・ファブルがあなたを助けるのをやめても、あなたは決して知りません
モデルカードでこれが読めるとは思いませんでした。
フェイブル5モデルカード：
フロンティア LLM 開発 (たとえば、事前トレーニング パイプライン、分散トレーニング インフラストラクチャ、または ML アクセラレータの設計の構築など) をターゲットとするリクエストに対するクロードの有効性を制限する新しい介入を実装しました。 Claude を使用して競合モデルを開発することは、すでに当社の利用規約に違反していますが、当社の安全策を通じてこの制限を強制することで、これらの規約に違反しようとする攻撃者が加速することを回避できます。サイバーセキュリティ、生物学と化学、蒸留の試みに対する当社の介入とは異なり、これらの保護手段はユーザーには見えません。 Fable 5 は、別のモデルにフォールバックしません。代わりに、安全策は、即時変更、ステアリングベクトル、パラメータ効率の良い微調整 (PEFT) などの方法を通じて有効性を制限します。
クロードは静かに弱体化できるようになりました。 Anthropic は、これがいつ起こるかをユーザーに知らせないと決定しました。
現代のソフトウェア会社は、独自の埋め込み、再ランキング、推奨システムを構築することが増えています。私の小さなブートストラップ アプリ、wanderfugl.com にも、私が自分でトレーニングしたカスタム リランカーと埋め込みアルゴリズムが含まれています。
Anthropic は、同社が考える「フロンティア AI 開発」の例をいくつか挙げていますが、明確な線引きは示していません。問題は、かつては AI ラボ専用とされていた技術の多くが、現在では一般のソフトウェア会社で使用されているということです。スタートアップは埋め込みモデルをトレーニングします。彼らはリランカーを構築します。彼らは微調整し、小規模な llms をホストします。 「フロンティアAI研究」と通常の製品開発との境界線は年々定義が難しくなっている。
これにより、企業にとってサプライチェーンに真のリスクが生じます。私が迷っている間にクロードが私に下手なアドバイスや間違ったアドバイスをしたら

AI コンポーネントを操作しているとき、モデルが混乱しているのか、問題が解決できないのか、それとも目に見えないポリシー制限が静かに発動しているのかを知る方法はありません。Anthropic は、これがいつ起こっているかをユーザーに通知しないことを明確に選択しています。
開発ツールがユーザーに通知することなく成功のための最適化を停止すると、インフラストラクチャを完全に信頼することは不可能になります。
人為的なサプライチェーンのリスク
Anthropic によれば、これらの安全対策は開発者の 0.03% にしか影響を与えないという。もしかしたら今日もそうなのかもしれない。
問題は、AI企業の定義が変わりつつあることだ。
もしかしたら、現在フロンティア モデルをトレーニングしていないかも知れません。ほとんどの企業はトレーニングしていません。しかし、最新のソフトウェアには AI モデルが含まれることが増えています。 5 年前、スタートアップを構築するということは、API と SQL クエリを作成することを意味していました。現在では、モデルのトレーニング、チューニング、デプロイを意味することが多くなっています。
5 年前、CLIP のようなモデルは最先端の AI 研究プロジェクトでした。今日、私は自社で旅行を始めるスタートアップ向けにそれらを微調整しています。
製品のモデル トレーニング パイプラインをデバッグしていて、クロードが間違った答えを返した場合、モデルは混乱していませんか?間違った文脈を与えましたか?それとも、隠されたポリシーによってクロードの支援能力が弱まったのでしょうか?

## Original Extract

Anthropic

← Back to blog
If Claude Fable stops helping you, you'll never know
I didn't expect to read this in a model card.
Fable 5 model card :
we’ve implemented new interventions that limit Claude’s effectiveness for requests targeting frontier LLM development (for example, on building pretraining pipelines, distributed training infrastructure, or ML accelerator design). Using Claude to develop competing models already violates our Terms of Service, but enforcing this restriction through our safeguards avoids accelerating the actors most willing to violate these terms. Unlike our interventions for cybersecurity, biology and chemistry, and distillation attempts, these safeguards will not be visible to the user. Fable 5 will not fall back to a different model. Instead, the safeguards will limit effectiveness through methods such as prompt modification, steering vectors, or parameter-efficient fine-tuning (PEFT).
Claude can now be silently nerfed. Anthropic has decided it won't tell users when this happens.
Modern software companies increasingly build their own embedding, reranking, and recommendation systems. Even my small bootstrapped app, wanderfugl.com , has a custom reranker and embedding algorithm that I trained myself.
Anthropic gives a few examples of what it considers "frontier AI development," but doesn’t provide a clear line. The problem is that many techniques once reserved for AI labs are now being used by ordinary software companies. Startups train embedding models. They build rerankers. They finetune and host small llms. The boundary between "frontier AI research" and normal product development is becoming harder to define every year.
That creates a real supply chain risk for businesses. If Claude gives me poor or incorrect advice while I’m working on an AI component, I have no way of knowing whether the model was confused, whether my problem is unsolvable, or if some invisible policy restriction quietly kicked in. Anthropic has explicitly chosen not to tell users when this is happening.
Once a development tool can stop optimizing for your success without telling you, it becomes impossible to fully trust your infrastructure.
The Anthropic supply chain risk
Anthropic says these safeguards only affect 0.03% of developers. Maybe that's true today.
The problem is that the definition of an AI company is changing.
Maybe you're not training frontier models today—most companies aren't. But modern software increasingly contains AI models. Five years ago, building a startup meant writing APIs and SQL queries. Today, it often means training, tuning, and deploying models.
Five years ago, models like CLIP were frontier AI research projects. Today I'm fine-tuning them for a bootstrapped travel startup.
If you're debugging a model training pipeline for your product and Claude gives a bad answer, was the model confused? Did you give it bad context? Or did a hidden policy nerf Claude's ability to assist you?
