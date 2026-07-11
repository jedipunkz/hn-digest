---
source: "https://github.com/logxio/picchio"
hn_url: "https://news.ycombinator.com/item?id=48874905"
title: "Show HN: Catch your local LLM falling back to CPU"
article_title: "GitHub - logxio/picchio: Catches silent CPU fallback and mislabeled tok/s in local LLMs. llama.cpp and ollama, one file, no deps. · GitHub"
author: "logickkk1"
captured_at: "2026-07-11T19:59:54Z"
capture_tool: "hn-digest"
hn_id: 48874905
score: 2
comments: 0
posted_at: "2026-07-11T19:17:35Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Catch your local LLM falling back to CPU

- HN: [48874905](https://news.ycombinator.com/item?id=48874905)
- Source: [github.com](https://github.com/logxio/picchio)
- Score: 2
- Comments: 0
- Posted: 2026-07-11T19:17:35Z

## Translation

タイトル: HN を表示: CPU にフォールバックするローカル LLM をキャッチ
記事のタイトル: GitHub - logxio/picchio: ローカル LLM でサイレント CPU フォールバックと誤ってラベル付けされた tok/s を検出します。 llama.cpp と ollam、1 つのファイル、deps なし。 · GitHub
説明: ローカル LLM でサイレント CPU フォールバックと誤ってラベル付けされた tok/s を捕捉します。 llama.cpp と ollam、1 つのファイル、deps なし。 - ログシオ/ピッキオ

記事本文:
GitHub - logxio/picchio: ローカル LLM でサイレント CPU フォールバックと誤ってラベル付けされた tok/s をキャッチします。 llama.cpp と ollam、1 つのファイル、deps なし。 · GitHub
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
ログシオ
/
ピッキオ
公共
通知
通知設定を変更するにはサインインする必要があります
さらに

l ナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
22 コミット 22 コミット .github .github アセット アセットの例 例 .gitignore .gitignore ライセンス ライセンス README.md README.md picchio.py picchio.py すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Picchio はイタリア語でキツツキを意味します。ノックオンする 1 つの Python ファイルです。
ローカルの LLM セットアップを実行し、空洞スポットをリッスンします。どのタスクが実行したか
実際に GPU が実際に動作したのでしょうか?
インストール · チェック内容 · 例
そのブロックが製品全体です。決して結合されることのない 3 つのレーンです。
エンジン自体のログからの配置の証拠、冷却の故障
可決、そして評決。 15 行なので、意図的に狭くしてあります。
コメントスレッドに貼り付けられたままでも残ります。長いモデル ファイル名は、
最初の行をさらに広く押します。実際の出力、未編集
(例/healthy-metal.txt);テキスト
以下のバージョンは貼り付けたものです。
モデル Qwen3.5-9B-Q4_K_M.gguf、8.95 B、5.28 GiB、llama.cpp b9430
GPU 使用中: GPU 上の 33/33 レイヤー (メタル: Apple M5)
プレフィルデコードウォールクロック
コールド 597.3 トーク/秒 17.2 トーク/秒 11.4 トーク/秒
暖かい中頃 558.9 tok/s 20.0 tok/s 14.4 tok/s
ウォームスパン 522~596 18.8~21.2 13.2~15.6
コールド パスがどこに移動したか (11.1 秒、4/10 スレッド)
負荷重量 1.8 秒 #####................................................ 16%
プレフィル 1.3 秒 ###................................. 11%
デコード 7.4 秒 ###################...... 66%
エンジンその他 0.7 秒 ##................................ 6%
評決: 健康。 GPU がその仕事をしてくれました。暖かい中央値を引用する
デコード: 20.0 トーク/秒。 559 tok/s はプレフィルです。読み取り速度ではなく、読み取り速度です。
書くこと。
-- picchio v0.1.0 mp1 (Apple M5、32 GB、macOS 26.5.1)
実行してみよう
カール -LO https://raw.githubusercontent.com/snxoi/picchio/main/picchio.py
python3 picchio.py
2 行目は引数なしで、マシンの周囲を調べます。
(ollam タグ、

現在のフォルダー、HF および LM Studio キャッシュ)、および
そのままコピーできるコマンドを出力します。またはモデルに向けてください
自分自身: .gguf パスは完全な llama.cpp 診断を実行します。
タグは測定モードを実行します。
python3 picchio.py /path/to/model.gguf
python3 picchio.py qwen3.5:9b
pip、依存関係、設定はありません。 1 つの Python ファイル、980 行、
標準ライブラリのみ。 python3 と llama.cpp または ollam のいずれかを持っている場合は、
必要なものはすでにすべて揃っています。モデルを 3 回実行します。
固定プロンプト (最初のパスは冷たく、残りは温かいもの)、
エンジン自体の番号を取得し、上記の判定ブロックを出力します。
実行コスト: 3 パス、このマシンで約 1 分
GPU が作動しており、CPU では数分かかります。重いものを読んでください。モデルファイルを読み込みます。
~/.cache/picchio の下に 1 つの小さなキャッシュ ファイルを書き込み、何も変更しません。
その後、プロセスは実行されません。
再実行するタイミング: llama.cpp または ollama のアップグレード後、OS のアップグレード後
同じモデルのクオントを切り替えた後、-ngl をタッチした後、更新します。
またはコンテキスト サイズ、そして tok/s 番号をどこかに投稿する前に 1 回。
パーサーは、このリポジトリ内の生のログによって固定されます。ランニング
python3 picchio.py --selftest は、編集されていないエンジン出力を再生します。
example/raw/ であり、コミットされたものをすべて再現する必要があります
ブロックを行ごとに判定します。現時点ではパスフィクスチャが 9 つとパスフィクスチャが 3 つです
ブロックすると、上のバッジはプッシュするたびに正確に実行されます。
GPU はプリフィルでは 22 倍を購入しましたが、デコードでは 2 倍しか購入しなかったので、この数字は
投稿者全員がボトルネックに気づいていません。 36 トーク/秒の読み取り I
信頼できるものは 32 個のマトリックス セルのうち 0 個で再現されます。バックグラウンドでのダウンロード
デコードをほぼ半分に削減するため、アイドル状態でベンチマークを実行します。
機械。以下はこれら 3 つを開梱したものです。
すべてのトーク/秒の数字は 3 つのレーンの 1 つに属し、ピッキオは決してありません
3 つのレーンを 1 つの番号にマージします。
プレフィルは

モデルがプロンプトを読み取る速度。デコードの速さ
答えを書きます。ウォールクロックは生成されたトークンをすべてで割ったもので、
負荷とウォームアップが含まれており、これがストップウォッチと腸の機能です
測る。暖かい中央値の上のマシンでは、559、20.0、そして
14.4.以下の CPU 実行では、それらは 26、11、および 3 です。ラベルのない単一の
30倍の範囲にわたる数値は測定値ではなく、噂です。
スクリーンショットで Mac が 500 tok/s を実行していることが示されている場合、それはほとんどの場合、
プレフィル。 llama-bench が tg128 を出力すると、それはデコードされます。アプリのとき
最初の単語が表示される前に遅く感じる、つまりコールドロードに加えて
プレフィルであり、デコード番号がないと説明されません。次の投稿をする前に
tok/s 番号。これを実行してどのレーンを確認するのに 1 分かかります
に住んでいます。
みんなが投稿した数字からはボトルネックが見えない
ここで測定すると、GPU はデコードで約 2 倍、プリフィルで約 22 倍のパフォーマンスを発揮します。
(どちらの実行も example/ にあり、CPU 上の 10 CPU スレッドのうち 4 つ)
側面）。オンラインに投稿されたほぼすべての tok/s 数値はデコードされています。
タイピングの速さを感じるものです。しかし、コンシューマ向けハードウェアでは、
痛みは主にプレフィルに存在します。プロンプトがどのくらい長く続くかを決定します。
最初の言葉の前に沈黙する。 2 つのセットアップで同じデコード番号をポストできる
一方、答え始めるまでに 10 倍の時間がかかります。もしあなたのエンジンが
静かに CPU に戻りました、デコードは数字ではわかりません
あなた。
穴場: サイレント CPU フォールバック
同じマシン、同じモデル、同じファイル、CPU に強制使用
(例/cpu-fallback.txt):
モデル Qwen3.5-9B-Q4_K_M.gguf、8.95 B、5.28 GiB、llama.cpp b9430
GPU が使用されていません: GPU 上の 0/33 レイヤー
プレフィルデコードウォールクロック
コールド 24.3 トーク/秒 9.3 トーク/秒 2.7 トーク/秒
暖かい中頃 25.7 tok/s 11.2 tok/s 2.9 tok/s
ウォームスパン 26~26 11.0~11.4 2.9~2.9
コールド パスがどこにいったか (47.9 秒、4/10 スレッド、キャッシュされた重み)
負荷をかける

重み 2.0 秒 #................................................ 4%
プレフィル 31.4 秒 ##################....... 65%
デコード 13.7 秒 ########................................ 29%
エンジンその他 0.9 秒 #................................. 2%
評決: サイレント CPU フォールバック。 33 レイヤーのうち 0 レイヤーが GPU に到達しました。
Decode (11.2) は問題ないように見えますが、これが隠れている方法です。プレフィル
26 トークン/秒でパークすると、2500 トークンがプロンプトで 97 秒アウトします。 -nglをチェックしてください。
-- picchio v0.1.0 mp1 (Apple M5、32 GB、macOS 26.5.1)
何が動いたのか、何が動かなかったのかを見てください。デコードが 2 回ドロップしました (チャット中)
肩をすくめるかも知れません。プレフィルは 22 倍に減少し、ロングワードの最初のワードが減少しました
プロンプトには数分かかるようになりました。ピッキオはこれを 2 つの方向から呼び出します。
1 回: エンジン自身のレイヤー配置ログ (0/33 オフロード) と
事前入力署名。この判決はどの Apple Silicon でも再現できます
マシン:
python3 picchio.py model.gguf -- --device none -ngl 0
bare -- 以降はすべて、llama.cpp バイナリに直接送られます。
ピッキオが行う 3 番目のことは、あなたに代わって番号を尋ねることです。誰か
tok/s の図を投稿するか、覚えていてそれが何であるかを知りたいとします
おそらく次のとおりでした:
python3 picchio.py --explain 36
あなたの数値: 36.0 tok/s -> ここで測定されたものと一致しません
36.0 tok/s は、ここで測定された値の 30% 以内ではありません
(最も近い: デコード、1.8x オフ; 測定: プレフィル 558.9、デコード
20.0、壁時計 14.4 トク/秒）。その番号を信用する前に聞いてください
それは 3 つのレートのどれでしたか、どのハードウェアでのクオンツでしたか、
そしてコンテキストの長さ。
(レート: Qwen3.5-9B-Q4_K_M.gguf、Apple M5、32 GB、2026-07-11)
この 36 は、このリポジトリが存在する正確な数です (ストーリー
は以下の「これを書いた理由」にあります)、マシンに対して尋ねられました。
から来たものと思われる。この短いチェックはそれ自体の出力です。
意図的に評決ブロックではありません: picchio はあなたのレートをキャッシュします。
最後の診断が実行されたため、チェックを再実行する必要はありません。 --expla を渡す

で
代わりにモデルパスを使用し、同じセクションが追加されます
完全な判定ブロックの下では、両方に 1 回の実行。
ファイルパスの代わりに ollam モデルタグを picchio に与えると、
同じものがローカルの ollam サーバーを通過します (デフォルト
127.0.0.1:11434 、または OLLAMA_HOST を設定します)。同じ 3 つのレーンがあり、
同じコールド パス ブレークダウンとメモリに基づく配置チェック
Split ollam 自体がレポートします: モデルのどのくらいが GPU メモリに存在するか
CPU メモリとの比較。
実際の実行、同じウェイトを ollam にインポート
(examples/ollama-qwen35.txt):
モデル qwen3.5:9b、9.0 B、Q4_K_M、5.55 GiB、ollam 0.31.1
GPU ENGAGED: GPU メモリの重みの 100% (ollam ps)
プレフィルデコードウォールクロック
コールド 539.2 トーク/秒 19.3 トーク/秒 12.2 トーク/秒
暖かい中頃 853.5 tok/s 19.9 tok/s 17.1 tok/s
ウォームスパン 847～860 19.5～20.4 16.8～17.4
コールドパスがどこへ行ったのか(10.5秒)
負荷重量 2.5 秒 #######................................ 23%
プレフィル 1.4 秒 ####................................. 14%
デコード 6.6 秒 ##################....... 63%
エンジンその他 0.0 秒 ................................... 0%
評決: 健康。 Ollama は GPU メモリの重みを 100% 報告します。
ウォーム デコードの中央値を引用します: 19.9 tok/s。 853トーク/秒は
プレフィル: 書き込みではなく読み取り。
-- picchio v0.1.0 mp1 (Apple M5、32 GB、macOS 26.5.1)
ollam は公開しないため、このモードでは表示できないものに注意してください。
それ: レイヤーごとの配置、デバイスの初期化ログ、およびスレッド構成。
そのため、llama.cpp モードが完全な診断であり、ollam モードが
測定と配置チェック。 ollam が次の時点でメモリ分割を行わない場合
すべて、ピッキオは配置を推測ではなく不明として報告します。そして
報告された分割自体が間違っている可能性があるため、ピッキオはそれをクロスチェックします
測定されたレートに対して: ollam が完全な GPU 配置を主張しているが、
プリフィルとデコードの比率は CPU に似ているように見え、判定はダウングレード
t

o 健康ではなく矛盾する証拠。
これら 2 つのモードがすべての範囲です。 picchio は 1 つの読み取り可能なファイルのままです。
picchio MODEL [フラグ] [-- エンジン引数]
MODEL .gguf パス (llama.cpp) または ollam モデル タグ。
引数なしで、検索できる実行可能なモデルをリストします。
--passes N 回の測定パス、最初の 1 回のコールド (デフォルトは 3)
--explain TOKS は、測定されたレーンに対して見た数値を分類します。
--keep-logs DIR 各パスの生のエンジン出力を DIR に保存します
--json ブロック後の機械可読測定値
--bin PATH 自分で llama.cpp バイナリを選択します
--selftest リプレイ例/生、コミットされた判定が再現されることを確認する
--version バージョンと測定プロトコルを表示します
bare -- 以降はすべて、llama.cpp バイナリに直接送られます。
色は端末上にのみ表示されます (NO_COLOR が尊重されます)。パイプ付き
出力は常にプレーン ASCII であるため、貼り付けられたブロックはバイト単位で表示されます。
セルフテストで検証されます。
私はアプリのローカル モデルを体系的に測定していました。
数週間かけて構築していたとき、私は自分のコードに対してバグを報告しそうになりました。
Bare llama.cpp では 36 tok/s でした。同じモデルがアプリを通じて提供されました
11.5。同じマシン、同じ日、そして 3x はあなたにとっての数字のようなものです
1週間を中心に再構成します。
修正を書く前に、両方のシステムを再実行しました

[切り捨てられた]

## Original Extract

Catches silent CPU fallback and mislabeled tok/s in local LLMs. llama.cpp and ollama, one file, no deps. - logxio/picchio

GitHub - logxio/picchio: Catches silent CPU fallback and mislabeled tok/s in local LLMs. llama.cpp and ollama, one file, no deps. · GitHub
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
logxio
/
picchio
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
22 Commits 22 Commits .github .github assets assets examples examples .gitignore .gitignore LICENSE LICENSE README.md README.md picchio.py picchio.py View all files Repository files navigation
Picchio is Italian for woodpecker: one Python file that knocks on
your local LLM setup and listens for hollow spots. Which tok/s did
you actually get, and did the GPU really do the work?
Install · What it checks · Examples
That block is the whole product: three lanes that never get merged,
placement evidence from the engine's own logs, a breakdown of the cold
pass, and a verdict. It is 15 lines, kept narrow on purpose so it
survives being pasted into a comment thread; a long model filename can
push the first line wider. Real output, unedited
( examples/healthy-metal.txt ); the text
version below is the one you paste:
model Qwen3.5-9B-Q4_K_M.gguf, 8.95 B, 5.28 GiB, llama.cpp b9430
gpu ENGAGED: 33/33 layers on GPU (Metal: Apple M5)
prefill decode wallclock
cold 597.3 tok/s 17.2 tok/s 11.4 tok/s
warm mid 558.9 tok/s 20.0 tok/s 14.4 tok/s
warm span 522~596 18.8~21.2 13.2~15.6
where the cold pass went (11.1 s, 4/10 threads)
load weights 1.8 s #####....................... 16%
prefill 1.3 s ###......................... 11%
decode 7.4 s ###################......... 66%
engine misc 0.7 s ##.......................... 6%
VERDICT: HEALTHY. The GPU did the work. Quote the warm median
decode: 20.0 tok/s. 559 tok/s is prefill: reading speed, not
writing.
-- picchio v0.1.0 mp1 on Apple M5, 32 GB, macOS 26.5.1
Get it running
curl -LO https://raw.githubusercontent.com/snxoi/picchio/main/picchio.py
python3 picchio.py
That second line, with no arguments, looks around your machine
(ollama tags, the current folder, the HF and LM Studio caches) and
prints commands you can copy as they are. Or point it at a model
yourself: a .gguf path runs the full llama.cpp diagnosis, an ollama
tag runs measurement mode.
python3 picchio.py /path/to/model.gguf
python3 picchio.py qwen3.5:9b
No pip, no dependencies, no config. One Python file, 980 lines,
stdlib only. If you have python3 plus either llama.cpp or ollama, you
already have everything it needs. It runs your model three times with
a fixed prompt (the first pass cold, the rest warm), reads the
engine's own numbers, and prints the verdict block above.
Cost of a run: three passes, about a minute on this machine with the
GPU engaged, a few minutes on CPU. Read heavy; it reads your model file,
writes one small cache file under ~/.cache/picchio , modifies nothing,
and leaves no process running afterwards.
When to rerun it: after a llama.cpp or ollama upgrade, after an OS
update, after switching quants of the same model, after touching -ngl
or context size, and once before you post a tok/s number anywhere.
The parser is pinned by the raw logs in this repo. Running
python3 picchio.py --selftest replays the unedited engine output in
examples/raw/ and has to reproduce every committed
verdict block line for line; right now that is 9 pass fixtures and 3
blocks, and the badge above runs exactly that on every push.
The GPU bought 22x on prefill but only 2x on decode, so the number
everyone posts is blind to the bottleneck. A 36 tok/s reading I
trusted reproduced in zero of 32 matrix cells. A background download
cut decode nearly in half, which is why you benchmark on an idle
machine. Everything below unpacks these three.
Every tok/s figure belongs to one of three lanes, and picchio never
merges the three lanes into one number.
Prefill is how fast the model reads your prompt. Decode is how fast it
writes the answer. Wallclock is generated tokens divided by everything,
load and warmup included, which is what your stopwatch and your gut
measure. On the machine above the warm medians land at 559, 20.0 and
14.4. On the CPU run below they are 26, 11 and 3. A single unlabeled
number spanning a 30x range is not a measurement, it is a rumor.
When a screenshot shows a Mac doing 500 tok/s, that is almost always
prefill. When llama-bench prints tg128, that is decode. When an app
feels slow before the first word appears, that is cold load plus
prefill, and no decode number will explain it. Before you post your next
tok/s number, it costs one minute to run this and find out which lane it
lives in.
The number everyone posts cannot see your bottleneck
Measured here, the GPU buys about 2x on decode and about 22x on prefill
(both runs are in examples/ , 4 of 10 cpu threads on the CPU
side). Nearly every tok/s figure posted online is decode, because decode
is the one that feels like typing speed. But on consumer hardware the
pain lives mostly in prefill: it decides how long a long prompt sits
silent before the first word. Two setups can post the same decode number
while one takes ten times longer to start answering. And if your engine
quietly fell back to CPU, decode is the number that will not tell
you.
The hollow spot: silent CPU fallback
Same machine, same model, same file, forced to CPU
( examples/cpu-fallback.txt ):
model Qwen3.5-9B-Q4_K_M.gguf, 8.95 B, 5.28 GiB, llama.cpp b9430
gpu NOT ENGAGED: 0/33 layers on GPU
prefill decode wallclock
cold 24.3 tok/s 9.3 tok/s 2.7 tok/s
warm mid 25.7 tok/s 11.2 tok/s 2.9 tok/s
warm span 26~26 11.0~11.4 2.9~2.9
where the cold pass went (47.9 s, 4/10 threads, weights cached)
load weights 2.0 s #........................... 4%
prefill 31.4 s ##################.......... 65%
decode 13.7 s ########.................... 29%
engine misc 0.9 s #........................... 2%
VERDICT: SILENT CPU FALLBACK. 0 of 33 layers reached the GPU.
Decode (11.2) looks passable, which is how this hides. Prefill
at 26 tok/s parks a 2500 token prompt 97 s out. Check -ngl.
-- picchio v0.1.0 mp1 on Apple M5, 32 GB, macOS 26.5.1
Look at what moved and what did not. Decode dropped 2x, which in a chat
you might shrug at. Prefill dropped 22x, and the first word of a long
prompt now takes minutes. picchio calls this from two directions at
once: the engine's own layer placement log (0/33 offloaded) and the
prefill signature. You can reproduce this verdict on any Apple Silicon
machine with:
python3 picchio.py model.gguf -- --device none -ngl 0
Anything after the bare -- goes straight to the llama.cpp binary.
The third thing picchio does is interrogate a number for you. Someone
posts a tok/s figure, or you remember one, and you want to know what it
probably was:
python3 picchio.py --explain 36
YOUR NUMBER: 36.0 tok/s -> MATCHES NOTHING MEASURED HERE
36.0 tok/s is not within 30% of anything measured here
(closest: decode, off by 1.8x; measured: prefill 558.9, decode
20.0, wallclock 14.4 tok/s). Before trusting that number, ask
which of the three rates it was, and on what hardware, quant,
and context length.
(rates: Qwen3.5-9B-Q4_K_M.gguf, Apple M5, 32 GB, 2026-07-11)
That 36 is the exact number this repo exists because of (the story
is in "Why I wrote this" below), asked against the machine it
supposedly came from. This short check is its own output,
deliberately not a verdict block: picchio caches the rates from your
last diagnostic run, so the check needs no rerun. Pass --explain
together with a model path instead and the same section is appended
under a full verdict block, one run for both.
Give picchio an ollama model tag instead of a file path and it runs the
same passes through your local ollama server (default
127.0.0.1:11434 , or set OLLAMA_HOST ). You get the same three lanes,
the same cold pass breakdown, and a placement check based on the memory
split ollama itself reports: how much of the model sits in GPU memory
versus CPU memory.
Real run, same weights imported into ollama
( examples/ollama-qwen35.txt ):
model qwen3.5:9b, 9.0 B, Q4_K_M, 5.55 GiB, ollama 0.31.1
gpu ENGAGED: 100% of weights in GPU memory (ollama ps)
prefill decode wallclock
cold 539.2 tok/s 19.3 tok/s 12.2 tok/s
warm mid 853.5 tok/s 19.9 tok/s 17.1 tok/s
warm span 847~860 19.5~20.4 16.8~17.4
where the cold pass went (10.5 s)
load weights 2.5 s #######..................... 23%
prefill 1.4 s ####........................ 14%
decode 6.6 s ##################.......... 63%
engine misc 0.0 s ............................ 0%
VERDICT: HEALTHY. Ollama reports 100% of weights in GPU memory.
Quote the warm median decode: 19.9 tok/s. 853 tok/s is
prefill: reading, not writing.
-- picchio v0.1.0 mp1 on Apple M5, 32 GB, macOS 26.5.1
Be aware of what this mode cannot see, because ollama does not expose
it: per layer placement, device init logs, and thread configuration.
That is why llama.cpp mode is the full diagnosis and ollama mode is
measurement plus a placement check. If ollama gives no memory split at
all, picchio reports the placement as unknown instead of guessing. And
because a reported split can itself be wrong, picchio cross checks it
against the measured rates: when ollama claims full GPU placement but
the prefill to decode ratio looks CPU shaped, the verdict downgrades
to CONFLICTING EVIDENCE instead of HEALTHY.
These two modes are the whole scope; picchio stays one readable file.
picchio MODEL [flags] [-- engine args]
MODEL a .gguf path (llama.cpp) or an ollama model tag;
with no arguments, lists runnable models it can find
--passes N measurement passes, first one cold (default 3)
--explain TOKS classify a number you saw against the measured lanes
--keep-logs DIR save each pass's raw engine output into DIR
--json machine readable measurements after the block
--bin PATH choose the llama.cpp binary yourself
--selftest replay examples/raw, verify committed verdicts reproduce
--version print version and measurement protocol
Anything after a bare -- goes straight to the llama.cpp binary.
Color appears only on a terminal ( NO_COLOR is respected); piped
output is always plain ASCII, so a pasted block is byte for byte what
the selftest verifies.
I had been systematically measuring local models for an app I am
building, weeks of it, when I nearly filed a bug against my own code.
Bare llama.cpp gave me 36 tok/s. The same model through the app gave
11.5. Same machine, same day, and 3x is the kind of number you
reorganize a week around.
Before writing the fix I reran both si

[truncated]
