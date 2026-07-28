---
source: "https://win-vector.com/2026/07/28/llms-will-benefit-from-scratch-workspaces/"
hn_url: "https://news.ycombinator.com/item?id=49090156"
title: "LLMs Will Benefit from Scratch Workspaces"
article_title: "LLMs Will Benefit from Scratch Workspaces – Win Vector LLC"
author: "jmount"
captured_at: "2026-07-28T22:04:53Z"
capture_tool: "hn-digest"
hn_id: 49090156
score: 3
comments: 0
posted_at: "2026-07-28T21:28:44Z"
tags:
  - hacker-news
  - translated
---

# LLMs Will Benefit from Scratch Workspaces

- HN: [49090156](https://news.ycombinator.com/item?id=49090156)
- Source: [win-vector.com](https://win-vector.com/2026/07/28/llms-will-benefit-from-scratch-workspaces/)
- Score: 3
- Comments: 0
- Posted: 2026-07-28T21:28:44Z

## Translation

タイトル: LLM はスクラッチ ワークスペースから恩恵を受ける
記事のタイトル: LLM はスクラッチ ワークスペースから恩恵を受ける – Win Vector LLC
説明: 現在 (2026 年半ば) の LLM は非常に強力です。ただし、(エージェント ハーネスとループを追加する前は) LLM ベンダーは、効果的なスクラッチ ワークスペースが含まれているとは主張していません。現在の LLM 設計は、私たちがそれが何であるかを知る限りでは、おそらく注意/変圧器設計の拡張にとどまります。
[切り捨てられた]

記事本文:
LLM はスクラッチ ワークスペースから恩恵を受ける – Win Vector LLC
« 戻る

私たちについて
会社情報
データサイエンスに関するアドバイス、コンサルティング、トレーニング
LLM はスクラッチ ワークスペースから恩恵を受ける
ジョン・マウント著、2026 年 7 月 28 日 • ( コメントを残す )
現在 (2026 年半ば) の LLM は非常に強力です。ただし、(エージェント ハーネスとループを追加する前は) LLM ベンダーは、効果的なスクラッチ ワークスペースが含まれているとは主張していません。現在の LLM 設計は、私たちがそれが何であるかを知る限りでは、おそらく限られた数のアテンション ヘッドを備えたアテンション/トランス設計の拡張のままです。これらの LLM は (「思考モード」であっても)、効果的な変更可能または編集可能な非ベクター スクラッチ ワークスペースを直接組み込んでいないようです。
このスクラッチ ワークスペースの欠如を確認するには、注目が難しく、スクラッチ ペーパーでは簡単であるはずのシンプルで簡単なクエリを設計します。私たちが選んだのは、テキストからユニークな単語をソートすることです。私の推測では、この種のタスクは出力トークンを入力トークンに一致させるために注意力を消費するだろうと考えています。これは、注意/変換 LLM に対する意図的なポンピング補題タイプのプロンプトと考えてください。
効果を確認するために、OpenAI の GPT-5.4 nano モデルに次のプロンプトを与えました。
以下のユニークな単語を並べ替えてください。
「おお、ミューズよ、有名なトロイの町を略奪した後、広範囲を旅したあの独創的な英雄のことを教えてください。彼が訪れた多くの都市、そして彼がその風習や慣習をよく知っていた国々も多くありました。さらに彼は、自分の命を救い、部下を安全に帰国させようとして海で多くの苦しみを味わいました。しかし、どんなに力を尽くしても部下を救うことはできませんでした。彼らは太陽神の牛を食べるという自らの愚かさによって命を落としたのですから」ハイペリオン、それで神は彼らが家に帰ることを妨げたのだ。

ああ、ジョーブの娘よ、あなたが知っている情報源が何であれ。
さて、戦闘や難破で死を免れた者はユリシーズを除いて全員無事に家に帰りました。ユリシーズは妻と国に帰りたいと切望していましたが、女神カリプソによって拘束され、女神カリプソは彼を大きな洞窟に閉じ込め、彼と結婚しようとしたのです。しかし、年月が経つにつれ、神々は彼がイサカに戻るべきであると決着する時が来ました。しかし、その時でさえ、彼が自分の民の中にいたときでさえ、彼の悩みはまだ終わっていませんでした。それにもかかわらず、ネプチューンを除くすべての神々が彼を憐れみ始めており、ネプチューンは依然として彼を絶え間なく迫害し、彼を家に帰そうとはしませんでした。」
このテキストは、ホメーロスの『オデュッセイア』のサミュエル・バトラー翻訳のプロジェクト・グーテンベルクのコピーの始まりです。
この作業は人間が頭の中で実行することはほとんど不可能ですが、人間にペンと 3×5 のインデックス カードのパックが許可されていれば簡単です。私の主張は、現在の LLM についても同様である可能性があるということです。
LLM はタスクで失敗しましたが、失敗の兆候はありませんでした。色と疑問符はセッションの保存後に私が追加したものです。 LLM 応答はここにリンクされています。
「洞窟」という単語は引用文の中にありますが、作成された単語リストには残っていません。
「had」などの単語が重複しています。
プロンプトは、改良の実行が可能であることを示唆していました。改良を許可しても何も解決せず、さらなる問題が発生しました (「go got go」という順序など)。
効果的な変更可能なワーキング メモリが LLM に追加されると、LLM を利用したエージェントの機能がさらに大幅に向上することになります。インスピレーションを得るために、Soar コグニティブ アーキテクチャのアイデアのいくつかに戻ることをお勧めします。
私が LLM の使い方を知らないからといってこれを無視しないでください。私は仕事でそれらを使用していますが、それらを上手に使用する方法を知るには、いくつかの注意点を持っている必要があります

自分にとって適切なタスクと不適切なタスクについての知識を覚えておく。注意力が限られている場合の結果についてもう少し推論すると、プロンプトでより一貫性と制約のチェックが求められるため、丸ごと捏造された回答がより蔓延するのではないかと推測します。その考え方は、応答の一貫性と制約のチェックによって消費されるアテンション ヘッドが増えるほど、応答がプロンプトに関連しているか、または応答がトレーニング データ (または事実) に関連しているかを確認するために利用できるアテンション ヘッドが少なくなるということです。つまり、より多くの制限 (「注意してください」、「博士号の物理学者のように書く」、「結果を確認してください」、「このスタイルガイドを使用してください」) を追加すると、実際にはモデルがより説得力のある間違った答えを生成する可能性があります。 I know this contrary to common advice.
補足。 This task is easy for classic Unix shell-tools.たとえば、tr -cs '[:alnum:]' '[\n*]' | を通じてテキストを実行した結果は次のようになります。 tr -s '\n' | sort -u can be found here .テキスト処理では、適切なトークン化と単語の正規化を達成するには、これよりも少し苦労する必要がありますが、少なくともこの結果には単語「cave」が含まれており、大文字と小文字を区別する重複がなく、信頼性の高い並べ替え順序が得られます。
X で共有 (新しいウィンドウで開きます)
×
Share on LinkedIn (Opens in new window)
リンクトイン
Facebook で共有 (新しいウィンドウで開きます)
フェイスブック
Share on Reddit (Opens in new window)
レディット
Email a link to a friend (Opens in new window)
電子メール
読み込み中… カテゴリ: コンピューター サイエンスに関する意見のチュートリアル
Tagged as: LLM pumping lemma text Tools
LLM はスクラッチ ワークスペースから恩恵を受ける
もしかしたらヤコビアン予想の反例は珍しくないかもしれない
適合しないのはボックスですか、それとも「はい」ですか?
アナリティクスから十分な価値が得られない理由
最強のバッターは誰だ？不均一に収集されたデータから確率を推定する
発見する

Win Vector LLC のその他の情報
今すぐ購読して読み続け、完全なアーカイブにアクセスしてください。

## Original Extract

Current (mid-2026) LLMs are quite powerful. However, (prior to addition of agent harnesses and loops) LLMs vendors don't claim to include an effective scratch workspace. Current LLM designs, to the extent we know what they are, remain extensions of the attention/transformer design probably with a li
[truncated]

LLMs Will Benefit from Scratch Workspaces – Win Vector LLC
‹ return

About Us
Company Information
Data science advising, consulting, and training
LLMs Will Benefit from Scratch Workspaces
By John Mount on July 28, 2026 • ( Leave a comment )
Current (mid-2026) LLMs are quite powerful. However, (prior to addition of agent harnesses and loops) LLMs vendors don’t claim to include an effective scratch workspace. Current LLM designs, to the extent we know what they are, remain extensions of the attention/transformer design probably with a limited number of attention heads. These LLMs (even in “thinking mode”) don’t seem to directly incorporate an effective mutable or editable non-vector scratch workspace.
We can try to confirm this lack of scratch workspace by designing a simple and easy query that should be hard for attention and easy with scratch paper. We picked: sorting unique words from a text. My assumption is that this sort of task will consume attention heads matching output tokens to input tokens. Think of this as a deliberate pumping lemma type prompt for attention/transformer LLMs.
To confirm the effect I gave OpenAI’s GPT-5.4 nano Model the following prompt.
Please sort the unique words in the following.
“Tell me, O Muse, of that ingenious hero who travelled far and wide after he had sacked the famous town of Troy. Many cities did he visit, and many were the nations with whose manners and customs he was acquainted ; moreover he suffered much by sea while trying to save his own life and bring his men safely home; but do what he might he could not save his men, for they perished through their own sheer folly in eating the cattle of the Sun-god Hyperion; so the god prevented them from ever reaching home. Tell me, too, about all these things, oh daughter of Jove, from whatsoever source you may know them.
So now all who escaped death in battle or by shipwreck had got safely home except Ulysses, and he, though he was longing to return to his wife and country, was detained by the goddess Calypso, who had got him into a large cave and wanted to marry him. But as years went by, there came a time when the gods settled that he should go back to Ithaca; even then, however, when he was among his own people, his troubles were not yet over; nevertheless all the gods had now begun to pity him except Neptune, who still persecuted him without ceasing and would not let him get home.”
This text is the beginning of Project Gutenberg’s copy of Samuel Butler’s translation of Homer’s The Odyssey .
This task is almost impossible for a human to do in their head, but trivial if the human is allowed a pen and pack of 3×5 index cards. My argument is that it may be similar for current LLMs.
The LLM failed at the task with no indication of failure. The color and question marks were added by me after the session was saved. The LLM response is linked here .
The word “cave” is in the quote, but doesn’t survive to the produced word list.
There are duplicate words such as “had”.
The prompt implied a refinement run was possible. Allowing the refinement fixed nothing and introduced more problems (such as the ordering “go got go”).
If and when an effective mutable working memory is added to LLMs, we will see another large increase in LLM powered agent capabilities. I would recommend going back to some of the Soar cognitive architecture ideas for inspiration.
I ask: please don’t dismiss this as me not knowing how to use LLMs. I do use them for work, but part of knowing how to use them well is having some concrete knowledge of tasks that are and are not appropriate for them. Reasoning a bit more on the consequences of limited attention heads I would speculate that whole-cloth fabricated answers become more prevalent as the prompt asks for more consistency and constraint checks. The idea being: the more attention heads consumed by response consistency and constraint checking, the fewer attention heads are available to ensure the response is related to the prompt or the response is related to training data (or even facts). That is: adding more strictures (“be careful”, “write like a PhD physicist”, “check your result”, “use this style guide”) may in fact drive the model to produce more convincing more wrong answers. I know this contrary to common advice.
Side note. This task is easy for classic Unix shell-tools. For example, the result of running the text through tr -cs '[:alnum:]' '[\n*]' | tr -s '\n' | sort -u can be found here . In text processing one does have to work a bit harder than this to achieve proper tokenization and word regularization, but at least this result has the word “cave”, no case-sensitive duplicates, and reliable sorting order.
Share on X (Opens in new window)
X
Share on LinkedIn (Opens in new window)
LinkedIn
Share on Facebook (Opens in new window)
Facebook
Share on Reddit (Opens in new window)
Reddit
Email a link to a friend (Opens in new window)
Email
Loading… Categories: Computer Science Opinion Tutorials
Tagged as: LLM pumping lemma text Tools
LLMs Will Benefit from Scratch Workspaces
Maybe Jacobian Conjecture Counterexamples are not Rare
Is it the box or the “yes” that does not fit?
Why You Are Not Getting Full Value From Your Analytics
Who’s the Best Batter? Estimating Probabilities from Unevenly Collected Data
Discover more from Win Vector LLC
Subscribe now to keep reading and get access to the full archive.
