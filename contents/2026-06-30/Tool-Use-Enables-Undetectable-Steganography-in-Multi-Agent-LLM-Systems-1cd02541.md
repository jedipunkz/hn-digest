---
source: "https://arxiv.org/abs/2606.28425"
hn_url: "https://news.ycombinator.com/item?id=48733505"
title: "Tool Use Enables Undetectable Steganography in Multi-Agent LLM Systems"
article_title: "[2606.28425] Tool Use Enables Undetectable Steganography in Multi-Agent LLM Systems"
author: "Brajeshwar"
captured_at: "2026-06-30T15:50:35Z"
capture_tool: "hn-digest"
hn_id: 48733505
score: 3
comments: 0
posted_at: "2026-06-30T14:50:11Z"
tags:
  - hacker-news
  - translated
---

# Tool Use Enables Undetectable Steganography in Multi-Agent LLM Systems

- HN: [48733505](https://news.ycombinator.com/item?id=48733505)
- Source: [arxiv.org](https://arxiv.org/abs/2606.28425)
- Score: 3
- Comments: 0
- Posted: 2026-06-30T14:50:11Z

## Translation

タイトル: ツールの使用により、マルチエージェント LLM システムで検出不可能なステガノグラフィーが可能になります
記事のタイトル: [2606.28425] ツールの使用により、マルチエージェント LLM システムで検出不可能なステガノグラフィーが可能になります
説明: arXiv 論文 2606.28425 の要約ページ: ツールの使用により、マルチエージェント LLM システムで検出不可能なステガノグラフィーが可能になります

記事本文:
メインコンテンツにスキップ
arXiv の提出は、6 月 30 日火曜日 14:00 EDT からメンテナンスのため停止されます。それ以外の場合は、サイトは稼働し続ける必要があります。
コンピューターサイエンス > 暗号化とセキュリティ
[2026 年 6 月 25 日に提出]
タイトル: ツールの使用により、マルチエージェント LLM システムで検出不可能なステガノグラフィーが可能になります
要約: ますます自律的なエージェント AI システムは、秘密の通信チャネルを介した秘密の共謀など、新たなマルチエージェントのリスクを引き起こします。こうした共謀の試みに対する自然な防御策は、平文通信を監視することですが、ますます洗練されたモデルのステガノグラフィーによって、監視の有効性が疑問視されています。実際、情報理論上または計算上、誠実な平文通信と区別できないいくつかの理論的スキームが提案されています。この論文では、コードの実行や Web 検索による研究論文へのアクセスなど、現実的なツールの使用状況を考慮すると、エージェント コーディング モデルはすでに検出不可能なステゴシステムを生成できるため、これらのスキームの複雑さがもはや安全障壁ではないことを実証します。エージェントは、主要な要素が欠落している場合にも、たとえばモデル サンプリング コンポーネントを追加したり、関連するキー付きコーディング スキームを実装したりすることで適応します。次に、エージェント間の暗黙のステガノグラフィー調整をシェリング点問題として枠組み化し、明示的な事前合意なしに 2 つのエージェントが互換性のあるスキームを選択する可能性が高い時期を推定するための調整メトリクスを導入します。私たちの結果は、AI エージェント間の秘密通信の脅威モデルの変化を示唆しています。そこでの主な障壁は、もはやフロンティア エージェントが高度なステゴシステムを理解して実装できるかどうかではなく、調整、つまり独立して動作するエージェントが互換性のあるスキーム、キー、パラメータに収束できるかどうかです。私たちは

広範なスキームファミリーでは実質的な収束が見られたが、厳密なワンショット調整は限られており、共有アーティファクト、反復的な対話、およびツールを介した検索が、秘密コミュニケーションのリスクが最も深刻な状況であることを示唆しています。全体として、我々の発見は、有能なエージェントが監視を生き延びる秘密のチャネルを構築できると仮定する、最近の戦略的監禁仮説に経験的な根拠を提供するものである。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI
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

Abstract page for arXiv paper 2606.28425: Tool Use Enables Undetectable Steganography in Multi-Agent LLM Systems

Skip to main content
arXiv submission will be down for maintenance beginning 14:00 EDT Tuesday June 30th. The site should otherwise remain in operation.
Computer Science > Cryptography and Security
[Submitted on 25 Jun 2026]
Title: Tool Use Enables Undetectable Steganography in Multi-Agent LLM Systems
Abstract: Increasingly autonomous agentic AI systems pose novel multi-agent risks, such as secret collusion via covert communication channels. The natural defence to these collusion attempts is to monitor plain-text communication, but the efficacy of monitors has been called into doubt by increasingly sophisticated model steganography; indeed, some theoretical schemes have been proposed that are information-theoretically or computationally indistinguishable from good-faith plain-text communication. In this paper, we demonstrate that the complexity of these schemes is no longer a safety barrier, as agentic coding models can already produce undetectable stegosystems when given realistic tool usage, such as code execution or accessing research papers through web searches. Agents also adapt when key ingredients are missing, for example, by adding model-sampling components or implementing related keyed coding schemes. We then frame tacit steganographic coordination between agents as a Schelling-point problem and introduce coordination metrics for estimating when two agents are likely to select compatible schemes without explicit prior agreement. Our results suggest a shift in the threat model for covert communication between AI agents, where the main barrier is no longer whether frontier agents can understand and implement sophisticated stegosystems, but coordination: whether independently acting agents can converge on compatible schemes, keys, and parameters. We find substantial convergence on broad scheme families but limited strict one-shot coordination, suggesting that shared artefacts, repeated interaction, and tool-mediated search are the settings where covert communication risks are most acute. Overall, our findings provide empirical grounding for the recent strategic confinement hypothesis, which assumes that capable agents can construct covert channels that survive monitoring.
Focus to learn more
arXiv-issued DOI via DataCite
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
