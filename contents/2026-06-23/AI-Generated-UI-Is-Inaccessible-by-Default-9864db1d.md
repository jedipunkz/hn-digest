---
source: "https://master.dev/blog/ai-generated-ui-is-inaccessible-by-default/"
hn_url: "https://news.ycombinator.com/item?id=48648479"
title: "AI-Generated UI Is Inaccessible by Default"
article_title: "AI-Generated UI Is Inaccessible by Default – Master.dev Blog"
author: "speckx"
captured_at: "2026-06-23T18:40:07Z"
capture_tool: "hn-digest"
hn_id: 48648479
score: 1
comments: 0
posted_at: "2026-06-23T17:39:15Z"
tags:
  - hacker-news
  - translated
---

# AI-Generated UI Is Inaccessible by Default

- HN: [48648479](https://news.ycombinator.com/item?id=48648479)
- Source: [master.dev](https://master.dev/blog/ai-generated-ui-is-inaccessible-by-default/)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T17:39:15Z

## Translation

タイトル: AI が生成した UI はデフォルトではアクセスできない
記事タイトル: AI が生成した UI はデフォルトではアクセスできない – Master.dev ブログ
説明: ありません

記事本文:
← Master.dev に戻る
コース
学ぶ
会員になる
ゲスト執筆
RSS
/ ブログ
アクセシビリティ AI React
AI によって生成された UI はデフォルトではアクセスできません
LLM で生成された React コンポーネントのセマンティックな正確性を実現する 5 層の強制システム
最近では、AI コード生成ツール (Claude Code、Codex、Cursor など) を使用して React サイドバー コンポーネントを 8 秒で生成できます。滑らかなホバー状態、回転する山形、調和のとれた間隔など、正しく見えます。ただし、DevTools でブラウザのアクセシビリティ ツリーを見てください。おそらく、ルート側の要素ツリーは、 role generic 、 name none 、 focusable false を受け取ります。スクリーン リーダー ユーザー、キーボード ナビゲーター、および音声コントロール ユーザーにとって、このコンポーネントは事実上存在しません。
これは、ほとんどの LLM が、支援技術が実際に読み取るレイヤーのセマンティック情報をほぼゼロに生成しながら、視覚的な出力を最適化するために発生します。この記事では、アーキテクチャ上の理由を説明し、プロンプト制約の 5 層の強制システムを示します。これらには、静的分析、ランタイム テスト、CI 統合、セマンティックな正確性を願望的ではなく自動的にするアクセス可能なコンポーネントの抽象化が含まれます。この例では React と Tailwind を使用しています。これは、ほとんどの AI ツールが出力するものだからですが、アクセシビリティ ツリーはフレームワークを考慮しません。受信した DOM を考慮します。
注意: 風景は一枚岩ではありません。 Vercel の v0 のような特殊なツールは、アクセス可能なプリミティブを生成パイプラインにハードコーディングし始めています。v0 は Radix 上に構築された shadcn/ui コンポーネントを出力します。つまり、その出力はデフォルトで Radix のアクセシビリティ コントラクトを継承します。
これは正しいアーキテクチャ上のアプローチであり、心強いものです。ただし、v0 は例外です。 ChatGPT、C など、ほとんどの開発者が日常的に使用する汎用ツール

laude、Copilot、および Cursor は、この記事で詳しく説明する <div> スープを依然として生成します。また、v0 の出力であっても、生成パイプラインがないため、出荷されたものが実際に動作することを確認する必要がなくなるため、ここで説明した検証レイヤーの恩恵を受けることができます。
これは、汎用 AI コード生成ツールが生成するものを表すナビゲーション サイドバーです。私は複数のツールでこのプロンプトのバリエーションをテストしました。構造パターンは一貫しています。
このコンポーネントのブラウザのアクセシビリティ ツリー
グループ
テキスト「設定」
グループ
グループ
テキスト「アカウント」
画像 (アクセス可能な名前なし)
グループ
テキスト「プロフィール」
テキスト「セキュリティ」
壊れたもの:
目印はありません。外側の <div> はナビゲーションの役割を生成しません。スクリーン リーダーのユーザーは、ランドマーク ナビゲーション経由でこのセクションにジャンプすることはできません。
見出しはありません。 「設定」は、 <h2> ではなく、スタイル付きの <div> です。見出しによるナビゲーションでは決して見つかりません。
リスト構造はありません。項目は <ul> / <li> にありません。 「リスト、2 項目」のコンテキストはありません。
トグルの役割が間違っています。アカウント <div> は generic にマップされます。インタラクティブとしては発表されていません。
集中できません。 <div> 要素はタブ オーダー内にありません。キーボード ユーザーはトグルに到達できません。
展開/折りたたみ状態はありません。 aria-expanded はありません。山形の回転は視覚的にのみ行われます。
支配関係はありません。トリガーをパネルにリンクする aria コントロールはありません。
キーボード操作はありません。 onKeyDown はありません。フォーカスされていても Enter キーや Space キーを押しても何も起こりません。
ラベルのないアイコン。 SVG には、aria-hidden 名とアクセス可能な名前の両方がありません。
偽リンク。 Profile と Security は、クリック ハンドラーを備えた <div> 要素です。リンクの役割がなく、フォーカス可能ではないため、新しいタブで開くことはできません。
29 行に 10 個の異なる障害が含まれています。正確な数は密度よりも重要です。ほぼすべての要素が意味的に間違っており、失敗はさらに増大します。スクリーン リーダー ユーザーとの遭遇

これでは、インタラクションのためのアフォーダンスのない、平坦で構造化されていないテキストが聞こえます。
過去 2 か月間、いくつかのツールをテストしたところ、このパターンが蔓延していました。 <button> や <a> の代わりに <div onClick> が、ほとんどの対話型コンポーネントで使用されます。 ARIA 状態属性の欠落はほぼ普遍的でした。ほとんどすべてのカスタム コントロールにはキーボード処理がありませんでした。ほとんどのレイアウトでランドマークが欠落していました。多くの場合、アイコンは代替テキストなしで出荷されます
汎用 AI ツールは、いくつかの最新モデルのバージョンが改善されており、6 か月前よりも優れたセマンティック HTML を生成しており、一部にはアクセシビリティを意識したシステム プロンプトが組み込まれ始めています。しかし、改善には一貫性がなく、デフォルトの出力には体系的な強制が必要なほどアクセスできないままです。
なぜそうなるのか: アクセシビリティ ツリー
ブラウザは HTML と CSS を受け取ると、2 つの主要な表現を構築します。レンダー ツリー (DOM および CSSOM から派生) によって、画面に描画される内容 (レイアウト、色、タイポグラフィ、ホバー状態、遷移) が決まります。これはビジュアル レイヤーであり、AI によって生成されたコードがうまく処理されます。
ただし、ブラウザは、アクセシビリティ ツリーという別個の並列構造も構築します。これは、フィルター処理され、意味的に強化された DOM 表現であり、 WAI-ARIA 、 HTML-AAM 、および Core AAM 仕様に従って構築されています。スクリーン リーダーがページを移動するとき、DOM やレンダー ツリーではなく、このツリーを読み取ります。各ノードには次の 4 つのプロパティがあります。
役割：これは何ですか？ (ボタン、リンク、ナビゲーション、タブ)
名前 : それは何と呼ばれますか? (テキストコンテンツ、aria-label、altから計算)
状態：どのような状態ですか？ (展開、選択、チェック、無効)
値: どのようなデータが保持されますか? (入力値、進捗値)
レンダートレ

e とアクセシビリティ ツリーは同じ DOM から構築されますが、異なるコンシューマにサービスを提供し、異なる情報を抽出します。レンダー ツリーは、cursor-pointer と hover:bg-gray-800 を考慮します。アクセシビリティ ツリーは、 <button> 、 aria-expanded 、および aria-controls を考慮します。 CSS を使用すると、<div> をボタンのように見せることができます。それを 1 つにできるのは HTML セマンティクスだけです。
< ボタン > アカウント </ ボタン >
<!--
役割: ボタン |名前: "アカウント" |フォーカス可能: true |起動: Enter、スペース
--> コード言語: HTML、XML ( xml )
< div onClick = {handleClick} > アカウント </ div >
<!--
役割: 汎用 |名前: なし |フォーカス可能: false |アクティベーション: なし
--> コード言語: HTML、XML ( xml )
同じピクセル。一つはドアです。もう一つは扉の絵です。
正確なメカニズムを経験的に特定することは困難ですが、一貫したバイアスを説明できるいくつかの要因が考えられます。
トレーニング データ : GitHub 上のほとんどの React コードは、CSS クラスを持つ <div> 要素を使用します。セマンティック HTML と ARIA は過小評価されており、これはおそらくトレーニング コーパスで過小評価されていることを意味します。
フィードバック ループ : 開発者と RLHF 評価者が AI 出力を評価するとき、ほぼ確実に視覚的に評価します。フィードバック信号は、意味論的な欠陥を損なうことなく、視覚的な忠実性を強化します。
トークンエコノミクス: <div onClick={fn}> は、 <button type="button" aria-expanded={isOpen} aria-controls="panel-id"> よりもトークンが少なくなります。特定の制約がなければ、モデルには追加のトークンを費やすインセンティブがありません。
AOM モデルがありません : モデルにはアクセシビリティ ツリーの表現がありません。支援技術にとってコードがどのような意味を持つかではなく、コードがどのようなものかをモデル化します。
これらは情報に基づいた仮説です。しかし、彼らが予測する高い視覚的忠実度、ほぼゼロの意味的忠実度を予測するパターンは、私が観察したものと一貫して一致しています。
これを理解すると、修正が明確になります。

意味上の正しさが失われる可能性があるすべての段階で強制します。
制約のないプロンプトは、制約のない出力を生成します。モデルは統計的なデフォルトの <div> スープに収束します。制約されたプロンプトは、出力スペースを意味的に有効なコンポーネントに絞り込みます。
このプロンプトは毎回手動で入力する必要はありません。これをワークスペース コンテキストにベイクします。カーソルは、プロジェクト ルートから .cursorrules ファイルを読み取ります。 GitHub Copilot は .github/copilot-instructions.md をサポートします。他のツールにも同様のメカニズムがあります。以下のプロンプトはこれらのファイルの 1 つに属し、すべての世代に自動的に適用されます。使用しているツールが永続コンテキストをサポートしていない場合は、レイヤー 2 5 への投資を強く主張する必要があります。
# コンポーネント生成ルール
React コンポーネントを生成しています。これらのルールを厳密に従ってください。
## HTML セマンティクス
- アクションには <button> を使用します。 <div onClick> や <span onClick> は絶対に行わないでください。
- ナビゲーションには <a href="..."> を使用します。決して <span onClick={navigate}> しないでください。
- ランドマークには <nav>、<main>、<aside>、<header>、<footer> を使用します。
- <h1> ～ <h6> は正しい階層順序で使用してください。レベルをスキップしないでください。
- リストには <ul>/<ol> と <li> を使用します。
- 表形式のデータには、<table>、<thead>、<tbody>、<th>、<td> を使用します。
- フォームには <form>、<fieldset>、<legend>、<label> を使用します。
- showModal() API を使用してモーダル ダイアログに <dialog> を使用します。
- 必要に応じて、簡単な開示には <details>/<summary> を使用します。
## アクセシビリティ
- すべてのインタラクティブな要素にはアクセス可能な名前が必要です
(表示テキスト、aria-label、または aria-labelledby)。
- すべてのフォーム入力には、関連付けられた <label> または aria-label が必要です。
- アイコンのみのボタン: ボタンに aria-label、アイコンに aria-hidden。
- 装飾画像: alt="" または aria-hidden="true"。
- 動的状態: aria-expanded、aria-selected、aria-checked、を使用します。
必要に応じて、aria-current、aria-disabled。
- aria-live="pol を使用します。

ite」はステータス メッセージの場合です。
- ヘルプ テキストとエラー メッセージには aria-descriptionby を使用します。
## キーボード操作
- すべてのインタラクティブな要素はキーボードでアクセスできる必要があります。
- フォーカスに表示されるスタイルを使用します。置き換えずにアウトラインを削除しないでください。
- 複合ウィジェット: WAI-ARIA オーサリング プラクティスに基づく矢印キー。
- モーダルはフォーカスをトラップし、閉じるときにフォーカスを復元する必要があります。
- エスケープではオーバーレイを閉じる必要があります。
## モーション
- 優先モーションを尊重します。モーションセーフを使用する: または
motion-reduce: トランジションに関する Tailwind バリアント
空間移動（変形、位置変更、スケーリング）。
ホバー/フォーカス時の単純な色の遷移は許容されます
モーションガードなし。
## ライブラリの設定
- 複雑なパターン (タブ、コンボボックス、ダイアログ、リストボックス、メニュー) の場合、
ビルドする代わりに Headless UI、Radix UI、または React Aria を使用してください
ゼロから。
- スタイル設定には Tailwind CSS を使用します。
- すべてのインタラクティブ要素にフォーカス表示リング スタイルを含めます。
## テスティン
[切り捨てられた]
これらの制約を使用して再生成されたサイドバーは次のとおりです。
ナビゲーションの「設定」
見出し「設定」（レベル 2）
リスト（2件）
リスト項目
「アカウント」ボタン（展開）
リージョン「アカウント」
リスト（2件）
リスト項目
リンク「プロフィール」
リスト項目
リンク「セキュリティ」
リスト項目
「設定」ボタン（折りたたまれています）
すべての要素には役割、名前、状態があります。これらはどれも React 固有のものではありません。プレーン HTML の同じ構造からは同じツリーが生成されます。
同じ <button> 、同じ aria-expanded 、同じ aria-controls 、同じアクセシビリティ ツリー。残りの例では React を使用しています。これは、AI 生成が最も普及している場所であるためですが、原則はすべてのエコシステムに同等のものがあります (Vue の eslint-plugin-vuejs-accessibility、SvelteKit の組み込み a11y 警告、および axe-core は、生成元に関係なく、レンダリングされた DOM に対して機能します)。
モデルが制約を無視する場合
これは定期的に発生します。 3 つの戦略:
ターグ

追加のフォローアップ: 最初から再生成しないでください。 「アカウントの切り替えは onClick を備えた div です。ボタンに置​​き換えて、aria-expanded と aria-controls を追加します。」という具体的な修正を求めるプロンプトが表示されます。
監査プロンプト: 生成後: 「このコンポーネントの WCAG 2.1 AA 違反を監査し、すべての問題を修正します。以下を確認してください: ボタンであるはずのインタラクティブ div、aria-expanded/aria-controls の欠落、キーボード サポートの欠落、フォーカス表示スタイルの欠落。」モデルは、正しいコードを最初から生成するよりも確実にコードをレビューします。
手動チェックリスト: コミットする前に、インタラクティブな要素 <button> または <a> はありますか?トグルには aria-expanded がありますか? Tab キーですべてのコントロールに移動してアクティブにすることができますか?ランドマークや見出しは存在しますか?アイコンには aria-hidden がありますか? 2分です。
レイヤ 1 が弱いか存在しない場合、インラインのコパイロットの提案とタブ入力が主な防御策となります。この場合、レイヤ 2 ～ 5 を制限するためのプロンプトは表示されません。このシステムは、個々の層で障害が発生する可能性があるという現実を考慮して設計されています。
ツールを使用してコードを直接チェックできます。これまで作業しているコードに最適なオプションは、eslint-plugin-jsx-a11y プラグインを使用した ESLint です。
npm install --save-dev eslint-plugin-jsx-a11y コード言語: Bash ( bash )
エラーとなる設定例を次に示します。

[切り捨てられた]

## Original Extract

It doesn

← Back to Master.dev
Courses
Learn
Become a Member
Guest Writing
RSS
/ BLOG
Accessibility AI React
AI-Generated UI Is Inaccessible by Default
A five-layer enforcement system for semantic correctness in LLM-generated React components
These days, an AI code-generation tool (e.g., Claude Code, Codex, Cursor) can produce a React sidebar component in 8 seconds. It looks correct: smooth hover states, rotating chevrons, harmonious spacing. But take a look at the browser’s accessibility tree in DevTools. Chances are the root side element tree receives: role generic , name none , focusable false . For screen reader users, keyboard navigators, and voice control users, the component effectively doesn’t exist.
This happens because most LLMs optimize for visual output while generating near-zero semantic information for the layer that assistive technologies actually read. This article explains the architectural reasons and presents a five-layer enforcement system of prompt constraints. These include static analysis, runtime testing, CI integration, and accessible component abstractions that make semantic correctness automatic rather than aspirational. The examples use React and Tailwind because that’s what most AI tools emit, but the accessibility tree doesn’t care about your framework. It cares about the DOM it receives.
A caveat: the landscape is not monolithic. Specialized tools like Vercel’s v0 have begun hardcoding accessible primitives into their generation pipelines v0 outputs shadcn/ui components built on Radix, which means its output inherits Radix’s accessibility contracts by default.
This is the right architectural approach, and it’s encouraging. But v0 is the exception. The general-purpose tools that most developers use daily, like ChatGPT, Claude, Copilot, and Cursor, still produce the <div> soup this article dissects. And even v0’s output benefits from the verification layers described here, because no generation pipeline eliminates the need to confirm that what is shipped actually works.
Here is a navigation sidebar representative of what general-purpose AI code generation tools produce. I’ve tested variations of this prompt across multiple tools; the structural patterns are consistent.
The browser’s accessibility tree for this component
group
text "Settings"
group
group
text "Account"
image (no accessible name)
group
text "Profile"
text "Security"
What’s broken:
No landmark. The outer <div> produces no navigation role. Screen reader users can’t jump to this section via landmark navigation.
No heading. “Settings” is a styled <div> , not an <h2> . Navigation by heading will never find it.
No list structure. The items aren’t in <ul> / <li> . No “list, 2 items” context.
Wrong role on the toggle. The Account <div> maps to generic . Not announced as interactive.
Not focusable. <div> elements aren’t in the tab order. Keyboard users can’t reach the toggle.
No expanded/collapsed state. No aria-expanded . The chevron rotation is visual-only.
No control relationship. No aria-controls linking the trigger to the panel.
No keyboard interaction. No onKeyDown . Even if focused, Enter and Space do nothing.
Unlabeled icon. The SVG lacks both aria-hidden and an accessible name.
Fake links. Profile and Security are <div> elements with click handlers. No link role, not focusable, can’t be opened in a new tab.
Ten distinct failures in twenty-nine lines. The exact count matters less than the density: nearly every element is semantically wrong, and the failures compound. A screen reader user encountering this hears flat, unstructured text with no affordance for interaction.
In my testing across several tools over the past two months, this pattern was pervasive. <div onClick> instead of <button> or <a> appeared in the vast majority of interactive components. Missing ARIA state attributes were nearly universal. Keyboard handling was absent from almost every custom control. Landmarks were missing from most layouts. Icons shipped without text alternatives more often than not
General-purpose AI tools are improving some recent model versions generate better semantic HTML than they did six months ago, and some are beginning to incorporate accessibility-aware system prompts. But improvement is inconsistent, and the default output remains inaccessible enough to require systematic enforcement.
Why This Happens: The Accessibility Tree
When the browser receives your HTML and CSS, it builds two primary representations. The render tree (derived from the DOM and CSSOM) determines what gets painted to the screen: layout, color, typography, hover states, and transitions. This is the visual layer, and it’s the one AI-generated code handles well.
But the browser also constructs a separate, parallel structure: the Accessibility Tree . This is a filtered, semantically-enriched representation of the DOM, built according to the WAI-ARIA , HTML-AAM , and Core AAM specifications. When a screen reader traverses a page, it reads this tree, not the DOM, not the render tree. Each node has four properties:
Role : What is this thing? ( button , link , navigation , tab )
Name : What is it called? (computed from text content, aria-label , alt )
State : What condition is it in? ( expanded , selected , checked , disabled )
Value : What data does it hold? (input value, progress value)
The render tree and the accessibility tree are built from the same DOM, but they serve different consumers and extract different information. The render tree cares about cursor-pointer and hover:bg-gray-800 . The accessibility tree cares about <button> , aria-expanded , and aria-controls . CSS can make a <div> look like a button. Only HTML semantics can make it be one.
< button > Account </ button >
<!--
role: button | name: "Account" | focusable: true | activation: Enter, Space
--> Code language: HTML, XML ( xml )
< div onClick = {handleClick} > Account </ div >
<!--
​​role: generic | name: none | focusable: false | activation: none
--> Code language: HTML, XML ( xml )
Same pixels. One is a door. The other is a painting of a door.
The exact mechanisms are difficult to isolate empirically, but several plausible factors explain the consistent bias:
Training data : Most React code on GitHub uses <div> elements with CSS classes. Semantic HTML and ARIA are underrepresented, which likely means they’re underrepresented in training corpora.
Feedback loops : When developers and RLHF evaluators assess AI output, they almost certainly evaluate visually. The feedback signal reinforces visual fidelity without penalizing semantic failures.
Token economics : <div onClick={fn}> is fewer tokens than <button type="button" aria-expanded={isOpen} aria-controls="panel-id"> . Absent specific constraints, the model has no incentive to spend extra tokens.
No AOM model : The model has no representation of the accessibility tree. It models what code looks like , not what code means to assistive technologies.
These are informed hypotheses. But the pattern they predict high visual fidelity, near-zero semantic fidelity, matches what I observe consistently.
Understanding this makes the fix clear: you need to enforce semantic correctness at every stage where it could be lost.
An unconstrained prompt produces unconstrained output. The model converges on its statistical default <div> soup. A constrained prompt narrows the output space to semantically valid components.
This prompt should not be typed manually each time. Bake it into your workspace context: Cursor reads a .cursorrules file from your project root. GitHub Copilot supports .github/copilot-instructions.md . Other tools have similar mechanisms. The prompt below belongs in one of those files, applied automatically to every generation. If your tool doesn’t support persistent context, you have a stronger argument for investing in Layers 2 5.
# Component Generation Rules
You are generating a React component. Follow these rules strictly.
## HTML Semantics
- Use <button> for actions. Never <div onClick> or <span onClick>.
- Use <a href="..."> for navigation. Never <span onClick={navigate}>.
- Use <nav>, <main>, <aside>, <header>, <footer> for landmarks.
- Use <h1>-<h6> in correct hierarchical order. Do not skip levels.
- Use <ul>/<ol> with <li> for lists.
- Use <table>, <thead>, <tbody>, <th>, <td> for tabular data.
- Use <form>, <fieldset>, <legend>, <label> for forms.
- Use <dialog> for modal dialogs with its showModal() API.
- Use <details>/<summary> for simple disclosures when appropriate.
## Accessibility
- Every interactive element must have an accessible name
(visible text, aria-label, or aria-labelledby).
- Every form input must have an associated <label> or aria-label.
- Icon-only buttons: aria-label on button, aria-hidden on icon.
- Decorative images: alt="" or aria-hidden="true".
- Dynamic state: use aria-expanded, aria-selected, aria-checked,
aria-current, aria-disabled as appropriate.
- Use aria-live="polite" for status messages.
- Use aria-describedby for help text and error messages.
## Keyboard Interaction
- All interactive elements must be keyboard accessible.
- Use focus-visible styles. Never remove outlines without replacement.
- Composite widgets: arrow keys per WAI-ARIA Authoring Practices.
- Modals must trap focus and restore it on close.
- Escape must close overlays.
## Motion
- Respect prefers-reduced-motion. Use motion-safe: or
motion-reduce: Tailwind variants on transitions involving
spatial movement (transforms, position changes, scaling).
Simple color transitions on hover/focus are acceptable
without motion guards.
## Library Preferences
- For complex patterns (tabs, combobox, dialog, listbox, menu),
use Headless UI, Radix UI, or React Aria instead of building
from scratch.
- Use Tailwind CSS for styling.
- Include focus-visible ring styles on all interactive elements.
## Testin
[truncated]
Here is the sidebar regenerated with these constraints:
navigation "Settings"
heading "Settings" (level 2)
list (2 items)
listitem
button "Account" (expanded)
region "Account"
list (2 items)
listitem
link "Profile"
listitem
link "Security"
listitem
button "Preferences" (collapsed)
Every element has a role, a name, and a state. None of this is React-specific; the same structure in plain HTML produces the same tree:
Same <button> , same aria-expanded , same aria-controls , same accessibility tree. The remaining examples use React because that’s where AI generation is most prevalent, but the principles have equivalents in every ecosystem ( eslint-plugin-vuejs-accessibility for Vue, SvelteKit’s built-in a11y warnings, and axe-core works against any rendered DOM regardless of origin).
When the Model Ignores Your Constraints
This happens regularly. Three strategies:
Targeted follow-up : Don’t regenerate from scratch. Prompt with a specific correction: “The Account toggle is a div with onClick. Replace it with a button and add aria-expanded and aria-controls.”
Audit prompt : After generation: “Audit this component for WCAG 2.1 AA violations and fix all issues. Check for: interactive divs that should be buttons, missing aria-expanded/aria-controls, missing keyboard support, missing focus-visible styles.” Models review code more reliably than they generate correct code from scratch.
Manual checklist : Before committing, are there interactive elements <button> or <a> ? Do toggles have aria-expanded ? Can you Tab to and activate every control? Do landmarks and headings exist? Do icons have aria-hidden ? Two minutes.
When Layer 1 is weak or absent, inline Copilot suggestions and tab completions, where there’s no prompt to constrain Layers 2 through 5, become your primary defense. The system is designed for the reality that any individual layer might fail.
You can be checking your code directly with tools. A great option for the code we’re working with so far is ESLint with the eslint-plugin-jsx-a11y plugin.
npm install --save-dev eslint-plugin-jsx-a11y Code language: Bash ( bash )
Here’s an example configuration that erro

[truncated]
