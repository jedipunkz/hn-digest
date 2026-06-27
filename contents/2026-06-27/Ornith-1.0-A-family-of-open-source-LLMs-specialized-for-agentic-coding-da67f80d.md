---
source: "https://twitter.com/ornith_/status/2070148887067963854"
hn_url: "https://news.ycombinator.com/item?id=48697068"
title: "Ornith-1.0: A family of open-source LLMs specialized for agentic coding"
article_title: "Ornith on X: \"Aloha! 🌺 Meet Ornith-1.0, a family of open-source LLMs specialized for agentic coding.\nOrnith-1.0 spans the full parameter sizes including 9B Dense, 31B Dense, 35B MoE, and 397B MoE. It achieves state-of-the-art performance among open-source models of comparable size on https://t.co/7g\n[truncated]"
author: "jalopy"
captured_at: "2026-06-27T11:30:13Z"
capture_tool: "hn-digest"
hn_id: 48697068
score: 3
comments: 0
posted_at: "2026-06-27T10:45:38Z"
tags:
  - hacker-news
  - translated
---

# Ornith-1.0: A family of open-source LLMs specialized for agentic coding

- HN: [48697068](https://news.ycombinator.com/item?id=48697068)
- Source: [twitter.com](https://twitter.com/ornith_/status/2070148887067963854)
- Score: 3
- Comments: 0
- Posted: 2026-06-27T10:45:38Z

## Translation

タイトル: Ornith-1.0: エージェント コーディングに特化したオープンソース LLM ファミリ
記事のタイトル: Ornith on X: 「アロハ! 🌺 エージェント コーディングに特化したオープンソース LLM ファミリである Ornith-1.0 をご紹介します。
Ornith-1.0 は、9B Dense、31B Dense、35B MoE、397B MoE を含む全パラメータ サイズに対応します。 https://t.co/7g の同等のサイズのオープンソース モデルの中で最先端のパフォーマンスを実現します。
[切り捨てられた]
説明: アロハ! 🌺 エージェントコーディングに特化したオープンソース LLM ファミリである Ornith-1.0 をご紹介します。
Ornith-1.0 は、9B Dense、31B Dense、35B MoE、397B MoE を含む全パラメータ サイズに対応します。 https://t.co/7g1rmacLps の同等のサイズのオープンソース モデルの中で最先端のパフォーマンスを実現します。

記事本文:
@ornith_ アロハ！ 🌺 エージェントコーディングに特化したオープンソース LLM ファミリである Ornith-1.0 をご紹介します。
Ornith-1.0 は、9B Dense、31B Dense、35B MoE、397B MoE を含む全パラメータ サイズに対応します。以下のようなコーディング ベンチマークにおいて、同等のサイズのオープンソース モデルの中で最先端のパフォーマンスを実現します。
✅ターミナルベンチ 2.1(77.5)
✅SWE-Bench (検証済みで 82.4、プロで 62.2、多言語で 78.9)
✅NL2リポジトリ(48.2)
✅SWE アトラス (QnA で 41.2、RF 42.6、TW 39.1)
✅ClawEval(77.1)
gemma4 と qwen3.5 上でポストトレーニングされた Ornith-1.0 は、強化学習を使用してソリューションのロールアウトだけでなく、それらのロールアウトを推進するタスク固有のスキャフォールドも生成する、新しい自己改善トレーニング戦略を採用しています。足場と結果として得られるソリューションを共同で最適化することで、モデルはエージェントティック コーディングで高品質のソリューションを生成します。😎
すべてのモデルは MIT ライセンスの下でリリースされており、商用および研究での完全な使用が可能です。
📖技術ブログ: deep-reinforce.com/ornith_1_0.html
🤗Huggingface:huggingface.co/collections/de… 2:15 PM · 2026 年 6 月 25 日 4.8M 再生数 4 4 8 448 9 0 8 908 6 K 6K 5 。 8 K 5.8K 448 件の返信を読む X は初めてですか?
今すぐサインアップして、自分だけのタイムラインを手に入れましょう!
Google でサインアップ Apple でサインアップ アカウントを作成 サインアップすると、Cookie の使用を含むサービス利用規約とプライバシー ポリシーに同意したことになります。

## Original Extract

Aloha! 🌺 Meet Ornith-1.0, a family of open-source LLMs specialized for agentic coding.
Ornith-1.0 spans the full parameter sizes including 9B Dense, 31B Dense, 35B MoE, and 397B MoE. It achieves state-of-the-art performance among open-source models of comparable size on https://t.co/7g1rmacLps

@ornith_ Aloha! 🌺 Meet Ornith-1.0, a family of open-source LLMs specialized for agentic coding.
Ornith-1.0 spans the full parameter sizes including 9B Dense, 31B Dense, 35B MoE, and 397B MoE. It achieves state-of-the-art performance among open-source models of comparable size on coding benchmarks including:
✅Terminal-Bench 2.1(77.5)
✅SWE-Bench(82.4 on verified, 62.2 on pro, 78.9 on Multilingual)
✅NL2Repo(48.2)
✅SWE Atlas(41.2 on QnA, 42.6 RF, 39.1 TW)
✅ClawEval(77.1)
Post-trained on top of gemma4 and qwen3.5, Ornith-1.0 employs a novel self-improving training strategy in which reinforcement learning is used to generate not only solution rollouts, but also the task-specific scaffolds that drive those rollouts. By jointly optimizing the scaffold and the resulting solution, the model generate higher-quality solutions in agentic coding.😎
All models are released under the MIT license, enabling full commercial and research use.
📖Tech Blog: deep-reinforce.com/ornith_1_0.html
🤗Huggingface: huggingface.co/collections/de… 2:15 PM · Jun 25, 2026 4.8M Views 4 4 8 448 9 0 8 908 6 K 6K 5 . 8 K 5.8K Read 448 replies New to X?
Sign up now to get your own personalized timeline!
Sign up with Google Sign up with Apple Create account By signing up, you agree to the Terms of Service and Privacy Policy , including Cookie Use.
