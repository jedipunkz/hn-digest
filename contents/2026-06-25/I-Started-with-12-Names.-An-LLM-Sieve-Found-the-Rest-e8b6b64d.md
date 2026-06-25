---
source: "https://zamechek.com/blog/the-sieve-growing-from-a-seed/"
hn_url: "https://news.ycombinator.com/item?id=48676951"
title: "I Started with 12 Names. An LLM Sieve Found the Rest"
article_title: "I Started With 12 Names. An LLM Sieve Found the Rest."
author: "shawnzam"
captured_at: "2026-06-25T18:43:32Z"
capture_tool: "hn-digest"
hn_id: 48676951
score: 2
comments: 1
posted_at: "2026-06-25T17:51:54Z"
tags:
  - hacker-news
  - translated
---

# I Started with 12 Names. An LLM Sieve Found the Rest

- HN: [48676951](https://news.ycombinator.com/item?id=48676951)
- Source: [zamechek.com](https://zamechek.com/blog/the-sieve-growing-from-a-seed/)
- Score: 2
- Comments: 1
- Posted: 2026-06-25T17:51:54Z

## Translation

タイトル：12の名前から始めました。 LLMふるいが残りを見つけた
記事タイトル: 12の名前から始めました。 LLM ふるいが残りを見つけました。
説明: 私は、信頼できる 12 人の名前を使用してインディーズ レコード メーカーのマップをシードし、LLM ふるい (フィルター、決して作成者ではない) でそれを 469 個の検証可能なノードまで成長させました。 2inch.fmでライブ。

記事本文:
↑トップ
shawn@zamechek :~$ whoami ショーン・ザメチェック
メニュー 仕事について 教育スキル ブログ 参考文献 cd .. エッセイ · 2026年6月25日 私は12の名前から始めました。 LLM ふるいが残りを見つけました。
私が子供の頃から聴いていたインディーズやポストロックのレコードを実際に作ったのは誰なのか、バンドやガラスの向こうの人々ではなく、地図を作りたかったのです。エンジニア、プロデューサー、スタジオ。私はデータベースやスプレッドシートから始めたわけではありません。私は何も考えずに保証したい 12 人の名前から始めました - スティーブ・アルビニ、ジョン・マッキンタイア、ボブ・ウェストン、そして彼らに似たその他の 9 人です。それが種全体です。私が信頼していた十数人が手書きでタイプしていました。他のすべて - 469 ノード、620 アルバム クレジット、すべて 2inch.fm でライブ配信 - は、私が「ふるい」と呼んでいるプロセスを通じて、これら 12 個から増えました。
重要なのは、決して機械に何かを作らせないことだった。
信頼できるものにする唯一のルール
「LLM にデータセットを構築してもらう」プロジェクトはどれも同じように失敗します。つまり、モデルが捏造されてしまいます。それは決して起こらなかったもっともらしい信用をでっち上げ、それを完全な自信を持ってあなたに渡します。そこで私はモデルにジョブを 1 つだけ与え、それをルールにしました。
LLM はフィルターであり、ソースではありません。データは構造化 API から取得され、コードによって解析されます。モデルはジャンクを捨てて、次にどこを探すべきかを提案するだけです。
クレジットを作成しないため、クレジットを発明することはできません。この 1 つの制約が、検証可能なグラフと確信のある幻覚の違いです。
グラフはラウンド単位で成長します。各ラウンドは、私がすでに信頼しているエンジニアから開始され、彼らのクレジットを MusicBrainz からクロールし、単一のノードが追加される前に 3 つのふるいを通過して戻ってきたすべてのものを実行します。
ふるい①は、1995 年から 2003 年までの厳しい年枠であり、私の偏見を認めますが、コードで強制されています。安い、正確、モデルなし。 Sieve ② は、LLM がデータに触れる唯一の場所です。新たに検出されたデータをそれぞれ調べます。

バンドに参加して、はい/いいえの質問に 1 つずつ答えます。これは本物の注目に値する行為ですか、それともノイズですか?これにより、「Various Artists」行、誤って解析されたゴミ、DJ ミックス ボリューム 7 が削除されます。好みやジャンルを判断せず、迷ったときはそのままにしておきます。 Sieve ③ は、インデックスに既に存在するものをすべて削除します。その後、生き残ったものはすべてマージされ、 src:"mb" のタグが付けられ、すべてのエッジに正確な MusicBrainz リリース ID が付けられます。そのため、各ファクトは元のレコードにディープリンクされます。私が手で入力した厳選されたコアにはタグがまったくありません。つまり、私が個人的に保証する部分とマシンが追加した部分は常に分離可能です。 1 行ですべてがトリミングされます。
賢い部分はループです。マージ後、私はモデルに、まだいないエンジニア、つまり同じシーンのハブをさらに指名するように依頼します。生成的な機能はこれだけであり、そこでも事実ではなく検索語を提案しています。これらの名前は次のラウンドのシードとなり、クロールされ、他のものと同様にふるいにかけられます。すでに持っている名前が候補に挙がり始めたら、終わりだとわかります。井戸は枯れてしまいます。シーンは飽和状態です。
種が実際に私にもたらしてくれたもの
機械が真実を告げるということを私がまったく信じなかったにもかかわらず、12 の名前が 469 になりました。シードはクローラーに開始点を与えるだけでなく、システム全体に基準を与えます。私が厳選したすべての名前が実際の協力者のクラスターにドラッグされ、それらの名前はカウントされる前に情報源と照合されました。拡張は、シードが優れているほど誠実であり、それにボルトを付けた出所が検証可能です。
それが私がいつも行う取引です。個人的に支持できる小さなことから始めて、それをふるいに任せて、決して作者ではなく、記憶だけでは書き留めることができないものに成長させます。
プロンプト付きのメソッド

逐語的に公開されており、2inch.fm/about にあります。

## Original Extract

I seeded a map of indie record-makers with 12 names I trusted, then let an LLM sieve (filter, never author) grow it to 469 verifiable nodes. Live at 2inch.fm.

↑ top
shawn@zamechek :~$ whoami Shawn Zamechek
menu about work edu skills blog refs cd .. essay · June 25, 2026 I Started With 12 Names. An LLM Sieve Found the Rest.
I wanted to map who actually made the indie and post-rock records I grew up on — not the bands, the people behind the glass. The engineers, the producers, the studios. I didn’t start with a database or a spreadsheet. I started with twelve names I’d vouch for without thinking — Steve Albini, John McEntire, Bob Weston, and nine more like them. That’s the whole seed. A dozen people I trusted, typed by hand. Everything else — 469 nodes, 620 album credits, all live at 2inch.fm — grew out of those twelve through a process I keep calling the sieve .
The trick was never letting the machine make anything up.
The one rule that makes it trustworthy
Every “let an LLM build my dataset” project fails the same way: the model fabricates. It invents a plausible credit that never happened and hands it to you with total confidence. So I gave the model exactly one job, and made it a rule:
The LLM is a filter, not a source. The data comes from a structured API and is parsed by code. The model only throws out junk and suggests where to look next.
It can’t invent a credit because it never writes one. That single constraint is the difference between a verifiable graph and a confident hallucination.
The graph grows in rounds. Each round starts from engineers I already trust, crawls their credits out of MusicBrainz , and runs everything that comes back through three sieves before a single node is added.
Sieve ① is a hard year window — 1995 to 2003, my admitted bias — enforced in code. Cheap, exact, no model. Sieve ② is the only place the LLM touches the data: it looks at each newly-discovered band and answers one yes/no question — is this a real, notable act, or is it noise? It kills the “Various Artists” rows, the mis-parsed garbage, the DJ-mix-volume-7s. It does not judge taste or genre, and when it’s unsure, it keeps. Sieve ③ drops anything already in the index. Then everything that survives gets merged in, tagged src:"mb" , with the exact MusicBrainz release ID on every edge — so each fact deep-links to the record it came from. The curated core I typed by hand carries no tag at all, which means the part I personally vouch for and the part the machine added are always separable. One line trims all of it.
The clever part is the loop. After merging, I ask the model to nominate more engineers I don’t have yet — hubs from the same scene. That’s the only generative thing it does, and even there it’s proposing search terms , not facts. Those names seed the next round, get crawled, and get sieved like everything else. You know you’re done when the nominations start coming back as names you already have. The well runs dry; the scene is saturated.
What the seed actually bought me
Twelve names became 469 without my ever trusting the machine to tell the truth. The seed didn’t just give the crawler a starting point — it gave the whole system its standards. Every name I hand-picked dragged in a cluster of real collaborators, and every one of those was checked against a source before it counted. The expansion is only as honest as the seed is good, and only as verifiable as the provenance you bolt onto it.
That’s the trade I’d make every time: start from a small thing you can personally stand behind, and let a sieve — never an author — grow it into something you couldn’t have written down from memory.
The method, with the prompts published verbatim, is at 2inch.fm/about .
