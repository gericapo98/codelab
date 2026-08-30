### in a Julia cell using Pluto
using PlutoUI
# bind φ to a slider from 0 to 2π, starting at π/4, and show its value
###############################################################################
@bind φ Slider(0:π/16:2π, π/4, true)
###############################################################################
begin
    using Plots
    v = (cos(φ), sin(φ))
    theta = range(0, 2π, length=200)
    circle_x, circle_y = cos.(theta), sin.(theta)
    plt = plot(circle_x, circle_y;
               aspect_ratio=1,
               legend=false,
               linestyle=:dash,
               xlabel="x₁", ylabel="x₂",
               title="φ = $(round(φ, digits=2))")
    plot!(plt, [0, v[1]], [0, v[2]]; lw=3, label="k=1")
    annotate!(plt, v[1], v[2], text("($(round(v[1], digits=2)), $(round(v[2], digits=2)))", 8))
    plt
end

@bind psi Slider(0:π/16:2π, π/4, true)
@bind k1 Slider(1:1:10, 1, true)
begin 
	theta = range(0, 2π, length=200)
    w = (cos(k1 * psi), sin(k1 * psi))
    plt1 = plot(
        cos.(theta), sin.(theta);
        aspect_ratio = 1,
        legend = false,
        linestyle = :dash,
        xlabel = "x1",
        ylabel = "x2",               # ← not ylabel1
        title  = "ψ = $(round(psi, digits=2))"
    )
    plot!(plt1, [0, w[1]], [0, w[2]]; lw = 3, label = "k1 = 1")
    annotate!(
      plt1,
      w[1], w[2],
      text("($(round(v[1], digits=2)), $(round(w[2], digits=2)))", 8)
    )
    plt1
end
