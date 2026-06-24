---
source: "https://github.com/wkswede/robot-deck-audio"
hn_url: "https://news.ycombinator.com/item?id=48659515"
title: "I built a local AI deck music assistant on a Mac mini"
article_title: "GitHub - wkswede/robot-deck-audio: A local-first home automation hub bridging Open WebUI, Flask, and AppleScript to orchestrate multi-room Airfoil audio and Chrome/YouTube playback. · GitHub"
author: "wkswede"
captured_at: "2026-06-24T13:42:46Z"
capture_tool: "hn-digest"
hn_id: 48659515
score: 1
comments: 0
posted_at: "2026-06-24T13:31:10Z"
tags:
  - hacker-news
  - translated
---

# I built a local AI deck music assistant on a Mac mini

- HN: [48659515](https://news.ycombinator.com/item?id=48659515)
- Source: [github.com](https://github.com/wkswede/robot-deck-audio)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T13:31:10Z

## Translation

タイトル: Mac mini でローカル AI デッキ音楽アシスタントを構築してみた
記事のタイトル: GitHub - wkswede/robot-deck-audio: Open WebUI、Flask、AppleScript をブリッジして、マルチルーム Airfoil オーディオと Chrome/YouTube 再生を調整するローカル ファーストのホーム オートメーション ハブ。 · GitHub
説明: Open WebUI、Flask、AppleScript を橋渡しして、マルチルーム Airfoil オーディオと Chrome/YouTube 再生を調整する、ローカル ファーストのホーム オートメーション ハブです。 - wkswede/robot-deck-audio

記事本文:
GitHub - wkswede/robot-deck-audio: Open WebUI、Flask、AppleScript をブリッジして、マルチルーム Airfoil オーディオと Chrome/YouTube 再生を調整するローカル ファーストのホーム オートメーション ハブ。 · GitHub
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
ウィクスヴェーデ
/
ロボットデッキオーディオ
公共
通知
Y

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
10 コミット 10 コミット README-assets: README-assets: README.md README.md all-off.sh all-off.sh app.py app.py backyard-chrome-on.sh backyard-chrome-on.sh backyard-living-room-on.sh backyard-living-room-on.sh backyard-off.sh backyard-off.sh next-track.sh next-track.sh play-deck-music.sh play-deck-music.sh play-pause-music.sh play-pause-music.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ロボット: ローカル AI デッキ音楽アシスタント
私は、Mac mini 上のローカル LLM を利用して、携帯電話、時計、またはラップトップからワンタップでデッキ音楽を再生したいと考えていました。
そこで私は、Open WebUI + Ollama + Flask + AppleScript + Airfoil + Chrome + Apple Shortcuts を組み合わせたホームラボ音楽アシスタント Robot を構築しました。
これで、「デッキ ミュージックを再生」と言うかタップすると、ロボットが次のようになります。
エアフォイルの裏庭スピーカーを接続します
Chrome で Deck Music YouTube プレイリストを開きます
挿入された JavaScript を介して YouTube の [すべて再生] ボタンをクリックします
携帯電話、時計、または Open WebUI からすべてを一時停止、スキップ、またはオフにできます
裏庭 + リビングルームのスピーカーをオンにする
面白いのは、同じシステムを制御する方法が複数あることです。
WebUI チャットを開く (「デッキ ミュージックを再生」)
iPhone / Apple Watch の Apple ショートカット
Mac から Robot を開くための Raycast ショートカット
テスト/デバッグのために Flask に直接 HTTP 呼び出しを行う
大まかに言うと、システムは次のようになります。
Apple ショートカット / オープン WebUI / レイキャスト
↓
フラスコ API
↓
シェルスクリプト + AppleScript
↓
エアフォイル + Chrome + YouTube
↓
裏庭/リビングルーム/デッキスピーカー
実際に
Open WebUI は Mac mini 上で実行され、play_deck_music() などのツール呼び出しを公開します。
これらのツール呼び出しは、ポート 5055 で実行されている Flask API にヒットします。
Flask ルートトリガーシェルスクリプト
シェルスクリプトはAppleを使用します

Airfoil と Google Chrome を制御するスクリプト
Chrome は YouTube プレイリストを開き、JavaScript を実行して [すべて再生] をクリックします
翼型ハンドルがオーディオを適切なスピーカーにルーティングします
現在の裏庭のセットアップには次のものが含まれます。
Mac mini (「ロボット」) — ローカル AI + 自動化の頭脳
軒下に設置されたKlipsch屋外スピーカー
再生ソースとしての Google Chrome
屋外オーディオセットアップの部品を収納するデッキボックス
iPhone / Apple Watch の Apple ショートカットによる素早いコントロール
デッキボックスの中には、オーディオ機器、ケーブル配線、電源、換気装置のコレクションが入っています。これは洗練された製品ではなく、本当に役立つようになった自宅の研究室/裏庭システムです。
重要なアイデア: AI は音楽を選択しません
私が気に入っているデザインの選択の 1 つは、ロボットが DJ になろうとしているわけではないということです。
私は今でも自分の携帯電話からプレイリストを自分で厳選しています。ロボットのジョブはアクティビティ層です。
一時停止/スキップ/すべてシャットダウン
これによりシステムがシンプルになり、信頼性が高まります。
ロボットに与えることができるコマンドの例:
「裏庭とリビングルームをつけてください」
現在、次のようなショートカットがあります。
Mac から Robot を即座に開く Raycast コマンドもあります。
微妙だが重要な教訓の 1 つは、Flask がどこに耳を傾けているかということでした。
元々は Flask を次のように実行していました。
アプリ。実行 (ホスト = "127.0.0.1"、ポート = 5055)
これは、Mac mini 自体と Docker で実行されている Open WebUI からは機能しましたが、iPhone Safari / Apple Shortcuts からは機能しませんでした。
アプリ 。実行 (ホスト = "0.0.0.0" 、ポート = 5055)
LAN 上のデバイスが次の API にアクセスできるようになりました。
http://192.168.x.x:5055
それが Apple ショートカットを機能させるための鍵でした。
リクエストの送信元に応じて次のようになります。
ロボット本体
http://127.0.0.1:5055
WebUIコンテナを開く
http://host.docker.internal:5055
iPhone / Apple ショートカット / その他の LAN デバイス
http://192.168.x.x:5055
それは結局b

これは、プロジェクトで最も役立つ「グルー コード」レッスンの 1 つです。
API 設計のレッスン: 長時間実行されるアクションはすぐに返される必要がある
一部のアクション、特に Play Deck Music は本質的に時間がかかります。
Apple ショートカットの場合、これらのエンドポイントがすぐに戻ってバックグラウンドで作業を続行できる方が良いでしょう。実際には、これは、長時間実行されるルートに対して、 subprocess.run() でブロックするのではなく、 subprocess.Popen() を使用することを意味します。
たとえば、play_pause ルートは次のようになります。
@アプリ。ルート ( "/music/play_pause" 、メソッド = [ "POST" ])
def play_pause():
サブプロセス 。ポペン（
[ スクリプト [ "play_pause_music" ]]、
stdout = サブプロセス。デブヌル、
stderr = サブプロセス。デブヌル
)
jsonify を返します ({
"アクション" : "音楽の再生一時停止" ,
"ステータス" : "開始済み"
})
ハッキングとは何か
これはまさに Mac mini + スクリプト + AppleScript + ローカル AI プロジェクトであり、洗練された製品ではありません
YouTube の自動化は本質的に少し脆弱です
ローカル モデルは、ツールについて話すのではなく、ツールを使用するよう促す必要がある場合があります。
デッキボックスには現実世界の妥協と「やりながら考え出す」エンジニアリングが数多く含まれています
でも、それが好きな理由でもあります。これはスマート ホーム プラットフォームを装っているのではなく、便利で非常に個人的なホームラボ システムです。
次に検討したいアイデアがいくつかあります。
「今何が再生されているの？」サポート
より堅牢なマルチゾーン プリセット
プロンプトやツールの説明が改善され、ローカル モデルが適切なツールをより確実に選択できるようになります。
実際に便利な機能を備えたコンピューターが家に欲しかったからです。
「インターネットのまとめ」ではありません。
「一般的な質問に答える」ものではありません。
ただ：
デッキの音楽をオンにして、裏庭に生き生きとした雰囲気を与えましょう
ロボットアクション/
§── 翼形/
│ §── backyard-chrome-on.sh
│ §── backyard-living-room-on.sh
│ §── backyard-off.sh

│ §── all-off.sh
│ §── play-deck-music.sh
│ §── play-pause-music.sh
│ └── next-track.sh
§── サーバー/
│ └─ airfoil_server.py
└── README-assets/
§── デッキ全体.jpg
└── デッキボックス内部.jpg
これを適応させる人へのメモ
同様のものを構築したい場合は、以下をお勧めします。
プレイリストの URL を環境変数または設定ファイルに保持します
高速コントロール (ショートカット) と柔軟なコントロール (チャット) を分離します。
すべてを AI で制御するのではなく、アクティビティを制御する
ネットワーキング、ブラウザの自動化、オーディオ ルーティングの間には、多少の接着作業が必要になることが予想されます
あなたがホームラボの人なら、とにかくそれは楽しみの半分です。
Open WebUI、Flask、AppleScript を橋渡しして、マルチルーム Airfoil オーディオと Chrome/YouTube 再生を調整する、ローカル ファーストのホーム オートメーション ハブです。
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

A local-first home automation hub bridging Open WebUI, Flask, and AppleScript to orchestrate multi-room Airfoil audio and Chrome/YouTube playback. - wkswede/robot-deck-audio

GitHub - wkswede/robot-deck-audio: A local-first home automation hub bridging Open WebUI, Flask, and AppleScript to orchestrate multi-room Airfoil audio and Chrome/YouTube playback. · GitHub
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
wkswede
/
robot-deck-audio
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
10 Commits 10 Commits README-assets: README-assets: README.md README.md all-off.sh all-off.sh app.py app.py backyard-chrome-on.sh backyard-chrome-on.sh backyard-living-room-on.sh backyard-living-room-on.sh backyard-off.sh backyard-off.sh next-track.sh next-track.sh play-deck-music.sh play-deck-music.sh play-pause-music.sh play-pause-music.sh View all files Repository files navigation
Robot: a local AI deck music assistant
I wanted one-tap deck music from my phone, watch, or laptop — powered by a local LLM on a Mac mini.
So I built Robot , a home-lab music assistant that combines Open WebUI + Ollama + Flask + AppleScript + Airfoil + Chrome + Apple Shortcuts .
Now I can say or tap “Play Deck Music” and Robot:
Connects my Backyard speakers in Airfoil
Opens my Deck Music YouTube playlist in Chrome
Clicks YouTube’s Play all button via injected JavaScript
Lets me pause, skip, or shut everything off from my phone, watch, or Open WebUI
Turn on Backyard + Living Room speakers
The fun part is that there are multiple ways to control the same system :
Open WebUI chat ( "Play deck music" )
Apple Shortcuts on iPhone / Apple Watch
Raycast shortcut to open Robot from Mac
direct HTTP calls to Flask for testing/debugging
At a high level, the system looks like this:
Apple Shortcut / Open WebUI / Raycast
↓
Flask API
↓
shell scripts + AppleScript
↓
Airfoil + Chrome + YouTube
↓
Backyard / Living Room / Deck speakers
In practice
Open WebUI runs on the Mac mini and exposes tool calls like play_deck_music()
those tool calls hit a Flask API running on port 5055
Flask routes trigger shell scripts
shell scripts use AppleScript to control Airfoil and Google Chrome
Chrome opens a YouTube playlist and executes JavaScript to click Play all
Airfoil handles routing the audio to the correct speakers
The current backyard setup includes:
Mac mini (“Robot”) — local AI + automation brain
Klipsch outdoor speakers mounted under the eaves
Google Chrome as the playback source
Deck box housing parts of the outdoor audio setup
Apple Shortcuts on iPhone / Apple Watch for quick controls
Inside the deck box I’ve got a collection of audio gear, cabling, power, and ventilation. It’s not a polished product build — it’s a home-lab / backyard system that became genuinely useful.
Key idea: AI doesn’t choose the music
One design choice I like is that Robot does not try to become a DJ.
I still curate the playlist myself from my phone. Robot’s job is the activity layer:
pause / skip / shut it all down
That keeps the system simple and makes it feel more reliable.
Examples of commands I can give Robot:
"Turn on the backyard and living room"
I currently have shortcuts like:
I also have a Raycast command that opens Robot instantly from my Mac.
One subtle but important lesson was where Flask is listening .
Originally I had Flask running like this:
app . run ( host = "127.0.0.1" , port = 5055 )
That worked from the Mac mini itself and from Open WebUI running in Docker, but did not work from iPhone Safari / Apple Shortcuts .
app . run ( host = "0.0.0.0" , port = 5055 )
allowed devices on my LAN to reach the API at:
http://192.168.x.x:5055
That was the key to getting Apple Shortcuts working.
Depending on where the request originates:
Robot itself
http://127.0.0.1:5055
Open WebUI container
http://host.docker.internal:5055
iPhone / Apple Shortcuts / other LAN devices
http://192.168.x.x:5055
That ended up being one of the most useful “glue code” lessons in the project.
API design lesson: long-running actions should return quickly
Some actions — especially Play Deck Music — are slow by nature:
For Apple Shortcuts, it’s better if those endpoints return immediately and let the work continue in the background. In practice, that means using subprocess.Popen() for long-running routes rather than blocking on subprocess.run() .
For example, the play_pause route looks like this:
@ app . route ( "/music/play_pause" , methods = [ "POST" ])
def play_pause ():
subprocess . Popen (
[ SCRIPTS [ "play_pause_music" ]],
stdout = subprocess . DEVNULL ,
stderr = subprocess . DEVNULL
)
return jsonify ({
"action" : "play_pause_music" ,
"status" : "started"
})
What’s hacky
This is very much a Mac mini + scripts + AppleScript + local AI project, not a polished product
YouTube automation is inherently a little brittle
local models sometimes need a nudge to use the tool instead of talking about the tool
the deck box contains a bunch of real-world compromise and “figure it out as you go” engineering
But that’s also why I like it. It’s not pretending to be a smart home platform — it’s a useful, very personal home-lab system.
A few ideas I want to explore next:
better “what’s currently playing?” support
more robust multi-zone presets
better prompting / tool descriptions so the local model chooses the right tool more reliably
Because I wanted a computer in my house that does something actually useful.
Not “summarize the internet.”
Not “answer generic questions.”
Just:
turn on the deck music and make the backyard feel alive
robot-actions/
├── airfoil/
│ ├── backyard-chrome-on.sh
│ ├── backyard-living-room-on.sh
│ ├── backyard-off.sh
│ ├── all-off.sh
│ ├── play-deck-music.sh
│ ├── play-pause-music.sh
│ └── next-track.sh
├── server/
│ └── airfoil_server.py
└── README-assets/
├── deck-wide.jpg
└── deck-box-inside.jpg
Notes for anyone adapting this
If you want to build something similar, I’d recommend:
keep your playlist URL in an environment variable or config file
separate fast controls (Shortcuts) from flexible controls (chat)
make the AI control activities , not everything
expect a bit of glue work between networking, browser automation, and audio routing
If you’re a home-lab person, that’s half the fun anyway.
A local-first home automation hub bridging Open WebUI, Flask, and AppleScript to orchestrate multi-room Airfoil audio and Chrome/YouTube playback.
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
