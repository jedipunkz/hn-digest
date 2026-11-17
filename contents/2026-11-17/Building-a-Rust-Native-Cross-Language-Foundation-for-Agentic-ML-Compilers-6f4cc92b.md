---
source: "https://tvm.apache.org/2026/11/16/rust-native-cross-language-ml-compiler-foundation"
hn_url: "https://news.ycombinator.com/item?id=50528233"
title: "Building a Rust-Native, Cross-Language Foundation for Agentic ML Compilers"
article_title: "Building a Rust-Native, Cross-Language Foundation for Agentic ML Compilers"
image: ""
author: "matt_d"
captured_at: "2026-11-17T06:51:32Z"
capture_tool: "hn-digest"
hn_id: 50528233
score: 1
comments: 0
posted_at: "2026-11-17T06:46:30Z"
tags:
  - hacker-news
---

# Building a Rust-Native, Cross-Language Foundation for Agentic ML Compilers

- HN: [50528233](https://news.ycombinator.com/item?id=50528233)
- Source: [tvm.apache.org](https://tvm.apache.org/2026/11/16/rust-native-cross-language-ml-compiler-foundation)
- Score: 1
- Comments: 0
- Posted: 2026-11-17T06:46:30Z

## Translation

Title: Building a Rust-Native, Cross-Language Foundation for Agentic ML Compilers

Article text:
Building a Rust-Native, Cross-Language Foundation for Agentic ML Compilers
Nov 16, 2026
•
Apache TVM Community
Recent work on agent-driven kernel generation and systems optimization has shown that coding agents can take on substantial systems development. With agents writing a growing share of systems code, validating its correctness has become an increasingly important problem. As we discussed in our recent post “Rethinking Abstractions in the Agentic Era” , ML compiler foundation will continue to play a central role in the agentic era, to inspect, verify, and transform the programs that agents generate.
One natural consequence is that agents should now be able to get access to the intermediate program representations (IR), such as loops and calls and write customized compiler routines. In this scenario, the compiler foundation no longer only serves the human developers and increasingly becomes agent facing. One natural question arises: which programming language should agents pick to write ML compiler routines in?
Many language options are available and each comes with its own tradeoffs:
Rust is a very natural choice, as its safety guarantees enable us to catch memory bugs early while maintaining high performance in programming. Its ownership and borrowing rules, static type system, and concurrency guarantees prevent many memory-safety errors and data races at compile time. Not surprisingly, we start to see a big movement of rewriting systems entirely onto Rust with the help of AI agents. The safety and efficiency seems to make Rust the language of choice.
Python remains a strong choice, and is still the language in machine learning land, thanks to its rich ecosystem, and dynamic nature to enable quick prototyping. The ability to quickly prototype helps to cut down agent iteration time greatly.
C++ continues to have a strong base and admittedly the best infrastructure support for low-level compiler tool chains. As a matter of fact, the vast majority of production compiler frameworks are still written in C++, thanks to established frameworks like LLVM infrastructure.
Traditionally, most compilers are built around a single main language. The IR and other core data structures are defined in that language, typically C++, and languages such as Python reach them through bindings or wrappers. As we can see in the above discussion, there is not a really clear winner in terms of languages. In our experience agents would favor Python for prototyping, and Rust/C++ for robust stable implementations.
The above observations make us ask one question: why not natively support multiple languages? Our main insight is that most of the compiler routines operate on trees of nodes that represent the program, which we call IR, if we can enable in-memory, zero-cost access to these IR node objects across languages , and build first-class ergonomic APIs around them , we will be able to enable a first class cross language foundation that works across Rust, C++ and Python.
To understand how it works, it is helpful to distinguish our approach from traditional ones. Most traditional ML compilers would build solutions in one primary language, while exposing bindings to others. Those bindings typically bring additional indirections and may not present an idiomatic API for all the features needed. Instead, we want to provide fine-grained native access to the IR objects . Concretely, a Rust module should work on the same objects as a C++ pass, whether IR nodes or other compiler data structures, and read their fields directly. This calls for an agreement that sits below any single language. We build our solution on top of TVM-FFI , an open ABI for ML systems. Through its C ABI, every language shares the same reference-counted objects, and a registry records each type’s fields and where they sit in memory . Any library can add types to this registry when the library loads. A Rust library can therefore read C++ objects in place and contribute new types of its own.
We call our effort TIRx Rust, a Rust-Native, Cross-Language Foundation for Agentic ML Compilers . TIRx Rust builds on top of TIRx, a minimal and stable ML compiler foundation, while providing cross-language native IR access, ergonomic Rust APIs, and dynamic extension.
We envision a fundamentally new way to explore ML compiler development with the help of agents: instead of completely rewriting everything into one language like Rust, we allow agents to pick the best language for different modules at different stages of development while being able to mix and match at fine grained levels. Our preliminary experiences show that this approach empowers agents to develop faster with python prototyping, building robust code with Rust, while connecting to stable compiler toolchains in C++ organically.
While the TVM-FFI ABI provides a basic way to access the data structures, we still need to provide idiomatic and safe Rust APIs to inspect, traverse and transform the program structures. Each of the IR objects in Rust is stored as a reference counted object that follows Rust’s ownership rules. Different IR nodes can be represented by different node types.
One typical need in compiler analysis is to take a node dispatch based on its type, whether it is a For node, a constant or expression. The match_any! macro matches the IR node type:
fn describe ( expr : & PrimExpr ) -> Result < String > {
let text = match_any! {
expr {
Add ( ref add ) => format! ( "({} + {})" , describe ( & add .a ) ? , describe ( & add .b ) ? ),
Mul ( ref mul ) => format! ( "({} * {})" , describe ( & mul .a ) ? , describe ( & mul .b ) ? ),
Var ( ref var ) => var .name .as_str () .to_owned (),
IntImm ( ref imm ) => imm .value .to_string (),
_ => "<other>" .to_owned (),
}
};
Ok ( text )
}
The match_any! behaves, much like a match on an enum, one difference is that in our case, we do not require to enumerate enum ahead of time, instead, different downstream extensions can build their own variant of IR objects in different dynamic libraries and match_any! will work out of the box for them. This dynamic extension ability is extremely important as agents organically grow the program representations without having to rebuild the full program.
Besides match_any! , we also provide #[dispatch] macro that enables us to build type based dispatch easily in Rust. Such a pattern is very commonly used in visitors in most compiler analysis.
#[dispatch(print, (AnyView<'_>, bool), String)]
impl Printer {
fn print_add ( & self , add : & Add , nested : bool ) -> String {
let text = format! ( "{} + {}" , self .show ( & add .a , true ), self .show ( & add .b , true ));
if nested { format! ( "({text})" ) } else { text }
}
fn print_var ( & self , var : & Var ) -> String {
var .name .as_str () .to_owned ()
}
fn print_imm ( & self , imm : & IntImm ) -> String {
imm .value .to_string ()
}
fn show ( & self , expr : & PrimExpr , nested : bool ) -> String {
let text = self .dispatch ( AnyView :: from ( expr ), nested );
text .unwrap_or_else (|| "<other>" .to_owned ())
}
}
match_any! and #[dispatch] provide us with the basic API to access and dispatch handling at each step. In practice, we also need to recursively traverse the IR nodes and perform inspection and rewriting. We provide structural traversal APIs through four structural_* :
structural_walk and structural_visit read the IR, and structural_map and structural_mutate rewrite it. Of these, structural_walk and structural_map traverse the child nodes automatically, and the other two give users full control over the traversal order. Here is an example to show the structural_walk :
#[derive(Default)]
struct Counter {
vars : usize ,
adds : usize ,
}
#[dispatch(walk, (AnyView<'_>, DefRegionKind), Result<WalkResult>)]
impl Counter {
fn walk_var ( & mut self , _var : & Var ) -> Result < WalkResult > {
self .vars += 1 ;
...
Ok ( WalkResult :: Advance )
}
fn walk_add ( & mut self , _add : & Add ) -> Result < WalkResult > {
self .adds += 1 ;
Ok ( WalkResult :: Advance )
}
}
fn count () -> Result < tvm_ffi :: String > {
...
tvm_ffi :: structural_walk ( & expr , & mut counter , WalkOrder :: PreOrder ) ? ;
}
Combined with #[dispatch] or typed function closures, these traversals can form the core of a pass.
Beyond the ergonomic APIs, two questions remain. How do Rust and C++ keep object layouts in sync, and how do the traversals reach nodes whose types they have never seen?
We will use the add_one example to show the APIs in several languages.
add_one uses a loop to compute B[i] = A[i] + 1 . In each language, the same analysis collects the tensors add_one writes and finds B :
writes = []
def on_store ( store ):
writes . append ( store . dest . name )
tvm_ffi . structural_walk ( add_one , ( ir . TensorStore , on_store ))
print ( writes )
C++
ffi :: Array < ffi :: String > TensorsWritten ( const tirx :: Function & func ) {
ffi :: Array < ffi :: String > writes ;
ffi :: StructuralWalk < ffi :: WalkOrder :: kPostOrder > (
func , [ & ]( const TensorStore & store ) -> ffi :: Expected < ffi :: WalkResult > {
writes . push_back ( store -> dest . as < VarObj > () -> name );
return ffi :: WalkResult :: Advance ();
});
return writes ;
}
Rust
fn written_tensors ( func : Function ) -> Result < Strings > {
let mut writes = Vec :: new ();
let on_store = | store : TensorStore | -> Result < WalkResult > {
let dest = Var :: try_from ( & store .dest ) ? ;
writes .push ( dest .name .clone ());
Ok ( WalkResult :: Advance )
};
tvm_ffi :: structural_walk ( & func , on_store , WalkOrder :: PostOrder ) ? ;
Ok ( writes .into_iter () .collect ())
}
In Rust, store.dest is a plain field read with no wrapper call, even though TensorStore and Var are defined in C++. This is the fine-grained access that bindings make expensive.
All three versions call structural_walk , one of the four traversal entry points introduced earlier. It walks the whole IR and calls the callback only on nodes of the selected type.
Verified Layout Equivalence: Native Access Without Manual Synchronization
Reading a C++ object’s field directly from Rust works only if the Rust mirror of the class has exactly the same layout as the C++ definition : the same size and alignment, and every field at the same offset. However, the C++ compiler determines this layout from details Rust never sees, such as inheritance, templates, and padding.
If we maintain this mirror by hand, it can silently fail: once a C++ field moves, Rust code keeps reading from the old offset, which is undefined behavior. To stay in sync, the mirror should ideally be derived from the actual memory layout of C++ objects and updated automatically when the C++ definition changes.
TIRx generates this mirror automatically . TVM-FFI records each reflected type’s layout, as the C++ compiler computed it, in the compiled library. When running tvm-ffi-stubgen --target rust , its stubgen tool reads this layout and generates matching type definitions for other languages, including Rust structs.
Every struct that stubgen generates matches the C++ layout exactly. To guarantee this, stubgen first checks that the type has a complete layout: apart from padding, its reflected fields fill every byte from the end of its parent to the end of the type. For example, prim.Add has a complete layout, since its parent ir.Expr ends at byte 40 and a and b fill [40, 56) . Stubgen will also automatically add assertions to check that the struct for prim.Add has the same size and alignment as in C++.
Rust thus works with shared IR objects natively. Both languages hold pointers to the same allocation, and each Rust field read is one memory load at the offset C++ uses. Rust can allocate these objects too: stubgen generates a constructor that sets each field directly. However, if a type needs the logic in its C++ constructor, we write its Rust constructor by hand.
A type with an incomplete layout gets no struct and stays opaque behind C accessors, but most node types the pass works with have complete layou

[truncated]

## Original Extract

Building a Rust-Native, Cross-Language Foundation for Agentic ML Compilers
Nov 16, 2026
•
Apache TVM Community
Recent work on agent-driven kernel generation and systems optimization has shown that coding agents can take on substantial systems development. With agents writing a growing share of systems code, validating its correctness has become an increasingly important problem. As we discussed in our recent post “Rethinking Abstractions in the Agentic Era” , ML compiler foundation will continue to play a central role in the agentic era, to inspect, verify, and transform the programs that agents generate.
One natural consequence is that agents should now be able to get access to the intermediate program representations (IR), such as loops and calls and write customized compiler routines. In this scenario, the compiler foundation no longer only serves the human developers and increasingly becomes agent facing. One natural question arises: which programming language should agents pick to write ML compiler routines in?
Many language options are available and each comes with its own tradeoffs:
Rust is a very natural choice, as its safety guarantees enable us to catch memory bugs early while maintaining high performance in programming. Its ownership and borrowing rules, static type system, and concurrency guarantees prevent many memory-safety errors and data races at compile time. Not surprisingly, we start to see a big movement of rewriting systems entirely onto Rust with the help of AI agents. The safety and efficiency seems to make Rust the language of choice.
Python remains a strong choice, and is still the language in machine learning land, thanks to its rich ecosystem, and dynamic nature to enable quick prototyping. The ability to quickly prototype helps to cut down agent iteration time greatly.
C++ continues to have a strong base and admittedly the best infrastructure support for low-level compiler tool chains. As a matter of fact, the vast majority of production compiler frameworks are still written in C++, thanks to established frameworks like LLVM infrastructure.
Traditionally, most compilers are built around a single main language. The IR and other core data structures are defined in that language, typically C++, and languages such as Python reach them through bindings or wrappers. As we can see in the above discussion, there is not a really clear winner in terms of languages. In our experience agents would favor Python for prototyping, and Rust/C++ for robust stable implementations.
The above observations make us ask one question: why not natively support multiple languages? Our main insight is that most of the compiler routines operate on trees of nodes that represent the program, which we call IR, if we can enable in-memory, zero-cost access to these IR node objects across languages , and build first-class ergonomic APIs around them , we will be able to enable a first class cross language foundation that works across Rust, C++ and Python.
To understand how it works, it is helpful to distinguish our approach from traditional ones. Most traditional ML compilers would build solutions in one primary language, while exposing bindings to others. Those bindings typically bring additional indirections and may not present an idiomatic API for all the features needed. Instead, we want to provide fine-grained native access to the IR objects . Concretely, a Rust module should work on the same objects as a C++ pass, whether IR nodes or other compiler data structures, and read their fields directly. This calls for an agreement that sits below any single language. We build our solution on top of TVM-FFI , an open ABI for ML systems. Through its C ABI, every language shares the same reference-counted objects, and a registry records each type’s fields and where they sit in memory . Any library can add types to this registry when the library loads. A Rust library can therefore read C++ objects in place and contribute new types of its own.
We call our effort TIRx Rust, a Rust-Native, Cross-Language Foundation for Agentic ML Compilers . TIRx Rust builds on top of TIRx, a minimal and stable ML compiler foundation, while providing cross-language native IR access, ergonomic Rust APIs, and dynamic extension.
We envision a fundamentally new way to explore ML compiler development with the help of agents: instead of completely rewriting everything into one language like Rust, we allow agents to pick the best language for different modules at different stages of development while being able to mix and match at fine grained levels. Our preliminary experiences show that this approach empowers agents to develop faster with python prototyping, building robust code with Rust, while connecting to stable compiler toolchains in C++ organically.
While the TVM-FFI ABI provides a basic way to access the data structures, we still need to provide idiomatic and safe Rust APIs to inspect, traverse and transform the program structures. Each of the IR objects in Rust is stored as a reference counted object that follows Rust’s ownership rules. Different IR nodes can be represented by different node types.
One typical need in compiler analysis is to take a node dispatch based on its type, whether it is a For node, a constant or expression. The match_any! macro matches the IR node type:
fn describe ( expr : & PrimExpr ) -> Result < String > {
let text = match_any! {
expr {
Add ( ref add ) => format! ( "({} + {})" , describe ( & add .a ) ? , describe ( & add .b ) ? ),
Mul ( ref mul ) => format! ( "({} * {})" , describe ( & mul .a ) ? , describe ( & mul .b ) ? ),
Var ( ref var ) => var .name .as_str () .to_owned (),
IntImm ( ref imm ) => imm .value .to_string (),
_ => "<other>" .to_owned (),
}
};
Ok ( text )
}
The match_any! behaves, much like a match on an enum, one difference is that in our case, we do not require to enumerate enum ahead of time, instead, different downstream extensions can build their own variant of IR objects in different dynamic libraries and match_any! will work out of the box for them. This dynamic extension ability is extremely important as agents organically grow the program representations without having to rebuild the full program.
Besides match_any! , we also provide #[dispatch] macro that enables us to build type based dispatch easily in Rust. Such a pattern is very commonly used in visitors in most compiler analysis.
#[dispatch(print, (AnyView<'_>, bool), String)]
impl Printer {
fn print_add ( & self , add : & Add , nested : bool ) -> String {
let text = format! ( "{} + {}" , self .show ( & add .a , true ), self .show ( & add .b , true ));
if nested { format! ( "({text})" ) } else { text }
}
fn print_var ( & self , var : & Var ) -> String {
var .name .as_str () .to_owned ()
}
fn print_imm ( & self , imm : & IntImm ) -> String {
imm .value .to_string ()
}
fn show ( & self , expr : & PrimExpr , nested : bool ) -> String {
let text = self .dispatch ( AnyView :: from ( expr ), nested );
text .unwrap_or_else (|| "<other>" .to_owned ())
}
}
match_any! and #[dispatch] provide us with the basic API to access and dispatch handling at each step. In practice, we also need to recursively traverse the IR nodes and perform inspection and rewriting. We provide structural traversal APIs through four structural_* :
structural_walk and structural_visit read the IR, and structural_map and structural_mutate rewrite it. Of these, structural_walk and structural_map traverse the child nodes automatically, and the other two give users full control over the traversal order. Here is an example to show the structural_walk :
#[derive(Default)]
struct Counter {
vars : usize ,
adds : usize ,
}
#[dispatch(walk, (AnyView<'_>, DefRegionKind), Result<WalkResult>)]
impl Counter {
fn walk_var ( & mut self , _var : & Var ) -> Result < WalkResult > {
self .vars += 1 ;
...
Ok ( WalkResult :: Advance )
}
fn walk_add ( & mut self , _add : & Add ) -> Result < WalkResult > {
self .adds += 1 ;
Ok ( WalkResult :: Advance )
}
}
fn count () -> Result < tvm_ffi :: String > {
...
tvm_ffi :: structural_walk ( & expr , & mut counter , WalkOrder :: PreOrder ) ? ;
}
Combined with #[dispatch] or typed function closures, these traversals can form the core of a pass.
Beyond the ergonomic APIs, two questions remain. How do Rust and C++ keep object layouts in sync, and how do the traversals reach nodes whose types they have never seen?
We will use the add_one example to show the APIs in several languages.
add_one uses a loop to compute B[i] = A[i] + 1 . In each language, the same analysis collects the tensors add_one writes and finds B :
writes = []
def on_store ( store ):
writes . append ( store . dest . name )
tvm_ffi . structural_walk ( add_one , ( ir . TensorStore , on_store ))
print ( writes )
C++
ffi :: Array < ffi :: String > TensorsWritten ( const tirx :: Function & func ) {
ffi :: Array < ffi :: String > writes ;
ffi :: StructuralWalk < ffi :: WalkOrder :: kPostOrder > (
func , [ & ]( const TensorStore & store ) -> ffi :: Expected < ffi :: WalkResult > {
writes . push_back ( store -> dest . as < VarObj > () -> name );
return ffi :: WalkResult :: Advance ();
});
return writes ;
}
Rust
fn written_tensors ( func : Function ) -> Result < Strings > {
let mut writes = Vec :: new ();
let on_store = | store : TensorStore | -> Result < WalkResult > {
let dest = Var :: try_from ( & store .dest ) ? ;
writes .push ( dest .name .clone ());
Ok ( WalkResult :: Advance )
};
tvm_ffi :: structural_walk ( & func , on_store , WalkOrder :: PostOrder ) ? ;
Ok ( writes .into_iter () .collect ())
}
In Rust, store.dest is a plain field read with no wrapper call, even though TensorStore and Var are defined in C++. This is the fine-grained access that bindings make expensive.
All three versions call structural_walk , one of the four traversal entry points introduced earlier. It walks the whole IR and calls the callback only on nodes of the selected type.
Verified Layout Equivalence: Native Access Without Manual Synchronization
Reading a C++ object’s field directly from Rust works only if the Rust mirror of the class has exactly the same layout as the C++ definition : the same size and alignment, and every field at the same offset. However, the C++ compiler determines this layout from details Rust never sees, such as inheritance, templates, and padding.
If we maintain this mirror by hand, it can silently fail: once a C++ field moves, Rust code keeps reading from the old offset, which is undefined behavior. To stay in sync, the mirror should ideally be derived from the actual memory layout of C++ objects and updated automatically when the C++ definition changes.
TIRx generates this mirror automatically . TVM-FFI records each reflected type’s layout, as the C++ compiler computed it, in the compiled library. When running tvm-ffi-stubgen --target rust , its stubgen tool reads this layout and generates matching type definitions for other languages, including Rust structs.
Every struct that stubgen generates matches the C++ layout exactly. To guarantee this, stubgen first checks that the type has a complete layout: apart from padding, its reflected fields fill every byte from the end of its parent to the end of the type. For example, prim.Add has a complete layout, since its parent ir.Expr ends at byte 40 and a and b fill [40, 56) . Stubgen will also automatically add assertions to check that the struct for prim.Add has the same size and alignment as in C++.
Rust thus works with shared IR objects natively. Both languages hold pointers to the same allocation, and each Rust field read is one memory load at the offset C++ uses. Rust can allocate these objects too: stubgen generates a constructor that sets each field directly. However, if a type needs the logic in its C++ constructor, we write its Rust constructor by hand.
A type with an incomplete layout gets no struct and stays opaque behind C accessors, but most node types the pass works with have complete layou

[truncated]
