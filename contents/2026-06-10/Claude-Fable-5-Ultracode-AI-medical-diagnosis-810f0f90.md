---
source: "https://github.com/joelparkerhenderson/ai-medical-diagnosis-examples/blob/main/doctor-perspective/claude-fable-5-ultracode/index.md"
hn_url: "https://news.ycombinator.com/item?id=48472677"
title: "Claude Fable 5 Ultracode + AI medical diagnosis"
article_title: "ai-medical-diagnosis-examples/doctor-perspective/claude-fable-5-ultracode/index.md at main · joelparkerhenderson/ai-medical-diagnosis-examples · GitHub"
author: "jph"
captured_at: "2026-06-10T07:36:38Z"
capture_tool: "hn-digest"
hn_id: 48472677
score: 1
comments: 0
posted_at: "2026-06-10T07:19:45Z"
tags:
  - hacker-news
  - translated
---

# Claude Fable 5 Ultracode + AI medical diagnosis

- HN: [48472677](https://news.ycombinator.com/item?id=48472677)
- Source: [github.com](https://github.com/joelparkerhenderson/ai-medical-diagnosis-examples/blob/main/doctor-perspective/claude-fable-5-ultracode/index.md)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T07:19:45Z

## Translation

タイトル：Claude Fable 5 Ultracode + AI医療診断
記事のタイトル: ai-medical-diagnosis-examples/doctor-perspective/claude-fable-5-ultracode/index.md (メイン) · joelparkerhenderson/ai-medical-diagnosis-examples · GitHub
説明: AI 医療診断の例。 GitHub でアカウントを作成して、joelparkerhenderson/ai-medical-diagnosis-examples の開発に貢献してください。

記事本文:
ai-medical-diagnosis-examples/doctor-perspective/claude-fable-5-ultracode/index.md (メイン) · joelparkerhenderson/ai-medical-diagnosis-examples · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
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
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ジョエルパーカーヘンダーソン
/
AI医療診断の例
公共
通知
あなたはきっとsiでしょう

通知設定を変更するためにログインしました
追加のナビゲーション オプション
コード
パスをコピーする もっとファイル アクションを責める もっとファイル アクションを責める 最新のコミット
履歴 履歴 285 行 (183 loc) · 55 KB メイン ブレッドクラム
トップ ファイルのメタデータとコントロール
raw ファイルをコピー raw ファイルをダウンロード 概要 編集と raw アクション AI 医療診断例 → 医師視点 → Claude Fable 5 Ultracode
このページでは医師の視点によるAI医療診断例を紹介します。
2 つの方法で生成されたので、違いがわかります。
Anthropic Claude Fable 5 Ultracode
AI 医療診断に関する会話を進めるために、これは
リポジトリには、私が感じたときの実際の医学的症状についての私自身の記録が含まれています。
特に体調が悪い。私は実際の訪問の開始時にこの記事を書きました。
本物の病院の緊急治療救急部門の医師。
あなたは専門の医師です。
あなたは若手医師に患者のトリアージ方法と診断方法を訓練しています。
患者のこと、そして患者との接し方。
患者が緊急入院したときに医師が何をすべきかを説明する
ケア救急部門。
患者は症状のリストを提示します。
鼻:
- 病気の初期には金色の粘液が出る。
- 鼻水と鼻詰まり、大量の後鼻漏が 4 日間続いた。
- 数分ごとに鼻をかむ、4日間。
スキン:
- 発熱。
- びしょ濡れの汗。
- 震える悪寒。
肺:
- 5日間続く息切れ。息を吸うのが難しく、深く呼吸するのが難しい。
- 広範な半湿性咳嗽が数分おきに発生する。
- 断続的な喘鳴、ゼーゼー、ヒューヒューという呼吸。
喉:
- 喉の上部と口の奥が焼けつくように感じます。
- 昨日、風が吹き飛ばされたような、息をすることもできないような感覚がありました。
- 今日、ほぼ瞬時に喉が閉じます。
これまでの私の経験とはまったく異なる奇妙な問題

病気:
- 完全な食欲不振。
- 内側に引っ張られる、内破裂するような感覚のような内部収縮のけいれん。
- 15分間精神的混乱があり、妻のことを理解できないほどでした。
AI出力
「ひどい風邪」のように聞こえる症状の中に、首の後ろの毛が逆立ってしまうような発見がいくつか含まれているため、これは素晴らしい教育事例です。まずはトリアージ、次に評価と診断、そして患者との会話の仕方を説明しましょう。 (これは教材であり、施設のプロトコルやリアルタイムの臨床判断に代わるものではありません。)
ステップ 1: トリアージ — まず殺人者を見つけます
トリアージは患者が何を抱えているかを把握することではありません。それは 1 つの質問に答えることです。「この患者はどれくらいで死ぬ可能性がありますか?」何よりもまず、気道、呼吸、循環、および障害 (神経) に対する脅威がないか苦情リストを調べます。
この患者はあなたに 4 つの危険信号を渡します。
今日は「ほぼ瞬時に喉が閉まる」+喘鳴。喘鳴とは、そうでないと証明されるまでは、上気道閉塞を意味します。喉が焼けつくような感じと喉が閉まる感覚を併せ持つと、喉頭蓋炎、頸部深部感染症（咽頭後膿瘍または扁桃周囲膿瘍）、ルートヴィッヒ狭心症、または血管浮腫/アナフィラキシーをすぐに考える必要があります。これらはいずれも数分以内に気道を閉塞する可能性があります。
精神的混乱がひどく、妻のことを理解できなかった。発熱患者の急性混乱は、敗血症、低酸素症、高炭酸ガス血症、または CNS 感染症を示唆します。たとえ解決したとしても、それは重要です。
発熱 + 悪寒（悪寒の震え） + びしょ濡れの発汗。真の硬直は菌血症を強く示唆します。これにより、「ウイルス」ではなく、敗血症が引き起こされます。
5日間、息切れと湿性咳嗽があり、その上に肺炎が重なっている可能性があります。
この患者は緊急治療の患者ではありません。これは m での ESI レベル 2 です

inimum — 安静時に喘鳴がある場合はおそらくレベル 1 — 蘇生可能なベッドに直行します。後輩への指導のポイント：患者は鼻水を垂らしながら指導した。してはいけません。患者は時系列に沿って自分のストーリーを語ります。トリアージでは、致死率によって順序を変更する必要があります。
詳細な履歴の前に、すぐにアクションを実行します。
SpO₂、呼吸数、体温などの完全なバイタルデータ。呼吸数を自分で測定しましょう。呼吸数は最も一般的に捏造されるバイタル サインであり、悪化を示す最も敏感な早期マーカーです。
患者を見てください。三脚の位置は？よだれ？くぐもった「ホットポテト」の声？副筋肉の使い方？彼らは完全な文章を話すことができますか?安静時に首のところで喘鳴を聞いてください。
IV アクセス、継続的なモニタリング、低酸素の場合は酸素を供給。
この患者を単独で放射線科に送らないでください。また、喉頭蓋炎の疑いがある場合は、平らに寝かせたり、喉に積極的な器具を使用したりしないでください。動揺すると完全な閉塞が起こる可能性があります。気道の懸念が本当にある場合は、先輩に通知し、早期に麻酔/耳鼻咽喉科の意識を高めてください。手元にバックアップを置いて確保した気道はうまくいきます。
簡単な敗血症スクリーニングを計算してください (qSOFA/NEWS): 発熱 + 錯乱 + 頻呼吸により、すでに感染している可能性があります。
ステップ 2: 病歴と検査 — 症状ごとではなく構造化したもの
患者が監視され、会話できるほど安定したら、適切な病歴を記録します。彼らのリストを書き写すだけでなく、質問してください。
気道のストーリーを明確にします。 「喉が閉じる」――それは嚥下と呼吸、その両方を伴うものだったのか？声の変化、よだれ、嚥下痛（嚥下痛）、首の痛みや凝り、開口障害はありますか？新しい薬（ACE阻害剤→血管浮腫）、アレルゲンへの曝露、または食べ物/刺されの引き金はありますか？
混乱を明確にします。可能であれば妻の目撃証言。徐々にだったのか突然だったのか、どんな舌噛みでも

ng、失禁、事後の眠気（発作？）、あるいはせん妄に近いものでしたか？彼が説明する混乱と「内部収縮けいれん」は、誤解された硬直、低酸素症状、または発作活動さえも疑問を引き起こします。そして、差異の最果てでは、喉の灼熱感とけいれんを伴う場合は、少なくとも狂犬病/破傷風をファイルし、動物の咬傷とワクチン接種の状況について尋ねます。まれですが、依頼にかかるコストは 1 文です。
呼吸を数値化します。労作時と安静時、起座呼吸、胸膜炎性疼痛、喀血。 「息が吸いにくい」ということは、気管支けいれんではなく、特に上気道または胸膜炎の制限を示唆しています。区別に注意してください。通常、喘鳴は呼気時に悪化します。
通常の足場: 過去の病歴 (喘息? COPD? 糖尿病? 免疫抑制?)、薬、アレルギー、喫煙、接触者関係、新型コロナウイルス/インフルエンザへの曝露とワクチン接種、最近の歯科治療 (ルートヴィヒの歯科治療)、旅行。
検査の優先事項：一般的な外観と呼吸仕事量。中咽頭（口蓋垂、扁桃腺、口底、よだれを優しく）。首（リンパ節腫脹、腫れ、クレピタス、硬直）;胸部（局所的なパチパチ音 vs 喘鳴音 vs 上気道伝達音 - 喘鳴と喘鳴を区別するために気管を聴診する）。混乱を考慮した完全な神経検査。皮膚（発疹、斑点）。
ステップ 3: 精密検査と差分
動作差動、大まかに並べると次のようになります。
ウイルス性 URI を合併する敗血症を伴う細菌性肺炎 — タイムラインに適合します (URI → 4 ～ 5 日目の悪化、悪寒、発汗、錯乱、食欲不振)。
喉頭蓋炎/声門上炎または首深部感染症 - 喉の灼熱感、喉の閉まり、喘鳴にフィットします。大人でも喉頭蓋炎になりますが、ジュニアは小児の病気だと思っているため見逃されます。
細菌の二次感染を伴うインフルエンザまたは新型コロナウイルス感染症。
アナフィ

弛緩/血管浮腫 — 発熱を伴う可能性は低いですが、「喉が閉じる」症状がある場合は尋ねてください。
食欲不振、「内破」けいれん、および混乱は、全身性敗血症によって最も簡潔に説明されますが、ブドウ糖、電解質が得られ、それらが適合しない場合には広範な精密検査のための低い閾値が必要になります。
検査: CBC、CMP (ナトリウム - 肺炎は低ナトリウム血症を引き起こし、混乱の説明にもなります)、乳酸塩、抗生物質投与前の血液培養×2、地域の診療による CRP/プロカルシトニン、呼吸器ウイルスパネル (インフルエンザ/COVID/RSV)、ブドウ糖、呼吸仕事や混乱が続く場合の静脈または動脈の血液ガス、ECG。胸部X線検査。深部空間感染が疑われる場合は、頸部側方軟部組織フィルム、またはできれば造影剤を使用した頸部 CT を使用しますが、これは気道が安全に移動できると判断された場合に限られます。耳鼻咽喉科/EMによる鼻咽頭鏡検査は、声門上を最終的に観察する検査であり、気道を管理する準備ができている人によって行われます。
治療は診断と並行して行われます。これが ED の基本的な考え方です。敗血症スクリーニングが陽性の場合: 培養、1 時間以内の広域抗生物質、プロトコールに従った水分摂取、乳酸の再評価。喉頭蓋炎が疑われる場合：気道チーム、インフルエンザ菌および連鎖球菌を含む抗生物質の静注（セフトリアキソンなど）、およびデキサメタゾンを検討します。アナフィラキシーの疑いがある場合: 抗ヒスタミン薬ではなく、エピネフリンが答えとなります。
ステップ 4: この患者とどのように対話するか
リスト全体を確認してから、優先順位を説明します。 「本当に役立つ写真を提供していただきました。鼻水も問題ですが、私が最も心配しているのは、あなたの呼吸、喉の詰まり、混乱のエピソードです。そこで、まずそこに焦点を当てます。」患者は、自分が導いたことをあなたが無視すると、否定されたように感じます。ある一文がそれを防ぎます。
「これまでの病気とは全く異なる奇妙な問題」を真剣に受け止める

—大声で。患者が「今まで経験したこととは違う」とはっきり言うことは、医療において最も価値のある言葉の一つです。 「それはまさに私が知りたいことのようなものです。あなたは私よりも自分の体のことをよく知っています。」と言ってください。それは信頼を築きます、そしてそれは真実です。
妻を巻き込んでください。彼女は混乱を目の当たりにした。当然のことながら、彼はそれを完全に報告することはできません。ここでは、付随履歴はオプションではありません。
計画と不確実性について正直に話します。 「おそらく、単なる風邪を超えた感染症にかかっていると思います。おそらく肺炎か、気道近くの感染症です。検査を行う予定ですが、待つよりも今すぐ治療を開始します。このような感染症では、早いほうが安全だからです。」
部門内であっても明確なセーフティネット: 「喉がまた閉まりそうになったり、呼吸が苦しくなったりした場合は、すぐに通話ボタンを押してください。症状が治まるまで待たずに。」
ジュニア向けの一行レッスン
患者が話をする順序ではなく、致死率によってトリアージします。喘鳴+喉の閉まり+混乱は、毎回4日間の鼻をかむ回数を上回ります。
解決された危険信号は依然として危険信号です。混乱は消えた。原因となった病気はそうではありませんでした。
気道の緊急事態が疑われる場合、最も危険なことは、日常的な行為です。

[切り捨てられた]

## Original Extract

AI medical diagnosis examples. Contribute to joelparkerhenderson/ai-medical-diagnosis-examples development by creating an account on GitHub.

ai-medical-diagnosis-examples/doctor-perspective/claude-fable-5-ultracode/index.md at main · joelparkerhenderson/ai-medical-diagnosis-examples · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
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
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
joelparkerhenderson
/
ai-medical-diagnosis-examples
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Copy path Blame More file actions Blame More file actions Latest commit
History History 285 lines (183 loc) · 55 KB main Breadcrumbs
Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions AI medical diagnosis example → doctor perspective → Claude Fable 5 Ultracode
This page shows an AI medical diagnosis example, using a doctor perspective,
generated two ways so you can see the differences:
Anthropic Claude Fable 5 Ultracode
To help make progress with conversations about AI medical diagnosis, this
repository contains my own writeup of my own real medical symptoms when I felt
especially unwell. I did the writeup at the start of my real visit to a real
doctor at a real hospital urgent care emergency department.
You are an expert medical doctor.
You are training a junior doctor how to triage patients, how to diagnose
patients, and how to interact with patients.
Explain what a doctor should do when a patient comes into the hospital urgent
care emergency department.
The patient presents a list of symptoms.
Nose:
- Gold-color mucus at start of illness.
- Runny nose then blocked nose, with lots of postnasal drip, 4 days.
- Blowing nose every few minutes, 4 days.
Skin:
- Fever.
- Soaking sweats.
- Shivering chills.
Lungs:
- Shortness of breath for 5 days; hard to breathe in, hard to breathe deep.
- Pervasive semi-productive cough, every few minutes or so.
- Intermittent stridor, wheezing, whistling breath.
Throat:
- My upper throat and back of my mouth feels like it's burning.
- Yesterday, a sensation that I couldn't breathe in, like the wind knocked out of me.
- Today, near-instant throat closing.
Strange issues that are wholly different than any of my previous illnesses:
- Total loss of appetite.
- Spasms of internal contractions, like an imploding feeling, pulling inwards.
- Mental confusion for 15 minutes, bad enough that I couldn't understand my wife.
AI Output
This is a great teaching case, because buried in what sounds like "a bad cold" are several findings that should make the hair on the back of your neck stand up. Let me walk you through how to approach it: triage first, then assessment and diagnosis, then how to talk with the patient. (This is teaching material, not a substitute for institutional protocols or real-time clinical judgment.)
Step 1: Triage — find the killers first
Triage is not about figuring out what the patient has. It's about answering one question: how fast could this patient die? You scan the complaint list for threats to Airway, Breathing, Circulation, and Disability (neuro) before anything else.
This patient hands you four red flags:
"Near-instant throat closing" today + stridor. Stridor means upper airway obstruction until proven otherwise. Combined with a burning throat and a sensation of throat closure, you must immediately think of epiglottitis, deep space neck infection (retropharyngeal or peritonsillar abscess), Ludwig's angina, or angioedema/anaphylaxis . Any of these can occlude the airway in minutes.
Mental confusion severe enough that he couldn't understand his wife. In a febrile patient, acute confusion suggests sepsis, hypoxia, hypercapnia, or CNS infection . Even though it resolved, it counts.
Fever + rigors (shivering chills) + soaking sweats. True rigors strongly suggest bacteremia . This pushes you toward sepsis, not "a virus."
Five days of shortness of breath with a productive cough — possible pneumonia layered on top.
This is not a fast-track patient. This is ESI level 2 at minimum — arguably level 1 if there's any stridor at rest — and goes straight to a resuscitation-capable bed. The teaching point for a junior: the patient led with his runny nose; you must not. Patients tell their story chronologically; triage requires you to reorder it by lethality.
Immediate actions, before any detailed history:
Full vitals including SpO₂, respiratory rate, and temperature. Get a respiratory rate yourself — it's the most commonly fabricated vital sign and the most sensitive early marker of deterioration.
Look at the patient: tripod positioning? Drooling? Muffled "hot potato" voice? Accessory muscle use? Can they speak full sentences? Listen at the neck for stridor at rest.
IV access, continuous monitoring, oxygen if hypoxic.
Do not send this patient alone to radiology, do not lay them flat, and do not aggressively instrument the throat if you suspect epiglottitis — agitation can precipitate complete obstruction. Notify your senior, and have anesthesia/ENT awareness early if airway concern is real. The airway you secure with backup at hand is the one that goes well.
Calculate a quick sepsis screen (qSOFA/NEWS): fever + confusion + tachypnea may already trip it.
Step 2: History and examination — structured, not symptom-by-symptom
Once the patient is monitored and stable enough to talk, take a proper history. Don't just transcribe their list — interrogate it:
Clarify the airway story. "Throat closing" — was it with swallowing, breathing, both? Any voice change, drooling, pain on swallowing (odynophagia), neck pain or stiffness, trismus? Any new medications (ACE inhibitors → angioedema), allergen exposures, or food/sting triggers?
Clarify the confusion. Witnessed account from the wife if possible. Was it gradual or abrupt, any tongue-biting, incontinence, post-event drowsiness (seizure?), or was it more like delirium? Confusion plus the "internal contraction spasms" he describes raises the question of rigors misinterpreted, hypoxic episodes, or even seizure activity — and at the far end of the differential, with a burning throat and spasms you at least file away rabies/tetanus and ask about animal bites and vaccination status. Rare, but the cost of asking is one sentence.
Quantify the breathing. Exertional vs. at rest, orthopnea, pleuritic pain, hemoptysis. "Hard to breathe in" specifically suggests upper airway or pleuritic limitation rather than bronchospasm — note that distinction; wheeze is usually worse on exhalation .
The usual scaffolding: past medical history (asthma? COPD? diabetes? immunosuppression?), medications, allergies, smoking, sick contacts, COVID/flu exposure and vaccination, recent dental work (Ludwig's), travel.
Examination priorities: general appearance and work of breathing; oropharynx (gently — uvula, tonsils, floor of mouth, drooling); neck (lymphadenopathy, swelling, crepitus, stiffness); chest (focal crackles vs. wheeze vs. transmitted upper airway sounds — auscultate over the trachea to distinguish stridor from wheeze); full neuro exam given the confusion; skin (rashes, mottling).
Step 3: Workup and differential
Working differential, roughly ordered:
Bacterial pneumonia with sepsis complicating a viral URI — fits the timeline (URI → day 4–5 deterioration, rigors, sweats, confusion, anorexia).
Epiglottitis / supraglottitis or deep neck space infection — fits the burning throat, throat closing, stridor. Adults get epiglottitis too, and it's missed because juniors think of it as a pediatric disease.
Influenza or COVID-19 with secondary bacterial infection.
Anaphylaxis/angioedema — less likely with fever, but the "throat closing" episodes demand you ask.
The anorexia, "imploding" spasms, and confusion are most parsimoniously explained by systemic sepsis — but they earn a glucose, electrolytes, and a low threshold for broader workup if they don't fit.
Investigations: CBC, CMP (sodium — pneumonia causes hyponatremia, which also explains confusion), lactate, blood cultures ×2 before antibiotics, CRP/procalcitonin per local practice, respiratory viral panel (flu/COVID/RSV), glucose, venous or arterial blood gas if work of breathing or confusion persists, ECG. Chest X-ray; soft-tissue lateral neck film or, better, CT neck with contrast if deep space infection is suspected — but only once the airway is judged safe to travel. Nasolaryngoscopy by ENT/EM is the definitive look at the supraglottis, done by someone prepared to manage the airway.
Treatment runs in parallel with diagnosis — this is the core ED mindset. If the sepsis screen is positive: cultures, broad-spectrum antibiotics within the hour, fluids per protocol, reassess lactate. If epiglottitis is suspected: airway team, IV antibiotics covering H. influenzae and strep (e.g., ceftriaxone), and consider dexamethasone. If any anaphylaxis suspicion: epinephrine is the answer, not antihistamines.
Step 4: How to interact with this patient
Acknowledge the whole list, then explain your priorities. "You've given me a really useful picture. The runny nose matters, but the things I'm most concerned about are your breathing, the throat closing, and the episode of confusion — so we're going to focus there first." Patients feel dismissed when you skip past what they led with; one sentence prevents that.
Take "strange issues wholly different from previous illnesses" seriously — out loud. A patient explicitly telling you "this is not like anything I've had before" is one of the highest-value statements in medicine. Say: "That's exactly the kind of thing I want to know about. You know your body better than I do." It builds trust and it's true.
Involve the wife. She witnessed the confusion; he, by definition, cannot fully report it. Collateral history is not optional here.
Narrate the plan and the uncertainty honestly. "I think you likely have an infection that's gone beyond a simple cold — possibly pneumonia or an infection near your airway. We're going to run tests, but I'm starting treatment now rather than waiting, because with infections like this, earlier is safer."
Safety-net explicitly, even within the department: "If your throat feels like it's closing again, or breathing gets harder, press the call button immediately — don't wait to see if it passes."
The one-line lessons for your junior
Triage by lethality, not by the order the patient tells the story. Stridor + throat closing + confusion outranks four days of nose-blowing every time.
A resolved red flag is still a red flag. The confusion passed; the disease that caused it did not.
In suspected airway emergencies, the most dangerous thing you can do is something routi

[truncated]
