package com.github.pranjalsaxena10.protobuf;

import example.simple.Simple.SimpleMessage;

import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;

public class SimpleProtobuf {
    public static void main(String[] args) {
        System.out.println("hello");
        SimpleMessage.Builder builder = SimpleMessage.newBuilder();
        builder.setText("hello protobuf");
        System.out.println(builder.toString());
        SimpleMessage message = builder.build();
        System.out.println("next:: " + message.getText());

        try {
            FileOutputStream outputStream = new FileOutputStream("simple_message.bin");
            message.writeTo(outputStream);
            outputStream.close();
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }

        try {
            System.out.println("reading binary file.....");
            FileInputStream inputStream = new FileInputStream("simple_message.bin");
            SimpleMessage simpleMessage = SimpleMessage.parseFrom(inputStream);
            System.out.println("Output from binary file: " + simpleMessage);
            inputStream.close();

        } catch (FileNotFoundException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
