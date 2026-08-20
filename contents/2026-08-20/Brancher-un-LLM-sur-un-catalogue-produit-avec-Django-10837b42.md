---
source: "https://tonydevweb.com/articles/brancher-llm-catalogue-produit.html"
hn_url: "https://news.ycombinator.com/item?id=49378450"
title: "Brancher un LLM sur un catalogue produit avec Django"
article_title: "Brancher un LLM sur un catalogue produit avec Django"
image: "https://tonydevweb.com/static/developpeur-python.png"
author: "tonydevweb"
captured_at: "2026-08-20T19:24:44Z"
capture_tool: "hn-digest"
hn_id: 49378450
score: 5
comments: 1
posted_at: "2026-08-20T18:39:16Z"
tags:
  - hacker-news
  - translated
---

# Brancher un LLM sur un catalogue produit avec Django

- HN: [49378450](https://news.ycombinator.com/item?id=49378450)
- Source: [tonydevweb.com](https://tonydevweb.com/articles/brancher-llm-catalogue-produit.html)
- Score: 5
- Comments: 1
- Posted: 2026-08-20T18:39:16Z

## Translation

タイトル: Django を使用して LLM を製品カタログに接続する
説明: の返品

記事本文:
Django を使用して LLM を製品カタログに接続する
Django を使用して LLM を製品カタログに接続する
Django を使用して LLM を製品カタログに接続する
診断エージェントのスペアパーツ市場への統合に関するフィードバック。スタック: 非同期ビューの Django、PostgreSQL フルテキスト、Redis、LLM API。
スペアパーツ市場には語彙の問題があります。探しているもの (「タイミング チェーン テンショナー」) を知っている購入者は、すでに古典的な全文検索を利用できます。入力すると、検索されます。問題は他の訪問者からのもので、「寒いときにバイクがアイドリングで失速する」と書いている人です。彼は部屋の名前を知らず、検索しても結果はゼロで、再び部屋を出てしまいました。症状とカタログ参照の間には、翻訳の層が欠落しています。
強化された研究ではなく LLM を使用する理由
3 つのアプローチが可能でした。手動のルックアップ テーブル。信頼性は高いが厳格で、計画された内容のみをカバーします。埋め込みによるセマンティック検索はエレガントですが、間違った問題を解決します。つまり、症状と部品の間のリンクは語彙的ではなく、因果関係です。または、構造化された出力を備えた LLM では、バックエンドがカタログとブリッジしながらモデルが推論します。
これは決定的なニュアンスを伴って保持された 3 つ目です。LLM はカタログを決して参照しません。標準化された診断と検索用語を生成するだけです。実際の株式との関係は確定的なコードのままです。この分離は、プロジェクトのアーキテクチャ上の最も重要なポイントです。モデルは、存在しない製品を幻覚させたり、価格をでっち上げたり、入手可能性を約束したりすることはできません。
フロリダ州

ux は意図的に線形です。エ​​ージェントの連鎖、関数呼び出し、RAG はありません。 1 つの LLM 呼び出し、1 つの変換、N 個の SQL クエリ。注意すべき点は、キャッシュはブラウザに送信される最終応答ではなく、LLM からの応答にあるということです。診断は時間の経過とともに安定しますが、在庫は安定しません。今朝オンラインに掲載された製品は、テキストが 6 日間キャッシュされた診断にすぐに表示されます。
訪問者 ──► 非同期ビュー
│
§─► レート制限 (Redis)
§─► LLM 応答キャッシュ (Redis、7 日間)
│ └─ ミス ─► API → 構造化 JSON
§─► カタログ検索（SQL、クエリごと）
└─► ベースログ
モデルの出力を制約する
システム プロンプトは、前後にテキストを含まない厳密な JSON を強制します。最も重要なフィールドは、ユーザーには最も見えにくいフィールドでもあります。search_name はデータベースのクエリにのみ使用されます。最初のバージョンでは、「燃料インジェクター」や「スパーク プラグ」などの説明的な用語が使用されていましたが、「2006 ～ 2008 kawasaki の er6 n スパーク プラグ」などのカタログ タイトルと一致することはありませんでした。したがって、明示的な制約: 単数、可能な限り短い用語。
# プロンプト.py
SYSTEM_DIAGNOSTIC = """あなたは専門の整備士です。
顧客は症状を説明し、場合によっては自分の車についても言及します。
JSON のみで回答し、前後のテキストは省略します。
{
"vehicle": "メッセージから抽出されたモデル、または null",
"診断": "わかりやすい説明 (最大 100 文字)"、
"可能性のある個数": [
{"検索名": "カタログ用語",
"reason": "最大 40 単語",
"緊急": "高|中|低"}
]、
"security_advice": "関連する場合はアラート、そうでない場合は null"
}
最大 4 個、確率順に並べられます。
「search_name」には、単数形の最短の用語を使用します。
「で

「燃料インジェクター」ではなく「インジェクター」、「点火プラグ」ではなく「プラグ」。
「」
この教訓は一般化可能です。LLM 出力フィールドが検索をフィードするときは、自然形式ではなくインデックス形式に制限する必要があります。これらは 2 つの異なるレジスタであり、モデルはそれを推測しません。
プロジェクト内のすべてのビューは非同期であるため、いくつかの注意が必要です。 HTTP クライアントは非同期である必要があります。そうでないと、イベント ループが生成の 3 ～ 6 秒間ハングします。レート制限とキャッシュは Django の非同期メソッド (cache.aget、cache.aset) を経由し、最終的なログは acreate を使用します。
#views.py
クラス AgentDiagnosticView(ビュー):
async def post(self, request):
body = json.loads(request.body)
症状 = body.get("症状", "").strip()[:500]
len(症状) < 5の場合:
return JsonResponse({"error": "症状を指定してください。"}, status=400)
ip = get_ip(リクエスト)
limit_key = f"diag:limit:{ip}"
カウンタ = 待機キャッシュ.aget(limit_key, 0)
# ... カタログを検索してから JsonResponse
非同期ビューに特有の 3 つの落とし穴に注意してください。 LoginRequiredMixin は使用できません。ビューが実行される前に request.user に同期的にアクセスし、 SynchronousOnlyOperation が発生します。これを await request.auser() に置き換えます。次に、非同期クエリセットで .union() を実行すると、同じエラーが発生します。最後に、クエリセットはクエリセットのままでなければなりません。スライスすると常にクエリセットが生成されるため、非同期で反復可能のままですが、list() によって実体化された Python リストはそうではなくなります。
検索では PostgreSQL のネイティブのフルテキストが使用され、その場でベクトルが構築されます。動作するモーターと空に戻るモーターの違いは 3 つの詳細によって決まります。フランス語の構成では、q を使用せずにステミングを有効にします。

uoi の「candle」が「candle」と一致しないため、失敗率が爆発的に増加します。 search_type="websearch" を指定すると、クエリが寛容になります。デフォルト モードではすべての用語が存在する必要があります。さらに Q(アイコンを含む) が全文に欠けている部分を補います。
#views.py
ベクトル = (
SearchVector("製品_タイトル", 重み = "A", config = "フランス語")
+ SearchVector("製品説明", 重み = "B", config = "フランス語")
)
非同期デフォルト検索(用語、制限=3):
sq = SearchQuery(用語、config="フランス語"、search_type="websearch")
...
関連性のある 3 つの失敗
その結果、最初の機能バージョンが返されました。悪い結果。プレーンモードでは、カタログに部品が含まれているにもかかわらず、車両モデルを含むすべての条件が 4 ブロック中 3 ブロックをクリアする必要がありました。訂正: ウェブ検索に切り替え、車両モデルなしのフォールバックを追加しました。
この修正後、すべてのブロックが埋まりました。良すぎます。 「温度センサー」については、エンジンには速度センサー、落下センサー、もう 1 つの速度センサーが用意されていました。3 つのセンサーがありましたが、どれも適切なものはありませんでした。原因は、最初の重要な単語だけに対する最後の手段と思われるカスケード フォールバックでした。そして、「センサー」は明らかにカタログ内のすべてのセンサーと一致します。訂正: このフォールバックを削除し、最小ランクしきい値を導入しました。主題から外れた部分よりも、「アラートの作成」リンクを含む空のブロックの方が優れています。最初のブロックはリードを生成し、2 番目のブロックは論争を生成します。
3 番目の失敗は最も危険です。 「タイミング チェーン」（エンジンの内部）と「セカンダリ チェーン」（トランスミッション）は互いに何の関係もありませんが、フランス語のステミングでは両方に「チェーン」が含まれるため区別できません。閾値調整なし

全文スコアが正当に高いため、この問題は修正されません。ドメインごとの除外リストが必要でした。
#views.py
寄生虫 = {
"ディストリビューション": ["リア", "セカンダリ", "ケーシング", "保護"],
「ブレーキ」: [「レバー」、「ホース」、「瓶」]、
}
キーの場合、PARASITES.items() の除外:
term. lower() のキーの場合:
除外の単語:
qs = qs.exclude(product_title__icontains=word)
それはエレガントではありません。これはハードコーディングされたビジネス知識であり、維持する必要があります。しかし、これは 2 つの異なる概念が同じ語彙を共有する場合にのみ機能し、この状況は技術カタログでは例外ではなく標準です。
これらの失敗の共通の原因
これら 3 つの問題は、カタログが小さすぎるという単一の原因を共有しています。 4 つまたは 5 つのモデルに数百の参考資料が集中しているため、しきい値の設定や除外範囲の細かさに関係なく、カバーされていない車両で訪問者が互換性のある部品を見つけることはできません。このタイプのエージェントは生成器ではなく乗算器です。変換すべきものがあった場合、トラフィックをコンバージョンに変換します。スパースカタログでは、主にトラフィックを空のページに変換します。
プロンプト内の HTML または構造化された JSON
LLM が表示を目的としたコンテンツを作成するときは常に、この疑問が生じます。 1 つ目のアプローチは、モデルから HTML を直接リクエストし、それを DOM に挿入することです。実装は即座に行われ、既存の CSS クラスを再利用しますが、結果は不安定です。モデルはドリフトし、クラスを追加し、終了タグを忘れ、ネストが異なります。レイアウトが予期せず壊れ、派生した HTML はキーの有効期間中キャッシュされたままになります。
2回目のAP

close は構造化された JSON を要求し、フロントにマークアップを構築させます。これにより、構造が保証され、各フィールドが個別にエスケープされ、要素ごとの処理 (オフセット アニメーション、タイプライター効果) が簡単になり、プロンプトに触れたりキャッシュを削除したりすることなくマークアップを進化させることができます。ステップの番号付けはモデルではなくコードに基づいているため、場合によっては間違ってしまうことがあります。
覚えておくべき原則は 1 つの文にあります。LLM はコンテンツを生成し、コードは構造を生成します。プレゼンテーションをモデルに委ねるたびに、不安定性の原因を導入して 20 行の JavaScript を節約します。
消費量はリクエストごとに約 300 の入力トークンと 300 の出力トークンです。エントリーレベルのモデルでは、これは診断 1 件あたり約 0.0017 ユーロ、または月々のリクエスト 1 万件あたり約 20 ユーロに相当します。症状が非常に反復的であるため、キャッシュによってこの量が半分になります。「アイドル ストール」、「スポンジ状のブレーキング」、およびその他の数十の症状がループで発生し、観測されたヒット率は 40 ～ 60% です。
ここでは、高速で経済的なモデルで十分です。タスクは、複雑な推論ではなく、制約された出力を使用してマッピングすることです。上位モデルとの比較テストでは、コストが 2 倍になることを正当化するような差異は示されませんでした。一方で、毎月の支出上限は初日から設定する必要があります。エンドポイントが制限なく有料 API を呼び出すと、請求が発生するのを待っています。
Redis キャッシュを削除せずにプロンプ​​トを変更すると、いくつかの反復が失われ、そのため、古い応答を受信しながら新しいバージョンをテストしました。反射的に採用する: プロンプト ファイルの変更には、

一致するキーをパージしません。
初期バージョンのプロンプトは冗長すぎて、応答時間が 20 秒を超えていました。各フィールドの長さを明示的に制限することで、知覚される品質を損なうことなく、全体が 4 ～ 6 秒に短縮されました。
前面には 2 つの古典的なトラップがあります。 script タグの async 属性は、DOM にフォームが含まれる前にコードを実行します。出力ガードはサイレントにアクティブになり、何も起こりません。コンソール エラーも動作も発生せず、最悪の種類のバグも発生しません。延期しなければなりません。また、プログラムで送信をトリガーするには、dispatchEvent(new Event("submit")) はハンドラーをトリガーしません。form.requestSubmit() を呼び出す必要があります。
このパターンは、ユーザーの語彙が索引の語彙と一致しない技術カタログでも再現可能です。 LLM はカタログをまったく参照せず、単に標準化された用語を生成するだけです。出力は厳密に制限された JSON であり、HTML ではありません。キャッシュは世代に関するものであり、最終結果ではありません。レート制限は、運用が開始されるとすぐに適用されます。
2 つの点が他の点よりも重要です。コンバージョン トラッキングは機能よりも優先する必要があります。エージェントは構築するのは簡単ですが、評価するのは難しく、測定がなければ、有用性がわからないものを盲目的に最適化することになります。そして何よりも、在庫が依然として制限要因です。プロンプトの品質が製品の欠如を補うことはできません。これは最も重要な観察であり、私たちが発見したものです

[切り捨てられた]

## Original Extract

Retour d

Brancher un LLM sur un catalogue produit avec Django
Brancher un LLM sur un catalogue produit avec Django
Brancher un LLM sur un catalogue produit avec Django
Retour d'expérience sur l'intégration d'un agent de diagnostic dans une marketplace de pièces détachées. Stack : Django en vues asynchrones, PostgreSQL full-text, Redis, API LLM.
Une marketplace de pièces détachées a un problème de vocabulaire. L'acheteur qui sait ce qu'il cherche — « tendeur de chaîne de distribution » — est déjà servi par une recherche full-text classique : il tape, il trouve. Le problème vient de l'autre visiteur, celui qui écrit « ma moto cale au ralenti quand elle est froide ». Il ne connaît pas le nom de la pièce, la recherche lui renvoie zéro résultat, et il repart. Entre le symptôme et la référence catalogue, il manque une couche de traduction.
Pourquoi un LLM plutôt qu'une recherche améliorée
Trois approches étaient envisageables. Une table de correspondance manuelle, fiable mais rigide, qui ne couvre que ce qui a été prévu. Une recherche sémantique par embeddings, élégante mais qui résout le mauvais problème : le lien entre un symptôme et une pièce n'est pas lexical, il est causal. Ou un LLM avec sortie structurée, où le modèle raisonne pendant que le backend fait le pont avec le catalogue.
C'est la troisième qui a été retenue, avec une nuance déterminante : le LLM ne voit jamais le catalogue . Il produit un diagnostic et des termes de recherche normalisés, rien de plus. La mise en relation avec le stock réel reste du code déterministe. Cette séparation est le point d'architecture le plus important du projet : le modèle ne peut pas halluciner un produit inexistant, inventer un prix ou promettre une disponibilité.
Le flux est volontairement linéaire : pas de chaînage d'agents, pas de function calling, pas de RAG. Un appel LLM, une transformation, N requêtes SQL. Le point à noter est que le cache porte sur la réponse du LLM , pas sur la réponse finale envoyée au navigateur. Le diagnostic est stable dans le temps, le stock ne l'est pas : un produit mis en ligne ce matin apparaît immédiatement dans un diagnostic dont le texte est en cache depuis six jours.
Visiteur ──► Vue async
│
├─► Rate limit (Redis)
├─► Cache réponse LLM (Redis, 7 jours)
│ └─ miss ─► API → JSON structuré
├─► Recherche catalogue (SQL, à chaque requête)
└─► Log en base
Contraindre la sortie du modèle
Le prompt système impose un JSON strict, sans texte avant ni après. Le champ le plus important est aussi le moins visible pour l'utilisateur : nom_recherche ne sert qu'à interroger la base. Les premières versions renvoyaient des termes descriptifs comme « injecteur carburant » ou « bougies allumage », qui ne matchaient jamais des titres catalogue du type « Bougie er6 n de 2006 a 2008 kawasaki ». D'où la contrainte explicite : singulier, terme le plus court possible.
# prompts.py
SYSTEM_DIAGNOSTIC = """Tu es un mécanicien expert.
Le client décrit un symptôme, parfois en mentionnant son véhicule.
Réponds UNIQUEMENT en JSON, sans texte avant/après :
{
"vehicule": "modèle extrait du message, ou null",
"diagnostic": "explication claire (100 mots maximum)",
"pieces_probables": [
{"nom_recherche": "terme catalogue",
"raison": "40 mots maximum",
"urgence": "haute|moyenne|basse"}
],
"conseil_securite": "alerte si pertinent, sinon null"
}
Maximum 4 pièces, classées par probabilité.
Pour "nom_recherche", utilise le SINGULIER et le terme le plus court :
"injecteur" pas "injecteur carburant", "bougie" pas "bougies allumage".
"""
La leçon est généralisable : quand un champ de sortie LLM alimente une recherche, il faut le contraindre au format de l'index , pas au format naturel. Ce sont deux registres différents et le modèle ne le devine pas.
Toutes les vues du projet sont async, ce qui impose quelques précautions. Le client HTTP doit être asynchrone, sinon la boucle d'événements se bloque pendant les trois à six secondes de génération. Le rate limit et le cache passent par les méthodes async de Django ( cache.aget , cache.aset ), et le log final utilise acreate .
# views.py
class AgentDiagnosticView(View):
async def post(self, request):
body = json.loads(request.body)
symptome = body.get("symptome", "").strip()[:500]
if len(symptome) < 5:
return JsonResponse({"erreur": "Précisez le symptôme."}, status=400)
ip = get_ip(request)
cle_limite = f"diag:limite:{ip}"
compteur = await cache.aget(cle_limite, 0)
# ... recherche catalogue puis JsonResponse
Trois pièges spécifiques aux vues async méritent d'être signalés. LoginRequiredMixin est inutilisable : il accède à request.user de manière synchrone avant l'exécution de la vue, ce qui lève une SynchronousOnlyOperation — il faut le remplacer par await request.auser() . Ensuite, .union() sur un queryset async provoque la même erreur. Enfin, les querysets doivent rester des querysets : le slicing en produit toujours un, donc reste itérable en async, alors qu'une liste Python matérialisée par list() ne l'est plus.
La recherche utilise le full-text natif de PostgreSQL, avec le vecteur construit à la volée. Trois détails font la différence entre un moteur qui fonctionne et un moteur qui renvoie du vide. La configuration french active le stemming, sans quoi « bougies » ne matche pas « bougie » et le taux d'échec explose. Le search_type="websearch" rend la requête tolérante, là où le mode par défaut exige la présence de tous les termes. Et le Q(icontains) en complément rattrape ce que le full-text laisse passer.
# views.py
vecteur = (
SearchVector("produit_titre", weight="A", config="french")
+ SearchVector("produit_description", weight="B", config="french")
)
async def chercher(termes, limite=3):
sq = SearchQuery(termes, config="french", search_type="websearch")
...
Trois échecs de pertinence
La première version fonctionnelle remontait des résultats. De mauvais résultats. Le mode plain exigeait tous les termes, modèle de véhicule compris, ce qui vidait trois blocs sur quatre alors que le catalogue contenait bien les pièces. Correction : passage en websearch et ajout d'un repli sans le modèle de véhicule.
Après cette correction, tous les blocs se remplissaient. Trop bien. Pour « capteur température », le moteur proposait un capteur de vitesse, un capteur de chute et un autre capteur de vitesse : trois capteurs, aucun le bon. La cause était un repli en cascade qui cherchait en dernier recours sur le premier mot significatif seul — et « capteur » matche évidemment tous les capteurs du catalogue. Correction : suppression de ce repli et introduction d'un seuil de rang minimal. Mieux vaut un bloc vide avec un lien « créer une alerte » qu'une pièce hors sujet : le premier génère un lead, le second génère un litige.
Le troisième échec est le plus insidieux. « Chaîne de distribution » (interne au moteur) et « chaîne secondaire » (transmission) n'ont rien à voir, mais le stemming français les rend indistinguables puisque les deux contiennent « chaîne ». Aucun réglage de seuil ne résout ça, parce que le score full-text est légitimement élevé. Il a fallu une liste d'exclusions par domaine.
# views.py
PARASITES = {
"distribution": ["arriere", "secondaire", "carter", "protection"],
"frein": ["levier", "durite", "bocal"],
}
for cle, exclusions in PARASITES.items():
if cle in termes.lower():
for mot in exclusions:
qs = qs.exclude(produit_titre__icontains=mot)
Ce n'est pas élégant : c'est de la connaissance métier codée en dur, qu'il faudra maintenir. Mais c'est la seule chose qui fonctionne quand deux concepts distincts partagent le même vocabulaire, et cette situation est la norme dans les catalogues techniques, pas l'exception.
La cause commune de ces échecs
Ces trois problèmes partagent une origine unique : le catalogue était trop petit. Avec quelques centaines de références concentrées sur quatre ou cinq modèles, un visiteur qui arrive avec un véhicule non couvert ne trouvera jamais de pièce compatible, quel que soit le réglage du seuil ou la finesse des exclusions. Un agent de ce type est un multiplicateur, pas un générateur : il transforme du trafic en conversions quand il y a de quoi convertir. Sur un catalogue clairsemé, il transforme surtout du trafic en pages vides.
HTML dans le prompt ou JSON structuré
La question revient dès qu'un LLM produit du contenu destiné à l'affichage. La première approche consiste à demander directement du HTML au modèle et à l'injecter dans le DOM. L'implémentation est immédiate et réutilise les classes CSS existantes, mais le résultat est instable : le modèle dérive, ajoute une classe, oublie une balise fermante, imbrique différemment. La mise en page casse de façon imprévisible, et le HTML dérivé reste en cache pendant toute la durée de vie de la clé.
La seconde approche demande un JSON structuré et laisse le front construire le markup. La structure est alors garantie, chaque champ est échappé individuellement, le traitement par élément devient trivial — animation décalée, effet machine à écrire — et le markup peut évoluer sans toucher au prompt ni purger le cache. La numérotation des étapes vient du code plutôt que du modèle, qui se trompe parfois.
Le principe à retenir tient en une phrase : le LLM produit du contenu, le code produit de la structure . Chaque fois qu'on confie la présentation au modèle, on introduit une source d'instabilité pour économiser vingt lignes de JavaScript.
La consommation tourne autour de 300 tokens en entrée et 300 en sortie par requête. Sur un modèle d'entrée de gamme, cela représente environ 0,0017 € par diagnostic, soit une vingtaine d'euros pour dix mille requêtes mensuelles. Le cache divise ce montant par deux, car les symptômes sont extrêmement répétitifs : « cale au ralenti », « freinage spongieux » et quelques dizaines d'autres reviennent en boucle, avec un taux de hit observé entre 40 et 60 %.
Un modèle rapide et économique suffit largement ici : la tâche est du mapping avec sortie contrainte, pas du raisonnement complexe. Les tests comparatifs avec un modèle supérieur n'ont pas montré d'écart justifiant le doublement du coût. En revanche, un plafond de dépense mensuel doit être posé dès le premier jour — un endpoint qui appelle une API payante sans limite est une facture qui attend de se produire.
Plusieurs itérations ont été perdues à modifier le prompt sans purger le cache Redis, donc à tester une nouvelle version en recevant les anciennes réponses. Le réflexe à adopter : toute modification du fichier de prompts s'accompagne d'une purge des clés correspondantes.
Les premières versions du prompt étaient trop verbeuses, ce qui portait les temps de réponse au-delà de vingt secondes. Contraindre explicitement la longueur de chaque champ a ramené l'ensemble à quatre ou six secondes sans perte de qualité perçue.
Côté front, deux pièges classiques. L'attribut async sur la balise script exécute le code avant que le DOM contienne le formulaire, la garde de sortie s'active silencieusement et rien ne se passe : aucune erreur console, aucun comportement, le pire type de bug. Il faut defer . Et pour déclencher une soumission par programme, dispatchEvent(new Event("submit")) ne déclenche pas le handler — c'est form.requestSubmit() qu'il faut appeler.
Le pattern est reproductible pour n'importe quel catalogue technique où le vocabulaire de l'utilisateur ne correspond pas à celui de l'index. Le LLM ne voit jamais le catalogue et se contente de produire des termes normalisés ; la sortie est du JSON strictement contraint, jamais du HTML ; le cache porte sur la génération, pas sur le résultat final ; le rate limit est en place dès la mise en production.
Deux points comptent plus que les autres. Le tracking de conversion doit précéder les fonctionnalités : un agent est facile à construire et difficile à évaluer, et sans mesure on optimise à l'aveugle quelque chose dont on ignore l'utilité. Et surtout, le stock reste le facteur limitant — la qualité du prompt ne compense jamais l'absence de produits. C'est le constat le plus important, et c'est celui qu'on décou

[truncated]
