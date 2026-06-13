---
source: "https://ma.ttias.be/transcribing-syscast-with-local-ai/"
hn_url: "https://news.ycombinator.com/item?id=48519611"
title: "Transcribing my old podcast locally with open-source AI"
article_title: "Transcribing my old podcast locally with open-source AI · ma.ttias.be"
author: "Mojah"
captured_at: "2026-06-13T18:34:52Z"
capture_tool: "hn-digest"
hn_id: 48519611
score: 2
comments: 1
posted_at: "2026-06-13T17:47:44Z"
tags:
  - hacker-news
  - translated
---

# Transcribing my old podcast locally with open-source AI

- HN: [48519611](https://news.ycombinator.com/item?id=48519611)
- Source: [ma.ttias.be](https://ma.ttias.be/transcribing-syscast-with-local-ai/)
- Score: 2
- Comments: 1
- Posted: 2026-06-13T17:47:44Z

## Translation

タイトル: オープンソース AI を使用して古いポッドキャストをローカルで文字起こしする
記事のタイトル: オープンソース AI を使用してローカルで古いポッドキャストを転写する · ma.ttias.be
説明: 2016 年と 2017 年に、私は [Syscast](/syscast/) というポッドキャストを録画しました。これは、Linux、オープンソース、インフラストラクチャの世界で尊敬する人々へのインタビューです。 …

記事本文:
オープンソース AI を使用して古いポッドキャストをローカルで文字起こしする
2016 年と 2017 年に、私は Syscast というポッドキャストを録音しました。
: Linux、オープンソース、インフラストラクチャの世界で尊敬する人々へのインタビュー。マット・ホルト、キャディについて
、ダニエル・ステンバーグ、カールについて
、Vault についての Seth Vargo
、さらにいくつか。 10 エピソード、約 10 時間の音声があったのですが、生活に支障が出て一時停止しました。
これらのエピソードに決して存在しなかったのは、トランスクリプトでした。ずっと欲しかったんです。音声は素晴らしいですが、検索したり流し読みしたりすることはできず、Google はそれを読むことができません。問題は、2016 年当時、10 時間にわたる 2 人によるインタビューを自分で文字に起こすのは現実的ではなかったということです。まともな音声合成は有料のクラウド サービスであり、2 人の話者を区別するのは基本的に研究プロジェクトでした。
今は 2026 年なので、オープンソース モデルで API 請求なしで、自分のラップトップで夕方に実行しました。その方法は次のとおりです。
次の 2 つのオープンソース部分が機能します。
ウィスパーX
OpenAI の Whisper をラップします
単語レベルのタイムスタンプを備えた、実際の音声テキスト変換用のlarge-v3モデル。
pyannote.audio
話者のダイアライゼーションを処理します。
これを始めたとき、「ダイアライゼーション」という言葉は私にとって新しい言葉でした。これは、録音を話者ごとに分割するステップです。この部分は 1 つの声、この部分は別の声ですが、どちらが誰であるかはまだわかりません。ウィスパーは言われたことを書き留めます。 pyannote は誰が言ったかを調べます。この 2 つを組み合わせると、2 人によるインタビューは、分割されていない 1 つの長いブロックではなく、実際のやり取りのようになります。
どちらもローカルで実行されます。音声がマシンから出ることはなく、分当たりの料金もかかりません。外部から必要なのは、無料の Hugging Face アカウントだけです。
ハグフェイスのゲートモデル #
ダイアライゼーション モデルは、Hugging Face でゲートされます。理由はわかりません。それらは危険ですか、そしてそうする必要がありますか

迷ってサインしますか？知るか。 🤷‍♂️
私がしたのは、(無料) アカウントを作成し、読み取りトークンを作成し、モデル ページで「同意する」をクリックしただけで、モデルをダウンロードできました。私は初めてでしたが、WhisperX は次のように挨拶してくれました。
「pyannote/speaker-diarization-3.1」パイプラインをダウンロードできませんでした。
パイプラインがプライベートまたはゲートされていることが原因である可能性があります...
これはネットワークまたは認証のバグのように見えますが、ライセンスがまだ受け入れられていないことを意味します。これを試す場合は、まず pyannote/speaker-diarization-3.1 と pyannote/segmentation-3.0 の条件を受け入れ、トークンを ~/.hf_token にドロップすると機能します。
セットアップは virtualenv と 1 回のインストールです (私は uv を使用しました)
):
uv venv .venv-Whisper
uv pip install --python .venv-whisper whisperx
コアは約 12 行の WhisperX です。 int8 量子化を使用してモデルを CPU にロードし (Apple Silicon にはこのスタックに使用できる CUDA パスがありません)、文字起こしし、正確な単語レベルのタイムスタンプに合わせて調整し、ダイアライズして各単語を話者に割り当てます。
インポートウィスパーx
オーディオ = ささやきx 。ロードオーディオ (mp3_path)
#1. Whisperlarge-v3で文字起こしする
モデル = ウィスパーx 。 load_model ( "large-v3" 、 "cpu" 、 compute_type = "int8" 、 language = "en" )
結果 = モデル。転写 (オーディオ、バッチサイズ = 1、言語 = "en")
# 2. 整列、正確なワードレベルのタイムスタンプ
align_model 、メタ = ウィスパーx 。 load_align_model (言語コード = "en"、デバイス = "cpu")
結果 = ささやきx 。 align ( result [ "segments" ] 、 align_model 、 meta 、 audio 、 "cpu" )
# 3. (誰がいつ話したか) を日記化し、各単語に話者をタグ付けします。
ダイアライズ = ささやきx 。 DiarizationPipeline ( use_auth_token = hf_token 、 device = "cpu" )
結果 = ささやきx 。 assign_word_speakers (diarize (オーディオ)、結果)
これにより、各セグメントに Speaker 、 start 、 end 、 text が含まれるセグメントのリストが得られます。 10 のエピソードすべてをバッチ処理するのは、単にループするだけです。

mp3s、ファイルごとの実時間を記録します (以下のベンチマークはそこから来ています):
static/podcast/episodes/*.mp3 のエピソードの場合。する
開始 = $( 日付 +%s )
.venv-whisper/bin/python scripts/transcribe-syscast.py " $episode "
printf '%s\t%ss\n' " $(basename " $episode " ) " " $(( $( date +%s ) - start )) "
完了しました
WhisperX の生の出力は途切れ途切れです。多くの短いセグメント、SPEAKER_00 / SPEAKER_01 ラベル (最初と 2 番目に話した人だけ)、段落はありません。
[0:00] SPEAKER_00: Syscast の新しいエピソードへようこそ。私の名前は Mattias Geniar です。今日は HashiCorp の Seth Vargo が参加します。
[0:14] SPEAKER_01: やあ、マティアス、元気だよ。ここピッツバーグではうまくやっている。
小さなクリーンアップ ステップにより、同じ話者による連続したセグメントが 1 つのターンにマージされ、その後、長いターンが数文ごとに段落に分割されます。
ターン = []
セグメント内のセグメントの場合:
if 何度も繰り返した場合 [ - 1 ][ "spk" ] == seg [ "speaker" ]:
[ - 1 ][ "テキスト" ] += " " + seg [ "テキスト" ] になります。ストリップ() # 同じ話者、マージを続ける
それ以外の場合:
ターンします。 append ({ "spk" : seg [ "speaker" ], "start" : int ( seg [ "start" ])、 "text" : seg [ "text" ] .strip ()})
最後の仕上げは、SPEAKER_00 を「Mattias」に、SPEAKER_01 をゲスト (エピソード タイトルに名前が記載されています) にマッピングし、明らかな聞き間違いを修正することです。ウィスパーは私の名前が「マティアス・ゲンジャール」であると非常に自信を持っていました。 😁
これは各エピソードのページに、読みやすい話者ラベル付きの会話として表示され、タイムスタンプをクリックすると音声に直接ジャンプできます。たとえば、Vault の Seth Vargo です。
または Jan-Piet Mens on Linux vs BSD
。
ただし、実現可能というのは速いという意味ではありません。 CPU (Apple Silicon、このスタックに使用可能な GPU パスがない) でのlarge-v3 プラスダイアライゼーションは遅いです。自分のマシン上でエピソードごとに:
リアルタイムで約 2 回呼び出し、全体で約 14 時間のコンピューティングを要します。

電子カタログ。一晩実行したらラップトップが暖かくなりました。数分で完了させたい場合は、数ドルでホスト型 API を使用することになるでしょうが、ここでのポイントはその逆で、これをローカルで無料で自分で実行できるか?ということです。はい。
トラフィックがそれほど多くなかった 10 個の古いエピソードの場合、値はコンピューティング時間ではありません。マットやダニエルのような人々との 10 年間の会話がテキストになり、検索、スキミング、インデックス付けが可能になりました。バック カタログは第 2 の人生を手に入れました。費用は暖かいラップトップと一晩の電気代だけでした。
私がこだわっているのはタイムラインです。 2016 年には 1 人の人間には手の届かなかったまさにそのタスクが、今では私の目の前のラップトップ上でエンドツーエンドで実行されています。それは継続的に起こっており、注目に値します。
フォローしていただくと、新しい投稿がフィードに表示されます。お好みであれば、RSS も機能します。
感謝してくれる人と共有してください。言及することには大きな意味があります。
×
リンクトイン
Facebook // 読み続けてください
私が毎日使用している Linux デバッグ ツール
2026 年 6 月 12 日
ポッドキャスト: 構成管理キャンプ: Kubernetes、Sysdig、および管理
2017 年 3 月 2 日
要約: Linux パフォーマンスの測定 - 典型的な間違いを回避する方法
2020年7月9日
Linux のパフォーマンスを測定する方法 最もよくある間違いを回避する: ネットワーク
2020年7月8日
サイトを維持することが Oh Dear の仕事です
Oh Dear は、私が構築を支援している監視プラットフォームで、グローバル企業、主要なオープンソース プロジェクト、公共部門のサービスから信頼され、サイトを 24 時間監視しています。訪問者が気付かないうちに、複数の場所からあなたのサーバーがダウンした瞬間に警告を発します。
ベルギーの開発者兼 Linux システム管理者。2008 年からサーバー、PHP、DNS、SSL、Web について執筆。Oh Dear の共同創設者。

## Original Extract

Back in 2016 and 2017 I recorded a podcast called [Syscast](/syscast/): interviews with people I admired in the Linux, open source and infrastructure world. …

Transcribing my old podcast locally with open-source AI
Back in 2016 and 2017 I recorded a podcast called Syscast
: interviews with people I admired in the Linux, open source and infrastructure world. Matt Holt about Caddy
, Daniel Stenberg about curl
, Seth Vargo about Vault
, and a handful more. Ten episodes, roughly ten hours of audio, and then life got in the way and I put it on pause.
The one thing those episodes never had was transcripts. I always wanted them. Audio is nice, but you can’t search it, you can’t skim it, and Google can’t read it. The problem was that in 2016, transcribing ten hours of two-person interviews yourself just wasn’t realistic. Decent speech-to-text was a paid cloud service, and telling two speakers apart was basically a research project.
It’s 2026 now, so I did it in an evening, on my own laptop, with open-source models and no API bill. Here’s how.
Two open-source pieces do the work:
WhisperX
wraps OpenAI’s Whisper
large-v3 model for the actual speech-to-text, with word-level timestamps.
pyannote.audio
handles the speaker diarization.
“Diarization” was a new word to me when I started this. It’s the step that splits a recording up by speaker: this stretch is one voice, this stretch is another , without knowing who either of them is yet. Whisper writes down what gets said; pyannote works out who said it. Put the two together and a two-person interview reads as a real back-and-forth instead of one long undivided block.
Both run locally. The audio never leaves the machine and there’s no per-minute cost. The only thing you need from the outside world is a free Hugging Face account.
Gated models on Hugging Face #
The diarization models are gated on Hugging Face. No idea why. Are they dangerous and you need to sign a waver? Who knows. 🤷‍♂️
All I did was create a (free) account, a read token and clicked “agree” on the model pages before I could download them. I hadn’t first, and WhisperX greeted me with this:
Could not download 'pyannote/speaker-diarization-3.1' pipeline.
It might be because the pipeline is private or gated...
That reads like a network or auth bug, but it just means the licence wasn’t accepted yet. If you try this, accept the conditions on pyannote/speaker-diarization-3.1 and pyannote/segmentation-3.0 first, drop your token in ~/.hf_token , and it works.
Setup is a virtualenv and one install (I used uv
):
uv venv .venv-whisper
uv pip install --python .venv-whisper whisperx
The core is about a dozen lines of WhisperX. Load the model on CPU with int8 quantization (Apple Silicon has no usable CUDA path for this stack), transcribe, align for accurate word-level timestamps, then diarize and assign each word to a speaker:
import whisperx
audio = whisperx . load_audio ( mp3_path )
# 1. transcribe with Whisper large-v3
model = whisperx . load_model ( "large-v3" , "cpu" , compute_type = "int8" , language = "en" )
result = model . transcribe ( audio , batch_size = 1 , language = "en" )
# 2. align, for accurate word-level timestamps
align_model , meta = whisperx . load_align_model ( language_code = "en" , device = "cpu" )
result = whisperx . align ( result [ "segments" ], align_model , meta , audio , "cpu" )
# 3. diarize (who spoke when), then tag each word with a speaker
diarize = whisperx . DiarizationPipeline ( use_auth_token = hf_token , device = "cpu" )
result = whisperx . assign_word_speakers ( diarize ( audio ), result )
That gives me a list of segments, each with a speaker , start , end and text . Batching all ten episodes is just a loop over the mp3s, logging the wall-clock time per file (that’s where the benchmark below comes from):
for episode in static/podcast/episodes/*.mp3 ; do
start = $( date +%s )
.venv-whisper/bin/python scripts/transcribe-syscast.py " $episode "
printf '%s\t%ss\n' " $( basename " $episode " ) " " $(( $( date +%s ) - start )) "
done
Raw WhisperX output is choppy: lots of short segments, SPEAKER_00 / SPEAKER_01 labels (just whoever talked first and second), no paragraphs:
[0:00] SPEAKER_00: Welcome to a new episode of Syscast. My name is Mattias Geniar and today I'm joined by Seth Vargo from HashiCorp.
[0:14] SPEAKER_01: Hey Mattias, I'm good. Doing well over here in Pittsburgh.
A small cleanup step merges consecutive segments from the same speaker into one turn, then splits long turns into paragraphs every few sentences:
turns = []
for seg in segments :
if turns and turns [ - 1 ][ "spk" ] == seg [ "speaker" ]:
turns [ - 1 ][ "text" ] += " " + seg [ "text" ] . strip () # same speaker, keep merging
else :
turns . append ({ "spk" : seg [ "speaker" ], "start" : int ( seg [ "start" ]), "text" : seg [ "text" ] . strip ()})
The last touch is mapping SPEAKER_00 to “Mattias” and SPEAKER_01 to the guest (whose name is right there in the episode title), and fixing the obvious mis-hearings. Whisper was very confident my name is “Matthias Genjar”. 😁
That lands on each episode page as a readable, speaker-labelled conversation with clickable timestamps to jump straight into the audio. For example, Seth Vargo on Vault
or Jan-Piet Mens on Linux vs BSD
.
Feasible doesn’t mean fast, though. large-v3 plus diarization on a CPU (Apple Silicon, no usable GPU path for this stack) is slow. Per episode, on my own machine:
Call it roughly twice real-time, and about 14 hours of compute for the whole catalogue. I ran it overnight and the laptop got warm. If I wanted it done in minutes I’d have used a hosted API for a few dollars, but the point here was the opposite: can I do this myself, locally, for free? Yes.
For ten old episodes that weren’t getting much traffic, the value isn’t the compute time. A decade of conversations with people like Matt and Daniel is now text: searchable, skimmable and indexable. The back catalogue gets a second life, and it cost me nothing but a warm laptop and a night of electricity.
The thing that sticks with me is the timeline. The exact task that was out of reach for one person in 2016 now runs, end to end, on the laptop in front of me. That keeps happening, and it’s worth noticing.
Follow along and new posts show up in your feed. RSS works too, if that's your thing.
Share it with someone who'd appreciate it. A mention means a lot.
X
LinkedIn
Facebook // keep reading
Linux debugging tools I use daily
Jun 12, 2026
Podcast: Config Management Camp: Kubernetes, Sysdig & Mgmt
Mar 2, 2017
Recap: measuring linux performance - how to avoid typical mistakes
Jul 9, 2020
How to measure Linux Performance Avoiding Most Typical Mistakes: Network
Jul 8, 2020
Keeping sites up is what Oh Dear does
Oh Dear is the monitoring platform I help build, trusted by global companies, major open-source projects and public-sector services to watch their sites around the clock. It alerts you the moment yours goes down, from multiple locations, before your visitors ever notice.
Belgian developer & Linux sysadmin, writing about servers, PHP, DNS, SSL and the web since 2008. Co-founder of Oh Dear .
