---
source: "https://www.lightningjar.com/blog/ast-as-html"
hn_url: "https://news.ycombinator.com/item?id=48794460"
title: "HTML as a native data format for LLMs"
article_title: "HTML as a Native Data Format for LLMs | AST-as-HTML - LJ"
author: "kevinpeckham"
captured_at: "2026-07-05T15:07:01Z"
capture_tool: "hn-digest"
hn_id: 48794460
score: 2
comments: 0
posted_at: "2026-07-05T14:17:03Z"
tags:
  - hacker-news
  - translated
---

# HTML as a native data format for LLMs

- HN: [48794460](https://news.ycombinator.com/item?id=48794460)
- Source: [www.lightningjar.com](https://www.lightningjar.com/blog/ast-as-html)
- Score: 2
- Comments: 0
- Posted: 2026-07-05T14:17:03Z

## Translation

タイトル: LLM のネイティブ データ形式としての HTML
記事のタイトル: LLM のネイティブ データ形式としての HTML | HTML としての AST - LJ
説明: LLM は、事前にトレーニングされた HTML の高度な流暢さを備えています。そのため、ドキュメント プラットフォームの AI アシスタントがテンプレート レイアウトを読み書きするときに、ツリーを JSON ではなくマークアップとしてエンコードします。ここ

記事本文:
-->
LLM のネイティブ データ形式としての HTML | HTML としての AST - LJ
Lightning Jar をハックする - Web スタジオ
サイトナビゲーション
LLM のネイティブ データ形式としての HTML: データを JSON ではなくマークアップでエンコードする理由
私たちは、ブランドが承認したレール内で AI アシスタントがマーケティング ドキュメント (チラシ、パンフレット、概要資料) をデザインし、人間がレンダリングされたページをクリックして入力することでドキュメント プラットフォームを構築しました。エージェントにテンプレート (ドキュメントを構築するための構造レイアウト) を作成してもらうことは、次の 1 つの時代遅れな決定に左右されることが判明しました。
テンプレートは JSON ではなく HTML としてエンコードされます。そして、エージェントの主な編集ツールは「全体を書き直す」ことです。
これは、構造化データに基づいてエージェントを構築することに関する現在のアドバイスのほとんどを覆すものです。また、これは、より安価で堅牢な選択肢であることも判明しました。つまり、消費されるトークンが少なく、応答が速く、編集ごとのデータの整合性が向上します。この投稿はその理由についてです。
私たちのシステムのドキュメント テンプレートはツリーです。ドキュメントにはページが含まれ、ページにはブロックが含まれ、ブロックにはテキスト アトム、画像アトム、スタイル付きコンテナ、および再利用可能なウィジェットを参照するスロットが含まれます。すべてのノードには、CSS クラス、コピーの長さの割り当て、タイプスケールの選択、スロットの制約などの属性があります。テーマは CSS 変数を通じてツリーをペイントするため、テンプレートで色がハードコードされることはありません。 bg-primary と書かれており、アクティブなブランド テーマがそれを意味します。
私たちは、会話形式でこれらのツリーを構築および再構築できるアシスタントを必要としていました。「左側に丸いコンテンツ ボックス、右側に統計パネルを備えた全幅のウィジェットを設計してください」。そして、人間の編集と同じ元に戻すセマンティクスと検証を使用して、その出力が同じエディターに表示されるようにしたいと考えていました。
明らかなデザインと出荷しなかった理由
教科書的なアプローチは、ツリーを JSON として保存し、

詳細なツール API をモデル化します。
insertNode(親Id、タイプ、インデックス)
setAttribute(nodeId, キー, 値)
moveNode(nodeId, newParentId, インデックス)
RemoveNode(nodeId) このようなエージェントを構築しました。これらは機能しますが、次の 3 つの予測可能な点でパフォーマンスが低下します。
オーダーメイドのスキーマをゼロから教えることになります。すべてのノード タイプ、すべての属性、すべての包含ルールをプロンプト内で詳しく説明する必要があり、モデルの流暢さはプロンプトが何を購入したかによってのみ決まります。トレーニング中に JSON スキーマが 0 回参照されました。
細分化されたツールは細分化された失敗を招きます。 12 ノードのレイアウトを構築するには、12 回の往復が必要です。各呼び出しは、古い ID、間違った親、2 つの呼び出し前にシフトされたインデックスを参照する可能性があります。ツリーは 11 の中間状態を通過します。それぞれの状態でエージェントが無効な場所に取り残される可能性があり、それぞれの状態でレンダラーが生き残らなければならない可能性があります。
モデルはその作業を「見る」ことができません。 Mutation-by-tool-call では、現在のツリーのモデルの画像は、それ自身の呼び出し履歴から精神的に再構築されたものになります。ドリフトは避けられません。
私たちのテンプレートは、小さな属性文法を使用してプレーン HTML にシリアル化されます。
<div data-type="block" data-name="feature-callout" id="wgt-root">
<div data-type="コンテナアトム" data-name="コンテンツボックス" id="wgtコンテンツ"
data-container-classes="rounded-2xl flex-[2] bg-primary p-6">
<div data-type="text-atom" data-name="Heading" id="wgt-Heading"
data-text-element-tag="h2" data-text-style="見出し-2"
data-max-length="60"></div>
</div>
</div> すべてのノードは要素です。 data-type はノードの種類を示し、data-name は安定した人間のラベル、data-* は属性を持ち、パーサーはこれを内部 AST との間で変換します。これは私たちのエージェントの仕事よりも前から行われていました。これは、テンプレートが人間が編集可能なマークアップ ビューを往復できるようにするために存在しました。
テンプレート作成エージェントを構築したときに、そのエージェントもプライマリにしました

私は恥ずかしいほど率直にこう言いました。
set_template_markup(マークアップ: 文字列、概要: 文字列)
// "テンプレート全体を新しいマークアップに置き換えます。"それでおしまい。モデルは現在のマークアップ (リクエストごとに新しく送信されるため、ユーザーが実際に見るものを常に編集します) を読み取り、完全な新しいツリーを書き込み、クライアントはそれをワンショットで検証して適用します。
結果は、プロンプトをほとんど必要としないことに私たちを驚かせました。 LLM は、ネスト、属性、クラス文字列、ID など、HTML に関して事前にトレーニングされた高度な流暢さを備えています。私たちはフォーマットを教えたわけではありません。モデルがすでにネイティブで話しているものを借りました。プロンプトは、構文ではなくセマンティクス (ウィジェット スロットの意味、どのテーマ トークンが存在するか、長さのバジェットがどのように機能するか) にバジェットを費やします。ビジュアル エディターが使用する同じ構成ファイルから生成された文法ダイジェストにより、この 2 つの相違が防止されます。
また、各編集はツリー全体であるため、中間状態は存在しません。編集が一貫しているか拒否され、境界で 1 つの解析が行われます。
const parsed = parseMarkupToAst(マークアップ);
if (parsed?.type !== "block") destroy("ウィジェットには単一のブロック ルートが必要です");
editor.markupCurrent = format(build(parsed)); // 人間の編集と同じファネル すべてのエージェントの編集は、すべての人間の編集が使用する正確なパイプラインを通過します。同じ検証、同じ元に戻す、同じリアクティブ プレビュー。エージェントは、独自の書き込みパスを持つ特権アクターではありません。それはただの別の作者です。
屋根裏の密閉コンテナ
JSON を扱うのは、屋根裏部屋で閉じた保管コンテナを探し回るようなものです。ラベルは蓋の内側にあります。それが何であるかを知るためには各コンテナを開かないといけません。また、JSON はネストしているため、コンテナ内のコンテナ内でコンテナを開くことになります。コンテナーは、パッキングリストを持参した場合にのみ意味を持ちます。キーの名前とスキーマを事前に知らなければ、適切な質問をすることさえできません。そしていつ

これで探索は完了です。出口は、ラベルのない同一の蓋 ( }]}} を並べることです。どの蓋にも、何を閉じるかは記載されていません。屋根裏部屋の奥深くで、蓋が 1 つ置き忘れられ、スタック全体が破損しています。
HTML はすべてのコンテナの外側にラベルを付けます。
<div data-type="text-atom" data-name="Heading" data-max-length="60"> ノードのタイプ、名前、および制約は、内部に入る前に読み取ることができ、コンテナーはその名前を付けたまま閉じます。モデルに JSON ツリーのインベントリを要求すると、モデルが再構築します。 HTML ツリーのインベントリを要求すると、ラベルが読み取られます。この違いが、ツリー全体ツールが機能する理由のほとんどです。ラベル付きの自己終了 HTML コンテナーの 60 個のノードを生成することは、モデルが 10 億回実行していることです。また、すべての蓋がそれが何に属しているかを示しているため、蓋を正しく取得するのは簡単です。
皮肉なことに、私たちは常にこのツリーを LLM に引き渡すことを計画していましたが、初期のプロトタイプはまだ引き渡されていませんでした。テンプレートは手動で作成およびデバッグする必要があったため、私たちが最初に解決していた問題は人間による可読性でした。ツリーの外側にラベルを付けると、人が読み書きしやすくなりました。この選択により、LLM の可読性に関する隠された真実が明らかになりました。人間の作成者の可読性とモデルの流暢性は同じ性質です。
(ただし、屋根裏部屋についての考えは抱いてください。また戻ってきます。屋根裏部屋は物を保管するのに最適な場所です。屋根裏部屋で仕事をしたくないだけです。)
見栄えは美しいだけでなく実用的です
燃焼されるトークンが少なくなりました。プロンプトは、モデルがすでに知っているスキーマを教える必要がなくなりました。
往復が少なくなります。これほど流暢なモデルは、突然変異ごとにツリーを組み立てるのではなく、1 回のパスでツリー全体を作成できます。
より確実な転送。スリップされたタグは、その後のすべてを静かに破損するのではなく、パーサーで大声で失敗します。
読み取りが速くなります。モデルはラベル i をスキャンします

構造を再構築するのではなく。
すべての解析と再構築のサイクルにおける整合性の向上。ノードが自らを宣言するツリーは、一目で確認できるツリーです。
1 つの属性の微調整にはツリー全体の書き換えは無駄であり、1 つのクラスを変更するために 60 個のノードを書き換えると、残りの 59 個のノードで転写ドリフトが発生します。したがって、手術器具は 1 つだけです。
set_node_attributes(nodeId: string,attributes: Record<string, value>) このツールの構築にはほとんど費用がかからないため、一時停止する価値があります。 Web アプリケーションでは、HTML を編集するための機構がすでにブラウザーに付属しています ( getElementById 、 setAttribute 、 querySelector )。これらの API は高速で、実戦テストされており、社内のツリー突然変異ライブラリでは匹敵するものがないほど詳細に文書化されています。ノード アドレス指定スキームを発明したり、属性パッチャーを作成したりする必要はありませんでした。プラットフォームはすでに数十年をかけてそれらを完成させてきました。
構造編集用の合計 2 つのツール。実際には、モデルはガイダンスなしで正しく選択されます。大きな変更には書き直しが必要で、微調整にはパッチが必要です。これを、モデルが 12 個の突然変異の中から選択して正しく構成する必要があるツール カタログ アプローチと比較してください。
その後、同じ文法が 2 番目の恩恵をもたらしました。後でドキュメント アシスタントに会話中にまったく新しいウィジェットを作成させると、新しい形式を教えることなく、同じマークアップでウィジェットが作成され、同じパーサーによって検証されました。
これは、ほとんどの投稿がスキップする部分です。 3 つの実際のバグはどれも有益です:
1. ID は神聖なものですが、フォーマッタは ID をそのように扱っていませんでした。システム内のコンテンツは ID によってテンプレート ノードと照合され、エージェントは作成した ID (「wgt-wrapper」) によってノードを参照します。私たちの HTML フォーマッタには、21 文字 (生成された nanoid の長さ) より短い ID を再生成するという古いルールがありました。したがって、エージェントの読み取り可能な ID は次のようになりました。

e は最初のフォーマット パスでランダムな文字列にサイレントに書き換えられ、その後の set_node_attributes("wgt-wrapper", …) は「ID を持つノードがありません」というエラーで失敗しました。ユーザーは最初の実際のセッションでこれをヒットしました。修正は 1 行 (ID が見つからない場合にのみ ID を生成する) でしたが、レッスンは一般化しています。エージェントが識別子を参照する場合、パイプライン内のすべてのパスで識別子をバイトごとに保持する必要があります。人間向けに書かれた正規化手順はあなたを裏切るでしょう。
2. HTML 属性は文字列で型指定されます。あなたの AST はおそらくそうではありません。私たちのパーサーは、読み取り時に「true」→ true、「1.5」→ 1.5 を強制します。人間にとっては便利ですが、一般的には危険です。「1.5」のようなバージョン文字列は数値 1.5 になり、たまたま数値に見えるテキスト属性は往復するたびに型が変換されます。私たちは文字列のみの属性のオプトアウト リストを維持しています。つまり、誰かがそのリストを思い出すまで、すべての新しい属性は潜在的なバグです。やり直すと、強制は属性ごとに行われ、推論ではなくノード構成で宣言されます。 (この記事の最後にあるリファレンス実装はまさにそれを行います。)
3. マークアップは交換形式であり、保存形式ではありません。実際、BLOB は解析された AST を JSON として保存します (屋根裏部屋は物を保管するのに最適な場所であることを覚えておいてください)。住所によって密封されたコンテナを取得するマシンは、ラベルが内側にあるかどうかを気にしません。著者だけがそうします。マークアップは、人間のマークアップ エディターとエージェントのツール I/O の 2 つの境界に存在します。誰かが実際に探し回っているのはこの 2 つの場所です。これは思っている以上に重要です。私たちのサーバーには DOMParser がないため、すべてのマークアップから AST への変換はそれらの境界でクライアント側で行われ、ダウンストリームのすべて (スロット解決、レンダリング、コンテンツ マッチング) は型指定された JSON で機能します。 「AST-as-HTML」とは、実際には AST のオーサリング言語としての HTML を意味し、1 つの保護された doo が付いています。

それらの間にr。その境界線を曖昧にすると、解析できない箇所で HTML を解析することになります。
もう 1 つの重要な分野は、ラウンドトリップ プロパティ テストです。 parse(build(tree)) は、ID、名前、属性を正​​確に保持する必要があります。上記のフォーマッタのバグは、後で作成した 5 行のテストによって発見されたでしょう。
深く型付けされたツリーまたは数値を多用するツリー。ノードのほとんどが float、enum、および相互参照である場合、HTML の stringly 属性が問題となります。これが機能するのは、ドキュメントのレイアウトがすでに HTML 形式になっているためです。
膨大な書類。フルツリーはツリーのサイズに合わせてスケールを書き換えます。私たちのものは有限です (テンプレートは数十のノードです)。 10,000 ノードのシーン グラフでは、結局のところ、粒度の高い API またはチャンク化された書き換えが必要になります。
モデルが完全に見ることができない木。ツリー全体の I/O は、ツリー全体がコンテキストに適合し、安全に表示できることを前提としています。
マルチライターの同時実行。 「すべてを置き換える」は、構造上、最後に書き込んだ方が有利です。セッション内で 1 人のユーザーと 1 人のアシスタントであれば問題ありません。リアルタイムのコラボレーションには間違っています。
モデルがすでに話している形式を選択してください。プロンプトを必要としない流暢さは、これまでに提供される中で最も安価な機能であり、人間の可読性と同じ性質です。手で読む人のためのデザインであり、モデルの利点です。

[切り捨てられた]

## Original Extract

LLMs have deep, pre-trained fluency in HTML. So when the AI assistant in our document platform reads and writes template layouts, we encode the tree as markup instead of JSON. Here

-->
HTML as a Native Data Format for LLMs | AST-as-HTML - LJ
hack Lightning Jar - Web Studio
Site Navigation
HTML as a Native Data Format for LLMs: Why We Encode Our Data in Markup Instead of JSON
We built a document platform where an AI assistant designs marketing documents (flyers, brochures, one-pagers) inside brand-approved rails, and humans finish them by clicking into the rendered page and typing. Getting the agent to author templates , the structural layouts those documents are built from, turned out to hinge on a single unfashionable decision:
We encode our templates as HTML, not JSON. And the agent's main editing tool is "rewrite the whole thing."
That inverts most of the current advice about building agents on structured data. It also turned out to be the cheaper, sturdier choice: fewer tokens burned, faster responses, and better data integrity on every edit. This post is about why.
A document template in our system is a tree: a document contains pages, pages contain blocks, blocks contain text atoms, image atoms, styled containers, and slots that reference reusable widgets. Every node carries attributes: CSS classes, length budgets for copy, type-scale choices, slot constraints. Themes paint the tree through CSS variables, so a template never hardcodes a color; it says bg-primary and the active brand theme decides what that means.
We wanted an assistant that could build and restructure these trees conversationally: "Design a full-width widget with a rounded content box on the left and a stat panel on the right." And we wanted its output to land in the same editor, with the same undo semantics and the same validation, as a human's edits.
The Obvious Design, and Why We Didn't Ship It
The textbook approach is to store the tree as JSON and give the model a granular tool API:
insertNode(parentId, type, index)
setAttribute(nodeId, key, value)
moveNode(nodeId, newParentId, index)
removeNode(nodeId) We've built agents like this. They work, but they under-perform in three predictable ways:
You're teaching a bespoke schema from scratch. Every node type, every attribute, every containment rule has to be spelled out in the prompt, and the model's only fluency is whatever your prompt bought. It has seen your JSON schema zero times in training.
Granular tools invite granular failure. Building a twelve-node layout takes a dozen round trips. Each call can reference a stale id, a wrong parent, an index that shifted two calls ago. The tree passes through eleven intermediate states, each a chance to strand the agent somewhere invalid, and each a state your renderer might have to survive.
The model can't "see" its work. With mutation-by-tool-call, the model's picture of the current tree is a mental reconstruction from its own call history. Drift is inevitable.
Our templates serialize to plain HTML with a small attribute grammar:
<div data-type="block" data-name="feature-callout" id="wgt-root">
<div data-type="container-atom" data-name="content-box" id="wgt-content"
data-container-classes="rounded-2xl flex-[2] bg-primary p-6">
<div data-type="text-atom" data-name="heading" id="wgt-heading"
data-text-element-tag="h2" data-text-style="heading-2"
data-max-length="60"></div>
</div>
</div> Every node is an element. data-type names the node kind, data-name is a stable human label, data-* carries attributes, and a parser converts this to and from the internal AST. This predates our agent work; it existed so templates could round-trip through a human-editable markup view.
When we built the template-authoring agent, we made its primary tool embarrassingly blunt:
set_template_markup(markup: string, summary: string)
// "Replace the ENTIRE template with new markup." That's it. The model reads the current markup (sent fresh with every request, so it always edits what the user actually sees), writes the complete new tree, and the client validates and applies it in one shot.
The result surprised us with how little prompting it needed. LLMs have deep, pre-trained fluency in HTML: nesting, attributes, class strings, ids. We didn't teach a format; we borrowed one the model already speaks natively. The prompt spends its budget on our semantics (what a widget-slot means, which theme tokens exist, how length budgets work) instead of on syntax. A grammar digest generated from the same config file the visual editor uses keeps the two from drifting.
And because each edit is a whole tree, there are no intermediate states. The edit is coherent or it's rejected, one parse at the boundary:
const parsed = parseMarkupToAst(markup);
if (parsed?.type !== "block") reject("widgets need a single block root");
editor.markupCurrent = format(build(parsed)); // same funnel as human edits Every agent edit flows through the exact pipeline every human edit uses. Same validation, same undo, same reactive preview. The agent isn't a privileged actor with its own write path; it's just another author.
Closed Containers in the Attic
Working with JSON is like hunting through closed storage containers in an attic. The labels are inside the lid; you have to open each container to learn what it is, and because JSON nests, you're opening containers inside containers inside containers. The containers only make sense if you brought the packing list: without knowing the key names and schema in advance, you can't even ask the right questions. And when you're done rummaging, the way out is a run of identical unlabeled lids ( }]}} ), none of which says what it closes. Deep in the attic, one misplaced lid and the whole stack is corrupt.
HTML labels the outside of every container:
<div data-type="text-atom" data-name="heading" data-max-length="60"> The node's type, name, and constraints are readable before you ever step inside, and the container closes with its name on it. Ask a model for an inventory of a JSON tree and it reconstructs ; ask for an inventory of an HTML tree and it reads the labels . That difference is most of why the whole-tree tool works: generating sixty nodes of labeled, self-closing HTML containers is something the model has done a billion times, and getting the lids right is easy when every lid says what it belongs to.
The irony is that we always planned to hand this tree to an LLM, but our early prototypes didn't yet; templates had to be authored and debugged by hand, so the problem we were solving first was human parsability. Labels on the outside made the tree easier for a person to read and write. That choice ended up revealing a hidden truth about LLM readability: legibility for the human author and fluency for the model are the same property.
(Hold that thought about the attic, though; it comes back. An attic is a perfectly good place to store things. You just don't want to work in one.)
The Payoff Is Practical, Not Just Aesthetic
Fewer tokens burned. The prompt no longer has to teach a schema the model already knows.
Fewer round trips. A model this fluent can author the whole tree in one pass instead of assembling it mutation by mutation.
More reliable transfer. A slipped tag fails loudly at the parser instead of silently corrupting everything after it.
Faster reads. The model scans labels instead of reconstructing structure.
Better integrity on every parse-and-rebuild cycle. A tree whose nodes announce themselves is a tree you can verify at a glance.
Full-tree rewrites are wasteful for one-attribute tweaks, and rewriting sixty nodes to change one class invites transcription drift in the other fifty-nine. So there's exactly one surgical tool:
set_node_attributes(nodeId: string, attributes: Record<string, value>) The tool cost almost nothing to build, and that's worth pausing on. In a web application, the machinery for editing HTML surgically already ships with the browser: getElementById , setAttribute , querySelector . These APIs are fast, battle-tested, and documented to a depth no in-house tree-mutation library will ever match. We didn't have to invent a node-addressing scheme or write an attribute patcher; the platform had already spent decades perfecting them.
Two tools total for structural editing. In practice the model picks correctly without guidance: big changes get a rewrite, tweaks get a patch. Compare that with the tool-catalog approach, where the model must choose among a dozen mutations and compose them correctly.
The same grammar then paid a second dividend: when we later let the document assistant mint brand-new widgets mid-conversation, it authored them in the same markup, validated by the same parser, with no new format to teach.
This is the part most posts skip. Three real bugs, all instructive:
1. Ids are sacred, and our formatter wasn't treating them that way. Content in our system is matched to template nodes by id, and the agent references nodes by the ids it authored ("wgt-wrapper"). Our HTML formatter had an old rule: regenerate any id shorter than 21 characters (the length of our generated nanoids). So the agent's readable ids were silently rewritten to random strings on the first format pass, and its follow-up set_node_attributes("wgt-wrapper", …) failed with "No node with id." A user hit this in the first real session. The fix was one line (only generate an id when one is missing ), but the lesson generalizes: if agents reference identifiers, every pass in your pipeline must preserve them byte-for-byte. Normalization steps written for humans will betray you.
2. HTML attributes are stringly typed; your AST probably isn't. Our parser coerces "true" → true and "1.5" → 1.5 on read. Useful for humans, hazardous in general: a version string like "1.5" becomes the number 1.5, and a text attribute that happens to look numeric gets type-bent on every round trip. We maintain an opt-out list of string-only attributes, which means every new attribute is a latent bug until someone remembers the list. If we started over, coercion would be per-attribute and declared in the node config, not inferred. (The reference implementation at the end of this article does exactly that.)
3. Markup is the interchange format, not the storage format. Our blobs actually store the parsed AST as JSON (the attic, remember, is a fine place to keep things). Machines fetching sealed containers by address don't care that the labels are inside; only authors do. Markup exists at exactly two boundaries: the human markup editor and the agent's tool I/O. Those are the two places where someone is actually rummaging. This matters more than it sounds: our servers have no DOMParser, so all markup→AST conversion happens client-side at those boundaries, and everything downstream (slot resolution, rendering, content matching) works on typed JSON. "AST-as-HTML" really means HTML as the authoring dialect of the AST , with one guarded door between them. Blur that line and you'll end up parsing HTML in places that can't.
One more discipline that earns its keep: round-trip property tests. parse(build(tree)) must preserve ids, names, and attributes exactly. The formatter bug above would have been caught by a five-line test we only wrote afterward.
Deeply typed or numeric-heavy trees. If your nodes are mostly floats, enums, and cross-references, HTML's stringly attributes fight you. This works because document layout is already HTML-shaped.
Huge documents. Full-tree rewrites scale with tree size. Ours are bounded (a template is a few dozen nodes); a 10,000-node scene graph would need the granular API after all, or chunked rewrites.
Trees the model shouldn't fully see. Whole-tree I/O assumes the whole tree fits in context and is safe to show.
Multi-writer concurrency. "Replace everything" is last-write-wins by construction. Fine for one user plus one assistant in a session; wrong for real-time collaboration.
Choose formats the model already speaks. Fluency you don't have to prompt for is the cheapest capability you will ever ship, and it's the same property as human legibility: design for the person reading by hand and the model benefits for

[truncated]
