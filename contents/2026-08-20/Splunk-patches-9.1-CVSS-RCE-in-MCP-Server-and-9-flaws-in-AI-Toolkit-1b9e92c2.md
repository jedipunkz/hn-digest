---
source: "https://cyberupdates365.com/splunk-mcp-server-rce-patch/"
hn_url: "https://news.ycombinator.com/item?id=49372596"
title: "Splunk patches 9.1 CVSS RCE in MCP Server and 9 flaws in AI Toolkit"
article_title: "Critical Splunk MCP Server RCE & AI Toolkit Patches"
image: "https://cyberupdates365.com/wp-content/uploads/2026/08/splunk-mcp-server-rce-vulnerability.jpg"
author: "udayhero"
captured_at: "2026-08-20T10:20:01Z"
capture_tool: "hn-digest"
hn_id: 49372596
score: 1
comments: 0
posted_at: "2026-08-20T10:10:49Z"
tags:
  - hacker-news
  - translated
---

# Splunk patches 9.1 CVSS RCE in MCP Server and 9 flaws in AI Toolkit

- HN: [49372596](https://news.ycombinator.com/item?id=49372596)
- Source: [cyberupdates365.com](https://cyberupdates365.com/splunk-mcp-server-rce-patch/)
- Score: 1
- Comments: 0
- Posted: 2026-08-20T10:10:49Z

## Translation

タイトル: Splunk が MCP サーバーの 9.1 CVSS RCE と AI Toolkit の 9 つの欠陥をパッチ
記事のタイトル: クリティカルな Splunk MCP サーバー RCE および AI ツールキット パッチ
説明: plunk は、重大な Splunk mcp サーバー rce (CVE-2026-76404) を含む 17 件の欠陥を修正します。 AI Toolkit、VictorOps、Kafka コンポーネントをすぐにアップグレードします。

記事本文:
重要な Splunk MCP サーバー RCE および AI ツールキット パッチ
ホーム
Splunk が AI ツールキット、Kafka アプリ全体にわたる重大な MCP サーバー RCE およびその他の 16 のセキュリティ欠陥にパッチを適用
Splunk は、Splunk MCP Server、Splunk AI Toolkit、Splunk Connect for Kafka など、いくつかのアプリやアドオンに影響を与える 17 件の脆弱性に対するセキュリティ アップデートをリリースしました。
CVE-2026-76404 として追跡される最も深刻な問題は、CVSS スコアが 9.1 の重大な Splunk mcp サーバー rce (リモート コード実行) の脆弱性です。 2026 年 8 月のアドバイザリでは、Enterprise Security Cloud および Splunk On-Call (VictorOps) 向けの Cisco Talos Intelligence も対象としています。
影響を受けるコンポーネントを使用している組織は、特に管理インターフェイス、REST API、または AI モデル管理機能が信頼できないユーザーやネットワークに公開されている場合、アップグレードを優先する必要があります。
完全な脆弱性の概要（モバイル フレンドリー） Cisco Talos Intelligence
FAQ: CyberUpdates365 デスクによって報告された Splunk の脆弱性
CVE-2026-76404 は、1.2.1 より前の Splunk MCP サーバー アプリのバージョンに影響します。 Splunkによると、管理者ロールを持つ認証されたユーザーは、基盤となるオペレーティングシステム上で任意のコマンドを実行できる可能性があるという。
この欠陥はアプリの認証情報管理コンポーネントに存在します。入力検証が欠落していると、提供されたコンテンツが予期されたタイプであることを確認せずに、保存されたデータの安全でない逆シリアル化が可能になります。
これにより、悪意のあるシリアル化されたオブジェクトがコマンドを実行するためのパスが作成されます。この脆弱性は、CWE-502 (信頼できないデータの逆シリアル化) として分類されます。管理者は、Splunk MCP サーバー アプリをすぐにバージョン 1.2.1 にアップグレードする必要があります。
パッチ適用が完了するまで、Splunk はアプリを無効にするか削除することをお勧めします。 Splunk AI Toolkit は、いくつかの高重大度のアクセス制御およびコード実行を含む 9 つの脆弱性に対する修正を受け取りました。

法律。
最も深刻な CVE-2026-76395 は、CVSS スコアが 8.8 であり、権限のあるユーザーが、悪意のある疎行列データを含む細工されたモデルをロードすることにより、Splunk サーバー上で任意のコードを実行できる可能性があります。この問題は、埋め込まれた pickle コンテンツの安全でない逆シリアル化に起因します。
AI Toolkit のその他の欠陥により、権限の低いユーザーがシステム レベルの権限で検索を実行したり、適切な承認なしにコンテナーや接続を管理したり、他のユーザーの実験履歴にアクセスまたは削除したり、検索所有者の権限で実行されるスケジュールされた検索を変更したりできる可能性があります。
Splunk AI Toolkit バージョン 6.0.0 より前のバージョンはこれらの問題のほとんどの影響を受けますが、2 つの欠陥が 6.0.1 より前のバージョンに影響します。 5.7 リリースラインを使用しているユーザーはバージョン 6.0.0 にアップグレードする必要があり、すでにバージョン 6.0 を使用しているユーザーは 6.0.1 に移行する必要があります。
完全な脆弱性の概要 (モバイル対応)
CVE-2026-76389 (高 – CVSS 8.8)
問題: SSRF はトークンを公開する可能性があります。修正バージョン: 1.0.3
CVE-2026-76390 (中 – CVSS 5.3)
問題: OpenAPI 仕様の公開。修正バージョン: 1.0.3
CVE-2026-76391 (高 - CVSS 8.3)
問題: システムレベルの検索アクセス。修正バージョン: 6.0.0
CVE-2026-76392 (中 – CVSS 5.4)
問題: 予測可能な認証情報。修正バージョン: 6.0.0
CVE-2026-76393 (中 – CVSS 5.9)
問題: 競合状態によるモデルの置き換え。修正バージョン: 6.0.0
CVE-2026-76394 (高 - CVSS 8.3)
問題: 不正なコンテナ制御。修正バージョン: 6.0.0
CVE-2026-76395 (高 - CVSS 8.8)
問題: 安全でない pickle 逆シリアル化を介した悪意のあるモデル RCE。修正バージョン: 6.0.0
CVE-2026-76396 (高 - CVSS 7.5)
問題: 安全でないモデルの逆シリアル化。修正バージョン: 6.0.0
CVE-2026-76397 (高 - CVSS 8.1)
問題: 実験データへの不正アクセス。修正バージョン: 6.0.0
CVE-2026-76398 (中 – CVSS 4.3)
問題: アンノート

恐ろしい履歴の削除。修正バージョン: 6.0.1
CVE-2026-76399 (高 - CVSS 8.1)
問題: スケジュールされた検索操作。修正バージョン: 6.0.1
CVE-2026-76400 (中 – CVSS 5.9)
問題: イベント配信の中断。修正バージョン: 2.2.7
CVE-2026-76401 (中 – CVSS 5.9)
問題: コネクタ ワーカーの DoS。修正バージョン: 2.2.7
CVE-2026-76402 (高 - CVSS 8.2)
問題: HEC を介した認証情報の漏洩。修正バージョン: 2.2.7
CVE-2026-76403 (高 - CVSS 7.4)
問題: データの傍受/改ざん。修正バージョン: 2.2.7
CVE-2026-76404 (重大 - CVSS 9.1)
問題: OS コマンドの実行 (RCE)。修正バージョン: 1.2.1
CVE-2026-76405 (中 – CVSS 4.3)
問題: Splunk On-Call で部分的な API キーが公開される。修正バージョン: 1.0.43
Splunk Connect for Kafka バージョン 2.2.7 は 4 つの脆弱性に対処します。これらには、SSRF の欠陥 CVE-2026-76402 (評価 8.2) が含まれており、これにより、認証されていない攻撃者が Kafka Connect REST API にアクセスして、安全でない HTTP イベント コレクター エンドポイントを構成し、コネクタの認証資格情報を取得できる可能性があります。
Kafka の修正は、Kerberos 認証された HTTP Event Collector デプロイメントにおける 2 つのサービス拒否の問題と不適切な証明書の検証にも対処します。
Splunk は、Kafka Connect REST API アクセスを信頼できる管理システムおよびネットワークに制限し、HTTP Event Collector エンドポイントに安全なトランスポートを強制し、有限の再試行制限を適用することを推奨しています。
追加の修正には、Cisco Talos Intelligence for Enterprise Security Cloud の SSRF 脆弱性と Splunk On-Call (VictorOps) での部分的な API キーの開示が含まれます。更新されたバージョンは、Talos アプリの場合は 1.0.3、Splunk On-Call の場合は 1.0.43 です。
組織は、インストールされている Splunk アプリのインベントリを作成し、関連するアップデートを適用し、管理 API へのアクセスを制限し、管理者および権限の役割の割り当てを確認する必要があります。

不必要な特権暴露の兆候。
関連リソース: クラウド テレメトリとログ管理パイプラインがネットワークを RCE 脅威にさらしていないことを確認します。詳細については、「Agentic SOC vs Traditional SOC 2026 Guide」を参照してください。
Splunk MCP サーバー RCE とは何ですか?
これは CVE-2026-76404 として追跡されており、Splunk MCP サーバー アプリの安全でない逆シリアル化によって引き起こされる重大なリモート コード実行の脆弱性 (CVSS 9.1) です。これにより、認証された管理者は任意の OS コマンドを実行できます。
Splunk AI ツールキットには脆弱性がありますか?
はい、Splunk は、CVE-2026-76395 を含む AI ツールキットの 9 つの脆弱性にパッチを適用しました。これにより、安全でない pickle データを含む悪意を持って作成された AI モデルを介して任意のコードが実行される可能性があります。
CyberUpdates365 デスクによる報告
企業のセキュリティ、連邦政府の AI 指令、IT インフラストラクチャの将来に関する最新の洞察を提供します。テクノロジーが企業の状況をどのように再形成しているかに関する最新情報を毎日入手してください。
この記事は面白かったですか? Twitter と LinkedIn でフォローして、当社が投稿する限定コンテンツをさらに読んでください。 Cyber​​Updates365 GitHub リポジトリで未加工の修復スクリプト アーカイブを表示することもできます。
Uday Patil Uday Patil は、サイバーセキュリティ研究者、DevSecOps エンジニア、Cyber​​Updates365 の創設者です。脅威インテリジェンスとゼロデイ脆弱性分析を専門とする Uday は、複雑なサイバー脅威を分析して実用的な洞察を得ることに専念しています。彼の使命は、世界中の開発者とセキュリティ チームに、迅速なアラート、修復スクリプト、実践的なガイダンスを提供して、進化する脅威の状況の先を行くことを支援することです。
Cyber​​ Updates 365 世界的なサイバー脅威、エンタープライズ セキュリティ、新たな AI テクノロジーに関する最新のインテリジェンスを提供します。 IT プロフェッショナルやテクノロジー愛好家に検証済みのニュースと実用的な情報を提供します。

デジタル戦場で優位に立つために。
記事をライブで検索するには、少なくとも 2 文字を入力してください...

## Original Extract

plunk fixes 17 flaws, including a critical splunk mcp server rce (CVE-2026-76404). Upgrade AI Toolkit, VictorOps, and Kafka components immediately.

Critical Splunk MCP Server RCE & AI Toolkit Patches
Home
Splunk Patches Critical MCP Server RCE and 16 Other Security Flaws Across AI Toolkit, Kafka Apps
Splunk has released security updates for 17 vulnerabilities affecting several apps and add-ons, including Splunk MCP Server, Splunk AI Toolkit, and Splunk Connect for Kafka.
The most severe issue, tracked as CVE-2026-76404 , is a critical splunk mcp server rce (remote code execution) vulnerability with a CVSS score of 9.1. The August 2026 advisory also covers Cisco Talos Intelligence for Enterprise Security Cloud and Splunk On-Call (VictorOps).
Organizations using the affected components should prioritize upgrades, especially where administrative interfaces, REST APIs, or AI model-management features are exposed to untrusted users or networks.
Complete Vulnerability Summary (Mobile-Friendly) Cisco Talos Intelligence
FAQ: Splunk Vulnerabilities Reported by CyberUpdates365 Desk
CVE-2026-76404 affects Splunk MCP Server app versions earlier than 1.2.1. Splunk said an authenticated user with the admin role could execute arbitrary commands on the underlying operating system.
The flaw exists in the app’s credential-management component. Missing input validation allows unsafe deserialization of stored data without confirming that the supplied content has the expected type.
This creates a path for malicious serialized objects to execute commands. The vulnerability is categorized as CWE-502, or deserialization of untrusted data. Administrators should upgrade the Splunk MCP Server app to version 1.2.1 immediately.
Until patching is complete, Splunk recommends disabling or removing the app. Splunk AI Toolkit received fixes for nine vulnerabilities, including several high-severity access-control and code-execution flaws.
The most serious, CVE-2026-76395, carries a CVSS score of 8.8 and could allow a user with the power role to run arbitrary code on a Splunk server by loading a crafted model containing malicious sparse matrix data. The issue stems from unsafe deserialization of embedded pickle content.
Other AI Toolkit flaws could allow lower-privileged users to perform searches with system-level privileges, manage containers and connections without proper authorization, access or delete other users’ experiment history, and modify scheduled searches that run under the search owner’s permissions.
Splunk AI Toolkit versions below 6.0.0 are affected by most of these issues, while two flaws affect versions below 6.0.1. Users on the 5.7 release line should upgrade to version 6.0.0, while users already on version 6.0 should move to 6.0.1.
Complete Vulnerability Summary (Mobile-Friendly)
CVE-2026-76389 (High – CVSS 8.8)
Issue: SSRF may expose tokens. Fixed Version: 1.0.3
CVE-2026-76390 (Medium – CVSS 5.3)
Issue: OpenAPI spec exposure. Fixed Version: 1.0.3
CVE-2026-76391 (High – CVSS 8.3)
Issue: System-level search access. Fixed Version: 6.0.0
CVE-2026-76392 (Medium – CVSS 5.4)
Issue: Predictable credentials. Fixed Version: 6.0.0
CVE-2026-76393 (Medium – CVSS 5.9)
Issue: Model replacement via race condition. Fixed Version: 6.0.0
CVE-2026-76394 (High – CVSS 8.3)
Issue: Unauthorized container control. Fixed Version: 6.0.0
CVE-2026-76395 (High – CVSS 8.8)
Issue: Malicious model RCE via unsafe pickle deserialization. Fixed Version: 6.0.0
CVE-2026-76396 (High – CVSS 7.5)
Issue: Unsafe model deserialization. Fixed Version: 6.0.0
CVE-2026-76397 (High – CVSS 8.1)
Issue: Unauthorized experiment data access. Fixed Version: 6.0.0
CVE-2026-76398 (Medium – CVSS 4.3)
Issue: Unauthorized history deletion. Fixed Version: 6.0.1
CVE-2026-76399 (High – CVSS 8.1)
Issue: Scheduled search manipulation. Fixed Version: 6.0.1
CVE-2026-76400 (Medium – CVSS 5.9)
Issue: Event delivery disruption. Fixed Version: 2.2.7
CVE-2026-76401 (Medium – CVSS 5.9)
Issue: Connector worker DoS. Fixed Version: 2.2.7
CVE-2026-76402 (High – CVSS 8.2)
Issue: Credential exposure via HEC. Fixed Version: 2.2.7
CVE-2026-76403 (High – CVSS 7.4)
Issue: Data interception/modification. Fixed Version: 2.2.7
CVE-2026-76404 (Critical – CVSS 9.1)
Issue: OS command execution (RCE). Fixed Version: 1.2.1
CVE-2026-76405 (Medium – CVSS 4.3)
Issue: Partial API key exposure in Splunk On-Call. Fixed Version: 1.0.43
Splunk Connect for Kafka version 2.2.7 addresses four vulnerabilities. These include an SSRF flaw, CVE-2026-76402, rated 8.2, which could allow an unauthenticated attacker with access to the Kafka Connect REST API to configure a non-secure HTTP Event Collector endpoint and capture connector authentication credentials.
The Kafka fixes also address two denial-of-service issues and improper certificate validation in Kerberos-authenticated HTTP Event Collector deployments.
Splunk advises restricting Kafka Connect REST API access to trusted administrative systems and networks, enforcing secure transport for HTTP Event Collector endpoints, and applying finite retry limits.
Additional fixes include an SSRF vulnerability in Cisco Talos Intelligence for Enterprise Security Cloud and partial API key disclosure in Splunk On-Call (VictorOps). Updated versions are 1.0.3 for the Talos app and 1.0.43 for Splunk On-Call.
Organizations should inventory installed Splunk apps, apply relevant updates, restrict access to the management API, and review administrative and power-role assignments for signs of unnecessary privilege exposure.
Related Resource: Ensure your cloud telemetry and log management pipelines aren’t exposing your network to RCE threats. Learn more in our Agentic SOC vs Traditional SOC 2026 Guide .
What is the Splunk MCP Server RCE?
Tracked as CVE-2026-76404, it is a critical remote code execution vulnerability (CVSS 9.1) caused by unsafe deserialization in the Splunk MCP Server app. It allows authenticated admins to run arbitrary OS commands.
Is the Splunk AI Toolkit vulnerable?
Yes, Splunk patched nine vulnerabilities in the AI Toolkit, including CVE-2026-76395, which allows arbitrary code execution via maliciously crafted AI models containing unsafe pickle data.
Reported by CyberUpdates365 Desk
Delivering the latest insights on enterprise security, federal AI directives, and the future of IT infrastructure. Follow us for daily updates on how technology is reshaping the corporate landscape.
Found this article interesting? Follow us on Twitter and LinkedIn to read more exclusive content we post. You can also view our raw remediation script archives on our CyberUpdates365 GitHub repository .
Uday Patil Uday Patil is a Cybersecurity Researcher, DevSecOps Engineer, and the Founder of CyberUpdates365. Specializing in Threat Intelligence and Zero-Day vulnerability analysis, Uday is dedicated to breaking down complex cyber threats into actionable insights. His mission is to empower developers and security teams worldwide with rapid alerts, remediation scripts, and practical guidance to stay ahead of the evolving threat landscape.
Cyber Updates 365 Delivering the latest intelligence on global cyber threats, enterprise security, and emerging AI technologies. We empower IT professionals and tech enthusiasts with verified news and actionable insights to stay ahead in the digital battlefield.
Type at least 2 characters to search articles live...
