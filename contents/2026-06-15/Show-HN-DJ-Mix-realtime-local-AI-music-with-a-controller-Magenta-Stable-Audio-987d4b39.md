---
source: "https://github.com/brxs/slipmate"
hn_url: "https://news.ycombinator.com/item?id=48539852"
title: "Show HN: DJ Mix realtime local AI music with a controller; Magenta, Stable Audio"
article_title: "GitHub - brxs/slipmate: An AI DJ application · GitHub"
author: "ttoinou"
captured_at: "2026-06-15T14:17:12Z"
capture_tool: "hn-digest"
hn_id: 48539852
score: 3
comments: 0
posted_at: "2026-06-15T11:40:35Z"
tags:
  - hacker-news
  - translated
---

# Show HN: DJ Mix realtime local AI music with a controller; Magenta, Stable Audio

- HN: [48539852](https://news.ycombinator.com/item?id=48539852)
- Source: [github.com](https://github.com/brxs/slipmate)
- Score: 3
- Comments: 0
- Posted: 2026-06-15T11:40:35Z

## Translation

タイトル: Show HN: DJ コントローラーを使用してリアルタイムのローカル AI 音楽をミックスします。マゼンタ、安定したオーディオ
記事タイトル: GitHub - brxs/ Slipmate: AI DJ アプリケーション · GitHub
説明: AI DJ アプリケーション。 GitHub でアカウントを作成して、brxs/スリップメイトの開発に貢献してください。
HN テキスト: SlipMate は 2 つのローカル AI 音楽モデルを DJ デッキに変え、テキスト プロンプトで操作し、ヴァイナル クロスフェーダー、3 バンド EQ、フリーズ ループ、Pioneer DDJ-FLX4 のフル コントロールのようにミックスします。すべてが Apple Silicon 上でローカルに実行され、マウスの代わりに実際のハードウェアでプレイすることが最も楽しいです。自分でコードを変更して他のモデルを追加したり、最近の Mac 以外の場所でも動作させることができます。禁止される前に Fable 5 で構築されました。このような関連プロジェクトを他にも知っているので教えてください。

記事本文:
GitHub - brxs/ Slipmate: AI DJ アプリケーション · GitHub
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
brxs
/
スリップメイト
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く Fo

フォルダとファイル
216 コミット 216 コミット .claude .claude backend backend docs docsfrontendfrontend .gitignore .gitignore CLAUDE.md CLAUDE.md README.md README.md justfile justfile すべてのファイルを表示 リポジトリ ファイルのナビゲーション
あなたの生成的な DJ メイト — リアルタイム AI 音楽のための DJ 楽器。
テキスト プロンプトによって操作され、ビニールのようにミックスされた、ローカルで実行される 2 つのモデル デッキ:
3バンドEQ、ワンノブカラーFX、クロスフェーダー、ヘッドフォンキュー、フル機能
パイオニアDDJ-FLX4コントロール。ライブデッキは上で動作します
マゼンタ リアルタイム 2 ;生成された
パッドと完成したトラックは Stable Audio 3 から取得されます。 を参照してください。
docs/ROADMAP.md に到達した経緯と
アーキテクチャの決定については docs/adr/ を参照してください。
Apple Silicon Mac (MLX バックエンド)
モデルの重み付け用に最大 13 GB のディスク (最初のセットアップ時にダウンロード: マゼンタ)
両方のデッキ モデルで最大 4.5 GB、安定したオーディオ 3 を含む最大 8 GB
ミディアムトラックモデル）
Chromium ベースのブラウザ (アプリは Web オーディオ ワークレットと Web MIDI に依存しています。
Chrome に対して開発および検証されています)
オプション: ハードウェア制御用の Pioneer DDJ-FLX4 とそのヘッドフォン ジャック
すべての一般的なタスクは justfile 内に存在します。タスクをリストするために実行するだけです。
バックエンド deps、すべてのモデルの重み (~13 GB)、フロントエンド deps + ビルドをセットアップするだけです
マゼンタ モデルは ~/Documents/Magenta/magenta-rt-v2 に配置されます (次のようにオーバーライドします)
MAGENTA_HOME ): 両方のデッキ モデル、デフォルトの mrt2_small と重い方、
高品質の mrt2_base 、UI でデッキごとに選択可能 — アプリは警告します
組み合わせた選択が RAM に対して厳しいと思われる場合。安定したオーディオ 3 —
生成されたパッドとトラック — ~/Repos/stable-audio-3 にクローン化されます (オーバーライド)
SA3_MLX_HOME を使用します。既存のチェックアウトが再利用されます)、その重みは
事前にウォームアップされているため、リクエストによってダウンロードの費用が発生することはありません。 setup-sa3 を再実行するだけです
その半分は一人です。
ただ走ってください
次に http://127.0.0.1:8000 を開きます — スタイル ターゲットをデッキのパッドに追加し、ヒットします
遊び、

カーソル (またはドット自体) をドラッグしてターゲットをブレンドします。
クラスター化して)、デッキ間でクロスフェーダーを動かします。
ミキサー — デッキごとのボリューム、Hi/Mid/Low EQ、クロスフェーダー、および Record 、
これにより、マスター バスがダウンロード可能な WAV にキャプチャされます。健康行には次のことが表示されます
ストリームバッファ、アンダーラン数、生成速度。
カラー FX — 選択したエフェクトに対するデッキごとに 1 つのノブ: フィルター (バイポーラー)
LPF/HPF)、ダブエコー、スペース、クラッシュ、ノイズ、スイープ。ノブのセンター/ゼロは、
ビット正確なバイパス ( ADR-0008 )。
フリーズ ループ — デッキの最後の小節を 4 つのうちの 1 つにキャプチャします。
スロットをループし、モデルを再操作する間、その瞬間を空中に保持します。
その下に。ループは設計上セッションのみです
( ADR-0009 )。
ビート検出 - 各デッキは、検出された BPM をビートの背後に表示します。
正直ゲート (間違った数字ではなくダッシュ);自信を持って
Dub Echo がビートに同期し、フリーズ キャプチャがクオンタイズされるテンポ
全ビート
( ADR-0010 )。
デッキからデッキへのスタイルのサンプリング — ワンプレスで「
「他のデッキ、今すぐ」をデッキのスタイル パッド上でブレンド可能なターゲットとして表示します。
サンプリングされたターゲットは設計によりセッションのみです
( ADR-0011 )。
クレート — デッキのパッド + カラー FX を名前付きプリセットとして保存し、参照します
FLX4 のロータリーからクレートを取り外し、セットの途中でいずれかのデッキにロードします。
バックアップと共有のために JSON としてエクスポート/インポートします。
マスターハウスキーピング — マスターのリミッター（メーター、
録音すると、すべての電話機が限定された信号を聞きます。その利益
リダクションはミキサーに表示されます）とチャンネルごとの自動ゲイントリム
手動オーバーライドにより、異なるラウドネスのデッキをレベルアップします。
ヘッドフォンキュー — チャンネルのキューを押し、キューミックスノブを動かします
キューとマスターの間でフォンアウトを選択します: 任意の出力デバイス
ブラウザがアクセスできるか、FLX4 自体のヘッドフォン ジャックに接続できます。
USB 経由のバックエンド ( ADR-0007 )。
設定（パッド配置、

ボリューム、クロスフェード) はリロード後も持続します。
ショートカット: A / B はデッキのスタイルターゲット入力にフォーカスし、X はデッキのスタイルターゲット入力にフォーカスします。
クロスフェーダー。
フロントエンド開発の場合: ある端末では dev-backend のみ、別の端末では dev-frontend のみ (Vite dev サーバーは /ws をバックエンドにプロキシします)。
ハードウェア制御（Pioneer DDJ-FLX4）
FLX4 を接続し、「MIDI を接続」をクリックします (Chrome は SysEx で MIDI を要求します。
通常の MIDI も機能します (位置同期はありません)。マップされたコントロール:
再生/一時停止、チャンネルフェーダー、3バンドEQ、クロスフェーダー
チャンネルCUEボタン（ヘッドフォンキュー）とトランスポートCUEボタン
(デッキの準備: 停止したデッキをオフエアで準備し、再生中のデッキを停止します)
SMART CFX ノブ — カラー FX 量。 SHIFT を押したままスタイル パッドをスイープします
代わりに
PAD FX パッド バンク — デッキのエフェクトを選択します (再押しするとオフになります)。
HOT CUEパッ​​ドはスタイルターゲットを選択します。 SAMPLER パッドのフリーズ ループ
(SHIFT + パッドでスロットをクリア)
ロータリー + LOAD ボタンを参照 — クレート プリセットをハイライト表示し、ロードします
デッキ1または2に
ノブとフェーダーの位置は、接続時のハードウェアと LED から同期されます。
アプリの状態をミラーリングします。測定されたバイトマップは次の場所にあります。
docs/midi-ddj-flx4.md 。
テストするだけ — バックエンド pytest + フロントエンド vitest
lint だけ — フォーマットチェック、ruff、eslint、tsc
上記の両方をチェックしてください。 PR が通過しなければならないもの
just verify-stream / just verify-ui — 実行中のサーバーに対する e2e
(UI e2e には Playwright Chromium が一度必要です: npx playwright install chrome
フロントエンドで/ )
just verify-worklets — audio-worklet モジュールのグラフが実際にロードされます
Chromium (自己完結型。jsdom はワークレット コードを実行しません)
ハードウェアの動作は、次のチェックリストに基づいて人間によって検証されます。
docs/ ( m7- 、 m9-m10- 、 m12-ハードウェア-チェックリスト.md )
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
貢献者
読み込み中にエラーが発生しました。このページをリロードしてください。

© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

An AI DJ application. Contribute to brxs/slipmate development by creating an account on GitHub.

SlipMate turns two local AI music models into DJ decks you steer with text prompts and mix like vinyl.crossfader, three-band EQ, freeze loops, and full Pioneer DDJ-FLX4 control. Everything runs locally on Apple Silicon, and playing it on real hardware instead of a mouse is most of the fun. You can change the code yourself to add others models or make it work elsewhere than on a recent mac. Built with Fable 5 before it got banned. Let me know others related projects you know like this !

GitHub - brxs/slipmate: An AI DJ application · GitHub
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
brxs
/
slipmate
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
216 Commits 216 Commits .claude .claude backend backend docs docs frontend frontend .gitignore .gitignore CLAUDE.md CLAUDE.md README.md README.md justfile justfile View all files Repository files navigation
Your generative DJ mate — a DJ instrument for real-time AI music.
Two locally-running model decks, steered by text prompts and mixed like vinyl:
three-band EQ, one-knob Color FX, a crossfader, headphone cue, and full
Pioneer DDJ-FLX4 control. The live decks run on
Magenta RealTime 2 ; generated
pads and finished tracks come from Stable Audio 3. See
docs/ROADMAP.md for how it got here and
docs/adr/ for the architecture decisions.
Apple Silicon Mac (MLX backend)
~13 GB disk for model weights (downloaded on first setup: Magenta
~4.5 GB for both deck models, Stable Audio 3 ~8 GB including the
medium track model)
A Chromium-based browser (the app leans on Web Audio worklets and Web MIDI;
it is developed and verified against Chrome)
Optional: a Pioneer DDJ-FLX4 for hardware control and its headphone jack
All common tasks live in the justfile — run just to list them.
just setup # backend deps, all model weights (~13 GB), frontend deps + build
Magenta models land in ~/Documents/Magenta/magenta-rt-v2 (override with
MAGENTA_HOME ): both deck models, the default mrt2_small and the heavier,
higher-quality mrt2_base , selectable per deck in the UI — the app warns
when the combined selection looks tight for your RAM. Stable Audio 3 —
generated pads and tracks — is cloned to ~/Repos/stable-audio-3 (override
with SA3_MLX_HOME ; an existing checkout is reused) and its weights are
pre-warmed so no request ever pays for a download; just setup-sa3 re-runs
that half alone.
just run
Then open http://127.0.0.1:8000 — add style targets to a deck's pad, hit
play, blend targets by dragging the cursor (or the dots themselves, to
cluster them), and ride the crossfader between decks.
Mixer — per-deck volume and Hi/Mid/Low EQ, crossfader, and Record ,
which captures the master bus to a downloadable WAV. The health row shows
the stream buffer, underrun count, and generation speed.
Color FX — one knob per deck over a chosen effect: Filter (bipolar
LPF/HPF), Dub Echo, Space, Crush, Noise, Sweep. The knob's centre/zero is a
bit-exact bypass ( ADR-0008 ).
Freeze loops — capture the last bars of a deck into one of four
loop slots and hold the moment on air while you re-steer the model
underneath; loops are session-only by design
( ADR-0009 ).
Beat detection — each deck shows its detected BPM behind an
honesty gate (a dash rather than a wrong number); with a confident
tempo the Dub Echo syncs to the beat and freeze captures quantise to
whole beats
( ADR-0010 ).
Deck-to-deck style sampling — one press puts "the sound of the
other deck, right now" on a deck's style pad as a blendable target;
sampled targets are session-only by design
( ADR-0011 ).
Crates — save a deck's pad + Color FX as a named preset, browse
the crate from the FLX4's rotary, and load onto either deck mid-set;
export/import as JSON for backup and sharing.
Master housekeeping — a limiter on the master (the meter, the
recording, and the phones all hear the limited signal; its gain
reduction shows in the mixer) and per-channel auto-gain Trim that
levels decks of different loudness, with a manual override.
Headphone cue — hit a channel's Cue , ride the Cue mix knob
between cue and master, and pick a Phones out : any output device the
browser can reach, or the FLX4's own headphone jack, which is fed by the
backend over USB ( ADR-0007 ).
Settings (pad arrangements, volumes, crossfade) persist across reloads.
Shortcuts: A / B focus a deck's style-target input, X focuses the
crossfader.
For frontend development: just dev-backend in one terminal, just dev-frontend in another (the Vite dev server proxies /ws to the backend).
Hardware control (Pioneer DDJ-FLX4)
Plug in the FLX4 and click Connect MIDI (Chrome asks for MIDI with SysEx;
plain MIDI works too, minus position sync). Mapped controls:
Play/pause, channel faders, three-band EQ, crossfader
Channel CUE buttons (headphone cue) and the transport CUE button
(deck prep: prime a stopped deck off-air, stop a playing one)
SMART CFX knob — Color FX amount; hold SHIFT to sweep the style pad
instead
PAD FX pad bank — select the deck's effect (re-press toggles it off);
HOT CUE pads pick style targets; SAMPLER pads freeze loops
(SHIFT + pad clears a slot)
Browse rotary + LOAD buttons — highlight a crate preset, load it
onto deck 1 or 2
Knob and fader positions sync from the hardware on connect, and the LEDs
mirror app state. The measured byte map lives in
docs/midi-ddj-flx4.md .
just test — backend pytest + frontend vitest
just lint — format check, ruff, eslint, tsc
just check — both of the above; what a PR must pass
just verify-stream / just verify-ui — e2e against a running server
(UI e2e needs Playwright Chromium once: npx playwright install chromium
in frontend/ )
just verify-worklets — the audio-worklet module graph loads in real
Chromium (self-contained; jsdom executes none of the worklet code)
Hardware behaviour is verified by a human against the checklists in
docs/ ( m7- , m9-m10- , m12-hardware-checklist.md )
There was an error while loading. Please reload this page .
0
forks
Report repository
Contributors
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
