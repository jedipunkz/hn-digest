---
source: "https://soullessmusic.com/"
hn_url: "https://news.ycombinator.com/item?id=48882955"
title: "Soulless – List of AI Artists Hiding on Spotify"
article_title: "Soulless: List of AI Artists Hiding on Spotify"
author: "ChrisArchitect"
captured_at: "2026-07-12T17:51:30Z"
capture_tool: "hn-digest"
hn_id: 48882955
score: 1
comments: 0
posted_at: "2026-07-12T17:46:57Z"
tags:
  - hacker-news
  - translated
---

# Soulless – List of AI Artists Hiding on Spotify

- HN: [48882955](https://news.ycombinator.com/item?id=48882955)
- Source: [soullessmusic.com](https://soullessmusic.com/)
- Score: 1
- Comments: 0
- Posted: 2026-07-12T17:46:57Z

## Translation

タイトル: Soulless – Spotify に隠れている AI アーティストのリスト
記事タイトル: Soulless: Spotify に隠れている AI アーティストのリスト
説明: Soulless は、Spotify 上の AI アーティストと AI が生成した音楽の厳選されたリストです。偽のアーティストを閲覧し、本物の毎月のリスナーを確認し、マシンに反対票を投じます。

記事本文:
Soulless: Spotify に隠れている AI アーティストのリスト SOULLESS MUSIC Analyze ホーム
Spotify に隠れている AI アーティスト。バンドもスタジオも魂も存在せず、ただマシンとメロディーがあるだけだ。
AI 音楽は本物のアーティストから盗んでいます。
曲を分析するのは人間ですか、それとも AI ですか?オープンソース 気に入らないマシンに反対票を投じてください。彼らの収入を見る→
AI が生成した下手なアーティストを公開する (232)
コメント「ベルベット・サンダウン」を呼び出す
シエナ・ローズを呼ぶコメント
コメント Breaking Rustを呼び出す
Soulless 用の AI 音楽検出器をどのように構築したか
誰のスノが作ったの？バージョンの帰属
検出器の出荷: Mac mini が生産中
偽陽性に対するモデルのキャリブレーション
私たちは本物の歌を AI と呼んだ: 検出器に自己疑念を教える
AI の帰属: Suno v4 ホールと Udio
自分の Spotify をスキャンします (そして Discover Weekly が立ち入り禁止になっている理由)
Spotify または Deezer のリンクを貼り付けると、検出器はそれを「人間」、または「Suno、Udio などが作成したもの」と呼びます。
曲を分析 → AI音楽風景
Soulless / ai-music-classifier
当社独自の検出器: SONICS スペクトログラム モデル、lofcz ボコーダー フェイクプリント、およびメタデータ スキャンのアンサンブルで、30 秒のストア プレビューで実行されます。 MITライセンス取得済み。
音楽基盤モデル (オーディオの BERT など)。検出器自体ではありませんが、その埋め込みは、最新の検出器がその上に分類器をトレーニングするものです。
内部で SH Labs アナライザーを使用して AI トラックにフラグを付ける Spotify スキャナー。
SubmitHub の検出器は、ラベル付きの実際の人間と AI による送信の膨大なストリームに基づいてトレーニングされています。強力ですが、独自のライセンスがあり、商用利用のみが許可されています。
ほとんどのオープン検出器が構築する SpecTTTra スペクトログラム トランスフォーマー モデル (ICLR 2025) と、大規模なラベル付き Suno/Udio データセット。当社のスペクトル検出のバックボーン。
ボコーダーの「フェイクプリント」検出器: ニューラル ボコーダーが残した周期的な 1 ～ 8 kHz のアーティファクトを検出します。 MITライセンス

ENSED、そしてアンサンブルの後半。
Spotify プレビュー上でブラウザ内でローカルに SONICS モデルを実行する Chrome 拡張機能。パブリックベータ版は誤検知の懸念を受けて廃止されました。ソースは起動したままです。
22 のジェネレーターにわたるコーデック不変の一般化を主張するフォレンジック アーティファクト検出器。 ONNX 推論ビルドのみが CC-BY-NC に基づいてリリースされており、特許出願中です。
約 97,000 曲 (およそ半分が本物、半分が Suno/Udio の AI)。私たちが実行する事前トレーニング済みモデルのソース。
161 ジャンルにわたる 106k のクリエイティブ コモンズ トラック。トレーニングやベンチマーク用の、クリーンで正規にライセンスを取得したヒューマン ミュージックのソース。
美しい注釈が付いた Udio オーディオ スニペット。Apache-2.0 ライセンスが付与されており、私たちが見つけた中で最もライセンスが許可された Udio オーディオです。
最も広く使用されているテキストを歌に変換するジェネレーター。ユーザーが検出器に貼り付けるほとんどのトラックはここから来ています。
もう 1 つの主要なテキストを歌に変換するジェネレーターで、高音質で知られています。
無料の音楽アーカイブ データセット
AIのトレーニングに音楽が密かに使用された2,100万人のアーティストの中にバッド・バニーとテイラー・スウィフトが含まれる
The Atlantic は、AI のトレーニングに使用される音楽の検索可能なデータベースを作成しました
AI が生成した音楽にマッシュされた数百万の曲
あなたのお気に入りの音楽アーティストが AI によって生成されたものであるかどうかを見分ける方法
SlopTracker: AI によって生成されたアーティストを Spotify で公開する
Soulless 用の AI 音楽検出器をどのように構築したか
誰のスノが作ったの？バージョンの帰属
雪の街の夜 魔法のクリスマス光と居心地の良いホリデーの雰囲気
コメント 時代を超えたサウンドを呼び起こす
コメント Breaking Rustを呼び出す
ザニア・モネを呼び出すコメント
ザニア・モネを呼び出すコメント
どうやって知る必要があったのでしょうか？ (人間バージョン)
もし最後が最後なら
それは愛ではなかった、それは生存だった
間違った方向 (シングル・バージョン)
未完成の未来 (Single Version)
スリー・セカンズ・トゥ・グッバイ (ライヴ)
ドント・セイ・グッバイ・マイ・ラブ (ライヴ)
時間がない

クソ野郎
選手に祝福の雨が降り注ぐ
私はプライムがあるとは信じていません
許してやる、ただあなたに振り回されたくないだけ
今日はマザーファッカーズを読むのをやめます
私は人よりも犬を信頼します
DJ にセックスしようとしていると伝えてください
今夜は外出したくない
愚かなたわごとに対処するには年をとりすぎています
人生がレモンを与えたら、テキーラと混ぜてください
ビッチをゲットするのにお金なんて必要なかった
寒くてくわがない
今日はダチを平手打ちしたい気分だ
でたらめを手放す
新年、私は同じクソ野郎です
今日はたわごとのことは心配しない
新年、私は同じクソ野郎です
寝るためにお尻を連れて行こうとしています
ジョージア州のダンサーからの賢明な言葉
誰もがクソみたいな意見を持っている
このたわごとはすべてうまくいく
でたらめを手放す
Niggas Trynna 私の祝福をブロックしてください
ガール、あと 1 パーレーです
教師よりも多くのお金を稼ぐ
止められないような気分
彼女の経歴には悪魔とのつながりがある
久しぶりのセックス
PH バランスをチェックする必要があります、ベイビー
もう食べ物には興味がない
ハスラーの魂 1巻。モチベーションを高める 70 年代の音楽
今日はダチを平手打ちしたい気分だ
朝起きて素晴らしい小便をする
ママ、私は鍬と結婚したと思う
コカインを嗅いでいるサンタクロースを捕まえた
みんなにとって私がすべてである必要はない
生き残ったことを証明する必要はない
今夜誰かがブルースを演奏している
ルーズ・コントロール (リアーナ・ウラニス・リミックス)
オーディナリー (リアーナ・ウラニス・リミックス)
ダイ・ホワイト・ア・スマイル (リアーナ・ウラニス・リミックス)
伝統的なクリスマスソング Vol. 1 (ソウル・ブルース・ヴァージョン)
ジングルベル (ソウル・ブルース・ヴァージョン)
エコーズに浮かぶ（人間バージョン）

## Original Extract

Soulless is a curated list of AI artists and AI-generated music on Spotify. Browse fake artists, see real monthly listeners, and downvote the machines.

Soulless: List of AI Artists Hiding on Spotify SOULLESS MUSIC Analyze Home
AI artists hiding on Spotify. No bands, no studios, no soul, just machines and melody.
AI music is stealing from real artists.
Analyze a song: human or AI? Open source Downvote any machine you dislike. See what they earn →
Expose the slop AI generated artists (232)
Comment Calling out The Velvet Sundown
Comment Calling out Sienna Rose
Comment Calling out Breaking Rust
How we built an AI-music detector for Soulless
Which Suno made it? Version attribution
Shipping the detector: a Mac mini in production
Calibrating the model for false positives
We called a real song AI: teaching the detector self-doubt
AI attribution: Suno v4 hole and Udio
Scan your own Spotify (and why Discover Weekly is off-limits)
Paste a Spotify or Deezer link and our detector calls it: human, or made by Suno, Udio and the rest.
Analyze a song → AI music landscape
Soulless / ai-music-classifier
Our own detector: an ensemble of the SONICS spectrogram models, the lofcz vocoder fakeprint, and a metadata scan, running on 30-second store previews. MIT licensed.
A music foundation model (like BERT for audio). Not a detector itself, but its embeddings are what modern detectors train a classifier on top of.
A Spotify scanner that flags AI tracks using the SH Labs analyzer under the hood.
SubmitHub's detector, trained on their huge stream of real, labelled human and AI submissions. Strong, but proprietary and licensed for commercial use only.
The SpecTTTra spectrogram-transformer models (ICLR 2025) that most open detectors build on, plus a large labelled Suno/Udio dataset. The backbone of our spectral detection.
A vocoder "fakeprint" detector: it finds the periodic 1-8 kHz artefacts neural vocoders leave behind. MIT licensed, and the second half of our ensemble.
A Chrome extension that runs the SONICS model locally in-browser over Spotify previews. Public beta was retired after false-positive concerns; source stays up.
A forensic-artefact detector claiming codec-invariant generalisation across 22 generators. Only an ONNX inference build is released, under CC-BY-NC with a patent pending.
~97k songs (roughly half real, half AI from Suno/Udio). The source of the pretrained models we run.
106k Creative-Commons tracks across 161 genres. The clean, genuinely-licensed source of human music for training and benchmarks.
Udio audio snippets with aesthetic annotations, Apache-2.0 licensed, the most permissively-licensed Udio audio we found.
The most widely-used text-to-song generator. Most tracks people paste into detectors come from here.
The other major text-to-song generator, known for high audio quality.
The Free Music Archive Dataset
Bad Bunny, Taylor Swift Among 21 Million Artists Whose Music Was Secretly Used to Train AI
The Atlantic created a searchable database of the music used to train AI
The Millions of Songs Mashed Into AI-Generated Music
How To Tell If Your Favorite Music Artist Is AI-Generated
SlopTracker: Exposing AI-Generated Artists on Spotify
How we built an AI-music detector for Soulless
Which Suno made it? Version attribution
Snowy City Nights Magical ChristmasLights & Cozy Holiday Vibes
Comment Calling out Timeless Sound
Comment Calling out Breaking Rust
Comment Calling out Xania Monet
Comment Calling out Xania Monet
How Was I supposed to Know? (human version)
if the last time is my last time
it wasn't love, it was survival
Wrong Direction (Single Version)
Unfinished Future (Single Version)
Three Seconds To Goodbye (Live)
Don't Say Goodbye My Love (Live)
I Ain't Got Time For the Fuckery
Blessing Rain Down On A Player
I Don't Believe in Having a Prime
I Forgive You I Just Don't Want You The Fuck Around
I'm Leaving Motherfuckers on Read Today
I Trust My Dog More Than I Trust People
Tell The DJ I’m Trying To Fuck
I Don't Feel Like Going Out Tonight
Too Old To Deal With Any Dumb Shit
When Life Gives Lemons, Mix Them With Tequila
I Never Needed Money To Get Bitches
It’s Cold And I Don’t Got No Hoes
I Feel Like Slapping A Nigga Today
I’m Letting Go Of The Bullshit
New Year, I’m The Same Motherfucker
I Ain’t Worried About Shit Today
New Year, I’m The Same Motherfucker
I’m About To Take My Ass To Sleeep
Wise Words From A Dancer In Georgia
Everybody Got A Fucking Opinion
All Of This Shit Is Gonna Work Out
I’m Letting Go Of The Bullshit
Niggas Trynna Block my Blessings
Girl, I’m Just One Parlay Away
Making More Money Than My Teachers
Feeling Like I Can’t Be Stopped
The Devil Has A Link In Her Bio
First Time In A Long Fucking Time
You Gotta Check That PH Balance, Baby
I Ain’t Fucking With Edibles No More
Soul of A Hustler Vol 1. | Motivational 70s Music
I Feel Like Slapping A Nigga Today
I Wake Up In The Morning and Piss Excellence
Momma, I Think I Married a Hoe
I Caught Santa Clause Sniffing Cocaine
I Don't Need to Be Everything to Everyone
I Don't Need to Prove I Survived
Someone’s playing blues tonight
Lose Control (Rihanna Uranis Remix)
Ordinary (Rihanna Uranis Remix)
Die White a Smile (Rihanna Uranis Remix)
Traditional Christmas Songs, Vol. 1 (Soul Blues Version)
Jingle Bells (Soul Blues Version)
Floating on Echoes (human version)
