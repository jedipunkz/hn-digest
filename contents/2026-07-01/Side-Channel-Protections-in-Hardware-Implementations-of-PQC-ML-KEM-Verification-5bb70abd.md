---
source: "https://arxiv.org/abs/2606.31681"
hn_url: "https://news.ycombinator.com/item?id=48743913"
title: "Side-Channel Protections in Hardware Implementations of PQC ML-KEM Verification"
article_title: "[2606.31681] Exploring Side-Channel Protections in Hardware Implementations of PQC ML-KEM Verification"
author: "austinallegro"
captured_at: "2026-07-01T09:06:01Z"
capture_tool: "hn-digest"
hn_id: 48743913
score: 2
comments: 0
posted_at: "2026-07-01T08:43:24Z"
tags:
  - hacker-news
  - translated
---

# Side-Channel Protections in Hardware Implementations of PQC ML-KEM Verification

- HN: [48743913](https://news.ycombinator.com/item?id=48743913)
- Source: [arxiv.org](https://arxiv.org/abs/2606.31681)
- Score: 2
- Comments: 0
- Posted: 2026-07-01T08:43:24Z

## Translation

タイトル: PQC ML-KEM 検証のハードウェア実装におけるサイドチャネル保護
記事のタイトル: [2606.31681] PQC ML-KEM 検証のハードウェア実装におけるサイドチャネル保護の探索
説明: arXiv 論文 2606.31681 の要約ページ: Exploring Side-Channel Protections in Hardware Implementations of PQC ML-KEM Verification

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
タイトル: PQC ML-KEM 検証のハードウェア実装におけるサイドチャネル保護の探求
要約: ML-KEM がポスト量子暗号標準として採用されるにつれて、物理サイドチャネル攻撃に対する回復力が不可欠になりました。構成ステップの中で、カプセル化解除の Fujisaki-Okamoto (FO) 検証は、サイドチャネル電力および電磁 (EM) 解析に対して特に脆弱です。この研究では、一般的な FPGA ベースの実装に焦点を当て、そのサイドチャネルの脆弱性を調査し、マイクロコントローラー実装の脆弱性と比較します。保護されていない、ハッシュベース (1 次)、および高次マスクされた 3 つの検証実装が、マイクロコントローラーと FPGA の両方でサイドチャネル セキュリティについて評価されます。 FPGA は高速性と並列性を実現しますが、特に高帯域幅構成ではサイドチャネル リークが強くなることがよくあります。高次のマスクされた設計でも、ハードウェア レベルの影響とデータ依存の処理により、基礎となるデータに関する情報が漏洩します。私たちの実験では、FPGA での並列処理により、秘密鍵を完全に回復するのに十分な一次漏洩が発生することがわかりました。これらの結果は、パフォーマンスに制約があり、並列化されたハードウェア環境で PQC アルゴリズムを保護するという永続的な課題を浮き彫りにしています。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティとのコラボレーションによる実験プロジェクト

レーター
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2606.31681: Exploring Side-Channel Protections in Hardware Implementations of PQC ML-KEM Verification

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
Title: Exploring Side-Channel Protections in Hardware Implementations of PQC ML-KEM Verification
Abstract: As ML-KEM is adopted as a post-quantum cryptographic standard, resilience against physical side-channel attacks has become essential. Among the constituent steps, the decapsulation Fujisaki-Okamoto (FO) verification is particularly vulnerable to side-channel power and electromagnetic (EM) analysis. In this work, we focus on common FPGA-based implementations and examine their side-channel vulnerabilities, and compare them with those of microcontroller implementations. Three verification implementations, unprotected, hash-based (first-order), and higher-order masked, are evaluated for side-channel security on both a microcontroller and an FPGA. While FPGAs offer higher speed and parallelism, they often exhibit stronger side-channel leakage, especially in high bandwidth configurations. The higher-order masked designs still leak information about the underlying data due to hardware-level effects and data-dependent processing. Our experiments show that their parallelized processing on FPGAs introduces sufficient first-order leakage for full secret-key recovery. These results underscore the persistent challenge of securing PQC algorithms in performance-constrained and parallelized hardware environments.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
