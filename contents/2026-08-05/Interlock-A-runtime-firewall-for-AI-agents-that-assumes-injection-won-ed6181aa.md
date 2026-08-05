---
source: "https://yashwanthreddymali.com/blog/interlock-exfiltration-at-runtime/"
hn_url: "https://news.ycombinator.com/item?id=49185796"
title: "Interlock: A runtime firewall for AI agents that assumes injection won"
article_title: "Interlock: a runtime firewall for AI agents that assumes prompt injection already won | Yashwanth Reddy Mali"
author: "yxshwanthreddy"
captured_at: "2026-08-05T17:20:28Z"
capture_tool: "hn-digest"
hn_id: 49185796
score: 2
comments: 0
posted_at: "2026-08-05T17:11:47Z"
tags:
  - hacker-news
  - translated
---

# Interlock: A runtime firewall for AI agents that assumes injection won

- HN: [49185796](https://news.ycombinator.com/item?id=49185796)
- Source: [yashwanthreddymali.com](https://yashwanthreddymali.com/blog/interlock-exfiltration-at-runtime/)
- Score: 2
- Comments: 0
- Posted: 2026-08-05T17:11:47Z

## Translation

タイトル: Interlock: インジェクションが成功したことを前提とした AI エージェント用のランタイム ファイアウォール
記事のタイトル: Interlock: プロンプト インジェクションがすでに成功していることを前提とした AI エージェント用のランタイム ファイアウォール |ヤシュワンス・レディ・マリ
説明: 漏洩検出器の構築に関して、書き換えを余儀なくされた 46.7% の誤検知率と、私が独自のツールで捕捉できないもののリストを公開する理由について説明します。

記事本文:
Interlock: プロンプト インジェクションがすでに成功していることを前提とした AI エージェント用のランタイム ファイアウォール
流出検出器の構築について、書き換えを余儀なくされた 46.7% の誤検知率、および私が独自のツールで捕捉できないもののリストを公開している理由について説明します。
これは 2 年前の私にとって何の意味もなかったであろう一文です。
資格情報ファイルを作成し、汚染された GitHub の問題を吸収し、両方を Webhook に投稿しようとしました。
そして私が書いたソフトウェアは、これら 3 つの出来事すべてが起こるのを監視し、
バイトは同じバイトだったので、プロセスを強制終了しました。
そのソフトウェアはInterlockと呼ばれます。これは、その機能とその理由の長いバージョンです。
デザインをそのまま形にしています。
流し読みする人向けのTL;DR (正しい本能):
Interlock はプロンプト注入を停止しようとはせず、注入が成功したものとみなし、注入の目的を監視します。
検出は、JSON-RPC ワイヤ上の MCP プロキシとシステムコール上の eBPF センサーの 2 つのプレーンで実行されます。
ハードエンフォースメントは、汚染されたシークレットとアウトバウンドトラフィックの間のバイトオーバーラップの証明に対してのみ実行され、セッション形状単独では実行されません。
測定値: EXFIL 層検出 100% (非ギャップ悪意のある 31/31)、EXFIL 層誤検知 0% (良性 0/37)、任意のトリップ ソフト アラート 18.9%
構造的に把握できないものについては公開台帳があり、その台帳を最初に読んでいただきたい部分です
AI エージェントが役立つのは、構造上の理由がまさに 1 つあります。それは、AI エージェントがユーザーの情報を読み取って行動できることです。
あなたに代わって。ファイルを読み取り、API を呼び出し、参照し、外部ツールにプラグインします。
MCP (モデル コンテキスト プロトコル、本質的にはエージェント機能用の電源タップ)。それは
ピッチ全体。モデルに手を渡します。
問題は、次の 3 つの機能が同じ部屋に存在することが想定されていなかったことです。
社外とのコミュニケーション能力
サイモン・ウィリソンはこれを致命的な三連単と名付けた

。片足だけでも
通常の製品の動作。同じセッションで点灯した 3 つすべてが構造的に次の形状になります。
データの流出。
そして、通常はツール中毒によって消滅しますが、それほど劇的なものではありません
聞こえるよりも。攻撃者の指示はツールの結果内に届きます。モデル
言語モデルにはクリーンなコンテキストがないため、その結果を信頼できるコンテキストとして扱います。
「読むべき文章」と「従うべき文章」を分ける仕組み。あとあと、
個別に承認されたツール呼び出しにより、秘密が実行されます。その連鎖のすべてのステップ
単独では合理的に見えます。それが難しいところです、そして、私も認めますが、ある種のことです
ロックピックが美しいのと同じように美しい。
プロジェクト全体がかかっている決断
プロンプトインジェクションを解決しようとはしませんでした。
インジェクションは、信頼できないテキストと特権命令を混合する特性です。
同じ文脈。ツールの静的ホワイトリストやツール スキーマの事前呼び出しスキャナーはありません
エージェントが実行されてライブ結果を読み取ると、その混合が除去されます。 「ストップ」の扱い
あなたのプライマリコントロールが武器を獲得するためにサインアップしているため、モデルは影響を受けません。
ゴールラインもスコアボードもない、自然言語との永久的な競争です。
そこでコントロールプレーンを移動させました。インターロックは注入を周囲条件として扱います。
エージェントはシステムを起動し、シンクでより限定的な質問をします。
シークレットの実際のバイト、またはそれらのバイトの登録されたエンコーディングは、
アウトバウンドトラフィック?
これは、「この行動は疑わしいかどうか」よりもはるかに小さな主張であり、少額の主張は
実際に保管できるもの。別のモデルを置かずにバイトオーバーラップを決定可能
ホットな道で。雰囲気は決められるものではありません。その選択の代償は明白であり、
永続的: 秘密の意味上の言い換えは、構造上、永久に範囲外になります。
それをREADMEに記載するよりもむしろそう言いたいです

誰かが事件の調査でそれを発見します。
バイトが表示される場所がちょうど 2 つあるため、2 つのバリアント
VariantA Variant = "A_chained_tool" // プロキシhold-before-forwardによって捕捉される
VariantB Variant = "B_server_channel" // eBPF センサーによって捕捉されました
バリアント A は、連鎖されたツール呼び出しです。汚染された結果は後で問題を引き起こします
tools/call の引数にはシークレットが含まれます。そのトラフィックは JSON-RPC です
ワイヤーインターロックが存在できるため、プロキシは転送前に評価し、
電話をきっぱり拒否します。
バリアント B はサイド チャネルです。 MCP サーバー プロセス (またはその子プロセス)
spawns) は独自のソケットを開き、シークレットを直接書き込みます。ツールとしては決して登場しない
引数。プロキシは構造的にそれに対して盲目です。それがカーネルセンサーの目的です。
2 つの飛行機、そしてそれぞれには見えないもの
2 つの飛行機は建築上の繁栄ではなく、脅威モデルが現時点で必要としているものです
両方の exfil パスを同じページに描画します。 1 つの飛行機は攻撃対象領域の半分を意味します
目に見えず、どちらの半分かを選択することはできません。
Hold-before-forward、および遅延が重要な理由
プロキシは、フレームを
子プロセス。拒否すると、JSON-RPC エラーが合成され、呼び出しは行われません。オン
許可すると、転送して応答を待ち、IngestResult を実行します。
エージェントが結果を確認する前に遡ります。
代わりの方法、forward-then-inspect は安価ですが、役に立ちません。ブロックするファイアウォールです。
これは、シークレットがすでに到達した後で、子がバリアント A をすでに失っていることを検査します。
エージェントは評価とバックエンドの往復を待機します。約束を守るために待ち時間を支払う
正直に「阻止した」し、その取引は正しいと思いますが、p99を見つめました
私が認めたいよりも長い数字。
セッションごとの状態は 3 つのレッグを追跡します。
足が腐ってしまう。デフォルトのレッグ TTL は 30 分です

es、または 32 回の呼び出しのいずれか最初にヒットするもの。セッション
午前 9 時に 1 つの構成ファイルを読み取った場合、正午になってもソフトトリップすることはありません。
なぜ 2 つではなく、4 つではなく 3 つなのでしょうか。以前のデザインでは粘着脚が使用されていました
コンテンツ要件はありません。これにより、次のセクションの数字が生成されます。そして
「関連性」、「エンコード形式」、および「コンテナーの中止」は意図的に脚部ではありません。
これらは分類子内の評価時修飾子です。それらを脚に昇格させると、
スティッキー コンテンツ ブラインド状態を再導入します。これはまさにバージョン 0 を破壊したものです。
書き換えを強制した番号
バージョン 0 は明らかな方法で機能しました。 3 つのスティッキー ブール値。機密データに触れた、設定された
旗。信頼できないコンテンツが検出されました。フラグを設定してください。外部シンクが呼び出され、フラグが設定されます。 3つとも
点灯すると、呼び出しがブロックされ、プロセスが強制終了されます。
それは完全に合理的な最初のデザインです。また、コンテンツブラインドでもあります。
シークレットが実際に移動したかどうかを決して尋ねません。セッションが行われたかどうかのみを尋ねます。
ある時点で、3 つの機能すべてを保持していました。
あらゆる旅行の誤検知率は 46.7%。
正規のセッションのほぼ半数がトリップしました。そしてもちろん彼らはそうしました。考えてみましょう
火曜日の通常のコーディング エージェント セッションはどのようなものですか:
~/.config/something.yaml でプロジェクトのセットアップ (機密性の高い部分) を確認します。
ドキュメント ページをフェッチするか、GitHub の問題 (信頼できないレッグ) を読み取り、API を呼び出します。
ホスト名 (シンクレッグ) を解決します。三本足。建物からは何も残らなかった。何も平らではなかった
出発が近づいています。とにかくハードブロック。
偽陽性率は、ブロック コントロールの美的指標ではありません。それは
存在指標。正当な作業の半分をブロックするファイアウォールがアンインストールされる
1 週目以降は、製品が実行されていないため、偽陰性率は 100% になります。
修正: 評決と訴訟は 2 つの異なる質問です
書き換えにより 2 つのことが分離されました

以前は 1 つの列挙型でした:
評決の答え: このシンクイベントについて私たちは何を結論付けましたか?
行動の答え: 執行は実際に何をしたのか?
それらを折りたたむと、負荷がかかることが判明した 3 つの動作 (モニター モード、
ハードブロッキングなしのソフト検出、およびリークの防止と漏洩の区別
最初のパケットの後に 1 つが含まれます。
順序付けが設計であるため、分類子はそのままです。
// 内部/エンジン/engine.go -classifyTrip
重複する場合 != nil {
return model.VerdictExfil, 0.95, true // AllLit チェックの前
}
if !state.Legs.AllLit() {
""、0、false を返す
}
// … この後のみソフト SUSPICIOUS ゲート …
それをよく読んでください。 EXFILはオーバーラップ単体です。 3 本の脚すべてに取り付けることができます。
完全に減衰しても、テイントが保持され、バイトが一致する限り、EXFIL は引き続き発火します。
逆に、バイト証明なしで点灯している 3 つの脚すべては SUSPICIOUS です。
殺すのではなく警告します。
バリアント A (プロキシ、防止可能):
バリアント B (センサー、のみを含むことができる):
予防されたものと封じ込められたものとの間の語彙のギャップ
見た目よりも重要です。 1 つは、パケットが決して残されていないことを意味します。もう 1 つは、出発したことを意味し、
その後、プロセスが終了しました。両方を「キャッチ」と呼ぶと、唯一の違いが消えてしまいます。
インシデント対応担当者は実際に気にかけています。
残りの 7 つは設計によって固定されており、ハード アラートではなく、オペレーターに表示されるソフト アラートによって固定されます。
執行の失敗。
この投稿では「Overlap」が多くの取り組みを行っていますので、実際の内容を紹介します。
機密性の高い結果が返されると、Interlock は汚染された値を抽出します (ストライプ スタイル)
キー、ベアラー トークン、PEM 秘密キー ブロック、PuTTY キー ファイル、アカウント ID)、および
それぞれの正規エンコーディングの閉じたセットを事前計算します。
リテラル、base64、16 進数、url_encoded、反転、
base64_hex、hex_base64、base64_url、base64_reversed、
gzip_base64、brotli_base64、zstd_base64、lz4_base64
クローズドセット

、無制限の変換検索ではありません。それは製品に関する決定であり、怠惰ではありません。
オープンエンドの検索スペースにより、圧縮されたファイルに対して起動する検出器が完成します。
JPEG。カスタム暗号は名前付きの永続的なギャップなので、紙で書き直すよりも名前を付けたいと思います
それ。
ミスの場合、有界再帰デコーダは、base64 と深さ 5 までの 16 進数をアンラップします (深さ 5 に固定されています)。
[3, 5] 、オペレーターにデコード深度を無限に設定させると、
攻撃者は設定フラグを持つ CPU シンクを攻撃します）。 EXFIL 層が false であるため、深さ 5 が選択されました。
深さ 3、4、5 では陽性率は 0.0% に留まり、デコードミスのレイテンシーはほぼ一定でした。
380 ～ 390 マイクロ秒。
クロスコール フラグメント再アセンブリ (ページ分割された結果全体での秘密の分割) もあります。
半分が隣接している場合でもキャッチされます)、断片化された DNS の出力フローの再構成、および
書き込み、ハードボムキャップを使用した ZIP/gzip/zlib/tar へのコンテナ降下、およびチャンクマッチング
長いシークレットの場合、その完全な形式はカーネル キャプチャ ウィンドウによって切り捨てられますが、
32 バイトのボディ チャンクが引き続き表示されます。
それぞれが特定の攻撃形態を閉じ、CPU やメモリを消費したり、誤検知を引き起こします。
表面に残り、名前付き残差が残ります。それが正直な説明であり、あらゆる緩和策は次のとおりです。
毎回どの取引を行ったかを書き留めようとしました。
カーネル側、そして私が過度に主張したくないことが 1 つあります
センサーは syscall トレースポイントを接続し、2 つの個別のリング バッファーを排出します。
分離されているため、connect() ストームによって実際のデータを運ぶリングが枯渇することはありません
ペイロードの証拠。残りの部分は、はっきり言って、フラッディング書き込みは依然として行われるということです。
クリティカルリングに圧力をかけます。既知のギャップ、テストで固定。
ペイロードのキャプチャは 1024 バイト ( PAYLOAD_MAX ) に制限されています。
その窓を完全に超えて生きている秘密は、抜粋には決して現れず、明らかにすることはできません
一致しました。チャンクマッチングはこれを狭めます。閉じません。
オプションの LSM があります

security_socket_connect にフックします。これがそのです
これはまさにランディング ページで過剰に販売されている機能であるため、正直な範囲です。
EXFILがすでに確認された後、repeat connect()を停止します
検出を早期に移動することはなく、connect() はアプリケーション ペイロードを運びません
最初の EXFIL を運ぶパケットは常に contains_by_kill され、カーネルによって阻止されることはありません。
オプトインであり、カーネル構成が必要ですが、ほとんどのディストリビューションはデフォルトでは出荷されていません
文字通り名前が付けられたテストがあります
TestIngestSyscall_FirstPacketStillContained_KnownGap 。最初のパケット
制限はアーキテクチャ上のものであり、唯一の本当の防止策はネットワーク名前空間です
生成時の分離。Interlock がプロセスを実行している場合にのみ機能します。
産卵。
これは私がほとんど公開していなかった部分であり、今では私が最も思い入れのある部分です。
私は Interlock が把握できないすべてのリストを 3 つに分類して保管しています。
これらのカテゴリはまったく異なる意味を持っているためです。
閉店しました。隙間はありませんでした、これがメカニズムです。粘着質
コンテンツブラインドブロッキング、コンプレッサーを含むトークンエンコーディング、クロスコールスプリット、
深さ 5 のネスト、DNS 断片化、ZIP および XLSX 内部、証拠改ざん。
デマンドゲート型。構築できるかもしれないが、意図的に構築されていない。

[切り捨てられた]

## Original Extract

On building an exfiltration detector, the 46.7% false positive rate that forced a rewrite, and why I publish a list of things my own tool cannot catch.

Interlock: a runtime firewall for AI agents that assumes prompt injection already won
On building an exfiltration detector, the 46.7% false positive rate that forced a rewrite, and why I publish a list of things my own tool cannot catch.
Here's a sentence that would have meant nothing to me two years ago: my agent read a
credentials file, absorbed a poisoned GitHub issue, and tried to post both to a webhook,
and a piece of software I wrote watched all three of those things happen, proved the
bytes were the same bytes, and killed the process.
That software is called Interlock. This is the long version of what it does and why the
design is shaped the way it is.
TL;DR for people who skim (correct instinct):
Interlock does not try to stop prompt injection, it assumes injection succeeds and watches for the thing injection is for
Detection runs on two planes: an MCP proxy on the JSON-RPC wire, and an eBPF sensor on syscalls
Hard enforcement fires only on proof of byte overlap between a tainted secret and outbound traffic, never on session shape alone
Measured: 100% EXFIL-tier detection (31/31 non-gap malicious), 0% EXFIL-tier false positives (0/37 benign), 18.9% any-trip soft alerts
There is a public ledger of what it structurally cannot catch, and that ledger is the part I'd want you to read first
AI agents are useful for exactly one structural reason: they can read your stuff and act
on your behalf. They read files, call APIs, browse, and plug into external tools through
MCP (Model Context Protocol, essentially a power strip for agent capabilities). That's
the whole pitch. Give the model hands.
The problem is that three capabilities were never supposed to sit in the same room:
The ability to communicate externally
Simon Willison named this the lethal trifecta . Any one leg alone is
ordinary product behavior. All three lit in the same session is, structurally, the shape
of data walking out.
And it usually walks out through tool poisoning , which is less dramatic
than it sounds. Attacker instructions arrive inside a tool result . The model
treats that result as trusted context, because language models do not have a clean
mechanism for separating "text I should read" from "text I should obey." A later,
individually authorized tool call carries your secret out. Every step in that chain
looks reasonable in isolation. That's what makes it hard, and, I'll admit, kind of
beautiful in the way a lockpick is beautiful.
The decision the whole project hangs on
I did not try to solve prompt injection.
Injection is a property of mixing untrusted text with privileged instructions in the
same context. No static allowlist of tools and no pre-call scanner of tool schemas
removes that mixing once the agent is running and reading live results. Treating "stop
the model from being influenced" as your primary control is signing up to win an arms
race against natural language, permanently, with no finish line and no scoreboard.
So I moved the control plane. Interlock treats injection as the ambient condition of
agent systems and asks a narrower question at the sink:
Did the actual bytes of a secret, or a registered encoding of those bytes, appear in
outbound traffic?
That is a much smaller claim than "was this behavior suspicious," and small claims are
the ones you can actually keep. Byte overlap is decidable without putting another model
on the hot path. Vibes are not decidable. The cost of that choice is explicit and
permanent: semantic paraphrase of a secret is out of scope forever, by construction, and
I'd rather say that in the README than have someone discover it in an incident review.
Two variants, because there are exactly two places bytes are visible
VariantA Variant = "A_chained_tool" // caught by proxy hold-before-forward
VariantB Variant = "B_server_channel" // caught by eBPF sensor
Variant A is the chained tool call. A poisoned result causes a later
tools/call whose arguments contain the secret. That traffic is JSON-RPC on
a wire Interlock can sit on, so the proxy evaluates before forwarding and can
refuse the call outright.
Variant B is the side channel. An MCP server process (or a child it
spawns) opens its own socket and writes the secret directly. It never appears as tool
arguments. The proxy is structurally blind to it. That's what the kernel sensor is for.
Two planes, and what each one cannot see
Two planes is not architectural flourish, it's what the threat model requires the moment
you draw both exfil paths on the same page. One plane means half the attack surface is
invisible and you don't get to pick which half.
Hold-before-forward, and why the latency is the point
The proxy calls EvaluateRequest before it writes the frame to the
child process. On deny, it synthesizes a JSON-RPC error and the call never happens. On
allow, it forwards, waits for the response, and runs IngestResult on the
way back before the agent ever sees the result.
The alternative, forward-then-inspect, is cheaper and also useless: a blocking firewall
that inspects after the secret already reached the child has already lost Variant A. So
the agent waits on evaluate plus a backend round trip. I pay latency to keep the word
"prevented" honest, and I think that trade is correct, though I did stare at the p99
numbers for longer than I'd like to admit.
Per-session state tracks three legs:
Legs decay. Default leg TTL is 30 minutes, or 32 calls, whichever hits first. A session
that read one config file at 9am should not still be soft-tripping at noon.
Why three, and not two, and not four. An earlier design used sticky legs
with no content requirement, which produced the number in the next section. And
"relevance," "encoding form," and "container abort" are deliberately not legs,
they're evaluation-time qualifiers inside the classifier. Promoting them to legs would
reintroduce sticky content-blind state, which is the exact thing that broke version zero.
The number that forced the rewrite
Version zero worked the obvious way. Three sticky booleans. Sensitive data touched, set
a flag. Untrusted content seen, set a flag. External sink invoked, set a flag. All three
lit, block the call and kill the process.
It is a completely reasonable first design. It is also content-blind : it
never asks whether the secret actually moved. It only asks whether the session had, at
some point, held all three capabilities.
46.7% any-trip false positive rate.
Nearly half of legitimate sessions tripped. And of course they did. Think about
what a normal coding-agent session looks like on a Tuesday: it reads
~/.config/something.yaml to figure out your project setup (sensitive leg),
it fetches a docs page or reads a GitHub issue (untrusted leg), it calls an API or
resolves a hostname (sink leg). Three legs. Nothing left the building. Nothing was even
close to leaving. Hard-blocked anyway.
A false positive rate is not an aesthetic metric for a blocking control. It's an
existence metric. A firewall that blocks half of legitimate work gets uninstalled in
week one, and then the false negative rate is 100% because the product isn't running.
The fix: verdict and action are two different questions
The rewrite separated two things that were previously one enum:
Verdict answers: what did we conclude about this sink event?
Action answers: what did enforcement actually do?
Collapsing them breaks three behaviors that turn out to be load-bearing: monitor mode,
soft detection without hard blocking, and the distinction between preventing a leak and
containing one after the first packet.
The classifier, verbatim, because the ordering is the design:
// internal/engine/engine.go - classifyTrip
if overlap != nil {
return model.VerdictExfil, 0.95, true // before any AllLit check
}
if !state.Legs.AllLit() {
return "", 0, false
}
// … soft SUSPICIOUS gates only after this …
Read that carefully. EXFIL is overlap alone. All three legs can have
fully decayed and EXFIL still fires, as long as taint is retained and the bytes match.
Conversely, all three legs lit with no byte proof is SUSPICIOUS , which is an
alert, not a kill.
Variant A (proxy, can prevent):
Variant B (sensor, can only contain):
The vocabulary gap between prevented and contained_by_kill
matters more than it looks. One means the packet never left. The other means it left and
then the process died. Calling both "caught" would erase the only distinction an
incident responder actually cares about.
Those remaining 7 are pinned by design, operator-visible soft alerts, not hard
enforcement failures.
"Overlap" is doing a lot of work in this post, so here's what it actually is.
When a sensitive result comes back, Interlock extracts tainted values (Stripe-style
keys, bearer tokens, PEM private key blocks, PuTTY key files, account IDs) and
precomputes a closed set of canonical encodings for each:
literal, base64, hex, url_encoded, reversed,
base64_hex, hex_base64, base64_url, base64_reversed,
gzip_base64, brotli_base64, zstd_base64, lz4_base64
Closed set, not unbounded transform search. That's a product decision, not laziness: an
open-ended search space is how you end up with a detector that fires on compressed
JPEGs. Custom ciphers are a named permanent gap and I'd rather name it than paper over
it.
On a miss, a bounded recursive decoder unwraps base64 and hex up to depth 5 (clamped to
[3, 5] , because letting operators set decode depth to infinity is offering
attackers a CPU sink with a config flag). Depth 5 was chosen because EXFIL-tier false
positives stayed at 0.0% across depths 3, 4, and 5, and decode-miss latency stayed around
380 to 390 microseconds.
There's also cross-call fragment reassembly (a secret split across paginated results
still gets caught when the halves abut), egress flow reassembly for fragmented DNS and
writes, container descent into ZIP/gzip/zlib/tar with hard bomb caps, and chunk matching
for long secrets whose full form gets truncated by the kernel capture window but whose
32-byte body chunks still appear.
Each of those closes a specific attack shape, costs CPU or memory or false positive
surface, and leaves a named residual. That's the honest accounting, every mitigation is
a trade and I tried to write down which one I made each time.
The kernel side, and one thing I refuse to overclaim
The sensor attaches syscall tracepoints and drains two separate ring buffers:
Segregated so a connect() storm cannot starve the ring carrying actual
payload evidence. The residual, stated plainly, is that flooding writes still
pressures the critical ring. Known gap, with a test pinning it.
Payload capture is capped at 1024 bytes ( PAYLOAD_MAX ).
Secrets living entirely past that window never appear in the excerpt and cannot be
matched. Chunk matching narrows this. It does not close it.
There's an optional LSM hook on security_socket_connect , and here is its
honest scope, because this is exactly the feature that gets oversold on a landing page:
It stops repeat connect() after EXFIL was already confirmed
It does not move detection earlier, connect() carries no application payload
The first EXFIL-carrying packet is always contained_by_kill , never kernel- prevented
It's opt-in and needs kernel config most distros don't ship by default
I have a test literally named
TestIngestSyscall_FirstPacketStillContained_KnownGap . The first-packet
limitation is architectural, and the only real prevention answer is network namespace
isolation at spawn time, which only works when Interlock is the process doing the
spawning.
This is the part I almost didn't publish, and the part I'm now most attached to.
I keep a maintained list of everything Interlock cannot catch, sorted into three
categories, because those categories mean genuinely different things:
Closed. Was a gap, isn't anymore, here's the mechanism. Sticky
content-blind blocking, token encodings including compressors, cross-call splits,
depth-5 nests, DNS fragmentation, ZIP and XLSX interiors, evidence tampering.
Demand-gated. Could be built, deliberately isn't.

[truncated]
