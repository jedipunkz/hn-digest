---
source: "https://www.vincentschmalbach.com/make-ai-writing-sound-human-logit-bias/"
hn_url: "https://news.ycombinator.com/item?id=49106070"
title: "I Tried to Make AI Writing Sound Human by Banning AI Words Through Logit_bias"
article_title: "I Tried to Make AI Writing Sound Human by Banning AI Words Through logit_bias - Vincent Schmalbach"
author: "vincent_s"
captured_at: "2026-07-30T04:48:45Z"
capture_tool: "hn-digest"
hn_id: 49106070
score: 1
comments: 0
posted_at: "2026-07-30T04:15:23Z"
tags:
  - hacker-news
  - translated
---

# I Tried to Make AI Writing Sound Human by Banning AI Words Through Logit_bias

- HN: [49106070](https://news.ycombinator.com/item?id=49106070)
- Source: [www.vincentschmalbach.com](https://www.vincentschmalbach.com/make-ai-writing-sound-human-logit-bias/)
- Score: 1
- Comments: 0
- Posted: 2026-07-30T04:15:23Z

## Translation

タイトル: Logit_bias を通じて AI の単語を禁止することで、AI の文章を人間らしく聞こえるようにしてみました
記事のタイトル: logit_bias を通じて AI の単語を禁止することで、AI の文章を人間らしく聞こえるようにしてみました - Vincent Schmalbach
説明: モデルが特定のトークンを選択する確率を変更する API 設定を使用して、AI の記述をより人間的に聞こえるようにしようとしました。言葉からブラックリストを作成しました…

記事本文:
logit_bias を通じて AI の単語を禁止することで、AI の文章を人間らしいものにしようとした - Vincent Schmalbach
コンテンツにスキップ
ソフトウェア開発者
ホーム
サービス
カスタム Laravel アプリ開発
Laravelのパフォーマンスの最適化
技術チームのリーダーシップとメンタリング
Laravel開発コンサルティング
logit_bias を通じて AI の言葉を禁止することで、AI の文章を人間らしいものにしようとした
2026 年 7 月 29 日
ヴィンセント・シュマルバック著
私は、モデルが特定のトークンを選択する確率を変更する API 設定である logit_bias を使用して、AI の記述をより人間的に聞こえるようにしようとしました。 AI が生成したテキストに異常に頻繁に出現する単語からブラックリストを作成し、その単語をトークン ID に変換し、それらの ID に負の値を割り当てました。
私は、タスクをプル リクエストに変換するための AI コーディング ツールである vroni.com で作業しながらこれを書いています。
これは、プロンプトに「これらの単語を使用しないでください」と入れることとは異なります。プロンプトはモデルに指示を与えます。モデルはそれを記事と一緒に読み取り、それを適用する方法を決定し、別の指示がより重要であると思われる場合はそれを無視することがあります。 logit_bias は尋ねません。次のすべてのトークンを選択するために使用される数値スコアが変更されます。
私のアイデアは単純でした。AI の記事が同じ単語が出現し続けることによって認識されやすいのであれば、API レベルでそれらの単語を削減できるかもしれません。私は、その主張をそのまま維持しながら、AI の散文に関連する単語を減らして、AI が生成した記事を書き直すモデルを望みました。
それをテストするために、AI が生成した 8 つの完全な記事から始めました。私は各記事全文を OpenRouter 経由で DeepSeek V4 Flash に送信し、すべての事実、名前、番号、例、資格、推奨事項、結論を保持しながら文章を編集するよう依頼しました。記事を段落ごとに処理したわけではありません。
各ソース記事は、m を経て、

二度オーデル。どちらの呼び出しでも、同じ編集プロンプトとサンプリング設定を使用しました。制御呼び出しではトークン バイアスは使用されませんでした。実験的な呼び出しにより、ブラックリストのトークン ID に -8 のバイアスが追加されました。これにより、同じ 8 つのソース記事に対して 8 つのコントロールの書き換えと 8 つのブラックリストアクティブな書き換えが発生しました。
言語モデルは、次のトークンを書き込む前に、選択できるすべてのトークンにスコアを割り当てます。そのスコアはロジットと呼ばれます。 API はそれらのスコアを確率に変換し、次のトークンをサンプリングします。
logit_bias は、これら 2 つのステップの間に別の数値を挿入します。 OpenAI の Chat Completions リファレンスでは、これを -100 から 100 までの値へのトークン ID のマップとして定義しています。 API は、サンプリング前に、指定された値をモデルのスコアに追加します。負の値を指定すると、トークンが選択される可能性が低くなります。 -100 に近い値を指定すると、実質的に禁止されます。
モデルが「これは重要な違いです」と書こうとしているとします。 「重要」のトークンが負のバイアスを受けると、別の継続の可能性が高くなります。モデルは「重要な」を選択するか、形容詞を削除するか、別の文を構築する可能性があります。
それは一度だけではありません。選択されたトークンはすべて、次のトークンのスコア付けに使用されるコンテキストの一部になります。したがって、1 回の強制置換により、文の残りの部分とそれ以降の文も変更される可能性があります。
また、単語はトークンにきちんとマッピングされません。文頭の「Crucial」は、スペースの後の「Crucial」とは異なるトークン ID を持つ場合があります。大文字の形式はまた異なる場合があります。一部の単語は複数のトークンに分割されます。共有フラグメントにペナルティを与えると、ブラックリストに載せられるはずのなかった単語に影響が出る可能性があります。
プロンプトブラックリストは同じものではありません
トークン化が必要ないため、プロンプトに単語をリストする方が簡単です。また、モデルにある程度の裁量権も与えられます。

「らしい」がリストに載っているが情報源が不確かな主張をしている場合、ブラックリストに完全に従うよりも主張を維持することの方が重要であるため、モデルはヘッジを維持することができます。
欠点は、迅速な指示には強制力がないことです。モデルは、特に他の多くの要件を伴う長い記事で、リストされた単語を使用する可能性があります。また、機械的に従って、通常の単語を厄介な同義語に置き換えることもあります。
これを GPT-5.6 Sol で個別にテストしましたが、 GPT-5.6 Sol では logit_bias を受け入れません。 5 つのコントロール記事では、通常の編集プロンプトが使用されました。 5 つの実験呼び出しでは、同じタスクに加えて明示的な 93 単語の抑制リストが使用されました。
プロンプト リストにより、ウォッチワードの出現数が 5 つの制御出力全体で 11 件から 5 つの実験出力全体で 1 件に減少しました。しかし、5 つの実験記事のうち 2 つは情報を削除または変更しました。 5 つのコントロールはすべてソースの意味を保持していました。
モデルとソースセットが異なるため、これは DeepSeek 実験との直接比較ではありませんでした。ブラックリストをプロンプトに移動するだけではトレードオフが解消されないことが示されました。モデルはリストに従っても、物品に損傷を与える可能性があります。
どの API とモデルがサポートしているか
実際のエンドポイントをテストせずに信頼できる単一のモデル リストはありません。
OpenAI には依然として Chat Completions スキーマに logit_bias が含まれていますが、それは現在の OpenAI モデルがそれを受け入れることを意味するものではありません。互換性プローブを GPT-5.2 チャット、GPT-5.3 チャット、GPT-5.4、GPT-5.4 Mini、GPT-5.4 Nano、GPT-5.5、GPT-5.6 Sol、GPT-5.6 Terra、および GPT-5.6 Luna に送信しました。すべてが unsupported_pa​​rameter エラーを返しました。現在の Responses API もフィールドを公開していません。
Anthropic の Messages API にはネイティブの logit_bias パラメータがありません。クロード プロンプトには単語リストを含めることができますが、API は単語リストを含めません。

同じトークン ID スコア調整を提供します。
OpenRouter は、多くのオープンでホストされたモデルにルーティングするため、より広範なサポートを備えています。その Models API は、supported_pa​​rameters によってモデルをフィルタリングできます。 2026 年 7 月 29 日、logit_bias のフィルタリングにより 115 個のモデル ID が返されました。リストには、DeepSeek V4 Flash、GLM 5.2、Kimi K3、Gemma 4 31B、Qwen 3.6 27B、Llama 3.3 70B、Mistral Small 3.2 24B、GPT-4o、およびいくつかの古い GPT-4o バージョンが含まれていました。
そのリストでさえ、特定のリクエストが機能することを証明するものではありません。 OpenRouter は、一部のモデルに対して複数のプロバイダーを組み合わせています。そのルーティングのドキュメントには、require_parameters が有効になっていない限り、プロバイダーはリクエストを受信して​​も不明なパラメーターを無視する可能性があると記載されています。
両方の障害モードが発生しました。 1 つの Qwen エンドポイントはサポートをアドバタイズしましたが、空ではないバイアス マップを含むすべてのリクエストを拒否しました。 GPT-OSS エンドポイントはパラメータを受け入れましたが、-100 禁止を無視しました。 DeepSeek 実験では、ルートを信頼する前に Cloudflare プロバイダーを固定し、フォールバックを無効にし、 require_parameters を有効にして、ハード禁止プローブを実行しました。 OpenRouter は現在、利用可能なモデルとして DeepSeek V4 Flash をリストしていますが、この実験の作業単位はモデルとその特定のプロバイダー ルートでした。
8つの記事はどうなったのか
ハード禁止プローブでは -100 が使用されました。監視されたすべてのワードが消え、プロバイダーがバイアスを適用したことが確認されました。その後、モデルは出力制限を使い果たすまで同じテキストを繰り返しました。パラメータは機能しましたが、結果は使用できませんでした。
8 つの記事の比較には -8 を使用しました。ソース記事には、私の監視リストに含まれる 264 件の記事が含まれていました。ブラックリストアクティブな書き換えは 180 個を保持し、264 個のうち 84 個を削除しました。
ブラックリストにより、モデルはさらに多くのテキストを変更しました。平均 5 グラムの重複。ソース記事とその編集版の間で共有される 5 つの単語シーケンスを測定します。

コントロール書き換えでは 0.982、ブラックリストアクティブ書き換えでは 0.896 でした。
ブラックリストアクティブな 8 件の書き換えのうち 7 件では、ソースの事実、エンティティ、および確実性のレベルが保存されていました。 8 番目はエンティティを変更し、「出現」という垣根を取り除き、限定された観察を直接的な主張に変えました。
別のレビューでは、文法、明瞭さ、繰り返し、文章の構成を比較しました。ブラックリストに基づくリライトのうち 3 件はソース記事よりも優れており、4 件はそれほど悪くなく、1 件はさらに悪かった。最悪の記事では、「その詳細が重要なのは…」が文法的ではない「その詳細がなぜ…」に変更され、他にもいくつかの厄介な置換が導入されていました。
そして、意味を変えたり散文を悪化させたりする記事はすべて拒否しました。 8 つの制御書き換えのうち 7 つが成功しました。ブラックリストアクティブな 8 件の書き換えのうち 6 件が成功しました。このバイアスにより、監視された出来事のおよそ 3 分の 1 が削除され、許容可能な記事 1 つが犠牲になりました。
モデルが時々知性が低く見えることがある理由
logit_bias はモデルの知識を減らしたり、重みを変更したりしませんでした。これにより、モデルの優先される次のトークンの一部を選択することが困難になりました。
モデルは依然として同じ事実を知っていました。しかし、ブラックリストには、どのトークンを避けるべきかだけが記載されていました。どの置き換えが文法的であるか、ヘッジが必要かどうか、固有名詞をそのままにしておく必要があるかどうかなどについての情報は提供されていませんでした。したがって、結果として得られる文は、能力の低いモデルからのものであるかのように見える可能性があります。
ペナルティを増やしても問題は解決しませんでした。別の一連のテストでは、-4 から -12 に移行しても、より多くのウォッチワードが一貫して削除されませんでした。 -100 では、モデルは単語を削除し、繰り返しループに入りました。バイアス値からは、次の記事がその主張と文法を維持するかどうかは分かりませんでした。
したがって、プロンプトブラックリストとロジットバイアス

さまざまな方法で失敗します。プロンプト バージョンは柔軟ですが、無視したり文字通りに従うこともできます。 API バージョンはトークンの選択に直接圧力を加えますが、ブラックリストに登録された単語がいつ正しい単語であるかを認識できません。
Vroni に GitHub の問題、バグレポート、仕様、または大まかなアイデアを提供します。リポジトリを読み取り、変更を計画し、コードを作成し、チェックを実行し、レビュー可能なプル リクエストに向けて作業します。
新しい投稿を公開すると取得します。
私はあなたのプライバシーを尊重します。いつでも購読を解除してください。
このトピックに関してさらに執筆します。
研究エージェントが検索クエリを通じて個人ファイルを漏洩する可能性がある
研究機関は、個人文書をアップロードせずに漏洩する可能性があります。詳細を読み取り、Web 検索に変換することができます。
Qualys は、Claude Mythos プレビューが Linux XFS での CVE-2026-64600 の発見に役立ったと言っています
2026 年 7 月 22 日、Qualys は、Linux ファイル システムである XFS の競合状態である CVE-2026-64600 を公開しました。欠陥は XFS のコピーオンライト パスにあります。
Google のマネージド エージェントがランタイムを Google のクラウドに配置します
現在パブリック プレビュー段階にある Google のマネージド エージェントは、Gemini エージェントに Google がホストする Linux サンドボックスを提供できます。 GoogleはデフォルトのAntigravityエージェントを文書化しています…
あなたのメールアドレスは公開されません。 * が付いているフィールドは必須です
ソフトウェア、AI、Laravel、SEO、オンライン製品の構築に必要なものについて書いています。 2025 年には、73,000 人がこのブログを読んでいます。一部の投稿は出版物や開発者コミュニティによって取り上げられました。
いくつかの古い投稿が最も注目を集めました。以下の最新の文章とは別に、ここに保管しておきます。
Google Now はデフォルトでコンテンツのインデックスを作成しません
この投稿は、Guardian のコラムと開発者コミュニティでの長い議論につながりました。
AI がテクノロジー負債を指数関数化する
AI によって脆弱なエンジニアリング基盤のコストが下がるのではなく、より高価になる理由。
専門家だけが適切なプロンプトを作成できる
AI の優れた機能は依然として技術に依存します

判断力と味覚。
実際のコードベースで作業する人向けの Laravel に関する具体的なアドバイス。
ブログの最新の投稿。これは公開すると自動的に更新されます。
logit_bias を通じて AI の言葉を禁止することで、AI の文章を人間らしいものにしようとした
私は、確率を変更する API 設定である logit_bias を使用して、AI の文章をより人間的に聞こえるようにしようとしました...
研究エージェントが検索クエリを通じて個人ファイルを漏洩する可能性がある
研究機関は、個人文書をアップロードせずに漏洩する可能性があります。詳細を読み取ったり、回転させたりすることができます...
Qualys は、Claude Mythos プレビューが Linux XFS での CVE-2026-64600 の発見に役立ったと言っています
2026 年 7 月 22 日、Qualys は、Linux ファイル システムである XFS の競合状態である CVE-2026-64600 を公開しました。欠陥は...
私の文章や解説について言及、引用、または議論した場所。
私がこれまでに仕事をしたクライアントからの実際のフィードバック。
「ヴィンセントを雇ってください。真剣に。私は何年にもわたって多くの開発者と仕事をしてきましたが（レンタコーダーを覚えていますか？）、ヴィンセントはそれを理解する稀な種類の Web 開発者です。チャットをし、必要なことを表現すると、彼は求めたものを返してきます。彼の英語は流暢で、あなたが何を必要としているのかを理解しており、物事を過度に複雑にするつもりはありません。彼と一緒に仕事をするのは喜びであり、細部まで管理する必要がなく、喜びです。全体的に素晴らしい経験でした

[切り捨てられた]

## Original Extract

I tried to make AI writing sound more human with , an API setting that changes how likely a model is to select specific tokens. I built a blacklist from words…

I Tried to Make AI Writing Sound Human by Banning AI Words Through logit_bias - Vincent Schmalbach
Skip to content
Software Developer
Home
Services
Custom Laravel App Development
Laravel Performance Optimization
Technical Team Leadership & Mentoring
Laravel Development Consulting
I Tried to Make AI Writing Sound Human by Banning AI Words Through logit_bias
July 29, 2026
by Vincent Schmalbach
I tried to make AI writing sound more human with logit_bias , an API setting that changes how likely a model is to select specific tokens. I built a blacklist from words that appeared unusually often in AI-generated text, converted the words into token IDs, and assigned those IDs negative values.
I'm writing this while working on vroni.com , my AI coding tool for turning tasks into pull requests.
This is different from putting “do not use these words” in the prompt. A prompt gives the model an instruction. The model reads it along with the article, decides how to apply it, and may ignore it when another instruction seems more important. logit_bias does not ask. It changes the numerical scores used to select every next token.
My idea was simple: if AI articles are easy to recognize partly because the same words keep appearing, perhaps I could reduce those words at the API level. I wanted the model to rewrite an AI-generated article while keeping its claims intact, but with fewer of the words associated with AI prose.
To test that, I started with eight complete AI-generated articles. I sent each full article to DeepSeek V4 Flash through OpenRouter and asked it to edit the prose while preserving every fact, name, number, example, qualification, recommendation, and conclusion. I did not process the articles paragraph by paragraph.
Each source article went through the model twice. Both calls used the same editing prompt and sampling settings. The control call used no token bias. The experimental call added a -8 bias to the token IDs on my blacklist. This produced eight control rewrites and eight blacklist-active rewrites of the same eight source articles.
Before a language model writes its next token, it assigns a score to every token it could choose. That score is called a logit. The API turns those scores into probabilities and samples the next token.
logit_bias inserts another number between those two steps. OpenAI's Chat Completions reference defines it as a map of token IDs to values from -100 to 100 . The API adds the supplied value to the model's score before sampling. A negative value reduces the chance that a token will be selected. A value near -100 can effectively ban it.
Suppose the model is about to write “This is a crucial distinction.” If the token for “ crucial” receives a negative bias, another continuation becomes more likely. The model might choose “important,” remove the adjective, or construct a different sentence.
That does not happen only once. Whatever token gets selected becomes part of the context used to score the following token. One forced substitution can therefore change the rest of the sentence and later sentences too.
Words also do not map neatly to tokens. “Crucial” at the start of a sentence may have a different token ID from “ crucial” after a space. A capitalized form may be different again. Some words split into several tokens. Penalizing a shared fragment can affect words that were never meant to be on the blacklist.
A prompt blacklist is not the same thing
Listing words in the prompt is easier because it does not require tokenization. It also gives the model some discretion. If “appears” is on the list but the source makes an uncertain claim, the model can keep the hedge because preserving the claim is more important than following the blacklist perfectly.
The disadvantage is that prompt instructions are not enforcement. The model may use a listed word anyway, especially in a long article with many other requirements. It may also obey mechanically and replace a normal word with an awkward synonym.
I tested this separately with GPT-5.6 Sol, which does not accept logit_bias . Five control articles used an ordinary editing prompt. Five experimental calls used the same task plus an explicit 93-word suppression list.
The prompt list reduced watched-word occurrences from 11 across the five control outputs to one across the five experimental outputs. But two of the five experimental articles dropped or changed information. All five controls preserved the source meaning.
That was not a direct comparison with the DeepSeek experiment because the model and source set were different. It did show that merely moving the blacklist into the prompt does not remove the tradeoff. The model can follow the list and still damage the article.
Which APIs and models support it
There is no single model list that can be trusted without testing the actual endpoint.
OpenAI still includes logit_bias in the Chat Completions schema, but that does not mean current OpenAI models accept it. I sent compatibility probes to GPT-5.2 Chat, GPT-5.3 Chat, GPT-5.4, GPT-5.4 Mini, GPT-5.4 Nano, GPT-5.5, GPT-5.6 Sol, GPT-5.6 Terra, and GPT-5.6 Luna. Every one returned an unsupported_parameter error. The current Responses API does not expose the field either.
Anthropic's Messages API has no native logit_bias parameter. A Claude prompt can contain a word list, but the API does not provide the same token-ID score adjustment.
OpenRouter has much broader support because it routes to many open and hosted models. Its Models API can filter models by supported_parameters . On July 29, 2026, filtering for logit_bias returned 115 model IDs. The list included DeepSeek V4 Flash, GLM 5.2, Kimi K3, Gemma 4 31B, Qwen 3.6 27B, Llama 3.3 70B, Mistral Small 3.2 24B, GPT-4o, and several dated GPT-4o versions.
Even that list is not proof that a particular request will work. OpenRouter combines several providers for some models. Its routing documentation says a provider may receive a request and ignore an unknown parameter unless require_parameters is enabled.
I saw both failure modes. One Qwen endpoint advertised support but rejected every request containing a non-empty bias map. A GPT-OSS endpoint accepted the parameter but ignored a -100 ban. For the DeepSeek experiment, I pinned the Cloudflare provider, disabled fallbacks, enabled require_parameters , and ran a hard-ban probe before trusting the route. OpenRouter currently lists DeepSeek V4 Flash as an available model, but the working unit in this experiment was the model plus that specific provider route.
What happened to the eight articles
The hard-ban probe used -100 . Every watched word disappeared, confirming that the provider applied the bias. The model then repeated the same text until it exhausted the output limit. The parameter worked, but the result was unusable.
I used -8 for the eight-article comparison. The source articles contained 264 occurrences from my watched list. The blacklist-active rewrites retained 180, removing 84 of 264.
The blacklist also caused the model to change more text. Average 5-gram overlap, which measures shared five-word sequences between a source article and its edited version, was 0.982 in the control rewrites and 0.896 in the blacklist-active rewrites.
Seven of the eight blacklist-active rewrites preserved the source facts, entities, and levels of certainty. The eighth changed an entity and removed the hedge “appears,” turning a qualified observation into a direct assertion.
A separate review compared grammar, clarity, repetition, and sentence construction. Three blacklist-active rewrites were better than their source articles, four were no worse, and one was worse. The worse article changed “That detail matters because…” into the ungrammatical “That details why…” and introduced several other awkward substitutions.
I then rejected any article that either changed the meaning or made the prose worse. Seven of the eight control rewrites passed. Six of the eight blacklist-active rewrites passed. The bias removed roughly a third of the watched occurrences and cost one acceptable article.
Why the model sometimes looked less intelligent
logit_bias did not reduce the model's knowledge or change its weights. It made some of the model's preferred next tokens harder to select.
The model still knew the same facts. But the blacklist only said which tokens to avoid. It supplied no information about which replacement was grammatical, whether a hedge was necessary, or whether a proper name had to remain untouched. The resulting sentences could therefore look as if they came from a less capable model.
Increasing the penalty did not solve this. In another set of tests, moving from -4 to -12 did not consistently remove more watched words. At -100 , the model removed the words and entered a repetition loop. The bias value did not tell me whether the next article would preserve its claims and grammar.
A prompt blacklist and logit_bias therefore fail in different ways. The prompt version is flexible but can be ignored or followed too literally. The API version applies direct pressure to token selection, but it cannot recognize when a blacklisted word is the correct word.
Give Vroni a GitHub issue, bug report, spec, or rough idea. It reads the repo, plans the change, writes code, runs checks, and works toward a review-ready pull request.
Get new posts when I publish them.
I respect your privacy. Unsubscribe at any time.
More writing around this topic.
A Research Agent Can Leak Private Files Through Its Search Queries
A research agent can leak a private document without uploading it. It can read a detail, turn it into a web search,…
Qualys Says Claude Mythos Preview Helped Find CVE-2026-64600 in Linux XFS
On July 22, 2026, Qualys disclosed CVE-2026-64600, a race condition in XFS, a Linux filesystem. The flaw is in XFS's copy-on-write path,…
Google’s Managed Agents Put Your Runtime in Google’s Cloud
Google’s Managed Agents, now in Public Preview, can give a Gemini agent a Google-hosted Linux sandbox. Google documents the default Antigravity agent…
Your email address will not be published. Required fields are marked *
I write about software, AI, Laravel, SEO, and what it takes to build online products. In 2025, 73k people read the blog. Some posts were picked up by publications and developer communities.
A few older posts got the most attention. I keep them here, separate from the newest writing below.
Google Now Defaults to Not Indexing Your Content
The post that led to a Guardian column and a long developer-community discussion.
AI Exponentializes Your Tech Debt
Why AI makes weak engineering foundations more expensive, not less.
Only Experts Can Write Good Prompts
Good AI work still depends on technical judgment and taste.
Specific Laravel advice for people working in real codebases.
The newest posts from the blog. This updates automatically when I publish.
I Tried to Make AI Writing Sound Human by Banning AI Words Through logit_bias
I tried to make AI writing sound more human with logit_bias, an API setting that changes how likely...
A Research Agent Can Leak Private Files Through Its Search Queries
A research agent can leak a private document without uploading it. It can read a detail, turn it...
Qualys Says Claude Mythos Preview Helped Find CVE-2026-64600 in Linux XFS
On July 22, 2026, Qualys disclosed CVE-2026-64600, a race condition in XFS, a Linux filesystem. The flaw is...
Places that have mentioned, quoted, or discussed my writing and commentary.
Actual feedback from clients I've worked with.
"HIRE VINCENT. Seriously. I've worked with many many developers over the years (remember rentacoder?), and Vincent is that rare breed of web developer who just gets it. You have a chat, you express what you need, and he comes back with what you asked for. His english is fluent, he understands what you need, and he's not looking to overcomplicate things. It's a joy and pleasure to work with him, not to have to micromanage and just overall a great expe

[truncated]
