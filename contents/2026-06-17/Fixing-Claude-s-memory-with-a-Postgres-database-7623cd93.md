---
source: "https://www.makeuseof.com/fixed-claudes-memory-problem-postgres-database-changed-everything/"
hn_url: "https://news.ycombinator.com/item?id=48568063"
title: "Fixing Claude's memory with a Postgres database"
article_title: "I fixed Claude's memory problem with a Postgres database and it changed everything"
author: "dstala"
captured_at: "2026-06-17T10:37:56Z"
capture_tool: "hn-digest"
hn_id: 48568063
score: 1
comments: 0
posted_at: "2026-06-17T09:55:04Z"
tags:
  - hacker-news
  - translated
---

# Fixing Claude's memory with a Postgres database

- HN: [48568063](https://news.ycombinator.com/item?id=48568063)
- Source: [www.makeuseof.com](https://www.makeuseof.com/fixed-claudes-memory-problem-postgres-database-changed-everything/)
- Score: 1
- Comments: 0
- Posted: 2026-06-17T09:55:04Z

## Translation

タイトル: Postgres データベースを使用してクロードの記憶を修正する
記事のタイトル: Postgres データベースに関するクロードのメモリ問題を修正したところ、すべてが変わりました
説明: 私の AI は金魚症候群を発症しなくなりました。

記事本文:
Postgres データベースに関するクロードのメモリ問題を修正したところ、すべてが変わりました
メニュー
サインイン
今すぐサインイン
閉じる
PC＆モバイル
サブメニュー
Windows2
技術的な
サブメニュー
技術の説明
フォローする
フォローしました
いいね
スレッド
1
さらなるアクション
概要
このストーリーの要約を生成する
サインイン
今すぐサインイン
生産性
アンドロイド
スマートテレビ
ネットワーキング
Windows 11
エンターテイメント
閉じる
Postgres データベースに関するクロードのメモリ問題を修正したところ、すべてが変わりました
写真提供: ヤドゥラ・アビディ |帰属は必要ありません。
によって
ヤドゥラ・アビディ
2026 年 6 月 16 日、午後 1 時 30 分 EDT に公開
Yadullah Abidi は、デリー大学でコンピューター サイエンスを卒業し、チェンナイのアジア ジャーナリズム大学でジャーナリズムの大学院の学位を取得しています。 Windows および Linux システム、プログラミング、PC ハードウェア、サイバーセキュリティ、マルウェア分析、ゲームに関して 10 年以上の経験を持つ彼は、深い技術知識と強力な編集本能を組み合わせています。
Yadullah は現在、スタッフ ライターとして MakeUseOf に執筆し、サイバーセキュリティ、ゲーム、消費者向けテクノロジーをカバーしています。彼は以前、Candid.Technology で副編集長として、また The Mac Observer でニュース編集者として働いており、そこで猛烈なサイバー攻撃から最新の Apple 技術まであらゆるものをレポートしていました。
Yadullah はジャーナリズムの仕事に加えて、JavaScript/TypeScript、Next.js、MERN スタック、Python、C/C++、AI/ML の経験を持つフルスタック開発者です。マルウェアの分析、ハードウェアのレビュー、GitHub でのツールの構築など、彼は実践的な開発者の視点をテクノロジー ジャーナリズムにもたらしています。
MakeUseOf アカウントにサインインする
私たちを追加してください
概要
このストーリーの要約を生成する
フォローする
フォローする
フォローしました
フォローしました
いいね
いいね
スレッド
1
ログイン
物語の内容を事実に基づいて要約すると次のとおりです。
何か違うことを試してください:
事実を見せてください
私が5歳であるかのように説明してください
G

気軽な総括をします
クロード コードに真剣に取り組んだことがある人なら、AI がすべてを忘れてしまうことにすでにイライラしているはずです。セッションの最初の 20 分は、プロジェクトの構造、コーディング規約、特定のライブラリやツールを使用する理由の説明に費やしますが、Claude は次のセッションではすべてを忘れてしまいます。
しかし、正しい記憶力があれば、クロード・コードはそれ自体で軍隊になることができます。したがって、クロード トークン ウィンドウがプロジェクトに対して小さすぎると思われる場合は、それを補う簡単な方法がいくつかあります。
メッセージの上限は、Pro に料金を支払う最も興味深い理由ではありません。
クロードは本当に忘れているのではなく、最初からやり直しているのです
コンテキスト ウィンドウとチャット履歴が記憶と同じものではない理由
クロードがあなたが説明したことを忘れるのは、従来の意味ではバグではありません。 Claude Code は、現在の会話、ファイルの読み取り、ツールの出力、指示を保持する作業メモリであるコンテキスト ウィンドウ上で動作します。プランとモデルに応じて、その期間は 200,00 トークンから 100 万トークンまでになります。すごく大きいように思えますが、予想よりもすぐにいっぱいになります。
あなたが開くすべてのファイル、クロードが実行するすべての bash コマンド、すべての長い往復操作がそのコンテキスト ウィンドウに食い込みます。ウィンドウがいっぱいに近づくと、クロードは自動的にチャットの圧縮を開始します。まず古いツールの出力を削除してから、以前の会話を要約します。ただし、問題は、圧縮には損失が多く、セッションの早い段階での詳細な指示や決定が静かに失われる可能性があることです。
この問題には 2 つの修正があります。すべての概要を手動で別のファイルに保存し、新しいセッションを開始するたびに詳細を思い出すためにそれを読むようにクロードに依頼するか、コンテキスト ウィンドウをクリアするか、クロードに本物の長期的な情報を提供する Postgres データベースを接続することができます。

すべてのセッション、すべての /clear コマンド、さらにはマシンの切り替えごとに存続するメモリ。
クロード
開発者
人間の PBC
価格モデル
無料、購読可能
Postgres データベースは方程式を変える
すべての会話を生き残る永続的な記憶をクロードに与える
写真提供: ヤドゥラ・アビディ |帰属は必要ありません。
長期にわたるプロジェクト、チーム環境、またはコンテキストのドリフトによって数時間と数千のトークンが犠牲になる可能性があるコードベースなど、深刻なものの場合は、クロードのコンテキスト ウィンドウの完全に外側に存在するメモリが必要になります。そこで、Postgres が登場します。MCP (モデル コンテキスト プロトコル) 経由でクロードに接続されます。MCP (モデル コンテキスト プロトコル) は、クロードのような AI ツールを外部ツールやサービスに接続できるようにするオープン スタンダードです。この場合、ファイルを保存およびフェッチするための Postgres データベースへの読み取りおよび書き込みアクセス権がクロードに与えられます。
これを機能させるツールは Memory Vault と呼ばれます。これは、ハイブリッド検索のために Postgres と pgvector を組み合わせたオープンソース プロジェクトです。意味上の類似性とキーワードの一致を組み合わせたものです。そのため、クロードは、正確な単語を覚えていない場合でも、適切な記憶を見つけることができます。セットアップは、公式 GitHub リポジトリのクローンを作成し、Docker コンテナを起動するだけで簡単です。これらのコマンドを順番に実行します。
Yadullah Abidi によるスクリーンショット |帰属は必要ありません。
git clone https://github.com/mihaibuilds/memory-vault.git
CD メモリー保管庫
ドッカー構成 -d
データベース側については以上です。 Docker はイメージをプルし、移行を実行し、すべてを自動的に開始します。次のコマンドを使用して、すべてが稼働しているかどうかを確認できます。
Yadullah Abidi によるスクリーンショット |帰属は必要ありません。
docker compose exec アプリのメモリボールトのステータス
「データベース: 正常」と表示され、新規インストールの場合はチャンク数がゼロになっているはずです。メモリへの保存を開始すると、このチャンク カウントは

上がるよ。 Memory Vault は、http://localhost:8000 で Web ダッシュボードも起動し、保存されたメモリを直接参照、検索、管理できます。
次に、それをクロード コードに接続する必要があります。これを行うには、次のスニペットを .claude.json ファイルに追加するだけです。
{
"mcpサーバー": {
"メモリ保管庫": {
"コマンド": "Python",
"args": ["-m", "src.mcp"],
"cwd": "/パス/to/memory-vault",
"環境": {
"PYTHONPATH": "/path/to/memory-vault",
"DB_HOST": "ローカルホスト",
"DB_PORT": "5432",
"DB_NAME": "memory_vault",
"DB_USER": "メモリボールト",
"DB_PASSWORD": "memory_vault"
}
}
}
}
/path/to/memory-vault を、リポジトリのクローンを作成した場所に必ず置き換えてください。また、Python 3.11 以降と、複製されたディレクトリ内から pip install -e ".[mcp]" を実行してインストールされる Memory Vault の依存関係も必要です。完了したら、Claude Code を再起動し、「/mcp」と入力してサーバーが接続されているかどうかを確認します。すべてがうまくいけば、すぐに Postgres データベースへの保存を開始できます。
Yadullah Abidi によるスクリーンショット |帰属は必要ありません。
Windows で Claude Code を使用しているが、Postgres データベースを Linux サーバーでホストしたい場合は、追加の問題に対処する必要があります。 Windows 上の Python のデフォルトの非同期ループ (ProactorEventLoop) は、Memory Vault が使用するデータベース ドライバーである psycopg3 と互換性がありません。そのため、データベース パスを正しく設定したとしても、有用なエラーが発生することなく 30 秒後にタイムアウトになります。修正するには、クローンされたリポジトリの src/mcp/__main__.py ファイル内の 1 行を変更します。それを開いて次のようにします。
"""次のコマンドを使用して MCP サーバーの実行を許可します: python -m src.mcp"""
非同期をインポートする
インポートシステム
sys.platform == "win32" の場合:
asyncio.set_event_loop_policy(asyncio.WindowsSelectorEventLoopPolicy())
src.mcp.server インポートメインから
メイン()
.env ファイル内の DB_HOST を Linux マシンに設定してください。

■ localhost の代わりにローカル IP アドレス。これらの変更により、MCP サーバーを Windows 上で起動して実行し、リモート データベースに正常に接続できるようになります。
単純なメモリ ファイルを過小評価しないでください
マークダウンドキュメントが依然として驚くほど多くの問題を解決できる理由
写真提供: ヤドゥラ・アビディ |帰属は必要ありません。
データベース インフラストラクチャを起動したくない場合は、端末から離れることなく問題の大部分を解決できます。 Claude Code には、CLAUDE.md (すべてのセッションの開始時に自動的に読み取られるプレーン テキスト ファイル) を介してこれを行うためのメカニズムが組み込まれています。また、特定のセッション用の特定のファイルを作成するように要求し、それらを読み取ってプロジェクト固有の情報を取得するように依頼することもできます。組み込みの CLAUDE.md ファイルを使用している場合は、/memory を実行してシステム エディターでファイルを直接開き、編集や再編成を行うこともできます。
Claude メモリと Postgres の実践的なガイドを購読してください
もう 1 つの制限は、これらのファイルにアクセスするたびに、これらのファイルのコンテンツがコンテキスト ウィンドウに挿入されることです。つまり、トークンが消費されます。メモが数千行に増えると、保護しようとしているのと同じ予算を食いつぶすことになります。ファイルに焦点を当てたままにし、定期的にトリミングすることが最善です。
違いは日常の使用で現れます
繰り返しが減り、継続性が向上し、会話自体が実際に構築される
これを実行すると、個別のファイルから記憶するようにクロードに手動で依頼したり、メモをコンテキスト ウィンドウに食い込ませたりする必要はなくなります。自分が取り組んだ内容を覚えておいてくださいと頼むだけで、以前のセッションのことをクロードに尋ねることができ、完全なコンテキストとともに答えてくれます。圧縮の問題は完全に解決するわけではありませんが、問題ではなくなります。クロードが覚えておくべき重要なことはすでに保存されています

dはPostgresデータベースにあります。
ChatGPT は素晴らしいです。誤解しないでください。でもクロードのほうがずっといいよ。
共有バックエンドを使用すると、プロジェクトの完全な履歴を把握している Claude インスタンスを複数の開発者が操作できます。メモリはセッションごとにサイロ化されるのではなく、プロジェクト全体にわたって蓄積されます。
手動でメモリ ファイルを作成する方法を使用すると、日常のほとんどのイライラからすぐに解放されますが、完全な解決策ではありません。 Postgres セットアップは、拡張性があり、すべてを存続させ、回避策というよりもむしろクロードがプロジェクトを実際に理解しているように感じられるものが本当に必要なときにたどり着くものです。
2,000 万ダウンロードのこの Steam アプリは誤ってマルウェアを拡散させています
高価なイヤホンでも通話時の音がひどい、それが理由です
インテルは誤って完璧な小型 PC を作ったが、それを放棄した
もっと見る
機能が組み込まれたのでついに削除できる Android アプリ 6 選
人生の仕事を節約できる Windows の機能はデフォルトでオフになっています
アプリグリッドのないランチャーに切り替えたところ、以前よりも速くすべてを見つけることができました
もっと見る
今のトレンド
Adobe Acrobat サブスクリプションを無料アプリに置き換えましたが、元に戻す必要はありません
今年私が見つけた最高の Gemini アップグレードは Google Keep に隠されていました
Excel の二重括弧が何を意味するのか全く分かりませんでしたが、今ではほぼすべての数式で二重括弧を使用しています。
スレッド
1
MakeUseOf アカウントにサインインする
皆様のご意見をお待ちしております。以下のコメント欄であなたの視点を共有し、敬意を持って会話をしてください。
添付ファイル
私たちのコミュニティガイドラインを尊重してください。リンク、不適切な言葉遣い、スパムはありません。
コメントは保存されていません
率直な
率直な
率直な
#KW870247
2026-02-13 からのメンバー
フォロー中
0
トピックス
0
ユーザー
フォローする
フォローしました
0 フォロワー
見る
素晴らしいアイデアですね。 Postgres は今ではどこにでもあるようです。私はします

これを試してください
VLC をやめて、ついに遅延のない高品質の HDR 再生が可能になりました
「それは私ですか、それともインターネットですか？」と尋ねる必要はありません。この 10 ドルのビルドのおかげで
Copilot サブスクリプションをキャンセルし、Office タスク用の Claude に切り替えました
私の PC には、この隠された CPU BIOS 設定を見つけるまで、説明できない迷惑なフレームドロップが発生していました

## Original Extract

My AI stopped having goldfish syndrome.

I fixed Claude's memory problem with a Postgres database and it changed everything
Menu
Sign in
Sign in now
Close
PC & Mobile
Submenu
Windows2
Technical
Submenu
Tech Explained
Follow
Followed
Like
Threads
1
More Action
Summary
Generate a summary of this story
Sign in
Sign in now
Productivity
Android
Smart TVs
Networking
Windows 11
Entertainment
Close
I fixed Claude's memory problem with a Postgres database and it changed everything
Photo by Yadullah Abidi | No Attribution Required.
By
Yadullah Abidi
Published Jun 16, 2026, 1:30 PM EDT
Yadullah Abidi is a Computer Science graduate from the University of Delhi and holds a postgraduate degree in Journalism from the Asian College of Journalism, Chennai. With over a decade of experience in Windows and Linux systems, programming, PC hardware, cybersecurity, malware analysis, and gaming, he combines deep technical knowledge with strong editorial instincts.
Yadullah currently writes for MakeUseOf as a Staff Writer, covering cybersecurity, gaming, and consumer tech. He formerly worked as Associate Editor at Candid.Technology and as News Editor at The Mac Observer , where he reported on everything from raging cyberattacks to the latest in Apple tech.
In addition to his journalism work, Yadullah is a full-stack developer with experience in JavaScript/TypeScript, Next.js, the MERN stack, Python, C/C++, and AI/ML. Whether he's analyzing malware, reviewing hardware, or building tools on GitHub, he brings a hands-on, developer’s perspective to tech journalism.
Sign in to your MakeUseOf account
Add Us On
Summary
Generate a summary of this story
follow
Follow
followed
Followed
Like
Like
Thread
1
Log in
Here is a fact-based summary of the story contents:
Try something different:
Show me the facts
Explain it like I’m 5
Give me a lighthearted recap
If you've spent any serious time with Claude Code, you've likely already been frustrated by the AI forgetting everything. You spend the first twenty minutes of the session describing project structure, coding conventions, and why you're using specific libraries or tools, and Claude forgets everything in the next session.
But with the right memory, Claude Code can be an army in itself . So if the Claude token window seems too small for your projects, there are some simple ways you can supplement it.
The message cap is the least interesting reason to pay for Pro.
Claude isn't really forgetting—it's starting over
Why context windows and chat history aren't the same thing as memory
Claude forgetting what you described isn't a bug in the traditional sense. Claude Code operates on a context window — a working memory that holds your current conversation, file reads, tool outputs, and instructions. Depending on your plan and model, that window can be anywhere from 200,00 tokens to a million. It sounds enormous, but it fills up faster than you'd expect.
Every file you open, every bash command Claude runs, every long back-and-forth eats into that context window. When the window gets close to full, Claude automatically starts compacting your chat. It deletes older tool outputs first, then summarizes the earlier conversation. The problem, however, is that compaction is lossy, and detailed instructions or decisions from earlier in the session can quietly get lost.
There are two fixes for this issue. You can either ask Claude to manually save a summary of everything in a separate file and read it to remember details every time you start a new session or clear the context window, or you can wire a Postgres database that gives Claude genuine long-term memory that survives every session, every /clear command, and even switching machines.
Claude
Developer
Anthropic PBC
Price model
Free, subscription available
A Postgres database changes the equation
Giving Claude persistent memory that survives every conversation
Photo by Yadullah Abidi | No Attribution Required.
For anything serious like a long-running project, a team environment, or a codebase where context drift can cost you hours and thousands of tokens, you'd want memory that lives completely outside Claude's context window. That's where Postgres comes in, wired to Claude via MCP (Model Context Protocol) — an open standard that lets AI tools like Claude connect to external tools and services. In this case, it gives Claude read and write access to a Postgres database to save and fetch files.
The tool that makes this work is called Memory Vault , an open-source project that combines Postgres with pgvector for hybrid search — semantic similarity and keyword matching combined, so Claude can find the right memory even when you don't remember the exact words. Setup is as simple as cloning the official GitHub repository and spinning up a docker container. Run these commands one after the other:
Screenshot by Yadullah Abidi | No Attribution Required.
git clone https://github.com/mihaibuilds/memory-vault.git
cd memory-vault
docker compose up -d
And that's it for the database side. Docker pulls the images, runs migrations, and starts everything automatically. You can verify whether everything is up and running by using this command:
Screenshot by Yadullah Abidi | No Attribution Required.
docker compose exec app memory-vault status
You should see a Database: healthy and a chunk count of zero for fresh installations. As you start saving things to memory, this chunk count will rise. Memory Vault also spins up a web dashboard at http://localhost:8000 where you can browse, search, and manage stored memories directly.
Next, you need to wire it to Claude Code. To do so, just add the following snippet to your .claude.json file:
{
"mcpServers": {
"memory-vault": {
"command": "python",
"args": ["-m", "src.mcp"],
"cwd": "/path/to/memory-vault",
"env": {
"PYTHONPATH": "/path/to/memory-vault",
"DB_HOST": "localhost",
"DB_PORT": "5432",
"DB_NAME": "memory_vault",
"DB_USER": "memory_vault",
"DB_PASSWORD": "memory_vault"
}
}
}
}
Make sure to replace /path/to/memory-vault with wherever you cloned the repository. You'll also need Python 3.11 or higher and the Memory Vault dependencies installed by running pip install -e ".[mcp]" from inside the cloned directory. Once done, restart Claude Code and type in /mcp to check if the server is connected. If everything works, you can start saving to your Postgres database right away.
Screenshot by Yadullah Abidi | No Attribution Required.
If you're using Claude Code on Windows but want to host your Postgres database on a Linux server, you will have to deal with an additional issue. Python's default async loop on Windows (ProactorEventLoop) is incompatible with psycopg3, the database driver Memory Vault uses. So even if you've got the database path set right, it'll time out after 30 seconds without a useful error. The fix is to change one line in the cloned repository's src/mcp/__main__.py file. Open it and make it look like this:
"""Allow running the MCP server with: python -m src.mcp"""
import asyncio
import sys
if sys.platform == "win32":
asyncio.set_event_loop_policy(asyncio.WindowsSelectorEventLoopPolicy())
from src.mcp.server import main
main()
Make sure to set DB_HOST in your .env file to your Linux machine's local IP address instead of localhost. With these changes, you should be able to get the MCP server up and running on Windows and connected to your remote database cleanly.
Don't underestimate simple memory files
Why markdown documents still solve a surprising number of problems
Photo by Yadullah Abidi | No Attribution Required.
If you don't want to spin up database infrastructure, you can solve a majority of the problem without ever leaving your terminal. Claude Code has a built-in mechanism for this via CLAUDE.md — a plain text file that reads automatically at the start of every session. You can also ask it to create specific files for specific sessions and have it read them to get project-specific information. If you're using the built-in CLAUDE.md file, you can also run /memory to open the file directly in your system editor for editing and reorganization.
Subscribe for hands-on Claude memory and Postgres guides
Another limitation is that the content from these files gets injected into the context window every time you access them, meaning it does consume tokens. If your notes grow to thousands of lines, you're eating into the same budget you're trying to protect. It's best to keep the file focused and trim it regularly.
The difference shows up in everyday use
Less repetition, better continuity, and conversations that actually build on themselves
Once this is running, you no longer have to manually ask Claude to remember from individual files or have your notes eat into your context window. Simply ask it to remember what you worked on, and you can ask Claude anything from your previous sessions, and it answers with full context. The compaction problem doesn't go away entirely, but it stops mattering. The important things Claude needs to remember are already saved in the Postgres database.
ChatGPT is great; don't get me wrong. But Claude is so much better.
With a shared backend, multiple developers can work with a Claude instance that knows the full project history. The memory will compound across the full project rather than staying siloed per session.
The manual memory file approach gets you out of most daily frustrations fast, but it's not a full solution. The Postgres setup is what you reach for when you genuinely need something that scales, survives everything, and starts feeling less like a workaround and more like Claude actually knowing your project.
This Steam app with 20 million downloads is accidentally spreading malware
Your expensive earbuds still sound awful on calls, and this is why
Intel accidentally made the perfect tiny PC, then abandoned it
See More
6 Android apps you can finally delete because the feature is now built in
The Windows feature that could save your life's work is off by default
I switched to a launcher with no app grid and I find everything faster than I ever did before
See More
Trending Now
I replaced my Adobe Acrobat subscription with a free app and haven't needed to go back
The best Gemini upgrade I've found this year was hiding in Google Keep
I never knew what those double brackets in Excel meant — now I use them in almost every formula
Thread
1
Sign in to your MakeUseOf account
We want to hear from you. Share your perspective in the comments below, and please keep the conversation respectful.
Attachment(s)
Please respect our community guidelines . No links, inappropriate language, or spam.
Your comment has not been saved
Candid
Candid
Candid
#KW870247
Member since 2026-02-13
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
Great idea. Postgres seems to be everywhere now. I will try this
I ditched VLC and finally got high-quality HDR playback without any lag
I never have to ask “is it me or the internet?” thanks to this $10 build
I cancelled my Copilot subscription and switched to Claude for Office tasks
My PC had annoying frame drops I couldn't explain — until I found this hidden CPU BIOS setting
