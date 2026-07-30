---
source: "https://chsxf.dev/2026/07/27/18-firefox-setting-up-google-search-without-ai.html"
hn_url: "https://news.ycombinator.com/item?id=49114955"
title: "Google Search on Firefox Without AI"
article_title: "Google Search on Firefox Without AI | chsxf.dev"
author: "speckx"
captured_at: "2026-07-30T20:56:55Z"
capture_tool: "hn-digest"
hn_id: 49114955
score: 2
comments: 0
posted_at: "2026-07-30T19:58:01Z"
tags:
  - hacker-news
  - translated
---

# Google Search on Firefox Without AI

- HN: [49114955](https://news.ycombinator.com/item?id=49114955)
- Source: [chsxf.dev](https://chsxf.dev/2026/07/27/18-firefox-setting-up-google-search-without-ai.html)
- Score: 2
- Comments: 0
- Posted: 2026-07-30T19:58:01Z

## Translation

タイトル: AI を使用しない Firefox での Google 検索
記事のタイトル: AI を使用しない Firefox での Google 検索 | chsxf.dev
説明: 現在、EU では Google 検索に AI 概要が含まれているため、AI のスロップを削除するために広告ブロッカーを使用している人を多く見てきました。私の謙虚な意見では、これは十分な解決策ではなく、この機能とそれに関連するリソースの消費を無効にするものではありません。今回はFirefoxの設定方法を解説します。
[切り捨てられた]

記事本文:
AI を使用しない Firefox での Google 検索 | chsxf.dev
chsxf 。開発者
GitHub
AI を使用しない Firefox での Google 検索
ここ数日間、EU では Google 検索に AI 概要が含まれるようになり、多くの人が UBlock Origin または同様の機能を備えた拡張機能を使用して、検索結果の上部にある AI のスロップを削除しているのを目にしました。ただし、これはバックグラウンドで実行される機能を実際に無効にすることなく出力をクリーンアップするだけであるため、十分な解決策ではないと思います。エネルギーは依然として浪費され、混乱だけが隠蔽されるだけです。
完璧ではないにしても、より良い解決策はあります。この記事では、AI を使用せずに Google 検索を使用するように Firefox を設定する方法を説明します。私自身試したことはありませんが、同じ方法が主要なブラウザに適用できるはずです。
このチュートリアルは macOS 上の Firefox 153 向けに書かれていますが、ブラウザの以前のバージョンと以降のバージョン、およびすべてのオペレーティング システムで動作するはずです。ただし、オプションページのデザインは異なる場合があります。
残念ながら、AI を使用せずに Google ホームページから直接検索を開始するソリューションは存在しません。したがって、Firefox が新しいウィンドウやタブを開くように設定されている場合は、変更する必要があると思います。幸いなことに、残りの部分を正しく設定すれば、「Firefox ホーム」ページでも検索に関してほぼ同じ機能が使用できます。したがって、私の推奨事項は、あらゆる状況で「Firefox ホーム」をホームページとして使用するように Firefox を設定し直すことです。
Google 検索で AI をバイパスする
Firefox には、誰でもブラウザに検索エンジンを追加できる非常に便利な機能があります。構文は非常に単純で、利用可能な場合は提案もサポートします。この機能を利用して、代替の「AI なしの Google」検索エンジン エントリを作成します。残念ながら、Firefox では cust が許可されていません

組み込みまたは拡張機能が提供する検索エンジン エントリの省略。
これにより、新しい検索エンジン エントリのさまざまなパラメータを要求するポップアップが開きます。入力するさまざまな値は次のとおりです。最後の 2 つに到達するには、「詳細」ボタンをクリックする必要がある場合があります。
「キーワード」パラメータに関する予備的な注意事項:
このパラメータを使用すると、必要なときにいつでもアドレス バーで特定の検索エンジンを呼び出すことができます。たとえば、私は主に Google ではなく Lilo を使用しますが、Google の方が優れている場合もあります。そこで、アドレスバーに「@g」と入力するだけで、Firefox は自動的に Google に切り替わります (G 文字で始まる他の検索エンジンはこれだけなので)。残念ながら、複数の検索エンジンを同じキーワードで構成することはできません (もちろん、組み込みの Google エントリには @google ショートカットがあります)。
検索エンジン名:
IA のない Google
検索語の代わりに %s を含む URL:
https://www.google.com/search?udm=web&q=%s
キーワード (オプション):
@google_no_ai
検索語の代わりに %s を使用してデータを POST (GET の場合は空のままにします):
<空のままにしておきます>
検索語の代わりに %s を含む候補 URL (オプション):
https://www.google.com/complete/search?client=chrome&q=%s
AI を使用しない Google をデフォルトの検索エンジンとして設定する (オプション)
Google をデフォルトの検索エンジンとして使用している場合は、AI を含まないバージョンを使用する必要があります。また、設定で通常の Google 検索エンジンをオフにすることをお勧めします。
これで準備完了です。当面は AI サマリーなしで Google を使用できるはずです。それが長く続くことを祈りましょう。

## Original Extract

Now that Google Search includes AI summaries in the EU, I’ve seen many people using adblockers to remove the AI slop. In my humble opinion, this is not a good enough solution and does not disable the feature and the associated resource consumption. In this post, I will explain how to set up Firefox
[truncated]

Google Search on Firefox Without AI | chsxf.dev
chsxf . dev
GitHub
Google Search on Firefox Without AI
For the past few days, now that Google Search includes AI summaries in the EU, I’ve seen many people using UBlock Origin or extensions with similar capabilities to remove the AI slop at the top of the search results. However, I think this is not a good enough solution as it only cleans the output without actually disabling the feature that runs in the background. The energy will still be wasted, and the disturbance only will be put under the rug.
There are better solutions for that, even if not perfect. This post will explain how to configure Firefox to use Google Search without AI. The same method should be applicable to any major browser, even though I’ve not tried it myself.
This tutorial has been written for Firefox 153 on macOS but should work in previous and later versions of the browser, and on all operating systems. The design of the options pages may differ though.
Unfortunately, no solution exists to start a search without AI from the Google homepage directly. So, if your Firefox is configured to open its new windows and tabs to that, I’m afraid it has to change. Fortunately, the “Firefox Home” page allows pretty much the same functionality regarding search, once we have setup the rest correctly. So my recommendation is to set Firefox back to use the “Firefox Home” as its homepage in all situations.
Bypassing AI for Google Search
Firefox has a very useful feature that allows anybody to add search engines to the browser. The syntax is pretty simple and it supports suggestions too if available. We will leverage this feature to create an alternate “Google without AI” search engine entry. Unfortunately, Firefox doesn’t allow customization on built-in or extension provided search engine entries.
This will open a popup requesting the various parameters for the new search engine entry. Here are the various values to enter. You may have to click on the “Advanced” button to reach the last two.
Some preliminary note on the “Keyword” parameter:
This parameter allows you to call for a specific search engine in the address bar whenever it is needed. For example, I mostly use Lilo instead of Google, but sometimes Google is just better. So I just type @g in the addresse bar and Firefox automatically switches to Google (as this is my only other search engine starting with the G letter). Unfortunately, you can’t have mulitple search engines configured with the same keyword (and of course the built-in Google entry has the @google shortcut).
Search engine name:
Google without IA
URL with %s in place of search term:
https://www.google.com/search?udm=web&q=%s
Keyword (optional):
@google_no_ai
POST data with %s in place of search term (leave empty for GET):
<leave empty>
Suggestions URL with %s in place of search term (optional):
https://www.google.com/complete/search?client=chrome&q=%s
Setting Google Without AI as Your Default Search Engine (Optional)
If you are using Google as your default search engine, you should then use the version without AI now. I also recommend switching off the regular Google search engine in the settings.
You are now good to go. You should be able to use Google without any AI summary for the foreseable future. Let’s hope it will last long.
