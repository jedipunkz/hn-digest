---
source: "https://blog.virustotal.com/2026/06/crowdsourced-ai-knostic.html"
hn_url: "https://news.ycombinator.com/item?id=48455863"
title: "Crowdsourced AI and= Knostic"
article_title: "Crowdsourced AI += Knostic ~ VirusTotal Blog"
author: "jruohonen"
captured_at: "2026-06-09T04:29:07Z"
capture_tool: "hn-digest"
hn_id: 48455863
score: 2
comments: 0
posted_at: "2026-06-09T03:12:48Z"
tags:
  - hacker-news
  - translated
---

# Crowdsourced AI and= Knostic

- HN: [48455863](https://news.ycombinator.com/item?id=48455863)
- Source: [blog.virustotal.com](https://blog.virustotal.com/2026/06/crowdsourced-ai-knostic.html)
- Score: 2
- Comments: 0
- Posted: 2026-06-09T03:12:48Z

## Translation

タイトル: クラウドソーシング AI と = Knostic
記事タイトル: クラウドソーシング AI += Knostic ~ VirusTotal Blog
説明: VirusTotal のクラウドソーシング AI ラインナップに新しいスペシャリスト、Knostic の AgentMesh Agentic Security Supply Chain Reputation Engine を追加します。 ...

記事本文:
クラウドソーシング AI += Knostic ~ VirusTotal ブログ
人気の投稿
VirusTotal からのアップデート
私たちの目標はシンプルです。Web 上での安全を確保することです。そして、私たちは提供するサービスが継続的に改善されるよう懸命に取り組んできました。しかし、...
自動化から感染まで: OpenClaw AI エージェントのスキルがどのように武器化されているか
急速に成長しているパーソナル AI エージェント エコシステムが、マルウェアの新しい配信チャネルになりました。ここ数日間、VirusTotal は...
VTPRACTITIONERS{SEQRITE}: UNG0002、Silent Lynx、DragonClone を追跡しています
はじめに VirusTotal (VT) に参加して最も良かったことの 1 つは、コミュニティがツールを使用してウイルスを追跡する驚くべき方法をすべて見ることができることです。
▼
2026年
(5)
▼
2026年6月
(1)
クラウドソーシング AI += Knostic
►
2025年
(17)
►
2025年12月
(1)
►
2024年
(18)
►
2024年11月
(1)
►
2023年
(35)
►
2023年12月
(5)
►
2022年
(23)
►
2022年12月
(1)
►
2021年
(19)
►
2021年12月
(2)
►
2020年
(15)
►
2020年12月
(2)
►
2019年
(19)
►
2019年11月
(2)
►
2018年
(10)
►
2018年12月
(1)
►
2017年
(13)
►
2017年11月
(1)
►
2016年
(6)
►
2016 年 11 月
(2)
►
2015年
(8)
►
2015 年 11 月
(1)
►
2014年
(24)
►
2014 年 12 月
(1)
►
2013年
(26)
►
2013 年 11 月
(2)
►
2012年
(26)
►
2012 年 12 月
(1)
►
2011年
(1)
►
2011 年 12 月
(1)
VirusTotal のクラウドソーシング AI ラインナップに新しいスペシャリストを追加します。それは、Knostic の AgentMesh Agentic Security Supply Chain Reputation Engine です。私たちは Visual Studio Code 拡張機能 (.VSIX) ファイルの分析で彼らと提携しています。これは、開発者、プラットフォーム エンジニア、セキュリティ チームが拡張機能のセキュリティ プロファイルをより深く理解し、拡張機能をインストールする前にサプライ チェーンの脅威を検出できるようにすることで、既存の Code Insight やその他の AI コントリビューターを補完します。
最近の GitHub データ侵害を脇に置いても、

悪意のある VS Code 拡張機能から始まり、IDE ベースの AI コーディング アシスタントや専用の開発者ツールの台頭により、Visual Studio Code 拡張機能は最新の開発ワークフローの中心となっています。しかし、これにより、サプライチェーン攻撃の主な標的にもなりました。悪意のある攻撃者が、密かにペイロードをダウンロードしたり、リモート コードを実行したり、資格情報を盗んだり、独自のソース コードや機密環境変数を黙って抽出したりする一見無害な拡張機能を公開していることが捕まりました。
.VSX に対するセカンドオピニオン: Knostic は、特に `.VSIX` パッケージに特化した AI 主導の分析ストリームを追加します。これにより、セキュリティ チームは拡張ファイルを独立して評価できるようになり、重大な脆弱性と意図的なバックドア動作の両方を特定するのに役立ちます。
明確な判定とリスク レベル : Knostic はファイルを分析し、検出されたリスク指標の詳細な説明とともに、リスク レベル (SAFE、MEDIUM、または CRITICAL など) と組み合わせた明確なスキャン判定 (BENIGN、SUSPICIOUS、または MALICIOUS) を割り当てます。
VT Intelligence のスキャン時のピボットと検索: セキュリティ アナリストは、新しくインデックス付けされた演算子を使用して、Knostic 結果全体を検索およびフィルターできるようになりました。
* `knostic_ai_verdict:悪意のある |疑わしい |良性です
* `knostic_ai_analysis: `
knostic_ai_verdict:悪意のある |疑わしい |良性の
knostic_ai_analysis:[キーワード]
Knostic の AgentMesh が実際にどのように機能するかを説明するために、分析された実際の VS Code 拡張機能をいくつか見てみましょう。
Knostic は、エージェントとコーディング アシスタント、およびその周辺サプライ チェーン (拡張機能、スキル、MCP サーバーなど) を検出して防御します。 Knostic は、Kirin プラットフォームと AgentMesh 脅威インテリジェンス エンジンを通じて、企業が AI エージェントと開発者拡張機能が内部システムとどのように対話するかを管理できるように支援します。

プライベートデータを管理し、ソースにおけるサプライチェーンのリスクを中和します。
VT クラウドソーシング AI は、多くのファイル タイプにわたって動作を説明し、判断を提供する独立した AI ソリューションを集約することで、なじみのないコードをより速く理解し、新たな脅威をより早く発見できるようにします。コミュニティに役立つ AI ソリューションを構築されている場合は、ぜひご意見をお待ちしています。

## Original Extract

We’re adding a new specialist to VirusTotal’s Crowdsourced AI lineup: Knostic 's AgentMesh Agentic Security Supply Chain Reputation Engine. ...

Crowdsourced AI += Knostic ~ VirusTotal Blog
Popular Posts
An update from VirusTotal
Our goal is simple: to help keep you safe on the web. And we’ve worked hard to ensure that the services we offer continually improve. But as...
From Automation to Infection: How OpenClaw AI Agent Skills Are Being Weaponized
The fastest-growing personal AI agent ecosystem just became a new delivery channel for malware. Over the last few days, VirusTotal has detec...
VTPRACTITIONERS{SEQRITE}: Tracking UNG0002, Silent Lynx and DragonClone
Introduction One of the best parts of being at VirusTotal (VT) is seeing all the amazing ways our community uses our tools to hunt down thr...
▼
2026
(5)
▼
June 2026
(1)
Crowdsourced AI += Knostic
►
2025
(17)
►
December 2025
(1)
►
2024
(18)
►
November 2024
(1)
►
2023
(35)
►
December 2023
(5)
►
2022
(23)
►
December 2022
(1)
►
2021
(19)
►
December 2021
(2)
►
2020
(15)
►
December 2020
(2)
►
2019
(19)
►
November 2019
(2)
►
2018
(10)
►
December 2018
(1)
►
2017
(13)
►
November 2017
(1)
►
2016
(6)
►
November 2016
(2)
►
2015
(8)
►
November 2015
(1)
►
2014
(24)
►
December 2014
(1)
►
2013
(26)
►
November 2013
(2)
►
2012
(26)
►
December 2012
(1)
►
2011
(1)
►
December 2011
(1)
We’re adding a new specialist to VirusTotal’s Crowdsourced AI lineup: Knostic 's AgentMesh Agentic Security Supply Chain Reputation Engine. We are partnering with them to analyze Visual Studio Code extension (.VSIX) files. This complements our existing Code Insight and other AI contributors by helping developers, platform engineers, and security teams better understand the security profile of extensions and detect supply-chain threats before installing them.
Even putting aside the recent GitHub data breach, resulting from a malicious VS Code extensions, with the rise of IDE-based AI coding assistants and specialized developer tools, Visual Studio Code extensions have become central to modern development workflows. However, this has also made them prime targets for supply-chain attacks. Malicious actors have been caught publishing seemingly benign extensions that secretly download payloads, perform remote code execution, steal credentials, or silently exfiltrate proprietary source code and sensitive environment variables.
Second opinion for .VSX : Knostic adds a specialized AI-driven analysis stream specifically for `.VSIX` packages. This provides security teams with an independent assessment of extension files, helping to identify both critical vulnerabilities and deliberate backdoor behaviors.
Clear Verdicts and Risk Levels : Knostic analyzes files and assigns a clear scan verdict (BENIGN, SUSPICIOUS, or MALICIOUS) coupled with a risk level (such as SAFE, MEDIUM, or CRITICAL) along with detailed descriptions of detected risk indicators.
Pivot and Search at Scane in VT Intelligence : Security analysts can now search and filter across Knostic results using newly indexed operators:
* `knostic_ai_verdict:malicious | suspicious | benign`
* `knostic_ai_analysis: `
knostic_ai_verdict:malicious | suspicious | benign
knostic_ai_analysis:[keywords]
To illustrate how Knostic’s AgentMesh works in practice, let’s explore some real VS Code extensions that have been analyzed::
Knostic discovers and defends agents and coding assistants, as well as their peripheral supply chain (e.g., extensions, skills, MCP servers). Through its Kirin platform and AgentMesh threat intelligence engine, Knostic helps enterprises govern how AI agents and developer extensions interact with internal systems and private data, neutralizing supply-chain risks at the source.
VT Crowdsourced AI is about aggregating independent AI solutions that explain behavior and provide judgments across many file types, helping you understand unfamiliar code faster and spot novel threats sooner. If you build AI solutions that can help the community, we want to hear from you.
