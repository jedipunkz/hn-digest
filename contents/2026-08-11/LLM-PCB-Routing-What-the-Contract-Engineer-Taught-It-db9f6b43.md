---
source: "https://edgelog.dev/blog/llm-pcb-routing-human-taught/"
hn_url: "https://news.ycombinator.com/item?id=49254735"
title: "LLM PCB Routing: What the Contract Engineer Taught It"
article_title: "LLM PCB Routing: What the Contract Engineer Taught It | EdgeLog"
author: "0xecro1"
captured_at: "2026-08-11T08:03:32Z"
capture_tool: "hn-digest"
hn_id: 49254735
score: 1
comments: 0
posted_at: "2026-08-11T07:55:39Z"
tags:
  - hacker-news
  - translated
---

# LLM PCB Routing: What the Contract Engineer Taught It

- HN: [49254735](https://news.ycombinator.com/item?id=49254735)
- Source: [edgelog.dev](https://edgelog.dev/blog/llm-pcb-routing-human-taught/)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T07:55:39Z

## Translation

タイトル: LLM PCB 配線: 契約エンジニアが教えてくれたこと
記事のタイトル: LLM PCB 配線: 契約エンジニアが教えてくれたこと |エッジログ
説明: 私のエージェントの 16 個の PCB 配線テクニックのうち 7 個

記事本文:
LLM PCB 配線: 契約エンジニアが教えてくれたこと |エッジログ
EdgeLog by Ecro Essays About Now 検索 Ctrl K Esc EN / KO ← エッセイ
/投稿 · 01/22 · 組み込み開発 LLM PCB ルーティング : 契約エンジニアが教えてくれたこと
私のエージェントのレジストリにある 16 個の PCB 配線テクニックのうち 7 個は、そのほとんどが契約エンジニアから学んだもので、エージェントでは実行できません。
-- · リンクをコピー X で共有 ハッカー ニュースに送信 Reddit に送信 · #pcb #kicad #ai-agents #hardware #build-in-public このページについて ボード
DRC でゼロを読み取ることもできましたが、そうしませんでした
電源とクロックは同じ幅でした
私よりもはるかに優れたコードは何ですか
2 つの約束があったが、厳しい約束は空に戻った
2枚の板を並べて置く
なぜこれら 11 件の違反が消えたのか、そして最初にどのように間違ってしまったのか
私たち自身の準備書面も間違っていました
失われたもの、そして不明なもの
失火も記録に残る
DRC でゼロを読み取ることもできましたが、そうしませんでした
電源とクロックは同じ幅でした
私よりもはるかに優れたコードは何ですか
2 つの約束があったが、厳しい約束は空に戻った
2枚の板を並べて置く
なぜこれら 11 件の違反が消えたのか、そして最初にどのように間違ってしまったのか
私たち自身の準備書面も間違っていました
失われたもの、そして不明なもの
失火も記録に残る
PCB 配線を LLM に任せました。うまくいかなかったので、採用しました。次に、返されたボードを前のボードの隣にネットごとに置き、その違いをエージェントが次回使用するためのルーティング技術のレジストリに変換しました。
16 のテクニックが使用されました。それぞれがエディター スロットに名前を付け、そのうちの 7 つは gui_card と読みます。これは、関数が存在せず、人が手動で行う必要がある場合に記述するものです。この盤面を救った手は7手のうちの1手です。
LLM は、熟練した人間のようにボードをレイアウトすることはできません。何U

契約エンジニアが購入した SD 145 は配線スキルではなく、どこで障害が発生したかを文書化したリストであり、そのリストはボードよりも長く保存されます。
記憶にある限り、モデルはフェイブル5でした。リポジトリには「Claude」だけが記録されていたため、これはログ行ではなく私の言葉です。
直径 21 mm、丸型、4 層、CR1632 コイン電池から動作します。メイン モジュールは、動作用の TDK ICM-42688-P を備えた Raytac AN54LV-15 (nRF54L15、内部チップ アンテナ) です。最初の内側の層はソリッド グラウンドであり、15 個のネット用の 2 層の実際の配線スペースが残ります。
残りは 3 期です。ネットは、信号を結び付ける回路図内の 1 本の線であり、ソフトウェアにおけるほぼ 1 つの変数です。ビアは層間でトレースを移動するために開けられる穴で、DRC はファブが拒否するものを基板上で掃引するチェックです。リンターに十分近い。
ボードはマウスで描画されません。 Python コードによって生成され、ルール ファイル内の定数は次のようになります。
TRACK_MIN_MM = 0.127 # 5 ミル
CLEARANCE_MM = 0.1 # ルーティング ルームを開くために締め付け (JLCPCB 対応のまま)
VIA_DRILL_MM = 0.3
VIA_DIA_MM = 0.5
2行目のコメントを読んでください。それは、それ自身の言葉で、配線のための余地を広げるためにクリアランスが厳しくなったと述べています。以下の内容の半分はその 1 行から始まります。
最初のボードには 143 件の DRC 違反が報告されました。診断では、そのうちの約 120 個 (84%) に機械的内容がないと分類されました。90 個は 1 つの不良ビア テンプレートからのもの、30 個のビアのそれぞれで 3 つのルール違反、16 個のファインピッチ クリアランス ルールの不一致、18 個のシルクスクリーン コスメティクス、3 個のプレーン フィル。これを加えると 120 ではなく 127 になります。これは、カテゴリーがいくつかを二重にカウントしており、私はそれを調整しなかったためです。実際に合計する計算では 16 個が残り、そのうち 8 個が実際の信号、つまり中庭の重なりでした。
ビア仕様を独自の 0.50/0.30 ルールに引き上げます。

上に印刷したものでは、1 回の変更で 90 がゼロになりました。診断では原因が 1 行で示されています。freerouting オートルーターの via テンプレートがボードの via ルールを無視しています。放出されたパッドスタックには、当社独自の 0.10 mm ルールに反して 0.075 mm の環状リングが残されており、診断の評決では、ボードはファブリーガルではないということでした。工場に入ったわけではない。
DRC は、最初の実行から 90 件すべてを正直に報告しました。青信号は嘘ではなかった。誰も読んでいませんでした。
この修正には代償が伴いました。ビアを仕様に合わせると、より太いビアが増加し、約 34 個の新たな密度違反が発生し、接続されていないネットは 6 個から 13 個になりました。合計は 143 から 64 に減少し、判決は DENSITY_LIMITED に変わりました。 64 はチューニング ルールで到達可能な下限であり、私が目標として設定した 20 には遠く及ばなかった。
数値は改善する一方、判決は悪化する。この組み合わせがオートルーターの終焉です。
DRC でゼロを読み取ることもできましたが、そうしませんでした
契約概要を書いているときに、私は一つのことを試みました。グローバル クリアランスをコードの 0.1 mm から 0.16 mm に広げると、ハンドオフ ボード上の 11 個の境界違反がゼロになります。また、接続されていないネットも 6 から 15 に引き上げられましたが、準備書面では今回ではなく前回の改訂でその数が考慮されています。
それは、ルーティングを放棄して DRC を通過することです。私はそれを使用しませんでしたが、概要にはそのように記載されています。
数字を見てみると、このトレードは改善しているように見えます。 １１人が０人になった。
電源とクロックは同じ幅でした
最も静かな欠陥がここにありました。バッテリ電源 VBAT、グランド GND、SPI クロック SPI_SCK。私が渡した基板では、3 つのネットすべてに同じ 0.2 mm のトレース幅が搭載されていました。
SPI クロックの幅で消費されるコイン型電池の電力。ルール違反がなかったため、DRC はこれを捕捉できません。戻ってきたボードでは、VBAT は 0.25 の間でステップしました。

1.0mm、GNDは0.3～0.7まで開きました。
私よりもはるかに優れたコードは何ですか
パーツローテーションの組み合わせを徹底的に調べても、コードは私にゲームを与えません。ここで機能したのは LLM ではなく for ループです。それらすべての 2^14、16384 を実行し、ネットの全長を 128.62 mm から 124.38 mm まで縮小しました。 3.3%の削減です。
人間がその検索を実行することはできません。人間にできることは、この基板が工場で受け入れられるものであるかどうかを知ることです。力技では超人的だが、許容性では無力、そしてその分裂こそが編集者の分野で最終的にエンコードされるものだ。 AI が依然としてハードウェア エンジニアを必要とする理由もそこにあります。
2 つの約束があったが、厳しい約束は空に戻った
1 つ目は、適切に計画された実験でした。外部のエンジニアが当社の基板を見ずに独自に基板をレイアウトしました。要旨は明示的でした。非表示のレイアウトと一致させようとしてはなりません。
採点基準は作業開始前に書かれ、封印されました。結果を見てから基準を書くと、好みの結果に向かって曲がっていくので、最初に基準を書いて、後戻りしていないことを確認できるようにしました。このメカニズムはフィンガープリントです。ハッシュを通じてファイルを実行し、結果の文字列を別に保持し、1 つの文字が変更されると別の文字列が得られます。私たち自身のボード ファイルも同様の扱いを受け、保管されていました。
戻ってきたのは、範囲外の回路図の書き換えでした。同日追加されたスコープゲートにより得点不可と判定された。
厳密な実験ではまったく比較できませんでした。
2つ目は通常の請負業務です。配置は固定され、自動ルーティングは 96% に達しました。 2 つの質問は、残りの 5 つのネットと 11 の DRC 違反を手動で完了することと、見落としていた設計上の問題を指摘することです。それはオープンで有料の取り組みでした。彼らは私たちのボードを見ました。これは盲目的ではなく、実験と呼ばれるべきではありません。
結果を出したのはセコ

nd。私は 200,000 ウォン、約 145 米ドルを支払いましたが、エンジニアは匿名のままです。
2枚の板を並べて置く
2 つのラウンドが一致しないため、2 つの列になっています。ラウンド 1 は積極的なもので、ラウンド 2 ではその一部が戻りました。モジュールのリセットは 14.4 mm から 18.3 mm に戻りました。デバッグ ネットは逆方向に移動し、ラウンド 2 でのみ最後の 2 つのビアを失います。これは、デバッグ ネットを内層から引き離す変更です。注文したボードはラウンド2です。
これら 4 つのネットはすべて、私が渡したボード上ですでに完全に接続されていました。オートルーターがそれらに触れる理由はありません。人間はそれらを削除し、再度描きました。それを待ってください。これは後にカタログが人間の名前で提出されることになる動きです。
どちらも、kicad-cli 9.0.8 を使用して実際のボード ファイルからレンダリングされ、同じジオメトリ、同じ角度です。
表側は話の半分にすぎません。本当の変化は内側の層 2 で起こりました。
ボード ファイルを解析して数えてみると、このレイヤを使用するネットは 6 から 8 に、セグメントは 12 から 79 になりました。元の 6 つのうち 3 つだけが残り、5 つの新しいものが移動されました。このレイヤの使用量が減ったわけではなく、別の方法で使用されました。残った 3 つのうち 2 つはデバッグ ネットです。
ボード全体のビアは 37 から 58 に、セグメントは 151 から 282 に増加しました。カウントが増加するにつれて、トレースは短くなりました。
なぜこれら 11 件の違反が消えたのか、そして最初にどのように間違ってしまったのか
ハンドオフ時に存在する 11 個の Hole_clearance 違反は、返されたボード上でゼロを読み取りました。ソース リポジトリの技術文書には、ビア直径のスイープ、つまりドリルを使用せずにすべてのビアでパッドを 0.5 mm から 0.6 mm に変更したことが記載されています。
それは間違いだと思いました。私が解析したものは次のとおりです。
11 件の違反は 5 つのビアから発生しています。同じビアに繰り返しフラグが立てられました。
測定されたクリアランスはすべて 0.25mm 要件をわずかに下回っており、0.0077 ～ 0.048 不足しています。

3mm。概要にはそれらがリストされています。
0.2017 0.2018 0.2092 0.2092 0.2112
0.2245 0.2248 0.2248 0.2280 0.2300 0.2423
問題のあるペア: C_SC1 パッドから INT_IMU (×2 経由)、TP5 パッドから SWDIO (×2 経由)、NRESET_MOD トラックから SWDCLK (×3 経由)、VBAT トラックから INT_IMU (×2 経由)、SPI_SCK/SPI_MISO トラックから GND (×2 経由)。
ドリル直径は 3 つのボードすべてで 0.3 mm です。パッドだけ変えました。
これら 5 つのビアは、返されたリビジョンのどこにも表示されません。座標の 0.05mm 以内には何も存在しません。
そのことから、私はパッドスイープではそれができなかったと結論付けました。hole_clearance はホールから銅までの距離を測定するため、ホールがそのままの状態でパッドを成長させても何も変わらないはずです。
その推論は、すでに取締役会に加わっている11人にも当てはまります。ライセンスが付与されていないのは、私が次にとったステップであり、したがってクレジットはぼったくりに属するということです。私が自分で書いた概要では、なぜパッドのサイズが重要なのかを説明しています。異質な銅が穴からはみ出さない。銅と銅のルールに従ってパッドの端から離れているため、穴と銅の間隔はクリアランスに環状リングを加えたものになります。
freerouting は銅線間の隙間 (0.1mm) のみを考慮し、KiCad の穴と銅線のルール (0.25mm) は決して考慮しません。ビアはパッド Ø0.5、ドリル Ø0.3 なので、環状リングはわずか 0.1mm です。銅をパッドの端から0.1mm離して保持すると、ドリルから0.2mmが得られ、0.05mm短くなります。
ドリル0.3でパッドをØ0.6にすると、アニュラーリングは0.15mmになります。クリアランスプラスリングは0.25mmになりました。
ただし、これは変更後にルーターが配置した銅線にのみ当てはまります。パッドを大きくしても基板上にすでにある銅は移動しないため、0.2017mm での既存の違反は 0.2017mm のままになります。スイープにより、今後のクラスが消滅します。すでにそこにある 11 個も同様に銅を動かす必要があり、両方のメカニズムが同じリワークで実行されました。
私の解析で判明したのは、これらの 5 つのサイトに銅があるということです

移動しました。サイズ変更は移動ではないため、スイープだけでイレブンを引退させることはできません。この不在によって分からないのは、誰が銅線を移動させたかということです。意図的な引き裂きと、新しいビア ルールに基づく完全な再配線では、同じ痕跡が残ります。
エラーの形状が興味深い部分であるため、私は間違ったバージョンを静かに置き換えるのではなく、上に残しておきます。私は実際に測定を行い、それが裏付けられていない因果関係の結論を導き出し、そのメカニズムを説明する私自身の要約の段落まであと 1 grep のところにありました。
私たち自身の準備書面も間違っていました
請負業者が私たちが送った掲示板で報告した 13 件の DRC 問題のうち 10 件は私たちのせいでした。バッテリーパッドは立ち入り禁止領域内にありますが、そこに痕跡を残さないように要求したため、最初から要求を満たすことは不可能です。私たちはパッド内ビアが許可されていることを決して書き留めませんでした。残りの 3 つは、リポジトリに存在しないフットプリント ファイルからのものです。
人間が「発見」したことの一部は、取締役会ではなく、事務手続きの欠陥でした。
基板はすでに完成していました。これに関してはもう何もすることがありませんでした。目標は特異でした。エージェントが次のプロトタイプで同じ間違いを繰り返さないようにすることです。
したがって、私は完成したボードをLLMに渡しませんでした。上の表にあげてみました。ネットごと、長さの差、cou経由

[切り捨てられた]

## Original Extract

Seven of the sixteen PCB routing techniques in my agent

LLM PCB Routing: What the Contract Engineer Taught It | EdgeLog
EdgeLog by Ecro Essays About Now Search Ctrl K Esc EN / KO ← Essays
/Post · 01 of 22 · Embedded Dev LLM PCB Routing : What the Contract Engineer Taught It
Seven of the sixteen PCB routing techniques in my agent's registry, most of them learned from a contract engineer, are ones it cannot execute.
-- · Copy link Share on X Submit to Hacker News Submit to Reddit · #pcb #kicad #ai-agents #hardware #build-in-public On this page The board
I could have made DRC read zero and did not
Power and clock were the same width
What code did far better than me
Two engagements, and the rigorous one came back empty
Putting the two boards side by side
Why those 11 violations disappeared, and how I got it wrong first
Our own brief was wrong as well
What it lost, and what is unknown
The misfires belong in the record too
I could have made DRC read zero and did not
Power and clock were the same width
What code did far better than me
Two engagements, and the rigorous one came back empty
Putting the two boards side by side
Why those 11 violations disappeared, and how I got it wrong first
Our own brief was wrong as well
What it lost, and what is unknown
The misfires belong in the record too
I handed PCB routing to an LLM. It did not go well, so I hired it out. Then I put the returned board next to the previous one, net by net, and turned the differences into a registry of routing techniques for the agent to use next time.
Sixteen techniques went in. Each one names an editor slot, and seven of them read gui_card , which is what you write when no function exists and a person has to do it by hand. The move that saved this board is one of the seven.
An LLM cannot lay out a board the way a practised human does. What USD 145 with a contract engineer bought is not routing skill but a written list of where it fails, and the list outlives the board.
The model was Fable 5, from memory. The repo only ever wrote down “Claude”, so that one is my word, not a log line.
21mm diameter, round, four layers, running off a CR1632 coin cell. The main module is a Raytac AN54LV-15 (nRF54L15, internal chip antenna) with a TDK ICM-42688-P for motion. The first inner layer is solid ground, which leaves two layers of real routing room for fifteen nets.
Three terms carry the rest. A net is one line in the schematic tying a signal together, roughly a single variable in software. A via is the hole drilled to move a trace between layers, and DRC is the check that sweeps the board for anything a fab would reject. Close enough to a linter.
The board is not drawn with a mouse. Python code generates it, and the constants in the rules file look like this.
TRACK_MIN_MM = 0.127 # 5 mil
CLEARANCE_MM = 0.1 # tightened (still JLCPCB-capable) to open routing room
VIA_DRILL_MM = 0.3
VIA_DIA_MM = 0.5
Read the comment on the second line. It says, in its own words, that the clearance was tightened to open up room for routing. Half of what follows starts from that one line.
The first board came back with 143 DRC violations. The diagnostic classified about 120 of them, 84%, as having no mechanical content: 90 from one bad via template, three rules broken on each of 30 vias, 16 fine-pitch clearance rule mismatches, 18 silkscreen cosmetics, 3 plane fills. That adds to 127, not 120, because the categories double-count a handful and I never reconciled it. Sixteen are left over on the arithmetic that actually adds up, and 8 of those were the real signal: courtyard overlaps.
Bringing the via spec up to our own 0.50/0.30 rule, the one printed above, took those 90 to zero in a single change. The diagnostic named the cause in one line: the freerouting autorouter’s via template ignores the board’s via rules. The padstack it emitted left a 0.075mm annular ring against our own 0.10mm rule, and the diagnostic’s verdict was that the board was not fab-legal. It was not entering a factory .
DRC reported all 90 of them honestly, from the first run. The green light did not lie. Nobody was reading it.
That fix came at a price. Bringing the vias to spec grew a field of fatter vias that created roughly 34 new density violations , and unconnected nets went from 6 to 13 . The total dropped from 143 to 64 while the verdict changed to DENSITY_LIMITED . Sixty-four was the floor reachable by tuning rules, nowhere near the 20 I had set as a target.
The number improves while the verdict gets worse. That combination is where the autorouter ended.
I could have made DRC read zero and did not
While writing the contract brief I tried one thing. Widening the global clearance from the 0.1mm in that code back to 0.16mm takes the 11 boundary violations on the handoff board to zero. It also pushed unconnected nets from 6 to 15, though the brief took that count on the previous revision, not this one.
That is passing DRC by abandoning routing. I did not use it, and the brief says so in those terms.
Watching one number, this trade reads as an improvement. Eleven became zero.
Power and clock were the same width
The quietest defect was here. Battery power VBAT, ground GND, SPI clock SPI_SCK. All three nets carried the same 0.2mm trace width on the board I handed over.
Coin-cell power drawn at the width of an SPI clock. DRC cannot catch this, because no rule was broken. On the board that came back, VBAT stepped between 0.25 and 1.0mm and GND opened up to between 0.3 and 0.7.
What code did far better than me
On exhaustively sweeping part-rotation combinations, code does not even give me a game. What worked here is a for loop, not an LLM. It ran all 2^14, 16384 of them, and pulled total net length from 128.62mm down to 124.38mm. A 3.3% reduction.
A person cannot run that search. What a person can do is know whether this board is a thing a factory will accept. Superhuman at brute force, helpless at admissibility, and that split is what the editor field ends up encoding. It is also the reason AI still needs hardware engineers .
Two engagements, and the rigorous one came back empty
The first was a properly designed experiment. An outside engineer laid the board out independently, without seeing ours. The brief was explicit: do not try to match a hidden layout.
The scoring criteria were written and sealed before the work started. Write criteria after seeing results and they bend toward the result you like, so I wrote them first and made it checkable that I had not gone back. The mechanism is a fingerprint: run the file through a hash, keep the resulting string separately, and one changed character gives a different string. Our own board file got the same treatment and was kept aside.
What came back was an out-of-scope schematic rewrite. A scope gate added the same day ruled it unscoreable.
The rigorous experiment produced no comparison at all.
The second was ordinary contract work. Placement was fixed and autorouting had reached 96%. Two asks: finish the remaining 5 nets and 11 DRC violations by hand, and point out design problems we had missed. It was an open, paid engagement. They saw our board. This is not blind and should not be called an experiment.
The one that produced results is the second. I paid KRW 200,000, about USD 145, and the engineer stays anonymous.
Putting the two boards side by side
Two columns because the two rounds disagree. Round 1 is the aggressive one, and round 2 gave some of it back: module reset went from 14.4 back up to 18.3mm. The debug net moves the other way, losing its last two vias only in round 2, which is the change that pulled it off the inner layer. The board that was ordered is round 2.
All four of these nets were already fully connected on the board I handed over. An autorouter has no reason to touch them. The human deleted them and drew them again. Hold onto that: it is the move the catalogue would later file under a human’s name.
Both are rendered from the actual board files with kicad-cli 9.0.8, same geometry, same angle.
The front side is only half the story. The real change happened on inner layer 2.
Parsing the board files and counting, the nets using this layer went from 6 to 8 and the segments from 12 to 79. Only three of the original six left, and five new ones moved in. The layer is not used less, it is used differently. Two of the three that left are debug nets.
Vias across the whole board went from 37 to 58, segments from 151 to 282. The traces got shorter while the count went up.
Why those 11 violations disappeared, and how I got it wrong first
The 11 hole_clearance violations present at handoff read zero on the board that came back. The technique document in the source repo credits the via diameter sweep, the change that took the pad from 0.5mm to 0.6mm, with the drill untouched, on every via.
I thought that was wrong. Here is what I parsed:
The 11 violations come from 5 vias . The same vias were flagged repeatedly.
The measured clearances are all just under the 0.25mm requirement, short by 0.0077 to 0.0483mm. The brief lists them:
0.2017 0.2018 0.2092 0.2092 0.2112
0.2245 0.2248 0.2248 0.2280 0.2300 0.2423
The offending pairs: C_SC1 pad to INT_IMU via ×2, TP5 pad to SWDIO via ×2, NRESET_MOD track to SWDCLK via ×3, VBAT track to INT_IMU via ×2, SPI_SCK/SPI_MISO track to GND via ×2.
Drill diameter is 0.3mm on all three boards . Only the pad changed.
Those 5 vias appear nowhere on either returned revision. Nothing sits within 0.05mm of their coordinates.
From that I concluded the pad sweep could not have done it: hole_clearance measures hole to copper, so growing the pad while the hole stays put should change nothing.
That reasoning holds for the eleven already on the board. What it does not license is the step I took next, that the credit therefore belongs to the rip-up. The brief I wrote myself explains why the pad size matters at all. Foreign copper is not held off the hole. It is held off the pad edge , by the copper-to-copper rule, so hole-to-copper is the clearance plus the annular ring:
freerouting only looks at copper-to-copper clearance (0.1mm) and never at KiCad’s hole-to-copper rule (0.25mm). The via is pad Ø0.5 with drill Ø0.3, so the annular ring is only 0.1mm. Hold copper 0.1mm off the pad edge and you get 0.2mm from the drill, which is 0.05mm short.
Take the pad to Ø0.6 with the drill at 0.3 and the annular ring becomes 0.15mm. Clearance plus ring is now 0.25mm.
But that only holds for copper the router places after the change. A bigger pad does not move copper already sitting on the board, so an existing violation at 0.2017mm stays at 0.2017mm. The sweep kills the class going forward. The eleven already there needed the copper to move as well, and both mechanisms ran in the same rework.
What my parse established is that copper at those five sites moved. A resize is not a move, so the sweep alone cannot have retired the eleven. What the absence does not tell me is who moved the copper: a deliberate rip-up and a full re-route under the new via rule leave the same trace.
I am leaving the wrong version above rather than quietly replacing it, because the shape of the error is the interesting part. I had a real measurement, drew a causal conclusion it does not support, and was one grep away from the paragraph in my own brief that explains the mechanism.
Our own brief was wrong as well
Ten of the 13 DRC issues the contractor reported on the board we sent were our fault. The battery pads sit inside a keep-out region while we asked for no traces there, which makes the request impossible to satisfy from the start. We never wrote down that via-in-pad was allowed. The other 3 came from footprint files missing from our repo.
Part of what the human “found” was a defect in our paperwork, not in our board.
The board was already finished. There was nothing left to do for this one. The goal was singular: stop the agent repeating the same mistakes on the next prototype.
So I did not give the LLM the finished board. I gave it the table above. Per net, the difference in length, via cou

[truncated]
