---
source: "https://arxiv.org/abs/2607.12624"
hn_url: "https://news.ycombinator.com/item?id=48921816"
title: "Detecting Prompt Injection Attacks on Purpose-Specific LLM Agents"
article_title: "[2607.12624] PVDetector: Detecting Prompt Injection Attacks on Purpose-Specific LLM Agents through Policy-Violation Concept Analysis"
author: "zhinit"
captured_at: "2026-07-15T15:19:02Z"
capture_tool: "hn-digest"
hn_id: 48921816
score: 1
comments: 0
posted_at: "2026-07-15T14:56:23Z"
tags:
  - hacker-news
  - translated
---

# Detecting Prompt Injection Attacks on Purpose-Specific LLM Agents

- HN: [48921816](https://news.ycombinator.com/item?id=48921816)
- Source: [arxiv.org](https://arxiv.org/abs/2607.12624)
- Score: 1
- Comments: 0
- Posted: 2026-07-15T14:56:23Z

## Translation

タイトル: 目的固有の LLM エージェントに対するプロンプト インジェクション攻撃の検出
記事タイトル: [2607.12624] PVDetector: ポリシー違反概念分析による目的固有の LLM エージェントに対するプロンプト インジェクション攻撃の検出
説明: arXiv 論文 2607.12624 の要約ページ: PVDetector: ポリシー違反概念分析による目的固有の LLM エージェントに対するプロンプト インジェクション攻撃の検出

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
[2026 年 7 月 14 日に提出]
タイトル: PVDetector: ポリシー違反概念分析による目的固有の LLM エージェントに対するプロンプト インジェクション攻撃の検出
要約: 大規模言語モデル (LLM) は、顧客サービスやコード生成などのドメイン固有のタスクを処理する目的固有のエージェントとして導入されることが増えています。これらのエージェントは、一般的な安全ガードレールだけでなく、指定された役割に合わせた目的固有の制限にも準拠することが期待されます。このような追加の制限により、特にプロンプ​​ト インジェクション (PI) 攻撃に対する攻撃対象領域が拡大します。このような攻撃を防ぐために、既存の検出方法は主に入出力パターンの分析に依存していますが、効果は限られています。この制限に対処するために、私たちは隠れたアクティベーション空間の分析に目を向け、LLM が指定された目的を超えたリクエストを要求された場合に本質的に潜在的なポリシー違反 (PV) の概念を保持していることを発見しました。特に、PV の概念は、ユーザーのクエリと事前定義された制限の間の競合のセマンティクスを捉えており、ポリシー違反を認識するという LLM の本質的な認識を暗黙的に反映しています。この洞察に基づいて、私たちは、ポリシー違反プロンプトとポリシー準拠プロンプトの対照的なペアからオフラインで導出される PV 概念との隠れ状態の整合性を測定することによって、LLM 推論中に PI 攻撃を検出するトレーニング不要のフレームワークである PVDetector を提案します。複数の LLM とデータセットにわたる実験では、PVDetector が最小限の補助オーバーヘッドで偽陰性率 <1\% を達成し、一貫して最先端の手法を上回るパフォーマンスを示しています。私たちのコードはこの https UR で入手できます

L .
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

Abstract page for arXiv paper 2607.12624: PVDetector: Detecting Prompt Injection Attacks on Purpose-Specific LLM Agents through Policy-Violation Concept Analysis

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Cryptography and Security
[Submitted on 14 Jul 2026]
Title: PVDetector: Detecting Prompt Injection Attacks on Purpose-Specific LLM Agents through Policy-Violation Concept Analysis
Abstract: Large language models (LLMs) are increasingly deployed as purpose-specific agents to handle domain-specific tasks such as customer service and code generation. These agents are expected to comply with not only generic safety guardrails but also purpose-specific restrictions tailored to their designated roles. Such additional restrictions enlarge the attack surface, particularly to prompt injection (PI) attacks. To defend against such attacks, existing detection methods primarily rely on analyzing input-output patterns, yet yield limited effectiveness. To address this limitation, we turn to analyzing the hidden activation space and discover that LLMs inherently retain latent policy-violation (PV) concepts when prompted with requests beyond their designated purpose. Particularly, PV concepts capture the semantics of conflicts between user queries and predefined restrictions, implicitly reflecting LLMs' intrinsic awareness of recognizing policy violations. Building on this insight, we propose PVDetector, a training-free framework that detects PI attacks during LLM inference by measuring hidden-state alignment with PV concepts, which are derived offline from the contrastive pairs of policy-violating and policy-compliant prompts. Experiments across multiple LLMs and datasets show that PVDetector achieves <1\% false negative rate with minimal auxiliary overhead, consistently outperforming state-of-the-art methods. Our code is available at this https URL .
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
