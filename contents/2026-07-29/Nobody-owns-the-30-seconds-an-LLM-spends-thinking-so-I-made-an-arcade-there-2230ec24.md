---
source: "https://github.com/The747Lab/747-terminal-games"
hn_url: "https://news.ycombinator.com/item?id=49101497"
title: "Nobody owns the 30 seconds an LLM spends thinking, so I made an arcade there"
article_title: "GitHub - The747Lab/747-terminal-games: Arcade, reborn in your terminal — five games that open while Claude thinks and vanish when it replies. Developed by The 747 Lab. · GitHub"
author: "the747lab"
captured_at: "2026-07-29T18:57:21Z"
capture_tool: "hn-digest"
hn_id: 49101497
score: 1
comments: 0
posted_at: "2026-07-29T18:53:10Z"
tags:
  - hacker-news
  - translated
---

# Nobody owns the 30 seconds an LLM spends thinking, so I made an arcade there

- HN: [49101497](https://news.ycombinator.com/item?id=49101497)
- Source: [github.com](https://github.com/The747Lab/747-terminal-games)
- Score: 1
- Comments: 0
- Posted: 2026-07-29T18:53:10Z

## Translation

タイトル: LLM が考える 30 秒は誰も所有していないので、そこにアーケードを作りました
記事のタイトル: GitHub - The747Lab/747-terminal-games: アーケード、ターミナルで生まれ変わる — クロードが考えている間に開き、応答すると消える 5 つのゲーム。 747ラボによって開発されました。 · GitHub
説明: アーケード、あなたのターミナルで生まれ変わります — クロードが考えている間に開き、応答すると消える 5 つのゲーム。 747ラボによって開発されました。 - The747Lab/747-ターミナル-ゲーム

記事本文:
GitHub - The747Lab/747-terminal-games: ターミナルで生まれ変わるアーケード — クロードが考えている間に開き、応答すると消える 5 つのゲーム。 747ラボによって開発されました。 · GitHub
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
エラーがありましたw

ハイルロード中。このページをリロードしてください。
ザ747ラボ
/
747 ターミナル ゲーム
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
20 コミット 20 コミット .claude-plugin .claude-plugin .github/ workflows .github/ workflows アセット アセット コマンド コマンド ゲーム ゲーム フック フック スクリプト スクリプト .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md ライセンスライセンス README.md README.md SECURITY.md SECURITY.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI がコーディングしている間、端末でゲームをプレイします。
クロード コードが考えているとき、作業の横のペインにゲームが表示されます。クロードが応答するとすぐに一時停止されるため、答えを見逃すことはありません。次のプロンプトを送信すると再開されます (同じ実行、同じスコア)。いつでも閉じてください。プラグインのインストール以外のセットアップは不要で、依存関係も不要です。Python 3 と Curses のみで、他には何もありません。
インストール — 2 行、シェル スクリプトなし
/プラグイン マーケットプレイス追加 The747Lab/747-terminal-games
/plugin install 747-terminal-games@747-terminal-games
それだけです。カールするものは何もなく、シェルにパイプするものは何もなく、ルートもネットワーク呼び出しもありません
ランタイム — CI は出荷されたコードを grep し、ソケットについて言及している場合はビルドに失敗します。
MIT であり、依存関係ツリーは空です。
次回クロードが考えたとき、あなたは一度プレイするかどうか尋ねられます - 試してみてください、
a は毎回自動的に開きます。o は何も尋ねません。 「はい」と答えるとピッカーが表示されます。
1 ～ 5 または矢印キーでタイトルを選択します。あなたの選択は記憶されます。
正直なところ、事前に制限が 1 つあります。自動オープンには tmux または iTerm2 が必要です。それ以外のどこでも、
ゲームは依然としてスタンドアロンで実行されますが、出現と消滅が失われるだけです。
初めてゲームが開くときに、いずれかを選択します

。 5 つすべてが HUD を備えた完成したゲームです。
スコア表、勝利（または深度）画面、自己ベストファイル - デモではありません。
以下のすべてのクリップは 1 つの連続実行であり、ペインが実際に開くサイズで記録されます。
実際に実行した速度で再生します。何も加速せず、何もカットされません。
壁をクリアすることは決してありません。天井のハッチを割って、ボールを元に戻します
穴を通って上の部屋に登ります。それからまたそれをやります。永遠に。
C 1 が重要な唯一の数字であり、それは常に増加するだけです。
▌ 侵入 · 160 ♥♥ C 1 違反 ▰▱▱▱▱
・・・・・・・・
▔▔▛▜▛▜▛▜▔▔▔▔▔▔▔▔▔▔▔▛▜▛▜▛▜ ▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔ ▛▜▛▜▛▜▔▔▔▔▔▔▔▔▔▔▔▛▜▛▜▛▜▔▔
▄▄▄▄▄ █████ █████ █████ █████ ▄▄▄▄▄ █████ ▄▄▄▄▄ █████ █████ █████ ▄▄▄▄▄
▄▄▄▄▄ ▄▄▄▄▄ ▄▄▄▄▄ █████ █████ ▄▄▄▄▄ █████ ▄▄▄▄▄ ▄▄▄▄▄ ▄▄▄▄▄ █████ ▄▄▄▄▄
▄▄▄▄▄ ▄▄▄▄▄ █████ ▄▄▄▄▄ █████ █████ █████ ▄▄▄▄▄ ▄▄▄▄▄ █████ ▄▄▄▄▄ ▄▄▄▄▄
▄▄▄▄▄ ▄▄▄▄▄ █████ ▄▄▄▄▄ ▄▄▄▄▄ █████ ▄▄▄▄▄ ▄▄▄▄▄ █████ ▄▄▄▄▄ ▄▄▄▄▄
●
▀▀▀▀▀▀▀▀
チャンバー1・ベスト1 THE 747 LAB
天井のシアン▛▜パネルが出口で、BREACH▰▱▱▱▱がどれだけ近づいているかを示しています
あなたはそうです。ライフを失うことなくチャンバーを占領し、連勝数を増やす

エス; 7 回ごと
チャンバーは金庫です。
通過する車のフロントガラスから見た 7 区間の配達風景
星間空間。エイリアンの宇宙船を撃ち、岩を避け、最後にあるゲートに糸を通します。
あらゆる分野。セクター 7 をクリアするとランは完了します。その後、必要に応じて OVERRUN で飛行を続けてください。
▌ スカイラン · 40 ▰▰▱S 1/7 ▸▹▹▹▹▹▹▹ 13% ×1
・☩・
☩◄◄‹◦›►►☩ ·
▏ · ▕ ·
▏ · ▕ ·
▏ ▕
▏▲▕
◉███████████████████████████◉
←→ 移動 · [スペース] シュート ☩ · 回避 █ · つかむ ◈ THE 747 LAB
S 1/7 と距離バーは最初のフレームから画面上に表示されるため、いつでもどこにあるかがわかります。
走行が終了します。シールドは 3 枚、命はありません。脱出する唯一の方法は、船体を使い果たすことです。
同じフィクション、反対側のカメラ。ボタンを 1 つ上に、ボタンを 1 つ下に移動: ソリッドをジャンプします
脆いディザーをブロックし、叩きつけます。 7 つのゲート、7,470 メートル、HUD に 1 つの数字。
推進力はスピードなので、推進力を集めると、走るのが速くなり、短くなります。
▌ ジェットウォッシュ · 741 ▣▣▣▢ G1 ▸▸▸▸▸▸▸▹▹▹ ▮▮▮▯▯▯▯▯▯▯
· · · -
· · · ·
▓▓▓▓ ▓▓▓▓▓
▓▓▓ ▓▓▓ ▓▓▓▓ ▓▓▓▓▓▓▓▓ ▓▓▓▓▓ ▓▓▓▓
═======================== ═======================== ═======================== ═========================
▛▜ ◈ ◈ ║║ ▛▜
║║
◈ ◈ ║║
► ║║ ◈ ▟▙
▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀ ▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀

▀▀▀▀▀▀▀▀▀▀
╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱
╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱
↑ジャンプ・↓スラム・█ソリッド・▒ブレイク THE 747 LAB
すべての危険は、あなたに届く前に 1 秒以上前に電報で伝えられるため、攻撃は常に発生します。
あなたが間違った決断をしたとしても、決して驚くべきことではありません。
7 つの波、そして大きな波。一番上の列は 3 倍の価値があるため、貪欲に選択するのが現実です。
1.2 秒以内にキルを連鎖させてマルチプライヤーを構築しましょう。ウェーブ3でバンカーが到着、
ウェーブ5のダイバー。
▌ アストロズ · 0 ▲▲ W1/7 ▰▰▰▰▰▰▰▰▰▰
▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼
▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼
▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼
▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼
▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼
！
←→移動[スペース]で射撃を回避！
▲
WAVE 1・BEST 0 THE 747 LAB
冒頭の数秒では、チュートリアル ボックスではなくレイアウトによってゲームを説明します: 1 つの爆弾
一人で何もない空間に落ちるので、あなたは何を学びますか？費用がかかる前にという意味です。
4 車線の交通、3 車線の川、7 つの湾が埋められます。道は待っている
ゲーム — 立ち止まって、ギャップが来るのを待ちましょう。川は反対です、なぜなら水だからです
それはあなたが立っていることができる車線ではありません。通り過ぎていくものに乗るか、渡らないかです。
▒ ▒ ▽ ▽ ▽ ▒ ·▽· ▽ ▒ ▽▒ ▽
▄▄▄▄▄▄▄▄ · ▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄▄ ·▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄▄ ·▄▄
▄▄▄▄▄ ▄▄▄▄▄▄▄▄ ~~▄▄▄▄▄▄▄▄~ ▄▄▄▄▄▄

▄▄ ▄▄▄▄▄▄▄▄ ·▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄
▄▄▄▄▄ ▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄▄ · ▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄
▄▄▄▄▄▄▄ · ▄▄▄▄▄▄▄▄ ·▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄▄~ ▄▄▄▄▄▄▄▄ ~▄▄▄▄▄▄▄▄ ~ ▄▄▄▄▄▄▄▄ ▄▄▄▄
▄▄ ▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄▄ · ▄▄▄▄▄▄▄▄ ·▄▄▄▄▄▄▄▄ ·▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄▄ · ▄▄▄▄▄▄▄▄
▒▒ ▒ ▒ ▒ ▒ ▒ ▒ ▒ ▒ ▒ ▒ ▒ ▒▒ ▒
╌ ╌▬▶ ╌ ╌▬▶ ╌ ╌▬▶ ╌ ╌▬▶ ╌ ╌▬▶
╌ ◀▬╌ ╌ ◀▬╌ ╌ ◀▬╌ ╌ ◀▬╌ ╌ ◀▬╌
╌ ╌ ╌ ╌ ╌▌✷▐ ╌ ╌ ╌ ╌ ╌
╌ ╌◀██ ╌ ╌ ◀██ ╌ ╌ ◀██ ╌ ╌ ◀██ ╌ ╌ ◀█
▶ ╌ ╌ ▬▶ ╌ ╌ ▬▶ ╌ ╌ ▬▶ ╌ ╌ ▬▶ ╌ ╌ ▬
▒▒▒ ▒ ▒ ▒ ▒ ▒ ▒ ▒ ▒
ジェイウォーク 80 ◆◆ ▽▽▽▽▽▽▽ R1 [↑] クロス [q] 閉じる
▽▽▽▽▽▽▽ はスコアボード全体です。開始時に 7 つの空のベイがあり、最後のベイが閉じられます。
車の速度が速くなり、あなたは縁石に戻ります。誰も言われる必要がなかった
この中で何をすべきか、まさにそれがここにある理由です。
このページの上部にあるクリップは、JAYウォークの 1 回の走行です。ベイを埋め、丸太を走り、そして 2 回の
そこに至るまでに費やした3つの人生。
どのゲームでも — q でゲームを閉じます。クロードが応答すると、ゲームは一時停止します。スペース（または
SKYRUN の p) はとにかく実行し続けます。
スラッシュ コマンドは、クロードの状態を無視して、いつでもフリープレイ ラウンドを開始します。
/breakin · /skyrun · /jetwas

h · /astros · /jaywalk — および /breakout は引き続き動作します
/breakin のエイリアスとして。
tmux (推奨) — クロードが考えている間、ゲームはウィンドウ内で分割ペインとして開き、クロードが応答するとすぐに消えます — ターミナルは元の状態に戻り、閉じるタブはありません。あなたの実行はバックグラウンドで維持され、ゲーム中に次のプロンプトで再び参加します。
tmux なしの iTerm2 — ゲームは現在の iTerm ウィンドウをネイティブに分割します (同じインプレース感覚)。
他の場所 — 自動オープンには tmux が必要です。ゲームはいつでもスタンドアロンで実行できます: python3 games/jetwash.py --free 。
Python 3.9 以降、curses (macOS および Linux の標準)。 pip インストールは一切ありません。
すべてのタイトルは 80×8 ペインまで読みやすいままで、16 色に戻り、また元に戻ります。
UTF-8 を実行できない端末では純粋な ASCII に変換します。
~/.747-terminal-games/mode にある小さなファイルが自動起動を制御します。
ピッカーで最後に選択したもの。 1つの単語として保存されます
~/.747-terminal-games/game なので、直接設定することもできます。
echo Jetwash > ~/.747-terminal-games/game
スラッシュ コマンドはこのファイルを無視します。常に指定したタイトルが開きます。
ピッカーが表示されるまで、テキストモードのフライスルー: 7.47 秒かかり、任意のキーを押すとスキップされます。
即座に。
747 ラボによって構築されました。端末ゲームのラインナップが増えています。
それぞれに 747 が隠されています。それはラベルには決して記載されていません - それは
ゲーム。見つけてください。それは何かをします。
MITライセンス取得済み。 PR と新しいゲームは歓迎です — COTRIBUTING.md を参照してください。
アーケード、あなたのターミナルに生まれ変わる — クロードが考えている間に開き、応答すると消える 5 つのゲーム。 747ラボによって開発されました。
github.com/The747Lab/747-terminal-games/releases/tag/v2.0.0 トピック
Readme MIT ライセンスの行動規範
セキュリティ ポリシー アクティビティ カスタム プロパティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション

イゲーション
私の個人情報を共有しないでください

## Original Extract

Arcade, reborn in your terminal — five games that open while Claude thinks and vanish when it replies. Developed by The 747 Lab. - The747Lab/747-terminal-games

GitHub - The747Lab/747-terminal-games: Arcade, reborn in your terminal — five games that open while Claude thinks and vanish when it replies. Developed by The 747 Lab. · GitHub
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
The747Lab
/
747-terminal-games
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
20 Commits 20 Commits .claude-plugin .claude-plugin .github/ workflows .github/ workflows assets assets commands commands games games hooks hooks scripts scripts .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md View all files Repository files navigation
Play a game in your terminal while your AI codes.
When Claude Code is thinking, a game appears in a pane beside your work. It pauses the instant Claude replies so you never miss the answer, and resumes when you send your next prompt — same run, same score. Close it any time. Zero setup beyond installing the plugin, and zero dependencies: Python 3 and curses , nothing else.
Install — two lines, no shell script
/plugin marketplace add The747Lab/747-terminal-games
/plugin install 747-terminal-games@747-terminal-games
That's it. Nothing to curl , nothing to pipe into a shell, no root, no network calls at
runtime — CI greps the shipped code and fails the build if it so much as mentions a socket.
MIT, and the dependency tree is empty.
The next time Claude thinks, you'll be asked once whether you want to play — y to try it,
a to auto-open every time, o to never ask. Say yes and the picker comes up: choose a
title with 1 – 5 or the arrow keys. Your choice is remembered.
One honest limitation up front: the auto-open needs tmux or iTerm2 . Anywhere else the
games still run standalone, you just lose the appearing and vanishing.
The first time a game opens, you pick one. All five are finished games with a HUD, a
scoring table, a win (or a depth) screen and a personal-best file — not demos.
Every clip below is one continuous run, recorded at the size the pane actually opens at
and played back at the speed it actually ran. Nothing is sped up and nothing is cut.
You never clear the wall. You crack a hatch in the ceiling , thread the ball back up
through the hole, and climb into the chamber above. Then you do it again. Forever.
C 1 is the only number that matters, and it only ever goes up.
▌ BREAK-IN · 160 ♥♥ C 1 BREACH ▰▱▱▱▱
······
▔▔▛▜▛▜▛▜▔▔▔▔▔▔▔▔▔▔▔▛▜▛▜▛▜ ▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▛▜▛▜▛▜▔▔▔▔▔▔▔▔▔▔▔▛▜▛▜▛▜▔▔
▄▄▄▄▄ █████ █████ █████ █████ ▄▄▄▄▄ █████ ▄▄▄▄▄ █████ █████ █████ ▄▄▄▄▄
▄▄▄▄▄ ▄▄▄▄▄ ▄▄▄▄▄ █████ █████ ▄▄▄▄▄ █████ ▄▄▄▄▄ ▄▄▄▄▄ ▄▄▄▄▄ █████ ▄▄▄▄▄
▄▄▄▄▄ ▄▄▄▄▄ █████ ▄▄▄▄▄ █████ █████ █████ ▄▄▄▄▄ ▄▄▄▄▄ █████ ▄▄▄▄▄ ▄▄▄▄▄
▄▄▄▄▄ ▄▄▄▄▄ █████ ▄▄▄▄▄ ▄▄▄▄▄ █████ ▄▄▄▄▄ ▄▄▄▄▄ █████ ▄▄▄▄▄ ▄▄▄▄▄
●
▀▀▀▀▀▀▀▀
CHAMBER 1 · BEST 1 THE 747 LAB
The cyan ▛▜ panels in the ceiling are the way out, and BREACH ▰▱▱▱▱ says how close
you are. Take a chamber without losing a life and your streak multiplies; every seventh
chamber is a vault.
A seven-sector delivery run , seen from the windshield of a car flying through
interstellar space. Shoot the alien craft, dodge the rock, thread the gate at the end of
every sector. Clear sector 7 and the run is complete — then keep flying in OVERRUN if you want.
▌ SKYRUN · 40 ▰▰▱ S 1/7 ▸▹▹▹▹▹▹▹ 13% ×1
· ☩ ·
☩◄◄‹◦›►►☩ ··
▏ · ▕ ·
▏ · ▕ ·
▏ ▕
▏ ▲ ▕
◉███████████████████████████◉
←→ move · [space] shoot ☩ · dodge █ · grab ◈ THE 747 LAB
S 1/7 and the distance bar are on screen from the first frame, so you always know where
the run ends. Three shields, no lives — the only way out is to run out of hull.
The same fiction, the opposite camera. One button up, one button down: jump the solid
block, slam through the brittle dither. Seven gates, 7,470 metres , one number on the HUD.
Thrust is speed, so collecting it makes the run both faster and shorter.
▌ JETWASH · 741 ▣▣▣▢ G1 ▸▸▸▸▸▸▸▹▹▹ ▮▮▮▯▯▯▯▯▯▯
· · · -
· · · ·
▓▓▓▓ ▓▓▓▓▓
▓▓▓ ▓▓▓ ▓▓▓▓ ▓▓▓▓▓▓▓▓ ▓▓▓▓▓ ▓▓▓▓
════════════════════════════════════════════════════════════════════════════════
▛▜ ◈ ◈ ║║ ▛▜
║║
◈ ◈ ║║
► ║║ ◈ ▟▙
▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀
╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱
╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱ ╱
↑ jump · ↓ slam · █ solid · ▒ breaks THE 747 LAB
Every hazard is telegraphed more than a second before it reaches you, so a hit is always
a decision you got wrong — never a surprise.
Seven waves, then the big one. The top row is worth triple, so greed is a real choice.
Chain your kills inside 1.2 seconds to build a multiplier; bunkers arrive in wave 3,
divers in wave 5.
▌ ASTROS · 0 ▲▲ W1/7 ▰▰▰▰▰▰▰▰▰▰
▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼
▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼
▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼
▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼
▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼ ▼▼
!
← → move [space] fire dodge the !
▲
WAVE 1 · BEST 0 THE 747 LAB
The opening seconds teach the game by layout rather than by a tutorial box: one bomb
falls, alone, into empty space, so you learn what ! means before it can cost you.
Four lanes of traffic, three lanes of river, seven bays to fill. The road is a waiting
game — stand still and let the gap come to you. The river is the opposite, because water
is not a lane you can stand in: you ride whatever floats past, or you don't get across.
▒ ▒ ▽ ▽ ▽ ▒ ·▽· ▽ ▒ ▽▒ ▽
▄▄▄▄▄▄▄▄· ▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄▄ ·▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄▄ ·▄▄
▄▄▄▄▄ ▄▄▄▄▄▄▄▄ ~~▄▄▄▄▄▄▄▄~ ▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄▄ ·▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄
▄▄▄▄▄ ▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄▄ · ▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄
▄▄▄▄▄▄▄ · ▄▄▄▄▄▄▄▄ ·▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄▄~ ▄▄▄▄▄▄▄▄ ~▄▄▄▄▄▄▄▄ ~ ▄▄▄▄▄▄▄▄ ▄▄▄▄
▄▄ ▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄▄ · ▄▄▄▄▄▄▄▄ ·▄▄▄▄▄▄▄▄ ·▄▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄▄· ▄▄▄▄▄▄▄▄
▒▒ ▒ ▒ ▒ ▒ ▒ ▒ ▒ ▒ ▒ ▒ ▒ ▒▒ ▒
╌ ╌▬▶ ╌ ╌▬▶ ╌ ╌▬▶ ╌ ╌▬▶ ╌ ╌▬▶
╌ ◀▬╌ ╌ ◀▬╌ ╌ ◀▬╌ ╌ ◀▬╌ ╌ ◀▬╌
╌ ╌ ╌ ╌ ╌▌✷▐ ╌ ╌ ╌ ╌ ╌
╌ ╌◀██ ╌ ╌ ◀██ ╌ ╌ ◀██ ╌ ╌ ◀██ ╌ ╌ ◀█
▶ ╌ ╌ ▬▶ ╌ ╌ ▬▶ ╌ ╌ ▬▶ ╌ ╌ ▬▶ ╌ ╌ ▬
▒▒▒ ▒ ▒ ▒ ▒ ▒ ▒ ▒ ▒
JAYWALK 80 ◆◆ ▽▽▽▽▽▽▽ R1 [↑] cross [q] close
▽▽▽▽▽▽▽ is the whole scoreboard: seven empty bays at the start, and closing the last one
sends you back to the kerb with the traffic running faster. Nobody has ever had to be told
what to do in this one, which is exactly why it is here.
The clip at the top of this page is one JAYWALK run: a bay filled, a log ridden, and two of
the three lives spent getting there.
In any game — q closes it. The game pauses itself when Claude replies; space (or
p in SKYRUN) keeps it running anyway.
Slash commands open a free-play round any time, ignoring Claude's state:
/breakin · /skyrun · /jetwash · /astros · /jaywalk — and /breakout still works
as an alias for /breakin .
tmux (recommended) — the game opens as a split pane right in your window while Claude thinks, and disappears the moment Claude replies — your terminal goes back to exactly how it was, no tab to close. Your run is kept alive in the background and rejoins, mid-game, on your next prompt.
iTerm2 without tmux — the game splits your current iTerm window natively (same in-place feel).
Elsewhere — auto-open needs tmux. The games still run standalone any time: python3 games/jetwash.py --free .
Python 3.9+ with curses (standard on macOS and Linux). No pip install, ever.
Every title stays readable down to an 80×8 pane, falls back to 16 colours, and falls back
to pure ASCII on a terminal that can't do UTF-8.
A tiny file at ~/.747-terminal-games/mode controls auto-launch:
Whichever you last chose in the picker. It is stored as one word in
~/.747-terminal-games/game , so you can also set it directly:
echo jetwash > ~/.747-terminal-games/game
The slash commands ignore this file — they always open the title you named.
Before the picker comes up, a textmode fly-through: 7.47 seconds, and any key skips it
instantly.
Built by The 747 Lab . A growing line of terminal games.
There is a 747 hidden in every one of them. It is never on the label — it is in the
game. Find it; it does something.
MIT licensed. PRs and new games welcome — see CONTRIBUTING.md .
Arcade, reborn in your terminal — five games that open while Claude thinks and vanish when it replies. Developed by The 747 Lab.
github.com/The747Lab/747-terminal-games/releases/tag/v2.0.0 Topics
Readme MIT license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
