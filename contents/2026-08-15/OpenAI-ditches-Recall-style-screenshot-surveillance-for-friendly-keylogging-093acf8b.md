---
source: "https://www.theregister.com/ai-and-ml/2026/08/14/openai-ditches-recall-style-screenshot-surveillance-for-friendly-keylogging/5287618"
hn_url: "https://news.ycombinator.com/item?id=49306929"
title: "OpenAI ditches Recall-style screenshot surveillance for friendly keylogging"
article_title: "OpenAI ditches Recall-style screenshot surveillance for friendly keylogging"
author: "sbulaev"
captured_at: "2026-08-15T02:08:11Z"
capture_tool: "hn-digest"
hn_id: 49306929
score: 1
comments: 0
posted_at: "2026-08-15T02:07:06Z"
tags:
  - hacker-news
  - translated
---

# OpenAI ditches Recall-style screenshot surveillance for friendly keylogging

- HN: [49306929](https://news.ycombinator.com/item?id=49306929)
- Source: [www.theregister.com](https://www.theregister.com/ai-and-ml/2026/08/14/openai-ditches-recall-style-screenshot-surveillance-for-friendly-keylogging/5287618)
- Score: 1
- Comments: 0
- Posted: 2026-08-15T02:07:06Z

## Translation

タイトル: OpenAI は、フレンドリーなキーロギングのために Recall スタイルのスクリーンショット監視を廃止

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
2026 年クラウド インフラストラクチャ月間
FIS と AWS による金融サービスの最新化
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
OpenAI はフレンドリーなキーロギングのために Recall スタイルのスクリーンショット監視を廃止
「コンピュータ履歴」はクリックとタイピングを記録して ChatGPT の記憶を構築します
コンピューター上で行うことをすべて記録し、それらの記録を OpenAI に送信し、より多くの ChatGPT トークンを使用し、インジェクションを促す脆弱性を高めたい場合は、OpenAI が役に立ちます。
これは コンピューター履歴 と呼ばれるもので、アプリや Web サイト間でのコンピューターの操作をタイムライン上に整理された思い出として記録するオプトイン方式です。
なぜそうしたいのですか?おそらく、スクリーンショットを使用して同様の履歴をまとめた Computer History の前身である Chronicle が少し押し付けがましすぎると感じたかもしれませんが、OpenAI のサーバーに短時間アクセスするだけで、入力イベントを記録し、暗号化されずにローカルに 48 時間 (またはそれ以上) 保存するという Computer History のアプローチは気にしません。
OpenAI のドキュメントにある次の警告は気にしていないかもしれません。「コンピュータ履歴ファイルには機密情報が含まれている可能性があります。これらはコンピュータ履歴によって暗号化されておらず、macOS ユーザーとして実行されている他のプログラムがアクセスできる可能性があります。」
おそらく、OpenAI の Codex と GPT Work をコンピュータで実行したことがある方は、コンピュータのアクティビティをメモリ ファイルに保存し、それらのインタラクションをタイムラインに配置すると、ChatGPT の応答が改善され、自動化の機会が表面化し、以前の作業の再開が容易になるという考えにすでに納得されているのではないでしょうか。
Computer History は、端的に言えば、キーロギングおよびイベント キャプチャ システムです。
メガネ型カメラ、ナンバープレートリーダー、監視カメラが登場する以前の時代がありました。

資本主義、そしてそのような覗き見がプライバシー擁護派からの抗議を引き起こした可能性があるときに警察のドローンが使用されました。しかしハイテク業界は、監視を自社の顧客に委託することで、パノプティコンが個人的な選択になった場合に抗議することが難しくなる可能性があることに気づいた。
「Computer History は、許可されたアプリや Web サイトからのインタラクション イベント ストリームを作成します」と OpenAI のドキュメントには説明されています。 「イベントには、クリック、タイピング、キーボード ショートカット、アプリの切り替え、macOS がアクセシビリティ システムを通じて公開するコンテキストなどが含まれます。コンピュータ履歴は、これらのイベントを定期的にテキストの概要とローカル メモリ ファイルに変換します。」
AI 業界では、Computer History では画面イメージ、マイク入力、システム オーディオがキャプチャされないことを強調しています。また、プライベート モードのブラウジングもキャプチャしません。
デフォルトではオフになっているコンピューター履歴は、macOS 上の ChatGPT デスクトップ アプリで ChatGPT Pro、Business、および Enterprise ユーザーが利用できます。 Pro ユーザーは個別に有効にすることができます。 Business および Enterprise ユーザーは、管理者による承認が必要です。現在、欧州経済領域 (EEA)、スイス、英国では利用できません。また、API キーまたは Amazon Bedrock 経由で ChatGPT にアクセスするユーザーも利用できません。
OpenAI は、コンピューター履歴ユーザー​​がアプリや Web サイトでのユーザー アクティビティを許可なく取得することに何らかの懸念がある場合、サービスの一時停止を希望する可能性があると示唆する状況があります。
おそらく法的リスクを認識して、同社は「事前に明示的な同意がない限り、他の人とのコミュニケーション中はこの機能をオフにしてください」とアドバイスしている。 「一時停止するか、機密性の高い健康情報、財務情報、個人情報を含むアプリを除外することを検討してください。」
コンピューター履歴のインタラクション イベントは、ChatGPT および Codex によって削除されるまで、最大 48 時間ローカルに保存されることになっています。
ただし、イベント

そうですね、メモリを生成するために OpenAI サーバーに送信され、それらは長期間ローカルに保存され、OpenAI に戻される将来のチャットで使用される可能性があります。
OpenAI 広告サービスは、キャンペーンを一時停止した後、顧客に最大 1 日分の料金を請求できる
ザック、ミューズ・グリマーと無差別級ラマ劇を再燃
AnthropicがFableの鎖を緩める中、OpenAIはAstraのセキュリティを強化することを約束
OpenAI、不正エージェントの群れがハグフェイスハッキングよりも少し先んじて行動していたことを明らかに
「OpenAIは、法律で義務付けられている場合を除き、処理後にこれらのイベントファイルを保持せず、トレーニングに使用しません」と同社は述べている。
同社はチャットログの要求には反対してきたが、それでも法的手続きに応じてチャットログを提供してきた。
コンピューター履歴では、アクティビティの要約とメモリ データの作成中にトークンが使用されるため、コストが増加します。また、プロンプト インジェクションの攻撃対象領域も拡大します。
「コンピュータの履歴により、アプリや Web サイトのコンテンツからの即時挿入のリスクが高まります」と同社は述べています。 「たとえば、悪意のある命令を含む Web サイトにアクセスすると、ChatGPT や Codex がその命令に従う可能性があります。」
ただし、少なくとも、最近のアクティビティに関する優れたタイムラインが得られます。 ®
aiとml
Anthropicは、テキスト透かしスキームは重要でない単語に依存していると述べています
そして他の AI モデル メーカーも同様のものを導入すると予想されています
DeepSeek の革新的なハーネスはすべてをプラグインとして扱います
中国のAI研究所は前進を続ける一方、米国の研究所は守りを固める
Microsoft が約束したことのないバックアップ
スポンサー機能: M365 と Azure のデータは、ランサムウェアに対して思っているほど安全ではない可能性があります。現実を確認する時間
ShinyHunters の恐喝攻撃により、160 万の RingCentral アカウントのデータがダンプされる
コラムニスト
エージェントは私のレトロな技術を安全に再度使用できるようにし、テスターとしての真の価値を示しました

アイデアの
みんなでちょっとマッドサイエンティストになって、ソフトウェアが私たちの最も突飛な理論を検証できるかどうか見てみましょう
ロシアのミサイルはウクライナを標的にするためにNvidia AIチップを使用
キエフはロシアの兵器に外国シリコンを持ち込まないよう、より厳格な管理を望んでいる
セキュリティ
ディープフェイクのしゃっくりがデジタル証明書詐欺容疑者の正体を暴く
セキュリティ
デルタ航空の機内Wi-Fiを乗っ取ろうとした疑いのあるDEF CONディンガス
セキュリティ
暗号化されていないコピーが GitHub に到達した後、Mozilla が Firefox 署名キーを取り消す
データベース
深く埋もれていた 16 年前の SQLite のバグが昨年の Tailscale の停止を引き起こした
OSプラットフォーム
Linus Torvalds氏、AIのおかげで「巨大な」Linuxカーネルのアップデートが新たな常態になったと語る
そして他の AI モデル メーカーも同様のものを導入すると予想されています
中国のAI研究所は前進を続ける一方、米国の研究所は守りを固める
武器を持ったエージェントはデジタル侵入を動的災害に変える可能性があると専門家が警告
そしてそこに到達するためにあらゆる経済的困難を乗り越えるだろう
開発者はクアルコム買収後の不確実性を払拭するためにオープンソースコンパイラのリリースを待っている
プラス：米国はイランのプロパガンダサイトを削除。マーケティング会社が「なぜあなたの情報を入手するのですか？」と尋ねます。さらに！
プラス：中国はスマートフォン監視ツールをアップグレード。リングは覗き見防止の姿勢を緩和します。などなど
ジェフ・モス氏によると、投票村のレポートは非常に成功しており、今後は DEF CON 全体が含まれることになる
会社全体の評価額は35億ドル以上に相当するが、売却部分は特定されていない
プラスの面としては、情報セキュリティは長く安定したキャリアを築くのに適しています。
FOSS は Microsoft の独占を 1 つ打ち破りました。 20年間の失敗を経て、次の失敗をする時が来た
一言
GNOME は Windows のように見えることができ、Flashback は拡張機能なしで実行できます
新しい「シンプルタスクバー」はオプションですが、よりシンプルで安定した方法があります
x での Debian の最終リリースに黙祷を捧げてください

86-32
新しい Debian バージョンが 13.6 および 12.15 の形で FOSSland に登場
脆弱な Joomla Web サイトで拡張機能のバグを悪用し、10 点満点を獲得した悪者を発見
iCagenda、Balbooa Forms 拡張機能の欠陥は、世界中の 100 万のサイトを支えるオープンソース CMS に影響を与える可能性があります
フレーム: 新しい X11 サーバー – アセンブリに直接実装
yserver、Phoenix、そしてもちろん XLibre、そして外れ値の Arcan に参加します
Cinnamon 6.8 は Wayland をサポートします – 必要に応じて
Linux Mint デスクトップの次期バージョンには両方の種類のディスプレイ サーバーが搭載されています
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

Jump to main content
Search
TOPICS
Special Features
All Special Features
Cloud Infrastructure Month 2026
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
OpenAI ditches Recall-style screenshot surveillance for friendly keylogging
'Computer History' records clicks and typing to build ChatGPT memories
If you want to record whatever you do on a computer, send those records to OpenAI, use more ChatGPT tokens, and increase your vulnerability to prompt injection, then OpenAI has something for you.
It's called Computer History , an opt-in way to record your computer interactions across apps and websites as memories organized on a timeline.
Why would you want to do so? Maybe you found Chronicle , the predecessor of Computer History which compiled similar histories using screenshots, a bit too intrusive but don't mind Computer History's approach – recording input events and storing them unencrypted locally for 48 hours (or more), with a brief visit to OpenAI's servers.
Maybe you're not bothered by the warning OpenAI includes in its documentation: "Computer History files can contain sensitive information. They are not encrypted by Computer History, and other programs running as your macOS user may be able to access them."
Perhaps, having given OpenAI's Codex and GPT Work the run of your computer, you're already sold on the suggestion that storing your computer activity in memory files and arranging those interactions in a timeline will improve ChatGPT responses, surface opportunities for automation, and make it easier to resume prior work.
Computer History is, to put it bluntly, a keylogging and event capture system.
There was a time before eyeglass cameras, license plate readers, surveillance capitalism, and police drones when such snooping might have provoked an outcry from privacy advocates. But the tech industry has found it can outsource surveillance to its own customers and in so doing make the panopticon harder to protest once it becomes a personal choice.
"Computer History creates an interaction-event stream from allowed apps and websites," OpenAI's documentation explains. "Events can include clicks, typing, keyboard shortcuts, app switches, and context that macOS exposes through its accessibility system. Computer History periodically turns these events into text summaries and local memory files."
The AI biz makes a point of noting that Computer History does not capture screen images, microphone input, or system audio. Nor does it capture private-mode browsing.
Off by default, Computer History is available for ChatGPT Pro, Business, and Enterprise users in the ChatGPT desktop app on macOS. Pro users can enable it individually; Business and Enterprise users need an admin to approve it. It's not available currently in the European Economic Area (EEA), Switzerland, or the United Kingdom or to those accessing ChatGPT via API key or Amazon Bedrock.
There are circumstances in which OpenAI suggests Computer History users might want to suspend the service if they have some scruples about capturing user activity in apps and websites without permission.
"Turn it off during communications with other people unless you have their prior express consent," the company advises, perhaps in acknowledgement of legal risk. "Consider pausing it or excluding apps that contain sensitive health, financial, or personal information."
Computer History interaction events are supposed to be saved locally for up to 48 hours before being deleted by ChatGPT and Codex.
Events, however, get sent to OpenAI servers to generate memories, and those may be stored locally for longer periods of time and may be used in future chats that get passed back to OpenAI.
OpenAI ad service can bill customers for up to one day after they pause campaigns
Zuck rekindles open weights Llama drama with Muse Glimmer
OpenAI pledges to add Astra security as Anthropic loosens Fable's leash
OpenAI reveals its rogue agent swarm went a little bit Borg ahead of Hugging Face hack
"OpenAI does not retain those event files after processing unless required by law and does not use them for training," the company says.
While it has opposed demands for chat logs, it has nonetheless provided chat logs in response to legal process.
Computer History adds cost because it uses tokens during the summarization of activities and the creation of memory data. It also expands the prompt injection attack surface.
"Computer History increases the risk of prompt injection from content in apps and websites," the company says. "For example, if you visit a website containing malicious instructions, ChatGPT or Codex might follow those instructions."
But at least you get a nice timeline of your recent activity. ®
ai and ml
Anthropic says text watermarking scheme relies on inconsequential words
And other AI model makers are expected to deploy something similar
DeepSeek's innovative harness treats everything as a plug-in
Chinese AI labs keep moving forward while US labs play defense
The backup Microsoft never promised you
SPONSORED FEATURE: Your M365 and Azure data might not be as safe as you think from ransomware; time for a reality check
1.6M RingCentral accounts' data dumped after ShinyHunters extortion attack
columnists
Agents made my retro tech safe to use again and showed their real value as testers of ideas
Let's all go a bit mad scientist and see if software can validate our wildest theories
Russian missile uses Nvidia AI chip to help target Ukraine
Kyiv wants tighter controls to keep foreign silicon out of Moscow's weapons
Security
Deepfake hiccup unmasks suspected digital certificate fraudster
security
DEF CON dingus suspected of trying to take over Delta in-flight Wi-Fi
SECURITY
Mozilla revokes Firefox signing key after unencrypted copy lands in GitHub
databases
Deeply buried 16-year-old SQLite bug caused last year's Tailscale outages
OS PLATFORMS
Linus Torvalds says AI has made 'huge' Linux kernel updates the new normal
And other AI model makers are expected to deploy something similar
Chinese AI labs keep moving forward while US labs play defense
Weaponized agents could turn digital intrusions into kinetic disasters, experts warn
And it'll jump through every financial hoop it can to get there
Developers await open source compiler release to dispel uncertainty following Qualcomm acquisition
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
FOSS smashed one Microsoft monopoly. After 20 years of failure, it's time to smash another
Word up
GNOME can look like Windows – and Flashback can do it without extensions
New 'Simple-taskbar' is an option, but there's a simpler, stabler way
A moment of silence, please, for the final release of Debian on x86-32
New Debian versions hit FOSSland in the form of 13.6 and 12.15
Baddies caught exploiting extensions bugs with perfect 10 scores on vulnerable Joomla websites
Flaws in iCagenda, Balbooa Forms extensions can impact open source CMS that powers a million sites worldwide
Frame: A new X11 server – implemented directly in assembly
Joins yserver, Phoenix, and of course XLibre – and outlier Arcan
Cinnamon 6.8 will support Wayland – if you want it
Next version of Linux Mint’s desktop has both kinds of display server
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
