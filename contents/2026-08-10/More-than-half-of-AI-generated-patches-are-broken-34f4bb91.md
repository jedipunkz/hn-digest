---
source: "https://cyberscoop.com/ai-code-patching-security-risks/"
hn_url: "https://news.ycombinator.com/item?id=49245826"
title: "More than half of AI-generated patches are broken"
article_title: "More than half of AI-generated patches are broken | CyberScoop"
author: "speckx"
captured_at: "2026-08-10T16:43:43Z"
capture_tool: "hn-digest"
hn_id: 49245826
score: 3
comments: 0
posted_at: "2026-08-10T16:23:16Z"
tags:
  - hacker-news
  - translated
---

# More than half of AI-generated patches are broken

- HN: [49245826](https://news.ycombinator.com/item?id=49245826)
- Source: [cyberscoop.com](https://cyberscoop.com/ai-code-patching-security-risks/)
- Score: 3
- Comments: 0
- Posted: 2026-08-10T16:23:16Z

## Translation

タイトル: AI が生成したパッチの半分以上が壊れている
記事タイトル: AI 生成パッチの半分以上が壊れている |サイバースクープ
説明: 新しい調査によると、ChatGPT や Claude などの AI モデルはサイバーセキュリティのバグを修正できないことが多く、代わりに新たなコードの脆弱性が発生することがよくあります。

記事本文:
メインコンテンツにスキップ
広告
サイバースクープ
検索
閉じる
検索:
検索
ナビゲーションを開く
トピックス
戻る
AI
購読する
広告
最新のサイバーセキュリティ ニュースをまず Google で入手してください。
AIが生成したパッチの半分以上が壊れている
この記事を聞く
0:00
さらに詳しく。
この機能では自動音声が使用されているため、発音、口調、感情に誤りが生じる場合があります。
(ゲッティイメージズ)
AI が生成したコードがインターネットの隅々に注入され続けるにつれ、悪意のあるハッカーが悪用する攻撃対象領域が拡大するのではないかとの懸念が高まっています。
大規模な言語モデルの強化されたサイバーセキュリティ機能がチェック機能として機能し、脆弱性が作成されるのとほぼ同じ速さで脆弱性を発見して修正できる可能性があると主張する人もいます。
しかし、OpenAI の ChatGPT 5.5 と Anthropic の Claude Opus 4.8 という 2 つの人気のある商用モデルのパッチ適用機能をテストした新しい研究では、生成 AI は脆弱性を封じ込めるよりも、悪用可能なパッチを作成したり、まったく新しいバグを導入したりする可能性が高いことが判明しました。
1Password の研究者は、攻撃者に Linux クラウド環境への root アクセスを与える可能性のあるカーネルの欠陥である「コピー失敗」脆弱性を含む、6 つの「影響が大きく、複雑性の高い」CVE にパッチを適用するモデルの機能をテストしました。全体的な成功率 (または、新たな問題を引き起こすことなく脆弱性に完全にパッチを適用すること) は、47% でコイントス未満でした。
「私たちの調査結果は、さまざまなシナリオを総合すると、Claude と ChatGPT の両方でパッチ生成の成功率が低いことを示しています。これを、アプリケーションの動作に誤った変更を加えずにすべての既知のエクスプロイト パスを完全に修復することと定義しています」と Keith Hoodlet、Axel Mierchuk、Spencer Michaels は書いています。
「モデルは多くの場合、脆弱なコードのサブセットのみに対応していました。

aths、脆弱性の根本原因に対処できずにテストを満たした脆弱なガード コードを追加し、当面の脆弱性にパッチを適用する際にアプリケーションの動作に微妙な変更を導入する場合がありました」と著者は続けました。
この研究は、AI 時代に急増する脆弱なコードの修正には、ほぼ自律的な脆弱性の発見とパッチ適用がまだ効果的ではない可能性があることを示唆しています。
他の民間部門の調査でも同様の問題が指摘されています。 Veracode の今年のレポートでは、LLM が実行可能なコードの作成において「大きな進歩」を遂げている一方で、「セキュリティは別の話」であることがわかりました。さまざまなフロンティア モデルにわたるテストの結果、AI によって生成されたコードの平均セキュリティ「合格率」は約 56% であることがわかりました。 GPT 5.5 などの新しいモデルは 70% に近づいていますが、半数以上は 50 ～ 53% の間にあります。
Veracode は 100 の異なるモデルをテストし、ばらつきはありましたが、一般的に少数のモデルでセキュリティ パッチ適用の進歩が見られましたが、残りのモデルでは「停滞」が見られました。 1Password の調査と同様に、Veracode テストの 44% で、モデルは検出可能な OWASP トップ 10 の脆弱性をコードベースに導入しました。
重要な注意点: どちらのレポートも、Anthropic の Mythos や OpenAI の GPT-5.6-Sol など、フロンティア企業が大幅に高いサイバーセキュリティ機能を備えていると宣伝している新しいモデルをテストしていません。
これらの高度なモデルは、脆弱なコードを特定して修正できます。 Anthropic と OpenAI は、外国製品やオープンソースの代替製品が競合する前に、Project Glasswing と Daybreak を通じて主要産業にそれらを配布しています。
Veracode の製品担当副社長 Tim Jarret 氏は CyberScoop に対し、AI ツールには依然としてさまざまな制限があり、サイバーセキュリティ パッチの信頼性が低くなる可能性があると語った。

g 知識豊富な人間が関与していない。
SQL インジェクションなどの一部の脆弱性は自動化によって簡単に修正できますが、クロスサイト スクリプティングなどの他のバグはさまざまな方法で悪用される可能性があり、完全に遮断するには人間による介入、追加のコンテキスト、またはその両方が必要です。さらに、モデルは時間の経過とともに以前のセッションのコンテキストを徐々に失う可能性があり、タスクを正しく完了する能力に影響を与え、失われたギャップを埋めるために幻覚を見る可能性が高まります。
「現時点では、エージェントに自由にコードをマージさせるのではなく、チームがレビューして承認する必要があるコードベースへの別のコード変更以外のものとしてこれらを扱うのは時期尚早だと言えると思います」とジャレット氏は述べた。
しかし、AI エージェントが人間の防御者がレビューできるコードを飛躍的に生成している世界では、それは不可能かもしれないことを同氏は認めました。何らかの自動化されたコードレビューが必要になりますが、コードを作成したのと同じ自動化ツールを使用しないことが望ましいです。最終的な目標は、セキュリティにおけるこれまでと同じ、「信頼するが検証する」です。
「90パーセントの場合、人間によるチェックは『クロスチェックはうまくいきましたか？ 「まだ何か問題がある場合は、そこにもう少し注意を集中してください。」
透明な AI エージェントが思っている以上に重要な理由
AI は選挙の事実をより正確に把握できるようになっていますが、有権者は AI に頼るべきではありません
国家サイバー局長、ホワイトハウスが新たなルールを策定せずにAIを保護する計画を説明
被害者117人を虐待した英国人、The Com関係者に有罪判決
AISI、OpenAI、さらなる「未承認」モデルハッキングを報告
民主党上院議員ら、AIのセキュリティリスクに関するトランプ政権の意思決定を批判

s
CrowdStrike: AI は今やサイバー攻撃の武器でもあり標的でもあります
Hugging Face 侵害がエージェント AI 時代の防御について明らかにしたこと
Anthropicは、安全性テスト中に自社のAIが誤って3社をハッキングしたと発表
CISA がオープンソース ソフトウェアのセキュリティに関する推奨事項を連邦政府機関に発行
あまり知られていない npm パッケージは、axios ハッキングに対する北朝鮮の準備行為だった
702 条の失効がサイバーセキュリティにとって何を意味するか
SOC はこのために構築されたものではありません
サイバーセキュリティが米中の AI 競争の中心である理由
AI 軍拡競争に対する建設者の視点
政府
国会議事堂は、行政府と外国の同盟国が詐欺と戦うために十分に連携しているかどうか知りたいと考えている
水道部門は警鐘を鳴らしました。また。
スノーフレークハッカーが有罪を認め、最長32年の懲役刑の可能性
有効期限が迫っているため、議員らは OPM 侵害の被害者のために ID 盗難サービスの保存に乗り出している
企業が機密を漏らさずにサイバーリスクを共有する方法
ワイデン上院議員、連邦政府に対し、古くて安全性の低い公共向けVPNを廃棄するよう要請
マイクロソフトとテクノロジー企業がオープンソースAIの普及を後押し
ホワイトハウス、アンスロピックの寓話を蒸留したとして中国企業を非難
沿岸警備隊、ノースカロライナ州の港を混乱させたサイバー攻撃を監視していると発表
連邦政府の警告にもかかわらず、水道システムで使用される数千の米国の産業用コントローラーがオンラインで公開されたまま
ランサムカルテル創設者に懲役16年の判決
オープンソース ソフトウェアの宿敵 TeamPCP は、誰もが考えていたよりもさらに遡ります
上院、プライバシー、AI、子供の安全に関する一連の法案を審議する予定
CIRCIA に関する業界のメッセージ: サイバー攻撃に関する質問を減らしてください
ルビオ氏、性的強要者やサイバー詐欺師のビザを制限
連邦サイバーセキュリティ報告規則のほとんどは重複していることが研究で判明

## Original Extract

New research shows AI models like ChatGPT and Claude frequently fail to fix cybersecurity bugs, often introducing fresh code vulnerabilities instead.

Skip to main content
Advertisement
CyberScoop
Search
Close
Search for:
Search
Open navigation
Topics
Back
AI
Subscribe
Advertisement
Get our latest cybersecurity news first on Google.
More than half of AI-generated patches are broken
Listen to this article
0:00
Learn more.
This feature uses an automated voice, which may result in occasional errors in pronunciation, tone, or sentiment.
(Getty Images)
As AI-generated code continues to be injected into all corners of the internet, concerns have risen about an expanding attack surface for malicious hackers to exploit.
Some have argued that the enhanced cybersecurity capabilities of large language models could serve as a check, finding and fixing vulnerabilities nearly as fast as they’re created.
But new research that tested the patching capabilities of two popular commercial models, OpenAI’s ChatGPT 5.5 and Anthropic’s Claude Opus 4.8, found that generative AI is more likely to create an exploitable patch or introduce entirely new bugs than close off a vulnerability.
Researchers at 1Password tested the models ability to patch six “high-impact, high-complexity” CVEs, including the “Copy Fail” vulnerability, a kernel flaw that can give an attacker root access to Linux cloud environments. The overall success rate (or fully patching the vulnerability without introducing new problems), was less than a coin flip at 47%.
“Our research findings show that, in aggregate across a variety of scenarios, both Claude and ChatGPT had a low rate of successful patch generation, which we define as full remediation of all known exploit paths with no erroneous changes to application behavior,” wrote Keith Hoodlet, Axel Mierczuk and Spencer Michaels.
“The models often addressed only a subset of vulnerable code paths, added fragile guard code that satisfied tests while failing to address the vulnerability’s root cause, and sometimes introduced subtle changes in the application’s behavior while patching the immediate vulnerability,” the authors continued.
The research suggests that largely autonomous vulnerability-discovery and patching may not yet be effective in fixing the explosion of vulnerable code that is being created in the AI era.
Other private sector research has pointed to a similar problem. A report this year from Veracode found that while LLMs have made “enormous strides” in crafting workable code, “security is a different story.” Testing across a range of frontier models found the average security “pass rate” for AI generated code is around 56%. Newer models like GPT 5.5 push closer to 70%, while more than half sit between 50-53%.
Veracode tested 100 different models and while there was variability, in general a small number of models were showing progress on security patching while the rest have experienced “stagnation.” Similar to the 1Password research, in 44% of Veracode tests the models introduced a detectable OWASP Top 10 vulnerability into the codebase.
An important caveat: neither report tested newer models, like Anthropic’s Mythos or OpenAI’s GPT-5.6-Sol, that frontier companies tout as having significantly higher cybersecurity capabilities.
Those advanced models can identify and fix vulnerable code. Anthropic and OpenAI are distributing them to key industries through Project Glasswing and Daybreak before foreign or open-source alternatives can compete.
Tim Jarret, vice president of product at Veracode, told CyberScoop that AI tools are still subject to a range of limitations that can make them unreliable for cybersecurity patching without knowledgeable humans in the loop.
While some vulnerabilities – like SQL injections – can be easily patched through automation, other bugs like cross-site scripting, can be exploitable in several different ways and require either a human touch, additional context or both to fully close off. Additionally, models can slowly lose context from prior sessions over time, affecting their ability to complete tasks correctly and raising the possibility they’ll hallucinate to fill in the missing gaps.
“I think we would say, at this point, that Iits premature to treat those as anything other than another code change to the code base that needs to be reviewed and accepted by the team, as opposed to letting the agent merge the code freely,” said Jarrett.
However, he acknowledged that may not be possible in a world where AI agents are generating exponentially more code for human defenders to review. Some kind of automated code review will be necessary – preferably not by the same automation tool that produced the code. The ultimate goal is the same as it has always been in security: “trust but verify.”
“Ninety percent of the time, the human check might just be ‘did the cross check look good?’ Do we have a thumbs up?’” Jarrett said. “In those cases where there’s still something wrong, that’s where you focus your attention a little bit more.”
Why transparent AI agents matter more than you think
AI is getting better at election facts, but voters shouldn’t rely on it
National cyber director lays out White House plans to secure AI without writing new rules
UK man tied to The Com sentenced for abusing 117 victims
AISI, OpenAI report more ‘unsanctioned’ model hacks
Dem senators criticize Trump administration decisionmaking on AI security risks
CrowdStrike: AI is now both the weapon and the target in cyberattacks
What the Hugging Face breach reveals about defense in the age of agentic AI
Anthropic says its AI accidentally hacked three companies during safety tests
CISA issues recommendations to federal agencies on open-source software security
A little-known npm package was North Korea’s warm-up act for the axios hack
What the Section 702 lapse means for cybersecurity
The SOC wasn’t built for this
Why Cybersecurity is at the heart of the US-China AI race
A builder’s view of the AI arms race
Government
Capitol Hill wants to know if executive branch, foreign allies coordinated enough to combat scams
The water sector just got it's wake-up call. Again.
Snowflake hacker pleads guilty, faces up to 32 years in prison
Lawmakers spring to save ID theft services for OPM breach victims, with expiration looming
How companies could share cyber risks without exposing their secrets
Sen. Wyden urges feds to discard older, insecure, public-facing VPNs
Microsoft, tech companies throw weight behind spread of open-source AI
White House accuses Chinese company of distilling Anthropic’s Fable
Coast Guard says it is monitoring cyberattack that disrupted North Carolina's ports
Despite federal warnings, thousands of U.S. industrial controllers used in water systems remain exposed online
Ransom Cartel creator sentenced to 16 years in prison
Open-source software’s archenemy TeamPCP goes back further than anyone thought
Senate set to debate package of bills on privacy, AI and kids safety
Industry's message on CIRCIA: Please ask us fewer questions about cyberattacks
Rubio restricts visas for sextortionists, cyber scammers
Most federal cybersecurity reporting rules are duplicative, study finds
