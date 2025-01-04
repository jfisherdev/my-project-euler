package com.jfisherdev.pejava;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class Problem8Test {

    private final Problem8 problem8 = new Problem8();

    @Test
    void solveProblem8Spec() {
        final long expected = 5832;
        final long actual = problem8.findLargestProduct(4);
        assertEquals(expected, actual);
    }

    @Test
    void solveProblem8() {
        final long expected = 23514624000L;
        final long actual = problem8.solveProblem8();
        assertEquals(expected, actual);
    }

}