---
source: "https://arxiv.org/abs/2606.31408"
hn_url: "https://news.ycombinator.com/item?id=48850154"
title: "EnclaveX: End-to-End Confidential AI with CPU/GPU Tees"
article_title: "[2606.31408] EnclaveX: End-to-End Confidential AI with CPU/GPU TEEs"
author: "Jimmc414"
captured_at: "2026-07-09T18:47:13Z"
capture_tool: "hn-digest"
hn_id: 48850154
score: 1
comments: 0
posted_at: "2026-07-09T18:11:12Z"
tags:
  - hacker-news
  - translated
---

# EnclaveX: End-to-End Confidential AI with CPU/GPU Tees

- HN: [48850154](https://news.ycombinator.com/item?id=48850154)
- Source: [arxiv.org](https://arxiv.org/abs/2606.31408)
- Score: 1
- Comments: 0
- Posted: 2026-07-09T18:11:12Z

## Translation

タイトル: EnclaveX: CPU/GPU T を使用したエンドツーエンドの機密 AI
記事のタイトル: [2606.31408] EnclaveX: CPU/GPU TEE を使用したエンドツーエンドの機密 AI
説明: arXiv 論文 2606.31408 の要約ページ: EnclaveX: CPU/GPU TEE を使用したエンドツーエンドの機密 AI

記事本文:
メインコンテンツにスキップ
arXiv は独立した非営利団体になりました。
さらに詳しく
×
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピューターサイエンス > 暗号化とセキュリティ
[2026 年 6 月 30 日に提出]
タイトル: EnclaveX: CPU/GPU TEE を使用したエンドツーエンドの機密 AI
要約: 大規模言語モデル (LLM) は急速に普及し、AI アプリケーションの広範な採用を推進しています。ほとんどの展開は Microsoft Azure、Google Cloud、AWS などの一元化されたインフラストラクチャに依存しており、ユーザーは機密データを共有し、コードをトレーニングまたは微調整する必要があります。クラウドプロバイダーは機密性と整合性を確保するために信頼されなければならないため、この依存性はセキュリティとプライバシーに関する重大な懸念を引き起こします。
これらのリスクを軽減するために、Intel SGX/TDX、AMD SEV-SNP、ARM CCA などの信頼された実行環境 (TEE) が導入されました。最近では、NVIDIA が GPU TEE (H100/H200 など) を開発しましたが、CPU と GPU TEE を統合するエンドツーエンドのワークフローの包括的な評価は依然として限られています。パフォーマンスのオーバーヘッド、リモート認証、AI/LLM アプリケーションのセキュリティ保証などの重要な側面は、十分に研究されていません。
このペーパーでは、CPU と GPU TEE を組み合わせたエンドツーエンドのワークフローを提示することで、このギャップに対処します。 VM レベル (Intel TDX および AMD SEV-SNP 経由) とアプリケーション レベルの両方で機密性と整合性を確保するメカニズムを提案し、Kubernetes 管理者が VM の機密コンテンツにアクセスする機能などの脆弱性を強調します。最後に、Intel TDX と NVIDIA H200 GPU を統合する構成に焦点を当て、業界ベンチマークを使用してシステムのパフォーマンス オーバーヘッドを評価します。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
投稿履歴
参考文献

aphic と引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2606.31408: EnclaveX: End-to-End Confidential AI with CPU/GPU TEEs

Skip to main content
arXiv is now an independent nonprofit!
Learn more
×
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Cryptography and Security
[Submitted on 30 Jun 2026]
Title: EnclaveX: End-to-End Confidential AI with CPU/GPU TEEs
Abstract: Large Language Models (LLMs) have rapidly proliferated, driving widespread adoption of AI applications. Most deployments rely on centralized infrastructures such as Microsoft Azure, Google Cloud, or AWS, requiring users to share sensitive data and training or fine-tuning code. This dependence raises significant security and privacy concerns, as cloud providers must be trusted to ensure confidentiality and integrity.
Trusted Execution Environments (TEEs) e.g., Intel SGX/TDX, AMD SEV-SNP, and ARM CCA have been introduced to mitigate these risks. More recently, NVIDIA has developed GPU TEEs (e.g., H100/H200), yet comprehensive evaluations of end-to-end workflows that integrate CPU and GPU TEEs remain limited. Critical aspects, including performance overhead, remote attestation, and security guarantees for AI/LLM applications, have not been sufficiently studied.
This paper addresses this gap by presenting an end-to-end workflow that combines CPU and GPU TEEs. We propose mechanisms to ensure confidentiality and integrity at both the VM level (via Intel TDX and AMD SEV-SNP) and the application level, highlighting vulnerabilities such as Kubernetes administrators' ability to access confidential VM contents. Finally, we evaluate the performance overhead of our system using industry benchmarks, focusing on configurations that integrate Intel TDX with NVIDIA H200 GPUs.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
