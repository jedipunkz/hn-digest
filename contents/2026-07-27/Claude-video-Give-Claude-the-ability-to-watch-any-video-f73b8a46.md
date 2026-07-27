---
source: "https://github.com/bradautomates/claude-video"
hn_url: "https://news.ycombinator.com/item?id=49070557"
title: "Claude-video – Give Claude the ability to watch any video"
article_title: "GitHub - bradautomates/claude-video: Give Claude the ability to watch any video. /watch downloads, extracts frames, transcribes, hands it all to Claude. · GitHub"
author: "opwizardx"
captured_at: "2026-07-27T15:38:13Z"
capture_tool: "hn-digest"
hn_id: 49070557
score: 1
comments: 0
posted_at: "2026-07-27T14:53:01Z"
tags:
  - hacker-news
  - translated
---

# Claude-video – Give Claude the ability to watch any video

- HN: [49070557](https://news.ycombinator.com/item?id=49070557)
- Source: [github.com](https://github.com/bradautomates/claude-video)
- Score: 1
- Comments: 0
- Posted: 2026-07-27T14:53:01Z

## Translation

タイトル: Claude-video – クロードに任意のビデオを視聴できるようにします
記事のタイトル: GitHub - bradautomates/claude-video: クロードに任意のビデオを視聴できるようにします。 /watch のダウンロード、フレームの抽出、転写をすべてクロードに渡します。 · GitHub
説明: クロードにあらゆるビデオを視聴できるようにします。 /watch のダウンロード、フレームの抽出、転写をすべてクロードに渡します。 - bradautomates/claude-video

記事本文:
GitHub - bradautomates/claude-video: クロードに任意のビデオを視聴できるようにします。 /watch のダウンロード、フレームの抽出、転写をすべてクロードに渡します。 · GitHub
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
ブラッドオートメイツ
/
クロードビデオ
公共
通知

ション
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
11 コミット 11 コミット .agents/ plugins .agents/ plugins .claude-plugin .claude-plugin .codex-plugin .codex-plugin .github/ workflows .github/ workflows フック フック スキル/ ウォッチ スキル/ ウォッチ テスト テスト .gitattributes .gitattributes .gitignore .gitignore .skillignore .skillignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md dev-sync.sh dev-sync.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロードにあらゆるビデオを視聴できるようにします。
Claude Code (推奨 - マーケットプレイス経由の自動更新):
/プラグイン マーケットプレイス追加 bradautomates/claude-video
/プラグインのインストール watch@claude-video
Codex、Cursor、Copilot、Gemini CLI、または 50 以上のエージェント スキル ホストのいずれか:
npx スキル追加 bradautomates/claude-video -g
( -g はユーザーに対してグローバルにインストールされ、すべてのプロジェクトで使用できます。プロジェクトごとのスコープにドロップします。)
その他のインストール オプション (claude.ai Web、手動) については、以下の「インストール」セクションを参照してください。
開始するための設定は不要 — macOS 上の brew を介して最初の実行時に yt-dlp と ffmpeg がインストールされます (Linux/Windows は正確なコマンドを出力します)。キャプションにはほとんどの公開動画が無料で含まれています。 Whisper API キーは、ビデオにキャプションがない場合にのみ必要です。
クロードは、Web ページを読んだり、スクリプトを実行したり、リポジトリを参照したりできます。そのままではできないのは、ビデオの視聴です。 YouTube のリンクを貼り付けると、タイトルから推測するか、画面上の内容の 90% が欠けているトランスクリプトを取得する必要があります。
Claude Video /watch を使用すると、URL またはローカル パスを貼り付けて質問することができます。すると、Claude が最初にキャプションを取得し、必要なものだけをダウンロードし、フレーム (シーンに応じた、または効率的なディテールでの高速キーフレーム) を抽出し、タイムスタンプ付きの TRA を取得します。

nscript (利用可能な場合は無料キャプション、フォールバックとして Whisper API)、およびすべてのフレームを画像として読み取ります。応答するまでに、ビデオは見られ、音声も聞こえます。
/watch https://youtu.be/dQw4w9WgXcQ 30 秒の時点で何が起こっていますか?
実際に人々がそれを何に使用しているか
他人のコンテンツを分析する。 /watch https://youtu.be/<viral-video> どのフックで始まりましたか?クロードは最初のフレームを見て、冒頭のトランスクリプトを読み、構造を分解します。広告クリエイティブ、競合他社のローンチ、ポッドキャストのイントロなど、内容だけでなく方法が重要なあらゆるものについても同様です。
ビデオからバグを診断します。誰かが壊れた画面の録画を送信してきました。 /watch bug-repro.mov 何が問題ですか?クロードは録画を見て、問題が発生しているフレームを見つけ、画面に表示されている内容を説明し、ファイルを開かなくても原因を特定することがよくあります。
動画を要約します。 /watch https://youtu.be/<長いもの> 要約すると、これは明白なことを行います - 構造、重要な瞬間、実際に言われ、示された内容を引き出します。 2倍速で見るよりも速いです。
アップデートビデオから誇大宣伝を削除します。 /watch https://youtu.be/<launch-video> 実際の新機能 — 誇大広告はスキップします 「ゲームチェンジャー」機能を取り除き、重要ないくつかの項目に絞り込みます。これにより、10 分間の紹介や過剰な宣伝をせずに内容を理解できるようになります。
プレイリストをメモに変えます。 /watch https://youtu.be/<ビデオ> これをメモに要約します。シリーズ全体で実行し、ビデオごとの要約をファイルします。これにより、チャンネルやコースが、何時間も座らなければならない代わりに、検索可能なメモのセットになります。
動画と質問を貼り付けます。 URL (yt-dlp がサポートするもの - YouTube、Loom、TikTok、X、Instagram、さらに数百以上) またはローカル パス ( .mp4 、 .mov 、 .mkv 、 .webm )。
yt-dlp は最初にキャプションをチェックします。トランスクリプトの詳細では、キャプション付きの URL が表示されます

ビデオをダウンロードせずに回転します。それ以外の場合、または Whisper がオーディオを必要とする場合、実行に必要なもののみがダウンロードされます。
ffmpeg は、選択された詳細でフレームを抽出します。キーフレームのみを効率的にデコードします (ほぼ瞬時に)。バランスの取れた/トークンバーナーはシーンチェンジフレームを好み、生成が不足している場合は持続時間を考慮した均一サンプラーにフォールバックします。 JPEG はデフォルトで幅 512 ピクセルですが、Claude Read との互換性のために高さ 1998 ピクセルに固定されています。
トランスクリプトは 2 つの場所のうちの 1 つから取得されます。最初に試してみましょう: yt-dlp はソースからネイティブ キャプション (手動または自動生成) を取得します。無料、インスタント、正確といった感じです。フォールバック: モノラル 16 kHz 64 kbps mp3 オーディオ クリップ (~480 kb/分) を抽出し、Whisper (Groq の Whisper-large-v3 (推奨 - より安価で高速) または OpenAI の Whisper-1 に送信します)。
フレームとトランスクリプトはクロードに渡されます。このスクリプトは、t=MM:SS マーカーを含むフレーム パスとタイムスタンプを含むトランスクリプトを出力します。クロード 各フレームを並行して読み取ります。JPEG はコンテキスト内の画像として直接レンダリングします。
クロードは、実際に画面上と音声にあるものに基づいて答えます。 「説明に基づいて」または「タイトルに従って」ではありません。それはフレームを見ました。トランスクリプトを聞いた。ビデオを見た人がそうするように答えます。
掃除。スクリプトは最後に作業ディレクトリを出力します。フォローアップを求めていない場合は、クロードが削除します。
トークンのコストはフレームによって支配されます。すべてのフレームは画像です。画像トークンはすぐに増えます。スクリプトの自動 fps ロジックが存在するため、30 分間のビデオをまばらにスキャンするだけでコンテキストの予算を使い果たすことがなくなり、30 秒のウィンドウに焦点を当てたほうが適切な応答が得られます。
ユーザーが瞬間 (「2:30 頃」、「最後の 30 秒」、「0:45 から 1:00 まで」) に名前を付ける場合は、 --start / --end を渡します。集中モードでは、1 秒あたりの予算がより密になり、2 fps に制限されます。まばらなパスよりもはるかに便利です

彼は全部。
フレーム選択 — キーフレーム (効率的)、シーンチェンジ検出 (バランス型 / トークンバーナー)、またはフォールバックする均一サンプラー — によっても、ほぼ同一のフレームが表示される可能性があります。1 枚のスライドを 90 秒間保持する画面録画では 12 枚のスライドが生成され、それぞれが別の画像として請求されます。重複除去パスは、フレームがクロードに到達する前にそれらをドロップします。デフォルトでは、すべてのフレーム モードで実行されます ( --no-dedup をオフにします)。
1 回の ffmpeg 呼び出しで、抽出された各 JPEG が 16×16 のグレースケール サムネイルにスケールされます。これ以降はすべて純粋な stdlib Python であり、画像ライブラリはありません。
各フレームについて、保持された最後のフレームに対する平均絶対差を計算します (ピクセルごとの平均輝度変化、0 ～ 255 スケール)。
その差がしきい値 ( 2.0 ) 以下の場合、フレームはほぼ重複しているため、ドロップされます。それ以外の場合、それは保持され、新しい参照になります。
フレームバジェットの上限は重複排除後に適用されるため、バジェットは個別のフレームに費やされます。
最後に保持されたフレーム (前のフレームではなく) と比較すると、フレーム間のしきい値を決してトリップしない遅いフェードが検出されます。しきい値は意図的に低く設定されており、構造ではなく絶対的な明るさを測定するため、1 行のコードの差分、行をスクロールする端末、または 2 つの異なる色のフラット スライドはすべて生き残ります。
Frames 行は、折りたたまれた内容を報告します。 14 個の候補から 6 個が選択されました (… 8 個のほぼ重複が除外されました…)。常に動いている映像では何も落とされず、とにかく支払うべき金額を支払うだけです。
--detail ダイヤルは、速度とトークン コストを犠牲にして、視覚的な忠実さを実現します。以下の数値は、49:08 の YouTube ビデオ (1280×720、英語の自動キャプション) に対する実際の実行によるものです。これは、長く、ほとんど静的な画面録画であり、キャップに最も大きな負担がかかるケースです。抽出時間は、事前にダウンロードされたコピーに対するローカル CPU の時間です。ワンタイムダウンロード

3 つのフレーム モードで共有される最大 37 秒 / 76 MB でした。
画像トークンは Anthropic の (幅 × 高さ) / 750 を使用します。デフォルトの幅 512 ピクセルでは、これらの 720p フレームは 512×288、≈197 トークン/フレームです。 --解像度 1024 はその約 4 倍です。トランスクリプトはすべての字幕付きモードで表示されますが、長いビデオではコストが高くなることがよくあります。
フレーム モード全体で 1 つのサンプリング ルール。それぞれが全範囲にわたってすべての候補を検出し、その後、その上限まで偶数サンプル (最初と最後が常に保持される) を検出します。モードの違いは候補ソース (キーフレームとシーン カット) とキャップだけであり、カバレッジの広がり方はまったく異なります。そのため、最後のフレームは途中ではなく、常に最後に表示されます。
効率的なのは速度層 (約 0.5 秒) です。キーフレームのみを再構築するため、すべてのフレームをデコードしてカットを見つけるシーン モードよりも約 40 倍高速です。また、動きの少ない映像ではバランスよりも多くのフレームを返すこともあります (キーフレームの数がシーン カットを上回ります)。 「効率的」とは、フレーム数の減少ではなく、抽出が高速であることを意味します。
token-burner は、上限を超えたバランスからのみ分岐します。このクリップには 116 のカットが含まれていたため、バランスのとれたサンプルは 100 で、token-burner は 116 をすべて維持しました。数百のカットがあるハイモーション ビデオでは、token-burner はすべてを維持します (そして 250 フレームを超えるトークンの警告が表示されます)。バランスは 100 に減らされます。
コールド URL からのエンドツーエンドのトランスクリプトは、これまでで最も安価なモードです。フレーム モードでは、上記の抽出時間に最大 37 秒の共有ダウンロードが追加されます。
表面
インストール
クロード・コード
/plugin Marketplace add bradautomates/claude-video then /plugin install watch@claude-video
Codex、Cursor、Copilot、Gemini CLI、その他 50 個以上
npx スキル追加 bradautomates/claude-video -g
クロード.ai (ウェブ)
watch.skillをダウンロード → 設定 → 機能 → スキル → +
マニュアル/開発
git clone してから、ホストのスキル ディレクトリにスキル/ウォッチをシンボリックリンクします (以下を参照)
クロード・コード
/プラグインマーケットp

レース追加 bradautomates/claude-video
/プラグインのインストール watch@claude-video
後で /plugin update watch@claude-video で更新します。
Codex、Cursor、Copilot、Gemini CLI、および 50 を超えるその他のホスト
エージェント スキル CLI は、検出したエージェントにスキルをインストールします。
npx スキル追加 bradautomates/claude-video -g
-g はユーザーに対してグローバルにインストールします ( ~/.codex/skills 、 ~/.cursor/skills など)。代わりにドロップして現在のプロジェクトにインストールします。便利なフラグ:
-a、--agent <names…> — 特定のホストをターゲットにします。 -コーデックス -カーソル
-l, --list — インストールせずにこのリポジトリ内のスキルを一覧表示します
--copy — シンボリックリンクの代わりにファイルをコピーします (シンボリックリンクをサポートしていないファイルシステムの場合)
CLI は、skills/watch/SKILL.md からスキルを検出し、フォルダー全体 (SKILL.md とそのスクリプト/ランタイム) を自己完結型ユニットとしてコピーします。 SKILL.md は、インストールされた場所に応じて独自のスクリプトを解決するため、どのホストでも同じように動作します。
後で npx skill update watch -g を使用して更新します。
最新リリースから watch.skill をダウンロードします。
「設定」→「機能」→「スキル」に移動します。
まず、[機能] で [コードの実行とファイルの作成] を有効にします。スキルは ffmpeg と yt-dlp にシェルアウトするため、それなしでは実行されません。
リポジトリのクローンを作成し、自己完結型のスキル フォルダーをホストにシンボリックリンクします。

[切り捨てられた]

## Original Extract

Give Claude the ability to watch any video. /watch downloads, extracts frames, transcribes, hands it all to Claude. - bradautomates/claude-video

GitHub - bradautomates/claude-video: Give Claude the ability to watch any video. /watch downloads, extracts frames, transcribes, hands it all to Claude. · GitHub
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
bradautomates
/
claude-video
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
11 Commits 11 Commits .agents/ plugins .agents/ plugins .claude-plugin .claude-plugin .codex-plugin .codex-plugin .github/ workflows .github/ workflows hooks hooks skills/ watch skills/ watch tests tests .gitattributes .gitattributes .gitignore .gitignore .skillignore .skillignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md dev-sync.sh dev-sync.sh View all files Repository files navigation
Give Claude the ability to watch any video.
Claude Code (recommended — auto-updates via marketplace):
/plugin marketplace add bradautomates/claude-video
/plugin install watch@claude-video
Codex, Cursor, Copilot, Gemini CLI, or any of 50+ Agent Skills hosts:
npx skills add bradautomates/claude-video -g
( -g installs globally for your user, available across all projects. Drop it to scope per-project.)
More install options (claude.ai web, manual) in the Install section below.
Zero config to start — yt-dlp and ffmpeg install on first run via brew on macOS (Linux/Windows print exact commands). Captions cover most public videos for free. Whisper API key is only needed when a video has no captions.
Claude can read a webpage, run a script, browse a repo. What it can't do, out of the box, is watch a video . You paste a YouTube link and it has to either guess from the title or pull a transcript that's missing 90% of what's on screen.
With Claude Video /watch you can paste a URL or a local path, ask a question, and Claude fetches captions first, downloads only what it needs, extracts frames (scene-aware, or fast keyframes at efficient detail), pulls a timestamped transcript (free captions when available, Whisper API as fallback), and Read s every frame as an image. By the time it answers, it has seen the video and heard the audio.
/watch https://youtu.be/dQw4w9WgXcQ what happens at the 30 second mark?
What people actually use it for
Analyze someone else's content. /watch https://youtu.be/<viral-video> what hook did they open with? Claude looks at the first frames, reads the opening transcript, breaks down the structure. Same for ad creative, competitor launches, podcast intros, anything where the how matters as much as the what .
Diagnose a bug from a video. Someone sends you a screen recording of something broken. /watch bug-repro.mov what's going wrong? Claude watches the recording, finds the frame where the issue appears, describes what's on screen, often catches the cause without you ever opening the file.
Summarize a video. /watch https://youtu.be/<long-thing> summarize this does the obvious thing — pulls the structure, the key moments, what was actually said and shown. Faster than watching at 2x.
Cut the hype out of an update video. /watch https://youtu.be/<launch-video> what's actually new — skip the hype Strip a "game-changer" feature drop down to the few things that matter, so you get the substance without ten minutes of intro and overselling.
Turn a playlist into notes. /watch https://youtu.be/<video> summarize this to a note Run it across a series and file a per-video summary, so a channel or course becomes a searchable set of notes instead of hours you have to sit through.
You paste a video and a question. URL (anything yt-dlp supports — YouTube, Loom, TikTok, X, Instagram, plus a few hundred more) or a local path ( .mp4 , .mov , .mkv , .webm ).
yt-dlp checks captions first. At transcript detail, captioned URLs return without downloading video. Otherwise, or when Whisper needs audio, it downloads only what the run needs.
ffmpeg extracts frames at the chosen detail. efficient decodes keyframes only (near-instant); balanced / token-burner prefer scene-change frames and fall back to the duration-aware uniform sampler when they under-produce. JPEGs are 512px wide by default and clamped to 1998px tall for Claude Read compatibility.
The transcript comes from one of two places. First try: yt-dlp pulls native captions (manual or auto-generated) from the source. Free, instant, accurate-ish. Fallback: extract a mono 16 kHz 64 kbps mp3 audio clip (~480 kB/min) and ship it to Whisper — Groq's whisper-large-v3 (preferred — cheaper and faster) or OpenAI's whisper-1 .
Frames + transcript are handed to Claude. The script prints frame paths with t=MM:SS markers and the transcript with timestamps. Claude Read s each frame in parallel — JPEGs render directly as images in its context.
Claude answers grounded in what's actually on screen and in the audio. Not "based on the description" or "according to the title." It saw the frames. It heard the transcript. It answers the way someone who watched the video would.
Cleanup. The script prints a working directory at the end. If you're not asking follow-ups, Claude removes it.
Token cost is dominated by frames. Every frame is an image; image tokens add up fast. The script's auto-fps logic exists so you don't blow your context budget on a sparse scan of a 30-minute video that would have been better answered by a focused 30-second window.
When the user names a moment ("around 2:30", "the last 30 seconds", "from 0:45 to 1:00"), pass --start / --end . Focused mode gets denser per-second budgets, capped at 2 fps. Far more useful than a sparse pass over the whole thing.
Frame selection — keyframes ( efficient ), scene-change detection ( balanced / token-burner ), or the uniform sampler it falls back to — can still surface near-identical frames: a screen recording that holds one slide for 90 seconds produces a dozen, each billed as a separate image. A dedup pass drops them before frames reach Claude. It runs by default on every frame mode ( --no-dedup turns it off):
One ffmpeg call scales each extracted JPEG to a 16×16 grayscale thumbnail. Everything after is pure-stdlib Python — no image libraries.
For each frame, compute the mean absolute difference against the last frame that was kept (average per-pixel brightness change, 0–255 scale).
If that difference is at or below the threshold ( 2.0 ), the frame is a near-duplicate and is dropped. Otherwise it's kept and becomes the new reference.
The frame-budget cap applies after dedup, so the budget is spent on distinct frames.
Comparing against the last kept frame (not the previous one) catches slow fades that never trip a frame-to-frame threshold. The threshold is deliberately low and measures absolute brightness rather than structure, so a one-line code diff, a terminal scrolling a row, or two differently-colored flat slides all survive.
The Frames line reports what was collapsed, e.g. 6 selected from 14 candidates (… 8 near-duplicates dropped …) . On always-moving footage nothing is dropped and you pay what you would have anyway.
The --detail dial trades speed and token cost for visual fidelity. Numbers below are from a real run against a 49:08 YouTube video (1280×720, English auto-captions) — a long, mostly-static screen recording, the case that stresses the caps hardest. Extraction times are local CPU against a pre-downloaded copy; the one-time download was ~37 s / 76 MB, shared by the three frame modes.
Image tokens use Anthropic's (width × height) / 750 — at the default 512px width these 720p frames are 512×288, ≈197 tokens/frame ; --resolution 1024 roughly 4×s that. The transcript is surfaced in every captioned mode and on long videos is often the larger cost.
One sampling rule across frame modes. Each detects all candidates across the full range, then even-samples (first + last always kept) down to its cap. The modes differ only in candidate source (keyframes vs. scene cuts) and cap, never in how coverage is spread — so the last frame always lands at the end, not partway through.
efficient is the speed tier (~0.5 s) — it only reconstructs keyframes, so it's ~40× faster than the scene modes, which decode every frame to find cuts. It can also return more frames than balanced on low-motion footage (keyframes outnumber scene cuts); "efficient" means fast extraction, not fewer frames.
token-burner only diverges from balanced past the cap. This clip had 116 cuts, so balanced sampled 100 and token-burner kept all 116. On high-motion video with hundreds of cuts, token-burner keeps everything (and trips the >250-frame token warning) while balanced thins to 100.
End-to-end from a cold URL, transcript is the cheapest mode by far; the frame modes add the shared ~37 s download on top of the extraction times above.
Surface
Install
Claude Code
/plugin marketplace add bradautomates/claude-video then /plugin install watch@claude-video
Codex, Cursor, Copilot, Gemini CLI, +50 more
npx skills add bradautomates/claude-video -g
claude.ai (web)
Download watch.skill → Settings → Capabilities → Skills → +
Manual / dev
git clone then symlink skills/watch into your host's skills dir (see below)
Claude Code
/plugin marketplace add bradautomates/claude-video
/plugin install watch@claude-video
Update later with /plugin update watch@claude-video .
Codex, Cursor, Copilot, Gemini CLI, and 50+ other hosts
The Agent Skills CLI installs the skill into whatever agents it detects:
npx skills add bradautomates/claude-video -g
-g installs globally for your user ( ~/.codex/skills , ~/.cursor/skills , etc.); drop it to install into the current project instead. Useful flags:
-a, --agent <names…> — target specific hosts, e.g. -a codex -a cursor
-l, --list — list the skills in this repo without installing
--copy — copy files instead of symlinking (for filesystems without symlink support)
The CLI discovers the skill from skills/watch/SKILL.md and copies the whole folder — SKILL.md plus its scripts/ runtime — as a self-contained unit. SKILL.md resolves its own scripts relative to wherever it was installed, so it works the same on every host.
Update later with npx skills update watch -g .
Download watch.skill from the latest release.
Go to Settings → Capabilities → Skills.
Enable "Code execution and file creation" under Capabilities first — the skill shells out to ffmpeg and yt-dlp , so it won't run without it.
Clone the repo and symlink the self-contained skill folder into your host'

[truncated]
