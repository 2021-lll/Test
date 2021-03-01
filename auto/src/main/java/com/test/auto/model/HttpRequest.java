package com.test.auto.model;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.github.pagehelper.StringUtil;
import lombok.Data;

import java.io.Serializable;
import java.util.HashMap;

@Data
//@DateTimeFormat(pattern = "yyyy-MM-dd")
@JsonFormat(pattern = "yyyy-MM-dd", timezone = "GMT+8")
public class HttpRequestForm implements Serializable {

//    private static final long serialVersionUID = 1826152029135090793L;
    private int id;
    private String guid;
    private String pguid;
    private int sort;
    private String groupName; //分组名
    private String apiName;  //监控名称
    private String url; //监控地址
    private HttpMethod httpMethod; //请求方法
    private String headers; //请求头
    private String parameters; //请求参数
    private boolean archived;//是否删除（0-有效，1-删除）
    private MonitorFrequency frequency = MonitorFrequency.THIRTY;//监控频率
    //    private String resultType;
//    private String checkType;
//    private Date createTime;
//    private HashMap<String, String> headersMap = new HashMap<String, String>(); //请求头部，多个头部
//    private HashMap<String,String> parametersMap = new HashMap<String,String>();//请求参数
//    private HashMap<String,String> variablesMap = new HashMap<String,String>();//变量

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public String getGuid() {
        return guid;
    }

    public void setGuid(String guid) {
        this.guid = guid;
    }

    public String getPguid() {
        return pguid;
    }

    public void setPguid(String pguid) {
        this.pguid = pguid;
    }

    public int getSort() {
        return sort;
    }

    public void setSort(int sort) {
        this.sort = sort;
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

    public String getUrl() {
        return url;
    }

    public void setUrl(String url) {
        this.url = url;
    }

    public HttpMethod getHttpMethod() {
        return httpMethod;
    }

    public void setHttpMethod(HttpMethod httpMethod) {
        this.httpMethod = httpMethod;
    }

    public String getHeaders() {
        return headers;
    }

    public void setHeaders(String headers) {
        this.headers = headers;
    }

    public String getParameters() {
        return parameters;
    }

    public void setParameters(String parameters) {
        this.parameters = parameters;
    }

    public boolean isArchived() {
        return archived;
    }

    public void setArchived(boolean archived) {
        this.archived = archived;
    }

    public String getFrequency() {
        return String.valueOf(frequency);
    }

    public void setFrequency(MonitorFrequency frequency) {
        this.frequency = frequency;
    }

    public enum HttpMethod{
        GET, HEAD, POST, PUT, DELETE
    }

    public String mapToString(HashMap<String, String> map){
        if(map == null || map.size() == 0)return null;
        StringBuffer sb = new StringBuffer();
        for (String key : map.keySet()) {
            if(sb.length()!=0)sb.append("\r\n");
            sb.append(key).append(":").append(map.get(key));
        }
        return sb.toString();
    }

    public HashMap<String, String> stringToMap(String text) {
        if (StringUtil.isEmpty(text)) return null;
        HashMap<String, String> map = new HashMap<String, String>();
        String[] strArray = text.split("\r\n");
        for (String str : strArray) {
            String[] header = str.split("::");
            map.put(header[0], header[1]);
        }
        return map;
    }
}
