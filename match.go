package matcher

import "github.com/filevich/combinatronics"

// dados dos arrays `xs` y `ys` retorna el primer elemento en comun que encuente
func lazy_intersection(xs, ys []int) (common int, compatiblbes bool) {
	for _, x := range xs {
		for _, y := range ys {
			if x == y {
				return x, true
			}
		}
	}
	return -1, false
}

type Matcher struct {
	params_len                     int
	option_prime_ix                []map[int]int          // (param_ix, option_id) -> option_prime_ix
	combos_prime_dot_compatibility map[int](map[int]bool) // (combo1_prime_DOT, combo2_prime_DOT) -> 0|1
}

func (m *Matcher) map_prime(param_ix int, opts []int) []int {
	res := make([]int, len(opts))
	for i, opt := range opts {
		res[i] = primes[m.option_prime_ix[param_ix][opt]]
	}
	return res
}

func (m *Matcher) to_prime_dots(config [][]int) []int {
	res := make([]int, len(config))
	for pix, opts := range config {
		param_opts_primes := m.map_prime(pix, opts)
		res[pix] = combinatronics.Dot(param_opts_primes)
	}

	return res
}

func (m *Matcher) opts_compatibles(opts1_prime_dot, opts2_prime_dot int) bool {
	if _, ok := m.combos_prime_dot_compatibility[opts1_prime_dot]; ok {
		if compatibles, ok2 := m.combos_prime_dot_compatibility[opts1_prime_dot][opts2_prime_dot]; ok2 {
			return compatibles
		}
	}

	return m.combos_prime_dot_compatibility[opts2_prime_dot][opts1_prime_dot]
}

func (m *Matcher) Prime_dots_compatibles(a_dots, b_dots []int) bool {
	for i := 0; i < m.params_len; i++ {
		if !m.opts_compatibles(a_dots[i], b_dots[i]) {
			return false
		}
	}

	return true
}

// // retorna true si a y b son configuraciones compatibles
func (m *Matcher) Compatibles(a, b [][]int) bool {
	a_dots := m.to_prime_dots(a)
	b_dots := m.to_prime_dots(b)

	return m.Prime_dots_compatibles(a_dots, b_dots)
}

// c = <p1,p2,p3> where
// p1 \subset {o11,o12,o13}, p2 \subset {o21,o22,o23} ...

// ejemplo: si los parametros son: <PUNTOS,CANT_JUGADORES>
// y los dominios son {A20, A30, A40} y {1vs1, 2vs2, 3vs3}
// entonces una configuracion es <{A20, A40}, {2vs2}>

// entonces params es [][]{{20,30,40},{2,4,6}}

func NewMatcher(params [][]int) *Matcher {

	m := Matcher{
		params_len:                     len(params),
		option_prime_ix:                make([]map[int]int, len(params)),
		combos_prime_dot_compatibility: make(map[int](map[int]bool)),
	}

	combinatronics.Dot([]int{1, 2, 3})

	max_param := 0
	for _, param := range params {
		if len(param) > max_param {
			max_param = len(param)
		}
	}

	if max_param > len(primes) {
		panic("number of parameter options exceded")
	}

	// ya se que no voy a necesitar mas de `max_param` primos
	primes = primes[:max_param]

	// luego, armo la tabla en que, a cada OPCION le asocio un indice de primo
	// esto lo hago porque el valor/dominio de las opciones podria ser: 100,200,300
	// a ellos los mapeo con los indices 0,1,2
	// luego esos indices los mapeo con los primos 2,3,5
	// option_prime_ix : PARAM_IX * OPTION -> PRIME_IX

	for pix, param := range params {
		m.option_prime_ix[pix] = make(map[int]int)
		for i, option := range param {
			m.option_prime_ix[pix][option] = i
		}
	}

	// ahora creo la tabla de option_prime_ix -> primo
	primes = primes[:max_param]

	// todas las combinaciones INDICES de opciones
	all_possible_combinations := combinatronics.Subsets(primes, len(primes))

	for i, combo1_primes := range all_possible_combinations {
		combo1_primes_DOT := combinatronics.Dot(combo1_primes)
		m.combos_prime_dot_compatibility[combo1_primes_DOT] = make(map[int]bool)

		for _, combo2_primes := range all_possible_combinations[i:] {
			_, compatibles := lazy_intersection(combo1_primes, combo2_primes)
			combo2_primes_DOT := combinatronics.Dot(combo2_primes)
			m.combos_prime_dot_compatibility[combo1_primes_DOT][combo2_primes_DOT] = compatibles
		}
	}

	return &m
}
