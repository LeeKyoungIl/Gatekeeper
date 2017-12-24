package com.leekyoungil.gatekeeper.jenkins.helper;

import java.io.*;
import java.nio.charset.Charset;

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
            InputStreamReader inputStreamReader = new InputStreamReader(inputStream, Charset.defaultCharset());
            BufferedReader bufferedReader = new BufferedReader(inputStreamReader);

            String line = "";
            while ((line = bufferedReader.readLine()) != null) {
                this.printStream.println(line);
            }

            bufferedReader.close();
            inputStreamReader.close();
            inputStream.close();

            proc.waitFor();
        } catch (IOException ioe) {
            this.printStream.println("Command execution failed : " + command + ", message : " + ioe.getMessage());
        } catch (InterruptedException ie) {
            this.printStream.println("Command execution failed : " + command + ", message : " + ie.getMessage());
        }
    }
}
