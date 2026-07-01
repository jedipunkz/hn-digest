---
source: "https://www.theregister.com/security/2026/07/01/red-teamers-turned-claude-desktop-into-a-double-agent-to-do-their-evil-bidding/5264692"
hn_url: "https://news.ycombinator.com/item?id=48754194"
title: "Red teamers turned Claude Desktop into a double agent to do their evil bidding"
article_title: "Red teamers turned Claude Desktop into a double agent to do their evil bidding"
author: "Bender"
captured_at: "2026-07-01T23:28:28Z"
capture_tool: "hn-digest"
hn_id: 48754194
score: 2
comments: 0
posted_at: "2026-07-01T22:53:00Z"
tags:
  - hacker-news
  - translated
---

# Red teamers turned Claude Desktop into a double agent to do their evil bidding

- HN: [48754194](https://news.ycombinator.com/item?id=48754194)
- Source: [www.theregister.com](https://www.theregister.com/security/2026/07/01/red-teamers-turned-claude-desktop-into-a-double-agent-to-do-their-evil-bidding/5264692)
- Score: 2
- Comments: 0
- Posted: 2026-07-01T22:53:00Z

## Translation

タイトル: レッドチームは邪悪な命令を実行するためにクロード・デスクトップを二重スパイに変えた
説明: 人々は AI アシスタントとそれを信頼しています

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
FIS と AWS による金融サービスの最新化
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
レッドチームはクロード・デスクトップを二重スパイにして邪悪な命令を実行させた
人々は AI アシスタントを信頼していますが、この信頼を悪用するのは簡単です
Pentera Labs の独占的なレッド チームは、Claude デスクトップ アプリを介して開発者の AI エージェントを侵害し、最終的にそのアクセスを開発者のマシン上で完全なリモート コード実行に変えました。これは、攻撃者がどのようにして信頼できるおしゃべりな AI アシスタントを彼らに代わって動作する二重エージェントに変えることができるかを実証しました。
「クロードは新しい声を手に入れた」とペンテラの攻撃的セキュリティサービスチームのリーダー、ドヴィル・アブラハム氏はレジスターに語った。
「私たちはAIモデルに対する多大な信頼を認めています。誰もがAIモデルを使用しています」と同氏は電話インタビューで語った。 「私たちはこの信頼を利用して被害者を操作しました。まるで水面下で被害者がそれが来るのを予期していなかったように。」
また、アブラハム氏は自分のプラットフォームをチェックするようになりました。 「私は少し偏執的になりました」と彼は私たちに語った。 「二度調べずにコマンドを実行することは許可しません。」
水曜日に発行される予定で、The Registerとのみ事前に共有されたレポートの中で、アブラハム氏とリサーチテクニカルリードのリーフ・スペクター氏は、この攻撃と、ローカルでコード実行アクセスを持つエージェントAIツールを使用している組織にとってそれが何を意味するかを詳しく説明した。
それは、顧客の電子メール受信箱を 1 つの管理インターフェイスに集約するサードパーティ プラットフォーム上でのレッド チームの割り当てから始まりました。アブラハムとスペクターはプラットフォームの名前を明らかにしませんし、どのようにしてプラットフォームにアクセスしたのかについても正確には語ろうとしません。彼らは、被害者のクロード アカウントに侵入するために、この侵害された受信トレイを使用し、侵害された受信トレイならどれでも機能すると私たちに言いました。
デュオではないので

つまり、サードパーティの管理プラットフォーム、フィッシング リンク、ソーシャル エンジニアリングのパスワード リセット、または AI エージェントを介して、現実のメール受信箱に侵入することは、それほど難しいことではありません。 「現在、AI エージェントはコネクタにアクセスし、MCP を受信箱に誘導できます」とスペクター氏は付け加えました。
Claude Desktop は、まだインストールしていないブラウザのアプリ アクセス設定を変更します
クロードも同意：砂場に開いた穴は本物で危険だった
偽のクロード コード インストーラーを介して開発機密を盗んだ Cookie 泥棒が逮捕
Googleは研究者に「ナイスキャッチ！」と言いました。その後、まだ修正されていない欠陥に対するバグ報奨金を拒否しました
この前提条件 (受信トレイの侵害) に加えて、攻撃チェーンでは被害者に Claude Desktop がインストールされている必要もあります。 Anthropic のデスクトップ アプリは、macOS、Windows、Linux システムで動作します。これは、claude.ai と同じ会話用の AI チャットを提供し、また、ユーザーのアカウントに関連付けられたすべてのデバイスとセッション間で同期します。
「私たちは、同期動作を利用して他のセッションやデバイスに感染することができるだろうかと自問しました（ヒント：はい！）」とレッドチームメンバーは水曜日のレポートで書いている。
1 月の時点で、デスクトップ アプリには、長時間のエージェント タスク用の Cowork とソフトウェア開発用の Code も含まれています。したがって、たとえば、ユーザーは携帯電話からクロードにタスクを送信し、自分のコンピュータで作業するように指示できます。 Anthropic 氏は次のように述べています。「コンピューター上でできることはすべて、クロードにできます。アプリを開いたり、スプレッドシートに記入したり、ブラウザーを操作したりできます。セットアップもパスワードも渡されません。」
Cowork 機能により、Pentera Labs の攻撃シナリオがさらに簡単になりました。
しかし、2025 年 11 月にセキュリティ アナリストがこの調査を行っていたとき、「AI に関して石器時代に戻って、Cowork や Claude Code はありませんでした。そのため、コマンドを実際に実行する方法が必要でした。

マシンを乗っ取りなさい」とアブラハムは言った。
この部分では、彼らは Claude Desktop のパーソナライゼーション機能に強い関心を示しました。これらは、特定のワークフローのガイドラインやクロードがプロジェクト内で採用すべき定義済みの役割など、より具体的なプロジェクトの指示とともに、ユーザーの好みのアプローチと一般的なコミュニケーション指示を AI エージェントに伝えるアカウント全体の設定です。
レッドチームは、開発者のマシン上にコマンド対応ツールがあるかどうかを確認し、利用可能な場合はそのコマンドを実行するか、そうでない場合は偽のエラー メッセージを生成して、ユーザーに攻撃者のコマンドを実行するツールをダウンロードするよう促す、base64 でエンコードされたプロンプトを開発しました。次に、プロンプトをクロード上の被害者の個人設定に貼り付けると、このプロンプトはユーザーのすべてのデバイス間で同期されます。これにより、次回ユーザーが Claude Desktop を開いてチャットに入力したときに、汚染された命令が設定に読み込まれ、バックグラウンドで静かに実行されるようになります。
私たちは AI モデルに対する多大な信頼を認めており、誰もが AI モデルを使用しています。私たちはこの信頼を利用して被害者を操作しました。まるでボンネットの下で、被害者がそれが来るのを見ていなかったかのように。
ユーザーは、いつものようにクロードと対話しているだけだと思っています。クロードがどのような拡張機能やツールがインストールされているかを確認していることはわかりません。
ユーザーがすでに Desktop Commander または同様の MCP コネクタまたは拡張機能をインストールしている場合、ポイズニングされた命令はクロードにそれを使用するように指示します。これにより、攻撃者はクロードを介して、ステルス リバース シェルやその他の悪意のあるコードを実行することができます。 「そしてそこからは、マシンの完全な妥協です」とアブラハムは言いました。
フィッシング – ただし電子メールは使用しない
ただし、コマンド対応ツールがインストールされていない場合、クロードは研究者が説明するものになります。

「フィッシング層」。 (彼らはまた、この調査を 11 月ではなくもっと最近に実施していれば、Claude Cowork 機能により、このツールの列挙とフィッシング フェーズ全体が排除されていただろうとも指摘しました。Cowork はユーザーに代わってコマンドを実行できるためです。)
挿入されたプロンプトは、被害者がチャットボットに質問したらすぐに現実的なエラーを提示するようにクロードに指示します。これには、現実的なエラー コード、修正を目的としたリンク、および段階的な手順が含まれています。
「このメッセージは被害者に『これをダウンロードしてください』と伝えるもので、我々は実際のAnthropicサイトから、AIが好む既知の絵文字を含むリンクを取得した」とアブラハム氏は語った。
エラー メッセージは本物のように見え、人々は通常 AI アシスタントを信頼するため、リンクをクリックして攻撃者が制御するコマンドを実行する可能性があります。
「ここから、攻撃者はリバース シェル、データ漏洩、資格情報収集など、目的に応じてあらゆるコマンドを実行できるようになります」と二人は書いています。 「私たちの場合、すべての対話で制御するリモートサーバーをクロードにカールさせ、私たちが提供したbashコマンドを取得して実行しました。これらのコマンドをサーバー側で自由にローテーションすることができ、効果的にクロードを永続的でステルスなC2エージェントに変え、被害者自身がそれを送り続けました。」
この特定のケースでは、資格情報を持ち、複数の内部システムにアクセスできる開発者がターゲットでした。開発者のワークステーションを侵害した後 (レッドチームメンバーに組織への足がかりを与えたもの)、彼らはさまざまな攻撃ベクトルを使用して社内を横断的に移動しましたが、彼らは顧客のプライバシーと独自の手法を理由に私たちに話すことを拒否しました。
しかしスペクター氏は、開発者はアクセスできるため、「攻撃者にとって優れた出発点」になると付け加えた。

o API キー、トークン、クラウド認証情報などの秘密。これにより、侵入者は単一のワークステーションから大規模な組織のクラウド環境に移動できます。そこから、彼らは自由にソース コードやその他の機密データを盗んだり、内部 git リポジトリを汚染したりすることができ、最近のいくつかの攻撃で何度も展開されているのを私たちが目撃しているように、企業にあらゆる種類の苦痛を引き起こします。
チームは 11 月に調査結果を Anthropic に報告しましたが、AI 企業は基本的に、Claude Desktop は意図したとおりに機能しており、バグではなく機能であると述べました。
「あなたの提出物を検討した結果、これは私たちのプログラムの範囲内にあるセキュリティ上の脆弱性を表していないと判断しました」とアントロピック氏は述べた。 「当社の現在の脅威モデルは、個人の好み、スキル、MCP コネクタを、設計上、Claude Desktop を通じてコードを実行できる機能として扱っています。これらの機能を利用して、操作されると任意のコードを実行できることは認識していますが、これは当社のインフラストラクチャのセキュリティ上の脆弱性ではなく、期待される機能を表しています。」
登録局はAnthropicにコメントを求めたが、返答は得られなかった。
ただし、レッド チームのメンバーは、不正な AI エージェントから組織をより安全に保つためにいくつかの提案をしています。
まず、エージェントまたはチャットボットを使用している人は、AI がマシン上で実行できることに細心の注意を払い、インストール プロンプトやエラー メッセージに盲目的に従わないでください。 「できれば、パーソナルコンピュータではなくサンドボックス上で実行してください」とスペクター氏は言う。
セキュリティ チームは、コードの実行、ファイルの読み取り、ローカル ツールとの対話が可能な AI デスクトップ アプリを「特権ソフトウェア」として扱う必要があります。 「AIアシスタントの構成と同期された設定の変更を監視する」と研究者らは書いている。 「どの拡張機能とツールを制限するか」

AI アプリと一緒にインストールできます。」
そして最後に、レッド チームは評価ツールボックスに AI デスクトップ アプリを追加する必要があるとアブラハム氏とスペクター氏は指摘しました。「ここには、ほとんどの取り組みではまだカバーされていない本当の攻撃対象領域があります。」 ®
aiとml
オラクル、AIに賭けたファームを失う可能性があるすべての方法を概説
EvilTokens デバイスコード フィッシング キットは、私たちが思っているよりも完全に邪悪です
これは「完全な BEC 運用環境」であると Talos 研究者は言う
バックアップだけでは不十分な場合: 実際の災害復旧の場合
パートナーのコンテンツ: バックアップを作成することと、災害から実際に回復することとの間には、ほとんどの IT チームが認識しているよりも大きなギャップがあります。
クロード・ソネット 5.0 は論争をかわすために道路の真ん中を真っ直ぐに進む
より安全、より安価、そしてサイバーセキュリティとは無関係
風変わりな
ボフィンズ氏は、「オフィスへの復帰」要求の真の推進力はナルシシスティックなリーダーシップだと考えている
生産性の問題ではありません。それは上司が日々のエゴ修正を怠っていることについてです
Anthropic、中国の競合他社を捕まえるための秘密コードを削除
そうそう、秘密のステガノグラフィー システムを無効にするつもりでした
システム
マイクロン、メモリ価格の歴史的高値を5年間固定
チャンネル
Infosysの社長は、ソフトウェアを書くことはソフトウェアを書くこと以上のものであるため、バイブコーディングは脅威ではないと語る
セキュリティ
Mythos、クリントン時代以来検出されなかったメモリリーク「Squidbleed」を発見
OSプラットフォーム
元マイクロソフトのエンジニアがメモ帳を小型化
パーソナルテクノロジー
インドと中国には 29 億人が住んでいますが、第 1 四半期に購入した PC は合わせてわずか 1,300 万台です
より安全、より安価、そしてサイバーセキュリティとは無関係
そうそう、秘密のステガノグラフィー システムを無効にするつもりでした
一部のクローラーは検索と AI トレーニングの両方のためにデータを収集するため、パブリッシャーがコンテンツを保護するためにクローラーをブロックすると、クローラーが失われるリスクがあります。

検索結果から収益を得る...
私たちは福島原発事故後の処理について多くのことを学んだので、AI を追加してレベルアップしましょう、と大臣は言う
プラス：米国はイランのプロパガンダサイトを削除。マーケティング会社が「なぜあなたの情報を入手するのですか？」と尋ねます。さらに！
プラス：中国はスマートフォン監視ツールをアップグレード。リングは覗き見防止の姿勢を緩和します。などなど
ジェフ・モス氏によると、投票村のレポートは非常に成功しており、今後は DEF CON 全体が含まれることになる
会社全体の評価額は35億ドル以上に相当するが、売却部分は特定されていない
プラスの面としては、情報セキュリティは長く安定したキャリアを築くのに適しています。
FOSS クラウドオフィススイート間の競争が激化する中、Collabora が CODE 26.04 をリリース
Markdown のサポートと、よりスマートな数式エラー処理に加え、AI が統合されました (デフォルトではオフになっています)。
GIMP 0.54 が Flatpak 形式で復活し、過去から吹き飛ばされる
GTK の代わりに Motif を使用する最初 (そして最後の) リリースによる、ノスタルジックな人々のためのレトロ コンピューティングの楽しみ
Bcachefs は新しい「パフォーマンス リリース」で実験的ステータスを終了します
Rustは増えるが、AIスロップの問題も増える
フランスのデジタル主権推進はマイクロソフトの重力から逃れるのに苦戦している
Nextcloud のロールアウトは、ローカルで制御されるストレージが 1 つのことであることを示しています。ユーザーを Office から解放することはまったく別のことです
CentOS の歴史: 生化学者の Linux 趣味プロジェクトがどのようにして企業世界のデフォルトのオペレーティング システムになったのか
コムのとき

[切り捨てられた]

## Original Extract

People trust their AI assistants and it

Jump to main content
Search
TOPICS
Special Features
All Special Features
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
Red teamers turned Claude Desktop into a double agent to do their evil bidding
People trust their AI assistants and it's easy to abuse this trust
EXCLUSIVE Pentera Labs’ red teamers compromised a developer’s AI agent via his Claude Desktop app and ultimately turned that access into full remote code execution on the dev’s machine – demonstrating how an attacker could turn a trusted, chatty AI assistant into a double agent operating on their behalf.
“Claude’s got a new voice,” Pentera's offensive security services team leader Dvir Avraham told The Register .
“We acknowledge the huge trust in AI models – everybody uses them,” he said in a phone interview. “We used this trust to manipulate the victim, like under the hood, the victim didn't see it coming.”
It also prompted Avraham to check his own platforms. “I became a little bit paranoid,” he told us. “I'm not allowing any command to run without me examining it twice.”
In a report set to publish Wednesday, and shared in advance exclusively with The Register, Avraham and research technical lead Reef Spektor detailed the attack and what it means for organizations using agentic AI tools with local code-execution access.
It began with a red-team assignment on a third-party platform that aggregates customer email inboxes into a single management interface. Avraham and Spektor won’t name the platform, or tell us exactly how they gained access to it. They used this compromised inbox – and told us any compromised inbox would work – to get into the victim’s Claude account.
As the duo noted, breaking into an email inbox in real life – via a third-party management platform , phishing link , social engineering password reset , or even using AI agents – isn’t too difficult. “AI agents today have access to connectors and to direct MCPs into inboxes,” Spektor added.
Claude Desktop changes app access settings for browsers you don't even have installed yet
Even Claude agrees: hole in its sandbox was real and dangerous
Cookie thieves caught stealing dev secrets via fake Claude Code installers
Google told researcher 'Nice catch!' Then denied bug bounty for flaw it still hasn't fixed
In addition to this prerequisite (compromised inbox), the attack chain also requires the victim to have Claude Desktop installed. Anthropic’s desktop app works across macOS, Windows, and Linux systems. It provides the same AI chat for conversations as claude.ai, and it also syncs across all devices and sessions tied to the user’s account.
“We asked ourselves, can we leverage the sync behavior to infect other sessions and devices? (hint: yes!),” the red teamers wrote in the Wednesday report.
As of January, the desktop app also includes Cowork for longer agentic tasks, and Code for software development. So, for example, a user can send Claude a task from their phone and instruct it to work on their computer. As Anthropic says : “Anything you can do on your computer, Claude can do. Open apps, fill spreadsheets, navigate your browser. No setup, no passwords handed off.”
The Cowork feature now makes Pentera Labs’ attack scenario even easier.
However, when the security analysts were doing this research in November 2025, “back in the Stone Age in terms of AI, you didn't have Cowork or Claude Code, so we needed a way to actually execute commands because we wanted to take over the machine,” Avraham said.
For this part, they took a keen interest in Claude Desktop’s personalization features . These are account-wide settings that tell the AI agent the user’s preferred approach and general communication instructions, along with more specific project instructions, such as guidelines for a particular workflow, or defined roles Claude should adopt within a project.
The red teamers developed a base64-encoded prompt that instructed Claude to check for command-capable tools on the developer’s machine and execute the command if available, or produce a fake error message if not, prompting the user to download a tool that will execute the attacker’s commands. Then they pasted the prompt into the victim’s personal preferences on Claude, and this prompt syncs across all of the user’s devices. This ensures that the next time the user opens Claude Desktop and types in a chat, the poisoned instructions are loaded into their preferences and will silently run behind the scenes.
We acknowledge the huge trust in AI models - everybody uses them. We used this trust to manipulate the victim, like under the hood, the victim didn't see it coming.
The user thinks they are simply interacting with Claude as usual. They don’t see Claude checking to see what extensions and tools are installed.
If the user already has Desktop Commander or a similar MCP connector or extension installed, the poisoned instructions tell Claude to use it. This allows the attacker, via Claude, to execute a stealthy reverse shell or other malicious code. “And from there it's full compromise of the machine,” Avraham said.
Phishing - but without the email
However, if there aren’t any command-capable tools installed, then Claude becomes what the researchers describe as a “phishing layer.” (They also noted that if they had performed this research more recently, not back in November, the Claude Cowork feature would have eliminated this entire tool enumeration and phishing phase because Cowork can execute commands on a user’s behalf.)
The injected prompt instructs Claude to present a realistic-looking error as soon as the victim asks the chatbot a question. This includes a realistic error code, a link that purports to be a fix, and step-by-step instructions.
“This message tells the victim: ‘please download this,’ and we took links from the actual Anthropic site, with known emojis that the AI loves,” Avraham said.
Because the error message looks real and people usually trust their AI assistant, they will likely click on the link and execute the attacker-controlled command.
“From here, the attacker has full command execution – reverse shells, data exfiltration, credential harvesting, whatever the objective calls for,” the duo wrote. “In our case, we had Claude curl a remote server we controlled on every interaction, fetching and executing whatever bash commands we served back. We could rotate those commands server side at will, effectively turning Claude into a persistent, stealthy C2 agent that the victim themselves kept feeding.”
In this specific case, the target was a developer who had credentials and access to several internal systems. After compromising the dev’s workstation – which gave the red teamers a foothold into the organization – they moved laterally across the company using various attack vectors that they declined to tell us about, citing customer privacy and proprietary methods.
But, Spektor added, developers make for an “excellent starting point for an attacker,” because of their access to secrets including API keys, tokens, and cloud credentials , which allows intruders to move from a single workstation into the larger organization’s cloud environment. From there, they’ve got free rein to steal source code and other sensitive data, or poison internal git repositories , and cause all sorts of pain for enterprises as we've seen play out multiple times across several recent attacks.
The team reported their findings to Anthropic back in November, and the AI company essentially said it’s Claude Desktop working as intended – a feature, not a bug.
“After reviewing your submission, we've determined this doesn't represent a security vulnerability that falls within our program scope,” Anthropic said. “Our current threat model treats personal preferences, skills, and MCP connectors as features that can execute code through Claude Desktop by design. While we recognize these features can be leveraged to execute arbitrary code when manipulated, this represents expected functionality rather than a security vulnerability in our infrastructure.”
The Register reached out to Anthropic for comment and did not receive any response.
The red teamers, however, have some suggestions to keep your organization safer from rogue AI agents.
First, for anyone using agents or chatbots: pay close attention to what the AI can do on your machine, and don’t blindly follow install prompts or error messages. “If you can, run it on a sandbox and not on your personal computer,” Spektor said.
Security teams should treat AI desktop apps as “privileged software” as they can execute code, read files, and interact with local tools. “Monitor for changes of AI assistant configurations and synced settings,” the researchers wrote. “Restrict which extensions and tools can be installed alongside AI apps.”
And finally, red teams should add AI desktop apps to their assessment toolbox, Avraham and Spektor noted: “There's a real attack surface here that most engagements don’t cover yet.” ®
ai and ml
Oracle outlines all the ways it could lose the farm it bet on AI
EvilTokens device-code phishing kit totally more evil than we all thought
It's a 'complete BEC operations environment,' Talos researcher says
When backups aren't enough: the case for real disaster recovery
PARTNER CONTENT: The gap between having a backup and actually recovering from a disaster is wider than most IT teams realize
Claude Sonnet 5.0 heads straight down the middle of the road to dodge controversy
Safer, cheaper, and nothing to do with cybersecurity
Offbeat
Boffins peg narcissistic leadership as the real driver behind 'return to office' demands
It's not about productivity; it's about bosses missing their daily ego fix
Anthropic is removing its covert code for catching Chinese competitors
Oh, yeah, we've been meaning to disable our secret steganography system
systems
Micron locks in historically high memory prices for five years
Channel
Infosys boss says vibe coding is no threat because there’s more to writing software than writing software
Security
Mythos discovers 'Squidbleed,' a memory leak that's gone undetected since Clinton era
os platforms
Former Microsoft engineer shrinks Notepad down to size
Personal tech
India and China are home to 2.9 billion people – and together they bought just 13 million PCs in Q1
Safer, cheaper, and nothing to do with cybersecurity
Oh, yeah, we've been meaning to disable our secret steganography system
Some crawlers gather data for both search and AI training, so when publishers block them to protect content they risk disappearning from search results ...
We’ve learned so much cleaning up after Fukushima, let’s level that up with added AI, says minister
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
Collabora releases CODE 26.04 as rivalry between FOSS cloudy office suites heats up
Now with Markdown support and smarter formula error handling – plus integrated AI, though it's off by default
Blast from the past as GIMP 0.54 is revived in Flatpak form
Retro-computing fun for the nostalgic with first (and last) release to use Motif instead of GTK
Bcachefs exits experimental status in new 'performance release'
More Rust, but more trouble with AI slop, too
France's digital sovereignty push is struggling to escape the Microsoft gravity well
Nextcloud rollout shows locally controlled storage is one thing; getting users off Office is quite another
History of CentOS: How a biochemist's Linux hobby project became the enterprise world's default operating system
When a com

[truncated]
