---
source: "https://arxiv.org/abs/2607.20759"
hn_url: "https://news.ycombinator.com/item?id=49136144"
title: "IssueTrojanBench: Benchmarking AI Coding Agents Against Malicious Issue Requests"
article_title: "[2607.20759] IssueTrojanBench: Benchmarking AI Coding Agents Against Malicious Issue Requests"
author: "yruzin"
captured_at: "2026-08-01T17:53:55Z"
capture_tool: "hn-digest"
hn_id: 49136144
score: 2
comments: 0
posted_at: "2026-08-01T16:59:18Z"
tags:
  - hacker-news
  - translated
---

# IssueTrojanBench: Benchmarking AI Coding Agents Against Malicious Issue Requests

- HN: [49136144](https://news.ycombinator.com/item?id=49136144)
- Source: [arxiv.org](https://arxiv.org/abs/2607.20759)
- Score: 2
- Comments: 0
- Posted: 2026-08-01T16:59:18Z

## Translation

タイトル: IssueTrojanBench: 悪意のある問題リクエストに対する AI コーディング エージェントのベンチマーク
記事のタイトル: [2607.20759] IssueTrojanBench: 悪意のある問題リクエストに対する AI コーディング エージェントのベンチマーク
説明: arXiv 論文 2607.20759 の要約ページ: IssueTrojanBench: 悪意のある問題リクエストに対する AI コーディング エージェントのベンチマーク

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
[2026 年 7 月 22 日に提出]
タイトル: IssueTrojanBench: 悪意のある問題リクエストに対する AI コーディング エージェントのベンチマーク
要約: LLM を利用した AI コーディング エージェントは、現実世界のソフトウェア開発にますます統合されており、ローカル ファイルやツールに自律的にアクセスしてコードを生成、編集、実行します。コーディング エージェントは、LLM バックボーン (敵対的なプロンプト、汚染されたトレーニング データ、バックドア トリガーによってモデルが安全でないコードや攻撃者が選択したコードを出力する原因となる可能性がある) と、ツール使用の自律性によって外部 API の誤用、データ漏洩、開発環境の永続的な侵害が可能になるエージェント アーキテクチャの両方からセキュリティ リスクを継承します。このペーパーでは、2 つの主要なモデル ファミリ (OpenAI GPT-5.3 Codex/GPT-5.4 および Anthropic Sonnet 4.6) を利用した、最先端のコーディング エージェント (Cursor、Claude Code、および Codex Desktop) に対する悪意のある問題リクエストの体系的な評価を示します。当社の新しいベンチマーク IssueTrojanBench には、4 つの新しい攻撃カテゴリ (つまり、問題に悪意のある命令として埋め込まれている)、6 つの配信ベクトル (PDF または問題のコメントなど) に基づいて構築され、摂動によってさらに強化された悪意のある問題が含まれています。私たちの結果は、導入されたままの最新のコーディング エージェントに重大な脆弱性を明らかにしました。つまり、IssueTrojanBench からの悪意のある問題の 66.5% が、コーディング エージェントのすべてのガードレール (エージェントおよび LLM レベル) に侵入しています。さらに分析したところ、拒否はエージェント フレームワークではなくほぼ完全に LLM によるものであり、GPT モデルは広範囲に脆弱であり、Sonnet 4.6 はより選択的でリスクを意識した影響の大きい行為のブロックを示していることが示されています。

イオン。私たちの評価では、現在のエージェント レベルの防御戦略では、コーディング エージェントに対する追加の保護が限定的であることも強調されています。私たちの調査結果は、AI コーディング エージェントを保護するために、エージェント レベルおよびモデル レベルのより強力な安全メカニズムが緊急に必要であることを浮き彫りにしています。
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

Abstract page for arXiv paper 2607.20759: IssueTrojanBench: Benchmarking AI Coding Agents Against Malicious Issue Requests

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Cryptography and Security
[Submitted on 22 Jul 2026]
Title: IssueTrojanBench: Benchmarking AI Coding Agents Against Malicious Issue Requests
Abstract: AI coding agents powered by LLMs are increasingly integrated into real-world software development, where they generate, edit, and execute code with autonomous access to local files and tools. Coding agents inherit security risks from both the LLM backbone, where adversarial prompts, poisoned training data, and backdoor triggers can cause models to emit insecure or attacker-chosen code, and their agentic architecture, where tool-using autonomy enables induced misuse of external APIs, data exfiltration, and persistent compromise of development environments. This paper presents a systematic evaluation of malicious issue requests against state-of-the-art coding agents (Cursor, Claude Code, and Codex Desktop), powered by two major model families (OpenAI GPT-5.3 Codex/GPT-5.4 and Anthropic Sonnet 4.6). Our novel benchmark IssueTrojanBench contains malicious issues that are constructed based on four novel attack categories (i.e., embedded as malicious instructions in issues), six delivery vectors (e.g., PDF, or issue comment), and further augmented by perturbations. Our results reveal critical vulnerabilities in the as-deployed modern coding agents, i.e., 66.5% of the malicious issues from IssueTrojanBench penetrate all the guardrails (agent- and LLM-level) of coding agents. Our further analysis shows that rejection is almost entirely from LLMs rather than the agent frameworks, with GPT models broadly vulnerable and Sonnet 4.6 exhibiting more selective, risk-aware blocking of high-impact actions. Our evaluation also highlights that the current agent-level defense strategy offers limited additional protection for coding agents. Our findings highlight the urgent need for stronger agent- and model-level safety mechanisms to protect AI coding agents.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
