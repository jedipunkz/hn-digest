---
source: "https://github.com/ifoxhz/pigame"
hn_url: "https://news.ycombinator.com/item?id=49382504"
title: "Show HN: Pigame – pi+LLM plays black-box browser games via observe/move tools"
article_title: "GitHub - ifoxhz/pigame: This is a Pi extension that uses an LLM to drive the game, starting with a simple 2D game. · GitHub"
image: "https://opengraph.githubassets.com/171011cbb16201251e58c60a646f2a3cb3a5e1c52866e711bdaceaf442c89e69/ifoxhz/pigame"
author: "ifoxhz"
captured_at: "2026-08-21T02:18:23Z"
capture_tool: "hn-digest"
hn_id: 49382504
score: 1
comments: 0
posted_at: "2026-08-21T01:20:26Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Pigame – pi+LLM plays black-box browser games via observe/move tools

- HN: [49382504](https://news.ycombinator.com/item?id=49382504)
- Source: [github.com](https://github.com/ifoxhz/pigame)
- Score: 1
- Comments: 0
- Posted: 2026-08-21T01:20:26Z

## Translation

タイトル: Show HN: Pigame – pi+LLM は観察/移動ツールを介してブラックボックス ブラウザ ゲームをプレイします
記事のタイトル: GitHub - ifoxhz/pigame: これは、LLM を使用してゲームを駆動する Pi 拡張機能で、単純な 2D ゲームから始まります。 · GitHub
説明: これは、LLM を使用してゲームを駆動する Pi 拡張機能で、単純な 2D ゲームから始まります。 - ifoxhz/ピゲーム
HN テキスト: LLM をニューロモーフィック チップに圧縮して推論遅延を短縮すれば、低コストの自動運転を実現するのは比較的簡単でしょうか?

記事本文:
GitHub - ifoxhz/pigame: これは、LLM を使用してゲームを駆動する Pi 拡張機能で、単純な 2D ゲームから始まります。 · GitHub
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
ifoxhz
/
ピゲーム
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
3 コミット 3 コミット フォルダーとファイル
.pi .pi docs docs pigame pigame .gitignore .gitignore README.md README.md README.zh-CN.md README.zh-CN.md package.json package.json すべてのファイルを表示 リポジトリ f

ファイルナビゲーション
Pi + LLM 上の一般的なゲーム エージェント (DeepSeek など)。
デザインコードネーム: GameMind 。最初のデモ ゲーム: ネオン スネーク。
リモート: git@github.com:ifoxhz/pigame.git
ようこそ。あらゆる種類の貢献やアイデアを歓迎します。Pi エージェントをもっと楽しくしましょう。
Pi + LLM = 脳 (戦略 / いつ停止するか)
ピゲーム = 足場のみ
道具＝目・手・足
レイヤー
役割
円周率 + LLM
何をするかを決めてください。目標が達成されたときに判断する
ゲームオブザーブ
目 — 画面を参照 → AgentState
ゲーム移動 / ゲーム待機
手と足 — 行動する / 待つ
npm run stop:* / /game Eat
テスト スキャフォールディングのみ - CI のハードコードされたループ。エージェントパスではありません
game_eat_food / game_survive ツールはありません。モデルは自身を観察→移動するというループを行います。
ツール
注意事項
Node.js
推奨 20 以上 ( --experimental-strip-types )
npm
ノードが付属しています
Pi CLI
エージェント パスに必須 (PATH の pi)
円周率のLLM
例: Pi 設定で構成された DeepSeek
WSLg / ディスプレイ
可視の Chromium のみ ( HEADED=1 )
セットアップ
git clone git@github.com:ifoxhz/pigame.git
cd pigame # またはローカルフォルダー名
2.pigame deps + Chromiumをインストールする
npm install --prefix pigame
npx --prefix pigame playwright インストール クロム
デフォルトのゲーム (追加のクローンなし): pigame/games/snake.html — エージェント モードを備えたベンダーの Neon Snake (以下を参照)。
このリポジトリを複製すると、パッチが適用されたゲームが自動的に取得されます。アダプターはそれを /game connect にロードします。
3. オプション — 上流の Mini Game Studio
比較のために元の連続ループ ゲームが必要な場合のみ:
git clone git@github.com:Digiman/mini-game-studio-ai-gen.git
# HTTPS: https://github.com/Digiman/mini-game-studio-ai-gen.git
SNAKE_HTML=./mini-game-studio-ai-gen/games/snake.html HEADED=1 pi -e ./.pi/extensions/pigame.ts
ライブサイト（参考）：https://digiman.github.io/mini-game-studio-ai-gen/
個別のコンパイル手順はありません。

または喫煙: ノード ストリップ タイプは TypeScript を直接実行します。
ルート package.json は、スクリプトを pigame/ にプロキシするだけです。
git clone git@github.com:ifoxhz/pigame.git
cd pigame # またはローカルフォルダー名
npm install --prefix pigame
npx --prefix pigame playwright インストール クロム
B — Pi + LLM (プライマリ パス)
cd /パス/へ/pigame
HEADED=1 pi -e ./.pi/extensions/pigame.ts
Pi TUI 内 (順番に):
/ゲーム接続
/スキル:食べ物を食べる
または、スキルの代わりに次のプロンプトを貼り付けます。
game_observe、game_move のみを使用し、オプションで game_wait も使用します。
groundTruth グリッド座標を優先します。一つの食べ物を食べる：観察→移動→繰り返し。
スコアが上がったら停止し、スコアの前後を報告します。
180°反転させないでください。
Playwright Chromium ウィンドウを見てください (スネークが動きの合間にフリーズします)。
完了したら:
/ゲームの切断
ライブセッションでコードを変更した後: /reload 、次に /game 切断 → /game 再度接続。
C — オプション: 非 LLM スモーク (足場)
# アダプター + 知覚 + I/O のみ (Pi なし / LLM なし)
npm 煙を実行: 食べる
npm run 煙:食べる:頭
D — Pi npm 拡張機能としてパック (その他の場合)
dist/pi-pigame/ — インストール可能なディレクトリ
dist/pigame-0.1.0.tgz — npm tarball
pi インストール /path/to/pigame/dist/pi-pigame
# または
pi インストール /path/to/pigame/dist/pigame-0.1.0.tgz
# npm に公開した後:
# npm 公開 ./dist/pigame-0.1.0.tgz
# pi インストール npm:pigame
git clone (パッキングなし) からは、次のこともできます。
pi インストール ./pigame
次に、Pi を再起動するか /reload します。初回: npx playwright で chromium をインストールします (必要に応じて、インストールされたパッケージ ディレクトリから)。
エージェントモード Snake (Pi + LLM レイテンシの場合)
問題: 上流の Snake が勝手に動きます。 LLM は長い観察→思考→移動の連鎖の中にありますが、スネークはすでに移動しているため、決定は古くなり、プレイはぎこちなく感じられます。
出荷内容: pigame/games/snake.html ( pigame/games/README.md を参照)
特別なスターはいない

エージェント モードには tup フラグが必要です。接続はデフォルトでベンダー提供のファイルを使用します。
これが意図した製品フローです。Pi を開きます。 LLM はツールを使用します。
ステップ A — 拡張子を付けて Pi を起動する
cd /パス/へ/pigame
HEADED=1 pi -e ./.pi/extensions/pigame.ts
# または: pi (プロジェクトが信頼されている場合は .pi/extensions/pigame.ts を自動ロードします)
コード変更後: /reload 。
次のような内容が表示されるはずです: GameMind 準備完了 — LLM ツール: game_observe、game_move、game_wait 。
ステップ B — ゲームを接続する (ヒューマン / スラッシュ コマンド)
/ゲーム接続
Playwright は Chromium を開き、pigame/games/snake.html (エージェント モード) をロードしてから開始します。
世界は各 game_move までフリーズします。
自分のブラウザでsnake.htmlを開かないでください。制御されるのはPlaywrightウィンドウのみです。
オプション: /game は AgentState の健全性をチェックするために観察します。
ステップ C — LLM に駆動させる (プロンプト)
例 (DeepSeek / Pi の任意のモデル):
game_observe 、 game_move のみを使用し、オプションで game_wait を使用します。
1つの食べ物を食べる：観察→方向を選択→移動→繰り返し。
スコアが上がったら停止し、スコアの前後を報告します。
前の動きから 180 度反転しないでください。
または、スキルをロードします: /skill:eat-food 。
ステップ D — モデルが行うべきこと
game_observe → 選手 / 食事 / スコアの読み取り
game_move → 食への一歩
ゲーム観察 → …
…モデルが目標が完了したと判断するまで…
スラッシュ コマンド (LLM ツールではなくセッション スキャフォールディング)
コマンド
アクション
/ゲーム
ステータス + ヘルプ
/ゲーム接続
劇作家ネオン・スネークを始める
/ゲーム観察
AgentState の圧縮 (手動チェック)
/ゲームの状態
完全な AgentState JSON
/ゲームリセット
ラウンドを再開する
/ゲームの切断
ブラウザを閉じる
/ゲームモック
モックアダプターに切り替える
/ゲームを食べる
テストのスキャフォールディング — スクリプトによる自動実行
/ゲームは生き残る
テスト足場 — 死ぬまで台本に沿ったプレイ
LLM ツール
ツール
役割
ゲームオブザーブ
目
ゲームムーブ
手/足（上/下）

n / 左 / 右）
ゲーム待機
待ってください
スキャフォールディングのテスト ( npm runsmoke:* )
これらのコマンドは Pi または LLM を使用しません。固定ポリシー ( EatFoodLoop / SurvivalUntilEnd ) を実行して、アダプター + 認識 + I/O を検証します。エージェントの設計としてではなく、CI とデバッグに使用してください。
# 食べ物を 1 つ食べる (スコア 0 → 10)、首なし
npm 煙を実行: 食べる
# 可視の Chromium (WSLg) と同じ
npm run 煙:食べる:頭
# 死ぬまでロングラン (またはスコアで停止)
npm run スモーク:survive:headed
MAX_SCORE=50 npm 煙を実行:生存:頭
コマンド
チェック内容 (足場)
npm 煙を実行する
モックアダプター/AgentStateパイプライン
npm run スモーク:スネーク
Playwright を開く + スタート + キー + スクリーンショット
npm run スモーク:知覚
CV フィクスチャ + ライブ知覚
npm 煙を実行: 食べる
スコアまでのスクリプトステア↑
npm 煙を実行:生き残る
ゲームオーバー (または MAX_SCORE ) までスクリプト化されたループ
npm run スモーク:*:headed
HEADED=1 と同じ
環境変数（煙/頭）
変数
デフォルト
意味
頭がついた
設定を解除する
1 / true → Chromium を表示
SNAKE_HTML
販売代理店のスネーク
パスのオーバーライド
MOVE_STEPS
1
移動ごとのステップ数 (エージェント モード)
MAX_SCORE
なし
スコア ≥ 値の場合、生存は停止します
MAX_STEPS
50000
生き残るステップキャップ
STEP_MS
200
煙の移動間の視覚的な遅延
HOLD_MS
5000 – 8000
ヘディングランの後はウィンドウを開いたままにしてください
ネオン・スネークには明確な要素はなく、死だけが存在します。 MAX_SCORE はテストのための人為的な停止です。
。
§── README.md / README.zh-CN.md
§── package.json # スクリプトプロキシ
§── docs/ # デザインとスキーマ
§── pigame/ # メインパッケージ
│ §── games/snake.html # DEFAULT エージェントモード Neon Snake (ベンダー)
│ §── ゲーム/README.md
│ §── extensions/pigame.ts
│ §── スキル/
│ §── scripts/ #スモーク:* (スキャフォールディングのテスト)
│ └── src/
§── .pi/extensions/pigame.ts # Pi エントリ
└── mini-game-studio-ai-gen/ # 上流側はオプション

m クローン (SNAKE_HTML=…)
設計ドキュメント: docs/README.md 。
問題
修正
pigame/games/snake.html がありません
最新のピゲームをプルします — ファイルはリポジトリ内にあります
上流の連続ループが必要
SNAKE_HTML=./mini-game-studio-ai-gen/games/snake.html (最初にデジマンのクローンを作成します)
LLM が考えている間、Snake は自動的に実行されます。
エージェント モード HTML を使用していません。デフォルトのパス / SNAKE_HTML を確認してください
ルートにある ENOENT package.json
root スクリプトまたは cd pigame を使用する
クロムウィンドウなし
HEADED=1 または *:headed ; WSLg をチェック / echo $DISPLAY
自分でSnakeを開いても何も起こりません
Playwright のウィンドウのみが制御されます
Pi 拡張機能がロードされていません
pi -e ./.pi/extensions/pigame.ts then /reload
謝辞
オリジナルの Neon Snake と HTML ゲーム コレクションを提供してくれた Mini Game Studio ( Digiman/mini-game-studio-ai-gen ) に感謝します。
Pi + LLM エージェント モード テスト (一時停止した自動ループ + シングルステップ移動) 用に、パッチを適用したコピーを Pigame/games/ で販売しています。
進行中の作業 (MVP: Pi + LLM は観察/移動を介して Snake を駆動します)。
フェーズノート: docs/implementation.md 。
これは、LLM を使用してゲームを駆動する Pi 拡張機能で、単純な 2D ゲームから始まります。
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

This is a Pi extension that uses an LLM to drive the game, starting with a simple 2D game. - ifoxhz/pigame

If we compress LLMs into neuromorphic chips to reduce inference latency, would it be relatively easy to achieve low-cost autonomous driving?

GitHub - ifoxhz/pigame: This is a Pi extension that uses an LLM to drive the game, starting with a simple 2D game. · GitHub
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
ifoxhz
/
pigame
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
3 Commits 3 Commits Folders and files
.pi .pi docs docs pigame pigame .gitignore .gitignore README.md README.md README.zh-CN.md README.zh-CN.md package.json package.json View all files Repository files navigation
General game agent on Pi + LLM (e.g. DeepSeek).
Design codename: GameMind . First demo game: Neon Snake .
Remote: git@github.com:ifoxhz/pigame.git
Welcome. Contributions of any kind and any ideas are welcome — let’s make the Pi agent more fun.
Pi + LLM = brain (strategy / when to stop)
pigame = scaffolding only
tools = eyes / hands / feet
Layer
Role
Pi + LLM
Decide what to do; judge when the goal is done
game_observe
Eyes — see the screen → AgentState
game_move / game_wait
Hands & feet — act / wait
npm run smoke:* / /game eat
Test scaffolding only — hardcoded loops for CI; not the agent path
There is no game_eat_food / game_survive tool. The model loops observe → move itself.
Tool
Notes
Node.js
≥ 20 recommended ( --experimental-strip-types )
npm
Comes with Node
Pi CLI
Required for the agent path ( pi on PATH)
LLM in Pi
e.g. DeepSeek configured in your Pi settings
WSLg / display
Only for visible Chromium ( HEADED=1 )
Setup
git clone git@github.com:ifoxhz/pigame.git
cd pigame # or your local folder name
2. Install pigame deps + Chromium
npm install --prefix pigame
npx --prefix pigame playwright install chromium
Default game (no extra clone): pigame/games/snake.html — a vendored Neon Snake with agent mode (see below).
Anyone who clones this repo gets the patched game automatically; Adapter loads it on /game connect .
3. Optional — upstream Mini Game Studio
Only if you want the original continuous-loop game for comparison:
git clone git@github.com:Digiman/mini-game-studio-ai-gen.git
# HTTPS: https://github.com/Digiman/mini-game-studio-ai-gen.git
SNAKE_HTML=./mini-game-studio-ai-gen/games/snake.html HEADED=1 pi -e ./.pi/extensions/pigame.ts
Live site (reference): https://digiman.github.io/mini-game-studio-ai-gen/
No separate compile step for smokes: Node strip-types runs TypeScript directly.
Root package.json only proxies scripts into pigame/ .
git clone git@github.com:ifoxhz/pigame.git
cd pigame # or your local folder name
npm install --prefix pigame
npx --prefix pigame playwright install chromium
B — Pi + LLM (primary path)
cd /path/to/pigame
HEADED=1 pi -e ./.pi/extensions/pigame.ts
In the Pi TUI (in order):
/game connect
/skill:eat-food
Or paste this prompt instead of the skill:
Use only game_observe, game_move, and optionally game_wait.
Prefer groundTruth grid coords. Eat one food: observe → move → repeat.
When score increases, stop and report before/after score.
Do not reverse 180°.
Watch the Playwright Chromium window (snake freezes between moves).
When done:
/game disconnect
After code changes in a live session: /reload , then /game disconnect → /game connect again.
C — Optional: non-LLM smoke (scaffolding)
# Adapter + perception + I/O only (no Pi / no LLM)
npm run smoke:eat
npm run smoke:eat:headed
D — Pack as a Pi npm extension (for others)
dist/pi-pigame/ — installable directory
dist/pigame-0.1.0.tgz — npm tarball
pi install /path/to/pigame/dist/pi-pigame
# or
pi install /path/to/pigame/dist/pigame-0.1.0.tgz
# after you publish to npm:
# npm publish ./dist/pigame-0.1.0.tgz
# pi install npm:pigame
From a git clone (without packing), they can also:
pi install ./pigame
Then restart Pi or /reload . First time: npx playwright install chromium (from the installed package dir if needed).
Agent-mode Snake (for Pi + LLM latency)
Problem: Upstream Snake ticks on its own. While the LLM is in a long observe → think → move chain, the snake has already moved, so decisions are stale and play feels clumsy.
What we ship: pigame/games/snake.html (see pigame/games/README.md )
No special startup flag is required for agent mode — connect uses the vendored file by default.
This is the intended product flow: you open Pi; the LLM uses tools .
Step A — Start Pi with the extension
cd /path/to/pigame
HEADED=1 pi -e ./.pi/extensions/pigame.ts
# or: pi (auto-loads .pi/extensions/pigame.ts if the project is trusted)
After code changes: /reload .
You should see something like: GameMind ready — LLM tools: game_observe, game_move, game_wait .
Step B — Connect the game (human / slash command)
/game connect
Playwright opens Chromium and loads pigame/games/snake.html (agent-mode), then Start.
World is frozen until each game_move .
Do not open snake.html in your own browser — only the Playwright window is controlled.
Optional: /game observe to sanity-check AgentState.
Step C — Let the LLM drive (prompt)
Example (DeepSeek / any model in Pi):
Use only game_observe , game_move , and optionally game_wait .
Eat one food: observe → choose a direction → move → repeat.
When score increases, stop and report before/after score.
Do not reverse 180° from the previous move.
Or load the skill: /skill:eat-food .
Step D — What the model should do
game_observe → read player / food / score
game_move → one step toward food
game_observe → …
… until the model decides the goal is done …
Slash commands (session scaffolding — not LLM tools)
Command
Action
/game
Status + help
/game connect
Start Playwright Neon Snake
/game observe
Compact AgentState (manual check)
/game state
Full AgentState JSON
/game reset
Restart round
/game disconnect
Close browser
/game mock
Switch to mock adapter
/game eat
Test scaffolding — scripted auto-eat
/game survive
Test scaffolding — scripted play until death
LLM tools
Tool
Role
game_observe
Eyes
game_move
Hands / feet ( up / down / left / right )
game_wait
Wait
Test scaffolding ( npm run smoke:* )
These commands do not use Pi or an LLM. They run fixed policies ( eatFoodLoop / surviveUntilEnd ) to verify Adapter + Perception + I/O. Use them for CI and debugging — not as the agent design.
# Eat one food (score 0 → 10), headless
npm run smoke:eat
# Same with visible Chromium (WSLg)
npm run smoke:eat:headed
# Long run until death (or stop at a score)
npm run smoke:survive:headed
MAX_SCORE=50 npm run smoke:survive:headed
Command
What it checks (scaffolding)
npm run smoke
Mock adapter / AgentState pipeline
npm run smoke:snake
Playwright open + Start + key + screenshot
npm run smoke:perception
CV fixture + live perception
npm run smoke:eat
Scripted steer until score↑
npm run smoke:survive
Scripted loop until game over (or MAX_SCORE )
npm run smoke:*:headed
Same with HEADED=1
Environment variables (smoke / headed)
Variable
Default
Meaning
HEADED
unset
1 / true → show Chromium
SNAKE_HTML
vendored agent snake
Path override
MOVE_STEPS
1
Steps per move (agent mode)
MAX_SCORE
none
Survive stops when score ≥ value
MAX_STEPS
50000
Survive step cap
STEP_MS
200
Visual delay between smoke moves
HOLD_MS
5000 – 8000
Keep window open after headed runs
Neon Snake has no clear — only death; MAX_SCORE is an artificial stop for tests.
.
├── README.md / README.zh-CN.md
├── package.json # script proxies
├── docs/ # design & schema
├── pigame/ # main package
│ ├── games/snake.html # DEFAULT agent-mode Neon Snake (vendored)
│ ├── games/README.md
│ ├── extensions/pigame.ts
│ ├── skills/
│ ├── scripts/ # smoke:* (test scaffolding)
│ └── src/
├── .pi/extensions/pigame.ts # Pi entry
└── mini-game-studio-ai-gen/ # optional upstream clone (SNAKE_HTML=…)
Design docs: docs/README.md .
Issue
Fix
Missing pigame/games/snake.html
Pull latest pigame — file is in-repo
Want upstream continuous loop
SNAKE_HTML=./mini-game-studio-ai-gen/games/snake.html (clone Digiman first)
Snake runs by itself while LLM thinks
You are not on agent-mode HTML; check default path / SNAKE_HTML
ENOENT package.json at root
Use root scripts or cd pigame
No Chromium window
HEADED=1 or *:headed ; check WSLg / echo $DISPLAY
Opened Snake yourself, nothing happens
Only Playwright’s window is controlled
Pi extension not loaded
pi -e ./.pi/extensions/pigame.ts then /reload
Acknowledgments
Thanks to Mini Game Studio ( Digiman/mini-game-studio-ai-gen ) for the original Neon Snake and HTML game collection.
We vendor a patched copy under pigame/games/ for Pi + LLM agent-mode testing (paused auto loop + single-step moves).
Work in progress (MVP: Pi + LLM drives Snake via observe / move).
Phase notes: docs/implementation.md .
This is a Pi extension that uses an LLM to drive the game, starting with a simple 2D game.
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
