---
source: "https://arxiv.org/abs/2608.09867"
hn_url: "https://news.ycombinator.com/item?id=49254407"
title: "Stealing Reasoning Traces from Proprietary LLM APIs"
article_title: "[2608.09867] Stealing Reasoning Traces from Proprietary LLM APIs"
author: "sbulaev"
captured_at: "2026-08-11T08:03:53Z"
capture_tool: "hn-digest"
hn_id: 49254407
score: 1
comments: 0
posted_at: "2026-08-11T07:07:06Z"
tags:
  - hacker-news
  - translated
---

# Stealing Reasoning Traces from Proprietary LLM APIs

- HN: [49254407](https://news.ycombinator.com/item?id=49254407)
- Source: [arxiv.org](https://arxiv.org/abs/2608.09867)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T07:07:06Z

## Translation

タイトル: 独自の LLM API から推論トレースを盗む
記事のタイトル: [2608.09867] 独自の LLM API から推論トレースを盗む
説明: arXiv 論文 2608.09867 の要約ページ: 独自の LLM API から推論トレースを盗む

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピューターサイエンス > 暗号化とセキュリティ
[2026 年 8 月 10 日に提出]
タイトル: 独自の LLM API から推論トレースを盗む
要約: 大手の大規模言語モデルプロバイダーは現在、知的財産を保護し、情報漏洩を制限するために、モデルの段階的な推論、つまり思考の連鎖を隠蔽しています。プロバイダーは、これらのトレースをサーバー側に保存するのではなく、暗号化されたテキストのブロックとしてクライアントに返し、クライアントはそれを後続の各リクエストで返します。以前の調査に基づいて、私たちはアーキテクチャ上の脆弱性を特定しました。これらの暗号化されたブロックは、プロバイダーのエコシステム内のさまざまなセッション、ユーザー、モデル間で完全に互換性があり、交換可能です。私たちはこの互換性を利用して、スケーラブルな復号化ジェイルブレイクを開発します。特定のモデルからの暗号化された推論トレースを、同じプロバイダーのより脆弱で安全性の低いモデルに注入することで、より有能なモデルを直接脱獄することなく、トレースをそのまま平文でデコードして出力するように強制します。この脆弱性により、4 つの異なる攻撃ベクトルが可能になります。まず、Anthropic、OpenAI、Google で実証されているように、これは反蒸留メカニズムを回避し、敵対者が独自のモデルの推論を抽出できるようにします。 2 番目に、大規模なプライベート データの抽出が可能になります。開発者は、暗号化されたブロックの内容を意識せずに、セッション ログをパブリックに共有することがよくあります。パブリック リポジトリから収集した 315,320 個の推論ブロックを解読することで、367 個の個人識別情報 (PII) アーティファクトと 182 個の資格情報を復元しました。第三に、たとえ次のような場合であっても、推論プロセス内に隠された危険な情報を誤って明らかにしてしまいます。

モデルの最終的な目に見える出力は、悪意のあるリクエストを安全に拒否します。第 4 に、攻撃者はこの欠陥を利用して目に見えないプロンプト インジェクションを実行し、暗号化されたブロック内に悪意のあるペイロード全体を埋め込み、パブリック エージェントのロールアウトを妨害する可能性があります。責任ある開示に続いて、クライアント側の推論を保護するための具体的な暗号化およびシステムレベルの緩和策を提案します。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2608.09867: Stealing Reasoning Traces from Proprietary LLM APIs

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Cryptography and Security
[Submitted on 10 Aug 2026]
Title: Stealing Reasoning Traces from Proprietary LLM APIs
Abstract: Leading large language model providers now conceal their models' step-by-step reasoning, or chain-of-thought, to protect intellectual property and limit information leakage. Rather than storing these traces server-side, providers return them to the client as blocks of encrypted text, which the client passes back with each subsequent request. Building on prior research, we identify an architectural vulnerability: these encrypted blocks are fully compatible and interchangeable across different sessions, users, and models within a provider's ecosystem. We exploit this compatibility to develop a scalable decryption jailbreak. By injecting an encrypted reasoning trace from a given model into a weaker, and less safeguarded model from the same provider, we force it to decode and output the trace verbatim in plaintext, without ever jailbreaking the more capable model directly. This vulnerability enables four distinct attack vectors. First, it circumvents anti-distillation mechanisms, allowing adversaries to extract a proprietary model's reasoning, as we demonstrate across Anthropic, OpenAI, and Google. Second, it allows for large-scale private data extraction. Developers frequently share session logs publicly, unaware of contents of the encrypted blocks. By decoding 315,320 reasoning blocks scraped from public repositories, we recovered 367 Personally Identifiable Information (PII) artifacts and 182 credentials. Third, it inadvertently reveals hazardous information hidden within the reasoning process, even in cases where the model's final, visible output safely rejects a malicious request. Fourth, attackers can leverage this flaw to execute invisible prompt injections, embedding malicious payloads entirely within encrypted blocks to poison public agentic rollouts. Following responsible disclosure, we propose concrete cryptographic and system-level mitigations to secure client-side reasoning.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
