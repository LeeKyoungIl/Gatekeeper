package com.kakaocorp.buy.illuminati.jenkins.model;

import org.kohsuke.stapler.DataBoundConstructor;

/**
 * Created by kellin on 29/06/2017.
 */
public class UserInfo {

    private String id;
    private String password;

    @DataBoundConstructor
    public UserInfo(String id, String password) {
        this.id = id;
        this.password = password;
    }

    public String getId() {
        return this.id;
    }

    public String getPassword() {
        return this.password;
    }
}
