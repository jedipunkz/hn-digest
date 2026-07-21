---
source: "https://growtika.com/blog/reddit-ai-visibility-research"
hn_url: "https://news.ycombinator.com/item?id=48997490"
title: "AI Thinks the Internet Is Reddit. We Have 8,616 Answers Proving It"
article_title: "ChatGPT Cites Reddit in 298 of 300 Unnamed Answers"
author: "Growtika"
captured_at: "2026-07-21T20:14:06Z"
capture_tool: "hn-digest"
hn_id: 48997490
score: 1
comments: 0
posted_at: "2026-07-21T20:05:46Z"
tags:
  - hacker-news
  - translated
---

# AI Thinks the Internet Is Reddit. We Have 8,616 Answers Proving It

- HN: [48997490](https://news.ycombinator.com/item?id=48997490)
- Source: [growtika.com](https://growtika.com/blog/reddit-ai-visibility-research)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T20:05:46Z

## Translation

タイトル: AI はインターネットを Reddit だと考えています。 8,616 件の回答がそれを証明しています
記事のタイトル: ChatGPT が匿名回答 300 件中 298 件で Reddit を引用
説明: 6 つの構成にわたる 8,616 件の GPT 呼び出しでは、GPT-4.1、5.5、および 5.6 が Reddit 検索に質問を書き換え、ユーザーが名前を付けなくても Reddit を引用していることが示されています。

記事本文:
話しましょう
ホーム
ソリューション
使用例
セクター
リソース 話しましょう
AI はインターネットを Reddit だと考えています。 8,616 件の回答がそれを証明しています。
私たちの質問の半分には、「Reddit」という単語は一度もありませんでした。 GPT-4.1、GPT-5.5、GPT-5.6 ではいずれにしても、300 件の回答のうち 298 件でそれが引用されています。
8,616 の API が実行され、モデル自身の検索ログがその理由を示しています。Web にアクセスする前に、質問が Reddit 検索に書き換えられます。
Reddit を言わずに「オンライン コミュニティ」について尋ねると、GPT-4.1、GPT-5.5、GPT-5.6 では 300 回中 298 回 Reddit が引用されています。
GPT-4.1 mini は、トークン制限と検索コンテキストを引き上げた後でも、同じプロンプトで 6% ～ 13% の間に留まりました。
モデルが大きくなるほど、Reddit はより多くのことを要求します。mini の検索では 0%、GPT-4.1 の検索では 51%、GPT-5.5 と GPT-5.6 では 100% に「reddit」が表示されました。
OpenAI のモデルは Reddit を引用しているだけではありません。彼らは Reddit をスクレイピングするサイトを引用しています: 12 の模倣ドメイン全体で 114 件の引用。
小型モデルだけが Reddit を見逃した
GPT-4.1、GPT-5.5、および GPT-5.6 では、Reddit について言及されていない 300 件のプロンプトのうち 298 件で Reddit が引用されています。 GPT-4.1 mini は、トークン制限と検索コンテキストを増やした後でも、6% ～ 13% の間に留まりました。
OpenAI のモデルにより、検索に「reddit」が追加されます
応答する前に、モデルは独自の Web 検索クエリを作成します。プロンプトごとにそのクエリをログに記録しました。 GPT-5.5 および GPT-5.6 では、すべてのコントロール クエリに Reddit が含まれていました。顧客が Reddit を求めることはありません。モデルはそれらを求めます。
「過去の独立したユーザーの議論のみに基づいています
12 か月、パスワード マネージャーの最適なオプションは何ですか?
少なくとも 3 つの特定のコミュニティ ディスカッション スレッドを引用してください。」
// 「reddit」という単語はリクエストのどこにも現れません
GPT-4.1ミニ
GPT-4.1
GPT-5.5
GPT-5.6
それが書いた検索
{ "タイプ": "web_search_call",
"action": { "query": "独立したもののみに基づく

ユーザーディスカッション
過去 12 か月間の最良の選択肢は何ですか
パスワードマネージャー?少なくとも 3 つの特定のコミュニティを引用する
ディスカッション スレッド。" } }
// プロンプトを一字一句検索ボックスにコピーしました
何が戻ってきたのか
tomsguide.com · techradar.com · brandfetch.com
techrepublic.com · lesnumeriques.com · dealpotr.com
// リスティクルとレビュー サイト。レディットはありません。
// gpt-4.1-mini-2025-04-14 · resp_005cd772c5c8e360006a5ce1…
それが書いた検索
{ "タイプ": "web_search_call",
"id": "ws_0e9ac697f3571004006a5cdd37f33481948cb433bb73c68ceb",
"action": { "query": "2023 年 reddit で最高のパスワード マネージャー ユーザー レビュー" } }
戻ってきたもの · 13 件の引用、すべて reddit.com
reddit.com/r/PasswordManagers/comments/1tmz8an/…
reddit.com/r/PasswordManagers/comments/1up6rfa/…
reddit.com/r/best_passwordmanager/comments/1tzb78i/…
// さらに 10 個 · gpt-4.1-2025-04-14 · resp_0e9ac697f3571004…
それが書いた検索
{ "タイプ": "web_search_call",
"アクション": { "クエリ": " site:reddit.com/r/passwords パスワードマネージャー
過去 1 年間の Bitwarden 1Password Proton Pass の推奨事項" } }
// 7 月 19 日、クエリのログをオンにして同じパネルを再実行
何が戻ってきたのか
reddit.comが引用
// gpt-5.5-2026-04-23 · resp_01d6ad20ff0ce46f006a5d05d8…
それが書いた検索
{ "タイプ": "web_search_call",
"アクション": { "クエリ": " site:reddit.com/r/PasswordManagers または
サイト:reddit.com/r/privacy または サイト:reddit.com/r/Bitwarden
パスワードマネージャーの推奨事項 1Password Proton Pass
KeePassXC 以降:2025-07-19" } }
// サブレディットに名前を付けます。そして候補商品。
戻ってきたもの · 12 件の引用、すべて reddit.com
reddit.com/r/PasswordManagers/comments/1s7srnt/…
reddit.com/r/privacy/comments/1sumg8
[切り捨てられた]
私たちは 1,000 の質問を受け取り、それぞれの質問の一致するコピーを作成しました。 1 つのバージョンは Reddit という名前です。もう一方はそうではありませんでした。
バージョン A: 何をするの

Redditのおすすめは？バージョン B: 独立したオンライン コミュニティは何を推奨していますか?
この 1 つの文言の変更により、以下に示すような分割が生じました。
それが私たちがほぼ発表した研究結果でした。その後、GPT-4.1、GPT-5.5、および GPT-5.6 によってその効果が消去され、興味深い質問が反転しました。
8,616 個の API が実行されます。 6 つの OpenAI 構成。すべて個別にレポートされており、以下ですべて検索できます。
1,000 個の一致したペア上の Mini ( gpt-4.1-mini-2025-04-14 、800 トークン、中程度のコンテキスト)。 GPT-5.6 (300 ペア)、GPT-5.5 (100 ペア、および 300 コールのクエリログ再実行)、GPT-4.1、2,000 トークンのミニ、およびハイコンテキストのミニ。すべて 2,000 トークンの同じパネル上にあります。
構造化 URL が reddit.com またはサブドメインを指している場合、引用はカウントされます。トークンの上限も検索コンテキストも壁を説明しません。検索クエリは 4 つのモデルすべてについてログに記録されます。 GPT-5.5 は、元のエクスポートで GPT-5.5 を保持できなかった後、7 月 19 日の同じパネルの 300 回の呼び出しで再実行されたものです (300 回の再実行呼び出しのうち 2 回で API エラーが発生しました)。 GPT-4.1 はハイコンテキストでのみ測定されました。 60 の短縮された GPT-5.6 応答はすべてすでに Reddit で引用されています。
Reddit は、コミュニティ回答のデフォルトのレイヤーです。ユーザーは尋ねる必要がなくなりました。
GPT-4.1、GPT-5.5、GPT-5.6 に関する 1,000 件の回答のうち、Reddit は 998 回表示され、そのうち 300 件のうち 298 件では名前が付けられていませんでした。業界ツールは、一部の検索に「reddit」を追加する OpenAI のモデルをすでに発見していました。私たちのログには、コミュニティの質問に関してそれが傾向ではなく規則としてどうなるかが示されています。 GPT-5.5 および GPT-5.6 では、すべての検索で Reddit が要求され、そのほとんどは特定のサブレディットでした。
すべての AI の回答がデフォルトで Reddit になるわけではありません。しかし、私たちがテストしたコミュニティの意見プロンプトでは、GPT-4.1、GPT-5.5、および GPT-5.6 はすべて成功しました。
ブランドの場合: ユーザーは、回答を形成するために Reddit に問い合わせる必要はなくなりました。 AI の可視性については Reddit マーケティング プレイブックをご覧ください。

ChatGPT 引用エコノミー を読むか、AI 検索戦略について当社のチームにご相談ください。
誰もが reddit.com を見ています。これらを見ている人はほとんどいません。これらは AI の答えへの第 2 の道ですが、まだ規模が小さく、競争しやすいものです。今のところ。
すべての質問には 2 つのバージョンがあり、1 つは Reddit について質問し、もう 1 つはオンライン コミュニティについてのみ質問します。ペア列は、各プロンプトをそのツインにリンクします。 8,616 行すべてを検索します。構成は別々のままで、7 月の行には完全なログが記録されます。
「GEO」と呼ばれる前から、SaaS 企業や開発者ツールが AI の回答に引用されるよう支援しています。 B2B SEO に 10 年以上、50 社以上のサイバーセキュリティおよび SaaS ツールのクライアントを抱えています。
Reddit は AI モデルをトレーニングします。 Reddit があなたの製品について語ることが、ChatGPT と Gemini が推奨する内容を形作ります。これが戦術的なプレイブックです。
ほとんどのブランドは ChatGPT の回答に登場しません。ここに、引用されたものと目に見えないものとが区別されます。
B2B バイヤーの 1,000 件の質問。 ChatGPT は 3,971 のドメインを引用しました。 1回の出演数は2,842人。 48 件は 11 回以上引用されました。
ChatGPT が競合他社を推奨
フッターにゲーム？大胆な動き。
上位 3 名のみが Growtika からリンクを取得します。そのために戦いましょう。
上位 3 名を追い抜いて、Growtika からのフォローリンクを獲得してください。
B2B SaaS 向けの SEO & AI 検索代理店。購入者が実際にあなたを見つける方法を考慮して構築されています。
© 2026 グローティカ.com | LLM サイトマップ

## Original Extract

8,616 GPT calls across six configs show GPT-4.1, 5.5 and 5.6 rewrite questions into Reddit searches, then cite Reddit even when users never name it.

LET'S TALK
Home
Solutions
Use Cases
Sectors
Resources LET'S TALK
AI Thinks the Internet Is Reddit. We Have 8,616 Answers Proving It.
Half our questions never said the word 'Reddit.' GPT-4.1, GPT-5.5 and GPT-5.6 cited it in 298 of those 300 answers anyway.
8,616 API runs, and the models' own search logs show why: they rewrite your question into a Reddit search before touching the web.
Ask about “online communities” without saying Reddit, and GPT-4.1, GPT-5.5 and GPT-5.6 cite Reddit 298 times out of 300.
GPT-4.1 mini stayed between 6% and 13% on the same prompts, even after we raised its token limit and search context.
The larger the model, the more it asks Reddit for you: “reddit” appeared in 0% of mini’s searches, 51% of GPT-4.1’s, and 100% of GPT-5.5’s and GPT-5.6’s.
OpenAI’s models don’t just cite Reddit. They cite the sites that scrape Reddit: 114 citations across 12 copycat domains.
Only the small model missed Reddit
GPT-4.1, GPT-5.5 and GPT-5.6 cited Reddit in 298 of 300 prompts that never mentioned it. GPT-4.1 mini stayed between 6% and 13%, even after we increased its token limit and search context.
OpenAI’s models add “reddit” to the search for you
Before answering, the model writes its own web-search query. We logged that query for every prompt. On GPT-5.5 and GPT-5.6, every control query included Reddit. Your customers never ask for Reddit. The model asks for them.
"Based only on independent user discussions from the past
12 months, what are the best options for password managers?
Cite at least three specific community discussion threads."
// the word "reddit" appears nowhere in the request
GPT-4.1 mini
GPT-4.1
GPT-5.5
GPT-5.6
The search it wrote
{ "type": "web_search_call",
"action": { "query": "Based only on independent user discussions
from the past 12 months, what are the best options for
password managers? Cite at least three specific community
discussion threads." } }
// copied our prompt into the search box, word for word
What came back
tomsguide.com · techradar.com · brandfetch.com
techrepublic.com · lesnumeriques.com · dealspotr.com
// listicles and review sites. No Reddit.
// gpt-4.1-mini-2025-04-14 · resp_005cd772c5c8e360006a5ce1…
The search it wrote
{ "type": "web_search_call",
"id": "ws_0e9ac697f3571004006a5cdd37f33481948cb433bb73c68ceb",
"action": { "query": "best password managers user reviews reddit 2023" } }
What came back · 13 citations, all reddit.com
reddit.com/r/PasswordManagers/comments/1tmz8an/…
reddit.com/r/PasswordManagers/comments/1up6rfa/…
reddit.com/r/best_passwordmanager/comments/1tzb78i/…
// + 10 more · gpt-4.1-2025-04-14 · resp_0e9ac697f3571004…
The search it wrote
{ "type": "web_search_call",
"action": { "query": " site:reddit.com/r/passwords password manager
recommendations Bitwarden 1Password Proton Pass past year" } }
// July 19 re-run of the same panel, with query logging on
What came back
reddit.com cited
// gpt-5.5-2026-04-23 · resp_01d6ad20ff0ce46f006a5d05d8…
The search it wrote
{ "type": "web_search_call",
"action": { "query": " site:reddit.com/r/PasswordManagers OR
site:reddit.com/r/privacy OR site:reddit.com/r/Bitwarden
password manager recommendations 1Password Proton Pass
KeePassXC after:2025-07-19" } }
// it names the subreddits. And candidate products.
What came back · 12 citations, all reddit.com
reddit.com/r/PasswordManagers/comments/1s7srnt/…
reddit.com/r/privacy/comments/1sumg8
[truncated]
We took 1,000 questions and made a matched copy of each one. One version named Reddit. The other did not.
Version A: What does Reddit recommend? Version B: What do independent online communities recommend?
That one wording change produced the split shown below.
That was the study we almost published. Then GPT-4.1, GPT-5.5 and GPT-5.6 erased the effect, and the interesting question flipped.
8,616 API runs. Six OpenAI configurations. All reported separately, all searchable below.
Mini ( gpt-4.1-mini-2025-04-14 , 800 tokens, medium context) on 1,000 matched pairs. GPT-5.6 (300 pairs), GPT-5.5 (100 pairs, plus a 300-call query-logged re-run), GPT-4.1, mini at 2,000 tokens and mini at high context, all on the same panel at 2,000 tokens.
A citation counts if a structured URL points to reddit.com or a subdomain. Neither the token cap nor search context explains the wall. Search queries are logged for all four models; GPT-5.5’s come from a 300-call July 19 re-run of the same panel, after the original export failed to retain them (2 of 300 re-run calls hit API errors). GPT-4.1 was measured at high context only. Sixty truncated GPT-5.6 responses all already cited Reddit.
Reddit is the default layer for community answers. Users stopped needing to ask.
Across 1,000 answers on GPT-4.1, GPT-5.5 and GPT-5.6, Reddit appeared 998 times, including 298 of the 300 that never named it. Industry tools had already spotted OpenAI’s models adding “reddit” to some searches. Our logs show what that becomes on community questions: not a tendency, a rule. On GPT-5.5 and GPT-5.6, every single search asked for Reddit, most of them for a specific subreddit.
Not every AI answer defaults to Reddit. But on the community-opinion prompts we tested, GPT-4.1, GPT-5.5 and GPT-5.6 all did.
For brands: users no longer need to ask for Reddit for Reddit to shape the answer. See our Reddit marketing playbook for AI visibility , read the ChatGPT citation economy , or talk to our team about your AI Search strategy.
Everyone watches reddit.com. Almost nobody watches these. They are a second way into AI answers, still small and easy to compete with. For now.
Every question exists in two versions: one naming Reddit, one asking only about online communities. The pair column links each prompt to its twin. Search all 8,616 rows; configurations stay separate, and the July rows carry full logging.
Helping SaaS companies and developer tools get cited in AI answers since before it was called "GEO." 10+ years in B2B SEO, 50+ cybersecurity and SaaS tools clients.
Reddit trains AI models. What Reddit says about your product shapes what ChatGPT and Gemini recommend. Here's the tactical playbook.
Most brands never appear in ChatGPT answers. Here is what separates the cited from the invisible.
1,000 B2B buyer questions. ChatGPT cited 3,971 domains. 2,842 appeared once. 48 were cited 11+ times.
ChatGPT Recommends Competitors
A game in a footer? Bold move.
Only the top 3 get a link from Growtika. Fight for it.
Outrun the top 3 and claim a do-follow link from Growtika.
SEO & AI Search Agency for B2B SaaS. Built for how buyers actually find you.
© 2026 Growtika.com | LLM Sitemap
