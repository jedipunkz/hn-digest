---
source: "https://readablebyai.com/index-report"
hn_url: "https://news.ycombinator.com/item?id=49291660"
title: "A quarter of YC Fall 2025 startups are blank pages to AI crawlers"
article_title: "The ReadableByAI Index — YC Fall 2025 vs. Established SaaS, Raw-HTML Readability | ReadableByAI"
author: "abouchard11"
captured_at: "2026-08-13T21:35:41Z"
capture_tool: "hn-digest"
hn_id: 49291660
score: 2
comments: 0
posted_at: "2026-08-13T20:48:43Z"
tags:
  - hacker-news
  - translated
---

# A quarter of YC Fall 2025 startups are blank pages to AI crawlers

- HN: [49291660](https://news.ycombinator.com/item?id=49291660)
- Source: [readablebyai.com](https://readablebyai.com/index-report)
- Score: 2
- Comments: 0
- Posted: 2026-08-13T20:48:43Z

## Translation

タイトル: YC 2025 年秋のスタートアップの 4 分の 1 は AI クローラーにとって白紙のページ
記事のタイトル: ReadableByAI Index — YC Fall 2025 vs. 確立された SaaS、Raw-HTML の可読性 | AI によって読み取り可能
説明: 145 の YC Fall 2025 ホームページと 455 の確立された SaaS ホームページが、非レンダリング AI クローラーが取得する方法で取得されました。 YC スタートアップの 25.5% がブランク シェルを出荷しているのに対し、SaaS の 2.9% は 8.9 倍の差があります。 2026 年 8 月 8 日に測定、オープンな方法論、6 つの名前付き例。

記事本文:
ReadableByAI インデックス — YC Fall 2025 vs. 確立された SaaS、Raw-HTML の可読性 | ReadableByAI メニュー ▾ 無料スキャン インデックス JavaScript なしでは見えない スタートアップはほぼ 9 倍見えない 鍵のかかったドアのメニュー GPTBot がブロックされる理由 ログ分析 ホストされたログ ドレイン プライバシー規約 修正 THE READABLEBYAI インデックス
YC 2025 年秋のスタートアップ企業の 4 分の 1 は、AI クローラーにとっては白紙の状態です。
145 の YC Fall 2025 ホームページの 25.5 % は、GPTBot、ClaudeBot、および PerplexityBot にクライアントでレンダリングされた空のシェルを出荷しています。見出しや製品説明はなく、 <div id="root"> だけが含まれています。同じ方法で測定した既存の SaaS 企業 455 社のうち、これは 2.9 % です。それは8.9倍の差です。 2026 年 8 月 8 日に測定。
何が測定され、何が測定されなかったのか
このインデックス内のすべてのホームページは 2 回取得されました。1 回目はベースライン ブラウザー ユーザー エージェントで、もう 1 回は 12 個のクローラー ユーザー エージェント (11 個の AI ベンダー クローラーと非 AI コントロールとしての bingbot) ReadableByAI プローブで、すべて同じデータセンター IP から取得されました。以下の可読性の数値 (表示される単語数、レンダリング分類、robots.txt の内容) はベースラインの取得から取得され、これらは ID に依存しません。ページがそのコンテンツを生の HTML で送信するか、クライアント側のレンダリングの背後に非表示にするかは、すべての訪問者、ボット、またはブラウザーに当てはまり、それ以上の確認は必要ありません。
私たちが意図的に公開しなかったのは、ボット固有のアクセス結果、つまりどのユーザー エージェントが 200 ページ、403 ページ、またはチャレンジ ページを取得したかです。 ClaudeBot であると主張するデータセンター IP は未検証のプローブであり、実際のクローラーではありません。ベンダーは公開された IP 範囲によってボットを認証するため、私たちの調査に対する挑戦では、本物のブロックと、単に偽者を疑う WAF を区別することはできません。この問題を解決できるのはサイト独自のサーバー ログだけであり、私たちにはそれがありません。したがって、このインデックスは取得した内容を報告します

コンテンツと許可を証明でき、アクセス権の主張は完全に除外されます。
さらにもう 1 つの制限があり、これはこの方法に対する最も強力な反対意見であるため、明確に述べられています。それは、サイトが異なるリクエスターに異なる HTML を提供できるということです。ユーザー エージェント文字列上の共通の実装キーと、このインデックスがそれを検出します。以下に挙げるすべての企業が、テストされたすべてのクローラ ID に同じコンテンツを返しました。これは、各 ID に返された応答をベースラインと比較することによって検証されました (4 つは、13 回のフェッチすべてでバイトごとに同一でした。他の 2 つは、サイズが 1 パーセント未満で異なり、リクエストごとのタイムスタンプと一致しました)。実質的に異なるページをクローラーに提供した 4 社は、失敗としてカウントされるのではなく、個別に報告されます。外部プローブが除外できないのは、ユーザー エージェントではなく検証済みのクローラー IP 範囲によってコンテンツを変更するサイトです。それはまれであり、このデータセットにはそれを示唆するものは何もありませんが、外部から反証することはできません。これが、サーバーのログが重要であり、プローブだけが最後の決定ではない本当の理由です。
この制限はこのインデックスの欠陥ではありません。それは、私たちのものも含め、外部の探査機が確立できるものの境界です。この問題を解決するには、私たちが持っていない 1 つのレコードが必要です。それは、各ベンダーが公開しているクローラー IP 範囲に対してすべてのヒットをチェックしたターゲット独自のサーバー ログです。これにより、検証済みのクローラーと偽者が区別され、実際に取得したページが表示され、途中で放棄したフェッチが明らかになります。これらはいずれも外部からは見えず、すべてが有料監査で読み取られる内容です。まだ構築されていないため販売されていないロードマップ項目: 住宅ネットワークとデータセンター ネットワークから同時にデュアルオリジン プローブを行うことで、IP に敏感なボット管理がログなしで公開されます。
145/147 YC 2025 年秋 co.

企業および確立された SaaS 企業 514 社のうち 455 社が、ベースライン フェッチに対してクリーンな 200 を返しており、以下にカウントされています。残りは、このページのすべてのパーセンテージから除外されており、黙って削除されません。2 つの YC ドメイン (2 非 200 応答) と 59 の SaaS ドメイン (51 非 200、8 到達不能)。
CSR_SHELL = 生の HTML で表示される単語が 150 未満。 SSR_THIN = 150 ～ 399。 SSR_FULL = 400+。表示ワードの中央値: YC の 630 対 SaaS の 1,239 — 一般的な SaaS ホームページには、一般的な YC の 2025 年秋のホームページの約 2 倍の可読コンテンツが配信されます。ただし、どちらかが到達可能かどうかによって判断されます。
比較は一方向ではありません。確立された SaaS セットに属する 4 社は、ベースラインのフェッチを提供するよりも多くのコンテンツを AI クローラーに提供します。ボットを認識した動的レンダリングで、ブラウザ以外のリクエストを検出し、より薄いマークアップではなくより完全なマークアップで応答します。私たちが測定した 2 つの最大のギャップです。給与プラットフォームは、プレーン ブラウザーのフェッチに提供するクローラーのバイト サイズの約 220 倍、ML-ops プラットフォームは約 141 倍です (どちらも上場企業です。企業ごとの詳細は上記の名前付きサンプル ポリシーに属しているため、ここでは名前を付けていません)。一部のチームはすでにこの問題を意図的に解決しています。これは、これが避けられない問題ではなく、解決可能なエンジニアリング問題であることの証拠です。
145 YC 2025 年秋のホームページ (匿名化)
これらは、匿名化されて公開された、クリーンなベースライン取得を返した 2025 年秋の Y Combinator 企業 145 社です。分布は結果であり、その中の単一の行ではありません。150 の可視単語未満の最近の YC バッチ全体の 25.5 % がこのインデックスの見出しです。 SaaS 名簿が完全に公開されていないのと同じ理由で、ここでは個々の企業の名前は挙げられていません。このインデックスの目的は、名前の壁を築くことではなく、ページを修正することです。あなたの会社がこのエリアにある場合

そうですね、ReadableByAI ホームページの無料スキャンは、私たちがこのホームページをチェックしたのと同じ方法であなた自身のホームページをチェックするので、私たちの連絡を待たずに自分の立ち位置を知ることができます。
yc-001 ... yc-145 ID はこのデータセット内で安定しています。同じ会社がリリースごとに同じ ID を維持しているため、再検証された行を複数のウェーブで追跡できます。しかし、これらの ID には私たちが公開するマッピングは含まれていません。このサイトのどこにも、ID をドメインに戻す検索テーブルはありません。
このインデックスは、測定した 600 ドメインの名簿を公開していません。その理由については、以下を参照してください。そのトレードオフなしに公開できるのは、読者が 10 秒以内に再現できる結果です。6 つの大規模でリソースが豊富で、公的に認知されている企業が、このデータセット内の他のすべてのホームページをチェックしたのとまったく同じ方法でチェックしました。
これらは無名なチームやリソースが不足しているチームではありません。それが重要です。 Palantir や Duolingo のようなエンジニアリングの人員を擁する企業は、サーバーサイド レンダリングを行う余裕がないため、AI クローラーにゼロワードを提供しているわけではありません。一般的なフロントエンド フレームワークのデフォルト設定では、誰かが意図的にサーバー レンダリングをオンにしない限り、クライアント レンダリングのシェルが同梱されているため、ワードがゼロになります。これはフレームワークのデフォルトの問題であり、能力の問題ではありません。6 つのいずれも robots.txt 内の AI クローラーを禁止していません。これは、これが事故ではなくポリシーである場合に予想されることです。 GPTBot に対して意図的に立ち入りを禁止されているページには、robots.txt でその旨が記載されています。単に空のページは何も言いません。誰もそれが空であることを意図していないからです。
これらのいずれかを自分で再現します。生の HTML、JavaScript が無効、表示される単語数のみです。
カール -s https://duolingo.com | python3 -c "import re,sys; h=sys.stdin.read(); h=re.sub(r'<(script|style)\ [^>]*>.*?</\1>','',h,flags=re.S|re.I)

; h=re.sub(r'<[^>]+>',' ',h); print(len(h.split()))" 公開しない内容とその理由
このページの以前のバージョンでは、測定された 600 個のドメインすべてがリストされていました。私たちはそのテーブルを下ろしました。このデータセットは、非公開のアウトリーチ リスト (AI クローラーにホームページが見えないことに気付いた企業に直接連絡する企業) としても機能します。完全な名簿を公開すると、問題を解決するという企業のインセンティブと、より大きなリストを公開するという私たちのインセンティブが引き換えになります。むしろ修正してもらいたいです。上記の SaaS 企業 6 社以外にも、このデータセット内の他の 16 社でも同じ結果が得られました。私たちは、名前の壁を公開するのではなく、それぞれに非公開で通知しています。これは隠蔽ではなく、はっきりと述べた実際のポリシーです。このインデックスの目的はページを修正することであり、誰が捕まったのかを示すリーダーボードではありません。あなたが会社を経営していて、自分の立場を知りたい場合は、ReadableByAI ホームページの無料スキャンで、私たちの連絡を待たずに、今すぐに同じ方法であなたのサイトをチェックしてください。
すべてのホームページは、ヘッドレス ブラウザや JavaScript の実行を行わず、プレーンな HTTP GET で取得されました。これは、GPTBot、ClaudeBot、および PerplexityBot が行うことだからです。表示される単語数は、生の HTML 応答から抽出されたテキスト コンテンツから、スクリプトとスタイルのコンテンツを差し引いたものです。 3 つの分類バンドは固定しきい値です: CSR_SHELL は 150 可視ワード未満、SSR_THIN 150 ～ 399、SSR_FULL 400 以上。 robots.txt は直接読み取られ、文字通りの事実として報告されます。存在するかどうか、意図については何も推測されません。このインデックスの背後にあるエンジンは、 github.com/abouchard11/geo-crawl-audit でオープンソースです。誰でもそれを複製し、自分のリストに指定して、私たちの数字を確認したり、独自のリストを作成したりできます。
修正しましたか？無料で再確認させていただきます。
ホームページを測定したときに CSR_SHELL または SSR_THIN が返された場合、上記の名前または

非公開で通知されます。その後、サーバーでレンダリングされたコンテンツを出荷した場合は、私たちに知らせてください。無料で再調査します。キャッチもアップセルも添付されません。このインデックスの目標は、正確な公的記録であり、誰もが固執し続けるリーダーボードではありません。ご連絡ください: alex+reverify@midnightdev.dev 。
ここで名前出た？直接ご相談ください。
あなたの会社がこのレポートに記載されている場合（上の表に名前が記載されているか、匿名で記載されている場合）、これはあなたの直通ラインです: alex+named@midnightdev.dev 。同じ週に、実在の人物が回答します。公開前に貴社の有効な公開連絡チャネルがなかった場合、このアドレスがそのチャネルになります。フォームもゲートキーパーもありません。紛争と複製の不一致は、修正ページのプロセスに従います。修正後の再検証はいつでも無料です。
4 つのケース ファイルでは、このインデックスの背後にある調査結果がさらに詳しく説明されており、ReadableByAI ホームページの無料スキャンでは、同じ方法で自分のサイトをチェックできます。
クライアントがレンダリングしたシェルは、スクリプトを実行しないクローラーにとってどのように見えるか、ブラウザーが描画したものと並べた生の HTML です。
スタートアップ企業は 9 倍近く目に見えない
この指数の背後にある YC Fall 2025 と確立された SaaS の比較の全容 (レビューに残らなかった 2 つの調査結果を含む)。
なぜ llms.txt ファイルを公開しても、その背後にあるホームページが空のシェル (許可するコンテンツのない許可) である場合、何も行われないのですか。
「GPTBot をブロックする理由」は通常間違っています
robots.txt の禁止、WAF チャレンジ、空のページの違いと、このデータセットに 3 つのうち 1 つだけが存在する理由。
この指数は一度ではなく定期的に測定されます。次のウェーブでは、何が変わったのか、どのドメインが CSR_SHELL ホームページを修正したか、どのドメインが修正しなかったのか、そして YC Fall 2025 と確立された SaaS の間のギャップが縮まったのか広がったのかが示されます。

## Original Extract

145 YC Fall 2025 homepages and 455 established SaaS homepages, fetched the way a non-rendering AI crawler fetches them. 25.5% of YC startups ship a blank shell vs. 2.9% of SaaS — a 8.9x gap. Measured 8 August 2026, open methodology, six named examples.

The ReadableByAI Index — YC Fall 2025 vs. Established SaaS, Raw-HTML Readability | ReadableByAI Menu ▾ Free scan The Index Invisible without JavaScript Startups nearly 9x more invisible A menu on a locked door Why they block GPTBot Log analysis Hosted log drain Privacy Terms Corrections THE READABLEBYAI INDEX
A quarter of YC Fall 2025 startups are blank pages to AI crawlers.
25.5 % of 145 YC Fall 2025 homepages ship an empty client-rendered shell to GPTBot, ClaudeBot and PerplexityBot — no headline, no product description, nothing but a <div id="root"> . Among 455 established SaaS companies measured the same way, it's 2.9 %. That's a 8.9 x gap. Measured 8 August 2026 .
What was measured, and what wasn't
Every homepage in this Index was fetched twice: once with a baseline browser user-agent, and once each with the twelve crawler user-agents (eleven AI-vendor crawlers plus bingbot as a non-AI control) ReadableByAI probes, all from the same datacenter IP. The readability numbers below — visible word count, rendering classification, robots.txt contents — come from the baseline fetch, and they are identity-independent : whether a page ships its content in raw HTML or hides it behind a client-side render is true for every visitor, bot or browser, and needs no further confirmation.
What we deliberately did not publish is bot-specific access results — which user-agents got a 200, a 403, or a challenge page. A datacenter IP claiming to be ClaudeBot is an unverified probe, not the real crawler; vendors authenticate their bots by published IP range, so a challenge to our probe cannot distinguish a genuine block from a WAF simply doubting an impostor. Only the site's own server logs settle that, and we don't have them. So this Index reports what any fetch can prove — content and permission — and leaves access claims out entirely.
One further limit, stated plainly because it is the strongest objection to this method: a site can serve different HTML to different requesters. The common implementation keys on the user-agent string, and this Index detects it — every company named below returned the same content to every crawler identity tested , verified by comparing the response returned to each identity against the baseline (four were byte-for-byte identical across all thirteen fetches; the other two varied by under one percent in size, consistent with per-request timestamps). The four companies that did serve crawlers a materially different page are reported separately rather than counted as failures. What an outside probe cannot rule out is a site that varies its content by verified crawler IP range rather than by user-agent. That is rare, and nothing in this dataset suggests it, but it cannot be disproven from outside — which is the honest reason server logs matter and probes alone are never the last word.
That limit is not a flaw in this Index; it is the boundary of what any outside probe can establish, ours included. Resolving it needs the one record we do not have: the target's own server logs, with every hit checked against each vendor's published crawler IP ranges. That separates verified crawlers from impostors, shows which pages they actually retrieved, and surfaces the fetches they abandoned midway — none of which is visible from the outside, and all of which is what the paid audit reads. A roadmap item, not yet built and therefore not sold: dual-origin probing from residential and datacenter networks simultaneously, which would expose IP-sensitive bot management without logs.
145 of 147 YC Fall 2025 companies and 455 of 514 established SaaS companies returned a clean 200 to the baseline fetch and are counted below. The rest are excluded from every percentage on this page, not silently dropped: 2 YC domain s ( 2 non-200 response s ) and 59 SaaS domains ( 51 non-200, 8 unreachable).
CSR_SHELL = under 150 visible words in raw HTML. SSR_THIN = 150–399. SSR_FULL = 400+. Median visible words: 630 for YC vs. 1239 for SaaS — the typical SaaS homepage ships roughly double the readable content of the typical YC Fall 2025 homepage, before either one is judged by whether it's reachable at all.
The comparison isn't one-directional. 4 companies in the established-SaaS set serve AI crawlers more content than they served our baseline fetch — bot-aware, dynamic rendering that detects a non-browser request and responds with fuller markup instead of a thinner one. The two largest gaps we measured: a payroll platform serving roughly 220x the byte size to a crawler that it serves to a plain browser fetch, and an ML-ops platform at roughly 141x (both public companies; unnamed here because per-company detail belongs to the named-examples policy above). Some teams have already solved this deliberately — it's evidence this is a solvable engineering problem, not an unavoidable one.
The 145 YC Fall 2025 homepages, anonymized
These are the 145 Y Combinator Fall 2025 companies that returned a clean baseline fetch, published anonymized. The distribution is the finding, not any single row in it: 25.5 % of a whole recent YC batch under 150 visible words is the headline of this Index. Individual companies aren't named here, for the same reason the SaaS roster isn't published in full — the point of this Index is to get pages fixed, not to build a wall of names. If your company is in this batch, the free scan on the ReadableByAI homepage checks your own homepage the same way we checked this one, so you can find out where you stand without waiting on us to reach out.
The yc-001 … yc-145 ids are stable within this dataset — the same company keeps the same id release over release, so a re-verified row can be tracked across waves — but they carry no mapping we publish. There is no lookup table connecting an id back to a domain anywhere on this site.
This Index doesn't publish the roster of 600 domains it measured — see below for why. What it can publish, without that tradeoff, are results any reader can reproduce in ten seconds: six large, well-resourced, publicly recognizable companies, each checked exactly the way every other homepage in this dataset was checked.
These aren't obscure or under-resourced teams — that's the point. A company with Palantir's or Duolingo's engineering headcount isn't serving zero words to AI crawlers because it can't afford server-side rendering. It's serving zero words because a popular frontend framework's default configuration ships a client-rendered shell unless someone deliberately turns on server rendering, and nobody happened to. This is a framework-default problem, not a competence problem — and none of the six disallows any AI crawler in robots.txt, which is what you'd expect to see if this were policy instead of an accident. A page that's deliberately kept off-limits to GPTBot says so in robots.txt; a page that's just empty says nothing, because nobody meant for it to be empty.
Reproduce any of these yourself — raw HTML, JavaScript disabled, visible word count only:
curl -s https://duolingo.com | python3 -c "import re,sys; h=sys.stdin.read(); h=re.sub(r'<(script|style)\[^>]*>.*?</\1>','',h,flags=re.S|re.I); h=re.sub(r'<[^>]+>',' ',h); print(len(h.split()))" What we're not publishing, and why
Earlier versions of this page listed all 600 measured domains. We took that table down. This dataset doubles as a private outreach list — companies we contact directly when we find their homepage invisible to AI crawlers — and publishing the full roster would trade a company's incentive to fix the problem for our incentive to publish a bigger list. We'd rather have the fix. Beyond the six SaaS companies named above: sixteen other companies in this dataset had the same finding. We are notifying each of them privately rather than publishing a wall of names. That's not concealment — it's the actual policy, stated plainly: the point of this Index is to get pages fixed, not a leaderboard of who's been caught. If you run a company and want to know where you stand, the free scan on the ReadableByAI homepage checks your own site the same way, right now, without waiting for us to reach out.
Every homepage was fetched with a plain HTTP GET — no headless browser, no JavaScript execution — because that's what GPTBot, ClaudeBot and PerplexityBot do. Visible word count is text content extracted from the raw HTML response, minus script and style contents. The three classification bands are fixed thresholds: CSR_SHELL under 150 visible words, SSR_THIN 150–399, SSR_FULL 400 or more. robots.txt is read directly and reported as a literal fact — present or absent, nothing inferred about intent. The engine behind this Index is open source at github.com/abouchard11/geo-crawl-audit . Anyone can clone it, point it at their own list, and check our numbers or produce their own.
Fixed it? We'll re-verify free.
If your homepage came back CSR_SHELL or SSR_THIN when we measured it — named above or notified privately — and you've since shipped server-rendered content, tell us and we'll re-probe it at no charge — no catch, no upsell attached. The goal of this Index is an accurate public record, not a leaderboard anyone stays stuck on. Tell us: alex+reverify@midnightdev.dev .
Named here? Talk to us directly.
If your company appears in this report — named in the table above or described anonymously — this is your direct line: alex+named@midnightdev.dev . A real person answers, same week. If we didn't have a working public contact channel for your company before publication, this address is that channel — no form, no gatekeeper. Disputes and reproduction mismatches follow the process on the Corrections page; re-verification after a fix is free, always.
Four case files go deeper on the findings behind this Index, and the free scan on the ReadableByAI homepage checks your own site the same way.
What a client-rendered shell looks like to a crawler that never runs the script — the raw HTML, side by side with what a browser paints.
Startups are nearly 9x more invisible
The YC Fall 2025 vs. established-SaaS comparison behind this Index, in full — including the two findings that didn't survive review.
Why publishing an llms.txt file does nothing if the homepage behind it is a blank shell — permission without content to permit.
“Why they block GPTBot” is usually wrong
The difference between a robots.txt disallow, a WAF challenge, and an empty page — and why only one of the three is in this dataset.
This Index is measured periodically, not once. The next wave will show what changed — which domains fixed a CSR_SHELL homepage, which didn't, and whether the gap between YC Fall 2025 and established SaaS narrowed or widened.
