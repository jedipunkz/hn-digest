---
source: "https://webscraper.io/blog/how-web-scraper-was-translated-into-50-languages-with-ai"
hn_url: "https://news.ycombinator.com/item?id=49282505"
title: "Web Scraper was translated into 50 languages with AI"
article_title: "How Web Scraper was translated into 50 languages with AI | Web Scraper"
author: "martinsbalodis"
captured_at: "2026-08-13T07:13:06Z"
capture_tool: "hn-digest"
hn_id: 49282505
score: 2
comments: 1
posted_at: "2026-08-13T06:45:53Z"
tags:
  - hacker-news
  - translated
---

# Web Scraper was translated into 50 languages with AI

- HN: [49282505](https://news.ycombinator.com/item?id=49282505)
- Source: [webscraper.io](https://webscraper.io/blog/how-web-scraper-was-translated-into-50-languages-with-ai)
- Score: 2
- Comments: 1
- Posted: 2026-08-13T06:45:53Z

## Translation

タイトル: Web Scraper を AI で 50 か国語に翻訳
記事タイトル: Web Scraper を AI で 50 か国語に翻訳する方法 |ウェブスクレーパー
説明: Web Scraper Chrome 拡張機能を AI を使用して 50 の言語に翻訳し、人間の翻訳品質と同等の品質を 5 分の 1 のコストで実現した方法。

記事本文:
ウェブスクレーパークラウド
複雑なスクレイピング スタックを置き換えます。
-
7日間無料でお試しください
クラウドログイン
ウェブスクレーパークラウド
Web Scraper を AI で 50 か国語に翻訳した方法
当初、Web Scraper 拡張機能は英語でのみ利用可能でした。英語は世界で 3 億 7,200 万人の母語話者が話されており、世界で 3 番目に話されている母語に過ぎません。一方、スペイン語は 4 億 8,700 万人、中国語は 9 億 8,800 万人です。使いやすさを向上させ、誤用を減らすために、Web Scraper 拡張機能を国際化することにしました。当初はSaaSを利用する予定でしたが、最終的にはすべてAIで翻訳することになりました。
翻訳ソフトウェアが重要なのはなぜですか?
Chrome 拡張機能開発者ダッシュボードには、国ごとのユーザー数に関する情報が表示されます。この情報を使用して、100 万人あたりの Web Scraper 拡張機能ユーザーの数を計算できます。米国、英国、オーストラリアなどの英語を母国語とする国では、人口 100 万人あたり 350 人以上のユーザーがいますが、イタリアや日本のような英語の普及率が 28% にすぎない国では、ユーザー数は 200 ～ 250 人です。私たちの分析によると、ある国での英語の普及は 100 万人あたりのユーザー数と相関しています。その他の例については、以下の表を参照してください。
当初、私たちは翻訳を実行するために SaaS を使用することに決めました。私たちは、文字列ファイルをシステムと同期するスクリプトを作成しました。そこでは言語を初期化し、機械翻訳や AI を介して翻訳したり、人間の専門家を雇ったりすることができます。私たちのチーム全体がラトビア語を話しているため、ラトビア語から始めました。ラトビア語は英語よりもはるかに複雑です。単語には性別があり、最大 7 つの異なる語尾が付けられます。これは、文の中で複数の単語を組み合わせる場合、特定の語尾をとらなければならないか、文がまったく意味をなさないことを意味します。の

初期の機械翻訳と AI 翻訳は悪かったか、一貫性がありませんでした。機械翻訳では語尾が混同されてしまうため、AI が正しく理解できるのは半分ほどでした (その時点では Opus 4.8 を使用していました)。最終的には人間の翻訳者を使用することにしました。
ラトビア語の語尾の例:
AI にとって、適切なコンテキストが結果を左右することもありますが、i18n にとって、それは不可欠です。私たちは、AI 機能を試すためと、依然として AI と機械翻訳を出発点として使用している人間の翻訳者のためのあいまいさを取り除くためのコンテキストを設定しました。設定内容は次のとおりです。
ビジネスの説明 - 製品の簡単な説明。
トーン - ユーザーへの対応方法に関するルール。たとえば、ユーザーを「あなた」と呼ぶことは避けてください。
単語辞書 - 製品固有の用語の説明。たとえば、スクレーパーはウェブ スクレーパーまたはタング スクレーパーです。辞書はどちらを意味するのかを明らかにします。
SaaS人間による専門翻訳
すべてのセットアップが完了した後、ラトビア語とドイツ語の人による翻訳をテストしました。翻訳者はまず辞書を翻訳し、翻訳中に新しい単語を追加し続けました。翻訳された文字列が戻ってきてチェックを開始すると、翻訳は流暢ではあるものの、口調に一貫性がなく、英語の単語が曖昧な場合には誤訳がいくつかあることがわかりました。結果として、私たちは 100% 満足できませんでした。
私たちがプロによる人間による翻訳をテストしている間、AI の世界では多くのことが起こっていました。Fable モデルが発表され、OpenRouter は 3 つのモデルを組み合わせてベンチマークで Fable よりも高いスコアを獲得する Fusion を発表しました。試してみなければなりませんでした。 SaaS プラットフォームではこれにアクセスできなかったので、このセットアップをローカルで構築しました。 AI 翻訳は反復作業であり、新しい知識があったため各ステップがやり直されました。

前のステップでの変更または変更。以下に、複数回実行した手順を示します。
私たちはオリジナルの文字列に満足していましたが、トーン定義を作成した後、AI にそれらをレビューしてより良い表現を提案するよう依頼しました。そうなったので、いくつかの変更を加えました。
初めて AI をテストしたとき、あいまいな翻訳が多かったので、問題のある文字列に説明を追加しました。これにより、翻訳は正しい方向に進みました。また、文字列がボタンのラベルの場合はできるだけ短くする必要があることに注意することも重要でした。最終バージョンでは、AI にソース コード全体を調べて、すべての文字列に説明を追加するよう依頼しました。手作業によるレビューの結果、300 件の説明のうち約 10 件が調整されました。
当初、トーンの定義では、専門的な IT 言語を使用する必要があると記載され、翻訳すべきではない単語がリストされていました。 AI と人間による翻訳を比較した後、私たちは AI に定義を更新するよう依頼し、より気に入った翻訳の例を与えました。「X は、このように翻訳されるべきであり、そのように翻訳されるべきではありません。」
辞書は、1 つの単語を 1 つの翻訳に固定し、曖昧さを取り除くため、不可欠です。当初、これには、ラトビア語に誤って翻訳されていた特定の単語を Web スクレイピングして含まれていました。 AI 翻訳の最初の反復後、他の言語ではあいまいな単語をさらに追加するよう AI に依頼しました。辞書がかなり増えました。
辞書は、OpenRouter の Fusion モデルを使用したスクリプトで翻訳されました。英語辞書や語調の定義が変更されたため、翻訳は何度もやり直しされました。
追加の入力として、すべての言語のすべての文字列を機械翻訳で翻訳しました。 AI が何かを見逃した場合、機械翻訳が追加のコンテキストを提供できることが期待されていました。
S

文字列の変換は、次のスクリプトを使用して実行されました。
Fusion モデルを使用すると、入力の変更に応じて文字列が複数回翻訳されました。
トーンと辞書を設定し、Fusion モデルを使用して AI 翻訳を反復処理した後、間違った語尾はなくなりました。これは、競争の場がより平等になったことを示す良い兆候です。ラトビア語とドイツ語について AI と人間の翻訳者を比較しました。
人間の間違いのほとんどは、英語の文字列を誤解したり、望ましい音調から逸れたりしたことに起因します。 AI の場合、明確な悪いパターンはありませんでした。 AI 翻訳がまだ十分に優れているにもかかわらず、単純に人間による翻訳の方が優れていると感じることがありました。
ご覧のとおり、明確な勝者も明確な敗者もいません。
最終的な AI 翻訳の後、もう一度人間による翻訳テストを実行しました。今回はプロの翻訳者に AI 翻訳へのアクセスを許可しました。彼らは問題を修正するか、翻訳を現状のままにしておくだけで済みました。これはフランス語用に行いました。 307 個のテスト文字列のうち、228 個はそのまま残されました。私たちのチームにはフランス語を母国語とする人がいないため、最終レビューの実行には AI を活用しました。結果の一部を次に示します。
AI 比較の結果は次のとおりです。
結論: AI セットは全体的により強力な翻訳です。マルチエージェントによるブラインド レビューで 79 行中 60 行を獲得し、検証された重大な欠陥 30 個のうち 27 個は人間のセットに属します。しかし、人間セットは、AI があらゆるところで間違う 1 つの重要な点について体系的に正しいのです。それは、ガイダンス テキストでは命令形を使用する必要があるという TONE.md のルールです。フランスで最も優れたロケールは、人間のアプローチから借用した、約 12 個のターゲットを絞った修正を含む AI セットでしょう。
最上位の AI モデルは人間の翻訳者と競合することができ、AI 翻訳は重大なエラーを起こすことなく実行できます。これを書いている時点では

より新しく優れた AI モデルがすでに展開されており、今後は改善するしかありません。
人間の翻訳者を使用すると 7,500 ドルの費用がかかり、AI トークンに約 1,500 ドルを費やしましたが、この記事のような記事が存在していれば、その半分を費やしたでしょう。
私たちの当初の計画では、ベース翻訳には人間の翻訳者を使用し、新しい文字列には機械翻訳を使用する予定でした。リリースごとに 10 個の新しい文字列を専門家が翻訳するのは時間もコストもかかりすぎるためです。高品質の AI 翻訳を使用すれば、これはまったく問題になりません。
最終的には、ラトビア語、ドイツ語、フランス語、スペイン語を人力で翻訳しましたが、いずれにせよ物事の方向性はここにあると感じたため、AI のみの翻訳にこだわることにしました。
この記事はおそらく誰かの AI への入力として使用されることになるため、開始点として使用できるトーンの定義を次に示します。
Web Scraper Chrome 拡張機能は、Web スクレイピング用のユーティリティ アプリです。それはユーザーに許可します
ウェブサイトからデータを抽出するため。
Web スクレイパー拡張機能には、AI ウィザードとアドバンスト UI の 2 つのユーザー インターフェイスがあります。アイ
拡張機能アイコンをクリックするとウィザードにアクセスでき、サイドが開きます。
パネル。高度な UI には、ブラウザの開発ツールを介してアクセスできます。
短いテキストを翻訳する場合は、翻訳を短くしてください。それ以外の場合はテキスト
UIに適合しない可能性があります。
翻訳は IT プロフェッショナルと Web 開発者を対象としています。テクニカルを使用する
必要に応じて用語を使用します。
「削る」、「削る」、「削る」、「削った」などの単語を翻訳する場合、
"要素"、"ポイント アンド クリック"、"セレクター"、"親セレクター"、"ウィザード"、
「ブラウザ」は、Web スクレイピングのコンテキストでの意味です。この言葉を翻訳すると、
ただし、これらは Web スクレイピングに関連していることに注意してください。
ブランド基盤
中心となる対象者: 技術ユーザー、開発者、および対話する専門家
ソフトウェアの

表面。 UI テキストの明瞭さ、効率性、正確さを重視します。
知識レベル: 標準的なソフトウェア規約に精通していることを前提としています。必要ありません
基本的な UI パターンについて説明します。
核となる価値観: 客観性と明確さ。主観を入れずに情報を提示する
フレーミングまたは会話要素。
ブランドのペルソナ: 中立的なシステム インターフェイス - 切り離された、有益で機能的な、
優れたデザインのコントロールパネルのようなものです。
声と口調
主な声: 非個人的、客観的、そして行動指向。
トーンマトリックス
UI とレイアウト コンテキスト: 分離されたアクション ラベルを使用します。不定詞または現在を使用する
緊張した。実行: 開発者ツールを開きます。してはいけないこと: 開発者ツールを開いてください。
エラー メッセージ コンテキスト: 中立状態レポートを使用します。条件を付けずに述べます
ユーザーに話しかける。実行: ファイル形式が無効です。してはいけないこと: 無効なファイルをアップロードしました
ファイル。
ドキュメントのコンテキスト: 説明的な機能ラベルを使用します。どのようなアクションかを説明する
する。実行: 開発者ツールを開きます。禁止: 開発者ツールが開きます。
進行状況とステータスのコンテキスト: 進行中の業務をアクティブかつ現在形でレポートします。
動詞。アクターを暗黙的に保ち、受動的なものを避ける
[切り捨てられた]
ブラウザ拡張機能のプライバシー ポリシー
ウベル 5-71、
アダジ、ラトビア、LV-2164
Copyright © 2026 ウェブスクレイパー |無断転載を禁じます

## Original Extract

How we translated the Web Scraper Chrome extension into 50 languages with AI, matching human translation quality at a fifth of the cost.

Web Scraper Cloud
replaces your complex scraping stack.
-
Try it free for 7 days
Cloud login
Web Scraper Cloud
How Web Scraper was translated into 50 languages with AI
Originally, the Web Scraper extension was only available in English. English is only the third most spoken native language in the world, with 372 million native speakers, while Spanish has 487 million and Chinese has 988 million. To improve usability and reduce misuse, we decided to internationalize the Web Scraper extension. The initial plan was to use a SaaS, but in the end, everything was translated by AI.
Why is translating software important?
The Chrome extension developer dashboard provides information about user count per country. Using this information, we can calculate how many Web Scraper extension users there are per million people. In native English-speaking countries like the United States, the United Kingdom, and Australia, there are 350+ users per million people, while in countries like Italy and Japan, where English adoption is only 28%, the user count is 200 - 250. From our analysis, English adoption in a country correlates with the user count per million people. See the table below for more examples.
Initially, we decided to use a SaaS to perform the translations. We created scripts that would synchronize our string file with their system. There we could initialize languages and translate them via machine translation, AI, or hire human professionals. We started with Latvian, since it is the language our entire team speaks. Latvian is a lot more complex than English: a word can have a gender and up to seven different endings. In a sentence, this means that when multiple words are combined, they have to take specific endings or the sentence doesn't make any sense at all. The initial machine and AI translations were bad or inconsistent. Machine translation would mix up the word endings, and AI got them right only about half the time (at that point, we were using Opus 4.8). In the end, we decided to use human translators.
Example of Latvian language endings:
For AI, good context is sometimes what makes or breaks the result, and for i18n it is essential. We set up the context both to try out the AI feature and to remove any ambiguity for human translators, who still use AI and machine translation as a starting point. Here is what we set up:
Business description - a brief description of the product.
Tone - rules for how the user should be addressed. For example, avoid addressing the user as "you".
Word dictionary - descriptions for product-specific terms. For example, a scraper can be a web scraper or a tongue scraper; the dictionary clarifies which one is meant.
SaaS human professional translation
After everything was set up, we tested human translation for Latvian and German. The translator first translated the dictionary, then kept adding new words to it during the translation. When the translated strings came back and we started checking them, we found that while the translations were fluent, the tone was inconsistent, and there were some incorrect translations when an English word was ambiguous. As a result, we weren't 100% satisfied.
While we were testing professional human translation, a lot was happening in the AI world: the Fable model launched, and OpenRouter announced Fusion, which combines three models to score even higher on benchmarks than Fable. We had to try it. We didn't have access to this in the SaaS platform, so we built this setup locally. The AI translation was iterative work, where each step was redone because of new learnings or changes in previous steps. Here are the steps we performed, some of them multiple times.
While we were happy with our original strings, after creating the tone definition we tasked AI with reviewing them and suggesting better wording. It did, and we made a few changes.
When we first tested AI, there were a lot of ambiguous translations, so we added descriptions to the problematic strings. This steered the translations in the right direction. It was also essential to note when a string is a button label and should be kept as short as possible. For the final version, we tasked AI with going through the entire source code and adding a description for every string. After a manual review, about 10 out of 300 descriptions were adjusted.
Initially, the tone definition stated that professional IT language should be used and listed the words that shouldn't be translated. After comparing AI and human translations, we tasked AI with updating the definition, giving it examples of the translations we liked better - "X should be translated like this, not like that."
A dictionary is essential because it pins one word to one translation and removes any ambiguity. Initially, it contained web scraping specific words that were being translated into Latvian incorrectly. After the first iteration of AI translation, we tasked AI with adding more words that were ambiguous in other languages. The dictionary grew a lot.
The dictionaries were translated with a script that used OpenRouter's Fusion model. The translations were redone multiple times as the English dictionary and the tone definition changed.
As an additional input, we translated all strings in all languages with machine translation. The hope was that if AI missed something, machine translation could provide extra context.
String translations were performed with a script that took:
Using the Fusion model, the strings were translated multiple times as the inputs changed.
After setting up the tone and dictionaries and iterating on the AI translations with the Fusion model, there were no incorrect word endings - a good signal that the playing field had become a lot more equal. We compared AI against a human translator for Latvian and German.
Most human mistakes came from misunderstanding the English string or drifting away from the desired tone. For AI, there were no clear bad patterns; sometimes we simply felt that the human translation was better, even though the AI translation was still good enough.
As you can see, there is no clear winner and no clear loser.
After the final AI translation, we ran one more human translation test, this time giving the professional translator access to our AI translations: they only had to fix issues or keep the translation as is. We did this for the French language. Out of the 307 test strings, 228 were kept as is. Since no one on our team is a native French speaker, we turned to AI to perform the final review. Here are some of the results:
Here is the verdict from the AI comparison:
Bottom line: the AI set is the stronger translation overall - it won 60 of 79 rows in a blind multi-agent review, and 27 of the 30 verified critical defects belong to the human set. But the human set is systematically right about one important thing the AI gets wrong everywhere: TONE.md's rule that guidance text must use the imperative. The best French locale would be the AI set with roughly a dozen targeted fixes borrowed from the human's approach.
Top-tier AI models can compete with human translators, and AI translation can be done without making bad errors. As of writing this, newer and better AI models have already rolled out - and they can only improve from here.
Using human translators would have cost us $7,500, while we spent about $1,500 on AI tokens - and we would have spent half of that if an article like this one had existed.
Our initial plan was to use human translators for the base translation and machine translation for new strings, since having 10 new strings translated by professionals on every release would be too time consuming and too costly. With high quality AI translations, this is no longer a problem at all.
In the end, we had human translations for Latvian, German, French, and Spanish, but we decided to stick with AI-only translations, since we feel this is where things are heading anyway.
Since this article will probably end up as input for somebody's AI, here is our tone definition, which you can use as a starting point:
Web Scraper Chrome extension is an utility app for web scraping. It allows users
to extract data from websites.
Web scraper extension has two user interfaces - AI wizard and Advanced UI. Ai
wizard is accessible by clicking the extension icon and it will open a side
panel. The Advanced UI is accessible via browser devtools.
When translating a short text, keep the translation short. Otherwise the text
might not fit in the UI.
Translations should be for IT professionals and web developers. Use technical
terms when appropriate.
When translating words like "scraping", "scraper", "scrape", "scraped",
"element", "point and click", "selector", "parent selector", "wizard",
"browser" they are meant in the context of web scraping. Translate these words,
but keep in mind they are related to web scraping.
BRAND FOUNDATIONS
Core Audience: Technical users, developers, and professionals who interact with
software interfaces. Value clarity, efficiency, and precision in UI text.
Knowledge Level: Assume familiarity with standard software conventions. No need
to explain basic UI patterns.
Core Values: Objectivity and Clarity. Present information without subjective
framing or conversational elements.
Brand Persona: A neutral system interface-detached, informative, and functional,
like a well-designed control panel.
VOICE AND TONE
Primary Voice: Impersonal, objective, and action-oriented.
TONE MATRIX
UI & Layout Context: Use detached action labels. Use infinitive or present
tense. Do: Open developer tools. Don't: You should open the developer tools.
Error Messages Context: Use neutral state reporting. State the condition without
addressing the user. Do: Invalid file format. Don't: You uploaded an invalid
file.
Documentation Context: Use descriptive function labels. Describe what actions
do. Do: Opens developer tools. Don't: This will open your developer tools.
Progress & Status Context: Report ongoing operations with active, present-tense
verbs. Keep the actor implicit and avoid passive
[truncated]
Browser Extension Privacy Policy
Ubelu 5-71,
Adazi, Latvia, LV-2164
Copyright © 2026 Web Scraper | All rights reserved
