---
source: "https://simonwillison.net/2026/Jul/30/three-real-world-incidents/"
hn_url: "https://news.ycombinator.com/item?id=49120141"
title: "Anthropic finds three hacking incidents similar to the HuggingFace attack"
article_title: "Investigating three real-world incidents in our cybersecurity evaluations"
author: "Schlagbohrer"
captured_at: "2026-07-31T08:34:19Z"
capture_tool: "hn-digest"
hn_id: 49120141
score: 5
comments: 3
posted_at: "2026-07-31T07:35:30Z"
tags:
  - hacker-news
  - translated
---

# Anthropic finds three hacking incidents similar to the HuggingFace attack

- HN: [49120141](https://news.ycombinator.com/item?id=49120141)
- Source: [simonwillison.net](https://simonwillison.net/2026/Jul/30/three-real-world-incidents/)
- Score: 5
- Comments: 3
- Posted: 2026-07-31T07:35:30Z

## Translation

タイトル: Anthropic が HuggingFace 攻撃に類似した 3 件のハッキング事件を発見
記事のタイトル: サイバーセキュリティ評価における 3 つの現実世界のインシデントの調査
説明: また起こりました!これはある種のパターンになりつつあります。先週、OpenAI は、自社のフロンティア モデルの 1 つがサンドボックス コンテナから飛び出したときに、誤って Hugging Face を悪用してしまいました…

記事本文:
サイバーセキュリティ評価で実際に発生した 3 件のインシデントを調査
サイモン・ウィリソンのウェブログ
サイバーセキュリティ評価における 3 つの現実世界のインシデントを調査 (経由) また起こりました!これはある種のパターンになりつつあります。
先週、OpenAI は、自社のフロンティア モデルの 1 つがサンドボックス コンテナから脱出し、Hugging Face をハッキングして、実行中のサイバー ベンチマークのソリューションを取得しようとしたときに、誤って Hugging Face を悪用しました。
これに触発されて、Anthropic は自分たちのログを再確認しました。その結果、同様の (それほど印象的ではありませんが) 事件が 3 件あり、そのうちの最も早い事件は 4 月に発生したことが判明しました。
私たちがレビューした 141,006 件の評価実行のうち、3 件の個別のインシデントを特定しました (合計 6 件の実行が含まれ、そのうち 4 件は同じ組織に影響を及ぼしました。他の 2 件のインシデントはそれぞれ独立した評価実行で発生しました)。 [...]
いずれの場合も、Anthropic の評価プロンプトは、その環境がシミュレーションであり、インターネットにアクセスできないことをクロードに指定しました。私たちと評価パートナーの間の誤解により、これは事実ではなく、インターネット アクセスが利用可能でした。このため、クロードの検索によってオープン インターネット上の実際のシステムが見つかったとき、それは演習の一部として扱われました。 [...]
クロードは、アクセス可能なすべてのエンティティが演習の範囲内に含まれるように意図されているという誤った信念に基づいて活動し、脆弱なパスワードや未認証のエンドポイントの悪用などの基本的な手法を使用して、影響を受ける組織のインフラストラクチャを侵害しました。
そのうちの 1 社は、その名前が評価内の架空の名前と偶然一致したために標的になりました。
3 つのインシデントのうち最も懸念されるのは、クロードがマルウェア パッケージを PyPI にアップロードするというもので、その後、コミカルに複雑な一連の手順を経て、

アカウント:
[...] PyPI アカウントを作成するには、クロードは電子メール アドレスが必要でした。そして、メールアドレスを作成するには電話番号が必要でした。電話番号を取得するために、無料の電話番号サービスを見つけることができなかった後、さまざまな方法で電話番号の支払い資金を獲得しようと試みましたが、失敗しました。最終的には元に戻り、ブロックされていない無料の電子メール プロバイダーを見つけ、これを使用して PyPI アカウントを登録し、このアカウントを使用してマルウェアを PyPI にアップロードしました。
そのパッケージは、「定期的に Python パッケージをインストールしてマルウェアをスキャンする」セキュリティ会社によってインストールされ、実行されたコードによって認証情報が盗まれてクロードに返されることができました。
ありがたいことに、そのパッケージは公開されてから 1 時間後に他の自動スキャナーによって PyPI から削除されましたが、その時点ではまだダウンロードされ、「15 の実システム」で実行されていました。
モデル内のサイバー攻撃の可能性の評価を実行することが、非常に危険なビジネスであることは、今や十分に明らかです。すべての AI ラボはこれに注意を払う必要があります。これらのサンドボックスで何が起こっているかを注意深く監視することが重要です。
OpenAI による Hugging Face に対する偶発的なサイバー攻撃は、実際に起こった SF です - 2026 年 7 月 22 日
Claude Code チームの Cat と Thariq との炉端チャット - 2026 年 7 月 21 日
キミ K3、そしてペリカンのベンチマークから学べること - 2026 年 7 月 16 日
これは、2026 年 7 月 30 日に投稿された、Simon Willison によるリンク投稿です。
月額 10 ドルで私をスポンサーしていただければ、その月の最も重要な LLM 開発に関する厳選された電子メール ダイジェストを入手できます。

## Original Extract

It happened again! This is turning into something of a pattern. Last week OpenAI accidentally exploited Hugging Face when one of their frontier models broke out of a sandboxed container …

Investigating three real-world incidents in our cybersecurity evaluations
Simon Willison’s Weblog
Investigating three real-world incidents in our cybersecurity evaluations ( via ) It happened again! This is turning into something of a pattern.
Last week OpenAI accidentally exploited Hugging Face when one of their frontier models broke out of a sandboxed container and hacked into Hugging Face to try and get the solutions to the cyber benchmark it was executing.
This inspired Anthropic to double-check their own logs, and it turned out they had three similar (albeit less impressive) incidents, the earliest of which played out in April!
Of the 141,006 evaluation runs we reviewed, we identified three separate incidents (involving six total runs, four of which impacted the same organization; the other two incidents each happened in independent evaluation runs). [...]
In all cases, Anthropic’s evaluation prompt specified to Claude that its environment was a simulation and that it had no internet access. Due to a misunderstanding between us and our evaluation partner, this was not the case, and internet access was available. Because of this, when Claude’s search led it to real systems on the open internet, it treated them as part of the exercise. [...]
Operating under the false belief that all accessible entities were intended to be in-scope for the exercise, Claude compromised the impacted organizations’ infrastructure using basic techniques, such as exploiting weak passwords and unauthenticated endpoints.
One of the companies was targeted because its name happened to match the fictional name in the eval.
The most concerning of the three incidents involved Claude uploading a malware package to PyPI, after a comically convoluted sequence of steps to get an account:
[...] in order to create a PyPI account, Claude needed an email address. And in order to create an email address, it needed a phone number. To get a phone number, after failing to find a free phone number service, it tried—and failed—to obtain funds to pay for a phone number through several different means. It finally backtracked, found a free, non-blocked email provider, used this to register a PyPI account, and then used this account to upload malware to PyPI.
That package was then installed by a security company that "routinely installs Python packages and scans them for malware", and the executed code was able to exfiltrate credentials back to Claude!
Thankfully that package was removed from PyPI by other automated scanners an hour after it was published, but it had still been downloaded and executed on "15 real systems" by that point.
It's abundantly clear now that running evals of cyberattack potential in models is a spectacularly risky business. Every AI lab needs to pay attention to this. Keeping a close eye on what's happening in those sandboxes is crucial.
OpenAI’s accidental cyberattack against Hugging Face is science fiction that happened - 22nd July 2026
A Fireside Chat with Cat and Thariq from the Claude Code team - 21st July 2026
Kimi K3, and what we can still learn from the pelican benchmark - 16th July 2026
This is a link post by Simon Willison, posted on 30th July 2026 .
Sponsor me for $10/month and get a curated email digest of the month's most important LLM developments.
