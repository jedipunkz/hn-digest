---
source: "https://github.com/bp-purana52127/openai-navier-stokes-critique/blob/main/README.md"
hn_url: "https://news.ycombinator.com/item?id=50170670"
title: "Structural and mathematical audit of OpenAI's Navier-Stokes formalization"
article_title: "openai-navier-stokes-critique/README.md at main · bp-purana52127/openai-navier-stokes-critique · GitHub"
image: "https://opengraph.githubassets.com/4b8f98933b5a6bac2f7b8706e6fe283882ad26dd9efb3ca6b6062fa70d1ca6e5/bp-purana52127/openai-navier-stokes-critique"
author: "BPPurana"
captured_at: "2026-10-20T18:57:44Z"
capture_tool: "hn-digest"
hn_id: 50170670
score: 1
comments: 0
posted_at: "2026-10-20T18:35:02Z"
tags:
  - hacker-news
---

# Structural and mathematical audit of OpenAI's Navier-Stokes formalization

- HN: [50170670](https://news.ycombinator.com/item?id=50170670)
- Source: [github.com](https://github.com/bp-purana52127/openai-navier-stokes-critique/blob/main/README.md)
- Score: 1
- Comments: 0
- Posted: 2026-10-20T18:35:02Z

## Translation

Title: Structural and mathematical audit of OpenAI's Navier-Stokes formalization
Article title: openai-navier-stokes-critique/README.md at main · bp-purana52127/openai-navier-stokes-critique · GitHub
Description: A structural audit demonstrating the infinite energy requirement hidden within OpenAI's forced Navier-Stokes Lean 4 formalization. - openai-navier-stokes-critique/README.md at main · bp-purana52127/openai-navier-stokes-critique

Article text:
openai-navier-stokes-critique/README.md at main · bp-purana52127/openai-navier-stokes-critique · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
bp-purana52127
Copy path Blame More file actions Blame More file actions Latest commit
History History 169 lines (123 loc) · 16.8 KB main Breadcrumbs
Copy path Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions Structural & Mathematical Audit: openai/NavierStokesAndEuler
This repository contains an architectural, mathematical, and code-level verification audit of the openai/NavierStokesAndEuler Lean 4 formalization and its accompanying LaTeX preprint.
By tracing the dependency graph down to its leaf entry points and inspecting the core definitions across both ./NavierStokes/ (Forced 3D Navier-Stokes) and ./Euler/ (Unforced 3D Euler), this audit documents the exact architectural mechanisms used in the formalization—specifically manuscript-to-codebase specification gaps , conditional structure hypotheses ( PhysicalFields , StageEstimates ) , annular zero-germ cutoffs , post-hoc diagonal cutoff schedules , and a $10^{25}$ parameter scale divergence between the LaTeX preprint's explicit constants and Lean's non-constructive atTop filter.
Executive Summary of Audit Findings
Audit Area
Target Directory & Primary Files
Core Technical & Verification Finding
1. Documented Manuscript vs. Codebase Gaps
./NavierStokes/
GraphRestriction.lean
CycleContinuationInvariant.lean
DiagonalJetBoundsNoOptions.lean
Multiple modules explicitly document in their headers that they do not supply the manuscript's unconditional all-order estimates. Notably, GraphRestriction.lean proves a counterexample ( zero_mean_does_not_force_zero ) showing that zero auxiliary mean on the torus does not force zero after restricting to the graph.
2. External Realization Hypotheses & Conditional Assembly
./NavierStokes/
ActualCycleResidualBounds.lean
GluedStageEstimates.lean
CandidateAssembly.lean
MixedCandidateAssembly.lean
In ActualCycleResidualBounds.lean and GluedStageEstimates.lean , the physical realization of the stage velocity and pressure fields is passed as an external structure hypothesis ( PhysicalFields / PhysicalData : "Only the actual physical realization remains external" ).
3. Annular Zero-Germ Cutoffs & Residual Forcing Architecture
./NavierStokes/
ComparatorR3Theorem.lean
PeriodicResidualLimits.lean
MixedAxisPreservation.lean
ActualCandidateAssembly.lean
Clay Statement (C) ( option_C_of_compact_candidate ) is formalized with $u_0 \equiv 0$ by defining the external force $f(t,x)$ on $[0, 1)$ as the PDE residual $\partial_t u + (u\cdot\nabla)u - \nu\Delta u + \nabla p$ . All infinite wave-packet perturbations have a zero germ on the symmetry axis $r = 0$ ( mixedDiagonal_eq_base_germ ), isolating the blowup at $(t=1, x=0)$ entirely to the Stage-0 self-similar profile FinalSlowBase .
4. Exponential Mode Growth & Post-Hoc Cutoff Schedules
./NavierStokes/
ActualCandidateConstruction.lean
GluedStageEstimates.lean
MixedCandidateAssembly.lean
At cycle $j$ , active harmonic modes scale up to $\Vert{}k\Vert{} \le 2^{j+1}$ ( cycle_residualBand ), causing stage derivative constants $C_{j,m}$ to grow exponentially with $j$ and $m$ . To establish VanishingJointJets at $(1
[truncated]
1. Self-Documented Gaps Between the Manuscript and the Lean Code
A direct audit of the module docstrings and theorem statements in ./NavierStokes/ reveals several places where the formalization explicitly diverges from or weakens the unconditional claims of the written manuscript:
GraphRestriction.lean (Lines 16–18):
"None of these statements supplies the manuscript's claimed all-order residual estimates."
Furthermore, GraphRestriction.lean formalizes an explicit counterexample ( zero_mean_does_not_force_zero ), demonstrating that a zero auxiliary mean on the torus does not force a zero value after restriction to the graph.
CycleContinuationInvariant.lean (Lines 12–17):
"This file does not infer that input from smooth incoming primitives alone: an inverse covariance factor may grow at a flat annulus edge, so such an inference would require additional weighted estimates. In particular, the declarations below are a reconstruction algebra, not an unconditional iteration theorem or an assumed physical endpoint extension."
DiagonalJetBoundsNoOptions.lean (Lines 9–12):
"Pointwise stage estimates then give quantitative tail estimates. The stage estimates themselves are explicit hypotheses, not conclusions of the numerical cutoff selection."
CandidateAssembly.lean (Lines 10–14, 214):
"The theorem proves all CandidateProperties for these explicit outputs. It does not construct the input fields or the future force... A conditional implication, not a proof that its analytic inputs exist."
2. External PhysicalFields Hypotheses in Stage Residual Bounds
To prove Clay Statement (C) via the method of manufactured solutions ( option_C_of_compact_candidate in ComparatorR3Theorem.lean , Lines 22–35), the formalization sets the initial velocity to identically zero ( $u_0 \equiv 0$ ) and defines the external body force $f(t,x)$ on $t \in [0, 1)$ as the Navier-Stokes residual of the candidate fields $(u, p)$ ( assembledForce in CandidateAssembly.lean , Lines 20–67; periodicResidual in PeriodicResidualLimits.lean , Lines 23–35):
$$f(t,x) := \partial_t u + (u \cdot \nabla)u - \nu \Delta u + \nabla p$$
Under Fefferman's Statement (C) formulation, defining $f$ as the PDE residual of a blowing-up field $u$ is valid if and only if one proves that all spacetime derivatives of the residual cancel to infinite order at the singularity $(t=1, x=0)$ ( VanishingJointJets ), allowing $f$ to extend smoothly ( $C^\infty$ ) across $t = 1$ via Borel/Whitney extension ( CandidateFromLimits.force in MixedPeriodicAssembly.lean , Line 268).
In MixedCandidateAssembly.lean (Lines 52–60), VanishingJointJets depends entirely on supplying a witness for StageEstimates.finite_residual :
finite_residual : ∀ J m,
JetRate (𝓝[SpacetimeEndpoint.openPast 1 ] ( 1 , ( 0 : Space))) (PhysicalWaveSum.physicalQ h)
(fun z => navierStokesResidual (MixedDiagonalResidual.uncutVelocity A B J)
(DiagonalJetBounds.uncutPrefix P (J + 1 )) z. 1 z. 2 ) m (gain J - residualLoss m)
Tracing finite_residual to its provider in GluedStageEstimates.lean (Line 298) and ActualCycleResidualBounds.lean (Line 473) reveals that actualStageEstimates is parameterized by an explicit external realization input d :
(d : ∀ J, ActualCycleResidualBounds.PhysicalData B (N + 1 )
(ActualCyclePreservation.state B N0 J).state
(MixedDiagonalResidual.uncutVelocity A Bdirect J)
(DiagonalJetBounds.uncutPrefix P (J + 1 )))
And in ActualCycleResidualBounds.lean (Line 473), PhysicalFields (which underlies PhysicalData ) is documented as:
/-- Only the actual physical realization remains external. Native regularity,
the fixed base equation, and all size estimates are derived in this module. -/
structure PhysicalFields (B N : ℕ) (U : Set Cylinder) (s : StatePoint)
(u : VelocityField) (P : PressureField) : Prop where
velocity_germ : ∀ n, N ≤ n → ∀ z ∈ preterminal,
z ∈ PhysicalResidualTZ.graphSourceTZ (actualBandGraph n) U → ...
Rather than unconditionally constructing the global physical primitives in Cartesian coordinates, the residual module assumes velocity_germ to identify the Cartesian fields with the cylindrical graph pullback ( polarGraph ) on each dyadic band.
3. Annular Zero-Germ Isolation & Post-Hoc Cutoff Absorption
Inspecting ActualCandidateAssembly.lean (Lines 331–381) and MixedAxisPreservation.lean (Lines 33, 115, 212) clarifies how the blowup at $(t=1, x=0)$ interacts with the infinite wave-packet series:
theorem axis_not_active {w : SpaceTime} (ht : w ∈ PhysicalWaveSum.preterminal)
(ha : PhysicalGraphBounds.radialProjection w = 0 ) :
w ∉ ActualPolarCoverage.active
theorem mixedDiagonal_eq_base_germ :
mixedDiagonal base p D scales =ᶠ[𝓝 w] SpatialCurl.spatialCurl base
Zero Germ on the Blowup Axis ( $r = 0$ ): Every wave-packet perturbation stage ( particularPotential , signedPotential , directStages ) is supported strictly inside an active radial annulus $r / q \in [\text{leftRadius}, \text{rightRadius}]$ with $\text{leftRadius} &gt; 0$ ( PrimaryTargetBounds.leftRadius_pos ). Consequently, in a neighborhood of the symmetry axis $r = 0$ , all infinite wave-packet corrections vanish identically. The blowup along $x = 0$ ( periodicVelocity_speed_unbounded , scaling as $(1-t)^{-A(h)}$ ) is generated purely by the unperturbed Stage-0 self-similar profile FinalSlowBase .
Exploding Harmonic Modes in the Active Annulus: Inside the active annulus $r / q \in [\text{leftRadius}, \text{rightRadius}]$ , FinalSlowBase produces a non-zero stress error ( stressForce ), which the wave packets iteratively cancel. In ActualCandidateConstruction.lean (Line 79), the harmonic mode band doubles at each cycle $j$ :
theorem cycle_residualBand (B N0 j : ℕ) :
(cycle B N0 j).coefficients.residualBand = 2 ^ (j + 1 )
Differentiating the harmonic phases $\exp(i k \Phi(z))$ up to order $m$ for $\vert{}k\vert{} \le 2^{j+1}$ introduces factors of $\vert{}k\vert{}^m \le 2^{(j+1)m}$ . Summing over these modes in GluedStageEstimates.lean (Line 192, localPotential_bound_of_modes ) causes the stage derivative constants $C_{j,m}$ ( potentialConstant CA j m ) to grow exponentially with $j$ and $m$ .
Post-Hoc Diagonal Cutoff Schedule ( exists_schedule ): To prevent $C_{j,m}$ from breaking convergence as $j \to \infty$ , MixedCandidateAssembly.lean (Lines 36–46) invokes exists_schedule after the sequence of constants $C_{j,m}$ is defined, selecting a rapidly growing sequence $a_j \to \infty$ so that the similarity factor $q(t,x)^{\text{gain}(j)}$ absorbs $C_{j,m}$ for $m \le j+2$ on the shrinking sublevel set ${q(t,x) &lt; 1/a_j}$ .
Part II: Audit of ./Euler/ (3D Euler Parameter Hierarchy)
1. Discrepancy Between the Preprint's Explicit Constants and Lean's ActualBounds
In Section 5.2 and 5.7 of the LaTeX preprint (and PacketInductionStage.lean , Line 110), the base core polynomial degree is hardcoded at $D_0 = 1000$ :
theorem core_cap : P.low.B ≤ gradientConstant * S.X ^ 1000 + 2
In PacketSourceScaleSequence.lean (Line 30) and PacketSourceScaleChoice.lean (Line 147), the previous shear at Stage $J$ ( $n = 0$ ) is initialized at $h_{J-1} = X^{1000}$ , and the positive pressure Hessian increment ( goodCost ) is:
$$\text{dH}_J = \delta_J \cdot h_J \cdot h_{J-1} = \exp\left(-\frac{X}{J^3}\right) \cdot \exp\left(\frac{X}{J^5}\right) \cdot X^{1000}$$
Simultaneously, ActualBounds in PacketSourceScaleActual.lean (Line 217) and NormalPacketFrequencyGuards.lean (Line 68) require at Stage $J$ :
structure ActualBounds (J D : ℕ) (C c X θ : ℝ) : Prop where
normal : UniformBounds J C c 60 (scaleSeque

[truncated]

## Original Extract

A structural audit demonstrating the infinite energy requirement hidden within OpenAI's forced Navier-Stokes Lean 4 formalization. - openai-navier-stokes-critique/README.md at main · bp-purana52127/openai-navier-stokes-critique

openai-navier-stokes-critique/README.md at main · bp-purana52127/openai-navier-stokes-critique · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
bp-purana52127
Copy path Blame More file actions Blame More file actions Latest commit
History History 169 lines (123 loc) · 16.8 KB main Breadcrumbs
Copy path Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions Structural & Mathematical Audit: openai/NavierStokesAndEuler
This repository contains an architectural, mathematical, and code-level verification audit of the openai/NavierStokesAndEuler Lean 4 formalization and its accompanying LaTeX preprint.
By tracing the dependency graph down to its leaf entry points and inspecting the core definitions across both ./NavierStokes/ (Forced 3D Navier-Stokes) and ./Euler/ (Unforced 3D Euler), this audit documents the exact architectural mechanisms used in the formalization—specifically manuscript-to-codebase specification gaps , conditional structure hypotheses ( PhysicalFields , StageEstimates ) , annular zero-germ cutoffs , post-hoc diagonal cutoff schedules , and a $10^{25}$ parameter scale divergence between the LaTeX preprint's explicit constants and Lean's non-constructive atTop filter.
Executive Summary of Audit Findings
Audit Area
Target Directory & Primary Files
Core Technical & Verification Finding
1. Documented Manuscript vs. Codebase Gaps
./NavierStokes/
GraphRestriction.lean
CycleContinuationInvariant.lean
DiagonalJetBoundsNoOptions.lean
Multiple modules explicitly document in their headers that they do not supply the manuscript's unconditional all-order estimates. Notably, GraphRestriction.lean proves a counterexample ( zero_mean_does_not_force_zero ) showing that zero auxiliary mean on the torus does not force zero after restricting to the graph.
2. External Realization Hypotheses & Conditional Assembly
./NavierStokes/
ActualCycleResidualBounds.lean
GluedStageEstimates.lean
CandidateAssembly.lean
MixedCandidateAssembly.lean
In ActualCycleResidualBounds.lean and GluedStageEstimates.lean , the physical realization of the stage velocity and pressure fields is passed as an external structure hypothesis ( PhysicalFields / PhysicalData : "Only the actual physical realization remains external" ).
3. Annular Zero-Germ Cutoffs & Residual Forcing Architecture
./NavierStokes/
ComparatorR3Theorem.lean
PeriodicResidualLimits.lean
MixedAxisPreservation.lean
ActualCandidateAssembly.lean
Clay Statement (C) ( option_C_of_compact_candidate ) is formalized with $u_0 \equiv 0$ by defining the external force $f(t,x)$ on $[0, 1)$ as the PDE residual $\partial_t u + (u\cdot\nabla)u - \nu\Delta u + \nabla p$ . All infinite wave-packet perturbations have a zero germ on the symmetry axis $r = 0$ ( mixedDiagonal_eq_base_germ ), isolating the blowup at $(t=1, x=0)$ entirely to the Stage-0 self-similar profile FinalSlowBase .
4. Exponential Mode Growth & Post-Hoc Cutoff Schedules
./NavierStokes/
ActualCandidateConstruction.lean
GluedStageEstimates.lean
MixedCandidateAssembly.lean
At cycle $j$ , active harmonic modes scale up to $\Vert{}k\Vert{} \le 2^{j+1}$ ( cycle_residualBand ), causing stage derivative constants $C_{j,m}$ to grow exponentially with $j$ and $m$ . To establish VanishingJointJets at $(1
[truncated]
1. Self-Documented Gaps Between the Manuscript and the Lean Code
A direct audit of the module docstrings and theorem statements in ./NavierStokes/ reveals several places where the formalization explicitly diverges from or weakens the unconditional claims of the written manuscript:
GraphRestriction.lean (Lines 16–18):
"None of these statements supplies the manuscript's claimed all-order residual estimates."
Furthermore, GraphRestriction.lean formalizes an explicit counterexample ( zero_mean_does_not_force_zero ), demonstrating that a zero auxiliary mean on the torus does not force a zero value after restriction to the graph.
CycleContinuationInvariant.lean (Lines 12–17):
"This file does not infer that input from smooth incoming primitives alone: an inverse covariance factor may grow at a flat annulus edge, so such an inference would require additional weighted estimates. In particular, the declarations below are a reconstruction algebra, not an unconditional iteration theorem or an assumed physical endpoint extension."
DiagonalJetBoundsNoOptions.lean (Lines 9–12):
"Pointwise stage estimates then give quantitative tail estimates. The stage estimates themselves are explicit hypotheses, not conclusions of the numerical cutoff selection."
CandidateAssembly.lean (Lines 10–14, 214):
"The theorem proves all CandidateProperties for these explicit outputs. It does not construct the input fields or the future force... A conditional implication, not a proof that its analytic inputs exist."
2. External PhysicalFields Hypotheses in Stage Residual Bounds
To prove Clay Statement (C) via the method of manufactured solutions ( option_C_of_compact_candidate in ComparatorR3Theorem.lean , Lines 22–35), the formalization sets the initial velocity to identically zero ( $u_0 \equiv 0$ ) and defines the external body force $f(t,x)$ on $t \in [0, 1)$ as the Navier-Stokes residual of the candidate fields $(u, p)$ ( assembledForce in CandidateAssembly.lean , Lines 20–67; periodicResidual in PeriodicResidualLimits.lean , Lines 23–35):
$$f(t,x) := \partial_t u + (u \cdot \nabla)u - \nu \Delta u + \nabla p$$
Under Fefferman's Statement (C) formulation, defining $f$ as the PDE residual of a blowing-up field $u$ is valid if and only if one proves that all spacetime derivatives of the residual cancel to infinite order at the singularity $(t=1, x=0)$ ( VanishingJointJets ), allowing $f$ to extend smoothly ( $C^\infty$ ) across $t = 1$ via Borel/Whitney extension ( CandidateFromLimits.force in MixedPeriodicAssembly.lean , Line 268).
In MixedCandidateAssembly.lean (Lines 52–60), VanishingJointJets depends entirely on supplying a witness for StageEstimates.finite_residual :
finite_residual : ∀ J m,
JetRate (𝓝[SpacetimeEndpoint.openPast 1 ] ( 1 , ( 0 : Space))) (PhysicalWaveSum.physicalQ h)
(fun z => navierStokesResidual (MixedDiagonalResidual.uncutVelocity A B J)
(DiagonalJetBounds.uncutPrefix P (J + 1 )) z. 1 z. 2 ) m (gain J - residualLoss m)
Tracing finite_residual to its provider in GluedStageEstimates.lean (Line 298) and ActualCycleResidualBounds.lean (Line 473) reveals that actualStageEstimates is parameterized by an explicit external realization input d :
(d : ∀ J, ActualCycleResidualBounds.PhysicalData B (N + 1 )
(ActualCyclePreservation.state B N0 J).state
(MixedDiagonalResidual.uncutVelocity A Bdirect J)
(DiagonalJetBounds.uncutPrefix P (J + 1 )))
And in ActualCycleResidualBounds.lean (Line 473), PhysicalFields (which underlies PhysicalData ) is documented as:
/-- Only the actual physical realization remains external. Native regularity,
the fixed base equation, and all size estimates are derived in this module. -/
structure PhysicalFields (B N : ℕ) (U : Set Cylinder) (s : StatePoint)
(u : VelocityField) (P : PressureField) : Prop where
velocity_germ : ∀ n, N ≤ n → ∀ z ∈ preterminal,
z ∈ PhysicalResidualTZ.graphSourceTZ (actualBandGraph n) U → ...
Rather than unconditionally constructing the global physical primitives in Cartesian coordinates, the residual module assumes velocity_germ to identify the Cartesian fields with the cylindrical graph pullback ( polarGraph ) on each dyadic band.
3. Annular Zero-Germ Isolation & Post-Hoc Cutoff Absorption
Inspecting ActualCandidateAssembly.lean (Lines 331–381) and MixedAxisPreservation.lean (Lines 33, 115, 212) clarifies how the blowup at $(t=1, x=0)$ interacts with the infinite wave-packet series:
theorem axis_not_active {w : SpaceTime} (ht : w ∈ PhysicalWaveSum.preterminal)
(ha : PhysicalGraphBounds.radialProjection w = 0 ) :
w ∉ ActualPolarCoverage.active
theorem mixedDiagonal_eq_base_germ :
mixedDiagonal base p D scales =ᶠ[𝓝 w] SpatialCurl.spatialCurl base
Zero Germ on the Blowup Axis ( $r = 0$ ): Every wave-packet perturbation stage ( particularPotential , signedPotential , directStages ) is supported strictly inside an active radial annulus $r / q \in [\text{leftRadius}, \text{rightRadius}]$ with $\text{leftRadius} &gt; 0$ ( PrimaryTargetBounds.leftRadius_pos ). Consequently, in a neighborhood of the symmetry axis $r = 0$ , all infinite wave-packet corrections vanish identically. The blowup along $x = 0$ ( periodicVelocity_speed_unbounded , scaling as $(1-t)^{-A(h)}$ ) is generated purely by the unperturbed Stage-0 self-similar profile FinalSlowBase .
Exploding Harmonic Modes in the Active Annulus: Inside the active annulus $r / q \in [\text{leftRadius}, \text{rightRadius}]$ , FinalSlowBase produces a non-zero stress error ( stressForce ), which the wave packets iteratively cancel. In ActualCandidateConstruction.lean (Line 79), the harmonic mode band doubles at each cycle $j$ :
theorem cycle_residualBand (B N0 j : ℕ) :
(cycle B N0 j).coefficients.residualBand = 2 ^ (j + 1 )
Differentiating the harmonic phases $\exp(i k \Phi(z))$ up to order $m$ for $\vert{}k\vert{} \le 2^{j+1}$ introduces factors of $\vert{}k\vert{}^m \le 2^{(j+1)m}$ . Summing over these modes in GluedStageEstimates.lean (Line 192, localPotential_bound_of_modes ) causes the stage derivative constants $C_{j,m}$ ( potentialConstant CA j m ) to grow exponentially with $j$ and $m$ .
Post-Hoc Diagonal Cutoff Schedule ( exists_schedule ): To prevent $C_{j,m}$ from breaking convergence as $j \to \infty$ , MixedCandidateAssembly.lean (Lines 36–46) invokes exists_schedule after the sequence of constants $C_{j,m}$ is defined, selecting a rapidly growing sequence $a_j \to \infty$ so that the similarity factor $q(t,x)^{\text{gain}(j)}$ absorbs $C_{j,m}$ for $m \le j+2$ on the shrinking sublevel set ${q(t,x) &lt; 1/a_j}$ .
Part II: Audit of ./Euler/ (3D Euler Parameter Hierarchy)
1. Discrepancy Between the Preprint's Explicit Constants and Lean's ActualBounds
In Section 5.2 and 5.7 of the LaTeX preprint (and PacketInductionStage.lean , Line 110), the base core polynomial degree is hardcoded at $D_0 = 1000$ :
theorem core_cap : P.low.B ≤ gradientConstant * S.X ^ 1000 + 2
In PacketSourceScaleSequence.lean (Line 30) and PacketSourceScaleChoice.lean (Line 147), the previous shear at Stage $J$ ( $n = 0$ ) is initialized at $h_{J-1} = X^{1000}$ , and the positive pressure Hessian increment ( goodCost ) is:
$$\text{dH}_J = \delta_J \cdot h_J \cdot h_{J-1} = \exp\left(-\frac{X}{J^3}\right) \cdot \exp\left(\frac{X}{J^5}\right) \cdot X^{1000}$$
Simultaneously, ActualBounds in PacketSourceScaleActual.lean (Line 217) and NormalPacketFrequencyGuards.lean (Line 68) require at Stage $J$ :
structure ActualBounds (J D : ℕ) (C c X θ : ℝ) : Prop where
normal : UniformBounds J C c 60 (scaleSeque

[truncated]
