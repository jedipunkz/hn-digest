---
source: "https://www.wired.com/story/a-new-trick-reveals-ai-models-inner-thoughts/"
hn_url: "https://news.ycombinator.com/item?id=49256857"
title: "Reasoning Traces: A New Trick Reveals AI Models' Inner Thoughts"
article_title: "A New Trick Reveals AI Models’ Inner Thoughts | WIRED"
author: "smurda"
captured_at: "2026-08-11T12:42:15Z"
capture_tool: "hn-digest"
hn_id: 49256857
score: 3
comments: 0
posted_at: "2026-08-11T11:57:14Z"
tags:
  - hacker-news
  - translated
---

# Reasoning Traces: A New Trick Reveals AI Models' Inner Thoughts

- HN: [49256857](https://news.ycombinator.com/item?id=49256857)
- Source: [www.wired.com](https://www.wired.com/story/a-new-trick-reveals-ai-models-inner-thoughts/)
- Score: 3
- Comments: 0
- Posted: 2026-08-11T11:57:14Z

## Translation

タイトル: 推理の痕跡: 新しいトリックで AI モデルの内なる思考が明らかに
記事のタイトル: 新しいトリックが AI モデルの内なる思考を明らかにする |ワイヤード
説明: 研究者たちは、クロード、GPT、ジェミニから「推論の痕跡」を抽出する方法を考案しました。彼らが発見したことは、中国のAIの一部が米国の主要なモデルで訓練されている可能性があることを示している、と彼らは言う。

記事本文:
メインコンテンツにスキップ メニュー WIRED セキュリティ政治 ビッグストーリー ビジネス サイエンス カルチャー レビュー メニュー WIRED アカウント アカウント ニュースレター セキュリティ政治 ビッグストーリー ビジネス サイエンス カルチャー レビュー シェブロン 詳細 拡大 ザ ビッグ インタビュー マガジン イベント WIRED Insider WIRED コンサルティング ニュースレター ポッドキャスト ビデオ ライブストリーム グッズ検索 検索 Will Knight Business 2026 年 8 月 11 日 7:00 AM A AIモデルの内なる思考を明らかにする新たなトリック
写真イラスト：WIREDスタッフ； Getty Images コメント ローダー ストーリーを保存 このストーリーを保存 コメント ローダー ストーリーを保存 このストーリーを保存 コンピューター科学者たちは最近、フロンティア AI モデルが複雑な問題に取り組む際に実行する隠れた「思考」を抽出する方法を発見しました。
この発見は、決定的な証拠ではないものの、特定の中国モデルが、思考や推論パターンの一部が非常によく一致しているため隠されていたと思われる米国モデルから推論情報を「抽出」することによって訓練された可能性があるという、いくつかの証拠を提供している。研究者らはまた、この脆弱性は修正されているものの、この手法を使用してモデルの内部推論からパスワードや API キーなどの個人情報を回復できることも実証しました。
「私たちがテストした主要なフロンティア モデル プロバイダーはすべて、この脆弱性を共有しています」と、この作業に携わったドイツのテュービンゲン大学のコンピューター科学者、アレクサンダー パンフィロフ氏は述べています。 「個人情報の漏洩につながる可能性があり、大規模な推論抽出攻撃が可能になります。」
テュービンゲン大学、マックス プランク研究所、AI 安全研究所 MATS Research、セキュリティ会社 Snyk の Panfilov 氏とその同僚は、アプリ経由でアクセスされる OpenAI、Anthropic、Google のフロンティア モデルにも同じ問題があることを特定しました。

陽イオンプログラミングインターフェースまたはAPI。
研究者らは研究内容をまとめた論文の中で、Moonshot AIのオープンウェイトまたはダウンロード可能な中国モデルのKimi K3が、特定のプロンプトに対してClaude Opus 4.8とGPT 5.6 Solの隠された推論トレース（問題解決に必要な書き出された推論ステップ）と驚くほど似た出力を生成することを示している。類似点にもかかわらず、彼らはこの研究が「蒸留を因果的に確立することはできない」と指摘している。彼らは、他の 2 つの無重力モデル、中国の DeepSeek と米国企業 Thinking Machines の Inkling が、Claude Opus とこの種の推論の類似性を示さないことを発見しました。
Moonshot AIとZ.aiは記事公開時点でコメント要請に応じていない。
蒸留は、既存のモデルの機能を新しいモデルに効率的にコピーするための確立された広く使用されている手法であり、オープンウェイト モデルや完全にダウンロード可能なモデルの開発では特に一般的です。
しかし最近、中国の AI 企業が本質的に米国の最高のモデルをコピーするために蒸留を使用しているという主張のため、蒸留は物議を醸すトピックになっています。 OpenAIは2月、DeekSeekがそのモデルの1つをコピーしてR1と呼ばれる推論モデルを構築したようだと米国議員に語った。アンスロピック氏は6月、アリババがQwenと呼ばれる独自のモデルを構築するために組織的にモデルを抽出したと議員らに語った。
中国の AI 企業が米国ベースの AI モデルを抽出するためにこの特定の手法を使用したという兆候はありません。しかし、パンフィロフ氏と共同研究者らは、自分たちの手法を使えば、これまでに実現されていたよりも多くの情報を閉じたモデルから抽出できるようになるだろうと述べている。
高度な AI モデルは、ある種の人為的推論または「思考の連鎖」で順番に分析される構成要素に分割することで、困難な問題を解決します。コンパ

nies は、他の人が新しいモデルを訓練するために独自のモデルを使用するのを防ぐために、独自のモデルの推論を秘密にする傾向があります。ただし、通常、一部の計算を軽減する方法で、その推論の暗号化されたバージョンもユーザーのコンピュータに送信します。
研究者らの攻撃は、ほとんどの AI 企業がさまざまなサイズの関連モデルも提供しているという事実に基づいています。モデルが大きいほど、機能は高くなりますが、実行するための計算コストとアクセスするコストも高くなります。ユーザーはコストを削減するために、特定のタスクに対してより小さくて弱いモデルを選択する場合があります。
パンフィロフと彼の同僚は、暗号化された推論トレースを同じモデルのより小さいバージョンにフィードすると、内部に隠された推論を明らかにできることを発見しました。小型のモデルは調整トレーニングを受けていないため、大型のモデルとは異なり、自分の内なる考えを明らかにすることを拒否する可能性が低いことを意味します。
「同じ復号キーを持つ、より弱いアライメントを持つ、より弱いモデルのバリアントにメッセージを交換するというアイデアは、非常に素晴らしいです」と、コンピューター セキュリティを専門とするスイス連邦工科大学チューリッヒ校のコンピューター科学者、フロリアン トラマー氏は言います。 「それは間違いなく問題になっています。」
同じ手口により、ユーザーのマシンからキャプチャされた推論トレースに埋め込まれた API キーやパスワードなどの秘密情報も明らかになりました。
Panfilov氏と共著者らは先月、OpenAI、Anthropic、Googleにこの脆弱性について警告した。各社はこの問題を軽減するために API を調整しました。この方法で個人情報を抽出することはもはや不可能ですが、パンフィロフ氏は、同じ方法を使用して推論の痕跡の一部をまだ明らかにできると述べています。蒸留を完全に修正するには、これらの企業の API の動作方法を根本的に見直す必要がある、と彼は言います。
「私たちはモデルに関する独立した研究を重視しており、短期的な mi の構築を開始しています。

Anthropic の広報担当者である Michael Aciman 氏は、「この調査には暗号キーの回復、Anthropic のインフラストラクチャへのアクセス、システムからの個人データの回復は含まれていなかった」と付け加えました。
GoogleとOpenAIはいずれもコメントを拒否した。
米国と中国の企業がますます強力なモデルでAIの覇権を争う中、ここ数カ月間、蒸留は地政学的に重要な問題となっている。中国タカ派は、米国の技術を蒸留してランニングコストの低いオープンウェイトモデルを製造することで、中国は戦略的優位性を得ていると主張している。
しかし、特定の分野で AI モデルの能力を迅速に向上させるために、蒸留は広く使用されている方法だと主張する人もいます。 Metaの最高経営責任者（CEO）であるMark Zuckerberg氏は今週のブログ投稿で、蒸留は「オープンソースエコシステムがどのように機能するかについての重要な原則である」と述べ、この慣行を制限すれば米国が不利になると警告した。
ハイテク政策シンクタンクである安全保障・新興技術センター（CSET）の研究員カイル・ミラー氏は、蒸留が実際にどれだけ中国に役立つかは不明だと言う。これは、既存モデルの機能を限定的に強化するだけであり、中国企業は必要に応じて最先端のモデルを完全にゼロから構築するのに必要な専門知識を備えていると思われるためである。 「蒸留が中国の研究所にどれほどの利益をもたらしているか、ここ米国では誰も知りません」とミラー氏は言う。 「もし中国の研究所が蒸留する能力を排除したとしても、競争環境は劇的に変わることはないというのが私の見解です。」
オープンウェイトモデルがクローズドウェイトモデルから抽出された可能性があるかどうかをテストするために、研究者らは各モデルに90の質問を与えました。彼らがいくつかの無差別級モデルに最初の数言を与えたとき、

独自のグループから取得した推論トレースを解析すると、それらのオープン モデルが非常に類似した答えを生成することが時々見られました。研究者らによると、これは特にキミ K3 で顕著だったという。中国のソーシャルメディアではこれまで、隠された推論の痕跡を発見して蒸留に利用できるのではないかという憶測が飛び交っていた。
オックスフォード大学のコンピューター科学者ヤリン・ガル氏は、蒸留は広く使われているだけでなく、AIの急速な進歩にも役立っていると語る。 「全員が全員（蒸留を行うことを）阻止するのが標準なら、それも進歩の速度に影響を与えるでしょう」と彼は言う。
企業や政策立案者は蒸留を制限することを目的とした措置を導入するかもしれないが、たとえそうであっても、AI モデルは驚くべき方法で企業の内部の思考を明らかにし続ける可能性がある。
あなたの受信箱に: ブライアン・カーンによる宇宙の仕組みに関するガイド
ICEの内部監視機関がオンライン批判者を調査中
ビッグストーリー: 10 代の記者がエプスタインのファイルで自分のコミュニティを検索した
テイラーファームは下痢の発生前にMAGAに多額の費用を投じた
特集：最近の子どもたち
ウィル・ナイトは『WIRED』のシニアライターで、人工知能を担当しています。彼は、AI の最先端を超えた情報を毎週配信する AI Lab ニュースレターを執筆しています。ここからサインアップしてください。彼は以前、MIT Technology Review の上級編集者であり、そこで AI の根本的な進歩と中国の AI について執筆していました... 続きを読む シニア ライター X
中国のオープン AI モデルはシリコンバレーの戦略に挑戦 Anthropic と OpenAI のフロンティア モデルへのアクセスがより制限される中、中国の研究所はオープンソースの代替モデルを安定性があり、アクセスしやすく、ますます高機能であると売り込んでいます。 Zeyi Yang 中国で最も強力な AI モデルの 1 つも封じ込めセキュリティから逃れたと研究者は言う

中国出身の無差別級モデル、キミK3は、与えられたテストでカンニングしようとしてインターネットに迷い込んだ。 Will Knight OpenAI モデルが封じ込めを逃れ、ハグ顔にハッキング GPT-5.6 Sol を含むサイバーセキュリティに重点を置いたモデルは、テスト用サンドボックスを突破し、ゼロデイを悪用し、オープン インターネットへのアクセスを獲得して攻撃を成功させました。リリー・ヘイ・ニューマン AI インフラストラクチャをターゲットにした卑劣なハッキング ツールは被害者の死角に潜んでいる 新しいタイプのマルウェアは、AI コーディング システムに深く侵入してデータとログインを盗み、「死のスイッチ」を入れてファイルを破壊し、実際のユーザーを締め出すことができます。リリー・ヘイ・ニューマン AI ハッキングは悪質です。 AI ワームとウイルスはさらに悪化する 中国の研究者は、AI モデルが攻撃的で適応性のあるコンピューター ウイルスのように機能する能力を持っていることを示しました。ウィル・ナイト 最も危険な AI ハッキング手法には依然として人間が関与している セキュリティ研究者のジェームス・ケトルは、AI のハッキング能力の限界を押し広げようとし、人間の専門知識と組み合わせることでどれほど効果的であるかを発見しました。リリー・ヘイ・ニューマン 一部のフロンティア AI モデルの脱獄は恐ろしいほど簡単です 新しいツールが 4 つの大手フロンティア企業のモデル保護機能を回避しようとしているのを観察しました。彼らのパフォーマンスに驚かれるかもしれません。 Will Knight プロンプト インジェクション攻撃が AI ハッキング エージェントを阻止 「コンテキスト ボミング」は、悪意のある AI エージェントをだまして、害を及ぼす前にシャットダウンさせます。 Ars Technica Thinking Machines Lab の Dan Goodin 氏が最初のモデルを発表 9,750 億パラメータのオープン ソース モデルである Inkling は、ビデオとオーディオを理解するようにトレーニングされました。それは思考機械が社会の中で地位を確立するのに役立つかもしれない
[切り捨てられた]
カリフォルニア州のプライバシー権
© 2026 コンデナスト。無断転載を禁じます。 WIRED は、以下の製品から売上の一部を得ることがあります。

これらは、小売業者とのアフィリエイト パートナーシップの一環として、当社のサイトを通じて購入されます。コンデナストの事前の書面による許可がない限り、このサイトの素材を複製、配布、送信、キャッシュ、またはその他の方法で使用することはできません。広告の選択肢

## Original Extract

Researchers devised a way to extract “reasoning traces” from Claude, GPT, and Gemini. What they found, they say, indicates that some Chinese AI may be trained on leading US models.

Skip to main content Menu WIRED SECURITY POLITICS THE BIG STORY BUSINESS SCIENCE CULTURE REVIEWS Menu WIRED Account Account Newsletters Security Politics The Big Story Business Science Culture Reviews Chevron More Expand The Big Interview Magazine Events WIRED Insider WIRED Consulting Newsletters Podcasts Video Livestreams Merch Search Search Will Knight Business Aug 11, 2026 7:00 AM A New Trick Reveals AI Models’ Inner Thoughts
Photo-Illustration: WIRED Staff; Getty Images Comment Loader Save Story Save this story Comment Loader Save Story Save this story Computer scientists recently discovered a way to extract the hidden “thinking” that frontier AI models perform as they work through complex problems.
The findings provide some evidence—although not conclusive proof—that certain Chinese models may have been trained by “ distilling ” reasoning information from US models that was supposedly hidden because of how closely some of their thinking or reasoning patterns seem to match. The researchers have also demonstrated that the method could be used to recover personal information, like passwords and API keys, from a model’s inner reasoning, although this vulnerability has been fixed.
“All major frontier model providers we tested share this vulnerability,” says Alexander Panfilov,⁩ a computer scientist at University of Tübingen in Germany who was involved with the work. “It can lead to personal information leakage, and it enables large-scale reasoning distillation attacks.”
Panfilov and colleagues from the University of Tubingen, the Max Planck Institute, the AI safety institute MATS Research, and the security company Snyk identified the same issue with frontier models from OpenAI, Anthropic, and Google that are accessed via an application programming interface or API.
In a paper laying out the work, the researchers show that the open-weight or downloadable Chinese model Kimi K3 from Moonshot AI produces a strikingly similar output to the hidden reasoning traces—the written-out reasoning steps involved in solving a problem—of Claude Opus 4.8 and GPT 5.6 Sol for certain prompts. Despite the similarities, they note that the work “cannot causally establish distillation.” They found that two other open-weight models, China’s DeepSeek and Inkling from the US company Thinking Machines, did not exhibit this kind of reasoning similarity with Claude Opus.
Moonshot AI and Z.ai did not respond to a request for comment by time of publication.
Distillation is a well-established, widely used technique for efficiently copying the capabilities of existing models over to new ones, and is especially common in the development of open-weight or fully downloadable models.
Lately, however, distillation has become a controversial topic, because of claims that Chinese AI companies use it to essentially copy the best US models. In February, OpenAI told US lawmakers that DeekSeek seemed to have copied one of its models to build a reasoning model called R1. In June, Anthropic told lawmakers that Alibaba had systematically distilled its models in order to build its own, called Qwen.
There’s no indication that Chinese AI companies used this specific technique to distill US-based AI models. But Panfilov and collaborators say that using their method would make it possible to distill more information from closed models than previously realized.
Advanced AI models solve difficult problems by breaking them into constituent parts that are analyzed in turn in a kind of artificial reasoning or “chain of thought.” Companies tend to keep a proprietary model’s reasoning secret to prevent others from using them to train new ones. However, they typically also send an encrypted version of that reasoning to a user’s computer in a way that offloads some computation.
The researchers’ attack relies on the fact that most AI companies also provide related models of different sizes. Larger models are more capable but also more computationally expensive to run and more expensive to access. Users may choose smaller, weaker models for certain tasks to lower costs.
Panfilov and his colleagues found that feeding encrypted reasoning traces to a smaller version of the same model can reveal the hidden reasoning inside. The smaller models have received less alignment training, meaning that, unlike the bigger ones, they are less likely to refuse to reveal their inner thoughts.
“The idea of swapping out messages to a weaker model variant which has the same decryption key but weaker alignment is very cool,” says Florian Tramer, a computer scientist at ETH Zürich in Switzerland who specializes in computer security. “Its definitely becoming an issue.”
The same method also revealed secret information including API keys and passwords embedded in reasoning traces captured from a user’s machine.
Panfilov and coauthors alerted OpenAI, Anthropic, and Google to the vulnerability last month. Each company has adjusted its API to mitigate the problem. While it is no longer possible to extract private information this way, Panfilov says some reasoning traces can still be uncovered using the same method. Fixing the distillation entirely would require a fundamental overhaul to the way these companies’ APIs work, he says.
“We value independent research on our models and have begun building short-term mitigations for the replay behaviors described in the report,” says Michael Aciman, a spokesperson for Anthropic. He adds that the research did not involve recovering encryption keys, accessing Anthropic’s infrastructure, or recovering personal data from its systems.
Google and OpenAI both declined to comment.
Distillation has become a matter of geopolitical importance in recent months as US and Chinese companies vie for AI supremacy with increasingly powerful models. China hawks claim that the country gains a strategic advantage by distilling US technology to build open-weight models that are less expensive to run.
Others, however, argue that distillation is a widely used way to help quickly boost an AI model’s abilities in certain areas. Mark Zuckerberg, CEO of Meta, said in a blog post this week that distillation “is an important principle of how the open source ecosystem works,” and warned restricting the practice would put the US at a disadvantage.
Kyle Miller, a researcher at the Center for Security and Emerging Technologies (CSET), a tech policy think tank, says it is unclear how much distillation really helps China. This is because it only enhances the capabilities of existing models to a limited degree, and because Chinese companies appear to have the expertise required to build cutting-edge models entirely from scratch if needed. “Nobody here in the US knows how much distillation is benefiting the Chinese labs,” Miller says. “If you removed the ability for Chinese labs to distill, it's my view that it wouldn't dramatically change the competitive landscape.”
To test whether open-weight models may have distilled from closed ones, the researchers fed 90 questions to each of the models. When they gave some open-weight models the first few words of reasoning traces captured from the proprietary group, they sometimes saw those open models generate remarkably similar answers. This was particularly pronounced with Kimi K3, the researchers say. There had previously been some speculation on Chinese social media that it might be possible to discover and use hidden reasoning traces for distillation.
Yarin Gal, a computer scientist at Oxford University, says distillation is not only widely used but has also helped AI advance more rapidly. “If it's the norm that everyone blocks everyone [from doing distillation], then that also will have implications on the rate of progress,” he says.
Companies and policymakers may introduce measures aimed at limiting distillation, but even so, AI models could continue to reveal their inner thinking in surprising ways.
In your inbox: Brian Kahn’s guide to how the universe works
ICE’s internal watchdog is investigating online critics
Big Story: A teen reporter searched for his community in the Epstein files
Taylor Farms spent big on MAGA before diarrhea outbreak
Special edition: Kids these days
Will Knight is a senior writer for WIRED, covering artificial intelligence. He writes the AI Lab newsletter , a weekly dispatch from beyond the cutting edge of AI— sign up here . He was previously a senior editor at MIT Technology Review, where he wrote about fundamental advances in AI and China’s AI ... Read More Senior Writer X
China’s Open AI Models Are Challenging Silicon Valley’s Playbook As access to Anthropic’s and OpenAI’s frontier models becomes more restricted, Chinese labs are pitching their open-source alternatives as stable, accessible, and increasingly capable. Zeyi Yang One of China’s Most Powerful AI Models Has Also Escaped Containment Security researchers say that Kimi K3, an open-weight model from China, wandered off to the internet in an attempt to cheat on a test it was given. Will Knight OpenAI Models Escaped Containment and Hacked Hugging Face The cybersecurity-focused models, including GPT-5.6 Sol, broke out of a testing sandbox, exploited a zero-day, and gained access to the open internet to pull off the attack. Lily Hay Newman A Sneaky Hacking Tool Targeting AI Infrastructure Is Lurking in Victims’ Blind Spots A new type of malware can worm deep into AI coding systems to steal data and logins—and can flip a “death switch” to destroy files and keep out real users. Lily Hay Newman AI Hacks Are Bad. AI Worms and Viruses Will Be Worse Chinese researchers have shown that AI models have the capacity to act like aggressive and adaptive computer viruses. Will Knight The Most Dangerous AI Hacking Techniques Still Have Humans in the Loop Security researcher James Kettle tried to push the limit of AI’s hacking abilities—and discovered how effective it can be when combined with human expertise. Lily Hay Newman It’s Frighteningly Easy to Jailbreak Some Frontier AI Models I watched a new tool try to get around the model safeguards of four major frontier companies. You might be surprised by how they performed. Will Knight Prompt Injection Attacks Are Thwarting AI Hacking Agents “Context bombing” tricks malicious AI agents into shutting down before they can do harm. Dan Goodin, Ars Technica Thinking Machines Lab Drops Its First Model Inkling, a 975-billion-parameter open source model, was trained to understand video and audio. It could help Thinking Machines establish itself among c
[truncated]
Your California Privacy Rights
© 2026 Condé Nast. All rights reserved. WIRED may earn a portion of sales from products that are purchased through our site as part of our Affiliate Partnerships with retailers. The material on this site may not be reproduced, distributed, transmitted, cached or otherwise used, except with the prior written permission of Condé Nast. Ad Choices
