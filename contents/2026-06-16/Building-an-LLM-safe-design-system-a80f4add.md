---
source: "https://polar.sh/blog/orbit-llm-safe-design-system"
hn_url: "https://news.ycombinator.com/item?id=48558134"
title: "Building an LLM safe design system"
article_title: "Building an LLM safe design system | Polar"
author: "steventey"
captured_at: "2026-06-16T19:21:38Z"
capture_tool: "hn-digest"
hn_id: 48558134
score: 1
comments: 0
posted_at: "2026-06-16T16:48:45Z"
tags:
  - hacker-news
  - translated
---

# Building an LLM safe design system

- HN: [48558134](https://news.ycombinator.com/item?id=48558134)
- Source: [polar.sh](https://polar.sh/blog/orbit-llm-safe-design-system)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T16:48:45Z

## Translation

タイトル: LLM 安全設計システムの構築
記事のタイトル: LLM 安全設計システムの構築 |ポーラー
説明: スケーラブルで LLM に安全な設計システムを構築するという私たちの探求

記事本文:
LLM 安全設計システムの構築 | Polar Polar スタートアップ プログラムの紹介
Toggle Sidebar 概要 ドキュメント 価格 ブログ 会社 オープンソース Polar on X ログイン LLM 安全設計システムの構築
スケーラブルで LLM に安全な設計システムを構築するという私たちの探求
現在、Polar で出荷されている UI コードのほとんどは、ループ内で LLM を使用して書かれています。それはスピードにとって素晴らしいことです。設計システムが一貫性を考慮して構築されていない限り、一貫性はさらに難しくなります。
私たちは Orbit という新しいプロジェクトの初期段階にあり、まだ多くのことを検討中です。私たちはおそらく、いくつかの点については正しく、他の点については間違っています。この投稿はその背後にある考え方について、後で議論できるように新鮮なうちに書き留めたものです。
観察を開始するのは簡単です。問題は、LLM が CSS クラスや Tailwind クラスを作成できないことではありません。彼らはそれを流暢に書きます。問題は、彼らが根底にある決定を意識せずにそれを書いていることです。
LLM にカードの作成を依頼すると、 p-4 、rounded-lg 、 bg-gray-100 、 dark:bg-zinc-90​​0 、 text-gray-500 に到達します。どの値も妥当です。それらはどれも必ずしもあなたのものではありません。それを何百ものコンポーネントと何千もの世代にわたって掛け合わせると、インターフェースはゆっくりと千のわずかに異なるグレーに変化していきます。 CLAUDE.mdでそれを防ごうとしたにもかかわらず
そこで、私たちが Orbit で行っている賭けは次のとおりです。そもそも、ブランド外の決定をコードで表現することを困難にすることです。理想的には不可能に近い。値が実際に行った設計上の決定ではない場合、その値は CI に合格しないはずです。
明確にしておきたいのは、これは Tailwind を批判するものではありません。それは素晴らしいことだと思います。これは、CSS がこれまでに備えた中で最も人間工学に基づいたユーティリティであり、多くの Polar がこれを使用して構築されており、人間がマークアップの大部分を入力するプロジェクトで再びこれに手を伸ばすことになります。その開放感はag

人がキーボードの前にいるときの便利な機能。
問題は狭くて具体的です。LLM が入力を行うと、その同じオープンさがまさに私たちに不利に作用します。私たちは、Tailwind が悪いからといって、Tailwind から離れるつもりはありません。著者が変更されたため、制限しています。
私たちは、迅速に行動して反復したい場合には、Tailwind が最適なスタイリング アプローチであると考えています。ただし、この投稿は、成長するチームのためにコードベースを将来も保証し、エージェント LLM の時代に一貫性を確保するために加えなければならなかった変更についてのものです。
Tailwind クラスは文字列です。 className="flex p-4 bg-blue-500" のようなクラスは、コンパイラにヒットするまでは単なるテキストです。これがまさに、記述を高速化する理由であり、生成されたコードのリスクを高める原因でもあります。
文字列表面は、LLM にわずかに間違っている余地を無限に与えます。
p-4 、 p-5 、 p-[17px] 、 px-4 py-3 、すべて有効、すべて異なる間隔
bg-gray-100 、 bg-zinc-100 、 bg-neutral-100 、すべて有効、正規なし
dark: LLM が追加することを覚えておく必要があり、半分の確率で間違ってしまうバリアント
パレットを完全にバイパスする text-[#3b82f6] のような任意の値
これらはいずれも構文エラーではありません。それらはすべて糸くずを通過します。それらはすべてレンダリングされます。これらは、静的分析では検出できない 1 つの点で間違っています。それは、システム外にあるということです。 LLM には、型システムに何も指示がないため、グレーが oklch(0.96 0.003 264) であり、 bg-gray-100 ではないことを知る方法がありません。
文字列は lint ルールを記述するのが複雑です。終わりのない追跡は、通常、正規表現が考慮していない特殊なケースに行き着きます。一方、小道具はそうではありません。
避難ハッチは私たちが何度も戻ってくる部分です。 LLM が生の className またはインライン スタイルに落ちた瞬間、それを中心に構築したすべての保証が弱まってしまいます。また、LLM は避難ハッチが大好きです。訓練データには避難ハッチがたくさんあるからです。
クラス

決定ではなく価値観です
p-4 と --color-gray-100 にはもっと基本的な問題があり、それは誰が入力しているかに関係なく当てはまります。LLM の視点から少し離れてください。
デザインシステムは値を積み上げたものではありません。それは一連の決定です。カードはこの表面に置かれます。強調解除されたテキストにはこの色が使用されます。積み重ねられた要素間の隙間はこれです。価値は決定の結果であり、決定そのものではありません。
p-4 は値です。 「16 ピクセルのパディング」と表示されます。その理由、どこでそれが許可されるか、または何と一致する必要があるかについては述べられていません。 bg-gray-100 は値です。特定の 1 つのグレーであり、そのグレーがカードなのか、ホバー状態なのか、無効なコントロールなのか、それとも偶然なのかはわかりません。 CSS 変数ではこの問題は解決されません。 --color-gray-100: #f3f4f6 は、さらに適切な名前の値です。それはその色が何であるかを示すものであり、それが何のためにあるのかを示すものではありません。
値を作成すると、その決定は使用時点で蒸発します。 6 か月後、bg-gray-100 を使用する場所が 40 個ありますが、そのうちのどれが「カード」を意味するのかを知る方法はありません。カードの背景について考えを変えると、決定を編集するのではなく、色を取得することになります。その意図は、ツール、チームメイト、モデルが読み取れる場所には決して書き留められませんでした。
これが、Orbit のトークンが価値ではなく意図に基づいて名前が付けられている理由です。背景カードは決定です。これはカードが置かれる表面です。ライトモードまたはダークモードでどのヘクスに解決されるかは、名前の背後にある実装の詳細です。間隔も同様に機能します。 m 、 l 、 xl はスケール上の役割であり、たまたま気に入ったピクセル数ではありません。両方ともpadding="l"を使用する2つの要素は、同じ決定を下したことを宣言していますが、偶然に両方とも16pxを望んでいたわけではありません。
LLM が bg-gray-100 を渡したとき、何百もの妥当な近傍がある棚から値を選択しましたが、適切に選択するにはセンスが必要です。 LLM が発生するとき

nded 背景カードは、すでに行った決定のリストから決定を選択しました。味を求めているわけではありません。私たちは、構築しているものの名​​前を尋ねています。
ドキュメントは提案です。 CIは契約です。
明らかに最初の行動は、ルールを書き留めることです。システム プロンプトのスタイル ガイドの CLAUDE.md に「bg-gray-100 ではなく、グレーを使用してください」と入力します。これらすべてのバージョンがあります。それらは保持できません。
ドキュメントに記載された内容はすべて確率であり、保証ではありません。 LLM はそれを読み取り、コンテキスト内の他のすべてのものと比較して、ほとんどの場合それに従います。ほとんどの場合、それはデザインシステムではありません。何千世代にもわたってミスが積み重なり、ドリフトのためにすべての差分を手動で確認する作業に戻ります。
そこで私たちはより厳しい線を引き、それが Orbit の残りの部分にかかっている線です。実際に重要なルールは英語で書かれておらず、CI で実行される ESLint ルールとしてエンコードされています。これにより、1 つの決定的な契約が得られます。 PR が緑色の場合は、マージしても安全です。そして、対偶は、私たちが和解した部分です。何かが間違っていて、どのルールもそれを捕捉しなかったとしても、それは私たちのルールのギャップであり、作者の失敗ではありません。
私たちはルールを書くか、その出力をそのまま受け入れるかのどちらかです。 「しかし、ガイドラインにはそうはならないと書かれている」ということはありません。
これにより、誰が注意しなければならないかが変わります。人間であれ LLM であれ、すべての作成者がすべてのプロンプトについて意見のある規則を覚えていると信頼するのではなく、忘れたり、スキップしたり、無視したりすることのできないチェックにその意見を移動します。 LLM は、望むものを自由に書くことができます。私たちは、CI に合格したものだけが、喜んで出荷できるものであることを確認しています。
賭け: トークンを唯一の語彙にする
私たちは、Tailwind の代わりに、Meta のコンパイル時にタイプセーフなスタイル ライブラリである StyleX を試しています。ただし、StyleX はメカニズムであり、ポイントではありません。重要なのは、それによって何を構築できるかということです

n 上: 単一のプリミティブ <Box /> は、デザイン トークンを型指定されたプロパティとして受け入れますが、それ以外は受け入れません。
私たちのスタイリング API は、Shopify の Restyle システムから大きく影響を受けています。
< ボックス
flexDirection = "列"
ギャップ = " l "
パディング = " m "
背景色 = " 背景カード "
境界半径 = " m "
borderColor = " ボーダープライマリ "
boxShadow = " m "
>
< テキスト バリアント = "Heading-xs" color = "text-primary" >
カードタイトル
</ テキスト >
< Text color = " text-Secondary " > 説明 </ Text >
</ Box > ここでのすべての値は決定から導き出されます。 「padding」には「16px」は必要ありません。事前定義された一連のサイズが必要です。 backgroundColor は 16 進コードを受け取りません。実際に定義した色の名前を受け取ります。タイプはトークン定義から直接取得されます。
const spacing = stylex をエクスポートします。定義変数 ( {
なし: ' 0 ' 、
xs : ' 4px ' 、
s: ' 8px ' 、
m: ' 12px ' 、
l: ' 16px ' 、
xl : ' 24px ' 、
' 2xl ' : ' 32px ' 、
' 3xl ' : ' 48px ' 、
' 4xl ' : ' 64px ' 、
' 5xl ' : ' 96px ' 、
} )
const backgroundColors = stylex をエクスポートします。定義変数 ( {
' 背景-プライマリ ' : ' light-dark(hsl(233, 4%, 81%), hsl(233, 4%, 3.5%)) ' ,
' 背景カード ' : ' light-dark(hsl(240, 2.90%, 72.50%), hsl(233, 4%, 9.5%)) ' ,
// ...
} ) これが核となるアイデアです。デザイン上の決定は 1 か所に保存され、プロップ タイプが受け入れる唯一のものです。 LLM 生成の Orbit コードは、CSS の空間全体から選択されるわけではありません。私たちが書いた短いメニューから選択することです。オートコンプリートは有効なトークンを示します。タイプミスはタイプエラーであり、3週間後に発見された視覚的な低下ではありません。
裸の <div> を禁止する理由
これが、私たちが最も議論を重ねた考え方の一部であり、現在までに最も確信を持っている部分です。
制約されていないものがまだ座っている場合、ボックス上のプロップを制約しても何も起こりません

そのすぐ隣にあります。生の <div> は、任意の className、任意のインライン スタイル、任意の属性を受け入れます。それは真っ白なキャンバスです。そして、空白のキャンバスこそがオフシステムコードを可能にするものなのです。トークンは、それを回避する方法がない場合にのみ制約されます。<div> はトークンを回避する手段です。
つまり、賭けは「プロップを型付けする」ことだけではありません。それは、「型指定されていないコンテナを完全に削除する」です。 Box は、以前 <div> に手を伸ばしていた正確な場所に手を伸ばすものになります。たまたま 1 キーストローク短い、その隣にある平行で制約のないパスはありません。
これは個人よりも LLM にとって重要です。人間はコントリビューションガイドを一度読んで、「ここには生の div は書かない」と認識します。 LLM はプロンプトごとにコードベースを再検出し、デフォルトで最も抵抗の少ないパスを設定します。
<div className="..."> が利用可能な場合、それは数十年にわたるトレーニング データが教えてくれたものであり、おそらくそれを利用することになります。実際にデフォルトを移動する唯一の方法は、単に推奨されないだけでなく、生の要素を使用不可にすることです。
明らかな反論はセマンティクスです。 <div>、<section>、<nav>、<ul> を禁止すると、意味のあるアクセス可能な HTML が失われますか?いいえ、これが取引の価値のある部分です。ボックスはポリモーフィックです。私たちが禁止する要素のリストは、Box が as-prop を通じてレンダリングできるリストとまったく同じです。
< Box as = " nav " alignItems = " center " columnGap = " m " > … </ Box >
< ボックス as = " ul " flexDirection = " 列 " rowGap = " s " >
< ボックス as = " li " > アイテム </ ボックス >
</ Box > 適切な DOM props が入力され転送された状態で、DOM 内に実際の <nav> と実際の <ul> を取得できます。失うのは開放弦の表面であり、セマンティクスではありません。許可された要素の閉じたセットとして。裸の <nav> は開いたドアです。私たちは意味を保ってドアを閉めようとしています。
ESLi を使用してそれを強制します

nt ルール。
' Polar/no-raw-html-layout ' : ' error ' 、 <div /> の代わりに @polar-sh/orbit の <Box /> を使用してください。
これにより、Orbit 設計システムに従うことが保証されます。
これまでのところ、LLM の新しいコンテキスト ウィンドウで生き残る命令は、間違ったものをコンパイルに失敗させることだけです。
設計システムにない値が必要な場合は、システムをバイパスするのではなく、トークンを追加するための信号となることを望みます。私たちは、正当な脱出ハッチがどこで終わり、怠惰が始まるかについての線引きをまだ行っておりません。そして、学びながらそれを動かし続けることを期待しています。
LLM が忘れられないダークモード
トークン値をよく見てください: light-dark(hsl(233, 4%, 81%), hsl(233, 4%, 3.5%)) 。各色には、ブラウザーのネイティブな light-dark() CSS 関数によって解決される明暗値の両方が含まれます。
つまり、2 番目のスタイリング パスがまったくないため、覚えておくべき dark: バリアントはありません。 backgroundColor="background-card" を一度記述すると、両方のテーマで正しい値がレンダリングされます。 LLM は、ライト モードでは正しく表示され、ダーク モードでは機能しないコンポーネントを出荷することはできません。これは、コンポーネントが誤って動作するための個別のダーク モード コードがないためです。最も一般的なテーマのバグは、単に表現することができません。
まだ時期尚早なので、これを決定ではなく方向性として捉えてください。

[切り捨てられた]

## Original Extract

Our quest to build a scalable, LLM-safe design system

Building an LLM safe design system | Polar Introducing the Polar Startup Program
Toggle Sidebar Overview Documentation Pricing Blog Company Open Source Polar on X Login Building an LLM safe design system
Our quest to build a scalable, LLM-safe design system
Most of the UI code shipped at Polar today is written with an LLM in the loop. That is great for speed. It is harder on consistency, unless your design system is built for it.
We're early on a new one, called Orbit, and still figuring a lot of it out. We are probably right about a few things, and wrong about other. This post is about the thinking behind it, written down while it's fresh, so we can argue with it later.
The starting observation is simple. The problem is not that LLMs can't write CSS or Tailwind classes. They write it fluently. The problem is that they write it without being aware of the underlying decisions.
Ask an LLM to build a card and it will reach for p-4 , rounded-lg , bg-gray-100 , dark:bg-zinc-900 , text-gray-500 . Every value is reasonable. None of them is necessarily yours. Multiply that across hundreds of components and thousands of generations, and your interface slowly drifts into a thousand slightly different grays. Even though you've tried to prevent it in CLAUDE.md
So the bet we're making with Orbit is this: make it hard to express an off-brand decision in code in the first place. Ideally close to impossible. If a value isn't a design decision we've actually made, it should not pass our CI.
We want to make something very clear, this is not a knock on Tailwind. We think it's outstanding. It's the most ergonomic utility CSS has ever had, it's what a lot of Polar was built with, and we'd reach for it again on a project where humans type most of the markup. Its openness is a genuine feature when a person is at the keyboard.
The catch is narrow and specific: that same openness is exactly what works against us once an LLM is doing the typing. We're not steering away from Tailwind because it's bad. We're constraining it because our author changed.
We believe that Tailwind is the styling-approach to pick if you want to move fast & iterate. This post is however about the changes we’ve had to make to future-proof our codebase for a growing team and ensuring consistency in an era of agentic LLMs.
Tailwind classes are strings. Classes like className="flex p-4 bg-blue-500" are just text until it hits the compiler. That is exactly what makes it fast to write, and exactly what makes it risky for generated code.
A string surface gives an LLM infinite room to be slightly wrong:
p-4 , p-5 , p-[17px] , px-4 py-3 , all valid, all different spacing
bg-gray-100 , bg-zinc-100 , bg-neutral-100 , all valid, none canonical
dark: variants the LLM has to remember to add, and gets wrong half the time
arbitrary values like text-[#3b82f6] that bypass your palette entirely
None of these are syntax errors. They all pass lint. They all render. They are wrong in the one way static analysis can't catch: they are off-system. An LLM has no way to know that your gray is oklch(0.96 0.003 264) and not bg-gray-100 , because nothing in the type system tells it.
Strings are complex to write lint-rules for. A never-ending chase which usually ends up in special-cases your regex didn’t account for. Props on the other hand are not.
The escape hatches are the part we keep coming back to. The moment an LLM can drop to a raw className or an inline style, every guarantee you built around it gets weaker. And LLMs love escape hatches, because their training data is full of them.
A class is a value, not a decision
Step back from the LLM angle for a second, because there's a more basic problem with p-4 and --color-gray-100 , and it's true no matter who is typing.
A design system is not a pile of values. It's a set of decisions. Cards sit on this surface. De-emphasised text uses this color. The gap between stacked elements is this. The value is the consequence of the decision, never the decision itself.
p-4 is a value. It says "16 pixels of padding." It does not say why, or where it's allowed, or what it should match. bg-gray-100 is a value: one specific gray, carrying no idea of whether that gray is a card, a hover state, a disabled control, or a coincidence. A CSS variable doesn't fix this. --color-gray-100: #f3f4f6 is still a value with a nicer name. It tells you what the color is, never what it's for.
When you author in values, the decision evaporates at the point of use. Six months later you have 40 places using bg-gray-100 and no way to know which of them meant "card." Change your mind about card backgrounds and you're grepping a color, not editing a decision. The intent was never written down anywhere a tool, a teammate, or a model could read it back.
This is why Orbit's tokens are named for intent, not for value. background-card is a decision: this is the surface a card sits on. Which hex it resolves to in light or dark mode is an implementation detail living behind the name. Spacing works the same way. m , l , xl are roles on a scale, not pixel counts you happened to like. Two elements that both use padding="l" are declaring they made the same decision, not that they coincidentally both wanted 16px.
When an LLM handed bg-gray-100 it chose a value off a shelf with hundreds of plausible neighbours, and it needs taste to choose well. When an LLM handed background-card it chose a decision off a list of decisions we've already made. We're not asking it to have taste. We're asking it to name what it's building.
Docs are a suggestion. CI is a contract.
The obvious first move is to just write the rules down. Put "use our gray, not bg-gray-100" in CLAUDE.md, in a style guide, in the system prompt. We have versions of all of those. They don't hold.
Anything you put in a doc is a probability, not a guarantee. The LLM reads it, weighs it against everything else in its context, and follows it most of the time. Most of the time it is not a design system. Across thousands of generations the misses pile up, and you are back to reviewing every diff by hand for drift.
So we drew a harder line, and it's the line the rest of Orbit hangs off: the rules that actually matter aren't written in English, they're encoded as ESLint rules that run in CI. That gives us one deterministic contract. If a PR is green, it is safe to merge. And the contrapositive is the part we've made peace with: if something is off and no rule catches it, that's a gap in our rules, not a failure of the author.
We either write the rule or we live with the output. There is no "but the guidelines said not to."
This flips who has to be careful. Instead of trusting every author, human or LLM, to remember an opinionated convention on every prompt, we move the opinion into a check that can't be forgotten, skipped, or talked out of. The LLM is free to write anything it wants. We just make sure the only things that pass CI are things we'd be happy to ship.
The bet: make tokens the only vocabulary
We're trying out StyleX, Meta's compile-time, type-safe styling library, in place of Tailwind. But StyleX is the mechanism, not the point. The point is what it lets us build on top: a single primitive, <Box /> , that accepts design tokens as typed props and not much else.
Our styling API is heavily inspired by Shopify’s Restyle system.
< Box
flexDirection = " column "
gap = " l "
padding = " m "
backgroundColor = " background-card "
borderRadius = " m "
borderColor = " border-primary "
boxShadow = " m "
>
< Text variant = " heading-xs " color = " text-primary " >
Card title
</ Text >
< Text color = " text-secondary " > Description </ Text >
</ Box > Every value here is derived from a decision. “padding” does not take “16px”, it takes a set of predefined sizes. backgroundColor does not take a hex code, it takes the names of colors we have actually defined. The types come straight from the token definitions:
export const spacing = stylex . defineVars ( {
none : ' 0 ' ,
xs : ' 4px ' ,
s : ' 8px ' ,
m : ' 12px ' ,
l : ' 16px ' ,
xl : ' 24px ' ,
' 2xl ' : ' 32px ' ,
' 3xl ' : ' 48px ' ,
' 4xl ' : ' 64px ' ,
' 5xl ' : ' 96px ' ,
} )
export const backgroundColors = stylex . defineVars ( {
' background-primary ' : ' light-dark(hsl(233, 4%, 81%), hsl(233, 4%, 3.5%)) ' ,
' background-card ' : ' light-dark(hsl(240, 2.90%, 72.50%), hsl(233, 4%, 9.5%)) ' ,
// ...
} ) This is the core idea. The design decisions live in one place, and they are the only thing the prop types will accept. An LLM generating Orbit code is not choosing from the entire space of CSS. It is choosing from a short menu we wrote. Autocomplete shows it the valid tokens. A typo is a type error, not a visual regression discovered three weeks later.
Why we're banning the bare <div>
Here is the part of the thinking we've gone back and forth on the most, and the part we're most convinced by now.
Constraining the props on Box does nothing if the unconstrained thing is still sitting right next to it. A raw <div> accepts any className, any inline style, any attribute. It is a blank canvas. And a blank canvas is precisely what makes off-system code possible. Tokens only constrain you if there is no way around them, and a <div> is the way around them.
So the bet isn't only "make the props typed." It's "remove the untyped container entirely." Box becomes the thing you reach for in the exact place you used to reach for <div>. There is no parallel, unconstrained path sitting beside it that happens to be one keystroke shorter.
This matters more for an LLM than for a person. A human reads the contribution guide once and internalizes "we don't write raw divs here." An LLM rediscovers the codebase on every prompt, and it defaults to the path of least resistance.
If <div className="..."> is available, that is what decades of training data teached it, and that is what it will most likely resort to. The only way to actually move the default is to make the raw element unavailable, not just discouraged.
The obvious objection is semantics. If you ban <div>, <section>, <nav>, <ul>, do you lose meaningful, accessible HTML? No, and this is the part that makes the trade worth it. Box is polymorphic. The list of elements we ban is exactly the list Box can render through its as-prop:
< Box as = " nav " alignItems = " center " columnGap = " m " > … </ Box >
< Box as = " ul " flexDirection = " column " rowGap = " s " >
< Box as = " li " > Item </ Box >
</ Box > You still get a real <nav> and a real <ul> in the DOM, with the right DOM props typed and forwarded. What you lose is the open string surface, not the semantics. as is a closed set of allowed elements; a bare <nav> is an open door. We're trying to keep the meaning and close the door.
And we enforce it using ESLint rules.
' polar/no-raw-html-layout ' : ' error ' , Use <Box /> from @polar-sh/orbit instead of <div />.
This ensures we follow the Orbit design system.
Making the wrong thing fail to compile is, so far, the only instruction we've found that survives an LLM’s fresh context window.
If you need a value the design system doesn't have, we'd rather that be a signal to add a token than to bypass the system. We're still drawing the line on where legitimate escape hatches end and laziness begins, and we expect to keep moving it as we learn.
Dark mode the LLM can't forget
Look closely at the token values: light-dark(hsl(233, 4%, 81%), hsl(233, 4%, 3.5%)) . Each color carries both its light and dark value, resolved by the browser's native light-dark() CSS function.
That means there is no dark: variant to remember, because there is no second styling pass at all. You write backgroundColor="background-card" once, and the correct value renders in both themes. An LLM can't ship a component that looks right in light mode and broken in dark mode, because there is no separate dark-mode code for it to get wrong. The most common class of theming bug is simply not expressible.
It's early, so take this as a direction rather than a

[truncated]
