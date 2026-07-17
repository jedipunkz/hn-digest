---
source: "https://github.com/sushilk1991/velora"
hn_url: "https://news.ycombinator.com/item?id=48944208"
title: "Show HN: Velora – On-device macOS dictation (Whisper and a local LLM, no cloud)"
article_title: "GitHub - sushilk1991/velora: Local-first, on-device dictation for macOS (Apple Silicon) — Superwhisper/Wispr Flow alternative. Multilingual STT + LLM cleanup via MLX, fully offline. · GitHub"
author: "sushilk1991"
captured_at: "2026-07-17T07:10:59Z"
capture_tool: "hn-digest"
hn_id: 48944208
score: 1
comments: 0
posted_at: "2026-07-17T07:07:38Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Velora – On-device macOS dictation (Whisper and a local LLM, no cloud)

- HN: [48944208](https://news.ycombinator.com/item?id=48944208)
- Source: [github.com](https://github.com/sushilk1991/velora)
- Score: 1
- Comments: 0
- Posted: 2026-07-17T07:07:38Z

## Translation

タイトル: Show HN: Velora – オンデバイス macOS ディクテーション (Whisper およびローカル LLM、クラウドなし)
記事のタイトル: GitHub -ushilk1991/velora: macOS (Apple Silicon) 向けローカルファーストのオンデバイスディクテーション — Superwisper/Wispr Flow の代替手段。 MLX を介した多言語 STT + LLM クリーンアップ、完全オフライン。 · GitHub
説明: macOS (Apple Silicon) 向けのローカルファーストのオンデバイスディクテーション — Superwhisper/Wispr Flow の代替品。 MLX を介した多言語 STT + LLM クリーンアップ、完全オフライン。 -ushilk1991/ヴェローラ

記事本文:
GitHub -ushilk1991/velora: macOS (Apple Silicon) 向けローカルファーストのオンデバイスディクテーション — Superwhisper/Wispr Flow の代替手段。 MLX を介した多言語 STT + LLM クリーンアップ、完全オフライン。 · GitHub
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
シュシク1991
/
ヴェローラ
公共
通知
あなたは

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
122 コミット 122 コミット .github .github リソース リソース ソース/ Velora ソース/ Velora ドキュメント ドキュメント エンジン エンジン スクリプト スクリプト .gitignore .gitignore CLAUDE.md CLAUDE.md CONTRIBUTING.md COTRIBUTING.md ライセンス ライセンス Makefile Makefile Package.swift Package.swift README.md README.md バージョン バージョン すべてのファイルを表示 リポジトリ ファイルのナビゲーション
macOS のローカルファーストディクテーション。キーを押したまま、話し、放すと、使用しているどのアプリでも洗練されたテキストが表示されます。音声からテキストへの変換から AI クリーンアップまで、ディクテーションのすべてのステップは、Apple Silicon 上の MLX を介してデバイス上で実行されます。オーディオ、トランスクリプト、履歴が Mac から離れることはありません。個人辞書の用語は、iCloud Drive を通じて非公開で同期できることが確認されました。
Superwisper や Wispr Flow などのディクテーション ツールは、どこにいても目に見えず、高速でスマートな音声入力ができることを証明しました。ただし、Wispr Flow はクラウド内で文字起こしを行い、どちらもクローズドソースのサブスクリプションです。 Velora はオープンソースの答えです。
トグルによるものではなく、アーキテクチャによるプライベート。ディクテーション時のネットワーク呼び出しはゼロです。モデルは Hugging Face から 1 回ダウンロードされます。その後、エンジンがネットワークに接続されることはありません。アカウントもテレメトリーも分析もありません。
自由でオープン。 MITライセンス取得済み。モード、プロンプト、ボキャブラリーは、あなたが所有するプレーンな JSON ファイルです。
信頼できるほど速い。会話中に音声からテキストへのストリームが送信されるため、キーを放した直後にテキストが表示されます (AI クリーンアップが含まれます)。
あなたを書き換えません。クリーンアップは意図的に保守的です。文字起こしは応答せず、コンテンツは追加せず、LLM がドリフトした場合は生の言葉に戻る発散ガードを備えています。生のトランスクリプトは常に履歴に保存されます。
押し続けて話すとディクテーションを切り替える — 右を押し続ける

t オプション (デフォルト) を録音し、放すと挿入します。またはメニューバーのアイコンをダブルタップ/クリックして切り替えます。 Esc は常にきれいにキャンセルされます。ホットキーの動作は構成可能です。
Capsule HUD — マイクからの 24 バーのライブ波形を備えた小さな浮遊カプセル。決してフォーカスを奪うことはなく、リスニング→転写→挿入/エラー状態と変化します。開始/停止音が含まれています (切り替え可能)。
多言語 — デフォルトの Whisper-large-v3-turbo モデルは、自動言語検出により英語、インド英語、ヒンディー語、その他数十の言語を処理します。非ラテン語転写物 (デーヴァナーガリー文字、CJK、アラビア語) は英語に調整されたクリーンアップをスキップして、忠実さを保ちます。
オプションのローマ字出力 — ディクテーション設定の切り替えにより、非ラテン語音声をラテン文字に音訳します (ヒンディー語→自然なヒングリッシュ、例: "नमस्ते आज मौसम" → "Namaste aaj mausam")、英語の単語は英語のままです。デフォルトではオフです。
アプリ対応のスマート フォーマット — Velora は最前面のアプリを検出し、モードを自動的に選択します。
Slack / Messages / Discord / Telegram / WhatsApp → 簡潔、カジュアル、単一の短い文の末尾にピリオドがない
メール → プロフェッショナルな構造と口調
Notes / Obsidian / Notion / Bear → マークダウン可、列挙する場合はリスト
VS Code / Cursor / Terminal / iTerm / Ghostty / Warp / Zed → raw モード、AI 書き換えなし
その他すべて → クリーンで区切られたデフォルト
オンデバイス STT + LLM — Whisper-large-v3-turbo (多言語、英語のストリーミングに利用可能な parakeet-tdt-0.6b-v2、さらにヒンディー語/ヒングリッシュのスペシャリスト) による文字起こしと Qwen3-4B-Instruct-2507-4bit によるクリーンアップ/フォーマット。クリーンアップでは、フィラーを削除し、自己修正 (「いや、待ってください、火曜日のことです」) を適用し、句読点を付け、発話された「改行」/「新しい段落」を尊重します。
履歴ブラウザ — すべてのディクテーション (生のテキストと最終的なテキスト、アプリ、月)

de、duration) はローカル SQLite データベースに保存されます。 「履歴」タブから参照、検索、コピー、または貼り付けを再度行います。メニューバーのメニューには、最後の 3 つが表示されます (クリックしてコピーします)。
音声インテリジェンス — オンデバイスのダッシュボードは、ローカル履歴を有用な傾向に変えます。時間の経過に伴う単語とディクテーション、推定節約時間、連続記録、アプリとモードの内訳、STT/クリーンアップの遅延、クリーンアップ率、および観測範囲を含む正直なゼロ編集率です。共有カードには集計番号のみが含まれており、トランスクリプト、アプリ、連絡先テキストは含まれません。
プライベート ミーティング メモリ — Velora は、Zoom、Google Meet、Teams、または Slack Huddle がアクティブなときに、オプションのカレンダー マッチングを使用して録画を提案できます。キャプチャは明示的な確認後にのみ開始され、マイク (「私」) とコンピューターの音声 (「彼ら」) を別個のローカル トラックとして保持し、検索可能なトランスクリプト、概要、決定事項、アクション アイテムをバックグラウンドで生成します。処理は再開可能で、ライブディクテーションが優先されます。
話者ダイアライゼーション — 通話の相手側に複数人がいる場合、トランスクリプトではデバイス上のダイアライゼーション (sherpa-onnx: pyannote セグメンテーション + titanet 埋め込み、最初の会議で約 46 MB がダウンロード、sha256 固定) を使用して、発言者 1 / 話者 2 / … に分割されます。 1 対 1 の通話は単純に「彼ら」のままです。ダイアライゼーションに失敗すると、正常に元に戻ります。 「設定」→「ミーティング」に切り替えます。
安全な音声編集 — 任意のアプリでテキストを選択し、⌥⇧E (設定可能) を押して、編集内容を読み上げます: 「これをより形式的にする」、「文法を修正する」、「これを箇条書きに変える」。選択範囲のみがタッチされ、結果によってその位置が置き換えられ、 ⌘Z で元に戻ります。編集プロンプトはベンチマークされており (50 個のコマンド スイート、spikes/engine/bench_voice_edit.py で 94%)、保護されています。結果が使用できない場合でもテキストは変更されず、編集されたテキストは常にクリップに保存されます。

バックアップとしてのオード。
「聞いたとおり」の脱出ハッチ — クリーンアップで問題が発生した場合は、メニューバーから未加工の生のトランスクリプトを貼り付けるか (「次のように再フォーマット」→「聞いたとおり」)、履歴で表示します。再処理は不要で、オーディオ クリップが削除された後でも機能します。
ローカル CLI + MCP — スクリプトとローカル エージェントは、許可リストに登録された履歴/統計を検査したり、音声ファイルを文字起こししたり、目に見えて承認された 1 つのライブ ディクテーションを要求したりできます。アクセスはデフォルトでオフになっており、ネットワーク サーバーではなく所有者専用の Unix ソケットを通じて実行され、生の音声、画面コンテキスト、連絡先、学習データが公開されることはありません。
個人辞書 — Velora に正確な名前、製品用語、およびオプションの「聞き方→書き方」の修正を教えます。編集で学習した単語と自動学習した単語は表示されたままで元に戻すことができ、iCloud が利用可能な場合は、確認済みの辞書エントリのみがアプリ固有の iCloud Drive フォルダーを介して同期されます。
オーディオ アーカイブ + 再処理 — クリップは ~/.velora/audio にコンパクト FLAC として保存されるため (保持期間は設定可能、デフォルトは 6 か月 / 4 GB 上限)、履歴タブから直接、より良いモデルで過去のディクテーションを再文字起こしできます。
カスタム モード エディター - すべてのモードは ~/.velora/modes/ にある JSON ファイルで、[モード] タブから編集できます: モードごとの LLM プロンプト (Superwisper スタイル機能)、書式設定レベル、アプリ バインディング、ボキャブラリー、置換。ファイルをドロップして独自のファイルを作成します (「カスタマイズ」を参照)。
モデル ピッカー — 管理されたダウンロードを使用して、エンジンのレジストリから設定で STT モデルを選択します。
ライブスペクトル波形 — HUD の 24 バー波形はマイクの実際の FFT によって駆動されるため、バーはラウドネスとピッチの両方に反応します。
安全な挿入 — クリップボードはスナップショットを作成され、合成された ⌘V を中心に復元され、貼り付けをブロックするアプリに対するキーストローク入力のフォールバックが行われます。安全な入力フィールド（パスワード）

が検出され、挿入が抑制されます。
Apple Silicon Mac (M1以降)
uv (Python エンジンを管理)
Swift ツールチェーン — コマンドライン ツールで十分です ( xcode-select --install )。 Xcodeは必要ありません
デフォルト モデルの場合、最大 4.4 GB のディスク (初回実行時にダウンロード)
# 1. 前提条件 (1 回限り)
xcode-select --install # Swift ツールチェーン (コマンド ライン ツール)
カール -LsSf https://astral.sh/uv/install.sh | sh # uv、Python エンジン用
#2. ビルドして実行する
git clone https://github.com/sushilk1991/velora
CD ヴェローラ
make app # builds build/Velora.app (SwiftPM リリース + ハンドロール開発バンドル)
build/Velora.app を開く
make app は Swift アプリをコンパイルし、Python エンジンをバンドルします。エンジンの依存関係は、最初の起動時に uv によって取得されます。最初の起動では、マイクの許可、アクセシビリティの許可 (許可するとライブで検出される)、ホットキーの選択、および試してみるプレイグラウンドなどのオンボーディングが説明されます。最後に実際のディクテーションを行います。エンジンは、最初の実行時に Hugging Face からデフォルトのモデルをダウンロードします (最大 6 GB。早期にディクテーションを試行すると、オンボーディング ウィンドウ、メニューバー メニュー、および HUD にライブの進行状況が表示されます。最初に音声認識のロックが解除され、数分後に AI がクリーンアップされます)。その後はすべてオフラインになります。
構築したくないですか? Homebrew でインストールします。
brew install --caskushilk1991/tap/velora
または、「リリース」から Velora-x.y.z.dmg を取得し、「アプリケーション」に Velora をドラッグして開きます。 v0.4.3 以降のリリースは、Apple によって開発者 ID が署名され、公証されているため、Gatekeeper を介して通常に開きます。古い v0.4.1 イメージは公証以前のものであるため、バイパスするのではなく置き換える必要があります。
ベアバイナリではなく、.app バンドルを実行します。macOS 権限付与 (マイク、アクセシビリティ) が署名されたバンドル ID に付加されます。
テスト: make test は Python エンジン スイートを実行します。アプリには埋め込み機能もあります

Swift セルフテストを追加し ( swift run Velora --selftest )、100,000 行の履歴に対してパフォーマンス テスト チェックを行います。
インストールされたアプリには、次の場所に CLI が含まれています。
/アプリケーション/Velora.app/コンテンツ/リソース/bin/velora --help
[設定] → [全般] で [ローカル CLI とエージェントを許可する] を有効にし、 status 、 Recent 、 search 、 stats 、 transcribe 、または listen を使用します。機械可読出力のために --json を追加します。 listen では、マイクが開始される前に常に承認プロンプトが表示されます。
Velora は、ローカル MCP stdio サーバーと同じ狭い表面も公開します。
/アプリケーション/Velora.app/コンテンツ/リソース/bin/velora mcp
アプリが実行されている必要があります。ネットワーク上では何もリッスンされず、設定を無効にすると、履歴、統計、およびアクションへのアクセスがすぐに削除されます。ステータスは利用可能なままであるため、何が欠落しているかをツールで説明できます。
クリーンアップ モデルは、初回起動時に RAM 層によって選択され (14 GB 以下の Mac では Qwen3-1.7B-8bit、24 GB までの Qwen3-4B-4bit、それ以上の Qwen3.5-4B-8bit)、[設定] → [モデル] で変更できます。 Apple Silicon Mシリーズで測定:
デフォルトの多言語モデルはバッチ (ストリーミングではなくリリース時に転写) であり、正確なヒンディー語/インド英語に対する意図的な品質のトレードオフです。ライブ HUD 部分を使用して英語をストリーミングするには、[設定] で parakeet-tdt-0.6b-v2 を選択します。クリーンアップによって予算が使い果たされない場合、Velora は生のトランスクリプトをすぐに挿入します。

[切り捨てられた]

## Original Extract

Local-first, on-device dictation for macOS (Apple Silicon) — Superwhisper/Wispr Flow alternative. Multilingual STT + LLM cleanup via MLX, fully offline. - sushilk1991/velora

GitHub - sushilk1991/velora: Local-first, on-device dictation for macOS (Apple Silicon) — Superwhisper/Wispr Flow alternative. Multilingual STT + LLM cleanup via MLX, fully offline. · GitHub
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
sushilk1991
/
velora
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
122 Commits 122 Commits .github .github Resources Resources Sources/ Velora Sources/ Velora docs docs engine engine scripts scripts .gitignore .gitignore CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile Package.swift Package.swift README.md README.md VERSION VERSION View all files Repository files navigation
Local-first dictation for macOS. Hold a key, speak, release — polished text appears in whatever app you're using. Every dictation step, from speech-to-text to AI cleanup, runs on-device via MLX on Apple Silicon. Audio, transcripts, and history never leave your Mac; confirmed Personal Dictionary terms can sync privately through your iCloud Drive.
Dictation tools like Superwhisper and Wispr Flow proved the product: invisible, fast, smart voice input everywhere. But Wispr Flow transcribes in the cloud, and both are closed-source subscriptions. Velora is the open-source answer:
Private by architecture, not by toggle. Zero network calls at dictation time. Models are downloaded once from Hugging Face; after that the engine never touches the network. No accounts, no telemetry, no analytics.
Free and open. MIT-licensed. Modes, prompts, and vocabulary are plain JSON files you own.
Fast enough to trust. Speech-to-text streams while you talk , so text lands moments after you release the key — AI cleanup included.
Doesn't rewrite you. Cleanup is deliberately conservative: transcribe-don't-answer, no added content, and a divergence guard that falls back to your raw words if the LLM drifts. The raw transcript is always kept in history.
Hold-to-talk and toggle dictation — hold Right-Option (default) to record, release to insert; or double-tap / click the menubar icon to toggle. Esc always cancels cleanly. Hotkey behavior is configurable.
Capsule HUD — a small floating capsule with a live 24-bar waveform from your mic. It never steals focus, and morphs through listening → transcribing → inserted/error states. Start/stop sounds included (toggleable).
Multilingual — the default whisper-large-v3-turbo model handles English, Indian English, Hindi, and dozens more languages with automatic language detection. Non-Latin transcripts (Devanagari, CJK, Arabic) skip the English-tuned cleanup so they stay faithful.
Optional romanized output — a Dictation-settings toggle transliterates non-Latin speech into Latin letters (Hindi → natural Hinglish, e.g. "नमस्ते आज मौसम" → "Namaste aaj mausam"), keeping English words English. Off by default.
App-aware smart formatting — Velora detects the frontmost app and picks a mode automatically:
Slack / Messages / Discord / Telegram / WhatsApp → terse, casual, no trailing period on a single short sentence
Mail → professional structure and tone
Notes / Obsidian / Notion / Bear → markdown allowed, lists when you enumerate
VS Code / Cursor / Terminal / iTerm / Ghostty / Warp / Zed → raw mode, no AI rewriting
everything else → clean, well-punctuated default
On-device STT + LLM — transcription with whisper-large-v3-turbo (multilingual; parakeet-tdt-0.6b-v2 available for streaming English, plus a Hindi/Hinglish specialist) and cleanup/formatting with Qwen3-4B-Instruct-2507-4bit . Cleanup removes fillers, applies self-corrections ("no wait, I meant Tuesday"), punctuates, and honors spoken "new line" / "new paragraph".
History browser — every dictation (raw + final text, app, mode, duration) is stored in a local SQLite database. Browse, search, copy, or paste-again from the History tab; the menubar menu shows your last three (click to copy).
Voice Intelligence — an on-device dashboard turns local history into useful trends: words and dictations over time, estimated time saved, streaks, app and mode breakdowns, STT/cleanup latency, cleanup rate, and an honest zero-edit rate with observation coverage. Share cards contain aggregate numbers only — never transcript, app, or contact text.
Private Meeting Memory — Velora can suggest recording when Zoom, Google Meet, Teams, or a Slack Huddle is active, with optional Calendar matching. Capture starts only after an explicit confirmation, keeps microphone ("Me") and computer audio ("Them") as separate local tracks, and produces a searchable transcript, summary, decisions, and action items in the background. Processing is resumable and live dictation takes priority.
Speaker diarization — when more than one person is on the other side of a call, the transcript splits them into Speaker 1 / Speaker 2 / … using on-device diarization (sherpa-onnx: pyannote segmentation + titanet embeddings, ~46 MB downloaded on the first meeting, sha256-pinned). One-on-one calls stay plain "Them"; any diarization failure falls back cleanly. Toggle in Settings → Meetings.
Safe Voice Edit — select text in any app, press ⌥⇧E (configurable), and speak an edit: "make this more formal", "fix the grammar", "turn this into bullet points". Only the selection is touched, the result replaces it in place, and ⌘Z undoes it. The edit prompt is benchmarked (94% on a 50-command suite, spikes/engine/bench_voice_edit.py ) and guarded — an unusable result keeps your text unchanged, and the edited text is always on the clipboard as backup.
"As Heard" escape hatch — when cleanup gets something wrong, paste the untouched raw transcript from the menubar ( Reformat Last as → As Heard ) or view it in History. No re-processing, works even after the audio clip has been pruned.
Local CLI + MCP — scripts and local agents can inspect allow-listed history/stats, transcribe an audio file, or request one visibly approved live dictation. Access is off by default, runs through an owner-only Unix socket instead of a network server, and never exposes raw audio, screen context, contacts, or learning data.
Personal Dictionary — teach Velora exact names, product terms, and optional “heard as → write as” corrections. Edit-learned and auto-learned words stay visible and reversible, and only confirmed dictionary entries sync through your app-specific iCloud Drive folder when iCloud is available.
Audio archive + reprocess — clips are saved as compact FLAC under ~/.velora/audio (configurable retention, default 6 months / 4 GB cap) so you can re-transcribe any past dictation with a better model straight from the History tab.
Custom modes editor — every mode is a JSON file in ~/.velora/modes/ , editable from the Modes tab: per-mode LLM prompt (the Superwhisper-style feature), formatting level, app bindings, vocabulary, and replacements. Drop in a file to create your own (see Customization ).
Model picker — choose your STT model in Settings from the engine's registry, with managed downloads.
Live spectrum waveform — the HUD's 24-bar waveform is driven by a real FFT of your mic, so bars react to both loudness and pitch.
Safe insertion — clipboard is snapshotted and restored around the synthesized ⌘V, with a keystroke-typing fallback for apps that block paste. Secure input fields (passwords) are detected and insertion is suppressed.
Apple Silicon Mac (M1 or later)
uv (manages the Python engine)
Swift toolchain — Command Line Tools are enough ( xcode-select --install ); no Xcode needed
~4.4 GB of disk for the default models (downloaded on first run)
# 1. Prerequisites (one-time)
xcode-select --install # Swift toolchain (Command Line Tools)
curl -LsSf https://astral.sh/uv/install.sh | sh # uv, for the Python engine
# 2. Build & run
git clone https://github.com/sushilk1991/velora
cd velora
make app # builds build/Velora.app (SwiftPM release + hand-rolled dev bundle)
open build/Velora.app
make app compiles the Swift app and bundles the Python engine; the engine's dependencies are fetched by uv on first launch. First launch then walks you through onboarding: microphone permission, accessibility permission (live-detected as you grant it), hotkey choice, and a try-it playground — you finish with a real dictation. The engine downloads the default models from Hugging Face on first run (~6 GB; live progress shows in the onboarding window, the menubar menu, and the HUD if you try dictating early — speech recognition unlocks first, AI cleanup a few minutes later). After that, everything is offline.
Prefer not to build? Install with Homebrew:
brew install --cask sushilk1991/tap/velora
Or grab Velora-x.y.z.dmg from Releases , drag Velora to Applications, and open it. Releases from v0.4.3 onward are Developer ID-signed and notarized by Apple, so they open normally through Gatekeeper. The older v0.4.1 image predates notarization and should be replaced rather than bypassed.
Run the .app bundle, not the bare binary — macOS permission grants (mic, accessibility) attach to the signed bundle identity.
Tests: make test runs the Python engine suite. The app also has an embedded Swift self-test ( swift run Velora --selftest ), and make perf-test checks Intelligence against a 100,000-row history.
The installed app includes a CLI at:
/Applications/Velora.app/Contents/Resources/bin/velora --help
Enable Allow local CLI and agents in Settings → General, then use status , recent , search , stats , transcribe , or listen . Add --json for machine-readable output. listen always displays an approval prompt before the microphone starts.
Velora also exposes the same narrow surface as a local MCP stdio server:
/Applications/Velora.app/Contents/Resources/bin/velora mcp
The app must be running. Nothing listens on the network, and disabling the setting immediately removes history, stats, and action access; status remains available so tools can explain what is missing.
The cleanup model is picked by RAM tier at first launch (Qwen3-1.7B-8bit on ≤14 GB Macs, Qwen3-4B-4bit up to 24 GB, Qwen3.5-4B-8bit above) and can be changed in Settings → Models. Measured on Apple Silicon M-series:
The default multilingual model is batch (transcribes on release rather than streaming), a deliberate quality tradeoff for accurate Hindi/Indian-English. For streaming English with live HUD partials, pick parakeet-tdt-0.6b-v2 in Settings. If cleanup would blow its budget, Velora inserts the raw transcript immediately inst

[truncated]
