---
source: "https://github.com/lbj96347/nemotron-3.5-asr-ios"
hn_url: "https://news.ycombinator.com/item?id=48471215"
title: "Offline streaming speech recognition on iOS with Nvidia Nemotron 3.5 and Core ML"
article_title: "GitHub - lbj96347/nemotron-3.5-asr-ios: On-device, offline speech recognition for iPhone/iPad using NVIDIA's Nemotron-3.5-ASR Streaming 0.6B (multilingual) via CoreML.SwiftUI app with mic capture + audio file import, RNN-Tdecoding, and live benchmark metrics (latency, RTF, memory). · GitHub"
author: "lbj96347"
captured_at: "2026-06-10T04:34:02Z"
capture_tool: "hn-digest"
hn_id: 48471215
score: 1
comments: 0
posted_at: "2026-06-10T03:52:49Z"
tags:
  - hacker-news
  - translated
---

# Offline streaming speech recognition on iOS with Nvidia Nemotron 3.5 and Core ML

- HN: [48471215](https://news.ycombinator.com/item?id=48471215)
- Source: [github.com](https://github.com/lbj96347/nemotron-3.5-asr-ios)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T03:52:49Z

## Translation

タイトル: Nvidia Nemotron 3.5 と Core ML を使用した iOS でのオフライン ストリーミング音声認識
記事のタイトル: GitHub - lbj96347/nemotron-3.5-asr-ios: NVIDIA の Nemotron-3.5-ASR ストリーミング 0.6B (多言語) を使用した iPhone/iPad 向けのオンデバイス、オフライン音声認識 (マイク キャプチャ + オーディオ ファイル インポート、RNN-T デコーディング、ライブ ベンチマーク メトリクス (レイテンシー、RTF、メモリ) を備えた CoreML.SwiftUI アプリ経由)。 · GitHub
説明: マイク キャプチャ + オーディオ ファイル インポート、RNN-T デコーディング、およびライブ ベンチマーク メトリクス (レイテンシ、RTF、メモリ) を備えた CoreML.SwiftUI アプリ経由で、NVIDIA の Nemotron-3.5-ASR ストリーミング 0.6B (多言語) を使用した、iPhone/iPad 向けのオンデバイスのオフライン音声認識。 - lbj96347/nemotron-3.5-asr-ios

記事本文:
GitHub - lbj96347/nemotron-3.5-asr-ios: NVIDIA の Nemotron-3.5-ASR ストリーミング 0.6B (多言語) を使用した、iPhone/iPad 向けのオンデバイスのオフライン音声認識 (マイク キャプチャ + オーディオ ファイル インポート、RNN-T デコーディング、ライブ ベンチマーク メトリクス (レイテンシ、RTF、メモリ) を備えた CoreML.SwiftUI アプリ経由)。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
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
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードt

o セッションをリフレッシュします。
アラートを閉じる
{{ メッセージ }}
lbj96347
/
nemotron-3.5-asr-ios
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミット モデル モデル NemotronASRPoC NemotronASRPoC NemotronASRPoCTests NemotronASRPoCTests スクリプト scripts .gitignore .gitignore CLAUDE.md CLAUDE.md IMPLEMENTATION-PLAN.md IMPLEMENTATION-PLAN.md ライセンス ライセンス README.md README.md project.yml project.yml 提案された-plan.md 提案された-plan.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
NVIDIA を使用したオンデバイス、オフライン、ストリーミング音声認識の PoC
Nemotron-3.5-ASR iPhone/iPad 上の CoreML 経由で 0.6B (多言語) ストリーミング。
両方の推論パスが物理デバイス上でエンドツーエンドで実行されるようになりました: ライブ マイク ストリーミング
(フェーズ 7) とオフライン ファイルの転写。シミュレータは引き続き完全なシェルを実行します
CPU/GPU 上の (オーディオ + ベンチマーク) ですが、ANE のパフォーマンスを表すものではありません。
完全なエンジニアリング計画については、IMPLEMENTATION-PLAN.md を参照してください。
元の製品意図についてはprovided-plan.md。
Xcode 16+ (Xcode 26.3 でビルド/テスト済み)
iPhone 15 Pro 以降を推奨 (ANE)。実際の推論にはデバイスが必要です
XcodeGen ( brew install xcodegen )
別途ダウンロードされた CoreML モデル ファイル (「モデルのセットアップ」を参照)
# 1. project.yml から Xcode プロジェクトを生成します (project.yml の変更後に実行)
xcodegen 生成
＃2a。シミュレーター用のビルド (シェル + オーディオ + ベンチマーク。実際の推論はありません)
xcodebuild -scheme NemotronASRPoC -destination ' generic/platform=iOS Simulator ' ビルド
＃2b。または、Xcode で開いて物理デバイス上で実行します
NemotronASRPoC.xcodeproj を開く
# 3. 単体テストを実行する (リサンプラー、チャンク バッファー、レイテンシー トラッカー)
xcodebuild test -scheme NemotronASRPoC -destination ' platform=iOS Si

エミュレータ、名前=iPhone 17 Pro '
# 単一のテストを実行する
xcodebuild テスト -スキーム NemotronASRPoC \
-destination ' プラットフォーム=iOS シミュレーター、名前=iPhone 17 Pro ' \
-only-testing:NemotronASRPoCTests/AudioPipelineTests/testResamplerDownsamplesTo16k
アプリはモデルが存在しなくても実行されます。「モデルが見つかりません」ステータスが表示されます。
ただし、マイクオーディオをキャプチャし、16 kHz にリサンプリングし、ストリーミング層にチャンクし、
ライブベンチマークメトリクスをレポートします。これにより、ダウンロード前にシェルを検証できます。
ファイル トランスクリプション テストは、TestAudio/ (
フォルダーは gitignored です — 独自のフォルダーを用意してください)。各クリップに言語ヒントを付けて名前を付けます。
スイートは適切なプロンプトを選択します。 Sample-en.m4a（英語）またはsample-yue.m4a
(広東語→自動)。一致するクリップ (またはモデル) が存在しない場合、テストはきれいにスキップされます。
デフォルトのターゲット: 多言語 @ 2240 ミリ秒 (~634 MB、EN / zh / ja / ko をカバー。
広東語には専用のプロンプトがなく、自動に戻ります)。
# 5 つの .mlmodelc バンドル + トークナイザー + メタデータをすべて Models/multilingual/2240ms/ にダウンロードします
./scripts/download_models.sh 多言語 2240
# 再検査して署名を更新します (すでにチェックインされていますが、層を変更した場合は再実行します)
python3 scripts/inspect_model.py Models/multilingual/2240ms --out NemotronASRPoC/ASR/ModelSignatures.json
# Models/ フォルダーをアプリに再バンドルします
xcodegen 生成
Models/multilingual/2240ms/ に予期されるファイル:
preprocessor.mlmodelc # audio[1,?] → mel[1,128,?], mel_length
encoder.mlmodelc # mel + キャッシュ + プロンプト ID → エンコード + 更新されたキャッシュ (ステートフル)
decoder_joint.mlmodelc # B1 融合デコーダ⊕ジョイント — デフォルトの RNN-T ステップ
decoder.mlmodelc # フォールバック (未融合)
Joint.mlmodelc # フォールバック (非融合)
tokenizer.json # 13,087 トークンの多言語語彙
metadata.json # プロンプト辞書、キャッシュ形状、vocab/b

ひょろひょろ
出典: https://huggingface.co/FluidInference/Nemotron-3.5-ASR-Streaming-Multilingual-0.6b-CoreML
モデル署名の検証 (フェーズ 4)
NemotronASRPoC/ASR/ModelSignatures.json は実際のモデルから生成され、
実行時に消費されるため、テンソル名/プロンプト ID はハードコーディングされません。ハイライト:
オーディオ: 16 kHz モノラル;プリプロセッサはオーディオ [1,?] (1 ～ 1,280,000 サンプル) を受け取ります。
Mel: 128 個の特徴、total_mel_frames = 233 (224 チャンク + 9 プリエンコード キャッシュ)。
エンコーダの状態 ( MLState ではなく、明示的なテンソル):cache_channel [1,24,42,1024]、
cache_time [1,24,1024,8]、cache_len [1] — 渡され、 *_out として返されます。
エンコーダー出力: エンコードされた [1,1024,28] (2.24 秒のチャンクあたり約 28 フレーム) + プロンプト ID 。
RNN-T ステップ ( decoder_joint ): LSTM 状態 h_in / c_in [2,1,640]、トークン [1,1]、
エンコーダ [1,1024,1] → ロジッツ [1,1,1,13088]、 h_out / c_out 。
語彙: 13,087 トークン、blank_idx = 13087 。
プロンプト ID: en-US=0、zh-CN=4、zh-TW=5、ja-JP=10、ko-KR=14、default(auto)=101。
フェーズ
範囲
州
1
XcodeGen スキャフォールド + SwiftUI シェル (開始/停止/クリア、言語 + 階層ピッカー、トランスクリプト + ベンチマーク パネル)
✅ 完了
2
マイクキャプチャ → 16 kHz モノラル → 階層化されたチャンクバッファ
✅ 完了
3
ベンチマーク計測 (ロード時間、レイテンシ p50/p90/p99、RTF、ピーク メモリ、サーマル)
✅ 完了
4
モデルのダウンロード + CoreML 署名検査 ( scripts/ 、 ModelSignatures.json )
✅ 完了
5
CoreML ロード + アダプティブ ランナー ( CoreMLModelRunner 、コンピューティング ユニット フォールバック)
✅ 完了
6
RNN-T グリーディ デコーダー + トークナイザー + ファイル トランスクリプション ( NemotronASRService )
✅ 完了
7
ライブマイクストリーミング統合 ( StreamingTranscriber ) — デバイス上で検証済み
✅ 完了
8
精度テストセット + Whisper ベースライン + ベンチマーク結果
⬜次へ
ライブ マイク ストリーミング ( StreamingTranscriber ) がリアルタイムで文字起こしできるようになりました。
物理デバイス: マイク チャンクは FIFO のキューに入れられ、

1 人の注文済み消費者によって雨が降った
( RecordingController.consume )、エンコーダー キャッシュ + RNN-T プレディクター状態のスレッド化
チャンク全体にわたって。インポートされたファイルの転写 (ファイルインポーター → transcribeManyportFile )
選択したクリップに対して同じオフライン パイプラインを実行します。両方ともデバッグされ、確認されました
コミット 6874345 で実際のハードウェアで作業中。
エンドツーエンドのファイル転写はテストとしても実行されます: EndToEndTranscriptionTests.testTranscribeMeetingClip
iOS シミュレーター (CPU/GPU) 上の会議クリップの最初の 60 秒を転写します。
RTF ≈ 0.14 。次のように実行します。
xcodebuild テスト -スキーム NemotronASRPoC \
-destination ' プラットフォーム=iOS シミュレーター、名前=iPhone 17 Pro ' \
-only-testing:NemotronASRPoCTests/EndToEndTranscriptionTests/testTranscribeMeetingClip
アーキテクチャ（現在）
AVAudioEngine ─▶ AudioResampler (16 kHz モノラル) ─▶ AudioChunkBuffer (階層調整)
│ │
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
│ (BenchmarkLogger + ASRState を駆動)
▼
SwiftUI: ContentView / TranscriptView / BenchmarkPanelView
RecordingController は両方の推論パスを駆動します。ライブマイクストリーミングの場合、オーディオ
キューはチャンクを AsyncStream に生成し、単一の Consumer(chunk:) コンシューマー ドレインを生成します。
それらは厳密に順番に StreamingTranscriber.ingest に格納され、完全な推論が実行されます。
(プリプロセッサ → エンコーダ(キャッシュ) → RNN-T デコード → トークナイザ) 主要なアクターとスレッドをオフにする
エンコーダ キャッシュ + チャンク全体の予測子の状態。モデルバンドルが存在しない場合は落ちます
チャンクごとのタイミングのみの記録に戻ります (グレースフル デグラデーション)。再入ガードがオン
start() は、ダブルタップで 2 番目の ~634 MB モデルをロードすることを防ぎます。
推論レシピ (フェーズ 5 ～ 6、実装)
NemotronASRService.transcribeFile(url: language:maxSeconds:) はフルパスを実行します。
デ

ファイルをコーディング → 16 kHz モノラル ( AudioFileLoader )、 maxSeconds にトリミング。
クリップ全体→連続メル [128, N] に対してプリプロセッサを 1 回実行します。
(プリプロセッサは最大 80 秒を受け入れます。中心にある STFT は 1 + サンプル/ホップを生成します
フレーム、例: 2240 ミリ秒のチャンクの場合は 225 — チャンクのストライドは chunk_mel_frames =224 です。)
メルを固定の 233 フレームのエンコーダ ウィンドウにスライスします ( pre_encode_cache =9 フレーム)
実際の左コンテキスト + 224 new)、224 ずつ進みます。開始の前にゼロが付きます。
最後のウィンドウはゼロパディングされ、より小さい mel_length を報告します。
NemotronEncoderRunner は 3 つのエンコーダー キャッシュをスレッド化します (cache_*→cache_*_out)。
ファイル全体に対してウィンドウ全体で実行します (ファイル内でリセットしないでください)。プロンプトID
エンコーダのみを調整します。
RNNTDecoder は、各ウィンドウのフレームを貪欲にデコードし、予測子の状態を永続化します。
(トークン + LSTM h/c) ウィンドウ全体 (発話ごとにのみリセット)、上限は 10
シンボル/フレーム。
Tokenizer はトークン化を解除します (SentencePiece →スペース、CJK には特別なケースは必要ありません)。
プロンプト ID マップ + CoreML テンソル名は実行時に次からロードされます。
ModelSignatures.json (inspect_model.py によって生成)。ハードコードされることはありません。 ✅
CoreML ログ行 [espresso] … ios17.slice_by_index: 実行中のゼロ形状エラー
エンコーダの負荷は無害な柔軟な形状の型推論ノイズです - モジュールの負荷
そして正しい [1,1024,28] 出力が生成されます。
トランスクリプトの精度はまだ測定されていません (WER/CER スコアラーや Whisper ベースラインはありません)。
それがフェーズ 8 です。60 年代の煙テストでは、正常で非変性転写物のみが確認されています。
広東語 (yue-Hant-HK) の範囲は多言語語彙によって異なります。ささやきを保つ
CER が弱い場合の計画によるフォールバックとして。
ANE の動作とメモリの上限は、デバイス (シミュレータ) でのみ検証できます。
CPU/GPU 上で CoreML を実行します (正確性は忠実ですが、ANE のパフォーマンスを表すものではありません)。
ソースコードは以下に基づいて公開されています

MIT ライセンス。これはアプリのコードのみを対象としています —
NVIDIA Nemotron-3.5-ASR モデルの重量は含まれていないため、引き続きその重量が適用されます。
ハグフェイスに関する独自のライセンス。
この PoC の開発は、次のツールによって支援されました。
WhisKey — デバイス上でのメモの素早いディクテーション、
コミットメッセージと問題の説明。
TokKong — オフライン文字起こしと
開発中の参考資料 (NVIDIA / CoreML ドキュメント、モデル カード) の翻訳。
ラウンジ — 長期にわたるビルド、テスト、および
デスクトップ上のエージェント ジョブ。
NotifyMe — 次の場合にアラートを電話にプッシュしました
長い文字起こしテストが実行され、モデルのダウンロードが完了しました。
CoreML.SwiftUI アプリ経由で NVIDIA の Nemotron-3.5-ASR ストリーミング 0.6B (多言語) を使用し、マイク キャプチャ + オーディオ ファイル インポート、RNN-T デコーディング、ライブ ベンチマーク メトリクス (レイテンシ、RTF、メモリ) を使用した、iPhone/iPad 向けのオンデバイスのオフライン音声認識。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

On-device, offline speech recognition for iPhone/iPad using NVIDIA's Nemotron-3.5-ASR Streaming 0.6B (multilingual) via CoreML.SwiftUI app with mic capture + audio file import, RNN-Tdecoding, and live benchmark metrics (latency, RTF, memory). - lbj96347/nemotron-3.5-asr-ios

GitHub - lbj96347/nemotron-3.5-asr-ios: On-device, offline speech recognition for iPhone/iPad using NVIDIA's Nemotron-3.5-ASR Streaming 0.6B (multilingual) via CoreML.SwiftUI app with mic capture + audio file import, RNN-Tdecoding, and live benchmark metrics (latency, RTF, memory). · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
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
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
lbj96347
/
nemotron-3.5-asr-ios
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits Models Models NemotronASRPoC NemotronASRPoC NemotronASRPoCTests NemotronASRPoCTests scripts scripts .gitignore .gitignore CLAUDE.md CLAUDE.md IMPLEMENTATION-PLAN.md IMPLEMENTATION-PLAN.md LICENSE LICENSE README.md README.md project.yml project.yml proposed-plan.md proposed-plan.md View all files Repository files navigation
On-device, offline, streaming speech recognition PoC using NVIDIA
Nemotron-3.5-ASR Streaming 0.6B (multilingual) via CoreML on iPhone/iPad.
Both inference paths now run end-to-end on a physical device: live mic streaming
(Phase 7) and offline file transcription. The simulator still runs the full shell
(audio + benchmark) on CPU/GPU but is not representative of ANE performance.
See IMPLEMENTATION-PLAN.md for the full engineering plan
and proposed-plan.md for the original product intent.
Xcode 16+ (built/tested with Xcode 26.3)
iPhone 15 Pro or newer recommended (ANE); device required for real inference
XcodeGen ( brew install xcodegen )
CoreML model files downloaded separately (see Model Setup )
# 1. Generate the Xcode project from project.yml (run after any project.yml change)
xcodegen generate
# 2a. Build for the simulator (shell + audio + benchmark; no real inference)
xcodebuild -scheme NemotronASRPoC -destination ' generic/platform=iOS Simulator ' build
# 2b. Or open in Xcode and run on a physical device
open NemotronASRPoC.xcodeproj
# 3. Run unit tests (resampler, chunk buffer, latency tracker)
xcodebuild test -scheme NemotronASRPoC -destination ' platform=iOS Simulator,name=iPhone 17 Pro '
# Run a single test
xcodebuild test -scheme NemotronASRPoC \
-destination ' platform=iOS Simulator,name=iPhone 17 Pro ' \
-only-testing:NemotronASRPoCTests/AudioPipelineTests/testResamplerDownsamplesTo16k
The app runs without the model present — it shows a "models not found" status
but still captures mic audio, resamples to 16 kHz, chunks to the streaming tier, and
reports live benchmark metrics. This lets you validate the shell before the download.
The file-transcription tests are gated on user-supplied clips in TestAudio/ (the
folder is gitignored — provide your own). Name each clip with a language hint so the
suite picks the right prompt, e.g. sample-en.m4a (English) or sample-yue.m4a
(Cantonese → auto ). Tests skip cleanly when no matching clip (or the model) is present.
Default target: multilingual @ 2240 ms (~634 MB; covers EN / zh / ja / ko;
Cantonese has no dedicated prompt and falls back to auto ).
# Download all five .mlmodelc bundles + tokenizer + metadata into Models/multilingual/2240ms/
./scripts/download_models.sh multilingual 2240
# Re-inspect to refresh signatures (already checked in, but re-run if you change tier)
python3 scripts/inspect_model.py Models/multilingual/2240ms --out NemotronASRPoC/ASR/ModelSignatures.json
# Re-bundle the Models/ folder into the app
xcodegen generate
Expected files in Models/multilingual/2240ms/ :
preprocessor.mlmodelc # audio[1,?] → mel[1,128,?], mel_length
encoder.mlmodelc # mel + caches + prompt_id → encoded + updated caches (stateful)
decoder_joint.mlmodelc # B1 fused decoder⊕joint — default RNN-T step
decoder.mlmodelc # fallback (unfused)
joint.mlmodelc # fallback (unfused)
tokenizer.json # 13,087-token multilingual vocab
metadata.json # prompt dictionary, cache shapes, vocab/blank
Source: https://huggingface.co/FluidInference/Nemotron-3.5-ASR-Streaming-Multilingual-0.6b-CoreML
Verified model signatures (Phase 4)
NemotronASRPoC/ASR/ModelSignatures.json is generated from the real model and
consumed at runtime so no tensor names / prompt IDs are hardcoded. Highlights:
Audio: 16 kHz mono; preprocessor takes audio [1,?] (1–1,280,000 samples).
Mel: 128 features, total_mel_frames = 233 (224 chunk + 9 pre-encode cache).
Encoder state (explicit tensors, not MLState ): cache_channel [1,24,42,1024],
cache_time [1,24,1024,8], cache_len [1] — passed in and returned as *_out .
Encoder output: encoded [1,1024,28] (≈28 frames per 2.24 s chunk) + prompt_id .
RNN-T step ( decoder_joint ): LSTM state h_in / c_in [2,1,640], token [1,1],
encoder [1,1024,1] → logits [1,1,1,13088], h_out / c_out .
Vocab: 13,087 tokens, blank_idx = 13087 .
Prompt IDs: en-US=0, zh-CN=4, zh-TW=5, ja-JP=10, ko-KR=14, default(auto)=101.
Phase
Scope
State
1
XcodeGen scaffold + SwiftUI shell (Start/Stop/Clear, language + tier pickers, transcript + benchmark panel)
✅ Done
2
Mic capture → 16 kHz mono → tier-aligned chunk buffer
✅ Done
3
Benchmark instrumentation (load time, latency p50/p90/p99, RTF, peak memory, thermal)
✅ Done
4
Model download + CoreML signature inspection ( scripts/ , ModelSignatures.json )
✅ Done
5
CoreML loading + adaptive runner ( CoreMLModelRunner , compute-unit fallback)
✅ Done
6
RNN-T greedy decoder + tokenizer + file transcription ( NemotronASRService )
✅ Done
7
Live mic streaming integration ( StreamingTranscriber ) — verified on device
✅ Done
8
Accuracy test set + Whisper baseline + benchmark results
⬜ Next
Live mic streaming ( StreamingTranscriber ) now transcribes in real time on a
physical device: mic chunks are queued FIFO and drained by a single in-order consumer
( RecordingController.consume ), threading the encoder caches + RNN-T predictor state
across chunks. Imported file transcription (Files importer → transcribeImportedFile )
runs the same offline pipeline on a picked clip. Both were debugged and confirmed
working on real hardware in commit 6874345 .
End-to-end file transcription also runs as a test: EndToEndTranscriptionTests.testTranscribeMeetingClip
transcribes the first 60 s of the meeting clip on the iOS Simulator (CPU/GPU) at
RTF ≈ 0.14 . Run it with:
xcodebuild test -scheme NemotronASRPoC \
-destination ' platform=iOS Simulator,name=iPhone 17 Pro ' \
-only-testing:NemotronASRPoCTests/EndToEndTranscriptionTests/testTranscribeMeetingClip
Architecture (current)
AVAudioEngine ─▶ AudioResampler (16 kHz mono) ─▶ AudioChunkBuffer (tier-aligned)
│ │
└──────────────▶ RecordingController ◀─────────────┘
│ (drives BenchmarkLogger + ASRState)
▼
SwiftUI: ContentView / TranscriptView / BenchmarkPanelView
RecordingController drives both inference paths. For live mic streaming, the audio
queue yields chunks into an AsyncStream and a single consume(chunk:) consumer drains
them strictly in order into StreamingTranscriber.ingest , which runs the full inference
(preprocessor → encoder(cache) → RNN-T decode → tokenizer) off the main actor and threads
encoder caches + predictor state across chunks. When the model bundle is absent it falls
back to recording per-chunk timing only (graceful degradation). A re-entrancy guard on
start() prevents a double tap from loading a second ~634 MB model.
Inference recipe (Phases 5–6, implemented)
NemotronASRService.transcribeFile(url:language:maxSeconds:) runs the full path:
Decode the file → 16 kHz mono ( AudioFileLoader ), trimmed to maxSeconds .
Run the preprocessor once on the whole clip → a continuous mel [128, N] .
(The preprocessor accepts up to ~80 s; a centered STFT yields 1 + samples/hop
frames, e.g. 225 for a 2240 ms chunk — the chunk stride is chunk_mel_frames =224.)
Slice the mel into fixed 233-frame encoder windows ( pre_encode_cache =9 frames
of real left context + 224 new), advancing by 224. Zeros precede the start;
the final window is zero-padded and reports a smaller mel_length .
NemotronEncoderRunner threads the three encoder caches ( cache_* → cache_*_out )
across windows for the whole file (never reset within a file); prompt_id
conditions the encoder only.
RNNTDecoder greedy-decodes each window's frames, persisting predictor state
(token + LSTM h/c) across windows (reset only per utterance), capped at 10
symbols/frame.
Tokenizer detokenizes (SentencePiece ▁ →space; CJK needs no special case).
prompt_id map + CoreML tensor names are loaded at runtime from
ModelSignatures.json (generated by inspect_model.py ), never hardcoded. ✅
The CoreML log line [espresso] … ios17.slice_by_index: zero shape error during
encoder load is benign flexible-shape type-inference noise — the module loads
and produces correct [1,1024,28] output.
Transcript accuracy isn't measured yet (no WER/CER scorer or Whisper baseline);
that's Phase 8. The 60 s smoke test only asserts a sane, non-degenerate transcript.
Cantonese ( yue-Hant-HK ) coverage depends on the multilingual vocab; keep Whisper
as a fallback per the plan if CER is weak.
ANE behavior and the memory ceiling can only be validated on device (the Simulator
runs CoreML on CPU/GPU — correctness-faithful but not representative of ANE perf).
Source code is released under the MIT License . This covers the app code only —
the NVIDIA Nemotron-3.5-ASR model weights are not included and remain subject to their
own license on Hugging Face .
Development of this PoC was assisted by these tools:
WhisKey — quick on-device dictation of notes,
commit messages, and issue descriptions.
TokKong — offline transcription and
translation of reference material (NVIDIA / CoreML docs, model cards) during development.
Lounge — surfaced long-running build, test, and
agent jobs on the desktop.
NotifyMe — pushed alerts to phone when
long transcription test runs and model downloads finished.
On-device, offline speech recognition for iPhone/iPad using NVIDIA's Nemotron-3.5-ASR Streaming 0.6B (multilingual) via CoreML.SwiftUI app with mic capture + audio file import, RNN-Tdecoding, and live benchmark metrics (latency, RTF, memory).
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
