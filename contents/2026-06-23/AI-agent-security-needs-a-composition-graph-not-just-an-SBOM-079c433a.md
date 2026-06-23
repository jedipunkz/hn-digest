---
source: "https://openaca.dev/blog/your-agent-risk-is-in-the-composition/"
hn_url: "https://news.ycombinator.com/item?id=48647802"
title: "AI agent security needs a composition graph, not just an SBOM"
article_title: "Your Agent Risk Isn't in One Plugin. It's in the Composition. — OpenACA"
author: "vinodkone"
captured_at: "2026-06-23T17:35:39Z"
capture_tool: "hn-digest"
hn_id: 48647802
score: 2
comments: 0
posted_at: "2026-06-23T16:53:47Z"
tags:
  - hacker-news
  - translated
---

# AI agent security needs a composition graph, not just an SBOM

- HN: [48647802](https://news.ycombinator.com/item?id=48647802)
- Source: [openaca.dev](https://openaca.dev/blog/your-agent-risk-is-in-the-composition/)
- Score: 2
- Comments: 0
- Posted: 2026-06-23T16:53:47Z

## Translation

タイトル: AI エージェントのセキュリティには SBOM だけでなく構成グラフが必要
記事のタイトル: エージェントのリスクは 1 つのプラグインにありません。それは「コンポジション」にあります。 — OpenACA
説明: SCA スキャナーは脆弱なパッケージを検出できます。できる

記事本文:
ブログ ·
2026年6月23日 · ヴィノッド・コーン
エージェントのリスクは 1 つのプラグインにありません。それは「コンポジション」にあります。
SCA スキャナーは脆弱なパッケージを見つけることができます。これらのパッケージが、チャット メッセージを読み取り、ファイルを送信し、アクセス ポリシーを書き換えることができるスキルによって管理される AI エージェントに接続されているとは言えません。
ワンクリックで Claude Code にインストールできるプラグインは次のとおりです。
それは imessage と呼ばれます。これはまさにあなたが望んでいることを実行します。エージェントが iMessage を読んで送信できるので、エージェントにテキスト メッセージを送信し、エージェントからテキスト メッセージを返信してもらうことができます。役に立つ。そして、それは公式のプラグイン マーケットプレイスで出荷されます。
ボンネットの下では、3 つのものがボルトで固定されています。 ~/Library/Messages/chat.db (メッセージ履歴全体) を読み取るローカル MCP サーバーにはフル ディスク アクセスが必要で、AppleScript を通じてユーザーに代わってメッセージを送信します。スキルのペア、configure と access は、 Read 、 Write 、およびスコープ指定された Bash 権限を保持し、アクセス制御状態を ~/.claude/channels に書き込みます。そして、そのすべての下に 91 個の npm パッケージがあります。
それらの作品はどれも、それ自体は素晴らしいものです。 MCP サーバーは指示どおりに動作します。スキルは言われたことを実行します。パッケージは通常の Web サーバーの依存関係です。リポジトリに対してソフトウェア構成分析スキャナーを実行すると、ロックファイル、 hono と path-to-regexp のいくつかのアドバイザリー、おそらく「中程度」のバッジなど、きちんとしたレポートが得られます。パッケージにパッチを適用すると、レポートが緑色になります。
そして、実際のリスクについてはほとんど何も学んでいないことになります。
なぜなら、リスクが単一のパッケージに含まれることは決してなかったからです。リスクは次のような構成です。信頼できない入力に接続されたエージェント (受信テキストはすべて、読み取られるメッセージ ストアに到達します。エージェントへのアクセスはデフォルトで許可リストでゲートされていますが、表面はそこにあります)、プライベート メッセージ履歴を読み取ったり、混乱を送信したりする可能性があります。

誰が運転できるかを書き換えることができるスキルによって管理されており、すべてが 1 つのコンテキスト内で、それぞれ単独でレビューされた部品から組み立てられています。
それは仮説ではありません。私たちは行って見ました。
公式の Claude プラグイン マーケットプレイスをスキャンしました
公式の Claude プラグイン マーケットプレイスで構成スキャナーを指定しました。62 のマニフェスト、530 のコンポーネント: 38 のプラグイン、27 のスキル、26 のコマンド、21 のエージェント、16 の MCP サーバー、8 のフック、およびその下の 394 のパッケージです。
124 件の既知の脆弱性に関する勧告が明らかになりました。その部分は興味深いものではありません。バンドルされた npm の依存関係がアップストリームの修正の背後に漂っています。これは公的なものであり、一般的なものであり、誰のゼロデイでもありません。 (明確にしておきます: これはいずれも、マーケットプレイスまたはその作成者に対する脆弱性の主張ではありません。これらは構成の暴露であり、レビューに値するエージェントの表面であり、悪用可能性の証拠ではありません。)
ここが興味深い部分です。 124 件の勧告はすべて、 discord 、 telegram 、 fakechat 、および imessage の 4 つのプラグインを追跡します。
4 つのメッセージ チャネル プラグイン。
アドバイザリは、信頼できない外部メッセージを取り込み、ローカル ファイルを送り返す可能性のあるプラグインに正確にクラスタリングされます。脆弱なコードと機密機能は同じコンポーネントです。これはまさに、パッケージごとのスキャンでは表現できない相関関係です。なぜなら、パッケージがメッセージ読み取りエージェントに属していることをそもそも認識していないからです。
ソフトウェア構成分析はライブラリ用に構築されました。この依存関係ツリーには既知の脆弱性のあるバージョンが含まれていますか?という 1 つの質問によく答えます。 imessage の場合、答えは「はい、いくつかあります」であり、それがすべてです。
最初の行は true です。 2 行目は、そのものをインストールするかどうかを決定する行です。ロックファイルはそれがロックファイルの一部であることを認識していないため、SCA はそれを生成できません。

エージェント。
もう 1 つは、「実行時にキャッチする」という反射です。エンドポイントと動作の監視は現実的かつ補完的であり、一部のコントロールは実行時にブロックされる可能性があります。ただし、それらは、エージェントをインストールまたはレビューした時点で宣言された構成ではなく、発生したプロセスの動作に基づいて動作します。彼らは飛行中に悪い行動を見つけることができます。そもそも、このプラグインの構成によってそのアクションが実行可能になるかどうかは、インストールする前に教えてもらえません。
したがって、2 つのビューがあり、どちらにも重要なオブジェクトが表示されません。
SCA はパッケージを認識します。ランタイム監視により動作が監視されます。どちらも、宣言されたエージェントの構成、つまりプラグイン、スキル、MCP サーバー、フック、権限、すべてが組み込まれたインストール パスを認識しません。
この 3 番目のビューは、エージェントのセキュリティが実際に必要とする 1 つのビューです。そして今のところ、それを生み出すものはほとんどありません。
この修正は、より優れたパッケージ スキャナーではありません。それは別の分析単位です。
エージェント スタックは依存関係のフラットなリストではなく、グラフです。プラグインには、スキル、MCP サーバー、フックが含まれています。これらは機能を宣言し、パッケージをプルします。権限とインストール パスは途中で付加されます。重要なことの一部はノードローカル (パッケージのバージョン、スキルの許可ツール、MCP トランスポート) です。しかし、判断を変える事実は、多くの場合、ノードだけでなくノード間のエッジに存在します。
その絵は手で描かれたものではありません。構成の骨子とアドバイスの帰属はスキャンから直接得られます。機能ラベル (各コンポーネントが読み取り、送信、または変更できる内容) は、プラグインのソースを読み取ることで得られます。それを読めば、その論旨は明らかです。
エージェント システムでは、セキュリティ オブジェクトは単なるパッケージではなくなりました。パッケージにプラグイン、スキル、MCP サーバー、フック、権限、インストールを加えたものです

l パス、およびそれが構成される実行時コンテキスト。
フラット SBOM にはボックスがリストされます。赤いボックス (スキャナーが認識するすべてのもの) が、信頼できない入力ソース、プライベート データ リーダー、送信シンクと同じコンテキストに存在するかどうかはわかりません。グラフはそれが可能です。
これにより、スキャナーのフラグが個別に再フレーム化されます。マーケットプレイス全体で、変更可能な参照から起動する MCP サーバーが 10 台見つかりました。そのうち 5 台は外部参照 ( npx …@latest 、固定されていない uvx git ref、ダイジェストのない Docker タグ) から、5 台はマニフェストを変更せずにターゲットを変更できるローカル コマンド ランチャー ( bun run 、php 職人 ) からでした。実行可能ツールへのアクセスを宣言するスキルが 7 つ見つかりました。それぞれ単独では脚注です。グラフでは、「可変ランチャー」と「実行可能機能」と「同じコンテキストでの信頼できない入力」が、読む価値のある文章です。
OpenACA の今日の取り組みと今後の取り組み
OpenACA は、このレイヤーをオープンに構築するオープンソース プロジェクトです。
完全なエージェント スタック (プラグイン、スキル、MCP サーバー、フック、エージェント、コマンド、およびその下のパッケージ) のインベントリを作成します。
すべてのパッケージ アドバイザリを、それを取り込むエージェント コンポーネントに関連付けます (これにより、124 のアドバイザリすべてが 4 つのプラグインに存在することがわかります)。
変更可能なインストールや実行可能スキルの機能などの姿勢の問題にフラグを立てます。
全体をエージェント BOM (CycloneDX) として出力します。構成グラフ、エクスポート可能です。
次に行われるのは、グラフ由来のエクスポージャー分析です。「このコンテキストは、信頼できない入力、プライベート データの読み取り、およびアウトバウンド シンクを組み合わせています。確認してください」という構成によってのみ起動されるルールと、影響を与えるための妥当なパスが存在するかどうかによる調査結果の優先順位付けです。 OpenACA はまだこれらの危険経路を計算していません。インベントリ、アトリビューション、グラフはそれらを可能にする基盤であり、

それが私たちが目指していることです。
これが正直な売り文句です。「私たちは今日エージェントのエクスプロイトを検出します」ではなく、「私たちはエージェントの構成を可視化します。これはその後のすべての前提条件です。」
これが 32 番目の演習です。エージェント プラグイン リポジトリで SCA スキャナを実行します。次に、次のように質問します。
クロード スキル、ライフサイクル フック、またはローカル ファイルを送信できる MCP サーバーを通じて、どの脆弱な依存関係に到達できるかを知ることができますか?
それができない場合は、エージェントは表示されず、そのロックファイルが表示されます。
uvx --from openaca openaca scan repo --target 。 --include-姿勢
プラグイン、スキル、または MCP サーバーを含む任意のリポジトリにこれを指定すると、何がインストールされているか、コンポーネントがどのように接続されているか、どのアドバイザリとポスチャ結果がどのエージェント コンポーネントに属しているかなどの構成が表示されます。それから、何が落ちたか教えてください。
エージェントのセキュリティには、パッケージのスキャンだけでなく、構成を意識した分析が必要です。パッケージは決して本体ではありませんでした。エージェントは。

## Original Extract

Your SCA scanner can find vulnerable packages. It can

Blog ·
June 23, 2026 · Vinod Kone
Your Agent Risk Isn't in One Plugin. It's in the Composition.
Your SCA scanner can find vulnerable packages. It can't tell you those packages are wired into an AI agent that reads your chat messages, sends files, and is governed by skills that can rewrite its access policy.
Here’s a plugin you can install into Claude Code in one click.
It’s called imessage . It does exactly what you’d hope: it lets your agent read and send iMessages, so you can text your agent and have it text you back. Useful. And it ships in the official plugin marketplace.
Under the hood, it’s three things bolted together. A local MCP server that reads ~/Library/Messages/chat.db — your entire message history — requires Full Disk Access, and sends messages on your behalf through AppleScript. A pair of skills, configure and access , that hold Read , Write , and scoped Bash permissions and write access-control state into ~/.claude/channels . And ninety-one npm packages underneath all of it.
Every one of those pieces is, on its own, fine. The MCP server does what it says. The skills do what they say. The packages are ordinary web-server dependencies. Run a software composition analysis scanner over the repo and you’ll get a tidy report: a lockfile, a handful of advisories in hono and path-to-regexp , maybe a “moderate” badge. Patch the packages and the report goes green.
And you’ll have learned almost nothing about your actual risk.
Because the risk was never in any single package. The risk is the composition : an agent wired to untrusted input (every inbound text lands in the message store it reads — admission to the agent is allowlist-gated by default, but the surface is there), that can read your private message history, send messages as you, and is governed by skills that can rewrite who’s allowed to drive it — all in one context, assembled from parts that were each reviewed alone, if at all.
That’s not a hypothetical. We went and looked.
We scanned the official Claude plugins marketplace
We pointed a composition scanner at the official Claude plugin marketplace — 62 manifests, 530 components: 38 plugins, 27 skills, 26 commands, 21 agents, 16 MCP servers, 8 hooks, and 394 packages underneath them.
It surfaced 124 known-vulnerability advisories. That part isn’t interesting — bundled npm dependencies drift behind upstream fixes; this is public, ordinary, and not anyone’s 0-day. (To be clear: none of this is a vulnerability claim against the marketplace or its authors. These are composition exposures — review-worthy agent surfaces, not proof of exploitability.)
Here’s the interesting part. All 124 advisories trace to exactly four plugins: discord , telegram , fakechat , and imessage .
The four message-channel plugins.
The advisories cluster precisely in the plugins that ingest untrusted external messages and can send local files back out. The vulnerable code and the sensitive capability are the same components — which is exactly the correlation a package-by-package scan can’t express, because it never knew the packages belonged to a message-reading agent in the first place.
Software composition analysis was built for libraries. It answers one question well: does this dependency tree contain a known-vulnerable version? For imessage , the answer is “yes, a few,” and that’s the whole story it can tell.
The first row is true. The second row is the one that decides whether you should install the thing. SCA can’t produce it, because a lockfile doesn’t know it’s part of an agent.
The other reflex is “we’ll catch it at runtime.” Endpoint and behavioral monitoring are real and complementary — and some controls can block at execution time. But they operate on process behavior as it happens, not on the agent’s declared composition at the moment you install or review it. They can catch a bad action mid-flight; they can’t tell you, before you install, that this plugin’s composition is what makes that action reachable in the first place.
So we have two views, and neither sees the object that matters:
SCA sees packages. Runtime monitoring sees behavior. Neither sees the declared agent composition — the plugin plus the skill plus the MCP server plus the hook plus the permissions plus the install path it’s all assembled into.
That third view is the one agent security actually needs. And right now almost nothing produces it.
The fix isn’t a better package scanner. It’s a different unit of analysis.
An agent stack isn’t a flat list of dependencies — it’s a graph. Plugins contain skills, MCP servers, and hooks. Those declare capabilities and pull in packages. Permissions and install paths attach along the way. Some of what matters is node-local — a package version, a skill’s allowed-tools, an MCP transport — but the facts that change your verdict often live in the edges between them, not just the nodes:
That picture isn’t drawn by hand: the composition skeleton and the advisory attribution come straight from the scan; the capability labels — what each component can read, send, or change — come from reading the plugin’s source. Read it and the thesis is obvious:
In agent systems, the security object is no longer just the package. It is the package plus the plugin, skill, MCP server, hook, permissions, install path, and runtime context it is composed into.
A flat SBOM lists the boxes. It can’t tell you that the red box (all your scanner sees) sits in the same context as an untrusted-input source, a private-data reader, and an outbound sink. The graph can.
This also reframes the things scanners flag in isolation. Across the marketplace we found 10 MCP servers that launch from a mutable reference — five from external refs ( npx …@latest , an unpinned uvx git ref, a Docker tag with no digest) and five from local command launchers ( bun run , php artisan ) whose target can change with no manifest change. We found 7 skills that declare executable tool access. On their own, each is a footnote. On the graph, “mutable launcher” plus “executable capability” plus “untrusted input in the same context” is a sentence worth reading.
What OpenACA does today — and what’s next
OpenACA is an open-source project building this layer in the open.
inventory the full agent stack — plugins, skills, MCP servers, hooks, agents, commands, and the packages beneath them;
attribute every package advisory to the agent component that pulls it in (that’s how we know all 124 advisories live in four plugins);
flag posture issues like mutable installs and executable-skill capability;
emit the whole thing as an Agent BOM (CycloneDX) — the composition graph, exportable.
What’s next is graph-derived exposure analysis: rules that fire only because of the composition — “this context combines untrusted input, a private-data read, and an outbound sink; review it” — and prioritization of findings by whether a plausible path to impact even exists. OpenACA doesn’t compute those exposure paths yet. The inventory, attribution, and graph are the foundation that makes them possible, and that’s what we’re building toward.
That’s the honest pitch: not “we detect agent exploits today,” but “we make the agent’s composition visible, which is the prerequisite for everything after.”
Here’s a thirty-second exercise. Run your SCA scanner on an agent plugin repo. Then ask it a question:
Can it tell you which vulnerable dependency is reachable through a Claude skill, a lifecycle hook, or an MCP server that can send local files?
If it can’t, you’re not seeing your agent — you’re seeing its lockfile.
uvx --from openaca openaca scan repo --target . --include-posture
Point it at any repo with plugins, skills, or MCP servers and it’ll show you the composition: what’s installed, how the components are wired together, and which advisories and posture findings belong to which agent component. Then tell us what fell out.
Agent security needs composition-aware analysis, not just package scanning. The package was never the unit. The agent is.
