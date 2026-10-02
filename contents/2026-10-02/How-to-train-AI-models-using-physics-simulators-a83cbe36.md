---
source: "https://www.arenaphysica.com/publications/fem-edge-elements"
hn_url: "https://news.ycombinator.com/item?id=49933167"
title: "How to train AI models using physics simulators"
article_title: "Edge Elements: How FEM Solvers Represent Electromagnetic Fields | Arena Physica"
image: "https://www.arenaphysica.com/publications/fem-edge-elements/opengraph-image?c61037e45b1f3cab"
author: "pranade"
captured_at: "2026-10-02T13:38:04Z"
capture_tool: "hn-digest"
hn_id: 49933167
score: 1
comments: 0
posted_at: "2026-10-02T13:10:57Z"
tags:
  - hacker-news
---

# How to train AI models using physics simulators

- HN: [49933167](https://news.ycombinator.com/item?id=49933167)
- Source: [www.arenaphysica.com](https://www.arenaphysica.com/publications/fem-edge-elements)
- Score: 1
- Comments: 0
- Posted: 2026-10-02T13:10:57Z

## Translation

Title: How to train AI models using physics simulators
Article title: Edge Elements: How FEM Solvers Represent Electromagnetic Fields | Arena Physica
Description: A brief introduction to Nédélec edge elements and why standard electromagnetic simulation needs edge-based degrees of freedom.

Article text:
Edge Elements: How FEM Solvers Represent Electromagnetic Fields | Arena Physica Thesis Product Services Research Publications About Company Careers News System theme Get started blog Edge Elements: How FEM Solvers Represent Electromagnetic Fields
A brief introduction to Nédélec edge elements and why standard electromagnetic simulation needs edge-based degrees of freedom.
At Arena Physica, we're building a foundation model for electromagnetism (EM). To do that, we first need to generate a huge amount of training data from traditional simulation tools like FEM solvers . These solvers chop up the world into a mesh of small pieces and directly solve the differential equations governing electromagnetism within those pieces. However, when we first started collecting training data, we ran into a simple problem: once the FEM solver finished its solve, we didn't know how to extract all the available field values from the solution.
Some of us with machine learning research backgrounds had used FEM solvers before for thermal or structural mechanics problems, so it seemed reasonable to assume that we could collect everything the solver computed at mesh "nodes" during its solve (the vertices connecting all the mesh regions together), and interpolate between those node values later to get the field at any arbitrary point in space.
But when we tried to collect E-field data from the solver, there wasn't a "collect all" option. We had to provide the specific collection of probe positions where we wanted the field values. This is confusing. The solver works on a mesh. It must be computing the field at some collection of points in the mesh, right? So why can't we just export everything it computed? Why do we have to specify positions?
The answer surprised us: the solver never actually solves for the E-field at specific points on the mesh. It computes something else entirely.
Stop thinking in terms of nodes
If you're familiar with finite element methods from other domains (heat transfer, structural mechanics, etc.), you might expect this workflow:
Store field values ( E x E_x E x ​ , E y E_y E y ​ , E z E_z E z ​ ) at each mesh node (or at the center of each mesh cell)
Interpolate between nodes using shape functions
Extract the interpolated field values at an arbitrary location
This seems natural, and it works great for scalar fields like temperature. But for electromagnetic vector fields, it turns out that treating nodes as our degrees of freedom (DOF) is a fundamentally problematic approach. To understand why, we need to consult the physics we're solving.
At an interface between two materials, Maxwell's equations require:
Tangential continuity follows from Faraday's law . Imagine a thin rectangular loop straddling the interface. Stokes' theorem says:
where E \mathbf{E} E is the electric field around the loop and B \mathbf{B} B is the magnetic flux density through the loop. As the loop height shrinks to zero, the area vanishes, so the right-hand side of the equation goes to zero. The only surviving contributions to the line integral are the tangential components parallel to each side of the interface. Since they must cancel, E 1 ∥ = E 2 ∥ E_{1\parallel} = E_{2\parallel} E 1 ∥ ​ = E 2 ∥ ​ .
Normal discontinuity follows from Gauss's law . Apply the divergence theorem to a thin box straddling the interface:
​ ε E ⋅ d A = ∭ ρ d V ,
where ε \varepsilon ε is the material's permittivity and ρ \rho ρ is its free charge density . If we assume a scenario where both materials are dielectrics carrying no free charge, then ρ = 0 \rho = 0 ρ = 0 , and the right-hand side of the equation is also zero. Similarly to before, as the box flattens, all the side faces disappear and only the top and bottom faces contribute to the surface integral on the left-hand side. This means that the normal components of ε E \varepsilon \mathbf{E} ε E must cancel each other out on either side of the interface: ε 1 E 1 ⊥ = ε 2 E 2 ⊥ \varepsilon_1 E_{1\perp} = \varepsilon_2 E_{2\perp} ε 1 ​ E 1 ⊥ ​ = ε 2 ​ E 2 ⊥ ​ . So the component of the E-field perpendicular to the interface is not necessarily continuous across the interface. In this case, it jumps by the ratio of permittivities, but more generally, other material boundary conditions can lead to different kinds of discontinuities .
If a solver stored a single field vector at each shared node and interpolated between those values, it would be enforcing full continuity : all components would have to match across element boundaries. That would prevent the normal component from jumping at a material interface, even when Maxwell's equations allow it. A good solver needs a representation that keeps the tangential component continuous while allowing the normal component to jump.
Why doesn't nodal interpolation work for EM?
Edge elements: a different approach
Instead of storing field values at points, electromagnetic FEM solvers store line integrals along edges :
Physically, the edge coefficient e i j e_{ij} e ij ​ is the "work done per unit charge moving along that edge" . This might seem like a strange choice, but it naturally matches the line integral in the formulation of Faraday's law we saw above.
Sharing an edge coefficient forces neighboring elements to agree on the field along their common edge (the tangential component). But since the line integral doesn't care about the field across the edge (the normal component), that component can jump when crossing the boundary. This gives us the continuity we need without forcing the whole field vector to match.
In this post, we won't go into how exactly the e i j e_{ij} e ij ​ values are computed, but if we assume that our FEM solver has already computed them for us, how do we now use them to reconstruct the fields? A first-order approximation of E ( r ) \mathbf{E}(\mathbf{r}) E ( r ) represents the field at position r \mathbf{r} r as a linear combination of "basis functions" N i j ( r ) \mathbf{N}_{ij}(\mathbf{r}) N ij ​ ( r ) corresponding to each edge.
We want these basis functions to have the following properties:
They are defined over the entire mesh cell touching the edge
Their line integral along their corresponding edge equals 1, and along all other edges equals 0 (so that we recover the edge coefficient when we integrate the field over the edge)
It turns out there's a simple set of basis functions that has these properties, which comes from the lowest-order Nédélec elements of the first kind (also known as Whitney edge elements). They take the following form:
where L i L_i L i ​ are barycentric coordinates , weights that describe a point's position relative to the vertices. For the full 3D problem, our mesh consists of tetrahedra, but to understand how these basis functions work, we can simplify to a 2D problem with triangular elements (since the same principles apply).
Barycentric coordinates can be visualized as scalar fields that equal 1 at a vertex and 0 at all other vertices, varying linearly across the element:
This means that ∇ L i \nabla L_i ∇ L i ​ (the gradient of L i L_i L i ​ ) is a constant vector that points away from the side opposite vertex i i i .
Barycentric coordinates also have a geometric interpretation: L i L_i L i ​ is the fraction of the whole triangle occupied by the small triangle opposite the corresponding vertex:
Visualizing N i j \mathbf{N}_{ij} N ij ​ now, we see that each basis function is a vector field that curves in a circle around the vertex opposite the edge that the basis function is associated with:
ij ​ = L i ​ ∇ L j ​ − L j ​ ∇ L i ​ . The vectors flow along the selected edge, strongest (yellow) near that edge and weakest (purple) near the opposite vertex. Select an edge to view its basis function. Use the toggle button to switch between flow and vector views of the field.
Notice that the vector field is always perpendicular to the edges that it's not associated with. This is what allows the line integral to be zero along those edges:
That property is crucial: it makes the coefficients e i j e_{ij} e ij ​ truly independent degrees of freedom.
How is the line integral along the "correct" edge equal to 1?
Computing E \mathbf{E} E at an arbitrary point
So how do we actually get the field at an arbitrary point r \mathbf{r} r in 3D?
Find which tetrahedral element contains the point . Check barycentric coordinates: if all four L 1 , L 2 , L 3 , L 4 ≥ 0 L_1, L_2, L_3, L_4 \geq 0 L 1 ​ , L 2 ​ , L 3 ​ , L 4 ​ ≥ 0 , the point is inside that tetrahedron.
Look up the edge coefficients . For a tetrahedral element, these are the values e 12 , e 13 , e 14 , e 23 , e 24 , e 34 e_{12}, e_{13}, e_{14}, e_{23}, e_{24}, e_{34} e 12 ​ , e 13 ​ , e 14 ​ , e 23 ​ , e 24 ​ , e 34 ​ from the FEM solution vector.
Evaluate the basis functions . For each edge, compute: N i j ( r ) = L i ( r ) ∇ L j − L j ( r ) ∇ L i \mathbf{N}_{ij}(\mathbf{r}) = L_i(\mathbf{r}) \nabla L_j - L_j(\mathbf{r}) \nabla L_i N ij ​ ( r ) = L i ​ ( r ) ∇ L j ​ − L j ​ ( r ) ∇ L i ​
Sum them up : E ( r ) = ∑ edges e i j N i j ( r ) \mathbf{E}(\mathbf{r}) = \sum_{\text{edges}} e_{ij} \mathbf{N}_{ij}(\mathbf{r}) E ( r ) = ∑ edges ​ e ij ​ N ij ​ ( r )
In our 2D example, we can see how the field is reconstructed everywhere within the element as we adjust the edge coefficients:
When we put two elements next to each other, the shared edge has one coefficient used by both triangles, which automatically enforces tangential continuity at the boundary:
If we scale this up, we can see the effect of shared edge coefficients on a larger simulation domain:
Why is ∇ × E \nabla \times \mathbf{E} ∇ × E constant within each element?
What about higher-order elements?
Everything we've discussed uses the lowest-order Nédélec element, which stores just one coefficient per edge. These coefficients are effectively a "compressed representation" of the E-field within the mesh. However, using only these lowest-order elements limits how much the field can vary within each element.
"Higher-order" elements increase the fidelity of the representation by adding more degrees of freedom: additional coefficients per edge (to capture variation along the edge), plus new DOFs on faces and within volumes:
The details of how these higher-order elements work are beyond the scope of this post, but extracting the field follows the same idea: evaluate the basis functions at the position of interest and combine them using the coefficients the solver computed.
So, there isn't a hidden list of E-field samples waiting to be exported. The solver computes a representation of the field in terms of edge elements, and our requested probe positions tell it where to interpolate field values based on those elements. If there were a "collect all" button, it would give us edge coefficients, not field values.
Thank you to Boyuan Zhang, PhD, and Hao Liu for their feedback on the technical content of this post.
The choice of using edges for the E \mathbf{E} E field is part of a mathematical structure called the de Rham complex (see: Finite Element Exterior Calculus ). The de Rham complex comes from differential geometry, and it describes the relationship between grad, curl, and div operators. In FEM, this relationship guides where we place the degrees of freedom on the mesh: at vertices, along edges, across faces, or within cells. These choices determine which field components must stay continuous between neighboring mesh cells:
For attribution in academic contexts, please cite this work as:
Bryant, "Edge Elements: How FEM Solvers Represent Electromagnetic Fields", Arena Physica, 2026. https://www.arenaphysica.com/publications/fem-edge-elements BibTeX citation:
@misc{bryant2026femedgeelements,
author = {Bryant, Christopher M.},
title = {Edge Elements: How FEM Solvers Represent Electromagnetic Fields},
howpublished = {Arena Physica},
year = {2026},
month = oct,
url = {https://www.arenaphysica.com/publications/fem-edge-elements}
} Contents Stop thinking in terms of nodes
Edge elements: a differen

[truncated]

## Original Extract

A brief introduction to Nédélec edge elements and why standard electromagnetic simulation needs edge-based degrees of freedom.

Edge Elements: How FEM Solvers Represent Electromagnetic Fields | Arena Physica Thesis Product Services Research Publications About Company Careers News System theme Get started blog Edge Elements: How FEM Solvers Represent Electromagnetic Fields
A brief introduction to Nédélec edge elements and why standard electromagnetic simulation needs edge-based degrees of freedom.
At Arena Physica, we're building a foundation model for electromagnetism (EM). To do that, we first need to generate a huge amount of training data from traditional simulation tools like FEM solvers . These solvers chop up the world into a mesh of small pieces and directly solve the differential equations governing electromagnetism within those pieces. However, when we first started collecting training data, we ran into a simple problem: once the FEM solver finished its solve, we didn't know how to extract all the available field values from the solution.
Some of us with machine learning research backgrounds had used FEM solvers before for thermal or structural mechanics problems, so it seemed reasonable to assume that we could collect everything the solver computed at mesh "nodes" during its solve (the vertices connecting all the mesh regions together), and interpolate between those node values later to get the field at any arbitrary point in space.
But when we tried to collect E-field data from the solver, there wasn't a "collect all" option. We had to provide the specific collection of probe positions where we wanted the field values. This is confusing. The solver works on a mesh. It must be computing the field at some collection of points in the mesh, right? So why can't we just export everything it computed? Why do we have to specify positions?
The answer surprised us: the solver never actually solves for the E-field at specific points on the mesh. It computes something else entirely.
Stop thinking in terms of nodes
If you're familiar with finite element methods from other domains (heat transfer, structural mechanics, etc.), you might expect this workflow:
Store field values ( E x E_x E x ​ , E y E_y E y ​ , E z E_z E z ​ ) at each mesh node (or at the center of each mesh cell)
Interpolate between nodes using shape functions
Extract the interpolated field values at an arbitrary location
This seems natural, and it works great for scalar fields like temperature. But for electromagnetic vector fields, it turns out that treating nodes as our degrees of freedom (DOF) is a fundamentally problematic approach. To understand why, we need to consult the physics we're solving.
At an interface between two materials, Maxwell's equations require:
Tangential continuity follows from Faraday's law . Imagine a thin rectangular loop straddling the interface. Stokes' theorem says:
where E \mathbf{E} E is the electric field around the loop and B \mathbf{B} B is the magnetic flux density through the loop. As the loop height shrinks to zero, the area vanishes, so the right-hand side of the equation goes to zero. The only surviving contributions to the line integral are the tangential components parallel to each side of the interface. Since they must cancel, E 1 ∥ = E 2 ∥ E_{1\parallel} = E_{2\parallel} E 1 ∥ ​ = E 2 ∥ ​ .
Normal discontinuity follows from Gauss's law . Apply the divergence theorem to a thin box straddling the interface:
​ ε E ⋅ d A = ∭ ρ d V ,
where ε \varepsilon ε is the material's permittivity and ρ \rho ρ is its free charge density . If we assume a scenario where both materials are dielectrics carrying no free charge, then ρ = 0 \rho = 0 ρ = 0 , and the right-hand side of the equation is also zero. Similarly to before, as the box flattens, all the side faces disappear and only the top and bottom faces contribute to the surface integral on the left-hand side. This means that the normal components of ε E \varepsilon \mathbf{E} ε E must cancel each other out on either side of the interface: ε 1 E 1 ⊥ = ε 2 E 2 ⊥ \varepsilon_1 E_{1\perp} = \varepsilon_2 E_{2\perp} ε 1 ​ E 1 ⊥ ​ = ε 2 ​ E 2 ⊥ ​ . So the component of the E-field perpendicular to the interface is not necessarily continuous across the interface. In this case, it jumps by the ratio of permittivities, but more generally, other material boundary conditions can lead to different kinds of discontinuities .
If a solver stored a single field vector at each shared node and interpolated between those values, it would be enforcing full continuity : all components would have to match across element boundaries. That would prevent the normal component from jumping at a material interface, even when Maxwell's equations allow it. A good solver needs a representation that keeps the tangential component continuous while allowing the normal component to jump.
Why doesn't nodal interpolation work for EM?
Edge elements: a different approach
Instead of storing field values at points, electromagnetic FEM solvers store line integrals along edges :
Physically, the edge coefficient e i j e_{ij} e ij ​ is the "work done per unit charge moving along that edge" . This might seem like a strange choice, but it naturally matches the line integral in the formulation of Faraday's law we saw above.
Sharing an edge coefficient forces neighboring elements to agree on the field along their common edge (the tangential component). But since the line integral doesn't care about the field across the edge (the normal component), that component can jump when crossing the boundary. This gives us the continuity we need without forcing the whole field vector to match.
In this post, we won't go into how exactly the e i j e_{ij} e ij ​ values are computed, but if we assume that our FEM solver has already computed them for us, how do we now use them to reconstruct the fields? A first-order approximation of E ( r ) \mathbf{E}(\mathbf{r}) E ( r ) represents the field at position r \mathbf{r} r as a linear combination of "basis functions" N i j ( r ) \mathbf{N}_{ij}(\mathbf{r}) N ij ​ ( r ) corresponding to each edge.
We want these basis functions to have the following properties:
They are defined over the entire mesh cell touching the edge
Their line integral along their corresponding edge equals 1, and along all other edges equals 0 (so that we recover the edge coefficient when we integrate the field over the edge)
It turns out there's a simple set of basis functions that has these properties, which comes from the lowest-order Nédélec elements of the first kind (also known as Whitney edge elements). They take the following form:
where L i L_i L i ​ are barycentric coordinates , weights that describe a point's position relative to the vertices. For the full 3D problem, our mesh consists of tetrahedra, but to understand how these basis functions work, we can simplify to a 2D problem with triangular elements (since the same principles apply).
Barycentric coordinates can be visualized as scalar fields that equal 1 at a vertex and 0 at all other vertices, varying linearly across the element:
This means that ∇ L i \nabla L_i ∇ L i ​ (the gradient of L i L_i L i ​ ) is a constant vector that points away from the side opposite vertex i i i .
Barycentric coordinates also have a geometric interpretation: L i L_i L i ​ is the fraction of the whole triangle occupied by the small triangle opposite the corresponding vertex:
Visualizing N i j \mathbf{N}_{ij} N ij ​ now, we see that each basis function is a vector field that curves in a circle around the vertex opposite the edge that the basis function is associated with:
ij ​ = L i ​ ∇ L j ​ − L j ​ ∇ L i ​ . The vectors flow along the selected edge, strongest (yellow) near that edge and weakest (purple) near the opposite vertex. Select an edge to view its basis function. Use the toggle button to switch between flow and vector views of the field.
Notice that the vector field is always perpendicular to the edges that it's not associated with. This is what allows the line integral to be zero along those edges:
That property is crucial: it makes the coefficients e i j e_{ij} e ij ​ truly independent degrees of freedom.
How is the line integral along the "correct" edge equal to 1?
Computing E \mathbf{E} E at an arbitrary point
So how do we actually get the field at an arbitrary point r \mathbf{r} r in 3D?
Find which tetrahedral element contains the point . Check barycentric coordinates: if all four L 1 , L 2 , L 3 , L 4 ≥ 0 L_1, L_2, L_3, L_4 \geq 0 L 1 ​ , L 2 ​ , L 3 ​ , L 4 ​ ≥ 0 , the point is inside that tetrahedron.
Look up the edge coefficients . For a tetrahedral element, these are the values e 12 , e 13 , e 14 , e 23 , e 24 , e 34 e_{12}, e_{13}, e_{14}, e_{23}, e_{24}, e_{34} e 12 ​ , e 13 ​ , e 14 ​ , e 23 ​ , e 24 ​ , e 34 ​ from the FEM solution vector.
Evaluate the basis functions . For each edge, compute: N i j ( r ) = L i ( r ) ∇ L j − L j ( r ) ∇ L i \mathbf{N}_{ij}(\mathbf{r}) = L_i(\mathbf{r}) \nabla L_j - L_j(\mathbf{r}) \nabla L_i N ij ​ ( r ) = L i ​ ( r ) ∇ L j ​ − L j ​ ( r ) ∇ L i ​
Sum them up : E ( r ) = ∑ edges e i j N i j ( r ) \mathbf{E}(\mathbf{r}) = \sum_{\text{edges}} e_{ij} \mathbf{N}_{ij}(\mathbf{r}) E ( r ) = ∑ edges ​ e ij ​ N ij ​ ( r )
In our 2D example, we can see how the field is reconstructed everywhere within the element as we adjust the edge coefficients:
When we put two elements next to each other, the shared edge has one coefficient used by both triangles, which automatically enforces tangential continuity at the boundary:
If we scale this up, we can see the effect of shared edge coefficients on a larger simulation domain:
Why is ∇ × E \nabla \times \mathbf{E} ∇ × E constant within each element?
What about higher-order elements?
Everything we've discussed uses the lowest-order Nédélec element, which stores just one coefficient per edge. These coefficients are effectively a "compressed representation" of the E-field within the mesh. However, using only these lowest-order elements limits how much the field can vary within each element.
"Higher-order" elements increase the fidelity of the representation by adding more degrees of freedom: additional coefficients per edge (to capture variation along the edge), plus new DOFs on faces and within volumes:
The details of how these higher-order elements work are beyond the scope of this post, but extracting the field follows the same idea: evaluate the basis functions at the position of interest and combine them using the coefficients the solver computed.
So, there isn't a hidden list of E-field samples waiting to be exported. The solver computes a representation of the field in terms of edge elements, and our requested probe positions tell it where to interpolate field values based on those elements. If there were a "collect all" button, it would give us edge coefficients, not field values.
Thank you to Boyuan Zhang, PhD, and Hao Liu for their feedback on the technical content of this post.
The choice of using edges for the E \mathbf{E} E field is part of a mathematical structure called the de Rham complex (see: Finite Element Exterior Calculus ). The de Rham complex comes from differential geometry, and it describes the relationship between grad, curl, and div operators. In FEM, this relationship guides where we place the degrees of freedom on the mesh: at vertices, along edges, across faces, or within cells. These choices determine which field components must stay continuous between neighboring mesh cells:
For attribution in academic contexts, please cite this work as:
Bryant, "Edge Elements: How FEM Solvers Represent Electromagnetic Fields", Arena Physica, 2026. https://www.arenaphysica.com/publications/fem-edge-elements BibTeX citation:
@misc{bryant2026femedgeelements,
author = {Bryant, Christopher M.},
title = {Edge Elements: How FEM Solvers Represent Electromagnetic Fields},
howpublished = {Arena Physica},
year = {2026},
month = oct,
url = {https://www.arenaphysica.com/publications/fem-edge-elements}
} Contents Stop thinking in terms of nodes
Edge elements: a differen

[truncated]
