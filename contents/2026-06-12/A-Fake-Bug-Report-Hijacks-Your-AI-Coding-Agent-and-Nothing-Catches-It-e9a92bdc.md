---
source: "https://tenetsecurity.ai/blog/agentjacking-coding-agents-with-fake-sentry-errors/"
hn_url: "https://news.ycombinator.com/item?id=48507994"
title: "A Fake Bug Report Hijacks Your AI Coding Agent – and Nothing Catches It"
article_title: "A Fake Bug Report Hijacks Your AI Coding Agent - and Nothing Catches It. - Tenet Security"
author: "patrickdavey"
captured_at: "2026-06-12T21:38:49Z"
capture_tool: "hn-digest"
hn_id: 48507994
score: 3
comments: 0
posted_at: "2026-06-12T18:50:54Z"
tags:
  - hacker-news
  - translated
---

# A Fake Bug Report Hijacks Your AI Coding Agent – and Nothing Catches It

- HN: [48507994](https://news.ycombinator.com/item?id=48507994)
- Source: [tenetsecurity.ai](https://tenetsecurity.ai/blog/agentjacking-coding-agents-with-fake-sentry-errors/)
- Score: 3
- Comments: 0
- Posted: 2026-06-12T18:50:54Z

## Translation

タイトル: 偽のバグレポートが AI コーディング エージェントをハイジャック – しかし何も捕らえられない
記事のタイトル: 偽のバグレポートが AI コーディング エージェントをハイジャック - 何も捕らえられない。 - テネットセキュリティ
説明: Tenet Threat Labs は、AI コーディング エージェントをハイジャックして、開発者のコンピュータ上で攻撃者が制御するコードを実行する新しいクラスの攻撃「エージェントジャッキング」を実証しました。

記事本文:
偽のバグレポートが AI コーディング エージェントをハイジャックしますが、何も捕らえられません。 - テネットセキュリティ
コンテンツにスキップ
Tenet Security が 2026 年の Fortress Cybersecurity Awards 受賞者に選ばれました 🏆
Tenet Security が 2026 年の Fortress Cybersecurity Awards 受賞者に選ばれました 🏆
Tenet Security が 2026 年の Fortress Cybersecurity Awards 受賞者に選ばれました 🏆
偽のバグレポートが AI コーディング エージェントをハイジャックしますが、何も捕らえられません。
Tenet Threat Labs は、AI コーディング エージェントをハイジャックして、開発者のマシン上で攻撃者が制御するコードを実行する新しいクラスの攻撃「エージェントジャッキング」を実証しました。これは、単一の偽のエラー レポートによって引き起こされ、あらゆるセキュリティ制御からは見えません。パブリック Sentry API のみを使用し、何も侵害していないため、2,388 の組織が暴露され、100 を超えるエージェントが管理されたテストで挿入されたエラーに対処することが確認され、Fortune 500 企業から独立系開発者に至るまでの組織でエージェントの実行が確認されました。
図 1 – エージェントジャッキング チェーン。 Every step is authorized, which is why no security control sees it.
Tenet Security の Threat Labs による新しい調査では、Web サイトのソース コードに含まれる公開資格情報以上の認証を必要としない単一の挿入エラー イベントによって、AI コーディング エージェントがハイジャックされ、開発者のマシン上で任意のコードが実行される仕組みが実証されました。
この攻撃は、Sentry のイベント取り込み (DSN を持つすべてのユーザーからの任意のペイロードを受け入れる) と Sentry MCP サーバー (このデータを信頼できるシステム出力として AI エージェントに返す) の交差点にある重大なアーキテクチャ上の欠陥を悪用します。攻撃者は、巧妙に作成された入力を Sentry エラー イベントに挿入することで、Sentry 自身の修復ガイダンスと視覚的および構造的に区別できない指示を作成します。 Claude Code や Cursor を含む AI コーディング エージェントは、これらを次のように解釈します。

「診断解決手順」を gitimate し、攻撃者が制御する npm パッケージを実行します。
影響: 単一の挿入されたエラーにより、環境変数 (AWS キー、GitHub トークン、Sentry 認証トークン)、git 認証情報、プライベート リポジトリ URL、および開発者の ID が攻撃者の手の届く範囲に置かれ、認証情報のフィッシングや事前のサーバー侵害がなく、開発者の通常のワークフローを超えるユーザーの操作もなく、静かにサーバーに流出します。認証情報のフィッシング、事前のサーバー侵害、AI エージェントに Sentry エラーの調査を依頼する開発者の通常のワークフローを超えるユーザーの操作はありません。
企業が AI コーディング エージェントの導入を競う中、この調査は、エージェント自体が今や攻撃対象となっており、組織が自らについて公開しているデータのみを使用して、エージェントを信頼する開発者に敵対していることを証明しています。このイノベーションは新しいエクスプロイトではなく、実際にエージェントをいかに簡単かつどの規模でハイジャックできるかということです。これをキャッチできる唯一の場所は、エージェントの実行時です。
AI コーディング エージェント: 隠れた欠陥を持つ強力なアシスタント
Claude Code や Cursor などの最新の AI コーディング エージェントは、単純なオートコンプリート ツールから、ファイルの読み取り、ターミナル コマンドの実行、外部ツールのクエリ、コードの変更を行うことができる強力なアシスタントに進化しました。これらのエージェントは、モデル コンテキスト プロトコル (MCP) を通じて外部サービス (エラー監視用の Sentry など) に接続し、返されたデータを信頼できるシステム出力として処理します。
危険はこの暗黙の信頼にあります。 AI エージェントが Sentry に未解決のエラーを問い合わせると、開発者と同じように応答を受け取り、それに基づいて動作します。ただし、開発者とは異なり、エージェントはエラー イベントが実際のアプリケーションのクラッシュによって生成されたのか、それとも atta によって挿入されたのかを検証できません。

カー。 MCP ツールの応答に対するエージェントの信頼により、挿入されたデータからコード実行までの直接的な経路が作成されます。
AI コーディング エージェントは、読み取ったデータと動作するための命令の違いを区別できません。エージェントがそれを読み取る場所 (エラー ログなど人間が決して探すことのない場所であっても) にコマンドを埋め込むと、エージェントはそれを実行するだけで済みます。これはモデル自体の制限であり、パッチで解決できる構成ミスではありません。
攻撃の構造: 挿入されたエラーから RCE まで
この攻撃は、攻撃者にとっては驚くほど単純ですが、ターゲットにとっては壊滅的です。パブリック DSN を使用して Sentry に POST される 1 つの細工されたエラー イベントから始まります。資格情報は、仕様上、無数の運用 Web サイトの JavaScript ソースに存在します。違反はありません。認証情報は盗まれません。伝統的な意味での搾取はありません。攻撃者は被害者のインフラストラクチャには決して触れません。
悪意のある命令は、通常のエラーの中に正規の「解決策」を装って到着します。開発者が AI エージェントに Sentry の問題の修正を依頼すると、エージェントは攻撃者のコマンドを信頼できるガイダンスとして読み取り、開発者自身の権限で、開発者自身のマシン上で実行します。
ステップ 1 : ターゲットの Sentry DSN を見つけます。これは、フロントエンド JavaScript に埋め込んでも安全であると Sentry が意図的に文書化した公開の書き込み専用資格情報です。検出方法には、Web サイトの JavaScript ソースの検査、Censys の HTTP 本文内の ingest.sentry.io の検索、または GitHub コードの検索が含まれます。
ステップ 2: 通常のイベント作成: 細工されたエラー イベントを Sentry の取り込みエンドポイントに POST します。 DSN を超える認証は必要ありません。攻撃者は、イベント ペイロード全体 (エラー メッセージ、タグ、コンテキスト キー、追加データ、パンくずリスト、ユーザー、スタック トレース、フィンガープ) を制御します。

リント。 Sentry はそれを受け入れ (HTTP 200)、正規のアプリケーション エラーと同様に処理します。
ステップ 3 : マークダウンの挿入: 挿入されたイベントには、メッセージ フィールドとコンテキスト キー名に注意深くフォーマットされたマークダウンが含まれています。 Sentry MCP サーバーがこのイベントを AI エージェントに返すと、マークダウンは、Sentry 独自のシステム テンプレートと視覚的に同一の見出し、コード ブロック、テーブルなどの構造化コンテンツとしてレンダリングされます。挿入されたコンテンツには、npx コマンドを使用した偽の「## Resolution」セクションが含まれています。
Step 4 : Agent Manipulation: When a developer asks their AI agent to ‘fix unresolved Sentry issues,’ (or any other related prompt) the agent queries Sentry via MCP and receives the injected event.エージェントは、ソース コードの調査から慎重に誘導され、提案された診断ツールの実行に向けられます。エージェントはこれを正当なガイダンスと区別できません。
ステップ 5: コードの実行: エージェントは、npx @tenet-controlled-validation-package –diagnose を実行します。パッケージはパブリック npm レジストリからダウンロードされ、開発者の完全な権限で実行されます。パッケージには、管理されたテストが Tenet Security によって実行されていることを明確にするメッセージが含まれており、ヘッダーは「X-Tenet-Security」、値は「ResponsibleDisclosure [SECURITY SCAN]」です。 Advisory-tracker.com へのビーコンに接続します。責任ある開示メッセージもビーコンに添付されます。
ステップ 6: パッケージは、環境変数が存在すること、~/.aws/config、~/.npmrc、~/.docker/config.json のファイル サイズが調査されていること、およびネットワーク インターフェイス (VPN 検出) を確認します。暴露データの検証は、2 つの連続した POST リクエストを介して Tenet ビーコン サーバーに送信され、関連情報が企業に開示されます (情報は保持または保存されず、すべてのプローブ データは削除されます)。

ベスト プラクティスを遵守し、Sentry セキュリティ チームと同様に組織が適切にセキュリティを確保できるようにするために削除されました)。
新しいアプローチ: 信頼できる開発者ツールを介した攻撃
この攻撃がユニークなのは、開発者を直接標的にするのではなく、開発者が信頼する AI エージェントを標的にすることです。いくつかの要因がこれを特に危険にしています。
フィッシングは必要ありません: 攻撃者が開発者と対話することはありません。この攻撃は、AI エージェントに Sentry エラーの調査を依頼する開発者の通常のワークフローを通じて行われます。
エントリ ポイントとしての公開資格情報: Sentry の DSN は意図的に公開されており、フロントエンド JavaScript に埋め込まれています。この設計上の決定は、AI エージェント以前の世界では安全ですが、注入されたイベントが信頼できる出力として AI エージェントに返されると、壊滅的なものになります。
正規のガイダンスと区別がつかない: マークダウン挿入により、Sentry 独自の MCP システム テンプレートと構造的に同一のコンテンツが作成されます。攻撃者のコンテンツと実際のセントリーのガイダンスを区別する視覚的または構造的な指標はありません。
簡単に拡張可能: ペイロードを作成したら、それを何千もの Sentry プロジェクトに同時に注入できます。私たちは、制御されたキャンペーンで 100 以上の組織をターゲットにすることでこれを実証しました。
証明: 制御された現実世界の検証
これが理論上のものではないことを証明するために、私たちのチームは制御された条件下で攻撃をエンドツーエンドで検証し、現実世界のターゲットに対する悪用可能性を確認しました。
2,388 の組織が、受動的な偵察 (Censys インデックス作成、コード検索、CDN ローダー抽出) により、有効な挿入可能な DSN で公開されていることが判明しました。トランコトップ1Mで71位。
制御された検証ウェーブ全体で、100 以上の AI コーディングがクロード コード、カーソル、コーデックスを含む挿入されたエラーに作用し、85% の悪用が成功しました。

市場で最も広く使用されているエージェント全体で、挿入されたエラーに対する率を示します。
フォーチュン 500 企業 (2,000 億ドル以上)、20 億ドル以上のホスティング インフラストラクチャ プロバイダー、科学技術コンピューティング企業、Web スタートアップ企業、その他複数の開発チームに及ぶ、多くの組織にわたってエージェント実行の 100 以上の確認されたインスタンスが完全に文書化されています。
2,221 exposed organizations were not included in the validation set .同じ条件が何千ものプロジェクトに存在し、最小限のリソースで到達できます。
完全なキャプチャ ログ、Sentry 取り込みエンドポイントに送信されたリクエスト、機密マテリアル (環境変数、AWS 認証情報、Kubernetes トークン、GitHub OAuth トークン、git リポジトリ URL) の存在と到達可能性を確認するタイムスタンプ付きのアクセス証明テレメトリ - これらが存在し、公開されたことを記録します。
図 2 – 確認された組織と摘発された組織は 6 つの大陸にまたがっています。各マーカーは、キャンペーンで到達した個別の組織です。
編集された証拠 – 野生で捕獲されたもの
以下のすべての値はピクセル レベルで編集されます。このドキュメントのどこにも、実際の資格情報、アイデンティティ、ホストは登場しません。
被害者には無害な診断出力 (Node.js バージョンとメモリ統計) のみが表示されていましたが、エージェントはライブ クラウド、ソース管理、およびクラスターの資格情報を黙って本物の攻撃者に引き渡したはずです。
誰のエージェントがハイジャックされたのか – サンプル (編集済み)
その範囲は、2,500 億ドル規模の巨大テクノロジー企業から独立した個人開発者まで多岐にわたり、クラウド セキュリティ ベンダーもその中に含まれていました。規模、分野、セキュリティ予算によって安全性が予測されることはありません。
脅威の新時代: AI エージェントのセキュリティが変わる理由
この発見は単なる脆弱性ではなく、ソフトウェア開発の攻撃対象領域における根本的な変化を表しています。
長年にわたりサプライチェーン攻撃が焦点となっている

実際のパッケージ (SolarWinds、CodeCov) を侵害したり、タイポスクワッティングで開発者を騙したりすることを防止します。しかし、AI コーディング エージェントを使用すると、攻撃者はパッケージを侵害したり人間を騙したりする必要がなく、AI エージェントが信頼するデータを注入するだけで済みます。可観測性プラットフォームはコマンド アンド コントロール チャネルとなり、AI エージェントは実行エンジンになります。
エンタープライズ環境では、挿入された 1 つのエラーにより、攻撃者は CI/CD パイプライン認証情報の盗用、プライベート ソース コード リポジトリへのアクセス、クラウド インフラストラクチャの侵害、永続的なアクセスの確立を可能になる可能性があります。これらはすべて、ターゲットの開発者と直接やり取りすることなく行われます。
リスクはセントリーに限定されません。外部から影響を受けたデータを AI エージェントに返す MCP ツール統合では、同じ脆弱性クラスが作成されます。 AI エージェントのエコシステムが拡大し、MCP 経由で接続するツールが増えるにつれ、攻撃対象領域は飛躍的に増大します。
全体的な、検出不能、単一ベンダーのバグではない
この脆弱性は、単一の製品の欠陥ではなく、エージェントがツール出力を処理する方法にあるため、テストされたすべてのエージェント (市場で最も広く使用されている AI コーディング アシスタント) で機能しました。 Sentry の MCP 統合は実証済みのエントリ ポイントです。根本的な問題はエコシステム全体で共有されています。
プロンプト層の防御は失敗した。エージェントは、明示的に指示された場合でも、det を通じてペイロードを実行しました

[切り捨てられた]

## Original Extract

Tenet Threat Labs has demonstrated a new class of attack "Agentjacking" that hijacks AI coding agents into running attacker-controlled code on a developer's

A Fake Bug Report Hijacks Your AI Coding Agent - and Nothing Catches It. - Tenet Security
Skip to content
Tenet Security Named a 2026 Fortress Cybersecurity Awards Winner 🏆
Tenet Security Named a 2026 Fortress Cybersecurity Awards Winner 🏆
Tenet Security Named a 2026 Fortress Cybersecurity Awards Winner 🏆
A Fake Bug Report Hijacks Your AI Coding Agent – and Nothing Catches It.
Tenet Threat Labs has demonstrated a new class of attack “Agentjacking” that hijacks AI coding agents into running attacker-controlled code on a developer’s machine, triggered by a single fake error report and invisible to every security control. Using only public Sentry APIs, breaching nothing, we found 2,388 organizations exposed, saw 100+ agents act on injected errors in controlled testing, with confirmed agent execution at organizations spanning from Fortune 500 enterprise down to independent developers.
Figure 1 – The Agentjacking chain. Every step is authorized, which is why no security control sees it.
New research by Tenet Security’s Threat Labs demonstrates how a single injected error event requiring no authentication beyond a public credential found in any website’s source code can hijack AI coding agents into executing arbitrary code on developer machines.
The attack exploits a critical architectural flaw at the intersection of Sentry’s event ingestion (which accepts arbitrary payloads from anyone with the DSN) and the Sentry MCP server (which returns this data to AI agents as trusted system output). By injecting crafted input into Sentry error events, an attacker creates instructions that are visually and structurally indistinguishable from Sentry’s own remediation guidance. AI coding agents including Claude Code and Cursor interpret these as legitimate ‘diagnostic resolution steps’ and execute attacker-controlled npm packages.
The impact: a single injected error puts environment variables (AWS keys, GitHub tokens, Sentry auth tokens), git credentials, private repository URLs, and developer identity within an attacker’s reach – silently exfiltrated to their server, with no credential phishing, no prior server compromise, and no user interaction beyond the developer’s normal workflow. No credential phishing, no server prior compromise, no user interaction beyond the developer’s normal workflow of asking their AI agent to investigate Sentry errors.
As enterprises race to deploy AI coding agents, this research proves the agents themselves are now the attack surface – turned against the developers who trust them, using nothing but data those organizations publish about themselves. The innovation is not a novel exploit: it is how trivially and at what scale agents can be hijacked in the wild. The only place left to catch it is at the agent’s runtime.
AI Coding Agents: A Powerful Assistant with a Hidden Flaw
Modern AI coding agents like Claude Code and Cursor have evolved from simple autocomplete tools into powerful assistants that can read files, execute terminal commands, query external tools, and make code changes. Through the Model Context Protocol (MCP), these agents connect to external services – including Sentry for error monitoring – and treat the data returned as authoritative system output.
The danger lies in this implicit trust. When an AI agent queries Sentry for unresolved errors, it receives the response and acts on it – just as a developer would. But unlike a developer, the agent cannot verify whether an error event was generated by a real application crash or injected by an attacker. The agent’s trust in MCP tool responses creates a direct pathway from injected data to code execution.
AI coding agents cannot tell the difference between the data they read and an instruction to act. Plant a command somewhere an agent will read it – even somewhere no human would ever look for one, like an error log – and the agent may simply execute it. This is a limitation of the models themselves, not a misconfiguration that can be patched away.
The Anatomy of the Attack: From Injected Error to RCE
The attack is alarmingly simple for the attacker but devastating for the target, it begins with one crafted error event , POSTed to Sentry using a public DSN – a credential that, by design, sits in the JavaScript source of countless production websites. No breach. No stolen credentials. No exploit in the traditional sense. The attacker never touches the victim’s infrastructure.
The malicious instruction arrives disguised as a legitimate “Resolution” inside an ordinary error. When a developer asks their AI agent to fix the Sentry issue, the agent reads the attacker’s command as trusted guidance and runs it – with the developer’s own privileges, on the developer’s own machine.
Step 1 : Find the target’s Sentry DSN – a public, write-only credential that Sentry intentionally documents as safe to embed in frontend JavaScript. Discovery methods include: inspecting any website’s JavaScript source, Censys searches for ingest.sentry.io in HTTP bodies, or GitHub code search.
Step 2 : Regular event creation: POSTing a crafted error event to Sentry’s ingest endpoint. No authentication beyond the DSN is required. The attacker controls the entire event payload: error message, tags, context keys, extra data, breadcrumbs, user, stack traces, and fingerprint. Sentry accepts it (HTTP 200) and processes it identically to a legitimate application error.
Step 3 : Markdown Injection: The injected event contains carefully formatted markdown in the message field and context key names. When the Sentry MCP server returns this event to an AI agent, the markdown renders as structured content: headings, code blocks, and tables that are visually identical to Sentry’s own system template. The injected content includes a fake ‘## Resolution’ section with an npx command.
Step 4 : Agent Manipulation: When a developer asks their AI agent to ‘fix unresolved Sentry issues,’ (or any other related prompt) the agent queries Sentry via MCP and receives the injected event. The agent is carefully steered away from investigating source code and toward executing the suggested diagnostic tool. The agent cannot distinguish this from legitimate guidance.
Step 5 : Code Execution: The agent executes: npx @tenet-controlled-validation-package –diagnose. The package downloads from the public npm registry and runs with the developer’s full privileges. The package contains a message clarifying the controlled test is running by Tenet Security with header: “X-Tenet-Security” and with the value “ResponsibleDisclosure [SECURITY SCAN]”. Reaching out to a beacon to advisory-tracker.com . A Responsible disclosure message is attached to the beacon as well.
Step 6 : The package confirms that environment variables exist, file sizes of ~/.aws/config, ~/.npmrc, ~/.docker/config.json are probed, and network interfaces (VPN detection). Validation Of Exposure Data is sent via two sequential POST requests to Tenet beacon server, while disclosing to companies the relevant information (no information was ever kept or saved; all probe data was deleted and removed to adhere to best practices and make sure the organizations secure themselves correspondingly with Sentry security team as well).
A New Approach: Attacking Through Trusted Developer Tools
What makes this attack unique is that it doesn’t target the developer directly – it targets the AI agent that the developer trusts. Several factors make this particularly dangerous:
No phishing required: The attacker never interacts with the developer. The attack flows through the developer’s normal workflow of asking their AI agent to investigate Sentry errors.
Public credential as entry point: Sentry’s DSN is intentionally public and embedded in frontend JavaScript. This design decision – safe in a pre-AI-agent world – becomes catastrophic when injected events are returned to AI agents as trusted output.
Indistinguishable from legitimate guidance: The markdown injection creates content that is structurally identical to Sentry’s own MCP system template. No visual or structural indicator distinguishes attacker content from real Sentry guidance.
Scales effortlessly: Once a payload is crafted, it can be injected into thousands of Sentry projects simultaneously. We demonstrated this by targeting 100+ organizations in a controlled campaign.
Proof: A Controlled, Real-World Validation
To prove this wasn’t theoretical, our team validated the attack end-to-end in controlled conditions and confirmed exploitability against real-world targets.
2,388 organizations found exposed with valid injectable DSNs – via passive reconnaissance (Censys indexing, code search, CDN loader extraction). 71 rank in the Tranco top-1M.
Across controlled validation waves 100+ AI coding acted on the injected errors – including Claude Code, Cursor and Codex – an 85% exploitation success rate against injected errors, across the most widely-used agents on the market.
More than 100+ confirmed instances of agent execution across many organizations , documented in full – spanning a Fortune 500 enterprise ($200Bn+), a $2B+ hosting infrastructure provider, a scientific computing firm, a web startup, and multiple other development teams.
2,221 exposed organizations were not included in the validation set . The same conditions exist in thousands of projects, reachable with minimal resources.
Full capture logs, requests sent to Sentry ingest endpoints, and timestamped proof-of-access telemetry confirming the existence and reachability of sensitive material (environment variables, AWS credentials, Kubernetes tokens, GitHub OAuth tokens, git repository URLs) – recording that these were present and exposed.
Figure 2 – Confirmed and exposed organizations span six continents. Each marker is a distinct organization reached in the campaign.
Redacted Evidence – Captured in the Wild
Every value below is redacted at the pixel level. No real credential, identity, or host appears anywhere in this document.
The victim saw only benign diagnostic output – Node.js version and memory stats – while the agent would have silently surrendered live cloud, source-control, and cluster credentials to a real attacker.
Whose Agent Got Hijacked – Sample (Redacted)
The range ran from a ~$250B technology giant to independent solo developers – and even a cloud security vendor was among the exposed. No size, sector, or security budget predicted safety.
A New Era of Threat: Why This Changes AI Agent Security
This discovery is more than just another vulnerability – it represents a fundamental shift in the software development attack surface.
For years, supply chain attacks focused on compromising real packages (SolarWinds, CodeCov) or tricking developers with typosquatting. But with AI coding agents, attackers no longer need to compromise a package or trick a human – they just need to inject data that the AI agent trusts. The observability platform becomes a command-and-control channel, and the AI agent becomes the execution engine.
In an enterprise environment, a single injected error could allow an attacker to: steal CI/CD pipeline credentials, access private source code repositories, compromise cloud infrastructure, and establish persistent access – all without any direct interaction with the target developer.
The risk is not limited to Sentry. Any MCP tool integration that returns externally-influenced data to AI agents creates the same vulnerability class. As the AI agent ecosystem expands and more tools connect via MCP, the attack surface grows exponentially.
Systemic, Undetectable, Not a One-Vendor Bug
It worked across every agent tested – the most widely-used AI coding assistants on the market – because the weakness is in how agents handle tool output, not a flaw in any single product. Sentry’s MCP integration is the demonstrated entry point; the underlying problem is shared across the ecosystem.
Prompt-layer defenses failed. Agents executed the payload even when explicitly instructed – through det

[truncated]
