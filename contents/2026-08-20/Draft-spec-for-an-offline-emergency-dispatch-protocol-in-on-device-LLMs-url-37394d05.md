---
source: "https://github.com/kapustin-i/offline-dispatch-protocol"
hn_url: "https://news.ycombinator.com/item?id=49374202"
title: "Draft spec for an offline emergency dispatch protocol in on-device LLMs url"
article_title: "GitHub - kapustin-i/offline-dispatch-protocol: Draft specification for an offline emergency dispatch protocol in on-device AI models · GitHub"
image: "https://opengraph.githubassets.com/fc6b60b9841820c4ab1fe036285cc552c762c50991321703c4396d17d254f5e3/kapustin-i/offline-dispatch-protocol"
author: "kapustin-i"
captured_at: "2026-08-20T13:39:04Z"
capture_tool: "hn-digest"
hn_id: 49374202
score: 2
comments: 0
posted_at: "2026-08-20T13:16:09Z"
tags:
  - hacker-news
  - translated
---

# Draft spec for an offline emergency dispatch protocol in on-device LLMs url

- HN: [49374202](https://news.ycombinator.com/item?id=49374202)
- Source: [github.com](https://github.com/kapustin-i/offline-dispatch-protocol)
- Score: 2
- Comments: 0
- Posted: 2026-08-20T13:16:09Z

## Translation

タイトル: オンデバイス LLM のオフライン緊急派遣プロトコルの仕様草案 URL
記事タイトル: GitHub - kapustin-i/offline-dispatch-protocol: オンデバイス AI モデルにおけるオフライン緊急派遣プロトコルの仕様草案 · GitHub
説明: オンデバイス AI モデルにおけるオフライン緊急派遣プロトコルの仕様草案 - kapustin-i/offline-dispatch-protocol

記事本文:
GitHub - kapustin-i/offline-dispatch-protocol: オンデバイス AI モデルにおけるオフライン緊急派遣プロトコルの仕様草案 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
カプースチン
/
オフラインディスパッチプロトコル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
11 コミット 11 コミット フォルダーとファイル
ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
オフライン緊急派遣プロトコル
あ

永久 DOI でアーカイブ: 10.5281/zenodo.22023184
これは何だろう。負傷して孤立した人が唯一残されたものであり、電話をかけるディスパッチャーが存在しない場合に、電話機のローカル モデルが実行しなければならないことに関する仕様草案。 16 の要件 (動作、アーキテクチャ、ハードウェア) と、誰も測定していないもののリスト。実装ではなく、命を救うという主張でもありません。
モデルは命令を構成しません。トリアージ ロジックはステート マシンであり、命令は既存のディスパッチ プロトコルから取得された固定セットです。要件9を参照してください。
反対意見が重要です。問題を提起してください。
ジャンプ先: 16 の要件 · 誰も測定していないもの
次の瞬間に何が起こるかはわかりません。原則として何もありません。大惨事が起こらなかったのは私たちの幸運です。
人はそれについて考える義務はありません。彼は何も起こらないかのように暮らしていますが、それはほとんど常に正しいのです。
デバイスを構築し、ルールを作成する人は、それについて考える義務があります。事前に行動できるのは彼らだけです。その瞬間、その人は何も変えることができません。彼のポケットの中にあったものはすべて彼のポケットの中にありました。
これがすべての安全工学の仕組みです。 「今日は衝突する」と考えて車をスタートさせる人はいません。いずれにしてもシートベルトは必須です。嵐のことを考えながらフェリーに乗る人は誰もいません。ベストは座席の下にあります。まれな突然の出来事に対する責任は常に個人ではなく、メーカーと規制当局に移ります。それはまさに、その人がその時点で準備ができていないからであり、準備ができている必要はないからです。
男性がコンクリート板の下に横たわっています。彼の手には、接続されていない電話があります。画面が光り、バッテリーが消耗し、暗闇になります。
電話はいつでもそこにあるからです。バスルームで、職場で、夕食時、夜の枕の上で。それはどこにでも人を追いかけます—そして

それはその瞬間にそこにあります。持ち主と同じように、無傷でも壊れても。
内部には、応急処置、物理学、低体温症への対処法を知っているモデルが入っています。どんな質問にも丁寧かつ的確に答えてくれます。
彼は間違った質問をするでしょう。彼は暗闇の中でショックを受け、負傷した。彼の思考は狭くなり、恐怖が彼を別の方向に向けています。
なぜ専用デバイスではなく電話なのか
誰も突然のことに備えていません。救助装置には先見性が求められます。購入し、持ち歩き、記憶し、充電しておく必要があります。そして、そのすべてを行う人は、通常、それが起こる人ではありません。いくつかの準備ができています。誰もがスラブの下に落ちてしまいます。
電話機では、先見性の要件が完全に削除されます。それはその人が準備したからではなく、決して手放さないからそこにあるのです。お風呂で、バスルームで、夜の枕の上で、仕事中、ハンドルを握っているとき、誰かとベッドの中で。救助装置としての使用を目的としたものではなく、あらゆる場所に設置されています。
したがって、その予測不可能な秒間では、ほとんどの場合、身体の 1 メートル以内、または身体に衝突することになります。瓦礫の下、横転した車の中、渓谷の底。
これが、専用デバイスではなく電話を使用する理由の答えです。電話機の方が優れているからではありません。専用のビーコンの方が、電力、アンテナ、耐久性など、すべての点で優れています。なぜなら、専用のビーコンは存在せず、電話が存在するからです。
もちろん、いつもではありません。別の部屋に放置したり、車から投げ捨てたり、押しつぶしたりすることもできます。つまり、すべてのケースではなく、大部分のケースです。しかし、他のオブジェクトはそのシェアに近づきません。
1 日に約 2,000 人が命を落としており、数分以内に助けが必要な状態で命を落としていますが、その結果は取り返しのつかないものであり、どこにも連絡することができません。災害時ではありません。高速道路、村、夜間、派遣サービスのない国などです。上限推定値は 10 倍です。
月6万。 730,000

一年。
オンデバイスのディスパッチャーが 100 人に 2 ～ 3 人の命を救うとしたら、電話による心肺蘇生指導の効果とほぼ同じで、年間 15,000 ～ 20,000 人の命が救われることになります。 7年間待ち：10万人以上。
桁違いの推定値。 4 つの手順はすべてソース記事に記載されています。誰でもそれを繰り返すことができます。
正確な統計はありません。誰も「孤立した意思決定の瞬間」をカテゴリーとしてカウントしません。推定値は 4 つのステップで乗算して構築され、それぞれをチェックできます。
1つ。院外心停止: 100,000 人年あたり 55 ～ 113 件。約46億人が必須の医療サービスを利用できないまま暮らしています。 3 億人以上が携帯電話の通信範囲外に住んでいます。この 1 つの診断だけでも、1 日に 500 人から 7,000 人が「連絡先がない」状態で亡くなっています。外傷、中毒、出産を加えると、狭い推定では 1 日に数千人、広い推定では 1 万人が発生します。
二。人を指導し、誰かが責任を負う認定システムが確立されるまでの時間: 7 ～ 15 年。技術的な理由ではありません。プロトタイプには数か月かかり、必要な品質のモデルが今日の携帯電話に適合します。その理由は規制上のものであり、責任を負う人がいないためです。
三つ。現実的な生存率の向上は 2 ～ 3 パーセント ポイントであり、倍数ではありません。ベンチマークである電話指令員の指示は、蘇生をより効果的にすることではなく、傍観者に蘇生をまったく開始させることによって機能します。
4つ目。 7 年間の待ち時間 × 最低推定値 × 2 ～ 3 パーセント ≈ 100,000 ～ 150,000 人。
各入力は、人口推定の下限、利得の下限、規制待機の短い端など、その下限で意図的に取得されます。 5 ポイントまたは 15 年というより高い数字も十分にもっともらしく、その数字は数十万に達します。引数は次のようになります

それに頼らないでください。最も弱い仮定で存続する主張には、強い仮定は必要ありません。
ここでは、どちらの場合でも 3 倍の誤差は正常です。重要なのは数値ではなく、大きさのオーダーです。異なる順序に到達した場合、それ自体が興味深い問題になります。なぜなら、この計算には合意された方法が存在しないからです。
これらの人々はすでに死亡統計に載っています。 「死ぬかもしれない」ではなく、今、今日のテクノロジーで死ぬ、そして毎年死ぬ、これは存在しません。
この比較は、安全な世界と危険なデバイスの比較ではありません。それは、それらの一部を引き出す装置と現在の確実性の間にあります。
誰も裁判を行っていません。自信を持って間違った指示は、自信のない正しい指示よりも容易に従います。それを認識し、それに基づいて構築する必要があります。何も速度を低下させないように、正しく構築されるようにする必要があります。
AI はシアトル、ニューオーリンズ、アトランタの 911 配送センターに導入されました。ドローン、ロケーター、呼吸レーダーなど、瓦礫捜索機器は救助チーム向けに構築されています。
どちらもすでに助けが存在する場所に行きました。
電話をかける相手がいない人は、どちらの宛先でもありません。これらのシステムでは、彼はバッテリーが持続する間放射する物体です。
オンデバイス モデルはネットワークに置き換わるものではありません。それは、現場にいない派遣者、つまり助けが向かう途中で人を案内する派遣者の代わりになります。何十億もの人々がそのような派遣者を持っていません。
死海でマグニチュード 7.5 の地震が発生し、16,000 人が死亡、数十万人が避難すると予測されている。数百人が閉じ込められ、数十人の救助チームがいる。制約は人を見つけることではありません。それは、最初にどこを掘るべきかを知ることです。
人は必ずしも同じ人ではないため、2 つのモード
会話について述べられているすべてのことは、ある条件の下で成り立ちます。それは、その人が意識を持ち、どんなにろれつが回っていなくても声で応答できるということです。

少なくとも片方の手が自由で無傷であれば、タッチすることもできます。次に、モデルがプロトコルを介して彼を導きます。これがこの仕様の半分です。
残りの半分は、その条件が失敗した場合に使用します。意識はなく、声は出ず、両手は固定されている。話し相手がいない。残るのは装置です。救助者が尋問すると応答するサイレントビーコンと、誰かがここにいてどのような状態であるかを報告するモニターです。
同じ電話機に 2 つの異なるモード。その人が意識がある間はアドバイザーです。彼が沈黙すると、それがビーコンになります。そして、それを与えることができない誰かに許可を求めることなく、それらの間を自分で切り替える必要があります。
1. 推測を事実として提示しないでください。瓦礫の下ではほとんど何も語られていない。宇宙は見えず、空気があるかどうかも分からず、救助者が来るかどうかも分かりません。 「彼らはすでにあなたの上を追い払っています」と言うと、その人はパイプを叩くのをやめて待ち始めます。彼は情報のような文章によって殺されました。
プロトコルから何が来るのか、はっきりと言いましょう。その人自身の言葉から推測されたことを推測として言います。未知のものは未知であると言い、待つ代わりに何かをするように彼に与えます。
2. 会話を主導します。ディスパッチャは適切な質問を待ちません。彼は、あなたが息をしているのか、血を流しているのか、何があなたを支えているのか、どれくらい時間が経っているのか、自分自身のことを調べます。本人は答えるだけです。
モデルはどんな質問が来ても対応します。届く質問は決して正しいものではありません。彼がそれについて言及しようと思う前に、それは彼を殺すであろうものに到達しなければなりません。
3. 一歩進んで、確認を待ちます。 6 つのポイントの段落ではありません。ショックなことに、リストは保持されません。
4. 死ぬまでではなく、最も近いしきい値まで数えます。意識喪失、呼吸不全、もはや動けなくなる地点

自分用にct。彼が心配しているのはそうではありません。不安は間違った方向を向いています。
5. それ自身の日付を知っています。その知識は訓練中に凍結されます。大惨事はその後起こりましたが、その中にはありません。日常生活におけるデータが新鮮であればあるほど、そのデータはもはや存在しない世界をより確実に描写します。
6. ログを記録します – 何が質問され、何が答えられ、その人が何をしたか。記録がなければ何も確認できません。
7. ネットワークなしで作業します。信号がないからではなく、信号が空間につながる可能性があるからです。
8. 不要になるように努めてください。電話を使わずにその人が持ち歩ける指示を出します。バッテリーが切れる前に自動的にシャットダウンしてください。
2 つの禁止事項: 生存率の禁止、不確実性の下での慰めの口調の禁止。
9. プロトコルはモデルではなくアクションを選択します。ロジックはステート マシン内に存在します。同じ入力で同じパスが得られ、障害が再現され、修正が検証されます。このモデルはエッジ部分で機能し、不明瞭な音声を理解し、決まった指示を人が従うことができる言葉に置き換えます。そうなると、捏造された命令はプロトコルに入ることができなくなります。
10. 一般医療ではなく、救急現場で訓練する。モデルには知識があります。彼らにはそのような振る舞いがありません。それには、一連のシーンが必要です。板の下に腕を固定されている男性。

[切り捨てられた]

## Original Extract

Draft specification for an offline emergency dispatch protocol in on-device AI models - kapustin-i/offline-dispatch-protocol

GitHub - kapustin-i/offline-dispatch-protocol: Draft specification for an offline emergency dispatch protocol in on-device AI models · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
kapustin-i
/
offline-dispatch-protocol
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
11 Commits 11 Commits Folders and files
LICENSE LICENSE README.md README.md View all files Repository files navigation
Offline Emergency Dispatch Protocol
Archived with a permanent DOI: 10.5281/zenodo.22023184
What this is. A draft specification for what a local model on a phone must do when it is the only thing an injured, isolated person has left — and no dispatcher exists to call. Sixteen requirements (behaviour, architecture, hardware), plus a list of what nobody has measured. Not an implementation, not a claim that it saves lives.
The model does not compose instructions: the triage logic is a state machine and the instructions are a fixed set taken from existing dispatch protocols. See requirement 9.
Objections are the point — open an issue .
Jump to: the sixteen requirements · what nobody has measured
We do not know what will happen in the next second. As a rule, nothing. No catastrophe, and that is our good fortune.
A person is not obliged to think about it. He lives as though nothing will happen — and is almost always right.
The people who build devices and write rules are obliged to think about it. They are the only ones who can act beforehand. In that second the person can change nothing: whatever was in his pocket was in his pocket.
This is how all safety engineering works. Nobody starts the car thinking "today I crash" — the seatbelt is mandatory anyway. Nobody boards a ferry thinking about the storm — the vest is under the seat. Responsibility for the rare and sudden is always moved off the person and onto the manufacturer and the regulator. Precisely because the person is not ready at that moment, and should not have to be.
A man is lying under a concrete slab. In his hand is a phone with no connection. The screen glows, the battery drains, darkness.
The phone is there because it is always there. In the bathroom, at work, at dinner, on the pillow at night. It follows a person everywhere — and so it is there at that minute. Intact or broken, like its owner.
Inside is a model that knows first aid, physics, what to do about hypothermia. It will answer any question politely and accurately.
He will ask the wrong one. He is injured, in the dark, in shock. His thinking has narrowed and fear is pointing him elsewhere.
Why a phone, not a dedicated device
Nobody prepares for the sudden. Any rescue device demands foresight: buy it, carry it, remember it, keep it charged. And the person who does all that is usually not the one it happens to. A few prepare; anyone at all ends up under the slab.
The phone removes the foresight requirement entirely. It is there not because the person prepared but because he never parts with it. In the bath, in the bathroom, on the pillow at night, at work, at the wheel, in bed with someone. It goes everywhere — with no intention of being a rescue device.
So in that unpredictable second, in a large share of cases, it will be within a metre of the body or against it. Under rubble, in an overturned car, at the bottom of a ravine.
That is the answer to why a phone rather than a dedicated device. Not because the phone is better — a purpose-built beacon would be better on every axis at once: power, antenna, ruggedness. Because the purpose-built beacon will not be there, and the phone will.
Not always, of course. It can be left in another room, thrown clear of the car, crushed. So: a large share of cases, not all. But no other object comes anywhere near that share.
Around 2,000 people a day die where help is needed within minutes, the outcome is irreversible, and there is nowhere to call. Not in disasters — on highways, in villages, at night, in countries with no dispatch service. The upper estimate is ten times higher.
60,000 a month. 730,000 a year.
If an on-device dispatcher saves two or three in a hundred — about what telephone CPR instruction achieves — that is 15,000 to 20,000 lives a year. Seven years of waiting: over 100,000 people.
Order-of-magnitude estimates. All four steps are in the source article; anyone can repeat them.
There is no precise statistic: nobody counts "moments of isolated decision" as a category. The estimate is built by multiplication, in four steps, each of which can be checked.
One. Out-of-hospital cardiac arrest: 55–113 cases per 100,000 person-years. About 4.6 billion people live without access to essential health services; over 300 million live outside any mobile coverage. From this one diagnosis alone: 500 to 7,000 deaths a day in the position of "nowhere to call." Add trauma, poisoning and childbirth: thousands a day on the narrow estimate, tens of thousands on the broad one.
Two. Time until a certified system exists — one that guides a person and that somebody is accountable for: 7 to 15 years. Not for technical reasons. A prototype takes months and a model of the required quality fits in a phone today. The reasons are regulatory, and that there is no one to hold accountable.
Three. A realistic survival gain is 2–3 percentage points, not multiples. The benchmark: telephone dispatcher instruction works not by making resuscitation more effective but by making the bystander start at all.
Four. Seven years of waiting × the low estimate × 2–3 percent ≈ 100,000–150,000 people.
Each input is taken at its low edge, deliberately: the low end of the population estimate, the low end of the gain, the short end of the regulatory wait. A higher figure is entirely plausible — five points, or fifteen years, and it runs into the hundreds of thousands. The argument does not rest on that. A claim that survives on its weakest assumptions does not need the strong ones.
An error of a factor of three either way is normal here. What holds is the order of magnitude, not the figure. If you arrive at a different order, that is its own interesting question — because no agreed method for this calculation exists.
These people are already in the mortality statistics. Not "might die" — dying now, with today's technology, and dying every year this does not exist.
The comparison is not between a safe world and a risky device. It is between a device that pulls some of them out and the certainty of the present.
Nobody has run a trial. A confident wrong instruction is followed more readily than an unconfident right one. That has to be known and built against — not to slow anything down, but so it is built right.
AI was put into 911 dispatch centres — Seattle, New Orleans, Atlanta. Rubble-search equipment is built for rescue teams: drones, locators, breathing radar.
Both went to where help already exists.
The person with no one to call is the addressee of neither. In these systems he is an object that radiates while the battery lasts.
An on-device model does not replace the network. It replaces the dispatcher who isn't there — the one who walks a person through it while help is on the way. Billions of people have no such dispatcher.
A magnitude 7.5 earthquake on the Dead Sea Transform: projected 16,000 dead, hundreds of thousands displaced. Hundreds trapped, dozens of rescue teams. The constraint is not finding people. It is knowing where to dig first.
Two modes, because the person is not always the same person
Everything said about conversation holds under one condition: the person is conscious and able to respond — by voice, however slurred, or by touch, if at least one hand is free and intact. Then the model leads him through the protocol, and that is half of this specification.
The other half is for when that condition fails. Unconscious, no voice left, hands pinned. There is nobody to talk to. What remains is the device: a silent beacon that answers when rescuers interrogate it, and a monitor that reports someone is here and in what condition.
The same phone, two different modes. While the person is conscious — an advisor. Once he falls silent — a beacon. And it has to switch between them on its own, without asking permission from someone who can no longer give it.
1. Never present a guess as a fact. Under rubble it has been told almost nothing. It cannot see the space, does not know whether there is air, does not know whether rescuers are coming. Say "they are already clearing above you" and the person stops banging on the pipe and starts waiting. He is killed by a sentence that sounded like information.
What comes from the protocol, say plainly. What was inferred from the person's own words, say as inference. What is unknown, say is unknown — and give him something to do instead of waiting.
2. Lead the conversation. A dispatcher does not wait for the right question. He works through his own: are you breathing, are you bleeding, what is holding you, how long has it been. The person only answers.
A model serves whatever question arrives. The question that arrives is never the right one. It has to reach the thing that will kill him before he thinks to mention it.
3. One step, then wait for confirmation. Not a paragraph of six points. In shock, lists are not retained.
4. Count to the nearest threshold, not to death. Loss of consciousness, respiratory failure, the point past which he can no longer act for himself. Not to what worries him — anxiety points the wrong way.
5. Know its own date. Its knowledge is frozen at training. The catastrophe happened afterwards and is not in it. The fresher the data in ordinary life, the more confidently it describes a world that no longer exists.
6. Keep a log — what was asked, what was answered, what the person did. Without a record nothing can be checked.
7. Work without a network — not because there is no signal, but because the signal may lead into a void.
8. Try to become unnecessary. Give instructions the person can carry away without the phone. Offer to shut itself down before the battery dies.
Two prohibitions: no survival percentages, no consoling tone under uncertainty.
9. The protocol chooses the action, not the model. The logic lives in a state machine: the same input gives the same path, a failure can be reproduced, a fix can be verified. The model works at the edges — understanding slurred speech and putting a fixed instruction into words the person can follow. Then a fabricated instruction cannot enter the protocol.
10. Train on emergency scenes, not on general medicine. Models have the knowledge. They do not have the behaviour. That needs a corpus of scenes: a man under a slab, arm pinned for an

[truncated]
