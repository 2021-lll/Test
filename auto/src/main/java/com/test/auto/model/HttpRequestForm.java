package com.test.auto.model;

import lombok.Data;
import org.apache.ibatis.annotations.ResultType;
import org.springframework.http.HttpMethod;


import java.util.Date;

@Data
public class HttpRequestForm {

    private int id;
    private String groupName;
    private String apiName;
    private String urlAddress;

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public String getGroupName() {
        return groupName;
    }

    public void setGroupName(String groupName) {
        this.groupName = groupName;
    }

    public String getApiName() {
        return apiName;
    }

    public void setApiName(String apiName) {
        this.apiName = apiName;
    }

    public String getUrlAddress() {
        return urlAddress;
    }

    public void setUrlAddress(String urlAddress) {
        this.urlAddress = urlAddress;
    }
}
