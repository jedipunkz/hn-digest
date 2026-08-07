---
source: "https://byteofdev.com/posts/turning-claude-postgres/"
hn_url: "https://news.ycombinator.com/item?id=49212117"
title: "Turning Claude into Postgres so I can raise a Series A"
article_title: "Turning Claude into Postgres so I can raise a Series A"
author: "JeanSebTr"
captured_at: "2026-08-07T15:44:30Z"
capture_tool: "hn-digest"
hn_id: 49212117
score: 1
comments: 0
posted_at: "2026-08-07T15:37:30Z"
tags:
  - hacker-news
  - translated
---

# Turning Claude into Postgres so I can raise a Series A

- HN: [49212117](https://news.ycombinator.com/item?id=49212117)
- Source: [byteofdev.com](https://byteofdev.com/posts/turning-claude-postgres/)
- Score: 1
- Comments: 0
- Posted: 2026-08-07T15:37:30Z

## Translation

タイトル: シリーズ A を上げるためにクロードを Postgres に変える
説明: Databricks さん、気をつけてください

記事本文:
ByteofDev シリーズ A を上げるためにクロードを Postgres に変えることについて
Web 開発やその他のプログラミングに関する興味深い投稿を直接入手できます
あなたの受信箱に！
シリーズ A を獲得できるようにクロードを Postgres に変える
現時点では、AI の機会を除けばデータベースのことを気にしている人はいないように思えます。 Databricks、PlanetScale、Supabase、さらには Andy Pavlo の謎の SYDHT さえも、AI 最適化、ベクター ストレージ、MCP サーバーなどを追加しています。そしてもちろん、投資家はそれらすべてに資金を投じています。そのお金が欲しいです。
しかし、AI が付属しているにもかかわらず、ほとんどのデータベースの中核機能は依然としてひどく退屈で AI がありません。確かに、LLM はデータベースを使用する可能性があり、場合によってはデータベース パラメーターを調整することもありますが、データベース自体は決定的なコードのみを実行します。ではなく、（ほぼ）すべてが AI だったらどうなるでしょうか?クロードはデータベース コードを書くのが得意ではないかもしれませんが、ファイル システムを渡せば確かにそれ自体がデータベースになる可能性があります。
高価で遅いですが、私はクソみたいなデータベースを構築することに慣れています。さらに、データベースの内部を表示する良い方法です。そこで、 Keenan Feldspar の精神に従って、素晴らしいパフォーマンスを発揮するトレンディなデモを構築してみましょう。
目標は、永続化データを含む実際の Postgres データベースと同様にクエリを処理して応答できるようにクロードに教えることです。すべての Postgres をカバーするには時間がかかりすぎるため、TPC-C Twitter ベンチマーク要件に焦点を当てました。
Sonnet 5 Medium を使用することにしました。これは、比較的高速で、コンテキスト ウィンドウが大きく、おそらくクエリ プランニングに十分スマートであるためです。そうですね、おそらく GLM-5.2 または GPT 5.6 Luna を使用するべきだったのですが、これらは私が始めたときには存在しませんでした。さらに重要なのは、GPTgres のサウンドがひどいことです。
クロードは HTTP 経由のテキスト応答でのみ対話できるため、Python Postgres プロキシである Buena Vista を使用して変換しました。

Postgres の TCP ベースのワイヤ プロトコル パケットを単純な SQL テキストに変換します。
私はクロードに、Postgres が COPY に使用する形式と同様の構造化テキスト形式で応答するように指示しました。この形式は、Buena Vista によって応答パケットに変換されます。これをプロンプトに追加しました:
結果のみを返信してください。それぞれのパーツに
独自の行:
STATUS: <コマンドタグ>
COLUMNS: [[" <col1 > "," < type1 > "],...] (結果セットがない場合は省略)
DATA (結果セットがない場合は省略)
< 1 行に 1 行、列は 1 つのタブで区切られ、テキスト COPY 形式 >
COLUMNS は、[名前、タイプ] ペアの JSON 配列です。各データ行は独自の行です。
Postgres のテキストの単一のタブ文字で結合された列の値 COPY
表現。数値は裸で、ブール値は t または f として、SQL NULL は \\N として記述されます。
(空のフィールドは長さ 0 の文字列であり、NULL とは異なります)。テキストは
周囲の引用符を付けずに文字通りに書かれています。特殊文字をエンコードする
バックスラッシュ エスケープ (\\t はタブ、\\n は改行、\\r はキャリッジ リターン、\\\\ は
リテラルのバックスラッシュ)。
例: SELECT id, name FROM users は 2 行を返します。
ステータス: セレクト 2
列: [["id","int4"],["name","text"]]
データ
1 リエトキンス
2 レトアトレイデス
そして…それはうまくいきます！いや、そうではありません。
psql - h 127.0 .0 .1 - p 5433
psql ( 18.4 、サーバー 14.0 ( BuenaVista / Claude ) )
ヘルプには「help」と入力します。
jacob => CREATE TABLE users ( users_id UUID PRIMARY KEY DEFAULT gen_random_uuid() , email TEXT UNIQUE NOT NULL , home_planet TEXT NOT NULL );
テーブルの作成
jacob => INSERT INTO users (email , home_planet ) VALUES (' [email protected] ' , 'Arrakis' ) ;
42 P01: リレーション「users」が存在しません
クロードは応答できますが、内部データベース ファイルを永続化する方法がまだありません。
ストレージとエンコードの実装
クロードはデータ自体を永続化する必要があります。ディスクへのアクセスを許可するのは非常に簡単です。私が作成します

d Postgres の fd.c ( md.c の一種) API に似た API。通常、ディスク上のページ (8kB チャンク) の形式でデータの取得と保存を処理します。次に、必要に応じて呼び出せるように、これらの関数をクロードに渡しました。ただし、ファイル内のデータをエンコードするのは少し複雑です。
通常の Postgres は行にバイナリ エンコード形式を使用します。固定長のバイナリ ヘッダーが行のメタデータを定義し、動的にサイズ変更されるすべての値 ( TEXT など) の長さは先頭の整数によって決まります。これは、どのテキスト エンコーディングよりもはるかに効率的です。整数はよりコンパクトで、固定サイズの列では、(CSV のカンマのように) 終わりを示すターミネータが必要ありません。この形式では、より単純なランダム アクセスも可能になるため、Postgres はページ全体を読み取ることなく、ページ内の特定のデータをクエリできます。残念ながら、クロードはそれをあまりうまく使いません。
クロードはバイナリを効率的にトークン化できません。さらに重要なことに、クロードはデータを検索するためにターミネータの代わりにオフセット/長さの整数を使用できるほどデータ サイズを十分に把握していません。さらに、JSON ベースのネットワークはバイナリを効率的にエンコードしません。したがって、テキスト形式が必要です。
私は最初、ファイル ストレージに同じタブベースの形式を使用してみましたが、うまくいくようでした。
残念なことに、クロードは少し反抗的になってしまいました（伏線？）。タブ付きの値が 2 つの値として読み取られないように、区切り文字にはタブ リテラルを使用し、列値内のタブには \t を使用するように指示しましたが、すべてにエスケープされたタブが使用され続けました。
ページ\t0\t2
ITEM\t1\t1\t0\t0\t1\t11\ttesting\ttesting -- 最後の値がタブで結合された両方の単語なのか、それともそれぞれの「テスト」が独自の値なのかを判断することはできません。
ああ、まあ。すべてのタプルを JSON 配列としてエンコードするように指示するように切り替えました。クロードはすでにエスケープ方法を知っており、パイプを分離として使用します。

トール。
ページ | 0 | 2
アイテム | 1 | 1 | 0 | 0 | 1 | [ 11 、 "テスト" 、 "テスト" ]
それが機能するのであれば文句は言いません。
psql ( 18.4 、サーバー 14.0 ( BuenaVista / Claude ) )
ヘルプには「help」と入力します。
jacob => CREATE TABLE users ( users_id UUID PRIMARY KEY DEFAULT gen_random_uuid() , email TEXT UNIQUE NOT NULL , home_planet TEXT NOT NULL );
テーブルの作成
jacob => INSERT INTO users (email , home_planet ) VALUES (' [email protected] ' , 'Arrakis' ) ;
42 P01: リレーション「users」が存在しません
まあ、そうではないと思います。ああ。
ブートストラッピング、カタログ、RelCaches、なんと！
Claude がテーブルを永続化できない問題は、ファイルのエンコーディングが原因ではありません。クロードは、そもそもどこからファイルを保存すればよいのかわかりません。データ ディレクトリに ls を与えて、自動的に解決するように指示することで回避することもできますが、それだと、現在よりもさらに遅くなります。代わりに、通常の Postgres の設計である RelCache を借用しました。
Postgres がテーブルまたは他のリレーションからデータを取得するときは、そのリレーションのデータが保存されている場所を見つける必要があります。この情報は Postgres カタログに保存されており、テーブル統計やストアド プロシージャなども含まれています。クロードはこれを知っています。ただし、ブートストラップの問題が発生します。 Postgres カタログはテーブル自体であるため、カタログのデータの場所情報がなければカタログにアクセスすることはできません。
ブートストラップを可能にするために、Postgres は一部のカタログの場所をハードコーディングします。私は、簡略化されたカタログ セットに同様のものを実装しました (ハードコードされた Python が必要でした、申し訳ありません)。ただし、クエリごとにコア カタログを読み取るのは非効率です。
Postgres は、ディスクからカタログを常に読み取って解析する代わりに、ストレージの場所などの関係データ (驚くべきことですが) をキャッシュする RelCache を使用します。 Postgres は RelCache をメモリ内に保持します。

ディスクからカタログをクエリする必要があり、データベースの知識をブートストラップします。同様に、クロードに初期ストレージの場所を与え、クエリごとの評価ループを保存する独自の RelCache のようなものを作成しました。
この場合、RelCache は単純です。 Python スクリプトはコア カタログのファイルを検索し、それをそのままプロンプトにダンプします。コンテキスト ウィンドウをもう少し占有しますが、実を言うと、リレーション データが大量のコンテキストになるまでこのデータベースをスケールアップしている人は誰もいません。
jacob => CREATE TABLE users ( users_id UUID PRIMARY KEY DEFAULT gen_random_uuid() , email TEXT UNIQUE NOT NULL , home_planet TEXT NOT NULL );
テーブルの作成
jacob => INSERT INTO users (email , home_planet ) VALUES (' [email protected] ' , 'Arrakis' ) ;
0 1 を挿入
jacob => SELECT * FROM ユーザー ;
ユーザーID |メール |ホーム_プラネット
--------------------------------------+-----------------------------+---------------
a3f1c2d4 - 5 e6b - 4 a1c - 9 f3d - 8 b2e7c1a4d9f | @corrino.govアラキス
( 1 行 )
Postgres ソース コードがまったく含まれていない、(実際には) 完全に機能する Postgres データベース。確かに、1 行を返す 1 つの SELECT を実行するには 10 秒かかり、0.03 ドルかかりますが、それは tokenmaxx を助けるだけです。
しかし、VCの資金を集めるためには規模を示す必要があります。
次に、より複雑なシナリオをテストします。より単純なユーザー スキーマを使用して、データベースに 5000 のプレースホルダー行を入力しました。
jacob = > CREATE TABLE ユーザー (
users_id UUID 主キー デフォルト gen_random_uuid () 、
名前のテキストが NULL ではありません、
home_planet テキストが NULL ではありません
) ;
テーブルの作成
jacob = > INSERT INTO users ( name , home_planet )
選択
substr (md5 (random() :: text || (i * 13) :: text ) , 1 , 8 ) AS 名 ,
substr ( md5 ( ランダム ( ) :: テキスト || ( i * 13 ) :: テキスト ) , 1 , 8 ) AS home_planet
F

ROM 生成シリーズ (1, 5000) AS i ;
0 5000を挿入
驚くべきことに、クロードは 408kB データ ファイルにすべての行を正しく作成し、保存しました。これはクエリの計算上の性質によって間違いなく助けられました (これにより、クロードは各行自体を処理する代わりにスクリプトを作成できるようになりました) が、それでも私が勝つつもりです。
ユーザーからカウント (*) を選択します。
数える
-------
5000
( 1 行 )
SELECT * FROM users ORDER BY users_id OFFSET 250 LIMIT 3 ;
ユーザーID |名前 |ホーム_プラネット
--------------------------------------+----------+---------------
0 debdf69 - 9 ad9 - 4787 - 887 e - 2060762696 d9 | 9 e227d09 | 087 f7577
0 デカブd1 - 4 f31 - 4 c69 - 8 f36 - f4ab62031572 | 3046 AF4B | 18 cc3f88
0 df18fad - 29 a5 - 449 b - b9f6 - 01 adda3295c2 | a025e8e8 | 4 d5616bb
( 3 行 )
grep による行のカウントは一貫しています。
$ grep -c "アイテム" ./data/base/16384
5000
ここで、基本的な挿入とスキャンを超えた実際の SQL 機能をいくつか試してみる必要があります。まず、クロードがインデックスをどのように処理するかを確認したかったのです。プロンプトにはインデックス カタログに関する指示がいくつかありますが、実際にインデックスを保存およびクエリする方法については何もありません。残念ながら、Postgres で最も一般的に使用されるインデックスである B ツリー インデックスの実装は、フラット テーブルよりも少し複雑なので、インデックスが正しく保存されるかどうかはあまり自信がありませんでした。
CREATE INDEX users_name ON users(name) ;
インデックスの作成
\di
インデックス一覧
スキーマ |名前 |タイプ |オーナー |テーブル
--------+-----------+----------+----------+----------
パブリック |ユーザー名 |インデックス |ポストグレ |ユーザー
パブリック |ユーザー_pkey |インデックス |ポストグレ |ユーザー
( 2 行 )
インデックスは正しく保存されているようです。次に、クエリ プランナーが実際にインデックスを使用するかどうかを確認します。
EXPLAIN ユーザーから名前を選択 ORDER BY 名前 LIMIT 3 ;
クエリプラン
-------------------------

-----------------------------------------------------
制限 (コスト = 0.15 . .0 .27 行 = 3 幅 = 32)
- > ユーザーの users_name を使用したインデックス スキャン (コスト = 0.15 . .20 .35 行 = 460 幅 = 32)
( 2 行 )
ユーザーから名前を選択 ORDER BY 名前 LIMIT 3 ;
名前
----------
00057 c04
000c060c
00291273
( 3 行 )
時間：212827.080ミリ秒（03：32.827）
それは…いいですね。ちょっと。 2 番目のクエリの時間を計測しましたが、3 分もかからないはずです。はい、これは厳密には高性能データベースではありませんが、このクエリでは 1 ページ (インデックスの最初のページ) の取得だけが必要なため、30 秒未満であるはずです。単一の小さなテーブルから 3 行を取得するだけのクエリについて、このようなことを言うことがどれほど奇妙に感じられるか、あなたにはわかりません。また、クロードが保存した唯一の新しいデータは、空のファイルノードがアタッチされたカタログ エントリです。つまり、クロードはインデックスを使用していると嘘をついているだけです。
クロードにインデックスの作成方法を伝えるための明示的な指示をいくつか追加してみました。
<!-- その他のプロンプト セグメント -->
## インデックス
インデックス (relkind='i') は B ツリーであり、Postgres 独自の nbtree と同じ構造です。 grep `manual/btree.md`
不明な場合は「B ツリー構造」。同じインデックスの独自のファイルに存在します。
{BLCKSZ} バイトのブロックですが、ページ レイアウトがテーブル ヒープとは異なります。
ブロック 0 はメタページです: a s

[切り捨てられた]

## Original Extract

Watch out, Databricks

ByteofDev About Turning Claude into Postgres so I can raise a Series A
Get interesting posts about web development and more programming straight
in your inbox!
Turning Claude into Postgres so I can raise a Series A
Right now, it seems like nobody cares about databases aside from their AI opportunities. Databricks, PlanetScale, Supabase, even Andy Pavlo’s mysterious SYDHT all are adding AI optimization, vector storage, MCP servers, etc. And, of course, investors dump money on all of them. I want that money.
But, for all of their AI appendages, most databases’ core functions are still woefully boring and AI-less. Sure, an LLM might use your database, and maybe even tune database parameters, but the database itself only runs deterministic code. What if, instead, (almost) everything was AI? Claude might not be great at writing database code, but surely it can be a database itself when I hand it a filesystem.
It will be expensive and slow, but I am not a stranger to building shitty databases . Besides, it is a good way to display database internals. So, in the spirit of Keenan Feldspar , let’s build a trendy demo that performs terribly.
The goal is to teach Claude to process and respond to queries like a real Postgres database would, including persisting data. Covering all of Postgres would take too long, so I focused on the TPC-C Twitter benchmark requirements.
I decided to use Sonnet 5 medium because it is relatively fast, has a larger context window, and is probably smart enough for query planning. Yeah, I probably should have used GLM-5.2 or GPT 5.6 Luna, but these didn’t exist when I started. More importantly, GPTgres sounds terrible.
Claude can only interact via text responses over HTTP, so I used Buena Vista , a Python Postgres proxy, to convert Postgres’s TCP-based wire protocol packets to simple SQL text.
I told Claude to respond with a structured textual format similar to the format Postgres uses for COPY , which is converted to response packets by Buena Vista. I added this to the prompt:
Reply with ONLY the result. Each part on
its own line:
STATUS: < command tag >
COLUMNS: [[" < col1 > "," < type1 > "],...] (omit if no result set)
DATA (omit if no result set)
< one row per line, columns separated by a single TAB, in the text COPY format >
COLUMNS is a JSON array of [name, type] pairs. Each data row is its own line:
column values joined by a single tab character, in Postgres's text COPY
representation. Numbers are written bare, booleans as t or f, and SQL NULL as \\N
(an empty field is a zero-length string, which is distinct from NULL). Text is
written literally with no surrounding quotes; encode special characters with
backslash escapes (\\t for tab, \\n for newline, \\r for carriage return, \\\\ for
a literal backslash).
Example: SELECT id, name FROM users returning two rows:
STATUS: SELECT 2
COLUMNS: [["id","int4"],["name","text"]]
DATA
1 lietkynes
2 letoatreides
And… it works! Well, not really.
psql - h 127.0 .0 .1 - p 5433
psql ( 18.4 , server 14.0 ( BuenaVista / Claude ) )
Type "help" for help .
jacob = > CREATE TABLE users ( users_id UUID PRIMARY KEY DEFAULT gen_random_uuid ( ) , email TEXT UNIQUE NOT NULL , home_planet TEXT NOT NULL ) ;
CREATE TABLE
jacob = > INSERT INTO users ( email , home_planet ) VALUES ( ' [email protected] ' , 'Arrakis' ) ;
42 P01: relation "users" does not exist
Claude can respond, but it still has no way of persisting the internal database files.
Implementing storage and encoding
Claude needs to persist the data itself. Giving it access to a disk is easy enough. I created an API similar to Postgres’s fd.c (also kind of md.c ) API, which normally handles data retrieval and storage in the form of pages (8kB chunks) on disk. I then passed those functions to Claude to call when needed. However, encoding the data in files is a little more complicated.
Normal Postgres uses a binary encoding format for rows, where fixed-length binary headers define the row’s metadata, and the length of all dynamically sized values (like TEXT ) is determined by leading integers. This is much more efficient than any text encoding: Integers are more compact, and fixed size columns don’t require a terminator to show their end (like a comma would in CSV). This format also allows for simpler random access, so Postgres can query specific data within a page without reading the whole page. Unfortunately, Claude doesn’t play so well with it.
Claude doesn’t tokenize binary efficiently, and, more importantly, Claude doesn’t have a strong enough grasp of data size to use offset/length integers instead of terminators to find data. Besides, the JSON-based networking doesn’t encode binary efficiently. So, I need a text format.
I first tried using the same tab-based format for file storage, which seemed to work.
Unfortunately, Claude got a bit rebellious (foreshadowing?). I told it to use tab literals for separators and \t for any tabs within column values so a value with a tab wouldn’t be read as two values, but it kept using escaped tabs for everything.
PAGE\t0\t2
ITEM\t1\t1\t0\t0\t1\t11\ttesting\ttesting -- it isn't possible to tell whether the last value is both words joined by a tab or if each "testing" is its own value
Oh well. I switched to telling it to encode all tuples as a JSON array, which Claude knows how to escape already, and use pipes as separators.
PAGE | 0 | 2
ITEM | 1 | 1 | 0 | 0 | 1 | [ 11 , "testing" , "testing" ]
I’m not complaining if it works.
psql ( 18.4 , server 14.0 ( BuenaVista / Claude ) )
Type "help" for help .
jacob = > CREATE TABLE users ( users_id UUID PRIMARY KEY DEFAULT gen_random_uuid ( ) , email TEXT UNIQUE NOT NULL , home_planet TEXT NOT NULL ) ;
CREATE TABLE
jacob = > INSERT INTO users ( email , home_planet ) VALUES ( ' [email protected] ' , 'Arrakis' ) ;
42 P01: relation "users" does not exist
Well, I guess not. Agh.
Bootstrapping, catalogs, and RelCaches, oh my!
The issue preventing Claude from persisting a table isn’t because of file encoding. Claude doesn’t know where to start storing files in the first place. I could get away with giving it ls on the data directory and telling it to figure things out itself, but that would be even slower than it already is. Instead, I borrowed from normal Postgres’s design: the RelCache.
When Postgres retrieves data from a table or other relation, it needs to find where that relation’s data is stored. This information lives in Postgres catalogs , which also contain table statistics, stored procedures, etc. Claude knows this. However, it runs into a bootstrapping issue. Postgres catalogs are tables themselves, and so, without the catalog’s data location info, it is impossible to access the catalog.
To allow bootstrapping, Postgres hardcodes some catalog locations. I implemented something similar (which did require some hardcoded Python, sorry) for a simplified catalog set. However, reading the core catalogs every query is inefficient.
Instead of reading and parsing catalogs off of disk constantly, Postgres uses a RelCache, which caches relation data (surprising, I know), like storage location. Postgres keeps the RelCache in memory, reducing its need to query catalogs off of disk and bootstrapping its database knowledge. Similarly, I created my own sort-of RelCache that gives Claude initial storage locations and saves it an evaluation loop on every query.
In this case, the RelCache is simple. A Python script finds the files of core catalogs and dumps them into the prompt verbatim. It takes up a bit more of the context window, but let’s be real, nobody is scaling this database up to the point where the relation data will be a significant amount of context.
jacob = > CREATE TABLE users ( users_id UUID PRIMARY KEY DEFAULT gen_random_uuid ( ) , email TEXT UNIQUE NOT NULL , home_planet TEXT NOT NULL ) ;
CREATE TABLE
jacob = > INSERT INTO users ( email , home_planet ) VALUES ( ' [email protected] ' , 'Arrakis' ) ;
INSERT 0 1
jacob = > SELECT * FROM users ;
users_id | email | home_planet
--------------------------------------+--------------------+-------------
a3f1c2d4 - 5 e6b - 4 a1c - 9 f3d - 8 b2e7c1a4d9f | lkynes @corrino.gov | Arrakis
( 1 row )
A (not really) fully functional Postgres database, with none of the Postgres source code. Sure, it takes 10 seconds and costs $0.03 to run a single SELECT that returns one row, but that is just helping you tokenmaxx.
But, in order to attract VC funding, I need to show scale.
Now it’s time to test more complex scenarios. I filled the database with 5000 placeholder rows using a simpler user schema.
jacob = > CREATE TABLE users (
users_id UUID PRIMARY KEY DEFAULT gen_random_uuid ( ) ,
name TEXT NOT NULL ,
home_planet TEXT NOT NULL
) ;
CREATE TABLE
jacob = > INSERT INTO users ( name , home_planet )
SELECT
substr ( md5 ( random ( ) :: text || ( i * 13 ) :: text ) , 1 , 8 ) AS name ,
substr ( md5 ( random ( ) :: text || ( i * 13 ) :: text ) , 1 , 8 ) AS home_planet
FROM generate_series ( 1 , 5000 ) AS i ;
INSERT 0 5000
Surprisingly enough, Claude created and stored every row correctly in a 408kB data file. It was definitely helped by the computational nature of the query (which allowed Claude to write a script for it instead of having to process each row itself), but I will take the win nonetheless.
SELECT count ( * ) FROM users ;
count
-------
5000
( 1 row )
SELECT * FROM users ORDER BY users_id OFFSET 250 LIMIT 3 ;
users_id | name | home_planet
--------------------------------------+----------+-------------
0 debdf69 - 9 ad9 - 4787 - 887 e - 2060762696 d9 | 9 e227d09 | 087 f7577
0 decabd1 - 4 f31 - 4 c69 - 8 f36 - f4ab62031572 | 3046 af4b | 18 cc3f88
0 df18fad - 29 a5 - 449 b - b9f6 - 01 adda3295c2 | a025e8e8 | 4 d5616bb
( 3 rows )
Counting the rows via grep is consistent.
$ grep -c "ITEM" ./data/base/16384
5000
Now I gotta try some actual SQL features beyond basic inserts and scans. First, I wanted to see how Claude handled indexes. There is some instruction in the prompt around index catalogs, but nothing about how to actually store and query indexes. Unfortunately, implementing B-tree indexes, which are the most commonly used indexes in Postgres, is a little more complicated than a flat table, so I wasn’t super confident the indexes would be stored correctly.
CREATE INDEX users_name ON users ( name ) ;
CREATE INDEX
\di
List of indexes
Schema | Name | Type | Owner | Table
--------+------------+-------+----------+-------
public | users_name | index | postgres | users
public | users_pkey | index | postgres | users
( 2 rows )
It appears that the indexes are being stored properly. Now, to see if the query planner actually uses the index.
EXPLAIN SELECT name FROM users ORDER BY name LIMIT 3 ;
QUERY PLAN
----------------------------------------------------------------------------------
Limit ( cost = 0.15 . .0 .27 rows = 3 width = 32 )
- > Index Scan using users_name on users ( cost = 0.15 . .20 .35 rows = 460 width = 32 )
( 2 rows )
SELECT name FROM users ORDER BY name LIMIT 3 ;
name
----------
00057 c04
000 c060c
00291273
( 3 rows )
Time : 212827.080 ms ( 03 : 32.827 )
That looks… good. Kinda. I timed the second query, which should not have taken three minutes. Yeah, this isn’t exactly a high performance database, but this query should only require one page retrieval (the first page of the index), so it should be <30 seconds. You have no idea how odd it feels to say that about a query that just retrieves 3 rows on a single small table. Also, the only new data Claude stored is a catalog entry with an empty filenode attached. So, Claude is just lying about using the index.
I tried adding some explicit instructions to tell Claude how to create indexes.
<!-- other prompt segments -->
## Indexes
An index (relkind='i') is a B-tree, the same structure as Postgres's own nbtree. grep `manual/btree.md`
"B-Tree Structure" if unsure. It lives in the index's own file in the same
{BLCKSZ}-byte blocks but with a page layout distinct from a table heap.
Block 0 is the METAPAGE: a s

[truncated]
