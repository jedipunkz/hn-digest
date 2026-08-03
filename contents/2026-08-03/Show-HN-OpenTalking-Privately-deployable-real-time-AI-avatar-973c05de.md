---
source: "https://github.com/datascale-ai/opentalking"
hn_url: "https://news.ycombinator.com/item?id=49150333"
title: "Show HN: OpenTalking – Privately deployable real-time AI avatar"
article_title: "GitHub - datascale-ai/opentalking: OpenTalking: An industrial-grade open-source AI digital human framework that supports real-time conversation, private deployment, and pluggable models. · GitHub"
author: "xx123122"
captured_at: "2026-08-03T01:52:12Z"
capture_tool: "hn-digest"
hn_id: 49150333
score: 1
comments: 0
posted_at: "2026-08-03T01:50:44Z"
tags:
  - hacker-news
  - translated
---

# Show HN: OpenTalking – Privately deployable real-time AI avatar

- HN: [49150333](https://news.ycombinator.com/item?id=49150333)
- Source: [github.com](https://github.com/datascale-ai/opentalking)
- Score: 1
- Comments: 0
- Posted: 2026-08-03T01:50:44Z

## Translation

タイトル: Show HN: OpenTalking – プライベートに展開可能なリアルタイム AI アバター
記事タイトル: GitHub - datascale-ai/opentalking: OpenTalking: リアルタイム会話、プライベート展開、プラグイン可能なモデルをサポートする産業グレードのオープンソース AI デジタル ヒューマン フレームワーク。 · GitHub
説明: OpenTalking: リアルタイムの会話、プライベート展開、プラグイン可能なモデルをサポートする産業グレードのオープンソース AI デジタル ヒューマン フレームワーク。 - データスケール-ai/オープントーキング

記事本文:
GitHub - datascale-ai/opentalking: OpenTalking: リアルタイムの会話、プライベート展開、プラグイン可能なモデルをサポートする産業グレードのオープンソース AI デジタル ヒューマン フレームワーク。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
ああ、ああ！
の

読み込み中にエラーが発生しました。このページをリロードしてください。
データスケール-ai
/
オープントーク
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
147 コミット 147 コミット .github .github apps apps configs configs docker docker docs docs 例 例 opentalking opentalking スクリプト スクリプト テスト テスト .env.example .env.example .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml AGENT.md AGENT.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス Makefile Makefile README.en.md README.en.md README.md README.md README.zh.md README.zh.md conftest.py conftest.py docker-compose.gpu.yml docker-compose.gpu.yml docker-compose.yml docker-compose.yml mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
オープンソースのリアルタイム デジタル ヒューマン パイプライン: LLM、TTS、WebRTC、キャラクター ボイス、プラグイン可能なモデル バックエンド
デモ ·
展開 ·
クイックスタート ·
モデル ·
ロードマップ ·
ドキュメントとコミュニティ
OpenTalking は、デジタル人間とリアルタイムの会話のためのオープンソース オーケストレーション フレームワークです。これは、デジタル ヒューマンの会話型製品のコア パスをカバーしています。つまり、フロントエンドの対話、セッション状態、LLM 応答、STT、TTS および音声選択、割り込み制御、字幕イベント、WebRTC オーディオ/ビデオ再生、ローカルまたはリモート モデル サービスへの呼び出しです。
OpenTalking は、実用的なデジタル ヒューマン制作スタックとして設計されています。 WebUI、アバターおよび音声アセット ライブラリ、ナレッジ ベース、メモリ、マルチセッション状態、LLM / STT / TTS プロバイダー、WebRTC 再生、モデル バックエンドは 1 つのプロジェクトに編成されています。軽量の Mock モードから開始し、ローカルの QuickTalk / Wav2 に接続できます。

Lip を使用するか、FlashTalk、FasterLivePortrait、その他の高品質またはより複雑なモデル ワークフローに OmniRT を使用します。
高速トライアル: モック/ドライバーレス モード。ビデオ モデルの重みをダウンロードする前に API、TTS、および WebRTC パスを検証するのに役立ちます。
リアルタイム会話: QuickTalk 、 Wav2Lip 、 FlashTalk 、およびその他のモデルを接続して、デジタルとヒューマンの対話型対話を実現します。
ビデオの作成とクローン作成: オーディオ/テキスト主導のビデオ作成とカメラ/アップロードされたビデオ主導のビデオ クローン ワークフローに FasterLivePortrait ランタイムを再利用します。
プライベート展開: ローカル STT/TTS、OpenAI 互換 LLM、ナレッジ ベース、メモリ、OmniRT リモート推論、Docker、分散展開をサポートします。
ドキュメント サイト: https://datascale-ai.github.io/opentalking/latest/en/
中国語ドキュメント: https://datascale-ai.github.io/opentalking/latest/
OpenTalking は、デジタル人間と人間の会話パイプラインを管理するための Web サービス インターフェイスを提供します。同じページ上で、アバターの選択または作成、音声、LLM、TTS、STT、デジタル ヒューマン ドライバー モデルの構成、モデルの接続ステータスの検査、リアルタイムの会話、字幕、オーディオ/ビデオ再生の検証を行うことができます。
これらのデモでは、リアルタイム会話、ビデオ作成、ビデオ クローンという 3 つの一般的なフロントエンド ワークフローを取り上げます。
ヘルスケア-zh.mp4
ライブコマース
ライブコマース.mp4
黄山観光ガイド
黄山ガイド.mp4
A. リアルタイムの会話
Eコマースのライブストリーム
eコマース.mp4
仲間キャラクター
コンパニオン.mp4
ニュースキャスター
ニュースキャスター.mp4
B. ビデオの作成
オーディオ主導型
オーディオドライブ.mp4
テキスト駆動
テキストドライブ.mp4
クローン化された音声駆動型
cloned_voice_drive.mp4
C. ビデオクローン
リアルタイムカメラ模倣
カメラドライブ.mp4
アップロードされた動画の模倣
ビデオドライブ.mp4
導入パスの選択
OpenTalking のオーケストレーション層 (API / ワーカー / フロントエンド) とデジタルヒューマン合成バックエンド

(mock 、 local 、 direct_ws 、または OmniRT ) は個別にデプロイできます。プロジェクトを初めて使用する場合は、モック モードから開始してフル パスを検証し、その後、GPU、モデル、およびプライベート デプロイの要件に基づいて実際のレンダリング モデルに切り替えます。
まず、2 つのクイックスタート パスのいずれかを選択します。
すべてを手動で設定する前に、OpenTalking + OmniRT + QuickTalk リアルタイム デジタル ヒューマン パスを試してみたい場合は、Compshare で公開されているコミュニティ イメージを使用してください。
ガイド: Compshare イメージのクイック エクスペリエンス
イメージには、OpenTalking、OmniRT、QuickTalk ランタイム環境、およびモデル ファイルが含まれています。インスタンスをデプロイした後、ポート 5173 を開き、プラットフォームによって提供されるインスタンス URL にアクセスします。サービスを手動で再起動する必要がある場合は、ガイドのコマンドに従ってください。
OpenTalking をソースから実行する場合は、このパスを使用します。ビデオ モデルの重みをまだダウンロードしたくない場合は、モック モードから始めてください。モック モードは組み込みの静的フレームを使用しますが、LLM 応答、ストリーミング TTS、字幕イベント、および WebRTC 配信は引き続き完全な製品パスを通じて実行されます。
git clone https://github.com/datascale-ai/opentalking.git
CD オープントーキング
uv sync --extra dev --python 3.11
ソース .venv/bin/activate
cp .env.example .env
.env を編集し、少なくとも LLM を構成します。デフォルトの TTS はキーレス エッジ ボイスを使用できます。 LLM、STT、および TTS は独立したプロバイダーです。 「構成」と「LLM/STT」を参照してください。
bash スクリプト/start_unified.sh --mock
デフォルトのフロントエンド URL は http://localhost:5173 です。ポートを指定するには:
bash スクリプト/start_unified.sh --mock --api-port 8210 --web-port 5280
サービスを停止します。
bash スクリプト/クイックスタート/stop_all.sh
実際のモデルのエントリポイント
モック モードが機能したら、マシンに基づいて実際のモデル パスを選択します。重みのダウンロード、ディレクトリ レイアウト、ミラー、チェック、トラブルシューティングが維持されます

ドキュメントにあります。 README には起動エントリポイントのみが保持されています。
# ローカル QuickTalk: コンシューマ-GPU 単一マシン パス
import DIGITAL_HUMAN_HOME= " ${DIGITAL_HUMAN_HOME :- $HOME / デジタルヒューマン} "
import OPENTALKING_MODEL_ROOT= " ${OPENTALKING_MODEL_ROOT :- $DIGITAL_HUMAN_HOME / モデル} "
エクスポート OPENTALKING_TORCH_DEVICE=cuda:0
import OPENTALKING_QUICKTALK_ASSET_ROOT= " $OPENTALKING_MODEL_ROOT /quicktalk "
エクスポート OPENTALKING_QUICKTALK_WORKER_CACHE=1
bash scripts/start_unified.sh --backend local --model Quicktalk --api-port 8210 --web-port 5280
# リモート OmniRT / FlashTalk: 高品質またはマルチカード パス
bash スクリプト/start_unified.sh \
--バックエンドオムニアート \
--モデル フラッシュトーク \
--API ポート 8210 \
--web ポート 5280 \
--omnirt http:// < gpu サーバー > :9000
追加のエントリポイント:
Docker Compose と本番デプロイメント
モデル
入力
推奨されるバックエンド
リソースのガイダンス
モック
参考画像・静止フレーム
モック
GPUは必要ありません
早口トーク
テンプレートビデオ + オーディオ
地元の
CUDA GPU、RTX 3090 / 4090 推奨
wav2lip
参考画像/フレーム+音声
ローカル / オムニアート
>= 8 GB GPU / NPU メモリ
ミューズトーク
フルフレーム + オーディオ
オムニアート / ローカル
>= 12 GB GPU メモリ
soulx-フラッシュトーク-14b
ポートレート + 音声
オムニアート
マルチGPU/NPU
soulx-フラッシュヘッド-1.3b
ポートレート + 音声
オムニアート
マルチGPU/NPU
より速いライブポートレート
ポートレート/走行ビデオ/オーディオ
オムニアート
シングル GPU のリアルタイム ポートレート ペーストバック、ビデオ作成、ビデオ クローン
コンシューマ向け GPU リファレンス
モデル
ハードウェア
入力
出力
VRAM
スループット
早口トーク
RTX3090
テンプレートビデオ + オーディオ
720x900 / 25fps
約3.8GiB
約35fps
重みのダウンロード、Docker、トラブルシューティング、およびモデルの構成については、「モデルのデプロイメント」を参照してください。
Atlas Cloud は、オールモーダル AI 推論プラットフォームです。 1 つの API でビデオ生成、画像生成、LLM にアクセスできるため、複数の API を統合する必要はありません。

別途承認します。 1 つの統合で、300 以上の厳選されたオールモーダル モデルにルーティングできます。
OpenTalking は、LLM に OpenAI 互換インターフェイスを使用します。 Atlas がホストする DeepSeek / Qwen モデルを使用するには、OPENTALKING_LLM_BASE_URL を https://api.atlascloud.ai/v1 に指定します。 LLM と STT を参照してください。予算に優しい API オプションについては、Atlas Cloud のコーディング プランを参照してください。
より自然なリアルタイム会話
割り込み処理、低遅延応答、オーディオ/ビデオ同期、長時間セッションの回復、実行時の可視性が向上します。
コンシューマ-GPU マルチモデル パス
QuickTalk / Wav2Lip / MuseTalk ローカル パスのアセット チェック、プリウォーム、キャッシュの再利用、低メモリ パラメーターなどの RTX 3090 / 4090 / WSL2 ベンチマークを改善し、より多くの FasterLivePortrait ビデオ作成とビデオ クローン測定を埋めます。
ワンコマンドによる Windows / WSL2 導入
現在の Windows ドキュメントとテスト記録に基づいて、モデルのダウンロード、ランタイム インストール、環境チェック、診断のハードルを下げ続けます。
高品質のプライベート展開
外部 OmniRT 推論サービス、マルチモデル エンドポイント、キャパシティ スケジューリング、ヘルス チェック、運用監視、GPU / NPU 導入ガイダンスを改善します。
より多くのクラウド音声およびマルチモーダルプロバイダー
現在の OpenAI 互換、DashScope、および Xiaomi MiMo プロファイルに加えて、プラグ可能な STT / TTS / LLM プロバイダー、統合されたフロントエンド選択、およびプロバイダー レベルのヘルス チェックを拡張します。
エージェント、メモリ、プラットフォームの機能
アセット ライブラリ、ナレッジ ベース、メモリ、マルチセッション スケジューリング、ツール呼び出し、OpenClaw / 外部エージェントの統合を製品化し、可観測性、安全性、ライセンス付き音声、合成コンテンツのラベル付けを埋め込みます。
2026-06-25: WeChat メモリのインポートとペルソナのワークフロー
WeChat メモリ ペルソナのインポート、ドキュメント、および関連するペルソナ ワークフローを追加しました。フロントエンドはダメだ

nger はペルソナの選択と運転モデル​​の選択を相互に排他的なものとして扱うため、ユーザーはインポートされたメモリ/ペルソナのコンテキストを選択したアバター ドライバーと組み合わせることができます。
2026-06-23: ローカル CosyVoice TRT サイドカーのデプロイメント
TensorRT / FP16 アクセラレーション ノート、ランタイム チューニング、専用環境の分離、起動チェック、およびローカル TTS と QuickTalk をペアリングするための測定済みデプロイメント ガイダンスを含むローカル CosyVoice サイドカー デプロイメント パスを追加しました。
2026-06-22: ランタイム構成、メモリリフレッシュ、イマーシブシーン
ランタイム API 構成ページを追加し、ランタイム更新中の mem0 プロバイダーのリリースを改善し、シーン アセット パイプラインを拡張しました。シーン アセット API、アセット ライブラリの統合、イマーシブ会話モード、シーン/アバターのアンカーリング、透明な背景の処理、ビューの切り替え全体でのリアルタイム メディアの保存です。
2026-06-18/19: クイックスタート分割、LightRAG ランタイム構成、およびシナリオ ガイド
クイックスタートを Compshare イメージと自己展開パスに分割し、LightRAG ランタイム構成とクイックスタート更新を追加し、mem0 / Hugging Face ダウンロード ツールの依存関係に関するメモを修正し、Huangshan デジタル ヒューマン ガイドを追加しました。
2026-06-12: QuickTalk ローカル アセットの修正と Apple Silicon のサポート

[切り捨てられた]

## Original Extract

OpenTalking: An industrial-grade open-source AI digital human framework that supports real-time conversation, private deployment, and pluggable models. - datascale-ai/opentalking

GitHub - datascale-ai/opentalking: OpenTalking: An industrial-grade open-source AI digital human framework that supports real-time conversation, private deployment, and pluggable models. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Uh oh!
There was an error while loading. Please reload this page .
datascale-ai
/
opentalking
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
147 Commits 147 Commits .github .github apps apps configs configs docker docker docs docs examples examples opentalking opentalking scripts scripts tests tests .env.example .env.example .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml AGENT.md AGENT.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.en.md README.en.md README.md README.md README.zh.md README.zh.md conftest.py conftest.py docker-compose.gpu.yml docker-compose.gpu.yml docker-compose.yml docker-compose.yml mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml View all files Repository files navigation
Open-source real-time digital-human pipeline: LLM, TTS, WebRTC, character voices, and pluggable model backends
Demos ·
Deployment ·
Quickstart ·
Models ·
Roadmap ·
Docs & Community
OpenTalking is an open-source orchestration framework for real-time digital-human conversations. It covers the core path of a digital-human conversational product : frontend interaction, session state, LLM replies, STT, TTS and voice selection, interruption control, subtitle events, WebRTC audio/video playback, and calls into local or remote model services.
OpenTalking is designed as a practical digital-human production stack. The WebUI, avatar and voice asset libraries, knowledge bases, memory, multi-session state, LLM / STT / TTS providers, WebRTC playback, and model backends are organized in one project. You can start with the lightweight Mock mode, connect local QuickTalk / Wav2Lip, or use OmniRT for FlashTalk, FasterLivePortrait, and other higher-quality or more complex model workflows.
Fast trial : mock / driverless mode , useful for validating the API, TTS, and WebRTC path before downloading video model weights.
Real-time conversation : connect QuickTalk , Wav2Lip , FlashTalk , and other models for interactive digital-human dialogue.
Video creation and cloning : reuse FasterLivePortrait runtime for audio/text-driven video creation and camera/uploaded-video-driven video clone workflows.
Private deployment : supports local STT/TTS, OpenAI-compatible LLMs, knowledge bases, memory, OmniRT remote inference, Docker, and distributed deployment.
Documentation site: https://datascale-ai.github.io/opentalking/latest/en/
Chinese docs: https://datascale-ai.github.io/opentalking/latest/
OpenTalking provides a Web service interface for managing the digital-human conversation pipeline. You can select or create avatars, configure voices, LLM, TTS, STT, and digital-human driver models, inspect model connection status, and validate real-time conversation, subtitles, and audio/video playback on the same page.
These demos cover three common frontend workflows: real-time conversation, video creation, and video clone.
healthcare-zh.mp4
Live commerce
live-commerce.mp4
Huangshan tourism guide
huangshan-guide.mp4
A. Real-time Conversation
E-commerce livestream
eCommerce.mp4
Companion character
companion.mp4
News anchor
newscaster.mp4
B. Video Creation
Audio driven
audio_drive.mp4
Text driven
text_drive.mp4
Cloned voice driven
cloned_voice_drive.mp4
C. Video Clone
Realtime camera imitation
camera_drive.mp4
Uploaded video imitation
video_drive.mp4
Choose A Deployment Path
OpenTalking's orchestration layer (API / Worker / frontend) and digital-human synthesis backend ( mock , local , direct_ws , or OmniRT ) can be deployed independently. If you are new to the project, start with Mock mode to validate the full path, then switch to a real rendering model based on your GPU, model, and private-deployment requirements.
Choose one of the two quickstart paths first:
If you want to try the OpenTalking + OmniRT + QuickTalk real-time digital-human path before setting up everything manually, use the community image we published on Compshare:
Guide: Compshare image quick experience
The image includes OpenTalking, OmniRT, the QuickTalk runtime environment, and model files. After deploying an instance, open port 5173 and visit the instance URL provided by the platform. If you need to restart services manually, follow the commands in the guide.
Use this path when you want to run OpenTalking from source. Start with Mock mode if you do not want to download video model weights yet: Mock mode uses the built-in static frame, while LLM replies, streaming TTS, subtitle events, and WebRTC delivery still run through the full product path.
git clone https://github.com/datascale-ai/opentalking.git
cd opentalking
uv sync --extra dev --python 3.11
source .venv/bin/activate
cp .env.example .env
Edit .env and configure at least an LLM. The default TTS can use the keyless edge voice. LLM, STT, and TTS are independent providers; see Configuration and LLM / STT .
bash scripts/start_unified.sh --mock
The default frontend URL is http://localhost:5173 . To specify ports:
bash scripts/start_unified.sh --mock --api-port 8210 --web-port 5280
Stop services:
bash scripts/quickstart/stop_all.sh
Real Model Entrypoints
After Mock mode works, choose a real model path based on your machine. Weight downloads, directory layout, mirrors, checks, and troubleshooting are maintained in the docs; the README keeps only the startup entrypoints:
# Local QuickTalk: consumer-GPU single-machine path
export DIGITAL_HUMAN_HOME= " ${DIGITAL_HUMAN_HOME :- $HOME / digital-human} "
export OPENTALKING_MODEL_ROOT= " ${OPENTALKING_MODEL_ROOT :- $DIGITAL_HUMAN_HOME / models} "
export OPENTALKING_TORCH_DEVICE=cuda:0
export OPENTALKING_QUICKTALK_ASSET_ROOT= " $OPENTALKING_MODEL_ROOT /quicktalk "
export OPENTALKING_QUICKTALK_WORKER_CACHE=1
bash scripts/start_unified.sh --backend local --model quicktalk --api-port 8210 --web-port 5280
# Remote OmniRT / FlashTalk: high-quality or multi-card path
bash scripts/start_unified.sh \
--backend omnirt \
--model flashtalk \
--api-port 8210 \
--web-port 5280 \
--omnirt http:// < gpu-server > :9000
More entrypoints:
Docker Compose and production deployment
Model
Input
Recommended backend
Resource guidance
mock
Reference image / static frame
mock
No GPU required
quicktalk
Template video + audio
local
CUDA GPU, RTX 3090 / 4090 recommended
wav2lip
Reference image / frames + audio
local / omnirt
>= 8 GB GPU / NPU memory
musetalk
Full frames + audio
omnirt / local
>= 12 GB GPU memory
soulx-flashtalk-14b
Portrait + audio
omnirt
Multi-GPU / NPU
soulx-flashhead-1.3b
Portrait + audio
omnirt
Multi-GPU / NPU
fasterliveportrait
Portrait / driving video / audio
omnirt
Single-GPU real-time portrait paste-back, video creation, video clone
Consumer-GPU Reference
Model
Hardware
Input
Output
VRAM
Throughput
quicktalk
RTX 3090
Template video + audio
720x900 / 25fps
About 3.8 GiB
About 35 fps
For weight downloads, Docker, troubleshooting, and model configuration, see Model deployment .
Atlas Cloud is an all-modal AI inference platform. One API gives you access to video generation, image generation, and LLMs, so you do not need to integrate multiple vendors separately. A single integration can route to 300+ curated all-modal models.
OpenTalking uses an OpenAI-compatible interface for LLMs. Point OPENTALKING_LLM_BASE_URL to https://api.atlascloud.ai/v1 to use Atlas-hosted DeepSeek / Qwen models. See LLM and STT . For budget-friendly API options, see Atlas Cloud's coding plan .
More natural real-time conversations
Improve interruption handling, low-latency response, audio/video sync, long-session recovery, and runtime visibility.
Consumer-GPU multi-model path
Improve asset checks, prewarm, cache reuse, low-memory parameters, and more RTX 3090 / 4090 / WSL2 benchmarks for QuickTalk / Wav2Lip / MuseTalk local paths, while filling in more FasterLivePortrait video creation and video clone measurements.
One-command Windows / WSL2 deployment
Continue lowering the barrier for model downloads, runtime installation, environment checks, and diagnostics based on the current Windows docs and test records.
High-quality private deployment
Improve external OmniRT inference services, multi-model endpoints, capacity scheduling, health checks, production monitoring, and GPU / NPU deployment guidance.
More cloud voice and multimodal providers
Extend pluggable STT / TTS / LLM providers, unified frontend selection, and provider-level health checks on top of the current OpenAI-compatible, DashScope, and Xiaomi MiMo profiles.
Agent, memory, and platform capabilities
Productize the asset library, knowledge bases, memory, multi-session scheduling, tool calling, and OpenClaw / external Agent integrations, then fill in observability, safety, licensed voices, and synthetic-content labeling.
2026-06-25: WeChat memory import and persona workflow
Added WeChat memory persona import, documentation, and the related persona workflow. The frontend no longer treats persona selection and driving-model selection as mutually exclusive, so users can combine imported memory/persona context with the selected avatar driver.
2026-06-23: Local CosyVoice TRT sidecar deployment
Added the local CosyVoice sidecar deployment path with TensorRT / FP16 acceleration notes, runtime tuning, dedicated environment isolation, startup checks, and measured deployment guidance for pairing local TTS with QuickTalk.
2026-06-22: Runtime configuration, memory refresh, and immersive scenes
Added the runtime API configuration page, improved mem0 provider release during runtime refresh, and expanded the scene asset pipeline: scene asset APIs, asset-library integration, immersive conversation mode, scene/avatar anchoring, transparent background handling, and realtime media preservation across view switches.
2026-06-18/19: Quickstart split, LightRAG runtime config, and scenario guides
Split the quickstart into Compshare image and self-deployment paths, added LightRAG runtime configuration and quickstart updates, fixed dependency notes for mem0 / Hugging Face download tooling, and added the Huangshan digital-human guide.
2026-06-12: QuickTalk local asset fixes and Apple Silicon suppor

[truncated]
