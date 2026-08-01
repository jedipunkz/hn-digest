---
source: "https://github.com/MiskaKan/thinair"
hn_url: "https://news.ycombinator.com/item?id=49133864"
title: "Integrate LLM into your Python runtime"
article_title: "GitHub - MiskaKan/thinair: Integrate LLM into your Python runtime · GitHub"
author: "ObviousDecline"
captured_at: "2026-08-01T12:59:01Z"
capture_tool: "hn-digest"
hn_id: 49133864
score: 1
comments: 0
posted_at: "2026-08-01T12:31:52Z"
tags:
  - hacker-news
  - translated
---

# Integrate LLM into your Python runtime

- HN: [49133864](https://news.ycombinator.com/item?id=49133864)
- Source: [github.com](https://github.com/MiskaKan/thinair)
- Score: 1
- Comments: 0
- Posted: 2026-08-01T12:31:52Z

## Translation

タイトル: LLM を Python ランタイムに統合する
記事のタイトル: GitHub - MiskaKan/thinair: LLM を Python ランタイムに統合する · GitHub
説明: LLM を Python ランタイムに統合します。 GitHub でアカウントを作成して、MiskaKan/thinair の開発に貢献してください。

記事本文:
GitHub - MiskaKan/thinair: LLM を Python ランタイムに統合する · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
ミスカカン
/
薄い空気
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
ママ

ブランチ内 タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
61 コミット 61 コミット docs docs .gitignore .gitignore ライセンス ライセンス README.md README.md SPEC.md SPEC.md bench.py bench.py car_chat.py car_chat.py pyproject.toml pyproject.toml Thinair.py Thinair.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
確率論的な Python オブジェクト。属性、メソッドなど何もないところから発明します。選択した LLM (ローカルまたはホスト) が自信を持って空白を埋めます。
特定のものをコード化し、残りを想像してください。
パブリック サーフェス全体が 1 つのクラスであり、アイデア全体が 1 行に収まります。つまり、「モノは確率のある値である」ということです。
color = Thing ( "rusty red" ,confidence = 0.4 ) # 想像された読み取り値
# 返却されたデモ内
+ color # 'rusty red': 値 (自由、推論なし)
~ color # 0.4: 確率 (無料、推論なし)
インターフェースはPythonそのものです。プロンプト文字列、メッセージ配列、出力解析はありません。通常のクラス、属性、およびメソッド呼び出しを記述します。モデルの機能は、空白のままにした場所で利用可能であり、常に確率で価格が設定されます。
あなたが自分で書くもの（コード、裸の値、あなた自身の言葉）はすべて確実であり、確率は 1.0 であり、モデルは決してそれに触れることができません。モデルが記入したすべてのもの (誰も定義していないフィールド、誰も書いていないメソッド) は、それがどの程度確実であるかを報告する Thing として返されます。このモデルは、オブジェクトはストーリーであり、すべてのインタラクションはその継続であるという 1 つの公理から答えを引き出します。他のすべてはそこから派生します。完全な仕様については、SPEC.md を参照してください。
3 つの演算子が表面をカバーし、推論呼び出しにかかるコストは @ <type> のみです。
クラスは必要ありません。言葉で説明したもの、自分自身の疑問を付加した裸の値、または凍結したオブジェクト全体 (以下の blob @ Car ) など、あらゆるものが Thing 空間に持ち上げられます。

最も単純なのは単語だけで、推論が影響するまでは自分の単語が確実なものとしてカウントされます。
シンエアーからの輸入物
Spider = Thing (「クモの足の数」)
+ スパイダー # 「スパイダーの足の数」: 自分の
# 言葉を返す;無料、推論なし
~ スパイダー #1.0: あなたの言葉は確かです
Legs = Spider @ int # 入力によって折りたたみが発生します: 1 つの推論
# 呼び出し、足は 8 を運ぶものになりました
+ 脚 # 8: 値、実数 int
~ 脚 # 0.99: 確率、裸のフロート
要件は連鎖し、満たされていないものは値を落としますが、確率は維持されるため、失敗はそれ自体で説明されます。
+ ( Spider @ int @ 0.8 ) # 8: p >= 0.8 で型付けされ、保証される
car = Thing ( "錆びた 1990 Toyota Hilux" ) # あなたが定義したことのない属性
推測＝車。 Price_eur @ float @ 0.9 # は最初の読み取り時に想像されます。
# したがって、これは入力された、ゲートされた推測です
+ 推測 # なし: バーをクリアしませんでした
~ 推測 # 0.1: 理由を説明する確率が生き残る
ifgue : ... # 失敗しました 物事は間違っています: ブランチ全体をゲートします
type オペランドは、プリミティブからスキーマを経て実際のクラスまでスケールします。
movie = Thing (「ゼノモーフが登場するリドリー・スコットの映画」)
+ (映画 @ { "タイトル" : str , "年" : int }) # {'タイトル': 'エイリアン', '年': 1979},
# スキーマ保証された辞書
@データクラス
クラスレコード:
タイトル : ストラ
年 : int
+ ( movie @ Record ) # Record(title='Alien', year=1979): モデル
# kwargs を想像し、コンストラクターがそれを構築します
自分自身の疑念を主張することができ、比較はモノ空間で行われます。比較は判断であり、バイト比較ではありません。
Price = Thing ( 19_990 ,confidence = 0.4 ) # 信念を Thing 空間に持ち上げます
+ ( 価格 @ 0.5 ) # なし: あなた自身がそう言っています
物 (「車」) < 物 (「猫」) # 誤り (p 0.75)、想像上の判断
+車。 Price < 20_000 # プレーンの値を最初に取り出します。
# 無料の Python セマン

チック
決定論と確率論の融合
Thing をサブクラス化し、確実な部分を書きます。書かれたコードと裸の値は特定のスケルトンです。これらは通常の Python として実行され、コストはかからず、モデルは決してそれらに触れることができません。それ以外はすべてオンデマンドで想像されます。
クラス 車 (物):
「」「道路車両です。」」
ホイール = 4 # 定義により確実です
def Horn ( self ): # 実際のコード: CPython で実行されます。
return "beep" # 推論は決して参照されません
car = Car ( 「錆びた 1990 年式トヨタ ハイラックス、エンジンが咳き込む、」
「フィンランドのシュラーガー放送局にラジオが引っかかった」）
車。ホイール # 4: 裸の int。推論は実行されず、何も請求されません
車。 Horn () # "ビープ音": 実際のコード、実際に実行される
+車。カラー # "ラスティ レッド": クラスが静まり返った場所を想像
～車。カラー#0.4：正直よくわからない
車。 owner = "Miska" # 裸の割り当て: 権威と
# ロックされています;どの計画もそれを上書きすることはできません
車。 Mood = Thing ( "unknown so far" ) # モデルが管理する可能性のあるスロット
来歴とは許可です。裸の値はプログラマに属し、物の値は想像力に属します。
作成したことのないメソッドは呼び出し時に想像され、動作します。状態を読み取り、状態を書き込み (自信を持ってジャーナルに記録)、実際のメソッドを呼び出して実際に実行します。結果は物であるため、チェーンを呼び出します。
問題 = 車。 list_your_problems (returns = [ str ])
+ 問題 # [「エンジンが咳き込んでいる」、「ラジオが故障している」
# フィンランドのシュラーガー駅'、'高い錆レベル']
車。 Repair_engine () # そのような方法はありません: 計画は想像されて実行されます
＃アウト;ストーリーには修復が含まれています
車。 ()を診断します。 severity_of_worst_issue () # 結果は Things:chain
# 想像上のコール上の想像上のコール
重要なブランチを保護します。内部では Thing.require(0.9) を使用します。0.9 未満の解像度では、フローする代わりに Thing.LowConfidence が発生します。
物と一緒に。必要 ( 0.9 ):
もしそうなら

r 。 can_drive (): # p 0.93: バーをクリアし、ブランチは信頼されます
plan_road_trip (車)
車。 vin_number # p 0.02: このワイルドが現在提起している推測
# Thing.LowConfidence を先に進める代わりに
それに話しかけてください
チャットボット フレームワークはなく、チャットは組み込みではありません。誰も作成していないため、他の欠落しているメソッドと同様に呼び出し時に想像されます。会話は単なるストーリーであるため、オブジェクトは会話を行うことができます。
True の場合:
print ( car . chat ( input ( "> " ))) # チャットもどこからともなく現れる
> どうして今朝私に暴言を吐いたのですか？
*咳咳* ...ビープ音。ほら、私は1990年式のトヨタ・ハイラックスです。私は錆びていて、疲れていて、
私のラジオはフィンランドのシュレーガーで止まっています。わざと崩れたわけではありません。
ただ…諦めました。エンジンは午前中ずっと咳き込んでいました。何かアイデアはありますか
これだけ年をとって、こんなに真っ赤になってから始めるのは、どれほど難しいことだろう？
すべてのターンが記録されるため、車はあなたの発言を記憶します。オブジェクトは、モノをモノに渡すことによって、同じ方法で互いのサイズを調整することもできます。
customer = Thing (「キャラバンと 2 匹の犬を連れた退職した夫婦、予算は控えめ」)
ピック = 顧客。好む（セダン、SUV、ロードスター、
戻り値 = { "選択" : str 、 "理由" : str })
ピックします。選択肢 # 「フルサイズ SUV: 7 席、牽引フック、喉が渇く」
ピックします。なぜ # 「キャラバンには牽引フックが必要ですが、これは SUV だけです」
# があり、2 匹の犬に必要なスペースを提供します。」
オブジェクトはドキュメントです
car.__getstate__() は、説明、状態、ストーリー、フラグなどの JSON 対応ドキュメントです。コードも重みもクライアントもありません。 blob @ Car は、書かれたメソッドを再アタッチして復元し、pickle はそのまま使用できます。 __source__ は、任意のオブジェクトを現在のクラスとしてレンダリングします。
print ( car . __source__ )
クラス 車 (物):
「」
道路車両。
錆びた1990年式トヨタ・ハイラックス、エンジンが咳き込む、フィンランドのシュラーガー放送局にラジオが引っかかる
「」
車輪 = 4
ああ

ner = 'Miska' # 書き込み (p = 1.0)
color = 'brown' #想像 (p = 0.10)
top_speed_kmh = 130 # 想像 (p = 0.10)
デフホーン（セルフ）：
「ピー」と返します
__story__ はもう 1 つのレンズです。すべての出来事、答え、想像上のステップを順番に記録した完全な日記です。一貫性と出所は無料で提供されます。
pip インストール Thinair
PyPI で利用可能です。依存関係はなく、ファイルは 1 つで、stdlib のみです。任意の OpenAI 互換エンドポイントを指します (デフォルトではローカル サーバーがターゲットになります)。
import THINAIR_BASE_URL= " http://127.0.0.1:8000/v1 " # デフォルト
エクスポート THINAIR_API_KEY= " 1234 "
import THINAIR_MODEL= " Qwen3.6-35B-A3B-oQ8-mtp "
import THINAIR_MAX_TOKENS=65536 # 完了ごとの合計予算 (考慮事項を含む)
リクエストはサーバーの JSON 出力モード ( response_format: json_object ) を要求し、サーバーがサポートしていない場合は静かにフリーフォームに戻ります。推論は 2 層のはしごを登ります。最初は、思考を無効にして質問に単発で回答します (サーバーがサポートしている場合は、応答スキーマに制限されたデコードが行われます)。つまずいた質問 (信頼度が低い、拒否の形をしたヌル、解析不可能な出力) のみが思考層を獲得します。そこでは、ストリームが到着するときに監視され、逐語的に繰り返し始めた瞬間にカットされます。モデルがループ内でリハーサルを続けた回答が回答として収集され、回答が途中でカットされた場合、予算全体ではなく草案に基づいて決定された完了補助金が与えられ、考え続けるだけのモデルは最終的に明らかな予算エラーを取得するため、暴走世代は決して予算を獲得できません。
またはコード内: Thing.defaults(model="...", Base_url="...", api_key="...") 。 URL、 complete(messages) -> text を持つプロバイダー オブジェクト、または単なる呼び出し可能なすべてのオブジェクトもインスタンスごとに機能します: Thing("a car", model=...) 。
下で何が起こっているかを確認するには、任意のブロックを Thing.debug() でラップします。

すべてのプロンプトと生の補完は、操作 ( read 、imaginative 、 judge 、 Collapse 、 condense ) によってラベル付けされて、標準エラー出力にダンプされます。 THINAIR_DEBUG=1 はグローバルに有効にします。
python car_chat.py # 錆びたハイラックスと話す;チャットは想像されたものであり、書かれたものではありません
ステータス
実験です。すべての未解決の属性には推論呼び出しが必要であり、答えはその背後にあるモデルと同程度にしか得られません。
LLM を Python ランタイムに統合する
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Integrate LLM into your Python runtime. Contribute to MiskaKan/thinair development by creating an account on GitHub.

GitHub - MiskaKan/thinair: Integrate LLM into your Python runtime · GitHub
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
MiskaKan
/
thinair
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
61 Commits 61 Commits docs docs .gitignore .gitignore LICENSE LICENSE README.md README.md SPEC.md SPEC.md bench.py bench.py car_chat.py car_chat.py pyproject.toml pyproject.toml thinair.py thinair.py View all files Repository files navigation
Probabilistic Python objects. Invent attributes, methods, anything out of thin air; an LLM of your choice (local or hosted) fills in the blanks, with confidence attached.
Code the certain, imagine the rest.
The entire public surface is one class, and the whole idea fits in one line: a Thing is a value with a probability .
color = Thing ( "rusty red" , confidence = 0.4 ) # the value the imagined read
# in the demo handed back
+ color # 'rusty red': the value (free, no inference)
~ color # 0.4: the probability (free, no inference)
The interface is Python itself. There are no prompt strings, no message arrays, and no output parsing: you write ordinary classes, attributes, and method calls, and the model's capabilities are available wherever you left a blank, always priced with a probability.
Everything you write yourself (code, bare values, your own words) is certain: probability 1.0, and the model can never touch it. Everything the model fills in (a field nobody defined, a method nobody wrote) comes back as a Thing that reports how sure it is. The model draws its answers from one axiom: an object is a story, and every interaction is a continuation of it. Everything else follows from it; see SPEC.md for the full spec.
Three operators cover the surface, and only @ <type> costs an inference call:
No class is required. Anything lifts into Thing space: a description in words, a bare value with your own doubt attached, or a whole frozen object ( blob @ Car , below). The simplest is words alone, and your own words count as certain until inference touches them:
from thinair import Thing
spider = Thing ( "the number of legs on a spider" )
+ spider # "the number of legs on a spider": your own
# words back; free, no inference
~ spider # 1.0: your words are certain
legs = spider @ int # collapsing happens through typing: one inference
# call, and legs is now a Thing carrying 8
+ legs # 8: the value, a real int
~ legs # 0.99: the probability, a bare float
Requirements chain, and an unmet one drops the value but keeps the probability, so failures explain themselves:
+ ( spider @ int @ 0.8 ) # 8: typed and vouched for at p >= 0.8
car = Thing ( "a rusty 1990 Toyota Hilux" ) # attributes you never defined
guess = car . price_eur @ float @ 0.9 # are imagined on first read,
# so this is a typed, gated guess
+ guess # None: did not clear the bar
~ guess # 0.1: the probability survives to say why
if guess : ... # failed Things are falsy: gate whole branches
The type operand scales from primitives through schemas to real classes:
movie = Thing ( "the Ridley Scott movie with the xenomorph" )
+ ( movie @ { "title" : str , "year" : int }) # {'title': 'Alien', 'year': 1979},
# a schema-guaranteed dict
@ dataclass
class Record :
title : str
year : int
+ ( movie @ Record ) # Record(title='Alien', year=1979): the model
# imagines the kwargs, your constructor builds it
You can assert your own doubt, and comparisons happen in Thing space: they are judgments, not byte comparisons:
price = Thing ( 19_990 , confidence = 0.4 ) # lift a belief into Thing space
+ ( price @ 0.5 ) # None: you said so yourself
Thing ( "a car" ) < Thing ( "a cat" ) # False (p 0.75), an imagined judgment
+ car . price < 20_000 # take the value out first for plain,
# free Python semantics
Deterministic meets probabilistic
Subclass Thing and write the parts you are sure of. Written code and bare values are the certain skeleton: they run as ordinary Python, cost nothing, and the model can never touch them . Everything else is imagined on demand:
class Car ( Thing ):
"""A road vehicle."""
wheels = 4 # certain by definition
def horn ( self ): # real code: runs in CPython,
return "beep" # inference is never consulted
car = Car ( "a rusty 1990 Toyota Hilux, engine coughs, "
"radio stuck on a Finnish schlager station" )
car . wheels # 4: a bare int; no inference ran, nothing billed
car . horn () # "beep": real code, really executed
+ car . color # "rusty red": imagined where the class was silent
~ car . color # 0.4: honestly unsure about it
car . owner = "Miska" # bare assignment: authoritative and
# locked; no plan may overwrite it
car . mood = Thing ( "unknown so far" ) # a slot the model may manage
Provenance is permission: bare values belong to the programmer, Thing values belong to the imagination.
Any method you never wrote is imagined at call time, and it acts: it reads state, writes state (with confidence, journaled), and calls your real methods, which actually execute. Results are Things, so calls chain:
problems = car . list_your_problems ( returns = [ str ])
+ problems # ['Engine is coughing', 'Radio is stuck on a
# Finnish schlager station', 'High rust level']
car . repair_engine () # no such method: a plan is imagined and acted
# out; the story now contains the repair
car . diagnose (). severity_of_worst_issue () # results are Things: chain
# imagined calls on imagined calls
Guard the branches that matter: inside with Thing.require(0.9): any resolution below 0.9 raises Thing.LowConfidence instead of flowing:
with Thing . require ( 0.9 ):
if car . can_drive (): # p 0.93: clears the bar, the branch is trusted
plan_road_trip ( car )
car . vin_number # p 0.02: a guess this wild now raises
# Thing.LowConfidence instead of flowing onward
Talk to it
There is no chatbot framework, and chat is not a built-in: nobody wrote it, so it is imagined at call time like any other missing method. The object can hold a conversation because a conversation is just more story:
while True :
print ( car . chat ( input ( "> " ))) # chat appears out of thin air too
> Why did you break down on me this morning?
*cough cough* ... beep. Look, I'm a 1990 Toyota Hilux. I'm rusty, I'm tired,
and my radio is stuck on Finnish schlager. I didn't break down on purpose;
I just... gave up. The engine was coughing all morning. Do you have any idea
how hard it is to start when you're this old and this red?
Every turn is journaled, so the car remembers what you said. Objects can also size each other up the same way, by handing Things to a Thing:
customer = Thing ( "a retired couple with a caravan and two dogs, modest budget" )
pick = customer . prefers ( sedan , suv , roadster ,
returns = { "choice" : str , "why" : str })
pick . choice # "a full-size SUV: seven seats, tow hook, thirsty"
pick . why # "The caravan requires a tow hook, which only the SUV
# has, and it provides necessary space for the two dogs."
Objects are documents
car.__getstate__() is a JSON-able document: description, state, story, flags. No code, no weights, no client. blob @ Car restores it with written methods reattached, and pickle works out of the box. __source__ renders any object as the class it currently is:
print ( car . __source__ )
class Car ( Thing ):
"""
A road vehicle.
a rusty 1990 Toyota Hilux, engine coughs, radio stuck on a Finnish schlager station
"""
wheels = 4
owner = 'Miska' # written (p = 1.0)
color = 'brown' # imagined (p = 0.10)
top_speed_kmh = 130 # imagined (p = 0.10)
def horn ( self ):
return "beep"
__story__ is the other lens: the full journal of every event, answer, and imagined step, in order. Consistency and provenance come for free.
pip install thinair
Available on PyPI . No dependencies, one file, stdlib only. Point it at any OpenAI-compatible endpoint (defaults target a local server):
export THINAIR_BASE_URL= " http://127.0.0.1:8000/v1 " # default
export THINAIR_API_KEY= " 1234 "
export THINAIR_MODEL= " Qwen3.6-35B-A3B-oQ8-mtp "
export THINAIR_MAX_TOKENS=65536 # total budget per completion, thinking included
Requests ask for the server's JSON output mode ( response_format: json_object ) and quietly fall back to freeform if the server doesn't support it. Inference climbs a two-tier ladder: questions are first answered single-shot with thinking disabled (and, where the server supports it, decoding constrained to the reply schema); only a question that stumbles (low confidence, a refusal-shaped null, unparseable output) earns the thinking tier. There the stream is supervised as it arrives and cut the moment it starts repeating itself verbatim. A reply the model kept rehearsing inside a loop is harvested as the answer, an answer cut mid-way gets a completion grant sized from the draft rather than the whole budget, and a model that only keeps thinking eventually gets a clear budget error, so runaway generation can never capture the budget.
Or in code: Thing.defaults(model="...", base_url="...", api_key="...") . A URL, a provider object with complete(messages) -> text , or a bare callable all work per instance too: Thing("a car", model=...) .
To see what is happening underneath, wrap any block in with Thing.debug(): and every prompt and raw completion is dumped to stderr, labeled by operation ( read , imagined , judge , collapse , condense ). THINAIR_DEBUG=1 turns it on globally.
python car_chat.py # talk to a rusty Hilux; chat is imagined, not written
Status
An experiment. Every unresolved attribute costs an inference call, and answers are only as good as the model behind them.
Integrate LLM into your Python runtime
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
