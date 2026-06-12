---
source: "https://github.com/oktaydbk54/vibeclip"
hn_url: "https://news.ycombinator.com/item?id=48505420"
title: "Show HN: VibeClip – open-source AI video editor you control by chatting"
article_title: "GitHub - oktaydbk54/vibeclip: Open-source, self-hosted AI video editor: turn long videos into captioned 9:16 shorts — and edit by chatting. BYO LLM key. AGPL-3.0. · GitHub"
author: "borandabak"
captured_at: "2026-06-12T16:06:41Z"
capture_tool: "hn-digest"
hn_id: 48505420
score: 1
comments: 0
posted_at: "2026-06-12T15:30:51Z"
tags:
  - hacker-news
  - translated
---

# Show HN: VibeClip – open-source AI video editor you control by chatting

- HN: [48505420](https://news.ycombinator.com/item?id=48505420)
- Source: [github.com](https://github.com/oktaydbk54/vibeclip)
- Score: 1
- Comments: 0
- Posted: 2026-06-12T15:30:51Z

## Translation

タイトル: Show HN: VibeClip – チャットでコントロールできるオープンソース AI ビデオ エディター
記事のタイトル: GitHub - oktaydbk54/vibeclip: オープンソース、セルフホスト AI ビデオ エディター: 長いビデオをキャプション付きの 9 分 16 秒の短編に変換し、チャットで編集します。 BYO LLM キー。 AGPL-3.0。 · GitHub
説明: オープンソースの自己ホスト型 AI ビデオ エディター: 長いビデオをキャプション付きの 9 分 16 秒の短編に変換し、チャットで編集します。 BYO LLM キー。 AGPL-3.0。 - oktaydbk54/vibeclip

記事本文:
GitHub - oktaydbk54/vibeclip: オープンソースの自己ホスト型 AI ビデオ エディター: 長いビデオをキャプション付きの 9 分 16 秒の短編に変換し、チャットで編集します。 BYO LLM キー。 AGPL-3.0。 · GitHub
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
oktaydbk54
/
バイブクリップ
公共
通知
サインインする必要があります

o 通知設定を変更する
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
10 コミット 10 コミット .github .github アセット アセット チャット チャット ドキュメント ドキュメント パイプライン パイプライン スクリプト スクリプト テスト テスト .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version .secrets.baseline .secrets.baseline CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml server.py server.py uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ポッドキャスト、インタビュー、トーク、ストリームなどの長いビデオをドロップすると、VibeClip がそれを切り出します。
縦型でキャプション付きの、すぐに投稿できるショートパンツ。次に、チャットによってすべてのクリップを調整します。
「クリップ 2 をより迫力のあるものにする」、「キャプションを大きくする」、「0:05 にズームを追加する」、「元に戻す」。
クイックスタート · 機能 · 仕組み · 独自のキーの持ち込み · 構成 · 貢献
左: 生のクリップ。右: 「mrbeast スタイルにして、その下にゲームプレイを追加する」という一文の後にキャプションが付けられ、9 分 16 秒にリフレームされ、分割画面になっています。モックアップではなく、実際のパイプライン出力。
映像: Andy Dickinson (CC-BY) · ゲームプレイ: Orbital - 著作権なしのゲームプレイ (CC-BY) · Minecraft © Mojang.
3 つのコマンドでプライベート インスタンスをスピンアップします。追加するのは 1 つの LLM キーだけです。
git クローン https://github.com/oktaydbk54/vibeclip.git
CDバイブクリップ
cp .env.example .env # 1 行追加: OPENAI_API_KEY=sk-...
docker 構成 -d --build
# → http://localhost:8765 を開く
デフォルト ( EMAIL_MODE=console 、 REQUIRE_EMAIL_VERIFICATION=false ) のサインアップ ログの場合
直接入力してください — メールプロバイダーは必要ありません

エド。 OpenAI または DeepSeek キーを持参する
(DeepSeek が安価です)、または OpenAI 互換サーバーの LLM_BASE_URL を指定します
(Ollama、LM Studio、OpenRouter…)。 Docker を使用しない方がよいですか? 「ローカルインストール」を参照してください。
🎬 ロング→ショート、自動的に
デバイス上で文字起こしし、最も強力な瞬間 (フック / フロー / 価値 - くだらないキーワード スキャンではありません) をスコア化し、スピーカーを中心に 9:16 にリフレームし、単語と同期したキャプションを書き込みます。
💬 チャットで編集
ツール呼び出しエージェントは、平易な言葉を実際の編集、つまりトリミング、つなぎ言葉の削除 (「えー」/「ええ」)、ズーム、スタイル、音楽、B ロール、ブランド オーバーレイに変換します。 1 回元に戻すと、複数ステップの計画全体が元に戻ります。
🎨 ワンショットでスタイル
hormozi 、 mrbeast 、 podcast_minimal 、 kinetic — キャプション、ペース、ズーム、音楽、SFX が一緒に適用されます。独自のプリセットを JSON ファイルとしてドロップします。
🖥️ 本物のスタジオ UI
9:16 のライブ プレビュー、クリップ カード、CapCut スタイルのタイムライン、およびそのすぐ横にあるチャット副操縦士を備えた Web アプリ。
🔑 あなたの鍵、あなたのデータ
独自の LLM キー (OpenAI、Gemini、Claude、DeepSeek、互換性のあるエンドポイント) を持参してください。私たちを介して代理されるものは何もありません。「私たち」は存在しません。
🏠 まずはセルフホスト
1 つの Docker コマンド。 Speech-to-Text とすべてのレンダリングは、faster-whisper + ffmpeg を介してローカルで実行されます。 AGPL-3.0、SaaS ロックインなし。
🛠 仕組み
アップロード
│
┌──────▼────────┐ 高速ウィスパー (ローカル、API キーなし)
│ 転写 │
━━━━┬───────┘
┌──────▼────────┐ LLM「脳」（キー） — 構造 + 得点された瞬間
│ 構造を分析する │
│ ハイライトを探す │
━━━━┬───────┘
┌──────▼────────┐ クリップごと、キャッシュされた中間ファイルから再生（編集あたり約 2 ～ 4 秒）
│ 自動編集 │ ジャンプカット → 9:16 ref

ラメ → キャプション → 音楽+アンビエンス (ダッキング)
│ │ → SFX → フェード · その後、チャット コマンドのレイヤーが上に表示されます
━━━━┬───────┘
エクスポート → 垂直 MP4、公開準備完了
ネットワークに到達するのは 2 つだけです。選択した LLM (意図を理解するため)
スコアモーメント)、およびオプションで Pexels (ストック B ロール)。音声からテキストへの変換など
レンダリングはマシン上に残ります。
VibeClip にはキーが付属しておらず、プロンプトをプロキシすることはありません。
あなたが選んだプロバイダー。 1 つを供給する 2 つの方法:
インスタンスごと — OPENAI_API_KEY (または DEEPSEEK_API_KEY 、または任意の値) を設定します
.env 内の LLM_BASE_URL 経由の OpenAI 互換エンドポイント。
ユーザーごと - 各アカウントは、独自のキーをアプリ内設定ページに貼り付けます。
ライブテスト接続。キーは保管時に暗号化され、ブラウザーに送り返されることはありません。
Speech-to-Text はローカルで実行され、キーは必要ありません。
すべては .env によって駆動されます (完全なコメント付きリストについては、.env.example を参照してください)。それら
最も重要なこと:
要件: Python 3.12+ 、 ffmpeg 、および DejaVu フォント (キャプションのレンダリング用)。
cp .env.example .env # LLM キーを追加します
uv sync # または: pip install -e 。
python -m chat.app # → http://127.0.0.1:8765
最初の実行では、Whisper モデルがダウンロードされます。ターミナルの方がいいですか？ python -m chat.cli <video.mp4> 。
このリポジトリには、ロイヤリティフリーのメディア (音楽、アンビエンス、SFX、デモ) の小さなライブラリがバンドルされています。
フッテージ) 組み込みスタイル用。一部のトラックは CC-BY (Kevin MacLeod) であり、
ビデオの説明にクレジットを記載します。assets/ の下にあるクレジット ファイルを参照してください。バイブクリップ
著作権で保護された/ブランド化されたゲーム映像をバンドルしたり使用したりすることはありません。
問題や PR は歓迎です — まずは CONTRIBUTING.md から始めてください。セキュリティレポート:
SECURITY.md を参照してください。お互いに優れた態度をとりましょう (行動規範)。
GNU AGPL-3.0 — ライセンス を参照してください。 VibeClip は自由に自己ホストして変更できます。
変更を実行すると

d バージョンをネットワーク サービスとして使用するには、その変更されたソースをネットワーク サービスに提供する必要があります。
そのユーザー。著作権 © 2026 VibeClip 作者。
オープンソースの自己ホスト型 AI ビデオ エディター: 長いビデオをキャプション付きの 9 分 16 秒の短編に変換し、チャットで編集できます。 BYO LLM キー。 AGPL-3.0。
AGPL-3.0ライセンス
行動規範
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

Open-source, self-hosted AI video editor: turn long videos into captioned 9:16 shorts — and edit by chatting. BYO LLM key. AGPL-3.0. - oktaydbk54/vibeclip

GitHub - oktaydbk54/vibeclip: Open-source, self-hosted AI video editor: turn long videos into captioned 9:16 shorts — and edit by chatting. BYO LLM key. AGPL-3.0. · GitHub
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
oktaydbk54
/
vibeclip
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
10 Commits 10 Commits .github .github assets assets chat chat docs docs pipeline pipeline scripts scripts tests tests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version .secrets.baseline .secrets.baseline CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml server.py server.py uv.lock uv.lock View all files Repository files navigation
Drop in a long video — podcast, interview, talk, stream — and VibeClip cuts it into
vertical, captioned, ready-to-post shorts. Then you refine every clip by chatting :
“make clip 2 punchier,” “bigger captions,” “add a zoom at 0:05,” “undo.”
Quick start · Features · How it works · Bring your own key · Configuration · Contributing
Left: the raw clip. Right: after one sentence — “make it mrbeast style and add gameplay underneath” — captioned, reframed to 9:16, and split-screened. Real pipeline output, not a mockup.
Footage: Andy Dickinson (CC-BY) · gameplay: Orbital - No Copyright Gameplay (CC-BY) · Minecraft © Mojang.
Spin up a private instance in three commands. All you add is one LLM key.
git clone https://github.com/oktaydbk54/vibeclip.git
cd vibeclip
cp .env.example .env # add ONE line: OPENAI_API_KEY=sk-...
docker compose up -d --build
# → open http://localhost:8765
With the defaults ( EMAIL_MODE=console , REQUIRE_EMAIL_VERIFICATION=false ) sign-up logs
you straight in — no email provider needed. Bring an OpenAI or DeepSeek key
(DeepSeek is the cheap one), or point LLM_BASE_URL at any OpenAI-compatible server
(Ollama, LM Studio, OpenRouter…). Prefer no Docker? See local install .
🎬 Long → shorts, automatically
Transcribes on-device, scores the strongest moments (hook / flow / value — not a dumb keyword scan), reframes to 9:16 around the speaker, and burns word-synced captions.
💬 Edit by chatting
A tool-calling agent turns plain language into real edits — trims, filler-word removal (“uhh”/“ee”), zooms, styles, music, b-roll, brand overlays. One undo reverts a whole multi-step plan.
🎨 Styles in one shot
hormozi , mrbeast , podcast_minimal , kinetic — captions, pace, zoom, music and SFX applied together. Drop in your own preset as a JSON file.
🖥️ A real studio UI
Web app with a live 9:16 preview, clip cards, a CapCut-style timeline, and the chat copilot right beside it.
🔑 Your key, your data
Bring your own LLM key (OpenAI · Gemini · Claude · DeepSeek · any compatible endpoint). Nothing is proxied through us — there is no “us.”
🏠 Self-host first
One Docker command. Speech-to-text and every render run locally via faster-whisper + ffmpeg. AGPL-3.0, no SaaS lock-in.
🛠 How it works
upload
│
┌──────▼───────┐ faster-whisper (local, no API key)
│ transcribe │
└──────┬───────┘
┌──────▼────────────┐ LLM "brain" (your key) — structure + scored moments
│ analyze structure │
│ find highlights │
└──────┬────────────┘
┌──────▼───────┐ per clip, replayed from cached intermediates (~2–4s/edit)
│ auto edit │ jumpcut → 9:16 reframe → captions → music+ambience (ducked)
│ │ → SFX → fades · then your chat commands layer on top
└──────┬───────┘
export → vertical MP4, publish-ready
Only two things ever hit the network: your chosen LLM (to understand intent and
score moments) and, optionally, Pexels (stock b-roll). Speech-to-text and all
rendering stay on your machine.
VibeClip never ships with a key and never proxies your prompts anywhere except the
provider you choose. Two ways to supply one:
Per instance — set OPENAI_API_KEY (or DEEPSEEK_API_KEY , or any
OpenAI-compatible endpoint via LLM_BASE_URL ) in .env .
Per user — each account pastes its own key on the in-app Settings page, with a
live test-connection . Keys are encrypted at rest and never sent back to the browser.
Speech-to-text runs locally and needs no key.
Everything is driven by .env (see .env.example for the full, commented list). The ones
that matter most:
Requirements: Python 3.12+ , ffmpeg , and the DejaVu fonts (for caption rendering).
cp .env.example .env # add your LLM key
uv sync # or: pip install -e .
python -m chat.app # → http://127.0.0.1:8765
First run downloads the Whisper model. Prefer the terminal? python -m chat.cli <video.mp4> .
The repo bundles a small library of royalty-free media (music, ambience, SFX, demo
footage) for the built-in styles. Some tracks are CC-BY (Kevin MacLeod) and require
crediting in your video description — see the CREDITS files under assets/ . VibeClip
never bundles or uses copyrighted/branded game footage.
Issues and PRs welcome — start with CONTRIBUTING.md . Security reports:
see SECURITY.md . Be excellent to each other ( code of conduct ).
GNU AGPL-3.0 — see LICENSE . You can self-host and modify VibeClip freely;
if you run a modified version as a network service, you must offer that modified source to
its users. Copyright © 2026 the VibeClip authors.
Open-source, self-hosted AI video editor: turn long videos into captioned 9:16 shorts — and edit by chatting. BYO LLM key. AGPL-3.0.
AGPL-3.0 license
Code of conduct
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
