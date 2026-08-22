---
source: "https://github.com/hsandhu/agent"
hn_url: "https://news.ycombinator.com/item?id=49401506"
title: "An iOS app that runs AI agents and a complete voice pipeline on the device"
article_title: "GitHub - hsandhu/agent · GitHub"
image: "https://opengraph.githubassets.com/f31e6776ec1d22158e7feded3729a9d28d872ce9e4d583ba177cfd4176966152/hsandhu/agent"
author: "rsandhu"
captured_at: "2026-08-22T17:12:13Z"
capture_tool: "hn-digest"
hn_id: 49401506
score: 1
comments: 0
posted_at: "2026-08-22T16:52:38Z"
tags:
  - hacker-news
  - translated
---

# An iOS app that runs AI agents and a complete voice pipeline on the device

- HN: [49401506](https://news.ycombinator.com/item?id=49401506)
- Source: [github.com](https://github.com/hsandhu/agent)
- Score: 1
- Comments: 0
- Posted: 2026-08-22T16:52:38Z

## Translation

タイトル: デバイス上で AI エージェントと完全な音声パイプラインを実行する iOS アプリ
記事タイトル: GitHub - hsandhu/agent · GitHub
説明: GitHub でアカウントを作成して、hsandhu/agent の開発に貢献します。

記事本文:
GitHub - hsandhu/エージェント · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
サンドゥ
/
エージェント
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
6 コミット 6 コミット フォルダーとファイル
Agent.xcodeproj Agent.xcodeproj Agent Agent ドキュメント/スクリーンショット ドキュメント/スクリーンショット スクリプト スクリプト .gitignore .gitignore README.md README.md project.yml project.yml すべてのファイルを表示 リポジトリ ファイル ナビゲーション
AI エージェントと完全な機能を実行する iOS アプリ

音声パイプライン全体が
デバイス。アカウントもモデル API もクラウドでの推論もありません。たった一つのこと
電話から離れるのはウェブ調査です。エージェントはページを検索してダウンロードします。
したがって、オンデバイスモデルには、推論すべき最新の何かがあり、そのスイッチは
オフにすることができます。
お互いを強化する 2 つの部分:
バックグラウンドエージェント。タスクについて説明します（「自分にとって最適なサマーキャンプを見つける）」
シカゴ近郊の 9 歳」)、エージェントがバックグラウンドで作業を行っています。
検索クエリを計画し、見つかったページを読み取り、それを推論します。
Apple Intelligence のオンデバイス モデルと、それが使用したものを引用します。戻ってきて
後で再生を押すと、結果が読み上げられます。
デバイス上のスピーチ。ストリーミング音声テキスト変換 (ディクテーションの両方に使用)
タスクおよびライブ文字起こし用）およびゼロショット音声クローン作成
テキスト読み上げ。エージェントの結果に音声を与えるものです (オプション)
10 秒のサンプルから複製されたあなたの声。
スピーチの半分は、長期的な地域の構成要素としても存在します。
アイデア: 音声を送信者のデバイスで文字起こしして送信する音声通話アプリ
ネットワーク上でテキストとして送信 (Opus の 6 ～ 24 kbps ではなく、約 50 バイト/秒)、および
受信者のデバイス上で送信者のクローン音声が再合成されます。の
ネットワーク層はまだ構築されていません。ライブ文字起こし画面のエコー
toggle は、そのパイプラインのローカル ループバックです。
5層。 UI はどのモデルでもネイティブ ランタイムに直接触れることはありません。
呼び出しはメインスレッド外で行われ、モデルはネットワークに到達しません。
それ自体 — Web 層がフェッチを実行し、テキストを渡します。
フローチャートTB
サブグラフ UI["UI — SwiftUI"]
A["AgentsView<br/>リスト · 作曲者 · 詳細"]
S[「設定表示<br/>マイボイス・文字起こし・話す」]
D["DesignSystem<br/>共有コンポーネント"]
終わり
サブグラフ ドメイン["エージェント — ドメイン"]
R["エージェントランナー<br/>BG処理

タスク + フォアグラウンド"]
B["エージェントブレイン<br/>プロトコル"]
ST[("エージェントストア<br/>モデルアクター")]
J["エージェントジョブ<br/>SwiftData モデル"]
終わり
サブグラフ Web[「ウェブ調査」]
WR["WebResearcher<br/>検索・閲覧・抜粋"]
SP["WebSearchProvider<br/>DuckDuckGo・Brave"]
PR["ページリーダー<br/>+ HTMLText"]
終わり
サブグラフ Speech["Speech — エンジン"]
STT["SttEngine<br/>シリアルキュー"]
TTS["TtsEngine<br/>シリアルキュー"]
CAP["オーディオキャプチャ"]
PLAY["オーディオプレーヤー"]
VP["音声プロファイルストア"]
終わり
サブグラフ Native["ネイティブ ブリッジ"]
W["SherpaOnnx.swift<br/>+ Zipvoice 拡張機能"]
C["ブリッジヘッダー経由のc-api.h"]
X["sherpa-onnx.xcframework<br/>onnxruntime.xcframework"]
終わり
A --> R
A --> TTS
A --> STT
S --> STT
S --> TTS
S --> 副社長
S --> WR
R --> B
R --> ST
R --> WR
B --> WR
WR → SP
WR --> PR
ST --- J
B -.->|iOS 26+| FM[「財団モデル」]
B -.->|フォールバック| MOCK["モックエージェントブレイン"]
STT --> キャップ
TTS --> プレイ
TTS --> 副社長
STT --> W
TTS --> W
W --> C --> X
読み込み中
エージェントの実行
エージェントは、ステータス (キューに入っている → 実行中 → 完了/失敗)、動作に応じて追加される進行状況ログ、および結果の分割を含む SwiftData 内の行です。
音声形式の resultsummary と長い resultDetail に変換します。
AgentRunner は起動時に構成されたシングルトンです (BGTaskScheduler には
起動が完了する前の登録）。 2 つのエントリからキューを排出します
ポイント:
割り込みは明示的に処理されます。タスクの ExpirationHandler が割り込みをキャンセルします。
作業中、drainQueue は cancelError をキャッチし、実行中のジョブを再度キューに入れます。
プロセスの停止により実行中に保留されたジョブは、次のキューに再キューされます。
起動 ( requeueOrphanedRunningJobs )。
エージェントの背後にあるモデルは AgentBrain プロトコルを介して交換可能であるため、
永続性、スケジューリング、および再生機構はどのモデルにも依存しません
実行します:
FoundationModelsBrain — Apple Intelligence のオンデバイス モデル、コンパイル済み
で

#if canImport(FoundationModels) の背後にあり、次の場合にのみ提供されます
SystemLanguageModel.default.availability レポートが利用可能です。 3本走る
パス: 検索クエリを計画し、読み取られた内容から結果を書き込みます。
声に出して読むために書かれた 3 文の要約。
MockAgentBrain — Apple Intelligence が機能しないフォールバック
利用可能です。ウェブリサーチを行っても、実際の情報を検索して読み取ります。
発見したもののダイジェストを報告します。研究をオフにして、ステージングをシミュレートします
研究のため、パイプライン全体をどのデバイスでもテストできます。
エージェントは考える前に検索します。モデルがネットワーク自体に触れることはありません
— アプリが検索を実行し、ダウンロードするページを決定し、応答を返します。
抜粋 — したがって、モデルが認識できるものは、そのポリシーではなくポリシーによって制限されます。
独自のツール呼び出し。
クエリの計画 → 検索 → インターリーブ + 重複排除 → N ページの読み取り → 抜粋 → グラウンド
ステップ
どこで
注意事項
クエリを計画する
FoundationModelsBrain.planクエリ
ガイド付き生成 ( @Generable ) では、クエリ文字列のリストが強制されます。フリーテキストのままにし、小さなモデルがクエリを作成する代わりにタスクに応答します。WebResearcher.normalize はその習慣の残り物 (マークダウン、ダッシュ句、長すぎる行) をスクラブします。
検索
ウェブ検索プロバイダー
DuckDuckGoSearch (キーレス) または BraveSearch (キーチェーン内の API キー)
マージ
ウェブリサーチャー
クエリ間でラウンドロビンを実行するため、1 つのクエリが読み取りバジェットを独占できなくなり、正規のホスト + パスによって重複排除が行われます。
読む
ページリーダー + HTML テキスト
ページあたり 1.2 MB に制限され、<article> / <main> が優先され、タグとエンティティはプレーン テキストに削減されます。読み込まれないページは、ジョブが失敗するのではなく、検索スニペットに劣化します。
地面
FoundationModelsBrain.findings
プロンプト内の番号付きの抜粋。インラインでの引用が必要です。コンテキスト ウィンドウのオーバーフローでは、ソースごとに 1100 → 600 → 300 文字で再試行し、その後 f

すべては研究されていない一般知識に戻ります
ソースはジョブ (sourcesJSON) 上に保持され、タップ可能なリンクとしてリストされます。
調査結果の下にあるため、すべての申し立てをそのページに遡って追跡できます
から。設定 › Web Research はマスター スイッチ、プロバイダーの選択を保持します。
デプスノブとライブテストボタン。
AudioCapture は AVAudioEngine を利用してハードウェア フォーマットを変換します
(44.1/48 kHz) からモデルが期待する 16 kHz モノラル Float32 まで。その出力フィード
認識機能または登録レコーダーのいずれか — 両方を使用することはできません。
1 つのキャプチャ セッション。
SttEngine は、ストリーミング Zipformer トランスデューサをエンドポイント検出でラップします。
ライブパーシャルを公開し、それぞれに最終的な TranscriptSegment を追加します。
一時停止を検出しました。ディクテーション モードがあります — when dictationOnPartial /
dictationOnUtterance が設定され、それらのコールバックへの音声ルートが認識されます
主要なトランスクリプトの代わりに、これが作曲家の薬が話し言葉を受け取る方法です
文字起こし画面を汚さずに入力できます。
TtsEngine は、シリアル キュー上で ZipVoice ゼロショット合成を実行します。音声クローン
トレーニングは必要ありません。VoiceProfile は単なる参照 wav にそのファイルを加えたものです。
トランスクリプト、および呼び出し時のそのペアの合成条件。プロンプトオーディオ
ファイルパスごとにキャッシュされます。登録では毎回新しいファイル名が書き込まれるため、
キャッシュが古くなることはありません。
コンポーネント
執行者
ビュー、@Published 状態
メイン
AudioCapture コールバック
AVAudioEngine レンダリング スレッド
SttEngine デコード ループ
プライベートシリアル DispatchQueue
Ttsエンジン合成
プライベートシリアルDispatchQueue
AgentStore (すべての SwiftData 書き込み)
@ModelActor
AgentRunner.drainQueue
Swift Task 、ストアアクターを待機します
SwiftData コンテキストはスレッドセーフではないため、バックグラウンドからのすべての変更は
タスクは AgentStore を経由します。 UI は @Query を独自に読み取ります
メインスレッドのコンテキスト。
Swift パッケージ Ma はありません

nager または CocoaPods の依存関係。その人
サードパーティのランタイムは、事前に構築された .xcframework バイナリとして販売されており、
プロジェクト ファイルはコミットされずに生成されます。
ネイティブ ランタイム (ベンダー、ベンダー/)
パッケージ
バージョン
ライセンス
役割
シェルパオンクス
1.12.21
アパッチ2.0
音声ランタイム - ストリーミング ASR + オフライン TTS (ZipVoice ゼロショット クローン作成を含む)。 csukuangfj/sherpa-onnx-libs から事前に構築された iOS xcframework
ONNX ランタイム
1.17.1
マサチューセッツ工科大学
ニューラル ネットワーク推論 (CPU 実行プロバイダー)。 sherpa-onnx iOS リリースに同梱
1.12.21 に固定されているのは、ビルド済み iOS を備えた最新バージョンであるためです
フレームワークが公開されました — 1.13.x には iOS ビルドがありません。 Swift ラッパー、C
ヘッダーとバイナリはすべて同じリリースからのものである必要があります。
モデル (ダウンロード、ベンダー/モデル/)
モデル
サイズ
ライセンス/トレーニングデータ
役割
ストリーミング Zipformer トランスデューサー、英語 20M (int8 エンコーダー + fp32 デコーダー + int8 ジョイナー)
34MB
Apache 2.0 · LibriSpeech
音声テキスト変換のストリーミング
ZipVoice-Distill int8、zh-en ( sherpa-onnx-zipvoice-distill-int8-zh-en-emilia )
126 MB + 18 MB のピークデータ
Apache 2.0 · エミリア
ゼロショット音声クローン TTS (123M パラメータ フローマッチング モデル)
vocos 24 kHz ボコーダー
52MB
アパッチ2.0
メルスペクトログラム→ZipVoice用波形
デモ参考音声
212KB
LibriSpeech サンプル
音声が組み込まれているため、登録前に TTS が機能します
モデルの合計は 229 MB で、インストールされるアプリは約 257 MB になります。ベンダー/
これは gitignored であり、 script/fetch-deps.sh から完全に再現可能です。
フレームワーク
用途
SwiftUI
UI全体
スイフトデータ
エージェントの永続性 ( @Model 、 @Query 、 @ModelActor )
バックグラウンドタスク
BGProcessingタスクのスケジューリングと実行
AV財団
マイクキャプチャ、フォーマット変換、PCM再生、オーディオセッション
基礎モデル
Apple Intelligence オンデバイス LLM — 弱い、条件付き ( #if canImport )、iOS 26 以降
URLセッション
ウェブ検索

およびページの取得 (一時的なセッション、Cookie またはキャッシュの永続性なし)
セキュリティ
オプションの Brave API キー用のキーチェーン ストレージ
結合する
ObservableObject エンジン
ビルドツール
ツール
役割
XcodeGen
project.yml から Agent.xcodeproj を生成します — プロジェクト ファイルは使い捨てであり、project.yml が信頼できる情報源です
導入対象はiOS 17.0、Swift 5.9、iPhoneのみ。 C API に到達しました
Objective-C ブリッジング ヘッダー経由
( Agent/Support/SherpaOnnx-Bridging-Header.h ) -lc++ がリンクされています。
./scripts/fetch-deps.sh
xcodegen を生成し、Agent.xcodeproj を開きます
フェッチ スクリプトはフレームワークとモデルをダウンロードします (最大 330 MB のアーカイブ)
は、vendor/ に変換され、冪等です。すでに存在するものはすべてスキップされます。
エージェント/
AgentApp.swift アプリのエントリ。 ModelContainer を構築し、AgentRunner を登録し、
TTS モデルをウォーミングし、相変化に関するバックグラウンド作業をスケジュールします。
ContentView.swift ルート (AgentsView)
エージェント/
AgentJob.swift @Model: ステータス、プロンプト、タイムスタンプ、進行状況ログ、結果
AgentStore.swift @ModelActor: すべてのバックグラウンド SwiftData アクセス
AgentRunner.swift BGProcessingTask 登録/スケジューリング + ワークループ
AgentBrain.swift Brain プロトコル、MockAgentBrain、FoundationModelsBrain
ウェブ/
WebSearch.swift WebResult/WebSource、プロバイダー プロトコル、WebSearchConfig、キーチェーン
WebResearcher.swift オーケストレーション:

[切り捨てられた]

## Original Extract

Contribute to hsandhu/agent development by creating an account on GitHub.

GitHub - hsandhu/agent · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
hsandhu
/
agent
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
6 Commits 6 Commits Folders and files
Agent.xcodeproj Agent.xcodeproj Agent Agent docs/ screenshots docs/ screenshots scripts scripts .gitignore .gitignore README.md README.md project.yml project.yml View all files Repository files navigation
An iOS app that runs AI agents and a complete voice pipeline entirely on the
device . No account, no model API, no inference in the cloud. The one thing
that leaves the phone is web research — the agent searches and downloads pages
so the on-device model has something current to reason over, and that switch
can be turned off.
Two halves that reinforce each other:
Background agents. Describe a task ("find the best summer camps for my
9-year-old near Chicago"), and an agent works on it in the background: it
plans search queries, reads the pages it finds, then reasons over them with
Apple Intelligence's on-device model and cites what it used. Come back
later and press play — the result is read aloud.
On-device speech. Streaming speech-to-text (used both for dictating
tasks and for live transcription) and zero-shot voice-cloning
text-to-speech, which is what gives agent results a voice — optionally
your voice, cloned from a ten-second sample.
The speech half also exists as the local building block for a longer-term
idea: a voice-call app where audio is transcribed on the sender's device, sent
over the network as text (~50 bytes/sec instead of 6–24 kbps of Opus), and
re-synthesized on the receiver's device in the sender's cloned voice. The
networking layer isn't built yet; the Live Transcription screen's Echo
toggle is a local loopback of that pipeline.
Five layers. The UI never touches the native runtime directly, every model
call happens off the main thread, and the model never reaches the network
itself — the web layer does the fetching and hands it text.
flowchart TB
subgraph UI["UI — SwiftUI"]
A["AgentsView<br/>list · composer · detail"]
S["SettingsView<br/>My Voice · Transcribe · Speak"]
D["DesignSystem<br/>shared components"]
end
subgraph Domain["Agents — domain"]
R["AgentRunner<br/>BGProcessingTask + foreground"]
B["AgentBrain<br/>protocol"]
ST[("AgentStore<br/>ModelActor")]
J["AgentJob<br/>SwiftData Model"]
end
subgraph Web["Web research"]
WR["WebResearcher<br/>search · read · excerpt"]
SP["WebSearchProvider<br/>DuckDuckGo · Brave"]
PR["PageReader<br/>+ HTMLText"]
end
subgraph Speech["Speech — engines"]
STT["SttEngine<br/>serial queue"]
TTS["TtsEngine<br/>serial queue"]
CAP["AudioCapture"]
PLAY["AudioPlayer"]
VP["VoiceProfileStore"]
end
subgraph Native["Native bridge"]
W["SherpaOnnx.swift<br/>+ Zipvoice extension"]
C["c-api.h via bridging header"]
X["sherpa-onnx.xcframework<br/>onnxruntime.xcframework"]
end
A --> R
A --> TTS
A --> STT
S --> STT
S --> TTS
S --> VP
S --> WR
R --> B
R --> ST
R --> WR
B --> WR
WR --> SP
WR --> PR
ST --- J
B -.->|iOS 26+| FM["FoundationModels"]
B -.->|fallback| MOCK["MockAgentBrain"]
STT --> CAP
TTS --> PLAY
TTS --> VP
STT --> W
TTS --> W
W --> C --> X
Loading
Agent execution
An agent is a row in SwiftData with a status ( queued → running → completed/failed ), a progress log appended as it works, and a result split
into a spoken-style resultSummary and a longer resultDetail .
AgentRunner is a singleton configured at launch (BGTaskScheduler requires
registration before launch finishes). It drains the queue from two entry
points:
Interruption is handled explicitly: the task's expirationHandler cancels the
work, drainQueue catches CancellationError and requeues the in-flight job,
and any job left stranded in running by a process death is requeued at next
launch ( requeueOrphanedRunningJobs ).
The model behind an agent is swappable via the AgentBrain protocol, so the
persistence, scheduling, and playback machinery is independent of which model
runs:
FoundationModelsBrain — Apple Intelligence's on-device model, compiled
in behind #if canImport(FoundationModels) and offered only when
SystemLanguageModel.default.availability reports available. Runs three
passes: plan the search queries, write the findings from what was read, then
a three-sentence summary written to be read aloud.
MockAgentBrain — the fallback where Apple Intelligence isn't
available. With web research on it still searches and reads for real and
reports a digest of what it found; with research off it simulates staged
research, so the whole pipeline is testable on any device.
Agents search before they think. The model never touches the network itself
— the app runs the searches, decides which pages to download, and hands back
excerpts — so what the model can see is bounded by policy rather than by its
own tool calls.
plan queries → search → interleave + dedupe → read N pages → excerpt → ground
Step
Where
Notes
Plan queries
FoundationModelsBrain.planQueries
Guided generation ( @Generable ) forces a list of query strings. Left to free text, a small model answers the task instead of writing queries for it — WebResearcher.normalize scrubs the leftovers of that habit (markdown, dash clauses, over-long lines)
Search
WebSearchProvider
DuckDuckGoSearch (keyless) or BraveSearch (API key in the keychain)
Merge
WebResearcher
Round-robins across queries so one query can't monopolize the read budget, and dedupes by canonical host+path
Read
PageReader + HTMLText
Capped at 1.2 MB per page, <article> / <main> preferred, tags and entities reduced to plain text. A page that won't load degrades to its search snippet rather than failing the job
Ground
FoundationModelsBrain.findings
Numbered excerpts in the prompt, citations required inline. On a context-window overflow it retries at 1100 → 600 → 300 chars per source, then falls back to unresearched general knowledge
Sources are persisted on the job ( sourcesJSON ) and listed as tappable links
under the findings, so every claim can be traced back to the page it came
from. Settings › Web Research holds the master switch, the provider choice,
the depth knobs, and a live test button.
AudioCapture taps AVAudioEngine and converts the hardware format
(44.1/48 kHz) to the 16 kHz mono Float32 the models expect. Its output feeds
either the recognizer or the enrollment recorder — never both, since there's
one capture session.
SttEngine wraps a streaming Zipformer transducer with endpoint detection: it
publishes a live partial and appends a finalized TranscriptSegment on each
detected pause. It has a dictation mode — when dictationOnPartial /
dictationOnUtterance are set, recognized speech routes to those callbacks
instead of the main transcript, which is how the composer pill takes spoken
input without polluting the transcription screen.
TtsEngine runs ZipVoice zero-shot synthesis on a serial queue. Voice cloning
needs no training: a VoiceProfile is just a reference wav plus its
transcript, and synthesis conditions on that pair at call time. Prompt audio
is cached per file path; enrollment writes a new filename each time so the
cache can never go stale.
Component
Executor
Views, @Published state
Main
AudioCapture callbacks
AVAudioEngine render thread
SttEngine decode loop
Private serial DispatchQueue
TtsEngine synthesis
Private serial DispatchQueue
AgentStore (all SwiftData writes)
@ModelActor
AgentRunner.drainQueue
Swift Task , awaits the store actor
SwiftData contexts aren't thread-safe, so every mutation from a background
task goes through AgentStore ; the UI reads through @Query on its own
main-thread context.
There are no Swift Package Manager or CocoaPods dependencies. The one
third-party runtime is vendored as prebuilt .xcframework binaries, and the
project file is generated rather than committed.
Native runtime (vendored, vendor/ )
Package
Version
License
Role
sherpa-onnx
1.12.21
Apache 2.0
Speech runtime — streaming ASR + offline TTS, including ZipVoice zero-shot cloning. Prebuilt iOS xcframework from csukuangfj/sherpa-onnx-libs
ONNX Runtime
1.17.1
MIT
Neural network inference (CPU execution provider). Ships inside the sherpa-onnx iOS release
Pinned to 1.12.21 because that's the newest version with prebuilt iOS
frameworks published — 1.13.x has no iOS build. The Swift wrapper, C
headers, and binaries must all come from the same release.
Models (downloaded, vendor/models/ )
Model
Size
License / training data
Role
Streaming Zipformer transducer, English 20M (int8 encoder + fp32 decoder + int8 joiner)
34 MB
Apache 2.0 · LibriSpeech
Streaming speech-to-text
ZipVoice-Distill int8, zh-en ( sherpa-onnx-zipvoice-distill-int8-zh-en-emilia )
126 MB + 18 MB espeak-ng data
Apache 2.0 · Emilia
Zero-shot voice-cloning TTS (123M-param flow-matching model)
vocos 24 kHz vocoder
52 MB
Apache 2.0
Mel spectrogram → waveform for ZipVoice
Demo reference voice
212 KB
LibriSpeech sample
Built-in voice so TTS works before you enroll
229 MB of models in total, which puts the installed app at ~257 MB. vendor/
is gitignored and fully reproducible from scripts/fetch-deps.sh .
Framework
Used for
SwiftUI
Entire UI
SwiftData
Agent persistence ( @Model , @Query , @ModelActor )
BackgroundTasks
BGProcessingTask scheduling and execution
AVFoundation
Mic capture, format conversion, PCM playback, audio session
FoundationModels
Apple Intelligence on-device LLM — weak, conditional ( #if canImport ), iOS 26+
URLSession
Web search and page fetching (ephemeral session, no cookie or cache persistence)
Security
Keychain storage for the optional Brave API key
Combine
ObservableObject engines
Build tooling
Tool
Role
XcodeGen
Generates Agent.xcodeproj from project.yml — the project file is disposable, project.yml is the source of truth
Deployment target iOS 17.0, Swift 5.9, iPhone only. The C API is reached
through an Objective-C bridging header
( Agent/Support/SherpaOnnx-Bridging-Header.h ) with -lc++ linked.
./scripts/fetch-deps.sh
xcodegen generate && open Agent.xcodeproj
The fetch script downloads the frameworks and models (~330 MB of archives)
into vendor/ , and is idempotent — it skips anything already present.
Agent/
AgentApp.swift App entry; builds ModelContainer, registers AgentRunner,
warms the TTS model, schedules background work on phase change
ContentView.swift Root (AgentsView)
Agents/
AgentJob.swift @Model: status, prompt, timestamps, progress log, result
AgentStore.swift @ModelActor: all background SwiftData access
AgentRunner.swift BGProcessingTask registration/scheduling + the work loop
AgentBrain.swift Brain protocol, MockAgentBrain, FoundationModelsBrain
Web/
WebSearch.swift WebResult/WebSource, provider protocol, WebSearchConfig, keychain
WebResearcher.swift Orchestration:

[truncated]
