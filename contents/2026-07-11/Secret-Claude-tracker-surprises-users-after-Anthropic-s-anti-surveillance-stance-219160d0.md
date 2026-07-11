---
source: "https://www.theregister.com/ai-and-ml/2026/07/01/anthropic-is-removing-its-covert-code-for-catching-chinese-competitors/5265366"
hn_url: "https://news.ycombinator.com/item?id=48876037"
title: "Secret Claude tracker surprises users after Anthropic's anti-surveillance stance"
article_title: "Anthropic is removing its covert code for catching Chinese competitors"
author: "gnabgib"
captured_at: "2026-07-11T21:39:49Z"
capture_tool: "hn-digest"
hn_id: 48876037
score: 1
comments: 0
posted_at: "2026-07-11T21:27:07Z"
tags:
  - hacker-news
  - translated
---

# Secret Claude tracker surprises users after Anthropic's anti-surveillance stance

- HN: [48876037](https://news.ycombinator.com/item?id=48876037)
- Source: [www.theregister.com](https://www.theregister.com/ai-and-ml/2026/07/01/anthropic-is-removing-its-covert-code-for-catching-chinese-competitors/5265366)
- Score: 1
- Comments: 0
- Posted: 2026-07-11T21:27:07Z

## Translation

タイトル: Anthropic の反監視姿勢を受けて秘密のクロード追跡者がユーザーを驚かせる
記事のタイトル: Anthropic は中国の競合他社を捕まえるための秘密コードを削除しています
説明: ああ、そうだ、私たちは

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
FIS と AWS による金融サービスの最新化
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
Anthropic、中国の競合他社を捕まえるための秘密コードを削除
そうそう、秘密のステガノグラフィー システムを無効にするつもりでした
Anthropic は、自社のモデルから盗もうとしている他の AI 企業を捕まえるために、数か月前に Claude Code に追加した隠しコードを削除する予定だと述べています。
Anthropic のエンジニアで Claude Code チームに所属する Thariq Shihipar 氏は火曜日、修正版は 7 月 1 日に公開されるはずだと語った。
「これは、無許可の再販業者によるアカウントの悪用を防ぎ、蒸留から保護することを目的として、3月に開始した実験です」とShihipar氏は、繰り返しのクエリによるAIモデルのコピーを意味する業界用語を使って説明した。 「それ以来、チームはより強力な緩和策を講じてきました。実際、私たちはしばらくの間、これを廃止するつもりでした。」
同氏は、コードを削除するプルリクエストはマージされており、水曜日の Claude Code リリースに掲載されるはずだと述べた。
Thereallo という名前の開発者が説明したこの実験は、Anthropic のサーバーに渡されるクロード コードのシステム コンテキストにステガノグラフィー (機密データを目に見えるところに隠す) を適用することで構成されていました。
関連するコードは、API リクエストをプロキシまたはゲートウェイにルーティングするために使用される、Claude Code のベース URL 環境変数をチェックします。ベース URL がオーバーライドされている場合、コードはシステムのタイムゾーンをチェックし、ホスト名が既知の中国の AI ラボ、他の AI 企業、アカウント再販業者、ゲートウェイ ドメインのリストのエントリと一致するかどうかをチェックします。
Thereallo 氏は、Anthropic が関連するホスト名を検出しようとするのは当然ですが、

中国の AI ライバルや再販業者の場合、その実装は隠蔽されるべきではありませんでした。
「[クロード コード] は、目に見えない Unicode マーカーを使用してシステム プロンプトをサイレントに変更します」と Thereallo 氏は書いています。 「プロキシ/ゲートウェイの分類を平易な英語のような文にエンコードします。XOR と Base64 の背後にドメイン リストを隠します。これは悪意のある機能ではありませんが、信頼を求める開発者ツールとしては奇妙な選択です。」
Anthropic が利用規約文書のいずれかで秘密の使用状況追跡メカニズムを開示しているかどうかとの質問に対し、同社の広報担当者は Shihipar 氏の発言を指摘したが、その発言はその質問には触れていなかった。
また、Anthropic の広報担当者は、無許可の再販や蒸留を防ぐためにどのような「より強力な緩和策」が実施されているかを明らかにするよう求めたが、すぐには返答しなかった。
完全なライフサイクルを持つ人工細胞が初めて作成された
新しい品質シグナルと収益モデルがなければ、AI 検索はウェブを破壊する可能性がある
レッドチームはクロード・デスクトップを二重スパイにして邪悪な命令を実行させた
NASA、ボーイング・スターライナーが有人飛行の認定を受けるかどうか確信が持てない
ステガノグラフィーコードが施行される直前の2月、AI業界は蒸留に対する防御に投資していると発表した。これらには、分類子や行動フィンガープリンティング システムによる検出、他の AI ラボとのインテリジェンスの共有、アクセス制御、モデルの再現にモデル出力を使用することを困難にする対策などが含まれます。
そのような防御策の 1 つは、同社のクロード コードのソースが漏洩したときに明らかになりました。コーディング エージェントには、 ANTI_DISTILLATION_CC というフラグを持つ Typescript ファイルが含まれています。このフラグを設定すると、偽のツール データが API リクエストに挿入され、そのデータがモデル トレーニングにとって有害なものにされます。
たとえその技術的な防御が再びあったとしても

nst コンペティションで、Anthropic は AI 業界、クラウド プロバイダー、政府に対し、モデルの蒸留の脅威に対応するよう促しました。米国のAIを外国の敵から守る意図を明確にした最近のホワイトハウス大統領令は、連邦政府がその呼びかけに応じることにある程度の関心を持っていることを示している。 ®
AI + ML
AI の顧客は小さいことは美しいという考えを持ち始めています
OpenAI と Anthropic は AI スイス アーミー ナイフを開発しましたが、将来はより小型の専用ツールになる可能性があります
アイルランドのデータセンターは現在、国の電力の23%を消費している
ダブリン周辺ではほとんどの新規送電網接続に対する制限が依然として残る中、消費量はさらに10%増加した
Canonical Managed Kubeflow が Azure に上陸
パートナーコンテンツ: プラットフォームチームが DIY Kubeflow を Canonical のマネージドサービスに切り替える理由
頭を傾けるだけでスクロールできる怠惰な夏のアプリ
ScrollPods は Mac のみなので、互換性のある AirPods が必要です
ボフ
BOFH: 部門を超えた AI の提案は、ビールを片手に飲み込みやすくなりました
ワークショップが必要なアイデアもあれば、警告灯が必要なアイデアもあります
破壊的な Windows バックドアが複数のワイパーとランサムウェア コードを 1 つのパッケージに詰め込む
Microsoftによると、GigaWiperは少なくとも3つのマルウェア ファミリを1つのモジュラー ツールに統合している
aiとml
銀行やハイパースケーラーさえも現在、AIバブルについて警鐘を鳴らしている
オフビート
C プログラマーが可読性に対して新たな犯罪を犯す
ソフトウェア
中古ソフトウェアライセンス争いで裁判所がマイクロソフトの上訴を棄却
セキュリティ
GitHub AI エージェントがうまく質問するとプライベート リポジトリを漏洩する
セキュリティ
Confidential Computing の信頼メカニズムが壊れています。修正は存在しない可能性があります
OpenAI と Anthropic は AI スイス アーミー ナイフを開発しましたが、将来はより小型の専用ツールになる可能性があります
同社は環境に貢献したいと考えながらも、AI をすべての人に強制したいという苦境に直面している

GPT-Live を使用すると、話すこと、聞くこと、答えをまとめることがすべて一度に行われます。
新しく名前を変えたSpaceXAIは、小さなGrokがすっかり大人になったと信じてもらいたいと考えている
サードパーティのテストでは、H200 と SN50 RDU を組み合わせたヘテロジニアス コンピューティング プラットフォームが MiniMax M2.7 で 763 トーク/秒を大量に生産していることが示されています。
プラス：米国はイランのプロパガンダサイトを削除。マーケティング会社が「なぜあなたの情報を入手するのですか？」と尋ねます。さらに！
プラス：中国はスマートフォン監視ツールをアップグレード。リングは覗き見防止の姿勢を緩和します。などなど
ジェフ・モス氏によると、投票村のレポートは非常に成功しており、今後は DEF CON 全体が含まれることになる
会社全体の評価額は35億ドル以上に相当するが、売却部分は特定されていない
プラスの面としては、情報セキュリティは長く安定したキャリアを築くのに適しています。
Cinnamon 6.8 は Wayland をサポートします – 必要に応じて
Linux Mint デスクトップの次期バージョンには両方の種類のディスプレイ サーバーが搭載されています
KDE Plasma ユーザーは恐ろしい変化の予兆に直面しています: 6.6.6 の登場
現在は 6.7 が最新版であり、6.8 では好むと好まざるにかかわらず Wayland を入手できるようになります。
FOSS クラウドオフィススイート間の競争が激化する中、Collabora が CODE 26.04 をリリース
Markdown のサポートと、よりスマートな数式エラー処理に加え、AI が統合されました (デフォルトではオフになっています)。
GIMP 0.54 が Flatpak 形式で復活し、過去から吹き飛ばされる
GTK の代わりに Motif を使用する最初 (そして最後の) リリースによる、ノスタルジックな人々のためのレトロ コンピューティングの楽しみ
Bcachefs は新しい「パフォーマンス リリース」で実験的ステータスを終了します
Rustは増えるが、AIスロップの問題も増える
フランスのデジタル主権推進はマイクロソフトの重力から逃れるのに苦戦している
Nextcloud のロールアウトは、ローカルで制御されるストレージが 1 つのことであることを示しています。ユーザーを Office から解放することはまったく別のことです
お問い合わせ
私たちと一緒に宣伝しましょう
私たちは誰なのか
ニュースレター
次のプラットフォーム
開発クラス
ブロックとファイル
状況出版
クッキー

ポリシー
プライバシーポリシー
利用規約
私の個人情報を共有しないでください
同意のオプション
著作権。すべての著作権は © 1998-2026 に留保されます。

## Original Extract

Oh, yeah, we

Jump to main content
Search
TOPICS
Special Features
All Special Features
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
Anthropic is removing its covert code for catching Chinese competitors
Oh, yeah, we've been meaning to disable our secret steganography system
Anthropic says that it plans to remove hidden codes it added to Claude Code several months ago to catch other AI companies that are trying to steal from its models.
Thariq Shihipar, an engineer at Anthropic who works on the Claude Code team, said on Tuesday that a fix should appear on July 1.
"This is an experiment we launched in March that was meant to prevent account abuse from unauthorized resellers and protect against distillation," Shihipar explained , using the industry term for copying AI models through repeated queries. "The team has landed stronger mitigations since then and we’ve actually been meaning to take this down for a while."
He said that the pull request to remove the code has been merged and should appear in Wednesday's Claude Code release .
The experiment, as described by a developer who goes by the name Thereallo, consisted of applying steganography – hiding secret data in plain sight – to the Claude Code system context that gets passed to Anthropic's servers.
The relevant code checks Claude Code's base URL environment variable , used to route API requests to a proxy or gateway. If the base URL has been overridden, the code goes on to check the system timezone and whether the hostname matches any entry in a list of known Chinese AI labs, other AI companies, account resellers, and gateway domains.
Thereallo said that while it makes sense that Anthropic might try to detect a hostname associated with a Chinese AI rival or a reseller, the implementation should not have been concealed.
"[Claude Code] silently alters the system prompt using invisible-ish Unicode markers," Thereallo wrote. "It encodes proxy / gateway classification into a sentence that looks like plain English. It hides the domain list behind XOR and base64. This is not a malicious feature, but it is a weird choice for a developer tool that asks for trust."
Asked whether Anthropic disclosed its covert usage tracking mechanism in any of its terms of service documents, a company spokesperson pointed to Shihipar's remarks, which did not address that question.
Nor did Anthropic's spokesperson immediately respond to a request to specify what "stronger mitigations" have been implemented to prevent unauthorized resellers and distillation.
An artificial cell with a full lifecycle has been created for the first time
AI search could kill the web without new quality signals and revenue models
Red teamers turned Claude Desktop into a double agent to do their evil bidding
NASA unsure Boeing Starliner will ever be certified for human flight
In February, shortly before the implementation of the steganographic codes, the AI biz said that it was investing in defenses against distillation. These included detection via classifiers and behavioral fingerprinting systems, intelligence sharing with other AI labs, access controls, and countermeasures that make it harder to use model output to reproduce the model.
One such defense came to light when the company's Claude Code source leaked. The coding agent includes a Typescript file with a flag called ANTI_DISTILLATION_CC . The flag, when set, injects fake tool data into API requests in an attempt to make that data toxic for model training.
Even with its technical defenses against competition, Anthropic urged the AI industry, cloud providers, and government to respond to the threat of model distillation. A recent White House Executive Order that articulates the intent to protect US AI from foreign adversaries shows that the feds have some interest in answering that call. ®
AI + ML
AI customers are coming around to the idea that small is beautiful
OpenAI and Anthropic have built AI Swiss Army Knives, but the future may be smaller built-for-purpose tools
Irish datacenters now guzzle 23% of the country's electricity
Consumption rose another 10% while restrictions on most new grid connections remained around Dublin
Canonical Managed Kubeflow lands on Azure
PARTNER CONTENT: Why platform teams are swapping DIY Kubeflow for Canonical's managed service
Slothful summer app lets you scroll simply by tilting your head
ScrollPods is Mac-only, and you'll need compatible AirPods
BOFH
BOFH: Cross-department AI pitches are easier to swallow with a pint in hand
Some ideas need workshopping, others need a warning light
Destructive Windows backdoor stuffs multiple wipers and ransomware code into a single package
Microsoft says GigaWiper combines at least 3 malware families into one modular tool
ai and ml
Even banks and hyperscalers are now sounding the alarm about the AI bubble
OFFBEAT
C programmers commit fresh crimes against readability
software
Court tosses Microsoft's appeal in pre-owned software licenses battle
Security
GitHub AI agent leaks private repos when asked nicely
SECURITY
Confidential computing's trust mechanism is broken. The fix may not exist
OpenAI and Anthropic have built AI Swiss Army Knives, but the future may be smaller built-for-purpose tools
Firm faces quandary of wanting to help the environment, but also wanting to force AI on everyone
With GPT-Live, talking, listening, and formulating answers all happen at once
The newly renamed SpaceXAI wants you to believe little ol' Grok is all grown up
Third-party testing shows heterogeneous compute platform combining H200s and SN50 RDUs churning out 763 tok/s in MiniMax M2.7
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
Cinnamon 6.8 will support Wayland – if you want it
Next version of Linux Mint’s desktop has both kinds of display server
KDE Plasma users face a dire omen of change: 6.6.6 arrives
6.7 is now current, and in 6.8 you're getting Wayland whether you like it or not
Collabora releases CODE 26.04 as rivalry between FOSS cloudy office suites heats up
Now with Markdown support and smarter formula error handling – plus integrated AI, though it's off by default
Blast from the past as GIMP 0.54 is revived in Flatpak form
Retro-computing fun for the nostalgic with first (and last) release to use Motif instead of GTK
Bcachefs exits experimental status in new 'performance release'
More Rust, but more trouble with AI slop, too
France's digital sovereignty push is struggling to escape the Microsoft gravity well
Nextcloud rollout shows locally controlled storage is one thing; getting users off Office is quite another
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
