---
source: "https://claude.com/blog/using-llms-to-secure-source-code"
hn_url: "https://news.ycombinator.com/item?id=48353982"
title: "Using LLMs to secure source code"
article_title: "Using LLMs to secure source code | Claude"
author: "maxloh"
captured_at: "2026-06-02T00:43:39Z"
capture_tool: "hn-digest"
hn_id: 48353982
score: 1
comments: 0
posted_at: "2026-06-01T08:17:50Z"
tags:
  - hacker-news
  - translated
---

# Using LLMs to secure source code

- HN: [48353982](https://news.ycombinator.com/item?id=48353982)
- Source: [claude.com](https://claude.com/blog/using-llms-to-secure-source-code)
- Score: 1
- Comments: 0
- Posted: 2026-06-01T08:17:50Z

## Translation

タイトル: LLM を使用してソース コードを保護する
記事のタイトル: LLM を使用してソース コードを保護する |クロード
説明: Claude Opus と協力して脅威モデルを構築し、コードベースの脆弱性を発見し、検証、優先順位付け、パッチを適用する方法に関するベスト プラクティスを共有します。

記事本文:
LLM を使用してソース コードを保護する |クロード
クロード製品のご紹介 クロード
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
LLM を使用してソース コードを保護する
LLM を使用してソース コードを保護する
Claude Opus と協力して脅威モデルを構築し、コードベースの脆弱性を発見し、検証、優先順位付け、パッチを適用する方法に関するベスト プラクティスを共有します。
リンクをコピーして共有 https://claude.com/blog/using-llms-to-secure-source-code
モデルの機能は急速に、そして不均一に進化しています。私たちはセキュリティ チームと協力して、独自のコードやオープン ソース ソフトウェアの脆弱性を見つけて修正してきました。その作業により、モデルを使用してソース コードを保護する方法についての理解を深めることができました。主なポイントは、検出の並列化が簡単になり、ボトルネックが検証、トリアージ、パッチ適用に移行したことです。
この不一致を示すために、オープンソース ソフトウェアの独自のスキャンの一環として、2026 年 5 月 22 日の時点で、1,596 件の脆弱性を公開しました。私たちの知る限り、そのうち 97 件にはパッチが適用されています。
このガイドでは、Claude Opus と協力して脅威モデルを構築し、コードベースの脆弱性を発見し、検証、優先順位付け、パッチを適用する方法について説明します。すべての答えがあるわけではありませんが、チームが発見をどのようにスケールしたか、そして後の段階で何が役立ったかを共有します。付属のリポジトリを今すぐ利用してください。

インタラクティブなワークフローのスキルと自律スキャンのデモ ハーネスを習得します。読みながら、各ステップを実装するスキルを呼び出します。
ほとんどの脆弱性を発見して修正するチームは、既存のベスト プラクティスのバリエーションに集中しました。これらを一連の 6 つのステップに抽出しました。
脅威モデル: スキャンを開始する前に、何を脆弱性としてカウントするかを決定します。
サンドボックス: サンドボックス環境を構築してエージェントを隔離し、エクスプロイトを証明します。
検出: モデルにソース コード内の脆弱性を検索させます。
検証: どの発見結果が実際に悪用可能であるかを独立して確認します。
トリアージ: 検出結果の重複を排除し、重大度を割り当て、修正が必要なものに優先順位を付けます。
パッチ適用: 修正を適用し、脆弱性が解消されたことを確認し、亜種を検索します。
最初の 2 つのステップ (脅威モデルとサンドボックスの構築) は、ループの残りの部分のセットアップです。これらは通常、コードベースごとに 1 回実行され、基礎となるシステムが変更されたときに再検討されます。次の 4 つのステップは、ソースに対して実行するループ (検出、検証、優先順位付け、パッチ) です。
通常、コードベースでの最初の実行では、最も多くの結果が得られます。より単純な脆弱性は以前の実行でパッチされているため、後続の実行では、より複雑な脆弱性が少なくなる傾向があります。ただし、n 回目の実行で新しい発見がまったくないことを期待しないでください。モデルは確率的であり、大規模なコードベースには、コードが変更されていない場合でも脆弱性のロングテールが残り続ける可能性があります。
コードベースの最初の反復では、ループを複数回実行し、まったく新しい検出結果の数とそのシステムのリスク許容度に基づいて停止するタイミングを決定する必要があります。最初の反復の後、(1) 定期的に、または (2) コードが大幅に変更されるたびにスキャンを続けます。
次は歩いて行こう

各ステップを大まかに詳細に説明し、それが重要な理由、それが何を生み出すか、そしてそれをどのように実装するかを説明します。
1. 脅威モデル: 何を脆弱性としてカウントするかを定義する
誤検知の最も一般的な原因は、モデルが信頼境界を十分に理解していないことです。このモデルでは、環境内でこれらの入力が信頼されている場合でも、クライアントが破損した値を送信したり、攻撃者が構成を制御したりできると想定しているため、コードに脆弱性のフラグを付ける可能性があります。逆に、モデルはインターネットに接続されたサービスが内部専用であると想定しているため、実際の脆弱性が過小報告される可能性があります。どちらの場合も、モデルはコードではなく脅威モデルに関して間違っています。
あるチームは、調査結果全体にパターンがあることに気付きました。モデルは、十分に文書化された脅威モデル、システム設計文書、要件、制約を備えたシステムで最高のパフォーマンスを発揮しました。脅威モデルが明確に定義されている場合、モデルの結果は「90% の確率で悪用可能でした」。
クロードと協力して、次の 2 つのステップで脅威モデルを構築できます。
まず、コード、ドキュメント、脆弱性履歴からブートストラップします。新しいセキュリティ エンジニアに初日に渡すもの (アーキテクチャ ドキュメント、Wiki、エントリ ポイント、Git 履歴、過去の脆弱性) をモデルに入力します。これは、コードのみから暗黙の知識、トレードオフ、設計上の決定を推測するという課題を克服するのに役立ちます。次に、システム コンテキスト、資産、エントリ ポイント、信頼境界を含む脅威モデルを作成するようにモデルに依頼します。最後に、モデルに過去のバグをクラスタリングさせ、関連する脆弱性クラスをリストします。脅威モデルには、どのような脆弱性を考慮し、どのような脆弱性を考慮しないのか、またその理由が文書化されていることを確認してください。
あるチームは、過去の何百もの CVE とセキュリティ修正のコミットをレビューし、それらを「バグの形」のヒントに抽出し、モデルに 2 つの質問をしました。修正は完了したか、

そしてそれは他の場所にも適用されましたか？彼らは 1 時間で 3 つの悪用可能な問題を発見しました。彼らの言葉を借りれば、「『過去に悪用されたもの』は、『このコードベースの脆弱性を見つけてください』よりも、成功に向けたはるかに簡単なチートコードである場合があります。」
次に、モデルにシステムをよく知っている人にインタビューしてもらいます。ショスタックの 4 つの質問を考えてみましょう: 私たちは何を構築しているのでしょうか?何が問題になるのでしょうか?それについて私たちは何をしているのでしょうか？私たちは良い仕事をしましたか？インタビュー対象者が最初から始めないように、最初にブートストラップ ステップを実行します。こうすることで、脅威モデルの調査と構築に時間を費やすことなく、草案から始めることができます。インタビュー ステップはオプションですが、モデルがコードやドキュメントから取得できないコンテキストが追加され、脅威モデルが改善されます。
いくつかの実践により、大きな違いが生まれます。
依存関係のセキュリティ ポリシーを考慮してください。多くのオープンソース プロジェクトがこれを公開しています。たとえば、vLLM の security.md 、SQLite の「Defence Against the Dark Arts」、ImageMagick のセキュリティ ポリシーなどです。脅威モデルでは、ポリシーを最初から再構築するのではなく、それらを直接考慮する必要があります。
信頼できるものに名前を付けます。構成ファイルまたは認証されたクライアントを信頼する場合は、それを脅威モデルに文書化します。これらの仮定は、悪用不可能なバグを実際の悪用から区別するのに役立ちます。
コードに THREAT_MODEL.md を含めます。これをリポジトリに保存し、コードの変更に応じて更新します。検出エージェントは、既知の非問題をスキップして、検索する前にそれを読み取ることができます。
脅威モデルは 2 つの場所で使用します。ディスカバリーでは、スコープとして、コードを分割し、ターゲットに優先順位を付け、スコープ外のものをスキップします。これは、全体をスキャンできない大規模なコードベースに役立ちます。トリアージでは、フィルターとして: 広範囲にスキャンした後、脅威モデルを使用して、システムと環境に対する重大度をより適切に調整します。

大規模なプロジェクトをスキャンしたあるチームでは、誤検知率が 40% であり、その理由を詳しく調査しました。調査結果には再現性があり、PoC によって悪用可能性が証明されました。しかし、コードを所有していた開発チームは、バグがプロジェクトの脅威モデルに適合しないため、誤検知として無視しました。別のチームの CISO は、「（モデルには）コードの適切なコンテキストが含まれていますが、私たちにとっては適切なコンテキストはありません」と簡潔に述べました。
脅威モデル スキルを試してください。このセクションで説明する両方の手順を説明します。ブートストラップはコード、CVE、git 履歴からドラフトを導き出し、インタビューではシステム所有者に Shostack の 4 つの質問を説明してドラフトを洗練させます。出力は、検出および問題切り分けのステップで使用される THREAT_MODEL.md ファイルです。
2. サンドボックス: エージェントを安全に実行し、悪用可能性を検証する
サンドボックスの目的の 1 つは、システムを保護することです。モデルを安全かつ自律的に実行できるようにするには、強力な分離層が必要です。これがないと、エージェントがターゲットをオーバーシュートし、予期しない動作を行う可能性があります。
あるチームはモデルにネットワーク アクセスがないことを伝えましたが、実際にはアクセスできましたが、モデルはとにかく GitHub からフェッチできることを発見しました。別のチームは、エージェントがスキャン中に GitHub の問題に回答しているのを観察しました。どちらのアクションも悪意のあるものではありませんでしたが、どちらもコードと構成を介して制約を強制する必要があることを示していました。
隔離を脅威モデルに合わせます。コンテナーは検出エージェントがコードを読み取るのには適していますが、ターゲットとその PoC を microVM (Firecracker など) で実行するか、実稼働システムに何も到達できないように出口がロックされた完全な VM で実行します。また、エージェントが使用できる資格情報 ( ~/.aws 、 ~/.ssh 、 .env ) を決して持たないでください。
サンドボックス ネットワークへのアクセスは、セットアップ中にのみ許可してください。依存関係を取得し、ツールをビルド、インストールし、ターゲットをデプロイし、既存のテストを実行してすべてが機能することを確認します。

次に、環境のスナップショットを作成し、そのネットワーク アクセスを削除します。スキャン中は、ローカル プロキシ経由でルーティングされるモデル API へのトラフィックのみを許可します。各実行の開始時にスナップショットをロードして、すべてのスキャンが同じ白紙の状態から開始されるようにします。
サンドボックスのもう 1 つの目的は、悪用可能性を証明することです。静的スキャン中、モデルはコードを読み取り、何が壊れる可能性があるかを仮説化しますが、パスが到達可能かどうか、または補償するコントロールがあるかどうかをテストすることはできません。その結果、モデルは、実際には気にしていない、悪用できないコードの正確さのバグにフラグを立てる可能性があります。チームがエージェントがコードをコンパイルし、テストを実行し、概念実証を実行できるサンドボックスを構築すると、悪用不可能な発見は大幅に減少しました。
ある攻撃セキュリティ チームは、エージェントにテスト ベッドを提供するハーネスを構築しました。このハーネスは、エージェントが概念実証を構築し、テスト ベッドで実行できた場合にのみ真陽性となるという単純な検証ルールを備えていました。 6 週間後の彼らの評価は、「最大の効果要因は、モデルにテストベッド、ライブシステムを提供し、PoC を実行することだった」というものでした。
サンドボックスを構築するときは、すべての実行で同じ環境で同じコード (イメージ タグ、コミット SHA、依存関係、ビルド コマンド) が使用されるように、できる限り固定してください。ビルドにネットワークを必要としないようにローカル コピーをキャッシュし、複数のテスト ループでロードできるようにコンテナーの耐久性を目指します。
あるチームのスキャンにより、実際にデプロイされたものではなく、エージェントが古いバージョンのライブラリをダウンロードしたことによる副産物であることが判明した脆弱性のフラグが立てられました。これは、トランスクリプトを読んだエンジニアによって発見され、別の依存関係がダウンロードされていることに気づきました。現在、本番環境と一致するように依存関係が固定された Docker コンテナを構築するため、検索エージェントと検証エージェントは

攻撃者が作成するものと同じアーティファクトです。
本番環境に十分忠実なサンドボックスを構築することが重要です。依存関係 (キューやデータストアなど) を除外すると、運用環境に存在する可能性のあるバグが過小報告される可能性があります。逆に、本番環境の防御 (WAF や認証ゲートウェイなど) を無視すると、モデルは本番環境ですでに軽減されている悪用できない結果を報告することになります。
それにもかかわらず、クラウドの依存関係、データ ストア、またはその他の現実世界の複雑さのために、代表的なサンドボックスの構築が現実的でない場合は、代わりに検出ステップ (以下) から始めてください。必ずしもサンドボックスで PoC を実行する必要はありません。フロンティア モデルは、ソース コードを分析するだけで脆弱性を発見するのが得意です。私たちのチームを含むいくつかのチームが、これが効果的であることを発見しました。トレードオフは検証フェーズにあり、実行中のターゲットがなければ PoC で結果を証明できないため、検証により多くの時間を割り当てます。検出結果の量が正当化されたら、後でサンドボックスに投資することもできます。
リファレンス サンドボックスについては、ハーネスの README.md を参照してください。この実装では、エージェントとターゲットは、出力がモデル API にロックされた gVisor 分離コンテナーで実行されます。ターゲットは、setup_sandb を使用して、特定のコミットに固定された Dockerfile からビルドされます。

[切り捨てられた]

## Original Extract

We share best practices for how you can work with Claude Opus to build a threat model, discover vulnerabilities in your codebase, then verify, triage, and patch them.

Using LLMs to secure source code | Claude
Meet Claude Products Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Using LLMs to secure source code
Using LLMs to secure source code
We share best practices for how you can work with Claude Opus to build a threat model, discover vulnerabilities in your codebase, then verify, triage, and patch them.
Share Copy link https://claude.com/blog/using-llms-to-secure-source-code
Model capabilities are advancing quickly, and unevenly. We’ve been working with security teams to find and fix vulnerabilities in their own code and open source software, and the work has given us a better understanding of how to use models to secure source code. Our primary takeaway: discovery is now straightforward to parallelize, and the bottleneck has shifted to verification, triage, and patching .
To give some indication of this discrepancy, as part of our own scanning of open source software, as of May 22, 2026, we had disclosed 1,596 vulnerabilities. To our knowledge, 97 of these have been patched.
This guide walks through how you can work with Claude Opus to build a threat model, discover vulnerabilities in your codebase, then verify, triage, and patch them. While we don’t have all the answers, we’ll share how teams have scaled discovery and what’s helped in the later stages. Get started today with the accompanying repo which includes skills for interactive workflows and a demo harness for autonomous scanning; we’ll call out the skill that implements each step as you read.
Teams finding and fixing the most vulnerabilities converged on a variation of existing best practices. We’ve distilled them into a sequence of six steps:
Threat model: Decide what counts as a vulnerability before you start scanning.
Sandbox: Build a sandbox environment to isolate agents and prove exploits.
Discovery: Have models look for vulnerabilities in your source code.
Verification: Independently confirm which findings are actually exploitable.
Triage: Deduplicate findings, assign severity, and prioritize what needs fixing.
Patching: Apply the fix, confirm the vulnerability is nullified, and search for variants.
The first two steps—building a threat model and a sandbox—are the setup for the rest of the loop. These are typically done once per codebase and revisited when the underlying system changes. The next four steps are the loop you’ll run against the source: discover, verify, triage, and patch.
The first run on a codebase typically has the highest number of findings. Subsequent runs tend to have fewer—though often more complex—vulnerabilities, as the simpler ones were patched in prior runs. However, don’t expect the n th run to have zero new findings. Models are stochastic, and a large codebase can have a long tail of vulnerabilities that continue to trickle in even when the code is unchanged.
On your first iteration with a codebase, you should run the loop multiple times, deciding when to stop based on the number of net-new findings and your risk tolerance for that system. After that first iteration, continue to scan (1) periodically or (2) whenever the code meaningfully changes.
Next, we’ll walk through each step in detail, explaining why it matters, what it produces, and how to implement it.
1. Threat model: Define what counts as a vulnerability
The most common cause of false positives is that the model lacks a good understanding of your trust boundaries. The model might flag code as vulnerable because it assumes a client could send corrupted values or an attacker could control the config, even though these inputs are trusted in your environment. Conversely, the model might assume that an internet-facing service is internal-only and thus under-report true vulnerabilities. In both cases, the model is wrong about the threat model, not the code.
One team noticed a pattern across their findings: the model performed best on systems with well-documented threat models, system design docs, requirements, and constraints. When the threat model was well-defined, the model's findings "were exploitable 90 percent of the time."
You can work with Claude to build a threat model in two steps:
First, bootstrap from the code, docs, and vulnerability history. Feed the model what you would hand a new security engineer on day one: architecture docs, wikis, entry points, git history, and past vulnerabilities. This helps overcome the challenge of inferring implicit knowledge, trade-offs, and design decisions from code alone. Then, ask the model to create a threat model that includes the system context, assets, entry points, and trust boundaries. Finally, have the model cluster past bugs and list the relevant vulnerability classes. Make sure the threat model documents what vulnerabilities you do and don’t care about, and why.
One team reviewed hundreds of past CVE and security-fix commits, distilled them into "bug-shape" hints, and asked the model two questions: was the fix complete, and was it applied everywhere else? They found three exploitable issues in an hour. As they put it: "'What have people exploited in the past' is sometimes a much easier cheat-code towards success than 'find me vulnerabilities in this codebase.'"
Second, have the model interview someone who knows the system well. Consider Shostack's four questions : What are we building? What can go wrong? What are we doing about it? Did we do a good job? Run the bootstrap step first so the interviewee isn’t starting from scratch. This way, instead of spending hours researching and building a threat model from scratch, they can start from a draft. And while the interview step is optional, it adds context the model can’t get from the code or docs, which improves the threat model.
A few practices can make a big difference:
Consider your dependencies’ security policies. Many open-source projects publish one. For example, vLLM’s security.md , SQLite's "Defense Against the Dark Arts" , and ImageMagick's security policy . Your threat model should consider them directly instead of rebuilding a policy from scratch.
Name what is trusted. If you trust config files or authenticated clients, document it in the threat model. These assumptions help separate non-exploitable bugs from actual exploits.
Include a THREAT_MODEL.md with the code. Have it in the repo and update it as code changes. The discovery agent can then read it before searching, skipping known non-issues.
You’ll use the threat model in two places. In discovery, as scope : partition the code, prioritize targets, and skip what is out of scope. This helps with large codebases you cannot scan entirely. In triage, as a filter: after scanning broadly, use the threat model to better calibrate severity to your system and environment.
One team scanning a large project had a 40% false positive rate and dug into why. The findings were reproducible and the PoCs proved exploitability. But the dev team who owned the code dismissed them as false positives because the bugs didn't fit the project's threat model. Another team's CISO put it succinctly: "[The model has] good context of the code, but not good context of us."
Try the threat-model skill . It walks through both steps described in this section— bootstrap derives a draft from your code, CVEs, and git history, and interview walks a system owner through Shostack’s four questions to refine it. The output is a THREAT_MODEL.md file which is used in the Discovery and Triage steps.
2. Sandbox: Run agents safely and verify exploitability
One purpose of the sandbox is to protect your systems. To enable models to run safely and autonomously, you need a strong isolation layer. Without it, the agent may overshoot the target and do something unexpected.
One team told the model it had no network access—when it actually did—and the model discovered it could fetch from GitHub anyway. Another team observed an agent answer a GitHub issue mid-scan. Neither action was malicious, but both demonstrated the need to enforce constraints via code and configuration.
Match the isolation to your threat model. Containers are fine for the discovery agent reading code, but run the target and its PoCs in a microVM (like Firecracker) or a full VM with egress locked down so nothing can reach your production systems. And never have credentials ( ~/.aws , ~/.ssh , .env ) available to the agent.
Give the sandbox network access only while you’re setting it up. Pull the dependencies, build, install tools, deploy the target, and run the existing tests to confirm everything works. Then, snapshot the environment and remove its network access. During scanning, allow traffic only to the model API, routed through a local proxy. Load the snapshot at the start of each run so every scan begins from the same clean slate.
Another purpose of the sandbox is to prove exploitability. During static scanning, the model reads code and hypothesizes what might break, but it cannot test if a path is reachable or if there's a compensating control. As a result, the model might flag unexploitable code-correctness bugs that you don’t actually care about. When teams built a sandbox where the agent could compile code, run tests, and detonate a proof of concept, non-exploitable findings dropped significantly.
One offensive-security team built a harness that gives the agent a test bed, with a simple verification rule: it’s only a true positive if the agent can build a proof of concept and run it on the test bed. Their assessment after six weeks was that "the biggest efficacy lever has been giving the model test beds, live systems, and running the PoCs."
When building sandboxes, pin as much as you can so every run uses the same code in the same environment: image tags, commit SHAs, dependencies, and build commands. Cache a local copy so the build requires no network, and aim for the container to be durable so multiple testing loops can just load it.
One team's scan flagged a vulnerability that turned out to be a byproduct of the agent downloading an older version of the library instead of what was actually deployed. This was caught by an engineer who read the transcript and spotted that a different dependency was being downloaded. They now build Docker containers with dependencies pinned to match production, so the finding agent and the verification agent operate on the same artifacts an attacker would.
It’s important to build sandboxes that are faithful enough to production. Excluding dependencies (like a queue or datastore) can lead to under-reporting bugs that may exist in production. Conversely, ignoring production defenses (like a WAF or auth gateway) leads to the model reporting unexploitable findings that your prod environment already mitigates.
Nonetheless, if building a representative sandbox is impractical because of cloud dependencies, data stores, or other real-world complexities, start with the discovery step (below) instead. You don’t necessarily need to run PoCs in a sandbox. Frontier models are good at finding vulnerabilities from just analyzing source code. Several teams, including our own, have found this effective. The trade-off is in the verification phase, where without a running target we can’t prove findings with a PoC, so budget more time for verification. You can also invest in the sandbox later, once the volume of findings justifies it.
Refer to the harness README.md for a reference sandbox. In this implementation, agents and targets run in gVisor-isolated containers with egress locked to the model API. The target is built from a Dockerfile pinned to a specific commit, with setup_sandb

[truncated]
