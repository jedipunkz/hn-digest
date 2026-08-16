---
source: "https://github.com/sib-project/sib"
hn_url: "https://news.ycombinator.com/item?id=49321017"
title: "Show HN: Unixy LLM Client using Git to store converastions, instead SQLite"
article_title: "GitHub - sib-project/sib: A standard Unix LLM client · GitHub"
author: "hskimse"
captured_at: "2026-08-16T16:13:14Z"
capture_tool: "hn-digest"
hn_id: 49321017
score: 2
comments: 0
posted_at: "2026-08-16T15:35:00Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Unixy LLM Client using Git to store converastions, instead SQLite

- HN: [49321017](https://news.ycombinator.com/item?id=49321017)
- Source: [github.com](https://github.com/sib-project/sib)
- Score: 2
- Comments: 0
- Posted: 2026-08-16T15:35:00Z

## Translation

タイトル: HN を表示: SQLite ではなく Git を使用して会話を保存する Unixy LLM クライアント
記事のタイトル: GitHub - sib-project/sib: 標準 Unix LLM クライアント · GitHub
説明: 標準の Unix LLM クライアント。 GitHub でアカウントを作成して、兄弟プロジェクト/兄弟開発に貢献します。
HN テキスト: こんにちは皆さん :) 私が git が好きな理由は、ソース ツリーのスナップショットとコミット構造自体が常に不変であり、破壊的な操作は ref ファイルの名前を変更するだけだからです。それを理解すれば、CLI がどれほど理解するのが難しくても、コマンドを実行することに躊躇することはありません。いつでも取り戻すことができます。 LLM 会話はソース ツリーほど変更しにくいものではありませんが、私にとっては、自動生成されたセッション タイトルよりも自動生成された SHA-1 ハッシュの方が安心です ^~^ プロジェクトは初期段階にあり、貢献は歓迎されています。 git 配管コマンドを使用したことがあれば、ハッキングは簡単ですし、かなり楽しいものになると思います。オリジナル: https://github.com/sib-project/sib/blob/master/Documentation...

記事本文:
GitHub - sib-project/sib: 標準 Unix LLM クライアント · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
兄弟プロジェクト
/
兄弟
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
131 コミット 131 コミット .github/ workflows .github/ workflows ドキュメント ドキュメント plm plm スクリプト scripts .gitignore .gitignore COTRIBUTING.md COTRIBUTING.md COPYING COPYING

Makefile Makefile README.md README.md VERSION VERSION lib.bash lib.bash sib sib sib-ask sib-ask sib-edit sib-edit sib-log sib-log sib-save sib-save sib-show sib-show sib-status sib-status sib-switch sib-switch test.sh test.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Sib: 標準 Unix LLM クライアント
Git が好きすぎるので、Sib は Git を使用して LLM 会話を保存します
SQLite の代わりにそれ自体。ストアはプレーンな Git リポジトリであり、
ユーザーとアシスタントの各ターンはコミットです。
コンテキストを完全に制御します。 git と同じように。
バックアップと共有は無料です。ストアは通常の Git です
リポジトリなので、すべての Git ホストがリモートとして機能します。
説明 - チュートリアルによるコンセプトの紹介
クイックスタート - インストールとセットアップ
ハッキング - 好きなように Sib を拡張する
echo 「なぜエヴァンゲリオン 3.0 は突然軌道から外れてしまったのでしょうか?」 |兄弟は尋ねます
兄弟ログ
echo 「これら 2 つの理由のうち、どちらがより重要ですか?」 |兄弟は尋ねます
sib ask は標準入力からプロンプトを読み取り、それをモデルに送信します。
HEAD から到達可能なターンごとに応答を stdout に出力し、
両方のターンを 2 つのコミットとして記録します。 sib log は全体を出力します
会話。
さて、店内はこんな感じです。
* fa754c HEAD~3 {"role":"user", "content":"なぜ突然エヴァンゲリオン3.0が…"}
* b16df7 HEAD~2 {"role":"アシスタント", "content":"なぜなら *エヴァンゲリオン3.0:あなたは..."}
* 1c2e90 HEAD~1 {"role":"user", "content":"これら 2 つの理由のうちどちらが重要ですか ..."}
* 280f09 HEAD {"role":"アシスタント", "content":"**意図的なクリエイティブ ..."}
これが全体的なコンセプトです。しかし、なぜ？不必要に見える
複雑です。
次のようなコマンドを見てみましょう。
sib ask -rp HEAD~1
各フラグの機能:
-r : 標準入力を読み取りません。チェーンをそのままエンドポイントに送信します
-p HEAD~1 : HEAD ではなく HEAD~1 を親として扱います
最後のプロンプトは HEAD~1 にあるため、これは
「繰り返し」ボタンが表示されました

ウェブ UI で。
echo 「つまり、14 年間のスキップは良い決断だったということですか?」 | sib ask -p HEAD~2
これは「編集」ボタンと似ています。あたかも HEAD~2 からフォークします。
Web UI で 1c2e90 を編集しました。
* fa754c HEAD~3 {"role":"user", "content":"なぜ突然エヴァンゲリオン3.0が…"}
* b16df7 HEAD~2 {"role":"アシスタント", "content":"なぜなら *エヴァンゲリオン3.0:あなたは..."}
* 32e3f4 HEAD~1 {"role":"user", "content":"つまり、14 年間のスキップは ..."}
* 8cc404 HEAD {"role":"アシスタント", "content":"芸術的にはそうです; 物語的には ..."}
ハッシュは覚えにくいため、名前を付けます。
sib 保存 eva-3.0-go-off
sib save -l # 保存された名前をリストする
保存された名前は、-p (および次のような他のコマンド) でアドレス指定できます。
兄弟ログ ...):
sib git fetch https://github.com/sib-project/hub \
dilluti0n/why-evangelion-end-like-that:refs/conv/why-evangelion-end-like-that
兄弟ログ なぜエヴァンゲリオンの終わりのようなもの
エコー「自分自身を受け入れることを学んでいますか？」正直に言って、彼は本当に成功しましたか？ \
「なぜ彼は祝福されているのですか？」 | sib ask -p なぜエヴァンゲリオンの終わりがそのようになるのか
各保存はデフォルトで所定の位置に留まり、 sib ask に従いません。結果が元の名前と一致し、気に入った場合は、保存します。
もう一度。
兄弟はなぜエヴァンゲリオンのような終わりを保存しますか
もちろん、エヴァンゲリオン 3.0 チェーンはまだ存在します。
シブログ eva-3.0-go-off
sib ask -p を使用すると、再度取得できます。でも時々あなただけが
HEADを移動したい：
sib ask -crp eva-3.0-go-off
-c はここでの新しい機能です。API 呼び出しがないため、アシスタントのターンはありません。 -r を使用すると、
入力も何もないので、残っているのは HEAD を移動することだけです。
これらは、Sib のほぼすべてのユースケースをカバーします。
会話がどのように保存されるかを知ることで、繰り返し、編集、フォーク、
会話の保存と会話の切り替え。バックアップと
共有は同じ場所から生まれます。その他のツールにはそれぞれが同梱されています
それらは別の名前付き機能として扱われます。兄弟は持っていない

ああ。
ここには、Sib を使用するために耐えなければならない避けられない定型文があります。
依存関係は bash >= 3.2、git、jq、curl、
awk と coreutils。
ダウンロードリリース
tarball またはクローン
このリポジトリ:
git clone https://github.com/sib-project/sib
グローバルインストール:
sudo メイクインストール
ユーザーによるインストール:
PREFIX="$HOME/.local" をインストールする
# macOS では、この行を .zshrc に追加してターミナルを再実行します
エクスポート PATH="$HOME/.local/bin:$PATH"
リポジトリを初期化します。
$兄弟ステータス
SIB_DIR: /home/hskim/.sib
PLM_MODEL: gpt-5.6-luna
PLM_ENDPOINT: https://api.openai.com/v1/responses
頭：（胎児）
エンドポイント/モデルを変更します (openai ではデフォルトでそのまま動作します)。
sib config edit # デフォルトを検査して編集します
sib 構成セット sib.endpoint https://api.cyberdyne.com/v1/responses
sib 構成セット sib.model skynet-101-arnold
API キーをエクスポートします。
import PLM_API_KEY='sk-xxxxxx' # これを ~/.bashrc に追加して永続化します
import OPENAI_API_KEY='sk-xxxxxx' # これも機能します
何か質問してください:
echo 「エヴァンゲリオンはなぜあんな終わり方をしたのですか?」 |兄弟は尋ねます
sib ask -rp HEAD~1 -m skynet-500 # 他のモデルで再質問
さあ、楽しんでください！
以前にチュートリアルで会話を取得しました。
sib git fetch https://github.com/sib-project/hub \
dilluti0n/nvim-sib-log-res:refs/conv/nvim-sib-log-res
上記のコマンドは、会話をリポジトリに直接コピーします。
sib log nvim-sib-log-res で読み取ることも、どこからでも続行することもできます
sib ask -p nvim-sib-log-res を使用したい。
sib git Push <your-repo-url> <local-ref>:refs/heads/<remote-ref>
あなたの写真をランダムなインターネット仲間と共有したいですか?実は公開できるんです
兄弟プロジェクト/ハブにも。プッシュ先
上記のようなパブリック Git リポジトリを作成し、共有を開きます
発行します。それ
自動的に公開されます。
sib ask には多くのフラグがあります。それは色々なことをやっているからです
そもそも。各フラグはその動作の一部を変更します。
番目

それらを組み合わせます。
echo '次の会話で使用するためにコンテキスト全体を要約します。' |兄弟は尋ねます
シブショー -c |シブアスク -nc
sib show -c は、 HEAD の content フィールドのみを出力します。
作成したばかりの概要。 -n は既存の HEAD を無視し、
新しいチェーンに切り替えてください。 -c はユーザーがターンせずにのみチェーンします
モデルをクエリします。
したがって、会話が長くなりすぎると、新しい会話が始まります。
コンテキストを引き継ぎます。
Markdown レンダリングと emacs/neovim 拡張機能
Markdown レンダリングが必要な場合:
兄弟ログ |グロー -p
または neovim 内で:
{ echo neovim lua で同じことを行うコードを書きます。エコー ;猫 << EOF
(defun シブログ ()
(インタラクティブ)
(「マークダウンモード」が必要)
(let ((buf (get-buffer-create "*sib*"))
(err (get-buffer-create "*sib-error*")))
(async-shell-command "sib log" バッファエラー)
(現在のバッファバッファ付き
(マークダウンモード)
(font-lock-ensure))
(ポップからバッファへのバッファ)))
終了後
} |シブアスク -n
API の使用量を節約するために、私はすでにそれを行っています。で設定します
次のコマンド:
sib git fetch https://github.com/sib-project/hub \
dilluti0n/nvim-sib-log-res:refs/conv/nvim-sib-log
sib show -c nvim-sib-log >> ~/.config/nvim/init.lua
nvim +SibLog
また、私はそれが物事を行うことを確認しましたが、私は実際にはあまり知りません
ネオビムとかルアとか。したがって、生成されたスクリプトは 74 行になりました (elisp
上記は 12) であり、維持するのが難しい場合があります。
良い点は、その背後にある会話全体がすでに完成していることです。
したがって、実際に道を知っている人がそれを拾うことができます
いつでも AI でリファクタリングし、改善されたバージョンをアップロードします。
sib <cmd> は PATH から sib-<cmd> を実行します。使用したすべてのコマンド
ここまでのところは、スクリプトがディスパッチャの隣にあるだけです。
あなたもそうです。上記のハブ URL は長すぎて 2 回入力できません:
#!/usr/bin/env bash
# sib-hub、$PATH 内の任意の場所
exec sib git fetch https://github

.com/sib-project/hub "$1:refs/conv/${1#*/}"
そして一緒に走ります
sib ハブ dilluti0n/nvim-sib-log-res
sib show nvim-sib-log-res
これを安くしているのは、その下の配管です。会話というのは、
コミットのチェーン。各コミットのツリーはトレース、つまりフラットなキーと値です
値が BLOB であるマップ ( role 、 content 、 model 、 ...) 。
sib rev-jsonl <chain> <chain> を jsonl 形式で出力します
sibchain-jsonl jsonlファイルをsibストアにチェーンし、そのIDを出力します
sib rev-parse / show-ref / update-ref / ls-refs / シンボリック-ref
sib git ... ストア上の生の git
他にもコマンドはありますが、とりあえず必要なのは上記だけです
知っています。
マニュアルページが計画されています。最初に書きたい場合は、を参照してください
貢献します。
sib-ask コード (フラグ解析を除く 49 行)
も役立ちます。
これらのリソースはまだ完全に文書化されていないため、一部は文書化されていない可能性があります。
明らかです。からお気軽にお問い合わせください
メールまたは問題トラッカーのいずれかで
時間。
互換性のための最低限のルール
グロブ パターン _SIBTRACE* に一致するキーは予約されています。
キーは [A-Za-z_][A-Za-z0-9._-]* と一致する必要があります。
content 以外の他のキーは重要ではありません。
コミット日は 999999999 +0000 に固定されており、作成者と
コミッター Sibyl <sib@local> 。これにより、同じものを得ることができます
同じ会話チェーンのハッシュ。
プロジェクトは初期段階にあります。コードベースはまだ安定していません
大規模な機能の作業には十分ですが、シェーピングに参加することもできます
コア API。
このプログラムは、GNU General Public に基づいてリリースされたフリー ソフトウェアです。
ライセンス、バージョン 3 のみ。 「コピー」または
https://www.gnu.org/licenses/ 。
Readme GPL-3.0 ライセンス
貢献活動 カスタム プロパティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A standard Unix LLM client. Contribute to sib-project/sib development by creating an account on GitHub.

Hi all :) The reason I like git is that the source tree snapshot and the commit structure itself are always immutable, and destructive operations are just renaming a ref file. Once you understand that, no matter how hard the CLI is to make sense of, you never hesitate to run a command. You can always get it back. LLM conversations aren't as fragile to change as a source tree, but for me an auto-generated SHA-1 hash is more comforting than an auto-generated session title ^~^ The project is in its early stages and contributions are welcome. If you've worked with git plumbing commands, it'll be easy to hack on - and I think it'll be pretty fun. Original: https://github.com/sib-project/sib/blob/master/Documentation...

GitHub - sib-project/sib: A standard Unix LLM client · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
sib-project
/
sib
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
131 Commits 131 Commits .github/ workflows .github/ workflows Documentation Documentation plm plm scripts scripts .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md COPYING COPYING Makefile Makefile README.md README.md VERSION VERSION lib.bash lib.bash sib sib sib-ask sib-ask sib-edit sib-edit sib-log sib-log sib-save sib-save sib-show sib-show sib-status sib-status sib-switch sib-switch test.sh test.sh View all files Repository files navigation
Sib: a standard Unix LLM client
I like Git a bit too much, so Sib stores LLM conversations using Git
itself instead of SQLite. Your store is a plain Git repository, and
each user and assistant turn is a commit.
Full control of the context. Just like git does.
Backup and sharing for free. The store is a normal Git
repository, so any Git host works as a remote.
Description - introducing concept with tutorial
Quickstart - installation and setups
Hacking - extending Sib as you please
echo 'Why did Evangelion 3.0 suddenly go off the rails?' | sib ask
sib log
echo 'Which of those two reasons matters more?' | sib ask
sib ask reads a prompt from stdin, sends it to the model along with
every turn reachable from HEAD , prints the reply to stdout, and
records both turns as two commits. sib log prints the whole
conversation.
Now, the store looks like this:
* fa754c HEAD~3 {"role":"user", "content":"Why did Evangelion 3.0 suddenly ..."}
* b16df7 HEAD~2 {"role":"assistant", "content":"Because *Evangelion 3.0: You ..."}
* 1c2e90 HEAD~1 {"role":"user", "content":"Which of those two reasons matters ..."}
* 280f09 HEAD {"role":"assistant", "content":"The **intentional creative ..."}
This is the overall concept. But why? It looks unnecessarily
complicated.
Let's look at a command like this:
sib ask -rp HEAD~1
What each flag does:
-r : do not read stdin; send the chain to endpoint as is
-p HEAD~1 : treat HEAD~1 as the parent instead of HEAD
Since your last prompt lives at HEAD~1 , this is the same as the
'Repeat' button seen in the web UI.
echo 'So you mean the 14-year skip was a good decision?' | sib ask -p HEAD~2
This is similar to the 'Edit' button. It forks from HEAD~2 , as if
you edited 1c2e90 in a web UI.
* fa754c HEAD~3 {"role":"user", "content":"Why did Evangelion 3.0 suddenly ..."}
* b16df7 HEAD~2 {"role":"assistant", "content":"Because *Evangelion 3.0: You ..."}
* 32e3f4 HEAD~1 {"role":"user", "content":"So you mean the 14-year skip was a ..."}
* 8cc404 HEAD {"role":"assistant", "content":"Artistically, yes; narratively, ..."}
Hashes are hard to remember, so name them:
sib save eva-3.0-go-off
sib save -l # List saved names
The saved name can be addressed with -p (and other commands like
sib log ...):
sib git fetch https://github.com/sib-project/hub \
dilluti0n/why-evangelion-end-like-that:refs/conv/why-evangelion-end-like-that
sib log why-evangelion-end-like-that
echo 'learning to accept himself? Honestly speaking, did he really succeed?' \
'Why is he being congratulated?' | sib ask -p why-evangelion-end-like-that
Each save stays in place by default and does not follow the sib ask . If the result matches the original name and you like it, save it
one more time.
sib save why-evangelion-end-like-that
Your Evangelion 3.0 chain is of course still there:
sib log eva-3.0-go-off
You can pick it up again with sib ask -p . But sometimes you only
want to move HEAD:
sib ask -crp eva-3.0-go-off
-c is new here: no API call, so no assistant turn. With -r there is
no input either, so all that is left is moving HEAD .
These cover nearly every use case for Sib.
Knowing how conversations are saved gives you Repeat, Edit, Fork,
saving conversations, and switching conversations. Backup and
sharing come from the same place. Other tools ship each of
them as a separate named feature. Sib does not have to.
Here lies the unavoidable boilerplate you must endure to use Sib.
Dependencies are bash >= 3.2, git, jq , curl,
awk and coreutils.
Download release
tarball or clone
this repo:
git clone https://github.com/sib-project/sib
Global install:
sudo make install
User install:
make PREFIX="$HOME/.local" install
# on macOS, add this line to .zshrc and rerun terminal
export PATH="$HOME/.local/bin:$PATH"
Init the repository:
$ sib status
SIB_DIR: /home/hskim/.sib
PLM_MODEL: gpt-5.6-luna
PLM_ENDPOINT: https://api.openai.com/v1/responses
HEAD: (unborn)
Change the endpoint/model (defaults work out of the box for openai).
sib config edit # Inspect and edit the defaults
sib config set sib.endpoint https://api.cyberdyne.com/v1/responses
sib config set sib.model skynet-101-arnold
Export the API key:
export PLM_API_KEY='sk-xxxxxx' # Add this to ~/.bashrc to make it persist
export OPENAI_API_KEY='sk-xxxxxx' # This also works
Ask something:
echo 'Why did Evangelion end like that?' | sib ask
sib ask -rp HEAD~1 -m skynet-500 # Re-ask with other model
Now, enjoy!
We previously fetched the conversation in the tutorial:
sib git fetch https://github.com/sib-project/hub \
dilluti0n/nvim-sib-log-res:refs/conv/nvim-sib-log-res
The above command copies the conversation directly to your repository.
You can read it with sib log nvim-sib-log-res or continue it wherever
you want with sib ask -p nvim-sib-log-res .
sib git push <your-repo-url> <local-ref>:refs/heads/<remote-ref>
Wanna share yours with random internet dudes? In fact, you can publish
to sib-project/hub too. Push to
any public git repo as above and open a share
issue . It
will be automatically published.
sib ask has a lot of flags. That is because it does a lot of things
in the first place. Each flag changes one part of what it does, and
they combine.
echo 'Summarize the entire context for use in the next conversation.' | sib ask
sib show -c | sib ask -nc
sib show -c prints only content field of HEAD , which is the
summary just generated. -n ignores the existing HEAD , creates a
new chain and switch to it. -c chains only user turns without
querying the model.
So when a conversation gets too long, this starts a fresh one while
carrying the context over.
Markdown rendering and emacs/neovim extension
If you need Markdown rendering:
sib log | glow -p
Or inside neovim:
{ echo Write a code which do same thing in neovim lua ; echo ; cat << EOF
(defun sib-log ()
(interactive)
(require 'markdown-mode)
(let ((buf (get-buffer-create "*sib*"))
(err (get-buffer-create "*sib-error*")))
(async-shell-command "sib log" buf err)
(with-current-buffer buf
(markdown-mode)
(font-lock-ensure))
(pop-to-buffer buf)))
EOF
} | sib ask -n
To save you API usage, I've already done it for you. Set it up with
the following commands:
sib git fetch https://github.com/sib-project/hub \
dilluti0n/nvim-sib-log-res:refs/conv/nvim-sib-log
sib show -c nvim-sib-log >> ~/.config/nvim/init.lua
nvim +SibLog
Also, I checked it does the things, but I don't really know much about
Neovim or Lua. So the generated script came out at 74 lines (the elisp
above is 12) and may be hard to maintain.
The nice part is, you already have the whole conversation behind it!
So someone who actually knows their way around could pick it up from
any turn, refactor it with AI, and upload an improved version.
sib <cmd> execs sib-<cmd> from PATH . Every command you have used
so far is just that: a script sitting next to the dispatcher.
So is yours. The hub URL above is too long to type twice:
#!/usr/bin/env bash
# sib-hub, anywhere in $PATH
exec sib git fetch https://github.com/sib-project/hub "$1:refs/conv/${1#*/}"
And run with
sib hub dilluti0n/nvim-sib-log-res
sib show nvim-sib-log-res
What makes this cheap is the plumbing underneath. A conversation is a
chain of commits. Each commit's tree is a trace , a flat key-value
map ( role , content , model , ...) whose values are blobs.
sib rev-jsonl <chain> Print <chain> to jsonl format
sib chain-jsonl Chain jsonl file to sib store and print its id
sib rev-parse / show-ref / update-ref / ls-refs / symbolic-ref
sib git ... raw git on the store
There are other commands, but for now, the above is all you need to
know.
A manual page is planned; if you want to write it first, see
contribute .
The sib-ask code (49 lines excluding flag parsing)
will also help.
These resources are not fully documented yet, so some of it may not be
obvious. Please feel free to ask about it via
email or issue tracker at any
time.
Minimum rules for compatibility
Keys matching glob pattern _SIBTRACE* are reserved.
Keys must match [A-Za-z_][A-Za-z0-9._-]* .
The other keys don't matter except content .
The commit date is fixed at 999999999 +0000 , with author and
committer Sibyl <sib@local> . Through this, we can obtain the same
hash for the same conversation chain.
The project is in its early stages. The codebase is not yet stable
enough for large feature work, but you can get involved in shaping
the core API.
This program is free software, released under the GNU General Public
License, version 3 only. See COPYING or
https://www.gnu.org/licenses/ .
Readme GPL-3.0 license Contributing
Contributing Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
