package com.kakaocorp.buy.illuminati.jenkins.helper;

import com.jcraft.jsch.ChannelExec;
import com.jcraft.jsch.JSch;
import com.jcraft.jsch.Session;

import com.kakaocorp.buy.illuminati.jenkins.model.UserInfo;
import java.io.BufferedReader;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.PrintStream;
import java.util.ArrayList;
import java.util.List;

/**
 * Created by kellin.me@kakaocorp.com on 27/06/2017.
 */
public class SSHConnection {

    private final String userId;
    private final int port;
    private final String hostName;
    private String prvkey;
    private String password;

    public SSHConnection(String hostName, int port, String userId) {
        this.hostName = hostName;
        this.port = port;
        this.userId = userId;
    }

    public void setUserInfo(UserInfo userInfo) {
        this.password = userInfo.getPassword();
    }

    public void setPrvkey(String prvkey) {
        this.prvkey = prvkey;
    }

    private Session getConnection() throws Exception {
        JSch jSch = new JSch();

        if (this.password == null && this.prvkey != null) {
            jSch.addIdentity(this.prvkey);
        }

        Session session = jSch.getSession(this.userId, this.hostName, this.port);
        session.setConfig("StrictHostKeyChecking", "no");

        if (this.password != null) {
            session.setPassword(this.password);
        }

        session.connect();

        if (session.isConnected()) {
            return session;
        }

        return null;
    }

    public boolean checkConnection() {
        try {
            if (getConnection() != null && getConnection().isConnected()) {
                return true;
            }
        } catch (Exception ex) {
            return false;
        }

        return false;
    }

    public List<String> executeFile(String scriptFileName, PrintStream printStream) {
        List<String> result = new ArrayList<String>();
        try {
            printStream.println("execute shell script : " + scriptFileName);

            Session session = getConnection();

            ChannelExec channelExec = (ChannelExec)session.openChannel("exec");

            InputStream in = channelExec.getInputStream();
            channelExec.setCommand("sh "+scriptFileName);
            channelExec.connect();

            BufferedReader reader = new BufferedReader(new InputStreamReader(in));
            String line;

            while ((line = reader.readLine()) != null) {
                result.add(line);
            }

            int exitStatus = channelExec.getExitStatus();

            channelExec.disconnect();
            session.disconnect();

            if(exitStatus < 0) {
                // printStream.printl("Done, but exit status not set!");
            } else if(exitStatus > 0) {
                // printStream.printl("Done, but with error!");
            } else{
                // printStream.printl("Done!");
            }

        } catch(Exception e) {
            printStream.println("Error: " + e);
        }

        return result;
    }
}
