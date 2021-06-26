package com.github.pranjalsaxena10.protobuf;

import example.complex.Complex;

import java.util.Arrays;

public class ComplexProtobuf {
    public static void main(String[] args) {
        System.out.println("Complex Protobuf");
        Complex.DummyMessage.Builder dummyMessageBuilder = Complex.DummyMessage.newBuilder();
        dummyMessageBuilder.setId(123);
        dummyMessageBuilder.setName("dummyMessage");
        Complex.ComplexMessage.Builder complexMessageBuilder = Complex.ComplexMessage.newBuilder();
        complexMessageBuilder.setOneDummy(dummyMessageBuilder.build());
        complexMessageBuilder.addAllMultipleDummy(Arrays.asList(
                Complex.DummyMessage.newBuilder().setId(432).setName("mutliple2").build(),
                Complex.DummyMessage.newBuilder().setId(532).setName("multiple3").build()
        ));

        System.out.println(complexMessageBuilder.build());
        System.out.println(complexMessageBuilder.getMultipleDummyList());
    }
}
