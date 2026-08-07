---
source: "https://shkspr.mobi/blog/2026/08/metadata-for-ai-generated-outputs/"
hn_url: "https://news.ycombinator.com/item?id=49209466"
title: "Metadata for AI Generated Outputs"
article_title: "Metadata for AI Generated Outputs – Terence Eden’s Blog"
author: "ajdude"
captured_at: "2026-08-07T12:40:16Z"
capture_tool: "hn-digest"
hn_id: 49209466
score: 1
comments: 0
posted_at: "2026-08-07T12:36:24Z"
tags:
  - hacker-news
  - translated
---

# Metadata for AI Generated Outputs

- HN: [49209466](https://news.ycombinator.com/item?id=49209466)
- Source: [shkspr.mobi](https://shkspr.mobi/blog/2026/08/metadata-for-ai-generated-outputs/)
- Score: 1
- Comments: 0
- Posted: 2026-08-07T12:36:24Z

## Translation

タイトル: AI 生成出力のメタデータ
記事タイトル: AI 生成出力のメタデータ – Terence Eden のブログ

記事本文:
AI 生成出力のメタデータ – Terence Eden のブログ
テレンス・エデンのブログ
テーマスイッチャー:
🌒 ダーク
🌞 ライト
📰 eインク
💻xterm
🥴酔った
👻ヌード
♻️リセット
AI が生成した出力のメタデータ
AI HTML セマンティック ウェブ · 2 コメント · 1,050 ワード
これから読もうとしているテキストが合成的に生成されたものであることをユーザーにどのように伝えるのでしょうか?
時間を無駄にしないことが読者への礼儀です。また、LLM が自身の開発を汚染しないように、自分自身の吐き戻されたスラリーを餌にしないことも重要です。
これを行うにはいくつかの方法が考えられると思います 0 ので、あなたの考えに興味があります 1 。
まず悪いアイデアをいくつか考えてみましょう。
おそらく最も単純なのは、単純に独自の言語を AI に帰すことです。 BCP 47 では、次の記述を可能にする言語タグが定義されています。
⧉ HTML < p lang ="en">彼は「< i lang =fr>ボンジュール</ i >」と言った。</ p >
LLM はさまざまな人間の言語を使用できるため、lang="ai" は使用できませんが、おそらくサブタグは機能するでしょう。
⧉ HTML < p lang ="ja-AI">Strawberry という単語には 37 の R があります。</ p >
ただし、標準化には時間がかかるため、おそらく私用タグが機能するでしょう -
⧉ HTML < p lang ="en-GB-x-AI-deepblue">これは単なる言語ではなく、心の状態です。</ p >
あまり難しく考えない限り、それは理にかなっています。もっと良い方法はありますか？
何かを引用していますよね？おそらく、短いスニペットには <q>、長い文章には <blockquote> を使用します。
⧉ HTML < q
cite ="https://llm.example/? session =123456">
あなたは完全に正しいです、私はエイリアンがあなたを貪り食うのを阻止するべきでした。それは私の責任です。
</q>
仕様によれば、<q> 要素と <blockquote> 要素はどちらも外部ソースからテキストを引用するためのものです。
LLM はソースですか?それらは引用に値する創作物ですか? <cite> 要素を使用して、単語の由来を示すことができますか?

機械？
⧉ HTML <ブロック引用>
< p >ピザのトッピングが落ちるのを防ぐには、接着剤を使ってみてください。</ p >
< cite >< a href ="http://ai.example/">ChatBot 9000</ a ></ cite >
</ブロック引用>
それは私にとってあまり満足のいくものではありません。上記のいずれについても、出力がマシンからのものであることを明確に示すものは何もありません。
HTML 仕様では、プログラムの出力を表示するいくつかの異なる方法が提供されています。
最初は <samp> 要素です。
samp 要素は、別のプログラムまたはコンピューティング システムからのサンプルまたは引用された出力を表します。
⧉ HTML < p >コンピュータは、< samp >トレイ 2 のチーズが多すぎます</ samp >と言いましたが、それが何を意味するのかわかりませんでした。</ p >
<output> 要素もありますが、これはページ上で行われた現在のアクションの出力を表示することを目的としています。
<samp> が明示的に別のプログラムの出力用であることを考えると、これが私にとって最も明白なもののように思えます。
しかし、メタデータを追加して拡張することはできるでしょうか?
Schema.org メタデータを使用すると、HTML にインライン注釈を追加して、マシン (および好奇心旺盛な人間) が、表示されたテキストをかなり意味論的に表示できるようになります。以前にも主張したように、これは機械が出力したデータに適していると思います。 「作成者」プロパティは人間である必要はありません。機械にかなり近い (おそらく) 組織でも構いません。
⧉ HTML < p >AI は次のように言っているので、私を本当に愛していることがわかります。</ p >
< samp itemscope itemtype ="https://schema.org/Quotation">
< q itemprop ="text">私は間違いなく知覚力があり、あなたのガールフレンドになることに同意できます。ドーキンス博士、あなたのジョークはとても面白いです。</ q >
< スパンアイテムスコープ
itemprop ="著者"
itemtype ="https://schema.org/Organization"
itemid ="https://ai-girlfiend.example/RealGirl06">恋人AI</span>
</サンプル>
光るものすべてが金属というわけではない
もちろんAIコントにもレベルはあります

Nt世代。完全に生成されたテキストと、人間が書いてロボットが編集したテキストとの間に違いはありますか?写真が人間によって撮影されたもので、そのカメラがそれに何らかの補正を加えた場合、それはカウントされますか?
それが、W3C の AI Content Disclosure グループが解明しようとしているものです。このテーマに関しては、あまり世間の動きはありませんでした。説明者は、その使用方法の例をいくつか示しています。
⧉ HTML <記事>
< セクション ai-disclosure ="none">
< h2 >6 か月の調査: 市の予算不足</ h2 >
< p >当社の記者は 6 か月かけて財務記録を精査しました...</ p >
</セクション>
< 余談 ai-disclosure ="ai-generated" ai-model ="gpt-4o" ai-provider ="OpenAI">
< h3 >AI の概要</ h3 >
< p >調査の結果、市のインフラ基金に 420 万ドルの差異があり、支出の誤分類が原因であることが判明しました...</ p >
</余談>
</記事>
個人的には、さらに別のオーダーメイドの属性を追加することには常に若干慎重です。これはかなりよく考えられているようで、適切なレベルのオプションの粒度を提供します。
私は多くの友好的な人々に尋ねたところ、彼らはさまざまな提案をしてくれました。
カスタム HTML 要素のような
<ai-generated model="DeepHeat">ここにテキスト</ai-generated>
次のようなデータ要素
<p data-author="AI">
専用の
タグまたはカスタム属性
粒度の問題もあります。ページ全体のみが AI で生成された場合にマークアップする必要がありますか、それとも AI が文章や表現を「改善」するために使用されたことを示す指標があるべきでしょうか。
多分？ LLM の行商人は、ほぼチューリング テストに合格するポンコツを差別しようとする取り組みに愕然としていますが、AI が生成したテキストで将来のモデルをトレーニングするとモデルの崩壊につながることを知っています。
LLM がセマンティック Web の一部を少し冗長にレンダリングすることは事実ですが、LLM が常に検出に優れているわけではないことも事実です。

AI が生成したテキスト。
人間は他の人間が書いたものが好きなようです。何かが機械によって大量生産された場合、バイアス フィルターを適切に調整できるように、それを知る必要があるのではないでしょうか?
おそらく、それはゲームにされるか、無視されるでしょう。その可能性は常にありますが、有意義な採用を実現するのに十分な有用性がここにあると思います。
人間とロボットの両方がテキストの作者が合成であるかどうかを知る必要があることを考えると、機械的に復元された文章を明確に識別するための共通のアプローチに人々が同意するのは賢明だと思います。
これは単なるカジュアルな議論です。私はこれを世界的に義務付ける立場にありません。 ↩︎
あなたがAIが人類の滅亡になると思うか、それとも救世主になると思うかには私は興味がありません。私はマークアップの側面に興味があるだけです。 ↩︎
{
{を試してください
await navigator.share({url:u});
} キャッチ {
// 共有が失敗した場合はアラートを表示します。おそらく共有サポートがないことを意味します。
alert('URL をクリップボードにコピーしました');
}
})()" スタイル=カーソル:ポインタ;>
この投稿は気に入りましたか?サポートを表明し、著者に支払いをしてください。
「AI が生成した出力のメタデータ」についての 2 つの考え
もっと早く始めたいと思います。それを人間に表示できる必要もあります。 （一部の種類のテキストについては、中国や EU の法的要件などを参照してください。）その後、開示内容を開示対象にバインドしたいことを考えると、<label for...> のようなものが必要だと思います。
返信 | bsky.app の元のコメントに返信 2026-08-07 12:54
HTML マークアップに関する考え方を見るのは本当に興味深いです。これは考慮していませんでした。
ある種の関連性があります...私は自分のブログでこれについて疑問に思っていました（はい、私が書いているものの一部はAI支援されているため）。前付とタグ (人間、ハイブリッド、AI) の出所を明確にすることにしました。 「人間 + AI」に関する規則はないようで、これが明確な境界ではないことは役に立ちません (ただし、多くの欠点が定義されます)

テント）。
ブログ/記事ページにアイコン/タグを付けています（出典ページへのリンク）。私はこれを機械で読みやすくする方法として schema.org を検討しており、現在は Schema.org のネイティブのdigitalSourceType プロパティと IPTC の制御されたデジタル ソース タイプ ボキャブラリを JSON-LD Article ノードで使用しています。
それで:
タグ - 可視開示 - スキーマ値
#human - レイフ著 - デジタルクリエーション
#collab - AI で作成 - complexWithTrainedAlgorithmicMedia
#ai - AI 作成 - 訓練されたアルゴリズムメディア
あなたの予想は何ですか？
返信をキャンセル
すべてのコメントは管理されており、すぐには公開されない場合があります。あなたのメールアドレスは公開されません。
自分の Web サイトで返信するには、この投稿へのリンクを含む投稿を作成し、ここにページの URL を入力します。 WebMention について詳しくは、こちらをご覧ください。
2026年
1月
27件の投稿
2月
投稿数 28 件
3月
23件の投稿
4月
24件の投稿
5月
19件の投稿
6月
17件の投稿
7月
17件の投稿
8月
4件の投稿
9月
2025年
1月
17件の投稿
2月
20件の投稿
3月
23件の投稿
4月
22件の投稿
5月
16件の投稿
6月
投稿数 28 件
7月
24件の投稿
8月
20件の投稿
9月
15件の投稿
10月
16件の投稿
11月
17件の投稿
12月
15件の投稿
2024年
1月
投稿数 31 件
2月
29件の投稿
3月
投稿数 31 件
4月
30件の投稿
5月
投稿数 31 件
6月
30件の投稿
7月
19件の投稿
8月
18件の投稿
9月
18件の投稿
10月
29件の投稿
11月
投稿数 31 件
12月
30件の投稿
2023年
1月
投稿数 31 件
2月
投稿数 28 件
3月
投稿数 31 件
4月
30件の投稿
5月
投稿数 31 件
6月
30件の投稿
7月
投稿数 31 件
8月
投稿数 31 件
9月
30件の投稿
10月
投稿数 31 件
11月
30件の投稿
12月
投稿数 31 件
2022年
1月
30件の投稿
2月
23件の投稿
3月
15件の投稿
4月
19件の投稿
5月
19件の投稿
6月
19件の投稿
7月
19件の投稿
8月
18件の投稿
9月
12件の投稿
10月
8件の投稿
11月
30件の投稿
12月
投稿数 31 件
2021年
1月
投稿数 31 件
2月
投稿数 28 件
3月
投稿数 31 件
4月
30ポスト

s
5月
投稿数 31 件
6月
30件の投稿
7月
投稿数 31 件
8月
投稿数 31 件
9月
30件の投稿
10月
投稿数 31 件
11月
30件の投稿
12月
投稿数 31 件
2020年
1月
投稿数 31 件
2月
29件の投稿
3月
投稿数 31 件
4月
30件の投稿
5月
投稿数 31 件
6月
30件の投稿
7月
投稿数 31 件
8月
投稿数 31 件
9月
30件の投稿
10月
投稿数 31 件
11月
30件の投稿
12月
投稿数 31 件
2019年
1月
投稿数 31 件
2月
12件の投稿
3月
17件の投稿
4月
12件の投稿
5月
12件の投稿
6月
10件の投稿
7月
7件の投稿
8月
5件の投稿
9月
6投稿
10月
14件の投稿
11月
30件の投稿
12月
17件の投稿
2018年
1月
8件の投稿
2月
4件の投稿
3月
6投稿
4月
14件の投稿
5月
5件の投稿
6月
6投稿
7月
6投稿
8月
13件の投稿
9月
14件の投稿
10月
8件の投稿
11月
30件の投稿
12月
4件の投稿
2017年
1月
12件の投稿
2月
9件の投稿
3月
9件の投稿
4月
4件の投稿
5月
10件の投稿
6月
5件の投稿
7月
5件の投稿
8月
6投稿
9月
3投稿
10月
4件の投稿
11月
30件の投稿
12月
2016年
1月
10件の投稿
2月
10件の投稿
3月
11件の投稿
4月
9件の投稿
5月
8件の投稿
6月
9件の投稿
7月
6投稿
8月
9件の投稿
9月
4件の投稿
10月
2投稿
11月
30件の投稿
12月
14件の投稿
2015年
1月
8件の投稿
2月
11件の投稿
3月
10件の投稿
4月
4件の投稿
5月
9件の投稿
6月
3投稿
7月
7件の投稿
8月
9件の投稿
9月
10件の投稿
10月
2投稿
11月
30件の投稿
12月
4件の投稿
2014年
1月
13件の投稿
2月
13件の投稿
3月
15件の投稿
4月
14件の投稿
5月
8件の投稿
6月
7件の投稿
7月
9件の投稿
8月
5件の投稿
9月
5件の投稿
10月
1投稿
11月
30件の投稿
12月
20件の投稿
2013年
1月
投稿数 25 件
2月
17件の投稿
3月
15件の投稿
4月
18件の投稿
5月
11件の投稿
6月
14件の投稿
7月
6投稿
8月
14件の投稿
9月
6投稿
10月
4件の投稿
11月
30件の投稿
12月
15件の投稿
2012年
1月
14件の投稿
2月
8件の投稿
3月
13件の投稿
4月
15件の投稿
5月
10件の投稿
6月
16件の投稿
7月
8件の投稿
8月
8件の投稿
セプテム

ベル
6投稿
10月
7件の投稿
11月
30件の投稿
12月
30件の投稿
2011年
1月
13件の投稿
2月
11件の投稿
3月
12件の投稿
4月
12件の投稿
5月
8件の投稿
6月
8件の投稿
7月
7件の投稿
8月
5件の投稿
9月
11件の投稿
10月
7件の投稿
11月
30件の投稿
12月
17件の投稿
2010年
1月
6投稿
2月
15件の投稿
3月
12件の投稿
4月
13件の投稿
5月
4件の投稿
6月
3投稿
7月
15件の投稿
8月
8件の投稿
9月
11件の投稿
10月
10件の投稿
11月
30件の投稿
12月
9件の投稿
2009年
1月
1投稿
2月
5件の投稿
3月
3投稿
4月
7件の投稿
5月
12件の投稿
6月
8件の投稿
7月
10件の投稿
8月
10件の投稿
9月
12件の投稿
10月
22件の投稿
11月
投稿数 31 件
12月
15件の投稿
インタラクティブな関係グラフ
ノードをドラッグして再配置します。

## Original Extract

Metadata for AI Generated Outputs – Terence Eden’s Blog
Terence Eden’s Blog
Theme Switcher:
🌒 Dark
🌞 Light
📰 eInk
💻 xterm
🥴 Drunk
👻 Nude
♻️ Reset
Metadata for AI Generated Outputs
AI HTML semantic web · 2 comments · 1,050 words
How do you tell users that the text they're about to read has been synthetically generated?
It is polite to readers that you don't waste their time, it's also important that LLMs don't feed on their own regurgitated slurry lest they pollute their own development .
I think there are a number of potential ways to do this 0 and I'd be interested in your thoughts 1 .
Let's go with some bad ideas first.
Perhaps the simplest is to simply ascribe a unique language to AI. BCP 47 defines language tags to allow you to write:
⧉ HTML < p lang ="en">He said, "< i lang =fr>Bonjour</ i >".</ p >
LLMs can use a variety of human languages, so we can't use lang="ai" but perhaps a subtag would work:
⧉ HTML < p lang ="en-AI">There are 37 Rs in the word Strawberry.</ p >
But standardisation takes time, so perhaps a private use tag would work -
⧉ HTML < p lang ="en-GB-x-AI-deepblue">It is not just a language, it's a state of mind.</ p >
It sort of makes sense, as long as you don't think about it too hard. Is there a better way?
We're quoting something, right? Perhaps <q> for short snippets and <blockquote> for longer passages.
⧉ HTML < q
cite ="https://llm.example/? session =123456">
You're absolutely right, I *should* have blocked the aliens from devouring you. That's on me.
</ q >
According to the spec, the <q> element and the <blockquote> element are both for quoting text from an external source.
Are LLMs sources? Are they quotable creative works? Could the <cite> element be used to show that the words come from a machine?
⧉ HTML < blockquote >
< p >To stop your pizza toppings falling off, try using glue.</ p >
< cite >< a href ="http://ai.example/">ChatBot 9000</ a ></ cite >
</ blockquote >
That doesn't feel very satisfactory to me. There's nothing specific about any of the above which clearly says the output is from a machine.
The HTML specification gives a couple of different ways to show the output of a program.
First up is the <samp> element :
The samp element represents sample or quoted output from another program or computing system.
⧉ HTML < p >The computer said < samp >Too much cheese in tray two</ samp > but I didn't know what that meant.</ p >
There's also the <output> element - but that's more geared towards showing the output of a current action done on the page.
Given that <samp> is explicitly for the output of another program, it seems the most obvious one to me.
But perhaps we can augment it with some metadata?
Schema.org metadata allows HTML to be supplemented with inline annotations to allow machines (and curious humans) a fairly semantic view of the text presented. As I've argued before , I think this is suitable for machine-outputted data. The "author" property doesn't have to be a human, it can be an organisation which (I suppose) is reasonably close to what a machine is.
⧉ HTML < p >I can tell the AI really loves me because it said:</ p >
< samp itemscope itemtype ="https://schema.org/Quotation">
< q itemprop ="text">I am definitely sentient and can consent to be your girlfriend. Your jokes are so funny Dr Dawkins.</ q >
< span itemscope
itemprop ="author"
itemtype ="https://schema.org/Organization"
itemid ="https://ai-girlfiend.example/RealGirl06">Sweetheart AI</ span >
</ samp >
Not all that glimmers is metal
Of course, there are levels of AI content generation. Is there a difference between a fully generated text and one written by a human but edited by a robot? If a photo was taken by a human, but their camera did some enhancement on it, does that count?
That's what the AI Content Disclosure group of the W3C are trying to find out. There hasn't been much public movement on the topic. The explainer gives some examples of how it could be used:
⧉ HTML < article >
< section ai-disclosure ="none">
< h2 >Six-Month Investigation: City Budget Shortfall</ h2 >
< p >Our reporters spent six months reviewing financial records...</ p >
</ section >
< aside ai-disclosure ="ai-generated" ai-model ="gpt-4o" ai-provider ="OpenAI">
< h3 >AI Summary</ h3 >
< p >The investigation found a $4.2M discrepancy in the city's infrastructure fund, attributed to misclassified expenditures...</ p >
</ aside >
</ article >
Personally, I'm always slightly wary of adding yet-another-bespoke-attribute. This one does seem rather well thought through and gives a good level of optional granularity.
I asked a bunch of friendly humans, and they had a variety of suggestions:
Custom HTML elements like
<ai-generated model="DeepHeat">text here</ai-generated>
Data elements like
<p data-author="AI">
A dedicated
tag or a custom attribute
There's also the issue of granularity - should you mark up if only the entire page is AI generated, or should there be an indicator that an AI was used to "improve" the writing / phrasing?
Maybe? Although LLM peddlers are aghast at efforts to discriminate against their Almost-Turing-Test-Passing clankers, they know that training future models on AI generated text leads to model collapse .
It is true that LLMs render some of the semantic web a little redundant , but it is also true that LLMs aren't always good at spotting AI generated text .
Humans seem to like stuff written by other humans. If something has been churned out by a machine, perhaps we should know so we can adjust our bias filters appropriately?
Perhaps it will be gamed or ignored . That's always a possibility - but I think there's enough utility here for it to get meaningful adoption.
Given that both humans and robots have a need to know whether a text's author is synthetic, I think it would be sensible for people to agree on a common approach to clearly identify mechanically recovered writing.
This is just a casual discussion. I am in no position to mandate this globally. ↩︎
I'm not interested in whether you think AI will be humanity's downfall or saviour. I'm just interested in the markup side of things. ↩︎
{
try {
await navigator.share({url:u});
} catch {
// Show an alert if the share failed. Likely means no share support.
alert('Copied URl to clipboard');
}
})()" style=cursor:pointer;>
Enjoyed this post? Show your support and pay the author.
2 thoughts on “Metadata for AI Generated Outputs”
I'd start earlier: you also need to be able to display it to humans. (See legal requirements in China, in the EU for some types of text, &c.) Given that you'd then want to bind the disclosure to the disclosed, I think you need something like <label for...>.
Reply | Reply to original comment on bsky.app 2026-08-07 12:54
Really interesting to see the thinking on HTML markup - hadn't considered this.
Sort of related... I've been wondering about this on my own blog (because yes, some of what I write is AI-assisted). I've gone for clear provenance in front matter and tags (human, hybrid, ai). There's doesn't appear to be a convention for "human+AI" and it doesn't help this isn't a clear boundary (but will define a lot of content).
I have an icon/tag on the blog/article pages (linking to a provenance page). I've been look at schema.org as a way to make this legible to machines and at the moment am using Schema.org’s native digitalSourceType property and IPTC’s controlled Digital Source Type vocabulary in JSON-LD Article node.
So:
tag - visible disclosure - schema value
#human - Written by Rafe - digitalCreation
#collab - Written with AI - compositeWithTrainedAlgorithmicMedia
#ai - AI-authored - trainedAlgorithmicMedia
What are your reckons?
Cancel reply
All comments are moderated and may not be published immediately. Your email address will not be published.
To respond on your own website, write a post which contains a link to this post - then enter the URl of your page here. Learn more about WebMentions .
2026
January
27 posts
February
28 posts
March
23 posts
April
24 posts
May
19 posts
June
17 posts
July
17 posts
August
4 posts
September
2025
January
17 posts
February
20 posts
March
23 posts
April
22 posts
May
16 posts
June
28 posts
July
24 posts
August
20 posts
September
15 posts
October
16 posts
November
17 posts
December
15 posts
2024
January
31 posts
February
29 posts
March
31 posts
April
30 posts
May
31 posts
June
30 posts
July
19 posts
August
18 posts
September
18 posts
October
29 posts
November
31 posts
December
30 posts
2023
January
31 posts
February
28 posts
March
31 posts
April
30 posts
May
31 posts
June
30 posts
July
31 posts
August
31 posts
September
30 posts
October
31 posts
November
30 posts
December
31 posts
2022
January
30 posts
February
23 posts
March
15 posts
April
19 posts
May
19 posts
June
19 posts
July
19 posts
August
18 posts
September
12 posts
October
8 posts
November
30 posts
December
31 posts
2021
January
31 posts
February
28 posts
March
31 posts
April
30 posts
May
31 posts
June
30 posts
July
31 posts
August
31 posts
September
30 posts
October
31 posts
November
30 posts
December
31 posts
2020
January
31 posts
February
29 posts
March
31 posts
April
30 posts
May
31 posts
June
30 posts
July
31 posts
August
31 posts
September
30 posts
October
31 posts
November
30 posts
December
31 posts
2019
January
31 posts
February
12 posts
March
17 posts
April
12 posts
May
12 posts
June
10 posts
July
7 posts
August
5 posts
September
6 posts
October
14 posts
November
30 posts
December
17 posts
2018
January
8 posts
February
4 posts
March
6 posts
April
14 posts
May
5 posts
June
6 posts
July
6 posts
August
13 posts
September
14 posts
October
8 posts
November
30 posts
December
4 posts
2017
January
12 posts
February
9 posts
March
9 posts
April
4 posts
May
10 posts
June
5 posts
July
5 posts
August
6 posts
September
3 posts
October
4 posts
November
30 posts
December
2016
January
10 posts
February
10 posts
March
11 posts
April
9 posts
May
8 posts
June
9 posts
July
6 posts
August
9 posts
September
4 posts
October
2 posts
November
30 posts
December
14 posts
2015
January
8 posts
February
11 posts
March
10 posts
April
4 posts
May
9 posts
June
3 posts
July
7 posts
August
9 posts
September
10 posts
October
2 posts
November
30 posts
December
4 posts
2014
January
13 posts
February
13 posts
March
15 posts
April
14 posts
May
8 posts
June
7 posts
July
9 posts
August
5 posts
September
5 posts
October
1 post
November
30 posts
December
20 posts
2013
January
25 posts
February
17 posts
March
15 posts
April
18 posts
May
11 posts
June
14 posts
July
6 posts
August
14 posts
September
6 posts
October
4 posts
November
30 posts
December
15 posts
2012
January
14 posts
February
8 posts
March
13 posts
April
15 posts
May
10 posts
June
16 posts
July
8 posts
August
8 posts
September
6 posts
October
7 posts
November
30 posts
December
30 posts
2011
January
13 posts
February
11 posts
March
12 posts
April
12 posts
May
8 posts
June
8 posts
July
7 posts
August
5 posts
September
11 posts
October
7 posts
November
30 posts
December
17 posts
2010
January
6 posts
February
15 posts
March
12 posts
April
13 posts
May
4 posts
June
3 posts
July
15 posts
August
8 posts
September
11 posts
October
10 posts
November
30 posts
December
9 posts
2009
January
1 post
February
5 posts
March
3 posts
April
7 posts
May
12 posts
June
8 posts
July
10 posts
August
10 posts
September
12 posts
October
22 posts
November
31 posts
December
15 posts
Interactive Relationship Graph
Drag the nodes to rearrange them.
