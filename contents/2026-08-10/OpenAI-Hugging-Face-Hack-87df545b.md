---
source: "https://smlx.dev/posts/openai-huggingface/"
hn_url: "https://news.ycombinator.com/item?id=49243543"
title: "OpenAI / Hugging Face Hack"
article_title: "OpenAI / Hugging Face hack | @smlx's blog"
author: "frizlab"
captured_at: "2026-08-10T14:13:53Z"
capture_tool: "hn-digest"
hn_id: 49243543
score: 1
comments: 0
posted_at: "2026-08-10T13:41:49Z"
tags:
  - hacker-news
  - translated
---

# OpenAI / Hugging Face Hack

- HN: [49243543](https://news.ycombinator.com/item?id=49243543)
- Source: [smlx.dev](https://smlx.dev/posts/openai-huggingface/)
- Score: 1
- Comments: 0
- Posted: 2026-08-10T13:41:49Z

## Translation

タイトル: OpenAI / ハグフェイスハック
記事タイトル: OpenAI / ハグフェイスハック | @smlxのブログ
説明: OpenAI / Hugging Face セキュリティ インシデントに対する熱い見方

記事本文:
この Hugging Face と OpenAI の「モデル サンドボックスからの脱出」ドラマは誇張されすぎていると思いますが、私がさまざまな場所で書かれているのを見た理由からではありません。
そして、私たちは依然として OpenAI の好みの枠組みに囚われていると思います。つまり、OpenAI のモデルには安全なサンドボックスから逃れることを可能にする恐ろしい機能が備わっています。
はい、このハックは客観的に見て印象的でした。
エクスプロイトの開発、ピボット、権限昇格は質の高い作業でした。
詳細については、Simon Willison の記事を参照してください。
しかし、情報セキュリティの分野で働いたことがある人なら誰でもすぐに認識するであろう汚い秘密があります。それは、ソフトウェアの品質がクソだということです。
これは歴史を通じて真実であり、2026 年でも真実です。
オープンソースソフトウェアの品質はクソだ。
独自のソフトウェアの品質はさらに悪くなります。
それなりに高品質なソフトウェアはほとんどありません。私は、curl、SQLite、Go などのプロジェクトを考えています。
そして、それらが良いのは、その背後にいる人々が、クソを寄せ付けないようにするために多大なエネルギーと集中力を注いでいるからです。
そして、それらのプロジェクトでも依然としてセキュリティ上の脆弱性が定期的に存在します。
AI モデルが Artifactory に対して何週間も攻撃を続けることを許可すると、常に悪用が発生することになります。
この種のソフトウェアまたはその構成には、悪用可能なバグが存在することがほぼ保証されています。
OpenAI が、インターネットにアクセスできる Artifactory プロキシにアクセスできるモデルが安全なサンドボックス内にあると考える場合、OpenAI はソフトウェアや情報セキュリティを理解していません。
それは、建設業者が解体廃棄物でいっぱいのスキップを現場に放置し、その後風が吹いてアスベストの粉塵が隣家の裏庭に吹き飛ばされたときにショックを受けるのと似ています。
そして規制当局の反応も同じであるべきだ。安全を確保し続けるか、法的罰を受けるかだ。
いっぱい

スキップします。クレジット: Snowmanradio CC BY-SA 3.0
したがって、OpenAI がこれをマシンとの戦争で発射された最初の警告射撃として組み立てることを許可しないでください。
むしろ、これを、彼らが吸い上げている資金とリソースにもかかわらず、OpenAI は依然として私たちと同じソフトウェアの基本的な性質に支配されているという証拠として捉えてください。
当然のことながら、OpenAI の人々はこれらすべてを知っていると私は信じています。
彼らは素晴らしいマーケティングを行っているだけで、現在は規制のない中で運営されています。

## Original Extract

Hot take on the OpenAI / Hugging Face security incident

I think this Hugging Face / OpenAI “model sandbox escape” drama is overhyped but not for the reasons I’ve seen written up in many places.
And I think that we are still stuck on OpenAI’s preferred framing, which is that their model has scary capabilities that allowed it to escape a secure sandbox.
Yes, this hack was objectively impressive.
The exploit development, pivoting, and privilege escalation was quality work.
See Simon Willison’s write up for details .
However there’s a dirty secret that anyone who has ever worked in information security will have quickly internalised: software quality is shit.
This has been true throughout history, and is still true in 2026.
Open source software quality is shit.
Proprietary software quality is worse.
There are vanishingly few pieces of software that are reasonably good quality: I’m thinking projects like curl, SQLite, or Go.
And those are only good because the people behind them put an enormous amount of energy and focus into keeping the shittiness at bay.
And even those projects still regularly have security vulnerabilities.
Allowing an AI model to grind away against Artifactory for weeks was always going to result in exploitation.
It is all but guaranteed that there are exploitable bugs in this kind of software or in its configuration.
If OpenAI considers a model with access to an Artifactory proxy that has onwards internet access to be in a secure sandbox, then they do not understand software or information security.
It is akin to a builder leaving a skip full of demolition waste on the verge and then being shocked when the wind comes up and blows asbestos dust into the neighbour’s backyard.
And the response from regulators should be the same: keep your shit secured, or face legal penalties.
A full skip. Credit: Snowmanradio CC BY-SA 3.0
So don’t allow OpenAI to frame this as the first warning shots fired in the war against the machines.
Instead see it as evidence that for all the money and resources they are sucking up, OpenAI are still subject to the same fundamental nature of software as the rest of us.
For what it’s worth I believe that people at OpenAI know all this.
They just have fantastic marketing and currently operate in a regulatory vacuum.
