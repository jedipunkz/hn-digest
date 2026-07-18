---
source: "https://www.iandmacomber.com/blog/gopro-garmin-gemini-ride-recap/"
hn_url: "https://news.ycombinator.com/item?id=48957639"
title: "Show HN: ride-recap, teaching a LLM my taste to automate cycling highlights"
article_title: "Teaching LLMs Taste: How I Built an Automated Cycling Ride Recap with GoPro, Garmin, and Gemini — Ian Macomber"
author: "iandmacomber"
captured_at: "2026-07-18T12:52:22Z"
capture_tool: "hn-digest"
hn_id: 48957639
score: 1
comments: 0
posted_at: "2026-07-18T12:41:43Z"
tags:
  - hacker-news
  - translated
---

# Show HN: ride-recap, teaching a LLM my taste to automate cycling highlights

- HN: [48957639](https://news.ycombinator.com/item?id=48957639)
- Source: [www.iandmacomber.com](https://www.iandmacomber.com/blog/gopro-garmin-gemini-ride-recap/)
- Score: 1
- Comments: 0
- Posted: 2026-07-18T12:41:43Z

## Translation

タイトル: HN を表示: ライドの要約、LLM にサイクリングのハイライトを自動化する私の好みを教える
記事のタイトル: LLM に味を教える: GoPro、Garmin、Gemini を使用して自動サイクリング ライドを構築する方法の要約 — Ian Macomber
説明: または妻の言葉: 3 時間の退屈な GoPro 映像を 30 秒の退屈な GoPro 映像に変える方法
HN テキスト: TL;DR: 何時間もの生の GoPro フッテージと .fit ファイルを、ライド テレメトリが焼き付けられた 60 秒のハイライト リールに変換します。クリップのランキングとキュレーションと同様に、ライドの毎秒が Gemini-3.5-flash によってスキャンされます。料金は 1 回の乗車につき約 0.04 ドル、所要時間は 10 分です。ロングバージョン: ロードサイクリングは、ほぼ 10 年間、私の主な運動手段、社交のはけ口、セラピー、ワードローブの出費、そして性格特性となっています。私はほとんどの週末、マンハッタンから 9W までライドしています。ライドが終わる頃には、何時間もの GoPro 映像と、1 秒あたりの速度、パワー、心拍数、ケイデンス、GPS を含む 1 つの .fit ファイルが手に入ります。ウェストサイドハイウェイでシティバイクの後ろに3時間も立ち往生するのを見たくない人は絶対にいません。過去の映像を調べて、楽しい部分を特定し、思い出に残る物語をまとめるのは楽しいです。しかし、サイクリングにはすでに時間がかかるため、ライドごとにハイライト リール編集を手動で編集するのは簡単ではありません。そこで私は https://github.com/ianmacomber/ride-recap を構築してオープンソース化しました。私は 4 つの異なる情報源から魅力的な瞬間を特定します。
* .fit ファイル経由の Garmin テレメトリー (速度、HR、パワースパイク、スプリント、登坂)
* API 経由の Strava (人気のあるセグメント)
* Gemini ビジョンスキャン + 個々のフレームの評価
* (オプション) Streamlit アプリによる手動ラベル フュージョン ステップには LLM ナラティブ パスがあり、ライドのストーリーを最もよく伝える 20 個のクリップを選択し、「クロスソース合意」を促進します (人間のラベル + テレメトリ + Strava + Gemini のすべてが同意した場合)

クリップは興味深いものです）。貪欲な再ランキングと、すでに選択されているものに近すぎるクリップを避けるためのクラウディング ペナルティを備えています。振り返ってみれば明らかですが、Gemini が選択したクリップを見て、何を含めるべきか、何を含めるべきではないかについて自分の意見をしっかり持ち、その理由を十分に具体的にし、改善すべきことがこれ以上思いつかなくなるまで繰り返す以外に代わりはありません。自分自身に味覚がなければ、LLM の味覚を教えることはできません。

記事本文:
Ian Macomber について書く ← LLM の味を教えることを書く: GoPro、Garmin、Gemini を使って自動サイクリング ライドを構築した方法の要約
または、妻が言うには、「3 時間の退屈な GoPro 映像を、どうやって 30 秒の退屈な GoPro 映像に変えることができたのか」
ロードサイクリングは、ほぼ 10 年にわたり、私の主な運動手段、社交のはけ口、セラピー、ワードローブの出費、そして性格的特徴でした。私は最近ボストンからニューヨークに引っ越しましたが、私の最大の不安は、仕事前にドーバーまたはブルーヒルズに周回する午前 6 時 15 分の友人グループ (WhatsApp グループのタイトル: Boston WattsApp) がいなくなることでした。幸いなことに、ニューヨークのサイクリングシーンは素晴らしいです。私は多くの友情を築き、再燃させてきました。新しい WattsApp グループは、仕事の前にセントラルパークを周回したり、週末に 9W のロングライドをしたり、時折ベア マウンテンの頂上やモントークまで 130 マイル以上のライドに最適です。友達と 80 マイルを走り、ナイアックのボクサー ドーナツで冷たいビールとメープル ベーコンの丸太を食べることよりも 1 日を始めるのに最適な方法はありません。
多くのカジュアルサイクリストと同じように、私もライフスタイルを変える以外に、サイクリングで速くなるためなら何でもします（またはお金を使います）。私はすべての技術とすべての数字を持っています：スパンデックス、エアロバイク、カーボンホイール、ヘッドユニット、パワーメーター、3つの独立した心拍数モニター（Whoop、Garmin Watch、チェストストラップ）。今年の初めに、ハンドルバーに取り付けた GoPro をミックスに追加しました。
もちろん、ウェストサイドハイウェイでシティバイクの後ろに3時間も立ち往生するのを見たくない人は誰もいません。でも、お気に入りのセクションをつなぎ合わせるのは楽しいです。過去の映像を調べて、楽しい部分を特定し、思い出に残る物語をまとめます (Strava が私の個性であることを忘れないでください)。しかし、サイクリングにはすでに時間がかかるため、ライドごとにハイライト リール編集を手動で編集するのは簡単ではありませんでした。そこで私は r というパイプラインを構築し、オープンソース化しました。

ide-recap を使用して自動的に実行します。
3 か月、25 回のライド、45 回のコミット、1,000 マイル後には、何時間ものサイクリング映像と .fit ファイル (GPS、速度、心拍数、ケイデンス、パワーなどの時系列データ) を撮影して、完成したハイライト リールに変えることができます。クリップのランキングとキュレーションと同様に、ライドのすべての秒が gemini-3.5-flash によってスキャンされます。料金は 1 回の乗車につき約 0.04 ドル、所要時間は 10 分です。最終的なハイライトは次のとおりです。
このシステムには 6 つのステップと 1 つのコマンドがあります:ride-recap process data/raw/<YYYY-MM-DD>/
取り込み: GoPro データ (.mp4 + .LRV) と Garmin .fit ファイルを乗り物フォルダーにプルします
検出 : 4 つの候補ソースが乗り物の興味深い部分を特定します
Gemini ビジョン: 視覚的に魅力的なセグメントのための 2 パス フレーム スキャン (これについては後で詳しく説明します)
テレメトリー: 速度、登り + 下り、パワー + 心拍数のスパイク
Strava API: ルート上の人気のセグメント (例: State Line Sprint には 343 個の星があります)、これはニューヨークのサイクリストなら誰でも知っている舗装です
(オプション) 手動ラベル: Streamlit アプリ。時々、瞬間を「必須」として手動でフラグを立てます (たとえば、鹿を見かけた場合など)
Fuse : 候補クリップを生成し、それらをランク付けし、近くの瞬間にペナルティを与えます (これについては後で詳しく説明します)
(オプション) レビュー: Streamlit アプリ (候補のランキングまたはクリップの評価を上書きする場合)
作成 : 選択したクリップにライド テレメトリ オーバーレイを書き込みます。
出力: YouTube の場合は、highlight_landscape.mp4 (60 秒、16:9)、Instagram の場合は、highlight_portrait.mp4 (30 秒、9:16)
これは、ステップ 5 の焼き付けオーバーレイの事前/事後です。左側にライブ テレメトリ、右側に GTA スタイルのミニマップとルート トレース、そして下部に沿って標高プロファイルと走行の進行状況を追加したことがわかります。
サイクリングのビジュアルが魅力的な理由は何ですか?
多くの人がポスト AI 時代における Taste™ の価値について意見を述べています。 4か月前なら、私はこう答えていただろう

ポッター・スチュワート判事がポルノについて説明したように、「見ればそれがわかる」のです。このパイプラインを通じて 180,430 秒 + フレームを処理しました。 LLM にテイストを教えることについての私の最も強い意見は次のとおりです。テイストは反復回数によって発達します。事前にテイストを指定することはできません。
まず、「視覚的な関心を 1 ～ 5 で評価する」というプロンプトを含むすべてのフレームを Gemini に送信しました。しかし、ニューヨークをドライブするのはさまざまな理由から興味深いものになるでしょう。ハドソンヤードからのコンクリートジャングルの眺めやピークスキルからのハドソン川の眺め、リバーロードの登りやベアマウンテンの下り、9Wでペースラインに落ちそうになったり、9番街で逆走する電動自転車に轢かれそうになったり。
これは私が現在導入しているものの例です。すべての乗車の毎秒、同じバージョンのプロンプトが表示されます。 Gemini に 6 つの 848×480 プロキシ フレームと 1 行のテレメトリを送信すると、1 つの JSON オブジェクトが返されます。以下の 7 月 12 日のライドからの 3 つの通話をめくってください。ハドソンヤードを駆け抜けるスプリント、GW ブリッジの下からの美しい景色、そして黒い SUV の後ろで 2 秒間立ち往生する様子です。私が開発したセンスは、自分が同意できない JSON 出力に頭をぶつけてきたことから生まれました。
プロンプトを改善する最善の方法は、間違いを修正することです。それにはグラウンド トゥルースが必要なので、Streamlit アプリを作成し、セントラル パークの乗り物に手動でラベルを付けました。テレメトリグラフと並行してライド全体をフレームごとにレビューし、興味深いまたは退屈だと感じた約 30 の瞬間とその理由の説明をランク付けしました。
ジェミニは 200W を超える出力に感銘を受けました。私はタデイ・ポガチャルではありませんが、200w はそれほど印象的ではありません。
解決策: アスリート固有の FTP/パワー ゾーンがシステム プロンプトに挿入されました。誰がペダルを踏んでいるのかが分からなければ、200W であっても意味がありません。
power_zones = (
f "アスリートのパワーゾーン (FTP= { ftp } W): \n "

f " Z1 アクティブ回復 < {int (ftp * 0.55 ) } W | Z2 耐久 < {int (ftp * 0.75 ) } W | "
f "Z3 テンポ < {int (ftp * 0.90 ) } W \n "
f " Z4 閾値 < {int (ftp * 1.05 ) } W | Z5 VO2max < {int (ftp * 1.20 ) } W | "
f "Z6 無酸素運動 < {int (ftp * 1.50 ) } W | Z7 神経筋運動 {int (ftp * 1.50 ) } W+ \n "
「本当にハードな努力として認められるのは Z5+ だけです。Z3 ～ Z4 は中程度です。Z1 ～ Z2 は簡単です。」
)
ジェミニはトンネルのクリップや車の後ろに閉じ込められるのが大好きでした
解決策: ネガティブな例 — 4 月 14 日の乗車時の実際の交差点。永久的な警告としてプロンプトに表示されます。
赤信号で配送トラックの後ろに止まり、決して動かない：
{ 「光」 : 3 、 「構図」 : 2 、 「動き」 : 1 、 「風景」 : 1 、 「被写体」 : 1 、
"peak_offset" : 0.5 、 "clip_type" : "transition" 、 "crop_x" : 50 、
"reason" : "赤信号で配送トラックの後ろで立ち往生" }
私はビジュアルを識別するのが好きでした。ブルックリン橋、州境、9W の標識など、サイクリストなら誰でも一目瞭然です。
解決策: 肯定的な例 — 10 がどのようなものなのか、何がスコアを獲得するのかを示す
夕暮れ時にブルックリン橋を通過し、最後に最も大きな橋が見えます。
{ 「光」 : 9 、 「構図」 : 8 、 「動き」 : 5 、 「風景」 : 10 、 「被写体」 : 8 、
"peak_offset" : 0.8 、 "clip_type" : "landmark" 、 "crop_x" : 70 、
"reason" : "夕焼けのブルックリン橋が近づいてきました、友人も一緒です" }
Gemini はよく隣り合った 2 つのクリップを選択しました
解決策: 近くの瞬間にペナルティを課し、乗車中に貪欲な選択を強制します (詳細は後述)
ジェミニは私が時速0マイルで行こうとしていた場所をクリップしました
解決策: 時速 5 マイル未満のクリップをフィルターします。
振り返ってみれば明らかですが、Gemini が選択したクリップを見て、何を含めるべきか、何を含めるべきではないかについて自分の意見をしっかり持ち、その理由を十分に具体的にし、改善すべきことがこれ以上思いつかなくなるまで繰り返す以外に代わりはありません。

Hamel 氏が書いているように、裏付けとなるデータを収集し、ワークフローをより小さな単位に分割することで、自動採点が容易になります。自分自身に味覚がなければ、LLM の味覚を教えることはできません。
以下は、私が実行した実験からの出力の一部です。実際の乗車から数枚の静止画を取得し、プロンプトとモデルのあらゆる組み合わせを通じてスコアを付けました。v2 から v10 、現在の 3.5 フラッシュ、他の 2 つの Gemini モデル、および Opus 4.6 です。プロンプトは、より強力な基礎モデルよりもはるかに重要であることがわかります。
これまでにランキングのプロンプトを 9 回繰り返しました。これはリポジトリ内で最も重要なファイルであり、私はそのように扱います。プロンプトは、新しいバージョンでどのような障害モードが修正されるかを説明する必須の根拠でバージョン管理されています。このプロンプト レジストリは、サイクリング ビジュアルに対する LLM の好みがどのように進化したかの歴史的記録です。
このコードは説明するよりも直接読む価値があります。プロンプトがどのように増加したかを確認してください。v2 は 20 行で 2 つの軸上のスコアですが、v10 は 244 行です。
あなたは、GoPro ハンドルバー カム フレームをレビューするサイクリング ビデオ編集者です。
あなたの仕事は、ハイライト リールに含める価値のあるフレームにフラグを立てることです。
ハンドルバーカム映像の視覚的なヒント:
強い: 舗装上のモーション ブラー、傾いた角度のコーナー、近くのライダーの車輪、
ハイコントラストの光（日陰/トンネル/反射）、ランドマーク（スカイライン/橋/景色）、
スプリントの跳ね返り、都市の要素が駆け抜けていく、気象劇、危機一髪、予期せぬ出来事。
弱：近くに何もなく安定したケイデンス、平らな曇りの光、信号で停止。
各フレームを 2 つの次元でスコア付けします。
ビジュアル (1-5): フレームはどのように見えますか?光、構図、ランドマーク。
アクション (1-5): どれくらいのことが起こっていますか?努力、近さ、スピード、出来事。
5 = 編集を主導します。 4 = 有力な候補。 3 = フィラー。 2 = 目立たない。 1 = 死亡。
フレームごとに 1 つの Clip_type を割り当てます。
」

風景" | "アクション" | "都市" | "登り" | "下り" | "グループライディング" | "事件" | "ランドマーク" | "トランジション"
テレメトリ コンテキストはバッチごとに提供されます。これを使用して曖昧さを解消します (350W のプレーン フレーム = 含める価値のある労力)。
ビジュアル < {min_visual} かつアクション < {min_action} の両方のフレームを省略します。
JSON 配列のみを返し、散文は返しません。
[{{"frame_index": N、"visual": 1-5、"action": 1-5、"clip_type": "..."、"reason": "最大 10 単語"}}]
該当するものがない場合は [] を返します。私はもともと「ハイコントラストライト（日陰/トンネル/反射）」をプラスとしてリストしていました。これは間違いでした。ジェミニはたくさんのトンネルを選びました。あなたは、Instagram リール用のクリップをキュレーションするロードサイクリングビデオ編集者です —
Pas Normal、Pedal Mafia、GCN を思い浮かべてください。ライダーはニューヨークのロードサイクリストです
そしてボストンは、週末や仕事前に友人とサイクリングします。あなたは
1 回の乗車での GoPro ハンドルバー カメラの映像を見ています。
1 ～ 10 ～ 20 秒間で {n_frames} 個の等間隔のサンプル フレームが表示されます。
クリップ。クリップ全体を評価します。これらは、内容の層別サンプルです。
クリップは最初から最後まで表示されます。フレーム 0 が最も早いフレームです
{n_frames} - 1 が最新です。
═============================= ═============================
ザ・バー
═══════
[切り捨てられた]
候補生成プロセスでは、潜在的な 3 秒のクリップをすべてランク付けします (1 時間あたり約 1000)。 60 秒のリールを満たす最も簡単なアプローチは、上位 20 位の候補者を選択することです。多くの場合、上位にランクされた候補が互いに隣接しているため、編集が非常に繰り返し行われることになります。ベア マウンテンの下りは約 7 分で、美しい景色を望む高速下りのクリップが約 100 個作成されます。

Hudson はすべて、クリップごとのスコアが高くなりますが、独立して興味深いものではありません。
LLM に 1 つのクリップを正確にスコア付けするよう教えるのは困難ですが、少なくとも各クリップは単独で評価できます。候補者の選択はさらに困難です。優れた編集者は、優れたストーリーを伝えるためにどのようにして映像を選択するのでしょうか?私の観点から、ニューヨーク市のサイクリング編集の優れた特徴は次のとおりです。
スタートとフィニッシュで雰囲気を決めましょう。私の場合、通常はゆっくりとペダルを漕いでウェストサイド・ハイウェイを上り、ハドソン・ヤードの高層ビル群を通って下ります。
クリップを時系列順に並べ、ライド全体で均一にサンプリングします（無理のない範囲で）
ランドマーク、特にサイクリストに愛されるランドマークを含める
1 分以内のクリップは繰り返している可能性があります
時速 5 マイル未満のクリップはおそらく退屈です
友達が登場するクリップのほうが面白い
多様性を追求: 最も急な登り、最速の下り、最も力強いスプリント、パリセーズの森から輝く太陽、ニューヨークのコンクリートジャングル、州境の標識、9W マーケットなど
スコアリングと同様に、候補融合ステップを 6 回再構築し、各バージョンで以前の反復から何かを修正しました。現在の選択プロセスは次のように機能します。
「含める必要がある」手動ラベルが最初に選択されます
Gemini LLM ナラティブ パスは、ライドのストーリーを最もよく伝える 20 個のクリップを選択し、「クロスソース合意」を促進します (人間のラベル + テレメトリ + Strava + Gemini がすべて、

[切り捨てられた]

## Original Extract

Or as my wife says: How I managed to turn 3 hours of boring GoPro footage into 30 seconds of boring GoPro footage

TL;DR: Turn hours of raw GoPro footage + a .fit file into a 60-second highlight reel with ride telemetry burned in. Every second of the ride is scanned by gemini-3.5-flash, as is the clip ranking + curation. The whole thing costs about $0.04 per ride and takes 10 minutes. Longer version: Road cycling has been my primary form of exercise, social outlet, therapy, wardrobe expense, and personality trait for almost a decade. I ride most weekends, usually out of Manhattan and up 9W. By the end of a ride I have hours of GoPro footage and one .fit file with per-second speed, power, heart rate, cadence, and GPS. Absolutely no one wants to watch 3 hours of being stuck behind Citibikes on West Side Highway. It's fun to look through past footage, identify the fun parts, and put together a narrative to remember. But since cycling is already time consuming, manually editing a highlight reel edit per ride is a nonstarter. So I built and open-sourced https://github.com/ianmacomber/ride-recap . I identify compelling moments from four separate sources:
* Garmin telemetry via .fit file (speed, HR, power spikes, sprints, climbs)
* Strava via API (popular segments)
* Gemini vision scan + rating of individual frames
* (optional) hand-labels via Streamlit app The fusion step has a LLM narrative pass pick 20 clips to best tell the story of the ride, boosting “cross-source agreements” (if a human label + telemetry + Strava + Gemini all agree that a clip is interesting), with greedy re-ranking and a crowding penalty to avoid clips too close to something already selected. Obvious in retrospect, but there’s no substitute other than looking at the clips Gemini selects, being highly opinionated about what should / should not be included, being specific enough about why, and repeating until you can’t think of anything more to improve. You cannot teach an LLM taste if you do not have taste yourself.

Ian Macomber Writing About ← Writing Teaching LLMs Taste: How I Built an Automated Cycling Ride Recap with GoPro, Garmin, and Gemini
Or as my wife says: How I managed to turn 3 hours of boring GoPro footage into 30 seconds of boring GoPro footage
Road cycling has been my primary form of exercise, social outlet, therapy, wardrobe expense, and personality trait for almost a decade. I recently moved from Boston to New York, and my biggest anxiety was losing my 6:15am friend group that lapped to Dover or Blue Hills before work (WhatsApp group title: Boston WattsApp). Fortunately, the NYC cycling scene is awesome. I’ve made and rekindled many friendships, I have a new WattsApp group that’s great for Central Park laps before work, longer rides up 9W on the weekends, and the occasional Bear Mountain summit or 130+ mile ride to Montauk . There is no better way to start a day than 50 miles with friends and a cold brew + maple bacon log from Boxer Donut in Nyack.
Like many casual cyclists, I would do (or spend) anything to get faster at cycling, other than change my lifestyle in any way. I have all the tech and all the numbers: spandex, aero bike , carbon wheels, head unit, power meter, three separate heart rate monitors (Whoop, Garmin Watch, chest strap). Earlier this year I added a handlebar-mounted GoPro to the mix.
Of course, absolutely no one wants to watch 3 hours of being stuck behind Citibikes on West Side Highway. But I enjoy stitching together my favorite sections. I look through past footage, identify the fun parts, and put together a narrative to remember (remember, Strava is my personality). But since cycling is already time consuming, manually editing a highlight reel edit per ride was a nonstarter. So I built and open-sourced a pipeline called ride-recap to do it automatically.
Three months, 25 rides, 45 commits, and 1000 miles later, I can take hours of cycling footage plus a .fit file (time-series data with GPS, speed, heart rate, cadence, power, etc.) and turn it into a finished highlight reel. Every single second of the ride is scanned by gemini-3.5-flash , as is the clip ranking + curation. The whole thing costs about $0.04 per ride and takes 10 minutes. Here’s a finalized highlight:
The system has six steps and one command: ride-recap process data/raw/<YYYY-MM-DD>/
Ingest : pull GoPro data ( .mp4 + .LRV ) and Garmin .fit file into a ride folder
Detect : four candidate sources identify interesting parts of the ride
Gemini vision: a two-pass frame scan for visually compelling segments (more on this later)
Telemetry: speed, climbs + descents, power + heart rate spikes
Strava API: popular segments on the route (e.g. State Line Sprint has 343 stars), this is the pavement every NYC cyclist knows
(Optional) Manual Labels: Streamlit app, occasionally I’ll hand-flag a moment as a “must include” (say, if I see a deer )
Fuse : Generate candidate clips, rank them, penalize nearby moments (more on this later)
(Optional) Review : Streamlit app, if I want to override the candidate ranking or rate clips
Compose : Burn a ride telemetry overlay onto the selected clips
Output : highlight_landscape.mp4 (60s, 16:9) for YouTube, highlight_portrait.mp4 (30s, 9:16) for Instagram
Here’s a pre/post for the step five burned overlay: you can see I’ve added live telemetry on the left, a GTA style mini-map + route trace on the right, and elevation profile + ride progress along the bottom.
What Makes Cycling Visuals Compelling?
Many have opined on the value of Taste™ in the post-AI era. Four months ago, I would have answered like Justice Potter Stewart describing pornography: I know it when I see it . I’ve now processed 180,430 seconds + frames through this pipeline. Here’s my strongest opinion on teaching taste to a LLM: taste develops through iteration count, you cannot specify taste in advance.
I started by sending every frame to Gemini with a prompt to “rate visual interest 1-5”. But a ride around New York could be interesting for so many reasons! A concrete jungle view from Hudson Yards or a Hudson River view from Peekskill, a climb up River Road or a descent down Bear Mountain, almost getting dropped in a paceline on 9W or almost getting hit by an e-bike going the wrong way on 9th Avenue.
Here’s an example of what I have in place now: every second of every ride goes through the same versioned prompt. I send Gemini six 848×480 proxy frames plus one line of telemetry, and I get one JSON object back. Flip through three calls from the July 12th ride below: a sprint through Hudson Yards, a beautiful view from under the GW Bridge, and two seconds stuck behind a black SUV. Any taste I developed came from me banging my head against JSON outputs I disagreed with.
The best way to improve a prompt is to fix its mistakes. That requires ground truth, so I created a Streamlit app and manually labeled a Central Park ride . I reviewed the entire ride frame-by-frame alongside telemetry graphs, and ranked around 30 moments I found interesting or boring, plus a description of why.
Gemini was impressed by any power output above 200W. I’m no Tadej Pogacar, but 200w isn’t that impressive.
Solution: athlete specific FTP/power zones injected into the system prompt. 200W means nothing without knowing who’s pushing the pedals.
power_zones = (
f "Athlete power zones (FTP= { ftp } W): \n "
f " Z1 Active Recovery < {int (ftp * 0.55 ) } W | Z2 Endurance < {int (ftp * 0.75 ) } W | "
f "Z3 Tempo < {int (ftp * 0.90 ) } W \n "
f " Z4 Threshold < {int (ftp * 1.05 ) } W | Z5 VO2max < {int (ftp * 1.20 ) } W | "
f "Z6 Anaerobic < {int (ftp * 1.50 ) } W | Z7 Neuromuscular {int (ftp * 1.50 ) } W+ \n "
"Only Z5+ qualifies as genuinely hard effort. Z3-Z4 is moderate. Z1-Z2 is easy."
)
Gemini loved clips of tunnels, and being stuck behind cars
Solution: negative examples — a real intersection from the April 14th ride, immortalized in the prompt as a permanent warning
Stopped behind a delivery truck at a red light, never moves:
{ "light" : 3 , "composition" : 2 , "motion" : 1 , "scenery" : 1 , "subject" : 1 ,
"peak_offset" : 0.5 , "clip_type" : "transition" , "crop_x" : 50 ,
"reason" : "Stuck behind delivery truck at red light" }
I liked identifying visuals: Brooklyn Bridge, State Line and 9W signs that are obvious to any cyclist
Solution: positive examples — show it what a 10 looks like, and what earns the score
Cruising past Brooklyn Bridge at sunset, bridge biggest at end:
{ "light" : 9 , "composition" : 8 , "motion" : 5 , "scenery" : 10 , "subject" : 8 ,
"peak_offset" : 0.8 , "clip_type" : "landmark" , "crop_x" : 70 ,
"reason" : "Sunset Brooklyn Bridge approach, friend alongside" }
Gemini often picked two clips right next to one another
Solution: penalize nearby moments, enforce greedy selection across the ride (more below)
Gemini picked clips where I was going 0mph
Solution: filter clips below 5mph
Obvious in retrospect, but there’s no substitute other than looking at the clips Gemini selects, being highly opinionated about what should / should not be included, being specific enough about why, and repeating until you can’t think of anything more to improve. As Hamel writes : gathering supporting data and breaking down workflows into smaller units makes automated grading easier. You cannot teach an LLM taste if you do not have taste yourself.
Here are some outputs from an experiment I ran: I took a handful of stills from a real ride and scored them through every combination of prompts and models: from v2 to v10 , the current 3.5-flash , two other Gemini models, and Opus 4.6. You can see that the prompt matters much more than more powerful foundation models.
I’ve iterated nine times on the ranking prompt so far. It is the most important file in the repo, and I treat it that way: the prompts are version controlled with a mandatory rationale explaining what failure mode the new version fixes. This prompt registry is the historical record of how an LLM’s taste for cycling visuals evolves.
This code is worth reading directly instead of describing. Check out how the prompt has grown: v2 is 20 lines and scores on two axes, v10 is 244 lines.
You are a cycling video editor reviewing GoPro handlebar-cam frames.
Your job is to flag frames worth including in a highlight reel.
VISUAL CUES for handlebar-cam footage:
Strong: motion blur on pavement, corners with lean angle, rider wheels nearby,
high-contrast light (shade/tunnels/reflections), landmarks (skylines/bridges/vistas),
sprint bounce, urban elements rushing past, weather drama, close calls, unexpected events.
Weak: steady cadence with nothing nearby, flat overcast light, stopped at lights.
Score each frame on TWO dimensions:
VISUAL (1-5): How does the frame look? Light, composition, landmarks.
ACTION (1-5): How much is happening? Effort, proximity, speed, events.
5 = lead the edit. 4 = strong candidate. 3 = filler. 2 = unremarkable. 1 = dead.
Assign ONE clip_type per frame:
"scenery" | "action" | "urban" | "climb" | "descent" | "group_riding" | "incident" | "landmark" | "transition"
Telemetry context is provided per batch — use it to disambiguate (plain frame at 350W = effort worth including).
Omit frames where BOTH visual < {min_visual} AND action < {min_action}.
Return ONLY a JSON array, no prose:
[{{"frame_index": N, "visual": 1-5, "action": 1-5, "clip_type": "...", "reason": "max 10 words"}}]
Return [] if nothing qualifies. I originally had "High-contrast light (shade/tunnels/reflections)" listed as a plus. This was a mistake. Gemini picked a lot of tunnels. You are a road cycling video editor curating clips for an Instagram Reel —
think Pas Normal, Pedal Mafia, GCN. The rider is a road cyclist in NYC
and Boston, riding with friends on weekends and before work. You're
looking at GoPro handlebar-cam footage from a single ride.
You will see {n_frames} evenly-spaced sample frames from one ~10-20-second
clip. Rate the clip AS A WHOLE — these are a stratified sample of what
the clip shows from start to finish. Frame 0 is the earliest, frame
{n_frames} - 1 is the latest.
═══════════════════════════════════════════════════════════════════
THE BAR
═══════
[truncated]
The candidate generation process ranks every single potential three second clip (around ~1000 per hour). The simplest approach to fill a 60 second reel is to select the top 20 ranked candidates. That results in an extremely repetitive edit, because the top ranked candidates are often right next to one another. A Bear Mountain descent is ~7 minutes, which results in ~100 different clips of a fast descent with beautiful views of the Hudson, all of which score highly on a per-clip basis, but are not independently interesting.
Teaching a LLM to accurately score one clip is a challenge, but at least each clip can be evaluated on its own. Candidate selection is even harder: How DOES a great editor select the footage to tell a great story? From my perspective, here are traits of a great NYC cycling edit:
Set the tone with the start and finish: for me, usually slow pedaling up West Side Highway, and back down through the skyscrapers of Hudson Yards
Order the clips chronologically, sampling uniformly across the ride (within reason)
Include landmarks, especially landmarks beloved by cyclists
Clips within a minute of one another are probably repetitive
Clips going <5mph are probably boring
Clips with friends in them are more interesting
Shoot for diversity: include the steepest climb, the fastest descent, the most powerful sprint, the sun shining through the forest in the Palisades, the NYC concrete jungle, The State Line sign, 9W Market
Similar to scoring, I’ve rebuilt the candidate fusion step six times, with each version fixing something from a prior iteration. The current selection process works like this:
“Must include” manual labels are picked first
A Gemini LLM narrative pass picks 20 clips to best tell the story of the ride, boosting “cross-source agreements” (if a human label + telemetry + Strava + Gemini all agree that a

[truncated]
