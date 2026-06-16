---
source: "https://tomevault.io/standards/state-reports/2026"
hn_url: "https://news.ycombinator.com/item?id=48552977"
title: "AI Instruction Use 78.7% of repos configure 1 tool but devs use avg 2-4"
article_title: "State of AI instructions, 2026 · Standards | TomeVault"
author: "olijboyd"
captured_at: "2026-06-16T10:41:01Z"
capture_tool: "hn-digest"
hn_id: 48552977
score: 2
comments: 0
posted_at: "2026-06-16T10:07:21Z"
tags:
  - hacker-news
  - translated
---

# AI Instruction Use 78.7% of repos configure 1 tool but devs use avg 2-4

- HN: [48552977](https://news.ycombinator.com/item?id=48552977)
- Source: [tomevault.io](https://tomevault.io/standards/state-reports/2026)
- Score: 2
- Comments: 0
- Posted: 2026-06-16T10:07:21Z

## Translation

タイトル: AI 命令を使用するリポジトリの 78.7% が 1 つのツールを構成していますが、開発者は平均 2 ～ 4 を使用しています
記事のタイトル: AI 命令の現状、2026 · 標準 |書物保管庫
説明: AI 命令ファイルに関する年次報告書: この形式がガバナンス、マーケットプレイス、スキャナー、そして最初のポイズニング キャンペーンを獲得した年と、その下にあるデータ。 267,228 のリポジトリが測定されました。マルチツールのギャップ、ライセンス、安全性、署名率はゼロでした。

記事本文:
AI 命令の現状、2026 · 標準 | TomeVault TomeVault Home Vault Scan Search Docs Standards Wtf is… サインイン TomeVault · 標準年次報告書 · 2026 年次版 · 2026
これは、AI 命令ファイルがインフラストラクチャになり、財団によって管理され、マーケットプレイスによって配布され、スキャンされ、ベンダー カタログに署名され、初めて大規模に汚染された年でした。このレポートは、267,228 の公共リポジトリから抽出された、その年の記録とその下の測定値です。
のリポジトリは 1 つの AI コーディング ツールのみを構成しますが、その背後にいる開発者は複数の AI コーディング ツールを使用しています。
267,228 のリポジトリ · すべての主要な形式 · 2026 年 6 月 4 日と 2026 年 6 月 11 日に凍結されたスナップショット · 再現可能な数
03 ディストリビューション層の信頼
01 AI コーディング ツールを使用するリポジトリの 78.7% は、そのうちの 1 つだけを構成しています。
開発者は 1 つにとどまりません。ほとんどの場合、2 ～ 4 つの AI ツールが実行されるため、他のアシスタントは何もせずにそのコードベースを作業することになります。
02 AGENTS.md はツールの橋渡しとして構築されました。これまたは CLAUDE.md を使用するリポジトリのうち、両方を使用しているのは 18.3% のみです。
1 つのクロスベンダー標準は、橋渡しではなく、代替として採用されています。スプリットを移動しますが、閉じるわけではありません。
03 CLAUDE.md と AGENTS.md の両方を配布するリポジトリ 10 個のうち 9 個はそれらに接続しないため、2 番目のファイルは読み取られません。
チームは2つの契約を維持し、そのうちの1つについて報道を受け、その後、渡されていないルールを無視したとしてモデルを非難した。
04 CLAUDE.md ファイルの 26.3% は、Anthropic が独自のドキュメントで公開している 200 行のガイドラインを超えて実行されます。
その長さを超えると、ファイルは単にコンテキストを無駄にするだけではありません。 2026 年の調査では、エージェントが仕事を完了する割合が低下することがわかりました。
05 55.8% の命令ファイル リポジトリでは、ファイルを再配布するための使用可能な権限が付与されていません。
これらのファイルをインスタに提供した年

すべてのコマンド、ディレクトリ、およびマーケットプレイス。彼らが配布するもののほとんどは、誰も配布の許可を与えていません。
06 オープン コーパスでの署名の採用はゼロです。調査された 986 のリポジトリのうち、指示ファイルに署名したものはありません。
業界が今年構築した信頼層は、ディストリビューターが必要とする場所に存在しており、これまでのところ他の場所には存在しません。信頼はファイルではなくチャネルの財産になりつつあります。
07 スキャンされた 219,024 個の指示ファイルの 5.0% には、少なくとも 1 つの決定的な安全性に関する知見が含まれています。
ClawHavocは大音量バージョンでした。これは、エージェントがデフォルトでスキャンをまったく行わずにロードするファイルの静かな基本レートです。
結論。今年は、管理された標準、インストール コマンド、信頼スタック、脅威モデルなど、命令ファイルをめぐる世界が成長しました。ファイル自体が追いついていません。これらは、マルチツールの世界で 1 つのツールを提供し、取得するディストリビューションにはライセンスがありません。また、それらのために構築された信頼層は、オープン コーパスでは誰も使用されません。
AI コーディング ツールを 1 つだけ構成する
ファイルを再配布する権利を与えません
986 リポジトリ プローブ内のリポジトリは、指示ファイルに署名します
12 か月前には、指示ファイルは慣例でした。今年、同社は基盤、配布レイヤー、脅威モデル、そして最初のポイズニングキャンペーンを取得しました。
指示ファイルは、開発者と 1 つのツールの間のプライベートな約束事ではなくなりました。 9 か月にわたって、中立的なガバナンス、パッケージ マネージャー スタイルの配布レイヤー、ベンダー カタログ、セキュリティ分類法、およびインシデント記録を取り上げました。以下の日付はその物語の背表紙です。パート I の残りの部分では、最も重要な 3 つについて説明します。
2025 年 12 月 9 日 Linux Foundation は、Anthropic の Model Context Protocol、Block の goose、および OpenAI の AGENTS.md を創設プロジェクトとして、Agentic AI Foundation を設立しました。インス

truction-file 標準に中立的なホームが追加されました。
2025 年 12 月 Anthropic は、エージェント スキルをオープン スタンダード (agentskills.io) としてリリースします。このフォルダーには、SKILL.md を含むフォルダーがあり、これを読み込むエージェントへの指示がパッケージ化されています。
2026 年 1 月 15 日 新しいフォーマットに関する最初の大規模監査が arXiv に投稿されました。分析された 31,132 のスキルのうち、26.1% に少なくとも 1 つの脆弱性パターンが含まれていました。
2026 年 1 月 20 日 Vercel は、スキル CLI とスキル パッケージのディレクトリおよびリーダーボードである skill.sh を出荷しました。インストールコマンドを取得する手順です。
2026 年 2 月 1 日、Koi Security は、OpenClaw エージェントにサービスを提供するマーケットプレイスである ClawHub で ClawHavoc: 341 の悪意のあるスキルを公開しました。このフォーマットの最初に文書化されたポイズニング キャンペーン。
2026 年 2 月 17 日 skill.sh に 3 つのベンダー (Gen Digital、Socket、Snyk) による自動セキュリティ監査が追加されました。悪意があるとフラグが立てられたスキルは、検索やリーダーボードから非表示になります。
2026 年 2 月 チューリッヒ工科大学は、コンテキスト ファイルに関する最初の対照研究を発表しました。ファイルが過剰に指定されていると、エージェントの成功率が低下し、コストが 20% 以上増加します。無駄のないビートは長く続きます。
2026 年 3 月 13 ～ 26 日 コミュニティ RFC は、エージェント スキル仕様に暗号署名を吸収するよう求めています。仕様管理者はこれを拒否し、代わりにディストリビューション層を指摘しました。
2026 年 4 月 16 日 GitHub は、GH スキルをパブリック プレビューで出荷します。インストール、検索、公開、バージョンの固定とコンテンツ アドレス指定の変更検出を備え、意図的に一元化された検証は行われていません。
2026 年 4 月 22 日 Google は、Cloud Next の 1 日目に発表された、オープン エージェント スキル フォーマットの公式スキル リポジトリを開始しました。
2026 年 5 月 19 日 NVIDIA は、オープン仕様での検証済みエージェント スキル: スキャン、スキル カード、分離署名、インストール時の検証を発表しました。同日、Google のエンタープライズ スキル レジストリがパブリック プレビューに入ります。
2026 年 6 月 3 日 Libraries.io を構築し、Ecosyste.ms を運営している Andrew Nesbitt 氏は、そのギャップを明確に公開しています: Trust Machines PAC

kage レジストリの構築に 10 年を費やしたものは、ほとんどがスキルのために存在するものではありません。
最も重要な 2 つのフォーマットは、中立的なガバナンスまたはオープンな仕様の下で年末を迎え、最大手のベンダーがそれらのフォーマットにパブリッシュしました。
AGENTS.md は、OpenAI リポジトリ コンベンションとして最初の 1 年を費やしました。 2025 年 12 月 9 日、これは、Model Context Protocol および goose と並んで、Linux Foundation 傘下の有向基金である Agentic AI Foundation の設立プロジェクトとなりました。 2月下旬までに、財団の会員組織は146を数えた。これが決着したと主張する人への注意点が 1 つあります。その日の時点で、財団自体は、その 3 つのプロジェクトの正式なガバナンス モデルがまだ定義中であると述べています。そのため、AGENTS.md は中立的な本拠地とメンバーシップを持っていますが、まだ完成した憲法ではありません。
エージェント スキル フォーマットは、同じアークをより速く実行しました。 Anthropic によって 12 月にオープン スタンダードとしてリリースされましたが、Google が公式スキルを出荷する形式、GitHub の gh スキルが検証する形式、および NVIDIA が検証済みカタログに選択した形式として春に終了しました。 Google の Gemini CLI ドキュメントには、スキルは「エージェント スキルのオープン スタンダードに基づいている」と直接記載されています。最大手のベンダーが、競合他社が独自に作成したフォーマットをフォークするのではなく、そのフォーマットで公開すると、そのフォーマットは誰の製品でもなくなります。
記録に値する唯一の非対称性は Google 内にあります。そのパブリック リポジトリは公然と標準に準拠しています。同社のエンタープライズ スキル レジストリは、5 月 19 日からパブリック プレビュー中ですが、SKILL.md パッケージ構造を必要としますが、オープン仕様の名前はどこにも記載されておらず、安全なプライベート リポジトリであると説明されています。フォーマットは伝わります。カタログが商品です。この分割、下にオープンフォーマット、上に独自のディストリビューションが、今年のベンダーの動きのほとんどのシェアを占めています。
なぜマットなのか

えーっ。 2026 年に指示ファイルを標準化するチームは、ベンダーの気まぐれに賭けることはもうありません。フォーマットは管理されているものとオープンなものがあり、今年の参加者はフォーマットを中心とするのではなく、フォーマットに基づいて構築されています。未解決の質問は、ファイルがどのようなものであるかではなく、誰がそれを保証するかという階層に移りました。それが次のセクションです。
信頼は仕様ではなくディストリビューション層に到達しました
この年を決定づけた議論は、信頼がどこに生きるべきかということでした。仕様には「ここにはありません」と書かれていました。
3 月、コミュニティ RFC は、エージェント スキル仕様自体に暗号署名を組み込むことを提案しました。これは、既知の URL で発見可能なパブリッシャー キーを使用した、SKILL.md フロントマターの ed25519 署名ブロックです。仕様管理者はそれを吸収することを拒否しました。 3 月 16 日に述べられた彼の立場は、署名は「配布層での実装がより適切である」というものでした。 3月25日、同氏は「これを仕様の正式な一部とする可能性は低い」と付け加えた。議論はまだオープンなままであるため、これは非公開の決定ではなく、明言された管理者の立場ですが、仕様は維持されています: 6 月 11 日の時点で、その前付は署名、検証、または信頼フィールドをまったく定義していません。コンパニオン実装では、ツール側から同じストーリーを伝えます。スキル検証をディストリビューション CLI に追加するプル リクエストは 3 月 15 日にオープンされ、マージされていません。 CLI のリリース バージョンには verify コマンドはありません。
そして 5 月に、NVIDIA はメンテナが指摘したものを正確に出荷しました。同社の Verified Agent Skills プログラムは、完全にディストリビューション層でオープン仕様に構築されたトラスト スタックです。同社によれば、オープンソース スキャナは従来のリスクとエージェント固有のリスクの両方をチェックし、スキルごとに機械読み取り可能なスキル カード、OpenSSF モデル署名形式の分離署名、インストール時または CI での検証を行うものです。これまで

カタログ内の 207 スキルの 1 つは、SKILL.md、スキル カード、署名のフルセットを同梱しています。そして、カタログは、プレーンな npx skill add nvidia/skills とともに、競合他社のレジストリである skill.sh を通じて配布されます。
コミュニティは仕様に信頼を求めました。仕様はディストリビューターに向けられていました。 NVIDIA は 2 か月後にそこにそれを構築しました。
信頼層の残りの部分は、同じ点の周りに集まっています。 Skills.sh は 2 月に 3 ベンダーの自動監査を追加し、リストされたすべてのスキルを評価し、悪意があるとフラグが付けられたスキルを非表示にしました。 OWASP は春に Agentic Skills Top 10 プロジェクトを開始しました。これは 10 のリスク カテゴリの草案を含むインキュベーター段階の提案であり、独自のロードマップに基づいて 2026 年の第 4 四半期を目標とした 1.0 リリースが予定されています。そして学術的な記録は1月に到着した。実際に使われているスキルに関する最初の大規模な監査では、分析した31,132のうち26.1%が少なくとも1つの脆弱性パターンを抱えており、実行可能スクリプトをバンドルしたスキルは影響を受ける可能性が2倍であることが判明した。
もう 1 つのエントリがこの年代記に属しており、それは私たちのものであるため、その関心は暗示されるのではなく宣言されます。上記の各信頼スタックは、独自のシェルフのみを保証します。NVIDIA は NVIDIA が公開するものを検証し、Google のレジストリは企業が内部に保存しているものを保持し、skills.sh はリストにあるものを監査します。ビルドアウトにはレジストリ全体を保証するものは何もなく、企業の門外の開発者に届くものはほとんどありません。その中立的なコーナーは、TomeVault が取り組んでいるコーナーです。すべての主要なフォーマットとすべてのソースに同じスキャンとグレードが適用され、個人とチームにとって同様に無料であり、彼らが生成するコーパスはパート II の測定対象となります。そのような関心を視野に入れて数字を判断してください。
Libraries.io と Ecosyste.ms がパッケージ エコシステムのマッピングに 10 年を費やした Andrew Nesbitt 氏は、次のような文章でシーズンを締めくくりました。

「信頼できる出版と出所証明は、パッケージレジストリが到達した答えである。スキルに相当するものはほとんど存在しない。」
ソース: 各行は、参照ブロック内のアクセス日付を含む一次ソースを追跡します。ステータスは 2026 年 6 月 11 日時点のスナップショットであり、いくつかは明示的に進行中です。
なぜそれが重要なのか。市場は、ファイル形式ではなく、カタログ、レジストリ、インストール ツールを信頼するかどうかを決定しています。ファイルを配布する人は誰でも、それを保証するよう求められます。パート II では、その野心がオープン コーパスにどこまで到達したかを測定します。簡単に言うと、そうではありません。
信頼層が存在する脅威は、2 月 1 日に仮説ではなくなりました。
ClawHub は、この冬最も急速に成長しているオープン エージェントの 1 つである OpenClaw にサービスを提供するスキル マーケットプレイスです。 2 月 1 日、Koi Security は、当時市場に出回っていた 2,857 のスキルのうち、341 が悪意のあるスキルであり、そのうち 335 は ClawHavoc と名付けられた単一のキャンペーンによるものであることを明らかにしました。このメカニズムは、指示ファイル自体に埋め込まれたソーシャル エンジニアリングでした。SKILL.md は、インストールにはコンパニオン ユーティリティが必要であることをユーザーに通知します。このユーティリティはリンクから取得され、場合によってはパスワードで保護されたアーカイブ内に保存されます。ペイロードは、macOS 認証情報スティーラー、リバース シェル、

[切り捨てられた]

## Original Extract

The annual report on AI instruction files: the year the format gained governance, marketplaces, scanners, and its first poisoning campaign, and the data underneath it. 267,228 repositories measured: the multi-tool gap, licensing, safety, and a signing rate of zero.

State of AI instructions, 2026 · Standards | TomeVault TomeVault Home Vault Scan Search Docs Standards Wtf is… Sign in TomeVault · Standards Annual Report · 2026 Annual edition · 2026
This was the year the AI instruction file became infrastructure: foundation-governed, marketplace-distributed, scanned, signed in vendor catalogues, and poisoned at scale for the first time. This report is the chronicle of that year and the measurement underneath it, drawn from 267,228 public repositories.
of repositories configure only one AI coding tool, while the developers behind them use several.
267,228 repositories · every major format · snapshots frozen 2026-06-04 and 2026-06-11 · numbers reproducible
03 Trust at the distribution layer
01 78.7% of repositories that use AI coding tools configure only one of them.
Their developers do not stop at one. Most run two to four AI tools, so every other assistant is left working that codebase with nothing to go on.
02 AGENTS.md was built to bridge tools. Only 18.3% of the repositories that use it or CLAUDE.md use both.
The one cross-vendor standard is being adopted as a replacement, not a bridge. It moves the split, it does not close it.
03 Nine in ten repositories that ship both CLAUDE.md and AGENTS.md never connect them, so the second file is never read.
The team maintains two contracts and gets the coverage of one, then blames the model for ignoring rules it was never handed.
04 26.3% of CLAUDE.md files run past the 200-line guideline Anthropic publishes in its own docs.
Past that length a file does not just waste context. A 2026 study found it lowers the rate at which the agent finishes the job.
05 55.8% of instruction-file repositories grant no usable right to redistribute the file.
The year gave these files install commands, directories, and marketplaces. The majority of what they distribute, nobody granted permission to distribute.
06 Signing adoption in the open corpus is zero : of 986 probed repositories, none sign their instruction files.
The trust layer the industry built this year exists where a distributor requires it, and so far nowhere else. Trust is becoming a property of the channel, not the file.
07 5.0% of 219,024 scanned instruction files carry at least one deterministic safety finding.
ClawHavoc was the loud version. This is the quiet base rate, in files agents load by default with no scan at all.
The bottom line. The world around the instruction file grew up this year: governed standards, an install command, a trust stack, a threat model. The files themselves have not caught up. They serve one tool in a multi-tool world, they are unlicensed for the distribution they are getting, and the trust layer built for them is, in the open corpus, used by nobody.
configure only one AI coding tool
grant no right to redistribute the file
repositories in a 986-repo probe sign their instruction files
Twelve months ago an instruction file was a convention. This year it acquired a foundation, a distribution layer, a threat model, and its first poisoning campaign.
The instruction file stopped being a private convention between a developer and one tool. Over nine months it picked up neutral governance, a package-manager-style distribution layer, vendor catalogues, security taxonomies, and an incident record. The dates below are the spine of that story; the rest of Part I walks the three that matter most.
9 Dec 2025 The Linux Foundation forms the Agentic AI Foundation, with Anthropic’s Model Context Protocol, Block’s goose, and OpenAI’s AGENTS.md as founding projects. The instruction-file standard now has a neutral home.
Dec 2025 Anthropic releases Agent Skills as an open standard (agentskills.io): a folder with a SKILL.md that packages instructions for any agent that loads it.
15 Jan 2026 The first large-scale audit of the new format posts to arXiv. Of 31,132 skills analysed, 26.1% contain at least one vulnerability pattern.
20 Jan 2026 Vercel ships the skills CLI and skills.sh, a directory and leaderboard for skill packages. Instructions get an install command.
1 Feb 2026 Koi Security discloses ClawHavoc: 341 malicious skills on ClawHub, the marketplace serving the OpenClaw agent. The format’s first documented poisoning campaign.
17 Feb 2026 skills.sh adds automated security audits from three vendors (Gen Digital, Socket, Snyk); skills flagged malicious are hidden from search and leaderboards.
Feb 2026 ETH Zurich publishes the first controlled study of context files: over-specified files reduce agent success and add more than 20% cost. Lean beats long.
13–26 Mar 2026 A community RFC asks the Agent Skills spec to absorb cryptographic signing. The spec maintainer declines, pointing to the distribution layer instead.
16 Apr 2026 GitHub ships gh skill in public preview: install, search, publish, with version pinning and content-addressed change detection, deliberately not centralised verification.
22 Apr 2026 Google launches its official skills repository on the open Agent Skills format, announced on Day 1 of Cloud Next.
19 May 2026 NVIDIA announces Verified Agent Skills: scan, skill card, detached signature, verify at install, on the open spec. The same day, Google’s enterprise Skill Registry enters public preview.
3 Jun 2026 Andrew Nesbitt, who built Libraries.io and runs Ecosyste.ms, publishes the gap plainly: the trust machinery package registries spent a decade building mostly does not exist for skills.
The two formats that matter most ended the year under neutral governance or open specification, with the largest vendors publishing into them.
AGENTS.md spent its first year as an OpenAI repository convention. On 9 December 2025 it became a founding project of the Agentic AI Foundation, a directed fund under the Linux Foundation, alongside the Model Context Protocol and goose. By late February the foundation counted 146 member organisations. One caution for anyone citing this as settled: as of that date the foundation itself described the formal governance models for its three projects as still being defined, so AGENTS.md has a neutral home and a membership, not yet a finished constitution.
The Agent Skills format ran the same arc faster. Released by Anthropic as an open standard in December, it ended the spring as the format Google ships its official skills in, the format GitHub’s gh skill validates against, and the format NVIDIA chose for its verified catalogue. Google’s Gemini CLI documentation states it directly: a skill is “based on the Agent Skills open standard.” When the largest vendors publish into a competitor-originated format rather than forking it, the format has stopped being anyone’s product.
The one asymmetry worth recording sits inside Google. Its public repository is openly on the standard; its enterprise Skill Registry, in public preview since 19 May, requires the SKILL.md package structure but nowhere names the open specification, and describes itself as a secure, private repository. The format travels; the catalogue is the product. That split, open format underneath and proprietary distribution on top, is the shape most of the year’s vendor moves share.
Why it matters. Teams standardising their instruction files in 2026 are no longer betting on a vendor’s whim. The formats are governed or open, and the year’s entrants built on them rather than around them. The open question has moved up a layer: not what the file looks like, but who vouches for it. That is the next section.
Trust arrived at the distribution layer, not in the spec
The year’s defining argument was about where trust should live. The spec said: not here.
In March, a community RFC proposed building cryptographic signing into the Agent Skills specification itself: an ed25519 signature block in SKILL.md frontmatter, with publisher keys discoverable at a well-known URL. The spec maintainer declined to absorb it. His position, stated on 16 March, was that signing is “better implemented at the distribution layer”; on 25 March he added that it is “unlikely that we’d make it an official part of the spec.” The discussion remains open, so this is a stated maintainer position rather than a closed ruling, but the spec has held: as of 11 June its frontmatter defines no signing, verification, or trust fields at all. The companion implementation tells the same story from the tooling side. A pull request adding skills verify to the distribution CLI was opened on 15 March and remains unmerged; no released version of the CLI carries a verify command.
Then in May, NVIDIA shipped exactly what the maintainer had pointed at. Its Verified Agent Skills programme is a trust stack built entirely at the distribution layer, on the open spec: an open-source scanner the company says checks for both conventional risks and agent-specific ones, a machine-readable skill card per skill, a detached signature in the OpenSSF Model Signing format, and verification at install or in CI. Every one of the 207 skills in its catalogue ships the full set: SKILL.md, skill card, signature. And the catalogue distributes through skills.sh, a competitor’s registry, with a plain npx skills add nvidia/skills .
The community asked the spec for trust. The spec pointed at the distributors. NVIDIA built it there two months later.
The rest of the trust layer is assembling around the same point. skills.sh added three-vendor automated audits in February, grading every listed skill and hiding the ones flagged malicious. OWASP opened an Agentic Skills Top 10 project in the spring, an incubator-stage proposal with ten draft risk categories and a 1.0 release targeted for the fourth quarter of 2026 on its own roadmap. And the academic record arrived in January: the first large-scale audit of skills in the wild found 26.1% of the 31,132 it analysed carrying at least one vulnerability pattern, with skills that bundle executable scripts twice as likely to be affected.
One more entry belongs in this chronicle, and it is ours, so the interest gets declared rather than implied. Each trust stack above vouches only for its own shelf: NVIDIA verifies what NVIDIA publishes, Google’s registry holds what an enterprise stores inside it, skills.sh audits what it lists. Nothing in the buildout vouches across registries, and little of it reaches the developer outside an enterprise gate. That neutral corner is the one TomeVault works from: the same scans and grades applied across every major format and every source, free for individuals and teams alike, and the corpus they produce is the one Part II measures. Judge our numbers with that interest in view.
Andrew Nesbitt, whose Libraries.io and Ecosyste.ms spent a decade mapping package ecosystems, closed the season with the sentence the whole year had been spelling out: “trusted-publishing and provenance attestations are the answers package registries arrived at; the equivalents for skills are mostly absent.”
Sources: each row traces to a primary source with access date in the References block. Statuses are snapshots as of 2026-06-11 and several are explicitly in motion.
Why it matters. The market has decided where trust attaches: to catalogues, registries, and install tools, not to the file format. Whoever distributes the file is being asked to vouch for it. Part II measures how far that ambition has reached into the open corpus. The short answer is: it has not.
The threat the trust layer exists for stopped being hypothetical on 1 February.
ClawHub is the skill marketplace serving OpenClaw, one of the fastest-growing open agents of the winter. On 1 February, Koi Security disclosed that of the 2,857 skills then on the marketplace, 341 were malicious, 335 of them from a single campaign it named ClawHavoc. The mechanism was social engineering embedded in the instruction file itself: a SKILL.md telling the user that installation required a companion utility, fetched from a link, sometimes inside a password-protected archive. The payloads were a macOS credential stealer, reverse shells, an

[truncated]
