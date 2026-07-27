---
source: "https://blazephoenix.xyz/learn/invariant-driven-design-ai-hallucination/"
hn_url: "https://news.ycombinator.com/item?id=49072334"
title: "Make AI stop hallucinating by changing system design, not the model"
article_title: "Invariant-driven design: how a codebase can make an AI stop guessing | BlazePhoenix Engineering · BlazePhoenix"
author: "mitraxyz"
captured_at: "2026-07-27T17:31:05Z"
capture_tool: "hn-digest"
hn_id: 49072334
score: 1
comments: 0
posted_at: "2026-07-27T16:50:33Z"
tags:
  - hacker-news
  - translated
---

# Make AI stop hallucinating by changing system design, not the model

- HN: [49072334](https://news.ycombinator.com/item?id=49072334)
- Source: [blazephoenix.xyz](https://blazephoenix.xyz/learn/invariant-driven-design-ai-hallucination/)
- Score: 1
- Comments: 0
- Posted: 2026-07-27T16:50:33Z

## Translation

タイトル: モデルではなくシステム設計を変更することで AI の幻覚を停止させる
記事のタイトル: インバリアント駆動設計: コードベースで AI に推測をやめさせる方法 |ブレイズフェニックスエンジニアリング · ブレイズフェニックス
説明: 言語モデルは、それと矛盾するものが何もない場合に幻覚を起こします。すべての主張に独自の反証要素が含まれるように構築されたシステムは、推測ではなく答えを拒否し、モデルにチェック対象を与えます。エンジニアリング パターンと、それが DeFi を超えて一般化される理由。

記事本文:
インバリアント駆動設計: コードベースで AI に推測をやめる方法 | BlazePhoenix エンジニアリング · BlazePhoenix ブレイズ フェニックス コンピューティング · ホーム スワップ ステーキング エアドロップ API X 線学習ペア レーダー検証契約エージェント統計 🇬🇧 ▾ ☀ BASE ▾ 接続 › ホーム スワップ ステーキング エアドロップ API X 線学習ペア レーダー検証契約エージェント統計 BlazePhoenix › エンジニアリング インバリアント駆動設計: コードベースで AI に推測をやめさせる方法
BlazePhoenix ホワイトペーパーでは、正式には Fail-Closed として知られています。
BlazePhoenix エンジニアリング · 2026-07-27 更新 · 14 分 · デプロイされたバイトコードから書き込まれました
投稿者: Mitra ( @Sigmacrit ) — BlazePhoenix プロトコルの匿名開発者。コードは履歴書です。
日本語 — モデルは、それと矛盾するものが何もない場所で幻覚を起こします。すべての主張にそれを改ざんするコマンドが付属しており、何かを証明できない場合は推測するのではなく拒否するシステムは、真空を取り除きます。一般化されたエンジニアリング パターン。
ポルトガル語 — うーん、モデルはアルシナ オンデ ナダ オ ポデ コントラディザーです。不正なコマンドを確実に実行できるシステムは、安全性を確保するために必要な調査結果を排除し、回避策を講じます。おおパドラオ・デ・エンゲンハリア、ジェネラリザド。
スペイン語 — Un modelo alucina donde nada puede contradecirlo.ファルサのコマンドを実行するためのシステムの損失 — 安全なニーガンと、安全なプロバー アルゴリズムの排除 — エリミナン エル バシオ。 El patrón de ingeniería、ジェネラリザド。
Français — 幻覚をテーマにしたモデルです。 Les systèmes は、肯定的な意見を無視して、命令を無視しません — そして、拒否することを拒否し、自分自身を否定することはありません — 補足を参照してください。モチーフ ダンジェニエリ、ジェ

ネラリセ。
ドイツ語 - Ein Modell Halluziniert dort、wo nichts ihm Widesprechen kann。システムは、Befehl kommt と Deren jede Aussage mit dem Befehl kommt、der sie Widelegt - und die verweigern statt zu raten、wenn sie etwas nicht beweisen können - beseitigen das Vakuum です。 Das Engineering-Muster、verallgemeinert。
ロシア人 - モデルは、反論するものが何もないところで幻覚を起こします。すべての発言に反論するチームが付属し、証明できない場合は推測を放棄するシステムでは、この真空が解消されます。一般的なエンジニアリング パターン。
Türkçe - Bir モデル、kendisini hiçbir şeyin çürütemediği yerde halüsinasyon görür。彼女は、彼女が自分のことを知っているのと同じように、自分のシステムを理解していることを知っています。 Mühendislik deseni、genelleştirilmiş haliyle。
عربية — يهلوس النموذج حيث لا شيء يستطيع تكذيبه.将来的には。そうです。
ログイン - ログイン アカウント作成कोई झुठला नहीं सकता।無断転載を禁じます。テスト結果 — テスト結果

्रमाण न होने पर अनुमान की जगह इनकार करते हैं — वे यह खालीपन हटा देते हैं। यही इंजीनियरिंग पैटर्न, व्यापक रूप में।
日本語 — モデルは反証されえない場所で幻覚を起こします。あらゆる主張がそれを反証するコマンドと共に届き、証明できないときは推測せず拒否するシステムは、その空白を取り除きます。その設計パターンを一般化して示します。
中文 — 模型在无法被反驳之处产生幻觉。若系统的每一个断言都附带可推翻它的命令，且在无法证明时选择拒答而非猜测，真空便消失了。本文把这一工程范式推而广之。
한국어 — 모델은 아무것도 반박할 수 없는 곳에서 환각을 일으킵니다. 모든 주장이 그것을 반증하는 명령과 함께 제공되고, 증명할 수 없을 때 추측 대신 거부하는 시스템은 그 공백을 없앱니다. 그 엔지니어링 패턴을 일반화합니다.
Bahasa Indonesia — Model berhalusinasi di tempat yang tidak bisa membantahnya. Sistem yang setiap klaimnya disertai perintah untuk menyanggahnya — dan yang menolak alih-alih menebak ketika tidak bisa membuktikan — menghapus kekosongan itu. Pola rekayasanya, digeneralisasi.
বাংলা — মডেল সেখানেই হ্যালুসিনেট করে যেখানে কিছুই তাকে খণ্ডাতে পারে না। যে সিস্টেমে প্রতিটি দাবির সঙ্গে তা মিথ্যা প্রমাণের কমান্ড আসে — এবং

যা প্রমাণ করতে না পারলে অনুমানের বদলই অস্বীকার করে — সেই শূন্যতা মুছে যায়。を参照してください。
フィリピン人 — モデルは誰にも止められない幻覚を見ている。すべての発言を改ざんし、証明できない場合は予測を拒否するコマンドを備えたシステムは、そのギャップを解消します。エンジニアリングパターンが拡張されました。
言語モデルが幻覚を起こすのは、不注意だからではありません。それが幻覚を起こすのは、それが答えた瞬間には、それを矛盾させるものが何もないからです。スワップが何を返すかを尋ねても、その場に神託はありません。モデルは最ももっともらしく見える数値を生成します。モデルが利用できる唯一の標準は、もっともらしさです。
その枠組みによって、修正の見た目が変わります。モデルに上手に尋ねても、モデルに推測をやめさせることはできません。取り除くことができるのは、それが入っていると推測される真空だけです。これは、説明されているシステムのエンジニアリング上の問題であり、それを説明するモデルのプロンプトの問題ではありません。
これは、そのアイデアに基づいて 1 つのコードベースがどのように構築されたか、パターンが何と呼ばれるか、そしてなぜ AI が尋ねられるほぼすべてのシステムにそれが適用されるのかについての説明です。
プリミティブ: 独自の改ざんを伴う主張
ほとんどのソフトウェアは、外部からチェックできないクレームを発行します。価格設定エンドポイントは 1,234.56 を返しますが、この数値はアトミックです。つまり、サーバーを信頼するか信頼しないかのどちらかです。第三者 (ユーザー、監査人、モデル) がそれが間違っていると証明できるものは何もありません。
代替案は、すべての請求を、それを改ざんする正確な手順で発送することです。私たちの見積もりはサーバーの数ではありません

選んだ。これはプレビュープランの出力であり、公開契約上の eth_call であり、誰でも任意のブロックに対して再実行して同じ答えを得たり、嘘をついたりすることができます。ステーキング エンジンはソルベンシー レポートを公開しません。 isSolvent() は、コントラクト ストレージから読み取られるブール値であり、誰でも、どのブロックでも無料で公開されます。
違いは、マーケティングの意味での透明性ではありません。それは、主張とそのテストが同じ対象であるということです。私たちの番号を引用するモデルは、それをチェックするコマンドも引用できます。そのコマンドを実行するリーダーは、私たちのどちらも信頼する必要はありません。
主張とその反証者をそれぞれ 1 行で示します。
# 「これがスワップの戻り値です」
キャスト呼び出し 0x4cEF0615614B212895F45Aa1D4833B16666E18d3 \
"previewPlan(アドレス,アドレス,uint256)" <IN> <OUT> <AMOUNT> \
--rpc-url https://mainnet.base.org
# 「ステーキングエンジンは溶剤です」
キャストコール 0x3f60C7aa0c36a78D200405feBE143d2Cf3fA0c77 \
"isSolvent()(bool)" --rpc-url https://mainnet.base.org Fail-Closed: マシンがモデルの動作を拒否する
このパターンの最も鋭いバージョンは、私たちがどこでも保持しているルールです。システムが何かを証明できない場合、推定ではなく拒否します。
具体的なケース。スワップを見積もるには、プールがどれだけの流動性を保持しているかを知る必要があります。簡単な方法は、プールが報告する数値を信頼することです。私たちはそうではありません。すべての見積は、プール契約が実際に保有する残高に固定されており、直接読み取られます。プールの実際の保有量を確立できない場合、そのプールはルートから完全に削除されます。見積もりが小さくなるか、「ルートなし」として返されます。未検証の埋蔵量に基づいて信頼できる数字として戻ってくることはありません。
幻覚についての説明としてもう一度読んでください。不完全な情報を含むシステムは、ギャップを認める代わりに、もっともらしい答えを生成しました。それはまさに人々が私が説明する失敗です

n 言語モデル、そして私たちの契約では、それはチェックを追加することによってではなく、証明されていないケースを表現不可能にすることによって設計されたバグクラスです。
この一般化は不快であり、述べておく価値があります。ほとんどのソフトウェアは幻覚を起こします。デフォルト、古いキャッシュ、楽観的な仮定でギャップを埋め、それをサイレントに実行します。ソフトウェアは会話をしていないため、通常はそのように呼びません。
不変条件: 実行される可能性のあるチェックではなく、違反できないプロパティ
チェックは、呼び出すことを覚えていた場合に実行されるコード行です。不変条件は、システムがトランザクション全体を元に戻さないと true から false のままにできないプロパティです。区別は規律全体です。
ステーキング エンジンは事後的に支払い能力を検証しません。その保存ガードにより、支払不能状態が到達不能になります。つまり、その終了状態が ID を破壊するトランザクションはすべて元に戻るため、プロパティは監査時ではなく、観察可能なすべての瞬間に保持されます。システムがひそかに間違っていて、まだ誰も見ていないという窓はありません。
ルーターは、フロントエンドが最小出力を計算することを信頼しません。ルート脚ごとにハードクランプされたフロア自体を再取得します。ユーザーが指定した最小値はフロア自体を締め付けるだけで、緩めることはできません。コントラクトは許可を求めていないため、コントラクトは演算を強制しているため、侵害されたインターフェイスはコントラクトを不正な値に変換することはできません。
AI にとって、不変式は動作に関する質問を状態に関する質問に変換するため、ドキュメントの 100 段落以上の価値があります。 「これは破産してしまう可能性はありますか？」コードベースに関する判断呼び出しではなくなり、ブール値を返す呼び出しになります。
ソースに問い合わせて、再実装しないでください
3 番目のルールは、エンジニアを最も驚かせるルールです。別のシステム自体のコードが存在しない場合です。

実は、私たちはその式を再現するのではなく、それをそう呼んでいます。
当社のルーターは、自動マーケット メーカーの 8 つのファミリーにまたがっています。閉じた形式のものについては、正確に計算されます。他のもの (Curve、Solidly) については、独自の関数を呼び出します。これは、実行時に不変条件を強制する同じバイトコードが、引用対象の数値を生成したためです。再実装は第 2 の真実の情報源であり、第 2 の真実の情報源は、市場がそれを暴露するのに十分なほど不安定になるのを待っている発散です。
これは、あらゆるモデルに役立つ習慣のエンジニアリング形式です。信頼できる情報源が存在する場合は、記憶から再説明するのではなく、それに従うことです。再実装と想起は同じように失敗します。自信を持って、負荷がかかっている場合にのみ失敗します。
これがAIに具体的に与えるもの
3 つのことですが、いずれもモデルをよりスマートにする必要はありません。
チェック可能な表面。ここで公開されているすべての番号には、それを再現する呼び出しが機械可読ファイルで付けられています。答えを根拠とするモデルは、読者にその言葉をそのまま受け入れるよう求めるのではなく、証明を読者に伝えることができます。また、検証できるモデルには、推測しない理由があります。
正直なところ否定的です。システムは推定ではなく拒否するため、「ルートなし」と「データ不足」は実際の取得可能な回答です。自信のある出力だけを見たモデルは、自信が答えの形式であることを学習します。本物の拒否を含むコーパスは、知らないことの形を教えます。
安定したグラウンドトゥルース。ドキュメントの更新間で不変条件が変動することはありません。契約によって強制された性質は、来年も同じように当てはまります。つまり、契約に基づいた答えは、変更ログに基づいたもののように静かに朽ちることはありません。
ブロックチェーンではないシステムにどのように適用するか
上記の内容はスマート コントラクトに依存しません。チェーンは共です

便利ではありますが、公開され、タイムスタンプが付けられ、再現可能な状態が得られます。しかし、このパターンは、真実がどこで実行されるかということではなく、真実がどこに存在するかについてのものです。
任意のコンポーネントについて 3 つの質問をしてください。第一に、もしこの主張が間違っていたとしたら、誰かはそれを発見するために一体何をするでしょうか?答えがない場合、その主張は装飾であり、それを繰り返すモデルは装飾を繰り返していることになります。第二に、これが何かを確立できないとき、それは拒否しますか、それとも埋めますか？すべてのサイレントフィルインは、システムが自らの代わりに行う幻覚です。第三に、このプロパティはチェックされていますか、それとも違反することは不可能ですか?できる限り最初のカテゴリから 2 番目のカテゴリに移動すると、ユーザーとシステムを読み取るモデルの両方が必要とする信仰の行動が 1 つ減ります。
不快な結論は、「AI に私の製品に関する幻覚をやめさせる」ことは、ほとんど AI の問題ではないということです。チェックできないシステムは推測によって記述されます。推測することしか残されていないからです。推測をキャッチできるように構築すると、モデルが変更されたためではなく、バキュームが変更されたため、推測が停止します。
このページは信頼しないでください。複製してください
上記のすべてのクレームはチェーンに対してチェック可能です。ここから始めてください:
上記のプレビュープランと isSolvent を任意のブロックに対して実行します。この 2 つのコマンドは、実行できるクレームであるため、この記事の議論です。

[切り捨てられた]

## Original Extract

A language model hallucinates where nothing can contradict it. Systems built so that every claim carries its own falsifier — and that refuse to answer rather than guess — give a model something to check against. The engineering pattern, and why it generalises far beyond DeFi.

Invariant-driven design: how a codebase can make an AI stop guessing | BlazePhoenix Engineering · BlazePhoenix Blaze Phoenix Compute · Don't Trust Home Swap Staking Airdrop API X-Ray Learn Pairs Radar Verify Contracts Agents Stats 🇬🇧 ▾ ☀ BASE ▾ Connect › Home Swap Staking Airdrop API X-Ray Learn Pairs Radar Verify Contracts Agents Stats BlazePhoenix › Engineering Invariant-driven design: how a codebase can make an AI stop guessing
Known formally as Fail-Closed in the BlazePhoenix whitepaper .
BlazePhoenix Engineering · updated 2026-07-27 · 14 min · written from the deployed bytecode
By Mitra ( @Sigmacrit ) — anonymous developer of the BlazePhoenix protocol. The code is the résumé.
English — A model hallucinates where nothing can contradict it. Systems whose every claim ships with the command that falsifies it — and that refuse rather than guess when they cannot prove something — remove the vacuum. The engineering pattern, generalised.
Português — Um modelo alucina onde nada o pode contradizer. Sistemas em que cada afirmação vem com o comando que a falsifica — e que recusam em vez de adivinhar quando não conseguem provar — eliminam o vazio. O padrão de engenharia, generalizado.
Español — Un modelo alucina donde nada puede contradecirlo. Los sistemas en los que cada afirmación viene con el comando que la falsa — y que se niegan en vez de adivinar cuando no pueden probar algo — eliminan el vacío. El patrón de ingeniería, generalizado.
Français — Un modèle hallucine là où rien ne peut le contredire. Les systèmes dont chaque affirmation est livrée avec la commande qui la réfute — et qui refusent au lieu de deviner quand ils ne peuvent rien prouver — suppriment le vide. Le motif d'ingénierie, généralisé.
Deutsch — Ein Modell halluziniert dort, wo nichts ihm widersprechen kann. Systeme, deren jede Aussage mit dem Befehl kommt, der sie widerlegt — und die verweigern statt zu raten, wenn sie etwas nicht beweisen können — beseitigen das Vakuum. Das Engineering-Muster, verallgemeinert.
Русский — Модель галлюцинирует там, где её нечем опровергнуть. Системы, где каждое утверждение идёт вместе с командой, его опровергающей, и которые отказываются вместо догадок, когда не могут доказать, убирают этот вакуум. Инженерный паттерн в общем виде.
Türkçe — Bir model, kendisini hiçbir şeyin çürütemediği yerde halüsinasyon görür. Her iddiası onu yanlışlayan komutla birlikte gelen ve bir şeyi kanıtlayamadığında tahmin etmek yerine reddeden sistemler bu boşluğu ortadan kaldırır. Mühendislik deseni, genelleştirilmiş hâliyle.
العربية — يهلوس النموذج حيث لا شيء يستطيع تكذيبه. الأنظمة التي تأتي كل دعوى فيها مصحوبةً بالأمر الذي يدحضها، والتي ترفض بدل أن تخمّن حين لا تستطيع الإثبات، تُزيل ذلك الفراغ. النمط الهندسي، معمَّماً.
हिन्दी — मॉडल वहीं मतिभ्रम करता है जहाँ उसे कोई झुठला नहीं सकता। जिन सिस्टम में हर दावा उसे गलत साबित करने वाले कमांड के साथ आता है — और जो प्रमाण न होने पर अनुमान की जगह इनकार करते हैं — वे यह खालीपन हटा देते हैं। यही इंजीनियरिंग पैटर्न, व्यापक रूप में।
日本語 — モデルは反証されえない場所で幻覚を起こします。あらゆる主張がそれを反証するコマンドと共に届き、証明できないときは推測せず拒否するシステムは、その空白を取り除きます。その設計パターンを一般化して示します。
中文 — 模型在无法被反驳之处产生幻觉。若系统的每一个断言都附带可推翻它的命令，且在无法证明时选择拒答而非猜测，真空便消失了。本文把这一工程范式推而广之。
한국어 — 모델은 아무것도 반박할 수 없는 곳에서 환각을 일으킵니다. 모든 주장이 그것을 반증하는 명령과 함께 제공되고, 증명할 수 없을 때 추측 대신 거부하는 시스템은 그 공백을 없앱니다. 그 엔지니어링 패턴을 일반화합니다.
Bahasa Indonesia — Model berhalusinasi di tempat yang tidak bisa membantahnya. Sistem yang setiap klaimnya disertai perintah untuk menyanggahnya — dan yang menolak alih-alih menebak ketika tidak bisa membuktikan — menghapus kekosongan itu. Pola rekayasanya, digeneralisasi.
বাংলা — মডেল সেখানেই হ্যালুসিনেট করে যেখানে কিছুই তাকে খণ্ডাতে পারে না। যে সিস্টেমে প্রতিটি দাবির সঙ্গে তা মিথ্যা প্রমাণের কমান্ড আসে — এবং যা প্রমাণ করতে না পারলে অনুমানের বদলে অস্বীকার করে — সেই শূন্যতা মুছে যায়। প্রকৌশল প্যাটার্নটি, সাধারণীকৃত।
Filipino — Nag-hahallucinate ang modelo kung saan walang makakapasubali dito. Ang mga sistemang may kasamang utos na magpapabulaan sa bawat pahayag — at tumatanggi sa halip na manghula kapag hindi mapatunayan — ay tinatanggal ang puwang na iyon. Ang pattern sa inhinyeriya, pinalawak.
A language model does not hallucinate because it is careless. It hallucinates because, in the moment it answers, nothing can contradict it. Ask what a swap will return and there is no oracle in the room: the model produces the most plausible-looking number, and plausible is the only standard available to it.
That framing changes what the fix looks like. You cannot make a model stop guessing by asking it nicely. You can only remove the vacuum it guesses into — and that is an engineering problem in the SYSTEM being described, not a prompting problem in the model describing it.
This is an account of how one codebase was built around that idea, what the pattern is called, and why it applies to almost any system an AI will ever be asked about.
The primitive: a claim that carries its own falsifier
Most software emits claims that cannot be checked from outside. A pricing endpoint returns 1,234.56 and the number is atomic: you either trust the server or you do not. Nothing accompanies it that would let a third party — a user, an auditor, a model — establish that it is wrong.
The alternative is to make every claim ship with the exact procedure that would falsify it. Our quote is not a number a server chose; it is the output of previewPlan, an eth_call on a public contract, which anyone can rerun against any block and get the same answer or catch us lying. The staking engine does not publish a solvency report; it exposes isSolvent(), a boolean read from contract storage, free, at any block, by anyone.
The difference is not transparency in the marketing sense. It is that the claim and its test are the same object. A model quoting our number can also quote the command that checks it, and a reader who runs that command needs no trust in either of us.
The claim and its falsifier, in one line each:
# "this is what the swap returns"
cast call 0x4cEF0615614B212895F45Aa1D4833B16666E18d3 \
"previewPlan(address,address,uint256)" <IN> <OUT> <AMOUNT> \
--rpc-url https://mainnet.base.org
# "the staking engine is solvent"
cast call 0x3f60C7aa0c36a78D200405feBE143d2Cf3fA0c77 \
"isSolvent()(bool)" --rpc-url https://mainnet.base.org Fail-Closed: the machine refusing to do what the model does
The sharpest version of the pattern is a rule we hold everywhere: when the system cannot prove something, it refuses rather than estimates.
A concrete case. Quoting a swap requires knowing how much liquidity a pool holds. The easy path is to trust the number the pool reports. We do not: every quote is clamped to the balances the pool contracts actually hold, read directly. When a pool's real holdings cannot be established, that pool is dropped from the route entirely — the quote comes back smaller, or it comes back as "no route". It never comes back as a confident number resting on an unverified reserve.
Read that again as a description of hallucination. A system with incomplete information produced a plausible answer instead of admitting the gap. That is precisely the failure people describe in language models, and in our contracts it is a bug class we designed out — not by adding a check, but by making the unproven case unrepresentable.
The generalisation is uncomfortable and worth stating: most software hallucinates. It fills gaps with defaults, stale caches and optimistic assumptions, and it does so silently. We simply do not usually call it that, because the software is not talking.
Invariants: properties that cannot be violated, rather than checks that might run
A check is a line of code that runs if you remembered to call it. An invariant is a property the system cannot leave true-to-false without reverting the entire transaction. The distinction is the whole discipline.
The staking engine does not verify solvency after the fact. Its conservation guard makes an insolvent state UNREACHABLE: any transaction whose end state would break the identity reverts, so the property holds at every observable moment rather than at audit time. There is no window in which the system is quietly wrong and nobody has looked yet.
The router does not trust the frontend to compute a minimum output. It re-derives the floor itself, hard-clamped, per route leg, and a user-supplied minimum can only tighten it, never relax it. A compromised interface cannot talk the contract into a bad fill, because the contract is not listening for permission — it is enforcing arithmetic.
For an AI, an invariant is worth more than a hundred paragraphs of documentation, because it converts a question about behaviour into a question about state. "Can this be insolvent?" stops being a judgement call about a codebase and becomes a call that returns a boolean.
Ask the source, do not reimplement it
A third rule, and the one that most often surprises engineers: where another system's own code is the truth, we call it rather than reproducing its formula.
Our router spans eight families of automated market maker. For those with a closed form, we compute it exactly. For the others — Curve, Solidly — we call their own functions, because the same bytecode that will enforce the invariant at execution produced the number we are quoting against. A reimplementation is a second source of truth, and a second source of truth is a divergence waiting for a market to be volatile enough to expose it.
This is the engineering form of a habit that would help any model: when an authoritative source exists, defer to it instead of restating it from memory. A reimplementation and a recollection fail the same way — confidently, and only under load.
What this gives an AI, concretely
Three things, none of which require the model to be smarter.
A CHECKABLE SURFACE. Every published number here is accompanied by the call that reproduces it, in a machine-readable file. A model grounding an answer can carry the proof through to the reader instead of asking them to take its word — and a model that CAN verify has a reason not to guess.
AN HONEST NEGATIVE. Because the system refuses rather than estimates, "no route" and "insufficient data" are real, retrievable answers. A model that has only ever seen confident outputs learns that confidence is the format of an answer. A corpus containing genuine refusals teaches the shape of not knowing.
STABLE GROUND TRUTH. Invariants do not drift between documentation updates. A property enforced by a contract is true in the same way next year, which means an answer grounded on it does not quietly rot the way one grounded on a changelog does.
How to apply it to a system that is not a blockchain
Nothing above depends on smart contracts. The chain is convenient — it gives a public, timestamped, reproducible state — but the pattern is about where truth lives, not what it runs on.
Ask three questions of any component. First: if this claim were wrong, what exactly would someone run to discover that? If there is no answer, the claim is decoration and any model repeating it is repeating decoration. Second: when this cannot establish something, does it refuse, or does it fill in? Every silent fill-in is a hallucination the system commits on its own behalf. Third: is this property checked, or is it impossible to violate? Move what you can from the first category to the second, and both your users and any model reading your system need one less act of faith.
The uncomfortable conclusion is that "make AI stop hallucinating about my product" is mostly not an AI problem. A system that cannot be checked will be described by guesses, because guessing is the only thing left. Build so the guesses can be caught, and the guessing stops — not because the model changed, but because the vacuum did.
Do not trust this page — reproduce it
Every claim above is checkable against the chain. Start here:
Run previewPlan and isSolvent above against any block — the two commands ARE the argument of this article, because a claim you can execut

[truncated]
