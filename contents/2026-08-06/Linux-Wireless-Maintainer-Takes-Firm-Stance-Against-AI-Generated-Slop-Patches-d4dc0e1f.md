---
source: "https://www.phoronix.com/news/Linux-WiFi-Strict-AI-Slop-Patch"
hn_url: "https://news.ycombinator.com/item?id=49197017"
title: "Linux Wireless Maintainer Takes Firm Stance Against AI Generated Slop Patches"
article_title: "Linux Wireless Maintainer Takes Firm Stance Against AI/LLM Generated Slop Patches - Phoronix"
author: "speckx"
captured_at: "2026-08-06T15:08:04Z"
capture_tool: "hn-digest"
hn_id: 49197017
score: 2
comments: 1
posted_at: "2026-08-06T14:17:02Z"
tags:
  - hacker-news
  - translated
---

# Linux Wireless Maintainer Takes Firm Stance Against AI Generated Slop Patches

- HN: [49197017](https://news.ycombinator.com/item?id=49197017)
- Source: [www.phoronix.com](https://www.phoronix.com/news/Linux-WiFi-Strict-AI-Slop-Patch)
- Score: 2
- Comments: 1
- Posted: 2026-08-06T14:17:02Z

## Translation

タイトル: Linux Wireless Maintainer、AI 生成のスロップパッチに対して断固とした態度を取る
記事タイトル: Linux Wireless Maintainer、AI/LLM 生成のスロップパッチに対して断固とした態度をとる - Phoronix
説明: Linux カーネルのステージング領域が、実際のセキュリティ修正を除き、AI/LLM で生成されたパッチを拒否することに加えて、Linux ワイヤレス ネットワーキング コードでも、AI/LLM で生成されたパッチの処理方法に若干の変化が見られます。

記事本文:
カテゴリー
コンピュータ ディスプレイ ドライバ グラフィックス カード Linux ゲーム メモリ マザーボード プロセッサ ソフトウェア ストレージ オペレーティング システム 周辺機器
Linux Wireless Maintainer は AI/LLM によって生成されたスロップ パッチに対して断固とした態度をとります
「記録のために言っておきますが、パッチを無視する機能を賢明に利用し、3 秒間のレビューで「明らかに正しい」と判断されない限り、syzbot-AI が生成したほぼすべてのパッチに適用するつもりです。
特に、私は「いつものようにパッチについてコメント」したり、他のコンピューターが電子メールで私に配信するよりも早くコードをデタラメに出力できる LLM と議論したりすることは絶対に*ありません*。
現在ループ内の人間のふりをしている人は、「絶対に」パッチについて考える必要があり、それが起こっている場合、それはおそらく AI 生成の syzbot によって送信されたものではなくなりますが、パッチを書き直して適切に送信することになります。
(IOW: 迷子になれ、syzbot)」
「このようなケースでは、システムはほぼ確実に非常に限定的で的を絞った修正を提供するだろう（それを正確に説明するうっとうしいテキストの壁がある）。しかし、少なくともループに参加している人間は実際にはそこから一歩下がって、コードのセマンティクスがどうあるべきかを自問すべきだと思う…私はこのゲームを Slawomir の最初のパッチでプレイしたことがあるが、元の仲介者がそれを望まない場合、明らかに拡張できない。
おそらく、何らかの方法で LLM にそうするように指示することもできるので、最初のドラフトの方が優れています。たとえば、この場合、同じ switch ステートメントに複数の分岐があることが意味があると、一体なぜ判断したのでしょうか。しかも分岐が 2 つしかないのです。 - 同じ検証を実装しますか?少なくとも私は、「人間の技術者」がその一歩を後退することを期待します。
これが、私がこれらのパッチを拒否する理由です。実際にわざわざ見ようとする人は明らかにいないからです。

パッチ適用前またはパッチ適用中のコードのセマンティクス。スイッチの外側のチェックを抜くとエラーの順序は変わりますか?はい。それは重要ですか？いいえ、NL80211_TDLS_ENABLE_LINK の新しいエラー順序は、実際には - よく考えてみれば (!) - はるかに理にかなっています。対象を絞った修正を行うように指示したときに、LLM がそれを実行しないことに驚いていますか?絶対に違います。
しかし、私はそのようなすべての問題について自分で判断することは実際にはできません。上記のように、できることなら、すべてを自分で行うことができます。それを行うには貢献者が必要です。私が彼にそうするよう促した後（冗談です！）、スワボミールはそれを行いましたが、それはあまりうまくいかなかったので、それについてよく話し合いましたが、繰り返しになりますが、私は常にそのサポートを提供することはできません。
...
対象となる修正であっても修正であるため、あなたの観点からはこれはまだ理にかなっていると理解していますが、私の観点からは、長期的なメンテナンス（アーキテクチャを無視してコード全体にチェックを散布する）の点でも、レビュー帯域幅などの点でも、まったく拡張できません。」
Michael Larabel は Phoronix.com の主な作成者であり、Linux ハードウェア エクスペリエンスの充実に重点を置いて 2004 年にサイトを設立しました。 Michael は、Linux ハードウェア サポートの状況、Linux パフォーマンス、グラフィックス ドライバー、その他のトピックをカバーする 20,000 件を超える記事を執筆してきました。 Michael は、Phoronix Test Suite、Phoromatic、および OpenBenchmarking.org 自動ベンチマーク ソフトウェアの主任開発者でもあります。 Twitter や LinkedIn を通じて彼をフォローしたり、MichaelLarabel.com を通じて連絡したりすることができます。
Phoronix プレミアムでは、このサイトの継続的な運営をサポートしながら、サイトへの広告なしのアクセス、単一ページの複数ページの記事、その他の機能が可能になります。
2004 年以来、Phoronix の使命は、Linux ハードウェアの強化を中心に据えてきました。

経験。広告を通じて私たちのサイトをサポートすることに加えて、 Phoronix Premium に登録することで支援することができます。 PayPal または Stripe を介したチップ/寄付を通じて Phoronix に貢献することもできます。
広告なしでブラウジングしながら、
法的免責事項、プライバシー ポリシー、Cookie |プライバシーマネージャー |お問い合わせ
著作権 © 2004 - 2026 by Phoronix Media。
使用されているすべての商標は、それぞれの所有者の財産です。無断転載を禁じます。

## Original Extract

In addition to the Linux kernel staging area now rejecting AI/LLM-generated patches except for real security fixes, the Linux wireless networking code is also seeing some shifts around how it will deal with AI/LLM generated patches.

Categories
Computers Display Drivers Graphics Cards Linux Gaming Memory Motherboards Processors Software Storage Operating Systems Peripherals
Linux Wireless Maintainer Takes Firm Stance Against AI/LLM Generated Slop Patches
"I'm going to state, for the record, that I'm going to make judicious use of ability to ignore patches, and apply it to pretty much all syzbot-AI-generated patches unless a 3-second review says "obviously right".
In particular, I will absolutely *not* "comment on the patches as usual" and argue with an LLM that can bullshit out code faster than another computer can even deliver it to me by email.
Whoever is currently pretending to be the human in the loop *absolutely* needs to think about the patches, and if that's happening it's probably no longer an AI-generated-sent-by-syzbot but you're going to rewrite it and send it properly.
(IOW: get lost, syzbot)"
"The system, in a case like this, is almost certainly going to provide a very narrow, targeted fix (with an annoying wall of text explaining exactly that), but I think that at least the human in the loop should actually take a step back from that and ask what the semantics of the code should be ... I've played this game with Slawomir's first patch myself, but that clearly cannot scale if the original intermediary doesn't want to do that.
Maybe it's something you can even tell the LLM to do, somehow, so the first draft is better. In this case, for example, why the hell did it decide that it made any sense to have multiple branches of the same switch statement - and there are even only two! - implement the same validation? At the very least I'd expect the "human engineer" to take that step back.
This is why I'm refusing these patches, because clearly nobody actually even bothers to look at the semantics of the code before or during the patching. Does pulling out the check outside of the switch change the order of errors? Yes. Does that matter? No, the new order of errors for NL80211_TDLS_ENABLE_LINK would actually - if you think about it (!) - make a lot more sense! Am I surprised the LLM doesn't do that when you tell it to make a targeted fix? Absolutely not.
But I really cannot make that judgement call myself for every single issue like that, if I could, see above, I could be doing all of this myself. Need the contributors to do that. Slawomir did that after I prompted (pun intended!) him to do that, and it didn't work out so well and we had a good discussion about it, but again, I can't provide that support all time.
...
Even the targeted fixes are still fixes, so I understand from your perspective this still made sense, but from mine it just absolutely cannot scale, both in terms of long-term maintenance (sprinkling checks all over the code disregarding the architecture) and also in terms of review bandwidth etc."
Michael Larabel is the principal author of Phoronix.com and founded the site in 2004 with a focus on enriching the Linux hardware experience. Michael has written more than 20,000 articles covering the state of Linux hardware support, Linux performance, graphics drivers, and other topics. Michael is also the lead developer of the Phoronix Test Suite, Phoromatic, and OpenBenchmarking.org automated benchmarking software. He can be followed via Twitter , LinkedIn , or contacted via MichaelLarabel.com .
Phoronix Premium allows ad-free access to the site, multi-page articles on a single page, and other features while supporting this site's continued operations.
The mission at Phoronix since 2004 has centered around enriching the Linux hardware experience. In addition to supporting our site through advertisements, you can help by subscribing to Phoronix Premium . You can also contribute to Phoronix through tips/donations via PayPal or Stripe .
While Having Ad-Free Browsing,
Legal Disclaimer, Privacy Policy, Cookies | Privacy Manager | Contact
Copyright © 2004 - 2026 by Phoronix Media .
All trademarks used are properties of their respective owners. All rights reserved.
