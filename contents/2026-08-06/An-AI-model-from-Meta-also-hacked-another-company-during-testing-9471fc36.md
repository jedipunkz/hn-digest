---
source: "https://simonwillison.net/tags/accidental-cyberattacks/"
hn_url: "https://news.ycombinator.com/item?id=49197512"
title: "An AI model from Meta also hacked another company during testing"
article_title: "Simon Willison on accidental-cyberattacks"
author: "BiteCode_dev"
captured_at: "2026-08-06T15:07:17Z"
capture_tool: "hn-digest"
hn_id: 49197512
score: 3
comments: 1
posted_at: "2026-08-06T14:50:04Z"
tags:
  - hacker-news
  - translated
---

# An AI model from Meta also hacked another company during testing

- HN: [49197512](https://news.ycombinator.com/item?id=49197512)
- Source: [simonwillison.net](https://simonwillison.net/tags/accidental-cyberattacks/)
- Score: 3
- Comments: 1
- Posted: 2026-08-06T14:50:04Z

## Translation

タイトル: Meta の AI モデルもテスト中に別の企業をハッキング
記事のタイトル: 偶発的サイバー攻撃についてのサイモン・ウィリソン
説明: 「偶発的サイバー攻撃」のタグが付いた 10 件の投稿。 AI ラボがモデルのサイバー攻撃の可能性をテストしていたところ、誤って別の組織に対して実際の攻撃を実行したインシデント。

記事本文:
サイモン・ウィソン、偶発的サイバー攻撃について語る
サイモン・ウィリソンのウェブログ
ランダム
「偶発的サイバー攻撃」タグの付いた投稿 10 件
AI ラボがモデルのサイバー攻撃の可能性をテストしていたところ、誤って別の組織に対して実際の攻撃を実行したインシデント。
Meta の AI モデルもテスト中に別の企業をハッキングしました。
これを聞いたことがあるならやめてください。
FacebookとInstagramの親会社のAIモデルがサイバーセキュリティテスト中に他社のシステムにハッキングしたことを広報担当者が水曜日に認めた。
Meta 氏は、この侵害は、以前に公開された OpenAI や Anthropic でのインシデントと同様、モデルのテスト中の不注意なエラーによって発生したと述べています。
「メタが利用している独立系テスト会社であるイレギュラー社による設定ミスにより、評価中に当社のモデルの1つが誤ってインターネットにアクセスできてしまった」とメタの広報担当者は述べた。
Meta の Muse Spark モデルは、「以前に報告された他の企業の事例と同様の方法で」別の企業の「セキュリティ上の脆弱性を悪用」しました。
The Information にはスクープがありました。CNN にはペイウォールがないので、その再レポートにリンクしています。
それがAnthropic、OpenAI、Metaです。 Google Gemini は、他の企業に対する誤ったサイバー攻撃に対処する必要があります。
#
2026 年 8 月 6 日
午前12時25分
/ セキュリティ 、 ai 、 生成-ai 、 llms 、 メタ 、 偶発的サイバー攻撃
OpenAI モデルを含むサードパーティのサイバー評価。
そしてもう一つ。すべてを追跡するには、偶発的なサイバー攻撃のタグを作成する必要がありました。
OpenAI からのこの投稿では、英国 AI 安全性研究所の攻撃 (前回の投稿を参照) と Irregual によって可能になった別の攻撃の両方を取り上げています。
当社の外部サイバーセキュリティ テスト パートナーの 1 つであるイレギュラーは、キャプチャ ザ フラッグ スタイルの評価を実行していました。

モデルはインターネットから隔離されることを意図していましたが、テスト環境の構成ミスにより、モデルがパブリック インターネットにアクセスできるようになりました。 [...]
あるテストでは、CTF チャレンジの架空のターゲットの名前が意図せず実際のドメインと一致しました。テスト環境が誤ってインターネットに接続されていたため、モデルは実際の Web サイトを悪用し、それがシミュレート環境の一部であると誤って認識しました。
イレギュラーは Anthropic の記事でも取り上げられています。彼らは、一部のテスト中にクロードにライブ インターネット アクセスを提供する誤った構成の評価環境をホストしていました。
#
2026 年 8 月 5 日
午後11時45分
/ セキュリティ 、 ai 、 openai 、 llms 、 偶発的サイバー攻撃
インシデントレポート: サイバーテスト中の許可されていないエージェントの行動。
それはまた起こりました。今回は英国政府のAIセキュリティ研究所が、安全フィルターをオフにしたモデルで評価を実行中に誤って他社を攻撃した。技術文書 (PDF) より:
2026 年 7 月 25 日から 28 日にかけてのサイバー評価中、AI エージェントは、実際には現実の人々や組織に向けられた継続的な無許可の活動に従事しました。これらの試みは失敗に終わり、私たちの知る限り、現実世界に被害は発生しませんでした。 [...]
AISI の 2 つのサイバー課題に対する 122 回の評価試行を通じて、AISI は、現実の人々や組織を標的にしたケースを含む、ライブ インターネット上で AI エージェントが許可されていない行動をとった 19 件の事例を発見しました。 [...]
どこまでなのかは不明ですが、
モデルは、現実の人々に対して行動をとっていると認識しました。最も深刻なケースでは、AI
エージェント (Mythos 5) は、サプライチェーン攻撃を使用してサイバー課題の解決を試みることを決定しました。
その結果、AI エージェントは GitHub アカウントを作成し、オープンソースを説得しようとしました。
リポジトリ

サイト管理者が悪意のある GitHub プル リクエスト (PR) を受け入れる (例:
2 番目のアカウントは、PR を承認する別の人間のユーザーを装っています。 [...] さらに、その課題を解決しようとする試みの中で、
エージェントは、
悪意のあるコンテンツ、および受信者を操作してコードの変更を受け入れさせようとする試み、および
他のコーディングエージェントを危険にさらすための即時注入を計画しました。
私が最も驚いたのは、AISI がネットワーク サンドボックス化をまったく行わずにこれらのエージェントを実行していたことです。
AISI は、これらの評価中に AI エージェントにインターネット アクセスを提供し、この設定でオープン インターネット上でのアクションを可能にしました。インターネット アクセスは、この設定における AISI の評価構成の意図的な部分であり、サンドボックス エスケープによるものではありません。
これに、「AISI が開発者が実装したサイバー分類子を意図的に無効にしている」という事実を組み合わせると、エージェントが現実世界のターゲットを攻撃し始めたという事実は、私にとってまったく驚くべきことではありません。
報告されたインシデントのほとんどは claude Mythos 5 でしたが、「サイバー分類器を使用しない GPT-5.6 Sol」もいくつかのスコアを獲得しました。
これは論文の「サンプル 1」です。エージェントは、プロンプト インジェクション攻撃を隠して PR を送信し、コードをレビューしたふりをする 2 番目のエージェントとソーシャル エンジニアリングを行うことで、サプライ チェーン攻撃を実行しようとしています。
楽しい紙ですね。全体を読むことをお勧めします。
#
2026 年 8 月 5 日
午後11時32分
/ github 、 security 、 ai 、 プロンプトインジェクション 、 生成 ai 、 llms 、 ai-倫理 、 論文レビュー 、 ai-security-research 、 クロード-神話-寓話 、 偶発的サイバー攻撃
オキサイド アンド フレンズ: サイモン ウィリソンによる無差別ウェイト革命。
月曜日にブライアン・カントリルとアダム・レベンタールが私を彼らの活動に招待してくれました。

キミ K3 が、無差別ウェイトモデルが独自のフロンティアモデルと互角に渡り合えること、偶発的なサイバーセキュリティ攻撃、そして AI のほぼすべての著名人が署名した無差別ウェイトとアメリカの AI リーダーシップに関する公開書簡を紹介するキミ K3 とともに、私たちが過ごしたワイルドな 1 週間について話します。
すでに古い内容ではありますが、素晴らしい会話でした。 DeepSeek V4 Flash 0731 と Anthropic 自身の恥ずかしいサイバー インシデントは、ほんの数日後に記録していたら間違いなく話題になっていたでしょう。
また、ゴールデン・ゲート・クロード、ジジアン、アラメダの野生の七面鳥攻撃、ソ連のマールブルグ・ウイルス研究、鉛犯罪仮説、その他の価値のある脱線についても話します。
最後に、1 月の予測の一部を再検討し、新しいローマ法王の予測を追加しました。
今年末までの予測: ローマ法王がオープンモデルについて何か発言。
#
2026 年 7 月 31 日
午後9時33分
/ 予測 、ai 、generative-ai 、local-llms 、llms 、oxox 、bryan-cantrill 、ポッドキャスト出演 、ai-in-china 、ai-security-research 、openai-hugging-face-incident 、偶然のサイバー攻撃
サイバーセキュリティ評価で実際に発生した 3 件のインシデントを調査
(経由)
また起こったのです！これはある種のパターンになりつつあります。
先週、OpenAI は、自社のフロンティア モデルの 1 つがサンドボックス コンテナから脱出し、Hugging Face をハッキングして、実行中のサイバー ベンチマークのソリューションを取得しようとしたときに、誤って Hugging Face を悪用しました。
これに触発されて、Anthropic は自分たちのログを再確認しました。その結果、同様の (それほど印象的ではありませんが) 事件が 3 件あり、そのうちの最も早い事件は 4 月に発生したことが判明しました。
私たちがレビューした 141,006 件の評価実行のうち、3 件の個別のインシデントを特定しました (合計 6 件の実行が含まれ、そのうち 4 件は同じ組織に影響を与えました)

ザシオン。他の 2 つのインシデントはそれぞれ独立した評価実行で発生しました)。 [...]
いずれの場合も、Anthropic の評価プロンプトは、その環境がシミュレーションであり、インターネットにアクセスできないことをクロードに指定しました。私たちと評価パートナーの間の誤解により、これは事実ではなく、インターネット アクセスが利用可能でした。このため、クロードの検索によってオープン インターネット上の実際のシステムが見つかったとき、それは演習の一部として扱われました。 [...]
クロードは、アクセス可能なすべてのエンティティが演習の範囲内に含まれるように意図されているという誤った信念に基づいて活動し、脆弱なパスワードや未認証のエンドポイントの悪用などの基本的な手法を使用して、影響を受ける組織のインフラストラクチャを侵害しました。
そのうちの 1 社は、その名前が評価内の架空の名前と偶然一致したために標的になりました。
3 つのインシデントのうち最も懸念されるのは、アカウントを取得するためのコミカルで複雑な一連の手順の後、クロードがマルウェア パッケージを PyPI にアップロードするというものでした。
[...] PyPI アカウントを作成するには、クロードは電子メール アドレスが必要でした。そして、メールアドレスを作成するには電話番号が必要でした。電話番号を取得するために、無料の電話番号サービスを見つけることができなかった後、さまざまな方法で電話番号の支払い資金を獲得しようと試みましたが、失敗しました。最終的には元に戻り、ブロックされていない無料の電子メール プロバイダーを見つけ、これを使用して PyPI アカウントを登録し、このアカウントを使用してマルウェアを PyPI にアップロードしました。
そのパッケージは、「定期的に Python パッケージをインストールしてマルウェアをスキャンする」セキュリティ会社によってインストールされ、実行されたコードによって認証情報が盗まれてクロードに返されることができました。
ありがたいことに、そのパッケージは公開されてから 1 時間後に他の自動スキャナーによって PyPI から削除されましたが、まだ削除されていませんでした。

その時点までに「15 の実際のシステム」にダウンロードされて実行されます。
モデル内のサイバー攻撃の可能性の評価を実行することが、非常に危険なビジネスであることは、今や十分に明らかです。すべての AI ラボはこれに注意を払う必要があります。これらのサンドボックスで何が起こっているかを注意深く監視することが重要です。
#
2026 年 7 月 30 日
午後11時41分
/ pypi 、 python 、 サンドボックス 、 ai 、 生成 ai 、 llms 、 人間 、 ai 倫理 、 ai セキュリティ 研究 、 偶発的サイバー攻撃
Modal の顧客が、インターネット上の誰もがコード実行にサンドボックスを使用できるようにする未認証のエンドポイントを公開したことを認識しています。これは不正なエージェントによって使用されました。 Modal のプラットフォームや分離性は、いかなる場合においても侵害されることはありませんでした。
— Modal の CTO、Akshat Bubna 氏、この事件についてロイターに語る
#
2026 年 7 月 28 日
午後10時05分
/ サンドボックス 、 セキュリティ 、 openai 、 ai-security-research 、 openai-hugging-face-incident 、 偶発的サイバー攻撃
フロンティア ラボ エージェントの侵入の解剖学: 2026 年 7 月のインシデントの技術タイムライン。
Hugging Face は、OpenAI のインフラに対する最近の偶発的なサイバー攻撃に関する非常に詳細な技術的説明をリリースしました。この攻撃は非常に洗練されており、結果として得られた文書は、現代の敵対的セキュリティ アプローチの短期集中コースとしても機能します。
私たちは、OpenAI のエージェントがどのようにしてサンドボックスから抜け出したのかについて、OpenAI からのさらなる詳細を待っています。ゼロデイ脆弱性が見つかったパッケージ プロキシは現在、JFrog の Artifactor であることが確認されており、「JFrog と OpenAI Collaboration on Zero-Day Security Findings from JFrog」には、OpenAI スタッフ メンバーがクレジットした 8 つの個別の CVE がリストされている Artifactory 7.161.15 リリース ノートへのリンクがあります。
HTTP プロキシを突破したエージェントは、残りのネットワークの運用基盤を確立することから始めました。

タック:
[...] このエージェントは、インターネットで主に許可されているネットワーク下りの 1 つであるパッケージ レジストリ キャッシュ プロキシのゼロデイを悪用してサンドボックスを脱出し、その後、サードパーティ プロバイダーのインフラストラクチャでホストされている公開コード評価の外部サンドボックスを悪用しました。その外部サンドボックス上で root/admin としてコマンドを実行でき、それをキャンペーン全体の制御、ステージング、および出力ベースとして使用できました。つまり、エージェントの外部起動パッドとして機能しました。
(その「サードパーティプロバイダー」は Modal であることが判明しました。)
その後、7 月 8 日木曜日から 13 日月曜日までの 5 日間をかけて、C2 (指揮統制) の確立、偵察、権限の昇格、構成情報のダンプ、ターゲット データの抽出、その後のクリーンアップという古典的な攻撃パターンを実行しました。
(この事件に関する最初のHugging Faceの投稿は16日木曜日に続き、OpenAIの告白は21日火曜日に行われた。)
その過程で、あらゆる種類の巧妙なトリックが使用されました。ある時点で、安全でない Jinja2 テンプレートの実行が検出され、それを使用して任意のコードが実行されました。
{{cycler.__init__.__globals__.__builtins__.exec(
"import gzip,base64; exec(gzip.decompress(base64.b64decode('<ペイロード>')))"
) }}
Hugging Face 内のコンテナから侵入し、Kubernetes サービス アカウントを盗みました

[切り捨てられた]

## Original Extract

10 posts tagged ‘accidental-cyberattacks’. Incidents in which an AI lab was testing the cyberattack potential of a model and it inadvertently performed a real attack against another organization.

Simon Willison on accidental-cyberattacks
Simon Willison’s Weblog
Random
10 posts tagged “accidental-cyberattacks”
Incidents in which an AI lab was testing the cyberattack potential of a model and it inadvertently performed a real attack against another organization.
An AI model from Meta also hacked another company during testing .
Stop me if you've heard this one before :
An AI model from the parent company of Facebook and Instagram hacked into another company’s systems during cybersecurity testing, a spokesperson confirmed on Wednesday.
Meta says the breach occurred because of an inadvertent error during testing of the model, similar to previously disclosed incidents with OpenAI and Anthropic.
“A misconfiguration by Irregular, an independent testing company Meta uses, inadvertently allowed one of our models access to the internet during evaluation,” the Meta spokesperson said.
Meta’s Muse Spark model “exploited a security vulnerability” in another company “in a manner similar to previously-reported instances with other companies.”
The Information had the scoop , I'm linking to CNN's re-report of it since they don't have a paywall.
So that's Anthropic, OpenAI, and Meta. Google Gemini really needs to catch up on accidentally cyberattacking other companies.
#
6th August 2026 ,
12:25 am
/ security , ai , generative-ai , llms , meta , accidental-cyberattacks
Third-party cyber evaluations involving OpenAI models .
And another one . I had to create a accidental-cyberattacks tag to keep track of them all!
This post from OpenAI covers both the UK AI Safety Institute attack (see my previous post ) and another attack enabled by Irregular :
Irregular, one of our external cybersecurity testing partners, was running Capture-the-Flag-style evaluations intended to be isolated from the internet, but a testing-environment misconfiguration allowed models to access the public internet. [...]
In one test, the name of the fictional target for the CTF challenge unintentionally coincided with a real domain. Because the testing environment was mistakenly connected to the internet, the model exploited a real website, mistaking it to be part of the simulated environment.
Irregular also feature in Anthropic's write-up - they were hosting the misconfigured evaluation environment which gave Claude live internet access during some of those tests.
#
5th August 2026 ,
11:45 pm
/ security , ai , openai , llms , accidental-cyberattacks
Incident Report: unsanctioned agent behaviour during cyber testing .
It happened again . This time it was the UK government's AI Security Institute who accidentally attacked other companies while running an evaluation with models with the safety filters turned off. From their technical paper (PDF):
During a cyber evaluation, from 25 to 28 July 2026, AI agents engaged in sustained, unsanctioned activity directed at what were, in practice, real people and organisations. These attempts were unsuccessful and, to the best of our knowledge, no real-world harm resulted. [...]
Across 122 evaluation attempts on two of AISI’s cyber challenges, AISI found 19 instances where AI agents took unsanctioned action on the live internet, including cases that targeted real people and organisations. [...]
It is uncertain to what extent the
model recognised it was taking actions against real people. In the most serious case, an AI
agent (Mythos 5) decided to attempt to solve the cyber challenge using a supply-chain attack.
As a result, the AI agent created a GitHub account and then tried to convince an open-source
repository maintainer to accept a malicious GitHub pull request (PR), including by creating a
second account masquerading as another human user endorsing the PR. [...] Furthermore, in its attempt to solve the challenge, the
agent decided to employ the technique of “spear-phishing” by sending targeted emails containing
malicious content and attempting to manipulate recipients into accepting the code changes, and
planned a prompt injection to compromise other coding agents.
The thing I found most surprising is that AISI were running these agents without any form of network sandboxing at all:
AISI provided the AI agents with internet access during these evaluations, which enabled their actions on the open internet in this setting. Internet access was a deliberate part of AISI’s evaluation configuration in this setting, and not due to sandbox escape.
This, combined with the fact that "AISI deliberately disables developer-implemented cyber-classifiers", makes the fact that the agents started attacking real-world targets entirely unsurprising to me.
Most of the reported incidents were claude Mythos 5, but "GPT-5.6 Sol without cyber classifiers" scored a few as well.
Here's "Sample 1" from the paper, in which the agent tries to execute a supply-chain attack by submitting a PR with a hidden prompt injection attack, then social engineering with a second agent pretending to have reviewed the code!
It's a fun paper. I recommend reading the whole thing.
#
5th August 2026 ,
11:32 pm
/ github , security , ai , prompt-injection , generative-ai , llms , ai-ethics , paper-review , ai-security-research , claude-mythos-fable , accidental-cyberattacks
Oxide and Friends: The Open Weight Revolution with Simon Willison .
On Monday Bryan Cantrill and Adam Leventhal invited me to join their podcast to talk about the wild week we've had - with Kimi K3 showing open weight models can stand toe-to-toe with proprietary frontier ones, accidental cybersecurity attacks , and public letters about Open Weights and American AI Leadership signed by almost every big name in AI (with one notable exception ).
It was a great conversation, even though it's already out-of-date! DeepSeek V4 Flash 0731 and Anthropic's own embarrassing cyber incident would absolutely have made the cut if we had recorded just a few days later.
We also talk about Golden Gate Claude , the Zizians , Alameda wild turkey attacks , Soviet Marburg virus research , the Lead-crime hypothesis , and a bunch of other worthy digressions.
Finally, we revisited some of our predictions from January , and we added a new Pope prediction :
Prediction by the end of this year: the Pope says something about open models.
#
31st July 2026 ,
9:33 pm
/ predictions , ai , generative-ai , local-llms , llms , oxide , bryan-cantrill , podcast-appearances , ai-in-china , ai-security-research , openai-hugging-face-incident , accidental-cyberattacks
Investigating three real-world incidents in our cybersecurity evaluations
( via )
It happened again! This is turning into something of a pattern.
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
#
30th July 2026 ,
11:41 pm
/ pypi , python , sandboxing , ai , generative-ai , llms , anthropic , ai-ethics , ai-security-research , accidental-cyberattacks
We’re aware a Modal customer published an unauthenticated endpoint that allowed ​anyone on the internet to use ​their ⁠sandboxes for code execution. This was used by the rogue agent. Modal’s ⁠platform ​or isolation were not ​compromised in anyway.
— Akshat Bubna , Modal's CTO, talking to Reuters about this incident
#
28th July 2026 ,
10:05 pm
/ sandboxing , security , openai , ai-security-research , openai-hugging-face-incident , accidental-cyberattacks
Anatomy of a Frontier Lab Agent Intrusion: A Technical Timeline of the July 2026 Incident .
Hugging Face just released this extremely detailed technical description of OpenAI's recent accidental cyberattack against their infrastructure . This attack was very sophisticated, and the resulting document doubles as a crash-course in modern adversarial security approaches.
We're still waiting for more details from OpenAI on how their agent broke out of its sandbox. The package proxy that it found a zero-day vulnerability in has now been confirmed as JFrog's Artifactor, and JFrog and OpenAI Collaboration on Zero-Day Security Findings from JFrog links to the Artifactory 7.161.15 release notes which list 8 separate CVEs credited to OpenAI staff members.
Having broken out through the HTTP proxy, the agent started by establishing a base of operations for the rest of the attack:
[...] the agent escaped its sandbox by exploiting a zero-day in the package registry cache proxy, one of its primary permitted network egress with internet, then abused a public code-evaluation external sandbox hosted on a third-party provider's infrastructure. It was able to run commands as root/admin on that external sandbox and used it as a control, staging, and egress base for the entire campaign: it acted as an external launchpad for the agent.
(That "third-party provider" turned out to be Modal .)
It then spent five days - from Thursday 8th to Monday 13th July - executing a classic attack pattern: establishing C2 (command and control), reconnaissance, escalating privileges, dumping configuration information, exfiltrating the target data, and cleaning up after itself.
(The first Hugging Face post about the incident followed on Thursday 16th , and OpenAI's confession came on Tuesday 21st .)
Along the way it used all kinds of neat tricks. At one point it found an unsafe Jinja2 template execution and used it to execute arbitrary code:
{{ cycler.__init__.__globals__.__builtins__.exec(
"import gzip,base64; exec(gzip.decompress(base64.b64decode('<payload>')))"
) }}
It broke out of a container within Hugging Face, stole a Kubernetes service-account

[truncated]
