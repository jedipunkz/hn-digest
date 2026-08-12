---
source: "https://johnjwang.com/post/2026/08/12/how-claude-watermarking-probably-works/"
hn_url: "https://news.ycombinator.com/item?id=49276728"
title: "How Claude's watermarking (probably) works"
article_title: "How Claude's watermarking (probably) works | John Wang"
author: "johnjwang"
captured_at: "2026-08-12T18:47:56Z"
capture_tool: "hn-digest"
hn_id: 49276728
score: 1
comments: 0
posted_at: "2026-08-12T18:32:23Z"
tags:
  - hacker-news
  - translated
---

# How Claude's watermarking (probably) works

- HN: [49276728](https://news.ycombinator.com/item?id=49276728)
- Source: [johnjwang.com](https://johnjwang.com/post/2026/08/12/how-claude-watermarking-probably-works/)
- Score: 1
- Comments: 0
- Posted: 2026-08-12T18:32:23Z

## Translation

タイトル: クロードの透かし (おそらく) の仕組み
記事のタイトル: クロードのウォーターマーク (おそらく) の仕組み |ジョン・ワン

記事本文:
クロードの透かしの仕組み (おそらく)
昨日、Anthropic は、AI が生成したコンテンツに透かしを入れ始めたと発表しました。インターネット上の人々は特にこれについて憤慨しており（私もそれは当然だと思います）、特にこれは EU 内かどうかに関係なく、明らかにすべてのクロード モデルで起こっているためです。彼らが実際に何をしているのか、そしてそれが知覚できるのか、それとも変化し得るのかを調査したかったのです。
Anthropic は、記事「 How Claude mark AI-generated content 」で、彼らのアプローチについての洞察を提供しています。実際には実装に関する技術的な詳細は提供されていませんが、使用しているスキームの周囲に幅広い網を引くのに役立ついくつかのガイダンスを提供します。 Anthropic がヘルプセンターの記事に残した重要な手がかり:
このスキームは「目に見えない透かしをテキスト自体に直接織り込みます。透かしは表示されず、クロードの応答の意味、質、読みやすさは変わりません。」
透かしは非常に短いテキストではうまく機能しません。
ウォーターマークは最近（8月2日以降）に始まったようです
これは可能性をかなり絞り込むのに役立ちます。
Anthropic が何をしているのかを実際に理解するために、実行できる興味深い実験がいくつかあることに気付きました。グローゲンら。統計的な透かしをチェックするための特定のテストを作成しました。また、Claude 出力の分析を実行して、隠された Unicode や空白の結果などをチェックすることもできます。
念のため、私はクロード チャットのダンプもダウンロードし (2025 年 2 月 4 日からクロード コードを使用し、1206 セッションを記録しました)、Fable 5 が出力するトークンの組み合わせ (私の通常の毎日のドライバー モデル) を明らかに変更する変更点が 8 月初旬頃にあったかどうかを簡単に比較しました。見つかりませんでした

この歴史的分析には知覚できる差異はなく、より具体的なテストを使用しない限り、透かしは一般的に知覚できないという Anthropic の主張が裏付けられます。
隠された Unicode や空白はありません
次にチェックしたのは、隠された Unicode、空白、句読点パターンがないかどうかです。これはかなり決定的でした。彼らはそれほど単純なことをやっていないのです。
抽出された 720 万の散文文字 (過去のクロード コード テキストと 8 月 11 日に生成されたクロード コード テキストの両方) にわたる分析では、クロードによって出力された唯一の Unicode 文字が合理的であり、日常使用の一部であることが示されました。
波引用符、アポストロフィ、および水平方向の省略記号
全英ダッシュ (残念ながらたくさんあります)
アクセント付き文字または英語以外の文字
Codex による監査では、私が約 10 個のサンプルを抜き取りチェックしましたが、隠された Unicode ウォーターマークと一致する異常なインスタンスは見つかりませんでした。同様に、空白エンコードマークも見つかりませんでした。
この証拠は、透かしが知覚できず、短いテキストでは機能しないという事実と組み合わせると、Anthropic が何らかの形式の統計トークン透かしを使用している可能性が高いことを意味します。
統計的透かし入れスキーム
テキストに透かしを追加できるスキームがいくつかあります。最も単純なバージョンである、キルヒェンバウアーらによって 2023 年に作成されたグリーン/レッド リストについて説明します。なぜなら、それが説明するのが最も簡単で、一度理解すれば、これらのスキームが一般的にどのように機能するかを理解できるからです。
このスキームは非常に基本的なものですが、非常に賢くて楽しいものです。手順は次のとおりです。
出力語彙を 2 つのセット (緑のリストと赤のリスト) に分割します。それらが均一にランダムに選択されていることを確認してください。
次に、緑色のリストの場合、すべてのロジットに $\delta$ を追加し、更新された dist からサンプリングします。

デコード時のリブレーション。
テキストに透かしが入っているかどうかを確認するには、緑色のセット内のトークンが表示される Z スコアを計算します。テキストに透かしが入っていない場合、緑のリスト内のテキストの期待値は $T/2$ です。$T$ はトークン数で、標準偏差は $\frac{\sqrt{T}}{2}$ です。したがって、この結果が得られるかどうかは、Z スコアにすぎません。
赤いトークンの確率がすでに 0.99 である場合 (これは大きなロジット リードになります)、緑への大きな $\delta$ ナッジがそれを追い越すことはありません。したがって、バイアスによって変更されるのは、選択肢が多く、一般的にエントロピーが高い単語だけです。
たとえば、LLM に詩を書くように依頼したとします。次のような文章が生成される可能性があります。
ウォーターマークが入っている場合は、「さわやかな」または「静かな」朝を生成する割合が、目に見えないほど高くなります (LLM プロバイダーが $\delta$ 値でウォーターマークをどの程度強く入れるかによって決まります)。 LLM が生成するすべての単語に対してこれを行うと、透かしに高いレベルの信頼性を得ることができます。
そうは言っても、グリーン/レッド リストは検出が簡単で、モデルの利用可能なエントロピーをより効率的に使用するスキームがあるため (したがって、検出が難しく、モデルの出力が変更される可能性が低くなります)、グリーン/レッド リストが実際に使用されるとは思えません。最もよく知られているスキームは、Google Deepmind によって開発され、Google によって運用環境で使用されている SynthID-Text です。
SynthID-Text はグリーン/レッド リストと同じ考え方ですが、Deepmind がトーナメント サンプリングと呼ぶわずかに異なるアプローチを使用しています。手順は次のとおりです。
各復号化ステップで、コンテキストの最後の $h$ トークンとともに秘密鍵をハッシュしてシードを生成します。そのシードは、すべての語彙トークンごとに $[0,1]$ の 1 つの $g$ 値を割り当てます。

アーナメント層。
通常どおりモデルのロジットを使用して候補出力トークンを描画します。
候補者がペアで対決するトーナメント ブラケットを実行し、そのレイヤーの $g$ 値が高い候補者が勝ち進みます。
トーナメントの勝者をデコードされたトークンとして出力します。
ウォーターマークを検出するには、秘密キーだけが必要です。テキスト内のすべてのトークンに関連付けられたシード値と $g$ 値を計算できます。
次に、結果をヌル分布と比較し、グリーン/レッド リスト検出と同様の標準化されたスコアを生成します。
SynthID は繰り返しコンテキスト マスキングも使用します。同じ $h$-token コンテキスト ウィンドウが応答中にすでに使用されている場合、実装はその位置にウォーターマークを付けることを拒否できます。これにより、繰り返されるコンテキストが同じバイアスを何度も受け取ることを防ぎますが、間違ったコンテキスト長を使用すると誤ってマスクがトリガーされ、信号が非表示になる可能性があるため、テストには非常に重要です。
SynthID は、すべての候補がモデル独自のディストリビューションから抽出されるため、グリーン/レッド リストよりも少し偽装されており、トーナメントでは、モデルがすでに発言したと考えられている単語のみを宣伝できます。モデルが次のトークンについてかなり確信を持っている場合、エントロピーは低く、透かしはあまりありません (グリーン/レッド リストと同様)。信号はエントロピーの高い単語ほど強くなります。これが、これらのスキームが検出を信頼できるようになるまでに相当量のテキストを必要とする理由でもあります。
モデルが次のトークンについてほぼ確実である場合、サンプリング ベースのウォーターマークで使用できるエントロピーはほとんどありません。信号は主にエントロピーの高い位置に蓄積されます。これが、これらのスキームが検出の信頼性を得る前に相当量のテキストを必要とする理由の 1 つです。
詩に戻って、モデルが 4 人の候補者を描画し、t を使用して 2 層トーナメントを実行するとします。

連続スコアのバリエーション:
したがって、トーナメントの $g$ 値により、「quiet」が発行されます。これは、単一の単語が出力の標準部分のように見えるため賢明ですが、秘密キーを使用すると、出力された単語がキー付きスコアと繰り返し一致するかどうかをテストできます。また、レイヤーの数を調整して検出可能性とモデルの品質の間のトレードオフを変更できるため、非常に優れています。
SynthID は、実稼働環境で実行されている私が知っている (または ChatGPT/Claude が知っている) 唯一のスキームであり、Google によって公開され、規模を拡大して実行されていることから、これが私の最良の候補であると思われます。
Anthropic がこれらの種類の統計的ウォーターマークのいずれかを実行しているかどうかを確認するために、Codex GPT-5.6-sol[0] を使用してウォーターマーク検出の文献にあるブラック ボックス テストのいくつかを実行しました。
Codex に Gloaguen らの検出アルゴリズムを実行してもらいました。この中で、クロードは次の 4 つの果物のいずれかを使用して文章を完成させるように強制されます。
111111111111111個の桃を食べました
私は222222222222222梅を選びました
...
接頭辞と繰り返される数字は変化しますが、果物の選択肢は固定されたままです。古典的なキー付き透かしがトークンを緑と赤のセットに分割する場合、特定のフルーツが異常に発生する可能性が高くなるか、低くなるはずです。
Sonnet テストでは、それぞれ 1,800 件の応答に対して、それぞれの組み合わせに 10 個のプレフィックス、9 桁のコンテキスト、および 20 個の呼び出しを使用しました。 Fable テストでは、1,200 件の応答に対して 6 桁のコンテキストを使用しました。各 p 値は 100,000 の順列から計算されました。
3 つとも明らかな陰性結果でした。ソネット テストは、この特定の設定におけるかなり大きく安定した緑/赤の影響を除外するのに十分な強度でした。弱いウォーターマーク、リクエスト間で変更されるキー、または異なる種類のウォーターマークを完全に除外するわけではありません。
次の実験では、SynthID のようなスキームで使用される固定コンテキスト ウィンドウを探しました。基本的なセ

tup はクロードに次のような行を強制的に返すようにしました。
赤緑のチェリーを食べました
ここで、赤は摂動ワードであり、その後に緑が $H$ 回繰り返されます。ウォーターマークが最後の $h$ トークンのみを参照する場合、$H$ が $h$ に達すると、赤はウォーターマークへの影響を停止するはずです。たとえば、$h=2$ の場合、$H=1$ では強い効果が期待されますが、$H=2$ ではほとんど効果がないか、まったく効果がないと予想されます。したがって、私たちが探しているのは、$H$ の値での急激な下落です。
私が実行した実際の実験例は次のとおりです。
より大きなスケールで、摂動ワードの変化がクロードの果物の選択をどれだけ強く変えるかを測定する $Z(H)$ を計算しました。ゼロに近い値は、検出可能な影響がないことを意味します。実験の実行で確認した $Z(1)=87.25$ は、テスト統計量が、ランダム化されたデータで期待される標準偏差を 87.25 (!!) 上回っていたことを意味します。
これは非常に強力な結果であり、$H=1$ では $p\およそ 0.00001$ になります。残念ながら、このテストは決定的なものではありませんでした。なぜなら、このテストは摂動ワードが重要であることを示しているだけであり、実際には $H$ の急激な減少は見られず、ゆっくりとした減少だけが見られたためです。これは、これがまったく透かしではなく、むしろモデルから来る単なる効果であることを意味している可能性があります。
フルスキャンでは、$H$ の各値で 2,592 個を含む、20,736 個の Sonnet 5 応答が使用されました。最も急激な曲がりは $H=2$ でした。そのため、新しい Sonnet 5 の応答と Sonnet 4.6 (Anthropic のヘルプセンター記事を読んだところによると、古いモデルなので透かしが入る可能性は低いです) を使用して再度テストしました。
この曲がりは新しい Sonnet 5 データに再び現れました。しかし、その効果は $H=2$ でも消えず、$Z(2)$ と $Z(3)$ の値は依然として巨大でした。 Sonnet 4.6 も非常に似た曲がりを示しました。
このため、結果は、最初に示唆されていた膨大な数よりもはるかに刺激的ではなくなります。 Sonnet 4.6 も同様に機能する可能性があります。

透かしが入っているため、真のネガティブ コントロールではありませんが、このパターンは Sonnet 5 に固有のものではないようで、きれいなコンテキスト ウィンドウの境界のようには見えません。最も可能性の高い説明は、赤や緑などの単語によって、クロードが果物の中から選ぶ方法が自然に変化するということです。
したがって、これはプロンプト効果の強力な検出でしたが、SynthID の確実な検出ではありませんでした。また、このテストでは確認できない別のウォーターマークも除外されません。
私の現在の最善の推測は、Anthropic がクロードがテキストを生成するときにトークンの選択を変更する秘密鍵の透かしを使用しているということです。しかし、その推測は主に Anthropic の説明と、私の実験での肯定的な結果とは対照的に、他のテストでの否定的な結果から来ています。
隠された Unicode や空白の証拠が見つからなかったため、いくつかのことを除外することができました。また、制約付き選択テストでは、安定した大きな緑/赤バイアスに反対することができました。そうは言っても、見かけの SynthID シグナルは、Sonnet 4.6 にも登場した強力なプロンプト エフェクトである可能性が高いことが判明したため、必ずしもウォーターマークを明確に識別できるわけではありません。
ウォーターマークのロールアウトがまだ完全に完了しているかどうか、またどのモデルで利用できるかは実際にはわかりません。 Anthropic 自身のヘルプセンターの記事によると、今後モデルには透かしが入れられ、

[切り捨てられた]

## Original Extract

How Claude’s watermarking (probably) works
Yesterday, Anthropic announced that they had started watermarking AI-generated content. Folks across the internet were particularly up in arms about it (I think rightfully so), especially because this apparently is happening to all Claude models whether or not you are in the EU. I wanted to investigate what they’re actually doing and whether it’s perceptible or changeable.
Anthropic provides some insight into their approach in their article How Claude marks AI-generated content . Though it doesn’t actually provide any technical details on the implementation, it provides some guidance to help draw a wide net around the scheme they’re using. The key clues Anthropic left in their help center article:
The scheme “weaves an imperceptible watermark directly into the text itself. You won’t see it, and it doesn’t change the meaning, quality, or readability of Claude’s response.”
The watermarking doesn’t work well on very short text.
The watermark seems to have started recently (August 2nd or later)
This helps us narrow down the possibilities quite a bit.
To actually get to the bottom of what Anthropic are doing, I realized that there are some interesting experiments you can run. Gloaguen et al. created specific tests to check for statistical watermarking, and we can also perform an analysis of Claude outputs to check for things like hidden unicode or whitespace results.
For good measure, I also downloaded a dump of my Claude chats (I’ve used Claude Code since 2/4/2025 and have recorded 1206 sessions) and did a quick comparison to see if there was any changepoint around early August that percepitbly changed the mix of tokens that Fable 5 output (my usual daily driver model). I wasn’t able to find any perceptible difference in this historical analysis, which confirms Anthropic’s claim that the watermarking is generally imperceptible unless using more specific tests.
No hidden unicode or whitespace
The second thing I checked is whether there’s hidden unicode or whitespace or punctuation patterns. This was pretty conclusive: they’re not doing something so simple.
An analysis across 7.2 million extracted prose characters (both on historical Claude Code text as well as generated Claude Code text on August 11) showed that the only unicode characters that were output by Claude were reasonable and part of day to day usage:
Curly quotation marks, apostrophes, and horizontal ellipses
Em/en dashes (A LOT of them unfortunately)
Accented or non-English characters
An audit by Codex, with me spot checking about 10 samples, found no anomalous instances that were consistent with a hidden Unicode watermark. I similarly found no whitespace encoding marks.
This evidence, combined with the fact that the watermarking is imperceptible and doesn’t work on short text, means that Anthropic is most likely using some form of statistical token watermarking.
Statistical watermarking schemes
There are a few different schemes that are available that can add watermarking to text. I’ll talk about the simplest version, the green/red list created in 2023 by Kirchenbauer et al. , because it’s the easiest to explain and once you understand it will allow you to understand how these schemes generally work.
This scheme is super basic, but it’s quite clever and fun. Here are the steps:
Split your output vocabulary into two sets: a green and a red list. Make sure they’re chosen uniformly at random.
Then for the green list, add $\delta$ to all of the logits and sample from the updated distribution at decoding time.
To figure out whether a text has been watermarked, then you compute the z-score that the tokens in the green set appear. If the text wasn’t watermarked, then the expected value of text in the green list is $T/2$, where $T$ is the token count, with a standard deviation of $\frac{\sqrt{T}}{2}$. So the suspiciousness of getting this outcome is just the z-score:
If some red token already has probability 0.99 (which would be a huge logit lead) a big $\delta$ nudge to the greens still wouldn’t overtake it. So the bias only changes words that have a lot of options and generally high entropy.
For an example, let’s say you asked your LLM to write a poem, you might have the following potential sentences that get generated:
If it was watermarked, you’d get an imperceptibly higher percentage of generating “crisp” or “quiet” morning (depending on how strongly the LLM provider decided to watermark with their $\delta$ value). Do this across all the words that an LLM is generating, and you can get high levels of confidence in your watermarking.
That being said, I don’t believe Green/red lists are used in practice because they’re easy to detect and there are schemes that use the model’s available entropy more efficiently (and thus harder to detect and less likely to change the outputs of the model). The most well known scheme is SynthID-Text which was developed by Google Deepmind and is used by Google in production.
SynthID-Text is the same idea as green/red lists, but using a slightly different approach that Deepmind calls tournament sampling. Here are the steps:
At each decoding step, hash a secret key together with the last $h$ tokens of context to produce a seed. That seed assigns every vocabulary token one $g$ value in $[0,1]$ per tournament layer.
Draw your candidate output tokens using the model’s logits as normal.
Run a tournament bracket where candidates face off in pairs, and the candidate with the higher $g$ value for that layer advances.
Output the tournament winner as the decoded token.
To detect the watermark, you just need the secret key. You can compute the seed value and the $g$ values associated with every token in the text:
Then compare the result with the null distribution and generate a standardized score similar to the Green/red list detection.
SynthID also uses repeated-context masking: if the same $h$-token context window has already been used during a response, the implementation can decline to watermark that position. This prevents a repeated context from receiving the same bias over and over, but it matters substantially for testing because using the wrong context length can accidentally trigger the mask and hide the signal.
SynthID is a bit more disguised than Green/red lists because every candidate is drawn from the model’s own distribution, so the tournament can only promote words the model already considered saying. When the model is pretty certain about the next token, there’s low entropy and not much watermarking (just like Green/red lists). The signal gets stronger for high entropy words, which is also why these schemes need a decent amount of text before detection becomes reliable.
When the model is nearly certain about the next token, there is little entropy available for any sampling-based watermark to use. The signal accumulates mainly at higher-entropy positions, which is one reason these schemes need a decent amount of text before detection becomes reliable.
Going back to our poem, say the model draws four candidates and we run a two-layer tournament using the continuous-score variant:
So “quiet” gets emitted because of its tournament $g$ values. It’s clever because any single word looks like a standard part of the output, but with the secret key you can test whether emitted words repeatedly align with the keyed scores. It’s also quite nice because you can adjust the number of layers to change the tradeoff between detectability and model quality.
SynthID is the only scheme that I know of (or ChatGPT/Claude knows of) which is running in production, and it seems like my best guess candidate given that it’s been publicized and run a scale by Google.
To see if Anthropic is running one of these kinds of statistical watermarks, I ran some of the black-box tests from the watermark detection literature with Codex GPT-5.6-sol[0].
I had Codex run Gloaguen et al.’s detection algorithm. In this, Claude is forced to complete sentences using one of four fruits:
I ate 111111111111111 peaches
I chose 222222222222222 plums
...
The prefix and repeated digit vary while the fruit alternatives remain fixed. If a classic keyed watermark partitions tokens into green and red sets, then a particular fruit should become unusually likely or unlikely.
The Sonnet tests used 10 prefixes, 9 digit contexts, and 20 calls for each combination, for 1,800 responses each. The Fable test used 6 digit contexts, for 1,200 responses. Each p-value was calculated from 100,000 permutations.
All three were clear negative results. The Sonnet test was strong enough to rule out a fairly large, stable Green/red effect in this particular setup. It does not rule out a weaker watermark, a key that changes between requests, or a different kind of watermark entirely.
The next experiment looked for the fixed context window used by SynthID-like schemes. The basic setup forced Claude to return a line such as
I ate red green cherries
Here red is a perturbation word and green is repeated $H$ times after it. If the watermark only looks at the last $h$ tokens, red should stop affecting the watermark once $H$ reaches $h$. For example, if $h=2$, I would expect a strong effect at $H=1$ and little or no effect at $H=2$. So the thing we’re looking for is a sharp drop at some value of $H$.
Here’s an actual example of the experiment I ran:
At larger scale, I calculated $Z(H)$, which measures how strongly changing the perturbation word changes Claude’s fruit choice. A value near zero would mean no detectable effect. $Z(1)=87.25$, which I saw in the experimental runs, means the test statistic was 87.25 (!!) standard deviations above what you would expect to see in randomized data:
That is an extremely strong result, with $H=1$ giving $p\approx0.00001$. Unfortunately, this test wasn’t conclusive because it only tells us that the perturbation word matters, and we didn’t actually see a sharp drop off on any $H$, only a slow decrease, which could mean that this isn’t a watermark at all, but rather just an effect that comes from the model.
The full scan used 20,736 Sonnet 5 responses, with 2,592 at each value of $H$. The sharpest bend was at $H=2$, so I tested it again using fresh Sonnet 5 responses and also on Sonnet 4.6 (which based on my reading of Anthropic’s help center article is less likely to be watermarked because it’s an older model):
The bend appeared again in the fresh Sonnet 5 data. But the effect did not disappear at $H=2$: the $Z(2)$ and $Z(3)$ values were still enormous. Sonnet 4.6 also showed a very similar bend.
This makes the result much less exciting than the huge numbers initially suggest. Sonnet 4.6 might also be watermarked, so it is not a true negative control, but it does seem the pattern is not unique to Sonnet 5 and does not look like a clean context-window boundary. The most likely explanation is that words like red and green naturally change how Claude chooses among fruits.
So this was a strong detection of a prompt effect, but not a positive detection of SynthID. It also doesn’t rule out a different watermark that this test cannot see.
My current best guess is that Anthropic is using a private-key watermark that changes token selection as Claude generates text. But that guess comes mostly from Anthropic’s description and negative results in other tests as opposed to a positive result in my experiments.
I was able to rule out a few things as I found no evidence of hidden Unicode or whitespace, and the constrained-choice tests argue against a large, stable Green/red bias. That said, the apparent SynthID signal turned out to be a likely strong prompt effect that also appeared in Sonnet 4.6, so it doesn’t necessarily positively identify a watermark.
Note that I’m not actually sure whether the watermarking rollout is fully complete yet and which models it’s available on. From Anthropic’s own help center article, it says that models are going to be watermarked going forward and that support for a

[truncated]
