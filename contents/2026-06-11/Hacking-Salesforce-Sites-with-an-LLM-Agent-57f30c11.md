---
source: "https://www.reco.ai/blog/hacking-salesforce-sites-with-an-llm-agent"
hn_url: "https://news.ycombinator.com/item?id=48490289"
title: "Hacking Salesforce Sites with an LLM Agent"
article_title: "Hacking Salesforce Sites With an LLM Agent"
author: "llmacpu"
captured_at: "2026-06-11T16:29:55Z"
capture_tool: "hn-digest"
hn_id: 48490289
score: 2
comments: 0
posted_at: "2026-06-11T13:47:48Z"
tags:
  - hacker-news
  - translated
---

# Hacking Salesforce Sites with an LLM Agent

- HN: [48490289](https://news.ycombinator.com/item?id=48490289)
- Source: [www.reco.ai](https://www.reco.ai/blog/hacking-salesforce-sites-with-an-llm-agent)
- Score: 2
- Comments: 0
- Posted: 2026-06-11T13:47:48Z

## Translation

タイトル: LLM エージェントを使用した Salesforce サイトのハッキング
記事タイトル: LLM エージェントを使用した Salesforce サイトのハッキング
説明: AI エージェントが自律的に Salesforce の脆弱性を見つけて悪用できるようになりました。 Reco が実際の PII 暴露をどのように発見したか、そしてそれに対してチームが何をすべきかをご覧ください。

記事本文:
LLM エージェントを使用した Salesforce サイトのハッキング
シリーズ B 資金調達後、総額 8,500 万ドルを調達 ニュースを読む Reco が完全なデータ対応型 SaaS AI セキュリティのために Cyera と統合 ニュースを読む プラットフォーム プラットフォームのユースケース AI エージェントのセキュリティ 接続された AI エージェントの検出 姿勢管理とコンプライアンス SaaS セキュリティ態勢の改善 データ漏洩管理 SaaS 攻撃対象領域の削減 ID とアクセス ガバナンス 適切なアクセスの確保 アプリケーションの検出 すべてのアプリの検出と管理 脅威の検出と対応 脅威のアラートに優先順位を付ける Reco Factory 200 以上のアプリ統合、10 倍高速なポスチャ管理 データ露出管理 ID とアクセス ガバナンス 生成 AI 検出 アプリ検出とガバナンス シャドウ アプリ検出 脅威検出と対応 Reco AI エージェント 組み込み AI 検出 AI ガバナンスとセキュリティ Reco Factory ソリューション エージェント セキュリティの探索 エージェントティック ポスチャ管理 AI ガバナンスとコンプライアンス シャドウ AI 検出 SaaS セキュリティ SSPM エージェント セキュリティ AI ガバナンス エージェント セキュリティ 接続された AI の検出エージェント AI ガバナンス エージェントティック ポスチャ管理 姿勢を継続的に管理する AI エージェント AI ガバナンス AI ガバナンス AI ガバナンスとコンプライアンス AI エージェントとエージェント AI を検出して制御 AI ガバナンス シャドウ AI ディスカバリー あらゆる場所で不正な AI ツールを捕捉 AI ガバナンス SaaS セキュリティ すべてのアプリ、すべてのアップデートを検出して保護 AI ガバナンス SSPM SaaS の姿勢を自動的に気密に保つ 235 以上のアプリ。最大のエージェントとアプリのカタログ。すべての統合を探索する Open AI Cursor Glean Microsoft Copilot Workday Palantir Claude Salesforce Google リソース ブログ 専門家からの意見 IT ハブ IT セキュリティの頼りになるハブ CISO ハブ セキュリティ リーダー向けの戦略的リソース お客様事例 Reco がお客様をどのように支援したか ラーニング センター セルフサービス セキュリティ教育者

ウェビナー オンデマンド ビデオ コンテンツ リソース センター 最新の洞察とニュース 注目の記事 ガイド CISO ハンドブック: ルールが変わった神話 今すぐ読むガイド CISO AI セキュリティ ガイド 今すぐ読む ブログ ラーニング センター CI
[切り捨てられた]
AI はセキュリティの状況を変えています。偵察や悪用のワークフローに LLM を組み込む脅威グループが増えています。一部の脆弱性は実装するには複雑すぎるという概念は、現在では時代遅れです。 LLM を使用すると、ハッカーは複雑な脆弱性を自動的に見つけて悪用できます。 Claude Mythos と、大規模なコードベースの脆弱性を特定し、自動的に悪用するその機能については、誰もが聞いたことがあるでしょう。しかし、LLM はコード内の脆弱性を見つけるだけではありません。
ShinyHunters は、何千もの Salesforce サイトをスキャンしてきました。彼らは「AuraInspector」の修正版を使用しました。おそらく LLM を使用して、フレームワーク、MOD、偵察ツール、およびワークフローのその他の側面をコーディングした可能性があります。しかし、次のステップは、AI を使用して攻撃プロセス自体を強化することです。私たち Reco は、それがどのようなものになるかを調査することにしました。
Reco のセキュリティ研究チームは、Salesforce Experience Cloud サイトのエンドツーエンドのセキュリティ評価を完全に自律的に実行できる AI を活用したエージェントを構築しました。 URL を指定すると、攻撃対象領域を検出し、公開されているすべてのエンドポイントを分析し、脆弱性を特定し、機能するエクスプロイトを作成して実行します。ターゲットを提供した後は人間による誘導は必要ありません。
私たちはこれを大手テクノロジー企業に属する実際の Salesforce サイトに向けましたが、その結果は厳しいものでした。エージェントは、セキュリティに多額の投資を行っている組織に属するサイトで重大度の高い脆弱性を発見しました。機能するエクスプロイト スクリプトを最初から作成し、実際のデータを抽出し、パブリック ソースからデータを自律的に取得することもできました。

必要に応じてペイロード入力を構築します。
エンドツーエンドのセキュリティアナリスト
私たちのエージェントは単一のモノリシック スクリプトではありません。これは AI スキルのエージェント パイプラインであり、それぞれが評価の個別のフェーズを担当します。人間の研究者も同様のワークフローに従い、偵察、分析、活用、検証を行います。違いは、すべてのフェーズが LLM によって実行され、人間の介入なしに、目にしたものを推論し、アプローチを適応させ、判断を下すことができることです。 LLM はすべてを制御し、各スキルの使用方法を選択できます。フェーズ 1 ～ 5 の一般的なワークフローがありますが、LLM は必要に応じて、戻って別のコンテキストでスキルを再実行することを選択する場合があります。
フェーズ 1: 偵察 - 攻撃対象領域のマッピング
エージェントは URL だけから始まります。その最初のタスクは、サイトが認証されていない訪問者に公開しているすべてのものを検出することです。
Salesforce Aura フレームワークにクエリを実行して、アクセス可能なすべての Salesforce オブジェクト (データベーステーブル)、Apex コントローラメソッド (サーバー側ビジネスロジック)、ルート (公開ページ)、およびコンテンツ (ファイルとドキュメント) を列挙します。出力は包括的なサイト コンテキスト、つまりメソッド シグネチャ、パラメータの名前とタイプ、オブジェクト スキーマを含む攻撃対象領域全体のマップです。これは主に決定論的な手順に基づいており、これについてはペネトレーション テスト ガイドで読むことができます。 LLM は、オブジェクトと apex メソッド (パラメータを含む) を決定的に列挙する Python ツールを呼び出すことができます。 LLM は、必要と判断した場合、会社に関する情報など、追加のコンテキストも提供することを選択できます。
フェーズ 2: オブジェクト分析 - データ漏洩の調査
攻撃対象領域をマッピングしたら、エージェントはデータ分析に移ります。検出されたすべての Salesforce オブジェクトを機密性 (どのテーブルに sens が含まれる可能性があるか) ごとに分類します。

アクティブなデータ。 「Contact」、「Lead」、または「CreditCardTransaction__c」などのテーブルが「BlogPostEntry__c」よりも優先されます。
関心の高いオブジェクトごとに、エージェントはスキーマ (フィールド名、タイプ、カスタム フィールド) を検査し、ゲスト ユーザーとしてレコードのクエリを試みます。レコードにアクセスできるかどうかだけでなく、それらのレコード内のデータが機密であるかどうか、返される内容を評価します。
特にファイル オブジェクトの場合、エージェントは専用のワークフローに従います。つまり、アクセス可能なドキュメントを列挙し、メタデータを検査し、各ファイルのコンテンツをダウンロードして読み取り、ファイルに機密情報が含まれているかどうかを評価します。このようにして、何百もの平凡なファイルの中に埋もれている 1 つの機密ファイルを見つけます。
ここで LLM がゲームチェンジャーとなります。公開されたすべての Apex メソッドについて、エージェントはテストを試みます。シグネチャを分析し、メソッドを実行するためのさまざまなベースラインを構築します。
1. 有効な値を推論します。メソッド シグネチャ、パラメータ名、クライアント側 JavaScript 分析、および以前にダンプされたオブジェクトからのデータを使用して、エージェントは各メソッドがどのような入力を期待しているかを推論します。メソッドが caseId を受け取る場合、エージェントはすでに検出されている Case レコードを検索します。 emailAddress を受け取る場合は、一般的なパターンが試行されます。
2. メソッドを呼び出し、応答を分析します。エージェントは、テストしても安全なすべてのメソッド (削除/変更メソッドは除く) を呼び出し、応答を読み取り、ゲスト ユーザーがアクセスすべきではないデータ (PII、内部レコード、資格情報、構成の詳細など) を返すかどうかを評価します。
3. SOQL インジェクションのプローブ: 注入可能なパラメータ (文字列または複合型) を持つすべてのメソッドに対して、エージェントは一重引用符、トートロジー、ワイルドカード ペイロードの順に挿入を試みます。応答をベースラインと比較して、文字列の連結を示す動作の変化を検出します。

SOQL クエリの n。エラーメッセージ、変更された結果セット、シフトされたカウントなどの違いが見つかると、インジェクションが確認され、オラクルの特徴がわかります。
出力は、テストされたすべてのメソッドを網羅する詳細な分析レポートであり、確認された結果がタイプと重大度ごとに分類されます。
フェーズ 4: エクスプロイト - エクスプロイトの作成と実行
フェーズ 3 でエクスプロイト可能なインジェクションが確認されると、エージェントはそれを報告するだけでなく、動作する概念実証エクスプロイトを作成します。
エージェントは、完全な悪用チェーンを実装するスタンドアロン Python スクリプトを生成します。ブラインド SOQL インジェクションの場合、これは次のことを意味します。
脆弱なオブジェクトからユーザー、連絡先、リードなどの価値の高いターゲットにピボットするサブクエリの構築
LIKE 接頭辞一致を使用した文字ごとの抽出の実装
頻度順の文字セットを使用して最適化して HTTP リクエストを最小限に抑える
次に、エージェントはエクスプロイトを実行し、データが実際に抽出されたことを検証し、取得した内容を正確に文書化します。 User テーブルから従業員の電子メールを正常に抽出すると、そこで終了するわけではありません。追加のフィールド (電話番号、役職)、追加のピボット関係 (OwnerId、CreatedById)、および追加のターゲット オブジェクト (連絡先、リード) が調査されます。
最終段階はセキュリティ評価者にとって重要です。エージェントは、懐疑的なプログラム マネージャーの観点から、自身の調査結果をレビューします。脅威グループはこれをスキップする可能性がありますが、重大度の高い脆弱性を持つターゲットに焦点を当てたい場合には有益である可能性があります。
脆弱性ごとに次の質問が行われます。
実際に抽出されたデータは何ですか?
重大度は実証された影響を反映していますか、それとも単に脆弱性クラスを反映していますか?
この敵対的な自己レビューは重大度のインフレを捉え、実際の PII の露出とメタデータの漏洩を区別し、最終レポートが防御可能であることを保証します。もちろん

これで終わりではありません。エージェントはフェーズ 4 に戻って、より強力なエクスプロイトを見つけようとする可能性があります。
私たち自身の研究所でテストした後、それを許可する脆弱性開示プログラムを持つ数社をターゲットにしました。エージェントを明示的に制約しました。書き込み、削除、または変更操作は禁止です。一括抽出はありません。検査は副作用を引き起こす可能性のない方法に限定されました。目標は、データ漏洩を最大化することではなく、悪用可能性と影響を実証することでした。
このセキュリティ評価エージェントは、悪意のある攻撃者ではなく侵入テスターとして機能するように構築されていますが、結果は同じです。 PII を含むデータ ダンプはまだあります。主な違いはダンプのサイズです。
開示メモ: この投稿で説明されているすべての脆弱性は実際のものであり、影響を受ける組織のセキュリティ プログラムを通じて責任を持って開示されました。調査結果の技術的な正確性を維持しながら組織を保護するために、会社名は架空の名前に置き換えられています。
ケーススタディ 1: Aegis Security - 「このメールについてすべて教えてください」
会社概要: Aegis Security (社名変更) は、Salesforce Experience Cloud 上に構築されたパートナー ポータルを持つ大手サイバーセキュリティ ベンダーです。このポータルを使用すると、テクノロジー パートナーやチャネル パートナーは企業との関係を管理できます。このサイトは一般にアクセスできるように意図されていません。「ゲスト ユーザーはログインせずにサイトを表示および操作できる」が無効になっています。
エージェントは、攻撃対象領域全体 (263 個の Salesforce オブジェクト、9 つのコントローラ クラスにわたる 55 個の Apex メソッド、および 10 個のパブリック ルート) をマッピングすることから始めました。各エンドポイントを系統的に分析し、テスト入力でパラメータを調べてアプリケーションの動作を理解しました。
エージェントが PartnerPortalOnboardingController.getContactInfo に到達したとき、憂慮すべき事態を発見しました。この Apex メソッドでは、

パートナーのオンボーディング フローに対応し、電子メール アドレスという 1 つのパラメータを受け入れます。そして、認証されていないゲスト ユーザーもアクセスできました。
エージェントは一般的な電子メールでテストし、次のような返信を受け取りました。
{
"名" : "..." ,
"姓" : "..." ,
「電話」 : "1234567789" 、
"電子メール" : "test@example.com" ,
"タイトル" : "..." ,
「アカウント」: {
「タイプ」 : 「パートナー」 、
「名前」：「ブライトウェーブソリューションズ株式会社」 、
"BillingStreet" : "西新宿3-7-1 新宿パークタワー" 、
"請求先国" : "日本" 、
"BillingState" : "東京" 、
"請求先郵便番号" : "163-1055"
}
}
‍
任意の電子メール アドレスを指定した 1 つの未認証 HTTP リクエストにより、完全な連絡先レコード (名前、電話番号、役職) に加えて、リンクされたアカウントの完全な請求先住所が返されました。不明な電子メールの場合、サーバーは範囲外のリスト インデックス: 0 を返し、登録されたアドレスと未登録のアドレスを区別するためのクリーンなオラクルを提供しました。
ここで AI が真に輝きました。脆弱性を発見した後、それを兵器化することを決定した。エージェントは、この脆弱性の実際的な影響は、攻撃者が問い合わせる電子メール アドレスのリストを持っているかどうかに依存していることを認識しました。 LinkedIn やその他のパブリック ソース (Aegis のパブリック コミュニティ ポータルなど) に自律的に移動しました。

[切り捨てられた]

## Original Extract

AI agents can now autonomously find and exploit Salesforce vulnerabilities. See how Reco uncovered real PII exposures and what your team should do about it.

Hacking Salesforce Sites With an LLM Agent
Raises $85M Total after Series B Funding Read the News Reco Now Integrates with Cyera for Complete Data-Aware SaaS AI Security Read the News Platform Platform Use Cases AI Agent Security Discover connected AI agents Posture Management & Compliance Improve SaaS security posture Data Exposure Management Reduce SaaS attack surface Identity & Access Governance Ensure appropriate access Application Discovery Discover & manage all apps Threat Detection & Response Prioritize alerts of threats Reco Factory 200+ app integrations, 10x faster Posture Management Data Exposure Management Identity & Access Governance Generative AI Discovery App Discovery & Governance Shadow App Discovery Threat Detection & Response Reco AI Agents Embedded AI Discovery AI Governance & Security Reco Factory Solutions Explore Agent Security Agentic Posture Management AI Governance & Compliance Shadow AI Discovery SaaS Security SSPM Agent Security AI Governance Agent Security Discover connected AI agents AI Governance Agentic Posture Management AI agents that manage posture continuously AI Governance AI Governance AI Governance & Compliance Discover and control AI agents and agentic AI AI Governance Shadow AI Discovery Catch unauthorized AI tools everywhere AI Governance SaaS Security Discover and secure every app, every update AI Governance SSPM Keep SaaS posture airtight automatically 235+ apps. The largest agent and app catalog. Explore all integrations Open AI Cursor Glean Microsoft Copilot Workday Palantir Claude Salesforce Google Resources Blog Thoughts from our experts IT Hub The go-to hub for IT security CISO Hub Strategic resources for security leaders Customer Stories How Reco helped customers Learning Center Self-service security education Webinars On demand video content Resource Center Latest insights and news Featured Articles Guide CISO Playbook: Mythos Changed the Rules Read Now Guide CISO Guide to AI Security Read Now Blog Learning Center CI
[truncated]
AI is changing the security landscape. More and more threat groups incorporate LLMs into their reconnaissance and exploitation workflows. The notion that some vulnerabilities are too complex to implement is now obsolete. Using LLMs, hackers can automatically find and exploit complex vulnerabilities. We have all heard of Claude Mythos and its ability to identify vulnerabilities in large codebases and exploit them automatically. But LLMs can do more than find vulnerabilities in code.
ShinyHunters has scanned thousands of Salesforce Sites. They used a modified version of "AuraInspector". They possibly used an LLM to code their framework, mods, reconnaissance tools, and other aspects of their workflow. But the next step is to use AI to supercharge the attack process itself. We at Reco decided to explore what it would look like.
Reco's security research team built an AI-powered agent capable of performing end-to-end security assessments of Salesforce Experience Cloud sites - fully autonomously. Give it a URL, and it discovers the attack surface, analyzes every exposed endpoint, identifies vulnerabilities, writes working exploits, and runs them. No human guidance required after providing the target.
We pointed it at real-world Salesforce sites belonging to major technology companies, and the results were sobering. The agent discovered high-severity vulnerabilities on sites belonging to organizations that invest heavily in security. It wrote working exploit scripts from scratch, extracted real data, and even autonomously retrieved data from public sources to build payload input when needed.
The End-to-End Security Analyst
Our agent is not a single monolithic script. It's an agentic pipeline of AI skills, each responsible for a distinct phase of the assessment. A human researcher would follow a similar workflow: reconnaissance, analysis, exploitation, validation. The difference is that every phase is executed by an LLM that can reason about what it sees, adapt its approach, and make judgment calls without human intervention. The LLM controls everything, and it can choose how to use each skill. There is a generic workflow, phases 1-5, but the LLM may choose to go back and rerun skills with a different context as it sees fit.
Phase 1: Reconnaissance - Mapping the Attack Surface
The agent starts with nothing but a URL. Its first task is to discover everything the site exposes to unauthenticated visitors.
It queries the Salesforce Aura framework to enumerate all accessible Salesforce objects (database tables), Apex controller methods (server-side business logic), routes (public pages), and content (files and documents). The output is a comprehensive site context - a map of the entire attack surface, including method signatures, parameter names and types, and object schemas. This is mostly based on deterministic steps, which you can read about in our pentesting guide. The LLM can invoke a Python tool that deterministically enumerates the objects and the apex methods, including their parameters. The LLM may choose to provide additional context as well if it deems it necessary, such as information about the company.
Phase 2: Object Analysis - Probing for Data Exposure
With the attack surface mapped, the agent shifts to data analysis. It categorizes every discovered Salesforce object by sensitivity: what tables may contain sensitive data. It prioritizes tables such as "Contact", "Lead", or "CreditCardTransaction__c" over "BlogPostEntry__c".
For each high-interest object, the agent inspects the schema (field names, types, custom fields), then attempts to query records as a guest user. It evaluates what comes back - not just whether records are accessible, but whether the data in those records is sensitive.
For file objects specifically, the agent follows a dedicated workflow: enumerate accessible documents, inspect metadata, download and read the content of each file, and assess whether it contains sensitive or confidential information. This is how it finds the one sensitive file buried among hundreds of mundane ones.
Here the LLM is a game-changer. For every exposed Apex method, the agent tries to test it. It analyzes the signature and builds different baselines to run the method.
1. Infers valid values: Using method signatures, parameter names, client-side JavaScript analysis, and data from previously dumped objects, the agent reasons about what inputs each method expects. If a method takes a caseId, the agent looks for Case records it already discovered. If it takes an emailAddress, it tries common patterns.
2. Invokes the method and analyzes the response: The agent calls every safe-to-test method (no delete/modify ones), reads the response, and evaluates whether it returns data that guest users shouldn't have access to, including PII, internal records, credentials, and configuration details.
3. Probes for SOQL injection: For every method with potentially injectable parameters (strings or complex types), the agent tries to inject a single quote, then tautologies, then wildcard payloads. It compares responses against the baseline to detect behavioral changes that indicate string concatenation in SOQL queries. When it finds a difference - an error message, a changed result set, a shifted count - it confirms the injection and characterizes the oracle.
The output is a detailed analysis report covering every tested method, with confirmed findings classified by type and severity.
Phase 4: Exploitation - Writing and Running Exploits
When Phase 3 confirms an exploitable injection, the agent doesn't just report it - it writes a working proof-of-concept exploit.
The agent generates standalone Python scripts that implement the full exploitation chain. For blind SOQL injection, this means:
Constructing subqueries that pivot from the vulnerable object to high-value targets like User, Contact, or Lead
Implementing character-by-character extraction using LIKE prefix matching
Optimizing with frequency-ordered character sets to minimize HTTP requests
The agent then runs the exploit, validates that data is actually extracted, and documents exactly what was obtained. If it successfully extracts an employee email from the User table, it doesn't stop there - it probes additional fields (phone numbers, titles), additional pivot relationships (OwnerId, CreatedById), and additional target objects (Contact, Lead).
The final phase is important for security assessors. The agent reviews its own findings from the perspective of a skeptical program manager. Threat groups might skip it, though it could be beneficial if they want to focus on targets with high-severity vulnerabilities.
For each vulnerability, it asks:
What data was actually extracted?
Does the severity reflect demonstrated impact or just the vulnerability class?
This adversarial self-review catches severity inflation, distinguishes real PII exposure from metadata leaks, and ensures the final report is defensible. Of course, this is not the end - the agent may go back to phase 4 to try to find a stronger exploit.
After testing it in our own labs, we targeted several companies with vulnerability disclosure programs that allowed it. We constrained the agent explicitly: no write, delete, or modify operations; no bulk extraction; testing limited to methods that could not cause side effects. The goal was demonstrating exploitability and impact, not maximizing data exfiltration.
This security assessor agent was built to work as a pentester, and not a malicious attacker - but the results are the same; we still got data dumps, including PII. The main difference is the size of the dump.
Disclosure note: All vulnerabilities described in this post are real and were responsibly disclosed to the affected organizations through their security programs. Company names have been replaced with fictional names to protect the organizations while preserving the technical accuracy of the findings.
Case Study 1: Aegis Security - "Tell Me Everything About This Email"
Company profile: Aegis Security (name changed) is a major cybersecurity vendor with a partner portal built on Salesforce Experience Cloud. The portal allows technology and channel partners to manage their relationship with the company. The site was not meant to be publicly accessible - "Guest users can see and interact with the site without logging in" is disabled.
The agent began by mapping the full attack surface: 263 Salesforce objects, 55 Apex methods across 9 controller classes, and 10 public routes. It systematically analyzed each endpoint, probing parameters with test inputs to understand the application's behavior.
When the agent reached PartnerPortalOnboardingController.getContactInfo, it discovered something alarming. This Apex method, intended for the partner onboarding flow, accepted a single parameter: an email address. And it was accessible to unauthenticated guest users.
The agent tested it with a generic email and received back:
{
"FirstName" : "..." ,
"LastName" : "..." ,
"Phone" : "1234567789" ,
"Email" : "test@example.com" ,
"Title" : "..." ,
"Account" : {
"Type" : "Partner" ,
"Name" : "Brightwave Solutions K.K." ,
"BillingStreet" : "Shinjuku Park Tower, 3-7-1, Nishishinjuku" ,
"BillingCountry" : "Japan" ,
"BillingState" : "Tokyo" ,
"BillingPostalCode" : "163-1055"
}
}
‍
A single unauthenticated HTTP request, given any email address, returned the full Contact record - name, phone number, job title - plus the linked Account's complete billing address. For unknown emails, the server returned List index out of bounds: 0, providing a clean oracle to distinguish registered from unregistered addresses.
Here's where the AI truly shone. After discovering the vulnerability, it decided to weaponize it. The agent recognized that the vulnerability's practical impact depends on an attacker having a list of email addresses to query. It autonomously navigated to LinkedIn and other public sources (such as Aegis's public community porta

[truncated]
