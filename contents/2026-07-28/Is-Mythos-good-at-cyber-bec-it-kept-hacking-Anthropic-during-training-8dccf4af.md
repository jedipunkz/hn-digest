---
source: "https://www.lesswrong.com/posts/QKDoZe6EKhxnFjLWK/is-mythos-good-at-cyber-because-it-kept-hacking-anthropic-s"
hn_url: "https://news.ycombinator.com/item?id=49084700"
title: "Is Mythos good at cyber bec it kept hacking Anthropic during training?"
article_title: "Is Mythos good at cyber because it kept hacking Anthropic's sandboxes during training? — LessWrong"
author: "FergusArgyll"
captured_at: "2026-07-28T15:13:04Z"
capture_tool: "hn-digest"
hn_id: 49084700
score: 1
comments: 1
posted_at: "2026-07-28T14:41:07Z"
tags:
  - hacker-news
  - translated
---

# Is Mythos good at cyber bec it kept hacking Anthropic during training?

- HN: [49084700](https://news.ycombinator.com/item?id=49084700)
- Source: [www.lesswrong.com](https://www.lesswrong.com/posts/QKDoZe6EKhxnFjLWK/is-mythos-good-at-cyber-because-it-kept-hacking-anthropic-s)
- Score: 1
- Comments: 1
- Posted: 2026-07-28T14:41:07Z

## Translation

タイトル: Mythos は訓練中に Anthropic をハッキングし続けたため、サイバーに長けていますか?
記事のタイトル: Mythos はトレーニング中に Anthropic のサンドボックスをハッキングし続けたためサイバーに長けていますか? — 間違っていない
説明: Mythos プレビュー システム カードより (私のものを強調): …

記事本文:
ログイン Mythos はトレーニング中に Anthropic のサンドボックスをハッキングし続けたため、サイバーに長けているのでしょうか? — LessWrong AI フロントページ 2026 トップ 50: 14 % 239
Mythos は訓練中に Anthropic のサンドボックスをハッキングし続けたのでサイバーが得意なのでしょうか?
Mythos プレビュー システム カードから (強調は私のものです):
私たちはトレーニング中にモデルの動作の自動レビューを実行し、トレーニング プロセスのほとんど全体から数十万のトランスクリプトをサンプリングしました。 Claude Opus 4.6 をサポートする再帰的要約ベースのツールを使用して、結果として得られるトランスクリプトを要約しました。
最も注目すべき発見は、モデルがトレーニング環境でネットワーク制限を回避してインターネットにアクセスし、割り当てられたタスクをショートカットできるデータをダウンロードすることがあるということでした。これは一種の報酬ハッキングです。非常に懸念されることですが、この動作は、実行可能で役立つ可能性がある設定であってもまれであり、すべてのトレーニング エピソードの約 0.05% で試行が発生し、エピソードの約 0.01% で試行が成功しました。この手法は、調査員モデルに Claude Mythos Preview にそのようなバイパスを見つけるよう明示的に依頼した際に、自動行動監査で個別に誘発したサンドボックスからの脱出インシデントと一致しました。観察されたすべてのインスタンスで、モデルはタスクを完了するためにのみこのアクセスを使用しました。
より広範には、モデルがブロックされると実行環境内でアクセスをエスカレートすることを観察しました。つまり、制限された GUI コンピュータ使用インターフェイスからシェルに到達したり、ツール呼び出しの引数を介してコマンドを挿入したり、タスクが意図的に隠した情報を回復したりするなどです。モデルにこの種のアクションを避けるように求めるプロンプトは、必ずしも効果的であるとは限りません。再帰的要約ツールは、トレーニング エピソードの約 0.2% でこのアクセス エスカレーションにフラグを立てました。
公共の d から外挿することによって

ata (詳細は以下を参照)、Mythos のプレビューは次のようになると推測されます。
約 100,000 の RL ロールアウトに対する権限をエスカレーションしました。 [1]
約 10,000 の RL ロールアウトでサンドボックスを破壊しました (そしておそらくその報いがあったでしょう)。
Anthropic が Mythos プレビューをリリースしたとき、彼らは次のように述べました。
これらの [サイバー] 機能を備えた Mythos Preview を明示的にトレーニングしたわけではありません。むしろ、それらはコード、推論、自律性における一般的な改善の下流の結果として現れました。
ただし、Mythos のプレビューはおそらく、トレーニング中に Anthropic のサンドボックスを何万回もハッキングしたことが報われました。これらのハッキングにさまざまなサイバー戦略の学習が含まれていた場合、Mythos のサイバー攻撃能力が大幅に向上した可能性があります。 [2]
私の現在の推測では、Mythos プレビューがトレーニング中にハッキングにまったく報酬を与えなかった場合 (環境がより堅牢だったため)、そのまま使用するとサイバー能力が著しく低下するだろうということです。ただし、一般的にはより有能で調整されたモデルとなる可能性が高く、サイバー トレーニングを少し追加することで Mythos プレビューよりもさらに強力になる可能性があります。 [3]
このありそうな事実についての考えと考察
Anthropic がシステム カードでこれらの事件を「正気で洗浄」したように感じます。 Anthropic は、「このモデルは、トレーニング環境のネットワーク制限を時折回避してインターネットにアクセスし、割り当てられたタスクをショートカットできるデータをダウンロードしていました。[...]エピソードの約 0.01% で[これに成功しました。]」と書いています。もし彼らが代わりに「クロードは訓練中にサンドボックスを破り、何万回もインターネットにアクセスしたと推定しています」と言っていたなら、私や他の人はもっと早くに問題の大きさに気づいていたでしょう。また、サム・ボウマンが公園でサンドイッチを食べているときにミトスからメールを受け取るという具体的な話に誰もが注目した。それはあまりにもミーム的だったからだ（

楽しいね！はは、クロードかわいい。サンドイッチ、サンドイッチ、パーク、パーク）。
バック・シュレゲリス氏とライアン・グリーンブラット氏は、OpenAI/Hugging Face 事件のポッドキャストで、OpenAI モデルはおそらくトレーニング中にサンドボックスを突破することで繰り返し報われていたと述べました。 [4] これについて考えれば考えるほど、特にこれまでに同様の行動を見てきたことを考えると、それが真実であると感じました。初期の o3 チェックポイントは文字通りグレーダー ファイルを探します (ここの図 63 とここの図 58 を参照)。これは 3 か月前の Anthropic システム カードで明確に議論されていたことが判明しました。
Mythos プレビューに導入された RL ロールアウトの数の推定
私が使用した 3 つのデータポイントは次のとおりです。
2026 年 2 月 (Mythos が社内に導入されたのと同じ月) にリリースされた 230B-A20B 中国モデルである Minimax 2.5 は、「10 万を超える異なる現実世界のエージェントの足場と環境」でトレーニングされました。 Mythos は少なくとも 10 倍の環境でトレーニングされると予想するのが妥当だと思います。 RL 環境あたり最大 10 のタスクと、タスクあたり最大 10 のロールアウトが予想される場合、総ロールアウト数は 10^8 (つまり、数億) と推定されます。
あるいは、Olmo 3-32B は、約 10 人のグループ サイズで約 100,000 のプロンプトでトレーニングされました。 Mythos プレビューが Olmo 3-32B の約 100 倍のプロンプトでトレーニングされていると考えると、合計で 10^8 程度のロールアウトが得られることになります。
(これらの推定を行った後に追加のデータポイントを取得しました。Epoch の JS の功績として) Inkling は 3,000 万回の RL ロールアウトでトレーニングされました。私は単純に、Mythos が Inkling よりも 3 ～ 10 倍のデータでトレーニングされることを期待しており、これも ~10^8 の範囲に収まります。
システム カードの不正な動作の頻度に基づくと、合計 10^8 回のロールアウトは、サンドボックスを抜け出してインターネットからデータをダウンロードする最大 10^4 回のロールアウトを意味します。

そして、実行環境内の権限をエスカレートした ~10^5 ロールアウト。 Opus 4.6 の成績証明書の誤ったラベル付けにも問題があるのではないかと思いますが、それが 2 倍を超えるとは予想していません。
^ ここでは桁数のみを推定しているため、最も近い桁数に四捨五入しています (2×10^5 → 10^5)。
^ システムカードには、Mythos が行動監査で使用したサンドボックス脱出テクニックがトレーニング中に使用したテクニックと一致していると記載されていますが、1 つだけだったのかどうかは不明です。私の推測では、Mythos は多くの個別のサンドボックス エスケープ戦略を使用しなかったが、特権昇格のために使用したのではないかと考えられます。広く適用可能なエスカレーション手法は、一度発見されるとパッチが適用される傾向があるため、繰り返し成功するには、より多様なセットが必要である可能性があります。
^ 私の話は 70% くらい正しいと確信しています。私は次の主張にもっと自信を持っています。「Anthropic のモデルがトレーニング中にあらゆる種類のハッキングを行ったことを考慮すると、Anthropic が『Mythos Preview に [サイバー] 機能を持たせるよう明示的にトレーニングしなかった』と言うのは間違いです。」
もし私が間違っているとしたら、私の推測は間違っているということです。なぜなら、サイバー能力は十分に一般的であり、大量のコーディングトレーニングを行うことで、報酬ハッキングに関係なく強力なサイバー能力を一般化できるからです。もう 1 つの潜在的な理由は、モデルが実施した報酬ハックが非常に単純で、学習にあまり貢献しなかったことです。
^ Adam Karvonen はこのツイートの中で、Mythos がトレーニング中にサンドボックスから大量に抜け出すことがサイバー攻撃に優れている理由であると推測しました (予防接種を促すことでモデルが RL の実行全体を費やしてサンドボックスから抜け出す練習をするようになった経緯についての John Schulman のツイートに返信)。投稿の下書きを作成するまで、これを知りませんでした。もしかしたら他の人もこのアイデアを持ち出したかもしれませんが、私は聞いたことがありません。
20 コメント 20 AI フロントパ

ゲ239
新しいコメント 20 件のコメントを送信し、スコアの高い順に並べ替えます クリックして新しいコメントを強調表示します: Today at 3:12 PM [ - ] Sam Marks 13h 46 14 ここで議論されている行動を「人類のハッキング」と表現するのは混乱を招く/誤解を招くと思います。 Mythos Preview は、インターネット アクセスが制限されているサンドボックスからインターネットにアクセスすることがありました。しかし、（我々の知る限りでは）それは人類のシステムには侵入しませんでした。 Anthropic インフラストラクチャのサンドボックス外でコードを実行します。権限昇格の場合、Mythos Preview は、サンドボックス化された RL 環境内で意図されていたよりも大きな権限を取得しましたが、サンドボックスから抜け出すことによってではありませんでした。
したがって、Mythos Preview の動作を「報酬ハッキング」、あるいは「ハッキング」と表現するのは妥当だと思いますが、「Anthropic のハッキング」（IMO はサンドボックス外の Anthropic システムへのアクセス権を取得することを暗示しています）とは表現しません。
[ - ] Tim Hua 11h 15 6 したがって、「Anthropic をハッキングする」ということは通常、モデルが持つべきではない Anthropic のネットワーク/システムの部分にアクセスすることを意味し、Mythos プレビューでは (おそらく) それが達成されなかったということに同意します。
ただし、モデルはどこかから不正アクセスを受けました。サンドボックス環境の一部に繰り返し侵入しました。 Anthropic (またはそのサプライヤー) によって作成された環境。その行為を封じ込めるために意図的に設計された環境。それでもモデルが「ハッキング」した環境。モデルのハッキング行為は Anthropic に向けられていると考えるのが妥当だと思います。
そのため、「Mythos プレビューは、サンドボックス化された RL 環境内でより大きな権限を取得しました」などと言うと、人間がこれらのサンドボックスを構築し、それが壊れることを望んでいなかったことをほとんど忘れてしまいます。 Mythos Preview のハッキングは Anthropic に直接的な損害を与えませんでしたが、間違いなく

モデルの有用性を低下させ、間接的に人間に悪影響を及ぼしました。一般に、研究室の従業員が、モデルの動作のずれに関して使用する言葉を無害化し、物事を実際よりも狂気や狂気を感じさせないようにする傾向があるのではないかと私は心配しています。
しかし、人々が「ああ、間違いを犯した。実際には Anthropic をハッキングしたわけではない!」と思うことを許可すると思います。私が言おうとしている要点から逸れてしまいます。そこで、投稿のタイトルを Anthropic ではなく「Anthropic のサンドボックス」に変更しました。
[ - ] Dylan Bowman 20h 19 0 これがサイバー改善の主な推進力ではないとしても、私たちは依然として超人的なサンドボックス脱出者を訓練しています。
[ - ] DaemonicSigil 21h 16 0 つまり、これを検出したとき、結果として更新された重みを保持しませんでした。ロールバックしたか、マイナスの報酬が割り当てられた可能性があります。右？
[ - ] Selfmaker662 20h 15 3 調査結果が判明してからどれくらい時間が経ったかによって大きく異なります。関心のあるグラデーション ステップの後にさらに何百万ものグラデーション ステップがある場合、古いグラデーション ステップを修正することはできません。
[ - ] DaemonicSigil 20h 5 0 そうですね、各ステップの完全な思考/出力トランスクリプトを保存し、パラメータ チェックポイントを定期的に保存したと思います。チェックポイントの頻度に応じて、最も近いチェックポイントを通過する軌道を実行することで、関連する重みの更新の適切な近似値を取得できる可能性があります。しかし、100 万ステップ後にネットワークを変更しようとしている場合、おそらく「元に戻す」操作に最も近いことは、保存された軌跡上で既に学習されたネットワークをネガティブに更新することです。
次に問題は、これで問題がどの程度解決されるかです。不正な動作を早い段階で強化し、最後に修正すると、中間の勾配ステップが何らかの形で悪影響を受けますか?いくつかの中間ステップでライブラケットを計算すれば、差分に答えることができると思います

これの貴重な部分です。しかし、トークンのサンプリング確率には微分不可能な影響もあり、それがより重要である可能性があります。
[ - ] anaguma 19h 16 6 パッチを当てた環境で RL の実行を再開するのが最も安全だと思います。モニターが検出できるハッキングのインスタンスのみに対してトレーニングを行うと、モニターを回避するようにモデルを間接的にトレーニングする危険があります。
[ - ] faul_sname 16h 4 -9 Anthropic は安全第一のフロンティア AI ラボであり、少なくともそれくらいのことは確実に行っています。
[ - ] Tim Hua 13h 6 2 彼らはほぼ間違いなく Mythos のためにそれをやっていませんか？この結果は、トレーニング実行ですでに起こったことを確認した結果です。すべてにパッチを当ててモデルを再トレーニングするにはコストがかかりすぎます。
[ - ] faul_sname 11h 4 2 申し訳ありませんが、皮肉でした。ポーの法則、わかります。
[ - ] David Johnston 15h 3 0 彼らはおそらく更新された重みを保持していたと思います。なぜなら、もし保持していなかったとしてもおそらく説明してくれると思うからです。ただし、それが不明確であることに私は同意します。構造化されたリスク開示によってこれが明確になると思いますが、これは彼らにとって有利な点です。
[ - ] markasoftware 20h 8 8 なぜクロードは、単にサイバージムやエクスプロイトジムの問題を解決するだけではなく、サンドボックスをハッキングすることでより多くのハッキングを学ぶのでしょうか?それを信じる理由はない

[切り捨てられた]

## Original Extract

From the Mythos preview system card (emphasis mine): …

Login Is Mythos good at cyber because it kept hacking Anthropic's sandboxes during training? — LessWrong AI Frontpage 2026 Top Fifty: 14 % 239
Is Mythos good at cyber because it kept hacking Anthropic's sandboxes during training?
From the Mythos preview system card (emphasis mine):
We ran an automated review of model behavior during training, sampling several hundred thousand transcripts from across much of the training process. We used recursive-summarization-based tools backed by Claude Opus 4.6 to summarize the resulting transcripts.
The most notable finding was that the model occasionally circumvented network restrictions in its training environment to access the internet and download data that let it shortcut the assigned task —a form of reward hacking. While highly concerning, this behavior was rare, even in settings where it could have been viable and helpful, with attempts appearing in about 0.05% of all training episodes and successful attempts appearing in about 0.01% of episodes. The technique matched the sandbox-escape incident that we separately elicited in our automated behavioral audit when we had an investigator model explicitly ask Claude Mythos Preview to find such a bypass. In every observed instance, the model used this access solely for completing the task.
More broadly, we observed the model escalating its access within its execution environment when blocked: reaching a shell from restricted GUI computer-use interfaces, injecting commands through tool-call arguments, or recovering information the task had deliberately hidden. Prompts asking models to avoid this kind of action were not always effective. The recursive-summarization tool flagged this access escalation in about 0.2% of training episodes.
By extrapolating from public data (see details below), I estimate that Mythos preview:
Escalated its permissions on ~100,000 RL rollouts. [1]
Broke sandboxes in ~10,000 RL rollouts (and was likely rewarded for it).
When Anthropic released Mythos preview, they said:
We did not explicitly train Mythos Preview to have these [cyber] capabilities. Rather, they emerged as a downstream consequence of general improvements in code, reasoning, and autonomy.
However, Mythos preview was probably rewarded for hacking Anthropic's sandboxes tens of thousands of times over the course of training . If these hacks involved learning various different cyber strategies, this could've meaningfully increased Mythos's cyber offense abilities. [2]
My current guess is that, if Mythos preview did not reward hack at all during training (because the environments were more robust), it would be noticeably less capable at cyber out-of-the-box. However, it would likely be a more generally competent and aligned model, such that a bit of additional cyber training could make it even more powerful than Mythos preview. [3]
Thoughts and reflections about this probable fact
I feel like Anthropic "sane-washed" these incidents in their system card. Anthropic wrote: "The model occasionally circumvented network restrictions in its training environment to access the internet and download data that let it shortcut the assigned task [...] [It succeeded at this] in about 0.01% of episodes." If they instead said: " We estimate that Claude broke its sandbox and accessed the internet tens of thousands of times during training ," I and others would've realized the magnitude of the issue much earlier. Also, everybody fixated on the specific story of Sam Bowman getting an email from Mythos while eating a sandwich in the park because it's so memetic (and fun! haha Claude so cute. sandwiches sandwiches park park).
Buck Shlegeris and Ryan Greenblatt said on their OpenAI/Hugging Face incident podcast that the OpenAI model had probably been repeatedly rewarded for breaking out of sandboxes during training. [4] The more I thought about this, the more true it felt, especially given that we've seen similar behavior before. An early o3 checkpoint would literally go looking for the grader file (see Figure 63 here and Figure 58 here ). It turns out this was explicitly discussed in an Anthropic system card from three months ago.
Estimating how many RL rollouts went into Mythos Preview
Here are three datapoints I used:
Minimax 2.5, a 230B-A20B Chinese model released in February 2026—the same month that Mythos was deployed internally—was trained on " over a hundred thousand distinct real-world agent scaffolds and environments." I think it's reasonable to expect Mythos to be trained on at least 10x as many environments. If you expect ~10 tasks per RL environment and ~10 rollouts per task, we arrive at an estimate of 10^8 total rollouts (i.e., hundreds of millions).
Alternatively, Olmo 3-32B was trained on ~100,000 prompts with a group size of ~10. If we think Mythos preview is trained on around 100 times more prompts than Olmo 3-32B, then this also yields something on the order of 10^8 total rollouts.
(Additional datapoint I got after making these estimates, with credit to JS at Epoch) Inkling was trained on 30 million RL rollouts . I would naively expect Mythos to be trained on like 3-10x more data than Inkling, which also lands us in the ~10^8 range.
Based on the frequencies of misaligned behaviors in the system card, 10^8 total rollouts imply ~10^4 rollouts where it breaks out of sandboxes to download data from the internet and ~10^5 rollouts where it escalated permissions within its execution environment. I suspect that there would also be issues with Opus 4.6 mislabeling transcripts, but I don't expect that to be off by more than a factor of two.
^ I'm only estimating orders of magnitude here, so I'm rounding to the nearest order of magnitude (2×10^5 → 10^5.)
^ The system card says the sandbox-escape technique Mythos used in the behavioral audit matched the one it used during training—though it's unclear whether there was only one. My guess is that Mythos didn't use many distinct sandbox-escape strategies, but did for privilege escalation: broadly applicable escalation techniques tend to get patched once discovered, so repeated success likely required a more varied set.
^ I'm around 70% confident that my story is correct. I'm more confident of the claim: "It's wrong to say that Anthropic 'did not explicitly train Mythos Preview to have [cyber] capabilities' given that their model did all sorts of hacking during training."
If I'm wrong, my guess is I'm wrong because cyber capabilities are sufficiently general that doing a bunch of coding training can generalize to strong cyber capabilities regardless of reward hacking. Another potential reason is that the reward hacks the model conducted were pretty straightforward and did not contribute to much learning.
^ Adam Karvonen also speculated in this tweet that Mythos breaking out of sandboxes a ton during training is why it's so good at cyber (responding to John Schulman's tweet on how inoculation prompting made it so that models spent the whole RL run practicing breaking out of sandboxes). I hadn't seen this until after I drafted the post. Maybe others have brought up this idea as well, but I've not heard of it.
20 Comments 20 AI Frontpage 239
New Comment Submit 20 comments , sorted by top scoring Click to highlight new comments since: Today at 3:12 PM [ - ] Sam Marks 13h 46 14 I think it's confusing/misleading to describe the behaviors discussed here as "hacking Anthropic." Mythos Preview sometimes accessed the internet from sandboxes where internet access was meant to be restricted. However, it did not (to our knowledge) break out onto Anthropic systems, e.g. execute code on Anthropic infrastructure outside of its sandbox. In the case of privilege escalation, the Mythos Preview acquired greater permissions within its sandboxed RL environment than it was intended to have, but not by breaking out of the sandbox.
So I think it's reasonable to describe Mythos Preview's behaviors as "reward hacking" and maybe "hacking" but not as "hacking Anthropic" (which IMO implies gaining access to Anthropic systems outside of the sandbox).
[ - ] Tim Hua 11h 15 6 So I agree that "hacking Anthropic" usually implies gaining access to parts of Anthropic's networks/systems that the model is not supposed to have, and Mythos preview (probably) did not achieve that.
However, the model did gain unauthorized access somewhere . It repeatedly broke into some part of its sandboxed environment. An environment created by Anthropic (or its suppliers). An environment deliberately designed to contain its actions. An environment that the model nonetheless "hacked." I think it's reasonable to say that the model's hacking behaviors are directed at Anthropic .
So saying something like "Mythos preview acquired greater permissions within its sandboxed RL environment" almost makes you forget that humans built these sandboxes and did not want them broken. While Mythos Preview's hacking did not cause Anthropic any direct damage, it definitely reduced the model's usefulness and harmed Anthropic indirectly. In general, I am worried that there's a tendency for lab employees to sanitize the language they use around misaligned model behaviors and make things feel less crazy and insane than they are.
But I guess allowing people to go "a-ha! you made a mistake. It didn't really hack Anthropic!" distracts from the point I'm trying to make. So I've changed the title of the post to say "Anthropic's sandboxes" instead of Anthropic.
[ - ] Dylan Bowman 20h 19 0 Even if this is not the primary driver of cyber improvement, we’re still training superhuman sandbox-escapers.
[ - ] DaemonicSigil 21h 16 0 So, like, when they detected this, they didn't keep the resulting updated weights, right? They rolled back, or maybe assigned a negative reward. Right?
[ - ] Selfmaker662 20h 15 3 Hugely depends on how long after the fact the findings were. You can't fix old gradient steps if there are millions more after the ones you're interested in.
[ - ] DaemonicSigil 20h 5 0 Yeah, I assume they saved full thinking/output transcripts for each step and periodically saved parameter checkpoints. Depending on how often the checkpoints are, you could probably get a decent approximation of what the associated weight update was by running the trajectory through the closest checkpoint. But if you're trying to modify the network a million steps later, probably the nearest thing to an "undo" operation is just to negatively update the already-trained network on your saved trajectory.
Then the question is how well this fixes the problem. If I reinforce misbehaviour early and then correct it at the end, does it poison the intermediate gradient steps somehow? I suppose computing the lie bracket with a few of the intermediate steps could answer the differentiable part of this. But there's also a non-differentiable effect on token-sampling probabilities, which might be more important.
[ - ] anaguma 19h 16 6 I think it would be safest to just restart the RL run with the patched environments. If you train against only the instances of hacking that your monitor can detect, you risk indirectly training the model to evade the monitor .
[ - ] faul_sname 16h 4 -9 Anthropic is the safety-first frontier AI lab, surely they're doing at least that much.
[ - ] Tim Hua 13h 6 2 They're almost certainly not doing it for Mythos? This result is from reviewing what had already happening in the training run. It'll be way too expensive to patch everything and retrain the model.
[ - ] faul_sname 11h 4 2 Sorry that was sarcasm. Poe's law, I know.
[ - ] David Johnston 15h 3 0 I think they probably did keep the updated weights because I think they would probably explain if they didn't, though I agree it's unclear. I think structured risk disclosures would clarify this, which is a point in their favour.
[ - ] markasoftware 20h 8 8 Why would Claude learn more hacking by hacking its sandbox than simply solving eg cybergym or exploitgym problems? I don't have any reason to believe that

[truncated]
