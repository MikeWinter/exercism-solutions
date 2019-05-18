class Year(private val year: Int) {
    val isLeap: Boolean get() = 4 divides year && (!(100 divides year) || 400 divides year)

    private infix fun Int.divides(numerator: Int): Boolean = numerator % this == 0
}
