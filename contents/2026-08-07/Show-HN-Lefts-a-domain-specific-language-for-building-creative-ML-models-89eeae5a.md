---
source: "https://nsmat.github.io/lefts/"
hn_url: "https://news.ycombinator.com/item?id=49209498"
title: "Show HN: Lefts – a domain specific language for building creative ML models"
article_title: "Lefts"
author: "niksmather"
captured_at: "2026-08-07T12:40:05Z"
capture_tool: "hn-digest"
hn_id: 49209498
score: 1
comments: 0
posted_at: "2026-08-07T12:38:54Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Lefts – a domain specific language for building creative ML models

- HN: [49209498](https://news.ycombinator.com/item?id=49209498)
- Source: [nsmat.github.io](https://nsmat.github.io/lefts/)
- Score: 1
- Comments: 0
- Posted: 2026-08-07T12:38:54Z

## Translation

タイトル: Show HN: Lefts – クリエイティブな ML モデルを構築するためのドメイン固有言語
記事タイトル: レフト
説明: ML モデル実験のためのコンポーザブル変換
HN テキスト: Lefts は、応用機械学習モデリング用の小規模なドメイン固有言語です。これは、予測モデルの構築を生業としており、トレイン/テスト パイプラインの構築やデータ漏洩の心配ではなく、モデルの動作に関する推論と創造的なアーキテクチャの構築に集中したい人を対象としています。これはシンプルですが非常に強力です。私は自分の仕事でこれを使用して、モデリングの新しい方法を模索し (幾何学モデルに関するチュートリアルをチェックしてください!)、機械学習エンジニアとしての最も面白くない部分を簡単に乗り越えています。また、内部ではクールな関数型プログラミングも行われています (設計哲学をチェックしてください!)。

記事本文:
左派
コンテンツにスキップ
左派
ホーム
検索を初期化しています
nsmat/左
ホーム
例
例
リフトとフィードによる分位回帰
早期停止を伴う LGBM ローリングリトレイン
フィード付き教師と生徒のモデル
APIリファレンス
APIリファレンス
コマンド
左: 創造的な機械学習のためのコンポーザブルな変換
Lefts は、単純な機械学習ワークフローから複雑な機械学習ワークフローを構築するための、非常にシンプルなドメイン固有言語です。お気に入りの機械学習モデルから始めて、Lefts 演算を使用して次のことを行うことができます。
複雑な相互検証およびハイパーパラメータ化手順を構築します。
モデルが別のモデルのフィーチャまたはターゲットを作成できるようにします。
上記の創造的な組み合わせ。
後続のモデルのフィッティング、評価、実験を元のモデルよりも複雑にすることなく。この実装は、優れた Polars DataFrame ライブラリの上に構築されています。
Lefts には 5 つのコマンドがあり、その名前が付けられています。
リフト: データの異なるサブセットにわたってモデルの複数のコピーをトレーニングします。
アンサンブル: 一連のモデルを取得し、それらを 1 つとして評価します。
調整: モデルがそのハイパーパラメーターを別のモデルから学習できるようにします。
フィード: 1 つのモデルの出力を別のモデルでフィーチャーまたはターゲットとして使用できるようにします。
S 分割: 指定されたトレーニング/テスト/検証分割でモデルをトレーニングします。
これらのコマンドは、以下によって定義される任意のモデルに適用できます。
Fit メソッド。トレーニング データからモデル パラメーターにマッピングします。
予測メソッド。モデル パラメーターとテスト データを予測にマッピングします。
Lefts コマンドは、これらの関数を新しい .fit および .predict に変換することにより、新しいモデルを作成します。この新しいモデルには .fit と .predict もあるため、さらに Lefts コマンドを使用して変換できます。
次のコードは、ローリング月次再トレーニング ワークフローを作成します。 Ri の 12 個のコピー

dge 回帰。各月の前のすべてのデータでトレーニングされ、その月のみで評価されます。
左からインポートリーフ、リフト
lefts.helpers から tabular_model をインポート
極座標をPLとしてインポート
sklearn.linear_model からインポート リッジ
日時を dt としてインポート
特徴 = [ "温度" 、 "湿度" 、 "時間" ]
ターゲット＝「需要」
日付 = [ dt .範囲 ( 1 , 13 ) の m の日付 ( 2024 , m , 1 )]
モデル = リフト (
葉 ( tabular_model ( Ridge 、 features = features 、 target = target )、 label = 'ridge' )、
名前 = 'monthly_retrain' ,
値 = 日付 、
train_filter = ラムダ d : pl 。 Col ( 'date' ) < d 、
test_filter = ラムダ d : pl 。列 ( '日付' ) 。 dt 。月 () == d .月、
集約_with = pl 。合体、
)
# 12 個の独立したモデルに適合
モデル。フィット (DF)
# 各行は、その月までにトレーニングされたモデルから予測を取得します
予測 = モデル 。予測 ( df )
ワークフローを実行する前に、model.print_tree() を使用して検査できます。
リフト 'monthly_retrain' (12 モデル): [2024-01-01, ..., 2024-12-01] ⇒ 合体 → "monthly_retrain"
└── 葉の「尾根」
より複雑な例については、ドキュメントの「例」セクションを参照してください。

## Original Extract

Composable transformations for ML model experimentation

Lefts is a small domain specific language for applied machine learning modelling. It is aimed at anyone that builds predictive models for a living and wants to focus on reasoning about model behaviour and building creative architectures, and not on building train/test pipelines or worrying about data leakage. It is simple but quite powerful - I have been using it in my own work to explore new ways of modelling (check out the tutorial on geometric models!), and to breeze past the least interesting parts of being a machine learning engineer. It also has some cool functional programming going on under the hood (check out the design philosophy!).

Lefts
Skip to content
Lefts
Home
Initializing search
nsmat/lefts
Home
Examples
Examples
Quantile regressions with lift and feed
LGBM rolling retrain with early stopping
Teacher-student models with feed
API Reference
API Reference
Commands
Lefts: composable transformations for creative machine learning
Lefts is a very simple domain specific language for building complex machine learning workflows from simple ones. Starting with your favourite machine learning models, you can use Lefts operations to:
Build complex cross validation and hyper-parametrisation procedures.
Allow a model to create features or targets for another model.
And any creative combination of the above.
Without making subsequent model fitting, evaluation or experimentation any more complex than it was with the original model. This implementation is built on top of the excellent Polars DataFrame library.
Lefts has five commands, which give it its name:
L ift: trains multiple copies of a model across different subsets of data.
E nsemble: Takes a set of models and makes them evaluate as one.
T une: Allows a model to learn its hyperparameters from another.
F eed: Allows the output of one model to be used as a feature or target by another.
S plit: Trains a model on a given train/test/validation split.
These commands can be applied to any model that is defined by:
a fit method, which maps from training data into the model parameters
a predict method, which maps from model parameters and test data into predictions.
A Lefts command creates a new model by transforming these functions into a new .fit and .predict. Because this new model also has a .fit and .predict, it can be transformed with further Lefts commands.
The following code creates a rolling monthly retrain workflow: twelve copies of a Ridge regression, each trained on all data before its month and evaluated on that month only.
from lefts import leaf , lift
from lefts.helpers import tabular_model
import polars as pl
from sklearn.linear_model import Ridge
import datetime as dt
features = [ "temp" , "humidity" , "hour" ]
target = "demand"
dates = [ dt . date ( 2024 , m , 1 ) for m in range ( 1 , 13 )]
model = lift (
leaf ( tabular_model ( Ridge , features = features , target = target ), label = 'ridge' ),
name = 'monthly_retrain' ,
values = dates ,
train_filter = lambda d : pl . col ( 'date' ) < d ,
test_filter = lambda d : pl . col ( 'date' ) . dt . month () == d . month ,
aggregate_with = pl . coalesce ,
)
# Fits 12 independent models
model . fit ( df )
# Each row gets the prediction from the model trained up to its month
predictions = model . predict ( df )
You can inspect the workflow before running it with model.print_tree() :
Lift 'monthly_retrain' (12 models): [2024-01-01, ..., 2024-12-01] ⇒ coalesce → "monthly_retrain"
└── Leaf 'ridge'
For more complex examples, see the 'Examples' section in the docs.
