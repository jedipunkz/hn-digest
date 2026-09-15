---
source: "https://qxotic.ai/"
hn_url: "https://news.ycombinator.com/item?id=49708628"
title: "Show HN: Jinfer – AI inference engine for the JVM. AI in a jar"
article_title: "Quixotic AI: AI sovereignty for the JVM"
image: "https://qxotic.ai/og.jpg"
author: "mukel"
captured_at: "2026-09-15T07:16:02Z"
capture_tool: "hn-digest"
hn_id: 49708628
score: 2
comments: 0
posted_at: "2026-09-15T06:42:32Z"
tags:
  - hacker-news
---

# Show HN: Jinfer – AI inference engine for the JVM. AI in a jar

- HN: [49708628](https://news.ycombinator.com/item?id=49708628)
- Source: [qxotic.ai](https://qxotic.ai/)
- Score: 2
- Comments: 0
- Posted: 2026-09-15T06:42:32Z

## Translation

Title: Show HN: Jinfer – AI inference engine for the JVM. AI in a jar
Article title: Quixotic AI: AI sovereignty for the JVM
Description: Run LLMs, vision, embeddings, and text-to-speech on the JVM at native speed. No Python runtime, ONNX bridges, or Docker sidecars.
HN text: Hi HN,
Some would call AI on the JVM quixotic. And so, Quixotic AI was born. jinfer is an inference engine for the JVM: chat, vision, audio, embeddings, reranking, and TTS. No Python runtime, no ONNX, no Docker containers, no sidecar process, no IPC. Finally, AI in a jar. The stack underneath is built for the JVM rather than bolted onto it: toknroll: pure-Java tokenizers, zero dependencies
gguf / safetensors: read and write llama.cpp and HuggingFace model formats
jam: quantized matmul kernels, competitive with llama.cpp on CPU
jota: Tensor API targeting Java, C, CUDA, HIP, Metal, OpenCL, and Mojo
It ships with integrations for Spring AI and LangChain4j, and is compatible with GraalVM Native Image for low-overhead, self-contained binaries with millisecond startup. This is an early release (CPU only): I'd especially like feedback on the API surface. Site: https://qxotic.ai
Repo: https://github.com/qxoticai/qxotic

Article text:
Quixotic AI
Stack
Examples
Benchmarks
AI sovereignty
for the JVM
Some call AI on the JVM quixotic.
We call it Quixotic AI.
Finally, AI in a jar.
Every component of the AI stack, a la carte.
AI inference engine. Chat, vision, embeddings and text-to-speech capabilities for the JVM, with Spring AI and LangChain4j integrations.
Fast quantized matrix-multiplication routines.
TikToken-compatible and customizable tokenizers for popular LLM models.
Multi-backend tensor engine, in the works. Java, C, CUDA, HIP, Metal, OpenCL, and Mojo. Write once, accelerate everywhere.
Pure Java, read and write support for llama.cpp's GGUF model format.
Pure Java, read and write support for HuggingFace's Safetensors format.
Runnable snippets using jbang .
jbang header + imports ///usr/bin/env jbang "$0" "$@" ; exit $?
//JAVA 25+
//RUNTIME_OPTIONS --add-modules jdk.incubator.vector --enable-native-access=ALL-UNNAMED
//DEPS com.qxotic:jinfer-bom:0.2.0@pom
//DEPS com.qxotic:jinfer-spring-ai com.qxotic:jinfer-models-all
//DEPS com.qxotic:jam-native com.qxotic:jam-vector
//DEPS org.slf4j:slf4j-nop:2.0.17
import com.qxotic.jinfer.spring.ai.*;
void main () {
try ( var model = JinferChatModel .builder()
.model( "LiquidAI/LFM2.5-350M-GGUF:Q8_0" )
.build()) {
System .out.println(model.call( "What is the capital of France?" ));
}
}
$ jbang Chat.java # "The capital of France is Paris."
jbang header + imports ///usr/bin/env jbang "$0" "$@" ; exit $?
//JAVA 25+
//RUNTIME_OPTIONS --add-modules jdk.incubator.vector --enable-native-access=ALL-UNNAMED
//DEPS com.qxotic:jinfer-bom:0.2.0@pom
//DEPS com.qxotic:jinfer-spring-ai
//DEPS com.qxotic:jinfer-kokoro
//DEPS com.qxotic:jam-native com.qxotic:jam-vector
//DEPS org.slf4j:slf4j-nop:2.0.17
import com.qxotic.jinfer.spring.ai. JinferSpeechModel ;
import java.nio.file.*;
void main () throws Exception {
try ( var tts = JinferSpeechModel .builder()
.model( "simonfxr/kokoro.cpp-GGUF:Q8_0" )
.companion( "voice" , "simonfxr/kokoro.cpp-GGUF/voices/kokoro-voice-af_heart.gguf" )
.build()) {
var line = "Turns out, matrix multiplications can talk. "
+ "The JVM just crunched a few billion numbers to say this." ;
Files .write( Path .of( "kokoro.wav" ), tts.call(line));
}
}
$ jbang TextToSpeech.java # kokoro.wav · seven seconds, in a warm voice · Kokoro-82M
jbang header + imports ///usr/bin/env jbang "$0" "$@" ; exit $?
//JAVA 25+
//RUNTIME_OPTIONS --add-modules jdk.incubator.vector --enable-native-access=ALL-UNNAMED
//DEPS com.qxotic:jinfer-bom:0.2.0@pom
//DEPS com.qxotic:jinfer-spring-ai com.qxotic:jinfer-models-all
//DEPS org.springframework.ai:spring-ai-client-chat:2.0.1
//DEPS com.qxotic:jam-native com.qxotic:jam-vector
//DEPS org.slf4j:slf4j-nop:2.0.17
import com.qxotic.jinfer.spring.ai.*;
import org.springframework.ai.chat.client. ChatClient ;
import org.springframework.core.io. UrlResource ;
import org.springframework.util. MimeType ;
void main () throws Exception {
try ( var gemma = JinferChatModel .builder()
.model( "unsloth/gemma-4-E2B-it-GGUF:Q4_K_M" )
.companion( "media" , "unsloth/gemma-4-E2B-it-GGUF/mmproj-BF16.gguf" )
.build()) {
// jfk.wav: JFK's own voice, inaugural address, January 20, 1961
var speech = new UrlResource ( "https://qxotic.ai/snippets/jfk.wav" );
var transcript = ChatClient .create(gemma).prompt()
.user(u -> u.text( "Transcribe this recording." )
.media( MimeType .valueOf( "audio/wav" ), speech))
.call().content();
System .out.println(transcript);
}
}
$ jbang Audio.java # "ask not what your country can do for you ..." · every word heard
jbang header + imports ///usr/bin/env jbang "$0" "$@" ; exit $?
//JAVA 25+
//RUNTIME_OPTIONS --add-modules jdk.incubator.vector --enable-native-access=ALL-UNNAMED
//DEPS com.qxotic:jinfer-bom:0.2.0@pom
//DEPS com.qxotic:jinfer-spring-ai com.qxotic:jinfer-models-all
//DEPS org.springframework.ai:spring-ai-client-chat:2.0.1
//DEPS com.qxotic:jam-native com.qxotic:jam-vector
//DEPS org.slf4j:slf4j-nop:2.0.17
import com.qxotic.jinfer.spring.ai.*;
import org.springframework.ai.chat.client. ChatClient ;
import org.springframework.core.io. UrlResource ;
import org.springframework.util. MimeTypeUtils ;
void main () throws Exception {
try ( var vision = JinferChatModel .builder()
.model( "LiquidAI/LFM2.5-VL-3B-GGUF:Q8_0" )
.companion( "media" , "LiquidAI/LFM2.5-VL-3B-GGUF/mmproj-LFM2.5-VL-3B-Q8_0.gguf" )
.build()) {
// The windmills of Consuegra, La Mancha: Don Quixote's "giants"
var windmills = new UrlResource ( "https://qxotic.ai/snippets/consuegra.jpg" );
var answer = ChatClient .create(vision).prompt()
.user(u -> u.text( "Don Quixote saw giants here. What do you see?" )
.media( MimeTypeUtils .IMAGE_JPEG, windmills))
.call().content();
System .out.println(answer);
}
}
$ jbang Vision.java # "several traditional windmills standing on a hill under a clear blue sky"
jbang header + imports ///usr/bin/env jbang "$0" "$@" ; exit $?
//JAVA 25+
//RUNTIME_OPTIONS --add-modules jdk.incubator.vector --enable-native-access=ALL-UNNAMED
//DEPS com.qxotic:jinfer-bom:0.2.0@pom
//DEPS com.qxotic:jinfer-spring-ai com.qxotic:jinfer-models-all
//DEPS com.qxotic:jam-native com.qxotic:jam-vector
//DEPS org.slf4j:slf4j-nop:2.0.17
import com.qxotic.jinfer.spring.ai. JinferEmbeddingModel ;
void main () {
try ( var emb = JinferEmbeddingModel .builder()
.model( "LiquidAI/LFM2.5-Embedding-350M-GGUF:Q8_0" )
.build()) {
var a = emb.embed( "AI on the JVM" );
var b = emb.embed( "the JVM thinks now" );
var c = emb.embed( "the python shed its skin" );
System .out.printf( "similar %.3f%n" , cosineSimilarity(a, b));
System .out.printf( "unrelated %.3f%n" , cosineSimilarity(a, c));
}
}
cosineSimilarity helper float cosineSimilarity ( float [] a, float [] b) {
float dot = 0 , na = 0 , nb = 0 ;
for ( int i = 0 ; i < a.length; i++) {
dot += a[i] * b[i]; na += a[i] * a[i]; nb += b[i] * b[i];
}
return ( float ) (dot / Math .sqrt(na * nb));
}
$ jbang Embed.java # similar 0.609 · unrelated 0.081
Benchmarking is hard ... take these numbers as indicative, not definitive, always measure yourself.
jinfer-bench vs. llama-bench on CPU, ↑ higher is better.
Built for the JVM from first principles
Local AI, end-to-end on the JVM: no sidecars, no services, no IPC.
No Python runtime, no ONNX, no glue code. Every component, from tokenizers to the inference engine, built for the JVM, not bolted onto it.
Ships as a single self-contained binary: millisecond startup, small footprint, no JVM at runtime.
Write Once, Accelerate Everywhere
One Tensor API, seven backends: Java, C, CUDA, HIP, Metal, OpenCL, and Mojo.
Built in 🇨🇭 Switzerland · 2026 Quixotic AI

## Original Extract

Run LLMs, vision, embeddings, and text-to-speech on the JVM at native speed. No Python runtime, ONNX bridges, or Docker sidecars.

Hi HN,
Some would call AI on the JVM quixotic. And so, Quixotic AI was born. jinfer is an inference engine for the JVM: chat, vision, audio, embeddings, reranking, and TTS. No Python runtime, no ONNX, no Docker containers, no sidecar process, no IPC. Finally, AI in a jar. The stack underneath is built for the JVM rather than bolted onto it: toknroll: pure-Java tokenizers, zero dependencies
gguf / safetensors: read and write llama.cpp and HuggingFace model formats
jam: quantized matmul kernels, competitive with llama.cpp on CPU
jota: Tensor API targeting Java, C, CUDA, HIP, Metal, OpenCL, and Mojo
It ships with integrations for Spring AI and LangChain4j, and is compatible with GraalVM Native Image for low-overhead, self-contained binaries with millisecond startup. This is an early release (CPU only): I'd especially like feedback on the API surface. Site: https://qxotic.ai
Repo: https://github.com/qxoticai/qxotic

Quixotic AI
Stack
Examples
Benchmarks
AI sovereignty
for the JVM
Some call AI on the JVM quixotic.
We call it Quixotic AI.
Finally, AI in a jar.
Every component of the AI stack, a la carte.
AI inference engine. Chat, vision, embeddings and text-to-speech capabilities for the JVM, with Spring AI and LangChain4j integrations.
Fast quantized matrix-multiplication routines.
TikToken-compatible and customizable tokenizers for popular LLM models.
Multi-backend tensor engine, in the works. Java, C, CUDA, HIP, Metal, OpenCL, and Mojo. Write once, accelerate everywhere.
Pure Java, read and write support for llama.cpp's GGUF model format.
Pure Java, read and write support for HuggingFace's Safetensors format.
Runnable snippets using jbang .
jbang header + imports ///usr/bin/env jbang "$0" "$@" ; exit $?
//JAVA 25+
//RUNTIME_OPTIONS --add-modules jdk.incubator.vector --enable-native-access=ALL-UNNAMED
//DEPS com.qxotic:jinfer-bom:0.2.0@pom
//DEPS com.qxotic:jinfer-spring-ai com.qxotic:jinfer-models-all
//DEPS com.qxotic:jam-native com.qxotic:jam-vector
//DEPS org.slf4j:slf4j-nop:2.0.17
import com.qxotic.jinfer.spring.ai.*;
void main () {
try ( var model = JinferChatModel .builder()
.model( "LiquidAI/LFM2.5-350M-GGUF:Q8_0" )
.build()) {
System .out.println(model.call( "What is the capital of France?" ));
}
}
$ jbang Chat.java # "The capital of France is Paris."
jbang header + imports ///usr/bin/env jbang "$0" "$@" ; exit $?
//JAVA 25+
//RUNTIME_OPTIONS --add-modules jdk.incubator.vector --enable-native-access=ALL-UNNAMED
//DEPS com.qxotic:jinfer-bom:0.2.0@pom
//DEPS com.qxotic:jinfer-spring-ai
//DEPS com.qxotic:jinfer-kokoro
//DEPS com.qxotic:jam-native com.qxotic:jam-vector
//DEPS org.slf4j:slf4j-nop:2.0.17
import com.qxotic.jinfer.spring.ai. JinferSpeechModel ;
import java.nio.file.*;
void main () throws Exception {
try ( var tts = JinferSpeechModel .builder()
.model( "simonfxr/kokoro.cpp-GGUF:Q8_0" )
.companion( "voice" , "simonfxr/kokoro.cpp-GGUF/voices/kokoro-voice-af_heart.gguf" )
.build()) {
var line = "Turns out, matrix multiplications can talk. "
+ "The JVM just crunched a few billion numbers to say this." ;
Files .write( Path .of( "kokoro.wav" ), tts.call(line));
}
}
$ jbang TextToSpeech.java # kokoro.wav · seven seconds, in a warm voice · Kokoro-82M
jbang header + imports ///usr/bin/env jbang "$0" "$@" ; exit $?
//JAVA 25+
//RUNTIME_OPTIONS --add-modules jdk.incubator.vector --enable-native-access=ALL-UNNAMED
//DEPS com.qxotic:jinfer-bom:0.2.0@pom
//DEPS com.qxotic:jinfer-spring-ai com.qxotic:jinfer-models-all
//DEPS org.springframework.ai:spring-ai-client-chat:2.0.1
//DEPS com.qxotic:jam-native com.qxotic:jam-vector
//DEPS org.slf4j:slf4j-nop:2.0.17
import com.qxotic.jinfer.spring.ai.*;
import org.springframework.ai.chat.client. ChatClient ;
import org.springframework.core.io. UrlResource ;
import org.springframework.util. MimeType ;
void main () throws Exception {
try ( var gemma = JinferChatModel .builder()
.model( "unsloth/gemma-4-E2B-it-GGUF:Q4_K_M" )
.companion( "media" , "unsloth/gemma-4-E2B-it-GGUF/mmproj-BF16.gguf" )
.build()) {
// jfk.wav: JFK's own voice, inaugural address, January 20, 1961
var speech = new UrlResource ( "https://qxotic.ai/snippets/jfk.wav" );
var transcript = ChatClient .create(gemma).prompt()
.user(u -> u.text( "Transcribe this recording." )
.media( MimeType .valueOf( "audio/wav" ), speech))
.call().content();
System .out.println(transcript);
}
}
$ jbang Audio.java # "ask not what your country can do for you ..." · every word heard
jbang header + imports ///usr/bin/env jbang "$0" "$@" ; exit $?
//JAVA 25+
//RUNTIME_OPTIONS --add-modules jdk.incubator.vector --enable-native-access=ALL-UNNAMED
//DEPS com.qxotic:jinfer-bom:0.2.0@pom
//DEPS com.qxotic:jinfer-spring-ai com.qxotic:jinfer-models-all
//DEPS org.springframework.ai:spring-ai-client-chat:2.0.1
//DEPS com.qxotic:jam-native com.qxotic:jam-vector
//DEPS org.slf4j:slf4j-nop:2.0.17
import com.qxotic.jinfer.spring.ai.*;
import org.springframework.ai.chat.client. ChatClient ;
import org.springframework.core.io. UrlResource ;
import org.springframework.util. MimeTypeUtils ;
void main () throws Exception {
try ( var vision = JinferChatModel .builder()
.model( "LiquidAI/LFM2.5-VL-3B-GGUF:Q8_0" )
.companion( "media" , "LiquidAI/LFM2.5-VL-3B-GGUF/mmproj-LFM2.5-VL-3B-Q8_0.gguf" )
.build()) {
// The windmills of Consuegra, La Mancha: Don Quixote's "giants"
var windmills = new UrlResource ( "https://qxotic.ai/snippets/consuegra.jpg" );
var answer = ChatClient .create(vision).prompt()
.user(u -> u.text( "Don Quixote saw giants here. What do you see?" )
.media( MimeTypeUtils .IMAGE_JPEG, windmills))
.call().content();
System .out.println(answer);
}
}
$ jbang Vision.java # "several traditional windmills standing on a hill under a clear blue sky"
jbang header + imports ///usr/bin/env jbang "$0" "$@" ; exit $?
//JAVA 25+
//RUNTIME_OPTIONS --add-modules jdk.incubator.vector --enable-native-access=ALL-UNNAMED
//DEPS com.qxotic:jinfer-bom:0.2.0@pom
//DEPS com.qxotic:jinfer-spring-ai com.qxotic:jinfer-models-all
//DEPS com.qxotic:jam-native com.qxotic:jam-vector
//DEPS org.slf4j:slf4j-nop:2.0.17
import com.qxotic.jinfer.spring.ai. JinferEmbeddingModel ;
void main () {
try ( var emb = JinferEmbeddingModel .builder()
.model( "LiquidAI/LFM2.5-Embedding-350M-GGUF:Q8_0" )
.build()) {
var a = emb.embed( "AI on the JVM" );
var b = emb.embed( "the JVM thinks now" );
var c = emb.embed( "the python shed its skin" );
System .out.printf( "similar %.3f%n" , cosineSimilarity(a, b));
System .out.printf( "unrelated %.3f%n" , cosineSimilarity(a, c));
}
}
cosineSimilarity helper float cosineSimilarity ( float [] a, float [] b) {
float dot = 0 , na = 0 , nb = 0 ;
for ( int i = 0 ; i < a.length; i++) {
dot += a[i] * b[i]; na += a[i] * a[i]; nb += b[i] * b[i];
}
return ( float ) (dot / Math .sqrt(na * nb));
}
$ jbang Embed.java # similar 0.609 · unrelated 0.081
Benchmarking is hard ... take these numbers as indicative, not definitive, always measure yourself.
jinfer-bench vs. llama-bench on CPU, ↑ higher is better.
Built for the JVM from first principles
Local AI, end-to-end on the JVM: no sidecars, no services, no IPC.
No Python runtime, no ONNX, no glue code. Every component, from tokenizers to the inference engine, built for the JVM, not bolted onto it.
Ships as a single self-contained binary: millisecond startup, small footprint, no JVM at runtime.
Write Once, Accelerate Everywhere
One Tensor API, seven backends: Java, C, CUDA, HIP, Metal, OpenCL, and Mojo.
Built in 🇨🇭 Switzerland · 2026 Quixotic AI
