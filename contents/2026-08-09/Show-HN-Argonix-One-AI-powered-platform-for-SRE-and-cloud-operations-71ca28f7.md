---
source: "https://argonix.io/"
hn_url: "https://news.ycombinator.com/item?id=49230251"
title: "Show HN: Argonix – One AI-powered platform for SRE and cloud operations"
article_title: "Argonix | Argos — Autonomous AI Agent for Monitoring, Security & FinOps"
author: "pydavid"
captured_at: "2026-08-09T11:21:10Z"
capture_tool: "hn-digest"
hn_id: 49230251
score: 1
comments: 0
posted_at: "2026-08-09T10:50:18Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Argonix – One AI-powered platform for SRE and cloud operations

- HN: [49230251](https://news.ycombinator.com/item?id=49230251)
- Source: [argonix.io](https://argonix.io/)
- Score: 1
- Comments: 0
- Posted: 2026-08-09T10:50:18Z

## Translation

タイトル: Show HN: Argonix – SRE およびクラウド運用のための AI を活用した 1 つのプラットフォーム
記事タイトル: Argonix | Argos — 監視、セキュリティ、FinOps のための自律型 AI エージェント
説明: Argos は、Datadog、PagerDuty、Wiz に代わる、スタックを監視し、クラウドを保護し (CSPM、ISO 27001 / SOC 2 / NIS2 / CIS、MITRE ATT&CK)、クラウド支出を制御 (Kubernetes のコスト割り当て、チャージバック PDF、予算、適正サイズ) する自律型 AI エージェントです。 800 以上のツール、50 個のコネクタ。セルフホ
[切り捨てられた]
HN テキスト: 皆さん、こんにちは。 https://argonix.io を構築しました。 Argonix は、SRE、セキュリティ、クラウド運用のための AI を活用したプラットフォームです。このアイデアは、監視、セキュリティ、Finops、ヘルスチェック、ワークフロー、インシデント調査など、通常はさまざまなツールに分散しているものをまとめて接続することです。これを使用してできることはいくつかあります。 コネクタ: Argonix を既存のスタック (Kubernetes、AWS、GCP、Azure、独自の MCP サーバーなど) に接続します。 ワークフロー: コネクタを使用して自動化フローを構築します。ヘルスノートブック: プラットフォームの健全性の毎日の概要を取得します。 Argonix はレポートを Slack や Confluence などに公開し、注意が必要な場合にチケットを作成できます。定期的なジョブ: 定期的なチェック、分析、自動化をスケジュールします。セキュリティ: コネクタのインフラストラクチャ、コードもスキャンし (新機能)、ネットワーク ログを分析し、脆弱性と脅威を探します。 Falco、Wazuh、Wiz、SonarQube などのツールからのデータを消費することもできます。
FinOps – AWS、GCP、およびクラウド全体のコストを分析します。
アズール。監視: 監視は UptimeKuma などと同様に機能し、さらに Kubernetes のイングレス/ゲートウェイを自動的に検出して監視し、合成テストを実行し、ログを分析します。アラート調査: Datadog、Grafana などから Argonix にアラートを送信できます。 AIが調査できる

アラートを読み、KB を読み、推奨事項を提供したり、PR を提案したりすることもできます。会話型 AI: 会話型インターフェイスを通じてインフラストラクチャや接続されたツールと対話します。内部的には、Argonix は Python/DRF、Vue.js、PostgreSQL で構築されています。 Argonix は、SaaS バージョンとセルフホスト バージョンの両方として利用できます。使用する AI モデルも選択します。独自の AI プロバイダー/アカウントに接続でき、セルフホスト バージョンではローカル モデルを実行できます。 Qwenを使用してローカルでテストしています。無料でアカウントを作成して、プレイを始めることができます。まだ進行中の作業であり、間違いなくいくつかのバグがあるため、現場で働いている人々からのフィードバックを非常に歓迎します。あなたがどう思うか、何を変えるつもりか、そしてそれが日々の仕事に真に役立つようになるかどうかをぜひ聞きたいです。ありがとう、
デビッド。

記事本文:
アルゴニクス | Argos — 監視、セキュリティ、FinOps のための自律型 AI エージェント

## Original Extract

Meet Argos, the autonomous AI agent that monitors your stack, secures your cloud (CSPM, ISO 27001 / SOC 2 / NIS2 / CIS, MITRE ATT&CK) and controls cloud spend (Kubernetes cost allocation, chargeback PDF, budgets, rightsizing) — replacing Datadog, PagerDuty and Wiz. 800+ tools, 50 connectors. Self-ho
[truncated]

Hi everybody, I built https://argonix.io . Argonix is an AI-powered platform for SRE, security and cloud operations. The idea is to bring and connect together things that are usually spread across many different tools: monitoring, security, Finops, health checks, workflows and incident investigation. A few things you can do with it: Connectors: Connect Argonix to your existing stack: Kubernetes, AWS, GCP, Azure, your own MCP servers, etc. Workflows: Build automation flows using the connectors. Health notebooks: Get a daily overview of the health of your platforms. Argonix can publish the report to Slack, Confluence, etc., and create tickets when something needs attention. Periodic jobs: Schedule recurring checks, analysis and automation. Security: Scan of your connectors infras, your code also (new feature), analyse network logs, and look for vulnerabilities and threats. It can also consume data from tools such as Falco, Wazuh, Wiz, SonarQube, etc.
FinOps – Analyse cloud costs across AWS, GCP and
Azure. Monitoring: Monitor works just like UptimeKuma or others, plus automatically discover and monitor Kubernetes ingresses/gateways, run synthetic tests, and analyse logs. Alert investigatio: You can send alerts from Datadog, Grafana and others to Argonix. The AI can investigate the alert, read the KB and provide recommendations or even propose a PR. Conversational AI: Talk to your infrastructure and connected tools through a conversational interface. Under the hood, Argonix is built with Python/DRF, Vue.js and PostgreSQL. Argonix is available as both a SaaS and a self-hosted version. You also choose which AI model you want to use. You can connect your own AI provider/account, and with the self-hosted version you can run a local model. I've been testing it locally with Qwen. You can create an account for free and start playing with it. It's still a work in progress and there are definitely some bugs, so I'd really appreciate feedback from people working in the field. I would love to hear what you think, what you would change, and what would make it genuinely useful in your day-to-day work. Thanks,
David.

Argonix | Argos — Autonomous AI Agent for Monitoring, Security & FinOps
