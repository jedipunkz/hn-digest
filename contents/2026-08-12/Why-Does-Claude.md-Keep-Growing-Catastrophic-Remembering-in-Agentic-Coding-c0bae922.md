---
source: "https://arxiv.org/abs/2608.11095"
hn_url: "https://news.ycombinator.com/item?id=49268790"
title: "Why Does Claude.md Keep Growing? Catastrophic Remembering in Agentic Coding"
article_title: "[2608.11095] Why Does CLAUDE.md Keep Growing? Catastrophic Remembering in Agentic Coding"
author: "sbulaev"
captured_at: "2026-08-12T07:07:58Z"
capture_tool: "hn-digest"
hn_id: 49268790
score: 1
comments: 0
posted_at: "2026-08-12T07:07:06Z"
tags:
  - hacker-news
  - translated
---

# Why Does Claude.md Keep Growing? Catastrophic Remembering in Agentic Coding

- HN: [49268790](https://news.ycombinator.com/item?id=49268790)
- Source: [arxiv.org](https://arxiv.org/abs/2608.11095)
- Score: 1
- Comments: 0
- Posted: 2026-08-12T07:07:06Z

## Translation

タイトル: なぜClaude.mdは成長し続けるのか?エージェントコーディングにおける壊滅的な記憶
記事タイトル: [2608.11095] CLAUDE.md はなぜ成長し続けるのか?エージェントコーディングにおける壊滅的な記憶
説明: arXiv 論文 2608.11095 の要約ページ: CLAUDE.md はなぜ成長し続けるのか?エージェントコーディングにおける壊滅的な記憶

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピュータサイエンス > 人工知能
[2026 年 8 月 11 日に提出]
タイトル: CLAUDE.md はなぜ成長を続けるのか?エージェントコーディングにおける壊滅的な記憶
要約: この http URL のようなエージェント コーディング README は、実際のリポジトリ内で無制限に増加し、リポジトリが廃止されるか、誰かがファイルを全面的に書き換えた場合にのみ停止します。これを不完全な再現にたどります。命令の追加は常に低コストですが、命令の根拠が失われると、正確性の低下の危険を冒さずに命令を削除すると、|D| のプロンプトで O(2^|D|) のコストがかかります。説明書。私たちは、その結果として生じる発散を、壊滅的な記憶、つまり継続的な学習が組織される壊滅的な忘却の逆と名付けます。まず、1,867 のリポジトリにおける 247,694 の命令の存続期間にわたってこの現象を特徴付けます。エージェント プロンプトは際限なく増加し、存続期間中に 3 倍以上 (+226%)、コミットごとに正味命令数が +4.9 増加します。さらに、命令が古くなるほど、削除される可能性は低くなります (log-hazard -0.032/commit)。次に、プロンプトコメントが成長を止めることができることを示します。IFEval を反転すると、最適なプロンプトが既知である検証可能な世界が得られ、潜在的な推論をエンコードしたコメントによって過剰な命令が 99.3% 除去されます (+211.3% から +1.4%)。最後に、同じ逆変換を WildIFEval に適用すると、即時コメントによって現実世界のエージェントによる指示のフォローが最大 23.1% 向上することがわかります。英語が新しいコードである場合、なぜまだコメントがないのでしょうか?
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: 実験的プロジェクト

h コミュニティの協力者
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2608.11095: Why Does CLAUDE.md Keep Growing? Catastrophic Remembering in Agentic Coding

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Artificial Intelligence
[Submitted on 11 Aug 2026]
Title: Why Does CLAUDE.md Keep Growing? Catastrophic Remembering in Agentic Coding
Abstract: Agentic coding READMEs like this http URL grow without bound in real repositories, stopping only when the repository retires or someone rewrites the file wholesale. We trace this to imperfect recall: appending an instruction is always cheap, but once an instruction's rationale is gone, deleting it without risking a correctness regression costs O(2^|D|) in a prompt of |D| instructions. We name the resulting divergence catastrophic remembering, the inverse of catastrophic forgetting around which continual learning is organized. First, we characterize this phenomenon across 247,694 instruction lifetimes in 1,867 repositories: agentic prompts grow without bound, more than tripling over their lifetime (+226%), gaining +4.9 net instructions every commit; further, the older an instruction gets, the less likely it is to be deleted (log-hazard -0.032/commit). Then, we show that prompt comments can halt the growth: inverting IFEval yields verifiable worlds whose optimal prompts are known, and there comments encoding latent reasoning remove 99.3% of excess instructions (+211.3% to +1.4%). Finally, applying the same inversion to WildIFEval, we show that prompt comments can improve real-world agentic instruction-following by up to 23.1%. If English is the new code, why don't we have comments yet?
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
