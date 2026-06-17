---
source: "https://provi.me/cc-like-video-games"
hn_url: "https://news.ycombinator.com/item?id=48567001"
title: "Claude Code is a video game"
article_title: "Claude Code is a video game - provi"
author: "pro-vi"
captured_at: "2026-06-17T07:58:46Z"
capture_tool: "hn-digest"
hn_id: 48567001
score: 1
comments: 0
posted_at: "2026-06-17T07:29:55Z"
tags:
  - hacker-news
  - translated
---

# Claude Code is a video game

- HN: [48567001](https://news.ycombinator.com/item?id=48567001)
- Source: [provi.me](https://provi.me/cc-like-video-games)
- Score: 1
- Comments: 0
- Posted: 2026-06-17T07:29:55Z

## Translation

タイトル: クロード・コードはビデオゲームです
記事のタイトル: クロード・コードはビデオゲームです - provi
説明: Civilization をプレイし続けるループは、クロード コードの下で高解像度で実行されます。 RTS のように複数のエージェントを指揮するところまで、ゲーム デザインが移植されました。

記事本文:
クロード・コードはビデオゲームです - プロヴィ
provi ビルドについてのEN | 書き込み 「クロード・コード」はビデオゲームです
Civilization をプレイし続けるループは、クロード コードの下で高解像度で実行されます。 RTS のように複数のエージェントを指揮するところまで、ゲーム デザインが移植されました。
就寝時刻から3時間経過。簡単なバグ修正がリファクタリング、ブレインストーミング、そして並行して実行される機能ブランチに変わり、私は時間の経過に気づきませんでした。
これは、Civilization をプレイしているときにのみ発生していました。
「もう 1 ターン」は、Civ プレイヤーなら誰でも知っているミームです。コマンドを出し、政策カードを選び、首都の次の建物の列に並び、次のターンを押すと、世界が動き始めます。気が付くと、あなたは新鮮なパズルを手渡されています。見慣れた形でありながら、微妙な新しさが詰まっています。したがって、ループは継続します。
LLM 会話は同じループ上で実行されます。プロンプトを送信すると、モデルが応答し、その応答が次に返信するプロンプトになります。コーディング エージェントは、方程式にツールの使用を追加します。モデルは、ルールベースのコンピューター プレーヤーができることをはるかに超えて、最新の AI プレーヤーの高度な機能で関数を呼び出し、その結果に反応することができます。このやりとりは、単純な質疑応答というより遊びのように感じられ始めます。あなたが行動すると、世界が変わり、新しい状況が与えられます。
このループは卓上に十分近いので、ゲーム メカニクスを直接そこにドロップできます。 cc-dice はターンが終了するたびにサイコロを振り、会話が長くなるほどチャンスが大きくなります。 Natural 20 がセッションの奥深くに到達すると、停止フックはエージェントに何が起こったのかを振り返り、基礎となるパターンを抽出し、今後のシステムを改善するアーティファクトを生成するように促します。その下のループは D&D キャンペーンとそれほど変わらないため、サイコロは機能します。
メカニックの移植は最も文字通りです

借りています。より深いものはループ自体の形状をしており、ゲームで変化する最初の点は時間です。文明は、チェスの精神に基づいて、時間を個別のターンに分割し、空間を六角形のタイルに分割します。 Dota 2 は、無限に近い位置の粒度で座標平面上で連続時間で実行されます。表面下では、どちらのゲームも同じことを要求します。つまり、リソースを管理し、スペースを制御し、制約の下で最適な決定を下すことです。しかし、解像度、つまり動作の粒度は、プレイ全体の感触を変えます。
同じ解像度の変更により、チャットがクロード コードから分離されます。チャットはカード ゲームです。ハンドをプレイし、ターンを渡し、コントロールが戻るのを待ちます。ターンはアトミックであり、ターンが戻るまでロックアウトされます。クロード コードは連続時間に近い状態で実行されます。ターンの途中でステアリングを操作したり、ツール呼び出しを途中で中断したり、ツールがまだ動作している間に新しいコンテキストをフィードしたりすることができます。コードベースはナビゲート可能なマップになり、ツール呼び出しは複数のファイルに対して一度に実行される詳細なアクションになります。もう、約束されたターンが終わるまで待つ必要はありません。あなたは、まだ動いているシステムと継続的に接触しています。
ここまでの登りは、時間という 1 つの軸上にあります。カード ゲームまたは Civilization が順番に実行され、Dota 2 は継続的に実行されます。それでも、Dota 2 と StarCraft は同じ連続クロックを共有していますが、RTS が 2 番目の軸を提供するため、まったく似たものは感じられません。 Dota 2 では、あなたは 1 人のヒーローを操縦し、1 つのユニット、1 セットのクールダウン、実行する 1 つの役割に全神経を集中します。 StarCraft では軍隊を指揮することができます。各部隊に命令を出し、グループを引き離して嫌がらせをし、ある基地で生産の順番を待ちながら別の基地で戦闘が勃発し、注意は十数ものものに分散され、見ているかどうかに関係なくすべてが動き続けます。同じリアルタイム クロックを、はるかに多くの機能を一度に手元に置くことができます。
として

ingle Claude Code セッションは Dota 2 バージョンです。 1 人のエージェント、1 人のヒーロー、あなたの全注意は次の動きに集中します。一度に複数を実行するのが StarCraft バージョンです。 1 つにリファクタリングを渡し、もう 1 つをテストの作成に送り、移行の途中で 3 つ目をちらっと見て、順調に進んでいるかどうかを確認します。あなたは、一人の呪文詠唱者の操縦から軍隊の指揮へと移行します。どの部隊が今注意を必要とし、どの部隊があなたなしで進むことができ、どの部隊が惨事に向かっていますか。あなたは地図の上にいて、次にどこを見るかを決めています。
最後の部分が私が心配している部分です。あるエージェントは就寝時刻を3時間過ぎても私を起きさせてくれました。今は7つ走ってます。

## Original Extract

The loop that keeps you playing Civilization runs under Claude Code, at higher resolution. Game design ports over, down to commanding several agents like an RTS.

Claude Code is a video game - provi
provi write build about EN | 简 Claude Code is a video game
The loop that keeps you playing Civilization runs under Claude Code, at higher resolution. Game design ports over, down to commanding several agents like an RTS.
Three hours past bedtime. A quick bug fix had turned into a refactor, then a brainstorm, then a feature branch running in parallel, and I had not noticed the passing of time.
This used to only happen when I played Civilization.
“One more turn” is a meme any Civ player will recognize. You issue your commands, pick your policy cards, queue the next building in your capital, press next turn, and the world starts moving. Before you realize it you are handed a fresh puzzle: familiar in shape, filled with subtle novelties. So the loop continues.
LLM conversations run on that same loop. You send a prompt, the model responds, and its response becomes the next prompt you reply to. Coding agents add tool use to the equation: the model can call functions and react to their results with the sophistication of a modern AI player, well past anything a rule-based computer player could do. The exchange starts feeling more like play than simple Q&A. You act, the world changes, you are handed a new situation.
That loop is close enough to a tabletop that you can drop a game mechanic straight into it. cc-dice rolls a die each time a turn ends, and the chance grows the longer a conversation runs. When a Natural 20 lands deep into a session, the stop hook prompts the agent to reflect on what happened , extract the underlying pattern, and generate an artifact that improves the system going forward. The dice works, because the loop underneath was not much different from a D&D campaign.
Porting a mechanic is the most literal borrowing. The deeper one is in the shape of the loop itself, and the first thing games vary is time. Civilization divides time into discrete turns and space into hexagonal tiles, in the spirit of Chess. Dota 2 runs in continuous time across a coordinate plane with near-infinite positional granularity. Beneath the surface, both games demand the same thing: manage resources, control space, and make optimal decisions under constraints. Yet resolution, the granularity at which they operate, transforms the entire feel of play.
The same resolution shift separates chat from Claude Code. Chat is the card game: you play your hand, pass your turn, and wait for control to return. The turn is atomic, and you are locked out until it comes back. Claude Code runs closer to continuous time. You can steer mid-turn, interrupt a tool call halfway, feed new context while it is still working. The codebase becomes a navigable map, and tool calls are granular actions taken across several files at once. You are not waiting out a committed turn anymore. You are in continuous contact with a system that is still moving.
That climb so far is on one axis: time. A card game or Civilization takes turns, and Dota 2 runs continuous. Yet Dota 2 and StarCraft share the same continuous clock and still feel nothing alike, because an RTS hands you a second axis. In Dota 2 you pilot one hero, your full attention on a single unit, one set of cooldowns, one role to perform. StarCraft makes you command an army. You give orders to each unit, peel a group off to harass, queue production in one base while a fight breaks out in another, and your attention is spread across a dozen things that all keep moving whether or not you are looking at them. Same real-time clock, far more of it under your hands at once.
A single Claude Code session is the Dota 2 version. One agent, one hero, your whole attention resolving onto its next move. Running several at once is the StarCraft version. You hand one a refactor, send another off to write tests, glance at a third halfway through a migration and check whether it is on track. You shift from piloting a single spell-caster to commanding an army: which unit demands attention now, which can proceed without you, which is heading for disaster. You are above the map, deciding where to look next.
That last part is the one I worry about. One agent kept me up three hours past bedtime. I now run seven.
