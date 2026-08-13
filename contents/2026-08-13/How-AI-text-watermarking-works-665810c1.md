---
source: "https://declaude.org/watermarking/"
hn_url: "https://news.ycombinator.com/item?id=49292932"
title: "How AI text watermarking works"
article_title: "how AI text watermarking works: a visual guide"
author: "padolsey"
captured_at: "2026-08-13T23:31:38Z"
capture_tool: "hn-digest"
hn_id: 49292932
score: 1
comments: 0
posted_at: "2026-08-13T23:16:55Z"
tags:
  - hacker-news
  - translated
---

# How AI text watermarking works

- HN: [49292932](https://news.ycombinator.com/item?id=49292932)
- Source: [declaude.org](https://declaude.org/watermarking/)
- Score: 1
- Comments: 0
- Posted: 2026-08-13T23:16:55Z

## Translation

タイトル: AI テキスト透かしの仕組み
記事のタイトル: AI テキスト透かしの仕組み: ビジュアルガイド
説明: AI が生成したテキスト内に統計的透かしがどのように隠れるか、またそれらを消去するものについての優しい視覚的なガイドです。

記事本文:
← デクロード
AI テキスト透かしの仕組み 。
統計マークがどのように内部に隠されているかを、優しく視覚的に説明します。
生成されたテキストとそれを消去するもの。
NOPEの人々から、仲間として
を宣言する。
プレーンテキストに透かしを入れるのは不可能に思えます。テキストにはデータを隠すピクセルがありません。
コピーアンドペーストではメタデータは残りません。すべてのキャラクターがあなたの目の前にいます。
マークは一体どこへ行くのでしょうか？
それでいて痕跡は本物だ。 Google は Gemini アプリからのテキストに透かしを入れています。
2024 年以降の Web エクスペリエンス (この記事の執筆時点では、その API は
文書化された
例外)、2026 年 8 月以降、新しいクロード モデルではモデルにテキストがマークされます。
レベルであり、以前のモデルがその後ろに持ち込まれています。彼らは
目に見えないので、コピーしても生き残りますが、
彼らはキャラクターの中にまったく住んでいないので機能します。彼らはに住んでいます
それらの間の選択肢。
5 つの短いステップで、それぞれに突っ込みどころがあります。数を数えるほど難しいことはありません。
書くことは小さな選択の連続です
秘密鍵はそれらの選択に依存します
鍵を持っている人は誰でも数えることができます
1. 書くことは小さな選択の連続です
このステップの 1 つのアイデア: モデルは重み付きサイコロを転がして書き込みます
いくつかの単語の間には、それぞれ問題ありません。
モデルは文の途中では「次の単語」を知りません。候補リストがあり、
オートコンプリートなど、設定を使用します。これが本当の瞬間です。
文の終わり:
研究の結果はかなりのものでした
テキストのページには、単語ごとに 1 つずつ、このような小さなフォークが数百個含まれています。
いくつかのオプションは同様に問題ありません。そのたるみが原料です。誰が到達しても
ダイスランドがテキストの内容を変えずにテキスト内のパターンをどのように隠すことができるかに注目してください。
2. 秘密鍵はこれらの選択に依存します
このステップの 1 つのアイデア: 鍵は密かに最終候補リストを色付けし、
1つの色を与える

優しい小言。テキストはまだ普通に読めます。
これが古典的なレシピです (Kirchenbauer 2023; Google の SynthID が
より微妙なトーナメント形式のルートで同じ結末を迎えます)。各分岐点で、秘密鍵による数学が分割されます
候補の単語を
緑と
赤、任意のカラーリングのみ
キーホルダーも再現可能。その後、サイコロは緑の方向に少し傾きます。
研究の結果はかなりのものでした
これを卑劣なものにしているのは 2 つの点です。ナッジは穏やかです。赤い言葉でもまだ勝つ可能性があります。
可能性は少し低くなります。そして、色は単語の固定されたプロパティではありません: キー
直前の短い単語の連続から計算されるため、同じ候補は緑になります
1 つのプレフィックスの後に赤が続きます。
（二人の兄弟、同じ原理。
Google の SynthID (本番環境のもの) はナッジを小さな秘密に置き換えます
トーナメント : モデル独自のオッズからいくつかの候補者が抽出され、キーがそれらをスコア付けします。
そしてブラケットは、キーのドローを平均して、すべての単語のオッズが得られるように配置されています。
モデルが意図したものを正確に維持します。 OpenAIで構築されたアーロンソンのスキームはスキップする
それさえもキーからサイコロの出目自体を導き出します。異なる数学でも同じ
原則: マークは選択肢の中にあります。)
3. 鍵を持っている人は誰でもカウントできます
このステップの 1 つのアイデア: キーを使用すると、テキストの色を変更したり、
ただ数えるだけです。マークされたテキストが緑色になることが多すぎるのは幸運です。
検出ではテキストを読んだり、そのスタイルを判断したりしません。検出器は、
キーホルダーが単語の上に色を塗り、緑色になった数を数えます。マークなし
（または正しいキーがない場合）、緑は約半分の確率で勝つはずです。コイン投げ。こちらです
普通に見える段落。両方のキーを試してください。
バーがどれほど厳しいものであるかに注目してください。この段落は明らかに傾いていますが、それでも傾きすぎています
フラグに短い。検知器は誤警報がなくなるように調整されています

レア;の
代償として、短いテキスト、編集されたテキスト、または弱いマークが付けられたテキストが見逃されることが多く、55 単語であることです。
軽度の痩せだけでは十分な証拠ではありません。読み続けるを押してください: リーンはそのままです
増加する証拠に合わせてバーが滑り落ちる間、まったく同じです。長さは
テストの一部。 (そして、このデモでは、それがわかるように傾きが強く描かれています。
プロダクションマークはより緩やかに傾いているため、それに応じてより多くのテキストが必要になります。この中で
デモの 50/50 モデルでは、1,500 ワードの文書はわずか約 55% で緑のフラグを立てます: 小
リーンは長さによってのみ説得力を持ちます。それが短いテキストが重要である理由です。
誰にとっても本当に電話するのは難しい。）
4. 編集によりマークはどうなるか
このステップの 1 つのアイデアは、マークが手つかずの文言の中に存在するということです。
編集すると、ランが途切れた場所だけが消去され、それ以外の場所は消去されません。
各単語の色は、その直前の短い単語の続きから派生します (1 つずつ)。
スキームによっては一握りまで）。したがって、ポジションはショートの場合にのみ証拠としてカウントされます。
元の文言のウィンドウ (単語とその隣接単語) はそのまま残ります。
ここでは、ステップ 3 と同じ段落を 5 つの編集深度で示しています。スライダーをドラッグして、
ハイライト表示されたランが縮小するのを確認してください。ハイライトは、まだ文言が続いていることを意味します
オリジナルと正確に一致するため、検出器はそこでカウントできます。色あせたものはすべて新しい
コイン投げの騒音だけが数えられる言葉遣いです。
タイプミスを修正する ·
生き残ったウィンドウ: – %
実際の実装 (オープンモデル上の MarkLLM の KGW および EXP スキーム、
declaude の完全書き換えルート): ウィンドウの約 0.5% が生き残り、検出器の精度
本質的に確実な状態からコイントスに落ちます。出版された文献は次のことに同意しています。
これの形。軽い言い換えやワンパスの言い換えは、むしろ評価を薄めます。
それを削除する。キルヒェンバウアーらの実験では、十分な量を与えれば検出が回復した
テキスト、偶数付き

人間の言い換えは約 800 トークン後に再び検出可能になる
（約600ワード）。何を取り除くのか
マークは、オリジナルと文言を共有せずに再構成されたものです。
だからこそ、意味から書き換えるツール( declaude のような)
完全な書き換えルート）は、このマークのファミリーを実際に消去するものであり、なぜライトが消去されるのか
ほとんどのフレージングを維持するパスはありません。
1 つの境界は明確に述べられています。これらの数値は、私たちができるオープンな実装から得られたものです。
測る。 Anthropic の制作計画は非公開であるため、Anthropic 以外の者は誰も公開しません。
クロード自身のマークに対してこのテストを実行することはまだ可能です。私たちの実験が裏付けているのは、
このページではスキームのファミリーについて説明します。
5. これが実際に何を意味するか
このステップの 1 つの考え方は、検出はプライベートであり、確率的であり、
加工については、著作権ではありません。
キーホルダーの所有者のみが確認できます。あなたの先生、編集者、またはお気に入りの「AI」
「detector」の Web サイトではこのテストを実行できません。本物のチェックにはプロバイダーの認証が必要です
秘密キー、またはプロバイダーが実行するチェック サービス。 Googleは早期アクセスを実施している
SynthID の検出ポータル。 Anthropic 氏は、検出ツールが近々登場すると述べています。
透かしチェックは「AI 検出器」ではありません。 GPTZero のようなツールは次から推測します
スタイルがあり、信頼性が低いことで有名です。ウォーターマークはその逆で、意図的なものです。
キーゲート統計テスト。この2つを曖昧にしないでください。
見つかったマークは、「によって書き込まれた」ではなく、「によって処理された」ことを意味します。アントロピック独自の
文書には、人間のテキストはクロードによって校正または翻訳されただけであると記載されています
マーク。そして、古いモデル、短いパッセージ、または大幅な編集など、不在であることはさらに重要ではありません。
すべて本物の AI テキストでクリーンな結果が得られます。
短くて選択肢の少ないテキストは、あまり意味がありません。証拠は長くなるにつれて増えていきますが、
正しい継続が 1 つだけあるテキスト (コード、引用、事実のリスト) は、オファーを提供します。
サイコロの緩みが少なすぎてこんにちは

何でも入ってます。
特定のマークは書き換えを超えて存続します。単語自体をキーにしたスキーム
隣接するものよりもはるかによく耐えられます。同じ意味の書き換えは保持されます。
言葉がたくさんあるので、痕跡の多くは残っています。
(弱点は異なります。どこでも再利用されているカラーリングはリバース エンジニアリングできます。
十分な出力から。）その他は意味の中に隠れており、同じ意味の部分を書き換えたもの
それらを保存します。それらは別の問題です。
プロダクションパラメータは非公開です (コンテキストの長さ、キー構造、
しきい値）、ここでのすべてはスキームのファミリーを説明するものであり、何かを説明するものではありません。
具体的な展開。
ジェームズ・パドルシー著
いや、としては
デクロードの伴奏。インタラクティブな図形は教育モデルです
プロバイダーの実際のスキームではなく、例示的なパラメーターを使用します。
出典と詳細情報。
Kirchenbauer et al.、大規模言語モデルのウォーターマーク (ICML 2023) ·
Dathri et al.、LLM 出力を識別するためのスケーラブルな透かし
(SynthID-Text、Nature 2024) ·
Aaronson & Kirchner、透かし GPT 出力 (2022) ·
Kirchenbauer et al.、大規模言語モデルの透かしの信頼性について
(ICLR 2024) ·
Sadasivan 他、AI 生成テキストは確実に検出できるか? (2023) ·
Zhao et al.、「The Mark Fades: Adaptive Eevolutionary Paraphrase-based Attack」
(ACL 調査結果 2026) ·
人間的、どのように
クロードは AI によって生成されたコンテンツをマークします (ヘルプセンター、2026 年 8 月) ·
私たち自身の既知の鍵の実験: 再構成により KGW/EXP 検出が偶然に崩壊する
(AUC 0.99 → ≈0.5)、文脈自由ユニグラム マークは存続します (0.73 ～ 0.84)。アウトラインレベル
意味空間マークに関して私たちが知っている唯一の答えは再生です。
専門家向け: ステップ 4 の背後にある残存証拠モデル
判定は z ≈ f・√N・z₁ (生き残る部分 f、文書の長さ N、トークンごとの強度)
z₁)。数字は単語をカウントします。本当の検出

tors はモデル自身のトークナイザーをカウントします。
トークン。同じ形状です。

## Original Extract

A gentle visual guide to how statistical watermarks hide inside AI-generated text, and what erases them.

← declaude
How AI text watermarking works .
A gentle, visual walk through how a statistical mark hides inside
generated text, and what erases it.
From the folks at NOPE , as a companion
to declaude .
A watermark in plain text sounds impossible. Text has no pixels to hide data in, and
no metadata survives copy-and-paste; every character is right there in front of you.
Where could a mark possibly go?
And yet the marks are real. Google has watermarked text from the Gemini app and
web experience since 2024 (its API is, at the time of writing, a
documented
exception ), and as of August 2026, new Claude models mark text at the model
level, with earlier models being brought in behind them. They're
invisible, they survive copying, and they
work because they don't live in the characters at all. They live in the
choices between them .
Five short steps, each with something to poke at. Nothing harder than counting.
Writing is a series of small choices
A secret key leans on those choices
Whoever holds the key can count
1. Writing is a series of small choices
The one idea in this step: a model writes by rolling weighted dice
between several words that would each be fine.
When a model is mid-sentence, it doesn't know "the next word." It has a shortlist,
like autocomplete, with preferences. Here's a real kind of moment, one word from the
end of a sentence:
The results of the study were quite
A page of text contains hundreds of these little forks, one per word, and at many of
them several options are equally fine. That slack is the raw material. Whoever gets to
lean on how the dice land can hide a pattern in the text without changing what it says.
2. A secret key leans on those choices
The one idea in this step: the key secretly colours the shortlist and
gives one colour a gentle nudge. The text still reads normally.
Here is the classic recipe (Kirchenbauer 2023; Google's SynthID reaches the
same end by a subtler, tournament-style route). At each fork, secret-keyed maths splits
the candidate words into
green and
red , an arbitrary colouring only the
key-holder can reproduce. Then the dice get tilted a little toward green.
The results of the study were quite
Two things make this sneaky. The nudge is mild: a red word can still win, it's just
a little less likely. And the colouring is not a fixed property of the word: the key
computes it from a short run of the words just before, so the same candidate is green
after one prefix and red after another:
(Two siblings, same principle.
Google's SynthID (the one in production) replaces the nudge with a tiny secret
tournament : a few candidates are drawn from the model's own odds, the key scores them,
and the bracket is arranged so that, averaged over the key's draws, every word's odds
stay exactly what the model intended. Aaronson's scheme, built at OpenAI, skips
even that and derives the dice-rolls themselves from the key. Different maths, same
principle: the mark lives in the choices.)
3. Whoever holds the key can count
The one idea in this step: with the key, you can re-colour any text and
simply count. Marked text lands green too often to be luck.
Detection doesn't read the text or judge its style. The detector replays the
key-holder's colouring over the words and counts how many came up green. Without a mark
(or without the right key), green should win about half the time. A coin flip. Here's
an ordinary-looking paragraph; try both keys on it:
Note how demanding the bar is. This paragraph is visibly tilted and still too
short to flag. Detectors are tuned so that false alarms are vanishingly rare; the
price is that short, edited, or weakly marked text is often missed, and 55 words
of a mild lean isn't enough evidence. Press keep reading : the lean stays
exactly the same while the bar slides down to meet the growing evidence. Length is
part of the test. (And this demo's tilt is drawn strong so you can see it; a
production mark leans far more gently and needs correspondingly more text. In this
demo's 50/50 model, a 1,500-word document would flag at only ~55% green: small
leans become persuasive only through length, which is why short texts are
genuinely hard to call, for anyone.)
4. What editing does to the mark
The one idea in this step: the mark lives in runs of untouched wording.
Editing erases it exactly where the runs break, and nowhere else.
Each word's colouring is derived from a short run of the words just before it (one
to a handful, depending on the scheme). So a position only counts as evidence if a short
window of the original wording (the word plus its neighbours) survives intact.
Here is the same paragraph from step 3, at five edit depths. Drag the slider and
watch the highlighted runs shrink. A highlight means that run of wording still
matches the original exactly, so the detector can count there. Everything faded is new
wording, where there is nothing but coin-flip noise left to count.
fix typos ·
surviving windows: – %
On real implementations (MarkLLM's KGW and EXP schemes on an open model, washed by
declaude's full-rewrite route): about 0.5% of windows survive, and detector accuracy
falls from essentially certain to a coin flip. The published literature agrees on the
shape of this. Light or one-pass paraphrase dilutes the mark rather than
deleting it; in Kirchenbauer et al.'s experiments, detection recovered given enough
text, with even human paraphrase becoming detectable again after roughly 800 tokens
(about 600 words). What removes
the mark is re-composition that shares no runs of wording with the original.
That is why a tool that rewrites from the meaning (like declaude 's
full-rewrite route) is what actually erases this family of mark, and why a light
pass that keeps most of the phrasing does not.
One boundary stated plainly: those numbers come from open implementations we can
measure. Anthropic's production scheme is undisclosed, so no one outside Anthropic
can yet run this test against Claude's own mark. What our experiments support is the
mechanism, for the family of schemes this page describes.
5. What this means in practice
The one idea in this step: detection is private, probabilistic, and
about processing , not authorship.
Only the key-holder can check. Your teacher, editor, or favourite "AI
detector" website cannot run this test; an authentic check needs the provider's
secret key, or a checking service the provider runs. Google runs an early-access
detector portal for SynthID; Anthropic says detection tooling is forthcoming.
A watermark check is not an "AI detector." Tools like GPTZero guess from
style and are famously unreliable. A watermark is the opposite: a deliberate,
key-gated statistical test. Don't let the two blur.
A found mark means "processed by," not "written by." Anthropic's own
documentation notes that human text merely proofread or translated by Claude picks up
the mark. And absence proves even less: old models, short passages, or heavy editing
all yield clean results on genuine AI text.
Short and low-choice text carries little mark. Evidence grows with length,
and text with only one right continuation (code, quotations, lists of facts) offers
the dice too little slack to hide anything in.
Certain marks outlive a rewrite. Schemes keyed on the word itself
rather than its neighbours hold up far better: a same-meaning rewrite keeps
enough of the words that much of the mark survives.
(Their weakness is different: a colouring reused everywhere can be reverse-engineered
from enough output.) Others hide in the meaning, and a same-meaning rewrite partly
preserves them. Those are a different problem.
Production parameters are unpublished (context length, key structure,
thresholds), so everything here describes the family of schemes, not any
specific deployment.
Written by James Padolsey at
NOPE as an
accompaniment to declaude . The interactive figures are a teaching model
with illustrative parameters, not any provider's actual scheme.
Sources & further reading.
Kirchenbauer et al., A Watermark for Large Language Models (ICML 2023) ·
Dathathri et al., Scalable watermarking for identifying LLM outputs
(SynthID-Text, Nature 2024) ·
Aaronson & Kirchner, Watermarking GPT outputs (2022) ·
Kirchenbauer et al., On the Reliability of Watermarks for Large Language Models
(ICLR 2024) ·
Sadasivan et al., Can AI-Generated Text be Reliably Detected? (2023) ·
Zhao et al., The Mark Fades: Adaptive Evolutionary Paraphrase-based Attack
(ACL Findings 2026) ·
Anthropic, How
Claude marks AI-generated content (Help Center, Aug 2026) ·
Our own known-key experiments: re-composition collapses KGW/EXP detection to chance
(AUC 0.99 → ≈0.5), context-free unigram marks survive (0.73–0.84); outline-level
regeneration is the only answer we know for meaning-space marks.
For the specialist: the residual-evidence model behind the step-4
verdict is z ≈ f·√N·z₁ (surviving fraction f, document length N, per-token strength
z₁). The figures count words; real detectors count the model's own tokenizer's
tokens. Same shape.
