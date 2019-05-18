fun transcribeToRna(dna: String): String = buildString {
    dna.map(::toRna).forEach { append(it) }
}

private fun toRna(nucleotide: Char): Char = when (nucleotide) {
    'G' -> 'C'
    'C' -> 'G'
    'T' -> 'A'
    'A' -> 'U'
    else -> throw IllegalArgumentException("invalid nucleotide: $nucleotide")
}
