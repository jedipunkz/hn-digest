---
source: "https://github.com/ronak-create/FableCut"
hn_url: "https://news.ycombinator.com/item?id=48845422"
title: "Show HN: FableCut – A browser video editor AI agents can drive (zero deps)"
article_title: "GitHub - ronak-create/FableCut: Zero-dependency browser video editor that AI agents can drive — JSON timeline, MCP + REST, live-reloading UI · GitHub"
author: "ronak_parmar"
captured_at: "2026-07-09T13:43:19Z"
capture_tool: "hn-digest"
hn_id: 48845422
score: 11
comments: 3
posted_at: "2026-07-09T13:23:10Z"
tags:
  - hacker-news
  - translated
---

# Show HN: FableCut – A browser video editor AI agents can drive (zero deps)

- HN: [48845422](https://news.ycombinator.com/item?id=48845422)
- Source: [github.com](https://github.com/ronak-create/FableCut)
- Score: 11
- Comments: 3
- Posted: 2026-07-09T13:23:10Z

## Translation

タイトル: Show HN: FableCut – AI エージェントが駆動できるブラウザー ビデオ エディター (ゼロ デプス)
記事タイトル: GitHub - ronak-create/FableCut: AI エージェントが駆動できる依存関係ゼロのブラウザー ビデオ エディター — JSON タイムライン、MCP + REST、ライブリロード UI · GitHub
説明: AI エージェントが駆動できる依存関係のないブラウザー ビデオ エディター - JSON タイムライン、MCP + REST、ライブリロード UI - ronak-create/FableCut

記事本文:
GitHub - ronak-create/FableCut: AI エージェントが駆動できる依存関係のないブラウザー ビデオ エディター - JSON タイムライン、MCP + REST、ライブリロード UI · GitHub
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
ロナク作成
/
寓話カット
公共
通知
通知を変更するにはサインインする必要があります

n設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
9 コミット 9 コミット .github .github docs docs エクスポート エクスポート アイコン アイコン ライブラリ ライブラリ メディア メディア .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md CONTRIBUTING.mdライセンス ライセンス README.md README.md SECURITY.md SECURITY.md Analyst.js Analyst.js app.js app.js favicon.svg favicon.svg Index.html Index.html mcp-server.js mcp-server.js package.json package.json server.js server.js style.css style.css すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントが駆動できるブラウザービデオエディター。
FableCut は、Premiere スタイルのノンリニア ビデオ エディターです。
ブラウザ — タイムライン全体を 1 つの JSON ドキュメントとして公開します。手作業で編集して、
UI から、または AI エージェント (Claude Code、Claude Desktop、または
MCP/REST を話します) タイムラインの更新を見ながらビデオをカットします
生きる。
npm 依存関係はゼロです。 1 つのノードの server.js 。それでおしまい。
ほとんどの「AI ビデオ」ツールは編集を API の背後に隠します。 FableCut は次のように反転します。
プロジェクト ファイルはインターフェイスです。 project.json はメディア、クリップ、
トラック、エフェクト、キーフレーム、トランジション - JSON を書き込むことができるプロセス
ビデオを編集でき、開いているブラウザ UI は 150 ミリ秒以内にホットリロードされます。
サーバーから送信されたイベント。人間とエージェントは同じタイムラインで作業できます。
同じ時間です。
4 ビデオ トラック + 3 オーディオ トラック、ドラッグ/トリム/分割/スナップ、元に戻す/やり直し
タイムラインの複数選択 - ラバーバンドのマーキー (空のトラック領域にドラッグ)、
Ctrl/Cmd/Shift+クリックしてクリップを追加/削除し、Ctrl+A でクリップを追加/削除します。
すべて選択し、Esc で選択を解除します。選択したクリップをドラッグして移動します。
グループ全体。削除は選択したすべてを削除します。 S はすべてを分割します
再生ヘッドで選択されます。シュ警部

「N クリップが選択されました」バナーが表示されます。
エッジスナップ付きビート＆キューマーカー（再生中にビート上でMをタップ）
クリップ上の実際のデコードされたオーディオ波形
キャンバス アスペクト プリセット (16:9、9:16 リール、4:5、1:1) + セーフエリア ガイド
12 のワンクリックフィルタープリセット (シネマティック、ティールオレンジ、ノワール、ヴィンテージ、サイバーパンクなど)
調整レイヤー — 1 つのクリップがその下にあるすべてのものをグレーディングします。Premiere スタイル
フルグレードコントロール: 明るさ/コントラスト/彩度/色相、温度と色合い、
ぼかし、グレースケール/セピア/反転、ビネット、アニメーション フィルム グレイン
ブレンド モード (スクリーン、乗算、オーバーレイなど)、フィット モード (包含/カバー/ストレッチ)、
エッジごとのトリミング、コーナー半径、H/V 反転
耐性/柔らかさ + こぼれ抑制を備えたクロマ キー (グリーン スクリーン)
AI 背景除去 (人物の切り抜き、MediaPipe 経由のブラウザ内)
イージングを使用した約 25 のプロパティのキーフレーム アニメーション
速度ランプ - キーフレーム速度とエンジンのタイムリマップビデオと
オーディオ ミックスのエクスポート (高速から低速へのリールの動き)
手振れと RGB 分割/色収差、両方ともアニメーション化可能
17 のトランジション: フェード、スライド、ワイプ (4 方向)、ズーム、アイリス、スピン、ブラー、
ホイップパン、グリッチ、ポップ
キネティックキャプション: タイプライター、ワードポップ、ワードスライド、カラオケ、レターポップ、
波、跳ねる、揺れる
TikTokのキャプションルックのネオンの輝き
フォント エディタ: システム フォント、ドロップイン カスタム フォント ( library/fonts/ )、およびその他
名前による Google Font — 自動的にロードされます
グラデーション塗りつぶし、輪郭、背景丸薬、文字間隔、行の高さ、
ウェイト、斜体、大文字、配置、ソフト シャドウ
ファーストクラスの SVG クリップの種類: CSS- @keyframes -animated SVGs render
プレビューおよびエクスポートでフレームを正確に実行します (コンポジターはフレームをフリーズします)
いつでもアニメーション）。エージェントは独自のベクター オーバーレイを作成できます —
ローワー サード、紙吹雪、輝き - プレーンな .svg ファイルとして。スターターが含まれています。
参照編集を加えてください (a ree

好みに応じて)、編集ブループリントを取得します。
ショット境界、音楽ビート + BPM、ラウドネス カーブ、ショットごとのエネルギー、
ドロップ — メディアに抽出されたリファレンスの音楽トラックを加えて準備完了
同じアイデアを自分の映像で再構築します。余分な依存関係はゼロ
(ffmpeg がデコードを行います。オンセット/テンポの検出はプレーン ノードです)。
ノードanalyze.js ref.mp4 、POST /api/analyze 、または
fablecut_analyze_reference MCP ツール。
ライブラリ/フォルダーは UI のタブとして表示されます: 要素 (オーバーレイ アート)、
Sound FX、SVG — ファイルをドロップすると、開いているエディターがライブで更新されます
高速エクスポート: ブラウザーはすべてのフレームとオフライン オーディオ ミックス、ffmpeg をレンダリングします。
フレーム精度の CRF-18 MP4 をエンコードします (タブを切り替えてもレンダリングを継続します)
ffmpeg が利用できない場合のリアルタイム MediaRecorder フォールバック
git clone https://github.com/ronak-create/FableCut.git
cd フェイブルカット
ノードserver.js # → http://localhost:7777
要件: ノード 18 以降および Chromium ベースのブラウザ。 PATH 上の ffmpeg は
オプションですが推奨されます (高速エクスポート + アップロードの再多重化)。 AIの背景
削除は、最初の使用時に CDN からそのモデルを取得します。
メディアをウィンドウ (または ./media/ ) にドロップし、クリップをタイムラインにドラッグし、編集します。
輸出する。
エージェントに必要なものはすべて CLAUDE.md にあります。
スキーマ、セマンティクス、レシピ。対応可能なモデルをそのファイルに指定すると、
エディタをエンドツーエンドで操作します。
3 つの同等のコントロール サーフェス:
MCP (Claude Code / Claude Desktop に最適) — バンドルされたものを登録します
依存性ゼロの MCP サーバーを 1 回:
クロード mcp add -s user fablecut -- ノード " <path-to>/fablecut/mcp-server.js "
ツール: fablecut_status (エディターを自動起動)、 fablecut_docs 、
fablecut_get_project 、 fablecut_set_project 、 fablecut_patch_project 、
fablecut_import_media 、 fablecut_analyze_reference 。
サーフェスは設計によりトークン効率が高くなります。エージェントはタイムラインにパッチを適用します。
スマ

全体をラウンドトリップする代わりに ll ops ( fablecut_patch_project )
ドキュメントでは、クリップごとに 1 行のコンパクトな要約を読んでください。
( fablecut_get_project {compact:true} )、マニュアルセクションのみをフェッチします
( fablecut_docs {section:"props"} ) が必要です。
ファイル — project.json を読み取り、変更し、リビジョンを変更し、書き込みます。 UI
ライブリロード。
REST — GET/PUT /api/project 、POST /api/upload 、GET /api/library 、
/api/events の SSE。完全なリストについては、CLAUDE.md を参照してください。
例: Claude Code に「これら 6 つのクリップをビート マーカーに合わせてカットし、
ティールオレンジのグレードで、上部にワードポップのキャプションを付け、すべてのカットにシューという音を立てます。」
そしてタイムラインが自動的に再構築されるのを見てください。
または、リファレンスを渡します。「これが私が気に入っているリールです。それを分析して、これを使って作り直してください」
私のクリップ、同じ音楽」。エージェントは fablecut_analyze_reference を呼び出し、
ブループリント (カット、ビート、BPM、エネルギー、ドロップ、抽出された音楽) を作成し、再構築します。
映像をショットごとに構造化します。
競合安全な同時編集: UI、MCP ツール、およびダイレクト
project.json はリビジョン カウンターにすべての同意を書き込みます。でクリップを編集すると、
エージェントがタスク中の UI では、エージェントの次の書き込みが拒否されます (からの 409
REST API / fablecut_set_project からの競合エラー) の代わりに
変更を黙って上書きします。 UI は、エージェントが書き込みを行ったときも同様に検出します。
まだ保存されていないローカル調整を置き換え、代わりにトーストで通知します。
黙って落とします。
server.js 依存関係のない HTTP サーバー: 静的ホスティング、REST API、SSE、
ffmpeg エクスポート パイプライン
app.js エディター: タイムライン UI、コンポジター、キーフレーム、テキスト エンジン、
SVGラスタライザー、クロマキー、エクスポーター
Index.html シングルページ UI
style.css ダークエディターテーマ
mcp-server.js stdio AI エージェントにエディターを公開する MCP サーバー
analyze.js リファレンスビデオアナライザー: ショット、ビート/BPM、エネルギー、ドロップ、
音楽抽出

(モジュール + CLI)
CLAUDE.md エージェント マニュアル (スキーマ + レシピ) — fablecut_docs からも提供されます
project.json タイムライン (最初の実行時に作成、gitignored)
メディア/プロジェクト映像 (Gitgnore)
/api/analyze からの分析/キャッシュされた編集ブループリント (gitignored)
ライブラリ/デフォルトのアセット: 要素/sfx/svg/フォント/
エクスポート/完了したレンダリング (gitignored)
アニメーション SVG オーバーレイのオーサリング
SVG はプレーン CSS @keyframes を使用してアニメーション化されます。 1 つの規則: 決してハードコーディングしない
アニメーション遅延 — 代わりに --d: 0.4 秒を設定すると、コンポジターは時間を次のように駆動します。
すべてのアニメーションを一時停止し、遅延をリベースします。完全なルール + スケルトン
クロード.md ;働いている
例は library/svg/ にあります。
このリポジトリには 20 個の Google フォントが付属しています ( library/fonts/ 、OFL — を参照)
LICENSES.md があります)、および自己作成の SVG オーバーレイとアニメーションのセット
要素 (リポジトリの残りの部分と同様に library/elements/ 、 library/svg/ 、MIT)。
library/sfx/ はあなたのものです (gitignored): 効果音サイトは通常
パブリック リポジトリでファイルを再配布することは許可されていないため、FableCut は許可しません —
library/sfx/README.md には、優れた無料のソースがリストされています。
コンポジターはブラウザーであるため、エクスポートはブラウザーで実行されます。エージェント
[エクスポート] をクリックするように求められます (または、 media/ から ffmpeg を使用して直接レンダリングします)。
AI エージェントが駆動できる依存関係のないブラウザー ビデオ エディター - JSON タイムライン、MCP + REST、ライブリロード UI
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

Zero-dependency browser video editor that AI agents can drive — JSON timeline, MCP + REST, live-reloading UI - ronak-create/FableCut

GitHub - ronak-create/FableCut: Zero-dependency browser video editor that AI agents can drive — JSON timeline, MCP + REST, live-reloading UI · GitHub
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
ronak-create
/
FableCut
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
9 Commits 9 Commits .github .github docs docs exports exports icons icons library library media media .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md analyze.js analyze.js app.js app.js favicon.svg favicon.svg index.html index.html mcp-server.js mcp-server.js package.json package.json server.js server.js style.css style.css View all files Repository files navigation
A browser video editor that AI agents can drive.
FableCut is a Premiere-style non-linear video editor that runs entirely in your
browser — and exposes its whole timeline as one JSON document. Edit it by hand,
from the UI, or let an AI agent (Claude Code, Claude Desktop, or anything that
speaks MCP/REST) cut your video for you while you watch the timeline update
live.
Zero npm dependencies. One node server.js . That's it.
Most "AI video" tools hide the edit behind an API. FableCut flips that: the
project file is the interface . project.json describes media, clips,
tracks, effects, keyframes and transitions — any process that can write JSON
can edit video, and the open browser UI hot-reloads within ~150 ms via
server-sent events. A human and an agent can work on the same timeline at the
same time.
4 video tracks + 3 audio tracks, drag/trim/split/snap, undo/redo
Timeline multi-select — rubber-band marquee (drag on empty track area),
Ctrl/Cmd/Shift+click to add/remove clips, Ctrl+A to
select all, Esc to deselect. Drag any selected clip to move the
whole group; Delete removes all selected; S splits all
selected at the playhead. Inspector shows an "N clips selected" banner.
Beat & cue markers (tap M on the beat during playback) with edge snapping
Real decoded audio waveforms on clips
Canvas aspect presets (16:9, 9:16 reels, 4:5, 1:1) + safe-area guides
12 one-click filter presets (cinematic, teal-orange, noir, vintage, cyberpunk…)
Adjustment layers — one clip grades everything below it, Premiere-style
Full grade controls: brightness/contrast/saturation/hue, temperature & tint ,
blur, grayscale/sepia/invert, vignette , animated film grain
Blend modes (screen, multiply, overlay…), fit modes (contain/cover/stretch),
per-edge cropping, corner radius, flip H/V
Chroma key (green screen) with tolerance/softness + spill suppression
AI background removal (person cut-out, in-browser via MediaPipe)
Keyframe animation on ~25 properties with easing
Speed ramps — keyframe speed and the engine time-remaps video and the
export audio mix (the fast-into-slow-mo reel move)
Camera shake and RGB-split/chromatic aberration , both animatable
17 transitions: fades, slides, wipes (4 directions), zoom, iris, spin, blur,
whip-pan, glitch , pop
Kinetic captions: typewriter, word-pop, word-slide, karaoke, letter-pop ,
wave , bounce , shake
Neon glow for that TikTok caption look
Font editor: system fonts, drop-in custom fonts ( library/fonts/ ), and any
Google Font by name — loaded automatically
Gradient fills, outline, background pills, letter-spacing, line-height,
weights, italic, uppercase, alignment, soft shadows
A first-class svg clip kind: CSS- @keyframes -animated SVGs render
frame-accurately in preview and export (the compositor freezes the
animation at any time). Agents can author their own vector overlays —
lower-thirds, confetti, sparkles — as plain .svg files. Starters included.
Give it a reference edit (a reel you like) and get back an edit blueprint :
shot boundaries, music beats + BPM, a loudness curve, per-shot energy, the
drop — plus the reference's music track extracted into your media, ready
to rebuild the same idea with your own footage. Zero extra dependencies
(ffmpeg does the decoding; onset/tempo detection is plain Node).
node analyze.js ref.mp4 , POST /api/analyze , or the
fablecut_analyze_reference MCP tool.
library/ folders surface as tabs in the UI: Elements (overlay art),
Sound FX , SVG — drop files in, the open editor refreshes live
Fast export: browser renders every frame + an offline audio mix, ffmpeg
encodes a frame-accurate CRF-18 MP4 (keeps rendering if you switch tabs)
Realtime MediaRecorder fallback when ffmpeg isn't available
git clone https://github.com/ronak-create/FableCut.git
cd FableCut
node server.js # → http://localhost:7777
Requirements: Node 18+ and a Chromium-based browser. ffmpeg on PATH is
optional but recommended (fast export + upload remuxing). AI background
removal fetches its model from a CDN on first use.
Drop media into the window (or ./media/ ), drag clips onto the timeline, edit,
export.
Everything an agent needs is in CLAUDE.md — the complete
schema, semantics and recipes. Point any capable model at that file and it can
operate the editor end to end.
Three equivalent control surfaces:
MCP (best for Claude Code / Claude Desktop) — register the bundled
zero-dependency MCP server once:
claude mcp add -s user fablecut -- node " <path-to>/fablecut/mcp-server.js "
Tools: fablecut_status (auto-starts the editor), fablecut_docs ,
fablecut_get_project , fablecut_set_project , fablecut_patch_project ,
fablecut_import_media , fablecut_analyze_reference .
The surface is token-efficient by design : agents patch the timeline with
small ops ( fablecut_patch_project ) instead of round-tripping the whole
document, read a compact one-line-per-clip summary
( fablecut_get_project {compact:true} ), and fetch only the manual sections
they need ( fablecut_docs {section:"props"} ).
The file — read project.json , modify, bump revision , write. The UI
live-reloads.
REST — GET/PUT /api/project , POST /api/upload , GET /api/library ,
SSE at /api/events . See CLAUDE.md for the full list.
Example: ask Claude Code "cut these six clips to the beat markers, add a
teal-orange grade, put a word-pop caption on top and a whoosh on every cut" —
and watch the timeline rebuild itself.
Or hand it a reference: "here's a reel I like — analyze it and remake it with
my clips, same music" . The agent calls fablecut_analyze_reference , gets the
blueprint (cuts, beats, BPM, energy, drop, extracted music), and rebuilds the
structure shot-for-shot with your footage.
Conflict-safe concurrent editing : the UI, the MCP tools, and direct
project.json writes all agree on a revision counter. If you edit a clip in
the UI while an agent is mid-task, the agent's next write is rejected (409 from
the REST API / a conflict error from fablecut_set_project ) instead of
silently overwriting your change. The UI similarly detects when an agent write
supersedes a not-yet-saved local tweak and tells you with a toast instead of
dropping it silently.
server.js zero-dependency HTTP server: static hosting, REST API, SSE,
ffmpeg export pipeline
app.js the editor: timeline UI, compositor, keyframes, text engine,
SVG rasterizer, chroma key, exporters
index.html single-page UI
style.css dark editor theme
mcp-server.js stdio MCP server exposing the editor to AI agents
analyze.js reference-video analyzer: shots, beats/BPM, energy, drop,
music extraction (module + CLI)
CLAUDE.md the agent manual (schema + recipes) — also served by fablecut_docs
project.json your timeline (created on first run; gitignored)
media/ project footage (gitignored)
analysis/ cached edit blueprints from /api/analyze (gitignored)
library/ default assets: elements/ sfx/ svg/ fonts/
exports/ finished renders (gitignored)
Authoring animated SVG overlays
SVGs animate with plain CSS @keyframes . One convention: never hardcode
animation-delay — set --d: 0.4s instead, and the compositor drives time by
pausing all animations and rebasing their delays. Full rules + a skeleton in
CLAUDE.md ; working
examples in library/svg/ .
The repo ships with 20 Google Fonts ( library/fonts/ , OFL — see
LICENSES.md there) and a set of self-authored SVG overlays and animated
elements ( library/elements/ , library/svg/ , MIT like the rest of the repo).
library/sfx/ is yours to fill (gitignored): sound-effect sites typically
don't allow redistributing their files in a public repo, so FableCut doesn't —
library/sfx/README.md lists good free sources.
Export runs in the browser because the compositor is the browser; agents
ask you to click Export (or render directly with ffmpeg from media/ ).
Zero-dependency browser video editor that AI agents can drive — JSON timeline, MCP + REST, live-reloading UI
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
