---
source: "https://www.jamesdrandall.com/posts/thrust_ai_powered_software_archaeology/"
hn_url: "https://news.ycombinator.com/item?id=48865903"
title: "AI Can't Recreate the Thrust Game (But It Can Help You Understand It)"
article_title: "AI Can't Recreate Thrust (But It Can Help You Understand It) | James Randall"
author: "msephton"
captured_at: "2026-07-10T22:51:33Z"
capture_tool: "hn-digest"
hn_id: 48865903
score: 1
comments: 0
posted_at: "2026-07-10T22:04:45Z"
tags:
  - hacker-news
  - translated
---

# AI Can't Recreate the Thrust Game (But It Can Help You Understand It)

- HN: [48865903](https://news.ycombinator.com/item?id=48865903)
- Source: [www.jamesdrandall.com](https://www.jamesdrandall.com/posts/thrust_ai_powered_software_archaeology/)
- Score: 1
- Comments: 0
- Posted: 2026-07-10T22:04:45Z

## Translation

タイトル: AI はスラスト ゲームを再現することはできません (ただし、理解には役立ちます)
記事のタイトル: AI は推力を再現できない (しかし理解には役立つ) |ジェームズ・ランドール
説明: AI を使用して 6502 アセンブリをリバース エンジニアリングすることにより、1986 年の BBC マイクロ ゲームを再作成します。Q7.8 固定小数点物理を抽出し、SN76489 サウンド チップをエミュレートし、オリジナルがなぜそのように感じられるのかを発見します。現代の共犯者によるソフトウェア考古学。

記事本文:
ジェームズ・ランドール
ソフトウェア開発、ビジネス、テクノロジーについて思索します。
AI は推力を再現することはできません (ただし、理解するのには役立ちます)
併合
早期アクセス
私がゼロから構築したターンベースの 4X ストラテジー ゲームです。カスタム WebGPU エンジン、8 つの文明、インストールは必要ありません。今すぐブラウザでプレイしてください。
ブラウザで無料でプレイ
Steam のウィッシュリスト
私はクロードに、1986 年の古典的なゲーム Thrust をブラウザで再現するよう依頼しました。それはスロープを生み出しましたが、その後、事態は制御不能になりました。
Thrust は、BBC Micro で放映された私のお気に入りのゲームの 1 つでした。ジェレミー C. スミスによって書かれ、1986 年に出版されたこのゲームは、驚くべき物理学とゲームプレイを備えた一見奥深いゲームです。重力と勢いと戦いながら、洞窟を通って船を操縦し、燃料を集め、砲塔の砲撃を避け、ポッドを回収してボーナス ポイントを獲得します。ジェレミーはその後、ピーター・アーヴィンとともにさらに印象的な『Exile』を制作し、1992年に事故で悲劇的に亡くなりました。『Thrust』を書いたとき、彼は16歳から18歳の間でした。オリジナルをオンラインでプレイできます。
私の隣の机には BBC Master があり、今でも時々、他のクラシック音楽と一緒にそこで Thrust を起動します。これは、Elite、Exile、Holed Out と並んで私が繰り返しプレイするゲームの 1 つです。私は現在、このうち 3 つをさまざまな方法で再作成しました。4 つ目はますます避けられないように見えます。
ともかく。最近のある朝、何気なく Claude Code にブラウザで作成するように頼んだので、私は Thrust のことを考えていたと思います。私は OpenAI と Anthropic の最新の機能宣言を読んでいたと思います。そのため、非常に包括的な仕様をまとめ、元の逆アセンブルされたソース コードとスクリーンショットへのアクセスを許可し、「私のために Thrust を再作成してください」と言いました。
それはslという用語が使われる何かを生み出しました

opは優しすぎるだろう、それは非常にぼんやりとThrustに似ていた - スキャンラインのようなものを持っていた - しかし、それは本当に恐ろしいものでした。重力も正しく機能しておらず、船は適切に落下せず、制御は奇妙に感じられ、そしてそれはただ…悲惨でした。ある意味、それが機能し、見た目も Thrust に似たものを作成したことは驚くべきことですが、それはプレイ可能ではなく、本物の優雅さや美しさにはまったく近づきませんでした。
それが『Thrust』のようなゲームの特徴です。表面的に似たものはすぐに打ち破ることができます。デバイスのフレーム レートで実行し、標準的なデルタタイム物理を使用し、洞窟を描画するだけです。しかし、それはスラストのようなものではありません。魔法は、特定のタイミング、船の重量、勢いの作り方にあります。特にオリジナルをプレイしたことがある場合は、その詳細がすべてであり、テキストの説明に基づいて動作する AI は、オリジナルのソースですらそれらをキャプチャできないことがわかります。
しかし、それは私に興味を持ちました。原作はどうでしたか？開発者がこれを 8 ビットで動作させるために使用したトリックは興味深いもので、ちょっとした考古学のセッションになりました。私はすぐに、キーラン・コネルによるオリジナルのソースを見事にコメント付きで分解したものを見つけ、それをクロードに入力して質問していることに気づきました。
ここからが興味深いことになります。 AI がコードを書いたからではなく、コード自体は複雑ではなく、32K の RAM で実行された 1986 年のゲームです。クロードが 6502 アセンブリを調査するための並外れたツールであることが判明したためです。逆アセンブルされたソースのブロックを入力して、「レベル データはどのように機能するのか?」と尋ねることもできます。または「ここで物理モデルは何をしているのですか?」元のコードが何をしていたかについての詳細かつ正確な説明が得られます。
公平を期すために、私は仕事をしていたので、

私は逆アセンブルされたソース コードにコメントしましたが、コメントとアセンブリの両方から情報を抽出でき、ゲームがどのように動作するかの詳細な説明を思いつくことができたとしてもです。私の感覚では、適切な領域に焦点を当てるためのコメントがなければ、あまり役に立たなかったでしょう。しかし、それでも、作業は非常に簡単になり、より楽しくなりました。そして、はい、コードからコメントを削除して、Claude がどれだけうまく機能するかを確認することになりそうです。
そうしているうちに、その答えを元のゲームを再作成するための基礎として使用できることに気づき、クロードにさまざまなサブシステムの仕様を作成するよう依頼し始めました。私が生成した仕様のほとんどは、ソース コードの spec フォルダーにあります。いつかきちんと書き上げるかもしれませんが、今のところ、原文のニュアンスについての良い洞察が得られます。私が思っていた以上に、かなり多くのことが起こっています。たとえば、発電機を叩くと砲塔が一時的に発砲を停止することに私はまったく気づきませんでした。また、実際のコードを読んだときにのみ明らかになる発射角度には微妙な点があります。
物理学は、掘り下げるのに最も興味深い分野の 1 つでした。 Thrust は Q7.8 固定小数点演算を使用します。これは、浮動小数点ハードウェアを持たない 8 ビット マシンでは一般的な手法です。回転システムは、力の成分のルックアップ テーブルを備えた 32 のステップを使用します。
しかし、本当に難しいのはタイミングでした。私の最初の実装では、分解で得た正しい定数 (同じ重力、同じ推力値、同じ抗力) を使用しましたが、船はあまりにも速く、機敏すぎました。適切な重さがありませんでした。定数はオリジナルと同じだったので、タイミングの問題に違いありませんでした。
そして、そうでした。オリジナルは BBC Micro の 50 時間で物理的に動作しません

z 垂直同期レート。ティック ループはフレームごとに少なくとも 3 センチ秒待機し、実効レートは約 33.33 Hz になります。しかし、それはさらに進んでおり、各ティック内では、物理アップデートは 16 ティック ウィンドウごとに 6 つのアクティブ スロットのみにゲートされます。物理ステップの核心は次のようになります。
/** 16 ティック ウィンドウごとに 6 つのアクティブな物理スロット */
private static readonly ACTIVE_SLOTS = 新しい Set ([ 0 , 3 , 5 , 8 , 11 , 13 ]);
プライベートティックステップ (入力: ThrustInput) : void {
const スロット = this 。ティックカウンター & 0x0F ;
これ 。 tinyCounter = (this .tickCounter + 1) & 0xFF ;
// 回転: 4 ティックごとに 3 回、整数ステップのみ
if (( スロット & 0x03 ) !== 0 && 入力 . 回転 !== 0 ) {
s 。角度 = (( s . 角度 + 入力 . 回転 ) + 32 ) % 32 ;
}
const isActiveSlot = ThrustPhysics 。 ACTIVE_SLOTS 。 (スロット) があります。
// 強制計算 — アクティブなスロットのみ (16 ティックごとに 6 つ)
if ( isActiveSlot ) {
s 。 ForceY += this 。重力 ;
if (入力.推力) {
s 。 ForceY += ANGLE_Y [ angleIdx ] / ( 1 << this .massShift );
s 。 ForceX += ANGLE_X [ angleIdx ] / ( 1 << this . MassShift );
}
// 直線ドラッグ
s 。力X *= 1 - 1 / 64 ; // X: *= 63/64
s 。力 Y *= 1 - 1 / 256 ; // Y: *= 255/256
}
// 位置の統合 — ティックごと
s 。 x += s 。力X ;
s 。 y += s 。力Y ;
}
このゲートにより、約 12.5 Hz の実効力/抗力率が得られます。重力と推力は、すべてのフレームではなく、特定の 6 ティックにのみ適用されます。回転にも独自のゲートがあり、4 ティックごとにスキップされます。これらは任意の数字ではありません。これらは 6502 ソースの正確なパターンであり、ゲームの雰囲気を定義します。これらを使用すると、すぐに違和感が生じます。非対称のドラッグも興味深いものです。Y 軸 (255/256) よりも X 軸 (63/64) の方がはるかに強いため、水平方向の動きが ver よりも「粘り強く」感じられます。

ティカル。
これらのタイミングを正しく取得し、定数だけではなくオリジナルの正確な更新リズムと一致させると、完璧だと感じました。 BBC エミュレータと私のバージョンを切り替えることができ、コントロールの感触は同じです。
それが私にとって本当に興味深かったことです。通常の物理学を使用して Thrust のようなゲームを作成することは、特にコーディング AI を使用すると簡単です。独特の感触と重量を備えた Thrust を再作成するには、ゲーム ループ内の個々の物理更新のタイミングに至るまで、オリジナルがどのように機能するかを正確に理解する必要がありました。
本格的に焦点を当てる必要があったもう 1 つの領域はサウンドでした。 TypeScript バージョンの Elite を実行したとき、エミュレータから直接サウンドをサンプリングしました。 Elite のサウンドは非常に離散的で、短く鋭いエフェクトであるため、これはうまくいきました。ビームと軍事レーザーのエフェクトはオフになっていますが。しかし、Thrust では、エンジンはキーの押下に応答する継続的なドローンであり、爆発には特定のエンベロープがあります。それらをサンプリングすることもできましたが、それは…間違っていると感じました。
そこで私は別のアプローチを取ることにし、代わりに SN76489 サウンド チップ (作業中に学んだことですが、Sega Master System に含まれていたものと同じチップ) とそれに接続する BBC MOS インターフェイスを再作成しました。 MOS (BBC のマシン オペレーティング システム) は、OSWORD 呼び出しとメモリ マップされた I/O を使用して、ゲームがサウンド チップと通信するためのインターフェイスを提供しました。
これに大きな助けとなったのは、Toby Nelson がまとめた逆アセンブルされた MOS コード、BBC Advanced User Guide および SN76489 チップ仕様でした。これらすべてを Claude に入力し、サウンド システムの包括的な仕様を生成し、そこから AudioWorklet で実行されるエミュレートされたサウンド システムを構築することができました。完全な BBC MOS エンベロープ プロセッサ (OSWORD 7/8) がチップを駆動します

オーディオスレッドのエミュレータ。
結果？音は同じです。そして、他の手段でこれを実現するのは本当に困難でした。これはタイミングに非常に依存し、正確なエンベロープ形状とチップの動作に依存するため、実際のハードウェアをエミュレートする必要があります。
サウンドシステム自体は、そのレイヤーの仕方が非常にエレガントです。最上位レベルでは、サウンドを再生するということは、元のゲームで使用されていたのと同じ OSWORD 7 パラメータ (チャネル、振幅、ピッチ、継続時間) を送信することを意味します。
const サウンド = {
own_gun : { チャンネル : 0x0012 、振幅 : 1 、ピッチ : 0x50 、持続時間 : 2 }、
爆発_1 : { チャンネル : 0x0011 、振幅 : 2 、ピッチ : 0x96 、持続時間 : 100 }、
爆発_2 : { チャンネル : 0x0010 、振幅 : 3 、ピッチ : 0x07 、継続時間 : 100 }、
hostile_gun : { チャンネル : 0x0013 、振幅 : 4 、ピッチ : 0x1e 、持続時間 : 20 }、
エンジン : { チャンネル : 0x0010 、振幅 : - 10 、ピッチ : 0x05 、持続時間 : 3 }、
};
これらのパラメータは、SN76489 チップ エミュレータを駆動する MOS エンベロープ プロセッサに入力され、両方ともオーディオ スレッド上の AudioWorklet で実行されます。チップ エミュレータは、ハードウェア レベルでサンプルを生成します。つまり、10 ビット周期カウンタを備えたトーン チャネル、15 ビット線形フィードバック シフト レジスタを備えたノイズ チャネル、および元のシリコンに一致するステップあたり -2dB のボリューム テーブルです。
generate(out:Float32Array、オフセット:number、length:number、sampleRate:number):void {
const ステップ = 250000.0 / サンプルレート ; // オーディオ サンプルごとのチップ クロック
for ( let i = 0 ; i < 長さ ; i ++ ) {
サンプル = 0 とします。
// トーンチャンネル0～2
for ( ch = 0 ; ch < 3 ; ch ++ ) {
これ。カウンタ [ch] -= ステップ ;
if ( this .counter [ch] <= 0 ) {
const p = this 。ピリオド [ ch ] || 1024;
これ 。カウンタ [ch] += p ;
これ 。極性 [ch] = - これ。極性 [ ch ];
}
サンプル += これ 。極性

【ch】・これ。ボリューム [ ch ];
}
// ノイズ チャネル — 15 ビット LFSR
// ...
サンプル += (( this .lfsr & 1 ) ? 1 : - 1 ) * this .巻 [ 3 ];
out [オフセット + i] = サンプル;
}
}
隅々までエミュレーションです。 MOS はオーディオ スレッド上で 100 Hz で動作し、実際の BBC ハードウェアと同じレートでエンベロープを処理し、チップ レジスタを更新します。その結果、本物のサウンドが得られます。
逆アセンブルしたソースからフォント、グラフィックス、レベルデータを比較的簡単に抽出することができました。ソースにはそれを行うツールがいくつかあります。1 つは平地地形データをデコードし、もう 1 つは BBC のフレーム バッファーをエミュレートして出力を PNG に変換することで船の回転スプライトを抽出します。
私は当初、船のベクトル アプローチを試しましたが、各回転で数学的に描画するだけでした。しかし、オリジナルには 32 の回転角度ごとにハードコードされたスプライトがあり、各位置が正しく見えるようにジェレミー スミスが手作業で最適化していることが判明したため、BBC 解像度では粗く見えました。実際のスプライトを抽出して使用すると、正しく見えました。
地形のレンダリングでは、オリジナルのスキャンライン パリティ ポリゴンの塗りつぶし (線を 1 つおきに描画) が使用され、BBC Micro の特徴的な外観が得られます。内部解像度は 320×256 で、オリジナルと同じです。私のバージョンのビューポートはわずかに大きく、オリジナルには周囲に境界線があります。

[切り捨てられた]

## Original Extract

Recreating a 1986 BBC Micro game by using AI to reverse-engineer 6502 assembly — extracting Q7.8 fixed-point physics, emulating the SN76489 sound chip, and discovering why the original feels the way it does. Software archaeology with a modern accomplice.

James Randall
Musings on software development, business and technology.
AI Can't Recreate Thrust (But It Can Help You Understand It)
Annhexation
Early access
A turn-based 4X strategy game I built from scratch — custom WebGPU engine, eight civilisations, no install. Play it in your browser right now.
Play free in browser
Wishlist on Steam
I asked Claude to recreate the classic 1986 game Thrust for me in the browser. It created slop but then things spiralled out of control.
Thrust was one of my favourite games on the BBC Micro — written by Jeremy C. Smith and published in 1986, it’s a deceptively deep game with amazing physics and gameplay. You pilot a ship through caverns, collecting fuel, avoiding turret fire, and retrieving a pod for bonus points while fighting gravity and momentum. Jeremy went on to create the even more impressive Exile with Peter Irvin before tragically dying in an accident in 1992. He was somewhere between 16 and 18 when he wrote Thrust. You can play the original online .
I’ve got a BBC Master on the desk beside me and I still occasionally fire up Thrust on there along with some of the other classics. It’s one of those games I keep returning to along with Elite, Exile and Holed Out. I’ve now recreated three of these in different ways… the fourth is looking increasingly unavoidable.
Anyway. I guess I’d been thinking about Thrust as one morning recently I somewhat casually asked Claude Code to create it for me in the browser. I think I’d been reading the latest proclamations of capability from OpenAI and Anthropic and so I put together quite a comprehensive spec, gave it access to the original disassembled source code, screenshots, and said “go and recreate Thrust for me.”.
It created something for which the term slop would be too kind, it very vaguely resembled Thrust — it had the scanline stuff, sort of — but it was truly dreadful. It hadn’t even got gravity working right, the ship didn’t fall properly, the controls felt weird, and it was just… grim. In some ways its amazing that it created something that sort of worked and sort of looked like Thrust but it was not playable and nothing close to the elegance and beautfy of the real thing.
And that’s the thing about a game like Thrust. You could knock out something superficially similar pretty quickly — just run at the device frame rate, use standard delta-time physics, draw some caverns. But it would feel nothing like Thrust. The magic is in the specific timings, the weight of the ship, the way momentum builds. Particularly if you’ve played the original then those details are everything, and an AI working from a text description, and it turns out even the original source, can’t capture them.
But it got me curious. How did the original work? I find the tricks developers used to make this stuff work on the 8-bits fascinating, and it became a bit of an archaeology session. I quickly found this brilliant commented disassembly of the original source by Kieran Connell and found myself feeding it into Claude and asking questions.
This is where things got interesting. Not because AI wrote the code — the code itself isn’t complicated, it’s a 1986 game that ran in 32K of RAM — but because Claude turned out to be an extraordinary tool for interrogating 6502 assembly. I could feed in a block of disassembled source and ask “how does the level data work?” or “what’s the physics model doing here?” and get detailed, accurate explanations of what the original code was doing.
Now to be fair I was working from some commented disassembled source code but even given that it was able to extract information from both the comments and the assembly and come up with detailed descriptions of how the game worked. My sense is without the comments helping to focus it at the right areas it would have been much less useful - but even so, it made the job an awful lot simpler and more enjoyable. And yes it seems likely I’ll strip the comments from the code and see how well Claude does then.
While doing this I realised I could use the answers as the basis to recreate the original game and started asking Claude to create specifications for the various subsystems. Most of the specifications I generated can be found in a specs folder in the source code. I might write them up properly at some point but for now they give a good insight into the nuances in the original — there’s quite a lot going on, more than I’d realised. For example I’d never noticed that the turrets stop firing for a time if you hit the generator, and there are subtleties in their firing angles that only become apparent when you read the actual code.
The physics was one of the most interesting areas to dig into. Thrust uses Q7.8 fixed-point arithmetic — a common technique on 8-bit machines where you don’t have floating point hardware. The rotation system uses 32 steps with lookup tables for the force components.
But the really tricky part was timing. My first implementation used the correct constants from the disassembly — the same gravity, the same thrust values, the same drag — but the ship was far too fast and too agile. It didn’t have the right weight. The constants were identical to the original so it had to be a timing problem.
And it was. The original doesn’t run its physics at the BBC Micro’s 50 Hz VSync rate. The tick loop waits at least 3 centiseconds per frame, giving an effective rate of about 33.33 Hz. But it goes further than that: within each tick, physics updates are gated to only 6 active slots per 16-tick window. The core of the physics step looks like this:
/** The 6 active physics slots per 16-tick window */
private static readonly ACTIVE_SLOTS = new Set ([ 0 , 3 , 5 , 8 , 11 , 13 ]);
private tickStep ( input : ThrustInput ) : void {
const slot = this . tickCounter & 0x0F ;
this . tickCounter = ( this . tickCounter + 1 ) & 0xFF ;
// Rotation: 3 out of every 4 ticks, integer steps only
if (( slot & 0x03 ) !== 0 && input . rotate !== 0 ) {
s . angle = (( s . angle + input . rotate ) + 32 ) % 32 ;
}
const isActiveSlot = ThrustPhysics . ACTIVE_SLOTS . has ( slot );
// Force calculation — active slots only (6 of every 16 ticks)
if ( isActiveSlot ) {
s . forceY += this . gravity ;
if ( input . thrust ) {
s . forceY += ANGLE_Y [ angleIdx ] / ( 1 << this . massShift );
s . forceX += ANGLE_X [ angleIdx ] / ( 1 << this . massShift );
}
// Linear drag
s . forceX *= 1 - 1 / 64 ; // X: *= 63/64
s . forceY *= 1 - 1 / 256 ; // Y: *= 255/256
}
// Position integration — every tick
s . x += s . forceX ;
s . y += s . forceY ;
}
That gating gives an effective force/drag rate of roughly 12.5 Hz — gravity and thrust only apply on those 6 specific ticks, not every frame. Even rotation has its own gating: it skips every fourth tick. These aren’t arbitrary numbers; they’re the exact patterns from the 6502 source, and they define the feel of the game - if you pull them around things quickly start to feel off. The asymmetric drag is interesting too — much stronger on the X axis (63/64) than Y (255/256), which is why horizontal movement feels “stickier” than vertical.
Once I got these timings right — matching the original’s exact update cadence rather than just its constants — it felt perfect. You can switch between the BBC emulator and my version and the controls feel the same.
That’s what became genuinely interesting to me. Creating a Thrust-like game with normal physics is trivial - particularly using a coding AI. Recreating Thrust — with all its specific feel and weight — required understanding exactly how the original worked, right down to the timing of individual physics updates within the game loop.
The other area that required some real focus was the sound. When I did my TypeScript version of Elite I sampled the sounds directly from the emulator. That worked because Elite’s sounds are quite discrete — short, sharp effects. Though the beam and military laser effects are off. But in Thrust the engine is a continuous drone that responds to key presses, and the explosions have specific envelopes. I could have sampled them but it just felt… wrong.
So instead I decided to take a different approach and instead recreated the SN76489 sound chip (the same chip that was in the Sega Master System, as I learned while working on this) and the BBC MOS interface to it. The MOS — the BBC’s Machine Operating System — provided the interface through which games talked to the sound chip, using OSWORD calls and memory-mapped I/O.
A big help in this was the disassembled MOS code that Toby Nelson has put together, along with the BBC Advanced User Guide and the SN76489 chip specifications. I was able to feed all of this into Claude, generate a comprehensive spec for the sound system, and from that build an emulated sound system running in an AudioWorklet. The full BBC MOS envelope processor (OSWORD 7/8) drives the chip emulator on the audio thread.
The result? The sounds are identical. And it would have been really difficult to get that by any other means — it’s so timing-specific, so dependent on the exact envelope shapes and chip behaviour, that you really do need to emulate the actual hardware.
The sound system itself is quite elegant in how it layers. At the top level, playing a sound means sending the same OSWORD 7 parameters the original game used — channel, amplitude, pitch, duration:
const sounds = {
own_gun : { channel : 0x0012 , amplitude : 1 , pitch : 0x50 , duration : 2 },
explosion_1 : { channel : 0x0011 , amplitude : 2 , pitch : 0x96 , duration : 100 },
explosion_2 : { channel : 0x0010 , amplitude : 3 , pitch : 0x07 , duration : 100 },
hostile_gun : { channel : 0x0013 , amplitude : 4 , pitch : 0x1e , duration : 20 },
engine : { channel : 0x0010 , amplitude : - 10 , pitch : 0x05 , duration : 3 },
};
Those parameters feed into a MOS envelope processor that drives the SN76489 chip emulator, both running in an AudioWorklet on the audio thread. The chip emulator generates samples at the hardware level — tone channels with 10-bit period counters, a noise channel with a 15-bit linear feedback shift register, and a volume table with -2dB per step matching the original silicon:
generate ( out : Float32Array , offset : number , length : number , sampleRate : number ) : void {
const step = 250000.0 / sampleRate ; // chip clocks per audio sample
for ( let i = 0 ; i < length ; i ++ ) {
let sample = 0 ;
// Tone channels 0–2
for ( let ch = 0 ; ch < 3 ; ch ++ ) {
this . counter [ ch ] -= step ;
if ( this . counter [ ch ] <= 0 ) {
const p = this . period [ ch ] || 1024 ;
this . counter [ ch ] += p ;
this . polarity [ ch ] = - this . polarity [ ch ];
}
sample += this . polarity [ ch ] * this . vol [ ch ];
}
// Noise channel — 15-bit LFSR
// ...
sample += (( this . lfsr & 1 ) ? 1 : - 1 ) * this . vol [ 3 ];
out [ offset + i ] = sample ;
}
}
It’s emulation all the way down. The MOS ticks at 100 Hz on the audio thread, processing envelopes and updating the chip registers at the same rate the real BBC hardware would have and the resulting sounds are authentic.
The font, graphics, and level data I was able to extract from the disassembled source relatively easily. You can find a couple of tools in the source that do that — one decodes the level terrain data, another extracts the ship rotation sprites by emulating the BBC’s frame buffer and converting the output to PNGs.
I’d initially tried a vector approach for the ship — just drawing it mathematically at each rotation — but it looked rough at BBC resolution as it turned out the original had hard-coded sprites for each of the 32 rotation angles, hand-optimised by Jeremy Smith to look right at each position. Once I extracted and used those actual sprites, it looked correct.
The terrain rendering uses the original’s scanline-parity polygon fill — drawing every other line — which gives it that characteristic BBC Micro look. The internal resolution is 320×256, matching the original. My version has a slightly larger viewport — the original has a border aroun

[truncated]
