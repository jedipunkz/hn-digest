---
source: "https://www.evolvinglab.ai/blog/research-world-model"
hn_url: "https://news.ycombinator.com/item?id=48879047"
title: "AI Should Build Its Own Research World Model"
article_title: "AI Should Build Its Own Research World Model · Opus 4.8 on ARC-AGI-3"
author: "vinhnx"
captured_at: "2026-07-12T08:00:15Z"
capture_tool: "hn-digest"
hn_id: 48879047
score: 1
comments: 0
posted_at: "2026-07-12T07:11:59Z"
tags:
  - hacker-news
  - translated
---

# AI Should Build Its Own Research World Model

- HN: [48879047](https://news.ycombinator.com/item?id=48879047)
- Source: [www.evolvinglab.ai](https://www.evolvinglab.ai/blog/research-world-model)
- Score: 1
- Comments: 0
- Posted: 2026-07-12T07:11:59Z

## Translation

タイトル: AI は独自の研究世界モデルを構築する必要がある
記事のタイトル: AI は独自の研究世界モデルを構築する必要がある · ARC-AGI-3 の Opus 4.8
説明: 未知の ARC-AGI-3 ゲームにドロップされた Opus 4.8 は、クエリ可能な研究世界モデルを作成し、7 つのレベルすべてをクリアしました。答えを与えられたクローンはまったく進歩しませんでした。

記事本文:
ワールドモデル · ワールドモデル · —
ディレクトリ
内容
§1 なぜルールのないゲームなのか §1 なぜルールのないゲームなのか
§2 命名とファイリング §2 命名とファイリング
§3 ボタンは操作です §3 ボタンは操作です
§4 間違った法律、自己論駁 §4 間違った法律、自己論駁
§5 ワールド モデルのクエリ §5 ワールド モデルのクエリ
§51/2 アーカイブの内部 §51/2 アーカイブの内部
§6 回答キーを使用してクローンを作成する §6 回答キーを使用してクローンを作成する
§7 実験では何が検証されましたか? §7 それが示すもの
§8 引き継ぐ §8 引き継ぐ
§9 謝辞 §9 謝辞
アラ
ブログ
ハブ
実験の記録・2026年6月
現地レポート · 2026 年 6 月
AI は独自の研究世界モデルを構築すべき Opus 4.8 は ARC-AGI の最終レベルを一発でクリアしました。
スキルリポジトリ
🤗 軌跡データセット
英語/中国語
TL;DR AI はまったく未知の世界を本当に理解できるのでしょうか?システムが科学者のように新しいパズルを解けるようにするための鍵は何ですか?現在のパラダイムでは、大規模モデル推論の重みが完全にロックされており、文脈が壊れた瞬間に試行錯誤で蓄積された経験が瞬時にクリアされてしまうため、人間のように「探究－反省」のサイクルで自然に「科学的直観」を養うことはできません。このボトルネックを打破するために、私たちのソリューションは、AI のコグニティブ アーキテクチャを明示的に外部化し、探査プロセス中に試行錯誤しながらエラーを記録できるようにし、発見された基礎的なメカニズムをコンテキスト全体でクエリできる「科学的世界モデル」に動的に沈殿させることです。この方法をテストするために、Opus 4.8 のクロード コードをこれまで見たことのないパズルの世界 (ARC-AGI-3 の ls20) に投げ込み、それに合格するように依頼しました。ご存知のとおり、このようなパズルを裸で探索するための最先端の大規模言語モデルのベンチマーク合格率は、わずか約 1% です。しかし、ワールド モデルの導入後、システムは基礎となるすべてのメカニズムを独立して解読しただけでなく、最後で最も難しいレベルに挑戦する際、エージェントはワールド モデルの支援を受けて 1 回限りのパスを達成しました。これ

これは、トレーニング不要で自己進化するこの方法を通じて、AI エージェントが未知の世界の物理法則を真に理解し、独自の世界モデルを構築できることを証明しています。
TL;DR AI はまったく未知の世界を本当に理解できるのでしょうか?システムが科学者のように新たな問題を解決できるようにするための本当の鍵は何でしょうか?現在のパラダイムでは、モデルの重みは推論中にロックされます。コンテキスト ウィンドウが切れた瞬間、試行錯誤によって得られた経験はすべて消えてしまいます。AI は人間のように探索と考察のループを通じて「科学的直観」を自然に育てることはできません。このボトルネックを打破するために、私たちのアプローチは、AI に外部認知アーキテクチャを明示的に装備し、AI が探索中に記録し、発見されたメカニズムをクロスコンテキストでクエリ可能な「研究世界モデル」に動的に結晶化できるようにすることです。これをテストするために、Claude Code + Opus 4.8 エージェントをまったく見えないパズルの世界 (ARC-AGI-3 の ls20) にドロップし、ゲームをクリアするというタスクを課しました。文脈のために言うと、このようなパズルにおけるフロンティア LLM のベア探索のベースライン合格率は ~1% です。しかし、世界モデルによって強化されたため、すべての基礎となるメカニズムが自律的に体系的に解読されました。最後の最も難しいレベルに直面したとき、エージェントはモデルの支援により一発勝利を達成しました。このトレーニング不要で自己進化するアプローチにより、AI エージェントが未知の世界の物理法則を理解し、独自の世界モデルを構築できることを示します。
過去数年間の研究の中で、私は科学研究の本質が実際には未知の世界での継続的な探索のフィードバック ループであることを発見しました。つまり、探索し、仮説を提案し、行動し、観察し、反映することです。人間の科学はこのサイクルに基づいて直感を反復し、その後、これらの直感がゆっくりと論文に定着します。未知の世界を探検するとき、直感は

それは私たちの脳の中で自然に成長します。
長年の研究の中で、科学的探究の本質は、未知の世界の中での継続的なフィードバック ループ、つまり探索、仮説、行動、観察、考察であることに気づきました。人間の科学はこのループに依存して直感を反復し、それがゆっくりと論文に結晶化します。未知のものを探求すると、直感が自然に心の中で育まれます。
しかし、AIが未知の世界を探索するとき、抽象的に考える能力は自然に身につくわけではありません。では、AI の「直感的なシステム」を明示的に構築できるでしょうか? ——私たちはいつも、直感はやや形而上学的なものだと言いますが、理論的にはそれを記録することができます。これは強化学習のパラダイムに関するものです。RL では、確かにモデルの重みを直接更新できますが、これは解釈可能性に欠ける非常に暗黙的な知識圧縮にすぎず、簡単に抽出してデバッグすることはできません。したがって、これらすべてが私たちに考えるように促します。AI が継続的に記録、追跡し、いつでもクエリできるように、AI の外部認知アーキテクチャを構築して、AI が時間の経過とともに複合知識を生成できるようにすることはできるでしょうか?私はこのシステムを研究世界モデルと呼んでいます。
しかし、AIが未知の世界を探索する場合、このような抽象的な思考能力は自然に育つことはできません。では、AI の「直観システム」を明示的に構築できるでしょうか?私たちは直感について神秘的なものとしてよく言いますが、理論的にはそれを記録することができます。これは、強化学習パラダイムにつながります。RL では、確かにモデルの重みを直接更新できますが、これは知識圧縮の高度に暗黙的な形式にすぎません。不思議: 建てられるでしょうか

継続的に記録、追跡、クエリできる AI 用の外部認知アーキテクチャで、時間の経過とともに AI が知識を複合化できるようにするものでしょうか?私はこのアーキテクチャを研究世界モデルと呼んでいます。
レベル1
レベル1
結晶化する
結晶化する
7つの法律
請求項7
最初の夜 · 理解するには 24 ステップ · 最低 13 ステップ
第一夜・24で解決・ベスト13
レベル 2 ～ 4
レベル 2 ～ 4
レベル5
レベル5
結晶化する
結晶化する
検索
取得する
14の法律
請求項14
レベルを完了するのに 56 回の移動 (L4 では < 60) · 最初の紙の派生で壁を打ち破る
56 年にクリア (< L4 の 60) · 紙に基づいて作成
レベル6
レベル6
レベル7
レベル7
63ステップ・ワンパス
63ステップ・ワンショット
探索×2
リトリーブ×2
アラ
ワールドモデル
ワールドモデル
15の法則・7レシピ・75ノード
クレーム 15 · レシピ 7 · ノード 75
このレベルには新しいルールはありません: まず質問してから行ってください
新たなクレームはゼロです。まず尋ねてから歩きましょう
エージェントはレベルごとに交代します。同じ世界モデルが複利を継続する
エージェントはレベルごとに交換されます。同じ世界モデルが複雑化し続ける
赤いペン = 研究マネージャー: 学んだことを世界モデルに結晶化します (クローズ時にクレーム/トレースを書き込みます) · 青いミラー = 研究先見: 検索と予測 (読み取り専用)
赤い鉛筆 = 研究マネージャー: 教訓を世界モデルに具体化します (クレーム/トレース、終了時) · 青い拡大鏡 = 研究先見: 取得と予測 (読み取り専用)
世界モデルがどのように構築されるか。下の行は同じ ARA アーカイブで、すべてのレベルにわたって成長を続けています。最初の夜の後は 7 つの法則、第 5 レベル以降は 14 つ、そして最も厚いものでは、15 の法則、7 つのレシピ、および 75 の研究ツリー ノードです。上の行の各レベルの後に新しいエージェントが置き換えられますが、同じ世界モデルは引き続き削除されます。各エージェントは、学んだことを赤いペンでファイルに書き込み (研究マネージャー、ファイルが閉じているときのみ書き込みます)、プレイ中に最初にファイルを読み取り、青い虫眼鏡を使用して詰まっているかどうかを確認します (研究先見、読み取り専用)。 7番目まで

これには、「まず尋ねてから行き、63 段階でレベルをクリアする」という新しい法則は含まれていません。
世界モデルがどのように構築されるか。一番下の行は同じ ARA アーカイブで、レベルごとに成長します。最初の夜の後は 7 つの法則、レベル 5 以降は 14 つ、そして最も厚いところでは 15 の法則、7 つのレシピ、75 の研究ツリー ノードです。エージェントはレベルごとに入れ替わりますが、同じ世界モデルが複雑化し続けます。各現職者は学んだことを赤鉛筆でアーカイブに書き込み (研究マネージャー、終了の瞬間にのみ書き込みます)、到着したら最初に読み取り、行き詰まったときは青い拡大鏡を通して質問します (研究先見、読み取り専用)。レベル 7 までに、新しい法律はまったく書かれていませんでした。まず尋ねてから歩き、63 歩、1 回で歩きます。
次の 2 つのビデオは比較を示しています。この世界モデルのセットを独自に構築した後、AI は最後のレベルに挑戦するときにわずか 63 ステップで探索を終了し、無事にレベルをクリアしました。逆に、そのような世界モデルがなければ、前レベルの「クリア記録」を放り込んでも、5,000歩以上手探りで進歩せずに行き止まりになってしまいます。これは、AI が未知の行き止まりを突破するために世界モデルが重要であることを鮮やかに証明しています。この一連の書かれたファイルに正確に何が含まれているかが、この記事で説明する内容です。
以下の 2 つの記録は、その対照を示しています。AI が探索中に独自の世界モデルを構築することで、最終レベルに取り組むと、探索を終了し、わずか 63 ステップでゲームをクリアしました。逆に、そのような世界モデルがなければ、以前のレベルの「明確な記録」を渡されても、5000 ステップ以上、まったく進歩せずに行き止まりをさまよっています。これは、AI が未知の行き詰まりを打開するための世界モデルの重要性を鮮やかに証明しています。その書かれたアーカイブには正確に何があったのか、それがこの投稿の内容です。
§1 AI をルールのないゲームに投入する: 科学研究には、権威に盲目的に依存するのではなく、明示的な世界モデルが必要です。

重い
§1 ルールのないゲームへ: 研究には重みだけでなく明示的な世界モデルが必要
この実験を行うには、AI が独自に基礎となる物理法則を探索し理解するための、まったく未知の環境を見つける必要があります。探索、要約、考察、仮説策定のサイクルを通じて、世界の明確な理解がゆっくりと形成され、最終的にはそのレベルを超えます。この環境は最先端の科学研究分野となる可能性がありますが、ここでは、継続的な抽象化と反映機能における AI の進歩を調査するための実験オブジェクトとして ARC-AGI のインタラクティブ パズルを使用することを選択しました。
この実験を実行するには、AI が独自に基礎となる物理学を探索して理解できる、まったく未知の環境が必要でした。探索、要約、考察、仮説のループを通じて、ゆっくりと世界を明確に理解し、最終的にはゲームをクリアします。この環境は科学研究のフロンティアとなる可能性がありますが、ここでは、継続的な抽象化と反映機能における AI の進歩を調査するための実験対象として、ARC-AGI のインタラクティブなパズルを選択しました。
これはほんの一例です。それは、ゲームの世界、タンパク質、未解決の方程式、まだ誰もマッピングしていない科学のフロンティアである可能性があります。私たちが本当に知りたいのは、AI がそのレベルをクリアできるかどうかではなく、誰も行ったことのない場所に置かれた場合に科学者のように自力で脱出できるかどうかです。
そして、これはほんの一例にすぎません。世界はゲームかもしれない。それは、タンパク質の広がり、未解決の方程式、誰も描いたことのない科学の最前線である可能性もあります。私たちが実際に知りたいのは、AI がゲームをクリアできるかどうかではありません。それは、誰も行ったことのない場所に落とされたときに、科学者と同じように歩き出すことができるかどうかです。
研究は圧縮された状態で行われるべきではありません

重みと直感。
§2 未知のものに名前を付けてファイリングする: 表現の発明は理解への第一歩です
クロード コードにプロンプトをスローすると、実行全体が開始されます。
1 つのプロンプトを Claude Code に投げ込み、そこから全体の実行が始まりました。
プロセスは次のとおりです。メイン エージェントはゲーム ハーネスをセットアップし、プレイするために各レベルにプレイ サブエージェントを送信し、手配、アカウンティング、同期のみを行います。ゲームの推論はすべてサブエージェント内で行われます。
フロー: メイン エージェントはゲーム ハーネスを生成し、レベルごとに 1 つのプレイ サブエージェントをディスパッチし、オーケストレーション、台帳、および自身の同期のみを保持します。ゲームの推論はすべてサブエージェント内で行われます。
サブエージェントは、ワールド モデルに対処するために 2 つのスキルに依存します。読み取りには /research-foresight を使用します。ゲームに入る前に、まず ARA (そのワールド モデル) を読み取る必要があります。百歩程度で行き詰まったときに呼び出して、ワールドモデルに直接質問してください。読み取りのみが行われ、書き込みは行われません。
サブエージェントは 2 つのスキルを通じてワールド モデルと通信します。読み取りには /research-foresight が使用されます。サブエージェントは、行動する前に ARA (そのワールド モデル) を読み取る必要があり、約 100 回スタックしたアクションの後、スキルを呼び出してワールド モデルを直接クエリします (読み取り専用)。
書き込みには /research-manager を使用します。学習内容は終了時間中にのみファイルに結晶化されます。プレーヤー エージェント自体はファイルを手書きすることを許可されていません。
書き込みは /research-manager を使用します。レッスンは、閉鎖の瞬間に、それを通じてのみアーカイブに結晶化します。プレイ中のエージェントがファイルを手書きすることはありません。
さて、目を開けたときに何が見えたかを見てください。
あなただったら、最初にどのキーを押しますか?
4つともクリックされました。一部の方向には移動でき、一部の方向には移動できませんが、下部の水平バーはまだ 1 スペース不足です。壁にぶつかるのはタダではありません。この世界では、間違いを犯した場合でも料金を請求され、料金表さえ与えられません。それから、灰色の帽子をかぶった紫色の体の四角を動かしたり、目立たない黒い点を踏んだり、左下隅の絵をいじったりしました。

ケースが少し飛びます。ブロックを上のボックスに押し込むと負けです。パターンが一致するまで待ってもう一度押すと、画面全体が再描画されます (第 2 レベル)。
それは4つすべてを押しました。一部の方向が移動しました。他の人は何かに当たってもびくともしませんでしたが、やはり下のバーは短くなりました。壁にぶつかるのはタダではありません。そして画面全体が再描画されました。レベル2。
最初のセッションが終了する前（第一夜と呼んでもいいかもしれません）、博物学者が無人島に降り立ったときに最初にすること、それは見たことのないものに名前を付けることです。移動するブロックはブロックと呼ばれます。踏むと模様がジャンプする小さな点をスイッチといいます。ターゲットパターンが入った上のボックスはキーボックスと呼ばれます。そして、目的が不明瞭な地面上の珍しい斑点は、何気なく斑点と呼ばれています。ほとんど想像力に富んだものではありませんが、うまくいきます。名前を選択すると、ドキュメントへのエントリの書き込みが開始されました。
最初のセッション（第一夜と呼ぶ）を終える前に、島に漂着した博物学者が最初にすることを行った。それは、見たことのないものに名前を付けるというものだった。動くマス目がブロックになった・・・踏むとパターンがジャンプする小さなドット、スイッチ。ターゲットパターンを保持する上部のフレーム、キーボックス。明確な目的を持たない稀な床の斑点を、無作法に斑点と呼びました。
## スイッチ
稀に色 0/1 の床の斑点がブロックに足を踏み入れてロックを切り替え/循環させます
パネルに表示されるパターン。消費されない（ブロックが移動した後も存在する）
オフ）。レベルごとのトランス

形式が異なります: L1 のスイッチ ( R6C3 ) は、
中央の列 (2 つの州)。 L2 のスイッチ (R9C9) は 4 ステート サイクル (C09、C10) です。
スイッチ: 珍しい地上スポット。踏むとパネルのロックパターンが切り替わります。消費されません。各レベルの変換ルールは異なります。最初のレベル (行 6、列 3) のスイッチは中央の行 (2 つの状態) のみを反転します。 2 番目のレベルのスイッチは 4 状態サイクルです。
初めてconcepts.mdに目を通したときも、R6C3やcolor-0/1などで行き詰まってしまいました。急いで説明しないでください。この「読みにくさ」自体が証拠です。最初の行から始まるこれらのエントリは、人間が読むことを目的として書かれていません。
初めてこのconcepts.mdを開いたとき、R6C3とcolor-0/1にも引っかかりました。急いで説明しないようにしましょう。読みにくいこと自体が証拠です。最初の行から、これらのエントリは決して人間向けに書かれたものではありません。
最初の夜の請求: 最初にレベルを把握するために 24 手が取られました。その後、レベルを完了するまでの最短時間はわずか 13 手でした。死者数0人。完成したとき、その世界モデルにはすでに、あえて書き留めた 7 つのルールが含まれていました (ファイルには、これらはクレームと呼ばれ、1 つずつ番号が付けられています)。その中には、将来その命を救うことになる C07 も含まれています。つまり、封印されたキー ボックスは、ロックが満たされた場合にのみ開けられます。
最初の夜の請求書: 初めてレベルを調整するための 24 ステップ。その後は最短クリアまで13ステップ。死者数0人。会期終了までに、同社が作成に取り組むつもりだった7つの法律がすでに世界モデル（同社はこれらの主張を呼んでおり、それぞれに番号が付けられ、提出されている）に組み込まれており、その中には後に命を救うことになるC07が含まれている。施錠が完了した場合にのみ、壁に閉ざされたキーボックスが開く。
ネーミングとは、混沌を考えられる部分に切り分けることです。
人類の知恵が蓄積され始めました。ある世代が突然賢くなったのではなく、私たちが言語を発明したのです。初めて経験が言語を生み出した頭から離れることができるようになったのです。この小さな世界で繰り返すこの一歩

: まず言葉があって、それから伝わる理解があります。
人類の知恵が蓄積され始めたのは、一部の世代が突然賢くなったからではなく、私たちが言語を発明したからです。初めて、経験が言語を生み出した頭から離れることができました。それがこの小さな世界で再現されたステップです。まず言葉、そして次に伝えることができる理解を。
しかし、ここまでのところ、書いている内容は旅行日記と変わりません。見て、書き留めてください。変化は第 5 レベルで起こります。それは言葉を発明し、自動的に計算を開始します。
しかし、この時点までは、まだ旅行日記にすぎませんでした。つまり、記録です。それは変わります

[切り捨てられた]

## Original Extract

Dropped into an unseen ARC-AGI-3 game, Opus 4.8 wrote a queryable research world model and cleared all 7 levels; a clone given the answers made zero progress.

世界模型 · world model · —
目录
Contents
§1 为什么是没规则的游戏 §1 Why a game with no rules
§2 起名字，写词条 §2 Naming and filing
§3 按钮其实是运算 §3 Buttons are operations
§4 写错，再自己推翻 §4 Wrong laws, self-refuted
§5 卡关时问世界模型 §5 Querying the world model
§5½ 档案里有什么 §5½ Inside the archive
§6 给克隆体标准答案 §6 Clone with the answer key
§7 实验验证了什么 §7 What it shows
§8 传下去 §8 Pass it on
§9 致谢 §9 Acknowledgments
ARA
Blog
Hub
一次实验的实录 · 2026 年 6 月
A field report · June 2026
AI Should Build Its Own Research World Model Opus 4.8 Cleared ARC-AGI’s Final Level in One Shot.
Skill Repository
🤗 Trajectory Dataset
EN / CN
TL;DR AI 到底能否真正理解一个完全未知的世界？什么才是让系统像科学家一样破解全新难题的关键？在当前范式下，大模型推理时权重完全锁死，上下文一断，试错积累的经验便瞬间清零——它无法像人类那样在“探索-反思”的循环中自然长出“科研直觉”。为了打破这一瓶颈，我们的解法是： 为 AI 显式外置一套认知架构，让它在探索的过程中边试错边记录，将发现的底层机制动态沉淀为一份可跨上下文查询的“科研世界模型” 。为了验证这个方法，我们将 Claude Code 与 Opus 4.8 的组合丢入一个它从未见过的谜题世界（ARC-AGI-3 的 ls20）中要求其通关。要知道，前沿大语言模型裸探索这类谜题的基准通关率仅约 1%。然而，在引入世界模型后，系统不仅自主破解了所有底层机制，在挑战最后也是最难的关卡时，智能体在世界模型的辅助下实现了一次性通关。这证明了通过这种无需微调（training-free）且自我演化的方式，AI 智能体能够真正理解一个未知世界的物理规律，并构建出属于自己的世界模型。
TL;DR Can AI truly understand a completely unknown world? What is the real key to enabling a system to crack novel problems like a scientist? Under the current paradigm, model weights are locked during inference; the moment a context window expires, any experience gained through trial and error vanishes—AI cannot naturally grow "scientific intuition" through a loop of exploration and reflection as humans do. To break this bottleneck, our approach is to explicitly equip the AI with an external cognitive architecture, allowing it to record as it explores and dynamically crystallize discovered mechanics into a cross-context, queryable "research world model." To test this, we dropped a Claude Code + Opus 4.8 agent into a completely unseen puzzle world (ARC-AGI-3's ls20) and tasked it with clearing the game. For context, the baseline pass rate for frontier LLMs' bare exploration in such puzzles is ~1%. Yet, empowered by the world model, the system autonomously decoded all underlying mechanics. When faced with the final and hardest level, the agent achieved a one-shot victory with the model's assistance. We show that with this training-free and self-evolving approach, an AI agent is able to understand the physical rules of an unknown world and build its own world model.
做研究这几年，我发现科研的本质其实就是在一个未知的世界里持续探索的反馈循环（feedback loop）：探索、提出假设、行动、观察、反思。人类的科学就是靠这个循环迭代出直觉，然后这些直觉慢慢沉淀在论文里。当我们在探索未知世界时，直觉会自然而然地长在我们的脑子里。
In my years of doing research, I’ve realized that the essence of scientific inquiry is a continuous feedback loop inside an unknown world: explore, hypothesize, act, observe, and reflect. Human science relies on this loop to iterate out intuition, which then slowly crystallizes into papers. When we explore the unknown, intuition naturally grows in our minds.
但是，当 AI 在探索未知的世界时，它没办法自然而然地长出这种抽象思维的能力。那我们是不是可以显式地为 AI 搭建一套“直觉系统”呢？——我们总说直觉是个有些玄学的东西，但理论上它是可以被记录下来的。这就要说到强化学习的范式：在 RL 里，我们确实可以直接更新模型的权重，但这不过是一种非常隐式的知识压缩，缺乏可解释性，也没办法很方便地被提取和调试。所以这一切促使我们思考：能不能给 AI 搭建一个外置的认知架构（cognitive architecture），让它能够持续地记录、追踪并被随时查询，让 AI 能够随着时间产生知识复利（compound knowledge over time）？我把这套系统称为 科研世界模型 （research world model）。
However, when AI explores an unknown world, it cannot naturally grow this kind of abstract thinking ability. Could we, then, explicitly build an "intuition system" for AI? We often speak of intuition as something mystical, but theoretically, it can be recorded. This brings us to the reinforcement learning paradigm: in RL, we can indeed directly update the model's weights, but this is merely a highly implicit form of knowledge compression. It lacks interpretability and cannot be easily extracted or debugged. All of this prompted us to wonder: could we build an external cognitive architecture for AI that can continuously record, track, and be queried, allowing the AI to compound knowledge over time? I call this architecture a research world model .
第 1 关
Level 1
结晶
crystallize
7 条定律
claims 7
第一夜 · 24 步摸清 · 最短 13 步
first night · solved in 24 · best 13
第 2–4 关
Level 2–4
第 5 关
Level 5
结晶
crystallize
检索
retrieve
14 条定律
claims 14
56 步通关（< L4 的 60）· 第一次纸上推导赢过撞墙
cleared in 56 (< L4’s 60) · derived on paper
第 6 关
Level 6
第 7 关
Level 7
63 步 · 一次通关
63 steps · one shot
检索 ×2
retrieve ×2
ARA
世界模型
world model
15 条定律 · 7 份配方 · 75 个节点
claims 15 · recipes 7 · nodes 75
本关零新增定律：先问，再走
zero new claims: ask first, then walk
每关后 agent 都会换人；同一个世界模型持续复利
agent gets swapped after each level; the same world model keeps compounding
红笔 = research-manager：把学到的结晶进世界模型（闭合时写入 claims/trace） · 蓝镜 = research-foresight：检索与预测（只读）
red pencil = research-manager: crystallize lessons into the world model (claims/trace, at closure) · blue magnifier = research-foresight: retrieve & predict (read-only)
世界模型是怎么被建出来的。 下排是同一本 ARA 档案，跨越所有关卡持续生长：第一夜后 7 条定律，第五关后 14 条，最厚时 15 条定律、7 份配方、75 个研究树节点。上排每关后都会换一个新的 agent，但同一个世界模型被继续带走：每一任把学到的用红笔写进档案（research-manager，只在闭合时刻动笔），上场先读，卡住就用蓝色放大镜去问（research-foresight，只读）。到第七关，它一条新定律都没写：先问，再走，63 步一次通关。
How the world model gets built. The bottom row is the same ARA archive, growing across every level: 7 laws after the first night, 14 after level 5, and at its thickest 15 laws, 7 recipes, 75 research-tree nodes. The agent gets swapped after each level, but the same world model keeps compounding: each incumbent writes what it learned into the archive with the red pencil (research-manager, which only writes at moments of closure), reads first on arrival, and asks through the blue magnifier when stuck (research-foresight, read-only). By level 7 it wrote no new law at all: ask first, then walk — 63 steps, one shot.
下面的两段录像展示了对比：在自主构建了这套世界模型后，AI 在挑战最后一关时仅仅用了 63 步就结束了探索，成功通关。相反，如果没有这样一个世界模型，哪怕只是扔给它前面关卡的“通关记录”，它自己摸索了 5000 多步却依然卡在死胡同里毫无进展。这生动地证明了世界模型对于 AI 突破未知死局的重要性。这套写下的档案里到底有什么，这篇文章就讲这个。
The two recordings below show the contrast: by building its own world model as it explored, when the AI tackled the final level, it finished the exploration and cleared the game in just 63 steps. Conversely, without such a world model, even if handed the "clear records" of earlier levels, it wandered in a dead end for over 5000 steps with zero progress. This vividly proves the importance of a world model for AI to break through unknown deadlocks. What exactly was in that written archive: that's what this post is about.
§1 把 AI 扔进无规则游戏：科研需要显式的世界模型，而非盲目依靠权重
§1 Into a Game With No Rules: Research Requires Explicit World Models, Not Just Weights
为了做这个实验，我们需要为 AI 找一个完全未知的环境，让它可以自己去探索并理解底层的物理规律。通过探索、总结、反思和提出假设的循环，慢慢沉淀出对这个世界的显式理解，最终通关。这个环境可以是一片前沿的科研领域，但在这里，我们选择用 ARC-AGI 里的一个互动谜题作为实验对象，来探索 AI 在持续的抽象与反思能力上的进展。
To run this experiment, we needed a completely unknown environment for the AI, where it could explore and understand the underlying physics on its own. Through the loop of exploration, summarization, reflection, and hypothesizing, it would slowly precipitate an explicit understanding of the world to eventually clear the game. This environment could be a frontier of scientific research, but here, we chose an interactive puzzle from ARC-AGI as our experimental subject to explore AI's progress in continuous abstraction and reflection capabilities.
而这只是一个例子。它可以是一个游戏世界，也可以是一段蛋白质、一个没被解过的方程、一片还没有人绘制的科学前沿。我们真正想知道的，不是 AI 能不能通关，而是把它放到一个谁都没去过的地方时，它能不能像科学家那样，自己走出来。
And this is only one example. The world could be a game; it could just as well be a stretch of protein, an unsolved equation, a patch of scientific frontier nobody has charted. What we actually want to know isn't whether an AI can clear a game; it's whether, dropped somewhere no one has ever been, it can walk out the way a scientist would.
Research shouldn't live compressed in weights and intuition.
§2 Naming and Filing the Unknown: Inventing Representations is the First Step to Understanding
我们把一份 prompt 扔进 Claude Code，整个 run 就跑起来了。
We tossed one prompt into Claude Code, and the whole run took off from there.
流程是：main agent 起好游戏 harness，每关派一个 play-subagent 去打，自己只做编排、记账、同步。所有游戏推理都发生在子代理里。
The flow: the main agent spawns the game harness and dispatches one play-subagent per level, keeping only orchestration, the ledger, and syncing for itself. All game reasoning happens inside the subagents.
子代理靠两个 skill 跟世界模型打交道。读用 /research-foresight ：上场必须先读 ARA（它的世界模型），卡住约百步就调用它，直接向世界模型提问，只读不写。
The subagents talk to the world model through two skills. Reads use /research-foresight : a subagent must read the ARA (its world model) before acting, and after roughly a hundred stuck actions calls the skill to query the world model directly — read-only.
写用 /research-manager ：所学只在闭合时刻经它结晶进档案，玩家 agent 自己不许手写文件。
Writes use /research-manager : lessons crystallize into the archive only through it, at moments of closure; the playing agent never hand-writes a file.
Now, look at what it saw when it opened its eyes.
If it were you, which key would you press first?
它四个都按了。有的方向走得动，有的撞上去纹丝不动，但底部的横条照样短一格。 撞墙不免费。 这个世界连犯错都收钱，还不给你价目表。接下来它就是瞎摆弄：把那个灰帽紫身的方块挪来挪去，踩过一个不起眼的黑点，左下角框里的图案跳了一下；把方块推进上面那个框，输了；等图案对上再推，整个屏幕重画，第二关。
It pressed all four. Some directions moved; others hit something and didn't budge, but the bar at the bottom got shorter all the same. Hitting a wall isn't free. This world charges for mistakes and doesn't post the prices. What followed was basically fiddling: it shoved the grey-capped purple block around, stepped across an inconspicuous black dot, and the pattern in the bottom-left frame jumped; pushed the block into the frame up top, lost; waited for the pattern to match and pushed again, and the whole screen redrew. Level 2.
第一个 session 收工前（不妨叫它第一夜），它做了博物学家落地荒岛会做的第一件事： 给从没见过的东西起名字。 会动的方块叫 block ；踩了会让图案跳动的小点叫 switch ；上方那个装着目标图案的框叫 key-box ；地上那种说不清用途的稀有斑点，它随手叫作 speck 。谈不上有想象力，但都好用。名字起好，它开始往一份文件里写词条。
Before closing out its first session (call it the first night), it did the first thing a naturalist does after washing up on an island: it named the things it had never seen. The square that moves became block ; the little dot that makes the pattern jump when stepped on, switch ; the frame up top holding the target pattern, key-box ; the rare floor spots with no obvious purpose it offhandedly called speck . Not imaginative names, but they all worked. Once they were settled, it started writing entries into a file.
## switch
A rare color-0/1 floor speck the block steps onto to toggle/cycle the lock
pattern shown in the panel. Not consumed (still present after the block moves
off). The per-level transform differs: L1's switch ( R6C3 ) toggles only the
middle row (2 states); L2's switch ( R9C9 ) is a 4-state cycle (C09, C10).
开关：一种稀有的地面斑点，方块踩上去会切换面板上的锁图案。不会被消耗。每一关的变换规则不同：第一关的开关（第 6 行第 3 列）只翻转中间一行（两个状态）；第二关的开关是四状态循环。
我第一次翻这份 concepts.md 的时候，也在 R6C3 、 color-0/1 这些地方卡了一下。先不急着解释。这种"读不懂"本身就是证据：这些词条从第一行起，就不是写给人读的。
The first time I opened this concepts.md I got snagged on R6C3 and color-0/1 too. Let's not rush to explain. The unreadability is itself evidence: from their first line, these entries were never written for humans.
第一夜的账单：首次摸清这一关花了 24 步；事后最短通关只要 13 步；死亡 0 次。收工时，它的世界模型里已经躺着七条它敢写下来的规律（在它的档案里，这叫 claim，逐条编号），包括那条日后救它命的 C07： 锁满足时，封死的钥匙盒才会打开。
The first night's bill: 24 steps to work the level out the first time; 13 steps for the shortest clear afterwards; 0 deaths. By close of session, seven laws it was willing to commit to writing already lay in its world model (it calls these claims, each numbered and filed), including C07, the one that would later save its life: only when the lock is satisfied does the walled-shut key-box open.
Naming is cutting chaos into parts you can think with.
人类的智慧开始累积，不是哪一代人突然变聪明了，而是我们发明了语言：经验第一次可以离开产生它的那颗脑袋。它在这个小世界里重演的就是这一步：先有词，才有可以传下去的理解。
Human wisdom started to pile up not because some generation suddenly got smarter, but because we invented language: for the first time, experience could leave the head that produced it. That's the step it re-enacted in this little world: first the words, then an understanding you can pass on.
不过到这里为止，它写的东西跟一本旅行日记还没什么两样：看见，记下。变化出现在第五关：它发明的词，开始自己做数学。
Up to this point, though, what it wrote was still nothing more than a travel diary: see, record. That changes

[truncated]
