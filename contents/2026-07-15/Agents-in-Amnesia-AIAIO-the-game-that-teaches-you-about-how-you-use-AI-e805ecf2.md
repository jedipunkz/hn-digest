---
source: "https://github.com/sene1337/aiaio"
hn_url: "https://news.ycombinator.com/item?id=48927378"
title: "Agents in Amnesia (AIAIO) – the game that teaches you about how you use AI"
article_title: "GitHub - sene1337/aiaio: Your agent sessions as playable levels — errors are monsters, tasks are objectives, and the wall of forgetting is right behind you. Claude Code / OpenClaw / Hermes. · GitHub"
author: "notjoemama"
captured_at: "2026-07-15T21:48:42Z"
capture_tool: "hn-digest"
hn_id: 48927378
score: 2
comments: 0
posted_at: "2026-07-15T21:37:15Z"
tags:
  - hacker-news
  - translated
---

# Agents in Amnesia (AIAIO) – the game that teaches you about how you use AI

- HN: [48927378](https://news.ycombinator.com/item?id=48927378)
- Source: [github.com](https://github.com/sene1337/aiaio)
- Score: 2
- Comments: 0
- Posted: 2026-07-15T21:37:15Z

## Translation

タイトル: Agents in Amnesia (AIAIO) – AI の使い方を学ぶゲーム
記事のタイトル: GitHub - sene1337/aiaio: プレイ可能なレベルとしてのエージェント セッション — エラーは怪物、タスクは目標、そして忘却の壁はすぐ後ろにあります。クロード・コード/オープンクロウ/エルメス。 · GitHub
説明: エージェント セッションはプレイ可能なレベルとして行われます。エラーは怪物、タスクは目標、そして忘却の壁がすぐ後ろにあります。クロード・コード/オープンクロウ/エルメス。 - sene1337/アイアイオ

記事本文:
GitHub - sene1337/aiaio: エージェント セッションはプレイ可能なレベルとして行われます。エラーは怪物、タスクは目標、そして忘却の壁がすぐ後ろにあります。クロード・コード/オープンクロウ/エルメス。 · GitHub
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
セネ1337
/
アイアイオ
公共
通知
Y

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
74 コミット 74 コミット .claude .claude .github/ workflows .github/ workflows docs docs サンプル サンプル モックアップ モックアップ qa qa スクリプト スクリプト スキル/ aiaio スキル/ aiaio src src テスト テスト .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md ライセンス ライセンス PRIOR-ART.md PRIOR-ART.md README.md README.md Index.html Index.html package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vite.config.ts vite.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AIAIO: 記憶喪失のエージェント、狂気の兵器
操作: 内部空間。ただし、世界はエージェントの実際のセッション ログです。
🎮 ホストされたデモをプレイする (**THE OPENCLAW
エルメス キャンペーン**、明示的に架空の 12 レベル)。本当のゲームはあなた自身の歴史です。このリポジトリをクローンして実行します。
npm run scan 、およびすべての Claude Code、OpenClaw、または Hermise セッション
マシンがプレイ可能なレベルになります。プロンプトはタスクであり、エラーは
モンスター、あなたの圧縮物はあなたを追いかける壁です。
ベータノート。デスクトップとキーボードが必要です (モバイルは正直なゲートを取得します)。
SessionCard はローカルで生成され、マシンから離れることはありません。カードを共有する
誰かと一緒にすることは可能ですが実験的です: カードには編集されたスニペットが含まれています
実際のプロンプトなので、送信する前に public/cards/<card>.json を自分で読んでください。
それを誰にでも。適切な事前共有レビュー UI が次のベータ版の目玉です。
実際のセッションはレベルです。そのタイムラインは、あなたが通過する地形です。あなたの
実際のエラーは、実際に発生した場所でモンスターとして発生します。あなたの
実際のタスクはワークステーションとして世界に存在します。そしてあなたの後ろにはいつも壁がある
忘却の(

コンテキストの圧力により空間的に) 前進、地形、タスクの摂取、
そして最終的にはあなた。
プロセス出口 0 が生きている状態で到達します。完璧に実行するには、途中でタスクキューをクリアしてください。
現実の歴史のキャンペーンで、少なくとも半分を回復した後に終了します。
記録されたタスク作業は進行クレジットを獲得します。
あなたが戦いに費やすすべてのトークンは、壁が獲得する距離です。
このゲームが気に入った場合、またはアイデアがある場合は、私に連絡してください
ツイッターとか
インスタグラム @bradmillscan 。
npmインストール
npm run scan # オプション: セッションからの自動ビルド レベル (以下を参照)
npm run dev # 出力された URL を開きます
npm run build # dist/ の静的バンドル
実稼働ビルドではネットワーク呼び出しは行われず、架空のパブリックのみが出荷されます。
キャンペーン。ローカル エンリッチメントは明示的なオプトイン開発アクションであり、AI を呼び出す場合があります。
プライバシー通知を表示した後に設定するコマンド。
またはエージェントにすべてを任せることもできます
このゲームはエージェントを実行する人向けに作られているため、エージェントがインストーラーになることができます。
レベルキュレーターと技術サポート。エルメスユーザー：
hermes スキルインストール https://raw.githubusercontent.com/sene1337/aiaio/main/skills/aiaio/SKILL.md
次に、エージェントに「aiaio を設定してください」、「最もドラマチックな 10 件を見つけてください」などと伝えます。
セッションを行ってレベルを上げます」または「アナウンサーを苦いゴルフにします」
コメンテーター」。他のエージェント (Claude Code、OpenClaw、Codex): をポイントします。
AGENTS.md — 同じプレイブックで、スキルは必要ありません。金庫が出てきたら
空の場合、npm run Doctor は、スキャンされた内容と各ファイルがスキャンされた理由を正確に出力します。
拒否されました。エージェントはそれを読んで原因を解決できます。
キャンペーンのプレゼンテーションはカスタマイズ可能です。エージェントはタイトルを書いたり、コピーを表示したり、
不変のソースカードのスナップショットに対するオブザーバーのライン。事実に基づくキャンペーン
ソース由来のメカニズムを保持します。明示的なリミックス リクエストでは、1 つを選択できます。
固定された穏やかな、バランスのとれた、または残忍な戦闘プロファイル。それは別個にあります
進歩

常に Remix というラベルが付いています。
←/→移動します。 ↑ジャンプ。 ↓ 急速落下（掘削されたクレーターを探索）。スペース
（対面に向かって）火をつけます。
W（長押し）ステーションでタスクを実行します。あなたは根を張って頭を下げていますが、
これは、+25% のダメージを受けることを意味します。働くことが勝利への近道であり、働くことは
自分が一番弱いとき。
U ⬆ クレートのコマンド メニューを開く (1/2/3 で選択)、または ◈ モデルをインストールします
アップグレードします。 C は /compact を実行します。
[ ] / 1-9 武器を切り替えます。スロット 1 は ∞ 印刷デバッグ ザッパーです。
印刷ステートメントがなくなることはありません。また、ステートメントだけでは決して勝つことはできません。
S はサブエージェントを生成します (900tk と ~16tk/s の維持費、推論が行われないため)
自由になり、点滴は文字通り壁を加速します）。これは下位モデルです: 弱いザップ、
上限は 2 で、幻覚幽霊が触れるか落ちるまで忠実です。
壁の中へ。そして、それは破損し、赤くなり、文字化けし、あなたに向かって撃ちます。
それを終了させるか、それを追い越します（不正者は約 18 秒後に OOM で自殺します）。
M ミュート。すべてのオーディオは WebAudio で合成されます (チップ/グリッチ、ゼロ アセット):
圧縮とは静的な状態への途切れ途切れの降下であり、壁には鼓動が起こります。
もうすぐです。タスクの発送時にチャイムが鳴ります。
忘却の壁は行動によって決まります。タイマーが作動することはありません。
あなたが燃やすすべてのトークンは壁の距離になります (トークンごとに 0.16 ピクセル): 発射、動作、
サブエージェントの維持、ダメージの吐き出し、さらには歩行（記録を読むと推測されます）
10pxあたり1tkでも)。完全に静止していれば、それはあなたと一緒に静止します。あなたの
トークン法案は嵐だ。オーバーフロー送信者のスパム トークンがメーターに送られてくる間、
あなたが近くにいると、壁が動きます。すべてがそうなるからです。渡されるタスクは次のとおりです
忘れられた（文字化け、回復不能）。サブエージェントを丸ごと食べます。あなたはただそれを
遺跡: ゾーン内では毎秒 9 馬力の出血があり、武器が激しく飛び散りますが、
そこにある未取得の木箱はまだ機能するため、まだ飛び込むことができます。関与

アンタリー
圧縮（しきい値を超える）すると、240 ピクセル以上に急上昇します。
記憶喪失: シールドがなくなり、進行状況が巻き戻され、3 回目の圧縮からは
完了したタスクは出荷解除できます。
/コンパクト(C)。任意の場所で、20 秒のクールダウン。メーターが消耗します
きれいに、急増も忘れ物もありませんが、要約パスにはコストがかかります
250tk、そしてあなたはビートのために頭を下げます。しきい値に近づきすぎると、
峠そのものが転倒してしまいます。早くコンパクトに、頻繁にコンパクトに。
コンテキストエコノミー。燃焼されるすべてのトークンは壁の距離とメーター自体です
次の非自発的な圧縮サージへのカウントダウンです。の
context-window-nuke は画面の半分のエラーを消去し、画面の 4 分の 1 をあふれさせます。
自分のメートル (壁の約 700 ピクセル) なので、暴力は慎重に選択してください。殴られる
コンテキストにエラー吐き出しを挿入する (スタック トレースが長い) ため、ダメージが発生します
自分自身の圧縮を加速します。
アップグレード。 ⬆ パッチクレートは 3 つのオプションのコマンド メニューを開きます: 古典的なランダム
ギャンブルは常にオプション 1 に加えて、/restore キャッシュされたサブエージェントからの 2 つのシード済みピックです。
(維持点滴なし)、/tune context-manager (+5% しきい値)、/patch Shield-buffer 、および /restock error-log 。レアな◈ MODEL UPGRADE 木箱
良好な状態を維持: より多くのコンテキスト バジェット、より高いしきい値、シールドの補充、
vN タイトルバーのチェックマーク。
音楽。生成的なアンビエント スコア (純粋な WebAudio): ウォーム パッドとペンタトニック
安全なときに摘み取り、壁のように暗く、まばらに、そしてより離調して変化します
閉じて、ドローンに至るまで、そして忘却の中で静的に。
ハンディキャップ。安定性の低いセッション（あなたのセッションは荒かった）は開始シールドを付与します
そしてダメージボーナス。「なぜ」の行に実際の数値が引用されています。
動物寓話 (エラーログから生成)
あなたのエラーのカテゴリ
になる
行動
タイムアウト
⏱ タイムアウトブロブ
タンキーロバー。ショットが遅く爆発する
幻覚
👻幻覚ゴ

セント
フェーズ、テレポート、それが存在することは確かです。あなたの硬化はその接触に抵抗します
回帰
☢ 回帰スプリッター
死ぬと2つのミニに分割される
再起動/クラッシュ
🔁 クローラーを再起動する
死んでから一度再起動する
偽陽性
⚡ 誤検知スナイパー
電信レーザー、100% の信頼性、~70% の精度
ツールエラー
🔧 ツールタレット
割り込みボルト。作業中にヒットする = タスクの進行状況が失われる
コンテキストオーバーフロー
📈 オーバーフローエミッター
優先ターゲット: あなたが近くにいる間にトークンをメーターにスパムし、壁を動かします。
回復
➕ 回復スプライト
フレンドリー、HP/シールドのタッチ
カテゴリごとの敵の数はエラー数に応じて調整され、武器の弾薬は次のように調整されます。
同じログなので、最悪のエラー カテゴリは最大の脅威と最大の脅威の両方になります。
最も深い雑誌。
セマンティックレイヤー: レベルはセッションのストーリーを伝えます
カードにそれらが含まれている場合 (エクストラクターはこれらすべてを自動的にマイニングします):
目標を達成するあなたの最初の本当の質問は、ミッションとして示されています。
自分の言葉：…」
タスクは実際のメッセージです。ログに構造化されたタスクがない場合、
実質的なユーザーの要求は、あなたの言葉で名付けられたタスクステーションとなり、次の場所に配置されます。
実際のタイムラインの位置、作業単位はどれだけの量に比例するか
各アスクが実際に消費したセッション。
敵は実際にエラーが発生した場所 (errors[].at) にスポーンします。
◇瞬間。本当の勝利 (「OK、うまくいきました。誰も何も触れません。」)
フラストレーション（「まだ壊れている。まだ壊れている。」）は、彼らがいる世界に立っています。
起こった。 1つを通り過ぎると、記録がそれを引用します。
圧縮すると自分の言葉が文字化けします。バナーにより実際の行が破損します。
缶詰めのフィラーではなくセッションであり、要約によって実際の部分がどこにあるかがわかります。
セッションの実行が終了しました。
すべてはヒューリスティックかつ決定論的で、オンマシンで行われ、LLM は関与しません。
墨消しはカードを除くすべてのスニペットに適用されます

あなたの断片がまだ残っています
実際のプロンプトなので、共有する前にざっと読んでください。
キャンペーンの強化: 明示的なローカル アクション
メイン メニューの「✦ 歴史を豊かにする」では、次の 2 つの正直な形状が提供されます。
SHAPE MY OPENING には 6 つの対象セッションが必要で、時系列が保持されます。
BUILD MY CAMPAIGN には 15 が必要で、15 ～ 24 の対象となるセッションが厳選されます。
編集された SessionCard の抜粋を送信する前に、同意通知が表示されます。
設定された AI コマンド。単にセッションが自動的に強化されることはありません
開くか再生するか。生のカードは変更されません。エンリッチメントは、
ソース ダイジェスト、順序付けされたエントリ、および
プレゼンテーションのコピー。ライターが利用できない場合、AIAIO はアトミックに
代わりに、決定論的なベースライン キャンペーンを完了してください。
エージェント主導のフローの場合は、同じエンジンを使用します。
npm run rich -- --selection Hardest --pace 激しさ --remix 残忍な
# デフォルトのライターは claude -p です。別の CLI エージェントの AIAIO_LLM_CMD を設定する
--remix 穏やか|バランスの取れた|残酷な表現が明示的です。ジェントルは敵対的な予算を 70% 使用し、
80% のダメージ、+10 シールド、オブザーバーの注入なし。バランスが保たれている
正常値。残忍な敵対予算を 150% 使用 (合計上限 36 / タイプごとに 10)、
125%のダメージを与え、安定性のハンディキャップを取り除きます。リミックスのロック解除は決して変更されません
ファ

[切り捨てられた]

## Original Extract

Your agent sessions as playable levels — errors are monsters, tasks are objectives, and the wall of forgetting is right behind you. Claude Code / OpenClaw / Hermes. - sene1337/aiaio

GitHub - sene1337/aiaio: Your agent sessions as playable levels — errors are monsters, tasks are objectives, and the wall of forgetting is right behind you. Claude Code / OpenClaw / Hermes. · GitHub
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
sene1337
/
aiaio
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
74 Commits 74 Commits .claude .claude .github/ workflows .github/ workflows docs docs examples examples mockups mockups qa qa scripts scripts skills/ aiaio skills/ aiaio src src tests tests .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md LICENSE LICENSE PRIOR-ART.md PRIOR-ART.md README.md README.md index.html index.html package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vite.config.ts vite.config.ts View all files Repository files navigation
AIAIO: Agents In Amnesia, Insane Ordnance
Operation: Inner Space, except the world is your agent's actual session log.
🎮 Play the hosted demo (**THE OPENCLAW
HERMES CAMPAIGN**, 12 explicitly fictional levels). The real game is your own history: clone this repo, run
npm run scan , and every Claude Code, OpenClaw, or Hermes session on your
machine becomes a playable level. Your prompts are the tasks, your errors are
the monsters, your compactions are the wall chasing you.
Beta notes. Desktop and keyboard required (mobile gets an honest gate).
SessionCards are generated locally and never leave your machine. Sharing a card
with someone is possible but experimental: cards contain redacted snippets of
your real prompts, so read public/cards/<card>.json yourself before you send
it to anyone. A proper pre-share review UI is the headline of the next beta.
Your real session is the level. Its timeline is the terrain you cross. Your
real errors spawn as monsters at the points where they actually happened. Your
real tasks sit in the world as work stations. And behind you, always, the WALL
OF FORGETTING (context pressure made spatial) advances, eating terrain, tasks,
and eventually you.
Reach process exit 0 alive. Clear your task queue on the way for a perfect run;
in your real-history campaign, exiting after recovering at least half of the
recorded task work earns progression credit.
Every token you spend fighting is distance the wall gains.
If you like this game or have ideas, hit me up on
Twitter or
Instagram @bradmillscan .
npm install
npm run scan # optional: auto-build levels from YOUR sessions (see below)
npm run dev # open the printed URL
npm run build # static bundle in dist/
Production builds make no network calls and ship only the fictional public
campaign. Local enrichment is an explicit opt-in dev action and may call the AI
command you configure after showing its privacy notice.
Or let your agent do all of it
This game is made for people who run agents, so the agent can be the installer,
the level curator, and the tech support. Hermes users:
hermes skills install https://raw.githubusercontent.com/sene1337/aiaio/main/skills/aiaio/SKILL.md
then tell your agent things like "set up aiaio" , "find my 10 most dramatic
sessions and make them levels" , or "make the announcer a bitter golf
commentator" . Any other agent (Claude Code, OpenClaw, Codex): point it at
AGENTS.md — same playbook, no skill required. If the vault comes up
empty, npm run doctor prints exactly what was scanned and why each file was
rejected; your agent can read it and fix the cause.
Campaign presentation is customizable: agents can write titles, display copy,
and Observer lines over an immutable source-card snapshot. Factual campaigns
retain source-derived mechanics. An explicit Remix request may select one
of the fixed gentle , balanced , or brutal combat profiles; it has separate
progress and is always labelled Remix.
←/→ move. ↑ jump. ↓ fast-fall (explore drilled craters). space
fire (toward facing).
W (hold) work the task at a station. You're rooted and heads-down while
working, which means +25% damage taken. Working is how you win, and working is
when you're weakest.
U open a ⬆ crate's command menu (choose with 1/2/3), or install a ◈ model
upgrade. C runs /compact .
[ ] / 1-9 switch weapons. Slot 1 is the ∞ print-debug zapper : you can
never run out of print statements, and you can never win with them alone.
S spawn a subagent (900tk plus ~16tk/s upkeep, because inference isn't
free and the drip literally speeds up the wall). It's a lower model: weak zaps,
cap of 2, loyal right up until a hallucination-ghost touches one or it falls
into the wall. Then it's corrupted : red, garbled, and shooting at you.
Terminate it or outrun it (rogues OOM-kill themselves after about 18s).
M mute. All audio is synthesized WebAudio (chip/glitch, zero assets):
compaction is a stuttering descent into static, the wall has a heartbeat when
it's close, tasks chime when they ship.
The wall of forgetting is action-driven. It does NOT creep on a timer.
Every token you burn becomes wall distance (0.16px per token): firing, working,
subagent upkeep, damage spew, even walking (reading the transcript is inference
too, at 1tk per 10px). Stand perfectly still and it stands still with you. Your
token bill is the storm. overflow-emitter s spam tokens into your meter while
you're near, which moves the wall, because everything does. Tasks it passes are
forgotten (garbled, unrecoverable). It eats subagents whole. You it merely
ruins: inside the zone you bleed 9hp/s and your weapons spray wildly, though you
can still dive in, since unclaimed crates in there still work. Involuntary
compaction (crossing your threshold) surges it 240px or more, plus the
amnesia: shield gone, progress rewound, and from your 3rd compaction even
completed tasks can un-ship.
/compact (C). Voluntary, anywhere, 20s cooldown. It drains your meter
cleanly, with no surge and nothing forgotten, but the summarization pass costs
250tk and roots you heads-down for a beat. Run it too close to the threshold and
the pass itself tips you over. Compact early, compact often.
Context economy. Every token burned is wall distance, and the meter itself
is your countdown to the next involuntary compaction surge. The
context-window-nuke erases half a screen of errors and floods a quarter of your
own meter (about 700px of wall), so choose violence carefully. Getting hit
injects error-spew into your context (stack traces are long), so damage
accelerates your own compaction.
Upgrades. ⬆ patch crates open a 3-option command menu: the classic random
gamble is always option 1, plus two seeded picks from /restore cached-subagent
(no upkeep drip), /tune context-manager (+5% threshold), /patch shield-buffer , and /restock error-log . The rare ◈ MODEL UPGRADE crate
stays guaranteed-good: more context budget, a higher threshold, a shield refill,
and a vN title-bar tick.
Music. A generative ambient score (pure WebAudio): a warm pad and pentatonic
plucks when you're safe, morphing darker, sparser, and more detuned as the wall
closes, down to drone and static inside the forgetting.
Handicap. Low-stability sessions (yours was rough) grant a starting shield
and a damage bonus, with the "why" line quoting your real numbers.
The bestiary (spawned from your error log)
Your error category
Becomes
Behavior
timeout
⏱ timeout-blob
tanky lobber; shots detonate late
hallucination
👻 hallucination-ghost
phases, teleports, is sure it exists; your hardening resists its touch
regression
☢ regression-splitter
splits into two minis on death
restart / crash
🔁 restart-crawler
relaunches itself once after dying
false_positive
⚡ false-positive-sniper
telegraphed laser, 100% confidence, ~70% accuracy
tool_error
🔧 tool-turret
interrupt bolts; hit while working = lose task progress
context_overflow
📈 overflow-emitter
priority target : spams tokens into your meter while you're near, which moves the wall
recovery
➕ recovery-sprite
friendly , touch for hp/shield
Enemy count per category scales with the error count, and weapon ammo scales from
the same log, so your worst error category is both your biggest threat and your
deepest magazine.
The semantic layer: the level tells your session's story
When a card carries them (the extractor mines all of this automatically):
goal . Your first real ask, shown as the mission: "the mission, in your
own words: …"
Tasks are your actual messages. When logs have no structured tasks, your
substantive user asks become the task stations, named in your words, placed at
their real timeline positions , with work-units proportional to how much of
the session each ask actually consumed.
Enemies spawn where the errors really happened ( errors[].at ).
◇ moments. Real wins ("ok it works. nobody touch anything.") and
frustrations ("still broken. still. broken.") stand in the world where they
happened. Walk past one and the transcript quotes it.
Compaction garbles your own words. The banner corrupts real lines from the
session instead of canned filler, and the recap tells you where in the real
session your run ended.
All of it is heuristic, deterministic, and on-machine, with no LLM involved.
Redaction applies to every snippet, but a card still contains fragments of your
actual prompts, so skim before sharing.
Campaign enrichment: an explicit local action
The main menu's ✦ ENRICH YOUR HISTORY offers two honest shapes:
SHAPE MY OPENING requires six eligible sessions and preserves chronological order.
BUILD MY CAMPAIGN requires fifteen and curates 15–24 eligible sessions.
It displays a consent notice before it sends redacted SessionCard excerpts to
the configured AI command. No session is automatically enriched merely by
opening or playing it. The raw card remains unchanged; enrichment publishes a
versioned CampaignManifest overlay with source digests, ordered entries, and
presentation copy. If the writer is unavailable, AIAIO atomically publishes a
complete deterministic baseline campaign instead.
For an agent-driven flow, use the same engine:
npm run enrich -- --selection hardest --pace intense --remix brutal
# default writer is claude -p; set AIAIO_LLM_CMD for another CLI agent
--remix gentle|balanced|brutal is explicit. gentle uses 70% hostile budget,
80% damage, +10 shield, and no Observer injections; balanced preserves the
normal values; brutal uses 150% hostile budget (cap 36 total / 10 per type),
125% damage, and removes the stability handicap. Remix unlocks never alter
fa

[truncated]
