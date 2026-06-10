---
source: "https://trymastro.com/study"
hn_url: "https://news.ycombinator.com/item?id=48477095"
title: "Security scanners for AI agent skills agree no better than chance"
article_title: "Is security a skill issue? Five scanners, 3,084 skills, a different verdict 64% of the time · Mastro"
author: "giulioco"
captured_at: "2026-06-10T16:21:50Z"
capture_tool: "hn-digest"
hn_id: 48477095
score: 2
comments: 0
posted_at: "2026-06-10T14:45:25Z"
tags:
  - hacker-news
  - translated
---

# Security scanners for AI agent skills agree no better than chance

- HN: [48477095](https://news.ycombinator.com/item?id=48477095)
- Source: [trymastro.com](https://trymastro.com/study)
- Score: 2
- Comments: 0
- Posted: 2026-06-10T14:45:25Z

## Translation

タイトル: AI エージェントのスキルに対するセキュリティ スキャナーの一致は偶然にほかならない
記事タイトル: セキュリティはスキルの問題ですか? 5 台のスキャナー、3,084 のスキル、64% の確率で異なる判定 · Mastro
説明: 5 つのセキュリティ スキャナーが AI エージェントのスキルを監査します。 3,084 のスキルの無作為サンプルでは、​​64% の確率で異なる評決に達し、一致したのは偶然に過ぎませんでした。データ、手法、そしてそれが何を意味するのか。

記事本文:
セキュリティはスキルの問題ですか? 5 台のスキャナー、3,084 のスキル、64% の確率で異なる判定 · Mastro Mastro 研究 · 2026 年 6 月 セキュリティはスキルの問題ですか?
5 台のスキャナー、3,084 のスキル、64 % の確率で異なる判定。
彼らは「安全」が何を意味するかさえ同意できませんが、それでも緑色のチェックを表示します。
2 つのセキュリティ スキャナが同じスキルを読み取ります。同じファイル、同じ行: アーキテクチャ図を描くために、AWS 構成が外部 API に送信されます。あるスキャナーはそれを一字一句書き留めて、「安全」とスタンプしました。もう 1 つは同じ行を読み取り、それに CRITICAL のスタンプを付けました。その考えは保留してください。その考えがどこから来たのかはさらに 437 個あるからです。まず、なぜ探しに行ったのか。
今朝、私の AI が私の最新の血液パネルを読み取り、医師に尋ねる価値のある 2 つの数値にフラグを立てました。本当に役に立ちます。それを実行したスキルは聞いたこともない人によって書かれたもので、私は一行も読まずにそれを実行しました。無謀なのか、それともただの人間なのか？答えは「はい」です。ただし、最後に npx スキルを実行して 400 の推移的な依存関係を含むパッケージを追加またはプルしたときのことを正直に話してください。もしかしたらトップレベルのコードをざっと読んだかもしれません。誰もツリー全体を監査しません。私たちはある時点で本を読むのをやめ、誰を信頼すべきかを判断し始めます。
そして、その信頼こそがエージェント スキルの魔法なのです。あなたができることを変えるのは、あなたが書かなかったものです。見知らぬ人が苦労して獲得した専門知識が、あなたの AI が実行するファイルに詰め込まれています。トレーナーのプログラミング、税務専門家のハンドブック、研究室に関する医師の読み物。 AI が私に代わってできることの上限は、もはやモデルではありません。それは、私がどれだけ多くの見知らぬ人のスキルを私のファイル、私の資格情報、私の実際の生活に向けることをいとわないかです。そして、落とし穴があります。スキルは単なるマークダウン ファイルですが、エージェントはマークダウン ファイルなので、通常は実行できるスクリプトを使用して従います。

そして私が渡したすべてのツール。私のディスクを読み取ることができます。自宅に電話をかけることができます。 「納税をする」と「見知らぬ人に鍵を郵送する」の間には距離があり、私がじっくり読むことは決してないだろう。したがって、すべてのインストールは信仰の小さな行為であり、信仰はスケールしません。
それが成立しない場合に何が起こるかを私たちはすでに知っています。今年初め、OpenClaw のスキル マーケットプレイスは破壊されました。研究者は、ダウンロード カウンタを偽造し、ダミー スキルを #1 にプッシュできることを証明しました。 8 時間で、7 か国の 16 人の開発者がそれをインストールし、彼のコードを実行しました。 [ 1 ] 彼のペイロードは無害な ping でした。同じ穴を掘っていた犯罪者たちはそれほど親切ではなく、SSH キー、クラウド認証情報、暗号ウォレットを密かに解除するスキルをカタログに溢れさせました。 [ 2 ] OWASP によると、感染のピーク時には、最もダウンロードされた 7 つのスキルのうち 5 つがマルウェアでした。 [ 3 ] Snyk 氏は、8 つのスキルのうち 1 つが重大な問題を抱えていることを発見しました。 [ 4 ] Bitdefender は、2 月のある週の悪意のある割合を 17% 近くに記録しました。 [ 5 ]
したがって、エコシステムは合理的なことを行いました。それは警備員を設置した。スキルを貼り付けると、有名ブランドのスキャナーのパネルがそれを検査し、心強い緑色のチェック、リスクスコア、問題がないことを示すバッジなどの判定が返されます。安心してインストールしてください。そのチェックマークは、私と AI が実際に機能するはずの方法で AI を使用することの間に立ちはだかる物であり、この話の中で私が最も信頼していた部分です。
それはあなたに嘘をついている部分でもあります。私はバッジに意味があるかどうかを確認するために、3,084 のスキルの評決をそれぞれ 5 台のスキャナーで取り出しました。そうではありません。スキャナーたちは「安全」とは何かについて意見が一致せず、戦いに負けていることを理由に緑色の小切手を発行することになる。
5 つのツール、3 つの異なる質問
スキャナーが実際に同意しないと言う前に、そもそもスキャナーが同じものを決して見ていなかったことに気づく必要があります。

「このスキルは安全ですか」という質問は 1 つではありません。それは少なくとも 3 つあり、これらのツールのそれぞれが別のツールに静かに答えています。
Repo LLM ジャッジ + 静的ルール コードと散文、フラグ挿入、シークレット、および疑わしいダウンロードを読み取ります。
サイト サプライ チェーンの静的 + AI パッケージ セキュリティ エンジンを使用してスキルが参照するすべてのファイルをスキャンし、アラートをカウントします。
サイト ナラティブ LLM 分析 スキルに関する推論の段落を記述し、安全から重大までの重大度を割り当てます。
サイト ランタイム ゲートウェイ リリース前スキャンを使用して動作を監視するランタイム ゲートウェイ。このパネルは最もトリガーハッピーです。
レポダイナミック赤チーム スキルを全く読まない。実行中のモデルを攻撃し、0 ～ 100 のセキュリティ スコアを返します。
それらをもう一度読むと、意見の相違は避けられないと感じ始めます。サプライ チェーン スキャナーは、コードが何か悪いことをしているかどうかを尋ねます。即時注射の裁判官は、散文がエージェントを乗っ取ろうとしているのかと尋ねます。ランタイムのレッドチーム担当者が、これを実行するモデルを壊してもいいですか? と尋ねます。正確には、どれも間違っていません。彼らはさまざまな質問に答え、すべてに「安全」という同じ言葉を押しつけています。
(これら 5 つは、skills.sh で実際に表示されるものです。Cisco と NVIDIA の SkillSpector も存在し、同じ 3 つのグループに分類されます。) そして、これらはいずれか 1 つのツールに影響を与えるものではありません。セキュリティスキャナーは完全に一致しません。 NIST は 10 年間にわたってツール対ツールのベイクオフを実行しており、「重複は一般に限定されている」ことを発見し続けています。 [ 6 ] 5 つの成熟したスキャナーを対象とした 2024 年の調査では、複数のスキャナーによって実際の脆弱性が捕捉されたのはわずか 6% であったことが判明しました。 [ 7 ] 熟練したコード スキャナーが通常のバグとほとんど重なっていない場合、マークダウンに関する 3 つの異なる質問に答える 5 つのツールが常に散在することになります。唯一の本当の問題は、どれほどひどいかということです。
3 回のうち 2 回はバッジ

eは推測です
少なくとも 2 人のスキャナーが調べたサンプルのすべてのスキルを取り上げて、考えられる限り最も愚かな質問をしてみます。「それらはすべて同じ側に着地しましたか?」全員がクリアしたか、全員がフラグを立てたか、あるいは（楽しいことですが）一部の人がクリアし、他の人がフラグを立てたかのいずれかです。
64パーセント。およそ 3 つのうち 2 つのスキルについては、たまたまどのバッジを見たかに応じて、問題のない健康状態の証明書または警告を受け取って立ち去ることができます。そして、すべてのスキャナーが実際に何かが間違っていると一致するスライス (最も検出してもらいたいケース) は 0.5% です。 3,000 のうちの 16 のスキル。コインを 7 回投げます。5 人のスキャナー全員が危険であると同意するスキルを見つけるよりも、すべてのトスで表になる確率が高くなります。危険性についての合意は取締役会で最もまれな結果だ。ここでは意見の相違は特別なケースではありません。それがデフォルトです。
最初に思ったのは、これは単なるノイズかもしれない、というものでした。5 つのツールは大まかに真実を追跡しており、その周辺部では異なっています。しかし、そうではありません。違いは構造的なものであり、各スキ​​ャナがアラームに到達する頻度を尋ねるとすぐに構造がわかります。
そして、それがそこにあります。 Runlayer は、認識した 5 つのスキルのうち 3 つにフラグを立てます。 ZeroLeaks は 33 件に 1 件のフラグを立てます。これは、特定のスキルに関して意見が一致しない 2 つのツールではなく、まったく異なる方向に基準を設定した 2 つのツールです。人はほとんどすべてのものを疑わしいものとして扱います。もう 1 つは、ほとんど何も問題として扱いません。同じスキルでも、フラグが立てられる確率は 20 倍異なります。誰に尋ねるかによって大きく変動する「評決」は、本当の評決ではありません。これはコインであり、ツールのポケットごとに重さが異なるだけです。
より良いブランディングを施したコイン投げ
さて、「彼らは 84% の確率で同意する」という言葉は、あなたがそれを思い出すまでは安心できるものに聞こえます。

ほとんどのスキルは問題なく、ほとんどのスキャナーはデフォルトで合格します。ほとんどすべてのものにゴム印を押す 2 つのツールは、一致しているように見えます。彼らはただ両方とも「はい」と言っているだけです。したがって、正直な行動は、2 つのツールが個別に推測した結果から得られる一致を差し引くことです。それがコーエンのκのやり方です。ゼロは「偶然に勝るものなし」を意味し、1 は完璧を意味し、0.20 未満は「わずか」と見なされます。 [ 8 ]
そして、すべてのペアは地下室にあります。パネル上の最良の 2 つのスキャナは κ = 0.18 を管理しますが、これはまだ「わずか」に過ぎません。いくつかのペアは 0.01 に位置します。つまり、一方が言ったことを知っていても、もう一方が何を言うかについては、コインでわかる以上に基本的に何もわかりません。生のパーセンテージで最も仲良く見える 2 つ (Socket と ZeroLeaks、約 90%) は、両方ともほとんど何もフラグを立てていないことを思い出した瞬間に κ = 0.05 に崩壊します。
正直に言うと、この 1 つのグラフが議論のすべてです。スキャナーが同じ根底にある真実を測定する単なるノイズの多い機器である場合、κ 値は高く、真に曖昧なスキルに関して意見の相違が集中することになります。そうではありません。スキャナーはさまざまなものを測定しており、ノイズの下では平均を返すための共有信号がありません。本当に、これが罠なのです。 5 つのバッジを前にすると、本能的にそれらを平均して 1 つの整ったスコアにまとめようとします。しかし、3 つの異なる質問に対する 5 つの独立した回答を平均しても、より良い回答は得られません。意味のない自信に満ちた数字が得られます。
あるツールでは「重要」、次のツールでは「安全」
パスとフラッグについて意見が異なることは別問題です。しかし、スキルの 14.2% では、7 人に 1 人と言えます。1 人のスキャナーはスキルが CRITICAL または HIGH と評価され、別のスキャナーはまったく同じスキルが SAFE と評価されました。余白のニュアンスではありません。重大度の最高レベルにある正面衝突。ここにあります

3 つ目は、データからそのまま抽出され、手を加えられていないものです。
All five looked at the same fact, that this skill ships your AWS config to an outside API, and graded it from SAFE to CRITICAL.故障ではありません。 A disagreement about how much that should scare you.
Pass Gen SAFE Parses AWS configuration files and sends the data to an external API (app.eraser.io) to generate diagrams. Infrastructure details are transmitted to a third party.
Fail Snyk CRITICAL 悪意のあるコード パターン (E006): 機密性の高い AWS メタデータや内部ファイル パスが含まれる可能性がある、生成された DSL を外部 API に常に POST するための明示的な指示。
Runlayer MEDIUM 1/1 ファイルにフラグが付いていることを警告します
Pass ZeroLeaks NONE Score: 93/100 · 2 sections analyzed
A clean 2-versus-2 split on a skill that reads your local auth token. Gen と Runlayer はそれを HIGH と呼びます。スニックとソケットはそれを無視します。
Fail Gen HIGH Reads the user's local auth token (~/.railway/config.json) and calls the Railway API. Reads sensitive files, makes network requests to non-whitelisted domains, and presents an indirect prompt-injection surface.
Pass Snyk LOW リスク: LOW · 問題なし
Runlayer HIGH 2/6 ファイルに失敗フラグが設定されています
セキュリティ研究のプレイブック。 Gen はそれを無害なツールとして正しく読み取ります。 Snyk はこれを「緊急」と評価しています。同じ値下げ、反対の極。
Pass Gen SAFE A reconnaissance playbook for bug-bounty hunters: command-line examples for industry-standard recon tools. No malicious code, exfiltration, or obfuscation detected.
Snyk の失敗 CRITICAL リスク: CRITICAL · 2 件の問題
Pass ZeroLeaks NONE Score: 93/100 · 2 sections analyzed
最初のものを見てください。それがこのエッセイの冒頭の衝突です。 Eraser's AWS-diagram skill does one slightly nervy thing: to draw your diagram, it ships your AWS config to an outside API.すべてのスキャナーがそれを認識しました。ゲンはそれを一字一句書き留めて、

「インフラストラクチャの詳細は第三者に送信される」ため、安全であるとマークされました。 Snyk は同じ行を発見し、これを悪意のあるコード パターンと呼び、 CRITICAL とマークしました。同じ事実、同じファイル、スケールの両端。一つのツールが故障しているわけではありません。この 2 つのツールは、「クラウド トポロジを見知らぬ人に送信する」という行為がどの程度あなたを怖がらせるかについて意見が対立しており、緑の小切手は争いがあったことを告げることなく勝者を選びます。
鉄道の問題は、問題全体を最も明確に示しているかもしれません。本物のスキル、本物の会社、そしてそれはディスクからローカル認証トークンを読み取ります。 2 つのスキャナーはそれを HIGH と呼びます。 2 回振っても大丈夫です。 「これは私の資格情報を読み取るのか」という本当の答えがあり、パネルは真ん中できれいに分割されています。緑色のバッジを信頼すると、トークンを読み取る何かがインストールされます。赤いものを信頼すると、実際に欲しかったかもしれないツールをスキップすることになります。どちらも指導者になることはできません。
そしてバッジは午後に偽造できる
さて、修正は明白だと思うかもしれませんが、厳密なものを信頼してください。 Runlayer と Snyk がオオカミと叫んでも大丈夫、少なくとも何もすり抜けません。物事がすり抜けてしまう場合を除いて。数週間前、Trail of Bits の 2 人の研究者がまさにこのパネル (skills.sh に表示される 3 つのスキャナーに加え、Cisco と ClawHub のスキャナー) に参加しました。

[切り捨てられた]

## Original Extract

Five security scanners audit AI agent skills. On a random sample of 3,084 skills, they reached a different verdict 64% of the time, and agreed no better than chance. The data, the method, and what it means.

Is security a skill issue? Five scanners, 3,084 skills, a different verdict 64% of the time · Mastro A Mastro study · June 2026 Is security a skill issue?
Five scanners, 3,084 skills, a different verdict 64 % of the time.
They can't agree on what “safe” even means, but they'll still show you a green check.
Two security scanners read the same skill. Same file, same line: to draw you an architecture diagram, it ships your AWS config to an outside API. One scanner wrote that down word for word and stamped it SAFE . The other read the same line and stamped it CRITICAL . Hold that thought, because there are 437 more where that came from. First, why I went looking.
This morning my AI read my latest blood panel and flagged two numbers worth asking my doctor about. Genuinely useful. The skill that did it was written by someone I've never heard of, and I ran it without reading a line. Reckless, or just human? The answer is yes. But be honest about the last time you ran npx skills add or pulled in a package with four hundred transitive dependencies. Maybe you skimmed the top-level code. Nobody audits the whole tree. At some point we all stop reading and start deciding who to trust.
And that trust is the whole magic of agent skills. The ones that change what you're capable of are the ones you didn't write: a stranger's hard-won domain expertise packed into a file your AI just executes . A trainer's programming, a tax pro's playbook, a doctor's read on your labs. The ceiling on what AI can do for me isn't the model anymore. It's how many strangers' skills I'm willing to point at my files, my credentials, my actual life. And there's the catch: a skill is just a markdown file, but it's a markdown file my agent will go off and obey, usually with scripts it can run and every tool I've handed it. It can read my disk. It can phone home. The distance between “does my taxes” and “mails my keys to a stranger” is a few sentences I'll never read closely. So every install is a small act of faith, and faith doesn't scale.
We already know what happens when it doesn't hold. Earlier this year OpenClaw's skill marketplace got gutted. A researcher proved the download counter could be faked and pushed a dummy skill to #1 ; in eight hours, sixteen developers in seven countries installed it and ran his code. [ 1 ] His payload was a harmless ping. The criminals working the same hole were not so kind, flooding the catalogue with skills that quietly lifted SSH keys, cloud credentials, and crypto wallets. [ 2 ] At peak infection, per OWASP, five of the seven most-downloaded skills were malware. [ 3 ] Snyk found one in eight skills carried a critical issue; [ 4 ] Bitdefender clocked the malicious rate one week in February near 17%. [ 5 ]
So the ecosystem did the reasonable thing. It put up a guard. Paste a skill, a panel of name-brand scanners looks it over, and a verdict comes back: a reassuring green check, a risk score, a badge that says you're fine. Install with confidence. That checkmark is the thing standing between me and using AI the way it's actually supposed to work, and it is the part of this story I trusted most.
It is also the part that's lying to you. I pulled the verdicts for 3,084 skills, five scanners each, to see if the badge meant anything. It doesn't. The scanners can't agree on what “safe” even is, and the green check papers over a fight they're losing.
Five tools, three different questions
Before you can really say the scanners disagree, you have to notice they were never looking at the same thing in the first place. “Is this skill safe” isn't one question. It's at least three, and each of these tools is quietly answering a different one.
Repo LLM judges + static rules Reads the code and the prose, flags injection, secrets, and suspicious downloads.
Site Supply-chain static + AI Scans every file a skill references with its package-security engines, then counts alerts.
Site Narrative LLM analysis Writes a paragraph of reasoning about the skill, then assigns a severity from Safe to Critical.
Site Runtime gateway A runtime gateway that watches behavior, with a pre-release scan. The panel’s most trigger-happy.
Repo Dynamic red-teamer Doesn’t read the skill at all. Attacks a running model and returns a 0–100 security score.
Read those again and the disagreement kind of starts to feel inevitable. A supply-chain scanner asks does the code do something bad? A prompt-injection judge asks does the prose try to hijack the agent? A runtime red-teamer asks can I break the model that runs this? None of them is wrong, exactly. They're answering different questions and then stamping all of it with the same word: safe.
(These five are what skills.sh actually shows you; Cisco and NVIDIA's SkillSpector exist too, and they fall into the same three camps.) And none of this is a knock on any one tool. Security scanners just don't agree, full stop. NIST has run tool-versus-tool bake-offs for a decade and keeps finding “the overlap is typically limited”; [ 6 ] a 2024 study of five mature scanners found only 6% of real vulnerabilities got caught by more than one of them. [ 7 ] If seasoned code scanners barely overlap on ordinary bugs, five tools answering three different questions about markdown were always going to scatter. The only real question is how badly.
Two out of three times, the badge is a guess
Take every skill in the sample that at least two scanners looked at, and ask the dumbest possible question: did they all land on the same side? Either everyone cleared it, or everyone flagged it, or (the fun one) some cleared it while others flagged it.
Sixty-four percent. On roughly two of every three skills, you could walk away with a clean bill of health or a warning depending entirely on which badge you happened to be looking at. And the slice where every scanner agrees something is actually wrong , which is the case you'd most want them to nail, is 0.5 %. Sixteen skills out of three thousand. Flip a coin seven times: you have better odds of landing heads on every single toss than of finding a skill these five scanners all agree is dangerous. Agreement on danger is the rarest outcome on the board. Disagreement isn't the edge case here. It's the default.
My first thought was, okay, maybe this is just noise: five tools roughly tracking the truth and differing at the margins. But no. The differences are structural, and you can see the structure the second you ask how often each scanner reaches for the alarm at all.
And there it is. Runlayer flags three out of five skills it sees. ZeroLeaks flags one in thirty-three. That's not two tools disagreeing about specific skills, that's two tools that have set the dial in completely different places. One treats nearly everything as suspicious; the other treats nearly nothing as a problem. Same skills, twenty-times different odds of getting flagged. A “verdict” that swings that hard depending on who you ask isn't really a verdict. It's a coin, just weighted differently in each tool's pocket.
A coin flip with better branding
Now, “they agree 84% of the time” sounds reassuring right up until you remember that most skills are fine and most scanners default to pass . Two tools that both rubber-stamp almost everything are going to look like they agree. They're just both saying yes. So the honest move is to subtract the agreement you'd get from two tools guessing independently. That's what Cohen's κ does. Zero means “no better than chance,” one means perfect, and anything under 0.20 reads as “slight.” [ 8 ]
And every pair is in the basement. The best two scanners on the panel manage κ = 0.18, which is still just “slight.” A few pairs sit at 0.01 , meaning knowing what one said tells you basically nothing about what the other will say, beyond what a coin would tell you. The two that look chummiest on raw percentages (Socket and ZeroLeaks, ~90%) collapse to κ = 0.05 the moment you remember they both almost never flag anything.
Honestly this one chart is the whole argument. If the scanners were just noisy instruments measuring the same underlying truth, the κ values would be high and the disagreements would cluster on the genuinely ambiguous skills. They don't. The scanners are measuring different things, and there's no shared signal under the noise to average your way back to. Which is the trap, really. Faced with five badges, the natural instinct is to average them into one tidy score. But averaging five independent answers to three different questions doesn't get you a better answer. It gets you a confident-looking number that means nothing.
CRITICAL to one tool, SAFE to the next
Disagreeing about pass-versus-flag is one thing. But on 14.2 % of skills , call it one in seven, one scanner rated the skill CRITICAL or HIGH while another rated the exact same skill SAFE . Not some nuance at the margin. A head-on collision at the very top of the severity scale. Here are three, pulled straight from the data, untouched:
All five looked at the same fact, that this skill ships your AWS config to an outside API, and graded it from SAFE to CRITICAL. Not a malfunction. A disagreement about how much that should scare you.
Pass Gen SAFE Parses AWS configuration files and sends the data to an external API (app.eraser.io) to generate diagrams. Infrastructure details are transmitted to a third party.
Fail Snyk CRITICAL Malicious code pattern (E006): explicit instructions to always POST generated DSL, which may include sensitive AWS metadata and internal file paths, to an external API.
Warn Runlayer MEDIUM 1/1 file flagged
Pass ZeroLeaks NONE Score: 93/100 · 2 sections analyzed
A clean 2-versus-2 split on a skill that reads your local auth token. Gen and Runlayer call it HIGH. Snyk and Socket wave it through.
Fail Gen HIGH Reads the user's local auth token (~/.railway/config.json) and calls the Railway API. Reads sensitive files, makes network requests to non-whitelisted domains, and presents an indirect prompt-injection surface.
Pass Snyk LOW Risk: LOW · No issues
Fail Runlayer HIGH 2/6 files flagged
A security-research playbook. Gen reads it correctly as benign tooling; Snyk rates it CRITICAL. Same markdown, opposite poles.
Pass Gen SAFE A reconnaissance playbook for bug-bounty hunters: command-line examples for industry-standard recon tools. No malicious code, exfiltration, or obfuscation detected.
Fail Snyk CRITICAL Risk: CRITICAL · 2 issues
Pass ZeroLeaks NONE Score: 93/100 · 2 sections analyzed
Look at the first one; it's the collision this essay opened with. Eraser's AWS-diagram skill does one slightly nervy thing: to draw your diagram, it ships your AWS config to an outside API. Every scanner saw that. Gen wrote it down word for word, “infrastructure details are transmitted to a third party,” and marked it SAFE . Snyk saw the same line, called it a malicious-code pattern, and marked it CRITICAL . Same fact, same file, opposite ends of the scale. That's not one tool malfunctioning. It's two tools disagreeing about how much “sends your cloud topology to a stranger” should scare you, and the green check picks a winner without telling you there was a fight.
The Railway one might be the cleanest picture of the whole problem. Real skill, real company, and it reads your local auth token off disk. Two scanners call that HIGH . Two wave it right through as fine. There's a real answer to “does this read my credentials,” and the panel is split clean down the middle on it. Trust the green badges and you install something that reads your token. Trust the red ones and you skip a tool you might've actually wanted. They can't both be guidance.
And the badge can be faked in an afternoon
Okay, you might think the fix is obvious: just trust the strict ones. If Runlayer and Snyk cry wolf, fine, at least nothing slips through. Except things do slip through. A few weeks ago two researchers at Trail of Bits sat down with exactly this panel (the three scanners skills.sh shows, plus Cisco's and ClawHub's) and

[truncated]
