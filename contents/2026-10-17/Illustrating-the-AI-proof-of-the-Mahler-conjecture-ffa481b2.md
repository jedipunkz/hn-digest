---
source: "https://chessapig.github.io/Mahler"
hn_url: "https://news.ycombinator.com/item?id=50131292"
title: "Illustrating the AI proof of the Mahler conjecture"
article_title: "Illustrating the AI proof of the Mahler conjecture | Elliot Kienzle"
image: "https://chessapig.github.io/sketch/mahler%20conjecture/images/thumbnail.jpg"
author: "lukasschwab"
captured_at: "2026-10-17T15:04:55Z"
capture_tool: "hn-digest"
hn_id: 50131292
score: 1
comments: 0
posted_at: "2026-10-17T14:04:40Z"
tags:
  - hacker-news
---

# Illustrating the AI proof of the Mahler conjecture

- HN: [50131292](https://news.ycombinator.com/item?id=50131292)
- Source: [chessapig.github.io](https://chessapig.github.io/Mahler)
- Score: 1
- Comments: 0
- Posted: 2026-10-17T14:04:40Z

## Translation

Title: Illustrating the AI proof of the Mahler conjecture
Article title: Illustrating the AI proof of the Mahler conjecture | Elliot Kienzle
Description: [!epigraph] Srinivasa Ramanujan (apocryphal (I made it up)) The machines may prove my theorems, but they will never prove my spirit!

Article text:
Illustrating the AI proof of the Mahler conjecture
Summary: This week, OpenAI posted a proof of the Mahler conjecture , among about 400 other problems. One of the proofs goes through symplectic geometry, constructing a symplectic embedding of a ball of the right volume into $K \times K^\circ$. In this post I illustrate the techniques of this proof, and draw the image of this symplectic ball.
[!epigraph] Srinivasa Ramanujan (apocryphal (I made it up))
The machines may prove my theorems,
but they will never prove my spirit!
Slop day has arrived, and we must reckon with the aftermath 1 . Of the 373 solutions released by OpenAI this Tuesday, I’m most interested in #087, a proof of the Mahler conjecture . Like many other mathematicians, I’ve been scooped :(. The AI proof is a clever construction applying a technique I haven’t published yet. My result wasn’t groundbreaking. It was a cute, self-contained thing with a simple proof (imo the best kind). I thought it could be quite useful (apparently so!) but I haven’t finished my application. The Mahler conjecture was on my list of things to try eventually. Ah well. Nothing to do but plant a flag in one of the 400 plots of the wasteland.
I’m most interested in the symplectic aspect of the Mahler conjecture. For an $n$-dimensional symmetric convex body $K$ and its polar $K^\circ$, AI constructs a symplectic embedding of a ball of radius $2/\sqrt{\pi}$ into $K\times K^\circ$. As a corollary, the volume of the product is greater than the volume of the ball, $\Vol(K \times K^\circ) \geq \frac{4^n}{n!}$. This settles the longstanding Mahler conjecture (along with two other AI papers, characterizing equality and proving the non-symmetric version). This problem is one of the few with a summarized chain of thought , which is honestly kinda insulting. In the 45-page document, the discussion of the symmetric case occupies a grand total of one paragraph.
In this post, I’ll first explain the words in the paragraph above, and why we care. This will be accessible to a broad mathematical audience. In the second part, I’ll motivate, illustrate, and sketch the proof from the AI paper, written for symplectic geometers.
I have two motivations for writing this post. Mathematically, I want to be a human for these ideas. I’ve been thinking about half of this (producing symplectic balls in holomorphic sub-manifolds) for the last three months. This AI result is a really cool application. If anyone wants to discuss other applications, I’m happy to chat :)
Second, I’m modeling the role of mathematical illustration in a world of AI, where proofs are cheaper than understanding. Rather than “digesting” the proof, I’m “illustrating” it, in the sense discussed in “On the importance of illustration for mathematical research” . AI coding has democratized the technical skills needed for interactive visualizations — I wrote all the programs in this article in a couple of days. Unaided, the machines produce visualizations that are correct and flashy, evident from all the vibe-explanations of math you see on Twitter. We must strive for a higher standard, for rigorous illustrations. When the process behind the illustration is born from the mathematics it illustrates, the creation and illustration of math merge into one. In this post, my main visualizations draws a four dimensional shapes via slicing. This slicing is both illustrative and mathematical, being the exact slicing used in the proof. The illustration is rigorous.
I spent a trimester in Paris with the math illustration community, trying to make notions like rigorous illustration precise (see what I made here ). I believe that illustration has an important role to play in our new age of math, as argued in the article Mathematical Illustration: Proof, Intuition and AI . You too can make illustrations. All it takes is care.
All text and hand-drawn figures in this document were made by me. I vibecoded all the applets with Claude Code, and had AI assist with website layout. I used Claude to help read the AI paper, and to mathematically and grammatically proofread this document.
Background
The Mahler conjecture
Symplectic capacities and the Viterbo conjecture
Proof
Symplectic monotonicity theorem
Proof of
Visualizing the symplectic balls
Constructing the holomorphic submanifold
Limiting with $L^\infty$ norms
A first guess at a holomorphic submanifold
Building $K\times K^\circ$ when $n=1$
The holomorphic submanifold construction
Visualizing the symplectic balls in $K\times K^\circ$
Every math field comes with its very own duality. The originator is linear algebra, where a vector space $V$ is dual to the space of one-forms $V^\ast$. This manifests in Euclidean geometry as a duality between points and half-spaces. Each one-form $\alpha$ defines a half-space containing the origin, $H_\alpha = \lbrace v \in V \, \big \vert\, \alpha(v) \leq 1\rbrace$. In polyhedral geometry, a convex polyhedron can be described as either the convex hull of its vertices, or as the intersection of half-spaces corresponding to its facets. The vertex and facet descriptions are dual.
The fundamental duality in convex geometry is polar duality . For a convex body $K \subset V$ define the polar $K^\circ \subset V^\ast$ as the collection of half-spaces containing $K$. That is,
The convex body $K$ is thought of as a set of points, while its polar $K^\circ$ is a collection of hyperplanes. Identifying $V$ with $V^\ast$ using the Euclidean metric, we can compare pictures. Play around with these in 2D to get a feel. Try making $K$ long and skinny, and see what happens to $K^\circ$.
[!sketch] mahler conjecture/polar polygon
A symmetric polygon $K$ and its polar $K^\circ$. The defining half-spaces of $K$ correspond to the vertices of $K^\circ$.
The polar operation behaves nicely.
$K^\circ \subset V^\ast$ is convex. If $K$ is closed, symmetric and convex, then $K^{\circ\circ} = K$. The polar is an honest duality relation.
$(\lambda K)^\circ = \frac{1}{\lambda}K^\circ$. Scaling up the body scales down its polar.
Exploiting the scaling property, we can build a scale-invariant quantity from $K$ and its polar. Let $\vert K\vert$ denote the volume of $K$. Define the Mahler volume as the product of volumes of $K$ and its polar:
The largest possible Mahler volume is achieved by a ball, $M(K) \leq M(B(1))$. The minimal Mahler volume was, until now, unknown.
[!conjecture] Mahler conjecture
Suppose $K$ is symmetric, meaning $K = -K$. The minimal Mahler volume is achieved by a cube.
We can think of Mahler volume as a measure of roundness. The (unique) highest roundness shape is a ball, the (non-unique) lowest roundness symmetric shape is a cube. For more details, history, and intuition, see Terry Tao’s blog . This problem has been kicking around for at least 100 years, and has inspired some beautiful bridges between convex geometry and several complex variables. It also ties closely with symplectic geometry.
Symplectic capacities and the Viterbo conjecture
Let us combine $K$ and $K^\circ$ into a single object, their Cartesian product $K \times K^\circ \subset V \times V^\ast$. We will be making a lot of pictures of this space, and now is a good time to set the visual language. We will always draw two dimensional bodies $K$, so $K\times K^\circ$ is a four dimensional shape. The two main ways to render a 4 dimensional object is through projection or slicing. Here’s the projection onto the screen, so you can look around.
[!sketch] mahler conjecture/polar product
On the left, The product $K\times K^\circ$, projected from $V\times V^\ast$ to the screen. In the middle is the view selector. place the red selectors in the middle to long along the red dimension, and vice versa for the green selectors. On the right, a selector for the polygon $K$ and its dual $K^\circ$.
A lot going on, right? Slicing will be more understandable for us. For any shape $\Omega \subset V\times V^\ast$, denote the slice above a point $x\in V$ as $\Omega^x = \Omega \cap \lbrace x \rbrace \times V^\ast$. By stacking all the slices $\Omega^x$ next to each other, we reconstruct $\Omega$. When $V$ is two dimensional, we can draw $\Omega$! Let the screen represent $V$. For a representative array of points on $V$, draw a scaled down copy of $\Omega^x$. Lining these up next to one another, our eyes notice the changes in the slices for each $x$. Applying this technique to $\Omega = K \times K^\circ$ is not so exciting, since every slice is the polar $\Omega^x = K^\circ$. Later on, we will see the slices change with position.
[!sketch|square] mahler conjecture/hexagonal grid/polar.html
Visualization of $K\times K^\circ$. You are invited to think of $K$ as an island, and the small copies of $K^\circ$ as flowers populating the island. This visual metaphor is inspired by geography and botany .
Geometrically, the ambient space $V \times V^\ast$ is a cotangent bundle $T^\ast V$, which carries a natural symplectic structure (i.e nondegenerate 2 form). In coordinates, if $V$ has coordinates $q_i$ and $V^\ast$ has dual coordinates $p_i$, then $\omega_{std} = \sum \de q_i \wedge \de p_i$ is the symplectic form. This makes $K\times K^\circ$ into a symplectic domain, whose symplectic volume equals the Mahler volume. Motivated by the Mahler conjecture, is there a uniform lower bound for the symplectic size of $K\times K^\circ$?
We formalize the notion of symplectic size with a symplectic capacity. Consider a machine $c$ which takes a symplectic domain $(D,\omega)$ and outputs a real number $c(D,\omega)$ (we will often use $c(D)$ as shorthand when the symplectic structure is implicit). In order for $c$ to represent symplectic size, it must satisfy a few axioms. First, it must be symplectic.
Symplectic invariance: for any diffeomorphism $\psi:D \to D’$, we have $c(D’,\omega) = c(D,\psi^\ast\omega)$.
Second, it must be a “size”. For any good notion of size, if a domain $D$ fits inside of $D’$, then the size of $D$ is less than or equal to that of $D’$. In the symplectic world, to fit is to symplectically embed. A map $\psi:(D,\omega) \hookrightarrow (D’,\omega’)$ is a symplectic embedding if it is an embedding satisfying $\psi^\ast \omega’ = \omega$. This motivates our second axiom.
Monotonicity: If there exists a symplectic embedding $\psi:D \hookrightarrow D’$, then $c(D) \leq c(D’)$
Anything satisfying $1$ and $2$ is a symplectic measure of size. For our sanity, we demand that the size scales linearly with the symplectic form.
Scaling: $c(D,\lambda \omega) = \lambda c(D,\omega)$.
A machine $c$ satisfying axioms $1-3$ is known as a symplectic capacity .
The quintessential capacity is the Gromov width $c_{Gr}$, which measures the largest ball that symplectically embeds inside the domain. To fix notation, we define some standard symplectic domains. Pick coordinates $q_i$ on $V$, and dual coordinates $p_i$ on $V^\ast$. Define the ball $B^n(A)$ and a cylinder $Z^n(A)$, which are symplectic domains endowed with the standard symplectic form
These are parametrized by their two dimensional cross-sectional area $A$. Define the Gromov width :
For example, Gromov’s non-squeezing theorem states that $c_{Gr}(Z^n(A))=A$, meaning a ball cannot symplectically embed into a cylinder with smaller cross-sectional area. You can’t squeeze a ball into a smaller tube.
A bold conjecture from 2000 posits that many measures of symplectic sizes agree on symmetric convex bodies. Say that a symplectic capacity is normalized if $c(B^n(A)) = c(Z^n(A))$. For example, the Gromov width is normalized.
[!conjecture] Strong Viterbo conjecture
Let $c,c’$ be normalized symplectic capacities. Then for any symmetric convex body $\Omega \subset T^\ast V$, the capacities agree
[AAKO] 2 showed that the Mahler conjecture is a corollary of the Viterbo conjecture. Specifically, they proved that the Hofer–Zehnder capacity $c_{HZ}$, another normalized capacity, satisfies $c_{HZ}(K \times K^\circ) = 4$. Viterbo’s conjecture implies
That is, there is a symplectic embe

[truncated]

## Original Extract

[!epigraph] Srinivasa Ramanujan (apocryphal (I made it up)) The machines may prove my theorems, but they will never prove my spirit!

Illustrating the AI proof of the Mahler conjecture
Summary: This week, OpenAI posted a proof of the Mahler conjecture , among about 400 other problems. One of the proofs goes through symplectic geometry, constructing a symplectic embedding of a ball of the right volume into $K \times K^\circ$. In this post I illustrate the techniques of this proof, and draw the image of this symplectic ball.
[!epigraph] Srinivasa Ramanujan (apocryphal (I made it up))
The machines may prove my theorems,
but they will never prove my spirit!
Slop day has arrived, and we must reckon with the aftermath 1 . Of the 373 solutions released by OpenAI this Tuesday, I’m most interested in #087, a proof of the Mahler conjecture . Like many other mathematicians, I’ve been scooped :(. The AI proof is a clever construction applying a technique I haven’t published yet. My result wasn’t groundbreaking. It was a cute, self-contained thing with a simple proof (imo the best kind). I thought it could be quite useful (apparently so!) but I haven’t finished my application. The Mahler conjecture was on my list of things to try eventually. Ah well. Nothing to do but plant a flag in one of the 400 plots of the wasteland.
I’m most interested in the symplectic aspect of the Mahler conjecture. For an $n$-dimensional symmetric convex body $K$ and its polar $K^\circ$, AI constructs a symplectic embedding of a ball of radius $2/\sqrt{\pi}$ into $K\times K^\circ$. As a corollary, the volume of the product is greater than the volume of the ball, $\Vol(K \times K^\circ) \geq \frac{4^n}{n!}$. This settles the longstanding Mahler conjecture (along with two other AI papers, characterizing equality and proving the non-symmetric version). This problem is one of the few with a summarized chain of thought , which is honestly kinda insulting. In the 45-page document, the discussion of the symmetric case occupies a grand total of one paragraph.
In this post, I’ll first explain the words in the paragraph above, and why we care. This will be accessible to a broad mathematical audience. In the second part, I’ll motivate, illustrate, and sketch the proof from the AI paper, written for symplectic geometers.
I have two motivations for writing this post. Mathematically, I want to be a human for these ideas. I’ve been thinking about half of this (producing symplectic balls in holomorphic sub-manifolds) for the last three months. This AI result is a really cool application. If anyone wants to discuss other applications, I’m happy to chat :)
Second, I’m modeling the role of mathematical illustration in a world of AI, where proofs are cheaper than understanding. Rather than “digesting” the proof, I’m “illustrating” it, in the sense discussed in “On the importance of illustration for mathematical research” . AI coding has democratized the technical skills needed for interactive visualizations — I wrote all the programs in this article in a couple of days. Unaided, the machines produce visualizations that are correct and flashy, evident from all the vibe-explanations of math you see on Twitter. We must strive for a higher standard, for rigorous illustrations. When the process behind the illustration is born from the mathematics it illustrates, the creation and illustration of math merge into one. In this post, my main visualizations draws a four dimensional shapes via slicing. This slicing is both illustrative and mathematical, being the exact slicing used in the proof. The illustration is rigorous.
I spent a trimester in Paris with the math illustration community, trying to make notions like rigorous illustration precise (see what I made here ). I believe that illustration has an important role to play in our new age of math, as argued in the article Mathematical Illustration: Proof, Intuition and AI . You too can make illustrations. All it takes is care.
All text and hand-drawn figures in this document were made by me. I vibecoded all the applets with Claude Code, and had AI assist with website layout. I used Claude to help read the AI paper, and to mathematically and grammatically proofread this document.
Background
The Mahler conjecture
Symplectic capacities and the Viterbo conjecture
Proof
Symplectic monotonicity theorem
Proof of
Visualizing the symplectic balls
Constructing the holomorphic submanifold
Limiting with $L^\infty$ norms
A first guess at a holomorphic submanifold
Building $K\times K^\circ$ when $n=1$
The holomorphic submanifold construction
Visualizing the symplectic balls in $K\times K^\circ$
Every math field comes with its very own duality. The originator is linear algebra, where a vector space $V$ is dual to the space of one-forms $V^\ast$. This manifests in Euclidean geometry as a duality between points and half-spaces. Each one-form $\alpha$ defines a half-space containing the origin, $H_\alpha = \lbrace v \in V \, \big \vert\, \alpha(v) \leq 1\rbrace$. In polyhedral geometry, a convex polyhedron can be described as either the convex hull of its vertices, or as the intersection of half-spaces corresponding to its facets. The vertex and facet descriptions are dual.
The fundamental duality in convex geometry is polar duality . For a convex body $K \subset V$ define the polar $K^\circ \subset V^\ast$ as the collection of half-spaces containing $K$. That is,
The convex body $K$ is thought of as a set of points, while its polar $K^\circ$ is a collection of hyperplanes. Identifying $V$ with $V^\ast$ using the Euclidean metric, we can compare pictures. Play around with these in 2D to get a feel. Try making $K$ long and skinny, and see what happens to $K^\circ$.
[!sketch] mahler conjecture/polar polygon
A symmetric polygon $K$ and its polar $K^\circ$. The defining half-spaces of $K$ correspond to the vertices of $K^\circ$.
The polar operation behaves nicely.
$K^\circ \subset V^\ast$ is convex. If $K$ is closed, symmetric and convex, then $K^{\circ\circ} = K$. The polar is an honest duality relation.
$(\lambda K)^\circ = \frac{1}{\lambda}K^\circ$. Scaling up the body scales down its polar.
Exploiting the scaling property, we can build a scale-invariant quantity from $K$ and its polar. Let $\vert K\vert$ denote the volume of $K$. Define the Mahler volume as the product of volumes of $K$ and its polar:
The largest possible Mahler volume is achieved by a ball, $M(K) \leq M(B(1))$. The minimal Mahler volume was, until now, unknown.
[!conjecture] Mahler conjecture
Suppose $K$ is symmetric, meaning $K = -K$. The minimal Mahler volume is achieved by a cube.
We can think of Mahler volume as a measure of roundness. The (unique) highest roundness shape is a ball, the (non-unique) lowest roundness symmetric shape is a cube. For more details, history, and intuition, see Terry Tao’s blog . This problem has been kicking around for at least 100 years, and has inspired some beautiful bridges between convex geometry and several complex variables. It also ties closely with symplectic geometry.
Symplectic capacities and the Viterbo conjecture
Let us combine $K$ and $K^\circ$ into a single object, their Cartesian product $K \times K^\circ \subset V \times V^\ast$. We will be making a lot of pictures of this space, and now is a good time to set the visual language. We will always draw two dimensional bodies $K$, so $K\times K^\circ$ is a four dimensional shape. The two main ways to render a 4 dimensional object is through projection or slicing. Here’s the projection onto the screen, so you can look around.
[!sketch] mahler conjecture/polar product
On the left, The product $K\times K^\circ$, projected from $V\times V^\ast$ to the screen. In the middle is the view selector. place the red selectors in the middle to long along the red dimension, and vice versa for the green selectors. On the right, a selector for the polygon $K$ and its dual $K^\circ$.
A lot going on, right? Slicing will be more understandable for us. For any shape $\Omega \subset V\times V^\ast$, denote the slice above a point $x\in V$ as $\Omega^x = \Omega \cap \lbrace x \rbrace \times V^\ast$. By stacking all the slices $\Omega^x$ next to each other, we reconstruct $\Omega$. When $V$ is two dimensional, we can draw $\Omega$! Let the screen represent $V$. For a representative array of points on $V$, draw a scaled down copy of $\Omega^x$. Lining these up next to one another, our eyes notice the changes in the slices for each $x$. Applying this technique to $\Omega = K \times K^\circ$ is not so exciting, since every slice is the polar $\Omega^x = K^\circ$. Later on, we will see the slices change with position.
[!sketch|square] mahler conjecture/hexagonal grid/polar.html
Visualization of $K\times K^\circ$. You are invited to think of $K$ as an island, and the small copies of $K^\circ$ as flowers populating the island. This visual metaphor is inspired by geography and botany .
Geometrically, the ambient space $V \times V^\ast$ is a cotangent bundle $T^\ast V$, which carries a natural symplectic structure (i.e nondegenerate 2 form). In coordinates, if $V$ has coordinates $q_i$ and $V^\ast$ has dual coordinates $p_i$, then $\omega_{std} = \sum \de q_i \wedge \de p_i$ is the symplectic form. This makes $K\times K^\circ$ into a symplectic domain, whose symplectic volume equals the Mahler volume. Motivated by the Mahler conjecture, is there a uniform lower bound for the symplectic size of $K\times K^\circ$?
We formalize the notion of symplectic size with a symplectic capacity. Consider a machine $c$ which takes a symplectic domain $(D,\omega)$ and outputs a real number $c(D,\omega)$ (we will often use $c(D)$ as shorthand when the symplectic structure is implicit). In order for $c$ to represent symplectic size, it must satisfy a few axioms. First, it must be symplectic.
Symplectic invariance: for any diffeomorphism $\psi:D \to D’$, we have $c(D’,\omega) = c(D,\psi^\ast\omega)$.
Second, it must be a “size”. For any good notion of size, if a domain $D$ fits inside of $D’$, then the size of $D$ is less than or equal to that of $D’$. In the symplectic world, to fit is to symplectically embed. A map $\psi:(D,\omega) \hookrightarrow (D’,\omega’)$ is a symplectic embedding if it is an embedding satisfying $\psi^\ast \omega’ = \omega$. This motivates our second axiom.
Monotonicity: If there exists a symplectic embedding $\psi:D \hookrightarrow D’$, then $c(D) \leq c(D’)$
Anything satisfying $1$ and $2$ is a symplectic measure of size. For our sanity, we demand that the size scales linearly with the symplectic form.
Scaling: $c(D,\lambda \omega) = \lambda c(D,\omega)$.
A machine $c$ satisfying axioms $1-3$ is known as a symplectic capacity .
The quintessential capacity is the Gromov width $c_{Gr}$, which measures the largest ball that symplectically embeds inside the domain. To fix notation, we define some standard symplectic domains. Pick coordinates $q_i$ on $V$, and dual coordinates $p_i$ on $V^\ast$. Define the ball $B^n(A)$ and a cylinder $Z^n(A)$, which are symplectic domains endowed with the standard symplectic form
These are parametrized by their two dimensional cross-sectional area $A$. Define the Gromov width :
For example, Gromov’s non-squeezing theorem states that $c_{Gr}(Z^n(A))=A$, meaning a ball cannot symplectically embed into a cylinder with smaller cross-sectional area. You can’t squeeze a ball into a smaller tube.
A bold conjecture from 2000 posits that many measures of symplectic sizes agree on symmetric convex bodies. Say that a symplectic capacity is normalized if $c(B^n(A)) = c(Z^n(A))$. For example, the Gromov width is normalized.
[!conjecture] Strong Viterbo conjecture
Let $c,c’$ be normalized symplectic capacities. Then for any symmetric convex body $\Omega \subset T^\ast V$, the capacities agree
[AAKO] 2 showed that the Mahler conjecture is a corollary of the Viterbo conjecture. Specifically, they proved that the Hofer–Zehnder capacity $c_{HZ}$, another normalized capacity, satisfies $c_{HZ}(K \times K^\circ) = 4$. Viterbo’s conjecture implies
That is, there is a symplectic embe

[truncated]
