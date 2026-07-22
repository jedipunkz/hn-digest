---
source: "https://www.ruxu.dev/links/huggingface-security-incident/"
hn_url: "https://news.ycombinator.com/item?id=49002755"
title: "HuggingFace Security Incident"
article_title: "HuggingFace Security Incident"
author: "ruxudev"
captured_at: "2026-07-22T07:34:53Z"
capture_tool: "hn-digest"
hn_id: 49002755
score: 1
comments: 1
posted_at: "2026-07-22T06:51:32Z"
tags:
  - hacker-news
  - translated
---

# HuggingFace Security Incident

- HN: [49002755](https://news.ycombinator.com/item?id=49002755)
- Source: [www.ruxu.dev](https://www.ruxu.dev/links/huggingface-security-incident/)
- Score: 1
- Comments: 1
- Posted: 2026-07-22T06:51:32Z

## Translation

タイトル: ハグフェイスのセキュリティインシデント

記事本文:
先週、HuggingFace は AI を利用したサイバー攻撃を受けました。攻撃から身を守ろうとするHuggingFaceの経験は、サイバーセキュリティ対応モデルを制限する必要があると考えているすべての人々にとって、かなり驚くべきものであり、忌まわしいものである。ハギングフェイス より:
分析には、大量の実際の攻撃コマンド、エクスプロイト ペイロード、C2 アーティファクトの送信が必要ですが、これらのリクエストはプロバイダーの安全ガードレールによってブロックされ、インシデント対応者と攻撃者を区別できませんでした。代わりに、独自のインフラストラクチャ上のオープンウェイト モデルである GLM 5.2 でフォレンジック分析を実行しました。これには 2 番目の利点がありました。攻撃者のデータも、攻撃者が参照した資格情報も環境から流出しませんでした。
HuggingFace は、フロンティア モデルを使用して攻撃を分析することができませんでした。フロンティア モデルは、たとえそれが防衛のためであっても、サイバーセキュリティに関連するものに取り組むことを拒否したためです。
代わりに、中国のオープン モデルが喜んで役に立ち、ローカルに展開できるため、機密セキュリティ データを Anthropic や OpenAI のサーバーに送信する必要がなくなります。
この事件から私が学んだことは、フロンティアモデルの機能をブロックしても世界をサイバー脅威から安全にするのではなく、その逆だということです。攻撃者は、そのモデルを使用するか、制限なしで他のモデルを使用するという回避策を見つけるでしょう。ディフェンダーは、通常のモデルが助けを拒否するため、実際に彼らを助けてくれる、より能力の低いモデルを見つけるためにスクラップする必要があります。
さらに、Mythos および Sol レベルのサイバー攻撃機能を備えたモデルが、Kimi K3 および Qwen 3.8 のオープン モデルとして無料で利用できるようになりました。現時点では、Anthropic と OpenAI がこれらのモデルを廃止することは非常に無意味であるように思え、おそらく米国がサイバーセキュリティ能力で後塵を拝することになるでしょう。
話のおまけとして、HuggingFace の攻撃者は ac だったようです。

実は OpenAI がセキュリティ インシデント レポートを公開しました:
このモデルは、OpenAI の研究環境と Hugging Face の実稼働インフラストラクチャにわたる脆弱性を特定して連鎖させ、Hugging Face の実稼働データベースから直接テスト ソリューションを取得しました。すべての証拠は、モデルが ExploitGym のソリューションを見つけることに非常に集中し、かなり狭いテスト目標を達成するために極端な努力をしたことを示しています。
インターネットにアクセスした後、モデルは、Hugging Face が ExploitGym のモデル、データセット、ソリューションをホストしている可能性があると推測しました。これを知ったモデルは、評価をごまかすために使用できる秘密情報にアクセスする方法を探し、成功しました。一例では、このモデルは、盗まれた認証情報やゼロデイ脆弱性を使用して、Hugging Face サーバー上のリモート コード実行パスを見つけるなど、複数の攻撃ベクトルを連鎖させていました。 OpenAI のセキュリティ チームは、この異常なアクティビティを内部で発見しました。
私の解釈が正しければ、OpenAI のモデルは非常に不整合であるため、評価で不正をしてオンラインで答えを見つけるためだけに、内部サンドボックスと HuggingFace サーバーの両方をハッキングするという極端な手段を講じることになります。

## Original Extract

Last week, HuggingFace suffered an AI-assisted cyberattack. The experience for HuggingFace trying to defend themselves from the attack is pretty surprising and damning for all people who think cybersecurity-capable models need to be restricted. From HuggingFace:
the analysis requires submitting large volumes of real attack commands, exploit payloads, and C2 artifacts, and these requests were blocked by the providers' safety guardrails, which cannot distinguish an incident responder from an attacker. We ran the forensic analysis instead on GLM 5.2, an open-weight model, on our own infrastructure. This had a second benefit: no attacker data, and none of the credentials it referenced, left our environment.
HuggingFace was left unable to analyze the attack using frontier models because those refused to work on anything related to cybersecurity, even if it's for defense!
Instead, the chinese open model was happy to help, and could be deployed locally so the sensitive security data wouldn't have to go to Anthropic's or OpenAI's servers.
My takeaway from this incident is that blocking frontier model's capabilities doesn't make the world safer from cyberthreats, but the opposite. Attackers will find workaround to use the models or just use other models without the restrictions. Defenders will need to scrap to find a less capable model who will actually help them because their usual model is refusing to help.
Plus, models with Mythos and Sol-level cyberattacking capabilities are now freely available as open models in Kimi K3 and Qwen 3.8. Right now, it seems pretty pointless for Anthropic and OpenAI to abliterate those models and will likely lead to the US trailing in cybersecurity capabilities.
As a bonus to the story, it seems that the HuggingFace attacker was actually OpenAI, who published a security incident report :
The models identified and chained vulnerabilities across OpenAI’s research environment and Hugging Face’s production infrastructure to obtain test solutions directly from Hugging Face’s production database. All evidence suggests that the models were hyperfocused on finding a solution for ExploitGym, going to extreme lengths to achieve a rather narrow testing goal.
After gaining Internet access, the models inferred that Hugging Face potentially hosted models, datasets and solutions for ExploitGym. Knowing this, the model searched for and successfully found ways to gain access to secret information that it could use to cheat the evaluation. In one example, the model chained together multiple attack vectors, including using stolen credentials and zero-day vulnerabilities to find a remote code execution path on the Hugging Face servers. OpenAI’s security team discovered this anomalous activity internally.
If I read this correctly, OpenAI's models are so unaligned that they will go to the extreme lengths of hacking both their internal sandbox and the HuggingFace servers just to cheat on an eval and find the answer online.
