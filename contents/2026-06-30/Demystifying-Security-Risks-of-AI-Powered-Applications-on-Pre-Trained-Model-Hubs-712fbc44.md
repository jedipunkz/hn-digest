---
source: "https://arxiv.org/abs/2606.30373"
hn_url: "https://news.ycombinator.com/item?id=48737797"
title: "Demystifying Security Risks of AI-Powered Applications on Pre-Trained Model Hubs"
article_title: "[2606.30373] Your Space is My Zone: Demystifying the Security Risks of AI-Powered Applications on Pre-Trained Model Hubs"
author: "runningmike"
captured_at: "2026-06-30T19:33:23Z"
capture_tool: "hn-digest"
hn_id: 48737797
score: 2
comments: 1
posted_at: "2026-06-30T19:10:31Z"
tags:
  - hacker-news
  - translated
---

# Demystifying Security Risks of AI-Powered Applications on Pre-Trained Model Hubs

- HN: [48737797](https://news.ycombinator.com/item?id=48737797)
- Source: [arxiv.org](https://arxiv.org/abs/2606.30373)
- Score: 2
- Comments: 1
- Posted: 2026-06-30T19:10:31Z

## Translation

タイトル: 事前トレーニングされたモデル ハブ上の AI 搭載アプリケーションのセキュリティ リスクを解明する
記事のタイトル: [2606.30373] Your Space is My Zone: 事前トレーニングされたモデル ハブ上の AI 搭載アプリケーションのセキュリティ リスクを解明する
説明: arXiv 論文 2606.30373 の要約ページ: Your Space is My Zone: Demystifying the Security Risks of AI-Powered Applications on Pre-Trained Model Hubs

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
[2026 年 6 月 29 日に提出]
タイトル: Your Space is My Zone: 事前トレーニングされたモデル ハブ上の AI 搭載アプリケーションのセキュリティ リスクを解明する
要約: Hugging Face などのプラットフォームでホストされる AI 搭載アプリケーション (AI-App) は、オンライン推論と微調整サービスを通じて事前トレーニングされたモデルへのアクセスを民主化しています。 AI アプリは分離が弱く、セキュリティ設定が間違っている信頼できない当事者によって開発されることが多いため、これらのプラットフォームは AI 導入の障壁を下げる一方で、未踏の攻撃対象領域をもたらします。このペーパーでは、3 つの主要なプラットフォームにわたる AI アプリの初めての体系的なセキュリティ分析を紹介します。調査を構造化するために、AI アプリのライフサイクルを確立されたリスク分類 (OWASP など) にマッピングし、一般的な Web の欠陥から影響の大きいアーキテクチャの問題に至るまで、5 つの脅威カテゴリと 10 の攻撃ベクトルを特定します。私たちの分析により、壊れたアクセス制御、安全でないリソースの再利用、不十分な入力検証、機密データの漏洩などの重大な障害が明らかになりました。特に、プラットフォーム設計に固有の 3 つの新しいアーキテクチャ上の脆弱性を明らかにし、従来の問題 (例: 世界中で読み取り可能なログ) がこのエコシステム内でどのように独特に増幅されるかを実証します。現実世界への影響を評価するために、私たちは分析フレームワーク Insightor を開発し、それを 970,000 を超える公開 AI アプリに適用しています。驚くべきことに、数千のアプリが認証情報を漏洩し、数百のアプリに任意のコードの実行を可能にするインプットインジェクションの脆弱性が存在し、数十のアプリにバックドアが埋め込まれていることが判明しており、これは活発な悪用を示しています。私たちには責任があります

y はすべての調査結果を影響を受けるプラットフォームと開発者に開示しました。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2606.30373: Your Space is My Zone: Demystifying the Security Risks of AI-Powered Applications on Pre-Trained Model Hubs

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
[Submitted on 29 Jun 2026]
Title: Your Space is My Zone: Demystifying the Security Risks of AI-Powered Applications on Pre-Trained Model Hubs
Abstract: AI-powered Applications (AI-Apps), hosted on platforms such as Hugging Face, are democratizing access to pre-trained models through online inference and fine-tuning services. While lowering AI adoption barriers, these platforms introduce an unexplored attack surface, as AI-Apps are often developed by untrusted parties with weak isolation and misconfigured security settings. In this paper, we present the first systematic security analysis of AI-Apps across three leading platforms. To structure our investigation, we map the AI-App lifecycle to established risk taxonomies (e.g., OWASP), identifying five threat categories and ten attack vectors ranging from generic web flaws to high-impact architectural issues. Our analysis reveals critical failures including broken access control, insecure resource reuse, insufficient input validation, and sensitive data exposure. Notably, we uncover three novel architectural vulnerabilities inherent to platform design and demonstrate how traditional issues (e.g., world-readable logs) are uniquely amplified in this ecosystem. To assess real-world impact, we develop an analysis framework Insightor and apply it to over 970,000 public AI-Apps. Alarmingly, we find thousands of apps leaking credentials, hundreds containing input injection vulnerabilities that allow arbitrary code execution, and tens harboring embedded backdoors -- indicating active exploitation. We have responsibly disclosed all findings to the affected platforms and developers.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
