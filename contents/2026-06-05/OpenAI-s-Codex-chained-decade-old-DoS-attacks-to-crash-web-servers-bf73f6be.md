---
source: "https://www.theregister.com/security/2026/06/04/openais-codex-chains-decade-old-dos-techniques-into-http/2-bomb/5251377"
hn_url: "https://news.ycombinator.com/item?id=48407277"
title: "OpenAI's Codex chained decade-old DoS attacks to crash web servers"
article_title: "OpenAI's Codex chains decade-old DoS techniques into HTTP/2 Bomb"
author: "sbulaev"
captured_at: "2026-06-05T04:34:15Z"
capture_tool: "hn-digest"
hn_id: 48407277
score: 1
comments: 0
posted_at: "2026-06-05T02:34:54Z"
tags:
  - hacker-news
  - translated
---

# OpenAI's Codex chained decade-old DoS attacks to crash web servers

- HN: [48407277](https://news.ycombinator.com/item?id=48407277)
- Source: [www.theregister.com](https://www.theregister.com/security/2026/06/04/openais-codex-chains-decade-old-dos-techniques-into-http/2-bomb/5251377)
- Score: 1
- Comments: 0
- Posted: 2026-06-05T02:34:54Z

## Translation

タイトル: OpenAI の Codex が 10 年にわたる DoS 攻撃を連鎖させて Web サーバーをクラッシュさせた
記事のタイトル: OpenAI の Codex が 10 年前の DoS テクニックを HTTP/2 Bomb に連鎖させる
説明: Codex が HTTP/2 爆弾を投下

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
OpenAI のエージェントが 10 年にわたる DoS 攻撃を連鎖させ、Web サーバーを数秒でクラッシュさせた
サーバーが直面する次の脅威は、ボットによって助けられた可能性があります。カリフォルニア州のセキュリティ研究者によると、OpenAI の Codex エージェントは、単一マシンから起動して脆弱な Web サーバーを数秒でアクセス不能にするリモート サービス拒否 (DoS) エクスプロイトの発見に貢献しました。
この攻撃は、nginx、Apache HTTP Server、Microsoft IIS、Envoy、Cloudflare Pingora などの主要な Web サーバーのデフォルトの HTTP/2 構成に対して作用します。研究者らによると、木曜日の時点でMicrosoft IISとCloudflare Pingoraにはまだパッチが存在しないが、Cloudflareはこの調査結果に異議を唱えている。
「Cloudflareの既存のアーキテクチャとDDoS軽減策は、この攻撃を自動的に検出して保護するため、顧客はこの脆弱性から回復できるようになります」と広報担当者はThe Registerに語った。 「パッチは必要ありません。」
Microsoft の広報担当者は、「当社は顧客の保護を維持するための適切な緩和策を認識しており、積極的に調査しています」と The Register に語った。
カリフォルニア州の研究者 Quang Luong 氏がこのエクスプロイトを発見し、HTTP/2 Bomb と名付けました。今月下旬に開催される Real World AI Security カンファレンスで攻撃の技術的な詳細を完全に発表する予定です。それまでのところ、GitHub には概念実証のエクスプロイト スクリプトがあり、AI レッドチームのセキュリティ ショップからの警告とともに「所有していないインフラストラクチャにこれらのスクリプトを指定しないでください」とされています。
ルオン氏は火曜日のブログで、Codex が 10 年以上前から知られている 2 つの既存の DoS 攻撃手法、HPACK 圧縮爆弾と Slowloris スタイルのホールドを連鎖させたと述べ、8 を超える攻撃が行われていると警告しました。

HTTP/2 をサポートし、脆弱な Web サーバーのいずれかを実行している 80,000 の Web サイトが影響を受ける可能性があります。
HPACK 爆弾攻撃 (CVE-2016-6581 としても知られています) は、HTTP/2 ヘッダー圧縮アルゴリズム (HPACK) を悪用して、何千もの小さなメッセージをサーバーに送信し、サーバーにメモリの急速な割り当てを強制し、最終的にクラッシュさせます。
次に、Slowloris DoS 攻撃 ( CVE-2016-8740 および CVE-2016-1546 ) は、正当な接続を開き、可能な限り長く維持することでサーバーを圧倒します。
この 2 つを組み合わせると、サーバーのメモリが使い果たされ、強制的にオフラインになります。
「100Mbps 接続の家庭用コンピュータは、数秒以内に脆弱なサーバーにアクセスできなくなる可能性があります」とルオン氏は書いています。 「Apache httpd や Envoy に対して、1 つのクライアントは約 20 秒で 32 GB のサーバー メモリを消費して保持できます。」
カリフォルニア州の研究チームは 4 月にこの問題を nginx に公表し、Web サーバーの保守担当者は翌日、freenginx から max_headers ディレクティブをインポートするバージョン 1.29.8 でこの問題を修正しました。
Apache は、Calif がレポートを提出したのと同じ日に修正プログラム ( mod_http2 v2.0.41 ) を発行し、CVE-2026-49975 を割り当てました。
「上記の修正コミットは公開されており、ベクターを直接公開しています。有能な AI モデルであれば、それらの差分を有効なエクスプロイトに変えることができます。これがまさに、Microsoft IIS、Envoy、Pingora にも脆弱性があることを発見した方法です」と脅威ハンティング チームは書いており、3 つすべてに通知されていると付け加えました。
水曜日の更新で、Calif は Envoy のパッチが「この攻撃を緩和すると思われる」と指摘し、修正が機能することを確認するために研究者らがまだ検証中であると述べた。
「このソフトウェアを起動しないでください」: 壊れたバックアップが rsync プロジェクトの AI コーディング行を引き起こす
AI エージェントは脆弱性を見つけるだけでなくエクスプロイトを作成できることを示す
OpenClaw のようなエージェント ハーネスは、ビジネスの方法を変えています。

AI モデルを作成して実行する
混乱を引き起こすコンピューター ワームを構築するのに Mythos やゼロデイは必要ありません – 無料のオープンソース モデルは問題なく動作します
Microsoft IIS と Cloudflare Pingora の場合、セキュリティ調査担当者は、可能であれば HTTP/2 を無効にするか、クライアントが 1 回のリクエストでサーバーに送信できる HTTP ヘッダーの数に上限を設けることを推奨しています。
カリフォルニア州によれば、人間ではなくコーディングエージェントがこの攻撃を発見したという事実は注目に値する。「両方の半分は10年前から公になっていた」とルオン氏は書いている。 「Codex が行ったことは、コードベースを読み取り、2 つが構成されていることを認識し、複合攻撃を構築することでした。その組み合わせは一度見れば明らかですが、私たちが知る限り、人間がこれらのサーバーに対してそれを組み合わせた人は一人もいません。」 ®
Microsoft からの声明により 2023 年に更新されました。
AI と ML
AIのスプリントを遅らせることが「世界にとって良いことだ」とアンスロピック氏は語る
警戒を求める嘆願は、AI最大のライバルであるOpenAIを破ってIPOを申請したのと同じ週に発表された
ピンクは、偽のヘルプデスクへの電話を使用して資格を盗む最新の悪党チームです
混沌とした犯罪集団 Lapsus$ によって広められたおなじみの戦術
ZTEとパートナーは2026年のエンジニアリング能力構築プログラムを通じて世界的なICT人材を育成
世界的な ICT 専門家が深センに集まり、最先端のエンジニアリング実践を習得し、国際協力を促進します
カナダは米国のボットから解放され、独自のAIを開発したいと考えている
別の同盟国もアメリカのAIへの依存に疑問を呈する
SaaS
AWSはエンタープライズ需要がゼロにもかかわらず、イーロン・マスク氏のGrokをBedrockに押し込むと報じられている
フロンティアモデルのエナジードリンク
OpenAI のエージェントが 10 年にわたる DoS 攻撃を連鎖させ、Web サーバーを数秒でクラッシュさせた
AI と ML
Netflix のウィズが AI 費用を削減するアプリを作成し、それをオープンソース化
セキュリティ
レドモンド市が警察に通報、マイクロソフト社に「屈辱を受けた」と不満を抱く0日ハンターが「骨が砕けるほどの衝撃」を誓う

s
セキュリティ
軍隊の携帯電話が位置データを外国の敵に提供した
セキュリティ
すべてのパスワードは Active Directory の説明フィールドに保存されました
パーソナルテクノロジー
カリフォルニア州、3Dプリントによるゴーストガンのアルゴリズムによる死亡を宣言する法案を可決
データ主権のトレードオフを克服する
避けられないトレードオフであるデータ主権はネットワークにとって実際に何を意味するのでしょうか?もっと詳しく知る。
プロンプトからエクスプロイトまで: LLM は API 攻撃をどのように変化させているか
最新のアプリケーションは API 主導で相互接続されており、多くの場合、過剰な権限が与えられているため、AI 支援攻撃の理想的な標的となっています。
未来の構築: Kubernetes 向けエンタープライズ データ サービスのロックを解除する
インフラストラクチャのサイロを排除し、標準化されたエンタープライズ グレードのクラウド ネイティブ プラットフォームを確立する方法を発見してください。
Microsoft 365 が見逃す高度な攻撃を Behavioral AI セキュリティで捕捉
Microsoft 365 は企業コミュニケーションのバックボーンであり、そのネイティブ セキュリティが既知のものや騒々しいものを除外します。
ランサムウェア侵害の混乱に足を踏み入れ、対応スキルをテストし、他の IT およびセキュリティの専門家とチームを組んでサイバー犯罪者を出し抜こう
仮想サイバー復旧シミュレーション
ランサムウェア攻撃の勢いは衰えていませんが、私たちも同様です。 Druva の人気イベント「Escape Ransomware」が完全にバーチャルになりました。
大規模なエージェント AI: パイロットから運用まで
AI の導入を大規模に推進することで、真の ROI を実現する方法を学びましょう。
警戒を求める嘆願は、AI最大のライバルであるOpenAIを破ってIPOを申請したのと同じ週に発表された
混沌とした犯罪集団 Lapsus$ によって広められたおなじみの戦術
別の同盟国もアメリカのAIへの依存に疑問を呈する
科学者と業界リーダーはDNA合成スクリーニングの義務化を推進
警戒を求める嘆願は、AI最大のライバルであるOpenAIを破ってIPOを申請したのと同じ週に発表された
人気のおなじみの戦術

混沌とした犯罪集団 Lapsus$ によって化されました
別の同盟国もアメリカのAIへの依存に疑問を呈する
科学者と業界リーダーはDNA合成スクリーニングの義務化を推進
AIのスプリントを遅らせることが「世界にとって良いことだ」とアンスロピック氏は語る
警戒を求める嘆願は、AI最大のライバルであるOpenAIを破ってIPOを申請したのと同じ週に発表された
ピンクは、偽のヘルプデスクへの電話を使用して資格を盗む最新の悪党チームです
混沌とした犯罪集団 Lapsus$ によって広められたおなじみの戦術
カナダは米国のボットから解放され、独自のAIを開発したいと考えている
別の同盟国もアメリカのAIへの依存に疑問を呈する
OpenAI のエージェントが 10 年にわたる DoS 攻撃を連鎖させ、Web サーバーを数秒でクラッシュさせた
Codex が HTTP/2 爆弾を投下
AIの重鎮、自社の技術がテロリストの生物兵器開発に役立つ可能性があると警告
科学者と業界リーダーはDNA合成スクリーニングの義務化を推進
慈悲深い独裁者ザック氏、キーロギングによるプライバシー侵害からメタスタッフに30分間の休憩を与える
スタッフの活動を盗み見てAIにコンピュータの使い方を教えるテックビジネス
お問い合わせ
私たちと一緒に宣伝しましょう
私たちは誰なのか
ニュースレター
次のプラットフォーム
開発クラス
ブロックとファイル
状況出版
クッキーポリシー
プライバシーポリシー
利用規約
私の個人情報を共有しないでください
同意のオプション
著作権。すべての著作権は © 1998-2026 に留保されます。

## Original Extract

Codex drops an HTTP/2 Bomb

Jump to main content
Search
TOPICS
Special Features
All Special Features
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
OpenAI's agent chained decade-old DoS attacks to crash web servers in seconds
The next threat your server faces may have been helped along by a bot. OpenAI's Codex agent helped uncover a remote denial-of-service (DoS) exploit that can be launched from a single machine to render vulnerable web servers inaccessible in seconds, according to Calif security researchers.
The attack works on default HTTP/2 configurations of major web servers including nginx, Apache HTTP Server, Microsoft IIS, Envoy, and Cloudflare Pingora. As of Thursday, Microsoft IIS and Cloudflare Pingora still don’t have a patch, according to the researchers, although Cloudflare disputes this finding.
“Cloudflare's existing architecture and DDoS mitigations automatically detect and protect against this attack, making customers resilient to this vulnerability,” a spokesperson told The Register . “No patch is needed.”
“We are aware and actively investigating appropriate mitigations to help keep customers protected," a Microsoft spokesperson told The Register .
Calif researcher Quang Luong discovered the exploit, named it HTTP/2 Bomb, and will present the full technical details of the attack at the Real World AI Security conference later this month. In the meantime, there are proof-of-concept exploit scripts on GitHub along with a warning from the AI red teaming security shop: “Please don't point these at infrastructure you don't own.”
In a Tuesday blog , Luong says Codex chained two existing DoS attack techniques that have been known for more than a decade - HPACK compression bomb and Slowloris-style hold - and warns that upwards of 880,000 websites supporting HTTP/2 and running one of the vulnerable web servers may be affected.
An HPACK bomb attack (also known as CVE-2016-6581 ) exploits the HTTP/2 header compression algorithm (HPACK) by sending thousands of tiny messages to the server, forcing it to rapidly allocate memory and ultimately crash.
Then the Slowloris DoS attack ( CVE-2016-8740 and CVE-2016-1546 ) overwhelms the server by opening legitimate connections and maintaining them as long as possible.
Combining the two exhausts the server’s memory and forces it offline.
“A home computer on a 100Mbps connection can render a vulnerable server inaccessible within seconds,” Luong wrote. “Against Apache httpd and Envoy, a single client can consume and hold 32GB of server memory in roughly 20 seconds.”
The Calif research team disclosed the issue to nginx in April, and the web server’s maintainers fixed it the next day in version 1.29.8, which imports the max_headers directive from freenginx.
Apache issued a fix ( mod_http2 v2.0.41 ) the same day that Calif submitted its report, and assigned it CVE-2026-49975.
“The fix commits above are public and disclose the vectors directly; any capable AI model can turn those diffs into a working exploit, which is exactly how we found that Microsoft IIS, Envoy, and Pingora are also vulnerable,” the threat hunting team wrote, adding that all three have been notified.
In a Wednesday update, Calif pointed to Envoy patches “that appear to mitigate this attack,” and notes that its researchers are still validating the fix to ensure it works.
'Please do not vibe f--- up this software': Broken backups spark AI coding row in rsync project
AI agents show they can create exploits, not just find vulns
Agent harnesses, like OpenClaw, are changing how we build and run AI models
Nobody needs Mythos or 0-days to build a chaos-causing computer worm – free open source models work just fine
For Microsoft IIS and Cloudflare Pingora, the security sleuths recommend disabling HTTP/2 if possible, or enforcing a cap on the number of HTTP headers a client can send in a single request to the server.
The fact that a coding agent - not a human - discovered this attack is notable, according to Calif. “Both halves have been public for a decade,” Luong wrote. “What Codex did was read the codebases, recognize that the two compose, and build the combined attack. That combination is obvious once you see it, and yet as far as we can tell no human had put it together against these servers.” ®
Updated at 2023 with statement from Microsoft.
AI AND ML
'It would be good for the world' to slow down AI sprints, Anthropic says
The plea for caution comes the same week it beat AI archrival OpenAI to filing for an IPO
Pink is the latest goon squad to use fake helpdesk calls to steal creds
A familiar tactic popularized by chaotic crime crew Lapsus$
ZTE and partners nurture global ICT talent through 2026 engineering capacity building program
Global ICT experts gather in Shenzhen to master cutting-edge engineering practices and foster international collaboration
Canada wants to make its own AI, break free from US bots
Another ally questions reliance on American AI
SaaS
AWS reportedly to tuck Elon Musk's Grok into Bedrock, despite zero enterprise demand
The energy drink of frontier models
OpenAI's agent chained decade-old DoS attacks to crash web servers in seconds
AI and ML
Netflix wiz creates app to slash AI bills, then open sources it
Security
Disgruntled 0-day hunter 'humiliated' by Microsoft pledges 'bone shattering drop' as Redmond calls cops
Security
Troops’ phones gave away location data to foreign adversaries
SECURITY
All the passwords were stored in Active Directory description fields
Personal tech
California passes bill declaring death-by-algorithm to 3D-printed ghost guns
Overcoming the trade-offs in data sovereignty
What does data sovereignty actually mean for your network, which trade-offs are unavoidable? Learn more.
From Prompt to Exploit: How LLMs Are Changing API Attacks
Modern applications are API-driven, interconnected, and often over-permissioned, making them an ideal target for AI-assisted attacks.
Architecting the Future: Unlocking Enterprise Data Services for Kubernetes
Join us to discover how to eliminate infrastructure silos and establish a standardized, enterprise-grade cloud-native platform.
Catch the Advanced Attacks Microsoft 365 Misses with Behavioral AI Security
Microsoft 365 is the backbone of enterprise communication, and its native security filters out the known and the noisy.
Step into the chaos of a live ransomware breach, test your response skills, and team up with other IT and security pros to outsmart cybercriminals
Virtual Cyber Recovery Simulation
Ransomware attacks aren’t slowing down, and neither are we. Druva’s hit event, Escape Ransomware, is now fully virtual.
Agentic AI at Scale: From Pilot to Production
Join us to learn how to unlock real ROI by driving adoption of AI at scale.
The plea for caution comes the same week it beat AI archrival OpenAI to filing for an IPO
A familiar tactic popularized by chaotic crime crew Lapsus$
Another ally questions reliance on American AI
Scientists and industry leaders push for mandatory DNA synthesis screening
The plea for caution comes the same week it beat AI archrival OpenAI to filing for an IPO
A familiar tactic popularized by chaotic crime crew Lapsus$
Another ally questions reliance on American AI
Scientists and industry leaders push for mandatory DNA synthesis screening
'It would be good for the world' to slow down AI sprints, Anthropic says
The plea for caution comes the same week it beat AI archrival OpenAI to filing for an IPO
Pink is the latest goon squad to use fake helpdesk calls to steal creds
A familiar tactic popularized by chaotic crime crew Lapsus$
Canada wants to make its own AI, break free from US bots
Another ally questions reliance on American AI
OpenAI's agent chained decade-old DoS attacks to crash web servers in seconds
Codex drops an HTTP/2 Bomb
AI heavyweights warn their tech could help terrorists develop bioweapons
Scientists and industry leaders push for mandatory DNA synthesis screening
Benevolent dictator Zuck will give Meta staff 30-minute breaks from keylogging privacy assault
Tech biz teaching AI to use computers by slurping staff activity
Contact us
Advertise with us
Who we are
Newsletter
The Next Platform
DevClass
Blocks and Files
Situation Publishing
Cookies Policy
Privacy Policy
Ts & Cs
Do not share my personal information
Your Consent Options
Copyright. All rights reserved © 1998-2026.
