---
source: "https://www.ruxu.dev/articles/ai/build-a-basic-ai-agent/"
hn_url: "https://news.ycombinator.com/item?id=48492273"
title: "Build a Basic AI Agent from Scratch"
article_title: "Build a Basic AI Agent From Scratch"
author: "scapecast"
captured_at: "2026-06-11T16:26:24Z"
capture_tool: "hn-digest"
hn_id: 48492273
score: 1
comments: 0
posted_at: "2026-06-11T16:08:20Z"
tags:
  - hacker-news
  - translated
---

# Build a Basic AI Agent from Scratch

- HN: [48492273](https://news.ycombinator.com/item?id=48492273)
- Source: [www.ruxu.dev](https://www.ruxu.dev/articles/ai/build-a-basic-ai-agent/)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T16:08:20Z

## Translation

タイトル: 基本的な AI エージェントを最初から構築する
記事のタイトル: 基本的な AI エージェントを最初から構築する

記事本文:
基本的な AI エージェントを最初から構築する
8分で読めます
·
人工知能
2026 年は間違いなく AI エージェントの年です。クロード コードのリリース以来、これらの AI エージェントの力は否定できないものになりました。 Claude Code、Codex、OpenCode は、今日の多くの開発者にとって必須です。 OpenClaw と Hermes は、多くの人々の AI アシスタントになりつつあります。エージェントは、Cowork などのツールを使用してナレッジワークにも参入しています。
この一連の投稿をフォローしていただければ、これらのエージェントが実際にどのように動作するかをよりよく理解するために、基本的な AI エージェントをゼロから構築します。内部にあるものを実際に理解するために、フレームワークやライブラリは使用せず、Python でエージェントを最初から作成します。これはエージェントをできるだけ早く出荷する方法ではありませんが、学習する方法です。
AI エージェントは、人工知能を使用して自律的に目標を達成するプログラムです。他のタイプのエージェントと同様に、環境を認識し、それについて推論し、それに基づいてアクションを実行します。プログラムは通常、目標に到達するまでループで実行されます。
ベアボーンエージェントには何が必要ですか?
エージェントを雇うために必要なものは 4 つだけです。
エージェントを実行し続けるためのループ。
有能な AI モデルへの LLM 接続。
ユーザー入力。ユーザーがエージェントに目標を伝える方法。
コンテクスト 。エージェントが何が起こったかを忘れないように、これまでの会話を続けてください。
繰り返しますが、これは可能な限り最も基本的なエージェントの実装にすぎません。今後の投稿では、さらにエキサイティングな機能を追加する予定です。
エージェントを構築するには、まずモデルにアクセスする必要があります。この例では、無料で実行でき、実際に自分のマシンで実行できるモデルを使用します。そのために、4B の実効パラメーターを持つモデル gemma4:e4b を実行する Ollama のローカル インスタンスを使用します。

。
OSをインポートする
openaiインポートからOpenAI
def get_llm_client():
OpenAI を返す (
Base_url = "http://localhost:11434/v1" ,
api_key = ""
)
def エージェントループ (クライアント):
メッセージ = [
{ "role" : "system" , "content" : "あなたは役に立つアシスタントです。" }
】
True の場合:
user_input = input ( "あなた: " )
user_input の場合。下 ( ) == "\\終了" :
休憩
メッセージ。 append ( { "ロール" : "ユーザー" , "コンテンツ" : user_input } )
応答 = クライアント 。チャット 。完成品。作成(
モデル = "gemma4" 、
メッセージ = メッセージ 、
温度 = 0.7 、
)
返信 = 応答。選択肢 [ 0 ] 。メッセージ 。内容
print ( f"アシスタント: { 返信 } " )
メッセージ。 append ( { "役割" : "アシスタント" , "コンテンツ" : 返信 } )
if __name__ == "__main__" :
client = get_llm_client ( )
エージェントループ (クライアント)
このコードは、このブログ シリーズの Github リポジトリで見つけて複製できます。
まず、get_llm_client 関数を使用して、Ollama のローカル インスタンスへの LLM 接続を作成します。
次に、メッセージ履歴配列を作成し、システム プロンプト内の AI アシスタントの基本的な命令から開始します。
ユーザー入力を取得し、それをユーザー メッセージとしてメッセージ履歴に追加します。
最後のユーザー メッセージを含む新しい会話全体を AI モデルに送信し、AI モデルからの応答を要求します。
AI モデルの応答は会話履歴に追加されます。
ユーザーが \exit と入力するまで、ループは永久に実行されます。
このエージェント ループを実行すると、順番にこのエージェントに質問することができます。このエージェントは外部情報にアクセスできないため、内部の知識に基づいてのみ回答できます。
$ python エージェント.py
あなた: ドイツの首都はどこですか?
アシスタント: ドイツの首都は **ベルリン** です。
あなたが構築したもの
これは最も基本的な AI エージェントですが、まだ非常に不完全です。現時点では単なるチャットボットです

これにより、モデルの知識の範囲内にあるものはすべて答えられます。しかし、依然として環境と対話することはできません。ファイルの読み取りや書き込み、コマンドの実行、クエリに答えるための検索の実行はできません。
次のステップでは、エージェントがその環境内でアクションを開始できるようにするツールをエージェントに提供します。ここからが興味深いところであり、AI エージェントの可能性が明らかになり始めます。
このシリーズの次のパートでは、エージェント ループにツール呼び出しを追加します。それではまた。

## Original Extract

Build a Basic AI Agent From Scratch
8 minute read
·
Artificial Intelligence
2026 is without a doubt the year of AI agents. Since the release of Claude Code, the power of these AI agents has become undeniable. Claude Code, Codex, OpenCode are a must for many developers nowadays. OpenClaw and Hermes are becoming many people's AI assistants. Agents are also breaking into knowledge work with tools like Cowork.
If you follow me in this series of posts, we will build a basic AI agent from scratch in order to better understand how these agents actually work. For the purpose of actually understanding what's under the hood, we won't be using frameworks or libraries, we will write the agent from scratch in Python. This is not how you ship an agent as fast as possible, but it is how you learn.
An AI agent is a program that uses artificial intelligence to autonomously achieve a goal. Like any other type of agent, it perceives its environment, reasons about it, and takes action on it. The program usually runs in a loop until a goal is reached.
What Does a Barebones Agent Need?
You only need four things to have a working agent:
A loop to keep the agent running.
An LLM connection to a capable AI model.
User input . A way to let the user communicate the goal to the agent.
Context . Keep conversation so far so the agent doesn't forget what has happened.
Again, this is just for the most basic agent implementation possible. In future posts we will be adding more exciting features into it.
To build the agent, first you will need to have access to a model. For this example, I will be using a model that is free to run and can actually run in your own machine. For that I will be using a local instance of Ollama, which is running gemma4:e4b , a model with 4B effective parameters.
import os
from openai import OpenAI
def get_llm_client ( ) :
return OpenAI (
base_url = "http://localhost:11434/v1" ,
api_key = ""
)
def agent_loop ( client ) :
messages = [
{ "role" : "system" , "content" : "You are a helpful assistant." }
]
while True :
user_input = input ( "You: " )
if user_input . lower ( ) == "\\exit" :
break
messages . append ( { "role" : "user" , "content" : user_input } )
response = client . chat . completions . create (
model = "gemma4" ,
messages = messages ,
temperature = 0.7 ,
)
reply = response . choices [ 0 ] . message . content
print ( f"Assistant: { reply } " )
messages . append ( { "role" : "assistant" , "content" : reply } )
if __name__ == "__main__" :
client = get_llm_client ( )
agent_loop ( client )
You can find and clone this code in this blog series' Github repo .
First, we use the get_llm_client function to create a LLM connection to the local instance of Ollama.
Then, we create the message history array, starting it with basic instructions for the AI assistant in the system prompt.
We take the user input and we append it as a user message to the message history.
We send the new whole conversation, including the last user message, to the AI model, requesting a response from it.
The AI model response is appended to the conversation history.
The loop runs forever until the user types \exit .
If we run the this agent loop, we will be able to take turns and ask this agent questions. This agent doesn't have access to outside information, so it will only be able to answer based on its internal knowledge:
$ python agent.py
You: What's the capital city of Germany?
Assistant: The capital city of Germany is **Berlin**.
What You've Built
This is the most basic AI agent possible, but is still very incomplete. Right now it is just a chatbot that will answer anything that falls within the model's knowledge. But it still cannot interact with its environment. It can not read or write files, execute commands or do searches to help it answer your queries.
The next step is giving the agent tools to allow it to start taking actions in its environment. That's where things get interesting, and the potential of AI agents starts to become apparent.
In the next part of this series, we'll add tool calling to our agent loop. See you then.
