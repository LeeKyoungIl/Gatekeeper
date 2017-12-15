package com.kakaocorp.buy.illuminati.jenkins.helper;

import java.io.BufferedReader;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.PrintStream;

/**
 * Created by kellin.me@kakaocorp.com on 27/06/2017.
 */
public class ShellExecutor {

    private PrintStream printStream;

    public ShellExecutor(PrintStream printStream) {
        this.printStream = printStream;
    }

    public void execute(String command) {
        try {
            Process proc = Runtime.getRuntime().exec(command);
            InputStream inputStream = proc.getInputStream();
            InputStreamReader inputStreamReader = new InputStreamReader(inputStream);
            BufferedReader bufferedReader = new BufferedReader(inputStreamReader);

            String line = "";
            while ((line = bufferedReader.readLine()) != null) {
                this.printStream.println(line);
            }
            proc.waitFor();
        } catch (Exception e) {
            this.printStream.println("Command execution failed : "+command);
        }
    }
}
