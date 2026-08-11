---
source: "https://github.com/karimm-ai/NiceShot_AI"
hn_url: "https://news.ycombinator.com/item?id=49253675"
title: "Show HN: NiceShot AI – The analytics layer competitive games forgot to add"
article_title: "GitHub - karimm-ai/NiceShot_AI: A Python tool powered by computer vision to analyze gameplay videos and auto-detect, track and clip key gameplay events as well as create compilations in horizontal & vertical formats. In addition to creating summary reports for gameplay sessions. · GitHub"
author: "niceshot-ai"
captured_at: "2026-08-11T05:52:28Z"
capture_tool: "hn-digest"
hn_id: 49253675
score: 1
comments: 0
posted_at: "2026-08-11T05:18:28Z"
tags:
  - hacker-news
  - translated
---

# Show HN: NiceShot AI – The analytics layer competitive games forgot to add

- HN: [49253675](https://news.ycombinator.com/item?id=49253675)
- Source: [github.com](https://github.com/karimm-ai/NiceShot_AI)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T05:18:28Z

## Translation

タイトル: Show HN: NiceShot AI – 追加するのを忘れていた分析レイヤーの競争力のあるゲーム
記事タイトル: GitHub - karimm-ai/NiceShot_AI: コンピュータ ビジョンを利用した Python ツールで、ゲームプレイ ビデオを分析し、主要なゲームプレイ イベントを自動検出、追跡、クリップし、水平および垂直フォーマットでコンピレーションを作成します。ゲームプレイ セッションの概要レポートの作成に加えて。 · GitHub
説明: コンピューター ビジョンを利用した Python ツールで、ゲームプレイ ビデオを分析し、主要なゲームプレイ イベントを自動検出、追跡、クリップし、水平および垂直フォーマットでコンピレーションを作成します。ゲームプレイ セッションの概要レポートの作成に加えて。 - karimm-ai/NiceShot_AI
HN テキスト: 私は、ゲームプレイ分析のためのコンピューター ビジョンを実験する方法として、過去 26 か月間、NiceShot AI に取り組んできました。基本的なアイデアは、記録された長時間のゲームプレイ セッションを取得し、それを構造化されたゲームプレイ情報とハイライトに自動的に変換することです。試合の統計 (短すぎる) と生涯の統計 (長すぎる) の間のパフォーマンス分析のギャップを埋める。現在のパイプラインは次のとおりです。 ゲームプレイ ビデオ → YOLO 検出 → イベント トラッキング → OCR/コンテキスト フィルタリング → イベント タイムスタンプ → クリップ → ハイライト コンパイル → セッション統計 たとえば、ゲーム固有の HUD 要素を認識するように YOLO モデルを微調整し、それらの検出をキル、死亡、メダルなどのイベントとして解釈するようにパイプラインを構成できます。私が避けたかったことの 1 つは、システム全体をゲーム固有にすることでした。別のゲームをサポートするには、パイプライン全体を書き直すのではなく、新しい検出器/モデルと構成が必要になることがほとんどです。そして現在、私はそれに取り組んでいます。また、ゲームプレイ イベントとしてカウントすべきではない状態にも OCR を使用します。たとえば、Call of Duty ゲームでは、「キル」中にキルインジケーターが表示されることがあります。

「カメラ」または「観戦中」状態なので、OCR レイヤーを使用してこれらのケースをフィルタリングできます。その後、システムは検出されたイベントの周囲のクリップを生成し、ショート/TikTok スタイルのコンテンツの垂直バージョンを含むハイライト リールにコンパイルできます。最近では、イベント検出後に軽量のビデオ LLM を使用する実験も行っています。数時間のゲームプレイ セッション全体をビデオ LLM モデルに送信する代わりに、興味深いイベントを中心に抽出された短いクリップを取得し、モデルに何が起こったのか説明してもらいます。目標は、最終的には、これを軽量の会話型ゲームプレイ コーチに変えます。このプロジェクトは現在、古い GTX1650 4GB VRAM ラップトップで実験しましたが、もちろん問題なく動作しますが、遅いです。最後に、これが私の当初の目標だったことはわかっていますが、コンピューター ビジョン推論がプレイヤーの実際のゲーム パフォーマンスに干渉することを望まなかったので、オフライン処理を選択しました。検出器/推論層は、顕著な GPU/CPU オーバーヘッドを引き起こすことなく、リアルタイム ゲームプレイに十分な速度を提供します。プロジェクトはまだ発展途上であるため、技術的な批判や提案は大歓迎です。

記事本文:
GitHub - karimm-ai/NiceShot_AI: コンピュータ ビジョンを利用した Python ツールで、ゲームプレイ ビデオを分析し、主要なゲームプレイ イベントを自動検出、追跡、クリップし、水平および垂直フォーマットでコンピレーションを作成します。ゲームプレイ セッションの概要レポートの作成に加えて。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
カリムアイ
/
NiceShot_AI
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
154 コミット

154 コミット game_models game_models src/ niceshot_ai src/ niceshot_ai .gitignore .gitignore ライセンス ライセンス README.md README.md icon.ico icon.ico要件.txt要件.txtサンプル_レポート.pngサンプル_レポート.pngバージョン.jsonバージョン.jsonすべてのファイルを表示リポジトリ ファイル ナビゲーション
NiceShot_AI: Python コンピューター ビジョン ツール
NiceShot AI は、ゲームプレイ ビデオを分析するためのコンピューター ビジョンを利用した Python ツールです。 NiceShot AI は、YOLO、OpenCV、FFmpeg、matplotlib などの最先端のツールを統合することで、主要なゲームプレイ イベントを自動的に検出、追跡、クリップし、ゲームプレイ セッション分析用の視覚的なレポートを作成するように設計されています。
ツールの結果を紹介する簡単なデモ: ( https://youtu.be/op1GDREXiOg )
Call of Duty: Black Ops 7 (2025) --> まだテスト中
コール オブ デューティ: ブラックオプス 6 (2024)
Call of Duty: Modern Warfare II (2022) --> まだテスト中
Ultralytics による YOLOv8n 。 CC ライセンスに基づいてカスタムで収集され、注釈が付けられたゲームプレイ ビデオのデータセットを微調整しました。
スケーラビリティに対する堅牢性 : 構成可能な変数を使用することで、大規模な変更を加えることなく、さまざまなゲーム モデルやイベントに適応できるようになります。 (テスト中)
正確なイベント確認 : EasyOCR を使用して、特殊なゲーム シーン (例: KILLCAMS や SPECTATING) で発生するイベントのカウントを防ぎます。
特別なイベントの検出: (例: 時間しきい値内の複数の連続キルの組み合わせから発生するキル ストリーク)。
イベントのタイムスタンプと CSV 出力: 検出されたイベントにタイムスタンプを付け、[タイムスタンプ、イベント] の 2 列の CSV ファイルにダンプし、さらなるゲームプレイ データの分析と検査を行います。
セッション分析 : セッション後の統計分析を提供する複数のグラフで構成される概要レポートを作成します。
イベントの自動クリッピング : イベントの開始時刻と終了時刻を使用して、検出されたイベントをクリッピングします。
16:9 および TikTok 形式でクリップをエクスポート
C

ハイライト リールの作成 : フォルダー内のすべてのクリップを 1 つのコンピレーション ビデオに連結し、垂直方向と水平方向の両方のフォーマットでクリップ間の簡単なフェードインとフェードアウトのトランジション編集を行います。
Custom Reel Lengths : 抽出されたクリップから任意の長さのコンピレーションを作成できます。
Twitch チャンネルからビデオを一括分析: Twitch チャンネルから目的のゲーム ストリームをダウンロードして分析し、ゲームプレイ ビデオの一括分析を実行します。 (テスト中)
ランキングスペシャルクリップ：Ex. (イベント中に複数のメダルがポップアップするホット キル クリップ)。
イベント検出は完璧ではありません: 私のテストによると、イベントは複数回検出される場合もあれば、まったく検出されない場合もあります。
NiceShot_AI を使い始めるには、まず公式 Web サイト (https://www.ffmpeg.org/download.html) から ffmpeg をダウンロードしてインストールし、PATH に追加します。
Python 仮想環境を作成します (オプションですが推奨)。私のPythonのバージョンは3.10.11です
Python -m venv venv
venv \S cripts \a アクティブ化
トーチcudaをインストールします。 GTX1650 4GBのcuda12.1を使用しました。現在、RTX5070 8GB に nightly cuda 12.8 を使用しています。
pip install torch torchvision torchaudio --index-url https://download.pytorch.org/whl/cu121
依存関係をインストールします。
pip install -r 要件.txt
ツールを実行する
src.niceshot_ai.detector から EventDetector をインポート
検出器 = EventDetector(
「Call of Duty Black Ops 6」、# 適切なモデルとチャート構成をインポートするためのゲーム名
"video1.mp4", # 分析するゲームプレイ ビデオ パス
total_hours=100, # ビデオの分析にかかる合計時間
save_clips=True, # イベントを自動クリップし、イベント クリップをローカルに保存します (コンパイルの作成に必要)
Output_dir=".", # クリップ、チャート、CSV ファイルの出力フォルダー
max_workers=2, # 自動クリッピングのデフォルト
Frame_idx_start=0, # ビデオの開始フレーム
Frames_to_skip=8, # 分析中にスキップするフレーム数 (

多ければ多いほど、分析が速くなります)
add_to_csv=True、# タイムスタンプ イベントと CSV への出力
create_montage=True, # クリップ可能なすべてのイベントのコンパイルを作成します
montage_length_sec=50, # コンパイルの長さ
max_videos=1、
vertical_format=False、# 垂直フォーマットでのみイベントを自動クリップします
Advanced_detection=True, # 一部のイベントを OCR で確認する
session_analysis=False # レポート作成者
)
detecter.detect_events()
処理速度
次の仕様のラップトップ_1 でテストしました。
次の仕様のlaptop_2でテストしました:
これは、イベントの検出後にイベントを確認するためにのみ実行されます。ビデオフレーム全体ではありません。イベントが確認されるまで、処理速度が一時的に 170 FPS から 30 FPS (新しいラップトップの場合) に低下する可能性があります。もちろんオフにすることもできますが、これにより「SPECTATING」中のキルイベントがカウントされます。
コンピューター ビジョンを利用した Python ツールで、ゲームプレイ ビデオを分析し、主要なゲームプレイ イベントを自動検出、追跡、クリップし、水平および垂直フォーマットでコンピレーションを作成します。ゲームプレイ セッションの概要レポートの作成に加えて。
Readme AGPL-3.0 ライセンス アクティビティ スター
1 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A Python tool powered by computer vision to analyze gameplay videos and auto-detect, track and clip key gameplay events as well as create compilations in horizontal & vertical formats. In addition to creating summary reports for gameplay sessions. - karimm-ai/NiceShot_AI

I've been working on NiceShot AI for the last 26 months as a way to experiment with computer vision for gameplay analysis. The basic idea is to take a recorded long gameplay session and automatically turn it into structured gameplay information and highlights. Bridging the performance analysis gap between match stats (too short) & lifetime stats (too long). The current pipeline is: Gameplay video → YOLO detection → event tracking → OCR/context filtering → event timestamps → clips → highlight compilation → session statistics For example, I can fine-tune a YOLO model to recognize game-specific HUD elements and then configure the pipeline to interpret those detections as events such as kills, deaths, or medals. One thing I wanted to avoid was making the entire system game-specific. Supporting another game should mostly require a new detector/model and configuration rather than rewriting the whole pipeline. And currently, I am working on that. I also use OCR for states that shouldn't be counted as gameplay events. For example, in Call of Duty games a kill indicator can appear during a "KillCam" or "Spectating" state, so the OCR layer can be used to filter those cases. The system can then produce clips around detected events and compile them into highlight reels, including vertical versions for Shorts/TikTok-style content. Recently, I have also been experimenting with using a lightweight Video LLM after event detection. Instead of sending the entire multi-hour gameplay session to the Video LLM model, I take the extracted short clips around an interesting event and ask the model to explain what happened. The goal is to eventually turn this into a lightweight conversational gameplay coach. The project currently runs locally on NVIDIA GPUs. Have experimented with an old GTX1650 4GB VRAM laptop and it works fine but slow ofcourse. Finally, I know that real-time event detection should be the way to go and this was my original goal, but I chose offline processing because I didn't want computer vision inference to interfere with the player's actual game performance. I'd also be interested in hearing how others would approach making the detector/inference layer fast enough for real-time gameplay without causing noticeable GPU/CPU overhead. Keeping in mind, I input 1080p frames into the model. The project is still evolving, so technical criticism and suggestions are very welcome. Thanks.

GitHub - karimm-ai/NiceShot_AI: A Python tool powered by computer vision to analyze gameplay videos and auto-detect, track and clip key gameplay events as well as create compilations in horizontal & vertical formats. In addition to creating summary reports for gameplay sessions. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
karimm-ai
/
NiceShot_AI
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
154 Commits 154 Commits game_models game_models src/ niceshot_ai src/ niceshot_ai .gitignore .gitignore License License README.md README.md icon.ico icon.ico requirements.txt requirements.txt sample_report.png sample_report.png version.json version.json View all files Repository files navigation
NiceShot_AI: Python Computer Vision Tool
NiceShot AI is a Python tool powered by computer vision to analyze gameplay videos. With the integration of cutting-edge tools like YOLO, OpenCV, FFmpeg, and matplotlib, NiceShot AI is designed to automatically detect, track and clip key gameplay events as well as create visual report for gameplay session analysis.
Simple demo showcasing tool results: ( https://youtu.be/op1GDREXiOg )
Call of Duty: Black Ops 7 (2025) --> Still in testing
Call of Duty: Black Ops 6 (2024)
Call of Duty: Modern Warfare II (2022) --> Still in testing
YOLOv8n by Ultralytics . Fine-tuned on custom collected & annotated dataset of gameplay videos under CC license.
Robust to Scalability : Uses configurable variables enabling it to adapt to different game models and events without massive changes. (In testing)
Accurate Event Confirmation : Using EasyOCR to prevent counting events occurring in special game scenes (ex.KILLCAMS and SPECTATING).
Special Events Detection : (Ex. Kill Streaks occurring from the combination of multiple consecutive kills within a time threshold).
Events Timestamping & CSV Output : Timestamps detected events and dumps into a CSV file with 2 columns [Timestamp, Event] for further gameplay data analysis and inspections.
Session Analysis : Creates a summary report consisting of multiple charts providing a post-session stats analysis.
Event Auto-Clipping : Clipping detected events using event's start time and end time.
Clips Export in 16:9 & TikTok formats
Creating Highlight Reels : Concatenating all clips within a folder into one compilation video with simple fade in & out transition edits between clips in both vertical & horizontal formats.
Custom Reel Lengths : Allowing for creating compilations of any length from the extracted clips.
Analyzing videos in bulk from a Twitch channel : Downloads and analyzes desired game streams from a Twitch channel performing bulk analysis of gameplay videos. (In testing)
Ranking Special Clips : Ex. (Hot Kill Clips where multiple medals pop up during the event).
Event detection is not perfect : From my testing, an event can get detected more than once or not detected at all.
To get started with NiceShot_AI , download & install ffmpeg from the official website: https://www.ffmpeg.org/download.html first and add it to your PATH.
Create a Python virtual environment (optional, but recommended). My Python version is 3.10.11
python -m venv venv
venv \S cripts \a ctivate
Install torch cuda. I used cuda 12.1 for GTX1650 4GB. Currently, I am using nightly cuda 12.8 for RTX5070 8GB:
pip install torch torchvision torchaudio --index-url https://download.pytorch.org/whl/cu121
Install dependencies:
pip install -r requirements.txt
Run the tool
from src.niceshot_ai.detector import EventDetector
detector = EventDetector(
"Call of Duty Black Ops 6", # Game name to import proper model and chart configuration
"video1.mp4", # Gameplay video path to analyze
total_hours=100, # Total hours to analyze of the video
save_clips=True, # Auto-clip events and save event clips locally (required for compilation making)
output_dir=".", # Output folder for clips, charts and csv file
max_workers=2, # default for auto-clipping
frame_idx_start=0, # Start frame of the video
frames_to_skip=8, # Frames to skip during analysis (the more, the faster tha analysis)
add_to_csv=True, # Timestamp events and output to csv
create_montage=True, # Create compilation of every clippable event
montage_length_sec=50, # Length of the compilation
max_videos=1,
vertical_format=False, # Auto-clip events in vertical format only
advanced_detection=True, # Confirm some events with OCR
session_analysis=False # Report maker
)
detector.detect_events()
Processing Speed
Tested on laptop_1 with the following specs:
Tested on laptop_2 with the following specs:
This is run only to confirm an event after it's detected. Not through the whole video frames. It can cause the processing speed to fall down from 170 FPS to 30 FPS (on new laptop) temporarily until event is confirmed. It can definitely be turned off, however this will cause a kill event during "SPECTATING" to be counted.
A Python tool powered by computer vision to analyze gameplay videos and auto-detect, track and clip key gameplay events as well as create compilations in horizontal & vertical formats. In addition to creating summary reports for gameplay sessions.
Readme AGPL-3.0 license Activity Stars
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
