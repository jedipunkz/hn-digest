---
source: "https://openhealthhub.org/t/build-your-own-ai-assisted-personal-health-record/3006"
hn_url: "https://news.ycombinator.com/item?id=48630812"
title: "Build your own AI-assisted personal health record"
article_title: "Build your own AI-assisted personal health record - open forum - openhealthhub.org"
author: "spdegabrielle"
captured_at: "2026-06-22T14:52:48Z"
capture_tool: "hn-digest"
hn_id: 48630812
score: 1
comments: 0
posted_at: "2026-06-22T14:39:46Z"
tags:
  - hacker-news
  - translated
---

# Build your own AI-assisted personal health record

- HN: [48630812](https://news.ycombinator.com/item?id=48630812)
- Source: [openhealthhub.org](https://openhealthhub.org/t/build-your-own-ai-assisted-personal-health-record/3006)
- Score: 1
- Comments: 0
- Posted: 2026-06-22T14:39:46Z

## Translation

タイトル: AI を活用した独自の個人健康記録を作成する
記事のタイトル: AI を活用した独自の個人健康記録を作成する - オープン フォーラム - openhealthhub.org
説明: AI を活用した独自の個人健康記録を作成します
ちょっとした実験として数ヶ月間続けてきたことを書きます。この「ハウツー」は、昨日 @Alan_Hassey @mike.bainbridge @nick.booth @JW148 によってリクエストされました…

記事本文:
= 40rem)" rel="スタイルシート" data-target="デスクトップ" />
= 40rem)" rel="スタイルシート" data-target="chat_desktop" />
= 40rem)" rel="スタイルシート" data-target="discourse-ai_desktop" />
= 40rem)" rel="スタイルシート" data-target="discourse-reactions_desktop" />
= 40rem)" rel="スタイルシート" data-target="poll_desktop" />
= 40rem)" rel="stylesheet" data-target="desktop_theme" data-theme-id="6" data-theme-name="topic-width-desktop"/>
openhealthhub.org
AI を活用した独自の個人健康記録を作成する
パチャラネロ
(マーカス・バウ)
2026年6月20日 13:26
1
AI を活用した独自の個人健康記録を作成する
ちょっとした実験として数ヶ月間続けてきたことを書きます。この「ハウツー」は、昨日の Virtual Mabel の電話で @Alan_Hassey @mike.bainbridge @nick.booth @JW148 によってリクエストされましたが、私は考えました - 誰もが使用できる公共の場所に置いてはどうだろうか?
アイデアは、自分の医療記録のすべてのスクラップをコンピューター上の 1 つのフォルダーに収集し、LLM にその中のパターンを読み取らせ、整理し、パターンを検出させるというものです。すべてがマシン上に残ります。
あなたは、LLM があなたの歴史に関する多くの事実にアクセスできるときの賢さに驚かれるでしょう。このアプローチには重大な事実上の誤りはありませんでした。私は、LLM の幻覚は事実情報が欠如しているときに起こりやすく、LLM はユーザーを喜ばせるためにブラフでギャップを補わなければならないと理論立てています。
必要なもの (1 回限り、約 20 分)
Claude Code のインストール - [拡張機能] パネルから VS Code 拡張機能をインストールする (「Claude Code」を検索) か、そのページからインストーラーを実行します。 Claude のサブスクリプションが必要になります (最初は Pro、月額約 20 ポンドで十分です)。
プライバシー : クロード アカウントで、[設定] > [プライバシー] に移動し、デフォルトでオンになっている [チャットの使用を許可する] 設定をオフにします。

人間的 AI モデルをトレーニングおよび改善するためのコーディング セッション。
すべてのものに対して 1 つのフォルダーを作成します。文書/健康。 VS Code の場合: [ファイル] > [フォルダーを開く] を選択し、フォルダーを選択します。クロード パネル (サイドバーのクロード アイコン) を開きます。クロードは、そのフォルダー内でのみ、そのフォルダー内でのみ読み取りと書き込みができるようになります。
すべての医療データを収集します。病院からの手紙、退院概要、クリニックからの手紙、血液検査結果、スキャンレポート、眼鏡店の処方箋、ワクチン接種/新型コロナウイルスの記録、労働衛生および雇用前の血清学的証明書のPDFをドラッグします。スクリーンショットや紙の手紙の写真でも大丈夫です。
まだ持っていないレコードを取得します。あなたの歴史のほとんどはNHSと民間システムに保存されています。クロードに、あなたが利用した各一般開業医、信託、クリニックに対する被験者アクセス要求 (英国 GDPR 第 15 条) の草案を作成するよう依頼してください。レポートだけでなく画像検査も依頼してください。 Claude ブラウザ拡張機能を使用すると、醜い Web UI (NHS アプリなど) にある場合でも、データの抽出を簡単に自動化できます。
クロードに整理を依頼してください。 PDF を読み取り、事実を抽出し、どのエディタでも永久に読み取れるプレーンな Markdown ファイルを書き込みます。次のような最初のプロンプトは機能します。
このフォルダー内のすべてのファイルを読み取ります。
各 PDF レターを、オリジナルの隣にある構造化された Markdown ファイルに転記します。
何も削除しないでください。
それから、現在の問題リスト、投薬リスト、血液検査結果の時系列表を作成してください
ルールを決めるのはあなたです！完全に時系列の記録、または問題指向の記録、またはハイブリッド アプローチを選択することもできます。薬リストなど、必要なあらゆる種類の要約文書を作成できます。
洞察力のために私のものです。これで記録が構造化されました。次のように質問してください。「私の血圧の傾向をプロットしてください」、「実際に記録されている免疫力と推測されている免疫力は何ですか?」 、「どの結果が漂流したか」

十年？」 、「追いかけるべきものは何ですか？」ここで、それはファイルキャビネットではなくなり、ツールになります。
原本を保管してください。 LLM は、各ソース PDF の代わりにではなく、その横に構造化コピーを書き込みます。派生したすべての事実は、そのソースまで追跡可能です。
複雑な臨床計算機には注意してください。単純なスコアリング ツール (BMI など) はおそらく正確です。ただし、より複雑なもの (例: QRISK2-3) の場合、LLM が正確に計算する能力があるかどうかは信頼できません。 Web ベースの計算機は、ブラウザ拡張機能を介して使用できる可能性があります。 (これは、私が取り組んでいる、高速で LLM フレンドリーなコマンドライン臨床計算機スイートの必要性を強調しています [1] [2] )
継続的にバックアップします。理想的には、自動同期ツールを使用して、ある種の安全なプライベート クラウド (Dropbox や Google Drive など) に接続します。
珍しい病気にはかなり強いようです。トレーニング データが豊富であるため、経験豊富な臨床医が「点と点を結び付けていない」場合でも、まれな疾患の可能性を指摘するのに非常に優れています。臨床医は、特定の症状の考えられる稀な遺伝的原因をすべて知ることはできません。 LLM は当然のことながらそうします。
あなたに関する最も重要な重要な医学的事実を凝縮したリスト、つまり「緊急時用」ウォレットの概要を求めて、それがどの程度効果があるかを確認してください。
存在すると思われるギャップや矛盾を埋めるために質問するように依頼します。
今後数回の一般医相談で対処する必要がある事項についての推奨事項を尋ねてください。
包括的な記録の基本を収集することがいかに簡単であるかに私は驚いています。実際、この非常に初歩的な形式であっても、それが他の誰のあなたの医療記録よりも包括的なものであることに気づき始めています。あなたの主治医もおそらくそうしているでしょう

それほど完全な記録を持っているわけではありませんし、簡単に参照、読み取り、検索できるようにすべての記録を持っているわけではありません。 AVT？まともなLLMの無駄。
Discourse によって提供されており、JavaScript を有効にして表示するのが最適です

## Original Extract

Build your own AI-assisted personal health record
A write-up of something I’ve been doing for a few months as a bit of an experiment. This ‘How-to’ requested by @Alan_Hassey @mike.bainbridge @nick.booth @JW148 on yesterd&hellip;

= 40rem)" rel="stylesheet" data-target="desktop" />
= 40rem)" rel="stylesheet" data-target="chat_desktop" />
= 40rem)" rel="stylesheet" data-target="discourse-ai_desktop" />
= 40rem)" rel="stylesheet" data-target="discourse-reactions_desktop" />
= 40rem)" rel="stylesheet" data-target="poll_desktop" />
= 40rem)" rel="stylesheet" data-target="desktop_theme" data-theme-id="6" data-theme-name="topic-width-desktop"/>
openhealthhub.org
Build your own AI-assisted personal health record
pacharanero
(Marcus Baw)
20 June 2026 13:26
1
Build your own AI-assisted personal health record
A write-up of something I’ve been doing for a few months as a bit of an experiment. This ‘How-to’ requested by @Alan_Hassey @mike.bainbridge @nick.booth @JW148 on yesterday’s Virtual Mabel's call, but I thought - why not put in in a public place for all to use?
The idea : collect every scrap of your own medical record into one folder on your computer, then let an LLM read, organise, and find patterns in it. Everything stays on your machine.
You’ll be surprised at how astute the LLM can be when it has access to lots of facts about your history. I haven’t noticed any significant factual inaccuracies with this approach. I theorise that LLM hallucination is more likely to occur when there is an absence of factual information, and LLMs have to make up the gaps with bluff in order to please you.
What you need (one-off, ~20 min)
Install Claude Code - install the VS Code extension from the Extensions panel (search “Claude Code”), or run the installer from that page. You will need a Claude subscription (Pro, at about £20 a month, is fine to start with).
Privacy : In your Claude account, got to Settings > Privacy and switch OFF the default-on setting ‘Allow the use of your chats and coding sessions to train and improve Anthropic AI models’.
Make one folder for everything, e.g. Documents/health . In VS Code: File > Open Folder and pick it. Open the Claude panel (the Claude icon in the sidebar). Claude can now read and write inside that folder, and only that folder.
Collect all your medical data . Drag in PDFs of hospital letters, discharge summaries, clinic letters, blood results, scan reports, optician prescriptions, vaccination/COVID records, occupational-health and pre-employment serology certs. Screenshots and photos of paper letters are fine too.
Get the records you don’t have yet. Most of your history sits in NHS and private systems. Ask Claude to draft a Subject Access Request (UK GDPR Article 15) to each GP practice, trust, and clinic you have used. Ask for imaging studies , not just the reports. Using the Claude Browser Extension you can easily automate extraction of data even when it’s in a hideous web UI (like the NHS App)
Ask Claude to organise it. It will read the PDFs, extracts the facts, and writes plain Markdown files you can read in any editor forever. A first prompt of anything like this will work:
Read every file in this folder.
Transcribe each PDF letter into a structured Markdown file next to the original.
Don't delete anything.
Then build me a current problem list, medication list, and a chronological blood-results table
You decide the rules! You can go for a completely chronological record - or a Problem Oriented Record, or some hybrid approach. You can have any kinds of summary documents you want - medication lists
Mine for insight. Now the record is structured, ask questions across it: “Plot my blood pressure trend” , “What immunity is actually documented vs inferred?” , “Which results have drifted over ten years?” , “What’s missing that I should chase?” This is where it stops being a filing cabinet and becomes a tool.
Keep originals. The LLM writes a structured copy alongside each source PDF, never instead of it. Every derived fact stays traceable to its source.
Beware complex clinical calculators. Simple scoring tools (eg BMI) will likely be accurate. However for more complex ones, (eg QRISK2-3) I would not trust that the LLM has the ability to calculate it accurately. Web-based calculators might be something it can use via the Browser Extension. (This highlights the need for a suite of fast, LLM-friendly command-line clinical calculators, which I am working on [1] [2] )
Backup continuously. Ideally to some kind of secure, private cloud - for example Dropbox or Google Drive, using an automatic sync tool.
It seems to be quite good at rare diseases. Due to the breadth of its training data, it is quite good at pointing out the possibility of a rare condition, even when experienced clinicians haven’t ‘connected the dots’. Clinicians simply can’t know about every possible rare genetic cause of a given symptom. LLMs naturally just do.
Ask for a condensed list of the most important key medical facts about you - the ‘In Case Of Emergency’ wallet summary - and see how well it does.
Ask it to ask you questions to fill in any gaps or inconsistencies that it considers to be present.
Ask it for recommendations for things that need to be addressed at your next few GP consultations!
I’ve been surprised by how easy it’s been to accrue the basics of a comprehensive record, and actually what you start to realise is that even in this very rudimentary form, it’s still more comprehensive than anyone else’s medical record about you! Even your GP probably doesn’t have as complete a record, and they absolutely certainly don’t have it all at their fingertips, easy to browse, read, and search. AVTs? waste of a decent LLM.
Powered by Discourse , best viewed with JavaScript enabled
