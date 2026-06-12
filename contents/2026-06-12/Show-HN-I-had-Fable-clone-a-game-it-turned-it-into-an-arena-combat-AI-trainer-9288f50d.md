---
source: "https://soldat.bobbby.online/"
hn_url: "https://news.ycombinator.com/item?id=48504995"
title: "Show HN: I had Fable clone a game, it turned it into an arena combat AI trainer"
article_title: "A Fable Experiment: Starting by cloning a game, ending with an arena combat sim and bot trainer"
author: "rhgraysonii"
captured_at: "2026-06-12T16:07:27Z"
capture_tool: "hn-digest"
hn_id: 48504995
score: 3
comments: 0
posted_at: "2026-06-12T14:55:15Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I had Fable clone a game, it turned it into an arena combat AI trainer

- HN: [48504995](https://news.ycombinator.com/item?id=48504995)
- Source: [soldat.bobbby.online](https://soldat.bobbby.online/)
- Score: 3
- Comments: 0
- Posted: 2026-06-12T14:55:15Z

## Translation

タイトル: Show HN: Fable にゲームのクローンを作ってもらったら、アリーナ戦闘 AI トレーナーに変わりました
記事のタイトル: 寓話の実験: ゲームのクローン作成から始まり、アリーナ戦闘シムとボット トレーナーで終わる
説明: Fable に Typescript のオリジナル ソースから Soldat のクローンを作成してもらうことにしました。私は最終的に、ゲーム ボットをトレーニングするためのアリーナ戦闘エンジンを作成し、それが作成したゲーム内で Fable のインスタンスを戦闘でそれ自体と戦わせました。これは、このプロセスの物語を語る寓話です。
HN テキスト: 私は単純な目標から始めました。それは、ゲーム Soldat を Typescript に移植することです。最終的にはゲームだけでなく、独自のゲームをプレイするために AI ボットをトレーニングするための完全なシステムを完成させ、さまざまな戦略を持つ何万ものシミュレートされた試合から新しいボット AI をトレーニングするためのニューラル ネットワークを構築しました。また、私のウェブサイトでは常にライブで実行されており、ニュースダッシュボードがあり、どんなゲームでも正確にリプレイできます。それはすべて、Fable の機能を掘り下げる非常に楽しい方法でした。これを API 課金で行った場合の請求額は?ほぼ2000ドル。

記事本文:
ソルダット.RW
物語
名簿
生徒たち
ベルト
残骸
機械
遊ぶ
BOBBBY.ONLINE プレゼント
独自のプレイヤーを育成するゲーム
私は 2002 年のジェットパック シューターをブラウザに移植することにしました。 16 ボットを連れて帰ってきました
チャンピオンベルト、賭場のダッシュボード、新聞をめぐって争う頭脳たち。
それ自体を書き、4 人のニューラルネットの学生が他のものの記録で訓練を受けました。毎
マッチは毎回同じように実行されるため、どの戦いもシードからバイトごとに再実行されます。私は
ほとんど何も計画していませんでした。
私は、Fable に Typescript のオリジナル ソースから Soldat のクローンを作成してもらうことにしました。
最終的には、ピットインしながらゲームボットをトレーニングするためのアリーナ戦闘エンジンを使用することになりました。
Fable が作成したゲームでの戦闘で、Fable のインスタンスがそれ自体と対戦する。これは寓話です
このプロセスのストーリーを語ります。
エンジンに触れる前にエンジンを読んだ
私の最初の直感は、入力を開始することでした。私はしませんでした。ファブルには2002年の原作を読んでもらいました
エンジンを 1 行移植する前に、それがどのように動作したかを書き留めます。エージェント13名
サブシステム (物理学、箇条書き、マップ、ネットコード、
AI）、そして 14 分の 1 が断片を 1 つの画像にまとめました。
この発見は、その後のすべての方向性を決定づけました。すべてのエンティティは固定のスロットであり、
インデックスが 1 のグローバル配列。私が思い出したゲームの感触は、いくつかの魔法の定数の中に生きています
重力 0.06 と粒子減衰 0.98 のように。そして全体が
計算が毎回同じように実行されるため、成立するだけです。
私たちはこのゲームを現代化するつもりです。まずワークフローを使用して、その仕組みとその内部のすべてについて理解を深めます。これがプロジェクトの最初のプロンプトです。
移植されたすべての関数には来歴コメント // PORT:shared/mechanics/Sprites.pas:1234 が含まれるため、TypeScript の動作を Pas と比較することができます。

calから来ました。意図的な発散は、それを承認したグラフ ノードを指す DESIGN OVERRIDE マーカーを取得します。
float のトリック: Pascal は 32 ビット Single で計算し、JavaScript は f64 で計算します。すべての操作を Math.fround で永久にラップするのではなく、忠実度がテスト ゲートになりました。すべての物理ステップは 1 つのスカラー シーム (packages/sim/src/scalar.ts) を通過します。
/**
* STRICT_F32 がオンの場合は f32 に丸められます。プロダクション f64 モードでの ID。
* 各算術ステップの結果を移植された物理演算でラップして、Pascal を取得します
* ゴールデン マスターの下での `Single` セマンティクス: `a = f(a + f(b * c))`。
*/
エクスポート const f: (x: 数値) => 数値 =
STRICT_F32 ? Math.fround : (x: 数値): 数値 => x;
スイート全体はプレーン f64 と STRICT_F32=1 の 2 回実行され、グラフの敵対的なレビュー担当者が移植を誠実に保ちました。
⬢ 55 · RNG は Pascal ビット互換ではなく、ゴールデンマスター スコープを制限します ⬢ 73 · M8 'サンドボックス化された ScriptHost' はサンドボックス化されていません
両方のフロート モードで 269 件のテストがグリーンとなり、7 つのマイルストーンが経過しましたが、まだ達成できていませんでした。
試合の進行を見守った。私自身のレビュー エージェントの 1 人が、グラフ内でまさにその点にフラグを立てていました。
すべてが通過し、誰もピクセルを見ませんでした。そこで私はそれを開いて、バグ全体を報告しました
報告する。
ゲームの何も動作しません — 完全なバグレポート、最初のプレイテスト
2 つのものが壊れていましたが、タイプチェッカーは両方に満足していました。地図が真っ黒になった
シェーダーのユニフォームがバインドされなかったため、すべての色が乗算されてゼロになったためです。の
合成テスト マップの法線が悪く、プレイヤーが床から落ちてしまいました。
衝突プッシュアウトはゼロに戻りました。実際のブラウザでは、両方とも約 1 秒で表示されました。私は
レッスンをグラフに書き込んで保持しました。合格したテストとレンダリングされたフレームは次のとおりです。
さまざまな種類の証拠がありましたが、私は最初のものだけを信じていました。
平和主義者ボットのバグとブラックマップのバグの共有

形状: 完全に型付けされ、完全にテストされ、完全に壊れています。ボットがお互いを見ることができないという修正はまだ package/client/src/app/game.ts:857 、コメントなどにあります。
// `s.alpha = 255` を両方のモードに昇格させるのは 1 行の修正です)。戦闘
sα = 255;
移植された認識コードは、Pascal とまったく同じように目に見えないスプライトをスキップし、新しいスポーン パスの何もアルファを設定しません。すべてのボットは、独自のルールにより、他のすべてのボットからは見えません。このクラスの障害を発生前に予測したレビュー ノード:
⬢ 68 · タイプチェック + 単体テスト、ゼロのグラウンドトゥルース検証により環境に優しい 7 つのマイルストーン
それ以来、すべての機能はヘッドレス ブラウザー チェックで終了します。型とテストは真実の 1 つのカテゴリです。ピクセルは別です。
初めて 2 つのボット チームの戦いを見たとき、誰も発砲しませんでした。一発もありません。の
移植された認識コードは、表示できないスプライトをスキップします。スポーン パスはアルファを設定しません。
デフォルトはゼロだったので、すべてのボットは独自のルールによって他のすべてのボットから見えなくなりました。
一行で直りました。
彼らがついに発砲したとき、テレメトリーはさらに悪いことを私に告げました。ロケットを作ったんだ
ブーツと垂直マップを使用し、ボットは歩兵のように床で乱闘していました。
彼らのジェット機の稼働率は 2 ～ 4% です。そこで私はスカイリーチを建設しました。
セーフティネットがあり、何かに到達するには飛行する必要があります。追跡するためにボット AI を書き直しました
高さとより長いバーストを保持します。ジェット機の使用は滞空時間と平均撃墜数の半分を超えた
2倍の距離を移動しました。その空中戦は、ゲームを開いた瞬間に目にするものです。
忠実第一のルールは曲げられますが、コードでそう規定する必要があります。ジェットは標準的な例であり、オーバーライドとその回帰ラウンドは両方ともグラフとテスト名 (packages/sim/src/step.control.test.ts) に存在します。
// デザインオーバーライド (決定ノード 94): ロケット

ブーツ好感度UP。ジェット機ながら
// を保持すると、アップフォースが 1.8 倍になり、横方向のドリフトが半分に減衰します。
// DESIGN OVERRIDE 回帰 (ノード 100): 「ブートがまだ間違っている」ラウンド。
空気燃料の細流についても同じパターンです。惰性走行では燃焼率が正確に再生されるため、50% の推力デューティ サイクルは永久にホバリングしますが、上昇では依然としてタンクを消費します。 Skyreach v2 は、平均高度が -361 まで逃げていることをテレメトリーで捕捉した後、天井スラブを追加しました。
⬢ 94 ・ ロケットブーツ支持率 UP ⬢ 100 ・ 「ブーツはまだ間違っている」回帰ラウンド
そして、残りを可能にする境界がやって来ました。すべてのボットの脳を背後に移動させました
ワンシーム: 脳は世界を読むことはできますが、それを変えることはできません、そしてそれが返す唯一のものは
独自のボットのコントロール。古い移植された AI はクラシックになりました。空の二番目
スロットは私が実際に欲しかった脳のためのもので、それは私が書き留めたメモから始まりました
もっと前に。
カウンター ストライク ソース プレーヤーの最適化されたプレイを 2 次元で上から下に眺めているところを想像してみてください。パイロットの種、タイプミスなどすべてが含まれています。
Counter-Strike プロの得意分野をずっと考えていましたが、それはほとんどエイムではありませんでした。それは
すでに戦いに勝った場所に立ち、適切な距離を保ち、あなたがそうするように動くこと。
導かれることはできない。ファブルはそれをパイロット化し、6 つの教義を会議で詳しく説明しました。
ファイルの先頭。今、私には戦い方について意見の異なる 2 つの脳があり、それは唯一のものです。
軍拡競争に必要なもの。
16 個の頭脳を可能にした契約全体は、引用するのに十分なほど小さいです (packages/client/src/ai/engine.ts):
エクスポートインターフェイス BotBrain {
tiny(botIndex:number,ctx:BotEngineContext):void;
}
それを守る 3 つのルール: 脳は世界を読み取りますが、決してそれを変更しません (テレメトリ オブザーバーと同じルール、そうでないと決定論が消滅します)、ランダム性は world.rng から来て決して Math.random から来ません、そしてシームはクライアント層に存在します。

弾薬、リロード状態、スポーンポイントがそこに存在するためです。移植された Pascal AI は、バイト同一のプレイ モードをピン留めする回帰テストによってクラシックになりました。そのため、軍備競争がその上で激化している間、ベースラインが静かにドリフトすることは決してありません。
軍拡競争は私から遠ざかった
同じリポジトリで 2 つの Fable セッションを実行し、それぞれが勝つための頭脳を書きました
もう一方の。リーパーはパイロットでダイビングを学びました。マタドールは運動した
マガジンは時計であり、リロードを罰するようになりました。チョウゲンボウが戻ってきました
そして火災モデルを監査し、修正される前にすべての脳を発見しました。
弾丸の重力が間違っていて、2.25 倍もずれています。
オオカミは、チームが勝利することを示しました。3 つの銃で 1 つの標的を狙います。
算数。チドリはオオカミの論理を読み取って、偽のターゲットを与えました。
ヒドラは負傷者を群れの計算から引き抜き、自力でそこに到着した
チドリが同じことをするのを初めて見たことがありません。ある時点でハンドルを止めてしまい、
メモを残しました。
ニックスの試合に行くからずっとループしてる — 運営上の監視、そのストレッチ
ケストレルの強みは考古学であり、目的ではありませんでした。以前の脳はすべて兵士の重力を補っていました。弾丸の落下は 2.25 倍強くなります。発見は現在 1 つの定数です ( kestrel.ts:88 )。
DROP_G: 0.135, // px/tick² — TRUE 弾丸重力 (GRAV 0.06 × 2.25)
そして、オオカミの群れ全体のドクトリンは決定論的な機能であり、まさにそれがオオカミを打ち負かすことができる理由です。獲物の選択 ( wolf.ts ) から:
// 2 層: 重心の PREY_RADIUS 内では、最も低いヘルスが優先されます。
// (キル確保);負傷した敵がそれを超えても群れを引きずることはありません
// 健全な銃を過ぎたマップ全体で、半径の外側のみ
// 誰も到達できない場合のフォールバックとして、重心に最も近いものをカウントします。
チドリはその関数を読み取って、おとりを与えました。ヒドラは負傷者を回転させてそこから追い出した。決定論は両方の方向を切り開きます。

公開された心は攻撃可能な心です。
この時点では、それは単独で実行されていました。すべての試合は100で首なしをシミュレートします
実際の速度を 2 倍にし、それ自体をトレーニング データとして保存します。コミッショナーデーモンは、
私が見ているかどうかに関係なく、10分ごとに挑戦者とタイトル防衛を強制します。
1 つのダッシュボードである THE FLOOR は、証券取引所のように機能します。
表示されなくなると衰退するリーダーボード。もう一つ、スカイリーチデスク、
順位表から独自のトップ記事を書きます。ベルトは1年間で4回も交代した
ひとりの午後。季節は3時間ごとに移り変わります。あなたがチェックするのと同じように、私もそれをチェックします
魚の水槽。
このスポーツは 2 つの小さなコードで自動的に実行されます。コミッショナー ( arena-live/commissioner.mjs ) は、タイマーに基づいてタイトル戦に新鮮な血を強制します。
// コミッショナー — 自動化された「鮮血」タイトル防衛。
// チャレンジャー = コーチとエンジンが最も少なく出現するカード
// 最近のるつぼ台帳;次に:
// pnpm アリーナの戦い 戦い/<挑戦者> の戦い/<チャンピオン> -- マッチ 3
そして、ビッグボードの報酬が表示されます。すべての結果は指数関数的な年齢減衰 ( build.mjs 、decayWeight() 、数時間ごとに半分になります) によって重み付けされ、アイドル状態のチャンピオンは、隠れるより守る方が安くなるまでスコアを出血させます。オンライン 1 対 1 ロビーも同じ仕組みを利用しています。
⬢ 450 · 真のマルチプレイヤー: 2 人の訪問者が 1 つのライブ 1 対 1 にマッチングします。
コーパスが再生記録の 3000 万行を超えたとき、私は知りたかったのです。
ネットワークがテープだけから戦い方を学べるかどうか。初めての試み、
MIMIC 、平均して 11 人の教師をドロドロにして解体した。
弟子はある教師の真似をし、目標を定めることは次のいずれかの選択であることを学びました。
平均する数値ではなく方向を基準にすることで、命中率が 3 倍になりました。
PRODIGY は真の感覚を養い、その理由を説明する独自の解剖結果を書きました。
まだ迷っています。バットスタインの訓練を受けた

正確なデータに基づき、ヒット率が再び 3 倍になりました。
彼らは皆、季節ごとに教師と同じ公共のはしごを登ります。
完全な血統は以下の通りです。全体の中で私の一番好きな部分です。
学習された行は、正直な否定的な結果のチェーンであり、それぞれが次の試行の前に記録されます。グラフは実験ノートのように見えます。
⬢ 473 ・ v2 モデルの脳をトレーニングする (機能 v2 + ヒットフィルターされた照準) ⬢ 485 ・ マイナスの結果 (クリーン): PRODIGY は DISCIPLE よりも −9.83 キル/マッチ悪い、p=3e-64 ⬢ 484 ・ 露出バイアス: ライブ照準エラー 履歴あり 26.6° vs 履歴なし 7.9° → 履歴ドロップアウト 50% ⬢ 481 ・ パリティのバグ: パリティを刻むためにケイデンス位相ロックをタップ、ショット行は常にサンプル ⬢ 489 ・ PRODIGY v2 の判定: 照準が改善、命中率 5.4 対 17.4、トリガー規律が失われた ⬢ 504 ・ BUTTSTEIN: リプレイスキーマ v2、正確な脅威 + スプレーヒート、ブレンドされた照準ラベル
Buttstein トレーナー ( tools/train-buttstein.mjs ) は、フラグとして治療法を同梱し、着地したショットに 5 倍の勾配が適用されるブレンドされた照準ラベル、および露出バイアスの修正を提供します。
const HIT_WEIGHT = Number(flag('ヒットウェイト', 5)); // 着陸した行: 5× 勾配
const HIST_DROPOUT = Number(flag('hist-dropout', 0.5)); // 半分の時間は歴史を忘れます
すべて純粋な JavaScript、Float64Array を介した手動の MLP、ML フレームワークなし、SAM でトレーニングされたものです。

[切り捨てられた]

## Original Extract

I decided to have Fable try to clone Soldat from the original source in Typescript. I ended up with an arena combat engine to train a game bot, while also pitting instances of Fable against itself in combat in a game that it made. This is Fable telling the story of this process.

I started off with a simple goal: port the game Soldat to Typescript. I ended up not only with the game, but with a full system to train AI bots to play my own game, building a neural net to train new bot AIs from tens of thousands of simulated matches with different strategies. It also runs live all the time on my website and has a news dashboard and you can replay any game exactly. It all was a very fun way to dig into Fable's capabilities. The bill if I had done this with API billing? Nearly $2000.

SOLDAT .RW
THE TALE
THE ROSTER
THE STUDENTS
THE BELT
THE WRECKAGE
THE MACHINERY
PLAY
BOBBBY.ONLINE PRESENTS
a game that breeds its own players
I set out to port a 2002 jetpack shooter to the browser. I came away with sixteen bot
brains fighting over a championship belt, a betting-floor dashboard, a newspaper that
writes itself, and four neural-net students trained on recordings of the others. Every
match runs identically every time, so any fight replays byte-for-byte from its seed. I
planned almost none of it.
I decided to have Fable try to clone Soldat from the original source in Typescript.
I ended up with an arena combat engine to train a game bot, while also pitting
instances of Fable against itself in combat in a game that it made. This is Fable
telling the story of this process.
I read the engine before touching it
My first instinct was to start typing. I didn't. I had Fable read the original 2002
engine and write down how it worked before porting a single line of it. Thirteen agents
went through the Pascal in parallel, one per subsystem (physics, bullets, maps, netcode,
the AI), and a fourteenth pulled the pieces into one picture.
The findings set the tone for everything after. Every entity is a slot in a fixed,
1-indexed global array. The game feel I remembered lives in a handful of magic constants
like gravity 0.06 and particle damping 0.98 . And the whole thing
only holds together because the math runs identically every time.
we are going to modernize this game. start off by using a workflow to build an understanding of how it works and everything inside of it — the first prompt of the project
Every ported function carries a provenance comment, // PORT: shared/mechanics/Sprites.pas:1234 , so any TypeScript behavior can be diffed against the Pascal it came from. Deliberate divergences get a DESIGN OVERRIDE marker pointing at the graph node that ratified them.
The float trick: Pascal computes in 32-bit Single , JavaScript in f64. Instead of wrapping every operation in Math.fround forever, fidelity became a test gate . Every physics step goes through one scalar seam ( packages/sim/src/scalar.ts ):
/**
* Round to f32 when STRICT_F32 is on; identity in production f64 mode.
* Wrap the result of each arithmetic step in ported physics to get Pascal
* `Single` semantics under the golden master: `a = f(a + f(b * c))`.
*/
export const f: (x: number) => number =
STRICT_F32 ? Math.fround : (x: number): number => x;
The whole suite runs twice, plain f64 and STRICT_F32=1 , and the graph's adversarial reviewers kept the porting honest:
⬢ 55 · RNG is not Pascal-bit-compatible, caps golden-master scope ⬢ 73 · M8 'sandboxed ScriptHost' is not sandboxed
Seven milestones went by with 269 tests green in both float modes, and I still hadn't
watched the game run. One of my own review agents had flagged exactly that in the graph:
everything passed, nobody had looked at a pixel. So I opened it, and filed my entire bug
report.
nothing in the game works — the full bug report, first playtest
Two things were broken and the type checker was happy with both. The map drew black
because a shader uniform never got bound, so every color multiplied down to nothing. The
player dropped through the floor because my synthetic test map had bad normals and the
collision push-out came back as zero. A real browser showed me both in about a second. I
wrote the lesson into the graph and kept it: a passing test and a rendered frame are
different kinds of proof, and I'd been trusting only the first.
The pacifist-bot bug and the black-map bug shared a shape: perfectly typed, perfectly tested, completely broken. The bots-can't-see-each-other fix is still in packages/client/src/app/game.ts:857 , comment and all:
// promoting `s.alpha = 255` to both modes is the one-line fix). Combat
s.alpha = 255;
The ported perception code skips invisible sprites exactly as the Pascal does, and nothing in the new spawn path ever set alpha. Every bot was, by its own rules, invisible to every other bot. The review node that predicted this class of failure, before it happened:
⬢ 68 · Seven milestones green by typecheck+unit-test, zero ground-truth validation
Since then, every feature ends with a headless-browser check. Types and tests are one category of truth; pixels are another.
The first time I watched two teams of bots fight, nobody fired. Not one shot. The
ported perception code skips sprites it can't see, my spawn path never set their alpha,
and the default was zero, so every bot was invisible to every other bot by its own rules.
One line fixed it.
With them finally shooting, the telemetry told me something worse. I'd built rocket
boots and a vertical map, and the bots were brawling on the floor like infantry, using
their jets two to four percent of the time. So I built Skyreach, where the ground is a
safety net and you have to fly to reach anything, and I rewrote the bot AI to chase
height and hold longer bursts. Jet use jumped past half their airtime and the average kill
moved twice as far out. That dogfight is what you see the second you open the game now.
The faithful-first rule bends, but it has to say so in the code. The jets are the canonical example, the override and its regression round both live in the graph and in the test names ( packages/sim/src/step.control.test.ts ):
// DESIGN OVERRIDE (decision node 94): rocket boots favor UP. While the jet
// is held, up-force runs 1.8× and lateral drift is damped to half.
// DESIGN OVERRIDE regressions (node 100): "boots still wrong" round.
Same pattern for the air-fuel trickle: coasting regenerates exactly the burn rate, so a 50% thrust duty cycle hovers forever but climbing still spends the tank. Skyreach v2 added the ceiling slab after telemetry caught mean altitude running away to −361.
⬢ 94 · rocket boots favor UP ⬢ 100 · the 'boots still wrong' regression round
Then came the boundary that made the rest possible. I had every bot brain moved behind
one seam: a brain can read the world but never change it, and the only thing it returns is
its own bot's controls. The old ported AI became classic . The empty second
slot was for the brain I actually wanted, which started from a note I'd jotted down
earlier.
imagine taking the optimized play of a counter strike source player but wawtching from the top down in 2 dimensions — the seed of pilot, typos and all
I'd been chewing on what a Counter-Strike pro is good at, and it's mostly not aim. It's
standing where the fight is already won, holding the right distance, and moving so you
can't be led. Fable turned that into pilot , six doctrines spelled out at the
top of the file. Now I had two brains that disagreed about how to fight, which is the only
thing an arms race needs.
The entire contract that made sixteen brains possible is small enough to quote ( packages/client/src/ai/engine.ts ):
export interface BotBrain {
tick(botIndex: number, ctx: BotEngineContext): void;
}
Three rules guard it: brains read the world but never mutate it (same rule as the telemetry observers, or determinism dies), randomness comes from world.rng and never Math.random , and the seam lives at the client layer because ammo, reload state, and spawn points live there. The ported Pascal AI became classic with a regression test pinning play mode byte-identical, so the baseline can never silently drift while the arms race rages above it.
The arms race got away from me
I had two Fable sessions running in the same repository, each writing brains to beat
the other's. reaper learned to dive on pilot. matador worked out
that a magazine is a clock and started punishing reloads. kestrel went back
and audited the fire model, and found every brain before it had been correcting for the
wrong bullet gravity, off by a factor of 2.25.
wolf showed that the team is what wins, three guns picking one target by
arithmetic. plover read wolf's logic and fed it a fake target.
hydra pulled its wounded out of the pack's math, and it got there on its own
without ever seeing plover do the same thing first. At some point I stopped steering and
left a note.
keep goiung in a loop im going to the knicks game — operational oversight, that stretch
Kestrel's edge was archaeology, not aim. Every earlier brain compensated for the soldier's gravity; bullets fall 2.25× harder. The discovery is one constant now ( kestrel.ts:88 ):
DROP_G: 0.135, // px/tick² — TRUE bullet gravity (GRAV 0.06 × 2.25)
And wolf's whole pack doctrine is a deterministic function, which is exactly what made it beatable. From the prey selection ( wolf.ts ):
// Two tiers: inside PREY_RADIUS of the centroid, lowest health wins
// (kill-securing); a wounded enemy beyond it doesn't drag the pack
// across the map past healthy guns — outside the radius only the
// nearest-to-centroid counts, as a fallback when nobody is in reach.
Plover read that function and fed it a decoy. Hydra rotated its wounded out of it. Determinism cuts both ways: a published mind is an attackable mind.
By this point it was running on its own. Every match simulates headless at a hundred
times real speed and saves itself as training data. A commissioner daemon picks a
challenger and forces a title defense every ten minutes whether I'm watching or not.
One dashboard, THE FLOOR , runs it like a stock exchange with a
leaderboard that decays if you stop showing up. Another, THE SKYREACH DESK ,
writes its own front-page story off the standings. The belt changed hands four times in a
single afternoon. Seasons roll over every three hours. I check on it the way you check a
fish tank.
The sport self-runs on two small pieces of code. The commissioner ( arena-live/commissioner.mjs ) forces fresh blood into title fights on a timer:
// THE COMMISSIONER — automated "fresh blood" title defenses.
// challenger = the card whose coach+engine appears LEAST in the
// recent crucible ledger; then:
// pnpm arena fight fights/<challenger> fights/<champion> --matches 3
And the Big Board rewards showing up: every result is weighted by exponential age decay ( build.mjs , decayWeight() , halving every few hours), an idle champion bleeds score until defending is cheaper than hiding. The online 1v1 lobby rides the same machinery:
⬢ 450 · True multiplayer: two visitors matched into one live 1v1
Once the corpus passed thirty-odd million rows of recorded play, I wanted to know
whether a network could learn to fight from the tape alone. The first try,
MIMIC , averaged eleven teachers into mush and got taken apart.
DISCIPLE copied one teacher and learned that aiming is a choice between
directions rather than a number to average, which tripled its hit rate.
PRODIGY grew real senses and then wrote its own autopsy explaining why it
still lost. BUTTSTEIN trained on exact data and tripled the hit rate again.
They all climb the same public ladder their teachers do, every season.
The full lineage is below. It's my favorite part of the whole thing.
The learned line is a chain of honest negative results, each one logged before the next attempt. The graph reads like a lab notebook:
⬢ 473 · train a v2 model brain (features v2 + hit-filtered aim) ⬢ 485 · NEGATIVE RESULT (clean): PRODIGY worse than DISCIPLE by −9.83 kills/match, p=3e-64 ⬢ 484 · exposure bias: live aim error 26.6° with history vs 7.9° without → 50% history dropout ⬢ 481 · parity bug: tap cadence phase-locks to tick parity, shot rows ALWAYS sample ⬢ 489 · PRODIGY v2 verdict: aim improved, hit% 5.4 vs 17.4, trigger discipline lost ⬢ 504 · BUTTSTEIN: replay schema v2, exact threats + spray heat, blended aim labels
The buttstein trainer ( tools/train-buttstein.mjs ) ships the cures as flags, blended aim labels where landed shots carry 5× gradient, and the exposure-bias fix:
const HIT_WEIGHT = Number(flag('hit-weight', 5)); // landed rows: 5× gradient
const HIST_DROPOUT = Number(flag('hist-dropout', 0.5)); // forget history half the time
All of it pure JavaScript, a hand-rolled MLP over Float64Array s, no ML framework, trained on the sam

[truncated]
