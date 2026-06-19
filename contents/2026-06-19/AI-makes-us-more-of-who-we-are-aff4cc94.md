---
source: "https://www.ricky-dev.com/ai/2026/06/ai-makes-us-more-of-ourselves/"
hn_url: "https://news.ycombinator.com/item?id=48601431"
title: "AI makes us more of who we are"
article_title: "AI makes us more of who we are - Ricky Smith"
author: "DigitallyBorn"
captured_at: "2026-06-19T18:42:29Z"
capture_tool: "hn-digest"
hn_id: 48601431
score: 1
comments: 1
posted_at: "2026-06-19T18:14:54Z"
tags:
  - hacker-news
  - translated
---

# AI makes us more of who we are

- HN: [48601431](https://news.ycombinator.com/item?id=48601431)
- Source: [www.ricky-dev.com](https://www.ricky-dev.com/ai/2026/06/ai-makes-us-more-of-ourselves/)
- Score: 1
- Comments: 1
- Posted: 2026-06-19T18:14:54Z

## Translation

タイトル: AI は私たちをより自分らしくしてくれる
記事のタイトル: AI は私たちをより自分らしくする - リッキー・スミス
説明: AI が行った

記事本文:
AI は私たちをより自分らしくしてくれる - リッキー・スミス
リッキー・スミス
について
ツイート
AI は私たちをより自分らしくしてくれる
そしてそれは間違った人々を生産的にする
AI によって悪いエンジニアが優秀になったわけではありません。それによって彼らは速くなったのです。それは私たちが何者であるかを変えるのではなく、私たちをもっと私たちらしくするのです。これは力を倍増させるものであり、良い意味で言っているわけではありません。
かつては、怠け者またはずさんなプログラマーは自己制限を行っていました。彼らはゆっくりと動きました。彼らは怠け者の性質上、大きな変化を起こす意欲がありませんでした。彼らのコードと構造は貧弱だったかもしれませんが、生産性も貧弱でした。彼らの誤った決定の影響範囲は小さいままでした。 AI エージェント (そしてそれを使用する管理者) が登場すると、同じ人間が同じ判断、同じ盲点で 10 倍のコードを出荷し、その速度がさらに速くなるだけです。
これは優れたイコライザーですが、最悪の場合、値に関係なく出力が上昇します。
これは、Kubernetes ワークロードの IAM ロールを管理する中央リポジトリや可観測性のためのサービス監視コードを管理する別のリポジトリなど、いくつかの大規模な共有リポジトリで機能していることがわかります。これらのリポジトリとパイプラインは集中管理されていますが、コードの所有権は分散化されています。ディレクトリ構造によって断片化されています。チームがサービスを追加する必要がある場合、一般的なアプローチは、既存のサービスのコードを見つけてコピーし、実行するために最小限の変更を加えるというものです。そもそもそのモジュールが適切に設計されているかどうかを立ち止まって尋ねる人はいません。パターンを見てそれをコピーするだけです。これは何年も続いています。そして、それが長引けば長引くほど、それが正しいアプローチであるに違いないと誰もが確信を持つようになります。結局のところ、それが誰もが行う方法です。彼らがフォークのフォークから作業していること、そして元のソースがまったくよく考えられていない可能性があることに誰も気づきません。
AI がすぐにスロットに入る

そのワークフロー。乱雑な共有リポジトリを指定してサービスを追加するように依頼すると、怠惰なエンジニアがやったこととまったく同じことを実行します。つまり、最も近い既存のサンプルにパターン マッチしてコピーします。 「このモジュールは本当に良いのか?」ということは問いません。 — それは「すでにここにあるものは何ですか？」と尋ねます。人間が最も抵抗の少ない道を選択するのと同様に、エントロピーが最も小さい経路よりも最も抵抗の小さい道をデフォルトで選択します。違いは、人間はそれに対して罪悪感がちらつくかもしれないということです。 AI は自信を持ってコピーを出荷するだけです。そのため、AI は怠惰なエンジニアの速度を上げるだけでなく、怠惰なエンジニアのように動作し、エンジニアとまったく同じように悪いパターンを破ることを避けます。それは技術的負債を永続させます。
そのため、二重の苦しみに見舞われることになります。悪いエンジニアは、改善することなくより速く行動し、問題を解決することを目的としたツールは、すでに存在する混乱をすべて石灰化するデフォルトの傾向を持っています。大規模に、自信を持って、人間が立ち止まる可能性がある自信を失うことはありません。
リッキーがこれを入力しました。編集にはLLMを使用しました。
シスコのスタッフ ソフトウェア エンジニア

## Original Extract

AI didn

AI makes us more of who we are - Ricky Smith
Ricky Smith
About
Tweets
AI makes us more of who we are
And it makes the wrong people productive
AI didn't make bad engineers good. It made them fast. It doesn't change who we are, it makes us more of who we are. It's a force multiplier and I don't mean that in a good way.
Used to be, a lazy or sloppy coder was self-limiting. They moved slowly. By nature of being lazy, they weren't motivated to make huge changes. Their code and structures may have been poor, but so, too, was their productivity. The blast radius of their bad decisions stayed small. Now with AI agents (and management breathing down our neck to use them) the same person ships 10x the code with the same judgment, same blind spots — just compounding faster.
It's the great equalizer, but in the worst way: it raises output regardless of value.
I see this at work with some of our large, shared repositories – like in a central repository that manages IAM roles for kubernetes workloads and another that manages service monitoring code for observability. These repositories and pipelines are centraly owned, but the code ownership is decentralized; fragmented by a directory structure. When a team needs to add their service, the common approach is to find an existing service's code, copy it, and make the minimum changes to get running. Nobody stops to ask if that module was well-designed in the first place — they just see a pattern and copy it. This has gone on for years. And the longer it goes on, the more confidence everyone has that it must be the right approach — after all, that's how everyone does it. Nobody notices they're working from forks-of-forks, and that the original source might never have been well thought out at all.
AI slots right into that workflow. Point it at a messy shared repo and ask it to add a service, and it'll do exactly what the lazy engineer did: pattern-match to the nearest existing example and copy it. It doesn't ask "is this module actually good?" — it asks "what's already here?" Like a human taking the path of least resistance, it defaults to the path of least resistance over the path of least entropy. The difference is a human might feel a flicker of guilt about it. AI just confidently ships the copy. And so, AI doesn't just speed up lazy engineers, it behaves like one — it avoids breaking bad patterns exactly like the engineer would. It perpetuates the tech debt.
So you get a double whammy: bad engineers move faster without getting better, and the tool meant to fix things has a default tendency to calcify whatever mess already exists — at scale, with confidence, and without the self-doubt that might've made a human pause.
Ricky typed this. An LLM was used for editing.
Staff Software Engineer at Cisco
