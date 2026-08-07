---
source: "https://github.com/ZiadAbdelkarim/beat-synced-edit"
hn_url: "https://news.ycombinator.com/item?id=49208441"
title: "Claude Code skill allowing orchestration of edits beat-synced to songs"
article_title: "GitHub - ZiadAbdelkarim/beat-synced-edit: Automatic beat-synced video editing: feed it a song and raw footage, it analyzes beats, energy, and scenes, then cuts a ready-to-post edit. Pure Python/ffmpeg CLI, plus a bundled Claude Code skill for natural-language editing. · GitHub"
author: "iwasastreamer"
captured_at: "2026-08-07T11:39:11Z"
capture_tool: "hn-digest"
hn_id: 49208441
score: 1
comments: 0
posted_at: "2026-08-07T10:49:38Z"
tags:
  - hacker-news
  - translated
---

# Claude Code skill allowing orchestration of edits beat-synced to songs

- HN: [49208441](https://news.ycombinator.com/item?id=49208441)
- Source: [github.com](https://github.com/ZiadAbdelkarim/beat-synced-edit)
- Score: 1
- Comments: 0
- Posted: 2026-08-07T10:49:38Z

## Translation

タイトル: 曲にビート同期した編集のオーケストレーションを可能にするクロード コード スキル
記事のタイトル: GitHub - ZiadAbdelkarim/beat-synced-edit: 自動ビート同期ビデオ編集: 曲と生の映像をフィードすると、ビート、エネルギー、シーンを分析し、すぐに投稿できる編集を切り出します。 Pure Python/ffmpeg CLI、および自然言語編集用のバンドルされた Claude Code スキル。 · GitHub
説明: 自動ビート同期ビデオ編集: 曲と生の映像をフィードすると、ビート、エネルギー、シーンを分析し、すぐに投稿できる編集を切り出します。 Pure Python/ffmpeg CLI、および自然言語編集用のバンドルされた Claude Code スキル。 - ZiadAbdelkarim/beat-synced-edit

記事本文:
GitHub - ZiadAbdelkarim/beat-synced-edit: 自動ビート同期ビデオ編集: 曲と生の映像をフィードすると、ビート、エネルギー、シーンを分析し、すぐに投稿できる編集を切り出します。 Pure Python/ffmpeg CLI、および自然言語編集用のバンドルされた Claude Code スキル。 · GitHub
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
ジアド・アブデルカリム
/
ビート同期編集
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
28 コミット 28

コミット .claude/ スキル/ ビート同期編集 .claude/ スキル/ ビート同期編集 サンプル サンプル .gitignore .gitignore ライセンス ライセンス README.md README.md Beat_map.py Beat_map.py Clip_tag.py Clip_tag.py overunder_stack.py overunder_stack.py plan_edit.py plan_edit.py render_edit.py render_edit.py要件.txt 要件.txt 垂直スタイル.py 垂直スタイル.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ビート同期編集 — クロード コード オートメーション
自動ビート同期ビデオ編集。
曲と生の映像をフィードします。ビートを分析し、エネルギーをマッピングし、すべてのシーンにタグを付けて、すぐに投稿できる編集をカットします。つまり、手動編集は必要ありません。
プレビュー_フラッシュ-モンタージュ-1x1.mp4
1:1 モンタージュ — ホワイトフラッシュ トランジション + パンチイン、バンドルされたスキルによって構築
プレビュー_シーライフ-ビート-編集-9x16.mp4
9:16 ビート同期編集 — グリーンストロンググレード、ストック映像 + 私自身のシングル
両方のデモのフル品質バージョンは、examples/ にあります。
1. 純粋な CLI — 決定論的、AI なし
4 つのステージを自分で実行します。毎回同じ入力、同じ編集。
2. クロード コード スキル — 会話
リポジトリには、 .claude/skills/beat-sync-edit/SKILL.md にスキルが同梱されています。 Claude Code でこのフォルダーを開き、必要な編集を記述します。
「映像をコーラス部分までカットし、一滴ごとに白いフラッシュを出し、ショットが呼吸するようにカットアウトの間隔をあけて、9分16秒にしました。」
クロードは同じパイプラインを実行し、ビート精度のタイムスタンプで ffmpeg エフェクト (ホワイト フラッシュ、ストレッチ パンチ、色相シフト、RGB 分割、スピード ランプ、シェイク、カラー グレード) をレイヤー化します。完全なエフェクト クックブックはスキル ファイル内にあります。
会話では次のように指示することもできます。
アスペクト比 — 垂直 9:16、正方形 1:1、上下に積み重ね、または 16:9 のまま
カット密度 — 「ビートごとにカット」 vs 「間隔をあけてショットを呼吸させる」
参照による色相/色合い — マシン上の他のビデオをクロードに向けます

(「このようにグレード付けします」) リファレンスからフレームをサンプリングし、一致するカラー グレードを構築します。 Claude-in-Chrome 拡張機能を使用すると、ウェブ上のビデオの外観を研究することもできます
Song.mp3 ──► Beat_map.py ──► ビート、エネルギー カーブ、山/谷 (JSON)
Clips.mp4 ──► Clip_tag.py ──► 動き/エネルギー/明るさでタグ付けされたシーン (JSON)
両方 ──► plan_edit.py ──► 編集決定リスト: どのクリップがどのビートにヒットするか
EDL ──► render_edit.py ──► 最終 MP4 (抽出 → 連結 → オーディオの多重化)
ステージ
何をするのか
ビートマップ.py
librosa オーディオ分析 — ビートのタイムスタンプ、テンポ、時間の経過に伴うエネルギー、山/谷、最高のハイライト セグメント
クリップタグ.py
PySceneDetect シーン検出 + クリップごとのモーション/エネルギー/明るさのスコアリング。 --thumbs はラベル付きのコンタクト シートをエクスポートします
plan_edit.py
マッチャー — 繰り返し防止、ソース間隔、およびブラックフレーム フィルタリングを使用して、高エネルギーのクリップはエネルギーのピークに、穏やかなクリップはエネルギーの谷に着地します。
render_edit.py
ffmpeg アセンブリ。 --overunder はスタックされた 960×1080 バリアントを出力します
垂直スタイル.py
TikTok / リール / ショート用の様式化されたスクイーズ + グレードを使用して 16:9 の出力を 9:16 にフィットさせます
クイックスタート
pip install -rrequirements.txt # ffmpeg は PATH 上にある必要があります (brew install ffmpeg)
python3 Beat_map.py Song.mp3
python3 Clip_tag.py 映像.mp4 --thumbs
python3 plan_edit.py Song_beatmap.json 映像_クリップ.json --html
python3 render_edit.py フッテージ_edl.json -a Song.mp3 -v フッテージ.mp4
カットをコントロールする
すべてのクリエイティブ コントロールは plan_edit.py にあります。
密度を比較するときに便利な規則: 両方をレンダリングし、接尾辞として _beat-full および _beat-thinned を付けます。
pip install -rrequirements.txt — librosa、PySceneDetect、OpenCV、NumPy
自動ビート同期ビデオ編集: 曲と生の映像をフィードすると、ビート、エネルギー、シーンを分析し、領域をカットします。

dy-to-post編集。 Pure Python/ffmpeg CLI、および自然言語編集用のバンドルされた Claude Code スキル。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Automatic beat-synced video editing: feed it a song and raw footage, it analyzes beats, energy, and scenes, then cuts a ready-to-post edit. Pure Python/ffmpeg CLI, plus a bundled Claude Code skill for natural-language editing. - ZiadAbdelkarim/beat-synced-edit

GitHub - ZiadAbdelkarim/beat-synced-edit: Automatic beat-synced video editing: feed it a song and raw footage, it analyzes beats, energy, and scenes, then cuts a ready-to-post edit. Pure Python/ffmpeg CLI, plus a bundled Claude Code skill for natural-language editing. · GitHub
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
ZiadAbdelkarim
/
beat-synced-edit
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
28 Commits 28 Commits .claude/ skills/ beat-sync-edit .claude/ skills/ beat-sync-edit examples examples .gitignore .gitignore LICENSE LICENSE README.md README.md beat_map.py beat_map.py clip_tag.py clip_tag.py overunder_stack.py overunder_stack.py plan_edit.py plan_edit.py render_edit.py render_edit.py requirements.txt requirements.txt vertical_style.py vertical_style.py View all files Repository files navigation
Beat-Synced Edit — Claude Code Automation
Automatic beat-synced video editing.
Feed it a song and raw footage. It analyzes the beats, maps the energy, tags every scene, and cuts a ready-to-post edit — zero manual editing.
preview_flash-montage-1x1.mp4
1:1 montage — white-flash transitions + punch-ins, built via the bundled skill
preview_sealife-beat-edit-9x16.mp4
9:16 beat-synced edit — green-strong grade, stock footage + my own single
Full-quality versions of both demos live in examples/ .
1. Pure CLI — deterministic, no AI
Run the four stages yourself. Same inputs, same edit, every time.
2. Claude Code skill — conversational
The repo ships a skill at .claude/skills/beat-sync-edit/SKILL.md . Open this folder in Claude Code and describe the edit you want:
"cut my footage to the chorus, white flash on every drop, space the cuts out so the shots breathe, then make it 9:16"
Claude runs the same pipeline and layers ffmpeg effects at beat-accurate timestamps — white flashes, stretch punches, hue shifts, RGB split, speed ramps, shakes, color grades. The full effects cookbook lives in the skill file.
In conversation you can also dictate:
Aspect ratio — 9:16 vertical, 1:1 square, stacked over-under, or leave it 16:9
Cut density — "a cut on every beat" vs "spaced out, let the shots breathe"
Hue / tint by reference — point Claude at any other video on your machine ("grade it like this one") and it will sample frames from the reference and build a matching color grade; with the Claude-in-Chrome extension it can even study a look from a video on the web
song.mp3 ──► beat_map.py ──► beats, energy curve, peaks/valleys (JSON)
clips.mp4 ──► clip_tag.py ──► scenes tagged by motion/energy/brightness (JSON)
both ──► plan_edit.py ──► edit decision list: which clip hits which beat
EDL ──► render_edit.py ──► final MP4 (extract → concat → mux audio)
Stage
What it does
beat_map.py
librosa audio analysis — beat timestamps, tempo, energy over time, peaks/valleys, best highlight segments
clip_tag.py
PySceneDetect scene detection + per-clip motion/energy/brightness scoring; --thumbs exports a labeled contact sheet
plan_edit.py
the matcher — high-energy clips land on energy peaks, calm clips on valleys, with anti-repetition, source spacing, and black-frame filtering
render_edit.py
ffmpeg assembly; --overunder emits a stacked 960×1080 variant
vertical_style.py
fits 16:9 output into 9:16 with a stylized squeeze + grade for TikTok / Reels / Shorts
Quickstart
pip install -r requirements.txt # ffmpeg must be on PATH (brew install ffmpeg)
python3 beat_map.py song.mp3
python3 clip_tag.py footage.mp4 --thumbs
python3 plan_edit.py song_beatmap.json footage_clips.json --html
python3 render_edit.py footage_edl.json -a song.mp3 -v footage.mp4
Controlling the cut
All creative control lives in plan_edit.py :
A useful convention when comparing densities: render both and suffix them _beat-full and _beat-thinned .
pip install -r requirements.txt — librosa, PySceneDetect, OpenCV, NumPy
Automatic beat-synced video editing: feed it a song and raw footage, it analyzes beats, energy, and scenes, then cuts a ready-to-post edit. Pure Python/ffmpeg CLI, plus a bundled Claude Code skill for natural-language editing.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
