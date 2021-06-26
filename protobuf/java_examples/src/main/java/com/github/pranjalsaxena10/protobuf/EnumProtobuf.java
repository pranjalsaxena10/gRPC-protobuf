package com.github.pranjalsaxena10.protobuf;

import example.enumerations.EnumExample;

public class EnumProtobuf {

    public static void main(String[] args) {
        EnumExample.EnumMessage.Builder enumBuilder = EnumExample.EnumMessage.newBuilder();
        enumBuilder.setId(1234);
        enumBuilder.setDayOfTheWeek(EnumExample.DayOfTheWeek.MONDAY);
        System.out.println(enumBuilder.build());
    }
}
