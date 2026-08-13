---
source: "https://www.bleepingcomputer.com/news/security/who-vets-ais-code-the-scale-challenge-facing-open-source-ingestion/"
hn_url: "https://news.ycombinator.com/item?id=49287586"
title: "Who Vets AI's Code? The Scale Challenge Facing Open Source Ingestion"
article_title: "Who Vets AI’s Code? The Scale Challenge Facing Open Source Ingestion"
author: "DemiGuru"
captured_at: "2026-08-13T15:49:56Z"
capture_tool: "hn-digest"
hn_id: 49287586
score: 2
comments: 0
posted_at: "2026-08-13T15:34:53Z"
tags:
  - hacker-news
  - translated
---

# Who Vets AI's Code? The Scale Challenge Facing Open Source Ingestion

- HN: [49287586](https://news.ycombinator.com/item?id=49287586)
- Source: [www.bleepingcomputer.com](https://www.bleepingcomputer.com/news/security/who-vets-ais-code-the-scale-challenge-facing-open-source-ingestion/)
- Score: 2
- Comments: 0
- Posted: 2026-08-13T15:34:53Z

## Translation

タイトル: AI のコードを精査するのは誰ですか?オープンソースの取り込みが直面するスケールの課題
記事のタイトル: AI のコードを精査するのは誰ですか?オープンソースの取り込みが直面するスケールの課題
説明: AI コーディング ツールは、従来のセキュリティ レビューが追いつくよりも早く、精査されていない、または幻覚に満ちたオープンソースの依存関係を導入する可能性があります。 ActiveState は、組織が開発パイプラインに入る前の選択時点でパッケージを管理する必要がある理由を説明しています。

記事本文:
AI のコードを精査するのは誰ですか?オープンソースの取り込みが直面するスケールの課題
Plug and Pwn 攻撃は Windows SYSTEM アクセスに偽の USB デバイスを使用します
Lazarus ハッカーが Windows ゼロデイを悪用して防衛企業を標的に
FBI：ハッカーがオンラインアカウントを標的にしてヌード写真を盗む
新しい Microsoft Defender 'ShieldBreak' ゼロデイは SYSTEM 権限を付与します
Trezor、約 14,000 人の顧客に影響を与えるデータ侵害を明らかに
AI のコードを精査するのは誰ですか?オープンソースの取り込みが直面するスケールの課題
ホワイトハウス、攻撃的なハッキング作戦でセキュリティ会社を起用
WhatsApp、潜在的な詐欺メッセージにフラグを立てる新機能を公開
Tor ブラウザを使用してダークウェブにアクセスする方法
Windows 11 でカーネルモードのハードウェア強制スタック保護を有効にする方法
Windows レジストリ エディタの使用方法
Windows レジストリをバックアップおよび復元する方法
Windows をセーフ モードで起動する方法
トロイの木馬、ウイルス、ワーム、その他のマルウェアを削除する方法
Windows 7で隠しファイルを表示する方法
Windows で隠しファイルを確認する方法
AI のコードを精査するのは誰ですか?オープンソースの取り込みが直面するスケールの課題
AI のコードを精査するのは誰ですか?オープンソースの取り込みが直面するスケールの課題
ActiveState 製品責任者 Jonny Rivera 著
先週、Black Hat の展示会場で私たちが会話している間、私たちのチームが AppSec リーダー、プラットフォーム エンジニア、CISO と交わしたほぼすべての議論で、ある疑問が浮上しました。それは、AI のコードを実際に精査しているのは誰ですか?
開発者による AI コーディング ツールの導入は減速していません。生産性の向上は現実のものであり、オープンソース ソフトウェアは依然として最新のエンタープライズ アプリケーションのバックボーンです。しかし、AI コーディング アシスタントがサードパーティの依存関係の提案をミリ秒単位で自動補完するため、企業のセキュリティ チームとオープンソースのメンテナーは共通の運用上の課題に直面しています。コード生成のペースが従来の Inge を完全に上回っているということです。

レビュー。
未調査の依存関係または幻覚のある依存関係がマシンの速度でコードベースに入ると、コミット後のソフトウェア構成分析 (SCA) スキャンが追いつくのに苦労します。
このパイプラインを保護することは、開発者の速度を低下させたり、オープンソースを制限したりすることを意味するものではありません。インポートによってビルドがトリガーされる前に、選択の時点で環境に何が入るかを管理する必要があります。
「ずぼらしゃがみ」のメカニズムと機械による摂取
大規模言語モデル (LLM) は、リアルタイムのパッケージ レジストリ検証ではなく、統計的確率と履歴コード パターンに基づいてソフトウェア ライブラリを推奨します。
モデルが PyPI や npm に存在しないパッケージ名を提案すると、スロップスクワッティング (または AI パッケージの幻覚悪用) として知られるサプライ チェーンの脆弱性が作成されます。
この脆弱性ベクトルの規模は、500,000 以上のコード サンプルにわたる 16 の一般的なコード生成モデルを分析した USENIX セキュリティ調査で明らかになりました。
AI が提案するパッケージ名のかなりの割合がパブリック レジストリに存在しません。
実際のパッケージに解決される推奨される依存関係のうち、ほぼ半数には既知の CVE または古いリリースが含まれています。
攻撃者は、公開されている LLM 出力パターンと開発者コード リポジトリを定期的に監視して、これらの幻覚パッケージ名を特定します。敵対者が特定されると、ダミー名を PyPI または npm に登録し、悪意のあるペイロードをアップロードし、自動化された開発環境または CI/CD ビルダーがそれを取得するのを待ちます。
[開発者ワークスペース] ---> [AI アシスタントのオートコンプリート パッケージ名]
|
v
[パッケージ名がレジストリに存在しません]
|
v
[攻撃者がペイロードを含むPyPI/npmに名前を登録]
|
v
[CI/CD パイプラインフェッチパッケージ] ---> [侵害されたビルド]
このベクターは、野生環境での展開において積極的に観察されています。 2026 年の初めに、セキュリティ研究者は単一のセキュリティ セキュリティを追跡しました。

単一のコミットで AI が生成した 47 のエージェント スキルに由来する、npm パッケージ名 (react-codeshift) を示します。
この幻覚は、人間が明示的に選択していないことにエンジニアが気づくまで、フォークを通じて 230 以上のリポジトリに有機的に広がりました。
この問題は開発者による悪意ではなく、取り込み制御が完全に欠如していたことによるものでした。
AI パッケージの幻覚をビルドに影響を与える前に阻止する
AI コーディング アシスタントはマシンの速度でソフトウェアを生成しますが、精査されていない依存関係により、パイプラインが不法占拠やサプライ チェーン攻撃にさらされることになります。
ActiveState は、ソースから構築されたクリーンなコンポーネントの安全なリポジトリを利用して、組織がソフトウェアの来歴とビルド レベルの証明を証明できると同時に、導入ステップで不法侵入ベクトルを排除します。
オープンソースレビューにおける摩擦係数
企業内の摂取の課題は、より広範なオープンソース エコシステムに直接影響を与えます。エンタープライズ ネットワーク内で精査されていない依存関係の提案を生成する同じ AI アシスタントが、コミュニティが管理するリポジトリに送信される自動プル リクエストも生成します。
この大量の自動化されたコントリビューションは、人間のメンテナに前例のない負担を与えます。
矛盾する AI ポリシー: Kubernetes、Linux カーネル、LLVM、Godot などの主要プロジェクトは、AI 支援による貢献に関して異なるポリシーを公開しています。 AI 生成コードを完全に禁止するところもあれば、人間の投稿者が追加されたすべての行に対して完全な責任を負う場合にのみ許可するところもあります。
より高い欠陥密度: CodeRabbit が 470 件のオープンソース プル リクエストをレビューしたところ、表面的にはきれいに見えたにもかかわらず、AI が共同作成したコントリビュートには人間が作成したコードよりも 70% 多くの欠陥が含まれていることがわかりました。
幻覚を起こしたパッケージや脆弱なパッケージが企業の摂取を通過すると、それらは必然的に滴下します。

アップストリームのオープンソース PR に参加し、ボランティアのメンテナーは人間が意図的に評価しなかった依存関係の検証に何時間も費やすことを強いられます。
ベロシティと検証: ガバナンスのギャップ
Kusari の Application Security in Practice レポートの最近のテレメトリは、ツールの導入が取り込み制御をどの程度超えているかを示しています。
AIコーディングアシスタントを使用している組織
AI を使用して PR 段階のコードレビューを支援する組織
専用の AI AppSec コントロールを備えた組織
従来の AppSec ワークフローは、コードの作成後、またはプル リクエストのオープン後にコードをスキャンすることに依存しています。コードがマシン速度で生成されると、最終段階のアラートはエンジニアが無視するバックログ ノイズを生成するだけです。
選択時点でのパイプラインの保護
LLM 幻覚率がゼロに低下するのを待つことは、AppSec 戦略ではありません。中心的な問題は速度であり、モデルの精度ではありません。出力を犠牲にすることなく開発パイプラインを保護するために、セキュリティ チームとプラットフォーム チームは防御を IDE の左側に移動させています。
レジストリの直接取得を制限する: 開発者ワークステーションと AI エージェントが、コード補完中に未調査のパブリック エンドポイントに直接クエリを実行するのをブロックします。
AI が提案する依存関係を分離する: 新たに導入された依存関係を、プライマリ ブランチに許可する前に、到達可能性と脆弱性の自動分析のために分離されたサンドボックスにルーティングします。
インジェスト ゲートウェイの管理: 事後対応的な CVE カウントからプロアクティブなソース キュレーションに移行し、AI モデルが推奨するすべてのパッケージが悪意のあるタイプポスクワットやスロップスクワッティングのターゲットに対して事前に精査されるようにします。
この取り込みレイヤーは、まさに ActiveState のセキュア オープン ソース ライブラリと厳選されたカタログが動作する場所です。エンタープライズ グレードの取り込みゲートウェイとして機能するように設計された ActiveState は、事前に精査され、継続的に修正されるオープン ソース パッケージを提供します

開発者ワークステーション、CI/CD パイプライン、AI エージェント環境に直接接続できます。
パブリック パッケージ レジストリと開発者ツールの間に位置することにより、厳選されたカタログにより、パッケージの幻覚リスクが選択境界で確実に遮断されます。
管理された取り込みソース上で実行されているエンタープライズ チームは、取り込みステップで不法侵入ベクトルを排除し、開発者に AI アシスタントをオフにすることを強いることなく、全体的な CVE エクスポージャを約 95% 削減します。
AI コーディング ツールを無効にすることは現実的ではなく、競争力もありません。ただし、ソフトウェア サプライ チェーンの取り込みルールを更新せずに、AI 統合を開発者の生産性指標として純粋に扱うと、運用ビルドは自動侵害に対して脆弱なままになります。
最新の開発パイプラインを保護するには、開発者またはエージェントが選択したすべてのパッケージがビルドに入る前にデフォルトで管理されるようにする必要があります。
ActiveState がオープン ソース アプリケーションの保護にどのように役立つかに興味がある場合は、今すぐデモをスケジュールしてください。
Jonny Rivera はプロダクト リーダーであり、そのキャリアはサイバーセキュリティ、デジタル ヘルスケア ソリューション、開発者ツールに及びます。彼は中学 2 年生のダイナモを持つ誇り高き演劇人の父親であり、World of Warcraft で出会った最愛の人と結婚して 18 年になります。
ActiveState によってスポンサーおよび執筆されました。
Microsoft 2026 年 8 月のパッチ火曜日で 400 件の欠陥、3 件のゼロデイを修正
デルタ航空、DEF CON参加者を乗せた航空機へのWi-Fi認証解除攻撃を調査
LexisNexis、サーバー上で不審なアクティビティが発生したためサービスを停止
Pixellot は、管理されていない数百の AI エージェント ID を数か月ではなく数週間で発見し、保護しました。その方法については、ケーススタディをダウンロードしてください。
パスワードのヘルスチェックの期限を過ぎましたか? Active Directory を無料で監査する
AI のずさんな攻撃を阻止します。オープンソース パッケージがビルドに到達する前に安全に取り込む

ld。
パスワードのヘルスチェックの期限を過ぎましたか? Active Directory を無料で監査する
本物の受信箱、偽のストア、AI 詐欺が 2026 年上半期のサイバー脅威をどのように形作ったかをご覧ください
利用規約 - プライバシー ポリシー - 倫理声明 - アフィリエイトの開示
著作権 @ 2003 - 2026 Bleeping Computer ® LLC - 全著作権所有
まだメンバーではありませんか?今すぐ登録
どのようなコンテンツが禁止されているかについては、投稿ガイドラインをお読みください。

## Original Extract

AI coding tools can introduce unvetted or hallucinated open source dependencies faster than traditional security reviews can keep pace. ActiveState explains why organizations should govern packages at the point of selection, before they enter the development pipeline.

Who Vets AI’s Code? The Scale Challenge Facing Open Source Ingestion
Plug and Pwn attack uses fake USB devices for Windows SYSTEM access
Lazarus hackers exploited Windows zero-day to target defense firms
FBI: Hackers target online accounts to steal nude photos
New Microsoft Defender 'ShieldBreak' zero-day grants SYSTEM privileges
Trezor discloses data breach affecting nearly 14,000 customers
Who Vets AI’s Code? The Scale Challenge Facing Open Source Ingestion
White House taps security firms for offensive hack-back operations
WhatsApp rolls out new feature that flags potential scam messages
How to access the Dark Web using the Tor Browser
How to enable Kernel-mode Hardware-enforced Stack Protection in Windows 11
How to use the Windows Registry Editor
How to backup and restore the Windows Registry
How to start Windows in Safe Mode
How to remove a Trojan, Virus, Worm, or other Malware
How to show hidden files in Windows 7
How to see hidden files in Windows
Who Vets AI’s Code? The Scale Challenge Facing Open Source Ingestion
Who Vets AI’s Code? The Scale Challenge Facing Open Source Ingestion
By Jonny Rivera, Head of Product at ActiveState
During our conversations on the show floor at Black Hat last week, one question came up in almost every discussion our team had with AppSec leads, platform engineers, and CISOs: Who is actually vetting AI’s code?
Developer adoption of AI coding tools isn't slowing down. The productivity gains are real, and open source software remains the backbone of modern enterprise applications. But as AI coding assistants auto-complete third-party dependency suggestions in milliseconds, enterprise security teams and open source maintainers face a shared operational challenge: code generation has completely outpaced legacy ingestion review.
When an unvetted or hallucinated dependency enters a codebase at machine speed, post-commit Software Composition Analysis (SCA) scans struggle to keep pace.
Securing this pipeline doesn't mean slowing developers down or restricting open source. It requires governing what enters the environment at the point of selection, before an import ever triggers a build.
The Mechanics of "Slopsquatting" and Machine Ingestion
Large language models (LLMs) recommend software libraries based on statistical probability and historical code patterns, not real-time package registry verification.
When a model suggests a package name that does not exist in PyPI or npm, it creates a supply-chain vulnerability known as slopsquatting (or AI package hallucination exploitation).
The scale of this vulnerability vector was highlighted in a USENIX Security study analyzing sixteen popular code-generation models across 500,000+ code samples:
A measurable percentage of AI-suggested package names do not exist in public registries.
Of the suggested dependencies that do resolve to real packages, nearly half contain known CVEs or outdated releases.
Attackers routinely monitor public LLM output patterns and developer code repositories to identify these hallucinated package names. Once identified, an adversary registers the dummy name on PyPI or npm, uploads a malicious payload, and waits for automated developer environments or CI/CD builders to fetch it.
[Developer Workspace] ---> [AI Assistant Auto-completes Package Name]
|
v
[Package Name Does Not Exist in Registry]
|
v
[Attacker Registers Name on PyPI/npm with Payload]
|
v
[CI/CD Pipeline Fetches Package] ---> [Compromised Build]
This vector is actively being observed in wild deployment. Early in 2026, security researchers tracked a single hallucinated npm package name (react-codeshift) originating from 47 AI-generated agent skills in a single commit.
The hallucination spread organically through forks to over 230 repositories before an engineer noticed a human had never explicitly selected it.
The issue was not malicious intent by the developer, rather, a complete absence of ingestion controls.
Stop AI Package Hallucinations Before They Hit Your Build
AI coding assistants generate software at machine speed, but unvetted dependencies expose your pipeline to slopsquatting and supply chain attacks.
Powered by a secure repository of clean, built-from-source components, ActiveState lets organizations prove software provenance and build-level attestation while eliminating slopsquatting vectors at the intake step.
The Friction Multiplier on Open Source Review
The intake challenge inside the enterprise directly impacts the broader open source ecosystem. The same AI assistants generating unvetted dependency suggestions inside enterprise networks are also generating automated pull requests submitted to community-maintained repositories.
This volume of automated contributions puts unprecedented strain on human maintainers:
Conflicting AI Policies: Major projects, including Kubernetes, the Linux kernel, LLVM, and Godot, have published diverging policies on AI-assisted contributions. While some ban AI-generated code outright, others permit it only if a human contributor takes full accountability for every line added.
Higher Defect Density: A CodeRabbit review of 470 open-source pull requests found that AI-co-authored contributions carried 70% more defects than human-authored code, despite reading clean on the surface.
When hallucinated or vulnerable packages pass through corporate ingestion, they inevitably trickle down into upstream open source PRs, forcing volunteer maintainers to spend hours validating dependencies that no human deliberately evaluated.
Velocity vs. Verification: The Governance Gap
Recent telemetry from Kusari’s Application Security in Practice report illustrates how far tooling deployment has outrun ingestion controls:
Organizations using AI coding assistants
Organizations using AI to assist PR-stage code review
Organizations with dedicated AI AppSec controls
Traditional AppSec workflows rely on scanning code after it is written or after a pull request is opened. When code is generated at machine speed, late-stage alerts simply create backlog noise that engineers ignore.
Securing the Pipeline at the Point of Selection
Waiting for LLM hallucination rates to drop to zero is not an AppSec strategy. The core issue is velocity, not model accuracy. To secure the development pipeline without sacrificing output, security and platform teams are moving defense left of the IDE:
Restrict Direct Registry Fetching: Block developer workstations and AI agents from querying unvetted public endpoints directly during code completion.
Isolate AI-Suggested Dependencies: Route newly introduced dependencies into an isolated sandbox for automated reachability and vulnerability analysis before allowing them into primary branches.
Govern the Ingestion Gateway: Shift from reactive CVE counting to proactive source curation, ensuring that every package an AI model recommends is pre-vetted against malicious typosquats and slopsquatting targets.
This ingestion layer is precisely where ActiveState’s Secure Open Source Library and Curated Catalog operates. Designed to function as an enterprise-grade ingestion gateway, ActiveState delivers pre-vetted, continuously remediated open source packages directly to developer workstations, CI/CD pipelines, and AI agent environments.
By sitting between public package registries and developer tools, a curated catalog ensures that hallucinated package risks are intercepted at the selection boundary.
Enterprise teams running on a governed ingestion source eliminate slopsquatting vectors at the intake step, reducing overall CVE exposure by roughly 95% without forcing developers to turn off their AI assistants.
Disabling AI coding tools is neither practical nor competitive. However, treating AI integration purely as a developer productivity metric, without updating software supply chain ingestion rules, leaves production builds vulnerable to automated compromise.
Securing the modern development pipeline requires ensuring that every package selected by a developer or an agent is governed by default before it ever hits a build.
If you’re interested in seeing how ActiveState can help secure your open source applications, schedule a demo today .
Jonny Rivera is a product leader whose career spans cybersecurity, digital healthcare solutions, and developer tooling. He is a proud theatre dad to an 8th grade dynamo and has been married for 18 years to the love of his life which he met on World of Warcraft.
Sponsored and written by ActiveState .
Microsoft August 2026 Patch Tuesday fixes 400 flaws, 3 zero-days
Delta probes Wi-Fi deauth attack on flight carrying DEF CON attendees
LexisNexis shuts down services after suspicious activity on servers
Pixellot discovered and secured hundreds of unmanaged AI agent identities in weeks, not months. Download the case study for how.
Overdue a password health-check? Audit your Active Directory for free
Stop AI slopsquatting attacks. Secure open source package ingestion before it hits your build.
Overdue a password health-check? Audit your Active Directory for free
See how real inboxes, fake stores and AI scams shaped H1 2026 cyber threats
Terms of Use - Privacy Policy - Ethics Statement - Affiliate Disclosure
Copyright @ 2003 - 2026 Bleeping Computer ® LLC - All Rights Reserved
Not a member yet? Register Now
Read our posting guidelinese to learn what content is prohibited.
