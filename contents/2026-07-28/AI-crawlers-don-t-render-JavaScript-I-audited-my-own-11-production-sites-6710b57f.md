---
source: "https://www.smithappstudio.com/blog/ai-crawlers-cant-see-your-spa-2026/"
hn_url: "https://news.ycombinator.com/item?id=49079477"
title: "AI crawlers don't render JavaScript – I audited my own 11 production sites"
article_title: "Your site welcomes AI crawlers. Mine was serving them six words. — Smith App Studio"
author: "docmaasi"
captured_at: "2026-07-28T04:55:06Z"
capture_tool: "hn-digest"
hn_id: 49079477
score: 1
comments: 0
posted_at: "2026-07-28T04:48:46Z"
tags:
  - hacker-news
  - translated
---

# AI crawlers don't render JavaScript – I audited my own 11 production sites

- HN: [49079477](https://news.ycombinator.com/item?id=49079477)
- Source: [www.smithappstudio.com](https://www.smithappstudio.com/blog/ai-crawlers-cant-see-your-spa-2026/)
- Score: 1
- Comments: 0
- Posted: 2026-07-28T04:48:46Z

## Translation

タイトル: AI クローラーは JavaScript をレンダリングしない – 私は自分自身の 11 の運用サイトを監査しました
記事のタイトル: あなたのサイトは AI クローラーを歓迎しています。私は彼らに6つの言葉を提供していました。 — スミス アプリ スタジオ
説明: GPTBot、ClaudeBot、PerplexityBot は使用できません

記事本文:
あなたのサイトでは AI クローラーを歓迎しています。私は彼らに6つの言葉を提供していました。 — スミス アプリ スタジオ
共有方法…
リンクをコピー
Xで共有する
LinkedIn で共有する
電子メールで共有
日記
アプリを詳しく見る →
日記より · 2026 年 7 月 27 日
あなたのサイトでは AI クローラーを歓迎しています。私は彼らに6つの言葉を提供していました。
Google のクローラーは JavaScript を実行します。 ChatGPT にはありません。私は、私自身の 11 の生産現場でその差がどれくらいかかるかを測定しましたが、そのうちの 3 つは目に見えませんでした。
Maasi J. Smith 博士著 · 9 分で読みました · 2026 年 7 月 27 日発行
AI クローラーは JavaScript を実行しません。 GPTBot、ClaudeBot、および PerplexityBot は、生の HTML を一度読み取ってから次に進みます。 Googlebot は JavaScript をレンダリングします。そうではありません。
したがって、クライアントがレンダリングした単一ページのアプリはクライアントには見えません。ランクが低くても見えないわけではありません。
私自身の生産現場 11 か所を測定しました。 3 つは 10 単語未満の AI クローラーにサービスを提供しました。 3 つすべてに、それらのクローラーを明示的に招待する完璧な robots.txt ファイルがありました。
ビルド時のプリレンダリングで 3 つすべてを修正しました: 5 → 1,442、6 → 3,048、8 → 3,021 ワード。レンダラーには午後かかりました。もっとひどいものを出荷するのを阻止する警備員の作業にはさらに時間がかかり、全員が本物のバグを発見しました。
どのサイトも約4秒でチェックできます。スクリプトは以下にあります。ここで私の言葉を鵜呑みにするよりも、自分のサイトで実行してもらいたいと思います。
内容に問題があるのかと思った
先週、私はChatGPTに対し、法廷でのコミュニケーション記録を必要とする別居中の親向けにアプリを推奨するよう依頼した。まさにその製品を作りました。何ヶ月もライブを続けてきました。思い浮かびませんでした。
内容に問題があるのではないかと思いました。ブログ投稿が足りない、バックリンクが足りない、これはいつものことです。
配管に問題がありました。私のサーバーは AI クローラーを空の部屋に送信していました。
現在、インターネット上には 2 種類のクローラーがあります。

彼らは全く異なる振る舞いをします。
Googlebot はヘッドレス Chrome を実行します。ページを取得し、JavaScript を実行し、フレームワークが DOM を構築するのを待ち、結果にインデックスを付けます。これが、シングルページ アプリが Google で 10 年間存続し続けてきた理由です。
AI クローラー (GPTBot、OAI-SearchBot、ClaudeBot、PerplexityBot、CCBot) はサポートしません。彼らは 1 つの HTTP リクエストを発行し、返される HTML を解析して終了します。レンダリングはありません。 2 回目の試行はありません。待つ必要はありません。
Google のクローラーは JavaScript を実行します。 ChatGPT にはありません。そのたった 1 つの違いによって、AI アシスタントが製品を推奨できるかどうかが決まります。
これは推測ではありません。 5 億件を超える GPTBot フェッチを分析した結果、JavaScript が実行された証拠はまったく見つかりませんでした。クローラーログの調査によると、GPTBot はリクエストの約 11.5% で JavaScript ファイルをダウンロードし、ClaudeBot は約 23.8% で JavaScript ファイルをダウンロードしますが、どちらもその実行は観察されていません。彼らはファイルをプルします。彼らはそれを開きません。
したがって、マーケティング ページが <div id="root"></div> で起動する React または Vue アプリの場合、AI アシスタントがあなたの会社について知っていること、つまり <title> タグとメタ ディスクリプションの完全かつ網羅的なリストがここにあります。
それだけです。それがコーパス全体です。
私は 8 つの消費者向け製品、B2B プラットフォーム、Web3 マーケットプレイス、スタジオ サイトを運営しています。私は4行のスクリプトを書き、それを11すべてに向けて、コーヒーを淹れに行きました。私はこれに戻りました：
3つの製品。 5語、6語、8語。
さて、ここが刺さった部分です。私はそこに問題があることを期待して、これら 3 つのサイトの robots.txt ファイルを調べました。代わりに、次のようなものを見つけました。
# AI アシスタントと応答エンジン — 発見しやすさのために明示的に歓迎されています
ユーザーエージェント: GPTBot
ユーザーエージェント: OAI-SearchBot
ユーザーエージェント: ChatGPT ユーザー
ユーザーエージェント: ClaudeBot
ユーザーエージェント: PerplexityBot
ユーザーエージェント: Google-Extended
ユーザーエージェント:

CCボット
許可: /
汚れのない。 3 つはすべて、有効な llms.txt も提供します。 3 つすべてに、 <head> 内に正規タグ、Open Graph イメージ、Twitter カード、および JSON-LD 構造化データがあります。
私はレッドカーペットを敷き、7か国語で歓迎のサインを印刷し、それをすべて空の部屋に向けました。
良いスコアを獲得した 8 つのサイトは、マーケティングが優れているわけではありません。これらはサーバー上でレンダリングされるフレームワーク上に構築されているだけです。それがすべての違いです。ある製品には 2 年間の取り組みがあり、その発見の可能性は、私が午後に下したビルド時の決定にかかっています。
なぜこれが昨年よりも今年より重要なのか
2024 年にはこれを無視できたはずですが、現在は無視できません。3 つの理由はいずれも測定可能です。
1 つ目は、ランキングと引用されることは切り離されていることです。 Seeer Interactive は、ChatGPT、Perplexity、AI 概要全体で 5,000 以上の URL を分析し、上位の Google 検索結果と AI が引用したソースとの重複が約 70% から 20% 未満に減少したことを発見しました。 Google ランクは AI の可視性を表すものではなくなりました。現在、これらは 2 つの別々の販売チャネルとなり、2 つの別々の仕組みを備えています。
2: AI アシスタントは圧倒的にブランド自身のサイトを引用します。 Yext は、ChatGPT、Gemini、Perplexity 全体で 680 万件の引用を分析し、86% がブランドが管理するソースからのものであることを発見しました。あなたの製品に関して最も多く引用されている唯一の情報源はあなたであるはずです。 6 ワードを提供している場合、デフォルトではその表面積を失ったことになります。
3 つ目: ここで製品の発見が進んでいます。人々はもはや「最高の共同子育てアプリ」を検索して 10 個の青いリンクを比較することはありません。彼らはアシスタントに依頼し、候補者リストを選びます。最終候補リストに含まれていない場合は、検討対象には含まれていないため、2 ページ目に進む必要はありません。
llms.txt が行うことと行わないこと
自信満々のナンセンスが多いので注意したい

これについて。
llms.txt は、提案されている規則です。サイトの言語モデルを要約した、ドメイン ルートにあるプレーンテキスト ファイルです。安価で、静的で、サーバーで提供されるため、1 つ持っておくべきだと思います。
しかし、それはこの問題の解決策ではありません。これは自主的な規約であり、一貫性のない採用があり、クローラーのインデックス内の実際のページを置き換えるものではありません。そして、ここが重要ですが、私のものはすでにそこにありました。 3 つの非表示のサイトはすべて、完全に良好な llms.txt を提供していましたが、保存されませんでした。110 ワードの概要はサイトの代わりにはならないからです。
llms.txt を名刺として扱います。役に立つ。建物ではありません。
自分のサイトを 4 秒でチェック
これがスクリプトです。 URL を GPTBot として取得し、スクリプト、スタイル、コメント、タグを取り除き、実際に残っているもの (クローラーが実際に読み取るもの) をカウントします。
#!/usr/bin/env bash
# ai-visibility.sh — AI クローラーは実際にいくつの単語を認識しますか?
「$@」の URL の場合;する
Words=$(curl -sL --max-time 20 \
-A "Mozilla/5.0 (互換性; GPTBot/1.0; +https://openai.com/gptbot)" "$url" \
| perl -0777 -pe の{<script\b.*?</script>}{}gsi;
s{<style\b.*?</style>}{}gsi;
s{<noscript\b.*?</noscript>}{}gsi;
s{<!--.*?-->}{}gs;
s{<[^>]+>}{ }gs' \
| tr -s " \t\r\n" " " |トイレ -w)
if [ "$words" -lt 200 ];その後、評決 = "INVISIBLE"
elif [ "$words" -lt 600 ];その後、評決 = "THIN"
それ以外の場合の判定 = "OK"
フィ
printf "%-10s %6s 単語 %s\n" "$verdict" "$words" "$url"
完了しました
実行してください:
chmod +x ai-visibility.sh
./ai-visibility.sh https://yoursite.com/ https://yoursite.com/pricing
結果を読む。これらのしきい値は判断基準であり、標準ではありませんが、私がテストしたすべての項目で維持されています。
正直な注意点が 2 つあります。まず、単語数は内容を表すものであり、尺度ではありません。2,000 単語のナビゲーション リンクと法的な定型文は n です。

価値のある 2,000 語。次に、Perl は macOS と Linux に同梱されています。 Windows では Git Bash または WSL を使用します。
努力によるランキングです。一番上から始めて、番号が移動したら停止します。
マーケティングルートを事前にレンダリングします。 Vite または CRA アプリの場合、これは通常半日で完了する変更であり、利用可能な修正の中で最も効果が高いものです。ビルド後のステップでは、ヘッドレス ブラウザでルートをクロールし、実際の HTML ファイルをディスクに書き込みます。アプリは現在とまったく同じように動作し続けます。サーバーは、JavaScript が起動する前に何かを渡す必要があるだけです。
マーケティング サイトをサーバー上でレンダリングされるフレームワークに移行します。 Next.js、Astro、リミックス、Nuxt。とにかく再構築する場合、これは永続的な答えです。これは公開ページの場合にのみ必要であることに注意してください。ログインの背後にあるものはすべて、クライアントによってレンダリングされたままになるため、クローラーがそれを参照する必要はありません。
マーケティング サイトをアプリから分割します。過小評価されている。 yourdomain.com 上の静的サイトと app.yourdomain.com 上のアプリにより、完璧なクローラビリティとより高速なマーケティング サイトが得られ、ルーティングの問題のカテゴリー全体が解消されます。
動的レンダリングには手を出さないでください。ユーザー エージェントに基づいて人間とは異なる HTML をボットに提供することが、かつては標準的なアドバイスでした。これは脆弱で、クローキングのリスクがあり、プリレンダリングを使用すると、ユーザー エージェントの検出を信頼するかどうかを誰にも尋ねることなく、同じ問題が解決されます。
だったら安いものをやれよ。すべてのサイトマップを Bing ウェブマスター ツールに送信します。ChatGPT の検索は Bing のインデックスに基づいており、これは 20 分間の作業で、ほとんど誰も行いません。 Article および FAQPage スキーマを追加します。本物の llms.txt をアップします。見出しを質問として書き、最初の文で答えてください。それが AI システムが実際に抽出する単位だからです。
3 つのサイトすべてのプリレンダラーを構築しました。デルタは次のとおりです。
現在、約 270 ページが実際の HTML として出荷されています。ゼロファル

l 200 ワードのしきい値を下回っています。
私が予想していなかった 3 つのことは、すべて実際の教訓です。
素朴なバージョンは静かに状況を悪化させました
私の最初の作業実行では、ホームページのタイトルを含む 54 ページを作成しました。これは、ルートごとのメタデータが適用されるまで一定の 500 ミリ秒待ったのですが、4 つのヘッドレス タブを同時に実行している場合、適用されない場合がありました。 1 つのタイトルの下にある 54 個の URL は重複コンテンツです。タイトルを間違えただけで分かりました。
プリレンダリングにより正規タグのバグがゼロから作成されました
HTML シェルは <link rel="canonical" href="/"> をハードコードします。クローラーが 1 つのファイルだけをフェッチする場合は無害です。すべてのルートが独自のファイルになった瞬間、すべてのページに 2 つの正規版が送信されます。1 つは正直に自分自身を指しており、もう 1 つはホームページであると主張しています。これは、プリレンダリングをまったく行わないよりも悪く、私たちのブログの 1 つは、以前にまさにそのバグによってすでにインデックスが解除されていました。
ビルドにより、誰も気づかなかったページが壊れていることが判明しました
PostPilot の /platforms ページは、独自のタイトル、説明、または正規を設定したことはなく、常に一般的なホームページのメタデータを提供していました。他の公開ページにはすべてそれがありました。シングルページのアプリでは、違いを探さなければ違いが分からないため、これまでそのことは何も表面化していませんでした。
正直にまとめると、プリレンダリングは簡単な部分でした。警備員がその仕事だった。これを行う場合は、レンダラーを作成する前に、ビルドを失敗させるチェックを作成してください。私のすべてのチェックは、実際のものを捉えています。
最初の実稼働デプロイのうち 2 つも完全に失敗しました。これは、ビルド イメージに Chrome が含まれておらず、それを実行するためのシステム ライブラリもなかったためです。どちらの失敗も誰にも影響しませんでした。スクリプトは問題が発生するとゼロ以外で終了するため、以前に動作していたデプロイは単にライブのままでした。これが、警告ではなく激しく失敗するビルドタイム ガードの主張です。
いいえ

。 GPTBot、ClaudeBot、および PerplexityBot は生の HTML をフェッチし、JavaScript を実行しません。 5 億件を超える GPTBot フェッチを分析した結果、JavaScript が実行された形跡は見つかりませんでした。 Googlebot は例外で、ヘッドレス Chrome エンジンを使用して JavaScript をレンダリングします。
サーバーが送信する HTML 内に存在する部分のみ。 React アプリが完全にブラウザー内でレンダリングされる場合、ChatGPT は <title> とメタ ディスクリプションのみを参照します。
いいえ、robots.txt はクロールの権限を付与します。クロールするコンテンツは作成されません。サイトは、存在するすべての AI クローラーを歓迎し、空のページを提供することができます。
いいえ。これは便利で安価な追加機能 (言語モデルの概要ファイル) ですが、採用には一貫性がなく、サーバーでレンダリングされたページの代わりにはなりません。私の非表示の 3 つのサイトにはすべて、すでに 1 つのサイトがありました。
GPTBot ユーザー エージェントを使用して、curl でそれを取得し、タグとスクリプトを削除した後、応答内の単語をカウントします。マーケティング ページの単語数が 200 語未満であるということは、実質的に目に見えないことを意味します。スクリプトは上記です。
思ったよりも少ない — Googlebot は JavaScript をレンダリングするため、よく構築された単一ページのアプリは上位にランクされる可能性があります。これがまさにこれを危険にしている理由です。アナリティクスは健全に見えますが、新興チャネル全体が黙って何も返しません。
AI クローラーは HTML を読み取ります。 JavaScript は実行されません。もし

[切り捨てられた]

## Original Extract

GPTBot, ClaudeBot and PerplexityBot don

Your site welcomes AI crawlers. Mine was serving them six words. — Smith App Studio
Share via…
Copy link
Share on X
Share on LinkedIn
Share via Email
Journal
Explore the apps →
From the Journal · July 27, 2026
Your site welcomes AI crawlers. Mine was serving them six words.
Google’s crawler runs JavaScript. ChatGPT’s does not. I measured what that difference costs across eleven of my own production sites — and three of them were invisible.
By Dr. Maasi J. Smith · 9 min read · Published July 27, 2026
AI crawlers do not execute JavaScript. GPTBot, ClaudeBot and PerplexityBot read your raw HTML once and move on. Googlebot renders JavaScript. They do not.
A client-rendered single-page app is therefore invisible to them — not badly ranked, invisible.
I measured eleven of my own production sites. Three served AI crawlers fewer than ten words. All three had perfect robots.txt files explicitly inviting those crawlers in.
I fixed all three with build-time prerendering: 5 → 1,442, 6 → 3,048, 8 → 3,021 words. The renderer took an afternoon. The guards that stop it shipping something worse took longer — and every one of them caught a real bug.
You can check any site in about four seconds. The script is below, and I would rather you ran it on your own site than took my word for anything here.
I thought I had a content problem
Last week I asked ChatGPT to recommend an app for separated parents who need court-ready records of their communication. I have built exactly that product. It has been live for months. It did not come up.
I assumed I had a content problem. Not enough blog posts, not enough backlinks, the usual.
I had a plumbing problem. My server was sending AI crawlers an empty room.
There are two kinds of crawler on the internet now, and they behave completely differently.
Googlebot runs a headless Chrome. It fetches your page, executes the JavaScript, waits for the framework to build the DOM, and indexes the result. This is why single-page apps have been survivable in Google for a decade.
AI crawlers — GPTBot, OAI-SearchBot, ClaudeBot, PerplexityBot, CCBot — do not. They issue one HTTP request, parse whatever HTML comes back, and leave. No rendering. No second attempt. No waiting.
Google’s crawler runs JavaScript. ChatGPT’s does not. That single difference decides whether an AI assistant can recommend your product.
This is not speculation. An analysis of over 500 million GPTBot fetches found zero evidence of JavaScript execution. Crawler-log studies show GPTBot downloads JavaScript files in roughly 11.5% of requests and ClaudeBot in about 23.8% — and neither has ever been observed running them. They pull the file. They do not open it.
So if your marketing page is a React or Vue app that boots into <div id="root"></div> , here is the complete, exhaustive list of what an AI assistant knows about your company: your <title> tag and your meta description.
That is it. That is the whole corpus.
I run eight consumer products, a B2B platform, a Web3 marketplace and a studio site. I wrote a four-line script, pointed it at all eleven, and went to make coffee. I came back to this:
Three products. Five, six, and eight words.
Now here is the part that stung. I went and looked at those three sites’ robots.txt files, expecting to find the problem there. Instead I found this:
# AI assistants & answer engines — explicitly welcomed for discoverability
User-agent: GPTBot
User-agent: OAI-SearchBot
User-agent: ChatGPT-User
User-agent: ClaudeBot
User-agent: PerplexityBot
User-agent: Google-Extended
User-agent: CCBot
Allow: /
Immaculate. All three also serve a valid llms.txt . All three have canonical tags, Open Graph images, Twitter cards, and JSON-LD structured data in the <head> .
I had rolled out a red carpet, printed a welcome sign in seven languages, and pointed it all at an empty room.
The eight sites that scored fine are not better marketed. They are just built on frameworks that render on the server. That is the entire difference. Two years of work on a product, and its discoverability came down to a build-time decision I made in an afternoon.
Why this matters more this year than last
You could reasonably have ignored this in 2024. You cannot now, for three reasons that are all measurable.
One: ranking and being cited have come apart. Seer Interactive analysed 5,000+ URLs across ChatGPT, Perplexity and AI Overviews and found the overlap between top Google results and AI-cited sources fell from around 70% to under 20% . Your Google rank is no longer a proxy for your AI visibility. They are now two separate distribution channels with two separate mechanics.
Two: AI assistants overwhelmingly cite the brand’s own site. Yext analysed 6.8 million citations across ChatGPT, Gemini and Perplexity and found 86% came from brand-managed sources . The single most-cited source about your product is supposed to be you. If you are serving six words, you have forfeited that surface area by default.
Three: this is where product discovery is moving. People no longer search “best co-parenting app” and compare ten blue links. They ask an assistant and take the shortlist. If you are not in the shortlist you are not in the consideration set, and there is no second page to be on.
What llms.txt does and does not do
I want to be careful here, because there is a lot of confident nonsense about this.
llms.txt is a proposed convention: a plain-text file at your domain root summarising your site for language models. It is cheap, it is static, it is server-served, and I think you should have one.
But it is not a fix for this problem. It is a voluntary convention with inconsistent adoption, it does not replace your actual pages in any crawler’s index, and — this is the important bit — mine were already there. All three invisible sites served a perfectly good llms.txt and it did not save them, because a 110-word summary is not a substitute for a site.
Treat llms.txt as a business card. Useful. Not a building.
Check your own site in four seconds
Here is the script. It fetches a URL as GPTBot, strips scripts, styles, comments and tags, and counts what is actually left — which is what the crawler actually reads.
#!/usr/bin/env bash
# ai-visibility.sh — how many words does an AI crawler actually see?
for url in "$@"; do
words=$(curl -sL --max-time 20 \
-A "Mozilla/5.0 (compatible; GPTBot/1.0; +https://openai.com/gptbot)" "$url" \
| perl -0777 -pe 's{<script\b.*?</script>}{}gsi;
s{<style\b.*?</style>}{}gsi;
s{<noscript\b.*?</noscript>}{}gsi;
s{<!--.*?-->}{}gs;
s{<[^>]+>}{ }gs' \
| tr -s " \t\r\n" " " | wc -w)
if [ "$words" -lt 200 ]; then verdict="INVISIBLE"
elif [ "$words" -lt 600 ]; then verdict="THIN"
else verdict="OK"
fi
printf "%-10s %6s words %s\n" "$verdict" "$words" "$url"
done
Run it:
chmod +x ai-visibility.sh
./ai-visibility.sh https://yoursite.com/ https://yoursite.com/pricing
Reading the result. These thresholds are judgement calls, not standards — but they have held up across everything I have tested:
Two honest caveats. First, word count is a proxy for substance, not a measure of it — 2,000 words of nav links and legal boilerplate is not 2,000 words of value. Second, perl ships with macOS and Linux; on Windows use Git Bash or WSL.
Ranked by effort. Start at the top and stop when your number moves.
Prerender your marketing routes. For a Vite or CRA app this is usually a half-day change and it is the highest-leverage fix available. A post-build step crawls your routes in a headless browser and writes real HTML files to disk. Your app keeps working exactly as it does now; the server just has something to hand over before the JavaScript boots.
Move the marketing site to a framework that renders on the server. Next.js, Astro, Remix, Nuxt. This is the durable answer if you are rebuilding anyway. Note that you only need this for the public pages — everything behind a login can stay client-rendered forever, because no crawler should be seeing it regardless.
Split the marketing site from the app. Underrated. A static site on yourdomain.com and the app on app.yourdomain.com gives you perfect crawlability and a faster marketing site, and removes an entire category of routing problem.
Do not reach for dynamic rendering. Serving different HTML to bots than to humans based on user-agent used to be standard advice. It is fragile, it is a cloaking risk, and prerendering solves the same problem without asking anyone to trust your user-agent detection.
Then do the cheap things. Submit every sitemap to Bing Webmaster Tools — ChatGPT’s search leans on Bing’s index and this is twenty minutes of work that almost nobody does. Add Article and FAQPage schema. Put a real llms.txt up. Write your headings as questions and answer them in the first sentence, because that is the unit AI systems actually extract.
I built the prerenderer for all three sites. Here is the delta:
Around 270 pages now ship as real HTML. Zero fall under the 200-word threshold.
Three things I did not expect, all of which are the actual lessons:
The naive version silently made things worse
My first working run wrote 54 pages that all carried the homepage’s title, because I waited a fixed 500ms for the per-route metadata to apply and — under four concurrent headless tabs — it sometimes had not. Fifty-four URLs under one title is duplicate content. I only caught it because I diffed the titles.
Prerendering created a canonical-tag bug out of nothing
The HTML shell hardcodes <link rel="canonical" href="/"> . Harmless when a crawler only ever fetches one file. The moment every route is its own file, every page ships two canonicals — one truthfully pointing at itself, one insisting it is the homepage. That is worse than no prerendering at all, and one of our blogs had already been deindexed by that exact bug once before.
The build found a page nobody had noticed was broken
PostPilot’s /platforms page had never set its own title, description or canonical — it had always served the generic homepage metadata. Every other public page had it. Nothing had ever surfaced that, because in a single-page app you cannot see the difference without going looking for it.
So the honest summary is: the prerendering was the easy part. The guards were the work. If you do this, write the checks that fail your build before you write the renderer — every one of mine caught something real.
Two of the first production deploys also failed outright, because the build image had no Chrome and then no system libraries to run it with. Neither failure reached anyone: the script exits non-zero on any problem, so the previous working deploy simply stayed live. That is the argument for build-time guards that fail hard rather than warn.
No. GPTBot, ClaudeBot and PerplexityBot fetch raw HTML and do not run JavaScript. Analysis of over 500 million GPTBot fetches has found no evidence of JavaScript execution. Googlebot is the exception — it renders JavaScript using a headless Chrome engine.
Only the parts present in the HTML your server sends. If your React app renders entirely in the browser, ChatGPT sees your <title> and meta description and nothing else.
No. robots.txt grants permission to crawl. It does not create content to crawl. A site can welcome every AI crawler in existence and still serve them an empty page.
No. It is a useful, cheap addition — a summary file for language models — but adoption is inconsistent and it does not substitute for server-rendered pages. All three of my invisible sites already had one.
Fetch it with curl using the GPTBot user-agent and count the words in the response after stripping tags and scripts. Under 200 words on a marketing page means you are effectively invisible. The script is above.
Less than you would think — Googlebot renders JavaScript, so a well-built single-page app can rank fine. That is exactly what makes this dangerous: your analytics look healthy while an entire emerging channel silently returns nothing.
AI crawlers read HTML. They do not run JavaScript. If

[truncated]
