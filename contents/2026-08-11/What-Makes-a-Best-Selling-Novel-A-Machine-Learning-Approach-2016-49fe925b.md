---
source: "https://marcinciura.wordpress.com/2016/04/17/what-makes-a-best-selling-novel/"
hn_url: "https://news.ycombinator.com/item?id=49256191"
title: "What Makes a Best-Selling Novel? A Machine Learning Approach (2016)"
article_title: "What Makes a Best-Selling Novel?"
author: "mci"
captured_at: "2026-08-11T11:38:52Z"
capture_tool: "hn-digest"
hn_id: 49256191
score: 1
comments: 0
posted_at: "2026-08-11T10:56:04Z"
tags:
  - hacker-news
  - translated
---

# What Makes a Best-Selling Novel? A Machine Learning Approach (2016)

- HN: [49256191](https://news.ycombinator.com/item?id=49256191)
- Source: [marcinciura.wordpress.com](https://marcinciura.wordpress.com/2016/04/17/what-makes-a-best-selling-novel/)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T10:56:04Z

## Translation

タイトル：ベストセラー小説とは何か？機械学習アプローチ (2016)
記事タイトル: ベストセラー小説とは何か?
説明: 機械学習アプローチ 2013 年に Ashok et al.は、文体に基づいてこの質問に 61 ～ 84% の精度で回答しました。一方、この記事では、ベストセラーのプロットのテーマを検討します。私のモデルでは、小説の商業的な成功をプロットから予測することはほとんどできないことに注意してください。それはかなりのことでしょう
[切り捨てられた]

記事本文:
ベストセラー小説とは何でしょうか?
コンテンツにスキップ
クラコビアン: 行列のねじれた双子
語彙距離を 3 次元で視覚化する
Python の 600 ワードと 60 行で構成される PageRank
ベストセラー小説とは何でしょうか?
2013 年に、Ashok ら。は、文体に基づいてこの質問に 61 ～ 84% の精度で回答しました。一方、この記事では、ベストセラーのプロットのテーマを検討します。私のモデルでは、小説の商業的な成功をプロットから予測することはほとんどできないことに注意してください。それは非常に驚くべき偉業であり、レビュアーを時代遅れにするでしょう。私の目標はもっと控えめで、統計的に有益な記事を見つけることでした。
PetScan と Wikipedia のページ エクスポートを使用して、Category:Novels に属する Wikipedia の記事を年ごとに 25,359 件ダウンロードしました。各記事から、 Plot 、 Plot summary 、 Symbol などの名前のセクション (存在する場合) を抽出し、MediaWiki マークアップを取り除き、小説のタイトル、出版年、ニューヨーク タイムズの小説ベストセラー リストでトップになったことがあるかどうかを示すブール値とともに SQLite データベースに保存しました。
小説からタイトル、年、ベストセラー、長さ(プロット)を選択します
ORDER BY random() LIMIT 5;
シャープの混乱 | 2003年 | 0 | 2759
ザ・レスキュー (スパークス小説) | 2000年 | 1 |
スレイヤーズ | 1989年 | 0 |
看守 | 1855年 | 0 | 2793
第 4 プロトコル | 1984年 | 1 | 5666
SELECT count(*) FROM Novels
WHERE プロットが NULL ではない;
17744
SELECT count(*) FROM Novels
WHERE プロットが NULL ではなく、かつ was_bestseller;
398
SELECT min(year) FROM Novels -- 出版年。
ベストセラーはどこにありましたか。 -- NYTのリストは1942年に始まります。
1941年
解釈しやすい結果を得るために、ポーター ステマーによって処理された記事の TF-IDF 変換に基づいてロジスティック回帰モデルを構築しました。パラメータにはデフォルト値があります。特にロジスティックスに関しては、

gression は L2 正則化を使用するため、ストップワードではないすべての小文字の単語がモデルに表示されます。
nltkをインポート
nltk.corpus からストップワードをインポート
nltk.stem インポートポーターから
sklearnインポートcross_validationから
sklearnインポートlinear_modelから
sklearnインポートパイプラインから
sklearn.feature_extraction インポート テキストから
def トークン化(
テキスト、
ステマー=ポーター.PorterStemmer(),
uppercase=set(string.uppercase),
stop_set=set(stopwords.words('english')),
punctuation_re = re.compile(
ur'['“”…–—!"#$%&\'()*+,\-./:;?@\[\\\]^_`{|}~]',
re.UNICODE)):
text = punctuation_re.sub(' ', text)
トークン = nltk.word_tokenize(text)
return [stemmer.stem(x) for x in tokens]
x. lower() が stop_set になく、x[0] が大文字でない場合]
X = []
y = []
接続 = sqlite3.connect('novels.sqlite')
connection.cursor().execute( の行の場合
"""小説からプロット、was_bestsellerを選択
WHERE 年 >= 1941 AND プロット IS NOT NULL"""):
X.append(行[0])
y.append(行[1])
connection.close()
X_train、X_test、y_train、y_test = (
cross_validation.train_test_split(X, y, test_size=0.3))
モデル = パイプライン.パイプライン(
[('tfidf', text.TfidfVectorizer(
lowercase=False、tokenizer=トークナイズ))、
('ロジスティック'、linear_model.LogisticRegression())])
model.fit(X_train, y_train)
このモデルは、小説 b がベストセラーになる確率をプロットの概要とともに返すことができます。
logit( b ) = −4.6 + 2.5 tfidf(弁護士, b ) + 2.4 tfidf(kill, b ) + ⋯ − 1.5 tfidf(惑星, b )
Pr(was_bestseller( b )|plot( b )) = e logit( b ) / (1 + e logit( b ) )
これらの係数を文脈に当てはめると、 tfidf(lawyer, The Firm ) ≈ 0.06 となります。偶然にも、モデルは logit( b ) > 0 を返します。つまり、トレーニング セットまたはテスト セットに新規 b がない場合、Pr(was_bestseller( b )|plot( b )) > 1/2 を返します。最も高い確率である 0.39 は、実際に 20 年 12 月のベストセラーである Cross Fire について予測されています。

10. TF-IDF の正規化を無効にするか、ロジスティック回帰の正則化を弱めた場合にのみ、モデルを列車セットに過剰適合させることができますが、テスト セットの精度と再現率は両方とも最大 20% になります。しかし、冒頭で書いたように、これはこの演習のポイントではありません。係数の絶対値が高い単語を見てみましょう。
明らかに、法的スリラーを書くのは利益があるようです。弁護士 +2.5、事件 +2.4、法律 +1.5、依頼人 +1.3、陪審 +1.3、裁判 +1.3、弁護士 +1.0、容疑者 +1.0、裁判官 +0.9、有罪判決者 +0.8。
殺害+2.4、殺人+1.8、テロリスト+1.2、銃撃+1.1、死体+1.1、死亡+1.0、シリアル+0.9、攻撃+0.9、暗殺者+0.8、誘拐+0.8、殺人者+0.8。
政治スリラーも悪くありません。エージェント +1.4、政治 +1.4、大統領 +1.3、亡命者 +1.2。
ビジネスが関係する可能性があります: 会社 +1.3、会社 +1.3、キャリア +1.1、百万 +1.0、成功 +1.0、ビジネス +0.9、お金 +0.9。
最後に、キャラクターには家族が必要です: 夫 +1.4、家族 +1.3、家 +1.2、夫婦 +1.2、娘 +1.2、赤ちゃん +1.1、妻 +1.0、父親 +1.0、子供 +0.9、出産 +0.8、妊娠中 +0.8、車の使用 +1.5、電話 +0.8。
ベストセラー作家になる可能性のある人が避けるべきジャンルは何ですか?
SF: 惑星 -1.5、人間 -1.0、宇宙 -0.7、星 -0.4、ロボット -0.3、軌道 -0.3。
児童文学：少年 -1.3、学校 -1.0、若者 -0.8、少女 -0.8、青少年 -0.4、教師 -0.4、叔母 -0.4、成長 -0.4。
地理と旅行: 村 -1.0、都市 -1.0、船 -0.8、道 -0.7、行き -0.7、土地 -0.6、冒険 -0.6、植民地 -0.5、原住民 -0.5、追跡 -0.5、山 -0.5、乗組員 -0.5、森林 -0.5、旅行 -0.5、居住 -0.4、航海−0.4、道路−0.4、地図−0.3、部族−0.3。
戦争：戦い -1.0、戦士 -0.6、戦争 -0.6、武器 -0.5、兵士 -0.5、軍隊 -0.5、味方

−0.4、敵−0.3、征服−0.3。
ファンタジー：魔法 -0.9、クリーチャー -0.5、マジシャン -0.4、ゾンビ -0.3、トレジャー -0.3、ドラゴン -0.3。
歴史: 王女 -0.5、統治 -0.5、王国 -0.4、城 -0.4、世紀 -0.4、統治者 -0.3、宮殿 -0.3 (ちなみに、『ゲーム・オブ・スローンズ』はリストの 3 位に入っただけなので、ベストセラーには数えられません)。
上記のコードは大文字の単語を無視することに注意してください。そうでない場合、最も重要な単語はベストセラーの書籍シリーズの登場人物の名前になります。スカーペッタ +3.0、ステファニー +2.9、エイラ +2.0 など。さらに、FBI +1.3、CIA +1.3、NATO +0.9、ソ連 +0.9、または地球 -1.1 などの追加の洞察が含まれます。
X で共有 (新しいウィンドウで開きます)
×
Facebook で共有 (新しいウィンドウで開きます)
フェイスブック
このサイトはスパムを低減するために Akismet を使っています。コメントデータがどのように処理されるかをご覧ください。
購読する
購読しました
marcinciura.wordpress.com
すでに WordPress.com アカウントをお持ちですか?今すぐログインしてください。

## Original Extract

A Machine Learning Approach In 2013, Ashok et al. answered this question basing on the writing style, with 61–84% accuracy. This post, on the other hand, examines plot themes in best sellers. Note that my model can hardly predict the commercial success of a novel from its plot. That would be quite a
[truncated]

What Makes a Best-Selling Novel?
Skip to content
Cracovians: The Twisted Twins of Matrices
Visualizing Lexical Distance in Three Dimensions
PageRank in 600 words and 60 lines of Python
What Makes a Best-Selling Novel?
In 2013, Ashok et al. answered this question basing on the writing style, with 61–84% accuracy. This post, on the other hand, examines plot themes in best sellers. Note that my model can hardly predict the commercial success of a novel from its plot. That would be quite a surprising feat, making reviewers obsolete. My goal was more modest: finding statistically profitable topics to write about.
Using PetScan and Wikipedia’s page export , I downloaded 25,359 Wikipedia articles belonging to Category:Novels by year . From each article, I extracted the section named Plot , Plot summary , Synopsis , etc. if present and, stripped of MediaWiki markup, saved it into an SQLite database along with the title of the novel, its year of publication, and a Boolean that indicates if it ever topped the New York Times Fiction Best Seller list :
SELECT title, year, was_bestseller, length(plot) FROM Novels
ORDER BY random() LIMIT 5;
Sharpe's Havoc | 2003 | 0 | 2759
The Rescue (Sparks novel) | 2000 | 1 |
Slayers | 1989 | 0 |
The Warden | 1855 | 0 | 2793
The Fourth Protocol | 1984 | 1 | 5666
SELECT count(*) FROM Novels
WHERE plot IS NOT NULL;
17744
SELECT count(*) FROM Novels
WHERE plot IS NOT NULL AND was_bestseller;
398
SELECT min(year) FROM Novels -- The year of publication.
WHERE was_bestseller; -- The NYT list starts in 1942.
1941
To obtain easy to interpret results, I have built a logistic regression model on top of the TF–IDF transformation of articles processed by the Porter stemmer . The parameters have default values. In particular, the logistic regression uses L2 regularization so all lowercase words that are not stopwords appear in the model.
import nltk
from nltk.corpus import stopwords
from nltk.stem import porter
from sklearn import cross_validation
from sklearn import linear_model
from sklearn import pipeline
from sklearn.feature_extraction import text
def Tokenize(
text,
stemmer=porter.PorterStemmer(),
uppercase=set(string.uppercase),
stop_set=set(stopwords.words('english')),
punctuation_re = re.compile(
ur'[’“”…–—!"#$%&\'()*+,\-./:;?@\[\\\]^_`{|}~]',
re.UNICODE)):
text = punctuation_re.sub(' ', text)
tokens = nltk.word_tokenize(text)
return [stemmer.stem(x) for x in tokens
if x.lower() not in stop_set and x[0] not in uppercase]
X = []
y = []
connection = sqlite3.connect('novels.sqlite')
for row in connection.cursor().execute(
"""SELECT plot, was_bestseller FROM Novels
WHERE year >= 1941 AND plot IS NOT NULL"""):
X.append(row[0])
y.append(row[1])
connection.close()
X_train, X_test, y_train, y_test = (
cross_validation.train_test_split(X, y, test_size=0.3))
model = pipeline.Pipeline(
[('tfidf', text.TfidfVectorizer(
lowercase=False, tokenizer=Tokenize)),
('logistic', linear_model.LogisticRegression())])
model.fit(X_train, y_train)
The model can return the probability of being a best seller for any novel b with a plot summary:
logit( b ) = −4.6 + 2.5 tfidf(lawyer, b ) + 2.4 tfidf(kill, b ) + ⋯ − 1.5 tfidf(planet, b )
Pr(was_bestseller( b )|plot( b )) = e logit( b ) / (1 + e logit( b ) )
To put these coefficients in context, tfidf(lawyer, The Firm ) ≈ 0.06. As it happens, the model returns logit( b ) > 0, that is Pr(was_bestseller( b )|plot( b )) > 1/2 for no novel b from the train or test set. The highest probability, 0.39, is predicted for Cross Fire , indeed a best seller in December 2010. Only if I disable the normalization in TF–IDF or weaken the regularization in the logistic regression, I can overfit the model to the train set while for the test set both its precision and recall would be at most 20%. But, like I wrote in the introduction, this is not the point of this exercise. Let us look at the words with high absolute value of coefficients.
Apparently, it pays off to write legal thrillers: lawyer +2.5, case +2.4, law +1.5, client +1.3, jury +1.3, trial +1.3, attorney +1.0, suspect +1.0, judge +0.9, convict +0.8;
kill +2.4, murder +1.8, terrorist +1.2, shoot +1.1, body +1.1, die +1.0, serial +0.9, attack +0.9, assassin +0.8, kidnap +0.8, killer +0.8.
Political thrillers are not bad either: agent +1.4, politics +1.4, president +1.3, defector +1.2.
Business may be involved: firm +1.3, company +1.3, career +1.1, million +1.0, success +1.0, business +0.9, money +0.9.
Finally, the characters should have families: husband +1.4, family +1.3, house +1.2, couple +1.2, daughter +1.2, baby +1.1, wife +1.0, father +1.0, child +0.9, birth +0.8, pregnant +0.8, and use a car +1.5 and a phone +0.8.
The genres to avoid for prospective best-selling authors?
Sci-fi: planet −1.5, human −1.0, space −0.7, star −0.4, robot −0.3, orbit −0.3.
Children’s literature: boy −1.3, school −1.0, young −0.8, girl −0.8, youth −0.4, teacher −0.4, aunt −0.4, grow −0.4.
Geography and travels: village −1.0, city −1.0, ship −0.8, way −0.7, go −0.7, land −0.6, adventure −0.6, colony −0.5, native −0.5, follow −0.5, mountain −0.5, crew −0.5, forest −0.5, travel −0.5, inhabit −0.4, sail −0.4, road −0.4, map −0.3, tribe −0.3.
War: fight −1.0, warrior −0.6, war −0.6, weapon −0.5, soldier −0.5, army −0.5, ally −0.4, enemy −0.3, conquer −0.3.
Fantasy: magic −0.9, creature −0.5, magician −0.4, zombie −0.3, treasure −0.3, dragon −0.3.
History: princess −0.5, rule −0.5, kingdom −0.4, castle −0.4, century −0.4, ruler −0.3, palace −0.3 (for what it’s worth, A Game of Thrones only made it to the third place on the list so it does not count as a best seller).
Note that the code above ignores capitalized words. If it does not, the most significant words become the names of characters from best selling book series: Scarpetta +3.0, Stephanie +2.9, Ayla +2.0, etc., with additional insights like FBI +1.3, CIA +1.3, NATO +0.9, Soviet +0.9, or Earth −1.1.
Share on X (Opens in new window)
X
Share on Facebook (Opens in new window)
Facebook
This site uses Akismet to reduce spam. Learn how your comment data is processed.
Subscribe
Subscribed
marcinciura.wordpress.com
Already have a WordPress.com account? Log in now.
