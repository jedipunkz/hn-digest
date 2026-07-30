---
source: "https://blog.google/security/chrome-stronger-with-every-update/"
hn_url: "https://news.ycombinator.com/item?id=49114031"
title: "We're making Chrome and the web safer in the AI Era"
article_title: "Stronger with every update: How we’re making Chrome and the web safer in the AI Era"
author: "xnx"
captured_at: "2026-07-30T19:09:55Z"
capture_tool: "hn-digest"
hn_id: 49114031
score: 2
comments: 1
posted_at: "2026-07-30T18:48:13Z"
tags:
  - hacker-news
  - translated
---

# We're making Chrome and the web safer in the AI Era

- HN: [49114031](https://news.ycombinator.com/item?id=49114031)
- Source: [blog.google](https://blog.google/security/chrome-stronger-with-every-update/)
- Score: 2
- Comments: 1
- Posted: 2026-07-30T18:48:13Z

## Translation

タイトル: AI 時代における Chrome とウェブの安全性を高める
記事のタイトル: アップデートごとに強化: AI 時代に Chrome とウェブをどのように安全にしているか
説明: Chrome は Gemini AI を使用して脆弱性の発見、優先順位付け、パッチ適用を自動化し、最新のセキュリティ リスクに合わせて更新を加速します。

記事本文:
メインコンテンツにスキップ
セキュリティ
アップデートするたびに強化: AI 時代に Chrome とウェブをどのように安全にするか
グローバル（英語）
アフリカ (英語)
オーストラリア (英語)
ブラジル (ポルトガル語)
カナダ (英語)
カナダ (フランス)
チェコ共和国
ドイチュラント (ドイツ)
スペイン (スペイン語)
フランス (フランセ)
ギリシャ (Ελληνικά)
インド (英語)
インドネシア (インドネシア語)
アイルランド (英語)
イタリア (イタリアーノ)
日本 (日本語)
대한국어 (韓国語)
ラテンアメリカ (スペイン語)
中東および北アフリカ (アラビア語)
メナ (英語)
オランダ
ニュージーランド (英語)
ポルスカ
ポルトガル (ポルトガル語)
ルーマニア
スヴェリゲ (スヴェンスカ)
タイ
トルコ (テュルクチェ)
台灣 (中文)
アップデートするたびに強化: AI 時代に Chrome とウェブをどのように安全にするか
Chrome が AI を使用して脆弱性の発見、優先順位付け、パッチ適用をどのように改善しているか。
私たちはソフトウェア セキュリティ業界の大規模な変化を経験しています。大規模言語モデル (LLM) は、自動化された脆弱性検出のための前例のない機能を解放し、人間のセキュリティ専門知識の限界をはるかに超えて拡張し、攻撃者の先を行くための新しいアプローチを必要としています。
これは、AI モデルを大規模に展開して、これまでよりも迅速に何百ものセキュリティ バグを発見して修正することを意味し、より高い回復力と包括的な修復を実現することを目標としています。
ソフトウェアのバグの中には、セキュリティに影響を及ぼすものもあります。純粋に機能的なバグによってイライラする UI フリーズが発生する可能性がありますが、セキュリティ バグ (または脆弱性) を利用してエクスプロイトが構築される可能性があります。エクスプロイトにより、攻撃者は被害者のコンピューター上で、個人データの読み取りや、知らないうちにマシンを制御するなどの悪意のある操作を実行できるようになります。
一度セキュリティ b

ug がコードベースに入ると、そのライフサイクルは次のように進行します。
バグ修正を含む Chrome の新しいアップデートがリリースされました。
Chrome が再起動され、アップデートが適用されます。
私たちの目標は、これらの各ステップをできるだけ早く実行することです。
Chrome セキュリティ チームは長年にわたって LLM を使用してきました。 2023 年に、私たちは LLM を使用してセキュリティ ファジングの適用範囲とパフォーマンスを向上させる方法を開発しました。 2024 年、私たちは Naptime で Project Zero と協力し、LLM に脆弱性研究用の特殊なツールを提供しました。そして 2025 年には、V8 JavaScript エンジンとグラフィックス スタックのバグを発見することに成功した AI 脆弱性発見エージェントである Big Sleep で DeepMind および Project Zero と協力しました。
2026 年の初めに、Gemini を使用して、より効率的に、より低い誤検知でより広範な Chrome コードベース全体の脆弱性を検出するエージェント ハーネスを構築しました。私たちが見つけたバグの 1 つは、侵害されたレンダラーがブラウザをだましてローカル ファイルを読み取ることを可能にするサンドボックス エスケープでした。このバグは、私たちのコードベースで 13 年以上ひっそりと生き残っていました。私たちの多くにとって、この瞬間は AI を活用した脆弱性検出の可能性を確固たるものにしました。
そこから、次のようにして脆弱性発見エージェントのハーネスを改善しました。
モデルの相互運用性のサポートを追加して、オープンウェイト モデルと独自モデルの両方の独自の強みを活用します。
以前に特定されたすべての CVE と Chrome の Git 履歴全体を含む Chrome のナレッジ ベースを構築し、トレーニング データを超えて LLM の推論能力を拡張します。
開発者に SECURITY.md ファイルを追加するよう奨励します。これは、モデルが信頼境界をより深く理解し、脅威モデルの正確なビューを開発するのに役立ちます。
これらの SECURITY.md ファイルを使用するために、別のコンテキストを持つ「クリティカル」エージェントを追加します。
脆弱性検索を実行する機能の導入

g は、モデルの非決定性と時間の経過に伴うモデルの改善を考慮して、コードベースを複数回モデル化します。
私たちはこれらすべてを安全性を念頭に置いて構築しており、AI が予期せぬ動作をするリスクを軽減するためにガードレールを設置しています。当社の AI は、一般的なインターネット アクセスができないロックダウンされたマシン上で動作し、ソース コードを厳密に保存状態で分析します。また、これらの内部スキャンには専用のセットアップを利用し、すべてのネットワーク リクエストを傍受し、開始アプリケーションと宛先に基づいた厳格な許可リストを採用し、疑わしいモデル アクティビティをブロックします。さらに、無制限モードでモデルを実行することは決してなく、サブエージェントがローカル システムを変更したり、指定されたソース コード ディレクトリ外のファイルにアクセスしたりすることを厳しく制限しています。
AI を活用した脆弱性検出は、既存のセキュリティ テスト インフラストラクチャを補完します。たとえば、ファジングは、コードベースの異なる部分間の長距離相互作用から発生するバグや、一見無関係な操作の組み合わせを必要とするバグを発見するのに特に効果的です。
また、最も困難で影響力のある脆弱性を発見する際の専門知識と創造性に対して、Chrome 脆弱性報奨プログラム (VRP) を通じて外部研究者に継続的に報奨していきたいと考えています。 2026 年の初めには、すべてのカテゴリのバグ レポートが徐々に増加していましたが、3 月までにその変化は明らかでした。2025 年全体で受け取ったバグ レポートよりも多くのバグ レポートを受け取りました。これにより、VRP を変更して、社内で発見したものに追加され、新しく自動化された処理パイプラインによって簡単に取り込めるバグの提出に研究者を集中させることにしました。
AI を利用したツールでさらに多くのセキュリティ脆弱性を発見するにつれて、私たちは同時に AI を使用して検証、トリアージ、

そしてバグの修正。これまで、単一のセキュリティ レポートの優先順位付けには 5 分から 30 分以上かかり、主に人間の専門知識に依存していました。私たちは、トリアージ プロセスを、ルールベースのシステムと AI を組み合わせてスループットと精度を向上させる自動化されたアプローチにますます移行しています。
自動トリアージ プロセスは、次の 4 つの主要なフェーズに分かれています。
ノイズをフィルタリングして除去します。システムは、受信したバグがスパムであるかどうかをチェックし、受け入れ基準を満たしていることを確認し（重複していないなど）、Chrome のセキュリティ脆弱性を明確に記述していることを検証します。
バグを再現しています。次に、システムは概念実証をチェックします。再現可能なバグは、影響を受ける特定のオペレーティング システムとブラウザのバージョンでテストされます。これに基づいて、システムはスタック トレースなどの詳細をバグに添付して、修正の通知に役立てます。
メタデータを使用してレポートを強化します。システムは、バグが最初に発生した時期や重大度評価などの重要なメタデータをレポートに追加します。このプロセスの拡大を支援するために、重大度ガイドラインをより明確にし、自動的に適用しやすくしました。開発者が重大度評価が正しくないと思われる場合は変更し、モデルが SECURITY.md ファイルを使用してセキュリティ境界を推論できるようにするためのコンテキストを追加することを引き続き許可します。
自動割り当て。システムは、問題を適切なコンポーネントと人間の所有者に自動的に転送します。
正確に測定することは困難ですが、この新しいプロセスにより月あたり数百時間の開発時間が節約され、チームが他のセキュリティ優先事項に集中できるようになったと推定されています。
Google 全体で、開発者はセキュリティ チームとセキュリティ修正の優先順位付けの責任を共有していますが、バグ発見の規模を拡大するには、同様にスケーラブルなバグ修正プロセスが必要です。
これを達成するために、私たちは mult に依存します。

i-agent のワークフロー全体:
特定の問題からコンテキストを取り込む最初のビルド ステップの後、複数の修正候補を返す修正エージェントを実行します。
次に、批評家エージェントがどれが最適かを評価し、開発者が修正を評価できるように他の関連成果物を生成します。
修正エージェントと批評エージェントは、一般的なコード レビュー プロセスを模倣したループで動作し、コードが機能し、Chromium や Google のスタイル ガイドライン、およびその他のローカル コード規約に準拠していることを確認します。
テスト作成エージェントは、修正のためのテストの作成を支援します。これらのエージェントは、開発者が修正をレビューする前に、Chrome がサポートするプラットフォームと構成全体でテストが機能することを確認できるため、開発者の時間を最大数週間節約できます。
現時点では、LLM がほとんどの脆弱性に対する修正候補を生成しており、最近の Chrome リリースでのセキュリティ修正の割合が劇的に増加しています。
最近の Chrome 安定版リリースのマイルストーンで修正されたセキュリティ バグの数
過去 2 つのマイルストーンである Chrome 149 と 150 では、1,072 件のセキュリティ バグが修正され、これまでの 23 マイルストーンで修正されたセキュリティ バグの合計数を超えています。
当社は、BigSleep や CodeMender を含め、Google DeepMind および Project Zero と長年にわたり緊密に提携してきました。これらのツールは継続的インテグレーション (CI) システムにネイティブに統合されており、すべての CL で 24 時間ごとに実行され、セキュリティ バグをプロアクティブに検出します。この統合は重要な成果をもたらしました。5 月だけで、重大な S1+ 問題を含む 20 を超える脆弱性が本番環境に到達するのをブロックしました。
修正が公開され、オープン ソース コードベースで公開されると、攻撃者は修正がユーザーのマシンに到達する前にリバース エンジニアリングを開始してバグを悪用することができます (いわゆる「N デイ」攻撃)。これは一般的に参照されます

「パッチギャップ」と誤って呼ばれます。メインの「ツリー」にコミットされた修正が Chrome Stable チャネル (大多数のユーザーが実行しているチャネル) に到達するまでに通常は数週間かかるため、このパッチのギャップを最小限に抑えることが私たちの戦略の重要な部分です。
重大度に基づいて、セキュリティ修正はメインの「ツリー」からアクティブな Chrome 安定版リリース ブランチに直接マージされ、新たなクラッシュやリグレッションを防ぐために継続的に監視されます。現在、Chrome の主要なマイルストーンを 2 週間の間隔で実施し、セキュリティ アップデートを毎週実施するように移行中です。しかし、AI を活用した急速に進む攻撃に直面して、私たちの配信ペースはさらに加速する必要があります。この瞬間に対応するために、私たちは週に 2 回のセキュリティ リリースへの移行を試験的に実施しています。
このペースであっても、適切な情報公開が最重要であることに変わりはありません。 Chrome 安定版に到達したすべてのセキュリティ バグは、内部で発見されたか外部に報告されたかに関係なく、文書化され、標準のベスト プラクティスとして公開されます。私たちは、手動によるボトルネックを排除し、脆弱性の発見から公開までの期間を短縮するために、セキュリティ バグ修正からのリリース ノートと CVE 説明の生成を自動化することに取り組んでいます。
2008 年、Chrome は、サイレントなバックグラウンド ソフトウェア アップデートの概念を先駆けて導入しました。つまり、ユーザーの介入を最小限に抑えながら、新しいバイナリが自動的にダウンロードされ、ディスク上にステージングされます。次回ブラウザを再起動すると、更新が適用され、ユーザーは保護されます。ただし、トリアージ、修正、テスト、リリースにかかる時間は 1 ～ 2 日ですが、ユーザーが Chrome を再起動するまでの待機時間は、N 日間の悪用リスクに大きく寄与する可能性があります。
Chrome の再起動を遅らせるには理解できる理由があります。再起動は混乱を招く可能性があるため、インベットのスケジュールを設定する必要があります

それが最優先のタスクであることはめったにありません。この摩擦を解消するために、私たちは次のような方法でユーザーの負担を軽減する先駆的な方法を講じています。
ほとんどの場合、ブラウザを完全に再起動する必要がなくなる「動的パッチ適用」への投資。 Chrome のマルチプロセス アーキテクチャを活用することで、動的パッチ適用により、バックグラウンドの子プロセス (レンダラーや GPU など) が更新されたバイナリでその場で順次置き換えられます。この機能の研究と開発については、引き続き詳細をご確認ください。
より多くの状態をローカルに保存することで、複雑な場合でもセッションを確実にシームレスに復元する方法を検討しています。
セッションのシームレスな復元が保証できる場合に、自動的に再起動する適切なタイミングを見つけます。たとえば、Chrome 150 では、通常、すべてのウィンドウが閉じられた後でもアプリケーションがバックグラウンドで実行され続ける macOS の固有のアプリケーション状態を利用するための変更がロールアウトされました。このウィンドウのない状態で Chrome が保留中のアップデートを検出すると、自動的に再起動されるようになりました。
macOS でのゼロウィンドウ自動再起動
私たちの長期的なビジョンは、継続的かつ動的にパッチが適用され、中断が最小限に抑えられた適切な期間に自動的に再起動される、常に最新のブラウザーであることです。現在この問題に取り組んでいる間、右上隅にある更新メッセージをクリックすることで Chrome を最新の状態に保つことができます。
エンタープライズ向け

[切り捨てられた]

## Original Extract

Chrome uses Gemini AI to automate vulnerability discovery, triage, and patching, accelerating updates to match modern security risks.

Skip to main content
Security
Stronger with every update: How we’re making Chrome and the web safer in the AI Era
Global (English)
Africa (English)
Australia (English)
Brasil (Português)
Canada (English)
Canada (Français)
Česko (Čeština)
Deutschland (Deutsch)
España (Español)
France (Français)
Greece (Ελληνικά)
India (English)
Indonesia (Bahasa Indonesia)
Ireland (English)
Italia (Italiano)
日本 (日本語)
대한민국 (한국어)
Latinoamérica (Español)
الشرق الأوسط وشمال أفريقيا (اللغة العربية)
MENA (English)
Nederlands (Nederland)
New Zealand (English)
Polska (Polski)
Portugal (Português)
România (Română)
Sverige (Svenska)
ประเทศไทย (ไทย)
Türkiye (Türkçe)
台灣 (中文)
Stronger with every update: How we’re making Chrome and the web safer in the AI Era
How Chrome is using AI to improve vulnerability discovery, triage, and patching.
We’re living through a massive shift in the software security industry. Large Language Models (LLMs) are unlocking unprecedented capabilities for automated vulnerability discovery, scaling far beyond the limits of human security expertise, and requiring new approaches for staying ahead of attackers.
This means deploying AI models at scale to find and fix hundreds of security bugs, faster than ever, with the goal of achieving greater resilience and comprehensive remediation.
Some software bugs have security implications. While a purely functional bug might result in a frustrating UI freeze, a security bug (or vulnerability) can be used to build an exploit. Exploits allow attackers to perform malicious actions on a victim’s computer, such as reading private data, or controlling their machine without their knowledge.
Once a security bug enters the codebase, its life cycle proceeds as follows:
A new update of Chrome with the bug fix is released.
Chrome is restarted and the update is applied.
Our goal is for every one of these steps to happen as quickly as possible.
The Chrome Security team has been using LLMs for years. In 2023 we developed ways to use LLMs to increase security fuzzing coverage and performance . In 2024, we worked with Project Zero on Naptime , giving LLMs specialized tools for vulnerability research. And in 2025, we collaborated with DeepMind and Project Zero on Big Sleep , an AI vulnerability discovery agent that successfully found bugs in the V8 JavaScript engine and graphics stack.
In early 2026, we built an agent harness that used Gemini to find vulnerabilities across the broader Chrome codebase with higher efficiency and lower false positives. One of the bugs we found was a sandbox escape that would allow a compromised renderer to trick the browser into reading local files — a bug that quietly survived in our codebase for more than 13 years! For many of us, this moment cemented the potential of AI-powered vulnerability detection.
From there, we improved on our vulnerability finding agent harness by:
Adding support for model interoperability to leverage the unique strengths of both open-weights and proprietary models.
Building a knowledge base of Chrome, including all previously identified CVEs and Chrome’s entire Git history, to extend the LLMs reasoning capacity past its training data.
Encouraging developers to add SECURITY.md files, which help models better understand trust boundaries and develop an accurate view of the threat model.
Adding a “critic” agent with a separate context to consume these SECURITY.md files .
Introducing the ability to run vulnerability finding models over the codebase multiple times to account for model non-determinism and model improvements over time.
We’ve built all of this with safety in mind, and have put in place guardrails to mitigate the risk of AI behaving unexpectedly. Our AI analyzes source code strictly at rest, operating on locked-down machines that lack general internet access. We also utilize a dedicated setup for these internal scans that intercepts all network requests, employing strict allowlists based on the initiating application and destination, blocking any suspicious model activity. Furthermore, we never run models in an unrestricted mode, and we strictly limit our subagents from modifying the local system or accessing files outside of designated source code directories.
AI-powered vulnerability detection complements our existing security testing infrastructure. For example, fuzzing continues to be especially effective at finding bugs that arise from long-range interactions between disparate parts of our codebase, or those requiring a combination of seemingly unrelated operations.
We also want to continue to reward external researchers for their expertise and creativity in finding the most challenging and impactful vulnerabilities via the Chrome Vulnerability Reward Program (VRP). In early 2026, we saw a gradual increase in all categories of bug reports, but by March, the shift was apparent: we received more bug reports than we had in the entirety of 2025. This led us to change our VRP to focus researchers on bug submissions that are additive to what we are finding internally, and easily ingestible by our newly automated processing pipelines.
As we discover more security vulnerabilities with AI-powered tools, we’ve simultaneously used AI to scale and automate validating, triaging, and fixing bugs. Historically, triaging a single security report took anywhere from 5 to 30 or more minutes, and relied primarily on human expertise. We have been increasingly shifting our triage process towards an automated approach that blends rule-based systems with AI to increase throughput and accuracy.
The automated triage process is broken down into four key phases:
Filtering out the noise. The system checks if an incoming bug is spam, ensures it meets intake criteria (e.g. is not a duplicate), and verifies that it clearly describes a Chrome security vulnerability.
Reproducing bugs. Next, the system checks for a proof of concept. Reproducible bugs are tested on the specific operating system and browser versions they affect. Based on this, the system attaches further details such as stack traces to the bug to help inform the fix.
Enriching the report with metadata. The system adds essential metadata to the report, such as when the bug was first introduced and its severity rating. To help this process scale, we’ve made our severity guidelines clearer and easier to apply automatically. We continue to allow developers to modify the severity rating if they believe it is incorrect, and to add context to help models reason about security boundaries using SECURITY.md files.
Automatic assigning. The system automatically routes the issue to the correct component and human owner.
While it's hard to measure precisely, we estimate that this new process is saving hundreds of hours of developer time per month, allowing our team to focus on other security priorities.
Across Google, developers share the responsibility of prioritizing security fixes with the security team, but scaling bug discovery requires an equally scalable bug fixing process.
To achieve this, we rely on multi-agent workflows throughout:
After initial build steps that bring in context from a specific issue, we run a fixing agent that returns multiple candidate fixes.
A critic agent then evaluates which would be the best fit, producing other relevant artifacts for developers to evaluate the fix.
The fixing and critic agents work in a loop that mimics a typical code review process to ensure that code is functional and compliant with Chromium and Google style guidelines, as well as other local code conventions.
Test-writing agents help write tests for fixes. These agents can ensure that tests work across the full array of Chrome supported platforms and configurations before a developer reviews the fix, saving up to weeks of developer time.
At this point, we have LLMs generating candidate fixes for most vulnerabilities, dramatically increasing the rate of security fixes in recent Chrome releases:
Number of security bugs fixed in recent Chrome Stable release milestones
In the last two milestones, Chrome 149 and 150, we have fixed 1072 security bugs, surpassing the total number of security bugs fixed across the prior 23 milestones combined.
We have partnered closely with Google DeepMind and Project Zero for years, including on BigSleep and CodeMender. These tools are natively integrated into our continuous integration (CI) system, running every 24 hours across all CLs to proactively detect security bugs. This integration has yielded significant results: in May alone, we blocked over 20 vulnerabilities from reaching production, including a critical S1+ issue.
Once a fix has landed and is visible in the public open source codebase, attackers can start to reverse engineer and exploit the bug before the fix reaches users’ machines — so called "N-day" attacks. This is commonly referred to as the “patch gap.” Since fixes committed to the main “tree” typically take weeks to reach the Chrome Stable channel (what the vast majority of our users run), minimizing this patch gap is a critical part of our strategy.
Based on their severity, security fixes are merged directly from the main “tree” into the active Chrome stable release branch, which is continuously monitored to prevent new crashes or regressions. We are in the process of transitioning to a two-week cadence for major Chrome milestones , with weekly security updates. However, in the face of fast-moving, AI-powered attacks, our delivery cadence must accelerate even further. To meet this moment, we are piloting a shift to two security releases per week.
Even with this pace, proper public disclosure remains paramount. Every security bug that reaches Chrome Stable, regardless of whether it was discovered internally or reported externally, is documented and disclosed publicly as a standard best practice. We are working on automating the generation of release notes and CVE descriptions from security bug fixes to eliminate manual bottlenecks and shorten the window between vulnerability discovery and public disclosure.
In 2008, Chrome pioneered the concept of silent, background software updates: new binaries are automatically downloaded and staged on disk with minimal user intervention. At the next restart of the browser, the update would be applied and the user would be protected. However, compared to the 1–2 days it takes for triage, fix, test and release, the time spent waiting for the user to restart Chrome can be a significant contributor to N-day exploitation risk.
People have understandable reasons to delay restarting Chrome. A restart can be disruptive, requires scheduling in-between tasks, and is rarely the top priority at any given moment. To eliminate this friction, we are pioneering ways to shift the burden away from the user by:
Investing in "dynamic patching" that will eliminate the need for a full browser restart in most cases. By leveraging Chrome’s multi-process architecture, dynamic patching sequentially replaces background child processes (like the Renderer and GPU) with updated binaries on the fly. Stay tuned to learn more as we research and develop this feature.
Exploring ways to ensure a seamless session restore even in complex cases, by saving more state locally.
Finding opportune moments to restart automatically, when we can guarantee a seamless session restore. For example, in Chrome 150, we rolled out a change to take advantage of the unique application state on macOS where applications typically continue running in the background even after all windows are closed. Now, if Chrome detects a pending update while in this windowless state, it automatically restarts.
Zero window auto-restart on macOS
Our long-term vision is a browser that is always up-to-date – continuously and dynamically patched, and automatically restarted during opportune periods of minimal disruption. While we’re working on this, you can keep your Chrome up to date by clicking on the update message in the top right corner.
For enterprise

[truncated]
