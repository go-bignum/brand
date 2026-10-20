# go-bignum — brand assets

Official logos for the **go-bignum** organization.

<img src="avatar/go-bignum.png" width="160" alt="go-bignum">

The mark is a very long number: a `1`, a `0`, and an ellipsis, because the
number does not stop. It is written by `mark/main.go`, which lays the digits and
the three dots on one baseline and centres them; nothing is placed by hand. Its
gradient (`#E27AB8` to `#A02A74`) is a magenta that sets go-bignum apart from
the violet of go-numerics, the emerald of go-ndarray and go-fft and the blue of
go-plot; its dark end reads on white at 6.82:1 and its light end on a dark page
at 7.09:1.

## Formats

- **SVG** — `svg/color`, `svg/white` (white, transparent), `svg/black` (black, transparent)
- **PNG** 16→1024 px, 3 variants — `png/<variant>/<size>/`
- **JPG** (color, white background) — `jpg/<size>/`
- **ICO** (Windows) — `ico/`
- **ICNS** (macOS) — `icns/`
- **GitHub avatar** 512 px — `avatar/`
- **Social preview** 1280×640 (repo banner) — `social/`

`check/kit.py` counts the whole matrix on every pull request, and the `mark`
job runs `mark/main.go` again and fails if `svg/color/go-bignum.svg` differs.

---
*Generated from `svg/color/go-bignum.svg` by the fleet's Go brand pipeline (brandgen) — rounded square badge, line glyph, color / white / black variants.*
