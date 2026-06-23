---
source: "https://github.com/attenlabs/saa-sdk"
hn_url: "https://news.ycombinator.com/item?id=48649105"
title: "Show HN: We Help Voice AI Handle Group Conversations"
article_title: "GitHub - attenlabs/saa-sdk: The addressee layer for voice agents: only speech meant for your agent reaches your STT, LLM, or TTS, on any stack. · GitHub"
author: "betweenDan"
captured_at: "2026-06-23T18:38:41Z"
capture_tool: "hn-digest"
hn_id: 48649105
score: 3
comments: 0
posted_at: "2026-06-23T18:20:00Z"
tags:
  - hacker-news
  - translated
---

# Show HN: We Help Voice AI Handle Group Conversations

- HN: [48649105](https://news.ycombinator.com/item?id=48649105)
- Source: [github.com](https://github.com/attenlabs/saa-sdk)
- Score: 3
- Comments: 0
- Posted: 2026-06-23T18:20:00Z

## Translation

タイトル: HN を表示: 音声 AI によるグループ会話の処理を支援します
記事のタイトル: GitHub - attenlabs/saa-sdk: 音声エージェントの宛先層: エージェントに向けられた音声のみが、どのスタックでも STT、LLM、または TTS に到達します。 · GitHub
説明: 音声エージェントの受信者層: エージェント宛ての音声のみが、スタック上で STT、LLM、または TTS に到達します。 - attenlabs/saa-sdk
HN テキスト: 皆さん。私たちは、複数のロボット/複数のエージェントで優れたエクスペリエンスを実現する方法を模索した結果、SAA (選択的聴覚注意) を構築しました。典型的には、彼らが決して話をやめないということが起こりました。これは、STT の前に配置できる SDK です。ウェイクワードなしで、デバイスに話しかけられているかどうかを知ることができます。次の用途に使用できます。
- 単一の AI、複数の人間
-複数のAI、単一の人間
-マルチAI、マルチヒューマン（より良いシステムのためにウェイクワードを追加することもお勧めします） 2つのモデルがあります。 1 つはビデオ + オーディオ、もう 1 つはオーディオのみです。全体的な仕組みとしては、注意のパターン (ボディランゲージの変化や音声のパターン) の変化を探して機能するようになっています。人間やデバイスとの関わり方は人それぞれ異なるため、これを解決するのは難しい問題です。様子を教えてください!

記事本文:
GitHub - attenlabs/saa-sdk: 音声エージェントの宛先層: エージェントに向けられた音声のみが、どのスタックでも STT、LLM、または TTS に到達します。 · GitHub
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
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
アテンラボ
/
saa-sdk
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
221 コミット 221 コミット .github .github アセット アセット サンプル サンプル パッケージ パッケージ .gitignore .gitignore CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンスMakefile Makefile NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md llms-full.txt llms-full.txt llms.txt llms.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
実際にどの音声が該当するかを音声エージェントに伝えます。
発話ごとに 1 つの決定: 宛先の発話のみが STT、LLM、および TTS に届きます。ウェイクワードはありません。
すでに使用している音声エージェント スタックのドロップイン:
音声エージェントのマイクは、部屋内のあらゆる声を聞きます。あなたの声、同僚の声、子供たちの声、ラップトップで再生中のポッドキャスト、スピーカーから流れ出るエージェント自身の TTS などです。ほとんどのパイプラインはそのいずれかに応答し、文字起こしされた 1 秒ごとに STT を支払い、デバイスに向けられていない音声に対して LLM をトリガーします。
SAA (選択的聴覚注意) は、STT の前に実行され、発話ごとに、音声がデバイスに向けられたものであるかどうかを判断するホストされた分類子です。サイド トーク、バックグラウンド メディア、およびエージェント自身の再生はフィルターで除外されるため、STT / LLM / TTS にはエージェント向けの音声のみが表示されます。
ウェイクワードはありません。 SAA は、音声 (およびオプションで低レートのビデオ) ストリームから発話ごとに決定します。
主催。注意研究所のクラウドへのリアルタイム WebSocket。オープン SDK はシン クライアントです。 STT の前にゲートするため、アドレス指定された音声のみがすでに実行されている STT、LLM、および TTS に到達するため、ダウンストリーム サービスとログに表示される音声は少なくなり、多くは表示されなくなります。オンデバイス展開は別のエンタープライズ ライセンスです。
アーキテクチャと評価

使用方法はテクニカルレポートに記載されています。
形状
パッケージ
こんなときに使います
ストリーミングSDK
@attenlabs/saa-js 、attenlabs-saa
アプリはオーディオ/ビデオ自体をキャプチャしており、型指定されたアテンション イベントを独自のパイプラインにゲートする必要があるとします。 Web エージェント、モールのキオスク、ドライブスルー エージェント、ロボットに適しています。
ライブキット
saa-livekit-クライアント
LiveKit Agents 音声エージェントを実行します。 SAA がルームに参加し、セッションをゲートします。
パイプキャット (毎日)
saa-pipecat-クライアント
Daily で Pipecat 音声エージェントを実行します。 SAA はデイリー ルームに参加し、「saa」アプリ メッセージ トピックを介してパイプラインをゲートします。
イレブンラボ
attenlabs-saa
あなたは、イレブンラボの会話型 AI エージェントを実行します。 SAA はストリーミング SDK の feed_audio を介してゲートします (部屋は密閉されているため、SAA は直接参加できません)。
トゥイリオ
attenlabs-saa
Twilio Media Streams テレフォニー エージェントを実行するとします。 SAA は、ストリーミング SDK の feed_audio を介して、着信/発信通話音声 (PCM16 にリサンプリングされた μ-law 8 kHz) をゲートします。
インストール
npm install @attenlabs/saa-js # JavaScript / ブラウザ
pip install attenlabs-saa # Python (ストリーミング SDK)
pip install saa-livekit-client # Python (LiveKit)
pip install saa-pipecat-client # Python (Pipecat on Daily)
notelabs.ai で API キーを取得します。
あなたはメディアをキャプチャします。 SAA は型付きイベントを発行します。重要なイベントは、turnReady /turn_ready です。これは、デバイス向けの 1 つの発話であり、キャプチャされて STT または LLM に転送する準備ができています。
"@attenlabs/saa-js" から { AttendantClient } をインポートします。
const client = new AttendantClient({トークン:プロセス.環境.SAA_API_KEY});
// デバイス主導のターンごとに 1 回起動します。 turn.audioBase64 は PCM16 @ 16 kHz
クライアント。 on ( "turnReady" , (turn ) => yourSTT . send (turn .audioBase64 ) ) ;
クライアントを待ちます。 start ( { videoElement : document . querySelector ( "video" ) } ) ;
OSをインポートする
saa インポート AttendantClient から
client = AttendantClient ( token = os . en

viron [ "SAA_API_KEY" ])
@クライアント 。オンターン準備完了
def _ (ターン):
#turn.audio_base64、PCM16 @ 16 kHz モノラル; turn.audio_pcm16、np.int16 配列
your_stt 。送信（ターン。audio_base64）
クライアント。開始()
音声のみの展開の場合は、videoElement (ブラウザー) を省略するか、enable_video=False (Python) を渡します。
どちらの SDK も、 precision 、 vad 、 state 、interrupt 、 config 、および stats イベントを発行し、 mute() / unmute() 、 setThreshold() / set_threshold() 、および markResponding() / mark_responding() を公開します。 「packages/saa-js」 と 「packages/saa-py」 を参照してください。
実行可能なエンドツーエンドのデモは、examples/web/ (ブラウザー) および example/python/ (ターミナル) にあります。
LiveKit Agent の場合、saa-livekit-client は SAA をルームに導入し、分類子を実行し、「saa」データ トピックでイベントを公開します。エージェントは tentionEngine を通じてそれらを消費し、セッションをゲートします。
saa_livekit_client からインポート AttendeeEngine 、attention_agent_token 、start_attention_session
saa = await start_attention_session (
api_key = SAA_API_KEY 、 livekit_url = LIVEKIT_URL 、
エージェントトークン = アテンションエージェントトークン ( api_key = LK_KEY 、 api_secret = LK_SECRET 、 room_name = ctx . room . name )、
部屋名 = ctx 。部屋。名前 、参加者 ID = ユーザー 。アイデンティティ、
)
エンジン = AttendeeEngine (ctx . room 、agent_identity = saa . agent_identity )
@エンジン。 on_prediction
def _ ( p ):
セッション。入力を行ってください。 set_audio_enabled ( p . aligned_class == 2 ) # ゲート
エンジンを待ちます。開始()
OpenAI Realtime エージェントとバニラ JS Web クライアントの 2 つの実行可能なサンプルは、examples/livekit/ にあります。
Daily で実行されている Pipecat 音声エージェントの場合、saa-pipecat-client は SAA を Daily ルームに導入し、Daily のアプリメッセージ チャネルで「saa」トピックの下でイベントを公開します。ボットは AttendeeEngine ( DailyTransport 経由でサブスクライブ) を通じてそれらを消費し、パイプラインをゲートします。

e.
saa_pipecat_client からインポート AttendeeEngine 、attention_agent_token 、start_attention_session
saa = await start_attention_session (
api_key = SAA_API_KEY 、 room_url = ROOM_URL 、
エージェントトークン = アテンションエージェントトークン ( daily_api_key = DAILY_API_KEY 、 room_name = room_name )、
参加者アイデンティティ = 人間アイデンティティ 、
)
エンジン = tentionEngine (トランスポート、agent_identity = saa .agent_identity)
エンジン。バインドタスク (タスク)
@エンジン。 on_prediction
def _ ( p ):
addressee_gate 。抑制 = ( p . aligned_class == 1 および p . 信頼度 > 0.7 )
エンジンを待ちます。開始()
実行可能な Web クライアントのサンプルは、examples/pipecat/ にあります。
Celebrities Conversational AI は密閉された WebRTC ルームでエージェントを実行するため、SAA が直接それに参加することはできません。代わりに、ストリーミング SDK はフィード モードで実行されます。feed_audio を通じてエージェントのマイク音声を渡し、SAA の予測イベントでエージェントをゲートします。
saa インポート AttendantClient から
# フィード モード: SDK 自体は何もキャプチャしません。あなたがオーディオを提供します
saa = AttendantClient (トークン = SAA_API_KEY 、enable_audio = False、enable_video = False)
@さぁ。 on_prediction
def _ ( p ):
mic_to_agent 。 Enabled = ( p . aligned_class == 2 ) # 2 = デバイスにアドレス指定
さぁ。開始()
# イレブンラボの AudioInterface 入力コールバック:
さぁ。 feed_audio ( mic_pcm16 )
実行可能なサンプルは、examples/elevenlabs/ にあります。
Twilio Media Streams テレフォニー エージェントの場合、ストリーミング SDK は通話音声を介してフィード モードで実行されます。このアダプターは、Twilio の μ-law 8 kHz フレームを PCM16 16 kHz にトランスコードして SAA にフィードし、デバイス主導のターンのみをブリッジに転送するため、サイド トーク、保留音楽、およびエージェント自身の TTS エコーはゲートアウトされます。
saa インポート AttendantClient から
saa = AttendantClient (トークン = SAA_API_KEY 、enable_audio = False、enable_video = False)
@さぁ。オンターン準備完了
def_(ターン

）：
橋 。 on_speech (turn .audio_base64) # デバイス主導の通話音声のみが継続されます
さぁ。開始()
# Twilio メディア ハンドラー内で、μ-law -> PCM16 をデコードした後:
さぁ。 feed_audio ( pcm16_frames )
実行可能な Media Streams ブリッジ (コーデック、ペースアウトバウンド、自動 mark_responding ) は、examples/twilio/ にあります。
積極的なエージェント (最初に話します)
ストリーミング SDK は、markResponding(true) / mark_responding(True) を公開するため、エージェントが自分が話しているときにアサートし、自身の TTS 中にゲートを抑制し、テールがクリアされると再開できるようになります。 LiveKit ブリッジと Pipecat ブリッジは、engine.responding_start() / Reply_stop() を介して同じライフサイクルを公開し、同一の表面を提供します。
SAA は、VAD と STT の間でモデルに依存しない宛先を決定します。これは、VAD (誰かが話していますか)、話者ダイアライゼーション (どの声であるか)、方向転換検出 (話し終えたか)、ウェイク ワード (フレーズを言ったか) とは異なる質問に答えるため、これらのレイヤーで構成され、ウェイク ワードを完全に置き換えることができます。
オープン SDK は SAA クラウドにストリーミングされます。オーディオをデバイス上に保持する必要がある展開 (電話、組み込みシステム、ウェアラブル、ロボット工学、キオスク) の場合は、attentionlabs.ai でデバイス上および組み込みアクセスをリクエストしてください。
package/saa-js/README.md 、packages/saa-py/README.md 、ストリーミング SDK リファレンス。
package/saa-livekit-client/README.md : LiveKit クライアント。
package/saa-pipecat-client/README.md : Pipecat-on-Daily クライアント。
Examples/README.md 、実行可能なサンプル。
example/twilio/README.md : Twilio Media Streams ブリッジ。
リポジトリ全体の Apache-2.0、各パッケージおよびサンプルはその下に出荷されます (各サブツリーの LICENSE を参照)。ホスト型クラウド サービスには、attention labs の利用規約が適用されます。
SECURITY.md · CONTRIBUTING.md · CODE_OF_CONDUCT.md · CHANGELOG.md · NOTICE · CITATION.cff
で

テンションラボプロジェクト。 © 2026 Socero Inc.
音声エージェントの宛先層: エージェント宛ての音声のみが、どのスタックでも STT、LLM、または TTS に到達します。
Apache-2.0ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
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

The addressee layer for voice agents: only speech meant for your agent reaches your STT, LLM, or TTS, on any stack. - attenlabs/saa-sdk

Hey folks. We built SAA (Selective Auditory Attention) after trying to find ways to make a good experience with multiple robots/multiple agents. What typically ended up happening is they'd never stop talking. This is an SDK you can put before your STT. It lets you know when your device is being spoken to or not without a wakeword. You can use it for:
-Single AI, Multi human
-Multi AI, Single human
-Multi AI, Multi human (we recommend also adding a wakeword on top for a better system) There are two models. One that is video + audio and one that is just audio. The way it overall works is that it looks for shifts in attention patterns (body language changes, vocal patterns) to work. It's a tough problem to nail as every human being is different in how they interact with people/devices. Let me know how it is!

GitHub - attenlabs/saa-sdk: The addressee layer for voice agents: only speech meant for your agent reaches your STT, LLM, or TTS, on any stack. · GitHub
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
attenlabs
/
saa-sdk
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
221 Commits 221 Commits .github .github assets assets examples examples packages packages .gitignore .gitignore CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md llms-full.txt llms-full.txt llms.txt llms.txt View all files Repository files navigation
Tells your voice agent which speech is actually for it.
One decision per utterance: only addressee speech reaches your STT, LLM, and TTS. No wake word.
Drop-in for the voice-agent stack you already use:
A voice agent's microphone hears every voice in the room: yours, a coworker's, the kids, a podcast playing on the laptop, the agent's own TTS bleeding back through the speakers. Most pipelines respond to any of it, paying STT for every transcribed second and triggering the LLM on speech that was never directed at the device.
SAA (Selective Auditory Attention) is a hosted classifier that runs before STT and decides, per utterance, whether the speech was directed at the device. Side talk, background media, and the agent's own playback are filtered out, so your STT / LLM / TTS only see audio meant for the agent.
No wake word. SAA decides per-utterance from the audio (and optionally low-rate video) stream.
Hosted. A real-time WebSocket to attention labs' cloud; the open SDKs are thin clients. Because it gates before STT, only addressed speech reaches the STT, LLM, and TTS you already run, so your downstream services and logs see less audio, not more. On-device deployment is a separate enterprise licence.
The architecture and evaluation are described in the technical report .
Shape
Package
Use it when
Streaming SDK
@attenlabs/saa-js , attenlabs-saa
your app captures the audio/video itself and you want typed attention events to gate your own pipeline. Good for web agents, mall kiosks, drive-through agents, and robots.
LiveKit
saa-livekit-client
you run a LiveKit Agents voice agent. SAA joins your room and gates the session.
Pipecat (Daily)
saa-pipecat-client
you run a Pipecat voice agent on Daily. SAA joins your Daily room and gates the pipeline through the "saa" app-message topic.
ElevenLabs
attenlabs-saa
you run an ElevenLabs Conversational AI agent. SAA gates it via the streaming SDK's feed_audio (its room is sealed, so SAA can't join it directly).
Twilio
attenlabs-saa
you run a Twilio Media Streams telephony agent. SAA gates inbound/outbound call audio (μ-law 8 kHz resampled to PCM16) via the streaming SDK's feed_audio .
Install
npm install @attenlabs/saa-js # JavaScript / browser
pip install attenlabs-saa # Python (streaming SDK)
pip install saa-livekit-client # Python (LiveKit)
pip install saa-pipecat-client # Python (Pipecat on Daily)
Get an API key at attentionlabs.ai .
You capture the media; SAA emits typed events. The key event is turnReady / turn_ready , one device-directed utterance, captured and ready to forward to your STT or LLM.
import { AttentionClient } from "@attenlabs/saa-js" ;
const client = new AttentionClient ( { token : process . env . SAA_API_KEY } ) ;
// fires once per device-directed turn; turn.audioBase64 is PCM16 @ 16 kHz
client . on ( "turnReady" , ( turn ) => yourSTT . send ( turn . audioBase64 ) ) ;
await client . start ( { videoElement : document . querySelector ( "video" ) } ) ;
import os
from saa import AttentionClient
client = AttentionClient ( token = os . environ [ "SAA_API_KEY" ])
@ client . on_turn_ready
def _ ( turn ):
# turn.audio_base64, PCM16 @ 16 kHz mono; turn.audio_pcm16, np.int16 array
your_stt . send ( turn . audio_base64 )
client . start ()
For audio-only deployments, omit videoElement (browser) or pass enable_video=False (Python).
Both SDKs also emit prediction , vad , state , interrupt , config , and stats events, and expose mute() / unmute() , setThreshold() / set_threshold() , and markResponding() / mark_responding() . See packages/saa-js and packages/saa-py .
Runnable end-to-end demos are in examples/web/ (browser) and examples/python/ (terminal).
For LiveKit Agents , saa-livekit-client brings SAA into your room to run the classifier and publish events on the "saa" data topic. Your agent consumes them through AttentionEngine and gates the session.
from saa_livekit_client import AttentionEngine , attention_agent_token , start_attention_session
saa = await start_attention_session (
api_key = SAA_API_KEY , livekit_url = LIVEKIT_URL ,
agent_token = attention_agent_token ( api_key = LK_KEY , api_secret = LK_SECRET , room_name = ctx . room . name ),
room_name = ctx . room . name , participant_identity = user . identity ,
)
engine = AttentionEngine ( ctx . room , agent_identity = saa . agent_identity )
@ engine . on_prediction
def _ ( p ):
session . input . set_audio_enabled ( p . aligned_class == 2 ) # the gate
await engine . start ()
Two runnable samples, an OpenAI Realtime agent and a vanilla-JS web client, are in examples/livekit/ .
For Pipecat voice agents running on Daily, saa-pipecat-client brings SAA into your Daily room and publishes events on Daily's app-message channel under the "saa" topic. Your bot consumes them through AttentionEngine (which subscribes via your DailyTransport ) and gates the pipeline.
from saa_pipecat_client import AttentionEngine , attention_agent_token , start_attention_session
saa = await start_attention_session (
api_key = SAA_API_KEY , room_url = ROOM_URL ,
agent_token = attention_agent_token ( daily_api_key = DAILY_API_KEY , room_name = room_name ),
participant_identity = human_identity ,
)
engine = AttentionEngine ( transport , agent_identity = saa . agent_identity )
engine . bind_task ( task )
@ engine . on_prediction
def _ ( p ):
addressee_gate . suppressed = ( p . aligned_class == 1 and p . confidence > 0.7 )
await engine . start ()
A runnable web-client sample is in examples/pipecat/ .
ElevenLabs Conversational AI runs its agent in a sealed WebRTC room, so SAA can't join it directly. Instead the streaming SDK runs in feed mode: you hand it the agent's microphone audio through feed_audio and gate the agent on SAA's prediction events.
from saa import AttentionClient
# feed mode: the SDK captures nothing itself; you supply the audio
saa = AttentionClient ( token = SAA_API_KEY , enable_audio = False , enable_video = False )
@ saa . on_prediction
def _ ( p ):
mic_to_agent . enabled = ( p . aligned_class == 2 ) # 2 = addressed to the device
saa . start ()
# in ElevenLabs' AudioInterface input callback:
saa . feed_audio ( mic_pcm16 )
A runnable sample is in examples/elevenlabs/ .
For Twilio Media Streams telephony agents, the streaming SDK runs in feed mode over the call audio. The adapter transcodes Twilio's μ-law 8 kHz frames to PCM16 16 kHz, feeds them to SAA, and forwards only device-directed turns to your bridge, so side talk, hold music, and the agent's own TTS echo are gated out.
from saa import AttentionClient
saa = AttentionClient ( token = SAA_API_KEY , enable_audio = False , enable_video = False )
@ saa . on_turn_ready
def _ ( turn ):
bridge . on_speech ( turn . audio_base64 ) # only device-directed call audio continues
saa . start ()
# in the Twilio media handler, after decoding μ-law -> PCM16:
saa . feed_audio ( pcm16_frames )
A runnable Media Streams bridge (codec, paced outbound, automatic mark_responding ) is in examples/twilio/ .
Proactive agents (speak first)
The streaming SDKs expose markResponding(true) / mark_responding(True) so the agent can assert when it is the one speaking, suppressing the gate during its own TTS and resuming once the tail clears. The LiveKit and Pipecat bridges expose the same lifecycle via engine.responding_start() / responding_stop() , identical surface.
SAA is the model-agnostic addressee decision between your VAD and STT. It answers a different question than VAD (is anyone speaking), speaker diarization (which voice it is), turn detection (have they finished), or a wake word (did they say the phrase), so it composes with those layers and can replace the wake word outright.
The open SDKs stream to the SAA cloud. For deployments where audio must stay on the device (telephony, embedded systems, wearables, robotics, kiosks), request on-device and embedded access at attentionlabs.ai .
packages/saa-js/README.md , packages/saa-py/README.md , streaming SDK reference.
packages/saa-livekit-client/README.md : the LiveKit client.
packages/saa-pipecat-client/README.md : the Pipecat-on-Daily client.
examples/README.md , runnable examples.
examples/twilio/README.md : the Twilio Media Streams bridge.
Apache-2.0 across the repo, each package and the examples ship under it (see each subtree's LICENSE ). The hosted cloud service is governed by the attention labs Terms of Service.
SECURITY.md · CONTRIBUTING.md · CODE_OF_CONDUCT.md · CHANGELOG.md · NOTICE · CITATION.cff
An attention labs project. © 2026 Socero Inc.
The addressee layer for voice agents: only speech meant for your agent reaches your STT, LLM, or TTS, on any stack.
Apache-2.0 license
Code of conduct
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
