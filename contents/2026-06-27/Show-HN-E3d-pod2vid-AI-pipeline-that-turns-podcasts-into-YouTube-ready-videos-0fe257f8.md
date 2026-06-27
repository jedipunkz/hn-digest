---
source: "https://github.com/spacepacket1/e3d-pod2vid"
hn_url: "https://news.ycombinator.com/item?id=48702250"
title: "Show HN: E3d-pod2vid – AI pipeline that turns podcasts into YouTube-ready videos"
article_title: "GitHub - spacepacket1/e3d-pod2vid: AI-powered podcast-to-video pipeline. Semantic B-roll, voice synthesis, burned subtitles, YouTube publishing. · GitHub"
author: "spacepacket"
captured_at: "2026-06-27T22:21:31Z"
capture_tool: "hn-digest"
hn_id: 48702250
score: 1
comments: 0
posted_at: "2026-06-27T22:09:16Z"
tags:
  - hacker-news
  - translated
---

# Show HN: E3d-pod2vid – AI pipeline that turns podcasts into YouTube-ready videos

- HN: [48702250](https://news.ycombinator.com/item?id=48702250)
- Source: [github.com](https://github.com/spacepacket1/e3d-pod2vid)
- Score: 1
- Comments: 0
- Posted: 2026-06-27T22:09:16Z

## Translation

タイトル: Show HN: E3d-pod2vid – ポッドキャストを YouTube 対応ビデオに変換する AI パイプライン
記事のタイトル: GitHub - spacepacket1/e3d-pod2vid: AI を活用したポッドキャストからビデオへのパイプライン。セマンティック B ロール、音声合成、書き込み字幕、YouTube 公開。 · GitHub
説明: AI を活用したポッドキャストからビデオへのパイプライン。セマンティック B ロール、音声合成、書き込み字幕、YouTube 公開。 - スペースパケット1/e3d-pod2vid
HN テキスト: .mpa ファイルをアニメーション ビデオに変換します。

記事本文:
GitHub - spacepacket1/e3d-pod2vid: AI を活用したポッドキャストからビデオへのパイプライン。セマンティック B ロール、音声合成、書き込み字幕、YouTube 公開。 · GitHub
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
スペースパケット1
/
e3d-pod2vid
公共
通知
通知を変更するにはサインインする必要があります

カチオン設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット .env.example .env.example .gitignore .gitignore README.md README.md cancel.js Anan.js linkedin_auth.js linkedin_auth.js make_thumbnail.py make_thumbnail.py package.json package.json pod2vid.py pod2vid.pyrequirements.txt requirements.txt tts_replace.py tts_replace.py yt_auth.js yt_auth.js yt_update.js yt_update.js yt_upload.js yt_upload.js すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI を活用したポッドキャストからビデオへのパイプライン。以下を使用して、日記化されたオーディオ ファイル (NotebookLM、ポッドキャスト、インタビュー) を YouTube 対応 MP4 に変換します。
発話ごとに意味的に一致した Pexels B ロール (GPT-4o-mini がクリップを選択)
焼き付けられた字幕 (ffmpeg libass は必要ありません - 純粋な Pillow)
オプションの OpenAI TTS 音声置換 (NotebookLM / AI 音声の交換)
YouTube アップロード + 説明/サムネイルの更新
ワンショットのマルチプラットフォームソーシャル投稿 (Discord、Telegram、X、Moltbook、LinkedIn)
git クローン https://github.com/spacepacket1/e3d-pod2vid.git
cd e3d-pod2vid
# Python デプス
pip install -r 要件.txt
ノード数 (YouTube + ソーシャル投稿のみ)
npmインストール
# API キーをコピーして入力します
cp .env.example .env
$EDITOR .env
ワークフロー
python3 pod2vid.py エピソード.m4a 出力/episode.mp4
この 1 つのコマンド:
音声を AssemblyAI にアップロードして話者ダイアライゼーションを行う
発話ごとに GPT-4o-mini に特定の Pexels 検索クエリを要求します
一致する B ロール クリップをダウンロードします (クエリごとにキャッシュされます)
焼き付けられた字幕を使用して各セグメントをレンダリングします
最終的な MP4 + SRT 字幕ファイルに連結します
ダイアライゼーションとクエリを JSON としてキャッシュするため、再実行が高速になります。
2. (オプション) OpenAI TTS で音声を置き換える
オリジナルのオーディオではなくカスタムの音声が必要な場合 (例: NotebookLM の音声を置き換える):
#シン

OpenAI TTS 音声を使用したサイズ
python3 tts_replace.py 出力/episode-diarization.json エピソード-tts
# TTS オーディオを使用してビデオをレンダリングする
python3 pod2vid.py 出力/episode-tts.mp3 出力/episode-tts.mp4
デフォルトの音声: onyx (スピーカー A) および nova (スピーカー B)。 VOICE_A / VOICE_B でオーバーライドします。
利用可能な音声: アロイ、エコー、フェイブル、オニキス、ノヴァ、シマー
python3 make_thumbnail.py "自律型 AI エージェント用の予測 GPS" sumnail.png /path/to/logo.png
タイトル、アクセント ストライプ、およびオプションのロゴ オーバーレイを含む 1280×720 PNG を出力します。 Pure Pillow — ブラウザやデザインツールは必要ありません。
初回: アカウントを認証する
ノードyt_auth.js
スクリプトは URL を出力します。任意のデバイス (電話、ブラウザ - スクリプトを実行するマシンにはブラウザは必要ありません) で開きます。承認後、リダイレクト URL をターミナルに貼り付けます。トークンは youtube-tokens.json に保存されます。
ノード yt_upload.js 出力/episode-tts.mp4 「私のエピソードのタイトル」
完了したら、ビデオの URL と ID を出力します。
説明とサムネイルを更新する
YT_DESCRIPTION= " 自動運転車用の AI 搭載 GPS である、maps.e3d.ai をチェックしてください。
フォローしてください:
• X: @e3dmaps
• Discord: https://discord.gg/your-server " \
ノード yt_update.js VIDEO_ID サムネイル.png
5.ソーシャルメディアで発表する
ノード発表.js https://www.youtube.com/watch ? v=VIDEO_ID " 新しいエピソード: 自律型 AI エージェントのための予測 GPS "
構成されているすべてのプラットフォームに同時に投稿します。認証情報のないプラットフォームは通知なくスキップされます。
LinkedIn の API では、announce.js が投稿できるようになる前に、いくつかの 1 回限りのセットアップ手順が必要です。
ステップ 1 — LinkedIn アプリを作成する
linkedin.com/developers/apps に移動し、アプリを作成します。 [認証] タブで、これを承認されたリダイレクト URL として追加します。
https://www.linkedin.com/developers/tools/oauth/redirect
ステップ 2 — 必要な製品を追加する
「製品」タブでアクセスをリクエストします

両方に:
LinkedIn で共有 — w_member_social スコープを付与します (ユーザーに代わって投稿)
OpenID Connect を使用して LinkedIn でサインイン — openid プロファイル スコープを付与します (個人 URN を解決するために必要)
通常、個人用アプリの場合はどちらも即座に承認されます。
ステップ 3 — 会社の関連付けを確認する (プロンプトが表示された場合)
LinkedIn では、企業ページの関連付けを確認するよう求められる場合があります。ページ管理者としてログインした状態で認証 URL を開き、承認します。
ステップ 4 — トークンを承認して取得する
アプリの認証情報を .env に追加します。
LINKEDIN_CLIENT_ID=your_client_id
LINKEDIN_CLIENT_SECRET=your_client_secret
次に、次を実行します。
ノード linkedin_auth.js
印刷された URL を任意のデバイスで開きます。承認したら、リダイレクト URL を貼り付けて戻します。トークンは linkedin-tokens.json に保存されます。
LinkedIn の API には、エンコードされた個人 ID (数値メンバー ID ではありません) が必要です。それを見つけるには:
ブラウザで LinkedIn プロフィールに移動します
ページ ソースを表示し (Cmd+U / Ctrl+U)、urn:li:member: を検索します。
数値 ID をメモします (例: 4435724 )
テスト API 呼び出しを実行します。エラー応答により、エンコードされた個人 URN (例: urn:li:person:2KqUAyg4oY ) が明らかになります。
または、トークンを取得した後に次のワンライナーを実行します。
ノード -e "
const https = require('https');
const t = JSON.parse(require('fs').readFileSync('linkedin-tokens.json'));
// MEMBER_ID をページ ソースの数値 ID に置き換えます
const body = JSON.stringify({author:'urn:li:member:MEMBER_ID',commentary:'test',visibility:'PUBLIC',distribution:{feedDistribution:'MAIN_FEED',targetEntities:[],thirdPartyDistributionChannels:[]},lifecycleState:'PUBLISHED',isReshareDisabledByAuthor:false});
const u = require('url').parse('https://api.linkedin.com/rest/posts');
const r = https.request(Object.assign(u,{method:'POST',headers:{'Authorization':'Bearer '+t.access_token,'Content-Type':'application/json','Content-Length':Buffer.b

yteLength(body),'LinkedIn-Version':'202506','X-Restli-Protocol-Version':'2.0.0'}}),res=>{let d='';res.on('data',c=>d+=c);res.on('end',()=>console.log(d.slice(0,300)));});
r.write(body);r.end();
」
エラー メッセージには、エンコードされた URN が含まれます。保存します:
ノード -e "
const fs = require('fs');
const t = JSON.parse(fs.readFileSync('linkedin-tokens.json'));
t.person_urn = 'urn:li:person:YOUR_ENCODED_ID';
fs.writeFileSync('linkedin-tokens.json', JSON.stringify(t, null, 2));
」
linkedin-tokens.json に person_urn が含まれると、announce.js が自動的に LinkedIn に投稿します。
.env.example を .env にコピーし、必要なキーを入力します。
このパイプラインは、固定クリップ ライブラリを介して回転する代わりに、GPT-4o-mini に発話ごとに特定の Pexels 検索クエリを生成するように要求します。
「EZPass のおかげで、すべての料金所で 90 秒節約できました」
→「料金所・高速道路料金支払い」
「二重証人問題」
→「法廷裁判官証言」
「機械学習による位置予測」
→「機械学習データトレーニングループ」
クエリはキャッシュされるため、再実行や TTS 音声スワップによって API クレジットが再消費されることはありません。 90 セグメントのエピソード全体で最大 82 個の固有のクリップが一般的です。
ffmpeg (任意のバージョン - 字幕レンダリングには libfreetype/libass は必要ありません)
Pexels (B ロール クリップ、個人使用の場合は無料枠)
YouTube Data API v3 (Google Cloud コンソール経由)
LinkedIn API (LinkedIn 開発者ポータル経由) — 投稿用のオプション
出力/
エピソード.mp4 最終ビデオ
YouTube CC 用のepisode.srt 字幕ファイル
エピソード-diarization.json がキャッシュした AssemblyAI の結果
episode-queries.json がキャッシュした GPT Pexels クエリ
ブロール/キャッシュされた B ロール クリップ (一意のクエリごとに 1 つ)
tts-cache/ キャッシュされた TTS 発話 (音声 + テキスト ハッシュごと)
クレジット
E3D Maps によって構築 — AI を活用した自動運転車用ナビゲーション。
AI を活用したポッドキャストからビデオへのパイプライン。セマンティック B ロール、音声合成

sis、焼き付けられた字幕、YouTube 公開。
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

AI-powered podcast-to-video pipeline. Semantic B-roll, voice synthesis, burned subtitles, YouTube publishing. - spacepacket1/e3d-pod2vid

turn your .mpa files into animated videos.

GitHub - spacepacket1/e3d-pod2vid: AI-powered podcast-to-video pipeline. Semantic B-roll, voice synthesis, burned subtitles, YouTube publishing. · GitHub
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
spacepacket1
/
e3d-pod2vid
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits .env.example .env.example .gitignore .gitignore README.md README.md announce.js announce.js linkedin_auth.js linkedin_auth.js make_thumbnail.py make_thumbnail.py package.json package.json pod2vid.py pod2vid.py requirements.txt requirements.txt tts_replace.py tts_replace.py yt_auth.js yt_auth.js yt_update.js yt_update.js yt_upload.js yt_upload.js View all files Repository files navigation
AI-powered podcast-to-video pipeline. Converts a diarized audio file (NotebookLM, podcast, interview) into a YouTube-ready MP4 with:
Semantically matched Pexels B-roll per utterance (GPT-4o-mini picks the clip)
Burned-in subtitles (no ffmpeg libass required — pure Pillow)
Optional OpenAI TTS voice replacement (swap out NotebookLM / AI voices)
YouTube upload + description/thumbnail update
One-shot multi-platform social posting (Discord, Telegram, X, Moltbook, LinkedIn)
git clone https://github.com/spacepacket1/e3d-pod2vid.git
cd e3d-pod2vid
# Python deps
pip install -r requirements.txt
# Node deps (YouTube + social posting only)
npm install
# Copy and fill in your API keys
cp .env.example .env
$EDITOR .env
Workflow
python3 pod2vid.py episode.m4a output/episode.mp4
This single command:
Uploads audio to AssemblyAI for speaker diarization
Asks GPT-4o-mini for a specific Pexels search query per utterance
Downloads matching B-roll clips (cached per query)
Renders each segment with burned-in subtitles
Concatenates into a final MP4 + SRT subtitle file
Caches diarization and queries as JSON so re-runs are fast.
2. (Optional) Replace voices with OpenAI TTS
If you want custom voices instead of the original audio (e.g. replace NotebookLM voices):
# Synthesize with OpenAI TTS voices
python3 tts_replace.py output/episode-diarization.json episode-tts
# Render video using TTS audio
python3 pod2vid.py output/episode-tts.mp3 output/episode-tts.mp4
Default voices: onyx (Speaker A) and nova (Speaker B). Override with VOICE_A / VOICE_B .
Available voices: alloy , echo , fable , onyx , nova , shimmer
python3 make_thumbnail.py " Predictive GPS for Autonomous AI Agents " thumbnail.png /path/to/logo.png
Outputs a 1280×720 PNG with title, accent stripe, and optional logo overlay. Pure Pillow — no browser or design tool required.
First time: authorize your account
node yt_auth.js
The script prints a URL. Open it on any device (phone, browser — the machine running the script doesn't need a browser). After approving, paste the redirect URL back into the terminal. Tokens are saved to youtube-tokens.json .
node yt_upload.js output/episode-tts.mp4 " My Episode Title "
Prints the video URL and ID when done.
Update description and thumbnail
YT_DESCRIPTION= " Check out maps.e3d.ai — AI-powered GPS for autonomous vehicles.
Follow us:
• X: @e3dmaps
• Discord: https://discord.gg/your-server " \
node yt_update.js VIDEO_ID thumbnail.png
5. Announce on social media
node announce.js https://www.youtube.com/watch ? v=VIDEO_ID " New episode: Predictive GPS for Autonomous AI Agents "
Posts simultaneously to all configured platforms. Platforms with no credentials are silently skipped.
LinkedIn's API requires a few one-time setup steps before announce.js can post there.
Step 1 — Create a LinkedIn app
Go to linkedin.com/developers/apps and create an app. Under the Auth tab, add this as an authorized redirect URL:
https://www.linkedin.com/developers/tools/oauth/redirect
Step 2 — Add required products
Under the Products tab, request access to both:
Share on LinkedIn — grants w_member_social scope (post on behalf of user)
Sign In with LinkedIn using OpenID Connect — grants openid profile scopes (needed to resolve your person URN)
Both are typically approved instantly for personal apps.
Step 3 — Verify company association (if prompted)
LinkedIn may ask you to verify a company page association. Open the verification URL while logged in as a Page Admin and approve it.
Step 4 — Authorize and get tokens
Add your app credentials to .env :
LINKEDIN_CLIENT_ID=your_client_id
LINKEDIN_CLIENT_SECRET=your_client_secret
Then run:
node linkedin_auth.js
Open the printed URL on any device. After approving, paste the redirect URL back. Tokens are saved to linkedin-tokens.json .
LinkedIn's API requires your encoded person ID (not your numeric member ID). To find it:
Go to your LinkedIn profile in a browser
View Page Source (Cmd+U / Ctrl+U) and search for urn:li:member:
Note the numeric ID (e.g. 4435724 )
Make a test API call — the error response will reveal your encoded person URN (e.g. urn:li:person:2KqUAyg4oY )
Or run this one-liner after getting a token:
node -e "
const https = require('https');
const t = JSON.parse(require('fs').readFileSync('linkedin-tokens.json'));
// Replace MEMBER_ID with your numeric ID from page source
const body = JSON.stringify({author:'urn:li:member:MEMBER_ID',commentary:'test',visibility:'PUBLIC',distribution:{feedDistribution:'MAIN_FEED',targetEntities:[],thirdPartyDistributionChannels:[]},lifecycleState:'PUBLISHED',isReshareDisabledByAuthor:false});
const u = require('url').parse('https://api.linkedin.com/rest/posts');
const r = https.request(Object.assign(u,{method:'POST',headers:{'Authorization':'Bearer '+t.access_token,'Content-Type':'application/json','Content-Length':Buffer.byteLength(body),'LinkedIn-Version':'202506','X-Restli-Protocol-Version':'2.0.0'}}),res=>{let d='';res.on('data',c=>d+=c);res.on('end',()=>console.log(d.slice(0,300)));});
r.write(body);r.end();
"
The error message will contain your encoded URN. Save it:
node -e "
const fs = require('fs');
const t = JSON.parse(fs.readFileSync('linkedin-tokens.json'));
t.person_urn = 'urn:li:person:YOUR_ENCODED_ID';
fs.writeFileSync('linkedin-tokens.json', JSON.stringify(t, null, 2));
"
Once linkedin-tokens.json contains person_urn , announce.js will post to LinkedIn automatically.
Copy .env.example to .env and fill in the keys you need.
Instead of rotating through a fixed clip library, this pipeline asks GPT-4o-mini to generate a specific Pexels search query for each utterance:
"EZPass saved us 90 seconds at every toll plaza"
→ "toll booth highway payment"
"the dual-witness problem"
→ "courtroom judge testimony"
"machine learning position predictions"
→ "machine learning data training loop"
Queries are cached so re-runs or TTS voice swaps don't re-spend API credits. ~82 unique clips across a 90-segment episode is typical.
ffmpeg (any version — subtitle rendering does not require libfreetype/libass)
Pexels (B-roll clips, free tier fine for personal use)
YouTube Data API v3 (via Google Cloud Console)
LinkedIn API (via LinkedIn Developer Portal ) — optional, for posting
output/
episode.mp4 final video
episode.srt subtitle file for YouTube CC
episode-diarization.json cached AssemblyAI result
episode-queries.json cached GPT Pexels queries
broll/ cached B-roll clips (one per unique query)
tts-cache/ cached TTS utterances (per voice+text hash)
Credits
Built by E3D Maps — AI-powered navigation for autonomous vehicles.
AI-powered podcast-to-video pipeline. Semantic B-roll, voice synthesis, burned subtitles, YouTube publishing.
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
