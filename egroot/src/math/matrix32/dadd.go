package matrix32

// Add performs: d = (a + b) * s
func (d Dense) Add(a, b Dense, s float32) {
	d.checkDim(a)
	d.checkDim(b)
	switch s {
	case 1:
		for i := 0; i < d.numrow; i++ {
			dr := d.v[i*d.stride:]
			ar := a.v[i*a.stride:]
			br := b.v[i*b.stride:]
			k := d.numcol - 1
			for k > 0 {
				dr[k] = ar[k] + br[k]
				k--
				dr[k] = ar[k] + br[k]
				k--
			}
			if k == 0 {
				dr[0] = ar[0] + br[0]
			}
		}
	default:
		for i := 0; i < d.numrow; i++ {
			dr := d.v[i*d.stride:]
			ar := a.v[i*a.stride:]
			br := b.v[i*b.stride:]
			k := d.numcol - 1
			for k > 0 {
				dr[k] = (ar[k] + br[k]) * s
				k--
				dr[k] = (ar[k] + br[k]) * s
				k--
			}
			if k == 0 {
				dr[0] = (ar[0] + br[0]) * s
			}
		}
	}
}
