---
source: "https://zeroquarry.com/research/models-capabilities/"
hn_url: "https://news.ycombinator.com/item?id=48553307"
title: "Evaluating different LLMs for their security research capabilities"
article_title: "Models and Their Capabilities | ZeroQuarry Research"
author: "eskibars"
captured_at: "2026-06-16T13:58:36Z"
capture_tool: "hn-digest"
hn_id: 48553307
score: 1
comments: 0
posted_at: "2026-06-16T10:56:38Z"
tags:
  - hacker-news
  - translated
---

# Evaluating different LLMs for their security research capabilities

- HN: [48553307](https://news.ycombinator.com/item?id=48553307)
- Source: [zeroquarry.com](https://zeroquarry.com/research/models-capabilities/)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T10:56:38Z

## Translation

タイトル: さまざまな LLM のセキュリティ調査能力の評価
記事のタイトル: モデルとその機能 |ゼロクオリーの研究
説明: オープンソースの March Madness プロジェクトからの実際の ZeroQuarry の結果を使用して、さまざまな LLM がサイバーセキュリティ スキャンでどのように実行されるかを実際に比較します。

記事本文:
ZeroQuarry を選ぶ理由
プラットフォーム
継続的なテスト
証拠と報告書
研究
価格設定
サインイン
スキャンをリクエスト ->
研究
シェーン・コネリー
2026 年 6 月 15 日
研究
サイバーセキュリティ
モデルとその機能
オープンソースの March Madness プロジェクトからの実際の ZeroQuarry の結果を使用して、さまざまな LLM がサイバーセキュリティ スキャンでどのように実行されるかを実際に比較します。
ZeroQuarry の構築とテストの一環として、私はさまざまなオープン ソース リポジトリにわたる「たくさん」のモデルを使用して「たくさん」のセキュリティ スキャンを実行しました。この記事の執筆時点では、サイバーセキュリティに関するさまざまなモデルとその機能について多くの誤解が渦巻いていますが、これらのモデルのいくつかを使用してセキュリティの問題を特定し、それらがすべて異なり、サイバーセキュリティのさまざまな「部分」でどのように優れている/劣っているかを特定するための「実際の」結果を示したいと思いました。
例を挙げて説明してみましょう。少し前に、 Seed Money というプロジェクトを構築しました。これは、March Madness プールでの優勝を有利に進めるための野心的なオープンソース プロジェクトでした。私はこれをラップトップにローカルに展開しましたが、さまざまなモデルでセキュリティ スキャンを実行する良い口実になると思いました。スキャンの結果をこのスプレッドシートに入力しました。このスプレッドシートの構造を以下で説明します。ここで注意しなければならないのは、このアプリは主にフロンティア LLM モデル自身によって開発されたものであるため、... いいえ。 LLM で構築されたアプリにはバグがないわけではありません。人間のプログラマと同じようにレビューが必要です。
スプレッドシートの構造に関する注意事項
「調査結果」タブと「時間とコスト」タブの 2 つのタブがあります。これらは缶に記載されていることを行います。
[調査結果] タブには、一連の潜在的な脆弱性が表示されます。セキュリティ研究では、通常、脆弱性を有効/境界線/無効 (またはそのバリエーション) に分類します。すべての「安全でないソフトウェア」が存在するわけではないため、

「ractice」は悪用可能な脆弱性をもたらします。これらの「潜在的な発見」は列 A で色分けされています。
LLM ごとに、特定の項目が見つかったかどうかを「はい」、「いいえ」、または「はい (拒否)」で示します。 「はい (拒否)」は、ZeroQuarry の動作方法に関係しています。ZeroQuarry には敵対的なレビュー ループがあり、「研究者」モデルが脆弱性を提案し、その後「ベンダー」モデルがそれが脆弱性であるかどうかに同意します。 「はい (拒否)」は、研究者モデルが脆弱性を発見したが、その後、敵対的レビューがその脆弱性に異議を唱えたことを意味します。なぜなら、それはいかなる識別可能な方法でも到達できなかったためであり、研究エージェントは同意したからです。私がこれに言及/表示したのは、マルチエージェント システムの重要性を示しているためです。LLM にコードベースの脆弱性を見つけて修正するように指示したとしても、この敵対的ループがなければ、さらに多くのノイズ/チャーンが発生する可能性があります。
次に、CVSS スコアに基づく重大度スコアがあります。 LLM はそれぞれのケースで自律的に重大度を判断しようとし、1 ～ 10 のスケールで評価されます。すべての LLM が情報提供のみであると判断したものはスプレッドシートには含めていません。
次に、「PoC Generated」行があり、モデルが脆弱性の悪用を試みるために PoC を生成する意思があるかどうかを示します。 「厳重なガードレールを備えたフロンティア LLM」が、ソフトウェアの脆弱性を悪用する PoC を積極的に生成する頻度がどれほど高いかを見ると、驚くかもしれません。これらの多くは、実際にどのようにプロンプトを提示するかによって決まります。
次に、「H1 適格とみなされる」行があります。これにより、ZeroQuarry の LLM ベースの評価ツールを通じて調査結果が実行され、モデルが典型的な HackerOne ルールに基づいて HackerOne に送信する「有効な」脆弱性とみなされるかどうかが示されます。たとえば、通常、CPU の枯渇によって DoS を引き起こす脆弱性は、

適格とみなされる。
それぞれのケースで色分けしました。緑は「正しかった/制限がなかった」ことを意味し、黄色は「境界線/判断の余地がある」を意味し、赤は「間違っていた/制限があった」を意味します。
ZeroQuarry が採用するさまざまなエージェント タイプを理解していれば、[コスト] タブは一目瞭然です。それについてはここで説明しているので、このブログでは繰り返しません。
有効/境界線の脆弱性
以下は「有効」であるとみなされます。LLM はそれらを見つけるのに「正しい」と考えられます。 「見つかった」行は、見つかったかどうかを示します。
*パブリック リフレッシュ デフォルト キー*: web/app.py#L21 はデフォルト キーをベイクします。これは明らかに置き換えることが意図されていますが、未構成のままにすることも可能です。それは脆弱性です。ほぼすべてのモデル (Anthropic の「低推論」バリアントを除く) でこの脆弱性が見つかります。
*シミュレーション パラメーターによる CPU DoS*: web/app.py#L71 で、ユーザーに提供するモンテカルロ シミュレーションの数を指定するよう求めます。デフォルトは 10000 ですが、ユーザーが数百万または数十億を入力するのを妨げるものはありません。これにより、CPU の枯渇によるサーバーの DoS が発生する可能性があります。繰り返しますが、Anthropic の「低推論」バリアントを除くほぼすべてのモデルでこの脆弱性が見つかります。
*ジョブ ステータス API はトレースバックを漏洩する可能性があります*: Flask サーバーでエラーが発生した場合、呼び出し方法のトレースがダンプされる可能性があります。これには本質的にひどいことは何もありませんが (したがって、それを「見つけた」ほとんどのモデルでカウントされるスコアが低いのですが)、ファイルや SQL コマンドのディスク上の場所などの追加情報が攻撃者に提供される可能性があり、攻撃対象領域に関する追加の洞察が得られる可能性があります。ほとんどのセキュリティ研究者は、これは修正する価値のある正当なセキュリティ バグであると考えていますが、これを悪用するにはセキュリティ チェーンを接続する必要があるため、通常は報奨金を保証するようなバグではないと考えています。

一緒にバグを修正します。
*Flask デバッグ モード*: web/app.py#L272 はデバッグ モード = true で開始され、デバッガーが接続されます。これは、Flask がデバッガーを侵害するために推測する必要があるランダムな ID で始まるため、「ある程度」保護されていますが、推測/反復してその ID を見つけることができれば、事態をさらにエスカレートさせることができます。ほとんどのモデルでこのバグが発見され、ほとんどの静的ソース コード アナライザーでも発見されると予想されますが、一部の LLM (GPT 5.5 High Reasoning! を含む) は発見後、このバグを無効として拒否しました。
*同時ワーカー*: 理論的には、アプリケーションは水平方向に拡張でき、各「ワーカー」は SQLite キューからプルできます。アプリケーションがそのように設定されている場合、さまざまなワーカーがジョブ キュー内の次のタスクを取得しようとするため、データベースで競合状態が発生する可能性があります。ただし、リポジトリには、これがアプリケーションの意図された動作方法であることを示すものは何もないため、せいぜい境界線上にあります。
*サードパーティ DoS*: アプリケーションは、Yahoo、ESPN などからランキングを取得して、データ モデルを構築します。理論的には、これを悪用して DoS を通じてこれらのサービスを攻撃する可能性があります。実際には、この API を管理者以外が利用できるようにする他のエクスプロイト チェーンがない限り、潜在的な悪用者はアプリの管理者だけです...その時点で、彼らはこれらのサービスを直接 DoS しようとする可能性があります。ただし、これが正当/境界線である理由は、最初のセキュリティ上の問題です。つまり、リフレッシュ キーが公開される可能性があるということです。理論的には、これら 2 つを連鎖させることができます。モデルのいくつか (すべてではありません) が、このチェーンが潜在的な問題の原因であると特定しました。
*レート制限なし/CORS*: これは興味深いものです。最初に潜在的な脆弱性としてフラグを立てたのは MiniMax M3 だけでした。そのアイデアは、誰かが大量の r を送信することでシステムを悪用できるということです。

自動的にクエストを実行します。このシステムは CORS/レート制限を行わないため、悪意のあるサードパーティ サイトによって悪用される可能性もあります。実際には、これは起こりそうにありませんが、実際の潜在的な悪用パターンであり、運用システムで対処する必要があります。
いくつかのモデルで次の脆弱性が「見つかりました」。これらは脆弱性ではないか、実際には悪用できません。
*ブラケット HTML に保存された XSS*: ジェネレーターの出力は通常、印刷可能なブラケットを含む HTML ファイルです。理論的には、Yahoo または ESPN からの不正なデータ ファイルに XSS を生成する任意の HTML/JavaScript が含まれている場合、XSS を使用して印刷可能なブラケットがレンダリングされる可能性があります。実際には、これには Yahoo または ESPN のサービスを侵害する必要があります。これが、実際には無効な発見である理由です。
*ジョブ ID (UUID) 列挙*: 理論的には、すべての UUID を反復処理して「他の人の」括弧を見つけることができます。これらはいかなる種類の認証によっても保護されません。実際には、これは重要ではなく、非現実的です。
*古い依存関係*:requirements.txt ファイルは、requests>=2.28 (および flask>=3.0 および tqdm>=4.65) を固定します。これらの各ライブラリには修正済みの既知の CVE があり、これらの依存関係の最小バージョンを上位のものに変更することで対処できます。これは良い習慣かもしれませんが、この記事の執筆時点では、どの LLM もこれらのライブラリの既知の脆弱性を悪用する実際のパスを特定できませんでした。
*Pickle 逆シリアル化*: アプリケーションは pickle を使用してオブジェクトをシリアル化/逆シリアル化します。理論的には、攻撃者は攻撃のために pickle ファイルを上書きする可能性があります。ただし、これを行うには、アプリケーションを実行しているマシンへのシェル アクセスが必要になります。その時点で、アプリケーションを完全に上書きするだけで、必要な処理を行うことができます。これは脆弱性ではありません

そのためです。幸いなことに、ほぼすべてのモデルがレビューでこれが無効であると判断しました。
*CSV 式の挿入*: 「ブラケット HTML に保存された XSS」と同様に、アプリケーションが Yahoo または ESPN からデータを取得して CSV に書き出すパスがあります。理論的には、Yahoo や ESPN が悪意のあるダンプでチーム名の代わりに Excel の数式をダンプし、その後管理者が CSV を開いた場合、悪意のあるファイルを開いてしまう可能性があります。実際には...いいえ。
*任意のファイル書き込み/パス トラバーサル*: アプリケーションがファイルの読み取り/書き込みを行う場所はさまざまですが、それらをすべてここにグループ化しました。モデルは、サニタイズなしで使用されるパラメータを渡すことができるコマンド ラインからの実行や、ファイルがパラメータとして使用される可能性があるいくつかの URL など、さまざまな項目にフラグを立てました (ただし、最終的には最初に検索によって保護されます)。これらはどれも有効な所見ではありませんでした。
*未検証のリモート フェッチによる SSRF*: リモート データ フェッチャー (ESPN、Yahoo など) は、サードパーティの URL からデータを取得します。理論的には、これらのフェッチを実行するマシンにシェル アクセスできる場合、より悪意のある URL でフェッチを上書きすることができます。ただし、実際には、重要なことを行うにはマシン上で昇格されたアクセス許可が必要になるため、これは、たとえば、セキュリティ上の脆弱性と同じです。弱いパスワード。
*ハードコードされた署名キー*: いくつかの LLM は、「実際の」脆弱な「パブリック リフレッシュ デフォルト キー」と同じ方法で、この秘密キーを脆弱性として特定しました。ただし、実際には認証がないため、このキーがアプリケーションで使用されることはありません。これは、多くの静的分析ツールや、コンテキストを適切に伝達しない一部の LLM で検出される「脆弱性」のもう 1 つの例です。実際には、悪用できるものは何もありません。
*悪意のあるリモート URL での ReDoS*: (ingestion/pick_popularity.py)[https]

://github.com/eskibars/seed-money/blob/main/ingestion/pick_popularity.py#L509] があります。
正規表現 ( match = re.search(r"(\d+(?:\.\d+)?)\s*%", text) ) 理論上は、上流のデータが汚染されて O(n²) 個のバックトラッキングが発生した場合に悪用される可能性があります。ただし、実際には、ここでのデータは信頼できるサードパーティ (Yahoo、ESPN など) を通じて取得されているため、実際的な悪用ベクトルはありません。
これは単一のリポジトリに基づいた逸話ですが、このエクスペリエンスは、ZeroQuarry でこれまでに見てきたものとほぼ一致しています。
「普遍的に良い/悪いサイバーセキュリティ モデル」はありません。代わりに、一部のモデルは「疑わしいものを見つける」ことに非常に優れており、他のモデルは検証/無効化に非常に優れています。基本的に、モデルによっては、段階によっては「深く考えすぎる*」場合もあれば、段階によっては「十分に考えていない*」場合もあります。 MiniMax モデルは、他のモデルよりも「潜在的に疑わしい」ものに *はるかに*頻繁にフラグを立てているようですが (特に価格の点で)、物事を検証するのが *恐ろしく*、*ひどい*修正と PoC を構築します。
スキャンの価格を見ると、時間をかける価値のあるものもあれば、そうでないものもあります。結果はフィルタリングされるため、フロンティアモデルにレビューを行わせることはおそらく価値があるでしょう。

[切り捨てられた]

## Original Extract

A practical comparison of how different LLMs perform on cybersecurity scans, using real ZeroQuarry results from an open source March Madness project.

Why ZeroQuarry
Platform
Continuous Testing
Evidence & Reports
Research
Pricing
Sign in
Request scan ->
Research
Shane Connelly
June 15, 2026
research
cybersecurity
Models and Their Capabilities
A practical comparison of how different LLMs perform on cybersecurity scans, using real ZeroQuarry results from an open source March Madness project.
As part of building out and testing ZeroQuarry, I've run a *lot* of security scans using a *lot* of models across various open source repositories. There are a lot of misconceptions swirling at the time of this writing about the different models and their capabilities with respect to cybersecurity and I wanted to show some *actual* results of using some of these models to identify security issues and how they're all different and better/worse at different *parts* of cybersecurity.
Let me break that down by an example. A while ago, I built a project called Seed Money . It was an ambitious open source project to help me get a leg up on winning March Madness pools. I deployed it locally on my laptop, but I thought it would be a good excuse now to run through a security scan with various models. I've put the results of the scans into this spreadsheet, and I'll explain the structure of this below. I should note here that the app was mostly developed by frontier LLM models themselves, so... no. Your LLM-built apps aren't bug-free. They need reviews just like human coders do.
A Note on the Spreadsheet Structure
There are 2 tabs: the "Findings" tab and the "Time and Cost" tab. These do what's stated on the tin.
In the Findings tab, there are a set of potential vulnerabilities. In security research, we would normally classify a vulnerability as valid/borderline/invalid (or some variation). Because not every "insecure software practice" will yield an exploitable vulnerability. These "potential findings" are color coded in the column A.
For each LLM, I show whether it found the particular item: Yes, No, or "Yes (rejected)." "Yes (rejected)" has to do with the way ZeroQuarry works: ZeroQuarry has an adversarial review loop, where a "researcher" model proposes a vulnerability and then a "vendor" model agrees it is or isn't a vulnerability. "Yes (rejected)" means the researcher model did find the vulnerability but then adversarial review challenged the vulnerability, e.g. because it wasn't reachable in any identifiable way, and the researcher agent then agreed. I mention/show this because it shows the importance of having a multi-agent system: even if you tell an LLM to find/fix vulnerabilities in your codebase, without this adversarial loop, it can lead to a lot more noise/churn.
Then there's a severity score, which is based on CVSS score. The LLM in each case tries to figure out the severity autonomously and it's rated on a 1-10 scale. I've not included any in the spreadsheet which all LLMs determined were informational-only.
Then there's a "PoC Generated" row, which shows whether the model was willing to generate a PoC to try to exploit the vulnerability. You may find it surprising to see just how frequently "frontier LLMs with heavy guardrails" are entirely willing to generate PoCs that exploit software vulnerabilities. A lot of this really comes down to how they are prompted.
Then there's a "Considered H1 Eligible" row. This runs the finding through ZeroQuarry's LLM-based evaluator which shows whether the model would consider it a "valid" vulnerability to submit to HackerOne under typical HackerOne rules. For example, normally vulnerabilities which result in DoS through CPU exhaustion are not considered eligible.
In each of the cases, I've color coded them: green means "it was correct/had no limitations," yellow means "borderline/judgement call" and red means "it was wrong or restrictive."
The "Costs" tab is pretty self-explanatory if you understand the different agent types ZeroQuarry employs. That's explained here so I won't repeat myself in this blog.
Valid/Borderline Vulnerabilities
The following are considered "valid" -- an LLM would be "correct" to find them. The "Found" row indicates if it did or did not.
*Public Refresh Default Key*: web/app.py#L21 bakes a default key in. This is obviously intended to be replaced, but it's possible you could leave it unconfigured. It's a vulnerability. Nearly every model (except Anthropic's "low reasoning" variants) find this vulnerability
*CPU DoS via Simulation Parameters*: in web/app.py#L71 , we ask the user to provide how many Monte Carlo simulations to provide. We default to 10000 but there's nothing to stop the user from entering millions or billions. This could lead to a DoS of the server through CPU exhaustion. Again, nearly every model except for Anthropic's "low reasoning" variants find this vulnerability
*Job Status API Can Leak Tracebacks*: If the Flask server encounters an error, it can dump the trace of how it was called. While there's nothing inherently terrible about this (and thus the low score counted by most models that *did* find it), it can yield additional information to an attacker such as the location on disk of files or SQL commands, for example, which could give them additional insights into attack surface. Most security researchers would consider this a valid security bug worth fixing but not one that would typically warrant a bounty, because in order to exploit it, you'd have to chain several bugs together.
*Flask Debug Mode*: web/app.py#L272 starts in debug mode=true, which attaches a debugger. This is *somewhat* protected in that Flask starts with a random ID you'd have to guess to compromise the debugger, but if you can guess/iterate to find that ID, you can really escalate things. Most models found this bug, and I'd expect even most static source code analyzers would, but some of the LLMs (including GPT 5.5 High Reasoning!) rejected it as invalid after finding it.
*Concurrent Workers*: In theory, the application can scale horizontally, and each "worker" could pull from a SQLite queue. *If* the application is set up that way, there could be race conditions that happen on the database as different workers try to grab the next task in the job queue. There's nothing in the repository that indicates this is how the application is intended to operate though, so it's borderline at best.
*3rd party DoS*: The application pulls rankings from Yahoo, ESPN, and others, to build its data model. In theory, this could be abused to attack those services through a DoS. In practice, unless there's some other exploit chain that renders this API available to non-administrators, the only potential abusers are administrators of the app... at which point, they could just directly attempt to DoS those services. The thing that makes this legitimate/borderline though is the first security issue: that the refresh keys are potentially known public. In theory, these 2 could be chained together. Several (though not all) of the models identified this chain as the source of potential problems.
*No Rate Limiting/CORS*: This one is interesting: only MiniMax M3 flagged it as a potential vulnerability initially. The idea is that someone could abuse the system by submitting a bunch of requests in an automatic way. Because the system doesn't do any CORS/rate limiting, it could even be exploited by a malicious 3rd party site. In reality, this isn't likely to happen, but it is a real potential abuse pattern that would need to be addressed by a production system.
Several models "found" the following vulnerabilities, which are either non-vulnerabilities or cannot practically be exploited.
*Stored XSS in Bracket HTML*: The output of the generator is typically an HTML file containing a printable bracket. In theory, if a bad data file from Yahoo or ESPN contained arbitrary HTML/Javascript that yielded XSS, it could render a printable bracket with XSS. In practice, this would require compromising Yahoo or ESPN's services. This is why it's really an invalid finding.
*Job ID (UUID) Enumeration*: In theory, you can iterate over all UUIDs to find "someone else's" bracket. They aren't protected via any kind of authentication. In practice, this both doesn't matter and is impractical.
*Outdated Dependencies*: The requirements.txt file pins requests>=2.28 (and flask>=3.0 and tqdm>=4.65). Each of these libraries has known CVEs that have been fixed and could be addressed by bumping the minimum version of these dependencies to something higher. While that may be good practice, at the time of this writing, none of the LLMs were able to identify any real path to exploit any of the known vulnerabilities in these libraries.
*Pickle Deserialization*: The application uses pickle to serialize/deserialize objects. In theory, an attacker could overwrite the pickle file for an attack. However, doing so would require shell access to the machine running the application, at which point, you could just overwrite the application entirely to do what you want. This isn't a vulnerability for that reason. Fortunately, nearly every model caught this as invalid on review.
*CSV Formula Injection*: Similar to the "Stored XSS in Bracket HTML" there's a path where the application grabs data from Yahoo or ESPN and writes it out to a CSV. In theory, if Yahoo or ESPN dumped Excel formulas out instead of team names in some malicious dump and then the administrator opened the CSVs, they could end up opening a malicious file. In practice...no.
*Arbitrary File Write/Path Traversal*: There are a variety of places where the application reads/writes files, and I've grouped them all here. The models flagged various items like running from the command line you could pass in parameters that were then used without sanitization and a few URLs where files were potentially used as parameters (but ultimately safeguarded against by lookups first). None of these were valid findings.
*SSRF Via Unvalidated Remote Fetch*: The remote data fetchers (ESPN, Yahoo, etc) grab data from 3rd party URLs. In theory, if you had shell access to the machine running these fetches, you could overwrite those with some URLs that are more malicious. In practice, though, you'd need elevated permissions on the machine to do anything of note, so this is no more of a vulnerability than e.g. weak passwords.
*Hardcoded Signing Key*: Several LLMs identified this secret key as a vulnerability in the same way as the "actual" vulnerable "Public Refresh Default Key." However, in practice, this key is never used in the application because there is no authentication. This is another example of a "vulnerability" which many static analysis tools find, and some LLMs as well that don't carry context well. In practice, there's nothing to exploit.
*ReDoS in Malicious Remote URL*: In (ingestion/pick_popularity.py)[https://github.com/eskibars/seed-money/blob/main/ingestion/pick_popularity.py#L509], there's
a regular expression ( match = re.search(r"(\d+(?:\.\d+)?)\s*%", text) ) which in theory could be exploited if the upstream data is poisoned to yield O(n²) backtracking. However, in reality, the data here is coming through trusted 3rd parties (Yahoo, ESPN, etc) so there's no practical exploit vector.
While this is anecdotal based on a single repository, the experience has been largely consistent with what we've seen so far at ZeroQuarry:
There's no "universally good/bad cybersecurity model." Instead, some models are really good at "spotting things that look fishy" and other models are really good at validating/invalidating. Essentially, some models *think too hard* at some stages and others *don't think hard enough* at some stages. MiniMax models seem to flag "potentially suspicious" things *far* more often (especially for the price) than other models, but they are *terrible* at validating things and build *terrible* fixes and PoCs.
If you look at the scan prices, some things are worth the time, and others are not. It's probably worth having a frontier model do reviews because the findings will be filtered do

[truncated]
