__attribute__ ((always_inline))
extern inline
uintptr syscall$b2u(bool b) {
	return (uintptr)(b);
}

__attribute__ ((always_inline))
extern inline
uintptr syscall$f2u(void (*f)()) {
	union {void (*in)(); uintptr out;} cast;
	cast.in = f;
	return cast.out;
}
