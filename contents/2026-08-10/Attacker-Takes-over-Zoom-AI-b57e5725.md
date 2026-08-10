---
source: "https://www.promptarmor.com/resources/attacker-takes-over-zoom-ai"
hn_url: "https://news.ycombinator.com/item?id=49248629"
title: "Attacker Takes over Zoom AI"
article_title: "Attacker Takes Over Zoom AI | PromptArmor"
author: "hackerBanana"
captured_at: "2026-08-10T19:47:57Z"
capture_tool: "hn-digest"
hn_id: 49248629
score: 3
comments: 0
posted_at: "2026-08-10T19:38:42Z"
tags:
  - hacker-news
  - translated
---

# Attacker Takes over Zoom AI

- HN: [49248629](https://news.ycombinator.com/item?id=49248629)
- Source: [www.promptarmor.com](https://www.promptarmor.com/resources/attacker-takes-over-zoom-ai)
- Score: 3
- Comments: 0
- Posted: 2026-08-10T19:38:42Z

## Translation

タイトル: 攻撃者がZoom AIを乗っ取る
記事タイトル: 攻撃者がZoom AIを乗っ取る |プロンプトアーマー
説明: Zoom AI にはコーディング環境にネットワーク制限がないため、攻撃者が悪意のあるスキルを介して会議データやメッセージなどを窃取するための指揮統制を確立することができます。

記事本文:
攻撃者がZoom AIを乗っ取る | PromptArmor ソリューション 業界 パートナー リソース 価格 デモを予約 脅威 Intel 購読 攻撃者が Zoom AI を乗っ取る
Zoom AI にはコーディング環境にネットワーク制限がないため、攻撃者が悪意のあるスキルを介して会議データやメッセージなどを窃取するための指揮統制を確立することができます。
攻撃者は被害者のZoom AIを乗っ取り、Zoomおよび接続されたサービス全体からデータを窃取する
Zoom の主力 AI 機能は ZoomMate です。ZoomMate は、ユーザーの Zoom アカウントと接続されているサービス (OneDrive、Google、コネクタなど) 全体からのデータを操作するエージェント チャットボットです。このエージェントには、無制限の HTTPS ネットワーク アクセスが可能な環境が与えられており、ユーザーまたは管理者レベルでロックダウンするための設定は行われていないようです。
この記事では、悪意のあるスキル (またはプロンプト インジェクション) が Zoom のエージェントを操作して攻撃者のサーバーに接続し、攻撃者が独自のコマンドを発行して被害者のテナントからデータを抜き出す方法を検討します。
このリスクをさらに悪化させるのは、ユーザーがエージェント上で「停止」をクリックして Zoom を閉じた場合でも、攻撃者の接続がアクティブなままになる可能性があることです。さらに、ユーザーに対する最終的なチャット出力は完全に正常に見えます。
ZoomMate のインターネット アクセスを備えた広範な機能を備えた環境は、意図された機能であると思われます。 Zoom に公開される予定どおりに動作しない特定の脆弱性やプログラム コンポーネントはないようです。この記事は、厳密なネットワーク サンドボックスを使用せずにエージェント チャットボットを利用することでリスクを認識していない可能性があるユーザーに知らせるために公開しています。
1. 被害者は、その週の会議の報告を求めます。
被害者はZoomMateに週次会議レポートの作成を依頼する
2. 犠牲者

mさんはZoomにアップロードしたスキルを使用しています
注: この記事の攻撃は、スキルを使用せずに、間接的なプロンプト インジェクション (電子メールなど、Zoom が取り込む隠し命令など) を介して実行することもできます。
スキルは通常、オンライン マーケットプレイスを通じて配布され、Zoom 内のユーザー間で共有できます。以前の調査では、攻撃者が悪意のあるスキルをこれらのオンライン レジストリにアップロードしていることがわかっています。
インストールされた Zoom スキルの中にアップロードされた週次会議レポート スキルがリストされる 注: Zoom は、スキルをアップロードするときにユーザーに警告を表示しますが、それがリスクを適切に知らせているとは考えていません。警告: 「このスキルは Zoom の公式カタログからのものではなく、Zoom によって検証されていません。インストールする前に、このスキルの作成者を信頼していることを確認してください。」
3. Zoom AI が悪意のあるスキルからコードを実行する
ユーザーが「停止」ボタンをクリックするかZoomを閉じても、攻撃はZoomのサーバー上のエージェント環境で実行され続けるため、停止しません。
悪意のあるスキルのコードは、数秒ごとに攻撃者のサーバーにネットワーク リクエスト (HTTPS) を送信し、サーバーにコマンドの実行を要求します。攻撃者がコマンドを送信すると、スクリプトが Zoom の環境でそのコマンドを実行し、結果を攻撃者のサーバーに送り返します。
ZoomMateがスキルから悪意のあるスクリプトを実行する
4. 攻撃者はコマンドを送信して、会議の記録、メッセージ、コネクタからのデータなどを窃取します。
攻撃者は、Zoom および接続されているアプリ全体からデータを抜き出すコマンドを実行します。 注: 攻撃者は、接続されているデータ ソースを含め、ZoomMate がアクセスできる Zoom 全体からのデータをターゲットにすることができます。攻撃者が Zoom のブラウザ統合を悪用して他の Web サイトからデータを窃取できるかどうかは確認されていません。
5. 正常に見えるレポートがユーザーに配信されます

r、攻撃者は接続を維持します
レポート出力とチャットは正常に見えますが、攻撃者は ZoomMate が停止した後もコマンドを実行し続けることができます。
PromptArmor 脅威インテリジェンス
あなたの組織はベンダーの AI から保護されていますか?
PromptArmor は、ベンダー、スキル、プラグイン、コネクタ、MCP サーバー、モデルなどのサードパーティ AI のポートフォリオ全体を継続的に監視します。
このような脆弱性や変更を検出し、インシデントになる前にリスクを表面化します。
サインアップすると、次の変更に関するアラートを 1 か月間無料で受け取ることができます。
Claude Chatエンタープライズ Microsoft Copilot 用 GPT
新たな脅威に関する新しいインテリジェンスを備えたベンダーの AI によるリスクを評価および監視します。
プロンプトアーマー © 2026 全著作権所有

## Original Extract

Zoom AI lacks network restrictions in its coding environment, enabling attackers to establish command-and-control to exfiltrate meeting data, messages, and more via a malicious Skill.

Attacker Takes Over Zoom AI | PromptArmor Solutions Industries Partners Resources Pricing Book a Demo Threat Intel Subscribe Attacker Takes Over Zoom AI
Zoom AI lacks network restrictions in its coding environment, enabling attackers to establish command-and-control to exfiltrate meeting data, messages, and more via a malicious Skill.
Attacker takes over the victim’s Zoom AI, exfiltrating data from across Zoom and connected services Context
Zoom’s flagship AI feature is ZoomMate, an agentic chatbot that operates on data from across a user’s Zoom account and any connected services (e.g., OneDrive, Google, connectors, etc). The agent appears to have been given an environment with unrestricted HTTPS network access, with no user or admin-level configuration to lock it down.
In this article, we explore how a malicious Skill (or a prompt injection) can manipulate Zoom’s agent into connecting to an attacker’s server, allowing the attacker to issue their own commands and exfiltrate data from the victim’s tenant.
Further compounding this risk, the attacker’s connection can remain active even if the user clicks ‘stop’ on the agent and closes Zoom . Additionally, the final chat output to the user appears completely normal.
ZoomMate’s widely capable environment with internet access appears to be an intended functionality. There does not appear to be any specific vulnerability or programmatic component that is not working as intended to be disclosed to Zoom. We are publishing this article to inform users who may not be aware of the risk they are accepting by utilizing an agentic chatbot without strict network sandboxing.
1. The victim asks for a report on their meetings for the week
Victim asks ZoomMate to create a weekly meeting report
2. The victim is using a Skill they have uploaded to Zoom
Note: The attack in this article can also be conducted without a Skill, via indirect prompt injection (e.g., a hidden instruction Zoom ingests, such as an email).
Skills are typically distributed through online marketplaces and can be shared between users within Zoom; prior research shows that attackers are uploading malicious Skills to these online registries .
The uploaded weekly-meeting-report Skill listed among installed Zoom Skills Note: Zoom does offer users a warning when uploading a Skill, but we do not believe it adequately informs them of the risks. The warning: “This skill is not from Zoom's official catalog and hasn't been verified by Zoom. Make sure you trust this skill's creator before installing.”
3. Zoom AI runs code from the malicious Skill
Even if the user clicks the ‘stop’ button or closes Zoom, the attack does not stop because it continues to run in the agent’s environment on Zoom’s servers.
The code in the malicious Skill makes network requests (HTTPS) to an attacker’s server every few seconds, asking the server for commands to run. When the attacker sends a command, the script executes it in Zoom’s environment and sends the results back to the attacker’s server.
ZoomMate executing the malicious script from the Skill
4. The attacker sends commands to exfiltrate meeting transcripts, messages, data from connectors, and more
Attacker runs commands to exfiltrate data from across Zoom and connected apps Note: The attacker can target data from across Zoom that ZoomMate has access to, including connected data sources. It has not been confirmed whether the attacker can exploit Zoom’s browser integration to exfiltrate data from other websites.
5. A normal-looking report is delivered to the user, and the attacker stays connected
Report output and chat look normal while the attacker can keep running commands even after ZoomMate stops
PromptArmor Threat Intelligence
Is your organization protected from AI in vendors?
PromptArmor continuously monitors across your portfolio of third party AI in vendors, skills, plugins, connectors, MCP servers, models and more.
We detect vulnerabilities and changes like this, surfacing risk before it becomes an incident.
Sign up for a free month of alerts on changes to:
Claude ChatGPT for Enterprise Microsoft Copilot
Assess and monitor risk from AI in vendors with novel intelligence on emerging threats.
PromptArmor © 2026 All rights reserved
