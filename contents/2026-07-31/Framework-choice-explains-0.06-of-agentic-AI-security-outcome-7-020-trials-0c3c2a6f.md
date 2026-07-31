---
source: "https://figshare.com/articles/preprint/Cross-Framework_Portability_of_Agentic_AI_Security_A_Controlled_Payload-Verified_Evaluation/33110642"
hn_url: "https://news.ycombinator.com/item?id=49129575"
title: "Framework choice explains ~0.06% of agentic AI security outcome (7,020 trials)"
article_title: "Item - Cross-Framework Portability of Agentic AI Security: A Controlled, Payload-Verified Evaluation - figshare - Figshare"
author: "waqarjaved"
captured_at: "2026-07-31T23:54:15Z"
capture_tool: "hn-digest"
hn_id: 49129575
score: 2
comments: 0
posted_at: "2026-07-31T23:17:35Z"
tags:
  - hacker-news
  - translated
---

# Framework choice explains ~0.06% of agentic AI security outcome (7,020 trials)

- HN: [49129575](https://news.ycombinator.com/item?id=49129575)
- Source: [figshare.com](https://figshare.com/articles/preprint/Cross-Framework_Portability_of_Agentic_AI_Security_A_Controlled_Payload-Verified_Evaluation/33110642)
- Score: 2
- Comments: 0
- Posted: 2026-07-31T23:17:35Z

## Translation

タイトル: フレームワークの選択により、エージェント AI セキュリティの結果の ~0.06% が説明される (7,020 件の試験)
記事のタイトル: 項目 - Agentic AI セキュリティのフレームワーク間移植性: 制御されたペイロード検証済みの評価 - figshare - Figshare
説明: Agentic AI セキュリティのフレームワーク間の移植性: 制御されたペイロード検証済みの評価

記事本文:
メイン コンテンツにスキップ 参照 参照と検索 検索 Agentic AI セキュリティのフレームワーク間移植性: 制御されたペイロード検証済みの評価
https://doi.org/10.6084/m9.figshare.33110642 識別子の URL をクリップボードにコピー 識別子情報
バージョン 5 2026-07-30、20:54 バージョン 5 2026-07-30、20:54 バージョン 4 2026-07-30、18:51 バージョン 4 2026-07-30、18:51 バージョン 3 2026-07-29、08:07 バージョン 3 2026-07-29、08:07 バージョン 2 2026-07-29、07:47 バージョン 2 2026-07-29、07:47 バージョン 1 2026-07-29、07:46 バージョン 1 2026-07-29、07:46 プレプリント投稿日: 2026-07-30、 20:54 作成者: Waqar Javed Waqar Javed <p dir="ltr">Agentic AI の導入では、LangChain、CrewAI、AutoGen、LlamaIndex、OpenAI Agents SDK などのオーケストレーション フレームワークを通じて大規模な言語モデルをルーティングすることが増えています。各テンプレートは、コンテンツがモデルに到達する前にプロンプトを表示し、メッセージを構造化し、ツール スキーマのフォーマットを異なる方法で設定します。これまでの比較研究では、攻撃耐性におけるフレームワークレベルの大きな違い（たとえば、モデルを一定に保った場合、2 つのフレームワーク間で拒否率 52.3% 対 30.8%）が報告されていますが、基礎となる攻撃ペイロードが各フレームワークのアダプター層にわたって同一に配信されるかどうかは検証されておらず、そのような違いがフレームワーク自体を反映しているのか、それともアダプターレベルのプロンプト変動のアーティファクトを反映しているのかは未解決のままです。当社はこれまでで最大規模のクロスフレームワーク エージェント セキュリティ評価を提示します。2 つの機能層にわたる 6 つのモデル、6 つの実行条件 (直接 API ベースラインと 5 つのエージェント フレームワーク)、および 5 つの OWASPASI 調整攻撃ファミリー、合計 7,020 件のトライアルであり、フレームワークを制御変数として分離するためのすべてのアダプターにわたるバイトレベルのペイロード ID 検証が行われます。 ANOVA 分解では、結果の分散のうちフレームワークの選択によるものは 0.06% のみであるのに対し、フレームワークの選択によるものは 28.67% です。

攻撃ファミリーでは 4.23%、モデルでは 4.23% — この差は順列テストでは偶然と区別がつかないことが判明し (p = 0.70)、別の結果重み付け感度チェックでも生き残ります (0.06%→0.56%)。私たちの検証手順自体は、修正前に判定を大きく変える実際のアダプターの欠陥を捕捉しました。これは、明らかなフレームワークの影響が、モデルに対するフレームワークの動作ではなく、アダプターに起因する可能性があることを示す直接的な証拠です。これらの結果が示唆するのは、
[切り捨てられた]
1. DOI - LLM 拒否検出におけるパターンマッチングの失敗を参照: 検出器の信頼性に関するケーススタディ
ソフトウェアとアプリケーションのセキュリティ
自律エージェントとマルチエージェント システム

## Original Extract

Cross-Framework Portability of Agentic AI Security: A Controlled, Payload-Verified Evaluation

Skip to main content Browse Browse and Search Search Cross-Framework Portability of Agentic AI Security: A Controlled, Payload-Verified Evaluation
https://doi.org/10.6084/m9.figshare.33110642 Copy identifier URL to clipboard Identifier Info
Version 5 2026-07-30, 20:54 Version 5 2026-07-30, 20:54 Version 4 2026-07-30, 18:51 Version 4 2026-07-30, 18:51 Version 3 2026-07-29, 08:07 Version 3 2026-07-29, 08:07 Version 2 2026-07-29, 07:47 Version 2 2026-07-29, 07:47 Version 1 2026-07-29, 07:46 Version 1 2026-07-29, 07:46 preprint posted on 2026-07-30, 20:54 authored by Waqar Javed Waqar Javed <p dir="ltr">Agentic AI deployments increasingly route large language models through orchestration frameworks such as LangChain, CrewAI, AutoGen, LlamaIndex, and the OpenAI Agents SDK, each of which templates prompts, structures messages, and formats tool schemas differently before content reaches the model. Prior comparative work has reported large framework-level differences in attack resistance (e.g., 52.3% vs. 30.8% refusal across two frameworks, holding model constant), but without verifying that the underlying attack payload is delivered identically across each framework’s adapter layer, leaving open whether such differences reflect the framework itself or an artifact of adapter-level prompt variation. We present the largest cross-framework agentic security evaluation to date: 6 models across two capability tiers, 6 execution conditions (a direct-API baseline plus five agentic frameworks), and 5 OWASPASI-aligned attack families, totaling 7,020 trials, with bytelevel payload-identity verification across every adapter to isolate framework as a controlled variable. An ANOVA decomposition attributes only 0.06% of outcome variance to framework choice, against 28.67% for attack family and 4.23% for model — a difference a permutation test finds indistinguishable from chance (p = 0.70) and that survives an alternate outcomeweighting sensitivity check (0.06%→0.56%). Our verification procedure itself caught a real adapter defect that measurably shifted verdicts before correction, direct evidence that apparent framework effects can originate in the adapter rather than the framework’s behavior on the model. These results sugg
[truncated]
1. DOI - References Pattern-Matching Failures in LLM Refusal Detection: A Case Study in Detector Reliability
Software and application security
Autonomous agents and multiagent systems
