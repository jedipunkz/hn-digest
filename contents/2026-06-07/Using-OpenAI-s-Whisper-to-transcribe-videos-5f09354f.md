---
source: "https://byandrev.dev/en/blog/using-whisper-to-transcribe-videos/"
hn_url: "https://news.ycombinator.com/item?id=48435988"
title: "Using OpenAI's Whisper to transcribe videos"
article_title: "Using OpenAI's Whisper to transcribe videos - byandrev"
author: "mcormik"
captured_at: "2026-06-07T18:34:03Z"
capture_tool: "hn-digest"
hn_id: 48435988
score: 1
comments: 1
posted_at: "2026-06-07T15:49:42Z"
tags:
  - hacker-news
  - translated
---

# Using OpenAI's Whisper to transcribe videos

- HN: [48435988](https://news.ycombinator.com/item?id=48435988)
- Source: [byandrev.dev](https://byandrev.dev/en/blog/using-whisper-to-transcribe-videos/)
- Score: 1
- Comments: 1
- Posted: 2026-06-07T15:49:42Z

## Translation

タイトル: OpenAI の Whisper を使用したビデオの文字起こし
記事タイトル: OpenAI の Whisper を使用してビデオを転写する - byandrev
説明: Whisper は OpenAI です

記事本文:
OpenAI の Whisper を使用してビデオを文字起こしする - byandrev
ブログ
テーマの切り替え OpenAI の Whisper を使用したビデオの文字起こし
Whisper は、音声とビデオの文字起こしに革命をもたらした OpenAI の人工知能ツールです。このガイドでは、このオープンソース システムを PC にインストールする方法、またはその API を使用して統合して字幕 (SRT/VTT) を自動的に生成する方法を学びます。
Whisper は、OpenAI によって開発されたオープンソースの自動音声認識 (ASR) システムであり、680,000 時間の多言語およびマルチタスクの監視データに基づいてトレーニングされています。これにより、音声を複数の言語で書き写すことができます。
次のような実際のアプリケーションが数多くあります。
ビデオの字幕作成: 複数の言語に翻訳できる字幕を生成します。
パーソナル アシスタント: 会議、インタビュー、音声メモの文字起こし。
最も良い点は、始めるのが複雑ではないことです。ここでは、最初の一歩を踏み出すのに役立つステップバイステップのガイドを示します。
技術リソースと必要なプライバシーのレベルに応じて、次の 3 つのパスのいずれかを選択できます。
Google Colab ノートブックを使用すると、PC に何もインストールせずにコードを実行でき、Google の無料 GPU を利用できます。
Whisper は PC に直接インストールできます。たとえば、16 GB の RAM を搭載した Ryzen 5 5600G では、ベース モデルは非常に優れたパフォーマンスを発揮します。専用のグラフィック カード (NVIDIA) をお持ちの場合、Whisper はより高速に実行されます。
Whisper をアプリケーションに統合したい場合、またはサーバーを管理したくない場合は、API が解決策です。ここでは音声の 1 分ごとに料金がかかりますが、費用対効果が非常に優れています。
openaiインポートからOpenAI
client = OpenAI( api_key = "YOUR_API_KEY_HERE" )
audio_file_path = "file.mp3"
(audio_file_path, "rb" ) を audio_file として開きます:
転写 = client.audio.transcriptions.create(
メートル

odel = "ささやき-1" ,
ファイル = オーディオファイル、
response_format = "text" # 字幕の場合は "json" または "vtt"
)
印刷（転写）
Linux および macOS へのインストール
Whisper をインストールするには、Python がインストールされている必要があります。ターミナルで次のコマンドを実行します。
pip install -U openai-whisper
オーディオ ファイルとビデオ ファイルを読み取るには、マルチメディア処理ツールである ffmpeg をインストールすることが不可欠です。
sudo apt update && sudo apt install ffmpeg
Homebrew を使用した macOS の場合:
ffmpegを醸造インストールする
基本的な使い方
インストールすると、ターミナルから Whisper コマンドにアクセスできるようになります。ファイルを処理するには、次のコマンドを使用します。
ささやきファイル.mp4 --言語 英語 --モデルベース
主なパラメータ:
-- language : 精度を向上させるためにオーディオの元の言語を設定します。
--model : ハードウェアと精度のニーズに基づいてモデル サイズを選択します。利用可能なモデルは次のとおりです: tiny 、base 、small 、medium 、large 。
詳細な技術情報とソース コードは、公式リポジトリで見つけることができます: https://github.com/openai/whisper
Whisper を使用してビデオの文字起こしを生成します。
コマンドを実行すると、次の形式が生成されます。
.txt : プレーンテキストのみ。タイムスタンプや追加情報はありません。メモや記事に最適です。
.srt : ユニバーサル字幕標準。 YouTubeやビデオプレーヤーと互換性があります。
.vtt : SRT に似ていますが、Web プレーヤー (HTML5) 用に最適化されています。
.json : すべてが含まれます (タイムスタンプ、信頼度、メタデータ)。開発者にとって理想的です。
.tsv : タブ区切りの値。 Excel や Google スプレッドシートで開くのに最適です。
新しいコンテンツを公開したときに通知を受け取ります。スパムなし、いつでも購読解除できます。
© 2026 アンドレス・パラ 全著作権所有。

## Original Extract

Whisper is OpenAI

Using OpenAI's Whisper to transcribe videos - byandrev
Blog
Toggle theme Using OpenAI's Whisper to transcribe videos
Whisper is OpenAI's artificial intelligence tool that has revolutionized audio and video transcription. In this guide, you will learn how to install this open-source system on your PC or integrate it using its API to automatically generate subtitles (SRT/VTT).
Whisper is an Open Source automatic speech recognition (ASR) system developed by OpenAI and trained on 680,000 hours of multilingual and multitask supervised data. It allows us to transcribe audio in multiple languages.
It has many real-world applications, such as:
Video Subtitling: Generating subtitles with the ability to translate into multiple languages.
Personal Assistants: Transcribing meetings, interviews, or voice notes.
The best part is that it is not complicated to get started. Here is a step-by-step guide to help you take your first steps.
Depending on your technical resources and the level of privacy you need, you can choose one of these three paths:
You can use a Google Colab notebook to run the code without installing anything on your PC, taking advantage of Google's free GPUs.
You can install Whisper directly on your PC. For example, on a Ryzen 5 5600G with 16GB of RAM , the base model performs very well. If you have a dedicated graphics card ( NVIDIA ), Whisper will run much faster.
If you are looking to integrate Whisper into an application or don't want to manage servers, the API is the solution. Here you pay per minute of audio, but it is extremely cost-effective.
from openai import OpenAI
client = OpenAI( api_key = "YOUR_API_KEY_HERE" )
audio_file_path = "file.mp3"
with open (audio_file_path, "rb" ) as audio_file:
transcription = client.audio.transcriptions.create(
model = "whisper-1" ,
file = audio_file,
response_format = "text" # "json" or "vtt" for subtitles
)
print (transcription)
Installing on Linux and macOS
To install Whisper, you need to have Python installed. Run the following command in your terminal:
pip install -U openai-whisper
It is essential to install ffmpeg , a multimedia processing tool, to read audio and video files.
sudo apt update && sudo apt install ffmpeg
On macOS with Homebrew:
brew install ffmpeg
Basic Usage
Once installed, you will have access to the whisper command from the terminal. To process a file, use the following command:
whisper file.mp4 --language English --model base
Main parameters:
--language : Sets the original language of the audio to improve accuracy.
--model : Selects the model size based on your hardware and accuracy needs. Available models are: tiny , base , small , medium , and large .
You can find more technical information and the source code in the official repository: https://github.com/openai/whisper
Generating a transcription of a video with Whisper.
When running the command, the following formats are generated:
.txt : Plain text only. No timestamps or extras. Ideal for notes or articles.
.srt : The universal subtitle standard. Compatible with YouTube and video players.
.vtt : Similar to SRT, but optimized for web players (HTML5).
.json : Contains everything (timestamps, confidence, metadata). Ideal for developers.
.tsv : Tab-separated values. Perfect for opening in Excel or Google Sheets.
Get notified when I publish new content. No spam, unsubscribe anytime.
© 2026 Andres Parra All rights reserved.
