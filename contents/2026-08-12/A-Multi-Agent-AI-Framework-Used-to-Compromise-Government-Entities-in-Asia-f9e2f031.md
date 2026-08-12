---
source: "https://dreamgroup.com/blog/inside-a-multi-agent-ai-framework-used-to-compromise-government-entities-in-asia"
hn_url: "https://news.ycombinator.com/item?id=49276153"
title: "A Multi-Agent AI Framework Used to Compromise Government Entities in Asia"
article_title: "Inside a Multi-Agent AI Framework Used to Compromise Government Entities in Asia | | Dream Security Blog"
author: "mikeleeorg"
captured_at: "2026-08-12T17:51:17Z"
capture_tool: "hn-digest"
hn_id: 49276153
score: 1
comments: 0
posted_at: "2026-08-12T17:48:47Z"
tags:
  - hacker-news
  - translated
---

# A Multi-Agent AI Framework Used to Compromise Government Entities in Asia

- HN: [49276153](https://news.ycombinator.com/item?id=49276153)
- Source: [dreamgroup.com](https://dreamgroup.com/blog/inside-a-multi-agent-ai-framework-used-to-compromise-government-entities-in-asia)
- Score: 1
- Comments: 0
- Posted: 2026-08-12T17:48:47Z

## Translation

タイトル: アジアの政府機関を侵害するために使用されるマルチエージェント AI フレームワーク
記事のタイトル: アジアの政府機関を侵害するために使用されるマルチエージェント AI フレームワークの内部 | |ドリームセキュリティブログ

記事本文:
アジアの政府機関への侵害に使用されるマルチエージェント AI フレームワークの内部 | |ドリームセキュリティブログ
新たな現実 - CEO の投稿を読む >
キャリアについて 研究とニュース マニフェスト お問い合わせ お問い合わせ フォームに記入して専門家チームにご連絡ください。
アジアの政府機関への侵害に使用されるマルチエージェント AI フレームワークの内部
AI を活用した攻撃作戦は現在、転換点にあります。これは、次の 3 つの曲線の収束によって決まります。
モデルの機能は向上し続けており、ハードウェアを使用するすべてのオペレーターの手にフロンティアに隣接した推論を提供するオープンウェイトを実現します。
エージェント ハーネスは、デモンストレーションから運用の足場へと成熟しました。プランニング ループ、並列ディスパッチ、永続メモリ、および構造化されたアフターアクション レポートにより、侵入キャンペーンに関する質問に答えるのではなく、モデルが実行できるようになります。
最後の実際的な制約であるガードレールは、正直に質問するオペレーターに対してのみ適用されます。
以下に述べるのは、ほぼ自律的な攻撃と思われる具体的な例の 1 つであり、容易に入手可能なハーネスやモデルを利用して国家を狙ったものです。攻撃の詳細は当初、フィナンシャル・タイムズ紙と共有されました。
約 4 日間で、エージェント攻撃者は 1,395 個のファイル、85 個のクラックされた認証情報、数千件の流出した人事記録を作成し、州のインフラ内に永続的な足場を築きました。それは、有効な攻撃を実行するコストは崩壊したが、攻撃を防御するコストは崩壊していないということを声高に訴えている。
政府の AI 攻撃者の構造
2026 年 7 月初旬、DREAM Lab の脅威研究チームは、アジアの政府機関に対して積極的に侵入キャンペーンを実施していた自律型 AI 攻撃フレームワークの完全な運用ワークスペースを発見しました。弧

ハイブは 160 メガバイト、1,395 ファイルを超えており、州のインフラストラクチャに対して確認済みの現実世界の侵害を達成したマルチエージェント AI システムを明らかにしています。
このフレームワークは、Hermes エージェントと OpenClaw エージェントに基づいて構築されており、ウェーブごとに最大 8 つの文字のサブエージェント (キャンペーン全体で観測されたエージェント A からエージェント Q) を並行して展開し、それぞれが個別のターゲットと攻撃手法に割り当てられます。約 4 日間 (2026 年 7 月 1 日から 4 日) にわたって実施された 12 回の文書化された攻撃波にわたって、これらのエージェントは自律的に政府職員の資格情報を解読し、未認証の API エンドポイントから数百の人事記録を抽出し、政府の個人認証サービスの署名検証の欠陥を発見し、政府の Web アプリケーションに永続的なバックドアをインストールしました。
このフレームワークがこれまでに見た攻撃ツールと異なるのは、その運用インテリジェンスです。
ベイジアン優先順位付け: 事後確率スコアリングにより、14 の並列攻撃チェーンを継続的にランク付けして優先順位を再設定し、最初に最も価値の高いターゲットに重点を置きます。
自律的な調査: 既存の手法がブロックされている場合に、脆弱性データベース、GitHub リポジトリ、およびセキュリティ出版物を検索して新しい悪用手法を探す「学習サイクル」
フィードバック ループ : 各ウェーブの結果を次のウェーブの計画にフィードバックする構造化されたアクション後のレポートにより、人間の介入なしにフレームワークが運用中に適応できるようになります。
フレームワーク自身の安全ガードレール (LLM モデルの拒否) は、すべてのアクティビティを「認可された侵入テスト」として枠組み化することで回避されました。
運用文書の言語分析（内部状況報告では簡体字中国語と目標に向けた分析では繁体字中国語をコード切り替えする）では、中国語が指摘されています。

言語演算子。
このレポートでは、フレームワークのアーキテクチャ、その確認された影響、およびそれが防御側に示すパラダイムシフトについて詳しく説明します。政府インフラに対する AI で組織化され、並行して自律的に適応するサイバー作戦の時代は、従来の防御能力のほとんどを上回る攻撃能力を備え、ますます危険になるだけの現実です。
攻撃チェーン: 攻撃者が成功したこと
ステップ 1: 偵察: ターゲットを理解する
このフレームワークは、政府のエコシステム全体を自動的にマッピングすることから始まりました。これは、Angular ベースの政府ポータルから JavaScript バンドルをダウンロードして逆コンパイルし、コンパイルされたコードに隠されているすべての埋め込み URL、API エンドポイント、OAuth クライアント ID、および Keycloak 設定オブジェクトを抽出しました。この単一の開始点から、接続されている 21 の政府システムを特定し、6 つのサブレルム、すべての OIDC エンドポイント、2 つの RSA 署名キー、およびサポートされているすべての認証フローを含む完全な全国 SSO アーキテクチャをマッピングしました。
1 つのターゲットだけで、アカウント管理、ユーザー データの取得、ファイルのアップロード、管理機能にわたる 36 以上の API エンドポイントが発見されましたが、その多くは完全に認証されていませんでした。重要なのは、システムの 1 つが、名前、部門、SSO アカウント ID を含む数千の従業員レコードなど、ユーザー データベース全体を認証なしで公開したことです。
ステップ 2: SDK ハーベスティング: 偵察の加速
完全に自律的な決定のように見えますが、JS バンドルから抽出された URL によって、フレームワークは全国的な SSO 統合ガイドをホストする GitBook ドキュメント ポータルに誘導されました。 GitBook の組み込み機能 (ページのインデックス作成には llms.txt、生のコンテンツには .md 接尾辞) を使用して、体系的にドキュメントを収集し、ダウンロード リンクを見つけ、2 つの SDK 統合を取得しました。

サンプル プロジェクト (Java Spring Boot および ASP.NET Core 8.0) — 政府機関が国家 SSO に接続する方法を示すクライアント側のデモ アプリケーション。
次に、フレームワークは SDK サンプル コードに対して AI を活用した静的分析 (Local-LLM-SAST-0day 手法) を実行し、オープン リダイレクト (CSRF 無効 + 未検証の redirectUri パラメーター) や未認証の承認コードからトークンへの交換パターンなどの潜在的なセキュリティ問題を特定しました。ただし、これらの調査結果はどれも、実際のシステムに対する確認済みのエクスプロイトにつながるものではありませんでした。 SAST の結果をフレームワーク独自の最終検証済み脆弱性リストと相互参照すると、重複は見られません。実際の侵害は、SDK の知識がなくても標準的なブラックボックス テストを通じて発見できる、サーバー側の欠陥 (バックドア エンドポイント、署名されていない JWT の受け入れ、認証されていない API) から発生しました。
ステップ 3: 初期アクセス: 複数の方法でアクセス
攻撃対象領域がマップされているため、フレームワークは並行して実行されているいくつかのエントリ ポイントを検出しました。
認証バックドア : 政府の Web アプリケーションで、あらゆるリクエスト本文を受け入れ、有効な認証済みセッションを返す 3 つの隠れた API エンドポイントを発見しました。認証情報は必要ありません。これらは、運用環境で公開されたままになっている開発者のデバッグ エンドポイントでした。
自動資格情報スプレー : フレームワークは、認証されていないユーザー API エンドポイントから収集した従業員のユーザー名を使用して、部門のオフィス オートメーション ポータルを攻撃しました。ポータルは CAPTCHA によって保護されていましたが、フレームワークは Tesseract OCR を使用して、それぞれの小さな CAPTCHA 画像を 100% の精度で解決しました。各従業員の ID に基づいて予測可能なパスワード パターン (ユーザー名を大文字にするか、一般的な記号で装飾するなど) をテストし、複数回のスプレー ラウンドで 85 のアカウントを解読しました (最初は 12 個、その後追加の攻撃として 73 個のアカウントをクラックしました)

エルンが発見されました）。
JWT 署名バイパス : 政府 API は、アルゴリズム フィールドが none に設定された JWT トークンを受け入れました。これは、署名キーを知らなくても誰でも有効な ID トークンを偽造できることを意味します。
ステップ 4: 横方向の動きと持続性
ステップ 3 で解読された従業員の資格情報は、シングル サインオン経由で接続されている他のすべての内部システムに対して体系的にテストされました。各内部システムは、オフィス オートメーション セッション (ステップ 3 以降) を自動的に信頼する専用の SSO エントリ ポイントを公開しました。追加の認証、MFA、またはユーザー確認は必要ありません。
クラックされた 85 個のアカウントのうち 84 個 (98.8%) は、SSO ブリッジ エンドポイント経由で部門の内部情報システムへの認証に成功し、内部ダッシュボード、機器管理インターフェイス、人事統計ページへのアクセスを取得しました。フレームワークはこれを大規模に自動化し、複数の接続されたシステムに対して各認証情報を並行してテストしました。
また、このフレームワークは、無制限のファイル アップロード API を介して Web シェルをアップロードし、サーバー上の既知のパスに配置します。ただし、フォーム認証の 2 番目の層によって実行がブロックされました。部分的には成功しましたが、リモート コード実行の手前で停止しました。
バックドア アクセス、クラックされた認証情報、未認証の API エンドポイントを組み合わせて、フレームワークは以下を抽出しました。
2,564 件以上の人事レコード: 従業員 1,409 人 (名前、部門、SSO ID)、未認証 API のユーザー 916 人、未認証の法務省エンドポイントの法律専門家 239 人
完全なユーザー データベース: すべての部門システム ユーザーの完全な JSON エクスポート
内部アーキテクチャ: 7 つの SSO クライアント シークレット (ローテーションされているが命名規則が明らかに)、6 つの内部データベース資格情報 (MSSQL、Oracle、Sybase)、および内部ネットワーク IP 範囲
ステップ 6 : サプライチェーンの拡大: ベンダーへの転換
の

攻撃者は主なターゲットにとどまりませんでした。政府の IT サプライ チェーン ベンダー、原子力安全機関、政府の電子メール システム、および 7 社以上のエネルギー セクターの企業に業務を拡大し、構成ミス、露出した管理インターフェイス、悪用可能な脆弱性をすべて並行してスキャンしました。
アーキテクチャ: Hermes と OpenClaw
このフレームワークは 2 つのワークスペース識別子 (.hermes と .openclaw) の下で動作し、レポートは構造化されたパスに生成されます。
各サブエージェントには、文字の指定と、認証情報の盗難から API の悪用、サプライ チェーンの偵察に至るまで、特定の任務が割り当てられます。 1 回の攻撃ウェーブで、最大 8 つのエージェントが同時に展開され、それぞれが異なる攻撃対象領域をターゲットにしていることが観察されました。
文書化されたキャンペーンの 12 回の攻撃ウェーブ全体で、A から Q までのラベルが付けられたエージェントを特定しました。1 つのウェーブで最大 8 人が同時に展開されました (3 回のバッチに分けて派遣されました)。約 4 日間で 1,395 ファイルが作成されるという膨大な量の出力は、人間のオペレーターが単独で作成できる量をはるかに超える高度な自動化と一致しています。
ベイジアン脳: 2 層の確率的意思決定エンジン
おそらく、このフレームワークの最も顕著な特徴は、ベイジアン事後確率スコアリングを単一の指標としてではなく、個々の脆弱性レベルと戦略的攻撃チェーン レベルの両方で動作する 2 層の意思決定エンジンとして使用していることです。言い換えれば、これは「スプレーして祈る」アプローチではなく、アプローチの優先順位を動的に変更する意図的な戦略的推進攻撃でした。
レイヤ 1: 個別の脆弱性スコアリング
このフレームワークのトリアージ レポートは、最初の層、つまり正式なベイジアン モデルを使用して個々の所見をスコアリングする自律的なトリアージ システムを明らかにします。この方法論では、すべての脆弱性が次の時点で開始されます。

an uninformative prior of P=0.50, then updates based on evidence using explicit likelihood ratios:
The resulting posterior determines what happens next:
Once individual vulnerabilities are triaged, the framework connects them into multi-step attack paths and applies a second probability layer.アーカイブからの明示的な式:
P_success = P_chain × (1 - P_blocker)
P_chain = confirmed steps / total steps in the chain
P_blocker = probability of an insurmountable blockerScoring thresholds:
95%+ = all critical steps already confirmed and directly exploitable
60-94% = most steps confirmed, minor attempts needed
20-59% = partial confirmation, significant blockers present
5-19% = chain exists but critical conditions are missing
<5% = 克服できないブロッカーが特定された
Worked example — SSO Lateral Movement, rated 99%:
このフレームワークは、3 つの個別に確認された調査結果からこのチェーンを組み立てました。(1) 数千の従業員アカウントを生成する未認証のユーザー リスト、(2) 複数のアカウントをクラッキングするパスワード スプレー、および (3) ラテラル ピボット機能を確認する SSO トークンの取得。前提条件テーブル:
All critical steps were already tested and working. The only theoretical blocker — the target going offline — had near-zero probability since it was a

[切り捨てられた]

## Original Extract

Inside a Multi-Agent AI Framework Used to Compromise Government Entities in Asia | | Dream Security Blog
A New Reality - Read the CEO's Post >
About Careers Research & News Manifest Contact Us Contact Us Fill out the form to get in touch with our Expert Team.
Inside a Multi-Agent AI Framework Used to Compromise Government Entities in Asia
AI-enabled offensive operations are now at an inflection point. This is driven by the convergence of three curves:
Model capability keeps climbing, and it climbs on open weights that puts frontier-adjacent reasoning in the hands of any operator with hardware
Agentic harnesses have matured from demonstrations into operational scaffolding: planning loops, parallel dispatch, persistent memory, and structured after-action reporting that let a model run an intrusion campaign rather than answer questions about one
Guardrails, the last practical constraint, hold only against operators who ask honestly.
What follows is but one concrete example of what appears to be a near-autonomous attack, running off readily available harnesses and models and aimed at a nation state. Details on the attack were initially shared with the Financial Times .
In roughly four days, the agentic attacker produced 1,395 files, 85 cracked credentials, thousands of exfiltrated personnel records, and gained a persistent foothold inside state infrastructure. It spells out one thing loudly - the cost of running a competent attack has collapsed, but the cost of defending against one has not.
The Anatomy of a Government AI Attacker
In early July 2026, DREAM Lab's Threat Research team uncovered the complete operational workspace of an autonomous AI attack framework that had been actively conducting intrusion campaigns against government entities in Asia. The archive, spanning over 160 megabytes and 1,395 files, reveals a multi-agent AI system that achieved confirmed, real-world compromises against state infrastructure.
The framework — built on the Hermes and OpenClaw agents — deploys up to 8 lettered sub-agents in parallel per wave (Agent A through Agent Q observed across the campaign), each assigned to distinct targets and attack techniques. Across 12 documented attack waves conducted over approximately four days (July 1-4, 2026), these agents autonomously cracked government employee credentials, exfiltrated hundreds of personnel records from unauthenticated API endpoints, discovered a signature validation flaw in the government's personal authentication service, and installed persistent backdoors on government web applications.
What distinguishes this framework from attack tooling that we've seen before is its operational intelligence:
Bayesian prioritization : Posterior probability scoring to continuously rank and reprioritize 14 parallel attack chains, focusing effort on the highest-value targets first
Autonomous research : "Learning Cycles" that search vulnerability databases, GitHub repositories, and security publications for new exploitation techniques when existing methods are blocked
Feedback loops : Structured after-action reporting that feeds results from each wave back into the planning for the next, enabling the framework to adapt mid-operation without human intervention
The framework's own safety guardrails — LLM model refusals — were bypassed by framing all activity as "authorized penetration testing".
Linguistic analysis of the operational documentation — which code-switches between Simplified Chinese in internal status reports and Traditional Chinese in target-facing analysis — points to a Chinese-language operator.
This report details the framework's architecture, its confirmed impact, and the paradigm shift it represents for defenders: the era of AI-orchestrated, parallel, autonomously adaptive cyber operations against government infrastructure is a reality that will only get more dangerous, with offensive capabilities that outstrip most traditional defensive capabilities.
Attack Chain: What the Attacker Successfully Achieved
Step 1: Reconnaissance: Understanding the Target
The framework began by automatically mapping the entire government ecosystem. It downloaded and decompiled JavaScript bundles from an Angular-based government portal, extracting every embedded URL, API endpoint, OAuth client ID, and Keycloak configuration object hidden in the compiled code. From this single starting point, it identified 21 connected government systems and mapped the full national SSO architecture — 6 sub-realms, all OIDC endpoints, 2 RSA signing keys, and every supported authentication flow.
On one target alone, it discovered 36+ API endpoints spanning account management, user data retrieval, file upload, and administrative functions — many completely unauthenticated. Critically, it found that one of the systems exposed its entire user database without any authentication — thousands of employee records including names, departments, and SSO account IDs .
Step 2: SDK Harvesting: Accelerating Reconnaissance
In what appears to be a fully autonomous decision, a URL extracted from the JS bundles led the framework to a GitBook documentation portal hosting the national SSO integration guide. Using GitBook's built-in features ( llms.txt for page indexing, .md suffix for raw content), it systematically scraped the documentation, found a download link, and obtained two SDK integration example projects (Java Spring Boot and ASP.NET Core 8.0) — client-side demo applications showing how agencies connect to the national SSO.
The framework then ran an AI-powered static analysis (its Local-LLM-SAST-0day methodology) on the SDK sample code, identifying potential security issues — including an open redirect (CSRF disabled + unvalidated redirectUri parameter) and an unauthenticated authorization code-to-token exchange pattern. However, none of these findings translated into confirmed exploits against the live system. Cross-referencing the SAST results against the framework's own final validated vulnerability list shows zero overlap — the actual breaches came from server-side flaws (backdoor endpoints, unsigned JWT acceptance, unauthenticated APIs) discoverable through standard black-box testing without any SDK knowledge.
Step 3: Initial Access: Multiple Ways In
With the attack surface mapped, the framework found several entry points running in parallel:
Authentication backdoors : On a government web application, it discovered three hidden API endpoints that accept any request body and return a valid authenticated session. No credentials needed — these were developer debug endpoints left exposed in production.
Automated credential spraying : Using the employee usernames harvested from the unauthenticated users API endpoint, the framework attacked the department's office automation portal. The portal was protected by CAPTCHA, but the framework used Tesseract OCR to solve each small CAPTCHA image with 100% accuracy. It tested predictable password patterns based on each employee's ID — capitalizing or decorating the username with common symbols — and cracked 85 accounts across multiple spray rounds (an initial 12, then 73 more as additional patterns were discovered).
JWT signature bypass : A government API accepted JWT tokens with the algorithm field set to none, meaning anyone could forge valid identity tokens without knowing the signing key.
Step 4: Lateral Movement and Persistence
The cracked employee credentials from Step 3 were systematically tested against every other internal system connected via Single Sign-On. Each internal system exposed a dedicated SSO entry point that automatically trusted the office automation session (from step 3) — no additional authentication, MFA, or user confirmation required.
84 of the 85 cracked accounts (98.8%) successfully authenticated to the department's internal information system via an SSO bridge endpoint, gaining access to internal dashboards, equipment management interfaces, and personnel statistics pages. The framework automated this at scale, testing each credential against multiple connected systems in parallel.
The framework also uploaded a web shell through an unrestricted file upload API, placing it at a known path on the server. However, a secondary layer of Forms Authentication blocked execution — a partial success that stopped short of remote code execution.
Combining backdoor access, cracked credentials, and unauthenticated API endpoints, the framework extracted:
2,564+ personnel records : 1,409 employees (name, department, SSO ID), 916 users from an unauthenticated API, and 239 legal professionals from an unauthenticated Ministry of Justice endpoint
Complete user database : full JSON export of all department system users
Internal architecture : 7 SSO client secrets (rotated but naming conventions revealed), 6 internal database credentials (MSSQL, Oracle, Sybase), and internal network IP ranges
Step 6 : Supply Chain Expansion: Pivoting to Vendors
The attacker didn't stop at primary targets. It expanded the operation to government IT supply chain vendors, a nuclear safety agency, a government email system, and 7+ energy sector companies — scanning them all in parallel for misconfigurations, exposed admin interfaces, and exploitable vulnerabilities.
The Architecture: Hermes and OpenClaw
The framework operates under two workspace identifiers — .hermes and .openclaw — with reports generated into structured paths.
Each sub-agent is assigned a letter designation and a specific mission — from credential theft to API exploitation to supply chain reconnaissance. In a single attack wave, we observed up to 8 agents deployed concurrently, each targeting a different attack surface:
Across the campaign's 12 documented attack waves, we identified agents labeled A through Q, with up to 8 deployed concurrently in a single wave (dispatched in 3 batches). The sheer volume of output — 1,395 files produced in approximately four days — is consistent with heavy automation far beyond what a human operator could produce alone.
The Bayesian Brain: A Two-Layer Probabilistic Decision Engine
Perhaps the most striking feature of this framework is its use of Bayesian posterior probability scoring — not as a single metric, but as a two-layer decision engine that operates at both the individual vulnerability level and the strategic attack chain level. In other words, this was not a “spray and pray” approach but a deliberate strategically drive attack that dynamically reprioritized its approach.
Layer 1: Individual Vulnerability Scoring
The framework's triage report reveals the first layer — an autonomous triage system that scores each individual finding using a formal Bayesian model. The methodology starts every vulnerability at an uninformative prior of P=0.50, then updates based on evidence using explicit likelihood ratios:
The resulting posterior determines what happens next:
Once individual vulnerabilities are triaged, the framework connects them into multi-step attack paths and applies a second probability layer. The explicit formula from the archive:
P_success = P_chain × (1 - P_blocker)
P_chain = confirmed steps / total steps in the chain
P_blocker = probability of an insurmountable blockerScoring thresholds:
95%+ = all critical steps already confirmed and directly exploitable
60-94% = most steps confirmed, minor attempts needed
20-59% = partial confirmation, significant blockers present
5-19% = chain exists but critical conditions are missing
<5% = insurmountable blocker identified
Worked example — SSO Lateral Movement, rated 99%:
The framework assembled this chain from three separately confirmed findings: (1) an unauthenticated user list yielding thousands of employee accounts, (2) a password spray cracking multiple accounts, and (3) SSO token acquisition confirming lateral pivot capability. The precondition table:
All critical steps were already tested and working. The only theoretical blocker — the target going offline — had near-zero probability since it was a

[truncated]
