package com.leekyoungil.gatekeeper.jenkins;

import com.leekyoungil.gatekeeper.jenkins.helper.SSHConnection;
import com.leekyoungil.gatekeeper.jenkins.helper.ShellExecutor;
import com.leekyoungil.gatekeeper.jenkins.model.UserInfo;
import hudson.Extension;
import hudson.FilePath;
import hudson.Launcher;
import hudson.model.AbstractProject;
import hudson.model.Run;
import hudson.model.TaskListener;
import hudson.tasks.BuildStepDescriptor;
import hudson.tasks.Builder;
import hudson.util.FormValidation;
import java.io.IOException;
import java.util.List;
import javax.servlet.ServletException;
import jenkins.tasks.SimpleBuildStep;
import net.sf.json.JSONObject;
import org.kohsuke.stapler.DataBoundConstructor;
import org.kohsuke.stapler.QueryParameter;
import org.kohsuke.stapler.StaplerRequest;

public class GatekeeperExecutor extends Builder implements SimpleBuildStep {

    private final String targetServer;
    private final String logServerHost;
    private final String excludePattern;
    private final String logfileName;
    private final String userId;
    private final String privateSSHKey;
    private final UserInfo userInfo;

    @DataBoundConstructor
    public GatekeeperExecutor(String targetServer, String logServerHost, String excludePattern, String logfileName, String userId, String privateSSHKey, UserInfo userInfo) {
        this.targetServer = targetServer;
        this.logServerHost = logServerHost;
        this.logfileName = logfileName;
        this.userId = userId;
        this.privateSSHKey = privateSSHKey;
        this.userInfo = userInfo;

        if (excludePattern == null) {
            this.excludePattern = "";
        } else {
            this.excludePattern = excludePattern;
        }
    }

    public String getUserId() {
        return this.userId;
    }
    public String getPrivateSSHKey() {
        return this.privateSSHKey;
    }
    public String getTargetServer() {
        return this.targetServer;
    }
    public String getLogServerHost() {
        return this.logServerHost;
    }
    public String getExcludePattern() {
        return this.excludePattern;
    }
    public String getLogfileName() {
        return this.logfileName;
    }
    public UserInfo getUserInfo() {
        return this.userInfo;
    }

    @Override
    public void perform(Run<?,?> build, FilePath workspace, Launcher launcher, TaskListener listener) {
        SSHConnection sshConnection = new SSHConnection(this.logServerHost, 22, this.userId);

        if (this.privateSSHKey != null) {
            sshConnection.setPrvkey(this.privateSSHKey);
        } else if (this.userInfo != null && this.getUserInfo().getPassword() != null) {
            sshConnection.setUserInfo(this.userInfo);
        } else {
            listener.getLogger().println("Please enter user info or ssh-key.");
            return;
        }

        boolean isConnected = sshConnection.checkConnection();

        if (isConnected) {
            ShellExecutor shellExecutor = new ShellExecutor(listener.getLogger());

            shellExecutor.execute("rsync -chavzP --stats "+this.userId+"@"+this.logServerHost+":/home/"+this.userId+"/gatekeeper/"+logfileName+" /home/"+this.userId+"/gatekeeper");
            shellExecutor.execute("/home/"+this.userId+"/gatekeeper/gatekeeperParseCore /home/"+this.userId+"/gatekeeper/ "+logfileName+" "+this.targetServer+" "+this.excludePattern);

            List<String> commandResult = sshConnection.executeFile("/home/"+this.userId+"/scripts/gatekeeperGetSampleCheck.sh", listener.getLogger());

            if (commandResult.size() == 0) {
                listener.getLogger().println("Failed to copy a log file");
            }
        } else {
            listener.getLogger().println("The log server host can not be contacted. Please check ssh-key or id and password.");
        }
    }

    @Override
    public DescriptorImpl getDescriptor() {
        return (DescriptorImpl)super.getDescriptor();
    }

    @Extension
    public static final class DescriptorImpl extends BuildStepDescriptor<Builder> {

        public DescriptorImpl() {
            load();
        }


        public FormValidation doCheckTargetServer(@QueryParameter String value)
                throws IOException, ServletException {
            if (value.length() == 0)
                return FormValidation.error("Please set a targetServer");
            return FormValidation.ok();
        }

        public FormValidation doCheckLogServerHost(@QueryParameter String value)
                throws IOException, ServletException {
            if (value.length() == 0)
                return FormValidation.error("Please set a logServerHost");
            return FormValidation.ok();
        }

        public FormValidation doCheckLogfileName(@QueryParameter String value)
                throws IOException, ServletException {
            if (value.length() == 0)
                return FormValidation.error("Please set a logfileName");
            return FormValidation.ok();
        }

        public FormValidation doCheckUserId(@QueryParameter String value)
                throws IOException, ServletException {
            if (value.length() == 0)
                return FormValidation.error("Please set a usserId");
            return FormValidation.ok();
        }

        public boolean isApplicable(Class<? extends AbstractProject> aClass) {
            return true;
        }

        public String getDisplayName() {
            return "GatekeeperExecutor";
        }

        @Override
        public boolean configure(StaplerRequest req, JSONObject formData) throws FormException {
            save();
            return super.configure(req,formData);
        }
    }
}

