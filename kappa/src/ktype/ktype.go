package ktype

type ref_t struct {
    ref string
    name string
    seq []int
}

type sam_t struct {
    ref string
    refs []ref_t
}

type sam_i struct {
    ref string
    info sam_t
    size []int
    name []string
}

type genome struct {
    ref string
    name string
    start int
    end int
    size int
}

type superg struct {
    ref string
    genes []genome
}

type kappa_t struct {
    ref string
    size int
}
